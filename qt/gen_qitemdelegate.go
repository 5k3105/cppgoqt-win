package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QItemDelegate_chain uintptr
	addr_QItemDelegate_new uintptr
	addr_QItemDelegate_new2 uintptr
	addr_QItemDelegate_tr uintptr
	addr_QItemDelegate_hasClipping uintptr
	addr_QItemDelegate_setClipping uintptr
	addr_QItemDelegate_itemEditorFactory uintptr
	addr_QItemDelegate_setItemEditorFactory uintptr
	addr_QItemDelegate_drawBackground uintptr
	addr_QItemDelegate_doLayout uintptr
	addr_QItemDelegate_rect uintptr
	addr_QItemDelegate_setOptions uintptr
	addr_QItemDelegate_decoration uintptr
	addr_QItemDelegate_doCheck uintptr
	addr_QItemDelegate_textRectangle uintptr
	addr_QItemDelegate_connect uintptr
	addr_QItemDelegate_connect2 uintptr
	addr_QItemDelegate_sender uintptr
	addr_QItemDelegate_senderSignalIndex uintptr
	addr_QItemDelegate_receivers uintptr
	addr_QItemDelegate_isSignalConnected uintptr
	addr_QItemDelegate_tr2 uintptr
	addr_QItemDelegate_tr3 uintptr
	addr_QItemDelegate_connect3 uintptr
	addr_QItemDelegate_connect4 uintptr
	addr_QItemDelegate_metaObject_default uintptr
	addr_QItemDelegate_paint_default uintptr
	addr_QItemDelegate_sizeHint_default uintptr
	addr_QItemDelegate_createEditor_default uintptr
	addr_QItemDelegate_setEditorData_default uintptr
	addr_QItemDelegate_setModelData_default uintptr
	addr_QItemDelegate_updateEditorGeometry_default uintptr
	addr_QItemDelegate_drawDisplay_default uintptr
	addr_QItemDelegate_drawDecoration_default uintptr
	addr_QItemDelegate_drawFocus_default uintptr
	addr_QItemDelegate_drawCheck_default uintptr
	addr_QItemDelegate_eventFilter_default uintptr
	addr_QItemDelegate_editorEvent_default uintptr
	addr_QItemDelegate_destroyEditor_default uintptr
	addr_QItemDelegate_helpEvent_default uintptr
	addr_QItemDelegate_paintingRoles_default uintptr
	addr_QItemDelegate_event_default uintptr
	addr_QItemDelegate_timerEvent_default uintptr
	addr_QItemDelegate_childEvent_default uintptr
	addr_QItemDelegate_customEvent_default uintptr
	addr_QItemDelegate_connectNotify_default uintptr
	addr_QItemDelegate_disconnectNotify_default uintptr
	pg_QItemDelegate_connect_commitData func(uintptr, int64)
	pg_QItemDelegate_connect_closeEditor func(uintptr, int64)
	pg_QItemDelegate_connect_closeEditor2 func(uintptr, int64)
	pg_QItemDelegate_connect_sizeHintChanged func(uintptr, int64)
	pg_QItemDelegate_connect_destroyed func(uintptr, int64)
	pg_QItemDelegate_connect_destroyed2 func(uintptr, int64)
	pg_QItemDelegate_connect_objectNameChanged func(uintptr, int64)
	pg_QItemDelegate__disconnect func(int64)
	pg_QItemDelegate__set_override func(uintptr, int32)
	pg_QItemDelegate_delete func(uintptr)
)

