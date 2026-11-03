package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QFlag_new uintptr
	addr_QFlag_new2 uintptr
	addr_QFlag_new3 uintptr
	addr_QFlag_new4 uintptr
	addr_QFlag_new5 uintptr
	addr_QFlag_new6 uintptr
	addr_QFlag_new7 uintptr
	addr_QFlag_ToInt uintptr
	addr_QFlag_ToUint uintptr
	pg_QFlag_delete func(uintptr)
)

func init() {
	symaddr(&addr_QFlag_new, "QFlag_new")
	symaddr(&addr_QFlag_new2, "QFlag_new2")
	symaddr(&addr_QFlag_new3, "QFlag_new3")
	symaddr(&addr_QFlag_new4, "QFlag_new4")
	symaddr(&addr_QFlag_new5, "QFlag_new5")
	symaddr(&addr_QFlag_new6, "QFlag_new6")
	symaddr(&addr_QFlag_new7, "QFlag_new7")
	symaddr(&addr_QFlag_ToInt, "QFlag_ToInt")
	symaddr(&addr_QFlag_ToUint, "QFlag_ToUint")
	reg(&pg_QFlag_delete, "QFlag_delete")
}

type QFlag struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QFlag) UnsafeGoID() int64     { return this.vid }
func (this *QFlag) UnsafeSetGoID(id int64) { this.vid = id }

func bindQFlag(h uintptr) *QFlag {
	if h == 0 {
		return nil
	}
	return &QFlag{h: h}
}

func newownedQFlag(h uintptr) *QFlag {
	o := bindQFlag(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QFlag).Delete)
	}
	return o
}

func (this *QFlag) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QFlag) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQFlag(h unsafe.Pointer) *QFlag { return bindQFlag(uintptr(h)) }
func UnsafeNewQFlagOwned(h unsafe.Pointer) *QFlag { return newownedQFlag(uintptr(h)) }

func (this *QFlag) UnsafeRelease() { this.h = 0 }

func NewQFlag(value int) *QFlag {
	__r, _, _ := purego.SyscallN(addr_QFlag_new, uintptr(value))
	__o := newownedQFlag(__r)
	return __o
}

func NewQFlag2(value int32) *QFlag {
	__r, _, _ := purego.SyscallN(addr_QFlag_new2, uintptr(value))
	__o := newownedQFlag(__r)
	return __o
}

func NewQFlag3(value uint32) *QFlag {
	__r, _, _ := purego.SyscallN(addr_QFlag_new3, uintptr(value))
	__o := newownedQFlag(__r)
	return __o
}

func NewQFlag4(value uint) *QFlag {
	__r, _, _ := purego.SyscallN(addr_QFlag_new4, uintptr(value))
	__o := newownedQFlag(__r)
	return __o
}

func NewQFlag5(value int16) *QFlag {
	__r, _, _ := purego.SyscallN(addr_QFlag_new5, uintptr(value))
	__o := newownedQFlag(__r)
	return __o
}

func NewQFlag6(value uint16) *QFlag {
	__r, _, _ := purego.SyscallN(addr_QFlag_new6, uintptr(value))
	__o := newownedQFlag(__r)
	return __o
}

func NewQFlag7(arg1 *QFlag) *QFlag {
	__r, _, _ := purego.SyscallN(addr_QFlag_new7, uintptr(arg1.UnsafePointer()))
	__o := newownedQFlag(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QFlag) ToInt() int {
	__r0, _, _ := purego.SyscallN(addr_QFlag_ToInt, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QFlag) ToUint() uint {
	__r0, _, _ := purego.SyscallN(addr_QFlag_ToUint, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint(__r)
}

func (this *QFlag) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QFlag_delete(this.h)
	}
	this.UnsafeRelease()
}
