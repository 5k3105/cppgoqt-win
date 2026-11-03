package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QOperatingSystemVersionBase_new uintptr
	addr_QOperatingSystemVersionBase_new2 uintptr
	addr_QOperatingSystemVersionBase_new3 uintptr
	addr_QOperatingSystemVersionBase_new4 uintptr
	addr_QOperatingSystemVersionBase_current uintptr
	addr_QOperatingSystemVersionBase_name uintptr
	addr_QOperatingSystemVersionBase_currentType uintptr
	addr_QOperatingSystemVersionBase_version uintptr
	addr_QOperatingSystemVersionBase_majorVersion uintptr
	addr_QOperatingSystemVersionBase_minorVersion uintptr
	addr_QOperatingSystemVersionBase_microVersion uintptr
	addr_QOperatingSystemVersionBase_segmentCount uintptr
	addr_QOperatingSystemVersionBase_type uintptr
	addr_QOperatingSystemVersionBase_name2 uintptr
	addr_QOperatingSystemVersionBase_OperatorGreater uintptr
	addr_QOperatingSystemVersionBase_OperatorGreaterOrEqual uintptr
	addr_QOperatingSystemVersionBase_OperatorLesser uintptr
	addr_QOperatingSystemVersionBase_OperatorLesserOrEqual uintptr
	pg_QOperatingSystemVersionBase_delete func(uintptr)
)

func init() {
	symaddr(&addr_QOperatingSystemVersionBase_new, "QOperatingSystemVersionBase_new")
	symaddr(&addr_QOperatingSystemVersionBase_new2, "QOperatingSystemVersionBase_new2")
	symaddr(&addr_QOperatingSystemVersionBase_new3, "QOperatingSystemVersionBase_new3")
	symaddr(&addr_QOperatingSystemVersionBase_new4, "QOperatingSystemVersionBase_new4")
	symaddr(&addr_QOperatingSystemVersionBase_current, "QOperatingSystemVersionBase_current")
	symaddr(&addr_QOperatingSystemVersionBase_name, "QOperatingSystemVersionBase_name")
	symaddr(&addr_QOperatingSystemVersionBase_currentType, "QOperatingSystemVersionBase_currentType")
	symaddr(&addr_QOperatingSystemVersionBase_version, "QOperatingSystemVersionBase_version")
	symaddr(&addr_QOperatingSystemVersionBase_majorVersion, "QOperatingSystemVersionBase_majorVersion")
	symaddr(&addr_QOperatingSystemVersionBase_minorVersion, "QOperatingSystemVersionBase_minorVersion")
	symaddr(&addr_QOperatingSystemVersionBase_microVersion, "QOperatingSystemVersionBase_microVersion")
	symaddr(&addr_QOperatingSystemVersionBase_segmentCount, "QOperatingSystemVersionBase_segmentCount")
	symaddr(&addr_QOperatingSystemVersionBase_type, "QOperatingSystemVersionBase_type")
	symaddr(&addr_QOperatingSystemVersionBase_name2, "QOperatingSystemVersionBase_name2")
	symaddr(&addr_QOperatingSystemVersionBase_OperatorGreater, "QOperatingSystemVersionBase_OperatorGreater")
	symaddr(&addr_QOperatingSystemVersionBase_OperatorGreaterOrEqual, "QOperatingSystemVersionBase_OperatorGreaterOrEqual")
	symaddr(&addr_QOperatingSystemVersionBase_OperatorLesser, "QOperatingSystemVersionBase_OperatorLesser")
	symaddr(&addr_QOperatingSystemVersionBase_OperatorLesserOrEqual, "QOperatingSystemVersionBase_OperatorLesserOrEqual")
	reg(&pg_QOperatingSystemVersionBase_delete, "QOperatingSystemVersionBase_delete")
}

type QOperatingSystemVersionBase struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QOperatingSystemVersionBase) UnsafeGoID() int64     { return this.vid }
func (this *QOperatingSystemVersionBase) UnsafeSetGoID(id int64) { this.vid = id }

func bindQOperatingSystemVersionBase(h uintptr) *QOperatingSystemVersionBase {
	if h == 0 {
		return nil
	}
	return &QOperatingSystemVersionBase{h: h}
}

func newownedQOperatingSystemVersionBase(h uintptr) *QOperatingSystemVersionBase {
	o := bindQOperatingSystemVersionBase(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QOperatingSystemVersionBase).Delete)
	}
	return o
}

func (this *QOperatingSystemVersionBase) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QOperatingSystemVersionBase) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQOperatingSystemVersionBase(h unsafe.Pointer) *QOperatingSystemVersionBase { return bindQOperatingSystemVersionBase(uintptr(h)) }
func UnsafeNewQOperatingSystemVersionBaseOwned(h unsafe.Pointer) *QOperatingSystemVersionBase { return newownedQOperatingSystemVersionBase(uintptr(h)) }

