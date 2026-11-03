package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QTimeZone_new uintptr
	addr_QTimeZone_new2 uintptr
	addr_QTimeZone_new3 uintptr
	addr_QTimeZone_new4 uintptr
	addr_QTimeZone_new5 uintptr
	addr_QTimeZone_new6 uintptr
	addr_QTimeZone_new7 uintptr
	addr_QTimeZone_new8 uintptr
	addr_QTimeZone_OperatorAssign uintptr
	addr_QTimeZone_swap uintptr
	addr_QTimeZone_isValid uintptr
	addr_QTimeZone_fromDurationAheadOfUtc uintptr
	addr_QTimeZone_fromSecondsAheadOfUtc uintptr
	addr_QTimeZone_timeSpec uintptr
	addr_QTimeZone_fixedSecondsAheadOfUtc uintptr
	addr_QTimeZone_isUtcOrFixedOffset uintptr
	addr_QTimeZone_isUtcOrFixedOffset2 uintptr
	addr_QTimeZone_asBackendZone uintptr
	addr_QTimeZone_hasAlternativeName uintptr
	addr_QTimeZone_id uintptr
	addr_QTimeZone_territory uintptr
	addr_QTimeZone_country uintptr
	addr_QTimeZone_comment uintptr
	addr_QTimeZone_displayName uintptr
	addr_QTimeZone_displayName2 uintptr
	addr_QTimeZone_abbreviation uintptr
	addr_QTimeZone_offsetFromUtc uintptr
	addr_QTimeZone_standardTimeOffset uintptr
	addr_QTimeZone_daylightTimeOffset uintptr
	addr_QTimeZone_hasDaylightTime uintptr
	addr_QTimeZone_isDaylightTime uintptr
	addr_QTimeZone_offsetData uintptr
	addr_QTimeZone_hasTransitions uintptr
	addr_QTimeZone_nextTransition uintptr
	addr_QTimeZone_previousTransition uintptr
	addr_QTimeZone_transitions uintptr
	addr_QTimeZone_systemTimeZoneId uintptr
	addr_QTimeZone_systemTimeZone uintptr
	addr_QTimeZone_utc uintptr
	addr_QTimeZone_isTimeZoneIdAvailable uintptr
	addr_QTimeZone_availableTimeZoneIds uintptr
	addr_QTimeZone_availableTimeZoneIds2 uintptr
	addr_QTimeZone_availableTimeZoneIds3 uintptr
	addr_QTimeZone_ianaIdToWindowsId uintptr
	addr_QTimeZone_windowsIdToDefaultIanaId uintptr
	addr_QTimeZone_windowsIdToDefaultIanaId2 uintptr
	addr_QTimeZone_windowsIdToIanaIds uintptr
	addr_QTimeZone_windowsIdToIanaIds2 uintptr
	addr_QTimeZone_displayName3 uintptr
	addr_QTimeZone_displayName4 uintptr
	addr_QTimeZone_displayName5 uintptr
	addr_QTimeZone_displayName6 uintptr
	addr_QTimeZone_OperatorEqual uintptr
	addr_QTimeZone_OperatorNotEqual uintptr
	pg_QTimeZone_delete func(uintptr)
)

