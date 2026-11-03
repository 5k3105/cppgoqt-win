package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QDateTime_new uintptr
	addr_QDateTime_new2 uintptr
	addr_QDateTime_new3 uintptr
	addr_QDateTime_new4 uintptr
	addr_QDateTime_new5 uintptr
	addr_QDateTime_new6 uintptr
	addr_QDateTime_new7 uintptr
	addr_QDateTime_new8 uintptr
	addr_QDateTime_OperatorAssign uintptr
	addr_QDateTime_swap uintptr
	addr_QDateTime_isNull uintptr
	addr_QDateTime_isValid uintptr
	addr_QDateTime_date uintptr
	addr_QDateTime_time uintptr
	addr_QDateTime_timeSpec uintptr
	addr_QDateTime_offsetFromUtc uintptr
	addr_QDateTime_timeRepresentation uintptr
	addr_QDateTime_timeZone uintptr
	addr_QDateTime_timeZoneAbbreviation uintptr
	addr_QDateTime_isDaylightTime uintptr
	addr_QDateTime_toMSecsSinceEpoch uintptr
	addr_QDateTime_toSecsSinceEpoch uintptr
	addr_QDateTime_setDate uintptr
	addr_QDateTime_setTime uintptr
	addr_QDateTime_setTimeSpec uintptr
	addr_QDateTime_setOffsetFromUtc uintptr
	addr_QDateTime_setTimeZone uintptr
	addr_QDateTime_setMSecsSinceEpoch uintptr
	addr_QDateTime_setSecsSinceEpoch uintptr
	addr_QDateTime_toString uintptr
	addr_QDateTime_toString2 uintptr
	addr_QDateTime_toString3 uintptr
	addr_QDateTime_toString4 uintptr
	addr_QDateTime_toString5 uintptr
	addr_QDateTime_addDays uintptr
	addr_QDateTime_addMonths uintptr
	addr_QDateTime_addYears uintptr
	addr_QDateTime_addSecs uintptr
	addr_QDateTime_addMSecs uintptr
	addr_QDateTime_addDuration uintptr
	addr_QDateTime_toTimeSpec uintptr
	addr_QDateTime_toLocalTime uintptr
	addr_QDateTime_toUTC uintptr
	addr_QDateTime_toOffsetFromUtc uintptr
	addr_QDateTime_toTimeZone uintptr
	addr_QDateTime_daysTo uintptr
	addr_QDateTime_secsTo uintptr
	addr_QDateTime_msecsTo uintptr
	addr_QDateTime_currentDateTime uintptr
	addr_QDateTime_currentDateTime2 uintptr
	addr_QDateTime_currentDateTimeUtc uintptr
	addr_QDateTime_fromString uintptr
	addr_QDateTime_fromString2 uintptr
	addr_QDateTime_fromString3 uintptr
	addr_QDateTime_fromString4 uintptr
	addr_QDateTime_fromString5 uintptr
	addr_QDateTime_fromString6 uintptr
	addr_QDateTime_fromString7 uintptr
	addr_QDateTime_fromString8 uintptr
	addr_QDateTime_fromString9 uintptr
	addr_QDateTime_fromString10 uintptr
	addr_QDateTime_fromString11 uintptr
	addr_QDateTime_fromMSecsSinceEpoch uintptr
	addr_QDateTime_fromSecsSinceEpoch uintptr
	addr_QDateTime_fromMSecsSinceEpoch2 uintptr
	addr_QDateTime_fromSecsSinceEpoch2 uintptr
	addr_QDateTime_fromMSecsSinceEpoch3 uintptr
	addr_QDateTime_fromSecsSinceEpoch3 uintptr
	addr_QDateTime_currentMSecsSinceEpoch uintptr
	addr_QDateTime_currentSecsSinceEpoch uintptr
	addr_QDateTime_OperatorPlusAssign uintptr
	addr_QDateTime_OperatorMinusAssign uintptr
	addr_QDateTime_setDate2 uintptr
	addr_QDateTime_setTime2 uintptr
	addr_QDateTime_setTimeZone2 uintptr
	addr_QDateTime_toString6 uintptr
	addr_QDateTime_fromString12 uintptr
	addr_QDateTime_fromString13 uintptr
	addr_QDateTime_fromString14 uintptr
	addr_QDateTime_fromString15 uintptr
	addr_QDateTime_fromString16 uintptr
	addr_QDateTime_fromMSecsSinceEpoch4 uintptr
	addr_QDateTime_fromSecsSinceEpoch4 uintptr
	addr_QDateTime_OperatorMinus uintptr
	addr_QDateTime_OperatorPlus uintptr
	addr_QDateTime_OperatorMinus2 uintptr
	addr_QDateTime_OperatorEqual uintptr
	addr_QDateTime_OperatorNotEqual uintptr
	addr_QDateTime_OperatorLesser uintptr
	addr_QDateTime_OperatorGreater uintptr
	addr_QDateTime_OperatorLesserOrEqual uintptr
	addr_QDateTime_OperatorGreaterOrEqual uintptr
	pg_QDateTime_delete func(uintptr)
)

