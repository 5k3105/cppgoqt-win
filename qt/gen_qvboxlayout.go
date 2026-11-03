package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QVBoxLayout_chain uintptr
	addr_QVBoxLayout_new uintptr
	addr_QVBoxLayout_new2 uintptr
	addr_QVBoxLayout_tr uintptr
	addr_QVBoxLayout_closestAcceptableSize uintptr
	addr_QVBoxLayout_widgetEvent uintptr
	addr_QVBoxLayout_addChildLayout uintptr
	addr_QVBoxLayout_addChildWidget uintptr
	addr_QVBoxLayout_adoptLayout uintptr
	addr_QVBoxLayout_alignmentRect uintptr
	addr_QVBoxLayout_connect uintptr
	addr_QVBoxLayout_connect2 uintptr
	addr_QVBoxLayout_sender uintptr
	addr_QVBoxLayout_senderSignalIndex uintptr
	addr_QVBoxLayout_receivers uintptr
	addr_QVBoxLayout_isSignalConnected uintptr
	addr_QVBoxLayout_tr2 uintptr
	addr_QVBoxLayout_tr3 uintptr
	addr_QVBoxLayout_connect3 uintptr
	addr_QVBoxLayout_connect4 uintptr
	addr_QVBoxLayout_metaObject_default uintptr
	addr_QVBoxLayout_addItem_default uintptr
	addr_QVBoxLayout_spacing_default uintptr
	addr_QVBoxLayout_setSpacing_default uintptr
	addr_QVBoxLayout_sizeHint_default uintptr
	addr_QVBoxLayout_minimumSize_default uintptr
	addr_QVBoxLayout_maximumSize_default uintptr
	addr_QVBoxLayout_hasHeightForWidth_default uintptr
	addr_QVBoxLayout_heightForWidth_default uintptr
	addr_QVBoxLayout_minimumHeightForWidth_default uintptr
	addr_QVBoxLayout_expandingDirections_default uintptr
	addr_QVBoxLayout_invalidate_default uintptr
	addr_QVBoxLayout_itemAt_default uintptr
	addr_QVBoxLayout_takeAt_default uintptr
	addr_QVBoxLayout_count_default uintptr
	addr_QVBoxLayout_setGeometry_default uintptr
	addr_QVBoxLayout_geometry_default uintptr
	addr_QVBoxLayout_indexOf_default uintptr
	addr_QVBoxLayout_isEmpty_default uintptr
	addr_QVBoxLayout_controlTypes_default uintptr
	addr_QVBoxLayout_replaceWidget_default uintptr
	addr_QVBoxLayout_layout_default uintptr
	addr_QVBoxLayout_childEvent_default uintptr
	addr_QVBoxLayout_event_default uintptr
	addr_QVBoxLayout_eventFilter_default uintptr
	addr_QVBoxLayout_timerEvent_default uintptr
	addr_QVBoxLayout_customEvent_default uintptr
	addr_QVBoxLayout_connectNotify_default uintptr
	addr_QVBoxLayout_disconnectNotify_default uintptr
	addr_QVBoxLayout_widget_default uintptr
	addr_QVBoxLayout_spacerItem_default uintptr
	pg_QVBoxLayout_connect_destroyed func(uintptr, int64)
	pg_QVBoxLayout_connect_destroyed2 func(uintptr, int64)
	pg_QVBoxLayout_connect_objectNameChanged func(uintptr, int64)
	pg_QVBoxLayout__disconnect func(int64)
	pg_QVBoxLayout__set_override func(uintptr, int32)
	pg_QVBoxLayout_delete func(uintptr)
)

