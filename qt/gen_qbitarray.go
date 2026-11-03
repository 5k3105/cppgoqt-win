package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QBitArray_new uintptr
	addr_QBitArray_new2 uintptr
	addr_QBitArray_new3 uintptr
	addr_QBitArray_new4 uintptr
	addr_QBitArray_OperatorAssign uintptr
	addr_QBitArray_swap uintptr
	addr_QBitArray_size uintptr
	addr_QBitArray_count uintptr
	addr_QBitArray_count2 uintptr
	addr_QBitArray_isEmpty uintptr
	addr_QBitArray_isNull uintptr
	addr_QBitArray_resize uintptr
	addr_QBitArray_detach uintptr
	addr_QBitArray_isDetached uintptr
	addr_QBitArray_clear uintptr
	addr_QBitArray_testBit uintptr
	addr_QBitArray_setBit uintptr
	addr_QBitArray_setBit2 uintptr
	addr_QBitArray_clearBit uintptr
	addr_QBitArray_toggleBit uintptr
	addr_QBitArray_at uintptr
	addr_QBitArray_OperatorSubscript uintptr
	addr_QBitArray_OperatorSubscript2 uintptr
	addr_QBitArray_OperatorBitwiseAndAssign uintptr
	addr_QBitArray_OperatorBitwiseOrAssign uintptr
	addr_QBitArray_OperatorBitwiseXorAssign uintptr
	addr_QBitArray_fill uintptr
	addr_QBitArray_fill2 uintptr
	addr_QBitArray_truncate uintptr
	addr_QBitArray_bits uintptr
	addr_QBitArray_fromBits uintptr
	addr_QBitArray_toUInt32 uintptr
	addr_QBitArray_fill3 uintptr
	addr_QBitArray_toUInt322 uintptr
	addr_QBitArray_OperatorBitwiseAnd uintptr
	addr_QBitArray_OperatorBitwiseOr uintptr
	addr_QBitArray_OperatorBitwiseXor uintptr
	addr_QBitArray_OperatorBitwiseNot uintptr
	addr_QBitArray_OperatorEqual uintptr
	addr_QBitArray_OperatorNotEqual uintptr
	pg_QBitArray_delete func(uintptr)
)

func init() {
	symaddr(&addr_QBitArray_new, "QBitArray_new")
	symaddr(&addr_QBitArray_new2, "QBitArray_new2")
	symaddr(&addr_QBitArray_new3, "QBitArray_new3")
	symaddr(&addr_QBitArray_new4, "QBitArray_new4")
	symaddr(&addr_QBitArray_OperatorAssign, "QBitArray_OperatorAssign")
	symaddr(&addr_QBitArray_swap, "QBitArray_swap")
	symaddr(&addr_QBitArray_size, "QBitArray_size")
	symaddr(&addr_QBitArray_count, "QBitArray_count")
	symaddr(&addr_QBitArray_count2, "QBitArray_count2")
	symaddr(&addr_QBitArray_isEmpty, "QBitArray_isEmpty")
	symaddr(&addr_QBitArray_isNull, "QBitArray_isNull")
	symaddr(&addr_QBitArray_resize, "QBitArray_resize")
	symaddr(&addr_QBitArray_detach, "QBitArray_detach")
	symaddr(&addr_QBitArray_isDetached, "QBitArray_isDetached")
	symaddr(&addr_QBitArray_clear, "QBitArray_clear")
	symaddr(&addr_QBitArray_testBit, "QBitArray_testBit")
	symaddr(&addr_QBitArray_setBit, "QBitArray_setBit")
	symaddr(&addr_QBitArray_setBit2, "QBitArray_setBit2")
	symaddr(&addr_QBitArray_clearBit, "QBitArray_clearBit")
	symaddr(&addr_QBitArray_toggleBit, "QBitArray_toggleBit")
	symaddr(&addr_QBitArray_at, "QBitArray_at")
	symaddr(&addr_QBitArray_OperatorSubscript, "QBitArray_OperatorSubscript")
	symaddr(&addr_QBitArray_OperatorSubscript2, "QBitArray_OperatorSubscript2")
	symaddr(&addr_QBitArray_OperatorBitwiseAndAssign, "QBitArray_OperatorBitwiseAndAssign")
	symaddr(&addr_QBitArray_OperatorBitwiseOrAssign, "QBitArray_OperatorBitwiseOrAssign")
	symaddr(&addr_QBitArray_OperatorBitwiseXorAssign, "QBitArray_OperatorBitwiseXorAssign")
	symaddr(&addr_QBitArray_fill, "QBitArray_fill")
	symaddr(&addr_QBitArray_fill2, "QBitArray_fill2")
	symaddr(&addr_QBitArray_truncate, "QBitArray_truncate")
	symaddr(&addr_QBitArray_bits, "QBitArray_bits")
	symaddr(&addr_QBitArray_fromBits, "QBitArray_fromBits")
	symaddr(&addr_QBitArray_toUInt32, "QBitArray_toUInt32")
	symaddr(&addr_QBitArray_fill3, "QBitArray_fill3")
	symaddr(&addr_QBitArray_toUInt322, "QBitArray_toUInt322")
	symaddr(&addr_QBitArray_OperatorBitwiseAnd, "QBitArray_OperatorBitwiseAnd")
	symaddr(&addr_QBitArray_OperatorBitwiseOr, "QBitArray_OperatorBitwiseOr")
	symaddr(&addr_QBitArray_OperatorBitwiseXor, "QBitArray_OperatorBitwiseXor")
	symaddr(&addr_QBitArray_OperatorBitwiseNot, "QBitArray_OperatorBitwiseNot")
	symaddr(&addr_QBitArray_OperatorEqual, "QBitArray_OperatorEqual")
	symaddr(&addr_QBitArray_OperatorNotEqual, "QBitArray_OperatorNotEqual")
	reg(&pg_QBitArray_delete, "QBitArray_delete")
}

