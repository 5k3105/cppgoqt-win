package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QAccessible__State_new uintptr
	addr_QAccessible__State_OperatorEqual uintptr
	addr_QAccessible__State_disabled_get uintptr
	addr_QAccessible__State_disabled_set uintptr
	addr_QAccessible__State_selected_get uintptr
	addr_QAccessible__State_selected_set uintptr
	addr_QAccessible__State_focusable_get uintptr
	addr_QAccessible__State_focusable_set uintptr
	addr_QAccessible__State_focused_get uintptr
	addr_QAccessible__State_focused_set uintptr
	addr_QAccessible__State_pressed_get uintptr
	addr_QAccessible__State_pressed_set uintptr
	addr_QAccessible__State_checkable_get uintptr
	addr_QAccessible__State_checkable_set uintptr
	addr_QAccessible__State_checked_get uintptr
	addr_QAccessible__State_checked_set uintptr
	addr_QAccessible__State_checkStateMixed_get uintptr
	addr_QAccessible__State_checkStateMixed_set uintptr
	addr_QAccessible__State_readOnly_get uintptr
	addr_QAccessible__State_readOnly_set uintptr
	addr_QAccessible__State_hotTracked_get uintptr
	addr_QAccessible__State_hotTracked_set uintptr
	addr_QAccessible__State_defaultButton_get uintptr
	addr_QAccessible__State_defaultButton_set uintptr
	addr_QAccessible__State_expanded_get uintptr
	addr_QAccessible__State_expanded_set uintptr
	addr_QAccessible__State_collapsed_get uintptr
	addr_QAccessible__State_collapsed_set uintptr
	addr_QAccessible__State_busy_get uintptr
	addr_QAccessible__State_busy_set uintptr
	addr_QAccessible__State_expandable_get uintptr
	addr_QAccessible__State_expandable_set uintptr
	addr_QAccessible__State_marqueed_get uintptr
	addr_QAccessible__State_marqueed_set uintptr
	addr_QAccessible__State_animated_get uintptr
	addr_QAccessible__State_animated_set uintptr
	addr_QAccessible__State_invisible_get uintptr
	addr_QAccessible__State_invisible_set uintptr
	addr_QAccessible__State_offscreen_get uintptr
	addr_QAccessible__State_offscreen_set uintptr
	addr_QAccessible__State_sizeable_get uintptr
	addr_QAccessible__State_sizeable_set uintptr
	addr_QAccessible__State_movable_get uintptr
	addr_QAccessible__State_movable_set uintptr
	addr_QAccessible__State_selfVoicing_get uintptr
	addr_QAccessible__State_selfVoicing_set uintptr
	addr_QAccessible__State_selectable_get uintptr
	addr_QAccessible__State_selectable_set uintptr
	addr_QAccessible__State_linked_get uintptr
	addr_QAccessible__State_linked_set uintptr
	addr_QAccessible__State_traversed_get uintptr
	addr_QAccessible__State_traversed_set uintptr
	addr_QAccessible__State_multiSelectable_get uintptr
	addr_QAccessible__State_multiSelectable_set uintptr
	addr_QAccessible__State_extSelectable_get uintptr
	addr_QAccessible__State_extSelectable_set uintptr
	addr_QAccessible__State_passwordEdit_get uintptr
	addr_QAccessible__State_passwordEdit_set uintptr
	addr_QAccessible__State_hasPopup_get uintptr
	addr_QAccessible__State_hasPopup_set uintptr
	addr_QAccessible__State_modal_get uintptr
	addr_QAccessible__State_modal_set uintptr
	addr_QAccessible__State_active_get uintptr
	addr_QAccessible__State_active_set uintptr
	addr_QAccessible__State_invalid_get uintptr
	addr_QAccessible__State_invalid_set uintptr
	addr_QAccessible__State_editable_get uintptr
	addr_QAccessible__State_editable_set uintptr
	addr_QAccessible__State_multiLine_get uintptr
	addr_QAccessible__State_multiLine_set uintptr
	addr_QAccessible__State_selectableText_get uintptr
	addr_QAccessible__State_selectableText_set uintptr
	addr_QAccessible__State_supportsAutoCompletion_get uintptr
	addr_QAccessible__State_supportsAutoCompletion_set uintptr
	addr_QAccessible__State_searchEdit_get uintptr
	addr_QAccessible__State_searchEdit_set uintptr
	addr_QAccessible__State_qt_reserved_get uintptr
	addr_QAccessible__State_qt_reserved_set uintptr
	pg_QAccessible__State_delete func(uintptr)
)

