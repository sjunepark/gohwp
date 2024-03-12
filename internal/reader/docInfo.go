package reader

import (
	"github.com/sjunepark/gohwp/internal/constant"
	"github.com/sjunepark/gohwp/internal/model"
)

type DocInfoReader struct {
	record  model.Record
	docInfo *model.DocInfo
}

func NewDocInfoReader(data []byte) (*DocInfoReader, error) {
	record, err := model.ReadRecordTree(data)
	if err != nil {
		return nil, err
	}
	return &DocInfoReader{record: *record, docInfo: &model.DocInfo{}}, nil
}

func (p *DocInfoReader) Read() (*model.DocInfo, error) {
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
	err := br.ReadStruct(dp)
	if err != nil {
		return err
	}

	docInfo.DocumentProperties = dp
	return nil
}
