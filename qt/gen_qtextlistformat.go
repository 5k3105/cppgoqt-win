package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QTextListFormat_base func(uintptr) uintptr
	addr_QTextListFormat_new uintptr
	addr_QTextListFormat_new2 uintptr
	addr_QTextListFormat_isValid uintptr
	addr_QTextListFormat_setStyle uintptr
	addr_QTextListFormat_style uintptr
	addr_QTextListFormat_setIndent uintptr
	addr_QTextListFormat_indent uintptr
	addr_QTextListFormat_setNumberPrefix uintptr
	addr_QTextListFormat_numberPrefix uintptr
	addr_QTextListFormat_setNumberSuffix uintptr
	addr_QTextListFormat_numberSuffix uintptr
	addr_QTextListFormat_setStart uintptr
	addr_QTextListFormat_start uintptr
	addr_QTextListFormat_OperatorAssign uintptr
	pg_QTextListFormat_delete func(uintptr)
)

func init() {
	reg(&pg_QTextListFormat_base, "QTextListFormat_base")
	symaddr(&addr_QTextListFormat_new, "QTextListFormat_new")
	symaddr(&addr_QTextListFormat_new2, "QTextListFormat_new2")
	symaddr(&addr_QTextListFormat_isValid, "QTextListFormat_isValid")
	symaddr(&addr_QTextListFormat_setStyle, "QTextListFormat_setStyle")
	symaddr(&addr_QTextListFormat_style, "QTextListFormat_style")
	symaddr(&addr_QTextListFormat_setIndent, "QTextListFormat_setIndent")
	symaddr(&addr_QTextListFormat_indent, "QTextListFormat_indent")
	symaddr(&addr_QTextListFormat_setNumberPrefix, "QTextListFormat_setNumberPrefix")
	symaddr(&addr_QTextListFormat_numberPrefix, "QTextListFormat_numberPrefix")
	symaddr(&addr_QTextListFormat_setNumberSuffix, "QTextListFormat_setNumberSuffix")
	symaddr(&addr_QTextListFormat_numberSuffix, "QTextListFormat_numberSuffix")
	symaddr(&addr_QTextListFormat_setStart, "QTextListFormat_setStart")
	symaddr(&addr_QTextListFormat_start, "QTextListFormat_start")
	symaddr(&addr_QTextListFormat_OperatorAssign, "QTextListFormat_OperatorAssign")
	reg(&pg_QTextListFormat_delete, "QTextListFormat_delete")
}

type QTextListFormat struct {
	h uintptr
	*QTextFormat
}

func bindQTextListFormat(h uintptr) *QTextListFormat {
	if h == 0 {
		return nil
	}
	return &QTextListFormat{h: h, QTextFormat: bindQTextFormat(pg_QTextListFormat_base(h))}
}

func newownedQTextListFormat(h uintptr) *QTextListFormat {
	o := bindQTextListFormat(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QTextListFormat).Delete)
	}
	return o
}

func (this *QTextListFormat) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QTextListFormat) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQTextListFormat(h unsafe.Pointer) *QTextListFormat { return bindQTextListFormat(uintptr(h)) }
func UnsafeNewQTextListFormatOwned(h unsafe.Pointer) *QTextListFormat { return newownedQTextListFormat(uintptr(h)) }

func (this *QTextListFormat) UnsafeRelease() {
	this.h = 0
	if this.QTextFormat != nil {
		this.QTextFormat.UnsafeRelease()
	}
}

func NewQTextListFormat() *QTextListFormat {
	__r, _, _ := purego.SyscallN(addr_QTextListFormat_new)
	__o := newownedQTextListFormat(__r)
	return __o
}

func NewQTextListFormat2(arg1 *QTextListFormat) *QTextListFormat {
	__r, _, _ := purego.SyscallN(addr_QTextListFormat_new2, uintptr(arg1.UnsafePointer()))
	__o := newownedQTextListFormat(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QTextListFormat) IsValid() bool {
	__r0, _, _ := purego.SyscallN(addr_QTextListFormat_isValid, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTextListFormat) SetStyle(style QTextListFormat__Style) {
	purego.SyscallN(addr_QTextListFormat_setStyle, this.h, uintptr(style))
	runtime.KeepAlive(this)
}

func (this *QTextListFormat) Style() QTextListFormat__Style {
	__r0, _, _ := purego.SyscallN(addr_QTextListFormat_style, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QTextListFormat__Style(__r)
}

func (this *QTextListFormat) SetIndent(indent int) {
	purego.SyscallN(addr_QTextListFormat_setIndent, this.h, uintptr(indent))
	runtime.KeepAlive(this)
}

func (this *QTextListFormat) Indent() int {
	__r0, _, _ := purego.SyscallN(addr_QTextListFormat_indent, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTextListFormat) SetNumberPrefix(numberPrefix string) {
	var numberPrefix_p uintptr
	if len(numberPrefix) > 0 {
		numberPrefix_p = uintptr(unsafe.Pointer(unsafe.StringData(numberPrefix)))
	}
	purego.SyscallN(addr_QTextListFormat_setNumberPrefix, this.h, numberPrefix_p, uintptr(len(numberPrefix)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(numberPrefix)
}

func (this *QTextListFormat) NumberPrefix() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTextListFormat_numberPrefix, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QTextListFormat) SetNumberSuffix(numberSuffix string) {
	var numberSuffix_p uintptr
	if len(numberSuffix) > 0 {
		numberSuffix_p = uintptr(unsafe.Pointer(unsafe.StringData(numberSuffix)))
	}
	purego.SyscallN(addr_QTextListFormat_setNumberSuffix, this.h, numberSuffix_p, uintptr(len(numberSuffix)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(numberSuffix)
}

func (this *QTextListFormat) NumberSuffix() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTextListFormat_numberSuffix, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QTextListFormat) SetStart(indent int) {
	purego.SyscallN(addr_QTextListFormat_setStart, this.h, uintptr(indent))
	runtime.KeepAlive(this)
}

func (this *QTextListFormat) Start() int {
	__r0, _, _ := purego.SyscallN(addr_QTextListFormat_start, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTextListFormat) OperatorAssign(arg1 *QTextListFormat) *QTextListFormat {
	__r, _, _ := purego.SyscallN(addr_QTextListFormat_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQTextListFormat(__r)
}

func (this *QTextListFormat) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QTextListFormat_delete(this.h)
	}
	this.UnsafeRelease()
}
