package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QBasicReadWriteLock_new uintptr
	addr_QBasicReadWriteLock_lockForRead uintptr
	addr_QBasicReadWriteLock_tryLockForRead uintptr
	addr_QBasicReadWriteLock_tryLockForRead2 uintptr
	addr_QBasicReadWriteLock_lockForWrite uintptr
	addr_QBasicReadWriteLock_tryLockForWrite uintptr
	addr_QBasicReadWriteLock_tryLockForWrite2 uintptr
	addr_QBasicReadWriteLock_unlock uintptr
	addr_QBasicReadWriteLock_lock uintptr
	addr_QBasicReadWriteLock_lock_shared uintptr
	addr_QBasicReadWriteLock_try_lock uintptr
	addr_QBasicReadWriteLock_try_lock_shared uintptr
	addr_QBasicReadWriteLock_unlock_shared uintptr
	pg_QBasicReadWriteLock_delete func(uintptr)
)

func init() {
	symaddr(&addr_QBasicReadWriteLock_new, "QBasicReadWriteLock_new")
	symaddr(&addr_QBasicReadWriteLock_lockForRead, "QBasicReadWriteLock_lockForRead")
	symaddr(&addr_QBasicReadWriteLock_tryLockForRead, "QBasicReadWriteLock_tryLockForRead")
	symaddr(&addr_QBasicReadWriteLock_tryLockForRead2, "QBasicReadWriteLock_tryLockForRead2")
	symaddr(&addr_QBasicReadWriteLock_lockForWrite, "QBasicReadWriteLock_lockForWrite")
	symaddr(&addr_QBasicReadWriteLock_tryLockForWrite, "QBasicReadWriteLock_tryLockForWrite")
	symaddr(&addr_QBasicReadWriteLock_tryLockForWrite2, "QBasicReadWriteLock_tryLockForWrite2")
	symaddr(&addr_QBasicReadWriteLock_unlock, "QBasicReadWriteLock_unlock")
	symaddr(&addr_QBasicReadWriteLock_lock, "QBasicReadWriteLock_lock")
	symaddr(&addr_QBasicReadWriteLock_lock_shared, "QBasicReadWriteLock_lock_shared")
	symaddr(&addr_QBasicReadWriteLock_try_lock, "QBasicReadWriteLock_try_lock")
	symaddr(&addr_QBasicReadWriteLock_try_lock_shared, "QBasicReadWriteLock_try_lock_shared")
	symaddr(&addr_QBasicReadWriteLock_unlock_shared, "QBasicReadWriteLock_unlock_shared")
	reg(&pg_QBasicReadWriteLock_delete, "QBasicReadWriteLock_delete")
}

type QBasicReadWriteLock struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QBasicReadWriteLock) UnsafeGoID() int64     { return this.vid }
func (this *QBasicReadWriteLock) UnsafeSetGoID(id int64) { this.vid = id }

func bindQBasicReadWriteLock(h uintptr) *QBasicReadWriteLock {
	if h == 0 {
		return nil
	}
	return &QBasicReadWriteLock{h: h}
}

func newownedQBasicReadWriteLock(h uintptr) *QBasicReadWriteLock {
	o := bindQBasicReadWriteLock(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QBasicReadWriteLock).Delete)
	}
	return o
}

func (this *QBasicReadWriteLock) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QBasicReadWriteLock) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQBasicReadWriteLock(h unsafe.Pointer) *QBasicReadWriteLock { return bindQBasicReadWriteLock(uintptr(h)) }
func UnsafeNewQBasicReadWriteLockOwned(h unsafe.Pointer) *QBasicReadWriteLock { return newownedQBasicReadWriteLock(uintptr(h)) }

func (this *QBasicReadWriteLock) UnsafeRelease() { this.h = 0 }

func NewQBasicReadWriteLock() *QBasicReadWriteLock {
	__r, _, _ := purego.SyscallN(addr_QBasicReadWriteLock_new)
	__o := newownedQBasicReadWriteLock(__r)
	return __o
}

func (this *QBasicReadWriteLock) LockForRead() {
	purego.SyscallN(addr_QBasicReadWriteLock_lockForRead, this.h)
	runtime.KeepAlive(this)
}

func (this *QBasicReadWriteLock) TryLockForRead() bool {
	__r0, _, _ := purego.SyscallN(addr_QBasicReadWriteLock_tryLockForRead, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QBasicReadWriteLock) TryLockForReadWithTimeout(timeout *QDeadlineTimer) bool {
	__r0, _, _ := purego.SyscallN(addr_QBasicReadWriteLock_tryLockForRead2, this.h, uintptr(timeout.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(timeout)
	return __r != 0
}

func (this *QBasicReadWriteLock) LockForWrite() {
	purego.SyscallN(addr_QBasicReadWriteLock_lockForWrite, this.h)
	runtime.KeepAlive(this)
}

func (this *QBasicReadWriteLock) TryLockForWrite() bool {
	__r0, _, _ := purego.SyscallN(addr_QBasicReadWriteLock_tryLockForWrite, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QBasicReadWriteLock) TryLockForWriteWithTimeout(timeout *QDeadlineTimer) bool {
	__r0, _, _ := purego.SyscallN(addr_QBasicReadWriteLock_tryLockForWrite2, this.h, uintptr(timeout.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(timeout)
	return __r != 0
}

func (this *QBasicReadWriteLock) Unlock() {
	purego.SyscallN(addr_QBasicReadWriteLock_unlock, this.h)
	runtime.KeepAlive(this)
}

func (this *QBasicReadWriteLock) Lock() {
	purego.SyscallN(addr_QBasicReadWriteLock_lock, this.h)
	runtime.KeepAlive(this)
}

func (this *QBasicReadWriteLock) Lock_shared() {
	purego.SyscallN(addr_QBasicReadWriteLock_lock_shared, this.h)
	runtime.KeepAlive(this)
}

func (this *QBasicReadWriteLock) Try_lock() bool {
	__r0, _, _ := purego.SyscallN(addr_QBasicReadWriteLock_try_lock, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QBasicReadWriteLock) Try_lock_shared() bool {
	__r0, _, _ := purego.SyscallN(addr_QBasicReadWriteLock_try_lock_shared, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QBasicReadWriteLock) Unlock_shared() {
	purego.SyscallN(addr_QBasicReadWriteLock_unlock_shared, this.h)
	runtime.KeepAlive(this)
}

func (this *QBasicReadWriteLock) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QBasicReadWriteLock_delete(this.h)
	}
	this.UnsafeRelease()
}
