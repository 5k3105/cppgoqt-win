package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QEvent_new uintptr
	addr_QEvent_type uintptr
	addr_QEvent_spontaneous uintptr
	addr_QEvent_setAccepted uintptr
	addr_QEvent_isAccepted uintptr
	addr_QEvent_accept uintptr
	addr_QEvent_ignore uintptr
	addr_QEvent_isInputEvent uintptr
	addr_QEvent_isPointerEvent uintptr
	addr_QEvent_isSinglePointEvent uintptr
	addr_QEvent_registerEventType uintptr
	addr_QEvent_clone uintptr
	addr_QEvent_registerEventType2 uintptr
	pg_QEvent_delete func(uintptr)
)

func init() {
	symaddr(&addr_QEvent_new, "QEvent_new")
	symaddr(&addr_QEvent_type, "QEvent_type")
	symaddr(&addr_QEvent_spontaneous, "QEvent_spontaneous")
	symaddr(&addr_QEvent_setAccepted, "QEvent_setAccepted")
	symaddr(&addr_QEvent_isAccepted, "QEvent_isAccepted")
	symaddr(&addr_QEvent_accept, "QEvent_accept")
	symaddr(&addr_QEvent_ignore, "QEvent_ignore")
	symaddr(&addr_QEvent_isInputEvent, "QEvent_isInputEvent")
	symaddr(&addr_QEvent_isPointerEvent, "QEvent_isPointerEvent")
	symaddr(&addr_QEvent_isSinglePointEvent, "QEvent_isSinglePointEvent")
	symaddr(&addr_QEvent_registerEventType, "QEvent_registerEventType")
	symaddr(&addr_QEvent_clone, "QEvent_clone")
	symaddr(&addr_QEvent_registerEventType2, "QEvent_registerEventType2")
	reg(&pg_QEvent_delete, "QEvent_delete")
}

type QEvent struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QEvent) UnsafeGoID() int64     { return this.vid }
func (this *QEvent) UnsafeSetGoID(id int64) { this.vid = id }

func bindQEvent(h uintptr) *QEvent {
	if h == 0 {
		return nil
	}
	return &QEvent{h: h}
}

func newownedQEvent(h uintptr) *QEvent {
	o := bindQEvent(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QEvent).Delete)
	}
	return o
}

func (this *QEvent) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QEvent) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQEvent(h unsafe.Pointer) *QEvent { return bindQEvent(uintptr(h)) }
func UnsafeNewQEventOwned(h unsafe.Pointer) *QEvent { return newownedQEvent(uintptr(h)) }

func (this *QEvent) UnsafeRelease() { this.h = 0 }

func NewQEvent(type_ QEvent__Type) *QEvent {
	__r, _, _ := purego.SyscallN(addr_QEvent_new, uintptr(type_))
	__o := newownedQEvent(__r)
	return __o
}

func (this *QEvent) Type() QEvent__Type {
	__r0, _, _ := purego.SyscallN(addr_QEvent_type, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QEvent__Type(__r)
}

func (this *QEvent) Spontaneous() bool {
	__r0, _, _ := purego.SyscallN(addr_QEvent_spontaneous, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QEvent) SetAccepted(accepted bool) {
	purego.SyscallN(addr_QEvent_setAccepted, this.h, uintptr(b2i(accepted)))
	runtime.KeepAlive(this)
}

func (this *QEvent) IsAccepted() bool {
	__r0, _, _ := purego.SyscallN(addr_QEvent_isAccepted, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QEvent) Accept() {
	purego.SyscallN(addr_QEvent_accept, this.h)
	runtime.KeepAlive(this)
}

func (this *QEvent) Ignore() {
	purego.SyscallN(addr_QEvent_ignore, this.h)
	runtime.KeepAlive(this)
}

func (this *QEvent) IsInputEvent() bool {
	__r0, _, _ := purego.SyscallN(addr_QEvent_isInputEvent, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QEvent) IsPointerEvent() bool {
	__r0, _, _ := purego.SyscallN(addr_QEvent_isPointerEvent, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QEvent) IsSinglePointEvent() bool {
	__r0, _, _ := purego.SyscallN(addr_QEvent_isSinglePointEvent, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QEvent_RegisterEventType() int {
	__r0, _, _ := purego.SyscallN(addr_QEvent_registerEventType)
	__r := int32(uint32(__r0))
	return int(__r)
}

func (this *QEvent) Clone() *QEvent {
	__r, _, _ := purego.SyscallN(addr_QEvent_clone, this.h)
	runtime.KeepAlive(this)
	return bindQEvent(__r)
}

func QEvent_RegisterEventTypeWithHint(hint int) int {
	__r0, _, _ := purego.SyscallN(addr_QEvent_registerEventType2, uintptr(hint))
	__r := int32(uint32(__r0))
	return int(__r)
}

func (this *QEvent) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QEvent_delete(this.h)
	}
	this.UnsafeRelease()
}
