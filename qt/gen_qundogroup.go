package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QUndoGroup_base func(uintptr) uintptr
	addr_QUndoGroup_new uintptr
	addr_QUndoGroup_new2 uintptr
	addr_QUndoGroup_tr uintptr
	addr_QUndoGroup_addStack uintptr
	addr_QUndoGroup_removeStack uintptr
	addr_QUndoGroup_stacks uintptr
	addr_QUndoGroup_activeStack uintptr
	addr_QUndoGroup_createUndoAction uintptr
	addr_QUndoGroup_createRedoAction uintptr
	addr_QUndoGroup_canUndo uintptr
	addr_QUndoGroup_canRedo uintptr
	addr_QUndoGroup_undoText uintptr
	addr_QUndoGroup_redoText uintptr
	addr_QUndoGroup_isClean uintptr
	addr_QUndoGroup_undo uintptr
	addr_QUndoGroup_redo uintptr
	addr_QUndoGroup_setActiveStack uintptr
	addr_QUndoGroup_activeStackChanged uintptr
	addr_QUndoGroup_indexChanged uintptr
	addr_QUndoGroup_cleanChanged uintptr
	addr_QUndoGroup_canUndoChanged uintptr
	addr_QUndoGroup_canRedoChanged uintptr
	addr_QUndoGroup_undoTextChanged uintptr
	addr_QUndoGroup_redoTextChanged uintptr
	addr_QUndoGroup_connect uintptr
	addr_QUndoGroup_connect2 uintptr
	addr_QUndoGroup_sender uintptr
	addr_QUndoGroup_senderSignalIndex uintptr
	addr_QUndoGroup_receivers uintptr
	addr_QUndoGroup_isSignalConnected uintptr
	addr_QUndoGroup_tr2 uintptr
	addr_QUndoGroup_tr3 uintptr
	addr_QUndoGroup_createUndoAction2 uintptr
	addr_QUndoGroup_createRedoAction2 uintptr
	addr_QUndoGroup_connect3 uintptr
	addr_QUndoGroup_connect4 uintptr
	addr_QUndoGroup_metaObject_default uintptr
	addr_QUndoGroup_event_default uintptr
	addr_QUndoGroup_eventFilter_default uintptr
	addr_QUndoGroup_timerEvent_default uintptr
	addr_QUndoGroup_childEvent_default uintptr
	addr_QUndoGroup_customEvent_default uintptr
	addr_QUndoGroup_connectNotify_default uintptr
	addr_QUndoGroup_disconnectNotify_default uintptr
	pg_QUndoGroup_connect_activeStackChanged func(uintptr, int64)
	pg_QUndoGroup_connect_indexChanged func(uintptr, int64)
	pg_QUndoGroup_connect_cleanChanged func(uintptr, int64)
	pg_QUndoGroup_connect_canUndoChanged func(uintptr, int64)
	pg_QUndoGroup_connect_canRedoChanged func(uintptr, int64)
	pg_QUndoGroup_connect_undoTextChanged func(uintptr, int64)
	pg_QUndoGroup_connect_redoTextChanged func(uintptr, int64)
	pg_QUndoGroup_connect_destroyed func(uintptr, int64)
	pg_QUndoGroup_connect_destroyed2 func(uintptr, int64)
	pg_QUndoGroup_connect_objectNameChanged func(uintptr, int64)
	pg_QUndoGroup__disconnect func(int64)
	pg_QUndoGroup__set_override func(uintptr, int32)
	pg_QUndoGroup_delete func(uintptr)
)

