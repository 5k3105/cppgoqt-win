package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QGraphicsSceneEvent_base func(uintptr) uintptr
	addr_QGraphicsSceneEvent_new uintptr
	addr_QGraphicsSceneEvent_widget uintptr
	addr_QGraphicsSceneEvent_setWidget uintptr
	addr_QGraphicsSceneEvent_timestamp uintptr
	addr_QGraphicsSceneEvent_setTimestamp uintptr
	addr_QGraphicsSceneEvent_registerEventType uintptr
	addr_QGraphicsSceneEvent_registerEventType2 uintptr
	pg_QGraphicsSceneEvent_delete func(uintptr)
)

func init() {
	reg(&pg_QGraphicsSceneEvent_base, "QGraphicsSceneEvent_base")
	symaddr(&addr_QGraphicsSceneEvent_new, "QGraphicsSceneEvent_new")
	symaddr(&addr_QGraphicsSceneEvent_widget, "QGraphicsSceneEvent_widget")
	symaddr(&addr_QGraphicsSceneEvent_setWidget, "QGraphicsSceneEvent_setWidget")
	symaddr(&addr_QGraphicsSceneEvent_timestamp, "QGraphicsSceneEvent_timestamp")
	symaddr(&addr_QGraphicsSceneEvent_setTimestamp, "QGraphicsSceneEvent_setTimestamp")
	symaddr(&addr_QGraphicsSceneEvent_registerEventType, "QGraphicsSceneEvent_registerEventType")
	symaddr(&addr_QGraphicsSceneEvent_registerEventType2, "QGraphicsSceneEvent_registerEventType2")
	reg(&pg_QGraphicsSceneEvent_delete, "QGraphicsSceneEvent_delete")
}

type QGraphicsSceneEvent struct {
	h uintptr
	*QEvent
}

func bindQGraphicsSceneEvent(h uintptr) *QGraphicsSceneEvent {
	if h == 0 {
		return nil
	}
	return &QGraphicsSceneEvent{h: h, QEvent: bindQEvent(pg_QGraphicsSceneEvent_base(h))}
}

func newownedQGraphicsSceneEvent(h uintptr) *QGraphicsSceneEvent {
	o := bindQGraphicsSceneEvent(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QGraphicsSceneEvent).Delete)
	}
	return o
}

func (this *QGraphicsSceneEvent) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QGraphicsSceneEvent) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQGraphicsSceneEvent(h unsafe.Pointer) *QGraphicsSceneEvent { return bindQGraphicsSceneEvent(uintptr(h)) }
func UnsafeNewQGraphicsSceneEventOwned(h unsafe.Pointer) *QGraphicsSceneEvent { return newownedQGraphicsSceneEvent(uintptr(h)) }

func (this *QGraphicsSceneEvent) UnsafeRelease() {
	this.h = 0
	if this.QEvent != nil {
		this.QEvent.UnsafeRelease()
	}
}

func NewQGraphicsSceneEvent(type_ QEvent__Type) *QGraphicsSceneEvent {
	__r, _, _ := purego.SyscallN(addr_QGraphicsSceneEvent_new, uintptr(type_))
	__o := newownedQGraphicsSceneEvent(__r)
	return __o
}

func (this *QGraphicsSceneEvent) Widget() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QGraphicsSceneEvent_widget, this.h)
	runtime.KeepAlive(this)
	return bindQWidget(__r)
}

func (this *QGraphicsSceneEvent) SetWidget(widget *QWidget) {
	purego.SyscallN(addr_QGraphicsSceneEvent_setWidget, this.h, uintptr(widget.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(widget)
}

func (this *QGraphicsSceneEvent) Timestamp() uint64 {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsSceneEvent_timestamp, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return uint64(__r)
}

func (this *QGraphicsSceneEvent) SetTimestamp(ts uint64) {
	purego.SyscallN(addr_QGraphicsSceneEvent_setTimestamp, this.h, uintptr(ts))
	runtime.KeepAlive(this)
}

func QGraphicsSceneEvent_RegisterEventType() int {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsSceneEvent_registerEventType)
	__r := int32(uint32(__r0))
	return int(__r)
}

func QGraphicsSceneEvent_RegisterEventTypeWithHint(hint int) int {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsSceneEvent_registerEventType2, uintptr(hint))
	__r := int32(uint32(__r0))
	return int(__r)
}

func (this *QGraphicsSceneEvent) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QGraphicsSceneEvent_delete(this.h)
	}
	this.UnsafeRelease()
}
