package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QTextLength_new uintptr
	pg_QTextLength_new2 func(int32, float64) uintptr
	addr_QTextLength_new3 uintptr
	addr_QTextLength_type uintptr
	pg_QTextLength_value func(uintptr, float64) float64
	pg_QTextLength_rawValue func(uintptr) float64
	addr_QTextLength_OperatorEqual uintptr
	addr_QTextLength_OperatorNotEqual uintptr
	addr_QTextLength_ToQVariant uintptr
	addr_QTextLength_toQVariant uintptr
	pg_QTextLength_delete func(uintptr)
)

func init() {
	symaddr(&addr_QTextLength_new, "QTextLength_new")
	reg(&pg_QTextLength_new2, "QTextLength_new2")
	symaddr(&addr_QTextLength_new3, "QTextLength_new3")
	symaddr(&addr_QTextLength_type, "QTextLength_type")
	reg(&pg_QTextLength_value, "QTextLength_value")
	reg(&pg_QTextLength_rawValue, "QTextLength_rawValue")
	symaddr(&addr_QTextLength_OperatorEqual, "QTextLength_OperatorEqual")
	symaddr(&addr_QTextLength_OperatorNotEqual, "QTextLength_OperatorNotEqual")
	symaddr(&addr_QTextLength_ToQVariant, "QTextLength_ToQVariant")
	symaddr(&addr_QTextLength_toQVariant, "QTextLength_toQVariant")
	reg(&pg_QTextLength_delete, "QTextLength_delete")
}

type QTextLength struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QTextLength) UnsafeGoID() int64     { return this.vid }
func (this *QTextLength) UnsafeSetGoID(id int64) { this.vid = id }

func bindQTextLength(h uintptr) *QTextLength {
	if h == 0 {
		return nil
	}
	return &QTextLength{h: h}
}

func newownedQTextLength(h uintptr) *QTextLength {
	o := bindQTextLength(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QTextLength).Delete)
	}
	return o
}

func (this *QTextLength) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QTextLength) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQTextLength(h unsafe.Pointer) *QTextLength { return bindQTextLength(uintptr(h)) }
func UnsafeNewQTextLengthOwned(h unsafe.Pointer) *QTextLength { return newownedQTextLength(uintptr(h)) }

func (this *QTextLength) UnsafeRelease() { this.h = 0 }

func NewQTextLength() *QTextLength {
	__r, _, _ := purego.SyscallN(addr_QTextLength_new)
	__o := newownedQTextLength(__r)
	return __o
}

func NewQTextLength2(type_ QTextLength__Type, value float64) *QTextLength {
	__o := newownedQTextLength(pg_QTextLength_new2(int32(type_), value))
	return __o
}

func NewQTextLength3(arg1 *QTextLength) *QTextLength {
	__r, _, _ := purego.SyscallN(addr_QTextLength_new3, uintptr(arg1.UnsafePointer()))
	__o := newownedQTextLength(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QTextLength) Type() QTextLength__Type {
	__r0, _, _ := purego.SyscallN(addr_QTextLength_type, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QTextLength__Type(__r)
}

func (this *QTextLength) Value(maximumLength float64) float64 {
	__r := pg_QTextLength_value(this.h, maximumLength)
	runtime.KeepAlive(this)
	return __r
}

func (this *QTextLength) RawValue() float64 {
	__r := pg_QTextLength_rawValue(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QTextLength) OperatorEqual(other *QTextLength) bool {
	__r0, _, _ := purego.SyscallN(addr_QTextLength_OperatorEqual, this.h, uintptr(other.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return __r != 0
}

func (this *QTextLength) OperatorNotEqual(other *QTextLength) bool {
	__r0, _, _ := purego.SyscallN(addr_QTextLength_OperatorNotEqual, this.h, uintptr(other.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return __r != 0
}

func (this *QTextLength) ToQVariant() *QVariant {
	__r, _, _ := purego.SyscallN(addr_QTextLength_ToQVariant, this.h)
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QTextLength) ToQVariant2() *QVariant {
	__r, _, _ := purego.SyscallN(addr_QTextLength_toQVariant, this.h)
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QTextLength) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QTextLength_delete(this.h)
	}
	this.UnsafeRelease()
}
