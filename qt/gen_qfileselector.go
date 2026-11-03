package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QFileSelector_base func(uintptr) uintptr
	addr_QFileSelector_new uintptr
	addr_QFileSelector_new2 uintptr
	addr_QFileSelector_tr uintptr
	addr_QFileSelector_select uintptr
	addr_QFileSelector_select2 uintptr
	addr_QFileSelector_extraSelectors uintptr
	addr_QFileSelector_setExtraSelectors uintptr
	addr_QFileSelector_allSelectors uintptr
	addr_QFileSelector_connect uintptr
	addr_QFileSelector_connect2 uintptr
	addr_QFileSelector_sender uintptr
	addr_QFileSelector_senderSignalIndex uintptr
	addr_QFileSelector_receivers uintptr
	addr_QFileSelector_isSignalConnected uintptr
	addr_QFileSelector_tr2 uintptr
	addr_QFileSelector_tr3 uintptr
	addr_QFileSelector_connect3 uintptr
	addr_QFileSelector_connect4 uintptr
	addr_QFileSelector_metaObject_default uintptr
	addr_QFileSelector_event_default uintptr
	addr_QFileSelector_eventFilter_default uintptr
	addr_QFileSelector_timerEvent_default uintptr
	addr_QFileSelector_childEvent_default uintptr
	addr_QFileSelector_customEvent_default uintptr
	addr_QFileSelector_connectNotify_default uintptr
	addr_QFileSelector_disconnectNotify_default uintptr
	pg_QFileSelector_connect_destroyed func(uintptr, int64)
	pg_QFileSelector_connect_destroyed2 func(uintptr, int64)
	pg_QFileSelector_connect_objectNameChanged func(uintptr, int64)
	pg_QFileSelector__disconnect func(int64)
	pg_QFileSelector__set_override func(uintptr, int32)
	pg_QFileSelector_delete func(uintptr)
)

func init() {
	reg(&pg_QFileSelector_base, "QFileSelector_base")
	symaddr(&addr_QFileSelector_new, "QFileSelector_new")
	symaddr(&addr_QFileSelector_new2, "QFileSelector_new2")
	symaddr(&addr_QFileSelector_tr, "QFileSelector_tr")
	symaddr(&addr_QFileSelector_select, "QFileSelector_select")
	symaddr(&addr_QFileSelector_select2, "QFileSelector_select2")
	symaddr(&addr_QFileSelector_extraSelectors, "QFileSelector_extraSelectors")
	symaddr(&addr_QFileSelector_setExtraSelectors, "QFileSelector_setExtraSelectors")
	symaddr(&addr_QFileSelector_allSelectors, "QFileSelector_allSelectors")
	symaddr(&addr_QFileSelector_connect, "QFileSelector_connect")
	symaddr(&addr_QFileSelector_connect2, "QFileSelector_connect2")
	symaddr(&addr_QFileSelector_sender, "QFileSelector_sender")
	symaddr(&addr_QFileSelector_senderSignalIndex, "QFileSelector_senderSignalIndex")
	symaddr(&addr_QFileSelector_receivers, "QFileSelector_receivers")
	symaddr(&addr_QFileSelector_isSignalConnected, "QFileSelector_isSignalConnected")
	symaddr(&addr_QFileSelector_tr2, "QFileSelector_tr2")
	symaddr(&addr_QFileSelector_tr3, "QFileSelector_tr3")
	symaddr(&addr_QFileSelector_connect3, "QFileSelector_connect3")
	symaddr(&addr_QFileSelector_connect4, "QFileSelector_connect4")
	symaddr(&addr_QFileSelector_metaObject_default, "QFileSelector_metaObject_default")
	symaddr(&addr_QFileSelector_event_default, "QFileSelector_event_default")
	symaddr(&addr_QFileSelector_eventFilter_default, "QFileSelector_eventFilter_default")
	symaddr(&addr_QFileSelector_timerEvent_default, "QFileSelector_timerEvent_default")
	symaddr(&addr_QFileSelector_childEvent_default, "QFileSelector_childEvent_default")
	symaddr(&addr_QFileSelector_customEvent_default, "QFileSelector_customEvent_default")
	symaddr(&addr_QFileSelector_connectNotify_default, "QFileSelector_connectNotify_default")
	symaddr(&addr_QFileSelector_disconnectNotify_default, "QFileSelector_disconnectNotify_default")
	reg(&pg_QFileSelector_connect_destroyed, "QFileSelector_connect_destroyed")
	reg(&pg_QFileSelector_connect_destroyed2, "QFileSelector_connect_destroyed2")
	reg(&pg_QFileSelector_connect_objectNameChanged, "QFileSelector_connect_objectNameChanged")
	reg(&pg_QFileSelector__disconnect, "QFileSelector__disconnect")
	reg(&pg_QFileSelector__set_override, "QFileSelector__set_override")
	reg(&pg_QFileSelector_delete, "QFileSelector_delete")
}

