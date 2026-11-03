package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QAccessibleBridge_setRootObject uintptr
	addr_QAccessibleBridge_notifyAccessibilityUpdate uintptr
	addr_QAccessibleBridge_OperatorAssign uintptr
	pg_QAccessibleBridge__set_override func(uintptr, int32)
	pg_QAccessibleBridge__Go_new func(int64) uintptr
	pg_QAccessibleBridge_delete func(uintptr)
)

func init() {
	symaddr(&addr_QAccessibleBridge_setRootObject, "QAccessibleBridge_setRootObject")
	symaddr(&addr_QAccessibleBridge_notifyAccessibilityUpdate, "QAccessibleBridge_notifyAccessibilityUpdate")
	symaddr(&addr_QAccessibleBridge_OperatorAssign, "QAccessibleBridge_OperatorAssign")
	reg(&pg_QAccessibleBridge__set_override, "QAccessibleBridge__set_override")
	reg(&pg_QAccessibleBridge__Go_new, "QAccessibleBridge__Go_new")
	reg(&pg_QAccessibleBridge_delete, "QAccessibleBridge_delete")
}

type QAccessibleBridge struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QAccessibleBridge) UnsafeGoID() int64     { return this.vid }
func (this *QAccessibleBridge) UnsafeSetGoID(id int64) { this.vid = id }

func bindQAccessibleBridge(h uintptr) *QAccessibleBridge {
	if h == 0 {
		return nil
	}
	return &QAccessibleBridge{h: h}
}

func newownedQAccessibleBridge(h uintptr) *QAccessibleBridge {
	o := bindQAccessibleBridge(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QAccessibleBridge).Delete)
	}
	return o
}

func (this *QAccessibleBridge) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QAccessibleBridge) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQAccessibleBridge(h unsafe.Pointer) *QAccessibleBridge { return bindQAccessibleBridge(uintptr(h)) }
func UnsafeNewQAccessibleBridgeOwned(h unsafe.Pointer) *QAccessibleBridge { return newownedQAccessibleBridge(uintptr(h)) }

func (this *QAccessibleBridge) UnsafeRelease() { this.h = 0 }

func (this *QAccessibleBridge) SetRootObject(arg1 *QAccessibleInterface) {
	purego.SyscallN(addr_QAccessibleBridge_setRootObject, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QAccessibleBridge) NotifyAccessibilityUpdate(event *QAccessibleEvent) {
	purego.SyscallN(addr_QAccessibleBridge_notifyAccessibilityUpdate, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAccessibleBridge) OperatorAssign(arg1 *QAccessibleBridge) *QAccessibleBridge {
	__r, _, _ := purego.SyscallN(addr_QAccessibleBridge_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQAccessibleBridge(__r)
}

func QAccessibleBridge_virtIdx(slot int64) int32 {
	switch slot {
	case 7245844568460288668:
		return 0
	case -4087650821821605459:
		return 1
	}
	return -1
}

func QAccessibleBridge_enable(h uintptr, idx int32) { pg_QAccessibleBridge__set_override(h, idx) }

func NewQAccessibleBridge() *QAccessibleBridge {
	id := nextVirtID()
	o := newownedQAccessibleBridge(pg_QAccessibleBridge__Go_new(id))
	o.UnsafeSetGoID(id)
	registerInstance(id, o.h, QAccessibleBridge_virtIdx, QAccessibleBridge_enable)
	return o
}

func (this *QAccessibleBridge) OnSetRootObject(fn func(*QAccessibleInterface)) {
	registerVirt(this.UnsafeGoID(), 7245844568460288668, func(arg1 uintptr) { fn(UnsafeNewQAccessibleInterface(unsafe.Pointer(arg1))) }) // slot "setRootObject"
}

func (this *QAccessibleBridge) OnNotifyAccessibilityUpdate(fn func(*QAccessibleEvent)) {
	registerVirt(this.UnsafeGoID(), -4087650821821605459, func(event uintptr) { fn(UnsafeNewQAccessibleEvent(unsafe.Pointer(event))) }) // slot "notifyAccessibilityUpdate"
}

func (this *QAccessibleBridge) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QAccessibleBridge_delete(this.h)
	}
	this.UnsafeRelease()
}
