package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QTableWidgetItem_new uintptr
	addr_QTableWidgetItem_new2 uintptr
	addr_QTableWidgetItem_new3 uintptr
	addr_QTableWidgetItem_new4 uintptr
	addr_QTableWidgetItem_new5 uintptr
	addr_QTableWidgetItem_new6 uintptr
	addr_QTableWidgetItem_new7 uintptr
	addr_QTableWidgetItem_clone uintptr
	addr_QTableWidgetItem_tableWidget uintptr
	addr_QTableWidgetItem_row uintptr
	addr_QTableWidgetItem_column uintptr
	addr_QTableWidgetItem_setSelected uintptr
	addr_QTableWidgetItem_isSelected uintptr
	addr_QTableWidgetItem_flags uintptr
	addr_QTableWidgetItem_setFlags uintptr
	addr_QTableWidgetItem_text uintptr
	addr_QTableWidgetItem_setText uintptr
	addr_QTableWidgetItem_icon uintptr
	addr_QTableWidgetItem_setIcon uintptr
	addr_QTableWidgetItem_statusTip uintptr
	addr_QTableWidgetItem_setStatusTip uintptr
	addr_QTableWidgetItem_toolTip uintptr
	addr_QTableWidgetItem_setToolTip uintptr
	addr_QTableWidgetItem_whatsThis uintptr
	addr_QTableWidgetItem_setWhatsThis uintptr
	addr_QTableWidgetItem_font uintptr
	addr_QTableWidgetItem_setFont uintptr
	addr_QTableWidgetItem_textAlignment uintptr
	addr_QTableWidgetItem_setTextAlignment uintptr
	addr_QTableWidgetItem_setTextAlignment2 uintptr
	addr_QTableWidgetItem_setTextAlignment3 uintptr
	addr_QTableWidgetItem_background uintptr
	addr_QTableWidgetItem_setBackground uintptr
	addr_QTableWidgetItem_foreground uintptr
	addr_QTableWidgetItem_setForeground uintptr
	addr_QTableWidgetItem_checkState uintptr
	addr_QTableWidgetItem_setCheckState uintptr
	addr_QTableWidgetItem_sizeHint uintptr
	addr_QTableWidgetItem_setSizeHint uintptr
	addr_QTableWidgetItem_data uintptr
	addr_QTableWidgetItem_setData uintptr
	addr_QTableWidgetItem_OperatorLesser uintptr
	addr_QTableWidgetItem_read uintptr
	addr_QTableWidgetItem_write uintptr
	addr_QTableWidgetItem_OperatorAssign uintptr
	addr_QTableWidgetItem_type uintptr
	addr_QTableWidgetItem_clone_default uintptr
	addr_QTableWidgetItem_data_default uintptr
	addr_QTableWidgetItem_setData_default uintptr
	addr_QTableWidgetItem_read_default uintptr
	addr_QTableWidgetItem_write_default uintptr
	pg_QTableWidgetItem__set_override func(uintptr, int32)
	pg_QTableWidgetItem_delete func(uintptr)
)

