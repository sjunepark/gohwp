package parser

import (
	"context"
	"fmt"
	"github.com/sjunepark/gohwp/internal/constants"
	"github.com/sjunepark/gohwp/internal/models"
)

type SectionParser struct {
	reader  *models.Record
	section *models.Section
}

func NewSectionParser(data []byte) (*SectionParser, error) {
	record, err := models.ParseRecordTree(data)
	if err != nil {
		return nil, err
	}
	return &SectionParser{reader: record, section: &models.Section{}}, nil
}

func (p *SectionParser) Parse(ctx context.Context) (*models.Section, error) {
	for _, child := range p.reader.Children {
		err := visitSection(child, p.section, ctx)
		if err != nil {
			return nil, err
		}
	}
	return p.section, nil
}

func visitSection(record *models.Record, section *models.Section, ctx context.Context) error {
	fmt.Println("TagID:", record.TagID)
	switch record.TagID {
	case constants.SECTION_HWPTAG_PARA_HEADER:
		paraHeader, err := visitParHeader(record, section, ctx)
		if err != nil {
			return err
		}
		fmt.Println(paraHeader)
	case constants.SECTION_HWPTAG_PARA_TEXT:
		paraText, err := visitParaText(record, section)
		if err != nil {
			return err
		}
		// todo: remove
		fmt.Println(paraText.String())
	default:
		fmt.Println("Unimplemented tag:", record.TagID)
		return nil
	}

	for _, child := range record.Children {
		err := visitSection(child, section, ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func visitParHeader(record *models.Record, section *models.Section, ctx context.Context) (*models.ParaHeader, error) {
	br := models.ByteReader{Data: record.Payload}

	var paraHeader models.ParaHeader

	hwpVersion, ok := VersionFromContext(ctx)
	if !ok {
		return nil, fmt.Errorf("hwpVersion not found in context")
	}
	if (hwpVersion.Gte(models.HWPVersion{Major: 5, Build: 3, Revision: 2})) {
		var paraHeaderV2 models.ParaHeaderV2

		err := br.ReadStruct(&paraHeaderV2)
		if err != nil {
			return nil, err
		}
		paraHeader = paraHeaderV2
	} else {
		var paraHeaderV1 models.ParaHeaderV1

		err := br.ReadStruct(&paraHeaderV1)
		if err != nil {
			return nil, err
		}
		paraHeader = paraHeaderV1
	}

	return &paraHeader, nil
}

func visitParaText(record *models.Record, section *models.Section) (models.ParaText, error) {
	// todo: implement paragraph size
	br := models.ByteReader{Data: record.Payload}

	var paraText models.ParaText

	for !br.IsEOF() {
		var wChar models.WChar
		err := br.ReadStruct(&wChar)
		if err != nil {
			return nil, err
		}

		charType := wChar.CharType()
		if charType == models.CharTypeInline || charType == models.CharTypeExtended {
			err = br.Skip(14)
			if err != nil {
				return nil, err
			}
		}
		paraText = append(paraText, wChar)
	}

	return paraText, nil
}
