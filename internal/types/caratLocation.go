package types

type CaratLocation struct {
	ListID      int
	ParagraphID int
	CharIndex   int // 문단 내에서의 글자 단위 위치
}
