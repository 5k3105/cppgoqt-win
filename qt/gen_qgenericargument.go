package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QGenericArgument_new uintptr
	addr_QGenericArgument_new2 uintptr
	addr_QGenericArgument_new3 uintptr
	addr_QGenericArgument_name uintptr
	pg_QGenericArgument_delete func(uintptr)
)

func init() {
	symaddr(&addr_QGenericArgument_new, "QGenericArgument_new")
	symaddr(&addr_QGenericArgument_new2, "QGenericArgument_new2")
	symaddr(&addr_QGenericArgument_new3, "QGenericArgument_new3")
	symaddr(&addr_QGenericArgument_name, "QGenericArgument_name")
	reg(&pg_QGenericArgument_delete, "QGenericArgument_delete")
}

type QGenericArgument struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QGenericArgument) UnsafeGoID() int64     { return this.vid }
func (this *QGenericArgument) UnsafeSetGoID(id int64) { this.vid = id }

func bindQGenericArgument(h uintptr) *QGenericArgument {
	if h == 0 {
		return nil
	}
	return &QGenericArgument{h: h}
}

func newownedQGenericArgument(h uintptr) *QGenericArgument {
	o := bindQGenericArgument(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QGenericArgument).Delete)
	}
	return o
}

func (this *QGenericArgument) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QGenericArgument) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQGenericArgument(h unsafe.Pointer) *QGenericArgument { return bindQGenericArgument(uintptr(h)) }
func UnsafeNewQGenericArgumentOwned(h unsafe.Pointer) *QGenericArgument { return newownedQGenericArgument(uintptr(h)) }

func (this *QGenericArgument) UnsafeRelease() { this.h = 0 }

func NewQGenericArgument() *QGenericArgument {
	__r, _, _ := purego.SyscallN(addr_QGenericArgument_new)
	__o := newownedQGenericArgument(__r)
	return __o
}

func NewQGenericArgument2(arg1 *QGenericArgument) *QGenericArgument {
	__r, _, _ := purego.SyscallN(addr_QGenericArgument_new2, uintptr(arg1.UnsafePointer()))
	__o := newownedQGenericArgument(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func NewQGenericArgument3(aName string) *QGenericArgument {
	aName_b := append([]byte(aName), 0)
	__r, _, _ := purego.SyscallN(addr_QGenericArgument_new3, uintptr(unsafe.Pointer(&aName_b[0])))
	__o := newownedQGenericArgument(__r)
	runtime.KeepAlive(aName_b)
	return __o
}

func (this *QGenericArgument) Name() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGenericArgument_name, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QGenericArgument) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QGenericArgument_delete(this.h)
	}
	this.UnsafeRelease()
}
