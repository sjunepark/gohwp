package models

import "fmt"

type DivisionType uint8

const (
	DivisionTypeSection DivisionType = iota
	DivisionTypeMultiColumn
	DivisionTypePage
	DivisionTypeColumn
)

func (dt DivisionType) String() string {
	switch dt {
	case DivisionTypeSection:
		return "Section"
	case DivisionTypeMultiColumn:
		return "MultiColumn"
	case DivisionTypePage:
		return "Page"
	case DivisionTypeColumn:
		return "Column"
	default:
		return fmt.Sprintf("Unknown DivisionType (%d)", dt)
	}
}

type ParaHeaderV1 struct {
	TextLength       uint32
	ControlMask      uint32
	ParaShapeIDRef   uint16
	ParaStyleIDRef   uint8
	DivisionType     DivisionType
	CharShapeInfoCnt uint16
	RangeTagInfoCnt  uint16
	AlignInfoCnt     uint16
	ParaInstanceID   uint32
}

func (ph ParaHeaderV1) String() string {
	return fmt.Sprintf("TextLength: %d, ControlMask: %d, ParaShapeIDRef: %d, ParaStyleIDRef: %d, DivisionType: %s, CharShapeInfoCnt: %d, RangeTagInfoCnt: %d, AlignInfoCnt: %d, ParaInstanceID: %d",
		ph.TextLength, ph.ControlMask, ph.ParaShapeIDRef, ph.ParaStyleIDRef, ph.DivisionType.String(), ph.CharShapeInfoCnt, ph.RangeTagInfoCnt, ph.AlignInfoCnt, ph.ParaInstanceID)
}

type ParaHeaderV2 struct {
	ParaHeaderV1
	IsMergedTrack uint16
}

func (ph ParaHeaderV2) String() string {
	phV1 := ph.ParaHeaderV1.String()
	return fmt.Sprintf("%s, IsMergedTrack: %d", phV1[:len(phV1)-1], ph.IsMergedTrack)
}

type ParaHeader interface {
	String() string
}

var _ ParaHeader = (*ParaHeaderV1)(nil)
var _ ParaHeader = (*ParaHeaderV2)(nil)
