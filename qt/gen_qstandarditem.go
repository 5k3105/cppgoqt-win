package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QStandardItem_new uintptr
	addr_QStandardItem_new2 uintptr
	addr_QStandardItem_new3 uintptr
	addr_QStandardItem_new4 uintptr
	addr_QStandardItem_new5 uintptr
	addr_QStandardItem_data uintptr
	addr_QStandardItem_multiData uintptr
	addr_QStandardItem_setData uintptr
	addr_QStandardItem_clearData uintptr
	addr_QStandardItem_text uintptr
	addr_QStandardItem_setText uintptr
	addr_QStandardItem_icon uintptr
	addr_QStandardItem_setIcon uintptr
	addr_QStandardItem_toolTip uintptr
	addr_QStandardItem_setToolTip uintptr
	addr_QStandardItem_statusTip uintptr
	addr_QStandardItem_setStatusTip uintptr
	addr_QStandardItem_whatsThis uintptr
	addr_QStandardItem_setWhatsThis uintptr
	addr_QStandardItem_sizeHint uintptr
	addr_QStandardItem_setSizeHint uintptr
	addr_QStandardItem_font uintptr
	addr_QStandardItem_setFont uintptr
	addr_QStandardItem_textAlignment uintptr
	addr_QStandardItem_setTextAlignment uintptr
	addr_QStandardItem_background uintptr
	addr_QStandardItem_setBackground uintptr
	addr_QStandardItem_foreground uintptr
	addr_QStandardItem_setForeground uintptr
	addr_QStandardItem_checkState uintptr
	addr_QStandardItem_setCheckState uintptr
	addr_QStandardItem_accessibleText uintptr
	addr_QStandardItem_setAccessibleText uintptr
	addr_QStandardItem_accessibleDescription uintptr
	addr_QStandardItem_setAccessibleDescription uintptr
	addr_QStandardItem_flags uintptr
	addr_QStandardItem_setFlags uintptr
	addr_QStandardItem_isEnabled uintptr
	addr_QStandardItem_setEnabled uintptr
	addr_QStandardItem_isEditable uintptr
	addr_QStandardItem_setEditable uintptr
	addr_QStandardItem_isSelectable uintptr
	addr_QStandardItem_setSelectable uintptr
	addr_QStandardItem_isCheckable uintptr
	addr_QStandardItem_setCheckable uintptr
	addr_QStandardItem_isAutoTristate uintptr
	addr_QStandardItem_setAutoTristate uintptr
	addr_QStandardItem_isUserTristate uintptr
	addr_QStandardItem_setUserTristate uintptr
	addr_QStandardItem_isDragEnabled uintptr
	addr_QStandardItem_setDragEnabled uintptr
	addr_QStandardItem_isDropEnabled uintptr
	addr_QStandardItem_setDropEnabled uintptr
	addr_QStandardItem_parent uintptr
	addr_QStandardItem_row uintptr
	addr_QStandardItem_column uintptr
	addr_QStandardItem_index uintptr
	addr_QStandardItem_model uintptr
	addr_QStandardItem_rowCount uintptr
	addr_QStandardItem_setRowCount uintptr
	addr_QStandardItem_columnCount uintptr
	addr_QStandardItem_setColumnCount uintptr
	addr_QStandardItem_hasChildren uintptr
	addr_QStandardItem_child uintptr
	addr_QStandardItem_setChild uintptr
	addr_QStandardItem_setChild2 uintptr
	addr_QStandardItem_insertRow uintptr
	addr_QStandardItem_insertColumn uintptr
	addr_QStandardItem_insertRows uintptr
	addr_QStandardItem_insertRows2 uintptr
	addr_QStandardItem_insertColumns uintptr
	addr_QStandardItem_removeRow uintptr
	addr_QStandardItem_removeColumn uintptr
	addr_QStandardItem_removeRows uintptr
	addr_QStandardItem_removeColumns uintptr
	addr_QStandardItem_appendRow uintptr
	addr_QStandardItem_appendRows uintptr
	addr_QStandardItem_appendColumn uintptr
	addr_QStandardItem_insertRow2 uintptr
	addr_QStandardItem_appendRow2 uintptr
	addr_QStandardItem_takeChild uintptr
	addr_QStandardItem_takeRow uintptr
	addr_QStandardItem_takeColumn uintptr
	addr_QStandardItem_sortChildren uintptr
	addr_QStandardItem_clone uintptr
	addr_QStandardItem_type uintptr
	addr_QStandardItem_read uintptr
	addr_QStandardItem_write uintptr
	addr_QStandardItem_OperatorLesser uintptr
	addr_QStandardItem_emitDataChanged uintptr
	addr_QStandardItem_child2 uintptr
	addr_QStandardItem_takeChild2 uintptr
	addr_QStandardItem_sortChildren2 uintptr
	addr_QStandardItem_data_default uintptr
	addr_QStandardItem_multiData_default uintptr
	addr_QStandardItem_setData_default uintptr
	addr_QStandardItem_clone_default uintptr
	addr_QStandardItem_type_default uintptr
	addr_QStandardItem_read_default uintptr
	addr_QStandardItem_write_default uintptr
	pg_QStandardItem__set_override func(uintptr, int32)
	pg_QStandardItem_delete func(uintptr)
)