type QBitArray struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QBitArray) UnsafeGoID() int64     { return this.vid }
func (this *QBitArray) UnsafeSetGoID(id int64) { this.vid = id }

func bindQBitArray(h uintptr) *QBitArray {
	if h == 0 {
		return nil
	}
	return &QBitArray{h: h}
}

func newownedQBitArray(h uintptr) *QBitArray {
	o := bindQBitArray(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QBitArray).Delete)
	}
	return o
}

func (this *QBitArray) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QBitArray) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQBitArray(h unsafe.Pointer) *QBitArray { return bindQBitArray(uintptr(h)) }
func UnsafeNewQBitArrayOwned(h unsafe.Pointer) *QBitArray { return newownedQBitArray(uintptr(h)) }

func (this *QBitArray) UnsafeRelease() { this.h = 0 }

func NewQBitArray() *QBitArray {
	__r, _, _ := purego.SyscallN(addr_QBitArray_new)
	__o := newownedQBitArray(__r)
	return __o
}

func NewQBitArray2(size int64) *QBitArray {
	__r, _, _ := purego.SyscallN(addr_QBitArray_new2, uintptr(size))
	__o := newownedQBitArray(__r)
	return __o
}

func NewQBitArray3(other *QBitArray) *QBitArray {
	__r, _, _ := purego.SyscallN(addr_QBitArray_new3, uintptr(other.UnsafePointer()))
	__o := newownedQBitArray(__r)
	runtime.KeepAlive(other)
	return __o
}

func NewQBitArray4(size int64, val bool) *QBitArray {
	__r, _, _ := purego.SyscallN(addr_QBitArray_new4, uintptr(size), uintptr(b2i(val)))
	__o := newownedQBitArray(__r)
	return __o
}

