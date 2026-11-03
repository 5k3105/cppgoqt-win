package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QTextDocumentFragment_new uintptr
	addr_QTextDocumentFragment_new2 uintptr
	addr_QTextDocumentFragment_new3 uintptr
	addr_QTextDocumentFragment_new4 uintptr
	addr_QTextDocumentFragment_OperatorAssign uintptr
	addr_QTextDocumentFragment_isEmpty uintptr
	addr_QTextDocumentFragment_toPlainText uintptr
	addr_QTextDocumentFragment_toRawText uintptr
	addr_QTextDocumentFragment_toHtml uintptr
	addr_QTextDocumentFragment_toMarkdown uintptr
	addr_QTextDocumentFragment_fromPlainText uintptr
	addr_QTextDocumentFragment_fromHtml uintptr
	addr_QTextDocumentFragment_fromMarkdown uintptr
	addr_QTextDocumentFragment_toMarkdown2 uintptr
	addr_QTextDocumentFragment_fromHtml2 uintptr
	addr_QTextDocumentFragment_fromMarkdown2 uintptr
	pg_QTextDocumentFragment_delete func(uintptr)
)

func init() {
	symaddr(&addr_QTextDocumentFragment_new, "QTextDocumentFragment_new")
	symaddr(&addr_QTextDocumentFragment_new2, "QTextDocumentFragment_new2")
	symaddr(&addr_QTextDocumentFragment_new3, "QTextDocumentFragment_new3")
	symaddr(&addr_QTextDocumentFragment_new4, "QTextDocumentFragment_new4")
	symaddr(&addr_QTextDocumentFragment_OperatorAssign, "QTextDocumentFragment_OperatorAssign")
	symaddr(&addr_QTextDocumentFragment_isEmpty, "QTextDocumentFragment_isEmpty")
	symaddr(&addr_QTextDocumentFragment_toPlainText, "QTextDocumentFragment_toPlainText")
	symaddr(&addr_QTextDocumentFragment_toRawText, "QTextDocumentFragment_toRawText")
	symaddr(&addr_QTextDocumentFragment_toHtml, "QTextDocumentFragment_toHtml")
	symaddr(&addr_QTextDocumentFragment_toMarkdown, "QTextDocumentFragment_toMarkdown")
	symaddr(&addr_QTextDocumentFragment_fromPlainText, "QTextDocumentFragment_fromPlainText")
	symaddr(&addr_QTextDocumentFragment_fromHtml, "QTextDocumentFragment_fromHtml")
	symaddr(&addr_QTextDocumentFragment_fromMarkdown, "QTextDocumentFragment_fromMarkdown")
	symaddr(&addr_QTextDocumentFragment_toMarkdown2, "QTextDocumentFragment_toMarkdown2")
	symaddr(&addr_QTextDocumentFragment_fromHtml2, "QTextDocumentFragment_fromHtml2")
	symaddr(&addr_QTextDocumentFragment_fromMarkdown2, "QTextDocumentFragment_fromMarkdown2")
	reg(&pg_QTextDocumentFragment_delete, "QTextDocumentFragment_delete")
}

type QTextDocumentFragment struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QTextDocumentFragment) UnsafeGoID() int64     { return this.vid }
func (this *QTextDocumentFragment) UnsafeSetGoID(id int64) { this.vid = id }

func bindQTextDocumentFragment(h uintptr) *QTextDocumentFragment {
	if h == 0 {
		return nil
	}
	return &QTextDocumentFragment{h: h}
}

func newownedQTextDocumentFragment(h uintptr) *QTextDocumentFragment {
	o := bindQTextDocumentFragment(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QTextDocumentFragment).Delete)
	}
	return o
}

func (this *QTextDocumentFragment) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QTextDocumentFragment) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQTextDocumentFragment(h unsafe.Pointer) *QTextDocumentFragment { return bindQTextDocumentFragment(uintptr(h)) }
func UnsafeNewQTextDocumentFragmentOwned(h unsafe.Pointer) *QTextDocumentFragment { return newownedQTextDocumentFragment(uintptr(h)) }

func (this *QTextDocumentFragment) UnsafeRelease() { this.h = 0 }

func NewQTextDocumentFragment() *QTextDocumentFragment {
	__r, _, _ := purego.SyscallN(addr_QTextDocumentFragment_new)
	__o := newownedQTextDocumentFragment(__r)
	return __o
}

func NewQTextDocumentFragment2(document *QTextDocument) *QTextDocumentFragment {
	__r, _, _ := purego.SyscallN(addr_QTextDocumentFragment_new2, uintptr(document.UnsafePointer()))
	__o := newownedQTextDocumentFragment(__r)
	runtime.KeepAlive(document)
	return __o
}

