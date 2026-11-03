package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QIntValidator_chain uintptr
	addr_QIntValidator_new uintptr
	addr_QIntValidator_new2 uintptr
	addr_QIntValidator_new3 uintptr
	addr_QIntValidator_new4 uintptr
	addr_QIntValidator_tr uintptr
	addr_QIntValidator_setBottom uintptr
	addr_QIntValidator_setTop uintptr
	addr_QIntValidator_setRange uintptr
	addr_QIntValidator_bottom uintptr
	addr_QIntValidator_top uintptr
	addr_QIntValidator_bottomChanged uintptr
	addr_QIntValidator_topChanged uintptr
	addr_QIntValidator_connect uintptr
	addr_QIntValidator_connect2 uintptr
	addr_QIntValidator_sender uintptr
	addr_QIntValidator_senderSignalIndex uintptr
	addr_QIntValidator_receivers uintptr
	addr_QIntValidator_isSignalConnected uintptr
	addr_QIntValidator_tr2 uintptr
	addr_QIntValidator_tr3 uintptr
	addr_QIntValidator_connect3 uintptr
	addr_QIntValidator_connect4 uintptr
	addr_QIntValidator_metaObject_default uintptr
	addr_QIntValidator_validate_default uintptr
	addr_QIntValidator_fixup_default uintptr
	addr_QIntValidator_event_default uintptr
	addr_QIntValidator_eventFilter_default uintptr
	addr_QIntValidator_timerEvent_default uintptr
	addr_QIntValidator_childEvent_default uintptr
	addr_QIntValidator_customEvent_default uintptr
	addr_QIntValidator_connectNotify_default uintptr
	addr_QIntValidator_disconnectNotify_default uintptr
	pg_QIntValidator_connect_bottomChanged func(uintptr, int64)
	pg_QIntValidator_connect_topChanged func(uintptr, int64)
	pg_QIntValidator_connect_changed func(uintptr, int64)
	pg_QIntValidator_connect_destroyed func(uintptr, int64)
	pg_QIntValidator_connect_destroyed2 func(uintptr, int64)
	pg_QIntValidator_connect_objectNameChanged func(uintptr, int64)
	pg_QIntValidator__disconnect func(int64)
	pg_QIntValidator__set_override func(uintptr, int32)
	pg_QIntValidator_delete func(uintptr)
)

