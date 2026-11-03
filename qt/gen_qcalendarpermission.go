package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QCalendarPermission_new uintptr
	addr_QCalendarPermission_new2 uintptr
	addr_QCalendarPermission_setAccessMode uintptr
	addr_QCalendarPermission_accessMode uintptr
	addr_QCalendarPermission_OperatorAssign uintptr
	addr_QCalendarPermission_swap uintptr
	pg_QCalendarPermission_delete func(uintptr)
)

func init() {
	symaddr(&addr_QCalendarPermission_new, "QCalendarPermission_new")
	symaddr(&addr_QCalendarPermission_new2, "QCalendarPermission_new2")
	symaddr(&addr_QCalendarPermission_setAccessMode, "QCalendarPermission_setAccessMode")
	symaddr(&addr_QCalendarPermission_accessMode, "QCalendarPermission_accessMode")
	symaddr(&addr_QCalendarPermission_OperatorAssign, "QCalendarPermission_OperatorAssign")
	symaddr(&addr_QCalendarPermission_swap, "QCalendarPermission_swap")
	reg(&pg_QCalendarPermission_delete, "QCalendarPermission_delete")
}

type QCalendarPermission struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QCalendarPermission) UnsafeGoID() int64     { return this.vid }
func (this *QCalendarPermission) UnsafeSetGoID(id int64) { this.vid = id }

func bindQCalendarPermission(h uintptr) *QCalendarPermission {
	if h == 0 {
		return nil
	}
	return &QCalendarPermission{h: h}
}

func newownedQCalendarPermission(h uintptr) *QCalendarPermission {
	o := bindQCalendarPermission(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QCalendarPermission).Delete)
	}
	return o
}

func (this *QCalendarPermission) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QCalendarPermission) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQCalendarPermission(h unsafe.Pointer) *QCalendarPermission { return bindQCalendarPermission(uintptr(h)) }
func UnsafeNewQCalendarPermissionOwned(h unsafe.Pointer) *QCalendarPermission { return newownedQCalendarPermission(uintptr(h)) }

func (this *QCalendarPermission) UnsafeRelease() { this.h = 0 }

func NewQCalendarPermission() *QCalendarPermission {
	__r, _, _ := purego.SyscallN(addr_QCalendarPermission_new)
	__o := newownedQCalendarPermission(__r)
	return __o
}

func NewQCalendarPermission2(other *QCalendarPermission) *QCalendarPermission {
	__r, _, _ := purego.SyscallN(addr_QCalendarPermission_new2, uintptr(other.UnsafePointer()))
	__o := newownedQCalendarPermission(__r)
	runtime.KeepAlive(other)
	return __o
}

func (this *QCalendarPermission) SetAccessMode(mode QCalendarPermission__AccessMode) {
	purego.SyscallN(addr_QCalendarPermission_setAccessMode, this.h, uintptr(mode))
	runtime.KeepAlive(this)
}

func (this *QCalendarPermission) AccessMode() QCalendarPermission__AccessMode {
	__r0, _, _ := purego.SyscallN(addr_QCalendarPermission_accessMode, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QCalendarPermission__AccessMode(__r)
}

func (this *QCalendarPermission) OperatorAssign(other *QCalendarPermission) *QCalendarPermission {
	__r, _, _ := purego.SyscallN(addr_QCalendarPermission_OperatorAssign, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return bindQCalendarPermission(__r)
}

func (this *QCalendarPermission) Swap(other *QCalendarPermission) {
	purego.SyscallN(addr_QCalendarPermission_swap, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
}

func (this *QCalendarPermission) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QCalendarPermission_delete(this.h)
	}
	this.UnsafeRelease()
}
