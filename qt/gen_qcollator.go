package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QCollator_new uintptr
	addr_QCollator_new2 uintptr
	addr_QCollator_new3 uintptr
	addr_QCollator_OperatorAssign uintptr
	addr_QCollator_swap uintptr
	addr_QCollator_setLocale uintptr
	addr_QCollator_locale uintptr
	addr_QCollator_caseSensitivity uintptr
	addr_QCollator_setCaseSensitivity uintptr
	addr_QCollator_setNumericMode uintptr
	addr_QCollator_numericMode uintptr
	addr_QCollator_setIgnorePunctuation uintptr
	addr_QCollator_ignorePunctuation uintptr
	addr_QCollator_compare uintptr
	addr_QCollator_compare2 uintptr
	addr_QCollator_OperatorCall uintptr
	addr_QCollator_compare3 uintptr
	addr_QCollator_OperatorCall2 uintptr
	addr_QCollator_sortKey uintptr
	addr_QCollator_defaultCompare uintptr
	addr_QCollator_defaultSortKey uintptr
	pg_QCollator_delete func(uintptr)
)

func init() {
	symaddr(&addr_QCollator_new, "QCollator_new")
	symaddr(&addr_QCollator_new2, "QCollator_new2")
	symaddr(&addr_QCollator_new3, "QCollator_new3")
	symaddr(&addr_QCollator_OperatorAssign, "QCollator_OperatorAssign")
	symaddr(&addr_QCollator_swap, "QCollator_swap")
	symaddr(&addr_QCollator_setLocale, "QCollator_setLocale")
	symaddr(&addr_QCollator_locale, "QCollator_locale")
	symaddr(&addr_QCollator_caseSensitivity, "QCollator_caseSensitivity")
	symaddr(&addr_QCollator_setCaseSensitivity, "QCollator_setCaseSensitivity")
	symaddr(&addr_QCollator_setNumericMode, "QCollator_setNumericMode")
	symaddr(&addr_QCollator_numericMode, "QCollator_numericMode")
	symaddr(&addr_QCollator_setIgnorePunctuation, "QCollator_setIgnorePunctuation")
	symaddr(&addr_QCollator_ignorePunctuation, "QCollator_ignorePunctuation")
	symaddr(&addr_QCollator_compare, "QCollator_compare")
	symaddr(&addr_QCollator_compare2, "QCollator_compare2")
	symaddr(&addr_QCollator_OperatorCall, "QCollator_OperatorCall")
	symaddr(&addr_QCollator_compare3, "QCollator_compare3")
	symaddr(&addr_QCollator_OperatorCall2, "QCollator_OperatorCall2")
	symaddr(&addr_QCollator_sortKey, "QCollator_sortKey")
	symaddr(&addr_QCollator_defaultCompare, "QCollator_defaultCompare")
	symaddr(&addr_QCollator_defaultSortKey, "QCollator_defaultSortKey")
	reg(&pg_QCollator_delete, "QCollator_delete")
}

type QCollator struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QCollator) UnsafeGoID() int64     { return this.vid }
func (this *QCollator) UnsafeSetGoID(id int64) { this.vid = id }

func bindQCollator(h uintptr) *QCollator {
	if h == 0 {
		return nil
	}
	return &QCollator{h: h}
}

func newownedQCollator(h uintptr) *QCollator {
	o := bindQCollator(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QCollator).Delete)
	}
	return o
}

func (this *QCollator) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QCollator) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQCollator(h unsafe.Pointer) *QCollator { return bindQCollator(uintptr(h)) }
func UnsafeNewQCollatorOwned(h unsafe.Pointer) *QCollator { return newownedQCollator(uintptr(h)) }

func (this *QCollator) UnsafeRelease() { this.h = 0 }

func NewQCollator() *QCollator {
	__r, _, _ := purego.SyscallN(addr_QCollator_new)
	__o := newownedQCollator(__r)
	return __o
}

func NewQCollator2(locale *QLocale) *QCollator {
	__r, _, _ := purego.SyscallN(addr_QCollator_new2, uintptr(locale.UnsafePointer()))
	__o := newownedQCollator(__r)
	runtime.KeepAlive(locale)
	return __o
}

