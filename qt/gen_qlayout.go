package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QLayout_base func(uintptr) uintptr
	addr_QLayout_tr uintptr
	addr_QLayout_spacing uintptr
	addr_QLayout_setSpacing uintptr
	addr_QLayout_setContentsMargins uintptr
	addr_QLayout_setContentsMargins2 uintptr
	addr_QLayout_unsetContentsMargins uintptr
	addr_QLayout_getContentsMargins uintptr
	addr_QLayout_contentsMargins uintptr
	addr_QLayout_contentsRect uintptr
	addr_QLayout_setAlignment uintptr
	addr_QLayout_setAlignment2 uintptr
	addr_QLayout_setSizeConstraint uintptr
	addr_QLayout_sizeConstraint uintptr
	addr_QLayout_setSizeConstraints uintptr
	addr_QLayout_setHorizontalSizeConstraint uintptr
	addr_QLayout_horizontalSizeConstraint uintptr
	addr_QLayout_setVerticalSizeConstraint uintptr
	addr_QLayout_verticalSizeConstraint uintptr
	addr_QLayout_setMenuBar uintptr
	addr_QLayout_menuBar uintptr
	addr_QLayout_parentWidget uintptr
	addr_QLayout_invalidate uintptr
	addr_QLayout_geometry uintptr
	addr_QLayout_activate uintptr
	addr_QLayout_update uintptr
	addr_QLayout_addWidget uintptr
	addr_QLayout_addItem uintptr
	addr_QLayout_removeWidget uintptr
	addr_QLayout_removeItem uintptr
	addr_QLayout_expandingDirections uintptr
	addr_QLayout_minimumSize uintptr
	addr_QLayout_maximumSize uintptr
	addr_QLayout_setGeometry uintptr
	addr_QLayout_itemAt uintptr
	addr_QLayout_takeAt uintptr
	addr_QLayout_indexOf uintptr
	addr_QLayout_indexOf2 uintptr
	addr_QLayout_count uintptr
	addr_QLayout_isEmpty uintptr
	addr_QLayout_controlTypes uintptr
	addr_QLayout_replaceWidget uintptr
	addr_QLayout_totalMinimumHeightForWidth uintptr
	addr_QLayout_totalHeightForWidth uintptr
	addr_QLayout_totalMinimumSize uintptr
	addr_QLayout_totalMaximumSize uintptr
	addr_QLayout_totalSizeHint uintptr
	addr_QLayout_layout uintptr
	addr_QLayout_setEnabled uintptr
	addr_QLayout_isEnabled uintptr
	addr_QLayout_closestAcceptableSize uintptr
	addr_QLayout_widgetEvent uintptr
	addr_QLayout_addChildLayout uintptr
	addr_QLayout_addChildWidget uintptr
	addr_QLayout_adoptLayout uintptr
	addr_QLayout_alignmentRect uintptr
	addr_QLayout_connect uintptr
	addr_QLayout_connect2 uintptr
	addr_QLayout_sender uintptr
	addr_QLayout_senderSignalIndex uintptr
	addr_QLayout_receivers uintptr
	addr_QLayout_isSignalConnected uintptr
	addr_QLayout_sizeHint uintptr
	addr_QLayout_hasHeightForWidth uintptr
	addr_QLayout_heightForWidth uintptr
	addr_QLayout_minimumHeightForWidth uintptr
	addr_QLayout_widget uintptr
	addr_QLayout_spacerItem uintptr
	addr_QLayout_alignment uintptr
	addr_QLayout_tr2 uintptr
	addr_QLayout_tr3 uintptr
	addr_QLayout_connect3 uintptr
	addr_QLayout_connect4 uintptr
	addr_QLayout_asQLayoutItem uintptr
	addr_QLayout_spacing_default uintptr
	addr_QLayout_setSpacing_default uintptr
	addr_QLayout_invalidate_default uintptr
	addr_QLayout_geometry_default uintptr
	addr_QLayout_expandingDirections_default uintptr
	addr_QLayout_minimumSize_default uintptr
	addr_QLayout_maximumSize_default uintptr
	addr_QLayout_setGeometry_default uintptr
	addr_QLayout_indexOf_default uintptr
	addr_QLayout_indexOf_default2 uintptr
	addr_QLayout_isEmpty_default uintptr
	addr_QLayout_controlTypes_default uintptr
	addr_QLayout_replaceWidget_default uintptr
	addr_QLayout_layout_default uintptr
	addr_QLayout_hasHeightForWidth_default uintptr
	addr_QLayout_heightForWidth_default uintptr
	addr_QLayout_minimumHeightForWidth_default uintptr
	addr_QLayout_widget_default uintptr
	addr_QLayout_spacerItem_default uintptr
	addr_QLayout_metaObject_default uintptr
	addr_QLayout_childEvent_default uintptr
	addr_QLayout_event_default uintptr
	addr_QLayout_eventFilter_default uintptr
	addr_QLayout_timerEvent_default uintptr
	addr_QLayout_customEvent_default uintptr
	addr_QLayout_connectNotify_default uintptr
	addr_QLayout_disconnectNotify_default uintptr
	pg_QLayout_connect_destroyed func(uintptr, int64)
	pg_QLayout_connect_destroyed2 func(uintptr, int64)
	pg_QLayout_connect_objectNameChanged func(uintptr, int64)
	pg_QLayout__disconnect func(int64)
	pg_QLayout__set_override func(uintptr, int32)
	pg_QLayout__Go_new func(int64) uintptr
	pg_QLayout_delete func(uintptr)
)

