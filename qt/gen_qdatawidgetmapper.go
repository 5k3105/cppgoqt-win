package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QDataWidgetMapper_base func(uintptr) uintptr
	addr_QDataWidgetMapper_new uintptr
	addr_QDataWidgetMapper_new2 uintptr
	addr_QDataWidgetMapper_tr uintptr
	addr_QDataWidgetMapper_setModel uintptr
	addr_QDataWidgetMapper_model uintptr
	addr_QDataWidgetMapper_setItemDelegate uintptr
	addr_QDataWidgetMapper_itemDelegate uintptr
	addr_QDataWidgetMapper_setRootIndex uintptr
	addr_QDataWidgetMapper_rootIndex uintptr
	addr_QDataWidgetMapper_setOrientation uintptr
	addr_QDataWidgetMapper_orientation uintptr
	addr_QDataWidgetMapper_setSubmitPolicy uintptr
	addr_QDataWidgetMapper_submitPolicy uintptr
	addr_QDataWidgetMapper_addMapping uintptr
	addr_QDataWidgetMapper_addMapping2 uintptr
	addr_QDataWidgetMapper_removeMapping uintptr
	addr_QDataWidgetMapper_mappedSection uintptr
	addr_QDataWidgetMapper_mappedPropertyName uintptr
	addr_QDataWidgetMapper_mappedWidgetAt uintptr
	addr_QDataWidgetMapper_clearMapping uintptr
	addr_QDataWidgetMapper_currentIndex uintptr
	addr_QDataWidgetMapper_revert uintptr
	addr_QDataWidgetMapper_submit uintptr
	addr_QDataWidgetMapper_toFirst uintptr
	addr_QDataWidgetMapper_toLast uintptr
	addr_QDataWidgetMapper_toNext uintptr
	addr_QDataWidgetMapper_toPrevious uintptr
	addr_QDataWidgetMapper_setCurrentIndex uintptr
	addr_QDataWidgetMapper_setCurrentModelIndex uintptr
	addr_QDataWidgetMapper_currentIndexChanged uintptr
	addr_QDataWidgetMapper_connect uintptr
	addr_QDataWidgetMapper_connect2 uintptr
	addr_QDataWidgetMapper_sender uintptr
	addr_QDataWidgetMapper_senderSignalIndex uintptr
	addr_QDataWidgetMapper_receivers uintptr
	addr_QDataWidgetMapper_isSignalConnected uintptr
	addr_QDataWidgetMapper_tr2 uintptr
	addr_QDataWidgetMapper_tr3 uintptr
	addr_QDataWidgetMapper_connect3 uintptr
	addr_QDataWidgetMapper_connect4 uintptr
	addr_QDataWidgetMapper_setCurrentIndex_default uintptr
	addr_QDataWidgetMapper_metaObject_default uintptr
	addr_QDataWidgetMapper_event_default uintptr
	addr_QDataWidgetMapper_eventFilter_default uintptr
	addr_QDataWidgetMapper_timerEvent_default uintptr
	addr_QDataWidgetMapper_childEvent_default uintptr
	addr_QDataWidgetMapper_customEvent_default uintptr
	addr_QDataWidgetMapper_connectNotify_default uintptr
	addr_QDataWidgetMapper_disconnectNotify_default uintptr
	pg_QDataWidgetMapper_connect_currentIndexChanged func(uintptr, int64)
	pg_QDataWidgetMapper_connect_destroyed func(uintptr, int64)
	pg_QDataWidgetMapper_connect_destroyed2 func(uintptr, int64)
	pg_QDataWidgetMapper_connect_objectNameChanged func(uintptr, int64)
	pg_QDataWidgetMapper__disconnect func(int64)
	pg_QDataWidgetMapper__set_override func(uintptr, int32)
	pg_QDataWidgetMapper_delete func(uintptr)
)

