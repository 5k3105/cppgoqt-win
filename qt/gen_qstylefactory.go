package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QStyleFactory_keys uintptr
	addr_QStyleFactory_create uintptr
	pg_QStyleFactory_delete func(uintptr)
)

func init() {
	symaddr(&addr_QStyleFactory_keys, "QStyleFactory_keys")
	symaddr(&addr_QStyleFactory_create, "QStyleFactory_create")
	reg(&pg_QStyleFactory_delete, "QStyleFactory_delete")
}

type QStyleFactory struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QStyleFactory) UnsafeGoID() int64     { return this.vid }
func (this *QStyleFactory) UnsafeSetGoID(id int64) { this.vid = id }

func bindQStyleFactory(h uintptr) *QStyleFactory {
	if h == 0 {
		return nil
	}
	return &QStyleFactory{h: h}
}

func newownedQStyleFactory(h uintptr) *QStyleFactory {
	o := bindQStyleFactory(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QStyleFactory).Delete)
	}
	return o
}

func (this *QStyleFactory) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QStyleFactory) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQStyleFactory(h unsafe.Pointer) *QStyleFactory { return bindQStyleFactory(uintptr(h)) }
func UnsafeNewQStyleFactoryOwned(h unsafe.Pointer) *QStyleFactory { return newownedQStyleFactory(uintptr(h)) }

func (this *QStyleFactory) UnsafeRelease() { this.h = 0 }

func QStyleFactory_Keys() []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStyleFactory_keys, uintptr(unsafe.Pointer(&__n)))
	__out := make([]string, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = pgGoStr(__sl[__i])
		pgFree(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QStyleFactory_Create(arg1 string) *QStyle {
	var arg1_p uintptr
	if len(arg1) > 0 {
		arg1_p = uintptr(unsafe.Pointer(unsafe.StringData(arg1)))
	}
	__r, _, _ := purego.SyscallN(addr_QStyleFactory_create, arg1_p, uintptr(len(arg1)))
	runtime.KeepAlive(arg1)
	return bindQStyle(__r)
}

func (this *QStyleFactory) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QStyleFactory_delete(this.h)
	}
	this.UnsafeRelease()
}
