package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QReadLocker_new uintptr
	addr_QReadLocker_unlock uintptr
	addr_QReadLocker_relock uintptr
	addr_QReadLocker_readWriteLock uintptr
	pg_QReadLocker_delete func(uintptr)
)

func init() {
	symaddr(&addr_QReadLocker_new, "QReadLocker_new")
	symaddr(&addr_QReadLocker_unlock, "QReadLocker_unlock")
	symaddr(&addr_QReadLocker_relock, "QReadLocker_relock")
	symaddr(&addr_QReadLocker_readWriteLock, "QReadLocker_readWriteLock")
	reg(&pg_QReadLocker_delete, "QReadLocker_delete")
}

type QReadLocker struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QReadLocker) UnsafeGoID() int64     { return this.vid }
func (this *QReadLocker) UnsafeSetGoID(id int64) { this.vid = id }

func bindQReadLocker(h uintptr) *QReadLocker {
	if h == 0 {
		return nil
	}
	return &QReadLocker{h: h}
}

func newownedQReadLocker(h uintptr) *QReadLocker {
	o := bindQReadLocker(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QReadLocker).Delete)
	}
	return o
}

func (this *QReadLocker) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QReadLocker) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQReadLocker(h unsafe.Pointer) *QReadLocker { return bindQReadLocker(uintptr(h)) }
func UnsafeNewQReadLockerOwned(h unsafe.Pointer) *QReadLocker { return newownedQReadLocker(uintptr(h)) }

func (this *QReadLocker) UnsafeRelease() { this.h = 0 }

func NewQReadLocker(readWriteLock *QReadWriteLock) *QReadLocker {
	__r, _, _ := purego.SyscallN(addr_QReadLocker_new, uintptr(readWriteLock.UnsafePointer()))
	__o := newownedQReadLocker(__r)
	runtime.KeepAlive(readWriteLock)
	return __o
}

func (this *QReadLocker) Unlock() {
	purego.SyscallN(addr_QReadLocker_unlock, this.h)
	runtime.KeepAlive(this)
}

func (this *QReadLocker) Relock() {
	purego.SyscallN(addr_QReadLocker_relock, this.h)
	runtime.KeepAlive(this)
}

func (this *QReadLocker) ReadWriteLock() *QReadWriteLock {
	__r, _, _ := purego.SyscallN(addr_QReadLocker_readWriteLock, this.h)
	runtime.KeepAlive(this)
	return bindQReadWriteLock(__r)
}

func (this *QReadLocker) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QReadLocker_delete(this.h)
	}
	this.UnsafeRelease()
}
