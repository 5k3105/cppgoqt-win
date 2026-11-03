package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QWidget_base func(uintptr) uintptr
	addr_QWidget_new uintptr
	addr_QWidget_new2 uintptr
	addr_QWidget_new3 uintptr
	addr_QWidget_tr uintptr
	addr_QWidget_devType uintptr
	addr_QWidget_winId uintptr
	addr_QWidget_createWinId uintptr
	addr_QWidget_internalWinId uintptr
	addr_QWidget_effectiveWinId uintptr
	addr_QWidget_style uintptr
	addr_QWidget_setStyle uintptr
	addr_QWidget_isTopLevel uintptr
	addr_QWidget_isWindow uintptr
	addr_QWidget_isModal uintptr
	addr_QWidget_windowModality uintptr
	addr_QWidget_setWindowModality uintptr
	addr_QWidget_isEnabled uintptr
	addr_QWidget_isEnabledTo uintptr
	addr_QWidget_setEnabled uintptr
	addr_QWidget_setDisabled uintptr
	addr_QWidget_setWindowModified uintptr
	addr_QWidget_frameGeometry uintptr
	addr_QWidget_geometry uintptr
	addr_QWidget_normalGeometry uintptr
	addr_QWidget_x uintptr
	addr_QWidget_y uintptr
	addr_QWidget_pos uintptr
	addr_QWidget_frameSize uintptr
	addr_QWidget_size uintptr
	addr_QWidget_width uintptr
	addr_QWidget_height uintptr
	addr_QWidget_rect uintptr
	addr_QWidget_childrenRect uintptr
	addr_QWidget_childrenRegion uintptr
	addr_QWidget_minimumSize uintptr
	addr_QWidget_maximumSize uintptr
	addr_QWidget_minimumWidth uintptr
	addr_QWidget_minimumHeight uintptr
	addr_QWidget_maximumWidth uintptr
	addr_QWidget_maximumHeight uintptr
	addr_QWidget_setMinimumSize uintptr
	addr_QWidget_setMinimumSize2 uintptr
	addr_QWidget_setMaximumSize uintptr
	addr_QWidget_setMaximumSize2 uintptr
	addr_QWidget_setMinimumWidth uintptr
	addr_QWidget_setMinimumHeight uintptr
	addr_QWidget_setMaximumWidth uintptr
	addr_QWidget_setMaximumHeight uintptr
	addr_QWidget_sizeIncrement uintptr
	addr_QWidget_setSizeIncrement uintptr
	addr_QWidget_setSizeIncrement2 uintptr
	addr_QWidget_baseSize uintptr
	addr_QWidget_setBaseSize uintptr
	addr_QWidget_setBaseSize2 uintptr
	addr_QWidget_setFixedSize uintptr
	addr_QWidget_setFixedSize2 uintptr
	addr_QWidget_setFixedWidth uintptr
	addr_QWidget_setFixedHeight uintptr
	addr_QWidget_mapToGlobal uintptr
	addr_QWidget_mapToGlobal2 uintptr
	addr_QWidget_mapFromGlobal uintptr
	addr_QWidget_mapFromGlobal2 uintptr
	addr_QWidget_mapToParent uintptr
	addr_QWidget_mapToParent2 uintptr
	addr_QWidget_mapFromParent uintptr
	addr_QWidget_mapFromParent2 uintptr
	addr_QWidget_mapTo uintptr
	addr_QWidget_mapTo2 uintptr
	addr_QWidget_mapFrom uintptr
	addr_QWidget_mapFrom2 uintptr
	addr_QWidget_window uintptr
	addr_QWidget_nativeParentWidget uintptr
	addr_QWidget_topLevelWidget uintptr
	addr_QWidget_palette uintptr
	addr_QWidget_setPalette uintptr
	addr_QWidget_setBackgroundRole uintptr
	addr_QWidget_backgroundRole uintptr
	addr_QWidget_setForegroundRole uintptr
	addr_QWidget_foregroundRole uintptr
	addr_QWidget_font uintptr
	addr_QWidget_setFont uintptr
	addr_QWidget_fontMetrics uintptr
	addr_QWidget_fontInfo uintptr
	addr_QWidget_cursor uintptr
	addr_QWidget_setCursor uintptr
	addr_QWidget_unsetCursor uintptr
	addr_QWidget_setMouseTracking uintptr
	addr_QWidget_hasMouseTracking uintptr
	addr_QWidget_underMouse uintptr
	addr_QWidget_setTabletTracking uintptr
	addr_QWidget_hasTabletTracking uintptr
	addr_QWidget_setMask uintptr
	addr_QWidget_setMask2 uintptr
	addr_QWidget_mask uintptr
	addr_QWidget_clearMask uintptr
	addr_QWidget_render uintptr
	addr_QWidget_render2 uintptr
	addr_QWidget_grab uintptr
	addr_QWidget_graphicsEffect uintptr
	addr_QWidget_setGraphicsEffect uintptr
	addr_QWidget_grabGesture uintptr
	addr_QWidget_ungrabGesture uintptr
	addr_QWidget_setWindowTitle uintptr
	addr_QWidget_setStyleSheet uintptr
	addr_QWidget_styleSheet uintptr
	addr_QWidget_windowTitle uintptr
	addr_QWidget_setWindowIcon uintptr
	addr_QWidget_windowIcon uintptr
	addr_QWidget_setWindowIconText uintptr
	addr_QWidget_windowIconText uintptr
	addr_QWidget_setWindowRole uintptr
	addr_QWidget_windowRole uintptr
	addr_QWidget_setWindowFilePath uintptr
	addr_QWidget_windowFilePath uintptr
	pg_QWidget_setWindowOpacity func(uintptr, float64)
	pg_QWidget_windowOpacity func(uintptr) float64
	addr_QWidget_isWindowModified uintptr
	addr_QWidget_setToolTip uintptr
	addr_QWidget_toolTip uintptr
	addr_QWidget_setToolTipDuration uintptr
	addr_QWidget_toolTipDuration uintptr
	addr_QWidget_setStatusTip uintptr
	addr_QWidget_statusTip uintptr
	addr_QWidget_setWhatsThis uintptr
	addr_QWidget_whatsThis uintptr
	addr_QWidget_accessibleName uintptr
	addr_QWidget_setAccessibleName uintptr
	addr_QWidget_accessibleDescription uintptr
	addr_QWidget_setAccessibleDescription uintptr
	addr_QWidget_accessibleIdentifier uintptr
	addr_QWidget_setAccessibleIdentifier uintptr
	addr_QWidget_setLayoutDirection uintptr
	addr_QWidget_layoutDirection uintptr
	addr_QWidget_unsetLayoutDirection uintptr
	addr_QWidget_setLocale uintptr
	addr_QWidget_locale uintptr
	addr_QWidget_unsetLocale uintptr
	addr_QWidget_isRightToLeft uintptr
	addr_QWidget_isLeftToRight uintptr
	addr_QWidget_setFocus uintptr
	addr_QWidget_isActiveWindow uintptr
	addr_QWidget_activateWindow uintptr
	addr_QWidget_clearFocus uintptr
	addr_QWidget_setFocus2 uintptr
	addr_QWidget_focusPolicy uintptr
	addr_QWidget_setFocusPolicy uintptr
	addr_QWidget_hasFocus uintptr
	addr_QWidget_setTabOrder uintptr
	addr_QWidget_setFocusProxy uintptr
	addr_QWidget_focusProxy uintptr
	addr_QWidget_contextMenuPolicy uintptr
	addr_QWidget_setContextMenuPolicy uintptr
	addr_QWidget_grabMouse uintptr
	addr_QWidget_grabMouse2 uintptr
	addr_QWidget_releaseMouse uintptr
	addr_QWidget_grabKeyboard uintptr
	addr_QWidget_releaseKeyboard uintptr
	addr_QWidget_grabShortcut uintptr
	addr_QWidget_releaseShortcut uintptr
	addr_QWidget_setShortcutEnabled uintptr
	addr_QWidget_setShortcutAutoRepeat uintptr
	addr_QWidget_mouseGrabber uintptr
	addr_QWidget_keyboardGrabber uintptr
	addr_QWidget_updatesEnabled uintptr
	addr_QWidget_setUpdatesEnabled uintptr
	addr_QWidget_graphicsProxyWidget uintptr
	addr_QWidget_update uintptr
	addr_QWidget_repaint uintptr
	addr_QWidget_update2 uintptr
	addr_QWidget_update3 uintptr
	addr_QWidget_update4 uintptr
	addr_QWidget_repaint2 uintptr
	addr_QWidget_repaint3 uintptr
	addr_QWidget_repaint4 uintptr
	addr_QWidget_setVisible uintptr
	addr_QWidget_setHidden uintptr
	addr_QWidget_show uintptr
	addr_QWidget_hide uintptr
	addr_QWidget_showMinimized uintptr
	addr_QWidget_showMaximized uintptr
	addr_QWidget_showFullScreen uintptr
	addr_QWidget_showNormal uintptr
	addr_QWidget_close uintptr
	addr_QWidget_raise uintptr
	addr_QWidget_lower uintptr
	addr_QWidget_stackUnder uintptr
	addr_QWidget_move uintptr
	addr_QWidget_move2 uintptr
	addr_QWidget_resize uintptr
	addr_QWidget_resize2 uintptr
	addr_QWidget_setGeometry uintptr
	addr_QWidget_setGeometry2 uintptr
	addr_QWidget_saveGeometry uintptr
	addr_QWidget_restoreGeometry uintptr
	addr_QWidget_adjustSize uintptr
	addr_QWidget_isVisible uintptr
	addr_QWidget_isVisibleTo uintptr
	addr_QWidget_isHidden uintptr
	addr_QWidget_isMinimized uintptr
	addr_QWidget_isMaximized uintptr
	addr_QWidget_isFullScreen uintptr
	addr_QWidget_windowState uintptr
	addr_QWidget_setWindowState uintptr
	addr_QWidget_overrideWindowState uintptr
	addr_QWidget_sizeHint uintptr
	addr_QWidget_minimumSizeHint uintptr
	addr_QWidget_sizePolicy uintptr
	addr_QWidget_setSizePolicy uintptr
	addr_QWidget_setSizePolicy2 uintptr
	addr_QWidget_heightForWidth uintptr
	addr_QWidget_hasHeightForWidth uintptr
	addr_QWidget_visibleRegion uintptr
	addr_QWidget_setContentsMargins uintptr
	addr_QWidget_setContentsMargins2 uintptr
	addr_QWidget_contentsMargins uintptr
	addr_QWidget_contentsRect uintptr
	addr_QWidget_layout uintptr
	addr_QWidget_setLayout uintptr
	addr_QWidget_updateGeometry uintptr
	addr_QWidget_setParent uintptr
	addr_QWidget_setParent2 uintptr
	addr_QWidget_scroll uintptr
	addr_QWidget_scroll2 uintptr
	addr_QWidget_focusWidget uintptr
	addr_QWidget_nextInFocusChain uintptr
	addr_QWidget_previousInFocusChain uintptr
	addr_QWidget_acceptDrops uintptr
	addr_QWidget_setAcceptDrops uintptr
	addr_QWidget_addAction uintptr
	addr_QWidget_addActions uintptr
	addr_QWidget_insertActions uintptr
	addr_QWidget_insertAction uintptr
	addr_QWidget_removeAction uintptr
	addr_QWidget_actions uintptr
	addr_QWidget_addAction2 uintptr
	addr_QWidget_addAction3 uintptr
	addr_QWidget_addAction4 uintptr
	addr_QWidget_addAction5 uintptr
	addr_QWidget_addAction6 uintptr
	addr_QWidget_addAction7 uintptr
	addr_QWidget_addAction8 uintptr
	addr_QWidget_addAction9 uintptr
	addr_QWidget_parentWidget uintptr
	addr_QWidget_setWindowFlags uintptr
	addr_QWidget_windowFlags uintptr
	addr_QWidget_setWindowFlag uintptr
	addr_QWidget_overrideWindowFlags uintptr
	addr_QWidget_windowType uintptr
	addr_QWidget_find uintptr
	addr_QWidget_childAt uintptr
	addr_QWidget_childAt2 uintptr
	addr_QWidget_childAt3 uintptr
	addr_QWidget_setAttribute uintptr
	addr_QWidget_testAttribute uintptr
	addr_QWidget_paintEngine uintptr
	addr_QWidget_ensurePolished uintptr
	addr_QWidget_isAncestorOf uintptr
	addr_QWidget_autoFillBackground uintptr
	addr_QWidget_setAutoFillBackground uintptr
	addr_QWidget_backingStore uintptr
	addr_QWidget_windowHandle uintptr
	addr_QWidget_screen uintptr
	addr_QWidget_setScreen uintptr
	addr_QWidget_createWindowContainer uintptr
	addr_QWidget_windowTitleChanged uintptr
	addr_QWidget_windowIconChanged uintptr
	addr_QWidget_windowIconTextChanged uintptr
	addr_QWidget_customContextMenuRequested uintptr
	addr_QWidget_inputMethodQuery uintptr
	addr_QWidget_inputMethodHints uintptr
	addr_QWidget_setInputMethodHints uintptr
	addr_QWidget_updateMicroFocus uintptr
	addr_QWidget_create uintptr
	addr_QWidget_destroy uintptr
	addr_QWidget_focusNextChild uintptr
	addr_QWidget_focusPreviousChild uintptr
	addr_QWidget_connect uintptr
	addr_QWidget_connect2 uintptr
	addr_QWidget_sender uintptr
	addr_QWidget_senderSignalIndex uintptr
	addr_QWidget_receivers uintptr
	addr_QWidget_isSignalConnected uintptr
	addr_QWidget_paintingActive uintptr
	addr_QWidget_widthMM uintptr
	addr_QWidget_heightMM uintptr
	addr_QWidget_logicalDpiX uintptr
	addr_QWidget_logicalDpiY uintptr
	addr_QWidget_physicalDpiX uintptr
	addr_QWidget_physicalDpiY uintptr
	pg_QWidget_devicePixelRatio func(uintptr) float64
	pg_QWidget_devicePixelRatioF func(uintptr) float64
	addr_QWidget_colorCount uintptr
	addr_QWidget_depth uintptr
	pg_QWidget_devicePixelRatioFScale func() float64
	pg_QWidget_encodeMetricF func(int32, float64) int32
	pg_QWidget_getDecodedMetricF func(uintptr, int32, int32) float64
	addr_QWidget_tr2 uintptr
	addr_QWidget_tr3 uintptr
	addr_QWidget_render3 uintptr
	addr_QWidget_render4 uintptr
	addr_QWidget_render5 uintptr
	addr_QWidget_render6 uintptr
	addr_QWidget_render7 uintptr
	addr_QWidget_render8 uintptr
	addr_QWidget_grab2 uintptr
	addr_QWidget_grabGesture2 uintptr
	addr_QWidget_grabShortcut2 uintptr
	addr_QWidget_setShortcutEnabled2 uintptr
	addr_QWidget_setShortcutAutoRepeat2 uintptr
	addr_QWidget_addAction10 uintptr
	addr_QWidget_addAction11 uintptr
	addr_QWidget_addAction12 uintptr
	addr_QWidget_addAction13 uintptr
	addr_QWidget_setWindowFlag2 uintptr
	addr_QWidget_setAttribute2 uintptr
	addr_QWidget_createWindowContainer2 uintptr
	addr_QWidget_createWindowContainer3 uintptr
	addr_QWidget_updateMicroFocus2 uintptr
	addr_QWidget_create2 uintptr
	addr_QWidget_create3 uintptr
	addr_QWidget_create4 uintptr
	addr_QWidget_destroy2 uintptr
	addr_QWidget_destroy3 uintptr
	addr_QWidget_connect3 uintptr
	addr_QWidget_connect4 uintptr
	addr_QWidget_asQPaintDevice uintptr
	addr_QWidget_devType_default uintptr
	addr_QWidget_setVisible_default uintptr
	addr_QWidget_sizeHint_default uintptr
	addr_QWidget_minimumSizeHint_default uintptr
	addr_QWidget_heightForWidth_default uintptr
	addr_QWidget_hasHeightForWidth_default uintptr
	addr_QWidget_paintEngine_default uintptr
	addr_QWidget_inputMethodQuery_default uintptr
	addr_QWidget_metaObject_default uintptr
	addr_QWidget_event_default uintptr
	addr_QWidget_mousePressEvent_default uintptr
	addr_QWidget_mouseReleaseEvent_default uintptr
	addr_QWidget_mouseDoubleClickEvent_default uintptr
	addr_QWidget_mouseMoveEvent_default uintptr
	addr_QWidget_wheelEvent_default uintptr
	addr_QWidget_keyPressEvent_default uintptr
	addr_QWidget_keyReleaseEvent_default uintptr
	addr_QWidget_focusInEvent_default uintptr
	addr_QWidget_focusOutEvent_default uintptr
	addr_QWidget_enterEvent_default uintptr
	addr_QWidget_leaveEvent_default uintptr
	addr_QWidget_paintEvent_default uintptr
	addr_QWidget_moveEvent_default uintptr
	addr_QWidget_resizeEvent_default uintptr
	addr_QWidget_closeEvent_default uintptr
	addr_QWidget_contextMenuEvent_default uintptr
	addr_QWidget_tabletEvent_default uintptr
	addr_QWidget_actionEvent_default uintptr
	addr_QWidget_dragEnterEvent_default uintptr
	addr_QWidget_dragMoveEvent_default uintptr
	addr_QWidget_dragLeaveEvent_default uintptr
	addr_QWidget_dropEvent_default uintptr
	addr_QWidget_showEvent_default uintptr
	addr_QWidget_hideEvent_default uintptr
	addr_QWidget_changeEvent_default uintptr
	addr_QWidget_metric_default uintptr
	addr_QWidget_initPainter_default uintptr
	addr_QWidget_redirected_default uintptr
	addr_QWidget_sharedPainter_default uintptr
	addr_QWidget_inputMethodEvent_default uintptr
	addr_QWidget_focusNextPrevChild_default uintptr
	addr_QWidget_eventFilter_default uintptr
	addr_QWidget_timerEvent_default uintptr
	addr_QWidget_childEvent_default uintptr
	addr_QWidget_customEvent_default uintptr
	addr_QWidget_connectNotify_default uintptr
	addr_QWidget_disconnectNotify_default uintptr
	pg_QWidget_connect_windowTitleChanged func(uintptr, int64)
	pg_QWidget_connect_windowIconChanged func(uintptr, int64)
	pg_QWidget_connect_windowIconTextChanged func(uintptr, int64)
	pg_QWidget_connect_customContextMenuRequested func(uintptr, int64)
	pg_QWidget_connect_destroyed func(uintptr, int64)
	pg_QWidget_connect_destroyed2 func(uintptr, int64)
	pg_QWidget_connect_objectNameChanged func(uintptr, int64)
	pg_QWidget__disconnect func(int64)
	pg_QWidget__set_override func(uintptr, int32)
	pg_QWidget_delete func(uintptr)
)

