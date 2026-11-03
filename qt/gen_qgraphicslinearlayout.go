package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QGraphicsLinearLayout_chain uintptr
	addr_QGraphicsLinearLayout_new uintptr
	addr_QGraphicsLinearLayout_new2 uintptr
	addr_QGraphicsLinearLayout_new3 uintptr
	addr_QGraphicsLinearLayout_new4 uintptr
	addr_QGraphicsLinearLayout_setOrientation uintptr
	addr_QGraphicsLinearLayout_orientation uintptr
	addr_QGraphicsLinearLayout_addItem uintptr
	addr_QGraphicsLinearLayout_addStretch uintptr
	addr_QGraphicsLinearLayout_insertItem uintptr
	addr_QGraphicsLinearLayout_insertStretch uintptr
	addr_QGraphicsLinearLayout_removeItem uintptr
	pg_QGraphicsLinearLayout_setSpacing func(uintptr, float64)
	pg_QGraphicsLinearLayout_spacing func(uintptr) float64
	pg_QGraphicsLinearLayout_setItemSpacing func(uintptr, int32, float64)
	pg_QGraphicsLinearLayout_itemSpacing func(uintptr, int32) float64
	addr_QGraphicsLinearLayout_setStretchFactor uintptr
	addr_QGraphicsLinearLayout_stretchFactor uintptr
	addr_QGraphicsLinearLayout_setAlignment uintptr
	addr_QGraphicsLinearLayout_alignment uintptr
	addr_QGraphicsLinearLayout_dump uintptr
	addr_QGraphicsLinearLayout_setInstantInvalidatePropagation uintptr
	addr_QGraphicsLinearLayout_instantInvalidatePropagation uintptr
	addr_QGraphicsLinearLayout_addChildLayoutItem uintptr
	addr_QGraphicsLinearLayout_setGraphicsItem uintptr
	addr_QGraphicsLinearLayout_setOwnedByLayout uintptr
	addr_QGraphicsLinearLayout_addStretch2 uintptr
	addr_QGraphicsLinearLayout_insertStretch2 uintptr
	addr_QGraphicsLinearLayout_dump2 uintptr
	addr_QGraphicsLinearLayout_removeAt_default uintptr
	addr_QGraphicsLinearLayout_setGeometry_default uintptr
	addr_QGraphicsLinearLayout_count_default uintptr
	addr_QGraphicsLinearLayout_itemAt_default uintptr
	addr_QGraphicsLinearLayout_invalidate_default uintptr
	addr_QGraphicsLinearLayout_sizeHint_default uintptr
	addr_QGraphicsLinearLayout_getContentsMargins_default uintptr
	addr_QGraphicsLinearLayout_updateGeometry_default uintptr
	addr_QGraphicsLinearLayout_widgetEvent_default uintptr
	addr_QGraphicsLinearLayout_isEmpty_default uintptr
	pg_QGraphicsLinearLayout__set_override func(uintptr, int32)
	pg_QGraphicsLinearLayout_delete func(uintptr)
)

