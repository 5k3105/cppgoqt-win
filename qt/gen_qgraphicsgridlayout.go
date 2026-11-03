package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QGraphicsGridLayout_chain uintptr
	addr_QGraphicsGridLayout_new uintptr
	addr_QGraphicsGridLayout_new2 uintptr
	addr_QGraphicsGridLayout_addItem uintptr
	addr_QGraphicsGridLayout_addItem2 uintptr
	pg_QGraphicsGridLayout_setHorizontalSpacing func(uintptr, float64)
	pg_QGraphicsGridLayout_horizontalSpacing func(uintptr) float64
	pg_QGraphicsGridLayout_setVerticalSpacing func(uintptr, float64)
	pg_QGraphicsGridLayout_verticalSpacing func(uintptr) float64
	pg_QGraphicsGridLayout_setSpacing func(uintptr, float64)
	pg_QGraphicsGridLayout_setRowSpacing func(uintptr, int32, float64)
	pg_QGraphicsGridLayout_rowSpacing func(uintptr, int32) float64
	pg_QGraphicsGridLayout_setColumnSpacing func(uintptr, int32, float64)
	pg_QGraphicsGridLayout_columnSpacing func(uintptr, int32) float64
	addr_QGraphicsGridLayout_setRowStretchFactor uintptr
	addr_QGraphicsGridLayout_rowStretchFactor uintptr
	addr_QGraphicsGridLayout_setColumnStretchFactor uintptr
	addr_QGraphicsGridLayout_columnStretchFactor uintptr
	pg_QGraphicsGridLayout_setRowMinimumHeight func(uintptr, int32, float64)
	pg_QGraphicsGridLayout_rowMinimumHeight func(uintptr, int32) float64
	pg_QGraphicsGridLayout_setRowPreferredHeight func(uintptr, int32, float64)
	pg_QGraphicsGridLayout_rowPreferredHeight func(uintptr, int32) float64
	pg_QGraphicsGridLayout_setRowMaximumHeight func(uintptr, int32, float64)
	pg_QGraphicsGridLayout_rowMaximumHeight func(uintptr, int32) float64
	pg_QGraphicsGridLayout_setRowFixedHeight func(uintptr, int32, float64)
	pg_QGraphicsGridLayout_setColumnMinimumWidth func(uintptr, int32, float64)
	pg_QGraphicsGridLayout_columnMinimumWidth func(uintptr, int32) float64
	pg_QGraphicsGridLayout_setColumnPreferredWidth func(uintptr, int32, float64)
	pg_QGraphicsGridLayout_columnPreferredWidth func(uintptr, int32) float64
	pg_QGraphicsGridLayout_setColumnMaximumWidth func(uintptr, int32, float64)
	pg_QGraphicsGridLayout_columnMaximumWidth func(uintptr, int32) float64
	pg_QGraphicsGridLayout_setColumnFixedWidth func(uintptr, int32, float64)
	addr_QGraphicsGridLayout_setRowAlignment uintptr
	addr_QGraphicsGridLayout_rowAlignment uintptr
	addr_QGraphicsGridLayout_setColumnAlignment uintptr
	addr_QGraphicsGridLayout_columnAlignment uintptr
	addr_QGraphicsGridLayout_setAlignment uintptr
	addr_QGraphicsGridLayout_alignment uintptr
	addr_QGraphicsGridLayout_rowCount uintptr
	addr_QGraphicsGridLayout_columnCount uintptr
	addr_QGraphicsGridLayout_itemAt uintptr
	addr_QGraphicsGridLayout_itemAt2 uintptr
	addr_QGraphicsGridLayout_removeItem uintptr
	addr_QGraphicsGridLayout_setInstantInvalidatePropagation uintptr
	addr_QGraphicsGridLayout_instantInvalidatePropagation uintptr
	addr_QGraphicsGridLayout_addChildLayoutItem uintptr
	addr_QGraphicsGridLayout_setGraphicsItem uintptr
	addr_QGraphicsGridLayout_setOwnedByLayout uintptr
	addr_QGraphicsGridLayout_addItem3 uintptr
	addr_QGraphicsGridLayout_addItem4 uintptr
	addr_QGraphicsGridLayout_itemAt_default uintptr
	addr_QGraphicsGridLayout_itemAt_default2 uintptr
	addr_QGraphicsGridLayout_count_default uintptr
	addr_QGraphicsGridLayout_removeAt_default uintptr
	addr_QGraphicsGridLayout_invalidate_default uintptr
	addr_QGraphicsGridLayout_setGeometry_default uintptr
	addr_QGraphicsGridLayout_sizeHint_default uintptr
	addr_QGraphicsGridLayout_getContentsMargins_default uintptr
	addr_QGraphicsGridLayout_updateGeometry_default uintptr
	addr_QGraphicsGridLayout_widgetEvent_default uintptr
	addr_QGraphicsGridLayout_isEmpty_default uintptr
	pg_QGraphicsGridLayout__set_override func(uintptr, int32)
	pg_QGraphicsGridLayout_delete func(uintptr)
)

