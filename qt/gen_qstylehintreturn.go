package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QStyleHintReturn_new uintptr
	addr_QStyleHintReturn_new2 uintptr
	addr_QStyleHintReturn_new3 uintptr
	addr_QStyleHintReturn_new4 uintptr
	addr_QStyleHintReturn_OperatorAssign uintptr
	addr_QStyleHintReturn_version_get uintptr
	addr_QStyleHintReturn_version_set uintptr
	addr_QStyleHintReturn_type_get uintptr
	addr_QStyleHintReturn_type_set uintptr
	pg_QStyleHintReturn_delete func(uintptr)
)

func init() {
	symaddr(&addr_QStyleHintReturn_new, "QStyleHintReturn_new")
	symaddr(&addr_QStyleHintReturn_new2, "QStyleHintReturn_new2")
	symaddr(&addr_QStyleHintReturn_new3, "QStyleHintReturn_new3")
	symaddr(&addr_QStyleHintReturn_new4, "QStyleHintReturn_new4")
	symaddr(&addr_QStyleHintReturn_OperatorAssign, "QStyleHintReturn_OperatorAssign")
	symaddr(&addr_QStyleHintReturn_version_get, "QStyleHintReturn_version_get")
	symaddr(&addr_QStyleHintReturn_version_set, "QStyleHintReturn_version_set")
	symaddr(&addr_QStyleHintReturn_type_get, "QStyleHintReturn_type_get")
	symaddr(&addr_QStyleHintReturn_type_set, "QStyleHintReturn_type_set")
	reg(&pg_QStyleHintReturn_delete, "QStyleHintReturn_delete")
}

type QStyleHintReturn struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QStyleHintReturn) UnsafeGoID() int64     { return this.vid }
func (this *QStyleHintReturn) UnsafeSetGoID(id int64) { this.vid = id }

func bindQStyleHintReturn(h uintptr) *QStyleHintReturn {
	if h == 0 {
		return nil
	}
	return &QStyleHintReturn{h: h}
}

func newownedQStyleHintReturn(h uintptr) *QStyleHintReturn {
	o := bindQStyleHintReturn(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QStyleHintReturn).Delete)
	}
	return o
}

func (this *QStyleHintReturn) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QStyleHintReturn) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQStyleHintReturn(h unsafe.Pointer) *QStyleHintReturn { return bindQStyleHintReturn(uintptr(h)) }
func UnsafeNewQStyleHintReturnOwned(h unsafe.Pointer) *QStyleHintReturn { return newownedQStyleHintReturn(uintptr(h)) }

func (this *QStyleHintReturn) UnsafeRelease() { this.h = 0 }

func NewQStyleHintReturn() *QStyleHintReturn {
	__r, _, _ := purego.SyscallN(addr_QStyleHintReturn_new)
	__o := newownedQStyleHintReturn(__r)
	return __o
}

func NewQStyleHintReturn2(arg1 *QStyleHintReturn) *QStyleHintReturn {
	__r, _, _ := purego.SyscallN(addr_QStyleHintReturn_new2, uintptr(arg1.UnsafePointer()))
	__o := newownedQStyleHintReturn(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func NewQStyleHintReturn3(version int) *QStyleHintReturn {
	__r, _, _ := purego.SyscallN(addr_QStyleHintReturn_new3, uintptr(version))
	__o := newownedQStyleHintReturn(__r)
	return __o
}

func NewQStyleHintReturn4(version int, type_ int) *QStyleHintReturn {
	__r, _, _ := purego.SyscallN(addr_QStyleHintReturn_new4, uintptr(version), uintptr(type_))
	__o := newownedQStyleHintReturn(__r)
	return __o
}

func (this *QStyleHintReturn) OperatorAssign(arg1 *QStyleHintReturn) *QStyleHintReturn {
	__r, _, _ := purego.SyscallN(addr_QStyleHintReturn_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQStyleHintReturn(__r)
}

func (this *QStyleHintReturn) Version() int {
	__r0, _, _ := purego.SyscallN(addr_QStyleHintReturn_version_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QStyleHintReturn) SetVersion(version int) {
	purego.SyscallN(addr_QStyleHintReturn_version_set, this.h, uintptr(version))
	runtime.KeepAlive(this)
}

func (this *QStyleHintReturn) Type() int {
	__r0, _, _ := purego.SyscallN(addr_QStyleHintReturn_type_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QStyleHintReturn) SetType(type_ int) {
	purego.SyscallN(addr_QStyleHintReturn_type_set, this.h, uintptr(type_))
	runtime.KeepAlive(this)
}

func (this *QStyleHintReturn) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QStyleHintReturn_delete(this.h)
	}
	this.UnsafeRelease()
}
