package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QTableWidgetSelectionRange_new uintptr
	addr_QTableWidgetSelectionRange_new2 uintptr
	addr_QTableWidgetSelectionRange_topRow uintptr
	addr_QTableWidgetSelectionRange_bottomRow uintptr
	addr_QTableWidgetSelectionRange_leftColumn uintptr
	addr_QTableWidgetSelectionRange_rightColumn uintptr
	addr_QTableWidgetSelectionRange_rowCount uintptr
	addr_QTableWidgetSelectionRange_columnCount uintptr
	addr_QTableWidgetSelectionRange_OperatorEqual uintptr
	addr_QTableWidgetSelectionRange_OperatorNotEqual uintptr
	pg_QTableWidgetSelectionRange_delete func(uintptr)
)

func init() {
	symaddr(&addr_QTableWidgetSelectionRange_new, "QTableWidgetSelectionRange_new")
	symaddr(&addr_QTableWidgetSelectionRange_new2, "QTableWidgetSelectionRange_new2")
	symaddr(&addr_QTableWidgetSelectionRange_topRow, "QTableWidgetSelectionRange_topRow")
	symaddr(&addr_QTableWidgetSelectionRange_bottomRow, "QTableWidgetSelectionRange_bottomRow")
	symaddr(&addr_QTableWidgetSelectionRange_leftColumn, "QTableWidgetSelectionRange_leftColumn")
	symaddr(&addr_QTableWidgetSelectionRange_rightColumn, "QTableWidgetSelectionRange_rightColumn")
	symaddr(&addr_QTableWidgetSelectionRange_rowCount, "QTableWidgetSelectionRange_rowCount")
	symaddr(&addr_QTableWidgetSelectionRange_columnCount, "QTableWidgetSelectionRange_columnCount")
	symaddr(&addr_QTableWidgetSelectionRange_OperatorEqual, "QTableWidgetSelectionRange_OperatorEqual")
	symaddr(&addr_QTableWidgetSelectionRange_OperatorNotEqual, "QTableWidgetSelectionRange_OperatorNotEqual")
	reg(&pg_QTableWidgetSelectionRange_delete, "QTableWidgetSelectionRange_delete")
}

type QTableWidgetSelectionRange struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QTableWidgetSelectionRange) UnsafeGoID() int64     { return this.vid }
func (this *QTableWidgetSelectionRange) UnsafeSetGoID(id int64) { this.vid = id }

func bindQTableWidgetSelectionRange(h uintptr) *QTableWidgetSelectionRange {
	if h == 0 {
		return nil
	}
	return &QTableWidgetSelectionRange{h: h}
}

func newownedQTableWidgetSelectionRange(h uintptr) *QTableWidgetSelectionRange {
	o := bindQTableWidgetSelectionRange(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QTableWidgetSelectionRange).Delete)
	}
	return o
}

func (this *QTableWidgetSelectionRange) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QTableWidgetSelectionRange) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQTableWidgetSelectionRange(h unsafe.Pointer) *QTableWidgetSelectionRange { return bindQTableWidgetSelectionRange(uintptr(h)) }
func UnsafeNewQTableWidgetSelectionRangeOwned(h unsafe.Pointer) *QTableWidgetSelectionRange { return newownedQTableWidgetSelectionRange(uintptr(h)) }

func (this *QTableWidgetSelectionRange) UnsafeRelease() { this.h = 0 }

func NewQTableWidgetSelectionRange() *QTableWidgetSelectionRange {
	__r, _, _ := purego.SyscallN(addr_QTableWidgetSelectionRange_new)
	__o := newownedQTableWidgetSelectionRange(__r)
	return __o
}

func NewQTableWidgetSelectionRange2(top int, left int, bottom int, right int) *QTableWidgetSelectionRange {
	__r, _, _ := purego.SyscallN(addr_QTableWidgetSelectionRange_new2, uintptr(top), uintptr(left), uintptr(bottom), uintptr(right))
	__o := newownedQTableWidgetSelectionRange(__r)
	return __o
}

func (this *QTableWidgetSelectionRange) TopRow() int {
	__r0, _, _ := purego.SyscallN(addr_QTableWidgetSelectionRange_topRow, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTableWidgetSelectionRange) BottomRow() int {
	__r0, _, _ := purego.SyscallN(addr_QTableWidgetSelectionRange_bottomRow, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTableWidgetSelectionRange) LeftColumn() int {
	__r0, _, _ := purego.SyscallN(addr_QTableWidgetSelectionRange_leftColumn, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTableWidgetSelectionRange) RightColumn() int {
	__r0, _, _ := purego.SyscallN(addr_QTableWidgetSelectionRange_rightColumn, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTableWidgetSelectionRange) RowCount() int {
	__r0, _, _ := purego.SyscallN(addr_QTableWidgetSelectionRange_rowCount, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTableWidgetSelectionRange) ColumnCount() int {
	__r0, _, _ := purego.SyscallN(addr_QTableWidgetSelectionRange_columnCount, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTableWidgetSelectionRange) OperatorEqual(rhs *QTableWidgetSelectionRange) bool {
	__r0, _, _ := purego.SyscallN(addr_QTableWidgetSelectionRange_OperatorEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QTableWidgetSelectionRange) OperatorNotEqual(rhs *QTableWidgetSelectionRange) bool {
	__r0, _, _ := purego.SyscallN(addr_QTableWidgetSelectionRange_OperatorNotEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QTableWidgetSelectionRange) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QTableWidgetSelectionRange_delete(this.h)
	}
	this.UnsafeRelease()
}
