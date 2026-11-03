package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QStyledItemDelegate_chain uintptr
	addr_QStyledItemDelegate_new uintptr
	addr_QStyledItemDelegate_new2 uintptr
	addr_QStyledItemDelegate_tr uintptr
	addr_QStyledItemDelegate_itemEditorFactory uintptr
	addr_QStyledItemDelegate_setItemEditorFactory uintptr
	addr_QStyledItemDelegate_displayText uintptr
	addr_QStyledItemDelegate_connect uintptr
	addr_QStyledItemDelegate_connect2 uintptr
	addr_QStyledItemDelegate_sender uintptr
	addr_QStyledItemDelegate_senderSignalIndex uintptr
	addr_QStyledItemDelegate_receivers uintptr
	addr_QStyledItemDelegate_isSignalConnected uintptr
	addr_QStyledItemDelegate_tr2 uintptr
	addr_QStyledItemDelegate_tr3 uintptr
	addr_QStyledItemDelegate_connect3 uintptr
	addr_QStyledItemDelegate_connect4 uintptr
	addr_QStyledItemDelegate_displayText_default uintptr
	addr_QStyledItemDelegate_metaObject_default uintptr
	addr_QStyledItemDelegate_paint_default uintptr
	addr_QStyledItemDelegate_sizeHint_default uintptr
	addr_QStyledItemDelegate_createEditor_default uintptr
	addr_QStyledItemDelegate_setEditorData_default uintptr
	addr_QStyledItemDelegate_setModelData_default uintptr
	addr_QStyledItemDelegate_updateEditorGeometry_default uintptr
	addr_QStyledItemDelegate_initStyleOption_default uintptr
	addr_QStyledItemDelegate_eventFilter_default uintptr
	addr_QStyledItemDelegate_editorEvent_default uintptr
	addr_QStyledItemDelegate_destroyEditor_default uintptr
	addr_QStyledItemDelegate_helpEvent_default uintptr
	addr_QStyledItemDelegate_paintingRoles_default uintptr
	addr_QStyledItemDelegate_event_default uintptr
	addr_QStyledItemDelegate_timerEvent_default uintptr
	addr_QStyledItemDelegate_childEvent_default uintptr
	addr_QStyledItemDelegate_customEvent_default uintptr
	addr_QStyledItemDelegate_connectNotify_default uintptr
	addr_QStyledItemDelegate_disconnectNotify_default uintptr
	pg_QStyledItemDelegate_connect_commitData func(uintptr, int64)
	pg_QStyledItemDelegate_connect_closeEditor func(uintptr, int64)
	pg_QStyledItemDelegate_connect_closeEditor2 func(uintptr, int64)
	pg_QStyledItemDelegate_connect_sizeHintChanged func(uintptr, int64)
	pg_QStyledItemDelegate_connect_destroyed func(uintptr, int64)
	pg_QStyledItemDelegate_connect_destroyed2 func(uintptr, int64)
	pg_QStyledItemDelegate_connect_objectNameChanged func(uintptr, int64)
	pg_QStyledItemDelegate__disconnect func(int64)
	pg_QStyledItemDelegate__set_override func(uintptr, int32)
	pg_QStyledItemDelegate_delete func(uintptr)
)

