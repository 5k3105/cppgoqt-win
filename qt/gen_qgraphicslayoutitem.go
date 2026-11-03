package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QGraphicsLayoutItem_setSizePolicy uintptr
	addr_QGraphicsLayoutItem_setSizePolicy2 uintptr
	addr_QGraphicsLayoutItem_sizePolicy uintptr
	addr_QGraphicsLayoutItem_setMinimumSize uintptr
	pg_QGraphicsLayoutItem_setMinimumSize2 func(uintptr, float64, float64)
	addr_QGraphicsLayoutItem_minimumSize uintptr
	pg_QGraphicsLayoutItem_setMinimumWidth func(uintptr, float64)
	pg_QGraphicsLayoutItem_minimumWidth func(uintptr) float64
	pg_QGraphicsLayoutItem_setMinimumHeight func(uintptr, float64)
	pg_QGraphicsLayoutItem_minimumHeight func(uintptr) float64
	addr_QGraphicsLayoutItem_setPreferredSize uintptr
	pg_QGraphicsLayoutItem_setPreferredSize2 func(uintptr, float64, float64)
	addr_QGraphicsLayoutItem_preferredSize uintptr
	pg_QGraphicsLayoutItem_setPreferredWidth func(uintptr, float64)
	pg_QGraphicsLayoutItem_preferredWidth func(uintptr) float64
	pg_QGraphicsLayoutItem_setPreferredHeight func(uintptr, float64)
	pg_QGraphicsLayoutItem_preferredHeight func(uintptr) float64
	addr_QGraphicsLayoutItem_setMaximumSize uintptr
	pg_QGraphicsLayoutItem_setMaximumSize2 func(uintptr, float64, float64)
	addr_QGraphicsLayoutItem_maximumSize uintptr
	pg_QGraphicsLayoutItem_setMaximumWidth func(uintptr, float64)
	pg_QGraphicsLayoutItem_maximumWidth func(uintptr) float64
	pg_QGraphicsLayoutItem_setMaximumHeight func(uintptr, float64)
	pg_QGraphicsLayoutItem_maximumHeight func(uintptr) float64
	addr_QGraphicsLayoutItem_setGeometry uintptr
	addr_QGraphicsLayoutItem_geometry uintptr
	addr_QGraphicsLayoutItem_getContentsMargins uintptr
	addr_QGraphicsLayoutItem_contentsRect uintptr
	addr_QGraphicsLayoutItem_effectiveSizeHint uintptr
	addr_QGraphicsLayoutItem_updateGeometry uintptr
	addr_QGraphicsLayoutItem_isEmpty uintptr
	addr_QGraphicsLayoutItem_parentLayoutItem uintptr
	addr_QGraphicsLayoutItem_setParentLayoutItem uintptr
	addr_QGraphicsLayoutItem_isLayout uintptr
	addr_QGraphicsLayoutItem_graphicsItem uintptr
	addr_QGraphicsLayoutItem_ownedByLayout uintptr
	addr_QGraphicsLayoutItem_setGraphicsItem uintptr
	addr_QGraphicsLayoutItem_setOwnedByLayout uintptr
	addr_QGraphicsLayoutItem_setSizePolicy3 uintptr
	addr_QGraphicsLayoutItem_effectiveSizeHint2 uintptr
	addr_QGraphicsLayoutItem_setGeometry_default uintptr
	addr_QGraphicsLayoutItem_getContentsMargins_default uintptr
	addr_QGraphicsLayoutItem_updateGeometry_default uintptr
	addr_QGraphicsLayoutItem_isEmpty_default uintptr
	pg_QGraphicsLayoutItem__set_override func(uintptr, int32)
	pg_QGraphicsLayoutItem__Go_new func(int64) uintptr
	pg_QGraphicsLayoutItem_delete func(uintptr)
)

