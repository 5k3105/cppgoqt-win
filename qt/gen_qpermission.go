package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QPermission_new uintptr
	addr_QPermission_new2 uintptr
	addr_QPermission_status uintptr
	addr_QPermission_type uintptr
	addr_QPermission_OperatorAssign uintptr
	pg_QPermission_delete func(uintptr)
)

func init() {
	symaddr(&addr_QPermission_new, "QPermission_new")
	symaddr(&addr_QPermission_new2, "QPermission_new2")
	symaddr(&addr_QPermission_status, "QPermission_status")
	symaddr(&addr_QPermission_type, "QPermission_type")
	symaddr(&addr_QPermission_OperatorAssign, "QPermission_OperatorAssign")
	reg(&pg_QPermission_delete, "QPermission_delete")
}

type QPermission struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QPermission) UnsafeGoID() int64     { return this.vid }
func (this *QPermission) UnsafeSetGoID(id int64) { this.vid = id }

func bindQPermission(h uintptr) *QPermission {
	if h == 0 {
		return nil
	}
	return &QPermission{h: h}
}

func newownedQPermission(h uintptr) *QPermission {
	o := bindQPermission(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QPermission).Delete)
	}
	return o
}

func (this *QPermission) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QPermission) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQPermission(h unsafe.Pointer) *QPermission { return bindQPermission(uintptr(h)) }
func UnsafeNewQPermissionOwned(h unsafe.Pointer) *QPermission { return newownedQPermission(uintptr(h)) }

func (this *QPermission) UnsafeRelease() { this.h = 0 }

func NewQPermission() *QPermission {
	__r, _, _ := purego.SyscallN(addr_QPermission_new)
	__o := newownedQPermission(__r)
	return __o
}

func NewQPermission2(arg1 *QPermission) *QPermission {
	__r, _, _ := purego.SyscallN(addr_QPermission_new2, uintptr(arg1.UnsafePointer()))
	__o := newownedQPermission(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QPermission) Status() PermissionStatus {
	__r0, _, _ := purego.SyscallN(addr_QPermission_status, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return PermissionStatus(__r)
}

func (this *QPermission) Type() *QMetaType {
	__r, _, _ := purego.SyscallN(addr_QPermission_type, this.h)
	runtime.KeepAlive(this)
	return newownedQMetaType(__r)
}

func (this *QPermission) OperatorAssign(arg1 *QPermission) *QPermission {
	__r, _, _ := purego.SyscallN(addr_QPermission_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQPermission(__r)
}

func (this *QPermission) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QPermission_delete(this.h)
	}
	this.UnsafeRelease()
}
