package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QElapsedTimer_new uintptr
	addr_QElapsedTimer_clockType uintptr
	addr_QElapsedTimer_isMonotonic uintptr
	addr_QElapsedTimer_start uintptr
	addr_QElapsedTimer_restart uintptr
	addr_QElapsedTimer_invalidate uintptr
	addr_QElapsedTimer_isValid uintptr
	addr_QElapsedTimer_nsecsElapsed uintptr
	addr_QElapsedTimer_elapsed uintptr
	addr_QElapsedTimer_hasExpired uintptr
	addr_QElapsedTimer_msecsSinceReference uintptr
	addr_QElapsedTimer_msecsTo uintptr
	addr_QElapsedTimer_secsTo uintptr
	addr_QElapsedTimer_OperatorLesser uintptr
	addr_QElapsedTimer_OperatorEqual uintptr
	addr_QElapsedTimer_OperatorNotEqual uintptr
	addr_QElapsedTimer_OperatorGreater uintptr
	addr_QElapsedTimer_OperatorLesserOrEqual uintptr
	addr_QElapsedTimer_OperatorGreaterOrEqual uintptr
	pg_QElapsedTimer_delete func(uintptr)
)

func init() {
	symaddr(&addr_QElapsedTimer_new, "QElapsedTimer_new")
	symaddr(&addr_QElapsedTimer_clockType, "QElapsedTimer_clockType")
	symaddr(&addr_QElapsedTimer_isMonotonic, "QElapsedTimer_isMonotonic")
	symaddr(&addr_QElapsedTimer_start, "QElapsedTimer_start")
	symaddr(&addr_QElapsedTimer_restart, "QElapsedTimer_restart")
	symaddr(&addr_QElapsedTimer_invalidate, "QElapsedTimer_invalidate")
	symaddr(&addr_QElapsedTimer_isValid, "QElapsedTimer_isValid")
	symaddr(&addr_QElapsedTimer_nsecsElapsed, "QElapsedTimer_nsecsElapsed")
	symaddr(&addr_QElapsedTimer_elapsed, "QElapsedTimer_elapsed")
	symaddr(&addr_QElapsedTimer_hasExpired, "QElapsedTimer_hasExpired")
	symaddr(&addr_QElapsedTimer_msecsSinceReference, "QElapsedTimer_msecsSinceReference")
	symaddr(&addr_QElapsedTimer_msecsTo, "QElapsedTimer_msecsTo")
	symaddr(&addr_QElapsedTimer_secsTo, "QElapsedTimer_secsTo")
	symaddr(&addr_QElapsedTimer_OperatorLesser, "QElapsedTimer_OperatorLesser")
	symaddr(&addr_QElapsedTimer_OperatorEqual, "QElapsedTimer_OperatorEqual")
	symaddr(&addr_QElapsedTimer_OperatorNotEqual, "QElapsedTimer_OperatorNotEqual")
	symaddr(&addr_QElapsedTimer_OperatorGreater, "QElapsedTimer_OperatorGreater")
	symaddr(&addr_QElapsedTimer_OperatorLesserOrEqual, "QElapsedTimer_OperatorLesserOrEqual")
	symaddr(&addr_QElapsedTimer_OperatorGreaterOrEqual, "QElapsedTimer_OperatorGreaterOrEqual")
	reg(&pg_QElapsedTimer_delete, "QElapsedTimer_delete")
}

type QElapsedTimer struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QElapsedTimer) UnsafeGoID() int64     { return this.vid }
func (this *QElapsedTimer) UnsafeSetGoID(id int64) { this.vid = id }

func bindQElapsedTimer(h uintptr) *QElapsedTimer {
	if h == 0 {
		return nil
	}
	return &QElapsedTimer{h: h}
}

func newownedQElapsedTimer(h uintptr) *QElapsedTimer {
	o := bindQElapsedTimer(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QElapsedTimer).Delete)
	}
	return o
}

