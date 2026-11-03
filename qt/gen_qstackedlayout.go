package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QStackedLayout_chain uintptr
	addr_QStackedLayout_new uintptr
	addr_QStackedLayout_new2 uintptr
	addr_QStackedLayout_new3 uintptr
	addr_QStackedLayout_tr uintptr
	addr_QStackedLayout_addWidget uintptr
	addr_QStackedLayout_insertWidget uintptr
	addr_QStackedLayout_currentWidget uintptr
	addr_QStackedLayout_currentIndex uintptr
	addr_QStackedLayout_widget uintptr
	addr_QStackedLayout_stackingMode uintptr
	addr_QStackedLayout_setStackingMode uintptr
	addr_QStackedLayout_widgetRemoved uintptr
	addr_QStackedLayout_currentChanged uintptr
	addr_QStackedLayout_widgetAdded uintptr
	addr_QStackedLayout_setCurrentIndex uintptr
	addr_QStackedLayout_setCurrentWidget uintptr
	addr_QStackedLayout_closestAcceptableSize uintptr
	addr_QStackedLayout_widgetEvent uintptr
	addr_QStackedLayout_addChildLayout uintptr
	addr_QStackedLayout_addChildWidget uintptr
	addr_QStackedLayout_adoptLayout uintptr
	addr_QStackedLayout_alignmentRect uintptr
	addr_QStackedLayout_connect uintptr
	addr_QStackedLayout_connect2 uintptr
	addr_QStackedLayout_sender uintptr
	addr_QStackedLayout_senderSignalIndex uintptr
	addr_QStackedLayout_receivers uintptr
	addr_QStackedLayout_isSignalConnected uintptr
	addr_QStackedLayout_tr2 uintptr
	addr_QStackedLayout_tr3 uintptr
	addr_QStackedLayout_connect3 uintptr
	addr_QStackedLayout_connect4 uintptr
	addr_QStackedLayout_metaObject_default uintptr
	addr_QStackedLayout_count_default uintptr
	addr_QStackedLayout_addItem_default uintptr
	addr_QStackedLayout_sizeHint_default uintptr
	addr_QStackedLayout_minimumSize_default uintptr
	addr_QStackedLayout_itemAt_default uintptr
	addr_QStackedLayout_takeAt_default uintptr
	addr_QStackedLayout_setGeometry_default uintptr
	addr_QStackedLayout_hasHeightForWidth_default uintptr
	addr_QStackedLayout_heightForWidth_default uintptr
	addr_QStackedLayout_spacing_default uintptr
	addr_QStackedLayout_setSpacing_default uintptr
	addr_QStackedLayout_invalidate_default uintptr
	addr_QStackedLayout_geometry_default uintptr
	addr_QStackedLayout_expandingDirections_default uintptr
	addr_QStackedLayout_maximumSize_default uintptr
	addr_QStackedLayout_indexOf_default uintptr
	addr_QStackedLayout_isEmpty_default uintptr
	addr_QStackedLayout_controlTypes_default uintptr
	addr_QStackedLayout_replaceWidget_default uintptr
	addr_QStackedLayout_layout_default uintptr
	addr_QStackedLayout_childEvent_default uintptr
	addr_QStackedLayout_event_default uintptr
	addr_QStackedLayout_eventFilter_default uintptr
	addr_QStackedLayout_timerEvent_default uintptr
	addr_QStackedLayout_customEvent_default uintptr
	addr_QStackedLayout_connectNotify_default uintptr
	addr_QStackedLayout_disconnectNotify_default uintptr
	addr_QStackedLayout_minimumHeightForWidth_default uintptr
	addr_QStackedLayout_spacerItem_default uintptr
	pg_QStackedLayout_connect_widgetRemoved func(uintptr, int64)
	pg_QStackedLayout_connect_currentChanged func(uintptr, int64)
	pg_QStackedLayout_connect_widgetAdded func(uintptr, int64)
	pg_QStackedLayout_connect_destroyed func(uintptr, int64)
	pg_QStackedLayout_connect_destroyed2 func(uintptr, int64)
	pg_QStackedLayout_connect_objectNameChanged func(uintptr, int64)
	pg_QStackedLayout__disconnect func(int64)
	pg_QStackedLayout__set_override func(uintptr, int32)
	pg_QStackedLayout_delete func(uintptr)
)

