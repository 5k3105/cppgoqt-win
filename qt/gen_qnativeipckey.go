package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QNativeIpcKey_new uintptr
	addr_QNativeIpcKey_new2 uintptr
	addr_QNativeIpcKey_new3 uintptr
	addr_QNativeIpcKey_new4 uintptr
	addr_QNativeIpcKey_new5 uintptr
	addr_QNativeIpcKey_legacyDefaultTypeForOs uintptr
	addr_QNativeIpcKey_OperatorAssign uintptr
	addr_QNativeIpcKey_swap uintptr
	addr_QNativeIpcKey_isEmpty uintptr
	addr_QNativeIpcKey_isValid uintptr
	addr_QNativeIpcKey_type uintptr
	addr_QNativeIpcKey_setType uintptr
	addr_QNativeIpcKey_nativeKey uintptr
	addr_QNativeIpcKey_setNativeKey uintptr
	addr_QNativeIpcKey_toString uintptr
	addr_QNativeIpcKey_fromString uintptr
	addr_QNativeIpcKey_OperatorEqual uintptr
	addr_QNativeIpcKey_OperatorNotEqual uintptr
	pg_QNativeIpcKey_delete func(uintptr)
)

func init() {
	symaddr(&addr_QNativeIpcKey_new, "QNativeIpcKey_new")
	symaddr(&addr_QNativeIpcKey_new2, "QNativeIpcKey_new2")
	symaddr(&addr_QNativeIpcKey_new3, "QNativeIpcKey_new3")
	symaddr(&addr_QNativeIpcKey_new4, "QNativeIpcKey_new4")
	symaddr(&addr_QNativeIpcKey_new5, "QNativeIpcKey_new5")
	symaddr(&addr_QNativeIpcKey_legacyDefaultTypeForOs, "QNativeIpcKey_legacyDefaultTypeForOs")
	symaddr(&addr_QNativeIpcKey_OperatorAssign, "QNativeIpcKey_OperatorAssign")
	symaddr(&addr_QNativeIpcKey_swap, "QNativeIpcKey_swap")
	symaddr(&addr_QNativeIpcKey_isEmpty, "QNativeIpcKey_isEmpty")
	symaddr(&addr_QNativeIpcKey_isValid, "QNativeIpcKey_isValid")
	symaddr(&addr_QNativeIpcKey_type, "QNativeIpcKey_type")
	symaddr(&addr_QNativeIpcKey_setType, "QNativeIpcKey_setType")
	symaddr(&addr_QNativeIpcKey_nativeKey, "QNativeIpcKey_nativeKey")
	symaddr(&addr_QNativeIpcKey_setNativeKey, "QNativeIpcKey_setNativeKey")
	symaddr(&addr_QNativeIpcKey_toString, "QNativeIpcKey_toString")
	symaddr(&addr_QNativeIpcKey_fromString, "QNativeIpcKey_fromString")
	symaddr(&addr_QNativeIpcKey_OperatorEqual, "QNativeIpcKey_OperatorEqual")
	symaddr(&addr_QNativeIpcKey_OperatorNotEqual, "QNativeIpcKey_OperatorNotEqual")
	reg(&pg_QNativeIpcKey_delete, "QNativeIpcKey_delete")
}

type QNativeIpcKey struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QNativeIpcKey) UnsafeGoID() int64     { return this.vid }
func (this *QNativeIpcKey) UnsafeSetGoID(id int64) { this.vid = id }

func bindQNativeIpcKey(h uintptr) *QNativeIpcKey {
	if h == 0 {
		return nil
	}
	return &QNativeIpcKey{h: h}
}

func newownedQNativeIpcKey(h uintptr) *QNativeIpcKey {
	o := bindQNativeIpcKey(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QNativeIpcKey).Delete)
	}
	return o
}

func (this *QNativeIpcKey) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QNativeIpcKey) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQNativeIpcKey(h unsafe.Pointer) *QNativeIpcKey { return bindQNativeIpcKey(uintptr(h)) }
func UnsafeNewQNativeIpcKeyOwned(h unsafe.Pointer) *QNativeIpcKey { return newownedQNativeIpcKey(uintptr(h)) }

func (this *QNativeIpcKey) UnsafeRelease() { this.h = 0 }

func NewQNativeIpcKey() *QNativeIpcKey {
	__r, _, _ := purego.SyscallN(addr_QNativeIpcKey_new)
	__o := newownedQNativeIpcKey(__r)
	return __o
}

