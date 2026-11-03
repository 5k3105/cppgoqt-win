package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QBluetoothPermission_new uintptr
	addr_QBluetoothPermission_new2 uintptr
	addr_QBluetoothPermission_setCommunicationModes uintptr
	addr_QBluetoothPermission_communicationModes uintptr
	addr_QBluetoothPermission_OperatorAssign uintptr
	addr_QBluetoothPermission_swap uintptr
	pg_QBluetoothPermission_delete func(uintptr)
)

func init() {
	symaddr(&addr_QBluetoothPermission_new, "QBluetoothPermission_new")
	symaddr(&addr_QBluetoothPermission_new2, "QBluetoothPermission_new2")
	symaddr(&addr_QBluetoothPermission_setCommunicationModes, "QBluetoothPermission_setCommunicationModes")
	symaddr(&addr_QBluetoothPermission_communicationModes, "QBluetoothPermission_communicationModes")
	symaddr(&addr_QBluetoothPermission_OperatorAssign, "QBluetoothPermission_OperatorAssign")
	symaddr(&addr_QBluetoothPermission_swap, "QBluetoothPermission_swap")
	reg(&pg_QBluetoothPermission_delete, "QBluetoothPermission_delete")
}

type QBluetoothPermission struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QBluetoothPermission) UnsafeGoID() int64     { return this.vid }
func (this *QBluetoothPermission) UnsafeSetGoID(id int64) { this.vid = id }

func bindQBluetoothPermission(h uintptr) *QBluetoothPermission {
	if h == 0 {
		return nil
	}
	return &QBluetoothPermission{h: h}
}

func newownedQBluetoothPermission(h uintptr) *QBluetoothPermission {
	o := bindQBluetoothPermission(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QBluetoothPermission).Delete)
	}
	return o
}

func (this *QBluetoothPermission) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QBluetoothPermission) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQBluetoothPermission(h unsafe.Pointer) *QBluetoothPermission { return bindQBluetoothPermission(uintptr(h)) }
func UnsafeNewQBluetoothPermissionOwned(h unsafe.Pointer) *QBluetoothPermission { return newownedQBluetoothPermission(uintptr(h)) }

func (this *QBluetoothPermission) UnsafeRelease() { this.h = 0 }

func NewQBluetoothPermission() *QBluetoothPermission {
	__r, _, _ := purego.SyscallN(addr_QBluetoothPermission_new)
	__o := newownedQBluetoothPermission(__r)
	return __o
}

func NewQBluetoothPermission2(other *QBluetoothPermission) *QBluetoothPermission {
	__r, _, _ := purego.SyscallN(addr_QBluetoothPermission_new2, uintptr(other.UnsafePointer()))
	__o := newownedQBluetoothPermission(__r)
	runtime.KeepAlive(other)
	return __o
}

func (this *QBluetoothPermission) SetCommunicationModes(modes QBluetoothPermission__CommunicationMode) {
	purego.SyscallN(addr_QBluetoothPermission_setCommunicationModes, this.h, uintptr(modes))
	runtime.KeepAlive(this)
}

func (this *QBluetoothPermission) CommunicationModes() QBluetoothPermission__CommunicationMode {
	__r0, _, _ := purego.SyscallN(addr_QBluetoothPermission_communicationModes, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QBluetoothPermission__CommunicationMode(__r)
}

func (this *QBluetoothPermission) OperatorAssign(other *QBluetoothPermission) *QBluetoothPermission {
	__r, _, _ := purego.SyscallN(addr_QBluetoothPermission_OperatorAssign, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return bindQBluetoothPermission(__r)
}

func (this *QBluetoothPermission) Swap(other *QBluetoothPermission) {
	purego.SyscallN(addr_QBluetoothPermission_swap, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
}

func (this *QBluetoothPermission) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QBluetoothPermission_delete(this.h)
	}
	this.UnsafeRelease()
}
