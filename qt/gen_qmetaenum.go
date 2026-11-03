package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QMetaEnum_new uintptr
	addr_QMetaEnum_new2 uintptr
	addr_QMetaEnum_name uintptr
	addr_QMetaEnum_enumName uintptr
	addr_QMetaEnum_metaType uintptr
	addr_QMetaEnum_isFlag uintptr
	addr_QMetaEnum_isScoped uintptr
	addr_QMetaEnum_is64Bit uintptr
	addr_QMetaEnum_keyCount uintptr
	addr_QMetaEnum_key uintptr
	addr_QMetaEnum_value uintptr
	addr_QMetaEnum_value64 uintptr
	addr_QMetaEnum_scope uintptr
	addr_QMetaEnum_keyToValue uintptr
	addr_QMetaEnum_keysToValue uintptr
	addr_QMetaEnum_keyToValue64 uintptr
	addr_QMetaEnum_keysToValue64 uintptr
	addr_QMetaEnum_valueToKey uintptr
	addr_QMetaEnum_valueToKeys uintptr
	addr_QMetaEnum_enclosingMetaObject uintptr
	addr_QMetaEnum_isValid uintptr
	addr_QMetaEnum_keyToValue2 uintptr
	addr_QMetaEnum_keysToValue2 uintptr
	pg_QMetaEnum_delete func(uintptr)
)

func init() {
	symaddr(&addr_QMetaEnum_new, "QMetaEnum_new")
	symaddr(&addr_QMetaEnum_new2, "QMetaEnum_new2")
	symaddr(&addr_QMetaEnum_name, "QMetaEnum_name")
	symaddr(&addr_QMetaEnum_enumName, "QMetaEnum_enumName")
	symaddr(&addr_QMetaEnum_metaType, "QMetaEnum_metaType")
	symaddr(&addr_QMetaEnum_isFlag, "QMetaEnum_isFlag")
	symaddr(&addr_QMetaEnum_isScoped, "QMetaEnum_isScoped")
	symaddr(&addr_QMetaEnum_is64Bit, "QMetaEnum_is64Bit")
	symaddr(&addr_QMetaEnum_keyCount, "QMetaEnum_keyCount")
	symaddr(&addr_QMetaEnum_key, "QMetaEnum_key")
	symaddr(&addr_QMetaEnum_value, "QMetaEnum_value")
	symaddr(&addr_QMetaEnum_value64, "QMetaEnum_value64")
	symaddr(&addr_QMetaEnum_scope, "QMetaEnum_scope")
	symaddr(&addr_QMetaEnum_keyToValue, "QMetaEnum_keyToValue")
	symaddr(&addr_QMetaEnum_keysToValue, "QMetaEnum_keysToValue")
	symaddr(&addr_QMetaEnum_keyToValue64, "QMetaEnum_keyToValue64")
	symaddr(&addr_QMetaEnum_keysToValue64, "QMetaEnum_keysToValue64")
	symaddr(&addr_QMetaEnum_valueToKey, "QMetaEnum_valueToKey")
	symaddr(&addr_QMetaEnum_valueToKeys, "QMetaEnum_valueToKeys")
	symaddr(&addr_QMetaEnum_enclosingMetaObject, "QMetaEnum_enclosingMetaObject")
	symaddr(&addr_QMetaEnum_isValid, "QMetaEnum_isValid")
	symaddr(&addr_QMetaEnum_keyToValue2, "QMetaEnum_keyToValue2")
	symaddr(&addr_QMetaEnum_keysToValue2, "QMetaEnum_keysToValue2")
	reg(&pg_QMetaEnum_delete, "QMetaEnum_delete")
}

type QMetaEnum struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QMetaEnum) UnsafeGoID() int64     { return this.vid }
func (this *QMetaEnum) UnsafeSetGoID(id int64) { this.vid = id }

func bindQMetaEnum(h uintptr) *QMetaEnum {
	if h == 0 {
		return nil
	}
	return &QMetaEnum{h: h}
}

func newownedQMetaEnum(h uintptr) *QMetaEnum {
	o := bindQMetaEnum(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QMetaEnum).Delete)
	}
	return o
}

func (this *QMetaEnum) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QMetaEnum) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQMetaEnum(h unsafe.Pointer) *QMetaEnum { return bindQMetaEnum(uintptr(h)) }
func UnsafeNewQMetaEnumOwned(h unsafe.Pointer) *QMetaEnum { return newownedQMetaEnum(uintptr(h)) }

