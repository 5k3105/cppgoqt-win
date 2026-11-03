package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QGridLayout_chain uintptr
	addr_QGridLayout_new uintptr
	addr_QGridLayout_new2 uintptr
	addr_QGridLayout_tr uintptr
	addr_QGridLayout_setHorizontalSpacing uintptr
	addr_QGridLayout_horizontalSpacing uintptr
	addr_QGridLayout_setVerticalSpacing uintptr
	addr_QGridLayout_verticalSpacing uintptr
	addr_QGridLayout_setRowStretch uintptr
	addr_QGridLayout_setColumnStretch uintptr
	addr_QGridLayout_rowStretch uintptr
	addr_QGridLayout_columnStretch uintptr
	addr_QGridLayout_setRowMinimumHeight uintptr
	addr_QGridLayout_setColumnMinimumWidth uintptr
	addr_QGridLayout_rowMinimumHeight uintptr
	addr_QGridLayout_columnMinimumWidth uintptr
	addr_QGridLayout_columnCount uintptr
	addr_QGridLayout_rowCount uintptr
	addr_QGridLayout_cellRect uintptr
	addr_QGridLayout_addWidget uintptr
	addr_QGridLayout_addWidget2 uintptr
	addr_QGridLayout_addWidget3 uintptr
	addr_QGridLayout_addLayout uintptr
	addr_QGridLayout_addLayout2 uintptr
	addr_QGridLayout_setOriginCorner uintptr
	addr_QGridLayout_originCorner uintptr
	addr_QGridLayout_itemAtPosition uintptr
	addr_QGridLayout_addItem uintptr
	addr_QGridLayout_setDefaultPositioning uintptr
	addr_QGridLayout_getItemPosition uintptr
	addr_QGridLayout_closestAcceptableSize uintptr
	addr_QGridLayout_widgetEvent uintptr
	addr_QGridLayout_addChildLayout uintptr
	addr_QGridLayout_addChildWidget uintptr
	addr_QGridLayout_adoptLayout uintptr
	addr_QGridLayout_alignmentRect uintptr
	addr_QGridLayout_connect uintptr
	addr_QGridLayout_connect2 uintptr
	addr_QGridLayout_sender uintptr
	addr_QGridLayout_senderSignalIndex uintptr
	addr_QGridLayout_receivers uintptr
	addr_QGridLayout_isSignalConnected uintptr
	addr_QGridLayout_tr2 uintptr
	addr_QGridLayout_tr3 uintptr
	addr_QGridLayout_addWidget4 uintptr
	addr_QGridLayout_addWidget5 uintptr
	addr_QGridLayout_addLayout3 uintptr
	addr_QGridLayout_addLayout4 uintptr
	addr_QGridLayout_addItem2 uintptr
	addr_QGridLayout_addItem3 uintptr
	addr_QGridLayout_addItem4 uintptr
	addr_QGridLayout_connect3 uintptr
	addr_QGridLayout_connect4 uintptr
	addr_QGridLayout_addItem_default uintptr
	addr_QGridLayout_addItem_default2 uintptr
	addr_QGridLayout_addItem_default3 uintptr
	addr_QGridLayout_addItem_default4 uintptr
	addr_QGridLayout_metaObject_default uintptr
	addr_QGridLayout_sizeHint_default uintptr
	addr_QGridLayout_minimumSize_default uintptr
	addr_QGridLayout_maximumSize_default uintptr
	addr_QGridLayout_setSpacing_default uintptr
	addr_QGridLayout_spacing_default uintptr
	addr_QGridLayout_hasHeightForWidth_default uintptr
	addr_QGridLayout_heightForWidth_default uintptr
	addr_QGridLayout_minimumHeightForWidth_default uintptr
	addr_QGridLayout_expandingDirections_default uintptr
	addr_QGridLayout_invalidate_default uintptr
	addr_QGridLayout_itemAt_default uintptr
	addr_QGridLayout_takeAt_default uintptr
	addr_QGridLayout_count_default uintptr
	addr_QGridLayout_setGeometry_default uintptr
	addr_QGridLayout_geometry_default uintptr
	addr_QGridLayout_indexOf_default uintptr
	addr_QGridLayout_isEmpty_default uintptr
	addr_QGridLayout_controlTypes_default uintptr
	addr_QGridLayout_replaceWidget_default uintptr
	addr_QGridLayout_layout_default uintptr
	addr_QGridLayout_childEvent_default uintptr
	addr_QGridLayout_event_default uintptr
	addr_QGridLayout_eventFilter_default uintptr
	addr_QGridLayout_timerEvent_default uintptr
	addr_QGridLayout_customEvent_default uintptr
	addr_QGridLayout_connectNotify_default uintptr
	addr_QGridLayout_disconnectNotify_default uintptr
	addr_QGridLayout_widget_default uintptr
	addr_QGridLayout_spacerItem_default uintptr
	pg_QGridLayout_connect_destroyed func(uintptr, int64)
	pg_QGridLayout_connect_destroyed2 func(uintptr, int64)
	pg_QGridLayout_connect_objectNameChanged func(uintptr, int64)
	pg_QGridLayout__disconnect func(int64)
	pg_QGridLayout__set_override func(uintptr, int32)
	pg_QGridLayout_delete func(uintptr)
)