func init() {
	symaddr(&addr_QDateTime_new, "QDateTime_new")
	symaddr(&addr_QDateTime_new2, "QDateTime_new2")
	symaddr(&addr_QDateTime_new3, "QDateTime_new3")
	symaddr(&addr_QDateTime_new4, "QDateTime_new4")
	symaddr(&addr_QDateTime_new5, "QDateTime_new5")
	symaddr(&addr_QDateTime_new6, "QDateTime_new6")
	symaddr(&addr_QDateTime_new7, "QDateTime_new7")
	symaddr(&addr_QDateTime_new8, "QDateTime_new8")
	symaddr(&addr_QDateTime_OperatorAssign, "QDateTime_OperatorAssign")
	symaddr(&addr_QDateTime_swap, "QDateTime_swap")
	symaddr(&addr_QDateTime_isNull, "QDateTime_isNull")
	symaddr(&addr_QDateTime_isValid, "QDateTime_isValid")
	symaddr(&addr_QDateTime_date, "QDateTime_date")
	symaddr(&addr_QDateTime_time, "QDateTime_time")
	symaddr(&addr_QDateTime_timeSpec, "QDateTime_timeSpec")
	symaddr(&addr_QDateTime_offsetFromUtc, "QDateTime_offsetFromUtc")
	symaddr(&addr_QDateTime_timeRepresentation, "QDateTime_timeRepresentation")
	symaddr(&addr_QDateTime_timeZone, "QDateTime_timeZone")
	symaddr(&addr_QDateTime_timeZoneAbbreviation, "QDateTime_timeZoneAbbreviation")
	symaddr(&addr_QDateTime_isDaylightTime, "QDateTime_isDaylightTime")
	symaddr(&addr_QDateTime_toMSecsSinceEpoch, "QDateTime_toMSecsSinceEpoch")
	symaddr(&addr_QDateTime_toSecsSinceEpoch, "QDateTime_toSecsSinceEpoch")
	symaddr(&addr_QDateTime_setDate, "QDateTime_setDate")
	symaddr(&addr_QDateTime_setTime, "QDateTime_setTime")
	symaddr(&addr_QDateTime_setTimeSpec, "QDateTime_setTimeSpec")
	symaddr(&addr_QDateTime_setOffsetFromUtc, "QDateTime_setOffsetFromUtc")
	symaddr(&addr_QDateTime_setTimeZone, "QDateTime_setTimeZone")
	symaddr(&addr_QDateTime_setMSecsSinceEpoch, "QDateTime_setMSecsSinceEpoch")
	symaddr(&addr_QDateTime_setSecsSinceEpoch, "QDateTime_setSecsSinceEpoch")
	symaddr(&addr_QDateTime_toString, "QDateTime_toString")
	symaddr(&addr_QDateTime_toString2, "QDateTime_toString2")
	symaddr(&addr_QDateTime_toString3, "QDateTime_toString3")
	symaddr(&addr_QDateTime_toString4, "QDateTime_toString4")
	symaddr(&addr_QDateTime_toString5, "QDateTime_toString5")
	symaddr(&addr_QDateTime_addDays, "QDateTime_addDays")
	symaddr(&addr_QDateTime_addMonths, "QDateTime_addMonths")
	symaddr(&addr_QDateTime_addYears, "QDateTime_addYears")
	symaddr(&addr_QDateTime_addSecs, "QDateTime_addSecs")
	symaddr(&addr_QDateTime_addMSecs, "QDateTime_addMSecs")
	symaddr(&addr_QDateTime_addDuration, "QDateTime_addDuration")
	symaddr(&addr_QDateTime_toTimeSpec, "QDateTime_toTimeSpec")
	symaddr(&addr_QDateTime_toLocalTime, "QDateTime_toLocalTime")
	symaddr(&addr_QDateTime_toUTC, "QDateTime_toUTC")
	symaddr(&addr_QDateTime_toOffsetFromUtc, "QDateTime_toOffsetFromUtc")
	symaddr(&addr_QDateTime_toTimeZone, "QDateTime_toTimeZone")
	symaddr(&addr_QDateTime_daysTo, "QDateTime_daysTo")
	symaddr(&addr_QDateTime_secsTo, "QDateTime_secsTo")
	symaddr(&addr_QDateTime_msecsTo, "QDateTime_msecsTo")
	symaddr(&addr_QDateTime_currentDateTime, "QDateTime_currentDateTime")
	symaddr(&addr_QDateTime_currentDateTime2, "QDateTime_currentDateTime2")
	symaddr(&addr_QDateTime_currentDateTimeUtc, "QDateTime_currentDateTimeUtc")
	symaddr(&addr_QDateTime_fromString, "QDateTime_fromString")
	symaddr(&addr_QDateTime_fromString2, "QDateTime_fromString2")
	symaddr(&addr_QDateTime_fromString3, "QDateTime_fromString3")
	symaddr(&addr_QDateTime_fromString4, "QDateTime_fromString4")
	symaddr(&addr_QDateTime_fromString5, "QDateTime_fromString5")
	symaddr(&addr_QDateTime_fromString6, "QDateTime_fromString6")
	symaddr(&addr_QDateTime_fromString7, "QDateTime_fromString7")
	symaddr(&addr_QDateTime_fromString8, "QDateTime_fromString8")
	symaddr(&addr_QDateTime_fromString9, "QDateTime_fromString9")
	symaddr(&addr_QDateTime_fromString10, "QDateTime_fromString10")
	symaddr(&addr_QDateTime_fromString11, "QDateTime_fromString11")
	symaddr(&addr_QDateTime_fromMSecsSinceEpoch, "QDateTime_fromMSecsSinceEpoch")
	symaddr(&addr_QDateTime_fromSecsSinceEpoch, "QDateTime_fromSecsSinceEpoch")
	symaddr(&addr_QDateTime_fromMSecsSinceEpoch2, "QDateTime_fromMSecsSinceEpoch2")
	symaddr(&addr_QDateTime_fromSecsSinceEpoch2, "QDateTime_fromSecsSinceEpoch2")
	symaddr(&addr_QDateTime_fromMSecsSinceEpoch3, "QDateTime_fromMSecsSinceEpoch3")
	symaddr(&addr_QDateTime_fromSecsSinceEpoch3, "QDateTime_fromSecsSinceEpoch3")
	symaddr(&addr_QDateTime_currentMSecsSinceEpoch, "QDateTime_currentMSecsSinceEpoch")
	symaddr(&addr_QDateTime_currentSecsSinceEpoch, "QDateTime_currentSecsSinceEpoch")
	symaddr(&addr_QDateTime_OperatorPlusAssign, "QDateTime_OperatorPlusAssign")
	symaddr(&addr_QDateTime_OperatorMinusAssign, "QDateTime_OperatorMinusAssign")
	symaddr(&addr_QDateTime_setDate2, "QDateTime_setDate2")
	symaddr(&addr_QDateTime_setTime2, "QDateTime_setTime2")
	symaddr(&addr_QDateTime_setTimeZone2, "QDateTime_setTimeZone2")
	symaddr(&addr_QDateTime_toString6, "QDateTime_toString6")
	symaddr(&addr_QDateTime_fromString12, "QDateTime_fromString12")
	symaddr(&addr_QDateTime_fromString13, "QDateTime_fromString13")
	symaddr(&addr_QDateTime_fromString14, "QDateTime_fromString14")
	symaddr(&addr_QDateTime_fromString15, "QDateTime_fromString15")
	symaddr(&addr_QDateTime_fromString16, "QDateTime_fromString16")
	symaddr(&addr_QDateTime_fromMSecsSinceEpoch4, "QDateTime_fromMSecsSinceEpoch4")
	symaddr(&addr_QDateTime_fromSecsSinceEpoch4, "QDateTime_fromSecsSinceEpoch4")
	symaddr(&addr_QDateTime_OperatorMinus, "QDateTime_OperatorMinus")
	symaddr(&addr_QDateTime_OperatorPlus, "QDateTime_OperatorPlus")
	symaddr(&addr_QDateTime_OperatorMinus2, "QDateTime_OperatorMinus2")
	symaddr(&addr_QDateTime_OperatorEqual, "QDateTime_OperatorEqual")
	symaddr(&addr_QDateTime_OperatorNotEqual, "QDateTime_OperatorNotEqual")
	symaddr(&addr_QDateTime_OperatorLesser, "QDateTime_OperatorLesser")
	symaddr(&addr_QDateTime_OperatorGreater, "QDateTime_OperatorGreater")
	symaddr(&addr_QDateTime_OperatorLesserOrEqual, "QDateTime_OperatorLesserOrEqual")
	symaddr(&addr_QDateTime_OperatorGreaterOrEqual, "QDateTime_OperatorGreaterOrEqual")
	reg(&pg_QDateTime_delete, "QDateTime_delete")
}

