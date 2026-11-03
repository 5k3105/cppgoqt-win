package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QLocationPermission_new uintptr
	addr_QLocationPermission_new2 uintptr
	addr_QLocationPermission_setAccuracy uintptr
	addr_QLocationPermission_accuracy uintptr
	addr_QLocationPermission_setAvailability uintptr
	addr_QLocationPermission_availability uintptr
	addr_QLocationPermission_OperatorAssign uintptr
	addr_QLocationPermission_swap uintptr
	pg_QLocationPermission_delete func(uintptr)
)

func init() {
	symaddr(&addr_QLocationPermission_new, "QLocationPermission_new")
	symaddr(&addr_QLocationPermission_new2, "QLocationPermission_new2")
	symaddr(&addr_QLocationPermission_setAccuracy, "QLocationPermission_setAccuracy")
	symaddr(&addr_QLocationPermission_accuracy, "QLocationPermission_accuracy")
	symaddr(&addr_QLocationPermission_setAvailability, "QLocationPermission_setAvailability")
	symaddr(&addr_QLocationPermission_availability, "QLocationPermission_availability")
	symaddr(&addr_QLocationPermission_OperatorAssign, "QLocationPermission_OperatorAssign")
	symaddr(&addr_QLocationPermission_swap, "QLocationPermission_swap")
	reg(&pg_QLocationPermission_delete, "QLocationPermission_delete")
}

type QLocationPermission struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QLocationPermission) UnsafeGoID() int64     { return this.vid }
func (this *QLocationPermission) UnsafeSetGoID(id int64) { this.vid = id }

func bindQLocationPermission(h uintptr) *QLocationPermission {
	if h == 0 {
		return nil
	}
	return &QLocationPermission{h: h}
}

func newownedQLocationPermission(h uintptr) *QLocationPermission {
	o := bindQLocationPermission(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QLocationPermission).Delete)
	}
	return o
}

func (this *QLocationPermission) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QLocationPermission) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQLocationPermission(h unsafe.Pointer) *QLocationPermission { return bindQLocationPermission(uintptr(h)) }
func UnsafeNewQLocationPermissionOwned(h unsafe.Pointer) *QLocationPermission { return newownedQLocationPermission(uintptr(h)) }

func (this *QLocationPermission) UnsafeRelease() { this.h = 0 }

func NewQLocationPermission() *QLocationPermission {
	__r, _, _ := purego.SyscallN(addr_QLocationPermission_new)
	__o := newownedQLocationPermission(__r)
	return __o
}

func NewQLocationPermission2(other *QLocationPermission) *QLocationPermission {
	__r, _, _ := purego.SyscallN(addr_QLocationPermission_new2, uintptr(other.UnsafePointer()))
	__o := newownedQLocationPermission(__r)
	runtime.KeepAlive(other)
	return __o
}

func (this *QLocationPermission) SetAccuracy(accuracy QLocationPermission__Accuracy) {
	purego.SyscallN(addr_QLocationPermission_setAccuracy, this.h, uintptr(accuracy))
	runtime.KeepAlive(this)
}

func (this *QLocationPermission) Accuracy() QLocationPermission__Accuracy {
	__r0, _, _ := purego.SyscallN(addr_QLocationPermission_accuracy, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QLocationPermission__Accuracy(__r)
}

func (this *QLocationPermission) SetAvailability(availability QLocationPermission__Availability) {
	purego.SyscallN(addr_QLocationPermission_setAvailability, this.h, uintptr(availability))
	runtime.KeepAlive(this)
}

func (this *QLocationPermission) Availability() QLocationPermission__Availability {
	__r0, _, _ := purego.SyscallN(addr_QLocationPermission_availability, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QLocationPermission__Availability(__r)
}

func (this *QLocationPermission) OperatorAssign(other *QLocationPermission) *QLocationPermission {
	__r, _, _ := purego.SyscallN(addr_QLocationPermission_OperatorAssign, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return bindQLocationPermission(__r)
}

func (this *QLocationPermission) Swap(other *QLocationPermission) {
	purego.SyscallN(addr_QLocationPermission_swap, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
}

func (this *QLocationPermission) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QLocationPermission_delete(this.h)
	}
	this.UnsafeRelease()
}
