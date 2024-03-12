package model

type Section struct {
	Paragraphs []*Paragraph `validate:"required"`
}

func (s *Section) CurrentParagraph() *Paragraph {
	return s.Paragraphs[len(s.Paragraphs)-1]
}

func (s *Section) String() string {
	var str string
	for i, para := range s.Paragraphs {
		if i > 0 {
			str += "\n" // Add a newline before each paragraph except the first one
		}
		str += para.String()
	}
	return str
}