func init() {
	reg(&pg_QLayout_base, "QLayout_base")
	symaddr(&addr_QLayout_tr, "QLayout_tr")
	symaddr(&addr_QLayout_spacing, "QLayout_spacing")
	symaddr(&addr_QLayout_setSpacing, "QLayout_setSpacing")
	symaddr(&addr_QLayout_setContentsMargins, "QLayout_setContentsMargins")
	symaddr(&addr_QLayout_setContentsMargins2, "QLayout_setContentsMargins2")
	symaddr(&addr_QLayout_unsetContentsMargins, "QLayout_unsetContentsMargins")
	symaddr(&addr_QLayout_getContentsMargins, "QLayout_getContentsMargins")
	symaddr(&addr_QLayout_contentsMargins, "QLayout_contentsMargins")
	symaddr(&addr_QLayout_contentsRect, "QLayout_contentsRect")
	symaddr(&addr_QLayout_setAlignment, "QLayout_setAlignment")
	symaddr(&addr_QLayout_setAlignment2, "QLayout_setAlignment2")
	symaddr(&addr_QLayout_setSizeConstraint, "QLayout_setSizeConstraint")
	symaddr(&addr_QLayout_sizeConstraint, "QLayout_sizeConstraint")
	symaddr(&addr_QLayout_setSizeConstraints, "QLayout_setSizeConstraints")
	symaddr(&addr_QLayout_setHorizontalSizeConstraint, "QLayout_setHorizontalSizeConstraint")
	symaddr(&addr_QLayout_horizontalSizeConstraint, "QLayout_horizontalSizeConstraint")
	symaddr(&addr_QLayout_setVerticalSizeConstraint, "QLayout_setVerticalSizeConstraint")
	symaddr(&addr_QLayout_verticalSizeConstraint, "QLayout_verticalSizeConstraint")
	symaddr(&addr_QLayout_setMenuBar, "QLayout_setMenuBar")
	symaddr(&addr_QLayout_menuBar, "QLayout_menuBar")
	symaddr(&addr_QLayout_parentWidget, "QLayout_parentWidget")
	symaddr(&addr_QLayout_invalidate, "QLayout_invalidate")
	symaddr(&addr_QLayout_geometry, "QLayout_geometry")
	symaddr(&addr_QLayout_activate, "QLayout_activate")
	symaddr(&addr_QLayout_update, "QLayout_update")
	symaddr(&addr_QLayout_addWidget, "QLayout_addWidget")
	symaddr(&addr_QLayout_addItem, "QLayout_addItem")
	symaddr(&addr_QLayout_removeWidget, "QLayout_removeWidget")
	symaddr(&addr_QLayout_removeItem, "QLayout_removeItem")
	symaddr(&addr_QLayout_expandingDirections, "QLayout_expandingDirections")
	symaddr(&addr_QLayout_minimumSize, "QLayout_minimumSize")
	symaddr(&addr_QLayout_maximumSize, "QLayout_maximumSize")
	symaddr(&addr_QLayout_setGeometry, "QLayout_setGeometry")
	symaddr(&addr_QLayout_itemAt, "QLayout_itemAt")
	symaddr(&addr_QLayout_takeAt, "QLayout_takeAt")
	symaddr(&addr_QLayout_indexOf, "QLayout_indexOf")
	symaddr(&addr_QLayout_indexOf2, "QLayout_indexOf2")
	symaddr(&addr_QLayout_count, "QLayout_count")
	symaddr(&addr_QLayout_isEmpty, "QLayout_isEmpty")
	symaddr(&addr_QLayout_controlTypes, "QLayout_controlTypes")
	symaddr(&addr_QLayout_replaceWidget, "QLayout_replaceWidget")
	symaddr(&addr_QLayout_totalMinimumHeightForWidth, "QLayout_totalMinimumHeightForWidth")
	symaddr(&addr_QLayout_totalHeightForWidth, "QLayout_totalHeightForWidth")
	symaddr(&addr_QLayout_totalMinimumSize, "QLayout_totalMinimumSize")
	symaddr(&addr_QLayout_totalMaximumSize, "QLayout_totalMaximumSize")
	symaddr(&addr_QLayout_totalSizeHint, "QLayout_totalSizeHint")
	symaddr(&addr_QLayout_layout, "QLayout_layout")
	symaddr(&addr_QLayout_setEnabled, "QLayout_setEnabled")
	symaddr(&addr_QLayout_isEnabled, "QLayout_isEnabled")
	symaddr(&addr_QLayout_closestAcceptableSize, "QLayout_closestAcceptableSize")
	symaddr(&addr_QLayout_widgetEvent, "QLayout_widgetEvent")
	symaddr(&addr_QLayout_addChildLayout, "QLayout_addChildLayout")
	symaddr(&addr_QLayout_addChildWidget, "QLayout_addChildWidget")
	symaddr(&addr_QLayout_adoptLayout, "QLayout_adoptLayout")
	symaddr(&addr_QLayout_alignmentRect, "QLayout_alignmentRect")
	symaddr(&addr_QLayout_connect, "QLayout_connect")
	symaddr(&addr_QLayout_connect2, "QLayout_connect2")
	symaddr(&addr_QLayout_sender, "QLayout_sender")
	symaddr(&addr_QLayout_senderSignalIndex, "QLayout_senderSignalIndex")
	symaddr(&addr_QLayout_receivers, "QLayout_receivers")
	symaddr(&addr_QLayout_isSignalConnected, "QLayout_isSignalConnected")
	symaddr(&addr_QLayout_sizeHint, "QLayout_sizeHint")
	symaddr(&addr_QLayout_hasHeightForWidth, "QLayout_hasHeightForWidth")
	symaddr(&addr_QLayout_heightForWidth, "QLayout_heightForWidth")
	symaddr(&addr_QLayout_minimumHeightForWidth, "QLayout_minimumHeightForWidth")
	symaddr(&addr_QLayout_widget, "QLayout_widget")
	symaddr(&addr_QLayout_spacerItem, "QLayout_spacerItem")
	symaddr(&addr_QLayout_alignment, "QLayout_alignment")
	symaddr(&addr_QLayout_tr2, "QLayout_tr2")
	symaddr(&addr_QLayout_tr3, "QLayout_tr3")
	symaddr(&addr_QLayout_connect3, "QLayout_connect3")
	symaddr(&addr_QLayout_connect4, "QLayout_connect4")
	symaddr(&addr_QLayout_asQLayoutItem, "QLayout_asQLayoutItem")
	symaddr(&addr_QLayout_spacing_default, "QLayout_spacing_default")
	symaddr(&addr_QLayout_setSpacing_default, "QLayout_setSpacing_default")
	symaddr(&addr_QLayout_invalidate_default, "QLayout_invalidate_default")
	symaddr(&addr_QLayout_geometry_default, "QLayout_geometry_default")
	symaddr(&addr_QLayout_expandingDirections_default, "QLayout_expandingDirections_default")
	symaddr(&addr_QLayout_minimumSize_default, "QLayout_minimumSize_default")
	symaddr(&addr_QLayout_maximumSize_default, "QLayout_maximumSize_default")
	symaddr(&addr_QLayout_setGeometry_default, "QLayout_setGeometry_default")
	symaddr(&addr_QLayout_indexOf_default, "QLayout_indexOf_default")
	symaddr(&addr_QLayout_indexOf_default2, "QLayout_indexOf_default2")
	symaddr(&addr_QLayout_isEmpty_default, "QLayout_isEmpty_default")
	symaddr(&addr_QLayout_controlTypes_default, "QLayout_controlTypes_default")
	symaddr(&addr_QLayout_replaceWidget_default, "QLayout_replaceWidget_default")
	symaddr(&addr_QLayout_layout_default, "QLayout_layout_default")
	symaddr(&addr_QLayout_hasHeightForWidth_default, "QLayout_hasHeightForWidth_default")
	symaddr(&addr_QLayout_heightForWidth_default, "QLayout_heightForWidth_default")
	symaddr(&addr_QLayout_minimumHeightForWidth_default, "QLayout_minimumHeightForWidth_default")
	symaddr(&addr_QLayout_widget_default, "QLayout_widget_default")
	symaddr(&addr_QLayout_spacerItem_default, "QLayout_spacerItem_default")
	symaddr(&addr_QLayout_metaObject_default, "QLayout_metaObject_default")
	symaddr(&addr_QLayout_childEvent_default, "QLayout_childEvent_default")
	symaddr(&addr_QLayout_event_default, "QLayout_event_default")
	symaddr(&addr_QLayout_eventFilter_default, "QLayout_eventFilter_default")
	symaddr(&addr_QLayout_timerEvent_default, "QLayout_timerEvent_default")
	symaddr(&addr_QLayout_customEvent_default, "QLayout_customEvent_default")
	symaddr(&addr_QLayout_connectNotify_default, "QLayout_connectNotify_default")
	symaddr(&addr_QLayout_disconnectNotify_default, "QLayout_disconnectNotify_default")
	reg(&pg_QLayout_connect_destroyed, "QLayout_connect_destroyed")
	reg(&pg_QLayout_connect_destroyed2, "QLayout_connect_destroyed2")
	reg(&pg_QLayout_connect_objectNameChanged, "QLayout_connect_objectNameChanged")
	reg(&pg_QLayout__disconnect, "QLayout__disconnect")
	reg(&pg_QLayout__set_override, "QLayout__set_override")
	reg(&pg_QLayout__Go_new, "QLayout__Go_new")
	reg(&pg_QLayout_delete, "QLayout_delete")
}

