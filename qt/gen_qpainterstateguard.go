package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QPainterStateGuard_new uintptr
	addr_QPainterStateGuard_new2 uintptr
	addr_QPainterStateGuard_swap uintptr
	addr_QPainterStateGuard_save uintptr
	addr_QPainterStateGuard_restore uintptr
	pg_QPainterStateGuard_delete func(uintptr)
)

func init() {
	symaddr(&addr_QPainterStateGuard_new, "QPainterStateGuard_new")
	symaddr(&addr_QPainterStateGuard_new2, "QPainterStateGuard_new2")
	symaddr(&addr_QPainterStateGuard_swap, "QPainterStateGuard_swap")
	symaddr(&addr_QPainterStateGuard_save, "QPainterStateGuard_save")
	symaddr(&addr_QPainterStateGuard_restore, "QPainterStateGuard_restore")
	reg(&pg_QPainterStateGuard_delete, "QPainterStateGuard_delete")
}

type QPainterStateGuard struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QPainterStateGuard) UnsafeGoID() int64     { return this.vid }
func (this *QPainterStateGuard) UnsafeSetGoID(id int64) { this.vid = id }

func bindQPainterStateGuard(h uintptr) *QPainterStateGuard {
	if h == 0 {
		return nil
	}
	return &QPainterStateGuard{h: h}
}

func newownedQPainterStateGuard(h uintptr) *QPainterStateGuard {
	o := bindQPainterStateGuard(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QPainterStateGuard).Delete)
	}
	return o
}

func (this *QPainterStateGuard) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QPainterStateGuard) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQPainterStateGuard(h unsafe.Pointer) *QPainterStateGuard { return bindQPainterStateGuard(uintptr(h)) }
func UnsafeNewQPainterStateGuardOwned(h unsafe.Pointer) *QPainterStateGuard { return newownedQPainterStateGuard(uintptr(h)) }

func (this *QPainterStateGuard) UnsafeRelease() { this.h = 0 }

func NewQPainterStateGuard(painter *QPainter) *QPainterStateGuard {
	__r, _, _ := purego.SyscallN(addr_QPainterStateGuard_new, uintptr(painter.UnsafePointer()))
	__o := newownedQPainterStateGuard(__r)
	runtime.KeepAlive(painter)
	return __o
}

func NewQPainterStateGuard2(painter *QPainter, state QPainterStateGuard__InitialState) *QPainterStateGuard {
	__r, _, _ := purego.SyscallN(addr_QPainterStateGuard_new2, uintptr(painter.UnsafePointer()), uintptr(state))
	__o := newownedQPainterStateGuard(__r)
	runtime.KeepAlive(painter)
	return __o
}

func (this *QPainterStateGuard) Swap(other *QPainterStateGuard) {
	purego.SyscallN(addr_QPainterStateGuard_swap, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
}

func (this *QPainterStateGuard) Save() {
	purego.SyscallN(addr_QPainterStateGuard_save, this.h)
	runtime.KeepAlive(this)
}

func (this *QPainterStateGuard) Restore() {
	purego.SyscallN(addr_QPainterStateGuard_restore, this.h)
	runtime.KeepAlive(this)
}

func (this *QPainterStateGuard) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QPainterStateGuard_delete(this.h)
	}
	this.UnsafeRelease()
}
