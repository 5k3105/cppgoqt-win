package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QUuid_new uintptr
	addr_QUuid_new2 uintptr
	addr_QUuid_new3 uintptr
	addr_QUuid_new4 uintptr
	addr_QUuid_new5 uintptr
	addr_QUuid_new6 uintptr
	addr_QUuid_fromString uintptr
	addr_QUuid_toString uintptr
	addr_QUuid_toByteArray uintptr
	addr_QUuid_toBytes uintptr
	addr_QUuid_toRfc4122 uintptr
	addr_QUuid_fromRfc4122 uintptr
	addr_QUuid_isNull uintptr
	addr_QUuid_ToGUID uintptr
	addr_QUuid_createUuid uintptr
	addr_QUuid_createUuidV5 uintptr
	addr_QUuid_createUuidV3 uintptr
	addr_QUuid_createUuidV7 uintptr
	addr_QUuid_variant uintptr
	addr_QUuid_version uintptr
	addr_QUuid_OperatorAssign uintptr
	addr_QUuid_toString2 uintptr
	addr_QUuid_toByteArray2 uintptr
	addr_QUuid_toBytes2 uintptr
	addr_QUuid_OperatorEqual uintptr
	addr_QUuid_OperatorNotEqual uintptr
	addr_QUuid_OperatorLesser uintptr
	addr_QUuid_OperatorGreater uintptr
	addr_QUuid_OperatorLesserOrEqual uintptr
	addr_QUuid_OperatorGreaterOrEqual uintptr
	addr_QUuid_data1_get uintptr
	addr_QUuid_data1_set uintptr
	addr_QUuid_data2_get uintptr
	addr_QUuid_data2_set uintptr
	addr_QUuid_data3_get uintptr
	addr_QUuid_data3_set uintptr
	pg_QUuid_delete func(uintptr)
)

func init() {
	symaddr(&addr_QUuid_new, "QUuid_new")
	symaddr(&addr_QUuid_new2, "QUuid_new2")
	symaddr(&addr_QUuid_new3, "QUuid_new3")
	symaddr(&addr_QUuid_new4, "QUuid_new4")
	symaddr(&addr_QUuid_new5, "QUuid_new5")
	symaddr(&addr_QUuid_new6, "QUuid_new6")
	symaddr(&addr_QUuid_fromString, "QUuid_fromString")
	symaddr(&addr_QUuid_toString, "QUuid_toString")
	symaddr(&addr_QUuid_toByteArray, "QUuid_toByteArray")
	symaddr(&addr_QUuid_toBytes, "QUuid_toBytes")
	symaddr(&addr_QUuid_toRfc4122, "QUuid_toRfc4122")
	symaddr(&addr_QUuid_fromRfc4122, "QUuid_fromRfc4122")
	symaddr(&addr_QUuid_isNull, "QUuid_isNull")
	symaddr(&addr_QUuid_ToGUID, "QUuid_ToGUID")
	symaddr(&addr_QUuid_createUuid, "QUuid_createUuid")
	symaddr(&addr_QUuid_createUuidV5, "QUuid_createUuidV5")
	symaddr(&addr_QUuid_createUuidV3, "QUuid_createUuidV3")
	symaddr(&addr_QUuid_createUuidV7, "QUuid_createUuidV7")
	symaddr(&addr_QUuid_variant, "QUuid_variant")
	symaddr(&addr_QUuid_version, "QUuid_version")
	symaddr(&addr_QUuid_OperatorAssign, "QUuid_OperatorAssign")
	symaddr(&addr_QUuid_toString2, "QUuid_toString2")
	symaddr(&addr_QUuid_toByteArray2, "QUuid_toByteArray2")
	symaddr(&addr_QUuid_toBytes2, "QUuid_toBytes2")
	symaddr(&addr_QUuid_OperatorEqual, "QUuid_OperatorEqual")
	symaddr(&addr_QUuid_OperatorNotEqual, "QUuid_OperatorNotEqual")
	symaddr(&addr_QUuid_OperatorLesser, "QUuid_OperatorLesser")
	symaddr(&addr_QUuid_OperatorGreater, "QUuid_OperatorGreater")
	symaddr(&addr_QUuid_OperatorLesserOrEqual, "QUuid_OperatorLesserOrEqual")
	symaddr(&addr_QUuid_OperatorGreaterOrEqual, "QUuid_OperatorGreaterOrEqual")
	symaddr(&addr_QUuid_data1_get, "QUuid_data1_get")
	symaddr(&addr_QUuid_data1_set, "QUuid_data1_set")
	symaddr(&addr_QUuid_data2_get, "QUuid_data2_get")
	symaddr(&addr_QUuid_data2_set, "QUuid_data2_set")
	symaddr(&addr_QUuid_data3_get, "QUuid_data3_get")
	symaddr(&addr_QUuid_data3_set, "QUuid_data3_set")
	reg(&pg_QUuid_delete, "QUuid_delete")
}