type QLayout struct {
	h uintptr
	*QObject
}

func bindQLayout(h uintptr) *QLayout {
	if h == 0 {
		return nil
	}
	return &QLayout{h: h, QObject: bindQObject(pg_QLayout_base(h))}
}

func newownedQLayout(h uintptr) *QLayout { return bindQLayout(h) }

func (this *QLayout) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QLayout) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQLayout(h unsafe.Pointer) *QLayout { return bindQLayout(uintptr(h)) }
func UnsafeNewQLayoutOwned(h unsafe.Pointer) *QLayout { return newownedQLayout(uintptr(h)) }

func (this *QLayout) UnsafeRelease() {
	this.h = 0
	if this.QObject != nil {
		this.QObject.UnsafeRelease()
	}
}

func QLayout_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLayout_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QLayout) Spacing() int {
	__r0, _, _ := purego.SyscallN(addr_QLayout_spacing, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QLayout) SetSpacing(arg1 int) {
	purego.SyscallN(addr_QLayout_setSpacing, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QLayout) SetContentsMargins(left int, top int, right int, bottom int) {
	purego.SyscallN(addr_QLayout_setContentsMargins, this.h, uintptr(left), uintptr(top), uintptr(right), uintptr(bottom))
	runtime.KeepAlive(this)
}

func (this *QLayout) SetContentsMarginsWithMargins(margins *QMargins) {
	purego.SyscallN(addr_QLayout_setContentsMargins2, this.h, uintptr(margins.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(margins)
}

func (this *QLayout) UnsetContentsMargins() {
	purego.SyscallN(addr_QLayout_unsetContentsMargins, this.h)
	runtime.KeepAlive(this)
}

func (this *QLayout) GetContentsMargins() (int, int, int, int) {
	var left_out int32
	var top_out int32
	var right_out int32
	var bottom_out int32
	purego.SyscallN(addr_QLayout_getContentsMargins, this.h, uintptr(unsafe.Pointer(&left_out)), uintptr(unsafe.Pointer(&top_out)), uintptr(unsafe.Pointer(&right_out)), uintptr(unsafe.Pointer(&bottom_out)))
	runtime.KeepAlive(this)
	return int(left_out), int(top_out), int(right_out), int(bottom_out)
}

func (this *QLayout) ContentsMargins() *QMargins {
	__r, _, _ := purego.SyscallN(addr_QLayout_contentsMargins, this.h)
	runtime.KeepAlive(this)
	return newownedQMargins(__r)
}

func (this *QLayout) ContentsRect() *QRect {
	__r, _, _ := purego.SyscallN(addr_QLayout_contentsRect, this.h)
	runtime.KeepAlive(this)
	return newownedQRect(__r)
}

func (this *QLayout) SetAlignment(w *QWidget, alignment AlignmentFlag) bool {
	__r0, _, _ := purego.SyscallN(addr_QLayout_setAlignment, this.h, uintptr(w.UnsafePointer()), uintptr(alignment))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(w)
	return __r != 0
}

func (this *QLayout) SetAlignment2(l *QLayout, alignment AlignmentFlag) bool {
	__r0, _, _ := purego.SyscallN(addr_QLayout_setAlignment2, this.h, uintptr(l.UnsafePointer()), uintptr(alignment))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(l)
	return __r != 0
}

func (this *QLayout) SetSizeConstraint(constraint QLayout__SizeConstraint) {
	purego.SyscallN(addr_QLayout_setSizeConstraint, this.h, uintptr(constraint))
	runtime.KeepAlive(this)
}

func (this *QLayout) SizeConstraint() QLayout__SizeConstraint {
	__r0, _, _ := purego.SyscallN(addr_QLayout_sizeConstraint, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QLayout__SizeConstraint(__r)
}

func (this *QLayout) SetSizeConstraints(horizontal QLayout__SizeConstraint, vertical QLayout__SizeConstraint) {
	purego.SyscallN(addr_QLayout_setSizeConstraints, this.h, uintptr(horizontal), uintptr(vertical))
	runtime.KeepAlive(this)
}

func (this *QLayout) SetHorizontalSizeConstraint(constraint QLayout__SizeConstraint) {
	purego.SyscallN(addr_QLayout_setHorizontalSizeConstraint, this.h, uintptr(constraint))
	runtime.KeepAlive(this)
}

func (this *QLayout) HorizontalSizeConstraint() QLayout__SizeConstraint {
	__r0, _, _ := purego.SyscallN(addr_QLayout_horizontalSizeConstraint, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QLayout__SizeConstraint(__r)
}

func (this *QLayout) SetVerticalSizeConstraint(constraint QLayout__SizeConstraint) {
	purego.SyscallN(addr_QLayout_setVerticalSizeConstraint, this.h, uintptr(constraint))
	runtime.KeepAlive(this)
}

func (this *QLayout) VerticalSizeConstraint() QLayout__SizeConstraint {
	__r0, _, _ := purego.SyscallN(addr_QLayout_verticalSizeConstraint, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QLayout__SizeConstraint(__r)
}

func (this *QLayout) SetMenuBar(w *QWidget) {
	purego.SyscallN(addr_QLayout_setMenuBar, this.h, uintptr(w.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(w)
}

func (this *QLayout) MenuBar() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QLayout_menuBar, this.h)
	runtime.KeepAlive(this)
	return bindQWidget(__r)
}

func (this *QLayout) ParentWidget() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QLayout_parentWidget, this.h)
	runtime.KeepAlive(this)
	return bindQWidget(__r)
}

func (this *QLayout) Invalidate() {
	purego.SyscallN(addr_QLayout_invalidate, this.h)
	runtime.KeepAlive(this)
}

func (this *QLayout) Geometry() *QRect {
	__r, _, _ := purego.SyscallN(addr_QLayout_geometry, this.h)
	runtime.KeepAlive(this)
	return newownedQRect(__r)
}

func (this *QLayout) Activate() bool {
	__r0, _, _ := purego.SyscallN(addr_QLayout_activate, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QLayout) Update() {
	purego.SyscallN(addr_QLayout_update, this.h)
	runtime.KeepAlive(this)
}

func (this *QLayout) AddWidget(w *QWidget) {
	purego.SyscallN(addr_QLayout_addWidget, this.h, uintptr(w.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(w)
}

func (this *QLayout) AddItem(arg1 *QLayoutItem) {
	purego.SyscallN(addr_QLayout_addItem, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QLayout) RemoveWidget(w *QWidget) {
	purego.SyscallN(addr_QLayout_removeWidget, this.h, uintptr(w.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(w)
}

func (this *QLayout) RemoveItem(arg1 *QLayoutItem) {
	purego.SyscallN(addr_QLayout_removeItem, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QLayout) ExpandingDirections() Orientation {
	__r0, _, _ := purego.SyscallN(addr_QLayout_expandingDirections, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return Orientation(__r)
}

func (this *QLayout) MinimumSize() *QSize {
	__r, _, _ := purego.SyscallN(addr_QLayout_minimumSize, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QLayout) MaximumSize() *QSize {
	__r, _, _ := purego.SyscallN(addr_QLayout_maximumSize, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QLayout) SetGeometry(arg1 *QRect) {
	purego.SyscallN(addr_QLayout_setGeometry, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QLayout) ItemAt(index int) *QLayoutItem {
	__r, _, _ := purego.SyscallN(addr_QLayout_itemAt, this.h, uintptr(index))
	runtime.KeepAlive(this)
	return bindQLayoutItem(__r)
}

func (this *QLayout) TakeAt(index int) *QLayoutItem {
	__r, _, _ := purego.SyscallN(addr_QLayout_takeAt, this.h, uintptr(index))
	runtime.KeepAlive(this)
	return bindQLayoutItem(__r)
}

func (this *QLayout) IndexOf(arg1 *QWidget) int {
	__r0, _, _ := purego.SyscallN(addr_QLayout_indexOf, this.h, uintptr(arg1.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return int(__r)
}

func (this *QLayout) IndexOfWithQLayoutItem(arg1 *QLayoutItem) int {
	__r0, _, _ := purego.SyscallN(addr_QLayout_indexOf2, this.h, uintptr(arg1.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return int(__r)
}

func (this *QLayout) Count() int {
	__r0, _, _ := purego.SyscallN(addr_QLayout_count, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QLayout) IsEmpty() bool {
	__r0, _, _ := purego.SyscallN(addr_QLayout_isEmpty, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QLayout) ControlTypes() QSizePolicy__ControlType {
	__r0, _, _ := purego.SyscallN(addr_QLayout_controlTypes, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QSizePolicy__ControlType(__r)
}

func (this *QLayout) ReplaceWidget(from *QWidget, to *QWidget, options FindChildOption) *QLayoutItem {
	__r, _, _ := purego.SyscallN(addr_QLayout_replaceWidget, this.h, uintptr(from.UnsafePointer()), uintptr(to.UnsafePointer()), uintptr(options))
	runtime.KeepAlive(this)
	runtime.KeepAlive(from)
	runtime.KeepAlive(to)
	return bindQLayoutItem(__r)
}

func (this *QLayout) TotalMinimumHeightForWidth(w int) int {
	__r0, _, _ := purego.SyscallN(addr_QLayout_totalMinimumHeightForWidth, this.h, uintptr(w))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QLayout) TotalHeightForWidth(w int) int {
	__r0, _, _ := purego.SyscallN(addr_QLayout_totalHeightForWidth, this.h, uintptr(w))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QLayout) TotalMinimumSize() *QSize {
	__r, _, _ := purego.SyscallN(addr_QLayout_totalMinimumSize, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QLayout) TotalMaximumSize() *QSize {
	__r, _, _ := purego.SyscallN(addr_QLayout_totalMaximumSize, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QLayout) TotalSizeHint() *QSize {
	__r, _, _ := purego.SyscallN(addr_QLayout_totalSizeHint, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QLayout) Layout() *QLayout {
	__r, _, _ := purego.SyscallN(addr_QLayout_layout, this.h)
	runtime.KeepAlive(this)
	return bindQLayout(__r)
}

func (this *QLayout) SetEnabled(arg1 bool) {
	purego.SyscallN(addr_QLayout_setEnabled, this.h, uintptr(b2i(arg1)))
	runtime.KeepAlive(this)
}

func (this *QLayout) IsEnabled() bool {
	__r0, _, _ := purego.SyscallN(addr_QLayout_isEnabled, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QLayout_ClosestAcceptableSize(w *QWidget, s *QSize) *QSize {
	__r, _, _ := purego.SyscallN(addr_QLayout_closestAcceptableSize, uintptr(w.UnsafePointer()), uintptr(s.UnsafePointer()))
	runtime.KeepAlive(w)
	runtime.KeepAlive(s)
	return newownedQSize(__r)
}

func (this *QLayout) WidgetEvent(arg1 *QEvent) {
	purego.SyscallN(addr_QLayout_widgetEvent, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QLayout) AddChildLayout(l *QLayout) {
	purego.SyscallN(addr_QLayout_addChildLayout, this.h, uintptr(l.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(l)
}

func (this *QLayout) AddChildWidget(w *QWidget) {
	purego.SyscallN(addr_QLayout_addChildWidget, this.h, uintptr(w.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(w)
}

func (this *QLayout) AdoptLayout(layout *QLayout) bool {
	__r0, _, _ := purego.SyscallN(addr_QLayout_adoptLayout, this.h, uintptr(layout.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(layout)
	return __r != 0
}

func (this *QLayout) AlignmentRect(arg1 *QRect) *QRect {
	__r, _, _ := purego.SyscallN(addr_QLayout_alignmentRect, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return newownedQRect(__r)
}

func QLayout_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QLayout_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QLayout_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QLayout_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QLayout) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QLayout_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QLayout) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QLayout_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QLayout) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QLayout_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QLayout) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QLayout_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func (this *QLayout) SizeHint() *QSize {
	__r, _, _ := purego.SyscallN(addr_QLayout_sizeHint, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QLayout) HasHeightForWidth() bool {
	__r0, _, _ := purego.SyscallN(addr_QLayout_hasHeightForWidth, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QLayout) HeightForWidth(arg1 int) int {
	__r0, _, _ := purego.SyscallN(addr_QLayout_heightForWidth, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QLayout) MinimumHeightForWidth(arg1 int) int {
	__r0, _, _ := purego.SyscallN(addr_QLayout_minimumHeightForWidth, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QLayout) Widget() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QLayout_widget, this.h)
	runtime.KeepAlive(this)
	return bindQWidget(__r)
}

func (this *QLayout) SpacerItem() *QSpacerItem {
	__r, _, _ := purego.SyscallN(addr_QLayout_spacerItem, this.h)
	runtime.KeepAlive(this)
	return bindQSpacerItem(__r)
}

func (this *QLayout) Alignment() AlignmentFlag {
	__r0, _, _ := purego.SyscallN(addr_QLayout_alignment, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return AlignmentFlag(__r)
}

func QLayout_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLayout_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QLayout_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLayout_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QLayout_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QLayout_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QLayout_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QLayout_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QLayout) AsQLayoutItem() *QLayoutItem {
	__r, _, _ := purego.SyscallN(addr_QLayout_asQLayoutItem, this.h)
	runtime.KeepAlive(this)
	return bindQLayoutItem(__r)
}

func (this *QLayout) SpacingDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QLayout_spacing_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QLayout) SetSpacingDefault(arg1 int) {
	purego.SyscallN(addr_QLayout_setSpacing_default, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QLayout) InvalidateDefault() {
	purego.SyscallN(addr_QLayout_invalidate_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QLayout) GeometryDefault() *QRect {
	__r, _, _ := purego.SyscallN(addr_QLayout_geometry_default, this.h)
	runtime.KeepAlive(this)
	return newownedQRect(__r)
}

func (this *QLayout) ExpandingDirectionsDefault() Orientation {
	__r0, _, _ := purego.SyscallN(addr_QLayout_expandingDirections_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return Orientation(__r)
}

func (this *QLayout) MinimumSizeDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QLayout_minimumSize_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QLayout) MaximumSizeDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QLayout_maximumSize_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QLayout) SetGeometryDefault(arg1 *QRect) {
	purego.SyscallN(addr_QLayout_setGeometry_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QLayout) IndexOfDefault(arg1 *QWidget) int {
	__r0, _, _ := purego.SyscallN(addr_QLayout_indexOf_default, this.h, uintptr(arg1.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return int(__r)
}

func (this *QLayout) IndexOfWithQLayoutItemDefault(arg1 *QLayoutItem) int {
	__r0, _, _ := purego.SyscallN(addr_QLayout_indexOf_default2, this.h, uintptr(arg1.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return int(__r)
}

func (this *QLayout) IsEmptyDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QLayout_isEmpty_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QLayout) ControlTypesDefault() QSizePolicy__ControlType {
	__r0, _, _ := purego.SyscallN(addr_QLayout_controlTypes_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QSizePolicy__ControlType(__r)
}

func (this *QLayout) ReplaceWidgetDefault(from *QWidget, to *QWidget, options FindChildOption) *QLayoutItem {
	__r, _, _ := purego.SyscallN(addr_QLayout_replaceWidget_default, this.h, uintptr(from.UnsafePointer()), uintptr(to.UnsafePointer()), uintptr(options))
	runtime.KeepAlive(this)
	runtime.KeepAlive(from)
	runtime.KeepAlive(to)
	return bindQLayoutItem(__r)
}

func (this *QLayout) LayoutDefault() *QLayout {
	__r, _, _ := purego.SyscallN(addr_QLayout_layout_default, this.h)
	runtime.KeepAlive(this)
	return bindQLayout(__r)
}

func (this *QLayout) HasHeightForWidthDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QLayout_hasHeightForWidth_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QLayout) HeightForWidthDefault(arg1 int) int {
	__r0, _, _ := purego.SyscallN(addr_QLayout_heightForWidth_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QLayout) MinimumHeightForWidthDefault(arg1 int) int {
	__r0, _, _ := purego.SyscallN(addr_QLayout_minimumHeightForWidth_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QLayout) WidgetDefault() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QLayout_widget_default, this.h)
	runtime.KeepAlive(this)
	return bindQWidget(__r)
}

func (this *QLayout) SpacerItemDefault() *QSpacerItem {
	__r, _, _ := purego.SyscallN(addr_QLayout_spacerItem_default, this.h)
	runtime.KeepAlive(this)
	return bindQSpacerItem(__r)
}

func (this *QLayout) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QLayout_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QLayout) ChildEventDefault(e *QChildEvent) {
	purego.SyscallN(addr_QLayout_childEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QLayout) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QLayout_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QLayout) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QLayout_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QLayout) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QLayout_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QLayout) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QLayout_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QLayout) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QLayout_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QLayout) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QLayout_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QLayout) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QLayout_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QLayout) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QLayout_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QLayout) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QLayout_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QLayout) Disconnect(handle int64) {
	pg_QLayout__disconnect(handle)
	unregisterCallback(handle)
}

func QLayout_virtIdx(slot int64) int32 {
	switch slot {
	case 729690261457684101:
		return 0
	case 6884465960736393867:
		return 1
	case 7508512214013391107:
		return 2
	case 7549889108533486643:
		return 3
	case -5627843058667511180:
		return 4
	case 1607084412591819503:
		return 5
	case -4587105490415799682:
		return 6
	case 4039479028406881496:
		return 7
	case 6991341317990632368:
		return 8
	case 1050490530495513297:
		return 9
	case 415665469015095569:
		return 10
	case -7246748916541192690:
		return 11
	case 4158099606088301136:
		return 12
	case 5684820848726969227:
		return 13
	case -7258473138030329570:
		return 14
	case -3276435765994963516:
		return 15
	case -626822430251098498:
		return 16
	case 3278947275714896887:
		return 17
	case -712862742159708061:
		return 18
	case -7183390618627992737:
		return 19
	case -1185315367526857833:
		return 20
	case 1237236456312824991:
		return 21
	case 5217623514630858297:
		return 22
	case -5409927581812374498:
		return 23
	case 3769868229123868474:
		return 24
	case 7379029982170503008:
		return 25
	case -568369772982215962:
		return 26
	case -8266762397995423915:
		return 27
	case 4311534060495737339:
		return 28
	case 5037837229528203129:
		return 29
	case -4575937485689502171:
		return 30
	case 2272240669359776278:
		return 31
	}
	return -1
}

func QLayout_enable(h uintptr, idx int32) { pg_QLayout__set_override(h, idx) }

func NewQLayout() *QLayout {
	id := nextVirtID()
	o := newownedQLayout(pg_QLayout__Go_new(id))
	o.UnsafeSetGoID(id)
	registerInstance(id, o.h, QLayout_virtIdx, QLayout_enable)
	return o
}

func (this *QLayout) OnSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 729690261457684101, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sizeHint"
}

func (this *QLayout) OnAddItem(fn func(*QLayoutItem)) {
	registerVirt(this.UnsafeGoID(), 6884465960736393867, func(arg1 uintptr) { fn(UnsafeNewQLayoutItem(unsafe.Pointer(arg1))) }) // slot "addItem"
}

func (this *QLayout) OnItemAt(fn func(int) *QLayoutItem) {
	registerVirt(this.UnsafeGoID(), 7508512214013391107, func(index int32) uintptr { return uintptr(fn(int(index)).UnsafePointer()) }) // slot "itemAt"
}

func (this *QLayout) OnTakeAt(fn func(int) *QLayoutItem) {
	registerVirt(this.UnsafeGoID(), 7549889108533486643, func(index int32) uintptr { return uintptr(fn(int(index)).UnsafePointer()) }) // slot "takeAt"
}

func (this *QLayout) OnCount(fn func() int) {
	registerVirt(this.UnsafeGoID(), -5627843058667511180, func() int32 { return int32(fn()) }) // slot "count"
}

func (this *QLayout) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QLayout) OnSpacing(fn func() int) {
	registerVirt(this.UnsafeGoID(), -4587105490415799682, func() int32 { return int32(fn()) }) // slot "spacing"
}

func (this *QLayout) OnSetSpacing(fn func(int)) {
	registerVirt(this.UnsafeGoID(), 4039479028406881496, func(arg1 int32) { fn(int(arg1)) }) // slot "setSpacing"
}

func (this *QLayout) OnInvalidate(fn func()) {
	registerVirt(this.UnsafeGoID(), 6991341317990632368, func() { fn() }) // slot "invalidate"
}

func (this *QLayout) OnGeometry(fn func() *QRect) {
	registerVirt(this.UnsafeGoID(), 1050490530495513297, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "geometry"
}

func (this *QLayout) OnExpandingDirections(fn func() Orientation) {
	registerVirt(this.UnsafeGoID(), 415665469015095569, func() int32 { return int32(fn()) }) // slot "expandingDirections"
}

func (this *QLayout) OnMinimumSize(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), -7246748916541192690, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "minimumSize"
}

func (this *QLayout) OnMaximumSize(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 4158099606088301136, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "maximumSize"
}

func (this *QLayout) OnSetGeometry(fn func(*QRect)) {
	registerVirt(this.UnsafeGoID(), 5684820848726969227, func(arg1 uintptr) { fn(UnsafeNewQRect(unsafe.Pointer(arg1))) }) // slot "setGeometry"
}

func (this *QLayout) OnIndexOf(fn func(*QWidget) int) {
	registerVirt(this.UnsafeGoID(), -7258473138030329570, func(arg1 uintptr) int32 { return int32(fn(UnsafeNewQWidget(unsafe.Pointer(arg1)))) }) // slot "indexOf"
}

func (this *QLayout) OnIndexOf2(fn func(*QLayoutItem) int) {
	registerVirt(this.UnsafeGoID(), -3276435765994963516, func(arg1 uintptr) int32 { return int32(fn(UnsafeNewQLayoutItem(unsafe.Pointer(arg1)))) }) // slot "indexOf2"
}

func (this *QLayout) OnIsEmpty(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -626822430251098498, func() int32 { return b2i(fn()) }) // slot "isEmpty"
}

func (this *QLayout) OnControlTypes(fn func() QSizePolicy__ControlType) {
	registerVirt(this.UnsafeGoID(), 3278947275714896887, func() int32 { return int32(fn()) }) // slot "controlTypes"
}

func (this *QLayout) OnReplaceWidget(fn func(*QWidget, *QWidget, FindChildOption) *QLayoutItem) {
	registerVirt(this.UnsafeGoID(), -712862742159708061, func(from uintptr, to uintptr, options int32) uintptr { return uintptr(fn(UnsafeNewQWidget(unsafe.Pointer(from)), UnsafeNewQWidget(unsafe.Pointer(to)), FindChildOption(options)).UnsafePointer()) }) // slot "replaceWidget"
}

func (this *QLayout) OnLayout(fn func() *QLayout) {
	registerVirt(this.UnsafeGoID(), -7183390618627992737, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "layout"
}

func (this *QLayout) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(e uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(e))) }) // slot "childEvent"
}

func (this *QLayout) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QLayout) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QLayout) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QLayout) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QLayout) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QLayout) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QLayout) OnHasHeightForWidth(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -8266762397995423915, func() int32 { return b2i(fn()) }) // slot "hasHeightForWidth"
}

func (this *QLayout) OnHeightForWidth(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), 4311534060495737339, func(arg1 int32) int32 { return int32(fn(int(arg1))) }) // slot "heightForWidth"
}

func (this *QLayout) OnMinimumHeightForWidth(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), 5037837229528203129, func(arg1 int32) int32 { return int32(fn(int(arg1))) }) // slot "minimumHeightForWidth"
}

func (this *QLayout) OnWidget(fn func() *QWidget) {
	registerVirt(this.UnsafeGoID(), -4575937485689502171, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "widget"
}

func (this *QLayout) OnSpacerItem(fn func() *QSpacerItem) {
	registerVirt(this.UnsafeGoID(), 2272240669359776278, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "spacerItem"
}

func (this *QLayout) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QLayout_delete(this.h)
	}
	this.UnsafeRelease()
}