func init() {
	symaddr(&addr_QGridLayout_chain, "QGridLayout_chain")
	symaddr(&addr_QGridLayout_new, "QGridLayout_new")
	symaddr(&addr_QGridLayout_new2, "QGridLayout_new2")
	symaddr(&addr_QGridLayout_tr, "QGridLayout_tr")
	symaddr(&addr_QGridLayout_setHorizontalSpacing, "QGridLayout_setHorizontalSpacing")
	symaddr(&addr_QGridLayout_horizontalSpacing, "QGridLayout_horizontalSpacing")
	symaddr(&addr_QGridLayout_setVerticalSpacing, "QGridLayout_setVerticalSpacing")
	symaddr(&addr_QGridLayout_verticalSpacing, "QGridLayout_verticalSpacing")
	symaddr(&addr_QGridLayout_setRowStretch, "QGridLayout_setRowStretch")
	symaddr(&addr_QGridLayout_setColumnStretch, "QGridLayout_setColumnStretch")
	symaddr(&addr_QGridLayout_rowStretch, "QGridLayout_rowStretch")
	symaddr(&addr_QGridLayout_columnStretch, "QGridLayout_columnStretch")
	symaddr(&addr_QGridLayout_setRowMinimumHeight, "QGridLayout_setRowMinimumHeight")
	symaddr(&addr_QGridLayout_setColumnMinimumWidth, "QGridLayout_setColumnMinimumWidth")
	symaddr(&addr_QGridLayout_rowMinimumHeight, "QGridLayout_rowMinimumHeight")
	symaddr(&addr_QGridLayout_columnMinimumWidth, "QGridLayout_columnMinimumWidth")
	symaddr(&addr_QGridLayout_columnCount, "QGridLayout_columnCount")
	symaddr(&addr_QGridLayout_rowCount, "QGridLayout_rowCount")
	symaddr(&addr_QGridLayout_cellRect, "QGridLayout_cellRect")
	symaddr(&addr_QGridLayout_addWidget, "QGridLayout_addWidget")
	symaddr(&addr_QGridLayout_addWidget2, "QGridLayout_addWidget2")
	symaddr(&addr_QGridLayout_addWidget3, "QGridLayout_addWidget3")
	symaddr(&addr_QGridLayout_addLayout, "QGridLayout_addLayout")
	symaddr(&addr_QGridLayout_addLayout2, "QGridLayout_addLayout2")
	symaddr(&addr_QGridLayout_setOriginCorner, "QGridLayout_setOriginCorner")
	symaddr(&addr_QGridLayout_originCorner, "QGridLayout_originCorner")
	symaddr(&addr_QGridLayout_itemAtPosition, "QGridLayout_itemAtPosition")
	symaddr(&addr_QGridLayout_addItem, "QGridLayout_addItem")
	symaddr(&addr_QGridLayout_setDefaultPositioning, "QGridLayout_setDefaultPositioning")
	symaddr(&addr_QGridLayout_getItemPosition, "QGridLayout_getItemPosition")
	symaddr(&addr_QGridLayout_closestAcceptableSize, "QGridLayout_closestAcceptableSize")
	symaddr(&addr_QGridLayout_widgetEvent, "QGridLayout_widgetEvent")
	symaddr(&addr_QGridLayout_addChildLayout, "QGridLayout_addChildLayout")
	symaddr(&addr_QGridLayout_addChildWidget, "QGridLayout_addChildWidget")
	symaddr(&addr_QGridLayout_adoptLayout, "QGridLayout_adoptLayout")
	symaddr(&addr_QGridLayout_alignmentRect, "QGridLayout_alignmentRect")
	symaddr(&addr_QGridLayout_connect, "QGridLayout_connect")
	symaddr(&addr_QGridLayout_connect2, "QGridLayout_connect2")
	symaddr(&addr_QGridLayout_sender, "QGridLayout_sender")
	symaddr(&addr_QGridLayout_senderSignalIndex, "QGridLayout_senderSignalIndex")
	symaddr(&addr_QGridLayout_receivers, "QGridLayout_receivers")
	symaddr(&addr_QGridLayout_isSignalConnected, "QGridLayout_isSignalConnected")
	symaddr(&addr_QGridLayout_tr2, "QGridLayout_tr2")
	symaddr(&addr_QGridLayout_tr3, "QGridLayout_tr3")
	symaddr(&addr_QGridLayout_addWidget4, "QGridLayout_addWidget4")
	symaddr(&addr_QGridLayout_addWidget5, "QGridLayout_addWidget5")
	symaddr(&addr_QGridLayout_addLayout3, "QGridLayout_addLayout3")
	symaddr(&addr_QGridLayout_addLayout4, "QGridLayout_addLayout4")
	symaddr(&addr_QGridLayout_addItem2, "QGridLayout_addItem2")
	symaddr(&addr_QGridLayout_addItem3, "QGridLayout_addItem3")
	symaddr(&addr_QGridLayout_addItem4, "QGridLayout_addItem4")
	symaddr(&addr_QGridLayout_connect3, "QGridLayout_connect3")
	symaddr(&addr_QGridLayout_connect4, "QGridLayout_connect4")
	symaddr(&addr_QGridLayout_addItem_default, "QGridLayout_addItem_default")
	symaddr(&addr_QGridLayout_addItem_default2, "QGridLayout_addItem_default2")
	symaddr(&addr_QGridLayout_addItem_default3, "QGridLayout_addItem_default3")
	symaddr(&addr_QGridLayout_addItem_default4, "QGridLayout_addItem_default4")
	symaddr(&addr_QGridLayout_metaObject_default, "QGridLayout_metaObject_default")
	symaddr(&addr_QGridLayout_sizeHint_default, "QGridLayout_sizeHint_default")
	symaddr(&addr_QGridLayout_minimumSize_default, "QGridLayout_minimumSize_default")
	symaddr(&addr_QGridLayout_maximumSize_default, "QGridLayout_maximumSize_default")
	symaddr(&addr_QGridLayout_setSpacing_default, "QGridLayout_setSpacing_default")
	symaddr(&addr_QGridLayout_spacing_default, "QGridLayout_spacing_default")
	symaddr(&addr_QGridLayout_hasHeightForWidth_default, "QGridLayout_hasHeightForWidth_default")
	symaddr(&addr_QGridLayout_heightForWidth_default, "QGridLayout_heightForWidth_default")
	symaddr(&addr_QGridLayout_minimumHeightForWidth_default, "QGridLayout_minimumHeightForWidth_default")
	symaddr(&addr_QGridLayout_expandingDirections_default, "QGridLayout_expandingDirections_default")
	symaddr(&addr_QGridLayout_invalidate_default, "QGridLayout_invalidate_default")
	symaddr(&addr_QGridLayout_itemAt_default, "QGridLayout_itemAt_default")
	symaddr(&addr_QGridLayout_takeAt_default, "QGridLayout_takeAt_default")
	symaddr(&addr_QGridLayout_count_default, "QGridLayout_count_default")
	symaddr(&addr_QGridLayout_setGeometry_default, "QGridLayout_setGeometry_default")
	symaddr(&addr_QGridLayout_geometry_default, "QGridLayout_geometry_default")
	symaddr(&addr_QGridLayout_indexOf_default, "QGridLayout_indexOf_default")
	symaddr(&addr_QGridLayout_isEmpty_default, "QGridLayout_isEmpty_default")
	symaddr(&addr_QGridLayout_controlTypes_default, "QGridLayout_controlTypes_default")
	symaddr(&addr_QGridLayout_replaceWidget_default, "QGridLayout_replaceWidget_default")
	symaddr(&addr_QGridLayout_layout_default, "QGridLayout_layout_default")
	symaddr(&addr_QGridLayout_childEvent_default, "QGridLayout_childEvent_default")
	symaddr(&addr_QGridLayout_event_default, "QGridLayout_event_default")
	symaddr(&addr_QGridLayout_eventFilter_default, "QGridLayout_eventFilter_default")
	symaddr(&addr_QGridLayout_timerEvent_default, "QGridLayout_timerEvent_default")
	symaddr(&addr_QGridLayout_customEvent_default, "QGridLayout_customEvent_default")
	symaddr(&addr_QGridLayout_connectNotify_default, "QGridLayout_connectNotify_default")
	symaddr(&addr_QGridLayout_disconnectNotify_default, "QGridLayout_disconnectNotify_default")
	symaddr(&addr_QGridLayout_widget_default, "QGridLayout_widget_default")
	symaddr(&addr_QGridLayout_spacerItem_default, "QGridLayout_spacerItem_default")
	reg(&pg_QGridLayout_connect_destroyed, "QGridLayout_connect_destroyed")
	reg(&pg_QGridLayout_connect_destroyed2, "QGridLayout_connect_destroyed2")
	reg(&pg_QGridLayout_connect_objectNameChanged, "QGridLayout_connect_objectNameChanged")
	reg(&pg_QGridLayout__disconnect, "QGridLayout__disconnect")
	reg(&pg_QGridLayout__set_override, "QGridLayout__set_override")
	reg(&pg_QGridLayout_delete, "QGridLayout_delete")
}