type QDateTime struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QDateTime) UnsafeGoID() int64     { return this.vid }
func (this *QDateTime) UnsafeSetGoID(id int64) { this.vid = id }

func bindQDateTime(h uintptr) *QDateTime {
	if h == 0 {
		return nil
	}
	return &QDateTime{h: h}
}

func newownedQDateTime(h uintptr) *QDateTime {
	o := bindQDateTime(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QDateTime).Delete)
	}
	return o
}

func (this *QDateTime) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QDateTime) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQDateTime(h unsafe.Pointer) *QDateTime { return bindQDateTime(uintptr(h)) }
func UnsafeNewQDateTimeOwned(h unsafe.Pointer) *QDateTime { return newownedQDateTime(uintptr(h)) }

func (this *QDateTime) UnsafeRelease() { this.h = 0 }

func NewQDateTime() *QDateTime {
	__r, _, _ := purego.SyscallN(addr_QDateTime_new)
	__o := newownedQDateTime(__r)
	return __o
}

func NewQDateTime2(date *QDate, time *QTime, spec TimeSpec) *QDateTime {
	__r, _, _ := purego.SyscallN(addr_QDateTime_new2, uintptr(date.UnsafePointer()), uintptr(time.UnsafePointer()), uintptr(spec))
	__o := newownedQDateTime(__r)
	runtime.KeepAlive(date)
	runtime.KeepAlive(time)
	return __o
}

func NewQDateTime3(date *QDate, time *QTime, timeZone *QTimeZone) *QDateTime {
	__r, _, _ := purego.SyscallN(addr_QDateTime_new3, uintptr(date.UnsafePointer()), uintptr(time.UnsafePointer()), uintptr(timeZone.UnsafePointer()))
	__o := newownedQDateTime(__r)
	runtime.KeepAlive(date)
	runtime.KeepAlive(time)
	runtime.KeepAlive(timeZone)
	return __o
}

