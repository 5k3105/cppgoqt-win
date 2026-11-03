package qt

type QTextCursor__MoveMode int64

const (
	QTextCursor__MoveAnchor QTextCursor__MoveMode = 0
	QTextCursor__KeepAnchor QTextCursor__MoveMode = 1
)

type QTextCursor__MoveOperation int64

const (
	QTextCursor__NoMove QTextCursor__MoveOperation = 0
	QTextCursor__Start QTextCursor__MoveOperation = 1
	QTextCursor__Up QTextCursor__MoveOperation = 2
	QTextCursor__StartOfLine QTextCursor__MoveOperation = 3
	QTextCursor__StartOfBlock QTextCursor__MoveOperation = 4
	QTextCursor__StartOfWord QTextCursor__MoveOperation = 5
	QTextCursor__PreviousBlock QTextCursor__MoveOperation = 6
	QTextCursor__PreviousCharacter QTextCursor__MoveOperation = 7
	QTextCursor__PreviousWord QTextCursor__MoveOperation = 8
	QTextCursor__Left QTextCursor__MoveOperation = 9
	QTextCursor__WordLeft QTextCursor__MoveOperation = 10
	QTextCursor__End QTextCursor__MoveOperation = 11
	QTextCursor__Down QTextCursor__MoveOperation = 12
	QTextCursor__EndOfLine QTextCursor__MoveOperation = 13
	QTextCursor__EndOfWord QTextCursor__MoveOperation = 14
	QTextCursor__EndOfBlock QTextCursor__MoveOperation = 15
	QTextCursor__NextBlock QTextCursor__MoveOperation = 16
	QTextCursor__NextCharacter QTextCursor__MoveOperation = 17
	QTextCursor__NextWord QTextCursor__MoveOperation = 18
	QTextCursor__Right QTextCursor__MoveOperation = 19
	QTextCursor__WordRight QTextCursor__MoveOperation = 20
	QTextCursor__NextCell QTextCursor__MoveOperation = 21
	QTextCursor__PreviousCell QTextCursor__MoveOperation = 22
	QTextCursor__NextRow QTextCursor__MoveOperation = 23
	QTextCursor__PreviousRow QTextCursor__MoveOperation = 24
)

type QTextCursor__SelectionType int64

const (
	QTextCursor__WordUnderCursor QTextCursor__SelectionType = 0
	QTextCursor__LineUnderCursor QTextCursor__SelectionType = 1
	QTextCursor__BlockUnderCursor QTextCursor__SelectionType = 2
	QTextCursor__Document QTextCursor__SelectionType = 3
)