func init() {
	symaddr(&addr_QStyledItemDelegate_chain, "QStyledItemDelegate_chain")
	symaddr(&addr_QStyledItemDelegate_new, "QStyledItemDelegate_new")
	symaddr(&addr_QStyledItemDelegate_new2, "QStyledItemDelegate_new2")
	symaddr(&addr_QStyledItemDelegate_tr, "QStyledItemDelegate_tr")
	symaddr(&addr_QStyledItemDelegate_itemEditorFactory, "QStyledItemDelegate_itemEditorFactory")
	symaddr(&addr_QStyledItemDelegate_setItemEditorFactory, "QStyledItemDelegate_setItemEditorFactory")
	symaddr(&addr_QStyledItemDelegate_displayText, "QStyledItemDelegate_displayText")
	symaddr(&addr_QStyledItemDelegate_connect, "QStyledItemDelegate_connect")
	symaddr(&addr_QStyledItemDelegate_connect2, "QStyledItemDelegate_connect2")
	symaddr(&addr_QStyledItemDelegate_sender, "QStyledItemDelegate_sender")
	symaddr(&addr_QStyledItemDelegate_senderSignalIndex, "QStyledItemDelegate_senderSignalIndex")
	symaddr(&addr_QStyledItemDelegate_receivers, "QStyledItemDelegate_receivers")
	symaddr(&addr_QStyledItemDelegate_isSignalConnected, "QStyledItemDelegate_isSignalConnected")
	symaddr(&addr_QStyledItemDelegate_tr2, "QStyledItemDelegate_tr2")
	symaddr(&addr_QStyledItemDelegate_tr3, "QStyledItemDelegate_tr3")
	symaddr(&addr_QStyledItemDelegate_connect3, "QStyledItemDelegate_connect3")
	symaddr(&addr_QStyledItemDelegate_connect4, "QStyledItemDelegate_connect4")
	symaddr(&addr_QStyledItemDelegate_displayText_default, "QStyledItemDelegate_displayText_default")
	symaddr(&addr_QStyledItemDelegate_metaObject_default, "QStyledItemDelegate_metaObject_default")
	symaddr(&addr_QStyledItemDelegate_paint_default, "QStyledItemDelegate_paint_default")
	symaddr(&addr_QStyledItemDelegate_sizeHint_default, "QStyledItemDelegate_sizeHint_default")
	symaddr(&addr_QStyledItemDelegate_createEditor_default, "QStyledItemDelegate_createEditor_default")
	symaddr(&addr_QStyledItemDelegate_setEditorData_default, "QStyledItemDelegate_setEditorData_default")
	symaddr(&addr_QStyledItemDelegate_setModelData_default, "QStyledItemDelegate_setModelData_default")
	symaddr(&addr_QStyledItemDelegate_updateEditorGeometry_default, "QStyledItemDelegate_updateEditorGeometry_default")
	symaddr(&addr_QStyledItemDelegate_initStyleOption_default, "QStyledItemDelegate_initStyleOption_default")
	symaddr(&addr_QStyledItemDelegate_eventFilter_default, "QStyledItemDelegate_eventFilter_default")
	symaddr(&addr_QStyledItemDelegate_editorEvent_default, "QStyledItemDelegate_editorEvent_default")
	symaddr(&addr_QStyledItemDelegate_destroyEditor_default, "QStyledItemDelegate_destroyEditor_default")
	symaddr(&addr_QStyledItemDelegate_helpEvent_default, "QStyledItemDelegate_helpEvent_default")
	symaddr(&addr_QStyledItemDelegate_paintingRoles_default, "QStyledItemDelegate_paintingRoles_default")
	symaddr(&addr_QStyledItemDelegate_event_default, "QStyledItemDelegate_event_default")
	symaddr(&addr_QStyledItemDelegate_timerEvent_default, "QStyledItemDelegate_timerEvent_default")
	symaddr(&addr_QStyledItemDelegate_childEvent_default, "QStyledItemDelegate_childEvent_default")
	symaddr(&addr_QStyledItemDelegate_customEvent_default, "QStyledItemDelegate_customEvent_default")
	symaddr(&addr_QStyledItemDelegate_connectNotify_default, "QStyledItemDelegate_connectNotify_default")
	symaddr(&addr_QStyledItemDelegate_disconnectNotify_default, "QStyledItemDelegate_disconnectNotify_default")
	reg(&pg_QStyledItemDelegate_connect_commitData, "QStyledItemDelegate_connect_commitData")
	reg(&pg_QStyledItemDelegate_connect_closeEditor, "QStyledItemDelegate_connect_closeEditor")
	reg(&pg_QStyledItemDelegate_connect_closeEditor2, "QStyledItemDelegate_connect_closeEditor2")
	reg(&pg_QStyledItemDelegate_connect_sizeHintChanged, "QStyledItemDelegate_connect_sizeHintChanged")
	reg(&pg_QStyledItemDelegate_connect_destroyed, "QStyledItemDelegate_connect_destroyed")
	reg(&pg_QStyledItemDelegate_connect_destroyed2, "QStyledItemDelegate_connect_destroyed2")
	reg(&pg_QStyledItemDelegate_connect_objectNameChanged, "QStyledItemDelegate_connect_objectNameChanged")
	reg(&pg_QStyledItemDelegate__disconnect, "QStyledItemDelegate__disconnect")
	reg(&pg_QStyledItemDelegate__set_override, "QStyledItemDelegate__set_override")
	reg(&pg_QStyledItemDelegate_delete, "QStyledItemDelegate_delete")
}

