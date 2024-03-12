package model

type Paragraph struct {
	ParaHeader *ParaHeader
	ParaText   ParaText
	// todo: implement other properties
}

func (p *Paragraph) String() string {
	return p.ParaText.String()
}
