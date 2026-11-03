package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QAbstractItemDelegate_base func(uintptr) uintptr
	addr_QAbstractItemDelegate_tr uintptr
	addr_QAbstractItemDelegate_paint uintptr
	addr_QAbstractItemDelegate_sizeHint uintptr
	addr_QAbstractItemDelegate_createEditor uintptr
	addr_QAbstractItemDelegate_destroyEditor uintptr
	addr_QAbstractItemDelegate_setEditorData uintptr
	addr_QAbstractItemDelegate_setModelData uintptr
	addr_QAbstractItemDelegate_updateEditorGeometry uintptr
	addr_QAbstractItemDelegate_editorEvent uintptr
	addr_QAbstractItemDelegate_handleEditorEvent uintptr
	addr_QAbstractItemDelegate_helpEvent uintptr
	addr_QAbstractItemDelegate_paintingRoles uintptr
	addr_QAbstractItemDelegate_commitData uintptr
	addr_QAbstractItemDelegate_closeEditor uintptr
	addr_QAbstractItemDelegate_sizeHintChanged uintptr
	addr_QAbstractItemDelegate_connect uintptr
	addr_QAbstractItemDelegate_connect2 uintptr
	addr_QAbstractItemDelegate_sender uintptr
	addr_QAbstractItemDelegate_senderSignalIndex uintptr
	addr_QAbstractItemDelegate_receivers uintptr
	addr_QAbstractItemDelegate_isSignalConnected uintptr
	addr_QAbstractItemDelegate_tr2 uintptr
	addr_QAbstractItemDelegate_tr3 uintptr
	addr_QAbstractItemDelegate_closeEditor2 uintptr
	addr_QAbstractItemDelegate_connect3 uintptr
	addr_QAbstractItemDelegate_connect4 uintptr
	addr_QAbstractItemDelegate_createEditor_default uintptr
	addr_QAbstractItemDelegate_destroyEditor_default uintptr
	addr_QAbstractItemDelegate_setEditorData_default uintptr
	addr_QAbstractItemDelegate_setModelData_default uintptr
	addr_QAbstractItemDelegate_updateEditorGeometry_default uintptr
	addr_QAbstractItemDelegate_editorEvent_default uintptr
	addr_QAbstractItemDelegate_helpEvent_default uintptr
	addr_QAbstractItemDelegate_paintingRoles_default uintptr
	addr_QAbstractItemDelegate_metaObject_default uintptr
	addr_QAbstractItemDelegate_event_default uintptr
	addr_QAbstractItemDelegate_eventFilter_default uintptr
	addr_QAbstractItemDelegate_timerEvent_default uintptr
	addr_QAbstractItemDelegate_childEvent_default uintptr
	addr_QAbstractItemDelegate_customEvent_default uintptr
	addr_QAbstractItemDelegate_connectNotify_default uintptr
	addr_QAbstractItemDelegate_disconnectNotify_default uintptr
	pg_QAbstractItemDelegate_connect_commitData func(uintptr, int64)
	pg_QAbstractItemDelegate_connect_closeEditor func(uintptr, int64)
	pg_QAbstractItemDelegate_connect_closeEditor2 func(uintptr, int64)
	pg_QAbstractItemDelegate_connect_sizeHintChanged func(uintptr, int64)
	pg_QAbstractItemDelegate_connect_destroyed func(uintptr, int64)
	pg_QAbstractItemDelegate_connect_destroyed2 func(uintptr, int64)
	pg_QAbstractItemDelegate_connect_objectNameChanged func(uintptr, int64)
	pg_QAbstractItemDelegate__disconnect func(int64)
	pg_QAbstractItemDelegate__set_override func(uintptr, int32)
	pg_QAbstractItemDelegate__Go_new func(int64, uintptr) uintptr
	pg_QAbstractItemDelegate_delete func(uintptr)
)

