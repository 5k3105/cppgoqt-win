package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QLayoutItem_sizeHint uintptr
	addr_QLayoutItem_minimumSize uintptr
	addr_QLayoutItem_maximumSize uintptr
	addr_QLayoutItem_expandingDirections uintptr
	addr_QLayoutItem_setGeometry uintptr
	addr_QLayoutItem_geometry uintptr
	addr_QLayoutItem_isEmpty uintptr
	addr_QLayoutItem_hasHeightForWidth uintptr
	addr_QLayoutItem_heightForWidth uintptr
	addr_QLayoutItem_minimumHeightForWidth uintptr
	addr_QLayoutItem_invalidate uintptr
	addr_QLayoutItem_widget uintptr
	addr_QLayoutItem_layout uintptr
	addr_QLayoutItem_spacerItem uintptr
	addr_QLayoutItem_alignment uintptr
	addr_QLayoutItem_setAlignment uintptr
	addr_QLayoutItem_controlTypes uintptr
	addr_QLayoutItem_OperatorAssign uintptr
	addr_QLayoutItem_hasHeightForWidth_default uintptr
	addr_QLayoutItem_heightForWidth_default uintptr
	addr_QLayoutItem_minimumHeightForWidth_default uintptr
	addr_QLayoutItem_invalidate_default uintptr
	addr_QLayoutItem_widget_default uintptr
	addr_QLayoutItem_layout_default uintptr
	addr_QLayoutItem_spacerItem_default uintptr
	addr_QLayoutItem_controlTypes_default uintptr
	pg_QLayoutItem__set_override func(uintptr, int32)
	pg_QLayoutItem__Go_new func(int64) uintptr
	pg_QLayoutItem_delete func(uintptr)
)

func init() {
	symaddr(&addr_QLayoutItem_sizeHint, "QLayoutItem_sizeHint")
	symaddr(&addr_QLayoutItem_minimumSize, "QLayoutItem_minimumSize")
	symaddr(&addr_QLayoutItem_maximumSize, "QLayoutItem_maximumSize")
	symaddr(&addr_QLayoutItem_expandingDirections, "QLayoutItem_expandingDirections")
	symaddr(&addr_QLayoutItem_setGeometry, "QLayoutItem_setGeometry")
	symaddr(&addr_QLayoutItem_geometry, "QLayoutItem_geometry")
	symaddr(&addr_QLayoutItem_isEmpty, "QLayoutItem_isEmpty")
	symaddr(&addr_QLayoutItem_hasHeightForWidth, "QLayoutItem_hasHeightForWidth")
	symaddr(&addr_QLayoutItem_heightForWidth, "QLayoutItem_heightForWidth")
	symaddr(&addr_QLayoutItem_minimumHeightForWidth, "QLayoutItem_minimumHeightForWidth")
	symaddr(&addr_QLayoutItem_invalidate, "QLayoutItem_invalidate")
	symaddr(&addr_QLayoutItem_widget, "QLayoutItem_widget")
	symaddr(&addr_QLayoutItem_layout, "QLayoutItem_layout")
	symaddr(&addr_QLayoutItem_spacerItem, "QLayoutItem_spacerItem")
	symaddr(&addr_QLayoutItem_alignment, "QLayoutItem_alignment")
	symaddr(&addr_QLayoutItem_setAlignment, "QLayoutItem_setAlignment")
	symaddr(&addr_QLayoutItem_controlTypes, "QLayoutItem_controlTypes")
	symaddr(&addr_QLayoutItem_OperatorAssign, "QLayoutItem_OperatorAssign")
	symaddr(&addr_QLayoutItem_hasHeightForWidth_default, "QLayoutItem_hasHeightForWidth_default")
	symaddr(&addr_QLayoutItem_heightForWidth_default, "QLayoutItem_heightForWidth_default")
	symaddr(&addr_QLayoutItem_minimumHeightForWidth_default, "QLayoutItem_minimumHeightForWidth_default")
	symaddr(&addr_QLayoutItem_invalidate_default, "QLayoutItem_invalidate_default")
	symaddr(&addr_QLayoutItem_widget_default, "QLayoutItem_widget_default")
	symaddr(&addr_QLayoutItem_layout_default, "QLayoutItem_layout_default")
	symaddr(&addr_QLayoutItem_spacerItem_default, "QLayoutItem_spacerItem_default")
	symaddr(&addr_QLayoutItem_controlTypes_default, "QLayoutItem_controlTypes_default")
	reg(&pg_QLayoutItem__set_override, "QLayoutItem__set_override")
	reg(&pg_QLayoutItem__Go_new, "QLayoutItem__Go_new")
	reg(&pg_QLayoutItem_delete, "QLayoutItem_delete")
}

