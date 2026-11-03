package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QDrag_base func(uintptr) uintptr
	addr_QDrag_new uintptr
	addr_QDrag_tr uintptr
	addr_QDrag_setMimeData uintptr
	addr_QDrag_mimeData uintptr
	addr_QDrag_setPixmap uintptr
	addr_QDrag_pixmap uintptr
	addr_QDrag_setHotSpot uintptr
	addr_QDrag_hotSpot uintptr
	addr_QDrag_source uintptr
	addr_QDrag_target uintptr
	addr_QDrag_exec uintptr
	addr_QDrag_exec2 uintptr
	addr_QDrag_setDragCursor uintptr
	addr_QDrag_dragCursor uintptr
	addr_QDrag_supportedActions uintptr
	addr_QDrag_defaultAction uintptr
	addr_QDrag_cancel uintptr
	addr_QDrag_actionChanged uintptr
	addr_QDrag_targetChanged uintptr
	addr_QDrag_connect uintptr
	addr_QDrag_connect2 uintptr
	addr_QDrag_sender uintptr
	addr_QDrag_senderSignalIndex uintptr
	addr_QDrag_receivers uintptr
	addr_QDrag_isSignalConnected uintptr
	addr_QDrag_tr2 uintptr
	addr_QDrag_tr3 uintptr
	addr_QDrag_exec3 uintptr
	addr_QDrag_connect3 uintptr
	addr_QDrag_connect4 uintptr
	addr_QDrag_metaObject_default uintptr
	addr_QDrag_event_default uintptr
	addr_QDrag_eventFilter_default uintptr
	addr_QDrag_timerEvent_default uintptr
	addr_QDrag_childEvent_default uintptr
	addr_QDrag_customEvent_default uintptr
	addr_QDrag_connectNotify_default uintptr
	addr_QDrag_disconnectNotify_default uintptr
	pg_QDrag_connect_actionChanged func(uintptr, int64)
	pg_QDrag_connect_targetChanged func(uintptr, int64)
	pg_QDrag_connect_destroyed func(uintptr, int64)
	pg_QDrag_connect_destroyed2 func(uintptr, int64)
	pg_QDrag_connect_objectNameChanged func(uintptr, int64)
	pg_QDrag__disconnect func(int64)
	pg_QDrag__set_override func(uintptr, int32)
	pg_QDrag_delete func(uintptr)
)