func init() {
	symaddr(&addr_QAccessible__State_new, "QAccessible__State_new")
	symaddr(&addr_QAccessible__State_OperatorEqual, "QAccessible__State_OperatorEqual")
	symaddr(&addr_QAccessible__State_disabled_get, "QAccessible__State_disabled_get")
	symaddr(&addr_QAccessible__State_disabled_set, "QAccessible__State_disabled_set")
	symaddr(&addr_QAccessible__State_selected_get, "QAccessible__State_selected_get")
	symaddr(&addr_QAccessible__State_selected_set, "QAccessible__State_selected_set")
	symaddr(&addr_QAccessible__State_focusable_get, "QAccessible__State_focusable_get")
	symaddr(&addr_QAccessible__State_focusable_set, "QAccessible__State_focusable_set")
	symaddr(&addr_QAccessible__State_focused_get, "QAccessible__State_focused_get")
	symaddr(&addr_QAccessible__State_focused_set, "QAccessible__State_focused_set")
	symaddr(&addr_QAccessible__State_pressed_get, "QAccessible__State_pressed_get")
	symaddr(&addr_QAccessible__State_pressed_set, "QAccessible__State_pressed_set")
	symaddr(&addr_QAccessible__State_checkable_get, "QAccessible__State_checkable_get")
	symaddr(&addr_QAccessible__State_checkable_set, "QAccessible__State_checkable_set")
	symaddr(&addr_QAccessible__State_checked_get, "QAccessible__State_checked_get")
	symaddr(&addr_QAccessible__State_checked_set, "QAccessible__State_checked_set")
	symaddr(&addr_QAccessible__State_checkStateMixed_get, "QAccessible__State_checkStateMixed_get")
	symaddr(&addr_QAccessible__State_checkStateMixed_set, "QAccessible__State_checkStateMixed_set")
	symaddr(&addr_QAccessible__State_readOnly_get, "QAccessible__State_readOnly_get")
	symaddr(&addr_QAccessible__State_readOnly_set, "QAccessible__State_readOnly_set")
	symaddr(&addr_QAccessible__State_hotTracked_get, "QAccessible__State_hotTracked_get")
	symaddr(&addr_QAccessible__State_hotTracked_set, "QAccessible__State_hotTracked_set")
	symaddr(&addr_QAccessible__State_defaultButton_get, "QAccessible__State_defaultButton_get")
	symaddr(&addr_QAccessible__State_defaultButton_set, "QAccessible__State_defaultButton_set")
	symaddr(&addr_QAccessible__State_expanded_get, "QAccessible__State_expanded_get")
	symaddr(&addr_QAccessible__State_expanded_set, "QAccessible__State_expanded_set")
	symaddr(&addr_QAccessible__State_collapsed_get, "QAccessible__State_collapsed_get")
	symaddr(&addr_QAccessible__State_collapsed_set, "QAccessible__State_collapsed_set")
	symaddr(&addr_QAccessible__State_busy_get, "QAccessible__State_busy_get")
	symaddr(&addr_QAccessible__State_busy_set, "QAccessible__State_busy_set")
	symaddr(&addr_QAccessible__State_expandable_get, "QAccessible__State_expandable_get")
	symaddr(&addr_QAccessible__State_expandable_set, "QAccessible__State_expandable_set")
	symaddr(&addr_QAccessible__State_marqueed_get, "QAccessible__State_marqueed_get")
	symaddr(&addr_QAccessible__State_marqueed_set, "QAccessible__State_marqueed_set")
	symaddr(&addr_QAccessible__State_animated_get, "QAccessible__State_animated_get")
	symaddr(&addr_QAccessible__State_animated_set, "QAccessible__State_animated_set")
	symaddr(&addr_QAccessible__State_invisible_get, "QAccessible__State_invisible_get")
	symaddr(&addr_QAccessible__State_invisible_set, "QAccessible__State_invisible_set")
	symaddr(&addr_QAccessible__State_offscreen_get, "QAccessible__State_offscreen_get")
	symaddr(&addr_QAccessible__State_offscreen_set, "QAccessible__State_offscreen_set")
	symaddr(&addr_QAccessible__State_sizeable_get, "QAccessible__State_sizeable_get")
	symaddr(&addr_QAccessible__State_sizeable_set, "QAccessible__State_sizeable_set")
	symaddr(&addr_QAccessible__State_movable_get, "QAccessible__State_movable_get")
	symaddr(&addr_QAccessible__State_movable_set, "QAccessible__State_movable_set")
	symaddr(&addr_QAccessible__State_selfVoicing_get, "QAccessible__State_selfVoicing_get")
	symaddr(&addr_QAccessible__State_selfVoicing_set, "QAccessible__State_selfVoicing_set")
	symaddr(&addr_QAccessible__State_selectable_get, "QAccessible__State_selectable_get")
	symaddr(&addr_QAccessible__State_selectable_set, "QAccessible__State_selectable_set")
	symaddr(&addr_QAccessible__State_linked_get, "QAccessible__State_linked_get")
	symaddr(&addr_QAccessible__State_linked_set, "QAccessible__State_linked_set")
	symaddr(&addr_QAccessible__State_traversed_get, "QAccessible__State_traversed_get")
	symaddr(&addr_QAccessible__State_traversed_set, "QAccessible__State_traversed_set")
	symaddr(&addr_QAccessible__State_multiSelectable_get, "QAccessible__State_multiSelectable_get")
	symaddr(&addr_QAccessible__State_multiSelectable_set, "QAccessible__State_multiSelectable_set")
	symaddr(&addr_QAccessible__State_extSelectable_get, "QAccessible__State_extSelectable_get")
	symaddr(&addr_QAccessible__State_extSelectable_set, "QAccessible__State_extSelectable_set")
	symaddr(&addr_QAccessible__State_passwordEdit_get, "QAccessible__State_passwordEdit_get")
	symaddr(&addr_QAccessible__State_passwordEdit_set, "QAccessible__State_passwordEdit_set")
	symaddr(&addr_QAccessible__State_hasPopup_get, "QAccessible__State_hasPopup_get")
	symaddr(&addr_QAccessible__State_hasPopup_set, "QAccessible__State_hasPopup_set")
	symaddr(&addr_QAccessible__State_modal_get, "QAccessible__State_modal_get")
	symaddr(&addr_QAccessible__State_modal_set, "QAccessible__State_modal_set")
	symaddr(&addr_QAccessible__State_active_get, "QAccessible__State_active_get")
	symaddr(&addr_QAccessible__State_active_set, "QAccessible__State_active_set")
	symaddr(&addr_QAccessible__State_invalid_get, "QAccessible__State_invalid_get")
	symaddr(&addr_QAccessible__State_invalid_set, "QAccessible__State_invalid_set")
	symaddr(&addr_QAccessible__State_editable_get, "QAccessible__State_editable_get")
	symaddr(&addr_QAccessible__State_editable_set, "QAccessible__State_editable_set")
	symaddr(&addr_QAccessible__State_multiLine_get, "QAccessible__State_multiLine_get")
	symaddr(&addr_QAccessible__State_multiLine_set, "QAccessible__State_multiLine_set")
	symaddr(&addr_QAccessible__State_selectableText_get, "QAccessible__State_selectableText_get")
	symaddr(&addr_QAccessible__State_selectableText_set, "QAccessible__State_selectableText_set")
	symaddr(&addr_QAccessible__State_supportsAutoCompletion_get, "QAccessible__State_supportsAutoCompletion_get")
	symaddr(&addr_QAccessible__State_supportsAutoCompletion_set, "QAccessible__State_supportsAutoCompletion_set")
	symaddr(&addr_QAccessible__State_searchEdit_get, "QAccessible__State_searchEdit_get")
	symaddr(&addr_QAccessible__State_searchEdit_set, "QAccessible__State_searchEdit_set")
	symaddr(&addr_QAccessible__State_qt_reserved_get, "QAccessible__State_qt_reserved_get")
	symaddr(&addr_QAccessible__State_qt_reserved_set, "QAccessible__State_qt_reserved_set")
	reg(&pg_QAccessible__State_delete, "QAccessible__State_delete")
}