func init() {
	symaddr(&addr_QStandardItem_new, "QStandardItem_new")
	symaddr(&addr_QStandardItem_new2, "QStandardItem_new2")
	symaddr(&addr_QStandardItem_new3, "QStandardItem_new3")
	symaddr(&addr_QStandardItem_new4, "QStandardItem_new4")
	symaddr(&addr_QStandardItem_new5, "QStandardItem_new5")
	symaddr(&addr_QStandardItem_data, "QStandardItem_data")
	symaddr(&addr_QStandardItem_multiData, "QStandardItem_multiData")
	symaddr(&addr_QStandardItem_setData, "QStandardItem_setData")
	symaddr(&addr_QStandardItem_clearData, "QStandardItem_clearData")
	symaddr(&addr_QStandardItem_text, "QStandardItem_text")
	symaddr(&addr_QStandardItem_setText, "QStandardItem_setText")
	symaddr(&addr_QStandardItem_icon, "QStandardItem_icon")
	symaddr(&addr_QStandardItem_setIcon, "QStandardItem_setIcon")
	symaddr(&addr_QStandardItem_toolTip, "QStandardItem_toolTip")
	symaddr(&addr_QStandardItem_setToolTip, "QStandardItem_setToolTip")
	symaddr(&addr_QStandardItem_statusTip, "QStandardItem_statusTip")
	symaddr(&addr_QStandardItem_setStatusTip, "QStandardItem_setStatusTip")
	symaddr(&addr_QStandardItem_whatsThis, "QStandardItem_whatsThis")
	symaddr(&addr_QStandardItem_setWhatsThis, "QStandardItem_setWhatsThis")
	symaddr(&addr_QStandardItem_sizeHint, "QStandardItem_sizeHint")
	symaddr(&addr_QStandardItem_setSizeHint, "QStandardItem_setSizeHint")
	symaddr(&addr_QStandardItem_font, "QStandardItem_font")
	symaddr(&addr_QStandardItem_setFont, "QStandardItem_setFont")
	symaddr(&addr_QStandardItem_textAlignment, "QStandardItem_textAlignment")
	symaddr(&addr_QStandardItem_setTextAlignment, "QStandardItem_setTextAlignment")
	symaddr(&addr_QStandardItem_background, "QStandardItem_background")
	symaddr(&addr_QStandardItem_setBackground, "QStandardItem_setBackground")
	symaddr(&addr_QStandardItem_foreground, "QStandardItem_foreground")
	symaddr(&addr_QStandardItem_setForeground, "QStandardItem_setForeground")
	symaddr(&addr_QStandardItem_checkState, "QStandardItem_checkState")
	symaddr(&addr_QStandardItem_setCheckState, "QStandardItem_setCheckState")
	symaddr(&addr_QStandardItem_accessibleText, "QStandardItem_accessibleText")
	symaddr(&addr_QStandardItem_setAccessibleText, "QStandardItem_setAccessibleText")
	symaddr(&addr_QStandardItem_accessibleDescription, "QStandardItem_accessibleDescription")
	symaddr(&addr_QStandardItem_setAccessibleDescription, "QStandardItem_setAccessibleDescription")
	symaddr(&addr_QStandardItem_flags, "QStandardItem_flags")
	symaddr(&addr_QStandardItem_setFlags, "QStandardItem_setFlags")
	symaddr(&addr_QStandardItem_isEnabled, "QStandardItem_isEnabled")
	symaddr(&addr_QStandardItem_setEnabled, "QStandardItem_setEnabled")
	symaddr(&addr_QStandardItem_isEditable, "QStandardItem_isEditable")
	symaddr(&addr_QStandardItem_setEditable, "QStandardItem_setEditable")
	symaddr(&addr_QStandardItem_isSelectable, "QStandardItem_isSelectable")
	symaddr(&addr_QStandardItem_setSelectable, "QStandardItem_setSelectable")
	symaddr(&addr_QStandardItem_isCheckable, "QStandardItem_isCheckable")
	symaddr(&addr_QStandardItem_setCheckable, "QStandardItem_setCheckable")
	symaddr(&addr_QStandardItem_isAutoTristate, "QStandardItem_isAutoTristate")
	symaddr(&addr_QStandardItem_setAutoTristate, "QStandardItem_setAutoTristate")
	symaddr(&addr_QStandardItem_isUserTristate, "QStandardItem_isUserTristate")
	symaddr(&addr_QStandardItem_setUserTristate, "QStandardItem_setUserTristate")
	symaddr(&addr_QStandardItem_isDragEnabled, "QStandardItem_isDragEnabled")
	symaddr(&addr_QStandardItem_setDragEnabled, "QStandardItem_setDragEnabled")
	symaddr(&addr_QStandardItem_isDropEnabled, "QStandardItem_isDropEnabled")
	symaddr(&addr_QStandardItem_setDropEnabled, "QStandardItem_setDropEnabled")
	symaddr(&addr_QStandardItem_parent, "QStandardItem_parent")
	symaddr(&addr_QStandardItem_row, "QStandardItem_row")
	symaddr(&addr_QStandardItem_column, "QStandardItem_column")
	symaddr(&addr_QStandardItem_index, "QStandardItem_index")
	symaddr(&addr_QStandardItem_model, "QStandardItem_model")
	symaddr(&addr_QStandardItem_rowCount, "QStandardItem_rowCount")
	symaddr(&addr_QStandardItem_setRowCount, "QStandardItem_setRowCount")
	symaddr(&addr_QStandardItem_columnCount, "QStandardItem_columnCount")
	symaddr(&addr_QStandardItem_setColumnCount, "QStandardItem_setColumnCount")
	symaddr(&addr_QStandardItem_hasChildren, "QStandardItem_hasChildren")
	symaddr(&addr_QStandardItem_child, "QStandardItem_child")
	symaddr(&addr_QStandardItem_setChild, "QStandardItem_setChild")
	symaddr(&addr_QStandardItem_setChild2, "QStandardItem_setChild2")
	symaddr(&addr_QStandardItem_insertRow, "QStandardItem_insertRow")
	symaddr(&addr_QStandardItem_insertColumn, "QStandardItem_insertColumn")
	symaddr(&addr_QStandardItem_insertRows, "QStandardItem_insertRows")
	symaddr(&addr_QStandardItem_insertRows2, "QStandardItem_insertRows2")
	symaddr(&addr_QStandardItem_insertColumns, "QStandardItem_insertColumns")
	symaddr(&addr_QStandardItem_removeRow, "QStandardItem_removeRow")
	symaddr(&addr_QStandardItem_removeColumn, "QStandardItem_removeColumn")
	symaddr(&addr_QStandardItem_removeRows, "QStandardItem_removeRows")
	symaddr(&addr_QStandardItem_removeColumns, "QStandardItem_removeColumns")
	symaddr(&addr_QStandardItem_appendRow, "QStandardItem_appendRow")
	symaddr(&addr_QStandardItem_appendRows, "QStandardItem_appendRows")
	symaddr(&addr_QStandardItem_appendColumn, "QStandardItem_appendColumn")
	symaddr(&addr_QStandardItem_insertRow2, "QStandardItem_insertRow2")
	symaddr(&addr_QStandardItem_appendRow2, "QStandardItem_appendRow2")
	symaddr(&addr_QStandardItem_takeChild, "QStandardItem_takeChild")
	symaddr(&addr_QStandardItem_takeRow, "QStandardItem_takeRow")
	symaddr(&addr_QStandardItem_takeColumn, "QStandardItem_takeColumn")
	symaddr(&addr_QStandardItem_sortChildren, "QStandardItem_sortChildren")
	symaddr(&addr_QStandardItem_clone, "QStandardItem_clone")
	symaddr(&addr_QStandardItem_type, "QStandardItem_type")
	symaddr(&addr_QStandardItem_read, "QStandardItem_read")
	symaddr(&addr_QStandardItem_write, "QStandardItem_write")
	symaddr(&addr_QStandardItem_OperatorLesser, "QStandardItem_OperatorLesser")
	symaddr(&addr_QStandardItem_emitDataChanged, "QStandardItem_emitDataChanged")
	symaddr(&addr_QStandardItem_child2, "QStandardItem_child2")
	symaddr(&addr_QStandardItem_takeChild2, "QStandardItem_takeChild2")
	symaddr(&addr_QStandardItem_sortChildren2, "QStandardItem_sortChildren2")
	symaddr(&addr_QStandardItem_data_default, "QStandardItem_data_default")
	symaddr(&addr_QStandardItem_multiData_default, "QStandardItem_multiData_default")
	symaddr(&addr_QStandardItem_setData_default, "QStandardItem_setData_default")
	symaddr(&addr_QStandardItem_clone_default, "QStandardItem_clone_default")
	symaddr(&addr_QStandardItem_type_default, "QStandardItem_type_default")
	symaddr(&addr_QStandardItem_read_default, "QStandardItem_read_default")
	symaddr(&addr_QStandardItem_write_default, "QStandardItem_write_default")
	reg(&pg_QStandardItem__set_override, "QStandardItem__set_override")
	reg(&pg_QStandardItem_delete, "QStandardItem_delete")
}

