package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QApplication_chain uintptr
	addr_QApplication_tr uintptr
	addr_QApplication_style uintptr
	addr_QApplication_setStyle uintptr
	addr_QApplication_setStyle2 uintptr
	addr_QApplication_palette uintptr
	addr_QApplication_palette2 uintptr
	addr_QApplication_setPalette uintptr
	addr_QApplication_font uintptr
	addr_QApplication_font2 uintptr
	addr_QApplication_font3 uintptr
	addr_QApplication_setFont uintptr
	addr_QApplication_fontMetrics uintptr
	addr_QApplication_allWidgets uintptr
	addr_QApplication_topLevelWidgets uintptr
	addr_QApplication_activePopupWidget uintptr
	addr_QApplication_activeModalWidget uintptr
	addr_QApplication_focusWidget uintptr
	addr_QApplication_activeWindow uintptr
	addr_QApplication_setActiveWindow uintptr
	addr_QApplication_widgetAt uintptr
	addr_QApplication_widgetAt2 uintptr
	addr_QApplication_topLevelAt uintptr
	addr_QApplication_topLevelAt2 uintptr
	addr_QApplication_beep uintptr
	addr_QApplication_alert uintptr
	addr_QApplication_setCursorFlashTime uintptr
	addr_QApplication_cursorFlashTime uintptr
	addr_QApplication_setDoubleClickInterval uintptr
	addr_QApplication_doubleClickInterval uintptr
	addr_QApplication_setKeyboardInputInterval uintptr
	addr_QApplication_keyboardInputInterval uintptr
	addr_QApplication_setWheelScrollLines uintptr
	addr_QApplication_wheelScrollLines uintptr
	addr_QApplication_setStartDragTime uintptr
	addr_QApplication_startDragTime uintptr
	addr_QApplication_setStartDragDistance uintptr
	addr_QApplication_startDragDistance uintptr
	addr_QApplication_isEffectEnabled uintptr
	addr_QApplication_setEffectEnabled uintptr
	addr_QApplication_exec uintptr
	addr_QApplication_focusChanged uintptr
	addr_QApplication_styleSheet uintptr
	addr_QApplication_autoSipEnabled uintptr
	addr_QApplication_setStyleSheet uintptr
	addr_QApplication_setAutoSipEnabled uintptr
	addr_QApplication_closeAllWindows uintptr
	addr_QApplication_aboutQt uintptr
	addr_QApplication_setApplicationDisplayName uintptr
	addr_QApplication_applicationDisplayName uintptr
	addr_QApplication_setDesktopFileName uintptr
	addr_QApplication_desktopFileName uintptr
	addr_QApplication_allWindows uintptr
	addr_QApplication_topLevelWindows uintptr
	addr_QApplication_setWindowIcon uintptr
	addr_QApplication_windowIcon uintptr
	addr_QApplication_platformName uintptr
	addr_QApplication_modalWindow uintptr
	addr_QApplication_focusWindow uintptr
	addr_QApplication_focusObject uintptr
	addr_QApplication_primaryScreen uintptr
	addr_QApplication_screens uintptr
	addr_QApplication_screenAt uintptr
	addr_QApplication_overrideCursor uintptr
	addr_QApplication_setOverrideCursor uintptr
	addr_QApplication_changeOverrideCursor uintptr
	addr_QApplication_restoreOverrideCursor uintptr
	addr_QApplication_clipboard uintptr
	addr_QApplication_keyboardModifiers uintptr
	addr_QApplication_queryKeyboardModifiers uintptr
	addr_QApplication_mouseButtons uintptr
	addr_QApplication_setLayoutDirection uintptr
	addr_QApplication_layoutDirection uintptr
	addr_QApplication_isRightToLeft uintptr
	addr_QApplication_isLeftToRight uintptr
	addr_QApplication_styleHints uintptr
	addr_QApplication_setDesktopSettingsAware uintptr
	addr_QApplication_desktopSettingsAware uintptr
	addr_QApplication_inputMethod uintptr
	addr_QApplication_setQuitOnLastWindowClosed uintptr
	addr_QApplication_quitOnLastWindowClosed uintptr
	addr_QApplication_applicationState uintptr
	addr_QApplication_setHighDpiScaleFactorRoundingPolicy uintptr
	addr_QApplication_highDpiScaleFactorRoundingPolicy uintptr
	addr_QApplication_sync uintptr
	addr_QApplication_arguments uintptr
	addr_QApplication_setAttribute uintptr
	addr_QApplication_testAttribute uintptr
	addr_QApplication_setOrganizationDomain uintptr
	addr_QApplication_organizationDomain uintptr
	addr_QApplication_setOrganizationName uintptr
	addr_QApplication_organizationName uintptr
	addr_QApplication_setApplicationName uintptr
	addr_QApplication_applicationName uintptr
	addr_QApplication_setApplicationVersion uintptr
	addr_QApplication_applicationVersion uintptr
	addr_QApplication_setSetuidAllowed uintptr
	addr_QApplication_isSetuidAllowed uintptr
	addr_QApplication_instance uintptr
	addr_QApplication_instanceExists uintptr
	addr_QApplication_processEvents uintptr
	addr_QApplication_processEvents2 uintptr
	addr_QApplication_processEvents3 uintptr
	addr_QApplication_sendEvent uintptr
	addr_QApplication_postEvent uintptr
	addr_QApplication_sendPostedEvents uintptr
	addr_QApplication_removePostedEvents uintptr
	addr_QApplication_eventDispatcher uintptr
	addr_QApplication_setEventDispatcher uintptr
	addr_QApplication_startingUp uintptr
	addr_QApplication_closingDown uintptr
	addr_QApplication_applicationDirPath uintptr
	addr_QApplication_applicationFilePath uintptr
	addr_QApplication_applicationPid uintptr
	addr_QApplication_setLibraryPaths uintptr
	addr_QApplication_libraryPaths uintptr
	addr_QApplication_addLibraryPath uintptr
	addr_QApplication_removeLibraryPath uintptr
	addr_QApplication_installTranslator uintptr
	addr_QApplication_removeTranslator uintptr
	addr_QApplication_translate uintptr
	addr_QApplication_isQuitLockEnabled uintptr
	addr_QApplication_setQuitLockEnabled uintptr
	addr_QApplication_quit uintptr
	addr_QApplication_exit uintptr
	addr_QApplication_connect uintptr
	addr_QApplication_connect2 uintptr
	addr_QApplication_tr2 uintptr
	addr_QApplication_tr3 uintptr
	addr_QApplication_setPalette2 uintptr
	addr_QApplication_setFont2 uintptr
	addr_QApplication_alert2 uintptr
	addr_QApplication_setEffectEnabled2 uintptr
	addr_QApplication_setAttribute2 uintptr
	addr_QApplication_processEvents4 uintptr
	addr_QApplication_postEvent2 uintptr
	addr_QApplication_sendPostedEvents2 uintptr
	addr_QApplication_sendPostedEvents3 uintptr
	addr_QApplication_removePostedEvents2 uintptr
	addr_QApplication_translate2 uintptr
	addr_QApplication_translate3 uintptr
	addr_QApplication_exit2 uintptr
	addr_QApplication_connect3 uintptr
	addr_QApplication_connect4 uintptr
	pg_QApplication_connect_focusChanged func(uintptr, int64)
	pg_QApplication_connect_fontDatabaseChanged func(uintptr, int64)
	pg_QApplication_connect_screenAdded func(uintptr, int64)
	pg_QApplication_connect_screenRemoved func(uintptr, int64)
	pg_QApplication_connect_primaryScreenChanged func(uintptr, int64)
	pg_QApplication_connect_lastWindowClosed func(uintptr, int64)
	pg_QApplication_connect_focusObjectChanged func(uintptr, int64)
	pg_QApplication_connect_focusWindowChanged func(uintptr, int64)
	pg_QApplication_connect_applicationStateChanged func(uintptr, int64)
	pg_QApplication_connect_layoutDirectionChanged func(uintptr, int64)
	pg_QApplication_connect_commitDataRequest func(uintptr, int64)
	pg_QApplication_connect_saveStateRequest func(uintptr, int64)
	pg_QApplication_connect_applicationDisplayNameChanged func(uintptr, int64)
	pg_QApplication_connect_paletteChanged func(uintptr, int64)
	pg_QApplication_connect_fontChanged func(uintptr, int64)
	pg_QApplication_connect_aboutToQuit func(uintptr, int64)
	pg_QApplication_connect_organizationNameChanged func(uintptr, int64)
	pg_QApplication_connect_organizationDomainChanged func(uintptr, int64)
	pg_QApplication_connect_applicationNameChanged func(uintptr, int64)
	pg_QApplication_connect_applicationVersionChanged func(uintptr, int64)
	pg_QApplication_connect_destroyed func(uintptr, int64)
	pg_QApplication_connect_destroyed2 func(uintptr, int64)
	pg_QApplication_connect_objectNameChanged func(uintptr, int64)
	pg_QApplication__disconnect func(int64)
	pg_QApplication_new_args func(uintptr, uintptr) uintptr
	pg_QApplication_delete func(uintptr)
)

