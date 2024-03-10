package types

import "github.com/sjunepark/gohwp/internal/validator"

// A byte is 8 bits. Multibyte types in this context are stored in little-endian order.

type Word uint16
type Dword uint32
type WChar rune
type HWPUnit uint32
type SHWPUnit int32
type HWPUnit16 int16

type ColorRef struct {
	Red   uint8 `validate:"gte=0,lte=255"`
	Green uint8 `validate:"gte=0,lte=255"`
	Blue  uint8 `validate:"gte=0,lte=255"`
}

func NewColorRef(red, green, blue uint8) (ColorRef, error) {
	err := validator.ValidateStruct(ColorRef{Red: red, Green: green, Blue: blue})
	if err != nil {
		return ColorRef{}, err
	}
	return ColorRef{Red: red, Green: green, Blue: blue}, nil
}

type ByteStream []byte