func init() {
	symaddr(&addr_QGraphicsLayoutItem_setSizePolicy, "QGraphicsLayoutItem_setSizePolicy")
	symaddr(&addr_QGraphicsLayoutItem_setSizePolicy2, "QGraphicsLayoutItem_setSizePolicy2")
	symaddr(&addr_QGraphicsLayoutItem_sizePolicy, "QGraphicsLayoutItem_sizePolicy")
	symaddr(&addr_QGraphicsLayoutItem_setMinimumSize, "QGraphicsLayoutItem_setMinimumSize")
	reg(&pg_QGraphicsLayoutItem_setMinimumSize2, "QGraphicsLayoutItem_setMinimumSize2")
	symaddr(&addr_QGraphicsLayoutItem_minimumSize, "QGraphicsLayoutItem_minimumSize")
	reg(&pg_QGraphicsLayoutItem_setMinimumWidth, "QGraphicsLayoutItem_setMinimumWidth")
	reg(&pg_QGraphicsLayoutItem_minimumWidth, "QGraphicsLayoutItem_minimumWidth")
	reg(&pg_QGraphicsLayoutItem_setMinimumHeight, "QGraphicsLayoutItem_setMinimumHeight")
	reg(&pg_QGraphicsLayoutItem_minimumHeight, "QGraphicsLayoutItem_minimumHeight")
	symaddr(&addr_QGraphicsLayoutItem_setPreferredSize, "QGraphicsLayoutItem_setPreferredSize")
	reg(&pg_QGraphicsLayoutItem_setPreferredSize2, "QGraphicsLayoutItem_setPreferredSize2")
	symaddr(&addr_QGraphicsLayoutItem_preferredSize, "QGraphicsLayoutItem_preferredSize")
	reg(&pg_QGraphicsLayoutItem_setPreferredWidth, "QGraphicsLayoutItem_setPreferredWidth")
	reg(&pg_QGraphicsLayoutItem_preferredWidth, "QGraphicsLayoutItem_preferredWidth")
	reg(&pg_QGraphicsLayoutItem_setPreferredHeight, "QGraphicsLayoutItem_setPreferredHeight")
	reg(&pg_QGraphicsLayoutItem_preferredHeight, "QGraphicsLayoutItem_preferredHeight")
	symaddr(&addr_QGraphicsLayoutItem_setMaximumSize, "QGraphicsLayoutItem_setMaximumSize")
	reg(&pg_QGraphicsLayoutItem_setMaximumSize2, "QGraphicsLayoutItem_setMaximumSize2")
	symaddr(&addr_QGraphicsLayoutItem_maximumSize, "QGraphicsLayoutItem_maximumSize")
	reg(&pg_QGraphicsLayoutItem_setMaximumWidth, "QGraphicsLayoutItem_setMaximumWidth")
	reg(&pg_QGraphicsLayoutItem_maximumWidth, "QGraphicsLayoutItem_maximumWidth")
	reg(&pg_QGraphicsLayoutItem_setMaximumHeight, "QGraphicsLayoutItem_setMaximumHeight")
	reg(&pg_QGraphicsLayoutItem_maximumHeight, "QGraphicsLayoutItem_maximumHeight")
	symaddr(&addr_QGraphicsLayoutItem_setGeometry, "QGraphicsLayoutItem_setGeometry")
	symaddr(&addr_QGraphicsLayoutItem_geometry, "QGraphicsLayoutItem_geometry")
	symaddr(&addr_QGraphicsLayoutItem_getContentsMargins, "QGraphicsLayoutItem_getContentsMargins")
	symaddr(&addr_QGraphicsLayoutItem_contentsRect, "QGraphicsLayoutItem_contentsRect")
	symaddr(&addr_QGraphicsLayoutItem_effectiveSizeHint, "QGraphicsLayoutItem_effectiveSizeHint")
	symaddr(&addr_QGraphicsLayoutItem_updateGeometry, "QGraphicsLayoutItem_updateGeometry")
	symaddr(&addr_QGraphicsLayoutItem_isEmpty, "QGraphicsLayoutItem_isEmpty")
	symaddr(&addr_QGraphicsLayoutItem_parentLayoutItem, "QGraphicsLayoutItem_parentLayoutItem")
	symaddr(&addr_QGraphicsLayoutItem_setParentLayoutItem, "QGraphicsLayoutItem_setParentLayoutItem")
	symaddr(&addr_QGraphicsLayoutItem_isLayout, "QGraphicsLayoutItem_isLayout")
	symaddr(&addr_QGraphicsLayoutItem_graphicsItem, "QGraphicsLayoutItem_graphicsItem")
	symaddr(&addr_QGraphicsLayoutItem_ownedByLayout, "QGraphicsLayoutItem_ownedByLayout")
	symaddr(&addr_QGraphicsLayoutItem_setGraphicsItem, "QGraphicsLayoutItem_setGraphicsItem")
	symaddr(&addr_QGraphicsLayoutItem_setOwnedByLayout, "QGraphicsLayoutItem_setOwnedByLayout")
	symaddr(&addr_QGraphicsLayoutItem_setSizePolicy3, "QGraphicsLayoutItem_setSizePolicy3")
	symaddr(&addr_QGraphicsLayoutItem_effectiveSizeHint2, "QGraphicsLayoutItem_effectiveSizeHint2")
	symaddr(&addr_QGraphicsLayoutItem_setGeometry_default, "QGraphicsLayoutItem_setGeometry_default")
	symaddr(&addr_QGraphicsLayoutItem_getContentsMargins_default, "QGraphicsLayoutItem_getContentsMargins_default")
	symaddr(&addr_QGraphicsLayoutItem_updateGeometry_default, "QGraphicsLayoutItem_updateGeometry_default")
	symaddr(&addr_QGraphicsLayoutItem_isEmpty_default, "QGraphicsLayoutItem_isEmpty_default")
	reg(&pg_QGraphicsLayoutItem__set_override, "QGraphicsLayoutItem__set_override")
	reg(&pg_QGraphicsLayoutItem__Go_new, "QGraphicsLayoutItem__Go_new")
	reg(&pg_QGraphicsLayoutItem_delete, "QGraphicsLayoutItem_delete")
}

