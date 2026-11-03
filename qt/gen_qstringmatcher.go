package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QStringMatcher_new uintptr
	addr_QStringMatcher_new2 uintptr
	addr_QStringMatcher_new3 uintptr
	addr_QStringMatcher_new4 uintptr
	addr_QStringMatcher_new5 uintptr
	addr_QStringMatcher_new6 uintptr
	addr_QStringMatcher_new7 uintptr
	addr_QStringMatcher_new8 uintptr
	addr_QStringMatcher_OperatorAssign uintptr
	addr_QStringMatcher_setPattern uintptr
	addr_QStringMatcher_setCaseSensitivity uintptr
	addr_QStringMatcher_indexIn uintptr
	addr_QStringMatcher_indexIn2 uintptr
	addr_QStringMatcher_indexIn3 uintptr
	addr_QStringMatcher_pattern uintptr
	addr_QStringMatcher_patternView uintptr
	addr_QStringMatcher_caseSensitivity uintptr
	addr_QStringMatcher_indexIn4 uintptr
	addr_QStringMatcher_indexIn5 uintptr
	addr_QStringMatcher_indexIn6 uintptr
	pg_QStringMatcher_delete func(uintptr)
)

func init() {
	symaddr(&addr_QStringMatcher_new, "QStringMatcher_new")
	symaddr(&addr_QStringMatcher_new2, "QStringMatcher_new2")
	symaddr(&addr_QStringMatcher_new3, "QStringMatcher_new3")
	symaddr(&addr_QStringMatcher_new4, "QStringMatcher_new4")
	symaddr(&addr_QStringMatcher_new5, "QStringMatcher_new5")
	symaddr(&addr_QStringMatcher_new6, "QStringMatcher_new6")
	symaddr(&addr_QStringMatcher_new7, "QStringMatcher_new7")
	symaddr(&addr_QStringMatcher_new8, "QStringMatcher_new8")
	symaddr(&addr_QStringMatcher_OperatorAssign, "QStringMatcher_OperatorAssign")
	symaddr(&addr_QStringMatcher_setPattern, "QStringMatcher_setPattern")
	symaddr(&addr_QStringMatcher_setCaseSensitivity, "QStringMatcher_setCaseSensitivity")
	symaddr(&addr_QStringMatcher_indexIn, "QStringMatcher_indexIn")
	symaddr(&addr_QStringMatcher_indexIn2, "QStringMatcher_indexIn2")
	symaddr(&addr_QStringMatcher_indexIn3, "QStringMatcher_indexIn3")
	symaddr(&addr_QStringMatcher_pattern, "QStringMatcher_pattern")
	symaddr(&addr_QStringMatcher_patternView, "QStringMatcher_patternView")
	symaddr(&addr_QStringMatcher_caseSensitivity, "QStringMatcher_caseSensitivity")
	symaddr(&addr_QStringMatcher_indexIn4, "QStringMatcher_indexIn4")
	symaddr(&addr_QStringMatcher_indexIn5, "QStringMatcher_indexIn5")
	symaddr(&addr_QStringMatcher_indexIn6, "QStringMatcher_indexIn6")
	reg(&pg_QStringMatcher_delete, "QStringMatcher_delete")
}

type QStringMatcher struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QStringMatcher) UnsafeGoID() int64     { return this.vid }
func (this *QStringMatcher) UnsafeSetGoID(id int64) { this.vid = id }

func bindQStringMatcher(h uintptr) *QStringMatcher {
	if h == 0 {
		return nil
	}
	return &QStringMatcher{h: h}
}

func newownedQStringMatcher(h uintptr) *QStringMatcher {
	o := bindQStringMatcher(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QStringMatcher).Delete)
	}
	return o
}

func (this *QStringMatcher) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QStringMatcher) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQStringMatcher(h unsafe.Pointer) *QStringMatcher { return bindQStringMatcher(uintptr(h)) }
func UnsafeNewQStringMatcherOwned(h unsafe.Pointer) *QStringMatcher { return newownedQStringMatcher(uintptr(h)) }

func (this *QStringMatcher) UnsafeRelease() { this.h = 0 }

func NewQStringMatcher() *QStringMatcher {
	__r, _, _ := purego.SyscallN(addr_QStringMatcher_new)
	__o := newownedQStringMatcher(__r)
	return __o
}

func NewQStringMatcher2(pattern string) *QStringMatcher {
	var pattern_p uintptr
	if len(pattern) > 0 {
		pattern_p = uintptr(unsafe.Pointer(unsafe.StringData(pattern)))
	}
	__r, _, _ := purego.SyscallN(addr_QStringMatcher_new2, pattern_p, uintptr(len(pattern)))
	__o := newownedQStringMatcher(__r)
	runtime.KeepAlive(pattern)
	return __o
}

func NewQStringMatcher3(uc *QChar, len_ int64) *QStringMatcher {
	__r, _, _ := purego.SyscallN(addr_QStringMatcher_new3, uintptr(uc.UnsafePointer()), uintptr(len_))
	__o := newownedQStringMatcher(__r)
	runtime.KeepAlive(uc)
	return __o
}

