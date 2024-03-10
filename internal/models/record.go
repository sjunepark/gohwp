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

	for !br.IsEOF() {
		record, err := br.ReadRecord()
		if err != nil {
			return nil, err
		}

		parent := rootRecord

		for i := uint32(0); i < record.Level; i++ {
			parent = parent.Children[len(parent.Children)-1]
		}

		parent.Children = append(parent.Children, record)
	}

	return rootRecord, nil
}
