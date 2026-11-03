package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QPropertyProxyBindingData_d_ptr_get uintptr
	addr_QPropertyProxyBindingData_d_ptr_set uintptr
	addr_QPropertyProxyBindingData_propertyData_get uintptr
	addr_QPropertyProxyBindingData_propertyData_set uintptr
	pg_QPropertyProxyBindingData_delete func(uintptr)
)

func init() {
	symaddr(&addr_QPropertyProxyBindingData_d_ptr_get, "QPropertyProxyBindingData_d_ptr_get")
	symaddr(&addr_QPropertyProxyBindingData_d_ptr_set, "QPropertyProxyBindingData_d_ptr_set")
	symaddr(&addr_QPropertyProxyBindingData_propertyData_get, "QPropertyProxyBindingData_propertyData_get")
	symaddr(&addr_QPropertyProxyBindingData_propertyData_set, "QPropertyProxyBindingData_propertyData_set")
	reg(&pg_QPropertyProxyBindingData_delete, "QPropertyProxyBindingData_delete")
}

type QPropertyProxyBindingData struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QPropertyProxyBindingData) UnsafeGoID() int64     { return this.vid }
func (this *QPropertyProxyBindingData) UnsafeSetGoID(id int64) { this.vid = id }

func bindQPropertyProxyBindingData(h uintptr) *QPropertyProxyBindingData {
	if h == 0 {
		return nil
	}
	return &QPropertyProxyBindingData{h: h}
}

func newownedQPropertyProxyBindingData(h uintptr) *QPropertyProxyBindingData {
	o := bindQPropertyProxyBindingData(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QPropertyProxyBindingData).Delete)
	}
	return o
}

func (this *QPropertyProxyBindingData) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QPropertyProxyBindingData) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQPropertyProxyBindingData(h unsafe.Pointer) *QPropertyProxyBindingData { return bindQPropertyProxyBindingData(uintptr(h)) }
func UnsafeNewQPropertyProxyBindingDataOwned(h unsafe.Pointer) *QPropertyProxyBindingData { return newownedQPropertyProxyBindingData(uintptr(h)) }

func (this *QPropertyProxyBindingData) UnsafeRelease() { this.h = 0 }

func (this *QPropertyProxyBindingData) D_ptr() uintptr {
	__r0, _, _ := purego.SyscallN(addr_QPropertyProxyBindingData_d_ptr_get, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return uintptr(__r)
}

func (this *QPropertyProxyBindingData) SetD_ptr(d_ptr uintptr) {
	purego.SyscallN(addr_QPropertyProxyBindingData_d_ptr_set, this.h, uintptr(d_ptr))
	runtime.KeepAlive(this)
}

func (this *QPropertyProxyBindingData) PropertyData() *QUntypedPropertyData {
	__r, _, _ := purego.SyscallN(addr_QPropertyProxyBindingData_propertyData_get, this.h)
	runtime.KeepAlive(this)
	return bindQUntypedPropertyData(__r)
}

func (this *QPropertyProxyBindingData) SetPropertyData(propertyData *QUntypedPropertyData) {
	purego.SyscallN(addr_QPropertyProxyBindingData_propertyData_set, this.h, uintptr(propertyData.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(propertyData)
}

func (this *QPropertyProxyBindingData) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QPropertyProxyBindingData_delete(this.h)
	}
	this.UnsafeRelease()
}
