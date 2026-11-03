package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QTreeWidgetItem_new uintptr
	addr_QTreeWidgetItem_new2 uintptr
	addr_QTreeWidgetItem_new3 uintptr
	addr_QTreeWidgetItem_new4 uintptr
	addr_QTreeWidgetItem_new5 uintptr
	addr_QTreeWidgetItem_new6 uintptr
	addr_QTreeWidgetItem_new7 uintptr
	addr_QTreeWidgetItem_new8 uintptr
	addr_QTreeWidgetItem_new9 uintptr
	addr_QTreeWidgetItem_new10 uintptr
	addr_QTreeWidgetItem_new11 uintptr
	addr_QTreeWidgetItem_new12 uintptr
	addr_QTreeWidgetItem_new13 uintptr
	addr_QTreeWidgetItem_new14 uintptr
	addr_QTreeWidgetItem_new15 uintptr
	addr_QTreeWidgetItem_new16 uintptr
	addr_QTreeWidgetItem_new17 uintptr
	addr_QTreeWidgetItem_clone uintptr
	addr_QTreeWidgetItem_treeWidget uintptr
	addr_QTreeWidgetItem_setSelected uintptr
	addr_QTreeWidgetItem_isSelected uintptr
	addr_QTreeWidgetItem_setHidden uintptr
	addr_QTreeWidgetItem_isHidden uintptr
	addr_QTreeWidgetItem_setExpanded uintptr
	addr_QTreeWidgetItem_isExpanded uintptr
	addr_QTreeWidgetItem_setFirstColumnSpanned uintptr
	addr_QTreeWidgetItem_isFirstColumnSpanned uintptr
	addr_QTreeWidgetItem_setDisabled uintptr
	addr_QTreeWidgetItem_isDisabled uintptr
	addr_QTreeWidgetItem_setChildIndicatorPolicy uintptr
	addr_QTreeWidgetItem_childIndicatorPolicy uintptr
	addr_QTreeWidgetItem_flags uintptr
	addr_QTreeWidgetItem_setFlags uintptr
	addr_QTreeWidgetItem_text uintptr
	addr_QTreeWidgetItem_setText uintptr
	addr_QTreeWidgetItem_icon uintptr
	addr_QTreeWidgetItem_setIcon uintptr
	addr_QTreeWidgetItem_statusTip uintptr
	addr_QTreeWidgetItem_setStatusTip uintptr
	addr_QTreeWidgetItem_toolTip uintptr
	addr_QTreeWidgetItem_setToolTip uintptr
	addr_QTreeWidgetItem_whatsThis uintptr
	addr_QTreeWidgetItem_setWhatsThis uintptr
	addr_QTreeWidgetItem_font uintptr
	addr_QTreeWidgetItem_setFont uintptr
	addr_QTreeWidgetItem_textAlignment uintptr
	addr_QTreeWidgetItem_setTextAlignment uintptr
	addr_QTreeWidgetItem_setTextAlignment2 uintptr
	addr_QTreeWidgetItem_setTextAlignment3 uintptr
	addr_QTreeWidgetItem_background uintptr
	addr_QTreeWidgetItem_setBackground uintptr
	addr_QTreeWidgetItem_foreground uintptr
	addr_QTreeWidgetItem_setForeground uintptr
	addr_QTreeWidgetItem_checkState uintptr
	addr_QTreeWidgetItem_setCheckState uintptr
	addr_QTreeWidgetItem_sizeHint uintptr
	addr_QTreeWidgetItem_setSizeHint uintptr
	addr_QTreeWidgetItem_data uintptr
	addr_QTreeWidgetItem_setData uintptr
	addr_QTreeWidgetItem_OperatorLesser uintptr
	addr_QTreeWidgetItem_read uintptr
	addr_QTreeWidgetItem_write uintptr
	addr_QTreeWidgetItem_OperatorAssign uintptr
	addr_QTreeWidgetItem_parent uintptr
	addr_QTreeWidgetItem_child uintptr
	addr_QTreeWidgetItem_childCount uintptr
	addr_QTreeWidgetItem_columnCount uintptr
	addr_QTreeWidgetItem_indexOfChild uintptr
	addr_QTreeWidgetItem_addChild uintptr
	addr_QTreeWidgetItem_insertChild uintptr
	addr_QTreeWidgetItem_removeChild uintptr
	addr_QTreeWidgetItem_takeChild uintptr
	addr_QTreeWidgetItem_addChildren uintptr
	addr_QTreeWidgetItem_insertChildren uintptr
	addr_QTreeWidgetItem_takeChildren uintptr
	addr_QTreeWidgetItem_type uintptr
	addr_QTreeWidgetItem_sortChildren uintptr
	addr_QTreeWidgetItem_emitDataChanged uintptr
	addr_QTreeWidgetItem_clone_default uintptr
	addr_QTreeWidgetItem_data_default uintptr
	addr_QTreeWidgetItem_setData_default uintptr
	addr_QTreeWidgetItem_read_default uintptr
	addr_QTreeWidgetItem_write_default uintptr
	pg_QTreeWidgetItem__set_override func(uintptr, int32)
	pg_QTreeWidgetItem_delete func(uintptr)
)