type QGridLayout struct {
	h uintptr
	*QLayout
}

func bindQGridLayout(h uintptr) *QGridLayout {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QGridLayout_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QGridLayout{h: h, QLayout: &QLayout{h: __c[0], QObject: &QObject{h: __c[1]}}}
}

func newownedQGridLayout(h uintptr) *QGridLayout { return bindQGridLayout(h) }

func (this *QGridLayout) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QGridLayout) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQGridLayout(h unsafe.Pointer) *QGridLayout { return bindQGridLayout(uintptr(h)) }
func UnsafeNewQGridLayoutOwned(h unsafe.Pointer) *QGridLayout { return newownedQGridLayout(uintptr(h)) }

func (this *QGridLayout) UnsafeRelease() {
	this.h = 0
	if this.QLayout != nil {
		this.QLayout.UnsafeRelease()
	}
}

func NewQGridLayout(parent *QWidget) *QGridLayout {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QGridLayout_new, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQGridLayout(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGridLayout_virtIdx, QGridLayout_enable)
	return __o
}

func NewQGridLayout2() *QGridLayout {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QGridLayout_new2, uintptr(__goid))
	__o := newownedQGridLayout(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGridLayout_virtIdx, QGridLayout_enable)
	return __o
}

func QGridLayout_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGridLayout_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QGridLayout) SetHorizontalSpacing(spacing int) {
	purego.SyscallN(addr_QGridLayout_setHorizontalSpacing, this.h, uintptr(spacing))
	runtime.KeepAlive(this)
}

