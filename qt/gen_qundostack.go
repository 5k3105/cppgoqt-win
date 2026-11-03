package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QUndoStack_base func(uintptr) uintptr
	addr_QUndoStack_new uintptr
	addr_QUndoStack_new2 uintptr
	addr_QUndoStack_tr uintptr
	addr_QUndoStack_clear uintptr
	addr_QUndoStack_push uintptr
	addr_QUndoStack_canUndo uintptr
	addr_QUndoStack_canRedo uintptr
	addr_QUndoStack_undoText uintptr
	addr_QUndoStack_redoText uintptr
	addr_QUndoStack_count uintptr
	addr_QUndoStack_index uintptr
	addr_QUndoStack_text uintptr
	addr_QUndoStack_createUndoAction uintptr
	addr_QUndoStack_createRedoAction uintptr
	addr_QUndoStack_isActive uintptr
	addr_QUndoStack_isClean uintptr
	addr_QUndoStack_cleanIndex uintptr
	addr_QUndoStack_beginMacro uintptr
	addr_QUndoStack_endMacro uintptr
	addr_QUndoStack_setUndoLimit uintptr
	addr_QUndoStack_undoLimit uintptr
	addr_QUndoStack_command uintptr
	addr_QUndoStack_setClean uintptr
	addr_QUndoStack_resetClean uintptr
	addr_QUndoStack_setIndex uintptr
	addr_QUndoStack_undo uintptr
	addr_QUndoStack_redo uintptr
	addr_QUndoStack_setActive uintptr
	addr_QUndoStack_indexChanged uintptr
	addr_QUndoStack_cleanChanged uintptr
	addr_QUndoStack_canUndoChanged uintptr
	addr_QUndoStack_canRedoChanged uintptr
	addr_QUndoStack_undoTextChanged uintptr
	addr_QUndoStack_redoTextChanged uintptr
	addr_QUndoStack_connect uintptr
	addr_QUndoStack_connect2 uintptr
	addr_QUndoStack_sender uintptr
	addr_QUndoStack_senderSignalIndex uintptr
	addr_QUndoStack_receivers uintptr
	addr_QUndoStack_isSignalConnected uintptr
	addr_QUndoStack_tr2 uintptr
	addr_QUndoStack_tr3 uintptr
	addr_QUndoStack_createUndoAction2 uintptr
	addr_QUndoStack_createRedoAction2 uintptr
	addr_QUndoStack_setActive2 uintptr
	addr_QUndoStack_connect3 uintptr
	addr_QUndoStack_connect4 uintptr
	addr_QUndoStack_metaObject_default uintptr
	addr_QUndoStack_event_default uintptr
	addr_QUndoStack_eventFilter_default uintptr
	addr_QUndoStack_timerEvent_default uintptr
	addr_QUndoStack_childEvent_default uintptr
	addr_QUndoStack_customEvent_default uintptr
	addr_QUndoStack_connectNotify_default uintptr
	addr_QUndoStack_disconnectNotify_default uintptr
	pg_QUndoStack_connect_indexChanged func(uintptr, int64)
	pg_QUndoStack_connect_cleanChanged func(uintptr, int64)
	pg_QUndoStack_connect_canUndoChanged func(uintptr, int64)
	pg_QUndoStack_connect_canRedoChanged func(uintptr, int64)
	pg_QUndoStack_connect_undoTextChanged func(uintptr, int64)
	pg_QUndoStack_connect_redoTextChanged func(uintptr, int64)
	pg_QUndoStack_connect_destroyed func(uintptr, int64)
	pg_QUndoStack_connect_destroyed2 func(uintptr, int64)
	pg_QUndoStack_connect_objectNameChanged func(uintptr, int64)
	pg_QUndoStack__disconnect func(int64)
	pg_QUndoStack__set_override func(uintptr, int32)
	pg_QUndoStack_delete func(uintptr)
)

