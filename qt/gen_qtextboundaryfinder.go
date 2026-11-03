package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QTextBoundaryFinder_new uintptr
	addr_QTextBoundaryFinder_new2 uintptr
	addr_QTextBoundaryFinder_new3 uintptr
	addr_QTextBoundaryFinder_new4 uintptr
	addr_QTextBoundaryFinder_new5 uintptr
	addr_QTextBoundaryFinder_new6 uintptr
	addr_QTextBoundaryFinder_new7 uintptr
	addr_QTextBoundaryFinder_new8 uintptr
	addr_QTextBoundaryFinder_new9 uintptr
	addr_QTextBoundaryFinder_OperatorAssign uintptr
	addr_QTextBoundaryFinder_swap uintptr
	addr_QTextBoundaryFinder_isValid uintptr
	addr_QTextBoundaryFinder_type uintptr
	addr_QTextBoundaryFinder_string uintptr
	addr_QTextBoundaryFinder_toStart uintptr
	addr_QTextBoundaryFinder_toEnd uintptr
	addr_QTextBoundaryFinder_position uintptr
	addr_QTextBoundaryFinder_setPosition uintptr
	addr_QTextBoundaryFinder_toNextBoundary uintptr
	addr_QTextBoundaryFinder_toPreviousBoundary uintptr
	addr_QTextBoundaryFinder_isAtBoundary uintptr
	addr_QTextBoundaryFinder_boundaryReasons uintptr
	pg_QTextBoundaryFinder_delete func(uintptr)
)

func init() {
	symaddr(&addr_QTextBoundaryFinder_new, "QTextBoundaryFinder_new")
	symaddr(&addr_QTextBoundaryFinder_new2, "QTextBoundaryFinder_new2")
	symaddr(&addr_QTextBoundaryFinder_new3, "QTextBoundaryFinder_new3")
	symaddr(&addr_QTextBoundaryFinder_new4, "QTextBoundaryFinder_new4")
	symaddr(&addr_QTextBoundaryFinder_new5, "QTextBoundaryFinder_new5")
	symaddr(&addr_QTextBoundaryFinder_new6, "QTextBoundaryFinder_new6")
	symaddr(&addr_QTextBoundaryFinder_new7, "QTextBoundaryFinder_new7")
	symaddr(&addr_QTextBoundaryFinder_new8, "QTextBoundaryFinder_new8")
	symaddr(&addr_QTextBoundaryFinder_new9, "QTextBoundaryFinder_new9")
	symaddr(&addr_QTextBoundaryFinder_OperatorAssign, "QTextBoundaryFinder_OperatorAssign")
	symaddr(&addr_QTextBoundaryFinder_swap, "QTextBoundaryFinder_swap")
	symaddr(&addr_QTextBoundaryFinder_isValid, "QTextBoundaryFinder_isValid")
	symaddr(&addr_QTextBoundaryFinder_type, "QTextBoundaryFinder_type")
	symaddr(&addr_QTextBoundaryFinder_string, "QTextBoundaryFinder_string")
	symaddr(&addr_QTextBoundaryFinder_toStart, "QTextBoundaryFinder_toStart")
	symaddr(&addr_QTextBoundaryFinder_toEnd, "QTextBoundaryFinder_toEnd")
	symaddr(&addr_QTextBoundaryFinder_position, "QTextBoundaryFinder_position")
	symaddr(&addr_QTextBoundaryFinder_setPosition, "QTextBoundaryFinder_setPosition")
	symaddr(&addr_QTextBoundaryFinder_toNextBoundary, "QTextBoundaryFinder_toNextBoundary")
	symaddr(&addr_QTextBoundaryFinder_toPreviousBoundary, "QTextBoundaryFinder_toPreviousBoundary")
	symaddr(&addr_QTextBoundaryFinder_isAtBoundary, "QTextBoundaryFinder_isAtBoundary")
	symaddr(&addr_QTextBoundaryFinder_boundaryReasons, "QTextBoundaryFinder_boundaryReasons")
	reg(&pg_QTextBoundaryFinder_delete, "QTextBoundaryFinder_delete")
}

