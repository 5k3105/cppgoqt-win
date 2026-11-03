package qt

import (
	"unsafe"
	"math"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QDoubleValidator_chain uintptr
	addr_QDoubleValidator_new uintptr
	pg_QDoubleValidator_new2 func(int64, float64, float64, int32) uintptr
	addr_QDoubleValidator_new3 uintptr
	pg_QDoubleValidator_new4 func(int64, float64, float64, int32, uintptr) uintptr
	addr_QDoubleValidator_tr uintptr
	pg_QDoubleValidator_setRange func(uintptr, float64, float64, int32)
	pg_QDoubleValidator_setRange2 func(uintptr, float64, float64)
	pg_QDoubleValidator_setBottom func(uintptr, float64)
	pg_QDoubleValidator_setTop func(uintptr, float64)
	addr_QDoubleValidator_setDecimals uintptr
	addr_QDoubleValidator_setNotation uintptr
	pg_QDoubleValidator_bottom func(uintptr) float64
	pg_QDoubleValidator_top func(uintptr) float64
	addr_QDoubleValidator_decimals uintptr
	addr_QDoubleValidator_notation uintptr
	pg_QDoubleValidator_bottomChanged func(uintptr, float64)
	pg_QDoubleValidator_topChanged func(uintptr, float64)
	addr_QDoubleValidator_decimalsChanged uintptr
	addr_QDoubleValidator_notationChanged uintptr
	addr_QDoubleValidator_connect uintptr
	addr_QDoubleValidator_connect2 uintptr
	addr_QDoubleValidator_sender uintptr
	addr_QDoubleValidator_senderSignalIndex uintptr
	addr_QDoubleValidator_receivers uintptr
	addr_QDoubleValidator_isSignalConnected uintptr
	addr_QDoubleValidator_tr2 uintptr
	addr_QDoubleValidator_tr3 uintptr
	addr_QDoubleValidator_connect3 uintptr
	addr_QDoubleValidator_connect4 uintptr
	addr_QDoubleValidator_metaObject_default uintptr
	addr_QDoubleValidator_validate_default uintptr
	addr_QDoubleValidator_fixup_default uintptr
	addr_QDoubleValidator_event_default uintptr
	addr_QDoubleValidator_eventFilter_default uintptr
	addr_QDoubleValidator_timerEvent_default uintptr
	addr_QDoubleValidator_childEvent_default uintptr
	addr_QDoubleValidator_customEvent_default uintptr
	addr_QDoubleValidator_connectNotify_default uintptr
	addr_QDoubleValidator_disconnectNotify_default uintptr
	pg_QDoubleValidator_connect_bottomChanged func(uintptr, int64)
	pg_QDoubleValidator_connect_topChanged func(uintptr, int64)
	pg_QDoubleValidator_connect_decimalsChanged func(uintptr, int64)
	pg_QDoubleValidator_connect_notationChanged func(uintptr, int64)
	pg_QDoubleValidator_connect_changed func(uintptr, int64)
	pg_QDoubleValidator_connect_destroyed func(uintptr, int64)
	pg_QDoubleValidator_connect_destroyed2 func(uintptr, int64)
	pg_QDoubleValidator_connect_objectNameChanged func(uintptr, int64)
	pg_QDoubleValidator__disconnect func(int64)
	pg_QDoubleValidator__set_override func(uintptr, int32)
	pg_QDoubleValidator_delete func(uintptr)
)