func init() {
	symaddr(&addr_QTableWidgetItem_new, "QTableWidgetItem_new")
	symaddr(&addr_QTableWidgetItem_new2, "QTableWidgetItem_new2")
	symaddr(&addr_QTableWidgetItem_new3, "QTableWidgetItem_new3")
	symaddr(&addr_QTableWidgetItem_new4, "QTableWidgetItem_new4")
	symaddr(&addr_QTableWidgetItem_new5, "QTableWidgetItem_new5")
	symaddr(&addr_QTableWidgetItem_new6, "QTableWidgetItem_new6")
	symaddr(&addr_QTableWidgetItem_new7, "QTableWidgetItem_new7")
	symaddr(&addr_QTableWidgetItem_clone, "QTableWidgetItem_clone")
	symaddr(&addr_QTableWidgetItem_tableWidget, "QTableWidgetItem_tableWidget")
	symaddr(&addr_QTableWidgetItem_row, "QTableWidgetItem_row")
	symaddr(&addr_QTableWidgetItem_column, "QTableWidgetItem_column")
	symaddr(&addr_QTableWidgetItem_setSelected, "QTableWidgetItem_setSelected")
	symaddr(&addr_QTableWidgetItem_isSelected, "QTableWidgetItem_isSelected")
	symaddr(&addr_QTableWidgetItem_flags, "QTableWidgetItem_flags")
	symaddr(&addr_QTableWidgetItem_setFlags, "QTableWidgetItem_setFlags")
	symaddr(&addr_QTableWidgetItem_text, "QTableWidgetItem_text")
	symaddr(&addr_QTableWidgetItem_setText, "QTableWidgetItem_setText")
	symaddr(&addr_QTableWidgetItem_icon, "QTableWidgetItem_icon")
	symaddr(&addr_QTableWidgetItem_setIcon, "QTableWidgetItem_setIcon")
	symaddr(&addr_QTableWidgetItem_statusTip, "QTableWidgetItem_statusTip")
	symaddr(&addr_QTableWidgetItem_setStatusTip, "QTableWidgetItem_setStatusTip")
	symaddr(&addr_QTableWidgetItem_toolTip, "QTableWidgetItem_toolTip")
	symaddr(&addr_QTableWidgetItem_setToolTip, "QTableWidgetItem_setToolTip")
	symaddr(&addr_QTableWidgetItem_whatsThis, "QTableWidgetItem_whatsThis")
	symaddr(&addr_QTableWidgetItem_setWhatsThis, "QTableWidgetItem_setWhatsThis")
	symaddr(&addr_QTableWidgetItem_font, "QTableWidgetItem_font")
	symaddr(&addr_QTableWidgetItem_setFont, "QTableWidgetItem_setFont")
	symaddr(&addr_QTableWidgetItem_textAlignment, "QTableWidgetItem_textAlignment")
	symaddr(&addr_QTableWidgetItem_setTextAlignment, "QTableWidgetItem_setTextAlignment")
	symaddr(&addr_QTableWidgetItem_setTextAlignment2, "QTableWidgetItem_setTextAlignment2")
	symaddr(&addr_QTableWidgetItem_setTextAlignment3, "QTableWidgetItem_setTextAlignment3")
	symaddr(&addr_QTableWidgetItem_background, "QTableWidgetItem_background")
	symaddr(&addr_QTableWidgetItem_setBackground, "QTableWidgetItem_setBackground")
	symaddr(&addr_QTableWidgetItem_foreground, "QTableWidgetItem_foreground")
	symaddr(&addr_QTableWidgetItem_setForeground, "QTableWidgetItem_setForeground")
	symaddr(&addr_QTableWidgetItem_checkState, "QTableWidgetItem_checkState")
	symaddr(&addr_QTableWidgetItem_setCheckState, "QTableWidgetItem_setCheckState")
	symaddr(&addr_QTableWidgetItem_sizeHint, "QTableWidgetItem_sizeHint")
	symaddr(&addr_QTableWidgetItem_setSizeHint, "QTableWidgetItem_setSizeHint")
	symaddr(&addr_QTableWidgetItem_data, "QTableWidgetItem_data")
	symaddr(&addr_QTableWidgetItem_setData, "QTableWidgetItem_setData")
	symaddr(&addr_QTableWidgetItem_OperatorLesser, "QTableWidgetItem_OperatorLesser")
	symaddr(&addr_QTableWidgetItem_read, "QTableWidgetItem_read")
	symaddr(&addr_QTableWidgetItem_write, "QTableWidgetItem_write")
	symaddr(&addr_QTableWidgetItem_OperatorAssign, "QTableWidgetItem_OperatorAssign")
	symaddr(&addr_QTableWidgetItem_type, "QTableWidgetItem_type")
	symaddr(&addr_QTableWidgetItem_clone_default, "QTableWidgetItem_clone_default")
	symaddr(&addr_QTableWidgetItem_data_default, "QTableWidgetItem_data_default")
	symaddr(&addr_QTableWidgetItem_setData_default, "QTableWidgetItem_setData_default")
	symaddr(&addr_QTableWidgetItem_read_default, "QTableWidgetItem_read_default")
	symaddr(&addr_QTableWidgetItem_write_default, "QTableWidgetItem_write_default")
	reg(&pg_QTableWidgetItem__set_override, "QTableWidgetItem__set_override")
	reg(&pg_QTableWidgetItem_delete, "QTableWidgetItem_delete")
}

type QTableWidgetItem struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QTableWidgetItem) UnsafeGoID() int64     { return this.vid }
func (this *QTableWidgetItem) UnsafeSetGoID(id int64) { this.vid = id }

func bindQTableWidgetItem(h uintptr) *QTableWidgetItem {
	if h == 0 {
		return nil
	}
	return &QTableWidgetItem{h: h}
}

func newownedQTableWidgetItem(h uintptr) *QTableWidgetItem {
	o := bindQTableWidgetItem(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QTableWidgetItem).Delete)
	}
	return o
}

