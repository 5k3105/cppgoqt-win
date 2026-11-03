package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QOperatingSystemVersionUnexported_base func(uintptr) uintptr
	addr_QOperatingSystemVersionUnexported_new uintptr
	addr_QOperatingSystemVersionUnexported_new2 uintptr
	addr_QOperatingSystemVersionUnexported_new3 uintptr
	addr_QOperatingSystemVersionUnexported_new4 uintptr
	addr_QOperatingSystemVersionUnexported_current uintptr
	addr_QOperatingSystemVersionUnexported_name uintptr
	addr_QOperatingSystemVersionUnexported_currentType uintptr
	pg_QOperatingSystemVersionUnexported_delete func(uintptr)
)

func init() {
	reg(&pg_QOperatingSystemVersionUnexported_base, "QOperatingSystemVersionUnexported_base")
	symaddr(&addr_QOperatingSystemVersionUnexported_new, "QOperatingSystemVersionUnexported_new")
	symaddr(&addr_QOperatingSystemVersionUnexported_new2, "QOperatingSystemVersionUnexported_new2")
	symaddr(&addr_QOperatingSystemVersionUnexported_new3, "QOperatingSystemVersionUnexported_new3")
	symaddr(&addr_QOperatingSystemVersionUnexported_new4, "QOperatingSystemVersionUnexported_new4")
	symaddr(&addr_QOperatingSystemVersionUnexported_current, "QOperatingSystemVersionUnexported_current")
	symaddr(&addr_QOperatingSystemVersionUnexported_name, "QOperatingSystemVersionUnexported_name")
	symaddr(&addr_QOperatingSystemVersionUnexported_currentType, "QOperatingSystemVersionUnexported_currentType")
	reg(&pg_QOperatingSystemVersionUnexported_delete, "QOperatingSystemVersionUnexported_delete")
}

type QOperatingSystemVersionUnexported struct {
	h uintptr
	*QOperatingSystemVersionBase
}

func bindQOperatingSystemVersionUnexported(h uintptr) *QOperatingSystemVersionUnexported {
	if h == 0 {
		return nil
	}
	return &QOperatingSystemVersionUnexported{h: h, QOperatingSystemVersionBase: bindQOperatingSystemVersionBase(pg_QOperatingSystemVersionUnexported_base(h))}
}

func newownedQOperatingSystemVersionUnexported(h uintptr) *QOperatingSystemVersionUnexported {
	o := bindQOperatingSystemVersionUnexported(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QOperatingSystemVersionUnexported).Delete)
	}
	return o
}

func (this *QOperatingSystemVersionUnexported) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QOperatingSystemVersionUnexported) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQOperatingSystemVersionUnexported(h unsafe.Pointer) *QOperatingSystemVersionUnexported { return bindQOperatingSystemVersionUnexported(uintptr(h)) }
func UnsafeNewQOperatingSystemVersionUnexportedOwned(h unsafe.Pointer) *QOperatingSystemVersionUnexported { return newownedQOperatingSystemVersionUnexported(uintptr(h)) }

func (this *QOperatingSystemVersionUnexported) UnsafeRelease() {
	this.h = 0
	if this.QOperatingSystemVersionBase != nil {
		this.QOperatingSystemVersionBase.UnsafeRelease()
	}
}

func NewQOperatingSystemVersionUnexported(other *QOperatingSystemVersionBase) *QOperatingSystemVersionUnexported {
	__r, _, _ := purego.SyscallN(addr_QOperatingSystemVersionUnexported_new, uintptr(other.UnsafePointer()))
	__o := newownedQOperatingSystemVersionUnexported(__r)
	runtime.KeepAlive(other)
	return __o
}

func NewQOperatingSystemVersionUnexported2() *QOperatingSystemVersionUnexported {
	__r, _, _ := purego.SyscallN(addr_QOperatingSystemVersionUnexported_new2)
	__o := newownedQOperatingSystemVersionUnexported(__r)
	return __o
}

func NewQOperatingSystemVersionUnexported3(arg1 *QOperatingSystemVersionUnexported) *QOperatingSystemVersionUnexported {
	__r, _, _ := purego.SyscallN(addr_QOperatingSystemVersionUnexported_new3, uintptr(arg1.UnsafePointer()))
	__o := newownedQOperatingSystemVersionUnexported(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func NewQOperatingSystemVersionUnexported4(arg1 QOperatingSystemVersionBase__OSType, arg2 int, arg3 int, arg4 int) *QOperatingSystemVersionUnexported {
	__r, _, _ := purego.SyscallN(addr_QOperatingSystemVersionUnexported_new4, uintptr(arg1), uintptr(arg2), uintptr(arg3), uintptr(arg4))
	__o := newownedQOperatingSystemVersionUnexported(__r)
	return __o
}

func QOperatingSystemVersionUnexported_Current() *QOperatingSystemVersionBase {
	__r, _, _ := purego.SyscallN(addr_QOperatingSystemVersionUnexported_current)
	return newownedQOperatingSystemVersionBase(__r)
}

func QOperatingSystemVersionUnexported_Name(osversion *QOperatingSystemVersionBase) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QOperatingSystemVersionUnexported_name, uintptr(osversion.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(osversion)
	return pgStr(__r, __n)
}

func QOperatingSystemVersionUnexported_CurrentType() QOperatingSystemVersionBase__OSType {
	__r0, _, _ := purego.SyscallN(addr_QOperatingSystemVersionUnexported_currentType)
	__r := int32(uint32(__r0))
	return QOperatingSystemVersionBase__OSType(__r)
}

func (this *QOperatingSystemVersionUnexported) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QOperatingSystemVersionUnexported_delete(this.h)
	}
	this.UnsafeRelease()
}
