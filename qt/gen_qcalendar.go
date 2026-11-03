package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QCalendar_new uintptr
	addr_QCalendar_new2 uintptr
	addr_QCalendar_new3 uintptr
	addr_QCalendar_new4 uintptr
	addr_QCalendar_isValid uintptr
	addr_QCalendar_daysInMonth uintptr
	addr_QCalendar_daysInYear uintptr
	addr_QCalendar_monthsInYear uintptr
	addr_QCalendar_isDateValid uintptr
	addr_QCalendar_isLeapYear uintptr
	addr_QCalendar_isGregorian uintptr
	addr_QCalendar_isLunar uintptr
	addr_QCalendar_isLuniSolar uintptr
	addr_QCalendar_isSolar uintptr
	addr_QCalendar_isProleptic uintptr
	addr_QCalendar_hasYearZero uintptr
	addr_QCalendar_maximumDaysInMonth uintptr
	addr_QCalendar_minimumDaysInMonth uintptr
	addr_QCalendar_maximumMonthsInYear uintptr
	addr_QCalendar_name uintptr
	addr_QCalendar_dateFromParts uintptr
	addr_QCalendar_dateFromParts2 uintptr
	addr_QCalendar_matchCenturyToWeekday uintptr
	addr_QCalendar_partsFromDate uintptr
	addr_QCalendar_dayOfWeek uintptr
	addr_QCalendar_monthName uintptr
	addr_QCalendar_standaloneMonthName uintptr
	addr_QCalendar_weekDayName uintptr
	addr_QCalendar_standaloneWeekDayName uintptr
	addr_QCalendar_dateTimeToString uintptr
	addr_QCalendar_availableCalendars uintptr
	addr_QCalendar_daysInMonth2 uintptr
	addr_QCalendar_monthName2 uintptr
	addr_QCalendar_monthName3 uintptr
	addr_QCalendar_standaloneMonthName2 uintptr
	addr_QCalendar_standaloneMonthName3 uintptr
	addr_QCalendar_weekDayName2 uintptr
	addr_QCalendar_standaloneWeekDayName2 uintptr
	pg_QCalendar_delete func(uintptr)
)

func init() {
	symaddr(&addr_QCalendar_new, "QCalendar_new")
	symaddr(&addr_QCalendar_new2, "QCalendar_new2")
	symaddr(&addr_QCalendar_new3, "QCalendar_new3")
	symaddr(&addr_QCalendar_new4, "QCalendar_new4")
	symaddr(&addr_QCalendar_isValid, "QCalendar_isValid")
	symaddr(&addr_QCalendar_daysInMonth, "QCalendar_daysInMonth")
	symaddr(&addr_QCalendar_daysInYear, "QCalendar_daysInYear")
	symaddr(&addr_QCalendar_monthsInYear, "QCalendar_monthsInYear")
	symaddr(&addr_QCalendar_isDateValid, "QCalendar_isDateValid")
	symaddr(&addr_QCalendar_isLeapYear, "QCalendar_isLeapYear")
	symaddr(&addr_QCalendar_isGregorian, "QCalendar_isGregorian")
	symaddr(&addr_QCalendar_isLunar, "QCalendar_isLunar")
	symaddr(&addr_QCalendar_isLuniSolar, "QCalendar_isLuniSolar")
	symaddr(&addr_QCalendar_isSolar, "QCalendar_isSolar")
	symaddr(&addr_QCalendar_isProleptic, "QCalendar_isProleptic")
	symaddr(&addr_QCalendar_hasYearZero, "QCalendar_hasYearZero")
	symaddr(&addr_QCalendar_maximumDaysInMonth, "QCalendar_maximumDaysInMonth")
	symaddr(&addr_QCalendar_minimumDaysInMonth, "QCalendar_minimumDaysInMonth")
	symaddr(&addr_QCalendar_maximumMonthsInYear, "QCalendar_maximumMonthsInYear")
	symaddr(&addr_QCalendar_name, "QCalendar_name")
	symaddr(&addr_QCalendar_dateFromParts, "QCalendar_dateFromParts")
	symaddr(&addr_QCalendar_dateFromParts2, "QCalendar_dateFromParts2")
	symaddr(&addr_QCalendar_matchCenturyToWeekday, "QCalendar_matchCenturyToWeekday")
	symaddr(&addr_QCalendar_partsFromDate, "QCalendar_partsFromDate")
	symaddr(&addr_QCalendar_dayOfWeek, "QCalendar_dayOfWeek")
	symaddr(&addr_QCalendar_monthName, "QCalendar_monthName")
	symaddr(&addr_QCalendar_standaloneMonthName, "QCalendar_standaloneMonthName")
	symaddr(&addr_QCalendar_weekDayName, "QCalendar_weekDayName")
	symaddr(&addr_QCalendar_standaloneWeekDayName, "QCalendar_standaloneWeekDayName")
	symaddr(&addr_QCalendar_dateTimeToString, "QCalendar_dateTimeToString")
	symaddr(&addr_QCalendar_availableCalendars, "QCalendar_availableCalendars")
	symaddr(&addr_QCalendar_daysInMonth2, "QCalendar_daysInMonth2")
	symaddr(&addr_QCalendar_monthName2, "QCalendar_monthName2")
	symaddr(&addr_QCalendar_monthName3, "QCalendar_monthName3")
	symaddr(&addr_QCalendar_standaloneMonthName2, "QCalendar_standaloneMonthName2")
	symaddr(&addr_QCalendar_standaloneMonthName3, "QCalendar_standaloneMonthName3")
	symaddr(&addr_QCalendar_weekDayName2, "QCalendar_weekDayName2")
	symaddr(&addr_QCalendar_standaloneWeekDayName2, "QCalendar_standaloneWeekDayName2")
	reg(&pg_QCalendar_delete, "QCalendar_delete")
}

