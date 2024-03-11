package models

type DivisionType uint8

const (
	DivisionTypeSection DivisionType = iota
	DivisionTypeMultiColumn
	DivisionTypePage
	DivisionTypeColumn
)

type ParaHeaderV1 struct {
	TextLength       uint32
	ControlMask      uint32
	ParaShapeIDRef   uint16
	ParaStyleIDRef   uint8
	DivisionKind     DivisionType
	CharShapeInfoCnt uint16
	RangeTagInfoCnt  uint16
	AlignInfoCnt     uint16
	ParaInstanceID   uint32
}

type ParaHeaderV2 struct {
	ParaHeaderV1
	IsMergedTrack uint16
}

type ParaHeader interface {
}

var _ ParaHeader = (*ParaHeaderV1)(nil)
var _ ParaHeader = (*ParaHeaderV2)(nil)