func init() {
	symaddr(&addr_QGraphicsGridLayout_chain, "QGraphicsGridLayout_chain")
	symaddr(&addr_QGraphicsGridLayout_new, "QGraphicsGridLayout_new")
	symaddr(&addr_QGraphicsGridLayout_new2, "QGraphicsGridLayout_new2")
	symaddr(&addr_QGraphicsGridLayout_addItem, "QGraphicsGridLayout_addItem")
	symaddr(&addr_QGraphicsGridLayout_addItem2, "QGraphicsGridLayout_addItem2")
	reg(&pg_QGraphicsGridLayout_setHorizontalSpacing, "QGraphicsGridLayout_setHorizontalSpacing")
	reg(&pg_QGraphicsGridLayout_horizontalSpacing, "QGraphicsGridLayout_horizontalSpacing")
	reg(&pg_QGraphicsGridLayout_setVerticalSpacing, "QGraphicsGridLayout_setVerticalSpacing")
	reg(&pg_QGraphicsGridLayout_verticalSpacing, "QGraphicsGridLayout_verticalSpacing")
	reg(&pg_QGraphicsGridLayout_setSpacing, "QGraphicsGridLayout_setSpacing")
	reg(&pg_QGraphicsGridLayout_setRowSpacing, "QGraphicsGridLayout_setRowSpacing")
	reg(&pg_QGraphicsGridLayout_rowSpacing, "QGraphicsGridLayout_rowSpacing")
	reg(&pg_QGraphicsGridLayout_setColumnSpacing, "QGraphicsGridLayout_setColumnSpacing")
	reg(&pg_QGraphicsGridLayout_columnSpacing, "QGraphicsGridLayout_columnSpacing")
	symaddr(&addr_QGraphicsGridLayout_setRowStretchFactor, "QGraphicsGridLayout_setRowStretchFactor")
	symaddr(&addr_QGraphicsGridLayout_rowStretchFactor, "QGraphicsGridLayout_rowStretchFactor")
	symaddr(&addr_QGraphicsGridLayout_setColumnStretchFactor, "QGraphicsGridLayout_setColumnStretchFactor")
	symaddr(&addr_QGraphicsGridLayout_columnStretchFactor, "QGraphicsGridLayout_columnStretchFactor")
	reg(&pg_QGraphicsGridLayout_setRowMinimumHeight, "QGraphicsGridLayout_setRowMinimumHeight")
	reg(&pg_QGraphicsGridLayout_rowMinimumHeight, "QGraphicsGridLayout_rowMinimumHeight")
	reg(&pg_QGraphicsGridLayout_setRowPreferredHeight, "QGraphicsGridLayout_setRowPreferredHeight")
	reg(&pg_QGraphicsGridLayout_rowPreferredHeight, "QGraphicsGridLayout_rowPreferredHeight")
	reg(&pg_QGraphicsGridLayout_setRowMaximumHeight, "QGraphicsGridLayout_setRowMaximumHeight")
	reg(&pg_QGraphicsGridLayout_rowMaximumHeight, "QGraphicsGridLayout_rowMaximumHeight")
	reg(&pg_QGraphicsGridLayout_setRowFixedHeight, "QGraphicsGridLayout_setRowFixedHeight")
	reg(&pg_QGraphicsGridLayout_setColumnMinimumWidth, "QGraphicsGridLayout_setColumnMinimumWidth")
	reg(&pg_QGraphicsGridLayout_columnMinimumWidth, "QGraphicsGridLayout_columnMinimumWidth")
	reg(&pg_QGraphicsGridLayout_setColumnPreferredWidth, "QGraphicsGridLayout_setColumnPreferredWidth")
	reg(&pg_QGraphicsGridLayout_columnPreferredWidth, "QGraphicsGridLayout_columnPreferredWidth")
	reg(&pg_QGraphicsGridLayout_setColumnMaximumWidth, "QGraphicsGridLayout_setColumnMaximumWidth")
	reg(&pg_QGraphicsGridLayout_columnMaximumWidth, "QGraphicsGridLayout_columnMaximumWidth")
	reg(&pg_QGraphicsGridLayout_setColumnFixedWidth, "QGraphicsGridLayout_setColumnFixedWidth")
	symaddr(&addr_QGraphicsGridLayout_setRowAlignment, "QGraphicsGridLayout_setRowAlignment")
	symaddr(&addr_QGraphicsGridLayout_rowAlignment, "QGraphicsGridLayout_rowAlignment")
	symaddr(&addr_QGraphicsGridLayout_setColumnAlignment, "QGraphicsGridLayout_setColumnAlignment")
	symaddr(&addr_QGraphicsGridLayout_columnAlignment, "QGraphicsGridLayout_columnAlignment")
	symaddr(&addr_QGraphicsGridLayout_setAlignment, "QGraphicsGridLayout_setAlignment")
	symaddr(&addr_QGraphicsGridLayout_alignment, "QGraphicsGridLayout_alignment")
	symaddr(&addr_QGraphicsGridLayout_rowCount, "QGraphicsGridLayout_rowCount")
	symaddr(&addr_QGraphicsGridLayout_columnCount, "QGraphicsGridLayout_columnCount")
	symaddr(&addr_QGraphicsGridLayout_itemAt, "QGraphicsGridLayout_itemAt")
	symaddr(&addr_QGraphicsGridLayout_itemAt2, "QGraphicsGridLayout_itemAt2")
	symaddr(&addr_QGraphicsGridLayout_removeItem, "QGraphicsGridLayout_removeItem")
	symaddr(&addr_QGraphicsGridLayout_setInstantInvalidatePropagation, "QGraphicsGridLayout_setInstantInvalidatePropagation")
	symaddr(&addr_QGraphicsGridLayout_instantInvalidatePropagation, "QGraphicsGridLayout_instantInvalidatePropagation")
	symaddr(&addr_QGraphicsGridLayout_addChildLayoutItem, "QGraphicsGridLayout_addChildLayoutItem")
	symaddr(&addr_QGraphicsGridLayout_setGraphicsItem, "QGraphicsGridLayout_setGraphicsItem")
	symaddr(&addr_QGraphicsGridLayout_setOwnedByLayout, "QGraphicsGridLayout_setOwnedByLayout")
	symaddr(&addr_QGraphicsGridLayout_addItem3, "QGraphicsGridLayout_addItem3")
	symaddr(&addr_QGraphicsGridLayout_addItem4, "QGraphicsGridLayout_addItem4")
	symaddr(&addr_QGraphicsGridLayout_itemAt_default, "QGraphicsGridLayout_itemAt_default")
	symaddr(&addr_QGraphicsGridLayout_itemAt_default2, "QGraphicsGridLayout_itemAt_default2")
	symaddr(&addr_QGraphicsGridLayout_count_default, "QGraphicsGridLayout_count_default")
	symaddr(&addr_QGraphicsGridLayout_removeAt_default, "QGraphicsGridLayout_removeAt_default")
	symaddr(&addr_QGraphicsGridLayout_invalidate_default, "QGraphicsGridLayout_invalidate_default")
	symaddr(&addr_QGraphicsGridLayout_setGeometry_default, "QGraphicsGridLayout_setGeometry_default")
	symaddr(&addr_QGraphicsGridLayout_sizeHint_default, "QGraphicsGridLayout_sizeHint_default")
	symaddr(&addr_QGraphicsGridLayout_getContentsMargins_default, "QGraphicsGridLayout_getContentsMargins_default")
	symaddr(&addr_QGraphicsGridLayout_updateGeometry_default, "QGraphicsGridLayout_updateGeometry_default")
	symaddr(&addr_QGraphicsGridLayout_widgetEvent_default, "QGraphicsGridLayout_widgetEvent_default")
	symaddr(&addr_QGraphicsGridLayout_isEmpty_default, "QGraphicsGridLayout_isEmpty_default")
	reg(&pg_QGraphicsGridLayout__set_override, "QGraphicsGridLayout__set_override")
	reg(&pg_QGraphicsGridLayout_delete, "QGraphicsGridLayout_delete")
}