type QUuid struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QUuid) UnsafeGoID() int64     { return this.vid }
func (this *QUuid) UnsafeSetGoID(id int64) { this.vid = id }

func bindQUuid(h uintptr) *QUuid {
	if h == 0 {
		return nil
	}
	return &QUuid{h: h}
}

func newownedQUuid(h uintptr) *QUuid {
	o := bindQUuid(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QUuid).Delete)
	}
	return o
}

func (this *QUuid) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QUuid) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQUuid(h unsafe.Pointer) *QUuid { return bindQUuid(uintptr(h)) }
func UnsafeNewQUuidOwned(h unsafe.Pointer) *QUuid { return newownedQUuid(uintptr(h)) }

func (this *QUuid) UnsafeRelease() { this.h = 0 }

func NewQUuid() *QUuid {
	__r, _, _ := purego.SyscallN(addr_QUuid_new)
	__o := newownedQUuid(__r)
	return __o
}

func NewQUuid2(l uint, w1 uint16, w2 uint16, b1 byte, b2 byte, b3 byte, b4 byte, b5 byte, b6 byte, b7 byte, b8 byte) *QUuid {
	__r, _, _ := purego.SyscallN(addr_QUuid_new2, uintptr(l), uintptr(w1), uintptr(w2), uintptr(b1), uintptr(b2), uintptr(b3), uintptr(b4), uintptr(b5), uintptr(b6), uintptr(b7), uintptr(b8))
	__o := newownedQUuid(__r)
	return __o
}

func NewQUuid3(id128 *QUuid__Id128Bytes) *QUuid {
	__r, _, _ := purego.SyscallN(addr_QUuid_new3, uintptr(id128.UnsafePointer()))
	__o := newownedQUuid(__r)
	runtime.KeepAlive(id128)
	return __o
}

func NewQUuid4(string_ string) *QUuid {
	var string__p uintptr
	if len(string_) > 0 {
		string__p = uintptr(unsafe.Pointer(unsafe.StringData(string_)))
	}
	__r, _, _ := purego.SyscallN(addr_QUuid_new4, string__p, uintptr(len(string_)))
	__o := newownedQUuid(__r)
	runtime.KeepAlive(string_)
	return __o
}

func NewQUuid5(arg1 *QUuid) *QUuid {
	__r, _, _ := purego.SyscallN(addr_QUuid_new5, uintptr(arg1.UnsafePointer()))
	__o := newownedQUuid(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func NewQUuid6(id128 *QUuid__Id128Bytes, order QSysInfo__Endian) *QUuid {
	__r, _, _ := purego.SyscallN(addr_QUuid_new6, uintptr(id128.UnsafePointer()), uintptr(order))
	__o := newownedQUuid(__r)
	runtime.KeepAlive(id128)
	return __o
}

func QUuid_FromString(string_ string) *QUuid {
	var string__p uintptr
	if len(string_) > 0 {
		string__p = uintptr(unsafe.Pointer(unsafe.StringData(string_)))
	}
	__r, _, _ := purego.SyscallN(addr_QUuid_fromString, string__p, uintptr(len(string_)))
	runtime.KeepAlive(string_)
	return newownedQUuid(__r)
}

func (this *QUuid) ToString() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QUuid_toString, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QUuid) ToByteArray() *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QUuid_toByteArray, this.h)
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QUuid) ToBytes() *QUuid__Id128Bytes {
	__r, _, _ := purego.SyscallN(addr_QUuid_toBytes, this.h)
	runtime.KeepAlive(this)
	return newownedQUuid__Id128Bytes(__r)
}

