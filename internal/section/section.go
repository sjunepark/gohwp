package section

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
