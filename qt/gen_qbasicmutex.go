package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QBasicMutex_new uintptr
	addr_QBasicMutex_lock uintptr
	addr_QBasicMutex_unlock uintptr
	addr_QBasicMutex_tryLock uintptr
	addr_QBasicMutex_try_lock uintptr
	pg_QBasicMutex_delete func(uintptr)
)

func init() {
	symaddr(&addr_QBasicMutex_new, "QBasicMutex_new")
	symaddr(&addr_QBasicMutex_lock, "QBasicMutex_lock")
	symaddr(&addr_QBasicMutex_unlock, "QBasicMutex_unlock")
	symaddr(&addr_QBasicMutex_tryLock, "QBasicMutex_tryLock")
	symaddr(&addr_QBasicMutex_try_lock, "QBasicMutex_try_lock")
	reg(&pg_QBasicMutex_delete, "QBasicMutex_delete")
}

type QBasicMutex struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QBasicMutex) UnsafeGoID() int64     { return this.vid }
func (this *QBasicMutex) UnsafeSetGoID(id int64) { this.vid = id }

func bindQBasicMutex(h uintptr) *QBasicMutex {
	if h == 0 {
		return nil
	}
	return &QBasicMutex{h: h}
}

func newownedQBasicMutex(h uintptr) *QBasicMutex {
	o := bindQBasicMutex(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QBasicMutex).Delete)
	}
	return o
}

func (this *QBasicMutex) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QBasicMutex) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQBasicMutex(h unsafe.Pointer) *QBasicMutex { return bindQBasicMutex(uintptr(h)) }
func UnsafeNewQBasicMutexOwned(h unsafe.Pointer) *QBasicMutex { return newownedQBasicMutex(uintptr(h)) }

func (this *QBasicMutex) UnsafeRelease() { this.h = 0 }

func NewQBasicMutex() *QBasicMutex {
	__r, _, _ := purego.SyscallN(addr_QBasicMutex_new)
	__o := newownedQBasicMutex(__r)
	return __o
}

func (this *QBasicMutex) Lock() {
	purego.SyscallN(addr_QBasicMutex_lock, this.h)
	runtime.KeepAlive(this)
}

func (this *QBasicMutex) Unlock() {
	purego.SyscallN(addr_QBasicMutex_unlock, this.h)
	runtime.KeepAlive(this)
}

func (this *QBasicMutex) TryLock() bool {
	__r0, _, _ := purego.SyscallN(addr_QBasicMutex_tryLock, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QBasicMutex) Try_lock() bool {
	__r0, _, _ := purego.SyscallN(addr_QBasicMutex_try_lock, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QBasicMutex) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QBasicMutex_delete(this.h)
	}
	this.UnsafeRelease()
}