func init() {
	symaddr(&addr_QGraphicsLinearLayout_chain, "QGraphicsLinearLayout_chain")
	symaddr(&addr_QGraphicsLinearLayout_new, "QGraphicsLinearLayout_new")
	symaddr(&addr_QGraphicsLinearLayout_new2, "QGraphicsLinearLayout_new2")
	symaddr(&addr_QGraphicsLinearLayout_new3, "QGraphicsLinearLayout_new3")
	symaddr(&addr_QGraphicsLinearLayout_new4, "QGraphicsLinearLayout_new4")
	symaddr(&addr_QGraphicsLinearLayout_setOrientation, "QGraphicsLinearLayout_setOrientation")
	symaddr(&addr_QGraphicsLinearLayout_orientation, "QGraphicsLinearLayout_orientation")
	symaddr(&addr_QGraphicsLinearLayout_addItem, "QGraphicsLinearLayout_addItem")
	symaddr(&addr_QGraphicsLinearLayout_addStretch, "QGraphicsLinearLayout_addStretch")
	symaddr(&addr_QGraphicsLinearLayout_insertItem, "QGraphicsLinearLayout_insertItem")
	symaddr(&addr_QGraphicsLinearLayout_insertStretch, "QGraphicsLinearLayout_insertStretch")
	symaddr(&addr_QGraphicsLinearLayout_removeItem, "QGraphicsLinearLayout_removeItem")
	reg(&pg_QGraphicsLinearLayout_setSpacing, "QGraphicsLinearLayout_setSpacing")
	reg(&pg_QGraphicsLinearLayout_spacing, "QGraphicsLinearLayout_spacing")
	reg(&pg_QGraphicsLinearLayout_setItemSpacing, "QGraphicsLinearLayout_setItemSpacing")
	reg(&pg_QGraphicsLinearLayout_itemSpacing, "QGraphicsLinearLayout_itemSpacing")
	symaddr(&addr_QGraphicsLinearLayout_setStretchFactor, "QGraphicsLinearLayout_setStretchFactor")
	symaddr(&addr_QGraphicsLinearLayout_stretchFactor, "QGraphicsLinearLayout_stretchFactor")
	symaddr(&addr_QGraphicsLinearLayout_setAlignment, "QGraphicsLinearLayout_setAlignment")
	symaddr(&addr_QGraphicsLinearLayout_alignment, "QGraphicsLinearLayout_alignment")
	symaddr(&addr_QGraphicsLinearLayout_dump, "QGraphicsLinearLayout_dump")
	symaddr(&addr_QGraphicsLinearLayout_setInstantInvalidatePropagation, "QGraphicsLinearLayout_setInstantInvalidatePropagation")
	symaddr(&addr_QGraphicsLinearLayout_instantInvalidatePropagation, "QGraphicsLinearLayout_instantInvalidatePropagation")
	symaddr(&addr_QGraphicsLinearLayout_addChildLayoutItem, "QGraphicsLinearLayout_addChildLayoutItem")
	symaddr(&addr_QGraphicsLinearLayout_setGraphicsItem, "QGraphicsLinearLayout_setGraphicsItem")
	symaddr(&addr_QGraphicsLinearLayout_setOwnedByLayout, "QGraphicsLinearLayout_setOwnedByLayout")
	symaddr(&addr_QGraphicsLinearLayout_addStretch2, "QGraphicsLinearLayout_addStretch2")
	symaddr(&addr_QGraphicsLinearLayout_insertStretch2, "QGraphicsLinearLayout_insertStretch2")
	symaddr(&addr_QGraphicsLinearLayout_dump2, "QGraphicsLinearLayout_dump2")
	symaddr(&addr_QGraphicsLinearLayout_removeAt_default, "QGraphicsLinearLayout_removeAt_default")
	symaddr(&addr_QGraphicsLinearLayout_setGeometry_default, "QGraphicsLinearLayout_setGeometry_default")
	symaddr(&addr_QGraphicsLinearLayout_count_default, "QGraphicsLinearLayout_count_default")
	symaddr(&addr_QGraphicsLinearLayout_itemAt_default, "QGraphicsLinearLayout_itemAt_default")
	symaddr(&addr_QGraphicsLinearLayout_invalidate_default, "QGraphicsLinearLayout_invalidate_default")
	symaddr(&addr_QGraphicsLinearLayout_sizeHint_default, "QGraphicsLinearLayout_sizeHint_default")
	symaddr(&addr_QGraphicsLinearLayout_getContentsMargins_default, "QGraphicsLinearLayout_getContentsMargins_default")
	symaddr(&addr_QGraphicsLinearLayout_updateGeometry_default, "QGraphicsLinearLayout_updateGeometry_default")
	symaddr(&addr_QGraphicsLinearLayout_widgetEvent_default, "QGraphicsLinearLayout_widgetEvent_default")
	symaddr(&addr_QGraphicsLinearLayout_isEmpty_default, "QGraphicsLinearLayout_isEmpty_default")
	reg(&pg_QGraphicsLinearLayout__set_override, "QGraphicsLinearLayout__set_override")
	reg(&pg_QGraphicsLinearLayout_delete, "QGraphicsLinearLayout_delete")
}

type QGraphicsLinearLayout struct {
	h uintptr
	*QGraphicsLayout
}

func bindQGraphicsLinearLayout(h uintptr) *QGraphicsLinearLayout {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QGraphicsLinearLayout_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QGraphicsLinearLayout{h: h, QGraphicsLayout: &QGraphicsLayout{h: __c[0], QGraphicsLayoutItem: &QGraphicsLayoutItem{h: __c[1]}}}
}