func init() {
	reg(&pg_QUndoGroup_base, "QUndoGroup_base")
	symaddr(&addr_QUndoGroup_new, "QUndoGroup_new")
	symaddr(&addr_QUndoGroup_new2, "QUndoGroup_new2")
	symaddr(&addr_QUndoGroup_tr, "QUndoGroup_tr")
	symaddr(&addr_QUndoGroup_addStack, "QUndoGroup_addStack")
	symaddr(&addr_QUndoGroup_removeStack, "QUndoGroup_removeStack")
	symaddr(&addr_QUndoGroup_stacks, "QUndoGroup_stacks")
	symaddr(&addr_QUndoGroup_activeStack, "QUndoGroup_activeStack")
	symaddr(&addr_QUndoGroup_createUndoAction, "QUndoGroup_createUndoAction")
	symaddr(&addr_QUndoGroup_createRedoAction, "QUndoGroup_createRedoAction")
	symaddr(&addr_QUndoGroup_canUndo, "QUndoGroup_canUndo")
	symaddr(&addr_QUndoGroup_canRedo, "QUndoGroup_canRedo")
	symaddr(&addr_QUndoGroup_undoText, "QUndoGroup_undoText")
	symaddr(&addr_QUndoGroup_redoText, "QUndoGroup_redoText")
	symaddr(&addr_QUndoGroup_isClean, "QUndoGroup_isClean")
	symaddr(&addr_QUndoGroup_undo, "QUndoGroup_undo")
	symaddr(&addr_QUndoGroup_redo, "QUndoGroup_redo")
	symaddr(&addr_QUndoGroup_setActiveStack, "QUndoGroup_setActiveStack")
	symaddr(&addr_QUndoGroup_activeStackChanged, "QUndoGroup_activeStackChanged")
	symaddr(&addr_QUndoGroup_indexChanged, "QUndoGroup_indexChanged")
	symaddr(&addr_QUndoGroup_cleanChanged, "QUndoGroup_cleanChanged")
	symaddr(&addr_QUndoGroup_canUndoChanged, "QUndoGroup_canUndoChanged")
	symaddr(&addr_QUndoGroup_canRedoChanged, "QUndoGroup_canRedoChanged")
	symaddr(&addr_QUndoGroup_undoTextChanged, "QUndoGroup_undoTextChanged")
	symaddr(&addr_QUndoGroup_redoTextChanged, "QUndoGroup_redoTextChanged")
	symaddr(&addr_QUndoGroup_connect, "QUndoGroup_connect")
	symaddr(&addr_QUndoGroup_connect2, "QUndoGroup_connect2")
	symaddr(&addr_QUndoGroup_sender, "QUndoGroup_sender")
	symaddr(&addr_QUndoGroup_senderSignalIndex, "QUndoGroup_senderSignalIndex")
	symaddr(&addr_QUndoGroup_receivers, "QUndoGroup_receivers")
	symaddr(&addr_QUndoGroup_isSignalConnected, "QUndoGroup_isSignalConnected")
	symaddr(&addr_QUndoGroup_tr2, "QUndoGroup_tr2")
	symaddr(&addr_QUndoGroup_tr3, "QUndoGroup_tr3")
	symaddr(&addr_QUndoGroup_createUndoAction2, "QUndoGroup_createUndoAction2")
	symaddr(&addr_QUndoGroup_createRedoAction2, "QUndoGroup_createRedoAction2")
	symaddr(&addr_QUndoGroup_connect3, "QUndoGroup_connect3")
	symaddr(&addr_QUndoGroup_connect4, "QUndoGroup_connect4")
	symaddr(&addr_QUndoGroup_metaObject_default, "QUndoGroup_metaObject_default")
	symaddr(&addr_QUndoGroup_event_default, "QUndoGroup_event_default")
	symaddr(&addr_QUndoGroup_eventFilter_default, "QUndoGroup_eventFilter_default")
	symaddr(&addr_QUndoGroup_timerEvent_default, "QUndoGroup_timerEvent_default")
	symaddr(&addr_QUndoGroup_childEvent_default, "QUndoGroup_childEvent_default")
	symaddr(&addr_QUndoGroup_customEvent_default, "QUndoGroup_customEvent_default")
	symaddr(&addr_QUndoGroup_connectNotify_default, "QUndoGroup_connectNotify_default")
	symaddr(&addr_QUndoGroup_disconnectNotify_default, "QUndoGroup_disconnectNotify_default")
	reg(&pg_QUndoGroup_connect_activeStackChanged, "QUndoGroup_connect_activeStackChanged")
	reg(&pg_QUndoGroup_connect_indexChanged, "QUndoGroup_connect_indexChanged")
	reg(&pg_QUndoGroup_connect_cleanChanged, "QUndoGroup_connect_cleanChanged")
	reg(&pg_QUndoGroup_connect_canUndoChanged, "QUndoGroup_connect_canUndoChanged")
	reg(&pg_QUndoGroup_connect_canRedoChanged, "QUndoGroup_connect_canRedoChanged")
	reg(&pg_QUndoGroup_connect_undoTextChanged, "QUndoGroup_connect_undoTextChanged")
	reg(&pg_QUndoGroup_connect_redoTextChanged, "QUndoGroup_connect_redoTextChanged")
	reg(&pg_QUndoGroup_connect_destroyed, "QUndoGroup_connect_destroyed")
	reg(&pg_QUndoGroup_connect_destroyed2, "QUndoGroup_connect_destroyed2")
	reg(&pg_QUndoGroup_connect_objectNameChanged, "QUndoGroup_connect_objectNameChanged")
	reg(&pg_QUndoGroup__disconnect, "QUndoGroup__disconnect")
	reg(&pg_QUndoGroup__set_override, "QUndoGroup__set_override")
	reg(&pg_QUndoGroup_delete, "QUndoGroup_delete")
}

type QUndoGroup struct {
	h uintptr
	*QObject
}

func bindQUndoGroup(h uintptr) *QUndoGroup {
	if h == 0 {
		return nil
	}
	return &QUndoGroup{h: h, QObject: bindQObject(pg_QUndoGroup_base(h))}
}

