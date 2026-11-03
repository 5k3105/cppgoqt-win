package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QSignalBlocker_new uintptr
	addr_QSignalBlocker_new2 uintptr
	addr_QSignalBlocker_reblock uintptr
	addr_QSignalBlocker_unblock uintptr
	addr_QSignalBlocker_dismiss uintptr
	pg_QSignalBlocker_delete func(uintptr)
)

func init() {
	symaddr(&addr_QSignalBlocker_new, "QSignalBlocker_new")
	symaddr(&addr_QSignalBlocker_new2, "QSignalBlocker_new2")
	symaddr(&addr_QSignalBlocker_reblock, "QSignalBlocker_reblock")
	symaddr(&addr_QSignalBlocker_unblock, "QSignalBlocker_unblock")
	symaddr(&addr_QSignalBlocker_dismiss, "QSignalBlocker_dismiss")
	reg(&pg_QSignalBlocker_delete, "QSignalBlocker_delete")
}

type QSignalBlocker struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QSignalBlocker) UnsafeGoID() int64     { return this.vid }
func (this *QSignalBlocker) UnsafeSetGoID(id int64) { this.vid = id }

func bindQSignalBlocker(h uintptr) *QSignalBlocker {
	if h == 0 {
		return nil
	}
	return &QSignalBlocker{h: h}
}

func newownedQSignalBlocker(h uintptr) *QSignalBlocker {
	o := bindQSignalBlocker(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QSignalBlocker).Delete)
	}
	return o
}

func (this *QSignalBlocker) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QSignalBlocker) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQSignalBlocker(h unsafe.Pointer) *QSignalBlocker { return bindQSignalBlocker(uintptr(h)) }
func UnsafeNewQSignalBlockerOwned(h unsafe.Pointer) *QSignalBlocker { return newownedQSignalBlocker(uintptr(h)) }

func (this *QSignalBlocker) UnsafeRelease() { this.h = 0 }

func NewQSignalBlocker(o *QObject) *QSignalBlocker {
	__r, _, _ := purego.SyscallN(addr_QSignalBlocker_new, uintptr(o.UnsafePointer()))
	__o := newownedQSignalBlocker(__r)
	runtime.KeepAlive(o)
	return __o
}

func NewQSignalBlocker2(o *QObject) *QSignalBlocker {
	__r, _, _ := purego.SyscallN(addr_QSignalBlocker_new2, uintptr(o.UnsafePointer()))
	__o := newownedQSignalBlocker(__r)
	runtime.KeepAlive(o)
	return __o
}

func (this *QSignalBlocker) Reblock() {
	purego.SyscallN(addr_QSignalBlocker_reblock, this.h)
	runtime.KeepAlive(this)
}

func (this *QSignalBlocker) Unblock() {
	purego.SyscallN(addr_QSignalBlocker_unblock, this.h)
	runtime.KeepAlive(this)
}

func (this *QSignalBlocker) Dismiss() {
	purego.SyscallN(addr_QSignalBlocker_dismiss, this.h)
	runtime.KeepAlive(this)
}

func (this *QSignalBlocker) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QSignalBlocker_delete(this.h)
	}
	this.UnsafeRelease()
}