type QGraphicsGridLayout struct {
	h uintptr
	*QGraphicsLayout
}

func bindQGraphicsGridLayout(h uintptr) *QGraphicsGridLayout {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QGraphicsGridLayout_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QGraphicsGridLayout{h: h, QGraphicsLayout: &QGraphicsLayout{h: __c[0], QGraphicsLayoutItem: &QGraphicsLayoutItem{h: __c[1]}}}
}

func newownedQGraphicsGridLayout(h uintptr) *QGraphicsGridLayout {
	o := bindQGraphicsGridLayout(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QGraphicsGridLayout).Delete)
	}
	return o
}

func (this *QGraphicsGridLayout) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QGraphicsGridLayout) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQGraphicsGridLayout(h unsafe.Pointer) *QGraphicsGridLayout { return bindQGraphicsGridLayout(uintptr(h)) }
func UnsafeNewQGraphicsGridLayoutOwned(h unsafe.Pointer) *QGraphicsGridLayout { return newownedQGraphicsGridLayout(uintptr(h)) }

func (this *QGraphicsGridLayout) UnsafeRelease() {
	this.h = 0
	if this.QGraphicsLayout != nil {
		this.QGraphicsLayout.UnsafeRelease()
	}
}

func NewQGraphicsGridLayout() *QGraphicsGridLayout {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QGraphicsGridLayout_new, uintptr(__goid))
	__o := newownedQGraphicsGridLayout(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGraphicsGridLayout_virtIdx, QGraphicsGridLayout_enable)
	return __o
}