func init() {
	reg(&pg_QAbstractItemDelegate_base, "QAbstractItemDelegate_base")
	symaddr(&addr_QAbstractItemDelegate_tr, "QAbstractItemDelegate_tr")
	symaddr(&addr_QAbstractItemDelegate_paint, "QAbstractItemDelegate_paint")
	symaddr(&addr_QAbstractItemDelegate_sizeHint, "QAbstractItemDelegate_sizeHint")
	symaddr(&addr_QAbstractItemDelegate_createEditor, "QAbstractItemDelegate_createEditor")
	symaddr(&addr_QAbstractItemDelegate_destroyEditor, "QAbstractItemDelegate_destroyEditor")
	symaddr(&addr_QAbstractItemDelegate_setEditorData, "QAbstractItemDelegate_setEditorData")
	symaddr(&addr_QAbstractItemDelegate_setModelData, "QAbstractItemDelegate_setModelData")
	symaddr(&addr_QAbstractItemDelegate_updateEditorGeometry, "QAbstractItemDelegate_updateEditorGeometry")
	symaddr(&addr_QAbstractItemDelegate_editorEvent, "QAbstractItemDelegate_editorEvent")
	symaddr(&addr_QAbstractItemDelegate_handleEditorEvent, "QAbstractItemDelegate_handleEditorEvent")
	symaddr(&addr_QAbstractItemDelegate_helpEvent, "QAbstractItemDelegate_helpEvent")
	symaddr(&addr_QAbstractItemDelegate_paintingRoles, "QAbstractItemDelegate_paintingRoles")
	symaddr(&addr_QAbstractItemDelegate_commitData, "QAbstractItemDelegate_commitData")
	symaddr(&addr_QAbstractItemDelegate_closeEditor, "QAbstractItemDelegate_closeEditor")
	symaddr(&addr_QAbstractItemDelegate_sizeHintChanged, "QAbstractItemDelegate_sizeHintChanged")
	symaddr(&addr_QAbstractItemDelegate_connect, "QAbstractItemDelegate_connect")
	symaddr(&addr_QAbstractItemDelegate_connect2, "QAbstractItemDelegate_connect2")
	symaddr(&addr_QAbstractItemDelegate_sender, "QAbstractItemDelegate_sender")
	symaddr(&addr_QAbstractItemDelegate_senderSignalIndex, "QAbstractItemDelegate_senderSignalIndex")
	symaddr(&addr_QAbstractItemDelegate_receivers, "QAbstractItemDelegate_receivers")
	symaddr(&addr_QAbstractItemDelegate_isSignalConnected, "QAbstractItemDelegate_isSignalConnected")
	symaddr(&addr_QAbstractItemDelegate_tr2, "QAbstractItemDelegate_tr2")
	symaddr(&addr_QAbstractItemDelegate_tr3, "QAbstractItemDelegate_tr3")
	symaddr(&addr_QAbstractItemDelegate_closeEditor2, "QAbstractItemDelegate_closeEditor2")
	symaddr(&addr_QAbstractItemDelegate_connect3, "QAbstractItemDelegate_connect3")
	symaddr(&addr_QAbstractItemDelegate_connect4, "QAbstractItemDelegate_connect4")
	symaddr(&addr_QAbstractItemDelegate_createEditor_default, "QAbstractItemDelegate_createEditor_default")
	symaddr(&addr_QAbstractItemDelegate_destroyEditor_default, "QAbstractItemDelegate_destroyEditor_default")
	symaddr(&addr_QAbstractItemDelegate_setEditorData_default, "QAbstractItemDelegate_setEditorData_default")
	symaddr(&addr_QAbstractItemDelegate_setModelData_default, "QAbstractItemDelegate_setModelData_default")
	symaddr(&addr_QAbstractItemDelegate_updateEditorGeometry_default, "QAbstractItemDelegate_updateEditorGeometry_default")
	symaddr(&addr_QAbstractItemDelegate_editorEvent_default, "QAbstractItemDelegate_editorEvent_default")
	symaddr(&addr_QAbstractItemDelegate_helpEvent_default, "QAbstractItemDelegate_helpEvent_default")
	symaddr(&addr_QAbstractItemDelegate_paintingRoles_default, "QAbstractItemDelegate_paintingRoles_default")
	symaddr(&addr_QAbstractItemDelegate_metaObject_default, "QAbstractItemDelegate_metaObject_default")
	symaddr(&addr_QAbstractItemDelegate_event_default, "QAbstractItemDelegate_event_default")
	symaddr(&addr_QAbstractItemDelegate_eventFilter_default, "QAbstractItemDelegate_eventFilter_default")
	symaddr(&addr_QAbstractItemDelegate_timerEvent_default, "QAbstractItemDelegate_timerEvent_default")
	symaddr(&addr_QAbstractItemDelegate_childEvent_default, "QAbstractItemDelegate_childEvent_default")
	symaddr(&addr_QAbstractItemDelegate_customEvent_default, "QAbstractItemDelegate_customEvent_default")
	symaddr(&addr_QAbstractItemDelegate_connectNotify_default, "QAbstractItemDelegate_connectNotify_default")
	symaddr(&addr_QAbstractItemDelegate_disconnectNotify_default, "QAbstractItemDelegate_disconnectNotify_default")
	reg(&pg_QAbstractItemDelegate_connect_commitData, "QAbstractItemDelegate_connect_commitData")
	reg(&pg_QAbstractItemDelegate_connect_closeEditor, "QAbstractItemDelegate_connect_closeEditor")
	reg(&pg_QAbstractItemDelegate_connect_closeEditor2, "QAbstractItemDelegate_connect_closeEditor2")
	reg(&pg_QAbstractItemDelegate_connect_sizeHintChanged, "QAbstractItemDelegate_connect_sizeHintChanged")
	reg(&pg_QAbstractItemDelegate_connect_destroyed, "QAbstractItemDelegate_connect_destroyed")
	reg(&pg_QAbstractItemDelegate_connect_destroyed2, "QAbstractItemDelegate_connect_destroyed2")
	reg(&pg_QAbstractItemDelegate_connect_objectNameChanged, "QAbstractItemDelegate_connect_objectNameChanged")
	reg(&pg_QAbstractItemDelegate__disconnect, "QAbstractItemDelegate__disconnect")
	reg(&pg_QAbstractItemDelegate__set_override, "QAbstractItemDelegate__set_override")
	reg(&pg_QAbstractItemDelegate__Go_new, "QAbstractItemDelegate__Go_new")
	reg(&pg_QAbstractItemDelegate_delete, "QAbstractItemDelegate_delete")
}

