package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QWidgetAction_chain uintptr
	addr_QWidgetAction_new uintptr
	addr_QWidgetAction_tr uintptr
	addr_QWidgetAction_setDefaultWidget uintptr
	addr_QWidgetAction_defaultWidget uintptr
	addr_QWidgetAction_requestWidget uintptr
	addr_QWidgetAction_releaseWidget uintptr
	addr_QWidgetAction_createdWidgets uintptr
	addr_QWidgetAction_connect uintptr
	addr_QWidgetAction_connect2 uintptr
	addr_QWidgetAction_sender uintptr
	addr_QWidgetAction_senderSignalIndex uintptr
	addr_QWidgetAction_receivers uintptr
	addr_QWidgetAction_isSignalConnected uintptr
	addr_QWidgetAction_tr2 uintptr
	addr_QWidgetAction_tr3 uintptr
	addr_QWidgetAction_connect3 uintptr
	addr_QWidgetAction_connect4 uintptr
	addr_QWidgetAction_metaObject_default uintptr
	addr_QWidgetAction_event_default uintptr
	addr_QWidgetAction_eventFilter_default uintptr
	addr_QWidgetAction_createWidget_default uintptr
	addr_QWidgetAction_deleteWidget_default uintptr
	addr_QWidgetAction_timerEvent_default uintptr
	addr_QWidgetAction_childEvent_default uintptr
	addr_QWidgetAction_customEvent_default uintptr
	addr_QWidgetAction_connectNotify_default uintptr
	addr_QWidgetAction_disconnectNotify_default uintptr
	pg_QWidgetAction_connect_changed func(uintptr, int64)
	pg_QWidgetAction_connect_enabledChanged func(uintptr, int64)
	pg_QWidgetAction_connect_checkableChanged func(uintptr, int64)
	pg_QWidgetAction_connect_visibleChanged func(uintptr, int64)
	pg_QWidgetAction_connect_triggered func(uintptr, int64)
	pg_QWidgetAction_connect_triggered2 func(uintptr, int64)
	pg_QWidgetAction_connect_hovered func(uintptr, int64)
	pg_QWidgetAction_connect_toggled func(uintptr, int64)
	pg_QWidgetAction_connect_destroyed func(uintptr, int64)
	pg_QWidgetAction_connect_destroyed2 func(uintptr, int64)
	pg_QWidgetAction_connect_objectNameChanged func(uintptr, int64)
	pg_QWidgetAction__disconnect func(int64)
	pg_QWidgetAction__set_override func(uintptr, int32)
	pg_QWidgetAction_delete func(uintptr)
)

