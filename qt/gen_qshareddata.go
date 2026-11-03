package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QSharedData_new uintptr
	addr_QSharedData_new2 uintptr
	addr_QSharedData_ref_get uintptr
	addr_QSharedData_ref_set uintptr
	pg_QSharedData_delete func(uintptr)
)

func init() {
	symaddr(&addr_QSharedData_new, "QSharedData_new")
	symaddr(&addr_QSharedData_new2, "QSharedData_new2")
	symaddr(&addr_QSharedData_ref_get, "QSharedData_ref_get")
	symaddr(&addr_QSharedData_ref_set, "QSharedData_ref_set")
	reg(&pg_QSharedData_delete, "QSharedData_delete")
}

type QSharedData struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QSharedData) UnsafeGoID() int64     { return this.vid }
func (this *QSharedData) UnsafeSetGoID(id int64) { this.vid = id }

func bindQSharedData(h uintptr) *QSharedData {
	if h == 0 {
		return nil
	}
	return &QSharedData{h: h}
}

func newownedQSharedData(h uintptr) *QSharedData {
	o := bindQSharedData(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QSharedData).Delete)
	}
	return o
}

func (this *QSharedData) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QSharedData) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQSharedData(h unsafe.Pointer) *QSharedData { return bindQSharedData(uintptr(h)) }
func UnsafeNewQSharedDataOwned(h unsafe.Pointer) *QSharedData { return newownedQSharedData(uintptr(h)) }

func (this *QSharedData) UnsafeRelease() { this.h = 0 }

func NewQSharedData() *QSharedData {
	__r, _, _ := purego.SyscallN(addr_QSharedData_new)
	__o := newownedQSharedData(__r)
	return __o
}

func NewQSharedData2(arg1 *QSharedData) *QSharedData {
	__r, _, _ := purego.SyscallN(addr_QSharedData_new2, uintptr(arg1.UnsafePointer()))
	__o := newownedQSharedData(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QSharedData) Ref() *QAtomicInt {
	__r, _, _ := purego.SyscallN(addr_QSharedData_ref_get, this.h)
	runtime.KeepAlive(this)
	return newownedQAtomicInt(__r)
}

func (this *QSharedData) SetRef(ref *QAtomicInt) {
	purego.SyscallN(addr_QSharedData_ref_set, this.h, uintptr(ref.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(ref)
}

func (this *QSharedData) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QSharedData_delete(this.h)
	}
	this.UnsafeRelease()
}
