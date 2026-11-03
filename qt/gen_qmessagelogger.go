package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QMessageLogger_new uintptr
	addr_QMessageLogger_new2 uintptr
	addr_QMessageLogger_new3 uintptr
	addr_QMessageLogger_debug uintptr
	addr_QMessageLogger_noDebug uintptr
	addr_QMessageLogger_info uintptr
	addr_QMessageLogger_warning uintptr
	addr_QMessageLogger_critical uintptr
	addr_QMessageLogger_fatal uintptr
	addr_QMessageLogger_debug2 uintptr
	addr_QMessageLogger_info2 uintptr
	addr_QMessageLogger_warning2 uintptr
	addr_QMessageLogger_critical2 uintptr
	addr_QMessageLogger_fatal2 uintptr
	addr_QMessageLogger_debug3 uintptr
	addr_QMessageLogger_debug4 uintptr
	addr_QMessageLogger_info3 uintptr
	addr_QMessageLogger_info4 uintptr
	addr_QMessageLogger_warning3 uintptr
	addr_QMessageLogger_warning4 uintptr
	addr_QMessageLogger_critical3 uintptr
	addr_QMessageLogger_critical4 uintptr
	addr_QMessageLogger_fatal3 uintptr
	addr_QMessageLogger_fatal4 uintptr
	addr_QMessageLogger_noDebug2 uintptr
	pg_QMessageLogger_delete func(uintptr)
)

func init() {
	symaddr(&addr_QMessageLogger_new, "QMessageLogger_new")
	symaddr(&addr_QMessageLogger_new2, "QMessageLogger_new2")
	symaddr(&addr_QMessageLogger_new3, "QMessageLogger_new3")
	symaddr(&addr_QMessageLogger_debug, "QMessageLogger_debug")
	symaddr(&addr_QMessageLogger_noDebug, "QMessageLogger_noDebug")
	symaddr(&addr_QMessageLogger_info, "QMessageLogger_info")
	symaddr(&addr_QMessageLogger_warning, "QMessageLogger_warning")
	symaddr(&addr_QMessageLogger_critical, "QMessageLogger_critical")
	symaddr(&addr_QMessageLogger_fatal, "QMessageLogger_fatal")
	symaddr(&addr_QMessageLogger_debug2, "QMessageLogger_debug2")
	symaddr(&addr_QMessageLogger_info2, "QMessageLogger_info2")
	symaddr(&addr_QMessageLogger_warning2, "QMessageLogger_warning2")
	symaddr(&addr_QMessageLogger_critical2, "QMessageLogger_critical2")
	symaddr(&addr_QMessageLogger_fatal2, "QMessageLogger_fatal2")
	symaddr(&addr_QMessageLogger_debug3, "QMessageLogger_debug3")
	symaddr(&addr_QMessageLogger_debug4, "QMessageLogger_debug4")
	symaddr(&addr_QMessageLogger_info3, "QMessageLogger_info3")
	symaddr(&addr_QMessageLogger_info4, "QMessageLogger_info4")
	symaddr(&addr_QMessageLogger_warning3, "QMessageLogger_warning3")
	symaddr(&addr_QMessageLogger_warning4, "QMessageLogger_warning4")
	symaddr(&addr_QMessageLogger_critical3, "QMessageLogger_critical3")
	symaddr(&addr_QMessageLogger_critical4, "QMessageLogger_critical4")
	symaddr(&addr_QMessageLogger_fatal3, "QMessageLogger_fatal3")
	symaddr(&addr_QMessageLogger_fatal4, "QMessageLogger_fatal4")
	symaddr(&addr_QMessageLogger_noDebug2, "QMessageLogger_noDebug2")
	reg(&pg_QMessageLogger_delete, "QMessageLogger_delete")
}

type QMessageLogger struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QMessageLogger) UnsafeGoID() int64     { return this.vid }
func (this *QMessageLogger) UnsafeSetGoID(id int64) { this.vid = id }

func bindQMessageLogger(h uintptr) *QMessageLogger {
	if h == 0 {
		return nil
	}
	return &QMessageLogger{h: h}
}

func newownedQMessageLogger(h uintptr) *QMessageLogger {
	o := bindQMessageLogger(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QMessageLogger).Delete)
	}
	return o
}

