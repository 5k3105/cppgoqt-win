package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QRegularExpressionValidator_chain uintptr
	addr_QRegularExpressionValidator_new uintptr
	addr_QRegularExpressionValidator_new2 uintptr
	addr_QRegularExpressionValidator_new3 uintptr
	addr_QRegularExpressionValidator_new4 uintptr
	addr_QRegularExpressionValidator_tr uintptr
	addr_QRegularExpressionValidator_regularExpression uintptr
	addr_QRegularExpressionValidator_setRegularExpression uintptr
	addr_QRegularExpressionValidator_regularExpressionChanged uintptr
	addr_QRegularExpressionValidator_connect uintptr
	addr_QRegularExpressionValidator_connect2 uintptr
	addr_QRegularExpressionValidator_sender uintptr
	addr_QRegularExpressionValidator_senderSignalIndex uintptr
	addr_QRegularExpressionValidator_receivers uintptr
	addr_QRegularExpressionValidator_isSignalConnected uintptr
	addr_QRegularExpressionValidator_tr2 uintptr
	addr_QRegularExpressionValidator_tr3 uintptr
	addr_QRegularExpressionValidator_connect3 uintptr
	addr_QRegularExpressionValidator_connect4 uintptr
	addr_QRegularExpressionValidator_metaObject_default uintptr
	addr_QRegularExpressionValidator_validate_default uintptr
	addr_QRegularExpressionValidator_fixup_default uintptr
	addr_QRegularExpressionValidator_event_default uintptr
	addr_QRegularExpressionValidator_eventFilter_default uintptr
	addr_QRegularExpressionValidator_timerEvent_default uintptr
	addr_QRegularExpressionValidator_childEvent_default uintptr
	addr_QRegularExpressionValidator_customEvent_default uintptr
	addr_QRegularExpressionValidator_connectNotify_default uintptr
	addr_QRegularExpressionValidator_disconnectNotify_default uintptr
	pg_QRegularExpressionValidator_connect_regularExpressionChanged func(uintptr, int64)
	pg_QRegularExpressionValidator_connect_changed func(uintptr, int64)
	pg_QRegularExpressionValidator_connect_destroyed func(uintptr, int64)
	pg_QRegularExpressionValidator_connect_destroyed2 func(uintptr, int64)
	pg_QRegularExpressionValidator_connect_objectNameChanged func(uintptr, int64)
	pg_QRegularExpressionValidator__disconnect func(int64)
	pg_QRegularExpressionValidator__set_override func(uintptr, int32)
	pg_QRegularExpressionValidator_delete func(uintptr)
)

