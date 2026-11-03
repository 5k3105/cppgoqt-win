package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QSemaphoreReleaser_new uintptr
	addr_QSemaphoreReleaser_new2 uintptr
	addr_QSemaphoreReleaser_new3 uintptr
	addr_QSemaphoreReleaser_new4 uintptr
	addr_QSemaphoreReleaser_new5 uintptr
	addr_QSemaphoreReleaser_swap uintptr
	addr_QSemaphoreReleaser_semaphore uintptr
	addr_QSemaphoreReleaser_cancel uintptr
	pg_QSemaphoreReleaser_delete func(uintptr)
)

func init() {
	symaddr(&addr_QSemaphoreReleaser_new, "QSemaphoreReleaser_new")
	symaddr(&addr_QSemaphoreReleaser_new2, "QSemaphoreReleaser_new2")
	symaddr(&addr_QSemaphoreReleaser_new3, "QSemaphoreReleaser_new3")
	symaddr(&addr_QSemaphoreReleaser_new4, "QSemaphoreReleaser_new4")
	symaddr(&addr_QSemaphoreReleaser_new5, "QSemaphoreReleaser_new5")
	symaddr(&addr_QSemaphoreReleaser_swap, "QSemaphoreReleaser_swap")
	symaddr(&addr_QSemaphoreReleaser_semaphore, "QSemaphoreReleaser_semaphore")
	symaddr(&addr_QSemaphoreReleaser_cancel, "QSemaphoreReleaser_cancel")
	reg(&pg_QSemaphoreReleaser_delete, "QSemaphoreReleaser_delete")
}

type QSemaphoreReleaser struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QSemaphoreReleaser) UnsafeGoID() int64     { return this.vid }
func (this *QSemaphoreReleaser) UnsafeSetGoID(id int64) { this.vid = id }

func bindQSemaphoreReleaser(h uintptr) *QSemaphoreReleaser {
	if h == 0 {
		return nil
	}
	return &QSemaphoreReleaser{h: h}
}

func newownedQSemaphoreReleaser(h uintptr) *QSemaphoreReleaser {
	o := bindQSemaphoreReleaser(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QSemaphoreReleaser).Delete)
	}
	return o
}

func (this *QSemaphoreReleaser) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QSemaphoreReleaser) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQSemaphoreReleaser(h unsafe.Pointer) *QSemaphoreReleaser { return bindQSemaphoreReleaser(uintptr(h)) }
func UnsafeNewQSemaphoreReleaserOwned(h unsafe.Pointer) *QSemaphoreReleaser { return newownedQSemaphoreReleaser(uintptr(h)) }

func (this *QSemaphoreReleaser) UnsafeRelease() { this.h = 0 }

func NewQSemaphoreReleaser() *QSemaphoreReleaser {
	__r, _, _ := purego.SyscallN(addr_QSemaphoreReleaser_new)
	__o := newownedQSemaphoreReleaser(__r)
	return __o
}

func NewQSemaphoreReleaser2(sem *QSemaphore) *QSemaphoreReleaser {
	__r, _, _ := purego.SyscallN(addr_QSemaphoreReleaser_new2, uintptr(sem.UnsafePointer()))
	__o := newownedQSemaphoreReleaser(__r)
	runtime.KeepAlive(sem)
	return __o
}

func NewQSemaphoreReleaser3(sem *QSemaphore) *QSemaphoreReleaser {
	__r, _, _ := purego.SyscallN(addr_QSemaphoreReleaser_new3, uintptr(sem.UnsafePointer()))
	__o := newownedQSemaphoreReleaser(__r)
	runtime.KeepAlive(sem)
	return __o
}

func NewQSemaphoreReleaser4(sem *QSemaphore, n int) *QSemaphoreReleaser {
	__r, _, _ := purego.SyscallN(addr_QSemaphoreReleaser_new4, uintptr(sem.UnsafePointer()), uintptr(n))
	__o := newownedQSemaphoreReleaser(__r)
	runtime.KeepAlive(sem)
	return __o
}

func NewQSemaphoreReleaser5(sem *QSemaphore, n int) *QSemaphoreReleaser {
	__r, _, _ := purego.SyscallN(addr_QSemaphoreReleaser_new5, uintptr(sem.UnsafePointer()), uintptr(n))
	__o := newownedQSemaphoreReleaser(__r)
	runtime.KeepAlive(sem)
	return __o
}

func (this *QSemaphoreReleaser) Swap(other *QSemaphoreReleaser) {
	purego.SyscallN(addr_QSemaphoreReleaser_swap, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
}

func (this *QSemaphoreReleaser) Semaphore() *QSemaphore {
	__r, _, _ := purego.SyscallN(addr_QSemaphoreReleaser_semaphore, this.h)
	runtime.KeepAlive(this)
	return bindQSemaphore(__r)
}

func (this *QSemaphoreReleaser) Cancel() *QSemaphore {
	__r, _, _ := purego.SyscallN(addr_QSemaphoreReleaser_cancel, this.h)
	runtime.KeepAlive(this)
	return bindQSemaphore(__r)
}

func (this *QSemaphoreReleaser) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QSemaphoreReleaser_delete(this.h)
	}
	this.UnsafeRelease()
}