type QGraphicsLayoutItem struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QGraphicsLayoutItem) UnsafeGoID() int64     { return this.vid }
func (this *QGraphicsLayoutItem) UnsafeSetGoID(id int64) { this.vid = id }

func bindQGraphicsLayoutItem(h uintptr) *QGraphicsLayoutItem {
	if h == 0 {
		return nil
	}
	return &QGraphicsLayoutItem{h: h}
}

func newownedQGraphicsLayoutItem(h uintptr) *QGraphicsLayoutItem {
	o := bindQGraphicsLayoutItem(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QGraphicsLayoutItem).Delete)
	}
	return o
}

func (this *QGraphicsLayoutItem) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QGraphicsLayoutItem) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQGraphicsLayoutItem(h unsafe.Pointer) *QGraphicsLayoutItem { return bindQGraphicsLayoutItem(uintptr(h)) }
func UnsafeNewQGraphicsLayoutItemOwned(h unsafe.Pointer) *QGraphicsLayoutItem { return newownedQGraphicsLayoutItem(uintptr(h)) }

func (this *QGraphicsLayoutItem) UnsafeRelease() { this.h = 0 }

func (this *QGraphicsLayoutItem) SetSizePolicy(policy *QSizePolicy) {
	purego.SyscallN(addr_QGraphicsLayoutItem_setSizePolicy, this.h, uintptr(policy.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(policy)
}

func (this *QGraphicsLayoutItem) SetSizePolicy2(hPolicy QSizePolicy__Policy, vPolicy QSizePolicy__Policy) {
	purego.SyscallN(addr_QGraphicsLayoutItem_setSizePolicy2, this.h, uintptr(hPolicy), uintptr(vPolicy))
	runtime.KeepAlive(this)
}

func (this *QGraphicsLayoutItem) SizePolicy() *QSizePolicy {
	__r, _, _ := purego.SyscallN(addr_QGraphicsLayoutItem_sizePolicy, this.h)
	runtime.KeepAlive(this)
	return newownedQSizePolicy(__r)
}

func (this *QGraphicsLayoutItem) SetMinimumSize(size *QSizeF) {
	purego.SyscallN(addr_QGraphicsLayoutItem_setMinimumSize, this.h, uintptr(size.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(size)
}

func (this *QGraphicsLayoutItem) SetMinimumSize2(w float64, h float64) {
	pg_QGraphicsLayoutItem_setMinimumSize2(this.h, w, h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsLayoutItem) MinimumSize() *QSizeF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsLayoutItem_minimumSize, this.h)
	runtime.KeepAlive(this)
	return newownedQSizeF(__r)
}

func (this *QGraphicsLayoutItem) SetMinimumWidth(width float64) {
	pg_QGraphicsLayoutItem_setMinimumWidth(this.h, width)
	runtime.KeepAlive(this)
}

func (this *QGraphicsLayoutItem) MinimumWidth() float64 {
	__r := pg_QGraphicsLayoutItem_minimumWidth(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QGraphicsLayoutItem) SetMinimumHeight(height float64) {
	pg_QGraphicsLayoutItem_setMinimumHeight(this.h, height)
	runtime.KeepAlive(this)
}

func (this *QGraphicsLayoutItem) MinimumHeight() float64 {
	__r := pg_QGraphicsLayoutItem_minimumHeight(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QGraphicsLayoutItem) SetPreferredSize(size *QSizeF) {
	purego.SyscallN(addr_QGraphicsLayoutItem_setPreferredSize, this.h, uintptr(size.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(size)
}

func (this *QGraphicsLayoutItem) SetPreferredSize2(w float64, h float64) {
	pg_QGraphicsLayoutItem_setPreferredSize2(this.h, w, h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsLayoutItem) PreferredSize() *QSizeF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsLayoutItem_preferredSize, this.h)
	runtime.KeepAlive(this)
	return newownedQSizeF(__r)
}

func (this *QGraphicsLayoutItem) SetPreferredWidth(width float64) {
	pg_QGraphicsLayoutItem_setPreferredWidth(this.h, width)
	runtime.KeepAlive(this)
}

func (this *QGraphicsLayoutItem) PreferredWidth() float64 {
	__r := pg_QGraphicsLayoutItem_preferredWidth(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QGraphicsLayoutItem) SetPreferredHeight(height float64) {
	pg_QGraphicsLayoutItem_setPreferredHeight(this.h, height)
	runtime.KeepAlive(this)
}

func (this *QGraphicsLayoutItem) PreferredHeight() float64 {
	__r := pg_QGraphicsLayoutItem_preferredHeight(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QGraphicsLayoutItem) SetMaximumSize(size *QSizeF) {
	purego.SyscallN(addr_QGraphicsLayoutItem_setMaximumSize, this.h, uintptr(size.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(size)
}

func (this *QGraphicsLayoutItem) SetMaximumSize2(w float64, h float64) {
	pg_QGraphicsLayoutItem_setMaximumSize2(this.h, w, h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsLayoutItem) MaximumSize() *QSizeF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsLayoutItem_maximumSize, this.h)
	runtime.KeepAlive(this)
	return newownedQSizeF(__r)
}

func (this *QGraphicsLayoutItem) SetMaximumWidth(width float64) {
	pg_QGraphicsLayoutItem_setMaximumWidth(this.h, width)
	runtime.KeepAlive(this)
}

func (this *QGraphicsLayoutItem) MaximumWidth() float64 {
	__r := pg_QGraphicsLayoutItem_maximumWidth(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QGraphicsLayoutItem) SetMaximumHeight(height float64) {
	pg_QGraphicsLayoutItem_setMaximumHeight(this.h, height)
	runtime.KeepAlive(this)
}

func (this *QGraphicsLayoutItem) MaximumHeight() float64 {
	__r := pg_QGraphicsLayoutItem_maximumHeight(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QGraphicsLayoutItem) SetGeometry(rect *QRectF) {
	purego.SyscallN(addr_QGraphicsLayoutItem_setGeometry, this.h, uintptr(rect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
}

func (this *QGraphicsLayoutItem) Geometry() *QRectF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsLayoutItem_geometry, this.h)
	runtime.KeepAlive(this)
	return newownedQRectF(__r)
}

func (this *QGraphicsLayoutItem) GetContentsMargins() (float64, float64, float64, float64) {
	var left_out float64
	var top_out float64
	var right_out float64
	var bottom_out float64
	purego.SyscallN(addr_QGraphicsLayoutItem_getContentsMargins, this.h, uintptr(unsafe.Pointer(&left_out)), uintptr(unsafe.Pointer(&top_out)), uintptr(unsafe.Pointer(&right_out)), uintptr(unsafe.Pointer(&bottom_out)))
	runtime.KeepAlive(this)
	return left_out, top_out, right_out, bottom_out
}

func (this *QGraphicsLayoutItem) ContentsRect() *QRectF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsLayoutItem_contentsRect, this.h)
	runtime.KeepAlive(this)
	return newownedQRectF(__r)
}

func (this *QGraphicsLayoutItem) EffectiveSizeHint(which SizeHint) *QSizeF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsLayoutItem_effectiveSizeHint, this.h, uintptr(which))
	runtime.KeepAlive(this)
	return newownedQSizeF(__r)
}

func (this *QGraphicsLayoutItem) UpdateGeometry() {
	purego.SyscallN(addr_QGraphicsLayoutItem_updateGeometry, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsLayoutItem) IsEmpty() bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsLayoutItem_isEmpty, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGraphicsLayoutItem) ParentLayoutItem() *QGraphicsLayoutItem {
	__r, _, _ := purego.SyscallN(addr_QGraphicsLayoutItem_parentLayoutItem, this.h)
	runtime.KeepAlive(this)
	return bindQGraphicsLayoutItem(__r)
}

func (this *QGraphicsLayoutItem) SetParentLayoutItem(parent *QGraphicsLayoutItem) {
	purego.SyscallN(addr_QGraphicsLayoutItem_setParentLayoutItem, this.h, uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
}

func (this *QGraphicsLayoutItem) IsLayout() bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsLayoutItem_isLayout, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGraphicsLayoutItem) GraphicsItem() *QGraphicsItem {
	__r, _, _ := purego.SyscallN(addr_QGraphicsLayoutItem_graphicsItem, this.h)
	runtime.KeepAlive(this)
	return bindQGraphicsItem(__r)
}

func (this *QGraphicsLayoutItem) OwnedByLayout() bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsLayoutItem_ownedByLayout, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGraphicsLayoutItem) SetGraphicsItem(item *QGraphicsItem) {
	purego.SyscallN(addr_QGraphicsLayoutItem_setGraphicsItem, this.h, uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QGraphicsLayoutItem) SetOwnedByLayout(ownedByLayout bool) {
	purego.SyscallN(addr_QGraphicsLayoutItem_setOwnedByLayout, this.h, uintptr(b2i(ownedByLayout)))
	runtime.KeepAlive(this)
}

func (this *QGraphicsLayoutItem) SetSizePolicy3(hPolicy QSizePolicy__Policy, vPolicy QSizePolicy__Policy, controlType QSizePolicy__ControlType) {
	purego.SyscallN(addr_QGraphicsLayoutItem_setSizePolicy3, this.h, uintptr(hPolicy), uintptr(vPolicy), uintptr(controlType))
	runtime.KeepAlive(this)
}

func (this *QGraphicsLayoutItem) EffectiveSizeHint2(which SizeHint, constraint *QSizeF) *QSizeF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsLayoutItem_effectiveSizeHint2, this.h, uintptr(which), uintptr(constraint.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(constraint)
	return newownedQSizeF(__r)
}

func (this *QGraphicsLayoutItem) SetGeometryDefault(rect *QRectF) {
	purego.SyscallN(addr_QGraphicsLayoutItem_setGeometry_default, this.h, uintptr(rect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
}

func (this *QGraphicsLayoutItem) GetContentsMarginsDefault() (float64, float64, float64, float64) {
	var left_out float64
	var top_out float64
	var right_out float64
	var bottom_out float64
	purego.SyscallN(addr_QGraphicsLayoutItem_getContentsMargins_default, this.h, uintptr(unsafe.Pointer(&left_out)), uintptr(unsafe.Pointer(&top_out)), uintptr(unsafe.Pointer(&right_out)), uintptr(unsafe.Pointer(&bottom_out)))
	runtime.KeepAlive(this)
	return left_out, top_out, right_out, bottom_out
}

func (this *QGraphicsLayoutItem) UpdateGeometryDefault() {
	purego.SyscallN(addr_QGraphicsLayoutItem_updateGeometry_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsLayoutItem) IsEmptyDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsLayoutItem_isEmpty_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QGraphicsLayoutItem_virtIdx(slot int64) int32 {
	switch slot {
	case 729690261457684101:
		return 0
	case 5684820848726969227:
		return 1
	case 164760286202646972:
		return 2
	case 8082455443689927668:
		return 3
	case -626822430251098498:
		return 4
	}
	return -1
}

func QGraphicsLayoutItem_enable(h uintptr, idx int32) { pg_QGraphicsLayoutItem__set_override(h, idx) }

func NewQGraphicsLayoutItem() *QGraphicsLayoutItem {
	id := nextVirtID()
	o := newownedQGraphicsLayoutItem(pg_QGraphicsLayoutItem__Go_new(id))
	o.UnsafeSetGoID(id)
	registerInstance(id, o.h, QGraphicsLayoutItem_virtIdx, QGraphicsLayoutItem_enable)
	return o
}

func (this *QGraphicsLayoutItem) OnSizeHint(fn func(SizeHint, *QSizeF) *QSizeF) {
	registerVirt(this.UnsafeGoID(), 729690261457684101, func(which int32, constraint uintptr) uintptr { return uintptr(fn(SizeHint(which), UnsafeNewQSizeF(unsafe.Pointer(constraint))).UnsafePointer()) }) // slot "sizeHint"
}

func (this *QGraphicsLayoutItem) OnSetGeometry(fn func(*QRectF)) {
	registerVirt(this.UnsafeGoID(), 5684820848726969227, func(rect uintptr) { fn(UnsafeNewQRectF(unsafe.Pointer(rect))) }) // slot "setGeometry"
}

func (this *QGraphicsLayoutItem) OnGetContentsMargins(fn func(*float64, *float64, *float64, *float64)) {
	registerVirt(this.UnsafeGoID(), 164760286202646972, func(left uintptr, top uintptr, right uintptr, bottom uintptr) { fn((*float64)(unsafe.Pointer(left)), (*float64)(unsafe.Pointer(top)), (*float64)(unsafe.Pointer(right)), (*float64)(unsafe.Pointer(bottom))) }) // slot "getContentsMargins"
}

func (this *QGraphicsLayoutItem) OnUpdateGeometry(fn func()) {
	registerVirt(this.UnsafeGoID(), 8082455443689927668, func() { fn() }) // slot "updateGeometry"
}

func (this *QGraphicsLayoutItem) OnIsEmpty(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -626822430251098498, func() int32 { return b2i(fn()) }) // slot "isEmpty"
}

func (this *QGraphicsLayoutItem) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QGraphicsLayoutItem_delete(this.h)
	}
	this.UnsafeRelease()
}
