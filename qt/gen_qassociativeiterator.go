package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QAssociativeIterator_new uintptr
	addr_QAssociativeIterator_key uintptr
	addr_QAssociativeIterator_metaContainer uintptr
	pg_QAssociativeIterator_delete func(uintptr)
)

func init() {
	symaddr(&addr_QAssociativeIterator_new, "QAssociativeIterator_new")
	symaddr(&addr_QAssociativeIterator_key, "QAssociativeIterator_key")
	symaddr(&addr_QAssociativeIterator_metaContainer, "QAssociativeIterator_metaContainer")
	reg(&pg_QAssociativeIterator_delete, "QAssociativeIterator_delete")
}

type QAssociativeIterator struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QAssociativeIterator) UnsafeGoID() int64     { return this.vid }
func (this *QAssociativeIterator) UnsafeSetGoID(id int64) { this.vid = id }

func bindQAssociativeIterator(h uintptr) *QAssociativeIterator {
	if h == 0 {
		return nil
	}
	return &QAssociativeIterator{h: h}
}

func newownedQAssociativeIterator(h uintptr) *QAssociativeIterator {
	o := bindQAssociativeIterator(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QAssociativeIterator).Delete)
	}
	return o
}

func (this *QAssociativeIterator) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QAssociativeIterator) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQAssociativeIterator(h unsafe.Pointer) *QAssociativeIterator { return bindQAssociativeIterator(uintptr(h)) }
func UnsafeNewQAssociativeIteratorOwned(h unsafe.Pointer) *QAssociativeIterator { return newownedQAssociativeIterator(uintptr(h)) }

func (this *QAssociativeIterator) UnsafeRelease() { this.h = 0 }

func NewQAssociativeIterator(arg1 *QAssociativeIterator) *QAssociativeIterator {
	__r, _, _ := purego.SyscallN(addr_QAssociativeIterator_new, uintptr(arg1.UnsafePointer()))
	__o := newownedQAssociativeIterator(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QAssociativeIterator) Key() *QVariant {
	__r, _, _ := purego.SyscallN(addr_QAssociativeIterator_key, this.h)
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QAssociativeIterator) MetaContainer() *QMetaAssociation {
	__r, _, _ := purego.SyscallN(addr_QAssociativeIterator_metaContainer, this.h)
	runtime.KeepAlive(this)
	return newownedQMetaAssociation(__r)
}

func (this *QAssociativeIterator) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QAssociativeIterator_delete(this.h)
	}
	this.UnsafeRelease()
}
