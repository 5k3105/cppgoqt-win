package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QProcess__CreateProcessArguments_inheritHandles_get uintptr
	addr_QProcess__CreateProcessArguments_inheritHandles_set uintptr
	addr_QProcess__CreateProcessArguments_flags_get uintptr
	addr_QProcess__CreateProcessArguments_flags_set uintptr
	pg_QProcess__CreateProcessArguments_delete func(uintptr)
)

func init() {
	symaddr(&addr_QProcess__CreateProcessArguments_inheritHandles_get, "QProcess__CreateProcessArguments_inheritHandles_get")
	symaddr(&addr_QProcess__CreateProcessArguments_inheritHandles_set, "QProcess__CreateProcessArguments_inheritHandles_set")
	symaddr(&addr_QProcess__CreateProcessArguments_flags_get, "QProcess__CreateProcessArguments_flags_get")
	symaddr(&addr_QProcess__CreateProcessArguments_flags_set, "QProcess__CreateProcessArguments_flags_set")
	reg(&pg_QProcess__CreateProcessArguments_delete, "QProcess__CreateProcessArguments_delete")
}

type QProcess__CreateProcessArguments struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QProcess__CreateProcessArguments) UnsafeGoID() int64     { return this.vid }
func (this *QProcess__CreateProcessArguments) UnsafeSetGoID(id int64) { this.vid = id }

func bindQProcess__CreateProcessArguments(h uintptr) *QProcess__CreateProcessArguments {
	if h == 0 {
		return nil
	}
	return &QProcess__CreateProcessArguments{h: h}
}

func newownedQProcess__CreateProcessArguments(h uintptr) *QProcess__CreateProcessArguments {
	o := bindQProcess__CreateProcessArguments(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QProcess__CreateProcessArguments).Delete)
	}
	return o
}

func (this *QProcess__CreateProcessArguments) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QProcess__CreateProcessArguments) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQProcess__CreateProcessArguments(h unsafe.Pointer) *QProcess__CreateProcessArguments { return bindQProcess__CreateProcessArguments(uintptr(h)) }
func UnsafeNewQProcess__CreateProcessArgumentsOwned(h unsafe.Pointer) *QProcess__CreateProcessArguments { return newownedQProcess__CreateProcessArguments(uintptr(h)) }

func (this *QProcess__CreateProcessArguments) UnsafeRelease() { this.h = 0 }

func (this *QProcess__CreateProcessArguments) InheritHandles() bool {
	__r0, _, _ := purego.SyscallN(addr_QProcess__CreateProcessArguments_inheritHandles_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QProcess__CreateProcessArguments) SetInheritHandles(inheritHandles bool) {
	purego.SyscallN(addr_QProcess__CreateProcessArguments_inheritHandles_set, this.h, uintptr(b2i(inheritHandles)))
	runtime.KeepAlive(this)
}

func (this *QProcess__CreateProcessArguments) Flags() uint32 {
	__r0, _, _ := purego.SyscallN(addr_QProcess__CreateProcessArguments_flags_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint32(__r)
}

func (this *QProcess__CreateProcessArguments) SetFlags(flags uint32) {
	purego.SyscallN(addr_QProcess__CreateProcessArguments_flags_set, this.h, uintptr(flags))
	runtime.KeepAlive(this)
}

func (this *QProcess__CreateProcessArguments) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QProcess__CreateProcessArguments_delete(this.h)
	}
	this.UnsafeRelease()
}