func (this *QMetaEnum) UnsafeRelease() { this.h = 0 }

func NewQMetaEnum() *QMetaEnum {
	__r, _, _ := purego.SyscallN(addr_QMetaEnum_new)
	__o := newownedQMetaEnum(__r)
	return __o
}

func NewQMetaEnum2(arg1 *QMetaEnum) *QMetaEnum {
	__r, _, _ := purego.SyscallN(addr_QMetaEnum_new2, uintptr(arg1.UnsafePointer()))
	__o := newownedQMetaEnum(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QMetaEnum) Name() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QMetaEnum_name, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QMetaEnum) EnumName() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QMetaEnum_enumName, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QMetaEnum) MetaType() *QMetaType {
	__r, _, _ := purego.SyscallN(addr_QMetaEnum_metaType, this.h)
	runtime.KeepAlive(this)
	return newownedQMetaType(__r)
}

func (this *QMetaEnum) IsFlag() bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaEnum_isFlag, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMetaEnum) IsScoped() bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaEnum_isScoped, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMetaEnum) Is64Bit() bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaEnum_is64Bit, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMetaEnum) KeyCount() int {
	__r0, _, _ := purego.SyscallN(addr_QMetaEnum_keyCount, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QMetaEnum) Key(index int) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QMetaEnum_key, this.h, uintptr(index), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QMetaEnum) Value(index int) int {
	__r0, _, _ := purego.SyscallN(addr_QMetaEnum_value, this.h, uintptr(index))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QMetaEnum) Value64(index int) (uint64, bool) {
	var __b int32
	__r0, _, _ := purego.SyscallN(addr_QMetaEnum_value64, this.h, uintptr(index), uintptr(unsafe.Pointer(&__b)))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return uint64(__r), __b != 0
}

func (this *QMetaEnum) Scope() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QMetaEnum_scope, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QMetaEnum) KeyToValue(key string) int {
	key_b := append([]byte(key), 0)
	__r0, _, _ := purego.SyscallN(addr_QMetaEnum_keyToValue, this.h, uintptr(unsafe.Pointer(&key_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key_b)
	return int(__r)
}

func (this *QMetaEnum) KeysToValue(keys string) int {
	keys_b := append([]byte(keys), 0)
	__r0, _, _ := purego.SyscallN(addr_QMetaEnum_keysToValue, this.h, uintptr(unsafe.Pointer(&keys_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(keys_b)
	return int(__r)
}

func (this *QMetaEnum) KeyToValue64(key string) (uint64, bool) {
	key_b := append([]byte(key), 0)
	var __b int32
	__r0, _, _ := purego.SyscallN(addr_QMetaEnum_keyToValue64, this.h, uintptr(unsafe.Pointer(&key_b[0])), uintptr(unsafe.Pointer(&__b)))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(key_b)
	return uint64(__r), __b != 0
}

func (this *QMetaEnum) KeysToValue64(keys string) (uint64, bool) {
	keys_b := append([]byte(keys), 0)
	var __b int32
	__r0, _, _ := purego.SyscallN(addr_QMetaEnum_keysToValue64, this.h, uintptr(unsafe.Pointer(&keys_b[0])), uintptr(unsafe.Pointer(&__b)))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(keys_b)
	return uint64(__r), __b != 0
}

func (this *QMetaEnum) ValueToKey(value uint64) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QMetaEnum_valueToKey, this.h, uintptr(value), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QMetaEnum) ValueToKeys(value uint64) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QMetaEnum_valueToKeys, this.h, uintptr(value))
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QMetaEnum) EnclosingMetaObject() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QMetaEnum_enclosingMetaObject, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QMetaEnum) IsValid() bool {
	__r0, _, _ := purego.SyscallN(addr_QMetaEnum_isValid, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMetaEnum) KeyToValue2(key string) (int, bool) {
	key_b := append([]byte(key), 0)
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QMetaEnum_keyToValue2, this.h, uintptr(unsafe.Pointer(&key_b[0])), uintptr(unsafe.Pointer(&ok_out)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key_b)
	return int(__r), ok_out != 0
}

func (this *QMetaEnum) KeysToValue2(keys string) (int, bool) {
	keys_b := append([]byte(keys), 0)
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QMetaEnum_keysToValue2, this.h, uintptr(unsafe.Pointer(&keys_b[0])), uintptr(unsafe.Pointer(&ok_out)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(keys_b)
	return int(__r), ok_out != 0
}

func (this *QMetaEnum) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QMetaEnum_delete(this.h)
	}
	this.UnsafeRelease()
}
