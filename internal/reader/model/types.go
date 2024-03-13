package model

import (
	"github.com/sjunepark/hwp/internal/validator"
	"unicode/utf16"
)

// A byte is 8 bits. Multibyte types in this context are stored in little-endian order.

type Word uint16
type Dword uint32
type WChar uint16
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

func (wc WChar) CharType() CharType {
	switch wc {
	case 0: // 불가능한 문자(unusable char)
		return CharTypeChar

	case 10: // 한 줄 끝(line break)
		return CharTypeChar

	case 13: // 문단 끝(para break)
		return CharTypeChar

	case 1, 2, 3, 11, 12, 14, 15, 16, 17, 18, 21, 22, 23: // 확장 컨트롤 (extended control)
		return CharTypeExtended

	case 4, 5, 6, 7, 8, 9, 19, 20: // 인라인 컨트롤 (inline control)
		return CharTypeInline

	case 24: // 하이픈 (hyphen)
		return CharTypeChar

	case 30: // 묶음 빈칸 (non-breaking space)
		return CharTypeChar

	case 31: // 고정폭 빈칸 (fixed-width space)
		return CharTypeChar

	default: // 0-31 범위 내의 다른 컨트롤 문자들은 예약되어 있거나 특수한 용도로 사용됩니다.
		// 이 경우에는 일반 문자(char)로 처리할 수 있습니다.
		return CharTypeDefault
	}
}

func (wc WChar) String() string {
	runeValue := rune(wc)
	encoded := utf16.Encode([]rune{runeValue})
	decodedRune := utf16.Decode(encoded)
	return string(decodedRune)
}
