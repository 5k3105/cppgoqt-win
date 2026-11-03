package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QStyleHints_base func(uintptr) uintptr
	addr_QStyleHints_tr uintptr
	addr_QStyleHints_setMouseDoubleClickInterval uintptr
	addr_QStyleHints_mouseDoubleClickInterval uintptr
	addr_QStyleHints_mouseDoubleClickDistance uintptr
	addr_QStyleHints_touchDoubleTapDistance uintptr
	addr_QStyleHints_setMousePressAndHoldInterval uintptr
	addr_QStyleHints_mousePressAndHoldInterval uintptr
	addr_QStyleHints_setStartDragDistance uintptr
	addr_QStyleHints_startDragDistance uintptr
	addr_QStyleHints_setStartDragTime uintptr
	addr_QStyleHints_startDragTime uintptr
	addr_QStyleHints_startDragVelocity uintptr
	addr_QStyleHints_setKeyboardInputInterval uintptr
	addr_QStyleHints_keyboardInputInterval uintptr
	addr_QStyleHints_keyboardAutoRepeatRate uintptr
	pg_QStyleHints_keyboardAutoRepeatRateF func(uintptr) float64
	addr_QStyleHints_setCursorFlashTime uintptr
	addr_QStyleHints_cursorFlashTime uintptr
	addr_QStyleHints_showIsFullScreen uintptr
	addr_QStyleHints_showIsMaximized uintptr
	addr_QStyleHints_showShortcutsInContextMenus uintptr
	addr_QStyleHints_setShowShortcutsInContextMenus uintptr
	addr_QStyleHints_contextMenuTrigger uintptr
	addr_QStyleHints_setContextMenuTrigger uintptr
	addr_QStyleHints_menuSelectionWraps uintptr
	addr_QStyleHints_passwordMaskDelay uintptr
	addr_QStyleHints_passwordMaskCharacter uintptr
	pg_QStyleHints_fontSmoothingGamma func(uintptr) float64
	addr_QStyleHints_useRtlExtensions uintptr
	addr_QStyleHints_setFocusOnTouchRelease uintptr
	addr_QStyleHints_tabFocusBehavior uintptr
	addr_QStyleHints_setTabFocusBehavior uintptr
	addr_QStyleHints_singleClickActivation uintptr
	addr_QStyleHints_useHoverEffects uintptr
	addr_QStyleHints_setUseHoverEffects uintptr
	addr_QStyleHints_wheelScrollLines uintptr
	addr_QStyleHints_setWheelScrollLines uintptr
	addr_QStyleHints_setMouseQuickSelectionThreshold uintptr
	addr_QStyleHints_mouseQuickSelectionThreshold uintptr
	addr_QStyleHints_colorScheme uintptr
	addr_QStyleHints_setColorScheme uintptr
	addr_QStyleHints_unsetColorScheme uintptr
	addr_QStyleHints_accessibility uintptr
	addr_QStyleHints_cursorFlashTimeChanged uintptr
	addr_QStyleHints_keyboardInputIntervalChanged uintptr
	addr_QStyleHints_mouseDoubleClickIntervalChanged uintptr
	addr_QStyleHints_mousePressAndHoldIntervalChanged uintptr
	addr_QStyleHints_startDragDistanceChanged uintptr
	addr_QStyleHints_startDragTimeChanged uintptr
	addr_QStyleHints_tabFocusBehaviorChanged uintptr
	addr_QStyleHints_useHoverEffectsChanged uintptr
	addr_QStyleHints_showShortcutsInContextMenusChanged uintptr
	addr_QStyleHints_contextMenuTriggerChanged uintptr
	addr_QStyleHints_wheelScrollLinesChanged uintptr
	addr_QStyleHints_mouseQuickSelectionThresholdChanged uintptr
	addr_QStyleHints_colorSchemeChanged uintptr
	addr_QStyleHints_connect uintptr
	addr_QStyleHints_connect2 uintptr
	addr_QStyleHints_tr2 uintptr
	addr_QStyleHints_tr3 uintptr
	addr_QStyleHints_connect3 uintptr
	addr_QStyleHints_connect4 uintptr
	pg_QStyleHints_connect_cursorFlashTimeChanged func(uintptr, int64)
	pg_QStyleHints_connect_keyboardInputIntervalChanged func(uintptr, int64)
	pg_QStyleHints_connect_mouseDoubleClickIntervalChanged func(uintptr, int64)
	pg_QStyleHints_connect_mousePressAndHoldIntervalChanged func(uintptr, int64)
	pg_QStyleHints_connect_startDragDistanceChanged func(uintptr, int64)
	pg_QStyleHints_connect_startDragTimeChanged func(uintptr, int64)
	pg_QStyleHints_connect_tabFocusBehaviorChanged func(uintptr, int64)
	pg_QStyleHints_connect_useHoverEffectsChanged func(uintptr, int64)
	pg_QStyleHints_connect_showShortcutsInContextMenusChanged func(uintptr, int64)
	pg_QStyleHints_connect_contextMenuTriggerChanged func(uintptr, int64)
	pg_QStyleHints_connect_wheelScrollLinesChanged func(uintptr, int64)
	pg_QStyleHints_connect_mouseQuickSelectionThresholdChanged func(uintptr, int64)
	pg_QStyleHints_connect_colorSchemeChanged func(uintptr, int64)
	pg_QStyleHints_connect_destroyed func(uintptr, int64)
	pg_QStyleHints_connect_destroyed2 func(uintptr, int64)
	pg_QStyleHints_connect_objectNameChanged func(uintptr, int64)
	pg_QStyleHints__disconnect func(int64)
	pg_QStyleHints_delete func(uintptr)
)

