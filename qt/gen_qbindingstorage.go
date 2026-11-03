package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QBindingStorage_new uintptr
	addr_QBindingStorage_isEmpty uintptr
	addr_QBindingStorage_isValid uintptr
	addr_QBindingStorage_registerDependency uintptr
	pg_QBindingStorage_delete func(uintptr)
)

func init() {
	symaddr(&addr_QBindingStorage_new, "QBindingStorage_new")
	symaddr(&addr_QBindingStorage_isEmpty, "QBindingStorage_isEmpty")
	symaddr(&addr_QBindingStorage_isValid, "QBindingStorage_isValid")
	symaddr(&addr_QBindingStorage_registerDependency, "QBindingStorage_registerDependency")
	reg(&pg_QBindingStorage_delete, "QBindingStorage_delete")
}

type QBindingStorage struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QBindingStorage) UnsafeGoID() int64     { return this.vid }
func (this *QBindingStorage) UnsafeSetGoID(id int64) { this.vid = id }

func bindQBindingStorage(h uintptr) *QBindingStorage {
	if h == 0 {
		return nil
	}
	return &QBindingStorage{h: h}
}

func newownedQBindingStorage(h uintptr) *QBindingStorage {
	o := bindQBindingStorage(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QBindingStorage).Delete)
	}
	return o
}

func (this *QBindingStorage) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QBindingStorage) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQBindingStorage(h unsafe.Pointer) *QBindingStorage { return bindQBindingStorage(uintptr(h)) }
func UnsafeNewQBindingStorageOwned(h unsafe.Pointer) *QBindingStorage { return newownedQBindingStorage(uintptr(h)) }

func (this *QBindingStorage) UnsafeRelease() { this.h = 0 }

func NewQBindingStorage() *QBindingStorage {
	__r, _, _ := purego.SyscallN(addr_QBindingStorage_new)
	__o := newownedQBindingStorage(__r)
	return __o
}

func (this *QBindingStorage) IsEmpty() bool {
	__r0, _, _ := purego.SyscallN(addr_QBindingStorage_isEmpty, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QBindingStorage) IsValid() bool {
	__r0, _, _ := purego.SyscallN(addr_QBindingStorage_isValid, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QBindingStorage) RegisterDependency(data *QUntypedPropertyData) {
	purego.SyscallN(addr_QBindingStorage_registerDependency, this.h, uintptr(data.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(data)
}

func (this *QBindingStorage) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QBindingStorage_delete(this.h)
	}
	this.UnsafeRelease()
}
