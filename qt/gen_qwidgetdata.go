package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QWidgetData_new uintptr
	addr_QWidgetData_OperatorAssign uintptr
	addr_QWidgetData_winid_get uintptr
	addr_QWidgetData_winid_set uintptr
	addr_QWidgetData_widget_attributes_get uintptr
	addr_QWidgetData_widget_attributes_set uintptr
	addr_QWidgetData_window_flags_get uintptr
	addr_QWidgetData_window_flags_set uintptr
	addr_QWidgetData_window_state_get uintptr
	addr_QWidgetData_window_state_set uintptr
	addr_QWidgetData_focus_policy_get uintptr
	addr_QWidgetData_focus_policy_set uintptr
	addr_QWidgetData_sizehint_forced_get uintptr
	addr_QWidgetData_sizehint_forced_set uintptr
	addr_QWidgetData_is_closing_get uintptr
	addr_QWidgetData_is_closing_set uintptr
	addr_QWidgetData_in_show_get uintptr
	addr_QWidgetData_in_show_set uintptr
	addr_QWidgetData_in_set_window_state_get uintptr
	addr_QWidgetData_in_set_window_state_set uintptr
	addr_QWidgetData_fstrut_dirty_get uintptr
	addr_QWidgetData_fstrut_dirty_set uintptr
	addr_QWidgetData_context_menu_policy_get uintptr
	addr_QWidgetData_context_menu_policy_set uintptr
	addr_QWidgetData_window_modality_get uintptr
	addr_QWidgetData_window_modality_set uintptr
	addr_QWidgetData_in_destructor_get uintptr
	addr_QWidgetData_in_destructor_set uintptr
	addr_QWidgetData_unused_get uintptr
	addr_QWidgetData_unused_set uintptr
	addr_QWidgetData_crect_get uintptr
	addr_QWidgetData_crect_set uintptr
	addr_QWidgetData_pal_get uintptr
	addr_QWidgetData_pal_set uintptr
	addr_QWidgetData_fnt_get uintptr
	addr_QWidgetData_fnt_set uintptr
	addr_QWidgetData_wrect_get uintptr
	addr_QWidgetData_wrect_set uintptr
	pg_QWidgetData_delete func(uintptr)
)

func init() {
	symaddr(&addr_QWidgetData_new, "QWidgetData_new")
	symaddr(&addr_QWidgetData_OperatorAssign, "QWidgetData_OperatorAssign")
	symaddr(&addr_QWidgetData_winid_get, "QWidgetData_winid_get")
	symaddr(&addr_QWidgetData_winid_set, "QWidgetData_winid_set")
	symaddr(&addr_QWidgetData_widget_attributes_get, "QWidgetData_widget_attributes_get")
	symaddr(&addr_QWidgetData_widget_attributes_set, "QWidgetData_widget_attributes_set")
	symaddr(&addr_QWidgetData_window_flags_get, "QWidgetData_window_flags_get")
	symaddr(&addr_QWidgetData_window_flags_set, "QWidgetData_window_flags_set")
	symaddr(&addr_QWidgetData_window_state_get, "QWidgetData_window_state_get")
	symaddr(&addr_QWidgetData_window_state_set, "QWidgetData_window_state_set")
	symaddr(&addr_QWidgetData_focus_policy_get, "QWidgetData_focus_policy_get")
	symaddr(&addr_QWidgetData_focus_policy_set, "QWidgetData_focus_policy_set")
	symaddr(&addr_QWidgetData_sizehint_forced_get, "QWidgetData_sizehint_forced_get")
	symaddr(&addr_QWidgetData_sizehint_forced_set, "QWidgetData_sizehint_forced_set")
	symaddr(&addr_QWidgetData_is_closing_get, "QWidgetData_is_closing_get")
	symaddr(&addr_QWidgetData_is_closing_set, "QWidgetData_is_closing_set")
	symaddr(&addr_QWidgetData_in_show_get, "QWidgetData_in_show_get")
	symaddr(&addr_QWidgetData_in_show_set, "QWidgetData_in_show_set")
	symaddr(&addr_QWidgetData_in_set_window_state_get, "QWidgetData_in_set_window_state_get")
	symaddr(&addr_QWidgetData_in_set_window_state_set, "QWidgetData_in_set_window_state_set")
	symaddr(&addr_QWidgetData_fstrut_dirty_get, "QWidgetData_fstrut_dirty_get")
	symaddr(&addr_QWidgetData_fstrut_dirty_set, "QWidgetData_fstrut_dirty_set")
	symaddr(&addr_QWidgetData_context_menu_policy_get, "QWidgetData_context_menu_policy_get")
	symaddr(&addr_QWidgetData_context_menu_policy_set, "QWidgetData_context_menu_policy_set")
	symaddr(&addr_QWidgetData_window_modality_get, "QWidgetData_window_modality_get")
	symaddr(&addr_QWidgetData_window_modality_set, "QWidgetData_window_modality_set")
	symaddr(&addr_QWidgetData_in_destructor_get, "QWidgetData_in_destructor_get")
	symaddr(&addr_QWidgetData_in_destructor_set, "QWidgetData_in_destructor_set")
	symaddr(&addr_QWidgetData_unused_get, "QWidgetData_unused_get")
	symaddr(&addr_QWidgetData_unused_set, "QWidgetData_unused_set")
	symaddr(&addr_QWidgetData_crect_get, "QWidgetData_crect_get")
	symaddr(&addr_QWidgetData_crect_set, "QWidgetData_crect_set")
	symaddr(&addr_QWidgetData_pal_get, "QWidgetData_pal_get")
	symaddr(&addr_QWidgetData_pal_set, "QWidgetData_pal_set")
	symaddr(&addr_QWidgetData_fnt_get, "QWidgetData_fnt_get")
	symaddr(&addr_QWidgetData_fnt_set, "QWidgetData_fnt_set")
	symaddr(&addr_QWidgetData_wrect_get, "QWidgetData_wrect_get")
	symaddr(&addr_QWidgetData_wrect_set, "QWidgetData_wrect_set")
	reg(&pg_QWidgetData_delete, "QWidgetData_delete")
}

