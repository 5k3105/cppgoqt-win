package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QWhatsThisClickedEvent_base func(uintptr) uintptr
	addr_QWhatsThisClickedEvent_new uintptr
	addr_QWhatsThisClickedEvent_href uintptr
	addr_QWhatsThisClickedEvent_registerEventType uintptr
	addr_QWhatsThisClickedEvent_registerEventType2 uintptr
	pg_QWhatsThisClickedEvent_delete func(uintptr)
)

func init() {
	reg(&pg_QWhatsThisClickedEvent_base, "QWhatsThisClickedEvent_base")
	symaddr(&addr_QWhatsThisClickedEvent_new, "QWhatsThisClickedEvent_new")
	symaddr(&addr_QWhatsThisClickedEvent_href, "QWhatsThisClickedEvent_href")
	symaddr(&addr_QWhatsThisClickedEvent_registerEventType, "QWhatsThisClickedEvent_registerEventType")
	symaddr(&addr_QWhatsThisClickedEvent_registerEventType2, "QWhatsThisClickedEvent_registerEventType2")
	reg(&pg_QWhatsThisClickedEvent_delete, "QWhatsThisClickedEvent_delete")
}

type QWhatsThisClickedEvent struct {
	h uintptr
	*QEvent
}

func bindQWhatsThisClickedEvent(h uintptr) *QWhatsThisClickedEvent {
	if h == 0 {
		return nil
	}
	return &QWhatsThisClickedEvent{h: h, QEvent: bindQEvent(pg_QWhatsThisClickedEvent_base(h))}
}

func newownedQWhatsThisClickedEvent(h uintptr) *QWhatsThisClickedEvent {
	o := bindQWhatsThisClickedEvent(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QWhatsThisClickedEvent).Delete)
	}
	return o
}

func (this *QWhatsThisClickedEvent) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QWhatsThisClickedEvent) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQWhatsThisClickedEvent(h unsafe.Pointer) *QWhatsThisClickedEvent { return bindQWhatsThisClickedEvent(uintptr(h)) }
func UnsafeNewQWhatsThisClickedEventOwned(h unsafe.Pointer) *QWhatsThisClickedEvent { return newownedQWhatsThisClickedEvent(uintptr(h)) }

func (this *QWhatsThisClickedEvent) UnsafeRelease() {
	this.h = 0
	if this.QEvent != nil {
		this.QEvent.UnsafeRelease()
	}
}

func NewQWhatsThisClickedEvent(href string) *QWhatsThisClickedEvent {
	var href_p uintptr
	if len(href) > 0 {
		href_p = uintptr(unsafe.Pointer(unsafe.StringData(href)))
	}
	__r, _, _ := purego.SyscallN(addr_QWhatsThisClickedEvent_new, href_p, uintptr(len(href)))
	__o := newownedQWhatsThisClickedEvent(__r)
	runtime.KeepAlive(href)
	return __o
}

func (this *QWhatsThisClickedEvent) Href() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QWhatsThisClickedEvent_href, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func QWhatsThisClickedEvent_RegisterEventType() int {
	__r0, _, _ := purego.SyscallN(addr_QWhatsThisClickedEvent_registerEventType)
	__r := int32(uint32(__r0))
	return int(__r)
}

func QWhatsThisClickedEvent_RegisterEventTypeWithHint(hint int) int {
	__r0, _, _ := purego.SyscallN(addr_QWhatsThisClickedEvent_registerEventType2, uintptr(hint))
	__r := int32(uint32(__r0))
	return int(__r)
}

func (this *QWhatsThisClickedEvent) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QWhatsThisClickedEvent_delete(this.h)
	}
	this.UnsafeRelease()
}
