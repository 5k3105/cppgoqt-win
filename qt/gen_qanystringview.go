package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QAnyStringView_new uintptr
	addr_QAnyStringView_new2 uintptr
	addr_QAnyStringView_new3 uintptr
	addr_QAnyStringView_new4 uintptr
	addr_QAnyStringView_new5 uintptr
	addr_QAnyStringView_new6 uintptr
	addr_QAnyStringView_mid uintptr
	addr_QAnyStringView_left uintptr
	addr_QAnyStringView_right uintptr
	addr_QAnyStringView_sliced uintptr
	addr_QAnyStringView_sliced2 uintptr
	addr_QAnyStringView_first uintptr
	addr_QAnyStringView_last uintptr
	addr_QAnyStringView_chopped uintptr
	addr_QAnyStringView_slice uintptr
	addr_QAnyStringView_slice2 uintptr
	addr_QAnyStringView_truncate uintptr
	addr_QAnyStringView_chop uintptr
	addr_QAnyStringView_toString uintptr
	addr_QAnyStringView_size uintptr
	addr_QAnyStringView_compare uintptr
	addr_QAnyStringView_equal uintptr
	addr_QAnyStringView_front uintptr
	addr_QAnyStringView_back uintptr
	addr_QAnyStringView_empty uintptr
	addr_QAnyStringView_size_bytes uintptr
	addr_QAnyStringView_max_size uintptr
	addr_QAnyStringView_isNull uintptr
	addr_QAnyStringView_isEmpty uintptr
	addr_QAnyStringView_length uintptr
	addr_QAnyStringView_OperatorAssign uintptr
	addr_QAnyStringView_mid2 uintptr
	addr_QAnyStringView_compare2 uintptr
	addr_QAnyStringView_OperatorEqual uintptr
	addr_QAnyStringView_OperatorNotEqual uintptr
	addr_QAnyStringView_OperatorLesser uintptr
	addr_QAnyStringView_OperatorGreater uintptr
	addr_QAnyStringView_OperatorLesserOrEqual uintptr
	addr_QAnyStringView_OperatorGreaterOrEqual uintptr
	pg_QAnyStringView_delete func(uintptr)
)

func init() {
	symaddr(&addr_QAnyStringView_new, "QAnyStringView_new")
	symaddr(&addr_QAnyStringView_new2, "QAnyStringView_new2")
	symaddr(&addr_QAnyStringView_new3, "QAnyStringView_new3")
	symaddr(&addr_QAnyStringView_new4, "QAnyStringView_new4")
	symaddr(&addr_QAnyStringView_new5, "QAnyStringView_new5")
	symaddr(&addr_QAnyStringView_new6, "QAnyStringView_new6")
	symaddr(&addr_QAnyStringView_mid, "QAnyStringView_mid")
	symaddr(&addr_QAnyStringView_left, "QAnyStringView_left")
	symaddr(&addr_QAnyStringView_right, "QAnyStringView_right")
	symaddr(&addr_QAnyStringView_sliced, "QAnyStringView_sliced")
	symaddr(&addr_QAnyStringView_sliced2, "QAnyStringView_sliced2")
	symaddr(&addr_QAnyStringView_first, "QAnyStringView_first")
	symaddr(&addr_QAnyStringView_last, "QAnyStringView_last")
	symaddr(&addr_QAnyStringView_chopped, "QAnyStringView_chopped")
	symaddr(&addr_QAnyStringView_slice, "QAnyStringView_slice")
	symaddr(&addr_QAnyStringView_slice2, "QAnyStringView_slice2")
	symaddr(&addr_QAnyStringView_truncate, "QAnyStringView_truncate")
	symaddr(&addr_QAnyStringView_chop, "QAnyStringView_chop")
	symaddr(&addr_QAnyStringView_toString, "QAnyStringView_toString")
	symaddr(&addr_QAnyStringView_size, "QAnyStringView_size")
	symaddr(&addr_QAnyStringView_compare, "QAnyStringView_compare")
	symaddr(&addr_QAnyStringView_equal, "QAnyStringView_equal")
	symaddr(&addr_QAnyStringView_front, "QAnyStringView_front")
	symaddr(&addr_QAnyStringView_back, "QAnyStringView_back")
	symaddr(&addr_QAnyStringView_empty, "QAnyStringView_empty")
	symaddr(&addr_QAnyStringView_size_bytes, "QAnyStringView_size_bytes")
	symaddr(&addr_QAnyStringView_max_size, "QAnyStringView_max_size")
	symaddr(&addr_QAnyStringView_isNull, "QAnyStringView_isNull")
	symaddr(&addr_QAnyStringView_isEmpty, "QAnyStringView_isEmpty")
	symaddr(&addr_QAnyStringView_length, "QAnyStringView_length")
	symaddr(&addr_QAnyStringView_OperatorAssign, "QAnyStringView_OperatorAssign")
	symaddr(&addr_QAnyStringView_mid2, "QAnyStringView_mid2")
	symaddr(&addr_QAnyStringView_compare2, "QAnyStringView_compare2")
	symaddr(&addr_QAnyStringView_OperatorEqual, "QAnyStringView_OperatorEqual")
	symaddr(&addr_QAnyStringView_OperatorNotEqual, "QAnyStringView_OperatorNotEqual")
	symaddr(&addr_QAnyStringView_OperatorLesser, "QAnyStringView_OperatorLesser")
	symaddr(&addr_QAnyStringView_OperatorGreater, "QAnyStringView_OperatorGreater")
	symaddr(&addr_QAnyStringView_OperatorLesserOrEqual, "QAnyStringView_OperatorLesserOrEqual")
	symaddr(&addr_QAnyStringView_OperatorGreaterOrEqual, "QAnyStringView_OperatorGreaterOrEqual")
	reg(&pg_QAnyStringView_delete, "QAnyStringView_delete")
}

