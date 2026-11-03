package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QPainterPathStroker_new uintptr
	addr_QPainterPathStroker_new2 uintptr
	pg_QPainterPathStroker_setWidth func(uintptr, float64)
	pg_QPainterPathStroker_width func(uintptr) float64
	addr_QPainterPathStroker_setCapStyle uintptr
	addr_QPainterPathStroker_capStyle uintptr
	addr_QPainterPathStroker_setJoinStyle uintptr
	addr_QPainterPathStroker_joinStyle uintptr
	pg_QPainterPathStroker_setMiterLimit func(uintptr, float64)
	pg_QPainterPathStroker_miterLimit func(uintptr) float64
	pg_QPainterPathStroker_setCurveThreshold func(uintptr, float64)
	pg_QPainterPathStroker_curveThreshold func(uintptr) float64
	addr_QPainterPathStroker_setDashPattern uintptr
	addr_QPainterPathStroker_setDashPattern2 uintptr
	addr_QPainterPathStroker_dashPattern uintptr
	pg_QPainterPathStroker_setDashOffset func(uintptr, float64)
	pg_QPainterPathStroker_dashOffset func(uintptr) float64
	addr_QPainterPathStroker_createStroke uintptr
	pg_QPainterPathStroker_delete func(uintptr)
)

func init() {
	symaddr(&addr_QPainterPathStroker_new, "QPainterPathStroker_new")
	symaddr(&addr_QPainterPathStroker_new2, "QPainterPathStroker_new2")
	reg(&pg_QPainterPathStroker_setWidth, "QPainterPathStroker_setWidth")
	reg(&pg_QPainterPathStroker_width, "QPainterPathStroker_width")
	symaddr(&addr_QPainterPathStroker_setCapStyle, "QPainterPathStroker_setCapStyle")
	symaddr(&addr_QPainterPathStroker_capStyle, "QPainterPathStroker_capStyle")
	symaddr(&addr_QPainterPathStroker_setJoinStyle, "QPainterPathStroker_setJoinStyle")
	symaddr(&addr_QPainterPathStroker_joinStyle, "QPainterPathStroker_joinStyle")
	reg(&pg_QPainterPathStroker_setMiterLimit, "QPainterPathStroker_setMiterLimit")
	reg(&pg_QPainterPathStroker_miterLimit, "QPainterPathStroker_miterLimit")
	reg(&pg_QPainterPathStroker_setCurveThreshold, "QPainterPathStroker_setCurveThreshold")
	reg(&pg_QPainterPathStroker_curveThreshold, "QPainterPathStroker_curveThreshold")
	symaddr(&addr_QPainterPathStroker_setDashPattern, "QPainterPathStroker_setDashPattern")
	symaddr(&addr_QPainterPathStroker_setDashPattern2, "QPainterPathStroker_setDashPattern2")
	symaddr(&addr_QPainterPathStroker_dashPattern, "QPainterPathStroker_dashPattern")
	reg(&pg_QPainterPathStroker_setDashOffset, "QPainterPathStroker_setDashOffset")
	reg(&pg_QPainterPathStroker_dashOffset, "QPainterPathStroker_dashOffset")
	symaddr(&addr_QPainterPathStroker_createStroke, "QPainterPathStroker_createStroke")
	reg(&pg_QPainterPathStroker_delete, "QPainterPathStroker_delete")
}

type QPainterPathStroker struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QPainterPathStroker) UnsafeGoID() int64     { return this.vid }
func (this *QPainterPathStroker) UnsafeSetGoID(id int64) { this.vid = id }

func bindQPainterPathStroker(h uintptr) *QPainterPathStroker {
	if h == 0 {
		return nil
	}
	return &QPainterPathStroker{h: h}
}

func newownedQPainterPathStroker(h uintptr) *QPainterPathStroker {
	o := bindQPainterPathStroker(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QPainterPathStroker).Delete)
	}
	return o
}

func (this *QPainterPathStroker) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QPainterPathStroker) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQPainterPathStroker(h unsafe.Pointer) *QPainterPathStroker { return bindQPainterPathStroker(uintptr(h)) }
func UnsafeNewQPainterPathStrokerOwned(h unsafe.Pointer) *QPainterPathStroker { return newownedQPainterPathStroker(uintptr(h)) }

