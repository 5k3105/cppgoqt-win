package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QQuaternion__Axis_fromVector3D uintptr
	addr_QQuaternion__Axis_toVector3D uintptr
	pg_QQuaternion__Axis_x_get func(uintptr) float32
	pg_QQuaternion__Axis_x_set func(uintptr, float32)
	pg_QQuaternion__Axis_y_get func(uintptr) float32
	pg_QQuaternion__Axis_y_set func(uintptr, float32)
	pg_QQuaternion__Axis_z_get func(uintptr) float32
	pg_QQuaternion__Axis_z_set func(uintptr, float32)
	pg_QQuaternion__Axis_delete func(uintptr)
)

func init() {
	symaddr(&addr_QQuaternion__Axis_fromVector3D, "QQuaternion__Axis_fromVector3D")
	symaddr(&addr_QQuaternion__Axis_toVector3D, "QQuaternion__Axis_toVector3D")
	reg(&pg_QQuaternion__Axis_x_get, "QQuaternion__Axis_x_get")
	reg(&pg_QQuaternion__Axis_x_set, "QQuaternion__Axis_x_set")
	reg(&pg_QQuaternion__Axis_y_get, "QQuaternion__Axis_y_get")
	reg(&pg_QQuaternion__Axis_y_set, "QQuaternion__Axis_y_set")
	reg(&pg_QQuaternion__Axis_z_get, "QQuaternion__Axis_z_get")
	reg(&pg_QQuaternion__Axis_z_set, "QQuaternion__Axis_z_set")
	reg(&pg_QQuaternion__Axis_delete, "QQuaternion__Axis_delete")
}

type QQuaternion__Axis struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QQuaternion__Axis) UnsafeGoID() int64     { return this.vid }
func (this *QQuaternion__Axis) UnsafeSetGoID(id int64) { this.vid = id }

func bindQQuaternion__Axis(h uintptr) *QQuaternion__Axis {
	if h == 0 {
		return nil
	}
	return &QQuaternion__Axis{h: h}
}

func newownedQQuaternion__Axis(h uintptr) *QQuaternion__Axis {
	o := bindQQuaternion__Axis(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QQuaternion__Axis).Delete)
	}
	return o
}

func (this *QQuaternion__Axis) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QQuaternion__Axis) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQQuaternion__Axis(h unsafe.Pointer) *QQuaternion__Axis { return bindQQuaternion__Axis(uintptr(h)) }
func UnsafeNewQQuaternion__AxisOwned(h unsafe.Pointer) *QQuaternion__Axis { return newownedQQuaternion__Axis(uintptr(h)) }

func (this *QQuaternion__Axis) UnsafeRelease() { this.h = 0 }

func QQuaternion__Axis_FromVector3D(v *QVector3D) *QQuaternion__Axis {
	__r, _, _ := purego.SyscallN(addr_QQuaternion__Axis_fromVector3D, uintptr(v.UnsafePointer()))
	runtime.KeepAlive(v)
	return newownedQQuaternion__Axis(__r)
}

func (this *QQuaternion__Axis) ToVector3D() *QVector3D {
	__r, _, _ := purego.SyscallN(addr_QQuaternion__Axis_toVector3D, this.h)
	runtime.KeepAlive(this)
	return newownedQVector3D(__r)
}

func (this *QQuaternion__Axis) X() float32 {
	__r := pg_QQuaternion__Axis_x_get(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QQuaternion__Axis) SetX(x float32) {
	pg_QQuaternion__Axis_x_set(this.h, x)
	runtime.KeepAlive(this)
}

func (this *QQuaternion__Axis) Y() float32 {
	__r := pg_QQuaternion__Axis_y_get(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QQuaternion__Axis) SetY(y float32) {
	pg_QQuaternion__Axis_y_set(this.h, y)
	runtime.KeepAlive(this)
}

func (this *QQuaternion__Axis) Z() float32 {
	__r := pg_QQuaternion__Axis_z_get(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QQuaternion__Axis) SetZ(z float32) {
	pg_QQuaternion__Axis_z_set(this.h, z)
	runtime.KeepAlive(this)
}

func (this *QQuaternion__Axis) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QQuaternion__Axis_delete(this.h)
	}
	this.UnsafeRelease()
}
