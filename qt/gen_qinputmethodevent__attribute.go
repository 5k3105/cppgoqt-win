package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QInputMethodEvent__Attribute_new uintptr
	addr_QInputMethodEvent__Attribute_new2 uintptr
	addr_QInputMethodEvent__Attribute_new3 uintptr
	addr_QInputMethodEvent__Attribute_OperatorAssign uintptr
	addr_QInputMethodEvent__Attribute_OperatorEqual uintptr
	addr_QInputMethodEvent__Attribute_OperatorNotEqual uintptr
	addr_QInputMethodEvent__Attribute_type_get uintptr
	addr_QInputMethodEvent__Attribute_type_set uintptr
	addr_QInputMethodEvent__Attribute_start_get uintptr
	addr_QInputMethodEvent__Attribute_start_set uintptr
	addr_QInputMethodEvent__Attribute_length_get uintptr
	addr_QInputMethodEvent__Attribute_length_set uintptr
	addr_QInputMethodEvent__Attribute_value_get uintptr
	addr_QInputMethodEvent__Attribute_value_set uintptr
	pg_QInputMethodEvent__Attribute_delete func(uintptr)
)

func init() {
	symaddr(&addr_QInputMethodEvent__Attribute_new, "QInputMethodEvent__Attribute_new")
	symaddr(&addr_QInputMethodEvent__Attribute_new2, "QInputMethodEvent__Attribute_new2")
	symaddr(&addr_QInputMethodEvent__Attribute_new3, "QInputMethodEvent__Attribute_new3")
	symaddr(&addr_QInputMethodEvent__Attribute_OperatorAssign, "QInputMethodEvent__Attribute_OperatorAssign")
	symaddr(&addr_QInputMethodEvent__Attribute_OperatorEqual, "QInputMethodEvent__Attribute_OperatorEqual")
	symaddr(&addr_QInputMethodEvent__Attribute_OperatorNotEqual, "QInputMethodEvent__Attribute_OperatorNotEqual")
	symaddr(&addr_QInputMethodEvent__Attribute_type_get, "QInputMethodEvent__Attribute_type_get")
	symaddr(&addr_QInputMethodEvent__Attribute_type_set, "QInputMethodEvent__Attribute_type_set")
	symaddr(&addr_QInputMethodEvent__Attribute_start_get, "QInputMethodEvent__Attribute_start_get")
	symaddr(&addr_QInputMethodEvent__Attribute_start_set, "QInputMethodEvent__Attribute_start_set")
	symaddr(&addr_QInputMethodEvent__Attribute_length_get, "QInputMethodEvent__Attribute_length_get")
	symaddr(&addr_QInputMethodEvent__Attribute_length_set, "QInputMethodEvent__Attribute_length_set")
	symaddr(&addr_QInputMethodEvent__Attribute_value_get, "QInputMethodEvent__Attribute_value_get")
	symaddr(&addr_QInputMethodEvent__Attribute_value_set, "QInputMethodEvent__Attribute_value_set")
	reg(&pg_QInputMethodEvent__Attribute_delete, "QInputMethodEvent__Attribute_delete")
}

type QInputMethodEvent__Attribute struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QInputMethodEvent__Attribute) UnsafeGoID() int64     { return this.vid }
func (this *QInputMethodEvent__Attribute) UnsafeSetGoID(id int64) { this.vid = id }

func bindQInputMethodEvent__Attribute(h uintptr) *QInputMethodEvent__Attribute {
	if h == 0 {
		return nil
	}
	return &QInputMethodEvent__Attribute{h: h}
}

func newownedQInputMethodEvent__Attribute(h uintptr) *QInputMethodEvent__Attribute {
	o := bindQInputMethodEvent__Attribute(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QInputMethodEvent__Attribute).Delete)
	}
	return o
}

func (this *QInputMethodEvent__Attribute) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QInputMethodEvent__Attribute) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQInputMethodEvent__Attribute(h unsafe.Pointer) *QInputMethodEvent__Attribute { return bindQInputMethodEvent__Attribute(uintptr(h)) }
func UnsafeNewQInputMethodEvent__AttributeOwned(h unsafe.Pointer) *QInputMethodEvent__Attribute { return newownedQInputMethodEvent__Attribute(uintptr(h)) }

