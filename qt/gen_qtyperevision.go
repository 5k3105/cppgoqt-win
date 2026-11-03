package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QTypeRevision_new uintptr
	addr_QTypeRevision_new2 uintptr
	addr_QTypeRevision_zero uintptr
	addr_QTypeRevision_hasMajorVersion uintptr
	addr_QTypeRevision_majorVersion uintptr
	addr_QTypeRevision_hasMinorVersion uintptr
	addr_QTypeRevision_minorVersion uintptr
	addr_QTypeRevision_isValid uintptr
	addr_QTypeRevision_OperatorEqual uintptr
	addr_QTypeRevision_OperatorNotEqual uintptr
	addr_QTypeRevision_OperatorLesser uintptr
	addr_QTypeRevision_OperatorGreater uintptr
	addr_QTypeRevision_OperatorLesserOrEqual uintptr
	addr_QTypeRevision_OperatorGreaterOrEqual uintptr
	pg_QTypeRevision_delete func(uintptr)
)

func init() {
	symaddr(&addr_QTypeRevision_new, "QTypeRevision_new")
	symaddr(&addr_QTypeRevision_new2, "QTypeRevision_new2")
	symaddr(&addr_QTypeRevision_zero, "QTypeRevision_zero")
	symaddr(&addr_QTypeRevision_hasMajorVersion, "QTypeRevision_hasMajorVersion")
	symaddr(&addr_QTypeRevision_majorVersion, "QTypeRevision_majorVersion")
	symaddr(&addr_QTypeRevision_hasMinorVersion, "QTypeRevision_hasMinorVersion")
	symaddr(&addr_QTypeRevision_minorVersion, "QTypeRevision_minorVersion")
	symaddr(&addr_QTypeRevision_isValid, "QTypeRevision_isValid")
	symaddr(&addr_QTypeRevision_OperatorEqual, "QTypeRevision_OperatorEqual")
	symaddr(&addr_QTypeRevision_OperatorNotEqual, "QTypeRevision_OperatorNotEqual")
	symaddr(&addr_QTypeRevision_OperatorLesser, "QTypeRevision_OperatorLesser")
	symaddr(&addr_QTypeRevision_OperatorGreater, "QTypeRevision_OperatorGreater")
	symaddr(&addr_QTypeRevision_OperatorLesserOrEqual, "QTypeRevision_OperatorLesserOrEqual")
	symaddr(&addr_QTypeRevision_OperatorGreaterOrEqual, "QTypeRevision_OperatorGreaterOrEqual")
	reg(&pg_QTypeRevision_delete, "QTypeRevision_delete")
}

type QTypeRevision struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QTypeRevision) UnsafeGoID() int64     { return this.vid }
func (this *QTypeRevision) UnsafeSetGoID(id int64) { this.vid = id }

func bindQTypeRevision(h uintptr) *QTypeRevision {
	if h == 0 {
		return nil
	}
	return &QTypeRevision{h: h}
}

func newownedQTypeRevision(h uintptr) *QTypeRevision {
	o := bindQTypeRevision(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QTypeRevision).Delete)
	}
	return o
}

func (this *QTypeRevision) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QTypeRevision) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQTypeRevision(h unsafe.Pointer) *QTypeRevision { return bindQTypeRevision(uintptr(h)) }
func UnsafeNewQTypeRevisionOwned(h unsafe.Pointer) *QTypeRevision { return newownedQTypeRevision(uintptr(h)) }

func (this *QTypeRevision) UnsafeRelease() { this.h = 0 }

func NewQTypeRevision() *QTypeRevision {
	__r, _, _ := purego.SyscallN(addr_QTypeRevision_new)
	__o := newownedQTypeRevision(__r)
	return __o
}

func NewQTypeRevision2(arg1 *QTypeRevision) *QTypeRevision {
	__r, _, _ := purego.SyscallN(addr_QTypeRevision_new2, uintptr(arg1.UnsafePointer()))
	__o := newownedQTypeRevision(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func QTypeRevision_Zero() *QTypeRevision {
	__r, _, _ := purego.SyscallN(addr_QTypeRevision_zero)
	return newownedQTypeRevision(__r)
}

func (this *QTypeRevision) HasMajorVersion() bool {
	__r0, _, _ := purego.SyscallN(addr_QTypeRevision_hasMajorVersion, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTypeRevision) MajorVersion() byte {
	__r0, _, _ := purego.SyscallN(addr_QTypeRevision_majorVersion, this.h)
	__r := uint8(__r0)
	runtime.KeepAlive(this)
	return byte(__r)
}

func (this *QTypeRevision) HasMinorVersion() bool {
	__r0, _, _ := purego.SyscallN(addr_QTypeRevision_hasMinorVersion, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTypeRevision) MinorVersion() byte {
	__r0, _, _ := purego.SyscallN(addr_QTypeRevision_minorVersion, this.h)
	__r := uint8(__r0)
	runtime.KeepAlive(this)
	return byte(__r)
}

func (this *QTypeRevision) IsValid() bool {
	__r0, _, _ := purego.SyscallN(addr_QTypeRevision_isValid, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTypeRevision) OperatorEqual(rhs *QTypeRevision) bool {
	__r0, _, _ := purego.SyscallN(addr_QTypeRevision_OperatorEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QTypeRevision) OperatorNotEqual(rhs *QTypeRevision) bool {
	__r0, _, _ := purego.SyscallN(addr_QTypeRevision_OperatorNotEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QTypeRevision) OperatorLesser(rhs *QTypeRevision) bool {
	__r0, _, _ := purego.SyscallN(addr_QTypeRevision_OperatorLesser, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QTypeRevision) OperatorGreater(rhs *QTypeRevision) bool {
	__r0, _, _ := purego.SyscallN(addr_QTypeRevision_OperatorGreater, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QTypeRevision) OperatorLesserOrEqual(rhs *QTypeRevision) bool {
	__r0, _, _ := purego.SyscallN(addr_QTypeRevision_OperatorLesserOrEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QTypeRevision) OperatorGreaterOrEqual(rhs *QTypeRevision) bool {
	__r0, _, _ := purego.SyscallN(addr_QTypeRevision_OperatorGreaterOrEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QTypeRevision) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QTypeRevision_delete(this.h)
	}
	this.UnsafeRelease()
}