type QStandardItem struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QStandardItem) UnsafeGoID() int64     { return this.vid }
func (this *QStandardItem) UnsafeSetGoID(id int64) { this.vid = id }

func bindQStandardItem(h uintptr) *QStandardItem {
	if h == 0 {
		return nil
	}
	return &QStandardItem{h: h}
}

func newownedQStandardItem(h uintptr) *QStandardItem {
	o := bindQStandardItem(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QStandardItem).Delete)
	}
	return o
}

func (this *QStandardItem) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QStandardItem) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQStandardItem(h unsafe.Pointer) *QStandardItem { return bindQStandardItem(uintptr(h)) }
func UnsafeNewQStandardItemOwned(h unsafe.Pointer) *QStandardItem { return newownedQStandardItem(uintptr(h)) }

func (this *QStandardItem) UnsafeRelease() { this.h = 0 }

func NewQStandardItem() *QStandardItem {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QStandardItem_new, uintptr(__goid))
	__o := newownedQStandardItem(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QStandardItem_virtIdx, QStandardItem_enable)
	return __o
}

func NewQStandardItem2(text string) *QStandardItem {
	__goid := nextVirtID()
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QStandardItem_new2, uintptr(__goid), text_p, uintptr(len(text)))
	__o := newownedQStandardItem(__r)
	runtime.KeepAlive(text)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QStandardItem_virtIdx, QStandardItem_enable)
	return __o
}

