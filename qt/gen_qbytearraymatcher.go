package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QByteArrayMatcher_new uintptr
	addr_QByteArrayMatcher_new2 uintptr
	addr_QByteArrayMatcher_new3 uintptr
	addr_QByteArrayMatcher_new4 uintptr
	addr_QByteArrayMatcher_new5 uintptr
	addr_QByteArrayMatcher_new6 uintptr
	addr_QByteArrayMatcher_OperatorAssign uintptr
	addr_QByteArrayMatcher_setPattern uintptr
	addr_QByteArrayMatcher_indexIn uintptr
	addr_QByteArrayMatcher_indexIn2 uintptr
	addr_QByteArrayMatcher_pattern uintptr
	addr_QByteArrayMatcher_indexIn3 uintptr
	addr_QByteArrayMatcher_indexIn4 uintptr
	pg_QByteArrayMatcher_delete func(uintptr)
)

func init() {
	symaddr(&addr_QByteArrayMatcher_new, "QByteArrayMatcher_new")
	symaddr(&addr_QByteArrayMatcher_new2, "QByteArrayMatcher_new2")
	symaddr(&addr_QByteArrayMatcher_new3, "QByteArrayMatcher_new3")
	symaddr(&addr_QByteArrayMatcher_new4, "QByteArrayMatcher_new4")
	symaddr(&addr_QByteArrayMatcher_new5, "QByteArrayMatcher_new5")
	symaddr(&addr_QByteArrayMatcher_new6, "QByteArrayMatcher_new6")
	symaddr(&addr_QByteArrayMatcher_OperatorAssign, "QByteArrayMatcher_OperatorAssign")
	symaddr(&addr_QByteArrayMatcher_setPattern, "QByteArrayMatcher_setPattern")
	symaddr(&addr_QByteArrayMatcher_indexIn, "QByteArrayMatcher_indexIn")
	symaddr(&addr_QByteArrayMatcher_indexIn2, "QByteArrayMatcher_indexIn2")
	symaddr(&addr_QByteArrayMatcher_pattern, "QByteArrayMatcher_pattern")
	symaddr(&addr_QByteArrayMatcher_indexIn3, "QByteArrayMatcher_indexIn3")
	symaddr(&addr_QByteArrayMatcher_indexIn4, "QByteArrayMatcher_indexIn4")
	reg(&pg_QByteArrayMatcher_delete, "QByteArrayMatcher_delete")
}

type QByteArrayMatcher struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QByteArrayMatcher) UnsafeGoID() int64     { return this.vid }
func (this *QByteArrayMatcher) UnsafeSetGoID(id int64) { this.vid = id }

func bindQByteArrayMatcher(h uintptr) *QByteArrayMatcher {
	if h == 0 {
		return nil
	}
	return &QByteArrayMatcher{h: h}
}

func newownedQByteArrayMatcher(h uintptr) *QByteArrayMatcher {
	o := bindQByteArrayMatcher(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QByteArrayMatcher).Delete)
	}
	return o
}

func (this *QByteArrayMatcher) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QByteArrayMatcher) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQByteArrayMatcher(h unsafe.Pointer) *QByteArrayMatcher { return bindQByteArrayMatcher(uintptr(h)) }
func UnsafeNewQByteArrayMatcherOwned(h unsafe.Pointer) *QByteArrayMatcher { return newownedQByteArrayMatcher(uintptr(h)) }

func (this *QByteArrayMatcher) UnsafeRelease() { this.h = 0 }

func NewQByteArrayMatcher() *QByteArrayMatcher {
	__r, _, _ := purego.SyscallN(addr_QByteArrayMatcher_new)
	__o := newownedQByteArrayMatcher(__r)
	return __o
}

func NewQByteArrayMatcher2(pattern *QByteArray) *QByteArrayMatcher {
	__r, _, _ := purego.SyscallN(addr_QByteArrayMatcher_new2, uintptr(pattern.UnsafePointer()))
	__o := newownedQByteArrayMatcher(__r)
	runtime.KeepAlive(pattern)
	return __o
}

func NewQByteArrayMatcher3(pattern *QByteArrayView) *QByteArrayMatcher {
	__r, _, _ := purego.SyscallN(addr_QByteArrayMatcher_new3, uintptr(pattern.UnsafePointer()))
	__o := newownedQByteArrayMatcher(__r)
	runtime.KeepAlive(pattern)
	return __o
}

func NewQByteArrayMatcher4(pattern string) *QByteArrayMatcher {
	pattern_b := append([]byte(pattern), 0)
	__r, _, _ := purego.SyscallN(addr_QByteArrayMatcher_new4, uintptr(unsafe.Pointer(&pattern_b[0])))
	__o := newownedQByteArrayMatcher(__r)
	runtime.KeepAlive(pattern_b)
	return __o
}

func NewQByteArrayMatcher5(other *QByteArrayMatcher) *QByteArrayMatcher {
	__r, _, _ := purego.SyscallN(addr_QByteArrayMatcher_new5, uintptr(other.UnsafePointer()))
	__o := newownedQByteArrayMatcher(__r)
	runtime.KeepAlive(other)
	return __o
}

func NewQByteArrayMatcher6(pattern string, length int64) *QByteArrayMatcher {
	pattern_b := append([]byte(pattern), 0)
	__r, _, _ := purego.SyscallN(addr_QByteArrayMatcher_new6, uintptr(unsafe.Pointer(&pattern_b[0])), uintptr(length))
	__o := newownedQByteArrayMatcher(__r)
	runtime.KeepAlive(pattern_b)
	return __o
}

func (this *QByteArrayMatcher) OperatorAssign(other *QByteArrayMatcher) *QByteArrayMatcher {
	__r, _, _ := purego.SyscallN(addr_QByteArrayMatcher_OperatorAssign, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return bindQByteArrayMatcher(__r)
}

func (this *QByteArrayMatcher) SetPattern(pattern *QByteArray) {
	purego.SyscallN(addr_QByteArrayMatcher_setPattern, this.h, uintptr(pattern.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pattern)
}

func (this *QByteArrayMatcher) IndexIn(str string, len_ int64) int64 {
	str_b := append([]byte(str), 0)
	__r0, _, _ := purego.SyscallN(addr_QByteArrayMatcher_indexIn, this.h, uintptr(unsafe.Pointer(&str_b[0])), uintptr(len_))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(str_b)
	return __r
}

func (this *QByteArrayMatcher) IndexInWithData(data *QByteArrayView) int64 {
	__r0, _, _ := purego.SyscallN(addr_QByteArrayMatcher_indexIn2, this.h, uintptr(data.UnsafePointer()))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(data)
	return __r
}

func (this *QByteArrayMatcher) Pattern() *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QByteArrayMatcher_pattern, this.h)
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QByteArrayMatcher) IndexIn2(str string, len_ int64, from int64) int64 {
	str_b := append([]byte(str), 0)
	__r0, _, _ := purego.SyscallN(addr_QByteArrayMatcher_indexIn3, this.h, uintptr(unsafe.Pointer(&str_b[0])), uintptr(len_), uintptr(from))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(str_b)
	return __r
}

func (this *QByteArrayMatcher) IndexIn3(data *QByteArrayView, from int64) int64 {
	__r0, _, _ := purego.SyscallN(addr_QByteArrayMatcher_indexIn4, this.h, uintptr(data.UnsafePointer()), uintptr(from))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(data)
	return __r
}

func (this *QByteArrayMatcher) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QByteArrayMatcher_delete(this.h)
	}
	this.UnsafeRelease()
}