func NewQNativeIpcKey2(type_ QNativeIpcKey__Type) *QNativeIpcKey {
	__r, _, _ := purego.SyscallN(addr_QNativeIpcKey_new2, uintptr(type_))
	__o := newownedQNativeIpcKey(__r)
	return __o
}

func NewQNativeIpcKey3(k string) *QNativeIpcKey {
	var k_p uintptr
	if len(k) > 0 {
		k_p = uintptr(unsafe.Pointer(unsafe.StringData(k)))
	}
	__r, _, _ := purego.SyscallN(addr_QNativeIpcKey_new3, k_p, uintptr(len(k)))
	__o := newownedQNativeIpcKey(__r)
	runtime.KeepAlive(k)
	return __o
}

func NewQNativeIpcKey4(other *QNativeIpcKey) *QNativeIpcKey {
	__r, _, _ := purego.SyscallN(addr_QNativeIpcKey_new4, uintptr(other.UnsafePointer()))
	__o := newownedQNativeIpcKey(__r)
	runtime.KeepAlive(other)
	return __o
}

func NewQNativeIpcKey5(k string, type_ QNativeIpcKey__Type) *QNativeIpcKey {
	var k_p uintptr
	if len(k) > 0 {
		k_p = uintptr(unsafe.Pointer(unsafe.StringData(k)))
	}
	__r, _, _ := purego.SyscallN(addr_QNativeIpcKey_new5, k_p, uintptr(len(k)), uintptr(type_))
	__o := newownedQNativeIpcKey(__r)
	runtime.KeepAlive(k)
	return __o
}

func QNativeIpcKey_LegacyDefaultTypeForOs() QNativeIpcKey__Type {
	__r0, _, _ := purego.SyscallN(addr_QNativeIpcKey_legacyDefaultTypeForOs)
	__r := int32(uint32(__r0))
	return QNativeIpcKey__Type(__r)
}

func (this *QNativeIpcKey) OperatorAssign(other *QNativeIpcKey) *QNativeIpcKey {
	__r, _, _ := purego.SyscallN(addr_QNativeIpcKey_OperatorAssign, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return bindQNativeIpcKey(__r)
}

func (this *QNativeIpcKey) Swap(other *QNativeIpcKey) {
	purego.SyscallN(addr_QNativeIpcKey_swap, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
}

func (this *QNativeIpcKey) IsEmpty() bool {
	__r0, _, _ := purego.SyscallN(addr_QNativeIpcKey_isEmpty, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QNativeIpcKey) IsValid() bool {
	__r0, _, _ := purego.SyscallN(addr_QNativeIpcKey_isValid, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QNativeIpcKey) Type() QNativeIpcKey__Type {
	__r0, _, _ := purego.SyscallN(addr_QNativeIpcKey_type, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QNativeIpcKey__Type(__r)
}

func (this *QNativeIpcKey) SetType(type_ QNativeIpcKey__Type) {
	purego.SyscallN(addr_QNativeIpcKey_setType, this.h, uintptr(type_))
	runtime.KeepAlive(this)
}

func (this *QNativeIpcKey) NativeKey() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QNativeIpcKey_nativeKey, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QNativeIpcKey) SetNativeKey(newKey string) {
	var newKey_p uintptr
	if len(newKey) > 0 {
		newKey_p = uintptr(unsafe.Pointer(unsafe.StringData(newKey)))
	}
	purego.SyscallN(addr_QNativeIpcKey_setNativeKey, this.h, newKey_p, uintptr(len(newKey)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(newKey)
}

func (this *QNativeIpcKey) ToString() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QNativeIpcKey_toString, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func QNativeIpcKey_FromString(string_ string) *QNativeIpcKey {
	var string__p uintptr
	if len(string_) > 0 {
		string__p = uintptr(unsafe.Pointer(unsafe.StringData(string_)))
	}
	__r, _, _ := purego.SyscallN(addr_QNativeIpcKey_fromString, string__p, uintptr(len(string_)))
	runtime.KeepAlive(string_)
	return newownedQNativeIpcKey(__r)
}

func (this *QNativeIpcKey) OperatorEqual(rhs *QNativeIpcKey) bool {
	__r0, _, _ := purego.SyscallN(addr_QNativeIpcKey_OperatorEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QNativeIpcKey) OperatorNotEqual(rhs *QNativeIpcKey) bool {
	__r0, _, _ := purego.SyscallN(addr_QNativeIpcKey_OperatorNotEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QNativeIpcKey) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QNativeIpcKey_delete(this.h)
	}
	this.UnsafeRelease()
}
