package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QTextTableFormat_chain uintptr
	addr_QTextTableFormat_new uintptr
	addr_QTextTableFormat_new2 uintptr
	addr_QTextTableFormat_isValid uintptr
	addr_QTextTableFormat_columns uintptr
	addr_QTextTableFormat_setColumns uintptr
	addr_QTextTableFormat_setColumnWidthConstraints uintptr
	addr_QTextTableFormat_columnWidthConstraints uintptr
	addr_QTextTableFormat_clearColumnWidthConstraints uintptr
	pg_QTextTableFormat_cellSpacing func(uintptr) float64
	pg_QTextTableFormat_setCellSpacing func(uintptr, float64)
	pg_QTextTableFormat_cellPadding func(uintptr) float64
	pg_QTextTableFormat_setCellPadding func(uintptr, float64)
	addr_QTextTableFormat_setAlignment uintptr
	addr_QTextTableFormat_alignment uintptr
	addr_QTextTableFormat_setHeaderRowCount uintptr
	addr_QTextTableFormat_headerRowCount uintptr
	addr_QTextTableFormat_setBorderCollapse uintptr
	addr_QTextTableFormat_borderCollapse uintptr
	pg_QTextTableFormat_delete func(uintptr)
)

func init() {
	symaddr(&addr_QTextTableFormat_chain, "QTextTableFormat_chain")
	symaddr(&addr_QTextTableFormat_new, "QTextTableFormat_new")
	symaddr(&addr_QTextTableFormat_new2, "QTextTableFormat_new2")
	symaddr(&addr_QTextTableFormat_isValid, "QTextTableFormat_isValid")
	symaddr(&addr_QTextTableFormat_columns, "QTextTableFormat_columns")
	symaddr(&addr_QTextTableFormat_setColumns, "QTextTableFormat_setColumns")
	symaddr(&addr_QTextTableFormat_setColumnWidthConstraints, "QTextTableFormat_setColumnWidthConstraints")
	symaddr(&addr_QTextTableFormat_columnWidthConstraints, "QTextTableFormat_columnWidthConstraints")
	symaddr(&addr_QTextTableFormat_clearColumnWidthConstraints, "QTextTableFormat_clearColumnWidthConstraints")
	reg(&pg_QTextTableFormat_cellSpacing, "QTextTableFormat_cellSpacing")
	reg(&pg_QTextTableFormat_setCellSpacing, "QTextTableFormat_setCellSpacing")
	reg(&pg_QTextTableFormat_cellPadding, "QTextTableFormat_cellPadding")
	reg(&pg_QTextTableFormat_setCellPadding, "QTextTableFormat_setCellPadding")
	symaddr(&addr_QTextTableFormat_setAlignment, "QTextTableFormat_setAlignment")
	symaddr(&addr_QTextTableFormat_alignment, "QTextTableFormat_alignment")
	symaddr(&addr_QTextTableFormat_setHeaderRowCount, "QTextTableFormat_setHeaderRowCount")
	symaddr(&addr_QTextTableFormat_headerRowCount, "QTextTableFormat_headerRowCount")
	symaddr(&addr_QTextTableFormat_setBorderCollapse, "QTextTableFormat_setBorderCollapse")
	symaddr(&addr_QTextTableFormat_borderCollapse, "QTextTableFormat_borderCollapse")
	reg(&pg_QTextTableFormat_delete, "QTextTableFormat_delete")
}

type QTextTableFormat struct {
	h uintptr
	*QTextFrameFormat
}

func bindQTextTableFormat(h uintptr) *QTextTableFormat {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QTextTableFormat_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QTextTableFormat{h: h, QTextFrameFormat: &QTextFrameFormat{h: __c[0], QTextFormat: &QTextFormat{h: __c[1]}}}
}

func newownedQTextTableFormat(h uintptr) *QTextTableFormat {
	o := bindQTextTableFormat(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QTextTableFormat).Delete)
	}
	return o
}

func (this *QTextTableFormat) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QTextTableFormat) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQTextTableFormat(h unsafe.Pointer) *QTextTableFormat { return bindQTextTableFormat(uintptr(h)) }
func UnsafeNewQTextTableFormatOwned(h unsafe.Pointer) *QTextTableFormat { return newownedQTextTableFormat(uintptr(h)) }

