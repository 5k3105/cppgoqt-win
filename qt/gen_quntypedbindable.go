package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QUntypedBindable_new uintptr
	addr_QUntypedBindable_new2 uintptr
	addr_QUntypedBindable_isValid uintptr
	addr_QUntypedBindable_isBindable uintptr
	addr_QUntypedBindable_isReadOnly uintptr
	addr_QUntypedBindable_makeBinding uintptr
	addr_QUntypedBindable_takeBinding uintptr
	addr_QUntypedBindable_observe uintptr
	addr_QUntypedBindable_binding uintptr
	addr_QUntypedBindable_setBinding uintptr
	addr_QUntypedBindable_hasBinding uintptr
	addr_QUntypedBindable_metaType uintptr
	addr_QUntypedBindable_makeBinding2 uintptr
	pg_QUntypedBindable_delete func(uintptr)
)

func init() {
	symaddr(&addr_QUntypedBindable_new, "QUntypedBindable_new")
	symaddr(&addr_QUntypedBindable_new2, "QUntypedBindable_new2")
	symaddr(&addr_QUntypedBindable_isValid, "QUntypedBindable_isValid")
	symaddr(&addr_QUntypedBindable_isBindable, "QUntypedBindable_isBindable")
	symaddr(&addr_QUntypedBindable_isReadOnly, "QUntypedBindable_isReadOnly")
	symaddr(&addr_QUntypedBindable_makeBinding, "QUntypedBindable_makeBinding")
	symaddr(&addr_QUntypedBindable_takeBinding, "QUntypedBindable_takeBinding")
	symaddr(&addr_QUntypedBindable_observe, "QUntypedBindable_observe")
	symaddr(&addr_QUntypedBindable_binding, "QUntypedBindable_binding")
	symaddr(&addr_QUntypedBindable_setBinding, "QUntypedBindable_setBinding")
	symaddr(&addr_QUntypedBindable_hasBinding, "QUntypedBindable_hasBinding")
	symaddr(&addr_QUntypedBindable_metaType, "QUntypedBindable_metaType")
	symaddr(&addr_QUntypedBindable_makeBinding2, "QUntypedBindable_makeBinding2")
	reg(&pg_QUntypedBindable_delete, "QUntypedBindable_delete")
}

type QUntypedBindable struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QUntypedBindable) UnsafeGoID() int64     { return this.vid }
func (this *QUntypedBindable) UnsafeSetGoID(id int64) { this.vid = id }

func bindQUntypedBindable(h uintptr) *QUntypedBindable {
	if h == 0 {
		return nil
	}
	return &QUntypedBindable{h: h}
}

func newownedQUntypedBindable(h uintptr) *QUntypedBindable {
	o := bindQUntypedBindable(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QUntypedBindable).Delete)
	}
	return o
}

func (this *QUntypedBindable) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QUntypedBindable) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQUntypedBindable(h unsafe.Pointer) *QUntypedBindable { return bindQUntypedBindable(uintptr(h)) }
func UnsafeNewQUntypedBindableOwned(h unsafe.Pointer) *QUntypedBindable { return newownedQUntypedBindable(uintptr(h)) }

func (this *QUntypedBindable) UnsafeRelease() { this.h = 0 }

func NewQUntypedBindable() *QUntypedBindable {
	__r, _, _ := purego.SyscallN(addr_QUntypedBindable_new)
	__o := newownedQUntypedBindable(__r)
	return __o
}

func NewQUntypedBindable2(arg1 *QUntypedBindable) *QUntypedBindable {
	__r, _, _ := purego.SyscallN(addr_QUntypedBindable_new2, uintptr(arg1.UnsafePointer()))
	__o := newownedQUntypedBindable(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QUntypedBindable) IsValid() bool {
	__r0, _, _ := purego.SyscallN(addr_QUntypedBindable_isValid, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QUntypedBindable) IsBindable() bool {
	__r0, _, _ := purego.SyscallN(addr_QUntypedBindable_isBindable, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QUntypedBindable) IsReadOnly() bool {
	__r0, _, _ := purego.SyscallN(addr_QUntypedBindable_isReadOnly, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QUntypedBindable) MakeBinding() *QUntypedPropertyBinding {
	__r, _, _ := purego.SyscallN(addr_QUntypedBindable_makeBinding, this.h)
	runtime.KeepAlive(this)
	return newownedQUntypedPropertyBinding(__r)
}

func (this *QUntypedBindable) TakeBinding() *QUntypedPropertyBinding {
	__r, _, _ := purego.SyscallN(addr_QUntypedBindable_takeBinding, this.h)
	runtime.KeepAlive(this)
	return newownedQUntypedPropertyBinding(__r)
}

func (this *QUntypedBindable) Observe(observer *QPropertyObserver) {
	purego.SyscallN(addr_QUntypedBindable_observe, this.h, uintptr(observer.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(observer)
}

func (this *QUntypedBindable) Binding() *QUntypedPropertyBinding {
	__r, _, _ := purego.SyscallN(addr_QUntypedBindable_binding, this.h)
	runtime.KeepAlive(this)
	return newownedQUntypedPropertyBinding(__r)
}

func (this *QUntypedBindable) SetBinding(binding *QUntypedPropertyBinding) bool {
	__r0, _, _ := purego.SyscallN(addr_QUntypedBindable_setBinding, this.h, uintptr(binding.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(binding)
	return __r != 0
}

func (this *QUntypedBindable) HasBinding() bool {
	__r0, _, _ := purego.SyscallN(addr_QUntypedBindable_hasBinding, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QUntypedBindable) MetaType() *QMetaType {
	__r, _, _ := purego.SyscallN(addr_QUntypedBindable_metaType, this.h)
	runtime.KeepAlive(this)
	return newownedQMetaType(__r)
}

func (this *QUntypedBindable) MakeBindingWithLocation(location *QPropertyBindingSourceLocation) *QUntypedPropertyBinding {
	__r, _, _ := purego.SyscallN(addr_QUntypedBindable_makeBinding2, this.h, uintptr(location.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(location)
	return newownedQUntypedPropertyBinding(__r)
}

func (this *QUntypedBindable) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QUntypedBindable_delete(this.h)
	}
	this.UnsafeRelease()
}