type QCalendar struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QCalendar) UnsafeGoID() int64     { return this.vid }
func (this *QCalendar) UnsafeSetGoID(id int64) { this.vid = id }

func bindQCalendar(h uintptr) *QCalendar {
	if h == 0 {
		return nil
	}
	return &QCalendar{h: h}
}

func newownedQCalendar(h uintptr) *QCalendar {
	o := bindQCalendar(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QCalendar).Delete)
	}
	return o
}

func (this *QCalendar) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QCalendar) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQCalendar(h unsafe.Pointer) *QCalendar { return bindQCalendar(uintptr(h)) }
func UnsafeNewQCalendarOwned(h unsafe.Pointer) *QCalendar { return newownedQCalendar(uintptr(h)) }

func (this *QCalendar) UnsafeRelease() { this.h = 0 }

func NewQCalendar() *QCalendar {
	__r, _, _ := purego.SyscallN(addr_QCalendar_new)
	__o := newownedQCalendar(__r)
	return __o
}

func NewQCalendar2(system QCalendar__System) *QCalendar {
	__r, _, _ := purego.SyscallN(addr_QCalendar_new2, uintptr(system))
	__o := newownedQCalendar(__r)
	return __o
}

func NewQCalendar3(name string) *QCalendar {
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	__r, _, _ := purego.SyscallN(addr_QCalendar_new3, name_p, uintptr(len(name)))
	__o := newownedQCalendar(__r)
	runtime.KeepAlive(name)
	return __o
}

func NewQCalendar4(id *QCalendar__SystemId) *QCalendar {
	__r, _, _ := purego.SyscallN(addr_QCalendar_new4, uintptr(id.UnsafePointer()))
	__o := newownedQCalendar(__r)
	runtime.KeepAlive(id)
	return __o
}