func NewQGraphicsGridLayout2(parent *QGraphicsLayoutItem) *QGraphicsGridLayout {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QGraphicsGridLayout_new2, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQGraphicsGridLayout(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGraphicsGridLayout_virtIdx, QGraphicsGridLayout_enable)
	return __o
}

func (this *QGraphicsGridLayout) AddItem(item *QGraphicsLayoutItem, row int, column int, rowSpan int, columnSpan int) {
	purego.SyscallN(addr_QGraphicsGridLayout_addItem, this.h, uintptr(item.UnsafePointer()), uintptr(row), uintptr(column), uintptr(rowSpan), uintptr(columnSpan))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QGraphicsGridLayout) AddItem2(item *QGraphicsLayoutItem, row int, column int) {
	purego.SyscallN(addr_QGraphicsGridLayout_addItem2, this.h, uintptr(item.UnsafePointer()), uintptr(row), uintptr(column))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QGraphicsGridLayout) SetHorizontalSpacing(spacing float64) {
	pg_QGraphicsGridLayout_setHorizontalSpacing(this.h, spacing)
	runtime.KeepAlive(this)
}

func (this *QGraphicsGridLayout) HorizontalSpacing() float64 {
	__r := pg_QGraphicsGridLayout_horizontalSpacing(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QGraphicsGridLayout) SetVerticalSpacing(spacing float64) {
	pg_QGraphicsGridLayout_setVerticalSpacing(this.h, spacing)
	runtime.KeepAlive(this)
}

func (this *QGraphicsGridLayout) VerticalSpacing() float64 {
	__r := pg_QGraphicsGridLayout_verticalSpacing(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QGraphicsGridLayout) SetSpacing(spacing float64) {
	pg_QGraphicsGridLayout_setSpacing(this.h, spacing)
	runtime.KeepAlive(this)
}

func (this *QGraphicsGridLayout) SetRowSpacing(row int, spacing float64) {
	pg_QGraphicsGridLayout_setRowSpacing(this.h, int32(row), spacing)
	runtime.KeepAlive(this)
}

func (this *QGraphicsGridLayout) RowSpacing(row int) float64 {
	__r := pg_QGraphicsGridLayout_rowSpacing(this.h, int32(row))
	runtime.KeepAlive(this)
	return __r
}

func (this *QGraphicsGridLayout) SetColumnSpacing(column int, spacing float64) {
	pg_QGraphicsGridLayout_setColumnSpacing(this.h, int32(column), spacing)
	runtime.KeepAlive(this)
}

func (this *QGraphicsGridLayout) ColumnSpacing(column int) float64 {
	__r := pg_QGraphicsGridLayout_columnSpacing(this.h, int32(column))
	runtime.KeepAlive(this)
	return __r
}

func (this *QGraphicsGridLayout) SetRowStretchFactor(row int, stretch int) {
	purego.SyscallN(addr_QGraphicsGridLayout_setRowStretchFactor, this.h, uintptr(row), uintptr(stretch))
	runtime.KeepAlive(this)
}

func (this *QGraphicsGridLayout) RowStretchFactor(row int) int {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsGridLayout_rowStretchFactor, this.h, uintptr(row))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QGraphicsGridLayout) SetColumnStretchFactor(column int, stretch int) {
	purego.SyscallN(addr_QGraphicsGridLayout_setColumnStretchFactor, this.h, uintptr(column), uintptr(stretch))
	runtime.KeepAlive(this)
}

func (this *QGraphicsGridLayout) ColumnStretchFactor(column int) int {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsGridLayout_columnStretchFactor, this.h, uintptr(column))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QGraphicsGridLayout) SetRowMinimumHeight(row int, height float64) {
	pg_QGraphicsGridLayout_setRowMinimumHeight(this.h, int32(row), height)
	runtime.KeepAlive(this)
}