func init() {
	reg(&pg_QStyleHints_base, "QStyleHints_base")
	symaddr(&addr_QStyleHints_tr, "QStyleHints_tr")
	symaddr(&addr_QStyleHints_setMouseDoubleClickInterval, "QStyleHints_setMouseDoubleClickInterval")
	symaddr(&addr_QStyleHints_mouseDoubleClickInterval, "QStyleHints_mouseDoubleClickInterval")
	symaddr(&addr_QStyleHints_mouseDoubleClickDistance, "QStyleHints_mouseDoubleClickDistance")
	symaddr(&addr_QStyleHints_touchDoubleTapDistance, "QStyleHints_touchDoubleTapDistance")
	symaddr(&addr_QStyleHints_setMousePressAndHoldInterval, "QStyleHints_setMousePressAndHoldInterval")
	symaddr(&addr_QStyleHints_mousePressAndHoldInterval, "QStyleHints_mousePressAndHoldInterval")
	symaddr(&addr_QStyleHints_setStartDragDistance, "QStyleHints_setStartDragDistance")
	symaddr(&addr_QStyleHints_startDragDistance, "QStyleHints_startDragDistance")
	symaddr(&addr_QStyleHints_setStartDragTime, "QStyleHints_setStartDragTime")
	symaddr(&addr_QStyleHints_startDragTime, "QStyleHints_startDragTime")
	symaddr(&addr_QStyleHints_startDragVelocity, "QStyleHints_startDragVelocity")
	symaddr(&addr_QStyleHints_setKeyboardInputInterval, "QStyleHints_setKeyboardInputInterval")
	symaddr(&addr_QStyleHints_keyboardInputInterval, "QStyleHints_keyboardInputInterval")
	symaddr(&addr_QStyleHints_keyboardAutoRepeatRate, "QStyleHints_keyboardAutoRepeatRate")
	reg(&pg_QStyleHints_keyboardAutoRepeatRateF, "QStyleHints_keyboardAutoRepeatRateF")
	symaddr(&addr_QStyleHints_setCursorFlashTime, "QStyleHints_setCursorFlashTime")
	symaddr(&addr_QStyleHints_cursorFlashTime, "QStyleHints_cursorFlashTime")
	symaddr(&addr_QStyleHints_showIsFullScreen, "QStyleHints_showIsFullScreen")
	symaddr(&addr_QStyleHints_showIsMaximized, "QStyleHints_showIsMaximized")
	symaddr(&addr_QStyleHints_showShortcutsInContextMenus, "QStyleHints_showShortcutsInContextMenus")
	symaddr(&addr_QStyleHints_setShowShortcutsInContextMenus, "QStyleHints_setShowShortcutsInContextMenus")
	symaddr(&addr_QStyleHints_contextMenuTrigger, "QStyleHints_contextMenuTrigger")
	symaddr(&addr_QStyleHints_setContextMenuTrigger, "QStyleHints_setContextMenuTrigger")
	symaddr(&addr_QStyleHints_menuSelectionWraps, "QStyleHints_menuSelectionWraps")
	symaddr(&addr_QStyleHints_passwordMaskDelay, "QStyleHints_passwordMaskDelay")
	symaddr(&addr_QStyleHints_passwordMaskCharacter, "QStyleHints_passwordMaskCharacter")
	reg(&pg_QStyleHints_fontSmoothingGamma, "QStyleHints_fontSmoothingGamma")
	symaddr(&addr_QStyleHints_useRtlExtensions, "QStyleHints_useRtlExtensions")
	symaddr(&addr_QStyleHints_setFocusOnTouchRelease, "QStyleHints_setFocusOnTouchRelease")
	symaddr(&addr_QStyleHints_tabFocusBehavior, "QStyleHints_tabFocusBehavior")
	symaddr(&addr_QStyleHints_setTabFocusBehavior, "QStyleHints_setTabFocusBehavior")
	symaddr(&addr_QStyleHints_singleClickActivation, "QStyleHints_singleClickActivation")
	symaddr(&addr_QStyleHints_useHoverEffects, "QStyleHints_useHoverEffects")
	symaddr(&addr_QStyleHints_setUseHoverEffects, "QStyleHints_setUseHoverEffects")
	symaddr(&addr_QStyleHints_wheelScrollLines, "QStyleHints_wheelScrollLines")
	symaddr(&addr_QStyleHints_setWheelScrollLines, "QStyleHints_setWheelScrollLines")
	symaddr(&addr_QStyleHints_setMouseQuickSelectionThreshold, "QStyleHints_setMouseQuickSelectionThreshold")
	symaddr(&addr_QStyleHints_mouseQuickSelectionThreshold, "QStyleHints_mouseQuickSelectionThreshold")
	symaddr(&addr_QStyleHints_colorScheme, "QStyleHints_colorScheme")
	symaddr(&addr_QStyleHints_setColorScheme, "QStyleHints_setColorScheme")
	symaddr(&addr_QStyleHints_unsetColorScheme, "QStyleHints_unsetColorScheme")
	symaddr(&addr_QStyleHints_accessibility, "QStyleHints_accessibility")
	symaddr(&addr_QStyleHints_cursorFlashTimeChanged, "QStyleHints_cursorFlashTimeChanged")
	symaddr(&addr_QStyleHints_keyboardInputIntervalChanged, "QStyleHints_keyboardInputIntervalChanged")
	symaddr(&addr_QStyleHints_mouseDoubleClickIntervalChanged, "QStyleHints_mouseDoubleClickIntervalChanged")
	symaddr(&addr_QStyleHints_mousePressAndHoldIntervalChanged, "QStyleHints_mousePressAndHoldIntervalChanged")
	symaddr(&addr_QStyleHints_startDragDistanceChanged, "QStyleHints_startDragDistanceChanged")
	symaddr(&addr_QStyleHints_startDragTimeChanged, "QStyleHints_startDragTimeChanged")
	symaddr(&addr_QStyleHints_tabFocusBehaviorChanged, "QStyleHints_tabFocusBehaviorChanged")
	symaddr(&addr_QStyleHints_useHoverEffectsChanged, "QStyleHints_useHoverEffectsChanged")
	symaddr(&addr_QStyleHints_showShortcutsInContextMenusChanged, "QStyleHints_showShortcutsInContextMenusChanged")
	symaddr(&addr_QStyleHints_contextMenuTriggerChanged, "QStyleHints_contextMenuTriggerChanged")
	symaddr(&addr_QStyleHints_wheelScrollLinesChanged, "QStyleHints_wheelScrollLinesChanged")
	symaddr(&addr_QStyleHints_mouseQuickSelectionThresholdChanged, "QStyleHints_mouseQuickSelectionThresholdChanged")
	symaddr(&addr_QStyleHints_colorSchemeChanged, "QStyleHints_colorSchemeChanged")
	symaddr(&addr_QStyleHints_connect, "QStyleHints_connect")
	symaddr(&addr_QStyleHints_connect2, "QStyleHints_connect2")
	symaddr(&addr_QStyleHints_tr2, "QStyleHints_tr2")
	symaddr(&addr_QStyleHints_tr3, "QStyleHints_tr3")
	symaddr(&addr_QStyleHints_connect3, "QStyleHints_connect3")
	symaddr(&addr_QStyleHints_connect4, "QStyleHints_connect4")
	reg(&pg_QStyleHints_connect_cursorFlashTimeChanged, "QStyleHints_connect_cursorFlashTimeChanged")
	reg(&pg_QStyleHints_connect_keyboardInputIntervalChanged, "QStyleHints_connect_keyboardInputIntervalChanged")
	reg(&pg_QStyleHints_connect_mouseDoubleClickIntervalChanged, "QStyleHints_connect_mouseDoubleClickIntervalChanged")
	reg(&pg_QStyleHints_connect_mousePressAndHoldIntervalChanged, "QStyleHints_connect_mousePressAndHoldIntervalChanged")
	reg(&pg_QStyleHints_connect_startDragDistanceChanged, "QStyleHints_connect_startDragDistanceChanged")
	reg(&pg_QStyleHints_connect_startDragTimeChanged, "QStyleHints_connect_startDragTimeChanged")
	reg(&pg_QStyleHints_connect_tabFocusBehaviorChanged, "QStyleHints_connect_tabFocusBehaviorChanged")
	reg(&pg_QStyleHints_connect_useHoverEffectsChanged, "QStyleHints_connect_useHoverEffectsChanged")
	reg(&pg_QStyleHints_connect_showShortcutsInContextMenusChanged, "QStyleHints_connect_showShortcutsInContextMenusChanged")
	reg(&pg_QStyleHints_connect_contextMenuTriggerChanged, "QStyleHints_connect_contextMenuTriggerChanged")
	reg(&pg_QStyleHints_connect_wheelScrollLinesChanged, "QStyleHints_connect_wheelScrollLinesChanged")
	reg(&pg_QStyleHints_connect_mouseQuickSelectionThresholdChanged, "QStyleHints_connect_mouseQuickSelectionThresholdChanged")
	reg(&pg_QStyleHints_connect_colorSchemeChanged, "QStyleHints_connect_colorSchemeChanged")
	reg(&pg_QStyleHints_connect_destroyed, "QStyleHints_connect_destroyed")
	reg(&pg_QStyleHints_connect_destroyed2, "QStyleHints_connect_destroyed2")
	reg(&pg_QStyleHints_connect_objectNameChanged, "QStyleHints_connect_objectNameChanged")
	reg(&pg_QStyleHints__disconnect, "QStyleHints__disconnect")
	reg(&pg_QStyleHints_delete, "QStyleHints_delete")
}

