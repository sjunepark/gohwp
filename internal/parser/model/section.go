package model

import (
	raw "github.com/sjunepark/hwp/internal/reader/model"
	"github.com/sjunepark/hwp/internal/util"
)

type Section struct {
	Paragraphs []*Paragraph
}

func (s *Section) String() (result string) {
	return util.JoinStringers(s.Paragraphs, "\n")
}

func (s *Section) parse(r *raw.Section) {
	for _, rawPara := range r.Paragraphs {
		para := &Paragraph{}
		para.parse(rawPara)
		s.Paragraphs = append(s.Paragraphs, para)
	}
}
