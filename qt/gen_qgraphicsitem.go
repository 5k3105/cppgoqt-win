package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QGraphicsItem_scene uintptr
	addr_QGraphicsItem_parentItem uintptr
	addr_QGraphicsItem_topLevelItem uintptr
	addr_QGraphicsItem_parentObject uintptr
	addr_QGraphicsItem_parentWidget uintptr
	addr_QGraphicsItem_topLevelWidget uintptr
	addr_QGraphicsItem_window uintptr
	addr_QGraphicsItem_panel uintptr
	addr_QGraphicsItem_setParentItem uintptr
	addr_QGraphicsItem_childItems uintptr
	addr_QGraphicsItem_isWidget uintptr
	addr_QGraphicsItem_isWindow uintptr
	addr_QGraphicsItem_isPanel uintptr
	addr_QGraphicsItem_toGraphicsObject uintptr
	addr_QGraphicsItem_toGraphicsObject2 uintptr
	addr_QGraphicsItem_group uintptr
	addr_QGraphicsItem_setGroup uintptr
	addr_QGraphicsItem_flags uintptr
	addr_QGraphicsItem_setFlag uintptr
	addr_QGraphicsItem_setFlags uintptr
	addr_QGraphicsItem_cacheMode uintptr
	addr_QGraphicsItem_setCacheMode uintptr
	addr_QGraphicsItem_panelModality uintptr
	addr_QGraphicsItem_setPanelModality uintptr
	addr_QGraphicsItem_isBlockedByModalPanel uintptr
	addr_QGraphicsItem_toolTip uintptr
	addr_QGraphicsItem_setToolTip uintptr
	addr_QGraphicsItem_cursor uintptr
	addr_QGraphicsItem_setCursor uintptr
	addr_QGraphicsItem_hasCursor uintptr
	addr_QGraphicsItem_unsetCursor uintptr
	addr_QGraphicsItem_isVisible uintptr
	addr_QGraphicsItem_isVisibleTo uintptr
	addr_QGraphicsItem_setVisible uintptr
	addr_QGraphicsItem_hide uintptr
	addr_QGraphicsItem_show uintptr
	addr_QGraphicsItem_isEnabled uintptr
	addr_QGraphicsItem_setEnabled uintptr
	addr_QGraphicsItem_isSelected uintptr
	addr_QGraphicsItem_setSelected uintptr
	addr_QGraphicsItem_acceptDrops uintptr
	addr_QGraphicsItem_setAcceptDrops uintptr
	pg_QGraphicsItem_opacity func(uintptr) float64
	pg_QGraphicsItem_effectiveOpacity func(uintptr) float64
	pg_QGraphicsItem_setOpacity func(uintptr, float64)
	addr_QGraphicsItem_graphicsEffect uintptr
	addr_QGraphicsItem_setGraphicsEffect uintptr
	addr_QGraphicsItem_acceptedMouseButtons uintptr
	addr_QGraphicsItem_setAcceptedMouseButtons uintptr
	addr_QGraphicsItem_acceptHoverEvents uintptr
	addr_QGraphicsItem_setAcceptHoverEvents uintptr
	addr_QGraphicsItem_acceptTouchEvents uintptr
	addr_QGraphicsItem_setAcceptTouchEvents uintptr
	addr_QGraphicsItem_filtersChildEvents uintptr
	addr_QGraphicsItem_setFiltersChildEvents uintptr
	addr_QGraphicsItem_handlesChildEvents uintptr
	addr_QGraphicsItem_setHandlesChildEvents uintptr
	addr_QGraphicsItem_isActive uintptr
	addr_QGraphicsItem_setActive uintptr
	addr_QGraphicsItem_hasFocus uintptr
	addr_QGraphicsItem_setFocus uintptr
	addr_QGraphicsItem_clearFocus uintptr
	addr_QGraphicsItem_focusProxy uintptr
	addr_QGraphicsItem_setFocusProxy uintptr
	addr_QGraphicsItem_focusItem uintptr
	addr_QGraphicsItem_focusScopeItem uintptr
	addr_QGraphicsItem_grabMouse uintptr
	addr_QGraphicsItem_ungrabMouse uintptr
	addr_QGraphicsItem_grabKeyboard uintptr
	addr_QGraphicsItem_ungrabKeyboard uintptr
	addr_QGraphicsItem_pos uintptr
	pg_QGraphicsItem_x func(uintptr) float64
	pg_QGraphicsItem_setX func(uintptr, float64)
	pg_QGraphicsItem_y func(uintptr) float64
	pg_QGraphicsItem_setY func(uintptr, float64)
	addr_QGraphicsItem_scenePos uintptr
	addr_QGraphicsItem_setPos uintptr
	pg_QGraphicsItem_setPos2 func(uintptr, float64, float64)
	pg_QGraphicsItem_moveBy func(uintptr, float64, float64)
	addr_QGraphicsItem_ensureVisible uintptr
	pg_QGraphicsItem_ensureVisible2 func(uintptr, float64, float64, float64, float64)
	addr_QGraphicsItem_transform uintptr
	addr_QGraphicsItem_sceneTransform uintptr
	addr_QGraphicsItem_deviceTransform uintptr
	addr_QGraphicsItem_itemTransform uintptr
	addr_QGraphicsItem_setTransform uintptr
	addr_QGraphicsItem_resetTransform uintptr
	pg_QGraphicsItem_setRotation func(uintptr, float64)
	pg_QGraphicsItem_rotation func(uintptr) float64
	pg_QGraphicsItem_setScale func(uintptr, float64)
	pg_QGraphicsItem_scale func(uintptr) float64
	addr_QGraphicsItem_transformations uintptr
	addr_QGraphicsItem_setTransformations uintptr
	addr_QGraphicsItem_transformOriginPoint uintptr
	addr_QGraphicsItem_setTransformOriginPoint uintptr
	pg_QGraphicsItem_setTransformOriginPoint2 func(uintptr, float64, float64)
	addr_QGraphicsItem_advance uintptr
	pg_QGraphicsItem_zValue func(uintptr) float64
	pg_QGraphicsItem_setZValue func(uintptr, float64)
	addr_QGraphicsItem_stackBefore uintptr
	addr_QGraphicsItem_boundingRect uintptr
	addr_QGraphicsItem_childrenBoundingRect uintptr
	addr_QGraphicsItem_sceneBoundingRect uintptr
	addr_QGraphicsItem_shape uintptr
	addr_QGraphicsItem_isClipped uintptr
	addr_QGraphicsItem_clipPath uintptr
	addr_QGraphicsItem_contains uintptr
	addr_QGraphicsItem_collidesWithItem uintptr
	addr_QGraphicsItem_collidesWithPath uintptr
	addr_QGraphicsItem_collidingItems uintptr
	addr_QGraphicsItem_isObscured uintptr
	pg_QGraphicsItem_isObscured2 func(uintptr, float64, float64, float64, float64) int32
	addr_QGraphicsItem_isObscuredBy uintptr
	addr_QGraphicsItem_opaqueArea uintptr
	addr_QGraphicsItem_boundingRegion uintptr
	pg_QGraphicsItem_boundingRegionGranularity func(uintptr) float64
	pg_QGraphicsItem_setBoundingRegionGranularity func(uintptr, float64)
	addr_QGraphicsItem_paint uintptr
	addr_QGraphicsItem_update uintptr
	pg_QGraphicsItem_update2 func(uintptr, float64, float64, float64, float64)
	pg_QGraphicsItem_scroll func(uintptr, float64, float64)
	addr_QGraphicsItem_mapToItem uintptr
	addr_QGraphicsItem_mapToParent uintptr
	addr_QGraphicsItem_mapToScene uintptr
	addr_QGraphicsItem_mapToItem2 uintptr
	addr_QGraphicsItem_mapToParent2 uintptr
	addr_QGraphicsItem_mapToScene2 uintptr
	addr_QGraphicsItem_mapRectToItem uintptr
	addr_QGraphicsItem_mapRectToParent uintptr
	addr_QGraphicsItem_mapRectToScene uintptr
	addr_QGraphicsItem_mapToItem3 uintptr
	addr_QGraphicsItem_mapToParent3 uintptr
	addr_QGraphicsItem_mapToScene3 uintptr
	addr_QGraphicsItem_mapToItem4 uintptr
	addr_QGraphicsItem_mapToParent4 uintptr
	addr_QGraphicsItem_mapToScene4 uintptr
	addr_QGraphicsItem_mapFromItem uintptr
	addr_QGraphicsItem_mapFromParent uintptr
	addr_QGraphicsItem_mapFromScene uintptr
	addr_QGraphicsItem_mapFromItem2 uintptr
	addr_QGraphicsItem_mapFromParent2 uintptr
	addr_QGraphicsItem_mapFromScene2 uintptr
	addr_QGraphicsItem_mapRectFromItem uintptr
	addr_QGraphicsItem_mapRectFromParent uintptr
	addr_QGraphicsItem_mapRectFromScene uintptr
	addr_QGraphicsItem_mapFromItem3 uintptr
	addr_QGraphicsItem_mapFromParent3 uintptr
	addr_QGraphicsItem_mapFromScene3 uintptr
	addr_QGraphicsItem_mapFromItem4 uintptr
	addr_QGraphicsItem_mapFromParent4 uintptr
	addr_QGraphicsItem_mapFromScene4 uintptr
	pg_QGraphicsItem_mapToItem5 func(uintptr, uintptr, float64, float64) uintptr
	pg_QGraphicsItem_mapToParent5 func(uintptr, float64, float64) uintptr
	pg_QGraphicsItem_mapToScene5 func(uintptr, float64, float64) uintptr
	pg_QGraphicsItem_mapToItem6 func(uintptr, uintptr, float64, float64, float64, float64) uintptr
	pg_QGraphicsItem_mapToParent6 func(uintptr, float64, float64, float64, float64) uintptr
	pg_QGraphicsItem_mapToScene6 func(uintptr, float64, float64, float64, float64) uintptr
	pg_QGraphicsItem_mapRectToItem2 func(uintptr, uintptr, float64, float64, float64, float64) uintptr
	pg_QGraphicsItem_mapRectToParent2 func(uintptr, float64, float64, float64, float64) uintptr
	pg_QGraphicsItem_mapRectToScene2 func(uintptr, float64, float64, float64, float64) uintptr
	pg_QGraphicsItem_mapFromItem5 func(uintptr, uintptr, float64, float64) uintptr
	pg_QGraphicsItem_mapFromParent5 func(uintptr, float64, float64) uintptr
	pg_QGraphicsItem_mapFromScene5 func(uintptr, float64, float64) uintptr
	pg_QGraphicsItem_mapFromItem6 func(uintptr, uintptr, float64, float64, float64, float64) uintptr
	pg_QGraphicsItem_mapFromParent6 func(uintptr, float64, float64, float64, float64) uintptr
	pg_QGraphicsItem_mapFromScene6 func(uintptr, float64, float64, float64, float64) uintptr
	pg_QGraphicsItem_mapRectFromItem2 func(uintptr, uintptr, float64, float64, float64, float64) uintptr
	pg_QGraphicsItem_mapRectFromParent2 func(uintptr, float64, float64, float64, float64) uintptr
	pg_QGraphicsItem_mapRectFromScene2 func(uintptr, float64, float64, float64, float64) uintptr
	addr_QGraphicsItem_isAncestorOf uintptr
	addr_QGraphicsItem_commonAncestorItem uintptr
	addr_QGraphicsItem_isUnderMouse uintptr
	addr_QGraphicsItem_data uintptr
	addr_QGraphicsItem_setData uintptr
	addr_QGraphicsItem_inputMethodHints uintptr
	addr_QGraphicsItem_setInputMethodHints uintptr
	addr_QGraphicsItem_type uintptr
	addr_QGraphicsItem_installSceneEventFilter uintptr
	addr_QGraphicsItem_removeSceneEventFilter uintptr
	addr_QGraphicsItem_updateMicroFocus uintptr
	addr_QGraphicsItem_addToIndex uintptr
	addr_QGraphicsItem_removeFromIndex uintptr
	addr_QGraphicsItem_prepareGeometryChange uintptr
	addr_QGraphicsItem_setFlag2 uintptr
	addr_QGraphicsItem_setCacheMode2 uintptr
	addr_QGraphicsItem_setFocus2 uintptr
	addr_QGraphicsItem_ensureVisible3 uintptr
	addr_QGraphicsItem_ensureVisible4 uintptr
	addr_QGraphicsItem_ensureVisible5 uintptr
	pg_QGraphicsItem_ensureVisible6 func(uintptr, float64, float64, float64, float64, int32)
	pg_QGraphicsItem_ensureVisible7 func(uintptr, float64, float64, float64, float64, int32, int32)
	addr_QGraphicsItem_itemTransform2 uintptr
	addr_QGraphicsItem_setTransform2 uintptr
	addr_QGraphicsItem_collidingItems2 uintptr
	addr_QGraphicsItem_isObscured3 uintptr
	addr_QGraphicsItem_update3 uintptr
	pg_QGraphicsItem_scroll2 func(uintptr, float64, float64, uintptr)
	addr_QGraphicsItem_advance_default uintptr
	addr_QGraphicsItem_shape_default uintptr
	addr_QGraphicsItem_contains_default uintptr
	addr_QGraphicsItem_collidesWithItem_default uintptr
	addr_QGraphicsItem_collidesWithPath_default uintptr
	addr_QGraphicsItem_isObscuredBy_default uintptr
	addr_QGraphicsItem_opaqueArea_default uintptr
	addr_QGraphicsItem_type_default uintptr
	addr_QGraphicsItem_sceneEventFilter_default uintptr
	addr_QGraphicsItem_sceneEvent_default uintptr
	addr_QGraphicsItem_contextMenuEvent_default uintptr
	addr_QGraphicsItem_dragEnterEvent_default uintptr
	addr_QGraphicsItem_dragLeaveEvent_default uintptr
	addr_QGraphicsItem_dragMoveEvent_default uintptr
	addr_QGraphicsItem_dropEvent_default uintptr
	addr_QGraphicsItem_focusInEvent_default uintptr
	addr_QGraphicsItem_focusOutEvent_default uintptr
	addr_QGraphicsItem_hoverEnterEvent_default uintptr
	addr_QGraphicsItem_hoverMoveEvent_default uintptr
	addr_QGraphicsItem_hoverLeaveEvent_default uintptr
	addr_QGraphicsItem_keyPressEvent_default uintptr
	addr_QGraphicsItem_keyReleaseEvent_default uintptr
	addr_QGraphicsItem_mousePressEvent_default uintptr
	addr_QGraphicsItem_mouseMoveEvent_default uintptr
	addr_QGraphicsItem_mouseReleaseEvent_default uintptr
	addr_QGraphicsItem_mouseDoubleClickEvent_default uintptr
	addr_QGraphicsItem_wheelEvent_default uintptr
	addr_QGraphicsItem_inputMethodEvent_default uintptr
	addr_QGraphicsItem_inputMethodQuery_default uintptr
	addr_QGraphicsItem_itemChange_default uintptr
	addr_QGraphicsItem_supportsExtension_default uintptr
	addr_QGraphicsItem_setExtension_default uintptr
	addr_QGraphicsItem_extension_default uintptr
	pg_QGraphicsItem__set_override func(uintptr, int32)
	pg_QGraphicsItem__Go_new func(int64) uintptr
	pg_QGraphicsItem_delete func(uintptr)
)

