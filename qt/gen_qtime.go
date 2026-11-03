package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QTime_new uintptr
	addr_QTime_new2 uintptr
	addr_QTime_new3 uintptr
	addr_QTime_new4 uintptr
	addr_QTime_new5 uintptr
	addr_QTime_isNull uintptr
	addr_QTime_isValid uintptr
	addr_QTime_hour uintptr
	addr_QTime_minute uintptr
	addr_QTime_second uintptr
	addr_QTime_msec uintptr
	addr_QTime_toString uintptr
	addr_QTime_toString2 uintptr
	addr_QTime_toString3 uintptr
	addr_QTime_setHMS uintptr
	addr_QTime_addSecs uintptr
	addr_QTime_secsTo uintptr
	addr_QTime_addMSecs uintptr
	addr_QTime_msecsTo uintptr
	addr_QTime_fromMSecsSinceStartOfDay uintptr
	addr_QTime_msecsSinceStartOfDay uintptr
	addr_QTime_currentTime uintptr
	addr_QTime_fromString uintptr
	addr_QTime_fromString2 uintptr
	addr_QTime_fromString3 uintptr
	addr_QTime_fromString4 uintptr
	addr_QTime_fromString5 uintptr
	addr_QTime_isValid2 uintptr
	addr_QTime_toString4 uintptr
	addr_QTime_setHMS2 uintptr
	addr_QTime_fromString6 uintptr
	addr_QTime_fromString7 uintptr
	addr_QTime_isValid3 uintptr
	addr_QTime_OperatorEqual uintptr
	addr_QTime_OperatorNotEqual uintptr
	addr_QTime_OperatorLesser uintptr
	addr_QTime_OperatorGreater uintptr
	addr_QTime_OperatorLesserOrEqual uintptr
	addr_QTime_OperatorGreaterOrEqual uintptr
	pg_QTime_delete func(uintptr)
)

func init() {
	symaddr(&addr_QTime_new, "QTime_new")
	symaddr(&addr_QTime_new2, "QTime_new2")
	symaddr(&addr_QTime_new3, "QTime_new3")
	symaddr(&addr_QTime_new4, "QTime_new4")
	symaddr(&addr_QTime_new5, "QTime_new5")
	symaddr(&addr_QTime_isNull, "QTime_isNull")
	symaddr(&addr_QTime_isValid, "QTime_isValid")
	symaddr(&addr_QTime_hour, "QTime_hour")
	symaddr(&addr_QTime_minute, "QTime_minute")
	symaddr(&addr_QTime_second, "QTime_second")
	symaddr(&addr_QTime_msec, "QTime_msec")
	symaddr(&addr_QTime_toString, "QTime_toString")
	symaddr(&addr_QTime_toString2, "QTime_toString2")
	symaddr(&addr_QTime_toString3, "QTime_toString3")
	symaddr(&addr_QTime_setHMS, "QTime_setHMS")
	symaddr(&addr_QTime_addSecs, "QTime_addSecs")
	symaddr(&addr_QTime_secsTo, "QTime_secsTo")
	symaddr(&addr_QTime_addMSecs, "QTime_addMSecs")
	symaddr(&addr_QTime_msecsTo, "QTime_msecsTo")
	symaddr(&addr_QTime_fromMSecsSinceStartOfDay, "QTime_fromMSecsSinceStartOfDay")
	symaddr(&addr_QTime_msecsSinceStartOfDay, "QTime_msecsSinceStartOfDay")
	symaddr(&addr_QTime_currentTime, "QTime_currentTime")
	symaddr(&addr_QTime_fromString, "QTime_fromString")
	symaddr(&addr_QTime_fromString2, "QTime_fromString2")
	symaddr(&addr_QTime_fromString3, "QTime_fromString3")
	symaddr(&addr_QTime_fromString4, "QTime_fromString4")
	symaddr(&addr_QTime_fromString5, "QTime_fromString5")
	symaddr(&addr_QTime_isValid2, "QTime_isValid2")
	symaddr(&addr_QTime_toString4, "QTime_toString4")
	symaddr(&addr_QTime_setHMS2, "QTime_setHMS2")
	symaddr(&addr_QTime_fromString6, "QTime_fromString6")
	symaddr(&addr_QTime_fromString7, "QTime_fromString7")
	symaddr(&addr_QTime_isValid3, "QTime_isValid3")
	symaddr(&addr_QTime_OperatorEqual, "QTime_OperatorEqual")
	symaddr(&addr_QTime_OperatorNotEqual, "QTime_OperatorNotEqual")
	symaddr(&addr_QTime_OperatorLesser, "QTime_OperatorLesser")
	symaddr(&addr_QTime_OperatorGreater, "QTime_OperatorGreater")
	symaddr(&addr_QTime_OperatorLesserOrEqual, "QTime_OperatorLesserOrEqual")
	symaddr(&addr_QTime_OperatorGreaterOrEqual, "QTime_OperatorGreaterOrEqual")
	reg(&pg_QTime_delete, "QTime_delete")
}