func init() {
	reg(&pg_QDrag_base, "QDrag_base")
	symaddr(&addr_QDrag_new, "QDrag_new")
	symaddr(&addr_QDrag_tr, "QDrag_tr")
	symaddr(&addr_QDrag_setMimeData, "QDrag_setMimeData")
	symaddr(&addr_QDrag_mimeData, "QDrag_mimeData")
	symaddr(&addr_QDrag_setPixmap, "QDrag_setPixmap")
	symaddr(&addr_QDrag_pixmap, "QDrag_pixmap")
	symaddr(&addr_QDrag_setHotSpot, "QDrag_setHotSpot")
	symaddr(&addr_QDrag_hotSpot, "QDrag_hotSpot")
	symaddr(&addr_QDrag_source, "QDrag_source")
	symaddr(&addr_QDrag_target, "QDrag_target")
	symaddr(&addr_QDrag_exec, "QDrag_exec")
	symaddr(&addr_QDrag_exec2, "QDrag_exec2")
	symaddr(&addr_QDrag_setDragCursor, "QDrag_setDragCursor")
	symaddr(&addr_QDrag_dragCursor, "QDrag_dragCursor")
	symaddr(&addr_QDrag_supportedActions, "QDrag_supportedActions")
	symaddr(&addr_QDrag_defaultAction, "QDrag_defaultAction")
	symaddr(&addr_QDrag_cancel, "QDrag_cancel")
	symaddr(&addr_QDrag_actionChanged, "QDrag_actionChanged")
	symaddr(&addr_QDrag_targetChanged, "QDrag_targetChanged")
	symaddr(&addr_QDrag_connect, "QDrag_connect")
	symaddr(&addr_QDrag_connect2, "QDrag_connect2")
	symaddr(&addr_QDrag_sender, "QDrag_sender")
	symaddr(&addr_QDrag_senderSignalIndex, "QDrag_senderSignalIndex")
	symaddr(&addr_QDrag_receivers, "QDrag_receivers")
	symaddr(&addr_QDrag_isSignalConnected, "QDrag_isSignalConnected")
	symaddr(&addr_QDrag_tr2, "QDrag_tr2")
	symaddr(&addr_QDrag_tr3, "QDrag_tr3")
	symaddr(&addr_QDrag_exec3, "QDrag_exec3")
	symaddr(&addr_QDrag_connect3, "QDrag_connect3")
	symaddr(&addr_QDrag_connect4, "QDrag_connect4")
	symaddr(&addr_QDrag_metaObject_default, "QDrag_metaObject_default")
	symaddr(&addr_QDrag_event_default, "QDrag_event_default")
	symaddr(&addr_QDrag_eventFilter_default, "QDrag_eventFilter_default")
	symaddr(&addr_QDrag_timerEvent_default, "QDrag_timerEvent_default")
	symaddr(&addr_QDrag_childEvent_default, "QDrag_childEvent_default")
	symaddr(&addr_QDrag_customEvent_default, "QDrag_customEvent_default")
	symaddr(&addr_QDrag_connectNotify_default, "QDrag_connectNotify_default")
	symaddr(&addr_QDrag_disconnectNotify_default, "QDrag_disconnectNotify_default")
	reg(&pg_QDrag_connect_actionChanged, "QDrag_connect_actionChanged")
	reg(&pg_QDrag_connect_targetChanged, "QDrag_connect_targetChanged")
	reg(&pg_QDrag_connect_destroyed, "QDrag_connect_destroyed")
	reg(&pg_QDrag_connect_destroyed2, "QDrag_connect_destroyed2")
	reg(&pg_QDrag_connect_objectNameChanged, "QDrag_connect_objectNameChanged")
	reg(&pg_QDrag__disconnect, "QDrag__disconnect")
	reg(&pg_QDrag__set_override, "QDrag__set_override")
	reg(&pg_QDrag_delete, "QDrag_delete")
}

type QDrag struct {
	h uintptr
	*QObject
}

func bindQDrag(h uintptr) *QDrag {
	if h == 0 {
		return nil
	}
	return &QDrag{h: h, QObject: bindQObject(pg_QDrag_base(h))}
}

func newownedQDrag(h uintptr) *QDrag { return bindQDrag(h) }

func (this *QDrag) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QDrag) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQDrag(h unsafe.Pointer) *QDrag { return bindQDrag(uintptr(h)) }
func UnsafeNewQDragOwned(h unsafe.Pointer) *QDrag { return newownedQDrag(uintptr(h)) }

func (this *QDrag) UnsafeRelease() {
	this.h = 0
	if this.QObject != nil {
		this.QObject.UnsafeRelease()
	}
}

func NewQDrag(dragSource *QObject) *QDrag {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QDrag_new, uintptr(__goid), uintptr(dragSource.UnsafePointer()))
	__o := newownedQDrag(__r)
	runtime.KeepAlive(dragSource)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QDrag_virtIdx, QDrag_enable)
	return __o
}

