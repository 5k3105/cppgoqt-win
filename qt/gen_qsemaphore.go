package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QSemaphore_new uintptr
	addr_QSemaphore_new2 uintptr
	addr_QSemaphore_acquire uintptr
	addr_QSemaphore_tryAcquire uintptr
	addr_QSemaphore_tryAcquire2 uintptr
	addr_QSemaphore_tryAcquire3 uintptr
	addr_QSemaphore_release uintptr
	addr_QSemaphore_available uintptr
	addr_QSemaphore_try_acquire uintptr
	addr_QSemaphore_acquire2 uintptr
	addr_QSemaphore_tryAcquire4 uintptr
	addr_QSemaphore_release2 uintptr
	pg_QSemaphore_delete func(uintptr)
)

func init() {
	symaddr(&addr_QSemaphore_new, "QSemaphore_new")
	symaddr(&addr_QSemaphore_new2, "QSemaphore_new2")
	symaddr(&addr_QSemaphore_acquire, "QSemaphore_acquire")
	symaddr(&addr_QSemaphore_tryAcquire, "QSemaphore_tryAcquire")
	symaddr(&addr_QSemaphore_tryAcquire2, "QSemaphore_tryAcquire2")
	symaddr(&addr_QSemaphore_tryAcquire3, "QSemaphore_tryAcquire3")
	symaddr(&addr_QSemaphore_release, "QSemaphore_release")
	symaddr(&addr_QSemaphore_available, "QSemaphore_available")
	symaddr(&addr_QSemaphore_try_acquire, "QSemaphore_try_acquire")
	symaddr(&addr_QSemaphore_acquire2, "QSemaphore_acquire2")
	symaddr(&addr_QSemaphore_tryAcquire4, "QSemaphore_tryAcquire4")
	symaddr(&addr_QSemaphore_release2, "QSemaphore_release2")
	reg(&pg_QSemaphore_delete, "QSemaphore_delete")
}

type QSemaphore struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QSemaphore) UnsafeGoID() int64     { return this.vid }
func (this *QSemaphore) UnsafeSetGoID(id int64) { this.vid = id }

func bindQSemaphore(h uintptr) *QSemaphore {
	if h == 0 {
		return nil
	}
	return &QSemaphore{h: h}
}

func newownedQSemaphore(h uintptr) *QSemaphore {
	o := bindQSemaphore(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QSemaphore).Delete)
	}
	return o
}

func (this *QSemaphore) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QSemaphore) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQSemaphore(h unsafe.Pointer) *QSemaphore { return bindQSemaphore(uintptr(h)) }
func UnsafeNewQSemaphoreOwned(h unsafe.Pointer) *QSemaphore { return newownedQSemaphore(uintptr(h)) }

func (this *QSemaphore) UnsafeRelease() { this.h = 0 }

func NewQSemaphore() *QSemaphore {
	__r, _, _ := purego.SyscallN(addr_QSemaphore_new)
	__o := newownedQSemaphore(__r)
	return __o
}

func NewQSemaphore2(n int) *QSemaphore {
	__r, _, _ := purego.SyscallN(addr_QSemaphore_new2, uintptr(n))
	__o := newownedQSemaphore(__r)
	return __o
}

func (this *QSemaphore) Acquire() {
	purego.SyscallN(addr_QSemaphore_acquire, this.h)
	runtime.KeepAlive(this)
}

func (this *QSemaphore) TryAcquire() bool {
	__r0, _, _ := purego.SyscallN(addr_QSemaphore_tryAcquire, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QSemaphore) TryAcquire2(n int, timeout int) bool {
	__r0, _, _ := purego.SyscallN(addr_QSemaphore_tryAcquire2, this.h, uintptr(n), uintptr(timeout))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QSemaphore) TryAcquire3(n int, timeout *QDeadlineTimer) bool {
	__r0, _, _ := purego.SyscallN(addr_QSemaphore_tryAcquire3, this.h, uintptr(n), uintptr(timeout.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(timeout)
	return __r != 0
}

func (this *QSemaphore) Release() {
	purego.SyscallN(addr_QSemaphore_release, this.h)
	runtime.KeepAlive(this)
}

func (this *QSemaphore) Available() int {
	__r0, _, _ := purego.SyscallN(addr_QSemaphore_available, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QSemaphore) Try_acquire() bool {
	__r0, _, _ := purego.SyscallN(addr_QSemaphore_try_acquire, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QSemaphore) AcquireWithInt(n int) {
	purego.SyscallN(addr_QSemaphore_acquire2, this.h, uintptr(n))
	runtime.KeepAlive(this)
}

func (this *QSemaphore) TryAcquireWithInt(n int) bool {
	__r0, _, _ := purego.SyscallN(addr_QSemaphore_tryAcquire4, this.h, uintptr(n))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QSemaphore) ReleaseWithInt(n int) {
	purego.SyscallN(addr_QSemaphore_release2, this.h, uintptr(n))
	runtime.KeepAlive(this)
}

func (this *QSemaphore) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QSemaphore_delete(this.h)
	}
	this.UnsafeRelease()
}