type QStyledItemDelegate struct {
	h uintptr
	*QAbstractItemDelegate
}

func bindQStyledItemDelegate(h uintptr) *QStyledItemDelegate {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QStyledItemDelegate_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QStyledItemDelegate{h: h, QAbstractItemDelegate: &QAbstractItemDelegate{h: __c[0], QObject: &QObject{h: __c[1]}}}
}

func newownedQStyledItemDelegate(h uintptr) *QStyledItemDelegate { return bindQStyledItemDelegate(h) }

func (this *QStyledItemDelegate) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QStyledItemDelegate) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQStyledItemDelegate(h unsafe.Pointer) *QStyledItemDelegate { return bindQStyledItemDelegate(uintptr(h)) }
func UnsafeNewQStyledItemDelegateOwned(h unsafe.Pointer) *QStyledItemDelegate { return newownedQStyledItemDelegate(uintptr(h)) }

func (this *QStyledItemDelegate) UnsafeRelease() {
	this.h = 0
	if this.QAbstractItemDelegate != nil {
		this.QAbstractItemDelegate.UnsafeRelease()
	}
}

func NewQStyledItemDelegate() *QStyledItemDelegate {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QStyledItemDelegate_new, uintptr(__goid))
	__o := newownedQStyledItemDelegate(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QStyledItemDelegate_virtIdx, QStyledItemDelegate_enable)
	return __o
}

func NewQStyledItemDelegate2(parent *QObject) *QStyledItemDelegate {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QStyledItemDelegate_new2, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQStyledItemDelegate(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QStyledItemDelegate_virtIdx, QStyledItemDelegate_enable)
	return __o
}

func QStyledItemDelegate_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStyledItemDelegate_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QStyledItemDelegate) ItemEditorFactory() *QItemEditorFactory {
	__r, _, _ := purego.SyscallN(addr_QStyledItemDelegate_itemEditorFactory, this.h)
	runtime.KeepAlive(this)
	return bindQItemEditorFactory(__r)
}

func (this *QStyledItemDelegate) SetItemEditorFactory(factory *QItemEditorFactory) {
	purego.SyscallN(addr_QStyledItemDelegate_setItemEditorFactory, this.h, uintptr(factory.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(factory)
}

func (this *QStyledItemDelegate) DisplayText(value *QVariant, locale *QLocale) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStyledItemDelegate_displayText, this.h, uintptr(value.UnsafePointer()), uintptr(locale.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(value)
	runtime.KeepAlive(locale)
	return pgStr(__r, __n)
}

func QStyledItemDelegate_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QStyledItemDelegate_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QStyledItemDelegate_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QStyledItemDelegate_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QStyledItemDelegate) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QStyledItemDelegate_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QStyledItemDelegate) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QStyledItemDelegate_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QStyledItemDelegate) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QStyledItemDelegate_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QStyledItemDelegate) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QStyledItemDelegate_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QStyledItemDelegate_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStyledItemDelegate_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QStyledItemDelegate_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStyledItemDelegate_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QStyledItemDelegate_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QStyledItemDelegate_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QStyledItemDelegate_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QStyledItemDelegate_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QStyledItemDelegate) DisplayTextDefault(value *QVariant, locale *QLocale) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStyledItemDelegate_displayText_default, this.h, uintptr(value.UnsafePointer()), uintptr(locale.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(value)
	runtime.KeepAlive(locale)
	return pgStr(__r, __n)
}

