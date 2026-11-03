package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QDeadlineTimer_new uintptr
	addr_QDeadlineTimer_new2 uintptr
	addr_QDeadlineTimer_new3 uintptr
	addr_QDeadlineTimer_new4 uintptr
	addr_QDeadlineTimer_new5 uintptr
	addr_QDeadlineTimer_new6 uintptr
	addr_QDeadlineTimer_new7 uintptr
	addr_QDeadlineTimer_swap uintptr
	addr_QDeadlineTimer_isForever uintptr
	addr_QDeadlineTimer_hasExpired uintptr
	addr_QDeadlineTimer_timerType uintptr
	addr_QDeadlineTimer_setTimerType uintptr
	addr_QDeadlineTimer_remainingTime uintptr
	addr_QDeadlineTimer_remainingTimeNSecs uintptr
	addr_QDeadlineTimer_setRemainingTime uintptr
	addr_QDeadlineTimer_setPreciseRemainingTime uintptr
	addr_QDeadlineTimer_deadline uintptr
	addr_QDeadlineTimer_deadlineNSecs uintptr
	addr_QDeadlineTimer_setDeadline uintptr
	addr_QDeadlineTimer_setPreciseDeadline uintptr
	addr_QDeadlineTimer_addNSecs uintptr
	addr_QDeadlineTimer_current uintptr
	addr_QDeadlineTimer_OperatorPlusAssign uintptr
	addr_QDeadlineTimer_OperatorMinusAssign uintptr
	addr_QDeadlineTimer_remainingTimeAsDuration uintptr
	addr_QDeadlineTimer_setRemainingTime2 uintptr
	addr_QDeadlineTimer_setPreciseRemainingTime2 uintptr
	addr_QDeadlineTimer_setPreciseRemainingTime3 uintptr
	addr_QDeadlineTimer_setDeadline2 uintptr
	addr_QDeadlineTimer_setPreciseDeadline2 uintptr
	addr_QDeadlineTimer_setPreciseDeadline3 uintptr
	addr_QDeadlineTimer_current2 uintptr
	addr_QDeadlineTimer_OperatorPlus uintptr
	addr_QDeadlineTimer_OperatorMinus uintptr
	addr_QDeadlineTimer_OperatorMinus2 uintptr
	addr_QDeadlineTimer_OperatorEqual uintptr
	addr_QDeadlineTimer_OperatorNotEqual uintptr
	addr_QDeadlineTimer_OperatorLesser uintptr
	addr_QDeadlineTimer_OperatorGreater uintptr
	addr_QDeadlineTimer_OperatorLesserOrEqual uintptr
	addr_QDeadlineTimer_OperatorGreaterOrEqual uintptr
	pg_QDeadlineTimer_delete func(uintptr)
)

