package parser

import "github.com/sjunepark/gohwp/internal/types"

type DocInfoParser struct {
	record  types.Record
	docInfo types.DocInfo
}

func NewDocInfoParser(data []byte) (*DocInfoParser, error) {
	record, err := types.ParseRecordTree(data)
	if err != nil {
		return nil, err
	}

	return &DocInfoParser{record: *record}, nil
}

func (d *DocInfoParser) Parse() (*types.DocInfo, error) {
	for _, child := range d.record.Children {
		err := visit(child)
		if err != nil {
			return nil, err
		}
	}
	return &d.docInfo, nil
}

func visit(record *types.Record) error {
	// todo: implement visit logic
	for _, child := range record.Children {
		err := visit(child)
		if err != nil {
			return err
		}
	}
	return nil
}