func init() {
	reg(&pg_QWidget_base, "QWidget_base")
	symaddr(&addr_QWidget_new, "QWidget_new")
	symaddr(&addr_QWidget_new2, "QWidget_new2")
	symaddr(&addr_QWidget_new3, "QWidget_new3")
	symaddr(&addr_QWidget_tr, "QWidget_tr")
	symaddr(&addr_QWidget_devType, "QWidget_devType")
	symaddr(&addr_QWidget_winId, "QWidget_winId")
	symaddr(&addr_QWidget_createWinId, "QWidget_createWinId")
	symaddr(&addr_QWidget_internalWinId, "QWidget_internalWinId")
	symaddr(&addr_QWidget_effectiveWinId, "QWidget_effectiveWinId")
	symaddr(&addr_QWidget_style, "QWidget_style")
	symaddr(&addr_QWidget_setStyle, "QWidget_setStyle")
	symaddr(&addr_QWidget_isTopLevel, "QWidget_isTopLevel")
	symaddr(&addr_QWidget_isWindow, "QWidget_isWindow")
	symaddr(&addr_QWidget_isModal, "QWidget_isModal")
	symaddr(&addr_QWidget_windowModality, "QWidget_windowModality")
	symaddr(&addr_QWidget_setWindowModality, "QWidget_setWindowModality")
	symaddr(&addr_QWidget_isEnabled, "QWidget_isEnabled")
	symaddr(&addr_QWidget_isEnabledTo, "QWidget_isEnabledTo")
	symaddr(&addr_QWidget_setEnabled, "QWidget_setEnabled")
	symaddr(&addr_QWidget_setDisabled, "QWidget_setDisabled")
	symaddr(&addr_QWidget_setWindowModified, "QWidget_setWindowModified")
	symaddr(&addr_QWidget_frameGeometry, "QWidget_frameGeometry")
	symaddr(&addr_QWidget_geometry, "QWidget_geometry")
	symaddr(&addr_QWidget_normalGeometry, "QWidget_normalGeometry")
	symaddr(&addr_QWidget_x, "QWidget_x")
	symaddr(&addr_QWidget_y, "QWidget_y")
	symaddr(&addr_QWidget_pos, "QWidget_pos")
	symaddr(&addr_QWidget_frameSize, "QWidget_frameSize")
	symaddr(&addr_QWidget_size, "QWidget_size")
	symaddr(&addr_QWidget_width, "QWidget_width")
	symaddr(&addr_QWidget_height, "QWidget_height")
	symaddr(&addr_QWidget_rect, "QWidget_rect")
	symaddr(&addr_QWidget_childrenRect, "QWidget_childrenRect")
	symaddr(&addr_QWidget_childrenRegion, "QWidget_childrenRegion")
	symaddr(&addr_QWidget_minimumSize, "QWidget_minimumSize")
	symaddr(&addr_QWidget_maximumSize, "QWidget_maximumSize")
	symaddr(&addr_QWidget_minimumWidth, "QWidget_minimumWidth")
	symaddr(&addr_QWidget_minimumHeight, "QWidget_minimumHeight")
	symaddr(&addr_QWidget_maximumWidth, "QWidget_maximumWidth")
	symaddr(&addr_QWidget_maximumHeight, "QWidget_maximumHeight")
	symaddr(&addr_QWidget_setMinimumSize, "QWidget_setMinimumSize")
	symaddr(&addr_QWidget_setMinimumSize2, "QWidget_setMinimumSize2")
	symaddr(&addr_QWidget_setMaximumSize, "QWidget_setMaximumSize")
	symaddr(&addr_QWidget_setMaximumSize2, "QWidget_setMaximumSize2")
	symaddr(&addr_QWidget_setMinimumWidth, "QWidget_setMinimumWidth")
	symaddr(&addr_QWidget_setMinimumHeight, "QWidget_setMinimumHeight")
	symaddr(&addr_QWidget_setMaximumWidth, "QWidget_setMaximumWidth")
	symaddr(&addr_QWidget_setMaximumHeight, "QWidget_setMaximumHeight")
	symaddr(&addr_QWidget_sizeIncrement, "QWidget_sizeIncrement")
	symaddr(&addr_QWidget_setSizeIncrement, "QWidget_setSizeIncrement")
	symaddr(&addr_QWidget_setSizeIncrement2, "QWidget_setSizeIncrement2")
	symaddr(&addr_QWidget_baseSize, "QWidget_baseSize")
	symaddr(&addr_QWidget_setBaseSize, "QWidget_setBaseSize")
	symaddr(&addr_QWidget_setBaseSize2, "QWidget_setBaseSize2")
	symaddr(&addr_QWidget_setFixedSize, "QWidget_setFixedSize")
	symaddr(&addr_QWidget_setFixedSize2, "QWidget_setFixedSize2")
	symaddr(&addr_QWidget_setFixedWidth, "QWidget_setFixedWidth")
	symaddr(&addr_QWidget_setFixedHeight, "QWidget_setFixedHeight")
	symaddr(&addr_QWidget_mapToGlobal, "QWidget_mapToGlobal")
	symaddr(&addr_QWidget_mapToGlobal2, "QWidget_mapToGlobal2")
	symaddr(&addr_QWidget_mapFromGlobal, "QWidget_mapFromGlobal")
	symaddr(&addr_QWidget_mapFromGlobal2, "QWidget_mapFromGlobal2")
	symaddr(&addr_QWidget_mapToParent, "QWidget_mapToParent")
	symaddr(&addr_QWidget_mapToParent2, "QWidget_mapToParent2")
	symaddr(&addr_QWidget_mapFromParent, "QWidget_mapFromParent")
	symaddr(&addr_QWidget_mapFromParent2, "QWidget_mapFromParent2")
	symaddr(&addr_QWidget_mapTo, "QWidget_mapTo")
	symaddr(&addr_QWidget_mapTo2, "QWidget_mapTo2")
	symaddr(&addr_QWidget_mapFrom, "QWidget_mapFrom")
	symaddr(&addr_QWidget_mapFrom2, "QWidget_mapFrom2")
	symaddr(&addr_QWidget_window, "QWidget_window")
	symaddr(&addr_QWidget_nativeParentWidget, "QWidget_nativeParentWidget")
	symaddr(&addr_QWidget_topLevelWidget, "QWidget_topLevelWidget")
	symaddr(&addr_QWidget_palette, "QWidget_palette")
	symaddr(&addr_QWidget_setPalette, "QWidget_setPalette")
	symaddr(&addr_QWidget_setBackgroundRole, "QWidget_setBackgroundRole")
	symaddr(&addr_QWidget_backgroundRole, "QWidget_backgroundRole")
	symaddr(&addr_QWidget_setForegroundRole, "QWidget_setForegroundRole")
	symaddr(&addr_QWidget_foregroundRole, "QWidget_foregroundRole")
	symaddr(&addr_QWidget_font, "QWidget_font")
	symaddr(&addr_QWidget_setFont, "QWidget_setFont")
	symaddr(&addr_QWidget_fontMetrics, "QWidget_fontMetrics")
	symaddr(&addr_QWidget_fontInfo, "QWidget_fontInfo")
	symaddr(&addr_QWidget_cursor, "QWidget_cursor")
	symaddr(&addr_QWidget_setCursor, "QWidget_setCursor")
	symaddr(&addr_QWidget_unsetCursor, "QWidget_unsetCursor")
	symaddr(&addr_QWidget_setMouseTracking, "QWidget_setMouseTracking")
	symaddr(&addr_QWidget_hasMouseTracking, "QWidget_hasMouseTracking")
	symaddr(&addr_QWidget_underMouse, "QWidget_underMouse")
	symaddr(&addr_QWidget_setTabletTracking, "QWidget_setTabletTracking")
	symaddr(&addr_QWidget_hasTabletTracking, "QWidget_hasTabletTracking")
	symaddr(&addr_QWidget_setMask, "QWidget_setMask")
	symaddr(&addr_QWidget_setMask2, "QWidget_setMask2")
	symaddr(&addr_QWidget_mask, "QWidget_mask")
	symaddr(&addr_QWidget_clearMask, "QWidget_clearMask")
	symaddr(&addr_QWidget_render, "QWidget_render")
	symaddr(&addr_QWidget_render2, "QWidget_render2")
	symaddr(&addr_QWidget_grab, "QWidget_grab")
	symaddr(&addr_QWidget_graphicsEffect, "QWidget_graphicsEffect")
	symaddr(&addr_QWidget_setGraphicsEffect, "QWidget_setGraphicsEffect")
	symaddr(&addr_QWidget_grabGesture, "QWidget_grabGesture")
	symaddr(&addr_QWidget_ungrabGesture, "QWidget_ungrabGesture")
	symaddr(&addr_QWidget_setWindowTitle, "QWidget_setWindowTitle")
	symaddr(&addr_QWidget_setStyleSheet, "QWidget_setStyleSheet")
	symaddr(&addr_QWidget_styleSheet, "QWidget_styleSheet")
	symaddr(&addr_QWidget_windowTitle, "QWidget_windowTitle")
	symaddr(&addr_QWidget_setWindowIcon, "QWidget_setWindowIcon")
	symaddr(&addr_QWidget_windowIcon, "QWidget_windowIcon")
	symaddr(&addr_QWidget_setWindowIconText, "QWidget_setWindowIconText")
	symaddr(&addr_QWidget_windowIconText, "QWidget_windowIconText")
	symaddr(&addr_QWidget_setWindowRole, "QWidget_setWindowRole")
	symaddr(&addr_QWidget_windowRole, "QWidget_windowRole")
	symaddr(&addr_QWidget_setWindowFilePath, "QWidget_setWindowFilePath")
	symaddr(&addr_QWidget_windowFilePath, "QWidget_windowFilePath")
	reg(&pg_QWidget_setWindowOpacity, "QWidget_setWindowOpacity")
	reg(&pg_QWidget_windowOpacity, "QWidget_windowOpacity")
	symaddr(&addr_QWidget_isWindowModified, "QWidget_isWindowModified")
	symaddr(&addr_QWidget_setToolTip, "QWidget_setToolTip")
	symaddr(&addr_QWidget_toolTip, "QWidget_toolTip")
	symaddr(&addr_QWidget_setToolTipDuration, "QWidget_setToolTipDuration")
	symaddr(&addr_QWidget_toolTipDuration, "QWidget_toolTipDuration")
	symaddr(&addr_QWidget_setStatusTip, "QWidget_setStatusTip")
	symaddr(&addr_QWidget_statusTip, "QWidget_statusTip")
	symaddr(&addr_QWidget_setWhatsThis, "QWidget_setWhatsThis")
	symaddr(&addr_QWidget_whatsThis, "QWidget_whatsThis")
	symaddr(&addr_QWidget_accessibleName, "QWidget_accessibleName")
	symaddr(&addr_QWidget_setAccessibleName, "QWidget_setAccessibleName")
	symaddr(&addr_QWidget_accessibleDescription, "QWidget_accessibleDescription")
	symaddr(&addr_QWidget_setAccessibleDescription, "QWidget_setAccessibleDescription")
	symaddr(&addr_QWidget_accessibleIdentifier, "QWidget_accessibleIdentifier")
	symaddr(&addr_QWidget_setAccessibleIdentifier, "QWidget_setAccessibleIdentifier")
	symaddr(&addr_QWidget_setLayoutDirection, "QWidget_setLayoutDirection")
	symaddr(&addr_QWidget_layoutDirection, "QWidget_layoutDirection")
	symaddr(&addr_QWidget_unsetLayoutDirection, "QWidget_unsetLayoutDirection")
	symaddr(&addr_QWidget_setLocale, "QWidget_setLocale")
	symaddr(&addr_QWidget_locale, "QWidget_locale")
	symaddr(&addr_QWidget_unsetLocale, "QWidget_unsetLocale")
	symaddr(&addr_QWidget_isRightToLeft, "QWidget_isRightToLeft")
	symaddr(&addr_QWidget_isLeftToRight, "QWidget_isLeftToRight")
	symaddr(&addr_QWidget_setFocus, "QWidget_setFocus")
	symaddr(&addr_QWidget_isActiveWindow, "QWidget_isActiveWindow")
	symaddr(&addr_QWidget_activateWindow, "QWidget_activateWindow")
	symaddr(&addr_QWidget_clearFocus, "QWidget_clearFocus")
	symaddr(&addr_QWidget_setFocus2, "QWidget_setFocus2")
	symaddr(&addr_QWidget_focusPolicy, "QWidget_focusPolicy")
	symaddr(&addr_QWidget_setFocusPolicy, "QWidget_setFocusPolicy")
	symaddr(&addr_QWidget_hasFocus, "QWidget_hasFocus")
	symaddr(&addr_QWidget_setTabOrder, "QWidget_setTabOrder")
	symaddr(&addr_QWidget_setFocusProxy, "QWidget_setFocusProxy")
	symaddr(&addr_QWidget_focusProxy, "QWidget_focusProxy")
	symaddr(&addr_QWidget_contextMenuPolicy, "QWidget_contextMenuPolicy")
	symaddr(&addr_QWidget_setContextMenuPolicy, "QWidget_setContextMenuPolicy")
	symaddr(&addr_QWidget_grabMouse, "QWidget_grabMouse")
	symaddr(&addr_QWidget_grabMouse2, "QWidget_grabMouse2")
	symaddr(&addr_QWidget_releaseMouse, "QWidget_releaseMouse")
	symaddr(&addr_QWidget_grabKeyboard, "QWidget_grabKeyboard")
	symaddr(&addr_QWidget_releaseKeyboard, "QWidget_releaseKeyboard")
	symaddr(&addr_QWidget_grabShortcut, "QWidget_grabShortcut")
	symaddr(&addr_QWidget_releaseShortcut, "QWidget_releaseShortcut")
	symaddr(&addr_QWidget_setShortcutEnabled, "QWidget_setShortcutEnabled")
	symaddr(&addr_QWidget_setShortcutAutoRepeat, "QWidget_setShortcutAutoRepeat")
	symaddr(&addr_QWidget_mouseGrabber, "QWidget_mouseGrabber")
	symaddr(&addr_QWidget_keyboardGrabber, "QWidget_keyboardGrabber")
	symaddr(&addr_QWidget_updatesEnabled, "QWidget_updatesEnabled")
	symaddr(&addr_QWidget_setUpdatesEnabled, "QWidget_setUpdatesEnabled")
	symaddr(&addr_QWidget_graphicsProxyWidget, "QWidget_graphicsProxyWidget")
	symaddr(&addr_QWidget_update, "QWidget_update")
	symaddr(&addr_QWidget_repaint, "QWidget_repaint")
	symaddr(&addr_QWidget_update2, "QWidget_update2")
	symaddr(&addr_QWidget_update3, "QWidget_update3")
	symaddr(&addr_QWidget_update4, "QWidget_update4")
	symaddr(&addr_QWidget_repaint2, "QWidget_repaint2")
	symaddr(&addr_QWidget_repaint3, "QWidget_repaint3")
	symaddr(&addr_QWidget_repaint4, "QWidget_repaint4")
	symaddr(&addr_QWidget_setVisible, "QWidget_setVisible")
	symaddr(&addr_QWidget_setHidden, "QWidget_setHidden")
	symaddr(&addr_QWidget_show, "QWidget_show")
	symaddr(&addr_QWidget_hide, "QWidget_hide")
	symaddr(&addr_QWidget_showMinimized, "QWidget_showMinimized")
	symaddr(&addr_QWidget_showMaximized, "QWidget_showMaximized")
	symaddr(&addr_QWidget_showFullScreen, "QWidget_showFullScreen")
	symaddr(&addr_QWidget_showNormal, "QWidget_showNormal")
	symaddr(&addr_QWidget_close, "QWidget_close")
	symaddr(&addr_QWidget_raise, "QWidget_raise")
	symaddr(&addr_QWidget_lower, "QWidget_lower")
	symaddr(&addr_QWidget_stackUnder, "QWidget_stackUnder")
	symaddr(&addr_QWidget_move, "QWidget_move")
	symaddr(&addr_QWidget_move2, "QWidget_move2")
	symaddr(&addr_QWidget_resize, "QWidget_resize")
	symaddr(&addr_QWidget_resize2, "QWidget_resize2")
	symaddr(&addr_QWidget_setGeometry, "QWidget_setGeometry")
	symaddr(&addr_QWidget_setGeometry2, "QWidget_setGeometry2")
	symaddr(&addr_QWidget_saveGeometry, "QWidget_saveGeometry")
	symaddr(&addr_QWidget_restoreGeometry, "QWidget_restoreGeometry")
	symaddr(&addr_QWidget_adjustSize, "QWidget_adjustSize")
	symaddr(&addr_QWidget_isVisible, "QWidget_isVisible")
	symaddr(&addr_QWidget_isVisibleTo, "QWidget_isVisibleTo")
	symaddr(&addr_QWidget_isHidden, "QWidget_isHidden")
	symaddr(&addr_QWidget_isMinimized, "QWidget_isMinimized")
	symaddr(&addr_QWidget_isMaximized, "QWidget_isMaximized")
	symaddr(&addr_QWidget_isFullScreen, "QWidget_isFullScreen")
	symaddr(&addr_QWidget_windowState, "QWidget_windowState")
	symaddr(&addr_QWidget_setWindowState, "QWidget_setWindowState")
	symaddr(&addr_QWidget_overrideWindowState, "QWidget_overrideWindowState")
	symaddr(&addr_QWidget_sizeHint, "QWidget_sizeHint")
	symaddr(&addr_QWidget_minimumSizeHint, "QWidget_minimumSizeHint")
	symaddr(&addr_QWidget_sizePolicy, "QWidget_sizePolicy")
	symaddr(&addr_QWidget_setSizePolicy, "QWidget_setSizePolicy")
	symaddr(&addr_QWidget_setSizePolicy2, "QWidget_setSizePolicy2")
	symaddr(&addr_QWidget_heightForWidth, "QWidget_heightForWidth")
	symaddr(&addr_QWidget_hasHeightForWidth, "QWidget_hasHeightForWidth")
	symaddr(&addr_QWidget_visibleRegion, "QWidget_visibleRegion")
	symaddr(&addr_QWidget_setContentsMargins, "QWidget_setContentsMargins")
	symaddr(&addr_QWidget_setContentsMargins2, "QWidget_setContentsMargins2")
	symaddr(&addr_QWidget_contentsMargins, "QWidget_contentsMargins")
	symaddr(&addr_QWidget_contentsRect, "QWidget_contentsRect")
	symaddr(&addr_QWidget_layout, "QWidget_layout")
	symaddr(&addr_QWidget_setLayout, "QWidget_setLayout")
	symaddr(&addr_QWidget_updateGeometry, "QWidget_updateGeometry")
	symaddr(&addr_QWidget_setParent, "QWidget_setParent")
	symaddr(&addr_QWidget_setParent2, "QWidget_setParent2")
	symaddr(&addr_QWidget_scroll, "QWidget_scroll")
	symaddr(&addr_QWidget_scroll2, "QWidget_scroll2")
	symaddr(&addr_QWidget_focusWidget, "QWidget_focusWidget")
	symaddr(&addr_QWidget_nextInFocusChain, "QWidget_nextInFocusChain")
	symaddr(&addr_QWidget_previousInFocusChain, "QWidget_previousInFocusChain")
	symaddr(&addr_QWidget_acceptDrops, "QWidget_acceptDrops")
	symaddr(&addr_QWidget_setAcceptDrops, "QWidget_setAcceptDrops")
	symaddr(&addr_QWidget_addAction, "QWidget_addAction")
	symaddr(&addr_QWidget_addActions, "QWidget_addActions")
	symaddr(&addr_QWidget_insertActions, "QWidget_insertActions")
	symaddr(&addr_QWidget_insertAction, "QWidget_insertAction")
	symaddr(&addr_QWidget_removeAction, "QWidget_removeAction")
	symaddr(&addr_QWidget_actions, "QWidget_actions")
	symaddr(&addr_QWidget_addAction2, "QWidget_addAction2")
	symaddr(&addr_QWidget_addAction3, "QWidget_addAction3")
	symaddr(&addr_QWidget_addAction4, "QWidget_addAction4")
	symaddr(&addr_QWidget_addAction5, "QWidget_addAction5")
	symaddr(&addr_QWidget_addAction6, "QWidget_addAction6")
	symaddr(&addr_QWidget_addAction7, "QWidget_addAction7")
	symaddr(&addr_QWidget_addAction8, "QWidget_addAction8")
	symaddr(&addr_QWidget_addAction9, "QWidget_addAction9")
	symaddr(&addr_QWidget_parentWidget, "QWidget_parentWidget")
	symaddr(&addr_QWidget_setWindowFlags, "QWidget_setWindowFlags")
	symaddr(&addr_QWidget_windowFlags, "QWidget_windowFlags")
	symaddr(&addr_QWidget_setWindowFlag, "QWidget_setWindowFlag")
	symaddr(&addr_QWidget_overrideWindowFlags, "QWidget_overrideWindowFlags")
	symaddr(&addr_QWidget_windowType, "QWidget_windowType")
	symaddr(&addr_QWidget_find, "QWidget_find")
	symaddr(&addr_QWidget_childAt, "QWidget_childAt")
	symaddr(&addr_QWidget_childAt2, "QWidget_childAt2")
	symaddr(&addr_QWidget_childAt3, "QWidget_childAt3")
	symaddr(&addr_QWidget_setAttribute, "QWidget_setAttribute")
	symaddr(&addr_QWidget_testAttribute, "QWidget_testAttribute")
	symaddr(&addr_QWidget_paintEngine, "QWidget_paintEngine")
	symaddr(&addr_QWidget_ensurePolished, "QWidget_ensurePolished")
	symaddr(&addr_QWidget_isAncestorOf, "QWidget_isAncestorOf")
	symaddr(&addr_QWidget_autoFillBackground, "QWidget_autoFillBackground")
	symaddr(&addr_QWidget_setAutoFillBackground, "QWidget_setAutoFillBackground")
	symaddr(&addr_QWidget_backingStore, "QWidget_backingStore")
	symaddr(&addr_QWidget_windowHandle, "QWidget_windowHandle")
	symaddr(&addr_QWidget_screen, "QWidget_screen")
	symaddr(&addr_QWidget_setScreen, "QWidget_setScreen")
	symaddr(&addr_QWidget_createWindowContainer, "QWidget_createWindowContainer")
	symaddr(&addr_QWidget_windowTitleChanged, "QWidget_windowTitleChanged")
	symaddr(&addr_QWidget_windowIconChanged, "QWidget_windowIconChanged")
	symaddr(&addr_QWidget_windowIconTextChanged, "QWidget_windowIconTextChanged")
	symaddr(&addr_QWidget_customContextMenuRequested, "QWidget_customContextMenuRequested")
	symaddr(&addr_QWidget_inputMethodQuery, "QWidget_inputMethodQuery")
	symaddr(&addr_QWidget_inputMethodHints, "QWidget_inputMethodHints")
	symaddr(&addr_QWidget_setInputMethodHints, "QWidget_setInputMethodHints")
	symaddr(&addr_QWidget_updateMicroFocus, "QWidget_updateMicroFocus")
	symaddr(&addr_QWidget_create, "QWidget_create")
	symaddr(&addr_QWidget_destroy, "QWidget_destroy")
	symaddr(&addr_QWidget_focusNextChild, "QWidget_focusNextChild")
	symaddr(&addr_QWidget_focusPreviousChild, "QWidget_focusPreviousChild")
	symaddr(&addr_QWidget_connect, "QWidget_connect")
	symaddr(&addr_QWidget_connect2, "QWidget_connect2")
	symaddr(&addr_QWidget_sender, "QWidget_sender")
	symaddr(&addr_QWidget_senderSignalIndex, "QWidget_senderSignalIndex")
	symaddr(&addr_QWidget_receivers, "QWidget_receivers")
	symaddr(&addr_QWidget_isSignalConnected, "QWidget_isSignalConnected")
	symaddr(&addr_QWidget_paintingActive, "QWidget_paintingActive")
	symaddr(&addr_QWidget_widthMM, "QWidget_widthMM")
	symaddr(&addr_QWidget_heightMM, "QWidget_heightMM")
	symaddr(&addr_QWidget_logicalDpiX, "QWidget_logicalDpiX")
	symaddr(&addr_QWidget_logicalDpiY, "QWidget_logicalDpiY")
	symaddr(&addr_QWidget_physicalDpiX, "QWidget_physicalDpiX")
	symaddr(&addr_QWidget_physicalDpiY, "QWidget_physicalDpiY")
	reg(&pg_QWidget_devicePixelRatio, "QWidget_devicePixelRatio")
	reg(&pg_QWidget_devicePixelRatioF, "QWidget_devicePixelRatioF")
	symaddr(&addr_QWidget_colorCount, "QWidget_colorCount")
	symaddr(&addr_QWidget_depth, "QWidget_depth")
	reg(&pg_QWidget_devicePixelRatioFScale, "QWidget_devicePixelRatioFScale")
	reg(&pg_QWidget_encodeMetricF, "QWidget_encodeMetricF")
	reg(&pg_QWidget_getDecodedMetricF, "QWidget_getDecodedMetricF")
	symaddr(&addr_QWidget_tr2, "QWidget_tr2")
	symaddr(&addr_QWidget_tr3, "QWidget_tr3")
	symaddr(&addr_QWidget_render3, "QWidget_render3")
	symaddr(&addr_QWidget_render4, "QWidget_render4")
	symaddr(&addr_QWidget_render5, "QWidget_render5")
	symaddr(&addr_QWidget_render6, "QWidget_render6")
	symaddr(&addr_QWidget_render7, "QWidget_render7")
	symaddr(&addr_QWidget_render8, "QWidget_render8")
	symaddr(&addr_QWidget_grab2, "QWidget_grab2")
	symaddr(&addr_QWidget_grabGesture2, "QWidget_grabGesture2")
	symaddr(&addr_QWidget_grabShortcut2, "QWidget_grabShortcut2")
	symaddr(&addr_QWidget_setShortcutEnabled2, "QWidget_setShortcutEnabled2")
	symaddr(&addr_QWidget_setShortcutAutoRepeat2, "QWidget_setShortcutAutoRepeat2")
	symaddr(&addr_QWidget_addAction10, "QWidget_addAction10")
	symaddr(&addr_QWidget_addAction11, "QWidget_addAction11")
	symaddr(&addr_QWidget_addAction12, "QWidget_addAction12")
	symaddr(&addr_QWidget_addAction13, "QWidget_addAction13")
	symaddr(&addr_QWidget_setWindowFlag2, "QWidget_setWindowFlag2")
	symaddr(&addr_QWidget_setAttribute2, "QWidget_setAttribute2")
	symaddr(&addr_QWidget_createWindowContainer2, "QWidget_createWindowContainer2")
	symaddr(&addr_QWidget_createWindowContainer3, "QWidget_createWindowContainer3")
	symaddr(&addr_QWidget_updateMicroFocus2, "QWidget_updateMicroFocus2")
	symaddr(&addr_QWidget_create2, "QWidget_create2")
	symaddr(&addr_QWidget_create3, "QWidget_create3")
	symaddr(&addr_QWidget_create4, "QWidget_create4")
	symaddr(&addr_QWidget_destroy2, "QWidget_destroy2")
	symaddr(&addr_QWidget_destroy3, "QWidget_destroy3")
	symaddr(&addr_QWidget_connect3, "QWidget_connect3")
	symaddr(&addr_QWidget_connect4, "QWidget_connect4")
	symaddr(&addr_QWidget_asQPaintDevice, "QWidget_asQPaintDevice")
	symaddr(&addr_QWidget_devType_default, "QWidget_devType_default")
	symaddr(&addr_QWidget_setVisible_default, "QWidget_setVisible_default")
	symaddr(&addr_QWidget_sizeHint_default, "QWidget_sizeHint_default")
	symaddr(&addr_QWidget_minimumSizeHint_default, "QWidget_minimumSizeHint_default")
	symaddr(&addr_QWidget_heightForWidth_default, "QWidget_heightForWidth_default")
	symaddr(&addr_QWidget_hasHeightForWidth_default, "QWidget_hasHeightForWidth_default")
	symaddr(&addr_QWidget_paintEngine_default, "QWidget_paintEngine_default")
	symaddr(&addr_QWidget_inputMethodQuery_default, "QWidget_inputMethodQuery_default")
	symaddr(&addr_QWidget_metaObject_default, "QWidget_metaObject_default")
	symaddr(&addr_QWidget_event_default, "QWidget_event_default")
	symaddr(&addr_QWidget_mousePressEvent_default, "QWidget_mousePressEvent_default")
	symaddr(&addr_QWidget_mouseReleaseEvent_default, "QWidget_mouseReleaseEvent_default")
	symaddr(&addr_QWidget_mouseDoubleClickEvent_default, "QWidget_mouseDoubleClickEvent_default")
	symaddr(&addr_QWidget_mouseMoveEvent_default, "QWidget_mouseMoveEvent_default")
	symaddr(&addr_QWidget_wheelEvent_default, "QWidget_wheelEvent_default")
	symaddr(&addr_QWidget_keyPressEvent_default, "QWidget_keyPressEvent_default")
	symaddr(&addr_QWidget_keyReleaseEvent_default, "QWidget_keyReleaseEvent_default")
	symaddr(&addr_QWidget_focusInEvent_default, "QWidget_focusInEvent_default")
	symaddr(&addr_QWidget_focusOutEvent_default, "QWidget_focusOutEvent_default")
	symaddr(&addr_QWidget_enterEvent_default, "QWidget_enterEvent_default")
	symaddr(&addr_QWidget_leaveEvent_default, "QWidget_leaveEvent_default")
	symaddr(&addr_QWidget_paintEvent_default, "QWidget_paintEvent_default")
	symaddr(&addr_QWidget_moveEvent_default, "QWidget_moveEvent_default")
	symaddr(&addr_QWidget_resizeEvent_default, "QWidget_resizeEvent_default")
	symaddr(&addr_QWidget_closeEvent_default, "QWidget_closeEvent_default")
	symaddr(&addr_QWidget_contextMenuEvent_default, "QWidget_contextMenuEvent_default")
	symaddr(&addr_QWidget_tabletEvent_default, "QWidget_tabletEvent_default")
	symaddr(&addr_QWidget_actionEvent_default, "QWidget_actionEvent_default")
	symaddr(&addr_QWidget_dragEnterEvent_default, "QWidget_dragEnterEvent_default")
	symaddr(&addr_QWidget_dragMoveEvent_default, "QWidget_dragMoveEvent_default")
	symaddr(&addr_QWidget_dragLeaveEvent_default, "QWidget_dragLeaveEvent_default")
	symaddr(&addr_QWidget_dropEvent_default, "QWidget_dropEvent_default")
	symaddr(&addr_QWidget_showEvent_default, "QWidget_showEvent_default")
	symaddr(&addr_QWidget_hideEvent_default, "QWidget_hideEvent_default")
	symaddr(&addr_QWidget_changeEvent_default, "QWidget_changeEvent_default")
	symaddr(&addr_QWidget_metric_default, "QWidget_metric_default")
	symaddr(&addr_QWidget_initPainter_default, "QWidget_initPainter_default")
	symaddr(&addr_QWidget_redirected_default, "QWidget_redirected_default")
	symaddr(&addr_QWidget_sharedPainter_default, "QWidget_sharedPainter_default")
	symaddr(&addr_QWidget_inputMethodEvent_default, "QWidget_inputMethodEvent_default")
	symaddr(&addr_QWidget_focusNextPrevChild_default, "QWidget_focusNextPrevChild_default")
	symaddr(&addr_QWidget_eventFilter_default, "QWidget_eventFilter_default")
	symaddr(&addr_QWidget_timerEvent_default, "QWidget_timerEvent_default")
	symaddr(&addr_QWidget_childEvent_default, "QWidget_childEvent_default")
	symaddr(&addr_QWidget_customEvent_default, "QWidget_customEvent_default")
	symaddr(&addr_QWidget_connectNotify_default, "QWidget_connectNotify_default")
	symaddr(&addr_QWidget_disconnectNotify_default, "QWidget_disconnectNotify_default")
	reg(&pg_QWidget_connect_windowTitleChanged, "QWidget_connect_windowTitleChanged")
	reg(&pg_QWidget_connect_windowIconChanged, "QWidget_connect_windowIconChanged")
	reg(&pg_QWidget_connect_windowIconTextChanged, "QWidget_connect_windowIconTextChanged")
	reg(&pg_QWidget_connect_customContextMenuRequested, "QWidget_connect_customContextMenuRequested")
	reg(&pg_QWidget_connect_destroyed, "QWidget_connect_destroyed")
	reg(&pg_QWidget_connect_destroyed2, "QWidget_connect_destroyed2")
	reg(&pg_QWidget_connect_objectNameChanged, "QWidget_connect_objectNameChanged")
	reg(&pg_QWidget__disconnect, "QWidget__disconnect")
	reg(&pg_QWidget__set_override, "QWidget__set_override")
	reg(&pg_QWidget_delete, "QWidget_delete")
}