func init() {
	symaddr(&addr_QItemDelegate_chain, "QItemDelegate_chain")
	symaddr(&addr_QItemDelegate_new, "QItemDelegate_new")
	symaddr(&addr_QItemDelegate_new2, "QItemDelegate_new2")
	symaddr(&addr_QItemDelegate_tr, "QItemDelegate_tr")
	symaddr(&addr_QItemDelegate_hasClipping, "QItemDelegate_hasClipping")
	symaddr(&addr_QItemDelegate_setClipping, "QItemDelegate_setClipping")
	symaddr(&addr_QItemDelegate_itemEditorFactory, "QItemDelegate_itemEditorFactory")
	symaddr(&addr_QItemDelegate_setItemEditorFactory, "QItemDelegate_setItemEditorFactory")
	symaddr(&addr_QItemDelegate_drawBackground, "QItemDelegate_drawBackground")
	symaddr(&addr_QItemDelegate_doLayout, "QItemDelegate_doLayout")
	symaddr(&addr_QItemDelegate_rect, "QItemDelegate_rect")
	symaddr(&addr_QItemDelegate_setOptions, "QItemDelegate_setOptions")
	symaddr(&addr_QItemDelegate_decoration, "QItemDelegate_decoration")
	symaddr(&addr_QItemDelegate_doCheck, "QItemDelegate_doCheck")
	symaddr(&addr_QItemDelegate_textRectangle, "QItemDelegate_textRectangle")
	symaddr(&addr_QItemDelegate_connect, "QItemDelegate_connect")
	symaddr(&addr_QItemDelegate_connect2, "QItemDelegate_connect2")
	symaddr(&addr_QItemDelegate_sender, "QItemDelegate_sender")
	symaddr(&addr_QItemDelegate_senderSignalIndex, "QItemDelegate_senderSignalIndex")
	symaddr(&addr_QItemDelegate_receivers, "QItemDelegate_receivers")
	symaddr(&addr_QItemDelegate_isSignalConnected, "QItemDelegate_isSignalConnected")
	symaddr(&addr_QItemDelegate_tr2, "QItemDelegate_tr2")
	symaddr(&addr_QItemDelegate_tr3, "QItemDelegate_tr3")
	symaddr(&addr_QItemDelegate_connect3, "QItemDelegate_connect3")
	symaddr(&addr_QItemDelegate_connect4, "QItemDelegate_connect4")
	symaddr(&addr_QItemDelegate_metaObject_default, "QItemDelegate_metaObject_default")
	symaddr(&addr_QItemDelegate_paint_default, "QItemDelegate_paint_default")
	symaddr(&addr_QItemDelegate_sizeHint_default, "QItemDelegate_sizeHint_default")
	symaddr(&addr_QItemDelegate_createEditor_default, "QItemDelegate_createEditor_default")
	symaddr(&addr_QItemDelegate_setEditorData_default, "QItemDelegate_setEditorData_default")
	symaddr(&addr_QItemDelegate_setModelData_default, "QItemDelegate_setModelData_default")
	symaddr(&addr_QItemDelegate_updateEditorGeometry_default, "QItemDelegate_updateEditorGeometry_default")
	symaddr(&addr_QItemDelegate_drawDisplay_default, "QItemDelegate_drawDisplay_default")
	symaddr(&addr_QItemDelegate_drawDecoration_default, "QItemDelegate_drawDecoration_default")
	symaddr(&addr_QItemDelegate_drawFocus_default, "QItemDelegate_drawFocus_default")
	symaddr(&addr_QItemDelegate_drawCheck_default, "QItemDelegate_drawCheck_default")
	symaddr(&addr_QItemDelegate_eventFilter_default, "QItemDelegate_eventFilter_default")
	symaddr(&addr_QItemDelegate_editorEvent_default, "QItemDelegate_editorEvent_default")
	symaddr(&addr_QItemDelegate_destroyEditor_default, "QItemDelegate_destroyEditor_default")
	symaddr(&addr_QItemDelegate_helpEvent_default, "QItemDelegate_helpEvent_default")
	symaddr(&addr_QItemDelegate_paintingRoles_default, "QItemDelegate_paintingRoles_default")
	symaddr(&addr_QItemDelegate_event_default, "QItemDelegate_event_default")
	symaddr(&addr_QItemDelegate_timerEvent_default, "QItemDelegate_timerEvent_default")
	symaddr(&addr_QItemDelegate_childEvent_default, "QItemDelegate_childEvent_default")
	symaddr(&addr_QItemDelegate_customEvent_default, "QItemDelegate_customEvent_default")
	symaddr(&addr_QItemDelegate_connectNotify_default, "QItemDelegate_connectNotify_default")
	symaddr(&addr_QItemDelegate_disconnectNotify_default, "QItemDelegate_disconnectNotify_default")
	reg(&pg_QItemDelegate_connect_commitData, "QItemDelegate_connect_commitData")
	reg(&pg_QItemDelegate_connect_closeEditor, "QItemDelegate_connect_closeEditor")
	reg(&pg_QItemDelegate_connect_closeEditor2, "QItemDelegate_connect_closeEditor2")
	reg(&pg_QItemDelegate_connect_sizeHintChanged, "QItemDelegate_connect_sizeHintChanged")
	reg(&pg_QItemDelegate_connect_destroyed, "QItemDelegate_connect_destroyed")
	reg(&pg_QItemDelegate_connect_destroyed2, "QItemDelegate_connect_destroyed2")
	reg(&pg_QItemDelegate_connect_objectNameChanged, "QItemDelegate_connect_objectNameChanged")
	reg(&pg_QItemDelegate__disconnect, "QItemDelegate__disconnect")
	reg(&pg_QItemDelegate__set_override, "QItemDelegate__set_override")
	reg(&pg_QItemDelegate_delete, "QItemDelegate_delete")
}

