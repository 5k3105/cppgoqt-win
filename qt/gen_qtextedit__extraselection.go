package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QTextEdit__ExtraSelection_new uintptr
	addr_QTextEdit__ExtraSelection_OperatorAssign uintptr
	addr_QTextEdit__ExtraSelection_cursor_get uintptr
	addr_QTextEdit__ExtraSelection_cursor_set uintptr
	addr_QTextEdit__ExtraSelection_format_get uintptr
	addr_QTextEdit__ExtraSelection_format_set uintptr
	pg_QTextEdit__ExtraSelection_delete func(uintptr)
)

func init() {
	symaddr(&addr_QTextEdit__ExtraSelection_new, "QTextEdit__ExtraSelection_new")
	symaddr(&addr_QTextEdit__ExtraSelection_OperatorAssign, "QTextEdit__ExtraSelection_OperatorAssign")
	symaddr(&addr_QTextEdit__ExtraSelection_cursor_get, "QTextEdit__ExtraSelection_cursor_get")
	symaddr(&addr_QTextEdit__ExtraSelection_cursor_set, "QTextEdit__ExtraSelection_cursor_set")
	symaddr(&addr_QTextEdit__ExtraSelection_format_get, "QTextEdit__ExtraSelection_format_get")
	symaddr(&addr_QTextEdit__ExtraSelection_format_set, "QTextEdit__ExtraSelection_format_set")
	reg(&pg_QTextEdit__ExtraSelection_delete, "QTextEdit__ExtraSelection_delete")
}

type QTextEdit__ExtraSelection struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QTextEdit__ExtraSelection) UnsafeGoID() int64     { return this.vid }
func (this *QTextEdit__ExtraSelection) UnsafeSetGoID(id int64) { this.vid = id }

func bindQTextEdit__ExtraSelection(h uintptr) *QTextEdit__ExtraSelection {
	if h == 0 {
		return nil
	}
	return &QTextEdit__ExtraSelection{h: h}
}

func newownedQTextEdit__ExtraSelection(h uintptr) *QTextEdit__ExtraSelection {
	o := bindQTextEdit__ExtraSelection(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QTextEdit__ExtraSelection).Delete)
	}
	return o
}

func (this *QTextEdit__ExtraSelection) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QTextEdit__ExtraSelection) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQTextEdit__ExtraSelection(h unsafe.Pointer) *QTextEdit__ExtraSelection { return bindQTextEdit__ExtraSelection(uintptr(h)) }
func UnsafeNewQTextEdit__ExtraSelectionOwned(h unsafe.Pointer) *QTextEdit__ExtraSelection { return newownedQTextEdit__ExtraSelection(uintptr(h)) }

func (this *QTextEdit__ExtraSelection) UnsafeRelease() { this.h = 0 }

func NewQTextEdit__ExtraSelection(arg1 *QTextEdit__ExtraSelection) *QTextEdit__ExtraSelection {
	__r, _, _ := purego.SyscallN(addr_QTextEdit__ExtraSelection_new, uintptr(arg1.UnsafePointer()))
	__o := newownedQTextEdit__ExtraSelection(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QTextEdit__ExtraSelection) OperatorAssign(arg1 *QTextEdit__ExtraSelection) *QTextEdit__ExtraSelection {
	__r, _, _ := purego.SyscallN(addr_QTextEdit__ExtraSelection_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQTextEdit__ExtraSelection(__r)
}

func (this *QTextEdit__ExtraSelection) Cursor() *QTextCursor {
	__r, _, _ := purego.SyscallN(addr_QTextEdit__ExtraSelection_cursor_get, this.h)
	runtime.KeepAlive(this)
	return newownedQTextCursor(__r)
}

func (this *QTextEdit__ExtraSelection) SetCursor(cursor *QTextCursor) {
	purego.SyscallN(addr_QTextEdit__ExtraSelection_cursor_set, this.h, uintptr(cursor.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(cursor)
}

func (this *QTextEdit__ExtraSelection) Format() *QTextCharFormat {
	__r, _, _ := purego.SyscallN(addr_QTextEdit__ExtraSelection_format_get, this.h)
	runtime.KeepAlive(this)
	return newownedQTextCharFormat(__r)
}

func (this *QTextEdit__ExtraSelection) SetFormat(format *QTextCharFormat) {
	purego.SyscallN(addr_QTextEdit__ExtraSelection_format_set, this.h, uintptr(format.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(format)
}

func (this *QTextEdit__ExtraSelection) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QTextEdit__ExtraSelection_delete(this.h)
	}
	this.UnsafeRelease()
}
