package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QScrollPrepareEvent_base func(uintptr) uintptr
	addr_QScrollPrepareEvent_new uintptr
	addr_QScrollPrepareEvent_startPos uintptr
	addr_QScrollPrepareEvent_viewportSize uintptr
	addr_QScrollPrepareEvent_contentPosRange uintptr
	addr_QScrollPrepareEvent_contentPos uintptr
	addr_QScrollPrepareEvent_setViewportSize uintptr
	addr_QScrollPrepareEvent_setContentPosRange uintptr
	addr_QScrollPrepareEvent_setContentPos uintptr
	addr_QScrollPrepareEvent_registerEventType uintptr
	addr_QScrollPrepareEvent_registerEventType2 uintptr
	pg_QScrollPrepareEvent_delete func(uintptr)
)

func init() {
	reg(&pg_QScrollPrepareEvent_base, "QScrollPrepareEvent_base")
	symaddr(&addr_QScrollPrepareEvent_new, "QScrollPrepareEvent_new")
	symaddr(&addr_QScrollPrepareEvent_startPos, "QScrollPrepareEvent_startPos")
	symaddr(&addr_QScrollPrepareEvent_viewportSize, "QScrollPrepareEvent_viewportSize")
	symaddr(&addr_QScrollPrepareEvent_contentPosRange, "QScrollPrepareEvent_contentPosRange")
	symaddr(&addr_QScrollPrepareEvent_contentPos, "QScrollPrepareEvent_contentPos")
	symaddr(&addr_QScrollPrepareEvent_setViewportSize, "QScrollPrepareEvent_setViewportSize")
	symaddr(&addr_QScrollPrepareEvent_setContentPosRange, "QScrollPrepareEvent_setContentPosRange")
	symaddr(&addr_QScrollPrepareEvent_setContentPos, "QScrollPrepareEvent_setContentPos")
	symaddr(&addr_QScrollPrepareEvent_registerEventType, "QScrollPrepareEvent_registerEventType")
	symaddr(&addr_QScrollPrepareEvent_registerEventType2, "QScrollPrepareEvent_registerEventType2")
	reg(&pg_QScrollPrepareEvent_delete, "QScrollPrepareEvent_delete")
}

type QScrollPrepareEvent struct {
	h uintptr
	*QEvent
}

func bindQScrollPrepareEvent(h uintptr) *QScrollPrepareEvent {
	if h == 0 {
		return nil
	}
	return &QScrollPrepareEvent{h: h, QEvent: bindQEvent(pg_QScrollPrepareEvent_base(h))}
}

func newownedQScrollPrepareEvent(h uintptr) *QScrollPrepareEvent {
	o := bindQScrollPrepareEvent(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QScrollPrepareEvent).Delete)
	}
	return o
}

func (this *QScrollPrepareEvent) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QScrollPrepareEvent) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQScrollPrepareEvent(h unsafe.Pointer) *QScrollPrepareEvent { return bindQScrollPrepareEvent(uintptr(h)) }
func UnsafeNewQScrollPrepareEventOwned(h unsafe.Pointer) *QScrollPrepareEvent { return newownedQScrollPrepareEvent(uintptr(h)) }

func (this *QScrollPrepareEvent) UnsafeRelease() {
	this.h = 0
	if this.QEvent != nil {
		this.QEvent.UnsafeRelease()
	}
}

func NewQScrollPrepareEvent(startPos *QPointF) *QScrollPrepareEvent {
	__r, _, _ := purego.SyscallN(addr_QScrollPrepareEvent_new, uintptr(startPos.UnsafePointer()))
	__o := newownedQScrollPrepareEvent(__r)
	runtime.KeepAlive(startPos)
	return __o
}

func (this *QScrollPrepareEvent) StartPos() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QScrollPrepareEvent_startPos, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QScrollPrepareEvent) ViewportSize() *QSizeF {
	__r, _, _ := purego.SyscallN(addr_QScrollPrepareEvent_viewportSize, this.h)
	runtime.KeepAlive(this)
	return newownedQSizeF(__r)
}

func (this *QScrollPrepareEvent) ContentPosRange() *QRectF {
	__r, _, _ := purego.SyscallN(addr_QScrollPrepareEvent_contentPosRange, this.h)
	runtime.KeepAlive(this)
	return newownedQRectF(__r)
}

func (this *QScrollPrepareEvent) ContentPos() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QScrollPrepareEvent_contentPos, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QScrollPrepareEvent) SetViewportSize(size *QSizeF) {
	purego.SyscallN(addr_QScrollPrepareEvent_setViewportSize, this.h, uintptr(size.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(size)
}

func (this *QScrollPrepareEvent) SetContentPosRange(rect *QRectF) {
	purego.SyscallN(addr_QScrollPrepareEvent_setContentPosRange, this.h, uintptr(rect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
}

func (this *QScrollPrepareEvent) SetContentPos(pos *QPointF) {
	purego.SyscallN(addr_QScrollPrepareEvent_setContentPos, this.h, uintptr(pos.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pos)
}

func QScrollPrepareEvent_RegisterEventType() int {
	__r0, _, _ := purego.SyscallN(addr_QScrollPrepareEvent_registerEventType)
	__r := int32(uint32(__r0))
	return int(__r)
}

func QScrollPrepareEvent_RegisterEventTypeWithHint(hint int) int {
	__r0, _, _ := purego.SyscallN(addr_QScrollPrepareEvent_registerEventType2, uintptr(hint))
	__r := int32(uint32(__r0))
	return int(__r)
}

func (this *QScrollPrepareEvent) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QScrollPrepareEvent_delete(this.h)
	}
	this.UnsafeRelease()
}