func init() {
	symaddr(&addr_QStackedLayout_chain, "QStackedLayout_chain")
	symaddr(&addr_QStackedLayout_new, "QStackedLayout_new")
	symaddr(&addr_QStackedLayout_new2, "QStackedLayout_new2")
	symaddr(&addr_QStackedLayout_new3, "QStackedLayout_new3")
	symaddr(&addr_QStackedLayout_tr, "QStackedLayout_tr")
	symaddr(&addr_QStackedLayout_addWidget, "QStackedLayout_addWidget")
	symaddr(&addr_QStackedLayout_insertWidget, "QStackedLayout_insertWidget")
	symaddr(&addr_QStackedLayout_currentWidget, "QStackedLayout_currentWidget")
	symaddr(&addr_QStackedLayout_currentIndex, "QStackedLayout_currentIndex")
	symaddr(&addr_QStackedLayout_widget, "QStackedLayout_widget")
	symaddr(&addr_QStackedLayout_stackingMode, "QStackedLayout_stackingMode")
	symaddr(&addr_QStackedLayout_setStackingMode, "QStackedLayout_setStackingMode")
	symaddr(&addr_QStackedLayout_widgetRemoved, "QStackedLayout_widgetRemoved")
	symaddr(&addr_QStackedLayout_currentChanged, "QStackedLayout_currentChanged")
	symaddr(&addr_QStackedLayout_widgetAdded, "QStackedLayout_widgetAdded")
	symaddr(&addr_QStackedLayout_setCurrentIndex, "QStackedLayout_setCurrentIndex")
	symaddr(&addr_QStackedLayout_setCurrentWidget, "QStackedLayout_setCurrentWidget")
	symaddr(&addr_QStackedLayout_closestAcceptableSize, "QStackedLayout_closestAcceptableSize")
	symaddr(&addr_QStackedLayout_widgetEvent, "QStackedLayout_widgetEvent")
	symaddr(&addr_QStackedLayout_addChildLayout, "QStackedLayout_addChildLayout")
	symaddr(&addr_QStackedLayout_addChildWidget, "QStackedLayout_addChildWidget")
	symaddr(&addr_QStackedLayout_adoptLayout, "QStackedLayout_adoptLayout")
	symaddr(&addr_QStackedLayout_alignmentRect, "QStackedLayout_alignmentRect")
	symaddr(&addr_QStackedLayout_connect, "QStackedLayout_connect")
	symaddr(&addr_QStackedLayout_connect2, "QStackedLayout_connect2")
	symaddr(&addr_QStackedLayout_sender, "QStackedLayout_sender")
	symaddr(&addr_QStackedLayout_senderSignalIndex, "QStackedLayout_senderSignalIndex")
	symaddr(&addr_QStackedLayout_receivers, "QStackedLayout_receivers")
	symaddr(&addr_QStackedLayout_isSignalConnected, "QStackedLayout_isSignalConnected")
	symaddr(&addr_QStackedLayout_tr2, "QStackedLayout_tr2")
	symaddr(&addr_QStackedLayout_tr3, "QStackedLayout_tr3")
	symaddr(&addr_QStackedLayout_connect3, "QStackedLayout_connect3")
	symaddr(&addr_QStackedLayout_connect4, "QStackedLayout_connect4")
	symaddr(&addr_QStackedLayout_metaObject_default, "QStackedLayout_metaObject_default")
	symaddr(&addr_QStackedLayout_count_default, "QStackedLayout_count_default")
	symaddr(&addr_QStackedLayout_addItem_default, "QStackedLayout_addItem_default")
	symaddr(&addr_QStackedLayout_sizeHint_default, "QStackedLayout_sizeHint_default")
	symaddr(&addr_QStackedLayout_minimumSize_default, "QStackedLayout_minimumSize_default")
	symaddr(&addr_QStackedLayout_itemAt_default, "QStackedLayout_itemAt_default")
	symaddr(&addr_QStackedLayout_takeAt_default, "QStackedLayout_takeAt_default")
	symaddr(&addr_QStackedLayout_setGeometry_default, "QStackedLayout_setGeometry_default")
	symaddr(&addr_QStackedLayout_hasHeightForWidth_default, "QStackedLayout_hasHeightForWidth_default")
	symaddr(&addr_QStackedLayout_heightForWidth_default, "QStackedLayout_heightForWidth_default")
	symaddr(&addr_QStackedLayout_spacing_default, "QStackedLayout_spacing_default")
	symaddr(&addr_QStackedLayout_setSpacing_default, "QStackedLayout_setSpacing_default")
	symaddr(&addr_QStackedLayout_invalidate_default, "QStackedLayout_invalidate_default")
	symaddr(&addr_QStackedLayout_geometry_default, "QStackedLayout_geometry_default")
	symaddr(&addr_QStackedLayout_expandingDirections_default, "QStackedLayout_expandingDirections_default")
	symaddr(&addr_QStackedLayout_maximumSize_default, "QStackedLayout_maximumSize_default")
	symaddr(&addr_QStackedLayout_indexOf_default, "QStackedLayout_indexOf_default")
	symaddr(&addr_QStackedLayout_isEmpty_default, "QStackedLayout_isEmpty_default")
	symaddr(&addr_QStackedLayout_controlTypes_default, "QStackedLayout_controlTypes_default")
	symaddr(&addr_QStackedLayout_replaceWidget_default, "QStackedLayout_replaceWidget_default")
	symaddr(&addr_QStackedLayout_layout_default, "QStackedLayout_layout_default")
	symaddr(&addr_QStackedLayout_childEvent_default, "QStackedLayout_childEvent_default")
	symaddr(&addr_QStackedLayout_event_default, "QStackedLayout_event_default")
	symaddr(&addr_QStackedLayout_eventFilter_default, "QStackedLayout_eventFilter_default")
	symaddr(&addr_QStackedLayout_timerEvent_default, "QStackedLayout_timerEvent_default")
	symaddr(&addr_QStackedLayout_customEvent_default, "QStackedLayout_customEvent_default")
	symaddr(&addr_QStackedLayout_connectNotify_default, "QStackedLayout_connectNotify_default")
	symaddr(&addr_QStackedLayout_disconnectNotify_default, "QStackedLayout_disconnectNotify_default")
	symaddr(&addr_QStackedLayout_minimumHeightForWidth_default, "QStackedLayout_minimumHeightForWidth_default")
	symaddr(&addr_QStackedLayout_spacerItem_default, "QStackedLayout_spacerItem_default")
	reg(&pg_QStackedLayout_connect_widgetRemoved, "QStackedLayout_connect_widgetRemoved")
	reg(&pg_QStackedLayout_connect_currentChanged, "QStackedLayout_connect_currentChanged")
	reg(&pg_QStackedLayout_connect_widgetAdded, "QStackedLayout_connect_widgetAdded")
	reg(&pg_QStackedLayout_connect_destroyed, "QStackedLayout_connect_destroyed")
	reg(&pg_QStackedLayout_connect_destroyed2, "QStackedLayout_connect_destroyed2")
	reg(&pg_QStackedLayout_connect_objectNameChanged, "QStackedLayout_connect_objectNameChanged")
	reg(&pg_QStackedLayout__disconnect, "QStackedLayout__disconnect")
	reg(&pg_QStackedLayout__set_override, "QStackedLayout__set_override")
	reg(&pg_QStackedLayout_delete, "QStackedLayout_delete")
}

