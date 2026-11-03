package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QSyntaxHighlighter_base func(uintptr) uintptr
	addr_QSyntaxHighlighter_tr uintptr
	addr_QSyntaxHighlighter_setDocument uintptr
	addr_QSyntaxHighlighter_document uintptr
	addr_QSyntaxHighlighter_rehighlight uintptr
	addr_QSyntaxHighlighter_rehighlightBlock uintptr
	addr_QSyntaxHighlighter_setFormat uintptr
	addr_QSyntaxHighlighter_setFormat2 uintptr
	addr_QSyntaxHighlighter_setFormat3 uintptr
	addr_QSyntaxHighlighter_format uintptr
	addr_QSyntaxHighlighter_previousBlockState uintptr
	addr_QSyntaxHighlighter_currentBlockState uintptr
	addr_QSyntaxHighlighter_setCurrentBlockState uintptr
	addr_QSyntaxHighlighter_setCurrentBlockUserData uintptr
	addr_QSyntaxHighlighter_currentBlockUserData uintptr
	addr_QSyntaxHighlighter_currentBlock uintptr
	addr_QSyntaxHighlighter_connect uintptr
	addr_QSyntaxHighlighter_connect2 uintptr
	addr_QSyntaxHighlighter_sender uintptr
	addr_QSyntaxHighlighter_senderSignalIndex uintptr
	addr_QSyntaxHighlighter_receivers uintptr
	addr_QSyntaxHighlighter_isSignalConnected uintptr
	addr_QSyntaxHighlighter_tr2 uintptr
	addr_QSyntaxHighlighter_tr3 uintptr
	addr_QSyntaxHighlighter_connect3 uintptr
	addr_QSyntaxHighlighter_connect4 uintptr
	addr_QSyntaxHighlighter_metaObject_default uintptr
	addr_QSyntaxHighlighter_event_default uintptr
	addr_QSyntaxHighlighter_eventFilter_default uintptr
	addr_QSyntaxHighlighter_timerEvent_default uintptr
	addr_QSyntaxHighlighter_childEvent_default uintptr
	addr_QSyntaxHighlighter_customEvent_default uintptr
	addr_QSyntaxHighlighter_connectNotify_default uintptr
	addr_QSyntaxHighlighter_disconnectNotify_default uintptr
	pg_QSyntaxHighlighter_connect_destroyed func(uintptr, int64)
	pg_QSyntaxHighlighter_connect_destroyed2 func(uintptr, int64)
	pg_QSyntaxHighlighter_connect_objectNameChanged func(uintptr, int64)
	pg_QSyntaxHighlighter__disconnect func(int64)
	pg_QSyntaxHighlighter__set_override func(uintptr, int32)
	pg_QSyntaxHighlighter__Go_new func(int64, uintptr) uintptr
	pg_QSyntaxHighlighter_delete func(uintptr)
)