func init() {
	symaddr(&addr_QApplication_chain, "QApplication_chain")
	symaddr(&addr_QApplication_tr, "QApplication_tr")
	symaddr(&addr_QApplication_style, "QApplication_style")
	symaddr(&addr_QApplication_setStyle, "QApplication_setStyle")
	symaddr(&addr_QApplication_setStyle2, "QApplication_setStyle2")
	symaddr(&addr_QApplication_palette, "QApplication_palette")
	symaddr(&addr_QApplication_palette2, "QApplication_palette2")
	symaddr(&addr_QApplication_setPalette, "QApplication_setPalette")
	symaddr(&addr_QApplication_font, "QApplication_font")
	symaddr(&addr_QApplication_font2, "QApplication_font2")
	symaddr(&addr_QApplication_font3, "QApplication_font3")
	symaddr(&addr_QApplication_setFont, "QApplication_setFont")
	symaddr(&addr_QApplication_fontMetrics, "QApplication_fontMetrics")
	symaddr(&addr_QApplication_allWidgets, "QApplication_allWidgets")
	symaddr(&addr_QApplication_topLevelWidgets, "QApplication_topLevelWidgets")
	symaddr(&addr_QApplication_activePopupWidget, "QApplication_activePopupWidget")
	symaddr(&addr_QApplication_activeModalWidget, "QApplication_activeModalWidget")
	symaddr(&addr_QApplication_focusWidget, "QApplication_focusWidget")
	symaddr(&addr_QApplication_activeWindow, "QApplication_activeWindow")
	symaddr(&addr_QApplication_setActiveWindow, "QApplication_setActiveWindow")
	symaddr(&addr_QApplication_widgetAt, "QApplication_widgetAt")
	symaddr(&addr_QApplication_widgetAt2, "QApplication_widgetAt2")
	symaddr(&addr_QApplication_topLevelAt, "QApplication_topLevelAt")
	symaddr(&addr_QApplication_topLevelAt2, "QApplication_topLevelAt2")
	symaddr(&addr_QApplication_beep, "QApplication_beep")
	symaddr(&addr_QApplication_alert, "QApplication_alert")
	symaddr(&addr_QApplication_setCursorFlashTime, "QApplication_setCursorFlashTime")
	symaddr(&addr_QApplication_cursorFlashTime, "QApplication_cursorFlashTime")
	symaddr(&addr_QApplication_setDoubleClickInterval, "QApplication_setDoubleClickInterval")
	symaddr(&addr_QApplication_doubleClickInterval, "QApplication_doubleClickInterval")
	symaddr(&addr_QApplication_setKeyboardInputInterval, "QApplication_setKeyboardInputInterval")
	symaddr(&addr_QApplication_keyboardInputInterval, "QApplication_keyboardInputInterval")
	symaddr(&addr_QApplication_setWheelScrollLines, "QApplication_setWheelScrollLines")
	symaddr(&addr_QApplication_wheelScrollLines, "QApplication_wheelScrollLines")
	symaddr(&addr_QApplication_setStartDragTime, "QApplication_setStartDragTime")
	symaddr(&addr_QApplication_startDragTime, "QApplication_startDragTime")
	symaddr(&addr_QApplication_setStartDragDistance, "QApplication_setStartDragDistance")
	symaddr(&addr_QApplication_startDragDistance, "QApplication_startDragDistance")
	symaddr(&addr_QApplication_isEffectEnabled, "QApplication_isEffectEnabled")
	symaddr(&addr_QApplication_setEffectEnabled, "QApplication_setEffectEnabled")
	symaddr(&addr_QApplication_exec, "QApplication_exec")
	symaddr(&addr_QApplication_focusChanged, "QApplication_focusChanged")
	symaddr(&addr_QApplication_styleSheet, "QApplication_styleSheet")
	symaddr(&addr_QApplication_autoSipEnabled, "QApplication_autoSipEnabled")
	symaddr(&addr_QApplication_setStyleSheet, "QApplication_setStyleSheet")
	symaddr(&addr_QApplication_setAutoSipEnabled, "QApplication_setAutoSipEnabled")
	symaddr(&addr_QApplication_closeAllWindows, "QApplication_closeAllWindows")
	symaddr(&addr_QApplication_aboutQt, "QApplication_aboutQt")
	symaddr(&addr_QApplication_setApplicationDisplayName, "QApplication_setApplicationDisplayName")
	symaddr(&addr_QApplication_applicationDisplayName, "QApplication_applicationDisplayName")
	symaddr(&addr_QApplication_setDesktopFileName, "QApplication_setDesktopFileName")
	symaddr(&addr_QApplication_desktopFileName, "QApplication_desktopFileName")
	symaddr(&addr_QApplication_allWindows, "QApplication_allWindows")
	symaddr(&addr_QApplication_topLevelWindows, "QApplication_topLevelWindows")
	symaddr(&addr_QApplication_setWindowIcon, "QApplication_setWindowIcon")
	symaddr(&addr_QApplication_windowIcon, "QApplication_windowIcon")
	symaddr(&addr_QApplication_platformName, "QApplication_platformName")
	symaddr(&addr_QApplication_modalWindow, "QApplication_modalWindow")
	symaddr(&addr_QApplication_focusWindow, "QApplication_focusWindow")
	symaddr(&addr_QApplication_focusObject, "QApplication_focusObject")
	symaddr(&addr_QApplication_primaryScreen, "QApplication_primaryScreen")
	symaddr(&addr_QApplication_screens, "QApplication_screens")
	symaddr(&addr_QApplication_screenAt, "QApplication_screenAt")
	symaddr(&addr_QApplication_overrideCursor, "QApplication_overrideCursor")
	symaddr(&addr_QApplication_setOverrideCursor, "QApplication_setOverrideCursor")
	symaddr(&addr_QApplication_changeOverrideCursor, "QApplication_changeOverrideCursor")
	symaddr(&addr_QApplication_restoreOverrideCursor, "QApplication_restoreOverrideCursor")
	symaddr(&addr_QApplication_clipboard, "QApplication_clipboard")
	symaddr(&addr_QApplication_keyboardModifiers, "QApplication_keyboardModifiers")
	symaddr(&addr_QApplication_queryKeyboardModifiers, "QApplication_queryKeyboardModifiers")
	symaddr(&addr_QApplication_mouseButtons, "QApplication_mouseButtons")
	symaddr(&addr_QApplication_setLayoutDirection, "QApplication_setLayoutDirection")
	symaddr(&addr_QApplication_layoutDirection, "QApplication_layoutDirection")
	symaddr(&addr_QApplication_isRightToLeft, "QApplication_isRightToLeft")
	symaddr(&addr_QApplication_isLeftToRight, "QApplication_isLeftToRight")
	symaddr(&addr_QApplication_styleHints, "QApplication_styleHints")
	symaddr(&addr_QApplication_setDesktopSettingsAware, "QApplication_setDesktopSettingsAware")
	symaddr(&addr_QApplication_desktopSettingsAware, "QApplication_desktopSettingsAware")
	symaddr(&addr_QApplication_inputMethod, "QApplication_inputMethod")
	symaddr(&addr_QApplication_setQuitOnLastWindowClosed, "QApplication_setQuitOnLastWindowClosed")
	symaddr(&addr_QApplication_quitOnLastWindowClosed, "QApplication_quitOnLastWindowClosed")
	symaddr(&addr_QApplication_applicationState, "QApplication_applicationState")
	symaddr(&addr_QApplication_setHighDpiScaleFactorRoundingPolicy, "QApplication_setHighDpiScaleFactorRoundingPolicy")
	symaddr(&addr_QApplication_highDpiScaleFactorRoundingPolicy, "QApplication_highDpiScaleFactorRoundingPolicy")
	symaddr(&addr_QApplication_sync, "QApplication_sync")
	symaddr(&addr_QApplication_arguments, "QApplication_arguments")
	symaddr(&addr_QApplication_setAttribute, "QApplication_setAttribute")
	symaddr(&addr_QApplication_testAttribute, "QApplication_testAttribute")
	symaddr(&addr_QApplication_setOrganizationDomain, "QApplication_setOrganizationDomain")
	symaddr(&addr_QApplication_organizationDomain, "QApplication_organizationDomain")
	symaddr(&addr_QApplication_setOrganizationName, "QApplication_setOrganizationName")
	symaddr(&addr_QApplication_organizationName, "QApplication_organizationName")
	symaddr(&addr_QApplication_setApplicationName, "QApplication_setApplicationName")
	symaddr(&addr_QApplication_applicationName, "QApplication_applicationName")
	symaddr(&addr_QApplication_setApplicationVersion, "QApplication_setApplicationVersion")
	symaddr(&addr_QApplication_applicationVersion, "QApplication_applicationVersion")
	symaddr(&addr_QApplication_setSetuidAllowed, "QApplication_setSetuidAllowed")
	symaddr(&addr_QApplication_isSetuidAllowed, "QApplication_isSetuidAllowed")
	symaddr(&addr_QApplication_instance, "QApplication_instance")
	symaddr(&addr_QApplication_instanceExists, "QApplication_instanceExists")
	symaddr(&addr_QApplication_processEvents, "QApplication_processEvents")
	symaddr(&addr_QApplication_processEvents2, "QApplication_processEvents2")
	symaddr(&addr_QApplication_processEvents3, "QApplication_processEvents3")
	symaddr(&addr_QApplication_sendEvent, "QApplication_sendEvent")
	symaddr(&addr_QApplication_postEvent, "QApplication_postEvent")
	symaddr(&addr_QApplication_sendPostedEvents, "QApplication_sendPostedEvents")
	symaddr(&addr_QApplication_removePostedEvents, "QApplication_removePostedEvents")
	symaddr(&addr_QApplication_eventDispatcher, "QApplication_eventDispatcher")
	symaddr(&addr_QApplication_setEventDispatcher, "QApplication_setEventDispatcher")
	symaddr(&addr_QApplication_startingUp, "QApplication_startingUp")
	symaddr(&addr_QApplication_closingDown, "QApplication_closingDown")
	symaddr(&addr_QApplication_applicationDirPath, "QApplication_applicationDirPath")
	symaddr(&addr_QApplication_applicationFilePath, "QApplication_applicationFilePath")
	symaddr(&addr_QApplication_applicationPid, "QApplication_applicationPid")
	symaddr(&addr_QApplication_setLibraryPaths, "QApplication_setLibraryPaths")
	symaddr(&addr_QApplication_libraryPaths, "QApplication_libraryPaths")
	symaddr(&addr_QApplication_addLibraryPath, "QApplication_addLibraryPath")
	symaddr(&addr_QApplication_removeLibraryPath, "QApplication_removeLibraryPath")
	symaddr(&addr_QApplication_installTranslator, "QApplication_installTranslator")
	symaddr(&addr_QApplication_removeTranslator, "QApplication_removeTranslator")
	symaddr(&addr_QApplication_translate, "QApplication_translate")
	symaddr(&addr_QApplication_isQuitLockEnabled, "QApplication_isQuitLockEnabled")
	symaddr(&addr_QApplication_setQuitLockEnabled, "QApplication_setQuitLockEnabled")
	symaddr(&addr_QApplication_quit, "QApplication_quit")
	symaddr(&addr_QApplication_exit, "QApplication_exit")
	symaddr(&addr_QApplication_connect, "QApplication_connect")
	symaddr(&addr_QApplication_connect2, "QApplication_connect2")
	symaddr(&addr_QApplication_tr2, "QApplication_tr2")
	symaddr(&addr_QApplication_tr3, "QApplication_tr3")
	symaddr(&addr_QApplication_setPalette2, "QApplication_setPalette2")
	symaddr(&addr_QApplication_setFont2, "QApplication_setFont2")
	symaddr(&addr_QApplication_alert2, "QApplication_alert2")
	symaddr(&addr_QApplication_setEffectEnabled2, "QApplication_setEffectEnabled2")
	symaddr(&addr_QApplication_setAttribute2, "QApplication_setAttribute2")
	symaddr(&addr_QApplication_processEvents4, "QApplication_processEvents4")
	symaddr(&addr_QApplication_postEvent2, "QApplication_postEvent2")
	symaddr(&addr_QApplication_sendPostedEvents2, "QApplication_sendPostedEvents2")
	symaddr(&addr_QApplication_sendPostedEvents3, "QApplication_sendPostedEvents3")
	symaddr(&addr_QApplication_removePostedEvents2, "QApplication_removePostedEvents2")
	symaddr(&addr_QApplication_translate2, "QApplication_translate2")
	symaddr(&addr_QApplication_translate3, "QApplication_translate3")
	symaddr(&addr_QApplication_exit2, "QApplication_exit2")
	symaddr(&addr_QApplication_connect3, "QApplication_connect3")
	symaddr(&addr_QApplication_connect4, "QApplication_connect4")
	reg(&pg_QApplication_connect_focusChanged, "QApplication_connect_focusChanged")
	reg(&pg_QApplication_connect_fontDatabaseChanged, "QApplication_connect_fontDatabaseChanged")
	reg(&pg_QApplication_connect_screenAdded, "QApplication_connect_screenAdded")
	reg(&pg_QApplication_connect_screenRemoved, "QApplication_connect_screenRemoved")
	reg(&pg_QApplication_connect_primaryScreenChanged, "QApplication_connect_primaryScreenChanged")
	reg(&pg_QApplication_connect_lastWindowClosed, "QApplication_connect_lastWindowClosed")
	reg(&pg_QApplication_connect_focusObjectChanged, "QApplication_connect_focusObjectChanged")
	reg(&pg_QApplication_connect_focusWindowChanged, "QApplication_connect_focusWindowChanged")
	reg(&pg_QApplication_connect_applicationStateChanged, "QApplication_connect_applicationStateChanged")
	reg(&pg_QApplication_connect_layoutDirectionChanged, "QApplication_connect_layoutDirectionChanged")
	reg(&pg_QApplication_connect_commitDataRequest, "QApplication_connect_commitDataRequest")
	reg(&pg_QApplication_connect_saveStateRequest, "QApplication_connect_saveStateRequest")
	reg(&pg_QApplication_connect_applicationDisplayNameChanged, "QApplication_connect_applicationDisplayNameChanged")
	reg(&pg_QApplication_connect_paletteChanged, "QApplication_connect_paletteChanged")
	reg(&pg_QApplication_connect_fontChanged, "QApplication_connect_fontChanged")
	reg(&pg_QApplication_connect_aboutToQuit, "QApplication_connect_aboutToQuit")
	reg(&pg_QApplication_connect_organizationNameChanged, "QApplication_connect_organizationNameChanged")
	reg(&pg_QApplication_connect_organizationDomainChanged, "QApplication_connect_organizationDomainChanged")
	reg(&pg_QApplication_connect_applicationNameChanged, "QApplication_connect_applicationNameChanged")
	reg(&pg_QApplication_connect_applicationVersionChanged, "QApplication_connect_applicationVersionChanged")
	reg(&pg_QApplication_connect_destroyed, "QApplication_connect_destroyed")
	reg(&pg_QApplication_connect_destroyed2, "QApplication_connect_destroyed2")
	reg(&pg_QApplication_connect_objectNameChanged, "QApplication_connect_objectNameChanged")
	reg(&pg_QApplication__disconnect, "QApplication__disconnect")
	reg(&pg_QApplication_new_args, "QApplication_new_args")
	reg(&pg_QApplication_delete, "QApplication_delete")
}