func init() {
	symaddr(&addr_QTimeZone_new, "QTimeZone_new")
	symaddr(&addr_QTimeZone_new2, "QTimeZone_new2")
	symaddr(&addr_QTimeZone_new3, "QTimeZone_new3")
	symaddr(&addr_QTimeZone_new4, "QTimeZone_new4")
	symaddr(&addr_QTimeZone_new5, "QTimeZone_new5")
	symaddr(&addr_QTimeZone_new6, "QTimeZone_new6")
	symaddr(&addr_QTimeZone_new7, "QTimeZone_new7")
	symaddr(&addr_QTimeZone_new8, "QTimeZone_new8")
	symaddr(&addr_QTimeZone_OperatorAssign, "QTimeZone_OperatorAssign")
	symaddr(&addr_QTimeZone_swap, "QTimeZone_swap")
	symaddr(&addr_QTimeZone_isValid, "QTimeZone_isValid")
	symaddr(&addr_QTimeZone_fromDurationAheadOfUtc, "QTimeZone_fromDurationAheadOfUtc")
	symaddr(&addr_QTimeZone_fromSecondsAheadOfUtc, "QTimeZone_fromSecondsAheadOfUtc")
	symaddr(&addr_QTimeZone_timeSpec, "QTimeZone_timeSpec")
	symaddr(&addr_QTimeZone_fixedSecondsAheadOfUtc, "QTimeZone_fixedSecondsAheadOfUtc")
	symaddr(&addr_QTimeZone_isUtcOrFixedOffset, "QTimeZone_isUtcOrFixedOffset")
	symaddr(&addr_QTimeZone_isUtcOrFixedOffset2, "QTimeZone_isUtcOrFixedOffset2")
	symaddr(&addr_QTimeZone_asBackendZone, "QTimeZone_asBackendZone")
	symaddr(&addr_QTimeZone_hasAlternativeName, "QTimeZone_hasAlternativeName")
	symaddr(&addr_QTimeZone_id, "QTimeZone_id")
	symaddr(&addr_QTimeZone_territory, "QTimeZone_territory")
	symaddr(&addr_QTimeZone_country, "QTimeZone_country")
	symaddr(&addr_QTimeZone_comment, "QTimeZone_comment")
	symaddr(&addr_QTimeZone_displayName, "QTimeZone_displayName")
	symaddr(&addr_QTimeZone_displayName2, "QTimeZone_displayName2")
	symaddr(&addr_QTimeZone_abbreviation, "QTimeZone_abbreviation")
	symaddr(&addr_QTimeZone_offsetFromUtc, "QTimeZone_offsetFromUtc")
	symaddr(&addr_QTimeZone_standardTimeOffset, "QTimeZone_standardTimeOffset")
	symaddr(&addr_QTimeZone_daylightTimeOffset, "QTimeZone_daylightTimeOffset")
	symaddr(&addr_QTimeZone_hasDaylightTime, "QTimeZone_hasDaylightTime")
	symaddr(&addr_QTimeZone_isDaylightTime, "QTimeZone_isDaylightTime")
	symaddr(&addr_QTimeZone_offsetData, "QTimeZone_offsetData")
	symaddr(&addr_QTimeZone_hasTransitions, "QTimeZone_hasTransitions")
	symaddr(&addr_QTimeZone_nextTransition, "QTimeZone_nextTransition")
	symaddr(&addr_QTimeZone_previousTransition, "QTimeZone_previousTransition")
	symaddr(&addr_QTimeZone_transitions, "QTimeZone_transitions")
	symaddr(&addr_QTimeZone_systemTimeZoneId, "QTimeZone_systemTimeZoneId")
	symaddr(&addr_QTimeZone_systemTimeZone, "QTimeZone_systemTimeZone")
	symaddr(&addr_QTimeZone_utc, "QTimeZone_utc")
	symaddr(&addr_QTimeZone_isTimeZoneIdAvailable, "QTimeZone_isTimeZoneIdAvailable")
	symaddr(&addr_QTimeZone_availableTimeZoneIds, "QTimeZone_availableTimeZoneIds")
	symaddr(&addr_QTimeZone_availableTimeZoneIds2, "QTimeZone_availableTimeZoneIds2")
	symaddr(&addr_QTimeZone_availableTimeZoneIds3, "QTimeZone_availableTimeZoneIds3")
	symaddr(&addr_QTimeZone_ianaIdToWindowsId, "QTimeZone_ianaIdToWindowsId")
	symaddr(&addr_QTimeZone_windowsIdToDefaultIanaId, "QTimeZone_windowsIdToDefaultIanaId")
	symaddr(&addr_QTimeZone_windowsIdToDefaultIanaId2, "QTimeZone_windowsIdToDefaultIanaId2")
	symaddr(&addr_QTimeZone_windowsIdToIanaIds, "QTimeZone_windowsIdToIanaIds")
	symaddr(&addr_QTimeZone_windowsIdToIanaIds2, "QTimeZone_windowsIdToIanaIds2")
	symaddr(&addr_QTimeZone_displayName3, "QTimeZone_displayName3")
	symaddr(&addr_QTimeZone_displayName4, "QTimeZone_displayName4")
	symaddr(&addr_QTimeZone_displayName5, "QTimeZone_displayName5")
	symaddr(&addr_QTimeZone_displayName6, "QTimeZone_displayName6")
	symaddr(&addr_QTimeZone_OperatorEqual, "QTimeZone_OperatorEqual")
	symaddr(&addr_QTimeZone_OperatorNotEqual, "QTimeZone_OperatorNotEqual")
	reg(&pg_QTimeZone_delete, "QTimeZone_delete")
}