func (this *QTextTableFormat) UnsafeRelease() {
	this.h = 0
	if this.QTextFrameFormat != nil {
		this.QTextFrameFormat.UnsafeRelease()
	}
}

func NewQTextTableFormat() *QTextTableFormat {
	__r, _, _ := purego.SyscallN(addr_QTextTableFormat_new)
	__o := newownedQTextTableFormat(__r)
	return __o
}

func NewQTextTableFormat2(arg1 *QTextTableFormat) *QTextTableFormat {
	__r, _, _ := purego.SyscallN(addr_QTextTableFormat_new2, uintptr(arg1.UnsafePointer()))
	__o := newownedQTextTableFormat(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QTextTableFormat) IsValid() bool {
	__r0, _, _ := purego.SyscallN(addr_QTextTableFormat_isValid, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTextTableFormat) Columns() int {
	__r0, _, _ := purego.SyscallN(addr_QTextTableFormat_columns, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTextTableFormat) SetColumns(columns int) {
	purego.SyscallN(addr_QTextTableFormat_setColumns, this.h, uintptr(columns))
	runtime.KeepAlive(this)
}

func (this *QTextTableFormat) SetColumnWidthConstraints(constraints []*QTextLength) {
	constraints_arr := make([]uintptr, len(constraints))
	for __i, __e := range constraints {
		constraints_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var constraints_p uintptr
	if len(constraints_arr) > 0 {
		constraints_p = uintptr(unsafe.Pointer(&constraints_arr[0]))
	}
	purego.SyscallN(addr_QTextTableFormat_setColumnWidthConstraints, this.h, constraints_p, uintptr(len(constraints)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(constraints)
	runtime.KeepAlive(constraints_arr)
}

func (this *QTextTableFormat) ColumnWidthConstraints() []*QTextLength {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTextTableFormat_columnWidthConstraints, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QTextLength, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQTextLength(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QTextTableFormat) ClearColumnWidthConstraints() {
	purego.SyscallN(addr_QTextTableFormat_clearColumnWidthConstraints, this.h)
	runtime.KeepAlive(this)
}

func (this *QTextTableFormat) CellSpacing() float64 {
	__r := pg_QTextTableFormat_cellSpacing(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QTextTableFormat) SetCellSpacing(spacing float64) {
	pg_QTextTableFormat_setCellSpacing(this.h, spacing)
	runtime.KeepAlive(this)
}

func (this *QTextTableFormat) CellPadding() float64 {
	__r := pg_QTextTableFormat_cellPadding(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QTextTableFormat) SetCellPadding(padding float64) {
	pg_QTextTableFormat_setCellPadding(this.h, padding)
	runtime.KeepAlive(this)
}

func (this *QTextTableFormat) SetAlignment(alignment AlignmentFlag) {
	purego.SyscallN(addr_QTextTableFormat_setAlignment, this.h, uintptr(alignment))
	runtime.KeepAlive(this)
}

func (this *QTextTableFormat) Alignment() AlignmentFlag {
	__r0, _, _ := purego.SyscallN(addr_QTextTableFormat_alignment, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return AlignmentFlag(__r)
}

func (this *QTextTableFormat) SetHeaderRowCount(count int) {
	purego.SyscallN(addr_QTextTableFormat_setHeaderRowCount, this.h, uintptr(count))
	runtime.KeepAlive(this)
}

func (this *QTextTableFormat) HeaderRowCount() int {
	__r0, _, _ := purego.SyscallN(addr_QTextTableFormat_headerRowCount, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTextTableFormat) SetBorderCollapse(borderCollapse bool) {
	purego.SyscallN(addr_QTextTableFormat_setBorderCollapse, this.h, uintptr(b2i(borderCollapse)))
	runtime.KeepAlive(this)
}

func (this *QTextTableFormat) BorderCollapse() bool {
	__r0, _, _ := purego.SyscallN(addr_QTextTableFormat_borderCollapse, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTextTableFormat) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QTextTableFormat_delete(this.h)
	}
	this.UnsafeRelease()
}