func (this *QOperatingSystemVersionBase) UnsafeRelease() { this.h = 0 }

func NewQOperatingSystemVersionBase(osType QOperatingSystemVersionBase__OSType, vmajor int) *QOperatingSystemVersionBase {
	__r, _, _ := purego.SyscallN(addr_QOperatingSystemVersionBase_new, uintptr(osType), uintptr(vmajor))
	__o := newownedQOperatingSystemVersionBase(__r)
	return __o
}

func NewQOperatingSystemVersionBase2(arg1 *QOperatingSystemVersionBase) *QOperatingSystemVersionBase {
	__r, _, _ := purego.SyscallN(addr_QOperatingSystemVersionBase_new2, uintptr(arg1.UnsafePointer()))
	__o := newownedQOperatingSystemVersionBase(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func NewQOperatingSystemVersionBase3(osType QOperatingSystemVersionBase__OSType, vmajor int, vminor int) *QOperatingSystemVersionBase {
	__r, _, _ := purego.SyscallN(addr_QOperatingSystemVersionBase_new3, uintptr(osType), uintptr(vmajor), uintptr(vminor))
	__o := newownedQOperatingSystemVersionBase(__r)
	return __o
}

func NewQOperatingSystemVersionBase4(osType QOperatingSystemVersionBase__OSType, vmajor int, vminor int, vmicro int) *QOperatingSystemVersionBase {
	__r, _, _ := purego.SyscallN(addr_QOperatingSystemVersionBase_new4, uintptr(osType), uintptr(vmajor), uintptr(vminor), uintptr(vmicro))
	__o := newownedQOperatingSystemVersionBase(__r)
	return __o
}

func QOperatingSystemVersionBase_Current() *QOperatingSystemVersionBase {
	__r, _, _ := purego.SyscallN(addr_QOperatingSystemVersionBase_current)
	return newownedQOperatingSystemVersionBase(__r)
}

func QOperatingSystemVersionBase_Name(osversion *QOperatingSystemVersionBase) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QOperatingSystemVersionBase_name, uintptr(osversion.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(osversion)
	return pgStr(__r, __n)
}

func QOperatingSystemVersionBase_CurrentType() QOperatingSystemVersionBase__OSType {
	__r0, _, _ := purego.SyscallN(addr_QOperatingSystemVersionBase_currentType)
	__r := int32(uint32(__r0))
	return QOperatingSystemVersionBase__OSType(__r)
}

func (this *QOperatingSystemVersionBase) Version() *QVersionNumber {
	__r, _, _ := purego.SyscallN(addr_QOperatingSystemVersionBase_version, this.h)
	runtime.KeepAlive(this)
	return newownedQVersionNumber(__r)
}

func (this *QOperatingSystemVersionBase) MajorVersion() int {
	__r0, _, _ := purego.SyscallN(addr_QOperatingSystemVersionBase_majorVersion, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QOperatingSystemVersionBase) MinorVersion() int {
	__r0, _, _ := purego.SyscallN(addr_QOperatingSystemVersionBase_minorVersion, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QOperatingSystemVersionBase) MicroVersion() int {
	__r0, _, _ := purego.SyscallN(addr_QOperatingSystemVersionBase_microVersion, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QOperatingSystemVersionBase) SegmentCount() int {
	__r0, _, _ := purego.SyscallN(addr_QOperatingSystemVersionBase_segmentCount, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QOperatingSystemVersionBase) Type() QOperatingSystemVersionBase__OSType {
	__r0, _, _ := purego.SyscallN(addr_QOperatingSystemVersionBase_type, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QOperatingSystemVersionBase__OSType(__r)
}

func (this *QOperatingSystemVersionBase) Name() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QOperatingSystemVersionBase_name2, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QOperatingSystemVersionBase) OperatorGreater(rhs *QOperatingSystemVersionBase) bool {
	__r0, _, _ := purego.SyscallN(addr_QOperatingSystemVersionBase_OperatorGreater, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QOperatingSystemVersionBase) OperatorGreaterOrEqual(rhs *QOperatingSystemVersionBase) bool {
	__r0, _, _ := purego.SyscallN(addr_QOperatingSystemVersionBase_OperatorGreaterOrEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QOperatingSystemVersionBase) OperatorLesser(rhs *QOperatingSystemVersionBase) bool {
	__r0, _, _ := purego.SyscallN(addr_QOperatingSystemVersionBase_OperatorLesser, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QOperatingSystemVersionBase) OperatorLesserOrEqual(rhs *QOperatingSystemVersionBase) bool {
	__r0, _, _ := purego.SyscallN(addr_QOperatingSystemVersionBase_OperatorLesserOrEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QOperatingSystemVersionBase) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QOperatingSystemVersionBase_delete(this.h)
	}
	this.UnsafeRelease()
}
