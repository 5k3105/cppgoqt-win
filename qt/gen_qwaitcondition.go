package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QWaitCondition_new uintptr
	addr_QWaitCondition_wait uintptr
	addr_QWaitCondition_wait2 uintptr
	addr_QWaitCondition_wait3 uintptr
	addr_QWaitCondition_wait4 uintptr
	addr_QWaitCondition_wakeOne uintptr
	addr_QWaitCondition_wakeAll uintptr
	addr_QWaitCondition_notify_one uintptr
	addr_QWaitCondition_notify_all uintptr
	addr_QWaitCondition_wait5 uintptr
	addr_QWaitCondition_wait6 uintptr
	pg_QWaitCondition_delete func(uintptr)
)

func init() {
	symaddr(&addr_QWaitCondition_new, "QWaitCondition_new")
	symaddr(&addr_QWaitCondition_wait, "QWaitCondition_wait")
	symaddr(&addr_QWaitCondition_wait2, "QWaitCondition_wait2")
	symaddr(&addr_QWaitCondition_wait3, "QWaitCondition_wait3")
	symaddr(&addr_QWaitCondition_wait4, "QWaitCondition_wait4")
	symaddr(&addr_QWaitCondition_wakeOne, "QWaitCondition_wakeOne")
	symaddr(&addr_QWaitCondition_wakeAll, "QWaitCondition_wakeAll")
	symaddr(&addr_QWaitCondition_notify_one, "QWaitCondition_notify_one")
	symaddr(&addr_QWaitCondition_notify_all, "QWaitCondition_notify_all")
	symaddr(&addr_QWaitCondition_wait5, "QWaitCondition_wait5")
	symaddr(&addr_QWaitCondition_wait6, "QWaitCondition_wait6")
	reg(&pg_QWaitCondition_delete, "QWaitCondition_delete")
}

type QWaitCondition struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QWaitCondition) UnsafeGoID() int64     { return this.vid }
func (this *QWaitCondition) UnsafeSetGoID(id int64) { this.vid = id }

func bindQWaitCondition(h uintptr) *QWaitCondition {
	if h == 0 {
		return nil
	}
	return &QWaitCondition{h: h}
}

func newownedQWaitCondition(h uintptr) *QWaitCondition {
	o := bindQWaitCondition(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QWaitCondition).Delete)
	}
	return o
}

func (this *QWaitCondition) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QWaitCondition) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQWaitCondition(h unsafe.Pointer) *QWaitCondition { return bindQWaitCondition(uintptr(h)) }
func UnsafeNewQWaitConditionOwned(h unsafe.Pointer) *QWaitCondition { return newownedQWaitCondition(uintptr(h)) }

func (this *QWaitCondition) UnsafeRelease() { this.h = 0 }

func NewQWaitCondition() *QWaitCondition {
	__r, _, _ := purego.SyscallN(addr_QWaitCondition_new)
	__o := newownedQWaitCondition(__r)
	return __o
}

func (this *QWaitCondition) Wait(lockedMutex *QMutex) bool {
	__r0, _, _ := purego.SyscallN(addr_QWaitCondition_wait, this.h, uintptr(lockedMutex.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(lockedMutex)
	return __r != 0
}

func (this *QWaitCondition) Wait2(lockedMutex *QMutex, time uint32) bool {
	__r0, _, _ := purego.SyscallN(addr_QWaitCondition_wait2, this.h, uintptr(lockedMutex.UnsafePointer()), uintptr(time))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(lockedMutex)
	return __r != 0
}

func (this *QWaitCondition) WaitWithLockedReadWriteLock(lockedReadWriteLock *QReadWriteLock) bool {
	__r0, _, _ := purego.SyscallN(addr_QWaitCondition_wait3, this.h, uintptr(lockedReadWriteLock.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(lockedReadWriteLock)
	return __r != 0
}

func (this *QWaitCondition) Wait3(lockedReadWriteLock *QReadWriteLock, time uint32) bool {
	__r0, _, _ := purego.SyscallN(addr_QWaitCondition_wait4, this.h, uintptr(lockedReadWriteLock.UnsafePointer()), uintptr(time))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(lockedReadWriteLock)
	return __r != 0
}

func (this *QWaitCondition) WakeOne() {
	purego.SyscallN(addr_QWaitCondition_wakeOne, this.h)
	runtime.KeepAlive(this)
}

func (this *QWaitCondition) WakeAll() {
	purego.SyscallN(addr_QWaitCondition_wakeAll, this.h)
	runtime.KeepAlive(this)
}

func (this *QWaitCondition) Notify_one() {
	purego.SyscallN(addr_QWaitCondition_notify_one, this.h)
	runtime.KeepAlive(this)
}

func (this *QWaitCondition) Notify_all() {
	purego.SyscallN(addr_QWaitCondition_notify_all, this.h)
	runtime.KeepAlive(this)
}

func (this *QWaitCondition) Wait4(lockedMutex *QMutex, deadline *QDeadlineTimer) bool {
	__r0, _, _ := purego.SyscallN(addr_QWaitCondition_wait5, this.h, uintptr(lockedMutex.UnsafePointer()), uintptr(deadline.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(lockedMutex)
	runtime.KeepAlive(deadline)
	return __r != 0
}

func (this *QWaitCondition) Wait5(lockedReadWriteLock *QReadWriteLock, deadline *QDeadlineTimer) bool {
	__r0, _, _ := purego.SyscallN(addr_QWaitCondition_wait6, this.h, uintptr(lockedReadWriteLock.UnsafePointer()), uintptr(deadline.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(lockedReadWriteLock)
	runtime.KeepAlive(deadline)
	return __r != 0
}

func (this *QWaitCondition) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QWaitCondition_delete(this.h)
	}
	this.UnsafeRelease()
}
