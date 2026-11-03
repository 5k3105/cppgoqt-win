package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QEventLoopLocker_new uintptr
	addr_QEventLoopLocker_new2 uintptr
	addr_QEventLoopLocker_new3 uintptr
	addr_QEventLoopLocker_swap uintptr
	pg_QEventLoopLocker_delete func(uintptr)
)

func init() {
	symaddr(&addr_QEventLoopLocker_new, "QEventLoopLocker_new")
	symaddr(&addr_QEventLoopLocker_new2, "QEventLoopLocker_new2")
	symaddr(&addr_QEventLoopLocker_new3, "QEventLoopLocker_new3")
	symaddr(&addr_QEventLoopLocker_swap, "QEventLoopLocker_swap")
	reg(&pg_QEventLoopLocker_delete, "QEventLoopLocker_delete")
}

type QEventLoopLocker struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QEventLoopLocker) UnsafeGoID() int64     { return this.vid }
func (this *QEventLoopLocker) UnsafeSetGoID(id int64) { this.vid = id }

func bindQEventLoopLocker(h uintptr) *QEventLoopLocker {
	if h == 0 {
		return nil
	}
	return &QEventLoopLocker{h: h}
}

func newownedQEventLoopLocker(h uintptr) *QEventLoopLocker {
	o := bindQEventLoopLocker(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QEventLoopLocker).Delete)
	}
	return o
}

func (this *QEventLoopLocker) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QEventLoopLocker) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQEventLoopLocker(h unsafe.Pointer) *QEventLoopLocker { return bindQEventLoopLocker(uintptr(h)) }
func UnsafeNewQEventLoopLockerOwned(h unsafe.Pointer) *QEventLoopLocker { return newownedQEventLoopLocker(uintptr(h)) }

func (this *QEventLoopLocker) UnsafeRelease() { this.h = 0 }

func NewQEventLoopLocker() *QEventLoopLocker {
	__r, _, _ := purego.SyscallN(addr_QEventLoopLocker_new)
	__o := newownedQEventLoopLocker(__r)
	return __o
}

func NewQEventLoopLocker2(loop *QEventLoop) *QEventLoopLocker {
	__r, _, _ := purego.SyscallN(addr_QEventLoopLocker_new2, uintptr(loop.UnsafePointer()))
	__o := newownedQEventLoopLocker(__r)
	runtime.KeepAlive(loop)
	return __o
}

func NewQEventLoopLocker3(thread *QThread) *QEventLoopLocker {
	__r, _, _ := purego.SyscallN(addr_QEventLoopLocker_new3, uintptr(thread.UnsafePointer()))
	__o := newownedQEventLoopLocker(__r)
	runtime.KeepAlive(thread)
	return __o
}

func (this *QEventLoopLocker) Swap(other *QEventLoopLocker) {
	purego.SyscallN(addr_QEventLoopLocker_swap, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
}

func (this *QEventLoopLocker) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QEventLoopLocker_delete(this.h)
	}
	this.UnsafeRelease()
}
