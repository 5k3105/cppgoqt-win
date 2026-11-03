package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QTextTableCell_new uintptr
	addr_QTextTableCell_new2 uintptr
	addr_QTextTableCell_OperatorAssign uintptr
	addr_QTextTableCell_setFormat uintptr
	addr_QTextTableCell_format uintptr
	addr_QTextTableCell_row uintptr
	addr_QTextTableCell_column uintptr
	addr_QTextTableCell_rowSpan uintptr
	addr_QTextTableCell_columnSpan uintptr
	addr_QTextTableCell_isValid uintptr
	addr_QTextTableCell_firstCursorPosition uintptr
	addr_QTextTableCell_lastCursorPosition uintptr
	addr_QTextTableCell_firstPosition uintptr
	addr_QTextTableCell_lastPosition uintptr
	addr_QTextTableCell_OperatorEqual uintptr
	addr_QTextTableCell_OperatorNotEqual uintptr
	addr_QTextTableCell_tableCellFormatIndex uintptr
	pg_QTextTableCell_delete func(uintptr)
)

func init() {
	symaddr(&addr_QTextTableCell_new, "QTextTableCell_new")
	symaddr(&addr_QTextTableCell_new2, "QTextTableCell_new2")
	symaddr(&addr_QTextTableCell_OperatorAssign, "QTextTableCell_OperatorAssign")
	symaddr(&addr_QTextTableCell_setFormat, "QTextTableCell_setFormat")
	symaddr(&addr_QTextTableCell_format, "QTextTableCell_format")
	symaddr(&addr_QTextTableCell_row, "QTextTableCell_row")
	symaddr(&addr_QTextTableCell_column, "QTextTableCell_column")
	symaddr(&addr_QTextTableCell_rowSpan, "QTextTableCell_rowSpan")
	symaddr(&addr_QTextTableCell_columnSpan, "QTextTableCell_columnSpan")
	symaddr(&addr_QTextTableCell_isValid, "QTextTableCell_isValid")
	symaddr(&addr_QTextTableCell_firstCursorPosition, "QTextTableCell_firstCursorPosition")
	symaddr(&addr_QTextTableCell_lastCursorPosition, "QTextTableCell_lastCursorPosition")
	symaddr(&addr_QTextTableCell_firstPosition, "QTextTableCell_firstPosition")
	symaddr(&addr_QTextTableCell_lastPosition, "QTextTableCell_lastPosition")
	symaddr(&addr_QTextTableCell_OperatorEqual, "QTextTableCell_OperatorEqual")
	symaddr(&addr_QTextTableCell_OperatorNotEqual, "QTextTableCell_OperatorNotEqual")
	symaddr(&addr_QTextTableCell_tableCellFormatIndex, "QTextTableCell_tableCellFormatIndex")
	reg(&pg_QTextTableCell_delete, "QTextTableCell_delete")
}

type QTextTableCell struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QTextTableCell) UnsafeGoID() int64     { return this.vid }
func (this *QTextTableCell) UnsafeSetGoID(id int64) { this.vid = id }

func bindQTextTableCell(h uintptr) *QTextTableCell {
	if h == 0 {
		return nil
	}
	return &QTextTableCell{h: h}
}

func newownedQTextTableCell(h uintptr) *QTextTableCell {
	o := bindQTextTableCell(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QTextTableCell).Delete)
	}
	return o
}

func (this *QTextTableCell) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QTextTableCell) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQTextTableCell(h unsafe.Pointer) *QTextTableCell { return bindQTextTableCell(uintptr(h)) }
func UnsafeNewQTextTableCellOwned(h unsafe.Pointer) *QTextTableCell { return newownedQTextTableCell(uintptr(h)) }

func (this *QTextTableCell) UnsafeRelease() { this.h = 0 }

func NewQTextTableCell() *QTextTableCell {
	__r, _, _ := purego.SyscallN(addr_QTextTableCell_new)
	__o := newownedQTextTableCell(__r)
	return __o
}

func NewQTextTableCell2(o *QTextTableCell) *QTextTableCell {
	__r, _, _ := purego.SyscallN(addr_QTextTableCell_new2, uintptr(o.UnsafePointer()))
	__o := newownedQTextTableCell(__r)
	runtime.KeepAlive(o)
	return __o
}

func (this *QTextTableCell) OperatorAssign(o *QTextTableCell) *QTextTableCell {
	__r, _, _ := purego.SyscallN(addr_QTextTableCell_OperatorAssign, this.h, uintptr(o.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(o)
	return bindQTextTableCell(__r)
}

func (this *QTextTableCell) SetFormat(format *QTextCharFormat) {
	purego.SyscallN(addr_QTextTableCell_setFormat, this.h, uintptr(format.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(format)
}

func (this *QTextTableCell) Format() *QTextCharFormat {
	__r, _, _ := purego.SyscallN(addr_QTextTableCell_format, this.h)
	runtime.KeepAlive(this)
	return newownedQTextCharFormat(__r)
}

func (this *QTextTableCell) Row() int {
	__r0, _, _ := purego.SyscallN(addr_QTextTableCell_row, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTextTableCell) Column() int {
	__r0, _, _ := purego.SyscallN(addr_QTextTableCell_column, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTextTableCell) RowSpan() int {
	__r0, _, _ := purego.SyscallN(addr_QTextTableCell_rowSpan, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTextTableCell) ColumnSpan() int {
	__r0, _, _ := purego.SyscallN(addr_QTextTableCell_columnSpan, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTextTableCell) IsValid() bool {
	__r0, _, _ := purego.SyscallN(addr_QTextTableCell_isValid, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTextTableCell) FirstCursorPosition() *QTextCursor {
	__r, _, _ := purego.SyscallN(addr_QTextTableCell_firstCursorPosition, this.h)
	runtime.KeepAlive(this)
	return newownedQTextCursor(__r)
}

func (this *QTextTableCell) LastCursorPosition() *QTextCursor {
	__r, _, _ := purego.SyscallN(addr_QTextTableCell_lastCursorPosition, this.h)
	runtime.KeepAlive(this)
	return newownedQTextCursor(__r)
}

func (this *QTextTableCell) FirstPosition() int {
	__r0, _, _ := purego.SyscallN(addr_QTextTableCell_firstPosition, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTextTableCell) LastPosition() int {
	__r0, _, _ := purego.SyscallN(addr_QTextTableCell_lastPosition, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTextTableCell) OperatorEqual(other *QTextTableCell) bool {
	__r0, _, _ := purego.SyscallN(addr_QTextTableCell_OperatorEqual, this.h, uintptr(other.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return __r != 0
}

func (this *QTextTableCell) OperatorNotEqual(other *QTextTableCell) bool {
	__r0, _, _ := purego.SyscallN(addr_QTextTableCell_OperatorNotEqual, this.h, uintptr(other.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return __r != 0
}

func (this *QTextTableCell) TableCellFormatIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QTextTableCell_tableCellFormatIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTextTableCell) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QTextTableCell_delete(this.h)
	}
	this.UnsafeRelease()
}