func init() {
	symaddr(&addr_QVBoxLayout_chain, "QVBoxLayout_chain")
	symaddr(&addr_QVBoxLayout_new, "QVBoxLayout_new")
	symaddr(&addr_QVBoxLayout_new2, "QVBoxLayout_new2")
	symaddr(&addr_QVBoxLayout_tr, "QVBoxLayout_tr")
	symaddr(&addr_QVBoxLayout_closestAcceptableSize, "QVBoxLayout_closestAcceptableSize")
	symaddr(&addr_QVBoxLayout_widgetEvent, "QVBoxLayout_widgetEvent")
	symaddr(&addr_QVBoxLayout_addChildLayout, "QVBoxLayout_addChildLayout")
	symaddr(&addr_QVBoxLayout_addChildWidget, "QVBoxLayout_addChildWidget")
	symaddr(&addr_QVBoxLayout_adoptLayout, "QVBoxLayout_adoptLayout")
	symaddr(&addr_QVBoxLayout_alignmentRect, "QVBoxLayout_alignmentRect")
	symaddr(&addr_QVBoxLayout_connect, "QVBoxLayout_connect")
	symaddr(&addr_QVBoxLayout_connect2, "QVBoxLayout_connect2")
	symaddr(&addr_QVBoxLayout_sender, "QVBoxLayout_sender")
	symaddr(&addr_QVBoxLayout_senderSignalIndex, "QVBoxLayout_senderSignalIndex")
	symaddr(&addr_QVBoxLayout_receivers, "QVBoxLayout_receivers")
	symaddr(&addr_QVBoxLayout_isSignalConnected, "QVBoxLayout_isSignalConnected")
	symaddr(&addr_QVBoxLayout_tr2, "QVBoxLayout_tr2")
	symaddr(&addr_QVBoxLayout_tr3, "QVBoxLayout_tr3")
	symaddr(&addr_QVBoxLayout_connect3, "QVBoxLayout_connect3")
	symaddr(&addr_QVBoxLayout_connect4, "QVBoxLayout_connect4")
	symaddr(&addr_QVBoxLayout_metaObject_default, "QVBoxLayout_metaObject_default")
	symaddr(&addr_QVBoxLayout_addItem_default, "QVBoxLayout_addItem_default")
	symaddr(&addr_QVBoxLayout_spacing_default, "QVBoxLayout_spacing_default")
	symaddr(&addr_QVBoxLayout_setSpacing_default, "QVBoxLayout_setSpacing_default")
	symaddr(&addr_QVBoxLayout_sizeHint_default, "QVBoxLayout_sizeHint_default")
	symaddr(&addr_QVBoxLayout_minimumSize_default, "QVBoxLayout_minimumSize_default")
	symaddr(&addr_QVBoxLayout_maximumSize_default, "QVBoxLayout_maximumSize_default")
	symaddr(&addr_QVBoxLayout_hasHeightForWidth_default, "QVBoxLayout_hasHeightForWidth_default")
	symaddr(&addr_QVBoxLayout_heightForWidth_default, "QVBoxLayout_heightForWidth_default")
	symaddr(&addr_QVBoxLayout_minimumHeightForWidth_default, "QVBoxLayout_minimumHeightForWidth_default")
	symaddr(&addr_QVBoxLayout_expandingDirections_default, "QVBoxLayout_expandingDirections_default")
	symaddr(&addr_QVBoxLayout_invalidate_default, "QVBoxLayout_invalidate_default")
	symaddr(&addr_QVBoxLayout_itemAt_default, "QVBoxLayout_itemAt_default")
	symaddr(&addr_QVBoxLayout_takeAt_default, "QVBoxLayout_takeAt_default")
	symaddr(&addr_QVBoxLayout_count_default, "QVBoxLayout_count_default")
	symaddr(&addr_QVBoxLayout_setGeometry_default, "QVBoxLayout_setGeometry_default")
	symaddr(&addr_QVBoxLayout_geometry_default, "QVBoxLayout_geometry_default")
	symaddr(&addr_QVBoxLayout_indexOf_default, "QVBoxLayout_indexOf_default")
	symaddr(&addr_QVBoxLayout_isEmpty_default, "QVBoxLayout_isEmpty_default")
	symaddr(&addr_QVBoxLayout_controlTypes_default, "QVBoxLayout_controlTypes_default")
	symaddr(&addr_QVBoxLayout_replaceWidget_default, "QVBoxLayout_replaceWidget_default")
	symaddr(&addr_QVBoxLayout_layout_default, "QVBoxLayout_layout_default")
	symaddr(&addr_QVBoxLayout_childEvent_default, "QVBoxLayout_childEvent_default")
	symaddr(&addr_QVBoxLayout_event_default, "QVBoxLayout_event_default")
	symaddr(&addr_QVBoxLayout_eventFilter_default, "QVBoxLayout_eventFilter_default")
	symaddr(&addr_QVBoxLayout_timerEvent_default, "QVBoxLayout_timerEvent_default")
	symaddr(&addr_QVBoxLayout_customEvent_default, "QVBoxLayout_customEvent_default")
	symaddr(&addr_QVBoxLayout_connectNotify_default, "QVBoxLayout_connectNotify_default")
	symaddr(&addr_QVBoxLayout_disconnectNotify_default, "QVBoxLayout_disconnectNotify_default")
	symaddr(&addr_QVBoxLayout_widget_default, "QVBoxLayout_widget_default")
	symaddr(&addr_QVBoxLayout_spacerItem_default, "QVBoxLayout_spacerItem_default")
	reg(&pg_QVBoxLayout_connect_destroyed, "QVBoxLayout_connect_destroyed")
	reg(&pg_QVBoxLayout_connect_destroyed2, "QVBoxLayout_connect_destroyed2")
	reg(&pg_QVBoxLayout_connect_objectNameChanged, "QVBoxLayout_connect_objectNameChanged")
	reg(&pg_QVBoxLayout__disconnect, "QVBoxLayout__disconnect")
	reg(&pg_QVBoxLayout__set_override, "QVBoxLayout__set_override")
	reg(&pg_QVBoxLayout_delete, "QVBoxLayout_delete")
}

