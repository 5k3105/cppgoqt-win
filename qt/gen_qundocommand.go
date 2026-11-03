package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QUndoCommand_new uintptr
	addr_QUndoCommand_new2 uintptr
	addr_QUndoCommand_new3 uintptr
	addr_QUndoCommand_new4 uintptr
	addr_QUndoCommand_undo uintptr
	addr_QUndoCommand_redo uintptr
	addr_QUndoCommand_text uintptr
	addr_QUndoCommand_actionText uintptr
	addr_QUndoCommand_setText uintptr
	addr_QUndoCommand_isObsolete uintptr
	addr_QUndoCommand_setObsolete uintptr
	addr_QUndoCommand_id uintptr
	addr_QUndoCommand_mergeWith uintptr
	addr_QUndoCommand_childCount uintptr
	addr_QUndoCommand_child uintptr
	addr_QUndoCommand_undo_default uintptr
	addr_QUndoCommand_redo_default uintptr
	addr_QUndoCommand_id_default uintptr
	addr_QUndoCommand_mergeWith_default uintptr
	pg_QUndoCommand__set_override func(uintptr, int32)
	pg_QUndoCommand_delete func(uintptr)
)

func init() {
	symaddr(&addr_QUndoCommand_new, "QUndoCommand_new")
	symaddr(&addr_QUndoCommand_new2, "QUndoCommand_new2")
	symaddr(&addr_QUndoCommand_new3, "QUndoCommand_new3")
	symaddr(&addr_QUndoCommand_new4, "QUndoCommand_new4")
	symaddr(&addr_QUndoCommand_undo, "QUndoCommand_undo")
	symaddr(&addr_QUndoCommand_redo, "QUndoCommand_redo")
	symaddr(&addr_QUndoCommand_text, "QUndoCommand_text")
	symaddr(&addr_QUndoCommand_actionText, "QUndoCommand_actionText")
	symaddr(&addr_QUndoCommand_setText, "QUndoCommand_setText")
	symaddr(&addr_QUndoCommand_isObsolete, "QUndoCommand_isObsolete")
	symaddr(&addr_QUndoCommand_setObsolete, "QUndoCommand_setObsolete")
	symaddr(&addr_QUndoCommand_id, "QUndoCommand_id")
	symaddr(&addr_QUndoCommand_mergeWith, "QUndoCommand_mergeWith")
	symaddr(&addr_QUndoCommand_childCount, "QUndoCommand_childCount")
	symaddr(&addr_QUndoCommand_child, "QUndoCommand_child")
	symaddr(&addr_QUndoCommand_undo_default, "QUndoCommand_undo_default")
	symaddr(&addr_QUndoCommand_redo_default, "QUndoCommand_redo_default")
	symaddr(&addr_QUndoCommand_id_default, "QUndoCommand_id_default")
	symaddr(&addr_QUndoCommand_mergeWith_default, "QUndoCommand_mergeWith_default")
	reg(&pg_QUndoCommand__set_override, "QUndoCommand__set_override")
	reg(&pg_QUndoCommand_delete, "QUndoCommand_delete")
}

type QUndoCommand struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QUndoCommand) UnsafeGoID() int64     { return this.vid }
func (this *QUndoCommand) UnsafeSetGoID(id int64) { this.vid = id }

func bindQUndoCommand(h uintptr) *QUndoCommand {
	if h == 0 {
		return nil
	}
	return &QUndoCommand{h: h}
}

func newownedQUndoCommand(h uintptr) *QUndoCommand {
	o := bindQUndoCommand(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QUndoCommand).Delete)
	}
	return o
}

func (this *QUndoCommand) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QUndoCommand) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQUndoCommand(h unsafe.Pointer) *QUndoCommand { return bindQUndoCommand(uintptr(h)) }
func UnsafeNewQUndoCommandOwned(h unsafe.Pointer) *QUndoCommand { return newownedQUndoCommand(uintptr(h)) }

func (this *QUndoCommand) UnsafeRelease() { this.h = 0 }

func NewQUndoCommand() *QUndoCommand {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QUndoCommand_new, uintptr(__goid))
	__o := newownedQUndoCommand(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QUndoCommand_virtIdx, QUndoCommand_enable)
	return __o
}

