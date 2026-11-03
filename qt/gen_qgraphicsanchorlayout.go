package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QGraphicsAnchorLayout_chain uintptr
	addr_QGraphicsAnchorLayout_new uintptr
	addr_QGraphicsAnchorLayout_new2 uintptr
	addr_QGraphicsAnchorLayout_addAnchor uintptr
	addr_QGraphicsAnchorLayout_anchor uintptr
	addr_QGraphicsAnchorLayout_addCornerAnchors uintptr
	addr_QGraphicsAnchorLayout_addAnchors uintptr
	pg_QGraphicsAnchorLayout_setHorizontalSpacing func(uintptr, float64)
	pg_QGraphicsAnchorLayout_setVerticalSpacing func(uintptr, float64)
	pg_QGraphicsAnchorLayout_setSpacing func(uintptr, float64)
	pg_QGraphicsAnchorLayout_horizontalSpacing func(uintptr) float64
	pg_QGraphicsAnchorLayout_verticalSpacing func(uintptr) float64
	addr_QGraphicsAnchorLayout_setInstantInvalidatePropagation uintptr
	addr_QGraphicsAnchorLayout_instantInvalidatePropagation uintptr
	addr_QGraphicsAnchorLayout_addChildLayoutItem uintptr
	addr_QGraphicsAnchorLayout_setGraphicsItem uintptr
	addr_QGraphicsAnchorLayout_setOwnedByLayout uintptr
	addr_QGraphicsAnchorLayout_addAnchors2 uintptr
	addr_QGraphicsAnchorLayout_removeAt_default uintptr
	addr_QGraphicsAnchorLayout_setGeometry_default uintptr
	addr_QGraphicsAnchorLayout_count_default uintptr
	addr_QGraphicsAnchorLayout_itemAt_default uintptr
	addr_QGraphicsAnchorLayout_invalidate_default uintptr
	addr_QGraphicsAnchorLayout_sizeHint_default uintptr
	addr_QGraphicsAnchorLayout_getContentsMargins_default uintptr
	addr_QGraphicsAnchorLayout_updateGeometry_default uintptr
	addr_QGraphicsAnchorLayout_widgetEvent_default uintptr
	addr_QGraphicsAnchorLayout_isEmpty_default uintptr
	pg_QGraphicsAnchorLayout__set_override func(uintptr, int32)
	pg_QGraphicsAnchorLayout_delete func(uintptr)
)

func init() {
	symaddr(&addr_QGraphicsAnchorLayout_chain, "QGraphicsAnchorLayout_chain")
	symaddr(&addr_QGraphicsAnchorLayout_new, "QGraphicsAnchorLayout_new")
	symaddr(&addr_QGraphicsAnchorLayout_new2, "QGraphicsAnchorLayout_new2")
	symaddr(&addr_QGraphicsAnchorLayout_addAnchor, "QGraphicsAnchorLayout_addAnchor")
	symaddr(&addr_QGraphicsAnchorLayout_anchor, "QGraphicsAnchorLayout_anchor")
	symaddr(&addr_QGraphicsAnchorLayout_addCornerAnchors, "QGraphicsAnchorLayout_addCornerAnchors")
	symaddr(&addr_QGraphicsAnchorLayout_addAnchors, "QGraphicsAnchorLayout_addAnchors")
	reg(&pg_QGraphicsAnchorLayout_setHorizontalSpacing, "QGraphicsAnchorLayout_setHorizontalSpacing")
	reg(&pg_QGraphicsAnchorLayout_setVerticalSpacing, "QGraphicsAnchorLayout_setVerticalSpacing")
	reg(&pg_QGraphicsAnchorLayout_setSpacing, "QGraphicsAnchorLayout_setSpacing")
	reg(&pg_QGraphicsAnchorLayout_horizontalSpacing, "QGraphicsAnchorLayout_horizontalSpacing")
	reg(&pg_QGraphicsAnchorLayout_verticalSpacing, "QGraphicsAnchorLayout_verticalSpacing")
	symaddr(&addr_QGraphicsAnchorLayout_setInstantInvalidatePropagation, "QGraphicsAnchorLayout_setInstantInvalidatePropagation")
	symaddr(&addr_QGraphicsAnchorLayout_instantInvalidatePropagation, "QGraphicsAnchorLayout_instantInvalidatePropagation")
	symaddr(&addr_QGraphicsAnchorLayout_addChildLayoutItem, "QGraphicsAnchorLayout_addChildLayoutItem")
	symaddr(&addr_QGraphicsAnchorLayout_setGraphicsItem, "QGraphicsAnchorLayout_setGraphicsItem")
	symaddr(&addr_QGraphicsAnchorLayout_setOwnedByLayout, "QGraphicsAnchorLayout_setOwnedByLayout")
	symaddr(&addr_QGraphicsAnchorLayout_addAnchors2, "QGraphicsAnchorLayout_addAnchors2")
	symaddr(&addr_QGraphicsAnchorLayout_removeAt_default, "QGraphicsAnchorLayout_removeAt_default")
	symaddr(&addr_QGraphicsAnchorLayout_setGeometry_default, "QGraphicsAnchorLayout_setGeometry_default")
	symaddr(&addr_QGraphicsAnchorLayout_count_default, "QGraphicsAnchorLayout_count_default")
	symaddr(&addr_QGraphicsAnchorLayout_itemAt_default, "QGraphicsAnchorLayout_itemAt_default")
	symaddr(&addr_QGraphicsAnchorLayout_invalidate_default, "QGraphicsAnchorLayout_invalidate_default")
	symaddr(&addr_QGraphicsAnchorLayout_sizeHint_default, "QGraphicsAnchorLayout_sizeHint_default")
	symaddr(&addr_QGraphicsAnchorLayout_getContentsMargins_default, "QGraphicsAnchorLayout_getContentsMargins_default")
	symaddr(&addr_QGraphicsAnchorLayout_updateGeometry_default, "QGraphicsAnchorLayout_updateGeometry_default")
	symaddr(&addr_QGraphicsAnchorLayout_widgetEvent_default, "QGraphicsAnchorLayout_widgetEvent_default")
	symaddr(&addr_QGraphicsAnchorLayout_isEmpty_default, "QGraphicsAnchorLayout_isEmpty_default")
	reg(&pg_QGraphicsAnchorLayout__set_override, "QGraphicsAnchorLayout__set_override")
	reg(&pg_QGraphicsAnchorLayout_delete, "QGraphicsAnchorLayout_delete")
}

