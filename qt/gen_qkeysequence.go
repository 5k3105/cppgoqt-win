package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QKeySequence_new uintptr
	addr_QKeySequence_new2 uintptr
	addr_QKeySequence_new3 uintptr
	addr_QKeySequence_new4 uintptr
	addr_QKeySequence_new5 uintptr
	addr_QKeySequence_new6 uintptr
	addr_QKeySequence_new7 uintptr
	addr_QKeySequence_new8 uintptr
	addr_QKeySequence_new9 uintptr
	addr_QKeySequence_new10 uintptr
	addr_QKeySequence_new11 uintptr
	addr_QKeySequence_new12 uintptr
	addr_QKeySequence_new13 uintptr
	addr_QKeySequence_count uintptr
	addr_QKeySequence_isEmpty uintptr
	addr_QKeySequence_toString uintptr
	addr_QKeySequence_fromString uintptr
	addr_QKeySequence_listFromString uintptr
	addr_QKeySequence_listToString uintptr
	addr_QKeySequence_matches uintptr
	addr_QKeySequence_mnemonic uintptr
	addr_QKeySequence_keyBindings uintptr
	addr_QKeySequence_ToQVariant uintptr
	addr_QKeySequence_OperatorSubscript uintptr
	addr_QKeySequence_OperatorAssign uintptr
	addr_QKeySequence_swap uintptr
	addr_QKeySequence_OperatorEqual uintptr
	addr_QKeySequence_OperatorNotEqual uintptr
	addr_QKeySequence_OperatorLesser uintptr
	addr_QKeySequence_OperatorGreater uintptr
	addr_QKeySequence_OperatorLesserOrEqual uintptr
	addr_QKeySequence_OperatorGreaterOrEqual uintptr
	addr_QKeySequence_isDetached uintptr
	addr_QKeySequence_toString2 uintptr
	addr_QKeySequence_fromString2 uintptr
	addr_QKeySequence_listFromString2 uintptr
	addr_QKeySequence_listToString2 uintptr
	addr_QKeySequence_toQVariant uintptr
	pg_QKeySequence_delete func(uintptr)
)

func init() {
	symaddr(&addr_QKeySequence_new, "QKeySequence_new")
	symaddr(&addr_QKeySequence_new2, "QKeySequence_new2")
	symaddr(&addr_QKeySequence_new3, "QKeySequence_new3")
	symaddr(&addr_QKeySequence_new4, "QKeySequence_new4")
	symaddr(&addr_QKeySequence_new5, "QKeySequence_new5")
	symaddr(&addr_QKeySequence_new6, "QKeySequence_new6")
	symaddr(&addr_QKeySequence_new7, "QKeySequence_new7")
	symaddr(&addr_QKeySequence_new8, "QKeySequence_new8")
	symaddr(&addr_QKeySequence_new9, "QKeySequence_new9")
	symaddr(&addr_QKeySequence_new10, "QKeySequence_new10")
	symaddr(&addr_QKeySequence_new11, "QKeySequence_new11")
	symaddr(&addr_QKeySequence_new12, "QKeySequence_new12")
	symaddr(&addr_QKeySequence_new13, "QKeySequence_new13")
	symaddr(&addr_QKeySequence_count, "QKeySequence_count")
	symaddr(&addr_QKeySequence_isEmpty, "QKeySequence_isEmpty")
	symaddr(&addr_QKeySequence_toString, "QKeySequence_toString")
	symaddr(&addr_QKeySequence_fromString, "QKeySequence_fromString")
	symaddr(&addr_QKeySequence_listFromString, "QKeySequence_listFromString")
	symaddr(&addr_QKeySequence_listToString, "QKeySequence_listToString")
	symaddr(&addr_QKeySequence_matches, "QKeySequence_matches")
	symaddr(&addr_QKeySequence_mnemonic, "QKeySequence_mnemonic")
	symaddr(&addr_QKeySequence_keyBindings, "QKeySequence_keyBindings")
	symaddr(&addr_QKeySequence_ToQVariant, "QKeySequence_ToQVariant")
	symaddr(&addr_QKeySequence_OperatorSubscript, "QKeySequence_OperatorSubscript")
	symaddr(&addr_QKeySequence_OperatorAssign, "QKeySequence_OperatorAssign")
	symaddr(&addr_QKeySequence_swap, "QKeySequence_swap")
	symaddr(&addr_QKeySequence_OperatorEqual, "QKeySequence_OperatorEqual")
	symaddr(&addr_QKeySequence_OperatorNotEqual, "QKeySequence_OperatorNotEqual")
	symaddr(&addr_QKeySequence_OperatorLesser, "QKeySequence_OperatorLesser")
	symaddr(&addr_QKeySequence_OperatorGreater, "QKeySequence_OperatorGreater")
	symaddr(&addr_QKeySequence_OperatorLesserOrEqual, "QKeySequence_OperatorLesserOrEqual")
	symaddr(&addr_QKeySequence_OperatorGreaterOrEqual, "QKeySequence_OperatorGreaterOrEqual")
	symaddr(&addr_QKeySequence_isDetached, "QKeySequence_isDetached")
	symaddr(&addr_QKeySequence_toString2, "QKeySequence_toString2")
	symaddr(&addr_QKeySequence_fromString2, "QKeySequence_fromString2")
	symaddr(&addr_QKeySequence_listFromString2, "QKeySequence_listFromString2")
	symaddr(&addr_QKeySequence_listToString2, "QKeySequence_listToString2")
	symaddr(&addr_QKeySequence_toQVariant, "QKeySequence_toQVariant")
	reg(&pg_QKeySequence_delete, "QKeySequence_delete")
}