func NewQStringMatcher4(pattern *QStringView) *QStringMatcher {
	__r, _, _ := purego.SyscallN(addr_QStringMatcher_new4, uintptr(pattern.UnsafePointer()))
	__o := newownedQStringMatcher(__r)
	runtime.KeepAlive(pattern)
	return __o
}

func NewQStringMatcher5(other *QStringMatcher) *QStringMatcher {
	__r, _, _ := purego.SyscallN(addr_QStringMatcher_new5, uintptr(other.UnsafePointer()))
	__o := newownedQStringMatcher(__r)
	runtime.KeepAlive(other)
	return __o
}

func NewQStringMatcher6(pattern string, cs CaseSensitivity) *QStringMatcher {
	var pattern_p uintptr
	if len(pattern) > 0 {
		pattern_p = uintptr(unsafe.Pointer(unsafe.StringData(pattern)))
	}
	__r, _, _ := purego.SyscallN(addr_QStringMatcher_new6, pattern_p, uintptr(len(pattern)), uintptr(cs))
	__o := newownedQStringMatcher(__r)
	runtime.KeepAlive(pattern)
	return __o
}

func NewQStringMatcher7(uc *QChar, len_ int64, cs CaseSensitivity) *QStringMatcher {
	__r, _, _ := purego.SyscallN(addr_QStringMatcher_new7, uintptr(uc.UnsafePointer()), uintptr(len_), uintptr(cs))
	__o := newownedQStringMatcher(__r)
	runtime.KeepAlive(uc)
	return __o
}

func NewQStringMatcher8(pattern *QStringView, cs CaseSensitivity) *QStringMatcher {
	__r, _, _ := purego.SyscallN(addr_QStringMatcher_new8, uintptr(pattern.UnsafePointer()), uintptr(cs))
	__o := newownedQStringMatcher(__r)
	runtime.KeepAlive(pattern)
	return __o
}

func (this *QStringMatcher) OperatorAssign(other *QStringMatcher) *QStringMatcher {
	__r, _, _ := purego.SyscallN(addr_QStringMatcher_OperatorAssign, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return bindQStringMatcher(__r)
}

func (this *QStringMatcher) SetPattern(pattern string) {
	var pattern_p uintptr
	if len(pattern) > 0 {
		pattern_p = uintptr(unsafe.Pointer(unsafe.StringData(pattern)))
	}
	purego.SyscallN(addr_QStringMatcher_setPattern, this.h, pattern_p, uintptr(len(pattern)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pattern)
}

func (this *QStringMatcher) SetCaseSensitivity(cs CaseSensitivity) {
	purego.SyscallN(addr_QStringMatcher_setCaseSensitivity, this.h, uintptr(cs))
	runtime.KeepAlive(this)
}

func (this *QStringMatcher) IndexIn(str string) int64 {
	var str_p uintptr
	if len(str) > 0 {
		str_p = uintptr(unsafe.Pointer(unsafe.StringData(str)))
	}
	__r0, _, _ := purego.SyscallN(addr_QStringMatcher_indexIn, this.h, str_p, uintptr(len(str)))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(str)
	return __r
}

func (this *QStringMatcher) IndexIn2(str *QChar, length int64) int64 {
	__r0, _, _ := purego.SyscallN(addr_QStringMatcher_indexIn2, this.h, uintptr(str.UnsafePointer()), uintptr(length))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(str)
	return __r
}

func (this *QStringMatcher) IndexInWithStr(str *QStringView) int64 {
	__r0, _, _ := purego.SyscallN(addr_QStringMatcher_indexIn3, this.h, uintptr(str.UnsafePointer()))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(str)
	return __r
}

func (this *QStringMatcher) Pattern() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStringMatcher_pattern, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QStringMatcher) PatternView() *QStringView {
	__r, _, _ := purego.SyscallN(addr_QStringMatcher_patternView, this.h)
	runtime.KeepAlive(this)
	return newownedQStringView(__r)
}

func (this *QStringMatcher) CaseSensitivity() CaseSensitivity {
	__r0, _, _ := purego.SyscallN(addr_QStringMatcher_caseSensitivity, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return CaseSensitivity(__r)
}

func (this *QStringMatcher) IndexIn3(str string, from int64) int64 {
	var str_p uintptr
	if len(str) > 0 {
		str_p = uintptr(unsafe.Pointer(unsafe.StringData(str)))
	}
	__r0, _, _ := purego.SyscallN(addr_QStringMatcher_indexIn4, this.h, str_p, uintptr(len(str)), uintptr(from))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(str)
	return __r
}

func (this *QStringMatcher) IndexIn4(str *QChar, length int64, from int64) int64 {
	__r0, _, _ := purego.SyscallN(addr_QStringMatcher_indexIn5, this.h, uintptr(str.UnsafePointer()), uintptr(length), uintptr(from))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(str)
	return __r
}

func (this *QStringMatcher) IndexIn5(str *QStringView, from int64) int64 {
	__r0, _, _ := purego.SyscallN(addr_QStringMatcher_indexIn6, this.h, uintptr(str.UnsafePointer()), uintptr(from))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(str)
	return __r
}

func (this *QStringMatcher) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QStringMatcher_delete(this.h)
	}
	this.UnsafeRelease()
}