func init() {
	symaddr(&addr_QIntValidator_chain, "QIntValidator_chain")
	symaddr(&addr_QIntValidator_new, "QIntValidator_new")
	symaddr(&addr_QIntValidator_new2, "QIntValidator_new2")
	symaddr(&addr_QIntValidator_new3, "QIntValidator_new3")
	symaddr(&addr_QIntValidator_new4, "QIntValidator_new4")
	symaddr(&addr_QIntValidator_tr, "QIntValidator_tr")
	symaddr(&addr_QIntValidator_setBottom, "QIntValidator_setBottom")
	symaddr(&addr_QIntValidator_setTop, "QIntValidator_setTop")
	symaddr(&addr_QIntValidator_setRange, "QIntValidator_setRange")
	symaddr(&addr_QIntValidator_bottom, "QIntValidator_bottom")
	symaddr(&addr_QIntValidator_top, "QIntValidator_top")
	symaddr(&addr_QIntValidator_bottomChanged, "QIntValidator_bottomChanged")
	symaddr(&addr_QIntValidator_topChanged, "QIntValidator_topChanged")
	symaddr(&addr_QIntValidator_connect, "QIntValidator_connect")
	symaddr(&addr_QIntValidator_connect2, "QIntValidator_connect2")
	symaddr(&addr_QIntValidator_sender, "QIntValidator_sender")
	symaddr(&addr_QIntValidator_senderSignalIndex, "QIntValidator_senderSignalIndex")
	symaddr(&addr_QIntValidator_receivers, "QIntValidator_receivers")
	symaddr(&addr_QIntValidator_isSignalConnected, "QIntValidator_isSignalConnected")
	symaddr(&addr_QIntValidator_tr2, "QIntValidator_tr2")
	symaddr(&addr_QIntValidator_tr3, "QIntValidator_tr3")
	symaddr(&addr_QIntValidator_connect3, "QIntValidator_connect3")
	symaddr(&addr_QIntValidator_connect4, "QIntValidator_connect4")
	symaddr(&addr_QIntValidator_metaObject_default, "QIntValidator_metaObject_default")
	symaddr(&addr_QIntValidator_validate_default, "QIntValidator_validate_default")
	symaddr(&addr_QIntValidator_fixup_default, "QIntValidator_fixup_default")
	symaddr(&addr_QIntValidator_event_default, "QIntValidator_event_default")
	symaddr(&addr_QIntValidator_eventFilter_default, "QIntValidator_eventFilter_default")
	symaddr(&addr_QIntValidator_timerEvent_default, "QIntValidator_timerEvent_default")
	symaddr(&addr_QIntValidator_childEvent_default, "QIntValidator_childEvent_default")
	symaddr(&addr_QIntValidator_customEvent_default, "QIntValidator_customEvent_default")
	symaddr(&addr_QIntValidator_connectNotify_default, "QIntValidator_connectNotify_default")
	symaddr(&addr_QIntValidator_disconnectNotify_default, "QIntValidator_disconnectNotify_default")
	reg(&pg_QIntValidator_connect_bottomChanged, "QIntValidator_connect_bottomChanged")
	reg(&pg_QIntValidator_connect_topChanged, "QIntValidator_connect_topChanged")
	reg(&pg_QIntValidator_connect_changed, "QIntValidator_connect_changed")
	reg(&pg_QIntValidator_connect_destroyed, "QIntValidator_connect_destroyed")
	reg(&pg_QIntValidator_connect_destroyed2, "QIntValidator_connect_destroyed2")
	reg(&pg_QIntValidator_connect_objectNameChanged, "QIntValidator_connect_objectNameChanged")
	reg(&pg_QIntValidator__disconnect, "QIntValidator__disconnect")
	reg(&pg_QIntValidator__set_override, "QIntValidator__set_override")
	reg(&pg_QIntValidator_delete, "QIntValidator_delete")
}

type QIntValidator struct {
	h uintptr
	*QValidator
}

func bindQIntValidator(h uintptr) *QIntValidator {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QIntValidator_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QIntValidator{h: h, QValidator: &QValidator{h: __c[0], QObject: &QObject{h: __c[1]}}}
}

func newownedQIntValidator(h uintptr) *QIntValidator { return bindQIntValidator(h) }

func (this *QIntValidator) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QIntValidator) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQIntValidator(h unsafe.Pointer) *QIntValidator { return bindQIntValidator(uintptr(h)) }
func UnsafeNewQIntValidatorOwned(h unsafe.Pointer) *QIntValidator { return newownedQIntValidator(uintptr(h)) }

func (this *QIntValidator) UnsafeRelease() {
	this.h = 0
	if this.QValidator != nil {
		this.QValidator.UnsafeRelease()
	}
}

func NewQIntValidator() *QIntValidator {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QIntValidator_new, uintptr(__goid))
	__o := newownedQIntValidator(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QIntValidator_virtIdx, QIntValidator_enable)
	return __o
}

func NewQIntValidator2(bottom int, top int) *QIntValidator {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QIntValidator_new2, uintptr(__goid), uintptr(bottom), uintptr(top))
	__o := newownedQIntValidator(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QIntValidator_virtIdx, QIntValidator_enable)
	return __o
}

func NewQIntValidator3(parent *QObject) *QIntValidator {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QIntValidator_new3, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQIntValidator(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QIntValidator_virtIdx, QIntValidator_enable)
	return __o
}

func NewQIntValidator4(bottom int, top int, parent *QObject) *QIntValidator {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QIntValidator_new4, uintptr(__goid), uintptr(bottom), uintptr(top), uintptr(parent.UnsafePointer()))
	__o := newownedQIntValidator(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QIntValidator_virtIdx, QIntValidator_enable)
	return __o
}

