package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QPen_new uintptr
	addr_QPen_new2 uintptr
	addr_QPen_new3 uintptr
	pg_QPen_new4 func(uintptr, float64) uintptr
	addr_QPen_new5 uintptr
	pg_QPen_new6 func(uintptr, float64, int32) uintptr
	pg_QPen_new7 func(uintptr, float64, int32, int32) uintptr
	pg_QPen_new8 func(uintptr, float64, int32, int32, int32) uintptr
	addr_QPen_OperatorAssign uintptr
	addr_QPen_swap uintptr
	addr_QPen_OperatorAssign2 uintptr
	addr_QPen_OperatorAssign3 uintptr
	addr_QPen_style uintptr
	addr_QPen_setStyle uintptr
	addr_QPen_dashPattern uintptr
	addr_QPen_setDashPattern uintptr
	pg_QPen_dashOffset func(uintptr) float64
	pg_QPen_setDashOffset func(uintptr, float64)
	pg_QPen_miterLimit func(uintptr) float64
	pg_QPen_setMiterLimit func(uintptr, float64)
	pg_QPen_widthF func(uintptr) float64
	pg_QPen_setWidthF func(uintptr, float64)
	addr_QPen_width uintptr
	addr_QPen_setWidth uintptr
	addr_QPen_color uintptr
	addr_QPen_setColor uintptr
	addr_QPen_brush uintptr
	addr_QPen_setBrush uintptr
	addr_QPen_isSolid uintptr
	addr_QPen_capStyle uintptr
	addr_QPen_setCapStyle uintptr
	addr_QPen_joinStyle uintptr
	addr_QPen_setJoinStyle uintptr
	addr_QPen_isCosmetic uintptr
	addr_QPen_setCosmetic uintptr
	addr_QPen_OperatorEqual uintptr
	addr_QPen_OperatorNotEqual uintptr
	addr_QPen_ToQVariant uintptr
	addr_QPen_isDetached uintptr
	addr_QPen_OperatorEqual2 uintptr
	addr_QPen_OperatorNotEqual2 uintptr
	addr_QPen_toQVariant uintptr
	pg_QPen_delete func(uintptr)
)

func init() {
	symaddr(&addr_QPen_new, "QPen_new")
	symaddr(&addr_QPen_new2, "QPen_new2")
	symaddr(&addr_QPen_new3, "QPen_new3")
	reg(&pg_QPen_new4, "QPen_new4")
	symaddr(&addr_QPen_new5, "QPen_new5")
	reg(&pg_QPen_new6, "QPen_new6")
	reg(&pg_QPen_new7, "QPen_new7")
	reg(&pg_QPen_new8, "QPen_new8")
	symaddr(&addr_QPen_OperatorAssign, "QPen_OperatorAssign")
	symaddr(&addr_QPen_swap, "QPen_swap")
	symaddr(&addr_QPen_OperatorAssign2, "QPen_OperatorAssign2")
	symaddr(&addr_QPen_OperatorAssign3, "QPen_OperatorAssign3")
	symaddr(&addr_QPen_style, "QPen_style")
	symaddr(&addr_QPen_setStyle, "QPen_setStyle")
	symaddr(&addr_QPen_dashPattern, "QPen_dashPattern")
	symaddr(&addr_QPen_setDashPattern, "QPen_setDashPattern")
	reg(&pg_QPen_dashOffset, "QPen_dashOffset")
	reg(&pg_QPen_setDashOffset, "QPen_setDashOffset")
	reg(&pg_QPen_miterLimit, "QPen_miterLimit")
	reg(&pg_QPen_setMiterLimit, "QPen_setMiterLimit")
	reg(&pg_QPen_widthF, "QPen_widthF")
	reg(&pg_QPen_setWidthF, "QPen_setWidthF")
	symaddr(&addr_QPen_width, "QPen_width")
	symaddr(&addr_QPen_setWidth, "QPen_setWidth")
	symaddr(&addr_QPen_color, "QPen_color")
	symaddr(&addr_QPen_setColor, "QPen_setColor")
	symaddr(&addr_QPen_brush, "QPen_brush")
	symaddr(&addr_QPen_setBrush, "QPen_setBrush")
	symaddr(&addr_QPen_isSolid, "QPen_isSolid")
	symaddr(&addr_QPen_capStyle, "QPen_capStyle")
	symaddr(&addr_QPen_setCapStyle, "QPen_setCapStyle")
	symaddr(&addr_QPen_joinStyle, "QPen_joinStyle")
	symaddr(&addr_QPen_setJoinStyle, "QPen_setJoinStyle")
	symaddr(&addr_QPen_isCosmetic, "QPen_isCosmetic")
	symaddr(&addr_QPen_setCosmetic, "QPen_setCosmetic")
	symaddr(&addr_QPen_OperatorEqual, "QPen_OperatorEqual")
	symaddr(&addr_QPen_OperatorNotEqual, "QPen_OperatorNotEqual")
	symaddr(&addr_QPen_ToQVariant, "QPen_ToQVariant")
	symaddr(&addr_QPen_isDetached, "QPen_isDetached")
	symaddr(&addr_QPen_OperatorEqual2, "QPen_OperatorEqual2")
	symaddr(&addr_QPen_OperatorNotEqual2, "QPen_OperatorNotEqual2")
	symaddr(&addr_QPen_toQVariant, "QPen_toQVariant")
	reg(&pg_QPen_delete, "QPen_delete")
}