func (this *QGridLayout) HorizontalSpacing() int {
	__r0, _, _ := purego.SyscallN(addr_QGridLayout_horizontalSpacing, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QGridLayout) SetVerticalSpacing(spacing int) {
	purego.SyscallN(addr_QGridLayout_setVerticalSpacing, this.h, uintptr(spacing))
	runtime.KeepAlive(this)
}

func (this *QGridLayout) VerticalSpacing() int {
	__r0, _, _ := purego.SyscallN(addr_QGridLayout_verticalSpacing, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QGridLayout) SetRowStretch(row int, stretch int) {
	purego.SyscallN(addr_QGridLayout_setRowStretch, this.h, uintptr(row), uintptr(stretch))
	runtime.KeepAlive(this)
}

func (this *QGridLayout) SetColumnStretch(column int, stretch int) {
	purego.SyscallN(addr_QGridLayout_setColumnStretch, this.h, uintptr(column), uintptr(stretch))
	runtime.KeepAlive(this)
}

func (this *QGridLayout) RowStretch(row int) int {
	__r0, _, _ := purego.SyscallN(addr_QGridLayout_rowStretch, this.h, uintptr(row))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QGridLayout) ColumnStretch(column int) int {
	__r0, _, _ := purego.SyscallN(addr_QGridLayout_columnStretch, this.h, uintptr(column))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QGridLayout) SetRowMinimumHeight(row int, minSize int) {
	purego.SyscallN(addr_QGridLayout_setRowMinimumHeight, this.h, uintptr(row), uintptr(minSize))
	runtime.KeepAlive(this)
}

func (this *QGridLayout) SetColumnMinimumWidth(column int, minSize int) {
	purego.SyscallN(addr_QGridLayout_setColumnMinimumWidth, this.h, uintptr(column), uintptr(minSize))
	runtime.KeepAlive(this)
}

func (this *QGridLayout) RowMinimumHeight(row int) int {
	__r0, _, _ := purego.SyscallN(addr_QGridLayout_rowMinimumHeight, this.h, uintptr(row))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QGridLayout) ColumnMinimumWidth(column int) int {
	__r0, _, _ := purego.SyscallN(addr_QGridLayout_columnMinimumWidth, this.h, uintptr(column))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QGridLayout) ColumnCount() int {
	__r0, _, _ := purego.SyscallN(addr_QGridLayout_columnCount, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QGridLayout) RowCount() int {
	__r0, _, _ := purego.SyscallN(addr_QGridLayout_rowCount, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QGridLayout) CellRect(row int, column int) *QRect {
	__r, _, _ := purego.SyscallN(addr_QGridLayout_cellRect, this.h, uintptr(row), uintptr(column))
	runtime.KeepAlive(this)
	return newownedQRect(__r)
}

func (this *QGridLayout) AddWidget(w *QWidget) {
	purego.SyscallN(addr_QGridLayout_addWidget, this.h, uintptr(w.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(w)
}

func (this *QGridLayout) AddWidget2(arg1 *QWidget, row int, column int) {
	purego.SyscallN(addr_QGridLayout_addWidget2, this.h, uintptr(arg1.UnsafePointer()), uintptr(row), uintptr(column))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QGridLayout) AddWidget3(arg1 *QWidget, row int, column int, rowSpan int, columnSpan int) {
	purego.SyscallN(addr_QGridLayout_addWidget3, this.h, uintptr(arg1.UnsafePointer()), uintptr(row), uintptr(column), uintptr(rowSpan), uintptr(columnSpan))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QGridLayout) AddLayout(arg1 *QLayout, row int, column int) {
	purego.SyscallN(addr_QGridLayout_addLayout, this.h, uintptr(arg1.UnsafePointer()), uintptr(row), uintptr(column))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QGridLayout) AddLayout2(arg1 *QLayout, row int, column int, rowSpan int, columnSpan int) {
	purego.SyscallN(addr_QGridLayout_addLayout2, this.h, uintptr(arg1.UnsafePointer()), uintptr(row), uintptr(column), uintptr(rowSpan), uintptr(columnSpan))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QGridLayout) SetOriginCorner(arg1 Corner) {
	purego.SyscallN(addr_QGridLayout_setOriginCorner, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QGridLayout) OriginCorner() Corner {
	__r0, _, _ := purego.SyscallN(addr_QGridLayout_originCorner, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return Corner(__r)
}

func (this *QGridLayout) ItemAtPosition(row int, column int) *QLayoutItem {
	__r, _, _ := purego.SyscallN(addr_QGridLayout_itemAtPosition, this.h, uintptr(row), uintptr(column))
	runtime.KeepAlive(this)
	return bindQLayoutItem(__r)
}

func (this *QGridLayout) AddItem(item *QLayoutItem, row int, column int) {
	purego.SyscallN(addr_QGridLayout_addItem, this.h, uintptr(item.UnsafePointer()), uintptr(row), uintptr(column))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QGridLayout) SetDefaultPositioning(n int, orient Orientation) {
	purego.SyscallN(addr_QGridLayout_setDefaultPositioning, this.h, uintptr(n), uintptr(orient))
	runtime.KeepAlive(this)
}

func (this *QGridLayout) GetItemPosition(idx int) (int, int, int, int) {
	var row_out int32
	var column_out int32
	var rowSpan_out int32
	var columnSpan_out int32
	purego.SyscallN(addr_QGridLayout_getItemPosition, this.h, uintptr(idx), uintptr(unsafe.Pointer(&row_out)), uintptr(unsafe.Pointer(&column_out)), uintptr(unsafe.Pointer(&rowSpan_out)), uintptr(unsafe.Pointer(&columnSpan_out)))
	runtime.KeepAlive(this)
	return int(row_out), int(column_out), int(rowSpan_out), int(columnSpan_out)
}

func QGridLayout_ClosestAcceptableSize(w *QWidget, s *QSize) *QSize {
	__r, _, _ := purego.SyscallN(addr_QGridLayout_closestAcceptableSize, uintptr(w.UnsafePointer()), uintptr(s.UnsafePointer()))
	runtime.KeepAlive(w)
	runtime.KeepAlive(s)
	return newownedQSize(__r)
}

func (this *QGridLayout) WidgetEvent(arg1 *QEvent) {
	purego.SyscallN(addr_QGridLayout_widgetEvent, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QGridLayout) AddChildLayout(l *QLayout) {
	purego.SyscallN(addr_QGridLayout_addChildLayout, this.h, uintptr(l.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(l)
}

func (this *QGridLayout) AddChildWidget(w *QWidget) {
	purego.SyscallN(addr_QGridLayout_addChildWidget, this.h, uintptr(w.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(w)
}

func (this *QGridLayout) AdoptLayout(layout *QLayout) bool {
	__r0, _, _ := purego.SyscallN(addr_QGridLayout_adoptLayout, this.h, uintptr(layout.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(layout)
	return __r != 0
}

func (this *QGridLayout) AlignmentRect(arg1 *QRect) *QRect {
	__r, _, _ := purego.SyscallN(addr_QGridLayout_alignmentRect, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return newownedQRect(__r)
}

func QGridLayout_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QGridLayout_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QGridLayout_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QGridLayout_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QGridLayout) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QGridLayout_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QGridLayout) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QGridLayout_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QGridLayout) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QGridLayout_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QGridLayout) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QGridLayout_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QGridLayout_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGridLayout_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QGridLayout_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGridLayout_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func (this *QGridLayout) AddWidget4(arg1 *QWidget, row int, column int, arg4 AlignmentFlag) {
	purego.SyscallN(addr_QGridLayout_addWidget4, this.h, uintptr(arg1.UnsafePointer()), uintptr(row), uintptr(column), uintptr(arg4))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QGridLayout) AddWidget5(arg1 *QWidget, row int, column int, rowSpan int, columnSpan int, arg6 AlignmentFlag) {
	purego.SyscallN(addr_QGridLayout_addWidget5, this.h, uintptr(arg1.UnsafePointer()), uintptr(row), uintptr(column), uintptr(rowSpan), uintptr(columnSpan), uintptr(arg6))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QGridLayout) AddLayout3(arg1 *QLayout, row int, column int, arg4 AlignmentFlag) {
	purego.SyscallN(addr_QGridLayout_addLayout3, this.h, uintptr(arg1.UnsafePointer()), uintptr(row), uintptr(column), uintptr(arg4))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QGridLayout) AddLayout4(arg1 *QLayout, row int, column int, rowSpan int, columnSpan int, arg6 AlignmentFlag) {
	purego.SyscallN(addr_QGridLayout_addLayout4, this.h, uintptr(arg1.UnsafePointer()), uintptr(row), uintptr(column), uintptr(rowSpan), uintptr(columnSpan), uintptr(arg6))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QGridLayout) AddItem2(item *QLayoutItem, row int, column int, rowSpan int) {
	purego.SyscallN(addr_QGridLayout_addItem2, this.h, uintptr(item.UnsafePointer()), uintptr(row), uintptr(column), uintptr(rowSpan))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QGridLayout) AddItem3(item *QLayoutItem, row int, column int, rowSpan int, columnSpan int) {
	purego.SyscallN(addr_QGridLayout_addItem3, this.h, uintptr(item.UnsafePointer()), uintptr(row), uintptr(column), uintptr(rowSpan), uintptr(columnSpan))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QGridLayout) AddItem4(item *QLayoutItem, row int, column int, rowSpan int, columnSpan int, arg6 AlignmentFlag) {
	purego.SyscallN(addr_QGridLayout_addItem4, this.h, uintptr(item.UnsafePointer()), uintptr(row), uintptr(column), uintptr(rowSpan), uintptr(columnSpan), uintptr(arg6))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func QGridLayout_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QGridLayout_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QGridLayout_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QGridLayout_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QGridLayout) AddItemDefault(item *QLayoutItem, row int, column int) {
	purego.SyscallN(addr_QGridLayout_addItem_default, this.h, uintptr(item.UnsafePointer()), uintptr(row), uintptr(column))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QGridLayout) AddItem2Default(item *QLayoutItem, row int, column int, rowSpan int) {
	purego.SyscallN(addr_QGridLayout_addItem_default2, this.h, uintptr(item.UnsafePointer()), uintptr(row), uintptr(column), uintptr(rowSpan))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QGridLayout) AddItem3Default(item *QLayoutItem, row int, column int, rowSpan int, columnSpan int) {
	purego.SyscallN(addr_QGridLayout_addItem_default3, this.h, uintptr(item.UnsafePointer()), uintptr(row), uintptr(column), uintptr(rowSpan), uintptr(columnSpan))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QGridLayout) AddItem4Default(item *QLayoutItem, row int, column int, rowSpan int, columnSpan int, arg6 AlignmentFlag) {
	purego.SyscallN(addr_QGridLayout_addItem_default4, this.h, uintptr(item.UnsafePointer()), uintptr(row), uintptr(column), uintptr(rowSpan), uintptr(columnSpan), uintptr(arg6))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QGridLayout) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QGridLayout_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QGridLayout) SizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QGridLayout_sizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QGridLayout) MinimumSizeDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QGridLayout_minimumSize_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QGridLayout) MaximumSizeDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QGridLayout_maximumSize_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QGridLayout) SetSpacingDefault(spacing int) {
	purego.SyscallN(addr_QGridLayout_setSpacing_default, this.h, uintptr(spacing))
	runtime.KeepAlive(this)
}

func (this *QGridLayout) SpacingDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QGridLayout_spacing_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QGridLayout) HasHeightForWidthDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QGridLayout_hasHeightForWidth_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGridLayout) HeightForWidthDefault(arg1 int) int {
	__r0, _, _ := purego.SyscallN(addr_QGridLayout_heightForWidth_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QGridLayout) MinimumHeightForWidthDefault(arg1 int) int {
	__r0, _, _ := purego.SyscallN(addr_QGridLayout_minimumHeightForWidth_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QGridLayout) ExpandingDirectionsDefault() Orientation {
	__r0, _, _ := purego.SyscallN(addr_QGridLayout_expandingDirections_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return Orientation(__r)
}

func (this *QGridLayout) InvalidateDefault() {
	purego.SyscallN(addr_QGridLayout_invalidate_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QGridLayout) ItemAtDefault(index int) *QLayoutItem {
	__r, _, _ := purego.SyscallN(addr_QGridLayout_itemAt_default, this.h, uintptr(index))
	runtime.KeepAlive(this)
	return bindQLayoutItem(__r)
}

func (this *QGridLayout) TakeAtDefault(index int) *QLayoutItem {
	__r, _, _ := purego.SyscallN(addr_QGridLayout_takeAt_default, this.h, uintptr(index))
	runtime.KeepAlive(this)
	return bindQLayoutItem(__r)
}

func (this *QGridLayout) CountDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QGridLayout_count_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QGridLayout) SetGeometryDefault(arg1 *QRect) {
	purego.SyscallN(addr_QGridLayout_setGeometry_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QGridLayout) GeometryDefault() *QRect {
	__r, _, _ := purego.SyscallN(addr_QGridLayout_geometry_default, this.h)
	runtime.KeepAlive(this)
	return newownedQRect(__r)
}

func (this *QGridLayout) IndexOfDefault(arg1 *QWidget) int {
	__r0, _, _ := purego.SyscallN(addr_QGridLayout_indexOf_default, this.h, uintptr(arg1.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return int(__r)
}

func (this *QGridLayout) IsEmptyDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QGridLayout_isEmpty_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGridLayout) ControlTypesDefault() QSizePolicy__ControlType {
	__r0, _, _ := purego.SyscallN(addr_QGridLayout_controlTypes_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QSizePolicy__ControlType(__r)
}

func (this *QGridLayout) ReplaceWidgetDefault(from *QWidget, to *QWidget, options FindChildOption) *QLayoutItem {
	__r, _, _ := purego.SyscallN(addr_QGridLayout_replaceWidget_default, this.h, uintptr(from.UnsafePointer()), uintptr(to.UnsafePointer()), uintptr(options))
	runtime.KeepAlive(this)
	runtime.KeepAlive(from)
	runtime.KeepAlive(to)
	return bindQLayoutItem(__r)
}

func (this *QGridLayout) LayoutDefault() *QLayout {
	__r, _, _ := purego.SyscallN(addr_QGridLayout_layout_default, this.h)
	runtime.KeepAlive(this)
	return bindQLayout(__r)
}

func (this *QGridLayout) ChildEventDefault(e *QChildEvent) {
	purego.SyscallN(addr_QGridLayout_childEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QGridLayout) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QGridLayout_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QGridLayout) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QGridLayout_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QGridLayout) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QGridLayout_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGridLayout) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QGridLayout_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGridLayout) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QGridLayout_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QGridLayout) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QGridLayout_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QGridLayout) WidgetDefault() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QGridLayout_widget_default, this.h)
	runtime.KeepAlive(this)
	return bindQWidget(__r)
}

func (this *QGridLayout) SpacerItemDefault() *QSpacerItem {
	__r, _, _ := purego.SyscallN(addr_QGridLayout_spacerItem_default, this.h)
	runtime.KeepAlive(this)
	return bindQSpacerItem(__r)
}

func (this *QGridLayout) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QGridLayout_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGridLayout) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QGridLayout_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGridLayout) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QGridLayout_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGridLayout) Disconnect(handle int64) {
	pg_QGridLayout__disconnect(handle)
	unregisterCallback(handle)
}

func QGridLayout_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case 729690261457684101:
		return 1
	case -7246748916541192690:
		return 2
	case 4158099606088301136:
		return 3
	case 4039479028406881496:
		return 4
	case -4587105490415799682:
		return 5
	case -8266762397995423915:
		return 6
	case 4311534060495737339:
		return 7
	case 5037837229528203129:
		return 8
	case 415665469015095569:
		return 9
	case 6991341317990632368:
		return 10
	case 7508512214013391107:
		return 11
	case 7549889108533486643:
		return 12
	case -5627843058667511180:
		return 13
	case 5684820848726969227:
		return 14
	case 6884465960736393867:
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

func QGridLayout_enable(h uintptr, idx int32) { pg_QGridLayout__set_override(h, idx) }

func (this *QGridLayout) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QGridLayout) OnSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 729690261457684101, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sizeHint"
}

