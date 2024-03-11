package models

type Section struct {
	Paragraphs []*Paragraph `validate:"required"`
}

func (s *Section) CurrentParagraph() *Paragraph {
	return s.Paragraphs[len(s.Paragraphs)-1]
}
