package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QGraphicsWidget_chain uintptr
	addr_QGraphicsWidget_new uintptr
	addr_QGraphicsWidget_new2 uintptr
	addr_QGraphicsWidget_new3 uintptr
	addr_QGraphicsWidget_tr uintptr
	addr_QGraphicsWidget_layout uintptr
	addr_QGraphicsWidget_setLayout uintptr
	addr_QGraphicsWidget_adjustSize uintptr
	addr_QGraphicsWidget_layoutDirection uintptr
	addr_QGraphicsWidget_setLayoutDirection uintptr
	addr_QGraphicsWidget_unsetLayoutDirection uintptr
	addr_QGraphicsWidget_style uintptr
	addr_QGraphicsWidget_setStyle uintptr
	addr_QGraphicsWidget_font uintptr
	addr_QGraphicsWidget_setFont uintptr
	addr_QGraphicsWidget_palette uintptr
	addr_QGraphicsWidget_setPalette uintptr
	addr_QGraphicsWidget_autoFillBackground uintptr
	addr_QGraphicsWidget_setAutoFillBackground uintptr
	addr_QGraphicsWidget_resize uintptr
	pg_QGraphicsWidget_resize2 func(uintptr, float64, float64)
	addr_QGraphicsWidget_size uintptr
	addr_QGraphicsWidget_setGeometry uintptr
	pg_QGraphicsWidget_setGeometry2 func(uintptr, float64, float64, float64, float64)
	addr_QGraphicsWidget_rect uintptr
	pg_QGraphicsWidget_setContentsMargins func(uintptr, float64, float64, float64, float64)
	addr_QGraphicsWidget_setContentsMargins2 uintptr
	addr_QGraphicsWidget_getContentsMargins uintptr
	pg_QGraphicsWidget_setWindowFrameMargins func(uintptr, float64, float64, float64, float64)
	addr_QGraphicsWidget_setWindowFrameMargins2 uintptr
	addr_QGraphicsWidget_getWindowFrameMargins uintptr
	addr_QGraphicsWidget_unsetWindowFrameMargins uintptr
	addr_QGraphicsWidget_windowFrameGeometry uintptr
	addr_QGraphicsWidget_windowFrameRect uintptr
	addr_QGraphicsWidget_windowFlags uintptr
	addr_QGraphicsWidget_windowType uintptr
	addr_QGraphicsWidget_setWindowFlags uintptr
	addr_QGraphicsWidget_isActiveWindow uintptr
	addr_QGraphicsWidget_setWindowTitle uintptr
	addr_QGraphicsWidget_windowTitle uintptr
	addr_QGraphicsWidget_focusPolicy uintptr
	addr_QGraphicsWidget_setFocusPolicy uintptr
	addr_QGraphicsWidget_setTabOrder uintptr
	addr_QGraphicsWidget_focusWidget uintptr
	addr_QGraphicsWidget_grabShortcut uintptr
	addr_QGraphicsWidget_releaseShortcut uintptr
	addr_QGraphicsWidget_setShortcutEnabled uintptr
	addr_QGraphicsWidget_setShortcutAutoRepeat uintptr
	addr_QGraphicsWidget_addAction uintptr
	addr_QGraphicsWidget_addActions uintptr
	addr_QGraphicsWidget_insertActions uintptr
	addr_QGraphicsWidget_insertAction uintptr
	addr_QGraphicsWidget_removeAction uintptr
	addr_QGraphicsWidget_actions uintptr
	addr_QGraphicsWidget_setAttribute uintptr
	addr_QGraphicsWidget_testAttribute uintptr
	addr_QGraphicsWidget_paintWindowFrame uintptr
	addr_QGraphicsWidget_geometryChanged uintptr
	addr_QGraphicsWidget_layoutChanged uintptr
	addr_QGraphicsWidget_close uintptr
	addr_QGraphicsWidget_updateMicroFocus uintptr
	addr_QGraphicsWidget_connect uintptr
	addr_QGraphicsWidget_connect2 uintptr
	addr_QGraphicsWidget_sender uintptr
	addr_QGraphicsWidget_senderSignalIndex uintptr
	addr_QGraphicsWidget_receivers uintptr
	addr_QGraphicsWidget_isSignalConnected uintptr
	addr_QGraphicsWidget_addToIndex uintptr
	addr_QGraphicsWidget_removeFromIndex uintptr
	addr_QGraphicsWidget_prepareGeometryChange uintptr
	addr_QGraphicsWidget_setSizePolicy uintptr
	addr_QGraphicsWidget_setSizePolicy2 uintptr
	addr_QGraphicsWidget_sizePolicy uintptr
	addr_QGraphicsWidget_setMinimumSize uintptr
	pg_QGraphicsWidget_setMinimumSize2 func(uintptr, float64, float64)
	addr_QGraphicsWidget_minimumSize uintptr
	pg_QGraphicsWidget_setMinimumWidth func(uintptr, float64)
	pg_QGraphicsWidget_minimumWidth func(uintptr) float64
	pg_QGraphicsWidget_setMinimumHeight func(uintptr, float64)
	pg_QGraphicsWidget_minimumHeight func(uintptr) float64
	addr_QGraphicsWidget_setPreferredSize uintptr
	pg_QGraphicsWidget_setPreferredSize2 func(uintptr, float64, float64)
	addr_QGraphicsWidget_preferredSize uintptr
	pg_QGraphicsWidget_setPreferredWidth func(uintptr, float64)
	pg_QGraphicsWidget_preferredWidth func(uintptr) float64
	pg_QGraphicsWidget_setPreferredHeight func(uintptr, float64)
	pg_QGraphicsWidget_preferredHeight func(uintptr) float64
	addr_QGraphicsWidget_setMaximumSize uintptr
	pg_QGraphicsWidget_setMaximumSize2 func(uintptr, float64, float64)
	addr_QGraphicsWidget_maximumSize uintptr
	pg_QGraphicsWidget_setMaximumWidth func(uintptr, float64)
	pg_QGraphicsWidget_maximumWidth func(uintptr) float64
	pg_QGraphicsWidget_setMaximumHeight func(uintptr, float64)
	pg_QGraphicsWidget_maximumHeight func(uintptr) float64
	addr_QGraphicsWidget_geometry uintptr
	addr_QGraphicsWidget_contentsRect uintptr
	addr_QGraphicsWidget_effectiveSizeHint uintptr
	addr_QGraphicsWidget_isEmpty uintptr
	addr_QGraphicsWidget_parentLayoutItem uintptr
	addr_QGraphicsWidget_setParentLayoutItem uintptr
	addr_QGraphicsWidget_isLayout uintptr
	addr_QGraphicsWidget_graphicsItem uintptr
	addr_QGraphicsWidget_ownedByLayout uintptr
	addr_QGraphicsWidget_setGraphicsItem uintptr
	addr_QGraphicsWidget_setOwnedByLayout uintptr
	addr_QGraphicsWidget_tr2 uintptr
	addr_QGraphicsWidget_tr3 uintptr
	addr_QGraphicsWidget_grabShortcut2 uintptr
	addr_QGraphicsWidget_setShortcutEnabled2 uintptr
	addr_QGraphicsWidget_setShortcutAutoRepeat2 uintptr
	addr_QGraphicsWidget_setAttribute2 uintptr
	addr_QGraphicsWidget_connect3 uintptr
	addr_QGraphicsWidget_connect4 uintptr
	addr_QGraphicsWidget_setSizePolicy3 uintptr
	addr_QGraphicsWidget_effectiveSizeHint2 uintptr
	addr_QGraphicsWidget_asQGraphicsLayoutItem uintptr
	addr_QGraphicsWidget_setGeometry_default uintptr
	pg_QGraphicsWidget_setGeometry_default2 func(uintptr, float64, float64, float64, float64)
	addr_QGraphicsWidget_getContentsMargins_default uintptr
	addr_QGraphicsWidget_paintWindowFrame_default uintptr
	addr_QGraphicsWidget_isEmpty_default uintptr
	addr_QGraphicsWidget_metaObject_default uintptr
	addr_QGraphicsWidget_type_default uintptr
	addr_QGraphicsWidget_paint_default uintptr
	addr_QGraphicsWidget_boundingRect_default uintptr
	addr_QGraphicsWidget_shape_default uintptr
	addr_QGraphicsWidget_initStyleOption_default uintptr
	addr_QGraphicsWidget_sizeHint_default uintptr
	addr_QGraphicsWidget_updateGeometry_default uintptr
	addr_QGraphicsWidget_itemChange_default uintptr
	addr_QGraphicsWidget_propertyChange_default uintptr
	addr_QGraphicsWidget_sceneEvent_default uintptr
	addr_QGraphicsWidget_windowFrameEvent_default uintptr
	addr_QGraphicsWidget_windowFrameSectionAt_default uintptr
	addr_QGraphicsWidget_event_default uintptr
	addr_QGraphicsWidget_changeEvent_default uintptr
	addr_QGraphicsWidget_closeEvent_default uintptr
	addr_QGraphicsWidget_focusInEvent_default uintptr
	addr_QGraphicsWidget_focusNextPrevChild_default uintptr
	addr_QGraphicsWidget_focusOutEvent_default uintptr
	addr_QGraphicsWidget_hideEvent_default uintptr
	addr_QGraphicsWidget_moveEvent_default uintptr
	addr_QGraphicsWidget_polishEvent_default uintptr
	addr_QGraphicsWidget_resizeEvent_default uintptr
	addr_QGraphicsWidget_showEvent_default uintptr
	addr_QGraphicsWidget_hoverMoveEvent_default uintptr
	addr_QGraphicsWidget_hoverLeaveEvent_default uintptr
	addr_QGraphicsWidget_grabMouseEvent_default uintptr
	addr_QGraphicsWidget_ungrabMouseEvent_default uintptr
	addr_QGraphicsWidget_grabKeyboardEvent_default uintptr
	addr_QGraphicsWidget_ungrabKeyboardEvent_default uintptr
	addr_QGraphicsWidget_eventFilter_default uintptr
	addr_QGraphicsWidget_timerEvent_default uintptr
	addr_QGraphicsWidget_childEvent_default uintptr
	addr_QGraphicsWidget_customEvent_default uintptr
	addr_QGraphicsWidget_connectNotify_default uintptr
	addr_QGraphicsWidget_disconnectNotify_default uintptr
	addr_QGraphicsWidget_advance_default uintptr
	addr_QGraphicsWidget_contains_default uintptr
	addr_QGraphicsWidget_collidesWithItem_default uintptr
	addr_QGraphicsWidget_collidesWithPath_default uintptr
	addr_QGraphicsWidget_isObscuredBy_default uintptr
	addr_QGraphicsWidget_opaqueArea_default uintptr
	addr_QGraphicsWidget_sceneEventFilter_default uintptr
	addr_QGraphicsWidget_contextMenuEvent_default uintptr
	addr_QGraphicsWidget_dragEnterEvent_default uintptr
	addr_QGraphicsWidget_dragLeaveEvent_default uintptr
	addr_QGraphicsWidget_dragMoveEvent_default uintptr
	addr_QGraphicsWidget_dropEvent_default uintptr
	addr_QGraphicsWidget_hoverEnterEvent_default uintptr
	addr_QGraphicsWidget_keyPressEvent_default uintptr
	addr_QGraphicsWidget_keyReleaseEvent_default uintptr
	addr_QGraphicsWidget_mousePressEvent_default uintptr
	addr_QGraphicsWidget_mouseMoveEvent_default uintptr
	addr_QGraphicsWidget_mouseReleaseEvent_default uintptr
	addr_QGraphicsWidget_mouseDoubleClickEvent_default uintptr
	addr_QGraphicsWidget_wheelEvent_default uintptr
	addr_QGraphicsWidget_inputMethodEvent_default uintptr
	addr_QGraphicsWidget_inputMethodQuery_default uintptr
	addr_QGraphicsWidget_supportsExtension_default uintptr
	addr_QGraphicsWidget_setExtension_default uintptr
	addr_QGraphicsWidget_extension_default uintptr
	pg_QGraphicsWidget_connect_geometryChanged func(uintptr, int64)
	pg_QGraphicsWidget_connect_layoutChanged func(uintptr, int64)
	pg_QGraphicsWidget_connect_parentChanged func(uintptr, int64)
	pg_QGraphicsWidget_connect_opacityChanged func(uintptr, int64)
	pg_QGraphicsWidget_connect_visibleChanged func(uintptr, int64)
	pg_QGraphicsWidget_connect_enabledChanged func(uintptr, int64)
	pg_QGraphicsWidget_connect_xChanged func(uintptr, int64)
	pg_QGraphicsWidget_connect_yChanged func(uintptr, int64)
	pg_QGraphicsWidget_connect_zChanged func(uintptr, int64)
	pg_QGraphicsWidget_connect_rotationChanged func(uintptr, int64)
	pg_QGraphicsWidget_connect_scaleChanged func(uintptr, int64)
	pg_QGraphicsWidget_connect_childrenChanged func(uintptr, int64)
	pg_QGraphicsWidget_connect_widthChanged func(uintptr, int64)
	pg_QGraphicsWidget_connect_heightChanged func(uintptr, int64)
	pg_QGraphicsWidget_connect_destroyed func(uintptr, int64)
	pg_QGraphicsWidget_connect_destroyed2 func(uintptr, int64)
	pg_QGraphicsWidget_connect_objectNameChanged func(uintptr, int64)
	pg_QGraphicsWidget__disconnect func(int64)
	pg_QGraphicsWidget__set_override func(uintptr, int32)
	pg_QGraphicsWidget_delete func(uintptr)
)