func init() {
	symaddr(&addr_QTreeWidgetItem_new, "QTreeWidgetItem_new")
	symaddr(&addr_QTreeWidgetItem_new2, "QTreeWidgetItem_new2")
	symaddr(&addr_QTreeWidgetItem_new3, "QTreeWidgetItem_new3")
	symaddr(&addr_QTreeWidgetItem_new4, "QTreeWidgetItem_new4")
	symaddr(&addr_QTreeWidgetItem_new5, "QTreeWidgetItem_new5")
	symaddr(&addr_QTreeWidgetItem_new6, "QTreeWidgetItem_new6")
	symaddr(&addr_QTreeWidgetItem_new7, "QTreeWidgetItem_new7")
	symaddr(&addr_QTreeWidgetItem_new8, "QTreeWidgetItem_new8")
	symaddr(&addr_QTreeWidgetItem_new9, "QTreeWidgetItem_new9")
	symaddr(&addr_QTreeWidgetItem_new10, "QTreeWidgetItem_new10")
	symaddr(&addr_QTreeWidgetItem_new11, "QTreeWidgetItem_new11")
	symaddr(&addr_QTreeWidgetItem_new12, "QTreeWidgetItem_new12")
	symaddr(&addr_QTreeWidgetItem_new13, "QTreeWidgetItem_new13")
	symaddr(&addr_QTreeWidgetItem_new14, "QTreeWidgetItem_new14")
	symaddr(&addr_QTreeWidgetItem_new15, "QTreeWidgetItem_new15")
	symaddr(&addr_QTreeWidgetItem_new16, "QTreeWidgetItem_new16")
	symaddr(&addr_QTreeWidgetItem_new17, "QTreeWidgetItem_new17")
	symaddr(&addr_QTreeWidgetItem_clone, "QTreeWidgetItem_clone")
	symaddr(&addr_QTreeWidgetItem_treeWidget, "QTreeWidgetItem_treeWidget")
	symaddr(&addr_QTreeWidgetItem_setSelected, "QTreeWidgetItem_setSelected")
	symaddr(&addr_QTreeWidgetItem_isSelected, "QTreeWidgetItem_isSelected")
	symaddr(&addr_QTreeWidgetItem_setHidden, "QTreeWidgetItem_setHidden")
	symaddr(&addr_QTreeWidgetItem_isHidden, "QTreeWidgetItem_isHidden")
	symaddr(&addr_QTreeWidgetItem_setExpanded, "QTreeWidgetItem_setExpanded")
	symaddr(&addr_QTreeWidgetItem_isExpanded, "QTreeWidgetItem_isExpanded")
	symaddr(&addr_QTreeWidgetItem_setFirstColumnSpanned, "QTreeWidgetItem_setFirstColumnSpanned")
	symaddr(&addr_QTreeWidgetItem_isFirstColumnSpanned, "QTreeWidgetItem_isFirstColumnSpanned")
	symaddr(&addr_QTreeWidgetItem_setDisabled, "QTreeWidgetItem_setDisabled")
	symaddr(&addr_QTreeWidgetItem_isDisabled, "QTreeWidgetItem_isDisabled")
	symaddr(&addr_QTreeWidgetItem_setChildIndicatorPolicy, "QTreeWidgetItem_setChildIndicatorPolicy")
	symaddr(&addr_QTreeWidgetItem_childIndicatorPolicy, "QTreeWidgetItem_childIndicatorPolicy")
	symaddr(&addr_QTreeWidgetItem_flags, "QTreeWidgetItem_flags")
	symaddr(&addr_QTreeWidgetItem_setFlags, "QTreeWidgetItem_setFlags")
	symaddr(&addr_QTreeWidgetItem_text, "QTreeWidgetItem_text")
	symaddr(&addr_QTreeWidgetItem_setText, "QTreeWidgetItem_setText")
	symaddr(&addr_QTreeWidgetItem_icon, "QTreeWidgetItem_icon")
	symaddr(&addr_QTreeWidgetItem_setIcon, "QTreeWidgetItem_setIcon")
	symaddr(&addr_QTreeWidgetItem_statusTip, "QTreeWidgetItem_statusTip")
	symaddr(&addr_QTreeWidgetItem_setStatusTip, "QTreeWidgetItem_setStatusTip")
	symaddr(&addr_QTreeWidgetItem_toolTip, "QTreeWidgetItem_toolTip")
	symaddr(&addr_QTreeWidgetItem_setToolTip, "QTreeWidgetItem_setToolTip")
	symaddr(&addr_QTreeWidgetItem_whatsThis, "QTreeWidgetItem_whatsThis")
	symaddr(&addr_QTreeWidgetItem_setWhatsThis, "QTreeWidgetItem_setWhatsThis")
	symaddr(&addr_QTreeWidgetItem_font, "QTreeWidgetItem_font")
	symaddr(&addr_QTreeWidgetItem_setFont, "QTreeWidgetItem_setFont")
	symaddr(&addr_QTreeWidgetItem_textAlignment, "QTreeWidgetItem_textAlignment")
	symaddr(&addr_QTreeWidgetItem_setTextAlignment, "QTreeWidgetItem_setTextAlignment")
	symaddr(&addr_QTreeWidgetItem_setTextAlignment2, "QTreeWidgetItem_setTextAlignment2")
	symaddr(&addr_QTreeWidgetItem_setTextAlignment3, "QTreeWidgetItem_setTextAlignment3")
	symaddr(&addr_QTreeWidgetItem_background, "QTreeWidgetItem_background")
	symaddr(&addr_QTreeWidgetItem_setBackground, "QTreeWidgetItem_setBackground")
	symaddr(&addr_QTreeWidgetItem_foreground, "QTreeWidgetItem_foreground")
	symaddr(&addr_QTreeWidgetItem_setForeground, "QTreeWidgetItem_setForeground")
	symaddr(&addr_QTreeWidgetItem_checkState, "QTreeWidgetItem_checkState")
	symaddr(&addr_QTreeWidgetItem_setCheckState, "QTreeWidgetItem_setCheckState")
	symaddr(&addr_QTreeWidgetItem_sizeHint, "QTreeWidgetItem_sizeHint")
	symaddr(&addr_QTreeWidgetItem_setSizeHint, "QTreeWidgetItem_setSizeHint")
	symaddr(&addr_QTreeWidgetItem_data, "QTreeWidgetItem_data")
	symaddr(&addr_QTreeWidgetItem_setData, "QTreeWidgetItem_setData")
	symaddr(&addr_QTreeWidgetItem_OperatorLesser, "QTreeWidgetItem_OperatorLesser")
	symaddr(&addr_QTreeWidgetItem_read, "QTreeWidgetItem_read")
	symaddr(&addr_QTreeWidgetItem_write, "QTreeWidgetItem_write")
	symaddr(&addr_QTreeWidgetItem_OperatorAssign, "QTreeWidgetItem_OperatorAssign")
	symaddr(&addr_QTreeWidgetItem_parent, "QTreeWidgetItem_parent")
	symaddr(&addr_QTreeWidgetItem_child, "QTreeWidgetItem_child")
	symaddr(&addr_QTreeWidgetItem_childCount, "QTreeWidgetItem_childCount")
	symaddr(&addr_QTreeWidgetItem_columnCount, "QTreeWidgetItem_columnCount")
	symaddr(&addr_QTreeWidgetItem_indexOfChild, "QTreeWidgetItem_indexOfChild")
	symaddr(&addr_QTreeWidgetItem_addChild, "QTreeWidgetItem_addChild")
	symaddr(&addr_QTreeWidgetItem_insertChild, "QTreeWidgetItem_insertChild")
	symaddr(&addr_QTreeWidgetItem_removeChild, "QTreeWidgetItem_removeChild")
	symaddr(&addr_QTreeWidgetItem_takeChild, "QTreeWidgetItem_takeChild")
	symaddr(&addr_QTreeWidgetItem_addChildren, "QTreeWidgetItem_addChildren")
	symaddr(&addr_QTreeWidgetItem_insertChildren, "QTreeWidgetItem_insertChildren")
	symaddr(&addr_QTreeWidgetItem_takeChildren, "QTreeWidgetItem_takeChildren")
	symaddr(&addr_QTreeWidgetItem_type, "QTreeWidgetItem_type")
	symaddr(&addr_QTreeWidgetItem_sortChildren, "QTreeWidgetItem_sortChildren")
	symaddr(&addr_QTreeWidgetItem_emitDataChanged, "QTreeWidgetItem_emitDataChanged")
	symaddr(&addr_QTreeWidgetItem_clone_default, "QTreeWidgetItem_clone_default")
	symaddr(&addr_QTreeWidgetItem_data_default, "QTreeWidgetItem_data_default")
	symaddr(&addr_QTreeWidgetItem_setData_default, "QTreeWidgetItem_setData_default")
	symaddr(&addr_QTreeWidgetItem_read_default, "QTreeWidgetItem_read_default")
	symaddr(&addr_QTreeWidgetItem_write_default, "QTreeWidgetItem_write_default")
	reg(&pg_QTreeWidgetItem__set_override, "QTreeWidgetItem__set_override")
	reg(&pg_QTreeWidgetItem_delete, "QTreeWidgetItem_delete")
}

