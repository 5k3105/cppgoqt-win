package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QAbstractTextDocumentLayout__Selection_new uintptr
	addr_QAbstractTextDocumentLayout__Selection_new2 uintptr
	addr_QAbstractTextDocumentLayout__Selection_OperatorAssign uintptr
	addr_QAbstractTextDocumentLayout__Selection_cursor_get uintptr
	addr_QAbstractTextDocumentLayout__Selection_cursor_set uintptr
	addr_QAbstractTextDocumentLayout__Selection_format_get uintptr
	addr_QAbstractTextDocumentLayout__Selection_format_set uintptr
	pg_QAbstractTextDocumentLayout__Selection_delete func(uintptr)
)

func init() {
	symaddr(&addr_QAbstractTextDocumentLayout__Selection_new, "QAbstractTextDocumentLayout__Selection_new")
	symaddr(&addr_QAbstractTextDocumentLayout__Selection_new2, "QAbstractTextDocumentLayout__Selection_new2")
	symaddr(&addr_QAbstractTextDocumentLayout__Selection_OperatorAssign, "QAbstractTextDocumentLayout__Selection_OperatorAssign")
	symaddr(&addr_QAbstractTextDocumentLayout__Selection_cursor_get, "QAbstractTextDocumentLayout__Selection_cursor_get")
	symaddr(&addr_QAbstractTextDocumentLayout__Selection_cursor_set, "QAbstractTextDocumentLayout__Selection_cursor_set")
	symaddr(&addr_QAbstractTextDocumentLayout__Selection_format_get, "QAbstractTextDocumentLayout__Selection_format_get")
	symaddr(&addr_QAbstractTextDocumentLayout__Selection_format_set, "QAbstractTextDocumentLayout__Selection_format_set")
	reg(&pg_QAbstractTextDocumentLayout__Selection_delete, "QAbstractTextDocumentLayout__Selection_delete")
}

type QAbstractTextDocumentLayout__Selection struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QAbstractTextDocumentLayout__Selection) UnsafeGoID() int64     { return this.vid }
func (this *QAbstractTextDocumentLayout__Selection) UnsafeSetGoID(id int64) { this.vid = id }

func bindQAbstractTextDocumentLayout__Selection(h uintptr) *QAbstractTextDocumentLayout__Selection {
	if h == 0 {
		return nil
	}
	return &QAbstractTextDocumentLayout__Selection{h: h}
}

func newownedQAbstractTextDocumentLayout__Selection(h uintptr) *QAbstractTextDocumentLayout__Selection {
	o := bindQAbstractTextDocumentLayout__Selection(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QAbstractTextDocumentLayout__Selection).Delete)
	}
	return o
}

func (this *QAbstractTextDocumentLayout__Selection) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QAbstractTextDocumentLayout__Selection) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQAbstractTextDocumentLayout__Selection(h unsafe.Pointer) *QAbstractTextDocumentLayout__Selection { return bindQAbstractTextDocumentLayout__Selection(uintptr(h)) }
func UnsafeNewQAbstractTextDocumentLayout__SelectionOwned(h unsafe.Pointer) *QAbstractTextDocumentLayout__Selection { return newownedQAbstractTextDocumentLayout__Selection(uintptr(h)) }

func (this *QAbstractTextDocumentLayout__Selection) UnsafeRelease() { this.h = 0 }

func NewQAbstractTextDocumentLayout__Selection(arg1 *QAbstractTextDocumentLayout__Selection) *QAbstractTextDocumentLayout__Selection {
	__r, _, _ := purego.SyscallN(addr_QAbstractTextDocumentLayout__Selection_new, uintptr(arg1.UnsafePointer()))
	__o := newownedQAbstractTextDocumentLayout__Selection(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func NewQAbstractTextDocumentLayout__Selection2() *QAbstractTextDocumentLayout__Selection {
	__r, _, _ := purego.SyscallN(addr_QAbstractTextDocumentLayout__Selection_new2)
	__o := newownedQAbstractTextDocumentLayout__Selection(__r)
	return __o
}

func (this *QAbstractTextDocumentLayout__Selection) OperatorAssign(arg1 *QAbstractTextDocumentLayout__Selection) *QAbstractTextDocumentLayout__Selection {
	__r, _, _ := purego.SyscallN(addr_QAbstractTextDocumentLayout__Selection_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQAbstractTextDocumentLayout__Selection(__r)
}

func (this *QAbstractTextDocumentLayout__Selection) Cursor() *QTextCursor {
	__r, _, _ := purego.SyscallN(addr_QAbstractTextDocumentLayout__Selection_cursor_get, this.h)
	runtime.KeepAlive(this)
	return newownedQTextCursor(__r)
}

func (this *QAbstractTextDocumentLayout__Selection) SetCursor(cursor *QTextCursor) {
	purego.SyscallN(addr_QAbstractTextDocumentLayout__Selection_cursor_set, this.h, uintptr(cursor.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(cursor)
}

func (this *QAbstractTextDocumentLayout__Selection) Format() *QTextCharFormat {
	__r, _, _ := purego.SyscallN(addr_QAbstractTextDocumentLayout__Selection_format_get, this.h)
	runtime.KeepAlive(this)
	return newownedQTextCharFormat(__r)
}

func (this *QAbstractTextDocumentLayout__Selection) SetFormat(format *QTextCharFormat) {
	purego.SyscallN(addr_QAbstractTextDocumentLayout__Selection_format_set, this.h, uintptr(format.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(format)
}

func (this *QAbstractTextDocumentLayout__Selection) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QAbstractTextDocumentLayout__Selection_delete(this.h)
	}
	this.UnsafeRelease()
}