type QItemDelegate struct {
	h uintptr
	*QAbstractItemDelegate
}

func bindQItemDelegate(h uintptr) *QItemDelegate {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QItemDelegate_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QItemDelegate{h: h, QAbstractItemDelegate: &QAbstractItemDelegate{h: __c[0], QObject: &QObject{h: __c[1]}}}
}

func newownedQItemDelegate(h uintptr) *QItemDelegate { return bindQItemDelegate(h) }

func (this *QItemDelegate) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QItemDelegate) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQItemDelegate(h unsafe.Pointer) *QItemDelegate { return bindQItemDelegate(uintptr(h)) }
func UnsafeNewQItemDelegateOwned(h unsafe.Pointer) *QItemDelegate { return newownedQItemDelegate(uintptr(h)) }

func (this *QItemDelegate) UnsafeRelease() {
	this.h = 0
	if this.QAbstractItemDelegate != nil {
		this.QAbstractItemDelegate.UnsafeRelease()
	}
}

func NewQItemDelegate() *QItemDelegate {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QItemDelegate_new, uintptr(__goid))
	__o := newownedQItemDelegate(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QItemDelegate_virtIdx, QItemDelegate_enable)
	return __o
}

func NewQItemDelegate2(parent *QObject) *QItemDelegate {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QItemDelegate_new2, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQItemDelegate(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QItemDelegate_virtIdx, QItemDelegate_enable)
	return __o
}

