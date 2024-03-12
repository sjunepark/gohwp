package model

import "github.com/sjunepark/gohwp/internal/util"

type Section struct {
	Paragraphs []*Paragraph
}

func (s *Section) CurrentParagraph() *Paragraph {
	return s.Paragraphs[len(s.Paragraphs)-1]
}

func (s *Section) String() string {
	return util.JoinStringers(s.Paragraphs, "\n")
}
