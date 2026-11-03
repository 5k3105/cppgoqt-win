package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QStatusTipEvent_base func(uintptr) uintptr
	addr_QStatusTipEvent_new uintptr
	addr_QStatusTipEvent_tip uintptr
	addr_QStatusTipEvent_registerEventType uintptr
	addr_QStatusTipEvent_registerEventType2 uintptr
	pg_QStatusTipEvent_delete func(uintptr)
)

func init() {
	reg(&pg_QStatusTipEvent_base, "QStatusTipEvent_base")
	symaddr(&addr_QStatusTipEvent_new, "QStatusTipEvent_new")
	symaddr(&addr_QStatusTipEvent_tip, "QStatusTipEvent_tip")
	symaddr(&addr_QStatusTipEvent_registerEventType, "QStatusTipEvent_registerEventType")
	symaddr(&addr_QStatusTipEvent_registerEventType2, "QStatusTipEvent_registerEventType2")
	reg(&pg_QStatusTipEvent_delete, "QStatusTipEvent_delete")
}

type QStatusTipEvent struct {
	h uintptr
	*QEvent
}

func bindQStatusTipEvent(h uintptr) *QStatusTipEvent {
	if h == 0 {
		return nil
	}
	return &QStatusTipEvent{h: h, QEvent: bindQEvent(pg_QStatusTipEvent_base(h))}
}

func newownedQStatusTipEvent(h uintptr) *QStatusTipEvent {
	o := bindQStatusTipEvent(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QStatusTipEvent).Delete)
	}
	return o
}

func (this *QStatusTipEvent) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QStatusTipEvent) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQStatusTipEvent(h unsafe.Pointer) *QStatusTipEvent { return bindQStatusTipEvent(uintptr(h)) }
func UnsafeNewQStatusTipEventOwned(h unsafe.Pointer) *QStatusTipEvent { return newownedQStatusTipEvent(uintptr(h)) }

func (this *QStatusTipEvent) UnsafeRelease() {
	this.h = 0
	if this.QEvent != nil {
		this.QEvent.UnsafeRelease()
	}
}

func NewQStatusTipEvent(tip string) *QStatusTipEvent {
	var tip_p uintptr
	if len(tip) > 0 {
		tip_p = uintptr(unsafe.Pointer(unsafe.StringData(tip)))
	}
	__r, _, _ := purego.SyscallN(addr_QStatusTipEvent_new, tip_p, uintptr(len(tip)))
	__o := newownedQStatusTipEvent(__r)
	runtime.KeepAlive(tip)
	return __o
}

func (this *QStatusTipEvent) Tip() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStatusTipEvent_tip, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func QStatusTipEvent_RegisterEventType() int {
	__r0, _, _ := purego.SyscallN(addr_QStatusTipEvent_registerEventType)
	__r := int32(uint32(__r0))
	return int(__r)
}

func QStatusTipEvent_RegisterEventTypeWithHint(hint int) int {
	__r0, _, _ := purego.SyscallN(addr_QStatusTipEvent_registerEventType2, uintptr(hint))
	__r := int32(uint32(__r0))
	return int(__r)
}

func (this *QStatusTipEvent) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QStatusTipEvent_delete(this.h)
	}
	this.UnsafeRelease()
}
