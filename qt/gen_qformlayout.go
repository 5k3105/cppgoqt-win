package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QFormLayout_chain uintptr
	addr_QFormLayout_new uintptr
	addr_QFormLayout_new2 uintptr
	addr_QFormLayout_tr uintptr
	addr_QFormLayout_setFieldGrowthPolicy uintptr
	addr_QFormLayout_fieldGrowthPolicy uintptr
	addr_QFormLayout_setRowWrapPolicy uintptr
	addr_QFormLayout_rowWrapPolicy uintptr
	addr_QFormLayout_setLabelAlignment uintptr
	addr_QFormLayout_labelAlignment uintptr
	addr_QFormLayout_setFormAlignment uintptr
	addr_QFormLayout_formAlignment uintptr
	addr_QFormLayout_setHorizontalSpacing uintptr
	addr_QFormLayout_horizontalSpacing uintptr
	addr_QFormLayout_setVerticalSpacing uintptr
	addr_QFormLayout_verticalSpacing uintptr
	addr_QFormLayout_addRow uintptr
	addr_QFormLayout_addRow2 uintptr
	addr_QFormLayout_addRow3 uintptr
	addr_QFormLayout_addRow4 uintptr
	addr_QFormLayout_addRow5 uintptr
	addr_QFormLayout_addRow6 uintptr
	addr_QFormLayout_insertRow uintptr
	addr_QFormLayout_insertRow2 uintptr
	addr_QFormLayout_insertRow3 uintptr
	addr_QFormLayout_insertRow4 uintptr
	addr_QFormLayout_insertRow5 uintptr
	addr_QFormLayout_insertRow6 uintptr
	addr_QFormLayout_removeRow uintptr
	addr_QFormLayout_removeRow2 uintptr
	addr_QFormLayout_removeRow3 uintptr
	addr_QFormLayout_takeRow uintptr
	addr_QFormLayout_takeRow2 uintptr
	addr_QFormLayout_takeRow3 uintptr
	addr_QFormLayout_setItem uintptr
	addr_QFormLayout_setWidget uintptr
	addr_QFormLayout_setLayout uintptr
	addr_QFormLayout_setRowVisible uintptr
	addr_QFormLayout_setRowVisible2 uintptr
	addr_QFormLayout_setRowVisible3 uintptr
	addr_QFormLayout_isRowVisible uintptr
	addr_QFormLayout_isRowVisible2 uintptr
	addr_QFormLayout_isRowVisible3 uintptr
	addr_QFormLayout_itemAt uintptr
	addr_QFormLayout_getItemPosition uintptr
	addr_QFormLayout_getWidgetPosition uintptr
	addr_QFormLayout_getLayoutPosition uintptr
	addr_QFormLayout_labelForField uintptr
	addr_QFormLayout_labelForField2 uintptr
	addr_QFormLayout_itemAt2 uintptr
	addr_QFormLayout_rowCount uintptr
	addr_QFormLayout_closestAcceptableSize uintptr
	addr_QFormLayout_widgetEvent uintptr
	addr_QFormLayout_addChildLayout uintptr
	addr_QFormLayout_addChildWidget uintptr
	addr_QFormLayout_adoptLayout uintptr
	addr_QFormLayout_alignmentRect uintptr
	addr_QFormLayout_connect uintptr
	addr_QFormLayout_connect2 uintptr
	addr_QFormLayout_sender uintptr
	addr_QFormLayout_senderSignalIndex uintptr
	addr_QFormLayout_receivers uintptr
	addr_QFormLayout_isSignalConnected uintptr
	addr_QFormLayout_tr2 uintptr
	addr_QFormLayout_tr3 uintptr
	addr_QFormLayout_connect3 uintptr
	addr_QFormLayout_connect4 uintptr
	addr_QFormLayout_itemAt_default uintptr
	addr_QFormLayout_itemAt_default2 uintptr
	addr_QFormLayout_metaObject_default uintptr
	addr_QFormLayout_spacing_default uintptr
	addr_QFormLayout_setSpacing_default uintptr
	addr_QFormLayout_addItem_default uintptr
	addr_QFormLayout_takeAt_default uintptr
	addr_QFormLayout_setGeometry_default uintptr
	addr_QFormLayout_minimumSize_default uintptr
	addr_QFormLayout_sizeHint_default uintptr
	addr_QFormLayout_invalidate_default uintptr
	addr_QFormLayout_hasHeightForWidth_default uintptr
	addr_QFormLayout_heightForWidth_default uintptr
	addr_QFormLayout_expandingDirections_default uintptr
	addr_QFormLayout_count_default uintptr
	addr_QFormLayout_geometry_default uintptr
	addr_QFormLayout_maximumSize_default uintptr
	addr_QFormLayout_indexOf_default uintptr
	addr_QFormLayout_isEmpty_default uintptr
	addr_QFormLayout_controlTypes_default uintptr
	addr_QFormLayout_replaceWidget_default uintptr
	addr_QFormLayout_layout_default uintptr
	addr_QFormLayout_childEvent_default uintptr
	addr_QFormLayout_event_default uintptr
	addr_QFormLayout_eventFilter_default uintptr
	addr_QFormLayout_timerEvent_default uintptr
	addr_QFormLayout_customEvent_default uintptr
	addr_QFormLayout_connectNotify_default uintptr
	addr_QFormLayout_disconnectNotify_default uintptr
	addr_QFormLayout_minimumHeightForWidth_default uintptr
	addr_QFormLayout_widget_default uintptr
	addr_QFormLayout_spacerItem_default uintptr
	pg_QFormLayout_connect_destroyed func(uintptr, int64)
	pg_QFormLayout_connect_destroyed2 func(uintptr, int64)
	pg_QFormLayout_connect_objectNameChanged func(uintptr, int64)
	pg_QFormLayout__disconnect func(int64)
	pg_QFormLayout__set_override func(uintptr, int32)
	pg_QFormLayout_delete func(uintptr)
)