func init() {
	symaddr(&addr_QWidgetAction_chain, "QWidgetAction_chain")
	symaddr(&addr_QWidgetAction_new, "QWidgetAction_new")
	symaddr(&addr_QWidgetAction_tr, "QWidgetAction_tr")
	symaddr(&addr_QWidgetAction_setDefaultWidget, "QWidgetAction_setDefaultWidget")
	symaddr(&addr_QWidgetAction_defaultWidget, "QWidgetAction_defaultWidget")
	symaddr(&addr_QWidgetAction_requestWidget, "QWidgetAction_requestWidget")
	symaddr(&addr_QWidgetAction_releaseWidget, "QWidgetAction_releaseWidget")
	symaddr(&addr_QWidgetAction_createdWidgets, "QWidgetAction_createdWidgets")
	symaddr(&addr_QWidgetAction_connect, "QWidgetAction_connect")
	symaddr(&addr_QWidgetAction_connect2, "QWidgetAction_connect2")
	symaddr(&addr_QWidgetAction_sender, "QWidgetAction_sender")
	symaddr(&addr_QWidgetAction_senderSignalIndex, "QWidgetAction_senderSignalIndex")
	symaddr(&addr_QWidgetAction_receivers, "QWidgetAction_receivers")
	symaddr(&addr_QWidgetAction_isSignalConnected, "QWidgetAction_isSignalConnected")
	symaddr(&addr_QWidgetAction_tr2, "QWidgetAction_tr2")
	symaddr(&addr_QWidgetAction_tr3, "QWidgetAction_tr3")
	symaddr(&addr_QWidgetAction_connect3, "QWidgetAction_connect3")
	symaddr(&addr_QWidgetAction_connect4, "QWidgetAction_connect4")
	symaddr(&addr_QWidgetAction_metaObject_default, "QWidgetAction_metaObject_default")
	symaddr(&addr_QWidgetAction_event_default, "QWidgetAction_event_default")
	symaddr(&addr_QWidgetAction_eventFilter_default, "QWidgetAction_eventFilter_default")
	symaddr(&addr_QWidgetAction_createWidget_default, "QWidgetAction_createWidget_default")
	symaddr(&addr_QWidgetAction_deleteWidget_default, "QWidgetAction_deleteWidget_default")
	symaddr(&addr_QWidgetAction_timerEvent_default, "QWidgetAction_timerEvent_default")
	symaddr(&addr_QWidgetAction_childEvent_default, "QWidgetAction_childEvent_default")
	symaddr(&addr_QWidgetAction_customEvent_default, "QWidgetAction_customEvent_default")
	symaddr(&addr_QWidgetAction_connectNotify_default, "QWidgetAction_connectNotify_default")
	symaddr(&addr_QWidgetAction_disconnectNotify_default, "QWidgetAction_disconnectNotify_default")
	reg(&pg_QWidgetAction_connect_changed, "QWidgetAction_connect_changed")
	reg(&pg_QWidgetAction_connect_enabledChanged, "QWidgetAction_connect_enabledChanged")
	reg(&pg_QWidgetAction_connect_checkableChanged, "QWidgetAction_connect_checkableChanged")
	reg(&pg_QWidgetAction_connect_visibleChanged, "QWidgetAction_connect_visibleChanged")
	reg(&pg_QWidgetAction_connect_triggered, "QWidgetAction_connect_triggered")
	reg(&pg_QWidgetAction_connect_triggered2, "QWidgetAction_connect_triggered2")
	reg(&pg_QWidgetAction_connect_hovered, "QWidgetAction_connect_hovered")
	reg(&pg_QWidgetAction_connect_toggled, "QWidgetAction_connect_toggled")
	reg(&pg_QWidgetAction_connect_destroyed, "QWidgetAction_connect_destroyed")
	reg(&pg_QWidgetAction_connect_destroyed2, "QWidgetAction_connect_destroyed2")
	reg(&pg_QWidgetAction_connect_objectNameChanged, "QWidgetAction_connect_objectNameChanged")
	reg(&pg_QWidgetAction__disconnect, "QWidgetAction__disconnect")
	reg(&pg_QWidgetAction__set_override, "QWidgetAction__set_override")
	reg(&pg_QWidgetAction_delete, "QWidgetAction_delete")
}

type QWidgetAction struct {
	h uintptr
	*QAction
}

func bindQWidgetAction(h uintptr) *QWidgetAction {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QWidgetAction_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QWidgetAction{h: h, QAction: &QAction{h: __c[0], QObject: &QObject{h: __c[1]}}}
}

func newownedQWidgetAction(h uintptr) *QWidgetAction { return bindQWidgetAction(h) }

func (this *QWidgetAction) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QWidgetAction) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQWidgetAction(h unsafe.Pointer) *QWidgetAction { return bindQWidgetAction(uintptr(h)) }
func UnsafeNewQWidgetActionOwned(h unsafe.Pointer) *QWidgetAction { return newownedQWidgetAction(uintptr(h)) }

func (this *QWidgetAction) UnsafeRelease() {
	this.h = 0
	if this.QAction != nil {
		this.QAction.UnsafeRelease()
	}
}

func NewQWidgetAction(parent *QObject) *QWidgetAction {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QWidgetAction_new, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQWidgetAction(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QWidgetAction_virtIdx, QWidgetAction_enable)
	return __o
}

func QWidgetAction_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QWidgetAction_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QWidgetAction) SetDefaultWidget(w *QWidget) {
	purego.SyscallN(addr_QWidgetAction_setDefaultWidget, this.h, uintptr(w.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(w)
}

func (this *QWidgetAction) DefaultWidget() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QWidgetAction_defaultWidget, this.h)
	runtime.KeepAlive(this)
	return bindQWidget(__r)
}

