package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QTextFrameLayoutData_OperatorAssign uintptr
	pg_QTextFrameLayoutData_delete func(uintptr)
)

func init() {
	symaddr(&addr_QTextFrameLayoutData_OperatorAssign, "QTextFrameLayoutData_OperatorAssign")
	reg(&pg_QTextFrameLayoutData_delete, "QTextFrameLayoutData_delete")
}

type QTextFrameLayoutData struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QTextFrameLayoutData) UnsafeGoID() int64     { return this.vid }
func (this *QTextFrameLayoutData) UnsafeSetGoID(id int64) { this.vid = id }

func bindQTextFrameLayoutData(h uintptr) *QTextFrameLayoutData {
	if h == 0 {
		return nil
	}
	return &QTextFrameLayoutData{h: h}
}

func newownedQTextFrameLayoutData(h uintptr) *QTextFrameLayoutData {
	o := bindQTextFrameLayoutData(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QTextFrameLayoutData).Delete)
	}
	return o
}

func (this *QTextFrameLayoutData) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QTextFrameLayoutData) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQTextFrameLayoutData(h unsafe.Pointer) *QTextFrameLayoutData { return bindQTextFrameLayoutData(uintptr(h)) }
func UnsafeNewQTextFrameLayoutDataOwned(h unsafe.Pointer) *QTextFrameLayoutData { return newownedQTextFrameLayoutData(uintptr(h)) }

func (this *QTextFrameLayoutData) UnsafeRelease() { this.h = 0 }

func (this *QTextFrameLayoutData) OperatorAssign(arg1 *QTextFrameLayoutData) *QTextFrameLayoutData {
	__r, _, _ := purego.SyscallN(addr_QTextFrameLayoutData_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQTextFrameLayoutData(__r)
}

func (this *QTextFrameLayoutData) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QTextFrameLayoutData_delete(this.h)
	}
	this.UnsafeRelease()
}
