package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QResizeEvent_base func(uintptr) uintptr
	addr_QResizeEvent_new uintptr
	addr_QResizeEvent_size uintptr
	addr_QResizeEvent_oldSize uintptr
	addr_QResizeEvent_registerEventType uintptr
	addr_QResizeEvent_registerEventType2 uintptr
	pg_QResizeEvent_delete func(uintptr)
)

func init() {
	reg(&pg_QResizeEvent_base, "QResizeEvent_base")
	symaddr(&addr_QResizeEvent_new, "QResizeEvent_new")
	symaddr(&addr_QResizeEvent_size, "QResizeEvent_size")
	symaddr(&addr_QResizeEvent_oldSize, "QResizeEvent_oldSize")
	symaddr(&addr_QResizeEvent_registerEventType, "QResizeEvent_registerEventType")
	symaddr(&addr_QResizeEvent_registerEventType2, "QResizeEvent_registerEventType2")
	reg(&pg_QResizeEvent_delete, "QResizeEvent_delete")
}

type QResizeEvent struct {
	h uintptr
	*QEvent
}

func bindQResizeEvent(h uintptr) *QResizeEvent {
	if h == 0 {
		return nil
	}
	return &QResizeEvent{h: h, QEvent: bindQEvent(pg_QResizeEvent_base(h))}
}

func newownedQResizeEvent(h uintptr) *QResizeEvent {
	o := bindQResizeEvent(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QResizeEvent).Delete)
	}
	return o
}

func (this *QResizeEvent) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QResizeEvent) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQResizeEvent(h unsafe.Pointer) *QResizeEvent { return bindQResizeEvent(uintptr(h)) }
func UnsafeNewQResizeEventOwned(h unsafe.Pointer) *QResizeEvent { return newownedQResizeEvent(uintptr(h)) }

func (this *QResizeEvent) UnsafeRelease() {
	this.h = 0
	if this.QEvent != nil {
		this.QEvent.UnsafeRelease()
	}
}

func NewQResizeEvent(size *QSize, oldSize *QSize) *QResizeEvent {
	__r, _, _ := purego.SyscallN(addr_QResizeEvent_new, uintptr(size.UnsafePointer()), uintptr(oldSize.UnsafePointer()))
	__o := newownedQResizeEvent(__r)
	runtime.KeepAlive(size)
	runtime.KeepAlive(oldSize)
	return __o
}

func (this *QResizeEvent) Size() *QSize {
	__r, _, _ := purego.SyscallN(addr_QResizeEvent_size, this.h)
	runtime.KeepAlive(this)
	return bindQSize(__r)
}

func (this *QResizeEvent) OldSize() *QSize {
	__r, _, _ := purego.SyscallN(addr_QResizeEvent_oldSize, this.h)
	runtime.KeepAlive(this)
	return bindQSize(__r)
}

func QResizeEvent_RegisterEventType() int {
	__r0, _, _ := purego.SyscallN(addr_QResizeEvent_registerEventType)
	__r := int32(uint32(__r0))
	return int(__r)
}

func QResizeEvent_RegisterEventTypeWithHint(hint int) int {
	__r0, _, _ := purego.SyscallN(addr_QResizeEvent_registerEventType2, uintptr(hint))
	__r := int32(uint32(__r0))
	return int(__r)
}

func (this *QResizeEvent) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QResizeEvent_delete(this.h)
	}
	this.UnsafeRelease()
}