type QFileSelector struct {
	h uintptr
	*QObject
}

func bindQFileSelector(h uintptr) *QFileSelector {
	if h == 0 {
		return nil
	}
	return &QFileSelector{h: h, QObject: bindQObject(pg_QFileSelector_base(h))}
}

func newownedQFileSelector(h uintptr) *QFileSelector { return bindQFileSelector(h) }

func (this *QFileSelector) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QFileSelector) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQFileSelector(h unsafe.Pointer) *QFileSelector { return bindQFileSelector(uintptr(h)) }
func UnsafeNewQFileSelectorOwned(h unsafe.Pointer) *QFileSelector { return newownedQFileSelector(uintptr(h)) }

func (this *QFileSelector) UnsafeRelease() {
	this.h = 0
	if this.QObject != nil {
		this.QObject.UnsafeRelease()
	}
}

func NewQFileSelector() *QFileSelector {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QFileSelector_new, uintptr(__goid))
	__o := newownedQFileSelector(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QFileSelector_virtIdx, QFileSelector_enable)
	return __o
}

func NewQFileSelector2(parent *QObject) *QFileSelector {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QFileSelector_new2, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQFileSelector(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QFileSelector_virtIdx, QFileSelector_enable)
	return __o
}

func QFileSelector_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFileSelector_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QFileSelector) Select(filePath string) string {
	var filePath_p uintptr
	if len(filePath) > 0 {
		filePath_p = uintptr(unsafe.Pointer(unsafe.StringData(filePath)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFileSelector_select, this.h, filePath_p, uintptr(len(filePath)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(filePath)
	return pgStr(__r, __n)
}

func (this *QFileSelector) SelectWithFilePath(filePath *QUrl) *QUrl {
	__r, _, _ := purego.SyscallN(addr_QFileSelector_select2, this.h, uintptr(filePath.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(filePath)
	return newownedQUrl(__r)
}

func (this *QFileSelector) ExtraSelectors() []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFileSelector_extraSelectors, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]string, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = pgGoStr(__sl[__i])
		pgFree(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QFileSelector) SetExtraSelectors(list []string) {
	list_arr := make([]uintptr, len(list))
	for __i, __e := range list {
		list_arr[__i] = pgCString(__e)
	}
	var list_p uintptr
	if len(list_arr) > 0 {
		list_p = uintptr(unsafe.Pointer(&list_arr[0]))
	}
	purego.SyscallN(addr_QFileSelector_setExtraSelectors, this.h, list_p, uintptr(len(list)))
	runtime.KeepAlive(this)
	for __i := range list_arr {
		pgFree(list_arr[__i])
	}
	runtime.KeepAlive(list_arr)
}

func (this *QFileSelector) AllSelectors() []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFileSelector_allSelectors, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]string, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = pgGoStr(__sl[__i])
		pgFree(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QFileSelector_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QFileSelector_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QFileSelector_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QFileSelector_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QFileSelector) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QFileSelector_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QFileSelector) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QFileSelector_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QFileSelector) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QFileSelector_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QFileSelector) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QFileSelector_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QFileSelector_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFileSelector_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QFileSelector_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFileSelector_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QFileSelector_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QFileSelector_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QFileSelector_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QFileSelector_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QFileSelector) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QFileSelector_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QFileSelector) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QFileSelector_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QFileSelector) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QFileSelector_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QFileSelector) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QFileSelector_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFileSelector) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QFileSelector_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFileSelector) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QFileSelector_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFileSelector) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QFileSelector_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QFileSelector) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QFileSelector_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QFileSelector) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QFileSelector_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFileSelector) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QFileSelector_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFileSelector) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QFileSelector_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFileSelector) Disconnect(handle int64) {
	pg_QFileSelector__disconnect(handle)
	unregisterCallback(handle)
}

func QFileSelector_virtIdx(slot int64) int32 {
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

func QFileSelector_enable(h uintptr, idx int32) { pg_QFileSelector__set_override(h, idx) }

func (this *QFileSelector) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QFileSelector) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QFileSelector) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QFileSelector) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QFileSelector) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QFileSelector) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QFileSelector) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QFileSelector) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QFileSelector) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QFileSelector_delete(this.h)
	}
	this.UnsafeRelease()
}
