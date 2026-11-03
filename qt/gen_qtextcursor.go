package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QTextCursor_new uintptr
	addr_QTextCursor_new2 uintptr
	addr_QTextCursor_new3 uintptr
	addr_QTextCursor_new4 uintptr
	addr_QTextCursor_new5 uintptr
	addr_QTextCursor_OperatorAssign uintptr
	addr_QTextCursor_swap uintptr
	addr_QTextCursor_isNull uintptr
	addr_QTextCursor_setPosition uintptr
	addr_QTextCursor_position uintptr
	addr_QTextCursor_positionInBlock uintptr
	addr_QTextCursor_anchor uintptr
	addr_QTextCursor_insertText uintptr
	addr_QTextCursor_insertText2 uintptr
	addr_QTextCursor_movePosition uintptr
	addr_QTextCursor_visualNavigation uintptr
	addr_QTextCursor_setVisualNavigation uintptr
	addr_QTextCursor_setVerticalMovementX uintptr
	addr_QTextCursor_verticalMovementX uintptr
	addr_QTextCursor_setKeepPositionOnInsert uintptr
	addr_QTextCursor_keepPositionOnInsert uintptr
	addr_QTextCursor_deleteChar uintptr
	addr_QTextCursor_deletePreviousChar uintptr
	addr_QTextCursor_select uintptr
	addr_QTextCursor_hasSelection uintptr
	addr_QTextCursor_hasComplexSelection uintptr
	addr_QTextCursor_removeSelectedText uintptr
	addr_QTextCursor_clearSelection uintptr
	addr_QTextCursor_selectionStart uintptr
	addr_QTextCursor_selectionEnd uintptr
	addr_QTextCursor_selectedText uintptr
	addr_QTextCursor_selection uintptr
	addr_QTextCursor_selectedTableCells uintptr
	addr_QTextCursor_block uintptr
	addr_QTextCursor_charFormat uintptr
	addr_QTextCursor_setCharFormat uintptr
	addr_QTextCursor_mergeCharFormat uintptr
	addr_QTextCursor_blockFormat uintptr
	addr_QTextCursor_setBlockFormat uintptr
	addr_QTextCursor_mergeBlockFormat uintptr
	addr_QTextCursor_blockCharFormat uintptr
	addr_QTextCursor_setBlockCharFormat uintptr
	addr_QTextCursor_mergeBlockCharFormat uintptr
	addr_QTextCursor_atBlockStart uintptr
	addr_QTextCursor_atBlockEnd uintptr
	addr_QTextCursor_atStart uintptr
	addr_QTextCursor_atEnd uintptr
	addr_QTextCursor_insertBlock uintptr
	addr_QTextCursor_insertBlock2 uintptr
	addr_QTextCursor_insertBlock3 uintptr
	addr_QTextCursor_insertList uintptr
	addr_QTextCursor_insertList2 uintptr
	addr_QTextCursor_createList uintptr
	addr_QTextCursor_createList2 uintptr
	addr_QTextCursor_currentList uintptr
	addr_QTextCursor_insertTable uintptr
	addr_QTextCursor_insertTable2 uintptr
	addr_QTextCursor_currentTable uintptr
	addr_QTextCursor_insertFrame uintptr
	addr_QTextCursor_currentFrame uintptr
	addr_QTextCursor_insertFragment uintptr
	addr_QTextCursor_insertHtml uintptr
	addr_QTextCursor_insertMarkdown uintptr
	addr_QTextCursor_insertImage uintptr
	addr_QTextCursor_insertImage2 uintptr
	addr_QTextCursor_insertImage3 uintptr
	addr_QTextCursor_insertImage4 uintptr
	addr_QTextCursor_beginEditBlock uintptr
	addr_QTextCursor_joinPreviousEditBlock uintptr
	addr_QTextCursor_endEditBlock uintptr
	addr_QTextCursor_OperatorNotEqual uintptr
	addr_QTextCursor_OperatorLesser uintptr
	addr_QTextCursor_OperatorLesserOrEqual uintptr
	addr_QTextCursor_OperatorEqual uintptr
	addr_QTextCursor_OperatorGreaterOrEqual uintptr
	addr_QTextCursor_OperatorGreater uintptr
	addr_QTextCursor_isCopyOf uintptr
	addr_QTextCursor_blockNumber uintptr
	addr_QTextCursor_columnNumber uintptr
	addr_QTextCursor_document uintptr
	addr_QTextCursor_setPosition2 uintptr
	addr_QTextCursor_movePosition2 uintptr
	addr_QTextCursor_movePosition3 uintptr
	addr_QTextCursor_insertMarkdown2 uintptr
	addr_QTextCursor_insertImage5 uintptr
	pg_QTextCursor_delete func(uintptr)
)