type QStackedLayout struct {
	h uintptr
	*QLayout
}

func bindQStackedLayout(h uintptr) *QStackedLayout {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QStackedLayout_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QStackedLayout{h: h, QLayout: &QLayout{h: __c[0], QObject: &QObject{h: __c[1]}}}
}

func newownedQStackedLayout(h uintptr) *QStackedLayout { return bindQStackedLayout(h) }

func (this *QStackedLayout) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QStackedLayout) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQStackedLayout(h unsafe.Pointer) *QStackedLayout { return bindQStackedLayout(uintptr(h)) }
func UnsafeNewQStackedLayoutOwned(h unsafe.Pointer) *QStackedLayout { return newownedQStackedLayout(uintptr(h)) }

func (this *QStackedLayout) UnsafeRelease() {
	this.h = 0
	if this.QLayout != nil {
		this.QLayout.UnsafeRelease()
	}
}

func NewQStackedLayout(parent *QWidget) *QStackedLayout {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QStackedLayout_new, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQStackedLayout(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QStackedLayout_virtIdx, QStackedLayout_enable)
	return __o
}

func NewQStackedLayout2() *QStackedLayout {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QStackedLayout_new2, uintptr(__goid))
	__o := newownedQStackedLayout(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QStackedLayout_virtIdx, QStackedLayout_enable)
	return __o
}