func (this *QStyledItemDelegate) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QStyledItemDelegate_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QStyledItemDelegate) PaintDefault(painter *QPainter, option *QStyleOptionViewItem, index *QModelIndex) {
	purego.SyscallN(addr_QStyledItemDelegate_paint_default, this.h, uintptr(painter.UnsafePointer()), uintptr(option.UnsafePointer()), uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
	runtime.KeepAlive(option)
	runtime.KeepAlive(index)
}

func (this *QStyledItemDelegate) SizeHintDefault(option *QStyleOptionViewItem, index *QModelIndex) *QSize {
	__r, _, _ := purego.SyscallN(addr_QStyledItemDelegate_sizeHint_default, this.h, uintptr(option.UnsafePointer()), uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(option)
	runtime.KeepAlive(index)
	return newownedQSize(__r)
}

func (this *QStyledItemDelegate) CreateEditorDefault(parent *QWidget, option *QStyleOptionViewItem, index *QModelIndex) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QStyledItemDelegate_createEditor_default, this.h, uintptr(parent.UnsafePointer()), uintptr(option.UnsafePointer()), uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	runtime.KeepAlive(option)
	runtime.KeepAlive(index)
	return bindQWidget(__r)
}

func (this *QStyledItemDelegate) SetEditorDataDefault(editor *QWidget, index *QModelIndex) {
	purego.SyscallN(addr_QStyledItemDelegate_setEditorData_default, this.h, uintptr(editor.UnsafePointer()), uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(editor)
	runtime.KeepAlive(index)
}

func (this *QStyledItemDelegate) SetModelDataDefault(editor *QWidget, model *QAbstractItemModel, index *QModelIndex) {
	purego.SyscallN(addr_QStyledItemDelegate_setModelData_default, this.h, uintptr(editor.UnsafePointer()), uintptr(model.UnsafePointer()), uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(editor)
	runtime.KeepAlive(model)
	runtime.KeepAlive(index)
}

func (this *QStyledItemDelegate) UpdateEditorGeometryDefault(editor *QWidget, option *QStyleOptionViewItem, index *QModelIndex) {
	purego.SyscallN(addr_QStyledItemDelegate_updateEditorGeometry_default, this.h, uintptr(editor.UnsafePointer()), uintptr(option.UnsafePointer()), uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(editor)
	runtime.KeepAlive(option)
	runtime.KeepAlive(index)
}

func (this *QStyledItemDelegate) InitStyleOptionDefault(option *QStyleOptionViewItem, index *QModelIndex) {
	purego.SyscallN(addr_QStyledItemDelegate_initStyleOption_default, this.h, uintptr(option.UnsafePointer()), uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(option)
	runtime.KeepAlive(index)
}

func (this *QStyledItemDelegate) EventFilterDefault(object *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QStyledItemDelegate_eventFilter_default, this.h, uintptr(object.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(object)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QStyledItemDelegate) EditorEventDefault(event *QEvent, model *QAbstractItemModel, option *QStyleOptionViewItem, index *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QStyledItemDelegate_editorEvent_default, this.h, uintptr(event.UnsafePointer()), uintptr(model.UnsafePointer()), uintptr(option.UnsafePointer()), uintptr(index.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	runtime.KeepAlive(model)
	runtime.KeepAlive(option)
	runtime.KeepAlive(index)
	return __r != 0
}

func (this *QStyledItemDelegate) DestroyEditorDefault(editor *QWidget, index *QModelIndex) {
	purego.SyscallN(addr_QStyledItemDelegate_destroyEditor_default, this.h, uintptr(editor.UnsafePointer()), uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(editor)
	runtime.KeepAlive(index)
}

func (this *QStyledItemDelegate) HelpEventDefault(event *QHelpEvent, view *QAbstractItemView, option *QStyleOptionViewItem, index *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QStyledItemDelegate_helpEvent_default, this.h, uintptr(event.UnsafePointer()), uintptr(view.UnsafePointer()), uintptr(option.UnsafePointer()), uintptr(index.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	runtime.KeepAlive(view)
	runtime.KeepAlive(option)
	runtime.KeepAlive(index)
	return __r != 0
}

func (this *QStyledItemDelegate) PaintingRolesDefault() []int {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStyledItemDelegate_paintingRoles_default, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]int, __n)
	__sl := unsafe.Slice((*int32)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = int(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QStyledItemDelegate) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QStyledItemDelegate_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QStyledItemDelegate) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QStyledItemDelegate_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QStyledItemDelegate) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QStyledItemDelegate_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QStyledItemDelegate) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QStyledItemDelegate_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QStyledItemDelegate) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QStyledItemDelegate_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QStyledItemDelegate) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QStyledItemDelegate_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QStyledItemDelegate) OnCommitData(fn func(*QWidget)) int64 {
	id := registerCallback(func(editor uintptr) { fn(UnsafeNewQWidget(unsafe.Pointer(editor))) })
	pg_QStyledItemDelegate_connect_commitData(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStyledItemDelegate) OnCloseEditor(fn func(*QWidget)) int64 {
	id := registerCallback(func(editor uintptr) { fn(UnsafeNewQWidget(unsafe.Pointer(editor))) })
	pg_QStyledItemDelegate_connect_closeEditor(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStyledItemDelegate) OnCloseEditor2(fn func(*QWidget, QAbstractItemDelegate__EndEditHint)) int64 {
	id := registerCallback(func(editor uintptr, hint int32) { fn(UnsafeNewQWidget(unsafe.Pointer(editor)), QAbstractItemDelegate__EndEditHint(hint)) })
	pg_QStyledItemDelegate_connect_closeEditor2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStyledItemDelegate) OnSizeHintChanged(fn func(*QModelIndex)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(arg1))) })
	pg_QStyledItemDelegate_connect_sizeHintChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStyledItemDelegate) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QStyledItemDelegate_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStyledItemDelegate) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QStyledItemDelegate_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStyledItemDelegate) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QStyledItemDelegate_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStyledItemDelegate) Disconnect(handle int64) {
	pg_QStyledItemDelegate__disconnect(handle)
	unregisterCallback(handle)
}