type QLayoutItem struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QLayoutItem) UnsafeGoID() int64     { return this.vid }
func (this *QLayoutItem) UnsafeSetGoID(id int64) { this.vid = id }

func bindQLayoutItem(h uintptr) *QLayoutItem {
	if h == 0 {
		return nil
	}
	return &QLayoutItem{h: h}
}

func newownedQLayoutItem(h uintptr) *QLayoutItem {
	o := bindQLayoutItem(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QLayoutItem).Delete)
	}
	return o
}

func (this *QLayoutItem) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QLayoutItem) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQLayoutItem(h unsafe.Pointer) *QLayoutItem { return bindQLayoutItem(uintptr(h)) }
func UnsafeNewQLayoutItemOwned(h unsafe.Pointer) *QLayoutItem { return newownedQLayoutItem(uintptr(h)) }

func (this *QLayoutItem) UnsafeRelease() { this.h = 0 }

func (this *QLayoutItem) SizeHint() *QSize {
	__r, _, _ := purego.SyscallN(addr_QLayoutItem_sizeHint, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QLayoutItem) MinimumSize() *QSize {
	__r, _, _ := purego.SyscallN(addr_QLayoutItem_minimumSize, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QLayoutItem) MaximumSize() *QSize {
	__r, _, _ := purego.SyscallN(addr_QLayoutItem_maximumSize, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QLayoutItem) ExpandingDirections() Orientation {
	__r0, _, _ := purego.SyscallN(addr_QLayoutItem_expandingDirections, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return Orientation(__r)
}

func (this *QLayoutItem) SetGeometry(arg1 *QRect) {
	purego.SyscallN(addr_QLayoutItem_setGeometry, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QLayoutItem) Geometry() *QRect {
	__r, _, _ := purego.SyscallN(addr_QLayoutItem_geometry, this.h)
	runtime.KeepAlive(this)
	return newownedQRect(__r)
}

func (this *QLayoutItem) IsEmpty() bool {
	__r0, _, _ := purego.SyscallN(addr_QLayoutItem_isEmpty, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QLayoutItem) HasHeightForWidth() bool {
	__r0, _, _ := purego.SyscallN(addr_QLayoutItem_hasHeightForWidth, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QLayoutItem) HeightForWidth(arg1 int) int {
	__r0, _, _ := purego.SyscallN(addr_QLayoutItem_heightForWidth, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QLayoutItem) MinimumHeightForWidth(arg1 int) int {
	__r0, _, _ := purego.SyscallN(addr_QLayoutItem_minimumHeightForWidth, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QLayoutItem) Invalidate() {
	purego.SyscallN(addr_QLayoutItem_invalidate, this.h)
	runtime.KeepAlive(this)
}

func (this *QLayoutItem) Widget() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QLayoutItem_widget, this.h)
	runtime.KeepAlive(this)
	return bindQWidget(__r)
}

func (this *QLayoutItem) Layout() *QLayout {
	__r, _, _ := purego.SyscallN(addr_QLayoutItem_layout, this.h)
	runtime.KeepAlive(this)
	return bindQLayout(__r)
}

func (this *QLayoutItem) SpacerItem() *QSpacerItem {
	__r, _, _ := purego.SyscallN(addr_QLayoutItem_spacerItem, this.h)
	runtime.KeepAlive(this)
	return bindQSpacerItem(__r)
}

func (this *QLayoutItem) Alignment() AlignmentFlag {
	__r0, _, _ := purego.SyscallN(addr_QLayoutItem_alignment, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return AlignmentFlag(__r)
}

func (this *QLayoutItem) SetAlignment(a AlignmentFlag) {
	purego.SyscallN(addr_QLayoutItem_setAlignment, this.h, uintptr(a))
	runtime.KeepAlive(this)
}

func (this *QLayoutItem) ControlTypes() QSizePolicy__ControlType {
	__r0, _, _ := purego.SyscallN(addr_QLayoutItem_controlTypes, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QSizePolicy__ControlType(__r)
}

func (this *QLayoutItem) OperatorAssign(arg1 *QLayoutItem) *QLayoutItem {
	__r, _, _ := purego.SyscallN(addr_QLayoutItem_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQLayoutItem(__r)
}

func (this *QLayoutItem) HasHeightForWidthDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QLayoutItem_hasHeightForWidth_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QLayoutItem) HeightForWidthDefault(arg1 int) int {
	__r0, _, _ := purego.SyscallN(addr_QLayoutItem_heightForWidth_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QLayoutItem) MinimumHeightForWidthDefault(arg1 int) int {
	__r0, _, _ := purego.SyscallN(addr_QLayoutItem_minimumHeightForWidth_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QLayoutItem) InvalidateDefault() {
	purego.SyscallN(addr_QLayoutItem_invalidate_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QLayoutItem) WidgetDefault() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QLayoutItem_widget_default, this.h)
	runtime.KeepAlive(this)
	return bindQWidget(__r)
}

func (this *QLayoutItem) LayoutDefault() *QLayout {
	__r, _, _ := purego.SyscallN(addr_QLayoutItem_layout_default, this.h)
	runtime.KeepAlive(this)
	return bindQLayout(__r)
}

func (this *QLayoutItem) SpacerItemDefault() *QSpacerItem {
	__r, _, _ := purego.SyscallN(addr_QLayoutItem_spacerItem_default, this.h)
	runtime.KeepAlive(this)
	return bindQSpacerItem(__r)
}

func (this *QLayoutItem) ControlTypesDefault() QSizePolicy__ControlType {
	__r0, _, _ := purego.SyscallN(addr_QLayoutItem_controlTypes_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QSizePolicy__ControlType(__r)
}

func QLayoutItem_virtIdx(slot int64) int32 {
	switch slot {
	case 729690261457684101:
		return 0
	case -7246748916541192690:
		return 1
	case 4158099606088301136:
		return 2
	case 415665469015095569:
		return 3
	case 5684820848726969227:
		return 4
	case 1050490530495513297:
		return 5
	case -626822430251098498:
		return 6
	case -8266762397995423915:
		return 7
	case 4311534060495737339:
		return 8
	case 5037837229528203129:
		return 9
	case 6991341317990632368:
		return 10
	case -4575937485689502171:
		return 11
	case -7183390618627992737:
		return 12
	case 2272240669359776278:
		return 13
	case 3278947275714896887:
		return 14
	}
	return -1
}

func QLayoutItem_enable(h uintptr, idx int32) { pg_QLayoutItem__set_override(h, idx) }

func NewQLayoutItem() *QLayoutItem {
	id := nextVirtID()
	o := newownedQLayoutItem(pg_QLayoutItem__Go_new(id))
	o.UnsafeSetGoID(id)
	registerInstance(id, o.h, QLayoutItem_virtIdx, QLayoutItem_enable)
	return o
}

func (this *QLayoutItem) OnSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 729690261457684101, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sizeHint"
}

func (this *QLayoutItem) OnMinimumSize(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), -7246748916541192690, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "minimumSize"
}

func (this *QLayoutItem) OnMaximumSize(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 4158099606088301136, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "maximumSize"
}

func (this *QLayoutItem) OnExpandingDirections(fn func() Orientation) {
	registerVirt(this.UnsafeGoID(), 415665469015095569, func() int32 { return int32(fn()) }) // slot "expandingDirections"
}

func (this *QLayoutItem) OnSetGeometry(fn func(*QRect)) {
	registerVirt(this.UnsafeGoID(), 5684820848726969227, func(arg1 uintptr) { fn(UnsafeNewQRect(unsafe.Pointer(arg1))) }) // slot "setGeometry"
}

func (this *QLayoutItem) OnGeometry(fn func() *QRect) {
	registerVirt(this.UnsafeGoID(), 1050490530495513297, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "geometry"
}

func (this *QLayoutItem) OnIsEmpty(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -626822430251098498, func() int32 { return b2i(fn()) }) // slot "isEmpty"
}

func (this *QLayoutItem) OnHasHeightForWidth(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -8266762397995423915, func() int32 { return b2i(fn()) }) // slot "hasHeightForWidth"
}

func (this *QLayoutItem) OnHeightForWidth(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), 4311534060495737339, func(arg1 int32) int32 { return int32(fn(int(arg1))) }) // slot "heightForWidth"
}

func (this *QLayoutItem) OnMinimumHeightForWidth(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), 5037837229528203129, func(arg1 int32) int32 { return int32(fn(int(arg1))) }) // slot "minimumHeightForWidth"
}

func (this *QLayoutItem) OnInvalidate(fn func()) {
	registerVirt(this.UnsafeGoID(), 6991341317990632368, func() { fn() }) // slot "invalidate"
}

func (this *QLayoutItem) OnWidget(fn func() *QWidget) {
	registerVirt(this.UnsafeGoID(), -4575937485689502171, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "widget"
}

func (this *QLayoutItem) OnLayout(fn func() *QLayout) {
	registerVirt(this.UnsafeGoID(), -7183390618627992737, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "layout"
}

func (this *QLayoutItem) OnSpacerItem(fn func() *QSpacerItem) {
	registerVirt(this.UnsafeGoID(), 2272240669359776278, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "spacerItem"
}

func (this *QLayoutItem) OnControlTypes(fn func() QSizePolicy__ControlType) {
	registerVirt(this.UnsafeGoID(), 3278947275714896887, func() int32 { return int32(fn()) }) // slot "controlTypes"
}

func (this *QLayoutItem) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QLayoutItem_delete(this.h)
	}
	this.UnsafeRelease()
}
