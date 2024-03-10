package docInfo

type DocumentProperties struct {
	SectionSize uint16
	CaratLocation
	StartingIndex
}

type StartingIndex struct {
	Page     uint16
	Footnote uint16
	Endnote  uint16
	Picture  uint16
	Table    uint16
	Equation uint16
}

type CaratLocation struct {
	ListID      uint32
	ParagraphID uint32
	CharIndex   uint32 // 문단 내에서의 글자 단위 위치
}