func NewQStackedLayout3(parentLayout *QLayout) *QStackedLayout {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QStackedLayout_new3, uintptr(__goid), uintptr(parentLayout.UnsafePointer()))
	__o := newownedQStackedLayout(__r)
	runtime.KeepAlive(parentLayout)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QStackedLayout_virtIdx, QStackedLayout_enable)
	return __o
}

func QStackedLayout_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStackedLayout_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QStackedLayout) AddWidget(w *QWidget) int {
	__r0, _, _ := purego.SyscallN(addr_QStackedLayout_addWidget, this.h, uintptr(w.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(w)
	return int(__r)
}

func (this *QStackedLayout) InsertWidget(index int, w *QWidget) int {
	__r0, _, _ := purego.SyscallN(addr_QStackedLayout_insertWidget, this.h, uintptr(index), uintptr(w.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(w)
	return int(__r)
}

func (this *QStackedLayout) CurrentWidget() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QStackedLayout_currentWidget, this.h)
	runtime.KeepAlive(this)
	return bindQWidget(__r)
}

func (this *QStackedLayout) CurrentIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QStackedLayout_currentIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QStackedLayout) Widget(arg1 int) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QStackedLayout_widget, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return bindQWidget(__r)
}

func (this *QStackedLayout) StackingMode() QStackedLayout__StackingMode {
	__r0, _, _ := purego.SyscallN(addr_QStackedLayout_stackingMode, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QStackedLayout__StackingMode(__r)
}

func (this *QStackedLayout) SetStackingMode(stackingMode QStackedLayout__StackingMode) {
	purego.SyscallN(addr_QStackedLayout_setStackingMode, this.h, uintptr(stackingMode))
	runtime.KeepAlive(this)
}

func (this *QStackedLayout) WidgetRemoved(index int) {
	purego.SyscallN(addr_QStackedLayout_widgetRemoved, this.h, uintptr(index))
	runtime.KeepAlive(this)
}

func (this *QStackedLayout) CurrentChanged(index int) {
	purego.SyscallN(addr_QStackedLayout_currentChanged, this.h, uintptr(index))
	runtime.KeepAlive(this)
}

func (this *QStackedLayout) WidgetAdded(index int) {
	purego.SyscallN(addr_QStackedLayout_widgetAdded, this.h, uintptr(index))
	runtime.KeepAlive(this)
}

func (this *QStackedLayout) SetCurrentIndex(index int) {
	purego.SyscallN(addr_QStackedLayout_setCurrentIndex, this.h, uintptr(index))
	runtime.KeepAlive(this)
}

func (this *QStackedLayout) SetCurrentWidget(w *QWidget) {
	purego.SyscallN(addr_QStackedLayout_setCurrentWidget, this.h, uintptr(w.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(w)
}

func QStackedLayout_ClosestAcceptableSize(w *QWidget, s *QSize) *QSize {
	__r, _, _ := purego.SyscallN(addr_QStackedLayout_closestAcceptableSize, uintptr(w.UnsafePointer()), uintptr(s.UnsafePointer()))
	runtime.KeepAlive(w)
	runtime.KeepAlive(s)
	return newownedQSize(__r)
}

func (this *QStackedLayout) WidgetEvent(arg1 *QEvent) {
	purego.SyscallN(addr_QStackedLayout_widgetEvent, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QStackedLayout) AddChildLayout(l *QLayout) {
	purego.SyscallN(addr_QStackedLayout_addChildLayout, this.h, uintptr(l.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(l)
}

func (this *QStackedLayout) AddChildWidget(w *QWidget) {
	purego.SyscallN(addr_QStackedLayout_addChildWidget, this.h, uintptr(w.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(w)
}

func (this *QStackedLayout) AdoptLayout(layout *QLayout) bool {
	__r0, _, _ := purego.SyscallN(addr_QStackedLayout_adoptLayout, this.h, uintptr(layout.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(layout)
	return __r != 0
}

func (this *QStackedLayout) AlignmentRect(arg1 *QRect) *QRect {
	__r, _, _ := purego.SyscallN(addr_QStackedLayout_alignmentRect, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return newownedQRect(__r)
}

func QStackedLayout_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QStackedLayout_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QStackedLayout_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QStackedLayout_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QStackedLayout) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QStackedLayout_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QStackedLayout) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QStackedLayout_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QStackedLayout) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QStackedLayout_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QStackedLayout) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QStackedLayout_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QStackedLayout_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStackedLayout_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QStackedLayout_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStackedLayout_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QStackedLayout_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QStackedLayout_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QStackedLayout_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QStackedLayout_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QStackedLayout) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QStackedLayout_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QStackedLayout) CountDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QStackedLayout_count_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QStackedLayout) AddItemDefault(item *QLayoutItem) {
	purego.SyscallN(addr_QStackedLayout_addItem_default, this.h, uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QStackedLayout) SizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QStackedLayout_sizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QStackedLayout) MinimumSizeDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QStackedLayout_minimumSize_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QStackedLayout) ItemAtDefault(arg1 int) *QLayoutItem {
	__r, _, _ := purego.SyscallN(addr_QStackedLayout_itemAt_default, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return bindQLayoutItem(__r)
}

func (this *QStackedLayout) TakeAtDefault(arg1 int) *QLayoutItem {
	__r, _, _ := purego.SyscallN(addr_QStackedLayout_takeAt_default, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return bindQLayoutItem(__r)
}

func (this *QStackedLayout) SetGeometryDefault(rect *QRect) {
	purego.SyscallN(addr_QStackedLayout_setGeometry_default, this.h, uintptr(rect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
}

func (this *QStackedLayout) HasHeightForWidthDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QStackedLayout_hasHeightForWidth_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QStackedLayout) HeightForWidthDefault(width int) int {
	__r0, _, _ := purego.SyscallN(addr_QStackedLayout_heightForWidth_default, this.h, uintptr(width))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QStackedLayout) SpacingDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QStackedLayout_spacing_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QStackedLayout) SetSpacingDefault(arg1 int) {
	purego.SyscallN(addr_QStackedLayout_setSpacing_default, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QStackedLayout) InvalidateDefault() {
	purego.SyscallN(addr_QStackedLayout_invalidate_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QStackedLayout) GeometryDefault() *QRect {
	__r, _, _ := purego.SyscallN(addr_QStackedLayout_geometry_default, this.h)
	runtime.KeepAlive(this)
	return newownedQRect(__r)
}

func (this *QStackedLayout) ExpandingDirectionsDefault() Orientation {
	__r0, _, _ := purego.SyscallN(addr_QStackedLayout_expandingDirections_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return Orientation(__r)
}

func (this *QStackedLayout) MaximumSizeDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QStackedLayout_maximumSize_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QStackedLayout) IndexOfDefault(arg1 *QWidget) int {
	__r0, _, _ := purego.SyscallN(addr_QStackedLayout_indexOf_default, this.h, uintptr(arg1.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return int(__r)
}

func (this *QStackedLayout) IsEmptyDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QStackedLayout_isEmpty_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QStackedLayout) ControlTypesDefault() QSizePolicy__ControlType {
	__r0, _, _ := purego.SyscallN(addr_QStackedLayout_controlTypes_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QSizePolicy__ControlType(__r)
}

func (this *QStackedLayout) ReplaceWidgetDefault(from *QWidget, to *QWidget, options FindChildOption) *QLayoutItem {
	__r, _, _ := purego.SyscallN(addr_QStackedLayout_replaceWidget_default, this.h, uintptr(from.UnsafePointer()), uintptr(to.UnsafePointer()), uintptr(options))
	runtime.KeepAlive(this)
	runtime.KeepAlive(from)
	runtime.KeepAlive(to)
	return bindQLayoutItem(__r)
}

func (this *QStackedLayout) LayoutDefault() *QLayout {
	__r, _, _ := purego.SyscallN(addr_QStackedLayout_layout_default, this.h)
	runtime.KeepAlive(this)
	return bindQLayout(__r)
}

func (this *QStackedLayout) ChildEventDefault(e *QChildEvent) {
	purego.SyscallN(addr_QStackedLayout_childEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QStackedLayout) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QStackedLayout_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QStackedLayout) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QStackedLayout_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QStackedLayout) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QStackedLayout_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QStackedLayout) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QStackedLayout_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QStackedLayout) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QStackedLayout_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QStackedLayout) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QStackedLayout_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QStackedLayout) MinimumHeightForWidthDefault(arg1 int) int {
	__r0, _, _ := purego.SyscallN(addr_QStackedLayout_minimumHeightForWidth_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QStackedLayout) SpacerItemDefault() *QSpacerItem {
	__r, _, _ := purego.SyscallN(addr_QStackedLayout_spacerItem_default, this.h)
	runtime.KeepAlive(this)
	return bindQSpacerItem(__r)
}

func (this *QStackedLayout) OnWidgetRemoved(fn func(int)) int64 {
	id := registerCallback(func(index int32) { fn(int(index)) })
	pg_QStackedLayout_connect_widgetRemoved(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStackedLayout) OnCurrentChanged(fn func(int)) int64 {
	id := registerCallback(func(index int32) { fn(int(index)) })
	pg_QStackedLayout_connect_currentChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStackedLayout) OnWidgetAdded(fn func(int)) int64 {
	id := registerCallback(func(index int32) { fn(int(index)) })
	pg_QStackedLayout_connect_widgetAdded(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStackedLayout) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QStackedLayout_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStackedLayout) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QStackedLayout_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStackedLayout) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QStackedLayout_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStackedLayout) Disconnect(handle int64) {
	pg_QStackedLayout__disconnect(handle)
	unregisterCallback(handle)
}

func QStackedLayout_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case -5627843058667511180:
		return 1
	case 6884465960736393867:
		return 2
	case 729690261457684101:
		return 3
	case -7246748916541192690:
		return 4
	case 7508512214013391107:
		return 5
	case 7549889108533486643:
		return 6
	case 5684820848726969227:
		return 7
	case -8266762397995423915:
		return 8
	case 4311534060495737339:
		return 9
	case -4587105490415799682:
		return 10
	case 4039479028406881496:
		return 11
	case 6991341317990632368:
		return 12
	case 1050490530495513297:
		return 13
	case 415665469015095569:
		return 14
	case 4158099606088301136:
		return 15
	case -7258473138030329570:
		return 16
	case -3276435765994963516:
		return 17
	case -626822430251098498:
		return 18
	case 3278947275714896887:
		return 19
	case -712862742159708061:
		return 20
	case -7183390618627992737:
		return 21
	case -1185315367526857833:
		return 22
	case 1237236456312824991:
		return 23
	case 5217623514630858297:
		return 24
	case -5409927581812374498:
		return 25
	case 3769868229123868474:
		return 26
	case 7379029982170503008:
		return 27
	case -568369772982215962:
		return 28
	case 5037837229528203129:
		return 29
	case 2272240669359776278:
		return 30
	}
	return -1
}

func QStackedLayout_enable(h uintptr, idx int32) { pg_QStackedLayout__set_override(h, idx) }

func (this *QStackedLayout) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QStackedLayout) OnCount(fn func() int) {
	registerVirt(this.UnsafeGoID(), -5627843058667511180, func() int32 { return int32(fn()) }) // slot "count"
}

func (this *QStackedLayout) OnAddItem(fn func(*QLayoutItem)) {
	registerVirt(this.UnsafeGoID(), 6884465960736393867, func(item uintptr) { fn(UnsafeNewQLayoutItem(unsafe.Pointer(item))) }) // slot "addItem"
}

func (this *QStackedLayout) OnSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 729690261457684101, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sizeHint"
}