func (this *QGraphicsGridLayout) RowMinimumHeight(row int) float64 {
	__r := pg_QGraphicsGridLayout_rowMinimumHeight(this.h, int32(row))
	runtime.KeepAlive(this)
	return __r
}

func (this *QGraphicsGridLayout) SetRowPreferredHeight(row int, height float64) {
	pg_QGraphicsGridLayout_setRowPreferredHeight(this.h, int32(row), height)
	runtime.KeepAlive(this)
}

func (this *QGraphicsGridLayout) RowPreferredHeight(row int) float64 {
	__r := pg_QGraphicsGridLayout_rowPreferredHeight(this.h, int32(row))
	runtime.KeepAlive(this)
	return __r
}

func (this *QGraphicsGridLayout) SetRowMaximumHeight(row int, height float64) {
	pg_QGraphicsGridLayout_setRowMaximumHeight(this.h, int32(row), height)
	runtime.KeepAlive(this)
}

func (this *QGraphicsGridLayout) RowMaximumHeight(row int) float64 {
	__r := pg_QGraphicsGridLayout_rowMaximumHeight(this.h, int32(row))
	runtime.KeepAlive(this)
	return __r
}

func (this *QGraphicsGridLayout) SetRowFixedHeight(row int, height float64) {
	pg_QGraphicsGridLayout_setRowFixedHeight(this.h, int32(row), height)
	runtime.KeepAlive(this)
}

func (this *QGraphicsGridLayout) SetColumnMinimumWidth(column int, width float64) {
	pg_QGraphicsGridLayout_setColumnMinimumWidth(this.h, int32(column), width)
	runtime.KeepAlive(this)
}