type QWidget struct {
	h uintptr
	*QObject
}

func bindQWidget(h uintptr) *QWidget {
	if h == 0 {
		return nil
	}
	return &QWidget{h: h, QObject: bindQObject(pg_QWidget_base(h))}
}

func newownedQWidget(h uintptr) *QWidget { return bindQWidget(h) }

func (this *QWidget) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QWidget) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQWidget(h unsafe.Pointer) *QWidget { return bindQWidget(uintptr(h)) }
func UnsafeNewQWidgetOwned(h unsafe.Pointer) *QWidget { return newownedQWidget(uintptr(h)) }

func (this *QWidget) UnsafeRelease() {
	this.h = 0
	if this.QObject != nil {
		this.QObject.UnsafeRelease()
	}
}

func NewQWidget(parent *QWidget) *QWidget {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QWidget_new, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQWidget(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QWidget_virtIdx, QWidget_enable)
	return __o
}

func NewQWidget2() *QWidget {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QWidget_new2, uintptr(__goid))
	__o := newownedQWidget(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QWidget_virtIdx, QWidget_enable)
	return __o
}

func NewQWidget3(parent *QWidget, f WindowType) *QWidget {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QWidget_new3, uintptr(__goid), uintptr(parent.UnsafePointer()), uintptr(f))
	__o := newownedQWidget(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QWidget_virtIdx, QWidget_enable)
	return __o
}

