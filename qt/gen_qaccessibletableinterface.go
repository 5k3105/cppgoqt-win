package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QAccessibleTableInterface_caption uintptr
	addr_QAccessibleTableInterface_summary uintptr
	addr_QAccessibleTableInterface_cellAt uintptr
	addr_QAccessibleTableInterface_selectedCellCount uintptr
	addr_QAccessibleTableInterface_selectedCells uintptr
	addr_QAccessibleTableInterface_columnDescription uintptr
	addr_QAccessibleTableInterface_rowDescription uintptr
	addr_QAccessibleTableInterface_selectedColumnCount uintptr
	addr_QAccessibleTableInterface_selectedRowCount uintptr
	addr_QAccessibleTableInterface_columnCount uintptr
	addr_QAccessibleTableInterface_rowCount uintptr
	addr_QAccessibleTableInterface_selectedColumns uintptr
	addr_QAccessibleTableInterface_selectedRows uintptr
	addr_QAccessibleTableInterface_isColumnSelected uintptr
	addr_QAccessibleTableInterface_isRowSelected uintptr
	addr_QAccessibleTableInterface_selectRow uintptr
	addr_QAccessibleTableInterface_selectColumn uintptr
	addr_QAccessibleTableInterface_unselectRow uintptr
	addr_QAccessibleTableInterface_unselectColumn uintptr
	addr_QAccessibleTableInterface_modelChange uintptr
	addr_QAccessibleTableInterface_OperatorAssign uintptr
	pg_QAccessibleTableInterface__set_override func(uintptr, int32)
	pg_QAccessibleTableInterface__Go_new func(int64) uintptr
	pg_QAccessibleTableInterface_delete func(uintptr)
)

func init() {
	symaddr(&addr_QAccessibleTableInterface_caption, "QAccessibleTableInterface_caption")
	symaddr(&addr_QAccessibleTableInterface_summary, "QAccessibleTableInterface_summary")
	symaddr(&addr_QAccessibleTableInterface_cellAt, "QAccessibleTableInterface_cellAt")
	symaddr(&addr_QAccessibleTableInterface_selectedCellCount, "QAccessibleTableInterface_selectedCellCount")
	symaddr(&addr_QAccessibleTableInterface_selectedCells, "QAccessibleTableInterface_selectedCells")
	symaddr(&addr_QAccessibleTableInterface_columnDescription, "QAccessibleTableInterface_columnDescription")
	symaddr(&addr_QAccessibleTableInterface_rowDescription, "QAccessibleTableInterface_rowDescription")
	symaddr(&addr_QAccessibleTableInterface_selectedColumnCount, "QAccessibleTableInterface_selectedColumnCount")
	symaddr(&addr_QAccessibleTableInterface_selectedRowCount, "QAccessibleTableInterface_selectedRowCount")
	symaddr(&addr_QAccessibleTableInterface_columnCount, "QAccessibleTableInterface_columnCount")
	symaddr(&addr_QAccessibleTableInterface_rowCount, "QAccessibleTableInterface_rowCount")
	symaddr(&addr_QAccessibleTableInterface_selectedColumns, "QAccessibleTableInterface_selectedColumns")
	symaddr(&addr_QAccessibleTableInterface_selectedRows, "QAccessibleTableInterface_selectedRows")
	symaddr(&addr_QAccessibleTableInterface_isColumnSelected, "QAccessibleTableInterface_isColumnSelected")
	symaddr(&addr_QAccessibleTableInterface_isRowSelected, "QAccessibleTableInterface_isRowSelected")
	symaddr(&addr_QAccessibleTableInterface_selectRow, "QAccessibleTableInterface_selectRow")
	symaddr(&addr_QAccessibleTableInterface_selectColumn, "QAccessibleTableInterface_selectColumn")
	symaddr(&addr_QAccessibleTableInterface_unselectRow, "QAccessibleTableInterface_unselectRow")
	symaddr(&addr_QAccessibleTableInterface_unselectColumn, "QAccessibleTableInterface_unselectColumn")
	symaddr(&addr_QAccessibleTableInterface_modelChange, "QAccessibleTableInterface_modelChange")
	symaddr(&addr_QAccessibleTableInterface_OperatorAssign, "QAccessibleTableInterface_OperatorAssign")
	reg(&pg_QAccessibleTableInterface__set_override, "QAccessibleTableInterface__set_override")
	reg(&pg_QAccessibleTableInterface__Go_new, "QAccessibleTableInterface__Go_new")
	reg(&pg_QAccessibleTableInterface_delete, "QAccessibleTableInterface_delete")
}

