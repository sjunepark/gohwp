package parser

import (
	"github.com/sjunepark/hwp/internal/parser/model"
	raw "github.com/sjunepark/hwp/internal/reader/model"
)

func Parse(r *raw.Document) *model.Document {
	var doc *model.Document = &model.Document{}
	doc.Parse(r)
	return doc
}