func init() {
	symaddr(&addr_QDoubleValidator_chain, "QDoubleValidator_chain")
	symaddr(&addr_QDoubleValidator_new, "QDoubleValidator_new")
	reg(&pg_QDoubleValidator_new2, "QDoubleValidator_new2")
	symaddr(&addr_QDoubleValidator_new3, "QDoubleValidator_new3")
	reg(&pg_QDoubleValidator_new4, "QDoubleValidator_new4")
	symaddr(&addr_QDoubleValidator_tr, "QDoubleValidator_tr")
	reg(&pg_QDoubleValidator_setRange, "QDoubleValidator_setRange")
	reg(&pg_QDoubleValidator_setRange2, "QDoubleValidator_setRange2")
	reg(&pg_QDoubleValidator_setBottom, "QDoubleValidator_setBottom")
	reg(&pg_QDoubleValidator_setTop, "QDoubleValidator_setTop")
	symaddr(&addr_QDoubleValidator_setDecimals, "QDoubleValidator_setDecimals")
	symaddr(&addr_QDoubleValidator_setNotation, "QDoubleValidator_setNotation")
	reg(&pg_QDoubleValidator_bottom, "QDoubleValidator_bottom")
	reg(&pg_QDoubleValidator_top, "QDoubleValidator_top")
	symaddr(&addr_QDoubleValidator_decimals, "QDoubleValidator_decimals")
	symaddr(&addr_QDoubleValidator_notation, "QDoubleValidator_notation")
	reg(&pg_QDoubleValidator_bottomChanged, "QDoubleValidator_bottomChanged")
	reg(&pg_QDoubleValidator_topChanged, "QDoubleValidator_topChanged")
	symaddr(&addr_QDoubleValidator_decimalsChanged, "QDoubleValidator_decimalsChanged")
	symaddr(&addr_QDoubleValidator_notationChanged, "QDoubleValidator_notationChanged")
	symaddr(&addr_QDoubleValidator_connect, "QDoubleValidator_connect")
	symaddr(&addr_QDoubleValidator_connect2, "QDoubleValidator_connect2")
	symaddr(&addr_QDoubleValidator_sender, "QDoubleValidator_sender")
	symaddr(&addr_QDoubleValidator_senderSignalIndex, "QDoubleValidator_senderSignalIndex")
	symaddr(&addr_QDoubleValidator_receivers, "QDoubleValidator_receivers")
	symaddr(&addr_QDoubleValidator_isSignalConnected, "QDoubleValidator_isSignalConnected")
	symaddr(&addr_QDoubleValidator_tr2, "QDoubleValidator_tr2")
	symaddr(&addr_QDoubleValidator_tr3, "QDoubleValidator_tr3")
	symaddr(&addr_QDoubleValidator_connect3, "QDoubleValidator_connect3")
	symaddr(&addr_QDoubleValidator_connect4, "QDoubleValidator_connect4")
	symaddr(&addr_QDoubleValidator_metaObject_default, "QDoubleValidator_metaObject_default")
	symaddr(&addr_QDoubleValidator_validate_default, "QDoubleValidator_validate_default")
	symaddr(&addr_QDoubleValidator_fixup_default, "QDoubleValidator_fixup_default")
	symaddr(&addr_QDoubleValidator_event_default, "QDoubleValidator_event_default")
	symaddr(&addr_QDoubleValidator_eventFilter_default, "QDoubleValidator_eventFilter_default")
	symaddr(&addr_QDoubleValidator_timerEvent_default, "QDoubleValidator_timerEvent_default")
	symaddr(&addr_QDoubleValidator_childEvent_default, "QDoubleValidator_childEvent_default")
	symaddr(&addr_QDoubleValidator_customEvent_default, "QDoubleValidator_customEvent_default")
	symaddr(&addr_QDoubleValidator_connectNotify_default, "QDoubleValidator_connectNotify_default")
	symaddr(&addr_QDoubleValidator_disconnectNotify_default, "QDoubleValidator_disconnectNotify_default")
	reg(&pg_QDoubleValidator_connect_bottomChanged, "QDoubleValidator_connect_bottomChanged")
	reg(&pg_QDoubleValidator_connect_topChanged, "QDoubleValidator_connect_topChanged")
	reg(&pg_QDoubleValidator_connect_decimalsChanged, "QDoubleValidator_connect_decimalsChanged")
	reg(&pg_QDoubleValidator_connect_notationChanged, "QDoubleValidator_connect_notationChanged")
	reg(&pg_QDoubleValidator_connect_changed, "QDoubleValidator_connect_changed")
	reg(&pg_QDoubleValidator_connect_destroyed, "QDoubleValidator_connect_destroyed")
	reg(&pg_QDoubleValidator_connect_destroyed2, "QDoubleValidator_connect_destroyed2")
	reg(&pg_QDoubleValidator_connect_objectNameChanged, "QDoubleValidator_connect_objectNameChanged")
	reg(&pg_QDoubleValidator__disconnect, "QDoubleValidator__disconnect")
	reg(&pg_QDoubleValidator__set_override, "QDoubleValidator__set_override")
	reg(&pg_QDoubleValidator_delete, "QDoubleValidator_delete")
}

