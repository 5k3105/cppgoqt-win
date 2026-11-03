package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QListWidgetItem_new uintptr
	addr_QListWidgetItem_new2 uintptr
	addr_QListWidgetItem_new3 uintptr
	addr_QListWidgetItem_new4 uintptr
	addr_QListWidgetItem_new5 uintptr
	addr_QListWidgetItem_new6 uintptr
	addr_QListWidgetItem_new7 uintptr
	addr_QListWidgetItem_new8 uintptr
	addr_QListWidgetItem_new9 uintptr
	addr_QListWidgetItem_new10 uintptr
	addr_QListWidgetItem_clone uintptr
	addr_QListWidgetItem_listWidget uintptr
	addr_QListWidgetItem_setSelected uintptr
	addr_QListWidgetItem_isSelected uintptr
	addr_QListWidgetItem_setHidden uintptr
	addr_QListWidgetItem_isHidden uintptr
	addr_QListWidgetItem_flags uintptr
	addr_QListWidgetItem_setFlags uintptr
	addr_QListWidgetItem_text uintptr
	addr_QListWidgetItem_setText uintptr
	addr_QListWidgetItem_icon uintptr
	addr_QListWidgetItem_setIcon uintptr
	addr_QListWidgetItem_statusTip uintptr
	addr_QListWidgetItem_setStatusTip uintptr
	addr_QListWidgetItem_toolTip uintptr
	addr_QListWidgetItem_setToolTip uintptr
	addr_QListWidgetItem_whatsThis uintptr
	addr_QListWidgetItem_setWhatsThis uintptr
	addr_QListWidgetItem_font uintptr
	addr_QListWidgetItem_setFont uintptr
	addr_QListWidgetItem_textAlignment uintptr
	addr_QListWidgetItem_setTextAlignment uintptr
	addr_QListWidgetItem_setTextAlignment2 uintptr
	addr_QListWidgetItem_setTextAlignment3 uintptr
	addr_QListWidgetItem_background uintptr
	addr_QListWidgetItem_setBackground uintptr
	addr_QListWidgetItem_foreground uintptr
	addr_QListWidgetItem_setForeground uintptr
	addr_QListWidgetItem_checkState uintptr
	addr_QListWidgetItem_setCheckState uintptr
	addr_QListWidgetItem_sizeHint uintptr
	addr_QListWidgetItem_setSizeHint uintptr
	addr_QListWidgetItem_data uintptr
	addr_QListWidgetItem_setData uintptr
	addr_QListWidgetItem_OperatorLesser uintptr
	addr_QListWidgetItem_read uintptr
	addr_QListWidgetItem_write uintptr
	addr_QListWidgetItem_OperatorAssign uintptr
	addr_QListWidgetItem_type uintptr
	addr_QListWidgetItem_clone_default uintptr
	addr_QListWidgetItem_data_default uintptr
	addr_QListWidgetItem_setData_default uintptr
	addr_QListWidgetItem_read_default uintptr
	addr_QListWidgetItem_write_default uintptr
	pg_QListWidgetItem__set_override func(uintptr, int32)
	pg_QListWidgetItem_delete func(uintptr)
)