type QTimeZone struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QTimeZone) UnsafeGoID() int64     { return this.vid }
func (this *QTimeZone) UnsafeSetGoID(id int64) { this.vid = id }

func bindQTimeZone(h uintptr) *QTimeZone {
	if h == 0 {
		return nil
	}
	return &QTimeZone{h: h}
}

func newownedQTimeZone(h uintptr) *QTimeZone {
	o := bindQTimeZone(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QTimeZone).Delete)
	}
	return o
}

func (this *QTimeZone) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QTimeZone) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQTimeZone(h unsafe.Pointer) *QTimeZone { return bindQTimeZone(uintptr(h)) }
func UnsafeNewQTimeZoneOwned(h unsafe.Pointer) *QTimeZone { return newownedQTimeZone(uintptr(h)) }

func (this *QTimeZone) UnsafeRelease() { this.h = 0 }

func NewQTimeZone() *QTimeZone {
	__r, _, _ := purego.SyscallN(addr_QTimeZone_new)
	__o := newownedQTimeZone(__r)
	return __o
}

func NewQTimeZone2(spec QTimeZone__Initialization) *QTimeZone {
	__r, _, _ := purego.SyscallN(addr_QTimeZone_new2, uintptr(spec))
	__o := newownedQTimeZone(__r)
	return __o
}

func NewQTimeZone3(offsetSeconds int) *QTimeZone {
	__r, _, _ := purego.SyscallN(addr_QTimeZone_new3, uintptr(offsetSeconds))
	__o := newownedQTimeZone(__r)
	return __o
}

func NewQTimeZone4(ianaId *QByteArray) *QTimeZone {
	__r, _, _ := purego.SyscallN(addr_QTimeZone_new4, uintptr(ianaId.UnsafePointer()))
	__o := newownedQTimeZone(__r)
	runtime.KeepAlive(ianaId)
	return __o
}

func NewQTimeZone5(zoneId *QByteArray, offsetSeconds int, name string, abbreviation string) *QTimeZone {
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	var abbreviation_p uintptr
	if len(abbreviation) > 0 {
		abbreviation_p = uintptr(unsafe.Pointer(unsafe.StringData(abbreviation)))
	}
	__r, _, _ := purego.SyscallN(addr_QTimeZone_new5, uintptr(zoneId.UnsafePointer()), uintptr(offsetSeconds), name_p, uintptr(len(name)), abbreviation_p, uintptr(len(abbreviation)))
	__o := newownedQTimeZone(__r)
	runtime.KeepAlive(zoneId)
	runtime.KeepAlive(name)
	runtime.KeepAlive(abbreviation)
	return __o
}

func NewQTimeZone6(other *QTimeZone) *QTimeZone {
	__r, _, _ := purego.SyscallN(addr_QTimeZone_new6, uintptr(other.UnsafePointer()))
	__o := newownedQTimeZone(__r)
	runtime.KeepAlive(other)
	return __o
}

func NewQTimeZone7(zoneId *QByteArray, offsetSeconds int, name string, abbreviation string, territory QLocale__Country) *QTimeZone {
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	var abbreviation_p uintptr
	if len(abbreviation) > 0 {
		abbreviation_p = uintptr(unsafe.Pointer(unsafe.StringData(abbreviation)))
	}
	__r, _, _ := purego.SyscallN(addr_QTimeZone_new7, uintptr(zoneId.UnsafePointer()), uintptr(offsetSeconds), name_p, uintptr(len(name)), abbreviation_p, uintptr(len(abbreviation)), uintptr(territory))
	__o := newownedQTimeZone(__r)
	runtime.KeepAlive(zoneId)
	runtime.KeepAlive(name)
	runtime.KeepAlive(abbreviation)
	return __o
}