func (this *QWidgetAction) RequestWidget(parent *QWidget) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QWidgetAction_requestWidget, this.h, uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func (this *QWidgetAction) ReleaseWidget(widget *QWidget) {
	purego.SyscallN(addr_QWidgetAction_releaseWidget, this.h, uintptr(widget.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(widget)
}

func (this *QWidgetAction) CreatedWidgets() []*QWidget {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QWidgetAction_createdWidgets, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QWidget, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = bindQWidget(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QWidgetAction_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QWidgetAction_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QWidgetAction_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QWidgetAction_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QWidgetAction) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QWidgetAction_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QWidgetAction) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QWidgetAction_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QWidgetAction) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QWidgetAction_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QWidgetAction) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QWidgetAction_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QWidgetAction_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QWidgetAction_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QWidgetAction_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QWidgetAction_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QWidgetAction_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QWidgetAction_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QWidgetAction_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QWidgetAction_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QWidgetAction) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QWidgetAction_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QWidgetAction) EventDefault(arg1 *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QWidgetAction_event_default, this.h, uintptr(arg1.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return __r != 0
}

func (this *QWidgetAction) EventFilterDefault(arg1 *QObject, arg2 *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QWidgetAction_eventFilter_default, this.h, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
	return __r != 0
}

func (this *QWidgetAction) CreateWidgetDefault(parent *QWidget) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QWidgetAction_createWidget_default, this.h, uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func (this *QWidgetAction) DeleteWidgetDefault(widget *QWidget) {
	purego.SyscallN(addr_QWidgetAction_deleteWidget_default, this.h, uintptr(widget.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(widget)
}

func (this *QWidgetAction) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QWidgetAction_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QWidgetAction) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QWidgetAction_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QWidgetAction) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QWidgetAction_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QWidgetAction) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QWidgetAction_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QWidgetAction) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QWidgetAction_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QWidgetAction) OnChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QWidgetAction_connect_changed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QWidgetAction) OnEnabledChanged(fn func(bool)) int64 {
	id := registerCallback(func(enabled int32) { fn(enabled != 0) })
	pg_QWidgetAction_connect_enabledChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QWidgetAction) OnCheckableChanged(fn func(bool)) int64 {
	id := registerCallback(func(checkable int32) { fn(checkable != 0) })
	pg_QWidgetAction_connect_checkableChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QWidgetAction) OnVisibleChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QWidgetAction_connect_visibleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QWidgetAction) OnTriggered(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QWidgetAction_connect_triggered(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QWidgetAction) OnTriggeredWithChecked(fn func(bool)) int64 {
	id := registerCallback(func(checked int32) { fn(checked != 0) })
	pg_QWidgetAction_connect_triggered2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QWidgetAction) OnHovered(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QWidgetAction_connect_hovered(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QWidgetAction) OnToggled(fn func(bool)) int64 {
	id := registerCallback(func(arg1 int32) { fn(arg1 != 0) })
	pg_QWidgetAction_connect_toggled(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QWidgetAction) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QWidgetAction_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QWidgetAction) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QWidgetAction_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QWidgetAction) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QWidgetAction_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QWidgetAction) Disconnect(handle int64) {
	pg_QWidgetAction__disconnect(handle)
	unregisterCallback(handle)
}

func QWidgetAction_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case 1237236456312824991:
		return 1
	case 5217623514630858297:
		return 2
	case -8470300370605365395:
		return 3
	case -6493542048625819130:
		return 4
	case -5409927581812374498:
		return 5
	case -1185315367526857833:
		return 6
	case 3769868229123868474:
		return 7
	case 7379029982170503008:
		return 8
	case -568369772982215962:
		return 9
	}
	return -1
}

func QWidgetAction_enable(h uintptr, idx int32) { pg_QWidgetAction__set_override(h, idx) }

func (this *QWidgetAction) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QWidgetAction) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(arg1 uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(arg1)))) }) // slot "event"
}

func (this *QWidgetAction) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(arg1 uintptr, arg2 uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(arg1)), UnsafeNewQEvent(unsafe.Pointer(arg2)))) }) // slot "eventFilter"
}

func (this *QWidgetAction) OnCreateWidget(fn func(*QWidget) *QWidget) {
	registerVirt(this.UnsafeGoID(), -8470300370605365395, func(parent uintptr) uintptr { return uintptr(fn(UnsafeNewQWidget(unsafe.Pointer(parent))).UnsafePointer()) }) // slot "createWidget"
}

func (this *QWidgetAction) OnDeleteWidget(fn func(*QWidget)) {
	registerVirt(this.UnsafeGoID(), -6493542048625819130, func(widget uintptr) { fn(UnsafeNewQWidget(unsafe.Pointer(widget))) }) // slot "deleteWidget"
}

func (this *QWidgetAction) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QWidgetAction) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QWidgetAction) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QWidgetAction) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QWidgetAction) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QWidgetAction) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QWidgetAction_delete(this.h)
	}
	this.UnsafeRelease()
}
