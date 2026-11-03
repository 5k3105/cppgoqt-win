package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QCameraPermission_new uintptr
	addr_QCameraPermission_new2 uintptr
	addr_QCameraPermission_OperatorAssign uintptr
	addr_QCameraPermission_swap uintptr
	pg_QCameraPermission_delete func(uintptr)
)

func init() {
	symaddr(&addr_QCameraPermission_new, "QCameraPermission_new")
	symaddr(&addr_QCameraPermission_new2, "QCameraPermission_new2")
	symaddr(&addr_QCameraPermission_OperatorAssign, "QCameraPermission_OperatorAssign")
	symaddr(&addr_QCameraPermission_swap, "QCameraPermission_swap")
	reg(&pg_QCameraPermission_delete, "QCameraPermission_delete")
}

type QCameraPermission struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QCameraPermission) UnsafeGoID() int64     { return this.vid }
func (this *QCameraPermission) UnsafeSetGoID(id int64) { this.vid = id }

func bindQCameraPermission(h uintptr) *QCameraPermission {
	if h == 0 {
		return nil
	}
	return &QCameraPermission{h: h}
}

func newownedQCameraPermission(h uintptr) *QCameraPermission {
	o := bindQCameraPermission(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QCameraPermission).Delete)
	}
	return o
}

func (this *QCameraPermission) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QCameraPermission) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQCameraPermission(h unsafe.Pointer) *QCameraPermission { return bindQCameraPermission(uintptr(h)) }
func UnsafeNewQCameraPermissionOwned(h unsafe.Pointer) *QCameraPermission { return newownedQCameraPermission(uintptr(h)) }

func (this *QCameraPermission) UnsafeRelease() { this.h = 0 }

func NewQCameraPermission() *QCameraPermission {
	__r, _, _ := purego.SyscallN(addr_QCameraPermission_new)
	__o := newownedQCameraPermission(__r)
	return __o
}

func NewQCameraPermission2(other *QCameraPermission) *QCameraPermission {
	__r, _, _ := purego.SyscallN(addr_QCameraPermission_new2, uintptr(other.UnsafePointer()))
	__o := newownedQCameraPermission(__r)
	runtime.KeepAlive(other)
	return __o
}

func (this *QCameraPermission) OperatorAssign(other *QCameraPermission) *QCameraPermission {
	__r, _, _ := purego.SyscallN(addr_QCameraPermission_OperatorAssign, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return bindQCameraPermission(__r)
}

func (this *QCameraPermission) Swap(other *QCameraPermission) {
	purego.SyscallN(addr_QCameraPermission_swap, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
}

func (this *QCameraPermission) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QCameraPermission_delete(this.h)
	}
	this.UnsafeRelease()
}
