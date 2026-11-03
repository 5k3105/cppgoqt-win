package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QContactsPermission_new uintptr
	addr_QContactsPermission_new2 uintptr
	addr_QContactsPermission_setAccessMode uintptr
	addr_QContactsPermission_accessMode uintptr
	addr_QContactsPermission_OperatorAssign uintptr
	addr_QContactsPermission_swap uintptr
	pg_QContactsPermission_delete func(uintptr)
)

func init() {
	symaddr(&addr_QContactsPermission_new, "QContactsPermission_new")
	symaddr(&addr_QContactsPermission_new2, "QContactsPermission_new2")
	symaddr(&addr_QContactsPermission_setAccessMode, "QContactsPermission_setAccessMode")
	symaddr(&addr_QContactsPermission_accessMode, "QContactsPermission_accessMode")
	symaddr(&addr_QContactsPermission_OperatorAssign, "QContactsPermission_OperatorAssign")
	symaddr(&addr_QContactsPermission_swap, "QContactsPermission_swap")
	reg(&pg_QContactsPermission_delete, "QContactsPermission_delete")
}

type QContactsPermission struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QContactsPermission) UnsafeGoID() int64     { return this.vid }
func (this *QContactsPermission) UnsafeSetGoID(id int64) { this.vid = id }

func bindQContactsPermission(h uintptr) *QContactsPermission {
	if h == 0 {
		return nil
	}
	return &QContactsPermission{h: h}
}

func newownedQContactsPermission(h uintptr) *QContactsPermission {
	o := bindQContactsPermission(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QContactsPermission).Delete)
	}
	return o
}

func (this *QContactsPermission) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QContactsPermission) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQContactsPermission(h unsafe.Pointer) *QContactsPermission { return bindQContactsPermission(uintptr(h)) }
func UnsafeNewQContactsPermissionOwned(h unsafe.Pointer) *QContactsPermission { return newownedQContactsPermission(uintptr(h)) }

func (this *QContactsPermission) UnsafeRelease() { this.h = 0 }

func NewQContactsPermission() *QContactsPermission {
	__r, _, _ := purego.SyscallN(addr_QContactsPermission_new)
	__o := newownedQContactsPermission(__r)
	return __o
}

func NewQContactsPermission2(other *QContactsPermission) *QContactsPermission {
	__r, _, _ := purego.SyscallN(addr_QContactsPermission_new2, uintptr(other.UnsafePointer()))
	__o := newownedQContactsPermission(__r)
	runtime.KeepAlive(other)
	return __o
}

func (this *QContactsPermission) SetAccessMode(mode QContactsPermission__AccessMode) {
	purego.SyscallN(addr_QContactsPermission_setAccessMode, this.h, uintptr(mode))
	runtime.KeepAlive(this)
}

func (this *QContactsPermission) AccessMode() QContactsPermission__AccessMode {
	__r0, _, _ := purego.SyscallN(addr_QContactsPermission_accessMode, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QContactsPermission__AccessMode(__r)
}

func (this *QContactsPermission) OperatorAssign(other *QContactsPermission) *QContactsPermission {
	__r, _, _ := purego.SyscallN(addr_QContactsPermission_OperatorAssign, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return bindQContactsPermission(__r)
}

func (this *QContactsPermission) Swap(other *QContactsPermission) {
	purego.SyscallN(addr_QContactsPermission_swap, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
}

func (this *QContactsPermission) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QContactsPermission_delete(this.h)
	}
	this.UnsafeRelease()
}