type QVBoxLayout struct {
	h uintptr
	*QBoxLayout
}

func bindQVBoxLayout(h uintptr) *QVBoxLayout {
	if h == 0 {
		return nil
	}
	var __c [3]uintptr
	purego.SyscallN(addr_QVBoxLayout_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QVBoxLayout{h: h, QBoxLayout: &QBoxLayout{h: __c[0], QLayout: &QLayout{h: __c[1], QObject: &QObject{h: __c[2]}}}}
}

func newownedQVBoxLayout(h uintptr) *QVBoxLayout { return bindQVBoxLayout(h) }

func (this *QVBoxLayout) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QVBoxLayout) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQVBoxLayout(h unsafe.Pointer) *QVBoxLayout { return bindQVBoxLayout(uintptr(h)) }
func UnsafeNewQVBoxLayoutOwned(h unsafe.Pointer) *QVBoxLayout { return newownedQVBoxLayout(uintptr(h)) }

func (this *QVBoxLayout) UnsafeRelease() {
	this.h = 0
	if this.QBoxLayout != nil {
		this.QBoxLayout.UnsafeRelease()
	}
}

func NewQVBoxLayout(parent *QWidget) *QVBoxLayout {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QVBoxLayout_new, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQVBoxLayout(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QVBoxLayout_virtIdx, QVBoxLayout_enable)
	return __o
}

func NewQVBoxLayout2() *QVBoxLayout {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QVBoxLayout_new2, uintptr(__goid))
	__o := newownedQVBoxLayout(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QVBoxLayout_virtIdx, QVBoxLayout_enable)
	return __o
}

func QVBoxLayout_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QVBoxLayout_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func QVBoxLayout_ClosestAcceptableSize(w *QWidget, s *QSize) *QSize {
	__r, _, _ := purego.SyscallN(addr_QVBoxLayout_closestAcceptableSize, uintptr(w.UnsafePointer()), uintptr(s.UnsafePointer()))
	runtime.KeepAlive(w)
	runtime.KeepAlive(s)
	return newownedQSize(__r)
}