func (this *QMessageLogger) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QMessageLogger) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQMessageLogger(h unsafe.Pointer) *QMessageLogger { return bindQMessageLogger(uintptr(h)) }
func UnsafeNewQMessageLoggerOwned(h unsafe.Pointer) *QMessageLogger { return newownedQMessageLogger(uintptr(h)) }

func (this *QMessageLogger) UnsafeRelease() { this.h = 0 }

func NewQMessageLogger() *QMessageLogger {
	__r, _, _ := purego.SyscallN(addr_QMessageLogger_new)
	__o := newownedQMessageLogger(__r)
	return __o
}

func NewQMessageLogger2(file string, line int, function string) *QMessageLogger {
	file_b := append([]byte(file), 0)
	function_b := append([]byte(function), 0)
	__r, _, _ := purego.SyscallN(addr_QMessageLogger_new2, uintptr(unsafe.Pointer(&file_b[0])), uintptr(line), uintptr(unsafe.Pointer(&function_b[0])))
	__o := newownedQMessageLogger(__r)
	runtime.KeepAlive(file_b)
	runtime.KeepAlive(function_b)
	return __o
}

func NewQMessageLogger3(file string, line int, function string, category string) *QMessageLogger {
	file_b := append([]byte(file), 0)
	function_b := append([]byte(function), 0)
	category_b := append([]byte(category), 0)
	__r, _, _ := purego.SyscallN(addr_QMessageLogger_new3, uintptr(unsafe.Pointer(&file_b[0])), uintptr(line), uintptr(unsafe.Pointer(&function_b[0])), uintptr(unsafe.Pointer(&category_b[0])))
	__o := newownedQMessageLogger(__r)
	runtime.KeepAlive(file_b)
	runtime.KeepAlive(function_b)
	runtime.KeepAlive(category_b)
	return __o
}

func (this *QMessageLogger) Debug(msg string) {
	msg_b := append([]byte(msg), 0)
	purego.SyscallN(addr_QMessageLogger_debug, this.h, uintptr(unsafe.Pointer(&msg_b[0])))
	runtime.KeepAlive(this)
	runtime.KeepAlive(msg_b)
}

func (this *QMessageLogger) NoDebug(arg1 string) {
	arg1_b := append([]byte(arg1), 0)
	purego.SyscallN(addr_QMessageLogger_noDebug, this.h, uintptr(unsafe.Pointer(&arg1_b[0])))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1_b)
}

func (this *QMessageLogger) Info(msg string) {
	msg_b := append([]byte(msg), 0)
	purego.SyscallN(addr_QMessageLogger_info, this.h, uintptr(unsafe.Pointer(&msg_b[0])))
	runtime.KeepAlive(this)
	runtime.KeepAlive(msg_b)
}

func (this *QMessageLogger) Warning(msg string) {
	msg_b := append([]byte(msg), 0)
	purego.SyscallN(addr_QMessageLogger_warning, this.h, uintptr(unsafe.Pointer(&msg_b[0])))
	runtime.KeepAlive(this)
	runtime.KeepAlive(msg_b)
}

func (this *QMessageLogger) Critical(msg string) {
	msg_b := append([]byte(msg), 0)
	purego.SyscallN(addr_QMessageLogger_critical, this.h, uintptr(unsafe.Pointer(&msg_b[0])))
	runtime.KeepAlive(this)
	runtime.KeepAlive(msg_b)
}

func (this *QMessageLogger) Fatal(msg string) {
	msg_b := append([]byte(msg), 0)
	purego.SyscallN(addr_QMessageLogger_fatal, this.h, uintptr(unsafe.Pointer(&msg_b[0])))
	runtime.KeepAlive(this)
	runtime.KeepAlive(msg_b)
}

func (this *QMessageLogger) Debug2(cat *QLoggingCategory, msg string) {
	msg_b := append([]byte(msg), 0)
	purego.SyscallN(addr_QMessageLogger_debug2, this.h, uintptr(cat.UnsafePointer()), uintptr(unsafe.Pointer(&msg_b[0])))
	runtime.KeepAlive(this)
	runtime.KeepAlive(cat)
	runtime.KeepAlive(msg_b)
}

