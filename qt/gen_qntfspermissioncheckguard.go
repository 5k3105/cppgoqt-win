package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QNtfsPermissionCheckGuard_new uintptr
	pg_QNtfsPermissionCheckGuard_delete func(uintptr)
)

func init() {
	symaddr(&addr_QNtfsPermissionCheckGuard_new, "QNtfsPermissionCheckGuard_new")
	reg(&pg_QNtfsPermissionCheckGuard_delete, "QNtfsPermissionCheckGuard_delete")
}

type QNtfsPermissionCheckGuard struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QNtfsPermissionCheckGuard) UnsafeGoID() int64     { return this.vid }
func (this *QNtfsPermissionCheckGuard) UnsafeSetGoID(id int64) { this.vid = id }

func bindQNtfsPermissionCheckGuard(h uintptr) *QNtfsPermissionCheckGuard {
	if h == 0 {
		return nil
	}
	return &QNtfsPermissionCheckGuard{h: h}
}

func newownedQNtfsPermissionCheckGuard(h uintptr) *QNtfsPermissionCheckGuard {
	o := bindQNtfsPermissionCheckGuard(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QNtfsPermissionCheckGuard).Delete)
	}
	return o
}

func (this *QNtfsPermissionCheckGuard) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QNtfsPermissionCheckGuard) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQNtfsPermissionCheckGuard(h unsafe.Pointer) *QNtfsPermissionCheckGuard { return bindQNtfsPermissionCheckGuard(uintptr(h)) }
func UnsafeNewQNtfsPermissionCheckGuardOwned(h unsafe.Pointer) *QNtfsPermissionCheckGuard { return newownedQNtfsPermissionCheckGuard(uintptr(h)) }

func (this *QNtfsPermissionCheckGuard) UnsafeRelease() { this.h = 0 }

func NewQNtfsPermissionCheckGuard() *QNtfsPermissionCheckGuard {
	__r, _, _ := purego.SyscallN(addr_QNtfsPermissionCheckGuard_new)
	__o := newownedQNtfsPermissionCheckGuard(__r)
	return __o
}

func (this *QNtfsPermissionCheckGuard) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QNtfsPermissionCheckGuard_delete(this.h)
	}
	this.UnsafeRelease()
}
