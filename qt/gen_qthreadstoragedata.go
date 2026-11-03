package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QThreadStorageData_id_get uintptr
	addr_QThreadStorageData_id_set uintptr
	pg_QThreadStorageData_delete func(uintptr)
)

func init() {
	symaddr(&addr_QThreadStorageData_id_get, "QThreadStorageData_id_get")
	symaddr(&addr_QThreadStorageData_id_set, "QThreadStorageData_id_set")
	reg(&pg_QThreadStorageData_delete, "QThreadStorageData_delete")
}

type QThreadStorageData struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QThreadStorageData) UnsafeGoID() int64     { return this.vid }
func (this *QThreadStorageData) UnsafeSetGoID(id int64) { this.vid = id }

func bindQThreadStorageData(h uintptr) *QThreadStorageData {
	if h == 0 {
		return nil
	}
	return &QThreadStorageData{h: h}
}

func newownedQThreadStorageData(h uintptr) *QThreadStorageData {
	o := bindQThreadStorageData(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QThreadStorageData).Delete)
	}
	return o
}

func (this *QThreadStorageData) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QThreadStorageData) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQThreadStorageData(h unsafe.Pointer) *QThreadStorageData { return bindQThreadStorageData(uintptr(h)) }
func UnsafeNewQThreadStorageDataOwned(h unsafe.Pointer) *QThreadStorageData { return newownedQThreadStorageData(uintptr(h)) }

func (this *QThreadStorageData) UnsafeRelease() { this.h = 0 }

func (this *QThreadStorageData) Id() int {
	__r0, _, _ := purego.SyscallN(addr_QThreadStorageData_id_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QThreadStorageData) SetId(id int) {
	purego.SyscallN(addr_QThreadStorageData_id_set, this.h, uintptr(id))
	runtime.KeepAlive(this)
}

func (this *QThreadStorageData) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QThreadStorageData_delete(this.h)
	}
	this.UnsafeRelease()
}