type QTreeWidgetItem struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QTreeWidgetItem) UnsafeGoID() int64     { return this.vid }
func (this *QTreeWidgetItem) UnsafeSetGoID(id int64) { this.vid = id }

func bindQTreeWidgetItem(h uintptr) *QTreeWidgetItem {
	if h == 0 {
		return nil
	}
	return &QTreeWidgetItem{h: h}
}

func newownedQTreeWidgetItem(h uintptr) *QTreeWidgetItem {
	o := bindQTreeWidgetItem(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QTreeWidgetItem).Delete)
	}
	return o
}

func (this *QTreeWidgetItem) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QTreeWidgetItem) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQTreeWidgetItem(h unsafe.Pointer) *QTreeWidgetItem { return bindQTreeWidgetItem(uintptr(h)) }
func UnsafeNewQTreeWidgetItemOwned(h unsafe.Pointer) *QTreeWidgetItem { return newownedQTreeWidgetItem(uintptr(h)) }

func (this *QTreeWidgetItem) UnsafeRelease() { this.h = 0 }

func NewQTreeWidgetItem() *QTreeWidgetItem {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QTreeWidgetItem_new, uintptr(__goid))
	__o := newownedQTreeWidgetItem(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QTreeWidgetItem_virtIdx, QTreeWidgetItem_enable)
	return __o
}