func init() {
	reg(&pg_QDataWidgetMapper_base, "QDataWidgetMapper_base")
	symaddr(&addr_QDataWidgetMapper_new, "QDataWidgetMapper_new")
	symaddr(&addr_QDataWidgetMapper_new2, "QDataWidgetMapper_new2")
	symaddr(&addr_QDataWidgetMapper_tr, "QDataWidgetMapper_tr")
	symaddr(&addr_QDataWidgetMapper_setModel, "QDataWidgetMapper_setModel")
	symaddr(&addr_QDataWidgetMapper_model, "QDataWidgetMapper_model")
	symaddr(&addr_QDataWidgetMapper_setItemDelegate, "QDataWidgetMapper_setItemDelegate")
	symaddr(&addr_QDataWidgetMapper_itemDelegate, "QDataWidgetMapper_itemDelegate")
	symaddr(&addr_QDataWidgetMapper_setRootIndex, "QDataWidgetMapper_setRootIndex")
	symaddr(&addr_QDataWidgetMapper_rootIndex, "QDataWidgetMapper_rootIndex")
	symaddr(&addr_QDataWidgetMapper_setOrientation, "QDataWidgetMapper_setOrientation")
	symaddr(&addr_QDataWidgetMapper_orientation, "QDataWidgetMapper_orientation")
	symaddr(&addr_QDataWidgetMapper_setSubmitPolicy, "QDataWidgetMapper_setSubmitPolicy")
	symaddr(&addr_QDataWidgetMapper_submitPolicy, "QDataWidgetMapper_submitPolicy")
	symaddr(&addr_QDataWidgetMapper_addMapping, "QDataWidgetMapper_addMapping")
	symaddr(&addr_QDataWidgetMapper_addMapping2, "QDataWidgetMapper_addMapping2")
	symaddr(&addr_QDataWidgetMapper_removeMapping, "QDataWidgetMapper_removeMapping")
	symaddr(&addr_QDataWidgetMapper_mappedSection, "QDataWidgetMapper_mappedSection")
	symaddr(&addr_QDataWidgetMapper_mappedPropertyName, "QDataWidgetMapper_mappedPropertyName")
	symaddr(&addr_QDataWidgetMapper_mappedWidgetAt, "QDataWidgetMapper_mappedWidgetAt")
	symaddr(&addr_QDataWidgetMapper_clearMapping, "QDataWidgetMapper_clearMapping")
	symaddr(&addr_QDataWidgetMapper_currentIndex, "QDataWidgetMapper_currentIndex")
	symaddr(&addr_QDataWidgetMapper_revert, "QDataWidgetMapper_revert")
	symaddr(&addr_QDataWidgetMapper_submit, "QDataWidgetMapper_submit")
	symaddr(&addr_QDataWidgetMapper_toFirst, "QDataWidgetMapper_toFirst")
	symaddr(&addr_QDataWidgetMapper_toLast, "QDataWidgetMapper_toLast")
	symaddr(&addr_QDataWidgetMapper_toNext, "QDataWidgetMapper_toNext")
	symaddr(&addr_QDataWidgetMapper_toPrevious, "QDataWidgetMapper_toPrevious")
	symaddr(&addr_QDataWidgetMapper_setCurrentIndex, "QDataWidgetMapper_setCurrentIndex")
	symaddr(&addr_QDataWidgetMapper_setCurrentModelIndex, "QDataWidgetMapper_setCurrentModelIndex")
	symaddr(&addr_QDataWidgetMapper_currentIndexChanged, "QDataWidgetMapper_currentIndexChanged")
	symaddr(&addr_QDataWidgetMapper_connect, "QDataWidgetMapper_connect")
	symaddr(&addr_QDataWidgetMapper_connect2, "QDataWidgetMapper_connect2")
	symaddr(&addr_QDataWidgetMapper_sender, "QDataWidgetMapper_sender")
	symaddr(&addr_QDataWidgetMapper_senderSignalIndex, "QDataWidgetMapper_senderSignalIndex")
	symaddr(&addr_QDataWidgetMapper_receivers, "QDataWidgetMapper_receivers")
	symaddr(&addr_QDataWidgetMapper_isSignalConnected, "QDataWidgetMapper_isSignalConnected")
	symaddr(&addr_QDataWidgetMapper_tr2, "QDataWidgetMapper_tr2")
	symaddr(&addr_QDataWidgetMapper_tr3, "QDataWidgetMapper_tr3")
	symaddr(&addr_QDataWidgetMapper_connect3, "QDataWidgetMapper_connect3")
	symaddr(&addr_QDataWidgetMapper_connect4, "QDataWidgetMapper_connect4")
	symaddr(&addr_QDataWidgetMapper_setCurrentIndex_default, "QDataWidgetMapper_setCurrentIndex_default")
	symaddr(&addr_QDataWidgetMapper_metaObject_default, "QDataWidgetMapper_metaObject_default")
	symaddr(&addr_QDataWidgetMapper_event_default, "QDataWidgetMapper_event_default")
	symaddr(&addr_QDataWidgetMapper_eventFilter_default, "QDataWidgetMapper_eventFilter_default")
	symaddr(&addr_QDataWidgetMapper_timerEvent_default, "QDataWidgetMapper_timerEvent_default")
	symaddr(&addr_QDataWidgetMapper_childEvent_default, "QDataWidgetMapper_childEvent_default")
	symaddr(&addr_QDataWidgetMapper_customEvent_default, "QDataWidgetMapper_customEvent_default")
	symaddr(&addr_QDataWidgetMapper_connectNotify_default, "QDataWidgetMapper_connectNotify_default")
	symaddr(&addr_QDataWidgetMapper_disconnectNotify_default, "QDataWidgetMapper_disconnectNotify_default")
	reg(&pg_QDataWidgetMapper_connect_currentIndexChanged, "QDataWidgetMapper_connect_currentIndexChanged")
	reg(&pg_QDataWidgetMapper_connect_destroyed, "QDataWidgetMapper_connect_destroyed")
	reg(&pg_QDataWidgetMapper_connect_destroyed2, "QDataWidgetMapper_connect_destroyed2")
	reg(&pg_QDataWidgetMapper_connect_objectNameChanged, "QDataWidgetMapper_connect_objectNameChanged")
	reg(&pg_QDataWidgetMapper__disconnect, "QDataWidgetMapper__disconnect")
	reg(&pg_QDataWidgetMapper__set_override, "QDataWidgetMapper__set_override")
	reg(&pg_QDataWidgetMapper_delete, "QDataWidgetMapper_delete")
}