type QStyleHints struct {
	h uintptr
	*QObject
}

func bindQStyleHints(h uintptr) *QStyleHints {
	if h == 0 {
		return nil
	}
	return &QStyleHints{h: h, QObject: bindQObject(pg_QStyleHints_base(h))}
}

func newownedQStyleHints(h uintptr) *QStyleHints { return bindQStyleHints(h) }

func (this *QStyleHints) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QStyleHints) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQStyleHints(h unsafe.Pointer) *QStyleHints { return bindQStyleHints(uintptr(h)) }
func UnsafeNewQStyleHintsOwned(h unsafe.Pointer) *QStyleHints { return newownedQStyleHints(uintptr(h)) }

func (this *QStyleHints) UnsafeRelease() {
	this.h = 0
	if this.QObject != nil {
		this.QObject.UnsafeRelease()
	}
}

func QStyleHints_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStyleHints_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QStyleHints) SetMouseDoubleClickInterval(mouseDoubleClickInterval int) {
	purego.SyscallN(addr_QStyleHints_setMouseDoubleClickInterval, this.h, uintptr(mouseDoubleClickInterval))
	runtime.KeepAlive(this)
}

func (this *QStyleHints) MouseDoubleClickInterval() int {
	__r0, _, _ := purego.SyscallN(addr_QStyleHints_mouseDoubleClickInterval, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QStyleHints) MouseDoubleClickDistance() int {
	__r0, _, _ := purego.SyscallN(addr_QStyleHints_mouseDoubleClickDistance, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QStyleHints) TouchDoubleTapDistance() int {
	__r0, _, _ := purego.SyscallN(addr_QStyleHints_touchDoubleTapDistance, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QStyleHints) SetMousePressAndHoldInterval(mousePressAndHoldInterval int) {
	purego.SyscallN(addr_QStyleHints_setMousePressAndHoldInterval, this.h, uintptr(mousePressAndHoldInterval))
	runtime.KeepAlive(this)
}

func (this *QStyleHints) MousePressAndHoldInterval() int {
	__r0, _, _ := purego.SyscallN(addr_QStyleHints_mousePressAndHoldInterval, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QStyleHints) SetStartDragDistance(startDragDistance int) {
	purego.SyscallN(addr_QStyleHints_setStartDragDistance, this.h, uintptr(startDragDistance))
	runtime.KeepAlive(this)
}

func (this *QStyleHints) StartDragDistance() int {
	__r0, _, _ := purego.SyscallN(addr_QStyleHints_startDragDistance, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QStyleHints) SetStartDragTime(startDragTime int) {
	purego.SyscallN(addr_QStyleHints_setStartDragTime, this.h, uintptr(startDragTime))
	runtime.KeepAlive(this)
}

func (this *QStyleHints) StartDragTime() int {
	__r0, _, _ := purego.SyscallN(addr_QStyleHints_startDragTime, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QStyleHints) StartDragVelocity() int {
	__r0, _, _ := purego.SyscallN(addr_QStyleHints_startDragVelocity, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QStyleHints) SetKeyboardInputInterval(keyboardInputInterval int) {
	purego.SyscallN(addr_QStyleHints_setKeyboardInputInterval, this.h, uintptr(keyboardInputInterval))
	runtime.KeepAlive(this)
}

func (this *QStyleHints) KeyboardInputInterval() int {
	__r0, _, _ := purego.SyscallN(addr_QStyleHints_keyboardInputInterval, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QStyleHints) KeyboardAutoRepeatRate() int {
	__r0, _, _ := purego.SyscallN(addr_QStyleHints_keyboardAutoRepeatRate, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QStyleHints) KeyboardAutoRepeatRateF() float64 {
	__r := pg_QStyleHints_keyboardAutoRepeatRateF(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QStyleHints) SetCursorFlashTime(cursorFlashTime int) {
	purego.SyscallN(addr_QStyleHints_setCursorFlashTime, this.h, uintptr(cursorFlashTime))
	runtime.KeepAlive(this)
}

func (this *QStyleHints) CursorFlashTime() int {
	__r0, _, _ := purego.SyscallN(addr_QStyleHints_cursorFlashTime, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QStyleHints) ShowIsFullScreen() bool {
	__r0, _, _ := purego.SyscallN(addr_QStyleHints_showIsFullScreen, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QStyleHints) ShowIsMaximized() bool {
	__r0, _, _ := purego.SyscallN(addr_QStyleHints_showIsMaximized, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QStyleHints) ShowShortcutsInContextMenus() bool {
	__r0, _, _ := purego.SyscallN(addr_QStyleHints_showShortcutsInContextMenus, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QStyleHints) SetShowShortcutsInContextMenus(showShortcutsInContextMenus bool) {
	purego.SyscallN(addr_QStyleHints_setShowShortcutsInContextMenus, this.h, uintptr(b2i(showShortcutsInContextMenus)))
	runtime.KeepAlive(this)
}

func (this *QStyleHints) ContextMenuTrigger() ContextMenuTrigger {
	__r0, _, _ := purego.SyscallN(addr_QStyleHints_contextMenuTrigger, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return ContextMenuTrigger(__r)
}

func (this *QStyleHints) SetContextMenuTrigger(contextMenuTrigger ContextMenuTrigger) {
	purego.SyscallN(addr_QStyleHints_setContextMenuTrigger, this.h, uintptr(contextMenuTrigger))
	runtime.KeepAlive(this)
}

func (this *QStyleHints) MenuSelectionWraps() bool {
	__r0, _, _ := purego.SyscallN(addr_QStyleHints_menuSelectionWraps, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QStyleHints) PasswordMaskDelay() int {
	__r0, _, _ := purego.SyscallN(addr_QStyleHints_passwordMaskDelay, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QStyleHints) PasswordMaskCharacter() *QChar {
	__r, _, _ := purego.SyscallN(addr_QStyleHints_passwordMaskCharacter, this.h)
	runtime.KeepAlive(this)
	return newownedQChar(__r)
}

func (this *QStyleHints) FontSmoothingGamma() float64 {
	__r := pg_QStyleHints_fontSmoothingGamma(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QStyleHints) UseRtlExtensions() bool {
	__r0, _, _ := purego.SyscallN(addr_QStyleHints_useRtlExtensions, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QStyleHints) SetFocusOnTouchRelease() bool {
	__r0, _, _ := purego.SyscallN(addr_QStyleHints_setFocusOnTouchRelease, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QStyleHints) TabFocusBehavior() TabFocusBehavior {
	__r0, _, _ := purego.SyscallN(addr_QStyleHints_tabFocusBehavior, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return TabFocusBehavior(__r)
}

func (this *QStyleHints) SetTabFocusBehavior(tabFocusBehavior TabFocusBehavior) {
	purego.SyscallN(addr_QStyleHints_setTabFocusBehavior, this.h, uintptr(tabFocusBehavior))
	runtime.KeepAlive(this)
}

func (this *QStyleHints) SingleClickActivation() bool {
	__r0, _, _ := purego.SyscallN(addr_QStyleHints_singleClickActivation, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QStyleHints) UseHoverEffects() bool {
	__r0, _, _ := purego.SyscallN(addr_QStyleHints_useHoverEffects, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QStyleHints) SetUseHoverEffects(useHoverEffects bool) {
	purego.SyscallN(addr_QStyleHints_setUseHoverEffects, this.h, uintptr(b2i(useHoverEffects)))
	runtime.KeepAlive(this)
}

func (this *QStyleHints) WheelScrollLines() int {
	__r0, _, _ := purego.SyscallN(addr_QStyleHints_wheelScrollLines, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QStyleHints) SetWheelScrollLines(scrollLines int) {
	purego.SyscallN(addr_QStyleHints_setWheelScrollLines, this.h, uintptr(scrollLines))
	runtime.KeepAlive(this)
}

func (this *QStyleHints) SetMouseQuickSelectionThreshold(threshold int) {
	purego.SyscallN(addr_QStyleHints_setMouseQuickSelectionThreshold, this.h, uintptr(threshold))
	runtime.KeepAlive(this)
}

func (this *QStyleHints) MouseQuickSelectionThreshold() int {
	__r0, _, _ := purego.SyscallN(addr_QStyleHints_mouseQuickSelectionThreshold, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QStyleHints) ColorScheme() ColorScheme {
	__r0, _, _ := purego.SyscallN(addr_QStyleHints_colorScheme, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return ColorScheme(__r)
}

func (this *QStyleHints) SetColorScheme(scheme ColorScheme) {
	purego.SyscallN(addr_QStyleHints_setColorScheme, this.h, uintptr(scheme))
	runtime.KeepAlive(this)
}

func (this *QStyleHints) UnsetColorScheme() {
	purego.SyscallN(addr_QStyleHints_unsetColorScheme, this.h)
	runtime.KeepAlive(this)
}

func (this *QStyleHints) Accessibility() *QAccessibilityHints {
	__r, _, _ := purego.SyscallN(addr_QStyleHints_accessibility, this.h)
	runtime.KeepAlive(this)
	return bindQAccessibilityHints(__r)
}

func (this *QStyleHints) CursorFlashTimeChanged(cursorFlashTime int) {
	purego.SyscallN(addr_QStyleHints_cursorFlashTimeChanged, this.h, uintptr(cursorFlashTime))
	runtime.KeepAlive(this)
}

func (this *QStyleHints) KeyboardInputIntervalChanged(keyboardInputInterval int) {
	purego.SyscallN(addr_QStyleHints_keyboardInputIntervalChanged, this.h, uintptr(keyboardInputInterval))
	runtime.KeepAlive(this)
}

func (this *QStyleHints) MouseDoubleClickIntervalChanged(mouseDoubleClickInterval int) {
	purego.SyscallN(addr_QStyleHints_mouseDoubleClickIntervalChanged, this.h, uintptr(mouseDoubleClickInterval))
	runtime.KeepAlive(this)
}

func (this *QStyleHints) MousePressAndHoldIntervalChanged(mousePressAndHoldInterval int) {
	purego.SyscallN(addr_QStyleHints_mousePressAndHoldIntervalChanged, this.h, uintptr(mousePressAndHoldInterval))
	runtime.KeepAlive(this)
}

func (this *QStyleHints) StartDragDistanceChanged(startDragDistance int) {
	purego.SyscallN(addr_QStyleHints_startDragDistanceChanged, this.h, uintptr(startDragDistance))
	runtime.KeepAlive(this)
}

func (this *QStyleHints) StartDragTimeChanged(startDragTime int) {
	purego.SyscallN(addr_QStyleHints_startDragTimeChanged, this.h, uintptr(startDragTime))
	runtime.KeepAlive(this)
}

func (this *QStyleHints) TabFocusBehaviorChanged(tabFocusBehavior TabFocusBehavior) {
	purego.SyscallN(addr_QStyleHints_tabFocusBehaviorChanged, this.h, uintptr(tabFocusBehavior))
	runtime.KeepAlive(this)
}

func (this *QStyleHints) UseHoverEffectsChanged(useHoverEffects bool) {
	purego.SyscallN(addr_QStyleHints_useHoverEffectsChanged, this.h, uintptr(b2i(useHoverEffects)))
	runtime.KeepAlive(this)
}

func (this *QStyleHints) ShowShortcutsInContextMenusChanged(arg1 bool) {
	purego.SyscallN(addr_QStyleHints_showShortcutsInContextMenusChanged, this.h, uintptr(b2i(arg1)))
	runtime.KeepAlive(this)
}

func (this *QStyleHints) ContextMenuTriggerChanged(contextMenuTrigger ContextMenuTrigger) {
	purego.SyscallN(addr_QStyleHints_contextMenuTriggerChanged, this.h, uintptr(contextMenuTrigger))
	runtime.KeepAlive(this)
}

func (this *QStyleHints) WheelScrollLinesChanged(scrollLines int) {
	purego.SyscallN(addr_QStyleHints_wheelScrollLinesChanged, this.h, uintptr(scrollLines))
	runtime.KeepAlive(this)
}

func (this *QStyleHints) MouseQuickSelectionThresholdChanged(threshold int) {
	purego.SyscallN(addr_QStyleHints_mouseQuickSelectionThresholdChanged, this.h, uintptr(threshold))
	runtime.KeepAlive(this)
}

func (this *QStyleHints) ColorSchemeChanged(colorScheme ColorScheme) {
	purego.SyscallN(addr_QStyleHints_colorSchemeChanged, this.h, uintptr(colorScheme))
	runtime.KeepAlive(this)
}

func QStyleHints_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QStyleHints_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QStyleHints_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QStyleHints_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func QStyleHints_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStyleHints_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QStyleHints_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStyleHints_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QStyleHints_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QStyleHints_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QStyleHints_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QStyleHints_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QStyleHints) OnCursorFlashTimeChanged(fn func(int)) int64 {
	id := registerCallback(func(cursorFlashTime int32) { fn(int(cursorFlashTime)) })
	pg_QStyleHints_connect_cursorFlashTimeChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStyleHints) OnKeyboardInputIntervalChanged(fn func(int)) int64 {
	id := registerCallback(func(keyboardInputInterval int32) { fn(int(keyboardInputInterval)) })
	pg_QStyleHints_connect_keyboardInputIntervalChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStyleHints) OnMouseDoubleClickIntervalChanged(fn func(int)) int64 {
	id := registerCallback(func(mouseDoubleClickInterval int32) { fn(int(mouseDoubleClickInterval)) })
	pg_QStyleHints_connect_mouseDoubleClickIntervalChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStyleHints) OnMousePressAndHoldIntervalChanged(fn func(int)) int64 {
	id := registerCallback(func(mousePressAndHoldInterval int32) { fn(int(mousePressAndHoldInterval)) })
	pg_QStyleHints_connect_mousePressAndHoldIntervalChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStyleHints) OnStartDragDistanceChanged(fn func(int)) int64 {
	id := registerCallback(func(startDragDistance int32) { fn(int(startDragDistance)) })
	pg_QStyleHints_connect_startDragDistanceChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStyleHints) OnStartDragTimeChanged(fn func(int)) int64 {
	id := registerCallback(func(startDragTime int32) { fn(int(startDragTime)) })
	pg_QStyleHints_connect_startDragTimeChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStyleHints) OnTabFocusBehaviorChanged(fn func(TabFocusBehavior)) int64 {
	id := registerCallback(func(tabFocusBehavior int32) { fn(TabFocusBehavior(tabFocusBehavior)) })
	pg_QStyleHints_connect_tabFocusBehaviorChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStyleHints) OnUseHoverEffectsChanged(fn func(bool)) int64 {
	id := registerCallback(func(useHoverEffects int32) { fn(useHoverEffects != 0) })
	pg_QStyleHints_connect_useHoverEffectsChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStyleHints) OnShowShortcutsInContextMenusChanged(fn func(bool)) int64 {
	id := registerCallback(func(arg1 int32) { fn(arg1 != 0) })
	pg_QStyleHints_connect_showShortcutsInContextMenusChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStyleHints) OnContextMenuTriggerChanged(fn func(ContextMenuTrigger)) int64 {
	id := registerCallback(func(contextMenuTrigger int32) { fn(ContextMenuTrigger(contextMenuTrigger)) })
	pg_QStyleHints_connect_contextMenuTriggerChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStyleHints) OnWheelScrollLinesChanged(fn func(int)) int64 {
	id := registerCallback(func(scrollLines int32) { fn(int(scrollLines)) })
	pg_QStyleHints_connect_wheelScrollLinesChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStyleHints) OnMouseQuickSelectionThresholdChanged(fn func(int)) int64 {
	id := registerCallback(func(threshold int32) { fn(int(threshold)) })
	pg_QStyleHints_connect_mouseQuickSelectionThresholdChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStyleHints) OnColorSchemeChanged(fn func(ColorScheme)) int64 {
	id := registerCallback(func(colorScheme int32) { fn(ColorScheme(colorScheme)) })
	pg_QStyleHints_connect_colorSchemeChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStyleHints) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QStyleHints_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStyleHints) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QStyleHints_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStyleHints) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QStyleHints_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStyleHints) Disconnect(handle int64) {
	pg_QStyleHints__disconnect(handle)
	unregisterCallback(handle)
}

func (this *QStyleHints) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QStyleHints_delete(this.h)
	}
	this.UnsafeRelease()
}
