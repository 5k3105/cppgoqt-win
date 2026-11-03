package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QAtomicInt_new uintptr
	addr_QAtomicInt_new2 uintptr
	addr_QAtomicInt_new3 uintptr
	addr_QAtomicInt_OperatorAssign uintptr
	addr_QAtomicInt_loadRelaxed uintptr
	addr_QAtomicInt_storeRelaxed uintptr
	addr_QAtomicInt_loadAcquire uintptr
	addr_QAtomicInt_storeRelease uintptr
	addr_QAtomicInt_ToInt uintptr
	addr_QAtomicInt_isReferenceCountingNative uintptr
	addr_QAtomicInt_isReferenceCountingWaitFree uintptr
	addr_QAtomicInt_ref uintptr
	addr_QAtomicInt_refRelaxed uintptr
	addr_QAtomicInt_deref uintptr
	addr_QAtomicInt_isTestAndSetNative uintptr
	addr_QAtomicInt_isTestAndSetWaitFree uintptr
	addr_QAtomicInt_testAndSetRelaxed uintptr
	addr_QAtomicInt_testAndSetAcquire uintptr
	addr_QAtomicInt_testAndSetRelease uintptr
	addr_QAtomicInt_testAndSetOrdered uintptr
	addr_QAtomicInt_testAndSetRelaxed2 uintptr
	addr_QAtomicInt_testAndSetAcquire2 uintptr
	addr_QAtomicInt_testAndSetRelease2 uintptr
	addr_QAtomicInt_testAndSetOrdered2 uintptr
	addr_QAtomicInt_isFetchAndStoreNative uintptr
	addr_QAtomicInt_isFetchAndStoreWaitFree uintptr
	addr_QAtomicInt_fetchAndStoreRelaxed uintptr
	addr_QAtomicInt_fetchAndStoreAcquire uintptr
	addr_QAtomicInt_fetchAndStoreRelease uintptr
	addr_QAtomicInt_fetchAndStoreOrdered uintptr
	addr_QAtomicInt_isFetchAndAddNative uintptr
	addr_QAtomicInt_isFetchAndAddWaitFree uintptr
	addr_QAtomicInt_fetchAndAddRelaxed uintptr
	addr_QAtomicInt_fetchAndAddAcquire uintptr
	addr_QAtomicInt_fetchAndAddRelease uintptr
	addr_QAtomicInt_fetchAndAddOrdered uintptr
	addr_QAtomicInt_fetchAndSubRelaxed uintptr
	addr_QAtomicInt_fetchAndSubAcquire uintptr
	addr_QAtomicInt_fetchAndSubRelease uintptr
	addr_QAtomicInt_fetchAndSubOrdered uintptr
	addr_QAtomicInt_fetchAndAndRelaxed uintptr
	addr_QAtomicInt_fetchAndAndAcquire uintptr
	addr_QAtomicInt_fetchAndAndRelease uintptr
	addr_QAtomicInt_fetchAndAndOrdered uintptr
	addr_QAtomicInt_fetchAndOrRelaxed uintptr
	addr_QAtomicInt_fetchAndOrAcquire uintptr
	addr_QAtomicInt_fetchAndOrRelease uintptr
	addr_QAtomicInt_fetchAndOrOrdered uintptr
	addr_QAtomicInt_fetchAndXorRelaxed uintptr
	addr_QAtomicInt_fetchAndXorAcquire uintptr
	addr_QAtomicInt_fetchAndXorRelease uintptr
	addr_QAtomicInt_fetchAndXorOrdered uintptr
	addr_QAtomicInt_OperatorPlusPlus uintptr
	addr_QAtomicInt_OperatorPlusPlus2 uintptr
	addr_QAtomicInt_OperatorMinusMinus uintptr
	addr_QAtomicInt_OperatorMinusMinus2 uintptr
	addr_QAtomicInt_OperatorPlusAssign uintptr
	addr_QAtomicInt_OperatorMinusAssign uintptr
	addr_QAtomicInt_OperatorBitwiseAndAssign uintptr
	addr_QAtomicInt_OperatorBitwiseOrAssign uintptr
	addr_QAtomicInt_OperatorBitwiseXorAssign uintptr
	pg_QAtomicInt_delete func(uintptr)
)

