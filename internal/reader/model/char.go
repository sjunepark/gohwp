package model

type CharType int

const (
	CharTypeChar CharType = iota
	CharTypeInline
	CharTypeExtended
	CharTypeDefault
)

// todo: implement CharType as a struct with