func QWidget_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QWidget_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QWidget) DevType() int {
	__r0, _, _ := purego.SyscallN(addr_QWidget_devType, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QWidget) WinId() uint64 {
	__r0, _, _ := purego.SyscallN(addr_QWidget_winId, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return uint64(__r)
}

func (this *QWidget) CreateWinId() {
	purego.SyscallN(addr_QWidget_createWinId, this.h)
	runtime.KeepAlive(this)
}

func (this *QWidget) InternalWinId() uint64 {
	__r0, _, _ := purego.SyscallN(addr_QWidget_internalWinId, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return uint64(__r)
}

func (this *QWidget) EffectiveWinId() uint64 {
	__r0, _, _ := purego.SyscallN(addr_QWidget_effectiveWinId, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return uint64(__r)
}

func (this *QWidget) Style() *QStyle {
	__r, _, _ := purego.SyscallN(addr_QWidget_style, this.h)
	runtime.KeepAlive(this)
	return bindQStyle(__r)
}

func (this *QWidget) SetStyle(arg1 *QStyle) {
	purego.SyscallN(addr_QWidget_setStyle, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QWidget) IsTopLevel() bool {
	__r0, _, _ := purego.SyscallN(addr_QWidget_isTopLevel, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QWidget) IsWindow() bool {
	__r0, _, _ := purego.SyscallN(addr_QWidget_isWindow, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QWidget) IsModal() bool {
	__r0, _, _ := purego.SyscallN(addr_QWidget_isModal, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QWidget) WindowModality() WindowModality {
	__r0, _, _ := purego.SyscallN(addr_QWidget_windowModality, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return WindowModality(__r)
}

func (this *QWidget) SetWindowModality(windowModality WindowModality) {
	purego.SyscallN(addr_QWidget_setWindowModality, this.h, uintptr(windowModality))
	runtime.KeepAlive(this)
}

func (this *QWidget) IsEnabled() bool {
	__r0, _, _ := purego.SyscallN(addr_QWidget_isEnabled, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QWidget) IsEnabledTo(arg1 *QWidget) bool {
	__r0, _, _ := purego.SyscallN(addr_QWidget_isEnabledTo, this.h, uintptr(arg1.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return __r != 0
}

func (this *QWidget) SetEnabled(arg1 bool) {
	purego.SyscallN(addr_QWidget_setEnabled, this.h, uintptr(b2i(arg1)))
	runtime.KeepAlive(this)
}

func (this *QWidget) SetDisabled(arg1 bool) {
	purego.SyscallN(addr_QWidget_setDisabled, this.h, uintptr(b2i(arg1)))
	runtime.KeepAlive(this)
}

func (this *QWidget) SetWindowModified(arg1 bool) {
	purego.SyscallN(addr_QWidget_setWindowModified, this.h, uintptr(b2i(arg1)))
	runtime.KeepAlive(this)
}

func (this *QWidget) FrameGeometry() *QRect {
	__r, _, _ := purego.SyscallN(addr_QWidget_frameGeometry, this.h)
	runtime.KeepAlive(this)
	return newownedQRect(__r)
}

func (this *QWidget) Geometry() *QRect {
	__r, _, _ := purego.SyscallN(addr_QWidget_geometry, this.h)
	runtime.KeepAlive(this)
	return bindQRect(__r)
}

func (this *QWidget) NormalGeometry() *QRect {
	__r, _, _ := purego.SyscallN(addr_QWidget_normalGeometry, this.h)
	runtime.KeepAlive(this)
	return newownedQRect(__r)
}

func (this *QWidget) X() int {
	__r0, _, _ := purego.SyscallN(addr_QWidget_x, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QWidget) Y() int {
	__r0, _, _ := purego.SyscallN(addr_QWidget_y, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QWidget) Pos() *QPoint {
	__r, _, _ := purego.SyscallN(addr_QWidget_pos, this.h)
	runtime.KeepAlive(this)
	return newownedQPoint(__r)
}

func (this *QWidget) FrameSize() *QSize {
	__r, _, _ := purego.SyscallN(addr_QWidget_frameSize, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QWidget) Size() *QSize {
	__r, _, _ := purego.SyscallN(addr_QWidget_size, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QWidget) Width() int {
	__r0, _, _ := purego.SyscallN(addr_QWidget_width, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QWidget) Height() int {
	__r0, _, _ := purego.SyscallN(addr_QWidget_height, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QWidget) Rect() *QRect {
	__r, _, _ := purego.SyscallN(addr_QWidget_rect, this.h)
	runtime.KeepAlive(this)
	return newownedQRect(__r)
}

func (this *QWidget) ChildrenRect() *QRect {
	__r, _, _ := purego.SyscallN(addr_QWidget_childrenRect, this.h)
	runtime.KeepAlive(this)
	return newownedQRect(__r)
}

func (this *QWidget) ChildrenRegion() *QRegion {
	__r, _, _ := purego.SyscallN(addr_QWidget_childrenRegion, this.h)
	runtime.KeepAlive(this)
	return newownedQRegion(__r)
}

func (this *QWidget) MinimumSize() *QSize {
	__r, _, _ := purego.SyscallN(addr_QWidget_minimumSize, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QWidget) MaximumSize() *QSize {
	__r, _, _ := purego.SyscallN(addr_QWidget_maximumSize, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QWidget) MinimumWidth() int {
	__r0, _, _ := purego.SyscallN(addr_QWidget_minimumWidth, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QWidget) MinimumHeight() int {
	__r0, _, _ := purego.SyscallN(addr_QWidget_minimumHeight, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QWidget) MaximumWidth() int {
	__r0, _, _ := purego.SyscallN(addr_QWidget_maximumWidth, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QWidget) MaximumHeight() int {
	__r0, _, _ := purego.SyscallN(addr_QWidget_maximumHeight, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QWidget) SetMinimumSize(arg1 *QSize) {
	purego.SyscallN(addr_QWidget_setMinimumSize, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QWidget) SetMinimumSize2(minw int, minh int) {
	purego.SyscallN(addr_QWidget_setMinimumSize2, this.h, uintptr(minw), uintptr(minh))
	runtime.KeepAlive(this)
}

func (this *QWidget) SetMaximumSize(arg1 *QSize) {
	purego.SyscallN(addr_QWidget_setMaximumSize, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QWidget) SetMaximumSize2(maxw int, maxh int) {
	purego.SyscallN(addr_QWidget_setMaximumSize2, this.h, uintptr(maxw), uintptr(maxh))
	runtime.KeepAlive(this)
}

func (this *QWidget) SetMinimumWidth(minw int) {
	purego.SyscallN(addr_QWidget_setMinimumWidth, this.h, uintptr(minw))
	runtime.KeepAlive(this)
}

func (this *QWidget) SetMinimumHeight(minh int) {
	purego.SyscallN(addr_QWidget_setMinimumHeight, this.h, uintptr(minh))
	runtime.KeepAlive(this)
}

func (this *QWidget) SetMaximumWidth(maxw int) {
	purego.SyscallN(addr_QWidget_setMaximumWidth, this.h, uintptr(maxw))
	runtime.KeepAlive(this)
}

func (this *QWidget) SetMaximumHeight(maxh int) {
	purego.SyscallN(addr_QWidget_setMaximumHeight, this.h, uintptr(maxh))
	runtime.KeepAlive(this)
}

func (this *QWidget) SizeIncrement() *QSize {
	__r, _, _ := purego.SyscallN(addr_QWidget_sizeIncrement, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QWidget) SetSizeIncrement(arg1 *QSize) {
	purego.SyscallN(addr_QWidget_setSizeIncrement, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QWidget) SetSizeIncrement2(w int, h int) {
	purego.SyscallN(addr_QWidget_setSizeIncrement2, this.h, uintptr(w), uintptr(h))
	runtime.KeepAlive(this)
}

func (this *QWidget) BaseSize() *QSize {
	__r, _, _ := purego.SyscallN(addr_QWidget_baseSize, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QWidget) SetBaseSize(arg1 *QSize) {
	purego.SyscallN(addr_QWidget_setBaseSize, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QWidget) SetBaseSize2(basew int, baseh int) {
	purego.SyscallN(addr_QWidget_setBaseSize2, this.h, uintptr(basew), uintptr(baseh))
	runtime.KeepAlive(this)
}

func (this *QWidget) SetFixedSize(arg1 *QSize) {
	purego.SyscallN(addr_QWidget_setFixedSize, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QWidget) SetFixedSize2(w int, h int) {
	purego.SyscallN(addr_QWidget_setFixedSize2, this.h, uintptr(w), uintptr(h))
	runtime.KeepAlive(this)
}

func (this *QWidget) SetFixedWidth(w int) {
	purego.SyscallN(addr_QWidget_setFixedWidth, this.h, uintptr(w))
	runtime.KeepAlive(this)
}

func (this *QWidget) SetFixedHeight(h int) {
	purego.SyscallN(addr_QWidget_setFixedHeight, this.h, uintptr(h))
	runtime.KeepAlive(this)
}

func (this *QWidget) MapToGlobal(arg1 *QPointF) *QPointF {
	__r, _, _ := purego.SyscallN(addr_QWidget_mapToGlobal, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return newownedQPointF(__r)
}

func (this *QWidget) MapToGlobalWithQPoint(arg1 *QPoint) *QPoint {
	__r, _, _ := purego.SyscallN(addr_QWidget_mapToGlobal2, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return newownedQPoint(__r)
}

func (this *QWidget) MapFromGlobal(arg1 *QPointF) *QPointF {
	__r, _, _ := purego.SyscallN(addr_QWidget_mapFromGlobal, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return newownedQPointF(__r)
}

func (this *QWidget) MapFromGlobalWithQPoint(arg1 *QPoint) *QPoint {
	__r, _, _ := purego.SyscallN(addr_QWidget_mapFromGlobal2, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return newownedQPoint(__r)
}

func (this *QWidget) MapToParent(arg1 *QPointF) *QPointF {
	__r, _, _ := purego.SyscallN(addr_QWidget_mapToParent, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return newownedQPointF(__r)
}

func (this *QWidget) MapToParentWithQPoint(arg1 *QPoint) *QPoint {
	__r, _, _ := purego.SyscallN(addr_QWidget_mapToParent2, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return newownedQPoint(__r)
}

func (this *QWidget) MapFromParent(arg1 *QPointF) *QPointF {
	__r, _, _ := purego.SyscallN(addr_QWidget_mapFromParent, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return newownedQPointF(__r)
}

func (this *QWidget) MapFromParentWithQPoint(arg1 *QPoint) *QPoint {
	__r, _, _ := purego.SyscallN(addr_QWidget_mapFromParent2, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return newownedQPoint(__r)
}

func (this *QWidget) MapTo(arg1 *QWidget, arg2 *QPointF) *QPointF {
	__r, _, _ := purego.SyscallN(addr_QWidget_mapTo, this.h, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
	return newownedQPointF(__r)
}

func (this *QWidget) MapTo2(arg1 *QWidget, arg2 *QPoint) *QPoint {
	__r, _, _ := purego.SyscallN(addr_QWidget_mapTo2, this.h, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
	return newownedQPoint(__r)
}

func (this *QWidget) MapFrom(arg1 *QWidget, arg2 *QPointF) *QPointF {
	__r, _, _ := purego.SyscallN(addr_QWidget_mapFrom, this.h, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
	return newownedQPointF(__r)
}

func (this *QWidget) MapFrom2(arg1 *QWidget, arg2 *QPoint) *QPoint {
	__r, _, _ := purego.SyscallN(addr_QWidget_mapFrom2, this.h, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
	return newownedQPoint(__r)
}

func (this *QWidget) Window() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QWidget_window, this.h)
	runtime.KeepAlive(this)
	return bindQWidget(__r)
}

func (this *QWidget) NativeParentWidget() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QWidget_nativeParentWidget, this.h)
	runtime.KeepAlive(this)
	return bindQWidget(__r)
}

func (this *QWidget) TopLevelWidget() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QWidget_topLevelWidget, this.h)
	runtime.KeepAlive(this)
	return bindQWidget(__r)
}

func (this *QWidget) Palette() *QPalette {
	__r, _, _ := purego.SyscallN(addr_QWidget_palette, this.h)
	runtime.KeepAlive(this)
	return bindQPalette(__r)
}

func (this *QWidget) SetPalette(arg1 *QPalette) {
	purego.SyscallN(addr_QWidget_setPalette, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QWidget) SetBackgroundRole(arg1 QPalette__ColorRole) {
	purego.SyscallN(addr_QWidget_setBackgroundRole, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QWidget) BackgroundRole() QPalette__ColorRole {
	__r0, _, _ := purego.SyscallN(addr_QWidget_backgroundRole, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QPalette__ColorRole(__r)
}

func (this *QWidget) SetForegroundRole(arg1 QPalette__ColorRole) {
	purego.SyscallN(addr_QWidget_setForegroundRole, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QWidget) ForegroundRole() QPalette__ColorRole {
	__r0, _, _ := purego.SyscallN(addr_QWidget_foregroundRole, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QPalette__ColorRole(__r)
}

func (this *QWidget) Font() *QFont {
	__r, _, _ := purego.SyscallN(addr_QWidget_font, this.h)
	runtime.KeepAlive(this)
	return bindQFont(__r)
}

func (this *QWidget) SetFont(arg1 *QFont) {
	purego.SyscallN(addr_QWidget_setFont, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QWidget) FontMetrics() *QFontMetrics {
	__r, _, _ := purego.SyscallN(addr_QWidget_fontMetrics, this.h)
	runtime.KeepAlive(this)
	return newownedQFontMetrics(__r)
}

func (this *QWidget) FontInfo() *QFontInfo {
	__r, _, _ := purego.SyscallN(addr_QWidget_fontInfo, this.h)
	runtime.KeepAlive(this)
	return newownedQFontInfo(__r)
}

func (this *QWidget) Cursor() *QCursor {
	__r, _, _ := purego.SyscallN(addr_QWidget_cursor, this.h)
	runtime.KeepAlive(this)
	return newownedQCursor(__r)
}

func (this *QWidget) SetCursor(arg1 *QCursor) {
	purego.SyscallN(addr_QWidget_setCursor, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QWidget) UnsetCursor() {
	purego.SyscallN(addr_QWidget_unsetCursor, this.h)
	runtime.KeepAlive(this)
}

func (this *QWidget) SetMouseTracking(enable bool) {
	purego.SyscallN(addr_QWidget_setMouseTracking, this.h, uintptr(b2i(enable)))
	runtime.KeepAlive(this)
}

func (this *QWidget) HasMouseTracking() bool {
	__r0, _, _ := purego.SyscallN(addr_QWidget_hasMouseTracking, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QWidget) UnderMouse() bool {
	__r0, _, _ := purego.SyscallN(addr_QWidget_underMouse, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QWidget) SetTabletTracking(enable bool) {
	purego.SyscallN(addr_QWidget_setTabletTracking, this.h, uintptr(b2i(enable)))
	runtime.KeepAlive(this)
}

func (this *QWidget) HasTabletTracking() bool {
	__r0, _, _ := purego.SyscallN(addr_QWidget_hasTabletTracking, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QWidget) SetMask(arg1 *QBitmap) {
	purego.SyscallN(addr_QWidget_setMask, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QWidget) SetMaskWithQRegion(arg1 *QRegion) {
	purego.SyscallN(addr_QWidget_setMask2, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QWidget) Mask() *QRegion {
	__r, _, _ := purego.SyscallN(addr_QWidget_mask, this.h)
	runtime.KeepAlive(this)
	return newownedQRegion(__r)
}

func (this *QWidget) ClearMask() {
	purego.SyscallN(addr_QWidget_clearMask, this.h)
	runtime.KeepAlive(this)
}

func (this *QWidget) Render(target *QPaintDevice) {
	purego.SyscallN(addr_QWidget_render, this.h, uintptr(target.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(target)
}

func (this *QWidget) RenderWithPainter(painter *QPainter) {
	purego.SyscallN(addr_QWidget_render2, this.h, uintptr(painter.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
}

func (this *QWidget) Grab() *QPixmap {
	__r, _, _ := purego.SyscallN(addr_QWidget_grab, this.h)
	runtime.KeepAlive(this)
	return newownedQPixmap(__r)
}

func (this *QWidget) GraphicsEffect() *QGraphicsEffect {
	__r, _, _ := purego.SyscallN(addr_QWidget_graphicsEffect, this.h)
	runtime.KeepAlive(this)
	return bindQGraphicsEffect(__r)
}

func (this *QWidget) SetGraphicsEffect(effect *QGraphicsEffect) {
	purego.SyscallN(addr_QWidget_setGraphicsEffect, this.h, uintptr(effect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(effect)
}

func (this *QWidget) GrabGesture(type_ GestureType) {
	purego.SyscallN(addr_QWidget_grabGesture, this.h, uintptr(type_))
	runtime.KeepAlive(this)
}

func (this *QWidget) UngrabGesture(type_ GestureType) {
	purego.SyscallN(addr_QWidget_ungrabGesture, this.h, uintptr(type_))
	runtime.KeepAlive(this)
}

func (this *QWidget) SetWindowTitle(arg1 string) {
	var arg1_p uintptr
	if len(arg1) > 0 {
		arg1_p = uintptr(unsafe.Pointer(unsafe.StringData(arg1)))
	}
	purego.SyscallN(addr_QWidget_setWindowTitle, this.h, arg1_p, uintptr(len(arg1)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QWidget) SetStyleSheet(styleSheet string) {
	var styleSheet_p uintptr
	if len(styleSheet) > 0 {
		styleSheet_p = uintptr(unsafe.Pointer(unsafe.StringData(styleSheet)))
	}
	purego.SyscallN(addr_QWidget_setStyleSheet, this.h, styleSheet_p, uintptr(len(styleSheet)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(styleSheet)
}

func (this *QWidget) StyleSheet() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QWidget_styleSheet, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QWidget) WindowTitle() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QWidget_windowTitle, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QWidget) SetWindowIcon(icon *QIcon) {
	purego.SyscallN(addr_QWidget_setWindowIcon, this.h, uintptr(icon.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(icon)
}

func (this *QWidget) WindowIcon() *QIcon {
	__r, _, _ := purego.SyscallN(addr_QWidget_windowIcon, this.h)
	runtime.KeepAlive(this)
	return newownedQIcon(__r)
}

func (this *QWidget) SetWindowIconText(arg1 string) {
	var arg1_p uintptr
	if len(arg1) > 0 {
		arg1_p = uintptr(unsafe.Pointer(unsafe.StringData(arg1)))
	}
	purego.SyscallN(addr_QWidget_setWindowIconText, this.h, arg1_p, uintptr(len(arg1)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QWidget) WindowIconText() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QWidget_windowIconText, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QWidget) SetWindowRole(arg1 string) {
	var arg1_p uintptr
	if len(arg1) > 0 {
		arg1_p = uintptr(unsafe.Pointer(unsafe.StringData(arg1)))
	}
	purego.SyscallN(addr_QWidget_setWindowRole, this.h, arg1_p, uintptr(len(arg1)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QWidget) WindowRole() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QWidget_windowRole, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QWidget) SetWindowFilePath(filePath string) {
	var filePath_p uintptr
	if len(filePath) > 0 {
		filePath_p = uintptr(unsafe.Pointer(unsafe.StringData(filePath)))
	}
	purego.SyscallN(addr_QWidget_setWindowFilePath, this.h, filePath_p, uintptr(len(filePath)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(filePath)
}

func (this *QWidget) WindowFilePath() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QWidget_windowFilePath, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QWidget) SetWindowOpacity(level float64) {
	pg_QWidget_setWindowOpacity(this.h, level)
	runtime.KeepAlive(this)
}

func (this *QWidget) WindowOpacity() float64 {
	__r := pg_QWidget_windowOpacity(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QWidget) IsWindowModified() bool {
	__r0, _, _ := purego.SyscallN(addr_QWidget_isWindowModified, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QWidget) SetToolTip(arg1 string) {
	var arg1_p uintptr
	if len(arg1) > 0 {
		arg1_p = uintptr(unsafe.Pointer(unsafe.StringData(arg1)))
	}
	purego.SyscallN(addr_QWidget_setToolTip, this.h, arg1_p, uintptr(len(arg1)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QWidget) ToolTip() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QWidget_toolTip, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QWidget) SetToolTipDuration(msec int) {
	purego.SyscallN(addr_QWidget_setToolTipDuration, this.h, uintptr(msec))
	runtime.KeepAlive(this)
}

func (this *QWidget) ToolTipDuration() int {
	__r0, _, _ := purego.SyscallN(addr_QWidget_toolTipDuration, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QWidget) SetStatusTip(arg1 string) {
	var arg1_p uintptr
	if len(arg1) > 0 {
		arg1_p = uintptr(unsafe.Pointer(unsafe.StringData(arg1)))
	}
	purego.SyscallN(addr_QWidget_setStatusTip, this.h, arg1_p, uintptr(len(arg1)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QWidget) StatusTip() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QWidget_statusTip, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QWidget) SetWhatsThis(arg1 string) {
	var arg1_p uintptr
	if len(arg1) > 0 {
		arg1_p = uintptr(unsafe.Pointer(unsafe.StringData(arg1)))
	}
	purego.SyscallN(addr_QWidget_setWhatsThis, this.h, arg1_p, uintptr(len(arg1)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QWidget) WhatsThis() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QWidget_whatsThis, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QWidget) AccessibleName() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QWidget_accessibleName, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QWidget) SetAccessibleName(name string) {
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	purego.SyscallN(addr_QWidget_setAccessibleName, this.h, name_p, uintptr(len(name)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(name)
}

func (this *QWidget) AccessibleDescription() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QWidget_accessibleDescription, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QWidget) SetAccessibleDescription(description string) {
	var description_p uintptr
	if len(description) > 0 {
		description_p = uintptr(unsafe.Pointer(unsafe.StringData(description)))
	}
	purego.SyscallN(addr_QWidget_setAccessibleDescription, this.h, description_p, uintptr(len(description)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(description)
}

func (this *QWidget) AccessibleIdentifier() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QWidget_accessibleIdentifier, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QWidget) SetAccessibleIdentifier(identifier string) {
	var identifier_p uintptr
	if len(identifier) > 0 {
		identifier_p = uintptr(unsafe.Pointer(unsafe.StringData(identifier)))
	}
	purego.SyscallN(addr_QWidget_setAccessibleIdentifier, this.h, identifier_p, uintptr(len(identifier)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(identifier)
}

func (this *QWidget) SetLayoutDirection(direction LayoutDirection) {
	purego.SyscallN(addr_QWidget_setLayoutDirection, this.h, uintptr(direction))
	runtime.KeepAlive(this)
}

func (this *QWidget) LayoutDirection() LayoutDirection {
	__r0, _, _ := purego.SyscallN(addr_QWidget_layoutDirection, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return LayoutDirection(__r)
}

func (this *QWidget) UnsetLayoutDirection() {
	purego.SyscallN(addr_QWidget_unsetLayoutDirection, this.h)
	runtime.KeepAlive(this)
}

func (this *QWidget) SetLocale(locale *QLocale) {
	purego.SyscallN(addr_QWidget_setLocale, this.h, uintptr(locale.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(locale)
}

func (this *QWidget) Locale() *QLocale {
	__r, _, _ := purego.SyscallN(addr_QWidget_locale, this.h)
	runtime.KeepAlive(this)
	return newownedQLocale(__r)
}

func (this *QWidget) UnsetLocale() {
	purego.SyscallN(addr_QWidget_unsetLocale, this.h)
	runtime.KeepAlive(this)
}

func (this *QWidget) IsRightToLeft() bool {
	__r0, _, _ := purego.SyscallN(addr_QWidget_isRightToLeft, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QWidget) IsLeftToRight() bool {
	__r0, _, _ := purego.SyscallN(addr_QWidget_isLeftToRight, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QWidget) SetFocus() {
	purego.SyscallN(addr_QWidget_setFocus, this.h)
	runtime.KeepAlive(this)
}

func (this *QWidget) IsActiveWindow() bool {
	__r0, _, _ := purego.SyscallN(addr_QWidget_isActiveWindow, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QWidget) ActivateWindow() {
	purego.SyscallN(addr_QWidget_activateWindow, this.h)
	runtime.KeepAlive(this)
}

func (this *QWidget) ClearFocus() {
	purego.SyscallN(addr_QWidget_clearFocus, this.h)
	runtime.KeepAlive(this)
}

func (this *QWidget) SetFocusWithReason(reason FocusReason) {
	purego.SyscallN(addr_QWidget_setFocus2, this.h, uintptr(reason))
	runtime.KeepAlive(this)
}

func (this *QWidget) FocusPolicy() FocusPolicy {
	__r0, _, _ := purego.SyscallN(addr_QWidget_focusPolicy, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return FocusPolicy(__r)
}

func (this *QWidget) SetFocusPolicy(policy FocusPolicy) {
	purego.SyscallN(addr_QWidget_setFocusPolicy, this.h, uintptr(policy))
	runtime.KeepAlive(this)
}

func (this *QWidget) HasFocus() bool {
	__r0, _, _ := purego.SyscallN(addr_QWidget_hasFocus, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QWidget_SetTabOrder(arg1 *QWidget, arg2 *QWidget) {
	purego.SyscallN(addr_QWidget_setTabOrder, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()))
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
}

func (this *QWidget) SetFocusProxy(arg1 *QWidget) {
	purego.SyscallN(addr_QWidget_setFocusProxy, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QWidget) FocusProxy() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QWidget_focusProxy, this.h)
	runtime.KeepAlive(this)
	return bindQWidget(__r)
}

func (this *QWidget) ContextMenuPolicy() ContextMenuPolicy {
	__r0, _, _ := purego.SyscallN(addr_QWidget_contextMenuPolicy, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return ContextMenuPolicy(__r)
}

func (this *QWidget) SetContextMenuPolicy(policy ContextMenuPolicy) {
	purego.SyscallN(addr_QWidget_setContextMenuPolicy, this.h, uintptr(policy))
	runtime.KeepAlive(this)
}

func (this *QWidget) GrabMouse() {
	purego.SyscallN(addr_QWidget_grabMouse, this.h)
	runtime.KeepAlive(this)
}

func (this *QWidget) GrabMouseWithQCursor(arg1 *QCursor) {
	purego.SyscallN(addr_QWidget_grabMouse2, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QWidget) ReleaseMouse() {
	purego.SyscallN(addr_QWidget_releaseMouse, this.h)
	runtime.KeepAlive(this)
}

func (this *QWidget) GrabKeyboard() {
	purego.SyscallN(addr_QWidget_grabKeyboard, this.h)
	runtime.KeepAlive(this)
}

func (this *QWidget) ReleaseKeyboard() {
	purego.SyscallN(addr_QWidget_releaseKeyboard, this.h)
	runtime.KeepAlive(this)
}

func (this *QWidget) GrabShortcut(key *QKeySequence) int {
	__r0, _, _ := purego.SyscallN(addr_QWidget_grabShortcut, this.h, uintptr(key.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
	return int(__r)
}

func (this *QWidget) ReleaseShortcut(id int) {
	purego.SyscallN(addr_QWidget_releaseShortcut, this.h, uintptr(id))
	runtime.KeepAlive(this)
}

func (this *QWidget) SetShortcutEnabled(id int) {
	purego.SyscallN(addr_QWidget_setShortcutEnabled, this.h, uintptr(id))
	runtime.KeepAlive(this)
}

func (this *QWidget) SetShortcutAutoRepeat(id int) {
	purego.SyscallN(addr_QWidget_setShortcutAutoRepeat, this.h, uintptr(id))
	runtime.KeepAlive(this)
}

func QWidget_MouseGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QWidget_mouseGrabber)
	return bindQWidget(__r)
}

func QWidget_KeyboardGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QWidget_keyboardGrabber)
	return bindQWidget(__r)
}

func (this *QWidget) UpdatesEnabled() bool {
	__r0, _, _ := purego.SyscallN(addr_QWidget_updatesEnabled, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QWidget) SetUpdatesEnabled(enable bool) {
	purego.SyscallN(addr_QWidget_setUpdatesEnabled, this.h, uintptr(b2i(enable)))
	runtime.KeepAlive(this)
}

func (this *QWidget) GraphicsProxyWidget() *QGraphicsProxyWidget {
	__r, _, _ := purego.SyscallN(addr_QWidget_graphicsProxyWidget, this.h)
	runtime.KeepAlive(this)
	return bindQGraphicsProxyWidget(__r)
}

func (this *QWidget) Update() {
	purego.SyscallN(addr_QWidget_update, this.h)
	runtime.KeepAlive(this)
}

func (this *QWidget) Repaint() {
	purego.SyscallN(addr_QWidget_repaint, this.h)
	runtime.KeepAlive(this)
}

func (this *QWidget) Update2(x int, y int, w int, h int) {
	purego.SyscallN(addr_QWidget_update2, this.h, uintptr(x), uintptr(y), uintptr(w), uintptr(h))
	runtime.KeepAlive(this)
}

func (this *QWidget) UpdateWithQRect(arg1 *QRect) {
	purego.SyscallN(addr_QWidget_update3, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QWidget) UpdateWithQRegion(arg1 *QRegion) {
	purego.SyscallN(addr_QWidget_update4, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QWidget) Repaint2(x int, y int, w int, h int) {
	purego.SyscallN(addr_QWidget_repaint2, this.h, uintptr(x), uintptr(y), uintptr(w), uintptr(h))
	runtime.KeepAlive(this)
}

func (this *QWidget) RepaintWithQRect(arg1 *QRect) {
	purego.SyscallN(addr_QWidget_repaint3, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QWidget) RepaintWithQRegion(arg1 *QRegion) {
	purego.SyscallN(addr_QWidget_repaint4, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QWidget) SetVisible(visible bool) {
	purego.SyscallN(addr_QWidget_setVisible, this.h, uintptr(b2i(visible)))
	runtime.KeepAlive(this)
}

func (this *QWidget) SetHidden(hidden bool) {
	purego.SyscallN(addr_QWidget_setHidden, this.h, uintptr(b2i(hidden)))
	runtime.KeepAlive(this)
}

func (this *QWidget) Show() {
	purego.SyscallN(addr_QWidget_show, this.h)
	runtime.KeepAlive(this)
}

func (this *QWidget) Hide() {
	purego.SyscallN(addr_QWidget_hide, this.h)
	runtime.KeepAlive(this)
}

func (this *QWidget) ShowMinimized() {
	purego.SyscallN(addr_QWidget_showMinimized, this.h)
	runtime.KeepAlive(this)
}

func (this *QWidget) ShowMaximized() {
	purego.SyscallN(addr_QWidget_showMaximized, this.h)
	runtime.KeepAlive(this)
}

func (this *QWidget) ShowFullScreen() {
	purego.SyscallN(addr_QWidget_showFullScreen, this.h)
	runtime.KeepAlive(this)
}

func (this *QWidget) ShowNormal() {
	purego.SyscallN(addr_QWidget_showNormal, this.h)
	runtime.KeepAlive(this)
}

func (this *QWidget) Close() bool {
	__r0, _, _ := purego.SyscallN(addr_QWidget_close, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QWidget) Raise() {
	purego.SyscallN(addr_QWidget_raise, this.h)
	runtime.KeepAlive(this)
}

func (this *QWidget) Lower() {
	purego.SyscallN(addr_QWidget_lower, this.h)
	runtime.KeepAlive(this)
}

func (this *QWidget) StackUnder(arg1 *QWidget) {
	purego.SyscallN(addr_QWidget_stackUnder, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QWidget) Move(x int, y int) {
	purego.SyscallN(addr_QWidget_move, this.h, uintptr(x), uintptr(y))
	runtime.KeepAlive(this)
}

func (this *QWidget) MoveWithQPoint(arg1 *QPoint) {
	purego.SyscallN(addr_QWidget_move2, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QWidget) Resize(w int, h int) {
	purego.SyscallN(addr_QWidget_resize, this.h, uintptr(w), uintptr(h))
	runtime.KeepAlive(this)
}

func (this *QWidget) ResizeWithQSize(arg1 *QSize) {
	purego.SyscallN(addr_QWidget_resize2, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QWidget) SetGeometry(x int, y int, w int, h int) {
	purego.SyscallN(addr_QWidget_setGeometry, this.h, uintptr(x), uintptr(y), uintptr(w), uintptr(h))
	runtime.KeepAlive(this)
}

func (this *QWidget) SetGeometryWithQRect(arg1 *QRect) {
	purego.SyscallN(addr_QWidget_setGeometry2, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QWidget) SaveGeometry() *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QWidget_saveGeometry, this.h)
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QWidget) RestoreGeometry(geometry *QByteArray) bool {
	__r0, _, _ := purego.SyscallN(addr_QWidget_restoreGeometry, this.h, uintptr(geometry.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(geometry)
	return __r != 0
}

func (this *QWidget) AdjustSize() {
	purego.SyscallN(addr_QWidget_adjustSize, this.h)
	runtime.KeepAlive(this)
}

func (this *QWidget) IsVisible() bool {
	__r0, _, _ := purego.SyscallN(addr_QWidget_isVisible, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QWidget) IsVisibleTo(arg1 *QWidget) bool {
	__r0, _, _ := purego.SyscallN(addr_QWidget_isVisibleTo, this.h, uintptr(arg1.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return __r != 0
}

func (this *QWidget) IsHidden() bool {
	__r0, _, _ := purego.SyscallN(addr_QWidget_isHidden, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QWidget) IsMinimized() bool {
	__r0, _, _ := purego.SyscallN(addr_QWidget_isMinimized, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QWidget) IsMaximized() bool {
	__r0, _, _ := purego.SyscallN(addr_QWidget_isMaximized, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QWidget) IsFullScreen() bool {
	__r0, _, _ := purego.SyscallN(addr_QWidget_isFullScreen, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QWidget) WindowState() WindowState {
	__r0, _, _ := purego.SyscallN(addr_QWidget_windowState, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return WindowState(__r)
}

func (this *QWidget) SetWindowState(state WindowState) {
	purego.SyscallN(addr_QWidget_setWindowState, this.h, uintptr(state))
	runtime.KeepAlive(this)
}

func (this *QWidget) OverrideWindowState(state WindowState) {
	purego.SyscallN(addr_QWidget_overrideWindowState, this.h, uintptr(state))
	runtime.KeepAlive(this)
}

func (this *QWidget) SizeHint() *QSize {
	__r, _, _ := purego.SyscallN(addr_QWidget_sizeHint, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QWidget) MinimumSizeHint() *QSize {
	__r, _, _ := purego.SyscallN(addr_QWidget_minimumSizeHint, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QWidget) SizePolicy() *QSizePolicy {
	__r, _, _ := purego.SyscallN(addr_QWidget_sizePolicy, this.h)
	runtime.KeepAlive(this)
	return newownedQSizePolicy(__r)
}

func (this *QWidget) SetSizePolicy(arg1 *QSizePolicy) {
	purego.SyscallN(addr_QWidget_setSizePolicy, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QWidget) SetSizePolicy2(horizontal QSizePolicy__Policy, vertical QSizePolicy__Policy) {
	purego.SyscallN(addr_QWidget_setSizePolicy2, this.h, uintptr(horizontal), uintptr(vertical))
	runtime.KeepAlive(this)
}

func (this *QWidget) HeightForWidth(arg1 int) int {
	__r0, _, _ := purego.SyscallN(addr_QWidget_heightForWidth, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QWidget) HasHeightForWidth() bool {
	__r0, _, _ := purego.SyscallN(addr_QWidget_hasHeightForWidth, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QWidget) VisibleRegion() *QRegion {
	__r, _, _ := purego.SyscallN(addr_QWidget_visibleRegion, this.h)
	runtime.KeepAlive(this)
	return newownedQRegion(__r)
}

func (this *QWidget) SetContentsMargins(left int, top int, right int, bottom int) {
	purego.SyscallN(addr_QWidget_setContentsMargins, this.h, uintptr(left), uintptr(top), uintptr(right), uintptr(bottom))
	runtime.KeepAlive(this)
}

func (this *QWidget) SetContentsMarginsWithMargins(margins *QMargins) {
	purego.SyscallN(addr_QWidget_setContentsMargins2, this.h, uintptr(margins.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(margins)
}

func (this *QWidget) ContentsMargins() *QMargins {
	__r, _, _ := purego.SyscallN(addr_QWidget_contentsMargins, this.h)
	runtime.KeepAlive(this)
	return newownedQMargins(__r)
}

func (this *QWidget) ContentsRect() *QRect {
	__r, _, _ := purego.SyscallN(addr_QWidget_contentsRect, this.h)
	runtime.KeepAlive(this)
	return newownedQRect(__r)
}

func (this *QWidget) Layout() *QLayout {
	__r, _, _ := purego.SyscallN(addr_QWidget_layout, this.h)
	runtime.KeepAlive(this)
	return bindQLayout(__r)
}

func (this *QWidget) SetLayout(arg1 *QLayout) {
	purego.SyscallN(addr_QWidget_setLayout, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QWidget) UpdateGeometry() {
	purego.SyscallN(addr_QWidget_updateGeometry, this.h)
	runtime.KeepAlive(this)
}

func (this *QWidget) SetParent(parent *QWidget) {
	purego.SyscallN(addr_QWidget_setParent, this.h, uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
}

func (this *QWidget) SetParent2(parent *QWidget, f WindowType) {
	purego.SyscallN(addr_QWidget_setParent2, this.h, uintptr(parent.UnsafePointer()), uintptr(f))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
}

func (this *QWidget) Scroll(dx int, dy int) {
	purego.SyscallN(addr_QWidget_scroll, this.h, uintptr(dx), uintptr(dy))
	runtime.KeepAlive(this)
}

func (this *QWidget) Scroll2(dx int, dy int, arg3 *QRect) {
	purego.SyscallN(addr_QWidget_scroll2, this.h, uintptr(dx), uintptr(dy), uintptr(arg3.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg3)
}

func (this *QWidget) FocusWidget() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QWidget_focusWidget, this.h)
	runtime.KeepAlive(this)
	return bindQWidget(__r)
}

func (this *QWidget) NextInFocusChain() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QWidget_nextInFocusChain, this.h)
	runtime.KeepAlive(this)
	return bindQWidget(__r)
}

func (this *QWidget) PreviousInFocusChain() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QWidget_previousInFocusChain, this.h)
	runtime.KeepAlive(this)
	return bindQWidget(__r)
}

func (this *QWidget) AcceptDrops() bool {
	__r0, _, _ := purego.SyscallN(addr_QWidget_acceptDrops, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QWidget) SetAcceptDrops(on bool) {
	purego.SyscallN(addr_QWidget_setAcceptDrops, this.h, uintptr(b2i(on)))
	runtime.KeepAlive(this)
}

func (this *QWidget) AddAction(action *QAction) {
	purego.SyscallN(addr_QWidget_addAction, this.h, uintptr(action.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(action)
}

func (this *QWidget) AddActions(actions []*QAction) {
	actions_arr := make([]uintptr, len(actions))
	for __i, __e := range actions {
		actions_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var actions_p uintptr
	if len(actions_arr) > 0 {
		actions_p = uintptr(unsafe.Pointer(&actions_arr[0]))
	}
	purego.SyscallN(addr_QWidget_addActions, this.h, actions_p, uintptr(len(actions)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(actions)
	runtime.KeepAlive(actions_arr)
}

func (this *QWidget) InsertActions(before *QAction, actions []*QAction) {
	actions_arr := make([]uintptr, len(actions))
	for __i, __e := range actions {
		actions_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var actions_p uintptr
	if len(actions_arr) > 0 {
		actions_p = uintptr(unsafe.Pointer(&actions_arr[0]))
	}
	purego.SyscallN(addr_QWidget_insertActions, this.h, uintptr(before.UnsafePointer()), actions_p, uintptr(len(actions)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(before)
	runtime.KeepAlive(actions)
	runtime.KeepAlive(actions_arr)
}

func (this *QWidget) InsertAction(before *QAction, action *QAction) {
	purego.SyscallN(addr_QWidget_insertAction, this.h, uintptr(before.UnsafePointer()), uintptr(action.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(before)
	runtime.KeepAlive(action)
}

func (this *QWidget) RemoveAction(action *QAction) {
	purego.SyscallN(addr_QWidget_removeAction, this.h, uintptr(action.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(action)
}

func (this *QWidget) Actions() []*QAction {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QWidget_actions, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QAction, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = bindQAction(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QWidget) AddActionWithText(text string) *QAction {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QWidget_addAction2, this.h, text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
	return bindQAction(__r)
}

func (this *QWidget) AddAction2(icon *QIcon, text string) *QAction {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QWidget_addAction3, this.h, uintptr(icon.UnsafePointer()), text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(icon)
	runtime.KeepAlive(text)
	return bindQAction(__r)
}

func (this *QWidget) AddAction3(text string, receiver *QObject, member string) *QAction {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QWidget_addAction4, this.h, text_p, uintptr(len(text)), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return bindQAction(__r)
}

func (this *QWidget) AddAction4(icon *QIcon, text string, receiver *QObject, member string) *QAction {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QWidget_addAction5, this.h, uintptr(icon.UnsafePointer()), text_p, uintptr(len(text)), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(this)
	runtime.KeepAlive(icon)
	runtime.KeepAlive(text)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return bindQAction(__r)
}

func (this *QWidget) AddAction5(text string, shortcut *QKeySequence) *QAction {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QWidget_addAction6, this.h, text_p, uintptr(len(text)), uintptr(shortcut.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
	runtime.KeepAlive(shortcut)
	return bindQAction(__r)
}

func (this *QWidget) AddAction6(icon *QIcon, text string, shortcut *QKeySequence) *QAction {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QWidget_addAction7, this.h, uintptr(icon.UnsafePointer()), text_p, uintptr(len(text)), uintptr(shortcut.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(icon)
	runtime.KeepAlive(text)
	runtime.KeepAlive(shortcut)
	return bindQAction(__r)
}

func (this *QWidget) AddAction7(text string, shortcut *QKeySequence, receiver *QObject, member string) *QAction {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QWidget_addAction8, this.h, text_p, uintptr(len(text)), uintptr(shortcut.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
	runtime.KeepAlive(shortcut)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return bindQAction(__r)
}

func (this *QWidget) AddAction8(icon *QIcon, text string, shortcut *QKeySequence, receiver *QObject, member string) *QAction {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QWidget_addAction9, this.h, uintptr(icon.UnsafePointer()), text_p, uintptr(len(text)), uintptr(shortcut.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(this)
	runtime.KeepAlive(icon)
	runtime.KeepAlive(text)
	runtime.KeepAlive(shortcut)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return bindQAction(__r)
}

func (this *QWidget) ParentWidget() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QWidget_parentWidget, this.h)
	runtime.KeepAlive(this)
	return bindQWidget(__r)
}

func (this *QWidget) SetWindowFlags(type_ WindowType) {
	purego.SyscallN(addr_QWidget_setWindowFlags, this.h, uintptr(type_))
	runtime.KeepAlive(this)
}

func (this *QWidget) WindowFlags() WindowType {
	__r0, _, _ := purego.SyscallN(addr_QWidget_windowFlags, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return WindowType(__r)
}

func (this *QWidget) SetWindowFlag(arg1 WindowType) {
	purego.SyscallN(addr_QWidget_setWindowFlag, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QWidget) OverrideWindowFlags(type_ WindowType) {
	purego.SyscallN(addr_QWidget_overrideWindowFlags, this.h, uintptr(type_))
	runtime.KeepAlive(this)
}

func (this *QWidget) WindowType() WindowType {
	__r0, _, _ := purego.SyscallN(addr_QWidget_windowType, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return WindowType(__r)
}

func QWidget_Find(arg1 uint64) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QWidget_find, uintptr(arg1))
	return bindQWidget(__r)
}

func (this *QWidget) ChildAt(x int, y int) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QWidget_childAt, this.h, uintptr(x), uintptr(y))
	runtime.KeepAlive(this)
	return bindQWidget(__r)
}

func (this *QWidget) ChildAtWithQPoint(p *QPoint) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QWidget_childAt2, this.h, uintptr(p.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
	return bindQWidget(__r)
}

func (this *QWidget) ChildAtWithQPointF(p *QPointF) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QWidget_childAt3, this.h, uintptr(p.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
	return bindQWidget(__r)
}

func (this *QWidget) SetAttribute(arg1 WidgetAttribute) {
	purego.SyscallN(addr_QWidget_setAttribute, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QWidget) TestAttribute(arg1 WidgetAttribute) bool {
	__r0, _, _ := purego.SyscallN(addr_QWidget_testAttribute, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QWidget) PaintEngine() *QPaintEngine {
	__r, _, _ := purego.SyscallN(addr_QWidget_paintEngine, this.h)
	runtime.KeepAlive(this)
	return bindQPaintEngine(__r)
}

func (this *QWidget) EnsurePolished() {
	purego.SyscallN(addr_QWidget_ensurePolished, this.h)
	runtime.KeepAlive(this)
}

func (this *QWidget) IsAncestorOf(child *QWidget) bool {
	__r0, _, _ := purego.SyscallN(addr_QWidget_isAncestorOf, this.h, uintptr(child.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(child)
	return __r != 0
}

func (this *QWidget) AutoFillBackground() bool {
	__r0, _, _ := purego.SyscallN(addr_QWidget_autoFillBackground, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QWidget) SetAutoFillBackground(enabled bool) {
	purego.SyscallN(addr_QWidget_setAutoFillBackground, this.h, uintptr(b2i(enabled)))
	runtime.KeepAlive(this)
}

func (this *QWidget) BackingStore() *QBackingStore {
	__r, _, _ := purego.SyscallN(addr_QWidget_backingStore, this.h)
	runtime.KeepAlive(this)
	return bindQBackingStore(__r)
}

func (this *QWidget) WindowHandle() *QWindow {
	__r, _, _ := purego.SyscallN(addr_QWidget_windowHandle, this.h)
	runtime.KeepAlive(this)
	return bindQWindow(__r)
}

func (this *QWidget) Screen() *QScreen {
	__r, _, _ := purego.SyscallN(addr_QWidget_screen, this.h)
	runtime.KeepAlive(this)
	return bindQScreen(__r)
}

func (this *QWidget) SetScreen(arg1 *QScreen) {
	purego.SyscallN(addr_QWidget_setScreen, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func QWidget_CreateWindowContainer(window *QWindow) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QWidget_createWindowContainer, uintptr(window.UnsafePointer()))
	runtime.KeepAlive(window)
	return bindQWidget(__r)
}

func (this *QWidget) WindowTitleChanged(title string) {
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	purego.SyscallN(addr_QWidget_windowTitleChanged, this.h, title_p, uintptr(len(title)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(title)
}

func (this *QWidget) WindowIconChanged(icon *QIcon) {
	purego.SyscallN(addr_QWidget_windowIconChanged, this.h, uintptr(icon.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(icon)
}

func (this *QWidget) WindowIconTextChanged(iconText string) {
	var iconText_p uintptr
	if len(iconText) > 0 {
		iconText_p = uintptr(unsafe.Pointer(unsafe.StringData(iconText)))
	}
	purego.SyscallN(addr_QWidget_windowIconTextChanged, this.h, iconText_p, uintptr(len(iconText)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(iconText)
}

func (this *QWidget) CustomContextMenuRequested(pos *QPoint) {
	purego.SyscallN(addr_QWidget_customContextMenuRequested, this.h, uintptr(pos.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pos)
}

func (this *QWidget) InputMethodQuery(arg1 InputMethodQuery) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QWidget_inputMethodQuery, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QWidget) InputMethodHints() InputMethodHint {
	__r0, _, _ := purego.SyscallN(addr_QWidget_inputMethodHints, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return InputMethodHint(__r)
}

func (this *QWidget) SetInputMethodHints(hints InputMethodHint) {
	purego.SyscallN(addr_QWidget_setInputMethodHints, this.h, uintptr(hints))
	runtime.KeepAlive(this)
}

func (this *QWidget) UpdateMicroFocus() {
	purego.SyscallN(addr_QWidget_updateMicroFocus, this.h)
	runtime.KeepAlive(this)
}

func (this *QWidget) Create() {
	purego.SyscallN(addr_QWidget_create, this.h)
	runtime.KeepAlive(this)
}

func (this *QWidget) Destroy() {
	purego.SyscallN(addr_QWidget_destroy, this.h)
	runtime.KeepAlive(this)
}

func (this *QWidget) FocusNextChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QWidget_focusNextChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QWidget) FocusPreviousChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QWidget_focusPreviousChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QWidget_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QWidget_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QWidget_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QWidget_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QWidget) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QWidget_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QWidget) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QWidget_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QWidget) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QWidget_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QWidget) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QWidget_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func (this *QWidget) PaintingActive() bool {
	__r0, _, _ := purego.SyscallN(addr_QWidget_paintingActive, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QWidget) WidthMM() int {
	__r0, _, _ := purego.SyscallN(addr_QWidget_widthMM, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QWidget) HeightMM() int {
	__r0, _, _ := purego.SyscallN(addr_QWidget_heightMM, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QWidget) LogicalDpiX() int {
	__r0, _, _ := purego.SyscallN(addr_QWidget_logicalDpiX, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QWidget) LogicalDpiY() int {
	__r0, _, _ := purego.SyscallN(addr_QWidget_logicalDpiY, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QWidget) PhysicalDpiX() int {
	__r0, _, _ := purego.SyscallN(addr_QWidget_physicalDpiX, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QWidget) PhysicalDpiY() int {
	__r0, _, _ := purego.SyscallN(addr_QWidget_physicalDpiY, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QWidget) DevicePixelRatio() float64 {
	__r := pg_QWidget_devicePixelRatio(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QWidget) DevicePixelRatioF() float64 {
	__r := pg_QWidget_devicePixelRatioF(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QWidget) ColorCount() int {
	__r0, _, _ := purego.SyscallN(addr_QWidget_colorCount, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QWidget) Depth() int {
	__r0, _, _ := purego.SyscallN(addr_QWidget_depth, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func QWidget_DevicePixelRatioFScale() float64 {
	__r := pg_QWidget_devicePixelRatioFScale()
	return __r
}

func QWidget_EncodeMetricF(metric QPaintDevice__PaintDeviceMetric, value float64) int {
	__r := pg_QWidget_encodeMetricF(int32(metric), value)
	return int(__r)
}

func (this *QWidget) GetDecodedMetricF(metricA QPaintDevice__PaintDeviceMetric, metricB QPaintDevice__PaintDeviceMetric) float64 {
	__r := pg_QWidget_getDecodedMetricF(this.h, int32(metricA), int32(metricB))
	runtime.KeepAlive(this)
	return __r
}

func QWidget_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QWidget_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QWidget_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QWidget_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func (this *QWidget) Render2(target *QPaintDevice, targetOffset *QPoint) {
	purego.SyscallN(addr_QWidget_render3, this.h, uintptr(target.UnsafePointer()), uintptr(targetOffset.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(target)
	runtime.KeepAlive(targetOffset)
}

func (this *QWidget) Render3(target *QPaintDevice, targetOffset *QPoint, sourceRegion *QRegion) {
	purego.SyscallN(addr_QWidget_render4, this.h, uintptr(target.UnsafePointer()), uintptr(targetOffset.UnsafePointer()), uintptr(sourceRegion.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(target)
	runtime.KeepAlive(targetOffset)
	runtime.KeepAlive(sourceRegion)
}

func (this *QWidget) Render4(target *QPaintDevice, targetOffset *QPoint, sourceRegion *QRegion, renderFlags QWidget__RenderFlag) {
	purego.SyscallN(addr_QWidget_render5, this.h, uintptr(target.UnsafePointer()), uintptr(targetOffset.UnsafePointer()), uintptr(sourceRegion.UnsafePointer()), uintptr(renderFlags))
	runtime.KeepAlive(this)
	runtime.KeepAlive(target)
	runtime.KeepAlive(targetOffset)
	runtime.KeepAlive(sourceRegion)
}

func (this *QWidget) Render5(painter *QPainter, targetOffset *QPoint) {
	purego.SyscallN(addr_QWidget_render6, this.h, uintptr(painter.UnsafePointer()), uintptr(targetOffset.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
	runtime.KeepAlive(targetOffset)
}

func (this *QWidget) Render6(painter *QPainter, targetOffset *QPoint, sourceRegion *QRegion) {
	purego.SyscallN(addr_QWidget_render7, this.h, uintptr(painter.UnsafePointer()), uintptr(targetOffset.UnsafePointer()), uintptr(sourceRegion.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
	runtime.KeepAlive(targetOffset)
	runtime.KeepAlive(sourceRegion)
}

func (this *QWidget) Render7(painter *QPainter, targetOffset *QPoint, sourceRegion *QRegion, renderFlags QWidget__RenderFlag) {
	purego.SyscallN(addr_QWidget_render8, this.h, uintptr(painter.UnsafePointer()), uintptr(targetOffset.UnsafePointer()), uintptr(sourceRegion.UnsafePointer()), uintptr(renderFlags))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
	runtime.KeepAlive(targetOffset)
	runtime.KeepAlive(sourceRegion)
}

func (this *QWidget) GrabWithRectangle(rectangle *QRect) *QPixmap {
	__r, _, _ := purego.SyscallN(addr_QWidget_grab2, this.h, uintptr(rectangle.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rectangle)
	return newownedQPixmap(__r)
}

func (this *QWidget) GrabGesture2(type_ GestureType, flags GestureFlag) {
	purego.SyscallN(addr_QWidget_grabGesture2, this.h, uintptr(type_), uintptr(flags))
	runtime.KeepAlive(this)
}

func (this *QWidget) GrabShortcut2(key *QKeySequence, context ShortcutContext) int {
	__r0, _, _ := purego.SyscallN(addr_QWidget_grabShortcut2, this.h, uintptr(key.UnsafePointer()), uintptr(context))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
	return int(__r)
}

func (this *QWidget) SetShortcutEnabled2(id int, enable bool) {
	purego.SyscallN(addr_QWidget_setShortcutEnabled2, this.h, uintptr(id), uintptr(b2i(enable)))
	runtime.KeepAlive(this)
}

func (this *QWidget) SetShortcutAutoRepeat2(id int, enable bool) {
	purego.SyscallN(addr_QWidget_setShortcutAutoRepeat2, this.h, uintptr(id), uintptr(b2i(enable)))
	runtime.KeepAlive(this)
}

func (this *QWidget) AddAction9(text string, receiver *QObject, member string, type_ ConnectionType) *QAction {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QWidget_addAction10, this.h, text_p, uintptr(len(text)), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(type_))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return bindQAction(__r)
}

func (this *QWidget) AddAction10(icon *QIcon, text string, receiver *QObject, member string, type_ ConnectionType) *QAction {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QWidget_addAction11, this.h, uintptr(icon.UnsafePointer()), text_p, uintptr(len(text)), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(type_))
	runtime.KeepAlive(this)
	runtime.KeepAlive(icon)
	runtime.KeepAlive(text)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return bindQAction(__r)
}

func (this *QWidget) AddAction11(text string, shortcut *QKeySequence, receiver *QObject, member string, type_ ConnectionType) *QAction {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QWidget_addAction12, this.h, text_p, uintptr(len(text)), uintptr(shortcut.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(type_))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
	runtime.KeepAlive(shortcut)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return bindQAction(__r)
}

func (this *QWidget) AddAction12(icon *QIcon, text string, shortcut *QKeySequence, receiver *QObject, member string, type_ ConnectionType) *QAction {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QWidget_addAction13, this.h, uintptr(icon.UnsafePointer()), text_p, uintptr(len(text)), uintptr(shortcut.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(type_))
	runtime.KeepAlive(this)
	runtime.KeepAlive(icon)
	runtime.KeepAlive(text)
	runtime.KeepAlive(shortcut)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return bindQAction(__r)
}

func (this *QWidget) SetWindowFlag2(arg1 WindowType, on bool) {
	purego.SyscallN(addr_QWidget_setWindowFlag2, this.h, uintptr(arg1), uintptr(b2i(on)))
	runtime.KeepAlive(this)
}

func (this *QWidget) SetAttribute2(arg1 WidgetAttribute, on bool) {
	purego.SyscallN(addr_QWidget_setAttribute2, this.h, uintptr(arg1), uintptr(b2i(on)))
	runtime.KeepAlive(this)
}

func QWidget_CreateWindowContainer2(window *QWindow, parent *QWidget) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QWidget_createWindowContainer2, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func QWidget_CreateWindowContainer3(window *QWindow, parent *QWidget, flags WindowType) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QWidget_createWindowContainer3, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()), uintptr(flags))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func (this *QWidget) UpdateMicroFocusWithQuery(query InputMethodQuery) {
	purego.SyscallN(addr_QWidget_updateMicroFocus2, this.h, uintptr(query))
	runtime.KeepAlive(this)
}

func (this *QWidget) CreateWith(arg1 uint64) {
	purego.SyscallN(addr_QWidget_create2, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QWidget) Create2(arg1 uint64, initializeWindow bool) {
	purego.SyscallN(addr_QWidget_create3, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)))
	runtime.KeepAlive(this)
}

func (this *QWidget) Create3(arg1 uint64, initializeWindow bool, destroyOldWindow bool) {
	purego.SyscallN(addr_QWidget_create4, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)), uintptr(b2i(destroyOldWindow)))
	runtime.KeepAlive(this)
}

func (this *QWidget) DestroyWithDestroyWindow(destroyWindow bool) {
	purego.SyscallN(addr_QWidget_destroy2, this.h, uintptr(b2i(destroyWindow)))
	runtime.KeepAlive(this)
}

func (this *QWidget) Destroy2(destroyWindow bool, destroySubWindows bool) {
	purego.SyscallN(addr_QWidget_destroy3, this.h, uintptr(b2i(destroyWindow)), uintptr(b2i(destroySubWindows)))
	runtime.KeepAlive(this)
}

func QWidget_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QWidget_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QWidget_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QWidget_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QWidget) AsQPaintDevice() *QPaintDevice {
	__r, _, _ := purego.SyscallN(addr_QWidget_asQPaintDevice, this.h)
	runtime.KeepAlive(this)
	return bindQPaintDevice(__r)
}

func (this *QWidget) DevTypeDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QWidget_devType_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QWidget) SetVisibleDefault(visible bool) {
	purego.SyscallN(addr_QWidget_setVisible_default, this.h, uintptr(b2i(visible)))
	runtime.KeepAlive(this)
}

func (this *QWidget) SizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QWidget_sizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QWidget) MinimumSizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QWidget_minimumSizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QWidget) HeightForWidthDefault(arg1 int) int {
	__r0, _, _ := purego.SyscallN(addr_QWidget_heightForWidth_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QWidget) HasHeightForWidthDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QWidget_hasHeightForWidth_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QWidget) PaintEngineDefault() *QPaintEngine {
	__r, _, _ := purego.SyscallN(addr_QWidget_paintEngine_default, this.h)
	runtime.KeepAlive(this)
	return bindQPaintEngine(__r)
}

func (this *QWidget) InputMethodQueryDefault(arg1 InputMethodQuery) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QWidget_inputMethodQuery_default, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QWidget) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QWidget_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QWidget) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QWidget_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QWidget) MousePressEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QWidget_mousePressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QWidget) MouseReleaseEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QWidget_mouseReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QWidget) MouseDoubleClickEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QWidget_mouseDoubleClickEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QWidget) MouseMoveEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QWidget_mouseMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QWidget) WheelEventDefault(event *QWheelEvent) {
	purego.SyscallN(addr_QWidget_wheelEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QWidget) KeyPressEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QWidget_keyPressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QWidget) KeyReleaseEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QWidget_keyReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QWidget) FocusInEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QWidget_focusInEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QWidget) FocusOutEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QWidget_focusOutEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QWidget) EnterEventDefault(event *QEnterEvent) {
	purego.SyscallN(addr_QWidget_enterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QWidget) LeaveEventDefault(event *QEvent) {
	purego.SyscallN(addr_QWidget_leaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QWidget) PaintEventDefault(event *QPaintEvent) {
	purego.SyscallN(addr_QWidget_paintEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QWidget) MoveEventDefault(event *QMoveEvent) {
	purego.SyscallN(addr_QWidget_moveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QWidget) ResizeEventDefault(event *QResizeEvent) {
	purego.SyscallN(addr_QWidget_resizeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QWidget) CloseEventDefault(event *QCloseEvent) {
	purego.SyscallN(addr_QWidget_closeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QWidget) ContextMenuEventDefault(event *QContextMenuEvent) {
	purego.SyscallN(addr_QWidget_contextMenuEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QWidget) TabletEventDefault(event *QTabletEvent) {
	purego.SyscallN(addr_QWidget_tabletEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QWidget) ActionEventDefault(event *QActionEvent) {
	purego.SyscallN(addr_QWidget_actionEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QWidget) DragEnterEventDefault(event *QDragEnterEvent) {
	purego.SyscallN(addr_QWidget_dragEnterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QWidget) DragMoveEventDefault(event *QDragMoveEvent) {
	purego.SyscallN(addr_QWidget_dragMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QWidget) DragLeaveEventDefault(event *QDragLeaveEvent) {
	purego.SyscallN(addr_QWidget_dragLeaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QWidget) DropEventDefault(event *QDropEvent) {
	purego.SyscallN(addr_QWidget_dropEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QWidget) ShowEventDefault(event *QShowEvent) {
	purego.SyscallN(addr_QWidget_showEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QWidget) HideEventDefault(event *QHideEvent) {
	purego.SyscallN(addr_QWidget_hideEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QWidget) ChangeEventDefault(arg1 *QEvent) {
	purego.SyscallN(addr_QWidget_changeEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QWidget) MetricDefault(arg1 QPaintDevice__PaintDeviceMetric) int {
	__r0, _, _ := purego.SyscallN(addr_QWidget_metric_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QWidget) InitPainterDefault(painter *QPainter) {
	purego.SyscallN(addr_QWidget_initPainter_default, this.h, uintptr(painter.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
}

func (this *QWidget) RedirectedDefault(offset *QPoint) *QPaintDevice {
	__r, _, _ := purego.SyscallN(addr_QWidget_redirected_default, this.h, uintptr(offset.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(offset)
	return bindQPaintDevice(__r)
}

func (this *QWidget) SharedPainterDefault() *QPainter {
	__r, _, _ := purego.SyscallN(addr_QWidget_sharedPainter_default, this.h)
	runtime.KeepAlive(this)
	return bindQPainter(__r)
}

func (this *QWidget) InputMethodEventDefault(arg1 *QInputMethodEvent) {
	purego.SyscallN(addr_QWidget_inputMethodEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QWidget) FocusNextPrevChildDefault(next bool) bool {
	__r0, _, _ := purego.SyscallN(addr_QWidget_focusNextPrevChild_default, this.h, uintptr(b2i(next)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QWidget) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QWidget_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QWidget) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QWidget_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QWidget) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QWidget_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QWidget) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QWidget_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QWidget) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QWidget_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QWidget) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QWidget_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QWidget) OnWindowTitleChanged(fn func(string)) int64 {
	id := registerCallback(func(title uintptr, title_len int64) { fn(pgStrNoFree(title, title_len)) })
	pg_QWidget_connect_windowTitleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QWidget) OnWindowIconChanged(fn func(*QIcon)) int64 {
	id := registerCallback(func(icon uintptr) { fn(UnsafeNewQIcon(unsafe.Pointer(icon))) })
	pg_QWidget_connect_windowIconChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QWidget) OnWindowIconTextChanged(fn func(string)) int64 {
	id := registerCallback(func(iconText uintptr, iconText_len int64) { fn(pgStrNoFree(iconText, iconText_len)) })
	pg_QWidget_connect_windowIconTextChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QWidget) OnCustomContextMenuRequested(fn func(*QPoint)) int64 {
	id := registerCallback(func(pos uintptr) { fn(UnsafeNewQPoint(unsafe.Pointer(pos))) })
	pg_QWidget_connect_customContextMenuRequested(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QWidget) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QWidget_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QWidget) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QWidget_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QWidget) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QWidget_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QWidget) Disconnect(handle int64) {
	pg_QWidget__disconnect(handle)
	unregisterCallback(handle)
}

func QWidget_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case -8659679893868524044:
		return 1
	case 2110182719375245575:
		return 2
	case 729690261457684101:
		return 3
	case 7049264599837532491:
		return 4
	case 4311534060495737339:
		return 5
	case -8266762397995423915:
		return 6
	case 6398007404756346381:
		return 7
	case 1237236456312824991:
		return 8
	case -332090898078543319:
		return 9
	case -4415920049297795213:
		return 10
	case -4029073905697023887:
		return 11
	case -4255494004099806581:
		return 12
	case 1919988340934511316:
		return 13
	case 1023791977679793185:
		return 14
	case 772598475950050299:
		return 15
	case 975043767033552552:
		return 16
	case -385596010576048993:
		return 17
	case 6316270346668237895:
		return 18
	case -551070009623711640:
		return 19
	case 5290442163167458465:
		return 20
	case -9161949150489525396:
		return 21
	case -4302917285655014661:
		return 22
	case -1989937050019464783:
		return 23
	case -6698043222227006015:
		return 24
	case -212068866860701141:
		return 25
	case 1697444140088555581:
		return 26
	case 206122715919917589:
		return 27
	case -1781545137352745978:
		return 28
	case -3139904158373725006:
		return 29
	case 9107058682754591716:
		return 30
	case -6825369174751916:
		return 31
	case -199718707876719129:
		return 32
	case -2039166586246443706:
		return 33
	case 4959376063826165943:
		return 34
	case 427629649267311777:
		return 35
	case 6477500296980050210:
		return 36
	case 4127240633310655048:
		return 37
	case 5792771328741774991:
		return 38
	case 5438819661942843222:
		return 39
	case 7345853097340709638:
		return 40
	case 5898328517505784191:
		return 41
	case 5217623514630858297:
		return 42
	case -5409927581812374498:
		return 43
	case -1185315367526857833:
		return 44
	case 3769868229123868474:
		return 45
	case 7379029982170503008:
		return 46
	case -568369772982215962:
		return 47
	}
	return -1
}

func QWidget_enable(h uintptr, idx int32) { pg_QWidget__set_override(h, idx) }

func (this *QWidget) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QWidget) OnDevType(fn func() int) {
	registerVirt(this.UnsafeGoID(), -8659679893868524044, func() int32 { return int32(fn()) }) // slot "devType"
}

func (this *QWidget) OnSetVisible(fn func(bool)) {
	registerVirt(this.UnsafeGoID(), 2110182719375245575, func(visible int32) { fn(visible != 0) }) // slot "setVisible"
}

func (this *QWidget) OnSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 729690261457684101, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sizeHint"
}

func (this *QWidget) OnMinimumSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 7049264599837532491, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "minimumSizeHint"
}

func (this *QWidget) OnHeightForWidth(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), 4311534060495737339, func(arg1 int32) int32 { return int32(fn(int(arg1))) }) // slot "heightForWidth"
}

func (this *QWidget) OnHasHeightForWidth(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -8266762397995423915, func() int32 { return b2i(fn()) }) // slot "hasHeightForWidth"
}

func (this *QWidget) OnPaintEngine(fn func() *QPaintEngine) {
	registerVirt(this.UnsafeGoID(), 6398007404756346381, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "paintEngine"
}

func (this *QWidget) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QWidget) OnMousePressEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -332090898078543319, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mousePressEvent"
}

func (this *QWidget) OnMouseReleaseEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4415920049297795213, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseReleaseEvent"
}

func (this *QWidget) OnMouseDoubleClickEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4029073905697023887, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseDoubleClickEvent"
}

func (this *QWidget) OnMouseMoveEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4255494004099806581, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseMoveEvent"
}

func (this *QWidget) OnWheelEvent(fn func(*QWheelEvent)) {
	registerVirt(this.UnsafeGoID(), 1919988340934511316, func(event uintptr) { fn(UnsafeNewQWheelEvent(unsafe.Pointer(event))) }) // slot "wheelEvent"
}

func (this *QWidget) OnKeyPressEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 1023791977679793185, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyPressEvent"
}

func (this *QWidget) OnKeyReleaseEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 772598475950050299, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyReleaseEvent"
}

func (this *QWidget) OnFocusInEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), 975043767033552552, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusInEvent"
}

func (this *QWidget) OnFocusOutEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), -385596010576048993, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusOutEvent"
}

func (this *QWidget) OnEnterEvent(fn func(*QEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 6316270346668237895, func(event uintptr) { fn(UnsafeNewQEnterEvent(unsafe.Pointer(event))) }) // slot "enterEvent"
}

func (this *QWidget) OnLeaveEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), -551070009623711640, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "leaveEvent"
}

func (this *QWidget) OnPaintEvent(fn func(*QPaintEvent)) {
	registerVirt(this.UnsafeGoID(), 5290442163167458465, func(event uintptr) { fn(UnsafeNewQPaintEvent(unsafe.Pointer(event))) }) // slot "paintEvent"
}

func (this *QWidget) OnMoveEvent(fn func(*QMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -9161949150489525396, func(event uintptr) { fn(UnsafeNewQMoveEvent(unsafe.Pointer(event))) }) // slot "moveEvent"
}

func (this *QWidget) OnResizeEvent(fn func(*QResizeEvent)) {
	registerVirt(this.UnsafeGoID(), -4302917285655014661, func(event uintptr) { fn(UnsafeNewQResizeEvent(unsafe.Pointer(event))) }) // slot "resizeEvent"
}

func (this *QWidget) OnCloseEvent(fn func(*QCloseEvent)) {
	registerVirt(this.UnsafeGoID(), -1989937050019464783, func(event uintptr) { fn(UnsafeNewQCloseEvent(unsafe.Pointer(event))) }) // slot "closeEvent"
}

func (this *QWidget) OnContextMenuEvent(fn func(*QContextMenuEvent)) {
	registerVirt(this.UnsafeGoID(), -6698043222227006015, func(event uintptr) { fn(UnsafeNewQContextMenuEvent(unsafe.Pointer(event))) }) // slot "contextMenuEvent"
}

func (this *QWidget) OnTabletEvent(fn func(*QTabletEvent)) {
	registerVirt(this.UnsafeGoID(), -212068866860701141, func(event uintptr) { fn(UnsafeNewQTabletEvent(unsafe.Pointer(event))) }) // slot "tabletEvent"
}

func (this *QWidget) OnActionEvent(fn func(*QActionEvent)) {
	registerVirt(this.UnsafeGoID(), 1697444140088555581, func(event uintptr) { fn(UnsafeNewQActionEvent(unsafe.Pointer(event))) }) // slot "actionEvent"
}

func (this *QWidget) OnDragEnterEvent(fn func(*QDragEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 206122715919917589, func(event uintptr) { fn(UnsafeNewQDragEnterEvent(unsafe.Pointer(event))) }) // slot "dragEnterEvent"
}

func (this *QWidget) OnDragMoveEvent(fn func(*QDragMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -1781545137352745978, func(event uintptr) { fn(UnsafeNewQDragMoveEvent(unsafe.Pointer(event))) }) // slot "dragMoveEvent"
}

func (this *QWidget) OnDragLeaveEvent(fn func(*QDragLeaveEvent)) {
	registerVirt(this.UnsafeGoID(), -3139904158373725006, func(event uintptr) { fn(UnsafeNewQDragLeaveEvent(unsafe.Pointer(event))) }) // slot "dragLeaveEvent"
}

func (this *QWidget) OnDropEvent(fn func(*QDropEvent)) {
	registerVirt(this.UnsafeGoID(), 9107058682754591716, func(event uintptr) { fn(UnsafeNewQDropEvent(unsafe.Pointer(event))) }) // slot "dropEvent"
}

func (this *QWidget) OnShowEvent(fn func(*QShowEvent)) {
	registerVirt(this.UnsafeGoID(), -6825369174751916, func(event uintptr) { fn(UnsafeNewQShowEvent(unsafe.Pointer(event))) }) // slot "showEvent"
}

func (this *QWidget) OnHideEvent(fn func(*QHideEvent)) {
	registerVirt(this.UnsafeGoID(), -199718707876719129, func(event uintptr) { fn(UnsafeNewQHideEvent(unsafe.Pointer(event))) }) // slot "hideEvent"
}

func (this *QWidget) OnNativeEvent(fn func(*QByteArray, unsafe.Pointer, *int64) bool) {
	registerVirt(this.UnsafeGoID(), -2039166586246443706, func(eventType uintptr, message uintptr, result uintptr) int32 { return b2i(fn(UnsafeNewQByteArray(unsafe.Pointer(eventType)), unsafe.Pointer(message), (*int64)(unsafe.Pointer(result)))) }) // slot "nativeEvent"
}

func (this *QWidget) OnChangeEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 4959376063826165943, func(arg1 uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(arg1))) }) // slot "changeEvent"
}

func (this *QWidget) OnMetric(fn func(QPaintDevice__PaintDeviceMetric) int) {
	registerVirt(this.UnsafeGoID(), 427629649267311777, func(arg1 int32) int32 { return int32(fn(QPaintDevice__PaintDeviceMetric(arg1))) }) // slot "metric"
}

func (this *QWidget) OnInitPainter(fn func(*QPainter)) {
	registerVirt(this.UnsafeGoID(), 6477500296980050210, func(painter uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter))) }) // slot "initPainter"
}

func (this *QWidget) OnRedirected(fn func(*QPoint) *QPaintDevice) {
	registerVirt(this.UnsafeGoID(), 4127240633310655048, func(offset uintptr) uintptr { return uintptr(fn(UnsafeNewQPoint(unsafe.Pointer(offset))).UnsafePointer()) }) // slot "redirected"
}

func (this *QWidget) OnSharedPainter(fn func() *QPainter) {
	registerVirt(this.UnsafeGoID(), 5792771328741774991, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sharedPainter"
}

func (this *QWidget) OnInputMethodEvent(fn func(*QInputMethodEvent)) {
	registerVirt(this.UnsafeGoID(), 5438819661942843222, func(arg1 uintptr) { fn(UnsafeNewQInputMethodEvent(unsafe.Pointer(arg1))) }) // slot "inputMethodEvent"
}

func (this *QWidget) OnInputMethodQuery(fn func(InputMethodQuery) *QVariant) {
	registerVirt(this.UnsafeGoID(), 7345853097340709638, func(arg1 int32) uintptr { return uintptr(fn(InputMethodQuery(arg1)).UnsafePointer()) }) // slot "inputMethodQuery"
}

func (this *QWidget) OnFocusNextPrevChild(fn func(bool) bool) {
	registerVirt(this.UnsafeGoID(), 5898328517505784191, func(next int32) int32 { return b2i(fn(next != 0)) }) // slot "focusNextPrevChild"
}

func (this *QWidget) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QWidget) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QWidget) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QWidget) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QWidget) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QWidget) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QWidget) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QWidget_delete(this.h)
	}
	this.UnsafeRelease()
}
