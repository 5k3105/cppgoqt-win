package qt

import (
	"unsafe"
	"runtime"
)

var (
	pg_QBindingStatus_delete func(uintptr)
)

func init() {
	reg(&pg_QBindingStatus_delete, "QBindingStatus_delete")
}

type QBindingStatus struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QBindingStatus) UnsafeGoID() int64     { return this.vid }
func (this *QBindingStatus) UnsafeSetGoID(id int64) { this.vid = id }

func bindQBindingStatus(h uintptr) *QBindingStatus {
	if h == 0 {
		return nil
	}
	return &QBindingStatus{h: h}
}

func newownedQBindingStatus(h uintptr) *QBindingStatus {
	o := bindQBindingStatus(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QBindingStatus).Delete)
	}
	return o
}

func (this *QBindingStatus) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QBindingStatus) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQBindingStatus(h unsafe.Pointer) *QBindingStatus { return bindQBindingStatus(uintptr(h)) }
func UnsafeNewQBindingStatusOwned(h unsafe.Pointer) *QBindingStatus { return newownedQBindingStatus(uintptr(h)) }

func (this *QBindingStatus) UnsafeRelease() { this.h = 0 }

func (this *QBindingStatus) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QBindingStatus_delete(this.h)
	}
	this.UnsafeRelease()
}
