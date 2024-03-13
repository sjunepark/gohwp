package reader

import (
	"context"
	"fmt"
	"github.com/sjunepark/hwp/internal/reader/constant"
	"github.com/sjunepark/hwp/internal/reader/model"
)

type SectionReader struct {
	record  *model.Record
	section *model.Section
}

func NewSectionReader(data []byte) (*SectionReader, error) {
	record, err := model.ReadRecordTree(data)
	if err != nil {
		return nil, err
	}
	return &SectionReader{record: record, section: &model.Section{}}, nil
}

func (p *SectionReader) Read(ctx context.Context) (*model.Section, error) {
	for _, child := range p.record.Children {
		err := visitSection(child, p.section, ctx)
		if err != nil {
			return nil, err
		}
	}
	return p.section, nil
}

func visitSection(record *model.Record, section *model.Section, ctx context.Context) error {
	switch record.TagID {
	case constant.HWPTAG_PARA_HEADER:
		err := visitParHeader(record, section, ctx)
		if err != nil {
			return err
		}
	default:
		return nil
	}
	return nil
}

func visitParHeader(record *model.Record, section *model.Section, ctx context.Context) error {
	if record.Level != 0 {
		return fmt.Errorf("invalid level: %d", record.Level)
	}
	br := model.ByteReader{Data: record.Payload}

	var paraHeader model.ParaHeader

	hwpVersion, ok := getVersion(ctx)
	if !ok {
		return fmt.Errorf("hwpVersion not found in context")
	}
	if (hwpVersion.Gte(model.HWPVersion{Major: 5, Build: 3, Revision: 2})) {
		var paraHeaderV2 model.ParaHeader

		err := br.ReadStruct(&paraHeaderV2)
		if err != nil {
			return err
		}
		paraHeader = paraHeaderV2
	} else {
		var paraHeaderV1 model.ParaHeaderV1

		err := br.ReadStruct(&paraHeaderV1)
		if err != nil {
			return err
		}
		paraHeader = model.ParaHeader{ParaHeaderV1: paraHeaderV1, IsMergedTrack: nil}
	}

	// Initialize paragraph
	paragraph := model.Paragraph{}
	paragraph.ParaHeader = &paraHeader
	section.Paragraphs = append(section.Paragraphs, &paragraph)

	for _, child := range record.Children {
		err := visitParaElem(child, section, ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func visitParaElem(record *model.Record, section *model.Section, ctx context.Context) error {
	switch record.TagID {
	case constant.HWPTAG_PARA_TEXT:
		err := visitParaText(record, section)
		if err != nil {
			return err
		}
	}
	return nil
}

func visitParaText(record *model.Record, section *model.Section) error {
	if record.Level != 1 {
		return fmt.Errorf("invalid level: %d", record.Level)
	}
	br := model.ByteReader{Data: record.Payload} // Is size 80

	currentPara := section.CurrentParagraph()
	textLength := currentPara.ParaHeader.TextLength
	const textBytes = 2

	var paraText model.ParaText

	bytesToRead := int(textLength) * textBytes
	// WChar is uint16, which is 2 bytes
	wCharBytes := 2

	for bytesToRead > 0 {
		var wChar model.WChar
		err := br.ReadStruct(&wChar)
		bytesToRead -= wCharBytes
		if err != nil {
			return err
		}

		switch wChar.CharType() {
		case model.CharTypeChar:
			//	Nothing to skip
		case model.CharTypeInline, model.CharTypeExtended:
			// Inline and extended controls are 12 bytes long, with an additional equivalent ch at the end
			bytesToSkip := 12
			err := br.Skip(bytesToSkip)
			if err != nil {
				return err
			}
			bytesToRead -= bytesToSkip

			var wCharEnd model.WChar
			err = br.ReadStruct(&wCharEnd)
			bytesToRead -= wCharBytes
			if err != nil {
				return err
			}

			if wChar != wCharEnd {
				return fmt.Errorf("during parsing control characters, start and end characters do not match: %d, %d", wChar, wCharEnd)
			}
		default:
			//	Nothing to skip
		}

		paraText = append(paraText, wChar)
	}

	section.CurrentParagraph().ParaText = paraText
	return nil
}
