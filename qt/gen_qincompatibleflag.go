package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QIncompatibleFlag_new uintptr
	addr_QIncompatibleFlag_new2 uintptr
	addr_QIncompatibleFlag_ToInt uintptr
	pg_QIncompatibleFlag_delete func(uintptr)
)

func init() {
	symaddr(&addr_QIncompatibleFlag_new, "QIncompatibleFlag_new")
	symaddr(&addr_QIncompatibleFlag_new2, "QIncompatibleFlag_new2")
	symaddr(&addr_QIncompatibleFlag_ToInt, "QIncompatibleFlag_ToInt")
	reg(&pg_QIncompatibleFlag_delete, "QIncompatibleFlag_delete")
}

type QIncompatibleFlag struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QIncompatibleFlag) UnsafeGoID() int64     { return this.vid }
func (this *QIncompatibleFlag) UnsafeSetGoID(id int64) { this.vid = id }

func bindQIncompatibleFlag(h uintptr) *QIncompatibleFlag {
	if h == 0 {
		return nil
	}
	return &QIncompatibleFlag{h: h}
}

func newownedQIncompatibleFlag(h uintptr) *QIncompatibleFlag {
	o := bindQIncompatibleFlag(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QIncompatibleFlag).Delete)
	}
	return o
}

func (this *QIncompatibleFlag) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QIncompatibleFlag) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQIncompatibleFlag(h unsafe.Pointer) *QIncompatibleFlag { return bindQIncompatibleFlag(uintptr(h)) }
func UnsafeNewQIncompatibleFlagOwned(h unsafe.Pointer) *QIncompatibleFlag { return newownedQIncompatibleFlag(uintptr(h)) }

func (this *QIncompatibleFlag) UnsafeRelease() { this.h = 0 }

func NewQIncompatibleFlag(i int) *QIncompatibleFlag {
	__r, _, _ := purego.SyscallN(addr_QIncompatibleFlag_new, uintptr(i))
	__o := newownedQIncompatibleFlag(__r)
	return __o
}

func NewQIncompatibleFlag2(arg1 *QIncompatibleFlag) *QIncompatibleFlag {
	__r, _, _ := purego.SyscallN(addr_QIncompatibleFlag_new2, uintptr(arg1.UnsafePointer()))
	__o := newownedQIncompatibleFlag(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QIncompatibleFlag) ToInt() int {
	__r0, _, _ := purego.SyscallN(addr_QIncompatibleFlag_ToInt, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QIncompatibleFlag) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QIncompatibleFlag_delete(this.h)
	}
	this.UnsafeRelease()
}
