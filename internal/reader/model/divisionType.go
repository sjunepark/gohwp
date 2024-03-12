package model

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
