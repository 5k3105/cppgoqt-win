package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QRecursiveMutex_new uintptr
	addr_QRecursiveMutex_lock uintptr
	addr_QRecursiveMutex_tryLock uintptr
	addr_QRecursiveMutex_tryLock2 uintptr
	addr_QRecursiveMutex_unlock uintptr
	addr_QRecursiveMutex_try_lock uintptr
	addr_QRecursiveMutex_tryLock3 uintptr
	pg_QRecursiveMutex_delete func(uintptr)
)

func init() {
	symaddr(&addr_QRecursiveMutex_new, "QRecursiveMutex_new")
	symaddr(&addr_QRecursiveMutex_lock, "QRecursiveMutex_lock")
	symaddr(&addr_QRecursiveMutex_tryLock, "QRecursiveMutex_tryLock")
	symaddr(&addr_QRecursiveMutex_tryLock2, "QRecursiveMutex_tryLock2")
	symaddr(&addr_QRecursiveMutex_unlock, "QRecursiveMutex_unlock")
	symaddr(&addr_QRecursiveMutex_try_lock, "QRecursiveMutex_try_lock")
	symaddr(&addr_QRecursiveMutex_tryLock3, "QRecursiveMutex_tryLock3")
	reg(&pg_QRecursiveMutex_delete, "QRecursiveMutex_delete")
}

type QRecursiveMutex struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QRecursiveMutex) UnsafeGoID() int64     { return this.vid }
func (this *QRecursiveMutex) UnsafeSetGoID(id int64) { this.vid = id }

func bindQRecursiveMutex(h uintptr) *QRecursiveMutex {
	if h == 0 {
		return nil
	}
	return &QRecursiveMutex{h: h}
}

func newownedQRecursiveMutex(h uintptr) *QRecursiveMutex {
	o := bindQRecursiveMutex(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QRecursiveMutex).Delete)
	}
	return o
}

func (this *QRecursiveMutex) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QRecursiveMutex) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQRecursiveMutex(h unsafe.Pointer) *QRecursiveMutex { return bindQRecursiveMutex(uintptr(h)) }
func UnsafeNewQRecursiveMutexOwned(h unsafe.Pointer) *QRecursiveMutex { return newownedQRecursiveMutex(uintptr(h)) }

func (this *QRecursiveMutex) UnsafeRelease() { this.h = 0 }

func NewQRecursiveMutex() *QRecursiveMutex {
	__r, _, _ := purego.SyscallN(addr_QRecursiveMutex_new)
	__o := newownedQRecursiveMutex(__r)
	return __o
}

func (this *QRecursiveMutex) Lock() {
	purego.SyscallN(addr_QRecursiveMutex_lock, this.h)
	runtime.KeepAlive(this)
}

func (this *QRecursiveMutex) TryLock(timeout int) bool {
	__r0, _, _ := purego.SyscallN(addr_QRecursiveMutex_tryLock, this.h, uintptr(timeout))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QRecursiveMutex) TryLock2() bool {
	__r0, _, _ := purego.SyscallN(addr_QRecursiveMutex_tryLock2, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QRecursiveMutex) Unlock() {
	purego.SyscallN(addr_QRecursiveMutex_unlock, this.h)
	runtime.KeepAlive(this)
}

func (this *QRecursiveMutex) Try_lock() bool {
	__r0, _, _ := purego.SyscallN(addr_QRecursiveMutex_try_lock, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QRecursiveMutex) TryLockWithTimer(timer *QDeadlineTimer) bool {
	__r0, _, _ := purego.SyscallN(addr_QRecursiveMutex_tryLock3, this.h, uintptr(timer.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(timer)
	return __r != 0
}

func (this *QRecursiveMutex) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QRecursiveMutex_delete(this.h)
	}
	this.UnsafeRelease()
}
