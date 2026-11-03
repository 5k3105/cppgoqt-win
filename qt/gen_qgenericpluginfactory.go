package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QGenericPluginFactory_keys uintptr
	addr_QGenericPluginFactory_create uintptr
	pg_QGenericPluginFactory_delete func(uintptr)
)

func init() {
	symaddr(&addr_QGenericPluginFactory_keys, "QGenericPluginFactory_keys")
	symaddr(&addr_QGenericPluginFactory_create, "QGenericPluginFactory_create")
	reg(&pg_QGenericPluginFactory_delete, "QGenericPluginFactory_delete")
}

type QGenericPluginFactory struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QGenericPluginFactory) UnsafeGoID() int64     { return this.vid }
func (this *QGenericPluginFactory) UnsafeSetGoID(id int64) { this.vid = id }

func bindQGenericPluginFactory(h uintptr) *QGenericPluginFactory {
	if h == 0 {
		return nil
	}
	return &QGenericPluginFactory{h: h}
}

func newownedQGenericPluginFactory(h uintptr) *QGenericPluginFactory {
	o := bindQGenericPluginFactory(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QGenericPluginFactory).Delete)
	}
	return o
}

func (this *QGenericPluginFactory) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QGenericPluginFactory) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQGenericPluginFactory(h unsafe.Pointer) *QGenericPluginFactory { return bindQGenericPluginFactory(uintptr(h)) }
func UnsafeNewQGenericPluginFactoryOwned(h unsafe.Pointer) *QGenericPluginFactory { return newownedQGenericPluginFactory(uintptr(h)) }

func (this *QGenericPluginFactory) UnsafeRelease() { this.h = 0 }

func QGenericPluginFactory_Keys() []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGenericPluginFactory_keys, uintptr(unsafe.Pointer(&__n)))
	__out := make([]string, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = pgGoStr(__sl[__i])
		pgFree(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QGenericPluginFactory_Create(arg1 string, arg2 string) *QObject {
	var arg1_p uintptr
	if len(arg1) > 0 {
		arg1_p = uintptr(unsafe.Pointer(unsafe.StringData(arg1)))
	}
	var arg2_p uintptr
	if len(arg2) > 0 {
		arg2_p = uintptr(unsafe.Pointer(unsafe.StringData(arg2)))
	}
	__r, _, _ := purego.SyscallN(addr_QGenericPluginFactory_create, arg1_p, uintptr(len(arg1)), arg2_p, uintptr(len(arg2)))
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
	return bindQObject(__r)
}

func (this *QGenericPluginFactory) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QGenericPluginFactory_delete(this.h)
	}
	this.UnsafeRelease()
}
