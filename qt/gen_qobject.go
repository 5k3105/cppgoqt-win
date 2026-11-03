package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QObject_new uintptr
	addr_QObject_new2 uintptr
	addr_QObject_metaObject uintptr
	addr_QObject_tr uintptr
	addr_QObject_event uintptr
	addr_QObject_eventFilter uintptr
	addr_QObject_objectName uintptr
	addr_QObject_setObjectName uintptr
	addr_QObject_isWidgetType uintptr
	addr_QObject_isWindowType uintptr
	addr_QObject_isQuickItemType uintptr
	addr_QObject_isQmlExposed uintptr
	addr_QObject_signalsBlocked uintptr
	addr_QObject_blockSignals uintptr
	addr_QObject_thread uintptr
	addr_QObject_moveToThread uintptr
	addr_QObject_startTimer uintptr
	addr_QObject_startTimer2 uintptr
	addr_QObject_killTimer uintptr
	addr_QObject_killTimer2 uintptr
	addr_QObject_children uintptr
	addr_QObject_setParent uintptr
	addr_QObject_installEventFilter uintptr
	addr_QObject_removeEventFilter uintptr
	addr_QObject_connect uintptr
	addr_QObject_connect2 uintptr
	addr_QObject_connect3 uintptr
	addr_QObject_dumpObjectTree uintptr
	addr_QObject_dumpObjectInfo uintptr
	addr_QObject_setProperty uintptr
	addr_QObject_property uintptr
	addr_QObject_dynamicPropertyNames uintptr
	addr_QObject_bindingStorage uintptr
	addr_QObject_bindingStorage2 uintptr
	addr_QObject_destroyed uintptr
	addr_QObject_parent uintptr
	addr_QObject_inherits uintptr
	addr_QObject_deleteLater uintptr
	addr_QObject_sender uintptr
	addr_QObject_senderSignalIndex uintptr
	addr_QObject_receivers uintptr
	addr_QObject_isSignalConnected uintptr
	addr_QObject_tr2 uintptr
	addr_QObject_tr3 uintptr
	addr_QObject_startTimer3 uintptr
	addr_QObject_startTimer4 uintptr
	addr_QObject_connect4 uintptr
	addr_QObject_connect5 uintptr
	addr_QObject_connect6 uintptr
	addr_QObject_destroyed2 uintptr
	addr_QObject_metaObject_default uintptr
	addr_QObject_event_default uintptr
	addr_QObject_eventFilter_default uintptr
	addr_QObject_timerEvent_default uintptr
	addr_QObject_childEvent_default uintptr
	addr_QObject_customEvent_default uintptr
	addr_QObject_connectNotify_default uintptr
	addr_QObject_disconnectNotify_default uintptr
	pg_QObject_connect_destroyed func(uintptr, int64)
	pg_QObject_connect_destroyed2 func(uintptr, int64)
	pg_QObject_connect_objectNameChanged func(uintptr, int64)
	pg_QObject__disconnect func(int64)
	pg_QObject__set_override func(uintptr, int32)
	pg_QObject_delete func(uintptr)
)