func NewQStandardItem3(icon *QIcon, text string) *QStandardItem {
	__goid := nextVirtID()
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QStandardItem_new3, uintptr(__goid), uintptr(icon.UnsafePointer()), text_p, uintptr(len(text)))
	__o := newownedQStandardItem(__r)
	runtime.KeepAlive(icon)
	runtime.KeepAlive(text)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QStandardItem_virtIdx, QStandardItem_enable)
	return __o
}

func NewQStandardItem4(rows int) *QStandardItem {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QStandardItem_new4, uintptr(__goid), uintptr(rows))
	__o := newownedQStandardItem(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QStandardItem_virtIdx, QStandardItem_enable)
	return __o
}

func NewQStandardItem5(rows int, columns int) *QStandardItem {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QStandardItem_new5, uintptr(__goid), uintptr(rows), uintptr(columns))
	__o := newownedQStandardItem(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QStandardItem_virtIdx, QStandardItem_enable)
	return __o
}

func (this *QStandardItem) Data(role int) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QStandardItem_data, this.h, uintptr(role))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QStandardItem) MultiData(roleDataSpan *QModelRoleDataSpan) {
	purego.SyscallN(addr_QStandardItem_multiData, this.h, uintptr(roleDataSpan.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(roleDataSpan)
}

func (this *QStandardItem) SetData(value *QVariant, role int) {
	purego.SyscallN(addr_QStandardItem_setData, this.h, uintptr(value.UnsafePointer()), uintptr(role))
	runtime.KeepAlive(this)
	runtime.KeepAlive(value)
}

func (this *QStandardItem) ClearData() {
	purego.SyscallN(addr_QStandardItem_clearData, this.h)
	runtime.KeepAlive(this)
}

func (this *QStandardItem) Text() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStandardItem_text, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QStandardItem) SetText(text string) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QStandardItem_setText, this.h, text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
}

func (this *QStandardItem) Icon() *QIcon {
	__r, _, _ := purego.SyscallN(addr_QStandardItem_icon, this.h)
	runtime.KeepAlive(this)
	return newownedQIcon(__r)
}

