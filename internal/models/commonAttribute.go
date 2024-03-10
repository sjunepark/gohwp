package models

import "fmt"

type vertRelTo int

const (
	vertRelToPaper vertRelTo = iota
	vertRelToPage
	vertRelToParagraph
)

type horzRelTo int

const (
	horzRelToPage horzRelTo = iota
	horzRelToColumn
	horzRelToParagraph
)

type widthCriterion int

const (
	widthCriterionPaper widthCriterion = iota
	widthCriterionPage
	widthCriterionColumn
	widthCriterionParagraph
	widthCriterionAbsolute
)

type heightCriterion int

const (
	heightCriterionPaper heightCriterion = iota
	heightCriterionPage
	heightCriterionAbsolute
)

type textFlowMethod int

const (
	textFlowMethodSquare textFlowMethod = iota
	textFlowMethodTight
	textFlowMethodThrough
	textFlowMethodTopAndBottom
	textFlowMethodBehindText
	textFlowMethodInFrontOfText
)

type textHorzArrange int

const (
	textHorzArrangeBothSides textHorzArrange = iota
	textHorzArrangeLeftOnly
	textHorzArrangeRightOnly
	textHorzArrangeLargestOnly
)

type ObjectType int

const (
	ObjectTypeNone ObjectType = iota
	ObjectTypeFigure
	ObjectTypeTable
	ObjectTypeEquation
)

type commonAttribute struct {
	isTextLike           bool
	isApplyLineSpace     bool
	vertRelTo            vertRelTo
	vertRelativeArrange  int
	horzRelTo            horzRelTo
	horzRelativeArrange  int
	isVertRelToParaLimit bool
	isAllowOverlap       bool
	widthCriterion       widthCriterion
	heightCriterion      heightCriterion
	isProtectSize        bool
	textFlowMethod       textFlowMethod
	textHorzArrange      textHorzArrange
	objectType           ObjectType
}

func (ca *commonAttribute) setHorzRelTo(value int) error {
	switch value {
	// 한글 표준 문서에따르면 0과 1 모두 page 이다
	case 0, 1:
		ca.horzRelTo = horzRelToPage
	case 2:
		ca.horzRelTo = horzRelToColumn
	case 3:
		ca.horzRelTo = horzRelToParagraph
	default:
		return fmt.Errorf("invalid horzRelTo: %d", value)
	}
	return nil
}

func (ca *commonAttribute) getVertAlign() (string, error) {
	switch ca.vertRelativeArrange {
	case 0:
		if ca.vertRelTo == vertRelToPaper || ca.vertRelTo == vertRelToPage {
			return "top", nil
		}
		return "left", nil
	case 1:
		if ca.vertRelTo == vertRelToPaper || ca.vertRelTo == vertRelToPage {
			return "center", nil
		}
	case 2:
		if ca.vertRelTo == vertRelToPaper || ca.vertRelTo == vertRelToPage {
			return "bottom", nil
		}
		return "right", nil
	case 3:
		if ca.vertRelTo == vertRelToPaper || ca.vertRelTo == vertRelToPage {
			return "inside", nil
		}
	case 4:
		if ca.vertRelTo == vertRelToPaper || ca.vertRelTo == vertRelToPage {
			return "outside", nil
		}
	}
	return "", fmt.Errorf("invalid vertRelativeArrange: %d", ca.vertRelativeArrange)
}