func init() {
	symaddr(&addr_QObject_new, "QObject_new")
	symaddr(&addr_QObject_new2, "QObject_new2")
	symaddr(&addr_QObject_metaObject, "QObject_metaObject")
	symaddr(&addr_QObject_tr, "QObject_tr")
	symaddr(&addr_QObject_event, "QObject_event")
	symaddr(&addr_QObject_eventFilter, "QObject_eventFilter")
	symaddr(&addr_QObject_objectName, "QObject_objectName")
	symaddr(&addr_QObject_setObjectName, "QObject_setObjectName")
	symaddr(&addr_QObject_isWidgetType, "QObject_isWidgetType")
	symaddr(&addr_QObject_isWindowType, "QObject_isWindowType")
	symaddr(&addr_QObject_isQuickItemType, "QObject_isQuickItemType")
	symaddr(&addr_QObject_isQmlExposed, "QObject_isQmlExposed")
	symaddr(&addr_QObject_signalsBlocked, "QObject_signalsBlocked")
	symaddr(&addr_QObject_blockSignals, "QObject_blockSignals")
	symaddr(&addr_QObject_thread, "QObject_thread")
	symaddr(&addr_QObject_moveToThread, "QObject_moveToThread")
	symaddr(&addr_QObject_startTimer, "QObject_startTimer")
	symaddr(&addr_QObject_startTimer2, "QObject_startTimer2")
	symaddr(&addr_QObject_killTimer, "QObject_killTimer")
	symaddr(&addr_QObject_killTimer2, "QObject_killTimer2")
	symaddr(&addr_QObject_children, "QObject_children")
	symaddr(&addr_QObject_setParent, "QObject_setParent")
	symaddr(&addr_QObject_installEventFilter, "QObject_installEventFilter")
	symaddr(&addr_QObject_removeEventFilter, "QObject_removeEventFilter")
	symaddr(&addr_QObject_connect, "QObject_connect")
	symaddr(&addr_QObject_connect2, "QObject_connect2")
	symaddr(&addr_QObject_connect3, "QObject_connect3")
	symaddr(&addr_QObject_dumpObjectTree, "QObject_dumpObjectTree")
	symaddr(&addr_QObject_dumpObjectInfo, "QObject_dumpObjectInfo")
	symaddr(&addr_QObject_setProperty, "QObject_setProperty")
	symaddr(&addr_QObject_property, "QObject_property")
	symaddr(&addr_QObject_dynamicPropertyNames, "QObject_dynamicPropertyNames")
	symaddr(&addr_QObject_bindingStorage, "QObject_bindingStorage")
	symaddr(&addr_QObject_bindingStorage2, "QObject_bindingStorage2")
	symaddr(&addr_QObject_destroyed, "QObject_destroyed")
	symaddr(&addr_QObject_parent, "QObject_parent")
	symaddr(&addr_QObject_inherits, "QObject_inherits")
	symaddr(&addr_QObject_deleteLater, "QObject_deleteLater")
	symaddr(&addr_QObject_sender, "QObject_sender")
	symaddr(&addr_QObject_senderSignalIndex, "QObject_senderSignalIndex")
	symaddr(&addr_QObject_receivers, "QObject_receivers")
	symaddr(&addr_QObject_isSignalConnected, "QObject_isSignalConnected")
	symaddr(&addr_QObject_tr2, "QObject_tr2")
	symaddr(&addr_QObject_tr3, "QObject_tr3")
	symaddr(&addr_QObject_startTimer3, "QObject_startTimer3")
	symaddr(&addr_QObject_startTimer4, "QObject_startTimer4")
	symaddr(&addr_QObject_connect4, "QObject_connect4")
	symaddr(&addr_QObject_connect5, "QObject_connect5")
	symaddr(&addr_QObject_connect6, "QObject_connect6")
	symaddr(&addr_QObject_destroyed2, "QObject_destroyed2")
	symaddr(&addr_QObject_metaObject_default, "QObject_metaObject_default")
	symaddr(&addr_QObject_event_default, "QObject_event_default")
	symaddr(&addr_QObject_eventFilter_default, "QObject_eventFilter_default")
	symaddr(&addr_QObject_timerEvent_default, "QObject_timerEvent_default")
	symaddr(&addr_QObject_childEvent_default, "QObject_childEvent_default")
	symaddr(&addr_QObject_customEvent_default, "QObject_customEvent_default")
	symaddr(&addr_QObject_connectNotify_default, "QObject_connectNotify_default")
	symaddr(&addr_QObject_disconnectNotify_default, "QObject_disconnectNotify_default")
	reg(&pg_QObject_connect_destroyed, "QObject_connect_destroyed")
	reg(&pg_QObject_connect_destroyed2, "QObject_connect_destroyed2")
	reg(&pg_QObject_connect_objectNameChanged, "QObject_connect_objectNameChanged")
	reg(&pg_QObject__disconnect, "QObject__disconnect")
	reg(&pg_QObject__set_override, "QObject__set_override")
	reg(&pg_QObject_delete, "QObject_delete")
}

