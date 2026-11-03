package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QCollatorSortKey_new uintptr
	addr_QCollatorSortKey_OperatorAssign uintptr
	addr_QCollatorSortKey_swap uintptr
	addr_QCollatorSortKey_compare uintptr
	addr_QCollatorSortKey_OperatorLesser uintptr
	pg_QCollatorSortKey_delete func(uintptr)
)

func init() {
	symaddr(&addr_QCollatorSortKey_new, "QCollatorSortKey_new")
	symaddr(&addr_QCollatorSortKey_OperatorAssign, "QCollatorSortKey_OperatorAssign")
	symaddr(&addr_QCollatorSortKey_swap, "QCollatorSortKey_swap")
	symaddr(&addr_QCollatorSortKey_compare, "QCollatorSortKey_compare")
	symaddr(&addr_QCollatorSortKey_OperatorLesser, "QCollatorSortKey_OperatorLesser")
	reg(&pg_QCollatorSortKey_delete, "QCollatorSortKey_delete")
}

type QCollatorSortKey struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QCollatorSortKey) UnsafeGoID() int64     { return this.vid }
func (this *QCollatorSortKey) UnsafeSetGoID(id int64) { this.vid = id }

func bindQCollatorSortKey(h uintptr) *QCollatorSortKey {
	if h == 0 {
		return nil
	}
	return &QCollatorSortKey{h: h}
}

func newownedQCollatorSortKey(h uintptr) *QCollatorSortKey {
	o := bindQCollatorSortKey(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QCollatorSortKey).Delete)
	}
	return o
}

func (this *QCollatorSortKey) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QCollatorSortKey) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQCollatorSortKey(h unsafe.Pointer) *QCollatorSortKey { return bindQCollatorSortKey(uintptr(h)) }
func UnsafeNewQCollatorSortKeyOwned(h unsafe.Pointer) *QCollatorSortKey { return newownedQCollatorSortKey(uintptr(h)) }

func (this *QCollatorSortKey) UnsafeRelease() { this.h = 0 }

func NewQCollatorSortKey(other *QCollatorSortKey) *QCollatorSortKey {
	__r, _, _ := purego.SyscallN(addr_QCollatorSortKey_new, uintptr(other.UnsafePointer()))
	__o := newownedQCollatorSortKey(__r)
	runtime.KeepAlive(other)
	return __o
}

func (this *QCollatorSortKey) OperatorAssign(other *QCollatorSortKey) *QCollatorSortKey {
	__r, _, _ := purego.SyscallN(addr_QCollatorSortKey_OperatorAssign, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return bindQCollatorSortKey(__r)
}

func (this *QCollatorSortKey) Swap(other *QCollatorSortKey) {
	purego.SyscallN(addr_QCollatorSortKey_swap, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
}

func (this *QCollatorSortKey) Compare(key *QCollatorSortKey) int {
	__r0, _, _ := purego.SyscallN(addr_QCollatorSortKey_compare, this.h, uintptr(key.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
	return int(__r)
}

func (this *QCollatorSortKey) OperatorLesser(rhs *QCollatorSortKey) bool {
	__r0, _, _ := purego.SyscallN(addr_QCollatorSortKey_OperatorLesser, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QCollatorSortKey) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QCollatorSortKey_delete(this.h)
	}
	this.UnsafeRelease()
}
