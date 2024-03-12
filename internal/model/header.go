package model

import (
	"fmt"
	"github.com/sjunepark/gohwp/internal/type"
	"strings"
)

type HWPHeader struct {
	Version     HWPVersion   `validate:"required"`
	Signature   string       `validate:"required"`
	Attributes1 *Attributes1 `validate:"required"`
}

func NewHWPHeader(data []byte) (*HWPHeader, error) {
	// todo: use byteReader to read bytes
	expectedLength := 256

	if len(data) != expectedLength {
		return nil, &_type.ByteLengthError{
			ExpectedLength: expectedLength,
			ActualLength:   len(data),
		}
	}

	header := &HWPHeader{}
	var err error

	header.Signature = strings.Trim(string(data[:32]), "\x00")
	header.Version, err = NewHWPVersion(data[32:36])
	if err != nil {
		return nil, err
	}
	header.Attributes1, err = NewAttributes1(data[36:40])
	if err != nil {
		return nil, err
	}

	// todo: implement other properties(attribute2, encryptVersion in data[36:44])

	reserved := data[48:256]
	for _, b := range reserved {
		if b != 0 {
			return nil, fmt.Errorf("bytes in reserved area: %v", reserved)
		}
	}

	return header, nil
}
