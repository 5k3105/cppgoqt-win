package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QBitRef_ToBool uintptr
	addr_QBitRef_OperatorNot uintptr
	addr_QBitRef_OperatorAssign uintptr
	addr_QBitRef_OperatorAssign2 uintptr
	pg_QBitRef_delete func(uintptr)
)

func init() {
	symaddr(&addr_QBitRef_ToBool, "QBitRef_ToBool")
	symaddr(&addr_QBitRef_OperatorNot, "QBitRef_OperatorNot")
	symaddr(&addr_QBitRef_OperatorAssign, "QBitRef_OperatorAssign")
	symaddr(&addr_QBitRef_OperatorAssign2, "QBitRef_OperatorAssign2")
	reg(&pg_QBitRef_delete, "QBitRef_delete")
}

type QBitRef struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QBitRef) UnsafeGoID() int64     { return this.vid }
func (this *QBitRef) UnsafeSetGoID(id int64) { this.vid = id }

func bindQBitRef(h uintptr) *QBitRef {
	if h == 0 {
		return nil
	}
	return &QBitRef{h: h}
}

func newownedQBitRef(h uintptr) *QBitRef {
	o := bindQBitRef(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QBitRef).Delete)
	}
	return o
}

func (this *QBitRef) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QBitRef) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQBitRef(h unsafe.Pointer) *QBitRef { return bindQBitRef(uintptr(h)) }
func UnsafeNewQBitRefOwned(h unsafe.Pointer) *QBitRef { return newownedQBitRef(uintptr(h)) }

func (this *QBitRef) UnsafeRelease() { this.h = 0 }

func (this *QBitRef) ToBool() bool {
	__r0, _, _ := purego.SyscallN(addr_QBitRef_ToBool, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QBitRef) OperatorNot() bool {
	__r0, _, _ := purego.SyscallN(addr_QBitRef_OperatorNot, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QBitRef) OperatorAssign(val *QBitRef) *QBitRef {
	__r, _, _ := purego.SyscallN(addr_QBitRef_OperatorAssign, this.h, uintptr(val.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(val)
	return bindQBitRef(__r)
}

func (this *QBitRef) OperatorAssignWithVal(val bool) *QBitRef {
	__r, _, _ := purego.SyscallN(addr_QBitRef_OperatorAssign2, this.h, uintptr(b2i(val)))
	runtime.KeepAlive(this)
	return bindQBitRef(__r)
}

func (this *QBitRef) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QBitRef_delete(this.h)
	}
	this.UnsafeRelease()
}
