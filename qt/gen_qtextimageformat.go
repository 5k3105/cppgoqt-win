package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QTextImageFormat_chain uintptr
	addr_QTextImageFormat_new uintptr
	addr_QTextImageFormat_new2 uintptr
	addr_QTextImageFormat_isValid uintptr
	addr_QTextImageFormat_setName uintptr
	addr_QTextImageFormat_name uintptr
	pg_QTextImageFormat_setWidth func(uintptr, float64)
	pg_QTextImageFormat_width func(uintptr) float64
	addr_QTextImageFormat_setMaximumWidth uintptr
	addr_QTextImageFormat_maximumWidth uintptr
	pg_QTextImageFormat_setHeight func(uintptr, float64)
	pg_QTextImageFormat_height func(uintptr) float64
	addr_QTextImageFormat_setQuality uintptr
	addr_QTextImageFormat_setQuality2 uintptr
	addr_QTextImageFormat_quality uintptr
	pg_QTextImageFormat_delete func(uintptr)
)

func init() {
	symaddr(&addr_QTextImageFormat_chain, "QTextImageFormat_chain")
	symaddr(&addr_QTextImageFormat_new, "QTextImageFormat_new")
	symaddr(&addr_QTextImageFormat_new2, "QTextImageFormat_new2")
	symaddr(&addr_QTextImageFormat_isValid, "QTextImageFormat_isValid")
	symaddr(&addr_QTextImageFormat_setName, "QTextImageFormat_setName")
	symaddr(&addr_QTextImageFormat_name, "QTextImageFormat_name")
	reg(&pg_QTextImageFormat_setWidth, "QTextImageFormat_setWidth")
	reg(&pg_QTextImageFormat_width, "QTextImageFormat_width")
	symaddr(&addr_QTextImageFormat_setMaximumWidth, "QTextImageFormat_setMaximumWidth")
	symaddr(&addr_QTextImageFormat_maximumWidth, "QTextImageFormat_maximumWidth")
	reg(&pg_QTextImageFormat_setHeight, "QTextImageFormat_setHeight")
	reg(&pg_QTextImageFormat_height, "QTextImageFormat_height")
	symaddr(&addr_QTextImageFormat_setQuality, "QTextImageFormat_setQuality")
	symaddr(&addr_QTextImageFormat_setQuality2, "QTextImageFormat_setQuality2")
	symaddr(&addr_QTextImageFormat_quality, "QTextImageFormat_quality")
	reg(&pg_QTextImageFormat_delete, "QTextImageFormat_delete")
}

type QTextImageFormat struct {
	h uintptr
	*QTextCharFormat
}

func bindQTextImageFormat(h uintptr) *QTextImageFormat {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QTextImageFormat_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QTextImageFormat{h: h, QTextCharFormat: &QTextCharFormat{h: __c[0], QTextFormat: &QTextFormat{h: __c[1]}}}
}

func newownedQTextImageFormat(h uintptr) *QTextImageFormat {
	o := bindQTextImageFormat(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QTextImageFormat).Delete)
	}
	return o
}

func (this *QTextImageFormat) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QTextImageFormat) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQTextImageFormat(h unsafe.Pointer) *QTextImageFormat { return bindQTextImageFormat(uintptr(h)) }
func UnsafeNewQTextImageFormatOwned(h unsafe.Pointer) *QTextImageFormat { return newownedQTextImageFormat(uintptr(h)) }

func (this *QTextImageFormat) UnsafeRelease() {
	this.h = 0
	if this.QTextCharFormat != nil {
		this.QTextCharFormat.UnsafeRelease()
	}
}

func NewQTextImageFormat() *QTextImageFormat {
	__r, _, _ := purego.SyscallN(addr_QTextImageFormat_new)
	__o := newownedQTextImageFormat(__r)
	return __o
}

func NewQTextImageFormat2(arg1 *QTextImageFormat) *QTextImageFormat {
	__r, _, _ := purego.SyscallN(addr_QTextImageFormat_new2, uintptr(arg1.UnsafePointer()))
	__o := newownedQTextImageFormat(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QTextImageFormat) IsValid() bool {
	__r0, _, _ := purego.SyscallN(addr_QTextImageFormat_isValid, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTextImageFormat) SetName(name string) {
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	purego.SyscallN(addr_QTextImageFormat_setName, this.h, name_p, uintptr(len(name)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(name)
}

func (this *QTextImageFormat) Name() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTextImageFormat_name, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QTextImageFormat) SetWidth(width float64) {
	pg_QTextImageFormat_setWidth(this.h, width)
	runtime.KeepAlive(this)
}

func (this *QTextImageFormat) Width() float64 {
	__r := pg_QTextImageFormat_width(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QTextImageFormat) SetMaximumWidth(maxWidth *QTextLength) {
	purego.SyscallN(addr_QTextImageFormat_setMaximumWidth, this.h, uintptr(maxWidth.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(maxWidth)
}

func (this *QTextImageFormat) MaximumWidth() *QTextLength {
	__r, _, _ := purego.SyscallN(addr_QTextImageFormat_maximumWidth, this.h)
	runtime.KeepAlive(this)
	return newownedQTextLength(__r)
}

func (this *QTextImageFormat) SetHeight(height float64) {
	pg_QTextImageFormat_setHeight(this.h, height)
	runtime.KeepAlive(this)
}

func (this *QTextImageFormat) Height() float64 {
	__r := pg_QTextImageFormat_height(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QTextImageFormat) SetQuality(quality int) {
	purego.SyscallN(addr_QTextImageFormat_setQuality, this.h, uintptr(quality))
	runtime.KeepAlive(this)
}

func (this *QTextImageFormat) SetQuality2() {
	purego.SyscallN(addr_QTextImageFormat_setQuality2, this.h)
	runtime.KeepAlive(this)
}

func (this *QTextImageFormat) Quality() int {
	__r0, _, _ := purego.SyscallN(addr_QTextImageFormat_quality, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTextImageFormat) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QTextImageFormat_delete(this.h)
	}
	this.UnsafeRelease()
}