func (this *QGraphicsGridLayout) ColumnMinimumWidth(column int) float64 {
	__r := pg_QGraphicsGridLayout_columnMinimumWidth(this.h, int32(column))
	runtime.KeepAlive(this)
	return __r
}

func (this *QGraphicsGridLayout) SetColumnPreferredWidth(column int, width float64) {
	pg_QGraphicsGridLayout_setColumnPreferredWidth(this.h, int32(column), width)
	runtime.KeepAlive(this)
}

func (this *QGraphicsGridLayout) ColumnPreferredWidth(column int) float64 {
	__r := pg_QGraphicsGridLayout_columnPreferredWidth(this.h, int32(column))
	runtime.KeepAlive(this)
	return __r
}

func (this *QGraphicsGridLayout) SetColumnMaximumWidth(column int, width float64) {
	pg_QGraphicsGridLayout_setColumnMaximumWidth(this.h, int32(column), width)
	runtime.KeepAlive(this)
}

func (this *QGraphicsGridLayout) ColumnMaximumWidth(column int) float64 {
	__r := pg_QGraphicsGridLayout_columnMaximumWidth(this.h, int32(column))
	runtime.KeepAlive(this)
	return __r
}

func (this *QGraphicsGridLayout) SetColumnFixedWidth(column int, width float64) {
	pg_QGraphicsGridLayout_setColumnFixedWidth(this.h, int32(column), width)
	runtime.KeepAlive(this)
}

func (this *QGraphicsGridLayout) SetRowAlignment(row int, alignment AlignmentFlag) {
	purego.SyscallN(addr_QGraphicsGridLayout_setRowAlignment, this.h, uintptr(row), uintptr(alignment))
	runtime.KeepAlive(this)
}

func (this *QGraphicsGridLayout) RowAlignment(row int) AlignmentFlag {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsGridLayout_rowAlignment, this.h, uintptr(row))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return AlignmentFlag(__r)
}

func (this *QGraphicsGridLayout) SetColumnAlignment(column int, alignment AlignmentFlag) {
	purego.SyscallN(addr_QGraphicsGridLayout_setColumnAlignment, this.h, uintptr(column), uintptr(alignment))
	runtime.KeepAlive(this)
}

func (this *QGraphicsGridLayout) ColumnAlignment(column int) AlignmentFlag {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsGridLayout_columnAlignment, this.h, uintptr(column))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return AlignmentFlag(__r)
}

