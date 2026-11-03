package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QException_new uintptr
	addr_QException_new2 uintptr
	addr_QException_OperatorAssign uintptr
	addr_QException_raise uintptr
	addr_QException_clone uintptr
	addr_QException_what uintptr
	addr_QException_raise_default uintptr
	addr_QException_clone_default uintptr
	pg_QException__set_override func(uintptr, int32)
	pg_QException_delete func(uintptr)
)

func init() {
	symaddr(&addr_QException_new, "QException_new")
	symaddr(&addr_QException_new2, "QException_new2")
	symaddr(&addr_QException_OperatorAssign, "QException_OperatorAssign")
	symaddr(&addr_QException_raise, "QException_raise")
	symaddr(&addr_QException_clone, "QException_clone")
	symaddr(&addr_QException_what, "QException_what")
	symaddr(&addr_QException_raise_default, "QException_raise_default")
	symaddr(&addr_QException_clone_default, "QException_clone_default")
	reg(&pg_QException__set_override, "QException__set_override")
	reg(&pg_QException_delete, "QException_delete")
}

type QException struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QException) UnsafeGoID() int64     { return this.vid }
func (this *QException) UnsafeSetGoID(id int64) { this.vid = id }

func bindQException(h uintptr) *QException {
	if h == 0 {
		return nil
	}
	return &QException{h: h}
}

func newownedQException(h uintptr) *QException {
	o := bindQException(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QException).Delete)
	}
	return o
}

func (this *QException) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QException) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQException(h unsafe.Pointer) *QException { return bindQException(uintptr(h)) }
func UnsafeNewQExceptionOwned(h unsafe.Pointer) *QException { return newownedQException(uintptr(h)) }

func (this *QException) UnsafeRelease() { this.h = 0 }

func NewQException() *QException {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QException_new, uintptr(__goid))
	__o := newownedQException(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QException_virtIdx, QException_enable)
	return __o
}

func NewQException2(arg1 *QException) *QException {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QException_new2, uintptr(__goid), uintptr(arg1.UnsafePointer()))
	__o := newownedQException(__r)
	runtime.KeepAlive(arg1)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QException_virtIdx, QException_enable)
	return __o
}

func (this *QException) OperatorAssign(arg1 *QException) *QException {
	__r, _, _ := purego.SyscallN(addr_QException_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQException(__r)
}

func (this *QException) Raise() {
	purego.SyscallN(addr_QException_raise, this.h)
	runtime.KeepAlive(this)
}

func (this *QException) Clone() *QException {
	__r, _, _ := purego.SyscallN(addr_QException_clone, this.h)
	runtime.KeepAlive(this)
	return bindQException(__r)
}

func (this *QException) What() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QException_what, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QException) RaiseDefault() {
	purego.SyscallN(addr_QException_raise_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QException) CloneDefault() *QException {
	__r, _, _ := purego.SyscallN(addr_QException_clone_default, this.h)
	runtime.KeepAlive(this)
	return bindQException(__r)
}

func QException_virtIdx(slot int64) int32 {
	switch slot {
	case 8013398049691041699:
		return 0
	case -7032813281048616478:
		return 1
	}
	return -1
}

func QException_enable(h uintptr, idx int32) { pg_QException__set_override(h, idx) }

func (this *QException) OnRaise(fn func()) {
	registerVirt(this.UnsafeGoID(), 8013398049691041699, func() { fn() }) // slot "raise"
}

func (this *QException) OnClone(fn func() *QException) {
	registerVirt(this.UnsafeGoID(), -7032813281048616478, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "clone"
}

func (this *QException) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QException_delete(this.h)
	}
	this.UnsafeRelease()
}
