package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QChildEvent_base func(uintptr) uintptr
	addr_QChildEvent_new uintptr
	addr_QChildEvent_child uintptr
	addr_QChildEvent_added uintptr
	addr_QChildEvent_polished uintptr
	addr_QChildEvent_removed uintptr
	addr_QChildEvent_registerEventType uintptr
	addr_QChildEvent_registerEventType2 uintptr
	pg_QChildEvent_delete func(uintptr)
)

func init() {
	reg(&pg_QChildEvent_base, "QChildEvent_base")
	symaddr(&addr_QChildEvent_new, "QChildEvent_new")
	symaddr(&addr_QChildEvent_child, "QChildEvent_child")
	symaddr(&addr_QChildEvent_added, "QChildEvent_added")
	symaddr(&addr_QChildEvent_polished, "QChildEvent_polished")
	symaddr(&addr_QChildEvent_removed, "QChildEvent_removed")
	symaddr(&addr_QChildEvent_registerEventType, "QChildEvent_registerEventType")
	symaddr(&addr_QChildEvent_registerEventType2, "QChildEvent_registerEventType2")
	reg(&pg_QChildEvent_delete, "QChildEvent_delete")
}

type QChildEvent struct {
	h uintptr
	*QEvent
}

func bindQChildEvent(h uintptr) *QChildEvent {
	if h == 0 {
		return nil
	}
	return &QChildEvent{h: h, QEvent: bindQEvent(pg_QChildEvent_base(h))}
}

func newownedQChildEvent(h uintptr) *QChildEvent {
	o := bindQChildEvent(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QChildEvent).Delete)
	}
	return o
}

func (this *QChildEvent) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QChildEvent) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQChildEvent(h unsafe.Pointer) *QChildEvent { return bindQChildEvent(uintptr(h)) }
func UnsafeNewQChildEventOwned(h unsafe.Pointer) *QChildEvent { return newownedQChildEvent(uintptr(h)) }

func (this *QChildEvent) UnsafeRelease() {
	this.h = 0
	if this.QEvent != nil {
		this.QEvent.UnsafeRelease()
	}
}

func NewQChildEvent(type_ QEvent__Type, child *QObject) *QChildEvent {
	__r, _, _ := purego.SyscallN(addr_QChildEvent_new, uintptr(type_), uintptr(child.UnsafePointer()))
	__o := newownedQChildEvent(__r)
	runtime.KeepAlive(child)
	return __o
}

func (this *QChildEvent) Child() *QObject {
	__r, _, _ := purego.SyscallN(addr_QChildEvent_child, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QChildEvent) Added() bool {
	__r0, _, _ := purego.SyscallN(addr_QChildEvent_added, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QChildEvent) Polished() bool {
	__r0, _, _ := purego.SyscallN(addr_QChildEvent_polished, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QChildEvent) Removed() bool {
	__r0, _, _ := purego.SyscallN(addr_QChildEvent_removed, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QChildEvent_RegisterEventType() int {
	__r0, _, _ := purego.SyscallN(addr_QChildEvent_registerEventType)
	__r := int32(uint32(__r0))
	return int(__r)
}

func QChildEvent_RegisterEventTypeWithHint(hint int) int {
	__r0, _, _ := purego.SyscallN(addr_QChildEvent_registerEventType2, uintptr(hint))
	__r := int32(uint32(__r0))
	return int(__r)
}

func (this *QChildEvent) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QChildEvent_delete(this.h)
	}
	this.UnsafeRelease()
}
