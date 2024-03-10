package section

import (
	"fmt"
	"github.com/sjunepark/gohwp/internal/constants"
	"github.com/sjunepark/gohwp/internal/models"
	"github.com/sjunepark/gohwp/internal/types"
)

type Parser struct {
	reader  *models.Record
	section *Section
}

func NewParser(data []byte) (*Parser, error) {
	record, err := models.ParseRecordTree(data)
	if err != nil {
		return nil, err
	}
	return &Parser{reader: record, section: &Section{}}, nil
}

func (p *Parser) Parse() (*Section, error) {
	for _, child := range p.reader.Children {
		err := visit(child, p.section)
		if err != nil {
			return nil, err
		}
	}
	return p.section, nil
}

func visit(record *models.Record, section *Section) error {
	switch record.TagID {
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
		err := visit(child, section)
		if err != nil {
			return err
		}
	}
	return nil
}

func visitParaText(record *models.Record, section *Section) (ParaText, error) {
	// todo: implement paragraph size
	br := models.ByteReader{Data: record.Payload}

	var paraText ParaText

	for !br.IsEOF() {
		var wChar types.WChar
		err := br.ReadStruct(&wChar)
		if err != nil {
			return nil, err
		}

		charType := wChar.GetCharType(wChar)
		if charType == types.CharTypeInline || charType == types.CharTypeExtended {
			err = br.Skip(14)
			if err != nil {
				return nil, err
			}
		}
		paraText = append(paraText, wChar)
	}

	return paraText, nil
}
