package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QMetaObject__Connection_new uintptr
	addr_QMetaObject__Connection_new2 uintptr
	addr_QMetaObject__Connection_OperatorAssign uintptr
	addr_QMetaObject__Connection_swap uintptr
	pg_QMetaObject__Connection_delete func(uintptr)
)

func init() {
	symaddr(&addr_QMetaObject__Connection_new, "QMetaObject__Connection_new")
	symaddr(&addr_QMetaObject__Connection_new2, "QMetaObject__Connection_new2")
	symaddr(&addr_QMetaObject__Connection_OperatorAssign, "QMetaObject__Connection_OperatorAssign")
	symaddr(&addr_QMetaObject__Connection_swap, "QMetaObject__Connection_swap")
	reg(&pg_QMetaObject__Connection_delete, "QMetaObject__Connection_delete")
}

type QMetaObject__Connection struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QMetaObject__Connection) UnsafeGoID() int64     { return this.vid }
func (this *QMetaObject__Connection) UnsafeSetGoID(id int64) { this.vid = id }

func bindQMetaObject__Connection(h uintptr) *QMetaObject__Connection {
	if h == 0 {
		return nil
	}
	return &QMetaObject__Connection{h: h}
}

func newownedQMetaObject__Connection(h uintptr) *QMetaObject__Connection {
	o := bindQMetaObject__Connection(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QMetaObject__Connection).Delete)
	}
	return o
}

func (this *QMetaObject__Connection) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QMetaObject__Connection) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQMetaObject__Connection(h unsafe.Pointer) *QMetaObject__Connection { return bindQMetaObject__Connection(uintptr(h)) }
func UnsafeNewQMetaObject__ConnectionOwned(h unsafe.Pointer) *QMetaObject__Connection { return newownedQMetaObject__Connection(uintptr(h)) }

func (this *QMetaObject__Connection) UnsafeRelease() { this.h = 0 }

func NewQMetaObject__Connection() *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QMetaObject__Connection_new)
	__o := newownedQMetaObject__Connection(__r)
	return __o
}

func NewQMetaObject__Connection2(other *QMetaObject__Connection) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QMetaObject__Connection_new2, uintptr(other.UnsafePointer()))
	__o := newownedQMetaObject__Connection(__r)
	runtime.KeepAlive(other)
	return __o
}

func (this *QMetaObject__Connection) OperatorAssign(other *QMetaObject__Connection) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QMetaObject__Connection_OperatorAssign, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return bindQMetaObject__Connection(__r)
}

func (this *QMetaObject__Connection) Swap(other *QMetaObject__Connection) {
	purego.SyscallN(addr_QMetaObject__Connection_swap, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
}

func (this *QMetaObject__Connection) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QMetaObject__Connection_delete(this.h)
	}
	this.UnsafeRelease()
}