type QDataWidgetMapper struct {
	h uintptr
	*QObject
}

func bindQDataWidgetMapper(h uintptr) *QDataWidgetMapper {
	if h == 0 {
		return nil
	}
	return &QDataWidgetMapper{h: h, QObject: bindQObject(pg_QDataWidgetMapper_base(h))}
}

func newownedQDataWidgetMapper(h uintptr) *QDataWidgetMapper { return bindQDataWidgetMapper(h) }

func (this *QDataWidgetMapper) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QDataWidgetMapper) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQDataWidgetMapper(h unsafe.Pointer) *QDataWidgetMapper { return bindQDataWidgetMapper(uintptr(h)) }
func UnsafeNewQDataWidgetMapperOwned(h unsafe.Pointer) *QDataWidgetMapper { return newownedQDataWidgetMapper(uintptr(h)) }

func (this *QDataWidgetMapper) UnsafeRelease() {
	this.h = 0
	if this.QObject != nil {
		this.QObject.UnsafeRelease()
	}
}

func NewQDataWidgetMapper() *QDataWidgetMapper {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QDataWidgetMapper_new, uintptr(__goid))
	__o := newownedQDataWidgetMapper(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QDataWidgetMapper_virtIdx, QDataWidgetMapper_enable)
	return __o
}

func NewQDataWidgetMapper2(parent *QObject) *QDataWidgetMapper {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QDataWidgetMapper_new2, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQDataWidgetMapper(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QDataWidgetMapper_virtIdx, QDataWidgetMapper_enable)
	return __o
}

