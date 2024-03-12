package parser

import (
	"github.com/sjunepark/gohwp/internal/constant"
	"github.com/sjunepark/gohwp/internal/model"
)

type DocInfoParser struct {
	record  model.Record
	docInfo *model.DocInfo
}

func NewDocInfoParser(data []byte) (*DocInfoParser, error) {
	record, err := model.ParseRecordTree(data)
	if err != nil {
		return nil, err
	}
	return &DocInfoParser{record: *record, docInfo: &model.DocInfo{}}, nil
}

func (p *DocInfoParser) Parse() (*model.DocInfo, error) {
	for _, child := range p.record.Children {
		err := visitDocInfo(child, p.docInfo)
		if err != nil {
			return nil, err
		}
	}
	return p.docInfo, nil
}

func visitDocInfo(record *model.Record, docInfo *model.DocInfo) error {
	switch record.TagID {
	case constant.HWPTAG_DOCUMENT_PROPERTIES:
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

func visitDocumentProperties(record *model.Record, docInfo *model.DocInfo) error {
	dp := &model.DocumentProperties{}
	br := model.ByteReader{Data: record.Payload}
	_, err := br.ReadStruct(dp)
	if err != nil {
		return err
	}

	docInfo.DocumentProperties = dp
	return nil
}
