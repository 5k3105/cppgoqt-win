package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QVersionNumber_new uintptr
	addr_QVersionNumber_new2 uintptr
	addr_QVersionNumber_new3 uintptr
	addr_QVersionNumber_new4 uintptr
	addr_QVersionNumber_new5 uintptr
	addr_QVersionNumber_isNull uintptr
	addr_QVersionNumber_isNormalized uintptr
	addr_QVersionNumber_majorVersion uintptr
	addr_QVersionNumber_minorVersion uintptr
	addr_QVersionNumber_microVersion uintptr
	addr_QVersionNumber_normalized uintptr
	addr_QVersionNumber_segments uintptr
	addr_QVersionNumber_segmentAt uintptr
	addr_QVersionNumber_segmentCount uintptr
	addr_QVersionNumber_isPrefixOf uintptr
	addr_QVersionNumber_compare uintptr
	addr_QVersionNumber_commonPrefix uintptr
	addr_QVersionNumber_toString uintptr
	addr_QVersionNumber_fromString uintptr
	addr_QVersionNumber_OperatorAssign uintptr
	addr_QVersionNumber_fromString2 uintptr
	addr_QVersionNumber_OperatorEqual uintptr
	addr_QVersionNumber_OperatorNotEqual uintptr
	addr_QVersionNumber_OperatorLesser uintptr
	addr_QVersionNumber_OperatorGreater uintptr
	addr_QVersionNumber_OperatorLesserOrEqual uintptr
	addr_QVersionNumber_OperatorGreaterOrEqual uintptr
	pg_QVersionNumber_delete func(uintptr)
)

func init() {
	symaddr(&addr_QVersionNumber_new, "QVersionNumber_new")
	symaddr(&addr_QVersionNumber_new2, "QVersionNumber_new2")
	symaddr(&addr_QVersionNumber_new3, "QVersionNumber_new3")
	symaddr(&addr_QVersionNumber_new4, "QVersionNumber_new4")
	symaddr(&addr_QVersionNumber_new5, "QVersionNumber_new5")
	symaddr(&addr_QVersionNumber_isNull, "QVersionNumber_isNull")
	symaddr(&addr_QVersionNumber_isNormalized, "QVersionNumber_isNormalized")
	symaddr(&addr_QVersionNumber_majorVersion, "QVersionNumber_majorVersion")
	symaddr(&addr_QVersionNumber_minorVersion, "QVersionNumber_minorVersion")
	symaddr(&addr_QVersionNumber_microVersion, "QVersionNumber_microVersion")
	symaddr(&addr_QVersionNumber_normalized, "QVersionNumber_normalized")
	symaddr(&addr_QVersionNumber_segments, "QVersionNumber_segments")
	symaddr(&addr_QVersionNumber_segmentAt, "QVersionNumber_segmentAt")
	symaddr(&addr_QVersionNumber_segmentCount, "QVersionNumber_segmentCount")
	symaddr(&addr_QVersionNumber_isPrefixOf, "QVersionNumber_isPrefixOf")
	symaddr(&addr_QVersionNumber_compare, "QVersionNumber_compare")
	symaddr(&addr_QVersionNumber_commonPrefix, "QVersionNumber_commonPrefix")
	symaddr(&addr_QVersionNumber_toString, "QVersionNumber_toString")
	symaddr(&addr_QVersionNumber_fromString, "QVersionNumber_fromString")
	symaddr(&addr_QVersionNumber_OperatorAssign, "QVersionNumber_OperatorAssign")
	symaddr(&addr_QVersionNumber_fromString2, "QVersionNumber_fromString2")
	symaddr(&addr_QVersionNumber_OperatorEqual, "QVersionNumber_OperatorEqual")
	symaddr(&addr_QVersionNumber_OperatorNotEqual, "QVersionNumber_OperatorNotEqual")
	symaddr(&addr_QVersionNumber_OperatorLesser, "QVersionNumber_OperatorLesser")
	symaddr(&addr_QVersionNumber_OperatorGreater, "QVersionNumber_OperatorGreater")
	symaddr(&addr_QVersionNumber_OperatorLesserOrEqual, "QVersionNumber_OperatorLesserOrEqual")
	symaddr(&addr_QVersionNumber_OperatorGreaterOrEqual, "QVersionNumber_OperatorGreaterOrEqual")
	reg(&pg_QVersionNumber_delete, "QVersionNumber_delete")
}

type QVersionNumber struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QVersionNumber) UnsafeGoID() int64     { return this.vid }
func (this *QVersionNumber) UnsafeSetGoID(id int64) { this.vid = id }

func bindQVersionNumber(h uintptr) *QVersionNumber {
	if h == 0 {
		return nil
	}
	return &QVersionNumber{h: h}
}

func newownedQVersionNumber(h uintptr) *QVersionNumber {
	o := bindQVersionNumber(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QVersionNumber).Delete)
	}
	return o
}

func (this *QVersionNumber) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QVersionNumber) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQVersionNumber(h unsafe.Pointer) *QVersionNumber { return bindQVersionNumber(uintptr(h)) }
func UnsafeNewQVersionNumberOwned(h unsafe.Pointer) *QVersionNumber { return newownedQVersionNumber(uintptr(h)) }

func (this *QVersionNumber) UnsafeRelease() { this.h = 0 }

func NewQVersionNumber() *QVersionNumber {
	__r, _, _ := purego.SyscallN(addr_QVersionNumber_new)
	__o := newownedQVersionNumber(__r)
	return __o
}

func NewQVersionNumber2(maj int) *QVersionNumber {
	__r, _, _ := purego.SyscallN(addr_QVersionNumber_new2, uintptr(maj))
	__o := newownedQVersionNumber(__r)
	return __o
}