func (this *QPainterPathStroker) UnsafeRelease() { this.h = 0 }

func NewQPainterPathStroker() *QPainterPathStroker {
	__r, _, _ := purego.SyscallN(addr_QPainterPathStroker_new)
	__o := newownedQPainterPathStroker(__r)
	return __o
}

func NewQPainterPathStroker2(pen *QPen) *QPainterPathStroker {
	__r, _, _ := purego.SyscallN(addr_QPainterPathStroker_new2, uintptr(pen.UnsafePointer()))
	__o := newownedQPainterPathStroker(__r)
	runtime.KeepAlive(pen)
	return __o
}

func (this *QPainterPathStroker) SetWidth(width float64) {
	pg_QPainterPathStroker_setWidth(this.h, width)
	runtime.KeepAlive(this)
}

func (this *QPainterPathStroker) Width() float64 {
	__r := pg_QPainterPathStroker_width(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QPainterPathStroker) SetCapStyle(style PenCapStyle) {
	purego.SyscallN(addr_QPainterPathStroker_setCapStyle, this.h, uintptr(style))
	runtime.KeepAlive(this)
}

func (this *QPainterPathStroker) CapStyle() PenCapStyle {
	__r0, _, _ := purego.SyscallN(addr_QPainterPathStroker_capStyle, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return PenCapStyle(__r)
}

func (this *QPainterPathStroker) SetJoinStyle(style PenJoinStyle) {
	purego.SyscallN(addr_QPainterPathStroker_setJoinStyle, this.h, uintptr(style))
	runtime.KeepAlive(this)
}

func (this *QPainterPathStroker) JoinStyle() PenJoinStyle {
	__r0, _, _ := purego.SyscallN(addr_QPainterPathStroker_joinStyle, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return PenJoinStyle(__r)
}

func (this *QPainterPathStroker) SetMiterLimit(length float64) {
	pg_QPainterPathStroker_setMiterLimit(this.h, length)
	runtime.KeepAlive(this)
}

func (this *QPainterPathStroker) MiterLimit() float64 {
	__r := pg_QPainterPathStroker_miterLimit(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QPainterPathStroker) SetCurveThreshold(threshold float64) {
	pg_QPainterPathStroker_setCurveThreshold(this.h, threshold)
	runtime.KeepAlive(this)
}

func (this *QPainterPathStroker) CurveThreshold() float64 {
	__r := pg_QPainterPathStroker_curveThreshold(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QPainterPathStroker) SetDashPattern(arg1 PenStyle) {
	purego.SyscallN(addr_QPainterPathStroker_setDashPattern, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QPainterPathStroker) SetDashPatternWithDashPattern(dashPattern []float64) {
	dashPattern_arr := make([]float64, len(dashPattern))
	for __i, __e := range dashPattern {
		dashPattern_arr[__i] = __e
	}
	var dashPattern_p uintptr
	if len(dashPattern_arr) > 0 {
		dashPattern_p = uintptr(unsafe.Pointer(&dashPattern_arr[0]))
	}
	purego.SyscallN(addr_QPainterPathStroker_setDashPattern2, this.h, dashPattern_p, uintptr(len(dashPattern)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(dashPattern_arr)
}

func (this *QPainterPathStroker) DashPattern() []float64 {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QPainterPathStroker_dashPattern, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]float64, __n)
	__sl := unsafe.Slice((*float64)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = __sl[__i]
	}
	pgFree(__r)
	return __out
}

func (this *QPainterPathStroker) SetDashOffset(offset float64) {
	pg_QPainterPathStroker_setDashOffset(this.h, offset)
	runtime.KeepAlive(this)
}

func (this *QPainterPathStroker) DashOffset() float64 {
	__r := pg_QPainterPathStroker_dashOffset(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QPainterPathStroker) CreateStroke(path *QPainterPath) *QPainterPath {
	__r, _, _ := purego.SyscallN(addr_QPainterPathStroker_createStroke, this.h, uintptr(path.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(path)
	return newownedQPainterPath(__r)
}

func (this *QPainterPathStroker) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QPainterPathStroker_delete(this.h)
	}
	this.UnsafeRelease()
}
