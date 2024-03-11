package models

import (
	"unicode/utf16"
)

type ParaText []WChar

func (p *ParaText) String() string {
	var uint16s []uint16
	for _, wc := range *p {
		uint16s = append(uint16s, uint16(wc))
	}

	runes := utf16.Decode(uint16s)
	return string(runes)
}