func init() {
	symaddr(&addr_QTextCursor_new, "QTextCursor_new")
	symaddr(&addr_QTextCursor_new2, "QTextCursor_new2")
	symaddr(&addr_QTextCursor_new3, "QTextCursor_new3")
	symaddr(&addr_QTextCursor_new4, "QTextCursor_new4")
	symaddr(&addr_QTextCursor_new5, "QTextCursor_new5")
	symaddr(&addr_QTextCursor_OperatorAssign, "QTextCursor_OperatorAssign")
	symaddr(&addr_QTextCursor_swap, "QTextCursor_swap")
	symaddr(&addr_QTextCursor_isNull, "QTextCursor_isNull")
	symaddr(&addr_QTextCursor_setPosition, "QTextCursor_setPosition")
	symaddr(&addr_QTextCursor_position, "QTextCursor_position")
	symaddr(&addr_QTextCursor_positionInBlock, "QTextCursor_positionInBlock")
	symaddr(&addr_QTextCursor_anchor, "QTextCursor_anchor")
	symaddr(&addr_QTextCursor_insertText, "QTextCursor_insertText")
	symaddr(&addr_QTextCursor_insertText2, "QTextCursor_insertText2")
	symaddr(&addr_QTextCursor_movePosition, "QTextCursor_movePosition")
	symaddr(&addr_QTextCursor_visualNavigation, "QTextCursor_visualNavigation")
	symaddr(&addr_QTextCursor_setVisualNavigation, "QTextCursor_setVisualNavigation")
	symaddr(&addr_QTextCursor_setVerticalMovementX, "QTextCursor_setVerticalMovementX")
	symaddr(&addr_QTextCursor_verticalMovementX, "QTextCursor_verticalMovementX")
	symaddr(&addr_QTextCursor_setKeepPositionOnInsert, "QTextCursor_setKeepPositionOnInsert")
	symaddr(&addr_QTextCursor_keepPositionOnInsert, "QTextCursor_keepPositionOnInsert")
	symaddr(&addr_QTextCursor_deleteChar, "QTextCursor_deleteChar")
	symaddr(&addr_QTextCursor_deletePreviousChar, "QTextCursor_deletePreviousChar")
	symaddr(&addr_QTextCursor_select, "QTextCursor_select")
	symaddr(&addr_QTextCursor_hasSelection, "QTextCursor_hasSelection")
	symaddr(&addr_QTextCursor_hasComplexSelection, "QTextCursor_hasComplexSelection")
	symaddr(&addr_QTextCursor_removeSelectedText, "QTextCursor_removeSelectedText")
	symaddr(&addr_QTextCursor_clearSelection, "QTextCursor_clearSelection")
	symaddr(&addr_QTextCursor_selectionStart, "QTextCursor_selectionStart")
	symaddr(&addr_QTextCursor_selectionEnd, "QTextCursor_selectionEnd")
	symaddr(&addr_QTextCursor_selectedText, "QTextCursor_selectedText")
	symaddr(&addr_QTextCursor_selection, "QTextCursor_selection")
	symaddr(&addr_QTextCursor_selectedTableCells, "QTextCursor_selectedTableCells")
	symaddr(&addr_QTextCursor_block, "QTextCursor_block")
	symaddr(&addr_QTextCursor_charFormat, "QTextCursor_charFormat")
	symaddr(&addr_QTextCursor_setCharFormat, "QTextCursor_setCharFormat")
	symaddr(&addr_QTextCursor_mergeCharFormat, "QTextCursor_mergeCharFormat")
	symaddr(&addr_QTextCursor_blockFormat, "QTextCursor_blockFormat")
	symaddr(&addr_QTextCursor_setBlockFormat, "QTextCursor_setBlockFormat")
	symaddr(&addr_QTextCursor_mergeBlockFormat, "QTextCursor_mergeBlockFormat")
	symaddr(&addr_QTextCursor_blockCharFormat, "QTextCursor_blockCharFormat")
	symaddr(&addr_QTextCursor_setBlockCharFormat, "QTextCursor_setBlockCharFormat")
	symaddr(&addr_QTextCursor_mergeBlockCharFormat, "QTextCursor_mergeBlockCharFormat")
	symaddr(&addr_QTextCursor_atBlockStart, "QTextCursor_atBlockStart")
	symaddr(&addr_QTextCursor_atBlockEnd, "QTextCursor_atBlockEnd")
	symaddr(&addr_QTextCursor_atStart, "QTextCursor_atStart")
	symaddr(&addr_QTextCursor_atEnd, "QTextCursor_atEnd")
	symaddr(&addr_QTextCursor_insertBlock, "QTextCursor_insertBlock")
	symaddr(&addr_QTextCursor_insertBlock2, "QTextCursor_insertBlock2")
	symaddr(&addr_QTextCursor_insertBlock3, "QTextCursor_insertBlock3")
	symaddr(&addr_QTextCursor_insertList, "QTextCursor_insertList")
	symaddr(&addr_QTextCursor_insertList2, "QTextCursor_insertList2")
	symaddr(&addr_QTextCursor_createList, "QTextCursor_createList")
	symaddr(&addr_QTextCursor_createList2, "QTextCursor_createList2")
	symaddr(&addr_QTextCursor_currentList, "QTextCursor_currentList")
	symaddr(&addr_QTextCursor_insertTable, "QTextCursor_insertTable")
	symaddr(&addr_QTextCursor_insertTable2, "QTextCursor_insertTable2")
	symaddr(&addr_QTextCursor_currentTable, "QTextCursor_currentTable")
	symaddr(&addr_QTextCursor_insertFrame, "QTextCursor_insertFrame")
	symaddr(&addr_QTextCursor_currentFrame, "QTextCursor_currentFrame")
	symaddr(&addr_QTextCursor_insertFragment, "QTextCursor_insertFragment")
	symaddr(&addr_QTextCursor_insertHtml, "QTextCursor_insertHtml")
	symaddr(&addr_QTextCursor_insertMarkdown, "QTextCursor_insertMarkdown")
	symaddr(&addr_QTextCursor_insertImage, "QTextCursor_insertImage")
	symaddr(&addr_QTextCursor_insertImage2, "QTextCursor_insertImage2")
	symaddr(&addr_QTextCursor_insertImage3, "QTextCursor_insertImage3")
	symaddr(&addr_QTextCursor_insertImage4, "QTextCursor_insertImage4")
	symaddr(&addr_QTextCursor_beginEditBlock, "QTextCursor_beginEditBlock")
	symaddr(&addr_QTextCursor_joinPreviousEditBlock, "QTextCursor_joinPreviousEditBlock")
	symaddr(&addr_QTextCursor_endEditBlock, "QTextCursor_endEditBlock")
	symaddr(&addr_QTextCursor_OperatorNotEqual, "QTextCursor_OperatorNotEqual")
	symaddr(&addr_QTextCursor_OperatorLesser, "QTextCursor_OperatorLesser")
	symaddr(&addr_QTextCursor_OperatorLesserOrEqual, "QTextCursor_OperatorLesserOrEqual")
	symaddr(&addr_QTextCursor_OperatorEqual, "QTextCursor_OperatorEqual")
	symaddr(&addr_QTextCursor_OperatorGreaterOrEqual, "QTextCursor_OperatorGreaterOrEqual")
	symaddr(&addr_QTextCursor_OperatorGreater, "QTextCursor_OperatorGreater")
	symaddr(&addr_QTextCursor_isCopyOf, "QTextCursor_isCopyOf")
	symaddr(&addr_QTextCursor_blockNumber, "QTextCursor_blockNumber")
	symaddr(&addr_QTextCursor_columnNumber, "QTextCursor_columnNumber")
	symaddr(&addr_QTextCursor_document, "QTextCursor_document")
	symaddr(&addr_QTextCursor_setPosition2, "QTextCursor_setPosition2")
	symaddr(&addr_QTextCursor_movePosition2, "QTextCursor_movePosition2")
	symaddr(&addr_QTextCursor_movePosition3, "QTextCursor_movePosition3")
	symaddr(&addr_QTextCursor_insertMarkdown2, "QTextCursor_insertMarkdown2")
	symaddr(&addr_QTextCursor_insertImage5, "QTextCursor_insertImage5")
	reg(&pg_QTextCursor_delete, "QTextCursor_delete")
}

