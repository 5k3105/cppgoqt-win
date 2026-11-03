package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QMetaObject__SuperData_new uintptr
	addr_QMetaObject__SuperData_new2 uintptr
	addr_QMetaObject__SuperData_new3 uintptr
	addr_QMetaObject__SuperData_OperatorMinusGreater uintptr
	addr_QMetaObject__SuperData_OperatorAssign uintptr
	addr_QMetaObject__SuperData_direct_get uintptr
	addr_QMetaObject__SuperData_direct_set uintptr
	pg_QMetaObject__SuperData_delete func(uintptr)
)

func init() {
	symaddr(&addr_QMetaObject__SuperData_new, "QMetaObject__SuperData_new")
	symaddr(&addr_QMetaObject__SuperData_new2, "QMetaObject__SuperData_new2")
	symaddr(&addr_QMetaObject__SuperData_new3, "QMetaObject__SuperData_new3")
	symaddr(&addr_QMetaObject__SuperData_OperatorMinusGreater, "QMetaObject__SuperData_OperatorMinusGreater")
	symaddr(&addr_QMetaObject__SuperData_OperatorAssign, "QMetaObject__SuperData_OperatorAssign")
	symaddr(&addr_QMetaObject__SuperData_direct_get, "QMetaObject__SuperData_direct_get")
	symaddr(&addr_QMetaObject__SuperData_direct_set, "QMetaObject__SuperData_direct_set")
	reg(&pg_QMetaObject__SuperData_delete, "QMetaObject__SuperData_delete")
}

type QMetaObject__SuperData struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QMetaObject__SuperData) UnsafeGoID() int64     { return this.vid }
func (this *QMetaObject__SuperData) UnsafeSetGoID(id int64) { this.vid = id }

func bindQMetaObject__SuperData(h uintptr) *QMetaObject__SuperData {
	if h == 0 {
		return nil
	}
	return &QMetaObject__SuperData{h: h}
}

func newownedQMetaObject__SuperData(h uintptr) *QMetaObject__SuperData {
	o := bindQMetaObject__SuperData(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QMetaObject__SuperData).Delete)
	}
	return o
}

func (this *QMetaObject__SuperData) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QMetaObject__SuperData) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQMetaObject__SuperData(h unsafe.Pointer) *QMetaObject__SuperData { return bindQMetaObject__SuperData(uintptr(h)) }
func UnsafeNewQMetaObject__SuperDataOwned(h unsafe.Pointer) *QMetaObject__SuperData { return newownedQMetaObject__SuperData(uintptr(h)) }

func (this *QMetaObject__SuperData) UnsafeRelease() { this.h = 0 }

func NewQMetaObject__SuperData() *QMetaObject__SuperData {
	__r, _, _ := purego.SyscallN(addr_QMetaObject__SuperData_new)
	__o := newownedQMetaObject__SuperData(__r)
	return __o
}

func NewQMetaObject__SuperData2(mo *QMetaObject) *QMetaObject__SuperData {
	__r, _, _ := purego.SyscallN(addr_QMetaObject__SuperData_new2, uintptr(mo.UnsafePointer()))
	__o := newownedQMetaObject__SuperData(__r)
	runtime.KeepAlive(mo)
	return __o
}

func NewQMetaObject__SuperData3(arg1 *QMetaObject__SuperData) *QMetaObject__SuperData {
	__r, _, _ := purego.SyscallN(addr_QMetaObject__SuperData_new3, uintptr(arg1.UnsafePointer()))
	__o := newownedQMetaObject__SuperData(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QMetaObject__SuperData) OperatorMinusGreater() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QMetaObject__SuperData_OperatorMinusGreater, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QMetaObject__SuperData) OperatorAssign(arg1 *QMetaObject__SuperData) *QMetaObject__SuperData {
	__r, _, _ := purego.SyscallN(addr_QMetaObject__SuperData_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQMetaObject__SuperData(__r)
}

func (this *QMetaObject__SuperData) Direct() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QMetaObject__SuperData_direct_get, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QMetaObject__SuperData) SetDirect(direct *QMetaObject) {
	purego.SyscallN(addr_QMetaObject__SuperData_direct_set, this.h, uintptr(direct.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(direct)
}

func (this *QMetaObject__SuperData) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QMetaObject__SuperData_delete(this.h)
	}
	this.UnsafeRelease()
}