func init() {
	symaddr(&addr_QListWidgetItem_new, "QListWidgetItem_new")
	symaddr(&addr_QListWidgetItem_new2, "QListWidgetItem_new2")
	symaddr(&addr_QListWidgetItem_new3, "QListWidgetItem_new3")
	symaddr(&addr_QListWidgetItem_new4, "QListWidgetItem_new4")
	symaddr(&addr_QListWidgetItem_new5, "QListWidgetItem_new5")
	symaddr(&addr_QListWidgetItem_new6, "QListWidgetItem_new6")
	symaddr(&addr_QListWidgetItem_new7, "QListWidgetItem_new7")
	symaddr(&addr_QListWidgetItem_new8, "QListWidgetItem_new8")
	symaddr(&addr_QListWidgetItem_new9, "QListWidgetItem_new9")
	symaddr(&addr_QListWidgetItem_new10, "QListWidgetItem_new10")
	symaddr(&addr_QListWidgetItem_clone, "QListWidgetItem_clone")
	symaddr(&addr_QListWidgetItem_listWidget, "QListWidgetItem_listWidget")
	symaddr(&addr_QListWidgetItem_setSelected, "QListWidgetItem_setSelected")
	symaddr(&addr_QListWidgetItem_isSelected, "QListWidgetItem_isSelected")
	symaddr(&addr_QListWidgetItem_setHidden, "QListWidgetItem_setHidden")
	symaddr(&addr_QListWidgetItem_isHidden, "QListWidgetItem_isHidden")
	symaddr(&addr_QListWidgetItem_flags, "QListWidgetItem_flags")
	symaddr(&addr_QListWidgetItem_setFlags, "QListWidgetItem_setFlags")
	symaddr(&addr_QListWidgetItem_text, "QListWidgetItem_text")
	symaddr(&addr_QListWidgetItem_setText, "QListWidgetItem_setText")
	symaddr(&addr_QListWidgetItem_icon, "QListWidgetItem_icon")
	symaddr(&addr_QListWidgetItem_setIcon, "QListWidgetItem_setIcon")
	symaddr(&addr_QListWidgetItem_statusTip, "QListWidgetItem_statusTip")
	symaddr(&addr_QListWidgetItem_setStatusTip, "QListWidgetItem_setStatusTip")
	symaddr(&addr_QListWidgetItem_toolTip, "QListWidgetItem_toolTip")
	symaddr(&addr_QListWidgetItem_setToolTip, "QListWidgetItem_setToolTip")
	symaddr(&addr_QListWidgetItem_whatsThis, "QListWidgetItem_whatsThis")
	symaddr(&addr_QListWidgetItem_setWhatsThis, "QListWidgetItem_setWhatsThis")
	symaddr(&addr_QListWidgetItem_font, "QListWidgetItem_font")
	symaddr(&addr_QListWidgetItem_setFont, "QListWidgetItem_setFont")
	symaddr(&addr_QListWidgetItem_textAlignment, "QListWidgetItem_textAlignment")
	symaddr(&addr_QListWidgetItem_setTextAlignment, "QListWidgetItem_setTextAlignment")
	symaddr(&addr_QListWidgetItem_setTextAlignment2, "QListWidgetItem_setTextAlignment2")
	symaddr(&addr_QListWidgetItem_setTextAlignment3, "QListWidgetItem_setTextAlignment3")
	symaddr(&addr_QListWidgetItem_background, "QListWidgetItem_background")
	symaddr(&addr_QListWidgetItem_setBackground, "QListWidgetItem_setBackground")
	symaddr(&addr_QListWidgetItem_foreground, "QListWidgetItem_foreground")
	symaddr(&addr_QListWidgetItem_setForeground, "QListWidgetItem_setForeground")
	symaddr(&addr_QListWidgetItem_checkState, "QListWidgetItem_checkState")
	symaddr(&addr_QListWidgetItem_setCheckState, "QListWidgetItem_setCheckState")
	symaddr(&addr_QListWidgetItem_sizeHint, "QListWidgetItem_sizeHint")
	symaddr(&addr_QListWidgetItem_setSizeHint, "QListWidgetItem_setSizeHint")
	symaddr(&addr_QListWidgetItem_data, "QListWidgetItem_data")
	symaddr(&addr_QListWidgetItem_setData, "QListWidgetItem_setData")
	symaddr(&addr_QListWidgetItem_OperatorLesser, "QListWidgetItem_OperatorLesser")
	symaddr(&addr_QListWidgetItem_read, "QListWidgetItem_read")
	symaddr(&addr_QListWidgetItem_write, "QListWidgetItem_write")
	symaddr(&addr_QListWidgetItem_OperatorAssign, "QListWidgetItem_OperatorAssign")
	symaddr(&addr_QListWidgetItem_type, "QListWidgetItem_type")
	symaddr(&addr_QListWidgetItem_clone_default, "QListWidgetItem_clone_default")
	symaddr(&addr_QListWidgetItem_data_default, "QListWidgetItem_data_default")
	symaddr(&addr_QListWidgetItem_setData_default, "QListWidgetItem_setData_default")
	symaddr(&addr_QListWidgetItem_read_default, "QListWidgetItem_read_default")
	symaddr(&addr_QListWidgetItem_write_default, "QListWidgetItem_write_default")
	reg(&pg_QListWidgetItem__set_override, "QListWidgetItem__set_override")
	reg(&pg_QListWidgetItem_delete, "QListWidgetItem_delete")
}

