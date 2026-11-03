package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QWhatsThis_enterWhatsThisMode uintptr
	addr_QWhatsThis_inWhatsThisMode uintptr
	addr_QWhatsThis_leaveWhatsThisMode uintptr
	addr_QWhatsThis_showText uintptr
	addr_QWhatsThis_hideText uintptr
	addr_QWhatsThis_createAction uintptr
	addr_QWhatsThis_showText2 uintptr
	addr_QWhatsThis_createAction2 uintptr
	pg_QWhatsThis_delete func(uintptr)
)

func init() {
	symaddr(&addr_QWhatsThis_enterWhatsThisMode, "QWhatsThis_enterWhatsThisMode")
	symaddr(&addr_QWhatsThis_inWhatsThisMode, "QWhatsThis_inWhatsThisMode")
	symaddr(&addr_QWhatsThis_leaveWhatsThisMode, "QWhatsThis_leaveWhatsThisMode")
	symaddr(&addr_QWhatsThis_showText, "QWhatsThis_showText")
	symaddr(&addr_QWhatsThis_hideText, "QWhatsThis_hideText")
	symaddr(&addr_QWhatsThis_createAction, "QWhatsThis_createAction")
	symaddr(&addr_QWhatsThis_showText2, "QWhatsThis_showText2")
	symaddr(&addr_QWhatsThis_createAction2, "QWhatsThis_createAction2")
	reg(&pg_QWhatsThis_delete, "QWhatsThis_delete")
}

type QWhatsThis struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QWhatsThis) UnsafeGoID() int64     { return this.vid }
func (this *QWhatsThis) UnsafeSetGoID(id int64) { this.vid = id }

func bindQWhatsThis(h uintptr) *QWhatsThis {
	if h == 0 {
		return nil
	}
	return &QWhatsThis{h: h}
}

func newownedQWhatsThis(h uintptr) *QWhatsThis {
	o := bindQWhatsThis(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QWhatsThis).Delete)
	}
	return o
}

func (this *QWhatsThis) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QWhatsThis) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQWhatsThis(h unsafe.Pointer) *QWhatsThis { return bindQWhatsThis(uintptr(h)) }
func UnsafeNewQWhatsThisOwned(h unsafe.Pointer) *QWhatsThis { return newownedQWhatsThis(uintptr(h)) }

func (this *QWhatsThis) UnsafeRelease() { this.h = 0 }

func QWhatsThis_EnterWhatsThisMode() {
	purego.SyscallN(addr_QWhatsThis_enterWhatsThisMode)
}

func QWhatsThis_InWhatsThisMode() bool {
	__r0, _, _ := purego.SyscallN(addr_QWhatsThis_inWhatsThisMode)
	__r := int32(uint32(__r0))
	return __r != 0
}

func QWhatsThis_LeaveWhatsThisMode() {
	purego.SyscallN(addr_QWhatsThis_leaveWhatsThisMode)
}

func QWhatsThis_ShowText(pos *QPoint, text string) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QWhatsThis_showText, uintptr(pos.UnsafePointer()), text_p, uintptr(len(text)))
	runtime.KeepAlive(pos)
	runtime.KeepAlive(text)
}

func QWhatsThis_HideText() {
	purego.SyscallN(addr_QWhatsThis_hideText)
}

func QWhatsThis_CreateAction() *QAction {
	__r, _, _ := purego.SyscallN(addr_QWhatsThis_createAction)
	return bindQAction(__r)
}

func QWhatsThis_ShowText2(pos *QPoint, text string, w *QWidget) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QWhatsThis_showText2, uintptr(pos.UnsafePointer()), text_p, uintptr(len(text)), uintptr(w.UnsafePointer()))
	runtime.KeepAlive(pos)
	runtime.KeepAlive(text)
	runtime.KeepAlive(w)
}

func QWhatsThis_CreateActionWithParent(parent *QObject) *QAction {
	__r, _, _ := purego.SyscallN(addr_QWhatsThis_createAction2, uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(parent)
	return bindQAction(__r)
}

func (this *QWhatsThis) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QWhatsThis_delete(this.h)
	}
	this.UnsafeRelease()
}