func NewQTreeWidgetItem2(strings []string) *QTreeWidgetItem {
	__goid := nextVirtID()
	strings_arr := make([]uintptr, len(strings))
	for __i, __e := range strings {
		strings_arr[__i] = pgCString(__e)
	}
	var strings_p uintptr
	if len(strings_arr) > 0 {
		strings_p = uintptr(unsafe.Pointer(&strings_arr[0]))
	}
	__r, _, _ := purego.SyscallN(addr_QTreeWidgetItem_new2, uintptr(__goid), strings_p, uintptr(len(strings)))
	__o := newownedQTreeWidgetItem(__r)
	for __i := range strings_arr {
		pgFree(strings_arr[__i])
	}
	runtime.KeepAlive(strings_arr)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QTreeWidgetItem_virtIdx, QTreeWidgetItem_enable)
	return __o
}

func NewQTreeWidgetItem3(treeview *QTreeWidget) *QTreeWidgetItem {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QTreeWidgetItem_new3, uintptr(__goid), uintptr(treeview.UnsafePointer()))
	__o := newownedQTreeWidgetItem(__r)
	runtime.KeepAlive(treeview)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QTreeWidgetItem_virtIdx, QTreeWidgetItem_enable)
	return __o
}

func NewQTreeWidgetItem4(treeview *QTreeWidget, strings []string) *QTreeWidgetItem {
	__goid := nextVirtID()
	strings_arr := make([]uintptr, len(strings))
	for __i, __e := range strings {
		strings_arr[__i] = pgCString(__e)
	}
	var strings_p uintptr
	if len(strings_arr) > 0 {
		strings_p = uintptr(unsafe.Pointer(&strings_arr[0]))
	}
	__r, _, _ := purego.SyscallN(addr_QTreeWidgetItem_new4, uintptr(__goid), uintptr(treeview.UnsafePointer()), strings_p, uintptr(len(strings)))
	__o := newownedQTreeWidgetItem(__r)
	runtime.KeepAlive(treeview)
	for __i := range strings_arr {
		pgFree(strings_arr[__i])
	}
	runtime.KeepAlive(strings_arr)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QTreeWidgetItem_virtIdx, QTreeWidgetItem_enable)
	return __o
}

func NewQTreeWidgetItem5(treeview *QTreeWidget, after *QTreeWidgetItem) *QTreeWidgetItem {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QTreeWidgetItem_new5, uintptr(__goid), uintptr(treeview.UnsafePointer()), uintptr(after.UnsafePointer()))
	__o := newownedQTreeWidgetItem(__r)
	runtime.KeepAlive(treeview)
	runtime.KeepAlive(after)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QTreeWidgetItem_virtIdx, QTreeWidgetItem_enable)
	return __o
}

func NewQTreeWidgetItem6(parent *QTreeWidgetItem) *QTreeWidgetItem {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QTreeWidgetItem_new6, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQTreeWidgetItem(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QTreeWidgetItem_virtIdx, QTreeWidgetItem_enable)
	return __o
}

func NewQTreeWidgetItem7(parent *QTreeWidgetItem, strings []string) *QTreeWidgetItem {
	__goid := nextVirtID()
	strings_arr := make([]uintptr, len(strings))
	for __i, __e := range strings {
		strings_arr[__i] = pgCString(__e)
	}
	var strings_p uintptr
	if len(strings_arr) > 0 {
		strings_p = uintptr(unsafe.Pointer(&strings_arr[0]))
	}
	__r, _, _ := purego.SyscallN(addr_QTreeWidgetItem_new7, uintptr(__goid), uintptr(parent.UnsafePointer()), strings_p, uintptr(len(strings)))
	__o := newownedQTreeWidgetItem(__r)
	runtime.KeepAlive(parent)
	for __i := range strings_arr {
		pgFree(strings_arr[__i])
	}
	runtime.KeepAlive(strings_arr)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QTreeWidgetItem_virtIdx, QTreeWidgetItem_enable)
	return __o
}

func NewQTreeWidgetItem8(parent *QTreeWidgetItem, after *QTreeWidgetItem) *QTreeWidgetItem {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QTreeWidgetItem_new8, uintptr(__goid), uintptr(parent.UnsafePointer()), uintptr(after.UnsafePointer()))
	__o := newownedQTreeWidgetItem(__r)
	runtime.KeepAlive(parent)
	runtime.KeepAlive(after)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QTreeWidgetItem_virtIdx, QTreeWidgetItem_enable)
	return __o
}

func NewQTreeWidgetItem9(other *QTreeWidgetItem) *QTreeWidgetItem {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QTreeWidgetItem_new9, uintptr(__goid), uintptr(other.UnsafePointer()))
	__o := newownedQTreeWidgetItem(__r)
	runtime.KeepAlive(other)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QTreeWidgetItem_virtIdx, QTreeWidgetItem_enable)
	return __o
}

func NewQTreeWidgetItem10(type_ int) *QTreeWidgetItem {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QTreeWidgetItem_new10, uintptr(__goid), uintptr(type_))
	__o := newownedQTreeWidgetItem(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QTreeWidgetItem_virtIdx, QTreeWidgetItem_enable)
	return __o
}

func NewQTreeWidgetItem11(strings []string, type_ int) *QTreeWidgetItem {
	__goid := nextVirtID()
	strings_arr := make([]uintptr, len(strings))
	for __i, __e := range strings {
		strings_arr[__i] = pgCString(__e)
	}
	var strings_p uintptr
	if len(strings_arr) > 0 {
		strings_p = uintptr(unsafe.Pointer(&strings_arr[0]))
	}
	__r, _, _ := purego.SyscallN(addr_QTreeWidgetItem_new11, uintptr(__goid), strings_p, uintptr(len(strings)), uintptr(type_))
	__o := newownedQTreeWidgetItem(__r)
	for __i := range strings_arr {
		pgFree(strings_arr[__i])
	}
	runtime.KeepAlive(strings_arr)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QTreeWidgetItem_virtIdx, QTreeWidgetItem_enable)
	return __o
}

