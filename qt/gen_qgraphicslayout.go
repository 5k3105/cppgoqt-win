package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QGraphicsLayout_base func(uintptr) uintptr
	pg_QGraphicsLayout_setContentsMargins func(uintptr, float64, float64, float64, float64)
	addr_QGraphicsLayout_activate uintptr
	addr_QGraphicsLayout_isActivated uintptr
	addr_QGraphicsLayout_invalidate uintptr
	addr_QGraphicsLayout_widgetEvent uintptr
	addr_QGraphicsLayout_count uintptr
	addr_QGraphicsLayout_itemAt uintptr
	addr_QGraphicsLayout_removeAt uintptr
	addr_QGraphicsLayout_setInstantInvalidatePropagation uintptr
	addr_QGraphicsLayout_instantInvalidatePropagation uintptr
	addr_QGraphicsLayout_addChildLayoutItem uintptr
	addr_QGraphicsLayout_setGraphicsItem uintptr
	addr_QGraphicsLayout_setOwnedByLayout uintptr
	addr_QGraphicsLayout_invalidate_default uintptr
	addr_QGraphicsLayout_widgetEvent_default uintptr
	addr_QGraphicsLayout_getContentsMargins_default uintptr
	addr_QGraphicsLayout_updateGeometry_default uintptr
	addr_QGraphicsLayout_setGeometry_default uintptr
	addr_QGraphicsLayout_isEmpty_default uintptr
	pg_QGraphicsLayout__set_override func(uintptr, int32)
	pg_QGraphicsLayout__Go_new func(int64) uintptr
	pg_QGraphicsLayout_delete func(uintptr)
)

func init() {
	reg(&pg_QGraphicsLayout_base, "QGraphicsLayout_base")
	reg(&pg_QGraphicsLayout_setContentsMargins, "QGraphicsLayout_setContentsMargins")
	symaddr(&addr_QGraphicsLayout_activate, "QGraphicsLayout_activate")
	symaddr(&addr_QGraphicsLayout_isActivated, "QGraphicsLayout_isActivated")
	symaddr(&addr_QGraphicsLayout_invalidate, "QGraphicsLayout_invalidate")
	symaddr(&addr_QGraphicsLayout_widgetEvent, "QGraphicsLayout_widgetEvent")
	symaddr(&addr_QGraphicsLayout_count, "QGraphicsLayout_count")
	symaddr(&addr_QGraphicsLayout_itemAt, "QGraphicsLayout_itemAt")
	symaddr(&addr_QGraphicsLayout_removeAt, "QGraphicsLayout_removeAt")
	symaddr(&addr_QGraphicsLayout_setInstantInvalidatePropagation, "QGraphicsLayout_setInstantInvalidatePropagation")
	symaddr(&addr_QGraphicsLayout_instantInvalidatePropagation, "QGraphicsLayout_instantInvalidatePropagation")
	symaddr(&addr_QGraphicsLayout_addChildLayoutItem, "QGraphicsLayout_addChildLayoutItem")
	symaddr(&addr_QGraphicsLayout_setGraphicsItem, "QGraphicsLayout_setGraphicsItem")
	symaddr(&addr_QGraphicsLayout_setOwnedByLayout, "QGraphicsLayout_setOwnedByLayout")
	symaddr(&addr_QGraphicsLayout_invalidate_default, "QGraphicsLayout_invalidate_default")
	symaddr(&addr_QGraphicsLayout_widgetEvent_default, "QGraphicsLayout_widgetEvent_default")
	symaddr(&addr_QGraphicsLayout_getContentsMargins_default, "QGraphicsLayout_getContentsMargins_default")
	symaddr(&addr_QGraphicsLayout_updateGeometry_default, "QGraphicsLayout_updateGeometry_default")
	symaddr(&addr_QGraphicsLayout_setGeometry_default, "QGraphicsLayout_setGeometry_default")
	symaddr(&addr_QGraphicsLayout_isEmpty_default, "QGraphicsLayout_isEmpty_default")
	reg(&pg_QGraphicsLayout__set_override, "QGraphicsLayout__set_override")
	reg(&pg_QGraphicsLayout__Go_new, "QGraphicsLayout__Go_new")
	reg(&pg_QGraphicsLayout_delete, "QGraphicsLayout_delete")
}