func NewQVersionNumber3(maj int, min int) *QVersionNumber {
	__r, _, _ := purego.SyscallN(addr_QVersionNumber_new3, uintptr(maj), uintptr(min))
	__o := newownedQVersionNumber(__r)
	return __o
}

func NewQVersionNumber4(maj int, min int, mic int) *QVersionNumber {
	__r, _, _ := purego.SyscallN(addr_QVersionNumber_new4, uintptr(maj), uintptr(min), uintptr(mic))
	__o := newownedQVersionNumber(__r)
	return __o
}

func NewQVersionNumber5(arg1 *QVersionNumber) *QVersionNumber {
	__r, _, _ := purego.SyscallN(addr_QVersionNumber_new5, uintptr(arg1.UnsafePointer()))
	__o := newownedQVersionNumber(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QVersionNumber) IsNull() bool {
	__r0, _, _ := purego.SyscallN(addr_QVersionNumber_isNull, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QVersionNumber) IsNormalized() bool {
	__r0, _, _ := purego.SyscallN(addr_QVersionNumber_isNormalized, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QVersionNumber) MajorVersion() int {
	__r0, _, _ := purego.SyscallN(addr_QVersionNumber_majorVersion, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QVersionNumber) MinorVersion() int {
	__r0, _, _ := purego.SyscallN(addr_QVersionNumber_minorVersion, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QVersionNumber) MicroVersion() int {
	__r0, _, _ := purego.SyscallN(addr_QVersionNumber_microVersion, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QVersionNumber) Normalized() *QVersionNumber {
	__r, _, _ := purego.SyscallN(addr_QVersionNumber_normalized, this.h)
	runtime.KeepAlive(this)
	return newownedQVersionNumber(__r)
}

func (this *QVersionNumber) Segments() []int {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QVersionNumber_segments, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]int, __n)
	__sl := unsafe.Slice((*int32)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = int(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QVersionNumber) SegmentAt(index int64) int {
	__r0, _, _ := purego.SyscallN(addr_QVersionNumber_segmentAt, this.h, uintptr(index))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QVersionNumber) SegmentCount() int64 {
	__r0, _, _ := purego.SyscallN(addr_QVersionNumber_segmentCount, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QVersionNumber) IsPrefixOf(other *QVersionNumber) bool {
	__r0, _, _ := purego.SyscallN(addr_QVersionNumber_isPrefixOf, this.h, uintptr(other.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return __r != 0
}

func QVersionNumber_Compare(v1 *QVersionNumber, v2 *QVersionNumber) int {
	__r0, _, _ := purego.SyscallN(addr_QVersionNumber_compare, uintptr(v1.UnsafePointer()), uintptr(v2.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(v1)
	runtime.KeepAlive(v2)
	return int(__r)
}

func QVersionNumber_CommonPrefix(v1 *QVersionNumber, v2 *QVersionNumber) *QVersionNumber {
	__r, _, _ := purego.SyscallN(addr_QVersionNumber_commonPrefix, uintptr(v1.UnsafePointer()), uintptr(v2.UnsafePointer()))
	runtime.KeepAlive(v1)
	runtime.KeepAlive(v2)
	return newownedQVersionNumber(__r)
}

func (this *QVersionNumber) ToString() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QVersionNumber_toString, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func QVersionNumber_FromString(string_ string) *QVersionNumber {
	var string__p uintptr
	if len(string_) > 0 {
		string__p = uintptr(unsafe.Pointer(unsafe.StringData(string_)))
	}
	__r, _, _ := purego.SyscallN(addr_QVersionNumber_fromString, string__p, uintptr(len(string_)))
	runtime.KeepAlive(string_)
	return newownedQVersionNumber(__r)
}

func (this *QVersionNumber) OperatorAssign(arg1 *QVersionNumber) *QVersionNumber {
	__r, _, _ := purego.SyscallN(addr_QVersionNumber_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQVersionNumber(__r)
}

func QVersionNumber_FromString2(string_ string) (*QVersionNumber, int64) {
	var string__p uintptr
	if len(string_) > 0 {
		string__p = uintptr(unsafe.Pointer(unsafe.StringData(string_)))
	}
	var suffixIndex_out int64
	__r, _, _ := purego.SyscallN(addr_QVersionNumber_fromString2, string__p, uintptr(len(string_)), uintptr(unsafe.Pointer(&suffixIndex_out)))
	runtime.KeepAlive(string_)
	return newownedQVersionNumber(__r), suffixIndex_out
}

func (this *QVersionNumber) OperatorEqual(rhs *QVersionNumber) bool {
	__r0, _, _ := purego.SyscallN(addr_QVersionNumber_OperatorEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QVersionNumber) OperatorNotEqual(rhs *QVersionNumber) bool {
	__r0, _, _ := purego.SyscallN(addr_QVersionNumber_OperatorNotEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QVersionNumber) OperatorLesser(rhs *QVersionNumber) bool {
	__r0, _, _ := purego.SyscallN(addr_QVersionNumber_OperatorLesser, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QVersionNumber) OperatorGreater(rhs *QVersionNumber) bool {
	__r0, _, _ := purego.SyscallN(addr_QVersionNumber_OperatorGreater, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QVersionNumber) OperatorLesserOrEqual(rhs *QVersionNumber) bool {
	__r0, _, _ := purego.SyscallN(addr_QVersionNumber_OperatorLesserOrEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QVersionNumber) OperatorGreaterOrEqual(rhs *QVersionNumber) bool {
	__r0, _, _ := purego.SyscallN(addr_QVersionNumber_OperatorGreaterOrEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QVersionNumber) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QVersionNumber_delete(this.h)
	}
	this.UnsafeRelease()
}