type QAccessibleTableInterface struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QAccessibleTableInterface) UnsafeGoID() int64     { return this.vid }
func (this *QAccessibleTableInterface) UnsafeSetGoID(id int64) { this.vid = id }

func bindQAccessibleTableInterface(h uintptr) *QAccessibleTableInterface {
	if h == 0 {
		return nil
	}
	return &QAccessibleTableInterface{h: h}
}

func newownedQAccessibleTableInterface(h uintptr) *QAccessibleTableInterface {
	o := bindQAccessibleTableInterface(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QAccessibleTableInterface).Delete)
	}
	return o
}

func (this *QAccessibleTableInterface) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QAccessibleTableInterface) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQAccessibleTableInterface(h unsafe.Pointer) *QAccessibleTableInterface { return bindQAccessibleTableInterface(uintptr(h)) }
func UnsafeNewQAccessibleTableInterfaceOwned(h unsafe.Pointer) *QAccessibleTableInterface { return newownedQAccessibleTableInterface(uintptr(h)) }

func (this *QAccessibleTableInterface) UnsafeRelease() { this.h = 0 }

func (this *QAccessibleTableInterface) Caption() *QAccessibleInterface {
	__r, _, _ := purego.SyscallN(addr_QAccessibleTableInterface_caption, this.h)
	runtime.KeepAlive(this)
	return bindQAccessibleInterface(__r)
}

func (this *QAccessibleTableInterface) Summary() *QAccessibleInterface {
	__r, _, _ := purego.SyscallN(addr_QAccessibleTableInterface_summary, this.h)
	runtime.KeepAlive(this)
	return bindQAccessibleInterface(__r)
}

func (this *QAccessibleTableInterface) CellAt(row int, column int) *QAccessibleInterface {
	__r, _, _ := purego.SyscallN(addr_QAccessibleTableInterface_cellAt, this.h, uintptr(row), uintptr(column))
	runtime.KeepAlive(this)
	return bindQAccessibleInterface(__r)
}

