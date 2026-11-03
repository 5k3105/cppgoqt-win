package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QAccessible__ActivationObserver_accessibilityActiveChanged uintptr
	addr_QAccessible__ActivationObserver_OperatorAssign uintptr
	pg_QAccessible__ActivationObserver__set_override func(uintptr, int32)
	pg_QAccessible__ActivationObserver__Go_new func(int64) uintptr
	pg_QAccessible__ActivationObserver_delete func(uintptr)
)

func init() {
	symaddr(&addr_QAccessible__ActivationObserver_accessibilityActiveChanged, "QAccessible__ActivationObserver_accessibilityActiveChanged")
	symaddr(&addr_QAccessible__ActivationObserver_OperatorAssign, "QAccessible__ActivationObserver_OperatorAssign")
	reg(&pg_QAccessible__ActivationObserver__set_override, "QAccessible__ActivationObserver__set_override")
	reg(&pg_QAccessible__ActivationObserver__Go_new, "QAccessible__ActivationObserver__Go_new")
	reg(&pg_QAccessible__ActivationObserver_delete, "QAccessible__ActivationObserver_delete")
}

type QAccessible__ActivationObserver struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QAccessible__ActivationObserver) UnsafeGoID() int64     { return this.vid }
func (this *QAccessible__ActivationObserver) UnsafeSetGoID(id int64) { this.vid = id }

func bindQAccessible__ActivationObserver(h uintptr) *QAccessible__ActivationObserver {
	if h == 0 {
		return nil
	}
	return &QAccessible__ActivationObserver{h: h}
}

func newownedQAccessible__ActivationObserver(h uintptr) *QAccessible__ActivationObserver {
	o := bindQAccessible__ActivationObserver(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QAccessible__ActivationObserver).Delete)
	}
	return o
}

func (this *QAccessible__ActivationObserver) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QAccessible__ActivationObserver) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQAccessible__ActivationObserver(h unsafe.Pointer) *QAccessible__ActivationObserver { return bindQAccessible__ActivationObserver(uintptr(h)) }
func UnsafeNewQAccessible__ActivationObserverOwned(h unsafe.Pointer) *QAccessible__ActivationObserver { return newownedQAccessible__ActivationObserver(uintptr(h)) }

func (this *QAccessible__ActivationObserver) UnsafeRelease() { this.h = 0 }

func (this *QAccessible__ActivationObserver) AccessibilityActiveChanged(active bool) {
	purego.SyscallN(addr_QAccessible__ActivationObserver_accessibilityActiveChanged, this.h, uintptr(b2i(active)))
	runtime.KeepAlive(this)
}

func (this *QAccessible__ActivationObserver) OperatorAssign(arg1 *QAccessible__ActivationObserver) *QAccessible__ActivationObserver {
	__r, _, _ := purego.SyscallN(addr_QAccessible__ActivationObserver_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQAccessible__ActivationObserver(__r)
}

func QAccessible__ActivationObserver_virtIdx(slot int64) int32 {
	switch slot {
	case 5505922000441301119:
		return 0
	}
	return -1
}

func QAccessible__ActivationObserver_enable(h uintptr, idx int32) { pg_QAccessible__ActivationObserver__set_override(h, idx) }

func NewQAccessible__ActivationObserver() *QAccessible__ActivationObserver {
	id := nextVirtID()
	o := newownedQAccessible__ActivationObserver(pg_QAccessible__ActivationObserver__Go_new(id))
	o.UnsafeSetGoID(id)
	registerInstance(id, o.h, QAccessible__ActivationObserver_virtIdx, QAccessible__ActivationObserver_enable)
	return o
}

func (this *QAccessible__ActivationObserver) OnAccessibilityActiveChanged(fn func(bool)) {
	registerVirt(this.UnsafeGoID(), 5505922000441301119, func(active int32) { fn(active != 0) }) // slot "accessibilityActiveChanged"
}

func (this *QAccessible__ActivationObserver) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QAccessible__ActivationObserver_delete(this.h)
	}
	this.UnsafeRelease()
}