func QItemDelegate_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QItemDelegate_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QItemDelegate) HasClipping() bool {
	__r0, _, _ := purego.SyscallN(addr_QItemDelegate_hasClipping, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QItemDelegate) SetClipping(clip bool) {
	purego.SyscallN(addr_QItemDelegate_setClipping, this.h, uintptr(b2i(clip)))
	runtime.KeepAlive(this)
}

func (this *QItemDelegate) ItemEditorFactory() *QItemEditorFactory {
	__r, _, _ := purego.SyscallN(addr_QItemDelegate_itemEditorFactory, this.h)
	runtime.KeepAlive(this)
	return bindQItemEditorFactory(__r)
}

func (this *QItemDelegate) SetItemEditorFactory(factory *QItemEditorFactory) {
	purego.SyscallN(addr_QItemDelegate_setItemEditorFactory, this.h, uintptr(factory.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(factory)
}

func (this *QItemDelegate) DrawBackground(painter *QPainter, option *QStyleOptionViewItem, index *QModelIndex) {
	purego.SyscallN(addr_QItemDelegate_drawBackground, this.h, uintptr(painter.UnsafePointer()), uintptr(option.UnsafePointer()), uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
	runtime.KeepAlive(option)
	runtime.KeepAlive(index)
}

func (this *QItemDelegate) DoLayout(option *QStyleOptionViewItem, checkRect *QRect, iconRect *QRect, textRect *QRect, hint bool) {
	purego.SyscallN(addr_QItemDelegate_doLayout, this.h, uintptr(option.UnsafePointer()), uintptr(checkRect.UnsafePointer()), uintptr(iconRect.UnsafePointer()), uintptr(textRect.UnsafePointer()), uintptr(b2i(hint)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(option)
	runtime.KeepAlive(checkRect)
	runtime.KeepAlive(iconRect)
	runtime.KeepAlive(textRect)
}

func (this *QItemDelegate) Rect(option *QStyleOptionViewItem, index *QModelIndex, role int) *QRect {
	__r, _, _ := purego.SyscallN(addr_QItemDelegate_rect, this.h, uintptr(option.UnsafePointer()), uintptr(index.UnsafePointer()), uintptr(role))
	runtime.KeepAlive(this)
	runtime.KeepAlive(option)
	runtime.KeepAlive(index)
	return newownedQRect(__r)
}

func (this *QItemDelegate) SetOptions(index *QModelIndex, option *QStyleOptionViewItem) *QStyleOptionViewItem {
	__r, _, _ := purego.SyscallN(addr_QItemDelegate_setOptions, this.h, uintptr(index.UnsafePointer()), uintptr(option.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	runtime.KeepAlive(option)
	return newownedQStyleOptionViewItem(__r)
}

func (this *QItemDelegate) Decoration(option *QStyleOptionViewItem, variant *QVariant) *QPixmap {
	__r, _, _ := purego.SyscallN(addr_QItemDelegate_decoration, this.h, uintptr(option.UnsafePointer()), uintptr(variant.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(option)
	runtime.KeepAlive(variant)
	return newownedQPixmap(__r)
}

func (this *QItemDelegate) DoCheck(option *QStyleOptionViewItem, bounding *QRect, variant *QVariant) *QRect {
	__r, _, _ := purego.SyscallN(addr_QItemDelegate_doCheck, this.h, uintptr(option.UnsafePointer()), uintptr(bounding.UnsafePointer()), uintptr(variant.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(option)
	runtime.KeepAlive(bounding)
	runtime.KeepAlive(variant)
	return newownedQRect(__r)
}

func (this *QItemDelegate) TextRectangle(painter *QPainter, rect *QRect, font *QFont, text string) *QRect {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QItemDelegate_textRectangle, this.h, uintptr(painter.UnsafePointer()), uintptr(rect.UnsafePointer()), uintptr(font.UnsafePointer()), text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
	runtime.KeepAlive(rect)
	runtime.KeepAlive(font)
	runtime.KeepAlive(text)
	return newownedQRect(__r)
}

func QItemDelegate_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QItemDelegate_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QItemDelegate_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QItemDelegate_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QItemDelegate) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QItemDelegate_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QItemDelegate) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QItemDelegate_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QItemDelegate) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QItemDelegate_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QItemDelegate) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QItemDelegate_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QItemDelegate_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QItemDelegate_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QItemDelegate_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QItemDelegate_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QItemDelegate_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QItemDelegate_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QItemDelegate_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QItemDelegate_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QItemDelegate) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QItemDelegate_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QItemDelegate) PaintDefault(painter *QPainter, option *QStyleOptionViewItem, index *QModelIndex) {
	purego.SyscallN(addr_QItemDelegate_paint_default, this.h, uintptr(painter.UnsafePointer()), uintptr(option.UnsafePointer()), uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
	runtime.KeepAlive(option)
	runtime.KeepAlive(index)
}

func (this *QItemDelegate) SizeHintDefault(option *QStyleOptionViewItem, index *QModelIndex) *QSize {
	__r, _, _ := purego.SyscallN(addr_QItemDelegate_sizeHint_default, this.h, uintptr(option.UnsafePointer()), uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(option)
	runtime.KeepAlive(index)
	return newownedQSize(__r)
}

func (this *QItemDelegate) CreateEditorDefault(parent *QWidget, option *QStyleOptionViewItem, index *QModelIndex) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QItemDelegate_createEditor_default, this.h, uintptr(parent.UnsafePointer()), uintptr(option.UnsafePointer()), uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	runtime.KeepAlive(option)
	runtime.KeepAlive(index)
	return bindQWidget(__r)
}

func (this *QItemDelegate) SetEditorDataDefault(editor *QWidget, index *QModelIndex) {
	purego.SyscallN(addr_QItemDelegate_setEditorData_default, this.h, uintptr(editor.UnsafePointer()), uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(editor)
	runtime.KeepAlive(index)
}

func (this *QItemDelegate) SetModelDataDefault(editor *QWidget, model *QAbstractItemModel, index *QModelIndex) {
	purego.SyscallN(addr_QItemDelegate_setModelData_default, this.h, uintptr(editor.UnsafePointer()), uintptr(model.UnsafePointer()), uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(editor)
	runtime.KeepAlive(model)
	runtime.KeepAlive(index)
}

func (this *QItemDelegate) UpdateEditorGeometryDefault(editor *QWidget, option *QStyleOptionViewItem, index *QModelIndex) {
	purego.SyscallN(addr_QItemDelegate_updateEditorGeometry_default, this.h, uintptr(editor.UnsafePointer()), uintptr(option.UnsafePointer()), uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(editor)
	runtime.KeepAlive(option)
	runtime.KeepAlive(index)
}

func (this *QItemDelegate) DrawDisplayDefault(painter *QPainter, option *QStyleOptionViewItem, rect *QRect, text string) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QItemDelegate_drawDisplay_default, this.h, uintptr(painter.UnsafePointer()), uintptr(option.UnsafePointer()), uintptr(rect.UnsafePointer()), text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
	runtime.KeepAlive(option)
	runtime.KeepAlive(rect)
	runtime.KeepAlive(text)
}

func (this *QItemDelegate) DrawDecorationDefault(painter *QPainter, option *QStyleOptionViewItem, rect *QRect, pixmap *QPixmap) {
	purego.SyscallN(addr_QItemDelegate_drawDecoration_default, this.h, uintptr(painter.UnsafePointer()), uintptr(option.UnsafePointer()), uintptr(rect.UnsafePointer()), uintptr(pixmap.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
	runtime.KeepAlive(option)
	runtime.KeepAlive(rect)
	runtime.KeepAlive(pixmap)
}

func (this *QItemDelegate) DrawFocusDefault(painter *QPainter, option *QStyleOptionViewItem, rect *QRect) {
	purego.SyscallN(addr_QItemDelegate_drawFocus_default, this.h, uintptr(painter.UnsafePointer()), uintptr(option.UnsafePointer()), uintptr(rect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
	runtime.KeepAlive(option)
	runtime.KeepAlive(rect)
}

func (this *QItemDelegate) DrawCheckDefault(painter *QPainter, option *QStyleOptionViewItem, rect *QRect, state CheckState) {
	purego.SyscallN(addr_QItemDelegate_drawCheck_default, this.h, uintptr(painter.UnsafePointer()), uintptr(option.UnsafePointer()), uintptr(rect.UnsafePointer()), uintptr(state))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
	runtime.KeepAlive(option)
	runtime.KeepAlive(rect)
}

func (this *QItemDelegate) EventFilterDefault(object *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QItemDelegate_eventFilter_default, this.h, uintptr(object.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(object)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QItemDelegate) EditorEventDefault(event *QEvent, model *QAbstractItemModel, option *QStyleOptionViewItem, index *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QItemDelegate_editorEvent_default, this.h, uintptr(event.UnsafePointer()), uintptr(model.UnsafePointer()), uintptr(option.UnsafePointer()), uintptr(index.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	runtime.KeepAlive(model)
	runtime.KeepAlive(option)
	runtime.KeepAlive(index)
	return __r != 0
}

func (this *QItemDelegate) DestroyEditorDefault(editor *QWidget, index *QModelIndex) {
	purego.SyscallN(addr_QItemDelegate_destroyEditor_default, this.h, uintptr(editor.UnsafePointer()), uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(editor)
	runtime.KeepAlive(index)
}

func (this *QItemDelegate) HelpEventDefault(event *QHelpEvent, view *QAbstractItemView, option *QStyleOptionViewItem, index *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QItemDelegate_helpEvent_default, this.h, uintptr(event.UnsafePointer()), uintptr(view.UnsafePointer()), uintptr(option.UnsafePointer()), uintptr(index.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	runtime.KeepAlive(view)
	runtime.KeepAlive(option)
	runtime.KeepAlive(index)
	return __r != 0
}

func (this *QItemDelegate) PaintingRolesDefault() []int {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QItemDelegate_paintingRoles_default, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]int, __n)
	__sl := unsafe.Slice((*int32)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = int(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QItemDelegate) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QItemDelegate_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QItemDelegate) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QItemDelegate_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QItemDelegate) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QItemDelegate_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QItemDelegate) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QItemDelegate_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QItemDelegate) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QItemDelegate_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QItemDelegate) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QItemDelegate_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QItemDelegate) OnCommitData(fn func(*QWidget)) int64 {
	id := registerCallback(func(editor uintptr) { fn(UnsafeNewQWidget(unsafe.Pointer(editor))) })
	pg_QItemDelegate_connect_commitData(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QItemDelegate) OnCloseEditor(fn func(*QWidget)) int64 {
	id := registerCallback(func(editor uintptr) { fn(UnsafeNewQWidget(unsafe.Pointer(editor))) })
	pg_QItemDelegate_connect_closeEditor(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QItemDelegate) OnCloseEditor2(fn func(*QWidget, QAbstractItemDelegate__EndEditHint)) int64 {
	id := registerCallback(func(editor uintptr, hint int32) { fn(UnsafeNewQWidget(unsafe.Pointer(editor)), QAbstractItemDelegate__EndEditHint(hint)) })
	pg_QItemDelegate_connect_closeEditor2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QItemDelegate) OnSizeHintChanged(fn func(*QModelIndex)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(arg1))) })
	pg_QItemDelegate_connect_sizeHintChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QItemDelegate) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QItemDelegate_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QItemDelegate) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QItemDelegate_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QItemDelegate) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QItemDelegate_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QItemDelegate) Disconnect(handle int64) {
	pg_QItemDelegate__disconnect(handle)
	unregisterCallback(handle)
}

func QItemDelegate_virtIdx(slot int64) int32 {
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
	case -5011630111077710685:
		return 7
	case -6476941580654430195:
		return 8
	case -5402985858130777771:
		return 9
	case -8060427496448215275:
		return 10
	case 5217623514630858297:
		return 11
	case -2261279149879782144:
		return 12
	case -7699250917160078568:
		return 13
	case 1696901855134526294:
		return 14
	case -4420854320663935586:
		return 15
	case 1237236456312824991:
		return 16
	case -5409927581812374498:
		return 17
	case -1185315367526857833:
		return 18
	case 3769868229123868474:
		return 19
	case 7379029982170503008:
		return 20
	case -568369772982215962:
		return 21
	}
	return -1
}

func QItemDelegate_enable(h uintptr, idx int32) { pg_QItemDelegate__set_override(h, idx) }

func (this *QItemDelegate) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QItemDelegate) OnPaint(fn func(*QPainter, *QStyleOptionViewItem, *QModelIndex)) {
	registerVirt(this.UnsafeGoID(), 6083584261053556819, func(painter uintptr, option uintptr, index uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter)), UnsafeNewQStyleOptionViewItem(unsafe.Pointer(option)), UnsafeNewQModelIndex(unsafe.Pointer(index))) }) // slot "paint"
}

func (this *QItemDelegate) OnSizeHint(fn func(*QStyleOptionViewItem, *QModelIndex) *QSize) {
	registerVirt(this.UnsafeGoID(), 729690261457684101, func(option uintptr, index uintptr) uintptr { return uintptr(fn(UnsafeNewQStyleOptionViewItem(unsafe.Pointer(option)), UnsafeNewQModelIndex(unsafe.Pointer(index))).UnsafePointer()) }) // slot "sizeHint"
}

func (this *QItemDelegate) OnCreateEditor(fn func(*QWidget, *QStyleOptionViewItem, *QModelIndex) *QWidget) {
	registerVirt(this.UnsafeGoID(), -2118980437705941456, func(parent uintptr, option uintptr, index uintptr) uintptr { return uintptr(fn(UnsafeNewQWidget(unsafe.Pointer(parent)), UnsafeNewQStyleOptionViewItem(unsafe.Pointer(option)), UnsafeNewQModelIndex(unsafe.Pointer(index))).UnsafePointer()) }) // slot "createEditor"
}

func (this *QItemDelegate) OnSetEditorData(fn func(*QWidget, *QModelIndex)) {
	registerVirt(this.UnsafeGoID(), -858237687273351194, func(editor uintptr, index uintptr) { fn(UnsafeNewQWidget(unsafe.Pointer(editor)), UnsafeNewQModelIndex(unsafe.Pointer(index))) }) // slot "setEditorData"
}

func (this *QItemDelegate) OnSetModelData(fn func(*QWidget, *QAbstractItemModel, *QModelIndex)) {
	registerVirt(this.UnsafeGoID(), -4740900545199884600, func(editor uintptr, model uintptr, index uintptr) { fn(UnsafeNewQWidget(unsafe.Pointer(editor)), UnsafeNewQAbstractItemModel(unsafe.Pointer(model)), UnsafeNewQModelIndex(unsafe.Pointer(index))) }) // slot "setModelData"
}

func (this *QItemDelegate) OnUpdateEditorGeometry(fn func(*QWidget, *QStyleOptionViewItem, *QModelIndex)) {
	registerVirt(this.UnsafeGoID(), -1005009386014896055, func(editor uintptr, option uintptr, index uintptr) { fn(UnsafeNewQWidget(unsafe.Pointer(editor)), UnsafeNewQStyleOptionViewItem(unsafe.Pointer(option)), UnsafeNewQModelIndex(unsafe.Pointer(index))) }) // slot "updateEditorGeometry"
}

func (this *QItemDelegate) OnDrawDisplay(fn func(*QPainter, *QStyleOptionViewItem, *QRect, string)) {
	registerVirt(this.UnsafeGoID(), -5011630111077710685, func(painter uintptr, option uintptr, rect uintptr, text uintptr, text_len int64) { fn(UnsafeNewQPainter(unsafe.Pointer(painter)), UnsafeNewQStyleOptionViewItem(unsafe.Pointer(option)), UnsafeNewQRect(unsafe.Pointer(rect)), pgStrNoFree(text, text_len)) }) // slot "drawDisplay"
}

func (this *QItemDelegate) OnDrawDecoration(fn func(*QPainter, *QStyleOptionViewItem, *QRect, *QPixmap)) {
	registerVirt(this.UnsafeGoID(), -6476941580654430195, func(painter uintptr, option uintptr, rect uintptr, pixmap uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter)), UnsafeNewQStyleOptionViewItem(unsafe.Pointer(option)), UnsafeNewQRect(unsafe.Pointer(rect)), UnsafeNewQPixmap(unsafe.Pointer(pixmap))) }) // slot "drawDecoration"
}

