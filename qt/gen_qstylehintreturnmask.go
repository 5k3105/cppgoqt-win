package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QStyleHintReturnMask_base func(uintptr) uintptr
	addr_QStyleHintReturnMask_new uintptr
	addr_QStyleHintReturnMask_new2 uintptr
	addr_QStyleHintReturnMask_OperatorAssign uintptr
	addr_QStyleHintReturnMask_region_get uintptr
	addr_QStyleHintReturnMask_region_set uintptr
	pg_QStyleHintReturnMask_delete func(uintptr)
)

func init() {
	reg(&pg_QStyleHintReturnMask_base, "QStyleHintReturnMask_base")
	symaddr(&addr_QStyleHintReturnMask_new, "QStyleHintReturnMask_new")
	symaddr(&addr_QStyleHintReturnMask_new2, "QStyleHintReturnMask_new2")
	symaddr(&addr_QStyleHintReturnMask_OperatorAssign, "QStyleHintReturnMask_OperatorAssign")
	symaddr(&addr_QStyleHintReturnMask_region_get, "QStyleHintReturnMask_region_get")
	symaddr(&addr_QStyleHintReturnMask_region_set, "QStyleHintReturnMask_region_set")
	reg(&pg_QStyleHintReturnMask_delete, "QStyleHintReturnMask_delete")
}

type QStyleHintReturnMask struct {
	h uintptr
	*QStyleHintReturn
}

func bindQStyleHintReturnMask(h uintptr) *QStyleHintReturnMask {
	if h == 0 {
		return nil
	}
	return &QStyleHintReturnMask{h: h, QStyleHintReturn: bindQStyleHintReturn(pg_QStyleHintReturnMask_base(h))}
}

func newownedQStyleHintReturnMask(h uintptr) *QStyleHintReturnMask {
	o := bindQStyleHintReturnMask(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QStyleHintReturnMask).Delete)
	}
	return o
}

func (this *QStyleHintReturnMask) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QStyleHintReturnMask) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQStyleHintReturnMask(h unsafe.Pointer) *QStyleHintReturnMask { return bindQStyleHintReturnMask(uintptr(h)) }
func UnsafeNewQStyleHintReturnMaskOwned(h unsafe.Pointer) *QStyleHintReturnMask { return newownedQStyleHintReturnMask(uintptr(h)) }

func (this *QStyleHintReturnMask) UnsafeRelease() {
	this.h = 0
	if this.QStyleHintReturn != nil {
		this.QStyleHintReturn.UnsafeRelease()
	}
}

func NewQStyleHintReturnMask() *QStyleHintReturnMask {
	__r, _, _ := purego.SyscallN(addr_QStyleHintReturnMask_new)
	__o := newownedQStyleHintReturnMask(__r)
	return __o
}

func NewQStyleHintReturnMask2(arg1 *QStyleHintReturnMask) *QStyleHintReturnMask {
	__r, _, _ := purego.SyscallN(addr_QStyleHintReturnMask_new2, uintptr(arg1.UnsafePointer()))
	__o := newownedQStyleHintReturnMask(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QStyleHintReturnMask) OperatorAssign(arg1 *QStyleHintReturnMask) *QStyleHintReturnMask {
	__r, _, _ := purego.SyscallN(addr_QStyleHintReturnMask_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQStyleHintReturnMask(__r)
}

func (this *QStyleHintReturnMask) Region() *QRegion {
	__r, _, _ := purego.SyscallN(addr_QStyleHintReturnMask_region_get, this.h)
	runtime.KeepAlive(this)
	return newownedQRegion(__r)
}

func (this *QStyleHintReturnMask) SetRegion(region *QRegion) {
	purego.SyscallN(addr_QStyleHintReturnMask_region_set, this.h, uintptr(region.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(region)
}

func (this *QStyleHintReturnMask) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QStyleHintReturnMask_delete(this.h)
	}
	this.UnsafeRelease()
}