func (this *QBitArray) OperatorAssign(other *QBitArray) *QBitArray {
	__r, _, _ := purego.SyscallN(addr_QBitArray_OperatorAssign, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return bindQBitArray(__r)
}

func (this *QBitArray) Swap(other *QBitArray) {
	purego.SyscallN(addr_QBitArray_swap, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
}

func (this *QBitArray) Size() int64 {
	__r0, _, _ := purego.SyscallN(addr_QBitArray_size, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QBitArray) Count() int64 {
	__r0, _, _ := purego.SyscallN(addr_QBitArray_count, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QBitArray) CountWithOn(on bool) int64 {
	__r0, _, _ := purego.SyscallN(addr_QBitArray_count2, this.h, uintptr(b2i(on)))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QBitArray) IsEmpty() bool {
	__r0, _, _ := purego.SyscallN(addr_QBitArray_isEmpty, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QBitArray) IsNull() bool {
	__r0, _, _ := purego.SyscallN(addr_QBitArray_isNull, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QBitArray) Resize(size int64) {
	purego.SyscallN(addr_QBitArray_resize, this.h, uintptr(size))
	runtime.KeepAlive(this)
}

func (this *QBitArray) Detach() {
	purego.SyscallN(addr_QBitArray_detach, this.h)
	runtime.KeepAlive(this)
}

func (this *QBitArray) IsDetached() bool {
	__r0, _, _ := purego.SyscallN(addr_QBitArray_isDetached, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QBitArray) Clear() {
	purego.SyscallN(addr_QBitArray_clear, this.h)
	runtime.KeepAlive(this)
}

func (this *QBitArray) TestBit(i int64) bool {
	__r0, _, _ := purego.SyscallN(addr_QBitArray_testBit, this.h, uintptr(i))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QBitArray) SetBit(i int64) {
	purego.SyscallN(addr_QBitArray_setBit, this.h, uintptr(i))
	runtime.KeepAlive(this)
}

func (this *QBitArray) SetBit2(i int64, val bool) {
	purego.SyscallN(addr_QBitArray_setBit2, this.h, uintptr(i), uintptr(b2i(val)))
	runtime.KeepAlive(this)
}

func (this *QBitArray) ClearBit(i int64) {
	purego.SyscallN(addr_QBitArray_clearBit, this.h, uintptr(i))
	runtime.KeepAlive(this)
}

func (this *QBitArray) ToggleBit(i int64) bool {
	__r0, _, _ := purego.SyscallN(addr_QBitArray_toggleBit, this.h, uintptr(i))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QBitArray) At(i int64) bool {
	__r0, _, _ := purego.SyscallN(addr_QBitArray_at, this.h, uintptr(i))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QBitArray) OperatorSubscript(i int64) *QBitRef {
	__r, _, _ := purego.SyscallN(addr_QBitArray_OperatorSubscript, this.h, uintptr(i))
	runtime.KeepAlive(this)
	return newownedQBitRef(__r)
}

func (this *QBitArray) OperatorSubscriptWith(i int64) bool {
	__r0, _, _ := purego.SyscallN(addr_QBitArray_OperatorSubscript2, this.h, uintptr(i))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QBitArray) OperatorBitwiseAndAssign(arg1 *QBitArray) *QBitArray {
	__r, _, _ := purego.SyscallN(addr_QBitArray_OperatorBitwiseAndAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQBitArray(__r)
}

func (this *QBitArray) OperatorBitwiseOrAssign(arg1 *QBitArray) *QBitArray {
	__r, _, _ := purego.SyscallN(addr_QBitArray_OperatorBitwiseOrAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQBitArray(__r)
}

func (this *QBitArray) OperatorBitwiseXorAssign(arg1 *QBitArray) *QBitArray {
	__r, _, _ := purego.SyscallN(addr_QBitArray_OperatorBitwiseXorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQBitArray(__r)
}

func (this *QBitArray) Fill(aval bool) bool {
	__r0, _, _ := purego.SyscallN(addr_QBitArray_fill, this.h, uintptr(b2i(aval)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QBitArray) Fill2(val bool, first int64, last int64) {
	purego.SyscallN(addr_QBitArray_fill2, this.h, uintptr(b2i(val)), uintptr(first), uintptr(last))
	runtime.KeepAlive(this)
}

func (this *QBitArray) Truncate(pos int64) {
	purego.SyscallN(addr_QBitArray_truncate, this.h, uintptr(pos))
	runtime.KeepAlive(this)
}

func (this *QBitArray) Bits() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QBitArray_bits, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func QBitArray_FromBits(data string, len_ int64) *QBitArray {
	data_b := append([]byte(data), 0)
	__r, _, _ := purego.SyscallN(addr_QBitArray_fromBits, uintptr(unsafe.Pointer(&data_b[0])), uintptr(len_))
	runtime.KeepAlive(data_b)
	return newownedQBitArray(__r)
}

func (this *QBitArray) ToUInt32(endianness QSysInfo__Endian) uint {
	__r0, _, _ := purego.SyscallN(addr_QBitArray_toUInt32, this.h, uintptr(endianness))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint(__r)
}

func (this *QBitArray) Fill3(aval bool, asize int64) bool {
	__r0, _, _ := purego.SyscallN(addr_QBitArray_fill3, this.h, uintptr(b2i(aval)), uintptr(asize))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QBitArray) ToUInt322(endianness QSysInfo__Endian) (uint, bool) {
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QBitArray_toUInt322, this.h, uintptr(endianness), uintptr(unsafe.Pointer(&ok_out)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint(__r), ok_out != 0
}

func (this *QBitArray) OperatorBitwiseAnd(a2 *QBitArray) *QBitArray {
	__r, _, _ := purego.SyscallN(addr_QBitArray_OperatorBitwiseAnd, this.h, uintptr(a2.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(a2)
	return newownedQBitArray(__r)
}

func (this *QBitArray) OperatorBitwiseOr(a2 *QBitArray) *QBitArray {
	__r, _, _ := purego.SyscallN(addr_QBitArray_OperatorBitwiseOr, this.h, uintptr(a2.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(a2)
	return newownedQBitArray(__r)
}

func (this *QBitArray) OperatorBitwiseXor(a2 *QBitArray) *QBitArray {
	__r, _, _ := purego.SyscallN(addr_QBitArray_OperatorBitwiseXor, this.h, uintptr(a2.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(a2)
	return newownedQBitArray(__r)
}

func (this *QBitArray) OperatorBitwiseNot() *QBitArray {
	__r, _, _ := purego.SyscallN(addr_QBitArray_OperatorBitwiseNot, this.h)
	runtime.KeepAlive(this)
	return newownedQBitArray(__r)
}

func (this *QBitArray) OperatorEqual(rhs *QBitArray) bool {
	__r0, _, _ := purego.SyscallN(addr_QBitArray_OperatorEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QBitArray) OperatorNotEqual(rhs *QBitArray) bool {
	__r0, _, _ := purego.SyscallN(addr_QBitArray_OperatorNotEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QBitArray) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QBitArray_delete(this.h)
	}
	this.UnsafeRelease()
}
