package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QAction_base func(uintptr) uintptr
	addr_QAction_new uintptr
	addr_QAction_new2 uintptr
	addr_QAction_new3 uintptr
	addr_QAction_new4 uintptr
	addr_QAction_new5 uintptr
	addr_QAction_new6 uintptr
	addr_QAction_tr uintptr
	addr_QAction_associatedObjects uintptr
	addr_QAction_setActionGroup uintptr
	addr_QAction_actionGroup uintptr
	addr_QAction_setIcon uintptr
	addr_QAction_icon uintptr
	addr_QAction_setText uintptr
	addr_QAction_text uintptr
	addr_QAction_setIconText uintptr
	addr_QAction_iconText uintptr
	addr_QAction_setToolTip uintptr
	addr_QAction_toolTip uintptr
	addr_QAction_setStatusTip uintptr
	addr_QAction_statusTip uintptr
	addr_QAction_setWhatsThis uintptr
	addr_QAction_whatsThis uintptr
	addr_QAction_setPriority uintptr
	addr_QAction_priority uintptr
	addr_QAction_setSeparator uintptr
	addr_QAction_isSeparator uintptr
	addr_QAction_setShortcut uintptr
	addr_QAction_shortcut uintptr
	addr_QAction_setShortcuts uintptr
	addr_QAction_setShortcuts2 uintptr
	addr_QAction_shortcuts uintptr
	addr_QAction_setShortcutContext uintptr
	addr_QAction_shortcutContext uintptr
	addr_QAction_setAutoRepeat uintptr
	addr_QAction_autoRepeat uintptr
	addr_QAction_setFont uintptr
	addr_QAction_font uintptr
	addr_QAction_setCheckable uintptr
	addr_QAction_isCheckable uintptr
	addr_QAction_data uintptr
	addr_QAction_setData uintptr
	addr_QAction_isChecked uintptr
	addr_QAction_isEnabled uintptr
	addr_QAction_isVisible uintptr
	addr_QAction_activate uintptr
	addr_QAction_setMenuRole uintptr
	addr_QAction_menuRole uintptr
	addr_QAction_setIconVisibleInMenu uintptr
	addr_QAction_isIconVisibleInMenu uintptr
	addr_QAction_setShortcutVisibleInContextMenu uintptr
	addr_QAction_isShortcutVisibleInContextMenu uintptr
	addr_QAction_showStatusText uintptr
	addr_QAction_trigger uintptr
	addr_QAction_hover uintptr
	addr_QAction_setChecked uintptr
	addr_QAction_toggle uintptr
	addr_QAction_setEnabled uintptr
	addr_QAction_resetEnabled uintptr
	addr_QAction_setDisabled uintptr
	addr_QAction_setVisible uintptr
	addr_QAction_changed uintptr
	addr_QAction_enabledChanged uintptr
	addr_QAction_checkableChanged uintptr
	addr_QAction_visibleChanged uintptr
	addr_QAction_triggered uintptr
	addr_QAction_hovered uintptr
	addr_QAction_toggled uintptr
	addr_QAction_connect uintptr
	addr_QAction_connect2 uintptr
	addr_QAction_sender uintptr
	addr_QAction_senderSignalIndex uintptr
	addr_QAction_receivers uintptr
	addr_QAction_isSignalConnected uintptr
	addr_QAction_tr2 uintptr
	addr_QAction_tr3 uintptr
	addr_QAction_showStatusText2 uintptr
	addr_QAction_triggered2 uintptr
	addr_QAction_connect3 uintptr
	addr_QAction_connect4 uintptr
	addr_QAction_metaObject_default uintptr
	addr_QAction_event_default uintptr
	addr_QAction_eventFilter_default uintptr
	addr_QAction_timerEvent_default uintptr
	addr_QAction_childEvent_default uintptr
	addr_QAction_customEvent_default uintptr
	addr_QAction_connectNotify_default uintptr
	addr_QAction_disconnectNotify_default uintptr
	pg_QAction_connect_changed func(uintptr, int64)
	pg_QAction_connect_enabledChanged func(uintptr, int64)
	pg_QAction_connect_checkableChanged func(uintptr, int64)
	pg_QAction_connect_visibleChanged func(uintptr, int64)
	pg_QAction_connect_triggered func(uintptr, int64)
	pg_QAction_connect_triggered2 func(uintptr, int64)
	pg_QAction_connect_hovered func(uintptr, int64)
	pg_QAction_connect_toggled func(uintptr, int64)
	pg_QAction_connect_destroyed func(uintptr, int64)
	pg_QAction_connect_destroyed2 func(uintptr, int64)
	pg_QAction_connect_objectNameChanged func(uintptr, int64)
	pg_QAction__disconnect func(int64)
	pg_QAction__set_override func(uintptr, int32)
	pg_QAction_delete func(uintptr)
)

