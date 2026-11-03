package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QUnhandledException_base func(uintptr) uintptr
	addr_QUnhandledException_new uintptr
	addr_QUnhandledException_new2 uintptr
	addr_QUnhandledException_swap uintptr
	addr_QUnhandledException_OperatorAssign uintptr
	pg_QUnhandledException_delete func(uintptr)
)

func init() {
	reg(&pg_QUnhandledException_base, "QUnhandledException_base")
	symaddr(&addr_QUnhandledException_new, "QUnhandledException_new")
	symaddr(&addr_QUnhandledException_new2, "QUnhandledException_new2")
	symaddr(&addr_QUnhandledException_swap, "QUnhandledException_swap")
	symaddr(&addr_QUnhandledException_OperatorAssign, "QUnhandledException_OperatorAssign")
	reg(&pg_QUnhandledException_delete, "QUnhandledException_delete")
}

type QUnhandledException struct {
	h uintptr
	*QException
}

func bindQUnhandledException(h uintptr) *QUnhandledException {
	if h == 0 {
		return nil
	}
	return &QUnhandledException{h: h, QException: bindQException(pg_QUnhandledException_base(h))}
}

func newownedQUnhandledException(h uintptr) *QUnhandledException {
	o := bindQUnhandledException(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QUnhandledException).Delete)
	}
	return o
}

func (this *QUnhandledException) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QUnhandledException) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQUnhandledException(h unsafe.Pointer) *QUnhandledException { return bindQUnhandledException(uintptr(h)) }
func UnsafeNewQUnhandledExceptionOwned(h unsafe.Pointer) *QUnhandledException { return newownedQUnhandledException(uintptr(h)) }

func (this *QUnhandledException) UnsafeRelease() {
	this.h = 0
	if this.QException != nil {
		this.QException.UnsafeRelease()
	}
}

func NewQUnhandledException() *QUnhandledException {
	__r, _, _ := purego.SyscallN(addr_QUnhandledException_new)
	__o := newownedQUnhandledException(__r)
	return __o
}

func NewQUnhandledException2(other *QUnhandledException) *QUnhandledException {
	__r, _, _ := purego.SyscallN(addr_QUnhandledException_new2, uintptr(other.UnsafePointer()))
	__o := newownedQUnhandledException(__r)
	runtime.KeepAlive(other)
	return __o
}

func (this *QUnhandledException) Swap(other *QUnhandledException) {
	purego.SyscallN(addr_QUnhandledException_swap, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
}

func (this *QUnhandledException) OperatorAssign(other *QUnhandledException) *QUnhandledException {
	__r, _, _ := purego.SyscallN(addr_QUnhandledException_OperatorAssign, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return bindQUnhandledException(__r)
}

func (this *QUnhandledException) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QUnhandledException_delete(this.h)
	}
	this.UnsafeRelease()
}
