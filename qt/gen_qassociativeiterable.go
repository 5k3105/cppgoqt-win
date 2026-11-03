package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QAssociativeIterable_new uintptr
	addr_QAssociativeIterable_new2 uintptr
	addr_QAssociativeIterable_containsKey uintptr
	addr_QAssociativeIterable_insertKey uintptr
	addr_QAssociativeIterable_removeKey uintptr
	addr_QAssociativeIterable_value uintptr
	addr_QAssociativeIterable_setValue uintptr
	addr_QAssociativeIterable_canRandomAccessIterate uintptr
	addr_QAssociativeIterable_canInputIterate uintptr
	addr_QAssociativeIterable_canForwardIterate uintptr
	addr_QAssociativeIterable_canReverseIterate uintptr
	addr_QAssociativeIterable_size uintptr
	addr_QAssociativeIterable_clear uintptr
	addr_QAssociativeIterable_metaContainer uintptr
	pg_QAssociativeIterable_delete func(uintptr)
)

func init() {
	symaddr(&addr_QAssociativeIterable_new, "QAssociativeIterable_new")
	symaddr(&addr_QAssociativeIterable_new2, "QAssociativeIterable_new2")
	symaddr(&addr_QAssociativeIterable_containsKey, "QAssociativeIterable_containsKey")
	symaddr(&addr_QAssociativeIterable_insertKey, "QAssociativeIterable_insertKey")
	symaddr(&addr_QAssociativeIterable_removeKey, "QAssociativeIterable_removeKey")
	symaddr(&addr_QAssociativeIterable_value, "QAssociativeIterable_value")
	symaddr(&addr_QAssociativeIterable_setValue, "QAssociativeIterable_setValue")
	symaddr(&addr_QAssociativeIterable_canRandomAccessIterate, "QAssociativeIterable_canRandomAccessIterate")
	symaddr(&addr_QAssociativeIterable_canInputIterate, "QAssociativeIterable_canInputIterate")
	symaddr(&addr_QAssociativeIterable_canForwardIterate, "QAssociativeIterable_canForwardIterate")
	symaddr(&addr_QAssociativeIterable_canReverseIterate, "QAssociativeIterable_canReverseIterate")
	symaddr(&addr_QAssociativeIterable_size, "QAssociativeIterable_size")
	symaddr(&addr_QAssociativeIterable_clear, "QAssociativeIterable_clear")
	symaddr(&addr_QAssociativeIterable_metaContainer, "QAssociativeIterable_metaContainer")
	reg(&pg_QAssociativeIterable_delete, "QAssociativeIterable_delete")
}

type QAssociativeIterable struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QAssociativeIterable) UnsafeGoID() int64     { return this.vid }
func (this *QAssociativeIterable) UnsafeSetGoID(id int64) { this.vid = id }

func bindQAssociativeIterable(h uintptr) *QAssociativeIterable {
	if h == 0 {
		return nil
	}
	return &QAssociativeIterable{h: h}
}

func newownedQAssociativeIterable(h uintptr) *QAssociativeIterable {
	o := bindQAssociativeIterable(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QAssociativeIterable).Delete)
	}
	return o
}

func (this *QAssociativeIterable) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QAssociativeIterable) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQAssociativeIterable(h unsafe.Pointer) *QAssociativeIterable { return bindQAssociativeIterable(uintptr(h)) }
func UnsafeNewQAssociativeIterableOwned(h unsafe.Pointer) *QAssociativeIterable { return newownedQAssociativeIterable(uintptr(h)) }

func (this *QAssociativeIterable) UnsafeRelease() { this.h = 0 }

func NewQAssociativeIterable() *QAssociativeIterable {
	__r, _, _ := purego.SyscallN(addr_QAssociativeIterable_new)
	__o := newownedQAssociativeIterable(__r)
	return __o
}

func NewQAssociativeIterable2(arg1 *QAssociativeIterable) *QAssociativeIterable {
	__r, _, _ := purego.SyscallN(addr_QAssociativeIterable_new2, uintptr(arg1.UnsafePointer()))
	__o := newownedQAssociativeIterable(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QAssociativeIterable) ContainsKey(key *QVariant) bool {
	__r0, _, _ := purego.SyscallN(addr_QAssociativeIterable_containsKey, this.h, uintptr(key.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
	return __r != 0
}

func (this *QAssociativeIterable) InsertKey(key *QVariant) {
	purego.SyscallN(addr_QAssociativeIterable_insertKey, this.h, uintptr(key.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
}

func (this *QAssociativeIterable) RemoveKey(key *QVariant) {
	purego.SyscallN(addr_QAssociativeIterable_removeKey, this.h, uintptr(key.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
}

func (this *QAssociativeIterable) Value(key *QVariant) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QAssociativeIterable_value, this.h, uintptr(key.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
	return newownedQVariant(__r)
}

func (this *QAssociativeIterable) SetValue(key *QVariant, mapped *QVariant) {
	purego.SyscallN(addr_QAssociativeIterable_setValue, this.h, uintptr(key.UnsafePointer()), uintptr(mapped.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
	runtime.KeepAlive(mapped)
}

func (this *QAssociativeIterable) CanRandomAccessIterate() bool {
	__r0, _, _ := purego.SyscallN(addr_QAssociativeIterable_canRandomAccessIterate, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QAssociativeIterable) CanInputIterate() bool {
	__r0, _, _ := purego.SyscallN(addr_QAssociativeIterable_canInputIterate, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QAssociativeIterable) CanForwardIterate() bool {
	__r0, _, _ := purego.SyscallN(addr_QAssociativeIterable_canForwardIterate, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QAssociativeIterable) CanReverseIterate() bool {
	__r0, _, _ := purego.SyscallN(addr_QAssociativeIterable_canReverseIterate, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QAssociativeIterable) Size() int64 {
	__r0, _, _ := purego.SyscallN(addr_QAssociativeIterable_size, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QAssociativeIterable) Clear() {
	purego.SyscallN(addr_QAssociativeIterable_clear, this.h)
	runtime.KeepAlive(this)
}

func (this *QAssociativeIterable) MetaContainer() *QMetaAssociation {
	__r, _, _ := purego.SyscallN(addr_QAssociativeIterable_metaContainer, this.h)
	runtime.KeepAlive(this)
	return newownedQMetaAssociation(__r)
}

func (this *QAssociativeIterable) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QAssociativeIterable_delete(this.h)
	}
	this.UnsafeRelease()
}
