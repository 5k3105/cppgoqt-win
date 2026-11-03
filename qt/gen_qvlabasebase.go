package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QVLABaseBase_capacity uintptr
	addr_QVLABaseBase_size uintptr
	addr_QVLABaseBase_empty uintptr
)

func init() {
	symaddr(&addr_QVLABaseBase_capacity, "QVLABaseBase_capacity")
	symaddr(&addr_QVLABaseBase_size, "QVLABaseBase_size")
	symaddr(&addr_QVLABaseBase_empty, "QVLABaseBase_empty")
}

type QVLABaseBase struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QVLABaseBase) UnsafeGoID() int64     { return this.vid }
func (this *QVLABaseBase) UnsafeSetGoID(id int64) { this.vid = id }

func bindQVLABaseBase(h uintptr) *QVLABaseBase {
	if h == 0 {
		return nil
	}
	return &QVLABaseBase{h: h}
}

func newownedQVLABaseBase(h uintptr) *QVLABaseBase { return bindQVLABaseBase(h) }

func (this *QVLABaseBase) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QVLABaseBase) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQVLABaseBase(h unsafe.Pointer) *QVLABaseBase { return bindQVLABaseBase(uintptr(h)) }
func UnsafeNewQVLABaseBaseOwned(h unsafe.Pointer) *QVLABaseBase { return newownedQVLABaseBase(uintptr(h)) }

func (this *QVLABaseBase) UnsafeRelease() { this.h = 0 }

func (this *QVLABaseBase) Capacity() int64 {
	__r0, _, _ := purego.SyscallN(addr_QVLABaseBase_capacity, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QVLABaseBase) Size() int64 {
	__r0, _, _ := purego.SyscallN(addr_QVLABaseBase_size, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QVLABaseBase) Empty() bool {
	__r0, _, _ := purego.SyscallN(addr_QVLABaseBase_empty, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QVLABaseBase) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	this.UnsafeRelease()
}