func (this *QGraphicsGridLayout) SetAlignment(item *QGraphicsLayoutItem, alignment AlignmentFlag) {
	purego.SyscallN(addr_QGraphicsGridLayout_setAlignment, this.h, uintptr(item.UnsafePointer()), uintptr(alignment))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QGraphicsGridLayout) Alignment(item *QGraphicsLayoutItem) AlignmentFlag {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsGridLayout_alignment, this.h, uintptr(item.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
	return AlignmentFlag(__r)
}

func (this *QGraphicsGridLayout) RowCount() int {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsGridLayout_rowCount, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QGraphicsGridLayout) ColumnCount() int {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsGridLayout_columnCount, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QGraphicsGridLayout) ItemAt(row int, column int) *QGraphicsLayoutItem {
	__r, _, _ := purego.SyscallN(addr_QGraphicsGridLayout_itemAt, this.h, uintptr(row), uintptr(column))
	runtime.KeepAlive(this)
	return bindQGraphicsLayoutItem(__r)
}

func (this *QGraphicsGridLayout) ItemAtWithIndex(index int) *QGraphicsLayoutItem {
	__r, _, _ := purego.SyscallN(addr_QGraphicsGridLayout_itemAt2, this.h, uintptr(index))
	runtime.KeepAlive(this)
	return bindQGraphicsLayoutItem(__r)
}

func (this *QGraphicsGridLayout) RemoveItem(item *QGraphicsLayoutItem) {
	purego.SyscallN(addr_QGraphicsGridLayout_removeItem, this.h, uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func QGraphicsGridLayout_SetInstantInvalidatePropagation(enable bool) {
	purego.SyscallN(addr_QGraphicsGridLayout_setInstantInvalidatePropagation, uintptr(b2i(enable)))
}

func QGraphicsGridLayout_InstantInvalidatePropagation() bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsGridLayout_instantInvalidatePropagation)
	__r := int32(uint32(__r0))
	return __r != 0
}

func (this *QGraphicsGridLayout) AddChildLayoutItem(layoutItem *QGraphicsLayoutItem) {
	purego.SyscallN(addr_QGraphicsGridLayout_addChildLayoutItem, this.h, uintptr(layoutItem.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(layoutItem)
}

func (this *QGraphicsGridLayout) SetGraphicsItem(item *QGraphicsItem) {
	purego.SyscallN(addr_QGraphicsGridLayout_setGraphicsItem, this.h, uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QGraphicsGridLayout) SetOwnedByLayout(ownedByLayout bool) {
	purego.SyscallN(addr_QGraphicsGridLayout_setOwnedByLayout, this.h, uintptr(b2i(ownedByLayout)))
	runtime.KeepAlive(this)
}

func (this *QGraphicsGridLayout) AddItem3(item *QGraphicsLayoutItem, row int, column int, rowSpan int, columnSpan int, alignment AlignmentFlag) {
	purego.SyscallN(addr_QGraphicsGridLayout_addItem3, this.h, uintptr(item.UnsafePointer()), uintptr(row), uintptr(column), uintptr(rowSpan), uintptr(columnSpan), uintptr(alignment))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QGraphicsGridLayout) AddItem4(item *QGraphicsLayoutItem, row int, column int, alignment AlignmentFlag) {
	purego.SyscallN(addr_QGraphicsGridLayout_addItem4, this.h, uintptr(item.UnsafePointer()), uintptr(row), uintptr(column), uintptr(alignment))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QGraphicsGridLayout) ItemAtDefault(row int, column int) *QGraphicsLayoutItem {
	__r, _, _ := purego.SyscallN(addr_QGraphicsGridLayout_itemAt_default, this.h, uintptr(row), uintptr(column))
	runtime.KeepAlive(this)
	return bindQGraphicsLayoutItem(__r)
}

func (this *QGraphicsGridLayout) ItemAtWithIndexDefault(index int) *QGraphicsLayoutItem {
	__r, _, _ := purego.SyscallN(addr_QGraphicsGridLayout_itemAt_default2, this.h, uintptr(index))
	runtime.KeepAlive(this)
	return bindQGraphicsLayoutItem(__r)
}

func (this *QGraphicsGridLayout) CountDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsGridLayout_count_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QGraphicsGridLayout) RemoveAtDefault(index int) {
	purego.SyscallN(addr_QGraphicsGridLayout_removeAt_default, this.h, uintptr(index))
	runtime.KeepAlive(this)
}

func (this *QGraphicsGridLayout) InvalidateDefault() {
	purego.SyscallN(addr_QGraphicsGridLayout_invalidate_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsGridLayout) SetGeometryDefault(rect *QRectF) {
	purego.SyscallN(addr_QGraphicsGridLayout_setGeometry_default, this.h, uintptr(rect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
}

func (this *QGraphicsGridLayout) SizeHintDefault(which SizeHint, constraint *QSizeF) *QSizeF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsGridLayout_sizeHint_default, this.h, uintptr(which), uintptr(constraint.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(constraint)
	return newownedQSizeF(__r)
}

func (this *QGraphicsGridLayout) GetContentsMarginsDefault() (float64, float64, float64, float64) {
	var left_out float64
	var top_out float64
	var right_out float64
	var bottom_out float64
	purego.SyscallN(addr_QGraphicsGridLayout_getContentsMargins_default, this.h, uintptr(unsafe.Pointer(&left_out)), uintptr(unsafe.Pointer(&top_out)), uintptr(unsafe.Pointer(&right_out)), uintptr(unsafe.Pointer(&bottom_out)))
	runtime.KeepAlive(this)
	return left_out, top_out, right_out, bottom_out
}

func (this *QGraphicsGridLayout) UpdateGeometryDefault() {
	purego.SyscallN(addr_QGraphicsGridLayout_updateGeometry_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsGridLayout) WidgetEventDefault(e *QEvent) {
	purego.SyscallN(addr_QGraphicsGridLayout_widgetEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QGraphicsGridLayout) IsEmptyDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsGridLayout_isEmpty_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QGraphicsGridLayout_virtIdx(slot int64) int32 {
	switch slot {
	case -5627843058667511180:
		return 0
	case 7508512214013391107:
		return 1
	case 1230346564330554752:
		return 2
	case 6991341317990632368:
		return 3
	case 5684820848726969227:
		return 4
	case 729690261457684101:
		return 5
	case 164760286202646972:
		return 6
	case 8082455443689927668:
		return 7
	case 8535746785413991679:
		return 8
	case -626822430251098498:
		return 9
	}
	return -1
}

func QGraphicsGridLayout_enable(h uintptr, idx int32) { pg_QGraphicsGridLayout__set_override(h, idx) }

func (this *QGraphicsGridLayout) OnCount(fn func() int) {
	registerVirt(this.UnsafeGoID(), -5627843058667511180, func() int32 { return int32(fn()) }) // slot "count"
}

func (this *QGraphicsGridLayout) OnItemAt(fn func(int) *QGraphicsLayoutItem) {
	registerVirt(this.UnsafeGoID(), 7508512214013391107, func(index int32) uintptr { return uintptr(fn(int(index)).UnsafePointer()) }) // slot "itemAt"
}

func (this *QGraphicsGridLayout) OnRemoveAt(fn func(int)) {
	registerVirt(this.UnsafeGoID(), 1230346564330554752, func(index int32) { fn(int(index)) }) // slot "removeAt"
}

func (this *QGraphicsGridLayout) OnInvalidate(fn func()) {
	registerVirt(this.UnsafeGoID(), 6991341317990632368, func() { fn() }) // slot "invalidate"
}

func (this *QGraphicsGridLayout) OnSetGeometry(fn func(*QRectF)) {
	registerVirt(this.UnsafeGoID(), 5684820848726969227, func(rect uintptr) { fn(UnsafeNewQRectF(unsafe.Pointer(rect))) }) // slot "setGeometry"
}

func (this *QGraphicsGridLayout) OnSizeHint(fn func(SizeHint, *QSizeF) *QSizeF) {
	registerVirt(this.UnsafeGoID(), 729690261457684101, func(which int32, constraint uintptr) uintptr { return uintptr(fn(SizeHint(which), UnsafeNewQSizeF(unsafe.Pointer(constraint))).UnsafePointer()) }) // slot "sizeHint"
}

func (this *QGraphicsGridLayout) OnGetContentsMargins(fn func(*float64, *float64, *float64, *float64)) {
	registerVirt(this.UnsafeGoID(), 164760286202646972, func(left uintptr, top uintptr, right uintptr, bottom uintptr) { fn((*float64)(unsafe.Pointer(left)), (*float64)(unsafe.Pointer(top)), (*float64)(unsafe.Pointer(right)), (*float64)(unsafe.Pointer(bottom))) }) // slot "getContentsMargins"
}

func (this *QGraphicsGridLayout) OnUpdateGeometry(fn func()) {
	registerVirt(this.UnsafeGoID(), 8082455443689927668, func() { fn() }) // slot "updateGeometry"
}

func (this *QGraphicsGridLayout) OnWidgetEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 8535746785413991679, func(e uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(e))) }) // slot "widgetEvent"
}

func (this *QGraphicsGridLayout) OnIsEmpty(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -626822430251098498, func() int32 { return b2i(fn()) }) // slot "isEmpty"
}

func (this *QGraphicsGridLayout) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QGraphicsGridLayout_delete(this.h)
	}
	this.UnsafeRelease()
}
