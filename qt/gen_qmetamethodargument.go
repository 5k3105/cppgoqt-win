package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QMetaMethodArgument_name_get uintptr
	addr_QMetaMethodArgument_name_set uintptr
	pg_QMetaMethodArgument_delete func(uintptr)
)

func init() {
	symaddr(&addr_QMetaMethodArgument_name_get, "QMetaMethodArgument_name_get")
	symaddr(&addr_QMetaMethodArgument_name_set, "QMetaMethodArgument_name_set")
	reg(&pg_QMetaMethodArgument_delete, "QMetaMethodArgument_delete")
}

type QMetaMethodArgument struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QMetaMethodArgument) UnsafeGoID() int64     { return this.vid }
func (this *QMetaMethodArgument) UnsafeSetGoID(id int64) { this.vid = id }

func bindQMetaMethodArgument(h uintptr) *QMetaMethodArgument {
	if h == 0 {
		return nil
	}
	return &QMetaMethodArgument{h: h}
}

func newownedQMetaMethodArgument(h uintptr) *QMetaMethodArgument {
	o := bindQMetaMethodArgument(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QMetaMethodArgument).Delete)
	}
	return o
}

func (this *QMetaMethodArgument) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QMetaMethodArgument) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQMetaMethodArgument(h unsafe.Pointer) *QMetaMethodArgument { return bindQMetaMethodArgument(uintptr(h)) }
func UnsafeNewQMetaMethodArgumentOwned(h unsafe.Pointer) *QMetaMethodArgument { return newownedQMetaMethodArgument(uintptr(h)) }

func (this *QMetaMethodArgument) UnsafeRelease() { this.h = 0 }

func (this *QMetaMethodArgument) Name() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QMetaMethodArgument_name_get, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QMetaMethodArgument) SetName(name string) {
	name_b := append([]byte(name), 0)
	purego.SyscallN(addr_QMetaMethodArgument_name_set, this.h, uintptr(unsafe.Pointer(&name_b[0])))
	runtime.KeepAlive(this)
	runtime.KeepAlive(name_b)
}

func (this *QMetaMethodArgument) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QMetaMethodArgument_delete(this.h)
	}
	this.UnsafeRelease()
}