func init() {
	reg(&pg_QUndoStack_base, "QUndoStack_base")
	symaddr(&addr_QUndoStack_new, "QUndoStack_new")
	symaddr(&addr_QUndoStack_new2, "QUndoStack_new2")
	symaddr(&addr_QUndoStack_tr, "QUndoStack_tr")
	symaddr(&addr_QUndoStack_clear, "QUndoStack_clear")
	symaddr(&addr_QUndoStack_push, "QUndoStack_push")
	symaddr(&addr_QUndoStack_canUndo, "QUndoStack_canUndo")
	symaddr(&addr_QUndoStack_canRedo, "QUndoStack_canRedo")
	symaddr(&addr_QUndoStack_undoText, "QUndoStack_undoText")
	symaddr(&addr_QUndoStack_redoText, "QUndoStack_redoText")
	symaddr(&addr_QUndoStack_count, "QUndoStack_count")
	symaddr(&addr_QUndoStack_index, "QUndoStack_index")
	symaddr(&addr_QUndoStack_text, "QUndoStack_text")
	symaddr(&addr_QUndoStack_createUndoAction, "QUndoStack_createUndoAction")
	symaddr(&addr_QUndoStack_createRedoAction, "QUndoStack_createRedoAction")
	symaddr(&addr_QUndoStack_isActive, "QUndoStack_isActive")
	symaddr(&addr_QUndoStack_isClean, "QUndoStack_isClean")
	symaddr(&addr_QUndoStack_cleanIndex, "QUndoStack_cleanIndex")
	symaddr(&addr_QUndoStack_beginMacro, "QUndoStack_beginMacro")
	symaddr(&addr_QUndoStack_endMacro, "QUndoStack_endMacro")
	symaddr(&addr_QUndoStack_setUndoLimit, "QUndoStack_setUndoLimit")
	symaddr(&addr_QUndoStack_undoLimit, "QUndoStack_undoLimit")
	symaddr(&addr_QUndoStack_command, "QUndoStack_command")
	symaddr(&addr_QUndoStack_setClean, "QUndoStack_setClean")
	symaddr(&addr_QUndoStack_resetClean, "QUndoStack_resetClean")
	symaddr(&addr_QUndoStack_setIndex, "QUndoStack_setIndex")
	symaddr(&addr_QUndoStack_undo, "QUndoStack_undo")
	symaddr(&addr_QUndoStack_redo, "QUndoStack_redo")
	symaddr(&addr_QUndoStack_setActive, "QUndoStack_setActive")
	symaddr(&addr_QUndoStack_indexChanged, "QUndoStack_indexChanged")
	symaddr(&addr_QUndoStack_cleanChanged, "QUndoStack_cleanChanged")
	symaddr(&addr_QUndoStack_canUndoChanged, "QUndoStack_canUndoChanged")
	symaddr(&addr_QUndoStack_canRedoChanged, "QUndoStack_canRedoChanged")
	symaddr(&addr_QUndoStack_undoTextChanged, "QUndoStack_undoTextChanged")
	symaddr(&addr_QUndoStack_redoTextChanged, "QUndoStack_redoTextChanged")
	symaddr(&addr_QUndoStack_connect, "QUndoStack_connect")
	symaddr(&addr_QUndoStack_connect2, "QUndoStack_connect2")
	symaddr(&addr_QUndoStack_sender, "QUndoStack_sender")
	symaddr(&addr_QUndoStack_senderSignalIndex, "QUndoStack_senderSignalIndex")
	symaddr(&addr_QUndoStack_receivers, "QUndoStack_receivers")
	symaddr(&addr_QUndoStack_isSignalConnected, "QUndoStack_isSignalConnected")
	symaddr(&addr_QUndoStack_tr2, "QUndoStack_tr2")
	symaddr(&addr_QUndoStack_tr3, "QUndoStack_tr3")
	symaddr(&addr_QUndoStack_createUndoAction2, "QUndoStack_createUndoAction2")
	symaddr(&addr_QUndoStack_createRedoAction2, "QUndoStack_createRedoAction2")
	symaddr(&addr_QUndoStack_setActive2, "QUndoStack_setActive2")
	symaddr(&addr_QUndoStack_connect3, "QUndoStack_connect3")
	symaddr(&addr_QUndoStack_connect4, "QUndoStack_connect4")
	symaddr(&addr_QUndoStack_metaObject_default, "QUndoStack_metaObject_default")
	symaddr(&addr_QUndoStack_event_default, "QUndoStack_event_default")
	symaddr(&addr_QUndoStack_eventFilter_default, "QUndoStack_eventFilter_default")
	symaddr(&addr_QUndoStack_timerEvent_default, "QUndoStack_timerEvent_default")
	symaddr(&addr_QUndoStack_childEvent_default, "QUndoStack_childEvent_default")
	symaddr(&addr_QUndoStack_customEvent_default, "QUndoStack_customEvent_default")
	symaddr(&addr_QUndoStack_connectNotify_default, "QUndoStack_connectNotify_default")
	symaddr(&addr_QUndoStack_disconnectNotify_default, "QUndoStack_disconnectNotify_default")
	reg(&pg_QUndoStack_connect_indexChanged, "QUndoStack_connect_indexChanged")
	reg(&pg_QUndoStack_connect_cleanChanged, "QUndoStack_connect_cleanChanged")
	reg(&pg_QUndoStack_connect_canUndoChanged, "QUndoStack_connect_canUndoChanged")
	reg(&pg_QUndoStack_connect_canRedoChanged, "QUndoStack_connect_canRedoChanged")
	reg(&pg_QUndoStack_connect_undoTextChanged, "QUndoStack_connect_undoTextChanged")
	reg(&pg_QUndoStack_connect_redoTextChanged, "QUndoStack_connect_redoTextChanged")
	reg(&pg_QUndoStack_connect_destroyed, "QUndoStack_connect_destroyed")
	reg(&pg_QUndoStack_connect_destroyed2, "QUndoStack_connect_destroyed2")
	reg(&pg_QUndoStack_connect_objectNameChanged, "QUndoStack_connect_objectNameChanged")
	reg(&pg_QUndoStack__disconnect, "QUndoStack__disconnect")
	reg(&pg_QUndoStack__set_override, "QUndoStack__set_override")
	reg(&pg_QUndoStack_delete, "QUndoStack_delete")
}

