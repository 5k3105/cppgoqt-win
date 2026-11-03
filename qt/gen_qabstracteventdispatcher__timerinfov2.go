package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QAbstractEventDispatcher__TimerInfoV2_new uintptr
	addr_QAbstractEventDispatcher__TimerInfoV2_new2 uintptr
	addr_QAbstractEventDispatcher__TimerInfoV2_OperatorAssign uintptr
	addr_QAbstractEventDispatcher__TimerInfoV2_timerId_get uintptr
	addr_QAbstractEventDispatcher__TimerInfoV2_timerId_set uintptr
	addr_QAbstractEventDispatcher__TimerInfoV2_timerType_get uintptr
	addr_QAbstractEventDispatcher__TimerInfoV2_timerType_set uintptr
	pg_QAbstractEventDispatcher__TimerInfoV2_delete func(uintptr)
)

func init() {
	symaddr(&addr_QAbstractEventDispatcher__TimerInfoV2_new, "QAbstractEventDispatcher__TimerInfoV2_new")
	symaddr(&addr_QAbstractEventDispatcher__TimerInfoV2_new2, "QAbstractEventDispatcher__TimerInfoV2_new2")
	symaddr(&addr_QAbstractEventDispatcher__TimerInfoV2_OperatorAssign, "QAbstractEventDispatcher__TimerInfoV2_OperatorAssign")
	symaddr(&addr_QAbstractEventDispatcher__TimerInfoV2_timerId_get, "QAbstractEventDispatcher__TimerInfoV2_timerId_get")
	symaddr(&addr_QAbstractEventDispatcher__TimerInfoV2_timerId_set, "QAbstractEventDispatcher__TimerInfoV2_timerId_set")
	symaddr(&addr_QAbstractEventDispatcher__TimerInfoV2_timerType_get, "QAbstractEventDispatcher__TimerInfoV2_timerType_get")
	symaddr(&addr_QAbstractEventDispatcher__TimerInfoV2_timerType_set, "QAbstractEventDispatcher__TimerInfoV2_timerType_set")
	reg(&pg_QAbstractEventDispatcher__TimerInfoV2_delete, "QAbstractEventDispatcher__TimerInfoV2_delete")
}

type QAbstractEventDispatcher__TimerInfoV2 struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QAbstractEventDispatcher__TimerInfoV2) UnsafeGoID() int64     { return this.vid }
func (this *QAbstractEventDispatcher__TimerInfoV2) UnsafeSetGoID(id int64) { this.vid = id }

func bindQAbstractEventDispatcher__TimerInfoV2(h uintptr) *QAbstractEventDispatcher__TimerInfoV2 {
	if h == 0 {
		return nil
	}
	return &QAbstractEventDispatcher__TimerInfoV2{h: h}
}

func newownedQAbstractEventDispatcher__TimerInfoV2(h uintptr) *QAbstractEventDispatcher__TimerInfoV2 {
	o := bindQAbstractEventDispatcher__TimerInfoV2(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QAbstractEventDispatcher__TimerInfoV2).Delete)
	}
	return o
}

func (this *QAbstractEventDispatcher__TimerInfoV2) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QAbstractEventDispatcher__TimerInfoV2) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQAbstractEventDispatcher__TimerInfoV2(h unsafe.Pointer) *QAbstractEventDispatcher__TimerInfoV2 { return bindQAbstractEventDispatcher__TimerInfoV2(uintptr(h)) }
func UnsafeNewQAbstractEventDispatcher__TimerInfoV2Owned(h unsafe.Pointer) *QAbstractEventDispatcher__TimerInfoV2 { return newownedQAbstractEventDispatcher__TimerInfoV2(uintptr(h)) }

func (this *QAbstractEventDispatcher__TimerInfoV2) UnsafeRelease() { this.h = 0 }

func NewQAbstractEventDispatcher__TimerInfoV2(arg1 *QAbstractEventDispatcher__TimerInfoV2) *QAbstractEventDispatcher__TimerInfoV2 {
	__r, _, _ := purego.SyscallN(addr_QAbstractEventDispatcher__TimerInfoV2_new, uintptr(arg1.UnsafePointer()))
	__o := newownedQAbstractEventDispatcher__TimerInfoV2(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func NewQAbstractEventDispatcher__TimerInfoV22() *QAbstractEventDispatcher__TimerInfoV2 {
	__r, _, _ := purego.SyscallN(addr_QAbstractEventDispatcher__TimerInfoV2_new2)
	__o := newownedQAbstractEventDispatcher__TimerInfoV2(__r)
	return __o
}

func (this *QAbstractEventDispatcher__TimerInfoV2) OperatorAssign(arg1 *QAbstractEventDispatcher__TimerInfoV2) *QAbstractEventDispatcher__TimerInfoV2 {
	__r, _, _ := purego.SyscallN(addr_QAbstractEventDispatcher__TimerInfoV2_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQAbstractEventDispatcher__TimerInfoV2(__r)
}

func (this *QAbstractEventDispatcher__TimerInfoV2) TimerId() TimerId {
	__r0, _, _ := purego.SyscallN(addr_QAbstractEventDispatcher__TimerInfoV2_timerId_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return TimerId(__r)
}

func (this *QAbstractEventDispatcher__TimerInfoV2) SetTimerId(timerId TimerId) {
	purego.SyscallN(addr_QAbstractEventDispatcher__TimerInfoV2_timerId_set, this.h, uintptr(timerId))
	runtime.KeepAlive(this)
}

func (this *QAbstractEventDispatcher__TimerInfoV2) TimerType() TimerType {
	__r0, _, _ := purego.SyscallN(addr_QAbstractEventDispatcher__TimerInfoV2_timerType_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return TimerType(__r)
}

func (this *QAbstractEventDispatcher__TimerInfoV2) SetTimerType(timerType TimerType) {
	purego.SyscallN(addr_QAbstractEventDispatcher__TimerInfoV2_timerType_set, this.h, uintptr(timerType))
	runtime.KeepAlive(this)
}

func (this *QAbstractEventDispatcher__TimerInfoV2) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QAbstractEventDispatcher__TimerInfoV2_delete(this.h)
	}
	this.UnsafeRelease()
}
