package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QRadialGradient_base func(uintptr) uintptr
	addr_QRadialGradient_new uintptr
	pg_QRadialGradient_new2 func(uintptr, float64, uintptr) uintptr
	pg_QRadialGradient_new3 func(float64, float64, float64, float64, float64) uintptr
	pg_QRadialGradient_new4 func(uintptr, float64) uintptr
	pg_QRadialGradient_new5 func(float64, float64, float64) uintptr
	pg_QRadialGradient_new6 func(uintptr, float64, uintptr, float64) uintptr
	pg_QRadialGradient_new7 func(float64, float64, float64, float64, float64, float64) uintptr
	addr_QRadialGradient_new8 uintptr
	addr_QRadialGradient_center uintptr
	addr_QRadialGradient_setCenter uintptr
	pg_QRadialGradient_setCenter2 func(uintptr, float64, float64)
	addr_QRadialGradient_focalPoint uintptr
	addr_QRadialGradient_setFocalPoint uintptr
	pg_QRadialGradient_setFocalPoint2 func(uintptr, float64, float64)
	pg_QRadialGradient_radius func(uintptr) float64
	pg_QRadialGradient_setRadius func(uintptr, float64)
	pg_QRadialGradient_centerRadius func(uintptr) float64
	pg_QRadialGradient_setCenterRadius func(uintptr, float64)
	pg_QRadialGradient_focalRadius func(uintptr) float64
	pg_QRadialGradient_setFocalRadius func(uintptr, float64)
	pg_QRadialGradient_delete func(uintptr)
)

func init() {
	reg(&pg_QRadialGradient_base, "QRadialGradient_base")
	symaddr(&addr_QRadialGradient_new, "QRadialGradient_new")
	reg(&pg_QRadialGradient_new2, "QRadialGradient_new2")
	reg(&pg_QRadialGradient_new3, "QRadialGradient_new3")
	reg(&pg_QRadialGradient_new4, "QRadialGradient_new4")
	reg(&pg_QRadialGradient_new5, "QRadialGradient_new5")
	reg(&pg_QRadialGradient_new6, "QRadialGradient_new6")
	reg(&pg_QRadialGradient_new7, "QRadialGradient_new7")
	symaddr(&addr_QRadialGradient_new8, "QRadialGradient_new8")
	symaddr(&addr_QRadialGradient_center, "QRadialGradient_center")
	symaddr(&addr_QRadialGradient_setCenter, "QRadialGradient_setCenter")
	reg(&pg_QRadialGradient_setCenter2, "QRadialGradient_setCenter2")
	symaddr(&addr_QRadialGradient_focalPoint, "QRadialGradient_focalPoint")
	symaddr(&addr_QRadialGradient_setFocalPoint, "QRadialGradient_setFocalPoint")
	reg(&pg_QRadialGradient_setFocalPoint2, "QRadialGradient_setFocalPoint2")
	reg(&pg_QRadialGradient_radius, "QRadialGradient_radius")
	reg(&pg_QRadialGradient_setRadius, "QRadialGradient_setRadius")
	reg(&pg_QRadialGradient_centerRadius, "QRadialGradient_centerRadius")
	reg(&pg_QRadialGradient_setCenterRadius, "QRadialGradient_setCenterRadius")
	reg(&pg_QRadialGradient_focalRadius, "QRadialGradient_focalRadius")
	reg(&pg_QRadialGradient_setFocalRadius, "QRadialGradient_setFocalRadius")
	reg(&pg_QRadialGradient_delete, "QRadialGradient_delete")
}

type QRadialGradient struct {
	h uintptr
	*QGradient
}

func bindQRadialGradient(h uintptr) *QRadialGradient {
	if h == 0 {
		return nil
	}
	return &QRadialGradient{h: h, QGradient: bindQGradient(pg_QRadialGradient_base(h))}
}

func newownedQRadialGradient(h uintptr) *QRadialGradient {
	o := bindQRadialGradient(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QRadialGradient).Delete)
	}
	return o
}

func (this *QRadialGradient) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QRadialGradient) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQRadialGradient(h unsafe.Pointer) *QRadialGradient { return bindQRadialGradient(uintptr(h)) }
func UnsafeNewQRadialGradientOwned(h unsafe.Pointer) *QRadialGradient { return newownedQRadialGradient(uintptr(h)) }