func (this *QCalendar) IsValid() bool {
	__r0, _, _ := purego.SyscallN(addr_QCalendar_isValid, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCalendar) DaysInMonth(month int) int {
	__r0, _, _ := purego.SyscallN(addr_QCalendar_daysInMonth, this.h, uintptr(month))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QCalendar) DaysInYear(year int) int {
	__r0, _, _ := purego.SyscallN(addr_QCalendar_daysInYear, this.h, uintptr(year))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QCalendar) MonthsInYear(year int) int {
	__r0, _, _ := purego.SyscallN(addr_QCalendar_monthsInYear, this.h, uintptr(year))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QCalendar) IsDateValid(year int, month int, day int) bool {
	__r0, _, _ := purego.SyscallN(addr_QCalendar_isDateValid, this.h, uintptr(year), uintptr(month), uintptr(day))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCalendar) IsLeapYear(year int) bool {
	__r0, _, _ := purego.SyscallN(addr_QCalendar_isLeapYear, this.h, uintptr(year))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCalendar) IsGregorian() bool {
	__r0, _, _ := purego.SyscallN(addr_QCalendar_isGregorian, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCalendar) IsLunar() bool {
	__r0, _, _ := purego.SyscallN(addr_QCalendar_isLunar, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCalendar) IsLuniSolar() bool {
	__r0, _, _ := purego.SyscallN(addr_QCalendar_isLuniSolar, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCalendar) IsSolar() bool {
	__r0, _, _ := purego.SyscallN(addr_QCalendar_isSolar, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCalendar) IsProleptic() bool {
	__r0, _, _ := purego.SyscallN(addr_QCalendar_isProleptic, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCalendar) HasYearZero() bool {
	__r0, _, _ := purego.SyscallN(addr_QCalendar_hasYearZero, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCalendar) MaximumDaysInMonth() int {
	__r0, _, _ := purego.SyscallN(addr_QCalendar_maximumDaysInMonth, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QCalendar) MinimumDaysInMonth() int {
	__r0, _, _ := purego.SyscallN(addr_QCalendar_minimumDaysInMonth, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QCalendar) MaximumMonthsInYear() int {
	__r0, _, _ := purego.SyscallN(addr_QCalendar_maximumMonthsInYear, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QCalendar) Name() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QCalendar_name, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QCalendar) DateFromParts(year int, month int, day int) *QDate {
	__r, _, _ := purego.SyscallN(addr_QCalendar_dateFromParts, this.h, uintptr(year), uintptr(month), uintptr(day))
	runtime.KeepAlive(this)
	return newownedQDate(__r)
}

func (this *QCalendar) DateFromPartsWithParts(parts *QCalendar__YearMonthDay) *QDate {
	__r, _, _ := purego.SyscallN(addr_QCalendar_dateFromParts2, this.h, uintptr(parts.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parts)
	return newownedQDate(__r)
}

func (this *QCalendar) MatchCenturyToWeekday(parts *QCalendar__YearMonthDay, dow int) *QDate {
	__r, _, _ := purego.SyscallN(addr_QCalendar_matchCenturyToWeekday, this.h, uintptr(parts.UnsafePointer()), uintptr(dow))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parts)
	return newownedQDate(__r)
}

func (this *QCalendar) PartsFromDate(date *QDate) *QCalendar__YearMonthDay {
	__r, _, _ := purego.SyscallN(addr_QCalendar_partsFromDate, this.h, uintptr(date.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(date)
	return newownedQCalendar__YearMonthDay(__r)
}

func (this *QCalendar) DayOfWeek(date *QDate) int {
	__r0, _, _ := purego.SyscallN(addr_QCalendar_dayOfWeek, this.h, uintptr(date.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(date)
	return int(__r)
}

func (this *QCalendar) MonthName(locale *QLocale, month int) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QCalendar_monthName, this.h, uintptr(locale.UnsafePointer()), uintptr(month), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(locale)
	return pgStr(__r, __n)
}

func (this *QCalendar) StandaloneMonthName(locale *QLocale, month int) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QCalendar_standaloneMonthName, this.h, uintptr(locale.UnsafePointer()), uintptr(month), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(locale)
	return pgStr(__r, __n)
}

func (this *QCalendar) WeekDayName(locale *QLocale, day int) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QCalendar_weekDayName, this.h, uintptr(locale.UnsafePointer()), uintptr(day), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(locale)
	return pgStr(__r, __n)
}

func (this *QCalendar) StandaloneWeekDayName(locale *QLocale, day int) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QCalendar_standaloneWeekDayName, this.h, uintptr(locale.UnsafePointer()), uintptr(day), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(locale)
	return pgStr(__r, __n)
}

func (this *QCalendar) DateTimeToString(format *QStringView, datetime *QDateTime, dateOnly *QDate, timeOnly *QTime, locale *QLocale) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QCalendar_dateTimeToString, this.h, uintptr(format.UnsafePointer()), uintptr(datetime.UnsafePointer()), uintptr(dateOnly.UnsafePointer()), uintptr(timeOnly.UnsafePointer()), uintptr(locale.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(format)
	runtime.KeepAlive(datetime)
	runtime.KeepAlive(dateOnly)
	runtime.KeepAlive(timeOnly)
	runtime.KeepAlive(locale)
	return pgStr(__r, __n)
}

func QCalendar_AvailableCalendars() []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QCalendar_availableCalendars, uintptr(unsafe.Pointer(&__n)))
	__out := make([]string, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = pgGoStr(__sl[__i])
		pgFree(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QCalendar) DaysInMonth2(month int, year int) int {
	__r0, _, _ := purego.SyscallN(addr_QCalendar_daysInMonth2, this.h, uintptr(month), uintptr(year))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QCalendar) MonthName2(locale *QLocale, month int, year int) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QCalendar_monthName2, this.h, uintptr(locale.UnsafePointer()), uintptr(month), uintptr(year), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(locale)
	return pgStr(__r, __n)
}

func (this *QCalendar) MonthName3(locale *QLocale, month int, year int, format QLocale__FormatType) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QCalendar_monthName3, this.h, uintptr(locale.UnsafePointer()), uintptr(month), uintptr(year), uintptr(format), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(locale)
	return pgStr(__r, __n)
}

func (this *QCalendar) StandaloneMonthName2(locale *QLocale, month int, year int) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QCalendar_standaloneMonthName2, this.h, uintptr(locale.UnsafePointer()), uintptr(month), uintptr(year), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(locale)
	return pgStr(__r, __n)
}

func (this *QCalendar) StandaloneMonthName3(locale *QLocale, month int, year int, format QLocale__FormatType) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QCalendar_standaloneMonthName3, this.h, uintptr(locale.UnsafePointer()), uintptr(month), uintptr(year), uintptr(format), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(locale)
	return pgStr(__r, __n)
}

func (this *QCalendar) WeekDayName2(locale *QLocale, day int, format QLocale__FormatType) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QCalendar_weekDayName2, this.h, uintptr(locale.UnsafePointer()), uintptr(day), uintptr(format), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(locale)
	return pgStr(__r, __n)
}

func (this *QCalendar) StandaloneWeekDayName2(locale *QLocale, day int, format QLocale__FormatType) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QCalendar_standaloneWeekDayName2, this.h, uintptr(locale.UnsafePointer()), uintptr(day), uintptr(format), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(locale)
	return pgStr(__r, __n)
}

func (this *QCalendar) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QCalendar_delete(this.h)
	}
	this.UnsafeRelease()
}