func newownedQGraphicsLinearLayout(h uintptr) *QGraphicsLinearLayout {
	o := bindQGraphicsLinearLayout(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QGraphicsLinearLayout).Delete)
	}
	return o
}

func (this *QGraphicsLinearLayout) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QGraphicsLinearLayout) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQGraphicsLinearLayout(h unsafe.Pointer) *QGraphicsLinearLayout { return bindQGraphicsLinearLayout(uintptr(h)) }
func UnsafeNewQGraphicsLinearLayoutOwned(h unsafe.Pointer) *QGraphicsLinearLayout { return newownedQGraphicsLinearLayout(uintptr(h)) }

func (this *QGraphicsLinearLayout) UnsafeRelease() {
	this.h = 0
	if this.QGraphicsLayout != nil {
		this.QGraphicsLayout.UnsafeRelease()
	}
}

func NewQGraphicsLinearLayout() *QGraphicsLinearLayout {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QGraphicsLinearLayout_new, uintptr(__goid))
	__o := newownedQGraphicsLinearLayout(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGraphicsLinearLayout_virtIdx, QGraphicsLinearLayout_enable)
	return __o
}

func NewQGraphicsLinearLayout2(orientation Orientation) *QGraphicsLinearLayout {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QGraphicsLinearLayout_new2, uintptr(__goid), uintptr(orientation))
	__o := newownedQGraphicsLinearLayout(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGraphicsLinearLayout_virtIdx, QGraphicsLinearLayout_enable)
	return __o
}

func NewQGraphicsLinearLayout3(parent *QGraphicsLayoutItem) *QGraphicsLinearLayout {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QGraphicsLinearLayout_new3, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQGraphicsLinearLayout(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGraphicsLinearLayout_virtIdx, QGraphicsLinearLayout_enable)
	return __o
}

func NewQGraphicsLinearLayout4(orientation Orientation, parent *QGraphicsLayoutItem) *QGraphicsLinearLayout {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QGraphicsLinearLayout_new4, uintptr(__goid), uintptr(orientation), uintptr(parent.UnsafePointer()))
	__o := newownedQGraphicsLinearLayout(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGraphicsLinearLayout_virtIdx, QGraphicsLinearLayout_enable)
	return __o
}

func (this *QGraphicsLinearLayout) SetOrientation(orientation Orientation) {
	purego.SyscallN(addr_QGraphicsLinearLayout_setOrientation, this.h, uintptr(orientation))
	runtime.KeepAlive(this)
}

func (this *QGraphicsLinearLayout) Orientation() Orientation {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsLinearLayout_orientation, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return Orientation(__r)
}

