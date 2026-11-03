package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QGraphicsObject_base func(uintptr) uintptr
	addr_QGraphicsObject_tr uintptr
	addr_QGraphicsObject_grabGesture uintptr
	addr_QGraphicsObject_ungrabGesture uintptr
	addr_QGraphicsObject_updateMicroFocus uintptr
	addr_QGraphicsObject_parentChanged uintptr
	addr_QGraphicsObject_opacityChanged uintptr
	addr_QGraphicsObject_visibleChanged uintptr
	addr_QGraphicsObject_enabledChanged uintptr
	addr_QGraphicsObject_xChanged uintptr
	addr_QGraphicsObject_yChanged uintptr
	addr_QGraphicsObject_zChanged uintptr
	addr_QGraphicsObject_rotationChanged uintptr
	addr_QGraphicsObject_scaleChanged uintptr
	addr_QGraphicsObject_childrenChanged uintptr
	addr_QGraphicsObject_widthChanged uintptr
	addr_QGraphicsObject_heightChanged uintptr
	addr_QGraphicsObject_connect uintptr
	addr_QGraphicsObject_connect2 uintptr
	addr_QGraphicsObject_sender uintptr
	addr_QGraphicsObject_senderSignalIndex uintptr
	addr_QGraphicsObject_receivers uintptr
	addr_QGraphicsObject_isSignalConnected uintptr
	addr_QGraphicsObject_scene uintptr
	addr_QGraphicsObject_parentItem uintptr
	addr_QGraphicsObject_topLevelItem uintptr
	addr_QGraphicsObject_parentObject uintptr
	addr_QGraphicsObject_parentWidget uintptr
	addr_QGraphicsObject_topLevelWidget uintptr
	addr_QGraphicsObject_window uintptr
	addr_QGraphicsObject_panel uintptr
	addr_QGraphicsObject_setParentItem uintptr
	addr_QGraphicsObject_childItems uintptr
	addr_QGraphicsObject_isWidget uintptr
	addr_QGraphicsObject_isWindow uintptr
	addr_QGraphicsObject_isPanel uintptr
	addr_QGraphicsObject_toGraphicsObject uintptr
	addr_QGraphicsObject_toGraphicsObject2 uintptr
	addr_QGraphicsObject_group uintptr
	addr_QGraphicsObject_setGroup uintptr
	addr_QGraphicsObject_flags uintptr
	addr_QGraphicsObject_setFlag uintptr
	addr_QGraphicsObject_setFlags uintptr
	addr_QGraphicsObject_cacheMode uintptr
	addr_QGraphicsObject_setCacheMode uintptr
	addr_QGraphicsObject_panelModality uintptr
	addr_QGraphicsObject_setPanelModality uintptr
	addr_QGraphicsObject_isBlockedByModalPanel uintptr
	addr_QGraphicsObject_toolTip uintptr
	addr_QGraphicsObject_setToolTip uintptr
	addr_QGraphicsObject_cursor uintptr
	addr_QGraphicsObject_setCursor uintptr
	addr_QGraphicsObject_hasCursor uintptr
	addr_QGraphicsObject_unsetCursor uintptr
	addr_QGraphicsObject_isVisible uintptr
	addr_QGraphicsObject_isVisibleTo uintptr
	addr_QGraphicsObject_setVisible uintptr
	addr_QGraphicsObject_hide uintptr
	addr_QGraphicsObject_show uintptr
	addr_QGraphicsObject_isEnabled uintptr
	addr_QGraphicsObject_setEnabled uintptr
	addr_QGraphicsObject_isSelected uintptr
	addr_QGraphicsObject_setSelected uintptr
	addr_QGraphicsObject_acceptDrops uintptr
	addr_QGraphicsObject_setAcceptDrops uintptr
	pg_QGraphicsObject_opacity func(uintptr) float64
	pg_QGraphicsObject_effectiveOpacity func(uintptr) float64
	pg_QGraphicsObject_setOpacity func(uintptr, float64)
	addr_QGraphicsObject_graphicsEffect uintptr
	addr_QGraphicsObject_setGraphicsEffect uintptr
	addr_QGraphicsObject_acceptedMouseButtons uintptr
	addr_QGraphicsObject_setAcceptedMouseButtons uintptr
	addr_QGraphicsObject_acceptHoverEvents uintptr
	addr_QGraphicsObject_setAcceptHoverEvents uintptr
	addr_QGraphicsObject_acceptTouchEvents uintptr
	addr_QGraphicsObject_setAcceptTouchEvents uintptr
	addr_QGraphicsObject_filtersChildEvents uintptr
	addr_QGraphicsObject_setFiltersChildEvents uintptr
	addr_QGraphicsObject_handlesChildEvents uintptr
	addr_QGraphicsObject_setHandlesChildEvents uintptr
	addr_QGraphicsObject_isActive uintptr
	addr_QGraphicsObject_setActive uintptr
	addr_QGraphicsObject_hasFocus uintptr
	addr_QGraphicsObject_setFocus uintptr
	addr_QGraphicsObject_clearFocus uintptr
	addr_QGraphicsObject_focusProxy uintptr
	addr_QGraphicsObject_setFocusProxy uintptr
	addr_QGraphicsObject_focusItem uintptr
	addr_QGraphicsObject_focusScopeItem uintptr
	addr_QGraphicsObject_grabMouse uintptr
	addr_QGraphicsObject_ungrabMouse uintptr
	addr_QGraphicsObject_grabKeyboard uintptr
	addr_QGraphicsObject_ungrabKeyboard uintptr
	addr_QGraphicsObject_pos uintptr
	pg_QGraphicsObject_x func(uintptr) float64
	pg_QGraphicsObject_setX func(uintptr, float64)
	pg_QGraphicsObject_y func(uintptr) float64
	pg_QGraphicsObject_setY func(uintptr, float64)
	addr_QGraphicsObject_scenePos uintptr
	addr_QGraphicsObject_setPos uintptr
	pg_QGraphicsObject_setPos2 func(uintptr, float64, float64)
	pg_QGraphicsObject_moveBy func(uintptr, float64, float64)
	addr_QGraphicsObject_ensureVisible uintptr
	pg_QGraphicsObject_ensureVisible2 func(uintptr, float64, float64, float64, float64)
	addr_QGraphicsObject_transform uintptr
	addr_QGraphicsObject_sceneTransform uintptr
	addr_QGraphicsObject_deviceTransform uintptr
	addr_QGraphicsObject_itemTransform uintptr
	addr_QGraphicsObject_setTransform uintptr
	addr_QGraphicsObject_resetTransform uintptr
	pg_QGraphicsObject_setRotation func(uintptr, float64)
	pg_QGraphicsObject_rotation func(uintptr) float64
	pg_QGraphicsObject_setScale func(uintptr, float64)
	pg_QGraphicsObject_scale func(uintptr) float64
	addr_QGraphicsObject_transformations uintptr
	addr_QGraphicsObject_setTransformations uintptr
	addr_QGraphicsObject_transformOriginPoint uintptr
	addr_QGraphicsObject_setTransformOriginPoint uintptr
	pg_QGraphicsObject_setTransformOriginPoint2 func(uintptr, float64, float64)
	addr_QGraphicsObject_advance uintptr
	pg_QGraphicsObject_zValue func(uintptr) float64
	pg_QGraphicsObject_setZValue func(uintptr, float64)
	addr_QGraphicsObject_stackBefore uintptr
	addr_QGraphicsObject_boundingRect uintptr
	addr_QGraphicsObject_childrenBoundingRect uintptr
	addr_QGraphicsObject_sceneBoundingRect uintptr
	addr_QGraphicsObject_shape uintptr
	addr_QGraphicsObject_isClipped uintptr
	addr_QGraphicsObject_clipPath uintptr
	addr_QGraphicsObject_contains uintptr
	addr_QGraphicsObject_collidesWithItem uintptr
	addr_QGraphicsObject_collidesWithPath uintptr
	addr_QGraphicsObject_collidingItems uintptr
	addr_QGraphicsObject_isObscured uintptr
	pg_QGraphicsObject_isObscured2 func(uintptr, float64, float64, float64, float64) int32
	addr_QGraphicsObject_isObscuredBy uintptr
	addr_QGraphicsObject_opaqueArea uintptr
	addr_QGraphicsObject_boundingRegion uintptr
	pg_QGraphicsObject_boundingRegionGranularity func(uintptr) float64
	pg_QGraphicsObject_setBoundingRegionGranularity func(uintptr, float64)
	addr_QGraphicsObject_paint uintptr
	addr_QGraphicsObject_update uintptr
	pg_QGraphicsObject_update2 func(uintptr, float64, float64, float64, float64)
	pg_QGraphicsObject_scroll func(uintptr, float64, float64)
	addr_QGraphicsObject_mapToItem uintptr
	addr_QGraphicsObject_mapToParent uintptr
	addr_QGraphicsObject_mapToScene uintptr
	addr_QGraphicsObject_mapToItem2 uintptr
	addr_QGraphicsObject_mapToParent2 uintptr
	addr_QGraphicsObject_mapToScene2 uintptr
	addr_QGraphicsObject_mapRectToItem uintptr
	addr_QGraphicsObject_mapRectToParent uintptr
	addr_QGraphicsObject_mapRectToScene uintptr
	addr_QGraphicsObject_mapToItem3 uintptr
	addr_QGraphicsObject_mapToParent3 uintptr
	addr_QGraphicsObject_mapToScene3 uintptr
	addr_QGraphicsObject_mapToItem4 uintptr
	addr_QGraphicsObject_mapToParent4 uintptr
	addr_QGraphicsObject_mapToScene4 uintptr
	addr_QGraphicsObject_mapFromItem uintptr
	addr_QGraphicsObject_mapFromParent uintptr
	addr_QGraphicsObject_mapFromScene uintptr
	addr_QGraphicsObject_mapFromItem2 uintptr
	addr_QGraphicsObject_mapFromParent2 uintptr
	addr_QGraphicsObject_mapFromScene2 uintptr
	addr_QGraphicsObject_mapRectFromItem uintptr
	addr_QGraphicsObject_mapRectFromParent uintptr
	addr_QGraphicsObject_mapRectFromScene uintptr
	addr_QGraphicsObject_mapFromItem3 uintptr
	addr_QGraphicsObject_mapFromParent3 uintptr
	addr_QGraphicsObject_mapFromScene3 uintptr
	addr_QGraphicsObject_mapFromItem4 uintptr
	addr_QGraphicsObject_mapFromParent4 uintptr
	addr_QGraphicsObject_mapFromScene4 uintptr
	pg_QGraphicsObject_mapToItem5 func(uintptr, uintptr, float64, float64) uintptr
	pg_QGraphicsObject_mapToParent5 func(uintptr, float64, float64) uintptr
	pg_QGraphicsObject_mapToScene5 func(uintptr, float64, float64) uintptr
	pg_QGraphicsObject_mapToItem6 func(uintptr, uintptr, float64, float64, float64, float64) uintptr
	pg_QGraphicsObject_mapToParent6 func(uintptr, float64, float64, float64, float64) uintptr
	pg_QGraphicsObject_mapToScene6 func(uintptr, float64, float64, float64, float64) uintptr
	pg_QGraphicsObject_mapRectToItem2 func(uintptr, uintptr, float64, float64, float64, float64) uintptr
	pg_QGraphicsObject_mapRectToParent2 func(uintptr, float64, float64, float64, float64) uintptr
	pg_QGraphicsObject_mapRectToScene2 func(uintptr, float64, float64, float64, float64) uintptr
	pg_QGraphicsObject_mapFromItem5 func(uintptr, uintptr, float64, float64) uintptr
	pg_QGraphicsObject_mapFromParent5 func(uintptr, float64, float64) uintptr
	pg_QGraphicsObject_mapFromScene5 func(uintptr, float64, float64) uintptr
	pg_QGraphicsObject_mapFromItem6 func(uintptr, uintptr, float64, float64, float64, float64) uintptr
	pg_QGraphicsObject_mapFromParent6 func(uintptr, float64, float64, float64, float64) uintptr
	pg_QGraphicsObject_mapFromScene6 func(uintptr, float64, float64, float64, float64) uintptr
	pg_QGraphicsObject_mapRectFromItem2 func(uintptr, uintptr, float64, float64, float64, float64) uintptr
	pg_QGraphicsObject_mapRectFromParent2 func(uintptr, float64, float64, float64, float64) uintptr
	pg_QGraphicsObject_mapRectFromScene2 func(uintptr, float64, float64, float64, float64) uintptr
	addr_QGraphicsObject_isAncestorOf uintptr
	addr_QGraphicsObject_commonAncestorItem uintptr
	addr_QGraphicsObject_isUnderMouse uintptr
	addr_QGraphicsObject_data uintptr
	addr_QGraphicsObject_setData uintptr
	addr_QGraphicsObject_inputMethodHints uintptr
	addr_QGraphicsObject_setInputMethodHints uintptr
	addr_QGraphicsObject_type uintptr
	addr_QGraphicsObject_installSceneEventFilter uintptr
	addr_QGraphicsObject_removeSceneEventFilter uintptr
	addr_QGraphicsObject_addToIndex uintptr
	addr_QGraphicsObject_removeFromIndex uintptr
	addr_QGraphicsObject_prepareGeometryChange uintptr
	addr_QGraphicsObject_tr2 uintptr
	addr_QGraphicsObject_tr3 uintptr
	addr_QGraphicsObject_grabGesture2 uintptr
	addr_QGraphicsObject_connect3 uintptr
	addr_QGraphicsObject_connect4 uintptr
	addr_QGraphicsObject_setFlag2 uintptr
	addr_QGraphicsObject_setCacheMode2 uintptr
	addr_QGraphicsObject_setFocus2 uintptr
	addr_QGraphicsObject_ensureVisible3 uintptr
	addr_QGraphicsObject_ensureVisible4 uintptr
	addr_QGraphicsObject_ensureVisible5 uintptr
	pg_QGraphicsObject_ensureVisible6 func(uintptr, float64, float64, float64, float64, int32)
	pg_QGraphicsObject_ensureVisible7 func(uintptr, float64, float64, float64, float64, int32, int32)
	addr_QGraphicsObject_itemTransform2 uintptr
	addr_QGraphicsObject_setTransform2 uintptr
	addr_QGraphicsObject_collidingItems2 uintptr
	addr_QGraphicsObject_isObscured3 uintptr
	addr_QGraphicsObject_update3 uintptr
	pg_QGraphicsObject_scroll2 func(uintptr, float64, float64, uintptr)
	addr_QGraphicsObject_asQGraphicsItem uintptr
	addr_QGraphicsObject_advance_default uintptr
	addr_QGraphicsObject_shape_default uintptr
	addr_QGraphicsObject_contains_default uintptr
	addr_QGraphicsObject_collidesWithItem_default uintptr
	addr_QGraphicsObject_collidesWithPath_default uintptr
	addr_QGraphicsObject_isObscuredBy_default uintptr
	addr_QGraphicsObject_opaqueArea_default uintptr
	addr_QGraphicsObject_type_default uintptr
	addr_QGraphicsObject_metaObject_default uintptr
	addr_QGraphicsObject_event_default uintptr
	addr_QGraphicsObject_eventFilter_default uintptr
	addr_QGraphicsObject_timerEvent_default uintptr
	addr_QGraphicsObject_childEvent_default uintptr
	addr_QGraphicsObject_customEvent_default uintptr
	addr_QGraphicsObject_connectNotify_default uintptr
	addr_QGraphicsObject_disconnectNotify_default uintptr
	addr_QGraphicsObject_sceneEventFilter_default uintptr
	addr_QGraphicsObject_sceneEvent_default uintptr
	addr_QGraphicsObject_contextMenuEvent_default uintptr
	addr_QGraphicsObject_dragEnterEvent_default uintptr
	addr_QGraphicsObject_dragLeaveEvent_default uintptr
	addr_QGraphicsObject_dragMoveEvent_default uintptr
	addr_QGraphicsObject_dropEvent_default uintptr
	addr_QGraphicsObject_focusInEvent_default uintptr
	addr_QGraphicsObject_focusOutEvent_default uintptr
	addr_QGraphicsObject_hoverEnterEvent_default uintptr
	addr_QGraphicsObject_hoverMoveEvent_default uintptr
	addr_QGraphicsObject_hoverLeaveEvent_default uintptr
	addr_QGraphicsObject_keyPressEvent_default uintptr
	addr_QGraphicsObject_keyReleaseEvent_default uintptr
	addr_QGraphicsObject_mousePressEvent_default uintptr
	addr_QGraphicsObject_mouseMoveEvent_default uintptr
	addr_QGraphicsObject_mouseReleaseEvent_default uintptr
	addr_QGraphicsObject_mouseDoubleClickEvent_default uintptr
	addr_QGraphicsObject_wheelEvent_default uintptr
	addr_QGraphicsObject_inputMethodEvent_default uintptr
	addr_QGraphicsObject_inputMethodQuery_default uintptr
	addr_QGraphicsObject_itemChange_default uintptr
	addr_QGraphicsObject_supportsExtension_default uintptr
	addr_QGraphicsObject_setExtension_default uintptr
	addr_QGraphicsObject_extension_default uintptr
	pg_QGraphicsObject_connect_parentChanged func(uintptr, int64)
	pg_QGraphicsObject_connect_opacityChanged func(uintptr, int64)
	pg_QGraphicsObject_connect_visibleChanged func(uintptr, int64)
	pg_QGraphicsObject_connect_enabledChanged func(uintptr, int64)
	pg_QGraphicsObject_connect_xChanged func(uintptr, int64)
	pg_QGraphicsObject_connect_yChanged func(uintptr, int64)
	pg_QGraphicsObject_connect_zChanged func(uintptr, int64)
	pg_QGraphicsObject_connect_rotationChanged func(uintptr, int64)
	pg_QGraphicsObject_connect_scaleChanged func(uintptr, int64)
	pg_QGraphicsObject_connect_childrenChanged func(uintptr, int64)
	pg_QGraphicsObject_connect_widthChanged func(uintptr, int64)
	pg_QGraphicsObject_connect_heightChanged func(uintptr, int64)
	pg_QGraphicsObject_connect_destroyed func(uintptr, int64)
	pg_QGraphicsObject_connect_destroyed2 func(uintptr, int64)
	pg_QGraphicsObject_connect_objectNameChanged func(uintptr, int64)
	pg_QGraphicsObject__disconnect func(int64)
	pg_QGraphicsObject__set_override func(uintptr, int32)
	pg_QGraphicsObject__Go_new func(int64) uintptr
	pg_QGraphicsObject_delete func(uintptr)
)

