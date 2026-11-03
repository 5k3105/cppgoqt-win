package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QGenericReturnArgument_base func(uintptr) uintptr
	addr_QGenericReturnArgument_new uintptr
	addr_QGenericReturnArgument_new2 uintptr
	pg_QGenericReturnArgument_delete func(uintptr)
)

func init() {
	reg(&pg_QGenericReturnArgument_base, "QGenericReturnArgument_base")
	symaddr(&addr_QGenericReturnArgument_new, "QGenericReturnArgument_new")
	symaddr(&addr_QGenericReturnArgument_new2, "QGenericReturnArgument_new2")
	reg(&pg_QGenericReturnArgument_delete, "QGenericReturnArgument_delete")
}

type QGenericReturnArgument struct {
	h uintptr
	*QGenericArgument
}

func bindQGenericReturnArgument(h uintptr) *QGenericReturnArgument {
	if h == 0 {
		return nil
	}
	return &QGenericReturnArgument{h: h, QGenericArgument: bindQGenericArgument(pg_QGenericReturnArgument_base(h))}
}

func newownedQGenericReturnArgument(h uintptr) *QGenericReturnArgument {
	o := bindQGenericReturnArgument(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QGenericReturnArgument).Delete)
	}
	return o
}

func (this *QGenericReturnArgument) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QGenericReturnArgument) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQGenericReturnArgument(h unsafe.Pointer) *QGenericReturnArgument { return bindQGenericReturnArgument(uintptr(h)) }
func UnsafeNewQGenericReturnArgumentOwned(h unsafe.Pointer) *QGenericReturnArgument { return newownedQGenericReturnArgument(uintptr(h)) }

func (this *QGenericReturnArgument) UnsafeRelease() {
	this.h = 0
	if this.QGenericArgument != nil {
		this.QGenericArgument.UnsafeRelease()
	}
}

func NewQGenericReturnArgument() *QGenericReturnArgument {
	__r, _, _ := purego.SyscallN(addr_QGenericReturnArgument_new)
	__o := newownedQGenericReturnArgument(__r)
	return __o
}

func NewQGenericReturnArgument2(aName string) *QGenericReturnArgument {
	aName_b := append([]byte(aName), 0)
	__r, _, _ := purego.SyscallN(addr_QGenericReturnArgument_new2, uintptr(unsafe.Pointer(&aName_b[0])))
	__o := newownedQGenericReturnArgument(__r)
	runtime.KeepAlive(aName_b)
	return __o
}

func (this *QGenericReturnArgument) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QGenericReturnArgument_delete(this.h)
	}
	this.UnsafeRelease()
}
