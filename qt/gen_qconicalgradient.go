package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QConicalGradient_base func(uintptr) uintptr
	addr_QConicalGradient_new uintptr
	pg_QConicalGradient_new2 func(uintptr, float64) uintptr
	pg_QConicalGradient_new3 func(float64, float64, float64) uintptr
	addr_QConicalGradient_new4 uintptr
	addr_QConicalGradient_center uintptr
	addr_QConicalGradient_setCenter uintptr
	pg_QConicalGradient_setCenter2 func(uintptr, float64, float64)
	pg_QConicalGradient_angle func(uintptr) float64
	pg_QConicalGradient_setAngle func(uintptr, float64)
	pg_QConicalGradient_delete func(uintptr)
)

func init() {
	reg(&pg_QConicalGradient_base, "QConicalGradient_base")
	symaddr(&addr_QConicalGradient_new, "QConicalGradient_new")
	reg(&pg_QConicalGradient_new2, "QConicalGradient_new2")
	reg(&pg_QConicalGradient_new3, "QConicalGradient_new3")
	symaddr(&addr_QConicalGradient_new4, "QConicalGradient_new4")
	symaddr(&addr_QConicalGradient_center, "QConicalGradient_center")
	symaddr(&addr_QConicalGradient_setCenter, "QConicalGradient_setCenter")
	reg(&pg_QConicalGradient_setCenter2, "QConicalGradient_setCenter2")
	reg(&pg_QConicalGradient_angle, "QConicalGradient_angle")
	reg(&pg_QConicalGradient_setAngle, "QConicalGradient_setAngle")
	reg(&pg_QConicalGradient_delete, "QConicalGradient_delete")
}

type QConicalGradient struct {
	h uintptr
	*QGradient
}

func bindQConicalGradient(h uintptr) *QConicalGradient {
	if h == 0 {
		return nil
	}
	return &QConicalGradient{h: h, QGradient: bindQGradient(pg_QConicalGradient_base(h))}
}

func newownedQConicalGradient(h uintptr) *QConicalGradient {
	o := bindQConicalGradient(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QConicalGradient).Delete)
	}
	return o
}

func (this *QConicalGradient) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QConicalGradient) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQConicalGradient(h unsafe.Pointer) *QConicalGradient { return bindQConicalGradient(uintptr(h)) }
func UnsafeNewQConicalGradientOwned(h unsafe.Pointer) *QConicalGradient { return newownedQConicalGradient(uintptr(h)) }

func (this *QConicalGradient) UnsafeRelease() {
	this.h = 0
	if this.QGradient != nil {
		this.QGradient.UnsafeRelease()
	}
}

func NewQConicalGradient() *QConicalGradient {
	__r, _, _ := purego.SyscallN(addr_QConicalGradient_new)
	__o := newownedQConicalGradient(__r)
	return __o
}

func NewQConicalGradient2(center *QPointF, startAngle float64) *QConicalGradient {
	__o := newownedQConicalGradient(pg_QConicalGradient_new2(uintptr(center.UnsafePointer()), startAngle))
	runtime.KeepAlive(center)
	return __o
}

func NewQConicalGradient3(cx float64, cy float64, startAngle float64) *QConicalGradient {
	__o := newownedQConicalGradient(pg_QConicalGradient_new3(cx, cy, startAngle))
	return __o
}

func NewQConicalGradient4(arg1 *QConicalGradient) *QConicalGradient {
	__r, _, _ := purego.SyscallN(addr_QConicalGradient_new4, uintptr(arg1.UnsafePointer()))
	__o := newownedQConicalGradient(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QConicalGradient) Center() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QConicalGradient_center, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QConicalGradient) SetCenter(center *QPointF) {
	purego.SyscallN(addr_QConicalGradient_setCenter, this.h, uintptr(center.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(center)
}

func (this *QConicalGradient) SetCenter2(x float64, y float64) {
	pg_QConicalGradient_setCenter2(this.h, x, y)
	runtime.KeepAlive(this)
}

func (this *QConicalGradient) Angle() float64 {
	__r := pg_QConicalGradient_angle(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QConicalGradient) SetAngle(angle float64) {
	pg_QConicalGradient_setAngle(this.h, angle)
	runtime.KeepAlive(this)
}

func (this *QConicalGradient) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QConicalGradient_delete(this.h)
	}
	this.UnsafeRelease()
}
