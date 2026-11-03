package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QSocketDescriptor_new uintptr
	addr_QSocketDescriptor_new2 uintptr
	addr_QSocketDescriptor_ToQintptr uintptr
	addr_QSocketDescriptor_isValid uintptr
	addr_QSocketDescriptor_OperatorEqual uintptr
	addr_QSocketDescriptor_OperatorNotEqual uintptr
	pg_QSocketDescriptor_delete func(uintptr)
)

func init() {
	symaddr(&addr_QSocketDescriptor_new, "QSocketDescriptor_new")
	symaddr(&addr_QSocketDescriptor_new2, "QSocketDescriptor_new2")
	symaddr(&addr_QSocketDescriptor_ToQintptr, "QSocketDescriptor_ToQintptr")
	symaddr(&addr_QSocketDescriptor_isValid, "QSocketDescriptor_isValid")
	symaddr(&addr_QSocketDescriptor_OperatorEqual, "QSocketDescriptor_OperatorEqual")
	symaddr(&addr_QSocketDescriptor_OperatorNotEqual, "QSocketDescriptor_OperatorNotEqual")
	reg(&pg_QSocketDescriptor_delete, "QSocketDescriptor_delete")
}

type QSocketDescriptor struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QSocketDescriptor) UnsafeGoID() int64     { return this.vid }
func (this *QSocketDescriptor) UnsafeSetGoID(id int64) { this.vid = id }

func bindQSocketDescriptor(h uintptr) *QSocketDescriptor {
	if h == 0 {
		return nil
	}
	return &QSocketDescriptor{h: h}
}

func newownedQSocketDescriptor(h uintptr) *QSocketDescriptor {
	o := bindQSocketDescriptor(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QSocketDescriptor).Delete)
	}
	return o
}

func (this *QSocketDescriptor) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QSocketDescriptor) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQSocketDescriptor(h unsafe.Pointer) *QSocketDescriptor { return bindQSocketDescriptor(uintptr(h)) }
func UnsafeNewQSocketDescriptorOwned(h unsafe.Pointer) *QSocketDescriptor { return newownedQSocketDescriptor(uintptr(h)) }

func (this *QSocketDescriptor) UnsafeRelease() { this.h = 0 }

func NewQSocketDescriptor() *QSocketDescriptor {
	__r, _, _ := purego.SyscallN(addr_QSocketDescriptor_new)
	__o := newownedQSocketDescriptor(__r)
	return __o
}

func NewQSocketDescriptor2(desc uintptr) *QSocketDescriptor {
	__r, _, _ := purego.SyscallN(addr_QSocketDescriptor_new2, uintptr(desc))
	__o := newownedQSocketDescriptor(__r)
	return __o
}

func (this *QSocketDescriptor) ToQintptr() uintptr {
	__r0, _, _ := purego.SyscallN(addr_QSocketDescriptor_ToQintptr, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return uintptr(__r)
}

func (this *QSocketDescriptor) IsValid() bool {
	__r0, _, _ := purego.SyscallN(addr_QSocketDescriptor_isValid, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QSocketDescriptor) OperatorEqual(rhs *QSocketDescriptor) bool {
	__r0, _, _ := purego.SyscallN(addr_QSocketDescriptor_OperatorEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QSocketDescriptor) OperatorNotEqual(rhs *QSocketDescriptor) bool {
	__r0, _, _ := purego.SyscallN(addr_QSocketDescriptor_OperatorNotEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QSocketDescriptor) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QSocketDescriptor_delete(this.h)
	}
	this.UnsafeRelease()
}