func NewQTreeWidgetItem12(treeview *QTreeWidget, type_ int) *QTreeWidgetItem {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QTreeWidgetItem_new12, uintptr(__goid), uintptr(treeview.UnsafePointer()), uintptr(type_))
	__o := newownedQTreeWidgetItem(__r)
	runtime.KeepAlive(treeview)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QTreeWidgetItem_virtIdx, QTreeWidgetItem_enable)
	return __o
}

func NewQTreeWidgetItem13(treeview *QTreeWidget, strings []string, type_ int) *QTreeWidgetItem {
	__goid := nextVirtID()
	strings_arr := make([]uintptr, len(strings))
	for __i, __e := range strings {
		strings_arr[__i] = pgCString(__e)
	}
	var strings_p uintptr
	if len(strings_arr) > 0 {
		strings_p = uintptr(unsafe.Pointer(&strings_arr[0]))
	}
	__r, _, _ := purego.SyscallN(addr_QTreeWidgetItem_new13, uintptr(__goid), uintptr(treeview.UnsafePointer()), strings_p, uintptr(len(strings)), uintptr(type_))
	__o := newownedQTreeWidgetItem(__r)
	runtime.KeepAlive(treeview)
	for __i := range strings_arr {
		pgFree(strings_arr[__i])
	}
	runtime.KeepAlive(strings_arr)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QTreeWidgetItem_virtIdx, QTreeWidgetItem_enable)
	return __o
}

func NewQTreeWidgetItem14(treeview *QTreeWidget, after *QTreeWidgetItem, type_ int) *QTreeWidgetItem {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QTreeWidgetItem_new14, uintptr(__goid), uintptr(treeview.UnsafePointer()), uintptr(after.UnsafePointer()), uintptr(type_))
	__o := newownedQTreeWidgetItem(__r)
	runtime.KeepAlive(treeview)
	runtime.KeepAlive(after)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QTreeWidgetItem_virtIdx, QTreeWidgetItem_enable)
	return __o
}

func NewQTreeWidgetItem15(parent *QTreeWidgetItem, type_ int) *QTreeWidgetItem {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QTreeWidgetItem_new15, uintptr(__goid), uintptr(parent.UnsafePointer()), uintptr(type_))
	__o := newownedQTreeWidgetItem(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QTreeWidgetItem_virtIdx, QTreeWidgetItem_enable)
	return __o
}

func NewQTreeWidgetItem16(parent *QTreeWidgetItem, strings []string, type_ int) *QTreeWidgetItem {
	__goid := nextVirtID()
	strings_arr := make([]uintptr, len(strings))
	for __i, __e := range strings {
		strings_arr[__i] = pgCString(__e)
	}
	var strings_p uintptr
	if len(strings_arr) > 0 {
		strings_p = uintptr(unsafe.Pointer(&strings_arr[0]))
	}
	__r, _, _ := purego.SyscallN(addr_QTreeWidgetItem_new16, uintptr(__goid), uintptr(parent.UnsafePointer()), strings_p, uintptr(len(strings)), uintptr(type_))
	__o := newownedQTreeWidgetItem(__r)
	runtime.KeepAlive(parent)
	for __i := range strings_arr {
		pgFree(strings_arr[__i])
	}
	runtime.KeepAlive(strings_arr)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QTreeWidgetItem_virtIdx, QTreeWidgetItem_enable)
	return __o
}

func NewQTreeWidgetItem17(parent *QTreeWidgetItem, after *QTreeWidgetItem, type_ int) *QTreeWidgetItem {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QTreeWidgetItem_new17, uintptr(__goid), uintptr(parent.UnsafePointer()), uintptr(after.UnsafePointer()), uintptr(type_))
	__o := newownedQTreeWidgetItem(__r)
	runtime.KeepAlive(parent)
	runtime.KeepAlive(after)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QTreeWidgetItem_virtIdx, QTreeWidgetItem_enable)
	return __o
}

func (this *QTreeWidgetItem) Clone() *QTreeWidgetItem {
	__r, _, _ := purego.SyscallN(addr_QTreeWidgetItem_clone, this.h)
	runtime.KeepAlive(this)
	return bindQTreeWidgetItem(__r)
}

func (this *QTreeWidgetItem) TreeWidget() *QTreeWidget {
	__r, _, _ := purego.SyscallN(addr_QTreeWidgetItem_treeWidget, this.h)
	runtime.KeepAlive(this)
	return bindQTreeWidget(__r)
}

func (this *QTreeWidgetItem) SetSelected(select_ bool) {
	purego.SyscallN(addr_QTreeWidgetItem_setSelected, this.h, uintptr(b2i(select_)))
	runtime.KeepAlive(this)
}