func init() {
	reg(&pg_QGraphicsObject_base, "QGraphicsObject_base")
	symaddr(&addr_QGraphicsObject_tr, "QGraphicsObject_tr")
	symaddr(&addr_QGraphicsObject_grabGesture, "QGraphicsObject_grabGesture")
	symaddr(&addr_QGraphicsObject_ungrabGesture, "QGraphicsObject_ungrabGesture")
	symaddr(&addr_QGraphicsObject_updateMicroFocus, "QGraphicsObject_updateMicroFocus")
	symaddr(&addr_QGraphicsObject_parentChanged, "QGraphicsObject_parentChanged")
	symaddr(&addr_QGraphicsObject_opacityChanged, "QGraphicsObject_opacityChanged")
	symaddr(&addr_QGraphicsObject_visibleChanged, "QGraphicsObject_visibleChanged")
	symaddr(&addr_QGraphicsObject_enabledChanged, "QGraphicsObject_enabledChanged")
	symaddr(&addr_QGraphicsObject_xChanged, "QGraphicsObject_xChanged")
	symaddr(&addr_QGraphicsObject_yChanged, "QGraphicsObject_yChanged")
	symaddr(&addr_QGraphicsObject_zChanged, "QGraphicsObject_zChanged")
	symaddr(&addr_QGraphicsObject_rotationChanged, "QGraphicsObject_rotationChanged")
	symaddr(&addr_QGraphicsObject_scaleChanged, "QGraphicsObject_scaleChanged")
	symaddr(&addr_QGraphicsObject_childrenChanged, "QGraphicsObject_childrenChanged")
	symaddr(&addr_QGraphicsObject_widthChanged, "QGraphicsObject_widthChanged")
	symaddr(&addr_QGraphicsObject_heightChanged, "QGraphicsObject_heightChanged")
	symaddr(&addr_QGraphicsObject_connect, "QGraphicsObject_connect")
	symaddr(&addr_QGraphicsObject_connect2, "QGraphicsObject_connect2")
	symaddr(&addr_QGraphicsObject_sender, "QGraphicsObject_sender")
	symaddr(&addr_QGraphicsObject_senderSignalIndex, "QGraphicsObject_senderSignalIndex")
	symaddr(&addr_QGraphicsObject_receivers, "QGraphicsObject_receivers")
	symaddr(&addr_QGraphicsObject_isSignalConnected, "QGraphicsObject_isSignalConnected")
	symaddr(&addr_QGraphicsObject_scene, "QGraphicsObject_scene")
	symaddr(&addr_QGraphicsObject_parentItem, "QGraphicsObject_parentItem")
	symaddr(&addr_QGraphicsObject_topLevelItem, "QGraphicsObject_topLevelItem")
	symaddr(&addr_QGraphicsObject_parentObject, "QGraphicsObject_parentObject")
	symaddr(&addr_QGraphicsObject_parentWidget, "QGraphicsObject_parentWidget")
	symaddr(&addr_QGraphicsObject_topLevelWidget, "QGraphicsObject_topLevelWidget")
	symaddr(&addr_QGraphicsObject_window, "QGraphicsObject_window")
	symaddr(&addr_QGraphicsObject_panel, "QGraphicsObject_panel")
	symaddr(&addr_QGraphicsObject_setParentItem, "QGraphicsObject_setParentItem")
	symaddr(&addr_QGraphicsObject_childItems, "QGraphicsObject_childItems")
	symaddr(&addr_QGraphicsObject_isWidget, "QGraphicsObject_isWidget")
	symaddr(&addr_QGraphicsObject_isWindow, "QGraphicsObject_isWindow")
	symaddr(&addr_QGraphicsObject_isPanel, "QGraphicsObject_isPanel")
	symaddr(&addr_QGraphicsObject_toGraphicsObject, "QGraphicsObject_toGraphicsObject")
	symaddr(&addr_QGraphicsObject_toGraphicsObject2, "QGraphicsObject_toGraphicsObject2")
	symaddr(&addr_QGraphicsObject_group, "QGraphicsObject_group")
	symaddr(&addr_QGraphicsObject_setGroup, "QGraphicsObject_setGroup")
	symaddr(&addr_QGraphicsObject_flags, "QGraphicsObject_flags")
	symaddr(&addr_QGraphicsObject_setFlag, "QGraphicsObject_setFlag")
	symaddr(&addr_QGraphicsObject_setFlags, "QGraphicsObject_setFlags")
	symaddr(&addr_QGraphicsObject_cacheMode, "QGraphicsObject_cacheMode")
	symaddr(&addr_QGraphicsObject_setCacheMode, "QGraphicsObject_setCacheMode")
	symaddr(&addr_QGraphicsObject_panelModality, "QGraphicsObject_panelModality")
	symaddr(&addr_QGraphicsObject_setPanelModality, "QGraphicsObject_setPanelModality")
	symaddr(&addr_QGraphicsObject_isBlockedByModalPanel, "QGraphicsObject_isBlockedByModalPanel")
	symaddr(&addr_QGraphicsObject_toolTip, "QGraphicsObject_toolTip")
	symaddr(&addr_QGraphicsObject_setToolTip, "QGraphicsObject_setToolTip")
	symaddr(&addr_QGraphicsObject_cursor, "QGraphicsObject_cursor")
	symaddr(&addr_QGraphicsObject_setCursor, "QGraphicsObject_setCursor")
	symaddr(&addr_QGraphicsObject_hasCursor, "QGraphicsObject_hasCursor")
	symaddr(&addr_QGraphicsObject_unsetCursor, "QGraphicsObject_unsetCursor")
	symaddr(&addr_QGraphicsObject_isVisible, "QGraphicsObject_isVisible")
	symaddr(&addr_QGraphicsObject_isVisibleTo, "QGraphicsObject_isVisibleTo")
	symaddr(&addr_QGraphicsObject_setVisible, "QGraphicsObject_setVisible")
	symaddr(&addr_QGraphicsObject_hide, "QGraphicsObject_hide")
	symaddr(&addr_QGraphicsObject_show, "QGraphicsObject_show")
	symaddr(&addr_QGraphicsObject_isEnabled, "QGraphicsObject_isEnabled")
	symaddr(&addr_QGraphicsObject_setEnabled, "QGraphicsObject_setEnabled")
	symaddr(&addr_QGraphicsObject_isSelected, "QGraphicsObject_isSelected")
	symaddr(&addr_QGraphicsObject_setSelected, "QGraphicsObject_setSelected")
	symaddr(&addr_QGraphicsObject_acceptDrops, "QGraphicsObject_acceptDrops")
	symaddr(&addr_QGraphicsObject_setAcceptDrops, "QGraphicsObject_setAcceptDrops")
	reg(&pg_QGraphicsObject_opacity, "QGraphicsObject_opacity")
	reg(&pg_QGraphicsObject_effectiveOpacity, "QGraphicsObject_effectiveOpacity")
	reg(&pg_QGraphicsObject_setOpacity, "QGraphicsObject_setOpacity")
	symaddr(&addr_QGraphicsObject_graphicsEffect, "QGraphicsObject_graphicsEffect")
	symaddr(&addr_QGraphicsObject_setGraphicsEffect, "QGraphicsObject_setGraphicsEffect")
	symaddr(&addr_QGraphicsObject_acceptedMouseButtons, "QGraphicsObject_acceptedMouseButtons")
	symaddr(&addr_QGraphicsObject_setAcceptedMouseButtons, "QGraphicsObject_setAcceptedMouseButtons")
	symaddr(&addr_QGraphicsObject_acceptHoverEvents, "QGraphicsObject_acceptHoverEvents")
	symaddr(&addr_QGraphicsObject_setAcceptHoverEvents, "QGraphicsObject_setAcceptHoverEvents")
	symaddr(&addr_QGraphicsObject_acceptTouchEvents, "QGraphicsObject_acceptTouchEvents")
	symaddr(&addr_QGraphicsObject_setAcceptTouchEvents, "QGraphicsObject_setAcceptTouchEvents")
	symaddr(&addr_QGraphicsObject_filtersChildEvents, "QGraphicsObject_filtersChildEvents")
	symaddr(&addr_QGraphicsObject_setFiltersChildEvents, "QGraphicsObject_setFiltersChildEvents")
	symaddr(&addr_QGraphicsObject_handlesChildEvents, "QGraphicsObject_handlesChildEvents")
	symaddr(&addr_QGraphicsObject_setHandlesChildEvents, "QGraphicsObject_setHandlesChildEvents")
	symaddr(&addr_QGraphicsObject_isActive, "QGraphicsObject_isActive")
	symaddr(&addr_QGraphicsObject_setActive, "QGraphicsObject_setActive")
	symaddr(&addr_QGraphicsObject_hasFocus, "QGraphicsObject_hasFocus")
	symaddr(&addr_QGraphicsObject_setFocus, "QGraphicsObject_setFocus")
	symaddr(&addr_QGraphicsObject_clearFocus, "QGraphicsObject_clearFocus")
	symaddr(&addr_QGraphicsObject_focusProxy, "QGraphicsObject_focusProxy")
	symaddr(&addr_QGraphicsObject_setFocusProxy, "QGraphicsObject_setFocusProxy")
	symaddr(&addr_QGraphicsObject_focusItem, "QGraphicsObject_focusItem")
	symaddr(&addr_QGraphicsObject_focusScopeItem, "QGraphicsObject_focusScopeItem")
	symaddr(&addr_QGraphicsObject_grabMouse, "QGraphicsObject_grabMouse")
	symaddr(&addr_QGraphicsObject_ungrabMouse, "QGraphicsObject_ungrabMouse")
	symaddr(&addr_QGraphicsObject_grabKeyboard, "QGraphicsObject_grabKeyboard")
	symaddr(&addr_QGraphicsObject_ungrabKeyboard, "QGraphicsObject_ungrabKeyboard")
	symaddr(&addr_QGraphicsObject_pos, "QGraphicsObject_pos")
	reg(&pg_QGraphicsObject_x, "QGraphicsObject_x")
	reg(&pg_QGraphicsObject_setX, "QGraphicsObject_setX")
	reg(&pg_QGraphicsObject_y, "QGraphicsObject_y")
	reg(&pg_QGraphicsObject_setY, "QGraphicsObject_setY")
	symaddr(&addr_QGraphicsObject_scenePos, "QGraphicsObject_scenePos")
	symaddr(&addr_QGraphicsObject_setPos, "QGraphicsObject_setPos")
	reg(&pg_QGraphicsObject_setPos2, "QGraphicsObject_setPos2")
	reg(&pg_QGraphicsObject_moveBy, "QGraphicsObject_moveBy")
	symaddr(&addr_QGraphicsObject_ensureVisible, "QGraphicsObject_ensureVisible")
	reg(&pg_QGraphicsObject_ensureVisible2, "QGraphicsObject_ensureVisible2")
	symaddr(&addr_QGraphicsObject_transform, "QGraphicsObject_transform")
	symaddr(&addr_QGraphicsObject_sceneTransform, "QGraphicsObject_sceneTransform")
	symaddr(&addr_QGraphicsObject_deviceTransform, "QGraphicsObject_deviceTransform")
	symaddr(&addr_QGraphicsObject_itemTransform, "QGraphicsObject_itemTransform")
	symaddr(&addr_QGraphicsObject_setTransform, "QGraphicsObject_setTransform")
	symaddr(&addr_QGraphicsObject_resetTransform, "QGraphicsObject_resetTransform")
	reg(&pg_QGraphicsObject_setRotation, "QGraphicsObject_setRotation")
	reg(&pg_QGraphicsObject_rotation, "QGraphicsObject_rotation")
	reg(&pg_QGraphicsObject_setScale, "QGraphicsObject_setScale")
	reg(&pg_QGraphicsObject_scale, "QGraphicsObject_scale")
	symaddr(&addr_QGraphicsObject_transformations, "QGraphicsObject_transformations")
	symaddr(&addr_QGraphicsObject_setTransformations, "QGraphicsObject_setTransformations")
	symaddr(&addr_QGraphicsObject_transformOriginPoint, "QGraphicsObject_transformOriginPoint")
	symaddr(&addr_QGraphicsObject_setTransformOriginPoint, "QGraphicsObject_setTransformOriginPoint")
	reg(&pg_QGraphicsObject_setTransformOriginPoint2, "QGraphicsObject_setTransformOriginPoint2")
	symaddr(&addr_QGraphicsObject_advance, "QGraphicsObject_advance")
	reg(&pg_QGraphicsObject_zValue, "QGraphicsObject_zValue")
	reg(&pg_QGraphicsObject_setZValue, "QGraphicsObject_setZValue")
	symaddr(&addr_QGraphicsObject_stackBefore, "QGraphicsObject_stackBefore")
	symaddr(&addr_QGraphicsObject_boundingRect, "QGraphicsObject_boundingRect")
	symaddr(&addr_QGraphicsObject_childrenBoundingRect, "QGraphicsObject_childrenBoundingRect")
	symaddr(&addr_QGraphicsObject_sceneBoundingRect, "QGraphicsObject_sceneBoundingRect")
	symaddr(&addr_QGraphicsObject_shape, "QGraphicsObject_shape")
	symaddr(&addr_QGraphicsObject_isClipped, "QGraphicsObject_isClipped")
	symaddr(&addr_QGraphicsObject_clipPath, "QGraphicsObject_clipPath")
	symaddr(&addr_QGraphicsObject_contains, "QGraphicsObject_contains")
	symaddr(&addr_QGraphicsObject_collidesWithItem, "QGraphicsObject_collidesWithItem")
	symaddr(&addr_QGraphicsObject_collidesWithPath, "QGraphicsObject_collidesWithPath")
	symaddr(&addr_QGraphicsObject_collidingItems, "QGraphicsObject_collidingItems")
	symaddr(&addr_QGraphicsObject_isObscured, "QGraphicsObject_isObscured")
	reg(&pg_QGraphicsObject_isObscured2, "QGraphicsObject_isObscured2")
	symaddr(&addr_QGraphicsObject_isObscuredBy, "QGraphicsObject_isObscuredBy")
	symaddr(&addr_QGraphicsObject_opaqueArea, "QGraphicsObject_opaqueArea")
	symaddr(&addr_QGraphicsObject_boundingRegion, "QGraphicsObject_boundingRegion")
	reg(&pg_QGraphicsObject_boundingRegionGranularity, "QGraphicsObject_boundingRegionGranularity")
	reg(&pg_QGraphicsObject_setBoundingRegionGranularity, "QGraphicsObject_setBoundingRegionGranularity")
	symaddr(&addr_QGraphicsObject_paint, "QGraphicsObject_paint")
	symaddr(&addr_QGraphicsObject_update, "QGraphicsObject_update")
	reg(&pg_QGraphicsObject_update2, "QGraphicsObject_update2")
	reg(&pg_QGraphicsObject_scroll, "QGraphicsObject_scroll")
	symaddr(&addr_QGraphicsObject_mapToItem, "QGraphicsObject_mapToItem")
	symaddr(&addr_QGraphicsObject_mapToParent, "QGraphicsObject_mapToParent")
	symaddr(&addr_QGraphicsObject_mapToScene, "QGraphicsObject_mapToScene")
	symaddr(&addr_QGraphicsObject_mapToItem2, "QGraphicsObject_mapToItem2")
	symaddr(&addr_QGraphicsObject_mapToParent2, "QGraphicsObject_mapToParent2")
	symaddr(&addr_QGraphicsObject_mapToScene2, "QGraphicsObject_mapToScene2")
	symaddr(&addr_QGraphicsObject_mapRectToItem, "QGraphicsObject_mapRectToItem")
	symaddr(&addr_QGraphicsObject_mapRectToParent, "QGraphicsObject_mapRectToParent")
	symaddr(&addr_QGraphicsObject_mapRectToScene, "QGraphicsObject_mapRectToScene")
	symaddr(&addr_QGraphicsObject_mapToItem3, "QGraphicsObject_mapToItem3")
	symaddr(&addr_QGraphicsObject_mapToParent3, "QGraphicsObject_mapToParent3")
	symaddr(&addr_QGraphicsObject_mapToScene3, "QGraphicsObject_mapToScene3")
	symaddr(&addr_QGraphicsObject_mapToItem4, "QGraphicsObject_mapToItem4")
	symaddr(&addr_QGraphicsObject_mapToParent4, "QGraphicsObject_mapToParent4")
	symaddr(&addr_QGraphicsObject_mapToScene4, "QGraphicsObject_mapToScene4")
	symaddr(&addr_QGraphicsObject_mapFromItem, "QGraphicsObject_mapFromItem")
	symaddr(&addr_QGraphicsObject_mapFromParent, "QGraphicsObject_mapFromParent")
	symaddr(&addr_QGraphicsObject_mapFromScene, "QGraphicsObject_mapFromScene")
	symaddr(&addr_QGraphicsObject_mapFromItem2, "QGraphicsObject_mapFromItem2")
	symaddr(&addr_QGraphicsObject_mapFromParent2, "QGraphicsObject_mapFromParent2")
	symaddr(&addr_QGraphicsObject_mapFromScene2, "QGraphicsObject_mapFromScene2")
	symaddr(&addr_QGraphicsObject_mapRectFromItem, "QGraphicsObject_mapRectFromItem")
	symaddr(&addr_QGraphicsObject_mapRectFromParent, "QGraphicsObject_mapRectFromParent")
	symaddr(&addr_QGraphicsObject_mapRectFromScene, "QGraphicsObject_mapRectFromScene")
	symaddr(&addr_QGraphicsObject_mapFromItem3, "QGraphicsObject_mapFromItem3")
	symaddr(&addr_QGraphicsObject_mapFromParent3, "QGraphicsObject_mapFromParent3")
	symaddr(&addr_QGraphicsObject_mapFromScene3, "QGraphicsObject_mapFromScene3")
	symaddr(&addr_QGraphicsObject_mapFromItem4, "QGraphicsObject_mapFromItem4")
	symaddr(&addr_QGraphicsObject_mapFromParent4, "QGraphicsObject_mapFromParent4")
	symaddr(&addr_QGraphicsObject_mapFromScene4, "QGraphicsObject_mapFromScene4")
	reg(&pg_QGraphicsObject_mapToItem5, "QGraphicsObject_mapToItem5")
	reg(&pg_QGraphicsObject_mapToParent5, "QGraphicsObject_mapToParent5")
	reg(&pg_QGraphicsObject_mapToScene5, "QGraphicsObject_mapToScene5")
	reg(&pg_QGraphicsObject_mapToItem6, "QGraphicsObject_mapToItem6")
	reg(&pg_QGraphicsObject_mapToParent6, "QGraphicsObject_mapToParent6")
	reg(&pg_QGraphicsObject_mapToScene6, "QGraphicsObject_mapToScene6")
	reg(&pg_QGraphicsObject_mapRectToItem2, "QGraphicsObject_mapRectToItem2")
	reg(&pg_QGraphicsObject_mapRectToParent2, "QGraphicsObject_mapRectToParent2")
	reg(&pg_QGraphicsObject_mapRectToScene2, "QGraphicsObject_mapRectToScene2")
	reg(&pg_QGraphicsObject_mapFromItem5, "QGraphicsObject_mapFromItem5")
	reg(&pg_QGraphicsObject_mapFromParent5, "QGraphicsObject_mapFromParent5")
	reg(&pg_QGraphicsObject_mapFromScene5, "QGraphicsObject_mapFromScene5")
	reg(&pg_QGraphicsObject_mapFromItem6, "QGraphicsObject_mapFromItem6")
	reg(&pg_QGraphicsObject_mapFromParent6, "QGraphicsObject_mapFromParent6")
	reg(&pg_QGraphicsObject_mapFromScene6, "QGraphicsObject_mapFromScene6")
	reg(&pg_QGraphicsObject_mapRectFromItem2, "QGraphicsObject_mapRectFromItem2")
	reg(&pg_QGraphicsObject_mapRectFromParent2, "QGraphicsObject_mapRectFromParent2")
	reg(&pg_QGraphicsObject_mapRectFromScene2, "QGraphicsObject_mapRectFromScene2")
	symaddr(&addr_QGraphicsObject_isAncestorOf, "QGraphicsObject_isAncestorOf")
	symaddr(&addr_QGraphicsObject_commonAncestorItem, "QGraphicsObject_commonAncestorItem")
	symaddr(&addr_QGraphicsObject_isUnderMouse, "QGraphicsObject_isUnderMouse")
	symaddr(&addr_QGraphicsObject_data, "QGraphicsObject_data")
	symaddr(&addr_QGraphicsObject_setData, "QGraphicsObject_setData")
	symaddr(&addr_QGraphicsObject_inputMethodHints, "QGraphicsObject_inputMethodHints")
	symaddr(&addr_QGraphicsObject_setInputMethodHints, "QGraphicsObject_setInputMethodHints")
	symaddr(&addr_QGraphicsObject_type, "QGraphicsObject_type")
	symaddr(&addr_QGraphicsObject_installSceneEventFilter, "QGraphicsObject_installSceneEventFilter")
	symaddr(&addr_QGraphicsObject_removeSceneEventFilter, "QGraphicsObject_removeSceneEventFilter")
	symaddr(&addr_QGraphicsObject_addToIndex, "QGraphicsObject_addToIndex")
	symaddr(&addr_QGraphicsObject_removeFromIndex, "QGraphicsObject_removeFromIndex")
	symaddr(&addr_QGraphicsObject_prepareGeometryChange, "QGraphicsObject_prepareGeometryChange")
	symaddr(&addr_QGraphicsObject_tr2, "QGraphicsObject_tr2")
	symaddr(&addr_QGraphicsObject_tr3, "QGraphicsObject_tr3")
	symaddr(&addr_QGraphicsObject_grabGesture2, "QGraphicsObject_grabGesture2")
	symaddr(&addr_QGraphicsObject_connect3, "QGraphicsObject_connect3")
	symaddr(&addr_QGraphicsObject_connect4, "QGraphicsObject_connect4")
	symaddr(&addr_QGraphicsObject_setFlag2, "QGraphicsObject_setFlag2")
	symaddr(&addr_QGraphicsObject_setCacheMode2, "QGraphicsObject_setCacheMode2")
	symaddr(&addr_QGraphicsObject_setFocus2, "QGraphicsObject_setFocus2")
	symaddr(&addr_QGraphicsObject_ensureVisible3, "QGraphicsObject_ensureVisible3")
	symaddr(&addr_QGraphicsObject_ensureVisible4, "QGraphicsObject_ensureVisible4")
	symaddr(&addr_QGraphicsObject_ensureVisible5, "QGraphicsObject_ensureVisible5")
	reg(&pg_QGraphicsObject_ensureVisible6, "QGraphicsObject_ensureVisible6")
	reg(&pg_QGraphicsObject_ensureVisible7, "QGraphicsObject_ensureVisible7")
	symaddr(&addr_QGraphicsObject_itemTransform2, "QGraphicsObject_itemTransform2")
	symaddr(&addr_QGraphicsObject_setTransform2, "QGraphicsObject_setTransform2")
	symaddr(&addr_QGraphicsObject_collidingItems2, "QGraphicsObject_collidingItems2")
	symaddr(&addr_QGraphicsObject_isObscured3, "QGraphicsObject_isObscured3")
	symaddr(&addr_QGraphicsObject_update3, "QGraphicsObject_update3")
	reg(&pg_QGraphicsObject_scroll2, "QGraphicsObject_scroll2")
	symaddr(&addr_QGraphicsObject_asQGraphicsItem, "QGraphicsObject_asQGraphicsItem")
	symaddr(&addr_QGraphicsObject_advance_default, "QGraphicsObject_advance_default")
	symaddr(&addr_QGraphicsObject_shape_default, "QGraphicsObject_shape_default")
	symaddr(&addr_QGraphicsObject_contains_default, "QGraphicsObject_contains_default")
	symaddr(&addr_QGraphicsObject_collidesWithItem_default, "QGraphicsObject_collidesWithItem_default")
	symaddr(&addr_QGraphicsObject_collidesWithPath_default, "QGraphicsObject_collidesWithPath_default")
	symaddr(&addr_QGraphicsObject_isObscuredBy_default, "QGraphicsObject_isObscuredBy_default")
	symaddr(&addr_QGraphicsObject_opaqueArea_default, "QGraphicsObject_opaqueArea_default")
	symaddr(&addr_QGraphicsObject_type_default, "QGraphicsObject_type_default")
	symaddr(&addr_QGraphicsObject_metaObject_default, "QGraphicsObject_metaObject_default")
	symaddr(&addr_QGraphicsObject_event_default, "QGraphicsObject_event_default")
	symaddr(&addr_QGraphicsObject_eventFilter_default, "QGraphicsObject_eventFilter_default")
	symaddr(&addr_QGraphicsObject_timerEvent_default, "QGraphicsObject_timerEvent_default")
	symaddr(&addr_QGraphicsObject_childEvent_default, "QGraphicsObject_childEvent_default")
	symaddr(&addr_QGraphicsObject_customEvent_default, "QGraphicsObject_customEvent_default")
	symaddr(&addr_QGraphicsObject_connectNotify_default, "QGraphicsObject_connectNotify_default")
	symaddr(&addr_QGraphicsObject_disconnectNotify_default, "QGraphicsObject_disconnectNotify_default")
	symaddr(&addr_QGraphicsObject_sceneEventFilter_default, "QGraphicsObject_sceneEventFilter_default")
	symaddr(&addr_QGraphicsObject_sceneEvent_default, "QGraphicsObject_sceneEvent_default")
	symaddr(&addr_QGraphicsObject_contextMenuEvent_default, "QGraphicsObject_contextMenuEvent_default")
	symaddr(&addr_QGraphicsObject_dragEnterEvent_default, "QGraphicsObject_dragEnterEvent_default")
	symaddr(&addr_QGraphicsObject_dragLeaveEvent_default, "QGraphicsObject_dragLeaveEvent_default")
	symaddr(&addr_QGraphicsObject_dragMoveEvent_default, "QGraphicsObject_dragMoveEvent_default")
	symaddr(&addr_QGraphicsObject_dropEvent_default, "QGraphicsObject_dropEvent_default")
	symaddr(&addr_QGraphicsObject_focusInEvent_default, "QGraphicsObject_focusInEvent_default")
	symaddr(&addr_QGraphicsObject_focusOutEvent_default, "QGraphicsObject_focusOutEvent_default")
	symaddr(&addr_QGraphicsObject_hoverEnterEvent_default, "QGraphicsObject_hoverEnterEvent_default")
	symaddr(&addr_QGraphicsObject_hoverMoveEvent_default, "QGraphicsObject_hoverMoveEvent_default")
	symaddr(&addr_QGraphicsObject_hoverLeaveEvent_default, "QGraphicsObject_hoverLeaveEvent_default")
	symaddr(&addr_QGraphicsObject_keyPressEvent_default, "QGraphicsObject_keyPressEvent_default")
	symaddr(&addr_QGraphicsObject_keyReleaseEvent_default, "QGraphicsObject_keyReleaseEvent_default")
	symaddr(&addr_QGraphicsObject_mousePressEvent_default, "QGraphicsObject_mousePressEvent_default")
	symaddr(&addr_QGraphicsObject_mouseMoveEvent_default, "QGraphicsObject_mouseMoveEvent_default")
	symaddr(&addr_QGraphicsObject_mouseReleaseEvent_default, "QGraphicsObject_mouseReleaseEvent_default")
	symaddr(&addr_QGraphicsObject_mouseDoubleClickEvent_default, "QGraphicsObject_mouseDoubleClickEvent_default")
	symaddr(&addr_QGraphicsObject_wheelEvent_default, "QGraphicsObject_wheelEvent_default")
	symaddr(&addr_QGraphicsObject_inputMethodEvent_default, "QGraphicsObject_inputMethodEvent_default")
	symaddr(&addr_QGraphicsObject_inputMethodQuery_default, "QGraphicsObject_inputMethodQuery_default")
	symaddr(&addr_QGraphicsObject_itemChange_default, "QGraphicsObject_itemChange_default")
	symaddr(&addr_QGraphicsObject_supportsExtension_default, "QGraphicsObject_supportsExtension_default")
	symaddr(&addr_QGraphicsObject_setExtension_default, "QGraphicsObject_setExtension_default")
	symaddr(&addr_QGraphicsObject_extension_default, "QGraphicsObject_extension_default")
	reg(&pg_QGraphicsObject_connect_parentChanged, "QGraphicsObject_connect_parentChanged")
	reg(&pg_QGraphicsObject_connect_opacityChanged, "QGraphicsObject_connect_opacityChanged")
	reg(&pg_QGraphicsObject_connect_visibleChanged, "QGraphicsObject_connect_visibleChanged")
	reg(&pg_QGraphicsObject_connect_enabledChanged, "QGraphicsObject_connect_enabledChanged")
	reg(&pg_QGraphicsObject_connect_xChanged, "QGraphicsObject_connect_xChanged")
	reg(&pg_QGraphicsObject_connect_yChanged, "QGraphicsObject_connect_yChanged")
	reg(&pg_QGraphicsObject_connect_zChanged, "QGraphicsObject_connect_zChanged")
	reg(&pg_QGraphicsObject_connect_rotationChanged, "QGraphicsObject_connect_rotationChanged")
	reg(&pg_QGraphicsObject_connect_scaleChanged, "QGraphicsObject_connect_scaleChanged")
	reg(&pg_QGraphicsObject_connect_childrenChanged, "QGraphicsObject_connect_childrenChanged")
	reg(&pg_QGraphicsObject_connect_widthChanged, "QGraphicsObject_connect_widthChanged")
	reg(&pg_QGraphicsObject_connect_heightChanged, "QGraphicsObject_connect_heightChanged")
	reg(&pg_QGraphicsObject_connect_destroyed, "QGraphicsObject_connect_destroyed")
	reg(&pg_QGraphicsObject_connect_destroyed2, "QGraphicsObject_connect_destroyed2")
	reg(&pg_QGraphicsObject_connect_objectNameChanged, "QGraphicsObject_connect_objectNameChanged")
	reg(&pg_QGraphicsObject__disconnect, "QGraphicsObject__disconnect")
	reg(&pg_QGraphicsObject__set_override, "QGraphicsObject__set_override")
	reg(&pg_QGraphicsObject__Go_new, "QGraphicsObject__Go_new")
	reg(&pg_QGraphicsObject_delete, "QGraphicsObject_delete")
}

