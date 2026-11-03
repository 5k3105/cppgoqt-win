package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QInputMethodQueryEvent_base func(uintptr) uintptr
	addr_QInputMethodQueryEvent_new uintptr
	addr_QInputMethodQueryEvent_queries uintptr
	addr_QInputMethodQueryEvent_setValue uintptr
	addr_QInputMethodQueryEvent_value uintptr
	addr_QInputMethodQueryEvent_registerEventType uintptr
	addr_QInputMethodQueryEvent_registerEventType2 uintptr
	pg_QInputMethodQueryEvent_delete func(uintptr)
)

func init() {
	reg(&pg_QInputMethodQueryEvent_base, "QInputMethodQueryEvent_base")
	symaddr(&addr_QInputMethodQueryEvent_new, "QInputMethodQueryEvent_new")
	symaddr(&addr_QInputMethodQueryEvent_queries, "QInputMethodQueryEvent_queries")
	symaddr(&addr_QInputMethodQueryEvent_setValue, "QInputMethodQueryEvent_setValue")
	symaddr(&addr_QInputMethodQueryEvent_value, "QInputMethodQueryEvent_value")
	symaddr(&addr_QInputMethodQueryEvent_registerEventType, "QInputMethodQueryEvent_registerEventType")
	symaddr(&addr_QInputMethodQueryEvent_registerEventType2, "QInputMethodQueryEvent_registerEventType2")
	reg(&pg_QInputMethodQueryEvent_delete, "QInputMethodQueryEvent_delete")
}

type QInputMethodQueryEvent struct {
	h uintptr
	*QEvent
}

func bindQInputMethodQueryEvent(h uintptr) *QInputMethodQueryEvent {
	if h == 0 {
		return nil
	}
	return &QInputMethodQueryEvent{h: h, QEvent: bindQEvent(pg_QInputMethodQueryEvent_base(h))}
}

func newownedQInputMethodQueryEvent(h uintptr) *QInputMethodQueryEvent {
	o := bindQInputMethodQueryEvent(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QInputMethodQueryEvent).Delete)
	}
	return o
}

func (this *QInputMethodQueryEvent) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QInputMethodQueryEvent) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQInputMethodQueryEvent(h unsafe.Pointer) *QInputMethodQueryEvent { return bindQInputMethodQueryEvent(uintptr(h)) }
func UnsafeNewQInputMethodQueryEventOwned(h unsafe.Pointer) *QInputMethodQueryEvent { return newownedQInputMethodQueryEvent(uintptr(h)) }

func (this *QInputMethodQueryEvent) UnsafeRelease() {
	this.h = 0
	if this.QEvent != nil {
		this.QEvent.UnsafeRelease()
	}
}

func NewQInputMethodQueryEvent(queries InputMethodQuery) *QInputMethodQueryEvent {
	__r, _, _ := purego.SyscallN(addr_QInputMethodQueryEvent_new, uintptr(queries))
	__o := newownedQInputMethodQueryEvent(__r)
	return __o
}

func (this *QInputMethodQueryEvent) Queries() InputMethodQuery {
	__r0, _, _ := purego.SyscallN(addr_QInputMethodQueryEvent_queries, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return InputMethodQuery(__r)
}

func (this *QInputMethodQueryEvent) SetValue(query InputMethodQuery, value *QVariant) {
	purego.SyscallN(addr_QInputMethodQueryEvent_setValue, this.h, uintptr(query), uintptr(value.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(value)
}

func (this *QInputMethodQueryEvent) Value(query InputMethodQuery) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QInputMethodQueryEvent_value, this.h, uintptr(query))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func QInputMethodQueryEvent_RegisterEventType() int {
	__r0, _, _ := purego.SyscallN(addr_QInputMethodQueryEvent_registerEventType)
	__r := int32(uint32(__r0))
	return int(__r)
}

func QInputMethodQueryEvent_RegisterEventTypeWithHint(hint int) int {
	__r0, _, _ := purego.SyscallN(addr_QInputMethodQueryEvent_registerEventType2, uintptr(hint))
	__r := int32(uint32(__r0))
	return int(__r)
}

func (this *QInputMethodQueryEvent) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QInputMethodQueryEvent_delete(this.h)
	}
	this.UnsafeRelease()
}
