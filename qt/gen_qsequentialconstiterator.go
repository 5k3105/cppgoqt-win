package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QSequentialConstIterator_new uintptr
	addr_QSequentialConstIterator_OperatorMultiply uintptr
	addr_QSequentialConstIterator_OperatorMinusGreater uintptr
	addr_QSequentialConstIterator_metaContainer uintptr
	pg_QSequentialConstIterator_delete func(uintptr)
)

func init() {
	symaddr(&addr_QSequentialConstIterator_new, "QSequentialConstIterator_new")
	symaddr(&addr_QSequentialConstIterator_OperatorMultiply, "QSequentialConstIterator_OperatorMultiply")
	symaddr(&addr_QSequentialConstIterator_OperatorMinusGreater, "QSequentialConstIterator_OperatorMinusGreater")
	symaddr(&addr_QSequentialConstIterator_metaContainer, "QSequentialConstIterator_metaContainer")
	reg(&pg_QSequentialConstIterator_delete, "QSequentialConstIterator_delete")
}

type QSequentialConstIterator struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QSequentialConstIterator) UnsafeGoID() int64     { return this.vid }
func (this *QSequentialConstIterator) UnsafeSetGoID(id int64) { this.vid = id }

func bindQSequentialConstIterator(h uintptr) *QSequentialConstIterator {
	if h == 0 {
		return nil
	}
	return &QSequentialConstIterator{h: h}
}

func newownedQSequentialConstIterator(h uintptr) *QSequentialConstIterator {
	o := bindQSequentialConstIterator(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QSequentialConstIterator).Delete)
	}
	return o
}

func (this *QSequentialConstIterator) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QSequentialConstIterator) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQSequentialConstIterator(h unsafe.Pointer) *QSequentialConstIterator { return bindQSequentialConstIterator(uintptr(h)) }
func UnsafeNewQSequentialConstIteratorOwned(h unsafe.Pointer) *QSequentialConstIterator { return newownedQSequentialConstIterator(uintptr(h)) }

func (this *QSequentialConstIterator) UnsafeRelease() { this.h = 0 }

func NewQSequentialConstIterator(arg1 *QSequentialConstIterator) *QSequentialConstIterator {
	__r, _, _ := purego.SyscallN(addr_QSequentialConstIterator_new, uintptr(arg1.UnsafePointer()))
	__o := newownedQSequentialConstIterator(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QSequentialConstIterator) OperatorMultiply() *QVariant {
	__r, _, _ := purego.SyscallN(addr_QSequentialConstIterator_OperatorMultiply, this.h)
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QSequentialConstIterator) OperatorMinusGreater() *QVariantConstPointer {
	__r, _, _ := purego.SyscallN(addr_QSequentialConstIterator_OperatorMinusGreater, this.h)
	runtime.KeepAlive(this)
	return newownedQVariantConstPointer(__r)
}

func (this *QSequentialConstIterator) MetaContainer() *QMetaSequence {
	__r, _, _ := purego.SyscallN(addr_QSequentialConstIterator_metaContainer, this.h)
	runtime.KeepAlive(this)
	return newownedQMetaSequence(__r)
}

func (this *QSequentialConstIterator) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QSequentialConstIterator_delete(this.h)
	}
	this.UnsafeRelease()
}