type QKeySequence struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QKeySequence) UnsafeGoID() int64     { return this.vid }
func (this *QKeySequence) UnsafeSetGoID(id int64) { this.vid = id }

func bindQKeySequence(h uintptr) *QKeySequence {
	if h == 0 {
		return nil
	}
	return &QKeySequence{h: h}
}

func newownedQKeySequence(h uintptr) *QKeySequence {
	o := bindQKeySequence(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QKeySequence).Delete)
	}
	return o
}

func (this *QKeySequence) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QKeySequence) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQKeySequence(h unsafe.Pointer) *QKeySequence { return bindQKeySequence(uintptr(h)) }
func UnsafeNewQKeySequenceOwned(h unsafe.Pointer) *QKeySequence { return newownedQKeySequence(uintptr(h)) }

func (this *QKeySequence) UnsafeRelease() { this.h = 0 }

func NewQKeySequence() *QKeySequence {
	__r, _, _ := purego.SyscallN(addr_QKeySequence_new)
	__o := newownedQKeySequence(__r)
	return __o
}

func NewQKeySequence2(key string) *QKeySequence {
	var key_p uintptr
	if len(key) > 0 {
		key_p = uintptr(unsafe.Pointer(unsafe.StringData(key)))
	}
	__r, _, _ := purego.SyscallN(addr_QKeySequence_new2, key_p, uintptr(len(key)))
	__o := newownedQKeySequence(__r)
	runtime.KeepAlive(key)
	return __o
}

func NewQKeySequence3(k1 int) *QKeySequence {
	__r, _, _ := purego.SyscallN(addr_QKeySequence_new3, uintptr(k1))
	__o := newownedQKeySequence(__r)
	return __o
}

func NewQKeySequence4(k1 *QKeyCombination) *QKeySequence {
	__r, _, _ := purego.SyscallN(addr_QKeySequence_new4, uintptr(k1.UnsafePointer()))
	__o := newownedQKeySequence(__r)
	runtime.KeepAlive(k1)
	return __o
}

func NewQKeySequence5(ks *QKeySequence) *QKeySequence {
	__r, _, _ := purego.SyscallN(addr_QKeySequence_new5, uintptr(ks.UnsafePointer()))
	__o := newownedQKeySequence(__r)
	runtime.KeepAlive(ks)
	return __o
}

func NewQKeySequence6(key QKeySequence__StandardKey) *QKeySequence {
	__r, _, _ := purego.SyscallN(addr_QKeySequence_new6, uintptr(key))
	__o := newownedQKeySequence(__r)
	return __o
}

func NewQKeySequence7(key string, format QKeySequence__SequenceFormat) *QKeySequence {
	var key_p uintptr
	if len(key) > 0 {
		key_p = uintptr(unsafe.Pointer(unsafe.StringData(key)))
	}
	__r, _, _ := purego.SyscallN(addr_QKeySequence_new7, key_p, uintptr(len(key)), uintptr(format))
	__o := newownedQKeySequence(__r)
	runtime.KeepAlive(key)
	return __o
}