func (this *QStackedLayout) OnMinimumSize(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), -7246748916541192690, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "minimumSize"
}

func (this *QStackedLayout) OnItemAt(fn func(int) *QLayoutItem) {
	registerVirt(this.UnsafeGoID(), 7508512214013391107, func(arg1 int32) uintptr { return uintptr(fn(int(arg1)).UnsafePointer()) }) // slot "itemAt"
}

func (this *QStackedLayout) OnTakeAt(fn func(int) *QLayoutItem) {
	registerVirt(this.UnsafeGoID(), 7549889108533486643, func(arg1 int32) uintptr { return uintptr(fn(int(arg1)).UnsafePointer()) }) // slot "takeAt"
}

func (this *QStackedLayout) OnSetGeometry(fn func(*QRect)) {
	registerVirt(this.UnsafeGoID(), 5684820848726969227, func(rect uintptr) { fn(UnsafeNewQRect(unsafe.Pointer(rect))) }) // slot "setGeometry"
}

func (this *QStackedLayout) OnHasHeightForWidth(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -8266762397995423915, func() int32 { return b2i(fn()) }) // slot "hasHeightForWidth"
}

func (this *QStackedLayout) OnHeightForWidth(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), 4311534060495737339, func(width int32) int32 { return int32(fn(int(width))) }) // slot "heightForWidth"
}