func QDataWidgetMapper_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDataWidgetMapper_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QDataWidgetMapper) SetModel(model *QAbstractItemModel) {
	purego.SyscallN(addr_QDataWidgetMapper_setModel, this.h, uintptr(model.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(model)
}

func (this *QDataWidgetMapper) Model() *QAbstractItemModel {
	__r, _, _ := purego.SyscallN(addr_QDataWidgetMapper_model, this.h)
	runtime.KeepAlive(this)
	return bindQAbstractItemModel(__r)
}

func (this *QDataWidgetMapper) SetItemDelegate(delegate *QAbstractItemDelegate) {
	purego.SyscallN(addr_QDataWidgetMapper_setItemDelegate, this.h, uintptr(delegate.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(delegate)
}

func (this *QDataWidgetMapper) ItemDelegate() *QAbstractItemDelegate {
	__r, _, _ := purego.SyscallN(addr_QDataWidgetMapper_itemDelegate, this.h)
	runtime.KeepAlive(this)
	return bindQAbstractItemDelegate(__r)
}

func (this *QDataWidgetMapper) SetRootIndex(index *QModelIndex) {
	purego.SyscallN(addr_QDataWidgetMapper_setRootIndex, this.h, uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
}

func (this *QDataWidgetMapper) RootIndex() *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QDataWidgetMapper_rootIndex, this.h)
	runtime.KeepAlive(this)
	return newownedQModelIndex(__r)
}

func (this *QDataWidgetMapper) SetOrientation(aOrientation Orientation) {
	purego.SyscallN(addr_QDataWidgetMapper_setOrientation, this.h, uintptr(aOrientation))
	runtime.KeepAlive(this)
}

func (this *QDataWidgetMapper) Orientation() Orientation {
	__r0, _, _ := purego.SyscallN(addr_QDataWidgetMapper_orientation, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return Orientation(__r)
}

func (this *QDataWidgetMapper) SetSubmitPolicy(policy QDataWidgetMapper__SubmitPolicy) {
	purego.SyscallN(addr_QDataWidgetMapper_setSubmitPolicy, this.h, uintptr(policy))
	runtime.KeepAlive(this)
}

func (this *QDataWidgetMapper) SubmitPolicy() QDataWidgetMapper__SubmitPolicy {
	__r0, _, _ := purego.SyscallN(addr_QDataWidgetMapper_submitPolicy, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QDataWidgetMapper__SubmitPolicy(__r)
}

func (this *QDataWidgetMapper) AddMapping(widget *QWidget, section int) {
	purego.SyscallN(addr_QDataWidgetMapper_addMapping, this.h, uintptr(widget.UnsafePointer()), uintptr(section))
	runtime.KeepAlive(this)
	runtime.KeepAlive(widget)
}

func (this *QDataWidgetMapper) AddMapping2(widget *QWidget, section int, propertyName *QByteArray) {
	purego.SyscallN(addr_QDataWidgetMapper_addMapping2, this.h, uintptr(widget.UnsafePointer()), uintptr(section), uintptr(propertyName.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(widget)
	runtime.KeepAlive(propertyName)
}

func (this *QDataWidgetMapper) RemoveMapping(widget *QWidget) {
	purego.SyscallN(addr_QDataWidgetMapper_removeMapping, this.h, uintptr(widget.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(widget)
}

func (this *QDataWidgetMapper) MappedSection(widget *QWidget) int {
	__r0, _, _ := purego.SyscallN(addr_QDataWidgetMapper_mappedSection, this.h, uintptr(widget.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(widget)
	return int(__r)
}

func (this *QDataWidgetMapper) MappedPropertyName(widget *QWidget) *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QDataWidgetMapper_mappedPropertyName, this.h, uintptr(widget.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(widget)
	return newownedQByteArray(__r)
}

func (this *QDataWidgetMapper) MappedWidgetAt(section int) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QDataWidgetMapper_mappedWidgetAt, this.h, uintptr(section))
	runtime.KeepAlive(this)
	return bindQWidget(__r)
}

func (this *QDataWidgetMapper) ClearMapping() {
	purego.SyscallN(addr_QDataWidgetMapper_clearMapping, this.h)
	runtime.KeepAlive(this)
}

func (this *QDataWidgetMapper) CurrentIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QDataWidgetMapper_currentIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QDataWidgetMapper) Revert() {
	purego.SyscallN(addr_QDataWidgetMapper_revert, this.h)
	runtime.KeepAlive(this)
}

func (this *QDataWidgetMapper) Submit() bool {
	__r0, _, _ := purego.SyscallN(addr_QDataWidgetMapper_submit, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QDataWidgetMapper) ToFirst() {
	purego.SyscallN(addr_QDataWidgetMapper_toFirst, this.h)
	runtime.KeepAlive(this)
}

func (this *QDataWidgetMapper) ToLast() {
	purego.SyscallN(addr_QDataWidgetMapper_toLast, this.h)
	runtime.KeepAlive(this)
}

func (this *QDataWidgetMapper) ToNext() {
	purego.SyscallN(addr_QDataWidgetMapper_toNext, this.h)
	runtime.KeepAlive(this)
}

func (this *QDataWidgetMapper) ToPrevious() {
	purego.SyscallN(addr_QDataWidgetMapper_toPrevious, this.h)
	runtime.KeepAlive(this)
}

func (this *QDataWidgetMapper) SetCurrentIndex(index int) {
	purego.SyscallN(addr_QDataWidgetMapper_setCurrentIndex, this.h, uintptr(index))
	runtime.KeepAlive(this)
}

func (this *QDataWidgetMapper) SetCurrentModelIndex(index *QModelIndex) {
	purego.SyscallN(addr_QDataWidgetMapper_setCurrentModelIndex, this.h, uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
}

func (this *QDataWidgetMapper) CurrentIndexChanged(index int) {
	purego.SyscallN(addr_QDataWidgetMapper_currentIndexChanged, this.h, uintptr(index))
	runtime.KeepAlive(this)
}

func QDataWidgetMapper_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QDataWidgetMapper_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QDataWidgetMapper_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QDataWidgetMapper_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QDataWidgetMapper) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QDataWidgetMapper_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QDataWidgetMapper) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QDataWidgetMapper_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QDataWidgetMapper) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QDataWidgetMapper_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QDataWidgetMapper) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QDataWidgetMapper_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QDataWidgetMapper_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDataWidgetMapper_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QDataWidgetMapper_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDataWidgetMapper_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QDataWidgetMapper_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QDataWidgetMapper_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QDataWidgetMapper_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QDataWidgetMapper_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QDataWidgetMapper) SetCurrentIndexDefault(index int) {
	purego.SyscallN(addr_QDataWidgetMapper_setCurrentIndex_default, this.h, uintptr(index))
	runtime.KeepAlive(this)
}

func (this *QDataWidgetMapper) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QDataWidgetMapper_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QDataWidgetMapper) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QDataWidgetMapper_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QDataWidgetMapper) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QDataWidgetMapper_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QDataWidgetMapper) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QDataWidgetMapper_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDataWidgetMapper) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QDataWidgetMapper_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDataWidgetMapper) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QDataWidgetMapper_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDataWidgetMapper) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QDataWidgetMapper_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QDataWidgetMapper) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QDataWidgetMapper_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QDataWidgetMapper) OnCurrentIndexChanged(fn func(int)) int64 {
	id := registerCallback(func(index int32) { fn(int(index)) })
	pg_QDataWidgetMapper_connect_currentIndexChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDataWidgetMapper) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QDataWidgetMapper_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDataWidgetMapper) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QDataWidgetMapper_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDataWidgetMapper) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QDataWidgetMapper_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDataWidgetMapper) Disconnect(handle int64) {
	pg_QDataWidgetMapper__disconnect(handle)
	unregisterCallback(handle)
}

func QDataWidgetMapper_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case -6491574503686613496:
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

func QDataWidgetMapper_enable(h uintptr, idx int32) { pg_QDataWidgetMapper__set_override(h, idx) }

func (this *QDataWidgetMapper) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QDataWidgetMapper) OnSetCurrentIndex(fn func(int)) {
	registerVirt(this.UnsafeGoID(), -6491574503686613496, func(index int32) { fn(int(index)) }) // slot "setCurrentIndex"
}

func (this *QDataWidgetMapper) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QDataWidgetMapper) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QDataWidgetMapper) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QDataWidgetMapper) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QDataWidgetMapper) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QDataWidgetMapper) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QDataWidgetMapper) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QDataWidgetMapper) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QDataWidgetMapper_delete(this.h)
	}
	this.UnsafeRelease()
}
