package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QLatin1StringMatcher_new uintptr
	addr_QLatin1StringMatcher_new2 uintptr
	addr_QLatin1StringMatcher_new3 uintptr
	addr_QLatin1StringMatcher_setPattern uintptr
	addr_QLatin1StringMatcher_pattern uintptr
	addr_QLatin1StringMatcher_setCaseSensitivity uintptr
	addr_QLatin1StringMatcher_caseSensitivity uintptr
	addr_QLatin1StringMatcher_indexIn uintptr
	addr_QLatin1StringMatcher_indexIn2 uintptr
	addr_QLatin1StringMatcher_indexIn3 uintptr
	addr_QLatin1StringMatcher_indexIn4 uintptr
	pg_QLatin1StringMatcher_delete func(uintptr)
)

func init() {
	symaddr(&addr_QLatin1StringMatcher_new, "QLatin1StringMatcher_new")
	symaddr(&addr_QLatin1StringMatcher_new2, "QLatin1StringMatcher_new2")
	symaddr(&addr_QLatin1StringMatcher_new3, "QLatin1StringMatcher_new3")
	symaddr(&addr_QLatin1StringMatcher_setPattern, "QLatin1StringMatcher_setPattern")
	symaddr(&addr_QLatin1StringMatcher_pattern, "QLatin1StringMatcher_pattern")
	symaddr(&addr_QLatin1StringMatcher_setCaseSensitivity, "QLatin1StringMatcher_setCaseSensitivity")
	symaddr(&addr_QLatin1StringMatcher_caseSensitivity, "QLatin1StringMatcher_caseSensitivity")
	symaddr(&addr_QLatin1StringMatcher_indexIn, "QLatin1StringMatcher_indexIn")
	symaddr(&addr_QLatin1StringMatcher_indexIn2, "QLatin1StringMatcher_indexIn2")
	symaddr(&addr_QLatin1StringMatcher_indexIn3, "QLatin1StringMatcher_indexIn3")
	symaddr(&addr_QLatin1StringMatcher_indexIn4, "QLatin1StringMatcher_indexIn4")
	reg(&pg_QLatin1StringMatcher_delete, "QLatin1StringMatcher_delete")
}

type QLatin1StringMatcher struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QLatin1StringMatcher) UnsafeGoID() int64     { return this.vid }
func (this *QLatin1StringMatcher) UnsafeSetGoID(id int64) { this.vid = id }

func bindQLatin1StringMatcher(h uintptr) *QLatin1StringMatcher {
	if h == 0 {
		return nil
	}
	return &QLatin1StringMatcher{h: h}
}

func newownedQLatin1StringMatcher(h uintptr) *QLatin1StringMatcher {
	o := bindQLatin1StringMatcher(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QLatin1StringMatcher).Delete)
	}
	return o
}

func (this *QLatin1StringMatcher) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QLatin1StringMatcher) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQLatin1StringMatcher(h unsafe.Pointer) *QLatin1StringMatcher { return bindQLatin1StringMatcher(uintptr(h)) }
func UnsafeNewQLatin1StringMatcherOwned(h unsafe.Pointer) *QLatin1StringMatcher { return newownedQLatin1StringMatcher(uintptr(h)) }

func (this *QLatin1StringMatcher) UnsafeRelease() { this.h = 0 }

func NewQLatin1StringMatcher() *QLatin1StringMatcher {
	__r, _, _ := purego.SyscallN(addr_QLatin1StringMatcher_new)
	__o := newownedQLatin1StringMatcher(__r)
	return __o
}

func NewQLatin1StringMatcher2(pattern *QLatin1String) *QLatin1StringMatcher {
	__r, _, _ := purego.SyscallN(addr_QLatin1StringMatcher_new2, uintptr(pattern.UnsafePointer()))
	__o := newownedQLatin1StringMatcher(__r)
	runtime.KeepAlive(pattern)
	return __o
}

func NewQLatin1StringMatcher3(pattern *QLatin1String, cs CaseSensitivity) *QLatin1StringMatcher {
	__r, _, _ := purego.SyscallN(addr_QLatin1StringMatcher_new3, uintptr(pattern.UnsafePointer()), uintptr(cs))
	__o := newownedQLatin1StringMatcher(__r)
	runtime.KeepAlive(pattern)
	return __o
}

func (this *QLatin1StringMatcher) SetPattern(pattern *QLatin1String) {
	purego.SyscallN(addr_QLatin1StringMatcher_setPattern, this.h, uintptr(pattern.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pattern)
}

func (this *QLatin1StringMatcher) Pattern() *QLatin1String {
	__r, _, _ := purego.SyscallN(addr_QLatin1StringMatcher_pattern, this.h)
	runtime.KeepAlive(this)
	return newownedQLatin1String(__r)
}

func (this *QLatin1StringMatcher) SetCaseSensitivity(cs CaseSensitivity) {
	purego.SyscallN(addr_QLatin1StringMatcher_setCaseSensitivity, this.h, uintptr(cs))
	runtime.KeepAlive(this)
}

func (this *QLatin1StringMatcher) CaseSensitivity() CaseSensitivity {
	__r0, _, _ := purego.SyscallN(addr_QLatin1StringMatcher_caseSensitivity, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return CaseSensitivity(__r)
}

func (this *QLatin1StringMatcher) IndexIn(haystack *QLatin1String) int64 {
	__r0, _, _ := purego.SyscallN(addr_QLatin1StringMatcher_indexIn, this.h, uintptr(haystack.UnsafePointer()))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(haystack)
	return __r
}

func (this *QLatin1StringMatcher) IndexInWithHaystack(haystack *QStringView) int64 {
	__r0, _, _ := purego.SyscallN(addr_QLatin1StringMatcher_indexIn2, this.h, uintptr(haystack.UnsafePointer()))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(haystack)
	return __r
}

func (this *QLatin1StringMatcher) IndexIn2(haystack *QLatin1String, from int64) int64 {
	__r0, _, _ := purego.SyscallN(addr_QLatin1StringMatcher_indexIn3, this.h, uintptr(haystack.UnsafePointer()), uintptr(from))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(haystack)
	return __r
}

func (this *QLatin1StringMatcher) IndexIn3(haystack *QStringView, from int64) int64 {
	__r0, _, _ := purego.SyscallN(addr_QLatin1StringMatcher_indexIn4, this.h, uintptr(haystack.UnsafePointer()), uintptr(from))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(haystack)
	return __r
}

func (this *QLatin1StringMatcher) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QLatin1StringMatcher_delete(this.h)
	}
	this.UnsafeRelease()
}
