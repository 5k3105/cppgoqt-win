package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QGradient_new uintptr
	addr_QGradient_new2 uintptr
	addr_QGradient_new3 uintptr
	addr_QGradient_type uintptr
	addr_QGradient_setSpread uintptr
	addr_QGradient_spread uintptr
	pg_QGradient_setColorAt func(uintptr, float64, uintptr)
	addr_QGradient_coordinateMode uintptr
	addr_QGradient_setCoordinateMode uintptr
	addr_QGradient_interpolationMode uintptr
	addr_QGradient_setInterpolationMode uintptr
	addr_QGradient_OperatorEqual uintptr
	addr_QGradient_OperatorNotEqual uintptr
	pg_QGradient_delete func(uintptr)
)

func init() {
	symaddr(&addr_QGradient_new, "QGradient_new")
	symaddr(&addr_QGradient_new2, "QGradient_new2")
	symaddr(&addr_QGradient_new3, "QGradient_new3")
	symaddr(&addr_QGradient_type, "QGradient_type")
	symaddr(&addr_QGradient_setSpread, "QGradient_setSpread")
	symaddr(&addr_QGradient_spread, "QGradient_spread")
	reg(&pg_QGradient_setColorAt, "QGradient_setColorAt")
	symaddr(&addr_QGradient_coordinateMode, "QGradient_coordinateMode")
	symaddr(&addr_QGradient_setCoordinateMode, "QGradient_setCoordinateMode")
	symaddr(&addr_QGradient_interpolationMode, "QGradient_interpolationMode")
	symaddr(&addr_QGradient_setInterpolationMode, "QGradient_setInterpolationMode")
	symaddr(&addr_QGradient_OperatorEqual, "QGradient_OperatorEqual")
	symaddr(&addr_QGradient_OperatorNotEqual, "QGradient_OperatorNotEqual")
	reg(&pg_QGradient_delete, "QGradient_delete")
}

type QGradient struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QGradient) UnsafeGoID() int64     { return this.vid }
func (this *QGradient) UnsafeSetGoID(id int64) { this.vid = id }

func bindQGradient(h uintptr) *QGradient {
	if h == 0 {
		return nil
	}
	return &QGradient{h: h}
}

func newownedQGradient(h uintptr) *QGradient {
	o := bindQGradient(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QGradient).Delete)
	}
	return o
}

func (this *QGradient) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QGradient) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQGradient(h unsafe.Pointer) *QGradient { return bindQGradient(uintptr(h)) }
func UnsafeNewQGradientOwned(h unsafe.Pointer) *QGradient { return newownedQGradient(uintptr(h)) }

func (this *QGradient) UnsafeRelease() { this.h = 0 }

func NewQGradient() *QGradient {
	__r, _, _ := purego.SyscallN(addr_QGradient_new)
	__o := newownedQGradient(__r)
	return __o
}

func NewQGradient2(arg1 QGradient__Preset) *QGradient {
	__r, _, _ := purego.SyscallN(addr_QGradient_new2, uintptr(arg1))
	__o := newownedQGradient(__r)
	return __o
}

func NewQGradient3(arg1 *QGradient) *QGradient {
	__r, _, _ := purego.SyscallN(addr_QGradient_new3, uintptr(arg1.UnsafePointer()))
	__o := newownedQGradient(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QGradient) Type() QGradient__Type {
	__r0, _, _ := purego.SyscallN(addr_QGradient_type, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QGradient__Type(__r)
}

func (this *QGradient) SetSpread(spread QGradient__Spread) {
	purego.SyscallN(addr_QGradient_setSpread, this.h, uintptr(spread))
	runtime.KeepAlive(this)
}

func (this *QGradient) Spread() QGradient__Spread {
	__r0, _, _ := purego.SyscallN(addr_QGradient_spread, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QGradient__Spread(__r)
}

func (this *QGradient) SetColorAt(pos float64, color *QColor) {
	pg_QGradient_setColorAt(this.h, pos, uintptr(color.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(color)
}

func (this *QGradient) CoordinateMode() QGradient__CoordinateMode {
	__r0, _, _ := purego.SyscallN(addr_QGradient_coordinateMode, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QGradient__CoordinateMode(__r)
}

func (this *QGradient) SetCoordinateMode(mode QGradient__CoordinateMode) {
	purego.SyscallN(addr_QGradient_setCoordinateMode, this.h, uintptr(mode))
	runtime.KeepAlive(this)
}

func (this *QGradient) InterpolationMode() QGradient__InterpolationMode {
	__r0, _, _ := purego.SyscallN(addr_QGradient_interpolationMode, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QGradient__InterpolationMode(__r)
}

func (this *QGradient) SetInterpolationMode(mode QGradient__InterpolationMode) {
	purego.SyscallN(addr_QGradient_setInterpolationMode, this.h, uintptr(mode))
	runtime.KeepAlive(this)
}

func (this *QGradient) OperatorEqual(gradient *QGradient) bool {
	__r0, _, _ := purego.SyscallN(addr_QGradient_OperatorEqual, this.h, uintptr(gradient.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(gradient)
	return __r != 0
}

func (this *QGradient) OperatorNotEqual(other *QGradient) bool {
	__r0, _, _ := purego.SyscallN(addr_QGradient_OperatorNotEqual, this.h, uintptr(other.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return __r != 0
}

func (this *QGradient) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QGradient_delete(this.h)
	}
	this.UnsafeRelease()
}
