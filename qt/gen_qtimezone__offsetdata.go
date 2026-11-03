package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QTimeZone__OffsetData_new uintptr
	addr_QTimeZone__OffsetData_new2 uintptr
	addr_QTimeZone__OffsetData_OperatorAssign uintptr
	addr_QTimeZone__OffsetData_abbreviation_get uintptr
	addr_QTimeZone__OffsetData_abbreviation_set uintptr
	addr_QTimeZone__OffsetData_atUtc_get uintptr
	addr_QTimeZone__OffsetData_atUtc_set uintptr
	addr_QTimeZone__OffsetData_offsetFromUtc_get uintptr
	addr_QTimeZone__OffsetData_offsetFromUtc_set uintptr
	addr_QTimeZone__OffsetData_standardTimeOffset_get uintptr
	addr_QTimeZone__OffsetData_standardTimeOffset_set uintptr
	addr_QTimeZone__OffsetData_daylightTimeOffset_get uintptr
	addr_QTimeZone__OffsetData_daylightTimeOffset_set uintptr
	pg_QTimeZone__OffsetData_delete func(uintptr)
)

func init() {
	symaddr(&addr_QTimeZone__OffsetData_new, "QTimeZone__OffsetData_new")
	symaddr(&addr_QTimeZone__OffsetData_new2, "QTimeZone__OffsetData_new2")
	symaddr(&addr_QTimeZone__OffsetData_OperatorAssign, "QTimeZone__OffsetData_OperatorAssign")
	symaddr(&addr_QTimeZone__OffsetData_abbreviation_get, "QTimeZone__OffsetData_abbreviation_get")
	symaddr(&addr_QTimeZone__OffsetData_abbreviation_set, "QTimeZone__OffsetData_abbreviation_set")
	symaddr(&addr_QTimeZone__OffsetData_atUtc_get, "QTimeZone__OffsetData_atUtc_get")
	symaddr(&addr_QTimeZone__OffsetData_atUtc_set, "QTimeZone__OffsetData_atUtc_set")
	symaddr(&addr_QTimeZone__OffsetData_offsetFromUtc_get, "QTimeZone__OffsetData_offsetFromUtc_get")
	symaddr(&addr_QTimeZone__OffsetData_offsetFromUtc_set, "QTimeZone__OffsetData_offsetFromUtc_set")
	symaddr(&addr_QTimeZone__OffsetData_standardTimeOffset_get, "QTimeZone__OffsetData_standardTimeOffset_get")
	symaddr(&addr_QTimeZone__OffsetData_standardTimeOffset_set, "QTimeZone__OffsetData_standardTimeOffset_set")
	symaddr(&addr_QTimeZone__OffsetData_daylightTimeOffset_get, "QTimeZone__OffsetData_daylightTimeOffset_get")
	symaddr(&addr_QTimeZone__OffsetData_daylightTimeOffset_set, "QTimeZone__OffsetData_daylightTimeOffset_set")
	reg(&pg_QTimeZone__OffsetData_delete, "QTimeZone__OffsetData_delete")
}

type QTimeZone__OffsetData struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QTimeZone__OffsetData) UnsafeGoID() int64     { return this.vid }
func (this *QTimeZone__OffsetData) UnsafeSetGoID(id int64) { this.vid = id }

func bindQTimeZone__OffsetData(h uintptr) *QTimeZone__OffsetData {
	if h == 0 {
		return nil
	}
	return &QTimeZone__OffsetData{h: h}
}

func newownedQTimeZone__OffsetData(h uintptr) *QTimeZone__OffsetData {
	o := bindQTimeZone__OffsetData(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QTimeZone__OffsetData).Delete)
	}
	return o
}

func (this *QTimeZone__OffsetData) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QTimeZone__OffsetData) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQTimeZone__OffsetData(h unsafe.Pointer) *QTimeZone__OffsetData { return bindQTimeZone__OffsetData(uintptr(h)) }
func UnsafeNewQTimeZone__OffsetDataOwned(h unsafe.Pointer) *QTimeZone__OffsetData { return newownedQTimeZone__OffsetData(uintptr(h)) }

func (this *QTimeZone__OffsetData) UnsafeRelease() { this.h = 0 }

func NewQTimeZone__OffsetData(arg1 *QTimeZone__OffsetData) *QTimeZone__OffsetData {
	__r, _, _ := purego.SyscallN(addr_QTimeZone__OffsetData_new, uintptr(arg1.UnsafePointer()))
	__o := newownedQTimeZone__OffsetData(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func NewQTimeZone__OffsetData2() *QTimeZone__OffsetData {
	__r, _, _ := purego.SyscallN(addr_QTimeZone__OffsetData_new2)
	__o := newownedQTimeZone__OffsetData(__r)
	return __o
}

func (this *QTimeZone__OffsetData) OperatorAssign(arg1 *QTimeZone__OffsetData) *QTimeZone__OffsetData {
	__r, _, _ := purego.SyscallN(addr_QTimeZone__OffsetData_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQTimeZone__OffsetData(__r)
}

func (this *QTimeZone__OffsetData) Abbreviation() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTimeZone__OffsetData_abbreviation_get, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QTimeZone__OffsetData) SetAbbreviation(abbreviation string) {
	var abbreviation_p uintptr
	if len(abbreviation) > 0 {
		abbreviation_p = uintptr(unsafe.Pointer(unsafe.StringData(abbreviation)))
	}
	purego.SyscallN(addr_QTimeZone__OffsetData_abbreviation_set, this.h, abbreviation_p, uintptr(len(abbreviation)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(abbreviation)
}

func (this *QTimeZone__OffsetData) AtUtc() *QDateTime {
	__r, _, _ := purego.SyscallN(addr_QTimeZone__OffsetData_atUtc_get, this.h)
	runtime.KeepAlive(this)
	return newownedQDateTime(__r)
}

func (this *QTimeZone__OffsetData) SetAtUtc(atUtc *QDateTime) {
	purego.SyscallN(addr_QTimeZone__OffsetData_atUtc_set, this.h, uintptr(atUtc.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(atUtc)
}

func (this *QTimeZone__OffsetData) OffsetFromUtc() int {
	__r0, _, _ := purego.SyscallN(addr_QTimeZone__OffsetData_offsetFromUtc_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTimeZone__OffsetData) SetOffsetFromUtc(offsetFromUtc int) {
	purego.SyscallN(addr_QTimeZone__OffsetData_offsetFromUtc_set, this.h, uintptr(offsetFromUtc))
	runtime.KeepAlive(this)
}

func (this *QTimeZone__OffsetData) StandardTimeOffset() int {
	__r0, _, _ := purego.SyscallN(addr_QTimeZone__OffsetData_standardTimeOffset_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTimeZone__OffsetData) SetStandardTimeOffset(standardTimeOffset int) {
	purego.SyscallN(addr_QTimeZone__OffsetData_standardTimeOffset_set, this.h, uintptr(standardTimeOffset))
	runtime.KeepAlive(this)
}

func (this *QTimeZone__OffsetData) DaylightTimeOffset() int {
	__r0, _, _ := purego.SyscallN(addr_QTimeZone__OffsetData_daylightTimeOffset_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTimeZone__OffsetData) SetDaylightTimeOffset(daylightTimeOffset int) {
	purego.SyscallN(addr_QTimeZone__OffsetData_daylightTimeOffset_set, this.h, uintptr(daylightTimeOffset))
	runtime.KeepAlive(this)
}

func (this *QTimeZone__OffsetData) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QTimeZone__OffsetData_delete(this.h)
	}
	this.UnsafeRelease()
}
