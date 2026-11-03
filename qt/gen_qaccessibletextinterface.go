package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QAccessibleTextInterface_selection uintptr
	addr_QAccessibleTextInterface_selectionCount uintptr
	addr_QAccessibleTextInterface_addSelection uintptr
	addr_QAccessibleTextInterface_removeSelection uintptr
	addr_QAccessibleTextInterface_setSelection uintptr
	addr_QAccessibleTextInterface_cursorPosition uintptr
	addr_QAccessibleTextInterface_setCursorPosition uintptr
	addr_QAccessibleTextInterface_text uintptr
	addr_QAccessibleTextInterface_textBeforeOffset uintptr
	addr_QAccessibleTextInterface_textAfterOffset uintptr
	addr_QAccessibleTextInterface_textAtOffset uintptr
	addr_QAccessibleTextInterface_characterCount uintptr
	addr_QAccessibleTextInterface_characterRect uintptr
	addr_QAccessibleTextInterface_offsetAtPoint uintptr
	addr_QAccessibleTextInterface_scrollToSubstring uintptr
	addr_QAccessibleTextInterface_attributes uintptr
	addr_QAccessibleTextInterface_OperatorAssign uintptr
	addr_QAccessibleTextInterface_textBeforeOffset_default uintptr
	addr_QAccessibleTextInterface_textAfterOffset_default uintptr
	addr_QAccessibleTextInterface_textAtOffset_default uintptr
	pg_QAccessibleTextInterface__set_override func(uintptr, int32)
	pg_QAccessibleTextInterface__Go_new func(int64) uintptr
	pg_QAccessibleTextInterface_delete func(uintptr)
)

func init() {
	symaddr(&addr_QAccessibleTextInterface_selection, "QAccessibleTextInterface_selection")
	symaddr(&addr_QAccessibleTextInterface_selectionCount, "QAccessibleTextInterface_selectionCount")
	symaddr(&addr_QAccessibleTextInterface_addSelection, "QAccessibleTextInterface_addSelection")
	symaddr(&addr_QAccessibleTextInterface_removeSelection, "QAccessibleTextInterface_removeSelection")
	symaddr(&addr_QAccessibleTextInterface_setSelection, "QAccessibleTextInterface_setSelection")
	symaddr(&addr_QAccessibleTextInterface_cursorPosition, "QAccessibleTextInterface_cursorPosition")
	symaddr(&addr_QAccessibleTextInterface_setCursorPosition, "QAccessibleTextInterface_setCursorPosition")
	symaddr(&addr_QAccessibleTextInterface_text, "QAccessibleTextInterface_text")
	symaddr(&addr_QAccessibleTextInterface_textBeforeOffset, "QAccessibleTextInterface_textBeforeOffset")
	symaddr(&addr_QAccessibleTextInterface_textAfterOffset, "QAccessibleTextInterface_textAfterOffset")
	symaddr(&addr_QAccessibleTextInterface_textAtOffset, "QAccessibleTextInterface_textAtOffset")
	symaddr(&addr_QAccessibleTextInterface_characterCount, "QAccessibleTextInterface_characterCount")
	symaddr(&addr_QAccessibleTextInterface_characterRect, "QAccessibleTextInterface_characterRect")
	symaddr(&addr_QAccessibleTextInterface_offsetAtPoint, "QAccessibleTextInterface_offsetAtPoint")
	symaddr(&addr_QAccessibleTextInterface_scrollToSubstring, "QAccessibleTextInterface_scrollToSubstring")
	symaddr(&addr_QAccessibleTextInterface_attributes, "QAccessibleTextInterface_attributes")
	symaddr(&addr_QAccessibleTextInterface_OperatorAssign, "QAccessibleTextInterface_OperatorAssign")
	symaddr(&addr_QAccessibleTextInterface_textBeforeOffset_default, "QAccessibleTextInterface_textBeforeOffset_default")
	symaddr(&addr_QAccessibleTextInterface_textAfterOffset_default, "QAccessibleTextInterface_textAfterOffset_default")
	symaddr(&addr_QAccessibleTextInterface_textAtOffset_default, "QAccessibleTextInterface_textAtOffset_default")
	reg(&pg_QAccessibleTextInterface__set_override, "QAccessibleTextInterface__set_override")
	reg(&pg_QAccessibleTextInterface__Go_new, "QAccessibleTextInterface__Go_new")
	reg(&pg_QAccessibleTextInterface_delete, "QAccessibleTextInterface_delete")
}