type QAbstractItemDelegate struct {
	h uintptr
	*QObject
}

func bindQAbstractItemDelegate(h uintptr) *QAbstractItemDelegate {
	if h == 0 {
		return nil
	}
	return &QAbstractItemDelegate{h: h, QObject: bindQObject(pg_QAbstractItemDelegate_base(h))}
}

func newownedQAbstractItemDelegate(h uintptr) *QAbstractItemDelegate { return bindQAbstractItemDelegate(h) }

func (this *QAbstractItemDelegate) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QAbstractItemDelegate) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQAbstractItemDelegate(h unsafe.Pointer) *QAbstractItemDelegate { return bindQAbstractItemDelegate(uintptr(h)) }
func UnsafeNewQAbstractItemDelegateOwned(h unsafe.Pointer) *QAbstractItemDelegate { return newownedQAbstractItemDelegate(uintptr(h)) }

func (this *QAbstractItemDelegate) UnsafeRelease() {
	this.h = 0
	if this.QObject != nil {
		this.QObject.UnsafeRelease()
	}
}

func QAbstractItemDelegate_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAbstractItemDelegate_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QAbstractItemDelegate) Paint(painter *QPainter, option *QStyleOptionViewItem, index *QModelIndex) {
	purego.SyscallN(addr_QAbstractItemDelegate_paint, this.h, uintptr(painter.UnsafePointer()), uintptr(option.UnsafePointer()), uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
	runtime.KeepAlive(option)
	runtime.KeepAlive(index)
}

func (this *QAbstractItemDelegate) SizeHint(option *QStyleOptionViewItem, index *QModelIndex) *QSize {
	__r, _, _ := purego.SyscallN(addr_QAbstractItemDelegate_sizeHint, this.h, uintptr(option.UnsafePointer()), uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(option)
	runtime.KeepAlive(index)
	return newownedQSize(__r)
}

func (this *QAbstractItemDelegate) CreateEditor(parent *QWidget, option *QStyleOptionViewItem, index *QModelIndex) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QAbstractItemDelegate_createEditor, this.h, uintptr(parent.UnsafePointer()), uintptr(option.UnsafePointer()), uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	runtime.KeepAlive(option)
	runtime.KeepAlive(index)
	return bindQWidget(__r)
}

func (this *QAbstractItemDelegate) DestroyEditor(editor *QWidget, index *QModelIndex) {
	purego.SyscallN(addr_QAbstractItemDelegate_destroyEditor, this.h, uintptr(editor.UnsafePointer()), uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(editor)
	runtime.KeepAlive(index)
}

func (this *QAbstractItemDelegate) SetEditorData(editor *QWidget, index *QModelIndex) {
	purego.SyscallN(addr_QAbstractItemDelegate_setEditorData, this.h, uintptr(editor.UnsafePointer()), uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(editor)
	runtime.KeepAlive(index)
}

func (this *QAbstractItemDelegate) SetModelData(editor *QWidget, model *QAbstractItemModel, index *QModelIndex) {
	purego.SyscallN(addr_QAbstractItemDelegate_setModelData, this.h, uintptr(editor.UnsafePointer()), uintptr(model.UnsafePointer()), uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(editor)
	runtime.KeepAlive(model)
	runtime.KeepAlive(index)
}

func (this *QAbstractItemDelegate) UpdateEditorGeometry(editor *QWidget, option *QStyleOptionViewItem, index *QModelIndex) {
	purego.SyscallN(addr_QAbstractItemDelegate_updateEditorGeometry, this.h, uintptr(editor.UnsafePointer()), uintptr(option.UnsafePointer()), uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(editor)
	runtime.KeepAlive(option)
	runtime.KeepAlive(index)
}

func (this *QAbstractItemDelegate) EditorEvent(event *QEvent, model *QAbstractItemModel, option *QStyleOptionViewItem, index *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemDelegate_editorEvent, this.h, uintptr(event.UnsafePointer()), uintptr(model.UnsafePointer()), uintptr(option.UnsafePointer()), uintptr(index.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	runtime.KeepAlive(model)
	runtime.KeepAlive(option)
	runtime.KeepAlive(index)
	return __r != 0
}

func (this *QAbstractItemDelegate) HandleEditorEvent(object *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemDelegate_handleEditorEvent, this.h, uintptr(object.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(object)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QAbstractItemDelegate) HelpEvent(event *QHelpEvent, view *QAbstractItemView, option *QStyleOptionViewItem, index *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemDelegate_helpEvent, this.h, uintptr(event.UnsafePointer()), uintptr(view.UnsafePointer()), uintptr(option.UnsafePointer()), uintptr(index.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	runtime.KeepAlive(view)
	runtime.KeepAlive(option)
	runtime.KeepAlive(index)
	return __r != 0
}

func (this *QAbstractItemDelegate) PaintingRoles() []int {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAbstractItemDelegate_paintingRoles, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]int, __n)
	__sl := unsafe.Slice((*int32)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = int(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QAbstractItemDelegate) CommitData(editor *QWidget) {
	purego.SyscallN(addr_QAbstractItemDelegate_commitData, this.h, uintptr(editor.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(editor)
}

func (this *QAbstractItemDelegate) CloseEditor(editor *QWidget) {
	purego.SyscallN(addr_QAbstractItemDelegate_closeEditor, this.h, uintptr(editor.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(editor)
}

func (this *QAbstractItemDelegate) SizeHintChanged(arg1 *QModelIndex) {
	purego.SyscallN(addr_QAbstractItemDelegate_sizeHintChanged, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func QAbstractItemDelegate_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QAbstractItemDelegate_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QAbstractItemDelegate_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QAbstractItemDelegate_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QAbstractItemDelegate) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QAbstractItemDelegate_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QAbstractItemDelegate) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemDelegate_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAbstractItemDelegate) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemDelegate_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QAbstractItemDelegate) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemDelegate_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QAbstractItemDelegate_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAbstractItemDelegate_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QAbstractItemDelegate_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAbstractItemDelegate_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func (this *QAbstractItemDelegate) CloseEditor2(editor *QWidget, hint QAbstractItemDelegate__EndEditHint) {
	purego.SyscallN(addr_QAbstractItemDelegate_closeEditor2, this.h, uintptr(editor.UnsafePointer()), uintptr(hint))
	runtime.KeepAlive(this)
	runtime.KeepAlive(editor)
}

func QAbstractItemDelegate_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QAbstractItemDelegate_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QAbstractItemDelegate_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QAbstractItemDelegate_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QAbstractItemDelegate) CreateEditorDefault(parent *QWidget, option *QStyleOptionViewItem, index *QModelIndex) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QAbstractItemDelegate_createEditor_default, this.h, uintptr(parent.UnsafePointer()), uintptr(option.UnsafePointer()), uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	runtime.KeepAlive(option)
	runtime.KeepAlive(index)
	return bindQWidget(__r)
}

func (this *QAbstractItemDelegate) DestroyEditorDefault(editor *QWidget, index *QModelIndex) {
	purego.SyscallN(addr_QAbstractItemDelegate_destroyEditor_default, this.h, uintptr(editor.UnsafePointer()), uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(editor)
	runtime.KeepAlive(index)
}

func (this *QAbstractItemDelegate) SetEditorDataDefault(editor *QWidget, index *QModelIndex) {
	purego.SyscallN(addr_QAbstractItemDelegate_setEditorData_default, this.h, uintptr(editor.UnsafePointer()), uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(editor)
	runtime.KeepAlive(index)
}

func (this *QAbstractItemDelegate) SetModelDataDefault(editor *QWidget, model *QAbstractItemModel, index *QModelIndex) {
	purego.SyscallN(addr_QAbstractItemDelegate_setModelData_default, this.h, uintptr(editor.UnsafePointer()), uintptr(model.UnsafePointer()), uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(editor)
	runtime.KeepAlive(model)
	runtime.KeepAlive(index)
}

func (this *QAbstractItemDelegate) UpdateEditorGeometryDefault(editor *QWidget, option *QStyleOptionViewItem, index *QModelIndex) {
	purego.SyscallN(addr_QAbstractItemDelegate_updateEditorGeometry_default, this.h, uintptr(editor.UnsafePointer()), uintptr(option.UnsafePointer()), uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(editor)
	runtime.KeepAlive(option)
	runtime.KeepAlive(index)
}

func (this *QAbstractItemDelegate) EditorEventDefault(event *QEvent, model *QAbstractItemModel, option *QStyleOptionViewItem, index *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemDelegate_editorEvent_default, this.h, uintptr(event.UnsafePointer()), uintptr(model.UnsafePointer()), uintptr(option.UnsafePointer()), uintptr(index.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	runtime.KeepAlive(model)
	runtime.KeepAlive(option)
	runtime.KeepAlive(index)
	return __r != 0
}

func (this *QAbstractItemDelegate) HelpEventDefault(event *QHelpEvent, view *QAbstractItemView, option *QStyleOptionViewItem, index *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemDelegate_helpEvent_default, this.h, uintptr(event.UnsafePointer()), uintptr(view.UnsafePointer()), uintptr(option.UnsafePointer()), uintptr(index.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	runtime.KeepAlive(view)
	runtime.KeepAlive(option)
	runtime.KeepAlive(index)
	return __r != 0
}

func (this *QAbstractItemDelegate) PaintingRolesDefault() []int {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAbstractItemDelegate_paintingRoles_default, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]int, __n)
	__sl := unsafe.Slice((*int32)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = int(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QAbstractItemDelegate) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QAbstractItemDelegate_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QAbstractItemDelegate) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemDelegate_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QAbstractItemDelegate) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemDelegate_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QAbstractItemDelegate) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QAbstractItemDelegate_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractItemDelegate) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QAbstractItemDelegate_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractItemDelegate) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QAbstractItemDelegate_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractItemDelegate) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QAbstractItemDelegate_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QAbstractItemDelegate) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QAbstractItemDelegate_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QAbstractItemDelegate) OnCommitData(fn func(*QWidget)) int64 {
	id := registerCallback(func(editor uintptr) { fn(UnsafeNewQWidget(unsafe.Pointer(editor))) })
	pg_QAbstractItemDelegate_connect_commitData(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractItemDelegate) OnCloseEditor(fn func(*QWidget)) int64 {
	id := registerCallback(func(editor uintptr) { fn(UnsafeNewQWidget(unsafe.Pointer(editor))) })
	pg_QAbstractItemDelegate_connect_closeEditor(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractItemDelegate) OnCloseEditor2(fn func(*QWidget, QAbstractItemDelegate__EndEditHint)) int64 {
	id := registerCallback(func(editor uintptr, hint int32) { fn(UnsafeNewQWidget(unsafe.Pointer(editor)), QAbstractItemDelegate__EndEditHint(hint)) })
	pg_QAbstractItemDelegate_connect_closeEditor2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractItemDelegate) OnSizeHintChanged(fn func(*QModelIndex)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(arg1))) })
	pg_QAbstractItemDelegate_connect_sizeHintChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractItemDelegate) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QAbstractItemDelegate_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractItemDelegate) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QAbstractItemDelegate_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractItemDelegate) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QAbstractItemDelegate_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractItemDelegate) Disconnect(handle int64) {
	pg_QAbstractItemDelegate__disconnect(handle)
	unregisterCallback(handle)
}

func QAbstractItemDelegate_virtIdx(slot int64) int32 {
	switch slot {
	case 6083584261053556819:
		return 0
	case 729690261457684101:
		return 1
	case 1607084412591819503:
		return 2
	case -2118980437705941456:
		return 3
	case -7699250917160078568:
		return 4
	case -858237687273351194:
		return 5
	case -4740900545199884600:
		return 6
	case -1005009386014896055:
		return 7
	case -2261279149879782144:
		return 8
	case 1696901855134526294:
		return 9
	case -4420854320663935586:
		return 10
	case 1237236456312824991:
		return 11
	case 5217623514630858297:
		return 12
	case -5409927581812374498:
		return 13
	case -1185315367526857833:
		return 14
	case 3769868229123868474:
		return 15
	case 7379029982170503008:
		return 16
	case -568369772982215962:
		return 17
	}
	return -1
}

func QAbstractItemDelegate_enable(h uintptr, idx int32) { pg_QAbstractItemDelegate__set_override(h, idx) }

func NewQAbstractItemDelegate() *QAbstractItemDelegate {
	id := nextVirtID()
	o := newownedQAbstractItemDelegate(pg_QAbstractItemDelegate__Go_new(id, 0))
	o.UnsafeSetGoID(id)
	registerInstance(id, o.h, QAbstractItemDelegate_virtIdx, QAbstractItemDelegate_enable)
	return o
}

func NewQAbstractItemDelegate2(parent *QObject) *QAbstractItemDelegate {
	id := nextVirtID()
	o := newownedQAbstractItemDelegate(pg_QAbstractItemDelegate__Go_new(id, uintptr(parent.UnsafePointer())))
	o.UnsafeSetGoID(id)
	registerInstance(id, o.h, QAbstractItemDelegate_virtIdx, QAbstractItemDelegate_enable)
	return o
}

func (this *QAbstractItemDelegate) OnPaint(fn func(*QPainter, *QStyleOptionViewItem, *QModelIndex)) {
	registerVirt(this.UnsafeGoID(), 6083584261053556819, func(painter uintptr, option uintptr, index uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter)), UnsafeNewQStyleOptionViewItem(unsafe.Pointer(option)), UnsafeNewQModelIndex(unsafe.Pointer(index))) }) // slot "paint"
}