func init() {
	symaddr(&addr_QGraphicsWidget_chain, "QGraphicsWidget_chain")
	symaddr(&addr_QGraphicsWidget_new, "QGraphicsWidget_new")
	symaddr(&addr_QGraphicsWidget_new2, "QGraphicsWidget_new2")
	symaddr(&addr_QGraphicsWidget_new3, "QGraphicsWidget_new3")
	symaddr(&addr_QGraphicsWidget_tr, "QGraphicsWidget_tr")
	symaddr(&addr_QGraphicsWidget_layout, "QGraphicsWidget_layout")
	symaddr(&addr_QGraphicsWidget_setLayout, "QGraphicsWidget_setLayout")
	symaddr(&addr_QGraphicsWidget_adjustSize, "QGraphicsWidget_adjustSize")
	symaddr(&addr_QGraphicsWidget_layoutDirection, "QGraphicsWidget_layoutDirection")
	symaddr(&addr_QGraphicsWidget_setLayoutDirection, "QGraphicsWidget_setLayoutDirection")
	symaddr(&addr_QGraphicsWidget_unsetLayoutDirection, "QGraphicsWidget_unsetLayoutDirection")
	symaddr(&addr_QGraphicsWidget_style, "QGraphicsWidget_style")
	symaddr(&addr_QGraphicsWidget_setStyle, "QGraphicsWidget_setStyle")
	symaddr(&addr_QGraphicsWidget_font, "QGraphicsWidget_font")
	symaddr(&addr_QGraphicsWidget_setFont, "QGraphicsWidget_setFont")
	symaddr(&addr_QGraphicsWidget_palette, "QGraphicsWidget_palette")
	symaddr(&addr_QGraphicsWidget_setPalette, "QGraphicsWidget_setPalette")
	symaddr(&addr_QGraphicsWidget_autoFillBackground, "QGraphicsWidget_autoFillBackground")
	symaddr(&addr_QGraphicsWidget_setAutoFillBackground, "QGraphicsWidget_setAutoFillBackground")
	symaddr(&addr_QGraphicsWidget_resize, "QGraphicsWidget_resize")
	reg(&pg_QGraphicsWidget_resize2, "QGraphicsWidget_resize2")
	symaddr(&addr_QGraphicsWidget_size, "QGraphicsWidget_size")
	symaddr(&addr_QGraphicsWidget_setGeometry, "QGraphicsWidget_setGeometry")
	reg(&pg_QGraphicsWidget_setGeometry2, "QGraphicsWidget_setGeometry2")
	symaddr(&addr_QGraphicsWidget_rect, "QGraphicsWidget_rect")
	reg(&pg_QGraphicsWidget_setContentsMargins, "QGraphicsWidget_setContentsMargins")
	symaddr(&addr_QGraphicsWidget_setContentsMargins2, "QGraphicsWidget_setContentsMargins2")
	symaddr(&addr_QGraphicsWidget_getContentsMargins, "QGraphicsWidget_getContentsMargins")
	reg(&pg_QGraphicsWidget_setWindowFrameMargins, "QGraphicsWidget_setWindowFrameMargins")
	symaddr(&addr_QGraphicsWidget_setWindowFrameMargins2, "QGraphicsWidget_setWindowFrameMargins2")
	symaddr(&addr_QGraphicsWidget_getWindowFrameMargins, "QGraphicsWidget_getWindowFrameMargins")
	symaddr(&addr_QGraphicsWidget_unsetWindowFrameMargins, "QGraphicsWidget_unsetWindowFrameMargins")
	symaddr(&addr_QGraphicsWidget_windowFrameGeometry, "QGraphicsWidget_windowFrameGeometry")
	symaddr(&addr_QGraphicsWidget_windowFrameRect, "QGraphicsWidget_windowFrameRect")
	symaddr(&addr_QGraphicsWidget_windowFlags, "QGraphicsWidget_windowFlags")
	symaddr(&addr_QGraphicsWidget_windowType, "QGraphicsWidget_windowType")
	symaddr(&addr_QGraphicsWidget_setWindowFlags, "QGraphicsWidget_setWindowFlags")
	symaddr(&addr_QGraphicsWidget_isActiveWindow, "QGraphicsWidget_isActiveWindow")
	symaddr(&addr_QGraphicsWidget_setWindowTitle, "QGraphicsWidget_setWindowTitle")
	symaddr(&addr_QGraphicsWidget_windowTitle, "QGraphicsWidget_windowTitle")
	symaddr(&addr_QGraphicsWidget_focusPolicy, "QGraphicsWidget_focusPolicy")
	symaddr(&addr_QGraphicsWidget_setFocusPolicy, "QGraphicsWidget_setFocusPolicy")
	symaddr(&addr_QGraphicsWidget_setTabOrder, "QGraphicsWidget_setTabOrder")
	symaddr(&addr_QGraphicsWidget_focusWidget, "QGraphicsWidget_focusWidget")
	symaddr(&addr_QGraphicsWidget_grabShortcut, "QGraphicsWidget_grabShortcut")
	symaddr(&addr_QGraphicsWidget_releaseShortcut, "QGraphicsWidget_releaseShortcut")
	symaddr(&addr_QGraphicsWidget_setShortcutEnabled, "QGraphicsWidget_setShortcutEnabled")
	symaddr(&addr_QGraphicsWidget_setShortcutAutoRepeat, "QGraphicsWidget_setShortcutAutoRepeat")
	symaddr(&addr_QGraphicsWidget_addAction, "QGraphicsWidget_addAction")
	symaddr(&addr_QGraphicsWidget_addActions, "QGraphicsWidget_addActions")
	symaddr(&addr_QGraphicsWidget_insertActions, "QGraphicsWidget_insertActions")
	symaddr(&addr_QGraphicsWidget_insertAction, "QGraphicsWidget_insertAction")
	symaddr(&addr_QGraphicsWidget_removeAction, "QGraphicsWidget_removeAction")
	symaddr(&addr_QGraphicsWidget_actions, "QGraphicsWidget_actions")
	symaddr(&addr_QGraphicsWidget_setAttribute, "QGraphicsWidget_setAttribute")
	symaddr(&addr_QGraphicsWidget_testAttribute, "QGraphicsWidget_testAttribute")
	symaddr(&addr_QGraphicsWidget_paintWindowFrame, "QGraphicsWidget_paintWindowFrame")
	symaddr(&addr_QGraphicsWidget_geometryChanged, "QGraphicsWidget_geometryChanged")
	symaddr(&addr_QGraphicsWidget_layoutChanged, "QGraphicsWidget_layoutChanged")
	symaddr(&addr_QGraphicsWidget_close, "QGraphicsWidget_close")
	symaddr(&addr_QGraphicsWidget_updateMicroFocus, "QGraphicsWidget_updateMicroFocus")
	symaddr(&addr_QGraphicsWidget_connect, "QGraphicsWidget_connect")
	symaddr(&addr_QGraphicsWidget_connect2, "QGraphicsWidget_connect2")
	symaddr(&addr_QGraphicsWidget_sender, "QGraphicsWidget_sender")
	symaddr(&addr_QGraphicsWidget_senderSignalIndex, "QGraphicsWidget_senderSignalIndex")
	symaddr(&addr_QGraphicsWidget_receivers, "QGraphicsWidget_receivers")
	symaddr(&addr_QGraphicsWidget_isSignalConnected, "QGraphicsWidget_isSignalConnected")
	symaddr(&addr_QGraphicsWidget_addToIndex, "QGraphicsWidget_addToIndex")
	symaddr(&addr_QGraphicsWidget_removeFromIndex, "QGraphicsWidget_removeFromIndex")
	symaddr(&addr_QGraphicsWidget_prepareGeometryChange, "QGraphicsWidget_prepareGeometryChange")
	symaddr(&addr_QGraphicsWidget_setSizePolicy, "QGraphicsWidget_setSizePolicy")
	symaddr(&addr_QGraphicsWidget_setSizePolicy2, "QGraphicsWidget_setSizePolicy2")
	symaddr(&addr_QGraphicsWidget_sizePolicy, "QGraphicsWidget_sizePolicy")
	symaddr(&addr_QGraphicsWidget_setMinimumSize, "QGraphicsWidget_setMinimumSize")
	reg(&pg_QGraphicsWidget_setMinimumSize2, "QGraphicsWidget_setMinimumSize2")
	symaddr(&addr_QGraphicsWidget_minimumSize, "QGraphicsWidget_minimumSize")
	reg(&pg_QGraphicsWidget_setMinimumWidth, "QGraphicsWidget_setMinimumWidth")
	reg(&pg_QGraphicsWidget_minimumWidth, "QGraphicsWidget_minimumWidth")
	reg(&pg_QGraphicsWidget_setMinimumHeight, "QGraphicsWidget_setMinimumHeight")
	reg(&pg_QGraphicsWidget_minimumHeight, "QGraphicsWidget_minimumHeight")
	symaddr(&addr_QGraphicsWidget_setPreferredSize, "QGraphicsWidget_setPreferredSize")
	reg(&pg_QGraphicsWidget_setPreferredSize2, "QGraphicsWidget_setPreferredSize2")
	symaddr(&addr_QGraphicsWidget_preferredSize, "QGraphicsWidget_preferredSize")
	reg(&pg_QGraphicsWidget_setPreferredWidth, "QGraphicsWidget_setPreferredWidth")
	reg(&pg_QGraphicsWidget_preferredWidth, "QGraphicsWidget_preferredWidth")
	reg(&pg_QGraphicsWidget_setPreferredHeight, "QGraphicsWidget_setPreferredHeight")
	reg(&pg_QGraphicsWidget_preferredHeight, "QGraphicsWidget_preferredHeight")
	symaddr(&addr_QGraphicsWidget_setMaximumSize, "QGraphicsWidget_setMaximumSize")
	reg(&pg_QGraphicsWidget_setMaximumSize2, "QGraphicsWidget_setMaximumSize2")
	symaddr(&addr_QGraphicsWidget_maximumSize, "QGraphicsWidget_maximumSize")
	reg(&pg_QGraphicsWidget_setMaximumWidth, "QGraphicsWidget_setMaximumWidth")
	reg(&pg_QGraphicsWidget_maximumWidth, "QGraphicsWidget_maximumWidth")
	reg(&pg_QGraphicsWidget_setMaximumHeight, "QGraphicsWidget_setMaximumHeight")
	reg(&pg_QGraphicsWidget_maximumHeight, "QGraphicsWidget_maximumHeight")
	symaddr(&addr_QGraphicsWidget_geometry, "QGraphicsWidget_geometry")
	symaddr(&addr_QGraphicsWidget_contentsRect, "QGraphicsWidget_contentsRect")
	symaddr(&addr_QGraphicsWidget_effectiveSizeHint, "QGraphicsWidget_effectiveSizeHint")
	symaddr(&addr_QGraphicsWidget_isEmpty, "QGraphicsWidget_isEmpty")
	symaddr(&addr_QGraphicsWidget_parentLayoutItem, "QGraphicsWidget_parentLayoutItem")
	symaddr(&addr_QGraphicsWidget_setParentLayoutItem, "QGraphicsWidget_setParentLayoutItem")
	symaddr(&addr_QGraphicsWidget_isLayout, "QGraphicsWidget_isLayout")
	symaddr(&addr_QGraphicsWidget_graphicsItem, "QGraphicsWidget_graphicsItem")
	symaddr(&addr_QGraphicsWidget_ownedByLayout, "QGraphicsWidget_ownedByLayout")
	symaddr(&addr_QGraphicsWidget_setGraphicsItem, "QGraphicsWidget_setGraphicsItem")
	symaddr(&addr_QGraphicsWidget_setOwnedByLayout, "QGraphicsWidget_setOwnedByLayout")
	symaddr(&addr_QGraphicsWidget_tr2, "QGraphicsWidget_tr2")
	symaddr(&addr_QGraphicsWidget_tr3, "QGraphicsWidget_tr3")
	symaddr(&addr_QGraphicsWidget_grabShortcut2, "QGraphicsWidget_grabShortcut2")
	symaddr(&addr_QGraphicsWidget_setShortcutEnabled2, "QGraphicsWidget_setShortcutEnabled2")
	symaddr(&addr_QGraphicsWidget_setShortcutAutoRepeat2, "QGraphicsWidget_setShortcutAutoRepeat2")
	symaddr(&addr_QGraphicsWidget_setAttribute2, "QGraphicsWidget_setAttribute2")
	symaddr(&addr_QGraphicsWidget_connect3, "QGraphicsWidget_connect3")
	symaddr(&addr_QGraphicsWidget_connect4, "QGraphicsWidget_connect4")
	symaddr(&addr_QGraphicsWidget_setSizePolicy3, "QGraphicsWidget_setSizePolicy3")
	symaddr(&addr_QGraphicsWidget_effectiveSizeHint2, "QGraphicsWidget_effectiveSizeHint2")
	symaddr(&addr_QGraphicsWidget_asQGraphicsLayoutItem, "QGraphicsWidget_asQGraphicsLayoutItem")
	symaddr(&addr_QGraphicsWidget_setGeometry_default, "QGraphicsWidget_setGeometry_default")
	reg(&pg_QGraphicsWidget_setGeometry_default2, "QGraphicsWidget_setGeometry_default2")
	symaddr(&addr_QGraphicsWidget_getContentsMargins_default, "QGraphicsWidget_getContentsMargins_default")
	symaddr(&addr_QGraphicsWidget_paintWindowFrame_default, "QGraphicsWidget_paintWindowFrame_default")
	symaddr(&addr_QGraphicsWidget_isEmpty_default, "QGraphicsWidget_isEmpty_default")
	symaddr(&addr_QGraphicsWidget_metaObject_default, "QGraphicsWidget_metaObject_default")
	symaddr(&addr_QGraphicsWidget_type_default, "QGraphicsWidget_type_default")
	symaddr(&addr_QGraphicsWidget_paint_default, "QGraphicsWidget_paint_default")
	symaddr(&addr_QGraphicsWidget_boundingRect_default, "QGraphicsWidget_boundingRect_default")
	symaddr(&addr_QGraphicsWidget_shape_default, "QGraphicsWidget_shape_default")
	symaddr(&addr_QGraphicsWidget_initStyleOption_default, "QGraphicsWidget_initStyleOption_default")
	symaddr(&addr_QGraphicsWidget_sizeHint_default, "QGraphicsWidget_sizeHint_default")
	symaddr(&addr_QGraphicsWidget_updateGeometry_default, "QGraphicsWidget_updateGeometry_default")
	symaddr(&addr_QGraphicsWidget_itemChange_default, "QGraphicsWidget_itemChange_default")
	symaddr(&addr_QGraphicsWidget_propertyChange_default, "QGraphicsWidget_propertyChange_default")
	symaddr(&addr_QGraphicsWidget_sceneEvent_default, "QGraphicsWidget_sceneEvent_default")
	symaddr(&addr_QGraphicsWidget_windowFrameEvent_default, "QGraphicsWidget_windowFrameEvent_default")
	symaddr(&addr_QGraphicsWidget_windowFrameSectionAt_default, "QGraphicsWidget_windowFrameSectionAt_default")
	symaddr(&addr_QGraphicsWidget_event_default, "QGraphicsWidget_event_default")
	symaddr(&addr_QGraphicsWidget_changeEvent_default, "QGraphicsWidget_changeEvent_default")
	symaddr(&addr_QGraphicsWidget_closeEvent_default, "QGraphicsWidget_closeEvent_default")
	symaddr(&addr_QGraphicsWidget_focusInEvent_default, "QGraphicsWidget_focusInEvent_default")
	symaddr(&addr_QGraphicsWidget_focusNextPrevChild_default, "QGraphicsWidget_focusNextPrevChild_default")
	symaddr(&addr_QGraphicsWidget_focusOutEvent_default, "QGraphicsWidget_focusOutEvent_default")
	symaddr(&addr_QGraphicsWidget_hideEvent_default, "QGraphicsWidget_hideEvent_default")
	symaddr(&addr_QGraphicsWidget_moveEvent_default, "QGraphicsWidget_moveEvent_default")
	symaddr(&addr_QGraphicsWidget_polishEvent_default, "QGraphicsWidget_polishEvent_default")
	symaddr(&addr_QGraphicsWidget_resizeEvent_default, "QGraphicsWidget_resizeEvent_default")
	symaddr(&addr_QGraphicsWidget_showEvent_default, "QGraphicsWidget_showEvent_default")
	symaddr(&addr_QGraphicsWidget_hoverMoveEvent_default, "QGraphicsWidget_hoverMoveEvent_default")
	symaddr(&addr_QGraphicsWidget_hoverLeaveEvent_default, "QGraphicsWidget_hoverLeaveEvent_default")
	symaddr(&addr_QGraphicsWidget_grabMouseEvent_default, "QGraphicsWidget_grabMouseEvent_default")
	symaddr(&addr_QGraphicsWidget_ungrabMouseEvent_default, "QGraphicsWidget_ungrabMouseEvent_default")
	symaddr(&addr_QGraphicsWidget_grabKeyboardEvent_default, "QGraphicsWidget_grabKeyboardEvent_default")
	symaddr(&addr_QGraphicsWidget_ungrabKeyboardEvent_default, "QGraphicsWidget_ungrabKeyboardEvent_default")
	symaddr(&addr_QGraphicsWidget_eventFilter_default, "QGraphicsWidget_eventFilter_default")
	symaddr(&addr_QGraphicsWidget_timerEvent_default, "QGraphicsWidget_timerEvent_default")
	symaddr(&addr_QGraphicsWidget_childEvent_default, "QGraphicsWidget_childEvent_default")
	symaddr(&addr_QGraphicsWidget_customEvent_default, "QGraphicsWidget_customEvent_default")
	symaddr(&addr_QGraphicsWidget_connectNotify_default, "QGraphicsWidget_connectNotify_default")
	symaddr(&addr_QGraphicsWidget_disconnectNotify_default, "QGraphicsWidget_disconnectNotify_default")
	symaddr(&addr_QGraphicsWidget_advance_default, "QGraphicsWidget_advance_default")
	symaddr(&addr_QGraphicsWidget_contains_default, "QGraphicsWidget_contains_default")
	symaddr(&addr_QGraphicsWidget_collidesWithItem_default, "QGraphicsWidget_collidesWithItem_default")
	symaddr(&addr_QGraphicsWidget_collidesWithPath_default, "QGraphicsWidget_collidesWithPath_default")
	symaddr(&addr_QGraphicsWidget_isObscuredBy_default, "QGraphicsWidget_isObscuredBy_default")
	symaddr(&addr_QGraphicsWidget_opaqueArea_default, "QGraphicsWidget_opaqueArea_default")
	symaddr(&addr_QGraphicsWidget_sceneEventFilter_default, "QGraphicsWidget_sceneEventFilter_default")
	symaddr(&addr_QGraphicsWidget_contextMenuEvent_default, "QGraphicsWidget_contextMenuEvent_default")
	symaddr(&addr_QGraphicsWidget_dragEnterEvent_default, "QGraphicsWidget_dragEnterEvent_default")
	symaddr(&addr_QGraphicsWidget_dragLeaveEvent_default, "QGraphicsWidget_dragLeaveEvent_default")
	symaddr(&addr_QGraphicsWidget_dragMoveEvent_default, "QGraphicsWidget_dragMoveEvent_default")
	symaddr(&addr_QGraphicsWidget_dropEvent_default, "QGraphicsWidget_dropEvent_default")
	symaddr(&addr_QGraphicsWidget_hoverEnterEvent_default, "QGraphicsWidget_hoverEnterEvent_default")
	symaddr(&addr_QGraphicsWidget_keyPressEvent_default, "QGraphicsWidget_keyPressEvent_default")
	symaddr(&addr_QGraphicsWidget_keyReleaseEvent_default, "QGraphicsWidget_keyReleaseEvent_default")
	symaddr(&addr_QGraphicsWidget_mousePressEvent_default, "QGraphicsWidget_mousePressEvent_default")
	symaddr(&addr_QGraphicsWidget_mouseMoveEvent_default, "QGraphicsWidget_mouseMoveEvent_default")
	symaddr(&addr_QGraphicsWidget_mouseReleaseEvent_default, "QGraphicsWidget_mouseReleaseEvent_default")
	symaddr(&addr_QGraphicsWidget_mouseDoubleClickEvent_default, "QGraphicsWidget_mouseDoubleClickEvent_default")
	symaddr(&addr_QGraphicsWidget_wheelEvent_default, "QGraphicsWidget_wheelEvent_default")
	symaddr(&addr_QGraphicsWidget_inputMethodEvent_default, "QGraphicsWidget_inputMethodEvent_default")
	symaddr(&addr_QGraphicsWidget_inputMethodQuery_default, "QGraphicsWidget_inputMethodQuery_default")
	symaddr(&addr_QGraphicsWidget_supportsExtension_default, "QGraphicsWidget_supportsExtension_default")
	symaddr(&addr_QGraphicsWidget_setExtension_default, "QGraphicsWidget_setExtension_default")
	symaddr(&addr_QGraphicsWidget_extension_default, "QGraphicsWidget_extension_default")
	reg(&pg_QGraphicsWidget_connect_geometryChanged, "QGraphicsWidget_connect_geometryChanged")
	reg(&pg_QGraphicsWidget_connect_layoutChanged, "QGraphicsWidget_connect_layoutChanged")
	reg(&pg_QGraphicsWidget_connect_parentChanged, "QGraphicsWidget_connect_parentChanged")
	reg(&pg_QGraphicsWidget_connect_opacityChanged, "QGraphicsWidget_connect_opacityChanged")
	reg(&pg_QGraphicsWidget_connect_visibleChanged, "QGraphicsWidget_connect_visibleChanged")
	reg(&pg_QGraphicsWidget_connect_enabledChanged, "QGraphicsWidget_connect_enabledChanged")
	reg(&pg_QGraphicsWidget_connect_xChanged, "QGraphicsWidget_connect_xChanged")
	reg(&pg_QGraphicsWidget_connect_yChanged, "QGraphicsWidget_connect_yChanged")
	reg(&pg_QGraphicsWidget_connect_zChanged, "QGraphicsWidget_connect_zChanged")
	reg(&pg_QGraphicsWidget_connect_rotationChanged, "QGraphicsWidget_connect_rotationChanged")
	reg(&pg_QGraphicsWidget_connect_scaleChanged, "QGraphicsWidget_connect_scaleChanged")
	reg(&pg_QGraphicsWidget_connect_childrenChanged, "QGraphicsWidget_connect_childrenChanged")
	reg(&pg_QGraphicsWidget_connect_widthChanged, "QGraphicsWidget_connect_widthChanged")
	reg(&pg_QGraphicsWidget_connect_heightChanged, "QGraphicsWidget_connect_heightChanged")
	reg(&pg_QGraphicsWidget_connect_destroyed, "QGraphicsWidget_connect_destroyed")
	reg(&pg_QGraphicsWidget_connect_destroyed2, "QGraphicsWidget_connect_destroyed2")
	reg(&pg_QGraphicsWidget_connect_objectNameChanged, "QGraphicsWidget_connect_objectNameChanged")
	reg(&pg_QGraphicsWidget__disconnect, "QGraphicsWidget__disconnect")
	reg(&pg_QGraphicsWidget__set_override, "QGraphicsWidget__set_override")
	reg(&pg_QGraphicsWidget_delete, "QGraphicsWidget_delete")
}