func (this *QVBoxLayout) WidgetEvent(arg1 *QEvent) {
	purego.SyscallN(addr_QVBoxLayout_widgetEvent, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QVBoxLayout) AddChildLayout(l *QLayout) {
	purego.SyscallN(addr_QVBoxLayout_addChildLayout, this.h, uintptr(l.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(l)
}

func (this *QVBoxLayout) AddChildWidget(w *QWidget) {
	purego.SyscallN(addr_QVBoxLayout_addChildWidget, this.h, uintptr(w.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(w)
}

func (this *QVBoxLayout) AdoptLayout(layout *QLayout) bool {
	__r0, _, _ := purego.SyscallN(addr_QVBoxLayout_adoptLayout, this.h, uintptr(layout.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(layout)
	return __r != 0
}

func (this *QVBoxLayout) AlignmentRect(arg1 *QRect) *QRect {
	__r, _, _ := purego.SyscallN(addr_QVBoxLayout_alignmentRect, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return newownedQRect(__r)
}

func QVBoxLayout_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QVBoxLayout_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QVBoxLayout_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QVBoxLayout_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QVBoxLayout) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QVBoxLayout_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QVBoxLayout) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QVBoxLayout_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QVBoxLayout) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QVBoxLayout_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QVBoxLayout) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QVBoxLayout_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QVBoxLayout_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QVBoxLayout_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QVBoxLayout_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QVBoxLayout_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QVBoxLayout_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QVBoxLayout_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QVBoxLayout_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QVBoxLayout_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QVBoxLayout) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QVBoxLayout_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QVBoxLayout) AddItemDefault(arg1 *QLayoutItem) {
	purego.SyscallN(addr_QVBoxLayout_addItem_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QVBoxLayout) SpacingDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QVBoxLayout_spacing_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QVBoxLayout) SetSpacingDefault(spacing int) {
	purego.SyscallN(addr_QVBoxLayout_setSpacing_default, this.h, uintptr(spacing))
	runtime.KeepAlive(this)
}

func (this *QVBoxLayout) SizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QVBoxLayout_sizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QVBoxLayout) MinimumSizeDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QVBoxLayout_minimumSize_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QVBoxLayout) MaximumSizeDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QVBoxLayout_maximumSize_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QVBoxLayout) HasHeightForWidthDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QVBoxLayout_hasHeightForWidth_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QVBoxLayout) HeightForWidthDefault(arg1 int) int {
	__r0, _, _ := purego.SyscallN(addr_QVBoxLayout_heightForWidth_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QVBoxLayout) MinimumHeightForWidthDefault(arg1 int) int {
	__r0, _, _ := purego.SyscallN(addr_QVBoxLayout_minimumHeightForWidth_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QVBoxLayout) ExpandingDirectionsDefault() Orientation {
	__r0, _, _ := purego.SyscallN(addr_QVBoxLayout_expandingDirections_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return Orientation(__r)
}

func (this *QVBoxLayout) InvalidateDefault() {
	purego.SyscallN(addr_QVBoxLayout_invalidate_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QVBoxLayout) ItemAtDefault(arg1 int) *QLayoutItem {
	__r, _, _ := purego.SyscallN(addr_QVBoxLayout_itemAt_default, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return bindQLayoutItem(__r)
}

func (this *QVBoxLayout) TakeAtDefault(arg1 int) *QLayoutItem {
	__r, _, _ := purego.SyscallN(addr_QVBoxLayout_takeAt_default, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return bindQLayoutItem(__r)
}

func (this *QVBoxLayout) CountDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QVBoxLayout_count_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QVBoxLayout) SetGeometryDefault(arg1 *QRect) {
	purego.SyscallN(addr_QVBoxLayout_setGeometry_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QVBoxLayout) GeometryDefault() *QRect {
	__r, _, _ := purego.SyscallN(addr_QVBoxLayout_geometry_default, this.h)
	runtime.KeepAlive(this)
	return newownedQRect(__r)
}

func (this *QVBoxLayout) IndexOfDefault(arg1 *QWidget) int {
	__r0, _, _ := purego.SyscallN(addr_QVBoxLayout_indexOf_default, this.h, uintptr(arg1.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return int(__r)
}

func (this *QVBoxLayout) IsEmptyDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QVBoxLayout_isEmpty_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QVBoxLayout) ControlTypesDefault() QSizePolicy__ControlType {
	__r0, _, _ := purego.SyscallN(addr_QVBoxLayout_controlTypes_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QSizePolicy__ControlType(__r)
}

func (this *QVBoxLayout) ReplaceWidgetDefault(from *QWidget, to *QWidget, options FindChildOption) *QLayoutItem {
	__r, _, _ := purego.SyscallN(addr_QVBoxLayout_replaceWidget_default, this.h, uintptr(from.UnsafePointer()), uintptr(to.UnsafePointer()), uintptr(options))
	runtime.KeepAlive(this)
	runtime.KeepAlive(from)
	runtime.KeepAlive(to)
	return bindQLayoutItem(__r)
}

func (this *QVBoxLayout) LayoutDefault() *QLayout {
	__r, _, _ := purego.SyscallN(addr_QVBoxLayout_layout_default, this.h)
	runtime.KeepAlive(this)
	return bindQLayout(__r)
}

func (this *QVBoxLayout) ChildEventDefault(e *QChildEvent) {
	purego.SyscallN(addr_QVBoxLayout_childEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QVBoxLayout) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QVBoxLayout_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QVBoxLayout) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QVBoxLayout_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QVBoxLayout) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QVBoxLayout_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QVBoxLayout) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QVBoxLayout_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QVBoxLayout) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QVBoxLayout_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QVBoxLayout) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QVBoxLayout_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QVBoxLayout) WidgetDefault() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QVBoxLayout_widget_default, this.h)
	runtime.KeepAlive(this)
	return bindQWidget(__r)
}

func (this *QVBoxLayout) SpacerItemDefault() *QSpacerItem {
	__r, _, _ := purego.SyscallN(addr_QVBoxLayout_spacerItem_default, this.h)
	runtime.KeepAlive(this)
	return bindQSpacerItem(__r)
}

func (this *QVBoxLayout) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QVBoxLayout_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QVBoxLayout) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QVBoxLayout_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QVBoxLayout) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QVBoxLayout_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QVBoxLayout) Disconnect(handle int64) {
	pg_QVBoxLayout__disconnect(handle)
	unregisterCallback(handle)
}