func init() {
	symaddr(&addr_QGraphicsItem_scene, "QGraphicsItem_scene")
	symaddr(&addr_QGraphicsItem_parentItem, "QGraphicsItem_parentItem")
	symaddr(&addr_QGraphicsItem_topLevelItem, "QGraphicsItem_topLevelItem")
	symaddr(&addr_QGraphicsItem_parentObject, "QGraphicsItem_parentObject")
	symaddr(&addr_QGraphicsItem_parentWidget, "QGraphicsItem_parentWidget")
	symaddr(&addr_QGraphicsItem_topLevelWidget, "QGraphicsItem_topLevelWidget")
	symaddr(&addr_QGraphicsItem_window, "QGraphicsItem_window")
	symaddr(&addr_QGraphicsItem_panel, "QGraphicsItem_panel")
	symaddr(&addr_QGraphicsItem_setParentItem, "QGraphicsItem_setParentItem")
	symaddr(&addr_QGraphicsItem_childItems, "QGraphicsItem_childItems")
	symaddr(&addr_QGraphicsItem_isWidget, "QGraphicsItem_isWidget")
	symaddr(&addr_QGraphicsItem_isWindow, "QGraphicsItem_isWindow")
	symaddr(&addr_QGraphicsItem_isPanel, "QGraphicsItem_isPanel")
	symaddr(&addr_QGraphicsItem_toGraphicsObject, "QGraphicsItem_toGraphicsObject")
	symaddr(&addr_QGraphicsItem_toGraphicsObject2, "QGraphicsItem_toGraphicsObject2")
	symaddr(&addr_QGraphicsItem_group, "QGraphicsItem_group")
	symaddr(&addr_QGraphicsItem_setGroup, "QGraphicsItem_setGroup")
	symaddr(&addr_QGraphicsItem_flags, "QGraphicsItem_flags")
	symaddr(&addr_QGraphicsItem_setFlag, "QGraphicsItem_setFlag")
	symaddr(&addr_QGraphicsItem_setFlags, "QGraphicsItem_setFlags")
	symaddr(&addr_QGraphicsItem_cacheMode, "QGraphicsItem_cacheMode")
	symaddr(&addr_QGraphicsItem_setCacheMode, "QGraphicsItem_setCacheMode")
	symaddr(&addr_QGraphicsItem_panelModality, "QGraphicsItem_panelModality")
	symaddr(&addr_QGraphicsItem_setPanelModality, "QGraphicsItem_setPanelModality")
	symaddr(&addr_QGraphicsItem_isBlockedByModalPanel, "QGraphicsItem_isBlockedByModalPanel")
	symaddr(&addr_QGraphicsItem_toolTip, "QGraphicsItem_toolTip")
	symaddr(&addr_QGraphicsItem_setToolTip, "QGraphicsItem_setToolTip")
	symaddr(&addr_QGraphicsItem_cursor, "QGraphicsItem_cursor")
	symaddr(&addr_QGraphicsItem_setCursor, "QGraphicsItem_setCursor")
	symaddr(&addr_QGraphicsItem_hasCursor, "QGraphicsItem_hasCursor")
	symaddr(&addr_QGraphicsItem_unsetCursor, "QGraphicsItem_unsetCursor")
	symaddr(&addr_QGraphicsItem_isVisible, "QGraphicsItem_isVisible")
	symaddr(&addr_QGraphicsItem_isVisibleTo, "QGraphicsItem_isVisibleTo")
	symaddr(&addr_QGraphicsItem_setVisible, "QGraphicsItem_setVisible")
	symaddr(&addr_QGraphicsItem_hide, "QGraphicsItem_hide")
	symaddr(&addr_QGraphicsItem_show, "QGraphicsItem_show")
	symaddr(&addr_QGraphicsItem_isEnabled, "QGraphicsItem_isEnabled")
	symaddr(&addr_QGraphicsItem_setEnabled, "QGraphicsItem_setEnabled")
	symaddr(&addr_QGraphicsItem_isSelected, "QGraphicsItem_isSelected")
	symaddr(&addr_QGraphicsItem_setSelected, "QGraphicsItem_setSelected")
	symaddr(&addr_QGraphicsItem_acceptDrops, "QGraphicsItem_acceptDrops")
	symaddr(&addr_QGraphicsItem_setAcceptDrops, "QGraphicsItem_setAcceptDrops")
	reg(&pg_QGraphicsItem_opacity, "QGraphicsItem_opacity")
	reg(&pg_QGraphicsItem_effectiveOpacity, "QGraphicsItem_effectiveOpacity")
	reg(&pg_QGraphicsItem_setOpacity, "QGraphicsItem_setOpacity")
	symaddr(&addr_QGraphicsItem_graphicsEffect, "QGraphicsItem_graphicsEffect")
	symaddr(&addr_QGraphicsItem_setGraphicsEffect, "QGraphicsItem_setGraphicsEffect")
	symaddr(&addr_QGraphicsItem_acceptedMouseButtons, "QGraphicsItem_acceptedMouseButtons")
	symaddr(&addr_QGraphicsItem_setAcceptedMouseButtons, "QGraphicsItem_setAcceptedMouseButtons")
	symaddr(&addr_QGraphicsItem_acceptHoverEvents, "QGraphicsItem_acceptHoverEvents")
	symaddr(&addr_QGraphicsItem_setAcceptHoverEvents, "QGraphicsItem_setAcceptHoverEvents")
	symaddr(&addr_QGraphicsItem_acceptTouchEvents, "QGraphicsItem_acceptTouchEvents")
	symaddr(&addr_QGraphicsItem_setAcceptTouchEvents, "QGraphicsItem_setAcceptTouchEvents")
	symaddr(&addr_QGraphicsItem_filtersChildEvents, "QGraphicsItem_filtersChildEvents")
	symaddr(&addr_QGraphicsItem_setFiltersChildEvents, "QGraphicsItem_setFiltersChildEvents")
	symaddr(&addr_QGraphicsItem_handlesChildEvents, "QGraphicsItem_handlesChildEvents")
	symaddr(&addr_QGraphicsItem_setHandlesChildEvents, "QGraphicsItem_setHandlesChildEvents")
	symaddr(&addr_QGraphicsItem_isActive, "QGraphicsItem_isActive")
	symaddr(&addr_QGraphicsItem_setActive, "QGraphicsItem_setActive")
	symaddr(&addr_QGraphicsItem_hasFocus, "QGraphicsItem_hasFocus")
	symaddr(&addr_QGraphicsItem_setFocus, "QGraphicsItem_setFocus")
	symaddr(&addr_QGraphicsItem_clearFocus, "QGraphicsItem_clearFocus")
	symaddr(&addr_QGraphicsItem_focusProxy, "QGraphicsItem_focusProxy")
	symaddr(&addr_QGraphicsItem_setFocusProxy, "QGraphicsItem_setFocusProxy")
	symaddr(&addr_QGraphicsItem_focusItem, "QGraphicsItem_focusItem")
	symaddr(&addr_QGraphicsItem_focusScopeItem, "QGraphicsItem_focusScopeItem")
	symaddr(&addr_QGraphicsItem_grabMouse, "QGraphicsItem_grabMouse")
	symaddr(&addr_QGraphicsItem_ungrabMouse, "QGraphicsItem_ungrabMouse")
	symaddr(&addr_QGraphicsItem_grabKeyboard, "QGraphicsItem_grabKeyboard")
	symaddr(&addr_QGraphicsItem_ungrabKeyboard, "QGraphicsItem_ungrabKeyboard")
	symaddr(&addr_QGraphicsItem_pos, "QGraphicsItem_pos")
	reg(&pg_QGraphicsItem_x, "QGraphicsItem_x")
	reg(&pg_QGraphicsItem_setX, "QGraphicsItem_setX")
	reg(&pg_QGraphicsItem_y, "QGraphicsItem_y")
	reg(&pg_QGraphicsItem_setY, "QGraphicsItem_setY")
	symaddr(&addr_QGraphicsItem_scenePos, "QGraphicsItem_scenePos")
	symaddr(&addr_QGraphicsItem_setPos, "QGraphicsItem_setPos")
	reg(&pg_QGraphicsItem_setPos2, "QGraphicsItem_setPos2")
	reg(&pg_QGraphicsItem_moveBy, "QGraphicsItem_moveBy")
	symaddr(&addr_QGraphicsItem_ensureVisible, "QGraphicsItem_ensureVisible")
	reg(&pg_QGraphicsItem_ensureVisible2, "QGraphicsItem_ensureVisible2")
	symaddr(&addr_QGraphicsItem_transform, "QGraphicsItem_transform")
	symaddr(&addr_QGraphicsItem_sceneTransform, "QGraphicsItem_sceneTransform")
	symaddr(&addr_QGraphicsItem_deviceTransform, "QGraphicsItem_deviceTransform")
	symaddr(&addr_QGraphicsItem_itemTransform, "QGraphicsItem_itemTransform")
	symaddr(&addr_QGraphicsItem_setTransform, "QGraphicsItem_setTransform")
	symaddr(&addr_QGraphicsItem_resetTransform, "QGraphicsItem_resetTransform")
	reg(&pg_QGraphicsItem_setRotation, "QGraphicsItem_setRotation")
	reg(&pg_QGraphicsItem_rotation, "QGraphicsItem_rotation")
	reg(&pg_QGraphicsItem_setScale, "QGraphicsItem_setScale")
	reg(&pg_QGraphicsItem_scale, "QGraphicsItem_scale")
	symaddr(&addr_QGraphicsItem_transformations, "QGraphicsItem_transformations")
	symaddr(&addr_QGraphicsItem_setTransformations, "QGraphicsItem_setTransformations")
	symaddr(&addr_QGraphicsItem_transformOriginPoint, "QGraphicsItem_transformOriginPoint")
	symaddr(&addr_QGraphicsItem_setTransformOriginPoint, "QGraphicsItem_setTransformOriginPoint")
	reg(&pg_QGraphicsItem_setTransformOriginPoint2, "QGraphicsItem_setTransformOriginPoint2")
	symaddr(&addr_QGraphicsItem_advance, "QGraphicsItem_advance")
	reg(&pg_QGraphicsItem_zValue, "QGraphicsItem_zValue")
	reg(&pg_QGraphicsItem_setZValue, "QGraphicsItem_setZValue")
	symaddr(&addr_QGraphicsItem_stackBefore, "QGraphicsItem_stackBefore")
	symaddr(&addr_QGraphicsItem_boundingRect, "QGraphicsItem_boundingRect")
	symaddr(&addr_QGraphicsItem_childrenBoundingRect, "QGraphicsItem_childrenBoundingRect")
	symaddr(&addr_QGraphicsItem_sceneBoundingRect, "QGraphicsItem_sceneBoundingRect")
	symaddr(&addr_QGraphicsItem_shape, "QGraphicsItem_shape")
	symaddr(&addr_QGraphicsItem_isClipped, "QGraphicsItem_isClipped")
	symaddr(&addr_QGraphicsItem_clipPath, "QGraphicsItem_clipPath")
	symaddr(&addr_QGraphicsItem_contains, "QGraphicsItem_contains")
	symaddr(&addr_QGraphicsItem_collidesWithItem, "QGraphicsItem_collidesWithItem")
	symaddr(&addr_QGraphicsItem_collidesWithPath, "QGraphicsItem_collidesWithPath")
	symaddr(&addr_QGraphicsItem_collidingItems, "QGraphicsItem_collidingItems")
	symaddr(&addr_QGraphicsItem_isObscured, "QGraphicsItem_isObscured")
	reg(&pg_QGraphicsItem_isObscured2, "QGraphicsItem_isObscured2")
	symaddr(&addr_QGraphicsItem_isObscuredBy, "QGraphicsItem_isObscuredBy")
	symaddr(&addr_QGraphicsItem_opaqueArea, "QGraphicsItem_opaqueArea")
	symaddr(&addr_QGraphicsItem_boundingRegion, "QGraphicsItem_boundingRegion")
	reg(&pg_QGraphicsItem_boundingRegionGranularity, "QGraphicsItem_boundingRegionGranularity")
	reg(&pg_QGraphicsItem_setBoundingRegionGranularity, "QGraphicsItem_setBoundingRegionGranularity")
	symaddr(&addr_QGraphicsItem_paint, "QGraphicsItem_paint")
	symaddr(&addr_QGraphicsItem_update, "QGraphicsItem_update")
	reg(&pg_QGraphicsItem_update2, "QGraphicsItem_update2")
	reg(&pg_QGraphicsItem_scroll, "QGraphicsItem_scroll")
	symaddr(&addr_QGraphicsItem_mapToItem, "QGraphicsItem_mapToItem")
	symaddr(&addr_QGraphicsItem_mapToParent, "QGraphicsItem_mapToParent")
	symaddr(&addr_QGraphicsItem_mapToScene, "QGraphicsItem_mapToScene")
	symaddr(&addr_QGraphicsItem_mapToItem2, "QGraphicsItem_mapToItem2")
	symaddr(&addr_QGraphicsItem_mapToParent2, "QGraphicsItem_mapToParent2")
	symaddr(&addr_QGraphicsItem_mapToScene2, "QGraphicsItem_mapToScene2")
	symaddr(&addr_QGraphicsItem_mapRectToItem, "QGraphicsItem_mapRectToItem")
	symaddr(&addr_QGraphicsItem_mapRectToParent, "QGraphicsItem_mapRectToParent")
	symaddr(&addr_QGraphicsItem_mapRectToScene, "QGraphicsItem_mapRectToScene")
	symaddr(&addr_QGraphicsItem_mapToItem3, "QGraphicsItem_mapToItem3")
	symaddr(&addr_QGraphicsItem_mapToParent3, "QGraphicsItem_mapToParent3")
	symaddr(&addr_QGraphicsItem_mapToScene3, "QGraphicsItem_mapToScene3")
	symaddr(&addr_QGraphicsItem_mapToItem4, "QGraphicsItem_mapToItem4")
	symaddr(&addr_QGraphicsItem_mapToParent4, "QGraphicsItem_mapToParent4")
	symaddr(&addr_QGraphicsItem_mapToScene4, "QGraphicsItem_mapToScene4")
	symaddr(&addr_QGraphicsItem_mapFromItem, "QGraphicsItem_mapFromItem")
	symaddr(&addr_QGraphicsItem_mapFromParent, "QGraphicsItem_mapFromParent")
	symaddr(&addr_QGraphicsItem_mapFromScene, "QGraphicsItem_mapFromScene")
	symaddr(&addr_QGraphicsItem_mapFromItem2, "QGraphicsItem_mapFromItem2")
	symaddr(&addr_QGraphicsItem_mapFromParent2, "QGraphicsItem_mapFromParent2")
	symaddr(&addr_QGraphicsItem_mapFromScene2, "QGraphicsItem_mapFromScene2")
	symaddr(&addr_QGraphicsItem_mapRectFromItem, "QGraphicsItem_mapRectFromItem")
	symaddr(&addr_QGraphicsItem_mapRectFromParent, "QGraphicsItem_mapRectFromParent")
	symaddr(&addr_QGraphicsItem_mapRectFromScene, "QGraphicsItem_mapRectFromScene")
	symaddr(&addr_QGraphicsItem_mapFromItem3, "QGraphicsItem_mapFromItem3")
	symaddr(&addr_QGraphicsItem_mapFromParent3, "QGraphicsItem_mapFromParent3")
	symaddr(&addr_QGraphicsItem_mapFromScene3, "QGraphicsItem_mapFromScene3")
	symaddr(&addr_QGraphicsItem_mapFromItem4, "QGraphicsItem_mapFromItem4")
	symaddr(&addr_QGraphicsItem_mapFromParent4, "QGraphicsItem_mapFromParent4")
	symaddr(&addr_QGraphicsItem_mapFromScene4, "QGraphicsItem_mapFromScene4")
	reg(&pg_QGraphicsItem_mapToItem5, "QGraphicsItem_mapToItem5")
	reg(&pg_QGraphicsItem_mapToParent5, "QGraphicsItem_mapToParent5")
	reg(&pg_QGraphicsItem_mapToScene5, "QGraphicsItem_mapToScene5")
	reg(&pg_QGraphicsItem_mapToItem6, "QGraphicsItem_mapToItem6")
	reg(&pg_QGraphicsItem_mapToParent6, "QGraphicsItem_mapToParent6")
	reg(&pg_QGraphicsItem_mapToScene6, "QGraphicsItem_mapToScene6")
	reg(&pg_QGraphicsItem_mapRectToItem2, "QGraphicsItem_mapRectToItem2")
	reg(&pg_QGraphicsItem_mapRectToParent2, "QGraphicsItem_mapRectToParent2")
	reg(&pg_QGraphicsItem_mapRectToScene2, "QGraphicsItem_mapRectToScene2")
	reg(&pg_QGraphicsItem_mapFromItem5, "QGraphicsItem_mapFromItem5")
	reg(&pg_QGraphicsItem_mapFromParent5, "QGraphicsItem_mapFromParent5")
	reg(&pg_QGraphicsItem_mapFromScene5, "QGraphicsItem_mapFromScene5")
	reg(&pg_QGraphicsItem_mapFromItem6, "QGraphicsItem_mapFromItem6")
	reg(&pg_QGraphicsItem_mapFromParent6, "QGraphicsItem_mapFromParent6")
	reg(&pg_QGraphicsItem_mapFromScene6, "QGraphicsItem_mapFromScene6")
	reg(&pg_QGraphicsItem_mapRectFromItem2, "QGraphicsItem_mapRectFromItem2")
	reg(&pg_QGraphicsItem_mapRectFromParent2, "QGraphicsItem_mapRectFromParent2")
	reg(&pg_QGraphicsItem_mapRectFromScene2, "QGraphicsItem_mapRectFromScene2")
	symaddr(&addr_QGraphicsItem_isAncestorOf, "QGraphicsItem_isAncestorOf")
	symaddr(&addr_QGraphicsItem_commonAncestorItem, "QGraphicsItem_commonAncestorItem")
	symaddr(&addr_QGraphicsItem_isUnderMouse, "QGraphicsItem_isUnderMouse")
	symaddr(&addr_QGraphicsItem_data, "QGraphicsItem_data")
	symaddr(&addr_QGraphicsItem_setData, "QGraphicsItem_setData")
	symaddr(&addr_QGraphicsItem_inputMethodHints, "QGraphicsItem_inputMethodHints")
	symaddr(&addr_QGraphicsItem_setInputMethodHints, "QGraphicsItem_setInputMethodHints")
	symaddr(&addr_QGraphicsItem_type, "QGraphicsItem_type")
	symaddr(&addr_QGraphicsItem_installSceneEventFilter, "QGraphicsItem_installSceneEventFilter")
	symaddr(&addr_QGraphicsItem_removeSceneEventFilter, "QGraphicsItem_removeSceneEventFilter")
	symaddr(&addr_QGraphicsItem_updateMicroFocus, "QGraphicsItem_updateMicroFocus")
	symaddr(&addr_QGraphicsItem_addToIndex, "QGraphicsItem_addToIndex")
	symaddr(&addr_QGraphicsItem_removeFromIndex, "QGraphicsItem_removeFromIndex")
	symaddr(&addr_QGraphicsItem_prepareGeometryChange, "QGraphicsItem_prepareGeometryChange")
	symaddr(&addr_QGraphicsItem_setFlag2, "QGraphicsItem_setFlag2")
	symaddr(&addr_QGraphicsItem_setCacheMode2, "QGraphicsItem_setCacheMode2")
	symaddr(&addr_QGraphicsItem_setFocus2, "QGraphicsItem_setFocus2")
	symaddr(&addr_QGraphicsItem_ensureVisible3, "QGraphicsItem_ensureVisible3")
	symaddr(&addr_QGraphicsItem_ensureVisible4, "QGraphicsItem_ensureVisible4")
	symaddr(&addr_QGraphicsItem_ensureVisible5, "QGraphicsItem_ensureVisible5")
	reg(&pg_QGraphicsItem_ensureVisible6, "QGraphicsItem_ensureVisible6")
	reg(&pg_QGraphicsItem_ensureVisible7, "QGraphicsItem_ensureVisible7")
	symaddr(&addr_QGraphicsItem_itemTransform2, "QGraphicsItem_itemTransform2")
	symaddr(&addr_QGraphicsItem_setTransform2, "QGraphicsItem_setTransform2")
	symaddr(&addr_QGraphicsItem_collidingItems2, "QGraphicsItem_collidingItems2")
	symaddr(&addr_QGraphicsItem_isObscured3, "QGraphicsItem_isObscured3")
	symaddr(&addr_QGraphicsItem_update3, "QGraphicsItem_update3")
	reg(&pg_QGraphicsItem_scroll2, "QGraphicsItem_scroll2")
	symaddr(&addr_QGraphicsItem_advance_default, "QGraphicsItem_advance_default")
	symaddr(&addr_QGraphicsItem_shape_default, "QGraphicsItem_shape_default")
	symaddr(&addr_QGraphicsItem_contains_default, "QGraphicsItem_contains_default")
	symaddr(&addr_QGraphicsItem_collidesWithItem_default, "QGraphicsItem_collidesWithItem_default")
	symaddr(&addr_QGraphicsItem_collidesWithPath_default, "QGraphicsItem_collidesWithPath_default")
	symaddr(&addr_QGraphicsItem_isObscuredBy_default, "QGraphicsItem_isObscuredBy_default")
	symaddr(&addr_QGraphicsItem_opaqueArea_default, "QGraphicsItem_opaqueArea_default")
	symaddr(&addr_QGraphicsItem_type_default, "QGraphicsItem_type_default")
	symaddr(&addr_QGraphicsItem_sceneEventFilter_default, "QGraphicsItem_sceneEventFilter_default")
	symaddr(&addr_QGraphicsItem_sceneEvent_default, "QGraphicsItem_sceneEvent_default")
	symaddr(&addr_QGraphicsItem_contextMenuEvent_default, "QGraphicsItem_contextMenuEvent_default")
	symaddr(&addr_QGraphicsItem_dragEnterEvent_default, "QGraphicsItem_dragEnterEvent_default")
	symaddr(&addr_QGraphicsItem_dragLeaveEvent_default, "QGraphicsItem_dragLeaveEvent_default")
	symaddr(&addr_QGraphicsItem_dragMoveEvent_default, "QGraphicsItem_dragMoveEvent_default")
	symaddr(&addr_QGraphicsItem_dropEvent_default, "QGraphicsItem_dropEvent_default")
	symaddr(&addr_QGraphicsItem_focusInEvent_default, "QGraphicsItem_focusInEvent_default")
	symaddr(&addr_QGraphicsItem_focusOutEvent_default, "QGraphicsItem_focusOutEvent_default")
	symaddr(&addr_QGraphicsItem_hoverEnterEvent_default, "QGraphicsItem_hoverEnterEvent_default")
	symaddr(&addr_QGraphicsItem_hoverMoveEvent_default, "QGraphicsItem_hoverMoveEvent_default")
	symaddr(&addr_QGraphicsItem_hoverLeaveEvent_default, "QGraphicsItem_hoverLeaveEvent_default")
	symaddr(&addr_QGraphicsItem_keyPressEvent_default, "QGraphicsItem_keyPressEvent_default")
	symaddr(&addr_QGraphicsItem_keyReleaseEvent_default, "QGraphicsItem_keyReleaseEvent_default")
	symaddr(&addr_QGraphicsItem_mousePressEvent_default, "QGraphicsItem_mousePressEvent_default")
	symaddr(&addr_QGraphicsItem_mouseMoveEvent_default, "QGraphicsItem_mouseMoveEvent_default")
	symaddr(&addr_QGraphicsItem_mouseReleaseEvent_default, "QGraphicsItem_mouseReleaseEvent_default")
	symaddr(&addr_QGraphicsItem_mouseDoubleClickEvent_default, "QGraphicsItem_mouseDoubleClickEvent_default")
	symaddr(&addr_QGraphicsItem_wheelEvent_default, "QGraphicsItem_wheelEvent_default")
	symaddr(&addr_QGraphicsItem_inputMethodEvent_default, "QGraphicsItem_inputMethodEvent_default")
	symaddr(&addr_QGraphicsItem_inputMethodQuery_default, "QGraphicsItem_inputMethodQuery_default")
	symaddr(&addr_QGraphicsItem_itemChange_default, "QGraphicsItem_itemChange_default")
	symaddr(&addr_QGraphicsItem_supportsExtension_default, "QGraphicsItem_supportsExtension_default")
	symaddr(&addr_QGraphicsItem_setExtension_default, "QGraphicsItem_setExtension_default")
	symaddr(&addr_QGraphicsItem_extension_default, "QGraphicsItem_extension_default")
	reg(&pg_QGraphicsItem__set_override, "QGraphicsItem__set_override")
	reg(&pg_QGraphicsItem__Go_new, "QGraphicsItem__Go_new")
	reg(&pg_QGraphicsItem_delete, "QGraphicsItem_delete")
}

