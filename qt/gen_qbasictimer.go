package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QBasicTimer_new uintptr
	addr_QBasicTimer_swap uintptr
	addr_QBasicTimer_isActive uintptr
	addr_QBasicTimer_timerId uintptr
	addr_QBasicTimer_id uintptr
	addr_QBasicTimer_start uintptr
	addr_QBasicTimer_start2 uintptr
	addr_QBasicTimer_stop uintptr
	pg_QBasicTimer_delete func(uintptr)
)

func init() {
	symaddr(&addr_QBasicTimer_new, "QBasicTimer_new")
	symaddr(&addr_QBasicTimer_swap, "QBasicTimer_swap")
	symaddr(&addr_QBasicTimer_isActive, "QBasicTimer_isActive")
	symaddr(&addr_QBasicTimer_timerId, "QBasicTimer_timerId")
	symaddr(&addr_QBasicTimer_id, "QBasicTimer_id")
	symaddr(&addr_QBasicTimer_start, "QBasicTimer_start")
	symaddr(&addr_QBasicTimer_start2, "QBasicTimer_start2")
	symaddr(&addr_QBasicTimer_stop, "QBasicTimer_stop")
	reg(&pg_QBasicTimer_delete, "QBasicTimer_delete")
}

type QBasicTimer struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QBasicTimer) UnsafeGoID() int64     { return this.vid }
func (this *QBasicTimer) UnsafeSetGoID(id int64) { this.vid = id }

func bindQBasicTimer(h uintptr) *QBasicTimer {
	if h == 0 {
		return nil
	}
	return &QBasicTimer{h: h}
}

func newownedQBasicTimer(h uintptr) *QBasicTimer {
	o := bindQBasicTimer(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QBasicTimer).Delete)
	}
	return o
}

func (this *QBasicTimer) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QBasicTimer) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQBasicTimer(h unsafe.Pointer) *QBasicTimer { return bindQBasicTimer(uintptr(h)) }
func UnsafeNewQBasicTimerOwned(h unsafe.Pointer) *QBasicTimer { return newownedQBasicTimer(uintptr(h)) }

func (this *QBasicTimer) UnsafeRelease() { this.h = 0 }

func NewQBasicTimer() *QBasicTimer {
	__r, _, _ := purego.SyscallN(addr_QBasicTimer_new)
	__o := newownedQBasicTimer(__r)
	return __o
}

func (this *QBasicTimer) Swap(other *QBasicTimer) {
	purego.SyscallN(addr_QBasicTimer_swap, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
}

func (this *QBasicTimer) IsActive() bool {
	__r0, _, _ := purego.SyscallN(addr_QBasicTimer_isActive, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QBasicTimer) TimerId() int {
	__r0, _, _ := purego.SyscallN(addr_QBasicTimer_timerId, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QBasicTimer) Id() TimerId {
	__r0, _, _ := purego.SyscallN(addr_QBasicTimer_id, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return TimerId(__r)
}

func (this *QBasicTimer) Start(msec int, obj *QObject) {
	purego.SyscallN(addr_QBasicTimer_start, this.h, uintptr(msec), uintptr(obj.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(obj)
}

func (this *QBasicTimer) Start2(msec int, timerType TimerType, obj *QObject) {
	purego.SyscallN(addr_QBasicTimer_start2, this.h, uintptr(msec), uintptr(timerType), uintptr(obj.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(obj)
}

func (this *QBasicTimer) Stop() {
	purego.SyscallN(addr_QBasicTimer_stop, this.h)
	runtime.KeepAlive(this)
}

func (this *QBasicTimer) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QBasicTimer_delete(this.h)
	}
	this.UnsafeRelease()
}
