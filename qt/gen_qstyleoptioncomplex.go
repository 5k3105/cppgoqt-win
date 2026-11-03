package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QStyleOptionComplex_base func(uintptr) uintptr
	addr_QStyleOptionComplex_new uintptr
	addr_QStyleOptionComplex_new2 uintptr
	addr_QStyleOptionComplex_new3 uintptr
	addr_QStyleOptionComplex_new4 uintptr
	addr_QStyleOptionComplex_OperatorAssign uintptr
	addr_QStyleOptionComplex_subControls_get uintptr
	addr_QStyleOptionComplex_subControls_set uintptr
	addr_QStyleOptionComplex_activeSubControls_get uintptr
	addr_QStyleOptionComplex_activeSubControls_set uintptr
	pg_QStyleOptionComplex_delete func(uintptr)
)

func init() {
	reg(&pg_QStyleOptionComplex_base, "QStyleOptionComplex_base")
	symaddr(&addr_QStyleOptionComplex_new, "QStyleOptionComplex_new")
	symaddr(&addr_QStyleOptionComplex_new2, "QStyleOptionComplex_new2")
	symaddr(&addr_QStyleOptionComplex_new3, "QStyleOptionComplex_new3")
	symaddr(&addr_QStyleOptionComplex_new4, "QStyleOptionComplex_new4")
	symaddr(&addr_QStyleOptionComplex_OperatorAssign, "QStyleOptionComplex_OperatorAssign")
	symaddr(&addr_QStyleOptionComplex_subControls_get, "QStyleOptionComplex_subControls_get")
	symaddr(&addr_QStyleOptionComplex_subControls_set, "QStyleOptionComplex_subControls_set")
	symaddr(&addr_QStyleOptionComplex_activeSubControls_get, "QStyleOptionComplex_activeSubControls_get")
	symaddr(&addr_QStyleOptionComplex_activeSubControls_set, "QStyleOptionComplex_activeSubControls_set")
	reg(&pg_QStyleOptionComplex_delete, "QStyleOptionComplex_delete")
}

type QStyleOptionComplex struct {
	h uintptr
	*QStyleOption
}

func bindQStyleOptionComplex(h uintptr) *QStyleOptionComplex {
	if h == 0 {
		return nil
	}
	return &QStyleOptionComplex{h: h, QStyleOption: bindQStyleOption(pg_QStyleOptionComplex_base(h))}
}

func newownedQStyleOptionComplex(h uintptr) *QStyleOptionComplex {
	o := bindQStyleOptionComplex(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QStyleOptionComplex).Delete)
	}
	return o
}

func (this *QStyleOptionComplex) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QStyleOptionComplex) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQStyleOptionComplex(h unsafe.Pointer) *QStyleOptionComplex { return bindQStyleOptionComplex(uintptr(h)) }
func UnsafeNewQStyleOptionComplexOwned(h unsafe.Pointer) *QStyleOptionComplex { return newownedQStyleOptionComplex(uintptr(h)) }

func (this *QStyleOptionComplex) UnsafeRelease() {
	this.h = 0
	if this.QStyleOption != nil {
		this.QStyleOption.UnsafeRelease()
	}
}

func NewQStyleOptionComplex() *QStyleOptionComplex {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionComplex_new)
	__o := newownedQStyleOptionComplex(__r)
	return __o
}

func NewQStyleOptionComplex2(other *QStyleOptionComplex) *QStyleOptionComplex {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionComplex_new2, uintptr(other.UnsafePointer()))
	__o := newownedQStyleOptionComplex(__r)
	runtime.KeepAlive(other)
	return __o
}

func NewQStyleOptionComplex3(version int) *QStyleOptionComplex {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionComplex_new3, uintptr(version))
	__o := newownedQStyleOptionComplex(__r)
	return __o
}

func NewQStyleOptionComplex4(version int, type_ int) *QStyleOptionComplex {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionComplex_new4, uintptr(version), uintptr(type_))
	__o := newownedQStyleOptionComplex(__r)
	return __o
}

func (this *QStyleOptionComplex) OperatorAssign(arg1 *QStyleOptionComplex) *QStyleOptionComplex {
	__r, _, _ := purego.SyscallN(addr_QStyleOptionComplex_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQStyleOptionComplex(__r)
}

func (this *QStyleOptionComplex) SubControls() QStyle__SubControl {
	__r0, _, _ := purego.SyscallN(addr_QStyleOptionComplex_subControls_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QStyle__SubControl(__r)
}

func (this *QStyleOptionComplex) SetSubControls(subControls QStyle__SubControl) {
	purego.SyscallN(addr_QStyleOptionComplex_subControls_set, this.h, uintptr(subControls))
	runtime.KeepAlive(this)
}

func (this *QStyleOptionComplex) ActiveSubControls() QStyle__SubControl {
	__r0, _, _ := purego.SyscallN(addr_QStyleOptionComplex_activeSubControls_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QStyle__SubControl(__r)
}

func (this *QStyleOptionComplex) SetActiveSubControls(activeSubControls QStyle__SubControl) {
	purego.SyscallN(addr_QStyleOptionComplex_activeSubControls_set, this.h, uintptr(activeSubControls))
	runtime.KeepAlive(this)
}

func (this *QStyleOptionComplex) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QStyleOptionComplex_delete(this.h)
	}
	this.UnsafeRelease()
}
