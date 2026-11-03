package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QGuiApplication_chain uintptr
	addr_QGuiApplication_tr uintptr
	addr_QGuiApplication_setApplicationDisplayName uintptr
	addr_QGuiApplication_applicationDisplayName uintptr
	addr_QGuiApplication_setBadgeNumber uintptr
	addr_QGuiApplication_setDesktopFileName uintptr
	addr_QGuiApplication_desktopFileName uintptr
	addr_QGuiApplication_allWindows uintptr
	addr_QGuiApplication_topLevelWindows uintptr
	addr_QGuiApplication_topLevelAt uintptr
	addr_QGuiApplication_setWindowIcon uintptr
	addr_QGuiApplication_windowIcon uintptr
	addr_QGuiApplication_platformName uintptr
	addr_QGuiApplication_modalWindow uintptr
	addr_QGuiApplication_focusWindow uintptr
	addr_QGuiApplication_focusObject uintptr
	addr_QGuiApplication_primaryScreen uintptr
	addr_QGuiApplication_screens uintptr
	addr_QGuiApplication_screenAt uintptr
	pg_QGuiApplication_devicePixelRatio func(uintptr) float64
	addr_QGuiApplication_overrideCursor uintptr
	addr_QGuiApplication_setOverrideCursor uintptr
	addr_QGuiApplication_changeOverrideCursor uintptr
	addr_QGuiApplication_restoreOverrideCursor uintptr
	addr_QGuiApplication_font uintptr
	addr_QGuiApplication_setFont uintptr
	addr_QGuiApplication_clipboard uintptr
	addr_QGuiApplication_palette uintptr
	addr_QGuiApplication_setPalette uintptr
	addr_QGuiApplication_keyboardModifiers uintptr
	addr_QGuiApplication_queryKeyboardModifiers uintptr
	addr_QGuiApplication_mouseButtons uintptr
	addr_QGuiApplication_setLayoutDirection uintptr
	addr_QGuiApplication_layoutDirection uintptr
	addr_QGuiApplication_isRightToLeft uintptr
	addr_QGuiApplication_isLeftToRight uintptr
	addr_QGuiApplication_styleHints uintptr
	addr_QGuiApplication_setDesktopSettingsAware uintptr
	addr_QGuiApplication_desktopSettingsAware uintptr
	addr_QGuiApplication_inputMethod uintptr
	addr_QGuiApplication_setQuitOnLastWindowClosed uintptr
	addr_QGuiApplication_quitOnLastWindowClosed uintptr
	addr_QGuiApplication_applicationState uintptr
	addr_QGuiApplication_setHighDpiScaleFactorRoundingPolicy uintptr
	addr_QGuiApplication_highDpiScaleFactorRoundingPolicy uintptr
	addr_QGuiApplication_exec uintptr
	addr_QGuiApplication_isSessionRestored uintptr
	addr_QGuiApplication_sessionId uintptr
	addr_QGuiApplication_sessionKey uintptr
	addr_QGuiApplication_isSavingSession uintptr
	addr_QGuiApplication_sync uintptr
	addr_QGuiApplication_fontDatabaseChanged uintptr
	addr_QGuiApplication_screenAdded uintptr
	addr_QGuiApplication_screenRemoved uintptr
	addr_QGuiApplication_primaryScreenChanged uintptr
	addr_QGuiApplication_lastWindowClosed uintptr
	addr_QGuiApplication_focusObjectChanged uintptr
	addr_QGuiApplication_focusWindowChanged uintptr
	addr_QGuiApplication_applicationStateChanged uintptr
	addr_QGuiApplication_layoutDirectionChanged uintptr
	addr_QGuiApplication_commitDataRequest uintptr
	addr_QGuiApplication_saveStateRequest uintptr
	addr_QGuiApplication_applicationDisplayNameChanged uintptr
	addr_QGuiApplication_paletteChanged uintptr
	addr_QGuiApplication_fontChanged uintptr
	addr_QGuiApplication_arguments uintptr
	addr_QGuiApplication_setAttribute uintptr
	addr_QGuiApplication_testAttribute uintptr
	addr_QGuiApplication_setOrganizationDomain uintptr
	addr_QGuiApplication_organizationDomain uintptr
	addr_QGuiApplication_setOrganizationName uintptr
	addr_QGuiApplication_organizationName uintptr
	addr_QGuiApplication_setApplicationName uintptr
	addr_QGuiApplication_applicationName uintptr
	addr_QGuiApplication_setApplicationVersion uintptr
	addr_QGuiApplication_applicationVersion uintptr
	addr_QGuiApplication_setSetuidAllowed uintptr
	addr_QGuiApplication_isSetuidAllowed uintptr
	addr_QGuiApplication_instance uintptr
	addr_QGuiApplication_instanceExists uintptr
	addr_QGuiApplication_processEvents uintptr
	addr_QGuiApplication_processEvents2 uintptr
	addr_QGuiApplication_processEvents3 uintptr
	addr_QGuiApplication_sendEvent uintptr
	addr_QGuiApplication_postEvent uintptr
	addr_QGuiApplication_sendPostedEvents uintptr
	addr_QGuiApplication_removePostedEvents uintptr
	addr_QGuiApplication_eventDispatcher uintptr
	addr_QGuiApplication_setEventDispatcher uintptr
	addr_QGuiApplication_startingUp uintptr
	addr_QGuiApplication_closingDown uintptr
	addr_QGuiApplication_applicationDirPath uintptr
	addr_QGuiApplication_applicationFilePath uintptr
	addr_QGuiApplication_applicationPid uintptr
	addr_QGuiApplication_setLibraryPaths uintptr
	addr_QGuiApplication_libraryPaths uintptr
	addr_QGuiApplication_addLibraryPath uintptr
	addr_QGuiApplication_removeLibraryPath uintptr
	addr_QGuiApplication_installTranslator uintptr
	addr_QGuiApplication_removeTranslator uintptr
	addr_QGuiApplication_translate uintptr
	addr_QGuiApplication_isQuitLockEnabled uintptr
	addr_QGuiApplication_setQuitLockEnabled uintptr
	addr_QGuiApplication_quit uintptr
	addr_QGuiApplication_exit uintptr
	addr_QGuiApplication_connect uintptr
	addr_QGuiApplication_connect2 uintptr
	addr_QGuiApplication_tr2 uintptr
	addr_QGuiApplication_tr3 uintptr
	addr_QGuiApplication_setAttribute2 uintptr
	addr_QGuiApplication_processEvents4 uintptr
	addr_QGuiApplication_postEvent2 uintptr
	addr_QGuiApplication_sendPostedEvents2 uintptr
	addr_QGuiApplication_sendPostedEvents3 uintptr
	addr_QGuiApplication_removePostedEvents2 uintptr
	addr_QGuiApplication_translate2 uintptr
	addr_QGuiApplication_translate3 uintptr
	addr_QGuiApplication_exit2 uintptr
	addr_QGuiApplication_connect3 uintptr
	addr_QGuiApplication_connect4 uintptr
	pg_QGuiApplication_connect_fontDatabaseChanged func(uintptr, int64)
	pg_QGuiApplication_connect_screenAdded func(uintptr, int64)
	pg_QGuiApplication_connect_screenRemoved func(uintptr, int64)
	pg_QGuiApplication_connect_primaryScreenChanged func(uintptr, int64)
	pg_QGuiApplication_connect_lastWindowClosed func(uintptr, int64)
	pg_QGuiApplication_connect_focusObjectChanged func(uintptr, int64)
	pg_QGuiApplication_connect_focusWindowChanged func(uintptr, int64)
	pg_QGuiApplication_connect_applicationStateChanged func(uintptr, int64)
	pg_QGuiApplication_connect_layoutDirectionChanged func(uintptr, int64)
	pg_QGuiApplication_connect_commitDataRequest func(uintptr, int64)
	pg_QGuiApplication_connect_saveStateRequest func(uintptr, int64)
	pg_QGuiApplication_connect_applicationDisplayNameChanged func(uintptr, int64)
	pg_QGuiApplication_connect_paletteChanged func(uintptr, int64)
	pg_QGuiApplication_connect_fontChanged func(uintptr, int64)
	pg_QGuiApplication_connect_aboutToQuit func(uintptr, int64)
	pg_QGuiApplication_connect_organizationNameChanged func(uintptr, int64)
	pg_QGuiApplication_connect_organizationDomainChanged func(uintptr, int64)
	pg_QGuiApplication_connect_applicationNameChanged func(uintptr, int64)
	pg_QGuiApplication_connect_applicationVersionChanged func(uintptr, int64)
	pg_QGuiApplication_connect_destroyed func(uintptr, int64)
	pg_QGuiApplication_connect_destroyed2 func(uintptr, int64)
	pg_QGuiApplication_connect_objectNameChanged func(uintptr, int64)
	pg_QGuiApplication__disconnect func(int64)
	pg_QGuiApplication_new_args func(uintptr, uintptr) uintptr
	pg_QGuiApplication_delete func(uintptr)
)

