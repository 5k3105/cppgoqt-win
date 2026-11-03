package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QMetaObject__Data_new uintptr
	addr_QMetaObject__Data_new2 uintptr
	addr_QMetaObject__Data_OperatorAssign uintptr
	addr_QMetaObject__Data_superdata_get uintptr
	addr_QMetaObject__Data_superdata_set uintptr
	addr_QMetaObject__Data_relatedMetaObjects_get uintptr
	addr_QMetaObject__Data_relatedMetaObjects_set uintptr
	pg_QMetaObject__Data_delete func(uintptr)
)

func init() {
	symaddr(&addr_QMetaObject__Data_new, "QMetaObject__Data_new")
	symaddr(&addr_QMetaObject__Data_new2, "QMetaObject__Data_new2")
	symaddr(&addr_QMetaObject__Data_OperatorAssign, "QMetaObject__Data_OperatorAssign")
	symaddr(&addr_QMetaObject__Data_superdata_get, "QMetaObject__Data_superdata_get")
	symaddr(&addr_QMetaObject__Data_superdata_set, "QMetaObject__Data_superdata_set")
	symaddr(&addr_QMetaObject__Data_relatedMetaObjects_get, "QMetaObject__Data_relatedMetaObjects_get")
	symaddr(&addr_QMetaObject__Data_relatedMetaObjects_set, "QMetaObject__Data_relatedMetaObjects_set")
	reg(&pg_QMetaObject__Data_delete, "QMetaObject__Data_delete")
}

type QMetaObject__Data struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QMetaObject__Data) UnsafeGoID() int64     { return this.vid }
func (this *QMetaObject__Data) UnsafeSetGoID(id int64) { this.vid = id }

func bindQMetaObject__Data(h uintptr) *QMetaObject__Data {
	if h == 0 {
		return nil
	}
	return &QMetaObject__Data{h: h}
}

func newownedQMetaObject__Data(h uintptr) *QMetaObject__Data {
	o := bindQMetaObject__Data(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QMetaObject__Data).Delete)
	}
	return o
}

func (this *QMetaObject__Data) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QMetaObject__Data) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQMetaObject__Data(h unsafe.Pointer) *QMetaObject__Data { return bindQMetaObject__Data(uintptr(h)) }
func UnsafeNewQMetaObject__DataOwned(h unsafe.Pointer) *QMetaObject__Data { return newownedQMetaObject__Data(uintptr(h)) }

func (this *QMetaObject__Data) UnsafeRelease() { this.h = 0 }

func NewQMetaObject__Data() *QMetaObject__Data {
	__r, _, _ := purego.SyscallN(addr_QMetaObject__Data_new)
	__o := newownedQMetaObject__Data(__r)
	return __o
}

func NewQMetaObject__Data2(arg1 *QMetaObject__Data) *QMetaObject__Data {
	__r, _, _ := purego.SyscallN(addr_QMetaObject__Data_new2, uintptr(arg1.UnsafePointer()))
	__o := newownedQMetaObject__Data(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QMetaObject__Data) OperatorAssign(arg1 *QMetaObject__Data) *QMetaObject__Data {
	__r, _, _ := purego.SyscallN(addr_QMetaObject__Data_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQMetaObject__Data(__r)
}

func (this *QMetaObject__Data) Superdata() *QMetaObject__SuperData {
	__r, _, _ := purego.SyscallN(addr_QMetaObject__Data_superdata_get, this.h)
	runtime.KeepAlive(this)
	return newownedQMetaObject__SuperData(__r)
}

func (this *QMetaObject__Data) SetSuperdata(superdata *QMetaObject__SuperData) {
	purego.SyscallN(addr_QMetaObject__Data_superdata_set, this.h, uintptr(superdata.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(superdata)
}

func (this *QMetaObject__Data) RelatedMetaObjects() *QMetaObject__SuperData {
	__r, _, _ := purego.SyscallN(addr_QMetaObject__Data_relatedMetaObjects_get, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject__SuperData(__r)
}

func (this *QMetaObject__Data) SetRelatedMetaObjects(relatedMetaObjects *QMetaObject__SuperData) {
	purego.SyscallN(addr_QMetaObject__Data_relatedMetaObjects_set, this.h, uintptr(relatedMetaObjects.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(relatedMetaObjects)
}

func (this *QMetaObject__Data) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QMetaObject__Data_delete(this.h)
	}
	this.UnsafeRelease()
}
