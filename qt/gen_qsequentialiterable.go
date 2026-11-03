package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QSequentialIterable_new uintptr
	addr_QSequentialIterable_new2 uintptr
	addr_QSequentialIterable_at uintptr
	addr_QSequentialIterable_set uintptr
	addr_QSequentialIterable_addValue uintptr
	addr_QSequentialIterable_removeValue uintptr
	addr_QSequentialIterable_valueMetaType uintptr
	addr_QSequentialIterable_canRandomAccessIterate uintptr
	addr_QSequentialIterable_canInputIterate uintptr
	addr_QSequentialIterable_canForwardIterate uintptr
	addr_QSequentialIterable_canReverseIterate uintptr
	addr_QSequentialIterable_size uintptr
	addr_QSequentialIterable_clear uintptr
	addr_QSequentialIterable_metaContainer uintptr
	addr_QSequentialIterable_addValue2 uintptr
	addr_QSequentialIterable_removeValue2 uintptr
	pg_QSequentialIterable_delete func(uintptr)
)

func init() {
	symaddr(&addr_QSequentialIterable_new, "QSequentialIterable_new")
	symaddr(&addr_QSequentialIterable_new2, "QSequentialIterable_new2")
	symaddr(&addr_QSequentialIterable_at, "QSequentialIterable_at")
	symaddr(&addr_QSequentialIterable_set, "QSequentialIterable_set")
	symaddr(&addr_QSequentialIterable_addValue, "QSequentialIterable_addValue")
	symaddr(&addr_QSequentialIterable_removeValue, "QSequentialIterable_removeValue")
	symaddr(&addr_QSequentialIterable_valueMetaType, "QSequentialIterable_valueMetaType")
	symaddr(&addr_QSequentialIterable_canRandomAccessIterate, "QSequentialIterable_canRandomAccessIterate")
	symaddr(&addr_QSequentialIterable_canInputIterate, "QSequentialIterable_canInputIterate")
	symaddr(&addr_QSequentialIterable_canForwardIterate, "QSequentialIterable_canForwardIterate")
	symaddr(&addr_QSequentialIterable_canReverseIterate, "QSequentialIterable_canReverseIterate")
	symaddr(&addr_QSequentialIterable_size, "QSequentialIterable_size")
	symaddr(&addr_QSequentialIterable_clear, "QSequentialIterable_clear")
	symaddr(&addr_QSequentialIterable_metaContainer, "QSequentialIterable_metaContainer")
	symaddr(&addr_QSequentialIterable_addValue2, "QSequentialIterable_addValue2")
	symaddr(&addr_QSequentialIterable_removeValue2, "QSequentialIterable_removeValue2")
	reg(&pg_QSequentialIterable_delete, "QSequentialIterable_delete")
}

type QSequentialIterable struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QSequentialIterable) UnsafeGoID() int64     { return this.vid }
func (this *QSequentialIterable) UnsafeSetGoID(id int64) { this.vid = id }

func bindQSequentialIterable(h uintptr) *QSequentialIterable {
	if h == 0 {
		return nil
	}
	return &QSequentialIterable{h: h}
}

func newownedQSequentialIterable(h uintptr) *QSequentialIterable {
	o := bindQSequentialIterable(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QSequentialIterable).Delete)
	}
	return o
}

func (this *QSequentialIterable) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QSequentialIterable) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQSequentialIterable(h unsafe.Pointer) *QSequentialIterable { return bindQSequentialIterable(uintptr(h)) }
func UnsafeNewQSequentialIterableOwned(h unsafe.Pointer) *QSequentialIterable { return newownedQSequentialIterable(uintptr(h)) }

func (this *QSequentialIterable) UnsafeRelease() { this.h = 0 }

func NewQSequentialIterable() *QSequentialIterable {
	__r, _, _ := purego.SyscallN(addr_QSequentialIterable_new)
	__o := newownedQSequentialIterable(__r)
	return __o
}

func NewQSequentialIterable2(arg1 *QSequentialIterable) *QSequentialIterable {
	__r, _, _ := purego.SyscallN(addr_QSequentialIterable_new2, uintptr(arg1.UnsafePointer()))
	__o := newownedQSequentialIterable(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QSequentialIterable) At(idx int64) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QSequentialIterable_at, this.h, uintptr(idx))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QSequentialIterable) Set(idx int64, value *QVariant) {
	purego.SyscallN(addr_QSequentialIterable_set, this.h, uintptr(idx), uintptr(value.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(value)
}

func (this *QSequentialIterable) AddValue(value *QVariant) {
	purego.SyscallN(addr_QSequentialIterable_addValue, this.h, uintptr(value.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(value)
}

func (this *QSequentialIterable) RemoveValue() {
	purego.SyscallN(addr_QSequentialIterable_removeValue, this.h)
	runtime.KeepAlive(this)
}

func (this *QSequentialIterable) ValueMetaType() *QMetaType {
	__r, _, _ := purego.SyscallN(addr_QSequentialIterable_valueMetaType, this.h)
	runtime.KeepAlive(this)
	return newownedQMetaType(__r)
}

func (this *QSequentialIterable) CanRandomAccessIterate() bool {
	__r0, _, _ := purego.SyscallN(addr_QSequentialIterable_canRandomAccessIterate, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QSequentialIterable) CanInputIterate() bool {
	__r0, _, _ := purego.SyscallN(addr_QSequentialIterable_canInputIterate, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QSequentialIterable) CanForwardIterate() bool {
	__r0, _, _ := purego.SyscallN(addr_QSequentialIterable_canForwardIterate, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QSequentialIterable) CanReverseIterate() bool {
	__r0, _, _ := purego.SyscallN(addr_QSequentialIterable_canReverseIterate, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QSequentialIterable) Size() int64 {
	__r0, _, _ := purego.SyscallN(addr_QSequentialIterable_size, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QSequentialIterable) Clear() {
	purego.SyscallN(addr_QSequentialIterable_clear, this.h)
	runtime.KeepAlive(this)
}

func (this *QSequentialIterable) MetaContainer() *QMetaSequence {
	__r, _, _ := purego.SyscallN(addr_QSequentialIterable_metaContainer, this.h)
	runtime.KeepAlive(this)
	return newownedQMetaSequence(__r)
}

func (this *QSequentialIterable) AddValue2(value *QVariant, position QSequentialIterable__Position) {
	purego.SyscallN(addr_QSequentialIterable_addValue2, this.h, uintptr(value.UnsafePointer()), uintptr(position))
	runtime.KeepAlive(this)
	runtime.KeepAlive(value)
}

func (this *QSequentialIterable) RemoveValueWithPosition(position QSequentialIterable__Position) {
	purego.SyscallN(addr_QSequentialIterable_removeValue2, this.h, uintptr(position))
	runtime.KeepAlive(this)
}

func (this *QSequentialIterable) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QSequentialIterable_delete(this.h)
	}
	this.UnsafeRelease()
}