func NewQTimeZone8(zoneId *QByteArray, offsetSeconds int, name string, abbreviation string, territory QLocale__Country, comment string) *QTimeZone {
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	var abbreviation_p uintptr
	if len(abbreviation) > 0 {
		abbreviation_p = uintptr(unsafe.Pointer(unsafe.StringData(abbreviation)))
	}
	var comment_p uintptr
	if len(comment) > 0 {
		comment_p = uintptr(unsafe.Pointer(unsafe.StringData(comment)))
	}
	__r, _, _ := purego.SyscallN(addr_QTimeZone_new8, uintptr(zoneId.UnsafePointer()), uintptr(offsetSeconds), name_p, uintptr(len(name)), abbreviation_p, uintptr(len(abbreviation)), uintptr(territory), comment_p, uintptr(len(comment)))
	__o := newownedQTimeZone(__r)
	runtime.KeepAlive(zoneId)
	runtime.KeepAlive(name)
	runtime.KeepAlive(abbreviation)
	runtime.KeepAlive(comment)
	return __o
}

func (this *QTimeZone) OperatorAssign(other *QTimeZone) *QTimeZone {
	__r, _, _ := purego.SyscallN(addr_QTimeZone_OperatorAssign, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return bindQTimeZone(__r)
}

func (this *QTimeZone) Swap(other *QTimeZone) {
	purego.SyscallN(addr_QTimeZone_swap, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
}

func (this *QTimeZone) IsValid() bool {
	__r0, _, _ := purego.SyscallN(addr_QTimeZone_isValid, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QTimeZone_FromDurationAheadOfUtc(offset int64) *QTimeZone {
	__r, _, _ := purego.SyscallN(addr_QTimeZone_fromDurationAheadOfUtc, uintptr(offset))
	return newownedQTimeZone(__r)
}

func QTimeZone_FromSecondsAheadOfUtc(offset int) *QTimeZone {
	__r, _, _ := purego.SyscallN(addr_QTimeZone_fromSecondsAheadOfUtc, uintptr(offset))
	return newownedQTimeZone(__r)
}

func (this *QTimeZone) TimeSpec() TimeSpec {
	__r0, _, _ := purego.SyscallN(addr_QTimeZone_timeSpec, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return TimeSpec(__r)
}

func (this *QTimeZone) FixedSecondsAheadOfUtc() int {
	__r0, _, _ := purego.SyscallN(addr_QTimeZone_fixedSecondsAheadOfUtc, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func QTimeZone_IsUtcOrFixedOffset(spec TimeSpec) bool {
	__r0, _, _ := purego.SyscallN(addr_QTimeZone_isUtcOrFixedOffset, uintptr(spec))
	__r := int32(uint32(__r0))
	return __r != 0
}

func (this *QTimeZone) IsUtcOrFixedOffset() bool {
	__r0, _, _ := purego.SyscallN(addr_QTimeZone_isUtcOrFixedOffset2, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTimeZone) AsBackendZone() *QTimeZone {
	__r, _, _ := purego.SyscallN(addr_QTimeZone_asBackendZone, this.h)
	runtime.KeepAlive(this)
	return newownedQTimeZone(__r)
}

func (this *QTimeZone) HasAlternativeName(alias *QByteArrayView) bool {
	__r0, _, _ := purego.SyscallN(addr_QTimeZone_hasAlternativeName, this.h, uintptr(alias.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(alias)
	return __r != 0
}

func (this *QTimeZone) Id() *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QTimeZone_id, this.h)
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QTimeZone) Territory() QLocale__Country {
	__r0, _, _ := purego.SyscallN(addr_QTimeZone_territory, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QLocale__Country(__r)
}

func (this *QTimeZone) Country() QLocale__Country {
	__r0, _, _ := purego.SyscallN(addr_QTimeZone_country, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QLocale__Country(__r)
}

func (this *QTimeZone) Comment() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTimeZone_comment, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QTimeZone) DisplayName(atDateTime *QDateTime) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTimeZone_displayName, this.h, uintptr(atDateTime.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(atDateTime)
	return pgStr(__r, __n)
}

func (this *QTimeZone) DisplayNameWithTimeType(timeType QTimeZone__TimeType) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTimeZone_displayName2, this.h, uintptr(timeType), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QTimeZone) Abbreviation(atDateTime *QDateTime) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTimeZone_abbreviation, this.h, uintptr(atDateTime.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(atDateTime)
	return pgStr(__r, __n)
}

func (this *QTimeZone) OffsetFromUtc(atDateTime *QDateTime) int {
	__r0, _, _ := purego.SyscallN(addr_QTimeZone_offsetFromUtc, this.h, uintptr(atDateTime.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(atDateTime)
	return int(__r)
}

func (this *QTimeZone) StandardTimeOffset(atDateTime *QDateTime) int {
	__r0, _, _ := purego.SyscallN(addr_QTimeZone_standardTimeOffset, this.h, uintptr(atDateTime.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(atDateTime)
	return int(__r)
}

func (this *QTimeZone) DaylightTimeOffset(atDateTime *QDateTime) int {
	__r0, _, _ := purego.SyscallN(addr_QTimeZone_daylightTimeOffset, this.h, uintptr(atDateTime.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(atDateTime)
	return int(__r)
}

func (this *QTimeZone) HasDaylightTime() bool {
	__r0, _, _ := purego.SyscallN(addr_QTimeZone_hasDaylightTime, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTimeZone) IsDaylightTime(atDateTime *QDateTime) bool {
	__r0, _, _ := purego.SyscallN(addr_QTimeZone_isDaylightTime, this.h, uintptr(atDateTime.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(atDateTime)
	return __r != 0
}

func (this *QTimeZone) OffsetData(forDateTime *QDateTime) *QTimeZone__OffsetData {
	__r, _, _ := purego.SyscallN(addr_QTimeZone_offsetData, this.h, uintptr(forDateTime.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(forDateTime)
	return newownedQTimeZone__OffsetData(__r)
}

func (this *QTimeZone) HasTransitions() bool {
	__r0, _, _ := purego.SyscallN(addr_QTimeZone_hasTransitions, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTimeZone) NextTransition(afterDateTime *QDateTime) *QTimeZone__OffsetData {
	__r, _, _ := purego.SyscallN(addr_QTimeZone_nextTransition, this.h, uintptr(afterDateTime.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(afterDateTime)
	return newownedQTimeZone__OffsetData(__r)
}

func (this *QTimeZone) PreviousTransition(beforeDateTime *QDateTime) *QTimeZone__OffsetData {
	__r, _, _ := purego.SyscallN(addr_QTimeZone_previousTransition, this.h, uintptr(beforeDateTime.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(beforeDateTime)
	return newownedQTimeZone__OffsetData(__r)
}

func (this *QTimeZone) Transitions(fromDateTime *QDateTime, toDateTime *QDateTime) []*QTimeZone__OffsetData {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTimeZone_transitions, this.h, uintptr(fromDateTime.UnsafePointer()), uintptr(toDateTime.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(fromDateTime)
	runtime.KeepAlive(toDateTime)
	__out := make([]*QTimeZone__OffsetData, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQTimeZone__OffsetData(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QTimeZone_SystemTimeZoneId() *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QTimeZone_systemTimeZoneId)
	return newownedQByteArray(__r)
}

func QTimeZone_SystemTimeZone() *QTimeZone {
	__r, _, _ := purego.SyscallN(addr_QTimeZone_systemTimeZone)
	return newownedQTimeZone(__r)
}

func QTimeZone_Utc() *QTimeZone {
	__r, _, _ := purego.SyscallN(addr_QTimeZone_utc)
	return newownedQTimeZone(__r)
}

func QTimeZone_IsTimeZoneIdAvailable(ianaId *QByteArray) bool {
	__r0, _, _ := purego.SyscallN(addr_QTimeZone_isTimeZoneIdAvailable, uintptr(ianaId.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(ianaId)
	return __r != 0
}

func QTimeZone_AvailableTimeZoneIds() []*QByteArray {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTimeZone_availableTimeZoneIds, uintptr(unsafe.Pointer(&__n)))
	__out := make([]*QByteArray, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQByteArray(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QTimeZone_AvailableTimeZoneIdsWithTerritory(territory QLocale__Country) []*QByteArray {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTimeZone_availableTimeZoneIds2, uintptr(territory), uintptr(unsafe.Pointer(&__n)))
	__out := make([]*QByteArray, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQByteArray(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QTimeZone_AvailableTimeZoneIdsWithOffsetSeconds(offsetSeconds int) []*QByteArray {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTimeZone_availableTimeZoneIds3, uintptr(offsetSeconds), uintptr(unsafe.Pointer(&__n)))
	__out := make([]*QByteArray, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQByteArray(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QTimeZone_IanaIdToWindowsId(ianaId *QByteArray) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QTimeZone_ianaIdToWindowsId, uintptr(ianaId.UnsafePointer()))
	runtime.KeepAlive(ianaId)
	return newownedQByteArray(__r)
}

func QTimeZone_WindowsIdToDefaultIanaId(windowsId *QByteArray) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QTimeZone_windowsIdToDefaultIanaId, uintptr(windowsId.UnsafePointer()))
	runtime.KeepAlive(windowsId)
	return newownedQByteArray(__r)
}

func QTimeZone_WindowsIdToDefaultIanaId2(windowsId *QByteArray, territory QLocale__Country) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QTimeZone_windowsIdToDefaultIanaId2, uintptr(windowsId.UnsafePointer()), uintptr(territory))
	runtime.KeepAlive(windowsId)
	return newownedQByteArray(__r)
}

func QTimeZone_WindowsIdToIanaIds(windowsId *QByteArray) []*QByteArray {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTimeZone_windowsIdToIanaIds, uintptr(windowsId.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(windowsId)
	__out := make([]*QByteArray, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQByteArray(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QTimeZone_WindowsIdToIanaIds2(windowsId *QByteArray, territory QLocale__Country) []*QByteArray {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTimeZone_windowsIdToIanaIds2, uintptr(windowsId.UnsafePointer()), uintptr(territory), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(windowsId)
	__out := make([]*QByteArray, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQByteArray(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QTimeZone) DisplayName2(atDateTime *QDateTime, nameType QTimeZone__NameType) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTimeZone_displayName3, this.h, uintptr(atDateTime.UnsafePointer()), uintptr(nameType), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(atDateTime)
	return pgStr(__r, __n)
}

func (this *QTimeZone) DisplayName3(atDateTime *QDateTime, nameType QTimeZone__NameType, locale *QLocale) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTimeZone_displayName4, this.h, uintptr(atDateTime.UnsafePointer()), uintptr(nameType), uintptr(locale.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(atDateTime)
	runtime.KeepAlive(locale)
	return pgStr(__r, __n)
}

func (this *QTimeZone) DisplayName4(timeType QTimeZone__TimeType, nameType QTimeZone__NameType) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTimeZone_displayName5, this.h, uintptr(timeType), uintptr(nameType), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QTimeZone) DisplayName5(timeType QTimeZone__TimeType, nameType QTimeZone__NameType, locale *QLocale) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTimeZone_displayName6, this.h, uintptr(timeType), uintptr(nameType), uintptr(locale.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(locale)
	return pgStr(__r, __n)
}

func (this *QTimeZone) OperatorEqual(rhs *QTimeZone) bool {
	__r0, _, _ := purego.SyscallN(addr_QTimeZone_OperatorEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QTimeZone) OperatorNotEqual(rhs *QTimeZone) bool {
	__r0, _, _ := purego.SyscallN(addr_QTimeZone_OperatorNotEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QTimeZone) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QTimeZone_delete(this.h)
	}
	this.UnsafeRelease()
}