func init() {
	symaddr(&addr_QAtomicInt_new, "QAtomicInt_new")
	symaddr(&addr_QAtomicInt_new2, "QAtomicInt_new2")
	symaddr(&addr_QAtomicInt_new3, "QAtomicInt_new3")
	symaddr(&addr_QAtomicInt_OperatorAssign, "QAtomicInt_OperatorAssign")
	symaddr(&addr_QAtomicInt_loadRelaxed, "QAtomicInt_loadRelaxed")
	symaddr(&addr_QAtomicInt_storeRelaxed, "QAtomicInt_storeRelaxed")
	symaddr(&addr_QAtomicInt_loadAcquire, "QAtomicInt_loadAcquire")
	symaddr(&addr_QAtomicInt_storeRelease, "QAtomicInt_storeRelease")
	symaddr(&addr_QAtomicInt_ToInt, "QAtomicInt_ToInt")
	symaddr(&addr_QAtomicInt_isReferenceCountingNative, "QAtomicInt_isReferenceCountingNative")
	symaddr(&addr_QAtomicInt_isReferenceCountingWaitFree, "QAtomicInt_isReferenceCountingWaitFree")
	symaddr(&addr_QAtomicInt_ref, "QAtomicInt_ref")
	symaddr(&addr_QAtomicInt_refRelaxed, "QAtomicInt_refRelaxed")
	symaddr(&addr_QAtomicInt_deref, "QAtomicInt_deref")
	symaddr(&addr_QAtomicInt_isTestAndSetNative, "QAtomicInt_isTestAndSetNative")
	symaddr(&addr_QAtomicInt_isTestAndSetWaitFree, "QAtomicInt_isTestAndSetWaitFree")
	symaddr(&addr_QAtomicInt_testAndSetRelaxed, "QAtomicInt_testAndSetRelaxed")
	symaddr(&addr_QAtomicInt_testAndSetAcquire, "QAtomicInt_testAndSetAcquire")
	symaddr(&addr_QAtomicInt_testAndSetRelease, "QAtomicInt_testAndSetRelease")
	symaddr(&addr_QAtomicInt_testAndSetOrdered, "QAtomicInt_testAndSetOrdered")
	symaddr(&addr_QAtomicInt_testAndSetRelaxed2, "QAtomicInt_testAndSetRelaxed2")
	symaddr(&addr_QAtomicInt_testAndSetAcquire2, "QAtomicInt_testAndSetAcquire2")
	symaddr(&addr_QAtomicInt_testAndSetRelease2, "QAtomicInt_testAndSetRelease2")
	symaddr(&addr_QAtomicInt_testAndSetOrdered2, "QAtomicInt_testAndSetOrdered2")
	symaddr(&addr_QAtomicInt_isFetchAndStoreNative, "QAtomicInt_isFetchAndStoreNative")
	symaddr(&addr_QAtomicInt_isFetchAndStoreWaitFree, "QAtomicInt_isFetchAndStoreWaitFree")
	symaddr(&addr_QAtomicInt_fetchAndStoreRelaxed, "QAtomicInt_fetchAndStoreRelaxed")
	symaddr(&addr_QAtomicInt_fetchAndStoreAcquire, "QAtomicInt_fetchAndStoreAcquire")
	symaddr(&addr_QAtomicInt_fetchAndStoreRelease, "QAtomicInt_fetchAndStoreRelease")
	symaddr(&addr_QAtomicInt_fetchAndStoreOrdered, "QAtomicInt_fetchAndStoreOrdered")
	symaddr(&addr_QAtomicInt_isFetchAndAddNative, "QAtomicInt_isFetchAndAddNative")
	symaddr(&addr_QAtomicInt_isFetchAndAddWaitFree, "QAtomicInt_isFetchAndAddWaitFree")
	symaddr(&addr_QAtomicInt_fetchAndAddRelaxed, "QAtomicInt_fetchAndAddRelaxed")
	symaddr(&addr_QAtomicInt_fetchAndAddAcquire, "QAtomicInt_fetchAndAddAcquire")
	symaddr(&addr_QAtomicInt_fetchAndAddRelease, "QAtomicInt_fetchAndAddRelease")
	symaddr(&addr_QAtomicInt_fetchAndAddOrdered, "QAtomicInt_fetchAndAddOrdered")
	symaddr(&addr_QAtomicInt_fetchAndSubRelaxed, "QAtomicInt_fetchAndSubRelaxed")
	symaddr(&addr_QAtomicInt_fetchAndSubAcquire, "QAtomicInt_fetchAndSubAcquire")
	symaddr(&addr_QAtomicInt_fetchAndSubRelease, "QAtomicInt_fetchAndSubRelease")
	symaddr(&addr_QAtomicInt_fetchAndSubOrdered, "QAtomicInt_fetchAndSubOrdered")
	symaddr(&addr_QAtomicInt_fetchAndAndRelaxed, "QAtomicInt_fetchAndAndRelaxed")
	symaddr(&addr_QAtomicInt_fetchAndAndAcquire, "QAtomicInt_fetchAndAndAcquire")
	symaddr(&addr_QAtomicInt_fetchAndAndRelease, "QAtomicInt_fetchAndAndRelease")
	symaddr(&addr_QAtomicInt_fetchAndAndOrdered, "QAtomicInt_fetchAndAndOrdered")
	symaddr(&addr_QAtomicInt_fetchAndOrRelaxed, "QAtomicInt_fetchAndOrRelaxed")
	symaddr(&addr_QAtomicInt_fetchAndOrAcquire, "QAtomicInt_fetchAndOrAcquire")
	symaddr(&addr_QAtomicInt_fetchAndOrRelease, "QAtomicInt_fetchAndOrRelease")
	symaddr(&addr_QAtomicInt_fetchAndOrOrdered, "QAtomicInt_fetchAndOrOrdered")
	symaddr(&addr_QAtomicInt_fetchAndXorRelaxed, "QAtomicInt_fetchAndXorRelaxed")
	symaddr(&addr_QAtomicInt_fetchAndXorAcquire, "QAtomicInt_fetchAndXorAcquire")
	symaddr(&addr_QAtomicInt_fetchAndXorRelease, "QAtomicInt_fetchAndXorRelease")
	symaddr(&addr_QAtomicInt_fetchAndXorOrdered, "QAtomicInt_fetchAndXorOrdered")
	symaddr(&addr_QAtomicInt_OperatorPlusPlus, "QAtomicInt_OperatorPlusPlus")
	symaddr(&addr_QAtomicInt_OperatorPlusPlus2, "QAtomicInt_OperatorPlusPlus2")
	symaddr(&addr_QAtomicInt_OperatorMinusMinus, "QAtomicInt_OperatorMinusMinus")
	symaddr(&addr_QAtomicInt_OperatorMinusMinus2, "QAtomicInt_OperatorMinusMinus2")
	symaddr(&addr_QAtomicInt_OperatorPlusAssign, "QAtomicInt_OperatorPlusAssign")
	symaddr(&addr_QAtomicInt_OperatorMinusAssign, "QAtomicInt_OperatorMinusAssign")
	symaddr(&addr_QAtomicInt_OperatorBitwiseAndAssign, "QAtomicInt_OperatorBitwiseAndAssign")
	symaddr(&addr_QAtomicInt_OperatorBitwiseOrAssign, "QAtomicInt_OperatorBitwiseOrAssign")
	symaddr(&addr_QAtomicInt_OperatorBitwiseXorAssign, "QAtomicInt_OperatorBitwiseXorAssign")
	reg(&pg_QAtomicInt_delete, "QAtomicInt_delete")
}