func (this *QItemDelegate) OnDrawFocus(fn func(*QPainter, *QStyleOptionViewItem, *QRect)) {
	registerVirt(this.UnsafeGoID(), -5402985858130777771, func(painter uintptr, option uintptr, rect uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter)), UnsafeNewQStyleOptionViewItem(unsafe.Pointer(option)), UnsafeNewQRect(unsafe.Pointer(rect))) }) // slot "drawFocus"
}

func (this *QItemDelegate) OnDrawCheck(fn func(*QPainter, *QStyleOptionViewItem, *QRect, CheckState)) {
	registerVirt(this.UnsafeGoID(), -8060427496448215275, func(painter uintptr, option uintptr, rect uintptr, state int32) { fn(UnsafeNewQPainter(unsafe.Pointer(painter)), UnsafeNewQStyleOptionViewItem(unsafe.Pointer(option)), UnsafeNewQRect(unsafe.Pointer(rect)), CheckState(state)) }) // slot "drawCheck"
}

func (this *QItemDelegate) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(object uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(object)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QItemDelegate) OnEditorEvent(fn func(*QEvent, *QAbstractItemModel, *QStyleOptionViewItem, *QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), -2261279149879782144, func(event uintptr, model uintptr, option uintptr, index uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)), UnsafeNewQAbstractItemModel(unsafe.Pointer(model)), UnsafeNewQStyleOptionViewItem(unsafe.Pointer(option)), UnsafeNewQModelIndex(unsafe.Pointer(index)))) }) // slot "editorEvent"
}

