package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QMimeType_new uintptr
	addr_QMimeType_new2 uintptr
	addr_QMimeType_OperatorAssign uintptr
	addr_QMimeType_swap uintptr
	addr_QMimeType_isValid uintptr
	addr_QMimeType_isDefault uintptr
	addr_QMimeType_name uintptr
	addr_QMimeType_comment uintptr
	addr_QMimeType_genericIconName uintptr
	addr_QMimeType_iconName uintptr
	addr_QMimeType_globPatterns uintptr
	addr_QMimeType_parentMimeTypes uintptr
	addr_QMimeType_allAncestors uintptr
	addr_QMimeType_aliases uintptr
	addr_QMimeType_suffixes uintptr
	addr_QMimeType_preferredSuffix uintptr
	addr_QMimeType_inherits uintptr
	addr_QMimeType_filterString uintptr
	addr_QMimeType_OperatorEqual uintptr
	addr_QMimeType_OperatorNotEqual uintptr
	pg_QMimeType_delete func(uintptr)
)

func init() {
	symaddr(&addr_QMimeType_new, "QMimeType_new")
	symaddr(&addr_QMimeType_new2, "QMimeType_new2")
	symaddr(&addr_QMimeType_OperatorAssign, "QMimeType_OperatorAssign")
	symaddr(&addr_QMimeType_swap, "QMimeType_swap")
	symaddr(&addr_QMimeType_isValid, "QMimeType_isValid")
	symaddr(&addr_QMimeType_isDefault, "QMimeType_isDefault")
	symaddr(&addr_QMimeType_name, "QMimeType_name")
	symaddr(&addr_QMimeType_comment, "QMimeType_comment")
	symaddr(&addr_QMimeType_genericIconName, "QMimeType_genericIconName")
	symaddr(&addr_QMimeType_iconName, "QMimeType_iconName")
	symaddr(&addr_QMimeType_globPatterns, "QMimeType_globPatterns")
	symaddr(&addr_QMimeType_parentMimeTypes, "QMimeType_parentMimeTypes")
	symaddr(&addr_QMimeType_allAncestors, "QMimeType_allAncestors")
	symaddr(&addr_QMimeType_aliases, "QMimeType_aliases")
	symaddr(&addr_QMimeType_suffixes, "QMimeType_suffixes")
	symaddr(&addr_QMimeType_preferredSuffix, "QMimeType_preferredSuffix")
	symaddr(&addr_QMimeType_inherits, "QMimeType_inherits")
	symaddr(&addr_QMimeType_filterString, "QMimeType_filterString")
	symaddr(&addr_QMimeType_OperatorEqual, "QMimeType_OperatorEqual")
	symaddr(&addr_QMimeType_OperatorNotEqual, "QMimeType_OperatorNotEqual")
	reg(&pg_QMimeType_delete, "QMimeType_delete")
}

type QMimeType struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QMimeType) UnsafeGoID() int64     { return this.vid }
func (this *QMimeType) UnsafeSetGoID(id int64) { this.vid = id }

func bindQMimeType(h uintptr) *QMimeType {
	if h == 0 {
		return nil
	}
	return &QMimeType{h: h}
}

func newownedQMimeType(h uintptr) *QMimeType {
	o := bindQMimeType(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QMimeType).Delete)
	}
	return o
}

func (this *QMimeType) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QMimeType) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQMimeType(h unsafe.Pointer) *QMimeType { return bindQMimeType(uintptr(h)) }
func UnsafeNewQMimeTypeOwned(h unsafe.Pointer) *QMimeType { return newownedQMimeType(uintptr(h)) }

func (this *QMimeType) UnsafeRelease() { this.h = 0 }

func NewQMimeType() *QMimeType {
	__r, _, _ := purego.SyscallN(addr_QMimeType_new)
	__o := newownedQMimeType(__r)
	return __o
}

func NewQMimeType2(other *QMimeType) *QMimeType {
	__r, _, _ := purego.SyscallN(addr_QMimeType_new2, uintptr(other.UnsafePointer()))
	__o := newownedQMimeType(__r)
	runtime.KeepAlive(other)
	return __o
}

func (this *QMimeType) OperatorAssign(other *QMimeType) *QMimeType {
	__r, _, _ := purego.SyscallN(addr_QMimeType_OperatorAssign, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return bindQMimeType(__r)
}

func (this *QMimeType) Swap(other *QMimeType) {
	purego.SyscallN(addr_QMimeType_swap, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
}

func (this *QMimeType) IsValid() bool {
	__r0, _, _ := purego.SyscallN(addr_QMimeType_isValid, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMimeType) IsDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QMimeType_isDefault, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMimeType) Name() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QMimeType_name, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QMimeType) Comment() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QMimeType_comment, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QMimeType) GenericIconName() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QMimeType_genericIconName, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QMimeType) IconName() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QMimeType_iconName, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QMimeType) GlobPatterns() []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QMimeType_globPatterns, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]string, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = pgGoStr(__sl[__i])
		pgFree(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QMimeType) ParentMimeTypes() []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QMimeType_parentMimeTypes, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]string, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = pgGoStr(__sl[__i])
		pgFree(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QMimeType) AllAncestors() []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QMimeType_allAncestors, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]string, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = pgGoStr(__sl[__i])
		pgFree(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QMimeType) Aliases() []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QMimeType_aliases, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]string, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = pgGoStr(__sl[__i])
		pgFree(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QMimeType) Suffixes() []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QMimeType_suffixes, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]string, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = pgGoStr(__sl[__i])
		pgFree(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QMimeType) PreferredSuffix() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QMimeType_preferredSuffix, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QMimeType) Inherits(mimeTypeName string) bool {
	var mimeTypeName_p uintptr
	if len(mimeTypeName) > 0 {
		mimeTypeName_p = uintptr(unsafe.Pointer(unsafe.StringData(mimeTypeName)))
	}
	__r0, _, _ := purego.SyscallN(addr_QMimeType_inherits, this.h, mimeTypeName_p, uintptr(len(mimeTypeName)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(mimeTypeName)
	return __r != 0
}

func (this *QMimeType) FilterString() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QMimeType_filterString, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QMimeType) OperatorEqual(rhs *QMimeType) bool {
	__r0, _, _ := purego.SyscallN(addr_QMimeType_OperatorEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QMimeType) OperatorNotEqual(rhs *QMimeType) bool {
	__r0, _, _ := purego.SyscallN(addr_QMimeType_OperatorNotEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QMimeType) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QMimeType_delete(this.h)
	}
	this.UnsafeRelease()
}
