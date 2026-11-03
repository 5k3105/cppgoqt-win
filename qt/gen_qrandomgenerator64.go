package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QRandomGenerator64_base func(uintptr) uintptr
	addr_QRandomGenerator64_new uintptr
	addr_QRandomGenerator64_new2 uintptr
	addr_QRandomGenerator64_new3 uintptr
	addr_QRandomGenerator64_new4 uintptr
	addr_QRandomGenerator64_generate uintptr
	addr_QRandomGenerator64_OperatorCall uintptr
	addr_QRandomGenerator64_discard uintptr
	addr_QRandomGenerator64_min uintptr
	addr_QRandomGenerator64_max uintptr
	addr_QRandomGenerator64_system uintptr
	addr_QRandomGenerator64_global uintptr
	addr_QRandomGenerator64_securelySeeded uintptr
	addr_QRandomGenerator64_OperatorAssign uintptr
	pg_QRandomGenerator64_delete func(uintptr)
)

func init() {
	reg(&pg_QRandomGenerator64_base, "QRandomGenerator64_base")
	symaddr(&addr_QRandomGenerator64_new, "QRandomGenerator64_new")
	symaddr(&addr_QRandomGenerator64_new2, "QRandomGenerator64_new2")
	symaddr(&addr_QRandomGenerator64_new3, "QRandomGenerator64_new3")
	symaddr(&addr_QRandomGenerator64_new4, "QRandomGenerator64_new4")
	symaddr(&addr_QRandomGenerator64_generate, "QRandomGenerator64_generate")
	symaddr(&addr_QRandomGenerator64_OperatorCall, "QRandomGenerator64_OperatorCall")
	symaddr(&addr_QRandomGenerator64_discard, "QRandomGenerator64_discard")
	symaddr(&addr_QRandomGenerator64_min, "QRandomGenerator64_min")
	symaddr(&addr_QRandomGenerator64_max, "QRandomGenerator64_max")
	symaddr(&addr_QRandomGenerator64_system, "QRandomGenerator64_system")
	symaddr(&addr_QRandomGenerator64_global, "QRandomGenerator64_global")
	symaddr(&addr_QRandomGenerator64_securelySeeded, "QRandomGenerator64_securelySeeded")
	symaddr(&addr_QRandomGenerator64_OperatorAssign, "QRandomGenerator64_OperatorAssign")
	reg(&pg_QRandomGenerator64_delete, "QRandomGenerator64_delete")
}

type QRandomGenerator64 struct {
	h uintptr
	*QRandomGenerator
}

func bindQRandomGenerator64(h uintptr) *QRandomGenerator64 {
	if h == 0 {
		return nil
	}
	return &QRandomGenerator64{h: h, QRandomGenerator: bindQRandomGenerator(pg_QRandomGenerator64_base(h))}
}

func newownedQRandomGenerator64(h uintptr) *QRandomGenerator64 {
	o := bindQRandomGenerator64(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QRandomGenerator64).Delete)
	}
	return o
}

func (this *QRandomGenerator64) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QRandomGenerator64) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQRandomGenerator64(h unsafe.Pointer) *QRandomGenerator64 { return bindQRandomGenerator64(uintptr(h)) }
func UnsafeNewQRandomGenerator64Owned(h unsafe.Pointer) *QRandomGenerator64 { return newownedQRandomGenerator64(uintptr(h)) }

func (this *QRandomGenerator64) UnsafeRelease() {
	this.h = 0
	if this.QRandomGenerator != nil {
		this.QRandomGenerator.UnsafeRelease()
	}
}

func NewQRandomGenerator64() *QRandomGenerator64 {
	__r, _, _ := purego.SyscallN(addr_QRandomGenerator64_new)
	__o := newownedQRandomGenerator64(__r)
	return __o
}

func NewQRandomGenerator642(other *QRandomGenerator) *QRandomGenerator64 {
	__r, _, _ := purego.SyscallN(addr_QRandomGenerator64_new2, uintptr(other.UnsafePointer()))
	__o := newownedQRandomGenerator64(__r)
	runtime.KeepAlive(other)
	return __o
}

func NewQRandomGenerator643(arg1 *QRandomGenerator64) *QRandomGenerator64 {
	__r, _, _ := purego.SyscallN(addr_QRandomGenerator64_new3, uintptr(arg1.UnsafePointer()))
	__o := newownedQRandomGenerator64(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func NewQRandomGenerator644(seedValue uint) *QRandomGenerator64 {
	__r, _, _ := purego.SyscallN(addr_QRandomGenerator64_new4, uintptr(seedValue))
	__o := newownedQRandomGenerator64(__r)
	return __o
}

func (this *QRandomGenerator64) Generate() uint64 {
	__r0, _, _ := purego.SyscallN(addr_QRandomGenerator64_generate, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return uint64(__r)
}

func (this *QRandomGenerator64) OperatorCall() uint64 {
	__r0, _, _ := purego.SyscallN(addr_QRandomGenerator64_OperatorCall, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return uint64(__r)
}

func (this *QRandomGenerator64) Discard(z uint64) {
	purego.SyscallN(addr_QRandomGenerator64_discard, this.h, uintptr(z))
	runtime.KeepAlive(this)
}

func QRandomGenerator64_Min() uint64 {
	__r0, _, _ := purego.SyscallN(addr_QRandomGenerator64_min)
	__r := int64(__r0)
	return uint64(__r)
}

func QRandomGenerator64_Max() uint64 {
	__r0, _, _ := purego.SyscallN(addr_QRandomGenerator64_max)
	__r := int64(__r0)
	return uint64(__r)
}

func QRandomGenerator64_System() *QRandomGenerator64 {
	__r, _, _ := purego.SyscallN(addr_QRandomGenerator64_system)
	return bindQRandomGenerator64(__r)
}

func QRandomGenerator64_Global() *QRandomGenerator64 {
	__r, _, _ := purego.SyscallN(addr_QRandomGenerator64_global)
	return bindQRandomGenerator64(__r)
}

func QRandomGenerator64_SecurelySeeded() *QRandomGenerator64 {
	__r, _, _ := purego.SyscallN(addr_QRandomGenerator64_securelySeeded)
	return newownedQRandomGenerator64(__r)
}

func (this *QRandomGenerator64) OperatorAssign(arg1 *QRandomGenerator64) *QRandomGenerator64 {
	__r, _, _ := purego.SyscallN(addr_QRandomGenerator64_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQRandomGenerator64(__r)
}

func (this *QRandomGenerator64) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QRandomGenerator64_delete(this.h)
	}
	this.UnsafeRelease()
}