type QAnyStringView struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QAnyStringView) UnsafeGoID() int64     { return this.vid }
func (this *QAnyStringView) UnsafeSetGoID(id int64) { this.vid = id }

func bindQAnyStringView(h uintptr) *QAnyStringView {
	if h == 0 {
		return nil
	}
	return &QAnyStringView{h: h}
}

func newownedQAnyStringView(h uintptr) *QAnyStringView {
	o := bindQAnyStringView(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QAnyStringView).Delete)
	}
	return o
}

func (this *QAnyStringView) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QAnyStringView) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQAnyStringView(h unsafe.Pointer) *QAnyStringView { return bindQAnyStringView(uintptr(h)) }
func UnsafeNewQAnyStringViewOwned(h unsafe.Pointer) *QAnyStringView { return newownedQAnyStringView(uintptr(h)) }

func (this *QAnyStringView) UnsafeRelease() { this.h = 0 }

func NewQAnyStringView() *QAnyStringView {
	__r, _, _ := purego.SyscallN(addr_QAnyStringView_new)
	__o := newownedQAnyStringView(__r)
	return __o
}

func NewQAnyStringView2(str *QByteArray) *QAnyStringView {
	__r, _, _ := purego.SyscallN(addr_QAnyStringView_new2, uintptr(str.UnsafePointer()))
	__o := newownedQAnyStringView(__r)
	runtime.KeepAlive(str)
	return __o
}

func NewQAnyStringView3(str string) *QAnyStringView {
	var str_p uintptr
	if len(str) > 0 {
		str_p = uintptr(unsafe.Pointer(unsafe.StringData(str)))
	}
	__r, _, _ := purego.SyscallN(addr_QAnyStringView_new3, str_p, uintptr(len(str)))
	__o := newownedQAnyStringView(__r)
	runtime.KeepAlive(str)
	return __o
}

func NewQAnyStringView4(str *QLatin1String) *QAnyStringView {
	__r, _, _ := purego.SyscallN(addr_QAnyStringView_new4, uintptr(str.UnsafePointer()))
	__o := newownedQAnyStringView(__r)
	runtime.KeepAlive(str)
	return __o
}

func NewQAnyStringView5(v *QStringView) *QAnyStringView {
	__r, _, _ := purego.SyscallN(addr_QAnyStringView_new5, uintptr(v.UnsafePointer()))
	__o := newownedQAnyStringView(__r)
	runtime.KeepAlive(v)
	return __o
}

