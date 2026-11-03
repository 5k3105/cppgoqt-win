package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QCalendar__YearMonthDay_new uintptr
	addr_QCalendar__YearMonthDay_new2 uintptr
	addr_QCalendar__YearMonthDay_new3 uintptr
	addr_QCalendar__YearMonthDay_new4 uintptr
	addr_QCalendar__YearMonthDay_isValid uintptr
	addr_QCalendar__YearMonthDay_year_get uintptr
	addr_QCalendar__YearMonthDay_year_set uintptr
	addr_QCalendar__YearMonthDay_month_get uintptr
	addr_QCalendar__YearMonthDay_month_set uintptr
	addr_QCalendar__YearMonthDay_day_get uintptr
	addr_QCalendar__YearMonthDay_day_set uintptr
	pg_QCalendar__YearMonthDay_delete func(uintptr)
)

func init() {
	symaddr(&addr_QCalendar__YearMonthDay_new, "QCalendar__YearMonthDay_new")
	symaddr(&addr_QCalendar__YearMonthDay_new2, "QCalendar__YearMonthDay_new2")
	symaddr(&addr_QCalendar__YearMonthDay_new3, "QCalendar__YearMonthDay_new3")
	symaddr(&addr_QCalendar__YearMonthDay_new4, "QCalendar__YearMonthDay_new4")
	symaddr(&addr_QCalendar__YearMonthDay_isValid, "QCalendar__YearMonthDay_isValid")
	symaddr(&addr_QCalendar__YearMonthDay_year_get, "QCalendar__YearMonthDay_year_get")
	symaddr(&addr_QCalendar__YearMonthDay_year_set, "QCalendar__YearMonthDay_year_set")
	symaddr(&addr_QCalendar__YearMonthDay_month_get, "QCalendar__YearMonthDay_month_get")
	symaddr(&addr_QCalendar__YearMonthDay_month_set, "QCalendar__YearMonthDay_month_set")
	symaddr(&addr_QCalendar__YearMonthDay_day_get, "QCalendar__YearMonthDay_day_get")
	symaddr(&addr_QCalendar__YearMonthDay_day_set, "QCalendar__YearMonthDay_day_set")
	reg(&pg_QCalendar__YearMonthDay_delete, "QCalendar__YearMonthDay_delete")
}

type QCalendar__YearMonthDay struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QCalendar__YearMonthDay) UnsafeGoID() int64     { return this.vid }
func (this *QCalendar__YearMonthDay) UnsafeSetGoID(id int64) { this.vid = id }

func bindQCalendar__YearMonthDay(h uintptr) *QCalendar__YearMonthDay {
	if h == 0 {
		return nil
	}
	return &QCalendar__YearMonthDay{h: h}
}

func newownedQCalendar__YearMonthDay(h uintptr) *QCalendar__YearMonthDay {
	o := bindQCalendar__YearMonthDay(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QCalendar__YearMonthDay).Delete)
	}
	return o
}

func (this *QCalendar__YearMonthDay) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QCalendar__YearMonthDay) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQCalendar__YearMonthDay(h unsafe.Pointer) *QCalendar__YearMonthDay { return bindQCalendar__YearMonthDay(uintptr(h)) }
func UnsafeNewQCalendar__YearMonthDayOwned(h unsafe.Pointer) *QCalendar__YearMonthDay { return newownedQCalendar__YearMonthDay(uintptr(h)) }

func (this *QCalendar__YearMonthDay) UnsafeRelease() { this.h = 0 }

func NewQCalendar__YearMonthDay() *QCalendar__YearMonthDay {
	__r, _, _ := purego.SyscallN(addr_QCalendar__YearMonthDay_new)
	__o := newownedQCalendar__YearMonthDay(__r)
	return __o
}

func NewQCalendar__YearMonthDay2(y int) *QCalendar__YearMonthDay {
	__r, _, _ := purego.SyscallN(addr_QCalendar__YearMonthDay_new2, uintptr(y))
	__o := newownedQCalendar__YearMonthDay(__r)
	return __o
}

func NewQCalendar__YearMonthDay3(y int, m int) *QCalendar__YearMonthDay {
	__r, _, _ := purego.SyscallN(addr_QCalendar__YearMonthDay_new3, uintptr(y), uintptr(m))
	__o := newownedQCalendar__YearMonthDay(__r)
	return __o
}

func NewQCalendar__YearMonthDay4(y int, m int, d int) *QCalendar__YearMonthDay {
	__r, _, _ := purego.SyscallN(addr_QCalendar__YearMonthDay_new4, uintptr(y), uintptr(m), uintptr(d))
	__o := newownedQCalendar__YearMonthDay(__r)
	return __o
}

func (this *QCalendar__YearMonthDay) IsValid() bool {
	__r0, _, _ := purego.SyscallN(addr_QCalendar__YearMonthDay_isValid, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCalendar__YearMonthDay) Year() int {
	__r0, _, _ := purego.SyscallN(addr_QCalendar__YearMonthDay_year_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QCalendar__YearMonthDay) SetYear(year int) {
	purego.SyscallN(addr_QCalendar__YearMonthDay_year_set, this.h, uintptr(year))
	runtime.KeepAlive(this)
}

func (this *QCalendar__YearMonthDay) Month() int {
	__r0, _, _ := purego.SyscallN(addr_QCalendar__YearMonthDay_month_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QCalendar__YearMonthDay) SetMonth(month int) {
	purego.SyscallN(addr_QCalendar__YearMonthDay_month_set, this.h, uintptr(month))
	runtime.KeepAlive(this)
}

func (this *QCalendar__YearMonthDay) Day() int {
	__r0, _, _ := purego.SyscallN(addr_QCalendar__YearMonthDay_day_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QCalendar__YearMonthDay) SetDay(day int) {
	purego.SyscallN(addr_QCalendar__YearMonthDay_day_set, this.h, uintptr(day))
	runtime.KeepAlive(this)
}

func (this *QCalendar__YearMonthDay) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QCalendar__YearMonthDay_delete(this.h)
	}
	this.UnsafeRelease()
}