func QIntValidator_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QIntValidator_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QIntValidator) SetBottom(arg1 int) {
	purego.SyscallN(addr_QIntValidator_setBottom, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QIntValidator) SetTop(arg1 int) {
	purego.SyscallN(addr_QIntValidator_setTop, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QIntValidator) SetRange(bottom int, top int) {
	purego.SyscallN(addr_QIntValidator_setRange, this.h, uintptr(bottom), uintptr(top))
	runtime.KeepAlive(this)
}

func (this *QIntValidator) Bottom() int {
	__r0, _, _ := purego.SyscallN(addr_QIntValidator_bottom, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QIntValidator) Top() int {
	__r0, _, _ := purego.SyscallN(addr_QIntValidator_top, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QIntValidator) BottomChanged(bottom int) {
	purego.SyscallN(addr_QIntValidator_bottomChanged, this.h, uintptr(bottom))
	runtime.KeepAlive(this)
}

func (this *QIntValidator) TopChanged(top int) {
	purego.SyscallN(addr_QIntValidator_topChanged, this.h, uintptr(top))
	runtime.KeepAlive(this)
}

func QIntValidator_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QIntValidator_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QIntValidator_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QIntValidator_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QIntValidator) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QIntValidator_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QIntValidator) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QIntValidator_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QIntValidator) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QIntValidator_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QIntValidator) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QIntValidator_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QIntValidator_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QIntValidator_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QIntValidator_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QIntValidator_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QIntValidator_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QIntValidator_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QIntValidator_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QIntValidator_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QIntValidator) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QIntValidator_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QIntValidator) ValidateDefault(arg1 *QString, arg2 int) (QValidator__State, int) {
	arg2_io := int32(arg2)
	__r0, _, _ := purego.SyscallN(addr_QIntValidator_validate_default, this.h, uintptr(arg1.UnsafePointer()), uintptr(unsafe.Pointer(&arg2_io)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return QValidator__State(__r), int(arg2_io)
}

func (this *QIntValidator) FixupDefault(input *QString) {
	purego.SyscallN(addr_QIntValidator_fixup_default, this.h, uintptr(input.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(input)
}

func (this *QIntValidator) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QIntValidator_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QIntValidator) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QIntValidator_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QIntValidator) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QIntValidator_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QIntValidator) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QIntValidator_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QIntValidator) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QIntValidator_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QIntValidator) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QIntValidator_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QIntValidator) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QIntValidator_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QIntValidator) OnBottomChanged(fn func(int)) int64 {
	id := registerCallback(func(bottom int32) { fn(int(bottom)) })
	pg_QIntValidator_connect_bottomChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QIntValidator) OnTopChanged(fn func(int)) int64 {
	id := registerCallback(func(top int32) { fn(int(top)) })
	pg_QIntValidator_connect_topChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QIntValidator) OnChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QIntValidator_connect_changed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QIntValidator) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QIntValidator_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QIntValidator) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QIntValidator_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QIntValidator) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QIntValidator_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QIntValidator) Disconnect(handle int64) {
	pg_QIntValidator__disconnect(handle)
	unregisterCallback(handle)
}

func QIntValidator_virtIdx(slot int64) int32 {
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

func QIntValidator_enable(h uintptr, idx int32) { pg_QIntValidator__set_override(h, idx) }

func (this *QIntValidator) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QIntValidator) OnValidate(fn func(*QString, *int32) QValidator__State) {
	registerVirt(this.UnsafeGoID(), 6784179835071232555, func(arg1 uintptr, arg2 uintptr) int32 { return int32(fn(UnsafeNewQString(unsafe.Pointer(arg1)), (*int32)(unsafe.Pointer(arg2)))) }) // slot "validate"
}

func (this *QIntValidator) OnFixup(fn func(*QString)) {
	registerVirt(this.UnsafeGoID(), -4855206404746966715, func(input uintptr) { fn(UnsafeNewQString(unsafe.Pointer(input))) }) // slot "fixup"
}

func (this *QIntValidator) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QIntValidator) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QIntValidator) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QIntValidator) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QIntValidator) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QIntValidator) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QIntValidator) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QIntValidator) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QIntValidator_delete(this.h)
	}
	this.UnsafeRelease()
}
