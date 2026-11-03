package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QArrayData_allocatedCapacity uintptr
	addr_QArrayData_constAllocatedCapacity uintptr
	addr_QArrayData_ref uintptr
	addr_QArrayData_deref uintptr
	addr_QArrayData_isShared uintptr
	addr_QArrayData_needsDetach uintptr
	addr_QArrayData_detachCapacity uintptr
	addr_QArrayData_deallocate uintptr
	addr_QArrayData_flags_get uintptr
	addr_QArrayData_flags_set uintptr
	addr_QArrayData_alloc_get uintptr
	addr_QArrayData_alloc_set uintptr
	pg_QArrayData_delete func(uintptr)
)

func init() {
	symaddr(&addr_QArrayData_allocatedCapacity, "QArrayData_allocatedCapacity")
	symaddr(&addr_QArrayData_constAllocatedCapacity, "QArrayData_constAllocatedCapacity")
	symaddr(&addr_QArrayData_ref, "QArrayData_ref")
	symaddr(&addr_QArrayData_deref, "QArrayData_deref")
	symaddr(&addr_QArrayData_isShared, "QArrayData_isShared")
	symaddr(&addr_QArrayData_needsDetach, "QArrayData_needsDetach")
	symaddr(&addr_QArrayData_detachCapacity, "QArrayData_detachCapacity")
	symaddr(&addr_QArrayData_deallocate, "QArrayData_deallocate")
	symaddr(&addr_QArrayData_flags_get, "QArrayData_flags_get")
	symaddr(&addr_QArrayData_flags_set, "QArrayData_flags_set")
	symaddr(&addr_QArrayData_alloc_get, "QArrayData_alloc_get")
	symaddr(&addr_QArrayData_alloc_set, "QArrayData_alloc_set")
	reg(&pg_QArrayData_delete, "QArrayData_delete")
}

type QArrayData struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QArrayData) UnsafeGoID() int64     { return this.vid }
func (this *QArrayData) UnsafeSetGoID(id int64) { this.vid = id }

func bindQArrayData(h uintptr) *QArrayData {
	if h == 0 {
		return nil
	}
	return &QArrayData{h: h}
}

func newownedQArrayData(h uintptr) *QArrayData {
	o := bindQArrayData(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QArrayData).Delete)
	}
	return o
}

func (this *QArrayData) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QArrayData) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQArrayData(h unsafe.Pointer) *QArrayData { return bindQArrayData(uintptr(h)) }
func UnsafeNewQArrayDataOwned(h unsafe.Pointer) *QArrayData { return newownedQArrayData(uintptr(h)) }

func (this *QArrayData) UnsafeRelease() { this.h = 0 }

func (this *QArrayData) AllocatedCapacity() int64 {
	__r0, _, _ := purego.SyscallN(addr_QArrayData_allocatedCapacity, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QArrayData) ConstAllocatedCapacity() int64 {
	__r0, _, _ := purego.SyscallN(addr_QArrayData_constAllocatedCapacity, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QArrayData) Ref() bool {
	__r0, _, _ := purego.SyscallN(addr_QArrayData_ref, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QArrayData) Deref() bool {
	__r0, _, _ := purego.SyscallN(addr_QArrayData_deref, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QArrayData) IsShared() bool {
	__r0, _, _ := purego.SyscallN(addr_QArrayData_isShared, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QArrayData) NeedsDetach() bool {
	__r0, _, _ := purego.SyscallN(addr_QArrayData_needsDetach, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QArrayData) DetachCapacity(newSize int64) int64 {
	__r0, _, _ := purego.SyscallN(addr_QArrayData_detachCapacity, this.h, uintptr(newSize))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func QArrayData_Deallocate(data *QArrayData, objectSize int64, alignment int64) {
	purego.SyscallN(addr_QArrayData_deallocate, uintptr(data.UnsafePointer()), uintptr(objectSize), uintptr(alignment))
	runtime.KeepAlive(data)
}

func (this *QArrayData) Flags() QArrayData__ArrayOption {
	__r0, _, _ := purego.SyscallN(addr_QArrayData_flags_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QArrayData__ArrayOption(__r)
}

func (this *QArrayData) SetFlags(flags QArrayData__ArrayOption) {
	purego.SyscallN(addr_QArrayData_flags_set, this.h, uintptr(flags))
	runtime.KeepAlive(this)
}

func (this *QArrayData) Alloc() int64 {
	__r0, _, _ := purego.SyscallN(addr_QArrayData_alloc_get, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QArrayData) SetAlloc(alloc int64) {
	purego.SyscallN(addr_QArrayData_alloc_set, this.h, uintptr(alloc))
	runtime.KeepAlive(this)
}

func (this *QArrayData) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QArrayData_delete(this.h)
	}
	this.UnsafeRelease()
}