type QListWidgetItem struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QListWidgetItem) UnsafeGoID() int64     { return this.vid }
func (this *QListWidgetItem) UnsafeSetGoID(id int64) { this.vid = id }

func bindQListWidgetItem(h uintptr) *QListWidgetItem {
	if h == 0 {
		return nil
	}
	return &QListWidgetItem{h: h}
}

func newownedQListWidgetItem(h uintptr) *QListWidgetItem {
	o := bindQListWidgetItem(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QListWidgetItem).Delete)
	}
	return o
}

func (this *QListWidgetItem) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QListWidgetItem) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQListWidgetItem(h unsafe.Pointer) *QListWidgetItem { return bindQListWidgetItem(uintptr(h)) }
func UnsafeNewQListWidgetItemOwned(h unsafe.Pointer) *QListWidgetItem { return newownedQListWidgetItem(uintptr(h)) }

func (this *QListWidgetItem) UnsafeRelease() { this.h = 0 }

func NewQListWidgetItem() *QListWidgetItem {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QListWidgetItem_new, uintptr(__goid))
	__o := newownedQListWidgetItem(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QListWidgetItem_virtIdx, QListWidgetItem_enable)
	return __o
}

func NewQListWidgetItem2(text string) *QListWidgetItem {
	__goid := nextVirtID()
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QListWidgetItem_new2, uintptr(__goid), text_p, uintptr(len(text)))
	__o := newownedQListWidgetItem(__r)
	runtime.KeepAlive(text)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QListWidgetItem_virtIdx, QListWidgetItem_enable)
	return __o
}

func NewQListWidgetItem3(icon *QIcon, text string) *QListWidgetItem {
	__goid := nextVirtID()
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QListWidgetItem_new3, uintptr(__goid), uintptr(icon.UnsafePointer()), text_p, uintptr(len(text)))
	__o := newownedQListWidgetItem(__r)
	runtime.KeepAlive(icon)
	runtime.KeepAlive(text)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QListWidgetItem_virtIdx, QListWidgetItem_enable)
	return __o
}

func NewQListWidgetItem4(other *QListWidgetItem) *QListWidgetItem {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QListWidgetItem_new4, uintptr(__goid), uintptr(other.UnsafePointer()))
	__o := newownedQListWidgetItem(__r)
	runtime.KeepAlive(other)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QListWidgetItem_virtIdx, QListWidgetItem_enable)
	return __o
}

func NewQListWidgetItem5(listview *QListWidget) *QListWidgetItem {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QListWidgetItem_new5, uintptr(__goid), uintptr(listview.UnsafePointer()))
	__o := newownedQListWidgetItem(__r)
	runtime.KeepAlive(listview)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QListWidgetItem_virtIdx, QListWidgetItem_enable)
	return __o
}

func NewQListWidgetItem6(listview *QListWidget, type_ int) *QListWidgetItem {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QListWidgetItem_new6, uintptr(__goid), uintptr(listview.UnsafePointer()), uintptr(type_))
	__o := newownedQListWidgetItem(__r)
	runtime.KeepAlive(listview)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QListWidgetItem_virtIdx, QListWidgetItem_enable)
	return __o
}

func NewQListWidgetItem7(text string, listview *QListWidget) *QListWidgetItem {
	__goid := nextVirtID()
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QListWidgetItem_new7, uintptr(__goid), text_p, uintptr(len(text)), uintptr(listview.UnsafePointer()))
	__o := newownedQListWidgetItem(__r)
	runtime.KeepAlive(text)
	runtime.KeepAlive(listview)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QListWidgetItem_virtIdx, QListWidgetItem_enable)
	return __o
}

