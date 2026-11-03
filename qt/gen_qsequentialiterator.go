package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QSequentialIterator_new uintptr
	addr_QSequentialIterator_metaContainer uintptr
	pg_QSequentialIterator_delete func(uintptr)
)

func init() {
	symaddr(&addr_QSequentialIterator_new, "QSequentialIterator_new")
	symaddr(&addr_QSequentialIterator_metaContainer, "QSequentialIterator_metaContainer")
	reg(&pg_QSequentialIterator_delete, "QSequentialIterator_delete")
}

type QSequentialIterator struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QSequentialIterator) UnsafeGoID() int64     { return this.vid }
func (this *QSequentialIterator) UnsafeSetGoID(id int64) { this.vid = id }

func bindQSequentialIterator(h uintptr) *QSequentialIterator {
	if h == 0 {
		return nil
	}
	return &QSequentialIterator{h: h}
}

func newownedQSequentialIterator(h uintptr) *QSequentialIterator {
	o := bindQSequentialIterator(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QSequentialIterator).Delete)
	}
	return o
}

func (this *QSequentialIterator) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QSequentialIterator) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQSequentialIterator(h unsafe.Pointer) *QSequentialIterator { return bindQSequentialIterator(uintptr(h)) }
func UnsafeNewQSequentialIteratorOwned(h unsafe.Pointer) *QSequentialIterator { return newownedQSequentialIterator(uintptr(h)) }

func (this *QSequentialIterator) UnsafeRelease() { this.h = 0 }

func NewQSequentialIterator(arg1 *QSequentialIterator) *QSequentialIterator {
	__r, _, _ := purego.SyscallN(addr_QSequentialIterator_new, uintptr(arg1.UnsafePointer()))
	__o := newownedQSequentialIterator(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QSequentialIterator) MetaContainer() *QMetaSequence {
	__r, _, _ := purego.SyscallN(addr_QSequentialIterator_metaContainer, this.h)
	runtime.KeepAlive(this)
	return newownedQMetaSequence(__r)
}

func (this *QSequentialIterator) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QSequentialIterator_delete(this.h)
	}
	this.UnsafeRelease()
}
