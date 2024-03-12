package model

import (
	"fmt"
)

type ParaHeaderV1 struct {
	TextLength       uint32
	ControlMask      ControlMask
	ParaShapeIDRef   uint16
	ParaStyleIDRef   uint8
	DivisionType     DivisionType
	CharShapeInfoCnt uint16
	RangeTagInfoCnt  uint16
	AlignInfoCnt     uint16
	ParaInstanceID   uint32
}

func (ph *ParaHeaderV1) String() string {
	return fmt.Sprintf("TextLength: %d, ControlMask: %d, ParaShapeIDRef: %d, ParaStyleIDRef: %d, DivisionType: %s, CharShapeInfoCnt: %d, RangeTagInfoCnt: %d, AlignInfoCnt: %d, ParaInstanceID: %d",
		ph.TextLength, ph.ControlMask, ph.ParaShapeIDRef, ph.ParaStyleIDRef, ph.DivisionType.String(), ph.CharShapeInfoCnt, ph.RangeTagInfoCnt, ph.AlignInfoCnt, ph.ParaInstanceID)
}

type ParaHeader struct {
	ParaHeaderV1
	IsMergedTrack *uint16
}

func (ph *ParaHeader) String() string {
	phV1 := ph.ParaHeaderV1.String()
	var trackInfo string
	if ph.IsMergedTrack != nil {
		trackInfo = fmt.Sprintf(", IsMergedTrack: %d", *ph.IsMergedTrack)
	} else {
		trackInfo = ", IsMergedTrack: nil"
	}
	return phV1 + trackInfo
}
