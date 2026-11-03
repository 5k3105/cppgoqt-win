package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QOperatingSystemVersion_chain uintptr
	addr_QOperatingSystemVersion_new uintptr
	addr_QOperatingSystemVersion_new2 uintptr
	addr_QOperatingSystemVersion_new3 uintptr
	addr_QOperatingSystemVersion_new4 uintptr
	addr_QOperatingSystemVersion_new5 uintptr
	addr_QOperatingSystemVersion_current uintptr
	addr_QOperatingSystemVersion_currentType uintptr
	addr_QOperatingSystemVersion_type uintptr
	addr_QOperatingSystemVersion_name uintptr
	pg_QOperatingSystemVersion_delete func(uintptr)
)

func init() {
	symaddr(&addr_QOperatingSystemVersion_chain, "QOperatingSystemVersion_chain")
	symaddr(&addr_QOperatingSystemVersion_new, "QOperatingSystemVersion_new")
	symaddr(&addr_QOperatingSystemVersion_new2, "QOperatingSystemVersion_new2")
	symaddr(&addr_QOperatingSystemVersion_new3, "QOperatingSystemVersion_new3")
	symaddr(&addr_QOperatingSystemVersion_new4, "QOperatingSystemVersion_new4")
	symaddr(&addr_QOperatingSystemVersion_new5, "QOperatingSystemVersion_new5")
	symaddr(&addr_QOperatingSystemVersion_current, "QOperatingSystemVersion_current")
	symaddr(&addr_QOperatingSystemVersion_currentType, "QOperatingSystemVersion_currentType")
	symaddr(&addr_QOperatingSystemVersion_type, "QOperatingSystemVersion_type")
	symaddr(&addr_QOperatingSystemVersion_name, "QOperatingSystemVersion_name")
	reg(&pg_QOperatingSystemVersion_delete, "QOperatingSystemVersion_delete")
}

type QOperatingSystemVersion struct {
	h uintptr
	*QOperatingSystemVersionUnexported
}

func bindQOperatingSystemVersion(h uintptr) *QOperatingSystemVersion {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QOperatingSystemVersion_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QOperatingSystemVersion{h: h, QOperatingSystemVersionUnexported: &QOperatingSystemVersionUnexported{h: __c[0], QOperatingSystemVersionBase: &QOperatingSystemVersionBase{h: __c[1]}}}
}

func newownedQOperatingSystemVersion(h uintptr) *QOperatingSystemVersion {
	o := bindQOperatingSystemVersion(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QOperatingSystemVersion).Delete)
	}
	return o
}

func (this *QOperatingSystemVersion) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QOperatingSystemVersion) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQOperatingSystemVersion(h unsafe.Pointer) *QOperatingSystemVersion { return bindQOperatingSystemVersion(uintptr(h)) }
func UnsafeNewQOperatingSystemVersionOwned(h unsafe.Pointer) *QOperatingSystemVersion { return newownedQOperatingSystemVersion(uintptr(h)) }

func (this *QOperatingSystemVersion) UnsafeRelease() {
	this.h = 0
	if this.QOperatingSystemVersionUnexported != nil {
		this.QOperatingSystemVersionUnexported.UnsafeRelease()
	}
}

func NewQOperatingSystemVersion(osversion *QOperatingSystemVersionBase) *QOperatingSystemVersion {
	__r, _, _ := purego.SyscallN(addr_QOperatingSystemVersion_new, uintptr(osversion.UnsafePointer()))
	__o := newownedQOperatingSystemVersion(__r)
	runtime.KeepAlive(osversion)
	return __o
}

func NewQOperatingSystemVersion2(osType QOperatingSystemVersion__OSType, vmajor int) *QOperatingSystemVersion {
	__r, _, _ := purego.SyscallN(addr_QOperatingSystemVersion_new2, uintptr(osType), uintptr(vmajor))
	__o := newownedQOperatingSystemVersion(__r)
	return __o
}

func NewQOperatingSystemVersion3(arg1 *QOperatingSystemVersion) *QOperatingSystemVersion {
	__r, _, _ := purego.SyscallN(addr_QOperatingSystemVersion_new3, uintptr(arg1.UnsafePointer()))
	__o := newownedQOperatingSystemVersion(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func NewQOperatingSystemVersion4(osType QOperatingSystemVersion__OSType, vmajor int, vminor int) *QOperatingSystemVersion {
	__r, _, _ := purego.SyscallN(addr_QOperatingSystemVersion_new4, uintptr(osType), uintptr(vmajor), uintptr(vminor))
	__o := newownedQOperatingSystemVersion(__r)
	return __o
}

func NewQOperatingSystemVersion5(osType QOperatingSystemVersion__OSType, vmajor int, vminor int, vmicro int) *QOperatingSystemVersion {
	__r, _, _ := purego.SyscallN(addr_QOperatingSystemVersion_new5, uintptr(osType), uintptr(vmajor), uintptr(vminor), uintptr(vmicro))
	__o := newownedQOperatingSystemVersion(__r)
	return __o
}

func QOperatingSystemVersion_Current() *QOperatingSystemVersion {
	__r, _, _ := purego.SyscallN(addr_QOperatingSystemVersion_current)
	return newownedQOperatingSystemVersion(__r)
}

func QOperatingSystemVersion_CurrentType() QOperatingSystemVersion__OSType {
	__r0, _, _ := purego.SyscallN(addr_QOperatingSystemVersion_currentType)
	__r := int32(uint32(__r0))
	return QOperatingSystemVersion__OSType(__r)
}

func (this *QOperatingSystemVersion) Type() QOperatingSystemVersion__OSType {
	__r0, _, _ := purego.SyscallN(addr_QOperatingSystemVersion_type, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QOperatingSystemVersion__OSType(__r)
}

func QOperatingSystemVersion_Name(osversion *QOperatingSystemVersionBase) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QOperatingSystemVersion_name, uintptr(osversion.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(osversion)
	return pgStr(__r, __n)
}

func (this *QOperatingSystemVersion) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QOperatingSystemVersion_delete(this.h)
	}
	this.UnsafeRelease()
}
