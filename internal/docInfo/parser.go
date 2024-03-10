package docInfo

import (
	"fmt"
	"github.com/sjunepark/gohwp/internal/constants"
	"github.com/sjunepark/gohwp/internal/models"
)

type Parser struct {
	record  models.Record
	docInfo *DocInfo
}

func NewParser(data []byte) (*Parser, error) {
	record, err := models.ParseRecordTree(data)
	if err != nil {
		return nil, err
	}

	return &Parser{record: *record, docInfo: &DocInfo{}}, nil
}

func (p *Parser) Parse() (*DocInfo, error) {
	for _, child := range p.record.Children {
		err := visit(child, p.docInfo)
		if err != nil {
			return nil, err
		}
	}
	return p.docInfo, nil
}

func visit(record *models.Record, docInfo *DocInfo) error {
	switch record.TagID {
	case constants.DOCINFO_HWPTAG_DOCUMENT_PROPERTIES:
		err := visitDocumentProperties(record, docInfo)
		if err != nil {
			return err
		}
	default:
		fmt.Println("Unimplemented tag:", record.TagID)
		return nil
	}

	for _, child := range record.Children {
		err := visit(child, docInfo)
		if err != nil {
			return err
		}
	}
	return nil
}

func visitDocumentProperties(record *models.Record, docInfo *DocInfo) error {
	dp := &DocumentProperties{}
	br := models.ByteReader{Data: record.Payload}
	err := br.ReadStruct(dp)
	if err != nil {
		return err
	}

	docInfo.DocumentProperties = dp
	return nil
}
