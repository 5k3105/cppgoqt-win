package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QHashDummyValue_new uintptr
	addr_QHashDummyValue_OperatorEqual uintptr
	addr_QHashDummyValue_OperatorNotEqual uintptr
	pg_QHashDummyValue_delete func(uintptr)
)

func init() {
	symaddr(&addr_QHashDummyValue_new, "QHashDummyValue_new")
	symaddr(&addr_QHashDummyValue_OperatorEqual, "QHashDummyValue_OperatorEqual")
	symaddr(&addr_QHashDummyValue_OperatorNotEqual, "QHashDummyValue_OperatorNotEqual")
	reg(&pg_QHashDummyValue_delete, "QHashDummyValue_delete")
}

type QHashDummyValue struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QHashDummyValue) UnsafeGoID() int64     { return this.vid }
func (this *QHashDummyValue) UnsafeSetGoID(id int64) { this.vid = id }

func bindQHashDummyValue(h uintptr) *QHashDummyValue {
	if h == 0 {
		return nil
	}
	return &QHashDummyValue{h: h}
}

func newownedQHashDummyValue(h uintptr) *QHashDummyValue {
	o := bindQHashDummyValue(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QHashDummyValue).Delete)
	}
	return o
}

func (this *QHashDummyValue) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QHashDummyValue) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQHashDummyValue(h unsafe.Pointer) *QHashDummyValue { return bindQHashDummyValue(uintptr(h)) }
func UnsafeNewQHashDummyValueOwned(h unsafe.Pointer) *QHashDummyValue { return newownedQHashDummyValue(uintptr(h)) }

func (this *QHashDummyValue) UnsafeRelease() { this.h = 0 }

func NewQHashDummyValue() *QHashDummyValue {
	__r, _, _ := purego.SyscallN(addr_QHashDummyValue_new)
	__o := newownedQHashDummyValue(__r)
	return __o
}

func (this *QHashDummyValue) OperatorEqual(arg1 *QHashDummyValue) bool {
	__r0, _, _ := purego.SyscallN(addr_QHashDummyValue_OperatorEqual, this.h, uintptr(arg1.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return __r != 0
}

func (this *QHashDummyValue) OperatorNotEqual(arg1 *QHashDummyValue) bool {
	__r0, _, _ := purego.SyscallN(addr_QHashDummyValue_OperatorNotEqual, this.h, uintptr(arg1.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return __r != 0
}

func (this *QHashDummyValue) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QHashDummyValue_delete(this.h)
	}
	this.UnsafeRelease()
}
