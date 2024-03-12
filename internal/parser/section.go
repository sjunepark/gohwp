package parser

import (
	"context"
	"fmt"
	"github.com/sjunepark/gohwp/internal/constant"
	"github.com/sjunepark/gohwp/internal/model"
)

type SectionParser struct {
	record  *model.Record
	section *model.Section
}

func NewSectionParser(data []byte) (*SectionParser, error) {
	record, err := model.ParseRecordTree(data)
	if err != nil {
		return nil, err
	}
	return &SectionParser{record: record, section: &model.Section{}}, nil
}

func (p *SectionParser) Parse(ctx context.Context) (*model.Section, error) {
	for _, child := range p.record.Children {
		err := visitSection(child, p.section, ctx)
		if err != nil {
			return nil, err
		}
	}
	fmt.Println(p.section)
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

		_, err := br.ReadStruct(&paraHeaderV2)
		if err != nil {
			return err
		}
		paraHeader = paraHeaderV2
	} else {
		var paraHeaderV1 model.ParaHeaderV1

		_, err := br.ReadStruct(&paraHeaderV1)
		if err != nil {
			return err
		}
		// Q: Fix this
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
	textLength := currentPara.ParaHeader.TextLength // Outputs 40

	var paraText model.ParaText
	var processedBytes int
	const textBytes = 2

	for processedBytes <= int(textLength)*textBytes {
		var wChar model.WChar
		offset, err := br.ReadStruct(&wChar)
		if err != nil {
			return err
		}
		processedBytes += offset

		charType := wChar.CharType()
		if charType == model.CharTypeInline || charType == model.CharTypeExtended {
			err = br.Skip(14)
			if err != nil {
				return err
			}
		}
		paraText = append(paraText, wChar)
	}

	section.CurrentParagraph().ParaText = paraText
	return nil
}