type QAccessibleTextInterface struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QAccessibleTextInterface) UnsafeGoID() int64     { return this.vid }
func (this *QAccessibleTextInterface) UnsafeSetGoID(id int64) { this.vid = id }

func bindQAccessibleTextInterface(h uintptr) *QAccessibleTextInterface {
	if h == 0 {
		return nil
	}
	return &QAccessibleTextInterface{h: h}
}

func newownedQAccessibleTextInterface(h uintptr) *QAccessibleTextInterface {
	o := bindQAccessibleTextInterface(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QAccessibleTextInterface).Delete)
	}
	return o
}

func (this *QAccessibleTextInterface) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QAccessibleTextInterface) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQAccessibleTextInterface(h unsafe.Pointer) *QAccessibleTextInterface { return bindQAccessibleTextInterface(uintptr(h)) }
func UnsafeNewQAccessibleTextInterfaceOwned(h unsafe.Pointer) *QAccessibleTextInterface { return newownedQAccessibleTextInterface(uintptr(h)) }

func (this *QAccessibleTextInterface) UnsafeRelease() { this.h = 0 }

func (this *QAccessibleTextInterface) Selection(selectionIndex int) (int, int) {
	var startOffset_out int32
	var endOffset_out int32
	purego.SyscallN(addr_QAccessibleTextInterface_selection, this.h, uintptr(selectionIndex), uintptr(unsafe.Pointer(&startOffset_out)), uintptr(unsafe.Pointer(&endOffset_out)))
	runtime.KeepAlive(this)
	return int(startOffset_out), int(endOffset_out)
}