func newownedQUndoGroup(h uintptr) *QUndoGroup { return bindQUndoGroup(h) }

func (this *QUndoGroup) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QUndoGroup) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQUndoGroup(h unsafe.Pointer) *QUndoGroup { return bindQUndoGroup(uintptr(h)) }
func UnsafeNewQUndoGroupOwned(h unsafe.Pointer) *QUndoGroup { return newownedQUndoGroup(uintptr(h)) }

func (this *QUndoGroup) UnsafeRelease() {
	this.h = 0
	if this.QObject != nil {
		this.QObject.UnsafeRelease()
	}
}

func NewQUndoGroup() *QUndoGroup {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QUndoGroup_new, uintptr(__goid))
	__o := newownedQUndoGroup(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QUndoGroup_virtIdx, QUndoGroup_enable)
	return __o
}

func NewQUndoGroup2(parent *QObject) *QUndoGroup {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QUndoGroup_new2, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQUndoGroup(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QUndoGroup_virtIdx, QUndoGroup_enable)
	return __o
}

func QUndoGroup_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QUndoGroup_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QUndoGroup) AddStack(stack *QUndoStack) {
	purego.SyscallN(addr_QUndoGroup_addStack, this.h, uintptr(stack.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(stack)
}

func (this *QUndoGroup) RemoveStack(stack *QUndoStack) {
	purego.SyscallN(addr_QUndoGroup_removeStack, this.h, uintptr(stack.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(stack)
}

func (this *QUndoGroup) Stacks() []*QUndoStack {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QUndoGroup_stacks, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QUndoStack, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = bindQUndoStack(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QUndoGroup) ActiveStack() *QUndoStack {
	__r, _, _ := purego.SyscallN(addr_QUndoGroup_activeStack, this.h)
	runtime.KeepAlive(this)
	return bindQUndoStack(__r)
}

func (this *QUndoGroup) CreateUndoAction(parent *QObject) *QAction {
	__r, _, _ := purego.SyscallN(addr_QUndoGroup_createUndoAction, this.h, uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return bindQAction(__r)
}

func (this *QUndoGroup) CreateRedoAction(parent *QObject) *QAction {
	__r, _, _ := purego.SyscallN(addr_QUndoGroup_createRedoAction, this.h, uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return bindQAction(__r)
}

func (this *QUndoGroup) CanUndo() bool {
	__r0, _, _ := purego.SyscallN(addr_QUndoGroup_canUndo, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QUndoGroup) CanRedo() bool {
	__r0, _, _ := purego.SyscallN(addr_QUndoGroup_canRedo, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QUndoGroup) UndoText() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QUndoGroup_undoText, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QUndoGroup) RedoText() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QUndoGroup_redoText, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QUndoGroup) IsClean() bool {
	__r0, _, _ := purego.SyscallN(addr_QUndoGroup_isClean, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QUndoGroup) Undo() {
	purego.SyscallN(addr_QUndoGroup_undo, this.h)
	runtime.KeepAlive(this)
}

func (this *QUndoGroup) Redo() {
	purego.SyscallN(addr_QUndoGroup_redo, this.h)
	runtime.KeepAlive(this)
}

func (this *QUndoGroup) SetActiveStack(stack *QUndoStack) {
	purego.SyscallN(addr_QUndoGroup_setActiveStack, this.h, uintptr(stack.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(stack)
}

func (this *QUndoGroup) ActiveStackChanged(stack *QUndoStack) {
	purego.SyscallN(addr_QUndoGroup_activeStackChanged, this.h, uintptr(stack.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(stack)
}

func (this *QUndoGroup) IndexChanged(idx int) {
	purego.SyscallN(addr_QUndoGroup_indexChanged, this.h, uintptr(idx))
	runtime.KeepAlive(this)
}

func (this *QUndoGroup) CleanChanged(clean bool) {
	purego.SyscallN(addr_QUndoGroup_cleanChanged, this.h, uintptr(b2i(clean)))
	runtime.KeepAlive(this)
}

func (this *QUndoGroup) CanUndoChanged(canUndo bool) {
	purego.SyscallN(addr_QUndoGroup_canUndoChanged, this.h, uintptr(b2i(canUndo)))
	runtime.KeepAlive(this)
}

func (this *QUndoGroup) CanRedoChanged(canRedo bool) {
	purego.SyscallN(addr_QUndoGroup_canRedoChanged, this.h, uintptr(b2i(canRedo)))
	runtime.KeepAlive(this)
}

func (this *QUndoGroup) UndoTextChanged(undoText string) {
	var undoText_p uintptr
	if len(undoText) > 0 {
		undoText_p = uintptr(unsafe.Pointer(unsafe.StringData(undoText)))
	}
	purego.SyscallN(addr_QUndoGroup_undoTextChanged, this.h, undoText_p, uintptr(len(undoText)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(undoText)
}

func (this *QUndoGroup) RedoTextChanged(redoText string) {
	var redoText_p uintptr
	if len(redoText) > 0 {
		redoText_p = uintptr(unsafe.Pointer(unsafe.StringData(redoText)))
	}
	purego.SyscallN(addr_QUndoGroup_redoTextChanged, this.h, redoText_p, uintptr(len(redoText)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(redoText)
}

func QUndoGroup_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QUndoGroup_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QUndoGroup_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QUndoGroup_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QUndoGroup) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QUndoGroup_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QUndoGroup) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QUndoGroup_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QUndoGroup) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QUndoGroup_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QUndoGroup) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QUndoGroup_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QUndoGroup_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QUndoGroup_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QUndoGroup_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QUndoGroup_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func (this *QUndoGroup) CreateUndoAction2(parent *QObject, prefix string) *QAction {
	var prefix_p uintptr
	if len(prefix) > 0 {
		prefix_p = uintptr(unsafe.Pointer(unsafe.StringData(prefix)))
	}
	__r, _, _ := purego.SyscallN(addr_QUndoGroup_createUndoAction2, this.h, uintptr(parent.UnsafePointer()), prefix_p, uintptr(len(prefix)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	runtime.KeepAlive(prefix)
	return bindQAction(__r)
}

func (this *QUndoGroup) CreateRedoAction2(parent *QObject, prefix string) *QAction {
	var prefix_p uintptr
	if len(prefix) > 0 {
		prefix_p = uintptr(unsafe.Pointer(unsafe.StringData(prefix)))
	}
	__r, _, _ := purego.SyscallN(addr_QUndoGroup_createRedoAction2, this.h, uintptr(parent.UnsafePointer()), prefix_p, uintptr(len(prefix)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	runtime.KeepAlive(prefix)
	return bindQAction(__r)
}

func QUndoGroup_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QUndoGroup_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QUndoGroup_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QUndoGroup_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QUndoGroup) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QUndoGroup_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QUndoGroup) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QUndoGroup_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QUndoGroup) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QUndoGroup_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QUndoGroup) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QUndoGroup_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QUndoGroup) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QUndoGroup_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QUndoGroup) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QUndoGroup_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QUndoGroup) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QUndoGroup_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QUndoGroup) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QUndoGroup_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QUndoGroup) OnActiveStackChanged(fn func(*QUndoStack)) int64 {
	id := registerCallback(func(stack uintptr) { fn(UnsafeNewQUndoStack(unsafe.Pointer(stack))) })
	pg_QUndoGroup_connect_activeStackChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QUndoGroup) OnIndexChanged(fn func(int)) int64 {
	id := registerCallback(func(idx int32) { fn(int(idx)) })
	pg_QUndoGroup_connect_indexChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QUndoGroup) OnCleanChanged(fn func(bool)) int64 {
	id := registerCallback(func(clean int32) { fn(clean != 0) })
	pg_QUndoGroup_connect_cleanChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QUndoGroup) OnCanUndoChanged(fn func(bool)) int64 {
	id := registerCallback(func(canUndo int32) { fn(canUndo != 0) })
	pg_QUndoGroup_connect_canUndoChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QUndoGroup) OnCanRedoChanged(fn func(bool)) int64 {
	id := registerCallback(func(canRedo int32) { fn(canRedo != 0) })
	pg_QUndoGroup_connect_canRedoChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QUndoGroup) OnUndoTextChanged(fn func(string)) int64 {
	id := registerCallback(func(undoText uintptr, undoText_len int64) { fn(pgStrNoFree(undoText, undoText_len)) })
	pg_QUndoGroup_connect_undoTextChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QUndoGroup) OnRedoTextChanged(fn func(string)) int64 {
	id := registerCallback(func(redoText uintptr, redoText_len int64) { fn(pgStrNoFree(redoText, redoText_len)) })
	pg_QUndoGroup_connect_redoTextChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QUndoGroup) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QUndoGroup_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QUndoGroup) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QUndoGroup_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QUndoGroup) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QUndoGroup_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QUndoGroup) Disconnect(handle int64) {
	pg_QUndoGroup__disconnect(handle)
	unregisterCallback(handle)
}

func QUndoGroup_virtIdx(slot int64) int32 {
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

func QUndoGroup_enable(h uintptr, idx int32) { pg_QUndoGroup__set_override(h, idx) }

func (this *QUndoGroup) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QUndoGroup) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QUndoGroup) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QUndoGroup) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QUndoGroup) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QUndoGroup) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QUndoGroup) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QUndoGroup) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QUndoGroup) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QUndoGroup_delete(this.h)
	}
	this.UnsafeRelease()
}
