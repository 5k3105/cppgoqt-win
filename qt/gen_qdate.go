package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QDate_new uintptr
	addr_QDate_new2 uintptr
	addr_QDate_new3 uintptr
	addr_QDate_new4 uintptr
	addr_QDate_isNull uintptr
	addr_QDate_isValid uintptr
	addr_QDate_year uintptr
	addr_QDate_month uintptr
	addr_QDate_day uintptr
	addr_QDate_dayOfWeek uintptr
	addr_QDate_dayOfYear uintptr
	addr_QDate_daysInMonth uintptr
	addr_QDate_daysInYear uintptr
	addr_QDate_weekNumber uintptr
	addr_QDate_year2 uintptr
	addr_QDate_month2 uintptr
	addr_QDate_day2 uintptr
	addr_QDate_dayOfWeek2 uintptr
	addr_QDate_dayOfYear2 uintptr
	addr_QDate_daysInMonth2 uintptr
	addr_QDate_daysInYear2 uintptr
	addr_QDate_startOfDay uintptr
	addr_QDate_endOfDay uintptr
	addr_QDate_startOfDay2 uintptr
	addr_QDate_endOfDay2 uintptr
	addr_QDate_startOfDay3 uintptr
	addr_QDate_endOfDay3 uintptr
	addr_QDate_toString uintptr
	addr_QDate_toString2 uintptr
	addr_QDate_toString3 uintptr
	addr_QDate_toString4 uintptr
	addr_QDate_toString5 uintptr
	addr_QDate_setDate uintptr
	addr_QDate_setDate2 uintptr
	addr_QDate_getDate uintptr
	addr_QDate_addDays uintptr
	addr_QDate_addMonths uintptr
	addr_QDate_addYears uintptr
	addr_QDate_addMonths2 uintptr
	addr_QDate_addYears2 uintptr
	addr_QDate_daysTo uintptr
	addr_QDate_currentDate uintptr
	addr_QDate_fromString uintptr
	addr_QDate_fromString2 uintptr
	addr_QDate_fromString3 uintptr
	addr_QDate_fromString4 uintptr
	addr_QDate_fromString5 uintptr
	addr_QDate_fromString6 uintptr
	addr_QDate_fromString7 uintptr
	addr_QDate_fromString8 uintptr
	addr_QDate_fromString9 uintptr
	addr_QDate_fromString10 uintptr
	addr_QDate_fromString11 uintptr
	addr_QDate_isValid2 uintptr
	addr_QDate_isLeapYear uintptr
	addr_QDate_fromJulianDay uintptr
	addr_QDate_toJulianDay uintptr
	addr_QDate_weekNumber2 uintptr
	addr_QDate_startOfDay4 uintptr
	addr_QDate_endOfDay4 uintptr
	addr_QDate_toString6 uintptr
	addr_QDate_fromString12 uintptr
	addr_QDate_fromString13 uintptr
	addr_QDate_fromString14 uintptr
	addr_QDate_fromString15 uintptr
	addr_QDate_fromString16 uintptr
	addr_QDate_OperatorEqual uintptr
	addr_QDate_OperatorNotEqual uintptr
	addr_QDate_OperatorLesser uintptr
	addr_QDate_OperatorGreater uintptr
	addr_QDate_OperatorLesserOrEqual uintptr
	addr_QDate_OperatorGreaterOrEqual uintptr
	addr_QDate_OperatorPlusPlus uintptr
	addr_QDate_OperatorMinusMinus uintptr
	addr_QDate_OperatorPlusPlus2 uintptr
	addr_QDate_OperatorMinusMinus2 uintptr
	pg_QDate_delete func(uintptr)
)