type QGraphicsWidget struct {
	h uintptr
	*QGraphicsObject
}

func bindQGraphicsWidget(h uintptr) *QGraphicsWidget {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QGraphicsWidget_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QGraphicsWidget{h: h, QGraphicsObject: &QGraphicsObject{h: __c[0], QObject: &QObject{h: __c[1]}}}
}

func newownedQGraphicsWidget(h uintptr) *QGraphicsWidget { return bindQGraphicsWidget(h) }

func (this *QGraphicsWidget) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QGraphicsWidget) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQGraphicsWidget(h unsafe.Pointer) *QGraphicsWidget { return bindQGraphicsWidget(uintptr(h)) }
func UnsafeNewQGraphicsWidgetOwned(h unsafe.Pointer) *QGraphicsWidget { return newownedQGraphicsWidget(uintptr(h)) }

func (this *QGraphicsWidget) UnsafeRelease() {
	this.h = 0
	if this.QGraphicsObject != nil {
		this.QGraphicsObject.UnsafeRelease()
	}
}

func NewQGraphicsWidget() *QGraphicsWidget {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QGraphicsWidget_new, uintptr(__goid))
	__o := newownedQGraphicsWidget(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGraphicsWidget_virtIdx, QGraphicsWidget_enable)
	return __o
}