func NewQUndoCommand2(text string) *QUndoCommand {
	__goid := nextVirtID()
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QUndoCommand_new2, uintptr(__goid), text_p, uintptr(len(text)))
	__o := newownedQUndoCommand(__r)
	runtime.KeepAlive(text)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QUndoCommand_virtIdx, QUndoCommand_enable)
	return __o
}

func NewQUndoCommand3(parent *QUndoCommand) *QUndoCommand {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QUndoCommand_new3, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQUndoCommand(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QUndoCommand_virtIdx, QUndoCommand_enable)
	return __o
}

func NewQUndoCommand4(text string, parent *QUndoCommand) *QUndoCommand {
	__goid := nextVirtID()
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QUndoCommand_new4, uintptr(__goid), text_p, uintptr(len(text)), uintptr(parent.UnsafePointer()))
	__o := newownedQUndoCommand(__r)
	runtime.KeepAlive(text)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QUndoCommand_virtIdx, QUndoCommand_enable)
	return __o
}

func (this *QUndoCommand) Undo() {
	purego.SyscallN(addr_QUndoCommand_undo, this.h)
	runtime.KeepAlive(this)
}

func (this *QUndoCommand) Redo() {
	purego.SyscallN(addr_QUndoCommand_redo, this.h)
	runtime.KeepAlive(this)
}

func (this *QUndoCommand) Text() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QUndoCommand_text, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QUndoCommand) ActionText() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QUndoCommand_actionText, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QUndoCommand) SetText(text string) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QUndoCommand_setText, this.h, text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
}

func (this *QUndoCommand) IsObsolete() bool {
	__r0, _, _ := purego.SyscallN(addr_QUndoCommand_isObsolete, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QUndoCommand) SetObsolete(obsolete bool) {
	purego.SyscallN(addr_QUndoCommand_setObsolete, this.h, uintptr(b2i(obsolete)))
	runtime.KeepAlive(this)
}

func (this *QUndoCommand) Id() int {
	__r0, _, _ := purego.SyscallN(addr_QUndoCommand_id, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QUndoCommand) MergeWith(other *QUndoCommand) bool {
	__r0, _, _ := purego.SyscallN(addr_QUndoCommand_mergeWith, this.h, uintptr(other.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return __r != 0
}

func (this *QUndoCommand) ChildCount() int {
	__r0, _, _ := purego.SyscallN(addr_QUndoCommand_childCount, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QUndoCommand) Child(index int) *QUndoCommand {
	__r, _, _ := purego.SyscallN(addr_QUndoCommand_child, this.h, uintptr(index))
	runtime.KeepAlive(this)
	return bindQUndoCommand(__r)
}

func (this *QUndoCommand) UndoDefault() {
	purego.SyscallN(addr_QUndoCommand_undo_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QUndoCommand) RedoDefault() {
	purego.SyscallN(addr_QUndoCommand_redo_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QUndoCommand) IdDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QUndoCommand_id_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QUndoCommand) MergeWithDefault(other *QUndoCommand) bool {
	__r0, _, _ := purego.SyscallN(addr_QUndoCommand_mergeWith_default, this.h, uintptr(other.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return __r != 0
}

func QUndoCommand_virtIdx(slot int64) int32 {
	switch slot {
	case 7259214242580717359:
		return 0
	case 5544063219051948393:
		return 1
	case 628021283683842752:
		return 2
	case -5894931084898981809:
		return 3
	}
	return -1
}

func QUndoCommand_enable(h uintptr, idx int32) { pg_QUndoCommand__set_override(h, idx) }

func (this *QUndoCommand) OnUndo(fn func()) {
	registerVirt(this.UnsafeGoID(), 7259214242580717359, func() { fn() }) // slot "undo"
}

func (this *QUndoCommand) OnRedo(fn func()) {
	registerVirt(this.UnsafeGoID(), 5544063219051948393, func() { fn() }) // slot "redo"
}

func (this *QUndoCommand) OnId(fn func() int) {
	registerVirt(this.UnsafeGoID(), 628021283683842752, func() int32 { return int32(fn()) }) // slot "id"
}

func (this *QUndoCommand) OnMergeWith(fn func(*QUndoCommand) bool) {
	registerVirt(this.UnsafeGoID(), -5894931084898981809, func(other uintptr) int32 { return b2i(fn(UnsafeNewQUndoCommand(unsafe.Pointer(other)))) }) // slot "mergeWith"
}

func (this *QUndoCommand) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QUndoCommand_delete(this.h)
	}
	this.UnsafeRelease()
}
