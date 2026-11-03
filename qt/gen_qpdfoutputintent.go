package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QPdfOutputIntent_new uintptr
	addr_QPdfOutputIntent_new2 uintptr
	addr_QPdfOutputIntent_OperatorAssign uintptr
	addr_QPdfOutputIntent_swap uintptr
	addr_QPdfOutputIntent_outputConditionIdentifier uintptr
	addr_QPdfOutputIntent_setOutputConditionIdentifier uintptr
	addr_QPdfOutputIntent_outputCondition uintptr
	addr_QPdfOutputIntent_setOutputCondition uintptr
	addr_QPdfOutputIntent_registryName uintptr
	addr_QPdfOutputIntent_setRegistryName uintptr
	addr_QPdfOutputIntent_outputProfile uintptr
	addr_QPdfOutputIntent_setOutputProfile uintptr
	pg_QPdfOutputIntent_delete func(uintptr)
)

func init() {
	symaddr(&addr_QPdfOutputIntent_new, "QPdfOutputIntent_new")
	symaddr(&addr_QPdfOutputIntent_new2, "QPdfOutputIntent_new2")
	symaddr(&addr_QPdfOutputIntent_OperatorAssign, "QPdfOutputIntent_OperatorAssign")
	symaddr(&addr_QPdfOutputIntent_swap, "QPdfOutputIntent_swap")
	symaddr(&addr_QPdfOutputIntent_outputConditionIdentifier, "QPdfOutputIntent_outputConditionIdentifier")
	symaddr(&addr_QPdfOutputIntent_setOutputConditionIdentifier, "QPdfOutputIntent_setOutputConditionIdentifier")
	symaddr(&addr_QPdfOutputIntent_outputCondition, "QPdfOutputIntent_outputCondition")
	symaddr(&addr_QPdfOutputIntent_setOutputCondition, "QPdfOutputIntent_setOutputCondition")
	symaddr(&addr_QPdfOutputIntent_registryName, "QPdfOutputIntent_registryName")
	symaddr(&addr_QPdfOutputIntent_setRegistryName, "QPdfOutputIntent_setRegistryName")
	symaddr(&addr_QPdfOutputIntent_outputProfile, "QPdfOutputIntent_outputProfile")
	symaddr(&addr_QPdfOutputIntent_setOutputProfile, "QPdfOutputIntent_setOutputProfile")
	reg(&pg_QPdfOutputIntent_delete, "QPdfOutputIntent_delete")
}

type QPdfOutputIntent struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QPdfOutputIntent) UnsafeGoID() int64     { return this.vid }
func (this *QPdfOutputIntent) UnsafeSetGoID(id int64) { this.vid = id }

func bindQPdfOutputIntent(h uintptr) *QPdfOutputIntent {
	if h == 0 {
		return nil
	}
	return &QPdfOutputIntent{h: h}
}

func newownedQPdfOutputIntent(h uintptr) *QPdfOutputIntent {
	o := bindQPdfOutputIntent(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QPdfOutputIntent).Delete)
	}
	return o
}

func (this *QPdfOutputIntent) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QPdfOutputIntent) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQPdfOutputIntent(h unsafe.Pointer) *QPdfOutputIntent { return bindQPdfOutputIntent(uintptr(h)) }
func UnsafeNewQPdfOutputIntentOwned(h unsafe.Pointer) *QPdfOutputIntent { return newownedQPdfOutputIntent(uintptr(h)) }

func (this *QPdfOutputIntent) UnsafeRelease() { this.h = 0 }

func NewQPdfOutputIntent() *QPdfOutputIntent {
	__r, _, _ := purego.SyscallN(addr_QPdfOutputIntent_new)
	__o := newownedQPdfOutputIntent(__r)
	return __o
}

func NewQPdfOutputIntent2(other *QPdfOutputIntent) *QPdfOutputIntent {
	__r, _, _ := purego.SyscallN(addr_QPdfOutputIntent_new2, uintptr(other.UnsafePointer()))
	__o := newownedQPdfOutputIntent(__r)
	runtime.KeepAlive(other)
	return __o
}

func (this *QPdfOutputIntent) OperatorAssign(other *QPdfOutputIntent) *QPdfOutputIntent {
	__r, _, _ := purego.SyscallN(addr_QPdfOutputIntent_OperatorAssign, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return bindQPdfOutputIntent(__r)
}

func (this *QPdfOutputIntent) Swap(other *QPdfOutputIntent) {
	purego.SyscallN(addr_QPdfOutputIntent_swap, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
}

func (this *QPdfOutputIntent) OutputConditionIdentifier() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QPdfOutputIntent_outputConditionIdentifier, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QPdfOutputIntent) SetOutputConditionIdentifier(identifier string) {
	var identifier_p uintptr
	if len(identifier) > 0 {
		identifier_p = uintptr(unsafe.Pointer(unsafe.StringData(identifier)))
	}
	purego.SyscallN(addr_QPdfOutputIntent_setOutputConditionIdentifier, this.h, identifier_p, uintptr(len(identifier)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(identifier)
}

func (this *QPdfOutputIntent) OutputCondition() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QPdfOutputIntent_outputCondition, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QPdfOutputIntent) SetOutputCondition(condition string) {
	var condition_p uintptr
	if len(condition) > 0 {
		condition_p = uintptr(unsafe.Pointer(unsafe.StringData(condition)))
	}
	purego.SyscallN(addr_QPdfOutputIntent_setOutputCondition, this.h, condition_p, uintptr(len(condition)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(condition)
}

func (this *QPdfOutputIntent) RegistryName() *QUrl {
	__r, _, _ := purego.SyscallN(addr_QPdfOutputIntent_registryName, this.h)
	runtime.KeepAlive(this)
	return newownedQUrl(__r)
}

func (this *QPdfOutputIntent) SetRegistryName(name *QUrl) {
	purego.SyscallN(addr_QPdfOutputIntent_setRegistryName, this.h, uintptr(name.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(name)
}

func (this *QPdfOutputIntent) OutputProfile() *QColorSpace {
	__r, _, _ := purego.SyscallN(addr_QPdfOutputIntent_outputProfile, this.h)
	runtime.KeepAlive(this)
	return newownedQColorSpace(__r)
}

func (this *QPdfOutputIntent) SetOutputProfile(profile *QColorSpace) {
	purego.SyscallN(addr_QPdfOutputIntent_setOutputProfile, this.h, uintptr(profile.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(profile)
}

func (this *QPdfOutputIntent) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QPdfOutputIntent_delete(this.h)
	}
	this.UnsafeRelease()
}
