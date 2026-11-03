package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QPropertyNotifier_chain uintptr
	addr_QPropertyNotifier_new uintptr
	pg_QPropertyNotifier_delete func(uintptr)
)

func init() {
	symaddr(&addr_QPropertyNotifier_chain, "QPropertyNotifier_chain")
	symaddr(&addr_QPropertyNotifier_new, "QPropertyNotifier_new")
	reg(&pg_QPropertyNotifier_delete, "QPropertyNotifier_delete")
}

type QPropertyNotifier struct {
	h uintptr
	*QPropertyObserver
}

func bindQPropertyNotifier(h uintptr) *QPropertyNotifier {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QPropertyNotifier_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QPropertyNotifier{h: h, QPropertyObserver: &QPropertyObserver{h: __c[0], QPropertyObserverBase: &QPropertyObserverBase{h: __c[1]}}}
}

func newownedQPropertyNotifier(h uintptr) *QPropertyNotifier {
	o := bindQPropertyNotifier(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QPropertyNotifier).Delete)
	}
	return o
}

func (this *QPropertyNotifier) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QPropertyNotifier) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQPropertyNotifier(h unsafe.Pointer) *QPropertyNotifier { return bindQPropertyNotifier(uintptr(h)) }
func UnsafeNewQPropertyNotifierOwned(h unsafe.Pointer) *QPropertyNotifier { return newownedQPropertyNotifier(uintptr(h)) }

func (this *QPropertyNotifier) UnsafeRelease() {
	this.h = 0
	if this.QPropertyObserver != nil {
		this.QPropertyObserver.UnsafeRelease()
	}
}

func NewQPropertyNotifier() *QPropertyNotifier {
	__r, _, _ := purego.SyscallN(addr_QPropertyNotifier_new)
	__o := newownedQPropertyNotifier(__r)
	return __o
}

func (this *QPropertyNotifier) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QPropertyNotifier_delete(this.h)
	}
	this.UnsafeRelease()
}