func (this *QGridLayout) OnMinimumSize(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), -7246748916541192690, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "minimumSize"
}

func (this *QGridLayout) OnMaximumSize(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 4158099606088301136, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "maximumSize"
}

func (this *QGridLayout) OnSetSpacing(fn func(int)) {
	registerVirt(this.UnsafeGoID(), 4039479028406881496, func(spacing int32) { fn(int(spacing)) }) // slot "setSpacing"
}

func (this *QGridLayout) OnSpacing(fn func() int) {
	registerVirt(this.UnsafeGoID(), -4587105490415799682, func() int32 { return int32(fn()) }) // slot "spacing"
}

func (this *QGridLayout) OnHasHeightForWidth(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -8266762397995423915, func() int32 { return b2i(fn()) }) // slot "hasHeightForWidth"
}

func (this *QGridLayout) OnHeightForWidth(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), 4311534060495737339, func(arg1 int32) int32 { return int32(fn(int(arg1))) }) // slot "heightForWidth"
}

func (this *QGridLayout) OnMinimumHeightForWidth(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), 5037837229528203129, func(arg1 int32) int32 { return int32(fn(int(arg1))) }) // slot "minimumHeightForWidth"
}

func (this *QGridLayout) OnExpandingDirections(fn func() Orientation) {
	registerVirt(this.UnsafeGoID(), 415665469015095569, func() int32 { return int32(fn()) }) // slot "expandingDirections"
}

