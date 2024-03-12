package parser

import (
	"github.com/sjunepark/gohwp/internal/parser/model"
	raw "github.com/sjunepark/gohwp/internal/reader/model"
)

func Parse(r *raw.Document) *model.Document {
	var doc *model.Document = &model.Document{}
	doc.Parse(r)
	return doc
}