func init() {
	symaddr(&addr_QDeadlineTimer_new, "QDeadlineTimer_new")
	symaddr(&addr_QDeadlineTimer_new2, "QDeadlineTimer_new2")
	symaddr(&addr_QDeadlineTimer_new3, "QDeadlineTimer_new3")
	symaddr(&addr_QDeadlineTimer_new4, "QDeadlineTimer_new4")
	symaddr(&addr_QDeadlineTimer_new5, "QDeadlineTimer_new5")
	symaddr(&addr_QDeadlineTimer_new6, "QDeadlineTimer_new6")
	symaddr(&addr_QDeadlineTimer_new7, "QDeadlineTimer_new7")
	symaddr(&addr_QDeadlineTimer_swap, "QDeadlineTimer_swap")
	symaddr(&addr_QDeadlineTimer_isForever, "QDeadlineTimer_isForever")
	symaddr(&addr_QDeadlineTimer_hasExpired, "QDeadlineTimer_hasExpired")
	symaddr(&addr_QDeadlineTimer_timerType, "QDeadlineTimer_timerType")
	symaddr(&addr_QDeadlineTimer_setTimerType, "QDeadlineTimer_setTimerType")
	symaddr(&addr_QDeadlineTimer_remainingTime, "QDeadlineTimer_remainingTime")
	symaddr(&addr_QDeadlineTimer_remainingTimeNSecs, "QDeadlineTimer_remainingTimeNSecs")
	symaddr(&addr_QDeadlineTimer_setRemainingTime, "QDeadlineTimer_setRemainingTime")
	symaddr(&addr_QDeadlineTimer_setPreciseRemainingTime, "QDeadlineTimer_setPreciseRemainingTime")
	symaddr(&addr_QDeadlineTimer_deadline, "QDeadlineTimer_deadline")
	symaddr(&addr_QDeadlineTimer_deadlineNSecs, "QDeadlineTimer_deadlineNSecs")
	symaddr(&addr_QDeadlineTimer_setDeadline, "QDeadlineTimer_setDeadline")
	symaddr(&addr_QDeadlineTimer_setPreciseDeadline, "QDeadlineTimer_setPreciseDeadline")
	symaddr(&addr_QDeadlineTimer_addNSecs, "QDeadlineTimer_addNSecs")
	symaddr(&addr_QDeadlineTimer_current, "QDeadlineTimer_current")
	symaddr(&addr_QDeadlineTimer_OperatorPlusAssign, "QDeadlineTimer_OperatorPlusAssign")
	symaddr(&addr_QDeadlineTimer_OperatorMinusAssign, "QDeadlineTimer_OperatorMinusAssign")
	symaddr(&addr_QDeadlineTimer_remainingTimeAsDuration, "QDeadlineTimer_remainingTimeAsDuration")
	symaddr(&addr_QDeadlineTimer_setRemainingTime2, "QDeadlineTimer_setRemainingTime2")
	symaddr(&addr_QDeadlineTimer_setPreciseRemainingTime2, "QDeadlineTimer_setPreciseRemainingTime2")
	symaddr(&addr_QDeadlineTimer_setPreciseRemainingTime3, "QDeadlineTimer_setPreciseRemainingTime3")
	symaddr(&addr_QDeadlineTimer_setDeadline2, "QDeadlineTimer_setDeadline2")
	symaddr(&addr_QDeadlineTimer_setPreciseDeadline2, "QDeadlineTimer_setPreciseDeadline2")
	symaddr(&addr_QDeadlineTimer_setPreciseDeadline3, "QDeadlineTimer_setPreciseDeadline3")
	symaddr(&addr_QDeadlineTimer_current2, "QDeadlineTimer_current2")
	symaddr(&addr_QDeadlineTimer_OperatorPlus, "QDeadlineTimer_OperatorPlus")
	symaddr(&addr_QDeadlineTimer_OperatorMinus, "QDeadlineTimer_OperatorMinus")
	symaddr(&addr_QDeadlineTimer_OperatorMinus2, "QDeadlineTimer_OperatorMinus2")
	symaddr(&addr_QDeadlineTimer_OperatorEqual, "QDeadlineTimer_OperatorEqual")
	symaddr(&addr_QDeadlineTimer_OperatorNotEqual, "QDeadlineTimer_OperatorNotEqual")
	symaddr(&addr_QDeadlineTimer_OperatorLesser, "QDeadlineTimer_OperatorLesser")
	symaddr(&addr_QDeadlineTimer_OperatorGreater, "QDeadlineTimer_OperatorGreater")
	symaddr(&addr_QDeadlineTimer_OperatorLesserOrEqual, "QDeadlineTimer_OperatorLesserOrEqual")
	symaddr(&addr_QDeadlineTimer_OperatorGreaterOrEqual, "QDeadlineTimer_OperatorGreaterOrEqual")
	reg(&pg_QDeadlineTimer_delete, "QDeadlineTimer_delete")
}

type QDeadlineTimer struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QDeadlineTimer) UnsafeGoID() int64     { return this.vid }
func (this *QDeadlineTimer) UnsafeSetGoID(id int64) { this.vid = id }

func bindQDeadlineTimer(h uintptr) *QDeadlineTimer {
	if h == 0 {
		return nil
	}
	return &QDeadlineTimer{h: h}
}

func newownedQDeadlineTimer(h uintptr) *QDeadlineTimer {
	o := bindQDeadlineTimer(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QDeadlineTimer).Delete)
	}
	return o
}

func (this *QDeadlineTimer) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QDeadlineTimer) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQDeadlineTimer(h unsafe.Pointer) *QDeadlineTimer { return bindQDeadlineTimer(uintptr(h)) }
func UnsafeNewQDeadlineTimerOwned(h unsafe.Pointer) *QDeadlineTimer { return newownedQDeadlineTimer(uintptr(h)) }

func (this *QDeadlineTimer) UnsafeRelease() { this.h = 0 }

func NewQDeadlineTimer() *QDeadlineTimer {
	__r, _, _ := purego.SyscallN(addr_QDeadlineTimer_new)
	__o := newownedQDeadlineTimer(__r)
	return __o
}

func NewQDeadlineTimer2(type_ TimerType) *QDeadlineTimer {
	__r, _, _ := purego.SyscallN(addr_QDeadlineTimer_new2, uintptr(type_))
	__o := newownedQDeadlineTimer(__r)
	return __o
}

func NewQDeadlineTimer3(arg1 QDeadlineTimer__ForeverConstant) *QDeadlineTimer {
	__r, _, _ := purego.SyscallN(addr_QDeadlineTimer_new3, uintptr(arg1))
	__o := newownedQDeadlineTimer(__r)
	return __o
}