func (this *QGridLayout) OnInvalidate(fn func()) {
	registerVirt(this.UnsafeGoID(), 6991341317990632368, func() { fn() }) // slot "invalidate"
}

func (this *QGridLayout) OnItemAt(fn func(int) *QLayoutItem) {
	registerVirt(this.UnsafeGoID(), 7508512214013391107, func(index int32) uintptr { return uintptr(fn(int(index)).UnsafePointer()) }) // slot "itemAt"
}

func (this *QGridLayout) OnTakeAt(fn func(int) *QLayoutItem) {
	registerVirt(this.UnsafeGoID(), 7549889108533486643, func(index int32) uintptr { return uintptr(fn(int(index)).UnsafePointer()) }) // slot "takeAt"
}

func (this *QGridLayout) OnCount(fn func() int) {
	registerVirt(this.UnsafeGoID(), -5627843058667511180, func() int32 { return int32(fn()) }) // slot "count"
}

func (this *QGridLayout) OnSetGeometry(fn func(*QRect)) {
	registerVirt(this.UnsafeGoID(), 5684820848726969227, func(arg1 uintptr) { fn(UnsafeNewQRect(unsafe.Pointer(arg1))) }) // slot "setGeometry"
}

func (this *QGridLayout) OnAddItem(fn func(*QLayoutItem)) {
	registerVirt(this.UnsafeGoID(), 6884465960736393867, func(arg1 uintptr) { fn(UnsafeNewQLayoutItem(unsafe.Pointer(arg1))) }) // slot "addItem"
}