type QUndoStack struct {
	h uintptr
	*QObject
}

func bindQUndoStack(h uintptr) *QUndoStack {
	if h == 0 {
		return nil
	}
	return &QUndoStack{h: h, QObject: bindQObject(pg_QUndoStack_base(h))}
}

func newownedQUndoStack(h uintptr) *QUndoStack { return bindQUndoStack(h) }

func (this *QUndoStack) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QUndoStack) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQUndoStack(h unsafe.Pointer) *QUndoStack { return bindQUndoStack(uintptr(h)) }
func UnsafeNewQUndoStackOwned(h unsafe.Pointer) *QUndoStack { return newownedQUndoStack(uintptr(h)) }

func (this *QUndoStack) UnsafeRelease() {
	this.h = 0
	if this.QObject != nil {
		this.QObject.UnsafeRelease()
	}
}

func NewQUndoStack() *QUndoStack {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QUndoStack_new, uintptr(__goid))
	__o := newownedQUndoStack(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QUndoStack_virtIdx, QUndoStack_enable)
	return __o
}

func NewQUndoStack2(parent *QObject) *QUndoStack {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QUndoStack_new2, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQUndoStack(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QUndoStack_virtIdx, QUndoStack_enable)
	return __o
}

func QUndoStack_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QUndoStack_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QUndoStack) Clear() {
	purego.SyscallN(addr_QUndoStack_clear, this.h)
	runtime.KeepAlive(this)
}