func QDrag_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDrag_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QDrag) SetMimeData(data *QMimeData) {
	purego.SyscallN(addr_QDrag_setMimeData, this.h, uintptr(data.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(data)
}

func (this *QDrag) MimeData() *QMimeData {
	__r, _, _ := purego.SyscallN(addr_QDrag_mimeData, this.h)
	runtime.KeepAlive(this)
	return bindQMimeData(__r)
}

func (this *QDrag) SetPixmap(arg1 *QPixmap) {
	purego.SyscallN(addr_QDrag_setPixmap, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QDrag) Pixmap() *QPixmap {
	__r, _, _ := purego.SyscallN(addr_QDrag_pixmap, this.h)
	runtime.KeepAlive(this)
	return newownedQPixmap(__r)
}

func (this *QDrag) SetHotSpot(hotspot *QPoint) {
	purego.SyscallN(addr_QDrag_setHotSpot, this.h, uintptr(hotspot.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(hotspot)
}

func (this *QDrag) HotSpot() *QPoint {
	__r, _, _ := purego.SyscallN(addr_QDrag_hotSpot, this.h)
	runtime.KeepAlive(this)
	return newownedQPoint(__r)
}

func (this *QDrag) Source() *QObject {
	__r, _, _ := purego.SyscallN(addr_QDrag_source, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QDrag) Target() *QObject {
	__r, _, _ := purego.SyscallN(addr_QDrag_target, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QDrag) Exec() DropAction {
	__r0, _, _ := purego.SyscallN(addr_QDrag_exec, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return DropAction(__r)
}

func (this *QDrag) Exec2(supportedActions DropAction, defaultAction DropAction) DropAction {
	__r0, _, _ := purego.SyscallN(addr_QDrag_exec2, this.h, uintptr(supportedActions), uintptr(defaultAction))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return DropAction(__r)
}

func (this *QDrag) SetDragCursor(cursor *QPixmap, action DropAction) {
	purego.SyscallN(addr_QDrag_setDragCursor, this.h, uintptr(cursor.UnsafePointer()), uintptr(action))
	runtime.KeepAlive(this)
	runtime.KeepAlive(cursor)
}

func (this *QDrag) DragCursor(action DropAction) *QPixmap {
	__r, _, _ := purego.SyscallN(addr_QDrag_dragCursor, this.h, uintptr(action))
	runtime.KeepAlive(this)
	return newownedQPixmap(__r)
}

func (this *QDrag) SupportedActions() DropAction {
	__r0, _, _ := purego.SyscallN(addr_QDrag_supportedActions, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return DropAction(__r)
}

func (this *QDrag) DefaultAction() DropAction {
	__r0, _, _ := purego.SyscallN(addr_QDrag_defaultAction, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return DropAction(__r)
}

func QDrag_Cancel() {
	purego.SyscallN(addr_QDrag_cancel)
}

func (this *QDrag) ActionChanged(action DropAction) {
	purego.SyscallN(addr_QDrag_actionChanged, this.h, uintptr(action))
	runtime.KeepAlive(this)
}

func (this *QDrag) TargetChanged(newTarget *QObject) {
	purego.SyscallN(addr_QDrag_targetChanged, this.h, uintptr(newTarget.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(newTarget)
}

func QDrag_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QDrag_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QDrag_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QDrag_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QDrag) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QDrag_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QDrag) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QDrag_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QDrag) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QDrag_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QDrag) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QDrag_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QDrag_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDrag_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QDrag_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDrag_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func (this *QDrag) ExecWithSupportedActions(supportedActions DropAction) DropAction {
	__r0, _, _ := purego.SyscallN(addr_QDrag_exec3, this.h, uintptr(supportedActions))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return DropAction(__r)
}

func QDrag_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QDrag_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QDrag_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QDrag_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QDrag) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QDrag_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QDrag) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QDrag_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QDrag) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QDrag_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QDrag) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QDrag_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDrag) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QDrag_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDrag) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QDrag_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDrag) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QDrag_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QDrag) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QDrag_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QDrag) OnActionChanged(fn func(DropAction)) int64 {
	id := registerCallback(func(action int32) { fn(DropAction(action)) })
	pg_QDrag_connect_actionChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDrag) OnTargetChanged(fn func(*QObject)) int64 {
	id := registerCallback(func(newTarget uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(newTarget))) })
	pg_QDrag_connect_targetChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDrag) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QDrag_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDrag) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QDrag_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDrag) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QDrag_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDrag) Disconnect(handle int64) {
	pg_QDrag__disconnect(handle)
	unregisterCallback(handle)
}

func QDrag_virtIdx(slot int64) int32 {
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

func QDrag_enable(h uintptr, idx int32) { pg_QDrag__set_override(h, idx) }

func (this *QDrag) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QDrag) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QDrag) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QDrag) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QDrag) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QDrag) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QDrag) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QDrag) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QDrag) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QDrag_delete(this.h)
	}
	this.UnsafeRelease()
}
