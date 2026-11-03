package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QIODeviceBase_new uintptr
	addr_QIODeviceBase_OperatorAssign uintptr
)

func init() {
	symaddr(&addr_QIODeviceBase_new, "QIODeviceBase_new")
	symaddr(&addr_QIODeviceBase_OperatorAssign, "QIODeviceBase_OperatorAssign")
}

type QIODeviceBase struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QIODeviceBase) UnsafeGoID() int64     { return this.vid }
func (this *QIODeviceBase) UnsafeSetGoID(id int64) { this.vid = id }

func bindQIODeviceBase(h uintptr) *QIODeviceBase {
	if h == 0 {
		return nil
	}
	return &QIODeviceBase{h: h}
}

func newownedQIODeviceBase(h uintptr) *QIODeviceBase { return bindQIODeviceBase(h) }

func (this *QIODeviceBase) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QIODeviceBase) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQIODeviceBase(h unsafe.Pointer) *QIODeviceBase { return bindQIODeviceBase(uintptr(h)) }
func UnsafeNewQIODeviceBaseOwned(h unsafe.Pointer) *QIODeviceBase { return newownedQIODeviceBase(uintptr(h)) }

func (this *QIODeviceBase) UnsafeRelease() { this.h = 0 }

func NewQIODeviceBase(arg1 *QIODeviceBase) *QIODeviceBase {
	__r, _, _ := purego.SyscallN(addr_QIODeviceBase_new, uintptr(arg1.UnsafePointer()))
	__o := newownedQIODeviceBase(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QIODeviceBase) OperatorAssign(arg1 *QIODeviceBase) *QIODeviceBase {
	__r, _, _ := purego.SyscallN(addr_QIODeviceBase_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQIODeviceBase(__r)
}

func (this *QIODeviceBase) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	this.UnsafeRelease()
}