func init() {
	symaddr(&addr_QRegularExpressionValidator_chain, "QRegularExpressionValidator_chain")
	symaddr(&addr_QRegularExpressionValidator_new, "QRegularExpressionValidator_new")
	symaddr(&addr_QRegularExpressionValidator_new2, "QRegularExpressionValidator_new2")
	symaddr(&addr_QRegularExpressionValidator_new3, "QRegularExpressionValidator_new3")
	symaddr(&addr_QRegularExpressionValidator_new4, "QRegularExpressionValidator_new4")
	symaddr(&addr_QRegularExpressionValidator_tr, "QRegularExpressionValidator_tr")
	symaddr(&addr_QRegularExpressionValidator_regularExpression, "QRegularExpressionValidator_regularExpression")
	symaddr(&addr_QRegularExpressionValidator_setRegularExpression, "QRegularExpressionValidator_setRegularExpression")
	symaddr(&addr_QRegularExpressionValidator_regularExpressionChanged, "QRegularExpressionValidator_regularExpressionChanged")
	symaddr(&addr_QRegularExpressionValidator_connect, "QRegularExpressionValidator_connect")
	symaddr(&addr_QRegularExpressionValidator_connect2, "QRegularExpressionValidator_connect2")
	symaddr(&addr_QRegularExpressionValidator_sender, "QRegularExpressionValidator_sender")
	symaddr(&addr_QRegularExpressionValidator_senderSignalIndex, "QRegularExpressionValidator_senderSignalIndex")
	symaddr(&addr_QRegularExpressionValidator_receivers, "QRegularExpressionValidator_receivers")
	symaddr(&addr_QRegularExpressionValidator_isSignalConnected, "QRegularExpressionValidator_isSignalConnected")
	symaddr(&addr_QRegularExpressionValidator_tr2, "QRegularExpressionValidator_tr2")
	symaddr(&addr_QRegularExpressionValidator_tr3, "QRegularExpressionValidator_tr3")
	symaddr(&addr_QRegularExpressionValidator_connect3, "QRegularExpressionValidator_connect3")
	symaddr(&addr_QRegularExpressionValidator_connect4, "QRegularExpressionValidator_connect4")
	symaddr(&addr_QRegularExpressionValidator_metaObject_default, "QRegularExpressionValidator_metaObject_default")
	symaddr(&addr_QRegularExpressionValidator_validate_default, "QRegularExpressionValidator_validate_default")
	symaddr(&addr_QRegularExpressionValidator_fixup_default, "QRegularExpressionValidator_fixup_default")
	symaddr(&addr_QRegularExpressionValidator_event_default, "QRegularExpressionValidator_event_default")
	symaddr(&addr_QRegularExpressionValidator_eventFilter_default, "QRegularExpressionValidator_eventFilter_default")
	symaddr(&addr_QRegularExpressionValidator_timerEvent_default, "QRegularExpressionValidator_timerEvent_default")
	symaddr(&addr_QRegularExpressionValidator_childEvent_default, "QRegularExpressionValidator_childEvent_default")
	symaddr(&addr_QRegularExpressionValidator_customEvent_default, "QRegularExpressionValidator_customEvent_default")
	symaddr(&addr_QRegularExpressionValidator_connectNotify_default, "QRegularExpressionValidator_connectNotify_default")
	symaddr(&addr_QRegularExpressionValidator_disconnectNotify_default, "QRegularExpressionValidator_disconnectNotify_default")
	reg(&pg_QRegularExpressionValidator_connect_regularExpressionChanged, "QRegularExpressionValidator_connect_regularExpressionChanged")
	reg(&pg_QRegularExpressionValidator_connect_changed, "QRegularExpressionValidator_connect_changed")
	reg(&pg_QRegularExpressionValidator_connect_destroyed, "QRegularExpressionValidator_connect_destroyed")
	reg(&pg_QRegularExpressionValidator_connect_destroyed2, "QRegularExpressionValidator_connect_destroyed2")
	reg(&pg_QRegularExpressionValidator_connect_objectNameChanged, "QRegularExpressionValidator_connect_objectNameChanged")
	reg(&pg_QRegularExpressionValidator__disconnect, "QRegularExpressionValidator__disconnect")
	reg(&pg_QRegularExpressionValidator__set_override, "QRegularExpressionValidator__set_override")
	reg(&pg_QRegularExpressionValidator_delete, "QRegularExpressionValidator_delete")
}

type QRegularExpressionValidator struct {
	h uintptr
	*QValidator
}

func bindQRegularExpressionValidator(h uintptr) *QRegularExpressionValidator {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QRegularExpressionValidator_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QRegularExpressionValidator{h: h, QValidator: &QValidator{h: __c[0], QObject: &QObject{h: __c[1]}}}
}

func newownedQRegularExpressionValidator(h uintptr) *QRegularExpressionValidator { return bindQRegularExpressionValidator(h) }

func (this *QRegularExpressionValidator) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QRegularExpressionValidator) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQRegularExpressionValidator(h unsafe.Pointer) *QRegularExpressionValidator { return bindQRegularExpressionValidator(uintptr(h)) }
func UnsafeNewQRegularExpressionValidatorOwned(h unsafe.Pointer) *QRegularExpressionValidator { return newownedQRegularExpressionValidator(uintptr(h)) }

func (this *QRegularExpressionValidator) UnsafeRelease() {
	this.h = 0
	if this.QValidator != nil {
		this.QValidator.UnsafeRelease()
	}
}

func NewQRegularExpressionValidator() *QRegularExpressionValidator {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QRegularExpressionValidator_new, uintptr(__goid))
	__o := newownedQRegularExpressionValidator(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QRegularExpressionValidator_virtIdx, QRegularExpressionValidator_enable)
	return __o
}

func NewQRegularExpressionValidator2(re *QRegularExpression) *QRegularExpressionValidator {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QRegularExpressionValidator_new2, uintptr(__goid), uintptr(re.UnsafePointer()))
	__o := newownedQRegularExpressionValidator(__r)
	runtime.KeepAlive(re)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QRegularExpressionValidator_virtIdx, QRegularExpressionValidator_enable)
	return __o
}

func NewQRegularExpressionValidator3(parent *QObject) *QRegularExpressionValidator {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QRegularExpressionValidator_new3, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQRegularExpressionValidator(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QRegularExpressionValidator_virtIdx, QRegularExpressionValidator_enable)
	return __o
}

func NewQRegularExpressionValidator4(re *QRegularExpression, parent *QObject) *QRegularExpressionValidator {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QRegularExpressionValidator_new4, uintptr(__goid), uintptr(re.UnsafePointer()), uintptr(parent.UnsafePointer()))
	__o := newownedQRegularExpressionValidator(__r)
	runtime.KeepAlive(re)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QRegularExpressionValidator_virtIdx, QRegularExpressionValidator_enable)
	return __o
}

