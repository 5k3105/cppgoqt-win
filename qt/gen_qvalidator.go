package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QValidator_base func(uintptr) uintptr
	addr_QValidator_tr uintptr
	addr_QValidator_setLocale uintptr
	addr_QValidator_locale uintptr
	addr_QValidator_validate uintptr
	addr_QValidator_fixup uintptr
	addr_QValidator_changed uintptr
	addr_QValidator_connect uintptr
	addr_QValidator_connect2 uintptr
	addr_QValidator_sender uintptr
	addr_QValidator_senderSignalIndex uintptr
	addr_QValidator_receivers uintptr
	addr_QValidator_isSignalConnected uintptr
	addr_QValidator_tr2 uintptr
	addr_QValidator_tr3 uintptr
	addr_QValidator_connect3 uintptr
	addr_QValidator_connect4 uintptr
	addr_QValidator_fixup_default uintptr
	addr_QValidator_metaObject_default uintptr
	addr_QValidator_event_default uintptr
	addr_QValidator_eventFilter_default uintptr
	addr_QValidator_timerEvent_default uintptr
	addr_QValidator_childEvent_default uintptr
	addr_QValidator_customEvent_default uintptr
	addr_QValidator_connectNotify_default uintptr
	addr_QValidator_disconnectNotify_default uintptr
	pg_QValidator_connect_changed func(uintptr, int64)
	pg_QValidator_connect_destroyed func(uintptr, int64)
	pg_QValidator_connect_destroyed2 func(uintptr, int64)
	pg_QValidator_connect_objectNameChanged func(uintptr, int64)
	pg_QValidator__disconnect func(int64)
	pg_QValidator__set_override func(uintptr, int32)
	pg_QValidator__Go_new func(int64, uintptr) uintptr
	pg_QValidator_delete func(uintptr)
)

func init() {
	reg(&pg_QValidator_base, "QValidator_base")
	symaddr(&addr_QValidator_tr, "QValidator_tr")
	symaddr(&addr_QValidator_setLocale, "QValidator_setLocale")
	symaddr(&addr_QValidator_locale, "QValidator_locale")
	symaddr(&addr_QValidator_validate, "QValidator_validate")
	symaddr(&addr_QValidator_fixup, "QValidator_fixup")
	symaddr(&addr_QValidator_changed, "QValidator_changed")
	symaddr(&addr_QValidator_connect, "QValidator_connect")
	symaddr(&addr_QValidator_connect2, "QValidator_connect2")
	symaddr(&addr_QValidator_sender, "QValidator_sender")
	symaddr(&addr_QValidator_senderSignalIndex, "QValidator_senderSignalIndex")
	symaddr(&addr_QValidator_receivers, "QValidator_receivers")
	symaddr(&addr_QValidator_isSignalConnected, "QValidator_isSignalConnected")
	symaddr(&addr_QValidator_tr2, "QValidator_tr2")
	symaddr(&addr_QValidator_tr3, "QValidator_tr3")
	symaddr(&addr_QValidator_connect3, "QValidator_connect3")
	symaddr(&addr_QValidator_connect4, "QValidator_connect4")
	symaddr(&addr_QValidator_fixup_default, "QValidator_fixup_default")
	symaddr(&addr_QValidator_metaObject_default, "QValidator_metaObject_default")
	symaddr(&addr_QValidator_event_default, "QValidator_event_default")
	symaddr(&addr_QValidator_eventFilter_default, "QValidator_eventFilter_default")
	symaddr(&addr_QValidator_timerEvent_default, "QValidator_timerEvent_default")
	symaddr(&addr_QValidator_childEvent_default, "QValidator_childEvent_default")
	symaddr(&addr_QValidator_customEvent_default, "QValidator_customEvent_default")
	symaddr(&addr_QValidator_connectNotify_default, "QValidator_connectNotify_default")
	symaddr(&addr_QValidator_disconnectNotify_default, "QValidator_disconnectNotify_default")
	reg(&pg_QValidator_connect_changed, "QValidator_connect_changed")
	reg(&pg_QValidator_connect_destroyed, "QValidator_connect_destroyed")
	reg(&pg_QValidator_connect_destroyed2, "QValidator_connect_destroyed2")
	reg(&pg_QValidator_connect_objectNameChanged, "QValidator_connect_objectNameChanged")
	reg(&pg_QValidator__disconnect, "QValidator__disconnect")
	reg(&pg_QValidator__set_override, "QValidator__set_override")
	reg(&pg_QValidator__Go_new, "QValidator__Go_new")
	reg(&pg_QValidator_delete, "QValidator_delete")
}

