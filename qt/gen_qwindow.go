package qt

import (
	"unsafe"
	"math"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QWindow_base func(uintptr) uintptr
	addr_QWindow_new uintptr
	addr_QWindow_new2 uintptr
	addr_QWindow_new3 uintptr
	addr_QWindow_tr uintptr
	addr_QWindow_setSurfaceType uintptr
	addr_QWindow_surfaceType uintptr
	addr_QWindow_isVisible uintptr
	addr_QWindow_visibility uintptr
	addr_QWindow_setVisibility uintptr
	addr_QWindow_create uintptr
	addr_QWindow_winId uintptr
	addr_QWindow_parent uintptr
	addr_QWindow_setParent uintptr
	addr_QWindow_isTopLevel uintptr
	addr_QWindow_isModal uintptr
	addr_QWindow_modality uintptr
	addr_QWindow_setModality uintptr
	addr_QWindow_setFormat uintptr
	addr_QWindow_format uintptr
	addr_QWindow_requestedFormat uintptr
	addr_QWindow_setFlags uintptr
	addr_QWindow_flags uintptr
	addr_QWindow_setFlag uintptr
	addr_QWindow_type uintptr
	addr_QWindow_title uintptr
	pg_QWindow_setOpacity func(uintptr, float64)
	pg_QWindow_opacity func(uintptr) float64
	addr_QWindow_setMask uintptr
	addr_QWindow_mask uintptr
	addr_QWindow_isActive uintptr
	addr_QWindow_reportContentOrientationChange uintptr
	addr_QWindow_contentOrientation uintptr
	pg_QWindow_devicePixelRatio func(uintptr) float64
	addr_QWindow_windowState uintptr
	addr_QWindow_windowStates uintptr
	addr_QWindow_setWindowState uintptr
	addr_QWindow_setWindowStates uintptr
	addr_QWindow_setTransientParent uintptr
	addr_QWindow_transientParent uintptr
	addr_QWindow_isAncestorOf uintptr
	addr_QWindow_isExposed uintptr
	addr_QWindow_minimumWidth uintptr
	addr_QWindow_minimumHeight uintptr
	addr_QWindow_maximumWidth uintptr
	addr_QWindow_maximumHeight uintptr
	addr_QWindow_minimumSize uintptr
	addr_QWindow_maximumSize uintptr
	addr_QWindow_baseSize uintptr
	addr_QWindow_sizeIncrement uintptr
	addr_QWindow_setMinimumSize uintptr
	addr_QWindow_setMaximumSize uintptr
	addr_QWindow_setBaseSize uintptr
	addr_QWindow_setSizeIncrement uintptr
	addr_QWindow_geometry uintptr
	addr_QWindow_frameMargins uintptr
	addr_QWindow_frameGeometry uintptr
	addr_QWindow_framePosition uintptr
	addr_QWindow_setFramePosition uintptr
	addr_QWindow_safeAreaMargins uintptr
	addr_QWindow_width uintptr
	addr_QWindow_height uintptr
	addr_QWindow_x uintptr
	addr_QWindow_y uintptr
	addr_QWindow_size uintptr
	addr_QWindow_position uintptr
	addr_QWindow_setPosition uintptr
	addr_QWindow_setPosition2 uintptr
	addr_QWindow_resize uintptr
	addr_QWindow_resize2 uintptr
	addr_QWindow_setFilePath uintptr
	addr_QWindow_filePath uintptr
	addr_QWindow_setIcon uintptr
	addr_QWindow_icon uintptr
	addr_QWindow_destroy uintptr
	addr_QWindow_setKeyboardGrabEnabled uintptr
	addr_QWindow_setMouseGrabEnabled uintptr
	addr_QWindow_screen uintptr
	addr_QWindow_setScreen uintptr
	addr_QWindow_accessibleRoot uintptr
	addr_QWindow_focusObject uintptr
	addr_QWindow_mapToGlobal uintptr
	addr_QWindow_mapFromGlobal uintptr
	addr_QWindow_mapToGlobal2 uintptr
	addr_QWindow_mapFromGlobal2 uintptr
	addr_QWindow_cursor uintptr
	addr_QWindow_setCursor uintptr
	addr_QWindow_unsetCursor uintptr
	addr_QWindow_fromWinId uintptr
	addr_QWindow_requestActivate uintptr
	addr_QWindow_setVisible uintptr
	addr_QWindow_show uintptr
	addr_QWindow_hide uintptr
	addr_QWindow_showMinimized uintptr
	addr_QWindow_showMaximized uintptr
	addr_QWindow_showFullScreen uintptr
	addr_QWindow_showNormal uintptr
	addr_QWindow_close uintptr
	addr_QWindow_raise uintptr
	addr_QWindow_lower uintptr
	addr_QWindow_startSystemResize uintptr
	addr_QWindow_startSystemMove uintptr
	addr_QWindow_setTitle uintptr
	addr_QWindow_setX uintptr
	addr_QWindow_setY uintptr
	addr_QWindow_setWidth uintptr
	addr_QWindow_setHeight uintptr
	addr_QWindow_setGeometry uintptr
	addr_QWindow_setGeometry2 uintptr
	addr_QWindow_setMinimumWidth uintptr
	addr_QWindow_setMinimumHeight uintptr
	addr_QWindow_setMaximumWidth uintptr
	addr_QWindow_setMaximumHeight uintptr
	addr_QWindow_alert uintptr
	addr_QWindow_requestUpdate uintptr
	addr_QWindow_screenChanged uintptr
	addr_QWindow_modalityChanged uintptr
	addr_QWindow_flagsChanged uintptr
	addr_QWindow_windowStateChanged uintptr
	addr_QWindow_windowTitleChanged uintptr
	addr_QWindow_xChanged uintptr
	addr_QWindow_yChanged uintptr
	addr_QWindow_widthChanged uintptr
	addr_QWindow_heightChanged uintptr
	addr_QWindow_minimumWidthChanged uintptr
	addr_QWindow_minimumHeightChanged uintptr
	addr_QWindow_maximumWidthChanged uintptr
	addr_QWindow_maximumHeightChanged uintptr
	addr_QWindow_safeAreaMarginsChanged uintptr
	addr_QWindow_visibleChanged uintptr
	addr_QWindow_visibilityChanged uintptr
	addr_QWindow_activeChanged uintptr
	addr_QWindow_contentOrientationChanged uintptr
	addr_QWindow_focusObjectChanged uintptr
	pg_QWindow_opacityChanged func(uintptr, float64)
	addr_QWindow_transientParentChanged uintptr
	addr_QWindow_connect uintptr
	addr_QWindow_connect2 uintptr
	addr_QWindow_sender uintptr
	addr_QWindow_senderSignalIndex uintptr
	addr_QWindow_receivers uintptr
	addr_QWindow_isSignalConnected uintptr
	addr_QWindow_surfaceClass uintptr
	addr_QWindow_supportsOpenGL uintptr
	addr_QWindow_tr2 uintptr
	addr_QWindow_tr3 uintptr
	addr_QWindow_parent2 uintptr
	addr_QWindow_setFlag2 uintptr
	addr_QWindow_isAncestorOf2 uintptr
	addr_QWindow_connect3 uintptr
	addr_QWindow_connect4 uintptr
	addr_QWindow_asQSurface uintptr
	addr_QWindow_surfaceType_default uintptr
	addr_QWindow_format_default uintptr
	addr_QWindow_size_default uintptr
	addr_QWindow_accessibleRoot_default uintptr
	addr_QWindow_focusObject_default uintptr
	addr_QWindow_metaObject_default uintptr
	addr_QWindow_exposeEvent_default uintptr
	addr_QWindow_resizeEvent_default uintptr
	addr_QWindow_paintEvent_default uintptr
	addr_QWindow_moveEvent_default uintptr
	addr_QWindow_focusInEvent_default uintptr
	addr_QWindow_focusOutEvent_default uintptr
	addr_QWindow_showEvent_default uintptr
	addr_QWindow_hideEvent_default uintptr
	addr_QWindow_closeEvent_default uintptr
	addr_QWindow_event_default uintptr
	addr_QWindow_keyPressEvent_default uintptr
	addr_QWindow_keyReleaseEvent_default uintptr
	addr_QWindow_mousePressEvent_default uintptr
	addr_QWindow_mouseReleaseEvent_default uintptr
	addr_QWindow_mouseDoubleClickEvent_default uintptr
	addr_QWindow_mouseMoveEvent_default uintptr
	addr_QWindow_wheelEvent_default uintptr
	addr_QWindow_touchEvent_default uintptr
	addr_QWindow_tabletEvent_default uintptr
	addr_QWindow_eventFilter_default uintptr
	addr_QWindow_timerEvent_default uintptr
	addr_QWindow_childEvent_default uintptr
	addr_QWindow_customEvent_default uintptr
	addr_QWindow_connectNotify_default uintptr
	addr_QWindow_disconnectNotify_default uintptr
	pg_QWindow_connect_screenChanged func(uintptr, int64)
	pg_QWindow_connect_modalityChanged func(uintptr, int64)
	pg_QWindow_connect_flagsChanged func(uintptr, int64)
	pg_QWindow_connect_windowStateChanged func(uintptr, int64)
	pg_QWindow_connect_windowTitleChanged func(uintptr, int64)
	pg_QWindow_connect_xChanged func(uintptr, int64)
	pg_QWindow_connect_yChanged func(uintptr, int64)
	pg_QWindow_connect_widthChanged func(uintptr, int64)
	pg_QWindow_connect_heightChanged func(uintptr, int64)
	pg_QWindow_connect_minimumWidthChanged func(uintptr, int64)
	pg_QWindow_connect_minimumHeightChanged func(uintptr, int64)
	pg_QWindow_connect_maximumWidthChanged func(uintptr, int64)
	pg_QWindow_connect_maximumHeightChanged func(uintptr, int64)
	pg_QWindow_connect_safeAreaMarginsChanged func(uintptr, int64)
	pg_QWindow_connect_visibleChanged func(uintptr, int64)
	pg_QWindow_connect_visibilityChanged func(uintptr, int64)
	pg_QWindow_connect_activeChanged func(uintptr, int64)
	pg_QWindow_connect_contentOrientationChanged func(uintptr, int64)
	pg_QWindow_connect_focusObjectChanged func(uintptr, int64)
	pg_QWindow_connect_opacityChanged func(uintptr, int64)
	pg_QWindow_connect_transientParentChanged func(uintptr, int64)
	pg_QWindow_connect_destroyed func(uintptr, int64)
	pg_QWindow_connect_destroyed2 func(uintptr, int64)
	pg_QWindow_connect_objectNameChanged func(uintptr, int64)
	pg_QWindow__disconnect func(int64)
	pg_QWindow__set_override func(uintptr, int32)
	pg_QWindow_delete func(uintptr)
)