type QTextCursor struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QTextCursor) UnsafeGoID() int64     { return this.vid }
func (this *QTextCursor) UnsafeSetGoID(id int64) { this.vid = id }

func bindQTextCursor(h uintptr) *QTextCursor {
	if h == 0 {
		return nil
	}
	return &QTextCursor{h: h}
}

func newownedQTextCursor(h uintptr) *QTextCursor {
	o := bindQTextCursor(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QTextCursor).Delete)
	}
	return o
}

func (this *QTextCursor) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QTextCursor) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQTextCursor(h unsafe.Pointer) *QTextCursor { return bindQTextCursor(uintptr(h)) }
func UnsafeNewQTextCursorOwned(h unsafe.Pointer) *QTextCursor { return newownedQTextCursor(uintptr(h)) }

func (this *QTextCursor) UnsafeRelease() { this.h = 0 }

func NewQTextCursor() *QTextCursor {
	__r, _, _ := purego.SyscallN(addr_QTextCursor_new)
	__o := newownedQTextCursor(__r)
	return __o
}

func NewQTextCursor2(document *QTextDocument) *QTextCursor {
	__r, _, _ := purego.SyscallN(addr_QTextCursor_new2, uintptr(document.UnsafePointer()))
	__o := newownedQTextCursor(__r)
	runtime.KeepAlive(document)
	return __o
}