func (this *QItemDelegate) OnDestroyEditor(fn func(*QWidget, *QModelIndex)) {
	registerVirt(this.UnsafeGoID(), -7699250917160078568, func(editor uintptr, index uintptr) { fn(UnsafeNewQWidget(unsafe.Pointer(editor)), UnsafeNewQModelIndex(unsafe.Pointer(index))) }) // slot "destroyEditor"
}

func (this *QItemDelegate) OnHelpEvent(fn func(*QHelpEvent, *QAbstractItemView, *QStyleOptionViewItem, *QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), 1696901855134526294, func(event uintptr, view uintptr, option uintptr, index uintptr) int32 { return b2i(fn(UnsafeNewQHelpEvent(unsafe.Pointer(event)), UnsafeNewQAbstractItemView(unsafe.Pointer(view)), UnsafeNewQStyleOptionViewItem(unsafe.Pointer(option)), UnsafeNewQModelIndex(unsafe.Pointer(index)))) }) // slot "helpEvent"
}

func (this *QItemDelegate) OnPaintingRoles(fn func() []int) {
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

func (this *QItemDelegate) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QItemDelegate) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QItemDelegate) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QItemDelegate) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QItemDelegate) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QItemDelegate) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QItemDelegate) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QItemDelegate_delete(this.h)
	}
	this.UnsafeRelease()
}