func init() {
	reg(&pg_QSyntaxHighlighter_base, "QSyntaxHighlighter_base")
	symaddr(&addr_QSyntaxHighlighter_tr, "QSyntaxHighlighter_tr")
	symaddr(&addr_QSyntaxHighlighter_setDocument, "QSyntaxHighlighter_setDocument")
	symaddr(&addr_QSyntaxHighlighter_document, "QSyntaxHighlighter_document")
	symaddr(&addr_QSyntaxHighlighter_rehighlight, "QSyntaxHighlighter_rehighlight")
	symaddr(&addr_QSyntaxHighlighter_rehighlightBlock, "QSyntaxHighlighter_rehighlightBlock")
	symaddr(&addr_QSyntaxHighlighter_setFormat, "QSyntaxHighlighter_setFormat")
	symaddr(&addr_QSyntaxHighlighter_setFormat2, "QSyntaxHighlighter_setFormat2")
	symaddr(&addr_QSyntaxHighlighter_setFormat3, "QSyntaxHighlighter_setFormat3")
	symaddr(&addr_QSyntaxHighlighter_format, "QSyntaxHighlighter_format")
	symaddr(&addr_QSyntaxHighlighter_previousBlockState, "QSyntaxHighlighter_previousBlockState")
	symaddr(&addr_QSyntaxHighlighter_currentBlockState, "QSyntaxHighlighter_currentBlockState")
	symaddr(&addr_QSyntaxHighlighter_setCurrentBlockState, "QSyntaxHighlighter_setCurrentBlockState")
	symaddr(&addr_QSyntaxHighlighter_setCurrentBlockUserData, "QSyntaxHighlighter_setCurrentBlockUserData")
	symaddr(&addr_QSyntaxHighlighter_currentBlockUserData, "QSyntaxHighlighter_currentBlockUserData")
	symaddr(&addr_QSyntaxHighlighter_currentBlock, "QSyntaxHighlighter_currentBlock")
	symaddr(&addr_QSyntaxHighlighter_connect, "QSyntaxHighlighter_connect")
	symaddr(&addr_QSyntaxHighlighter_connect2, "QSyntaxHighlighter_connect2")
	symaddr(&addr_QSyntaxHighlighter_sender, "QSyntaxHighlighter_sender")
	symaddr(&addr_QSyntaxHighlighter_senderSignalIndex, "QSyntaxHighlighter_senderSignalIndex")
	symaddr(&addr_QSyntaxHighlighter_receivers, "QSyntaxHighlighter_receivers")
	symaddr(&addr_QSyntaxHighlighter_isSignalConnected, "QSyntaxHighlighter_isSignalConnected")
	symaddr(&addr_QSyntaxHighlighter_tr2, "QSyntaxHighlighter_tr2")
	symaddr(&addr_QSyntaxHighlighter_tr3, "QSyntaxHighlighter_tr3")
	symaddr(&addr_QSyntaxHighlighter_connect3, "QSyntaxHighlighter_connect3")
	symaddr(&addr_QSyntaxHighlighter_connect4, "QSyntaxHighlighter_connect4")
	symaddr(&addr_QSyntaxHighlighter_metaObject_default, "QSyntaxHighlighter_metaObject_default")
	symaddr(&addr_QSyntaxHighlighter_event_default, "QSyntaxHighlighter_event_default")
	symaddr(&addr_QSyntaxHighlighter_eventFilter_default, "QSyntaxHighlighter_eventFilter_default")
	symaddr(&addr_QSyntaxHighlighter_timerEvent_default, "QSyntaxHighlighter_timerEvent_default")
	symaddr(&addr_QSyntaxHighlighter_childEvent_default, "QSyntaxHighlighter_childEvent_default")
	symaddr(&addr_QSyntaxHighlighter_customEvent_default, "QSyntaxHighlighter_customEvent_default")
	symaddr(&addr_QSyntaxHighlighter_connectNotify_default, "QSyntaxHighlighter_connectNotify_default")
	symaddr(&addr_QSyntaxHighlighter_disconnectNotify_default, "QSyntaxHighlighter_disconnectNotify_default")
	reg(&pg_QSyntaxHighlighter_connect_destroyed, "QSyntaxHighlighter_connect_destroyed")
	reg(&pg_QSyntaxHighlighter_connect_destroyed2, "QSyntaxHighlighter_connect_destroyed2")
	reg(&pg_QSyntaxHighlighter_connect_objectNameChanged, "QSyntaxHighlighter_connect_objectNameChanged")
	reg(&pg_QSyntaxHighlighter__disconnect, "QSyntaxHighlighter__disconnect")
	reg(&pg_QSyntaxHighlighter__set_override, "QSyntaxHighlighter__set_override")
	reg(&pg_QSyntaxHighlighter__Go_new, "QSyntaxHighlighter__Go_new")
	reg(&pg_QSyntaxHighlighter_delete, "QSyntaxHighlighter_delete")
}