func (this *QAbstractItemDelegate) OnSizeHint(fn func(*QStyleOptionViewItem, *QModelIndex) *QSize) {
	registerVirt(this.UnsafeGoID(), 729690261457684101, func(option uintptr, index uintptr) uintptr { return uintptr(fn(UnsafeNewQStyleOptionViewItem(unsafe.Pointer(option)), UnsafeNewQModelIndex(unsafe.Pointer(index))).UnsafePointer()) }) // slot "sizeHint"
}

func (this *QAbstractItemDelegate) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QAbstractItemDelegate) OnCreateEditor(fn func(*QWidget, *QStyleOptionViewItem, *QModelIndex) *QWidget) {
	registerVirt(this.UnsafeGoID(), -2118980437705941456, func(parent uintptr, option uintptr, index uintptr) uintptr { return uintptr(fn(UnsafeNewQWidget(unsafe.Pointer(parent)), UnsafeNewQStyleOptionViewItem(unsafe.Pointer(option)), UnsafeNewQModelIndex(unsafe.Pointer(index))).UnsafePointer()) }) // slot "createEditor"
}

func (this *QAbstractItemDelegate) OnDestroyEditor(fn func(*QWidget, *QModelIndex)) {
	registerVirt(this.UnsafeGoID(), -7699250917160078568, func(editor uintptr, index uintptr) { fn(UnsafeNewQWidget(unsafe.Pointer(editor)), UnsafeNewQModelIndex(unsafe.Pointer(index))) }) // slot "destroyEditor"
}