type QApplication struct {
	h uintptr
	*QGuiApplication
}

func bindQApplication(h uintptr) *QApplication {
	if h == 0 {
		return nil
	}
	var __c [3]uintptr
	purego.SyscallN(addr_QApplication_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QApplication{h: h, QGuiApplication: &QGuiApplication{h: __c[0], QCoreApplication: &QCoreApplication{h: __c[1], QObject: &QObject{h: __c[2]}}}}
}

func newownedQApplication(h uintptr) *QApplication { return bindQApplication(h) }

func (this *QApplication) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QApplication) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQApplication(h unsafe.Pointer) *QApplication { return bindQApplication(uintptr(h)) }
func UnsafeNewQApplicationOwned(h unsafe.Pointer) *QApplication { return newownedQApplication(uintptr(h)) }

func (this *QApplication) UnsafeRelease() {
	this.h = 0
	if this.QGuiApplication != nil {
		this.QGuiApplication.UnsafeRelease()
	}
}

func QApplication_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QApplication_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func QApplication_Style() *QStyle {
	__r, _, _ := purego.SyscallN(addr_QApplication_style)
	return bindQStyle(__r)
}

func QApplication_SetStyle(arg1 *QStyle) {
	purego.SyscallN(addr_QApplication_setStyle, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(arg1)
}

func QApplication_SetStyleWithQString(arg1 string) *QStyle {
	var arg1_p uintptr
	if len(arg1) > 0 {
		arg1_p = uintptr(unsafe.Pointer(unsafe.StringData(arg1)))
	}
	__r, _, _ := purego.SyscallN(addr_QApplication_setStyle2, arg1_p, uintptr(len(arg1)))
	runtime.KeepAlive(arg1)
	return bindQStyle(__r)
}

func QApplication_Palette(arg1 *QWidget) *QPalette {
	__r, _, _ := purego.SyscallN(addr_QApplication_palette, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(arg1)
	return newownedQPalette(__r)
}

func QApplication_PaletteWithClassName(className string) *QPalette {
	className_b := append([]byte(className), 0)
	__r, _, _ := purego.SyscallN(addr_QApplication_palette2, uintptr(unsafe.Pointer(&className_b[0])))
	runtime.KeepAlive(className_b)
	return newownedQPalette(__r)
}

func QApplication_SetPalette(arg1 *QPalette) {
	purego.SyscallN(addr_QApplication_setPalette, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(arg1)
}

func QApplication_Font() *QFont {
	__r, _, _ := purego.SyscallN(addr_QApplication_font)
	return newownedQFont(__r)
}

func QApplication_FontWithQWidget(arg1 *QWidget) *QFont {
	__r, _, _ := purego.SyscallN(addr_QApplication_font2, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(arg1)
	return newownedQFont(__r)
}

func QApplication_FontWithClassName(className string) *QFont {
	className_b := append([]byte(className), 0)
	__r, _, _ := purego.SyscallN(addr_QApplication_font3, uintptr(unsafe.Pointer(&className_b[0])))
	runtime.KeepAlive(className_b)
	return newownedQFont(__r)
}

func QApplication_SetFont(arg1 *QFont) {
	purego.SyscallN(addr_QApplication_setFont, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(arg1)
}

func QApplication_FontMetrics() *QFontMetrics {
	__r, _, _ := purego.SyscallN(addr_QApplication_fontMetrics)
	return newownedQFontMetrics(__r)
}

func QApplication_AllWidgets() []*QWidget {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QApplication_allWidgets, uintptr(unsafe.Pointer(&__n)))
	__out := make([]*QWidget, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = bindQWidget(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QApplication_TopLevelWidgets() []*QWidget {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QApplication_topLevelWidgets, uintptr(unsafe.Pointer(&__n)))
	__out := make([]*QWidget, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = bindQWidget(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QApplication_ActivePopupWidget() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QApplication_activePopupWidget)
	return bindQWidget(__r)
}

func QApplication_ActiveModalWidget() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QApplication_activeModalWidget)
	return bindQWidget(__r)
}

func QApplication_FocusWidget() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QApplication_focusWidget)
	return bindQWidget(__r)
}

func QApplication_ActiveWindow() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QApplication_activeWindow)
	return bindQWidget(__r)
}

func QApplication_SetActiveWindow(act *QWidget) {
	purego.SyscallN(addr_QApplication_setActiveWindow, uintptr(act.UnsafePointer()))
	runtime.KeepAlive(act)
}

func QApplication_WidgetAt(p *QPoint) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QApplication_widgetAt, uintptr(p.UnsafePointer()))
	runtime.KeepAlive(p)
	return bindQWidget(__r)
}

