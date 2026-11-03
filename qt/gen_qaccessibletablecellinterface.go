package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QAccessibleTableCellInterface_isSelected uintptr
	addr_QAccessibleTableCellInterface_columnHeaderCells uintptr
	addr_QAccessibleTableCellInterface_rowHeaderCells uintptr
	addr_QAccessibleTableCellInterface_columnIndex uintptr
	addr_QAccessibleTableCellInterface_rowIndex uintptr
	addr_QAccessibleTableCellInterface_columnExtent uintptr
	addr_QAccessibleTableCellInterface_rowExtent uintptr
	addr_QAccessibleTableCellInterface_table uintptr
	addr_QAccessibleTableCellInterface_OperatorAssign uintptr
	pg_QAccessibleTableCellInterface__set_override func(uintptr, int32)
	pg_QAccessibleTableCellInterface__Go_new func(int64) uintptr
	pg_QAccessibleTableCellInterface_delete func(uintptr)
)

func init() {
	symaddr(&addr_QAccessibleTableCellInterface_isSelected, "QAccessibleTableCellInterface_isSelected")
	symaddr(&addr_QAccessibleTableCellInterface_columnHeaderCells, "QAccessibleTableCellInterface_columnHeaderCells")
	symaddr(&addr_QAccessibleTableCellInterface_rowHeaderCells, "QAccessibleTableCellInterface_rowHeaderCells")
	symaddr(&addr_QAccessibleTableCellInterface_columnIndex, "QAccessibleTableCellInterface_columnIndex")
	symaddr(&addr_QAccessibleTableCellInterface_rowIndex, "QAccessibleTableCellInterface_rowIndex")
	symaddr(&addr_QAccessibleTableCellInterface_columnExtent, "QAccessibleTableCellInterface_columnExtent")
	symaddr(&addr_QAccessibleTableCellInterface_rowExtent, "QAccessibleTableCellInterface_rowExtent")
	symaddr(&addr_QAccessibleTableCellInterface_table, "QAccessibleTableCellInterface_table")
	symaddr(&addr_QAccessibleTableCellInterface_OperatorAssign, "QAccessibleTableCellInterface_OperatorAssign")
	reg(&pg_QAccessibleTableCellInterface__set_override, "QAccessibleTableCellInterface__set_override")
	reg(&pg_QAccessibleTableCellInterface__Go_new, "QAccessibleTableCellInterface__Go_new")
	reg(&pg_QAccessibleTableCellInterface_delete, "QAccessibleTableCellInterface_delete")
}

type QAccessibleTableCellInterface struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QAccessibleTableCellInterface) UnsafeGoID() int64     { return this.vid }
func (this *QAccessibleTableCellInterface) UnsafeSetGoID(id int64) { this.vid = id }

func bindQAccessibleTableCellInterface(h uintptr) *QAccessibleTableCellInterface {
	if h == 0 {
		return nil
	}
	return &QAccessibleTableCellInterface{h: h}
}

func newownedQAccessibleTableCellInterface(h uintptr) *QAccessibleTableCellInterface {
	o := bindQAccessibleTableCellInterface(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QAccessibleTableCellInterface).Delete)
	}
	return o
}

func (this *QAccessibleTableCellInterface) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QAccessibleTableCellInterface) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQAccessibleTableCellInterface(h unsafe.Pointer) *QAccessibleTableCellInterface { return bindQAccessibleTableCellInterface(uintptr(h)) }
func UnsafeNewQAccessibleTableCellInterfaceOwned(h unsafe.Pointer) *QAccessibleTableCellInterface { return newownedQAccessibleTableCellInterface(uintptr(h)) }

func (this *QAccessibleTableCellInterface) UnsafeRelease() { this.h = 0 }