func init() {
	reg(&pg_QWindow_base, "QWindow_base")
	symaddr(&addr_QWindow_new, "QWindow_new")
	symaddr(&addr_QWindow_new2, "QWindow_new2")
	symaddr(&addr_QWindow_new3, "QWindow_new3")
	symaddr(&addr_QWindow_tr, "QWindow_tr")
	symaddr(&addr_QWindow_setSurfaceType, "QWindow_setSurfaceType")
	symaddr(&addr_QWindow_surfaceType, "QWindow_surfaceType")
	symaddr(&addr_QWindow_isVisible, "QWindow_isVisible")
	symaddr(&addr_QWindow_visibility, "QWindow_visibility")
	symaddr(&addr_QWindow_setVisibility, "QWindow_setVisibility")
	symaddr(&addr_QWindow_create, "QWindow_create")
	symaddr(&addr_QWindow_winId, "QWindow_winId")
	symaddr(&addr_QWindow_parent, "QWindow_parent")
	symaddr(&addr_QWindow_setParent, "QWindow_setParent")
	symaddr(&addr_QWindow_isTopLevel, "QWindow_isTopLevel")
	symaddr(&addr_QWindow_isModal, "QWindow_isModal")
	symaddr(&addr_QWindow_modality, "QWindow_modality")
	symaddr(&addr_QWindow_setModality, "QWindow_setModality")
	symaddr(&addr_QWindow_setFormat, "QWindow_setFormat")
	symaddr(&addr_QWindow_format, "QWindow_format")
	symaddr(&addr_QWindow_requestedFormat, "QWindow_requestedFormat")
	symaddr(&addr_QWindow_setFlags, "QWindow_setFlags")
	symaddr(&addr_QWindow_flags, "QWindow_flags")
	symaddr(&addr_QWindow_setFlag, "QWindow_setFlag")
	symaddr(&addr_QWindow_type, "QWindow_type")
	symaddr(&addr_QWindow_title, "QWindow_title")
	reg(&pg_QWindow_setOpacity, "QWindow_setOpacity")
	reg(&pg_QWindow_opacity, "QWindow_opacity")
	symaddr(&addr_QWindow_setMask, "QWindow_setMask")
	symaddr(&addr_QWindow_mask, "QWindow_mask")
	symaddr(&addr_QWindow_isActive, "QWindow_isActive")
	symaddr(&addr_QWindow_reportContentOrientationChange, "QWindow_reportContentOrientationChange")
	symaddr(&addr_QWindow_contentOrientation, "QWindow_contentOrientation")
	reg(&pg_QWindow_devicePixelRatio, "QWindow_devicePixelRatio")
	symaddr(&addr_QWindow_windowState, "QWindow_windowState")
	symaddr(&addr_QWindow_windowStates, "QWindow_windowStates")
	symaddr(&addr_QWindow_setWindowState, "QWindow_setWindowState")
	symaddr(&addr_QWindow_setWindowStates, "QWindow_setWindowStates")
	symaddr(&addr_QWindow_setTransientParent, "QWindow_setTransientParent")
	symaddr(&addr_QWindow_transientParent, "QWindow_transientParent")
	symaddr(&addr_QWindow_isAncestorOf, "QWindow_isAncestorOf")
	symaddr(&addr_QWindow_isExposed, "QWindow_isExposed")
	symaddr(&addr_QWindow_minimumWidth, "QWindow_minimumWidth")
	symaddr(&addr_QWindow_minimumHeight, "QWindow_minimumHeight")
	symaddr(&addr_QWindow_maximumWidth, "QWindow_maximumWidth")
	symaddr(&addr_QWindow_maximumHeight, "QWindow_maximumHeight")
	symaddr(&addr_QWindow_minimumSize, "QWindow_minimumSize")
	symaddr(&addr_QWindow_maximumSize, "QWindow_maximumSize")
	symaddr(&addr_QWindow_baseSize, "QWindow_baseSize")
	symaddr(&addr_QWindow_sizeIncrement, "QWindow_sizeIncrement")
	symaddr(&addr_QWindow_setMinimumSize, "QWindow_setMinimumSize")
	symaddr(&addr_QWindow_setMaximumSize, "QWindow_setMaximumSize")
	symaddr(&addr_QWindow_setBaseSize, "QWindow_setBaseSize")
	symaddr(&addr_QWindow_setSizeIncrement, "QWindow_setSizeIncrement")
	symaddr(&addr_QWindow_geometry, "QWindow_geometry")
	symaddr(&addr_QWindow_frameMargins, "QWindow_frameMargins")
	symaddr(&addr_QWindow_frameGeometry, "QWindow_frameGeometry")
	symaddr(&addr_QWindow_framePosition, "QWindow_framePosition")
	symaddr(&addr_QWindow_setFramePosition, "QWindow_setFramePosition")
	symaddr(&addr_QWindow_safeAreaMargins, "QWindow_safeAreaMargins")
	symaddr(&addr_QWindow_width, "QWindow_width")
	symaddr(&addr_QWindow_height, "QWindow_height")
	symaddr(&addr_QWindow_x, "QWindow_x")
	symaddr(&addr_QWindow_y, "QWindow_y")
	symaddr(&addr_QWindow_size, "QWindow_size")
	symaddr(&addr_QWindow_position, "QWindow_position")
	symaddr(&addr_QWindow_setPosition, "QWindow_setPosition")
	symaddr(&addr_QWindow_setPosition2, "QWindow_setPosition2")
	symaddr(&addr_QWindow_resize, "QWindow_resize")
	symaddr(&addr_QWindow_resize2, "QWindow_resize2")
	symaddr(&addr_QWindow_setFilePath, "QWindow_setFilePath")
	symaddr(&addr_QWindow_filePath, "QWindow_filePath")
	symaddr(&addr_QWindow_setIcon, "QWindow_setIcon")
	symaddr(&addr_QWindow_icon, "QWindow_icon")
	symaddr(&addr_QWindow_destroy, "QWindow_destroy")
	symaddr(&addr_QWindow_setKeyboardGrabEnabled, "QWindow_setKeyboardGrabEnabled")
	symaddr(&addr_QWindow_setMouseGrabEnabled, "QWindow_setMouseGrabEnabled")
	symaddr(&addr_QWindow_screen, "QWindow_screen")
	symaddr(&addr_QWindow_setScreen, "QWindow_setScreen")
	symaddr(&addr_QWindow_accessibleRoot, "QWindow_accessibleRoot")
	symaddr(&addr_QWindow_focusObject, "QWindow_focusObject")
	symaddr(&addr_QWindow_mapToGlobal, "QWindow_mapToGlobal")
	symaddr(&addr_QWindow_mapFromGlobal, "QWindow_mapFromGlobal")
	symaddr(&addr_QWindow_mapToGlobal2, "QWindow_mapToGlobal2")
	symaddr(&addr_QWindow_mapFromGlobal2, "QWindow_mapFromGlobal2")
	symaddr(&addr_QWindow_cursor, "QWindow_cursor")
	symaddr(&addr_QWindow_setCursor, "QWindow_setCursor")
	symaddr(&addr_QWindow_unsetCursor, "QWindow_unsetCursor")
	symaddr(&addr_QWindow_fromWinId, "QWindow_fromWinId")
	symaddr(&addr_QWindow_requestActivate, "QWindow_requestActivate")
	symaddr(&addr_QWindow_setVisible, "QWindow_setVisible")
	symaddr(&addr_QWindow_show, "QWindow_show")
	symaddr(&addr_QWindow_hide, "QWindow_hide")
	symaddr(&addr_QWindow_showMinimized, "QWindow_showMinimized")
	symaddr(&addr_QWindow_showMaximized, "QWindow_showMaximized")
	symaddr(&addr_QWindow_showFullScreen, "QWindow_showFullScreen")
	symaddr(&addr_QWindow_showNormal, "QWindow_showNormal")
	symaddr(&addr_QWindow_close, "QWindow_close")
	symaddr(&addr_QWindow_raise, "QWindow_raise")
	symaddr(&addr_QWindow_lower, "QWindow_lower")
	symaddr(&addr_QWindow_startSystemResize, "QWindow_startSystemResize")
	symaddr(&addr_QWindow_startSystemMove, "QWindow_startSystemMove")
	symaddr(&addr_QWindow_setTitle, "QWindow_setTitle")
	symaddr(&addr_QWindow_setX, "QWindow_setX")
	symaddr(&addr_QWindow_setY, "QWindow_setY")
	symaddr(&addr_QWindow_setWidth, "QWindow_setWidth")
	symaddr(&addr_QWindow_setHeight, "QWindow_setHeight")
	symaddr(&addr_QWindow_setGeometry, "QWindow_setGeometry")
	symaddr(&addr_QWindow_setGeometry2, "QWindow_setGeometry2")
	symaddr(&addr_QWindow_setMinimumWidth, "QWindow_setMinimumWidth")
	symaddr(&addr_QWindow_setMinimumHeight, "QWindow_setMinimumHeight")
	symaddr(&addr_QWindow_setMaximumWidth, "QWindow_setMaximumWidth")
	symaddr(&addr_QWindow_setMaximumHeight, "QWindow_setMaximumHeight")
	symaddr(&addr_QWindow_alert, "QWindow_alert")
	symaddr(&addr_QWindow_requestUpdate, "QWindow_requestUpdate")
	symaddr(&addr_QWindow_screenChanged, "QWindow_screenChanged")
	symaddr(&addr_QWindow_modalityChanged, "QWindow_modalityChanged")
	symaddr(&addr_QWindow_flagsChanged, "QWindow_flagsChanged")
	symaddr(&addr_QWindow_windowStateChanged, "QWindow_windowStateChanged")
	symaddr(&addr_QWindow_windowTitleChanged, "QWindow_windowTitleChanged")
	symaddr(&addr_QWindow_xChanged, "QWindow_xChanged")
	symaddr(&addr_QWindow_yChanged, "QWindow_yChanged")
	symaddr(&addr_QWindow_widthChanged, "QWindow_widthChanged")
	symaddr(&addr_QWindow_heightChanged, "QWindow_heightChanged")
	symaddr(&addr_QWindow_minimumWidthChanged, "QWindow_minimumWidthChanged")
	symaddr(&addr_QWindow_minimumHeightChanged, "QWindow_minimumHeightChanged")
	symaddr(&addr_QWindow_maximumWidthChanged, "QWindow_maximumWidthChanged")
	symaddr(&addr_QWindow_maximumHeightChanged, "QWindow_maximumHeightChanged")
	symaddr(&addr_QWindow_safeAreaMarginsChanged, "QWindow_safeAreaMarginsChanged")
	symaddr(&addr_QWindow_visibleChanged, "QWindow_visibleChanged")
	symaddr(&addr_QWindow_visibilityChanged, "QWindow_visibilityChanged")
	symaddr(&addr_QWindow_activeChanged, "QWindow_activeChanged")
	symaddr(&addr_QWindow_contentOrientationChanged, "QWindow_contentOrientationChanged")
	symaddr(&addr_QWindow_focusObjectChanged, "QWindow_focusObjectChanged")
	reg(&pg_QWindow_opacityChanged, "QWindow_opacityChanged")
	symaddr(&addr_QWindow_transientParentChanged, "QWindow_transientParentChanged")
	symaddr(&addr_QWindow_connect, "QWindow_connect")
	symaddr(&addr_QWindow_connect2, "QWindow_connect2")
	symaddr(&addr_QWindow_sender, "QWindow_sender")
	symaddr(&addr_QWindow_senderSignalIndex, "QWindow_senderSignalIndex")
	symaddr(&addr_QWindow_receivers, "QWindow_receivers")
	symaddr(&addr_QWindow_isSignalConnected, "QWindow_isSignalConnected")
	symaddr(&addr_QWindow_surfaceClass, "QWindow_surfaceClass")
	symaddr(&addr_QWindow_supportsOpenGL, "QWindow_supportsOpenGL")
	symaddr(&addr_QWindow_tr2, "QWindow_tr2")
	symaddr(&addr_QWindow_tr3, "QWindow_tr3")
	symaddr(&addr_QWindow_parent2, "QWindow_parent2")
	symaddr(&addr_QWindow_setFlag2, "QWindow_setFlag2")
	symaddr(&addr_QWindow_isAncestorOf2, "QWindow_isAncestorOf2")
	symaddr(&addr_QWindow_connect3, "QWindow_connect3")
	symaddr(&addr_QWindow_connect4, "QWindow_connect4")
	symaddr(&addr_QWindow_asQSurface, "QWindow_asQSurface")
	symaddr(&addr_QWindow_surfaceType_default, "QWindow_surfaceType_default")
	symaddr(&addr_QWindow_format_default, "QWindow_format_default")
	symaddr(&addr_QWindow_size_default, "QWindow_size_default")
	symaddr(&addr_QWindow_accessibleRoot_default, "QWindow_accessibleRoot_default")
	symaddr(&addr_QWindow_focusObject_default, "QWindow_focusObject_default")
	symaddr(&addr_QWindow_metaObject_default, "QWindow_metaObject_default")
	symaddr(&addr_QWindow_exposeEvent_default, "QWindow_exposeEvent_default")
	symaddr(&addr_QWindow_resizeEvent_default, "QWindow_resizeEvent_default")
	symaddr(&addr_QWindow_paintEvent_default, "QWindow_paintEvent_default")
	symaddr(&addr_QWindow_moveEvent_default, "QWindow_moveEvent_default")
	symaddr(&addr_QWindow_focusInEvent_default, "QWindow_focusInEvent_default")
	symaddr(&addr_QWindow_focusOutEvent_default, "QWindow_focusOutEvent_default")
	symaddr(&addr_QWindow_showEvent_default, "QWindow_showEvent_default")
	symaddr(&addr_QWindow_hideEvent_default, "QWindow_hideEvent_default")
	symaddr(&addr_QWindow_closeEvent_default, "QWindow_closeEvent_default")
	symaddr(&addr_QWindow_event_default, "QWindow_event_default")
	symaddr(&addr_QWindow_keyPressEvent_default, "QWindow_keyPressEvent_default")
	symaddr(&addr_QWindow_keyReleaseEvent_default, "QWindow_keyReleaseEvent_default")
	symaddr(&addr_QWindow_mousePressEvent_default, "QWindow_mousePressEvent_default")
	symaddr(&addr_QWindow_mouseReleaseEvent_default, "QWindow_mouseReleaseEvent_default")
	symaddr(&addr_QWindow_mouseDoubleClickEvent_default, "QWindow_mouseDoubleClickEvent_default")
	symaddr(&addr_QWindow_mouseMoveEvent_default, "QWindow_mouseMoveEvent_default")
	symaddr(&addr_QWindow_wheelEvent_default, "QWindow_wheelEvent_default")
	symaddr(&addr_QWindow_touchEvent_default, "QWindow_touchEvent_default")
	symaddr(&addr_QWindow_tabletEvent_default, "QWindow_tabletEvent_default")
	symaddr(&addr_QWindow_eventFilter_default, "QWindow_eventFilter_default")
	symaddr(&addr_QWindow_timerEvent_default, "QWindow_timerEvent_default")
	symaddr(&addr_QWindow_childEvent_default, "QWindow_childEvent_default")
	symaddr(&addr_QWindow_customEvent_default, "QWindow_customEvent_default")
	symaddr(&addr_QWindow_connectNotify_default, "QWindow_connectNotify_default")
	symaddr(&addr_QWindow_disconnectNotify_default, "QWindow_disconnectNotify_default")
	reg(&pg_QWindow_connect_screenChanged, "QWindow_connect_screenChanged")
	reg(&pg_QWindow_connect_modalityChanged, "QWindow_connect_modalityChanged")
	reg(&pg_QWindow_connect_flagsChanged, "QWindow_connect_flagsChanged")
	reg(&pg_QWindow_connect_windowStateChanged, "QWindow_connect_windowStateChanged")
	reg(&pg_QWindow_connect_windowTitleChanged, "QWindow_connect_windowTitleChanged")
	reg(&pg_QWindow_connect_xChanged, "QWindow_connect_xChanged")
	reg(&pg_QWindow_connect_yChanged, "QWindow_connect_yChanged")
	reg(&pg_QWindow_connect_widthChanged, "QWindow_connect_widthChanged")
	reg(&pg_QWindow_connect_heightChanged, "QWindow_connect_heightChanged")
	reg(&pg_QWindow_connect_minimumWidthChanged, "QWindow_connect_minimumWidthChanged")
	reg(&pg_QWindow_connect_minimumHeightChanged, "QWindow_connect_minimumHeightChanged")
	reg(&pg_QWindow_connect_maximumWidthChanged, "QWindow_connect_maximumWidthChanged")
	reg(&pg_QWindow_connect_maximumHeightChanged, "QWindow_connect_maximumHeightChanged")
	reg(&pg_QWindow_connect_safeAreaMarginsChanged, "QWindow_connect_safeAreaMarginsChanged")
	reg(&pg_QWindow_connect_visibleChanged, "QWindow_connect_visibleChanged")
	reg(&pg_QWindow_connect_visibilityChanged, "QWindow_connect_visibilityChanged")
	reg(&pg_QWindow_connect_activeChanged, "QWindow_connect_activeChanged")
	reg(&pg_QWindow_connect_contentOrientationChanged, "QWindow_connect_contentOrientationChanged")
	reg(&pg_QWindow_connect_focusObjectChanged, "QWindow_connect_focusObjectChanged")
	reg(&pg_QWindow_connect_opacityChanged, "QWindow_connect_opacityChanged")
	reg(&pg_QWindow_connect_transientParentChanged, "QWindow_connect_transientParentChanged")
	reg(&pg_QWindow_connect_destroyed, "QWindow_connect_destroyed")
	reg(&pg_QWindow_connect_destroyed2, "QWindow_connect_destroyed2")
	reg(&pg_QWindow_connect_objectNameChanged, "QWindow_connect_objectNameChanged")
	reg(&pg_QWindow__disconnect, "QWindow__disconnect")
	reg(&pg_QWindow__set_override, "QWindow__set_override")
	reg(&pg_QWindow_delete, "QWindow_delete")
}