func QRegularExpressionValidator_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QRegularExpressionValidator_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QRegularExpressionValidator) RegularExpression() *QRegularExpression {
	__r, _, _ := purego.SyscallN(addr_QRegularExpressionValidator_regularExpression, this.h)
	runtime.KeepAlive(this)
	return newownedQRegularExpression(__r)
}

func (this *QRegularExpressionValidator) SetRegularExpression(re *QRegularExpression) {
	purego.SyscallN(addr_QRegularExpressionValidator_setRegularExpression, this.h, uintptr(re.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(re)
}

func (this *QRegularExpressionValidator) RegularExpressionChanged(re *QRegularExpression) {
	purego.SyscallN(addr_QRegularExpressionValidator_regularExpressionChanged, this.h, uintptr(re.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(re)
}

func QRegularExpressionValidator_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QRegularExpressionValidator_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QRegularExpressionValidator_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QRegularExpressionValidator_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QRegularExpressionValidator) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QRegularExpressionValidator_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QRegularExpressionValidator) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QRegularExpressionValidator_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QRegularExpressionValidator) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QRegularExpressionValidator_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QRegularExpressionValidator) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QRegularExpressionValidator_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QRegularExpressionValidator_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QRegularExpressionValidator_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QRegularExpressionValidator_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QRegularExpressionValidator_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QRegularExpressionValidator_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QRegularExpressionValidator_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QRegularExpressionValidator_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QRegularExpressionValidator_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QRegularExpressionValidator) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QRegularExpressionValidator_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QRegularExpressionValidator) ValidateDefault(input *QString, pos int) (QValidator__State, int) {
	pos_io := int32(pos)
	__r0, _, _ := purego.SyscallN(addr_QRegularExpressionValidator_validate_default, this.h, uintptr(input.UnsafePointer()), uintptr(unsafe.Pointer(&pos_io)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(input)
	return QValidator__State(__r), int(pos_io)
}

func (this *QRegularExpressionValidator) FixupDefault(arg1 *QString) {
	purego.SyscallN(addr_QRegularExpressionValidator_fixup_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QRegularExpressionValidator) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QRegularExpressionValidator_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QRegularExpressionValidator) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QRegularExpressionValidator_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QRegularExpressionValidator) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QRegularExpressionValidator_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QRegularExpressionValidator) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QRegularExpressionValidator_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QRegularExpressionValidator) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QRegularExpressionValidator_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QRegularExpressionValidator) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QRegularExpressionValidator_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QRegularExpressionValidator) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QRegularExpressionValidator_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QRegularExpressionValidator) OnRegularExpressionChanged(fn func(*QRegularExpression)) int64 {
	id := registerCallback(func(re uintptr) { fn(UnsafeNewQRegularExpression(unsafe.Pointer(re))) })
	pg_QRegularExpressionValidator_connect_regularExpressionChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QRegularExpressionValidator) OnChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QRegularExpressionValidator_connect_changed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QRegularExpressionValidator) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QRegularExpressionValidator_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QRegularExpressionValidator) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QRegularExpressionValidator_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QRegularExpressionValidator) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QRegularExpressionValidator_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QRegularExpressionValidator) Disconnect(handle int64) {
	pg_QRegularExpressionValidator__disconnect(handle)
	unregisterCallback(handle)
}

func QRegularExpressionValidator_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case 6784179835071232555:
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

func QRegularExpressionValidator_enable(h uintptr, idx int32) { pg_QRegularExpressionValidator__set_override(h, idx) }

func (this *QRegularExpressionValidator) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QRegularExpressionValidator) OnValidate(fn func(*QString, *int32) QValidator__State) {
	registerVirt(this.UnsafeGoID(), 6784179835071232555, func(input uintptr, pos uintptr) int32 { return int32(fn(UnsafeNewQString(unsafe.Pointer(input)), (*int32)(unsafe.Pointer(pos)))) }) // slot "validate"
}

func (this *QRegularExpressionValidator) OnFixup(fn func(*QString)) {
	registerVirt(this.UnsafeGoID(), -4855206404746966715, func(arg1 uintptr) { fn(UnsafeNewQString(unsafe.Pointer(arg1))) }) // slot "fixup"
}

func (this *QRegularExpressionValidator) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QRegularExpressionValidator) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QRegularExpressionValidator) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QRegularExpressionValidator) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QRegularExpressionValidator) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QRegularExpressionValidator) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QRegularExpressionValidator) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QRegularExpressionValidator) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QRegularExpressionValidator_delete(this.h)
	}
	this.UnsafeRelease()
}