type QTime struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QTime) UnsafeGoID() int64     { return this.vid }
func (this *QTime) UnsafeSetGoID(id int64) { this.vid = id }

func bindQTime(h uintptr) *QTime {
	if h == 0 {
		return nil
	}
	return &QTime{h: h}
}

func newownedQTime(h uintptr) *QTime {
	o := bindQTime(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QTime).Delete)
	}
	return o
}

func (this *QTime) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QTime) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQTime(h unsafe.Pointer) *QTime { return bindQTime(uintptr(h)) }
func UnsafeNewQTimeOwned(h unsafe.Pointer) *QTime { return newownedQTime(uintptr(h)) }

func (this *QTime) UnsafeRelease() { this.h = 0 }

func NewQTime() *QTime {
	__r, _, _ := purego.SyscallN(addr_QTime_new)
	__o := newownedQTime(__r)
	return __o
}

func NewQTime2(h int, m int) *QTime {
	__r, _, _ := purego.SyscallN(addr_QTime_new2, uintptr(h), uintptr(m))
	__o := newownedQTime(__r)
	return __o
}

func NewQTime3(arg1 *QTime) *QTime {
	__r, _, _ := purego.SyscallN(addr_QTime_new3, uintptr(arg1.UnsafePointer()))
	__o := newownedQTime(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func NewQTime4(h int, m int, s int) *QTime {
	__r, _, _ := purego.SyscallN(addr_QTime_new4, uintptr(h), uintptr(m), uintptr(s))
	__o := newownedQTime(__r)
	return __o
}

func NewQTime5(h int, m int, s int, ms int) *QTime {
	__r, _, _ := purego.SyscallN(addr_QTime_new5, uintptr(h), uintptr(m), uintptr(s), uintptr(ms))
	__o := newownedQTime(__r)
	return __o
}

func (this *QTime) IsNull() bool {
	__r0, _, _ := purego.SyscallN(addr_QTime_isNull, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTime) IsValid() bool {
	__r0, _, _ := purego.SyscallN(addr_QTime_isValid, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTime) Hour() int {
	__r0, _, _ := purego.SyscallN(addr_QTime_hour, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTime) Minute() int {
	__r0, _, _ := purego.SyscallN(addr_QTime_minute, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTime) Second() int {
	__r0, _, _ := purego.SyscallN(addr_QTime_second, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTime) Msec() int {
	__r0, _, _ := purego.SyscallN(addr_QTime_msec, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTime) ToString() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTime_toString, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QTime) ToStringWithFormat(format string) string {
	var format_p uintptr
	if len(format) > 0 {
		format_p = uintptr(unsafe.Pointer(unsafe.StringData(format)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTime_toString2, this.h, format_p, uintptr(len(format)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(format)
	return pgStr(__r, __n)
}

func (this *QTime) ToString2(format *QStringView) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTime_toString3, this.h, uintptr(format.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(format)
	return pgStr(__r, __n)
}

func (this *QTime) SetHMS(h int, m int, s int) bool {
	__r0, _, _ := purego.SyscallN(addr_QTime_setHMS, this.h, uintptr(h), uintptr(m), uintptr(s))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTime) AddSecs(secs int) *QTime {
	__r, _, _ := purego.SyscallN(addr_QTime_addSecs, this.h, uintptr(secs))
	runtime.KeepAlive(this)
	return newownedQTime(__r)
}

func (this *QTime) SecsTo(t *QTime) int {
	__r0, _, _ := purego.SyscallN(addr_QTime_secsTo, this.h, uintptr(t.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(t)
	return int(__r)
}

func (this *QTime) AddMSecs(ms int) *QTime {
	__r, _, _ := purego.SyscallN(addr_QTime_addMSecs, this.h, uintptr(ms))
	runtime.KeepAlive(this)
	return newownedQTime(__r)
}

func (this *QTime) MsecsTo(t *QTime) int {
	__r0, _, _ := purego.SyscallN(addr_QTime_msecsTo, this.h, uintptr(t.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(t)
	return int(__r)
}

func QTime_FromMSecsSinceStartOfDay(msecs int) *QTime {
	__r, _, _ := purego.SyscallN(addr_QTime_fromMSecsSinceStartOfDay, uintptr(msecs))
	return newownedQTime(__r)
}

func (this *QTime) MsecsSinceStartOfDay() int {
	__r0, _, _ := purego.SyscallN(addr_QTime_msecsSinceStartOfDay, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func QTime_CurrentTime() *QTime {
	__r, _, _ := purego.SyscallN(addr_QTime_currentTime)
	return newownedQTime(__r)
}

func QTime_FromString(string_ *QStringView) *QTime {
	__r, _, _ := purego.SyscallN(addr_QTime_fromString, uintptr(string_.UnsafePointer()))
	runtime.KeepAlive(string_)
	return newownedQTime(__r)
}

func QTime_FromString2(string_ *QStringView, format *QStringView) *QTime {
	__r, _, _ := purego.SyscallN(addr_QTime_fromString2, uintptr(string_.UnsafePointer()), uintptr(format.UnsafePointer()))
	runtime.KeepAlive(string_)
	runtime.KeepAlive(format)
	return newownedQTime(__r)
}

func QTime_FromString3(string_ string, format *QStringView) *QTime {
	var string__p uintptr
	if len(string_) > 0 {
		string__p = uintptr(unsafe.Pointer(unsafe.StringData(string_)))
	}
	__r, _, _ := purego.SyscallN(addr_QTime_fromString3, string__p, uintptr(len(string_)), uintptr(format.UnsafePointer()))
	runtime.KeepAlive(string_)
	runtime.KeepAlive(format)
	return newownedQTime(__r)
}

func QTime_FromStringWithString(string_ string) *QTime {
	var string__p uintptr
	if len(string_) > 0 {
		string__p = uintptr(unsafe.Pointer(unsafe.StringData(string_)))
	}
	__r, _, _ := purego.SyscallN(addr_QTime_fromString4, string__p, uintptr(len(string_)))
	runtime.KeepAlive(string_)
	return newownedQTime(__r)
}

func QTime_FromString4(string_ string, format string) *QTime {
	var string__p uintptr
	if len(string_) > 0 {
		string__p = uintptr(unsafe.Pointer(unsafe.StringData(string_)))
	}
	var format_p uintptr
	if len(format) > 0 {
		format_p = uintptr(unsafe.Pointer(unsafe.StringData(format)))
	}
	__r, _, _ := purego.SyscallN(addr_QTime_fromString5, string__p, uintptr(len(string_)), format_p, uintptr(len(format)))
	runtime.KeepAlive(string_)
	runtime.KeepAlive(format)
	return newownedQTime(__r)
}

func QTime_IsValid(h int, m int, s int) bool {
	__r0, _, _ := purego.SyscallN(addr_QTime_isValid2, uintptr(h), uintptr(m), uintptr(s))
	__r := int32(uint32(__r0))
	return __r != 0
}

func (this *QTime) ToStringWithQtDateFormat(f DateFormat) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTime_toString4, this.h, uintptr(f), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QTime) SetHMS2(h int, m int, s int, ms int) bool {
	__r0, _, _ := purego.SyscallN(addr_QTime_setHMS2, this.h, uintptr(h), uintptr(m), uintptr(s), uintptr(ms))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QTime_FromString5(string_ *QStringView, format DateFormat) *QTime {
	__r, _, _ := purego.SyscallN(addr_QTime_fromString6, uintptr(string_.UnsafePointer()), uintptr(format))
	runtime.KeepAlive(string_)
	return newownedQTime(__r)
}

func QTime_FromString6(string_ string, format DateFormat) *QTime {
	var string__p uintptr
	if len(string_) > 0 {
		string__p = uintptr(unsafe.Pointer(unsafe.StringData(string_)))
	}
	__r, _, _ := purego.SyscallN(addr_QTime_fromString7, string__p, uintptr(len(string_)), uintptr(format))
	runtime.KeepAlive(string_)
	return newownedQTime(__r)
}

func QTime_IsValid2(h int, m int, s int, ms int) bool {
	__r0, _, _ := purego.SyscallN(addr_QTime_isValid3, uintptr(h), uintptr(m), uintptr(s), uintptr(ms))
	__r := int32(uint32(__r0))
	return __r != 0
}

func (this *QTime) OperatorEqual(rhs *QTime) bool {
	__r0, _, _ := purego.SyscallN(addr_QTime_OperatorEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QTime) OperatorNotEqual(rhs *QTime) bool {
	__r0, _, _ := purego.SyscallN(addr_QTime_OperatorNotEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QTime) OperatorLesser(rhs *QTime) bool {
	__r0, _, _ := purego.SyscallN(addr_QTime_OperatorLesser, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QTime) OperatorGreater(rhs *QTime) bool {
	__r0, _, _ := purego.SyscallN(addr_QTime_OperatorGreater, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QTime) OperatorLesserOrEqual(rhs *QTime) bool {
	__r0, _, _ := purego.SyscallN(addr_QTime_OperatorLesserOrEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QTime) OperatorGreaterOrEqual(rhs *QTime) bool {
	__r0, _, _ := purego.SyscallN(addr_QTime_OperatorGreaterOrEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QTime) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QTime_delete(this.h)
	}
	this.UnsafeRelease()
}