func (this *QInputMethodEvent__Attribute) UnsafeRelease() { this.h = 0 }

func NewQInputMethodEvent__Attribute(typ QInputMethodEvent__AttributeType, s int, l int, val *QVariant) *QInputMethodEvent__Attribute {
	__r, _, _ := purego.SyscallN(addr_QInputMethodEvent__Attribute_new, uintptr(typ), uintptr(s), uintptr(l), uintptr(val.UnsafePointer()))
	__o := newownedQInputMethodEvent__Attribute(__r)
	runtime.KeepAlive(val)
	return __o
}

func NewQInputMethodEvent__Attribute2(typ QInputMethodEvent__AttributeType, s int, l int) *QInputMethodEvent__Attribute {
	__r, _, _ := purego.SyscallN(addr_QInputMethodEvent__Attribute_new2, uintptr(typ), uintptr(s), uintptr(l))
	__o := newownedQInputMethodEvent__Attribute(__r)
	return __o
}

func NewQInputMethodEvent__Attribute3(arg1 *QInputMethodEvent__Attribute) *QInputMethodEvent__Attribute {
	__r, _, _ := purego.SyscallN(addr_QInputMethodEvent__Attribute_new3, uintptr(arg1.UnsafePointer()))
	__o := newownedQInputMethodEvent__Attribute(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QInputMethodEvent__Attribute) OperatorAssign(arg1 *QInputMethodEvent__Attribute) *QInputMethodEvent__Attribute {
	__r, _, _ := purego.SyscallN(addr_QInputMethodEvent__Attribute_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQInputMethodEvent__Attribute(__r)
}

func (this *QInputMethodEvent__Attribute) OperatorEqual(rhs *QInputMethodEvent__Attribute) bool {
	__r0, _, _ := purego.SyscallN(addr_QInputMethodEvent__Attribute_OperatorEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QInputMethodEvent__Attribute) OperatorNotEqual(rhs *QInputMethodEvent__Attribute) bool {
	__r0, _, _ := purego.SyscallN(addr_QInputMethodEvent__Attribute_OperatorNotEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QInputMethodEvent__Attribute) Type() QInputMethodEvent__AttributeType {
	__r0, _, _ := purego.SyscallN(addr_QInputMethodEvent__Attribute_type_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QInputMethodEvent__AttributeType(__r)
}

func (this *QInputMethodEvent__Attribute) SetType(type_ QInputMethodEvent__AttributeType) {
	purego.SyscallN(addr_QInputMethodEvent__Attribute_type_set, this.h, uintptr(type_))
	runtime.KeepAlive(this)
}

func (this *QInputMethodEvent__Attribute) Start() int {
	__r0, _, _ := purego.SyscallN(addr_QInputMethodEvent__Attribute_start_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QInputMethodEvent__Attribute) SetStart(start int) {
	purego.SyscallN(addr_QInputMethodEvent__Attribute_start_set, this.h, uintptr(start))
	runtime.KeepAlive(this)
}

func (this *QInputMethodEvent__Attribute) Length() int {
	__r0, _, _ := purego.SyscallN(addr_QInputMethodEvent__Attribute_length_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QInputMethodEvent__Attribute) SetLength(length int) {
	purego.SyscallN(addr_QInputMethodEvent__Attribute_length_set, this.h, uintptr(length))
	runtime.KeepAlive(this)
}

func (this *QInputMethodEvent__Attribute) Value() *QVariant {
	__r, _, _ := purego.SyscallN(addr_QInputMethodEvent__Attribute_value_get, this.h)
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QInputMethodEvent__Attribute) SetValue(value *QVariant) {
	purego.SyscallN(addr_QInputMethodEvent__Attribute_value_set, this.h, uintptr(value.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(value)
}

func (this *QInputMethodEvent__Attribute) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QInputMethodEvent__Attribute_delete(this.h)
	}
	this.UnsafeRelease()
}
