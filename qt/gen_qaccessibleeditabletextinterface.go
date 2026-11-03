package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QAccessibleEditableTextInterface_deleteText uintptr
	addr_QAccessibleEditableTextInterface_insertText uintptr
	addr_QAccessibleEditableTextInterface_replaceText uintptr
	addr_QAccessibleEditableTextInterface_OperatorAssign uintptr
	pg_QAccessibleEditableTextInterface__set_override func(uintptr, int32)
	pg_QAccessibleEditableTextInterface__Go_new func(int64) uintptr
	pg_QAccessibleEditableTextInterface_delete func(uintptr)
)

func init() {
	symaddr(&addr_QAccessibleEditableTextInterface_deleteText, "QAccessibleEditableTextInterface_deleteText")
	symaddr(&addr_QAccessibleEditableTextInterface_insertText, "QAccessibleEditableTextInterface_insertText")
	symaddr(&addr_QAccessibleEditableTextInterface_replaceText, "QAccessibleEditableTextInterface_replaceText")
	symaddr(&addr_QAccessibleEditableTextInterface_OperatorAssign, "QAccessibleEditableTextInterface_OperatorAssign")
	reg(&pg_QAccessibleEditableTextInterface__set_override, "QAccessibleEditableTextInterface__set_override")
	reg(&pg_QAccessibleEditableTextInterface__Go_new, "QAccessibleEditableTextInterface__Go_new")
	reg(&pg_QAccessibleEditableTextInterface_delete, "QAccessibleEditableTextInterface_delete")
}

type QAccessibleEditableTextInterface struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QAccessibleEditableTextInterface) UnsafeGoID() int64     { return this.vid }
func (this *QAccessibleEditableTextInterface) UnsafeSetGoID(id int64) { this.vid = id }

func bindQAccessibleEditableTextInterface(h uintptr) *QAccessibleEditableTextInterface {
	if h == 0 {
		return nil
	}
	return &QAccessibleEditableTextInterface{h: h}
}

func newownedQAccessibleEditableTextInterface(h uintptr) *QAccessibleEditableTextInterface {
	o := bindQAccessibleEditableTextInterface(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QAccessibleEditableTextInterface).Delete)
	}
	return o
}

func (this *QAccessibleEditableTextInterface) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QAccessibleEditableTextInterface) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQAccessibleEditableTextInterface(h unsafe.Pointer) *QAccessibleEditableTextInterface { return bindQAccessibleEditableTextInterface(uintptr(h)) }
func UnsafeNewQAccessibleEditableTextInterfaceOwned(h unsafe.Pointer) *QAccessibleEditableTextInterface { return newownedQAccessibleEditableTextInterface(uintptr(h)) }

func (this *QAccessibleEditableTextInterface) UnsafeRelease() { this.h = 0 }

func (this *QAccessibleEditableTextInterface) DeleteText(startOffset int, endOffset int) {
	purego.SyscallN(addr_QAccessibleEditableTextInterface_deleteText, this.h, uintptr(startOffset), uintptr(endOffset))
	runtime.KeepAlive(this)
}

func (this *QAccessibleEditableTextInterface) InsertText(offset int, text string) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QAccessibleEditableTextInterface_insertText, this.h, uintptr(offset), text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
}

func (this *QAccessibleEditableTextInterface) ReplaceText(startOffset int, endOffset int, text string) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QAccessibleEditableTextInterface_replaceText, this.h, uintptr(startOffset), uintptr(endOffset), text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
}

func (this *QAccessibleEditableTextInterface) OperatorAssign(arg1 *QAccessibleEditableTextInterface) *QAccessibleEditableTextInterface {
	__r, _, _ := purego.SyscallN(addr_QAccessibleEditableTextInterface_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQAccessibleEditableTextInterface(__r)
}

func QAccessibleEditableTextInterface_virtIdx(slot int64) int32 {
	switch slot {
	case 4554921347280463265:
		return 0
	case -7342818475556881681:
		return 1
	case 7391831930358964452:
		return 2
	}
	return -1
}

func QAccessibleEditableTextInterface_enable(h uintptr, idx int32) { pg_QAccessibleEditableTextInterface__set_override(h, idx) }

func NewQAccessibleEditableTextInterface() *QAccessibleEditableTextInterface {
	id := nextVirtID()
	o := newownedQAccessibleEditableTextInterface(pg_QAccessibleEditableTextInterface__Go_new(id))
	o.UnsafeSetGoID(id)
	registerInstance(id, o.h, QAccessibleEditableTextInterface_virtIdx, QAccessibleEditableTextInterface_enable)
	return o
}

func (this *QAccessibleEditableTextInterface) OnDeleteText(fn func(int, int)) {
	registerVirt(this.UnsafeGoID(), 4554921347280463265, func(startOffset int32, endOffset int32) { fn(int(startOffset), int(endOffset)) }) // slot "deleteText"
}

func (this *QAccessibleEditableTextInterface) OnInsertText(fn func(int, string)) {
	registerVirt(this.UnsafeGoID(), -7342818475556881681, func(offset int32, text uintptr, text_len int64) { fn(int(offset), pgStrNoFree(text, text_len)) }) // slot "insertText"
}

func (this *QAccessibleEditableTextInterface) OnReplaceText(fn func(int, int, string)) {
	registerVirt(this.UnsafeGoID(), 7391831930358964452, func(startOffset int32, endOffset int32, text uintptr, text_len int64) { fn(int(startOffset), int(endOffset), pgStrNoFree(text, text_len)) }) // slot "replaceText"
}

func (this *QAccessibleEditableTextInterface) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QAccessibleEditableTextInterface_delete(this.h)
	}
	this.UnsafeRelease()
}
