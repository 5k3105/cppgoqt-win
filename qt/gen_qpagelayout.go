package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QPageLayout_new uintptr
	addr_QPageLayout_new2 uintptr
	addr_QPageLayout_new3 uintptr
	addr_QPageLayout_new4 uintptr
	addr_QPageLayout_new5 uintptr
	addr_QPageLayout_OperatorAssign uintptr
	addr_QPageLayout_swap uintptr
	addr_QPageLayout_isEquivalentTo uintptr
	addr_QPageLayout_isValid uintptr
	addr_QPageLayout_setMode uintptr
	addr_QPageLayout_mode uintptr
	addr_QPageLayout_setPageSize uintptr
	addr_QPageLayout_pageSize uintptr
	addr_QPageLayout_setOrientation uintptr
	addr_QPageLayout_orientation uintptr
	addr_QPageLayout_setUnits uintptr
	addr_QPageLayout_units uintptr
	addr_QPageLayout_setMargins uintptr
	pg_QPageLayout_setLeftMargin func(uintptr, float64) int32
	pg_QPageLayout_setRightMargin func(uintptr, float64) int32
	pg_QPageLayout_setTopMargin func(uintptr, float64) int32
	pg_QPageLayout_setBottomMargin func(uintptr, float64) int32
	addr_QPageLayout_margins uintptr
	addr_QPageLayout_margins2 uintptr
	addr_QPageLayout_marginsPoints uintptr
	addr_QPageLayout_marginsPixels uintptr
	addr_QPageLayout_setMinimumMargins uintptr
	addr_QPageLayout_minimumMargins uintptr
	addr_QPageLayout_maximumMargins uintptr
	addr_QPageLayout_fullRect uintptr
	addr_QPageLayout_fullRect2 uintptr
	addr_QPageLayout_fullRectPoints uintptr
	addr_QPageLayout_fullRectPixels uintptr
	addr_QPageLayout_paintRect uintptr
	addr_QPageLayout_paintRect2 uintptr
	addr_QPageLayout_paintRectPoints uintptr
	addr_QPageLayout_paintRectPixels uintptr
	addr_QPageLayout_setPageSize2 uintptr
	addr_QPageLayout_setMargins2 uintptr
	pg_QPageLayout_setLeftMargin2 func(uintptr, float64, int32) int32
	pg_QPageLayout_setRightMargin2 func(uintptr, float64, int32) int32
	pg_QPageLayout_setTopMargin2 func(uintptr, float64, int32) int32
	pg_QPageLayout_setBottomMargin2 func(uintptr, float64, int32) int32
	addr_QPageLayout_OperatorEqual uintptr
	addr_QPageLayout_OperatorNotEqual uintptr
	pg_QPageLayout_delete func(uintptr)
)

