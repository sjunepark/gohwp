package model

type Document struct {
	Header   *HWPHeader
	DocInfo  *DocInfo
	BodyText []*Section
}
