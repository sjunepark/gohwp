package parser

import (
	"github.com/sjunepark/gohwp/internal/constants"
	"github.com/sjunepark/gohwp/internal/models"
)

type DocInfoParser struct {
	record  models.Record
	docInfo *models.DocInfo
}

func NewDocInfoParser(data []byte) (*DocInfoParser, error) {
	record, err := models.ParseRecordTree(data)
	if err != nil {
		return nil, err
	}
	return &DocInfoParser{record: *record, docInfo: &models.DocInfo{}}, nil
}

func (p *DocInfoParser) Parse() (*models.DocInfo, error) {
	for _, child := range p.record.Children {
		err := visitDocInfo(child, p.docInfo)
		if err != nil {
			return nil, err
		}
	}
	return p.docInfo, nil
}

func visitDocInfo(record *models.Record, docInfo *models.DocInfo) error {
	switch record.TagID {
	case constants.DOCINFO_HWPTAG_DOCUMENT_PROPERTIES:
		err := visitDocumentProperties(record, docInfo)
		if err != nil {
			return err
		}
	default:
		return nil
	}

	for _, child := range record.Children {
		err := visitDocInfo(child, docInfo)
		if err != nil {
			return err
		}
	}
	return nil
}

func visitDocumentProperties(record *models.Record, docInfo *models.DocInfo) error {
	dp := &models.DocumentProperties{}
	br := models.ByteReader{Data: record.Payload}
	_, err := br.ReadStruct(dp)
	if err != nil {
		return err
	}

	docInfo.DocumentProperties = dp
	return nil
}
