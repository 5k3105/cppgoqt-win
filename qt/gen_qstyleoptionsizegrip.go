package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QStyleOptionSizeGrip_chain uintptr
	addr_QStyleOptionSizeGrip_new uintptr
	addr_QStyleOptionSizeGrip_new2 uintptr
	addr_QStyleOptionSizeGrip_OperatorAssign uintptr
	addr_QStyleOptionSizeGrip_corner_get uintptr
	addr_QStyleOptionSizeGrip_corner_set uintptr
	pg_QStyleOptionSizeGrip_delete func(uintptr)
)

func init() {
	symaddr(&addr_QStyleOptionSizeGrip_chain, "QStyleOptionSizeGrip_chain")
	symaddr(&addr_QStyleOptionSizeGrip_new, "QStyleOptionSizeGrip_new")
	symaddr(&addr_QStyleOptionSizeGrip_new2, "QStyleOptionSizeGrip_new2")
	symaddr(&addr_QStyleOptionSizeGrip_OperatorAssign, "QStyleOptionSizeGrip_OperatorAssign")
	symaddr(&addr_QStyleOptionSizeGrip_corner_get, "QStyleOptionSizeGrip_corner_get")
	symaddr(&addr_QStyleOptionSizeGrip_corner_set, "QStyleOptionSizeGrip_corner_set")
	reg(&pg_QStyleOptionSizeGrip_delete, "QStyleOptionSizeGrip_delete")
}

type QStyleOptionSizeGrip struct {
	h uintptr
	*QStyleOptionComplex
}

func bindQStyleOptionSizeGrip(h uintptr) *QStyleOptionSizeGrip {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QStyleOptionSizeGrip_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QStyleOptionSizeGrip{h: h, QStyleOptionComplex: &QStyleOptionComplex{h: __c[0], QStyleOption: &QStyleOption{h: __c[1]}}}
}

func newownedQStyleOptionSizeGrip(h uintptr) *QStyleOptionSizeGrip {
	o := bindQStyleOptionSizeGrip(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QStyleOptionSizeGrip).Delete)
	}
	return o
}

func (this *QStyleOptionSizeGrip) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QStyleOptionSizeGrip) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQStyleOptionSizeGrip(h unsafe.Pointer) *QStyleOptionSizeGrip { return bindQStyleOptionSizeGrip(uintptr(h)) }
func UnsafeNewQStyleOptionSizeGripOwned(h unsafe.Pointer) *QStyleOptionSizeGrip { return newownedQStyleOptionSizeGrip(uintptr(h)) }

func (this *QStyleOptionSizeGrip) UnsafeRelease() {
	this.h = 0
	if this.QStyleOptionComplex != nil {
		this.QStyleOptionComplex.UnsafeRelease()
	}
}

func NewQStyleOptionSizeGrip() *QStyleOptionSizeGrip {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionSizeGrip_new)
	__o := newownedQStyleOptionSizeGrip(__r)
	return __o
}

func NewQStyleOptionSizeGrip2(other *QStyleOptionSizeGrip) *QStyleOptionSizeGrip {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionSizeGrip_new2, uintptr(other.UnsafePointer()))
	__o := newownedQStyleOptionSizeGrip(__r)
	runtime.KeepAlive(other)
	return __o
}

func (this *QStyleOptionSizeGrip) OperatorAssign(arg1 *QStyleOptionSizeGrip) *QStyleOptionSizeGrip {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionSizeGrip_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQStyleOptionSizeGrip(__r)
}

func (this *QStyleOptionSizeGrip) Corner() Corner {
	__r0, _, _ := purego.SyscallN(addr_QStyleOptionSizeGrip_corner_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return Corner(__r)
}

func (this *QStyleOptionSizeGrip) SetCorner(corner Corner) {
	purego.SyscallN(addr_QStyleOptionSizeGrip_corner_set, this.h, uintptr(corner))
	runtime.KeepAlive(this)
}

func (this *QStyleOptionSizeGrip) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QStyleOptionSizeGrip_delete(this.h)
	}
	this.UnsafeRelease()
}
