package models

type Record struct {
	RecordHeader
	Payload  []byte
	Children []*Record
}

type RecordHeader struct {
	TagID uint32 `validate:"gte=0,lte=1023"` // TagID must be between 0 and 0x3FF (1023 in decimal)
	Level uint32 `validate:"gte=0,lte=1023"` // Level must be between 0 and 0x3FF
	Size  uint32 `validate:"gte=0,lte=4095"` // Size must be between 0 and 0xFFF (4095 in decimal)
}

func ParseRecordTree(data []byte) (*Record, error) {
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

		// Pop from the stack until we find the correct parent level, and get the parent
		for len(stack) > 1 && stack[len(stack)-1].Level <= record.Level {
			stack = stack[:len(stack)-1]
		}
		parent := stack[len(stack)-1]

		parent.Children = append(parent.Children, record)
		stack = append(stack, record)
	}

	return rootRecord, nil
}
