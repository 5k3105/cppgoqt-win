package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QMutex_base func(uintptr) uintptr
	addr_QMutex_new uintptr
	addr_QMutex_try_lock uintptr
	addr_QMutex_tryLock uintptr
	addr_QMutex_tryLock2 uintptr
	pg_QMutex_delete func(uintptr)
)

func init() {
	reg(&pg_QMutex_base, "QMutex_base")
	symaddr(&addr_QMutex_new, "QMutex_new")
	symaddr(&addr_QMutex_try_lock, "QMutex_try_lock")
	symaddr(&addr_QMutex_tryLock, "QMutex_tryLock")
	symaddr(&addr_QMutex_tryLock2, "QMutex_tryLock2")
	reg(&pg_QMutex_delete, "QMutex_delete")
}

type QMutex struct {
	h uintptr
	*QBasicMutex
}

func bindQMutex(h uintptr) *QMutex {
	if h == 0 {
		return nil
	}
	return &QMutex{h: h, QBasicMutex: bindQBasicMutex(pg_QMutex_base(h))}
}

func newownedQMutex(h uintptr) *QMutex {
	o := bindQMutex(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QMutex).Delete)
	}
	return o
}

func (this *QMutex) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QMutex) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQMutex(h unsafe.Pointer) *QMutex { return bindQMutex(uintptr(h)) }
func UnsafeNewQMutexOwned(h unsafe.Pointer) *QMutex { return newownedQMutex(uintptr(h)) }

func (this *QMutex) UnsafeRelease() {
	this.h = 0
	if this.QBasicMutex != nil {
		this.QBasicMutex.UnsafeRelease()
	}
}

func NewQMutex() *QMutex {
	__r, _, _ := purego.SyscallN(addr_QMutex_new)
	__o := newownedQMutex(__r)
	return __o
}

func (this *QMutex) Try_lock() bool {
	__r0, _, _ := purego.SyscallN(addr_QMutex_try_lock, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMutex) TryLock(timeout int) bool {
	__r0, _, _ := purego.SyscallN(addr_QMutex_tryLock, this.h, uintptr(timeout))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMutex) TryLockWithTimeout(timeout *QDeadlineTimer) bool {
	__r0, _, _ := purego.SyscallN(addr_QMutex_tryLock2, this.h, uintptr(timeout.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(timeout)
	return __r != 0
}

func (this *QMutex) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QMutex_delete(this.h)
	}
	this.UnsafeRelease()
}