func NewQGraphicsWidget2(parent *QGraphicsItem) *QGraphicsWidget {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QGraphicsWidget_new2, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQGraphicsWidget(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGraphicsWidget_virtIdx, QGraphicsWidget_enable)
	return __o
}

func NewQGraphicsWidget3(parent *QGraphicsItem, wFlags WindowType) *QGraphicsWidget {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QGraphicsWidget_new3, uintptr(__goid), uintptr(parent.UnsafePointer()), uintptr(wFlags))
	__o := newownedQGraphicsWidget(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGraphicsWidget_virtIdx, QGraphicsWidget_enable)
	return __o
}

func QGraphicsWidget_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGraphicsWidget_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QGraphicsWidget) Layout() *QGraphicsLayout {
	__r, _, _ := purego.SyscallN(addr_QGraphicsWidget_layout, this.h)
	runtime.KeepAlive(this)
	return bindQGraphicsLayout(__r)
}

func (this *QGraphicsWidget) SetLayout(layout *QGraphicsLayout) {
	purego.SyscallN(addr_QGraphicsWidget_setLayout, this.h, uintptr(layout.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(layout)
}

func (this *QGraphicsWidget) AdjustSize() {
	purego.SyscallN(addr_QGraphicsWidget_adjustSize, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsWidget) LayoutDirection() LayoutDirection {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsWidget_layoutDirection, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return LayoutDirection(__r)
}

func (this *QGraphicsWidget) SetLayoutDirection(direction LayoutDirection) {
	purego.SyscallN(addr_QGraphicsWidget_setLayoutDirection, this.h, uintptr(direction))
	runtime.KeepAlive(this)
}

func (this *QGraphicsWidget) UnsetLayoutDirection() {
	purego.SyscallN(addr_QGraphicsWidget_unsetLayoutDirection, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsWidget) Style() *QStyle {
	__r, _, _ := purego.SyscallN(addr_QGraphicsWidget_style, this.h)
	runtime.KeepAlive(this)
	return bindQStyle(__r)
}

func (this *QGraphicsWidget) SetStyle(style *QStyle) {
	purego.SyscallN(addr_QGraphicsWidget_setStyle, this.h, uintptr(style.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(style)
}

func (this *QGraphicsWidget) Font() *QFont {
	__r, _, _ := purego.SyscallN(addr_QGraphicsWidget_font, this.h)
	runtime.KeepAlive(this)
	return newownedQFont(__r)
}

func (this *QGraphicsWidget) SetFont(font *QFont) {
	purego.SyscallN(addr_QGraphicsWidget_setFont, this.h, uintptr(font.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(font)
}

func (this *QGraphicsWidget) Palette() *QPalette {
	__r, _, _ := purego.SyscallN(addr_QGraphicsWidget_palette, this.h)
	runtime.KeepAlive(this)
	return newownedQPalette(__r)
}

func (this *QGraphicsWidget) SetPalette(palette *QPalette) {
	purego.SyscallN(addr_QGraphicsWidget_setPalette, this.h, uintptr(palette.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(palette)
}

func (this *QGraphicsWidget) AutoFillBackground() bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsWidget_autoFillBackground, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGraphicsWidget) SetAutoFillBackground(enabled bool) {
	purego.SyscallN(addr_QGraphicsWidget_setAutoFillBackground, this.h, uintptr(b2i(enabled)))
	runtime.KeepAlive(this)
}

func (this *QGraphicsWidget) Resize(size *QSizeF) {
	purego.SyscallN(addr_QGraphicsWidget_resize, this.h, uintptr(size.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(size)
}

func (this *QGraphicsWidget) Resize2(w float64, h float64) {
	pg_QGraphicsWidget_resize2(this.h, w, h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsWidget) Size() *QSizeF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsWidget_size, this.h)
	runtime.KeepAlive(this)
	return newownedQSizeF(__r)
}

func (this *QGraphicsWidget) SetGeometry(rect *QRectF) {
	purego.SyscallN(addr_QGraphicsWidget_setGeometry, this.h, uintptr(rect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
}

func (this *QGraphicsWidget) SetGeometry2(x float64, y float64, w float64, h float64) {
	pg_QGraphicsWidget_setGeometry2(this.h, x, y, w, h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsWidget) Rect() *QRectF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsWidget_rect, this.h)
	runtime.KeepAlive(this)
	return newownedQRectF(__r)
}

func (this *QGraphicsWidget) SetContentsMargins(left float64, top float64, right float64, bottom float64) {
	pg_QGraphicsWidget_setContentsMargins(this.h, left, top, right, bottom)
	runtime.KeepAlive(this)
}

func (this *QGraphicsWidget) SetContentsMarginsWithMargins(margins *QMarginsF) {
	purego.SyscallN(addr_QGraphicsWidget_setContentsMargins2, this.h, uintptr(margins.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(margins)
}

func (this *QGraphicsWidget) GetContentsMargins() (float64, float64, float64, float64) {
	var left_out float64
	var top_out float64
	var right_out float64
	var bottom_out float64
	purego.SyscallN(addr_QGraphicsWidget_getContentsMargins, this.h, uintptr(unsafe.Pointer(&left_out)), uintptr(unsafe.Pointer(&top_out)), uintptr(unsafe.Pointer(&right_out)), uintptr(unsafe.Pointer(&bottom_out)))
	runtime.KeepAlive(this)
	return left_out, top_out, right_out, bottom_out
}

func (this *QGraphicsWidget) SetWindowFrameMargins(left float64, top float64, right float64, bottom float64) {
	pg_QGraphicsWidget_setWindowFrameMargins(this.h, left, top, right, bottom)
	runtime.KeepAlive(this)
}

func (this *QGraphicsWidget) SetWindowFrameMarginsWithMargins(margins *QMarginsF) {
	purego.SyscallN(addr_QGraphicsWidget_setWindowFrameMargins2, this.h, uintptr(margins.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(margins)
}

func (this *QGraphicsWidget) GetWindowFrameMargins() (float64, float64, float64, float64) {
	var left_out float64
	var top_out float64
	var right_out float64
	var bottom_out float64
	purego.SyscallN(addr_QGraphicsWidget_getWindowFrameMargins, this.h, uintptr(unsafe.Pointer(&left_out)), uintptr(unsafe.Pointer(&top_out)), uintptr(unsafe.Pointer(&right_out)), uintptr(unsafe.Pointer(&bottom_out)))
	runtime.KeepAlive(this)
	return left_out, top_out, right_out, bottom_out
}

func (this *QGraphicsWidget) UnsetWindowFrameMargins() {
	purego.SyscallN(addr_QGraphicsWidget_unsetWindowFrameMargins, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsWidget) WindowFrameGeometry() *QRectF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsWidget_windowFrameGeometry, this.h)
	runtime.KeepAlive(this)
	return newownedQRectF(__r)
}

func (this *QGraphicsWidget) WindowFrameRect() *QRectF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsWidget_windowFrameRect, this.h)
	runtime.KeepAlive(this)
	return newownedQRectF(__r)
}

func (this *QGraphicsWidget) WindowFlags() WindowType {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsWidget_windowFlags, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return WindowType(__r)
}

func (this *QGraphicsWidget) WindowType() WindowType {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsWidget_windowType, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return WindowType(__r)
}

func (this *QGraphicsWidget) SetWindowFlags(wFlags WindowType) {
	purego.SyscallN(addr_QGraphicsWidget_setWindowFlags, this.h, uintptr(wFlags))
	runtime.KeepAlive(this)
}

func (this *QGraphicsWidget) IsActiveWindow() bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsWidget_isActiveWindow, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGraphicsWidget) SetWindowTitle(title string) {
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	purego.SyscallN(addr_QGraphicsWidget_setWindowTitle, this.h, title_p, uintptr(len(title)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(title)
}

func (this *QGraphicsWidget) WindowTitle() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGraphicsWidget_windowTitle, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QGraphicsWidget) FocusPolicy() FocusPolicy {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsWidget_focusPolicy, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return FocusPolicy(__r)
}

func (this *QGraphicsWidget) SetFocusPolicy(policy FocusPolicy) {
	purego.SyscallN(addr_QGraphicsWidget_setFocusPolicy, this.h, uintptr(policy))
	runtime.KeepAlive(this)
}

func QGraphicsWidget_SetTabOrder(first *QGraphicsWidget, second *QGraphicsWidget) {
	purego.SyscallN(addr_QGraphicsWidget_setTabOrder, uintptr(first.UnsafePointer()), uintptr(second.UnsafePointer()))
	runtime.KeepAlive(first)
	runtime.KeepAlive(second)
}

func (this *QGraphicsWidget) FocusWidget() *QGraphicsWidget {
	__r, _, _ := purego.SyscallN(addr_QGraphicsWidget_focusWidget, this.h)
	runtime.KeepAlive(this)
	return bindQGraphicsWidget(__r)
}

func (this *QGraphicsWidget) GrabShortcut(sequence *QKeySequence) int {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsWidget_grabShortcut, this.h, uintptr(sequence.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sequence)
	return int(__r)
}

func (this *QGraphicsWidget) ReleaseShortcut(id int) {
	purego.SyscallN(addr_QGraphicsWidget_releaseShortcut, this.h, uintptr(id))
	runtime.KeepAlive(this)
}

func (this *QGraphicsWidget) SetShortcutEnabled(id int) {
	purego.SyscallN(addr_QGraphicsWidget_setShortcutEnabled, this.h, uintptr(id))
	runtime.KeepAlive(this)
}

func (this *QGraphicsWidget) SetShortcutAutoRepeat(id int) {
	purego.SyscallN(addr_QGraphicsWidget_setShortcutAutoRepeat, this.h, uintptr(id))
	runtime.KeepAlive(this)
}

func (this *QGraphicsWidget) AddAction(action *QAction) {
	purego.SyscallN(addr_QGraphicsWidget_addAction, this.h, uintptr(action.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(action)
}

func (this *QGraphicsWidget) AddActions(actions []*QAction) {
	actions_arr := make([]uintptr, len(actions))
	for __i, __e := range actions {
		actions_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var actions_p uintptr
	if len(actions_arr) > 0 {
		actions_p = uintptr(unsafe.Pointer(&actions_arr[0]))
	}
	purego.SyscallN(addr_QGraphicsWidget_addActions, this.h, actions_p, uintptr(len(actions)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(actions)
	runtime.KeepAlive(actions_arr)
}

func (this *QGraphicsWidget) InsertActions(before *QAction, actions []*QAction) {
	actions_arr := make([]uintptr, len(actions))
	for __i, __e := range actions {
		actions_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var actions_p uintptr
	if len(actions_arr) > 0 {
		actions_p = uintptr(unsafe.Pointer(&actions_arr[0]))
	}
	purego.SyscallN(addr_QGraphicsWidget_insertActions, this.h, uintptr(before.UnsafePointer()), actions_p, uintptr(len(actions)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(before)
	runtime.KeepAlive(actions)
	runtime.KeepAlive(actions_arr)
}

func (this *QGraphicsWidget) InsertAction(before *QAction, action *QAction) {
	purego.SyscallN(addr_QGraphicsWidget_insertAction, this.h, uintptr(before.UnsafePointer()), uintptr(action.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(before)
	runtime.KeepAlive(action)
}

func (this *QGraphicsWidget) RemoveAction(action *QAction) {
	purego.SyscallN(addr_QGraphicsWidget_removeAction, this.h, uintptr(action.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(action)
}

func (this *QGraphicsWidget) Actions() []*QAction {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGraphicsWidget_actions, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QAction, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = bindQAction(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QGraphicsWidget) SetAttribute(attribute WidgetAttribute) {
	purego.SyscallN(addr_QGraphicsWidget_setAttribute, this.h, uintptr(attribute))
	runtime.KeepAlive(this)
}

func (this *QGraphicsWidget) TestAttribute(attribute WidgetAttribute) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsWidget_testAttribute, this.h, uintptr(attribute))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGraphicsWidget) PaintWindowFrame(painter *QPainter, option *QStyleOptionGraphicsItem, widget *QWidget) {
	purego.SyscallN(addr_QGraphicsWidget_paintWindowFrame, this.h, uintptr(painter.UnsafePointer()), uintptr(option.UnsafePointer()), uintptr(widget.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
	runtime.KeepAlive(option)
	runtime.KeepAlive(widget)
}

func (this *QGraphicsWidget) GeometryChanged() {
	purego.SyscallN(addr_QGraphicsWidget_geometryChanged, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsWidget) LayoutChanged() {
	purego.SyscallN(addr_QGraphicsWidget_layoutChanged, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsWidget) Close() bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsWidget_close, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGraphicsWidget) UpdateMicroFocus() {
	purego.SyscallN(addr_QGraphicsWidget_updateMicroFocus, this.h)
	runtime.KeepAlive(this)
}

func QGraphicsWidget_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QGraphicsWidget_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QGraphicsWidget_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QGraphicsWidget_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QGraphicsWidget) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QGraphicsWidget_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QGraphicsWidget) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsWidget_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QGraphicsWidget) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QGraphicsWidget_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QGraphicsWidget) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsWidget_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func (this *QGraphicsWidget) AddToIndex() {
	purego.SyscallN(addr_QGraphicsWidget_addToIndex, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsWidget) RemoveFromIndex() {
	purego.SyscallN(addr_QGraphicsWidget_removeFromIndex, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsWidget) PrepareGeometryChange() {
	purego.SyscallN(addr_QGraphicsWidget_prepareGeometryChange, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsWidget) SetSizePolicy(policy *QSizePolicy) {
	purego.SyscallN(addr_QGraphicsWidget_setSizePolicy, this.h, uintptr(policy.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(policy)
}

func (this *QGraphicsWidget) SetSizePolicy2(hPolicy QSizePolicy__Policy, vPolicy QSizePolicy__Policy) {
	purego.SyscallN(addr_QGraphicsWidget_setSizePolicy2, this.h, uintptr(hPolicy), uintptr(vPolicy))
	runtime.KeepAlive(this)
}

func (this *QGraphicsWidget) SizePolicy() *QSizePolicy {
	__r, _, _ := purego.SyscallN(addr_QGraphicsWidget_sizePolicy, this.h)
	runtime.KeepAlive(this)
	return newownedQSizePolicy(__r)
}

func (this *QGraphicsWidget) SetMinimumSize(size *QSizeF) {
	purego.SyscallN(addr_QGraphicsWidget_setMinimumSize, this.h, uintptr(size.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(size)
}

func (this *QGraphicsWidget) SetMinimumSize2(w float64, h float64) {
	pg_QGraphicsWidget_setMinimumSize2(this.h, w, h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsWidget) MinimumSize() *QSizeF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsWidget_minimumSize, this.h)
	runtime.KeepAlive(this)
	return newownedQSizeF(__r)
}

func (this *QGraphicsWidget) SetMinimumWidth(width float64) {
	pg_QGraphicsWidget_setMinimumWidth(this.h, width)
	runtime.KeepAlive(this)
}

func (this *QGraphicsWidget) MinimumWidth() float64 {
	__r := pg_QGraphicsWidget_minimumWidth(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QGraphicsWidget) SetMinimumHeight(height float64) {
	pg_QGraphicsWidget_setMinimumHeight(this.h, height)
	runtime.KeepAlive(this)
}

func (this *QGraphicsWidget) MinimumHeight() float64 {
	__r := pg_QGraphicsWidget_minimumHeight(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QGraphicsWidget) SetPreferredSize(size *QSizeF) {
	purego.SyscallN(addr_QGraphicsWidget_setPreferredSize, this.h, uintptr(size.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(size)
}

func (this *QGraphicsWidget) SetPreferredSize2(w float64, h float64) {
	pg_QGraphicsWidget_setPreferredSize2(this.h, w, h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsWidget) PreferredSize() *QSizeF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsWidget_preferredSize, this.h)
	runtime.KeepAlive(this)
	return newownedQSizeF(__r)
}

func (this *QGraphicsWidget) SetPreferredWidth(width float64) {
	pg_QGraphicsWidget_setPreferredWidth(this.h, width)
	runtime.KeepAlive(this)
}

func (this *QGraphicsWidget) PreferredWidth() float64 {
	__r := pg_QGraphicsWidget_preferredWidth(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QGraphicsWidget) SetPreferredHeight(height float64) {
	pg_QGraphicsWidget_setPreferredHeight(this.h, height)
	runtime.KeepAlive(this)
}

func (this *QGraphicsWidget) PreferredHeight() float64 {
	__r := pg_QGraphicsWidget_preferredHeight(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QGraphicsWidget) SetMaximumSize(size *QSizeF) {
	purego.SyscallN(addr_QGraphicsWidget_setMaximumSize, this.h, uintptr(size.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(size)
}

func (this *QGraphicsWidget) SetMaximumSize2(w float64, h float64) {
	pg_QGraphicsWidget_setMaximumSize2(this.h, w, h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsWidget) MaximumSize() *QSizeF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsWidget_maximumSize, this.h)
	runtime.KeepAlive(this)
	return newownedQSizeF(__r)
}

func (this *QGraphicsWidget) SetMaximumWidth(width float64) {
	pg_QGraphicsWidget_setMaximumWidth(this.h, width)
	runtime.KeepAlive(this)
}

func (this *QGraphicsWidget) MaximumWidth() float64 {
	__r := pg_QGraphicsWidget_maximumWidth(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QGraphicsWidget) SetMaximumHeight(height float64) {
	pg_QGraphicsWidget_setMaximumHeight(this.h, height)
	runtime.KeepAlive(this)
}

func (this *QGraphicsWidget) MaximumHeight() float64 {
	__r := pg_QGraphicsWidget_maximumHeight(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QGraphicsWidget) Geometry() *QRectF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsWidget_geometry, this.h)
	runtime.KeepAlive(this)
	return newownedQRectF(__r)
}

func (this *QGraphicsWidget) ContentsRect() *QRectF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsWidget_contentsRect, this.h)
	runtime.KeepAlive(this)
	return newownedQRectF(__r)
}

func (this *QGraphicsWidget) EffectiveSizeHint(which SizeHint) *QSizeF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsWidget_effectiveSizeHint, this.h, uintptr(which))
	runtime.KeepAlive(this)
	return newownedQSizeF(__r)
}

func (this *QGraphicsWidget) IsEmpty() bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsWidget_isEmpty, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGraphicsWidget) ParentLayoutItem() *QGraphicsLayoutItem {
	__r, _, _ := purego.SyscallN(addr_QGraphicsWidget_parentLayoutItem, this.h)
	runtime.KeepAlive(this)
	return bindQGraphicsLayoutItem(__r)
}

func (this *QGraphicsWidget) SetParentLayoutItem(parent *QGraphicsLayoutItem) {
	purego.SyscallN(addr_QGraphicsWidget_setParentLayoutItem, this.h, uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
}

func (this *QGraphicsWidget) IsLayout() bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsWidget_isLayout, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGraphicsWidget) GraphicsItem() *QGraphicsItem {
	__r, _, _ := purego.SyscallN(addr_QGraphicsWidget_graphicsItem, this.h)
	runtime.KeepAlive(this)
	return bindQGraphicsItem(__r)
}

func (this *QGraphicsWidget) OwnedByLayout() bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsWidget_ownedByLayout, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGraphicsWidget) SetGraphicsItem(item *QGraphicsItem) {
	purego.SyscallN(addr_QGraphicsWidget_setGraphicsItem, this.h, uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QGraphicsWidget) SetOwnedByLayout(ownedByLayout bool) {
	purego.SyscallN(addr_QGraphicsWidget_setOwnedByLayout, this.h, uintptr(b2i(ownedByLayout)))
	runtime.KeepAlive(this)
}

func QGraphicsWidget_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGraphicsWidget_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QGraphicsWidget_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGraphicsWidget_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func (this *QGraphicsWidget) GrabShortcut2(sequence *QKeySequence, context ShortcutContext) int {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsWidget_grabShortcut2, this.h, uintptr(sequence.UnsafePointer()), uintptr(context))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sequence)
	return int(__r)
}

func (this *QGraphicsWidget) SetShortcutEnabled2(id int, enabled bool) {
	purego.SyscallN(addr_QGraphicsWidget_setShortcutEnabled2, this.h, uintptr(id), uintptr(b2i(enabled)))
	runtime.KeepAlive(this)
}

func (this *QGraphicsWidget) SetShortcutAutoRepeat2(id int, enabled bool) {
	purego.SyscallN(addr_QGraphicsWidget_setShortcutAutoRepeat2, this.h, uintptr(id), uintptr(b2i(enabled)))
	runtime.KeepAlive(this)
}

func (this *QGraphicsWidget) SetAttribute2(attribute WidgetAttribute, on bool) {
	purego.SyscallN(addr_QGraphicsWidget_setAttribute2, this.h, uintptr(attribute), uintptr(b2i(on)))
	runtime.KeepAlive(this)
}

func QGraphicsWidget_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QGraphicsWidget_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QGraphicsWidget_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QGraphicsWidget_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QGraphicsWidget) SetSizePolicy3(hPolicy QSizePolicy__Policy, vPolicy QSizePolicy__Policy, controlType QSizePolicy__ControlType) {
	purego.SyscallN(addr_QGraphicsWidget_setSizePolicy3, this.h, uintptr(hPolicy), uintptr(vPolicy), uintptr(controlType))
	runtime.KeepAlive(this)
}

func (this *QGraphicsWidget) EffectiveSizeHint2(which SizeHint, constraint *QSizeF) *QSizeF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsWidget_effectiveSizeHint2, this.h, uintptr(which), uintptr(constraint.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(constraint)
	return newownedQSizeF(__r)
}

func (this *QGraphicsWidget) AsQGraphicsLayoutItem() *QGraphicsLayoutItem {
	__r, _, _ := purego.SyscallN(addr_QGraphicsWidget_asQGraphicsLayoutItem, this.h)
	runtime.KeepAlive(this)
	return bindQGraphicsLayoutItem(__r)
}

func (this *QGraphicsWidget) SetGeometryDefault(rect *QRectF) {
	purego.SyscallN(addr_QGraphicsWidget_setGeometry_default, this.h, uintptr(rect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
}

func (this *QGraphicsWidget) SetGeometry2Default(x float64, y float64, w float64, h float64) {
	pg_QGraphicsWidget_setGeometry_default2(this.h, x, y, w, h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsWidget) GetContentsMarginsDefault() (float64, float64, float64, float64) {
	var left_out float64
	var top_out float64
	var right_out float64
	var bottom_out float64
	purego.SyscallN(addr_QGraphicsWidget_getContentsMargins_default, this.h, uintptr(unsafe.Pointer(&left_out)), uintptr(unsafe.Pointer(&top_out)), uintptr(unsafe.Pointer(&right_out)), uintptr(unsafe.Pointer(&bottom_out)))
	runtime.KeepAlive(this)
	return left_out, top_out, right_out, bottom_out
}

func (this *QGraphicsWidget) PaintWindowFrameDefault(painter *QPainter, option *QStyleOptionGraphicsItem, widget *QWidget) {
	purego.SyscallN(addr_QGraphicsWidget_paintWindowFrame_default, this.h, uintptr(painter.UnsafePointer()), uintptr(option.UnsafePointer()), uintptr(widget.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
	runtime.KeepAlive(option)
	runtime.KeepAlive(widget)
}

func (this *QGraphicsWidget) IsEmptyDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsWidget_isEmpty_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGraphicsWidget) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QGraphicsWidget_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QGraphicsWidget) TypeDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsWidget_type_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QGraphicsWidget) PaintDefault(painter *QPainter, option *QStyleOptionGraphicsItem, widget *QWidget) {
	purego.SyscallN(addr_QGraphicsWidget_paint_default, this.h, uintptr(painter.UnsafePointer()), uintptr(option.UnsafePointer()), uintptr(widget.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
	runtime.KeepAlive(option)
	runtime.KeepAlive(widget)
}

func (this *QGraphicsWidget) BoundingRectDefault() *QRectF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsWidget_boundingRect_default, this.h)
	runtime.KeepAlive(this)
	return newownedQRectF(__r)
}

func (this *QGraphicsWidget) ShapeDefault() *QPainterPath {
	__r, _, _ := purego.SyscallN(addr_QGraphicsWidget_shape_default, this.h)
	runtime.KeepAlive(this)
	return newownedQPainterPath(__r)
}

func (this *QGraphicsWidget) InitStyleOptionDefault(option *QStyleOption) {
	purego.SyscallN(addr_QGraphicsWidget_initStyleOption_default, this.h, uintptr(option.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(option)
}

func (this *QGraphicsWidget) SizeHintDefault(which SizeHint, constraint *QSizeF) *QSizeF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsWidget_sizeHint_default, this.h, uintptr(which), uintptr(constraint.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(constraint)
	return newownedQSizeF(__r)
}

func (this *QGraphicsWidget) UpdateGeometryDefault() {
	purego.SyscallN(addr_QGraphicsWidget_updateGeometry_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsWidget) ItemChangeDefault(change QGraphicsItem__GraphicsItemChange, value *QVariant) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QGraphicsWidget_itemChange_default, this.h, uintptr(change), uintptr(value.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(value)
	return newownedQVariant(__r)
}

func (this *QGraphicsWidget) PropertyChangeDefault(propertyName string, value *QVariant) *QVariant {
	var propertyName_p uintptr
	if len(propertyName) > 0 {
		propertyName_p = uintptr(unsafe.Pointer(unsafe.StringData(propertyName)))
	}
	__r, _, _ := purego.SyscallN(addr_QGraphicsWidget_propertyChange_default, this.h, propertyName_p, uintptr(len(propertyName)), uintptr(value.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(propertyName)
	runtime.KeepAlive(value)
	return newownedQVariant(__r)
}

func (this *QGraphicsWidget) SceneEventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsWidget_sceneEvent_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QGraphicsWidget) WindowFrameEventDefault(e *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsWidget_windowFrameEvent_default, this.h, uintptr(e.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
	return __r != 0
}

func (this *QGraphicsWidget) WindowFrameSectionAtDefault(pos *QPointF) WindowFrameSection {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsWidget_windowFrameSectionAt_default, this.h, uintptr(pos.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pos)
	return WindowFrameSection(__r)
}

func (this *QGraphicsWidget) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsWidget_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QGraphicsWidget) ChangeEventDefault(event *QEvent) {
	purego.SyscallN(addr_QGraphicsWidget_changeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsWidget) CloseEventDefault(event *QCloseEvent) {
	purego.SyscallN(addr_QGraphicsWidget_closeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsWidget) FocusInEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QGraphicsWidget_focusInEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsWidget) FocusNextPrevChildDefault(next bool) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsWidget_focusNextPrevChild_default, this.h, uintptr(b2i(next)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGraphicsWidget) FocusOutEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QGraphicsWidget_focusOutEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsWidget) HideEventDefault(event *QHideEvent) {
	purego.SyscallN(addr_QGraphicsWidget_hideEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsWidget) MoveEventDefault(event *QGraphicsSceneMoveEvent) {
	purego.SyscallN(addr_QGraphicsWidget_moveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsWidget) PolishEventDefault() {
	purego.SyscallN(addr_QGraphicsWidget_polishEvent_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsWidget) ResizeEventDefault(event *QGraphicsSceneResizeEvent) {
	purego.SyscallN(addr_QGraphicsWidget_resizeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsWidget) ShowEventDefault(event *QShowEvent) {
	purego.SyscallN(addr_QGraphicsWidget_showEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsWidget) HoverMoveEventDefault(event *QGraphicsSceneHoverEvent) {
	purego.SyscallN(addr_QGraphicsWidget_hoverMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsWidget) HoverLeaveEventDefault(event *QGraphicsSceneHoverEvent) {
	purego.SyscallN(addr_QGraphicsWidget_hoverLeaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsWidget) GrabMouseEventDefault(event *QEvent) {
	purego.SyscallN(addr_QGraphicsWidget_grabMouseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsWidget) UngrabMouseEventDefault(event *QEvent) {
	purego.SyscallN(addr_QGraphicsWidget_ungrabMouseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsWidget) GrabKeyboardEventDefault(event *QEvent) {
	purego.SyscallN(addr_QGraphicsWidget_grabKeyboardEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsWidget) UngrabKeyboardEventDefault(event *QEvent) {
	purego.SyscallN(addr_QGraphicsWidget_ungrabKeyboardEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsWidget) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsWidget_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QGraphicsWidget) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QGraphicsWidget_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsWidget) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QGraphicsWidget_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsWidget) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QGraphicsWidget_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsWidget) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QGraphicsWidget_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QGraphicsWidget) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QGraphicsWidget_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QGraphicsWidget) AdvanceDefault(phase int) {
	purego.SyscallN(addr_QGraphicsWidget_advance_default, this.h, uintptr(phase))
	runtime.KeepAlive(this)
}

func (this *QGraphicsWidget) ContainsDefault(point *QPointF) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsWidget_contains_default, this.h, uintptr(point.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(point)
	return __r != 0
}

func (this *QGraphicsWidget) CollidesWithItemDefault(other *QGraphicsItem, mode ItemSelectionMode) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsWidget_collidesWithItem_default, this.h, uintptr(other.UnsafePointer()), uintptr(mode))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return __r != 0
}

func (this *QGraphicsWidget) CollidesWithPathDefault(path *QPainterPath, mode ItemSelectionMode) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsWidget_collidesWithPath_default, this.h, uintptr(path.UnsafePointer()), uintptr(mode))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(path)
	return __r != 0
}

func (this *QGraphicsWidget) IsObscuredByDefault(item *QGraphicsItem) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsWidget_isObscuredBy_default, this.h, uintptr(item.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
	return __r != 0
}

func (this *QGraphicsWidget) OpaqueAreaDefault() *QPainterPath {
	__r, _, _ := purego.SyscallN(addr_QGraphicsWidget_opaqueArea_default, this.h)
	runtime.KeepAlive(this)
	return newownedQPainterPath(__r)
}

func (this *QGraphicsWidget) SceneEventFilterDefault(watched *QGraphicsItem, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsWidget_sceneEventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QGraphicsWidget) ContextMenuEventDefault(event *QGraphicsSceneContextMenuEvent) {
	purego.SyscallN(addr_QGraphicsWidget_contextMenuEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsWidget) DragEnterEventDefault(event *QGraphicsSceneDragDropEvent) {
	purego.SyscallN(addr_QGraphicsWidget_dragEnterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsWidget) DragLeaveEventDefault(event *QGraphicsSceneDragDropEvent) {
	purego.SyscallN(addr_QGraphicsWidget_dragLeaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsWidget) DragMoveEventDefault(event *QGraphicsSceneDragDropEvent) {
	purego.SyscallN(addr_QGraphicsWidget_dragMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsWidget) DropEventDefault(event *QGraphicsSceneDragDropEvent) {
	purego.SyscallN(addr_QGraphicsWidget_dropEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsWidget) HoverEnterEventDefault(event *QGraphicsSceneHoverEvent) {
	purego.SyscallN(addr_QGraphicsWidget_hoverEnterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsWidget) KeyPressEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QGraphicsWidget_keyPressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsWidget) KeyReleaseEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QGraphicsWidget_keyReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsWidget) MousePressEventDefault(event *QGraphicsSceneMouseEvent) {
	purego.SyscallN(addr_QGraphicsWidget_mousePressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsWidget) MouseMoveEventDefault(event *QGraphicsSceneMouseEvent) {
	purego.SyscallN(addr_QGraphicsWidget_mouseMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsWidget) MouseReleaseEventDefault(event *QGraphicsSceneMouseEvent) {
	purego.SyscallN(addr_QGraphicsWidget_mouseReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsWidget) MouseDoubleClickEventDefault(event *QGraphicsSceneMouseEvent) {
	purego.SyscallN(addr_QGraphicsWidget_mouseDoubleClickEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsWidget) WheelEventDefault(event *QGraphicsSceneWheelEvent) {
	purego.SyscallN(addr_QGraphicsWidget_wheelEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsWidget) InputMethodEventDefault(event *QInputMethodEvent) {
	purego.SyscallN(addr_QGraphicsWidget_inputMethodEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsWidget) InputMethodQueryDefault(query InputMethodQuery) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QGraphicsWidget_inputMethodQuery_default, this.h, uintptr(query))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QGraphicsWidget) SupportsExtensionDefault(extension QGraphicsItem__Extension) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsWidget_supportsExtension_default, this.h, uintptr(extension))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGraphicsWidget) SetExtensionDefault(extension QGraphicsItem__Extension, variant *QVariant) {
	purego.SyscallN(addr_QGraphicsWidget_setExtension_default, this.h, uintptr(extension), uintptr(variant.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(variant)
}

func (this *QGraphicsWidget) ExtensionDefault(variant *QVariant) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QGraphicsWidget_extension_default, this.h, uintptr(variant.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(variant)
	return newownedQVariant(__r)
}

func (this *QGraphicsWidget) OnGeometryChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QGraphicsWidget_connect_geometryChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsWidget) OnLayoutChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QGraphicsWidget_connect_layoutChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsWidget) OnParentChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QGraphicsWidget_connect_parentChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsWidget) OnOpacityChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QGraphicsWidget_connect_opacityChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsWidget) OnVisibleChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QGraphicsWidget_connect_visibleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsWidget) OnEnabledChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QGraphicsWidget_connect_enabledChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsWidget) OnXChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QGraphicsWidget_connect_xChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsWidget) OnYChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QGraphicsWidget_connect_yChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsWidget) OnZChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QGraphicsWidget_connect_zChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsWidget) OnRotationChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QGraphicsWidget_connect_rotationChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsWidget) OnScaleChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QGraphicsWidget_connect_scaleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsWidget) OnChildrenChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QGraphicsWidget_connect_childrenChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsWidget) OnWidthChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QGraphicsWidget_connect_widthChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsWidget) OnHeightChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QGraphicsWidget_connect_heightChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsWidget) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QGraphicsWidget_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsWidget) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QGraphicsWidget_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsWidget) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QGraphicsWidget_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsWidget) Disconnect(handle int64) {
	pg_QGraphicsWidget__disconnect(handle)
	unregisterCallback(handle)
}

func QGraphicsWidget_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case 5684820848726969227:
		return 1
	case 164760286202646972:
		return 2
	case -6371403872082420691:
		return 3
	case 6083584261053556819:
		return 4
	case -1640122132753285616:
		return 5
	case 2510008782785038831:
		return 6
	case 7938019069247444294:
		return 7
	case -488693769771308531:
		return 8
	case 729690261457684101:
		return 9
	case 8082455443689927668:
		return 10
	case -1888572168794878830:
		return 11
	case 5289089528078387296:
		return 12
	case -2846371507838587071:
		return 13
	case -7362375632910814166:
		return 14
	case 8834866722574844732:
		return 15
	case 1237236456312824991:
		return 16
	case 4959376063826165943:
		return 17
	case -1989937050019464783:
		return 18
	case 975043767033552552:
		return 19
	case 5898328517505784191:
		return 20
	case -385596010576048993:
		return 21
	case -199718707876719129:
		return 22
	case -9161949150489525396:
		return 23
	case -214704830098031580:
		return 24
	case -4302917285655014661:
		return 25
	case -6825369174751916:
		return 26
	case -2745191522144316396:
		return 27
	case 7944408482203603984:
		return 28
	case 7928460550911621078:
		return 29
	case -5662911227992581871:
		return 30
	case 6423297148758141226:
		return 31
	case 2912496686457383679:
		return 32
	case 5217623514630858297:
		return 33
	case -5409927581812374498:
		return 34
	case -1185315367526857833:
		return 35
	case 3769868229123868474:
		return 36
	case 7379029982170503008:
		return 37
	case -568369772982215962:
		return 38
	case -4830836341480073475:
		return 39
	case 5645366511944179768:
		return 40
	case 6469736184362593173:
		return 41
	case 6636575182718753853:
		return 42
	case -7215030271914903671:
		return 43
	case 399370970023989435:
		return 44
	case -4584203218241215725:
		return 45
	case -6698043222227006015:
		return 46
	case 206122715919917589:
		return 47
	case -3139904158373725006:
		return 48
	case -1781545137352745978:
		return 49
	case 9107058682754591716:
		return 50
	case 6932941943906681679:
		return 51
	case 1023791977679793185:
		return 52
	case 772598475950050299:
		return 53
	case -332090898078543319:
		return 54
	case -4255494004099806581:
		return 55
	case -4415920049297795213:
		return 56
	case -4029073905697023887:
		return 57
	case 1919988340934511316:
		return 58
	case 5438819661942843222:
		return 59
	case 7345853097340709638:
		return 60
	case -5729166443651306418:
		return 61
	case -6730181576106171580:
		return 62
	case 2270616358835251226:
		return 63
	case -626822430251098498:
		return 64
	}
	return -1
}

func QGraphicsWidget_enable(h uintptr, idx int32) { pg_QGraphicsWidget__set_override(h, idx) }

func (this *QGraphicsWidget) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QGraphicsWidget) OnSetGeometry(fn func(*QRectF)) {
	registerVirt(this.UnsafeGoID(), 5684820848726969227, func(rect uintptr) { fn(UnsafeNewQRectF(unsafe.Pointer(rect))) }) // slot "setGeometry"
}

func (this *QGraphicsWidget) OnGetContentsMargins(fn func(*float64, *float64, *float64, *float64)) {
	registerVirt(this.UnsafeGoID(), 164760286202646972, func(left uintptr, top uintptr, right uintptr, bottom uintptr) { fn((*float64)(unsafe.Pointer(left)), (*float64)(unsafe.Pointer(top)), (*float64)(unsafe.Pointer(right)), (*float64)(unsafe.Pointer(bottom))) }) // slot "getContentsMargins"
}

func (this *QGraphicsWidget) OnType(fn func() int) {
	registerVirt(this.UnsafeGoID(), -6371403872082420691, func() int32 { return int32(fn()) }) // slot "type"
}

func (this *QGraphicsWidget) OnPaint(fn func(*QPainter, *QStyleOptionGraphicsItem, *QWidget)) {
	registerVirt(this.UnsafeGoID(), 6083584261053556819, func(painter uintptr, option uintptr, widget uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter)), UnsafeNewQStyleOptionGraphicsItem(unsafe.Pointer(option)), UnsafeNewQWidget(unsafe.Pointer(widget))) }) // slot "paint"
}

func (this *QGraphicsWidget) OnPaintWindowFrame(fn func(*QPainter, *QStyleOptionGraphicsItem, *QWidget)) {
	registerVirt(this.UnsafeGoID(), -1640122132753285616, func(painter uintptr, option uintptr, widget uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter)), UnsafeNewQStyleOptionGraphicsItem(unsafe.Pointer(option)), UnsafeNewQWidget(unsafe.Pointer(widget))) }) // slot "paintWindowFrame"
}

func (this *QGraphicsWidget) OnBoundingRect(fn func() *QRectF) {
	registerVirt(this.UnsafeGoID(), 2510008782785038831, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "boundingRect"
}

func (this *QGraphicsWidget) OnShape(fn func() *QPainterPath) {
	registerVirt(this.UnsafeGoID(), 7938019069247444294, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "shape"
}

func (this *QGraphicsWidget) OnInitStyleOption(fn func(*QStyleOption)) {
	registerVirt(this.UnsafeGoID(), -488693769771308531, func(option uintptr) { fn(UnsafeNewQStyleOption(unsafe.Pointer(option))) }) // slot "initStyleOption"
}

func (this *QGraphicsWidget) OnSizeHint(fn func(SizeHint, *QSizeF) *QSizeF) {
	registerVirt(this.UnsafeGoID(), 729690261457684101, func(which int32, constraint uintptr) uintptr { return uintptr(fn(SizeHint(which), UnsafeNewQSizeF(unsafe.Pointer(constraint))).UnsafePointer()) }) // slot "sizeHint"
}

func (this *QGraphicsWidget) OnUpdateGeometry(fn func()) {
	registerVirt(this.UnsafeGoID(), 8082455443689927668, func() { fn() }) // slot "updateGeometry"
}

func (this *QGraphicsWidget) OnItemChange(fn func(QGraphicsItem__GraphicsItemChange, *QVariant) *QVariant) {
	registerVirt(this.UnsafeGoID(), -1888572168794878830, func(change int32, value uintptr) uintptr { return uintptr(fn(QGraphicsItem__GraphicsItemChange(change), UnsafeNewQVariant(unsafe.Pointer(value))).UnsafePointer()) }) // slot "itemChange"
}

func (this *QGraphicsWidget) OnPropertyChange(fn func(string, *QVariant) *QVariant) {
	registerVirt(this.UnsafeGoID(), 5289089528078387296, func(propertyName uintptr, propertyName_len int64, value uintptr) uintptr { return uintptr(fn(pgStrNoFree(propertyName, propertyName_len), UnsafeNewQVariant(unsafe.Pointer(value))).UnsafePointer()) }) // slot "propertyChange"
}

func (this *QGraphicsWidget) OnSceneEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), -2846371507838587071, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "sceneEvent"
}

func (this *QGraphicsWidget) OnWindowFrameEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), -7362375632910814166, func(e uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(e)))) }) // slot "windowFrameEvent"
}

func (this *QGraphicsWidget) OnWindowFrameSectionAt(fn func(*QPointF) WindowFrameSection) {
	registerVirt(this.UnsafeGoID(), 8834866722574844732, func(pos uintptr) int32 { return int32(fn(UnsafeNewQPointF(unsafe.Pointer(pos)))) }) // slot "windowFrameSectionAt"
}

func (this *QGraphicsWidget) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QGraphicsWidget) OnChangeEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 4959376063826165943, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "changeEvent"
}

func (this *QGraphicsWidget) OnCloseEvent(fn func(*QCloseEvent)) {
	registerVirt(this.UnsafeGoID(), -1989937050019464783, func(event uintptr) { fn(UnsafeNewQCloseEvent(unsafe.Pointer(event))) }) // slot "closeEvent"
}

func (this *QGraphicsWidget) OnFocusInEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), 975043767033552552, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusInEvent"
}

func (this *QGraphicsWidget) OnFocusNextPrevChild(fn func(bool) bool) {
	registerVirt(this.UnsafeGoID(), 5898328517505784191, func(next int32) int32 { return b2i(fn(next != 0)) }) // slot "focusNextPrevChild"
}

func (this *QGraphicsWidget) OnFocusOutEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), -385596010576048993, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusOutEvent"
}

func (this *QGraphicsWidget) OnHideEvent(fn func(*QHideEvent)) {
	registerVirt(this.UnsafeGoID(), -199718707876719129, func(event uintptr) { fn(UnsafeNewQHideEvent(unsafe.Pointer(event))) }) // slot "hideEvent"
}

func (this *QGraphicsWidget) OnMoveEvent(fn func(*QGraphicsSceneMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -9161949150489525396, func(event uintptr) { fn(UnsafeNewQGraphicsSceneMoveEvent(unsafe.Pointer(event))) }) // slot "moveEvent"
}

func (this *QGraphicsWidget) OnPolishEvent(fn func()) {
	registerVirt(this.UnsafeGoID(), -214704830098031580, func() { fn() }) // slot "polishEvent"
}

func (this *QGraphicsWidget) OnResizeEvent(fn func(*QGraphicsSceneResizeEvent)) {
	registerVirt(this.UnsafeGoID(), -4302917285655014661, func(event uintptr) { fn(UnsafeNewQGraphicsSceneResizeEvent(unsafe.Pointer(event))) }) // slot "resizeEvent"
}

func (this *QGraphicsWidget) OnShowEvent(fn func(*QShowEvent)) {
	registerVirt(this.UnsafeGoID(), -6825369174751916, func(event uintptr) { fn(UnsafeNewQShowEvent(unsafe.Pointer(event))) }) // slot "showEvent"
}

func (this *QGraphicsWidget) OnHoverMoveEvent(fn func(*QGraphicsSceneHoverEvent)) {
	registerVirt(this.UnsafeGoID(), -2745191522144316396, func(event uintptr) { fn(UnsafeNewQGraphicsSceneHoverEvent(unsafe.Pointer(event))) }) // slot "hoverMoveEvent"
}

func (this *QGraphicsWidget) OnHoverLeaveEvent(fn func(*QGraphicsSceneHoverEvent)) {
	registerVirt(this.UnsafeGoID(), 7944408482203603984, func(event uintptr) { fn(UnsafeNewQGraphicsSceneHoverEvent(unsafe.Pointer(event))) }) // slot "hoverLeaveEvent"
}

func (this *QGraphicsWidget) OnGrabMouseEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 7928460550911621078, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "grabMouseEvent"
}

func (this *QGraphicsWidget) OnUngrabMouseEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), -5662911227992581871, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "ungrabMouseEvent"
}

