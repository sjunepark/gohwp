package models

import "github.com/sjunepark/gohwp/internal/validator"

type Section struct {
	Width             uint32      `validate:"required"`
	Height            uint32      `validate:"required"`
	PaddingLeft       uint32      `validate:"required"`
	PaddingRight      uint32      `validate:"required"`
	PaddingTop        uint32      `validate:"required"`
	PaddingBottom     uint32      `validate:"required"`
	HeaderPadding     uint32      `validate:"required"`
	FooterPadding     uint32      `validate:"required"`
	Content           []Paragraph `validate:"required"`
	Orientation       uint32      `validate:"required"`
	BookBindingMethod uint32      `validate:"required"`
}

func NewSection(width, height, paddingLeft, paddingRight, paddingTop, paddingBottom, headerPadding, footerPadding, orientation, bookBindingMethod uint32, content []Paragraph) (Section, error) {
	var section Section
	section = Section{
		Width:             width,
		Height:            height,
		PaddingLeft:       paddingLeft,
		PaddingRight:      paddingRight,
		PaddingTop:        paddingTop,
		PaddingBottom:     paddingBottom,
		HeaderPadding:     headerPadding,
		FooterPadding:     footerPadding,
		Content:           content,
		Orientation:       orientation,
		BookBindingMethod: bookBindingMethod,
	}

	err := validator.ValidateStruct(section)
	if err != nil {
		return Section{}, err
	}
	return section, nil
}