type QPen struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QPen) UnsafeGoID() int64     { return this.vid }
func (this *QPen) UnsafeSetGoID(id int64) { this.vid = id }

func bindQPen(h uintptr) *QPen {
	if h == 0 {
		return nil
	}
	return &QPen{h: h}
}

func newownedQPen(h uintptr) *QPen {
	o := bindQPen(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QPen).Delete)
	}
	return o
}

func (this *QPen) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QPen) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQPen(h unsafe.Pointer) *QPen { return bindQPen(uintptr(h)) }
func UnsafeNewQPenOwned(h unsafe.Pointer) *QPen { return newownedQPen(uintptr(h)) }

func (this *QPen) UnsafeRelease() { this.h = 0 }

func NewQPen() *QPen {
	__r, _, _ := purego.SyscallN(addr_QPen_new)
	__o := newownedQPen(__r)
	return __o
}

func NewQPen2(arg1 PenStyle) *QPen {
	__r, _, _ := purego.SyscallN(addr_QPen_new2, uintptr(arg1))
	__o := newownedQPen(__r)
	return __o
}

func NewQPen3(color *QColor) *QPen {
	__r, _, _ := purego.SyscallN(addr_QPen_new3, uintptr(color.UnsafePointer()))
	__o := newownedQPen(__r)
	runtime.KeepAlive(color)
	return __o
}

func NewQPen4(brush *QBrush, width float64) *QPen {
	__o := newownedQPen(pg_QPen_new4(uintptr(brush.UnsafePointer()), width))
	runtime.KeepAlive(brush)
	return __o
}

func NewQPen5(pen *QPen) *QPen {
	__r, _, _ := purego.SyscallN(addr_QPen_new5, uintptr(pen.UnsafePointer()))
	__o := newownedQPen(__r)
	runtime.KeepAlive(pen)
	return __o
}

func NewQPen6(brush *QBrush, width float64, s PenStyle) *QPen {
	__o := newownedQPen(pg_QPen_new6(uintptr(brush.UnsafePointer()), width, int32(s)))
	runtime.KeepAlive(brush)
	return __o
}

func NewQPen7(brush *QBrush, width float64, s PenStyle, c PenCapStyle) *QPen {
	__o := newownedQPen(pg_QPen_new7(uintptr(brush.UnsafePointer()), width, int32(s), int32(c)))
	runtime.KeepAlive(brush)
	return __o
}

func NewQPen8(brush *QBrush, width float64, s PenStyle, c PenCapStyle, j PenJoinStyle) *QPen {
	__o := newownedQPen(pg_QPen_new8(uintptr(brush.UnsafePointer()), width, int32(s), int32(c), int32(j)))
	runtime.KeepAlive(brush)
	return __o
}