func (this *QGraphicsWidget) OnGrabKeyboardEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 6423297148758141226, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "grabKeyboardEvent"
}

func (this *QGraphicsWidget) OnUngrabKeyboardEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 2912496686457383679, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "ungrabKeyboardEvent"
}

func (this *QGraphicsWidget) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QGraphicsWidget) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QGraphicsWidget) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QGraphicsWidget) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QGraphicsWidget) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QGraphicsWidget) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QGraphicsWidget) OnAdvance(fn func(int)) {
	registerVirt(this.UnsafeGoID(), -4830836341480073475, func(phase int32) { fn(int(phase)) }) // slot "advance"
}

func (this *QGraphicsWidget) OnContains(fn func(*QPointF) bool) {
	registerVirt(this.UnsafeGoID(), 5645366511944179768, func(point uintptr) int32 { return b2i(fn(UnsafeNewQPointF(unsafe.Pointer(point)))) }) // slot "contains"
}

func (this *QGraphicsWidget) OnCollidesWithItem(fn func(*QGraphicsItem, ItemSelectionMode) bool) {
	registerVirt(this.UnsafeGoID(), 6469736184362593173, func(other uintptr, mode int32) int32 { return b2i(fn(UnsafeNewQGraphicsItem(unsafe.Pointer(other)), ItemSelectionMode(mode))) }) // slot "collidesWithItem"
}

