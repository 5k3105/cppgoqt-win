package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QWriteLocker_new uintptr
	addr_QWriteLocker_unlock uintptr
	addr_QWriteLocker_relock uintptr
	addr_QWriteLocker_readWriteLock uintptr
	pg_QWriteLocker_delete func(uintptr)
)

func init() {
	symaddr(&addr_QWriteLocker_new, "QWriteLocker_new")
	symaddr(&addr_QWriteLocker_unlock, "QWriteLocker_unlock")
	symaddr(&addr_QWriteLocker_relock, "QWriteLocker_relock")
	symaddr(&addr_QWriteLocker_readWriteLock, "QWriteLocker_readWriteLock")
	reg(&pg_QWriteLocker_delete, "QWriteLocker_delete")
}

type QWriteLocker struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QWriteLocker) UnsafeGoID() int64     { return this.vid }
func (this *QWriteLocker) UnsafeSetGoID(id int64) { this.vid = id }

func bindQWriteLocker(h uintptr) *QWriteLocker {
	if h == 0 {
		return nil
	}
	return &QWriteLocker{h: h}
}

func newownedQWriteLocker(h uintptr) *QWriteLocker {
	o := bindQWriteLocker(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QWriteLocker).Delete)
	}
	return o
}

func (this *QWriteLocker) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QWriteLocker) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQWriteLocker(h unsafe.Pointer) *QWriteLocker { return bindQWriteLocker(uintptr(h)) }
func UnsafeNewQWriteLockerOwned(h unsafe.Pointer) *QWriteLocker { return newownedQWriteLocker(uintptr(h)) }

func (this *QWriteLocker) UnsafeRelease() { this.h = 0 }

func NewQWriteLocker(readWriteLock *QReadWriteLock) *QWriteLocker {
	__r, _, _ := purego.SyscallN(addr_QWriteLocker_new, uintptr(readWriteLock.UnsafePointer()))
	__o := newownedQWriteLocker(__r)
	runtime.KeepAlive(readWriteLock)
	return __o
}

func (this *QWriteLocker) Unlock() {
	purego.SyscallN(addr_QWriteLocker_unlock, this.h)
	runtime.KeepAlive(this)
}

func (this *QWriteLocker) Relock() {
	purego.SyscallN(addr_QWriteLocker_relock, this.h)
	runtime.KeepAlive(this)
}

func (this *QWriteLocker) ReadWriteLock() *QReadWriteLock {
	__r, _, _ := purego.SyscallN(addr_QWriteLocker_readWriteLock, this.h)
	runtime.KeepAlive(this)
	return bindQReadWriteLock(__r)
}

func (this *QWriteLocker) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QWriteLocker_delete(this.h)
	}
	this.UnsafeRelease()
}
