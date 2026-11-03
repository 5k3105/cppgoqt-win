package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QEasingCurve_new uintptr
	addr_QEasingCurve_new2 uintptr
	addr_QEasingCurve_new3 uintptr
	addr_QEasingCurve_OperatorAssign uintptr
	addr_QEasingCurve_swap uintptr
	pg_QEasingCurve_amplitude func(uintptr) float64
	pg_QEasingCurve_setAmplitude func(uintptr, float64)
	pg_QEasingCurve_period func(uintptr) float64
	pg_QEasingCurve_setPeriod func(uintptr, float64)
	pg_QEasingCurve_overshoot func(uintptr) float64
	pg_QEasingCurve_setOvershoot func(uintptr, float64)
	addr_QEasingCurve_addCubicBezierSegment uintptr
	pg_QEasingCurve_addTCBSegment func(uintptr, uintptr, float64, float64, float64)
	addr_QEasingCurve_toCubicSpline uintptr
	addr_QEasingCurve_type uintptr
	addr_QEasingCurve_setType uintptr
	pg_QEasingCurve_valueForProgress func(uintptr, float64) float64
	addr_QEasingCurve_OperatorEqual uintptr
	addr_QEasingCurve_OperatorNotEqual uintptr
	pg_QEasingCurve_delete func(uintptr)
)

func init() {
	symaddr(&addr_QEasingCurve_new, "QEasingCurve_new")
	symaddr(&addr_QEasingCurve_new2, "QEasingCurve_new2")
	symaddr(&addr_QEasingCurve_new3, "QEasingCurve_new3")
	symaddr(&addr_QEasingCurve_OperatorAssign, "QEasingCurve_OperatorAssign")
	symaddr(&addr_QEasingCurve_swap, "QEasingCurve_swap")
	reg(&pg_QEasingCurve_amplitude, "QEasingCurve_amplitude")
	reg(&pg_QEasingCurve_setAmplitude, "QEasingCurve_setAmplitude")
	reg(&pg_QEasingCurve_period, "QEasingCurve_period")
	reg(&pg_QEasingCurve_setPeriod, "QEasingCurve_setPeriod")
	reg(&pg_QEasingCurve_overshoot, "QEasingCurve_overshoot")
	reg(&pg_QEasingCurve_setOvershoot, "QEasingCurve_setOvershoot")
	symaddr(&addr_QEasingCurve_addCubicBezierSegment, "QEasingCurve_addCubicBezierSegment")
	reg(&pg_QEasingCurve_addTCBSegment, "QEasingCurve_addTCBSegment")
	symaddr(&addr_QEasingCurve_toCubicSpline, "QEasingCurve_toCubicSpline")
	symaddr(&addr_QEasingCurve_type, "QEasingCurve_type")
	symaddr(&addr_QEasingCurve_setType, "QEasingCurve_setType")
	reg(&pg_QEasingCurve_valueForProgress, "QEasingCurve_valueForProgress")
	symaddr(&addr_QEasingCurve_OperatorEqual, "QEasingCurve_OperatorEqual")
	symaddr(&addr_QEasingCurve_OperatorNotEqual, "QEasingCurve_OperatorNotEqual")
	reg(&pg_QEasingCurve_delete, "QEasingCurve_delete")
}

type QEasingCurve struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QEasingCurve) UnsafeGoID() int64     { return this.vid }
func (this *QEasingCurve) UnsafeSetGoID(id int64) { this.vid = id }

func bindQEasingCurve(h uintptr) *QEasingCurve {
	if h == 0 {
		return nil
	}
	return &QEasingCurve{h: h}
}

func newownedQEasingCurve(h uintptr) *QEasingCurve {
	o := bindQEasingCurve(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QEasingCurve).Delete)
	}
	return o
}

func (this *QEasingCurve) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QEasingCurve) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQEasingCurve(h unsafe.Pointer) *QEasingCurve { return bindQEasingCurve(uintptr(h)) }
func UnsafeNewQEasingCurveOwned(h unsafe.Pointer) *QEasingCurve { return newownedQEasingCurve(uintptr(h)) }