type QTextBoundaryFinder struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QTextBoundaryFinder) UnsafeGoID() int64     { return this.vid }
func (this *QTextBoundaryFinder) UnsafeSetGoID(id int64) { this.vid = id }

func bindQTextBoundaryFinder(h uintptr) *QTextBoundaryFinder {
	if h == 0 {
		return nil
	}
	return &QTextBoundaryFinder{h: h}
}

func newownedQTextBoundaryFinder(h uintptr) *QTextBoundaryFinder {
	o := bindQTextBoundaryFinder(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QTextBoundaryFinder).Delete)
	}
	return o
}

func (this *QTextBoundaryFinder) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QTextBoundaryFinder) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQTextBoundaryFinder(h unsafe.Pointer) *QTextBoundaryFinder { return bindQTextBoundaryFinder(uintptr(h)) }
func UnsafeNewQTextBoundaryFinderOwned(h unsafe.Pointer) *QTextBoundaryFinder { return newownedQTextBoundaryFinder(uintptr(h)) }

func (this *QTextBoundaryFinder) UnsafeRelease() { this.h = 0 }

func NewQTextBoundaryFinder() *QTextBoundaryFinder {
	__r, _, _ := purego.SyscallN(addr_QTextBoundaryFinder_new)
	__o := newownedQTextBoundaryFinder(__r)
	return __o
}

func NewQTextBoundaryFinder2(other *QTextBoundaryFinder) *QTextBoundaryFinder {
	__r, _, _ := purego.SyscallN(addr_QTextBoundaryFinder_new2, uintptr(other.UnsafePointer()))
	__o := newownedQTextBoundaryFinder(__r)
	runtime.KeepAlive(other)
	return __o
}

func NewQTextBoundaryFinder3(type_ QTextBoundaryFinder__BoundaryType, string_ string) *QTextBoundaryFinder {
	var string__p uintptr
	if len(string_) > 0 {
		string__p = uintptr(unsafe.Pointer(unsafe.StringData(string_)))
	}
	__r, _, _ := purego.SyscallN(addr_QTextBoundaryFinder_new3, uintptr(type_), string__p, uintptr(len(string_)))
	__o := newownedQTextBoundaryFinder(__r)
	runtime.KeepAlive(string_)
	return __o
}

func NewQTextBoundaryFinder4(type_ QTextBoundaryFinder__BoundaryType, chars *QChar, length int64) *QTextBoundaryFinder {
	__r, _, _ := purego.SyscallN(addr_QTextBoundaryFinder_new4, uintptr(type_), uintptr(chars.UnsafePointer()), uintptr(length))
	__o := newownedQTextBoundaryFinder(__r)
	runtime.KeepAlive(chars)
	return __o
}

func NewQTextBoundaryFinder5(type_ QTextBoundaryFinder__BoundaryType, str *QStringView) *QTextBoundaryFinder {
	__r, _, _ := purego.SyscallN(addr_QTextBoundaryFinder_new5, uintptr(type_), uintptr(str.UnsafePointer()))
	__o := newownedQTextBoundaryFinder(__r)
	runtime.KeepAlive(str)
	return __o
}

func NewQTextBoundaryFinder6(type_ QTextBoundaryFinder__BoundaryType, chars *QChar, length int64, buffer []byte) *QTextBoundaryFinder {
	var buffer_p uintptr
	if len(buffer) > 0 {
		buffer_p = uintptr(unsafe.Pointer(&buffer[0]))
	}
	__r, _, _ := purego.SyscallN(addr_QTextBoundaryFinder_new6, uintptr(type_), uintptr(chars.UnsafePointer()), uintptr(length), buffer_p)
	__o := newownedQTextBoundaryFinder(__r)
	runtime.KeepAlive(chars)
	runtime.KeepAlive(buffer)
	return __o
}

func NewQTextBoundaryFinder7(type_ QTextBoundaryFinder__BoundaryType, chars *QChar, length int64, buffer []byte, bufferSize int64) *QTextBoundaryFinder {
	var buffer_p uintptr
	if len(buffer) > 0 {
		buffer_p = uintptr(unsafe.Pointer(&buffer[0]))
	}
	__r, _, _ := purego.SyscallN(addr_QTextBoundaryFinder_new7, uintptr(type_), uintptr(chars.UnsafePointer()), uintptr(length), buffer_p, uintptr(bufferSize))
	__o := newownedQTextBoundaryFinder(__r)
	runtime.KeepAlive(chars)
	runtime.KeepAlive(buffer)
	return __o
}

