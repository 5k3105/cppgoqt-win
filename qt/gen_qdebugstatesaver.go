package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QDebugStateSaver_new uintptr
	pg_QDebugStateSaver_delete func(uintptr)
)

func init() {
	symaddr(&addr_QDebugStateSaver_new, "QDebugStateSaver_new")
	reg(&pg_QDebugStateSaver_delete, "QDebugStateSaver_delete")
}

type QDebugStateSaver struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QDebugStateSaver) UnsafeGoID() int64     { return this.vid }
func (this *QDebugStateSaver) UnsafeSetGoID(id int64) { this.vid = id }

func bindQDebugStateSaver(h uintptr) *QDebugStateSaver {
	if h == 0 {
		return nil
	}
	return &QDebugStateSaver{h: h}
}

func newownedQDebugStateSaver(h uintptr) *QDebugStateSaver {
	o := bindQDebugStateSaver(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QDebugStateSaver).Delete)
	}
	return o
}

func (this *QDebugStateSaver) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QDebugStateSaver) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQDebugStateSaver(h unsafe.Pointer) *QDebugStateSaver { return bindQDebugStateSaver(uintptr(h)) }
func UnsafeNewQDebugStateSaverOwned(h unsafe.Pointer) *QDebugStateSaver { return newownedQDebugStateSaver(uintptr(h)) }

func (this *QDebugStateSaver) UnsafeRelease() { this.h = 0 }

func NewQDebugStateSaver(dbg *QDebug) *QDebugStateSaver {
	__r, _, _ := purego.SyscallN(addr_QDebugStateSaver_new, uintptr(dbg.UnsafePointer()))
	__o := newownedQDebugStateSaver(__r)
	runtime.KeepAlive(dbg)
	return __o
}

func (this *QDebugStateSaver) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QDebugStateSaver_delete(this.h)
	}
	this.UnsafeRelease()
}
