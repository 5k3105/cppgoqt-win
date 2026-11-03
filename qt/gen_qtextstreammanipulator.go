package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QTextStreamManipulator_exec uintptr
	pg_QTextStreamManipulator_delete func(uintptr)
)

func init() {
	symaddr(&addr_QTextStreamManipulator_exec, "QTextStreamManipulator_exec")
	reg(&pg_QTextStreamManipulator_delete, "QTextStreamManipulator_delete")
}

type QTextStreamManipulator struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QTextStreamManipulator) UnsafeGoID() int64     { return this.vid }
func (this *QTextStreamManipulator) UnsafeSetGoID(id int64) { this.vid = id }

func bindQTextStreamManipulator(h uintptr) *QTextStreamManipulator {
	if h == 0 {
		return nil
	}
	return &QTextStreamManipulator{h: h}
}

func newownedQTextStreamManipulator(h uintptr) *QTextStreamManipulator {
	o := bindQTextStreamManipulator(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QTextStreamManipulator).Delete)
	}
	return o
}

func (this *QTextStreamManipulator) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QTextStreamManipulator) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQTextStreamManipulator(h unsafe.Pointer) *QTextStreamManipulator { return bindQTextStreamManipulator(uintptr(h)) }
func UnsafeNewQTextStreamManipulatorOwned(h unsafe.Pointer) *QTextStreamManipulator { return newownedQTextStreamManipulator(uintptr(h)) }

func (this *QTextStreamManipulator) UnsafeRelease() { this.h = 0 }

func (this *QTextStreamManipulator) Exec(s *QTextStream) {
	purego.SyscallN(addr_QTextStreamManipulator_exec, this.h, uintptr(s.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
}

func (this *QTextStreamManipulator) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QTextStreamManipulator_delete(this.h)
	}
	this.UnsafeRelease()
}