func NewQDateTime4(date *QDate, time *QTime) *QDateTime {
	__r, _, _ := purego.SyscallN(addr_QDateTime_new4, uintptr(date.UnsafePointer()), uintptr(time.UnsafePointer()))
	__o := newownedQDateTime(__r)
	runtime.KeepAlive(date)
	runtime.KeepAlive(time)
	return __o
}

func NewQDateTime5(other *QDateTime) *QDateTime {
	__r, _, _ := purego.SyscallN(addr_QDateTime_new5, uintptr(other.UnsafePointer()))
	__o := newownedQDateTime(__r)
	runtime.KeepAlive(other)
	return __o
}

func NewQDateTime6(date *QDate, time *QTime, spec TimeSpec, offsetSeconds int) *QDateTime {
	__r, _, _ := purego.SyscallN(addr_QDateTime_new6, uintptr(date.UnsafePointer()), uintptr(time.UnsafePointer()), uintptr(spec), uintptr(offsetSeconds))
	__o := newownedQDateTime(__r)
	runtime.KeepAlive(date)
	runtime.KeepAlive(time)
	return __o
}

func NewQDateTime7(date *QDate, time *QTime, timeZone *QTimeZone, resolve QDateTime__TransitionResolution) *QDateTime {
	__r, _, _ := purego.SyscallN(addr_QDateTime_new7, uintptr(date.UnsafePointer()), uintptr(time.UnsafePointer()), uintptr(timeZone.UnsafePointer()), uintptr(resolve))
	__o := newownedQDateTime(__r)
	runtime.KeepAlive(date)
	runtime.KeepAlive(time)
	runtime.KeepAlive(timeZone)
	return __o
}

func NewQDateTime8(date *QDate, time *QTime, resolve QDateTime__TransitionResolution) *QDateTime {
	__r, _, _ := purego.SyscallN(addr_QDateTime_new8, uintptr(date.UnsafePointer()), uintptr(time.UnsafePointer()), uintptr(resolve))
	__o := newownedQDateTime(__r)
	runtime.KeepAlive(date)
	runtime.KeepAlive(time)
	return __o
}