type QDoubleValidator struct {
	h uintptr
	*QValidator
}

func bindQDoubleValidator(h uintptr) *QDoubleValidator {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QDoubleValidator_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QDoubleValidator{h: h, QValidator: &QValidator{h: __c[0], QObject: &QObject{h: __c[1]}}}
}

func newownedQDoubleValidator(h uintptr) *QDoubleValidator { return bindQDoubleValidator(h) }

func (this *QDoubleValidator) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QDoubleValidator) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQDoubleValidator(h unsafe.Pointer) *QDoubleValidator { return bindQDoubleValidator(uintptr(h)) }
func UnsafeNewQDoubleValidatorOwned(h unsafe.Pointer) *QDoubleValidator { return newownedQDoubleValidator(uintptr(h)) }

func (this *QDoubleValidator) UnsafeRelease() {
	this.h = 0
	if this.QValidator != nil {
		this.QValidator.UnsafeRelease()
	}
}

func NewQDoubleValidator() *QDoubleValidator {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QDoubleValidator_new, uintptr(__goid))
	__o := newownedQDoubleValidator(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QDoubleValidator_virtIdx, QDoubleValidator_enable)
	return __o
}

func NewQDoubleValidator2(bottom float64, top float64, decimals int) *QDoubleValidator {
	__goid := nextVirtID()
	__o := newownedQDoubleValidator(pg_QDoubleValidator_new2(__goid, bottom, top, int32(decimals)))
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QDoubleValidator_virtIdx, QDoubleValidator_enable)
	return __o
}

func NewQDoubleValidator3(parent *QObject) *QDoubleValidator {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QDoubleValidator_new3, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQDoubleValidator(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QDoubleValidator_virtIdx, QDoubleValidator_enable)
	return __o
}

func NewQDoubleValidator4(bottom float64, top float64, decimals int, parent *QObject) *QDoubleValidator {
	__goid := nextVirtID()
	__o := newownedQDoubleValidator(pg_QDoubleValidator_new4(__goid, bottom, top, int32(decimals), uintptr(parent.UnsafePointer())))
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QDoubleValidator_virtIdx, QDoubleValidator_enable)
	return __o
}

func QDoubleValidator_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDoubleValidator_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QDoubleValidator) SetRange(bottom float64, top float64, decimals int) {
	pg_QDoubleValidator_setRange(this.h, bottom, top, int32(decimals))
	runtime.KeepAlive(this)
}

func (this *QDoubleValidator) SetRange2(bottom float64, top float64) {
	pg_QDoubleValidator_setRange2(this.h, bottom, top)
	runtime.KeepAlive(this)
}

func (this *QDoubleValidator) SetBottom(arg1 float64) {
	pg_QDoubleValidator_setBottom(this.h, arg1)
	runtime.KeepAlive(this)
}

func (this *QDoubleValidator) SetTop(arg1 float64) {
	pg_QDoubleValidator_setTop(this.h, arg1)
	runtime.KeepAlive(this)
}