func (this *QEasingCurve) UnsafeRelease() { this.h = 0 }

func NewQEasingCurve() *QEasingCurve {
	__r, _, _ := purego.SyscallN(addr_QEasingCurve_new)
	__o := newownedQEasingCurve(__r)
	return __o
}

func NewQEasingCurve2(other *QEasingCurve) *QEasingCurve {
	__r, _, _ := purego.SyscallN(addr_QEasingCurve_new2, uintptr(other.UnsafePointer()))
	__o := newownedQEasingCurve(__r)
	runtime.KeepAlive(other)
	return __o
}

func NewQEasingCurve3(type_ QEasingCurve__Type) *QEasingCurve {
	__r, _, _ := purego.SyscallN(addr_QEasingCurve_new3, uintptr(type_))
	__o := newownedQEasingCurve(__r)
	return __o
}

func (this *QEasingCurve) OperatorAssign(other *QEasingCurve) *QEasingCurve {
	__r, _, _ := purego.SyscallN(addr_QEasingCurve_OperatorAssign, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return bindQEasingCurve(__r)
}

func (this *QEasingCurve) Swap(other *QEasingCurve) {
	purego.SyscallN(addr_QEasingCurve_swap, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
}

func (this *QEasingCurve) Amplitude() float64 {
	__r := pg_QEasingCurve_amplitude(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QEasingCurve) SetAmplitude(amplitude float64) {
	pg_QEasingCurve_setAmplitude(this.h, amplitude)
	runtime.KeepAlive(this)
}

func (this *QEasingCurve) Period() float64 {
	__r := pg_QEasingCurve_period(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QEasingCurve) SetPeriod(period float64) {
	pg_QEasingCurve_setPeriod(this.h, period)
	runtime.KeepAlive(this)
}

func (this *QEasingCurve) Overshoot() float64 {
	__r := pg_QEasingCurve_overshoot(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QEasingCurve) SetOvershoot(overshoot float64) {
	pg_QEasingCurve_setOvershoot(this.h, overshoot)
	runtime.KeepAlive(this)
}

func (this *QEasingCurve) AddCubicBezierSegment(c1 *QPointF, c2 *QPointF, endPoint *QPointF) {
	purego.SyscallN(addr_QEasingCurve_addCubicBezierSegment, this.h, uintptr(c1.UnsafePointer()), uintptr(c2.UnsafePointer()), uintptr(endPoint.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(c1)
	runtime.KeepAlive(c2)
	runtime.KeepAlive(endPoint)
}

func (this *QEasingCurve) AddTCBSegment(nextPoint *QPointF, t float64, c float64, b float64) {
	pg_QEasingCurve_addTCBSegment(this.h, uintptr(nextPoint.UnsafePointer()), t, c, b)
	runtime.KeepAlive(this)
	runtime.KeepAlive(nextPoint)
}

func (this *QEasingCurve) ToCubicSpline() []*QPointF {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QEasingCurve_toCubicSpline, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QPointF, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQPointF(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QEasingCurve) Type() QEasingCurve__Type {
	__r0, _, _ := purego.SyscallN(addr_QEasingCurve_type, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QEasingCurve__Type(__r)
}

func (this *QEasingCurve) SetType(type_ QEasingCurve__Type) {
	purego.SyscallN(addr_QEasingCurve_setType, this.h, uintptr(type_))
	runtime.KeepAlive(this)
}

func (this *QEasingCurve) ValueForProgress(progress float64) float64 {
	__r := pg_QEasingCurve_valueForProgress(this.h, progress)
	runtime.KeepAlive(this)
	return __r
}

func (this *QEasingCurve) OperatorEqual(rhs *QEasingCurve) bool {
	__r0, _, _ := purego.SyscallN(addr_QEasingCurve_OperatorEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QEasingCurve) OperatorNotEqual(rhs *QEasingCurve) bool {
	__r0, _, _ := purego.SyscallN(addr_QEasingCurve_OperatorNotEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QEasingCurve) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QEasingCurve_delete(this.h)
	}
	this.UnsafeRelease()
}
