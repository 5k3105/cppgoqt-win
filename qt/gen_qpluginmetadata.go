package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QPluginMetaData_archRequirements uintptr
	addr_QPluginMetaData_size_get uintptr
	addr_QPluginMetaData_size_set uintptr
	pg_QPluginMetaData_delete func(uintptr)
)

func init() {
	symaddr(&addr_QPluginMetaData_archRequirements, "QPluginMetaData_archRequirements")
	symaddr(&addr_QPluginMetaData_size_get, "QPluginMetaData_size_get")
	symaddr(&addr_QPluginMetaData_size_set, "QPluginMetaData_size_set")
	reg(&pg_QPluginMetaData_delete, "QPluginMetaData_delete")
}

type QPluginMetaData struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QPluginMetaData) UnsafeGoID() int64     { return this.vid }
func (this *QPluginMetaData) UnsafeSetGoID(id int64) { this.vid = id }

func bindQPluginMetaData(h uintptr) *QPluginMetaData {
	if h == 0 {
		return nil
	}
	return &QPluginMetaData{h: h}
}

func newownedQPluginMetaData(h uintptr) *QPluginMetaData {
	o := bindQPluginMetaData(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QPluginMetaData).Delete)
	}
	return o
}

func (this *QPluginMetaData) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QPluginMetaData) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQPluginMetaData(h unsafe.Pointer) *QPluginMetaData { return bindQPluginMetaData(uintptr(h)) }
func UnsafeNewQPluginMetaDataOwned(h unsafe.Pointer) *QPluginMetaData { return newownedQPluginMetaData(uintptr(h)) }

func (this *QPluginMetaData) UnsafeRelease() { this.h = 0 }

func QPluginMetaData_ArchRequirements() byte {
	__r0, _, _ := purego.SyscallN(addr_QPluginMetaData_archRequirements)
	__r := uint8(__r0)
	return byte(__r)
}

func (this *QPluginMetaData) Size() uint64 {
	__r0, _, _ := purego.SyscallN(addr_QPluginMetaData_size_get, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return uint64(__r)
}

func (this *QPluginMetaData) SetSize(size uint64) {
	purego.SyscallN(addr_QPluginMetaData_size_set, this.h, uintptr(size))
	runtime.KeepAlive(this)
}

func (this *QPluginMetaData) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QPluginMetaData_delete(this.h)
	}
	this.UnsafeRelease()
}