func QStyledItemDelegate_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case 6083584261053556819:
		return 1
	case 729690261457684101:
		return 2
	case -2118980437705941456:
		return 3
	case -858237687273351194:
		return 4
	case -4740900545199884600:
		return 5
	case -1005009386014896055:
		return 6
	case -6300405565848239506:
		return 7
	case -488693769771308531:
		return 8
	case 5217623514630858297:
		return 9
	case -2261279149879782144:
		return 10
	case -7699250917160078568:
		return 11
	case 1696901855134526294:
		return 12
	case -4420854320663935586:
		return 13
	case 1237236456312824991:
		return 14
	case -5409927581812374498:
		return 15
	case -1185315367526857833:
		return 16
	case 3769868229123868474:
		return 17
	case 7379029982170503008:
		return 18
	case -568369772982215962:
		return 19
	}
	return -1
}

func QStyledItemDelegate_enable(h uintptr, idx int32) { pg_QStyledItemDelegate__set_override(h, idx) }

func (this *QStyledItemDelegate) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QStyledItemDelegate) OnPaint(fn func(*QPainter, *QStyleOptionViewItem, *QModelIndex)) {
	registerVirt(this.UnsafeGoID(), 6083584261053556819, func(painter uintptr, option uintptr, index uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter)), UnsafeNewQStyleOptionViewItem(unsafe.Pointer(option)), UnsafeNewQModelIndex(unsafe.Pointer(index))) }) // slot "paint"
}

func (this *QStyledItemDelegate) OnSizeHint(fn func(*QStyleOptionViewItem, *QModelIndex) *QSize) {
	registerVirt(this.UnsafeGoID(), 729690261457684101, func(option uintptr, index uintptr) uintptr { return uintptr(fn(UnsafeNewQStyleOptionViewItem(unsafe.Pointer(option)), UnsafeNewQModelIndex(unsafe.Pointer(index))).UnsafePointer()) }) // slot "sizeHint"
}

func (this *QStyledItemDelegate) OnCreateEditor(fn func(*QWidget, *QStyleOptionViewItem, *QModelIndex) *QWidget) {
	registerVirt(this.UnsafeGoID(), -2118980437705941456, func(parent uintptr, option uintptr, index uintptr) uintptr { return uintptr(fn(UnsafeNewQWidget(unsafe.Pointer(parent)), UnsafeNewQStyleOptionViewItem(unsafe.Pointer(option)), UnsafeNewQModelIndex(unsafe.Pointer(index))).UnsafePointer()) }) // slot "createEditor"
}

func (this *QStyledItemDelegate) OnSetEditorData(fn func(*QWidget, *QModelIndex)) {
	registerVirt(this.UnsafeGoID(), -858237687273351194, func(editor uintptr, index uintptr) { fn(UnsafeNewQWidget(unsafe.Pointer(editor)), UnsafeNewQModelIndex(unsafe.Pointer(index))) }) // slot "setEditorData"
}