func (this *QMessageLogger) Info2(cat *QLoggingCategory, msg string) {
	msg_b := append([]byte(msg), 0)
	purego.SyscallN(addr_QMessageLogger_info2, this.h, uintptr(cat.UnsafePointer()), uintptr(unsafe.Pointer(&msg_b[0])))
	runtime.KeepAlive(this)
	runtime.KeepAlive(cat)
	runtime.KeepAlive(msg_b)
}

func (this *QMessageLogger) Warning2(cat *QLoggingCategory, msg string) {
	msg_b := append([]byte(msg), 0)
	purego.SyscallN(addr_QMessageLogger_warning2, this.h, uintptr(cat.UnsafePointer()), uintptr(unsafe.Pointer(&msg_b[0])))
	runtime.KeepAlive(this)
	runtime.KeepAlive(cat)
	runtime.KeepAlive(msg_b)
}

func (this *QMessageLogger) Critical2(cat *QLoggingCategory, msg string) {
	msg_b := append([]byte(msg), 0)
	purego.SyscallN(addr_QMessageLogger_critical2, this.h, uintptr(cat.UnsafePointer()), uintptr(unsafe.Pointer(&msg_b[0])))
	runtime.KeepAlive(this)
	runtime.KeepAlive(cat)
	runtime.KeepAlive(msg_b)
}

func (this *QMessageLogger) Fatal2(cat *QLoggingCategory, msg string) {
	msg_b := append([]byte(msg), 0)
	purego.SyscallN(addr_QMessageLogger_fatal2, this.h, uintptr(cat.UnsafePointer()), uintptr(unsafe.Pointer(&msg_b[0])))
	runtime.KeepAlive(this)
	runtime.KeepAlive(cat)
	runtime.KeepAlive(msg_b)
}

func (this *QMessageLogger) Debug3() *QDebug {
	__r, _, _ := purego.SyscallN(addr_QMessageLogger_debug3, this.h)
	runtime.KeepAlive(this)
	return newownedQDebug(__r)
}

func (this *QMessageLogger) DebugWithCat(cat *QLoggingCategory) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QMessageLogger_debug4, this.h, uintptr(cat.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(cat)
	return newownedQDebug(__r)
}

func (this *QMessageLogger) Info3() *QDebug {
	__r, _, _ := purego.SyscallN(addr_QMessageLogger_info3, this.h)
	runtime.KeepAlive(this)
	return newownedQDebug(__r)
}

func (this *QMessageLogger) InfoWithCat(cat *QLoggingCategory) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QMessageLogger_info4, this.h, uintptr(cat.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(cat)
	return newownedQDebug(__r)
}

func (this *QMessageLogger) Warning3() *QDebug {
	__r, _, _ := purego.SyscallN(addr_QMessageLogger_warning3, this.h)
	runtime.KeepAlive(this)
	return newownedQDebug(__r)
}

func (this *QMessageLogger) WarningWithCat(cat *QLoggingCategory) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QMessageLogger_warning4, this.h, uintptr(cat.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(cat)
	return newownedQDebug(__r)
}

func (this *QMessageLogger) Critical3() *QDebug {
	__r, _, _ := purego.SyscallN(addr_QMessageLogger_critical3, this.h)
	runtime.KeepAlive(this)
	return newownedQDebug(__r)
}

func (this *QMessageLogger) CriticalWithCat(cat *QLoggingCategory) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QMessageLogger_critical4, this.h, uintptr(cat.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(cat)
	return newownedQDebug(__r)
}

func (this *QMessageLogger) Fatal3() *QDebug {
	__r, _, _ := purego.SyscallN(addr_QMessageLogger_fatal3, this.h)
	runtime.KeepAlive(this)
	return newownedQDebug(__r)
}

func (this *QMessageLogger) FatalWithCat(cat *QLoggingCategory) *QDebug {
	__r, _, _ := purego.SyscallN(addr_QMessageLogger_fatal4, this.h, uintptr(cat.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(cat)
	return newownedQDebug(__r)
}

func (this *QMessageLogger) NoDebug2() *QNoDebug {
	__r, _, _ := purego.SyscallN(addr_QMessageLogger_noDebug2, this.h)
	runtime.KeepAlive(this)
	return newownedQNoDebug(__r)
}

func (this *QMessageLogger) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QMessageLogger_delete(this.h)
	}
	this.UnsafeRelease()
}
