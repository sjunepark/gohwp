package types

type paragraphList[P any] struct {
	Attribute P
	Items     []Paragraph
}