func (this *QElapsedTimer) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QElapsedTimer) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQElapsedTimer(h unsafe.Pointer) *QElapsedTimer { return bindQElapsedTimer(uintptr(h)) }
func UnsafeNewQElapsedTimerOwned(h unsafe.Pointer) *QElapsedTimer { return newownedQElapsedTimer(uintptr(h)) }

func (this *QElapsedTimer) UnsafeRelease() { this.h = 0 }

func NewQElapsedTimer() *QElapsedTimer {
	__r, _, _ := purego.SyscallN(addr_QElapsedTimer_new)
	__o := newownedQElapsedTimer(__r)
	return __o
}

func QElapsedTimer_ClockType() QElapsedTimer__ClockType {
	__r0, _, _ := purego.SyscallN(addr_QElapsedTimer_clockType)
	__r := int32(uint32(__r0))
	return QElapsedTimer__ClockType(__r)
}

func QElapsedTimer_IsMonotonic() bool {
	__r0, _, _ := purego.SyscallN(addr_QElapsedTimer_isMonotonic)
	__r := int32(uint32(__r0))
	return __r != 0
}

func (this *QElapsedTimer) Start() {
	purego.SyscallN(addr_QElapsedTimer_start, this.h)
	runtime.KeepAlive(this)
}

func (this *QElapsedTimer) Restart() int64 {
	__r0, _, _ := purego.SyscallN(addr_QElapsedTimer_restart, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QElapsedTimer) Invalidate() {
	purego.SyscallN(addr_QElapsedTimer_invalidate, this.h)
	runtime.KeepAlive(this)
}

func (this *QElapsedTimer) IsValid() bool {
	__r0, _, _ := purego.SyscallN(addr_QElapsedTimer_isValid, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QElapsedTimer) NsecsElapsed() int64 {
	__r0, _, _ := purego.SyscallN(addr_QElapsedTimer_nsecsElapsed, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QElapsedTimer) Elapsed() int64 {
	__r0, _, _ := purego.SyscallN(addr_QElapsedTimer_elapsed, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QElapsedTimer) HasExpired(timeout int64) bool {
	__r0, _, _ := purego.SyscallN(addr_QElapsedTimer_hasExpired, this.h, uintptr(timeout))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QElapsedTimer) MsecsSinceReference() int64 {
	__r0, _, _ := purego.SyscallN(addr_QElapsedTimer_msecsSinceReference, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QElapsedTimer) MsecsTo(other *QElapsedTimer) int64 {
	__r0, _, _ := purego.SyscallN(addr_QElapsedTimer_msecsTo, this.h, uintptr(other.UnsafePointer()))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return __r
}

func (this *QElapsedTimer) SecsTo(other *QElapsedTimer) int64 {
	__r0, _, _ := purego.SyscallN(addr_QElapsedTimer_secsTo, this.h, uintptr(other.UnsafePointer()))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return __r
}

func (this *QElapsedTimer) OperatorLesser(rhs *QElapsedTimer) bool {
	__r0, _, _ := purego.SyscallN(addr_QElapsedTimer_OperatorLesser, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QElapsedTimer) OperatorEqual(rhs *QElapsedTimer) bool {
	__r0, _, _ := purego.SyscallN(addr_QElapsedTimer_OperatorEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QElapsedTimer) OperatorNotEqual(rhs *QElapsedTimer) bool {
	__r0, _, _ := purego.SyscallN(addr_QElapsedTimer_OperatorNotEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QElapsedTimer) OperatorGreater(rhs *QElapsedTimer) bool {
	__r0, _, _ := purego.SyscallN(addr_QElapsedTimer_OperatorGreater, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QElapsedTimer) OperatorLesserOrEqual(rhs *QElapsedTimer) bool {
	__r0, _, _ := purego.SyscallN(addr_QElapsedTimer_OperatorLesserOrEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QElapsedTimer) OperatorGreaterOrEqual(rhs *QElapsedTimer) bool {
	__r0, _, _ := purego.SyscallN(addr_QElapsedTimer_OperatorGreaterOrEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QElapsedTimer) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QElapsedTimer_delete(this.h)
	}
	this.UnsafeRelease()
}