func (this *QDoubleValidator) SetDecimals(arg1 int) {
	purego.SyscallN(addr_QDoubleValidator_setDecimals, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QDoubleValidator) SetNotation(arg1 QDoubleValidator__Notation) {
	purego.SyscallN(addr_QDoubleValidator_setNotation, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QDoubleValidator) Bottom() float64 {
	__r := pg_QDoubleValidator_bottom(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QDoubleValidator) Top() float64 {
	__r := pg_QDoubleValidator_top(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QDoubleValidator) Decimals() int {
	__r0, _, _ := purego.SyscallN(addr_QDoubleValidator_decimals, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QDoubleValidator) Notation() QDoubleValidator__Notation {
	__r0, _, _ := purego.SyscallN(addr_QDoubleValidator_notation, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QDoubleValidator__Notation(__r)
}

func (this *QDoubleValidator) BottomChanged(bottom float64) {
	pg_QDoubleValidator_bottomChanged(this.h, bottom)
	runtime.KeepAlive(this)
}

func (this *QDoubleValidator) TopChanged(top float64) {
	pg_QDoubleValidator_topChanged(this.h, top)
	runtime.KeepAlive(this)
}

func (this *QDoubleValidator) DecimalsChanged(decimals int) {
	purego.SyscallN(addr_QDoubleValidator_decimalsChanged, this.h, uintptr(decimals))
	runtime.KeepAlive(this)
}

func (this *QDoubleValidator) NotationChanged(notation QDoubleValidator__Notation) {
	purego.SyscallN(addr_QDoubleValidator_notationChanged, this.h, uintptr(notation))
	runtime.KeepAlive(this)
}

func QDoubleValidator_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QDoubleValidator_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QDoubleValidator_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QDoubleValidator_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QDoubleValidator) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QDoubleValidator_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QDoubleValidator) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QDoubleValidator_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QDoubleValidator) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QDoubleValidator_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QDoubleValidator) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QDoubleValidator_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QDoubleValidator_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDoubleValidator_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QDoubleValidator_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDoubleValidator_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QDoubleValidator_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QDoubleValidator_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QDoubleValidator_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QDoubleValidator_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QDoubleValidator) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QDoubleValidator_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QDoubleValidator) ValidateDefault(arg1 *QString, arg2 int) (QValidator__State, int) {
	arg2_io := int32(arg2)
	__r0, _, _ := purego.SyscallN(addr_QDoubleValidator_validate_default, this.h, uintptr(arg1.UnsafePointer()), uintptr(unsafe.Pointer(&arg2_io)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return QValidator__State(__r), int(arg2_io)
}

func (this *QDoubleValidator) FixupDefault(input *QString) {
	purego.SyscallN(addr_QDoubleValidator_fixup_default, this.h, uintptr(input.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(input)
}

func (this *QDoubleValidator) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QDoubleValidator_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QDoubleValidator) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QDoubleValidator_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QDoubleValidator) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QDoubleValidator_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDoubleValidator) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QDoubleValidator_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDoubleValidator) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QDoubleValidator_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDoubleValidator) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QDoubleValidator_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QDoubleValidator) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QDoubleValidator_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QDoubleValidator) OnBottomChanged(fn func(float64)) int64 {
	id := registerCallback(func(bottom uintptr) { fn(math.Float64frombits(uint64(bottom))) })
	pg_QDoubleValidator_connect_bottomChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDoubleValidator) OnTopChanged(fn func(float64)) int64 {
	id := registerCallback(func(top uintptr) { fn(math.Float64frombits(uint64(top))) })
	pg_QDoubleValidator_connect_topChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDoubleValidator) OnDecimalsChanged(fn func(int)) int64 {
	id := registerCallback(func(decimals int32) { fn(int(decimals)) })
	pg_QDoubleValidator_connect_decimalsChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDoubleValidator) OnNotationChanged(fn func(QDoubleValidator__Notation)) int64 {
	id := registerCallback(func(notation int32) { fn(QDoubleValidator__Notation(notation)) })
	pg_QDoubleValidator_connect_notationChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDoubleValidator) OnChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QDoubleValidator_connect_changed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDoubleValidator) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QDoubleValidator_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDoubleValidator) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QDoubleValidator_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDoubleValidator) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QDoubleValidator_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDoubleValidator) Disconnect(handle int64) {
	pg_QDoubleValidator__disconnect(handle)
	unregisterCallback(handle)
}

func QDoubleValidator_virtIdx(slot int64) int32 {
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

func QDoubleValidator_enable(h uintptr, idx int32) { pg_QDoubleValidator__set_override(h, idx) }

func (this *QDoubleValidator) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QDoubleValidator) OnValidate(fn func(*QString, *int32) QValidator__State) {
	registerVirt(this.UnsafeGoID(), 6784179835071232555, func(arg1 uintptr, arg2 uintptr) int32 { return int32(fn(UnsafeNewQString(unsafe.Pointer(arg1)), (*int32)(unsafe.Pointer(arg2)))) }) // slot "validate"
}

func (this *QDoubleValidator) OnFixup(fn func(*QString)) {
	registerVirt(this.UnsafeGoID(), -4855206404746966715, func(input uintptr) { fn(UnsafeNewQString(unsafe.Pointer(input))) }) // slot "fixup"
}

func (this *QDoubleValidator) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QDoubleValidator) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QDoubleValidator) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QDoubleValidator) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QDoubleValidator) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QDoubleValidator) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QDoubleValidator) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QDoubleValidator) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QDoubleValidator_delete(this.h)
	}
	this.UnsafeRelease()
}
