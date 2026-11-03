package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QGestureRecognizer_create uintptr
	addr_QGestureRecognizer_recognize uintptr
	addr_QGestureRecognizer_reset uintptr
	addr_QGestureRecognizer_registerRecognizer uintptr
	addr_QGestureRecognizer_unregisterRecognizer uintptr
	addr_QGestureRecognizer_OperatorAssign uintptr
	addr_QGestureRecognizer_create_default uintptr
	addr_QGestureRecognizer_reset_default uintptr
	pg_QGestureRecognizer__set_override func(uintptr, int32)
	pg_QGestureRecognizer__Go_new func(int64) uintptr
	pg_QGestureRecognizer_delete func(uintptr)
)

func init() {
	symaddr(&addr_QGestureRecognizer_create, "QGestureRecognizer_create")
	symaddr(&addr_QGestureRecognizer_recognize, "QGestureRecognizer_recognize")
	symaddr(&addr_QGestureRecognizer_reset, "QGestureRecognizer_reset")
	symaddr(&addr_QGestureRecognizer_registerRecognizer, "QGestureRecognizer_registerRecognizer")
	symaddr(&addr_QGestureRecognizer_unregisterRecognizer, "QGestureRecognizer_unregisterRecognizer")
	symaddr(&addr_QGestureRecognizer_OperatorAssign, "QGestureRecognizer_OperatorAssign")
	symaddr(&addr_QGestureRecognizer_create_default, "QGestureRecognizer_create_default")
	symaddr(&addr_QGestureRecognizer_reset_default, "QGestureRecognizer_reset_default")
	reg(&pg_QGestureRecognizer__set_override, "QGestureRecognizer__set_override")
	reg(&pg_QGestureRecognizer__Go_new, "QGestureRecognizer__Go_new")
	reg(&pg_QGestureRecognizer_delete, "QGestureRecognizer_delete")
}

type QGestureRecognizer struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QGestureRecognizer) UnsafeGoID() int64     { return this.vid }
func (this *QGestureRecognizer) UnsafeSetGoID(id int64) { this.vid = id }

func bindQGestureRecognizer(h uintptr) *QGestureRecognizer {
	if h == 0 {
		return nil
	}
	return &QGestureRecognizer{h: h}
}

func newownedQGestureRecognizer(h uintptr) *QGestureRecognizer {
	o := bindQGestureRecognizer(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QGestureRecognizer).Delete)
	}
	return o
}

func (this *QGestureRecognizer) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QGestureRecognizer) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQGestureRecognizer(h unsafe.Pointer) *QGestureRecognizer { return bindQGestureRecognizer(uintptr(h)) }
func UnsafeNewQGestureRecognizerOwned(h unsafe.Pointer) *QGestureRecognizer { return newownedQGestureRecognizer(uintptr(h)) }

func (this *QGestureRecognizer) UnsafeRelease() { this.h = 0 }

func (this *QGestureRecognizer) Create(target *QObject) *QGesture {
	__r, _, _ := purego.SyscallN(addr_QGestureRecognizer_create, this.h, uintptr(target.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(target)
	return bindQGesture(__r)
}

func (this *QGestureRecognizer) Recognize(state *QGesture, watched *QObject, event *QEvent) QGestureRecognizer__ResultFlag {
	__r0, _, _ := purego.SyscallN(addr_QGestureRecognizer_recognize, this.h, uintptr(state.UnsafePointer()), uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(state)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return QGestureRecognizer__ResultFlag(__r)
}

func (this *QGestureRecognizer) Reset(state *QGesture) {
	purego.SyscallN(addr_QGestureRecognizer_reset, this.h, uintptr(state.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(state)
}

func QGestureRecognizer_RegisterRecognizer(recognizer *QGestureRecognizer) GestureType {
	__r0, _, _ := purego.SyscallN(addr_QGestureRecognizer_registerRecognizer, uintptr(recognizer.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(recognizer)
	return GestureType(__r)
}

func QGestureRecognizer_UnregisterRecognizer(type_ GestureType) {
	purego.SyscallN(addr_QGestureRecognizer_unregisterRecognizer, uintptr(type_))
}

func (this *QGestureRecognizer) OperatorAssign(arg1 *QGestureRecognizer) *QGestureRecognizer {
	__r, _, _ := purego.SyscallN(addr_QGestureRecognizer_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQGestureRecognizer(__r)
}

func (this *QGestureRecognizer) CreateDefault(target *QObject) *QGesture {
	__r, _, _ := purego.SyscallN(addr_QGestureRecognizer_create_default, this.h, uintptr(target.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(target)
	return bindQGesture(__r)
}

func (this *QGestureRecognizer) ResetDefault(state *QGesture) {
	purego.SyscallN(addr_QGestureRecognizer_reset_default, this.h, uintptr(state.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(state)
}

func QGestureRecognizer_virtIdx(slot int64) int32 {
	switch slot {
	case -8783688087013084677:
		return 0
	case -626457835795418115:
		return 1
	case 1086335023529244512:
		return 2
	}
	return -1
}

func QGestureRecognizer_enable(h uintptr, idx int32) { pg_QGestureRecognizer__set_override(h, idx) }

func NewQGestureRecognizer() *QGestureRecognizer {
	id := nextVirtID()
	o := newownedQGestureRecognizer(pg_QGestureRecognizer__Go_new(id))
	o.UnsafeSetGoID(id)
	registerInstance(id, o.h, QGestureRecognizer_virtIdx, QGestureRecognizer_enable)
	return o
}

func (this *QGestureRecognizer) OnRecognize(fn func(*QGesture, *QObject, *QEvent) QGestureRecognizer__ResultFlag) {
	registerVirt(this.UnsafeGoID(), -8783688087013084677, func(state uintptr, watched uintptr, event uintptr) int32 { return int32(fn(UnsafeNewQGesture(unsafe.Pointer(state)), UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "recognize"
}

func (this *QGestureRecognizer) OnCreate(fn func(*QObject) *QGesture) {
	registerVirt(this.UnsafeGoID(), -626457835795418115, func(target uintptr) uintptr { return uintptr(fn(UnsafeNewQObject(unsafe.Pointer(target))).UnsafePointer()) }) // slot "create"
}

func (this *QGestureRecognizer) OnReset(fn func(*QGesture)) {
	registerVirt(this.UnsafeGoID(), 1086335023529244512, func(state uintptr) { fn(UnsafeNewQGesture(unsafe.Pointer(state))) }) // slot "reset"
}

func (this *QGestureRecognizer) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QGestureRecognizer_delete(this.h)
	}
	this.UnsafeRelease()
}
