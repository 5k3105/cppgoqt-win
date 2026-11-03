package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QFactoryInterface_keys uintptr
	addr_QFactoryInterface_OperatorAssign uintptr
	pg_QFactoryInterface__set_override func(uintptr, int32)
	pg_QFactoryInterface__Go_new func(int64) uintptr
	pg_QFactoryInterface_delete func(uintptr)
)

func init() {
	symaddr(&addr_QFactoryInterface_keys, "QFactoryInterface_keys")
	symaddr(&addr_QFactoryInterface_OperatorAssign, "QFactoryInterface_OperatorAssign")
	reg(&pg_QFactoryInterface__set_override, "QFactoryInterface__set_override")
	reg(&pg_QFactoryInterface__Go_new, "QFactoryInterface__Go_new")
	reg(&pg_QFactoryInterface_delete, "QFactoryInterface_delete")
}

type QFactoryInterface struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QFactoryInterface) UnsafeGoID() int64     { return this.vid }
func (this *QFactoryInterface) UnsafeSetGoID(id int64) { this.vid = id }

func bindQFactoryInterface(h uintptr) *QFactoryInterface {
	if h == 0 {
		return nil
	}
	return &QFactoryInterface{h: h}
}

func newownedQFactoryInterface(h uintptr) *QFactoryInterface {
	o := bindQFactoryInterface(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QFactoryInterface).Delete)
	}
	return o
}

func (this *QFactoryInterface) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QFactoryInterface) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQFactoryInterface(h unsafe.Pointer) *QFactoryInterface { return bindQFactoryInterface(uintptr(h)) }
func UnsafeNewQFactoryInterfaceOwned(h unsafe.Pointer) *QFactoryInterface { return newownedQFactoryInterface(uintptr(h)) }

func (this *QFactoryInterface) UnsafeRelease() { this.h = 0 }

func (this *QFactoryInterface) Keys() []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFactoryInterface_keys, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]string, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = pgGoStr(__sl[__i])
		pgFree(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QFactoryInterface) OperatorAssign(arg1 *QFactoryInterface) *QFactoryInterface {
	__r, _, _ := purego.SyscallN(addr_QFactoryInterface_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQFactoryInterface(__r)
}

func QFactoryInterface_virtIdx(slot int64) int32 {
	switch slot {
	case 6348254508799049261:
		return 0
	}
	return -1
}

func QFactoryInterface_enable(h uintptr, idx int32) { pg_QFactoryInterface__set_override(h, idx) }

func NewQFactoryInterface() *QFactoryInterface {
	id := nextVirtID()
	o := newownedQFactoryInterface(pg_QFactoryInterface__Go_new(id))
	o.UnsafeSetGoID(id)
	registerInstance(id, o.h, QFactoryInterface_virtIdx, QFactoryInterface_enable)
	return o
}

func (this *QFactoryInterface) OnKeys(fn func() []string) {
	registerVirt(this.UnsafeGoID(), 6348254508799049261, func() (uintptr, int64) {
		__rs := fn()
		if len(__rs) == 0 {
			return 0, 0
		}
		__a := pgMalloc(uintptr(len(__rs)) * unsafe.Sizeof(uintptr(0)))
		__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__a)), len(__rs))
		for __i, __e := range __rs {
			__sl[__i] = pgCString(__e)
		}
		return __a, int64(len(__rs))
	}) // slot "keys"
}

func (this *QFactoryInterface) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QFactoryInterface_delete(this.h)
	}
	this.UnsafeRelease()
}