func (this *QUndoStack) Push(cmd *QUndoCommand) {
	purego.SyscallN(addr_QUndoStack_push, this.h, uintptr(cmd.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(cmd)
}

func (this *QUndoStack) CanUndo() bool {
	__r0, _, _ := purego.SyscallN(addr_QUndoStack_canUndo, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QUndoStack) CanRedo() bool {
	__r0, _, _ := purego.SyscallN(addr_QUndoStack_canRedo, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QUndoStack) UndoText() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QUndoStack_undoText, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QUndoStack) RedoText() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QUndoStack_redoText, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QUndoStack) Count() int {
	__r0, _, _ := purego.SyscallN(addr_QUndoStack_count, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QUndoStack) Index() int {
	__r0, _, _ := purego.SyscallN(addr_QUndoStack_index, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QUndoStack) Text(idx int) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QUndoStack_text, this.h, uintptr(idx), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QUndoStack) CreateUndoAction(parent *QObject) *QAction {
	__r, _, _ := purego.SyscallN(addr_QUndoStack_createUndoAction, this.h, uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return bindQAction(__r)
}

func (this *QUndoStack) CreateRedoAction(parent *QObject) *QAction {
	__r, _, _ := purego.SyscallN(addr_QUndoStack_createRedoAction, this.h, uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return bindQAction(__r)
}

func (this *QUndoStack) IsActive() bool {
	__r0, _, _ := purego.SyscallN(addr_QUndoStack_isActive, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QUndoStack) IsClean() bool {
	__r0, _, _ := purego.SyscallN(addr_QUndoStack_isClean, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QUndoStack) CleanIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QUndoStack_cleanIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QUndoStack) BeginMacro(text string) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QUndoStack_beginMacro, this.h, text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
}

func (this *QUndoStack) EndMacro() {
	purego.SyscallN(addr_QUndoStack_endMacro, this.h)
	runtime.KeepAlive(this)
}

func (this *QUndoStack) SetUndoLimit(limit int) {
	purego.SyscallN(addr_QUndoStack_setUndoLimit, this.h, uintptr(limit))
	runtime.KeepAlive(this)
}

func (this *QUndoStack) UndoLimit() int {
	__r0, _, _ := purego.SyscallN(addr_QUndoStack_undoLimit, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QUndoStack) Command(index int) *QUndoCommand {
	__r, _, _ := purego.SyscallN(addr_QUndoStack_command, this.h, uintptr(index))
	runtime.KeepAlive(this)
	return bindQUndoCommand(__r)
}

func (this *QUndoStack) SetClean() {
	purego.SyscallN(addr_QUndoStack_setClean, this.h)
	runtime.KeepAlive(this)
}

func (this *QUndoStack) ResetClean() {
	purego.SyscallN(addr_QUndoStack_resetClean, this.h)
	runtime.KeepAlive(this)
}

func (this *QUndoStack) SetIndex(idx int) {
	purego.SyscallN(addr_QUndoStack_setIndex, this.h, uintptr(idx))
	runtime.KeepAlive(this)
}

func (this *QUndoStack) Undo() {
	purego.SyscallN(addr_QUndoStack_undo, this.h)
	runtime.KeepAlive(this)
}

func (this *QUndoStack) Redo() {
	purego.SyscallN(addr_QUndoStack_redo, this.h)
	runtime.KeepAlive(this)
}

func (this *QUndoStack) SetActive() {
	purego.SyscallN(addr_QUndoStack_setActive, this.h)
	runtime.KeepAlive(this)
}

func (this *QUndoStack) IndexChanged(idx int) {
	purego.SyscallN(addr_QUndoStack_indexChanged, this.h, uintptr(idx))
	runtime.KeepAlive(this)
}

func (this *QUndoStack) CleanChanged(clean bool) {
	purego.SyscallN(addr_QUndoStack_cleanChanged, this.h, uintptr(b2i(clean)))
	runtime.KeepAlive(this)
}

func (this *QUndoStack) CanUndoChanged(canUndo bool) {
	purego.SyscallN(addr_QUndoStack_canUndoChanged, this.h, uintptr(b2i(canUndo)))
	runtime.KeepAlive(this)
}

func (this *QUndoStack) CanRedoChanged(canRedo bool) {
	purego.SyscallN(addr_QUndoStack_canRedoChanged, this.h, uintptr(b2i(canRedo)))
	runtime.KeepAlive(this)
}

func (this *QUndoStack) UndoTextChanged(undoText string) {
	var undoText_p uintptr
	if len(undoText) > 0 {
		undoText_p = uintptr(unsafe.Pointer(unsafe.StringData(undoText)))
	}
	purego.SyscallN(addr_QUndoStack_undoTextChanged, this.h, undoText_p, uintptr(len(undoText)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(undoText)
}

func (this *QUndoStack) RedoTextChanged(redoText string) {
	var redoText_p uintptr
	if len(redoText) > 0 {
		redoText_p = uintptr(unsafe.Pointer(unsafe.StringData(redoText)))
	}
	purego.SyscallN(addr_QUndoStack_redoTextChanged, this.h, redoText_p, uintptr(len(redoText)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(redoText)
}

func QUndoStack_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QUndoStack_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QUndoStack_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QUndoStack_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QUndoStack) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QUndoStack_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QUndoStack) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QUndoStack_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QUndoStack) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QUndoStack_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QUndoStack) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QUndoStack_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QUndoStack_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QUndoStack_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QUndoStack_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QUndoStack_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func (this *QUndoStack) CreateUndoAction2(parent *QObject, prefix string) *QAction {
	var prefix_p uintptr
	if len(prefix) > 0 {
		prefix_p = uintptr(unsafe.Pointer(unsafe.StringData(prefix)))
	}
	__r, _, _ := purego.SyscallN(addr_QUndoStack_createUndoAction2, this.h, uintptr(parent.UnsafePointer()), prefix_p, uintptr(len(prefix)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	runtime.KeepAlive(prefix)
	return bindQAction(__r)
}

func (this *QUndoStack) CreateRedoAction2(parent *QObject, prefix string) *QAction {
	var prefix_p uintptr
	if len(prefix) > 0 {
		prefix_p = uintptr(unsafe.Pointer(unsafe.StringData(prefix)))
	}
	__r, _, _ := purego.SyscallN(addr_QUndoStack_createRedoAction2, this.h, uintptr(parent.UnsafePointer()), prefix_p, uintptr(len(prefix)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	runtime.KeepAlive(prefix)
	return bindQAction(__r)
}

func (this *QUndoStack) SetActiveWithActive(active bool) {
	purego.SyscallN(addr_QUndoStack_setActive2, this.h, uintptr(b2i(active)))
	runtime.KeepAlive(this)
}

func QUndoStack_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QUndoStack_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QUndoStack_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QUndoStack_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QUndoStack) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QUndoStack_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QUndoStack) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QUndoStack_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QUndoStack) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QUndoStack_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QUndoStack) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QUndoStack_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QUndoStack) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QUndoStack_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QUndoStack) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QUndoStack_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QUndoStack) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QUndoStack_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QUndoStack) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QUndoStack_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QUndoStack) OnIndexChanged(fn func(int)) int64 {
	id := registerCallback(func(idx int32) { fn(int(idx)) })
	pg_QUndoStack_connect_indexChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QUndoStack) OnCleanChanged(fn func(bool)) int64 {
	id := registerCallback(func(clean int32) { fn(clean != 0) })
	pg_QUndoStack_connect_cleanChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QUndoStack) OnCanUndoChanged(fn func(bool)) int64 {
	id := registerCallback(func(canUndo int32) { fn(canUndo != 0) })
	pg_QUndoStack_connect_canUndoChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QUndoStack) OnCanRedoChanged(fn func(bool)) int64 {
	id := registerCallback(func(canRedo int32) { fn(canRedo != 0) })
	pg_QUndoStack_connect_canRedoChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QUndoStack) OnUndoTextChanged(fn func(string)) int64 {
	id := registerCallback(func(undoText uintptr, undoText_len int64) { fn(pgStrNoFree(undoText, undoText_len)) })
	pg_QUndoStack_connect_undoTextChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QUndoStack) OnRedoTextChanged(fn func(string)) int64 {
	id := registerCallback(func(redoText uintptr, redoText_len int64) { fn(pgStrNoFree(redoText, redoText_len)) })
	pg_QUndoStack_connect_redoTextChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QUndoStack) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QUndoStack_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QUndoStack) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QUndoStack_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QUndoStack) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QUndoStack_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QUndoStack) Disconnect(handle int64) {
	pg_QUndoStack__disconnect(handle)
	unregisterCallback(handle)
}

func QUndoStack_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case 1237236456312824991:
		return 1
	case 5217623514630858297:
		return 2
	case -5409927581812374498:
		return 3
	case -1185315367526857833:
		return 4
	case 3769868229123868474:
		return 5
	case 7379029982170503008:
		return 6
	case -568369772982215962:
		return 7
	}
	return -1
}

func QUndoStack_enable(h uintptr, idx int32) { pg_QUndoStack__set_override(h, idx) }

func (this *QUndoStack) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QUndoStack) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QUndoStack) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QUndoStack) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QUndoStack) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QUndoStack) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QUndoStack) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QUndoStack) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QUndoStack) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QUndoStack_delete(this.h)
	}
	this.UnsafeRelease()
}