func init() {
	symaddr(&addr_QDate_new, "QDate_new")
	symaddr(&addr_QDate_new2, "QDate_new2")
	symaddr(&addr_QDate_new3, "QDate_new3")
	symaddr(&addr_QDate_new4, "QDate_new4")
	symaddr(&addr_QDate_isNull, "QDate_isNull")
	symaddr(&addr_QDate_isValid, "QDate_isValid")
	symaddr(&addr_QDate_year, "QDate_year")
	symaddr(&addr_QDate_month, "QDate_month")
	symaddr(&addr_QDate_day, "QDate_day")
	symaddr(&addr_QDate_dayOfWeek, "QDate_dayOfWeek")
	symaddr(&addr_QDate_dayOfYear, "QDate_dayOfYear")
	symaddr(&addr_QDate_daysInMonth, "QDate_daysInMonth")
	symaddr(&addr_QDate_daysInYear, "QDate_daysInYear")
	symaddr(&addr_QDate_weekNumber, "QDate_weekNumber")
	symaddr(&addr_QDate_year2, "QDate_year2")
	symaddr(&addr_QDate_month2, "QDate_month2")
	symaddr(&addr_QDate_day2, "QDate_day2")
	symaddr(&addr_QDate_dayOfWeek2, "QDate_dayOfWeek2")
	symaddr(&addr_QDate_dayOfYear2, "QDate_dayOfYear2")
	symaddr(&addr_QDate_daysInMonth2, "QDate_daysInMonth2")
	symaddr(&addr_QDate_daysInYear2, "QDate_daysInYear2")
	symaddr(&addr_QDate_startOfDay, "QDate_startOfDay")
	symaddr(&addr_QDate_endOfDay, "QDate_endOfDay")
	symaddr(&addr_QDate_startOfDay2, "QDate_startOfDay2")
	symaddr(&addr_QDate_endOfDay2, "QDate_endOfDay2")
	symaddr(&addr_QDate_startOfDay3, "QDate_startOfDay3")
	symaddr(&addr_QDate_endOfDay3, "QDate_endOfDay3")
	symaddr(&addr_QDate_toString, "QDate_toString")
	symaddr(&addr_QDate_toString2, "QDate_toString2")
	symaddr(&addr_QDate_toString3, "QDate_toString3")
	symaddr(&addr_QDate_toString4, "QDate_toString4")
	symaddr(&addr_QDate_toString5, "QDate_toString5")
	symaddr(&addr_QDate_setDate, "QDate_setDate")
	symaddr(&addr_QDate_setDate2, "QDate_setDate2")
	symaddr(&addr_QDate_getDate, "QDate_getDate")
	symaddr(&addr_QDate_addDays, "QDate_addDays")
	symaddr(&addr_QDate_addMonths, "QDate_addMonths")
	symaddr(&addr_QDate_addYears, "QDate_addYears")
	symaddr(&addr_QDate_addMonths2, "QDate_addMonths2")
	symaddr(&addr_QDate_addYears2, "QDate_addYears2")
	symaddr(&addr_QDate_daysTo, "QDate_daysTo")
	symaddr(&addr_QDate_currentDate, "QDate_currentDate")
	symaddr(&addr_QDate_fromString, "QDate_fromString")
	symaddr(&addr_QDate_fromString2, "QDate_fromString2")
	symaddr(&addr_QDate_fromString3, "QDate_fromString3")
	symaddr(&addr_QDate_fromString4, "QDate_fromString4")
	symaddr(&addr_QDate_fromString5, "QDate_fromString5")
	symaddr(&addr_QDate_fromString6, "QDate_fromString6")
	symaddr(&addr_QDate_fromString7, "QDate_fromString7")
	symaddr(&addr_QDate_fromString8, "QDate_fromString8")
	symaddr(&addr_QDate_fromString9, "QDate_fromString9")
	symaddr(&addr_QDate_fromString10, "QDate_fromString10")
	symaddr(&addr_QDate_fromString11, "QDate_fromString11")
	symaddr(&addr_QDate_isValid2, "QDate_isValid2")
	symaddr(&addr_QDate_isLeapYear, "QDate_isLeapYear")
	symaddr(&addr_QDate_fromJulianDay, "QDate_fromJulianDay")
	symaddr(&addr_QDate_toJulianDay, "QDate_toJulianDay")
	symaddr(&addr_QDate_weekNumber2, "QDate_weekNumber2")
	symaddr(&addr_QDate_startOfDay4, "QDate_startOfDay4")
	symaddr(&addr_QDate_endOfDay4, "QDate_endOfDay4")
	symaddr(&addr_QDate_toString6, "QDate_toString6")
	symaddr(&addr_QDate_fromString12, "QDate_fromString12")
	symaddr(&addr_QDate_fromString13, "QDate_fromString13")
	symaddr(&addr_QDate_fromString14, "QDate_fromString14")
	symaddr(&addr_QDate_fromString15, "QDate_fromString15")
	symaddr(&addr_QDate_fromString16, "QDate_fromString16")
	symaddr(&addr_QDate_OperatorEqual, "QDate_OperatorEqual")
	symaddr(&addr_QDate_OperatorNotEqual, "QDate_OperatorNotEqual")
	symaddr(&addr_QDate_OperatorLesser, "QDate_OperatorLesser")
	symaddr(&addr_QDate_OperatorGreater, "QDate_OperatorGreater")
	symaddr(&addr_QDate_OperatorLesserOrEqual, "QDate_OperatorLesserOrEqual")
	symaddr(&addr_QDate_OperatorGreaterOrEqual, "QDate_OperatorGreaterOrEqual")
	symaddr(&addr_QDate_OperatorPlusPlus, "QDate_OperatorPlusPlus")
	symaddr(&addr_QDate_OperatorMinusMinus, "QDate_OperatorMinusMinus")
	symaddr(&addr_QDate_OperatorPlusPlus2, "QDate_OperatorPlusPlus2")
	symaddr(&addr_QDate_OperatorMinusMinus2, "QDate_OperatorMinusMinus2")
	reg(&pg_QDate_delete, "QDate_delete")
}