func (this *QGridLayout) OnGeometry(fn func() *QRect) {
	registerVirt(this.UnsafeGoID(), 1050490530495513297, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "geometry"
}

func (this *QGridLayout) OnIndexOf(fn func(*QWidget) int) {
	registerVirt(this.UnsafeGoID(), -7258473138030329570, func(arg1 uintptr) int32 { return int32(fn(UnsafeNewQWidget(unsafe.Pointer(arg1)))) }) // slot "indexOf"
}

func (this *QGridLayout) OnIndexOf2(fn func(*QLayoutItem) int) {
	registerVirt(this.UnsafeGoID(), -3276435765994963516, func(arg1 uintptr) int32 { return int32(fn(UnsafeNewQLayoutItem(unsafe.Pointer(arg1)))) }) // slot "indexOf2"
}

func (this *QGridLayout) OnIsEmpty(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -626822430251098498, func() int32 { return b2i(fn()) }) // slot "isEmpty"
}

func (this *QGridLayout) OnControlTypes(fn func() QSizePolicy__ControlType) {
	registerVirt(this.UnsafeGoID(), 3278947275714896887, func() int32 { return int32(fn()) }) // slot "controlTypes"
}

func (this *QGridLayout) OnReplaceWidget(fn func(*QWidget, *QWidget, FindChildOption) *QLayoutItem) {
	registerVirt(this.UnsafeGoID(), -712862742159708061, func(from uintptr, to uintptr, options int32) uintptr { return uintptr(fn(UnsafeNewQWidget(unsafe.Pointer(from)), UnsafeNewQWidget(unsafe.Pointer(to)), FindChildOption(options)).UnsafePointer()) }) // slot "replaceWidget"
}

func (this *QGridLayout) OnLayout(fn func() *QLayout) {
	registerVirt(this.UnsafeGoID(), -7183390618627992737, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "layout"
}

func (this *QGridLayout) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(e uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(e))) }) // slot "childEvent"
}

func (this *QGridLayout) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QGridLayout) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QGridLayout) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QGridLayout) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QGridLayout) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QGridLayout) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QGridLayout) OnWidget(fn func() *QWidget) {
	registerVirt(this.UnsafeGoID(), -4575937485689502171, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "widget"
}

func (this *QGridLayout) OnSpacerItem(fn func() *QSpacerItem) {
	registerVirt(this.UnsafeGoID(), 2272240669359776278, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "spacerItem"
}

func (this *QGridLayout) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QGridLayout_delete(this.h)
	}
	this.UnsafeRelease()
}
