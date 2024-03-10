package models

type Control struct {
	baseControl
	tableControl
	shapeControls
	columnControl
}

type baseControl struct {
	id int
}

type tableControl struct {
	commonControl
	// todo: define setter
	tableAttribute int
	rowCount       int
	columnCount    int
	borderFillID   int
	content        [][]*paragraphList[tableColumnOption]
}

func (t *tableControl) addRow(row int, list *paragraphList[tableColumnOption]) {
	// Ensure that the content slice has enough rows to accommodate the new entry
	if len(t.content) <= row {
		t.content = append(t.content, []*paragraphList[tableColumnOption]{})
	}

	// Append the new list to the specified row in the content slice
	t.content[row] = append(t.content[row], list)

	// Optionally, update rowCount and columnCount if necessary
	t.rowCount = len(t.content)
	if len(t.content[row]) > t.columnCount {
		t.columnCount = len(t.content[row])
	}
}

type tableColumnOption struct {
	column       int
	row          int
	colSpan      int
	rowSpan      int
	width        int
	height       int
	padding      [4]int
	borderFillID *int
}

type shapeControls = shapeControl[picture]

type shapeControl[P any] struct {
	commonControl
	shapeType int
	info      *P
	content   []*paragraphList[P]
}

type picture struct {
	binItemID int
}

type columnControl struct {
	id          int
	columnType  columnType
	count       int
	direction   columnDirection
	isSameWidth bool
	gap         int
	widths      []int
	borderStyle int
	borderWidth int
	borderColor int
}

type columnType int

const (
	columnTypeNormal columnType = iota
	columnTypeFixed
	columnTypeAuto
)

type columnDirection int

const (
	columnDirectionLeft columnDirection = iota
	columnDirectionRight
	columnDirectionJustify
)

type commonControl struct {
	id               int
	attribute        commonAttribute
	verticalOffset   int
	horizontalOffset int
	width            int
	height           int
	zIndex           int
	margin           [4]int
	uid              int
	split            int
}