type QDate struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QDate) UnsafeGoID() int64     { return this.vid }
func (this *QDate) UnsafeSetGoID(id int64) { this.vid = id }

func bindQDate(h uintptr) *QDate {
	if h == 0 {
		return nil
	}
	return &QDate{h: h}
}

func newownedQDate(h uintptr) *QDate {
	o := bindQDate(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QDate).Delete)
	}
	return o
}

func (this *QDate) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QDate) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQDate(h unsafe.Pointer) *QDate { return bindQDate(uintptr(h)) }
func UnsafeNewQDateOwned(h unsafe.Pointer) *QDate { return newownedQDate(uintptr(h)) }

func (this *QDate) UnsafeRelease() { this.h = 0 }

func NewQDate() *QDate {
	__r, _, _ := purego.SyscallN(addr_QDate_new)
	__o := newownedQDate(__r)
	return __o
}

func NewQDate2(y int, m int, d int) *QDate {
	__r, _, _ := purego.SyscallN(addr_QDate_new2, uintptr(y), uintptr(m), uintptr(d))
	__o := newownedQDate(__r)
	return __o
}

func NewQDate3(y int, m int, d int, cal *QCalendar) *QDate {
	__r, _, _ := purego.SyscallN(addr_QDate_new3, uintptr(y), uintptr(m), uintptr(d), uintptr(cal.UnsafePointer()))
	__o := newownedQDate(__r)
	runtime.KeepAlive(cal)
	return __o
}

