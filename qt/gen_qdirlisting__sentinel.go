package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QDirListing__sentinel_OperatorEqual uintptr
	addr_QDirListing__sentinel_OperatorNotEqual uintptr
	pg_QDirListing__sentinel_delete func(uintptr)
)

func init() {
	symaddr(&addr_QDirListing__sentinel_OperatorEqual, "QDirListing__sentinel_OperatorEqual")
	symaddr(&addr_QDirListing__sentinel_OperatorNotEqual, "QDirListing__sentinel_OperatorNotEqual")
	reg(&pg_QDirListing__sentinel_delete, "QDirListing__sentinel_delete")
}

type QDirListing__sentinel struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QDirListing__sentinel) UnsafeGoID() int64     { return this.vid }
func (this *QDirListing__sentinel) UnsafeSetGoID(id int64) { this.vid = id }

func bindQDirListing__sentinel(h uintptr) *QDirListing__sentinel {
	if h == 0 {
		return nil
	}
	return &QDirListing__sentinel{h: h}
}

func newownedQDirListing__sentinel(h uintptr) *QDirListing__sentinel {
	o := bindQDirListing__sentinel(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QDirListing__sentinel).Delete)
	}
	return o
}

func (this *QDirListing__sentinel) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QDirListing__sentinel) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQDirListing__sentinel(h unsafe.Pointer) *QDirListing__sentinel { return bindQDirListing__sentinel(uintptr(h)) }
func UnsafeNewQDirListing__sentinelOwned(h unsafe.Pointer) *QDirListing__sentinel { return newownedQDirListing__sentinel(uintptr(h)) }

func (this *QDirListing__sentinel) UnsafeRelease() { this.h = 0 }

func (this *QDirListing__sentinel) OperatorEqual(arg1 *QDirListing__sentinel) bool {
	__r0, _, _ := purego.SyscallN(addr_QDirListing__sentinel_OperatorEqual, this.h, uintptr(arg1.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return __r != 0
}

func (this *QDirListing__sentinel) OperatorNotEqual(arg1 *QDirListing__sentinel) bool {
	__r0, _, _ := purego.SyscallN(addr_QDirListing__sentinel_OperatorNotEqual, this.h, uintptr(arg1.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return __r != 0
}

func (this *QDirListing__sentinel) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QDirListing__sentinel_delete(this.h)
	}
	this.UnsafeRelease()
}
