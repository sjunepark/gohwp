package model

import raw "github.com/sjunepark/gohwp/internal/reader/model"

type Paragraph struct {
	ParaText string
}

func (p *Paragraph) String() string {
	return p.ParaText
}

func (p *Paragraph) parse(r *raw.Paragraph) {
	p.ParaText = r.ParaText.String()
}