func QApplication_WidgetAt2(x int, y int) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QApplication_widgetAt2, uintptr(x), uintptr(y))
	return bindQWidget(__r)
}

func QApplication_TopLevelAt(p *QPoint) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QApplication_topLevelAt, uintptr(p.UnsafePointer()))
	runtime.KeepAlive(p)
	return bindQWidget(__r)
}

func QApplication_TopLevelAt2(x int, y int) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QApplication_topLevelAt2, uintptr(x), uintptr(y))
	return bindQWidget(__r)
}

func QApplication_Beep() {
	purego.SyscallN(addr_QApplication_beep)
}

func QApplication_Alert(widget *QWidget) {
	purego.SyscallN(addr_QApplication_alert, uintptr(widget.UnsafePointer()))
	runtime.KeepAlive(widget)
}

func QApplication_SetCursorFlashTime(arg1 int) {
	purego.SyscallN(addr_QApplication_setCursorFlashTime, uintptr(arg1))
}

func QApplication_CursorFlashTime() int {
	__r0, _, _ := purego.SyscallN(addr_QApplication_cursorFlashTime)
	__r := int32(uint32(__r0))
	return int(__r)
}

func QApplication_SetDoubleClickInterval(arg1 int) {
	purego.SyscallN(addr_QApplication_setDoubleClickInterval, uintptr(arg1))
}