func (this *QTreeWidgetItem) IsSelected() bool {
	__r0, _, _ := purego.SyscallN(addr_QTreeWidgetItem_isSelected, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTreeWidgetItem) SetHidden(hide bool) {
	purego.SyscallN(addr_QTreeWidgetItem_setHidden, this.h, uintptr(b2i(hide)))
	runtime.KeepAlive(this)
}

func (this *QTreeWidgetItem) IsHidden() bool {
	__r0, _, _ := purego.SyscallN(addr_QTreeWidgetItem_isHidden, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTreeWidgetItem) SetExpanded(expand bool) {
	purego.SyscallN(addr_QTreeWidgetItem_setExpanded, this.h, uintptr(b2i(expand)))
	runtime.KeepAlive(this)
}

func (this *QTreeWidgetItem) IsExpanded() bool {
	__r0, _, _ := purego.SyscallN(addr_QTreeWidgetItem_isExpanded, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTreeWidgetItem) SetFirstColumnSpanned(span bool) {
	purego.SyscallN(addr_QTreeWidgetItem_setFirstColumnSpanned, this.h, uintptr(b2i(span)))
	runtime.KeepAlive(this)
}

func (this *QTreeWidgetItem) IsFirstColumnSpanned() bool {
	__r0, _, _ := purego.SyscallN(addr_QTreeWidgetItem_isFirstColumnSpanned, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTreeWidgetItem) SetDisabled(disabled bool) {
	purego.SyscallN(addr_QTreeWidgetItem_setDisabled, this.h, uintptr(b2i(disabled)))
	runtime.KeepAlive(this)
}

func (this *QTreeWidgetItem) IsDisabled() bool {
	__r0, _, _ := purego.SyscallN(addr_QTreeWidgetItem_isDisabled, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTreeWidgetItem) SetChildIndicatorPolicy(policy QTreeWidgetItem__ChildIndicatorPolicy) {
	purego.SyscallN(addr_QTreeWidgetItem_setChildIndicatorPolicy, this.h, uintptr(policy))
	runtime.KeepAlive(this)
}

func (this *QTreeWidgetItem) ChildIndicatorPolicy() QTreeWidgetItem__ChildIndicatorPolicy {
	__r0, _, _ := purego.SyscallN(addr_QTreeWidgetItem_childIndicatorPolicy, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QTreeWidgetItem__ChildIndicatorPolicy(__r)
}

func (this *QTreeWidgetItem) Flags() ItemFlag {
	__r0, _, _ := purego.SyscallN(addr_QTreeWidgetItem_flags, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return ItemFlag(__r)
}

func (this *QTreeWidgetItem) SetFlags(flags ItemFlag) {
	purego.SyscallN(addr_QTreeWidgetItem_setFlags, this.h, uintptr(flags))
	runtime.KeepAlive(this)
}

func (this *QTreeWidgetItem) Text(column int) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTreeWidgetItem_text, this.h, uintptr(column), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QTreeWidgetItem) SetText(column int, text string) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QTreeWidgetItem_setText, this.h, uintptr(column), text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
}

func (this *QTreeWidgetItem) Icon(column int) *QIcon {
	__r, _, _ := purego.SyscallN(addr_QTreeWidgetItem_icon, this.h, uintptr(column))
	runtime.KeepAlive(this)
	return newownedQIcon(__r)
}

func (this *QTreeWidgetItem) SetIcon(column int, icon *QIcon) {
	purego.SyscallN(addr_QTreeWidgetItem_setIcon, this.h, uintptr(column), uintptr(icon.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(icon)
}

func (this *QTreeWidgetItem) StatusTip(column int) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTreeWidgetItem_statusTip, this.h, uintptr(column), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QTreeWidgetItem) SetStatusTip(column int, statusTip string) {
	var statusTip_p uintptr
	if len(statusTip) > 0 {
		statusTip_p = uintptr(unsafe.Pointer(unsafe.StringData(statusTip)))
	}
	purego.SyscallN(addr_QTreeWidgetItem_setStatusTip, this.h, uintptr(column), statusTip_p, uintptr(len(statusTip)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(statusTip)
}

func (this *QTreeWidgetItem) ToolTip(column int) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTreeWidgetItem_toolTip, this.h, uintptr(column), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QTreeWidgetItem) SetToolTip(column int, toolTip string) {
	var toolTip_p uintptr
	if len(toolTip) > 0 {
		toolTip_p = uintptr(unsafe.Pointer(unsafe.StringData(toolTip)))
	}
	purego.SyscallN(addr_QTreeWidgetItem_setToolTip, this.h, uintptr(column), toolTip_p, uintptr(len(toolTip)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(toolTip)
}

func (this *QTreeWidgetItem) WhatsThis(column int) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTreeWidgetItem_whatsThis, this.h, uintptr(column), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QTreeWidgetItem) SetWhatsThis(column int, whatsThis string) {
	var whatsThis_p uintptr
	if len(whatsThis) > 0 {
		whatsThis_p = uintptr(unsafe.Pointer(unsafe.StringData(whatsThis)))
	}
	purego.SyscallN(addr_QTreeWidgetItem_setWhatsThis, this.h, uintptr(column), whatsThis_p, uintptr(len(whatsThis)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(whatsThis)
}

func (this *QTreeWidgetItem) Font(column int) *QFont {
	__r, _, _ := purego.SyscallN(addr_QTreeWidgetItem_font, this.h, uintptr(column))
	runtime.KeepAlive(this)
	return newownedQFont(__r)
}

func (this *QTreeWidgetItem) SetFont(column int, font *QFont) {
	purego.SyscallN(addr_QTreeWidgetItem_setFont, this.h, uintptr(column), uintptr(font.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(font)
}

func (this *QTreeWidgetItem) TextAlignment(column int) int {
	__r0, _, _ := purego.SyscallN(addr_QTreeWidgetItem_textAlignment, this.h, uintptr(column))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTreeWidgetItem) SetTextAlignment(column int, alignment int) {
	purego.SyscallN(addr_QTreeWidgetItem_setTextAlignment, this.h, uintptr(column), uintptr(alignment))
	runtime.KeepAlive(this)
}

func (this *QTreeWidgetItem) SetTextAlignment2(column int, alignment AlignmentFlag) {
	purego.SyscallN(addr_QTreeWidgetItem_setTextAlignment2, this.h, uintptr(column), uintptr(alignment))
	runtime.KeepAlive(this)
}

func (this *QTreeWidgetItem) SetTextAlignment3(column int, alignment AlignmentFlag) {
	purego.SyscallN(addr_QTreeWidgetItem_setTextAlignment3, this.h, uintptr(column), uintptr(alignment))
	runtime.KeepAlive(this)
}

func (this *QTreeWidgetItem) Background(column int) *QBrush {
	__r, _, _ := purego.SyscallN(addr_QTreeWidgetItem_background, this.h, uintptr(column))
	runtime.KeepAlive(this)
	return newownedQBrush(__r)
}

func (this *QTreeWidgetItem) SetBackground(column int, brush *QBrush) {
	purego.SyscallN(addr_QTreeWidgetItem_setBackground, this.h, uintptr(column), uintptr(brush.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(brush)
}

func (this *QTreeWidgetItem) Foreground(column int) *QBrush {
	__r, _, _ := purego.SyscallN(addr_QTreeWidgetItem_foreground, this.h, uintptr(column))
	runtime.KeepAlive(this)
	return newownedQBrush(__r)
}

func (this *QTreeWidgetItem) SetForeground(column int, brush *QBrush) {
	purego.SyscallN(addr_QTreeWidgetItem_setForeground, this.h, uintptr(column), uintptr(brush.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(brush)
}

func (this *QTreeWidgetItem) CheckState(column int) CheckState {
	__r0, _, _ := purego.SyscallN(addr_QTreeWidgetItem_checkState, this.h, uintptr(column))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return CheckState(__r)
}

func (this *QTreeWidgetItem) SetCheckState(column int, state CheckState) {
	purego.SyscallN(addr_QTreeWidgetItem_setCheckState, this.h, uintptr(column), uintptr(state))
	runtime.KeepAlive(this)
}

func (this *QTreeWidgetItem) SizeHint(column int) *QSize {
	__r, _, _ := purego.SyscallN(addr_QTreeWidgetItem_sizeHint, this.h, uintptr(column))
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QTreeWidgetItem) SetSizeHint(column int, size *QSize) {
	purego.SyscallN(addr_QTreeWidgetItem_setSizeHint, this.h, uintptr(column), uintptr(size.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(size)
}

func (this *QTreeWidgetItem) Data(column int, role int) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QTreeWidgetItem_data, this.h, uintptr(column), uintptr(role))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QTreeWidgetItem) SetData(column int, role int, value *QVariant) {
	purego.SyscallN(addr_QTreeWidgetItem_setData, this.h, uintptr(column), uintptr(role), uintptr(value.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(value)
}

func (this *QTreeWidgetItem) OperatorLesser(other *QTreeWidgetItem) bool {
	__r0, _, _ := purego.SyscallN(addr_QTreeWidgetItem_OperatorLesser, this.h, uintptr(other.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return __r != 0
}

func (this *QTreeWidgetItem) Read(in *QDataStream) {
	purego.SyscallN(addr_QTreeWidgetItem_read, this.h, uintptr(in.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(in)
}

func (this *QTreeWidgetItem) Write(out *QDataStream) {
	purego.SyscallN(addr_QTreeWidgetItem_write, this.h, uintptr(out.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(out)
}

func (this *QTreeWidgetItem) OperatorAssign(other *QTreeWidgetItem) *QTreeWidgetItem {
	__r, _, _ := purego.SyscallN(addr_QTreeWidgetItem_OperatorAssign, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return bindQTreeWidgetItem(__r)
}

func (this *QTreeWidgetItem) Parent() *QTreeWidgetItem {
	__r, _, _ := purego.SyscallN(addr_QTreeWidgetItem_parent, this.h)
	runtime.KeepAlive(this)
	return bindQTreeWidgetItem(__r)
}

func (this *QTreeWidgetItem) Child(index int) *QTreeWidgetItem {
	__r, _, _ := purego.SyscallN(addr_QTreeWidgetItem_child, this.h, uintptr(index))
	runtime.KeepAlive(this)
	return bindQTreeWidgetItem(__r)
}

func (this *QTreeWidgetItem) ChildCount() int {
	__r0, _, _ := purego.SyscallN(addr_QTreeWidgetItem_childCount, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTreeWidgetItem) ColumnCount() int {
	__r0, _, _ := purego.SyscallN(addr_QTreeWidgetItem_columnCount, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTreeWidgetItem) IndexOfChild(child *QTreeWidgetItem) int {
	__r0, _, _ := purego.SyscallN(addr_QTreeWidgetItem_indexOfChild, this.h, uintptr(child.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(child)
	return int(__r)
}

func (this *QTreeWidgetItem) AddChild(child *QTreeWidgetItem) {
	purego.SyscallN(addr_QTreeWidgetItem_addChild, this.h, uintptr(child.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(child)
}

func (this *QTreeWidgetItem) InsertChild(index int, child *QTreeWidgetItem) {
	purego.SyscallN(addr_QTreeWidgetItem_insertChild, this.h, uintptr(index), uintptr(child.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(child)
}

func (this *QTreeWidgetItem) RemoveChild(child *QTreeWidgetItem) {
	purego.SyscallN(addr_QTreeWidgetItem_removeChild, this.h, uintptr(child.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(child)
}

func (this *QTreeWidgetItem) TakeChild(index int) *QTreeWidgetItem {
	__r, _, _ := purego.SyscallN(addr_QTreeWidgetItem_takeChild, this.h, uintptr(index))
	runtime.KeepAlive(this)
	return bindQTreeWidgetItem(__r)
}

func (this *QTreeWidgetItem) AddChildren(children []*QTreeWidgetItem) {
	children_arr := make([]uintptr, len(children))
	for __i, __e := range children {
		children_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var children_p uintptr
	if len(children_arr) > 0 {
		children_p = uintptr(unsafe.Pointer(&children_arr[0]))
	}
	purego.SyscallN(addr_QTreeWidgetItem_addChildren, this.h, children_p, uintptr(len(children)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(children)
	runtime.KeepAlive(children_arr)
}

func (this *QTreeWidgetItem) InsertChildren(index int, children []*QTreeWidgetItem) {
	children_arr := make([]uintptr, len(children))
	for __i, __e := range children {
		children_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var children_p uintptr
	if len(children_arr) > 0 {
		children_p = uintptr(unsafe.Pointer(&children_arr[0]))
	}
	purego.SyscallN(addr_QTreeWidgetItem_insertChildren, this.h, uintptr(index), children_p, uintptr(len(children)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(children)
	runtime.KeepAlive(children_arr)
}

func (this *QTreeWidgetItem) TakeChildren() []*QTreeWidgetItem {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTreeWidgetItem_takeChildren, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QTreeWidgetItem, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = bindQTreeWidgetItem(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QTreeWidgetItem) Type() int {
	__r0, _, _ := purego.SyscallN(addr_QTreeWidgetItem_type, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTreeWidgetItem) SortChildren(column int, order SortOrder) {
	purego.SyscallN(addr_QTreeWidgetItem_sortChildren, this.h, uintptr(column), uintptr(order))
	runtime.KeepAlive(this)
}

func (this *QTreeWidgetItem) EmitDataChanged() {
	purego.SyscallN(addr_QTreeWidgetItem_emitDataChanged, this.h)
	runtime.KeepAlive(this)
}

func (this *QTreeWidgetItem) CloneDefault() *QTreeWidgetItem {
	__r, _, _ := purego.SyscallN(addr_QTreeWidgetItem_clone_default, this.h)
	runtime.KeepAlive(this)
	return bindQTreeWidgetItem(__r)
}

func (this *QTreeWidgetItem) DataDefault(column int, role int) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QTreeWidgetItem_data_default, this.h, uintptr(column), uintptr(role))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QTreeWidgetItem) SetDataDefault(column int, role int, value *QVariant) {
	purego.SyscallN(addr_QTreeWidgetItem_setData_default, this.h, uintptr(column), uintptr(role), uintptr(value.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(value)
}

func (this *QTreeWidgetItem) ReadDefault(in *QDataStream) {
	purego.SyscallN(addr_QTreeWidgetItem_read_default, this.h, uintptr(in.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(in)
}

func (this *QTreeWidgetItem) WriteDefault(out *QDataStream) {
	purego.SyscallN(addr_QTreeWidgetItem_write_default, this.h, uintptr(out.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(out)
}

func QTreeWidgetItem_virtIdx(slot int64) int32 {
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

func QTreeWidgetItem_enable(h uintptr, idx int32) { pg_QTreeWidgetItem__set_override(h, idx) }

func (this *QTreeWidgetItem) OnClone(fn func() *QTreeWidgetItem) {
	registerVirt(this.UnsafeGoID(), -7032813281048616478, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "clone"
}

func (this *QTreeWidgetItem) OnData(fn func(int, int) *QVariant) {
	registerVirt(this.UnsafeGoID(), -8837375942098466299, func(column int32, role int32) uintptr { return uintptr(fn(int(column), int(role)).UnsafePointer()) }) // slot "data"
}

func (this *QTreeWidgetItem) OnSetData(fn func(int, int, *QVariant)) {
	registerVirt(this.UnsafeGoID(), 35191454575202219, func(column int32, role int32, value uintptr) { fn(int(column), int(role), UnsafeNewQVariant(unsafe.Pointer(value))) }) // slot "setData"
}

func (this *QTreeWidgetItem) OnRead(fn func(*QDataStream)) {
	registerVirt(this.UnsafeGoID(), 5541207787354052101, func(in uintptr) { fn(UnsafeNewQDataStream(unsafe.Pointer(in))) }) // slot "read"
}

func (this *QTreeWidgetItem) OnWrite(fn func(*QDataStream)) {
	registerVirt(this.UnsafeGoID(), -5099743077423337732, func(out uintptr) { fn(UnsafeNewQDataStream(unsafe.Pointer(out))) }) // slot "write"
}

func (this *QTreeWidgetItem) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QTreeWidgetItem_delete(this.h)
	}
	this.UnsafeRelease()
}