type QSyntaxHighlighter struct {
	h uintptr
	*QObject
}

func bindQSyntaxHighlighter(h uintptr) *QSyntaxHighlighter {
	if h == 0 {
		return nil
	}
	return &QSyntaxHighlighter{h: h, QObject: bindQObject(pg_QSyntaxHighlighter_base(h))}
}

func newownedQSyntaxHighlighter(h uintptr) *QSyntaxHighlighter { return bindQSyntaxHighlighter(h) }

func (this *QSyntaxHighlighter) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QSyntaxHighlighter) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQSyntaxHighlighter(h unsafe.Pointer) *QSyntaxHighlighter { return bindQSyntaxHighlighter(uintptr(h)) }
func UnsafeNewQSyntaxHighlighterOwned(h unsafe.Pointer) *QSyntaxHighlighter { return newownedQSyntaxHighlighter(uintptr(h)) }

func (this *QSyntaxHighlighter) UnsafeRelease() {
	this.h = 0
	if this.QObject != nil {
		this.QObject.UnsafeRelease()
	}
}

func QSyntaxHighlighter_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QSyntaxHighlighter_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QSyntaxHighlighter) SetDocument(doc *QTextDocument) {
	purego.SyscallN(addr_QSyntaxHighlighter_setDocument, this.h, uintptr(doc.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(doc)
}

func (this *QSyntaxHighlighter) Document() *QTextDocument {
	__r, _, _ := purego.SyscallN(addr_QSyntaxHighlighter_document, this.h)
	runtime.KeepAlive(this)
	return bindQTextDocument(__r)
}

func (this *QSyntaxHighlighter) Rehighlight() {
	purego.SyscallN(addr_QSyntaxHighlighter_rehighlight, this.h)
	runtime.KeepAlive(this)
}

func (this *QSyntaxHighlighter) RehighlightBlock(block *QTextBlock) {
	purego.SyscallN(addr_QSyntaxHighlighter_rehighlightBlock, this.h, uintptr(block.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(block)
}

func (this *QSyntaxHighlighter) SetFormat(start int, count int, format *QTextCharFormat) {
	purego.SyscallN(addr_QSyntaxHighlighter_setFormat, this.h, uintptr(start), uintptr(count), uintptr(format.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(format)
}

func (this *QSyntaxHighlighter) SetFormat2(start int, count int, color *QColor) {
	purego.SyscallN(addr_QSyntaxHighlighter_setFormat2, this.h, uintptr(start), uintptr(count), uintptr(color.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(color)
}

func (this *QSyntaxHighlighter) SetFormat3(start int, count int, font *QFont) {
	purego.SyscallN(addr_QSyntaxHighlighter_setFormat3, this.h, uintptr(start), uintptr(count), uintptr(font.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(font)
}

func (this *QSyntaxHighlighter) Format(pos int) *QTextCharFormat {
	__r, _, _ := purego.SyscallN(addr_QSyntaxHighlighter_format, this.h, uintptr(pos))
	runtime.KeepAlive(this)
	return newownedQTextCharFormat(__r)
}

func (this *QSyntaxHighlighter) PreviousBlockState() int {
	__r0, _, _ := purego.SyscallN(addr_QSyntaxHighlighter_previousBlockState, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QSyntaxHighlighter) CurrentBlockState() int {
	__r0, _, _ := purego.SyscallN(addr_QSyntaxHighlighter_currentBlockState, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QSyntaxHighlighter) SetCurrentBlockState(newState int) {
	purego.SyscallN(addr_QSyntaxHighlighter_setCurrentBlockState, this.h, uintptr(newState))
	runtime.KeepAlive(this)
}

func (this *QSyntaxHighlighter) SetCurrentBlockUserData(data *QTextBlockUserData) {
	purego.SyscallN(addr_QSyntaxHighlighter_setCurrentBlockUserData, this.h, uintptr(data.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(data)
}

func (this *QSyntaxHighlighter) CurrentBlockUserData() *QTextBlockUserData {
	__r, _, _ := purego.SyscallN(addr_QSyntaxHighlighter_currentBlockUserData, this.h)
	runtime.KeepAlive(this)
	return bindQTextBlockUserData(__r)
}

func (this *QSyntaxHighlighter) CurrentBlock() *QTextBlock {
	__r, _, _ := purego.SyscallN(addr_QSyntaxHighlighter_currentBlock, this.h)
	runtime.KeepAlive(this)
	return newownedQTextBlock(__r)
}

func QSyntaxHighlighter_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QSyntaxHighlighter_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QSyntaxHighlighter_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QSyntaxHighlighter_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QSyntaxHighlighter) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QSyntaxHighlighter_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QSyntaxHighlighter) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QSyntaxHighlighter_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QSyntaxHighlighter) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QSyntaxHighlighter_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QSyntaxHighlighter) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QSyntaxHighlighter_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QSyntaxHighlighter_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QSyntaxHighlighter_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QSyntaxHighlighter_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QSyntaxHighlighter_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QSyntaxHighlighter_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QSyntaxHighlighter_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QSyntaxHighlighter_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QSyntaxHighlighter_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QSyntaxHighlighter) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QSyntaxHighlighter_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QSyntaxHighlighter) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QSyntaxHighlighter_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QSyntaxHighlighter) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QSyntaxHighlighter_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QSyntaxHighlighter) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QSyntaxHighlighter_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSyntaxHighlighter) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QSyntaxHighlighter_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSyntaxHighlighter) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QSyntaxHighlighter_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSyntaxHighlighter) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QSyntaxHighlighter_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QSyntaxHighlighter) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QSyntaxHighlighter_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QSyntaxHighlighter) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QSyntaxHighlighter_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSyntaxHighlighter) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QSyntaxHighlighter_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSyntaxHighlighter) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QSyntaxHighlighter_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSyntaxHighlighter) Disconnect(handle int64) {
	pg_QSyntaxHighlighter__disconnect(handle)
	unregisterCallback(handle)
}

