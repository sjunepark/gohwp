package types

type CharType int

const (
	CharTypeChar CharType = iota
	CharTypeInline
	CharTypeExtended
)

type HWPChar struct {
	CharType
	Value string
}
