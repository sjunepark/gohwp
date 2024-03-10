package models

import "fmt"

type IndexOutOfRangeError struct {
	Index  int
	Length int
}

func (e *IndexOutOfRangeError) Error() string {
	return fmt.Sprintf("index out of range: %d (length: %d)", e.Index, e.Length)
}

type ByteLengthError struct {
	ExpectedLength int
	ActualLength   int
}

func (e *ByteLengthError) Error() string {
	return fmt.Sprintf("invalid byte length: expected %d, actual %d", e.ExpectedLength, e.ActualLength)
}

type OutOfBoundsError struct {
	Requested int
	Max       int
}

func (e *OutOfBoundsError) Error() string {
	return fmt.Sprintf("out of bounds: requested %d, max %d", e.Requested, e.Max)
}