type QGraphicsAnchorLayout struct {
	h uintptr
	*QGraphicsLayout
}

func bindQGraphicsAnchorLayout(h uintptr) *QGraphicsAnchorLayout {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QGraphicsAnchorLayout_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QGraphicsAnchorLayout{h: h, QGraphicsLayout: &QGraphicsLayout{h: __c[0], QGraphicsLayoutItem: &QGraphicsLayoutItem{h: __c[1]}}}
}

func newownedQGraphicsAnchorLayout(h uintptr) *QGraphicsAnchorLayout {
	o := bindQGraphicsAnchorLayout(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QGraphicsAnchorLayout).Delete)
	}
	return o
}

func (this *QGraphicsAnchorLayout) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QGraphicsAnchorLayout) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQGraphicsAnchorLayout(h unsafe.Pointer) *QGraphicsAnchorLayout { return bindQGraphicsAnchorLayout(uintptr(h)) }
func UnsafeNewQGraphicsAnchorLayoutOwned(h unsafe.Pointer) *QGraphicsAnchorLayout { return newownedQGraphicsAnchorLayout(uintptr(h)) }

func (this *QGraphicsAnchorLayout) UnsafeRelease() {
	this.h = 0
	if this.QGraphicsLayout != nil {
		this.QGraphicsLayout.UnsafeRelease()
	}
}

func NewQGraphicsAnchorLayout() *QGraphicsAnchorLayout {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QGraphicsAnchorLayout_new, uintptr(__goid))
	__o := newownedQGraphicsAnchorLayout(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGraphicsAnchorLayout_virtIdx, QGraphicsAnchorLayout_enable)
	return __o
}

func NewQGraphicsAnchorLayout2(parent *QGraphicsLayoutItem) *QGraphicsAnchorLayout {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QGraphicsAnchorLayout_new2, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQGraphicsAnchorLayout(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGraphicsAnchorLayout_virtIdx, QGraphicsAnchorLayout_enable)
	return __o
}

func (this *QGraphicsAnchorLayout) AddAnchor(firstItem *QGraphicsLayoutItem, firstEdge AnchorPoint, secondItem *QGraphicsLayoutItem, secondEdge AnchorPoint) *QGraphicsAnchor {
	__r, _, _ := purego.SyscallN(addr_QGraphicsAnchorLayout_addAnchor, this.h, uintptr(firstItem.UnsafePointer()), uintptr(firstEdge), uintptr(secondItem.UnsafePointer()), uintptr(secondEdge))
	runtime.KeepAlive(this)
	runtime.KeepAlive(firstItem)
	runtime.KeepAlive(secondItem)
	return bindQGraphicsAnchor(__r)
}

func (this *QGraphicsAnchorLayout) Anchor(firstItem *QGraphicsLayoutItem, firstEdge AnchorPoint, secondItem *QGraphicsLayoutItem, secondEdge AnchorPoint) *QGraphicsAnchor {
	__r, _, _ := purego.SyscallN(addr_QGraphicsAnchorLayout_anchor, this.h, uintptr(firstItem.UnsafePointer()), uintptr(firstEdge), uintptr(secondItem.UnsafePointer()), uintptr(secondEdge))
	runtime.KeepAlive(this)
	runtime.KeepAlive(firstItem)
	runtime.KeepAlive(secondItem)
	return bindQGraphicsAnchor(__r)
}