func NewQKeySequence8(k1 int, k2 int) *QKeySequence {
	__r, _, _ := purego.SyscallN(addr_QKeySequence_new8, uintptr(k1), uintptr(k2))
	__o := newownedQKeySequence(__r)
	return __o
}

func NewQKeySequence9(k1 int, k2 int, k3 int) *QKeySequence {
	__r, _, _ := purego.SyscallN(addr_QKeySequence_new9, uintptr(k1), uintptr(k2), uintptr(k3))
	__o := newownedQKeySequence(__r)
	return __o
}

func NewQKeySequence10(k1 int, k2 int, k3 int, k4 int) *QKeySequence {
	__r, _, _ := purego.SyscallN(addr_QKeySequence_new10, uintptr(k1), uintptr(k2), uintptr(k3), uintptr(k4))
	__o := newownedQKeySequence(__r)
	return __o
}

func NewQKeySequence11(k1 *QKeyCombination, k2 *QKeyCombination) *QKeySequence {
	__r, _, _ := purego.SyscallN(addr_QKeySequence_new11, uintptr(k1.UnsafePointer()), uintptr(k2.UnsafePointer()))
	__o := newownedQKeySequence(__r)
	runtime.KeepAlive(k1)
	runtime.KeepAlive(k2)
	return __o
}

func NewQKeySequence12(k1 *QKeyCombination, k2 *QKeyCombination, k3 *QKeyCombination) *QKeySequence {
	__r, _, _ := purego.SyscallN(addr_QKeySequence_new12, uintptr(k1.UnsafePointer()), uintptr(k2.UnsafePointer()), uintptr(k3.UnsafePointer()))
	__o := newownedQKeySequence(__r)
	runtime.KeepAlive(k1)
	runtime.KeepAlive(k2)
	runtime.KeepAlive(k3)
	return __o
}

func NewQKeySequence13(k1 *QKeyCombination, k2 *QKeyCombination, k3 *QKeyCombination, k4 *QKeyCombination) *QKeySequence {
	__r, _, _ := purego.SyscallN(addr_QKeySequence_new13, uintptr(k1.UnsafePointer()), uintptr(k2.UnsafePointer()), uintptr(k3.UnsafePointer()), uintptr(k4.UnsafePointer()))
	__o := newownedQKeySequence(__r)
	runtime.KeepAlive(k1)
	runtime.KeepAlive(k2)
	runtime.KeepAlive(k3)
	runtime.KeepAlive(k4)
	return __o
}

