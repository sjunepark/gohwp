package model

import "fmt"

type ControlMask uint32

const (
	ControlMaskSectionOrColumn    ControlMask = 1 << 2  // Corresponding to "Ch" value 2
	ControlMaskFieldStart         ControlMask = 1 << 3  // Corresponding to "Ch" value 3
	ControlMaskDrawingObjectTable ControlMask = 1 << 11 // Corresponding to "Ch" value 11
	ControlMaskComment            ControlMask = 1 << 15 // Corresponding to "Ch" value 15
	ControlMaskHeaderFooter       ControlMask = 1 << 16 // Corresponding to "Ch" value 16
	ControlMaskFootnoteEndnote    ControlMask = 1 << 17 // Corresponding to "Ch" value 17
	ControlMaskAutoNumber         ControlMask = 1 << 18 // Corresponding to "Ch" value 18
	ControlMaskNewNumberHideNum   ControlMask = 1 << 21 // Corresponding to "Ch" value 21
	ControlMaskBookmarkIndexMark  ControlMask = 1 << 22 // Corresponding to "Ch" value 22
	ControlMaskAnnotationOverlap  ControlMask = 1 << 23 // Corresponding to "Ch" value 23
)

func (cm ControlMask) String() string {
	switch cm {
	case ControlMaskSectionOrColumn:
		return "Section or Column Definition"
	case ControlMaskFieldStart:
		return "Field Start"
	case ControlMaskDrawingObjectTable:
		return "Drawing Object or Table"
	case ControlMaskComment:
		return "Comment"
	case ControlMaskHeaderFooter:
		return "Header or Footer"
	case ControlMaskFootnoteEndnote:
		return "Footnote or Endnote"
	case ControlMaskAutoNumber:
		return "Auto Number"
	case ControlMaskNewNumberHideNum:
		return "New Number, Hide Number, or Number Position"
	case ControlMaskBookmarkIndexMark:
		return "Bookmark or Index Mark"
	case ControlMaskAnnotationOverlap:
		return "Annotation or Overlap"
	default:
		return fmt.Sprintf("Unknown ControlMask (0x%X)", uint32(cm))
	}
}

func (cm ControlMask) HasFeature(feature ControlMask) bool {
	return cm&feature != 0
}