func (this *QRadialGradient) UnsafeRelease() {
	this.h = 0
	if this.QGradient != nil {
		this.QGradient.UnsafeRelease()
	}
}

func NewQRadialGradient() *QRadialGradient {
	__r, _, _ := purego.SyscallN(addr_QRadialGradient_new)
	__o := newownedQRadialGradient(__r)
	return __o
}

func NewQRadialGradient2(center *QPointF, radius float64, focalPoint *QPointF) *QRadialGradient {
	__o := newownedQRadialGradient(pg_QRadialGradient_new2(uintptr(center.UnsafePointer()), radius, uintptr(focalPoint.UnsafePointer())))
	runtime.KeepAlive(center)
	runtime.KeepAlive(focalPoint)
	return __o
}

func NewQRadialGradient3(cx float64, cy float64, radius float64, fx float64, fy float64) *QRadialGradient {
	__o := newownedQRadialGradient(pg_QRadialGradient_new3(cx, cy, radius, fx, fy))
	return __o
}

func NewQRadialGradient4(center *QPointF, radius float64) *QRadialGradient {
	__o := newownedQRadialGradient(pg_QRadialGradient_new4(uintptr(center.UnsafePointer()), radius))
	runtime.KeepAlive(center)
	return __o
}

func NewQRadialGradient5(cx float64, cy float64, radius float64) *QRadialGradient {
	__o := newownedQRadialGradient(pg_QRadialGradient_new5(cx, cy, radius))
	return __o
}

func NewQRadialGradient6(center *QPointF, centerRadius float64, focalPoint *QPointF, focalRadius float64) *QRadialGradient {
	__o := newownedQRadialGradient(pg_QRadialGradient_new6(uintptr(center.UnsafePointer()), centerRadius, uintptr(focalPoint.UnsafePointer()), focalRadius))
	runtime.KeepAlive(center)
	runtime.KeepAlive(focalPoint)
	return __o
}

func NewQRadialGradient7(cx float64, cy float64, centerRadius float64, fx float64, fy float64, focalRadius float64) *QRadialGradient {
	__o := newownedQRadialGradient(pg_QRadialGradient_new7(cx, cy, centerRadius, fx, fy, focalRadius))
	return __o
}

func NewQRadialGradient8(arg1 *QRadialGradient) *QRadialGradient {
	__r, _, _ := purego.SyscallN(addr_QRadialGradient_new8, uintptr(arg1.UnsafePointer()))
	__o := newownedQRadialGradient(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QRadialGradient) Center() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QRadialGradient_center, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QRadialGradient) SetCenter(center *QPointF) {
	purego.SyscallN(addr_QRadialGradient_setCenter, this.h, uintptr(center.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(center)
}

func (this *QRadialGradient) SetCenter2(x float64, y float64) {
	pg_QRadialGradient_setCenter2(this.h, x, y)
	runtime.KeepAlive(this)
}

func (this *QRadialGradient) FocalPoint() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QRadialGradient_focalPoint, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QRadialGradient) SetFocalPoint(focalPoint *QPointF) {
	purego.SyscallN(addr_QRadialGradient_setFocalPoint, this.h, uintptr(focalPoint.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(focalPoint)
}

func (this *QRadialGradient) SetFocalPoint2(x float64, y float64) {
	pg_QRadialGradient_setFocalPoint2(this.h, x, y)
	runtime.KeepAlive(this)
}

func (this *QRadialGradient) Radius() float64 {
	__r := pg_QRadialGradient_radius(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QRadialGradient) SetRadius(radius float64) {
	pg_QRadialGradient_setRadius(this.h, radius)
	runtime.KeepAlive(this)
}

func (this *QRadialGradient) CenterRadius() float64 {
	__r := pg_QRadialGradient_centerRadius(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QRadialGradient) SetCenterRadius(radius float64) {
	pg_QRadialGradient_setCenterRadius(this.h, radius)
	runtime.KeepAlive(this)
}

func (this *QRadialGradient) FocalRadius() float64 {
	__r := pg_QRadialGradient_focalRadius(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QRadialGradient) SetFocalRadius(radius float64) {
	pg_QRadialGradient_setFocalRadius(this.h, radius)
	runtime.KeepAlive(this)
}

func (this *QRadialGradient) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QRadialGradient_delete(this.h)
	}
	this.UnsafeRelease()
}