func NewQTextDocumentFragment3(range_ *QTextCursor) *QTextDocumentFragment {
	__r, _, _ := purego.SyscallN(addr_QTextDocumentFragment_new3, uintptr(range_.UnsafePointer()))
	__o := newownedQTextDocumentFragment(__r)
	runtime.KeepAlive(range_)
	return __o
}

func NewQTextDocumentFragment4(rhs *QTextDocumentFragment) *QTextDocumentFragment {
	__r, _, _ := purego.SyscallN(addr_QTextDocumentFragment_new4, uintptr(rhs.UnsafePointer()))
	__o := newownedQTextDocumentFragment(__r)
	runtime.KeepAlive(rhs)
	return __o
}

func (this *QTextDocumentFragment) OperatorAssign(rhs *QTextDocumentFragment) *QTextDocumentFragment {
	__r, _, _ := purego.SyscallN(addr_QTextDocumentFragment_OperatorAssign, this.h, uintptr(rhs.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return bindQTextDocumentFragment(__r)
}

func (this *QTextDocumentFragment) IsEmpty() bool {
	__r0, _, _ := purego.SyscallN(addr_QTextDocumentFragment_isEmpty, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTextDocumentFragment) ToPlainText() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTextDocumentFragment_toPlainText, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QTextDocumentFragment) ToRawText() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTextDocumentFragment_toRawText, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QTextDocumentFragment) ToHtml() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTextDocumentFragment_toHtml, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QTextDocumentFragment) ToMarkdown() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTextDocumentFragment_toMarkdown, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func QTextDocumentFragment_FromPlainText(plainText string) *QTextDocumentFragment {
	var plainText_p uintptr
	if len(plainText) > 0 {
		plainText_p = uintptr(unsafe.Pointer(unsafe.StringData(plainText)))
	}
	__r, _, _ := purego.SyscallN(addr_QTextDocumentFragment_fromPlainText, plainText_p, uintptr(len(plainText)))
	runtime.KeepAlive(plainText)
	return newownedQTextDocumentFragment(__r)
}

func QTextDocumentFragment_FromHtml(html string) *QTextDocumentFragment {
	var html_p uintptr
	if len(html) > 0 {
		html_p = uintptr(unsafe.Pointer(unsafe.StringData(html)))
	}
	__r, _, _ := purego.SyscallN(addr_QTextDocumentFragment_fromHtml, html_p, uintptr(len(html)))
	runtime.KeepAlive(html)
	return newownedQTextDocumentFragment(__r)
}

func QTextDocumentFragment_FromMarkdown(markdown string) *QTextDocumentFragment {
	var markdown_p uintptr
	if len(markdown) > 0 {
		markdown_p = uintptr(unsafe.Pointer(unsafe.StringData(markdown)))
	}
	__r, _, _ := purego.SyscallN(addr_QTextDocumentFragment_fromMarkdown, markdown_p, uintptr(len(markdown)))
	runtime.KeepAlive(markdown)
	return newownedQTextDocumentFragment(__r)
}

func (this *QTextDocumentFragment) ToMarkdownWithFeatures(features QTextDocument__MarkdownFeature) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTextDocumentFragment_toMarkdown2, this.h, uintptr(features), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func QTextDocumentFragment_FromHtml2(html string, resourceProvider *QTextDocument) *QTextDocumentFragment {
	var html_p uintptr
	if len(html) > 0 {
		html_p = uintptr(unsafe.Pointer(unsafe.StringData(html)))
	}
	__r, _, _ := purego.SyscallN(addr_QTextDocumentFragment_fromHtml2, html_p, uintptr(len(html)), uintptr(resourceProvider.UnsafePointer()))
	runtime.KeepAlive(html)
	runtime.KeepAlive(resourceProvider)
	return newownedQTextDocumentFragment(__r)
}

func QTextDocumentFragment_FromMarkdown2(markdown string, features QTextDocument__MarkdownFeature) *QTextDocumentFragment {
	var markdown_p uintptr
	if len(markdown) > 0 {
		markdown_p = uintptr(unsafe.Pointer(unsafe.StringData(markdown)))
	}
	__r, _, _ := purego.SyscallN(addr_QTextDocumentFragment_fromMarkdown2, markdown_p, uintptr(len(markdown)), uintptr(features))
	runtime.KeepAlive(markdown)
	return newownedQTextDocumentFragment(__r)
}

func (this *QTextDocumentFragment) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QTextDocumentFragment_delete(this.h)
	}
	this.UnsafeRelease()
}