func init() {
	reg(&pg_QAction_base, "QAction_base")
	symaddr(&addr_QAction_new, "QAction_new")
	symaddr(&addr_QAction_new2, "QAction_new2")
	symaddr(&addr_QAction_new3, "QAction_new3")
	symaddr(&addr_QAction_new4, "QAction_new4")
	symaddr(&addr_QAction_new5, "QAction_new5")
	symaddr(&addr_QAction_new6, "QAction_new6")
	symaddr(&addr_QAction_tr, "QAction_tr")
	symaddr(&addr_QAction_associatedObjects, "QAction_associatedObjects")
	symaddr(&addr_QAction_setActionGroup, "QAction_setActionGroup")
	symaddr(&addr_QAction_actionGroup, "QAction_actionGroup")
	symaddr(&addr_QAction_setIcon, "QAction_setIcon")
	symaddr(&addr_QAction_icon, "QAction_icon")
	symaddr(&addr_QAction_setText, "QAction_setText")
	symaddr(&addr_QAction_text, "QAction_text")
	symaddr(&addr_QAction_setIconText, "QAction_setIconText")
	symaddr(&addr_QAction_iconText, "QAction_iconText")
	symaddr(&addr_QAction_setToolTip, "QAction_setToolTip")
	symaddr(&addr_QAction_toolTip, "QAction_toolTip")
	symaddr(&addr_QAction_setStatusTip, "QAction_setStatusTip")
	symaddr(&addr_QAction_statusTip, "QAction_statusTip")
	symaddr(&addr_QAction_setWhatsThis, "QAction_setWhatsThis")
	symaddr(&addr_QAction_whatsThis, "QAction_whatsThis")
	symaddr(&addr_QAction_setPriority, "QAction_setPriority")
	symaddr(&addr_QAction_priority, "QAction_priority")
	symaddr(&addr_QAction_setSeparator, "QAction_setSeparator")
	symaddr(&addr_QAction_isSeparator, "QAction_isSeparator")
	symaddr(&addr_QAction_setShortcut, "QAction_setShortcut")
	symaddr(&addr_QAction_shortcut, "QAction_shortcut")
	symaddr(&addr_QAction_setShortcuts, "QAction_setShortcuts")
	symaddr(&addr_QAction_setShortcuts2, "QAction_setShortcuts2")
	symaddr(&addr_QAction_shortcuts, "QAction_shortcuts")
	symaddr(&addr_QAction_setShortcutContext, "QAction_setShortcutContext")
	symaddr(&addr_QAction_shortcutContext, "QAction_shortcutContext")
	symaddr(&addr_QAction_setAutoRepeat, "QAction_setAutoRepeat")
	symaddr(&addr_QAction_autoRepeat, "QAction_autoRepeat")
	symaddr(&addr_QAction_setFont, "QAction_setFont")
	symaddr(&addr_QAction_font, "QAction_font")
	symaddr(&addr_QAction_setCheckable, "QAction_setCheckable")
	symaddr(&addr_QAction_isCheckable, "QAction_isCheckable")
	symaddr(&addr_QAction_data, "QAction_data")
	symaddr(&addr_QAction_setData, "QAction_setData")
	symaddr(&addr_QAction_isChecked, "QAction_isChecked")
	symaddr(&addr_QAction_isEnabled, "QAction_isEnabled")
	symaddr(&addr_QAction_isVisible, "QAction_isVisible")
	symaddr(&addr_QAction_activate, "QAction_activate")
	symaddr(&addr_QAction_setMenuRole, "QAction_setMenuRole")
	symaddr(&addr_QAction_menuRole, "QAction_menuRole")
	symaddr(&addr_QAction_setIconVisibleInMenu, "QAction_setIconVisibleInMenu")
	symaddr(&addr_QAction_isIconVisibleInMenu, "QAction_isIconVisibleInMenu")
	symaddr(&addr_QAction_setShortcutVisibleInContextMenu, "QAction_setShortcutVisibleInContextMenu")
	symaddr(&addr_QAction_isShortcutVisibleInContextMenu, "QAction_isShortcutVisibleInContextMenu")
	symaddr(&addr_QAction_showStatusText, "QAction_showStatusText")
	symaddr(&addr_QAction_trigger, "QAction_trigger")
	symaddr(&addr_QAction_hover, "QAction_hover")
	symaddr(&addr_QAction_setChecked, "QAction_setChecked")
	symaddr(&addr_QAction_toggle, "QAction_toggle")
	symaddr(&addr_QAction_setEnabled, "QAction_setEnabled")
	symaddr(&addr_QAction_resetEnabled, "QAction_resetEnabled")
	symaddr(&addr_QAction_setDisabled, "QAction_setDisabled")
	symaddr(&addr_QAction_setVisible, "QAction_setVisible")
	symaddr(&addr_QAction_changed, "QAction_changed")
	symaddr(&addr_QAction_enabledChanged, "QAction_enabledChanged")
	symaddr(&addr_QAction_checkableChanged, "QAction_checkableChanged")
	symaddr(&addr_QAction_visibleChanged, "QAction_visibleChanged")
	symaddr(&addr_QAction_triggered, "QAction_triggered")
	symaddr(&addr_QAction_hovered, "QAction_hovered")
	symaddr(&addr_QAction_toggled, "QAction_toggled")
	symaddr(&addr_QAction_connect, "QAction_connect")
	symaddr(&addr_QAction_connect2, "QAction_connect2")
	symaddr(&addr_QAction_sender, "QAction_sender")
	symaddr(&addr_QAction_senderSignalIndex, "QAction_senderSignalIndex")
	symaddr(&addr_QAction_receivers, "QAction_receivers")
	symaddr(&addr_QAction_isSignalConnected, "QAction_isSignalConnected")
	symaddr(&addr_QAction_tr2, "QAction_tr2")
	symaddr(&addr_QAction_tr3, "QAction_tr3")
	symaddr(&addr_QAction_showStatusText2, "QAction_showStatusText2")
	symaddr(&addr_QAction_triggered2, "QAction_triggered2")
	symaddr(&addr_QAction_connect3, "QAction_connect3")
	symaddr(&addr_QAction_connect4, "QAction_connect4")
	symaddr(&addr_QAction_metaObject_default, "QAction_metaObject_default")
	symaddr(&addr_QAction_event_default, "QAction_event_default")
	symaddr(&addr_QAction_eventFilter_default, "QAction_eventFilter_default")
	symaddr(&addr_QAction_timerEvent_default, "QAction_timerEvent_default")
	symaddr(&addr_QAction_childEvent_default, "QAction_childEvent_default")
	symaddr(&addr_QAction_customEvent_default, "QAction_customEvent_default")
	symaddr(&addr_QAction_connectNotify_default, "QAction_connectNotify_default")
	symaddr(&addr_QAction_disconnectNotify_default, "QAction_disconnectNotify_default")
	reg(&pg_QAction_connect_changed, "QAction_connect_changed")
	reg(&pg_QAction_connect_enabledChanged, "QAction_connect_enabledChanged")
	reg(&pg_QAction_connect_checkableChanged, "QAction_connect_checkableChanged")
	reg(&pg_QAction_connect_visibleChanged, "QAction_connect_visibleChanged")
	reg(&pg_QAction_connect_triggered, "QAction_connect_triggered")
	reg(&pg_QAction_connect_triggered2, "QAction_connect_triggered2")
	reg(&pg_QAction_connect_hovered, "QAction_connect_hovered")
	reg(&pg_QAction_connect_toggled, "QAction_connect_toggled")
	reg(&pg_QAction_connect_destroyed, "QAction_connect_destroyed")
	reg(&pg_QAction_connect_destroyed2, "QAction_connect_destroyed2")
	reg(&pg_QAction_connect_objectNameChanged, "QAction_connect_objectNameChanged")
	reg(&pg_QAction__disconnect, "QAction__disconnect")
	reg(&pg_QAction__set_override, "QAction__set_override")
	reg(&pg_QAction_delete, "QAction_delete")
}