type QAtomicInt struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QAtomicInt) UnsafeGoID() int64     { return this.vid }
func (this *QAtomicInt) UnsafeSetGoID(id int64) { this.vid = id }

func bindQAtomicInt(h uintptr) *QAtomicInt {
	if h == 0 {
		return nil
	}
	return &QAtomicInt{h: h}
}

func newownedQAtomicInt(h uintptr) *QAtomicInt {
	o := bindQAtomicInt(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QAtomicInt).Delete)
	}
	return o
}

func (this *QAtomicInt) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QAtomicInt) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQAtomicInt(h unsafe.Pointer) *QAtomicInt { return bindQAtomicInt(uintptr(h)) }
func UnsafeNewQAtomicIntOwned(h unsafe.Pointer) *QAtomicInt { return newownedQAtomicInt(uintptr(h)) }

func (this *QAtomicInt) UnsafeRelease() { this.h = 0 }

func NewQAtomicInt() *QAtomicInt {
	__r, _, _ := purego.SyscallN(addr_QAtomicInt_new)
	__o := newownedQAtomicInt(__r)
	return __o
}

func NewQAtomicInt2(arg1 *QAtomicInt) *QAtomicInt {
	__r, _, _ := purego.SyscallN(addr_QAtomicInt_new2, uintptr(arg1.UnsafePointer()))
	__o := newownedQAtomicInt(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func NewQAtomicInt3(value int) *QAtomicInt {
	__r, _, _ := purego.SyscallN(addr_QAtomicInt_new3, uintptr(value))
	__o := newownedQAtomicInt(__r)
	return __o
}

func (this *QAtomicInt) OperatorAssign(arg1 *QAtomicInt) *QAtomicInt {
	__r, _, _ := purego.SyscallN(addr_QAtomicInt_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQAtomicInt(__r)
}

func (this *QAtomicInt) LoadRelaxed() int {
	__r0, _, _ := purego.SyscallN(addr_QAtomicInt_loadRelaxed, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAtomicInt) StoreRelaxed(newValue int) {
	purego.SyscallN(addr_QAtomicInt_storeRelaxed, this.h, uintptr(newValue))
	runtime.KeepAlive(this)
}

func (this *QAtomicInt) LoadAcquire() int {
	__r0, _, _ := purego.SyscallN(addr_QAtomicInt_loadAcquire, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAtomicInt) StoreRelease(newValue int) {
	purego.SyscallN(addr_QAtomicInt_storeRelease, this.h, uintptr(newValue))
	runtime.KeepAlive(this)
}

func (this *QAtomicInt) ToInt() int {
	__r0, _, _ := purego.SyscallN(addr_QAtomicInt_ToInt, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func QAtomicInt_IsReferenceCountingNative() bool {
	__r0, _, _ := purego.SyscallN(addr_QAtomicInt_isReferenceCountingNative)
	__r := int32(uint32(__r0))
	return __r != 0
}

func QAtomicInt_IsReferenceCountingWaitFree() bool {
	__r0, _, _ := purego.SyscallN(addr_QAtomicInt_isReferenceCountingWaitFree)
	__r := int32(uint32(__r0))
	return __r != 0
}

func (this *QAtomicInt) Ref() bool {
	__r0, _, _ := purego.SyscallN(addr_QAtomicInt_ref, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QAtomicInt) RefRelaxed() {
	purego.SyscallN(addr_QAtomicInt_refRelaxed, this.h)
	runtime.KeepAlive(this)
}

func (this *QAtomicInt) Deref() bool {
	__r0, _, _ := purego.SyscallN(addr_QAtomicInt_deref, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QAtomicInt_IsTestAndSetNative() bool {
	__r0, _, _ := purego.SyscallN(addr_QAtomicInt_isTestAndSetNative)
	__r := int32(uint32(__r0))
	return __r != 0
}

func QAtomicInt_IsTestAndSetWaitFree() bool {
	__r0, _, _ := purego.SyscallN(addr_QAtomicInt_isTestAndSetWaitFree)
	__r := int32(uint32(__r0))
	return __r != 0
}

func (this *QAtomicInt) TestAndSetRelaxed(expectedValue int, newValue int) bool {
	__r0, _, _ := purego.SyscallN(addr_QAtomicInt_testAndSetRelaxed, this.h, uintptr(expectedValue), uintptr(newValue))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QAtomicInt) TestAndSetAcquire(expectedValue int, newValue int) bool {
	__r0, _, _ := purego.SyscallN(addr_QAtomicInt_testAndSetAcquire, this.h, uintptr(expectedValue), uintptr(newValue))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QAtomicInt) TestAndSetRelease(expectedValue int, newValue int) bool {
	__r0, _, _ := purego.SyscallN(addr_QAtomicInt_testAndSetRelease, this.h, uintptr(expectedValue), uintptr(newValue))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QAtomicInt) TestAndSetOrdered(expectedValue int, newValue int) bool {
	__r0, _, _ := purego.SyscallN(addr_QAtomicInt_testAndSetOrdered, this.h, uintptr(expectedValue), uintptr(newValue))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QAtomicInt) TestAndSetRelaxed2(expectedValue int, newValue int, currentValue int) (bool, int) {
	currentValue_io := int32(currentValue)
	__r0, _, _ := purego.SyscallN(addr_QAtomicInt_testAndSetRelaxed2, this.h, uintptr(expectedValue), uintptr(newValue), uintptr(unsafe.Pointer(&currentValue_io)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0, int(currentValue_io)
}

func (this *QAtomicInt) TestAndSetAcquire2(expectedValue int, newValue int, currentValue int) (bool, int) {
	currentValue_io := int32(currentValue)
	__r0, _, _ := purego.SyscallN(addr_QAtomicInt_testAndSetAcquire2, this.h, uintptr(expectedValue), uintptr(newValue), uintptr(unsafe.Pointer(&currentValue_io)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0, int(currentValue_io)
}

func (this *QAtomicInt) TestAndSetRelease2(expectedValue int, newValue int, currentValue int) (bool, int) {
	currentValue_io := int32(currentValue)
	__r0, _, _ := purego.SyscallN(addr_QAtomicInt_testAndSetRelease2, this.h, uintptr(expectedValue), uintptr(newValue), uintptr(unsafe.Pointer(&currentValue_io)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0, int(currentValue_io)
}

func (this *QAtomicInt) TestAndSetOrdered2(expectedValue int, newValue int, currentValue int) (bool, int) {
	currentValue_io := int32(currentValue)
	__r0, _, _ := purego.SyscallN(addr_QAtomicInt_testAndSetOrdered2, this.h, uintptr(expectedValue), uintptr(newValue), uintptr(unsafe.Pointer(&currentValue_io)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0, int(currentValue_io)
}

func QAtomicInt_IsFetchAndStoreNative() bool {
	__r0, _, _ := purego.SyscallN(addr_QAtomicInt_isFetchAndStoreNative)
	__r := int32(uint32(__r0))
	return __r != 0
}

func QAtomicInt_IsFetchAndStoreWaitFree() bool {
	__r0, _, _ := purego.SyscallN(addr_QAtomicInt_isFetchAndStoreWaitFree)
	__r := int32(uint32(__r0))
	return __r != 0
}

func (this *QAtomicInt) FetchAndStoreRelaxed(newValue int) int {
	__r0, _, _ := purego.SyscallN(addr_QAtomicInt_fetchAndStoreRelaxed, this.h, uintptr(newValue))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAtomicInt) FetchAndStoreAcquire(newValue int) int {
	__r0, _, _ := purego.SyscallN(addr_QAtomicInt_fetchAndStoreAcquire, this.h, uintptr(newValue))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAtomicInt) FetchAndStoreRelease(newValue int) int {
	__r0, _, _ := purego.SyscallN(addr_QAtomicInt_fetchAndStoreRelease, this.h, uintptr(newValue))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAtomicInt) FetchAndStoreOrdered(newValue int) int {
	__r0, _, _ := purego.SyscallN(addr_QAtomicInt_fetchAndStoreOrdered, this.h, uintptr(newValue))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func QAtomicInt_IsFetchAndAddNative() bool {
	__r0, _, _ := purego.SyscallN(addr_QAtomicInt_isFetchAndAddNative)
	__r := int32(uint32(__r0))
	return __r != 0
}

func QAtomicInt_IsFetchAndAddWaitFree() bool {
	__r0, _, _ := purego.SyscallN(addr_QAtomicInt_isFetchAndAddWaitFree)
	__r := int32(uint32(__r0))
	return __r != 0
}

func (this *QAtomicInt) FetchAndAddRelaxed(valueToAdd int) int {
	__r0, _, _ := purego.SyscallN(addr_QAtomicInt_fetchAndAddRelaxed, this.h, uintptr(valueToAdd))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAtomicInt) FetchAndAddAcquire(valueToAdd int) int {
	__r0, _, _ := purego.SyscallN(addr_QAtomicInt_fetchAndAddAcquire, this.h, uintptr(valueToAdd))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAtomicInt) FetchAndAddRelease(valueToAdd int) int {
	__r0, _, _ := purego.SyscallN(addr_QAtomicInt_fetchAndAddRelease, this.h, uintptr(valueToAdd))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAtomicInt) FetchAndAddOrdered(valueToAdd int) int {
	__r0, _, _ := purego.SyscallN(addr_QAtomicInt_fetchAndAddOrdered, this.h, uintptr(valueToAdd))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAtomicInt) FetchAndSubRelaxed(valueToAdd int) int {
	__r0, _, _ := purego.SyscallN(addr_QAtomicInt_fetchAndSubRelaxed, this.h, uintptr(valueToAdd))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAtomicInt) FetchAndSubAcquire(valueToAdd int) int {
	__r0, _, _ := purego.SyscallN(addr_QAtomicInt_fetchAndSubAcquire, this.h, uintptr(valueToAdd))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAtomicInt) FetchAndSubRelease(valueToAdd int) int {
	__r0, _, _ := purego.SyscallN(addr_QAtomicInt_fetchAndSubRelease, this.h, uintptr(valueToAdd))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAtomicInt) FetchAndSubOrdered(valueToAdd int) int {
	__r0, _, _ := purego.SyscallN(addr_QAtomicInt_fetchAndSubOrdered, this.h, uintptr(valueToAdd))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAtomicInt) FetchAndAndRelaxed(valueToAdd int) int {
	__r0, _, _ := purego.SyscallN(addr_QAtomicInt_fetchAndAndRelaxed, this.h, uintptr(valueToAdd))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAtomicInt) FetchAndAndAcquire(valueToAdd int) int {
	__r0, _, _ := purego.SyscallN(addr_QAtomicInt_fetchAndAndAcquire, this.h, uintptr(valueToAdd))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAtomicInt) FetchAndAndRelease(valueToAdd int) int {
	__r0, _, _ := purego.SyscallN(addr_QAtomicInt_fetchAndAndRelease, this.h, uintptr(valueToAdd))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAtomicInt) FetchAndAndOrdered(valueToAdd int) int {
	__r0, _, _ := purego.SyscallN(addr_QAtomicInt_fetchAndAndOrdered, this.h, uintptr(valueToAdd))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAtomicInt) FetchAndOrRelaxed(valueToAdd int) int {
	__r0, _, _ := purego.SyscallN(addr_QAtomicInt_fetchAndOrRelaxed, this.h, uintptr(valueToAdd))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAtomicInt) FetchAndOrAcquire(valueToAdd int) int {
	__r0, _, _ := purego.SyscallN(addr_QAtomicInt_fetchAndOrAcquire, this.h, uintptr(valueToAdd))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAtomicInt) FetchAndOrRelease(valueToAdd int) int {
	__r0, _, _ := purego.SyscallN(addr_QAtomicInt_fetchAndOrRelease, this.h, uintptr(valueToAdd))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAtomicInt) FetchAndOrOrdered(valueToAdd int) int {
	__r0, _, _ := purego.SyscallN(addr_QAtomicInt_fetchAndOrOrdered, this.h, uintptr(valueToAdd))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAtomicInt) FetchAndXorRelaxed(valueToAdd int) int {
	__r0, _, _ := purego.SyscallN(addr_QAtomicInt_fetchAndXorRelaxed, this.h, uintptr(valueToAdd))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAtomicInt) FetchAndXorAcquire(valueToAdd int) int {
	__r0, _, _ := purego.SyscallN(addr_QAtomicInt_fetchAndXorAcquire, this.h, uintptr(valueToAdd))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAtomicInt) FetchAndXorRelease(valueToAdd int) int {
	__r0, _, _ := purego.SyscallN(addr_QAtomicInt_fetchAndXorRelease, this.h, uintptr(valueToAdd))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAtomicInt) FetchAndXorOrdered(valueToAdd int) int {
	__r0, _, _ := purego.SyscallN(addr_QAtomicInt_fetchAndXorOrdered, this.h, uintptr(valueToAdd))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAtomicInt) OperatorPlusPlus() int {
	__r0, _, _ := purego.SyscallN(addr_QAtomicInt_OperatorPlusPlus, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAtomicInt) OperatorPlusPlusWithInt(arg1 int) int {
	__r0, _, _ := purego.SyscallN(addr_QAtomicInt_OperatorPlusPlus2, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAtomicInt) OperatorMinusMinus() int {
	__r0, _, _ := purego.SyscallN(addr_QAtomicInt_OperatorMinusMinus, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAtomicInt) OperatorMinusMinusWithInt(arg1 int) int {
	__r0, _, _ := purego.SyscallN(addr_QAtomicInt_OperatorMinusMinus2, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAtomicInt) OperatorPlusAssign(v int) int {
	__r0, _, _ := purego.SyscallN(addr_QAtomicInt_OperatorPlusAssign, this.h, uintptr(v))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAtomicInt) OperatorMinusAssign(v int) int {
	__r0, _, _ := purego.SyscallN(addr_QAtomicInt_OperatorMinusAssign, this.h, uintptr(v))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAtomicInt) OperatorBitwiseAndAssign(v int) int {
	__r0, _, _ := purego.SyscallN(addr_QAtomicInt_OperatorBitwiseAndAssign, this.h, uintptr(v))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAtomicInt) OperatorBitwiseOrAssign(v int) int {
	__r0, _, _ := purego.SyscallN(addr_QAtomicInt_OperatorBitwiseOrAssign, this.h, uintptr(v))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAtomicInt) OperatorBitwiseXorAssign(v int) int {
	__r0, _, _ := purego.SyscallN(addr_QAtomicInt_OperatorBitwiseXorAssign, this.h, uintptr(v))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAtomicInt) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QAtomicInt_delete(this.h)
	}
	this.UnsafeRelease()
}