type QGraphicsLayout struct {
	h uintptr
	*QGraphicsLayoutItem
}

func bindQGraphicsLayout(h uintptr) *QGraphicsLayout {
	if h == 0 {
		return nil
	}
	return &QGraphicsLayout{h: h, QGraphicsLayoutItem: bindQGraphicsLayoutItem(pg_QGraphicsLayout_base(h))}
}

func newownedQGraphicsLayout(h uintptr) *QGraphicsLayout {
	o := bindQGraphicsLayout(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QGraphicsLayout).Delete)
	}
	return o
}

func (this *QGraphicsLayout) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QGraphicsLayout) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQGraphicsLayout(h unsafe.Pointer) *QGraphicsLayout { return bindQGraphicsLayout(uintptr(h)) }
func UnsafeNewQGraphicsLayoutOwned(h unsafe.Pointer) *QGraphicsLayout { return newownedQGraphicsLayout(uintptr(h)) }

func (this *QGraphicsLayout) UnsafeRelease() {
	this.h = 0
	if this.QGraphicsLayoutItem != nil {
		this.QGraphicsLayoutItem.UnsafeRelease()
	}
}

func (this *QGraphicsLayout) SetContentsMargins(left float64, top float64, right float64, bottom float64) {
	pg_QGraphicsLayout_setContentsMargins(this.h, left, top, right, bottom)
	runtime.KeepAlive(this)
}

func (this *QGraphicsLayout) Activate() {
	purego.SyscallN(addr_QGraphicsLayout_activate, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsLayout) IsActivated() bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsLayout_isActivated, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGraphicsLayout) Invalidate() {
	purego.SyscallN(addr_QGraphicsLayout_invalidate, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsLayout) WidgetEvent(e *QEvent) {
	purego.SyscallN(addr_QGraphicsLayout_widgetEvent, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QGraphicsLayout) Count() int {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsLayout_count, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QGraphicsLayout) ItemAt(i int) *QGraphicsLayoutItem {
	__r, _, _ := purego.SyscallN(addr_QGraphicsLayout_itemAt, this.h, uintptr(i))
	runtime.KeepAlive(this)
	return bindQGraphicsLayoutItem(__r)
}

func (this *QGraphicsLayout) RemoveAt(index int) {
	purego.SyscallN(addr_QGraphicsLayout_removeAt, this.h, uintptr(index))
	runtime.KeepAlive(this)
}

func QGraphicsLayout_SetInstantInvalidatePropagation(enable bool) {
	purego.SyscallN(addr_QGraphicsLayout_setInstantInvalidatePropagation, uintptr(b2i(enable)))
}

func QGraphicsLayout_InstantInvalidatePropagation() bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsLayout_instantInvalidatePropagation)
	__r := int32(uint32(__r0))
	return __r != 0
}

