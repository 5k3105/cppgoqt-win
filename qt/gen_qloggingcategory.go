package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QLoggingCategory_new uintptr
	addr_QLoggingCategory_new2 uintptr
	addr_QLoggingCategory_isEnabled uintptr
	addr_QLoggingCategory_setEnabled uintptr
	addr_QLoggingCategory_isDebugEnabled uintptr
	addr_QLoggingCategory_isInfoEnabled uintptr
	addr_QLoggingCategory_isWarningEnabled uintptr
	addr_QLoggingCategory_isCriticalEnabled uintptr
	addr_QLoggingCategory_categoryName uintptr
	addr_QLoggingCategory_OperatorCall uintptr
	addr_QLoggingCategory_OperatorCall2 uintptr
	addr_QLoggingCategory_defaultCategory uintptr
	addr_QLoggingCategory_setFilterRules uintptr
	pg_QLoggingCategory_delete func(uintptr)
)

func init() {
	symaddr(&addr_QLoggingCategory_new, "QLoggingCategory_new")
	symaddr(&addr_QLoggingCategory_new2, "QLoggingCategory_new2")
	symaddr(&addr_QLoggingCategory_isEnabled, "QLoggingCategory_isEnabled")
	symaddr(&addr_QLoggingCategory_setEnabled, "QLoggingCategory_setEnabled")
	symaddr(&addr_QLoggingCategory_isDebugEnabled, "QLoggingCategory_isDebugEnabled")
	symaddr(&addr_QLoggingCategory_isInfoEnabled, "QLoggingCategory_isInfoEnabled")
	symaddr(&addr_QLoggingCategory_isWarningEnabled, "QLoggingCategory_isWarningEnabled")
	symaddr(&addr_QLoggingCategory_isCriticalEnabled, "QLoggingCategory_isCriticalEnabled")
	symaddr(&addr_QLoggingCategory_categoryName, "QLoggingCategory_categoryName")
	symaddr(&addr_QLoggingCategory_OperatorCall, "QLoggingCategory_OperatorCall")
	symaddr(&addr_QLoggingCategory_OperatorCall2, "QLoggingCategory_OperatorCall2")
	symaddr(&addr_QLoggingCategory_defaultCategory, "QLoggingCategory_defaultCategory")
	symaddr(&addr_QLoggingCategory_setFilterRules, "QLoggingCategory_setFilterRules")
	reg(&pg_QLoggingCategory_delete, "QLoggingCategory_delete")
}

type QLoggingCategory struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QLoggingCategory) UnsafeGoID() int64     { return this.vid }
func (this *QLoggingCategory) UnsafeSetGoID(id int64) { this.vid = id }

func bindQLoggingCategory(h uintptr) *QLoggingCategory {
	if h == 0 {
		return nil
	}
	return &QLoggingCategory{h: h}
}

func newownedQLoggingCategory(h uintptr) *QLoggingCategory {
	o := bindQLoggingCategory(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QLoggingCategory).Delete)
	}
	return o
}

func (this *QLoggingCategory) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QLoggingCategory) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQLoggingCategory(h unsafe.Pointer) *QLoggingCategory { return bindQLoggingCategory(uintptr(h)) }
func UnsafeNewQLoggingCategoryOwned(h unsafe.Pointer) *QLoggingCategory { return newownedQLoggingCategory(uintptr(h)) }

func (this *QLoggingCategory) UnsafeRelease() { this.h = 0 }

func NewQLoggingCategory(category string) *QLoggingCategory {
	category_b := append([]byte(category), 0)
	__r, _, _ := purego.SyscallN(addr_QLoggingCategory_new, uintptr(unsafe.Pointer(&category_b[0])))
	__o := newownedQLoggingCategory(__r)
	runtime.KeepAlive(category_b)
	return __o
}

func NewQLoggingCategory2(category string, severityLevel QtMsgType) *QLoggingCategory {
	category_b := append([]byte(category), 0)
	__r, _, _ := purego.SyscallN(addr_QLoggingCategory_new2, uintptr(unsafe.Pointer(&category_b[0])), uintptr(severityLevel))
	__o := newownedQLoggingCategory(__r)
	runtime.KeepAlive(category_b)
	return __o
}

func (this *QLoggingCategory) IsEnabled(type_ QtMsgType) bool {
	__r0, _, _ := purego.SyscallN(addr_QLoggingCategory_isEnabled, this.h, uintptr(type_))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QLoggingCategory) SetEnabled(type_ QtMsgType, enable bool) {
	purego.SyscallN(addr_QLoggingCategory_setEnabled, this.h, uintptr(type_), uintptr(b2i(enable)))
	runtime.KeepAlive(this)
}

func (this *QLoggingCategory) IsDebugEnabled() bool {
	__r0, _, _ := purego.SyscallN(addr_QLoggingCategory_isDebugEnabled, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QLoggingCategory) IsInfoEnabled() bool {
	__r0, _, _ := purego.SyscallN(addr_QLoggingCategory_isInfoEnabled, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QLoggingCategory) IsWarningEnabled() bool {
	__r0, _, _ := purego.SyscallN(addr_QLoggingCategory_isWarningEnabled, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QLoggingCategory) IsCriticalEnabled() bool {
	__r0, _, _ := purego.SyscallN(addr_QLoggingCategory_isCriticalEnabled, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QLoggingCategory) CategoryName() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLoggingCategory_categoryName, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QLoggingCategory) OperatorCall() *QLoggingCategory {
	__r, _, _ := purego.SyscallN(addr_QLoggingCategory_OperatorCall, this.h)
	runtime.KeepAlive(this)
	return bindQLoggingCategory(__r)
}

func (this *QLoggingCategory) OperatorCall2() *QLoggingCategory {
	__r, _, _ := purego.SyscallN(addr_QLoggingCategory_OperatorCall2, this.h)
	runtime.KeepAlive(this)
	return bindQLoggingCategory(__r)
}

func QLoggingCategory_DefaultCategory() *QLoggingCategory {
	__r, _, _ := purego.SyscallN(addr_QLoggingCategory_defaultCategory)
	return bindQLoggingCategory(__r)
}

func QLoggingCategory_SetFilterRules(rules string) {
	var rules_p uintptr
	if len(rules) > 0 {
		rules_p = uintptr(unsafe.Pointer(unsafe.StringData(rules)))
	}
	purego.SyscallN(addr_QLoggingCategory_setFilterRules, rules_p, uintptr(len(rules)))
	runtime.KeepAlive(rules)
}

func (this *QLoggingCategory) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QLoggingCategory_delete(this.h)
	}
	this.UnsafeRelease()
}