func NewQListWidgetItem8(text string, listview *QListWidget, type_ int) *QListWidgetItem {
	__goid := nextVirtID()
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QListWidgetItem_new8, uintptr(__goid), text_p, uintptr(len(text)), uintptr(listview.UnsafePointer()), uintptr(type_))
	__o := newownedQListWidgetItem(__r)
	runtime.KeepAlive(text)
	runtime.KeepAlive(listview)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QListWidgetItem_virtIdx, QListWidgetItem_enable)
	return __o
}

func NewQListWidgetItem9(icon *QIcon, text string, listview *QListWidget) *QListWidgetItem {
	__goid := nextVirtID()
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QListWidgetItem_new9, uintptr(__goid), uintptr(icon.UnsafePointer()), text_p, uintptr(len(text)), uintptr(listview.UnsafePointer()))
	__o := newownedQListWidgetItem(__r)
	runtime.KeepAlive(icon)
	runtime.KeepAlive(text)
	runtime.KeepAlive(listview)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QListWidgetItem_virtIdx, QListWidgetItem_enable)
	return __o
}

func NewQListWidgetItem10(icon *QIcon, text string, listview *QListWidget, type_ int) *QListWidgetItem {
	__goid := nextVirtID()
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QListWidgetItem_new10, uintptr(__goid), uintptr(icon.UnsafePointer()), text_p, uintptr(len(text)), uintptr(listview.UnsafePointer()), uintptr(type_))
	__o := newownedQListWidgetItem(__r)
	runtime.KeepAlive(icon)
	runtime.KeepAlive(text)
	runtime.KeepAlive(listview)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QListWidgetItem_virtIdx, QListWidgetItem_enable)
	return __o
}

func (this *QListWidgetItem) Clone() *QListWidgetItem {
	__r, _, _ := purego.SyscallN(addr_QListWidgetItem_clone, this.h)
	runtime.KeepAlive(this)
	return bindQListWidgetItem(__r)
}

func (this *QListWidgetItem) ListWidget() *QListWidget {
	__r, _, _ := purego.SyscallN(addr_QListWidgetItem_listWidget, this.h)
	runtime.KeepAlive(this)
	return bindQListWidget(__r)
}

func (this *QListWidgetItem) SetSelected(select_ bool) {
	purego.SyscallN(addr_QListWidgetItem_setSelected, this.h, uintptr(b2i(select_)))
	runtime.KeepAlive(this)
}