func QVBoxLayout_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case 6884465960736393867:
		return 1
	case -4587105490415799682:
		return 2
	case 4039479028406881496:
		return 3
	case 729690261457684101:
		return 4
	case -7246748916541192690:
		return 5
	case 4158099606088301136:
		return 6
	case -8266762397995423915:
		return 7
	case 4311534060495737339:
		return 8
	case 5037837229528203129:
		return 9
	case 415665469015095569:
		return 10
	case 6991341317990632368:
		return 11
	case 7508512214013391107:
		return 12
	case 7549889108533486643:
		return 13
	case -5627843058667511180:
		return 14
	case 5684820848726969227:
		return 15
	case 1050490530495513297:
		return 16
	case -7258473138030329570:
		return 17
	case -3276435765994963516:
		return 18
	case -626822430251098498:
		return 19
	case 3278947275714896887:
		return 20
	case -712862742159708061:
		return 21
	case -7183390618627992737:
		return 22
	case -1185315367526857833:
		return 23
	case 1237236456312824991:
		return 24
	case 5217623514630858297:
		return 25
	case -5409927581812374498:
		return 26
	case 3769868229123868474:
		return 27
	case 7379029982170503008:
		return 28
	case -568369772982215962:
		return 29
	case -4575937485689502171:
		return 30
	case 2272240669359776278:
		return 31
	}
	return -1
}

func QVBoxLayout_enable(h uintptr, idx int32) { pg_QVBoxLayout__set_override(h, idx) }

func (this *QVBoxLayout) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QVBoxLayout) OnAddItem(fn func(*QLayoutItem)) {
	registerVirt(this.UnsafeGoID(), 6884465960736393867, func(arg1 uintptr) { fn(UnsafeNewQLayoutItem(unsafe.Pointer(arg1))) }) // slot "addItem"
}

func (this *QVBoxLayout) OnSpacing(fn func() int) {
	registerVirt(this.UnsafeGoID(), -4587105490415799682, func() int32 { return int32(fn()) }) // slot "spacing"
}

func (this *QVBoxLayout) OnSetSpacing(fn func(int)) {
	registerVirt(this.UnsafeGoID(), 4039479028406881496, func(spacing int32) { fn(int(spacing)) }) // slot "setSpacing"
}

func (this *QVBoxLayout) OnSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 729690261457684101, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sizeHint"
}

func (this *QVBoxLayout) OnMinimumSize(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), -7246748916541192690, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "minimumSize"
}

func (this *QVBoxLayout) OnMaximumSize(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 4158099606088301136, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "maximumSize"
}

func (this *QVBoxLayout) OnHasHeightForWidth(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -8266762397995423915, func() int32 { return b2i(fn()) }) // slot "hasHeightForWidth"
}

func (this *QVBoxLayout) OnHeightForWidth(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), 4311534060495737339, func(arg1 int32) int32 { return int32(fn(int(arg1))) }) // slot "heightForWidth"
}