func NewQTextCursor3(frame *QTextFrame) *QTextCursor {
	__r, _, _ := purego.SyscallN(addr_QTextCursor_new3, uintptr(frame.UnsafePointer()))
	__o := newownedQTextCursor(__r)
	runtime.KeepAlive(frame)
	return __o
}

func NewQTextCursor4(block *QTextBlock) *QTextCursor {
	__r, _, _ := purego.SyscallN(addr_QTextCursor_new4, uintptr(block.UnsafePointer()))
	__o := newownedQTextCursor(__r)
	runtime.KeepAlive(block)
	return __o
}

func NewQTextCursor5(cursor *QTextCursor) *QTextCursor {
	__r, _, _ := purego.SyscallN(addr_QTextCursor_new5, uintptr(cursor.UnsafePointer()))
	__o := newownedQTextCursor(__r)
	runtime.KeepAlive(cursor)
	return __o
}

func (this *QTextCursor) OperatorAssign(other *QTextCursor) *QTextCursor {
	__r, _, _ := purego.SyscallN(addr_QTextCursor_OperatorAssign, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return bindQTextCursor(__r)
}

func (this *QTextCursor) Swap(other *QTextCursor) {
	purego.SyscallN(addr_QTextCursor_swap, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
}

func (this *QTextCursor) IsNull() bool {
	__r0, _, _ := purego.SyscallN(addr_QTextCursor_isNull, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTextCursor) SetPosition(pos int) {
	purego.SyscallN(addr_QTextCursor_setPosition, this.h, uintptr(pos))
	runtime.KeepAlive(this)
}

func (this *QTextCursor) Position() int {
	__r0, _, _ := purego.SyscallN(addr_QTextCursor_position, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTextCursor) PositionInBlock() int {
	__r0, _, _ := purego.SyscallN(addr_QTextCursor_positionInBlock, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTextCursor) Anchor() int {
	__r0, _, _ := purego.SyscallN(addr_QTextCursor_anchor, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTextCursor) InsertText(text string) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QTextCursor_insertText, this.h, text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
}

func (this *QTextCursor) InsertText2(text string, format *QTextCharFormat) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QTextCursor_insertText2, this.h, text_p, uintptr(len(text)), uintptr(format.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
	runtime.KeepAlive(format)
}

func (this *QTextCursor) MovePosition(op QTextCursor__MoveOperation) bool {
	__r0, _, _ := purego.SyscallN(addr_QTextCursor_movePosition, this.h, uintptr(op))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTextCursor) VisualNavigation() bool {
	__r0, _, _ := purego.SyscallN(addr_QTextCursor_visualNavigation, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTextCursor) SetVisualNavigation(b bool) {
	purego.SyscallN(addr_QTextCursor_setVisualNavigation, this.h, uintptr(b2i(b)))
	runtime.KeepAlive(this)
}

func (this *QTextCursor) SetVerticalMovementX(x int) {
	purego.SyscallN(addr_QTextCursor_setVerticalMovementX, this.h, uintptr(x))
	runtime.KeepAlive(this)
}

func (this *QTextCursor) VerticalMovementX() int {
	__r0, _, _ := purego.SyscallN(addr_QTextCursor_verticalMovementX, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTextCursor) SetKeepPositionOnInsert(b bool) {
	purego.SyscallN(addr_QTextCursor_setKeepPositionOnInsert, this.h, uintptr(b2i(b)))
	runtime.KeepAlive(this)
}

func (this *QTextCursor) KeepPositionOnInsert() bool {
	__r0, _, _ := purego.SyscallN(addr_QTextCursor_keepPositionOnInsert, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTextCursor) DeleteChar() {
	purego.SyscallN(addr_QTextCursor_deleteChar, this.h)
	runtime.KeepAlive(this)
}

func (this *QTextCursor) DeletePreviousChar() {
	purego.SyscallN(addr_QTextCursor_deletePreviousChar, this.h)
	runtime.KeepAlive(this)
}

func (this *QTextCursor) Select(selection QTextCursor__SelectionType) {
	purego.SyscallN(addr_QTextCursor_select, this.h, uintptr(selection))
	runtime.KeepAlive(this)
}

func (this *QTextCursor) HasSelection() bool {
	__r0, _, _ := purego.SyscallN(addr_QTextCursor_hasSelection, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTextCursor) HasComplexSelection() bool {
	__r0, _, _ := purego.SyscallN(addr_QTextCursor_hasComplexSelection, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTextCursor) RemoveSelectedText() {
	purego.SyscallN(addr_QTextCursor_removeSelectedText, this.h)
	runtime.KeepAlive(this)
}

func (this *QTextCursor) ClearSelection() {
	purego.SyscallN(addr_QTextCursor_clearSelection, this.h)
	runtime.KeepAlive(this)
}

func (this *QTextCursor) SelectionStart() int {
	__r0, _, _ := purego.SyscallN(addr_QTextCursor_selectionStart, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTextCursor) SelectionEnd() int {
	__r0, _, _ := purego.SyscallN(addr_QTextCursor_selectionEnd, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTextCursor) SelectedText() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTextCursor_selectedText, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QTextCursor) Selection() *QTextDocumentFragment {
	__r, _, _ := purego.SyscallN(addr_QTextCursor_selection, this.h)
	runtime.KeepAlive(this)
	return newownedQTextDocumentFragment(__r)
}

func (this *QTextCursor) SelectedTableCells() (int, int, int, int) {
	var firstRow_out int32
	var numRows_out int32
	var firstColumn_out int32
	var numColumns_out int32
	purego.SyscallN(addr_QTextCursor_selectedTableCells, this.h, uintptr(unsafe.Pointer(&firstRow_out)), uintptr(unsafe.Pointer(&numRows_out)), uintptr(unsafe.Pointer(&firstColumn_out)), uintptr(unsafe.Pointer(&numColumns_out)))
	runtime.KeepAlive(this)
	return int(firstRow_out), int(numRows_out), int(firstColumn_out), int(numColumns_out)
}

func (this *QTextCursor) Block() *QTextBlock {
	__r, _, _ := purego.SyscallN(addr_QTextCursor_block, this.h)
	runtime.KeepAlive(this)
	return newownedQTextBlock(__r)
}

func (this *QTextCursor) CharFormat() *QTextCharFormat {
	__r, _, _ := purego.SyscallN(addr_QTextCursor_charFormat, this.h)
	runtime.KeepAlive(this)
	return newownedQTextCharFormat(__r)
}

func (this *QTextCursor) SetCharFormat(format *QTextCharFormat) {
	purego.SyscallN(addr_QTextCursor_setCharFormat, this.h, uintptr(format.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(format)
}

func (this *QTextCursor) MergeCharFormat(modifier *QTextCharFormat) {
	purego.SyscallN(addr_QTextCursor_mergeCharFormat, this.h, uintptr(modifier.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(modifier)
}

func (this *QTextCursor) BlockFormat() *QTextBlockFormat {
	__r, _, _ := purego.SyscallN(addr_QTextCursor_blockFormat, this.h)
	runtime.KeepAlive(this)
	return newownedQTextBlockFormat(__r)
}

func (this *QTextCursor) SetBlockFormat(format *QTextBlockFormat) {
	purego.SyscallN(addr_QTextCursor_setBlockFormat, this.h, uintptr(format.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(format)
}

func (this *QTextCursor) MergeBlockFormat(modifier *QTextBlockFormat) {
	purego.SyscallN(addr_QTextCursor_mergeBlockFormat, this.h, uintptr(modifier.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(modifier)
}

func (this *QTextCursor) BlockCharFormat() *QTextCharFormat {
	__r, _, _ := purego.SyscallN(addr_QTextCursor_blockCharFormat, this.h)
	runtime.KeepAlive(this)
	return newownedQTextCharFormat(__r)
}

func (this *QTextCursor) SetBlockCharFormat(format *QTextCharFormat) {
	purego.SyscallN(addr_QTextCursor_setBlockCharFormat, this.h, uintptr(format.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(format)
}

func (this *QTextCursor) MergeBlockCharFormat(modifier *QTextCharFormat) {
	purego.SyscallN(addr_QTextCursor_mergeBlockCharFormat, this.h, uintptr(modifier.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(modifier)
}

func (this *QTextCursor) AtBlockStart() bool {
	__r0, _, _ := purego.SyscallN(addr_QTextCursor_atBlockStart, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTextCursor) AtBlockEnd() bool {
	__r0, _, _ := purego.SyscallN(addr_QTextCursor_atBlockEnd, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTextCursor) AtStart() bool {
	__r0, _, _ := purego.SyscallN(addr_QTextCursor_atStart, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTextCursor) AtEnd() bool {
	__r0, _, _ := purego.SyscallN(addr_QTextCursor_atEnd, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTextCursor) InsertBlock() {
	purego.SyscallN(addr_QTextCursor_insertBlock, this.h)
	runtime.KeepAlive(this)
}

func (this *QTextCursor) InsertBlockWithFormat(format *QTextBlockFormat) {
	purego.SyscallN(addr_QTextCursor_insertBlock2, this.h, uintptr(format.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(format)
}

func (this *QTextCursor) InsertBlock2(format *QTextBlockFormat, charFormat *QTextCharFormat) {
	purego.SyscallN(addr_QTextCursor_insertBlock3, this.h, uintptr(format.UnsafePointer()), uintptr(charFormat.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(format)
	runtime.KeepAlive(charFormat)
}

func (this *QTextCursor) InsertList(format *QTextListFormat) *QTextList {
	__r, _, _ := purego.SyscallN(addr_QTextCursor_insertList, this.h, uintptr(format.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(format)
	return bindQTextList(__r)
}

func (this *QTextCursor) InsertListWithStyle(style QTextListFormat__Style) *QTextList {
	__r, _, _ := purego.SyscallN(addr_QTextCursor_insertList2, this.h, uintptr(style))
	runtime.KeepAlive(this)
	return bindQTextList(__r)
}

func (this *QTextCursor) CreateList(format *QTextListFormat) *QTextList {
	__r, _, _ := purego.SyscallN(addr_QTextCursor_createList, this.h, uintptr(format.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(format)
	return bindQTextList(__r)
}

func (this *QTextCursor) CreateListWithStyle(style QTextListFormat__Style) *QTextList {
	__r, _, _ := purego.SyscallN(addr_QTextCursor_createList2, this.h, uintptr(style))
	runtime.KeepAlive(this)
	return bindQTextList(__r)
}

func (this *QTextCursor) CurrentList() *QTextList {
	__r, _, _ := purego.SyscallN(addr_QTextCursor_currentList, this.h)
	runtime.KeepAlive(this)
	return bindQTextList(__r)
}

func (this *QTextCursor) InsertTable(rows int, cols int, format *QTextTableFormat) *QTextTable {
	__r, _, _ := purego.SyscallN(addr_QTextCursor_insertTable, this.h, uintptr(rows), uintptr(cols), uintptr(format.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(format)
	return bindQTextTable(__r)
}

func (this *QTextCursor) InsertTable2(rows int, cols int) *QTextTable {
	__r, _, _ := purego.SyscallN(addr_QTextCursor_insertTable2, this.h, uintptr(rows), uintptr(cols))
	runtime.KeepAlive(this)
	return bindQTextTable(__r)
}

func (this *QTextCursor) CurrentTable() *QTextTable {
	__r, _, _ := purego.SyscallN(addr_QTextCursor_currentTable, this.h)
	runtime.KeepAlive(this)
	return bindQTextTable(__r)
}

func (this *QTextCursor) InsertFrame(format *QTextFrameFormat) *QTextFrame {
	__r, _, _ := purego.SyscallN(addr_QTextCursor_insertFrame, this.h, uintptr(format.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(format)
	return bindQTextFrame(__r)
}

func (this *QTextCursor) CurrentFrame() *QTextFrame {
	__r, _, _ := purego.SyscallN(addr_QTextCursor_currentFrame, this.h)
	runtime.KeepAlive(this)
	return bindQTextFrame(__r)
}

func (this *QTextCursor) InsertFragment(fragment *QTextDocumentFragment) {
	purego.SyscallN(addr_QTextCursor_insertFragment, this.h, uintptr(fragment.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(fragment)
}

func (this *QTextCursor) InsertHtml(html string) {
	var html_p uintptr
	if len(html) > 0 {
		html_p = uintptr(unsafe.Pointer(unsafe.StringData(html)))
	}
	purego.SyscallN(addr_QTextCursor_insertHtml, this.h, html_p, uintptr(len(html)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(html)
}

func (this *QTextCursor) InsertMarkdown(markdown string) {
	var markdown_p uintptr
	if len(markdown) > 0 {
		markdown_p = uintptr(unsafe.Pointer(unsafe.StringData(markdown)))
	}
	purego.SyscallN(addr_QTextCursor_insertMarkdown, this.h, markdown_p, uintptr(len(markdown)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(markdown)
}

func (this *QTextCursor) InsertImage(format *QTextImageFormat, alignment QTextFrameFormat__Position) {
	purego.SyscallN(addr_QTextCursor_insertImage, this.h, uintptr(format.UnsafePointer()), uintptr(alignment))
	runtime.KeepAlive(this)
	runtime.KeepAlive(format)
}

func (this *QTextCursor) InsertImageWithFormat(format *QTextImageFormat) {
	purego.SyscallN(addr_QTextCursor_insertImage2, this.h, uintptr(format.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(format)
}

func (this *QTextCursor) InsertImageWithName(name string) {
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	purego.SyscallN(addr_QTextCursor_insertImage3, this.h, name_p, uintptr(len(name)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(name)
}

func (this *QTextCursor) InsertImageWithImage(image *QImage) {
	purego.SyscallN(addr_QTextCursor_insertImage4, this.h, uintptr(image.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(image)
}

func (this *QTextCursor) BeginEditBlock() {
	purego.SyscallN(addr_QTextCursor_beginEditBlock, this.h)
	runtime.KeepAlive(this)
}

func (this *QTextCursor) JoinPreviousEditBlock() {
	purego.SyscallN(addr_QTextCursor_joinPreviousEditBlock, this.h)
	runtime.KeepAlive(this)
}

func (this *QTextCursor) EndEditBlock() {
	purego.SyscallN(addr_QTextCursor_endEditBlock, this.h)
	runtime.KeepAlive(this)
}

func (this *QTextCursor) OperatorNotEqual(rhs *QTextCursor) bool {
	__r0, _, _ := purego.SyscallN(addr_QTextCursor_OperatorNotEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QTextCursor) OperatorLesser(rhs *QTextCursor) bool {
	__r0, _, _ := purego.SyscallN(addr_QTextCursor_OperatorLesser, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QTextCursor) OperatorLesserOrEqual(rhs *QTextCursor) bool {
	__r0, _, _ := purego.SyscallN(addr_QTextCursor_OperatorLesserOrEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QTextCursor) OperatorEqual(rhs *QTextCursor) bool {
	__r0, _, _ := purego.SyscallN(addr_QTextCursor_OperatorEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QTextCursor) OperatorGreaterOrEqual(rhs *QTextCursor) bool {
	__r0, _, _ := purego.SyscallN(addr_QTextCursor_OperatorGreaterOrEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QTextCursor) OperatorGreater(rhs *QTextCursor) bool {
	__r0, _, _ := purego.SyscallN(addr_QTextCursor_OperatorGreater, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QTextCursor) IsCopyOf(other *QTextCursor) bool {
	__r0, _, _ := purego.SyscallN(addr_QTextCursor_isCopyOf, this.h, uintptr(other.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return __r != 0
}

func (this *QTextCursor) BlockNumber() int {
	__r0, _, _ := purego.SyscallN(addr_QTextCursor_blockNumber, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTextCursor) ColumnNumber() int {
	__r0, _, _ := purego.SyscallN(addr_QTextCursor_columnNumber, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTextCursor) Document() *QTextDocument {
	__r, _, _ := purego.SyscallN(addr_QTextCursor_document, this.h)
	runtime.KeepAlive(this)
	return bindQTextDocument(__r)
}

func (this *QTextCursor) SetPosition2(pos int, mode QTextCursor__MoveMode) {
	purego.SyscallN(addr_QTextCursor_setPosition2, this.h, uintptr(pos), uintptr(mode))
	runtime.KeepAlive(this)
}

func (this *QTextCursor) MovePosition2(op QTextCursor__MoveOperation, arg2 QTextCursor__MoveMode) bool {
	__r0, _, _ := purego.SyscallN(addr_QTextCursor_movePosition2, this.h, uintptr(op), uintptr(arg2))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTextCursor) MovePosition3(op QTextCursor__MoveOperation, arg2 QTextCursor__MoveMode, n int) bool {
	__r0, _, _ := purego.SyscallN(addr_QTextCursor_movePosition3, this.h, uintptr(op), uintptr(arg2), uintptr(n))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTextCursor) InsertMarkdown2(markdown string, features QTextDocument__MarkdownFeature) {
	var markdown_p uintptr
	if len(markdown) > 0 {
		markdown_p = uintptr(unsafe.Pointer(unsafe.StringData(markdown)))
	}
	purego.SyscallN(addr_QTextCursor_insertMarkdown2, this.h, markdown_p, uintptr(len(markdown)), uintptr(features))
	runtime.KeepAlive(this)
	runtime.KeepAlive(markdown)
}

func (this *QTextCursor) InsertImage2(image *QImage, name string) {
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	purego.SyscallN(addr_QTextCursor_insertImage5, this.h, uintptr(image.UnsafePointer()), name_p, uintptr(len(name)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(image)
	runtime.KeepAlive(name)
}

func (this *QTextCursor) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QTextCursor_delete(this.h)
	}
	this.UnsafeRelease()
}