type QWidgetData struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QWidgetData) UnsafeGoID() int64     { return this.vid }
func (this *QWidgetData) UnsafeSetGoID(id int64) { this.vid = id }

func bindQWidgetData(h uintptr) *QWidgetData {
	if h == 0 {
		return nil
	}
	return &QWidgetData{h: h}
}

func newownedQWidgetData(h uintptr) *QWidgetData {
	o := bindQWidgetData(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QWidgetData).Delete)
	}
	return o
}

func (this *QWidgetData) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QWidgetData) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQWidgetData(h unsafe.Pointer) *QWidgetData { return bindQWidgetData(uintptr(h)) }
func UnsafeNewQWidgetDataOwned(h unsafe.Pointer) *QWidgetData { return newownedQWidgetData(uintptr(h)) }

func (this *QWidgetData) UnsafeRelease() { this.h = 0 }

func NewQWidgetData(arg1 *QWidgetData) *QWidgetData {
	__r, _, _ := purego.SyscallN(addr_QWidgetData_new, uintptr(arg1.UnsafePointer()))
	__o := newownedQWidgetData(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QWidgetData) OperatorAssign(arg1 *QWidgetData) *QWidgetData {
	__r, _, _ := purego.SyscallN(addr_QWidgetData_OperatorAssign, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQWidgetData(__r)
}

func (this *QWidgetData) Winid() uint64 {
	__r0, _, _ := purego.SyscallN(addr_QWidgetData_winid_get, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return uint64(__r)
}

func (this *QWidgetData) SetWinid(winid uint64) {
	purego.SyscallN(addr_QWidgetData_winid_set, this.h, uintptr(winid))
	runtime.KeepAlive(this)
}

func (this *QWidgetData) Widget_attributes() uint {
	__r0, _, _ := purego.SyscallN(addr_QWidgetData_widget_attributes_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint(__r)
}

func (this *QWidgetData) SetWidget_attributes(widget_attributes uint) {
	purego.SyscallN(addr_QWidgetData_widget_attributes_set, this.h, uintptr(widget_attributes))
	runtime.KeepAlive(this)
}

func (this *QWidgetData) Window_flags() WindowType {
	__r0, _, _ := purego.SyscallN(addr_QWidgetData_window_flags_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return WindowType(__r)
}

func (this *QWidgetData) SetWindow_flags(window_flags WindowType) {
	purego.SyscallN(addr_QWidgetData_window_flags_set, this.h, uintptr(window_flags))
	runtime.KeepAlive(this)
}

func (this *QWidgetData) Window_state() uint {
	__r0, _, _ := purego.SyscallN(addr_QWidgetData_window_state_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint(__r)
}

func (this *QWidgetData) SetWindow_state(window_state uint) {
	purego.SyscallN(addr_QWidgetData_window_state_set, this.h, uintptr(window_state))
	runtime.KeepAlive(this)
}

func (this *QWidgetData) Focus_policy() uint {
	__r0, _, _ := purego.SyscallN(addr_QWidgetData_focus_policy_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint(__r)
}

func (this *QWidgetData) SetFocus_policy(focus_policy uint) {
	purego.SyscallN(addr_QWidgetData_focus_policy_set, this.h, uintptr(focus_policy))
	runtime.KeepAlive(this)
}

func (this *QWidgetData) Sizehint_forced() uint {
	__r0, _, _ := purego.SyscallN(addr_QWidgetData_sizehint_forced_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint(__r)
}

func (this *QWidgetData) SetSizehint_forced(sizehint_forced uint) {
	purego.SyscallN(addr_QWidgetData_sizehint_forced_set, this.h, uintptr(sizehint_forced))
	runtime.KeepAlive(this)
}

func (this *QWidgetData) Is_closing() uint {
	__r0, _, _ := purego.SyscallN(addr_QWidgetData_is_closing_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint(__r)
}

func (this *QWidgetData) SetIs_closing(is_closing uint) {
	purego.SyscallN(addr_QWidgetData_is_closing_set, this.h, uintptr(is_closing))
	runtime.KeepAlive(this)
}

func (this *QWidgetData) In_show() uint {
	__r0, _, _ := purego.SyscallN(addr_QWidgetData_in_show_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint(__r)
}

func (this *QWidgetData) SetIn_show(in_show uint) {
	purego.SyscallN(addr_QWidgetData_in_show_set, this.h, uintptr(in_show))
	runtime.KeepAlive(this)
}

func (this *QWidgetData) In_set_window_state() uint {
	__r0, _, _ := purego.SyscallN(addr_QWidgetData_in_set_window_state_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint(__r)
}

func (this *QWidgetData) SetIn_set_window_state(in_set_window_state uint) {
	purego.SyscallN(addr_QWidgetData_in_set_window_state_set, this.h, uintptr(in_set_window_state))
	runtime.KeepAlive(this)
}

func (this *QWidgetData) Fstrut_dirty() uint {
	__r0, _, _ := purego.SyscallN(addr_QWidgetData_fstrut_dirty_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint(__r)
}

func (this *QWidgetData) SetFstrut_dirty(fstrut_dirty uint) {
	purego.SyscallN(addr_QWidgetData_fstrut_dirty_set, this.h, uintptr(fstrut_dirty))
	runtime.KeepAlive(this)
}

func (this *QWidgetData) Context_menu_policy() uint {
	__r0, _, _ := purego.SyscallN(addr_QWidgetData_context_menu_policy_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint(__r)
}

func (this *QWidgetData) SetContext_menu_policy(context_menu_policy uint) {
	purego.SyscallN(addr_QWidgetData_context_menu_policy_set, this.h, uintptr(context_menu_policy))
	runtime.KeepAlive(this)
}

func (this *QWidgetData) Window_modality() uint {
	__r0, _, _ := purego.SyscallN(addr_QWidgetData_window_modality_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint(__r)
}

func (this *QWidgetData) SetWindow_modality(window_modality uint) {
	purego.SyscallN(addr_QWidgetData_window_modality_set, this.h, uintptr(window_modality))
	runtime.KeepAlive(this)
}

func (this *QWidgetData) In_destructor() uint {
	__r0, _, _ := purego.SyscallN(addr_QWidgetData_in_destructor_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint(__r)
}

func (this *QWidgetData) SetIn_destructor(in_destructor uint) {
	purego.SyscallN(addr_QWidgetData_in_destructor_set, this.h, uintptr(in_destructor))
	runtime.KeepAlive(this)
}

func (this *QWidgetData) Unused() uint {
	__r0, _, _ := purego.SyscallN(addr_QWidgetData_unused_get, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint(__r)
}

func (this *QWidgetData) SetUnused(unused uint) {
	purego.SyscallN(addr_QWidgetData_unused_set, this.h, uintptr(unused))
	runtime.KeepAlive(this)
}

func (this *QWidgetData) Crect() *QRect {
	__r, _, _ := purego.SyscallN(addr_QWidgetData_crect_get, this.h)
	runtime.KeepAlive(this)
	return newownedQRect(__r)
}

func (this *QWidgetData) SetCrect(crect *QRect) {
	purego.SyscallN(addr_QWidgetData_crect_set, this.h, uintptr(crect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(crect)
}

func (this *QWidgetData) Pal() *QPalette {
	__r, _, _ := purego.SyscallN(addr_QWidgetData_pal_get, this.h)
	runtime.KeepAlive(this)
	return newownedQPalette(__r)
}

func (this *QWidgetData) SetPal(pal *QPalette) {
	purego.SyscallN(addr_QWidgetData_pal_set, this.h, uintptr(pal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pal)
}

func (this *QWidgetData) Fnt() *QFont {
	__r, _, _ := purego.SyscallN(addr_QWidgetData_fnt_get, this.h)
	runtime.KeepAlive(this)
	return newownedQFont(__r)
}

func (this *QWidgetData) SetFnt(fnt *QFont) {
	purego.SyscallN(addr_QWidgetData_fnt_set, this.h, uintptr(fnt.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(fnt)
}

func (this *QWidgetData) Wrect() *QRect {
	__r, _, _ := purego.SyscallN(addr_QWidgetData_wrect_get, this.h)
	runtime.KeepAlive(this)
	return newownedQRect(__r)
}

func (this *QWidgetData) SetWrect(wrect *QRect) {
	purego.SyscallN(addr_QWidgetData_wrect_set, this.h, uintptr(wrect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(wrect)
}

func (this *QWidgetData) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QWidgetData_delete(this.h)
	}
	this.UnsafeRelease()
}
