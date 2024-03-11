package parser

import (
	"context"
	"fmt"
	"github.com/sjunepark/gohwp/internal/constants"
	"github.com/sjunepark/gohwp/internal/models"
)

type SectionParser struct {
	record  *models.Record
	section *models.Section
}

func NewSectionParser(data []byte) (*SectionParser, error) {
	record, err := models.ParseRecordTree(data)
	if err != nil {
		return nil, err
	}
	return &SectionParser{record: record, section: &models.Section{}}, nil
}

func (p *SectionParser) Parse(ctx context.Context) (*models.Section, error) {
	for _, child := range p.record.Children {
		err := visitSection(child, p.section, ctx)
		if err != nil {
			return nil, err
		}
	}
	fmt.Println(p.section)
	return p.section, nil
}

func visitSection(record *models.Record, section *models.Section, ctx context.Context) error {
	switch record.TagID {
	case constants.SECTION_HWPTAG_PARA_HEADER:
		err := visitParHeader(record, section, ctx)
		if err != nil {
			return err
		}
	default:
		return nil
	}

	return nil
}

func visitParHeader(record *models.Record, section *models.Section, ctx context.Context) error {
	if record.Level != 0 {
		return fmt.Errorf("invalid level: %d", record.Level)
	}
	br := models.ByteReader{Data: record.Payload}

	var paraHeader models.ParaHeader

	hwpVersion, ok := getVersion(ctx)
	if !ok {
		return fmt.Errorf("hwpVersion not found in context")
	}
	if (hwpVersion.Gte(models.HWPVersion{Major: 5, Build: 3, Revision: 2})) {
		var paraHeaderV2 models.ParaHeader

		_, err := br.ReadStruct(&paraHeaderV2)
		if err != nil {
			return err
		}
		paraHeader = paraHeaderV2
	} else {
		var paraHeaderV1 models.ParaHeaderV1

		_, err := br.ReadStruct(&paraHeaderV1)
		if err != nil {
			return err
		}
		// Q: Fix this
		paraHeader = models.ParaHeader{ParaHeaderV1: paraHeaderV1, IsMergedTrack: nil}
	}

	// Initialize paragraph
	paragraph := models.Paragraph{}
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

func visitParaElem(record *models.Record, section *models.Section, ctx context.Context) error {
	switch record.TagID {
	case constants.SECTION_HWPTAG_PARA_TEXT:
		err := visitParaText(record, section)
		if err != nil {
			return err
		}
	}
	return nil
}

func visitParaText(record *models.Record, section *models.Section) error {
	if record.Level != 1 {
		return fmt.Errorf("invalid level: %d", record.Level)
	}
	br := models.ByteReader{Data: record.Payload} // Is size 80

	currentPara := section.CurrentParagraph()
	textLength := currentPara.ParaHeader.TextLength // Outputs 40

	var paraText models.ParaText
	var processedBytes int
	const textBytes = 2

	for processedBytes <= int(textLength)*textBytes {
		var wChar models.WChar
		offset, err := br.ReadStruct(&wChar)
		if err != nil {
			return err
		}
		processedBytes += offset

		charType := wChar.CharType()
		if charType == models.CharTypeInline || charType == models.CharTypeExtended {
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