type QAccessible__State struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QAccessible__State) UnsafeGoID() int64     { return this.vid }
func (this *QAccessible__State) UnsafeSetGoID(id int64) { this.vid = id }

func bindQAccessible__State(h uintptr) *QAccessible__State {
	if h == 0 {
		return nil
	}
	return &QAccessible__State{h: h}
}

func newownedQAccessible__State(h uintptr) *QAccessible__State {
	o := bindQAccessible__State(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QAccessible__State).Delete)
	}
	return o
}

func (this *QAccessible__State) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QAccessible__State) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQAccessible__State(h unsafe.Pointer) *QAccessible__State { return bindQAccessible__State(uintptr(h)) }
func UnsafeNewQAccessible__StateOwned(h unsafe.Pointer) *QAccessible__State { return newownedQAccessible__State(uintptr(h)) }

func (this *QAccessible__State) UnsafeRelease() { this.h = 0 }

func NewQAccessible__State() *QAccessible__State {
	__r, _, _ := purego.SyscallN(addr_QAccessible__State_new)
	__o := newownedQAccessible__State(__r)
	return __o
}

func (this *QAccessible__State) OperatorEqual(second *QAccessible__State) bool {
	__r0, _, _ := purego.SyscallN(addr_QAccessible__State_OperatorEqual, this.h, uintptr(second.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(second)
	return __r != 0
}

func (this *QAccessible__State) Disabled() uint64 {
	__r0, _, _ := purego.SyscallN(addr_QAccessible__State_disabled_get, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return uint64(__r)
}

func (this *QAccessible__State) SetDisabled(disabled uint64) {
	purego.SyscallN(addr_QAccessible__State_disabled_set, this.h, uintptr(disabled))
	runtime.KeepAlive(this)
}

func (this *QAccessible__State) Selected() uint64 {
	__r0, _, _ := purego.SyscallN(addr_QAccessible__State_selected_get, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return uint64(__r)
}

func (this *QAccessible__State) SetSelected(selected uint64) {
	purego.SyscallN(addr_QAccessible__State_selected_set, this.h, uintptr(selected))
	runtime.KeepAlive(this)
}

func (this *QAccessible__State) Focusable() uint64 {
	__r0, _, _ := purego.SyscallN(addr_QAccessible__State_focusable_get, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return uint64(__r)
}

func (this *QAccessible__State) SetFocusable(focusable uint64) {
	purego.SyscallN(addr_QAccessible__State_focusable_set, this.h, uintptr(focusable))
	runtime.KeepAlive(this)
}

func (this *QAccessible__State) Focused() uint64 {
	__r0, _, _ := purego.SyscallN(addr_QAccessible__State_focused_get, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return uint64(__r)
}

func (this *QAccessible__State) SetFocused(focused uint64) {
	purego.SyscallN(addr_QAccessible__State_focused_set, this.h, uintptr(focused))
	runtime.KeepAlive(this)
}

func (this *QAccessible__State) Pressed() uint64 {
	__r0, _, _ := purego.SyscallN(addr_QAccessible__State_pressed_get, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return uint64(__r)
}

func (this *QAccessible__State) SetPressed(pressed uint64) {
	purego.SyscallN(addr_QAccessible__State_pressed_set, this.h, uintptr(pressed))
	runtime.KeepAlive(this)
}

func (this *QAccessible__State) Checkable() uint64 {
	__r0, _, _ := purego.SyscallN(addr_QAccessible__State_checkable_get, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return uint64(__r)
}

func (this *QAccessible__State) SetCheckable(checkable uint64) {
	purego.SyscallN(addr_QAccessible__State_checkable_set, this.h, uintptr(checkable))
	runtime.KeepAlive(this)
}

func (this *QAccessible__State) Checked() uint64 {
	__r0, _, _ := purego.SyscallN(addr_QAccessible__State_checked_get, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return uint64(__r)
}

func (this *QAccessible__State) SetChecked(checked uint64) {
	purego.SyscallN(addr_QAccessible__State_checked_set, this.h, uintptr(checked))
	runtime.KeepAlive(this)
}

func (this *QAccessible__State) CheckStateMixed() uint64 {
	__r0, _, _ := purego.SyscallN(addr_QAccessible__State_checkStateMixed_get, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return uint64(__r)
}

func (this *QAccessible__State) SetCheckStateMixed(checkStateMixed uint64) {
	purego.SyscallN(addr_QAccessible__State_checkStateMixed_set, this.h, uintptr(checkStateMixed))
	runtime.KeepAlive(this)
}

func (this *QAccessible__State) ReadOnly() uint64 {
	__r0, _, _ := purego.SyscallN(addr_QAccessible__State_readOnly_get, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return uint64(__r)
}

func (this *QAccessible__State) SetReadOnly(readOnly uint64) {
	purego.SyscallN(addr_QAccessible__State_readOnly_set, this.h, uintptr(readOnly))
	runtime.KeepAlive(this)
}

func (this *QAccessible__State) HotTracked() uint64 {
	__r0, _, _ := purego.SyscallN(addr_QAccessible__State_hotTracked_get, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return uint64(__r)
}

func (this *QAccessible__State) SetHotTracked(hotTracked uint64) {
	purego.SyscallN(addr_QAccessible__State_hotTracked_set, this.h, uintptr(hotTracked))
	runtime.KeepAlive(this)
}

func (this *QAccessible__State) DefaultButton() uint64 {
	__r0, _, _ := purego.SyscallN(addr_QAccessible__State_defaultButton_get, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return uint64(__r)
}

func (this *QAccessible__State) SetDefaultButton(defaultButton uint64) {
	purego.SyscallN(addr_QAccessible__State_defaultButton_set, this.h, uintptr(defaultButton))
	runtime.KeepAlive(this)
}

func (this *QAccessible__State) Expanded() uint64 {
	__r0, _, _ := purego.SyscallN(addr_QAccessible__State_expanded_get, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return uint64(__r)
}

func (this *QAccessible__State) SetExpanded(expanded uint64) {
	purego.SyscallN(addr_QAccessible__State_expanded_set, this.h, uintptr(expanded))
	runtime.KeepAlive(this)
}

func (this *QAccessible__State) Collapsed() uint64 {
	__r0, _, _ := purego.SyscallN(addr_QAccessible__State_collapsed_get, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return uint64(__r)
}

func (this *QAccessible__State) SetCollapsed(collapsed uint64) {
	purego.SyscallN(addr_QAccessible__State_collapsed_set, this.h, uintptr(collapsed))
	runtime.KeepAlive(this)
}

func (this *QAccessible__State) Busy() uint64 {
	__r0, _, _ := purego.SyscallN(addr_QAccessible__State_busy_get, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return uint64(__r)
}

func (this *QAccessible__State) SetBusy(busy uint64) {
	purego.SyscallN(addr_QAccessible__State_busy_set, this.h, uintptr(busy))
	runtime.KeepAlive(this)
}

func (this *QAccessible__State) Expandable() uint64 {
	__r0, _, _ := purego.SyscallN(addr_QAccessible__State_expandable_get, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return uint64(__r)
}

func (this *QAccessible__State) SetExpandable(expandable uint64) {
	purego.SyscallN(addr_QAccessible__State_expandable_set, this.h, uintptr(expandable))
	runtime.KeepAlive(this)
}

func (this *QAccessible__State) Marqueed() uint64 {
	__r0, _, _ := purego.SyscallN(addr_QAccessible__State_marqueed_get, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return uint64(__r)
}

func (this *QAccessible__State) SetMarqueed(marqueed uint64) {
	purego.SyscallN(addr_QAccessible__State_marqueed_set, this.h, uintptr(marqueed))
	runtime.KeepAlive(this)
}

func (this *QAccessible__State) Animated() uint64 {
	__r0, _, _ := purego.SyscallN(addr_QAccessible__State_animated_get, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return uint64(__r)
}

func (this *QAccessible__State) SetAnimated(animated uint64) {
	purego.SyscallN(addr_QAccessible__State_animated_set, this.h, uintptr(animated))
	runtime.KeepAlive(this)
}

func (this *QAccessible__State) Invisible() uint64 {
	__r0, _, _ := purego.SyscallN(addr_QAccessible__State_invisible_get, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return uint64(__r)
}

func (this *QAccessible__State) SetInvisible(invisible uint64) {
	purego.SyscallN(addr_QAccessible__State_invisible_set, this.h, uintptr(invisible))
	runtime.KeepAlive(this)
}

func (this *QAccessible__State) Offscreen() uint64 {
	__r0, _, _ := purego.SyscallN(addr_QAccessible__State_offscreen_get, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return uint64(__r)
}

func (this *QAccessible__State) SetOffscreen(offscreen uint64) {
	purego.SyscallN(addr_QAccessible__State_offscreen_set, this.h, uintptr(offscreen))
	runtime.KeepAlive(this)
}

func (this *QAccessible__State) Sizeable() uint64 {
	__r0, _, _ := purego.SyscallN(addr_QAccessible__State_sizeable_get, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return uint64(__r)
}

func (this *QAccessible__State) SetSizeable(sizeable uint64) {
	purego.SyscallN(addr_QAccessible__State_sizeable_set, this.h, uintptr(sizeable))
	runtime.KeepAlive(this)
}

func (this *QAccessible__State) Movable() uint64 {
	__r0, _, _ := purego.SyscallN(addr_QAccessible__State_movable_get, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return uint64(__r)
}

func (this *QAccessible__State) SetMovable(movable uint64) {
	purego.SyscallN(addr_QAccessible__State_movable_set, this.h, uintptr(movable))
	runtime.KeepAlive(this)
}

func (this *QAccessible__State) SelfVoicing() uint64 {
	__r0, _, _ := purego.SyscallN(addr_QAccessible__State_selfVoicing_get, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return uint64(__r)
}

func (this *QAccessible__State) SetSelfVoicing(selfVoicing uint64) {
	purego.SyscallN(addr_QAccessible__State_selfVoicing_set, this.h, uintptr(selfVoicing))
	runtime.KeepAlive(this)
}

func (this *QAccessible__State) Selectable() uint64 {
	__r0, _, _ := purego.SyscallN(addr_QAccessible__State_selectable_get, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return uint64(__r)
}

func (this *QAccessible__State) SetSelectable(selectable uint64) {
	purego.SyscallN(addr_QAccessible__State_selectable_set, this.h, uintptr(selectable))
	runtime.KeepAlive(this)
}

func (this *QAccessible__State) Linked() uint64 {
	__r0, _, _ := purego.SyscallN(addr_QAccessible__State_linked_get, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return uint64(__r)
}

func (this *QAccessible__State) SetLinked(linked uint64) {
	purego.SyscallN(addr_QAccessible__State_linked_set, this.h, uintptr(linked))
	runtime.KeepAlive(this)
}

func (this *QAccessible__State) Traversed() uint64 {
	__r0, _, _ := purego.SyscallN(addr_QAccessible__State_traversed_get, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return uint64(__r)
}

func (this *QAccessible__State) SetTraversed(traversed uint64) {
	purego.SyscallN(addr_QAccessible__State_traversed_set, this.h, uintptr(traversed))
	runtime.KeepAlive(this)
}

func (this *QAccessible__State) MultiSelectable() uint64 {
	__r0, _, _ := purego.SyscallN(addr_QAccessible__State_multiSelectable_get, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return uint64(__r)
}

func (this *QAccessible__State) SetMultiSelectable(multiSelectable uint64) {
	purego.SyscallN(addr_QAccessible__State_multiSelectable_set, this.h, uintptr(multiSelectable))
	runtime.KeepAlive(this)
}

func (this *QAccessible__State) ExtSelectable() uint64 {
	__r0, _, _ := purego.SyscallN(addr_QAccessible__State_extSelectable_get, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return uint64(__r)
}

func (this *QAccessible__State) SetExtSelectable(extSelectable uint64) {
	purego.SyscallN(addr_QAccessible__State_extSelectable_set, this.h, uintptr(extSelectable))
	runtime.KeepAlive(this)
}

func (this *QAccessible__State) PasswordEdit() uint64 {
	__r0, _, _ := purego.SyscallN(addr_QAccessible__State_passwordEdit_get, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return uint64(__r)
}

func (this *QAccessible__State) SetPasswordEdit(passwordEdit uint64) {
	purego.SyscallN(addr_QAccessible__State_passwordEdit_set, this.h, uintptr(passwordEdit))
	runtime.KeepAlive(this)
}

func (this *QAccessible__State) HasPopup() uint64 {
	__r0, _, _ := purego.SyscallN(addr_QAccessible__State_hasPopup_get, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return uint64(__r)
}

func (this *QAccessible__State) SetHasPopup(hasPopup uint64) {
	purego.SyscallN(addr_QAccessible__State_hasPopup_set, this.h, uintptr(hasPopup))
	runtime.KeepAlive(this)
}

func (this *QAccessible__State) Modal() uint64 {
	__r0, _, _ := purego.SyscallN(addr_QAccessible__State_modal_get, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return uint64(__r)
}

func (this *QAccessible__State) SetModal(modal uint64) {
	purego.SyscallN(addr_QAccessible__State_modal_set, this.h, uintptr(modal))
	runtime.KeepAlive(this)
}

func (this *QAccessible__State) Active() uint64 {
	__r0, _, _ := purego.SyscallN(addr_QAccessible__State_active_get, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return uint64(__r)
}

func (this *QAccessible__State) SetActive(active uint64) {
	purego.SyscallN(addr_QAccessible__State_active_set, this.h, uintptr(active))
	runtime.KeepAlive(this)
}

func (this *QAccessible__State) Invalid() uint64 {
	__r0, _, _ := purego.SyscallN(addr_QAccessible__State_invalid_get, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return uint64(__r)
}

func (this *QAccessible__State) SetInvalid(invalid uint64) {
	purego.SyscallN(addr_QAccessible__State_invalid_set, this.h, uintptr(invalid))
	runtime.KeepAlive(this)
}

func (this *QAccessible__State) Editable() uint64 {
	__r0, _, _ := purego.SyscallN(addr_QAccessible__State_editable_get, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return uint64(__r)
}

func (this *QAccessible__State) SetEditable(editable uint64) {
	purego.SyscallN(addr_QAccessible__State_editable_set, this.h, uintptr(editable))
	runtime.KeepAlive(this)
}

func (this *QAccessible__State) MultiLine() uint64 {
	__r0, _, _ := purego.SyscallN(addr_QAccessible__State_multiLine_get, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return uint64(__r)
}

func (this *QAccessible__State) SetMultiLine(multiLine uint64) {
	purego.SyscallN(addr_QAccessible__State_multiLine_set, this.h, uintptr(multiLine))
	runtime.KeepAlive(this)
}

func (this *QAccessible__State) SelectableText() uint64 {
	__r0, _, _ := purego.SyscallN(addr_QAccessible__State_selectableText_get, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return uint64(__r)
}

func (this *QAccessible__State) SetSelectableText(selectableText uint64) {
	purego.SyscallN(addr_QAccessible__State_selectableText_set, this.h, uintptr(selectableText))
	runtime.KeepAlive(this)
}

func (this *QAccessible__State) SupportsAutoCompletion() uint64 {
	__r0, _, _ := purego.SyscallN(addr_QAccessible__State_supportsAutoCompletion_get, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return uint64(__r)
}

func (this *QAccessible__State) SetSupportsAutoCompletion(supportsAutoCompletion uint64) {
	purego.SyscallN(addr_QAccessible__State_supportsAutoCompletion_set, this.h, uintptr(supportsAutoCompletion))
	runtime.KeepAlive(this)
}

func (this *QAccessible__State) SearchEdit() uint64 {
	__r0, _, _ := purego.SyscallN(addr_QAccessible__State_searchEdit_get, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return uint64(__r)
}

func (this *QAccessible__State) SetSearchEdit(searchEdit uint64) {
	purego.SyscallN(addr_QAccessible__State_searchEdit_set, this.h, uintptr(searchEdit))
	runtime.KeepAlive(this)
}

func (this *QAccessible__State) Qt_reserved() uint64 {
	__r0, _, _ := purego.SyscallN(addr_QAccessible__State_qt_reserved_get, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return uint64(__r)
}

func (this *QAccessible__State) SetQt_reserved(qt_reserved uint64) {
	purego.SyscallN(addr_QAccessible__State_qt_reserved_set, this.h, uintptr(qt_reserved))
	runtime.KeepAlive(this)
}

func (this *QAccessible__State) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QAccessible__State_delete(this.h)
	}
	this.UnsafeRelease()
}