func (this *QUuid) ToRfc4122() *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QUuid_toRfc4122, this.h)
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func QUuid_FromRfc4122(arg1 *QByteArrayView) *QUuid {
	__r, _, _ := purego.SyscallN(addr_QUuid_fromRfc4122, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(arg1)
	return newownedQUuid(__r)
}

func (this *QUuid) IsNull() bool {
	__r0, _, _ := purego.SyscallN(addr_QUuid_isNull, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QUuid) ToGUID() *_GUID {
	__r, _, _ := purego.SyscallN(addr_QUuid_ToGUID, this.h)
	runtime.KeepAlive(this)
	return newowned_GUID(__r)
}

func QUuid_CreateUuid() *QUuid {
	__r, _, _ := purego.SyscallN(addr_QUuid_createUuid)
	return newownedQUuid(__r)
}

func QUuid_CreateUuidV5(ns *QUuid, baseData *QByteArrayView) *QUuid {
	__r, _, _ := purego.SyscallN(addr_QUuid_createUuidV5, uintptr(ns.UnsafePointer()), uintptr(baseData.UnsafePointer()))
	runtime.KeepAlive(ns)
	runtime.KeepAlive(baseData)
	return newownedQUuid(__r)
}

func QUuid_CreateUuidV3(ns *QUuid, baseData *QByteArrayView) *QUuid {
	__r, _, _ := purego.SyscallN(addr_QUuid_createUuidV3, uintptr(ns.UnsafePointer()), uintptr(baseData.UnsafePointer()))
	runtime.KeepAlive(ns)
	runtime.KeepAlive(baseData)
	return newownedQUuid(__r)
}

func QUuid_CreateUuidV7() *QUuid {
	__r, _, _ := purego.SyscallN(addr_QUuid_createUuidV7)
	return newownedQUuid(__r)
}

func (this *QUuid) Variant() QUuid__Variant {
	__r0, _, _ := purego.SyscallN(addr_QUuid_variant, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QUuid__Variant(__r)
}

func (this *QUuid) Version() QUuid__Version {
	__r0, _, _ := purego.SyscallN(addr_QUuid_version, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QUuid__Version(__r)
}

func (this *QUuid) OperatorAssign(arg1 *QUuid) *QUuid {
	__r, _, _ := purego.SyscallN(addr_QUuid_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQUuid(__r)
}

func (this *QUuid) ToStringWithMode(mode QUuid__StringFormat) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QUuid_toString2, this.h, uintptr(mode), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QUuid) ToByteArrayWithMode(mode QUuid__StringFormat) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QUuid_toByteArray2, this.h, uintptr(mode))
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QUuid) ToBytesWithOrder(order QSysInfo__Endian) *QUuid__Id128Bytes {
	__r, _, _ := purego.SyscallN(addr_QUuid_toBytes2, this.h, uintptr(order))
	runtime.KeepAlive(this)
	return newownedQUuid__Id128Bytes(__r)
}

func (this *QUuid) OperatorEqual(rhs *QUuid) bool {
	__r0, _, _ := purego.SyscallN(addr_QUuid_OperatorEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QUuid) OperatorNotEqual(rhs *QUuid) bool {
	__r0, _, _ := purego.SyscallN(addr_QUuid_OperatorNotEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QUuid) OperatorLesser(rhs *QUuid) bool {
	__r0, _, _ := purego.SyscallN(addr_QUuid_OperatorLesser, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QUuid) OperatorGreater(rhs *QUuid) bool {
	__r0, _, _ := purego.SyscallN(addr_QUuid_OperatorGreater, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QUuid) OperatorLesserOrEqual(rhs *QUuid) bool {
	__r0, _, _ := purego.SyscallN(addr_QUuid_OperatorLesserOrEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QUuid) OperatorGreaterOrEqual(rhs *QUuid) bool {
	__r0, _, _ := purego.SyscallN(addr_QUuid_OperatorGreaterOrEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QUuid) Data1() uint {
	__r0, _, _ := purego.SyscallN(addr_QUuid_data1_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint(__r)
}

func (this *QUuid) SetData1(data1 uint) {
	purego.SyscallN(addr_QUuid_data1_set, this.h, uintptr(data1))
	runtime.KeepAlive(this)
}

func (this *QUuid) Data2() uint16 {
	__r0, _, _ := purego.SyscallN(addr_QUuid_data2_get, this.h)
	__r := int16(uint16(__r0))
	runtime.KeepAlive(this)
	return uint16(__r)
}

func (this *QUuid) SetData2(data2 uint16) {
	purego.SyscallN(addr_QUuid_data2_set, this.h, uintptr(data2))
	runtime.KeepAlive(this)
}

func (this *QUuid) Data3() uint16 {
	__r0, _, _ := purego.SyscallN(addr_QUuid_data3_get, this.h)
	__r := int16(uint16(__r0))
	runtime.KeepAlive(this)
	return uint16(__r)
}

func (this *QUuid) SetData3(data3 uint16) {
	purego.SyscallN(addr_QUuid_data3_set, this.h, uintptr(data3))
	runtime.KeepAlive(this)
}

func (this *QUuid) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QUuid_delete(this.h)
	}
	this.UnsafeRelease()
}
