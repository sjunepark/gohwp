package models

type Paragraph struct {
	content      []HWPChar
	shapeBuffer  []ShapePointer
	controls     []Control
	lineSegments []*LineSegment
	shapeIndex   int
	aligns       int
	textSize     int
}

/**
etc
*/

func (p *Paragraph) GetShapeEndPos(index int) (int, error) {
	if index >= len(p.shapeBuffer) {
		return len(p.content) - 1, &IndexOutOfRangeError{Index: index, Length: len(p.shapeBuffer)}
	}

	if index == len(p.shapeBuffer)-1 {
		return len(p.content) - 1, nil
	}

	return p.shapeBuffer[index+1].Pos - 1, nil
}

func (p *Paragraph) GetNextSize(index int) (int, error) {
	nextIndex := index + 1

	if nextIndex >= len(p.lineSegments) {
		return 0, &IndexOutOfRangeError{Index: nextIndex, Length: len(p.lineSegments)}
	}

	next := p.lineSegments[nextIndex]

	if next == nil {
		return p.textSize, nil
	}

	return next.Start, nil
}