func (this *QGraphicsWidget) OnCollidesWithPath(fn func(*QPainterPath, ItemSelectionMode) bool) {
	registerVirt(this.UnsafeGoID(), 6636575182718753853, func(path uintptr, mode int32) int32 { return b2i(fn(UnsafeNewQPainterPath(unsafe.Pointer(path)), ItemSelectionMode(mode))) }) // slot "collidesWithPath"
}

func (this *QGraphicsWidget) OnIsObscuredBy(fn func(*QGraphicsItem) bool) {
	registerVirt(this.UnsafeGoID(), -7215030271914903671, func(item uintptr) int32 { return b2i(fn(UnsafeNewQGraphicsItem(unsafe.Pointer(item)))) }) // slot "isObscuredBy"
}

func (this *QGraphicsWidget) OnOpaqueArea(fn func() *QPainterPath) {
	registerVirt(this.UnsafeGoID(), 399370970023989435, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "opaqueArea"
}

func (this *QGraphicsWidget) OnSceneEventFilter(fn func(*QGraphicsItem, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), -4584203218241215725, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQGraphicsItem(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "sceneEventFilter"
}

func (this *QGraphicsWidget) OnContextMenuEvent(fn func(*QGraphicsSceneContextMenuEvent)) {
	registerVirt(this.UnsafeGoID(), -6698043222227006015, func(event uintptr) { fn(UnsafeNewQGraphicsSceneContextMenuEvent(unsafe.Pointer(event))) }) // slot "contextMenuEvent"
}

func (this *QGraphicsWidget) OnDragEnterEvent(fn func(*QGraphicsSceneDragDropEvent)) {
	registerVirt(this.UnsafeGoID(), 206122715919917589, func(event uintptr) { fn(UnsafeNewQGraphicsSceneDragDropEvent(unsafe.Pointer(event))) }) // slot "dragEnterEvent"
}

func (this *QGraphicsWidget) OnDragLeaveEvent(fn func(*QGraphicsSceneDragDropEvent)) {
	registerVirt(this.UnsafeGoID(), -3139904158373725006, func(event uintptr) { fn(UnsafeNewQGraphicsSceneDragDropEvent(unsafe.Pointer(event))) }) // slot "dragLeaveEvent"
}

func (this *QGraphicsWidget) OnDragMoveEvent(fn func(*QGraphicsSceneDragDropEvent)) {
	registerVirt(this.UnsafeGoID(), -1781545137352745978, func(event uintptr) { fn(UnsafeNewQGraphicsSceneDragDropEvent(unsafe.Pointer(event))) }) // slot "dragMoveEvent"
}

func (this *QGraphicsWidget) OnDropEvent(fn func(*QGraphicsSceneDragDropEvent)) {
	registerVirt(this.UnsafeGoID(), 9107058682754591716, func(event uintptr) { fn(UnsafeNewQGraphicsSceneDragDropEvent(unsafe.Pointer(event))) }) // slot "dropEvent"
}

func (this *QGraphicsWidget) OnHoverEnterEvent(fn func(*QGraphicsSceneHoverEvent)) {
	registerVirt(this.UnsafeGoID(), 6932941943906681679, func(event uintptr) { fn(UnsafeNewQGraphicsSceneHoverEvent(unsafe.Pointer(event))) }) // slot "hoverEnterEvent"
}

func (this *QGraphicsWidget) OnKeyPressEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 1023791977679793185, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyPressEvent"
}

func (this *QGraphicsWidget) OnKeyReleaseEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 772598475950050299, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyReleaseEvent"
}