func (this *QGraphicsLinearLayout) AddItem(item *QGraphicsLayoutItem) {
	purego.SyscallN(addr_QGraphicsLinearLayout_addItem, this.h, uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QGraphicsLinearLayout) AddStretch() {
	purego.SyscallN(addr_QGraphicsLinearLayout_addStretch, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsLinearLayout) InsertItem(index int, item *QGraphicsLayoutItem) {
	purego.SyscallN(addr_QGraphicsLinearLayout_insertItem, this.h, uintptr(index), uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QGraphicsLinearLayout) InsertStretch(index int) {
	purego.SyscallN(addr_QGraphicsLinearLayout_insertStretch, this.h, uintptr(index))
	runtime.KeepAlive(this)
}

func (this *QGraphicsLinearLayout) RemoveItem(item *QGraphicsLayoutItem) {
	purego.SyscallN(addr_QGraphicsLinearLayout_removeItem, this.h, uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QGraphicsLinearLayout) SetSpacing(spacing float64) {
	pg_QGraphicsLinearLayout_setSpacing(this.h, spacing)
	runtime.KeepAlive(this)
}

func (this *QGraphicsLinearLayout) Spacing() float64 {
	__r := pg_QGraphicsLinearLayout_spacing(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QGraphicsLinearLayout) SetItemSpacing(index int, spacing float64) {
	pg_QGraphicsLinearLayout_setItemSpacing(this.h, int32(index), spacing)
	runtime.KeepAlive(this)
}

func (this *QGraphicsLinearLayout) ItemSpacing(index int) float64 {
	__r := pg_QGraphicsLinearLayout_itemSpacing(this.h, int32(index))
	runtime.KeepAlive(this)
	return __r
}

func (this *QGraphicsLinearLayout) SetStretchFactor(item *QGraphicsLayoutItem, stretch int) {
	purego.SyscallN(addr_QGraphicsLinearLayout_setStretchFactor, this.h, uintptr(item.UnsafePointer()), uintptr(stretch))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QGraphicsLinearLayout) StretchFactor(item *QGraphicsLayoutItem) int {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsLinearLayout_stretchFactor, this.h, uintptr(item.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
	return int(__r)
}

func (this *QGraphicsLinearLayout) SetAlignment(item *QGraphicsLayoutItem, alignment AlignmentFlag) {
	purego.SyscallN(addr_QGraphicsLinearLayout_setAlignment, this.h, uintptr(item.UnsafePointer()), uintptr(alignment))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QGraphicsLinearLayout) Alignment(item *QGraphicsLayoutItem) AlignmentFlag {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsLinearLayout_alignment, this.h, uintptr(item.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
	return AlignmentFlag(__r)
}

func (this *QGraphicsLinearLayout) Dump() {
	purego.SyscallN(addr_QGraphicsLinearLayout_dump, this.h)
	runtime.KeepAlive(this)
}

func QGraphicsLinearLayout_SetInstantInvalidatePropagation(enable bool) {
	purego.SyscallN(addr_QGraphicsLinearLayout_setInstantInvalidatePropagation, uintptr(b2i(enable)))
}

func QGraphicsLinearLayout_InstantInvalidatePropagation() bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsLinearLayout_instantInvalidatePropagation)
	__r := int32(uint32(__r0))
	return __r != 0
}

func (this *QGraphicsLinearLayout) AddChildLayoutItem(layoutItem *QGraphicsLayoutItem) {
	purego.SyscallN(addr_QGraphicsLinearLayout_addChildLayoutItem, this.h, uintptr(layoutItem.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(layoutItem)
}

func (this *QGraphicsLinearLayout) SetGraphicsItem(item *QGraphicsItem) {
	purego.SyscallN(addr_QGraphicsLinearLayout_setGraphicsItem, this.h, uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QGraphicsLinearLayout) SetOwnedByLayout(ownedByLayout bool) {
	purego.SyscallN(addr_QGraphicsLinearLayout_setOwnedByLayout, this.h, uintptr(b2i(ownedByLayout)))
	runtime.KeepAlive(this)
}

func (this *QGraphicsLinearLayout) AddStretchWithStretch(stretch int) {
	purego.SyscallN(addr_QGraphicsLinearLayout_addStretch2, this.h, uintptr(stretch))
	runtime.KeepAlive(this)
}

func (this *QGraphicsLinearLayout) InsertStretch2(index int, stretch int) {
	purego.SyscallN(addr_QGraphicsLinearLayout_insertStretch2, this.h, uintptr(index), uintptr(stretch))
	runtime.KeepAlive(this)
}

func (this *QGraphicsLinearLayout) DumpWithIndent(indent int) {
	purego.SyscallN(addr_QGraphicsLinearLayout_dump2, this.h, uintptr(indent))
	runtime.KeepAlive(this)
}

func (this *QGraphicsLinearLayout) RemoveAtDefault(index int) {
	purego.SyscallN(addr_QGraphicsLinearLayout_removeAt_default, this.h, uintptr(index))
	runtime.KeepAlive(this)
}

func (this *QGraphicsLinearLayout) SetGeometryDefault(rect *QRectF) {
	purego.SyscallN(addr_QGraphicsLinearLayout_setGeometry_default, this.h, uintptr(rect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
}

func (this *QGraphicsLinearLayout) CountDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsLinearLayout_count_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QGraphicsLinearLayout) ItemAtDefault(index int) *QGraphicsLayoutItem {
	__r, _, _ := purego.SyscallN(addr_QGraphicsLinearLayout_itemAt_default, this.h, uintptr(index))
	runtime.KeepAlive(this)
	return bindQGraphicsLayoutItem(__r)
}

func (this *QGraphicsLinearLayout) InvalidateDefault() {
	purego.SyscallN(addr_QGraphicsLinearLayout_invalidate_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsLinearLayout) SizeHintDefault(which SizeHint, constraint *QSizeF) *QSizeF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsLinearLayout_sizeHint_default, this.h, uintptr(which), uintptr(constraint.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(constraint)
	return newownedQSizeF(__r)
}

func (this *QGraphicsLinearLayout) GetContentsMarginsDefault() (float64, float64, float64, float64) {
	var left_out float64
	var top_out float64
	var right_out float64
	var bottom_out float64
	purego.SyscallN(addr_QGraphicsLinearLayout_getContentsMargins_default, this.h, uintptr(unsafe.Pointer(&left_out)), uintptr(unsafe.Pointer(&top_out)), uintptr(unsafe.Pointer(&right_out)), uintptr(unsafe.Pointer(&bottom_out)))
	runtime.KeepAlive(this)
	return left_out, top_out, right_out, bottom_out
}

func (this *QGraphicsLinearLayout) UpdateGeometryDefault() {
	purego.SyscallN(addr_QGraphicsLinearLayout_updateGeometry_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsLinearLayout) WidgetEventDefault(e *QEvent) {
	purego.SyscallN(addr_QGraphicsLinearLayout_widgetEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QGraphicsLinearLayout) IsEmptyDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsLinearLayout_isEmpty_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QGraphicsLinearLayout_virtIdx(slot int64) int32 {
	switch slot {
	case 1230346564330554752:
		return 0
	case 5684820848726969227:
		return 1
	case -5627843058667511180:
		return 2
	case 7508512214013391107:
		return 3
	case 6991341317990632368:
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

func QGraphicsLinearLayout_enable(h uintptr, idx int32) { pg_QGraphicsLinearLayout__set_override(h, idx) }

func (this *QGraphicsLinearLayout) OnRemoveAt(fn func(int)) {
	registerVirt(this.UnsafeGoID(), 1230346564330554752, func(index int32) { fn(int(index)) }) // slot "removeAt"
}

func (this *QGraphicsLinearLayout) OnSetGeometry(fn func(*QRectF)) {
	registerVirt(this.UnsafeGoID(), 5684820848726969227, func(rect uintptr) { fn(UnsafeNewQRectF(unsafe.Pointer(rect))) }) // slot "setGeometry"
}

func (this *QGraphicsLinearLayout) OnCount(fn func() int) {
	registerVirt(this.UnsafeGoID(), -5627843058667511180, func() int32 { return int32(fn()) }) // slot "count"
}

func (this *QGraphicsLinearLayout) OnItemAt(fn func(int) *QGraphicsLayoutItem) {
	registerVirt(this.UnsafeGoID(), 7508512214013391107, func(index int32) uintptr { return uintptr(fn(int(index)).UnsafePointer()) }) // slot "itemAt"
}

func (this *QGraphicsLinearLayout) OnInvalidate(fn func()) {
	registerVirt(this.UnsafeGoID(), 6991341317990632368, func() { fn() }) // slot "invalidate"
}

func (this *QGraphicsLinearLayout) OnSizeHint(fn func(SizeHint, *QSizeF) *QSizeF) {
	registerVirt(this.UnsafeGoID(), 729690261457684101, func(which int32, constraint uintptr) uintptr { return uintptr(fn(SizeHint(which), UnsafeNewQSizeF(unsafe.Pointer(constraint))).UnsafePointer()) }) // slot "sizeHint"
}

func (this *QGraphicsLinearLayout) OnGetContentsMargins(fn func(*float64, *float64, *float64, *float64)) {
	registerVirt(this.UnsafeGoID(), 164760286202646972, func(left uintptr, top uintptr, right uintptr, bottom uintptr) { fn((*float64)(unsafe.Pointer(left)), (*float64)(unsafe.Pointer(top)), (*float64)(unsafe.Pointer(right)), (*float64)(unsafe.Pointer(bottom))) }) // slot "getContentsMargins"
}

func (this *QGraphicsLinearLayout) OnUpdateGeometry(fn func()) {
	registerVirt(this.UnsafeGoID(), 8082455443689927668, func() { fn() }) // slot "updateGeometry"
}

func (this *QGraphicsLinearLayout) OnWidgetEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 8535746785413991679, func(e uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(e))) }) // slot "widgetEvent"
}

func (this *QGraphicsLinearLayout) OnIsEmpty(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -626822430251098498, func() int32 { return b2i(fn()) }) // slot "isEmpty"
}

func (this *QGraphicsLinearLayout) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QGraphicsLinearLayout_delete(this.h)
	}
	this.UnsafeRelease()
}