type QGraphicsItem struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QGraphicsItem) UnsafeGoID() int64     { return this.vid }
func (this *QGraphicsItem) UnsafeSetGoID(id int64) { this.vid = id }

func bindQGraphicsItem(h uintptr) *QGraphicsItem {
	if h == 0 {
		return nil
	}
	return &QGraphicsItem{h: h}
}

func newownedQGraphicsItem(h uintptr) *QGraphicsItem {
	o := bindQGraphicsItem(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QGraphicsItem).Delete)
	}
	return o
}

func (this *QGraphicsItem) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QGraphicsItem) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQGraphicsItem(h unsafe.Pointer) *QGraphicsItem { return bindQGraphicsItem(uintptr(h)) }
func UnsafeNewQGraphicsItemOwned(h unsafe.Pointer) *QGraphicsItem { return newownedQGraphicsItem(uintptr(h)) }

func (this *QGraphicsItem) UnsafeRelease() { this.h = 0 }

func (this *QGraphicsItem) Scene() *QGraphicsScene {
	__r, _, _ := purego.SyscallN(addr_QGraphicsItem_scene, this.h)
	runtime.KeepAlive(this)
	return bindQGraphicsScene(__r)
}

func (this *QGraphicsItem) ParentItem() *QGraphicsItem {
	__r, _, _ := purego.SyscallN(addr_QGraphicsItem_parentItem, this.h)
	runtime.KeepAlive(this)
	return bindQGraphicsItem(__r)
}

