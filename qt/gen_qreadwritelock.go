package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QReadWriteLock_base func(uintptr) uintptr
	addr_QReadWriteLock_new uintptr
	addr_QReadWriteLock_new2 uintptr
	addr_QReadWriteLock_tryLockForRead uintptr
	addr_QReadWriteLock_tryLockForWrite uintptr
	pg_QReadWriteLock_delete func(uintptr)
)

func init() {
	reg(&pg_QReadWriteLock_base, "QReadWriteLock_base")
	symaddr(&addr_QReadWriteLock_new, "QReadWriteLock_new")
	symaddr(&addr_QReadWriteLock_new2, "QReadWriteLock_new2")
	symaddr(&addr_QReadWriteLock_tryLockForRead, "QReadWriteLock_tryLockForRead")
	symaddr(&addr_QReadWriteLock_tryLockForWrite, "QReadWriteLock_tryLockForWrite")
	reg(&pg_QReadWriteLock_delete, "QReadWriteLock_delete")
}

type QReadWriteLock struct {
	h uintptr
	*QBasicReadWriteLock
}

func bindQReadWriteLock(h uintptr) *QReadWriteLock {
	if h == 0 {
		return nil
	}
	return &QReadWriteLock{h: h, QBasicReadWriteLock: bindQBasicReadWriteLock(pg_QReadWriteLock_base(h))}
}

func newownedQReadWriteLock(h uintptr) *QReadWriteLock {
	o := bindQReadWriteLock(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QReadWriteLock).Delete)
	}
	return o
}

func (this *QReadWriteLock) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QReadWriteLock) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQReadWriteLock(h unsafe.Pointer) *QReadWriteLock { return bindQReadWriteLock(uintptr(h)) }
func UnsafeNewQReadWriteLockOwned(h unsafe.Pointer) *QReadWriteLock { return newownedQReadWriteLock(uintptr(h)) }

func (this *QReadWriteLock) UnsafeRelease() {
	this.h = 0
	if this.QBasicReadWriteLock != nil {
		this.QBasicReadWriteLock.UnsafeRelease()
	}
}

func NewQReadWriteLock() *QReadWriteLock {
	__r, _, _ := purego.SyscallN(addr_QReadWriteLock_new)
	__o := newownedQReadWriteLock(__r)
	return __o
}

func NewQReadWriteLock2(recursionMode QReadWriteLock__RecursionMode) *QReadWriteLock {
	__r, _, _ := purego.SyscallN(addr_QReadWriteLock_new2, uintptr(recursionMode))
	__o := newownedQReadWriteLock(__r)
	return __o
}

func (this *QReadWriteLock) TryLockForRead(timeout int) bool {
	__r0, _, _ := purego.SyscallN(addr_QReadWriteLock_tryLockForRead, this.h, uintptr(timeout))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QReadWriteLock) TryLockForWrite(timeout int) bool {
	__r0, _, _ := purego.SyscallN(addr_QReadWriteLock_tryLockForWrite, this.h, uintptr(timeout))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QReadWriteLock) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QReadWriteLock_delete(this.h)
	}
	this.UnsafeRelease()
}