func init() {
	symaddr(&addr_QGuiApplication_chain, "QGuiApplication_chain")
	symaddr(&addr_QGuiApplication_tr, "QGuiApplication_tr")
	symaddr(&addr_QGuiApplication_setApplicationDisplayName, "QGuiApplication_setApplicationDisplayName")
	symaddr(&addr_QGuiApplication_applicationDisplayName, "QGuiApplication_applicationDisplayName")
	symaddr(&addr_QGuiApplication_setBadgeNumber, "QGuiApplication_setBadgeNumber")
	symaddr(&addr_QGuiApplication_setDesktopFileName, "QGuiApplication_setDesktopFileName")
	symaddr(&addr_QGuiApplication_desktopFileName, "QGuiApplication_desktopFileName")
	symaddr(&addr_QGuiApplication_allWindows, "QGuiApplication_allWindows")
	symaddr(&addr_QGuiApplication_topLevelWindows, "QGuiApplication_topLevelWindows")
	symaddr(&addr_QGuiApplication_topLevelAt, "QGuiApplication_topLevelAt")
	symaddr(&addr_QGuiApplication_setWindowIcon, "QGuiApplication_setWindowIcon")
	symaddr(&addr_QGuiApplication_windowIcon, "QGuiApplication_windowIcon")
	symaddr(&addr_QGuiApplication_platformName, "QGuiApplication_platformName")
	symaddr(&addr_QGuiApplication_modalWindow, "QGuiApplication_modalWindow")
	symaddr(&addr_QGuiApplication_focusWindow, "QGuiApplication_focusWindow")
	symaddr(&addr_QGuiApplication_focusObject, "QGuiApplication_focusObject")
	symaddr(&addr_QGuiApplication_primaryScreen, "QGuiApplication_primaryScreen")
	symaddr(&addr_QGuiApplication_screens, "QGuiApplication_screens")
	symaddr(&addr_QGuiApplication_screenAt, "QGuiApplication_screenAt")
	reg(&pg_QGuiApplication_devicePixelRatio, "QGuiApplication_devicePixelRatio")
	symaddr(&addr_QGuiApplication_overrideCursor, "QGuiApplication_overrideCursor")
	symaddr(&addr_QGuiApplication_setOverrideCursor, "QGuiApplication_setOverrideCursor")
	symaddr(&addr_QGuiApplication_changeOverrideCursor, "QGuiApplication_changeOverrideCursor")
	symaddr(&addr_QGuiApplication_restoreOverrideCursor, "QGuiApplication_restoreOverrideCursor")
	symaddr(&addr_QGuiApplication_font, "QGuiApplication_font")
	symaddr(&addr_QGuiApplication_setFont, "QGuiApplication_setFont")
	symaddr(&addr_QGuiApplication_clipboard, "QGuiApplication_clipboard")
	symaddr(&addr_QGuiApplication_palette, "QGuiApplication_palette")
	symaddr(&addr_QGuiApplication_setPalette, "QGuiApplication_setPalette")
	symaddr(&addr_QGuiApplication_keyboardModifiers, "QGuiApplication_keyboardModifiers")
	symaddr(&addr_QGuiApplication_queryKeyboardModifiers, "QGuiApplication_queryKeyboardModifiers")
	symaddr(&addr_QGuiApplication_mouseButtons, "QGuiApplication_mouseButtons")
	symaddr(&addr_QGuiApplication_setLayoutDirection, "QGuiApplication_setLayoutDirection")
	symaddr(&addr_QGuiApplication_layoutDirection, "QGuiApplication_layoutDirection")
	symaddr(&addr_QGuiApplication_isRightToLeft, "QGuiApplication_isRightToLeft")
	symaddr(&addr_QGuiApplication_isLeftToRight, "QGuiApplication_isLeftToRight")
	symaddr(&addr_QGuiApplication_styleHints, "QGuiApplication_styleHints")
	symaddr(&addr_QGuiApplication_setDesktopSettingsAware, "QGuiApplication_setDesktopSettingsAware")
	symaddr(&addr_QGuiApplication_desktopSettingsAware, "QGuiApplication_desktopSettingsAware")
	symaddr(&addr_QGuiApplication_inputMethod, "QGuiApplication_inputMethod")
	symaddr(&addr_QGuiApplication_setQuitOnLastWindowClosed, "QGuiApplication_setQuitOnLastWindowClosed")
	symaddr(&addr_QGuiApplication_quitOnLastWindowClosed, "QGuiApplication_quitOnLastWindowClosed")
	symaddr(&addr_QGuiApplication_applicationState, "QGuiApplication_applicationState")
	symaddr(&addr_QGuiApplication_setHighDpiScaleFactorRoundingPolicy, "QGuiApplication_setHighDpiScaleFactorRoundingPolicy")
	symaddr(&addr_QGuiApplication_highDpiScaleFactorRoundingPolicy, "QGuiApplication_highDpiScaleFactorRoundingPolicy")
	symaddr(&addr_QGuiApplication_exec, "QGuiApplication_exec")
	symaddr(&addr_QGuiApplication_isSessionRestored, "QGuiApplication_isSessionRestored")
	symaddr(&addr_QGuiApplication_sessionId, "QGuiApplication_sessionId")
	symaddr(&addr_QGuiApplication_sessionKey, "QGuiApplication_sessionKey")
	symaddr(&addr_QGuiApplication_isSavingSession, "QGuiApplication_isSavingSession")
	symaddr(&addr_QGuiApplication_sync, "QGuiApplication_sync")
	symaddr(&addr_QGuiApplication_fontDatabaseChanged, "QGuiApplication_fontDatabaseChanged")
	symaddr(&addr_QGuiApplication_screenAdded, "QGuiApplication_screenAdded")
	symaddr(&addr_QGuiApplication_screenRemoved, "QGuiApplication_screenRemoved")
	symaddr(&addr_QGuiApplication_primaryScreenChanged, "QGuiApplication_primaryScreenChanged")
	symaddr(&addr_QGuiApplication_lastWindowClosed, "QGuiApplication_lastWindowClosed")
	symaddr(&addr_QGuiApplication_focusObjectChanged, "QGuiApplication_focusObjectChanged")
	symaddr(&addr_QGuiApplication_focusWindowChanged, "QGuiApplication_focusWindowChanged")
	symaddr(&addr_QGuiApplication_applicationStateChanged, "QGuiApplication_applicationStateChanged")
	symaddr(&addr_QGuiApplication_layoutDirectionChanged, "QGuiApplication_layoutDirectionChanged")
	symaddr(&addr_QGuiApplication_commitDataRequest, "QGuiApplication_commitDataRequest")
	symaddr(&addr_QGuiApplication_saveStateRequest, "QGuiApplication_saveStateRequest")
	symaddr(&addr_QGuiApplication_applicationDisplayNameChanged, "QGuiApplication_applicationDisplayNameChanged")
	symaddr(&addr_QGuiApplication_paletteChanged, "QGuiApplication_paletteChanged")
	symaddr(&addr_QGuiApplication_fontChanged, "QGuiApplication_fontChanged")
	symaddr(&addr_QGuiApplication_arguments, "QGuiApplication_arguments")
	symaddr(&addr_QGuiApplication_setAttribute, "QGuiApplication_setAttribute")
	symaddr(&addr_QGuiApplication_testAttribute, "QGuiApplication_testAttribute")
	symaddr(&addr_QGuiApplication_setOrganizationDomain, "QGuiApplication_setOrganizationDomain")
	symaddr(&addr_QGuiApplication_organizationDomain, "QGuiApplication_organizationDomain")
	symaddr(&addr_QGuiApplication_setOrganizationName, "QGuiApplication_setOrganizationName")
	symaddr(&addr_QGuiApplication_organizationName, "QGuiApplication_organizationName")
	symaddr(&addr_QGuiApplication_setApplicationName, "QGuiApplication_setApplicationName")
	symaddr(&addr_QGuiApplication_applicationName, "QGuiApplication_applicationName")
	symaddr(&addr_QGuiApplication_setApplicationVersion, "QGuiApplication_setApplicationVersion")
	symaddr(&addr_QGuiApplication_applicationVersion, "QGuiApplication_applicationVersion")
	symaddr(&addr_QGuiApplication_setSetuidAllowed, "QGuiApplication_setSetuidAllowed")
	symaddr(&addr_QGuiApplication_isSetuidAllowed, "QGuiApplication_isSetuidAllowed")
	symaddr(&addr_QGuiApplication_instance, "QGuiApplication_instance")
	symaddr(&addr_QGuiApplication_instanceExists, "QGuiApplication_instanceExists")
	symaddr(&addr_QGuiApplication_processEvents, "QGuiApplication_processEvents")
	symaddr(&addr_QGuiApplication_processEvents2, "QGuiApplication_processEvents2")
	symaddr(&addr_QGuiApplication_processEvents3, "QGuiApplication_processEvents3")
	symaddr(&addr_QGuiApplication_sendEvent, "QGuiApplication_sendEvent")
	symaddr(&addr_QGuiApplication_postEvent, "QGuiApplication_postEvent")
	symaddr(&addr_QGuiApplication_sendPostedEvents, "QGuiApplication_sendPostedEvents")
	symaddr(&addr_QGuiApplication_removePostedEvents, "QGuiApplication_removePostedEvents")
	symaddr(&addr_QGuiApplication_eventDispatcher, "QGuiApplication_eventDispatcher")
	symaddr(&addr_QGuiApplication_setEventDispatcher, "QGuiApplication_setEventDispatcher")
	symaddr(&addr_QGuiApplication_startingUp, "QGuiApplication_startingUp")
	symaddr(&addr_QGuiApplication_closingDown, "QGuiApplication_closingDown")
	symaddr(&addr_QGuiApplication_applicationDirPath, "QGuiApplication_applicationDirPath")
	symaddr(&addr_QGuiApplication_applicationFilePath, "QGuiApplication_applicationFilePath")
	symaddr(&addr_QGuiApplication_applicationPid, "QGuiApplication_applicationPid")
	symaddr(&addr_QGuiApplication_setLibraryPaths, "QGuiApplication_setLibraryPaths")
	symaddr(&addr_QGuiApplication_libraryPaths, "QGuiApplication_libraryPaths")
	symaddr(&addr_QGuiApplication_addLibraryPath, "QGuiApplication_addLibraryPath")
	symaddr(&addr_QGuiApplication_removeLibraryPath, "QGuiApplication_removeLibraryPath")
	symaddr(&addr_QGuiApplication_installTranslator, "QGuiApplication_installTranslator")
	symaddr(&addr_QGuiApplication_removeTranslator, "QGuiApplication_removeTranslator")
	symaddr(&addr_QGuiApplication_translate, "QGuiApplication_translate")
	symaddr(&addr_QGuiApplication_isQuitLockEnabled, "QGuiApplication_isQuitLockEnabled")
	symaddr(&addr_QGuiApplication_setQuitLockEnabled, "QGuiApplication_setQuitLockEnabled")
	symaddr(&addr_QGuiApplication_quit, "QGuiApplication_quit")
	symaddr(&addr_QGuiApplication_exit, "QGuiApplication_exit")
	symaddr(&addr_QGuiApplication_connect, "QGuiApplication_connect")
	symaddr(&addr_QGuiApplication_connect2, "QGuiApplication_connect2")
	symaddr(&addr_QGuiApplication_tr2, "QGuiApplication_tr2")
	symaddr(&addr_QGuiApplication_tr3, "QGuiApplication_tr3")
	symaddr(&addr_QGuiApplication_setAttribute2, "QGuiApplication_setAttribute2")
	symaddr(&addr_QGuiApplication_processEvents4, "QGuiApplication_processEvents4")
	symaddr(&addr_QGuiApplication_postEvent2, "QGuiApplication_postEvent2")
	symaddr(&addr_QGuiApplication_sendPostedEvents2, "QGuiApplication_sendPostedEvents2")
	symaddr(&addr_QGuiApplication_sendPostedEvents3, "QGuiApplication_sendPostedEvents3")
	symaddr(&addr_QGuiApplication_removePostedEvents2, "QGuiApplication_removePostedEvents2")
	symaddr(&addr_QGuiApplication_translate2, "QGuiApplication_translate2")
	symaddr(&addr_QGuiApplication_translate3, "QGuiApplication_translate3")
	symaddr(&addr_QGuiApplication_exit2, "QGuiApplication_exit2")
	symaddr(&addr_QGuiApplication_connect3, "QGuiApplication_connect3")
	symaddr(&addr_QGuiApplication_connect4, "QGuiApplication_connect4")
	reg(&pg_QGuiApplication_connect_fontDatabaseChanged, "QGuiApplication_connect_fontDatabaseChanged")
	reg(&pg_QGuiApplication_connect_screenAdded, "QGuiApplication_connect_screenAdded")
	reg(&pg_QGuiApplication_connect_screenRemoved, "QGuiApplication_connect_screenRemoved")
	reg(&pg_QGuiApplication_connect_primaryScreenChanged, "QGuiApplication_connect_primaryScreenChanged")
	reg(&pg_QGuiApplication_connect_lastWindowClosed, "QGuiApplication_connect_lastWindowClosed")
	reg(&pg_QGuiApplication_connect_focusObjectChanged, "QGuiApplication_connect_focusObjectChanged")
	reg(&pg_QGuiApplication_connect_focusWindowChanged, "QGuiApplication_connect_focusWindowChanged")
	reg(&pg_QGuiApplication_connect_applicationStateChanged, "QGuiApplication_connect_applicationStateChanged")
	reg(&pg_QGuiApplication_connect_layoutDirectionChanged, "QGuiApplication_connect_layoutDirectionChanged")
	reg(&pg_QGuiApplication_connect_commitDataRequest, "QGuiApplication_connect_commitDataRequest")
	reg(&pg_QGuiApplication_connect_saveStateRequest, "QGuiApplication_connect_saveStateRequest")
	reg(&pg_QGuiApplication_connect_applicationDisplayNameChanged, "QGuiApplication_connect_applicationDisplayNameChanged")
	reg(&pg_QGuiApplication_connect_paletteChanged, "QGuiApplication_connect_paletteChanged")
	reg(&pg_QGuiApplication_connect_fontChanged, "QGuiApplication_connect_fontChanged")
	reg(&pg_QGuiApplication_connect_aboutToQuit, "QGuiApplication_connect_aboutToQuit")
	reg(&pg_QGuiApplication_connect_organizationNameChanged, "QGuiApplication_connect_organizationNameChanged")
	reg(&pg_QGuiApplication_connect_organizationDomainChanged, "QGuiApplication_connect_organizationDomainChanged")
	reg(&pg_QGuiApplication_connect_applicationNameChanged, "QGuiApplication_connect_applicationNameChanged")
	reg(&pg_QGuiApplication_connect_applicationVersionChanged, "QGuiApplication_connect_applicationVersionChanged")
	reg(&pg_QGuiApplication_connect_destroyed, "QGuiApplication_connect_destroyed")
	reg(&pg_QGuiApplication_connect_destroyed2, "QGuiApplication_connect_destroyed2")
	reg(&pg_QGuiApplication_connect_objectNameChanged, "QGuiApplication_connect_objectNameChanged")
	reg(&pg_QGuiApplication__disconnect, "QGuiApplication__disconnect")
	reg(&pg_QGuiApplication_new_args, "QGuiApplication_new_args")
	reg(&pg_QGuiApplication_delete, "QGuiApplication_delete")
}

