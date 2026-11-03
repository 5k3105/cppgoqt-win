package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QScopedPropertyUpdateGroup_new uintptr
	pg_QScopedPropertyUpdateGroup_delete func(uintptr)
)

func init() {
	symaddr(&addr_QScopedPropertyUpdateGroup_new, "QScopedPropertyUpdateGroup_new")
	reg(&pg_QScopedPropertyUpdateGroup_delete, "QScopedPropertyUpdateGroup_delete")
}

type QScopedPropertyUpdateGroup struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QScopedPropertyUpdateGroup) UnsafeGoID() int64     { return this.vid }
func (this *QScopedPropertyUpdateGroup) UnsafeSetGoID(id int64) { this.vid = id }

func bindQScopedPropertyUpdateGroup(h uintptr) *QScopedPropertyUpdateGroup {
	if h == 0 {
		return nil
	}
	return &QScopedPropertyUpdateGroup{h: h}
}

func newownedQScopedPropertyUpdateGroup(h uintptr) *QScopedPropertyUpdateGroup {
	o := bindQScopedPropertyUpdateGroup(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QScopedPropertyUpdateGroup).Delete)
	}
	return o
}

func (this *QScopedPropertyUpdateGroup) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QScopedPropertyUpdateGroup) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQScopedPropertyUpdateGroup(h unsafe.Pointer) *QScopedPropertyUpdateGroup { return bindQScopedPropertyUpdateGroup(uintptr(h)) }
func UnsafeNewQScopedPropertyUpdateGroupOwned(h unsafe.Pointer) *QScopedPropertyUpdateGroup { return newownedQScopedPropertyUpdateGroup(uintptr(h)) }

func (this *QScopedPropertyUpdateGroup) UnsafeRelease() { this.h = 0 }

func NewQScopedPropertyUpdateGroup() *QScopedPropertyUpdateGroup {
	__r, _, _ := purego.SyscallN(addr_QScopedPropertyUpdateGroup_new)
	__o := newownedQScopedPropertyUpdateGroup(__r)
	return __o
}

func (this *QScopedPropertyUpdateGroup) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QScopedPropertyUpdateGroup_delete(this.h)
	}
	this.UnsafeRelease()
}