type QGraphicsObject struct {
	h uintptr
	*QObject
}

func bindQGraphicsObject(h uintptr) *QGraphicsObject {
	if h == 0 {
		return nil
	}
	return &QGraphicsObject{h: h, QObject: bindQObject(pg_QGraphicsObject_base(h))}
}

func newownedQGraphicsObject(h uintptr) *QGraphicsObject { return bindQGraphicsObject(h) }

func (this *QGraphicsObject) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QGraphicsObject) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQGraphicsObject(h unsafe.Pointer) *QGraphicsObject { return bindQGraphicsObject(uintptr(h)) }
func UnsafeNewQGraphicsObjectOwned(h unsafe.Pointer) *QGraphicsObject { return newownedQGraphicsObject(uintptr(h)) }

func (this *QGraphicsObject) UnsafeRelease() {
	this.h = 0
	if this.QObject != nil {
		this.QObject.UnsafeRelease()
	}
}

func QGraphicsObject_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGraphicsObject_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QGraphicsObject) GrabGesture(type_ GestureType) {
	purego.SyscallN(addr_QGraphicsObject_grabGesture, this.h, uintptr(type_))
	runtime.KeepAlive(this)
}

func (this *QGraphicsObject) UngrabGesture(type_ GestureType) {
	purego.SyscallN(addr_QGraphicsObject_ungrabGesture, this.h, uintptr(type_))
	runtime.KeepAlive(this)
}

