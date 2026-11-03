package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QFont__Tag_new uintptr
	addr_QFont__Tag_new2 uintptr
	addr_QFont__Tag_isValid uintptr
	addr_QFont__Tag_value uintptr
	addr_QFont__Tag_toString uintptr
	addr_QFont__Tag_fromValue uintptr
	addr_QFont__Tag_fromString uintptr
	addr_QFont__Tag_OperatorAssign uintptr
	addr_QFont__Tag_OperatorEqual uintptr
	addr_QFont__Tag_OperatorNotEqual uintptr
	addr_QFont__Tag_OperatorLesser uintptr
	addr_QFont__Tag_OperatorGreater uintptr
	addr_QFont__Tag_OperatorLesserOrEqual uintptr
	addr_QFont__Tag_OperatorGreaterOrEqual uintptr
	pg_QFont__Tag_delete func(uintptr)
)

func init() {
	symaddr(&addr_QFont__Tag_new, "QFont__Tag_new")
	symaddr(&addr_QFont__Tag_new2, "QFont__Tag_new2")
	symaddr(&addr_QFont__Tag_isValid, "QFont__Tag_isValid")
	symaddr(&addr_QFont__Tag_value, "QFont__Tag_value")
	symaddr(&addr_QFont__Tag_toString, "QFont__Tag_toString")
	symaddr(&addr_QFont__Tag_fromValue, "QFont__Tag_fromValue")
	symaddr(&addr_QFont__Tag_fromString, "QFont__Tag_fromString")
	symaddr(&addr_QFont__Tag_OperatorAssign, "QFont__Tag_OperatorAssign")
	symaddr(&addr_QFont__Tag_OperatorEqual, "QFont__Tag_OperatorEqual")
	symaddr(&addr_QFont__Tag_OperatorNotEqual, "QFont__Tag_OperatorNotEqual")
	symaddr(&addr_QFont__Tag_OperatorLesser, "QFont__Tag_OperatorLesser")
	symaddr(&addr_QFont__Tag_OperatorGreater, "QFont__Tag_OperatorGreater")
	symaddr(&addr_QFont__Tag_OperatorLesserOrEqual, "QFont__Tag_OperatorLesserOrEqual")
	symaddr(&addr_QFont__Tag_OperatorGreaterOrEqual, "QFont__Tag_OperatorGreaterOrEqual")
	reg(&pg_QFont__Tag_delete, "QFont__Tag_delete")
}

type QFont__Tag struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QFont__Tag) UnsafeGoID() int64     { return this.vid }
func (this *QFont__Tag) UnsafeSetGoID(id int64) { this.vid = id }

func bindQFont__Tag(h uintptr) *QFont__Tag {
	if h == 0 {
		return nil
	}
	return &QFont__Tag{h: h}
}

func newownedQFont__Tag(h uintptr) *QFont__Tag {
	o := bindQFont__Tag(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QFont__Tag).Delete)
	}
	return o
}

func (this *QFont__Tag) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QFont__Tag) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQFont__Tag(h unsafe.Pointer) *QFont__Tag { return bindQFont__Tag(uintptr(h)) }
func UnsafeNewQFont__TagOwned(h unsafe.Pointer) *QFont__Tag { return newownedQFont__Tag(uintptr(h)) }

func (this *QFont__Tag) UnsafeRelease() { this.h = 0 }

func NewQFont__Tag() *QFont__Tag {
	__r, _, _ := purego.SyscallN(addr_QFont__Tag_new)
	__o := newownedQFont__Tag(__r)
	return __o
}

func NewQFont__Tag2(arg1 *QFont__Tag) *QFont__Tag {
	__r, _, _ := purego.SyscallN(addr_QFont__Tag_new2, uintptr(arg1.UnsafePointer()))
	__o := newownedQFont__Tag(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QFont__Tag) IsValid() bool {
	__r0, _, _ := purego.SyscallN(addr_QFont__Tag_isValid, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFont__Tag) Value() uint {
	__r0, _, _ := purego.SyscallN(addr_QFont__Tag_value, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint(__r)
}

func (this *QFont__Tag) ToString() *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QFont__Tag_toString, this.h)
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func QFont__Tag_FromValue(value uint) (*QFont__Tag, bool) {
	var __b int32
	__r, _, _ := purego.SyscallN(addr_QFont__Tag_fromValue, uintptr(value), uintptr(unsafe.Pointer(&__b)))
	return newownedQFont__Tag(__r), __b != 0
}

func QFont__Tag_FromString(view string) (*QFont__Tag, bool) {
	var view_p uintptr
	if len(view) > 0 {
		view_p = uintptr(unsafe.Pointer(unsafe.StringData(view)))
	}
	var __b int32
	__r, _, _ := purego.SyscallN(addr_QFont__Tag_fromString, view_p, uintptr(len(view)), uintptr(unsafe.Pointer(&__b)))
	runtime.KeepAlive(view)
	return newownedQFont__Tag(__r), __b != 0
}

func (this *QFont__Tag) OperatorAssign(arg1 *QFont__Tag) *QFont__Tag {
	__r, _, _ := purego.SyscallN(addr_QFont__Tag_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQFont__Tag(__r)
}

func (this *QFont__Tag) OperatorEqual(rhs *QFont__Tag) bool {
	__r0, _, _ := purego.SyscallN(addr_QFont__Tag_OperatorEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QFont__Tag) OperatorNotEqual(rhs *QFont__Tag) bool {
	__r0, _, _ := purego.SyscallN(addr_QFont__Tag_OperatorNotEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QFont__Tag) OperatorLesser(rhs *QFont__Tag) bool {
	__r0, _, _ := purego.SyscallN(addr_QFont__Tag_OperatorLesser, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QFont__Tag) OperatorGreater(rhs *QFont__Tag) bool {
	__r0, _, _ := purego.SyscallN(addr_QFont__Tag_OperatorGreater, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QFont__Tag) OperatorLesserOrEqual(rhs *QFont__Tag) bool {
	__r0, _, _ := purego.SyscallN(addr_QFont__Tag_OperatorLesserOrEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QFont__Tag) OperatorGreaterOrEqual(rhs *QFont__Tag) bool {
	__r0, _, _ := purego.SyscallN(addr_QFont__Tag_OperatorGreaterOrEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QFont__Tag) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QFont__Tag_delete(this.h)
	}
	this.UnsafeRelease()
}