func QApplication_DoubleClickInterval() int {
	__r0, _, _ := purego.SyscallN(addr_QApplication_doubleClickInterval)
	__r := int32(uint32(__r0))
	return int(__r)
}

func QApplication_SetKeyboardInputInterval(arg1 int) {
	purego.SyscallN(addr_QApplication_setKeyboardInputInterval, uintptr(arg1))
}

func QApplication_KeyboardInputInterval() int {
	__r0, _, _ := purego.SyscallN(addr_QApplication_keyboardInputInterval)
	__r := int32(uint32(__r0))
	return int(__r)
}

func QApplication_SetWheelScrollLines(arg1 int) {
	purego.SyscallN(addr_QApplication_setWheelScrollLines, uintptr(arg1))
}

func QApplication_WheelScrollLines() int {
	__r0, _, _ := purego.SyscallN(addr_QApplication_wheelScrollLines)
	__r := int32(uint32(__r0))
	return int(__r)
}

func QApplication_SetStartDragTime(ms int) {
	purego.SyscallN(addr_QApplication_setStartDragTime, uintptr(ms))
}

func QApplication_StartDragTime() int {
	__r0, _, _ := purego.SyscallN(addr_QApplication_startDragTime)
	__r := int32(uint32(__r0))
	return int(__r)
}

func QApplication_SetStartDragDistance(l int) {
	purego.SyscallN(addr_QApplication_setStartDragDistance, uintptr(l))
}

func QApplication_StartDragDistance() int {
	__r0, _, _ := purego.SyscallN(addr_QApplication_startDragDistance)
	__r := int32(uint32(__r0))
	return int(__r)
}

func QApplication_IsEffectEnabled(arg1 UIEffect) bool {
	__r0, _, _ := purego.SyscallN(addr_QApplication_isEffectEnabled, uintptr(arg1))
	__r := int32(uint32(__r0))
	return __r != 0
}

func QApplication_SetEffectEnabled(arg1 UIEffect) {
	purego.SyscallN(addr_QApplication_setEffectEnabled, uintptr(arg1))
}

func QApplication_Exec() int {
	__r0, _, _ := purego.SyscallN(addr_QApplication_exec)
	__r := int32(uint32(__r0))
	return int(__r)
}