func NewQDate4(arg1 *QDate) *QDate {
	__r, _, _ := purego.SyscallN(addr_QDate_new4, uintptr(arg1.UnsafePointer()))
	__o := newownedQDate(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QDate) IsNull() bool {
	__r0, _, _ := purego.SyscallN(addr_QDate_isNull, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QDate) IsValid() bool {
	__r0, _, _ := purego.SyscallN(addr_QDate_isValid, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QDate) Year() int {
	__r0, _, _ := purego.SyscallN(addr_QDate_year, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QDate) Month() int {
	__r0, _, _ := purego.SyscallN(addr_QDate_month, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QDate) Day() int {
	__r0, _, _ := purego.SyscallN(addr_QDate_day, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QDate) DayOfWeek() int {
	__r0, _, _ := purego.SyscallN(addr_QDate_dayOfWeek, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QDate) DayOfYear() int {
	__r0, _, _ := purego.SyscallN(addr_QDate_dayOfYear, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QDate) DaysInMonth() int {
	__r0, _, _ := purego.SyscallN(addr_QDate_daysInMonth, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QDate) DaysInYear() int {
	__r0, _, _ := purego.SyscallN(addr_QDate_daysInYear, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QDate) WeekNumber() int {
	__r0, _, _ := purego.SyscallN(addr_QDate_weekNumber, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QDate) YearWithCal(cal *QCalendar) int {
	__r0, _, _ := purego.SyscallN(addr_QDate_year2, this.h, uintptr(cal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(cal)
	return int(__r)
}

func (this *QDate) MonthWithCal(cal *QCalendar) int {
	__r0, _, _ := purego.SyscallN(addr_QDate_month2, this.h, uintptr(cal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(cal)
	return int(__r)
}

func (this *QDate) DayWithCal(cal *QCalendar) int {
	__r0, _, _ := purego.SyscallN(addr_QDate_day2, this.h, uintptr(cal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(cal)
	return int(__r)
}

func (this *QDate) DayOfWeekWithCal(cal *QCalendar) int {
	__r0, _, _ := purego.SyscallN(addr_QDate_dayOfWeek2, this.h, uintptr(cal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(cal)
	return int(__r)
}

func (this *QDate) DayOfYearWithCal(cal *QCalendar) int {
	__r0, _, _ := purego.SyscallN(addr_QDate_dayOfYear2, this.h, uintptr(cal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(cal)
	return int(__r)
}

func (this *QDate) DaysInMonthWithCal(cal *QCalendar) int {
	__r0, _, _ := purego.SyscallN(addr_QDate_daysInMonth2, this.h, uintptr(cal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(cal)
	return int(__r)
}

func (this *QDate) DaysInYearWithCal(cal *QCalendar) int {
	__r0, _, _ := purego.SyscallN(addr_QDate_daysInYear2, this.h, uintptr(cal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(cal)
	return int(__r)
}

func (this *QDate) StartOfDay(spec TimeSpec) *QDateTime {
	__r, _, _ := purego.SyscallN(addr_QDate_startOfDay, this.h, uintptr(spec))
	runtime.KeepAlive(this)
	return newownedQDateTime(__r)
}

func (this *QDate) EndOfDay(spec TimeSpec) *QDateTime {
	__r, _, _ := purego.SyscallN(addr_QDate_endOfDay, this.h, uintptr(spec))
	runtime.KeepAlive(this)
	return newownedQDateTime(__r)
}

func (this *QDate) StartOfDayWithZone(zone *QTimeZone) *QDateTime {
	__r, _, _ := purego.SyscallN(addr_QDate_startOfDay2, this.h, uintptr(zone.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(zone)
	return newownedQDateTime(__r)
}

func (this *QDate) EndOfDayWithZone(zone *QTimeZone) *QDateTime {
	__r, _, _ := purego.SyscallN(addr_QDate_endOfDay2, this.h, uintptr(zone.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(zone)
	return newownedQDateTime(__r)
}

func (this *QDate) StartOfDay2() *QDateTime {
	__r, _, _ := purego.SyscallN(addr_QDate_startOfDay3, this.h)
	runtime.KeepAlive(this)
	return newownedQDateTime(__r)
}

func (this *QDate) EndOfDay2() *QDateTime {
	__r, _, _ := purego.SyscallN(addr_QDate_endOfDay3, this.h)
	runtime.KeepAlive(this)
	return newownedQDateTime(__r)
}

func (this *QDate) ToString() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDate_toString, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QDate) ToStringWithFormat(format string) string {
	var format_p uintptr
	if len(format) > 0 {
		format_p = uintptr(unsafe.Pointer(unsafe.StringData(format)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDate_toString2, this.h, format_p, uintptr(len(format)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(format)
	return pgStr(__r, __n)
}

func (this *QDate) ToString2(format string, cal *QCalendar) string {
	var format_p uintptr
	if len(format) > 0 {
		format_p = uintptr(unsafe.Pointer(unsafe.StringData(format)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDate_toString3, this.h, format_p, uintptr(len(format)), uintptr(cal.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(format)
	runtime.KeepAlive(cal)
	return pgStr(__r, __n)
}

func (this *QDate) ToString3(format *QStringView) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDate_toString4, this.h, uintptr(format.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(format)
	return pgStr(__r, __n)
}

func (this *QDate) ToString4(format *QStringView, cal *QCalendar) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDate_toString5, this.h, uintptr(format.UnsafePointer()), uintptr(cal.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(format)
	runtime.KeepAlive(cal)
	return pgStr(__r, __n)
}

func (this *QDate) SetDate(year int, month int, day int) bool {
	__r0, _, _ := purego.SyscallN(addr_QDate_setDate, this.h, uintptr(year), uintptr(month), uintptr(day))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QDate) SetDate2(year int, month int, day int, cal *QCalendar) bool {
	__r0, _, _ := purego.SyscallN(addr_QDate_setDate2, this.h, uintptr(year), uintptr(month), uintptr(day), uintptr(cal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(cal)
	return __r != 0
}

func (this *QDate) GetDate() (int, int, int) {
	var year_out int32
	var month_out int32
	var day_out int32
	purego.SyscallN(addr_QDate_getDate, this.h, uintptr(unsafe.Pointer(&year_out)), uintptr(unsafe.Pointer(&month_out)), uintptr(unsafe.Pointer(&day_out)))
	runtime.KeepAlive(this)
	return int(year_out), int(month_out), int(day_out)
}

func (this *QDate) AddDays(days int64) *QDate {
	__r, _, _ := purego.SyscallN(addr_QDate_addDays, this.h, uintptr(days))
	runtime.KeepAlive(this)
	return newownedQDate(__r)
}

func (this *QDate) AddMonths(months int) *QDate {
	__r, _, _ := purego.SyscallN(addr_QDate_addMonths, this.h, uintptr(months))
	runtime.KeepAlive(this)
	return newownedQDate(__r)
}

func (this *QDate) AddYears(years int) *QDate {
	__r, _, _ := purego.SyscallN(addr_QDate_addYears, this.h, uintptr(years))
	runtime.KeepAlive(this)
	return newownedQDate(__r)
}

func (this *QDate) AddMonths2(months int, cal *QCalendar) *QDate {
	__r, _, _ := purego.SyscallN(addr_QDate_addMonths2, this.h, uintptr(months), uintptr(cal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(cal)
	return newownedQDate(__r)
}

func (this *QDate) AddYears2(years int, cal *QCalendar) *QDate {
	__r, _, _ := purego.SyscallN(addr_QDate_addYears2, this.h, uintptr(years), uintptr(cal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(cal)
	return newownedQDate(__r)
}

func (this *QDate) DaysTo(d *QDate) int64 {
	__r0, _, _ := purego.SyscallN(addr_QDate_daysTo, this.h, uintptr(d.UnsafePointer()))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(d)
	return __r
}

func QDate_CurrentDate() *QDate {
	__r, _, _ := purego.SyscallN(addr_QDate_currentDate)
	return newownedQDate(__r)
}

func QDate_FromString(string_ *QStringView) *QDate {
	__r, _, _ := purego.SyscallN(addr_QDate_fromString, uintptr(string_.UnsafePointer()))
	runtime.KeepAlive(string_)
	return newownedQDate(__r)
}

func QDate_FromStringWithString(string_ string) *QDate {
	var string__p uintptr
	if len(string_) > 0 {
		string__p = uintptr(unsafe.Pointer(unsafe.StringData(string_)))
	}
	__r, _, _ := purego.SyscallN(addr_QDate_fromString2, string__p, uintptr(len(string_)))
	runtime.KeepAlive(string_)
	return newownedQDate(__r)
}

func QDate_FromString2(string_ *QStringView, format *QStringView, cal *QCalendar) *QDate {
	__r, _, _ := purego.SyscallN(addr_QDate_fromString3, uintptr(string_.UnsafePointer()), uintptr(format.UnsafePointer()), uintptr(cal.UnsafePointer()))
	runtime.KeepAlive(string_)
	runtime.KeepAlive(format)
	runtime.KeepAlive(cal)
	return newownedQDate(__r)
}

func QDate_FromString3(string_ string, format *QStringView, cal *QCalendar) *QDate {
	var string__p uintptr
	if len(string_) > 0 {
		string__p = uintptr(unsafe.Pointer(unsafe.StringData(string_)))
	}
	__r, _, _ := purego.SyscallN(addr_QDate_fromString4, string__p, uintptr(len(string_)), uintptr(format.UnsafePointer()), uintptr(cal.UnsafePointer()))
	runtime.KeepAlive(string_)
	runtime.KeepAlive(format)
	runtime.KeepAlive(cal)
	return newownedQDate(__r)
}

func QDate_FromString4(string_ string, format string, cal *QCalendar) *QDate {
	var string__p uintptr
	if len(string_) > 0 {
		string__p = uintptr(unsafe.Pointer(unsafe.StringData(string_)))
	}
	var format_p uintptr
	if len(format) > 0 {
		format_p = uintptr(unsafe.Pointer(unsafe.StringData(format)))
	}
	__r, _, _ := purego.SyscallN(addr_QDate_fromString5, string__p, uintptr(len(string_)), format_p, uintptr(len(format)), uintptr(cal.UnsafePointer()))
	runtime.KeepAlive(string_)
	runtime.KeepAlive(format)
	runtime.KeepAlive(cal)
	return newownedQDate(__r)
}

func QDate_FromString5(string_ *QStringView, format *QStringView) *QDate {
	__r, _, _ := purego.SyscallN(addr_QDate_fromString6, uintptr(string_.UnsafePointer()), uintptr(format.UnsafePointer()))
	runtime.KeepAlive(string_)
	runtime.KeepAlive(format)
	return newownedQDate(__r)
}

func QDate_FromString6(string_ *QStringView, format *QStringView, baseYear int, cal *QCalendar) *QDate {
	__r, _, _ := purego.SyscallN(addr_QDate_fromString7, uintptr(string_.UnsafePointer()), uintptr(format.UnsafePointer()), uintptr(baseYear), uintptr(cal.UnsafePointer()))
	runtime.KeepAlive(string_)
	runtime.KeepAlive(format)
	runtime.KeepAlive(cal)
	return newownedQDate(__r)
}

func QDate_FromString7(string_ string, format *QStringView) *QDate {
	var string__p uintptr
	if len(string_) > 0 {
		string__p = uintptr(unsafe.Pointer(unsafe.StringData(string_)))
	}
	__r, _, _ := purego.SyscallN(addr_QDate_fromString8, string__p, uintptr(len(string_)), uintptr(format.UnsafePointer()))
	runtime.KeepAlive(string_)
	runtime.KeepAlive(format)
	return newownedQDate(__r)
}

func QDate_FromString8(string_ string, format *QStringView, baseYear int, cal *QCalendar) *QDate {
	var string__p uintptr
	if len(string_) > 0 {
		string__p = uintptr(unsafe.Pointer(unsafe.StringData(string_)))
	}
	__r, _, _ := purego.SyscallN(addr_QDate_fromString9, string__p, uintptr(len(string_)), uintptr(format.UnsafePointer()), uintptr(baseYear), uintptr(cal.UnsafePointer()))
	runtime.KeepAlive(string_)
	runtime.KeepAlive(format)
	runtime.KeepAlive(cal)
	return newownedQDate(__r)
}

func QDate_FromString9(string_ string, format string) *QDate {
	var string__p uintptr
	if len(string_) > 0 {
		string__p = uintptr(unsafe.Pointer(unsafe.StringData(string_)))
	}
	var format_p uintptr
	if len(format) > 0 {
		format_p = uintptr(unsafe.Pointer(unsafe.StringData(format)))
	}
	__r, _, _ := purego.SyscallN(addr_QDate_fromString10, string__p, uintptr(len(string_)), format_p, uintptr(len(format)))
	runtime.KeepAlive(string_)
	runtime.KeepAlive(format)
	return newownedQDate(__r)
}

func QDate_FromString10(string_ string, format string, baseYear int, cal *QCalendar) *QDate {
	var string__p uintptr
	if len(string_) > 0 {
		string__p = uintptr(unsafe.Pointer(unsafe.StringData(string_)))
	}
	var format_p uintptr
	if len(format) > 0 {
		format_p = uintptr(unsafe.Pointer(unsafe.StringData(format)))
	}
	__r, _, _ := purego.SyscallN(addr_QDate_fromString11, string__p, uintptr(len(string_)), format_p, uintptr(len(format)), uintptr(baseYear), uintptr(cal.UnsafePointer()))
	runtime.KeepAlive(string_)
	runtime.KeepAlive(format)
	runtime.KeepAlive(cal)
	return newownedQDate(__r)
}

func QDate_IsValid(y int, m int, d int) bool {
	__r0, _, _ := purego.SyscallN(addr_QDate_isValid2, uintptr(y), uintptr(m), uintptr(d))
	__r := int32(uint32(__r0))
	return __r != 0
}

func QDate_IsLeapYear(year int) bool {
	__r0, _, _ := purego.SyscallN(addr_QDate_isLeapYear, uintptr(year))
	__r := int32(uint32(__r0))
	return __r != 0
}

func QDate_FromJulianDay(jd_ int64) *QDate {
	__r, _, _ := purego.SyscallN(addr_QDate_fromJulianDay, uintptr(jd_))
	return newownedQDate(__r)
}

func (this *QDate) ToJulianDay() int64 {
	__r0, _, _ := purego.SyscallN(addr_QDate_toJulianDay, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QDate) WeekNumberWithYearNum() (int, int) {
	var yearNum_out int32
	__r0, _, _ := purego.SyscallN(addr_QDate_weekNumber2, this.h, uintptr(unsafe.Pointer(&yearNum_out)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r), int(yearNum_out)
}

func (this *QDate) StartOfDay3(spec TimeSpec, offsetSeconds int) *QDateTime {
	__r, _, _ := purego.SyscallN(addr_QDate_startOfDay4, this.h, uintptr(spec), uintptr(offsetSeconds))
	runtime.KeepAlive(this)
	return newownedQDateTime(__r)
}

func (this *QDate) EndOfDay3(spec TimeSpec, offsetSeconds int) *QDateTime {
	__r, _, _ := purego.SyscallN(addr_QDate_endOfDay4, this.h, uintptr(spec), uintptr(offsetSeconds))
	runtime.KeepAlive(this)
	return newownedQDateTime(__r)
}

func (this *QDate) ToString5(format DateFormat) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDate_toString6, this.h, uintptr(format), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func QDate_FromString11(string_ *QStringView, format DateFormat) *QDate {
	__r, _, _ := purego.SyscallN(addr_QDate_fromString12, uintptr(string_.UnsafePointer()), uintptr(format))
	runtime.KeepAlive(string_)
	return newownedQDate(__r)
}

func QDate_FromString12(string_ string, format DateFormat) *QDate {
	var string__p uintptr
	if len(string_) > 0 {
		string__p = uintptr(unsafe.Pointer(unsafe.StringData(string_)))
	}
	__r, _, _ := purego.SyscallN(addr_QDate_fromString13, string__p, uintptr(len(string_)), uintptr(format))
	runtime.KeepAlive(string_)
	return newownedQDate(__r)
}

func QDate_FromString13(string_ *QStringView, format *QStringView, baseYear int) *QDate {
	__r, _, _ := purego.SyscallN(addr_QDate_fromString14, uintptr(string_.UnsafePointer()), uintptr(format.UnsafePointer()), uintptr(baseYear))
	runtime.KeepAlive(string_)
	runtime.KeepAlive(format)
	return newownedQDate(__r)
}

func QDate_FromString14(string_ string, format *QStringView, baseYear int) *QDate {
	var string__p uintptr
	if len(string_) > 0 {
		string__p = uintptr(unsafe.Pointer(unsafe.StringData(string_)))
	}
	__r, _, _ := purego.SyscallN(addr_QDate_fromString15, string__p, uintptr(len(string_)), uintptr(format.UnsafePointer()), uintptr(baseYear))
	runtime.KeepAlive(string_)
	runtime.KeepAlive(format)
	return newownedQDate(__r)
}

func QDate_FromString15(string_ string, format string, baseYear int) *QDate {
	var string__p uintptr
	if len(string_) > 0 {
		string__p = uintptr(unsafe.Pointer(unsafe.StringData(string_)))
	}
	var format_p uintptr
	if len(format) > 0 {
		format_p = uintptr(unsafe.Pointer(unsafe.StringData(format)))
	}
	__r, _, _ := purego.SyscallN(addr_QDate_fromString16, string__p, uintptr(len(string_)), format_p, uintptr(len(format)), uintptr(baseYear))
	runtime.KeepAlive(string_)
	runtime.KeepAlive(format)
	return newownedQDate(__r)
}

func (this *QDate) OperatorEqual(rhs *QDate) bool {
	__r0, _, _ := purego.SyscallN(addr_QDate_OperatorEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QDate) OperatorNotEqual(rhs *QDate) bool {
	__r0, _, _ := purego.SyscallN(addr_QDate_OperatorNotEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QDate) OperatorLesser(rhs *QDate) bool {
	__r0, _, _ := purego.SyscallN(addr_QDate_OperatorLesser, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QDate) OperatorGreater(rhs *QDate) bool {
	__r0, _, _ := purego.SyscallN(addr_QDate_OperatorGreater, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QDate) OperatorLesserOrEqual(rhs *QDate) bool {
	__r0, _, _ := purego.SyscallN(addr_QDate_OperatorLesserOrEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QDate) OperatorGreaterOrEqual(rhs *QDate) bool {
	__r0, _, _ := purego.SyscallN(addr_QDate_OperatorGreaterOrEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QDate) OperatorPlusPlus() *QDate {
	__r, _, _ := purego.SyscallN(addr_QDate_OperatorPlusPlus, this.h)
	runtime.KeepAlive(this)
	return bindQDate(__r)
}

func (this *QDate) OperatorMinusMinus() *QDate {
	__r, _, _ := purego.SyscallN(addr_QDate_OperatorMinusMinus, this.h)
	runtime.KeepAlive(this)
	return bindQDate(__r)
}

func (this *QDate) OperatorPlusPlusWithInt(arg1 int) *QDate {
	__r, _, _ := purego.SyscallN(addr_QDate_OperatorPlusPlus2, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return newownedQDate(__r)
}

func (this *QDate) OperatorMinusMinusWithInt(arg1 int) *QDate {
	__r, _, _ := purego.SyscallN(addr_QDate_OperatorMinusMinus2, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return newownedQDate(__r)
}

func (this *QDate) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QDate_delete(this.h)
	}
	this.UnsafeRelease()
}
