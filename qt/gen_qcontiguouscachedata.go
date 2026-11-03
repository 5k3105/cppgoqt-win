package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QContiguousCacheData_allocateData uintptr
	addr_QContiguousCacheData_freeData uintptr
	addr_QContiguousCacheData_alloc_get uintptr
	addr_QContiguousCacheData_alloc_set uintptr
	addr_QContiguousCacheData_count_get uintptr
	addr_QContiguousCacheData_count_set uintptr
	addr_QContiguousCacheData_start_get uintptr
	addr_QContiguousCacheData_start_set uintptr
	addr_QContiguousCacheData_offset_get uintptr
	addr_QContiguousCacheData_offset_set uintptr
	pg_QContiguousCacheData_delete func(uintptr)
)

func init() {
	symaddr(&addr_QContiguousCacheData_allocateData, "QContiguousCacheData_allocateData")
	symaddr(&addr_QContiguousCacheData_freeData, "QContiguousCacheData_freeData")
	symaddr(&addr_QContiguousCacheData_alloc_get, "QContiguousCacheData_alloc_get")
	symaddr(&addr_QContiguousCacheData_alloc_set, "QContiguousCacheData_alloc_set")
	symaddr(&addr_QContiguousCacheData_count_get, "QContiguousCacheData_count_get")
	symaddr(&addr_QContiguousCacheData_count_set, "QContiguousCacheData_count_set")
	symaddr(&addr_QContiguousCacheData_start_get, "QContiguousCacheData_start_get")
	symaddr(&addr_QContiguousCacheData_start_set, "QContiguousCacheData_start_set")
	symaddr(&addr_QContiguousCacheData_offset_get, "QContiguousCacheData_offset_get")
	symaddr(&addr_QContiguousCacheData_offset_set, "QContiguousCacheData_offset_set")
	reg(&pg_QContiguousCacheData_delete, "QContiguousCacheData_delete")
}

type QContiguousCacheData struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QContiguousCacheData) UnsafeGoID() int64     { return this.vid }
func (this *QContiguousCacheData) UnsafeSetGoID(id int64) { this.vid = id }

func bindQContiguousCacheData(h uintptr) *QContiguousCacheData {
	if h == 0 {
		return nil
	}
	return &QContiguousCacheData{h: h}
}

func newownedQContiguousCacheData(h uintptr) *QContiguousCacheData {
	o := bindQContiguousCacheData(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QContiguousCacheData).Delete)
	}
	return o
}

func (this *QContiguousCacheData) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QContiguousCacheData) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQContiguousCacheData(h unsafe.Pointer) *QContiguousCacheData { return bindQContiguousCacheData(uintptr(h)) }
func UnsafeNewQContiguousCacheDataOwned(h unsafe.Pointer) *QContiguousCacheData { return newownedQContiguousCacheData(uintptr(h)) }

func (this *QContiguousCacheData) UnsafeRelease() { this.h = 0 }

func QContiguousCacheData_AllocateData(size int64, alignment int64) *QContiguousCacheData {
	__r, _, _ := purego.SyscallN(addr_QContiguousCacheData_allocateData, uintptr(size), uintptr(alignment))
	return bindQContiguousCacheData(__r)
}

func QContiguousCacheData_FreeData(data *QContiguousCacheData) {
	purego.SyscallN(addr_QContiguousCacheData_freeData, uintptr(data.UnsafePointer()))
	runtime.KeepAlive(data)
}

func (this *QContiguousCacheData) Alloc() int64 {
	__r0, _, _ := purego.SyscallN(addr_QContiguousCacheData_alloc_get, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QContiguousCacheData) SetAlloc(alloc int64) {
	purego.SyscallN(addr_QContiguousCacheData_alloc_set, this.h, uintptr(alloc))
	runtime.KeepAlive(this)
}

func (this *QContiguousCacheData) Count() int64 {
	__r0, _, _ := purego.SyscallN(addr_QContiguousCacheData_count_get, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QContiguousCacheData) SetCount(count int64) {
	purego.SyscallN(addr_QContiguousCacheData_count_set, this.h, uintptr(count))
	runtime.KeepAlive(this)
}

func (this *QContiguousCacheData) Start() int64 {
	__r0, _, _ := purego.SyscallN(addr_QContiguousCacheData_start_get, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QContiguousCacheData) SetStart(start int64) {
	purego.SyscallN(addr_QContiguousCacheData_start_set, this.h, uintptr(start))
	runtime.KeepAlive(this)
}

func (this *QContiguousCacheData) Offset() int64 {
	__r0, _, _ := purego.SyscallN(addr_QContiguousCacheData_offset_get, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QContiguousCacheData) SetOffset(offset int64) {
	purego.SyscallN(addr_QContiguousCacheData_offset_set, this.h, uintptr(offset))
	runtime.KeepAlive(this)
}

func (this *QContiguousCacheData) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QContiguousCacheData_delete(this.h)
	}
	this.UnsafeRelease()
}