func (this *QAbstractItemDelegate) OnSetEditorData(fn func(*QWidget, *QModelIndex)) {
	registerVirt(this.UnsafeGoID(), -858237687273351194, func(editor uintptr, index uintptr) { fn(UnsafeNewQWidget(unsafe.Pointer(editor)), UnsafeNewQModelIndex(unsafe.Pointer(index))) }) // slot "setEditorData"
}

func (this *QAbstractItemDelegate) OnSetModelData(fn func(*QWidget, *QAbstractItemModel, *QModelIndex)) {
	registerVirt(this.UnsafeGoID(), -4740900545199884600, func(editor uintptr, model uintptr, index uintptr) { fn(UnsafeNewQWidget(unsafe.Pointer(editor)), UnsafeNewQAbstractItemModel(unsafe.Pointer(model)), UnsafeNewQModelIndex(unsafe.Pointer(index))) }) // slot "setModelData"
}

func (this *QAbstractItemDelegate) OnUpdateEditorGeometry(fn func(*QWidget, *QStyleOptionViewItem, *QModelIndex)) {
	registerVirt(this.UnsafeGoID(), -1005009386014896055, func(editor uintptr, option uintptr, index uintptr) { fn(UnsafeNewQWidget(unsafe.Pointer(editor)), UnsafeNewQStyleOptionViewItem(unsafe.Pointer(option)), UnsafeNewQModelIndex(unsafe.Pointer(index))) }) // slot "updateEditorGeometry"
}

