package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QTextBlockUserData_OperatorAssign uintptr
	pg_QTextBlockUserData_delete func(uintptr)
)

func init() {
	symaddr(&addr_QTextBlockUserData_OperatorAssign, "QTextBlockUserData_OperatorAssign")
	reg(&pg_QTextBlockUserData_delete, "QTextBlockUserData_delete")
}

type QTextBlockUserData struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QTextBlockUserData) UnsafeGoID() int64     { return this.vid }
func (this *QTextBlockUserData) UnsafeSetGoID(id int64) { this.vid = id }

func bindQTextBlockUserData(h uintptr) *QTextBlockUserData {
	if h == 0 {
		return nil
	}
	return &QTextBlockUserData{h: h}
}

func newownedQTextBlockUserData(h uintptr) *QTextBlockUserData {
	o := bindQTextBlockUserData(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QTextBlockUserData).Delete)
	}
	return o
}

func (this *QTextBlockUserData) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QTextBlockUserData) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQTextBlockUserData(h unsafe.Pointer) *QTextBlockUserData { return bindQTextBlockUserData(uintptr(h)) }
func UnsafeNewQTextBlockUserDataOwned(h unsafe.Pointer) *QTextBlockUserData { return newownedQTextBlockUserData(uintptr(h)) }

func (this *QTextBlockUserData) UnsafeRelease() { this.h = 0 }

func (this *QTextBlockUserData) OperatorAssign(arg1 *QTextBlockUserData) *QTextBlockUserData {
	__r, _, _ := purego.SyscallN(addr_QTextBlockUserData_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQTextBlockUserData(__r)
}

func (this *QTextBlockUserData) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QTextBlockUserData_delete(this.h)
	}
	this.UnsafeRelease()
}