func NewQDeadlineTimer4(msecs int64) *QDeadlineTimer {
	__r, _, _ := purego.SyscallN(addr_QDeadlineTimer_new4, uintptr(msecs))
	__o := newownedQDeadlineTimer(__r)
	return __o
}

func NewQDeadlineTimer5(arg1 *QDeadlineTimer) *QDeadlineTimer {
	__r, _, _ := purego.SyscallN(addr_QDeadlineTimer_new5, uintptr(arg1.UnsafePointer()))
	__o := newownedQDeadlineTimer(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func NewQDeadlineTimer6(arg1 QDeadlineTimer__ForeverConstant, type_ TimerType) *QDeadlineTimer {
	__r, _, _ := purego.SyscallN(addr_QDeadlineTimer_new6, uintptr(arg1), uintptr(type_))
	__o := newownedQDeadlineTimer(__r)
	return __o
}

func NewQDeadlineTimer7(msecs int64, type_ TimerType) *QDeadlineTimer {
	__r, _, _ := purego.SyscallN(addr_QDeadlineTimer_new7, uintptr(msecs), uintptr(type_))
	__o := newownedQDeadlineTimer(__r)
	return __o
}

func (this *QDeadlineTimer) Swap(other *QDeadlineTimer) {
	purego.SyscallN(addr_QDeadlineTimer_swap, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
}

func (this *QDeadlineTimer) IsForever() bool {
	__r0, _, _ := purego.SyscallN(addr_QDeadlineTimer_isForever, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QDeadlineTimer) HasExpired() bool {
	__r0, _, _ := purego.SyscallN(addr_QDeadlineTimer_hasExpired, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QDeadlineTimer) TimerType() TimerType {
	__r0, _, _ := purego.SyscallN(addr_QDeadlineTimer_timerType, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return TimerType(__r)
}

func (this *QDeadlineTimer) SetTimerType(type_ TimerType) {
	purego.SyscallN(addr_QDeadlineTimer_setTimerType, this.h, uintptr(type_))
	runtime.KeepAlive(this)
}

func (this *QDeadlineTimer) RemainingTime() int64 {
	__r0, _, _ := purego.SyscallN(addr_QDeadlineTimer_remainingTime, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QDeadlineTimer) RemainingTimeNSecs() int64 {
	__r0, _, _ := purego.SyscallN(addr_QDeadlineTimer_remainingTimeNSecs, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QDeadlineTimer) SetRemainingTime(msecs int64) {
	purego.SyscallN(addr_QDeadlineTimer_setRemainingTime, this.h, uintptr(msecs))
	runtime.KeepAlive(this)
}

func (this *QDeadlineTimer) SetPreciseRemainingTime(secs int64) {
	purego.SyscallN(addr_QDeadlineTimer_setPreciseRemainingTime, this.h, uintptr(secs))
	runtime.KeepAlive(this)
}

func (this *QDeadlineTimer) Deadline() int64 {
	__r0, _, _ := purego.SyscallN(addr_QDeadlineTimer_deadline, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QDeadlineTimer) DeadlineNSecs() int64 {
	__r0, _, _ := purego.SyscallN(addr_QDeadlineTimer_deadlineNSecs, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QDeadlineTimer) SetDeadline(msecs int64) {
	purego.SyscallN(addr_QDeadlineTimer_setDeadline, this.h, uintptr(msecs))
	runtime.KeepAlive(this)
}

func (this *QDeadlineTimer) SetPreciseDeadline(secs int64) {
	purego.SyscallN(addr_QDeadlineTimer_setPreciseDeadline, this.h, uintptr(secs))
	runtime.KeepAlive(this)
}

func QDeadlineTimer_AddNSecs(dt *QDeadlineTimer, nsecs int64) *QDeadlineTimer {
	__r, _, _ := purego.SyscallN(addr_QDeadlineTimer_addNSecs, uintptr(dt.UnsafePointer()), uintptr(nsecs))
	runtime.KeepAlive(dt)
	return newownedQDeadlineTimer(__r)
}

func QDeadlineTimer_Current() *QDeadlineTimer {
	__r, _, _ := purego.SyscallN(addr_QDeadlineTimer_current)
	return newownedQDeadlineTimer(__r)
}

func (this *QDeadlineTimer) OperatorPlusAssign(msecs int64) *QDeadlineTimer {
	__r, _, _ := purego.SyscallN(addr_QDeadlineTimer_OperatorPlusAssign, this.h, uintptr(msecs))
	runtime.KeepAlive(this)
	return bindQDeadlineTimer(__r)
}

func (this *QDeadlineTimer) OperatorMinusAssign(msecs int64) *QDeadlineTimer {
	__r, _, _ := purego.SyscallN(addr_QDeadlineTimer_OperatorMinusAssign, this.h, uintptr(msecs))
	runtime.KeepAlive(this)
	return bindQDeadlineTimer(__r)
}

func (this *QDeadlineTimer) RemainingTimeAsDuration() int64 {
	__r0, _, _ := purego.SyscallN(addr_QDeadlineTimer_remainingTimeAsDuration, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QDeadlineTimer) SetRemainingTime2(msecs int64, type_ TimerType) {
	purego.SyscallN(addr_QDeadlineTimer_setRemainingTime2, this.h, uintptr(msecs), uintptr(type_))
	runtime.KeepAlive(this)
}

func (this *QDeadlineTimer) SetPreciseRemainingTime2(secs int64, nsecs int64) {
	purego.SyscallN(addr_QDeadlineTimer_setPreciseRemainingTime2, this.h, uintptr(secs), uintptr(nsecs))
	runtime.KeepAlive(this)
}

func (this *QDeadlineTimer) SetPreciseRemainingTime3(secs int64, nsecs int64, type_ TimerType) {
	purego.SyscallN(addr_QDeadlineTimer_setPreciseRemainingTime3, this.h, uintptr(secs), uintptr(nsecs), uintptr(type_))
	runtime.KeepAlive(this)
}

func (this *QDeadlineTimer) SetDeadline2(msecs int64, timerType TimerType) {
	purego.SyscallN(addr_QDeadlineTimer_setDeadline2, this.h, uintptr(msecs), uintptr(timerType))
	runtime.KeepAlive(this)
}

func (this *QDeadlineTimer) SetPreciseDeadline2(secs int64, nsecs int64) {
	purego.SyscallN(addr_QDeadlineTimer_setPreciseDeadline2, this.h, uintptr(secs), uintptr(nsecs))
	runtime.KeepAlive(this)
}

func (this *QDeadlineTimer) SetPreciseDeadline3(secs int64, nsecs int64, type_ TimerType) {
	purego.SyscallN(addr_QDeadlineTimer_setPreciseDeadline3, this.h, uintptr(secs), uintptr(nsecs), uintptr(type_))
	runtime.KeepAlive(this)
}

func QDeadlineTimer_CurrentWithTimerType(timerType TimerType) *QDeadlineTimer {
	__r, _, _ := purego.SyscallN(addr_QDeadlineTimer_current2, uintptr(timerType))
	return newownedQDeadlineTimer(__r)
}

func (this *QDeadlineTimer) OperatorPlus(msecs int64) *QDeadlineTimer {
	__r, _, _ := purego.SyscallN(addr_QDeadlineTimer_OperatorPlus, this.h, uintptr(msecs))
	runtime.KeepAlive(this)
	return newownedQDeadlineTimer(__r)
}

func (this *QDeadlineTimer) OperatorMinus(msecs int64) *QDeadlineTimer {
	__r, _, _ := purego.SyscallN(addr_QDeadlineTimer_OperatorMinus, this.h, uintptr(msecs))
	runtime.KeepAlive(this)
	return newownedQDeadlineTimer(__r)
}

func (this *QDeadlineTimer) OperatorMinusWithDt2(dt2 *QDeadlineTimer) int64 {
	__r0, _, _ := purego.SyscallN(addr_QDeadlineTimer_OperatorMinus2, this.h, uintptr(dt2.UnsafePointer()))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(dt2)
	return __r
}

func (this *QDeadlineTimer) OperatorEqual(rhs *QDeadlineTimer) bool {
	__r0, _, _ := purego.SyscallN(addr_QDeadlineTimer_OperatorEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QDeadlineTimer) OperatorNotEqual(rhs *QDeadlineTimer) bool {
	__r0, _, _ := purego.SyscallN(addr_QDeadlineTimer_OperatorNotEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QDeadlineTimer) OperatorLesser(rhs *QDeadlineTimer) bool {
	__r0, _, _ := purego.SyscallN(addr_QDeadlineTimer_OperatorLesser, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QDeadlineTimer) OperatorGreater(rhs *QDeadlineTimer) bool {
	__r0, _, _ := purego.SyscallN(addr_QDeadlineTimer_OperatorGreater, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QDeadlineTimer) OperatorLesserOrEqual(rhs *QDeadlineTimer) bool {
	__r0, _, _ := purego.SyscallN(addr_QDeadlineTimer_OperatorLesserOrEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QDeadlineTimer) OperatorGreaterOrEqual(rhs *QDeadlineTimer) bool {
	__r0, _, _ := purego.SyscallN(addr_QDeadlineTimer_OperatorGreaterOrEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QDeadlineTimer) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QDeadlineTimer_delete(this.h)
	}
	this.UnsafeRelease()
}
