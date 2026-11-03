package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QUntypedPropertyBinding_new uintptr
	addr_QUntypedPropertyBinding_new2 uintptr
	addr_QUntypedPropertyBinding_OperatorAssign uintptr
	addr_QUntypedPropertyBinding_isNull uintptr
	addr_QUntypedPropertyBinding_error uintptr
	addr_QUntypedPropertyBinding_valueMetaType uintptr
	pg_QUntypedPropertyBinding_delete func(uintptr)
)

func init() {
	symaddr(&addr_QUntypedPropertyBinding_new, "QUntypedPropertyBinding_new")
	symaddr(&addr_QUntypedPropertyBinding_new2, "QUntypedPropertyBinding_new2")
	symaddr(&addr_QUntypedPropertyBinding_OperatorAssign, "QUntypedPropertyBinding_OperatorAssign")
	symaddr(&addr_QUntypedPropertyBinding_isNull, "QUntypedPropertyBinding_isNull")
	symaddr(&addr_QUntypedPropertyBinding_error, "QUntypedPropertyBinding_error")
	symaddr(&addr_QUntypedPropertyBinding_valueMetaType, "QUntypedPropertyBinding_valueMetaType")
	reg(&pg_QUntypedPropertyBinding_delete, "QUntypedPropertyBinding_delete")
}

type QUntypedPropertyBinding struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QUntypedPropertyBinding) UnsafeGoID() int64     { return this.vid }
func (this *QUntypedPropertyBinding) UnsafeSetGoID(id int64) { this.vid = id }

func bindQUntypedPropertyBinding(h uintptr) *QUntypedPropertyBinding {
	if h == 0 {
		return nil
	}
	return &QUntypedPropertyBinding{h: h}
}

func newownedQUntypedPropertyBinding(h uintptr) *QUntypedPropertyBinding {
	o := bindQUntypedPropertyBinding(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QUntypedPropertyBinding).Delete)
	}
	return o
}

func (this *QUntypedPropertyBinding) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QUntypedPropertyBinding) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQUntypedPropertyBinding(h unsafe.Pointer) *QUntypedPropertyBinding { return bindQUntypedPropertyBinding(uintptr(h)) }
func UnsafeNewQUntypedPropertyBindingOwned(h unsafe.Pointer) *QUntypedPropertyBinding { return newownedQUntypedPropertyBinding(uintptr(h)) }

func (this *QUntypedPropertyBinding) UnsafeRelease() { this.h = 0 }

func NewQUntypedPropertyBinding() *QUntypedPropertyBinding {
	__r, _, _ := purego.SyscallN(addr_QUntypedPropertyBinding_new)
	__o := newownedQUntypedPropertyBinding(__r)
	return __o
}

func NewQUntypedPropertyBinding2(other *QUntypedPropertyBinding) *QUntypedPropertyBinding {
	__r, _, _ := purego.SyscallN(addr_QUntypedPropertyBinding_new2, uintptr(other.UnsafePointer()))
	__o := newownedQUntypedPropertyBinding(__r)
	runtime.KeepAlive(other)
	return __o
}

func (this *QUntypedPropertyBinding) OperatorAssign(other *QUntypedPropertyBinding) *QUntypedPropertyBinding {
	__r, _, _ := purego.SyscallN(addr_QUntypedPropertyBinding_OperatorAssign, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return bindQUntypedPropertyBinding(__r)
}

func (this *QUntypedPropertyBinding) IsNull() bool {
	__r0, _, _ := purego.SyscallN(addr_QUntypedPropertyBinding_isNull, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QUntypedPropertyBinding) Error() *QPropertyBindingError {
	__r, _, _ := purego.SyscallN(addr_QUntypedPropertyBinding_error, this.h)
	runtime.KeepAlive(this)
	return newownedQPropertyBindingError(__r)
}

func (this *QUntypedPropertyBinding) ValueMetaType() *QMetaType {
	__r, _, _ := purego.SyscallN(addr_QUntypedPropertyBinding_valueMetaType, this.h)
	runtime.KeepAlive(this)
	return newownedQMetaType(__r)
}

func (this *QUntypedPropertyBinding) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QUntypedPropertyBinding_delete(this.h)
	}
	this.UnsafeRelease()
}
