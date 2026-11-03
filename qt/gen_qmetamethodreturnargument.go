package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QMetaMethodReturnArgument_name_get uintptr
	addr_QMetaMethodReturnArgument_name_set uintptr
	pg_QMetaMethodReturnArgument_delete func(uintptr)
)

func init() {
	symaddr(&addr_QMetaMethodReturnArgument_name_get, "QMetaMethodReturnArgument_name_get")
	symaddr(&addr_QMetaMethodReturnArgument_name_set, "QMetaMethodReturnArgument_name_set")
	reg(&pg_QMetaMethodReturnArgument_delete, "QMetaMethodReturnArgument_delete")
}

type QMetaMethodReturnArgument struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QMetaMethodReturnArgument) UnsafeGoID() int64     { return this.vid }
func (this *QMetaMethodReturnArgument) UnsafeSetGoID(id int64) { this.vid = id }

func bindQMetaMethodReturnArgument(h uintptr) *QMetaMethodReturnArgument {
	if h == 0 {
		return nil
	}
	return &QMetaMethodReturnArgument{h: h}
}

func newownedQMetaMethodReturnArgument(h uintptr) *QMetaMethodReturnArgument {
	o := bindQMetaMethodReturnArgument(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QMetaMethodReturnArgument).Delete)
	}
	return o
}

func (this *QMetaMethodReturnArgument) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QMetaMethodReturnArgument) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQMetaMethodReturnArgument(h unsafe.Pointer) *QMetaMethodReturnArgument { return bindQMetaMethodReturnArgument(uintptr(h)) }
func UnsafeNewQMetaMethodReturnArgumentOwned(h unsafe.Pointer) *QMetaMethodReturnArgument { return newownedQMetaMethodReturnArgument(uintptr(h)) }

func (this *QMetaMethodReturnArgument) UnsafeRelease() { this.h = 0 }

func (this *QMetaMethodReturnArgument) Name() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QMetaMethodReturnArgument_name_get, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QMetaMethodReturnArgument) SetName(name string) {
	name_b := append([]byte(name), 0)
	purego.SyscallN(addr_QMetaMethodReturnArgument_name_set, this.h, uintptr(unsafe.Pointer(&name_b[0])))
	runtime.KeepAlive(this)
	runtime.KeepAlive(name_b)
}

func (this *QMetaMethodReturnArgument) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QMetaMethodReturnArgument_delete(this.h)
	}
	this.UnsafeRelease()
}