func (this *QDateTime) OperatorAssign(other *QDateTime) *QDateTime {
	__r, _, _ := purego.SyscallN(addr_QDateTime_OperatorAssign, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return bindQDateTime(__r)
}

func (this *QDateTime) Swap(other *QDateTime) {
	purego.SyscallN(addr_QDateTime_swap, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
}

func (this *QDateTime) IsNull() bool {
	__r0, _, _ := purego.SyscallN(addr_QDateTime_isNull, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QDateTime) IsValid() bool {
	__r0, _, _ := purego.SyscallN(addr_QDateTime_isValid, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QDateTime) Date() *QDate {
	__r, _, _ := purego.SyscallN(addr_QDateTime_date, this.h)
	runtime.KeepAlive(this)
	return newownedQDate(__r)
}

func (this *QDateTime) Time() *QTime {
	__r, _, _ := purego.SyscallN(addr_QDateTime_time, this.h)
	runtime.KeepAlive(this)
	return newownedQTime(__r)
}

func (this *QDateTime) TimeSpec() TimeSpec {
	__r0, _, _ := purego.SyscallN(addr_QDateTime_timeSpec, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return TimeSpec(__r)
}

func (this *QDateTime) OffsetFromUtc() int {
	__r0, _, _ := purego.SyscallN(addr_QDateTime_offsetFromUtc, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QDateTime) TimeRepresentation() *QTimeZone {
	__r, _, _ := purego.SyscallN(addr_QDateTime_timeRepresentation, this.h)
	runtime.KeepAlive(this)
	return newownedQTimeZone(__r)
}

func (this *QDateTime) TimeZone() *QTimeZone {
	__r, _, _ := purego.SyscallN(addr_QDateTime_timeZone, this.h)
	runtime.KeepAlive(this)
	return newownedQTimeZone(__r)
}

func (this *QDateTime) TimeZoneAbbreviation() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDateTime_timeZoneAbbreviation, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QDateTime) IsDaylightTime() bool {
	__r0, _, _ := purego.SyscallN(addr_QDateTime_isDaylightTime, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QDateTime) ToMSecsSinceEpoch() int64 {
	__r0, _, _ := purego.SyscallN(addr_QDateTime_toMSecsSinceEpoch, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QDateTime) ToSecsSinceEpoch() int64 {
	__r0, _, _ := purego.SyscallN(addr_QDateTime_toSecsSinceEpoch, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QDateTime) SetDate(date *QDate) {
	purego.SyscallN(addr_QDateTime_setDate, this.h, uintptr(date.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(date)
}

func (this *QDateTime) SetTime(time *QTime) {
	purego.SyscallN(addr_QDateTime_setTime, this.h, uintptr(time.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(time)
}

func (this *QDateTime) SetTimeSpec(spec TimeSpec) {
	purego.SyscallN(addr_QDateTime_setTimeSpec, this.h, uintptr(spec))
	runtime.KeepAlive(this)
}

func (this *QDateTime) SetOffsetFromUtc(offsetSeconds int) {
	purego.SyscallN(addr_QDateTime_setOffsetFromUtc, this.h, uintptr(offsetSeconds))
	runtime.KeepAlive(this)
}

func (this *QDateTime) SetTimeZone(toZone *QTimeZone) {
	purego.SyscallN(addr_QDateTime_setTimeZone, this.h, uintptr(toZone.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(toZone)
}

func (this *QDateTime) SetMSecsSinceEpoch(msecs int64) {
	purego.SyscallN(addr_QDateTime_setMSecsSinceEpoch, this.h, uintptr(msecs))
	runtime.KeepAlive(this)
}

func (this *QDateTime) SetSecsSinceEpoch(secs int64) {
	purego.SyscallN(addr_QDateTime_setSecsSinceEpoch, this.h, uintptr(secs))
	runtime.KeepAlive(this)
}

func (this *QDateTime) ToString() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDateTime_toString, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QDateTime) ToStringWithFormat(format string) string {
	var format_p uintptr
	if len(format) > 0 {
		format_p = uintptr(unsafe.Pointer(unsafe.StringData(format)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDateTime_toString2, this.h, format_p, uintptr(len(format)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(format)
	return pgStr(__r, __n)
}

func (this *QDateTime) ToString2(format string, cal *QCalendar) string {
	var format_p uintptr
	if len(format) > 0 {
		format_p = uintptr(unsafe.Pointer(unsafe.StringData(format)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDateTime_toString3, this.h, format_p, uintptr(len(format)), uintptr(cal.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(format)
	runtime.KeepAlive(cal)
	return pgStr(__r, __n)
}

func (this *QDateTime) ToString3(format *QStringView) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDateTime_toString4, this.h, uintptr(format.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(format)
	return pgStr(__r, __n)
}

func (this *QDateTime) ToString4(format *QStringView, cal *QCalendar) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDateTime_toString5, this.h, uintptr(format.UnsafePointer()), uintptr(cal.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(format)
	runtime.KeepAlive(cal)
	return pgStr(__r, __n)
}

func (this *QDateTime) AddDays(days int64) *QDateTime {
	__r, _, _ := purego.SyscallN(addr_QDateTime_addDays, this.h, uintptr(days))
	runtime.KeepAlive(this)
	return newownedQDateTime(__r)
}

func (this *QDateTime) AddMonths(months int) *QDateTime {
	__r, _, _ := purego.SyscallN(addr_QDateTime_addMonths, this.h, uintptr(months))
	runtime.KeepAlive(this)
	return newownedQDateTime(__r)
}

func (this *QDateTime) AddYears(years int) *QDateTime {
	__r, _, _ := purego.SyscallN(addr_QDateTime_addYears, this.h, uintptr(years))
	runtime.KeepAlive(this)
	return newownedQDateTime(__r)
}

func (this *QDateTime) AddSecs(secs int64) *QDateTime {
	__r, _, _ := purego.SyscallN(addr_QDateTime_addSecs, this.h, uintptr(secs))
	runtime.KeepAlive(this)
	return newownedQDateTime(__r)
}

func (this *QDateTime) AddMSecs(msecs int64) *QDateTime {
	__r, _, _ := purego.SyscallN(addr_QDateTime_addMSecs, this.h, uintptr(msecs))
	runtime.KeepAlive(this)
	return newownedQDateTime(__r)
}

func (this *QDateTime) AddDuration(msecs int64) *QDateTime {
	__r, _, _ := purego.SyscallN(addr_QDateTime_addDuration, this.h, uintptr(msecs))
	runtime.KeepAlive(this)
	return newownedQDateTime(__r)
}

func (this *QDateTime) ToTimeSpec(spec TimeSpec) *QDateTime {
	__r, _, _ := purego.SyscallN(addr_QDateTime_toTimeSpec, this.h, uintptr(spec))
	runtime.KeepAlive(this)
	return newownedQDateTime(__r)
}

func (this *QDateTime) ToLocalTime() *QDateTime {
	__r, _, _ := purego.SyscallN(addr_QDateTime_toLocalTime, this.h)
	runtime.KeepAlive(this)
	return newownedQDateTime(__r)
}

func (this *QDateTime) ToUTC() *QDateTime {
	__r, _, _ := purego.SyscallN(addr_QDateTime_toUTC, this.h)
	runtime.KeepAlive(this)
	return newownedQDateTime(__r)
}

func (this *QDateTime) ToOffsetFromUtc(offsetSeconds int) *QDateTime {
	__r, _, _ := purego.SyscallN(addr_QDateTime_toOffsetFromUtc, this.h, uintptr(offsetSeconds))
	runtime.KeepAlive(this)
	return newownedQDateTime(__r)
}

func (this *QDateTime) ToTimeZone(toZone *QTimeZone) *QDateTime {
	__r, _, _ := purego.SyscallN(addr_QDateTime_toTimeZone, this.h, uintptr(toZone.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(toZone)
	return newownedQDateTime(__r)
}

func (this *QDateTime) DaysTo(arg1 *QDateTime) int64 {
	__r0, _, _ := purego.SyscallN(addr_QDateTime_daysTo, this.h, uintptr(arg1.UnsafePointer()))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return __r
}

func (this *QDateTime) SecsTo(arg1 *QDateTime) int64 {
	__r0, _, _ := purego.SyscallN(addr_QDateTime_secsTo, this.h, uintptr(arg1.UnsafePointer()))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return __r
}

func (this *QDateTime) MsecsTo(arg1 *QDateTime) int64 {
	__r0, _, _ := purego.SyscallN(addr_QDateTime_msecsTo, this.h, uintptr(arg1.UnsafePointer()))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return __r
}

func QDateTime_CurrentDateTime(zone *QTimeZone) *QDateTime {
	__r, _, _ := purego.SyscallN(addr_QDateTime_currentDateTime, uintptr(zone.UnsafePointer()))
	runtime.KeepAlive(zone)
	return newownedQDateTime(__r)
}

func QDateTime_CurrentDateTime2() *QDateTime {
	__r, _, _ := purego.SyscallN(addr_QDateTime_currentDateTime2)
	return newownedQDateTime(__r)
}

func QDateTime_CurrentDateTimeUtc() *QDateTime {
	__r, _, _ := purego.SyscallN(addr_QDateTime_currentDateTimeUtc)
	return newownedQDateTime(__r)
}

func QDateTime_FromString(string_ *QStringView) *QDateTime {
	__r, _, _ := purego.SyscallN(addr_QDateTime_fromString, uintptr(string_.UnsafePointer()))
	runtime.KeepAlive(string_)
	return newownedQDateTime(__r)
}

func QDateTime_FromStringWithString(string_ string) *QDateTime {
	var string__p uintptr
	if len(string_) > 0 {
		string__p = uintptr(unsafe.Pointer(unsafe.StringData(string_)))
	}
	__r, _, _ := purego.SyscallN(addr_QDateTime_fromString2, string__p, uintptr(len(string_)))
	runtime.KeepAlive(string_)
	return newownedQDateTime(__r)
}

func QDateTime_FromString2(string_ *QStringView, format *QStringView, cal *QCalendar) *QDateTime {
	__r, _, _ := purego.SyscallN(addr_QDateTime_fromString3, uintptr(string_.UnsafePointer()), uintptr(format.UnsafePointer()), uintptr(cal.UnsafePointer()))
	runtime.KeepAlive(string_)
	runtime.KeepAlive(format)
	runtime.KeepAlive(cal)
	return newownedQDateTime(__r)
}

func QDateTime_FromString3(string_ string, format *QStringView, cal *QCalendar) *QDateTime {
	var string__p uintptr
	if len(string_) > 0 {
		string__p = uintptr(unsafe.Pointer(unsafe.StringData(string_)))
	}
	__r, _, _ := purego.SyscallN(addr_QDateTime_fromString4, string__p, uintptr(len(string_)), uintptr(format.UnsafePointer()), uintptr(cal.UnsafePointer()))
	runtime.KeepAlive(string_)
	runtime.KeepAlive(format)
	runtime.KeepAlive(cal)
	return newownedQDateTime(__r)
}

func QDateTime_FromString4(string_ string, format string, cal *QCalendar) *QDateTime {
	var string__p uintptr
	if len(string_) > 0 {
		string__p = uintptr(unsafe.Pointer(unsafe.StringData(string_)))
	}
	var format_p uintptr
	if len(format) > 0 {
		format_p = uintptr(unsafe.Pointer(unsafe.StringData(format)))
	}
	__r, _, _ := purego.SyscallN(addr_QDateTime_fromString5, string__p, uintptr(len(string_)), format_p, uintptr(len(format)), uintptr(cal.UnsafePointer()))
	runtime.KeepAlive(string_)
	runtime.KeepAlive(format)
	runtime.KeepAlive(cal)
	return newownedQDateTime(__r)
}

func QDateTime_FromString5(string_ *QStringView, format *QStringView) *QDateTime {
	__r, _, _ := purego.SyscallN(addr_QDateTime_fromString6, uintptr(string_.UnsafePointer()), uintptr(format.UnsafePointer()))
	runtime.KeepAlive(string_)
	runtime.KeepAlive(format)
	return newownedQDateTime(__r)
}

func QDateTime_FromString6(string_ *QStringView, format *QStringView, baseYear int, cal *QCalendar) *QDateTime {
	__r, _, _ := purego.SyscallN(addr_QDateTime_fromString7, uintptr(string_.UnsafePointer()), uintptr(format.UnsafePointer()), uintptr(baseYear), uintptr(cal.UnsafePointer()))
	runtime.KeepAlive(string_)
	runtime.KeepAlive(format)
	runtime.KeepAlive(cal)
	return newownedQDateTime(__r)
}

func QDateTime_FromString7(string_ string, format *QStringView) *QDateTime {
	var string__p uintptr
	if len(string_) > 0 {
		string__p = uintptr(unsafe.Pointer(unsafe.StringData(string_)))
	}
	__r, _, _ := purego.SyscallN(addr_QDateTime_fromString8, string__p, uintptr(len(string_)), uintptr(format.UnsafePointer()))
	runtime.KeepAlive(string_)
	runtime.KeepAlive(format)
	return newownedQDateTime(__r)
}

func QDateTime_FromString8(string_ string, format *QStringView, baseYear int, cal *QCalendar) *QDateTime {
	var string__p uintptr
	if len(string_) > 0 {
		string__p = uintptr(unsafe.Pointer(unsafe.StringData(string_)))
	}
	__r, _, _ := purego.SyscallN(addr_QDateTime_fromString9, string__p, uintptr(len(string_)), uintptr(format.UnsafePointer()), uintptr(baseYear), uintptr(cal.UnsafePointer()))
	runtime.KeepAlive(string_)
	runtime.KeepAlive(format)
	runtime.KeepAlive(cal)
	return newownedQDateTime(__r)
}

func QDateTime_FromString9(string_ string, format string) *QDateTime {
	var string__p uintptr
	if len(string_) > 0 {
		string__p = uintptr(unsafe.Pointer(unsafe.StringData(string_)))
	}
	var format_p uintptr
	if len(format) > 0 {
		format_p = uintptr(unsafe.Pointer(unsafe.StringData(format)))
	}
	__r, _, _ := purego.SyscallN(addr_QDateTime_fromString10, string__p, uintptr(len(string_)), format_p, uintptr(len(format)))
	runtime.KeepAlive(string_)
	runtime.KeepAlive(format)
	return newownedQDateTime(__r)
}

func QDateTime_FromString10(string_ string, format string, baseYear int, cal *QCalendar) *QDateTime {
	var string__p uintptr
	if len(string_) > 0 {
		string__p = uintptr(unsafe.Pointer(unsafe.StringData(string_)))
	}
	var format_p uintptr
	if len(format) > 0 {
		format_p = uintptr(unsafe.Pointer(unsafe.StringData(format)))
	}
	__r, _, _ := purego.SyscallN(addr_QDateTime_fromString11, string__p, uintptr(len(string_)), format_p, uintptr(len(format)), uintptr(baseYear), uintptr(cal.UnsafePointer()))
	runtime.KeepAlive(string_)
	runtime.KeepAlive(format)
	runtime.KeepAlive(cal)
	return newownedQDateTime(__r)
}

func QDateTime_FromMSecsSinceEpoch(msecs int64, spec TimeSpec) *QDateTime {
	__r, _, _ := purego.SyscallN(addr_QDateTime_fromMSecsSinceEpoch, uintptr(msecs), uintptr(spec))
	return newownedQDateTime(__r)
}

func QDateTime_FromSecsSinceEpoch(secs int64, spec TimeSpec) *QDateTime {
	__r, _, _ := purego.SyscallN(addr_QDateTime_fromSecsSinceEpoch, uintptr(secs), uintptr(spec))
	return newownedQDateTime(__r)
}

func QDateTime_FromMSecsSinceEpoch2(msecs int64, timeZone *QTimeZone) *QDateTime {
	__r, _, _ := purego.SyscallN(addr_QDateTime_fromMSecsSinceEpoch2, uintptr(msecs), uintptr(timeZone.UnsafePointer()))
	runtime.KeepAlive(timeZone)
	return newownedQDateTime(__r)
}

func QDateTime_FromSecsSinceEpoch2(secs int64, timeZone *QTimeZone) *QDateTime {
	__r, _, _ := purego.SyscallN(addr_QDateTime_fromSecsSinceEpoch2, uintptr(secs), uintptr(timeZone.UnsafePointer()))
	runtime.KeepAlive(timeZone)
	return newownedQDateTime(__r)
}

func QDateTime_FromMSecsSinceEpochWithMsecs(msecs int64) *QDateTime {
	__r, _, _ := purego.SyscallN(addr_QDateTime_fromMSecsSinceEpoch3, uintptr(msecs))
	return newownedQDateTime(__r)
}

func QDateTime_FromSecsSinceEpochWithSecs(secs int64) *QDateTime {
	__r, _, _ := purego.SyscallN(addr_QDateTime_fromSecsSinceEpoch3, uintptr(secs))
	return newownedQDateTime(__r)
}

func QDateTime_CurrentMSecsSinceEpoch() int64 {
	__r0, _, _ := purego.SyscallN(addr_QDateTime_currentMSecsSinceEpoch)
	__r := int64(__r0)
	return __r
}

func QDateTime_CurrentSecsSinceEpoch() int64 {
	__r0, _, _ := purego.SyscallN(addr_QDateTime_currentSecsSinceEpoch)
	__r := int64(__r0)
	return __r
}

func (this *QDateTime) OperatorPlusAssign(duration int64) *QDateTime {
	__r, _, _ := purego.SyscallN(addr_QDateTime_OperatorPlusAssign, this.h, uintptr(duration))
	runtime.KeepAlive(this)
	return bindQDateTime(__r)
}

func (this *QDateTime) OperatorMinusAssign(duration int64) *QDateTime {
	__r, _, _ := purego.SyscallN(addr_QDateTime_OperatorMinusAssign, this.h, uintptr(duration))
	runtime.KeepAlive(this)
	return bindQDateTime(__r)
}

func (this *QDateTime) SetDate2(date *QDate, resolve QDateTime__TransitionResolution) {
	purego.SyscallN(addr_QDateTime_setDate2, this.h, uintptr(date.UnsafePointer()), uintptr(resolve))
	runtime.KeepAlive(this)
	runtime.KeepAlive(date)
}

func (this *QDateTime) SetTime2(time *QTime, resolve QDateTime__TransitionResolution) {
	purego.SyscallN(addr_QDateTime_setTime2, this.h, uintptr(time.UnsafePointer()), uintptr(resolve))
	runtime.KeepAlive(this)
	runtime.KeepAlive(time)
}

func (this *QDateTime) SetTimeZone2(toZone *QTimeZone, resolve QDateTime__TransitionResolution) {
	purego.SyscallN(addr_QDateTime_setTimeZone2, this.h, uintptr(toZone.UnsafePointer()), uintptr(resolve))
	runtime.KeepAlive(this)
	runtime.KeepAlive(toZone)
}

func (this *QDateTime) ToString5(format DateFormat) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDateTime_toString6, this.h, uintptr(format), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func QDateTime_FromString11(string_ *QStringView, format DateFormat) *QDateTime {
	__r, _, _ := purego.SyscallN(addr_QDateTime_fromString12, uintptr(string_.UnsafePointer()), uintptr(format))
	runtime.KeepAlive(string_)
	return newownedQDateTime(__r)
}

func QDateTime_FromString12(string_ string, format DateFormat) *QDateTime {
	var string__p uintptr
	if len(string_) > 0 {
		string__p = uintptr(unsafe.Pointer(unsafe.StringData(string_)))
	}
	__r, _, _ := purego.SyscallN(addr_QDateTime_fromString13, string__p, uintptr(len(string_)), uintptr(format))
	runtime.KeepAlive(string_)
	return newownedQDateTime(__r)
}

func QDateTime_FromString13(string_ *QStringView, format *QStringView, baseYear int) *QDateTime {
	__r, _, _ := purego.SyscallN(addr_QDateTime_fromString14, uintptr(string_.UnsafePointer()), uintptr(format.UnsafePointer()), uintptr(baseYear))
	runtime.KeepAlive(string_)
	runtime.KeepAlive(format)
	return newownedQDateTime(__r)
}

func QDateTime_FromString14(string_ string, format *QStringView, baseYear int) *QDateTime {
	var string__p uintptr
	if len(string_) > 0 {
		string__p = uintptr(unsafe.Pointer(unsafe.StringData(string_)))
	}
	__r, _, _ := purego.SyscallN(addr_QDateTime_fromString15, string__p, uintptr(len(string_)), uintptr(format.UnsafePointer()), uintptr(baseYear))
	runtime.KeepAlive(string_)
	runtime.KeepAlive(format)
	return newownedQDateTime(__r)
}

func QDateTime_FromString15(string_ string, format string, baseYear int) *QDateTime {
	var string__p uintptr
	if len(string_) > 0 {
		string__p = uintptr(unsafe.Pointer(unsafe.StringData(string_)))
	}
	var format_p uintptr
	if len(format) > 0 {
		format_p = uintptr(unsafe.Pointer(unsafe.StringData(format)))
	}
	__r, _, _ := purego.SyscallN(addr_QDateTime_fromString16, string__p, uintptr(len(string_)), format_p, uintptr(len(format)), uintptr(baseYear))
	runtime.KeepAlive(string_)
	runtime.KeepAlive(format)
	return newownedQDateTime(__r)
}

func QDateTime_FromMSecsSinceEpoch3(msecs int64, spec TimeSpec, offsetFromUtc int) *QDateTime {
	__r, _, _ := purego.SyscallN(addr_QDateTime_fromMSecsSinceEpoch4, uintptr(msecs), uintptr(spec), uintptr(offsetFromUtc))
	return newownedQDateTime(__r)
}

func QDateTime_FromSecsSinceEpoch3(secs int64, spec TimeSpec, offsetFromUtc int) *QDateTime {
	__r, _, _ := purego.SyscallN(addr_QDateTime_fromSecsSinceEpoch4, uintptr(secs), uintptr(spec), uintptr(offsetFromUtc))
	return newownedQDateTime(__r)
}

func (this *QDateTime) OperatorMinus(rhs *QDateTime) int64 {
	__r0, _, _ := purego.SyscallN(addr_QDateTime_OperatorMinus, this.h, uintptr(rhs.UnsafePointer()))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r
}

func (this *QDateTime) OperatorPlus(duration int64) *QDateTime {
	__r, _, _ := purego.SyscallN(addr_QDateTime_OperatorPlus, this.h, uintptr(duration))
	runtime.KeepAlive(this)
	return newownedQDateTime(__r)
}

func (this *QDateTime) OperatorMinusWithDuration(duration int64) *QDateTime {
	__r, _, _ := purego.SyscallN(addr_QDateTime_OperatorMinus2, this.h, uintptr(duration))
	runtime.KeepAlive(this)
	return newownedQDateTime(__r)
}

func (this *QDateTime) OperatorEqual(rhs *QDateTime) bool {
	__r0, _, _ := purego.SyscallN(addr_QDateTime_OperatorEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QDateTime) OperatorNotEqual(rhs *QDateTime) bool {
	__r0, _, _ := purego.SyscallN(addr_QDateTime_OperatorNotEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QDateTime) OperatorLesser(rhs *QDateTime) bool {
	__r0, _, _ := purego.SyscallN(addr_QDateTime_OperatorLesser, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QDateTime) OperatorGreater(rhs *QDateTime) bool {
	__r0, _, _ := purego.SyscallN(addr_QDateTime_OperatorGreater, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QDateTime) OperatorLesserOrEqual(rhs *QDateTime) bool {
	__r0, _, _ := purego.SyscallN(addr_QDateTime_OperatorLesserOrEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QDateTime) OperatorGreaterOrEqual(rhs *QDateTime) bool {
	__r0, _, _ := purego.SyscallN(addr_QDateTime_OperatorGreaterOrEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QDateTime) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QDateTime_delete(this.h)
	}
	this.UnsafeRelease()
}
