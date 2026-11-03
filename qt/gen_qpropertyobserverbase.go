package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QPropertyObserverBase_new uintptr
	addr_QPropertyObserverBase_new2 uintptr
	addr_QPropertyObserverBase_OperatorAssign uintptr
	pg_QPropertyObserverBase_delete func(uintptr)
)

func init() {
	symaddr(&addr_QPropertyObserverBase_new, "QPropertyObserverBase_new")
	symaddr(&addr_QPropertyObserverBase_new2, "QPropertyObserverBase_new2")
	symaddr(&addr_QPropertyObserverBase_OperatorAssign, "QPropertyObserverBase_OperatorAssign")
	reg(&pg_QPropertyObserverBase_delete, "QPropertyObserverBase_delete")
}

type QPropertyObserverBase struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QPropertyObserverBase) UnsafeGoID() int64     { return this.vid }
func (this *QPropertyObserverBase) UnsafeSetGoID(id int64) { this.vid = id }

func bindQPropertyObserverBase(h uintptr) *QPropertyObserverBase {
	if h == 0 {
		return nil
	}
	return &QPropertyObserverBase{h: h}
}

func newownedQPropertyObserverBase(h uintptr) *QPropertyObserverBase {
	o := bindQPropertyObserverBase(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QPropertyObserverBase).Delete)
	}
	return o
}

func (this *QPropertyObserverBase) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QPropertyObserverBase) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQPropertyObserverBase(h unsafe.Pointer) *QPropertyObserverBase { return bindQPropertyObserverBase(uintptr(h)) }
func UnsafeNewQPropertyObserverBaseOwned(h unsafe.Pointer) *QPropertyObserverBase { return newownedQPropertyObserverBase(uintptr(h)) }

func (this *QPropertyObserverBase) UnsafeRelease() { this.h = 0 }

func NewQPropertyObserverBase() *QPropertyObserverBase {
	__r, _, _ := purego.SyscallN(addr_QPropertyObserverBase_new)
	__o := newownedQPropertyObserverBase(__r)
	return __o
}

func NewQPropertyObserverBase2(arg1 *QPropertyObserverBase) *QPropertyObserverBase {
	__r, _, _ := purego.SyscallN(addr_QPropertyObserverBase_new2, uintptr(arg1.UnsafePointer()))
	__o := newownedQPropertyObserverBase(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QPropertyObserverBase) OperatorAssign(arg1 *QPropertyObserverBase) *QPropertyObserverBase {
	__r, _, _ := purego.SyscallN(addr_QPropertyObserverBase_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQPropertyObserverBase(__r)
}

func (this *QPropertyObserverBase) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QPropertyObserverBase_delete(this.h)
	}
	this.UnsafeRelease()
}