func (this *QApplication) FocusChanged(old *QWidget, now *QWidget) {
	purego.SyscallN(addr_QApplication_focusChanged, this.h, uintptr(old.UnsafePointer()), uintptr(now.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(old)
	runtime.KeepAlive(now)
}

func (this *QApplication) StyleSheet() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QApplication_styleSheet, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QApplication) AutoSipEnabled() bool {
	__r0, _, _ := purego.SyscallN(addr_QApplication_autoSipEnabled, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QApplication) SetStyleSheet(sheet string) {
	var sheet_p uintptr
	if len(sheet) > 0 {
		sheet_p = uintptr(unsafe.Pointer(unsafe.StringData(sheet)))
	}
	purego.SyscallN(addr_QApplication_setStyleSheet, this.h, sheet_p, uintptr(len(sheet)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sheet)
}

func (this *QApplication) SetAutoSipEnabled(enabled bool) {
	purego.SyscallN(addr_QApplication_setAutoSipEnabled, this.h, uintptr(b2i(enabled)))
	runtime.KeepAlive(this)
}

func QApplication_CloseAllWindows() {
	purego.SyscallN(addr_QApplication_closeAllWindows)
}

func QApplication_AboutQt() {
	purego.SyscallN(addr_QApplication_aboutQt)
}

func QApplication_SetApplicationDisplayName(name string) {
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	purego.SyscallN(addr_QApplication_setApplicationDisplayName, name_p, uintptr(len(name)))
	runtime.KeepAlive(name)
}

func QApplication_ApplicationDisplayName() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QApplication_applicationDisplayName, uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QApplication_SetDesktopFileName(name string) {
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	purego.SyscallN(addr_QApplication_setDesktopFileName, name_p, uintptr(len(name)))
	runtime.KeepAlive(name)
}

func QApplication_DesktopFileName() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QApplication_desktopFileName, uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QApplication_AllWindows() []*QWindow {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QApplication_allWindows, uintptr(unsafe.Pointer(&__n)))
	__out := make([]*QWindow, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = bindQWindow(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QApplication_TopLevelWindows() []*QWindow {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QApplication_topLevelWindows, uintptr(unsafe.Pointer(&__n)))
	__out := make([]*QWindow, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = bindQWindow(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QApplication_SetWindowIcon(icon *QIcon) {
	purego.SyscallN(addr_QApplication_setWindowIcon, uintptr(icon.UnsafePointer()))
	runtime.KeepAlive(icon)
}

func QApplication_WindowIcon() *QIcon {
	__r, _, _ := purego.SyscallN(addr_QApplication_windowIcon)
	return newownedQIcon(__r)
}

func QApplication_PlatformName() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QApplication_platformName, uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QApplication_ModalWindow() *QWindow {
	__r, _, _ := purego.SyscallN(addr_QApplication_modalWindow)
	return bindQWindow(__r)
}

func QApplication_FocusWindow() *QWindow {
	__r, _, _ := purego.SyscallN(addr_QApplication_focusWindow)
	return bindQWindow(__r)
}

func QApplication_FocusObject() *QObject {
	__r, _, _ := purego.SyscallN(addr_QApplication_focusObject)
	return bindQObject(__r)
}

func QApplication_PrimaryScreen() *QScreen {
	__r, _, _ := purego.SyscallN(addr_QApplication_primaryScreen)
	return bindQScreen(__r)
}

func QApplication_Screens() []*QScreen {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QApplication_screens, uintptr(unsafe.Pointer(&__n)))
	__out := make([]*QScreen, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = bindQScreen(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QApplication_ScreenAt(point *QPoint) *QScreen {
	__r, _, _ := purego.SyscallN(addr_QApplication_screenAt, uintptr(point.UnsafePointer()))
	runtime.KeepAlive(point)
	return bindQScreen(__r)
}

func QApplication_OverrideCursor() *QCursor {
	__r, _, _ := purego.SyscallN(addr_QApplication_overrideCursor)
	return bindQCursor(__r)
}

func QApplication_SetOverrideCursor(arg1 *QCursor) {
	purego.SyscallN(addr_QApplication_setOverrideCursor, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(arg1)
}

func QApplication_ChangeOverrideCursor(arg1 *QCursor) {
	purego.SyscallN(addr_QApplication_changeOverrideCursor, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(arg1)
}

func QApplication_RestoreOverrideCursor() {
	purego.SyscallN(addr_QApplication_restoreOverrideCursor)
}

func QApplication_Clipboard() *QClipboard {
	__r, _, _ := purego.SyscallN(addr_QApplication_clipboard)
	return bindQClipboard(__r)
}

func QApplication_KeyboardModifiers() KeyboardModifier {
	__r0, _, _ := purego.SyscallN(addr_QApplication_keyboardModifiers)
	__r := int32(uint32(__r0))
	return KeyboardModifier(__r)
}

func QApplication_QueryKeyboardModifiers() KeyboardModifier {
	__r0, _, _ := purego.SyscallN(addr_QApplication_queryKeyboardModifiers)
	__r := int32(uint32(__r0))
	return KeyboardModifier(__r)
}

func QApplication_MouseButtons() MouseButton {
	__r0, _, _ := purego.SyscallN(addr_QApplication_mouseButtons)
	__r := int32(uint32(__r0))
	return MouseButton(__r)
}

func QApplication_SetLayoutDirection(direction LayoutDirection) {
	purego.SyscallN(addr_QApplication_setLayoutDirection, uintptr(direction))
}

func QApplication_LayoutDirection() LayoutDirection {
	__r0, _, _ := purego.SyscallN(addr_QApplication_layoutDirection)
	__r := int32(uint32(__r0))
	return LayoutDirection(__r)
}

func QApplication_IsRightToLeft() bool {
	__r0, _, _ := purego.SyscallN(addr_QApplication_isRightToLeft)
	__r := int32(uint32(__r0))
	return __r != 0
}

func QApplication_IsLeftToRight() bool {
	__r0, _, _ := purego.SyscallN(addr_QApplication_isLeftToRight)
	__r := int32(uint32(__r0))
	return __r != 0
}

func QApplication_StyleHints() *QStyleHints {
	__r, _, _ := purego.SyscallN(addr_QApplication_styleHints)
	return bindQStyleHints(__r)
}

func QApplication_SetDesktopSettingsAware(on bool) {
	purego.SyscallN(addr_QApplication_setDesktopSettingsAware, uintptr(b2i(on)))
}

func QApplication_DesktopSettingsAware() bool {
	__r0, _, _ := purego.SyscallN(addr_QApplication_desktopSettingsAware)
	__r := int32(uint32(__r0))
	return __r != 0
}

func QApplication_InputMethod() *QInputMethod {
	__r, _, _ := purego.SyscallN(addr_QApplication_inputMethod)
	return bindQInputMethod(__r)
}

func QApplication_SetQuitOnLastWindowClosed(quit bool) {
	purego.SyscallN(addr_QApplication_setQuitOnLastWindowClosed, uintptr(b2i(quit)))
}

func QApplication_QuitOnLastWindowClosed() bool {
	__r0, _, _ := purego.SyscallN(addr_QApplication_quitOnLastWindowClosed)
	__r := int32(uint32(__r0))
	return __r != 0
}

func QApplication_ApplicationState() ApplicationState {
	__r0, _, _ := purego.SyscallN(addr_QApplication_applicationState)
	__r := int32(uint32(__r0))
	return ApplicationState(__r)
}

func QApplication_SetHighDpiScaleFactorRoundingPolicy(policy HighDpiScaleFactorRoundingPolicy) {
	purego.SyscallN(addr_QApplication_setHighDpiScaleFactorRoundingPolicy, uintptr(policy))
}

func QApplication_HighDpiScaleFactorRoundingPolicy() HighDpiScaleFactorRoundingPolicy {
	__r0, _, _ := purego.SyscallN(addr_QApplication_highDpiScaleFactorRoundingPolicy)
	__r := int32(uint32(__r0))
	return HighDpiScaleFactorRoundingPolicy(__r)
}

func QApplication_Sync() {
	purego.SyscallN(addr_QApplication_sync)
}

func QApplication_Arguments() []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QApplication_arguments, uintptr(unsafe.Pointer(&__n)))
	__out := make([]string, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = pgGoStr(__sl[__i])
		pgFree(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QApplication_SetAttribute(attribute ApplicationAttribute) {
	purego.SyscallN(addr_QApplication_setAttribute, uintptr(attribute))
}

func QApplication_TestAttribute(attribute ApplicationAttribute) bool {
	__r0, _, _ := purego.SyscallN(addr_QApplication_testAttribute, uintptr(attribute))
	__r := int32(uint32(__r0))
	return __r != 0
}

func QApplication_SetOrganizationDomain(orgDomain string) {
	var orgDomain_p uintptr
	if len(orgDomain) > 0 {
		orgDomain_p = uintptr(unsafe.Pointer(unsafe.StringData(orgDomain)))
	}
	purego.SyscallN(addr_QApplication_setOrganizationDomain, orgDomain_p, uintptr(len(orgDomain)))
	runtime.KeepAlive(orgDomain)
}

func QApplication_OrganizationDomain() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QApplication_organizationDomain, uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QApplication_SetOrganizationName(orgName string) {
	var orgName_p uintptr
	if len(orgName) > 0 {
		orgName_p = uintptr(unsafe.Pointer(unsafe.StringData(orgName)))
	}
	purego.SyscallN(addr_QApplication_setOrganizationName, orgName_p, uintptr(len(orgName)))
	runtime.KeepAlive(orgName)
}

func QApplication_OrganizationName() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QApplication_organizationName, uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QApplication_SetApplicationName(application string) {
	var application_p uintptr
	if len(application) > 0 {
		application_p = uintptr(unsafe.Pointer(unsafe.StringData(application)))
	}
	purego.SyscallN(addr_QApplication_setApplicationName, application_p, uintptr(len(application)))
	runtime.KeepAlive(application)
}

func QApplication_ApplicationName() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QApplication_applicationName, uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QApplication_SetApplicationVersion(version string) {
	var version_p uintptr
	if len(version) > 0 {
		version_p = uintptr(unsafe.Pointer(unsafe.StringData(version)))
	}
	purego.SyscallN(addr_QApplication_setApplicationVersion, version_p, uintptr(len(version)))
	runtime.KeepAlive(version)
}

func QApplication_ApplicationVersion() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QApplication_applicationVersion, uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QApplication_SetSetuidAllowed(allow bool) {
	purego.SyscallN(addr_QApplication_setSetuidAllowed, uintptr(b2i(allow)))
}

func QApplication_IsSetuidAllowed() bool {
	__r0, _, _ := purego.SyscallN(addr_QApplication_isSetuidAllowed)
	__r := int32(uint32(__r0))
	return __r != 0
}

func QApplication_Instance() *QCoreApplication {
	__r, _, _ := purego.SyscallN(addr_QApplication_instance)
	return bindQCoreApplication(__r)
}

func QApplication_InstanceExists() bool {
	__r0, _, _ := purego.SyscallN(addr_QApplication_instanceExists)
	__r := int32(uint32(__r0))
	return __r != 0
}

func QApplication_ProcessEvents() {
	purego.SyscallN(addr_QApplication_processEvents)
}

func QApplication_ProcessEvents2(flags QEventLoop__ProcessEventsFlag, maxtime int) {
	purego.SyscallN(addr_QApplication_processEvents2, uintptr(flags), uintptr(maxtime))
}

func QApplication_ProcessEvents3(flags QEventLoop__ProcessEventsFlag, deadline *QDeadlineTimer) {
	purego.SyscallN(addr_QApplication_processEvents3, uintptr(flags), uintptr(deadline.UnsafePointer()))
	runtime.KeepAlive(deadline)
}

func QApplication_SendEvent(receiver *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QApplication_sendEvent, uintptr(receiver.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(event)
	return __r != 0
}

func QApplication_PostEvent(receiver *QObject, event *QEvent) {
	purego.SyscallN(addr_QApplication_postEvent, uintptr(receiver.UnsafePointer()), uintptr(event.UnsafePointer()))
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(event)
}

func QApplication_SendPostedEvents() {
	purego.SyscallN(addr_QApplication_sendPostedEvents)
}

func QApplication_RemovePostedEvents(receiver *QObject) {
	purego.SyscallN(addr_QApplication_removePostedEvents, uintptr(receiver.UnsafePointer()))
	runtime.KeepAlive(receiver)
}

func QApplication_EventDispatcher() *QAbstractEventDispatcher {
	__r, _, _ := purego.SyscallN(addr_QApplication_eventDispatcher)
	return bindQAbstractEventDispatcher(__r)
}

func QApplication_SetEventDispatcher(eventDispatcher *QAbstractEventDispatcher) {
	purego.SyscallN(addr_QApplication_setEventDispatcher, uintptr(eventDispatcher.UnsafePointer()))
	runtime.KeepAlive(eventDispatcher)
}

func QApplication_StartingUp() bool {
	__r0, _, _ := purego.SyscallN(addr_QApplication_startingUp)
	__r := int32(uint32(__r0))
	return __r != 0
}

func QApplication_ClosingDown() bool {
	__r0, _, _ := purego.SyscallN(addr_QApplication_closingDown)
	__r := int32(uint32(__r0))
	return __r != 0
}

func QApplication_ApplicationDirPath() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QApplication_applicationDirPath, uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QApplication_ApplicationFilePath() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QApplication_applicationFilePath, uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QApplication_ApplicationPid() int64 {
	__r0, _, _ := purego.SyscallN(addr_QApplication_applicationPid)
	__r := int64(__r0)
	return __r
}

func QApplication_SetLibraryPaths(arg1 []string) {
	arg1_arr := make([]uintptr, len(arg1))
	for __i, __e := range arg1 {
		arg1_arr[__i] = pgCString(__e)
	}
	var arg1_p uintptr
	if len(arg1_arr) > 0 {
		arg1_p = uintptr(unsafe.Pointer(&arg1_arr[0]))
	}
	purego.SyscallN(addr_QApplication_setLibraryPaths, arg1_p, uintptr(len(arg1)))
	for __i := range arg1_arr {
		pgFree(arg1_arr[__i])
	}
	runtime.KeepAlive(arg1_arr)
}

func QApplication_LibraryPaths() []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QApplication_libraryPaths, uintptr(unsafe.Pointer(&__n)))
	__out := make([]string, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = pgGoStr(__sl[__i])
		pgFree(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QApplication_AddLibraryPath(arg1 string) {
	var arg1_p uintptr
	if len(arg1) > 0 {
		arg1_p = uintptr(unsafe.Pointer(unsafe.StringData(arg1)))
	}
	purego.SyscallN(addr_QApplication_addLibraryPath, arg1_p, uintptr(len(arg1)))
	runtime.KeepAlive(arg1)
}

func QApplication_RemoveLibraryPath(arg1 string) {
	var arg1_p uintptr
	if len(arg1) > 0 {
		arg1_p = uintptr(unsafe.Pointer(unsafe.StringData(arg1)))
	}
	purego.SyscallN(addr_QApplication_removeLibraryPath, arg1_p, uintptr(len(arg1)))
	runtime.KeepAlive(arg1)
}

func QApplication_InstallTranslator(messageFile *QTranslator) bool {
	__r0, _, _ := purego.SyscallN(addr_QApplication_installTranslator, uintptr(messageFile.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(messageFile)
	return __r != 0
}

func QApplication_RemoveTranslator(messageFile *QTranslator) bool {
	__r0, _, _ := purego.SyscallN(addr_QApplication_removeTranslator, uintptr(messageFile.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(messageFile)
	return __r != 0
}

func QApplication_Translate(context string, key string) string {
	context_b := append([]byte(context), 0)
	key_b := append([]byte(key), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QApplication_translate, uintptr(unsafe.Pointer(&context_b[0])), uintptr(unsafe.Pointer(&key_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(context_b)
	runtime.KeepAlive(key_b)
	return pgStr(__r, __n)
}

func QApplication_IsQuitLockEnabled() bool {
	__r0, _, _ := purego.SyscallN(addr_QApplication_isQuitLockEnabled)
	__r := int32(uint32(__r0))
	return __r != 0
}

func QApplication_SetQuitLockEnabled(enabled bool) {
	purego.SyscallN(addr_QApplication_setQuitLockEnabled, uintptr(b2i(enabled)))
}

func QApplication_Quit() {
	purego.SyscallN(addr_QApplication_quit)
}

func QApplication_Exit() {
	purego.SyscallN(addr_QApplication_exit)
}

func QApplication_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QApplication_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QApplication_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QApplication_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func QApplication_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QApplication_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QApplication_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QApplication_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QApplication_SetPalette2(arg1 *QPalette, className string) {
	className_b := append([]byte(className), 0)
	purego.SyscallN(addr_QApplication_setPalette2, uintptr(arg1.UnsafePointer()), uintptr(unsafe.Pointer(&className_b[0])))
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(className_b)
}

func QApplication_SetFont2(arg1 *QFont, className string) {
	className_b := append([]byte(className), 0)
	purego.SyscallN(addr_QApplication_setFont2, uintptr(arg1.UnsafePointer()), uintptr(unsafe.Pointer(&className_b[0])))
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(className_b)
}

func QApplication_Alert2(widget *QWidget, duration int) {
	purego.SyscallN(addr_QApplication_alert2, uintptr(widget.UnsafePointer()), uintptr(duration))
	runtime.KeepAlive(widget)
}

func QApplication_SetEffectEnabled2(arg1 UIEffect, enable bool) {
	purego.SyscallN(addr_QApplication_setEffectEnabled2, uintptr(arg1), uintptr(b2i(enable)))
}

func QApplication_SetAttribute2(attribute ApplicationAttribute, on bool) {
	purego.SyscallN(addr_QApplication_setAttribute2, uintptr(attribute), uintptr(b2i(on)))
}

func QApplication_ProcessEventsWithFlags(flags QEventLoop__ProcessEventsFlag) {
	purego.SyscallN(addr_QApplication_processEvents4, uintptr(flags))
}

func QApplication_PostEvent2(receiver *QObject, event *QEvent, priority int) {
	purego.SyscallN(addr_QApplication_postEvent2, uintptr(receiver.UnsafePointer()), uintptr(event.UnsafePointer()), uintptr(priority))
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(event)
}

func QApplication_SendPostedEventsWithReceiver(receiver *QObject) {
	purego.SyscallN(addr_QApplication_sendPostedEvents2, uintptr(receiver.UnsafePointer()))
	runtime.KeepAlive(receiver)
}

func QApplication_SendPostedEvents2(receiver *QObject, event_type int) {
	purego.SyscallN(addr_QApplication_sendPostedEvents3, uintptr(receiver.UnsafePointer()), uintptr(event_type))
	runtime.KeepAlive(receiver)
}

func QApplication_RemovePostedEvents2(receiver *QObject, eventType int) {
	purego.SyscallN(addr_QApplication_removePostedEvents2, uintptr(receiver.UnsafePointer()), uintptr(eventType))
	runtime.KeepAlive(receiver)
}

func QApplication_Translate2(context string, key string, disambiguation string) string {
	context_b := append([]byte(context), 0)
	key_b := append([]byte(key), 0)
	disambiguation_b := append([]byte(disambiguation), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QApplication_translate2, uintptr(unsafe.Pointer(&context_b[0])), uintptr(unsafe.Pointer(&key_b[0])), uintptr(unsafe.Pointer(&disambiguation_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(context_b)
	runtime.KeepAlive(key_b)
	runtime.KeepAlive(disambiguation_b)
	return pgStr(__r, __n)
}

func QApplication_Translate3(context string, key string, disambiguation string, n int) string {
	context_b := append([]byte(context), 0)
	key_b := append([]byte(key), 0)
	disambiguation_b := append([]byte(disambiguation), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QApplication_translate3, uintptr(unsafe.Pointer(&context_b[0])), uintptr(unsafe.Pointer(&key_b[0])), uintptr(unsafe.Pointer(&disambiguation_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(context_b)
	runtime.KeepAlive(key_b)
	runtime.KeepAlive(disambiguation_b)
	return pgStr(__r, __n)
}

func QApplication_ExitWithRetcode(retcode int) {
	purego.SyscallN(addr_QApplication_exit2, uintptr(retcode))
}

func QApplication_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QApplication_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QApplication_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QApplication_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QApplication) OnFocusChanged(fn func(*QWidget, *QWidget)) int64 {
	id := registerCallback(func(old uintptr, now uintptr) { fn(UnsafeNewQWidget(unsafe.Pointer(old)), UnsafeNewQWidget(unsafe.Pointer(now))) })
	pg_QApplication_connect_focusChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QApplication) OnFontDatabaseChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QApplication_connect_fontDatabaseChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QApplication) OnScreenAdded(fn func(*QScreen)) int64 {
	id := registerCallback(func(screen uintptr) { fn(UnsafeNewQScreen(unsafe.Pointer(screen))) })
	pg_QApplication_connect_screenAdded(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QApplication) OnScreenRemoved(fn func(*QScreen)) int64 {
	id := registerCallback(func(screen uintptr) { fn(UnsafeNewQScreen(unsafe.Pointer(screen))) })
	pg_QApplication_connect_screenRemoved(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QApplication) OnPrimaryScreenChanged(fn func(*QScreen)) int64 {
	id := registerCallback(func(screen uintptr) { fn(UnsafeNewQScreen(unsafe.Pointer(screen))) })
	pg_QApplication_connect_primaryScreenChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QApplication) OnLastWindowClosed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QApplication_connect_lastWindowClosed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QApplication) OnFocusObjectChanged(fn func(*QObject)) int64 {
	id := registerCallback(func(focusObject uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(focusObject))) })
	pg_QApplication_connect_focusObjectChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QApplication) OnFocusWindowChanged(fn func(*QWindow)) int64 {
	id := registerCallback(func(focusWindow uintptr) { fn(UnsafeNewQWindow(unsafe.Pointer(focusWindow))) })
	pg_QApplication_connect_focusWindowChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QApplication) OnApplicationStateChanged(fn func(ApplicationState)) int64 {
	id := registerCallback(func(state int32) { fn(ApplicationState(state)) })
	pg_QApplication_connect_applicationStateChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QApplication) OnLayoutDirectionChanged(fn func(LayoutDirection)) int64 {
	id := registerCallback(func(direction int32) { fn(LayoutDirection(direction)) })
	pg_QApplication_connect_layoutDirectionChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QApplication) OnCommitDataRequest(fn func(*QSessionManager)) int64 {
	id := registerCallback(func(sessionManager uintptr) { fn(UnsafeNewQSessionManager(unsafe.Pointer(sessionManager))) })
	pg_QApplication_connect_commitDataRequest(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QApplication) OnSaveStateRequest(fn func(*QSessionManager)) int64 {
	id := registerCallback(func(sessionManager uintptr) { fn(UnsafeNewQSessionManager(unsafe.Pointer(sessionManager))) })
	pg_QApplication_connect_saveStateRequest(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QApplication) OnApplicationDisplayNameChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QApplication_connect_applicationDisplayNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QApplication) OnPaletteChanged(fn func(*QPalette)) int64 {
	id := registerCallback(func(pal uintptr) { fn(UnsafeNewQPalette(unsafe.Pointer(pal))) })
	pg_QApplication_connect_paletteChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QApplication) OnFontChanged(fn func(*QFont)) int64 {
	id := registerCallback(func(font uintptr) { fn(UnsafeNewQFont(unsafe.Pointer(font))) })
	pg_QApplication_connect_fontChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QApplication) OnAboutToQuit(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QApplication_connect_aboutToQuit(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QApplication) OnOrganizationNameChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QApplication_connect_organizationNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QApplication) OnOrganizationDomainChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QApplication_connect_organizationDomainChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QApplication) OnApplicationNameChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QApplication_connect_applicationNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QApplication) OnApplicationVersionChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QApplication_connect_applicationVersionChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QApplication) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QApplication_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QApplication) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QApplication_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QApplication) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QApplication_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QApplication) Disconnect(handle int64) {
	pg_QApplication__disconnect(handle)
	unregisterCallback(handle)
}

func NewQApplication(args []string) *QApplication {
	argc := pgMalloc(unsafe.Sizeof(int32(0)))
	*(*int32)(unsafe.Pointer(argc)) = int32(len(args))
	argv := pgMalloc(uintptr(len(args)+1) * unsafe.Sizeof(uintptr(0)))
	__av := unsafe.Slice((*uintptr)(unsafe.Pointer(argv)), len(args)+1)
	for i := range args {
		__av[i] = pgCString(args[i])
	}
	__av[len(args)] = 0
	runtime.LockOSThread()
	return newownedQApplication(pg_QApplication_new_args(argc, argv))
}

func (this *QApplication) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QApplication_delete(this.h)
	}
	this.UnsafeRelease()
}