func (this *QStyledItemDelegate) OnSetModelData(fn func(*QWidget, *QAbstractItemModel, *QModelIndex)) {
	registerVirt(this.UnsafeGoID(), -4740900545199884600, func(editor uintptr, model uintptr, index uintptr) { fn(UnsafeNewQWidget(unsafe.Pointer(editor)), UnsafeNewQAbstractItemModel(unsafe.Pointer(model)), UnsafeNewQModelIndex(unsafe.Pointer(index))) }) // slot "setModelData"
}

func (this *QStyledItemDelegate) OnUpdateEditorGeometry(fn func(*QWidget, *QStyleOptionViewItem, *QModelIndex)) {
	registerVirt(this.UnsafeGoID(), -1005009386014896055, func(editor uintptr, option uintptr, index uintptr) { fn(UnsafeNewQWidget(unsafe.Pointer(editor)), UnsafeNewQStyleOptionViewItem(unsafe.Pointer(option)), UnsafeNewQModelIndex(unsafe.Pointer(index))) }) // slot "updateEditorGeometry"
}

func (this *QStyledItemDelegate) OnDisplayText(fn func(*QVariant, *QLocale) string) {
	registerVirt(this.UnsafeGoID(), -6300405565848239506, func(value uintptr, locale uintptr) (uintptr, int64) {
		__s := fn(UnsafeNewQVariant(unsafe.Pointer(value)), UnsafeNewQLocale(unsafe.Pointer(locale)))
		return pgCString(__s), int64(len(__s))
	}) // slot "displayText"
}

func (this *QStyledItemDelegate) OnInitStyleOption(fn func(*QStyleOptionViewItem, *QModelIndex)) {
	registerVirt(this.UnsafeGoID(), -488693769771308531, func(option uintptr, index uintptr) { fn(UnsafeNewQStyleOptionViewItem(unsafe.Pointer(option)), UnsafeNewQModelIndex(unsafe.Pointer(index))) }) // slot "initStyleOption"
}

func (this *QStyledItemDelegate) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(object uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(object)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QStyledItemDelegate) OnEditorEvent(fn func(*QEvent, *QAbstractItemModel, *QStyleOptionViewItem, *QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), -2261279149879782144, func(event uintptr, model uintptr, option uintptr, index uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)), UnsafeNewQAbstractItemModel(unsafe.Pointer(model)), UnsafeNewQStyleOptionViewItem(unsafe.Pointer(option)), UnsafeNewQModelIndex(unsafe.Pointer(index)))) }) // slot "editorEvent"
}

func (this *QStyledItemDelegate) OnDestroyEditor(fn func(*QWidget, *QModelIndex)) {
	registerVirt(this.UnsafeGoID(), -7699250917160078568, func(editor uintptr, index uintptr) { fn(UnsafeNewQWidget(unsafe.Pointer(editor)), UnsafeNewQModelIndex(unsafe.Pointer(index))) }) // slot "destroyEditor"
}

func (this *QStyledItemDelegate) OnHelpEvent(fn func(*QHelpEvent, *QAbstractItemView, *QStyleOptionViewItem, *QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), 1696901855134526294, func(event uintptr, view uintptr, option uintptr, index uintptr) int32 { return b2i(fn(UnsafeNewQHelpEvent(unsafe.Pointer(event)), UnsafeNewQAbstractItemView(unsafe.Pointer(view)), UnsafeNewQStyleOptionViewItem(unsafe.Pointer(option)), UnsafeNewQModelIndex(unsafe.Pointer(index)))) }) // slot "helpEvent"
}

func (this *QStyledItemDelegate) OnPaintingRoles(fn func() []int) {
	registerVirt(this.UnsafeGoID(), -4420854320663935586, func() (uintptr, int64) {
		__rs := fn()
		if len(__rs) == 0 {
			return 0, 0
		}
		__a := pgMalloc(uintptr(len(__rs)) * unsafe.Sizeof(int32(0)))
		__sl := unsafe.Slice((*int32)(unsafe.Pointer(__a)), len(__rs))
		for __i, __e := range __rs {
			__sl[__i] = int32(__e)
		}
		return __a, int64(len(__rs))
	}) // slot "paintingRoles"
}

func (this *QStyledItemDelegate) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QStyledItemDelegate) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QStyledItemDelegate) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QStyledItemDelegate) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QStyledItemDelegate) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QStyledItemDelegate) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QStyledItemDelegate) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QStyledItemDelegate_delete(this.h)
	}
	this.UnsafeRelease()
}