type QGuiApplication struct {
	h uintptr
	*QCoreApplication
}

func bindQGuiApplication(h uintptr) *QGuiApplication {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QGuiApplication_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QGuiApplication{h: h, QCoreApplication: &QCoreApplication{h: __c[0], QObject: &QObject{h: __c[1]}}}
}

func newownedQGuiApplication(h uintptr) *QGuiApplication { return bindQGuiApplication(h) }

func (this *QGuiApplication) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QGuiApplication) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQGuiApplication(h unsafe.Pointer) *QGuiApplication { return bindQGuiApplication(uintptr(h)) }
func UnsafeNewQGuiApplicationOwned(h unsafe.Pointer) *QGuiApplication { return newownedQGuiApplication(uintptr(h)) }

func (this *QGuiApplication) UnsafeRelease() {
	this.h = 0
	if this.QCoreApplication != nil {
		this.QCoreApplication.UnsafeRelease()
	}
}

func QGuiApplication_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGuiApplication_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func QGuiApplication_SetApplicationDisplayName(name string) {
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	purego.SyscallN(addr_QGuiApplication_setApplicationDisplayName, name_p, uintptr(len(name)))
	runtime.KeepAlive(name)
}

func QGuiApplication_ApplicationDisplayName() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGuiApplication_applicationDisplayName, uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func (this *QGuiApplication) SetBadgeNumber(number int64) {
	purego.SyscallN(addr_QGuiApplication_setBadgeNumber, this.h, uintptr(number))
	runtime.KeepAlive(this)
}