func (this *QGraphicsObject) UpdateMicroFocus() {
	purego.SyscallN(addr_QGraphicsObject_updateMicroFocus, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsObject) ParentChanged() {
	purego.SyscallN(addr_QGraphicsObject_parentChanged, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsObject) OpacityChanged() {
	purego.SyscallN(addr_QGraphicsObject_opacityChanged, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsObject) VisibleChanged() {
	purego.SyscallN(addr_QGraphicsObject_visibleChanged, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsObject) EnabledChanged() {
	purego.SyscallN(addr_QGraphicsObject_enabledChanged, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsObject) XChanged() {
	purego.SyscallN(addr_QGraphicsObject_xChanged, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsObject) YChanged() {
	purego.SyscallN(addr_QGraphicsObject_yChanged, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsObject) ZChanged() {
	purego.SyscallN(addr_QGraphicsObject_zChanged, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsObject) RotationChanged() {
	purego.SyscallN(addr_QGraphicsObject_rotationChanged, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsObject) ScaleChanged() {
	purego.SyscallN(addr_QGraphicsObject_scaleChanged, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsObject) ChildrenChanged() {
	purego.SyscallN(addr_QGraphicsObject_childrenChanged, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsObject) WidthChanged() {
	purego.SyscallN(addr_QGraphicsObject_widthChanged, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsObject) HeightChanged() {
	purego.SyscallN(addr_QGraphicsObject_heightChanged, this.h)
	runtime.KeepAlive(this)
}

func QGraphicsObject_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QGraphicsObject_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QGraphicsObject_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QGraphicsObject_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QGraphicsObject) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QGraphicsObject_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QGraphicsObject) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsObject_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QGraphicsObject) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QGraphicsObject_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QGraphicsObject) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsObject_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func (this *QGraphicsObject) Scene() *QGraphicsScene {
	__r, _, _ := purego.SyscallN(addr_QGraphicsObject_scene, this.h)
	runtime.KeepAlive(this)
	return bindQGraphicsScene(__r)
}

func (this *QGraphicsObject) ParentItem() *QGraphicsItem {
	__r, _, _ := purego.SyscallN(addr_QGraphicsObject_parentItem, this.h)
	runtime.KeepAlive(this)
	return bindQGraphicsItem(__r)
}

func (this *QGraphicsObject) TopLevelItem() *QGraphicsItem {
	__r, _, _ := purego.SyscallN(addr_QGraphicsObject_topLevelItem, this.h)
	runtime.KeepAlive(this)
	return bindQGraphicsItem(__r)
}

func (this *QGraphicsObject) ParentObject() *QGraphicsObject {
	__r, _, _ := purego.SyscallN(addr_QGraphicsObject_parentObject, this.h)
	runtime.KeepAlive(this)
	return bindQGraphicsObject(__r)
}

func (this *QGraphicsObject) ParentWidget() *QGraphicsWidget {
	__r, _, _ := purego.SyscallN(addr_QGraphicsObject_parentWidget, this.h)
	runtime.KeepAlive(this)
	return bindQGraphicsWidget(__r)
}

func (this *QGraphicsObject) TopLevelWidget() *QGraphicsWidget {
	__r, _, _ := purego.SyscallN(addr_QGraphicsObject_topLevelWidget, this.h)
	runtime.KeepAlive(this)
	return bindQGraphicsWidget(__r)
}

func (this *QGraphicsObject) Window() *QGraphicsWidget {
	__r, _, _ := purego.SyscallN(addr_QGraphicsObject_window, this.h)
	runtime.KeepAlive(this)
	return bindQGraphicsWidget(__r)
}

func (this *QGraphicsObject) Panel() *QGraphicsItem {
	__r, _, _ := purego.SyscallN(addr_QGraphicsObject_panel, this.h)
	runtime.KeepAlive(this)
	return bindQGraphicsItem(__r)
}

func (this *QGraphicsObject) SetParentItem(parent *QGraphicsItem) {
	purego.SyscallN(addr_QGraphicsObject_setParentItem, this.h, uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
}

func (this *QGraphicsObject) ChildItems() []*QGraphicsItem {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGraphicsObject_childItems, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QGraphicsItem, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = bindQGraphicsItem(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QGraphicsObject) IsWidget() bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsObject_isWidget, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGraphicsObject) IsWindow() bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsObject_isWindow, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGraphicsObject) IsPanel() bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsObject_isPanel, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGraphicsObject) ToGraphicsObject() *QGraphicsObject {
	__r, _, _ := purego.SyscallN(addr_QGraphicsObject_toGraphicsObject, this.h)
	runtime.KeepAlive(this)
	return bindQGraphicsObject(__r)
}

func (this *QGraphicsObject) ToGraphicsObject2() *QGraphicsObject {
	__r, _, _ := purego.SyscallN(addr_QGraphicsObject_toGraphicsObject2, this.h)
	runtime.KeepAlive(this)
	return bindQGraphicsObject(__r)
}

func (this *QGraphicsObject) Group() *QGraphicsItemGroup {
	__r, _, _ := purego.SyscallN(addr_QGraphicsObject_group, this.h)
	runtime.KeepAlive(this)
	return bindQGraphicsItemGroup(__r)
}

func (this *QGraphicsObject) SetGroup(group *QGraphicsItemGroup) {
	purego.SyscallN(addr_QGraphicsObject_setGroup, this.h, uintptr(group.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(group)
}

func (this *QGraphicsObject) Flags() QGraphicsItem__GraphicsItemFlag {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsObject_flags, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QGraphicsItem__GraphicsItemFlag(__r)
}

func (this *QGraphicsObject) SetFlag(flag QGraphicsItem__GraphicsItemFlag) {
	purego.SyscallN(addr_QGraphicsObject_setFlag, this.h, uintptr(flag))
	runtime.KeepAlive(this)
}

func (this *QGraphicsObject) SetFlags(flags QGraphicsItem__GraphicsItemFlag) {
	purego.SyscallN(addr_QGraphicsObject_setFlags, this.h, uintptr(flags))
	runtime.KeepAlive(this)
}

func (this *QGraphicsObject) CacheMode() QGraphicsItem__CacheMode {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsObject_cacheMode, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QGraphicsItem__CacheMode(__r)
}

func (this *QGraphicsObject) SetCacheMode(mode QGraphicsItem__CacheMode) {
	purego.SyscallN(addr_QGraphicsObject_setCacheMode, this.h, uintptr(mode))
	runtime.KeepAlive(this)
}

func (this *QGraphicsObject) PanelModality() QGraphicsItem__PanelModality {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsObject_panelModality, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QGraphicsItem__PanelModality(__r)
}

func (this *QGraphicsObject) SetPanelModality(panelModality QGraphicsItem__PanelModality) {
	purego.SyscallN(addr_QGraphicsObject_setPanelModality, this.h, uintptr(panelModality))
	runtime.KeepAlive(this)
}

func (this *QGraphicsObject) IsBlockedByModalPanel() bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsObject_isBlockedByModalPanel, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGraphicsObject) ToolTip() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGraphicsObject_toolTip, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QGraphicsObject) SetToolTip(toolTip string) {
	var toolTip_p uintptr
	if len(toolTip) > 0 {
		toolTip_p = uintptr(unsafe.Pointer(unsafe.StringData(toolTip)))
	}
	purego.SyscallN(addr_QGraphicsObject_setToolTip, this.h, toolTip_p, uintptr(len(toolTip)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(toolTip)
}

func (this *QGraphicsObject) Cursor() *QCursor {
	__r, _, _ := purego.SyscallN(addr_QGraphicsObject_cursor, this.h)
	runtime.KeepAlive(this)
	return newownedQCursor(__r)
}

func (this *QGraphicsObject) SetCursor(cursor *QCursor) {
	purego.SyscallN(addr_QGraphicsObject_setCursor, this.h, uintptr(cursor.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(cursor)
}

func (this *QGraphicsObject) HasCursor() bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsObject_hasCursor, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGraphicsObject) UnsetCursor() {
	purego.SyscallN(addr_QGraphicsObject_unsetCursor, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsObject) IsVisible() bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsObject_isVisible, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGraphicsObject) IsVisibleTo(parent *QGraphicsItem) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsObject_isVisibleTo, this.h, uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QGraphicsObject) SetVisible(visible bool) {
	purego.SyscallN(addr_QGraphicsObject_setVisible, this.h, uintptr(b2i(visible)))
	runtime.KeepAlive(this)
}

func (this *QGraphicsObject) Hide() {
	purego.SyscallN(addr_QGraphicsObject_hide, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsObject) Show() {
	purego.SyscallN(addr_QGraphicsObject_show, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsObject) IsEnabled() bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsObject_isEnabled, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGraphicsObject) SetEnabled(enabled bool) {
	purego.SyscallN(addr_QGraphicsObject_setEnabled, this.h, uintptr(b2i(enabled)))
	runtime.KeepAlive(this)
}

func (this *QGraphicsObject) IsSelected() bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsObject_isSelected, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGraphicsObject) SetSelected(selected bool) {
	purego.SyscallN(addr_QGraphicsObject_setSelected, this.h, uintptr(b2i(selected)))
	runtime.KeepAlive(this)
}

func (this *QGraphicsObject) AcceptDrops() bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsObject_acceptDrops, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGraphicsObject) SetAcceptDrops(on bool) {
	purego.SyscallN(addr_QGraphicsObject_setAcceptDrops, this.h, uintptr(b2i(on)))
	runtime.KeepAlive(this)
}

func (this *QGraphicsObject) Opacity() float64 {
	__r := pg_QGraphicsObject_opacity(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QGraphicsObject) EffectiveOpacity() float64 {
	__r := pg_QGraphicsObject_effectiveOpacity(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QGraphicsObject) SetOpacity(opacity float64) {
	pg_QGraphicsObject_setOpacity(this.h, opacity)
	runtime.KeepAlive(this)
}

func (this *QGraphicsObject) GraphicsEffect() *QGraphicsEffect {
	__r, _, _ := purego.SyscallN(addr_QGraphicsObject_graphicsEffect, this.h)
	runtime.KeepAlive(this)
	return bindQGraphicsEffect(__r)
}

func (this *QGraphicsObject) SetGraphicsEffect(effect *QGraphicsEffect) {
	purego.SyscallN(addr_QGraphicsObject_setGraphicsEffect, this.h, uintptr(effect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(effect)
}

func (this *QGraphicsObject) AcceptedMouseButtons() MouseButton {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsObject_acceptedMouseButtons, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return MouseButton(__r)
}

func (this *QGraphicsObject) SetAcceptedMouseButtons(buttons MouseButton) {
	purego.SyscallN(addr_QGraphicsObject_setAcceptedMouseButtons, this.h, uintptr(buttons))
	runtime.KeepAlive(this)
}

func (this *QGraphicsObject) AcceptHoverEvents() bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsObject_acceptHoverEvents, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGraphicsObject) SetAcceptHoverEvents(enabled bool) {
	purego.SyscallN(addr_QGraphicsObject_setAcceptHoverEvents, this.h, uintptr(b2i(enabled)))
	runtime.KeepAlive(this)
}

func (this *QGraphicsObject) AcceptTouchEvents() bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsObject_acceptTouchEvents, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGraphicsObject) SetAcceptTouchEvents(enabled bool) {
	purego.SyscallN(addr_QGraphicsObject_setAcceptTouchEvents, this.h, uintptr(b2i(enabled)))
	runtime.KeepAlive(this)
}

func (this *QGraphicsObject) FiltersChildEvents() bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsObject_filtersChildEvents, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGraphicsObject) SetFiltersChildEvents(enabled bool) {
	purego.SyscallN(addr_QGraphicsObject_setFiltersChildEvents, this.h, uintptr(b2i(enabled)))
	runtime.KeepAlive(this)
}

func (this *QGraphicsObject) HandlesChildEvents() bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsObject_handlesChildEvents, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGraphicsObject) SetHandlesChildEvents(enabled bool) {
	purego.SyscallN(addr_QGraphicsObject_setHandlesChildEvents, this.h, uintptr(b2i(enabled)))
	runtime.KeepAlive(this)
}

func (this *QGraphicsObject) IsActive() bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsObject_isActive, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGraphicsObject) SetActive(active bool) {
	purego.SyscallN(addr_QGraphicsObject_setActive, this.h, uintptr(b2i(active)))
	runtime.KeepAlive(this)
}

func (this *QGraphicsObject) HasFocus() bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsObject_hasFocus, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGraphicsObject) SetFocus() {
	purego.SyscallN(addr_QGraphicsObject_setFocus, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsObject) ClearFocus() {
	purego.SyscallN(addr_QGraphicsObject_clearFocus, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsObject) FocusProxy() *QGraphicsItem {
	__r, _, _ := purego.SyscallN(addr_QGraphicsObject_focusProxy, this.h)
	runtime.KeepAlive(this)
	return bindQGraphicsItem(__r)
}

func (this *QGraphicsObject) SetFocusProxy(item *QGraphicsItem) {
	purego.SyscallN(addr_QGraphicsObject_setFocusProxy, this.h, uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QGraphicsObject) FocusItem() *QGraphicsItem {
	__r, _, _ := purego.SyscallN(addr_QGraphicsObject_focusItem, this.h)
	runtime.KeepAlive(this)
	return bindQGraphicsItem(__r)
}

func (this *QGraphicsObject) FocusScopeItem() *QGraphicsItem {
	__r, _, _ := purego.SyscallN(addr_QGraphicsObject_focusScopeItem, this.h)
	runtime.KeepAlive(this)
	return bindQGraphicsItem(__r)
}

func (this *QGraphicsObject) GrabMouse() {
	purego.SyscallN(addr_QGraphicsObject_grabMouse, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsObject) UngrabMouse() {
	purego.SyscallN(addr_QGraphicsObject_ungrabMouse, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsObject) GrabKeyboard() {
	purego.SyscallN(addr_QGraphicsObject_grabKeyboard, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsObject) UngrabKeyboard() {
	purego.SyscallN(addr_QGraphicsObject_ungrabKeyboard, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsObject) Pos() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsObject_pos, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QGraphicsObject) X() float64 {
	__r := pg_QGraphicsObject_x(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QGraphicsObject) SetX(x float64) {
	pg_QGraphicsObject_setX(this.h, x)
	runtime.KeepAlive(this)
}

func (this *QGraphicsObject) Y() float64 {
	__r := pg_QGraphicsObject_y(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QGraphicsObject) SetY(y float64) {
	pg_QGraphicsObject_setY(this.h, y)
	runtime.KeepAlive(this)
}

func (this *QGraphicsObject) ScenePos() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsObject_scenePos, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QGraphicsObject) SetPos(pos *QPointF) {
	purego.SyscallN(addr_QGraphicsObject_setPos, this.h, uintptr(pos.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pos)
}

func (this *QGraphicsObject) SetPos2(x float64, y float64) {
	pg_QGraphicsObject_setPos2(this.h, x, y)
	runtime.KeepAlive(this)
}

func (this *QGraphicsObject) MoveBy(dx float64, dy float64) {
	pg_QGraphicsObject_moveBy(this.h, dx, dy)
	runtime.KeepAlive(this)
}

func (this *QGraphicsObject) EnsureVisible() {
	purego.SyscallN(addr_QGraphicsObject_ensureVisible, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsObject) EnsureVisible2(x float64, y float64, w float64, h float64) {
	pg_QGraphicsObject_ensureVisible2(this.h, x, y, w, h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsObject) Transform() *QTransform {
	__r, _, _ := purego.SyscallN(addr_QGraphicsObject_transform, this.h)
	runtime.KeepAlive(this)
	return newownedQTransform(__r)
}

func (this *QGraphicsObject) SceneTransform() *QTransform {
	__r, _, _ := purego.SyscallN(addr_QGraphicsObject_sceneTransform, this.h)
	runtime.KeepAlive(this)
	return newownedQTransform(__r)
}

func (this *QGraphicsObject) DeviceTransform(viewportTransform *QTransform) *QTransform {
	__r, _, _ := purego.SyscallN(addr_QGraphicsObject_deviceTransform, this.h, uintptr(viewportTransform.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(viewportTransform)
	return newownedQTransform(__r)
}

func (this *QGraphicsObject) ItemTransform(other *QGraphicsItem) *QTransform {
	__r, _, _ := purego.SyscallN(addr_QGraphicsObject_itemTransform, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return newownedQTransform(__r)
}

func (this *QGraphicsObject) SetTransform(matrix *QTransform) {
	purego.SyscallN(addr_QGraphicsObject_setTransform, this.h, uintptr(matrix.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(matrix)
}

func (this *QGraphicsObject) ResetTransform() {
	purego.SyscallN(addr_QGraphicsObject_resetTransform, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsObject) SetRotation(angle float64) {
	pg_QGraphicsObject_setRotation(this.h, angle)
	runtime.KeepAlive(this)
}

func (this *QGraphicsObject) Rotation() float64 {
	__r := pg_QGraphicsObject_rotation(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QGraphicsObject) SetScale(scale float64) {
	pg_QGraphicsObject_setScale(this.h, scale)
	runtime.KeepAlive(this)
}

func (this *QGraphicsObject) Scale() float64 {
	__r := pg_QGraphicsObject_scale(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QGraphicsObject) Transformations() []*QGraphicsTransform {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGraphicsObject_transformations, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QGraphicsTransform, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = bindQGraphicsTransform(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QGraphicsObject) SetTransformations(transformations []*QGraphicsTransform) {
	transformations_arr := make([]uintptr, len(transformations))
	for __i, __e := range transformations {
		transformations_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var transformations_p uintptr
	if len(transformations_arr) > 0 {
		transformations_p = uintptr(unsafe.Pointer(&transformations_arr[0]))
	}
	purego.SyscallN(addr_QGraphicsObject_setTransformations, this.h, transformations_p, uintptr(len(transformations)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(transformations)
	runtime.KeepAlive(transformations_arr)
}

func (this *QGraphicsObject) TransformOriginPoint() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsObject_transformOriginPoint, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QGraphicsObject) SetTransformOriginPoint(origin *QPointF) {
	purego.SyscallN(addr_QGraphicsObject_setTransformOriginPoint, this.h, uintptr(origin.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(origin)
}

func (this *QGraphicsObject) SetTransformOriginPoint2(ax float64, ay float64) {
	pg_QGraphicsObject_setTransformOriginPoint2(this.h, ax, ay)
	runtime.KeepAlive(this)
}

func (this *QGraphicsObject) Advance(phase int) {
	purego.SyscallN(addr_QGraphicsObject_advance, this.h, uintptr(phase))
	runtime.KeepAlive(this)
}

func (this *QGraphicsObject) ZValue() float64 {
	__r := pg_QGraphicsObject_zValue(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QGraphicsObject) SetZValue(z float64) {
	pg_QGraphicsObject_setZValue(this.h, z)
	runtime.KeepAlive(this)
}

func (this *QGraphicsObject) StackBefore(sibling *QGraphicsItem) {
	purego.SyscallN(addr_QGraphicsObject_stackBefore, this.h, uintptr(sibling.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sibling)
}

func (this *QGraphicsObject) BoundingRect() *QRectF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsObject_boundingRect, this.h)
	runtime.KeepAlive(this)
	return newownedQRectF(__r)
}

func (this *QGraphicsObject) ChildrenBoundingRect() *QRectF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsObject_childrenBoundingRect, this.h)
	runtime.KeepAlive(this)
	return newownedQRectF(__r)
}

func (this *QGraphicsObject) SceneBoundingRect() *QRectF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsObject_sceneBoundingRect, this.h)
	runtime.KeepAlive(this)
	return newownedQRectF(__r)
}

func (this *QGraphicsObject) Shape() *QPainterPath {
	__r, _, _ := purego.SyscallN(addr_QGraphicsObject_shape, this.h)
	runtime.KeepAlive(this)
	return newownedQPainterPath(__r)
}

func (this *QGraphicsObject) IsClipped() bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsObject_isClipped, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGraphicsObject) ClipPath() *QPainterPath {
	__r, _, _ := purego.SyscallN(addr_QGraphicsObject_clipPath, this.h)
	runtime.KeepAlive(this)
	return newownedQPainterPath(__r)
}

func (this *QGraphicsObject) Contains(point *QPointF) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsObject_contains, this.h, uintptr(point.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(point)
	return __r != 0
}

func (this *QGraphicsObject) CollidesWithItem(other *QGraphicsItem, mode ItemSelectionMode) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsObject_collidesWithItem, this.h, uintptr(other.UnsafePointer()), uintptr(mode))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return __r != 0
}

func (this *QGraphicsObject) CollidesWithPath(path *QPainterPath, mode ItemSelectionMode) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsObject_collidesWithPath, this.h, uintptr(path.UnsafePointer()), uintptr(mode))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(path)
	return __r != 0
}

func (this *QGraphicsObject) CollidingItems() []*QGraphicsItem {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGraphicsObject_collidingItems, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QGraphicsItem, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = bindQGraphicsItem(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QGraphicsObject) IsObscured() bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsObject_isObscured, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGraphicsObject) IsObscured2(x float64, y float64, w float64, h float64) bool {
	__r := pg_QGraphicsObject_isObscured2(this.h, x, y, w, h)
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGraphicsObject) IsObscuredBy(item *QGraphicsItem) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsObject_isObscuredBy, this.h, uintptr(item.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
	return __r != 0
}

func (this *QGraphicsObject) OpaqueArea() *QPainterPath {
	__r, _, _ := purego.SyscallN(addr_QGraphicsObject_opaqueArea, this.h)
	runtime.KeepAlive(this)
	return newownedQPainterPath(__r)
}

func (this *QGraphicsObject) BoundingRegion(itemToDeviceTransform *QTransform) *QRegion {
	__r, _, _ := purego.SyscallN(addr_QGraphicsObject_boundingRegion, this.h, uintptr(itemToDeviceTransform.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(itemToDeviceTransform)
	return newownedQRegion(__r)
}

func (this *QGraphicsObject) BoundingRegionGranularity() float64 {
	__r := pg_QGraphicsObject_boundingRegionGranularity(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QGraphicsObject) SetBoundingRegionGranularity(granularity float64) {
	pg_QGraphicsObject_setBoundingRegionGranularity(this.h, granularity)
	runtime.KeepAlive(this)
}

func (this *QGraphicsObject) Paint(painter *QPainter, option *QStyleOptionGraphicsItem, widget *QWidget) {
	purego.SyscallN(addr_QGraphicsObject_paint, this.h, uintptr(painter.UnsafePointer()), uintptr(option.UnsafePointer()), uintptr(widget.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
	runtime.KeepAlive(option)
	runtime.KeepAlive(widget)
}

func (this *QGraphicsObject) Update() {
	purego.SyscallN(addr_QGraphicsObject_update, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsObject) Update2(x float64, y float64, width float64, height float64) {
	pg_QGraphicsObject_update2(this.h, x, y, width, height)
	runtime.KeepAlive(this)
}

func (this *QGraphicsObject) Scroll(dx float64, dy float64) {
	pg_QGraphicsObject_scroll(this.h, dx, dy)
	runtime.KeepAlive(this)
}

func (this *QGraphicsObject) MapToItem(item *QGraphicsItem, point *QPointF) *QPointF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsObject_mapToItem, this.h, uintptr(item.UnsafePointer()), uintptr(point.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
	runtime.KeepAlive(point)
	return newownedQPointF(__r)
}

func (this *QGraphicsObject) MapToParent(point *QPointF) *QPointF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsObject_mapToParent, this.h, uintptr(point.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(point)
	return newownedQPointF(__r)
}

func (this *QGraphicsObject) MapToScene(point *QPointF) *QPointF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsObject_mapToScene, this.h, uintptr(point.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(point)
	return newownedQPointF(__r)
}

func (this *QGraphicsObject) MapToItem2(item *QGraphicsItem, rect *QRectF) *QPolygonF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsObject_mapToItem2, this.h, uintptr(item.UnsafePointer()), uintptr(rect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
	runtime.KeepAlive(rect)
	return newownedQPolygonF(__r)
}

func (this *QGraphicsObject) MapToParentWithRect(rect *QRectF) *QPolygonF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsObject_mapToParent2, this.h, uintptr(rect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
	return newownedQPolygonF(__r)
}

func (this *QGraphicsObject) MapToSceneWithRect(rect *QRectF) *QPolygonF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsObject_mapToScene2, this.h, uintptr(rect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
	return newownedQPolygonF(__r)
}

func (this *QGraphicsObject) MapRectToItem(item *QGraphicsItem, rect *QRectF) *QRectF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsObject_mapRectToItem, this.h, uintptr(item.UnsafePointer()), uintptr(rect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
	runtime.KeepAlive(rect)
	return newownedQRectF(__r)
}

func (this *QGraphicsObject) MapRectToParent(rect *QRectF) *QRectF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsObject_mapRectToParent, this.h, uintptr(rect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
	return newownedQRectF(__r)
}

func (this *QGraphicsObject) MapRectToScene(rect *QRectF) *QRectF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsObject_mapRectToScene, this.h, uintptr(rect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
	return newownedQRectF(__r)
}

func (this *QGraphicsObject) MapToItem3(item *QGraphicsItem, polygon *QPolygonF) *QPolygonF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsObject_mapToItem3, this.h, uintptr(item.UnsafePointer()), uintptr(polygon.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
	runtime.KeepAlive(polygon)
	return newownedQPolygonF(__r)
}

func (this *QGraphicsObject) MapToParentWithPolygon(polygon *QPolygonF) *QPolygonF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsObject_mapToParent3, this.h, uintptr(polygon.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(polygon)
	return newownedQPolygonF(__r)
}

func (this *QGraphicsObject) MapToSceneWithPolygon(polygon *QPolygonF) *QPolygonF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsObject_mapToScene3, this.h, uintptr(polygon.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(polygon)
	return newownedQPolygonF(__r)
}

func (this *QGraphicsObject) MapToItem4(item *QGraphicsItem, path *QPainterPath) *QPainterPath {
	__r, _, _ := purego.SyscallN(addr_QGraphicsObject_mapToItem4, this.h, uintptr(item.UnsafePointer()), uintptr(path.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
	runtime.KeepAlive(path)
	return newownedQPainterPath(__r)
}

func (this *QGraphicsObject) MapToParentWithPath(path *QPainterPath) *QPainterPath {
	__r, _, _ := purego.SyscallN(addr_QGraphicsObject_mapToParent4, this.h, uintptr(path.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(path)
	return newownedQPainterPath(__r)
}

func (this *QGraphicsObject) MapToSceneWithPath(path *QPainterPath) *QPainterPath {
	__r, _, _ := purego.SyscallN(addr_QGraphicsObject_mapToScene4, this.h, uintptr(path.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(path)
	return newownedQPainterPath(__r)
}

func (this *QGraphicsObject) MapFromItem(item *QGraphicsItem, point *QPointF) *QPointF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsObject_mapFromItem, this.h, uintptr(item.UnsafePointer()), uintptr(point.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
	runtime.KeepAlive(point)
	return newownedQPointF(__r)
}

func (this *QGraphicsObject) MapFromParent(point *QPointF) *QPointF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsObject_mapFromParent, this.h, uintptr(point.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(point)
	return newownedQPointF(__r)
}

func (this *QGraphicsObject) MapFromScene(point *QPointF) *QPointF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsObject_mapFromScene, this.h, uintptr(point.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(point)
	return newownedQPointF(__r)
}

func (this *QGraphicsObject) MapFromItem2(item *QGraphicsItem, rect *QRectF) *QPolygonF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsObject_mapFromItem2, this.h, uintptr(item.UnsafePointer()), uintptr(rect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
	runtime.KeepAlive(rect)
	return newownedQPolygonF(__r)
}

func (this *QGraphicsObject) MapFromParentWithRect(rect *QRectF) *QPolygonF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsObject_mapFromParent2, this.h, uintptr(rect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
	return newownedQPolygonF(__r)
}

func (this *QGraphicsObject) MapFromSceneWithRect(rect *QRectF) *QPolygonF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsObject_mapFromScene2, this.h, uintptr(rect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
	return newownedQPolygonF(__r)
}

func (this *QGraphicsObject) MapRectFromItem(item *QGraphicsItem, rect *QRectF) *QRectF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsObject_mapRectFromItem, this.h, uintptr(item.UnsafePointer()), uintptr(rect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
	runtime.KeepAlive(rect)
	return newownedQRectF(__r)
}

func (this *QGraphicsObject) MapRectFromParent(rect *QRectF) *QRectF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsObject_mapRectFromParent, this.h, uintptr(rect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
	return newownedQRectF(__r)
}

func (this *QGraphicsObject) MapRectFromScene(rect *QRectF) *QRectF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsObject_mapRectFromScene, this.h, uintptr(rect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
	return newownedQRectF(__r)
}

func (this *QGraphicsObject) MapFromItem3(item *QGraphicsItem, polygon *QPolygonF) *QPolygonF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsObject_mapFromItem3, this.h, uintptr(item.UnsafePointer()), uintptr(polygon.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
	runtime.KeepAlive(polygon)
	return newownedQPolygonF(__r)
}

func (this *QGraphicsObject) MapFromParentWithPolygon(polygon *QPolygonF) *QPolygonF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsObject_mapFromParent3, this.h, uintptr(polygon.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(polygon)
	return newownedQPolygonF(__r)
}

func (this *QGraphicsObject) MapFromSceneWithPolygon(polygon *QPolygonF) *QPolygonF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsObject_mapFromScene3, this.h, uintptr(polygon.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(polygon)
	return newownedQPolygonF(__r)
}

func (this *QGraphicsObject) MapFromItem4(item *QGraphicsItem, path *QPainterPath) *QPainterPath {
	__r, _, _ := purego.SyscallN(addr_QGraphicsObject_mapFromItem4, this.h, uintptr(item.UnsafePointer()), uintptr(path.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
	runtime.KeepAlive(path)
	return newownedQPainterPath(__r)
}

func (this *QGraphicsObject) MapFromParentWithPath(path *QPainterPath) *QPainterPath {
	__r, _, _ := purego.SyscallN(addr_QGraphicsObject_mapFromParent4, this.h, uintptr(path.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(path)
	return newownedQPainterPath(__r)
}

func (this *QGraphicsObject) MapFromSceneWithPath(path *QPainterPath) *QPainterPath {
	__r, _, _ := purego.SyscallN(addr_QGraphicsObject_mapFromScene4, this.h, uintptr(path.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(path)
	return newownedQPainterPath(__r)
}

func (this *QGraphicsObject) MapToItem5(item *QGraphicsItem, x float64, y float64) *QPointF {
	__r := pg_QGraphicsObject_mapToItem5(this.h, uintptr(item.UnsafePointer()), x, y)
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
	return newownedQPointF(__r)
}

func (this *QGraphicsObject) MapToParent2(x float64, y float64) *QPointF {
	__r := pg_QGraphicsObject_mapToParent5(this.h, x, y)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QGraphicsObject) MapToScene2(x float64, y float64) *QPointF {
	__r := pg_QGraphicsObject_mapToScene5(this.h, x, y)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QGraphicsObject) MapToItem6(item *QGraphicsItem, x float64, y float64, w float64, h float64) *QPolygonF {
	__r := pg_QGraphicsObject_mapToItem6(this.h, uintptr(item.UnsafePointer()), x, y, w, h)
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
	return newownedQPolygonF(__r)
}

func (this *QGraphicsObject) MapToParent3(x float64, y float64, w float64, h float64) *QPolygonF {
	__r := pg_QGraphicsObject_mapToParent6(this.h, x, y, w, h)
	runtime.KeepAlive(this)
	return newownedQPolygonF(__r)
}

func (this *QGraphicsObject) MapToScene3(x float64, y float64, w float64, h float64) *QPolygonF {
	__r := pg_QGraphicsObject_mapToScene6(this.h, x, y, w, h)
	runtime.KeepAlive(this)
	return newownedQPolygonF(__r)
}

func (this *QGraphicsObject) MapRectToItem2(item *QGraphicsItem, x float64, y float64, w float64, h float64) *QRectF {
	__r := pg_QGraphicsObject_mapRectToItem2(this.h, uintptr(item.UnsafePointer()), x, y, w, h)
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
	return newownedQRectF(__r)
}

func (this *QGraphicsObject) MapRectToParent2(x float64, y float64, w float64, h float64) *QRectF {
	__r := pg_QGraphicsObject_mapRectToParent2(this.h, x, y, w, h)
	runtime.KeepAlive(this)
	return newownedQRectF(__r)
}

func (this *QGraphicsObject) MapRectToScene2(x float64, y float64, w float64, h float64) *QRectF {
	__r := pg_QGraphicsObject_mapRectToScene2(this.h, x, y, w, h)
	runtime.KeepAlive(this)
	return newownedQRectF(__r)
}

func (this *QGraphicsObject) MapFromItem5(item *QGraphicsItem, x float64, y float64) *QPointF {
	__r := pg_QGraphicsObject_mapFromItem5(this.h, uintptr(item.UnsafePointer()), x, y)
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
	return newownedQPointF(__r)
}

func (this *QGraphicsObject) MapFromParent2(x float64, y float64) *QPointF {
	__r := pg_QGraphicsObject_mapFromParent5(this.h, x, y)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QGraphicsObject) MapFromScene2(x float64, y float64) *QPointF {
	__r := pg_QGraphicsObject_mapFromScene5(this.h, x, y)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QGraphicsObject) MapFromItem6(item *QGraphicsItem, x float64, y float64, w float64, h float64) *QPolygonF {
	__r := pg_QGraphicsObject_mapFromItem6(this.h, uintptr(item.UnsafePointer()), x, y, w, h)
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
	return newownedQPolygonF(__r)
}

func (this *QGraphicsObject) MapFromParent3(x float64, y float64, w float64, h float64) *QPolygonF {
	__r := pg_QGraphicsObject_mapFromParent6(this.h, x, y, w, h)
	runtime.KeepAlive(this)
	return newownedQPolygonF(__r)
}

func (this *QGraphicsObject) MapFromScene3(x float64, y float64, w float64, h float64) *QPolygonF {
	__r := pg_QGraphicsObject_mapFromScene6(this.h, x, y, w, h)
	runtime.KeepAlive(this)
	return newownedQPolygonF(__r)
}

func (this *QGraphicsObject) MapRectFromItem2(item *QGraphicsItem, x float64, y float64, w float64, h float64) *QRectF {
	__r := pg_QGraphicsObject_mapRectFromItem2(this.h, uintptr(item.UnsafePointer()), x, y, w, h)
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
	return newownedQRectF(__r)
}

func (this *QGraphicsObject) MapRectFromParent2(x float64, y float64, w float64, h float64) *QRectF {
	__r := pg_QGraphicsObject_mapRectFromParent2(this.h, x, y, w, h)
	runtime.KeepAlive(this)
	return newownedQRectF(__r)
}

func (this *QGraphicsObject) MapRectFromScene2(x float64, y float64, w float64, h float64) *QRectF {
	__r := pg_QGraphicsObject_mapRectFromScene2(this.h, x, y, w, h)
	runtime.KeepAlive(this)
	return newownedQRectF(__r)
}

func (this *QGraphicsObject) IsAncestorOf(child *QGraphicsItem) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsObject_isAncestorOf, this.h, uintptr(child.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(child)
	return __r != 0
}

func (this *QGraphicsObject) CommonAncestorItem(other *QGraphicsItem) *QGraphicsItem {
	__r, _, _ := purego.SyscallN(addr_QGraphicsObject_commonAncestorItem, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return bindQGraphicsItem(__r)
}

func (this *QGraphicsObject) IsUnderMouse() bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsObject_isUnderMouse, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGraphicsObject) Data(key int) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QGraphicsObject_data, this.h, uintptr(key))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QGraphicsObject) SetData(key int, value *QVariant) {
	purego.SyscallN(addr_QGraphicsObject_setData, this.h, uintptr(key), uintptr(value.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(value)
}

func (this *QGraphicsObject) InputMethodHints() InputMethodHint {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsObject_inputMethodHints, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return InputMethodHint(__r)
}

func (this *QGraphicsObject) SetInputMethodHints(hints InputMethodHint) {
	purego.SyscallN(addr_QGraphicsObject_setInputMethodHints, this.h, uintptr(hints))
	runtime.KeepAlive(this)
}

func (this *QGraphicsObject) Type() int {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsObject_type, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QGraphicsObject) InstallSceneEventFilter(filterItem *QGraphicsItem) {
	purego.SyscallN(addr_QGraphicsObject_installSceneEventFilter, this.h, uintptr(filterItem.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(filterItem)
}

func (this *QGraphicsObject) RemoveSceneEventFilter(filterItem *QGraphicsItem) {
	purego.SyscallN(addr_QGraphicsObject_removeSceneEventFilter, this.h, uintptr(filterItem.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(filterItem)
}

func (this *QGraphicsObject) AddToIndex() {
	purego.SyscallN(addr_QGraphicsObject_addToIndex, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsObject) RemoveFromIndex() {
	purego.SyscallN(addr_QGraphicsObject_removeFromIndex, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsObject) PrepareGeometryChange() {
	purego.SyscallN(addr_QGraphicsObject_prepareGeometryChange, this.h)
	runtime.KeepAlive(this)
}

func QGraphicsObject_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGraphicsObject_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QGraphicsObject_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGraphicsObject_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func (this *QGraphicsObject) GrabGesture2(type_ GestureType, flags GestureFlag) {
	purego.SyscallN(addr_QGraphicsObject_grabGesture2, this.h, uintptr(type_), uintptr(flags))
	runtime.KeepAlive(this)
}

func QGraphicsObject_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QGraphicsObject_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QGraphicsObject_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QGraphicsObject_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QGraphicsObject) SetFlag2(flag QGraphicsItem__GraphicsItemFlag, enabled bool) {
	purego.SyscallN(addr_QGraphicsObject_setFlag2, this.h, uintptr(flag), uintptr(b2i(enabled)))
	runtime.KeepAlive(this)
}

func (this *QGraphicsObject) SetCacheMode2(mode QGraphicsItem__CacheMode, cacheSize *QSize) {
	purego.SyscallN(addr_QGraphicsObject_setCacheMode2, this.h, uintptr(mode), uintptr(cacheSize.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(cacheSize)
}

func (this *QGraphicsObject) SetFocusWithFocusReason(focusReason FocusReason) {
	purego.SyscallN(addr_QGraphicsObject_setFocus2, this.h, uintptr(focusReason))
	runtime.KeepAlive(this)
}

func (this *QGraphicsObject) EnsureVisibleWithRect(rect *QRectF) {
	purego.SyscallN(addr_QGraphicsObject_ensureVisible3, this.h, uintptr(rect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
}

func (this *QGraphicsObject) EnsureVisible3(rect *QRectF, xmargin int) {
	purego.SyscallN(addr_QGraphicsObject_ensureVisible4, this.h, uintptr(rect.UnsafePointer()), uintptr(xmargin))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
}

func (this *QGraphicsObject) EnsureVisible4(rect *QRectF, xmargin int, ymargin int) {
	purego.SyscallN(addr_QGraphicsObject_ensureVisible5, this.h, uintptr(rect.UnsafePointer()), uintptr(xmargin), uintptr(ymargin))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
}

func (this *QGraphicsObject) EnsureVisible5(x float64, y float64, w float64, h float64, xmargin int) {
	pg_QGraphicsObject_ensureVisible6(this.h, x, y, w, h, int32(xmargin))
	runtime.KeepAlive(this)
}

func (this *QGraphicsObject) EnsureVisible6(x float64, y float64, w float64, h float64, xmargin int, ymargin int) {
	pg_QGraphicsObject_ensureVisible7(this.h, x, y, w, h, int32(xmargin), int32(ymargin))
	runtime.KeepAlive(this)
}

func (this *QGraphicsObject) ItemTransform2(other *QGraphicsItem) (*QTransform, bool) {
	var ok_out int32
	__r, _, _ := purego.SyscallN(addr_QGraphicsObject_itemTransform2, this.h, uintptr(other.UnsafePointer()), uintptr(unsafe.Pointer(&ok_out)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return newownedQTransform(__r), ok_out != 0
}

func (this *QGraphicsObject) SetTransform2(matrix *QTransform, combine bool) {
	purego.SyscallN(addr_QGraphicsObject_setTransform2, this.h, uintptr(matrix.UnsafePointer()), uintptr(b2i(combine)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(matrix)
}

func (this *QGraphicsObject) CollidingItemsWithMode(mode ItemSelectionMode) []*QGraphicsItem {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGraphicsObject_collidingItems2, this.h, uintptr(mode), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QGraphicsItem, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = bindQGraphicsItem(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QGraphicsObject) IsObscuredWithRect(rect *QRectF) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsObject_isObscured3, this.h, uintptr(rect.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
	return __r != 0
}

func (this *QGraphicsObject) UpdateWithRect(rect *QRectF) {
	purego.SyscallN(addr_QGraphicsObject_update3, this.h, uintptr(rect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
}

func (this *QGraphicsObject) Scroll2(dx float64, dy float64, rect *QRectF) {
	pg_QGraphicsObject_scroll2(this.h, dx, dy, uintptr(rect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
}

func (this *QGraphicsObject) AsQGraphicsItem() *QGraphicsItem {
	__r, _, _ := purego.SyscallN(addr_QGraphicsObject_asQGraphicsItem, this.h)
	runtime.KeepAlive(this)
	return bindQGraphicsItem(__r)
}

func (this *QGraphicsObject) AdvanceDefault(phase int) {
	purego.SyscallN(addr_QGraphicsObject_advance_default, this.h, uintptr(phase))
	runtime.KeepAlive(this)
}

func (this *QGraphicsObject) ShapeDefault() *QPainterPath {
	__r, _, _ := purego.SyscallN(addr_QGraphicsObject_shape_default, this.h)
	runtime.KeepAlive(this)
	return newownedQPainterPath(__r)
}

func (this *QGraphicsObject) ContainsDefault(point *QPointF) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsObject_contains_default, this.h, uintptr(point.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(point)
	return __r != 0
}

func (this *QGraphicsObject) CollidesWithItemDefault(other *QGraphicsItem, mode ItemSelectionMode) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsObject_collidesWithItem_default, this.h, uintptr(other.UnsafePointer()), uintptr(mode))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return __r != 0
}

func (this *QGraphicsObject) CollidesWithPathDefault(path *QPainterPath, mode ItemSelectionMode) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsObject_collidesWithPath_default, this.h, uintptr(path.UnsafePointer()), uintptr(mode))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(path)
	return __r != 0
}

func (this *QGraphicsObject) IsObscuredByDefault(item *QGraphicsItem) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsObject_isObscuredBy_default, this.h, uintptr(item.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
	return __r != 0
}

func (this *QGraphicsObject) OpaqueAreaDefault() *QPainterPath {
	__r, _, _ := purego.SyscallN(addr_QGraphicsObject_opaqueArea_default, this.h)
	runtime.KeepAlive(this)
	return newownedQPainterPath(__r)
}

func (this *QGraphicsObject) TypeDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsObject_type_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QGraphicsObject) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QGraphicsObject_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QGraphicsObject) EventDefault(ev *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsObject_event_default, this.h, uintptr(ev.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(ev)
	return __r != 0
}

func (this *QGraphicsObject) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsObject_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QGraphicsObject) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QGraphicsObject_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsObject) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QGraphicsObject_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsObject) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QGraphicsObject_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsObject) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QGraphicsObject_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QGraphicsObject) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QGraphicsObject_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QGraphicsObject) SceneEventFilterDefault(watched *QGraphicsItem, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsObject_sceneEventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QGraphicsObject) SceneEventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsObject_sceneEvent_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QGraphicsObject) ContextMenuEventDefault(event *QGraphicsSceneContextMenuEvent) {
	purego.SyscallN(addr_QGraphicsObject_contextMenuEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsObject) DragEnterEventDefault(event *QGraphicsSceneDragDropEvent) {
	purego.SyscallN(addr_QGraphicsObject_dragEnterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsObject) DragLeaveEventDefault(event *QGraphicsSceneDragDropEvent) {
	purego.SyscallN(addr_QGraphicsObject_dragLeaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsObject) DragMoveEventDefault(event *QGraphicsSceneDragDropEvent) {
	purego.SyscallN(addr_QGraphicsObject_dragMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsObject) DropEventDefault(event *QGraphicsSceneDragDropEvent) {
	purego.SyscallN(addr_QGraphicsObject_dropEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsObject) FocusInEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QGraphicsObject_focusInEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsObject) FocusOutEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QGraphicsObject_focusOutEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsObject) HoverEnterEventDefault(event *QGraphicsSceneHoverEvent) {
	purego.SyscallN(addr_QGraphicsObject_hoverEnterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsObject) HoverMoveEventDefault(event *QGraphicsSceneHoverEvent) {
	purego.SyscallN(addr_QGraphicsObject_hoverMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsObject) HoverLeaveEventDefault(event *QGraphicsSceneHoverEvent) {
	purego.SyscallN(addr_QGraphicsObject_hoverLeaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsObject) KeyPressEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QGraphicsObject_keyPressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsObject) KeyReleaseEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QGraphicsObject_keyReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsObject) MousePressEventDefault(event *QGraphicsSceneMouseEvent) {
	purego.SyscallN(addr_QGraphicsObject_mousePressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsObject) MouseMoveEventDefault(event *QGraphicsSceneMouseEvent) {
	purego.SyscallN(addr_QGraphicsObject_mouseMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsObject) MouseReleaseEventDefault(event *QGraphicsSceneMouseEvent) {
	purego.SyscallN(addr_QGraphicsObject_mouseReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsObject) MouseDoubleClickEventDefault(event *QGraphicsSceneMouseEvent) {
	purego.SyscallN(addr_QGraphicsObject_mouseDoubleClickEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsObject) WheelEventDefault(event *QGraphicsSceneWheelEvent) {
	purego.SyscallN(addr_QGraphicsObject_wheelEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsObject) InputMethodEventDefault(event *QInputMethodEvent) {
	purego.SyscallN(addr_QGraphicsObject_inputMethodEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsObject) InputMethodQueryDefault(query InputMethodQuery) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QGraphicsObject_inputMethodQuery_default, this.h, uintptr(query))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QGraphicsObject) ItemChangeDefault(change QGraphicsItem__GraphicsItemChange, value *QVariant) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QGraphicsObject_itemChange_default, this.h, uintptr(change), uintptr(value.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(value)
	return newownedQVariant(__r)
}

func (this *QGraphicsObject) SupportsExtensionDefault(extension QGraphicsItem__Extension) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsObject_supportsExtension_default, this.h, uintptr(extension))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGraphicsObject) SetExtensionDefault(extension QGraphicsItem__Extension, variant *QVariant) {
	purego.SyscallN(addr_QGraphicsObject_setExtension_default, this.h, uintptr(extension), uintptr(variant.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(variant)
}

func (this *QGraphicsObject) ExtensionDefault(variant *QVariant) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QGraphicsObject_extension_default, this.h, uintptr(variant.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(variant)
	return newownedQVariant(__r)
}

func (this *QGraphicsObject) OnParentChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QGraphicsObject_connect_parentChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsObject) OnOpacityChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QGraphicsObject_connect_opacityChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsObject) OnVisibleChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QGraphicsObject_connect_visibleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsObject) OnEnabledChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QGraphicsObject_connect_enabledChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsObject) OnXChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QGraphicsObject_connect_xChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsObject) OnYChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QGraphicsObject_connect_yChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsObject) OnZChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QGraphicsObject_connect_zChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsObject) OnRotationChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QGraphicsObject_connect_rotationChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsObject) OnScaleChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QGraphicsObject_connect_scaleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsObject) OnChildrenChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QGraphicsObject_connect_childrenChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsObject) OnWidthChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QGraphicsObject_connect_widthChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsObject) OnHeightChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QGraphicsObject_connect_heightChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsObject) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QGraphicsObject_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsObject) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QGraphicsObject_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsObject) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QGraphicsObject_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsObject) Disconnect(handle int64) {
	pg_QGraphicsObject__disconnect(handle)
	unregisterCallback(handle)
}

func QGraphicsObject_virtIdx(slot int64) int32 {
	switch slot {
	case 2510008782785038831:
		return 0
	case 6083584261053556819:
		return 1
	case 1607084412591819503:
		return 2
	case 1237236456312824991:
		return 3
	case 5217623514630858297:
		return 4
	case -5409927581812374498:
		return 5
	case -1185315367526857833:
		return 6
	case 3769868229123868474:
		return 7
	case 7379029982170503008:
		return 8
	case -568369772982215962:
		return 9
	case -4830836341480073475:
		return 10
	case 7938019069247444294:
		return 11
	case 5645366511944179768:
		return 12
	case 6469736184362593173:
		return 13
	case 6636575182718753853:
		return 14
	case -7215030271914903671:
		return 15
	case 399370970023989435:
		return 16
	case -6371403872082420691:
		return 17
	case -4584203218241215725:
		return 18
	case -2846371507838587071:
		return 19
	case -6698043222227006015:
		return 20
	case 206122715919917589:
		return 21
	case -3139904158373725006:
		return 22
	case -1781545137352745978:
		return 23
	case 9107058682754591716:
		return 24
	case 975043767033552552:
		return 25
	case -385596010576048993:
		return 26
	case 6932941943906681679:
		return 27
	case -2745191522144316396:
		return 28
	case 7944408482203603984:
		return 29
	case 1023791977679793185:
		return 30
	case 772598475950050299:
		return 31
	case -332090898078543319:
		return 32
	case -4255494004099806581:
		return 33
	case -4415920049297795213:
		return 34
	case -4029073905697023887:
		return 35
	case 1919988340934511316:
		return 36
	case 5438819661942843222:
		return 37
	case 7345853097340709638:
		return 38
	case -1888572168794878830:
		return 39
	case -5729166443651306418:
		return 40
	case -6730181576106171580:
		return 41
	case 2270616358835251226:
		return 42
	}
	return -1
}

func QGraphicsObject_enable(h uintptr, idx int32) { pg_QGraphicsObject__set_override(h, idx) }

func NewQGraphicsObject() *QGraphicsObject {
	id := nextVirtID()
	o := newownedQGraphicsObject(pg_QGraphicsObject__Go_new(id))
	o.UnsafeSetGoID(id)
	registerInstance(id, o.h, QGraphicsObject_virtIdx, QGraphicsObject_enable)
	return o
}

func (this *QGraphicsObject) OnBoundingRect(fn func() *QRectF) {
	registerVirt(this.UnsafeGoID(), 2510008782785038831, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "boundingRect"
}

func (this *QGraphicsObject) OnPaint(fn func(*QPainter, *QStyleOptionGraphicsItem, *QWidget)) {
	registerVirt(this.UnsafeGoID(), 6083584261053556819, func(painter uintptr, option uintptr, widget uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter)), UnsafeNewQStyleOptionGraphicsItem(unsafe.Pointer(option)), UnsafeNewQWidget(unsafe.Pointer(widget))) }) // slot "paint"
}

func (this *QGraphicsObject) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QGraphicsObject) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(ev uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(ev)))) }) // slot "event"
}

func (this *QGraphicsObject) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QGraphicsObject) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QGraphicsObject) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QGraphicsObject) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QGraphicsObject) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QGraphicsObject) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QGraphicsObject) OnAdvance(fn func(int)) {
	registerVirt(this.UnsafeGoID(), -4830836341480073475, func(phase int32) { fn(int(phase)) }) // slot "advance"
}

func (this *QGraphicsObject) OnShape(fn func() *QPainterPath) {
	registerVirt(this.UnsafeGoID(), 7938019069247444294, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "shape"
}

func (this *QGraphicsObject) OnContains(fn func(*QPointF) bool) {
	registerVirt(this.UnsafeGoID(), 5645366511944179768, func(point uintptr) int32 { return b2i(fn(UnsafeNewQPointF(unsafe.Pointer(point)))) }) // slot "contains"
}

func (this *QGraphicsObject) OnCollidesWithItem(fn func(*QGraphicsItem, ItemSelectionMode) bool) {
	registerVirt(this.UnsafeGoID(), 6469736184362593173, func(other uintptr, mode int32) int32 { return b2i(fn(UnsafeNewQGraphicsItem(unsafe.Pointer(other)), ItemSelectionMode(mode))) }) // slot "collidesWithItem"
}

func (this *QGraphicsObject) OnCollidesWithPath(fn func(*QPainterPath, ItemSelectionMode) bool) {
	registerVirt(this.UnsafeGoID(), 6636575182718753853, func(path uintptr, mode int32) int32 { return b2i(fn(UnsafeNewQPainterPath(unsafe.Pointer(path)), ItemSelectionMode(mode))) }) // slot "collidesWithPath"
}

func (this *QGraphicsObject) OnIsObscuredBy(fn func(*QGraphicsItem) bool) {
	registerVirt(this.UnsafeGoID(), -7215030271914903671, func(item uintptr) int32 { return b2i(fn(UnsafeNewQGraphicsItem(unsafe.Pointer(item)))) }) // slot "isObscuredBy"
}

func (this *QGraphicsObject) OnOpaqueArea(fn func() *QPainterPath) {
	registerVirt(this.UnsafeGoID(), 399370970023989435, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "opaqueArea"
}

func (this *QGraphicsObject) OnType(fn func() int) {
	registerVirt(this.UnsafeGoID(), -6371403872082420691, func() int32 { return int32(fn()) }) // slot "type"
}

func (this *QGraphicsObject) OnSceneEventFilter(fn func(*QGraphicsItem, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), -4584203218241215725, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQGraphicsItem(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "sceneEventFilter"
}

func (this *QGraphicsObject) OnSceneEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), -2846371507838587071, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "sceneEvent"
}

func (this *QGraphicsObject) OnContextMenuEvent(fn func(*QGraphicsSceneContextMenuEvent)) {
	registerVirt(this.UnsafeGoID(), -6698043222227006015, func(event uintptr) { fn(UnsafeNewQGraphicsSceneContextMenuEvent(unsafe.Pointer(event))) }) // slot "contextMenuEvent"
}

func (this *QGraphicsObject) OnDragEnterEvent(fn func(*QGraphicsSceneDragDropEvent)) {
	registerVirt(this.UnsafeGoID(), 206122715919917589, func(event uintptr) { fn(UnsafeNewQGraphicsSceneDragDropEvent(unsafe.Pointer(event))) }) // slot "dragEnterEvent"
}

func (this *QGraphicsObject) OnDragLeaveEvent(fn func(*QGraphicsSceneDragDropEvent)) {
	registerVirt(this.UnsafeGoID(), -3139904158373725006, func(event uintptr) { fn(UnsafeNewQGraphicsSceneDragDropEvent(unsafe.Pointer(event))) }) // slot "dragLeaveEvent"
}

func (this *QGraphicsObject) OnDragMoveEvent(fn func(*QGraphicsSceneDragDropEvent)) {
	registerVirt(this.UnsafeGoID(), -1781545137352745978, func(event uintptr) { fn(UnsafeNewQGraphicsSceneDragDropEvent(unsafe.Pointer(event))) }) // slot "dragMoveEvent"
}

func (this *QGraphicsObject) OnDropEvent(fn func(*QGraphicsSceneDragDropEvent)) {
	registerVirt(this.UnsafeGoID(), 9107058682754591716, func(event uintptr) { fn(UnsafeNewQGraphicsSceneDragDropEvent(unsafe.Pointer(event))) }) // slot "dropEvent"
}

func (this *QGraphicsObject) OnFocusInEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), 975043767033552552, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusInEvent"
}

func (this *QGraphicsObject) OnFocusOutEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), -385596010576048993, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusOutEvent"
}

func (this *QGraphicsObject) OnHoverEnterEvent(fn func(*QGraphicsSceneHoverEvent)) {
	registerVirt(this.UnsafeGoID(), 6932941943906681679, func(event uintptr) { fn(UnsafeNewQGraphicsSceneHoverEvent(unsafe.Pointer(event))) }) // slot "hoverEnterEvent"
}

func (this *QGraphicsObject) OnHoverMoveEvent(fn func(*QGraphicsSceneHoverEvent)) {
	registerVirt(this.UnsafeGoID(), -2745191522144316396, func(event uintptr) { fn(UnsafeNewQGraphicsSceneHoverEvent(unsafe.Pointer(event))) }) // slot "hoverMoveEvent"
}

func (this *QGraphicsObject) OnHoverLeaveEvent(fn func(*QGraphicsSceneHoverEvent)) {
	registerVirt(this.UnsafeGoID(), 7944408482203603984, func(event uintptr) { fn(UnsafeNewQGraphicsSceneHoverEvent(unsafe.Pointer(event))) }) // slot "hoverLeaveEvent"
}

func (this *QGraphicsObject) OnKeyPressEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 1023791977679793185, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyPressEvent"
}

func (this *QGraphicsObject) OnKeyReleaseEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 772598475950050299, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyReleaseEvent"
}

func (this *QGraphicsObject) OnMousePressEvent(fn func(*QGraphicsSceneMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -332090898078543319, func(event uintptr) { fn(UnsafeNewQGraphicsSceneMouseEvent(unsafe.Pointer(event))) }) // slot "mousePressEvent"
}

func (this *QGraphicsObject) OnMouseMoveEvent(fn func(*QGraphicsSceneMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4255494004099806581, func(event uintptr) { fn(UnsafeNewQGraphicsSceneMouseEvent(unsafe.Pointer(event))) }) // slot "mouseMoveEvent"
}

func (this *QGraphicsObject) OnMouseReleaseEvent(fn func(*QGraphicsSceneMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4415920049297795213, func(event uintptr) { fn(UnsafeNewQGraphicsSceneMouseEvent(unsafe.Pointer(event))) }) // slot "mouseReleaseEvent"
}

func (this *QGraphicsObject) OnMouseDoubleClickEvent(fn func(*QGraphicsSceneMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4029073905697023887, func(event uintptr) { fn(UnsafeNewQGraphicsSceneMouseEvent(unsafe.Pointer(event))) }) // slot "mouseDoubleClickEvent"
}

func (this *QGraphicsObject) OnWheelEvent(fn func(*QGraphicsSceneWheelEvent)) {
	registerVirt(this.UnsafeGoID(), 1919988340934511316, func(event uintptr) { fn(UnsafeNewQGraphicsSceneWheelEvent(unsafe.Pointer(event))) }) // slot "wheelEvent"
}

func (this *QGraphicsObject) OnInputMethodEvent(fn func(*QInputMethodEvent)) {
	registerVirt(this.UnsafeGoID(), 5438819661942843222, func(event uintptr) { fn(UnsafeNewQInputMethodEvent(unsafe.Pointer(event))) }) // slot "inputMethodEvent"
}

func (this *QGraphicsObject) OnInputMethodQuery(fn func(InputMethodQuery) *QVariant) {
	registerVirt(this.UnsafeGoID(), 7345853097340709638, func(query int32) uintptr { return uintptr(fn(InputMethodQuery(query)).UnsafePointer()) }) // slot "inputMethodQuery"
}

func (this *QGraphicsObject) OnItemChange(fn func(QGraphicsItem__GraphicsItemChange, *QVariant) *QVariant) {
	registerVirt(this.UnsafeGoID(), -1888572168794878830, func(change int32, value uintptr) uintptr { return uintptr(fn(QGraphicsItem__GraphicsItemChange(change), UnsafeNewQVariant(unsafe.Pointer(value))).UnsafePointer()) }) // slot "itemChange"
}

func (this *QGraphicsObject) OnSupportsExtension(fn func(QGraphicsItem__Extension) bool) {
	registerVirt(this.UnsafeGoID(), -5729166443651306418, func(extension int32) int32 { return b2i(fn(QGraphicsItem__Extension(extension))) }) // slot "supportsExtension"
}

func (this *QGraphicsObject) OnSetExtension(fn func(QGraphicsItem__Extension, *QVariant)) {
	registerVirt(this.UnsafeGoID(), -6730181576106171580, func(extension int32, variant uintptr) { fn(QGraphicsItem__Extension(extension), UnsafeNewQVariant(unsafe.Pointer(variant))) }) // slot "setExtension"
}

func (this *QGraphicsObject) OnExtension(fn func(*QVariant) *QVariant) {
	registerVirt(this.UnsafeGoID(), 2270616358835251226, func(variant uintptr) uintptr { return uintptr(fn(UnsafeNewQVariant(unsafe.Pointer(variant))).UnsafePointer()) }) // slot "extension"
}

func (this *QGraphicsObject) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QGraphicsObject_delete(this.h)
	}
	this.UnsafeRelease()
}