func init() {
	symaddr(&addr_QFormLayout_chain, "QFormLayout_chain")
	symaddr(&addr_QFormLayout_new, "QFormLayout_new")
	symaddr(&addr_QFormLayout_new2, "QFormLayout_new2")
	symaddr(&addr_QFormLayout_tr, "QFormLayout_tr")
	symaddr(&addr_QFormLayout_setFieldGrowthPolicy, "QFormLayout_setFieldGrowthPolicy")
	symaddr(&addr_QFormLayout_fieldGrowthPolicy, "QFormLayout_fieldGrowthPolicy")
	symaddr(&addr_QFormLayout_setRowWrapPolicy, "QFormLayout_setRowWrapPolicy")
	symaddr(&addr_QFormLayout_rowWrapPolicy, "QFormLayout_rowWrapPolicy")
	symaddr(&addr_QFormLayout_setLabelAlignment, "QFormLayout_setLabelAlignment")
	symaddr(&addr_QFormLayout_labelAlignment, "QFormLayout_labelAlignment")
	symaddr(&addr_QFormLayout_setFormAlignment, "QFormLayout_setFormAlignment")
	symaddr(&addr_QFormLayout_formAlignment, "QFormLayout_formAlignment")
	symaddr(&addr_QFormLayout_setHorizontalSpacing, "QFormLayout_setHorizontalSpacing")
	symaddr(&addr_QFormLayout_horizontalSpacing, "QFormLayout_horizontalSpacing")
	symaddr(&addr_QFormLayout_setVerticalSpacing, "QFormLayout_setVerticalSpacing")
	symaddr(&addr_QFormLayout_verticalSpacing, "QFormLayout_verticalSpacing")
	symaddr(&addr_QFormLayout_addRow, "QFormLayout_addRow")
	symaddr(&addr_QFormLayout_addRow2, "QFormLayout_addRow2")
	symaddr(&addr_QFormLayout_addRow3, "QFormLayout_addRow3")
	symaddr(&addr_QFormLayout_addRow4, "QFormLayout_addRow4")
	symaddr(&addr_QFormLayout_addRow5, "QFormLayout_addRow5")
	symaddr(&addr_QFormLayout_addRow6, "QFormLayout_addRow6")
	symaddr(&addr_QFormLayout_insertRow, "QFormLayout_insertRow")
	symaddr(&addr_QFormLayout_insertRow2, "QFormLayout_insertRow2")
	symaddr(&addr_QFormLayout_insertRow3, "QFormLayout_insertRow3")
	symaddr(&addr_QFormLayout_insertRow4, "QFormLayout_insertRow4")
	symaddr(&addr_QFormLayout_insertRow5, "QFormLayout_insertRow5")
	symaddr(&addr_QFormLayout_insertRow6, "QFormLayout_insertRow6")
	symaddr(&addr_QFormLayout_removeRow, "QFormLayout_removeRow")
	symaddr(&addr_QFormLayout_removeRow2, "QFormLayout_removeRow2")
	symaddr(&addr_QFormLayout_removeRow3, "QFormLayout_removeRow3")
	symaddr(&addr_QFormLayout_takeRow, "QFormLayout_takeRow")
	symaddr(&addr_QFormLayout_takeRow2, "QFormLayout_takeRow2")
	symaddr(&addr_QFormLayout_takeRow3, "QFormLayout_takeRow3")
	symaddr(&addr_QFormLayout_setItem, "QFormLayout_setItem")
	symaddr(&addr_QFormLayout_setWidget, "QFormLayout_setWidget")
	symaddr(&addr_QFormLayout_setLayout, "QFormLayout_setLayout")
	symaddr(&addr_QFormLayout_setRowVisible, "QFormLayout_setRowVisible")
	symaddr(&addr_QFormLayout_setRowVisible2, "QFormLayout_setRowVisible2")
	symaddr(&addr_QFormLayout_setRowVisible3, "QFormLayout_setRowVisible3")
	symaddr(&addr_QFormLayout_isRowVisible, "QFormLayout_isRowVisible")
	symaddr(&addr_QFormLayout_isRowVisible2, "QFormLayout_isRowVisible2")
	symaddr(&addr_QFormLayout_isRowVisible3, "QFormLayout_isRowVisible3")
	symaddr(&addr_QFormLayout_itemAt, "QFormLayout_itemAt")
	symaddr(&addr_QFormLayout_getItemPosition, "QFormLayout_getItemPosition")
	symaddr(&addr_QFormLayout_getWidgetPosition, "QFormLayout_getWidgetPosition")
	symaddr(&addr_QFormLayout_getLayoutPosition, "QFormLayout_getLayoutPosition")
	symaddr(&addr_QFormLayout_labelForField, "QFormLayout_labelForField")
	symaddr(&addr_QFormLayout_labelForField2, "QFormLayout_labelForField2")
	symaddr(&addr_QFormLayout_itemAt2, "QFormLayout_itemAt2")
	symaddr(&addr_QFormLayout_rowCount, "QFormLayout_rowCount")
	symaddr(&addr_QFormLayout_closestAcceptableSize, "QFormLayout_closestAcceptableSize")
	symaddr(&addr_QFormLayout_widgetEvent, "QFormLayout_widgetEvent")
	symaddr(&addr_QFormLayout_addChildLayout, "QFormLayout_addChildLayout")
	symaddr(&addr_QFormLayout_addChildWidget, "QFormLayout_addChildWidget")
	symaddr(&addr_QFormLayout_adoptLayout, "QFormLayout_adoptLayout")
	symaddr(&addr_QFormLayout_alignmentRect, "QFormLayout_alignmentRect")
	symaddr(&addr_QFormLayout_connect, "QFormLayout_connect")
	symaddr(&addr_QFormLayout_connect2, "QFormLayout_connect2")
	symaddr(&addr_QFormLayout_sender, "QFormLayout_sender")
	symaddr(&addr_QFormLayout_senderSignalIndex, "QFormLayout_senderSignalIndex")
	symaddr(&addr_QFormLayout_receivers, "QFormLayout_receivers")
	symaddr(&addr_QFormLayout_isSignalConnected, "QFormLayout_isSignalConnected")
	symaddr(&addr_QFormLayout_tr2, "QFormLayout_tr2")
	symaddr(&addr_QFormLayout_tr3, "QFormLayout_tr3")
	symaddr(&addr_QFormLayout_connect3, "QFormLayout_connect3")
	symaddr(&addr_QFormLayout_connect4, "QFormLayout_connect4")
	symaddr(&addr_QFormLayout_itemAt_default, "QFormLayout_itemAt_default")
	symaddr(&addr_QFormLayout_itemAt_default2, "QFormLayout_itemAt_default2")
	symaddr(&addr_QFormLayout_metaObject_default, "QFormLayout_metaObject_default")
	symaddr(&addr_QFormLayout_spacing_default, "QFormLayout_spacing_default")
	symaddr(&addr_QFormLayout_setSpacing_default, "QFormLayout_setSpacing_default")
	symaddr(&addr_QFormLayout_addItem_default, "QFormLayout_addItem_default")
	symaddr(&addr_QFormLayout_takeAt_default, "QFormLayout_takeAt_default")
	symaddr(&addr_QFormLayout_setGeometry_default, "QFormLayout_setGeometry_default")
	symaddr(&addr_QFormLayout_minimumSize_default, "QFormLayout_minimumSize_default")
	symaddr(&addr_QFormLayout_sizeHint_default, "QFormLayout_sizeHint_default")
	symaddr(&addr_QFormLayout_invalidate_default, "QFormLayout_invalidate_default")
	symaddr(&addr_QFormLayout_hasHeightForWidth_default, "QFormLayout_hasHeightForWidth_default")
	symaddr(&addr_QFormLayout_heightForWidth_default, "QFormLayout_heightForWidth_default")
	symaddr(&addr_QFormLayout_expandingDirections_default, "QFormLayout_expandingDirections_default")
	symaddr(&addr_QFormLayout_count_default, "QFormLayout_count_default")
	symaddr(&addr_QFormLayout_geometry_default, "QFormLayout_geometry_default")
	symaddr(&addr_QFormLayout_maximumSize_default, "QFormLayout_maximumSize_default")
	symaddr(&addr_QFormLayout_indexOf_default, "QFormLayout_indexOf_default")
	symaddr(&addr_QFormLayout_isEmpty_default, "QFormLayout_isEmpty_default")
	symaddr(&addr_QFormLayout_controlTypes_default, "QFormLayout_controlTypes_default")
	symaddr(&addr_QFormLayout_replaceWidget_default, "QFormLayout_replaceWidget_default")
	symaddr(&addr_QFormLayout_layout_default, "QFormLayout_layout_default")
	symaddr(&addr_QFormLayout_childEvent_default, "QFormLayout_childEvent_default")
	symaddr(&addr_QFormLayout_event_default, "QFormLayout_event_default")
	symaddr(&addr_QFormLayout_eventFilter_default, "QFormLayout_eventFilter_default")
	symaddr(&addr_QFormLayout_timerEvent_default, "QFormLayout_timerEvent_default")
	symaddr(&addr_QFormLayout_customEvent_default, "QFormLayout_customEvent_default")
	symaddr(&addr_QFormLayout_connectNotify_default, "QFormLayout_connectNotify_default")
	symaddr(&addr_QFormLayout_disconnectNotify_default, "QFormLayout_disconnectNotify_default")
	symaddr(&addr_QFormLayout_minimumHeightForWidth_default, "QFormLayout_minimumHeightForWidth_default")
	symaddr(&addr_QFormLayout_widget_default, "QFormLayout_widget_default")
	symaddr(&addr_QFormLayout_spacerItem_default, "QFormLayout_spacerItem_default")
	reg(&pg_QFormLayout_connect_destroyed, "QFormLayout_connect_destroyed")
	reg(&pg_QFormLayout_connect_destroyed2, "QFormLayout_connect_destroyed2")
	reg(&pg_QFormLayout_connect_objectNameChanged, "QFormLayout_connect_objectNameChanged")
	reg(&pg_QFormLayout__disconnect, "QFormLayout__disconnect")
	reg(&pg_QFormLayout__set_override, "QFormLayout__set_override")
	reg(&pg_QFormLayout_delete, "QFormLayout_delete")
}

