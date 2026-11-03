package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QMicrophonePermission_new uintptr
	addr_QMicrophonePermission_new2 uintptr
	addr_QMicrophonePermission_OperatorAssign uintptr
	addr_QMicrophonePermission_swap uintptr
	pg_QMicrophonePermission_delete func(uintptr)
)

func init() {
	symaddr(&addr_QMicrophonePermission_new, "QMicrophonePermission_new")
	symaddr(&addr_QMicrophonePermission_new2, "QMicrophonePermission_new2")
	symaddr(&addr_QMicrophonePermission_OperatorAssign, "QMicrophonePermission_OperatorAssign")
	symaddr(&addr_QMicrophonePermission_swap, "QMicrophonePermission_swap")
	reg(&pg_QMicrophonePermission_delete, "QMicrophonePermission_delete")
}

type QMicrophonePermission struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QMicrophonePermission) UnsafeGoID() int64     { return this.vid }
func (this *QMicrophonePermission) UnsafeSetGoID(id int64) { this.vid = id }

func bindQMicrophonePermission(h uintptr) *QMicrophonePermission {
	if h == 0 {
		return nil
	}
	return &QMicrophonePermission{h: h}
}

func newownedQMicrophonePermission(h uintptr) *QMicrophonePermission {
	o := bindQMicrophonePermission(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QMicrophonePermission).Delete)
	}
	return o
}

func (this *QMicrophonePermission) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QMicrophonePermission) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQMicrophonePermission(h unsafe.Pointer) *QMicrophonePermission { return bindQMicrophonePermission(uintptr(h)) }
func UnsafeNewQMicrophonePermissionOwned(h unsafe.Pointer) *QMicrophonePermission { return newownedQMicrophonePermission(uintptr(h)) }

func (this *QMicrophonePermission) UnsafeRelease() { this.h = 0 }

func NewQMicrophonePermission() *QMicrophonePermission {
	__r, _, _ := purego.SyscallN(addr_QMicrophonePermission_new)
	__o := newownedQMicrophonePermission(__r)
	return __o
}

func NewQMicrophonePermission2(other *QMicrophonePermission) *QMicrophonePermission {
	__r, _, _ := purego.SyscallN(addr_QMicrophonePermission_new2, uintptr(other.UnsafePointer()))
	__o := newownedQMicrophonePermission(__r)
	runtime.KeepAlive(other)
	return __o
}

func (this *QMicrophonePermission) OperatorAssign(other *QMicrophonePermission) *QMicrophonePermission {
	__r, _, _ := purego.SyscallN(addr_QMicrophonePermission_OperatorAssign, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return bindQMicrophonePermission(__r)
}

func (this *QMicrophonePermission) Swap(other *QMicrophonePermission) {
	purego.SyscallN(addr_QMicrophonePermission_swap, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
}

func (this *QMicrophonePermission) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QMicrophonePermission_delete(this.h)
	}
	this.UnsafeRelease()
}