func (this *QGraphicsWidget) OnMousePressEvent(fn func(*QGraphicsSceneMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -332090898078543319, func(event uintptr) { fn(UnsafeNewQGraphicsSceneMouseEvent(unsafe.Pointer(event))) }) // slot "mousePressEvent"
}

func (this *QGraphicsWidget) OnMouseMoveEvent(fn func(*QGraphicsSceneMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4255494004099806581, func(event uintptr) { fn(UnsafeNewQGraphicsSceneMouseEvent(unsafe.Pointer(event))) }) // slot "mouseMoveEvent"
}

func (this *QGraphicsWidget) OnMouseReleaseEvent(fn func(*QGraphicsSceneMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4415920049297795213, func(event uintptr) { fn(UnsafeNewQGraphicsSceneMouseEvent(unsafe.Pointer(event))) }) // slot "mouseReleaseEvent"
}

func (this *QGraphicsWidget) OnMouseDoubleClickEvent(fn func(*QGraphicsSceneMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4029073905697023887, func(event uintptr) { fn(UnsafeNewQGraphicsSceneMouseEvent(unsafe.Pointer(event))) }) // slot "mouseDoubleClickEvent"
}

func (this *QGraphicsWidget) OnWheelEvent(fn func(*QGraphicsSceneWheelEvent)) {
	registerVirt(this.UnsafeGoID(), 1919988340934511316, func(event uintptr) { fn(UnsafeNewQGraphicsSceneWheelEvent(unsafe.Pointer(event))) }) // slot "wheelEvent"
}

