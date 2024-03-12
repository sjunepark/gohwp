package model

type HWPDocument struct {
	Header   *HWPHeader `validate:"required"`
	DocInfo  *DocInfo   `validate:"required"`
	BodyText []*Section `validate:"required"`
}