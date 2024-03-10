package types

type DocInfo struct {
	SectionSize         int
	CharShapes          []CharShape
	FontFaces           []FontFace
	BinData             []binData
	BorderFills         []BorderFill
	ParagraphShapes     []ParagraphShape
	StartingIndex       StartingIndex
	CaratLocation       CaratLocation
	CompatibleDocument  int
	LayoutCompatibility LayoutCompatibility
}

func (d *DocInfo) GetCharShape(index int) *CharShape {
	if index < len(d.CharShapes) {
		return &d.CharShapes[index]
	}
	return nil
}