func (this *QVBoxLayout) OnMinimumHeightForWidth(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), 5037837229528203129, func(arg1 int32) int32 { return int32(fn(int(arg1))) }) // slot "minimumHeightForWidth"
}

func (this *QVBoxLayout) OnExpandingDirections(fn func() Orientation) {
	registerVirt(this.UnsafeGoID(), 415665469015095569, func() int32 { return int32(fn()) }) // slot "expandingDirections"
}

func (this *QVBoxLayout) OnInvalidate(fn func()) {
	registerVirt(this.UnsafeGoID(), 6991341317990632368, func() { fn() }) // slot "invalidate"
}

func (this *QVBoxLayout) OnItemAt(fn func(int) *QLayoutItem) {
	registerVirt(this.UnsafeGoID(), 7508512214013391107, func(arg1 int32) uintptr { return uintptr(fn(int(arg1)).UnsafePointer()) }) // slot "itemAt"
}

func (this *QVBoxLayout) OnTakeAt(fn func(int) *QLayoutItem) {
	registerVirt(this.UnsafeGoID(), 7549889108533486643, func(arg1 int32) uintptr { return uintptr(fn(int(arg1)).UnsafePointer()) }) // slot "takeAt"
}

func (this *QVBoxLayout) OnCount(fn func() int) {
	registerVirt(this.UnsafeGoID(), -5627843058667511180, func() int32 { return int32(fn()) }) // slot "count"
}

func (this *QVBoxLayout) OnSetGeometry(fn func(*QRect)) {
	registerVirt(this.UnsafeGoID(), 5684820848726969227, func(arg1 uintptr) { fn(UnsafeNewQRect(unsafe.Pointer(arg1))) }) // slot "setGeometry"
}

func (this *QVBoxLayout) OnGeometry(fn func() *QRect) {
	registerVirt(this.UnsafeGoID(), 1050490530495513297, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "geometry"
}

func (this *QVBoxLayout) OnIndexOf(fn func(*QWidget) int) {
	registerVirt(this.UnsafeGoID(), -7258473138030329570, func(arg1 uintptr) int32 { return int32(fn(UnsafeNewQWidget(unsafe.Pointer(arg1)))) }) // slot "indexOf"
}

func (this *QVBoxLayout) OnIndexOf2(fn func(*QLayoutItem) int) {
	registerVirt(this.UnsafeGoID(), -3276435765994963516, func(arg1 uintptr) int32 { return int32(fn(UnsafeNewQLayoutItem(unsafe.Pointer(arg1)))) }) // slot "indexOf2"
}

func (this *QVBoxLayout) OnIsEmpty(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -626822430251098498, func() int32 { return b2i(fn()) }) // slot "isEmpty"
}

func (this *QVBoxLayout) OnControlTypes(fn func() QSizePolicy__ControlType) {
	registerVirt(this.UnsafeGoID(), 3278947275714896887, func() int32 { return int32(fn()) }) // slot "controlTypes"
}

func (this *QVBoxLayout) OnReplaceWidget(fn func(*QWidget, *QWidget, FindChildOption) *QLayoutItem) {
	registerVirt(this.UnsafeGoID(), -712862742159708061, func(from uintptr, to uintptr, options int32) uintptr { return uintptr(fn(UnsafeNewQWidget(unsafe.Pointer(from)), UnsafeNewQWidget(unsafe.Pointer(to)), FindChildOption(options)).UnsafePointer()) }) // slot "replaceWidget"
}

func (this *QVBoxLayout) OnLayout(fn func() *QLayout) {
	registerVirt(this.UnsafeGoID(), -7183390618627992737, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "layout"
}

func (this *QVBoxLayout) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(e uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(e))) }) // slot "childEvent"
}

func (this *QVBoxLayout) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QVBoxLayout) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QVBoxLayout) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QVBoxLayout) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QVBoxLayout) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QVBoxLayout) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QVBoxLayout) OnWidget(fn func() *QWidget) {
	registerVirt(this.UnsafeGoID(), -4575937485689502171, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "widget"
}

func (this *QVBoxLayout) OnSpacerItem(fn func() *QSpacerItem) {
	registerVirt(this.UnsafeGoID(), 2272240669359776278, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "spacerItem"
}

func (this *QVBoxLayout) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QVBoxLayout_delete(this.h)
	}
	this.UnsafeRelease()
}