type QObject struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QObject) UnsafeGoID() int64     { return this.vid }
func (this *QObject) UnsafeSetGoID(id int64) { this.vid = id }

func bindQObject(h uintptr) *QObject {
	if h == 0 {
		return nil
	}
	return &QObject{h: h}
}

func newownedQObject(h uintptr) *QObject { return bindQObject(h) }

func (this *QObject) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QObject) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQObject(h unsafe.Pointer) *QObject { return bindQObject(uintptr(h)) }
func UnsafeNewQObjectOwned(h unsafe.Pointer) *QObject { return newownedQObject(uintptr(h)) }

func (this *QObject) UnsafeRelease() { this.h = 0 }

func NewQObject() *QObject {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QObject_new, uintptr(__goid))
	__o := newownedQObject(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QObject_virtIdx, QObject_enable)
	return __o
}

func NewQObject2(parent *QObject) *QObject {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QObject_new2, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQObject(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QObject_virtIdx, QObject_enable)
	return __o
}

func (this *QObject) MetaObject() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QObject_metaObject, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func QObject_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QObject_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QObject) Event(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QObject_event, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QObject) EventFilter(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QObject_eventFilter, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QObject) ObjectName() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QObject_objectName, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QObject) SetObjectName(name string) {
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	purego.SyscallN(addr_QObject_setObjectName, this.h, name_p, uintptr(len(name)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(name)
}

func (this *QObject) IsWidgetType() bool {
	__r0, _, _ := purego.SyscallN(addr_QObject_isWidgetType, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QObject) IsWindowType() bool {
	__r0, _, _ := purego.SyscallN(addr_QObject_isWindowType, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QObject) IsQuickItemType() bool {
	__r0, _, _ := purego.SyscallN(addr_QObject_isQuickItemType, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QObject) IsQmlExposed() bool {
	__r0, _, _ := purego.SyscallN(addr_QObject_isQmlExposed, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QObject) SignalsBlocked() bool {
	__r0, _, _ := purego.SyscallN(addr_QObject_signalsBlocked, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QObject) BlockSignals(b bool) bool {
	__r0, _, _ := purego.SyscallN(addr_QObject_blockSignals, this.h, uintptr(b2i(b)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QObject) Thread() *QThread {
	__r, _, _ := purego.SyscallN(addr_QObject_thread, this.h)
	runtime.KeepAlive(this)
	return bindQThread(__r)
}

func (this *QObject) MoveToThread(thread *QThread) bool {
	__r0, _, _ := purego.SyscallN(addr_QObject_moveToThread, this.h, uintptr(thread.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(thread)
	return __r != 0
}

func (this *QObject) StartTimer(interval int) int {
	__r0, _, _ := purego.SyscallN(addr_QObject_startTimer, this.h, uintptr(interval))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QObject) StartTimerWithTime(time int64) int {
	__r0, _, _ := purego.SyscallN(addr_QObject_startTimer2, this.h, uintptr(time))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QObject) KillTimer(id int) {
	purego.SyscallN(addr_QObject_killTimer, this.h, uintptr(id))
	runtime.KeepAlive(this)
}

func (this *QObject) KillTimerWithId(id TimerId) {
	purego.SyscallN(addr_QObject_killTimer2, this.h, uintptr(id))
	runtime.KeepAlive(this)
}

func (this *QObject) Children() []*QObject {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QObject_children, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QObject, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = bindQObject(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QObject) SetParent(parent *QObject) {
	purego.SyscallN(addr_QObject_setParent, this.h, uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
}

func (this *QObject) InstallEventFilter(filterObj *QObject) {
	purego.SyscallN(addr_QObject_installEventFilter, this.h, uintptr(filterObj.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(filterObj)
}

func (this *QObject) RemoveEventFilter(obj *QObject) {
	purego.SyscallN(addr_QObject_removeEventFilter, this.h, uintptr(obj.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(obj)
}

func QObject_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QObject_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QObject_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QObject_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QObject) Connect(sender *QObject, signal string, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QObject_connect3, this.h, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func (this *QObject) DumpObjectTree() {
	purego.SyscallN(addr_QObject_dumpObjectTree, this.h)
	runtime.KeepAlive(this)
}

func (this *QObject) DumpObjectInfo() {
	purego.SyscallN(addr_QObject_dumpObjectInfo, this.h)
	runtime.KeepAlive(this)
}

func (this *QObject) SetProperty(name string, value *QVariant) bool {
	name_b := append([]byte(name), 0)
	__r0, _, _ := purego.SyscallN(addr_QObject_setProperty, this.h, uintptr(unsafe.Pointer(&name_b[0])), uintptr(value.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(name_b)
	runtime.KeepAlive(value)
	return __r != 0
}

func (this *QObject) Property(name string) *QVariant {
	name_b := append([]byte(name), 0)
	__r, _, _ := purego.SyscallN(addr_QObject_property, this.h, uintptr(unsafe.Pointer(&name_b[0])))
	runtime.KeepAlive(this)
	runtime.KeepAlive(name_b)
	return newownedQVariant(__r)
}

func (this *QObject) DynamicPropertyNames() []*QByteArray {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QObject_dynamicPropertyNames, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QByteArray, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQByteArray(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QObject) BindingStorage() *QBindingStorage {
	__r, _, _ := purego.SyscallN(addr_QObject_bindingStorage, this.h)
	runtime.KeepAlive(this)
	return bindQBindingStorage(__r)
}

func (this *QObject) BindingStorage2() *QBindingStorage {
	__r, _, _ := purego.SyscallN(addr_QObject_bindingStorage2, this.h)
	runtime.KeepAlive(this)
	return bindQBindingStorage(__r)
}

func (this *QObject) Destroyed() {
	purego.SyscallN(addr_QObject_destroyed, this.h)
	runtime.KeepAlive(this)
}

func (this *QObject) Parent() *QObject {
	__r, _, _ := purego.SyscallN(addr_QObject_parent, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QObject) Inherits(classname string) bool {
	classname_b := append([]byte(classname), 0)
	__r0, _, _ := purego.SyscallN(addr_QObject_inherits, this.h, uintptr(unsafe.Pointer(&classname_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(classname_b)
	return __r != 0
}

func (this *QObject) DeleteLater() {
	purego.SyscallN(addr_QObject_deleteLater, this.h)
	runtime.KeepAlive(this)
}

func (this *QObject) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QObject_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QObject) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QObject_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QObject) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QObject_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QObject) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QObject_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QObject_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QObject_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QObject_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QObject_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func (this *QObject) StartTimer2(interval int, timerType TimerType) int {
	__r0, _, _ := purego.SyscallN(addr_QObject_startTimer3, this.h, uintptr(interval), uintptr(timerType))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QObject) StartTimer3(time int64, timerType TimerType) int {
	__r0, _, _ := purego.SyscallN(addr_QObject_startTimer4, this.h, uintptr(time), uintptr(timerType))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func QObject_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QObject_connect4, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QObject_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QObject_connect5, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QObject) Connect2(sender *QObject, signal string, member string, type_ ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QObject_connect6, this.h, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(unsafe.Pointer(&member_b[0])), uintptr(type_))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func (this *QObject) DestroyedWithQObject(arg1 *QObject) {
	purego.SyscallN(addr_QObject_destroyed2, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QObject) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QObject_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QObject) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QObject_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QObject) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QObject_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QObject) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QObject_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QObject) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QObject_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QObject) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QObject_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QObject) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QObject_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QObject) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QObject_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QObject) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QObject_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QObject) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QObject_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QObject) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QObject_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QObject) Disconnect(handle int64) {
	pg_QObject__disconnect(handle)
	unregisterCallback(handle)
}

func QObject_virtIdx(slot int64) int32 {
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

func QObject_enable(h uintptr, idx int32) { pg_QObject__set_override(h, idx) }

func (this *QObject) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QObject) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QObject) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QObject) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QObject) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QObject) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QObject) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QObject) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QObject) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QObject_delete(this.h)
	}
	this.UnsafeRelease()
}