func (this *QGraphicsItem) TopLevelItem() *QGraphicsItem {
	__r, _, _ := purego.SyscallN(addr_QGraphicsItem_topLevelItem, this.h)
	runtime.KeepAlive(this)
	return bindQGraphicsItem(__r)
}

func (this *QGraphicsItem) ParentObject() *QGraphicsObject {
	__r, _, _ := purego.SyscallN(addr_QGraphicsItem_parentObject, this.h)
	runtime.KeepAlive(this)
	return bindQGraphicsObject(__r)
}

func (this *QGraphicsItem) ParentWidget() *QGraphicsWidget {
	__r, _, _ := purego.SyscallN(addr_QGraphicsItem_parentWidget, this.h)
	runtime.KeepAlive(this)
	return bindQGraphicsWidget(__r)
}

func (this *QGraphicsItem) TopLevelWidget() *QGraphicsWidget {
	__r, _, _ := purego.SyscallN(addr_QGraphicsItem_topLevelWidget, this.h)
	runtime.KeepAlive(this)
	return bindQGraphicsWidget(__r)
}

func (this *QGraphicsItem) Window() *QGraphicsWidget {
	__r, _, _ := purego.SyscallN(addr_QGraphicsItem_window, this.h)
	runtime.KeepAlive(this)
	return bindQGraphicsWidget(__r)
}

func (this *QGraphicsItem) Panel() *QGraphicsItem {
	__r, _, _ := purego.SyscallN(addr_QGraphicsItem_panel, this.h)
	runtime.KeepAlive(this)
	return bindQGraphicsItem(__r)
}

func (this *QGraphicsItem) SetParentItem(parent *QGraphicsItem) {
	purego.SyscallN(addr_QGraphicsItem_setParentItem, this.h, uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
}

func (this *QGraphicsItem) ChildItems() []*QGraphicsItem {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGraphicsItem_childItems, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QGraphicsItem, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = bindQGraphicsItem(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QGraphicsItem) IsWidget() bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsItem_isWidget, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGraphicsItem) IsWindow() bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsItem_isWindow, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGraphicsItem) IsPanel() bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsItem_isPanel, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGraphicsItem) ToGraphicsObject() *QGraphicsObject {
	__r, _, _ := purego.SyscallN(addr_QGraphicsItem_toGraphicsObject, this.h)
	runtime.KeepAlive(this)
	return bindQGraphicsObject(__r)
}

func (this *QGraphicsItem) ToGraphicsObject2() *QGraphicsObject {
	__r, _, _ := purego.SyscallN(addr_QGraphicsItem_toGraphicsObject2, this.h)
	runtime.KeepAlive(this)
	return bindQGraphicsObject(__r)
}

func (this *QGraphicsItem) Group() *QGraphicsItemGroup {
	__r, _, _ := purego.SyscallN(addr_QGraphicsItem_group, this.h)
	runtime.KeepAlive(this)
	return bindQGraphicsItemGroup(__r)
}

func (this *QGraphicsItem) SetGroup(group *QGraphicsItemGroup) {
	purego.SyscallN(addr_QGraphicsItem_setGroup, this.h, uintptr(group.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(group)
}

func (this *QGraphicsItem) Flags() QGraphicsItem__GraphicsItemFlag {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsItem_flags, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QGraphicsItem__GraphicsItemFlag(__r)
}

func (this *QGraphicsItem) SetFlag(flag QGraphicsItem__GraphicsItemFlag) {
	purego.SyscallN(addr_QGraphicsItem_setFlag, this.h, uintptr(flag))
	runtime.KeepAlive(this)
}

func (this *QGraphicsItem) SetFlags(flags QGraphicsItem__GraphicsItemFlag) {
	purego.SyscallN(addr_QGraphicsItem_setFlags, this.h, uintptr(flags))
	runtime.KeepAlive(this)
}

func (this *QGraphicsItem) CacheMode() QGraphicsItem__CacheMode {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsItem_cacheMode, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QGraphicsItem__CacheMode(__r)
}

func (this *QGraphicsItem) SetCacheMode(mode QGraphicsItem__CacheMode) {
	purego.SyscallN(addr_QGraphicsItem_setCacheMode, this.h, uintptr(mode))
	runtime.KeepAlive(this)
}

func (this *QGraphicsItem) PanelModality() QGraphicsItem__PanelModality {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsItem_panelModality, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QGraphicsItem__PanelModality(__r)
}

func (this *QGraphicsItem) SetPanelModality(panelModality QGraphicsItem__PanelModality) {
	purego.SyscallN(addr_QGraphicsItem_setPanelModality, this.h, uintptr(panelModality))
	runtime.KeepAlive(this)
}

func (this *QGraphicsItem) IsBlockedByModalPanel() bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsItem_isBlockedByModalPanel, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGraphicsItem) ToolTip() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGraphicsItem_toolTip, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QGraphicsItem) SetToolTip(toolTip string) {
	var toolTip_p uintptr
	if len(toolTip) > 0 {
		toolTip_p = uintptr(unsafe.Pointer(unsafe.StringData(toolTip)))
	}
	purego.SyscallN(addr_QGraphicsItem_setToolTip, this.h, toolTip_p, uintptr(len(toolTip)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(toolTip)
}

func (this *QGraphicsItem) Cursor() *QCursor {
	__r, _, _ := purego.SyscallN(addr_QGraphicsItem_cursor, this.h)
	runtime.KeepAlive(this)
	return newownedQCursor(__r)
}

func (this *QGraphicsItem) SetCursor(cursor *QCursor) {
	purego.SyscallN(addr_QGraphicsItem_setCursor, this.h, uintptr(cursor.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(cursor)
}

func (this *QGraphicsItem) HasCursor() bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsItem_hasCursor, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGraphicsItem) UnsetCursor() {
	purego.SyscallN(addr_QGraphicsItem_unsetCursor, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsItem) IsVisible() bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsItem_isVisible, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGraphicsItem) IsVisibleTo(parent *QGraphicsItem) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsItem_isVisibleTo, this.h, uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QGraphicsItem) SetVisible(visible bool) {
	purego.SyscallN(addr_QGraphicsItem_setVisible, this.h, uintptr(b2i(visible)))
	runtime.KeepAlive(this)
}

func (this *QGraphicsItem) Hide() {
	purego.SyscallN(addr_QGraphicsItem_hide, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsItem) Show() {
	purego.SyscallN(addr_QGraphicsItem_show, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsItem) IsEnabled() bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsItem_isEnabled, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGraphicsItem) SetEnabled(enabled bool) {
	purego.SyscallN(addr_QGraphicsItem_setEnabled, this.h, uintptr(b2i(enabled)))
	runtime.KeepAlive(this)
}

func (this *QGraphicsItem) IsSelected() bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsItem_isSelected, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGraphicsItem) SetSelected(selected bool) {
	purego.SyscallN(addr_QGraphicsItem_setSelected, this.h, uintptr(b2i(selected)))
	runtime.KeepAlive(this)
}

func (this *QGraphicsItem) AcceptDrops() bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsItem_acceptDrops, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGraphicsItem) SetAcceptDrops(on bool) {
	purego.SyscallN(addr_QGraphicsItem_setAcceptDrops, this.h, uintptr(b2i(on)))
	runtime.KeepAlive(this)
}

func (this *QGraphicsItem) Opacity() float64 {
	__r := pg_QGraphicsItem_opacity(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QGraphicsItem) EffectiveOpacity() float64 {
	__r := pg_QGraphicsItem_effectiveOpacity(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QGraphicsItem) SetOpacity(opacity float64) {
	pg_QGraphicsItem_setOpacity(this.h, opacity)
	runtime.KeepAlive(this)
}

func (this *QGraphicsItem) GraphicsEffect() *QGraphicsEffect {
	__r, _, _ := purego.SyscallN(addr_QGraphicsItem_graphicsEffect, this.h)
	runtime.KeepAlive(this)
	return bindQGraphicsEffect(__r)
}

func (this *QGraphicsItem) SetGraphicsEffect(effect *QGraphicsEffect) {
	purego.SyscallN(addr_QGraphicsItem_setGraphicsEffect, this.h, uintptr(effect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(effect)
}

func (this *QGraphicsItem) AcceptedMouseButtons() MouseButton {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsItem_acceptedMouseButtons, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return MouseButton(__r)
}

func (this *QGraphicsItem) SetAcceptedMouseButtons(buttons MouseButton) {
	purego.SyscallN(addr_QGraphicsItem_setAcceptedMouseButtons, this.h, uintptr(buttons))
	runtime.KeepAlive(this)
}

func (this *QGraphicsItem) AcceptHoverEvents() bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsItem_acceptHoverEvents, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGraphicsItem) SetAcceptHoverEvents(enabled bool) {
	purego.SyscallN(addr_QGraphicsItem_setAcceptHoverEvents, this.h, uintptr(b2i(enabled)))
	runtime.KeepAlive(this)
}

func (this *QGraphicsItem) AcceptTouchEvents() bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsItem_acceptTouchEvents, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGraphicsItem) SetAcceptTouchEvents(enabled bool) {
	purego.SyscallN(addr_QGraphicsItem_setAcceptTouchEvents, this.h, uintptr(b2i(enabled)))
	runtime.KeepAlive(this)
}

func (this *QGraphicsItem) FiltersChildEvents() bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsItem_filtersChildEvents, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGraphicsItem) SetFiltersChildEvents(enabled bool) {
	purego.SyscallN(addr_QGraphicsItem_setFiltersChildEvents, this.h, uintptr(b2i(enabled)))
	runtime.KeepAlive(this)
}

func (this *QGraphicsItem) HandlesChildEvents() bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsItem_handlesChildEvents, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGraphicsItem) SetHandlesChildEvents(enabled bool) {
	purego.SyscallN(addr_QGraphicsItem_setHandlesChildEvents, this.h, uintptr(b2i(enabled)))
	runtime.KeepAlive(this)
}

func (this *QGraphicsItem) IsActive() bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsItem_isActive, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGraphicsItem) SetActive(active bool) {
	purego.SyscallN(addr_QGraphicsItem_setActive, this.h, uintptr(b2i(active)))
	runtime.KeepAlive(this)
}