type QFormLayout struct {
	h uintptr
	*QLayout
}

func bindQFormLayout(h uintptr) *QFormLayout {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QFormLayout_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QFormLayout{h: h, QLayout: &QLayout{h: __c[0], QObject: &QObject{h: __c[1]}}}
}

func newownedQFormLayout(h uintptr) *QFormLayout { return bindQFormLayout(h) }

func (this *QFormLayout) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QFormLayout) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQFormLayout(h unsafe.Pointer) *QFormLayout { return bindQFormLayout(uintptr(h)) }
func UnsafeNewQFormLayoutOwned(h unsafe.Pointer) *QFormLayout { return newownedQFormLayout(uintptr(h)) }

func (this *QFormLayout) UnsafeRelease() {
	this.h = 0
	if this.QLayout != nil {
		this.QLayout.UnsafeRelease()
	}
}

func NewQFormLayout(parent *QWidget) *QFormLayout {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QFormLayout_new, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQFormLayout(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QFormLayout_virtIdx, QFormLayout_enable)
	return __o
}

func NewQFormLayout2() *QFormLayout {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QFormLayout_new2, uintptr(__goid))
	__o := newownedQFormLayout(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QFormLayout_virtIdx, QFormLayout_enable)
	return __o
}

func QFormLayout_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFormLayout_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QFormLayout) SetFieldGrowthPolicy(policy QFormLayout__FieldGrowthPolicy) {
	purego.SyscallN(addr_QFormLayout_setFieldGrowthPolicy, this.h, uintptr(policy))
	runtime.KeepAlive(this)
}