type QWindow struct {
	h uintptr
	*QObject
}

func bindQWindow(h uintptr) *QWindow {
	if h == 0 {
		return nil
	}
	return &QWindow{h: h, QObject: bindQObject(pg_QWindow_base(h))}
}

func newownedQWindow(h uintptr) *QWindow { return bindQWindow(h) }

func (this *QWindow) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QWindow) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQWindow(h unsafe.Pointer) *QWindow { return bindQWindow(uintptr(h)) }
func UnsafeNewQWindowOwned(h unsafe.Pointer) *QWindow { return newownedQWindow(uintptr(h)) }

func (this *QWindow) UnsafeRelease() {
	this.h = 0
	if this.QObject != nil {
		this.QObject.UnsafeRelease()
	}
}

func NewQWindow() *QWindow {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QWindow_new, uintptr(__goid))
	__o := newownedQWindow(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QWindow_virtIdx, QWindow_enable)
	return __o
}

func NewQWindow2(parent *QWindow) *QWindow {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QWindow_new2, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQWindow(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QWindow_virtIdx, QWindow_enable)
	return __o
}

func NewQWindow3(screen *QScreen) *QWindow {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QWindow_new3, uintptr(__goid), uintptr(screen.UnsafePointer()))
	__o := newownedQWindow(__r)
	runtime.KeepAlive(screen)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QWindow_virtIdx, QWindow_enable)
	return __o
}

