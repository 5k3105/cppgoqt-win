package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QPointingDeviceUniqueId_new uintptr
	addr_QPointingDeviceUniqueId_new2 uintptr
	addr_QPointingDeviceUniqueId_fromNumericId uintptr
	addr_QPointingDeviceUniqueId_isValid uintptr
	addr_QPointingDeviceUniqueId_numericId uintptr
	addr_QPointingDeviceUniqueId_OperatorEqual uintptr
	addr_QPointingDeviceUniqueId_OperatorNotEqual uintptr
	pg_QPointingDeviceUniqueId_delete func(uintptr)
)

func init() {
	symaddr(&addr_QPointingDeviceUniqueId_new, "QPointingDeviceUniqueId_new")
	symaddr(&addr_QPointingDeviceUniqueId_new2, "QPointingDeviceUniqueId_new2")
	symaddr(&addr_QPointingDeviceUniqueId_fromNumericId, "QPointingDeviceUniqueId_fromNumericId")
	symaddr(&addr_QPointingDeviceUniqueId_isValid, "QPointingDeviceUniqueId_isValid")
	symaddr(&addr_QPointingDeviceUniqueId_numericId, "QPointingDeviceUniqueId_numericId")
	symaddr(&addr_QPointingDeviceUniqueId_OperatorEqual, "QPointingDeviceUniqueId_OperatorEqual")
	symaddr(&addr_QPointingDeviceUniqueId_OperatorNotEqual, "QPointingDeviceUniqueId_OperatorNotEqual")
	reg(&pg_QPointingDeviceUniqueId_delete, "QPointingDeviceUniqueId_delete")
}

type QPointingDeviceUniqueId struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QPointingDeviceUniqueId) UnsafeGoID() int64     { return this.vid }
func (this *QPointingDeviceUniqueId) UnsafeSetGoID(id int64) { this.vid = id }

func bindQPointingDeviceUniqueId(h uintptr) *QPointingDeviceUniqueId {
	if h == 0 {
		return nil
	}
	return &QPointingDeviceUniqueId{h: h}
}

func newownedQPointingDeviceUniqueId(h uintptr) *QPointingDeviceUniqueId {
	o := bindQPointingDeviceUniqueId(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QPointingDeviceUniqueId).Delete)
	}
	return o
}

func (this *QPointingDeviceUniqueId) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QPointingDeviceUniqueId) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQPointingDeviceUniqueId(h unsafe.Pointer) *QPointingDeviceUniqueId { return bindQPointingDeviceUniqueId(uintptr(h)) }
func UnsafeNewQPointingDeviceUniqueIdOwned(h unsafe.Pointer) *QPointingDeviceUniqueId { return newownedQPointingDeviceUniqueId(uintptr(h)) }

func (this *QPointingDeviceUniqueId) UnsafeRelease() { this.h = 0 }

func NewQPointingDeviceUniqueId() *QPointingDeviceUniqueId {
	__r, _, _ := purego.SyscallN(addr_QPointingDeviceUniqueId_new)
	__o := newownedQPointingDeviceUniqueId(__r)
	return __o
}

func NewQPointingDeviceUniqueId2(arg1 *QPointingDeviceUniqueId) *QPointingDeviceUniqueId {
	__r, _, _ := purego.SyscallN(addr_QPointingDeviceUniqueId_new2, uintptr(arg1.UnsafePointer()))
	__o := newownedQPointingDeviceUniqueId(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func QPointingDeviceUniqueId_FromNumericId(id int64) *QPointingDeviceUniqueId {
	__r, _, _ := purego.SyscallN(addr_QPointingDeviceUniqueId_fromNumericId, uintptr(id))
	return newownedQPointingDeviceUniqueId(__r)
}

func (this *QPointingDeviceUniqueId) IsValid() bool {
	__r0, _, _ := purego.SyscallN(addr_QPointingDeviceUniqueId_isValid, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QPointingDeviceUniqueId) NumericId() int64 {
	__r0, _, _ := purego.SyscallN(addr_QPointingDeviceUniqueId_numericId, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QPointingDeviceUniqueId) OperatorEqual(rhs *QPointingDeviceUniqueId) bool {
	__r0, _, _ := purego.SyscallN(addr_QPointingDeviceUniqueId_OperatorEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QPointingDeviceUniqueId) OperatorNotEqual(rhs *QPointingDeviceUniqueId) bool {
	__r0, _, _ := purego.SyscallN(addr_QPointingDeviceUniqueId_OperatorNotEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QPointingDeviceUniqueId) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QPointingDeviceUniqueId_delete(this.h)
	}
	this.UnsafeRelease()
}