func (this *QTableWidgetItem) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QTableWidgetItem) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQTableWidgetItem(h unsafe.Pointer) *QTableWidgetItem { return bindQTableWidgetItem(uintptr(h)) }
func UnsafeNewQTableWidgetItemOwned(h unsafe.Pointer) *QTableWidgetItem { return newownedQTableWidgetItem(uintptr(h)) }

func (this *QTableWidgetItem) UnsafeRelease() { this.h = 0 }

func NewQTableWidgetItem() *QTableWidgetItem {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QTableWidgetItem_new, uintptr(__goid))
	__o := newownedQTableWidgetItem(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QTableWidgetItem_virtIdx, QTableWidgetItem_enable)
	return __o
}

func NewQTableWidgetItem2(text string) *QTableWidgetItem {
	__goid := nextVirtID()
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QTableWidgetItem_new2, uintptr(__goid), text_p, uintptr(len(text)))
	__o := newownedQTableWidgetItem(__r)
	runtime.KeepAlive(text)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QTableWidgetItem_virtIdx, QTableWidgetItem_enable)
	return __o
}

func NewQTableWidgetItem3(icon *QIcon, text string) *QTableWidgetItem {
	__goid := nextVirtID()
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QTableWidgetItem_new3, uintptr(__goid), uintptr(icon.UnsafePointer()), text_p, uintptr(len(text)))
	__o := newownedQTableWidgetItem(__r)
	runtime.KeepAlive(icon)
	runtime.KeepAlive(text)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QTableWidgetItem_virtIdx, QTableWidgetItem_enable)
	return __o
}

func NewQTableWidgetItem4(other *QTableWidgetItem) *QTableWidgetItem {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QTableWidgetItem_new4, uintptr(__goid), uintptr(other.UnsafePointer()))
	__o := newownedQTableWidgetItem(__r)
	runtime.KeepAlive(other)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QTableWidgetItem_virtIdx, QTableWidgetItem_enable)
	return __o
}

func NewQTableWidgetItem5(type_ int) *QTableWidgetItem {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QTableWidgetItem_new5, uintptr(__goid), uintptr(type_))
	__o := newownedQTableWidgetItem(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QTableWidgetItem_virtIdx, QTableWidgetItem_enable)
	return __o
}

func NewQTableWidgetItem6(text string, type_ int) *QTableWidgetItem {
	__goid := nextVirtID()
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QTableWidgetItem_new6, uintptr(__goid), text_p, uintptr(len(text)), uintptr(type_))
	__o := newownedQTableWidgetItem(__r)
	runtime.KeepAlive(text)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QTableWidgetItem_virtIdx, QTableWidgetItem_enable)
	return __o
}

func NewQTableWidgetItem7(icon *QIcon, text string, type_ int) *QTableWidgetItem {
	__goid := nextVirtID()
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QTableWidgetItem_new7, uintptr(__goid), uintptr(icon.UnsafePointer()), text_p, uintptr(len(text)), uintptr(type_))
	__o := newownedQTableWidgetItem(__r)
	runtime.KeepAlive(icon)
	runtime.KeepAlive(text)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QTableWidgetItem_virtIdx, QTableWidgetItem_enable)
	return __o
}

func (this *QTableWidgetItem) Clone() *QTableWidgetItem {
	__r, _, _ := purego.SyscallN(addr_QTableWidgetItem_clone, this.h)
	runtime.KeepAlive(this)
	return bindQTableWidgetItem(__r)
}

func (this *QTableWidgetItem) TableWidget() *QTableWidget {
	__r, _, _ := purego.SyscallN(addr_QTableWidgetItem_tableWidget, this.h)
	runtime.KeepAlive(this)
	return bindQTableWidget(__r)
}