func (this *QAccessibleTableCellInterface) IsSelected() bool {
	__r0, _, _ := purego.SyscallN(addr_QAccessibleTableCellInterface_isSelected, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QAccessibleTableCellInterface) ColumnHeaderCells() []*QAccessibleInterface {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAccessibleTableCellInterface_columnHeaderCells, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QAccessibleInterface, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = bindQAccessibleInterface(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QAccessibleTableCellInterface) RowHeaderCells() []*QAccessibleInterface {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAccessibleTableCellInterface_rowHeaderCells, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QAccessibleInterface, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = bindQAccessibleInterface(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QAccessibleTableCellInterface) ColumnIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QAccessibleTableCellInterface_columnIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAccessibleTableCellInterface) RowIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QAccessibleTableCellInterface_rowIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAccessibleTableCellInterface) ColumnExtent() int {
	__r0, _, _ := purego.SyscallN(addr_QAccessibleTableCellInterface_columnExtent, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAccessibleTableCellInterface) RowExtent() int {
	__r0, _, _ := purego.SyscallN(addr_QAccessibleTableCellInterface_rowExtent, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAccessibleTableCellInterface) Table() *QAccessibleInterface {
	__r, _, _ := purego.SyscallN(addr_QAccessibleTableCellInterface_table, this.h)
	runtime.KeepAlive(this)
	return bindQAccessibleInterface(__r)
}

func (this *QAccessibleTableCellInterface) OperatorAssign(arg1 *QAccessibleTableCellInterface) *QAccessibleTableCellInterface {
	__r, _, _ := purego.SyscallN(addr_QAccessibleTableCellInterface_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQAccessibleTableCellInterface(__r)
}

func QAccessibleTableCellInterface_virtIdx(slot int64) int32 {
	switch slot {
	case -3350427747592150876:
		return 0
	case 1398826710565483375:
		return 1
	case -2170165325896972357:
		return 2
	case -1856647841747687763:
		return 3
	case 5428465724232560177:
		return 4
	case -3228800890628578745:
		return 5
	case 6265358599971552931:
		return 6
	case 8583921542012250175:
		return 7
	}
	return -1
}

func QAccessibleTableCellInterface_enable(h uintptr, idx int32) { pg_QAccessibleTableCellInterface__set_override(h, idx) }

func NewQAccessibleTableCellInterface() *QAccessibleTableCellInterface {
	id := nextVirtID()
	o := newownedQAccessibleTableCellInterface(pg_QAccessibleTableCellInterface__Go_new(id))
	o.UnsafeSetGoID(id)
	registerInstance(id, o.h, QAccessibleTableCellInterface_virtIdx, QAccessibleTableCellInterface_enable)
	return o
}

func (this *QAccessibleTableCellInterface) OnIsSelected(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -3350427747592150876, func() int32 { return b2i(fn()) }) // slot "isSelected"
}

func (this *QAccessibleTableCellInterface) OnColumnHeaderCells(fn func() []*QAccessibleInterface) {
	registerVirt(this.UnsafeGoID(), 1398826710565483375, func() (uintptr, int64) {
		__rs := fn()
		if len(__rs) == 0 {
			return 0, 0
		}
		__a := pgMalloc(uintptr(len(__rs)) * unsafe.Sizeof(uintptr(0)))
		__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__a)), len(__rs))
		for __i, __e := range __rs {
			__sl[__i] = uintptr(__e.UnsafePointer())
		}
		return __a, int64(len(__rs))
	}) // slot "columnHeaderCells"
}

func (this *QAccessibleTableCellInterface) OnRowHeaderCells(fn func() []*QAccessibleInterface) {
	registerVirt(this.UnsafeGoID(), -2170165325896972357, func() (uintptr, int64) {
		__rs := fn()
		if len(__rs) == 0 {
			return 0, 0
		}
		__a := pgMalloc(uintptr(len(__rs)) * unsafe.Sizeof(uintptr(0)))
		__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__a)), len(__rs))
		for __i, __e := range __rs {
			__sl[__i] = uintptr(__e.UnsafePointer())
		}
		return __a, int64(len(__rs))
	}) // slot "rowHeaderCells"
}

func (this *QAccessibleTableCellInterface) OnColumnIndex(fn func() int) {
	registerVirt(this.UnsafeGoID(), -1856647841747687763, func() int32 { return int32(fn()) }) // slot "columnIndex"
}

func (this *QAccessibleTableCellInterface) OnRowIndex(fn func() int) {
	registerVirt(this.UnsafeGoID(), 5428465724232560177, func() int32 { return int32(fn()) }) // slot "rowIndex"
}

func (this *QAccessibleTableCellInterface) OnColumnExtent(fn func() int) {
	registerVirt(this.UnsafeGoID(), -3228800890628578745, func() int32 { return int32(fn()) }) // slot "columnExtent"
}

func (this *QAccessibleTableCellInterface) OnRowExtent(fn func() int) {
	registerVirt(this.UnsafeGoID(), 6265358599971552931, func() int32 { return int32(fn()) }) // slot "rowExtent"
}

func (this *QAccessibleTableCellInterface) OnTable(fn func() *QAccessibleInterface) {
	registerVirt(this.UnsafeGoID(), 8583921542012250175, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "table"
}

func (this *QAccessibleTableCellInterface) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QAccessibleTableCellInterface_delete(this.h)
	}
	this.UnsafeRelease()
}