type QAction struct {
	h uintptr
	*QObject
}

func bindQAction(h uintptr) *QAction {
	if h == 0 {
		return nil
	}
	return &QAction{h: h, QObject: bindQObject(pg_QAction_base(h))}
}

func newownedQAction(h uintptr) *QAction { return bindQAction(h) }

func (this *QAction) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QAction) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQAction(h unsafe.Pointer) *QAction { return bindQAction(uintptr(h)) }
func UnsafeNewQActionOwned(h unsafe.Pointer) *QAction { return newownedQAction(uintptr(h)) }

func (this *QAction) UnsafeRelease() {
	this.h = 0
	if this.QObject != nil {
		this.QObject.UnsafeRelease()
	}
}

func NewQAction() *QAction {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QAction_new, uintptr(__goid))
	__o := newownedQAction(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QAction_virtIdx, QAction_enable)
	return __o
}

func NewQAction2(text string) *QAction {
	__goid := nextVirtID()
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QAction_new2, uintptr(__goid), text_p, uintptr(len(text)))
	__o := newownedQAction(__r)
	runtime.KeepAlive(text)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QAction_virtIdx, QAction_enable)
	return __o
}

func NewQAction3(icon *QIcon, text string) *QAction {
	__goid := nextVirtID()
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QAction_new3, uintptr(__goid), uintptr(icon.UnsafePointer()), text_p, uintptr(len(text)))
	__o := newownedQAction(__r)
	runtime.KeepAlive(icon)
	runtime.KeepAlive(text)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QAction_virtIdx, QAction_enable)
	return __o
}

