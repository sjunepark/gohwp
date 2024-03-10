package models

import (
	"github.com/sjunepark/gohwp/internal/types"
)

type BorderFillStyle struct {
	Left   BorderStyle
	Right  BorderStyle
	Top    BorderStyle
	Bottom BorderStyle
}

type BorderStyle struct {
	Type  int
	Width int
	Color types.ColorRef
}

type BorderFill struct {
	// todo: create getters and setters
	Attribute int
	Style     BorderFillStyle
	// todo: implement gradation
	BackgroundColor *types.ColorRef
}
