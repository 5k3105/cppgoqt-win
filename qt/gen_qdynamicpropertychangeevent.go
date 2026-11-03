package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QDynamicPropertyChangeEvent_base func(uintptr) uintptr
	addr_QDynamicPropertyChangeEvent_new uintptr
	addr_QDynamicPropertyChangeEvent_propertyName uintptr
	addr_QDynamicPropertyChangeEvent_registerEventType uintptr
	addr_QDynamicPropertyChangeEvent_registerEventType2 uintptr
	pg_QDynamicPropertyChangeEvent_delete func(uintptr)
)

func init() {
	reg(&pg_QDynamicPropertyChangeEvent_base, "QDynamicPropertyChangeEvent_base")
	symaddr(&addr_QDynamicPropertyChangeEvent_new, "QDynamicPropertyChangeEvent_new")
	symaddr(&addr_QDynamicPropertyChangeEvent_propertyName, "QDynamicPropertyChangeEvent_propertyName")
	symaddr(&addr_QDynamicPropertyChangeEvent_registerEventType, "QDynamicPropertyChangeEvent_registerEventType")
	symaddr(&addr_QDynamicPropertyChangeEvent_registerEventType2, "QDynamicPropertyChangeEvent_registerEventType2")
	reg(&pg_QDynamicPropertyChangeEvent_delete, "QDynamicPropertyChangeEvent_delete")
}

type QDynamicPropertyChangeEvent struct {
	h uintptr
	*QEvent
}

func bindQDynamicPropertyChangeEvent(h uintptr) *QDynamicPropertyChangeEvent {
	if h == 0 {
		return nil
	}
	return &QDynamicPropertyChangeEvent{h: h, QEvent: bindQEvent(pg_QDynamicPropertyChangeEvent_base(h))}
}

func newownedQDynamicPropertyChangeEvent(h uintptr) *QDynamicPropertyChangeEvent {
	o := bindQDynamicPropertyChangeEvent(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QDynamicPropertyChangeEvent).Delete)
	}
	return o
}

func (this *QDynamicPropertyChangeEvent) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QDynamicPropertyChangeEvent) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQDynamicPropertyChangeEvent(h unsafe.Pointer) *QDynamicPropertyChangeEvent { return bindQDynamicPropertyChangeEvent(uintptr(h)) }
func UnsafeNewQDynamicPropertyChangeEventOwned(h unsafe.Pointer) *QDynamicPropertyChangeEvent { return newownedQDynamicPropertyChangeEvent(uintptr(h)) }

func (this *QDynamicPropertyChangeEvent) UnsafeRelease() {
	this.h = 0
	if this.QEvent != nil {
		this.QEvent.UnsafeRelease()
	}
}

func NewQDynamicPropertyChangeEvent(name *QByteArray) *QDynamicPropertyChangeEvent {
	__r, _, _ := purego.SyscallN(addr_QDynamicPropertyChangeEvent_new, uintptr(name.UnsafePointer()))
	__o := newownedQDynamicPropertyChangeEvent(__r)
	runtime.KeepAlive(name)
	return __o
}

func (this *QDynamicPropertyChangeEvent) PropertyName() *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QDynamicPropertyChangeEvent_propertyName, this.h)
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func QDynamicPropertyChangeEvent_RegisterEventType() int {
	__r0, _, _ := purego.SyscallN(addr_QDynamicPropertyChangeEvent_registerEventType)
	__r := int32(uint32(__r0))
	return int(__r)
}

func QDynamicPropertyChangeEvent_RegisterEventTypeWithHint(hint int) int {
	__r0, _, _ := purego.SyscallN(addr_QDynamicPropertyChangeEvent_registerEventType2, uintptr(hint))
	__r := int32(uint32(__r0))
	return int(__r)
}

func (this *QDynamicPropertyChangeEvent) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QDynamicPropertyChangeEvent_delete(this.h)
	}
	this.UnsafeRelease()
}
