package model

import (
	raw "github.com/sjunepark/gohwp/internal/reader/model"
	"github.com/sjunepark/gohwp/internal/util"
)

type Document struct {
	Header   *raw.HWPHeader
	DocInfo  *raw.DocInfo
	BodyText []*Section
}

func (d *Document) String() string {
	return util.JoinStringers(d.BodyText, "\n")
}

func (d *Document) Parse(r *raw.Document) {
	d.Header = r.Header
	d.DocInfo = r.DocInfo
	for _, rawSection := range r.BodyText {
		section := &Section{}
		section.parse(rawSection)
		d.BodyText = append(d.BodyText, section)
	}
}
