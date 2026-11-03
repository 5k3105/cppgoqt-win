package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QQuaternion__Axes_x_get uintptr
	addr_QQuaternion__Axes_x_set uintptr
	addr_QQuaternion__Axes_y_get uintptr
	addr_QQuaternion__Axes_y_set uintptr
	addr_QQuaternion__Axes_z_get uintptr
	addr_QQuaternion__Axes_z_set uintptr
	pg_QQuaternion__Axes_delete func(uintptr)
)

func init() {
	symaddr(&addr_QQuaternion__Axes_x_get, "QQuaternion__Axes_x_get")
	symaddr(&addr_QQuaternion__Axes_x_set, "QQuaternion__Axes_x_set")
	symaddr(&addr_QQuaternion__Axes_y_get, "QQuaternion__Axes_y_get")
	symaddr(&addr_QQuaternion__Axes_y_set, "QQuaternion__Axes_y_set")
	symaddr(&addr_QQuaternion__Axes_z_get, "QQuaternion__Axes_z_get")
	symaddr(&addr_QQuaternion__Axes_z_set, "QQuaternion__Axes_z_set")
	reg(&pg_QQuaternion__Axes_delete, "QQuaternion__Axes_delete")
}

type QQuaternion__Axes struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QQuaternion__Axes) UnsafeGoID() int64     { return this.vid }
func (this *QQuaternion__Axes) UnsafeSetGoID(id int64) { this.vid = id }

func bindQQuaternion__Axes(h uintptr) *QQuaternion__Axes {
	if h == 0 {
		return nil
	}
	return &QQuaternion__Axes{h: h}
}

func newownedQQuaternion__Axes(h uintptr) *QQuaternion__Axes {
	o := bindQQuaternion__Axes(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QQuaternion__Axes).Delete)
	}
	return o
}

func (this *QQuaternion__Axes) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QQuaternion__Axes) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQQuaternion__Axes(h unsafe.Pointer) *QQuaternion__Axes { return bindQQuaternion__Axes(uintptr(h)) }
func UnsafeNewQQuaternion__AxesOwned(h unsafe.Pointer) *QQuaternion__Axes { return newownedQQuaternion__Axes(uintptr(h)) }

func (this *QQuaternion__Axes) UnsafeRelease() { this.h = 0 }

func (this *QQuaternion__Axes) X() *QQuaternion__Axis {
	__r, _, _ := purego.SyscallN(addr_QQuaternion__Axes_x_get, this.h)
	runtime.KeepAlive(this)
	return newownedQQuaternion__Axis(__r)
}

func (this *QQuaternion__Axes) SetX(x *QQuaternion__Axis) {
	purego.SyscallN(addr_QQuaternion__Axes_x_set, this.h, uintptr(x.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(x)
}

func (this *QQuaternion__Axes) Y() *QQuaternion__Axis {
	__r, _, _ := purego.SyscallN(addr_QQuaternion__Axes_y_get, this.h)
	runtime.KeepAlive(this)
	return newownedQQuaternion__Axis(__r)
}

func (this *QQuaternion__Axes) SetY(y *QQuaternion__Axis) {
	purego.SyscallN(addr_QQuaternion__Axes_y_set, this.h, uintptr(y.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(y)
}

func (this *QQuaternion__Axes) Z() *QQuaternion__Axis {
	__r, _, _ := purego.SyscallN(addr_QQuaternion__Axes_z_get, this.h)
	runtime.KeepAlive(this)
	return newownedQQuaternion__Axis(__r)
}

func (this *QQuaternion__Axes) SetZ(z *QQuaternion__Axis) {
	purego.SyscallN(addr_QQuaternion__Axes_z_set, this.h, uintptr(z.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(z)
}

func (this *QQuaternion__Axes) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QQuaternion__Axes_delete(this.h)
	}
	this.UnsafeRelease()
}