func init() {
	symaddr(&addr_QPageLayout_new, "QPageLayout_new")
	symaddr(&addr_QPageLayout_new2, "QPageLayout_new2")
	symaddr(&addr_QPageLayout_new3, "QPageLayout_new3")
	symaddr(&addr_QPageLayout_new4, "QPageLayout_new4")
	symaddr(&addr_QPageLayout_new5, "QPageLayout_new5")
	symaddr(&addr_QPageLayout_OperatorAssign, "QPageLayout_OperatorAssign")
	symaddr(&addr_QPageLayout_swap, "QPageLayout_swap")
	symaddr(&addr_QPageLayout_isEquivalentTo, "QPageLayout_isEquivalentTo")
	symaddr(&addr_QPageLayout_isValid, "QPageLayout_isValid")
	symaddr(&addr_QPageLayout_setMode, "QPageLayout_setMode")
	symaddr(&addr_QPageLayout_mode, "QPageLayout_mode")
	symaddr(&addr_QPageLayout_setPageSize, "QPageLayout_setPageSize")
	symaddr(&addr_QPageLayout_pageSize, "QPageLayout_pageSize")
	symaddr(&addr_QPageLayout_setOrientation, "QPageLayout_setOrientation")
	symaddr(&addr_QPageLayout_orientation, "QPageLayout_orientation")
	symaddr(&addr_QPageLayout_setUnits, "QPageLayout_setUnits")
	symaddr(&addr_QPageLayout_units, "QPageLayout_units")
	symaddr(&addr_QPageLayout_setMargins, "QPageLayout_setMargins")
	reg(&pg_QPageLayout_setLeftMargin, "QPageLayout_setLeftMargin")
	reg(&pg_QPageLayout_setRightMargin, "QPageLayout_setRightMargin")
	reg(&pg_QPageLayout_setTopMargin, "QPageLayout_setTopMargin")
	reg(&pg_QPageLayout_setBottomMargin, "QPageLayout_setBottomMargin")
	symaddr(&addr_QPageLayout_margins, "QPageLayout_margins")
	symaddr(&addr_QPageLayout_margins2, "QPageLayout_margins2")
	symaddr(&addr_QPageLayout_marginsPoints, "QPageLayout_marginsPoints")
	symaddr(&addr_QPageLayout_marginsPixels, "QPageLayout_marginsPixels")
	symaddr(&addr_QPageLayout_setMinimumMargins, "QPageLayout_setMinimumMargins")
	symaddr(&addr_QPageLayout_minimumMargins, "QPageLayout_minimumMargins")
	symaddr(&addr_QPageLayout_maximumMargins, "QPageLayout_maximumMargins")
	symaddr(&addr_QPageLayout_fullRect, "QPageLayout_fullRect")
	symaddr(&addr_QPageLayout_fullRect2, "QPageLayout_fullRect2")
	symaddr(&addr_QPageLayout_fullRectPoints, "QPageLayout_fullRectPoints")
	symaddr(&addr_QPageLayout_fullRectPixels, "QPageLayout_fullRectPixels")
	symaddr(&addr_QPageLayout_paintRect, "QPageLayout_paintRect")
	symaddr(&addr_QPageLayout_paintRect2, "QPageLayout_paintRect2")
	symaddr(&addr_QPageLayout_paintRectPoints, "QPageLayout_paintRectPoints")
	symaddr(&addr_QPageLayout_paintRectPixels, "QPageLayout_paintRectPixels")
	symaddr(&addr_QPageLayout_setPageSize2, "QPageLayout_setPageSize2")
	symaddr(&addr_QPageLayout_setMargins2, "QPageLayout_setMargins2")
	reg(&pg_QPageLayout_setLeftMargin2, "QPageLayout_setLeftMargin2")
	reg(&pg_QPageLayout_setRightMargin2, "QPageLayout_setRightMargin2")
	reg(&pg_QPageLayout_setTopMargin2, "QPageLayout_setTopMargin2")
	reg(&pg_QPageLayout_setBottomMargin2, "QPageLayout_setBottomMargin2")
	symaddr(&addr_QPageLayout_OperatorEqual, "QPageLayout_OperatorEqual")
	symaddr(&addr_QPageLayout_OperatorNotEqual, "QPageLayout_OperatorNotEqual")
	reg(&pg_QPageLayout_delete, "QPageLayout_delete")
}

type QPageLayout struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QPageLayout) UnsafeGoID() int64     { return this.vid }
func (this *QPageLayout) UnsafeSetGoID(id int64) { this.vid = id }

func bindQPageLayout(h uintptr) *QPageLayout {
	if h == 0 {
		return nil
	}
	return &QPageLayout{h: h}
}

func newownedQPageLayout(h uintptr) *QPageLayout {
	o := bindQPageLayout(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QPageLayout).Delete)
	}
	return o
}

func (this *QPageLayout) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QPageLayout) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQPageLayout(h unsafe.Pointer) *QPageLayout { return bindQPageLayout(uintptr(h)) }
func UnsafeNewQPageLayoutOwned(h unsafe.Pointer) *QPageLayout { return newownedQPageLayout(uintptr(h)) }

