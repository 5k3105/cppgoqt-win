package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QAssociativeConstIterator_new uintptr
	addr_QAssociativeConstIterator_key uintptr
	addr_QAssociativeConstIterator_value uintptr
	addr_QAssociativeConstIterator_OperatorMultiply uintptr
	addr_QAssociativeConstIterator_OperatorMinusGreater uintptr
	addr_QAssociativeConstIterator_metaContainer uintptr
	pg_QAssociativeConstIterator_delete func(uintptr)
)

func init() {
	symaddr(&addr_QAssociativeConstIterator_new, "QAssociativeConstIterator_new")
	symaddr(&addr_QAssociativeConstIterator_key, "QAssociativeConstIterator_key")
	symaddr(&addr_QAssociativeConstIterator_value, "QAssociativeConstIterator_value")
	symaddr(&addr_QAssociativeConstIterator_OperatorMultiply, "QAssociativeConstIterator_OperatorMultiply")
	symaddr(&addr_QAssociativeConstIterator_OperatorMinusGreater, "QAssociativeConstIterator_OperatorMinusGreater")
	symaddr(&addr_QAssociativeConstIterator_metaContainer, "QAssociativeConstIterator_metaContainer")
	reg(&pg_QAssociativeConstIterator_delete, "QAssociativeConstIterator_delete")
}

type QAssociativeConstIterator struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QAssociativeConstIterator) UnsafeGoID() int64     { return this.vid }
func (this *QAssociativeConstIterator) UnsafeSetGoID(id int64) { this.vid = id }

func bindQAssociativeConstIterator(h uintptr) *QAssociativeConstIterator {
	if h == 0 {
		return nil
	}
	return &QAssociativeConstIterator{h: h}
}

func newownedQAssociativeConstIterator(h uintptr) *QAssociativeConstIterator {
	o := bindQAssociativeConstIterator(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QAssociativeConstIterator).Delete)
	}
	return o
}

func (this *QAssociativeConstIterator) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QAssociativeConstIterator) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQAssociativeConstIterator(h unsafe.Pointer) *QAssociativeConstIterator { return bindQAssociativeConstIterator(uintptr(h)) }
func UnsafeNewQAssociativeConstIteratorOwned(h unsafe.Pointer) *QAssociativeConstIterator { return newownedQAssociativeConstIterator(uintptr(h)) }

func (this *QAssociativeConstIterator) UnsafeRelease() { this.h = 0 }

func NewQAssociativeConstIterator(arg1 *QAssociativeConstIterator) *QAssociativeConstIterator {
	__r, _, _ := purego.SyscallN(addr_QAssociativeConstIterator_new, uintptr(arg1.UnsafePointer()))
	__o := newownedQAssociativeConstIterator(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QAssociativeConstIterator) Key() *QVariant {
	__r, _, _ := purego.SyscallN(addr_QAssociativeConstIterator_key, this.h)
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QAssociativeConstIterator) Value() *QVariant {
	__r, _, _ := purego.SyscallN(addr_QAssociativeConstIterator_value, this.h)
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QAssociativeConstIterator) OperatorMultiply() *QVariant {
	__r, _, _ := purego.SyscallN(addr_QAssociativeConstIterator_OperatorMultiply, this.h)
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QAssociativeConstIterator) OperatorMinusGreater() *QVariantConstPointer {
	__r, _, _ := purego.SyscallN(addr_QAssociativeConstIterator_OperatorMinusGreater, this.h)
	runtime.KeepAlive(this)
	return newownedQVariantConstPointer(__r)
}

func (this *QAssociativeConstIterator) MetaContainer() *QMetaAssociation {
	__r, _, _ := purego.SyscallN(addr_QAssociativeConstIterator_metaContainer, this.h)
	runtime.KeepAlive(this)
	return newownedQMetaAssociation(__r)
}

func (this *QAssociativeConstIterator) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QAssociativeConstIterator_delete(this.h)
	}
	this.UnsafeRelease()
}