func (this *QGraphicsAnchorLayout) AddCornerAnchors(firstItem *QGraphicsLayoutItem, firstCorner Corner, secondItem *QGraphicsLayoutItem, secondCorner Corner) {
	purego.SyscallN(addr_QGraphicsAnchorLayout_addCornerAnchors, this.h, uintptr(firstItem.UnsafePointer()), uintptr(firstCorner), uintptr(secondItem.UnsafePointer()), uintptr(secondCorner))
	runtime.KeepAlive(this)
	runtime.KeepAlive(firstItem)
	runtime.KeepAlive(secondItem)
}

func (this *QGraphicsAnchorLayout) AddAnchors(firstItem *QGraphicsLayoutItem, secondItem *QGraphicsLayoutItem) {
	purego.SyscallN(addr_QGraphicsAnchorLayout_addAnchors, this.h, uintptr(firstItem.UnsafePointer()), uintptr(secondItem.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(firstItem)
	runtime.KeepAlive(secondItem)
}

func (this *QGraphicsAnchorLayout) SetHorizontalSpacing(spacing float64) {
	pg_QGraphicsAnchorLayout_setHorizontalSpacing(this.h, spacing)
	runtime.KeepAlive(this)
}

func (this *QGraphicsAnchorLayout) SetVerticalSpacing(spacing float64) {
	pg_QGraphicsAnchorLayout_setVerticalSpacing(this.h, spacing)
	runtime.KeepAlive(this)
}

func (this *QGraphicsAnchorLayout) SetSpacing(spacing float64) {
	pg_QGraphicsAnchorLayout_setSpacing(this.h, spacing)
	runtime.KeepAlive(this)
}

func (this *QGraphicsAnchorLayout) HorizontalSpacing() float64 {
	__r := pg_QGraphicsAnchorLayout_horizontalSpacing(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QGraphicsAnchorLayout) VerticalSpacing() float64 {
	__r := pg_QGraphicsAnchorLayout_verticalSpacing(this.h)
	runtime.KeepAlive(this)
	return __r
}

func QGraphicsAnchorLayout_SetInstantInvalidatePropagation(enable bool) {
	purego.SyscallN(addr_QGraphicsAnchorLayout_setInstantInvalidatePropagation, uintptr(b2i(enable)))
}

func QGraphicsAnchorLayout_InstantInvalidatePropagation() bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsAnchorLayout_instantInvalidatePropagation)
	__r := int32(uint32(__r0))
	return __r != 0
}

func (this *QGraphicsAnchorLayout) AddChildLayoutItem(layoutItem *QGraphicsLayoutItem) {
	purego.SyscallN(addr_QGraphicsAnchorLayout_addChildLayoutItem, this.h, uintptr(layoutItem.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(layoutItem)
}

func (this *QGraphicsAnchorLayout) SetGraphicsItem(item *QGraphicsItem) {
	purego.SyscallN(addr_QGraphicsAnchorLayout_setGraphicsItem, this.h, uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QGraphicsAnchorLayout) SetOwnedByLayout(ownedByLayout bool) {
	purego.SyscallN(addr_QGraphicsAnchorLayout_setOwnedByLayout, this.h, uintptr(b2i(ownedByLayout)))
	runtime.KeepAlive(this)
}

func (this *QGraphicsAnchorLayout) AddAnchors2(firstItem *QGraphicsLayoutItem, secondItem *QGraphicsLayoutItem, orientations Orientation) {
	purego.SyscallN(addr_QGraphicsAnchorLayout_addAnchors2, this.h, uintptr(firstItem.UnsafePointer()), uintptr(secondItem.UnsafePointer()), uintptr(orientations))
	runtime.KeepAlive(this)
	runtime.KeepAlive(firstItem)
	runtime.KeepAlive(secondItem)
}

func (this *QGraphicsAnchorLayout) RemoveAtDefault(index int) {
	purego.SyscallN(addr_QGraphicsAnchorLayout_removeAt_default, this.h, uintptr(index))
	runtime.KeepAlive(this)
}

func (this *QGraphicsAnchorLayout) SetGeometryDefault(rect *QRectF) {
	purego.SyscallN(addr_QGraphicsAnchorLayout_setGeometry_default, this.h, uintptr(rect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
}

func (this *QGraphicsAnchorLayout) CountDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsAnchorLayout_count_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QGraphicsAnchorLayout) ItemAtDefault(index int) *QGraphicsLayoutItem {
	__r, _, _ := purego.SyscallN(addr_QGraphicsAnchorLayout_itemAt_default, this.h, uintptr(index))
	runtime.KeepAlive(this)
	return bindQGraphicsLayoutItem(__r)
}

func (this *QGraphicsAnchorLayout) InvalidateDefault() {
	purego.SyscallN(addr_QGraphicsAnchorLayout_invalidate_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsAnchorLayout) SizeHintDefault(which SizeHint, constraint *QSizeF) *QSizeF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsAnchorLayout_sizeHint_default, this.h, uintptr(which), uintptr(constraint.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(constraint)
	return newownedQSizeF(__r)
}

func (this *QGraphicsAnchorLayout) GetContentsMarginsDefault() (float64, float64, float64, float64) {
	var left_out float64
	var top_out float64
	var right_out float64
	var bottom_out float64
	purego.SyscallN(addr_QGraphicsAnchorLayout_getContentsMargins_default, this.h, uintptr(unsafe.Pointer(&left_out)), uintptr(unsafe.Pointer(&top_out)), uintptr(unsafe.Pointer(&right_out)), uintptr(unsafe.Pointer(&bottom_out)))
	runtime.KeepAlive(this)
	return left_out, top_out, right_out, bottom_out
}

func (this *QGraphicsAnchorLayout) UpdateGeometryDefault() {
	purego.SyscallN(addr_QGraphicsAnchorLayout_updateGeometry_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsAnchorLayout) WidgetEventDefault(e *QEvent) {
	purego.SyscallN(addr_QGraphicsAnchorLayout_widgetEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QGraphicsAnchorLayout) IsEmptyDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsAnchorLayout_isEmpty_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QGraphicsAnchorLayout_virtIdx(slot int64) int32 {
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

func QGraphicsAnchorLayout_enable(h uintptr, idx int32) { pg_QGraphicsAnchorLayout__set_override(h, idx) }

func (this *QGraphicsAnchorLayout) OnRemoveAt(fn func(int)) {
	registerVirt(this.UnsafeGoID(), 1230346564330554752, func(index int32) { fn(int(index)) }) // slot "removeAt"
}

func (this *QGraphicsAnchorLayout) OnSetGeometry(fn func(*QRectF)) {
	registerVirt(this.UnsafeGoID(), 5684820848726969227, func(rect uintptr) { fn(UnsafeNewQRectF(unsafe.Pointer(rect))) }) // slot "setGeometry"
}

func (this *QGraphicsAnchorLayout) OnCount(fn func() int) {
	registerVirt(this.UnsafeGoID(), -5627843058667511180, func() int32 { return int32(fn()) }) // slot "count"
}

func (this *QGraphicsAnchorLayout) OnItemAt(fn func(int) *QGraphicsLayoutItem) {
	registerVirt(this.UnsafeGoID(), 7508512214013391107, func(index int32) uintptr { return uintptr(fn(int(index)).UnsafePointer()) }) // slot "itemAt"
}

func (this *QGraphicsAnchorLayout) OnInvalidate(fn func()) {
	registerVirt(this.UnsafeGoID(), 6991341317990632368, func() { fn() }) // slot "invalidate"
}

func (this *QGraphicsAnchorLayout) OnSizeHint(fn func(SizeHint, *QSizeF) *QSizeF) {
	registerVirt(this.UnsafeGoID(), 729690261457684101, func(which int32, constraint uintptr) uintptr { return uintptr(fn(SizeHint(which), UnsafeNewQSizeF(unsafe.Pointer(constraint))).UnsafePointer()) }) // slot "sizeHint"
}

func (this *QGraphicsAnchorLayout) OnGetContentsMargins(fn func(*float64, *float64, *float64, *float64)) {
	registerVirt(this.UnsafeGoID(), 164760286202646972, func(left uintptr, top uintptr, right uintptr, bottom uintptr) { fn((*float64)(unsafe.Pointer(left)), (*float64)(unsafe.Pointer(top)), (*float64)(unsafe.Pointer(right)), (*float64)(unsafe.Pointer(bottom))) }) // slot "getContentsMargins"
}

func (this *QGraphicsAnchorLayout) OnUpdateGeometry(fn func()) {
	registerVirt(this.UnsafeGoID(), 8082455443689927668, func() { fn() }) // slot "updateGeometry"
}

func (this *QGraphicsAnchorLayout) OnWidgetEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 8535746785413991679, func(e uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(e))) }) // slot "widgetEvent"
}

func (this *QGraphicsAnchorLayout) OnIsEmpty(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -626822430251098498, func() int32 { return b2i(fn()) }) // slot "isEmpty"
}

func (this *QGraphicsAnchorLayout) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QGraphicsAnchorLayout_delete(this.h)
	}
	this.UnsafeRelease()
}