func (this *QGraphicsItem) HasFocus() bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsItem_hasFocus, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGraphicsItem) SetFocus() {
	purego.SyscallN(addr_QGraphicsItem_setFocus, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsItem) ClearFocus() {
	purego.SyscallN(addr_QGraphicsItem_clearFocus, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsItem) FocusProxy() *QGraphicsItem {
	__r, _, _ := purego.SyscallN(addr_QGraphicsItem_focusProxy, this.h)
	runtime.KeepAlive(this)
	return bindQGraphicsItem(__r)
}

func (this *QGraphicsItem) SetFocusProxy(item *QGraphicsItem) {
	purego.SyscallN(addr_QGraphicsItem_setFocusProxy, this.h, uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QGraphicsItem) FocusItem() *QGraphicsItem {
	__r, _, _ := purego.SyscallN(addr_QGraphicsItem_focusItem, this.h)
	runtime.KeepAlive(this)
	return bindQGraphicsItem(__r)
}

func (this *QGraphicsItem) FocusScopeItem() *QGraphicsItem {
	__r, _, _ := purego.SyscallN(addr_QGraphicsItem_focusScopeItem, this.h)
	runtime.KeepAlive(this)
	return bindQGraphicsItem(__r)
}

func (this *QGraphicsItem) GrabMouse() {
	purego.SyscallN(addr_QGraphicsItem_grabMouse, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsItem) UngrabMouse() {
	purego.SyscallN(addr_QGraphicsItem_ungrabMouse, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsItem) GrabKeyboard() {
	purego.SyscallN(addr_QGraphicsItem_grabKeyboard, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsItem) UngrabKeyboard() {
	purego.SyscallN(addr_QGraphicsItem_ungrabKeyboard, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsItem) Pos() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsItem_pos, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QGraphicsItem) X() float64 {
	__r := pg_QGraphicsItem_x(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QGraphicsItem) SetX(x float64) {
	pg_QGraphicsItem_setX(this.h, x)
	runtime.KeepAlive(this)
}

func (this *QGraphicsItem) Y() float64 {
	__r := pg_QGraphicsItem_y(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QGraphicsItem) SetY(y float64) {
	pg_QGraphicsItem_setY(this.h, y)
	runtime.KeepAlive(this)
}

func (this *QGraphicsItem) ScenePos() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsItem_scenePos, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QGraphicsItem) SetPos(pos *QPointF) {
	purego.SyscallN(addr_QGraphicsItem_setPos, this.h, uintptr(pos.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pos)
}

func (this *QGraphicsItem) SetPos2(x float64, y float64) {
	pg_QGraphicsItem_setPos2(this.h, x, y)
	runtime.KeepAlive(this)
}

func (this *QGraphicsItem) MoveBy(dx float64, dy float64) {
	pg_QGraphicsItem_moveBy(this.h, dx, dy)
	runtime.KeepAlive(this)
}

func (this *QGraphicsItem) EnsureVisible() {
	purego.SyscallN(addr_QGraphicsItem_ensureVisible, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsItem) EnsureVisible2(x float64, y float64, w float64, h float64) {
	pg_QGraphicsItem_ensureVisible2(this.h, x, y, w, h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsItem) Transform() *QTransform {
	__r, _, _ := purego.SyscallN(addr_QGraphicsItem_transform, this.h)
	runtime.KeepAlive(this)
	return newownedQTransform(__r)
}

func (this *QGraphicsItem) SceneTransform() *QTransform {
	__r, _, _ := purego.SyscallN(addr_QGraphicsItem_sceneTransform, this.h)
	runtime.KeepAlive(this)
	return newownedQTransform(__r)
}

func (this *QGraphicsItem) DeviceTransform(viewportTransform *QTransform) *QTransform {
	__r, _, _ := purego.SyscallN(addr_QGraphicsItem_deviceTransform, this.h, uintptr(viewportTransform.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(viewportTransform)
	return newownedQTransform(__r)
}

func (this *QGraphicsItem) ItemTransform(other *QGraphicsItem) *QTransform {
	__r, _, _ := purego.SyscallN(addr_QGraphicsItem_itemTransform, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return newownedQTransform(__r)
}

func (this *QGraphicsItem) SetTransform(matrix *QTransform) {
	purego.SyscallN(addr_QGraphicsItem_setTransform, this.h, uintptr(matrix.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(matrix)
}

func (this *QGraphicsItem) ResetTransform() {
	purego.SyscallN(addr_QGraphicsItem_resetTransform, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsItem) SetRotation(angle float64) {
	pg_QGraphicsItem_setRotation(this.h, angle)
	runtime.KeepAlive(this)
}

func (this *QGraphicsItem) Rotation() float64 {
	__r := pg_QGraphicsItem_rotation(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QGraphicsItem) SetScale(scale float64) {
	pg_QGraphicsItem_setScale(this.h, scale)
	runtime.KeepAlive(this)
}

func (this *QGraphicsItem) Scale() float64 {
	__r := pg_QGraphicsItem_scale(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QGraphicsItem) Transformations() []*QGraphicsTransform {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGraphicsItem_transformations, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QGraphicsTransform, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = bindQGraphicsTransform(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QGraphicsItem) SetTransformations(transformations []*QGraphicsTransform) {
	transformations_arr := make([]uintptr, len(transformations))
	for __i, __e := range transformations {
		transformations_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var transformations_p uintptr
	if len(transformations_arr) > 0 {
		transformations_p = uintptr(unsafe.Pointer(&transformations_arr[0]))
	}
	purego.SyscallN(addr_QGraphicsItem_setTransformations, this.h, transformations_p, uintptr(len(transformations)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(transformations)
	runtime.KeepAlive(transformations_arr)
}

func (this *QGraphicsItem) TransformOriginPoint() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsItem_transformOriginPoint, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QGraphicsItem) SetTransformOriginPoint(origin *QPointF) {
	purego.SyscallN(addr_QGraphicsItem_setTransformOriginPoint, this.h, uintptr(origin.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(origin)
}

func (this *QGraphicsItem) SetTransformOriginPoint2(ax float64, ay float64) {
	pg_QGraphicsItem_setTransformOriginPoint2(this.h, ax, ay)
	runtime.KeepAlive(this)
}

func (this *QGraphicsItem) Advance(phase int) {
	purego.SyscallN(addr_QGraphicsItem_advance, this.h, uintptr(phase))
	runtime.KeepAlive(this)
}

func (this *QGraphicsItem) ZValue() float64 {
	__r := pg_QGraphicsItem_zValue(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QGraphicsItem) SetZValue(z float64) {
	pg_QGraphicsItem_setZValue(this.h, z)
	runtime.KeepAlive(this)
}

func (this *QGraphicsItem) StackBefore(sibling *QGraphicsItem) {
	purego.SyscallN(addr_QGraphicsItem_stackBefore, this.h, uintptr(sibling.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sibling)
}

func (this *QGraphicsItem) BoundingRect() *QRectF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsItem_boundingRect, this.h)
	runtime.KeepAlive(this)
	return newownedQRectF(__r)
}

func (this *QGraphicsItem) ChildrenBoundingRect() *QRectF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsItem_childrenBoundingRect, this.h)
	runtime.KeepAlive(this)
	return newownedQRectF(__r)
}

func (this *QGraphicsItem) SceneBoundingRect() *QRectF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsItem_sceneBoundingRect, this.h)
	runtime.KeepAlive(this)
	return newownedQRectF(__r)
}

func (this *QGraphicsItem) Shape() *QPainterPath {
	__r, _, _ := purego.SyscallN(addr_QGraphicsItem_shape, this.h)
	runtime.KeepAlive(this)
	return newownedQPainterPath(__r)
}

func (this *QGraphicsItem) IsClipped() bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsItem_isClipped, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGraphicsItem) ClipPath() *QPainterPath {
	__r, _, _ := purego.SyscallN(addr_QGraphicsItem_clipPath, this.h)
	runtime.KeepAlive(this)
	return newownedQPainterPath(__r)
}

func (this *QGraphicsItem) Contains(point *QPointF) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsItem_contains, this.h, uintptr(point.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(point)
	return __r != 0
}

func (this *QGraphicsItem) CollidesWithItem(other *QGraphicsItem, mode ItemSelectionMode) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsItem_collidesWithItem, this.h, uintptr(other.UnsafePointer()), uintptr(mode))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return __r != 0
}

func (this *QGraphicsItem) CollidesWithPath(path *QPainterPath, mode ItemSelectionMode) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsItem_collidesWithPath, this.h, uintptr(path.UnsafePointer()), uintptr(mode))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(path)
	return __r != 0
}

func (this *QGraphicsItem) CollidingItems() []*QGraphicsItem {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGraphicsItem_collidingItems, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QGraphicsItem, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = bindQGraphicsItem(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QGraphicsItem) IsObscured() bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsItem_isObscured, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGraphicsItem) IsObscured2(x float64, y float64, w float64, h float64) bool {
	__r := pg_QGraphicsItem_isObscured2(this.h, x, y, w, h)
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGraphicsItem) IsObscuredBy(item *QGraphicsItem) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsItem_isObscuredBy, this.h, uintptr(item.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
	return __r != 0
}

func (this *QGraphicsItem) OpaqueArea() *QPainterPath {
	__r, _, _ := purego.SyscallN(addr_QGraphicsItem_opaqueArea, this.h)
	runtime.KeepAlive(this)
	return newownedQPainterPath(__r)
}

func (this *QGraphicsItem) BoundingRegion(itemToDeviceTransform *QTransform) *QRegion {
	__r, _, _ := purego.SyscallN(addr_QGraphicsItem_boundingRegion, this.h, uintptr(itemToDeviceTransform.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(itemToDeviceTransform)
	return newownedQRegion(__r)
}

func (this *QGraphicsItem) BoundingRegionGranularity() float64 {
	__r := pg_QGraphicsItem_boundingRegionGranularity(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QGraphicsItem) SetBoundingRegionGranularity(granularity float64) {
	pg_QGraphicsItem_setBoundingRegionGranularity(this.h, granularity)
	runtime.KeepAlive(this)
}

func (this *QGraphicsItem) Paint(painter *QPainter, option *QStyleOptionGraphicsItem, widget *QWidget) {
	purego.SyscallN(addr_QGraphicsItem_paint, this.h, uintptr(painter.UnsafePointer()), uintptr(option.UnsafePointer()), uintptr(widget.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
	runtime.KeepAlive(option)
	runtime.KeepAlive(widget)
}

func (this *QGraphicsItem) Update() {
	purego.SyscallN(addr_QGraphicsItem_update, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsItem) Update2(x float64, y float64, width float64, height float64) {
	pg_QGraphicsItem_update2(this.h, x, y, width, height)
	runtime.KeepAlive(this)
}

func (this *QGraphicsItem) Scroll(dx float64, dy float64) {
	pg_QGraphicsItem_scroll(this.h, dx, dy)
	runtime.KeepAlive(this)
}

func (this *QGraphicsItem) MapToItem(item *QGraphicsItem, point *QPointF) *QPointF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsItem_mapToItem, this.h, uintptr(item.UnsafePointer()), uintptr(point.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
	runtime.KeepAlive(point)
	return newownedQPointF(__r)
}

func (this *QGraphicsItem) MapToParent(point *QPointF) *QPointF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsItem_mapToParent, this.h, uintptr(point.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(point)
	return newownedQPointF(__r)
}

func (this *QGraphicsItem) MapToScene(point *QPointF) *QPointF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsItem_mapToScene, this.h, uintptr(point.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(point)
	return newownedQPointF(__r)
}

func (this *QGraphicsItem) MapToItem2(item *QGraphicsItem, rect *QRectF) *QPolygonF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsItem_mapToItem2, this.h, uintptr(item.UnsafePointer()), uintptr(rect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
	runtime.KeepAlive(rect)
	return newownedQPolygonF(__r)
}

func (this *QGraphicsItem) MapToParentWithRect(rect *QRectF) *QPolygonF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsItem_mapToParent2, this.h, uintptr(rect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
	return newownedQPolygonF(__r)
}

func (this *QGraphicsItem) MapToSceneWithRect(rect *QRectF) *QPolygonF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsItem_mapToScene2, this.h, uintptr(rect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
	return newownedQPolygonF(__r)
}

func (this *QGraphicsItem) MapRectToItem(item *QGraphicsItem, rect *QRectF) *QRectF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsItem_mapRectToItem, this.h, uintptr(item.UnsafePointer()), uintptr(rect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
	runtime.KeepAlive(rect)
	return newownedQRectF(__r)
}

func (this *QGraphicsItem) MapRectToParent(rect *QRectF) *QRectF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsItem_mapRectToParent, this.h, uintptr(rect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
	return newownedQRectF(__r)
}

func (this *QGraphicsItem) MapRectToScene(rect *QRectF) *QRectF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsItem_mapRectToScene, this.h, uintptr(rect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
	return newownedQRectF(__r)
}

func (this *QGraphicsItem) MapToItem3(item *QGraphicsItem, polygon *QPolygonF) *QPolygonF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsItem_mapToItem3, this.h, uintptr(item.UnsafePointer()), uintptr(polygon.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
	runtime.KeepAlive(polygon)
	return newownedQPolygonF(__r)
}

func (this *QGraphicsItem) MapToParentWithPolygon(polygon *QPolygonF) *QPolygonF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsItem_mapToParent3, this.h, uintptr(polygon.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(polygon)
	return newownedQPolygonF(__r)
}

func (this *QGraphicsItem) MapToSceneWithPolygon(polygon *QPolygonF) *QPolygonF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsItem_mapToScene3, this.h, uintptr(polygon.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(polygon)
	return newownedQPolygonF(__r)
}

func (this *QGraphicsItem) MapToItem4(item *QGraphicsItem, path *QPainterPath) *QPainterPath {
	__r, _, _ := purego.SyscallN(addr_QGraphicsItem_mapToItem4, this.h, uintptr(item.UnsafePointer()), uintptr(path.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
	runtime.KeepAlive(path)
	return newownedQPainterPath(__r)
}

func (this *QGraphicsItem) MapToParentWithPath(path *QPainterPath) *QPainterPath {
	__r, _, _ := purego.SyscallN(addr_QGraphicsItem_mapToParent4, this.h, uintptr(path.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(path)
	return newownedQPainterPath(__r)
}

func (this *QGraphicsItem) MapToSceneWithPath(path *QPainterPath) *QPainterPath {
	__r, _, _ := purego.SyscallN(addr_QGraphicsItem_mapToScene4, this.h, uintptr(path.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(path)
	return newownedQPainterPath(__r)
}

func (this *QGraphicsItem) MapFromItem(item *QGraphicsItem, point *QPointF) *QPointF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsItem_mapFromItem, this.h, uintptr(item.UnsafePointer()), uintptr(point.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
	runtime.KeepAlive(point)
	return newownedQPointF(__r)
}

func (this *QGraphicsItem) MapFromParent(point *QPointF) *QPointF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsItem_mapFromParent, this.h, uintptr(point.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(point)
	return newownedQPointF(__r)
}

func (this *QGraphicsItem) MapFromScene(point *QPointF) *QPointF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsItem_mapFromScene, this.h, uintptr(point.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(point)
	return newownedQPointF(__r)
}

func (this *QGraphicsItem) MapFromItem2(item *QGraphicsItem, rect *QRectF) *QPolygonF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsItem_mapFromItem2, this.h, uintptr(item.UnsafePointer()), uintptr(rect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
	runtime.KeepAlive(rect)
	return newownedQPolygonF(__r)
}

func (this *QGraphicsItem) MapFromParentWithRect(rect *QRectF) *QPolygonF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsItem_mapFromParent2, this.h, uintptr(rect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
	return newownedQPolygonF(__r)
}

func (this *QGraphicsItem) MapFromSceneWithRect(rect *QRectF) *QPolygonF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsItem_mapFromScene2, this.h, uintptr(rect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
	return newownedQPolygonF(__r)
}

func (this *QGraphicsItem) MapRectFromItem(item *QGraphicsItem, rect *QRectF) *QRectF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsItem_mapRectFromItem, this.h, uintptr(item.UnsafePointer()), uintptr(rect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
	runtime.KeepAlive(rect)
	return newownedQRectF(__r)
}

func (this *QGraphicsItem) MapRectFromParent(rect *QRectF) *QRectF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsItem_mapRectFromParent, this.h, uintptr(rect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
	return newownedQRectF(__r)
}

func (this *QGraphicsItem) MapRectFromScene(rect *QRectF) *QRectF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsItem_mapRectFromScene, this.h, uintptr(rect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
	return newownedQRectF(__r)
}

func (this *QGraphicsItem) MapFromItem3(item *QGraphicsItem, polygon *QPolygonF) *QPolygonF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsItem_mapFromItem3, this.h, uintptr(item.UnsafePointer()), uintptr(polygon.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
	runtime.KeepAlive(polygon)
	return newownedQPolygonF(__r)
}

func (this *QGraphicsItem) MapFromParentWithPolygon(polygon *QPolygonF) *QPolygonF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsItem_mapFromParent3, this.h, uintptr(polygon.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(polygon)
	return newownedQPolygonF(__r)
}

func (this *QGraphicsItem) MapFromSceneWithPolygon(polygon *QPolygonF) *QPolygonF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsItem_mapFromScene3, this.h, uintptr(polygon.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(polygon)
	return newownedQPolygonF(__r)
}

func (this *QGraphicsItem) MapFromItem4(item *QGraphicsItem, path *QPainterPath) *QPainterPath {
	__r, _, _ := purego.SyscallN(addr_QGraphicsItem_mapFromItem4, this.h, uintptr(item.UnsafePointer()), uintptr(path.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
	runtime.KeepAlive(path)
	return newownedQPainterPath(__r)
}

func (this *QGraphicsItem) MapFromParentWithPath(path *QPainterPath) *QPainterPath {
	__r, _, _ := purego.SyscallN(addr_QGraphicsItem_mapFromParent4, this.h, uintptr(path.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(path)
	return newownedQPainterPath(__r)
}

func (this *QGraphicsItem) MapFromSceneWithPath(path *QPainterPath) *QPainterPath {
	__r, _, _ := purego.SyscallN(addr_QGraphicsItem_mapFromScene4, this.h, uintptr(path.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(path)
	return newownedQPainterPath(__r)
}

func (this *QGraphicsItem) MapToItem5(item *QGraphicsItem, x float64, y float64) *QPointF {
	__r := pg_QGraphicsItem_mapToItem5(this.h, uintptr(item.UnsafePointer()), x, y)
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
	return newownedQPointF(__r)
}

func (this *QGraphicsItem) MapToParent2(x float64, y float64) *QPointF {
	__r := pg_QGraphicsItem_mapToParent5(this.h, x, y)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QGraphicsItem) MapToScene2(x float64, y float64) *QPointF {
	__r := pg_QGraphicsItem_mapToScene5(this.h, x, y)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QGraphicsItem) MapToItem6(item *QGraphicsItem, x float64, y float64, w float64, h float64) *QPolygonF {
	__r := pg_QGraphicsItem_mapToItem6(this.h, uintptr(item.UnsafePointer()), x, y, w, h)
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
	return newownedQPolygonF(__r)
}

func (this *QGraphicsItem) MapToParent3(x float64, y float64, w float64, h float64) *QPolygonF {
	__r := pg_QGraphicsItem_mapToParent6(this.h, x, y, w, h)
	runtime.KeepAlive(this)
	return newownedQPolygonF(__r)
}

func (this *QGraphicsItem) MapToScene3(x float64, y float64, w float64, h float64) *QPolygonF {
	__r := pg_QGraphicsItem_mapToScene6(this.h, x, y, w, h)
	runtime.KeepAlive(this)
	return newownedQPolygonF(__r)
}

func (this *QGraphicsItem) MapRectToItem2(item *QGraphicsItem, x float64, y float64, w float64, h float64) *QRectF {
	__r := pg_QGraphicsItem_mapRectToItem2(this.h, uintptr(item.UnsafePointer()), x, y, w, h)
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
	return newownedQRectF(__r)
}

func (this *QGraphicsItem) MapRectToParent2(x float64, y float64, w float64, h float64) *QRectF {
	__r := pg_QGraphicsItem_mapRectToParent2(this.h, x, y, w, h)
	runtime.KeepAlive(this)
	return newownedQRectF(__r)
}

func (this *QGraphicsItem) MapRectToScene2(x float64, y float64, w float64, h float64) *QRectF {
	__r := pg_QGraphicsItem_mapRectToScene2(this.h, x, y, w, h)
	runtime.KeepAlive(this)
	return newownedQRectF(__r)
}

func (this *QGraphicsItem) MapFromItem5(item *QGraphicsItem, x float64, y float64) *QPointF {
	__r := pg_QGraphicsItem_mapFromItem5(this.h, uintptr(item.UnsafePointer()), x, y)
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
	return newownedQPointF(__r)
}

func (this *QGraphicsItem) MapFromParent2(x float64, y float64) *QPointF {
	__r := pg_QGraphicsItem_mapFromParent5(this.h, x, y)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QGraphicsItem) MapFromScene2(x float64, y float64) *QPointF {
	__r := pg_QGraphicsItem_mapFromScene5(this.h, x, y)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QGraphicsItem) MapFromItem6(item *QGraphicsItem, x float64, y float64, w float64, h float64) *QPolygonF {
	__r := pg_QGraphicsItem_mapFromItem6(this.h, uintptr(item.UnsafePointer()), x, y, w, h)
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
	return newownedQPolygonF(__r)
}

func (this *QGraphicsItem) MapFromParent3(x float64, y float64, w float64, h float64) *QPolygonF {
	__r := pg_QGraphicsItem_mapFromParent6(this.h, x, y, w, h)
	runtime.KeepAlive(this)
	return newownedQPolygonF(__r)
}

func (this *QGraphicsItem) MapFromScene3(x float64, y float64, w float64, h float64) *QPolygonF {
	__r := pg_QGraphicsItem_mapFromScene6(this.h, x, y, w, h)
	runtime.KeepAlive(this)
	return newownedQPolygonF(__r)
}

func (this *QGraphicsItem) MapRectFromItem2(item *QGraphicsItem, x float64, y float64, w float64, h float64) *QRectF {
	__r := pg_QGraphicsItem_mapRectFromItem2(this.h, uintptr(item.UnsafePointer()), x, y, w, h)
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
	return newownedQRectF(__r)
}

func (this *QGraphicsItem) MapRectFromParent2(x float64, y float64, w float64, h float64) *QRectF {
	__r := pg_QGraphicsItem_mapRectFromParent2(this.h, x, y, w, h)
	runtime.KeepAlive(this)
	return newownedQRectF(__r)
}

func (this *QGraphicsItem) MapRectFromScene2(x float64, y float64, w float64, h float64) *QRectF {
	__r := pg_QGraphicsItem_mapRectFromScene2(this.h, x, y, w, h)
	runtime.KeepAlive(this)
	return newownedQRectF(__r)
}

func (this *QGraphicsItem) IsAncestorOf(child *QGraphicsItem) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsItem_isAncestorOf, this.h, uintptr(child.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(child)
	return __r != 0
}

func (this *QGraphicsItem) CommonAncestorItem(other *QGraphicsItem) *QGraphicsItem {
	__r, _, _ := purego.SyscallN(addr_QGraphicsItem_commonAncestorItem, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return bindQGraphicsItem(__r)
}

func (this *QGraphicsItem) IsUnderMouse() bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsItem_isUnderMouse, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGraphicsItem) Data(key int) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QGraphicsItem_data, this.h, uintptr(key))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QGraphicsItem) SetData(key int, value *QVariant) {
	purego.SyscallN(addr_QGraphicsItem_setData, this.h, uintptr(key), uintptr(value.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(value)
}

func (this *QGraphicsItem) InputMethodHints() InputMethodHint {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsItem_inputMethodHints, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return InputMethodHint(__r)
}

func (this *QGraphicsItem) SetInputMethodHints(hints InputMethodHint) {
	purego.SyscallN(addr_QGraphicsItem_setInputMethodHints, this.h, uintptr(hints))
	runtime.KeepAlive(this)
}

func (this *QGraphicsItem) Type() int {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsItem_type, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QGraphicsItem) InstallSceneEventFilter(filterItem *QGraphicsItem) {
	purego.SyscallN(addr_QGraphicsItem_installSceneEventFilter, this.h, uintptr(filterItem.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(filterItem)
}

func (this *QGraphicsItem) RemoveSceneEventFilter(filterItem *QGraphicsItem) {
	purego.SyscallN(addr_QGraphicsItem_removeSceneEventFilter, this.h, uintptr(filterItem.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(filterItem)
}

func (this *QGraphicsItem) UpdateMicroFocus() {
	purego.SyscallN(addr_QGraphicsItem_updateMicroFocus, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsItem) AddToIndex() {
	purego.SyscallN(addr_QGraphicsItem_addToIndex, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsItem) RemoveFromIndex() {
	purego.SyscallN(addr_QGraphicsItem_removeFromIndex, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsItem) PrepareGeometryChange() {
	purego.SyscallN(addr_QGraphicsItem_prepareGeometryChange, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsItem) SetFlag2(flag QGraphicsItem__GraphicsItemFlag, enabled bool) {
	purego.SyscallN(addr_QGraphicsItem_setFlag2, this.h, uintptr(flag), uintptr(b2i(enabled)))
	runtime.KeepAlive(this)
}

func (this *QGraphicsItem) SetCacheMode2(mode QGraphicsItem__CacheMode, cacheSize *QSize) {
	purego.SyscallN(addr_QGraphicsItem_setCacheMode2, this.h, uintptr(mode), uintptr(cacheSize.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(cacheSize)
}

func (this *QGraphicsItem) SetFocusWithFocusReason(focusReason FocusReason) {
	purego.SyscallN(addr_QGraphicsItem_setFocus2, this.h, uintptr(focusReason))
	runtime.KeepAlive(this)
}

func (this *QGraphicsItem) EnsureVisibleWithRect(rect *QRectF) {
	purego.SyscallN(addr_QGraphicsItem_ensureVisible3, this.h, uintptr(rect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
}

func (this *QGraphicsItem) EnsureVisible3(rect *QRectF, xmargin int) {
	purego.SyscallN(addr_QGraphicsItem_ensureVisible4, this.h, uintptr(rect.UnsafePointer()), uintptr(xmargin))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
}

func (this *QGraphicsItem) EnsureVisible4(rect *QRectF, xmargin int, ymargin int) {
	purego.SyscallN(addr_QGraphicsItem_ensureVisible5, this.h, uintptr(rect.UnsafePointer()), uintptr(xmargin), uintptr(ymargin))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
}

func (this *QGraphicsItem) EnsureVisible5(x float64, y float64, w float64, h float64, xmargin int) {
	pg_QGraphicsItem_ensureVisible6(this.h, x, y, w, h, int32(xmargin))
	runtime.KeepAlive(this)
}

func (this *QGraphicsItem) EnsureVisible6(x float64, y float64, w float64, h float64, xmargin int, ymargin int) {
	pg_QGraphicsItem_ensureVisible7(this.h, x, y, w, h, int32(xmargin), int32(ymargin))
	runtime.KeepAlive(this)
}

func (this *QGraphicsItem) ItemTransform2(other *QGraphicsItem) (*QTransform, bool) {
	var ok_out int32
	__r, _, _ := purego.SyscallN(addr_QGraphicsItem_itemTransform2, this.h, uintptr(other.UnsafePointer()), uintptr(unsafe.Pointer(&ok_out)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return newownedQTransform(__r), ok_out != 0
}

func (this *QGraphicsItem) SetTransform2(matrix *QTransform, combine bool) {
	purego.SyscallN(addr_QGraphicsItem_setTransform2, this.h, uintptr(matrix.UnsafePointer()), uintptr(b2i(combine)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(matrix)
}

func (this *QGraphicsItem) CollidingItemsWithMode(mode ItemSelectionMode) []*QGraphicsItem {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGraphicsItem_collidingItems2, this.h, uintptr(mode), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QGraphicsItem, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = bindQGraphicsItem(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QGraphicsItem) IsObscuredWithRect(rect *QRectF) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsItem_isObscured3, this.h, uintptr(rect.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
	return __r != 0
}

func (this *QGraphicsItem) UpdateWithRect(rect *QRectF) {
	purego.SyscallN(addr_QGraphicsItem_update3, this.h, uintptr(rect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
}

func (this *QGraphicsItem) Scroll2(dx float64, dy float64, rect *QRectF) {
	pg_QGraphicsItem_scroll2(this.h, dx, dy, uintptr(rect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
}

func (this *QGraphicsItem) AdvanceDefault(phase int) {
	purego.SyscallN(addr_QGraphicsItem_advance_default, this.h, uintptr(phase))
	runtime.KeepAlive(this)
}

func (this *QGraphicsItem) ShapeDefault() *QPainterPath {
	__r, _, _ := purego.SyscallN(addr_QGraphicsItem_shape_default, this.h)
	runtime.KeepAlive(this)
	return newownedQPainterPath(__r)
}

func (this *QGraphicsItem) ContainsDefault(point *QPointF) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsItem_contains_default, this.h, uintptr(point.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(point)
	return __r != 0
}

func (this *QGraphicsItem) CollidesWithItemDefault(other *QGraphicsItem, mode ItemSelectionMode) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsItem_collidesWithItem_default, this.h, uintptr(other.UnsafePointer()), uintptr(mode))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return __r != 0
}

func (this *QGraphicsItem) CollidesWithPathDefault(path *QPainterPath, mode ItemSelectionMode) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsItem_collidesWithPath_default, this.h, uintptr(path.UnsafePointer()), uintptr(mode))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(path)
	return __r != 0
}

func (this *QGraphicsItem) IsObscuredByDefault(item *QGraphicsItem) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsItem_isObscuredBy_default, this.h, uintptr(item.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
	return __r != 0
}

func (this *QGraphicsItem) OpaqueAreaDefault() *QPainterPath {
	__r, _, _ := purego.SyscallN(addr_QGraphicsItem_opaqueArea_default, this.h)
	runtime.KeepAlive(this)
	return newownedQPainterPath(__r)
}

func (this *QGraphicsItem) TypeDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsItem_type_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QGraphicsItem) SceneEventFilterDefault(watched *QGraphicsItem, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsItem_sceneEventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QGraphicsItem) SceneEventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsItem_sceneEvent_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QGraphicsItem) ContextMenuEventDefault(event *QGraphicsSceneContextMenuEvent) {
	purego.SyscallN(addr_QGraphicsItem_contextMenuEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsItem) DragEnterEventDefault(event *QGraphicsSceneDragDropEvent) {
	purego.SyscallN(addr_QGraphicsItem_dragEnterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsItem) DragLeaveEventDefault(event *QGraphicsSceneDragDropEvent) {
	purego.SyscallN(addr_QGraphicsItem_dragLeaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsItem) DragMoveEventDefault(event *QGraphicsSceneDragDropEvent) {
	purego.SyscallN(addr_QGraphicsItem_dragMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsItem) DropEventDefault(event *QGraphicsSceneDragDropEvent) {
	purego.SyscallN(addr_QGraphicsItem_dropEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsItem) FocusInEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QGraphicsItem_focusInEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsItem) FocusOutEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QGraphicsItem_focusOutEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsItem) HoverEnterEventDefault(event *QGraphicsSceneHoverEvent) {
	purego.SyscallN(addr_QGraphicsItem_hoverEnterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsItem) HoverMoveEventDefault(event *QGraphicsSceneHoverEvent) {
	purego.SyscallN(addr_QGraphicsItem_hoverMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsItem) HoverLeaveEventDefault(event *QGraphicsSceneHoverEvent) {
	purego.SyscallN(addr_QGraphicsItem_hoverLeaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsItem) KeyPressEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QGraphicsItem_keyPressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsItem) KeyReleaseEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QGraphicsItem_keyReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsItem) MousePressEventDefault(event *QGraphicsSceneMouseEvent) {
	purego.SyscallN(addr_QGraphicsItem_mousePressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsItem) MouseMoveEventDefault(event *QGraphicsSceneMouseEvent) {
	purego.SyscallN(addr_QGraphicsItem_mouseMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsItem) MouseReleaseEventDefault(event *QGraphicsSceneMouseEvent) {
	purego.SyscallN(addr_QGraphicsItem_mouseReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsItem) MouseDoubleClickEventDefault(event *QGraphicsSceneMouseEvent) {
	purego.SyscallN(addr_QGraphicsItem_mouseDoubleClickEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsItem) WheelEventDefault(event *QGraphicsSceneWheelEvent) {
	purego.SyscallN(addr_QGraphicsItem_wheelEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsItem) InputMethodEventDefault(event *QInputMethodEvent) {
	purego.SyscallN(addr_QGraphicsItem_inputMethodEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsItem) InputMethodQueryDefault(query InputMethodQuery) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QGraphicsItem_inputMethodQuery_default, this.h, uintptr(query))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QGraphicsItem) ItemChangeDefault(change QGraphicsItem__GraphicsItemChange, value *QVariant) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QGraphicsItem_itemChange_default, this.h, uintptr(change), uintptr(value.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(value)
	return newownedQVariant(__r)
}

func (this *QGraphicsItem) SupportsExtensionDefault(extension QGraphicsItem__Extension) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsItem_supportsExtension_default, this.h, uintptr(extension))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGraphicsItem) SetExtensionDefault(extension QGraphicsItem__Extension, variant *QVariant) {
	purego.SyscallN(addr_QGraphicsItem_setExtension_default, this.h, uintptr(extension), uintptr(variant.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(variant)
}

func (this *QGraphicsItem) ExtensionDefault(variant *QVariant) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QGraphicsItem_extension_default, this.h, uintptr(variant.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(variant)
	return newownedQVariant(__r)
}

func QGraphicsItem_virtIdx(slot int64) int32 {
	switch slot {
	case 2510008782785038831:
		return 0
	case 6083584261053556819:
		return 1
	case -4830836341480073475:
		return 2
	case 7938019069247444294:
		return 3
	case 5645366511944179768:
		return 4
	case 6469736184362593173:
		return 5
	case 6636575182718753853:
		return 6
	case -7215030271914903671:
		return 7
	case 399370970023989435:
		return 8
	case -6371403872082420691:
		return 9
	case -4584203218241215725:
		return 10
	case -2846371507838587071:
		return 11
	case -6698043222227006015:
		return 12
	case 206122715919917589:
		return 13
	case -3139904158373725006:
		return 14
	case -1781545137352745978:
		return 15
	case 9107058682754591716:
		return 16
	case 975043767033552552:
		return 17
	case -385596010576048993:
		return 18
	case 6932941943906681679:
		return 19
	case -2745191522144316396:
		return 20
	case 7944408482203603984:
		return 21
	case 1023791977679793185:
		return 22
	case 772598475950050299:
		return 23
	case -332090898078543319:
		return 24
	case -4255494004099806581:
		return 25
	case -4415920049297795213:
		return 26
	case -4029073905697023887:
		return 27
	case 1919988340934511316:
		return 28
	case 5438819661942843222:
		return 29
	case 7345853097340709638:
		return 30
	case -1888572168794878830:
		return 31
	case -5729166443651306418:
		return 32
	case -6730181576106171580:
		return 33
	case 2270616358835251226:
		return 34
	}
	return -1
}

func QGraphicsItem_enable(h uintptr, idx int32) { pg_QGraphicsItem__set_override(h, idx) }

func NewQGraphicsItem() *QGraphicsItem {
	id := nextVirtID()
	o := newownedQGraphicsItem(pg_QGraphicsItem__Go_new(id))
	o.UnsafeSetGoID(id)
	registerInstance(id, o.h, QGraphicsItem_virtIdx, QGraphicsItem_enable)
	return o
}

func (this *QGraphicsItem) OnBoundingRect(fn func() *QRectF) {
	registerVirt(this.UnsafeGoID(), 2510008782785038831, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "boundingRect"
}

func (this *QGraphicsItem) OnPaint(fn func(*QPainter, *QStyleOptionGraphicsItem, *QWidget)) {
	registerVirt(this.UnsafeGoID(), 6083584261053556819, func(painter uintptr, option uintptr, widget uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter)), UnsafeNewQStyleOptionGraphicsItem(unsafe.Pointer(option)), UnsafeNewQWidget(unsafe.Pointer(widget))) }) // slot "paint"
}

func (this *QGraphicsItem) OnAdvance(fn func(int)) {
	registerVirt(this.UnsafeGoID(), -4830836341480073475, func(phase int32) { fn(int(phase)) }) // slot "advance"
}

func (this *QGraphicsItem) OnShape(fn func() *QPainterPath) {
	registerVirt(this.UnsafeGoID(), 7938019069247444294, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "shape"
}

func (this *QGraphicsItem) OnContains(fn func(*QPointF) bool) {
	registerVirt(this.UnsafeGoID(), 5645366511944179768, func(point uintptr) int32 { return b2i(fn(UnsafeNewQPointF(unsafe.Pointer(point)))) }) // slot "contains"
}

func (this *QGraphicsItem) OnCollidesWithItem(fn func(*QGraphicsItem, ItemSelectionMode) bool) {
	registerVirt(this.UnsafeGoID(), 6469736184362593173, func(other uintptr, mode int32) int32 { return b2i(fn(UnsafeNewQGraphicsItem(unsafe.Pointer(other)), ItemSelectionMode(mode))) }) // slot "collidesWithItem"
}

func (this *QGraphicsItem) OnCollidesWithPath(fn func(*QPainterPath, ItemSelectionMode) bool) {
	registerVirt(this.UnsafeGoID(), 6636575182718753853, func(path uintptr, mode int32) int32 { return b2i(fn(UnsafeNewQPainterPath(unsafe.Pointer(path)), ItemSelectionMode(mode))) }) // slot "collidesWithPath"
}

func (this *QGraphicsItem) OnIsObscuredBy(fn func(*QGraphicsItem) bool) {
	registerVirt(this.UnsafeGoID(), -7215030271914903671, func(item uintptr) int32 { return b2i(fn(UnsafeNewQGraphicsItem(unsafe.Pointer(item)))) }) // slot "isObscuredBy"
}

func (this *QGraphicsItem) OnOpaqueArea(fn func() *QPainterPath) {
	registerVirt(this.UnsafeGoID(), 399370970023989435, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "opaqueArea"
}

func (this *QGraphicsItem) OnType(fn func() int) {
	registerVirt(this.UnsafeGoID(), -6371403872082420691, func() int32 { return int32(fn()) }) // slot "type"
}

func (this *QGraphicsItem) OnSceneEventFilter(fn func(*QGraphicsItem, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), -4584203218241215725, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQGraphicsItem(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "sceneEventFilter"
}

func (this *QGraphicsItem) OnSceneEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), -2846371507838587071, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "sceneEvent"
}

func (this *QGraphicsItem) OnContextMenuEvent(fn func(*QGraphicsSceneContextMenuEvent)) {
	registerVirt(this.UnsafeGoID(), -6698043222227006015, func(event uintptr) { fn(UnsafeNewQGraphicsSceneContextMenuEvent(unsafe.Pointer(event))) }) // slot "contextMenuEvent"
}

func (this *QGraphicsItem) OnDragEnterEvent(fn func(*QGraphicsSceneDragDropEvent)) {
	registerVirt(this.UnsafeGoID(), 206122715919917589, func(event uintptr) { fn(UnsafeNewQGraphicsSceneDragDropEvent(unsafe.Pointer(event))) }) // slot "dragEnterEvent"
}

func (this *QGraphicsItem) OnDragLeaveEvent(fn func(*QGraphicsSceneDragDropEvent)) {
	registerVirt(this.UnsafeGoID(), -3139904158373725006, func(event uintptr) { fn(UnsafeNewQGraphicsSceneDragDropEvent(unsafe.Pointer(event))) }) // slot "dragLeaveEvent"
}

func (this *QGraphicsItem) OnDragMoveEvent(fn func(*QGraphicsSceneDragDropEvent)) {
	registerVirt(this.UnsafeGoID(), -1781545137352745978, func(event uintptr) { fn(UnsafeNewQGraphicsSceneDragDropEvent(unsafe.Pointer(event))) }) // slot "dragMoveEvent"
}

func (this *QGraphicsItem) OnDropEvent(fn func(*QGraphicsSceneDragDropEvent)) {
	registerVirt(this.UnsafeGoID(), 9107058682754591716, func(event uintptr) { fn(UnsafeNewQGraphicsSceneDragDropEvent(unsafe.Pointer(event))) }) // slot "dropEvent"
}

func (this *QGraphicsItem) OnFocusInEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), 975043767033552552, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusInEvent"
}

func (this *QGraphicsItem) OnFocusOutEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), -385596010576048993, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusOutEvent"
}

func (this *QGraphicsItem) OnHoverEnterEvent(fn func(*QGraphicsSceneHoverEvent)) {
	registerVirt(this.UnsafeGoID(), 6932941943906681679, func(event uintptr) { fn(UnsafeNewQGraphicsSceneHoverEvent(unsafe.Pointer(event))) }) // slot "hoverEnterEvent"
}

func (this *QGraphicsItem) OnHoverMoveEvent(fn func(*QGraphicsSceneHoverEvent)) {
	registerVirt(this.UnsafeGoID(), -2745191522144316396, func(event uintptr) { fn(UnsafeNewQGraphicsSceneHoverEvent(unsafe.Pointer(event))) }) // slot "hoverMoveEvent"
}

func (this *QGraphicsItem) OnHoverLeaveEvent(fn func(*QGraphicsSceneHoverEvent)) {
	registerVirt(this.UnsafeGoID(), 7944408482203603984, func(event uintptr) { fn(UnsafeNewQGraphicsSceneHoverEvent(unsafe.Pointer(event))) }) // slot "hoverLeaveEvent"
}

func (this *QGraphicsItem) OnKeyPressEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 1023791977679793185, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyPressEvent"
}

func (this *QGraphicsItem) OnKeyReleaseEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 772598475950050299, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyReleaseEvent"
}

func (this *QGraphicsItem) OnMousePressEvent(fn func(*QGraphicsSceneMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -332090898078543319, func(event uintptr) { fn(UnsafeNewQGraphicsSceneMouseEvent(unsafe.Pointer(event))) }) // slot "mousePressEvent"
}

func (this *QGraphicsItem) OnMouseMoveEvent(fn func(*QGraphicsSceneMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4255494004099806581, func(event uintptr) { fn(UnsafeNewQGraphicsSceneMouseEvent(unsafe.Pointer(event))) }) // slot "mouseMoveEvent"
}

func (this *QGraphicsItem) OnMouseReleaseEvent(fn func(*QGraphicsSceneMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4415920049297795213, func(event uintptr) { fn(UnsafeNewQGraphicsSceneMouseEvent(unsafe.Pointer(event))) }) // slot "mouseReleaseEvent"
}

func (this *QGraphicsItem) OnMouseDoubleClickEvent(fn func(*QGraphicsSceneMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4029073905697023887, func(event uintptr) { fn(UnsafeNewQGraphicsSceneMouseEvent(unsafe.Pointer(event))) }) // slot "mouseDoubleClickEvent"
}

func (this *QGraphicsItem) OnWheelEvent(fn func(*QGraphicsSceneWheelEvent)) {
	registerVirt(this.UnsafeGoID(), 1919988340934511316, func(event uintptr) { fn(UnsafeNewQGraphicsSceneWheelEvent(unsafe.Pointer(event))) }) // slot "wheelEvent"
}

func (this *QGraphicsItem) OnInputMethodEvent(fn func(*QInputMethodEvent)) {
	registerVirt(this.UnsafeGoID(), 5438819661942843222, func(event uintptr) { fn(UnsafeNewQInputMethodEvent(unsafe.Pointer(event))) }) // slot "inputMethodEvent"
}

func (this *QGraphicsItem) OnInputMethodQuery(fn func(InputMethodQuery) *QVariant) {
	registerVirt(this.UnsafeGoID(), 7345853097340709638, func(query int32) uintptr { return uintptr(fn(InputMethodQuery(query)).UnsafePointer()) }) // slot "inputMethodQuery"
}

func (this *QGraphicsItem) OnItemChange(fn func(QGraphicsItem__GraphicsItemChange, *QVariant) *QVariant) {
	registerVirt(this.UnsafeGoID(), -1888572168794878830, func(change int32, value uintptr) uintptr { return uintptr(fn(QGraphicsItem__GraphicsItemChange(change), UnsafeNewQVariant(unsafe.Pointer(value))).UnsafePointer()) }) // slot "itemChange"
}

func (this *QGraphicsItem) OnSupportsExtension(fn func(QGraphicsItem__Extension) bool) {
	registerVirt(this.UnsafeGoID(), -5729166443651306418, func(extension int32) int32 { return b2i(fn(QGraphicsItem__Extension(extension))) }) // slot "supportsExtension"
}

func (this *QGraphicsItem) OnSetExtension(fn func(QGraphicsItem__Extension, *QVariant)) {
	registerVirt(this.UnsafeGoID(), -6730181576106171580, func(extension int32, variant uintptr) { fn(QGraphicsItem__Extension(extension), UnsafeNewQVariant(unsafe.Pointer(variant))) }) // slot "setExtension"
}

func (this *QGraphicsItem) OnExtension(fn func(*QVariant) *QVariant) {
	registerVirt(this.UnsafeGoID(), 2270616358835251226, func(variant uintptr) uintptr { return uintptr(fn(UnsafeNewQVariant(unsafe.Pointer(variant))).UnsafePointer()) }) // slot "extension"
}

func (this *QGraphicsItem) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QGraphicsItem_delete(this.h)
	}
	this.UnsafeRelease()
}