func (this *QPen) OperatorAssign(pen *QPen) *QPen {
	__r, _, _ := purego.SyscallN(addr_QPen_OperatorAssign, this.h, uintptr(pen.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pen)
	return bindQPen(__r)
}

func (this *QPen) Swap(other *QPen) {
	purego.SyscallN(addr_QPen_swap, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
}

func (this *QPen) OperatorAssignWithColor(color *QColor) *QPen {
	__r, _, _ := purego.SyscallN(addr_QPen_OperatorAssign2, this.h, uintptr(color.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(color)
	return bindQPen(__r)
}

func (this *QPen) OperatorAssignWithStyle(style PenStyle) *QPen {
	__r, _, _ := purego.SyscallN(addr_QPen_OperatorAssign3, this.h, uintptr(style))
	runtime.KeepAlive(this)
	return bindQPen(__r)
}

func (this *QPen) Style() PenStyle {
	__r0, _, _ := purego.SyscallN(addr_QPen_style, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return PenStyle(__r)
}

func (this *QPen) SetStyle(arg1 PenStyle) {
	purego.SyscallN(addr_QPen_setStyle, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QPen) DashPattern() []float64 {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QPen_dashPattern, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]float64, __n)
	__sl := unsafe.Slice((*float64)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = __sl[__i]
	}
	pgFree(__r)
	return __out
}

func (this *QPen) SetDashPattern(pattern []float64) {
	pattern_arr := make([]float64, len(pattern))
	for __i, __e := range pattern {
		pattern_arr[__i] = __e
	}
	var pattern_p uintptr
	if len(pattern_arr) > 0 {
		pattern_p = uintptr(unsafe.Pointer(&pattern_arr[0]))
	}
	purego.SyscallN(addr_QPen_setDashPattern, this.h, pattern_p, uintptr(len(pattern)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pattern_arr)
}

func (this *QPen) DashOffset() float64 {
	__r := pg_QPen_dashOffset(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QPen) SetDashOffset(doffset float64) {
	pg_QPen_setDashOffset(this.h, doffset)
	runtime.KeepAlive(this)
}

func (this *QPen) MiterLimit() float64 {
	__r := pg_QPen_miterLimit(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QPen) SetMiterLimit(limit float64) {
	pg_QPen_setMiterLimit(this.h, limit)
	runtime.KeepAlive(this)
}

func (this *QPen) WidthF() float64 {
	__r := pg_QPen_widthF(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QPen) SetWidthF(width float64) {
	pg_QPen_setWidthF(this.h, width)
	runtime.KeepAlive(this)
}

func (this *QPen) Width() int {
	__r0, _, _ := purego.SyscallN(addr_QPen_width, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QPen) SetWidth(width int) {
	purego.SyscallN(addr_QPen_setWidth, this.h, uintptr(width))
	runtime.KeepAlive(this)
}

func (this *QPen) Color() *QColor {
	__r, _, _ := purego.SyscallN(addr_QPen_color, this.h)
	runtime.KeepAlive(this)
	return newownedQColor(__r)
}

func (this *QPen) SetColor(color *QColor) {
	purego.SyscallN(addr_QPen_setColor, this.h, uintptr(color.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(color)
}

func (this *QPen) Brush() *QBrush {
	__r, _, _ := purego.SyscallN(addr_QPen_brush, this.h)
	runtime.KeepAlive(this)
	return newownedQBrush(__r)
}

func (this *QPen) SetBrush(brush *QBrush) {
	purego.SyscallN(addr_QPen_setBrush, this.h, uintptr(brush.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(brush)
}

func (this *QPen) IsSolid() bool {
	__r0, _, _ := purego.SyscallN(addr_QPen_isSolid, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QPen) CapStyle() PenCapStyle {
	__r0, _, _ := purego.SyscallN(addr_QPen_capStyle, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return PenCapStyle(__r)
}

func (this *QPen) SetCapStyle(pcs PenCapStyle) {
	purego.SyscallN(addr_QPen_setCapStyle, this.h, uintptr(pcs))
	runtime.KeepAlive(this)
}

func (this *QPen) JoinStyle() PenJoinStyle {
	__r0, _, _ := purego.SyscallN(addr_QPen_joinStyle, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return PenJoinStyle(__r)
}

func (this *QPen) SetJoinStyle(pcs PenJoinStyle) {
	purego.SyscallN(addr_QPen_setJoinStyle, this.h, uintptr(pcs))
	runtime.KeepAlive(this)
}

func (this *QPen) IsCosmetic() bool {
	__r0, _, _ := purego.SyscallN(addr_QPen_isCosmetic, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QPen) SetCosmetic(cosmetic bool) {
	purego.SyscallN(addr_QPen_setCosmetic, this.h, uintptr(b2i(cosmetic)))
	runtime.KeepAlive(this)
}

func (this *QPen) OperatorEqual(p *QPen) bool {
	__r0, _, _ := purego.SyscallN(addr_QPen_OperatorEqual, this.h, uintptr(p.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
	return __r != 0
}

func (this *QPen) OperatorNotEqual(p *QPen) bool {
	__r0, _, _ := purego.SyscallN(addr_QPen_OperatorNotEqual, this.h, uintptr(p.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
	return __r != 0
}

func (this *QPen) ToQVariant() *QVariant {
	__r, _, _ := purego.SyscallN(addr_QPen_ToQVariant, this.h)
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QPen) IsDetached() bool {
	__r0, _, _ := purego.SyscallN(addr_QPen_isDetached, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QPen) OperatorEqualWithRhs(rhs *QColor) bool {
	__r0, _, _ := purego.SyscallN(addr_QPen_OperatorEqual2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QPen) OperatorNotEqualWithRhs(rhs *QColor) bool {
	__r0, _, _ := purego.SyscallN(addr_QPen_OperatorNotEqual2, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QPen) ToQVariant2() *QVariant {
	__r, _, _ := purego.SyscallN(addr_QPen_toQVariant, this.h)
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QPen) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QPen_delete(this.h)
	}
	this.UnsafeRelease()
}
