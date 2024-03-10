package models

type CharType int

const (
	Char CharType = iota
	Inline
	Extened
)

type HWPChar struct {
	Type CharType
	// This is implemented as number | string in hwp.js
	Value string
}