func (this *QAccessibleTableInterface) SelectedCellCount() int {
	__r0, _, _ := purego.SyscallN(addr_QAccessibleTableInterface_selectedCellCount, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAccessibleTableInterface) SelectedCells() []*QAccessibleInterface {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAccessibleTableInterface_selectedCells, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QAccessibleInterface, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = bindQAccessibleInterface(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QAccessibleTableInterface) ColumnDescription(column int) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAccessibleTableInterface_columnDescription, this.h, uintptr(column), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QAccessibleTableInterface) RowDescription(row int) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAccessibleTableInterface_rowDescription, this.h, uintptr(row), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QAccessibleTableInterface) SelectedColumnCount() int {
	__r0, _, _ := purego.SyscallN(addr_QAccessibleTableInterface_selectedColumnCount, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAccessibleTableInterface) SelectedRowCount() int {
	__r0, _, _ := purego.SyscallN(addr_QAccessibleTableInterface_selectedRowCount, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAccessibleTableInterface) ColumnCount() int {
	__r0, _, _ := purego.SyscallN(addr_QAccessibleTableInterface_columnCount, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAccessibleTableInterface) RowCount() int {
	__r0, _, _ := purego.SyscallN(addr_QAccessibleTableInterface_rowCount, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAccessibleTableInterface) SelectedColumns() []int {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAccessibleTableInterface_selectedColumns, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]int, __n)
	__sl := unsafe.Slice((*int32)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = int(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QAccessibleTableInterface) SelectedRows() []int {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAccessibleTableInterface_selectedRows, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]int, __n)
	__sl := unsafe.Slice((*int32)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = int(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QAccessibleTableInterface) IsColumnSelected(column int) bool {
	__r0, _, _ := purego.SyscallN(addr_QAccessibleTableInterface_isColumnSelected, this.h, uintptr(column))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QAccessibleTableInterface) IsRowSelected(row int) bool {
	__r0, _, _ := purego.SyscallN(addr_QAccessibleTableInterface_isRowSelected, this.h, uintptr(row))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QAccessibleTableInterface) SelectRow(row int) bool {
	__r0, _, _ := purego.SyscallN(addr_QAccessibleTableInterface_selectRow, this.h, uintptr(row))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QAccessibleTableInterface) SelectColumn(column int) bool {
	__r0, _, _ := purego.SyscallN(addr_QAccessibleTableInterface_selectColumn, this.h, uintptr(column))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QAccessibleTableInterface) UnselectRow(row int) bool {
	__r0, _, _ := purego.SyscallN(addr_QAccessibleTableInterface_unselectRow, this.h, uintptr(row))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QAccessibleTableInterface) UnselectColumn(column int) bool {
	__r0, _, _ := purego.SyscallN(addr_QAccessibleTableInterface_unselectColumn, this.h, uintptr(column))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QAccessibleTableInterface) ModelChange(event *QAccessibleTableModelChangeEvent) {
	purego.SyscallN(addr_QAccessibleTableInterface_modelChange, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAccessibleTableInterface) OperatorAssign(arg1 *QAccessibleTableInterface) *QAccessibleTableInterface {
	__r, _, _ := purego.SyscallN(addr_QAccessibleTableInterface_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQAccessibleTableInterface(__r)
}

func QAccessibleTableInterface_virtIdx(slot int64) int32 {
	switch slot {
	case 6749251983603663525:
		return 0
	case -6646608006195182317:
		return 1
	case 6973530125813990328:
		return 2
	case -1985908008123410713:
		return 3
	case -2099517733663994687:
		return 4
	case -4866595225830892309:
		return 5
	case -2563375781797925217:
		return 6
	case -4569576858401950571:
		return 7
	case -2566786583858001057:
		return 8
	case -8256963024669584834:
		return 9
	case 1960960171956884010:
		return 10
	case -278488305062277297:
		return 11
	case -5252052890467341575:
		return 12
	case 1772671393755888886:
		return 13
	case -8255619831741278870:
		return 14
	case 2033665621955588835:
		return 15
	case 3930440648437166287:
		return 16
	case -7465067139058793362:
		return 17
	case 5189646775272561424:
		return 18
	case -8156005454275955122:
		return 19
	}
	return -1
}

func QAccessibleTableInterface_enable(h uintptr, idx int32) { pg_QAccessibleTableInterface__set_override(h, idx) }

func NewQAccessibleTableInterface() *QAccessibleTableInterface {
	id := nextVirtID()
	o := newownedQAccessibleTableInterface(pg_QAccessibleTableInterface__Go_new(id))
	o.UnsafeSetGoID(id)
	registerInstance(id, o.h, QAccessibleTableInterface_virtIdx, QAccessibleTableInterface_enable)
	return o
}

func (this *QAccessibleTableInterface) OnCaption(fn func() *QAccessibleInterface) {
	registerVirt(this.UnsafeGoID(), 6749251983603663525, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "caption"
}

func (this *QAccessibleTableInterface) OnSummary(fn func() *QAccessibleInterface) {
	registerVirt(this.UnsafeGoID(), -6646608006195182317, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "summary"
}

func (this *QAccessibleTableInterface) OnCellAt(fn func(int, int) *QAccessibleInterface) {
	registerVirt(this.UnsafeGoID(), 6973530125813990328, func(row int32, column int32) uintptr { return uintptr(fn(int(row), int(column)).UnsafePointer()) }) // slot "cellAt"
}

func (this *QAccessibleTableInterface) OnSelectedCellCount(fn func() int) {
	registerVirt(this.UnsafeGoID(), -1985908008123410713, func() int32 { return int32(fn()) }) // slot "selectedCellCount"
}

func (this *QAccessibleTableInterface) OnSelectedCells(fn func() []*QAccessibleInterface) {
	registerVirt(this.UnsafeGoID(), -2099517733663994687, func() (uintptr, int64) {
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
	}) // slot "selectedCells"
}

func (this *QAccessibleTableInterface) OnColumnDescription(fn func(int) string) {
	registerVirt(this.UnsafeGoID(), -4866595225830892309, func(column int32) (uintptr, int64) {
		__s := fn(int(column))
		return pgCString(__s), int64(len(__s))
	}) // slot "columnDescription"
}

func (this *QAccessibleTableInterface) OnRowDescription(fn func(int) string) {
	registerVirt(this.UnsafeGoID(), -2563375781797925217, func(row int32) (uintptr, int64) {
		__s := fn(int(row))
		return pgCString(__s), int64(len(__s))
	}) // slot "rowDescription"
}

func (this *QAccessibleTableInterface) OnSelectedColumnCount(fn func() int) {
	registerVirt(this.UnsafeGoID(), -4569576858401950571, func() int32 { return int32(fn()) }) // slot "selectedColumnCount"
}

func (this *QAccessibleTableInterface) OnSelectedRowCount(fn func() int) {
	registerVirt(this.UnsafeGoID(), -2566786583858001057, func() int32 { return int32(fn()) }) // slot "selectedRowCount"
}

func (this *QAccessibleTableInterface) OnColumnCount(fn func() int) {
	registerVirt(this.UnsafeGoID(), -8256963024669584834, func() int32 { return int32(fn()) }) // slot "columnCount"
}

func (this *QAccessibleTableInterface) OnRowCount(fn func() int) {
	registerVirt(this.UnsafeGoID(), 1960960171956884010, func() int32 { return int32(fn()) }) // slot "rowCount"
}

func (this *QAccessibleTableInterface) OnSelectedColumns(fn func() []int) {
	registerVirt(this.UnsafeGoID(), -278488305062277297, func() (uintptr, int64) {
		__rs := fn()
		if len(__rs) == 0 {
			return 0, 0
		}
		__a := pgMalloc(uintptr(len(__rs)) * unsafe.Sizeof(int32(0)))
		__sl := unsafe.Slice((*int32)(unsafe.Pointer(__a)), len(__rs))
		for __i, __e := range __rs {
			__sl[__i] = int32(__e)
		}
		return __a, int64(len(__rs))
	}) // slot "selectedColumns"
}

func (this *QAccessibleTableInterface) OnSelectedRows(fn func() []int) {
	registerVirt(this.UnsafeGoID(), -5252052890467341575, func() (uintptr, int64) {
		__rs := fn()
		if len(__rs) == 0 {
			return 0, 0
		}
		__a := pgMalloc(uintptr(len(__rs)) * unsafe.Sizeof(int32(0)))
		__sl := unsafe.Slice((*int32)(unsafe.Pointer(__a)), len(__rs))
		for __i, __e := range __rs {
			__sl[__i] = int32(__e)
		}
		return __a, int64(len(__rs))
	}) // slot "selectedRows"
}

func (this *QAccessibleTableInterface) OnIsColumnSelected(fn func(int) bool) {
	registerVirt(this.UnsafeGoID(), 1772671393755888886, func(column int32) int32 { return b2i(fn(int(column))) }) // slot "isColumnSelected"
}

func (this *QAccessibleTableInterface) OnIsRowSelected(fn func(int) bool) {
	registerVirt(this.UnsafeGoID(), -8255619831741278870, func(row int32) int32 { return b2i(fn(int(row))) }) // slot "isRowSelected"
}

func (this *QAccessibleTableInterface) OnSelectRow(fn func(int) bool) {
	registerVirt(this.UnsafeGoID(), 2033665621955588835, func(row int32) int32 { return b2i(fn(int(row))) }) // slot "selectRow"
}

func (this *QAccessibleTableInterface) OnSelectColumn(fn func(int) bool) {
	registerVirt(this.UnsafeGoID(), 3930440648437166287, func(column int32) int32 { return b2i(fn(int(column))) }) // slot "selectColumn"
}

func (this *QAccessibleTableInterface) OnUnselectRow(fn func(int) bool) {
	registerVirt(this.UnsafeGoID(), -7465067139058793362, func(row int32) int32 { return b2i(fn(int(row))) }) // slot "unselectRow"
}

func (this *QAccessibleTableInterface) OnUnselectColumn(fn func(int) bool) {
	registerVirt(this.UnsafeGoID(), 5189646775272561424, func(column int32) int32 { return b2i(fn(int(column))) }) // slot "unselectColumn"
}

func (this *QAccessibleTableInterface) OnModelChange(fn func(*QAccessibleTableModelChangeEvent)) {
	registerVirt(this.UnsafeGoID(), -8156005454275955122, func(event uintptr) { fn(UnsafeNewQAccessibleTableModelChangeEvent(unsafe.Pointer(event))) }) // slot "modelChange"
}

func (this *QAccessibleTableInterface) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QAccessibleTableInterface_delete(this.h)
	}
	this.UnsafeRelease()
}