func QGuiApplication_SetDesktopFileName(name string) {
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	purego.SyscallN(addr_QGuiApplication_setDesktopFileName, name_p, uintptr(len(name)))
	runtime.KeepAlive(name)
}

func QGuiApplication_DesktopFileName() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGuiApplication_desktopFileName, uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QGuiApplication_AllWindows() []*QWindow {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGuiApplication_allWindows, uintptr(unsafe.Pointer(&__n)))
	__out := make([]*QWindow, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = bindQWindow(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QGuiApplication_TopLevelWindows() []*QWindow {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGuiApplication_topLevelWindows, uintptr(unsafe.Pointer(&__n)))
	__out := make([]*QWindow, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = bindQWindow(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QGuiApplication_TopLevelAt(pos *QPoint) *QWindow {
	__r, _, _ := purego.SyscallN(addr_QGuiApplication_topLevelAt, uintptr(pos.UnsafePointer()))
	runtime.KeepAlive(pos)
	return bindQWindow(__r)
}

func QGuiApplication_SetWindowIcon(icon *QIcon) {
	purego.SyscallN(addr_QGuiApplication_setWindowIcon, uintptr(icon.UnsafePointer()))
	runtime.KeepAlive(icon)
}

func QGuiApplication_WindowIcon() *QIcon {
	__r, _, _ := purego.SyscallN(addr_QGuiApplication_windowIcon)
	return newownedQIcon(__r)
}

func QGuiApplication_PlatformName() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGuiApplication_platformName, uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QGuiApplication_ModalWindow() *QWindow {
	__r, _, _ := purego.SyscallN(addr_QGuiApplication_modalWindow)
	return bindQWindow(__r)
}

func QGuiApplication_FocusWindow() *QWindow {
	__r, _, _ := purego.SyscallN(addr_QGuiApplication_focusWindow)
	return bindQWindow(__r)
}

func QGuiApplication_FocusObject() *QObject {
	__r, _, _ := purego.SyscallN(addr_QGuiApplication_focusObject)
	return bindQObject(__r)
}

func QGuiApplication_PrimaryScreen() *QScreen {
	__r, _, _ := purego.SyscallN(addr_QGuiApplication_primaryScreen)
	return bindQScreen(__r)
}

func QGuiApplication_Screens() []*QScreen {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGuiApplication_screens, uintptr(unsafe.Pointer(&__n)))
	__out := make([]*QScreen, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = bindQScreen(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QGuiApplication_ScreenAt(point *QPoint) *QScreen {
	__r, _, _ := purego.SyscallN(addr_QGuiApplication_screenAt, uintptr(point.UnsafePointer()))
	runtime.KeepAlive(point)
	return bindQScreen(__r)
}

func (this *QGuiApplication) DevicePixelRatio() float64 {
	__r := pg_QGuiApplication_devicePixelRatio(this.h)
	runtime.KeepAlive(this)
	return __r
}

func QGuiApplication_OverrideCursor() *QCursor {
	__r, _, _ := purego.SyscallN(addr_QGuiApplication_overrideCursor)
	return bindQCursor(__r)
}

func QGuiApplication_SetOverrideCursor(arg1 *QCursor) {
	purego.SyscallN(addr_QGuiApplication_setOverrideCursor, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(arg1)
}

func QGuiApplication_ChangeOverrideCursor(arg1 *QCursor) {
	purego.SyscallN(addr_QGuiApplication_changeOverrideCursor, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(arg1)
}

func QGuiApplication_RestoreOverrideCursor() {
	purego.SyscallN(addr_QGuiApplication_restoreOverrideCursor)
}

func QGuiApplication_Font() *QFont {
	__r, _, _ := purego.SyscallN(addr_QGuiApplication_font)
	return newownedQFont(__r)
}

func QGuiApplication_SetFont(arg1 *QFont) {
	purego.SyscallN(addr_QGuiApplication_setFont, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(arg1)
}

func QGuiApplication_Clipboard() *QClipboard {
	__r, _, _ := purego.SyscallN(addr_QGuiApplication_clipboard)
	return bindQClipboard(__r)
}

func QGuiApplication_Palette() *QPalette {
	__r, _, _ := purego.SyscallN(addr_QGuiApplication_palette)
	return newownedQPalette(__r)
}

func QGuiApplication_SetPalette(pal *QPalette) {
	purego.SyscallN(addr_QGuiApplication_setPalette, uintptr(pal.UnsafePointer()))
	runtime.KeepAlive(pal)
}

func QGuiApplication_KeyboardModifiers() KeyboardModifier {
	__r0, _, _ := purego.SyscallN(addr_QGuiApplication_keyboardModifiers)
	__r := int32(uint32(__r0))
	return KeyboardModifier(__r)
}

func QGuiApplication_QueryKeyboardModifiers() KeyboardModifier {
	__r0, _, _ := purego.SyscallN(addr_QGuiApplication_queryKeyboardModifiers)
	__r := int32(uint32(__r0))
	return KeyboardModifier(__r)
}

func QGuiApplication_MouseButtons() MouseButton {
	__r0, _, _ := purego.SyscallN(addr_QGuiApplication_mouseButtons)
	__r := int32(uint32(__r0))
	return MouseButton(__r)
}

func QGuiApplication_SetLayoutDirection(direction LayoutDirection) {
	purego.SyscallN(addr_QGuiApplication_setLayoutDirection, uintptr(direction))
}

func QGuiApplication_LayoutDirection() LayoutDirection {
	__r0, _, _ := purego.SyscallN(addr_QGuiApplication_layoutDirection)
	__r := int32(uint32(__r0))
	return LayoutDirection(__r)
}

func QGuiApplication_IsRightToLeft() bool {
	__r0, _, _ := purego.SyscallN(addr_QGuiApplication_isRightToLeft)
	__r := int32(uint32(__r0))
	return __r != 0
}

func QGuiApplication_IsLeftToRight() bool {
	__r0, _, _ := purego.SyscallN(addr_QGuiApplication_isLeftToRight)
	__r := int32(uint32(__r0))
	return __r != 0
}

func QGuiApplication_StyleHints() *QStyleHints {
	__r, _, _ := purego.SyscallN(addr_QGuiApplication_styleHints)
	return bindQStyleHints(__r)
}

func QGuiApplication_SetDesktopSettingsAware(on bool) {
	purego.SyscallN(addr_QGuiApplication_setDesktopSettingsAware, uintptr(b2i(on)))
}

func QGuiApplication_DesktopSettingsAware() bool {
	__r0, _, _ := purego.SyscallN(addr_QGuiApplication_desktopSettingsAware)
	__r := int32(uint32(__r0))
	return __r != 0
}

func QGuiApplication_InputMethod() *QInputMethod {
	__r, _, _ := purego.SyscallN(addr_QGuiApplication_inputMethod)
	return bindQInputMethod(__r)
}

func QGuiApplication_SetQuitOnLastWindowClosed(quit bool) {
	purego.SyscallN(addr_QGuiApplication_setQuitOnLastWindowClosed, uintptr(b2i(quit)))
}

func QGuiApplication_QuitOnLastWindowClosed() bool {
	__r0, _, _ := purego.SyscallN(addr_QGuiApplication_quitOnLastWindowClosed)
	__r := int32(uint32(__r0))
	return __r != 0
}

func QGuiApplication_ApplicationState() ApplicationState {
	__r0, _, _ := purego.SyscallN(addr_QGuiApplication_applicationState)
	__r := int32(uint32(__r0))
	return ApplicationState(__r)
}

func QGuiApplication_SetHighDpiScaleFactorRoundingPolicy(policy HighDpiScaleFactorRoundingPolicy) {
	purego.SyscallN(addr_QGuiApplication_setHighDpiScaleFactorRoundingPolicy, uintptr(policy))
}

func QGuiApplication_HighDpiScaleFactorRoundingPolicy() HighDpiScaleFactorRoundingPolicy {
	__r0, _, _ := purego.SyscallN(addr_QGuiApplication_highDpiScaleFactorRoundingPolicy)
	__r := int32(uint32(__r0))
	return HighDpiScaleFactorRoundingPolicy(__r)
}

func QGuiApplication_Exec() int {
	__r0, _, _ := purego.SyscallN(addr_QGuiApplication_exec)
	__r := int32(uint32(__r0))
	return int(__r)
}

func (this *QGuiApplication) IsSessionRestored() bool {
	__r0, _, _ := purego.SyscallN(addr_QGuiApplication_isSessionRestored, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGuiApplication) SessionId() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGuiApplication_sessionId, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QGuiApplication) SessionKey() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGuiApplication_sessionKey, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QGuiApplication) IsSavingSession() bool {
	__r0, _, _ := purego.SyscallN(addr_QGuiApplication_isSavingSession, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QGuiApplication_Sync() {
	purego.SyscallN(addr_QGuiApplication_sync)
}

func (this *QGuiApplication) FontDatabaseChanged() {
	purego.SyscallN(addr_QGuiApplication_fontDatabaseChanged, this.h)
	runtime.KeepAlive(this)
}

func (this *QGuiApplication) ScreenAdded(screen *QScreen) {
	purego.SyscallN(addr_QGuiApplication_screenAdded, this.h, uintptr(screen.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(screen)
}

func (this *QGuiApplication) ScreenRemoved(screen *QScreen) {
	purego.SyscallN(addr_QGuiApplication_screenRemoved, this.h, uintptr(screen.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(screen)
}

func (this *QGuiApplication) PrimaryScreenChanged(screen *QScreen) {
	purego.SyscallN(addr_QGuiApplication_primaryScreenChanged, this.h, uintptr(screen.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(screen)
}

func (this *QGuiApplication) LastWindowClosed() {
	purego.SyscallN(addr_QGuiApplication_lastWindowClosed, this.h)
	runtime.KeepAlive(this)
}

func (this *QGuiApplication) FocusObjectChanged(focusObject *QObject) {
	purego.SyscallN(addr_QGuiApplication_focusObjectChanged, this.h, uintptr(focusObject.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(focusObject)
}

func (this *QGuiApplication) FocusWindowChanged(focusWindow *QWindow) {
	purego.SyscallN(addr_QGuiApplication_focusWindowChanged, this.h, uintptr(focusWindow.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(focusWindow)
}

func (this *QGuiApplication) ApplicationStateChanged(state ApplicationState) {
	purego.SyscallN(addr_QGuiApplication_applicationStateChanged, this.h, uintptr(state))
	runtime.KeepAlive(this)
}

func (this *QGuiApplication) LayoutDirectionChanged(direction LayoutDirection) {
	purego.SyscallN(addr_QGuiApplication_layoutDirectionChanged, this.h, uintptr(direction))
	runtime.KeepAlive(this)
}

func (this *QGuiApplication) CommitDataRequest(sessionManager *QSessionManager) {
	purego.SyscallN(addr_QGuiApplication_commitDataRequest, this.h, uintptr(sessionManager.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sessionManager)
}

func (this *QGuiApplication) SaveStateRequest(sessionManager *QSessionManager) {
	purego.SyscallN(addr_QGuiApplication_saveStateRequest, this.h, uintptr(sessionManager.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sessionManager)
}

func (this *QGuiApplication) ApplicationDisplayNameChanged() {
	purego.SyscallN(addr_QGuiApplication_applicationDisplayNameChanged, this.h)
	runtime.KeepAlive(this)
}

func (this *QGuiApplication) PaletteChanged(pal *QPalette) {
	purego.SyscallN(addr_QGuiApplication_paletteChanged, this.h, uintptr(pal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pal)
}

func (this *QGuiApplication) FontChanged(font *QFont) {
	purego.SyscallN(addr_QGuiApplication_fontChanged, this.h, uintptr(font.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(font)
}

func QGuiApplication_Arguments() []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGuiApplication_arguments, uintptr(unsafe.Pointer(&__n)))
	__out := make([]string, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = pgGoStr(__sl[__i])
		pgFree(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QGuiApplication_SetAttribute(attribute ApplicationAttribute) {
	purego.SyscallN(addr_QGuiApplication_setAttribute, uintptr(attribute))
}

func QGuiApplication_TestAttribute(attribute ApplicationAttribute) bool {
	__r0, _, _ := purego.SyscallN(addr_QGuiApplication_testAttribute, uintptr(attribute))
	__r := int32(uint32(__r0))
	return __r != 0
}

func QGuiApplication_SetOrganizationDomain(orgDomain string) {
	var orgDomain_p uintptr
	if len(orgDomain) > 0 {
		orgDomain_p = uintptr(unsafe.Pointer(unsafe.StringData(orgDomain)))
	}
	purego.SyscallN(addr_QGuiApplication_setOrganizationDomain, orgDomain_p, uintptr(len(orgDomain)))
	runtime.KeepAlive(orgDomain)
}

func QGuiApplication_OrganizationDomain() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGuiApplication_organizationDomain, uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QGuiApplication_SetOrganizationName(orgName string) {
	var orgName_p uintptr
	if len(orgName) > 0 {
		orgName_p = uintptr(unsafe.Pointer(unsafe.StringData(orgName)))
	}
	purego.SyscallN(addr_QGuiApplication_setOrganizationName, orgName_p, uintptr(len(orgName)))
	runtime.KeepAlive(orgName)
}

func QGuiApplication_OrganizationName() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGuiApplication_organizationName, uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QGuiApplication_SetApplicationName(application string) {
	var application_p uintptr
	if len(application) > 0 {
		application_p = uintptr(unsafe.Pointer(unsafe.StringData(application)))
	}
	purego.SyscallN(addr_QGuiApplication_setApplicationName, application_p, uintptr(len(application)))
	runtime.KeepAlive(application)
}

func QGuiApplication_ApplicationName() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGuiApplication_applicationName, uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QGuiApplication_SetApplicationVersion(version string) {
	var version_p uintptr
	if len(version) > 0 {
		version_p = uintptr(unsafe.Pointer(unsafe.StringData(version)))
	}
	purego.SyscallN(addr_QGuiApplication_setApplicationVersion, version_p, uintptr(len(version)))
	runtime.KeepAlive(version)
}

func QGuiApplication_ApplicationVersion() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGuiApplication_applicationVersion, uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QGuiApplication_SetSetuidAllowed(allow bool) {
	purego.SyscallN(addr_QGuiApplication_setSetuidAllowed, uintptr(b2i(allow)))
}

func QGuiApplication_IsSetuidAllowed() bool {
	__r0, _, _ := purego.SyscallN(addr_QGuiApplication_isSetuidAllowed)
	__r := int32(uint32(__r0))
	return __r != 0
}

func QGuiApplication_Instance() *QCoreApplication {
	__r, _, _ := purego.SyscallN(addr_QGuiApplication_instance)
	return bindQCoreApplication(__r)
}

func QGuiApplication_InstanceExists() bool {
	__r0, _, _ := purego.SyscallN(addr_QGuiApplication_instanceExists)
	__r := int32(uint32(__r0))
	return __r != 0
}

func QGuiApplication_ProcessEvents() {
	purego.SyscallN(addr_QGuiApplication_processEvents)
}

func QGuiApplication_ProcessEvents2(flags QEventLoop__ProcessEventsFlag, maxtime int) {
	purego.SyscallN(addr_QGuiApplication_processEvents2, uintptr(flags), uintptr(maxtime))
}

func QGuiApplication_ProcessEvents3(flags QEventLoop__ProcessEventsFlag, deadline *QDeadlineTimer) {
	purego.SyscallN(addr_QGuiApplication_processEvents3, uintptr(flags), uintptr(deadline.UnsafePointer()))
	runtime.KeepAlive(deadline)
}

func QGuiApplication_SendEvent(receiver *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QGuiApplication_sendEvent, uintptr(receiver.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(event)
	return __r != 0
}

func QGuiApplication_PostEvent(receiver *QObject, event *QEvent) {
	purego.SyscallN(addr_QGuiApplication_postEvent, uintptr(receiver.UnsafePointer()), uintptr(event.UnsafePointer()))
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(event)
}

func QGuiApplication_SendPostedEvents() {
	purego.SyscallN(addr_QGuiApplication_sendPostedEvents)
}

func QGuiApplication_RemovePostedEvents(receiver *QObject) {
	purego.SyscallN(addr_QGuiApplication_removePostedEvents, uintptr(receiver.UnsafePointer()))
	runtime.KeepAlive(receiver)
}

func QGuiApplication_EventDispatcher() *QAbstractEventDispatcher {
	__r, _, _ := purego.SyscallN(addr_QGuiApplication_eventDispatcher)
	return bindQAbstractEventDispatcher(__r)
}

func QGuiApplication_SetEventDispatcher(eventDispatcher *QAbstractEventDispatcher) {
	purego.SyscallN(addr_QGuiApplication_setEventDispatcher, uintptr(eventDispatcher.UnsafePointer()))
	runtime.KeepAlive(eventDispatcher)
}

func QGuiApplication_StartingUp() bool {
	__r0, _, _ := purego.SyscallN(addr_QGuiApplication_startingUp)
	__r := int32(uint32(__r0))
	return __r != 0
}

func QGuiApplication_ClosingDown() bool {
	__r0, _, _ := purego.SyscallN(addr_QGuiApplication_closingDown)
	__r := int32(uint32(__r0))
	return __r != 0
}

func QGuiApplication_ApplicationDirPath() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGuiApplication_applicationDirPath, uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QGuiApplication_ApplicationFilePath() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGuiApplication_applicationFilePath, uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QGuiApplication_ApplicationPid() int64 {
	__r0, _, _ := purego.SyscallN(addr_QGuiApplication_applicationPid)
	__r := int64(__r0)
	return __r
}

func QGuiApplication_SetLibraryPaths(arg1 []string) {
	arg1_arr := make([]uintptr, len(arg1))
	for __i, __e := range arg1 {
		arg1_arr[__i] = pgCString(__e)
	}
	var arg1_p uintptr
	if len(arg1_arr) > 0 {
		arg1_p = uintptr(unsafe.Pointer(&arg1_arr[0]))
	}
	purego.SyscallN(addr_QGuiApplication_setLibraryPaths, arg1_p, uintptr(len(arg1)))
	for __i := range arg1_arr {
		pgFree(arg1_arr[__i])
	}
	runtime.KeepAlive(arg1_arr)
}

func QGuiApplication_LibraryPaths() []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGuiApplication_libraryPaths, uintptr(unsafe.Pointer(&__n)))
	__out := make([]string, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = pgGoStr(__sl[__i])
		pgFree(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QGuiApplication_AddLibraryPath(arg1 string) {
	var arg1_p uintptr
	if len(arg1) > 0 {
		arg1_p = uintptr(unsafe.Pointer(unsafe.StringData(arg1)))
	}
	purego.SyscallN(addr_QGuiApplication_addLibraryPath, arg1_p, uintptr(len(arg1)))
	runtime.KeepAlive(arg1)
}

func QGuiApplication_RemoveLibraryPath(arg1 string) {
	var arg1_p uintptr
	if len(arg1) > 0 {
		arg1_p = uintptr(unsafe.Pointer(unsafe.StringData(arg1)))
	}
	purego.SyscallN(addr_QGuiApplication_removeLibraryPath, arg1_p, uintptr(len(arg1)))
	runtime.KeepAlive(arg1)
}

func QGuiApplication_InstallTranslator(messageFile *QTranslator) bool {
	__r0, _, _ := purego.SyscallN(addr_QGuiApplication_installTranslator, uintptr(messageFile.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(messageFile)
	return __r != 0
}

func QGuiApplication_RemoveTranslator(messageFile *QTranslator) bool {
	__r0, _, _ := purego.SyscallN(addr_QGuiApplication_removeTranslator, uintptr(messageFile.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(messageFile)
	return __r != 0
}

func QGuiApplication_Translate(context string, key string) string {
	context_b := append([]byte(context), 0)
	key_b := append([]byte(key), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGuiApplication_translate, uintptr(unsafe.Pointer(&context_b[0])), uintptr(unsafe.Pointer(&key_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(context_b)
	runtime.KeepAlive(key_b)
	return pgStr(__r, __n)
}

func QGuiApplication_IsQuitLockEnabled() bool {
	__r0, _, _ := purego.SyscallN(addr_QGuiApplication_isQuitLockEnabled)
	__r := int32(uint32(__r0))
	return __r != 0
}

func QGuiApplication_SetQuitLockEnabled(enabled bool) {
	purego.SyscallN(addr_QGuiApplication_setQuitLockEnabled, uintptr(b2i(enabled)))
}

func QGuiApplication_Quit() {
	purego.SyscallN(addr_QGuiApplication_quit)
}

func QGuiApplication_Exit() {
	purego.SyscallN(addr_QGuiApplication_exit)
}

func QGuiApplication_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QGuiApplication_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QGuiApplication_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QGuiApplication_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func QGuiApplication_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGuiApplication_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QGuiApplication_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGuiApplication_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QGuiApplication_SetAttribute2(attribute ApplicationAttribute, on bool) {
	purego.SyscallN(addr_QGuiApplication_setAttribute2, uintptr(attribute), uintptr(b2i(on)))
}

func QGuiApplication_ProcessEventsWithFlags(flags QEventLoop__ProcessEventsFlag) {
	purego.SyscallN(addr_QGuiApplication_processEvents4, uintptr(flags))
}

func QGuiApplication_PostEvent2(receiver *QObject, event *QEvent, priority int) {
	purego.SyscallN(addr_QGuiApplication_postEvent2, uintptr(receiver.UnsafePointer()), uintptr(event.UnsafePointer()), uintptr(priority))
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(event)
}

func QGuiApplication_SendPostedEventsWithReceiver(receiver *QObject) {
	purego.SyscallN(addr_QGuiApplication_sendPostedEvents2, uintptr(receiver.UnsafePointer()))
	runtime.KeepAlive(receiver)
}

func QGuiApplication_SendPostedEvents2(receiver *QObject, event_type int) {
	purego.SyscallN(addr_QGuiApplication_sendPostedEvents3, uintptr(receiver.UnsafePointer()), uintptr(event_type))
	runtime.KeepAlive(receiver)
}

func QGuiApplication_RemovePostedEvents2(receiver *QObject, eventType int) {
	purego.SyscallN(addr_QGuiApplication_removePostedEvents2, uintptr(receiver.UnsafePointer()), uintptr(eventType))
	runtime.KeepAlive(receiver)
}

func QGuiApplication_Translate2(context string, key string, disambiguation string) string {
	context_b := append([]byte(context), 0)
	key_b := append([]byte(key), 0)
	disambiguation_b := append([]byte(disambiguation), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGuiApplication_translate2, uintptr(unsafe.Pointer(&context_b[0])), uintptr(unsafe.Pointer(&key_b[0])), uintptr(unsafe.Pointer(&disambiguation_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(context_b)
	runtime.KeepAlive(key_b)
	runtime.KeepAlive(disambiguation_b)
	return pgStr(__r, __n)
}

func QGuiApplication_Translate3(context string, key string, disambiguation string, n int) string {
	context_b := append([]byte(context), 0)
	key_b := append([]byte(key), 0)
	disambiguation_b := append([]byte(disambiguation), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGuiApplication_translate3, uintptr(unsafe.Pointer(&context_b[0])), uintptr(unsafe.Pointer(&key_b[0])), uintptr(unsafe.Pointer(&disambiguation_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(context_b)
	runtime.KeepAlive(key_b)
	runtime.KeepAlive(disambiguation_b)
	return pgStr(__r, __n)
}

func QGuiApplication_ExitWithRetcode(retcode int) {
	purego.SyscallN(addr_QGuiApplication_exit2, uintptr(retcode))
}

func QGuiApplication_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QGuiApplication_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QGuiApplication_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QGuiApplication_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QGuiApplication) OnFontDatabaseChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QGuiApplication_connect_fontDatabaseChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGuiApplication) OnScreenAdded(fn func(*QScreen)) int64 {
	id := registerCallback(func(screen uintptr) { fn(UnsafeNewQScreen(unsafe.Pointer(screen))) })
	pg_QGuiApplication_connect_screenAdded(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGuiApplication) OnScreenRemoved(fn func(*QScreen)) int64 {
	id := registerCallback(func(screen uintptr) { fn(UnsafeNewQScreen(unsafe.Pointer(screen))) })
	pg_QGuiApplication_connect_screenRemoved(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGuiApplication) OnPrimaryScreenChanged(fn func(*QScreen)) int64 {
	id := registerCallback(func(screen uintptr) { fn(UnsafeNewQScreen(unsafe.Pointer(screen))) })
	pg_QGuiApplication_connect_primaryScreenChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGuiApplication) OnLastWindowClosed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QGuiApplication_connect_lastWindowClosed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGuiApplication) OnFocusObjectChanged(fn func(*QObject)) int64 {
	id := registerCallback(func(focusObject uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(focusObject))) })
	pg_QGuiApplication_connect_focusObjectChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGuiApplication) OnFocusWindowChanged(fn func(*QWindow)) int64 {
	id := registerCallback(func(focusWindow uintptr) { fn(UnsafeNewQWindow(unsafe.Pointer(focusWindow))) })
	pg_QGuiApplication_connect_focusWindowChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGuiApplication) OnApplicationStateChanged(fn func(ApplicationState)) int64 {
	id := registerCallback(func(state int32) { fn(ApplicationState(state)) })
	pg_QGuiApplication_connect_applicationStateChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGuiApplication) OnLayoutDirectionChanged(fn func(LayoutDirection)) int64 {
	id := registerCallback(func(direction int32) { fn(LayoutDirection(direction)) })
	pg_QGuiApplication_connect_layoutDirectionChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGuiApplication) OnCommitDataRequest(fn func(*QSessionManager)) int64 {
	id := registerCallback(func(sessionManager uintptr) { fn(UnsafeNewQSessionManager(unsafe.Pointer(sessionManager))) })
	pg_QGuiApplication_connect_commitDataRequest(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGuiApplication) OnSaveStateRequest(fn func(*QSessionManager)) int64 {
	id := registerCallback(func(sessionManager uintptr) { fn(UnsafeNewQSessionManager(unsafe.Pointer(sessionManager))) })
	pg_QGuiApplication_connect_saveStateRequest(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGuiApplication) OnApplicationDisplayNameChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QGuiApplication_connect_applicationDisplayNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGuiApplication) OnPaletteChanged(fn func(*QPalette)) int64 {
	id := registerCallback(func(pal uintptr) { fn(UnsafeNewQPalette(unsafe.Pointer(pal))) })
	pg_QGuiApplication_connect_paletteChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGuiApplication) OnFontChanged(fn func(*QFont)) int64 {
	id := registerCallback(func(font uintptr) { fn(UnsafeNewQFont(unsafe.Pointer(font))) })
	pg_QGuiApplication_connect_fontChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGuiApplication) OnAboutToQuit(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QGuiApplication_connect_aboutToQuit(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGuiApplication) OnOrganizationNameChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QGuiApplication_connect_organizationNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGuiApplication) OnOrganizationDomainChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QGuiApplication_connect_organizationDomainChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGuiApplication) OnApplicationNameChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QGuiApplication_connect_applicationNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGuiApplication) OnApplicationVersionChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QGuiApplication_connect_applicationVersionChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGuiApplication) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QGuiApplication_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGuiApplication) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QGuiApplication_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGuiApplication) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QGuiApplication_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGuiApplication) Disconnect(handle int64) {
	pg_QGuiApplication__disconnect(handle)
	unregisterCallback(handle)
}

func NewQGuiApplication(args []string) *QGuiApplication {
	argc := pgMalloc(unsafe.Sizeof(int32(0)))
	*(*int32)(unsafe.Pointer(argc)) = int32(len(args))
	argv := pgMalloc(uintptr(len(args)+1) * unsafe.Sizeof(uintptr(0)))
	__av := unsafe.Slice((*uintptr)(unsafe.Pointer(argv)), len(args)+1)
	for i := range args {
		__av[i] = pgCString(args[i])
	}
	__av[len(args)] = 0
	runtime.LockOSThread()
	return newownedQGuiApplication(pg_QGuiApplication_new_args(argc, argv))
}

func (this *QGuiApplication) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QGuiApplication_delete(this.h)
	}
	this.UnsafeRelease()
}