func (this *QFormLayout) FieldGrowthPolicy() QFormLayout__FieldGrowthPolicy {
	__r0, _, _ := purego.SyscallN(addr_QFormLayout_fieldGrowthPolicy, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QFormLayout__FieldGrowthPolicy(__r)
}

func (this *QFormLayout) SetRowWrapPolicy(policy QFormLayout__RowWrapPolicy) {
	purego.SyscallN(addr_QFormLayout_setRowWrapPolicy, this.h, uintptr(policy))
	runtime.KeepAlive(this)
}

func (this *QFormLayout) RowWrapPolicy() QFormLayout__RowWrapPolicy {
	__r0, _, _ := purego.SyscallN(addr_QFormLayout_rowWrapPolicy, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QFormLayout__RowWrapPolicy(__r)
}

func (this *QFormLayout) SetLabelAlignment(alignment AlignmentFlag) {
	purego.SyscallN(addr_QFormLayout_setLabelAlignment, this.h, uintptr(alignment))
	runtime.KeepAlive(this)
}

func (this *QFormLayout) LabelAlignment() AlignmentFlag {
	__r0, _, _ := purego.SyscallN(addr_QFormLayout_labelAlignment, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return AlignmentFlag(__r)
}

func (this *QFormLayout) SetFormAlignment(alignment AlignmentFlag) {
	purego.SyscallN(addr_QFormLayout_setFormAlignment, this.h, uintptr(alignment))
	runtime.KeepAlive(this)
}

func (this *QFormLayout) FormAlignment() AlignmentFlag {
	__r0, _, _ := purego.SyscallN(addr_QFormLayout_formAlignment, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return AlignmentFlag(__r)
}

func (this *QFormLayout) SetHorizontalSpacing(spacing int) {
	purego.SyscallN(addr_QFormLayout_setHorizontalSpacing, this.h, uintptr(spacing))
	runtime.KeepAlive(this)
}

func (this *QFormLayout) HorizontalSpacing() int {
	__r0, _, _ := purego.SyscallN(addr_QFormLayout_horizontalSpacing, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QFormLayout) SetVerticalSpacing(spacing int) {
	purego.SyscallN(addr_QFormLayout_setVerticalSpacing, this.h, uintptr(spacing))
	runtime.KeepAlive(this)
}

func (this *QFormLayout) VerticalSpacing() int {
	__r0, _, _ := purego.SyscallN(addr_QFormLayout_verticalSpacing, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QFormLayout) AddRow(label *QWidget, field *QWidget) {
	purego.SyscallN(addr_QFormLayout_addRow, this.h, uintptr(label.UnsafePointer()), uintptr(field.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(label)
	runtime.KeepAlive(field)
}

func (this *QFormLayout) AddRow2(label *QWidget, field *QLayout) {
	purego.SyscallN(addr_QFormLayout_addRow2, this.h, uintptr(label.UnsafePointer()), uintptr(field.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(label)
	runtime.KeepAlive(field)
}

func (this *QFormLayout) AddRow3(labelText string, field *QWidget) {
	var labelText_p uintptr
	if len(labelText) > 0 {
		labelText_p = uintptr(unsafe.Pointer(unsafe.StringData(labelText)))
	}
	purego.SyscallN(addr_QFormLayout_addRow3, this.h, labelText_p, uintptr(len(labelText)), uintptr(field.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(labelText)
	runtime.KeepAlive(field)
}

func (this *QFormLayout) AddRow4(labelText string, field *QLayout) {
	var labelText_p uintptr
	if len(labelText) > 0 {
		labelText_p = uintptr(unsafe.Pointer(unsafe.StringData(labelText)))
	}
	purego.SyscallN(addr_QFormLayout_addRow4, this.h, labelText_p, uintptr(len(labelText)), uintptr(field.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(labelText)
	runtime.KeepAlive(field)
}

func (this *QFormLayout) AddRowWithWidget(widget *QWidget) {
	purego.SyscallN(addr_QFormLayout_addRow5, this.h, uintptr(widget.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(widget)
}

func (this *QFormLayout) AddRowWithLayout(layout *QLayout) {
	purego.SyscallN(addr_QFormLayout_addRow6, this.h, uintptr(layout.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(layout)
}

func (this *QFormLayout) InsertRow(row int, label *QWidget, field *QWidget) {
	purego.SyscallN(addr_QFormLayout_insertRow, this.h, uintptr(row), uintptr(label.UnsafePointer()), uintptr(field.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(label)
	runtime.KeepAlive(field)
}

func (this *QFormLayout) InsertRow2(row int, label *QWidget, field *QLayout) {
	purego.SyscallN(addr_QFormLayout_insertRow2, this.h, uintptr(row), uintptr(label.UnsafePointer()), uintptr(field.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(label)
	runtime.KeepAlive(field)
}

func (this *QFormLayout) InsertRow3(row int, labelText string, field *QWidget) {
	var labelText_p uintptr
	if len(labelText) > 0 {
		labelText_p = uintptr(unsafe.Pointer(unsafe.StringData(labelText)))
	}
	purego.SyscallN(addr_QFormLayout_insertRow3, this.h, uintptr(row), labelText_p, uintptr(len(labelText)), uintptr(field.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(labelText)
	runtime.KeepAlive(field)
}

func (this *QFormLayout) InsertRow4(row int, labelText string, field *QLayout) {
	var labelText_p uintptr
	if len(labelText) > 0 {
		labelText_p = uintptr(unsafe.Pointer(unsafe.StringData(labelText)))
	}
	purego.SyscallN(addr_QFormLayout_insertRow4, this.h, uintptr(row), labelText_p, uintptr(len(labelText)), uintptr(field.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(labelText)
	runtime.KeepAlive(field)
}

func (this *QFormLayout) InsertRow5(row int, widget *QWidget) {
	purego.SyscallN(addr_QFormLayout_insertRow5, this.h, uintptr(row), uintptr(widget.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(widget)
}

func (this *QFormLayout) InsertRow6(row int, layout *QLayout) {
	purego.SyscallN(addr_QFormLayout_insertRow6, this.h, uintptr(row), uintptr(layout.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(layout)
}

func (this *QFormLayout) RemoveRow(row int) {
	purego.SyscallN(addr_QFormLayout_removeRow, this.h, uintptr(row))
	runtime.KeepAlive(this)
}

func (this *QFormLayout) RemoveRowWithWidget(widget *QWidget) {
	purego.SyscallN(addr_QFormLayout_removeRow2, this.h, uintptr(widget.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(widget)
}

func (this *QFormLayout) RemoveRowWithLayout(layout *QLayout) {
	purego.SyscallN(addr_QFormLayout_removeRow3, this.h, uintptr(layout.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(layout)
}

func (this *QFormLayout) TakeRow(row int) *QFormLayout__TakeRowResult {
	__r, _, _ := purego.SyscallN(addr_QFormLayout_takeRow, this.h, uintptr(row))
	runtime.KeepAlive(this)
	return newownedQFormLayout__TakeRowResult(__r)
}

func (this *QFormLayout) TakeRowWithWidget(widget *QWidget) *QFormLayout__TakeRowResult {
	__r, _, _ := purego.SyscallN(addr_QFormLayout_takeRow2, this.h, uintptr(widget.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(widget)
	return newownedQFormLayout__TakeRowResult(__r)
}

func (this *QFormLayout) TakeRowWithLayout(layout *QLayout) *QFormLayout__TakeRowResult {
	__r, _, _ := purego.SyscallN(addr_QFormLayout_takeRow3, this.h, uintptr(layout.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(layout)
	return newownedQFormLayout__TakeRowResult(__r)
}

func (this *QFormLayout) SetItem(row int, role QFormLayout__ItemRole, item *QLayoutItem) {
	purego.SyscallN(addr_QFormLayout_setItem, this.h, uintptr(row), uintptr(role), uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QFormLayout) SetWidget(row int, role QFormLayout__ItemRole, widget *QWidget) {
	purego.SyscallN(addr_QFormLayout_setWidget, this.h, uintptr(row), uintptr(role), uintptr(widget.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(widget)
}

func (this *QFormLayout) SetLayout(row int, role QFormLayout__ItemRole, layout *QLayout) {
	purego.SyscallN(addr_QFormLayout_setLayout, this.h, uintptr(row), uintptr(role), uintptr(layout.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(layout)
}

func (this *QFormLayout) SetRowVisible(row int, on bool) {
	purego.SyscallN(addr_QFormLayout_setRowVisible, this.h, uintptr(row), uintptr(b2i(on)))
	runtime.KeepAlive(this)
}

func (this *QFormLayout) SetRowVisible2(widget *QWidget, on bool) {
	purego.SyscallN(addr_QFormLayout_setRowVisible2, this.h, uintptr(widget.UnsafePointer()), uintptr(b2i(on)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(widget)
}

func (this *QFormLayout) SetRowVisible3(layout *QLayout, on bool) {
	purego.SyscallN(addr_QFormLayout_setRowVisible3, this.h, uintptr(layout.UnsafePointer()), uintptr(b2i(on)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(layout)
}

func (this *QFormLayout) IsRowVisible(row int) bool {
	__r0, _, _ := purego.SyscallN(addr_QFormLayout_isRowVisible, this.h, uintptr(row))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFormLayout) IsRowVisibleWithWidget(widget *QWidget) bool {
	__r0, _, _ := purego.SyscallN(addr_QFormLayout_isRowVisible2, this.h, uintptr(widget.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(widget)
	return __r != 0
}

func (this *QFormLayout) IsRowVisibleWithLayout(layout *QLayout) bool {
	__r0, _, _ := purego.SyscallN(addr_QFormLayout_isRowVisible3, this.h, uintptr(layout.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(layout)
	return __r != 0
}

func (this *QFormLayout) ItemAt(row int, role QFormLayout__ItemRole) *QLayoutItem {
	__r, _, _ := purego.SyscallN(addr_QFormLayout_itemAt, this.h, uintptr(row), uintptr(role))
	runtime.KeepAlive(this)
	return bindQLayoutItem(__r)
}

func (this *QFormLayout) GetItemPosition(index int) (int, QFormLayout__ItemRole) {
	var rowPtr_out int32
	var rolePtr_out int32
	purego.SyscallN(addr_QFormLayout_getItemPosition, this.h, uintptr(index), uintptr(unsafe.Pointer(&rowPtr_out)), uintptr(unsafe.Pointer(&rolePtr_out)))
	runtime.KeepAlive(this)
	return int(rowPtr_out), QFormLayout__ItemRole(rolePtr_out)
}

func (this *QFormLayout) GetWidgetPosition(widget *QWidget) (int, QFormLayout__ItemRole) {
	var rowPtr_out int32
	var rolePtr_out int32
	purego.SyscallN(addr_QFormLayout_getWidgetPosition, this.h, uintptr(widget.UnsafePointer()), uintptr(unsafe.Pointer(&rowPtr_out)), uintptr(unsafe.Pointer(&rolePtr_out)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(widget)
	return int(rowPtr_out), QFormLayout__ItemRole(rolePtr_out)
}

func (this *QFormLayout) GetLayoutPosition(layout *QLayout) (int, QFormLayout__ItemRole) {
	var rowPtr_out int32
	var rolePtr_out int32
	purego.SyscallN(addr_QFormLayout_getLayoutPosition, this.h, uintptr(layout.UnsafePointer()), uintptr(unsafe.Pointer(&rowPtr_out)), uintptr(unsafe.Pointer(&rolePtr_out)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(layout)
	return int(rowPtr_out), QFormLayout__ItemRole(rolePtr_out)
}

func (this *QFormLayout) LabelForField(field *QWidget) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QFormLayout_labelForField, this.h, uintptr(field.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(field)
	return bindQWidget(__r)
}

func (this *QFormLayout) LabelForFieldWithField(field *QLayout) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QFormLayout_labelForField2, this.h, uintptr(field.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(field)
	return bindQWidget(__r)
}

func (this *QFormLayout) ItemAtWithIndex(index int) *QLayoutItem {
	__r, _, _ := purego.SyscallN(addr_QFormLayout_itemAt2, this.h, uintptr(index))
	runtime.KeepAlive(this)
	return bindQLayoutItem(__r)
}

func (this *QFormLayout) RowCount() int {
	__r0, _, _ := purego.SyscallN(addr_QFormLayout_rowCount, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func QFormLayout_ClosestAcceptableSize(w *QWidget, s *QSize) *QSize {
	__r, _, _ := purego.SyscallN(addr_QFormLayout_closestAcceptableSize, uintptr(w.UnsafePointer()), uintptr(s.UnsafePointer()))
	runtime.KeepAlive(w)
	runtime.KeepAlive(s)
	return newownedQSize(__r)
}

func (this *QFormLayout) WidgetEvent(arg1 *QEvent) {
	purego.SyscallN(addr_QFormLayout_widgetEvent, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QFormLayout) AddChildLayout(l *QLayout) {
	purego.SyscallN(addr_QFormLayout_addChildLayout, this.h, uintptr(l.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(l)
}

func (this *QFormLayout) AddChildWidget(w *QWidget) {
	purego.SyscallN(addr_QFormLayout_addChildWidget, this.h, uintptr(w.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(w)
}

func (this *QFormLayout) AdoptLayout(layout *QLayout) bool {
	__r0, _, _ := purego.SyscallN(addr_QFormLayout_adoptLayout, this.h, uintptr(layout.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(layout)
	return __r != 0
}

func (this *QFormLayout) AlignmentRect(arg1 *QRect) *QRect {
	__r, _, _ := purego.SyscallN(addr_QFormLayout_alignmentRect, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return newownedQRect(__r)
}

func QFormLayout_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QFormLayout_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QFormLayout_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QFormLayout_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QFormLayout) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QFormLayout_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QFormLayout) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QFormLayout_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QFormLayout) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QFormLayout_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QFormLayout) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QFormLayout_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QFormLayout_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFormLayout_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QFormLayout_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFormLayout_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QFormLayout_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QFormLayout_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QFormLayout_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QFormLayout_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QFormLayout) ItemAtDefault(row int, role QFormLayout__ItemRole) *QLayoutItem {
	__r, _, _ := purego.SyscallN(addr_QFormLayout_itemAt_default, this.h, uintptr(row), uintptr(role))
	runtime.KeepAlive(this)
	return bindQLayoutItem(__r)
}

func (this *QFormLayout) ItemAtWithIndexDefault(index int) *QLayoutItem {
	__r, _, _ := purego.SyscallN(addr_QFormLayout_itemAt_default2, this.h, uintptr(index))
	runtime.KeepAlive(this)
	return bindQLayoutItem(__r)
}

func (this *QFormLayout) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QFormLayout_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QFormLayout) SpacingDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QFormLayout_spacing_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QFormLayout) SetSpacingDefault(arg1 int) {
	purego.SyscallN(addr_QFormLayout_setSpacing_default, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QFormLayout) AddItemDefault(item *QLayoutItem) {
	purego.SyscallN(addr_QFormLayout_addItem_default, this.h, uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QFormLayout) TakeAtDefault(index int) *QLayoutItem {
	__r, _, _ := purego.SyscallN(addr_QFormLayout_takeAt_default, this.h, uintptr(index))
	runtime.KeepAlive(this)
	return bindQLayoutItem(__r)
}

func (this *QFormLayout) SetGeometryDefault(rect *QRect) {
	purego.SyscallN(addr_QFormLayout_setGeometry_default, this.h, uintptr(rect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
}

func (this *QFormLayout) MinimumSizeDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QFormLayout_minimumSize_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QFormLayout) SizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QFormLayout_sizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QFormLayout) InvalidateDefault() {
	purego.SyscallN(addr_QFormLayout_invalidate_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QFormLayout) HasHeightForWidthDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QFormLayout_hasHeightForWidth_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFormLayout) HeightForWidthDefault(width int) int {
	__r0, _, _ := purego.SyscallN(addr_QFormLayout_heightForWidth_default, this.h, uintptr(width))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QFormLayout) ExpandingDirectionsDefault() Orientation {
	__r0, _, _ := purego.SyscallN(addr_QFormLayout_expandingDirections_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return Orientation(__r)
}

func (this *QFormLayout) CountDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QFormLayout_count_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QFormLayout) GeometryDefault() *QRect {
	__r, _, _ := purego.SyscallN(addr_QFormLayout_geometry_default, this.h)
	runtime.KeepAlive(this)
	return newownedQRect(__r)
}

func (this *QFormLayout) MaximumSizeDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QFormLayout_maximumSize_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QFormLayout) IndexOfDefault(arg1 *QWidget) int {
	__r0, _, _ := purego.SyscallN(addr_QFormLayout_indexOf_default, this.h, uintptr(arg1.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return int(__r)
}

func (this *QFormLayout) IsEmptyDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QFormLayout_isEmpty_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFormLayout) ControlTypesDefault() QSizePolicy__ControlType {
	__r0, _, _ := purego.SyscallN(addr_QFormLayout_controlTypes_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QSizePolicy__ControlType(__r)
}

func (this *QFormLayout) ReplaceWidgetDefault(from *QWidget, to *QWidget, options FindChildOption) *QLayoutItem {
	__r, _, _ := purego.SyscallN(addr_QFormLayout_replaceWidget_default, this.h, uintptr(from.UnsafePointer()), uintptr(to.UnsafePointer()), uintptr(options))
	runtime.KeepAlive(this)
	runtime.KeepAlive(from)
	runtime.KeepAlive(to)
	return bindQLayoutItem(__r)
}

func (this *QFormLayout) LayoutDefault() *QLayout {
	__r, _, _ := purego.SyscallN(addr_QFormLayout_layout_default, this.h)
	runtime.KeepAlive(this)
	return bindQLayout(__r)
}

func (this *QFormLayout) ChildEventDefault(e *QChildEvent) {
	purego.SyscallN(addr_QFormLayout_childEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QFormLayout) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QFormLayout_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QFormLayout) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QFormLayout_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QFormLayout) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QFormLayout_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFormLayout) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QFormLayout_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFormLayout) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QFormLayout_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QFormLayout) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QFormLayout_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QFormLayout) MinimumHeightForWidthDefault(arg1 int) int {
	__r0, _, _ := purego.SyscallN(addr_QFormLayout_minimumHeightForWidth_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QFormLayout) WidgetDefault() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QFormLayout_widget_default, this.h)
	runtime.KeepAlive(this)
	return bindQWidget(__r)
}

func (this *QFormLayout) SpacerItemDefault() *QSpacerItem {
	__r, _, _ := purego.SyscallN(addr_QFormLayout_spacerItem_default, this.h)
	runtime.KeepAlive(this)
	return bindQSpacerItem(__r)
}

func (this *QFormLayout) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QFormLayout_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFormLayout) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QFormLayout_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFormLayout) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QFormLayout_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFormLayout) Disconnect(handle int64) {
	pg_QFormLayout__disconnect(handle)
	unregisterCallback(handle)
}

func QFormLayout_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case -4587105490415799682:
		return 1
	case 4039479028406881496:
		return 2
	case 6884465960736393867:
		return 3
	case 7508512214013391107:
		return 4
	case 7549889108533486643:
		return 5
	case 5684820848726969227:
		return 6
	case -7246748916541192690:
		return 7
	case 729690261457684101:
		return 8
	case 6991341317990632368:
		return 9
	case -8266762397995423915:
		return 10
	case 4311534060495737339:
		return 11
	case 415665469015095569:
		return 12
	case -5627843058667511180:
		return 13
	case 1050490530495513297:
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
	case -4575937485689502171:
		return 30
	case 2272240669359776278:
		return 31
	}
	return -1
}

func QFormLayout_enable(h uintptr, idx int32) { pg_QFormLayout__set_override(h, idx) }

func (this *QFormLayout) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QFormLayout) OnSpacing(fn func() int) {
	registerVirt(this.UnsafeGoID(), -4587105490415799682, func() int32 { return int32(fn()) }) // slot "spacing"
}

func (this *QFormLayout) OnSetSpacing(fn func(int)) {
	registerVirt(this.UnsafeGoID(), 4039479028406881496, func(arg1 int32) { fn(int(arg1)) }) // slot "setSpacing"
}

func (this *QFormLayout) OnAddItem(fn func(*QLayoutItem)) {
	registerVirt(this.UnsafeGoID(), 6884465960736393867, func(item uintptr) { fn(UnsafeNewQLayoutItem(unsafe.Pointer(item))) }) // slot "addItem"
}

func (this *QFormLayout) OnItemAt(fn func(int) *QLayoutItem) {
	registerVirt(this.UnsafeGoID(), 7508512214013391107, func(index int32) uintptr { return uintptr(fn(int(index)).UnsafePointer()) }) // slot "itemAt"
}

func (this *QFormLayout) OnTakeAt(fn func(int) *QLayoutItem) {
	registerVirt(this.UnsafeGoID(), 7549889108533486643, func(index int32) uintptr { return uintptr(fn(int(index)).UnsafePointer()) }) // slot "takeAt"
}

func (this *QFormLayout) OnSetGeometry(fn func(*QRect)) {
	registerVirt(this.UnsafeGoID(), 5684820848726969227, func(rect uintptr) { fn(UnsafeNewQRect(unsafe.Pointer(rect))) }) // slot "setGeometry"
}

func (this *QFormLayout) OnMinimumSize(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), -7246748916541192690, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "minimumSize"
}

func (this *QFormLayout) OnSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 729690261457684101, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sizeHint"
}

func (this *QFormLayout) OnInvalidate(fn func()) {
	registerVirt(this.UnsafeGoID(), 6991341317990632368, func() { fn() }) // slot "invalidate"
}

func (this *QFormLayout) OnHasHeightForWidth(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -8266762397995423915, func() int32 { return b2i(fn()) }) // slot "hasHeightForWidth"
}

func (this *QFormLayout) OnHeightForWidth(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), 4311534060495737339, func(width int32) int32 { return int32(fn(int(width))) }) // slot "heightForWidth"
}

func (this *QFormLayout) OnExpandingDirections(fn func() Orientation) {
	registerVirt(this.UnsafeGoID(), 415665469015095569, func() int32 { return int32(fn()) }) // slot "expandingDirections"
}

func (this *QFormLayout) OnCount(fn func() int) {
	registerVirt(this.UnsafeGoID(), -5627843058667511180, func() int32 { return int32(fn()) }) // slot "count"
}

func (this *QFormLayout) OnGeometry(fn func() *QRect) {
	registerVirt(this.UnsafeGoID(), 1050490530495513297, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "geometry"
}

func (this *QFormLayout) OnMaximumSize(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 4158099606088301136, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "maximumSize"
}

func (this *QFormLayout) OnIndexOf(fn func(*QWidget) int) {
	registerVirt(this.UnsafeGoID(), -7258473138030329570, func(arg1 uintptr) int32 { return int32(fn(UnsafeNewQWidget(unsafe.Pointer(arg1)))) }) // slot "indexOf"
}

func (this *QFormLayout) OnIndexOf2(fn func(*QLayoutItem) int) {
	registerVirt(this.UnsafeGoID(), -3276435765994963516, func(arg1 uintptr) int32 { return int32(fn(UnsafeNewQLayoutItem(unsafe.Pointer(arg1)))) }) // slot "indexOf2"
}

func (this *QFormLayout) OnIsEmpty(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -626822430251098498, func() int32 { return b2i(fn()) }) // slot "isEmpty"
}

func (this *QFormLayout) OnControlTypes(fn func() QSizePolicy__ControlType) {
	registerVirt(this.UnsafeGoID(), 3278947275714896887, func() int32 { return int32(fn()) }) // slot "controlTypes"
}

func (this *QFormLayout) OnReplaceWidget(fn func(*QWidget, *QWidget, FindChildOption) *QLayoutItem) {
	registerVirt(this.UnsafeGoID(), -712862742159708061, func(from uintptr, to uintptr, options int32) uintptr { return uintptr(fn(UnsafeNewQWidget(unsafe.Pointer(from)), UnsafeNewQWidget(unsafe.Pointer(to)), FindChildOption(options)).UnsafePointer()) }) // slot "replaceWidget"
}

func (this *QFormLayout) OnLayout(fn func() *QLayout) {
	registerVirt(this.UnsafeGoID(), -7183390618627992737, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "layout"
}

func (this *QFormLayout) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(e uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(e))) }) // slot "childEvent"
}

func (this *QFormLayout) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QFormLayout) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QFormLayout) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QFormLayout) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QFormLayout) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QFormLayout) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QFormLayout) OnMinimumHeightForWidth(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), 5037837229528203129, func(arg1 int32) int32 { return int32(fn(int(arg1))) }) // slot "minimumHeightForWidth"
}

func (this *QFormLayout) OnWidget(fn func() *QWidget) {
	registerVirt(this.UnsafeGoID(), -4575937485689502171, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "widget"
}

func (this *QFormLayout) OnSpacerItem(fn func() *QSpacerItem) {
	registerVirt(this.UnsafeGoID(), 2272240669359776278, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "spacerItem"
}

func (this *QFormLayout) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QFormLayout_delete(this.h)
	}
	this.UnsafeRelease()
}