func (this *QAbstractItemDelegate) OnEditorEvent(fn func(*QEvent, *QAbstractItemModel, *QStyleOptionViewItem, *QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), -2261279149879782144, func(event uintptr, model uintptr, option uintptr, index uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)), UnsafeNewQAbstractItemModel(unsafe.Pointer(model)), UnsafeNewQStyleOptionViewItem(unsafe.Pointer(option)), UnsafeNewQModelIndex(unsafe.Pointer(index)))) }) // slot "editorEvent"
}

func (this *QAbstractItemDelegate) OnHelpEvent(fn func(*QHelpEvent, *QAbstractItemView, *QStyleOptionViewItem, *QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), 1696901855134526294, func(event uintptr, view uintptr, option uintptr, index uintptr) int32 { return b2i(fn(UnsafeNewQHelpEvent(unsafe.Pointer(event)), UnsafeNewQAbstractItemView(unsafe.Pointer(view)), UnsafeNewQStyleOptionViewItem(unsafe.Pointer(option)), UnsafeNewQModelIndex(unsafe.Pointer(index)))) }) // slot "helpEvent"
}

func (this *QAbstractItemDelegate) OnPaintingRoles(fn func() []int) {
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

func (this *QAbstractItemDelegate) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QAbstractItemDelegate) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QAbstractItemDelegate) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QAbstractItemDelegate) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QAbstractItemDelegate) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QAbstractItemDelegate) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QAbstractItemDelegate) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QAbstractItemDelegate) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QAbstractItemDelegate_delete(this.h)
	}
	this.UnsafeRelease()
}
