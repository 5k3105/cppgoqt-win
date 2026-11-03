package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QGradient__QGradientData_new uintptr
	pg_QGradient__QGradientData_delete func(uintptr)
)

func init() {
	symaddr(&addr_QGradient__QGradientData_new, "QGradient__QGradientData_new")
	reg(&pg_QGradient__QGradientData_delete, "QGradient__QGradientData_delete")
}

type QGradient__QGradientData struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QGradient__QGradientData) UnsafeGoID() int64     { return this.vid }
func (this *QGradient__QGradientData) UnsafeSetGoID(id int64) { this.vid = id }

func bindQGradient__QGradientData(h uintptr) *QGradient__QGradientData {
	if h == 0 {
		return nil
	}
	return &QGradient__QGradientData{h: h}
}

func newownedQGradient__QGradientData(h uintptr) *QGradient__QGradientData {
	o := bindQGradient__QGradientData(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QGradient__QGradientData).Delete)
	}
	return o
}

func (this *QGradient__QGradientData) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QGradient__QGradientData) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQGradient__QGradientData(h unsafe.Pointer) *QGradient__QGradientData { return bindQGradient__QGradientData(uintptr(h)) }
func UnsafeNewQGradient__QGradientDataOwned(h unsafe.Pointer) *QGradient__QGradientData { return newownedQGradient__QGradientData(uintptr(h)) }

func (this *QGradient__QGradientData) UnsafeRelease() { this.h = 0 }

func NewQGradient__QGradientData(arg1 *QGradient__QGradientData) *QGradient__QGradientData {
	__r, _, _ := purego.SyscallN(addr_QGradient__QGradientData_new, uintptr(arg1.UnsafePointer()))
	__o := newownedQGradient__QGradientData(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QGradient__QGradientData) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QGradient__QGradientData_delete(this.h)
	}
	this.UnsafeRelease()
}
