package model

type HWPDocument struct {
	Header   *HWPHeader
	DocInfo  *DocInfo
	BodyText []*Section
}