func (this *QGraphicsLayout) AddChildLayoutItem(layoutItem *QGraphicsLayoutItem) {
	purego.SyscallN(addr_QGraphicsLayout_addChildLayoutItem, this.h, uintptr(layoutItem.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(layoutItem)
}

func (this *QGraphicsLayout) SetGraphicsItem(item *QGraphicsItem) {
	purego.SyscallN(addr_QGraphicsLayout_setGraphicsItem, this.h, uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QGraphicsLayout) SetOwnedByLayout(ownedByLayout bool) {
	purego.SyscallN(addr_QGraphicsLayout_setOwnedByLayout, this.h, uintptr(b2i(ownedByLayout)))
	runtime.KeepAlive(this)
}

func (this *QGraphicsLayout) InvalidateDefault() {
	purego.SyscallN(addr_QGraphicsLayout_invalidate_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsLayout) WidgetEventDefault(e *QEvent) {
	purego.SyscallN(addr_QGraphicsLayout_widgetEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QGraphicsLayout) GetContentsMarginsDefault() (float64, float64, float64, float64) {
	var left_out float64
	var top_out float64
	var right_out float64
	var bottom_out float64
	purego.SyscallN(addr_QGraphicsLayout_getContentsMargins_default, this.h, uintptr(unsafe.Pointer(&left_out)), uintptr(unsafe.Pointer(&top_out)), uintptr(unsafe.Pointer(&right_out)), uintptr(unsafe.Pointer(&bottom_out)))
	runtime.KeepAlive(this)
	return left_out, top_out, right_out, bottom_out
}

func (this *QGraphicsLayout) UpdateGeometryDefault() {
	purego.SyscallN(addr_QGraphicsLayout_updateGeometry_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsLayout) SetGeometryDefault(rect *QRectF) {
	purego.SyscallN(addr_QGraphicsLayout_setGeometry_default, this.h, uintptr(rect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
}

func (this *QGraphicsLayout) IsEmptyDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsLayout_isEmpty_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QGraphicsLayout_virtIdx(slot int64) int32 {
	switch slot {
	case 729690261457684101:
		return 0
	case -5627843058667511180:
		return 1
	case 7508512214013391107:
		return 2
	case 1230346564330554752:
		return 3
	case 164760286202646972:
		return 4
	case 6991341317990632368:
		return 5
	case 8082455443689927668:
		return 6
	case 8535746785413991679:
		return 7
	case 5684820848726969227:
		return 8
	case -626822430251098498:
		return 9
	}
	return -1
}

func QGraphicsLayout_enable(h uintptr, idx int32) { pg_QGraphicsLayout__set_override(h, idx) }

func NewQGraphicsLayout() *QGraphicsLayout {
	id := nextVirtID()
	o := newownedQGraphicsLayout(pg_QGraphicsLayout__Go_new(id))
	o.UnsafeSetGoID(id)
	registerInstance(id, o.h, QGraphicsLayout_virtIdx, QGraphicsLayout_enable)
	return o
}

func (this *QGraphicsLayout) OnSizeHint(fn func(SizeHint, *QSizeF) *QSizeF) {
	registerVirt(this.UnsafeGoID(), 729690261457684101, func(which int32, constraint uintptr) uintptr { return uintptr(fn(SizeHint(which), UnsafeNewQSizeF(unsafe.Pointer(constraint))).UnsafePointer()) }) // slot "sizeHint"
}

func (this *QGraphicsLayout) OnCount(fn func() int) {
	registerVirt(this.UnsafeGoID(), -5627843058667511180, func() int32 { return int32(fn()) }) // slot "count"
}

func (this *QGraphicsLayout) OnItemAt(fn func(int) *QGraphicsLayoutItem) {
	registerVirt(this.UnsafeGoID(), 7508512214013391107, func(i int32) uintptr { return uintptr(fn(int(i)).UnsafePointer()) }) // slot "itemAt"
}

func (this *QGraphicsLayout) OnRemoveAt(fn func(int)) {
	registerVirt(this.UnsafeGoID(), 1230346564330554752, func(index int32) { fn(int(index)) }) // slot "removeAt"
}

func (this *QGraphicsLayout) OnGetContentsMargins(fn func(*float64, *float64, *float64, *float64)) {
	registerVirt(this.UnsafeGoID(), 164760286202646972, func(left uintptr, top uintptr, right uintptr, bottom uintptr) { fn((*float64)(unsafe.Pointer(left)), (*float64)(unsafe.Pointer(top)), (*float64)(unsafe.Pointer(right)), (*float64)(unsafe.Pointer(bottom))) }) // slot "getContentsMargins"
}

func (this *QGraphicsLayout) OnInvalidate(fn func()) {
	registerVirt(this.UnsafeGoID(), 6991341317990632368, func() { fn() }) // slot "invalidate"
}

func (this *QGraphicsLayout) OnUpdateGeometry(fn func()) {
	registerVirt(this.UnsafeGoID(), 8082455443689927668, func() { fn() }) // slot "updateGeometry"
}

func (this *QGraphicsLayout) OnWidgetEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 8535746785413991679, func(e uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(e))) }) // slot "widgetEvent"
}

func (this *QGraphicsLayout) OnSetGeometry(fn func(*QRectF)) {
	registerVirt(this.UnsafeGoID(), 5684820848726969227, func(rect uintptr) { fn(UnsafeNewQRectF(unsafe.Pointer(rect))) }) // slot "setGeometry"
}

func (this *QGraphicsLayout) OnIsEmpty(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -626822430251098498, func() int32 { return b2i(fn()) }) // slot "isEmpty"
}

func (this *QGraphicsLayout) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QGraphicsLayout_delete(this.h)
	}
	this.UnsafeRelease()
}