func (this *QPageLayout) UnsafeRelease() { this.h = 0 }

func NewQPageLayout() *QPageLayout {
	__r, _, _ := purego.SyscallN(addr_QPageLayout_new)
	__o := newownedQPageLayout(__r)
	return __o
}

func NewQPageLayout2(pageSize *QPageSize, orientation QPageLayout__Orientation, margins *QMarginsF) *QPageLayout {
	__r, _, _ := purego.SyscallN(addr_QPageLayout_new2, uintptr(pageSize.UnsafePointer()), uintptr(orientation), uintptr(margins.UnsafePointer()))
	__o := newownedQPageLayout(__r)
	runtime.KeepAlive(pageSize)
	runtime.KeepAlive(margins)
	return __o
}

func NewQPageLayout3(other *QPageLayout) *QPageLayout {
	__r, _, _ := purego.SyscallN(addr_QPageLayout_new3, uintptr(other.UnsafePointer()))
	__o := newownedQPageLayout(__r)
	runtime.KeepAlive(other)
	return __o
}

func NewQPageLayout4(pageSize *QPageSize, orientation QPageLayout__Orientation, margins *QMarginsF, units QPageLayout__Unit) *QPageLayout {
	__r, _, _ := purego.SyscallN(addr_QPageLayout_new4, uintptr(pageSize.UnsafePointer()), uintptr(orientation), uintptr(margins.UnsafePointer()), uintptr(units))
	__o := newownedQPageLayout(__r)
	runtime.KeepAlive(pageSize)
	runtime.KeepAlive(margins)
	return __o
}

func NewQPageLayout5(pageSize *QPageSize, orientation QPageLayout__Orientation, margins *QMarginsF, units QPageLayout__Unit, minMargins *QMarginsF) *QPageLayout {
	__r, _, _ := purego.SyscallN(addr_QPageLayout_new5, uintptr(pageSize.UnsafePointer()), uintptr(orientation), uintptr(margins.UnsafePointer()), uintptr(units), uintptr(minMargins.UnsafePointer()))
	__o := newownedQPageLayout(__r)
	runtime.KeepAlive(pageSize)
	runtime.KeepAlive(margins)
	runtime.KeepAlive(minMargins)
	return __o
}

