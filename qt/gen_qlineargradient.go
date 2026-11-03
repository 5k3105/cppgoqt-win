package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QLinearGradient_base func(uintptr) uintptr
	addr_QLinearGradient_new uintptr
	addr_QLinearGradient_new2 uintptr
	pg_QLinearGradient_new3 func(float64, float64, float64, float64) uintptr
	addr_QLinearGradient_new4 uintptr
	addr_QLinearGradient_start uintptr
	addr_QLinearGradient_setStart uintptr
	pg_QLinearGradient_setStart2 func(uintptr, float64, float64)
	addr_QLinearGradient_finalStop uintptr
	addr_QLinearGradient_setFinalStop uintptr
	pg_QLinearGradient_setFinalStop2 func(uintptr, float64, float64)
	pg_QLinearGradient_delete func(uintptr)
)

func init() {
	reg(&pg_QLinearGradient_base, "QLinearGradient_base")
	symaddr(&addr_QLinearGradient_new, "QLinearGradient_new")
	symaddr(&addr_QLinearGradient_new2, "QLinearGradient_new2")
	reg(&pg_QLinearGradient_new3, "QLinearGradient_new3")
	symaddr(&addr_QLinearGradient_new4, "QLinearGradient_new4")
	symaddr(&addr_QLinearGradient_start, "QLinearGradient_start")
	symaddr(&addr_QLinearGradient_setStart, "QLinearGradient_setStart")
	reg(&pg_QLinearGradient_setStart2, "QLinearGradient_setStart2")
	symaddr(&addr_QLinearGradient_finalStop, "QLinearGradient_finalStop")
	symaddr(&addr_QLinearGradient_setFinalStop, "QLinearGradient_setFinalStop")
	reg(&pg_QLinearGradient_setFinalStop2, "QLinearGradient_setFinalStop2")
	reg(&pg_QLinearGradient_delete, "QLinearGradient_delete")
}

type QLinearGradient struct {
	h uintptr
	*QGradient
}

func bindQLinearGradient(h uintptr) *QLinearGradient {
	if h == 0 {
		return nil
	}
	return &QLinearGradient{h: h, QGradient: bindQGradient(pg_QLinearGradient_base(h))}
}

func newownedQLinearGradient(h uintptr) *QLinearGradient {
	o := bindQLinearGradient(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QLinearGradient).Delete)
	}
	return o
}

func (this *QLinearGradient) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QLinearGradient) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQLinearGradient(h unsafe.Pointer) *QLinearGradient { return bindQLinearGradient(uintptr(h)) }
func UnsafeNewQLinearGradientOwned(h unsafe.Pointer) *QLinearGradient { return newownedQLinearGradient(uintptr(h)) }

func (this *QLinearGradient) UnsafeRelease() {
	this.h = 0
	if this.QGradient != nil {
		this.QGradient.UnsafeRelease()
	}
}

func NewQLinearGradient() *QLinearGradient {
	__r, _, _ := purego.SyscallN(addr_QLinearGradient_new)
	__o := newownedQLinearGradient(__r)
	return __o
}

func NewQLinearGradient2(start *QPointF, finalStop *QPointF) *QLinearGradient {
	__r, _, _ := purego.SyscallN(addr_QLinearGradient_new2, uintptr(start.UnsafePointer()), uintptr(finalStop.UnsafePointer()))
	__o := newownedQLinearGradient(__r)
	runtime.KeepAlive(start)
	runtime.KeepAlive(finalStop)
	return __o
}

func NewQLinearGradient3(xStart float64, yStart float64, xFinalStop float64, yFinalStop float64) *QLinearGradient {
	__o := newownedQLinearGradient(pg_QLinearGradient_new3(xStart, yStart, xFinalStop, yFinalStop))
	return __o
}

func NewQLinearGradient4(arg1 *QLinearGradient) *QLinearGradient {
	__r, _, _ := purego.SyscallN(addr_QLinearGradient_new4, uintptr(arg1.UnsafePointer()))
	__o := newownedQLinearGradient(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QLinearGradient) Start() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QLinearGradient_start, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QLinearGradient) SetStart(start *QPointF) {
	purego.SyscallN(addr_QLinearGradient_setStart, this.h, uintptr(start.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(start)
}

func (this *QLinearGradient) SetStart2(x float64, y float64) {
	pg_QLinearGradient_setStart2(this.h, x, y)
	runtime.KeepAlive(this)
}

func (this *QLinearGradient) FinalStop() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QLinearGradient_finalStop, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QLinearGradient) SetFinalStop(stop *QPointF) {
	purego.SyscallN(addr_QLinearGradient_setFinalStop, this.h, uintptr(stop.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(stop)
}

func (this *QLinearGradient) SetFinalStop2(x float64, y float64) {
	pg_QLinearGradient_setFinalStop2(this.h, x, y)
	runtime.KeepAlive(this)
}

func (this *QLinearGradient) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QLinearGradient_delete(this.h)
	}
	this.UnsafeRelease()
}