func (this *QListWidgetItem) IsSelected() bool {
	__r0, _, _ := purego.SyscallN(addr_QListWidgetItem_isSelected, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QListWidgetItem) SetHidden(hide bool) {
	purego.SyscallN(addr_QListWidgetItem_setHidden, this.h, uintptr(b2i(hide)))
	runtime.KeepAlive(this)
}

func (this *QListWidgetItem) IsHidden() bool {
	__r0, _, _ := purego.SyscallN(addr_QListWidgetItem_isHidden, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QListWidgetItem) Flags() ItemFlag {
	__r0, _, _ := purego.SyscallN(addr_QListWidgetItem_flags, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return ItemFlag(__r)
}

func (this *QListWidgetItem) SetFlags(flags ItemFlag) {
	purego.SyscallN(addr_QListWidgetItem_setFlags, this.h, uintptr(flags))
	runtime.KeepAlive(this)
}

func (this *QListWidgetItem) Text() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QListWidgetItem_text, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QListWidgetItem) SetText(text string) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QListWidgetItem_setText, this.h, text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
}

func (this *QListWidgetItem) Icon() *QIcon {
	__r, _, _ := purego.SyscallN(addr_QListWidgetItem_icon, this.h)
	runtime.KeepAlive(this)
	return newownedQIcon(__r)
}

func (this *QListWidgetItem) SetIcon(icon *QIcon) {
	purego.SyscallN(addr_QListWidgetItem_setIcon, this.h, uintptr(icon.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(icon)
}

func (this *QListWidgetItem) StatusTip() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QListWidgetItem_statusTip, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QListWidgetItem) SetStatusTip(statusTip string) {
	var statusTip_p uintptr
	if len(statusTip) > 0 {
		statusTip_p = uintptr(unsafe.Pointer(unsafe.StringData(statusTip)))
	}
	purego.SyscallN(addr_QListWidgetItem_setStatusTip, this.h, statusTip_p, uintptr(len(statusTip)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(statusTip)
}

func (this *QListWidgetItem) ToolTip() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QListWidgetItem_toolTip, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QListWidgetItem) SetToolTip(toolTip string) {
	var toolTip_p uintptr
	if len(toolTip) > 0 {
		toolTip_p = uintptr(unsafe.Pointer(unsafe.StringData(toolTip)))
	}
	purego.SyscallN(addr_QListWidgetItem_setToolTip, this.h, toolTip_p, uintptr(len(toolTip)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(toolTip)
}

func (this *QListWidgetItem) WhatsThis() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QListWidgetItem_whatsThis, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QListWidgetItem) SetWhatsThis(whatsThis string) {
	var whatsThis_p uintptr
	if len(whatsThis) > 0 {
		whatsThis_p = uintptr(unsafe.Pointer(unsafe.StringData(whatsThis)))
	}
	purego.SyscallN(addr_QListWidgetItem_setWhatsThis, this.h, whatsThis_p, uintptr(len(whatsThis)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(whatsThis)
}

func (this *QListWidgetItem) Font() *QFont {
	__r, _, _ := purego.SyscallN(addr_QListWidgetItem_font, this.h)
	runtime.KeepAlive(this)
	return newownedQFont(__r)
}

func (this *QListWidgetItem) SetFont(font *QFont) {
	purego.SyscallN(addr_QListWidgetItem_setFont, this.h, uintptr(font.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(font)
}

func (this *QListWidgetItem) TextAlignment() int {
	__r0, _, _ := purego.SyscallN(addr_QListWidgetItem_textAlignment, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QListWidgetItem) SetTextAlignment(alignment int) {
	purego.SyscallN(addr_QListWidgetItem_setTextAlignment, this.h, uintptr(alignment))
	runtime.KeepAlive(this)
}

func (this *QListWidgetItem) SetTextAlignmentWithAlignment(alignment AlignmentFlag) {
	purego.SyscallN(addr_QListWidgetItem_setTextAlignment2, this.h, uintptr(alignment))
	runtime.KeepAlive(this)
}

func (this *QListWidgetItem) SetTextAlignment2(alignment AlignmentFlag) {
	purego.SyscallN(addr_QListWidgetItem_setTextAlignment3, this.h, uintptr(alignment))
	runtime.KeepAlive(this)
}

func (this *QListWidgetItem) Background() *QBrush {
	__r, _, _ := purego.SyscallN(addr_QListWidgetItem_background, this.h)
	runtime.KeepAlive(this)
	return newownedQBrush(__r)
}

func (this *QListWidgetItem) SetBackground(brush *QBrush) {
	purego.SyscallN(addr_QListWidgetItem_setBackground, this.h, uintptr(brush.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(brush)
}

func (this *QListWidgetItem) Foreground() *QBrush {
	__r, _, _ := purego.SyscallN(addr_QListWidgetItem_foreground, this.h)
	runtime.KeepAlive(this)
	return newownedQBrush(__r)
}

func (this *QListWidgetItem) SetForeground(brush *QBrush) {
	purego.SyscallN(addr_QListWidgetItem_setForeground, this.h, uintptr(brush.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(brush)
}

func (this *QListWidgetItem) CheckState() CheckState {
	__r0, _, _ := purego.SyscallN(addr_QListWidgetItem_checkState, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return CheckState(__r)
}

func (this *QListWidgetItem) SetCheckState(state CheckState) {
	purego.SyscallN(addr_QListWidgetItem_setCheckState, this.h, uintptr(state))
	runtime.KeepAlive(this)
}

func (this *QListWidgetItem) SizeHint() *QSize {
	__r, _, _ := purego.SyscallN(addr_QListWidgetItem_sizeHint, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QListWidgetItem) SetSizeHint(size *QSize) {
	purego.SyscallN(addr_QListWidgetItem_setSizeHint, this.h, uintptr(size.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(size)
}

func (this *QListWidgetItem) Data(role int) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QListWidgetItem_data, this.h, uintptr(role))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QListWidgetItem) SetData(role int, value *QVariant) {
	purego.SyscallN(addr_QListWidgetItem_setData, this.h, uintptr(role), uintptr(value.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(value)
}

func (this *QListWidgetItem) OperatorLesser(other *QListWidgetItem) bool {
	__r0, _, _ := purego.SyscallN(addr_QListWidgetItem_OperatorLesser, this.h, uintptr(other.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return __r != 0
}

func (this *QListWidgetItem) Read(in *QDataStream) {
	purego.SyscallN(addr_QListWidgetItem_read, this.h, uintptr(in.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(in)
}

func (this *QListWidgetItem) Write(out *QDataStream) {
	purego.SyscallN(addr_QListWidgetItem_write, this.h, uintptr(out.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(out)
}

func (this *QListWidgetItem) OperatorAssign(other *QListWidgetItem) *QListWidgetItem {
	__r, _, _ := purego.SyscallN(addr_QListWidgetItem_OperatorAssign, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return bindQListWidgetItem(__r)
}

func (this *QListWidgetItem) Type() int {
	__r0, _, _ := purego.SyscallN(addr_QListWidgetItem_type, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QListWidgetItem) CloneDefault() *QListWidgetItem {
	__r, _, _ := purego.SyscallN(addr_QListWidgetItem_clone_default, this.h)
	runtime.KeepAlive(this)
	return bindQListWidgetItem(__r)
}

func (this *QListWidgetItem) DataDefault(role int) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QListWidgetItem_data_default, this.h, uintptr(role))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QListWidgetItem) SetDataDefault(role int, value *QVariant) {
	purego.SyscallN(addr_QListWidgetItem_setData_default, this.h, uintptr(role), uintptr(value.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(value)
}

func (this *QListWidgetItem) ReadDefault(in *QDataStream) {
	purego.SyscallN(addr_QListWidgetItem_read_default, this.h, uintptr(in.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(in)
}

func (this *QListWidgetItem) WriteDefault(out *QDataStream) {
	purego.SyscallN(addr_QListWidgetItem_write_default, this.h, uintptr(out.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(out)
}

func QListWidgetItem_virtIdx(slot int64) int32 {
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

func QListWidgetItem_enable(h uintptr, idx int32) { pg_QListWidgetItem__set_override(h, idx) }

func (this *QListWidgetItem) OnClone(fn func() *QListWidgetItem) {
	registerVirt(this.UnsafeGoID(), -7032813281048616478, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "clone"
}

func (this *QListWidgetItem) OnData(fn func(int) *QVariant) {
	registerVirt(this.UnsafeGoID(), -8837375942098466299, func(role int32) uintptr { return uintptr(fn(int(role)).UnsafePointer()) }) // slot "data"
}

func (this *QListWidgetItem) OnSetData(fn func(int, *QVariant)) {
	registerVirt(this.UnsafeGoID(), 35191454575202219, func(role int32, value uintptr) { fn(int(role), UnsafeNewQVariant(unsafe.Pointer(value))) }) // slot "setData"
}

func (this *QListWidgetItem) OnRead(fn func(*QDataStream)) {
	registerVirt(this.UnsafeGoID(), 5541207787354052101, func(in uintptr) { fn(UnsafeNewQDataStream(unsafe.Pointer(in))) }) // slot "read"
}

func (this *QListWidgetItem) OnWrite(fn func(*QDataStream)) {
	registerVirt(this.UnsafeGoID(), -5099743077423337732, func(out uintptr) { fn(UnsafeNewQDataStream(unsafe.Pointer(out))) }) // slot "write"
}

func (this *QListWidgetItem) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QListWidgetItem_delete(this.h)
	}
	this.UnsafeRelease()
}
