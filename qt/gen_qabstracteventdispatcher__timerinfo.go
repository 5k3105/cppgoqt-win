package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QAbstractEventDispatcher__TimerInfo_new uintptr
	addr_QAbstractEventDispatcher__TimerInfo_new2 uintptr
	addr_QAbstractEventDispatcher__TimerInfo_new3 uintptr
	addr_QAbstractEventDispatcher__TimerInfo_timerId_get uintptr
	addr_QAbstractEventDispatcher__TimerInfo_timerId_set uintptr
	addr_QAbstractEventDispatcher__TimerInfo_interval_get uintptr
	addr_QAbstractEventDispatcher__TimerInfo_interval_set uintptr
	addr_QAbstractEventDispatcher__TimerInfo_timerType_get uintptr
	addr_QAbstractEventDispatcher__TimerInfo_timerType_set uintptr
	pg_QAbstractEventDispatcher__TimerInfo_delete func(uintptr)
)

func init() {
	symaddr(&addr_QAbstractEventDispatcher__TimerInfo_new, "QAbstractEventDispatcher__TimerInfo_new")
	symaddr(&addr_QAbstractEventDispatcher__TimerInfo_new2, "QAbstractEventDispatcher__TimerInfo_new2")
	symaddr(&addr_QAbstractEventDispatcher__TimerInfo_new3, "QAbstractEventDispatcher__TimerInfo_new3")
	symaddr(&addr_QAbstractEventDispatcher__TimerInfo_timerId_get, "QAbstractEventDispatcher__TimerInfo_timerId_get")
	symaddr(&addr_QAbstractEventDispatcher__TimerInfo_timerId_set, "QAbstractEventDispatcher__TimerInfo_timerId_set")
	symaddr(&addr_QAbstractEventDispatcher__TimerInfo_interval_get, "QAbstractEventDispatcher__TimerInfo_interval_get")
	symaddr(&addr_QAbstractEventDispatcher__TimerInfo_interval_set, "QAbstractEventDispatcher__TimerInfo_interval_set")
	symaddr(&addr_QAbstractEventDispatcher__TimerInfo_timerType_get, "QAbstractEventDispatcher__TimerInfo_timerType_get")
	symaddr(&addr_QAbstractEventDispatcher__TimerInfo_timerType_set, "QAbstractEventDispatcher__TimerInfo_timerType_set")
	reg(&pg_QAbstractEventDispatcher__TimerInfo_delete, "QAbstractEventDispatcher__TimerInfo_delete")
}

type QAbstractEventDispatcher__TimerInfo struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QAbstractEventDispatcher__TimerInfo) UnsafeGoID() int64     { return this.vid }
func (this *QAbstractEventDispatcher__TimerInfo) UnsafeSetGoID(id int64) { this.vid = id }

func bindQAbstractEventDispatcher__TimerInfo(h uintptr) *QAbstractEventDispatcher__TimerInfo {
	if h == 0 {
		return nil
	}
	return &QAbstractEventDispatcher__TimerInfo{h: h}
}

func newownedQAbstractEventDispatcher__TimerInfo(h uintptr) *QAbstractEventDispatcher__TimerInfo {
	o := bindQAbstractEventDispatcher__TimerInfo(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QAbstractEventDispatcher__TimerInfo).Delete)
	}
	return o
}

func (this *QAbstractEventDispatcher__TimerInfo) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QAbstractEventDispatcher__TimerInfo) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQAbstractEventDispatcher__TimerInfo(h unsafe.Pointer) *QAbstractEventDispatcher__TimerInfo { return bindQAbstractEventDispatcher__TimerInfo(uintptr(h)) }
func UnsafeNewQAbstractEventDispatcher__TimerInfoOwned(h unsafe.Pointer) *QAbstractEventDispatcher__TimerInfo { return newownedQAbstractEventDispatcher__TimerInfo(uintptr(h)) }

func (this *QAbstractEventDispatcher__TimerInfo) UnsafeRelease() { this.h = 0 }

func NewQAbstractEventDispatcher__TimerInfo() *QAbstractEventDispatcher__TimerInfo {
	__r, _, _ := purego.SyscallN(addr_QAbstractEventDispatcher__TimerInfo_new)
	__o := newownedQAbstractEventDispatcher__TimerInfo(__r)
	return __o
}

func NewQAbstractEventDispatcher__TimerInfo2(id int, i int, t TimerType) *QAbstractEventDispatcher__TimerInfo {
	__r, _, _ := purego.SyscallN(addr_QAbstractEventDispatcher__TimerInfo_new2, uintptr(id), uintptr(i), uintptr(t))
	__o := newownedQAbstractEventDispatcher__TimerInfo(__r)
	return __o
}

func NewQAbstractEventDispatcher__TimerInfo3(arg1 *QAbstractEventDispatcher__TimerInfo) *QAbstractEventDispatcher__TimerInfo {
	__r, _, _ := purego.SyscallN(addr_QAbstractEventDispatcher__TimerInfo_new3, uintptr(arg1.UnsafePointer()))
	__o := newownedQAbstractEventDispatcher__TimerInfo(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QAbstractEventDispatcher__TimerInfo) TimerId() int {
	__r0, _, _ := purego.SyscallN(addr_QAbstractEventDispatcher__TimerInfo_timerId_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAbstractEventDispatcher__TimerInfo) SetTimerId(timerId int) {
	purego.SyscallN(addr_QAbstractEventDispatcher__TimerInfo_timerId_set, this.h, uintptr(timerId))
	runtime.KeepAlive(this)
}

func (this *QAbstractEventDispatcher__TimerInfo) Interval() int {
	__r0, _, _ := purego.SyscallN(addr_QAbstractEventDispatcher__TimerInfo_interval_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAbstractEventDispatcher__TimerInfo) SetInterval(interval int) {
	purego.SyscallN(addr_QAbstractEventDispatcher__TimerInfo_interval_set, this.h, uintptr(interval))
	runtime.KeepAlive(this)
}

func (this *QAbstractEventDispatcher__TimerInfo) TimerType() TimerType {
	__r0, _, _ := purego.SyscallN(addr_QAbstractEventDispatcher__TimerInfo_timerType_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return TimerType(__r)
}

func (this *QAbstractEventDispatcher__TimerInfo) SetTimerType(timerType TimerType) {
	purego.SyscallN(addr_QAbstractEventDispatcher__TimerInfo_timerType_set, this.h, uintptr(timerType))
	runtime.KeepAlive(this)
}

func (this *QAbstractEventDispatcher__TimerInfo) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QAbstractEventDispatcher__TimerInfo_delete(this.h)
	}
	this.UnsafeRelease()
}