func (this *QAccessibleTextInterface) SelectionCount() int {
	__r0, _, _ := purego.SyscallN(addr_QAccessibleTextInterface_selectionCount, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAccessibleTextInterface) AddSelection(startOffset int, endOffset int) {
	purego.SyscallN(addr_QAccessibleTextInterface_addSelection, this.h, uintptr(startOffset), uintptr(endOffset))
	runtime.KeepAlive(this)
}

func (this *QAccessibleTextInterface) RemoveSelection(selectionIndex int) {
	purego.SyscallN(addr_QAccessibleTextInterface_removeSelection, this.h, uintptr(selectionIndex))
	runtime.KeepAlive(this)
}

func (this *QAccessibleTextInterface) SetSelection(selectionIndex int, startOffset int, endOffset int) {
	purego.SyscallN(addr_QAccessibleTextInterface_setSelection, this.h, uintptr(selectionIndex), uintptr(startOffset), uintptr(endOffset))
	runtime.KeepAlive(this)
}

func (this *QAccessibleTextInterface) CursorPosition() int {
	__r0, _, _ := purego.SyscallN(addr_QAccessibleTextInterface_cursorPosition, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAccessibleTextInterface) SetCursorPosition(position int) {
	purego.SyscallN(addr_QAccessibleTextInterface_setCursorPosition, this.h, uintptr(position))
	runtime.KeepAlive(this)
}

func (this *QAccessibleTextInterface) Text(startOffset int, endOffset int) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAccessibleTextInterface_text, this.h, uintptr(startOffset), uintptr(endOffset), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QAccessibleTextInterface) TextBeforeOffset(offset int, boundaryType QAccessible__TextBoundaryType) (string, int, int) {
	var startOffset_out int32
	var endOffset_out int32
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAccessibleTextInterface_textBeforeOffset, this.h, uintptr(offset), uintptr(boundaryType), uintptr(unsafe.Pointer(&startOffset_out)), uintptr(unsafe.Pointer(&endOffset_out)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n), int(startOffset_out), int(endOffset_out)
}

func (this *QAccessibleTextInterface) TextAfterOffset(offset int, boundaryType QAccessible__TextBoundaryType) (string, int, int) {
	var startOffset_out int32
	var endOffset_out int32
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAccessibleTextInterface_textAfterOffset, this.h, uintptr(offset), uintptr(boundaryType), uintptr(unsafe.Pointer(&startOffset_out)), uintptr(unsafe.Pointer(&endOffset_out)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n), int(startOffset_out), int(endOffset_out)
}

func (this *QAccessibleTextInterface) TextAtOffset(offset int, boundaryType QAccessible__TextBoundaryType) (string, int, int) {
	var startOffset_out int32
	var endOffset_out int32
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAccessibleTextInterface_textAtOffset, this.h, uintptr(offset), uintptr(boundaryType), uintptr(unsafe.Pointer(&startOffset_out)), uintptr(unsafe.Pointer(&endOffset_out)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n), int(startOffset_out), int(endOffset_out)
}

func (this *QAccessibleTextInterface) CharacterCount() int {
	__r0, _, _ := purego.SyscallN(addr_QAccessibleTextInterface_characterCount, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAccessibleTextInterface) CharacterRect(offset int) *QRect {
	__r, _, _ := purego.SyscallN(addr_QAccessibleTextInterface_characterRect, this.h, uintptr(offset))
	runtime.KeepAlive(this)
	return newownedQRect(__r)
}

func (this *QAccessibleTextInterface) OffsetAtPoint(point *QPoint) int {
	__r0, _, _ := purego.SyscallN(addr_QAccessibleTextInterface_offsetAtPoint, this.h, uintptr(point.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(point)
	return int(__r)
}

func (this *QAccessibleTextInterface) ScrollToSubstring(startIndex int, endIndex int) {
	purego.SyscallN(addr_QAccessibleTextInterface_scrollToSubstring, this.h, uintptr(startIndex), uintptr(endIndex))
	runtime.KeepAlive(this)
}

func (this *QAccessibleTextInterface) Attributes(offset int) (string, int, int) {
	var startOffset_out int32
	var endOffset_out int32
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAccessibleTextInterface_attributes, this.h, uintptr(offset), uintptr(unsafe.Pointer(&startOffset_out)), uintptr(unsafe.Pointer(&endOffset_out)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n), int(startOffset_out), int(endOffset_out)
}

func (this *QAccessibleTextInterface) OperatorAssign(arg1 *QAccessibleTextInterface) *QAccessibleTextInterface {
	__r, _, _ := purego.SyscallN(addr_QAccessibleTextInterface_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQAccessibleTextInterface(__r)
}

func (this *QAccessibleTextInterface) TextBeforeOffsetDefault(offset int, boundaryType QAccessible__TextBoundaryType) (string, int, int) {
	var startOffset_out int32
	var endOffset_out int32
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAccessibleTextInterface_textBeforeOffset_default, this.h, uintptr(offset), uintptr(boundaryType), uintptr(unsafe.Pointer(&startOffset_out)), uintptr(unsafe.Pointer(&endOffset_out)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n), int(startOffset_out), int(endOffset_out)
}

func (this *QAccessibleTextInterface) TextAfterOffsetDefault(offset int, boundaryType QAccessible__TextBoundaryType) (string, int, int) {
	var startOffset_out int32
	var endOffset_out int32
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAccessibleTextInterface_textAfterOffset_default, this.h, uintptr(offset), uintptr(boundaryType), uintptr(unsafe.Pointer(&startOffset_out)), uintptr(unsafe.Pointer(&endOffset_out)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n), int(startOffset_out), int(endOffset_out)
}

func (this *QAccessibleTextInterface) TextAtOffsetDefault(offset int, boundaryType QAccessible__TextBoundaryType) (string, int, int) {
	var startOffset_out int32
	var endOffset_out int32
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAccessibleTextInterface_textAtOffset_default, this.h, uintptr(offset), uintptr(boundaryType), uintptr(unsafe.Pointer(&startOffset_out)), uintptr(unsafe.Pointer(&endOffset_out)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n), int(startOffset_out), int(endOffset_out)
}

func QAccessibleTextInterface_virtIdx(slot int64) int32 {
	switch slot {
	case 5156601901670015669:
		return 0
	case -4199621612085967932:
		return 1
	case 1701418371756208442:
		return 2
	case 4093222286646031805:
		return 3
	case -1539266838375037305:
		return 4
	case -4273823765597644148:
		return 5
	case -6512451105733458522:
		return 6
	case -430950356557152130:
		return 7
	case -6653538151610911041:
		return 8
	case 3334036117981237742:
		return 9
	case -4703424646947788745:
		return 10
	case 6934233772508985160:
		return 11
	case 6181953782893286932:
		return 12
	case 2244918558822102534:
		return 13
	case 3240586353052927855:
		return 14
	case -6929277707497019272:
		return 15
	}
	return -1
}

func QAccessibleTextInterface_enable(h uintptr, idx int32) { pg_QAccessibleTextInterface__set_override(h, idx) }

func NewQAccessibleTextInterface() *QAccessibleTextInterface {
	id := nextVirtID()
	o := newownedQAccessibleTextInterface(pg_QAccessibleTextInterface__Go_new(id))
	o.UnsafeSetGoID(id)
	registerInstance(id, o.h, QAccessibleTextInterface_virtIdx, QAccessibleTextInterface_enable)
	return o
}

func (this *QAccessibleTextInterface) OnSelection(fn func(int, *int32, *int32)) {
	registerVirt(this.UnsafeGoID(), 5156601901670015669, func(selectionIndex int32, startOffset uintptr, endOffset uintptr) { fn(int(selectionIndex), (*int32)(unsafe.Pointer(startOffset)), (*int32)(unsafe.Pointer(endOffset))) }) // slot "selection"
}

func (this *QAccessibleTextInterface) OnSelectionCount(fn func() int) {
	registerVirt(this.UnsafeGoID(), -4199621612085967932, func() int32 { return int32(fn()) }) // slot "selectionCount"
}

func (this *QAccessibleTextInterface) OnAddSelection(fn func(int, int)) {
	registerVirt(this.UnsafeGoID(), 1701418371756208442, func(startOffset int32, endOffset int32) { fn(int(startOffset), int(endOffset)) }) // slot "addSelection"
}

func (this *QAccessibleTextInterface) OnRemoveSelection(fn func(int)) {
	registerVirt(this.UnsafeGoID(), 4093222286646031805, func(selectionIndex int32) { fn(int(selectionIndex)) }) // slot "removeSelection"
}

func (this *QAccessibleTextInterface) OnSetSelection(fn func(int, int, int)) {
	registerVirt(this.UnsafeGoID(), -1539266838375037305, func(selectionIndex int32, startOffset int32, endOffset int32) { fn(int(selectionIndex), int(startOffset), int(endOffset)) }) // slot "setSelection"
}

func (this *QAccessibleTextInterface) OnCursorPosition(fn func() int) {
	registerVirt(this.UnsafeGoID(), -4273823765597644148, func() int32 { return int32(fn()) }) // slot "cursorPosition"
}

func (this *QAccessibleTextInterface) OnSetCursorPosition(fn func(int)) {
	registerVirt(this.UnsafeGoID(), -6512451105733458522, func(position int32) { fn(int(position)) }) // slot "setCursorPosition"
}

func (this *QAccessibleTextInterface) OnText(fn func(int, int) string) {
	registerVirt(this.UnsafeGoID(), -430950356557152130, func(startOffset int32, endOffset int32) (uintptr, int64) {
		__s := fn(int(startOffset), int(endOffset))
		return pgCString(__s), int64(len(__s))
	}) // slot "text"
}

func (this *QAccessibleTextInterface) OnCharacterCount(fn func() int) {
	registerVirt(this.UnsafeGoID(), -6653538151610911041, func() int32 { return int32(fn()) }) // slot "characterCount"
}

func (this *QAccessibleTextInterface) OnCharacterRect(fn func(int) *QRect) {
	registerVirt(this.UnsafeGoID(), 3334036117981237742, func(offset int32) uintptr { return uintptr(fn(int(offset)).UnsafePointer()) }) // slot "characterRect"
}

func (this *QAccessibleTextInterface) OnOffsetAtPoint(fn func(*QPoint) int) {
	registerVirt(this.UnsafeGoID(), -4703424646947788745, func(point uintptr) int32 { return int32(fn(UnsafeNewQPoint(unsafe.Pointer(point)))) }) // slot "offsetAtPoint"
}

func (this *QAccessibleTextInterface) OnScrollToSubstring(fn func(int, int)) {
	registerVirt(this.UnsafeGoID(), 6934233772508985160, func(startIndex int32, endIndex int32) { fn(int(startIndex), int(endIndex)) }) // slot "scrollToSubstring"
}

func (this *QAccessibleTextInterface) OnAttributes(fn func(int, *int32, *int32) string) {
	registerVirt(this.UnsafeGoID(), 6181953782893286932, func(offset int32, startOffset uintptr, endOffset uintptr) (uintptr, int64) {
		__s := fn(int(offset), (*int32)(unsafe.Pointer(startOffset)), (*int32)(unsafe.Pointer(endOffset)))
		return pgCString(__s), int64(len(__s))
	}) // slot "attributes"
}

func (this *QAccessibleTextInterface) OnTextBeforeOffset(fn func(int, QAccessible__TextBoundaryType, *int32, *int32) string) {
	registerVirt(this.UnsafeGoID(), 2244918558822102534, func(offset int32, boundaryType int32, startOffset uintptr, endOffset uintptr) (uintptr, int64) {
		__s := fn(int(offset), QAccessible__TextBoundaryType(boundaryType), (*int32)(unsafe.Pointer(startOffset)), (*int32)(unsafe.Pointer(endOffset)))
		return pgCString(__s), int64(len(__s))
	}) // slot "textBeforeOffset"
}

func (this *QAccessibleTextInterface) OnTextAfterOffset(fn func(int, QAccessible__TextBoundaryType, *int32, *int32) string) {
	registerVirt(this.UnsafeGoID(), 3240586353052927855, func(offset int32, boundaryType int32, startOffset uintptr, endOffset uintptr) (uintptr, int64) {
		__s := fn(int(offset), QAccessible__TextBoundaryType(boundaryType), (*int32)(unsafe.Pointer(startOffset)), (*int32)(unsafe.Pointer(endOffset)))
		return pgCString(__s), int64(len(__s))
	}) // slot "textAfterOffset"
}

func (this *QAccessibleTextInterface) OnTextAtOffset(fn func(int, QAccessible__TextBoundaryType, *int32, *int32) string) {
	registerVirt(this.UnsafeGoID(), -6929277707497019272, func(offset int32, boundaryType int32, startOffset uintptr, endOffset uintptr) (uintptr, int64) {
		__s := fn(int(offset), QAccessible__TextBoundaryType(boundaryType), (*int32)(unsafe.Pointer(startOffset)), (*int32)(unsafe.Pointer(endOffset)))
		return pgCString(__s), int64(len(__s))
	}) // slot "textAtOffset"
}

func (this *QAccessibleTextInterface) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QAccessibleTextInterface_delete(this.h)
	}
	this.UnsafeRelease()
}
