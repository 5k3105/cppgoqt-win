package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QTextLayout__FormatRange_new uintptr
	addr_QTextLayout__FormatRange_new2 uintptr
	addr_QTextLayout__FormatRange_OperatorEqual uintptr
	addr_QTextLayout__FormatRange_OperatorNotEqual uintptr
	addr_QTextLayout__FormatRange_start_get uintptr
	addr_QTextLayout__FormatRange_start_set uintptr
	addr_QTextLayout__FormatRange_length_get uintptr
	addr_QTextLayout__FormatRange_length_set uintptr
	addr_QTextLayout__FormatRange_format_get uintptr
	addr_QTextLayout__FormatRange_format_set uintptr
	pg_QTextLayout__FormatRange_delete func(uintptr)
)

func init() {
	symaddr(&addr_QTextLayout__FormatRange_new, "QTextLayout__FormatRange_new")
	symaddr(&addr_QTextLayout__FormatRange_new2, "QTextLayout__FormatRange_new2")
	symaddr(&addr_QTextLayout__FormatRange_OperatorEqual, "QTextLayout__FormatRange_OperatorEqual")
	symaddr(&addr_QTextLayout__FormatRange_OperatorNotEqual, "QTextLayout__FormatRange_OperatorNotEqual")
	symaddr(&addr_QTextLayout__FormatRange_start_get, "QTextLayout__FormatRange_start_get")
	symaddr(&addr_QTextLayout__FormatRange_start_set, "QTextLayout__FormatRange_start_set")
	symaddr(&addr_QTextLayout__FormatRange_length_get, "QTextLayout__FormatRange_length_get")
	symaddr(&addr_QTextLayout__FormatRange_length_set, "QTextLayout__FormatRange_length_set")
	symaddr(&addr_QTextLayout__FormatRange_format_get, "QTextLayout__FormatRange_format_get")
	symaddr(&addr_QTextLayout__FormatRange_format_set, "QTextLayout__FormatRange_format_set")
	reg(&pg_QTextLayout__FormatRange_delete, "QTextLayout__FormatRange_delete")
}

type QTextLayout__FormatRange struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QTextLayout__FormatRange) UnsafeGoID() int64     { return this.vid }
func (this *QTextLayout__FormatRange) UnsafeSetGoID(id int64) { this.vid = id }

func bindQTextLayout__FormatRange(h uintptr) *QTextLayout__FormatRange {
	if h == 0 {
		return nil
	}
	return &QTextLayout__FormatRange{h: h}
}

func newownedQTextLayout__FormatRange(h uintptr) *QTextLayout__FormatRange {
	o := bindQTextLayout__FormatRange(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QTextLayout__FormatRange).Delete)
	}
	return o
}

func (this *QTextLayout__FormatRange) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QTextLayout__FormatRange) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQTextLayout__FormatRange(h unsafe.Pointer) *QTextLayout__FormatRange { return bindQTextLayout__FormatRange(uintptr(h)) }
func UnsafeNewQTextLayout__FormatRangeOwned(h unsafe.Pointer) *QTextLayout__FormatRange { return newownedQTextLayout__FormatRange(uintptr(h)) }

func (this *QTextLayout__FormatRange) UnsafeRelease() { this.h = 0 }

func NewQTextLayout__FormatRange() *QTextLayout__FormatRange {
	__r, _, _ := purego.SyscallN(addr_QTextLayout__FormatRange_new)
	__o := newownedQTextLayout__FormatRange(__r)
	return __o
}

func NewQTextLayout__FormatRange2(arg1 *QTextLayout__FormatRange) *QTextLayout__FormatRange {
	__r, _, _ := purego.SyscallN(addr_QTextLayout__FormatRange_new2, uintptr(arg1.UnsafePointer()))
	__o := newownedQTextLayout__FormatRange(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QTextLayout__FormatRange) OperatorEqual(rhs *QTextLayout__FormatRange) bool {
	__r0, _, _ := purego.SyscallN(addr_QTextLayout__FormatRange_OperatorEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QTextLayout__FormatRange) OperatorNotEqual(rhs *QTextLayout__FormatRange) bool {
	__r0, _, _ := purego.SyscallN(addr_QTextLayout__FormatRange_OperatorNotEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QTextLayout__FormatRange) Start() int {
	__r0, _, _ := purego.SyscallN(addr_QTextLayout__FormatRange_start_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTextLayout__FormatRange) SetStart(start int) {
	purego.SyscallN(addr_QTextLayout__FormatRange_start_set, this.h, uintptr(start))
	runtime.KeepAlive(this)
}

func (this *QTextLayout__FormatRange) Length() int {
	__r0, _, _ := purego.SyscallN(addr_QTextLayout__FormatRange_length_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTextLayout__FormatRange) SetLength(length int) {
	purego.SyscallN(addr_QTextLayout__FormatRange_length_set, this.h, uintptr(length))
	runtime.KeepAlive(this)
}

func (this *QTextLayout__FormatRange) Format() *QTextCharFormat {
	__r, _, _ := purego.SyscallN(addr_QTextLayout__FormatRange_format_get, this.h)
	runtime.KeepAlive(this)
	return newownedQTextCharFormat(__r)
}

func (this *QTextLayout__FormatRange) SetFormat(format *QTextCharFormat) {
	purego.SyscallN(addr_QTextLayout__FormatRange_format_set, this.h, uintptr(format.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(format)
}

func (this *QTextLayout__FormatRange) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QTextLayout__FormatRange_delete(this.h)
	}
	this.UnsafeRelease()
}