func (this *QGraphicsWidget) OnInputMethodEvent(fn func(*QInputMethodEvent)) {
	registerVirt(this.UnsafeGoID(), 5438819661942843222, func(event uintptr) { fn(UnsafeNewQInputMethodEvent(unsafe.Pointer(event))) }) // slot "inputMethodEvent"
}

func (this *QGraphicsWidget) OnInputMethodQuery(fn func(InputMethodQuery) *QVariant) {
	registerVirt(this.UnsafeGoID(), 7345853097340709638, func(query int32) uintptr { return uintptr(fn(InputMethodQuery(query)).UnsafePointer()) }) // slot "inputMethodQuery"
}

func (this *QGraphicsWidget) OnSupportsExtension(fn func(QGraphicsItem__Extension) bool) {
	registerVirt(this.UnsafeGoID(), -5729166443651306418, func(extension int32) int32 { return b2i(fn(QGraphicsItem__Extension(extension))) }) // slot "supportsExtension"
}

func (this *QGraphicsWidget) OnSetExtension(fn func(QGraphicsItem__Extension, *QVariant)) {
	registerVirt(this.UnsafeGoID(), -6730181576106171580, func(extension int32, variant uintptr) { fn(QGraphicsItem__Extension(extension), UnsafeNewQVariant(unsafe.Pointer(variant))) }) // slot "setExtension"
}

func (this *QGraphicsWidget) OnExtension(fn func(*QVariant) *QVariant) {
	registerVirt(this.UnsafeGoID(), 2270616358835251226, func(variant uintptr) uintptr { return uintptr(fn(UnsafeNewQVariant(unsafe.Pointer(variant))).UnsafePointer()) }) // slot "extension"
}

func (this *QGraphicsWidget) OnIsEmpty(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -626822430251098498, func() int32 { return b2i(fn()) }) // slot "isEmpty"
}

func (this *QGraphicsWidget) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QGraphicsWidget_delete(this.h)
	}
	this.UnsafeRelease()
}