func NewQTextBoundaryFinder8(type_ QTextBoundaryFinder__BoundaryType, str *QStringView, buffer []byte) *QTextBoundaryFinder {
	var buffer_p uintptr
	if len(buffer) > 0 {
		buffer_p = uintptr(unsafe.Pointer(&buffer[0]))
	}
	__r, _, _ := purego.SyscallN(addr_QTextBoundaryFinder_new8, uintptr(type_), uintptr(str.UnsafePointer()), buffer_p)
	__o := newownedQTextBoundaryFinder(__r)
	runtime.KeepAlive(str)
	runtime.KeepAlive(buffer)
	return __o
}

func NewQTextBoundaryFinder9(type_ QTextBoundaryFinder__BoundaryType, str *QStringView, buffer []byte, bufferSize int64) *QTextBoundaryFinder {
	var buffer_p uintptr
	if len(buffer) > 0 {
		buffer_p = uintptr(unsafe.Pointer(&buffer[0]))
	}
	__r, _, _ := purego.SyscallN(addr_QTextBoundaryFinder_new9, uintptr(type_), uintptr(str.UnsafePointer()), buffer_p, uintptr(bufferSize))
	__o := newownedQTextBoundaryFinder(__r)
	runtime.KeepAlive(str)
	runtime.KeepAlive(buffer)
	return __o
}

func (this *QTextBoundaryFinder) OperatorAssign(other *QTextBoundaryFinder) *QTextBoundaryFinder {
	__r, _, _ := purego.SyscallN(addr_QTextBoundaryFinder_OperatorAssign, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return bindQTextBoundaryFinder(__r)
}

func (this *QTextBoundaryFinder) Swap(other *QTextBoundaryFinder) {
	purego.SyscallN(addr_QTextBoundaryFinder_swap, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
}

func (this *QTextBoundaryFinder) IsValid() bool {
	__r0, _, _ := purego.SyscallN(addr_QTextBoundaryFinder_isValid, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTextBoundaryFinder) Type() QTextBoundaryFinder__BoundaryType {
	__r0, _, _ := purego.SyscallN(addr_QTextBoundaryFinder_type, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QTextBoundaryFinder__BoundaryType(__r)
}

func (this *QTextBoundaryFinder) String() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTextBoundaryFinder_string, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QTextBoundaryFinder) ToStart() {
	purego.SyscallN(addr_QTextBoundaryFinder_toStart, this.h)
	runtime.KeepAlive(this)
}

func (this *QTextBoundaryFinder) ToEnd() {
	purego.SyscallN(addr_QTextBoundaryFinder_toEnd, this.h)
	runtime.KeepAlive(this)
}

func (this *QTextBoundaryFinder) Position() int64 {
	__r0, _, _ := purego.SyscallN(addr_QTextBoundaryFinder_position, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QTextBoundaryFinder) SetPosition(position int64) {
	purego.SyscallN(addr_QTextBoundaryFinder_setPosition, this.h, uintptr(position))
	runtime.KeepAlive(this)
}

func (this *QTextBoundaryFinder) ToNextBoundary() int64 {
	__r0, _, _ := purego.SyscallN(addr_QTextBoundaryFinder_toNextBoundary, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QTextBoundaryFinder) ToPreviousBoundary() int64 {
	__r0, _, _ := purego.SyscallN(addr_QTextBoundaryFinder_toPreviousBoundary, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QTextBoundaryFinder) IsAtBoundary() bool {
	__r0, _, _ := purego.SyscallN(addr_QTextBoundaryFinder_isAtBoundary, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTextBoundaryFinder) BoundaryReasons() QTextBoundaryFinder__BoundaryReason {
	__r0, _, _ := purego.SyscallN(addr_QTextBoundaryFinder_boundaryReasons, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QTextBoundaryFinder__BoundaryReason(__r)
}

func (this *QTextBoundaryFinder) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QTextBoundaryFinder_delete(this.h)
	}
	this.UnsafeRelease()
}
