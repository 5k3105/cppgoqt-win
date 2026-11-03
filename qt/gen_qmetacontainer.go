package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QMetaContainer_new uintptr
	addr_QMetaContainer_new2 uintptr
	addr_QMetaContainer_hasInputIterator uintptr
	addr_QMetaContainer_hasForwardIterator uintptr
	addr_QMetaContainer_hasBidirectionalIterator uintptr
	addr_QMetaContainer_hasRandomAccessIterator uintptr
	addr_QMetaContainer_hasSize uintptr
	addr_QMetaContainer_canClear uintptr
	addr_QMetaContainer_hasIterator uintptr
	addr_QMetaContainer_hasConstIterator uintptr
	pg_QMetaContainer_delete func(uintptr)
)

func init() {
	symaddr(&addr_QMetaContainer_new, "QMetaContainer_new")
	symaddr(&addr_QMetaContainer_new2, "QMetaContainer_new2")
	symaddr(&addr_QMetaContainer_hasInputIterator, "QMetaContainer_hasInputIterator")
	symaddr(&addr_QMetaContainer_hasForwardIterator, "QMetaContainer_hasForwardIterator")
	symaddr(&addr_QMetaContainer_hasBidirectionalIterator, "QMetaContainer_hasBidirectionalIterator")
	symaddr(&addr_QMetaContainer_hasRandomAccessIterator, "QMetaContainer_hasRandomAccessIterator")
	symaddr(&addr_QMetaContainer_hasSize, "QMetaContainer_hasSize")
	symaddr(&addr_QMetaContainer_canClear, "QMetaContainer_canClear")
	symaddr(&addr_QMetaContainer_hasIterator, "QMetaContainer_hasIterator")
	symaddr(&addr_QMetaContainer_hasConstIterator, "QMetaContainer_hasConstIterator")
	reg(&pg_QMetaContainer_delete, "QMetaContainer_delete")
}

type QMetaContainer struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QMetaContainer) UnsafeGoID() int64     { return this.vid }
func (this *QMetaContainer) UnsafeSetGoID(id int64) { this.vid = id }

func bindQMetaContainer(h uintptr) *QMetaContainer {
	if h == 0 {
		return nil
	}
	return &QMetaContainer{h: h}
}

func newownedQMetaContainer(h uintptr) *QMetaContainer {
	o := bindQMetaContainer(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QMetaContainer).Delete)
	}
	return o
}

func (this *QMetaContainer) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QMetaContainer) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQMetaContainer(h unsafe.Pointer) *QMetaContainer { return bindQMetaContainer(uintptr(h)) }
func UnsafeNewQMetaContainerOwned(h unsafe.Pointer) *QMetaContainer { return newownedQMetaContainer(uintptr(h)) }

func (this *QMetaContainer) UnsafeRelease() { this.h = 0 }

func NewQMetaContainer() *QMetaContainer {
	__r, _, _ := purego.SyscallN(addr_QMetaContainer_new)
	__o := newownedQMetaContainer(__r)
	return __o
}

func NewQMetaContainer2(arg1 *QMetaContainer) *QMetaContainer {
	__r, _, _ := purego.SyscallN(addr_QMetaContainer_new2, uintptr(arg1.UnsafePointer()))
	__o := newownedQMetaContainer(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QMetaContainer) HasInputIterator() bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaContainer_hasInputIterator, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMetaContainer) HasForwardIterator() bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaContainer_hasForwardIterator, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMetaContainer) HasBidirectionalIterator() bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaContainer_hasBidirectionalIterator, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMetaContainer) HasRandomAccessIterator() bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaContainer_hasRandomAccessIterator, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMetaContainer) HasSize() bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaContainer_hasSize, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMetaContainer) CanClear() bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaContainer_canClear, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMetaContainer) HasIterator() bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaContainer_hasIterator, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMetaContainer) HasConstIterator() bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaContainer_hasConstIterator, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMetaContainer) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QMetaContainer_delete(this.h)
	}
	this.UnsafeRelease()
}