func QWindow_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QWindow_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QWindow) SetSurfaceType(surfaceType QSurface__SurfaceType) {
	purego.SyscallN(addr_QWindow_setSurfaceType, this.h, uintptr(surfaceType))
	runtime.KeepAlive(this)
}

func (this *QWindow) SurfaceType() QSurface__SurfaceType {
	__r0, _, _ := purego.SyscallN(addr_QWindow_surfaceType, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QSurface__SurfaceType(__r)
}

func (this *QWindow) IsVisible() bool {
	__r0, _, _ := purego.SyscallN(addr_QWindow_isVisible, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QWindow) Visibility() QWindow__Visibility {
	__r0, _, _ := purego.SyscallN(addr_QWindow_visibility, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QWindow__Visibility(__r)
}

func (this *QWindow) SetVisibility(v QWindow__Visibility) {
	purego.SyscallN(addr_QWindow_setVisibility, this.h, uintptr(v))
	runtime.KeepAlive(this)
}

func (this *QWindow) Create() {
	purego.SyscallN(addr_QWindow_create, this.h)
	runtime.KeepAlive(this)
}

func (this *QWindow) WinId() uint64 {
	__r0, _, _ := purego.SyscallN(addr_QWindow_winId, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return uint64(__r)
}

func (this *QWindow) Parent() *QWindow {
	__r, _, _ := purego.SyscallN(addr_QWindow_parent, this.h)
	runtime.KeepAlive(this)
	return bindQWindow(__r)
}

func (this *QWindow) SetParent(parent *QWindow) {
	purego.SyscallN(addr_QWindow_setParent, this.h, uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
}

func (this *QWindow) IsTopLevel() bool {
	__r0, _, _ := purego.SyscallN(addr_QWindow_isTopLevel, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QWindow) IsModal() bool {
	__r0, _, _ := purego.SyscallN(addr_QWindow_isModal, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QWindow) Modality() WindowModality {
	__r0, _, _ := purego.SyscallN(addr_QWindow_modality, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return WindowModality(__r)
}

func (this *QWindow) SetModality(modality WindowModality) {
	purego.SyscallN(addr_QWindow_setModality, this.h, uintptr(modality))
	runtime.KeepAlive(this)
}

func (this *QWindow) SetFormat(format *QSurfaceFormat) {
	purego.SyscallN(addr_QWindow_setFormat, this.h, uintptr(format.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(format)
}

func (this *QWindow) Format() *QSurfaceFormat {
	__r, _, _ := purego.SyscallN(addr_QWindow_format, this.h)
	runtime.KeepAlive(this)
	return newownedQSurfaceFormat(__r)
}

func (this *QWindow) RequestedFormat() *QSurfaceFormat {
	__r, _, _ := purego.SyscallN(addr_QWindow_requestedFormat, this.h)
	runtime.KeepAlive(this)
	return newownedQSurfaceFormat(__r)
}

func (this *QWindow) SetFlags(flags WindowType) {
	purego.SyscallN(addr_QWindow_setFlags, this.h, uintptr(flags))
	runtime.KeepAlive(this)
}

func (this *QWindow) Flags() WindowType {
	__r0, _, _ := purego.SyscallN(addr_QWindow_flags, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return WindowType(__r)
}

func (this *QWindow) SetFlag(arg1 WindowType) {
	purego.SyscallN(addr_QWindow_setFlag, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QWindow) Type() WindowType {
	__r0, _, _ := purego.SyscallN(addr_QWindow_type, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return WindowType(__r)
}

func (this *QWindow) Title() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QWindow_title, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QWindow) SetOpacity(level float64) {
	pg_QWindow_setOpacity(this.h, level)
	runtime.KeepAlive(this)
}

func (this *QWindow) Opacity() float64 {
	__r := pg_QWindow_opacity(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QWindow) SetMask(region *QRegion) {
	purego.SyscallN(addr_QWindow_setMask, this.h, uintptr(region.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(region)
}

func (this *QWindow) Mask() *QRegion {
	__r, _, _ := purego.SyscallN(addr_QWindow_mask, this.h)
	runtime.KeepAlive(this)
	return newownedQRegion(__r)
}

func (this *QWindow) IsActive() bool {
	__r0, _, _ := purego.SyscallN(addr_QWindow_isActive, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QWindow) ReportContentOrientationChange(orientation ScreenOrientation) {
	purego.SyscallN(addr_QWindow_reportContentOrientationChange, this.h, uintptr(orientation))
	runtime.KeepAlive(this)
}

func (this *QWindow) ContentOrientation() ScreenOrientation {
	__r0, _, _ := purego.SyscallN(addr_QWindow_contentOrientation, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return ScreenOrientation(__r)
}

func (this *QWindow) DevicePixelRatio() float64 {
	__r := pg_QWindow_devicePixelRatio(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QWindow) WindowState() WindowState {
	__r0, _, _ := purego.SyscallN(addr_QWindow_windowState, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return WindowState(__r)
}

func (this *QWindow) WindowStates() WindowState {
	__r0, _, _ := purego.SyscallN(addr_QWindow_windowStates, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return WindowState(__r)
}

func (this *QWindow) SetWindowState(state WindowState) {
	purego.SyscallN(addr_QWindow_setWindowState, this.h, uintptr(state))
	runtime.KeepAlive(this)
}

func (this *QWindow) SetWindowStates(states WindowState) {
	purego.SyscallN(addr_QWindow_setWindowStates, this.h, uintptr(states))
	runtime.KeepAlive(this)
}

func (this *QWindow) SetTransientParent(parent *QWindow) {
	purego.SyscallN(addr_QWindow_setTransientParent, this.h, uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
}

func (this *QWindow) TransientParent() *QWindow {
	__r, _, _ := purego.SyscallN(addr_QWindow_transientParent, this.h)
	runtime.KeepAlive(this)
	return bindQWindow(__r)
}

func (this *QWindow) IsAncestorOf(child *QWindow) bool {
	__r0, _, _ := purego.SyscallN(addr_QWindow_isAncestorOf, this.h, uintptr(child.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(child)
	return __r != 0
}

func (this *QWindow) IsExposed() bool {
	__r0, _, _ := purego.SyscallN(addr_QWindow_isExposed, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QWindow) MinimumWidth() int {
	__r0, _, _ := purego.SyscallN(addr_QWindow_minimumWidth, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QWindow) MinimumHeight() int {
	__r0, _, _ := purego.SyscallN(addr_QWindow_minimumHeight, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QWindow) MaximumWidth() int {
	__r0, _, _ := purego.SyscallN(addr_QWindow_maximumWidth, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QWindow) MaximumHeight() int {
	__r0, _, _ := purego.SyscallN(addr_QWindow_maximumHeight, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QWindow) MinimumSize() *QSize {
	__r, _, _ := purego.SyscallN(addr_QWindow_minimumSize, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QWindow) MaximumSize() *QSize {
	__r, _, _ := purego.SyscallN(addr_QWindow_maximumSize, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QWindow) BaseSize() *QSize {
	__r, _, _ := purego.SyscallN(addr_QWindow_baseSize, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QWindow) SizeIncrement() *QSize {
	__r, _, _ := purego.SyscallN(addr_QWindow_sizeIncrement, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QWindow) SetMinimumSize(size *QSize) {
	purego.SyscallN(addr_QWindow_setMinimumSize, this.h, uintptr(size.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(size)
}

func (this *QWindow) SetMaximumSize(size *QSize) {
	purego.SyscallN(addr_QWindow_setMaximumSize, this.h, uintptr(size.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(size)
}

func (this *QWindow) SetBaseSize(size *QSize) {
	purego.SyscallN(addr_QWindow_setBaseSize, this.h, uintptr(size.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(size)
}

func (this *QWindow) SetSizeIncrement(size *QSize) {
	purego.SyscallN(addr_QWindow_setSizeIncrement, this.h, uintptr(size.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(size)
}

func (this *QWindow) Geometry() *QRect {
	__r, _, _ := purego.SyscallN(addr_QWindow_geometry, this.h)
	runtime.KeepAlive(this)
	return newownedQRect(__r)
}

func (this *QWindow) FrameMargins() *QMargins {
	__r, _, _ := purego.SyscallN(addr_QWindow_frameMargins, this.h)
	runtime.KeepAlive(this)
	return newownedQMargins(__r)
}

func (this *QWindow) FrameGeometry() *QRect {
	__r, _, _ := purego.SyscallN(addr_QWindow_frameGeometry, this.h)
	runtime.KeepAlive(this)
	return newownedQRect(__r)
}

func (this *QWindow) FramePosition() *QPoint {
	__r, _, _ := purego.SyscallN(addr_QWindow_framePosition, this.h)
	runtime.KeepAlive(this)
	return newownedQPoint(__r)
}

func (this *QWindow) SetFramePosition(point *QPoint) {
	purego.SyscallN(addr_QWindow_setFramePosition, this.h, uintptr(point.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(point)
}

func (this *QWindow) SafeAreaMargins() *QMargins {
	__r, _, _ := purego.SyscallN(addr_QWindow_safeAreaMargins, this.h)
	runtime.KeepAlive(this)
	return newownedQMargins(__r)
}

func (this *QWindow) Width() int {
	__r0, _, _ := purego.SyscallN(addr_QWindow_width, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QWindow) Height() int {
	__r0, _, _ := purego.SyscallN(addr_QWindow_height, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QWindow) X() int {
	__r0, _, _ := purego.SyscallN(addr_QWindow_x, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QWindow) Y() int {
	__r0, _, _ := purego.SyscallN(addr_QWindow_y, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QWindow) Size() *QSize {
	__r, _, _ := purego.SyscallN(addr_QWindow_size, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QWindow) Position() *QPoint {
	__r, _, _ := purego.SyscallN(addr_QWindow_position, this.h)
	runtime.KeepAlive(this)
	return newownedQPoint(__r)
}

func (this *QWindow) SetPosition(pt *QPoint) {
	purego.SyscallN(addr_QWindow_setPosition, this.h, uintptr(pt.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pt)
}

func (this *QWindow) SetPosition2(posx int, posy int) {
	purego.SyscallN(addr_QWindow_setPosition2, this.h, uintptr(posx), uintptr(posy))
	runtime.KeepAlive(this)
}

func (this *QWindow) Resize(newSize *QSize) {
	purego.SyscallN(addr_QWindow_resize, this.h, uintptr(newSize.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(newSize)
}

func (this *QWindow) Resize2(w int, h int) {
	purego.SyscallN(addr_QWindow_resize2, this.h, uintptr(w), uintptr(h))
	runtime.KeepAlive(this)
}

func (this *QWindow) SetFilePath(filePath string) {
	var filePath_p uintptr
	if len(filePath) > 0 {
		filePath_p = uintptr(unsafe.Pointer(unsafe.StringData(filePath)))
	}
	purego.SyscallN(addr_QWindow_setFilePath, this.h, filePath_p, uintptr(len(filePath)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(filePath)
}

func (this *QWindow) FilePath() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QWindow_filePath, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QWindow) SetIcon(icon *QIcon) {
	purego.SyscallN(addr_QWindow_setIcon, this.h, uintptr(icon.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(icon)
}

func (this *QWindow) Icon() *QIcon {
	__r, _, _ := purego.SyscallN(addr_QWindow_icon, this.h)
	runtime.KeepAlive(this)
	return newownedQIcon(__r)
}

func (this *QWindow) Destroy() {
	purego.SyscallN(addr_QWindow_destroy, this.h)
	runtime.KeepAlive(this)
}

func (this *QWindow) SetKeyboardGrabEnabled(grab bool) bool {
	__r0, _, _ := purego.SyscallN(addr_QWindow_setKeyboardGrabEnabled, this.h, uintptr(b2i(grab)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QWindow) SetMouseGrabEnabled(grab bool) bool {
	__r0, _, _ := purego.SyscallN(addr_QWindow_setMouseGrabEnabled, this.h, uintptr(b2i(grab)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QWindow) Screen() *QScreen {
	__r, _, _ := purego.SyscallN(addr_QWindow_screen, this.h)
	runtime.KeepAlive(this)
	return bindQScreen(__r)
}

func (this *QWindow) SetScreen(screen *QScreen) {
	purego.SyscallN(addr_QWindow_setScreen, this.h, uintptr(screen.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(screen)
}

func (this *QWindow) AccessibleRoot() *QAccessibleInterface {
	__r, _, _ := purego.SyscallN(addr_QWindow_accessibleRoot, this.h)
	runtime.KeepAlive(this)
	return bindQAccessibleInterface(__r)
}

func (this *QWindow) FocusObject() *QObject {
	__r, _, _ := purego.SyscallN(addr_QWindow_focusObject, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QWindow) MapToGlobal(pos *QPointF) *QPointF {
	__r, _, _ := purego.SyscallN(addr_QWindow_mapToGlobal, this.h, uintptr(pos.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pos)
	return newownedQPointF(__r)
}

func (this *QWindow) MapFromGlobal(pos *QPointF) *QPointF {
	__r, _, _ := purego.SyscallN(addr_QWindow_mapFromGlobal, this.h, uintptr(pos.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pos)
	return newownedQPointF(__r)
}

func (this *QWindow) MapToGlobalWithPos(pos *QPoint) *QPoint {
	__r, _, _ := purego.SyscallN(addr_QWindow_mapToGlobal2, this.h, uintptr(pos.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pos)
	return newownedQPoint(__r)
}

func (this *QWindow) MapFromGlobalWithPos(pos *QPoint) *QPoint {
	__r, _, _ := purego.SyscallN(addr_QWindow_mapFromGlobal2, this.h, uintptr(pos.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pos)
	return newownedQPoint(__r)
}

func (this *QWindow) Cursor() *QCursor {
	__r, _, _ := purego.SyscallN(addr_QWindow_cursor, this.h)
	runtime.KeepAlive(this)
	return newownedQCursor(__r)
}

func (this *QWindow) SetCursor(arg1 *QCursor) {
	purego.SyscallN(addr_QWindow_setCursor, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QWindow) UnsetCursor() {
	purego.SyscallN(addr_QWindow_unsetCursor, this.h)
	runtime.KeepAlive(this)
}

func QWindow_FromWinId(id uint64) *QWindow {
	__r, _, _ := purego.SyscallN(addr_QWindow_fromWinId, uintptr(id))
	return bindQWindow(__r)
}

func (this *QWindow) RequestActivate() {
	purego.SyscallN(addr_QWindow_requestActivate, this.h)
	runtime.KeepAlive(this)
}

func (this *QWindow) SetVisible(visible bool) {
	purego.SyscallN(addr_QWindow_setVisible, this.h, uintptr(b2i(visible)))
	runtime.KeepAlive(this)
}

func (this *QWindow) Show() {
	purego.SyscallN(addr_QWindow_show, this.h)
	runtime.KeepAlive(this)
}

func (this *QWindow) Hide() {
	purego.SyscallN(addr_QWindow_hide, this.h)
	runtime.KeepAlive(this)
}

func (this *QWindow) ShowMinimized() {
	purego.SyscallN(addr_QWindow_showMinimized, this.h)
	runtime.KeepAlive(this)
}

func (this *QWindow) ShowMaximized() {
	purego.SyscallN(addr_QWindow_showMaximized, this.h)
	runtime.KeepAlive(this)
}

func (this *QWindow) ShowFullScreen() {
	purego.SyscallN(addr_QWindow_showFullScreen, this.h)
	runtime.KeepAlive(this)
}

func (this *QWindow) ShowNormal() {
	purego.SyscallN(addr_QWindow_showNormal, this.h)
	runtime.KeepAlive(this)
}

func (this *QWindow) Close() bool {
	__r0, _, _ := purego.SyscallN(addr_QWindow_close, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QWindow) Raise() {
	purego.SyscallN(addr_QWindow_raise, this.h)
	runtime.KeepAlive(this)
}

func (this *QWindow) Lower() {
	purego.SyscallN(addr_QWindow_lower, this.h)
	runtime.KeepAlive(this)
}

func (this *QWindow) StartSystemResize(edges Edge) bool {
	__r0, _, _ := purego.SyscallN(addr_QWindow_startSystemResize, this.h, uintptr(edges))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QWindow) StartSystemMove() bool {
	__r0, _, _ := purego.SyscallN(addr_QWindow_startSystemMove, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QWindow) SetTitle(arg1 string) {
	var arg1_p uintptr
	if len(arg1) > 0 {
		arg1_p = uintptr(unsafe.Pointer(unsafe.StringData(arg1)))
	}
	purego.SyscallN(addr_QWindow_setTitle, this.h, arg1_p, uintptr(len(arg1)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QWindow) SetX(arg int) {
	purego.SyscallN(addr_QWindow_setX, this.h, uintptr(arg))
	runtime.KeepAlive(this)
}

func (this *QWindow) SetY(arg int) {
	purego.SyscallN(addr_QWindow_setY, this.h, uintptr(arg))
	runtime.KeepAlive(this)
}

func (this *QWindow) SetWidth(arg int) {
	purego.SyscallN(addr_QWindow_setWidth, this.h, uintptr(arg))
	runtime.KeepAlive(this)
}

func (this *QWindow) SetHeight(arg int) {
	purego.SyscallN(addr_QWindow_setHeight, this.h, uintptr(arg))
	runtime.KeepAlive(this)
}

func (this *QWindow) SetGeometry(posx int, posy int, w int, h int) {
	purego.SyscallN(addr_QWindow_setGeometry, this.h, uintptr(posx), uintptr(posy), uintptr(w), uintptr(h))
	runtime.KeepAlive(this)
}

func (this *QWindow) SetGeometryWithRect(rect *QRect) {
	purego.SyscallN(addr_QWindow_setGeometry2, this.h, uintptr(rect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
}

func (this *QWindow) SetMinimumWidth(w int) {
	purego.SyscallN(addr_QWindow_setMinimumWidth, this.h, uintptr(w))
	runtime.KeepAlive(this)
}

func (this *QWindow) SetMinimumHeight(h int) {
	purego.SyscallN(addr_QWindow_setMinimumHeight, this.h, uintptr(h))
	runtime.KeepAlive(this)
}

func (this *QWindow) SetMaximumWidth(w int) {
	purego.SyscallN(addr_QWindow_setMaximumWidth, this.h, uintptr(w))
	runtime.KeepAlive(this)
}

func (this *QWindow) SetMaximumHeight(h int) {
	purego.SyscallN(addr_QWindow_setMaximumHeight, this.h, uintptr(h))
	runtime.KeepAlive(this)
}

func (this *QWindow) Alert(msec int) {
	purego.SyscallN(addr_QWindow_alert, this.h, uintptr(msec))
	runtime.KeepAlive(this)
}

func (this *QWindow) RequestUpdate() {
	purego.SyscallN(addr_QWindow_requestUpdate, this.h)
	runtime.KeepAlive(this)
}

func (this *QWindow) ScreenChanged(screen *QScreen) {
	purego.SyscallN(addr_QWindow_screenChanged, this.h, uintptr(screen.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(screen)
}

func (this *QWindow) ModalityChanged(modality WindowModality) {
	purego.SyscallN(addr_QWindow_modalityChanged, this.h, uintptr(modality))
	runtime.KeepAlive(this)
}

func (this *QWindow) FlagsChanged(flags WindowType) {
	purego.SyscallN(addr_QWindow_flagsChanged, this.h, uintptr(flags))
	runtime.KeepAlive(this)
}

func (this *QWindow) WindowStateChanged(windowState WindowState) {
	purego.SyscallN(addr_QWindow_windowStateChanged, this.h, uintptr(windowState))
	runtime.KeepAlive(this)
}

func (this *QWindow) WindowTitleChanged(title string) {
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	purego.SyscallN(addr_QWindow_windowTitleChanged, this.h, title_p, uintptr(len(title)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(title)
}

func (this *QWindow) XChanged(arg int) {
	purego.SyscallN(addr_QWindow_xChanged, this.h, uintptr(arg))
	runtime.KeepAlive(this)
}

func (this *QWindow) YChanged(arg int) {
	purego.SyscallN(addr_QWindow_yChanged, this.h, uintptr(arg))
	runtime.KeepAlive(this)
}

func (this *QWindow) WidthChanged(arg int) {
	purego.SyscallN(addr_QWindow_widthChanged, this.h, uintptr(arg))
	runtime.KeepAlive(this)
}

func (this *QWindow) HeightChanged(arg int) {
	purego.SyscallN(addr_QWindow_heightChanged, this.h, uintptr(arg))
	runtime.KeepAlive(this)
}

func (this *QWindow) MinimumWidthChanged(arg int) {
	purego.SyscallN(addr_QWindow_minimumWidthChanged, this.h, uintptr(arg))
	runtime.KeepAlive(this)
}

func (this *QWindow) MinimumHeightChanged(arg int) {
	purego.SyscallN(addr_QWindow_minimumHeightChanged, this.h, uintptr(arg))
	runtime.KeepAlive(this)
}

func (this *QWindow) MaximumWidthChanged(arg int) {
	purego.SyscallN(addr_QWindow_maximumWidthChanged, this.h, uintptr(arg))
	runtime.KeepAlive(this)
}

func (this *QWindow) MaximumHeightChanged(arg int) {
	purego.SyscallN(addr_QWindow_maximumHeightChanged, this.h, uintptr(arg))
	runtime.KeepAlive(this)
}

func (this *QWindow) SafeAreaMarginsChanged(arg *QMargins) {
	purego.SyscallN(addr_QWindow_safeAreaMarginsChanged, this.h, uintptr(arg.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg)
}

func (this *QWindow) VisibleChanged(arg bool) {
	purego.SyscallN(addr_QWindow_visibleChanged, this.h, uintptr(b2i(arg)))
	runtime.KeepAlive(this)
}

func (this *QWindow) VisibilityChanged(visibility QWindow__Visibility) {
	purego.SyscallN(addr_QWindow_visibilityChanged, this.h, uintptr(visibility))
	runtime.KeepAlive(this)
}

func (this *QWindow) ActiveChanged() {
	purego.SyscallN(addr_QWindow_activeChanged, this.h)
	runtime.KeepAlive(this)
}

func (this *QWindow) ContentOrientationChanged(orientation ScreenOrientation) {
	purego.SyscallN(addr_QWindow_contentOrientationChanged, this.h, uintptr(orientation))
	runtime.KeepAlive(this)
}

func (this *QWindow) FocusObjectChanged(object *QObject) {
	purego.SyscallN(addr_QWindow_focusObjectChanged, this.h, uintptr(object.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(object)
}

func (this *QWindow) OpacityChanged(opacity float64) {
	pg_QWindow_opacityChanged(this.h, opacity)
	runtime.KeepAlive(this)
}

func (this *QWindow) TransientParentChanged(transientParent *QWindow) {
	purego.SyscallN(addr_QWindow_transientParentChanged, this.h, uintptr(transientParent.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(transientParent)
}

func QWindow_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QWindow_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QWindow_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QWindow_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QWindow) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QWindow_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QWindow) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QWindow_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QWindow) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QWindow_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QWindow) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QWindow_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func (this *QWindow) SurfaceClass() QSurface__SurfaceClass {
	__r0, _, _ := purego.SyscallN(addr_QWindow_surfaceClass, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QSurface__SurfaceClass(__r)
}

func (this *QWindow) SupportsOpenGL() bool {
	__r0, _, _ := purego.SyscallN(addr_QWindow_supportsOpenGL, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QWindow_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QWindow_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QWindow_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QWindow_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func (this *QWindow) ParentWithMode(mode QWindow__AncestorMode) *QWindow {
	__r, _, _ := purego.SyscallN(addr_QWindow_parent2, this.h, uintptr(mode))
	runtime.KeepAlive(this)
	return bindQWindow(__r)
}

func (this *QWindow) SetFlag2(arg1 WindowType, on bool) {
	purego.SyscallN(addr_QWindow_setFlag2, this.h, uintptr(arg1), uintptr(b2i(on)))
	runtime.KeepAlive(this)
}

func (this *QWindow) IsAncestorOf2(child *QWindow, mode QWindow__AncestorMode) bool {
	__r0, _, _ := purego.SyscallN(addr_QWindow_isAncestorOf2, this.h, uintptr(child.UnsafePointer()), uintptr(mode))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(child)
	return __r != 0
}

func QWindow_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QWindow_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QWindow_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QWindow_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QWindow) AsQSurface() *QSurface {
	__r, _, _ := purego.SyscallN(addr_QWindow_asQSurface, this.h)
	runtime.KeepAlive(this)
	return bindQSurface(__r)
}

func (this *QWindow) SurfaceTypeDefault() QSurface__SurfaceType {
	__r0, _, _ := purego.SyscallN(addr_QWindow_surfaceType_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QSurface__SurfaceType(__r)
}

func (this *QWindow) FormatDefault() *QSurfaceFormat {
	__r, _, _ := purego.SyscallN(addr_QWindow_format_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSurfaceFormat(__r)
}

func (this *QWindow) SizeDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QWindow_size_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QWindow) AccessibleRootDefault() *QAccessibleInterface {
	__r, _, _ := purego.SyscallN(addr_QWindow_accessibleRoot_default, this.h)
	runtime.KeepAlive(this)
	return bindQAccessibleInterface(__r)
}

func (this *QWindow) FocusObjectDefault() *QObject {
	__r, _, _ := purego.SyscallN(addr_QWindow_focusObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QWindow) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QWindow_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QWindow) ExposeEventDefault(arg1 *QExposeEvent) {
	purego.SyscallN(addr_QWindow_exposeEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QWindow) ResizeEventDefault(arg1 *QResizeEvent) {
	purego.SyscallN(addr_QWindow_resizeEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QWindow) PaintEventDefault(arg1 *QPaintEvent) {
	purego.SyscallN(addr_QWindow_paintEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QWindow) MoveEventDefault(arg1 *QMoveEvent) {
	purego.SyscallN(addr_QWindow_moveEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QWindow) FocusInEventDefault(arg1 *QFocusEvent) {
	purego.SyscallN(addr_QWindow_focusInEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QWindow) FocusOutEventDefault(arg1 *QFocusEvent) {
	purego.SyscallN(addr_QWindow_focusOutEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QWindow) ShowEventDefault(arg1 *QShowEvent) {
	purego.SyscallN(addr_QWindow_showEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QWindow) HideEventDefault(arg1 *QHideEvent) {
	purego.SyscallN(addr_QWindow_hideEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QWindow) CloseEventDefault(arg1 *QCloseEvent) {
	purego.SyscallN(addr_QWindow_closeEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QWindow) EventDefault(arg1 *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QWindow_event_default, this.h, uintptr(arg1.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return __r != 0
}

func (this *QWindow) KeyPressEventDefault(arg1 *QKeyEvent) {
	purego.SyscallN(addr_QWindow_keyPressEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QWindow) KeyReleaseEventDefault(arg1 *QKeyEvent) {
	purego.SyscallN(addr_QWindow_keyReleaseEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QWindow) MousePressEventDefault(arg1 *QMouseEvent) {
	purego.SyscallN(addr_QWindow_mousePressEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QWindow) MouseReleaseEventDefault(arg1 *QMouseEvent) {
	purego.SyscallN(addr_QWindow_mouseReleaseEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QWindow) MouseDoubleClickEventDefault(arg1 *QMouseEvent) {
	purego.SyscallN(addr_QWindow_mouseDoubleClickEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QWindow) MouseMoveEventDefault(arg1 *QMouseEvent) {
	purego.SyscallN(addr_QWindow_mouseMoveEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QWindow) WheelEventDefault(arg1 *QWheelEvent) {
	purego.SyscallN(addr_QWindow_wheelEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QWindow) TouchEventDefault(arg1 *QTouchEvent) {
	purego.SyscallN(addr_QWindow_touchEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QWindow) TabletEventDefault(arg1 *QTabletEvent) {
	purego.SyscallN(addr_QWindow_tabletEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QWindow) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QWindow_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QWindow) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QWindow_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QWindow) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QWindow_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QWindow) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QWindow_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QWindow) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QWindow_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QWindow) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QWindow_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QWindow) OnScreenChanged(fn func(*QScreen)) int64 {
	id := registerCallback(func(screen uintptr) { fn(UnsafeNewQScreen(unsafe.Pointer(screen))) })
	pg_QWindow_connect_screenChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QWindow) OnModalityChanged(fn func(WindowModality)) int64 {
	id := registerCallback(func(modality int32) { fn(WindowModality(modality)) })
	pg_QWindow_connect_modalityChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QWindow) OnFlagsChanged(fn func(WindowType)) int64 {
	id := registerCallback(func(flags int32) { fn(WindowType(flags)) })
	pg_QWindow_connect_flagsChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QWindow) OnWindowStateChanged(fn func(WindowState)) int64 {
	id := registerCallback(func(windowState int32) { fn(WindowState(windowState)) })
	pg_QWindow_connect_windowStateChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QWindow) OnWindowTitleChanged(fn func(string)) int64 {
	id := registerCallback(func(title uintptr, title_len int64) { fn(pgStrNoFree(title, title_len)) })
	pg_QWindow_connect_windowTitleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QWindow) OnXChanged(fn func(int)) int64 {
	id := registerCallback(func(arg int32) { fn(int(arg)) })
	pg_QWindow_connect_xChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QWindow) OnYChanged(fn func(int)) int64 {
	id := registerCallback(func(arg int32) { fn(int(arg)) })
	pg_QWindow_connect_yChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QWindow) OnWidthChanged(fn func(int)) int64 {
	id := registerCallback(func(arg int32) { fn(int(arg)) })
	pg_QWindow_connect_widthChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QWindow) OnHeightChanged(fn func(int)) int64 {
	id := registerCallback(func(arg int32) { fn(int(arg)) })
	pg_QWindow_connect_heightChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QWindow) OnMinimumWidthChanged(fn func(int)) int64 {
	id := registerCallback(func(arg int32) { fn(int(arg)) })
	pg_QWindow_connect_minimumWidthChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QWindow) OnMinimumHeightChanged(fn func(int)) int64 {
	id := registerCallback(func(arg int32) { fn(int(arg)) })
	pg_QWindow_connect_minimumHeightChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QWindow) OnMaximumWidthChanged(fn func(int)) int64 {
	id := registerCallback(func(arg int32) { fn(int(arg)) })
	pg_QWindow_connect_maximumWidthChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QWindow) OnMaximumHeightChanged(fn func(int)) int64 {
	id := registerCallback(func(arg int32) { fn(int(arg)) })
	pg_QWindow_connect_maximumHeightChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QWindow) OnSafeAreaMarginsChanged(fn func(*QMargins)) int64 {
	id := registerCallback(func(arg uintptr) { fn(UnsafeNewQMargins(unsafe.Pointer(arg))) })
	pg_QWindow_connect_safeAreaMarginsChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QWindow) OnVisibleChanged(fn func(bool)) int64 {
	id := registerCallback(func(arg int32) { fn(arg != 0) })
	pg_QWindow_connect_visibleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QWindow) OnVisibilityChanged(fn func(QWindow__Visibility)) int64 {
	id := registerCallback(func(visibility int32) { fn(QWindow__Visibility(visibility)) })
	pg_QWindow_connect_visibilityChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QWindow) OnActiveChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QWindow_connect_activeChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QWindow) OnContentOrientationChanged(fn func(ScreenOrientation)) int64 {
	id := registerCallback(func(orientation int32) { fn(ScreenOrientation(orientation)) })
	pg_QWindow_connect_contentOrientationChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QWindow) OnFocusObjectChanged(fn func(*QObject)) int64 {
	id := registerCallback(func(object uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(object))) })
	pg_QWindow_connect_focusObjectChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QWindow) OnOpacityChanged(fn func(float64)) int64 {
	id := registerCallback(func(opacity uintptr) { fn(math.Float64frombits(uint64(opacity))) })
	pg_QWindow_connect_opacityChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QWindow) OnTransientParentChanged(fn func(*QWindow)) int64 {
	id := registerCallback(func(transientParent uintptr) { fn(UnsafeNewQWindow(unsafe.Pointer(transientParent))) })
	pg_QWindow_connect_transientParentChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QWindow) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QWindow_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QWindow) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QWindow_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QWindow) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QWindow_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QWindow) Disconnect(handle int64) {
	pg_QWindow__disconnect(handle)
	unregisterCallback(handle)
}

func QWindow_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case 7339588379958472532:
		return 1
	case 7267508590720539762:
		return 2
	case 5614464919154503228:
		return 3
	case -2255663269888600317:
		return 4
	case 8913927178009815944:
		return 5
	case 3239118852587300383:
		return 6
	case -4302917285655014661:
		return 7
	case 5290442163167458465:
		return 8
	case -9161949150489525396:
		return 9
	case 975043767033552552:
		return 10
	case -385596010576048993:
		return 11
	case -6825369174751916:
		return 12
	case -199718707876719129:
		return 13
	case -1989937050019464783:
		return 14
	case 1237236456312824991:
		return 15
	case 1023791977679793185:
		return 16
	case 772598475950050299:
		return 17
	case -332090898078543319:
		return 18
	case -4415920049297795213:
		return 19
	case -4029073905697023887:
		return 20
	case -4255494004099806581:
		return 21
	case 1919988340934511316:
		return 22
	case -2204038732330958816:
		return 23
	case -212068866860701141:
		return 24
	case -2039166586246443706:
		return 25
	case 5217623514630858297:
		return 26
	case -5409927581812374498:
		return 27
	case -1185315367526857833:
		return 28
	case 3769868229123868474:
		return 29
	case 7379029982170503008:
		return 30
	case -568369772982215962:
		return 31
	}
	return -1
}

func QWindow_enable(h uintptr, idx int32) { pg_QWindow__set_override(h, idx) }

func (this *QWindow) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QWindow) OnSurfaceType(fn func() QSurface__SurfaceType) {
	registerVirt(this.UnsafeGoID(), 7339588379958472532, func() int32 { return int32(fn()) }) // slot "surfaceType"
}

func (this *QWindow) OnFormat(fn func() *QSurfaceFormat) {
	registerVirt(this.UnsafeGoID(), 7267508590720539762, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "format"
}

func (this *QWindow) OnSize(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 5614464919154503228, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "size"
}

func (this *QWindow) OnAccessibleRoot(fn func() *QAccessibleInterface) {
	registerVirt(this.UnsafeGoID(), -2255663269888600317, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "accessibleRoot"
}

func (this *QWindow) OnFocusObject(fn func() *QObject) {
	registerVirt(this.UnsafeGoID(), 8913927178009815944, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "focusObject"
}

func (this *QWindow) OnExposeEvent(fn func(*QExposeEvent)) {
	registerVirt(this.UnsafeGoID(), 3239118852587300383, func(arg1 uintptr) { fn(UnsafeNewQExposeEvent(unsafe.Pointer(arg1))) }) // slot "exposeEvent"
}

func (this *QWindow) OnResizeEvent(fn func(*QResizeEvent)) {
	registerVirt(this.UnsafeGoID(), -4302917285655014661, func(arg1 uintptr) { fn(UnsafeNewQResizeEvent(unsafe.Pointer(arg1))) }) // slot "resizeEvent"
}

func (this *QWindow) OnPaintEvent(fn func(*QPaintEvent)) {
	registerVirt(this.UnsafeGoID(), 5290442163167458465, func(arg1 uintptr) { fn(UnsafeNewQPaintEvent(unsafe.Pointer(arg1))) }) // slot "paintEvent"
}

func (this *QWindow) OnMoveEvent(fn func(*QMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -9161949150489525396, func(arg1 uintptr) { fn(UnsafeNewQMoveEvent(unsafe.Pointer(arg1))) }) // slot "moveEvent"
}

func (this *QWindow) OnFocusInEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), 975043767033552552, func(arg1 uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(arg1))) }) // slot "focusInEvent"
}

func (this *QWindow) OnFocusOutEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), -385596010576048993, func(arg1 uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(arg1))) }) // slot "focusOutEvent"
}

func (this *QWindow) OnShowEvent(fn func(*QShowEvent)) {
	registerVirt(this.UnsafeGoID(), -6825369174751916, func(arg1 uintptr) { fn(UnsafeNewQShowEvent(unsafe.Pointer(arg1))) }) // slot "showEvent"
}

func (this *QWindow) OnHideEvent(fn func(*QHideEvent)) {
	registerVirt(this.UnsafeGoID(), -199718707876719129, func(arg1 uintptr) { fn(UnsafeNewQHideEvent(unsafe.Pointer(arg1))) }) // slot "hideEvent"
}

func (this *QWindow) OnCloseEvent(fn func(*QCloseEvent)) {
	registerVirt(this.UnsafeGoID(), -1989937050019464783, func(arg1 uintptr) { fn(UnsafeNewQCloseEvent(unsafe.Pointer(arg1))) }) // slot "closeEvent"
}

func (this *QWindow) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(arg1 uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(arg1)))) }) // slot "event"
}

func (this *QWindow) OnKeyPressEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 1023791977679793185, func(arg1 uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(arg1))) }) // slot "keyPressEvent"
}

func (this *QWindow) OnKeyReleaseEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 772598475950050299, func(arg1 uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(arg1))) }) // slot "keyReleaseEvent"
}

func (this *QWindow) OnMousePressEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -332090898078543319, func(arg1 uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(arg1))) }) // slot "mousePressEvent"
}

func (this *QWindow) OnMouseReleaseEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4415920049297795213, func(arg1 uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(arg1))) }) // slot "mouseReleaseEvent"
}

func (this *QWindow) OnMouseDoubleClickEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4029073905697023887, func(arg1 uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(arg1))) }) // slot "mouseDoubleClickEvent"
}

func (this *QWindow) OnMouseMoveEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4255494004099806581, func(arg1 uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(arg1))) }) // slot "mouseMoveEvent"
}

func (this *QWindow) OnWheelEvent(fn func(*QWheelEvent)) {
	registerVirt(this.UnsafeGoID(), 1919988340934511316, func(arg1 uintptr) { fn(UnsafeNewQWheelEvent(unsafe.Pointer(arg1))) }) // slot "wheelEvent"
}

func (this *QWindow) OnTouchEvent(fn func(*QTouchEvent)) {
	registerVirt(this.UnsafeGoID(), -2204038732330958816, func(arg1 uintptr) { fn(UnsafeNewQTouchEvent(unsafe.Pointer(arg1))) }) // slot "touchEvent"
}

func (this *QWindow) OnTabletEvent(fn func(*QTabletEvent)) {
	registerVirt(this.UnsafeGoID(), -212068866860701141, func(arg1 uintptr) { fn(UnsafeNewQTabletEvent(unsafe.Pointer(arg1))) }) // slot "tabletEvent"
}

func (this *QWindow) OnNativeEvent(fn func(*QByteArray, unsafe.Pointer, *int64) bool) {
	registerVirt(this.UnsafeGoID(), -2039166586246443706, func(eventType uintptr, message uintptr, result uintptr) int32 { return b2i(fn(UnsafeNewQByteArray(unsafe.Pointer(eventType)), unsafe.Pointer(message), (*int64)(unsafe.Pointer(result)))) }) // slot "nativeEvent"
}

func (this *QWindow) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QWindow) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QWindow) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QWindow) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QWindow) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QWindow) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QWindow) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QWindow_delete(this.h)
	}
	this.UnsafeRelease()
}