func (this *QKeySequence) Count() int {
	__r0, _, _ := purego.SyscallN(addr_QKeySequence_count, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QKeySequence) IsEmpty() bool {
	__r0, _, _ := purego.SyscallN(addr_QKeySequence_isEmpty, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QKeySequence) ToString() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QKeySequence_toString, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func QKeySequence_FromString(str string) *QKeySequence {
	var str_p uintptr
	if len(str) > 0 {
		str_p = uintptr(unsafe.Pointer(unsafe.StringData(str)))
	}
	__r, _, _ := purego.SyscallN(addr_QKeySequence_fromString, str_p, uintptr(len(str)))
	runtime.KeepAlive(str)
	return newownedQKeySequence(__r)
}

func QKeySequence_ListFromString(str string) []*QKeySequence {
	var str_p uintptr
	if len(str) > 0 {
		str_p = uintptr(unsafe.Pointer(unsafe.StringData(str)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QKeySequence_listFromString, str_p, uintptr(len(str)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(str)
	__out := make([]*QKeySequence, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQKeySequence(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QKeySequence_ListToString(list []*QKeySequence) string {
	list_arr := make([]uintptr, len(list))
	for __i, __e := range list {
		list_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var list_p uintptr
	if len(list_arr) > 0 {
		list_p = uintptr(unsafe.Pointer(&list_arr[0]))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QKeySequence_listToString, list_p, uintptr(len(list)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(list)
	runtime.KeepAlive(list_arr)
	return pgStr(__r, __n)
}

func (this *QKeySequence) Matches(seq *QKeySequence) QKeySequence__SequenceMatch {
	__r0, _, _ := purego.SyscallN(addr_QKeySequence_matches, this.h, uintptr(seq.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(seq)
	return QKeySequence__SequenceMatch(__r)
}

func QKeySequence_Mnemonic(text string) *QKeySequence {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QKeySequence_mnemonic, text_p, uintptr(len(text)))
	runtime.KeepAlive(text)
	return newownedQKeySequence(__r)
}

func QKeySequence_KeyBindings(key QKeySequence__StandardKey) []*QKeySequence {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QKeySequence_keyBindings, uintptr(key), uintptr(unsafe.Pointer(&__n)))
	__out := make([]*QKeySequence, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQKeySequence(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QKeySequence) ToQVariant() *QVariant {
	__r, _, _ := purego.SyscallN(addr_QKeySequence_ToQVariant, this.h)
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QKeySequence) OperatorSubscript(i uint) *QKeyCombination {
	__r, _, _ := purego.SyscallN(addr_QKeySequence_OperatorSubscript, this.h, uintptr(i))
	runtime.KeepAlive(this)
	return newownedQKeyCombination(__r)
}

func (this *QKeySequence) OperatorAssign(other *QKeySequence) *QKeySequence {
	__r, _, _ := purego.SyscallN(addr_QKeySequence_OperatorAssign, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return bindQKeySequence(__r)
}

func (this *QKeySequence) Swap(other *QKeySequence) {
	purego.SyscallN(addr_QKeySequence_swap, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
}

func (this *QKeySequence) OperatorEqual(other *QKeySequence) bool {
	__r0, _, _ := purego.SyscallN(addr_QKeySequence_OperatorEqual, this.h, uintptr(other.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return __r != 0
}

func (this *QKeySequence) OperatorNotEqual(other *QKeySequence) bool {
	__r0, _, _ := purego.SyscallN(addr_QKeySequence_OperatorNotEqual, this.h, uintptr(other.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return __r != 0
}

func (this *QKeySequence) OperatorLesser(ks *QKeySequence) bool {
	__r0, _, _ := purego.SyscallN(addr_QKeySequence_OperatorLesser, this.h, uintptr(ks.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(ks)
	return __r != 0
}

func (this *QKeySequence) OperatorGreater(other *QKeySequence) bool {
	__r0, _, _ := purego.SyscallN(addr_QKeySequence_OperatorGreater, this.h, uintptr(other.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return __r != 0
}

func (this *QKeySequence) OperatorLesserOrEqual(other *QKeySequence) bool {
	__r0, _, _ := purego.SyscallN(addr_QKeySequence_OperatorLesserOrEqual, this.h, uintptr(other.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return __r != 0
}

func (this *QKeySequence) OperatorGreaterOrEqual(other *QKeySequence) bool {
	__r0, _, _ := purego.SyscallN(addr_QKeySequence_OperatorGreaterOrEqual, this.h, uintptr(other.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return __r != 0
}

func (this *QKeySequence) IsDetached() bool {
	__r0, _, _ := purego.SyscallN(addr_QKeySequence_isDetached, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QKeySequence) ToStringWithFormat(format QKeySequence__SequenceFormat) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QKeySequence_toString2, this.h, uintptr(format), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func QKeySequence_FromString2(str string, format QKeySequence__SequenceFormat) *QKeySequence {
	var str_p uintptr
	if len(str) > 0 {
		str_p = uintptr(unsafe.Pointer(unsafe.StringData(str)))
	}
	__r, _, _ := purego.SyscallN(addr_QKeySequence_fromString2, str_p, uintptr(len(str)), uintptr(format))
	runtime.KeepAlive(str)
	return newownedQKeySequence(__r)
}

func QKeySequence_ListFromString2(str string, format QKeySequence__SequenceFormat) []*QKeySequence {
	var str_p uintptr
	if len(str) > 0 {
		str_p = uintptr(unsafe.Pointer(unsafe.StringData(str)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QKeySequence_listFromString2, str_p, uintptr(len(str)), uintptr(format), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(str)
	__out := make([]*QKeySequence, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQKeySequence(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QKeySequence_ListToString2(list []*QKeySequence, format QKeySequence__SequenceFormat) string {
	list_arr := make([]uintptr, len(list))
	for __i, __e := range list {
		list_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var list_p uintptr
	if len(list_arr) > 0 {
		list_p = uintptr(unsafe.Pointer(&list_arr[0]))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QKeySequence_listToString2, list_p, uintptr(len(list)), uintptr(format), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(list)
	runtime.KeepAlive(list_arr)
	return pgStr(__r, __n)
}

func (this *QKeySequence) ToQVariant2() *QVariant {
	__r, _, _ := purego.SyscallN(addr_QKeySequence_toQVariant, this.h)
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QKeySequence) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QKeySequence_delete(this.h)
	}
	this.UnsafeRelease()
}
