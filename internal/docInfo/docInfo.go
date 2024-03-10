package docInfo

import "github.com/sjunepark/gohwp/internal/models"

type DocInfo struct {
	DocumentProperties  *DocumentProperties
	CharShapes          []models.CharShape
	FontFaces           []models.FontFace
	BinData             []models.BinData
	BorderFills         []models.BorderFill
	ParagraphShapes     []models.ParagraphShape
	CompatibleDocument  int
	LayoutCompatibility models.LayoutCompatibility
}

func (d *DocInfo) GetCharShape(index int) *models.CharShape {
	if index < len(d.CharShapes) {
		return &d.CharShapes[index]
	}
	return nil
}