func QSyntaxHighlighter_virtIdx(slot int64) int32 {
	switch slot {
	case -2478846640236730472:
		return 0
	case 1607084412591819503:
		return 1
	case 1237236456312824991:
		return 2
	case 5217623514630858297:
		return 3
	case -5409927581812374498:
		return 4
	case -1185315367526857833:
		return 5
	case 3769868229123868474:
		return 6
	case 7379029982170503008:
		return 7
	case -568369772982215962:
		return 8
	}
	return -1
}

func QSyntaxHighlighter_enable(h uintptr, idx int32) { pg_QSyntaxHighlighter__set_override(h, idx) }

func NewQSyntaxHighlighter(parent *QObject) *QSyntaxHighlighter {
	id := nextVirtID()
	o := newownedQSyntaxHighlighter(pg_QSyntaxHighlighter__Go_new(id, uintptr(parent.UnsafePointer())))
	o.UnsafeSetGoID(id)
	registerInstance(id, o.h, QSyntaxHighlighter_virtIdx, QSyntaxHighlighter_enable)
	return o
}

func (this *QSyntaxHighlighter) OnHighlightBlock(fn func(string)) {
	registerVirt(this.UnsafeGoID(), -2478846640236730472, func(text uintptr, text_len int64) { fn(pgStrNoFree(text, text_len)) }) // slot "highlightBlock"
}

func (this *QSyntaxHighlighter) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QSyntaxHighlighter) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QSyntaxHighlighter) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QSyntaxHighlighter) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QSyntaxHighlighter) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QSyntaxHighlighter) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QSyntaxHighlighter) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QSyntaxHighlighter) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QSyntaxHighlighter) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QSyntaxHighlighter_delete(this.h)
	}
	this.UnsafeRelease()
}