func NewQAction4(parent *QObject) *QAction {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QAction_new4, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQAction(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QAction_virtIdx, QAction_enable)
	return __o
}

func NewQAction5(text string, parent *QObject) *QAction {
	__goid := nextVirtID()
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QAction_new5, uintptr(__goid), text_p, uintptr(len(text)), uintptr(parent.UnsafePointer()))
	__o := newownedQAction(__r)
	runtime.KeepAlive(text)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QAction_virtIdx, QAction_enable)
	return __o
}

func NewQAction6(icon *QIcon, text string, parent *QObject) *QAction {
	__goid := nextVirtID()
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QAction_new6, uintptr(__goid), uintptr(icon.UnsafePointer()), text_p, uintptr(len(text)), uintptr(parent.UnsafePointer()))
	__o := newownedQAction(__r)
	runtime.KeepAlive(icon)
	runtime.KeepAlive(text)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QAction_virtIdx, QAction_enable)
	return __o
}

func QAction_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAction_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QAction) AssociatedObjects() []*QObject {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAction_associatedObjects, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QObject, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = bindQObject(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QAction) SetActionGroup(group *QActionGroup) {
	purego.SyscallN(addr_QAction_setActionGroup, this.h, uintptr(group.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(group)
}

func (this *QAction) ActionGroup() *QActionGroup {
	__r, _, _ := purego.SyscallN(addr_QAction_actionGroup, this.h)
	runtime.KeepAlive(this)
	return bindQActionGroup(__r)
}

func (this *QAction) SetIcon(icon *QIcon) {
	purego.SyscallN(addr_QAction_setIcon, this.h, uintptr(icon.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(icon)
}

func (this *QAction) Icon() *QIcon {
	__r, _, _ := purego.SyscallN(addr_QAction_icon, this.h)
	runtime.KeepAlive(this)
	return newownedQIcon(__r)
}

func (this *QAction) SetText(text string) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QAction_setText, this.h, text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
}

func (this *QAction) Text() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAction_text, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QAction) SetIconText(text string) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QAction_setIconText, this.h, text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
}

func (this *QAction) IconText() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAction_iconText, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QAction) SetToolTip(tip string) {
	var tip_p uintptr
	if len(tip) > 0 {
		tip_p = uintptr(unsafe.Pointer(unsafe.StringData(tip)))
	}
	purego.SyscallN(addr_QAction_setToolTip, this.h, tip_p, uintptr(len(tip)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(tip)
}

func (this *QAction) ToolTip() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAction_toolTip, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QAction) SetStatusTip(statusTip string) {
	var statusTip_p uintptr
	if len(statusTip) > 0 {
		statusTip_p = uintptr(unsafe.Pointer(unsafe.StringData(statusTip)))
	}
	purego.SyscallN(addr_QAction_setStatusTip, this.h, statusTip_p, uintptr(len(statusTip)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(statusTip)
}

func (this *QAction) StatusTip() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAction_statusTip, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QAction) SetWhatsThis(what string) {
	var what_p uintptr
	if len(what) > 0 {
		what_p = uintptr(unsafe.Pointer(unsafe.StringData(what)))
	}
	purego.SyscallN(addr_QAction_setWhatsThis, this.h, what_p, uintptr(len(what)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(what)
}

func (this *QAction) WhatsThis() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAction_whatsThis, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QAction) SetPriority(priority QAction__Priority) {
	purego.SyscallN(addr_QAction_setPriority, this.h, uintptr(priority))
	runtime.KeepAlive(this)
}

func (this *QAction) Priority() QAction__Priority {
	__r0, _, _ := purego.SyscallN(addr_QAction_priority, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QAction__Priority(__r)
}

func (this *QAction) SetSeparator(b bool) {
	purego.SyscallN(addr_QAction_setSeparator, this.h, uintptr(b2i(b)))
	runtime.KeepAlive(this)
}

func (this *QAction) IsSeparator() bool {
	__r0, _, _ := purego.SyscallN(addr_QAction_isSeparator, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QAction) SetShortcut(shortcut *QKeySequence) {
	purego.SyscallN(addr_QAction_setShortcut, this.h, uintptr(shortcut.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(shortcut)
}

func (this *QAction) Shortcut() *QKeySequence {
	__r, _, _ := purego.SyscallN(addr_QAction_shortcut, this.h)
	runtime.KeepAlive(this)
	return newownedQKeySequence(__r)
}

func (this *QAction) SetShortcuts(shortcuts []*QKeySequence) {
	shortcuts_arr := make([]uintptr, len(shortcuts))
	for __i, __e := range shortcuts {
		shortcuts_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var shortcuts_p uintptr
	if len(shortcuts_arr) > 0 {
		shortcuts_p = uintptr(unsafe.Pointer(&shortcuts_arr[0]))
	}
	purego.SyscallN(addr_QAction_setShortcuts, this.h, shortcuts_p, uintptr(len(shortcuts)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(shortcuts)
	runtime.KeepAlive(shortcuts_arr)
}

func (this *QAction) SetShortcutsWithQKeySequenceStandardKey(arg1 QKeySequence__StandardKey) {
	purego.SyscallN(addr_QAction_setShortcuts2, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QAction) Shortcuts() []*QKeySequence {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAction_shortcuts, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QKeySequence, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQKeySequence(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QAction) SetShortcutContext(context ShortcutContext) {
	purego.SyscallN(addr_QAction_setShortcutContext, this.h, uintptr(context))
	runtime.KeepAlive(this)
}

func (this *QAction) ShortcutContext() ShortcutContext {
	__r0, _, _ := purego.SyscallN(addr_QAction_shortcutContext, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return ShortcutContext(__r)
}

func (this *QAction) SetAutoRepeat(arg1 bool) {
	purego.SyscallN(addr_QAction_setAutoRepeat, this.h, uintptr(b2i(arg1)))
	runtime.KeepAlive(this)
}

func (this *QAction) AutoRepeat() bool {
	__r0, _, _ := purego.SyscallN(addr_QAction_autoRepeat, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QAction) SetFont(font *QFont) {
	purego.SyscallN(addr_QAction_setFont, this.h, uintptr(font.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(font)
}

func (this *QAction) Font() *QFont {
	__r, _, _ := purego.SyscallN(addr_QAction_font, this.h)
	runtime.KeepAlive(this)
	return newownedQFont(__r)
}

func (this *QAction) SetCheckable(arg1 bool) {
	purego.SyscallN(addr_QAction_setCheckable, this.h, uintptr(b2i(arg1)))
	runtime.KeepAlive(this)
}

func (this *QAction) IsCheckable() bool {
	__r0, _, _ := purego.SyscallN(addr_QAction_isCheckable, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QAction) Data() *QVariant {
	__r, _, _ := purego.SyscallN(addr_QAction_data, this.h)
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QAction) SetData(var_ *QVariant) {
	purego.SyscallN(addr_QAction_setData, this.h, uintptr(var_.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(var_)
}

func (this *QAction) IsChecked() bool {
	__r0, _, _ := purego.SyscallN(addr_QAction_isChecked, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QAction) IsEnabled() bool {
	__r0, _, _ := purego.SyscallN(addr_QAction_isEnabled, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QAction) IsVisible() bool {
	__r0, _, _ := purego.SyscallN(addr_QAction_isVisible, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QAction) Activate(event QAction__ActionEvent) {
	purego.SyscallN(addr_QAction_activate, this.h, uintptr(event))
	runtime.KeepAlive(this)
}

func (this *QAction) SetMenuRole(menuRole QAction__MenuRole) {
	purego.SyscallN(addr_QAction_setMenuRole, this.h, uintptr(menuRole))
	runtime.KeepAlive(this)
}

func (this *QAction) MenuRole() QAction__MenuRole {
	__r0, _, _ := purego.SyscallN(addr_QAction_menuRole, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QAction__MenuRole(__r)
}

func (this *QAction) SetIconVisibleInMenu(visible bool) {
	purego.SyscallN(addr_QAction_setIconVisibleInMenu, this.h, uintptr(b2i(visible)))
	runtime.KeepAlive(this)
}

func (this *QAction) IsIconVisibleInMenu() bool {
	__r0, _, _ := purego.SyscallN(addr_QAction_isIconVisibleInMenu, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QAction) SetShortcutVisibleInContextMenu(show bool) {
	purego.SyscallN(addr_QAction_setShortcutVisibleInContextMenu, this.h, uintptr(b2i(show)))
	runtime.KeepAlive(this)
}

func (this *QAction) IsShortcutVisibleInContextMenu() bool {
	__r0, _, _ := purego.SyscallN(addr_QAction_isShortcutVisibleInContextMenu, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QAction) ShowStatusText() bool {
	__r0, _, _ := purego.SyscallN(addr_QAction_showStatusText, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QAction) Trigger() {
	purego.SyscallN(addr_QAction_trigger, this.h)
	runtime.KeepAlive(this)
}

func (this *QAction) Hover() {
	purego.SyscallN(addr_QAction_hover, this.h)
	runtime.KeepAlive(this)
}

func (this *QAction) SetChecked(arg1 bool) {
	purego.SyscallN(addr_QAction_setChecked, this.h, uintptr(b2i(arg1)))
	runtime.KeepAlive(this)
}

func (this *QAction) Toggle() {
	purego.SyscallN(addr_QAction_toggle, this.h)
	runtime.KeepAlive(this)
}

func (this *QAction) SetEnabled(arg1 bool) {
	purego.SyscallN(addr_QAction_setEnabled, this.h, uintptr(b2i(arg1)))
	runtime.KeepAlive(this)
}

func (this *QAction) ResetEnabled() {
	purego.SyscallN(addr_QAction_resetEnabled, this.h)
	runtime.KeepAlive(this)
}

func (this *QAction) SetDisabled(b bool) {
	purego.SyscallN(addr_QAction_setDisabled, this.h, uintptr(b2i(b)))
	runtime.KeepAlive(this)
}

func (this *QAction) SetVisible(arg1 bool) {
	purego.SyscallN(addr_QAction_setVisible, this.h, uintptr(b2i(arg1)))
	runtime.KeepAlive(this)
}

func (this *QAction) Changed() {
	purego.SyscallN(addr_QAction_changed, this.h)
	runtime.KeepAlive(this)
}

func (this *QAction) EnabledChanged(enabled bool) {
	purego.SyscallN(addr_QAction_enabledChanged, this.h, uintptr(b2i(enabled)))
	runtime.KeepAlive(this)
}

func (this *QAction) CheckableChanged(checkable bool) {
	purego.SyscallN(addr_QAction_checkableChanged, this.h, uintptr(b2i(checkable)))
	runtime.KeepAlive(this)
}

func (this *QAction) VisibleChanged() {
	purego.SyscallN(addr_QAction_visibleChanged, this.h)
	runtime.KeepAlive(this)
}

func (this *QAction) Triggered() {
	purego.SyscallN(addr_QAction_triggered, this.h)
	runtime.KeepAlive(this)
}

func (this *QAction) Hovered() {
	purego.SyscallN(addr_QAction_hovered, this.h)
	runtime.KeepAlive(this)
}

func (this *QAction) Toggled(arg1 bool) {
	purego.SyscallN(addr_QAction_toggled, this.h, uintptr(b2i(arg1)))
	runtime.KeepAlive(this)
}

func QAction_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QAction_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QAction_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QAction_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QAction) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QAction_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QAction) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QAction_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAction) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QAction_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QAction) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QAction_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QAction_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAction_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QAction_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAction_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func (this *QAction) ShowStatusTextWithObject(object *QObject) bool {
	__r0, _, _ := purego.SyscallN(addr_QAction_showStatusText2, this.h, uintptr(object.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(object)
	return __r != 0
}

func (this *QAction) TriggeredWithChecked(checked bool) {
	purego.SyscallN(addr_QAction_triggered2, this.h, uintptr(b2i(checked)))
	runtime.KeepAlive(this)
}

func QAction_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QAction_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QAction_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QAction_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QAction) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QAction_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QAction) EventDefault(arg1 *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QAction_event_default, this.h, uintptr(arg1.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return __r != 0
}

func (this *QAction) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QAction_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QAction) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QAction_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAction) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QAction_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAction) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QAction_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAction) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QAction_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QAction) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QAction_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QAction) OnChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QAction_connect_changed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAction) OnEnabledChanged(fn func(bool)) int64 {
	id := registerCallback(func(enabled int32) { fn(enabled != 0) })
	pg_QAction_connect_enabledChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAction) OnCheckableChanged(fn func(bool)) int64 {
	id := registerCallback(func(checkable int32) { fn(checkable != 0) })
	pg_QAction_connect_checkableChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAction) OnVisibleChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QAction_connect_visibleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAction) OnTriggered(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QAction_connect_triggered(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAction) OnTriggeredWithChecked(fn func(bool)) int64 {
	id := registerCallback(func(checked int32) { fn(checked != 0) })
	pg_QAction_connect_triggered2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAction) OnHovered(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QAction_connect_hovered(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAction) OnToggled(fn func(bool)) int64 {
	id := registerCallback(func(arg1 int32) { fn(arg1 != 0) })
	pg_QAction_connect_toggled(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAction) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QAction_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAction) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QAction_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAction) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QAction_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAction) Disconnect(handle int64) {
	pg_QAction__disconnect(handle)
	unregisterCallback(handle)
}

func QAction_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case 1237236456312824991:
		return 1
	case 5217623514630858297:
		return 2
	case -5409927581812374498:
		return 3
	case -1185315367526857833:
		return 4
	case 3769868229123868474:
		return 5
	case 7379029982170503008:
		return 6
	case -568369772982215962:
		return 7
	}
	return -1
}

func QAction_enable(h uintptr, idx int32) { pg_QAction__set_override(h, idx) }

func (this *QAction) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QAction) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(arg1 uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(arg1)))) }) // slot "event"
}

func (this *QAction) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QAction) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QAction) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QAction) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QAction) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QAction) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QAction) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QAction_delete(this.h)
	}
	this.UnsafeRelease()
}