func NewQCollator3(arg1 *QCollator) *QCollator {
	__r, _, _ := purego.SyscallN(addr_QCollator_new3, uintptr(arg1.UnsafePointer()))
	__o := newownedQCollator(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QCollator) OperatorAssign(arg1 *QCollator) *QCollator {
	__r, _, _ := purego.SyscallN(addr_QCollator_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQCollator(__r)
}

func (this *QCollator) Swap(other *QCollator) {
	purego.SyscallN(addr_QCollator_swap, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
}

func (this *QCollator) SetLocale(locale *QLocale) {
	purego.SyscallN(addr_QCollator_setLocale, this.h, uintptr(locale.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(locale)
}

func (this *QCollator) Locale() *QLocale {
	__r, _, _ := purego.SyscallN(addr_QCollator_locale, this.h)
	runtime.KeepAlive(this)
	return newownedQLocale(__r)
}

func (this *QCollator) CaseSensitivity() CaseSensitivity {
	__r0, _, _ := purego.SyscallN(addr_QCollator_caseSensitivity, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return CaseSensitivity(__r)
}

func (this *QCollator) SetCaseSensitivity(cs CaseSensitivity) {
	purego.SyscallN(addr_QCollator_setCaseSensitivity, this.h, uintptr(cs))
	runtime.KeepAlive(this)
}

func (this *QCollator) SetNumericMode(on bool) {
	purego.SyscallN(addr_QCollator_setNumericMode, this.h, uintptr(b2i(on)))
	runtime.KeepAlive(this)
}

func (this *QCollator) NumericMode() bool {
	__r0, _, _ := purego.SyscallN(addr_QCollator_numericMode, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCollator) SetIgnorePunctuation(on bool) {
	purego.SyscallN(addr_QCollator_setIgnorePunctuation, this.h, uintptr(b2i(on)))
	runtime.KeepAlive(this)
}

func (this *QCollator) IgnorePunctuation() bool {
	__r0, _, _ := purego.SyscallN(addr_QCollator_ignorePunctuation, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCollator) Compare(s1 string, s2 string) int {
	var s1_p uintptr
	if len(s1) > 0 {
		s1_p = uintptr(unsafe.Pointer(unsafe.StringData(s1)))
	}
	var s2_p uintptr
	if len(s2) > 0 {
		s2_p = uintptr(unsafe.Pointer(unsafe.StringData(s2)))
	}
	__r0, _, _ := purego.SyscallN(addr_QCollator_compare, this.h, s1_p, uintptr(len(s1)), s2_p, uintptr(len(s2)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s1)
	runtime.KeepAlive(s2)
	return int(__r)
}

func (this *QCollator) Compare2(s1 *QChar, len1 int64, s2 *QChar, len2 int64) int {
	__r0, _, _ := purego.SyscallN(addr_QCollator_compare2, this.h, uintptr(s1.UnsafePointer()), uintptr(len1), uintptr(s2.UnsafePointer()), uintptr(len2))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s1)
	runtime.KeepAlive(s2)
	return int(__r)
}

func (this *QCollator) OperatorCall(s1 string, s2 string) bool {
	var s1_p uintptr
	if len(s1) > 0 {
		s1_p = uintptr(unsafe.Pointer(unsafe.StringData(s1)))
	}
	var s2_p uintptr
	if len(s2) > 0 {
		s2_p = uintptr(unsafe.Pointer(unsafe.StringData(s2)))
	}
	__r0, _, _ := purego.SyscallN(addr_QCollator_OperatorCall, this.h, s1_p, uintptr(len(s1)), s2_p, uintptr(len(s2)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s1)
	runtime.KeepAlive(s2)
	return __r != 0
}

func (this *QCollator) Compare3(s1 *QStringView, s2 *QStringView) int {
	__r0, _, _ := purego.SyscallN(addr_QCollator_compare3, this.h, uintptr(s1.UnsafePointer()), uintptr(s2.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s1)
	runtime.KeepAlive(s2)
	return int(__r)
}

func (this *QCollator) OperatorCall2(s1 *QStringView, s2 *QStringView) bool {
	__r0, _, _ := purego.SyscallN(addr_QCollator_OperatorCall2, this.h, uintptr(s1.UnsafePointer()), uintptr(s2.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s1)
	runtime.KeepAlive(s2)
	return __r != 0
}

func (this *QCollator) SortKey(string_ string) *QCollatorSortKey {
	var string__p uintptr
	if len(string_) > 0 {
		string__p = uintptr(unsafe.Pointer(unsafe.StringData(string_)))
	}
	__r, _, _ := purego.SyscallN(addr_QCollator_sortKey, this.h, string__p, uintptr(len(string_)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(string_)
	return newownedQCollatorSortKey(__r)
}

func QCollator_DefaultCompare(s1 *QStringView, s2 *QStringView) int {
	__r0, _, _ := purego.SyscallN(addr_QCollator_defaultCompare, uintptr(s1.UnsafePointer()), uintptr(s2.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(s1)
	runtime.KeepAlive(s2)
	return int(__r)
}

func QCollator_DefaultSortKey(key *QStringView) *QCollatorSortKey {
	__r, _, _ := purego.SyscallN(addr_QCollator_defaultSortKey, uintptr(key.UnsafePointer()))
	runtime.KeepAlive(key)
	return newownedQCollatorSortKey(__r)
}

func (this *QCollator) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QCollator_delete(this.h)
	}
	this.UnsafeRelease()
}
