package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QPropertyObserver_base func(uintptr) uintptr
	addr_QPropertyObserver_new uintptr
	pg_QPropertyObserver_delete func(uintptr)
)

func init() {
	reg(&pg_QPropertyObserver_base, "QPropertyObserver_base")
	symaddr(&addr_QPropertyObserver_new, "QPropertyObserver_new")
	reg(&pg_QPropertyObserver_delete, "QPropertyObserver_delete")
}

type QPropertyObserver struct {
	h uintptr
	*QPropertyObserverBase
}

func bindQPropertyObserver(h uintptr) *QPropertyObserver {
	if h == 0 {
		return nil
	}
	return &QPropertyObserver{h: h, QPropertyObserverBase: bindQPropertyObserverBase(pg_QPropertyObserver_base(h))}
}

func newownedQPropertyObserver(h uintptr) *QPropertyObserver {
	o := bindQPropertyObserver(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QPropertyObserver).Delete)
	}
	return o
}

func (this *QPropertyObserver) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QPropertyObserver) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQPropertyObserver(h unsafe.Pointer) *QPropertyObserver { return bindQPropertyObserver(uintptr(h)) }
func UnsafeNewQPropertyObserverOwned(h unsafe.Pointer) *QPropertyObserver { return newownedQPropertyObserver(uintptr(h)) }

func (this *QPropertyObserver) UnsafeRelease() {
	this.h = 0
	if this.QPropertyObserverBase != nil {
		this.QPropertyObserverBase.UnsafeRelease()
	}
}

func NewQPropertyObserver() *QPropertyObserver {
	__r, _, _ := purego.SyscallN(addr_QPropertyObserver_new)
	__o := newownedQPropertyObserver(__r)
	return __o
}

func (this *QPropertyObserver) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QPropertyObserver_delete(this.h)
	}
	this.UnsafeRelease()
}