func NewQAnyStringView6(arg1 string) *QAnyStringView {
	var arg1_p uintptr
	if len(arg1) > 0 {
		arg1_p = uintptr(unsafe.Pointer(unsafe.StringData(arg1)))
	}
	__r, _, _ := purego.SyscallN(addr_QAnyStringView_new6, arg1_p, uintptr(len(arg1)))
	__o := newownedQAnyStringView(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QAnyStringView) Mid(pos int64) *QAnyStringView {
	__r, _, _ := purego.SyscallN(addr_QAnyStringView_mid, this.h, uintptr(pos))
	runtime.KeepAlive(this)
	return newownedQAnyStringView(__r)
}

func (this *QAnyStringView) Left(n int64) *QAnyStringView {
	__r, _, _ := purego.SyscallN(addr_QAnyStringView_left, this.h, uintptr(n))
	runtime.KeepAlive(this)
	return newownedQAnyStringView(__r)
}

func (this *QAnyStringView) Right(n int64) *QAnyStringView {
	__r, _, _ := purego.SyscallN(addr_QAnyStringView_right, this.h, uintptr(n))
	runtime.KeepAlive(this)
	return newownedQAnyStringView(__r)
}

func (this *QAnyStringView) Sliced(pos int64) *QAnyStringView {
	__r, _, _ := purego.SyscallN(addr_QAnyStringView_sliced, this.h, uintptr(pos))
	runtime.KeepAlive(this)
	return newownedQAnyStringView(__r)
}

func (this *QAnyStringView) Sliced2(pos int64, n int64) *QAnyStringView {
	__r, _, _ := purego.SyscallN(addr_QAnyStringView_sliced2, this.h, uintptr(pos), uintptr(n))
	runtime.KeepAlive(this)
	return newownedQAnyStringView(__r)
}

func (this *QAnyStringView) First(n int64) *QAnyStringView {
	__r, _, _ := purego.SyscallN(addr_QAnyStringView_first, this.h, uintptr(n))
	runtime.KeepAlive(this)
	return newownedQAnyStringView(__r)
}

func (this *QAnyStringView) Last(n int64) *QAnyStringView {
	__r, _, _ := purego.SyscallN(addr_QAnyStringView_last, this.h, uintptr(n))
	runtime.KeepAlive(this)
	return newownedQAnyStringView(__r)
}

func (this *QAnyStringView) Chopped(n int64) *QAnyStringView {
	__r, _, _ := purego.SyscallN(addr_QAnyStringView_chopped, this.h, uintptr(n))
	runtime.KeepAlive(this)
	return newownedQAnyStringView(__r)
}

func (this *QAnyStringView) Slice(pos int64) *QAnyStringView {
	__r, _, _ := purego.SyscallN(addr_QAnyStringView_slice, this.h, uintptr(pos))
	runtime.KeepAlive(this)
	return bindQAnyStringView(__r)
}

func (this *QAnyStringView) Slice2(pos int64, n int64) *QAnyStringView {
	__r, _, _ := purego.SyscallN(addr_QAnyStringView_slice2, this.h, uintptr(pos), uintptr(n))
	runtime.KeepAlive(this)
	return bindQAnyStringView(__r)
}

func (this *QAnyStringView) Truncate(n int64) {
	purego.SyscallN(addr_QAnyStringView_truncate, this.h, uintptr(n))
	runtime.KeepAlive(this)
}

func (this *QAnyStringView) Chop(n int64) {
	purego.SyscallN(addr_QAnyStringView_chop, this.h, uintptr(n))
	runtime.KeepAlive(this)
}

func (this *QAnyStringView) ToString() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAnyStringView_toString, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QAnyStringView) Size() int64 {
	__r0, _, _ := purego.SyscallN(addr_QAnyStringView_size, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func QAnyStringView_Compare(lhs string, rhs string) int {
	var lhs_p uintptr
	if len(lhs) > 0 {
		lhs_p = uintptr(unsafe.Pointer(unsafe.StringData(lhs)))
	}
	var rhs_p uintptr
	if len(rhs) > 0 {
		rhs_p = uintptr(unsafe.Pointer(unsafe.StringData(rhs)))
	}
	__r0, _, _ := purego.SyscallN(addr_QAnyStringView_compare, lhs_p, uintptr(len(lhs)), rhs_p, uintptr(len(rhs)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(lhs)
	runtime.KeepAlive(rhs)
	return int(__r)
}

func QAnyStringView_Equal(lhs string, rhs string) bool {
	var lhs_p uintptr
	if len(lhs) > 0 {
		lhs_p = uintptr(unsafe.Pointer(unsafe.StringData(lhs)))
	}
	var rhs_p uintptr
	if len(rhs) > 0 {
		rhs_p = uintptr(unsafe.Pointer(unsafe.StringData(rhs)))
	}
	__r0, _, _ := purego.SyscallN(addr_QAnyStringView_equal, lhs_p, uintptr(len(lhs)), rhs_p, uintptr(len(rhs)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(lhs)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QAnyStringView) Front() *QChar {
	__r, _, _ := purego.SyscallN(addr_QAnyStringView_front, this.h)
	runtime.KeepAlive(this)
	return newownedQChar(__r)
}

func (this *QAnyStringView) Back() *QChar {
	__r, _, _ := purego.SyscallN(addr_QAnyStringView_back, this.h)
	runtime.KeepAlive(this)
	return newownedQChar(__r)
}

func (this *QAnyStringView) Empty() bool {
	__r0, _, _ := purego.SyscallN(addr_QAnyStringView_empty, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QAnyStringView) Size_bytes() int64 {
	__r0, _, _ := purego.SyscallN(addr_QAnyStringView_size_bytes, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QAnyStringView) Max_size() int64 {
	__r0, _, _ := purego.SyscallN(addr_QAnyStringView_max_size, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QAnyStringView) IsNull() bool {
	__r0, _, _ := purego.SyscallN(addr_QAnyStringView_isNull, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QAnyStringView) IsEmpty() bool {
	__r0, _, _ := purego.SyscallN(addr_QAnyStringView_isEmpty, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QAnyStringView) Length() int64 {
	__r0, _, _ := purego.SyscallN(addr_QAnyStringView_length, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QAnyStringView) OperatorAssign(arg1 string) *QAnyStringView {
	var arg1_p uintptr
	if len(arg1) > 0 {
		arg1_p = uintptr(unsafe.Pointer(unsafe.StringData(arg1)))
	}
	__r, _, _ := purego.SyscallN(addr_QAnyStringView_OperatorAssign, this.h, arg1_p, uintptr(len(arg1)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQAnyStringView(__r)
}

func (this *QAnyStringView) Mid2(pos int64, n int64) *QAnyStringView {
	__r, _, _ := purego.SyscallN(addr_QAnyStringView_mid2, this.h, uintptr(pos), uintptr(n))
	runtime.KeepAlive(this)
	return newownedQAnyStringView(__r)
}

func QAnyStringView_Compare2(lhs string, rhs string, cs CaseSensitivity) int {
	var lhs_p uintptr
	if len(lhs) > 0 {
		lhs_p = uintptr(unsafe.Pointer(unsafe.StringData(lhs)))
	}
	var rhs_p uintptr
	if len(rhs) > 0 {
		rhs_p = uintptr(unsafe.Pointer(unsafe.StringData(rhs)))
	}
	__r0, _, _ := purego.SyscallN(addr_QAnyStringView_compare2, lhs_p, uintptr(len(lhs)), rhs_p, uintptr(len(rhs)), uintptr(cs))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(lhs)
	runtime.KeepAlive(rhs)
	return int(__r)
}

func (this *QAnyStringView) OperatorEqual(rhs string) bool {
	var rhs_p uintptr
	if len(rhs) > 0 {
		rhs_p = uintptr(unsafe.Pointer(unsafe.StringData(rhs)))
	}
	__r0, _, _ := purego.SyscallN(addr_QAnyStringView_OperatorEqual, this.h, rhs_p, uintptr(len(rhs)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QAnyStringView) OperatorNotEqual(rhs string) bool {
	var rhs_p uintptr
	if len(rhs) > 0 {
		rhs_p = uintptr(unsafe.Pointer(unsafe.StringData(rhs)))
	}
	__r0, _, _ := purego.SyscallN(addr_QAnyStringView_OperatorNotEqual, this.h, rhs_p, uintptr(len(rhs)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QAnyStringView) OperatorLesser(rhs string) bool {
	var rhs_p uintptr
	if len(rhs) > 0 {
		rhs_p = uintptr(unsafe.Pointer(unsafe.StringData(rhs)))
	}
	__r0, _, _ := purego.SyscallN(addr_QAnyStringView_OperatorLesser, this.h, rhs_p, uintptr(len(rhs)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QAnyStringView) OperatorGreater(rhs string) bool {
	var rhs_p uintptr
	if len(rhs) > 0 {
		rhs_p = uintptr(unsafe.Pointer(unsafe.StringData(rhs)))
	}
	__r0, _, _ := purego.SyscallN(addr_QAnyStringView_OperatorGreater, this.h, rhs_p, uintptr(len(rhs)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QAnyStringView) OperatorLesserOrEqual(rhs string) bool {
	var rhs_p uintptr
	if len(rhs) > 0 {
		rhs_p = uintptr(unsafe.Pointer(unsafe.StringData(rhs)))
	}
	__r0, _, _ := purego.SyscallN(addr_QAnyStringView_OperatorLesserOrEqual, this.h, rhs_p, uintptr(len(rhs)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QAnyStringView) OperatorGreaterOrEqual(rhs string) bool {
	var rhs_p uintptr
	if len(rhs) > 0 {
		rhs_p = uintptr(unsafe.Pointer(unsafe.StringData(rhs)))
	}
	__r0, _, _ := purego.SyscallN(addr_QAnyStringView_OperatorGreaterOrEqual, this.h, rhs_p, uintptr(len(rhs)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QAnyStringView) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QAnyStringView_delete(this.h)
	}
	this.UnsafeRelease()
}