type QValidator struct {
	h uintptr
	*QObject
}

func bindQValidator(h uintptr) *QValidator {
	if h == 0 {
		return nil
	}
	return &QValidator{h: h, QObject: bindQObject(pg_QValidator_base(h))}
}

func newownedQValidator(h uintptr) *QValidator { return bindQValidator(h) }

func (this *QValidator) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QValidator) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQValidator(h unsafe.Pointer) *QValidator { return bindQValidator(uintptr(h)) }
func UnsafeNewQValidatorOwned(h unsafe.Pointer) *QValidator { return newownedQValidator(uintptr(h)) }

func (this *QValidator) UnsafeRelease() {
	this.h = 0
	if this.QObject != nil {
		this.QObject.UnsafeRelease()
	}
}

func QValidator_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QValidator_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QValidator) SetLocale(locale *QLocale) {
	purego.SyscallN(addr_QValidator_setLocale, this.h, uintptr(locale.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(locale)
}

func (this *QValidator) Locale() *QLocale {
	__r, _, _ := purego.SyscallN(addr_QValidator_locale, this.h)
	runtime.KeepAlive(this)
	return newownedQLocale(__r)
}

func (this *QValidator) Validate(arg1 *QString, arg2 int) (QValidator__State, int) {
	arg2_io := int32(arg2)
	__r0, _, _ := purego.SyscallN(addr_QValidator_validate, this.h, uintptr(arg1.UnsafePointer()), uintptr(unsafe.Pointer(&arg2_io)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return QValidator__State(__r), int(arg2_io)
}

func (this *QValidator) Fixup(arg1 *QString) {
	purego.SyscallN(addr_QValidator_fixup, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QValidator) Changed() {
	purego.SyscallN(addr_QValidator_changed, this.h)
	runtime.KeepAlive(this)
}

func QValidator_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QValidator_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QValidator_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QValidator_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QValidator) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QValidator_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QValidator) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QValidator_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QValidator) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QValidator_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QValidator) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QValidator_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QValidator_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QValidator_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QValidator_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QValidator_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QValidator_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QValidator_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QValidator_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QValidator_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QValidator) FixupDefault(arg1 *QString) {
	purego.SyscallN(addr_QValidator_fixup_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QValidator) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QValidator_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QValidator) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QValidator_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QValidator) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QValidator_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QValidator) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QValidator_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QValidator) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QValidator_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QValidator) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QValidator_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QValidator) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QValidator_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QValidator) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QValidator_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QValidator) OnChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QValidator_connect_changed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QValidator) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QValidator_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QValidator) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QValidator_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QValidator) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QValidator_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QValidator) Disconnect(handle int64) {
	pg_QValidator__disconnect(handle)
	unregisterCallback(handle)
}

func QValidator_virtIdx(slot int64) int32 {
	switch slot {
	case 6784179835071232555:
		return 0
	case 1607084412591819503:
		return 1
	case -4855206404746966715:
		return 2
	case 1237236456312824991:
		return 3
	case 5217623514630858297:
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

func QValidator_enable(h uintptr, idx int32) { pg_QValidator__set_override(h, idx) }

func NewQValidator() *QValidator {
	id := nextVirtID()
	o := newownedQValidator(pg_QValidator__Go_new(id, 0))
	o.UnsafeSetGoID(id)
	registerInstance(id, o.h, QValidator_virtIdx, QValidator_enable)
	return o
}

func NewQValidator2(parent *QObject) *QValidator {
	id := nextVirtID()
	o := newownedQValidator(pg_QValidator__Go_new(id, uintptr(parent.UnsafePointer())))
	o.UnsafeSetGoID(id)
	registerInstance(id, o.h, QValidator_virtIdx, QValidator_enable)
	return o
}

func (this *QValidator) OnValidate(fn func(*QString, *int32) QValidator__State) {
	registerVirt(this.UnsafeGoID(), 6784179835071232555, func(arg1 uintptr, arg2 uintptr) int32 { return int32(fn(UnsafeNewQString(unsafe.Pointer(arg1)), (*int32)(unsafe.Pointer(arg2)))) }) // slot "validate"
}

func (this *QValidator) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QValidator) OnFixup(fn func(*QString)) {
	registerVirt(this.UnsafeGoID(), -4855206404746966715, func(arg1 uintptr) { fn(UnsafeNewQString(unsafe.Pointer(arg1))) }) // slot "fixup"
}

func (this *QValidator) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QValidator) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QValidator) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QValidator) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QValidator) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QValidator) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QValidator) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QValidator) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QValidator_delete(this.h)
	}
	this.UnsafeRelease()
}