func (this *QStackedLayout) OnSpacing(fn func() int) {
	registerVirt(this.UnsafeGoID(), -4587105490415799682, func() int32 { return int32(fn()) }) // slot "spacing"
}

func (this *QStackedLayout) OnSetSpacing(fn func(int)) {
	registerVirt(this.UnsafeGoID(), 4039479028406881496, func(arg1 int32) { fn(int(arg1)) }) // slot "setSpacing"
}

func (this *QStackedLayout) OnInvalidate(fn func()) {
	registerVirt(this.UnsafeGoID(), 6991341317990632368, func() { fn() }) // slot "invalidate"
}

func (this *QStackedLayout) OnGeometry(fn func() *QRect) {
	registerVirt(this.UnsafeGoID(), 1050490530495513297, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "geometry"
}

func (this *QStackedLayout) OnExpandingDirections(fn func() Orientation) {
	registerVirt(this.UnsafeGoID(), 415665469015095569, func() int32 { return int32(fn()) }) // slot "expandingDirections"
}

func (this *QStackedLayout) OnMaximumSize(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 4158099606088301136, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "maximumSize"
}

func (this *QStackedLayout) OnIndexOf(fn func(*QWidget) int) {
	registerVirt(this.UnsafeGoID(), -7258473138030329570, func(arg1 uintptr) int32 { return int32(fn(UnsafeNewQWidget(unsafe.Pointer(arg1)))) }) // slot "indexOf"
}

func (this *QStackedLayout) OnIndexOf2(fn func(*QLayoutItem) int) {
	registerVirt(this.UnsafeGoID(), -3276435765994963516, func(arg1 uintptr) int32 { return int32(fn(UnsafeNewQLayoutItem(unsafe.Pointer(arg1)))) }) // slot "indexOf2"
}

func (this *QStackedLayout) OnIsEmpty(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -626822430251098498, func() int32 { return b2i(fn()) }) // slot "isEmpty"
}

