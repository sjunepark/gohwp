package model

type DocInfo struct {
	DocumentProperties  *DocumentProperties
	CharShapes          []CharShape
	FontFaces           []FontFace
	BinData             []BinData
	BorderFills         []BorderFill
	ParagraphShapes     []ParagraphShape
	CompatibleDocument  int
	LayoutCompatibility LayoutCompatibility
}
