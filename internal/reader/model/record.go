package model

import (
	"fmt"
	"github.com/sjunepark/hwp/internal/reader/constant"
	"strings"
)

type Record struct {
	RecordHeader
	Payload  []byte
	Children []*Record
}

func (r *Record) String() string {
	var prettyPrint func(*Record, int) string
	prettyPrint = func(record *Record, indentLevel int) string {
		indent := strings.Repeat("  ", indentLevel) // Creates an indentation string
		var builder strings.Builder

		// Append current record's information
		builder.WriteString(fmt.Sprintf("%s%s\n", indent, record.RecordHeader.String()))

		// Recursively append children's information, with increased indentation
		for _, child := range record.Children {
			builder.WriteString(prettyPrint(child, indentLevel+1))
		}

		return builder.String()
	}

	// Start the pretty print from the current record with no indentation
	return prettyPrint(r, 0)
}

type RecordHeader struct {
	TagID constant.TagID `validate:"gte=0,lte=1023"` // TagID must be between 0 and 0x3FF (1023 in decimal)
	Level uint32         `validate:"gte=0,lte=1023"` // Level must be between 0 and 0x3FF
	Size  uint32         `validate:"gte=0,lte=4095"` // Size must be between 0 and 0xFFF (4095 in decimal)
}

func (rh *RecordHeader) String() string {
	return fmt.Sprintf("TagID: %s, Level: %d, Size: %d", rh.TagID, rh.Level, rh.Size)
}

func ReadRecordTree(data []byte) (*Record, error) {
	br := &ByteReader{Data: data}

	rootRecord := &Record{
		RecordHeader: RecordHeader{
			TagID: 0,
			Level: 0,
			Size:  0,
		},
		Payload:  []byte{},
		Children: []*Record{},
	}

	stack := []*Record{rootRecord}

	for !br.IsEOF() {
		record, err := br.ReadRecord()
		if err != nil {
			return nil, err
		}

		for (len(stack) > 1) && (stack[len(stack)-1].Level >= record.Level) {
			stack = stack[:len(stack)-1]
		}
		parent := stack[len(stack)-1]
		parent.Children = append(parent.Children, record)
		stack = append(stack, record)
	}

	return rootRecord, nil
}
