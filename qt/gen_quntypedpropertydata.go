package qt

import (
	"unsafe"
	"runtime"
)

var (
	pg_QUntypedPropertyData_delete func(uintptr)
)

func init() {
	reg(&pg_QUntypedPropertyData_delete, "QUntypedPropertyData_delete")
}

type QUntypedPropertyData struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QUntypedPropertyData) UnsafeGoID() int64     { return this.vid }
func (this *QUntypedPropertyData) UnsafeSetGoID(id int64) { this.vid = id }

func bindQUntypedPropertyData(h uintptr) *QUntypedPropertyData {
	if h == 0 {
		return nil
	}
	return &QUntypedPropertyData{h: h}
}

func newownedQUntypedPropertyData(h uintptr) *QUntypedPropertyData {
	o := bindQUntypedPropertyData(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QUntypedPropertyData).Delete)
	}
	return o
}

func (this *QUntypedPropertyData) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QUntypedPropertyData) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQUntypedPropertyData(h unsafe.Pointer) *QUntypedPropertyData { return bindQUntypedPropertyData(uintptr(h)) }
func UnsafeNewQUntypedPropertyDataOwned(h unsafe.Pointer) *QUntypedPropertyData { return newownedQUntypedPropertyData(uintptr(h)) }

func (this *QUntypedPropertyData) UnsafeRelease() { this.h = 0 }

func (this *QUntypedPropertyData) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QUntypedPropertyData_delete(this.h)
	}
	this.UnsafeRelease()
}