func (this *QTableWidgetItem) Row() int {
	__r0, _, _ := purego.SyscallN(addr_QTableWidgetItem_row, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTableWidgetItem) Column() int {
	__r0, _, _ := purego.SyscallN(addr_QTableWidgetItem_column, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTableWidgetItem) SetSelected(select_ bool) {
	purego.SyscallN(addr_QTableWidgetItem_setSelected, this.h, uintptr(b2i(select_)))
	runtime.KeepAlive(this)
}

func (this *QTableWidgetItem) IsSelected() bool {
	__r0, _, _ := purego.SyscallN(addr_QTableWidgetItem_isSelected, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTableWidgetItem) Flags() ItemFlag {
	__r0, _, _ := purego.SyscallN(addr_QTableWidgetItem_flags, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return ItemFlag(__r)
}

func (this *QTableWidgetItem) SetFlags(flags ItemFlag) {
	purego.SyscallN(addr_QTableWidgetItem_setFlags, this.h, uintptr(flags))
	runtime.KeepAlive(this)
}

func (this *QTableWidgetItem) Text() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTableWidgetItem_text, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QTableWidgetItem) SetText(text string) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QTableWidgetItem_setText, this.h, text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
}

func (this *QTableWidgetItem) Icon() *QIcon {
	__r, _, _ := purego.SyscallN(addr_QTableWidgetItem_icon, this.h)
	runtime.KeepAlive(this)
	return newownedQIcon(__r)
}

func (this *QTableWidgetItem) SetIcon(icon *QIcon) {
	purego.SyscallN(addr_QTableWidgetItem_setIcon, this.h, uintptr(icon.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(icon)
}

func (this *QTableWidgetItem) StatusTip() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTableWidgetItem_statusTip, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QTableWidgetItem) SetStatusTip(statusTip string) {
	var statusTip_p uintptr
	if len(statusTip) > 0 {
		statusTip_p = uintptr(unsafe.Pointer(unsafe.StringData(statusTip)))
	}
	purego.SyscallN(addr_QTableWidgetItem_setStatusTip, this.h, statusTip_p, uintptr(len(statusTip)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(statusTip)
}

func (this *QTableWidgetItem) ToolTip() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTableWidgetItem_toolTip, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QTableWidgetItem) SetToolTip(toolTip string) {
	var toolTip_p uintptr
	if len(toolTip) > 0 {
		toolTip_p = uintptr(unsafe.Pointer(unsafe.StringData(toolTip)))
	}
	purego.SyscallN(addr_QTableWidgetItem_setToolTip, this.h, toolTip_p, uintptr(len(toolTip)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(toolTip)
}

func (this *QTableWidgetItem) WhatsThis() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTableWidgetItem_whatsThis, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QTableWidgetItem) SetWhatsThis(whatsThis string) {
	var whatsThis_p uintptr
	if len(whatsThis) > 0 {
		whatsThis_p = uintptr(unsafe.Pointer(unsafe.StringData(whatsThis)))
	}
	purego.SyscallN(addr_QTableWidgetItem_setWhatsThis, this.h, whatsThis_p, uintptr(len(whatsThis)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(whatsThis)
}

func (this *QTableWidgetItem) Font() *QFont {
	__r, _, _ := purego.SyscallN(addr_QTableWidgetItem_font, this.h)
	runtime.KeepAlive(this)
	return newownedQFont(__r)
}

func (this *QTableWidgetItem) SetFont(font *QFont) {
	purego.SyscallN(addr_QTableWidgetItem_setFont, this.h, uintptr(font.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(font)
}

func (this *QTableWidgetItem) TextAlignment() int {
	__r0, _, _ := purego.SyscallN(addr_QTableWidgetItem_textAlignment, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTableWidgetItem) SetTextAlignment(alignment int) {
	purego.SyscallN(addr_QTableWidgetItem_setTextAlignment, this.h, uintptr(alignment))
	runtime.KeepAlive(this)
}

func (this *QTableWidgetItem) SetTextAlignmentWithAlignment(alignment AlignmentFlag) {
	purego.SyscallN(addr_QTableWidgetItem_setTextAlignment2, this.h, uintptr(alignment))
	runtime.KeepAlive(this)
}

func (this *QTableWidgetItem) SetTextAlignment2(alignment AlignmentFlag) {
	purego.SyscallN(addr_QTableWidgetItem_setTextAlignment3, this.h, uintptr(alignment))
	runtime.KeepAlive(this)
}

func (this *QTableWidgetItem) Background() *QBrush {
	__r, _, _ := purego.SyscallN(addr_QTableWidgetItem_background, this.h)
	runtime.KeepAlive(this)
	return newownedQBrush(__r)
}

func (this *QTableWidgetItem) SetBackground(brush *QBrush) {
	purego.SyscallN(addr_QTableWidgetItem_setBackground, this.h, uintptr(brush.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(brush)
}

func (this *QTableWidgetItem) Foreground() *QBrush {
	__r, _, _ := purego.SyscallN(addr_QTableWidgetItem_foreground, this.h)
	runtime.KeepAlive(this)
	return newownedQBrush(__r)
}

func (this *QTableWidgetItem) SetForeground(brush *QBrush) {
	purego.SyscallN(addr_QTableWidgetItem_setForeground, this.h, uintptr(brush.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(brush)
}

func (this *QTableWidgetItem) CheckState() CheckState {
	__r0, _, _ := purego.SyscallN(addr_QTableWidgetItem_checkState, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return CheckState(__r)
}

func (this *QTableWidgetItem) SetCheckState(state CheckState) {
	purego.SyscallN(addr_QTableWidgetItem_setCheckState, this.h, uintptr(state))
	runtime.KeepAlive(this)
}

func (this *QTableWidgetItem) SizeHint() *QSize {
	__r, _, _ := purego.SyscallN(addr_QTableWidgetItem_sizeHint, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QTableWidgetItem) SetSizeHint(size *QSize) {
	purego.SyscallN(addr_QTableWidgetItem_setSizeHint, this.h, uintptr(size.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(size)
}

func (this *QTableWidgetItem) Data(role int) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QTableWidgetItem_data, this.h, uintptr(role))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QTableWidgetItem) SetData(role int, value *QVariant) {
	purego.SyscallN(addr_QTableWidgetItem_setData, this.h, uintptr(role), uintptr(value.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(value)
}

func (this *QTableWidgetItem) OperatorLesser(other *QTableWidgetItem) bool {
	__r0, _, _ := purego.SyscallN(addr_QTableWidgetItem_OperatorLesser, this.h, uintptr(other.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return __r != 0
}

func (this *QTableWidgetItem) Read(in *QDataStream) {
	purego.SyscallN(addr_QTableWidgetItem_read, this.h, uintptr(in.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(in)
}

func (this *QTableWidgetItem) Write(out *QDataStream) {
	purego.SyscallN(addr_QTableWidgetItem_write, this.h, uintptr(out.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(out)
}

func (this *QTableWidgetItem) OperatorAssign(other *QTableWidgetItem) *QTableWidgetItem {
	__r, _, _ := purego.SyscallN(addr_QTableWidgetItem_OperatorAssign, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return bindQTableWidgetItem(__r)
}

func (this *QTableWidgetItem) Type() int {
	__r0, _, _ := purego.SyscallN(addr_QTableWidgetItem_type, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTableWidgetItem) CloneDefault() *QTableWidgetItem {
	__r, _, _ := purego.SyscallN(addr_QTableWidgetItem_clone_default, this.h)
	runtime.KeepAlive(this)
	return bindQTableWidgetItem(__r)
}

func (this *QTableWidgetItem) DataDefault(role int) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QTableWidgetItem_data_default, this.h, uintptr(role))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QTableWidgetItem) SetDataDefault(role int, value *QVariant) {
	purego.SyscallN(addr_QTableWidgetItem_setData_default, this.h, uintptr(role), uintptr(value.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(value)
}

func (this *QTableWidgetItem) ReadDefault(in *QDataStream) {
	purego.SyscallN(addr_QTableWidgetItem_read_default, this.h, uintptr(in.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(in)
}

func (this *QTableWidgetItem) WriteDefault(out *QDataStream) {
	purego.SyscallN(addr_QTableWidgetItem_write_default, this.h, uintptr(out.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(out)
}

func QTableWidgetItem_virtIdx(slot int64) int32 {
	switch slot {
	case -7032813281048616478:
		return 0
	case -8837375942098466299:
		return 1
	case 35191454575202219:
		return 2
	case 5541207787354052101:
		return 3
	case -5099743077423337732:
		return 4
	}
	return -1
}

func QTableWidgetItem_enable(h uintptr, idx int32) { pg_QTableWidgetItem__set_override(h, idx) }

func (this *QTableWidgetItem) OnClone(fn func() *QTableWidgetItem) {
	registerVirt(this.UnsafeGoID(), -7032813281048616478, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "clone"
}

func (this *QTableWidgetItem) OnData(fn func(int) *QVariant) {
	registerVirt(this.UnsafeGoID(), -8837375942098466299, func(role int32) uintptr { return uintptr(fn(int(role)).UnsafePointer()) }) // slot "data"
}

func (this *QTableWidgetItem) OnSetData(fn func(int, *QVariant)) {
	registerVirt(this.UnsafeGoID(), 35191454575202219, func(role int32, value uintptr) { fn(int(role), UnsafeNewQVariant(unsafe.Pointer(value))) }) // slot "setData"
}

func (this *QTableWidgetItem) OnRead(fn func(*QDataStream)) {
	registerVirt(this.UnsafeGoID(), 5541207787354052101, func(in uintptr) { fn(UnsafeNewQDataStream(unsafe.Pointer(in))) }) // slot "read"
}

func (this *QTableWidgetItem) OnWrite(fn func(*QDataStream)) {
	registerVirt(this.UnsafeGoID(), -5099743077423337732, func(out uintptr) { fn(UnsafeNewQDataStream(unsafe.Pointer(out))) }) // slot "write"
}

func (this *QTableWidgetItem) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QTableWidgetItem_delete(this.h)
	}
	this.UnsafeRelease()
}