func (this *QStackedLayout) OnControlTypes(fn func() QSizePolicy__ControlType) {
	registerVirt(this.UnsafeGoID(), 3278947275714896887, func() int32 { return int32(fn()) }) // slot "controlTypes"
}

func (this *QStackedLayout) OnReplaceWidget(fn func(*QWidget, *QWidget, FindChildOption) *QLayoutItem) {
	registerVirt(this.UnsafeGoID(), -712862742159708061, func(from uintptr, to uintptr, options int32) uintptr { return uintptr(fn(UnsafeNewQWidget(unsafe.Pointer(from)), UnsafeNewQWidget(unsafe.Pointer(to)), FindChildOption(options)).UnsafePointer()) }) // slot "replaceWidget"
}

func (this *QStackedLayout) OnLayout(fn func() *QLayout) {
	registerVirt(this.UnsafeGoID(), -7183390618627992737, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "layout"
}

func (this *QStackedLayout) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(e uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(e))) }) // slot "childEvent"
}

func (this *QStackedLayout) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QStackedLayout) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QStackedLayout) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QStackedLayout) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QStackedLayout) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QStackedLayout) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QStackedLayout) OnMinimumHeightForWidth(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), 5037837229528203129, func(arg1 int32) int32 { return int32(fn(int(arg1))) }) // slot "minimumHeightForWidth"
}

func (this *QStackedLayout) OnSpacerItem(fn func() *QSpacerItem) {
	registerVirt(this.UnsafeGoID(), 2272240669359776278, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "spacerItem"
}

func (this *QStackedLayout) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QStackedLayout_delete(this.h)
	}
	this.UnsafeRelease()
}