func (this *QStandardItem) SetIcon(icon *QIcon) {
	purego.SyscallN(addr_QStandardItem_setIcon, this.h, uintptr(icon.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(icon)
}

func (this *QStandardItem) ToolTip() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStandardItem_toolTip, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QStandardItem) SetToolTip(toolTip string) {
	var toolTip_p uintptr
	if len(toolTip) > 0 {
		toolTip_p = uintptr(unsafe.Pointer(unsafe.StringData(toolTip)))
	}
	purego.SyscallN(addr_QStandardItem_setToolTip, this.h, toolTip_p, uintptr(len(toolTip)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(toolTip)
}

func (this *QStandardItem) StatusTip() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStandardItem_statusTip, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QStandardItem) SetStatusTip(statusTip string) {
	var statusTip_p uintptr
	if len(statusTip) > 0 {
		statusTip_p = uintptr(unsafe.Pointer(unsafe.StringData(statusTip)))
	}
	purego.SyscallN(addr_QStandardItem_setStatusTip, this.h, statusTip_p, uintptr(len(statusTip)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(statusTip)
}

func (this *QStandardItem) WhatsThis() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStandardItem_whatsThis, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QStandardItem) SetWhatsThis(whatsThis string) {
	var whatsThis_p uintptr
	if len(whatsThis) > 0 {
		whatsThis_p = uintptr(unsafe.Pointer(unsafe.StringData(whatsThis)))
	}
	purego.SyscallN(addr_QStandardItem_setWhatsThis, this.h, whatsThis_p, uintptr(len(whatsThis)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(whatsThis)
}

func (this *QStandardItem) SizeHint() *QSize {
	__r, _, _ := purego.SyscallN(addr_QStandardItem_sizeHint, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QStandardItem) SetSizeHint(sizeHint *QSize) {
	purego.SyscallN(addr_QStandardItem_setSizeHint, this.h, uintptr(sizeHint.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sizeHint)
}

func (this *QStandardItem) Font() *QFont {
	__r, _, _ := purego.SyscallN(addr_QStandardItem_font, this.h)
	runtime.KeepAlive(this)
	return newownedQFont(__r)
}

func (this *QStandardItem) SetFont(font *QFont) {
	purego.SyscallN(addr_QStandardItem_setFont, this.h, uintptr(font.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(font)
}

func (this *QStandardItem) TextAlignment() AlignmentFlag {
	__r0, _, _ := purego.SyscallN(addr_QStandardItem_textAlignment, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return AlignmentFlag(__r)
}

func (this *QStandardItem) SetTextAlignment(textAlignment AlignmentFlag) {
	purego.SyscallN(addr_QStandardItem_setTextAlignment, this.h, uintptr(textAlignment))
	runtime.KeepAlive(this)
}

func (this *QStandardItem) Background() *QBrush {
	__r, _, _ := purego.SyscallN(addr_QStandardItem_background, this.h)
	runtime.KeepAlive(this)
	return newownedQBrush(__r)
}

func (this *QStandardItem) SetBackground(brush *QBrush) {
	purego.SyscallN(addr_QStandardItem_setBackground, this.h, uintptr(brush.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(brush)
}

func (this *QStandardItem) Foreground() *QBrush {
	__r, _, _ := purego.SyscallN(addr_QStandardItem_foreground, this.h)
	runtime.KeepAlive(this)
	return newownedQBrush(__r)
}

func (this *QStandardItem) SetForeground(brush *QBrush) {
	purego.SyscallN(addr_QStandardItem_setForeground, this.h, uintptr(brush.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(brush)
}

func (this *QStandardItem) CheckState() CheckState {
	__r0, _, _ := purego.SyscallN(addr_QStandardItem_checkState, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return CheckState(__r)
}

func (this *QStandardItem) SetCheckState(checkState CheckState) {
	purego.SyscallN(addr_QStandardItem_setCheckState, this.h, uintptr(checkState))
	runtime.KeepAlive(this)
}

func (this *QStandardItem) AccessibleText() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStandardItem_accessibleText, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QStandardItem) SetAccessibleText(accessibleText string) {
	var accessibleText_p uintptr
	if len(accessibleText) > 0 {
		accessibleText_p = uintptr(unsafe.Pointer(unsafe.StringData(accessibleText)))
	}
	purego.SyscallN(addr_QStandardItem_setAccessibleText, this.h, accessibleText_p, uintptr(len(accessibleText)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(accessibleText)
}

func (this *QStandardItem) AccessibleDescription() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStandardItem_accessibleDescription, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QStandardItem) SetAccessibleDescription(accessibleDescription string) {
	var accessibleDescription_p uintptr
	if len(accessibleDescription) > 0 {
		accessibleDescription_p = uintptr(unsafe.Pointer(unsafe.StringData(accessibleDescription)))
	}
	purego.SyscallN(addr_QStandardItem_setAccessibleDescription, this.h, accessibleDescription_p, uintptr(len(accessibleDescription)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(accessibleDescription)
}

func (this *QStandardItem) Flags() ItemFlag {
	__r0, _, _ := purego.SyscallN(addr_QStandardItem_flags, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return ItemFlag(__r)
}

func (this *QStandardItem) SetFlags(flags ItemFlag) {
	purego.SyscallN(addr_QStandardItem_setFlags, this.h, uintptr(flags))
	runtime.KeepAlive(this)
}

func (this *QStandardItem) IsEnabled() bool {
	__r0, _, _ := purego.SyscallN(addr_QStandardItem_isEnabled, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QStandardItem) SetEnabled(enabled bool) {
	purego.SyscallN(addr_QStandardItem_setEnabled, this.h, uintptr(b2i(enabled)))
	runtime.KeepAlive(this)
}

func (this *QStandardItem) IsEditable() bool {
	__r0, _, _ := purego.SyscallN(addr_QStandardItem_isEditable, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QStandardItem) SetEditable(editable bool) {
	purego.SyscallN(addr_QStandardItem_setEditable, this.h, uintptr(b2i(editable)))
	runtime.KeepAlive(this)
}

func (this *QStandardItem) IsSelectable() bool {
	__r0, _, _ := purego.SyscallN(addr_QStandardItem_isSelectable, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QStandardItem) SetSelectable(selectable bool) {
	purego.SyscallN(addr_QStandardItem_setSelectable, this.h, uintptr(b2i(selectable)))
	runtime.KeepAlive(this)
}

func (this *QStandardItem) IsCheckable() bool {
	__r0, _, _ := purego.SyscallN(addr_QStandardItem_isCheckable, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QStandardItem) SetCheckable(checkable bool) {
	purego.SyscallN(addr_QStandardItem_setCheckable, this.h, uintptr(b2i(checkable)))
	runtime.KeepAlive(this)
}

func (this *QStandardItem) IsAutoTristate() bool {
	__r0, _, _ := purego.SyscallN(addr_QStandardItem_isAutoTristate, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QStandardItem) SetAutoTristate(tristate bool) {
	purego.SyscallN(addr_QStandardItem_setAutoTristate, this.h, uintptr(b2i(tristate)))
	runtime.KeepAlive(this)
}

func (this *QStandardItem) IsUserTristate() bool {
	__r0, _, _ := purego.SyscallN(addr_QStandardItem_isUserTristate, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QStandardItem) SetUserTristate(tristate bool) {
	purego.SyscallN(addr_QStandardItem_setUserTristate, this.h, uintptr(b2i(tristate)))
	runtime.KeepAlive(this)
}

func (this *QStandardItem) IsDragEnabled() bool {
	__r0, _, _ := purego.SyscallN(addr_QStandardItem_isDragEnabled, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QStandardItem) SetDragEnabled(dragEnabled bool) {
	purego.SyscallN(addr_QStandardItem_setDragEnabled, this.h, uintptr(b2i(dragEnabled)))
	runtime.KeepAlive(this)
}

func (this *QStandardItem) IsDropEnabled() bool {
	__r0, _, _ := purego.SyscallN(addr_QStandardItem_isDropEnabled, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QStandardItem) SetDropEnabled(dropEnabled bool) {
	purego.SyscallN(addr_QStandardItem_setDropEnabled, this.h, uintptr(b2i(dropEnabled)))
	runtime.KeepAlive(this)
}

func (this *QStandardItem) Parent() *QStandardItem {
	__r, _, _ := purego.SyscallN(addr_QStandardItem_parent, this.h)
	runtime.KeepAlive(this)
	return bindQStandardItem(__r)
}

func (this *QStandardItem) Row() int {
	__r0, _, _ := purego.SyscallN(addr_QStandardItem_row, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QStandardItem) Column() int {
	__r0, _, _ := purego.SyscallN(addr_QStandardItem_column, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QStandardItem) Index() *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QStandardItem_index, this.h)
	runtime.KeepAlive(this)
	return newownedQModelIndex(__r)
}

func (this *QStandardItem) Model() *QStandardItemModel {
	__r, _, _ := purego.SyscallN(addr_QStandardItem_model, this.h)
	runtime.KeepAlive(this)
	return bindQStandardItemModel(__r)
}

func (this *QStandardItem) RowCount() int {
	__r0, _, _ := purego.SyscallN(addr_QStandardItem_rowCount, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QStandardItem) SetRowCount(rows int) {
	purego.SyscallN(addr_QStandardItem_setRowCount, this.h, uintptr(rows))
	runtime.KeepAlive(this)
}

func (this *QStandardItem) ColumnCount() int {
	__r0, _, _ := purego.SyscallN(addr_QStandardItem_columnCount, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QStandardItem) SetColumnCount(columns int) {
	purego.SyscallN(addr_QStandardItem_setColumnCount, this.h, uintptr(columns))
	runtime.KeepAlive(this)
}

func (this *QStandardItem) HasChildren() bool {
	__r0, _, _ := purego.SyscallN(addr_QStandardItem_hasChildren, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QStandardItem) Child(row int) *QStandardItem {
	__r, _, _ := purego.SyscallN(addr_QStandardItem_child, this.h, uintptr(row))
	runtime.KeepAlive(this)
	return bindQStandardItem(__r)
}

func (this *QStandardItem) SetChild(row int, column int, item *QStandardItem) {
	purego.SyscallN(addr_QStandardItem_setChild, this.h, uintptr(row), uintptr(column), uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QStandardItem) SetChild2(row int, item *QStandardItem) {
	purego.SyscallN(addr_QStandardItem_setChild2, this.h, uintptr(row), uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QStandardItem) InsertRow(row int, items []*QStandardItem) {
	items_arr := make([]uintptr, len(items))
	for __i, __e := range items {
		items_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var items_p uintptr
	if len(items_arr) > 0 {
		items_p = uintptr(unsafe.Pointer(&items_arr[0]))
	}
	purego.SyscallN(addr_QStandardItem_insertRow, this.h, uintptr(row), items_p, uintptr(len(items)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(items)
	runtime.KeepAlive(items_arr)
}

func (this *QStandardItem) InsertColumn(column int, items []*QStandardItem) {
	items_arr := make([]uintptr, len(items))
	for __i, __e := range items {
		items_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var items_p uintptr
	if len(items_arr) > 0 {
		items_p = uintptr(unsafe.Pointer(&items_arr[0]))
	}
	purego.SyscallN(addr_QStandardItem_insertColumn, this.h, uintptr(column), items_p, uintptr(len(items)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(items)
	runtime.KeepAlive(items_arr)
}

func (this *QStandardItem) InsertRows(row int, items []*QStandardItem) {
	items_arr := make([]uintptr, len(items))
	for __i, __e := range items {
		items_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var items_p uintptr
	if len(items_arr) > 0 {
		items_p = uintptr(unsafe.Pointer(&items_arr[0]))
	}
	purego.SyscallN(addr_QStandardItem_insertRows, this.h, uintptr(row), items_p, uintptr(len(items)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(items)
	runtime.KeepAlive(items_arr)
}

func (this *QStandardItem) InsertRows2(row int, count int) {
	purego.SyscallN(addr_QStandardItem_insertRows2, this.h, uintptr(row), uintptr(count))
	runtime.KeepAlive(this)
}

func (this *QStandardItem) InsertColumns(column int, count int) {
	purego.SyscallN(addr_QStandardItem_insertColumns, this.h, uintptr(column), uintptr(count))
	runtime.KeepAlive(this)
}

func (this *QStandardItem) RemoveRow(row int) {
	purego.SyscallN(addr_QStandardItem_removeRow, this.h, uintptr(row))
	runtime.KeepAlive(this)
}

func (this *QStandardItem) RemoveColumn(column int) {
	purego.SyscallN(addr_QStandardItem_removeColumn, this.h, uintptr(column))
	runtime.KeepAlive(this)
}

func (this *QStandardItem) RemoveRows(row int, count int) {
	purego.SyscallN(addr_QStandardItem_removeRows, this.h, uintptr(row), uintptr(count))
	runtime.KeepAlive(this)
}

func (this *QStandardItem) RemoveColumns(column int, count int) {
	purego.SyscallN(addr_QStandardItem_removeColumns, this.h, uintptr(column), uintptr(count))
	runtime.KeepAlive(this)
}

func (this *QStandardItem) AppendRow(items []*QStandardItem) {
	items_arr := make([]uintptr, len(items))
	for __i, __e := range items {
		items_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var items_p uintptr
	if len(items_arr) > 0 {
		items_p = uintptr(unsafe.Pointer(&items_arr[0]))
	}
	purego.SyscallN(addr_QStandardItem_appendRow, this.h, items_p, uintptr(len(items)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(items)
	runtime.KeepAlive(items_arr)
}

func (this *QStandardItem) AppendRows(items []*QStandardItem) {
	items_arr := make([]uintptr, len(items))
	for __i, __e := range items {
		items_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var items_p uintptr
	if len(items_arr) > 0 {
		items_p = uintptr(unsafe.Pointer(&items_arr[0]))
	}
	purego.SyscallN(addr_QStandardItem_appendRows, this.h, items_p, uintptr(len(items)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(items)
	runtime.KeepAlive(items_arr)
}

func (this *QStandardItem) AppendColumn(items []*QStandardItem) {
	items_arr := make([]uintptr, len(items))
	for __i, __e := range items {
		items_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var items_p uintptr
	if len(items_arr) > 0 {
		items_p = uintptr(unsafe.Pointer(&items_arr[0]))
	}
	purego.SyscallN(addr_QStandardItem_appendColumn, this.h, items_p, uintptr(len(items)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(items)
	runtime.KeepAlive(items_arr)
}

func (this *QStandardItem) InsertRow2(row int, item *QStandardItem) {
	purego.SyscallN(addr_QStandardItem_insertRow2, this.h, uintptr(row), uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QStandardItem) AppendRowWithItem(item *QStandardItem) {
	purego.SyscallN(addr_QStandardItem_appendRow2, this.h, uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QStandardItem) TakeChild(row int) *QStandardItem {
	__r, _, _ := purego.SyscallN(addr_QStandardItem_takeChild, this.h, uintptr(row))
	runtime.KeepAlive(this)
	return bindQStandardItem(__r)
}

func (this *QStandardItem) TakeRow(row int) []*QStandardItem {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStandardItem_takeRow, this.h, uintptr(row), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QStandardItem, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = bindQStandardItem(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QStandardItem) TakeColumn(column int) []*QStandardItem {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStandardItem_takeColumn, this.h, uintptr(column), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QStandardItem, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = bindQStandardItem(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QStandardItem) SortChildren(column int) {
	purego.SyscallN(addr_QStandardItem_sortChildren, this.h, uintptr(column))
	runtime.KeepAlive(this)
}

func (this *QStandardItem) Clone() *QStandardItem {
	__r, _, _ := purego.SyscallN(addr_QStandardItem_clone, this.h)
	runtime.KeepAlive(this)
	return bindQStandardItem(__r)
}

func (this *QStandardItem) Type() int {
	__r0, _, _ := purego.SyscallN(addr_QStandardItem_type, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QStandardItem) Read(in *QDataStream) {
	purego.SyscallN(addr_QStandardItem_read, this.h, uintptr(in.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(in)
}

func (this *QStandardItem) Write(out *QDataStream) {
	purego.SyscallN(addr_QStandardItem_write, this.h, uintptr(out.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(out)
}

func (this *QStandardItem) OperatorLesser(other *QStandardItem) bool {
	__r0, _, _ := purego.SyscallN(addr_QStandardItem_OperatorLesser, this.h, uintptr(other.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return __r != 0
}

func (this *QStandardItem) EmitDataChanged() {
	purego.SyscallN(addr_QStandardItem_emitDataChanged, this.h)
	runtime.KeepAlive(this)
}

func (this *QStandardItem) Child2(row int, column int) *QStandardItem {
	__r, _, _ := purego.SyscallN(addr_QStandardItem_child2, this.h, uintptr(row), uintptr(column))
	runtime.KeepAlive(this)
	return bindQStandardItem(__r)
}

func (this *QStandardItem) TakeChild2(row int, column int) *QStandardItem {
	__r, _, _ := purego.SyscallN(addr_QStandardItem_takeChild2, this.h, uintptr(row), uintptr(column))
	runtime.KeepAlive(this)
	return bindQStandardItem(__r)
}

func (this *QStandardItem) SortChildren2(column int, order SortOrder) {
	purego.SyscallN(addr_QStandardItem_sortChildren2, this.h, uintptr(column), uintptr(order))
	runtime.KeepAlive(this)
}

func (this *QStandardItem) DataDefault(role int) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QStandardItem_data_default, this.h, uintptr(role))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QStandardItem) MultiDataDefault(roleDataSpan *QModelRoleDataSpan) {
	purego.SyscallN(addr_QStandardItem_multiData_default, this.h, uintptr(roleDataSpan.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(roleDataSpan)
}

func (this *QStandardItem) SetDataDefault(value *QVariant, role int) {
	purego.SyscallN(addr_QStandardItem_setData_default, this.h, uintptr(value.UnsafePointer()), uintptr(role))
	runtime.KeepAlive(this)
	runtime.KeepAlive(value)
}

func (this *QStandardItem) CloneDefault() *QStandardItem {
	__r, _, _ := purego.SyscallN(addr_QStandardItem_clone_default, this.h)
	runtime.KeepAlive(this)
	return bindQStandardItem(__r)
}

func (this *QStandardItem) TypeDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QStandardItem_type_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QStandardItem) ReadDefault(in *QDataStream) {
	purego.SyscallN(addr_QStandardItem_read_default, this.h, uintptr(in.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(in)
}

func (this *QStandardItem) WriteDefault(out *QDataStream) {
	purego.SyscallN(addr_QStandardItem_write_default, this.h, uintptr(out.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(out)
}

func QStandardItem_virtIdx(slot int64) int32 {
	switch slot {
	case -8837375942098466299:
		return 0
	case 731311060944096642:
		return 1
	case 35191454575202219:
		return 2
	case -7032813281048616478:
		return 3
	case -6371403872082420691:
		return 4
	case 5541207787354052101:
		return 5
	case -5099743077423337732:
		return 6
	}
	return -1
}

func QStandardItem_enable(h uintptr, idx int32) { pg_QStandardItem__set_override(h, idx) }

func (this *QStandardItem) OnData(fn func(int) *QVariant) {
	registerVirt(this.UnsafeGoID(), -8837375942098466299, func(role int32) uintptr { return uintptr(fn(int(role)).UnsafePointer()) }) // slot "data"
}

func (this *QStandardItem) OnMultiData(fn func(*QModelRoleDataSpan)) {
	registerVirt(this.UnsafeGoID(), 731311060944096642, func(roleDataSpan uintptr) { fn(UnsafeNewQModelRoleDataSpan(unsafe.Pointer(roleDataSpan))) }) // slot "multiData"
}

func (this *QStandardItem) OnSetData(fn func(*QVariant, int)) {
	registerVirt(this.UnsafeGoID(), 35191454575202219, func(value uintptr, role int32) { fn(UnsafeNewQVariant(unsafe.Pointer(value)), int(role)) }) // slot "setData"
}

func (this *QStandardItem) OnClone(fn func() *QStandardItem) {
	registerVirt(this.UnsafeGoID(), -7032813281048616478, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "clone"
}

func (this *QStandardItem) OnType(fn func() int) {
	registerVirt(this.UnsafeGoID(), -6371403872082420691, func() int32 { return int32(fn()) }) // slot "type"
}

func (this *QStandardItem) OnRead(fn func(*QDataStream)) {
	registerVirt(this.UnsafeGoID(), 5541207787354052101, func(in uintptr) { fn(UnsafeNewQDataStream(unsafe.Pointer(in))) }) // slot "read"
}

func (this *QStandardItem) OnWrite(fn func(*QDataStream)) {
	registerVirt(this.UnsafeGoID(), -5099743077423337732, func(out uintptr) { fn(UnsafeNewQDataStream(unsafe.Pointer(out))) }) // slot "write"
}

func (this *QStandardItem) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QStandardItem_delete(this.h)
	}
	this.UnsafeRelease()
}