func (this *QPageLayout) OperatorAssign(other *QPageLayout) *QPageLayout {
	__r, _, _ := purego.SyscallN(addr_QPageLayout_OperatorAssign, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return bindQPageLayout(__r)
}

func (this *QPageLayout) Swap(other *QPageLayout) {
	purego.SyscallN(addr_QPageLayout_swap, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
}

func (this *QPageLayout) IsEquivalentTo(other *QPageLayout) bool {
	__r0, _, _ := purego.SyscallN(addr_QPageLayout_isEquivalentTo, this.h, uintptr(other.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return __r != 0
}

func (this *QPageLayout) IsValid() bool {
	__r0, _, _ := purego.SyscallN(addr_QPageLayout_isValid, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QPageLayout) SetMode(mode QPageLayout__Mode) {
	purego.SyscallN(addr_QPageLayout_setMode, this.h, uintptr(mode))
	runtime.KeepAlive(this)
}

func (this *QPageLayout) Mode() QPageLayout__Mode {
	__r0, _, _ := purego.SyscallN(addr_QPageLayout_mode, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QPageLayout__Mode(__r)
}

func (this *QPageLayout) SetPageSize(pageSize *QPageSize) {
	purego.SyscallN(addr_QPageLayout_setPageSize, this.h, uintptr(pageSize.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pageSize)
}

func (this *QPageLayout) PageSize() *QPageSize {
	__r, _, _ := purego.SyscallN(addr_QPageLayout_pageSize, this.h)
	runtime.KeepAlive(this)
	return newownedQPageSize(__r)
}

func (this *QPageLayout) SetOrientation(orientation QPageLayout__Orientation) {
	purego.SyscallN(addr_QPageLayout_setOrientation, this.h, uintptr(orientation))
	runtime.KeepAlive(this)
}

func (this *QPageLayout) Orientation() QPageLayout__Orientation {
	__r0, _, _ := purego.SyscallN(addr_QPageLayout_orientation, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QPageLayout__Orientation(__r)
}

func (this *QPageLayout) SetUnits(units QPageLayout__Unit) {
	purego.SyscallN(addr_QPageLayout_setUnits, this.h, uintptr(units))
	runtime.KeepAlive(this)
}

func (this *QPageLayout) Units() QPageLayout__Unit {
	__r0, _, _ := purego.SyscallN(addr_QPageLayout_units, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QPageLayout__Unit(__r)
}

func (this *QPageLayout) SetMargins(margins *QMarginsF) bool {
	__r0, _, _ := purego.SyscallN(addr_QPageLayout_setMargins, this.h, uintptr(margins.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(margins)
	return __r != 0
}

func (this *QPageLayout) SetLeftMargin(leftMargin float64) bool {
	__r := pg_QPageLayout_setLeftMargin(this.h, leftMargin)
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QPageLayout) SetRightMargin(rightMargin float64) bool {
	__r := pg_QPageLayout_setRightMargin(this.h, rightMargin)
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QPageLayout) SetTopMargin(topMargin float64) bool {
	__r := pg_QPageLayout_setTopMargin(this.h, topMargin)
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QPageLayout) SetBottomMargin(bottomMargin float64) bool {
	__r := pg_QPageLayout_setBottomMargin(this.h, bottomMargin)
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QPageLayout) Margins() *QMarginsF {
	__r, _, _ := purego.SyscallN(addr_QPageLayout_margins, this.h)
	runtime.KeepAlive(this)
	return newownedQMarginsF(__r)
}

func (this *QPageLayout) MarginsWithUnits(units QPageLayout__Unit) *QMarginsF {
	__r, _, _ := purego.SyscallN(addr_QPageLayout_margins2, this.h, uintptr(units))
	runtime.KeepAlive(this)
	return newownedQMarginsF(__r)
}

func (this *QPageLayout) MarginsPoints() *QMargins {
	__r, _, _ := purego.SyscallN(addr_QPageLayout_marginsPoints, this.h)
	runtime.KeepAlive(this)
	return newownedQMargins(__r)
}

func (this *QPageLayout) MarginsPixels(resolution int) *QMargins {
	__r, _, _ := purego.SyscallN(addr_QPageLayout_marginsPixels, this.h, uintptr(resolution))
	runtime.KeepAlive(this)
	return newownedQMargins(__r)
}

func (this *QPageLayout) SetMinimumMargins(minMargins *QMarginsF) {
	purego.SyscallN(addr_QPageLayout_setMinimumMargins, this.h, uintptr(minMargins.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(minMargins)
}

func (this *QPageLayout) MinimumMargins() *QMarginsF {
	__r, _, _ := purego.SyscallN(addr_QPageLayout_minimumMargins, this.h)
	runtime.KeepAlive(this)
	return newownedQMarginsF(__r)
}

func (this *QPageLayout) MaximumMargins() *QMarginsF {
	__r, _, _ := purego.SyscallN(addr_QPageLayout_maximumMargins, this.h)
	runtime.KeepAlive(this)
	return newownedQMarginsF(__r)
}

func (this *QPageLayout) FullRect() *QRectF {
	__r, _, _ := purego.SyscallN(addr_QPageLayout_fullRect, this.h)
	runtime.KeepAlive(this)
	return newownedQRectF(__r)
}

func (this *QPageLayout) FullRectWithUnits(units QPageLayout__Unit) *QRectF {
	__r, _, _ := purego.SyscallN(addr_QPageLayout_fullRect2, this.h, uintptr(units))
	runtime.KeepAlive(this)
	return newownedQRectF(__r)
}

func (this *QPageLayout) FullRectPoints() *QRect {
	__r, _, _ := purego.SyscallN(addr_QPageLayout_fullRectPoints, this.h)
	runtime.KeepAlive(this)
	return newownedQRect(__r)
}

func (this *QPageLayout) FullRectPixels(resolution int) *QRect {
	__r, _, _ := purego.SyscallN(addr_QPageLayout_fullRectPixels, this.h, uintptr(resolution))
	runtime.KeepAlive(this)
	return newownedQRect(__r)
}

func (this *QPageLayout) PaintRect() *QRectF {
	__r, _, _ := purego.SyscallN(addr_QPageLayout_paintRect, this.h)
	runtime.KeepAlive(this)
	return newownedQRectF(__r)
}

func (this *QPageLayout) PaintRectWithUnits(units QPageLayout__Unit) *QRectF {
	__r, _, _ := purego.SyscallN(addr_QPageLayout_paintRect2, this.h, uintptr(units))
	runtime.KeepAlive(this)
	return newownedQRectF(__r)
}

func (this *QPageLayout) PaintRectPoints() *QRect {
	__r, _, _ := purego.SyscallN(addr_QPageLayout_paintRectPoints, this.h)
	runtime.KeepAlive(this)
	return newownedQRect(__r)
}

func (this *QPageLayout) PaintRectPixels(resolution int) *QRect {
	__r, _, _ := purego.SyscallN(addr_QPageLayout_paintRectPixels, this.h, uintptr(resolution))
	runtime.KeepAlive(this)
	return newownedQRect(__r)
}

func (this *QPageLayout) SetPageSize2(pageSize *QPageSize, minMargins *QMarginsF) {
	purego.SyscallN(addr_QPageLayout_setPageSize2, this.h, uintptr(pageSize.UnsafePointer()), uintptr(minMargins.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pageSize)
	runtime.KeepAlive(minMargins)
}

func (this *QPageLayout) SetMargins2(margins *QMarginsF, outOfBoundsPolicy QPageLayout__OutOfBoundsPolicy) bool {
	__r0, _, _ := purego.SyscallN(addr_QPageLayout_setMargins2, this.h, uintptr(margins.UnsafePointer()), uintptr(outOfBoundsPolicy))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(margins)
	return __r != 0
}

func (this *QPageLayout) SetLeftMargin2(leftMargin float64, outOfBoundsPolicy QPageLayout__OutOfBoundsPolicy) bool {
	__r := pg_QPageLayout_setLeftMargin2(this.h, leftMargin, int32(outOfBoundsPolicy))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QPageLayout) SetRightMargin2(rightMargin float64, outOfBoundsPolicy QPageLayout__OutOfBoundsPolicy) bool {
	__r := pg_QPageLayout_setRightMargin2(this.h, rightMargin, int32(outOfBoundsPolicy))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QPageLayout) SetTopMargin2(topMargin float64, outOfBoundsPolicy QPageLayout__OutOfBoundsPolicy) bool {
	__r := pg_QPageLayout_setTopMargin2(this.h, topMargin, int32(outOfBoundsPolicy))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QPageLayout) SetBottomMargin2(bottomMargin float64, outOfBoundsPolicy QPageLayout__OutOfBoundsPolicy) bool {
	__r := pg_QPageLayout_setBottomMargin2(this.h, bottomMargin, int32(outOfBoundsPolicy))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QPageLayout) OperatorEqual(rhs *QPageLayout) bool {
	__r0, _, _ := purego.SyscallN(addr_QPageLayout_OperatorEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QPageLayout) OperatorNotEqual(rhs *QPageLayout) bool {
	__r0, _, _ := purego.SyscallN(addr_QPageLayout_OperatorNotEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QPageLayout) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QPageLayout_delete(this.h)
	}
	this.UnsafeRelease()
}
