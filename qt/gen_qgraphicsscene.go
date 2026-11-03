package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QGraphicsScene_base func(uintptr) uintptr
	addr_QGraphicsScene_new uintptr
	addr_QGraphicsScene_new2 uintptr
	pg_QGraphicsScene_new3 func(int64, float64, float64, float64, float64) uintptr
	addr_QGraphicsScene_new4 uintptr
	addr_QGraphicsScene_new5 uintptr
	pg_QGraphicsScene_new6 func(int64, float64, float64, float64, float64, uintptr) uintptr
	addr_QGraphicsScene_tr uintptr
	addr_QGraphicsScene_sceneRect uintptr
	pg_QGraphicsScene_width func(uintptr) float64
	pg_QGraphicsScene_height func(uintptr) float64
	addr_QGraphicsScene_setSceneRect uintptr
	pg_QGraphicsScene_setSceneRect2 func(uintptr, float64, float64, float64, float64)
	addr_QGraphicsScene_render uintptr
	addr_QGraphicsScene_itemIndexMethod uintptr
	addr_QGraphicsScene_setItemIndexMethod uintptr
	addr_QGraphicsScene_bspTreeDepth uintptr
	addr_QGraphicsScene_setBspTreeDepth uintptr
	addr_QGraphicsScene_itemsBoundingRect uintptr
	addr_QGraphicsScene_items uintptr
	addr_QGraphicsScene_items2 uintptr
	addr_QGraphicsScene_items3 uintptr
	addr_QGraphicsScene_items4 uintptr
	addr_QGraphicsScene_items5 uintptr
	pg_QGraphicsScene_items6 func(uintptr, float64, float64, float64, float64, int32, int32, *int64) uintptr
	addr_QGraphicsScene_collidingItems uintptr
	addr_QGraphicsScene_itemAt uintptr
	pg_QGraphicsScene_itemAt2 func(uintptr, float64, float64, uintptr) uintptr
	addr_QGraphicsScene_selectedItems uintptr
	addr_QGraphicsScene_selectionArea uintptr
	addr_QGraphicsScene_setSelectionArea uintptr
	addr_QGraphicsScene_setSelectionArea2 uintptr
	addr_QGraphicsScene_createItemGroup uintptr
	addr_QGraphicsScene_destroyItemGroup uintptr
	addr_QGraphicsScene_addItem uintptr
	addr_QGraphicsScene_addEllipse uintptr
	addr_QGraphicsScene_addLine uintptr
	addr_QGraphicsScene_addPath uintptr
	addr_QGraphicsScene_addPixmap uintptr
	addr_QGraphicsScene_addPolygon uintptr
	addr_QGraphicsScene_addRect uintptr
	addr_QGraphicsScene_addText uintptr
	addr_QGraphicsScene_addSimpleText uintptr
	addr_QGraphicsScene_addWidget uintptr
	pg_QGraphicsScene_addEllipse2 func(uintptr, float64, float64, float64, float64) uintptr
	pg_QGraphicsScene_addLine2 func(uintptr, float64, float64, float64, float64) uintptr
	pg_QGraphicsScene_addRect2 func(uintptr, float64, float64, float64, float64) uintptr
	addr_QGraphicsScene_removeItem uintptr
	addr_QGraphicsScene_focusItem uintptr
	addr_QGraphicsScene_setFocusItem uintptr
	addr_QGraphicsScene_hasFocus uintptr
	addr_QGraphicsScene_setFocus uintptr
	addr_QGraphicsScene_clearFocus uintptr
	addr_QGraphicsScene_setStickyFocus uintptr
	addr_QGraphicsScene_stickyFocus uintptr
	addr_QGraphicsScene_mouseGrabberItem uintptr
	addr_QGraphicsScene_backgroundBrush uintptr
	addr_QGraphicsScene_setBackgroundBrush uintptr
	addr_QGraphicsScene_foregroundBrush uintptr
	addr_QGraphicsScene_setForegroundBrush uintptr
	addr_QGraphicsScene_inputMethodQuery uintptr
	addr_QGraphicsScene_views uintptr
	pg_QGraphicsScene_update func(uintptr, float64, float64, float64, float64)
	pg_QGraphicsScene_invalidate func(uintptr, float64, float64, float64, float64)
	addr_QGraphicsScene_style uintptr
	addr_QGraphicsScene_setStyle uintptr
	addr_QGraphicsScene_font uintptr
	addr_QGraphicsScene_setFont uintptr
	addr_QGraphicsScene_palette uintptr
	addr_QGraphicsScene_setPalette uintptr
	addr_QGraphicsScene_isActive uintptr
	addr_QGraphicsScene_activePanel uintptr
	addr_QGraphicsScene_setActivePanel uintptr
	addr_QGraphicsScene_activeWindow uintptr
	addr_QGraphicsScene_setActiveWindow uintptr
	addr_QGraphicsScene_sendEvent uintptr
	pg_QGraphicsScene_minimumRenderSize func(uintptr) float64
	pg_QGraphicsScene_setMinimumRenderSize func(uintptr, float64)
	addr_QGraphicsScene_focusOnTouch uintptr
	addr_QGraphicsScene_setFocusOnTouch uintptr
	addr_QGraphicsScene_update2 uintptr
	addr_QGraphicsScene_invalidate2 uintptr
	addr_QGraphicsScene_advance uintptr
	addr_QGraphicsScene_clearSelection uintptr
	addr_QGraphicsScene_clear uintptr
	addr_QGraphicsScene_changed uintptr
	addr_QGraphicsScene_sceneRectChanged uintptr
	addr_QGraphicsScene_selectionChanged uintptr
	addr_QGraphicsScene_focusItemChanged uintptr
	addr_QGraphicsScene_connect uintptr
	addr_QGraphicsScene_connect2 uintptr
	addr_QGraphicsScene_sender uintptr
	addr_QGraphicsScene_senderSignalIndex uintptr
	addr_QGraphicsScene_receivers uintptr
	addr_QGraphicsScene_isSignalConnected uintptr
	addr_QGraphicsScene_tr2 uintptr
	addr_QGraphicsScene_tr3 uintptr
	addr_QGraphicsScene_render2 uintptr
	addr_QGraphicsScene_render3 uintptr
	addr_QGraphicsScene_render4 uintptr
	addr_QGraphicsScene_items7 uintptr
	addr_QGraphicsScene_items8 uintptr
	addr_QGraphicsScene_items9 uintptr
	addr_QGraphicsScene_items10 uintptr
	addr_QGraphicsScene_items11 uintptr
	addr_QGraphicsScene_items12 uintptr
	addr_QGraphicsScene_items13 uintptr
	addr_QGraphicsScene_items14 uintptr
	addr_QGraphicsScene_items15 uintptr
	addr_QGraphicsScene_items16 uintptr
	addr_QGraphicsScene_items17 uintptr
	addr_QGraphicsScene_items18 uintptr
	addr_QGraphicsScene_items19 uintptr
	pg_QGraphicsScene_items20 func(uintptr, float64, float64, float64, float64, int32, int32, uintptr, *int64) uintptr
	addr_QGraphicsScene_collidingItems2 uintptr
	addr_QGraphicsScene_setSelectionArea3 uintptr
	addr_QGraphicsScene_setSelectionArea4 uintptr
	addr_QGraphicsScene_setSelectionArea5 uintptr
	addr_QGraphicsScene_addEllipse3 uintptr
	addr_QGraphicsScene_addEllipse4 uintptr
	addr_QGraphicsScene_addLine3 uintptr
	addr_QGraphicsScene_addPath2 uintptr
	addr_QGraphicsScene_addPath3 uintptr
	addr_QGraphicsScene_addPolygon2 uintptr
	addr_QGraphicsScene_addPolygon3 uintptr
	addr_QGraphicsScene_addRect3 uintptr
	addr_QGraphicsScene_addRect4 uintptr
	addr_QGraphicsScene_addText2 uintptr
	addr_QGraphicsScene_addSimpleText2 uintptr
	addr_QGraphicsScene_addWidget2 uintptr
	pg_QGraphicsScene_addEllipse5 func(uintptr, float64, float64, float64, float64, uintptr) uintptr
	pg_QGraphicsScene_addEllipse6 func(uintptr, float64, float64, float64, float64, uintptr, uintptr) uintptr
	pg_QGraphicsScene_addLine4 func(uintptr, float64, float64, float64, float64, uintptr) uintptr
	pg_QGraphicsScene_addRect5 func(uintptr, float64, float64, float64, float64, uintptr) uintptr
	pg_QGraphicsScene_addRect6 func(uintptr, float64, float64, float64, float64, uintptr, uintptr) uintptr
	addr_QGraphicsScene_setFocusItem2 uintptr
	addr_QGraphicsScene_setFocus2 uintptr
	pg_QGraphicsScene_invalidate3 func(uintptr, float64, float64, float64, float64, int32)
	addr_QGraphicsScene_update3 uintptr
	addr_QGraphicsScene_invalidate4 uintptr
	addr_QGraphicsScene_invalidate5 uintptr
	addr_QGraphicsScene_connect3 uintptr
	addr_QGraphicsScene_connect4 uintptr
	addr_QGraphicsScene_inputMethodQuery_default uintptr
	addr_QGraphicsScene_metaObject_default uintptr
	addr_QGraphicsScene_event_default uintptr
	addr_QGraphicsScene_eventFilter_default uintptr
	addr_QGraphicsScene_contextMenuEvent_default uintptr
	addr_QGraphicsScene_dragEnterEvent_default uintptr
	addr_QGraphicsScene_dragMoveEvent_default uintptr
	addr_QGraphicsScene_dragLeaveEvent_default uintptr
	addr_QGraphicsScene_dropEvent_default uintptr
	addr_QGraphicsScene_focusInEvent_default uintptr
	addr_QGraphicsScene_focusOutEvent_default uintptr
	addr_QGraphicsScene_helpEvent_default uintptr
	addr_QGraphicsScene_keyPressEvent_default uintptr
	addr_QGraphicsScene_keyReleaseEvent_default uintptr
	addr_QGraphicsScene_mousePressEvent_default uintptr
	addr_QGraphicsScene_mouseMoveEvent_default uintptr
	addr_QGraphicsScene_mouseReleaseEvent_default uintptr
	addr_QGraphicsScene_mouseDoubleClickEvent_default uintptr
	addr_QGraphicsScene_wheelEvent_default uintptr
	addr_QGraphicsScene_inputMethodEvent_default uintptr
	addr_QGraphicsScene_drawBackground_default uintptr
	addr_QGraphicsScene_drawForeground_default uintptr
	addr_QGraphicsScene_focusNextPrevChild_default uintptr
	addr_QGraphicsScene_timerEvent_default uintptr
	addr_QGraphicsScene_childEvent_default uintptr
	addr_QGraphicsScene_customEvent_default uintptr
	addr_QGraphicsScene_connectNotify_default uintptr
	addr_QGraphicsScene_disconnectNotify_default uintptr
	pg_QGraphicsScene_connect_changed func(uintptr, int64)
	pg_QGraphicsScene_connect_sceneRectChanged func(uintptr, int64)
	pg_QGraphicsScene_connect_selectionChanged func(uintptr, int64)
	pg_QGraphicsScene_connect_focusItemChanged func(uintptr, int64)
	pg_QGraphicsScene_connect_destroyed func(uintptr, int64)
	pg_QGraphicsScene_connect_destroyed2 func(uintptr, int64)
	pg_QGraphicsScene_connect_objectNameChanged func(uintptr, int64)
	pg_QGraphicsScene__disconnect func(int64)
	pg_QGraphicsScene__set_override func(uintptr, int32)
	pg_QGraphicsScene_delete func(uintptr)
)

func init() {
	reg(&pg_QGraphicsScene_base, "QGraphicsScene_base")
	symaddr(&addr_QGraphicsScene_new, "QGraphicsScene_new")
	symaddr(&addr_QGraphicsScene_new2, "QGraphicsScene_new2")
	reg(&pg_QGraphicsScene_new3, "QGraphicsScene_new3")
	symaddr(&addr_QGraphicsScene_new4, "QGraphicsScene_new4")
	symaddr(&addr_QGraphicsScene_new5, "QGraphicsScene_new5")
	reg(&pg_QGraphicsScene_new6, "QGraphicsScene_new6")
	symaddr(&addr_QGraphicsScene_tr, "QGraphicsScene_tr")
	symaddr(&addr_QGraphicsScene_sceneRect, "QGraphicsScene_sceneRect")
	reg(&pg_QGraphicsScene_width, "QGraphicsScene_width")
	reg(&pg_QGraphicsScene_height, "QGraphicsScene_height")
	symaddr(&addr_QGraphicsScene_setSceneRect, "QGraphicsScene_setSceneRect")
	reg(&pg_QGraphicsScene_setSceneRect2, "QGraphicsScene_setSceneRect2")
	symaddr(&addr_QGraphicsScene_render, "QGraphicsScene_render")
	symaddr(&addr_QGraphicsScene_itemIndexMethod, "QGraphicsScene_itemIndexMethod")
	symaddr(&addr_QGraphicsScene_setItemIndexMethod, "QGraphicsScene_setItemIndexMethod")
	symaddr(&addr_QGraphicsScene_bspTreeDepth, "QGraphicsScene_bspTreeDepth")
	symaddr(&addr_QGraphicsScene_setBspTreeDepth, "QGraphicsScene_setBspTreeDepth")
	symaddr(&addr_QGraphicsScene_itemsBoundingRect, "QGraphicsScene_itemsBoundingRect")
	symaddr(&addr_QGraphicsScene_items, "QGraphicsScene_items")
	symaddr(&addr_QGraphicsScene_items2, "QGraphicsScene_items2")
	symaddr(&addr_QGraphicsScene_items3, "QGraphicsScene_items3")
	symaddr(&addr_QGraphicsScene_items4, "QGraphicsScene_items4")
	symaddr(&addr_QGraphicsScene_items5, "QGraphicsScene_items5")
	reg(&pg_QGraphicsScene_items6, "QGraphicsScene_items6")
	symaddr(&addr_QGraphicsScene_collidingItems, "QGraphicsScene_collidingItems")
	symaddr(&addr_QGraphicsScene_itemAt, "QGraphicsScene_itemAt")
	reg(&pg_QGraphicsScene_itemAt2, "QGraphicsScene_itemAt2")
	symaddr(&addr_QGraphicsScene_selectedItems, "QGraphicsScene_selectedItems")
	symaddr(&addr_QGraphicsScene_selectionArea, "QGraphicsScene_selectionArea")
	symaddr(&addr_QGraphicsScene_setSelectionArea, "QGraphicsScene_setSelectionArea")
	symaddr(&addr_QGraphicsScene_setSelectionArea2, "QGraphicsScene_setSelectionArea2")
	symaddr(&addr_QGraphicsScene_createItemGroup, "QGraphicsScene_createItemGroup")
	symaddr(&addr_QGraphicsScene_destroyItemGroup, "QGraphicsScene_destroyItemGroup")
	symaddr(&addr_QGraphicsScene_addItem, "QGraphicsScene_addItem")
	symaddr(&addr_QGraphicsScene_addEllipse, "QGraphicsScene_addEllipse")
	symaddr(&addr_QGraphicsScene_addLine, "QGraphicsScene_addLine")
	symaddr(&addr_QGraphicsScene_addPath, "QGraphicsScene_addPath")
	symaddr(&addr_QGraphicsScene_addPixmap, "QGraphicsScene_addPixmap")
	symaddr(&addr_QGraphicsScene_addPolygon, "QGraphicsScene_addPolygon")
	symaddr(&addr_QGraphicsScene_addRect, "QGraphicsScene_addRect")
	symaddr(&addr_QGraphicsScene_addText, "QGraphicsScene_addText")
	symaddr(&addr_QGraphicsScene_addSimpleText, "QGraphicsScene_addSimpleText")
	symaddr(&addr_QGraphicsScene_addWidget, "QGraphicsScene_addWidget")
	reg(&pg_QGraphicsScene_addEllipse2, "QGraphicsScene_addEllipse2")
	reg(&pg_QGraphicsScene_addLine2, "QGraphicsScene_addLine2")
	reg(&pg_QGraphicsScene_addRect2, "QGraphicsScene_addRect2")
	symaddr(&addr_QGraphicsScene_removeItem, "QGraphicsScene_removeItem")
	symaddr(&addr_QGraphicsScene_focusItem, "QGraphicsScene_focusItem")
	symaddr(&addr_QGraphicsScene_setFocusItem, "QGraphicsScene_setFocusItem")
	symaddr(&addr_QGraphicsScene_hasFocus, "QGraphicsScene_hasFocus")
	symaddr(&addr_QGraphicsScene_setFocus, "QGraphicsScene_setFocus")
	symaddr(&addr_QGraphicsScene_clearFocus, "QGraphicsScene_clearFocus")
	symaddr(&addr_QGraphicsScene_setStickyFocus, "QGraphicsScene_setStickyFocus")
	symaddr(&addr_QGraphicsScene_stickyFocus, "QGraphicsScene_stickyFocus")
	symaddr(&addr_QGraphicsScene_mouseGrabberItem, "QGraphicsScene_mouseGrabberItem")
	symaddr(&addr_QGraphicsScene_backgroundBrush, "QGraphicsScene_backgroundBrush")
	symaddr(&addr_QGraphicsScene_setBackgroundBrush, "QGraphicsScene_setBackgroundBrush")
	symaddr(&addr_QGraphicsScene_foregroundBrush, "QGraphicsScene_foregroundBrush")
	symaddr(&addr_QGraphicsScene_setForegroundBrush, "QGraphicsScene_setForegroundBrush")
	symaddr(&addr_QGraphicsScene_inputMethodQuery, "QGraphicsScene_inputMethodQuery")
	symaddr(&addr_QGraphicsScene_views, "QGraphicsScene_views")
	reg(&pg_QGraphicsScene_update, "QGraphicsScene_update")
	reg(&pg_QGraphicsScene_invalidate, "QGraphicsScene_invalidate")
	symaddr(&addr_QGraphicsScene_style, "QGraphicsScene_style")
	symaddr(&addr_QGraphicsScene_setStyle, "QGraphicsScene_setStyle")
	symaddr(&addr_QGraphicsScene_font, "QGraphicsScene_font")
	symaddr(&addr_QGraphicsScene_setFont, "QGraphicsScene_setFont")
	symaddr(&addr_QGraphicsScene_palette, "QGraphicsScene_palette")
	symaddr(&addr_QGraphicsScene_setPalette, "QGraphicsScene_setPalette")
	symaddr(&addr_QGraphicsScene_isActive, "QGraphicsScene_isActive")
	symaddr(&addr_QGraphicsScene_activePanel, "QGraphicsScene_activePanel")
	symaddr(&addr_QGraphicsScene_setActivePanel, "QGraphicsScene_setActivePanel")
	symaddr(&addr_QGraphicsScene_activeWindow, "QGraphicsScene_activeWindow")
	symaddr(&addr_QGraphicsScene_setActiveWindow, "QGraphicsScene_setActiveWindow")
	symaddr(&addr_QGraphicsScene_sendEvent, "QGraphicsScene_sendEvent")
	reg(&pg_QGraphicsScene_minimumRenderSize, "QGraphicsScene_minimumRenderSize")
	reg(&pg_QGraphicsScene_setMinimumRenderSize, "QGraphicsScene_setMinimumRenderSize")
	symaddr(&addr_QGraphicsScene_focusOnTouch, "QGraphicsScene_focusOnTouch")
	symaddr(&addr_QGraphicsScene_setFocusOnTouch, "QGraphicsScene_setFocusOnTouch")
	symaddr(&addr_QGraphicsScene_update2, "QGraphicsScene_update2")
	symaddr(&addr_QGraphicsScene_invalidate2, "QGraphicsScene_invalidate2")
	symaddr(&addr_QGraphicsScene_advance, "QGraphicsScene_advance")
	symaddr(&addr_QGraphicsScene_clearSelection, "QGraphicsScene_clearSelection")
	symaddr(&addr_QGraphicsScene_clear, "QGraphicsScene_clear")
	symaddr(&addr_QGraphicsScene_changed, "QGraphicsScene_changed")
	symaddr(&addr_QGraphicsScene_sceneRectChanged, "QGraphicsScene_sceneRectChanged")
	symaddr(&addr_QGraphicsScene_selectionChanged, "QGraphicsScene_selectionChanged")
	symaddr(&addr_QGraphicsScene_focusItemChanged, "QGraphicsScene_focusItemChanged")
	symaddr(&addr_QGraphicsScene_connect, "QGraphicsScene_connect")
	symaddr(&addr_QGraphicsScene_connect2, "QGraphicsScene_connect2")
	symaddr(&addr_QGraphicsScene_sender, "QGraphicsScene_sender")
	symaddr(&addr_QGraphicsScene_senderSignalIndex, "QGraphicsScene_senderSignalIndex")
	symaddr(&addr_QGraphicsScene_receivers, "QGraphicsScene_receivers")
	symaddr(&addr_QGraphicsScene_isSignalConnected, "QGraphicsScene_isSignalConnected")
	symaddr(&addr_QGraphicsScene_tr2, "QGraphicsScene_tr2")
	symaddr(&addr_QGraphicsScene_tr3, "QGraphicsScene_tr3")
	symaddr(&addr_QGraphicsScene_render2, "QGraphicsScene_render2")
	symaddr(&addr_QGraphicsScene_render3, "QGraphicsScene_render3")
	symaddr(&addr_QGraphicsScene_render4, "QGraphicsScene_render4")
	symaddr(&addr_QGraphicsScene_items7, "QGraphicsScene_items7")
	symaddr(&addr_QGraphicsScene_items8, "QGraphicsScene_items8")
	symaddr(&addr_QGraphicsScene_items9, "QGraphicsScene_items9")
	symaddr(&addr_QGraphicsScene_items10, "QGraphicsScene_items10")
	symaddr(&addr_QGraphicsScene_items11, "QGraphicsScene_items11")
	symaddr(&addr_QGraphicsScene_items12, "QGraphicsScene_items12")
	symaddr(&addr_QGraphicsScene_items13, "QGraphicsScene_items13")
	symaddr(&addr_QGraphicsScene_items14, "QGraphicsScene_items14")
	symaddr(&addr_QGraphicsScene_items15, "QGraphicsScene_items15")
	symaddr(&addr_QGraphicsScene_items16, "QGraphicsScene_items16")
	symaddr(&addr_QGraphicsScene_items17, "QGraphicsScene_items17")
	symaddr(&addr_QGraphicsScene_items18, "QGraphicsScene_items18")
	symaddr(&addr_QGraphicsScene_items19, "QGraphicsScene_items19")
	reg(&pg_QGraphicsScene_items20, "QGraphicsScene_items20")
	symaddr(&addr_QGraphicsScene_collidingItems2, "QGraphicsScene_collidingItems2")
	symaddr(&addr_QGraphicsScene_setSelectionArea3, "QGraphicsScene_setSelectionArea3")
	symaddr(&addr_QGraphicsScene_setSelectionArea4, "QGraphicsScene_setSelectionArea4")
	symaddr(&addr_QGraphicsScene_setSelectionArea5, "QGraphicsScene_setSelectionArea5")
	symaddr(&addr_QGraphicsScene_addEllipse3, "QGraphicsScene_addEllipse3")
	symaddr(&addr_QGraphicsScene_addEllipse4, "QGraphicsScene_addEllipse4")
	symaddr(&addr_QGraphicsScene_addLine3, "QGraphicsScene_addLine3")
	symaddr(&addr_QGraphicsScene_addPath2, "QGraphicsScene_addPath2")
	symaddr(&addr_QGraphicsScene_addPath3, "QGraphicsScene_addPath3")
	symaddr(&addr_QGraphicsScene_addPolygon2, "QGraphicsScene_addPolygon2")
	symaddr(&addr_QGraphicsScene_addPolygon3, "QGraphicsScene_addPolygon3")
	symaddr(&addr_QGraphicsScene_addRect3, "QGraphicsScene_addRect3")
	symaddr(&addr_QGraphicsScene_addRect4, "QGraphicsScene_addRect4")
	symaddr(&addr_QGraphicsScene_addText2, "QGraphicsScene_addText2")
	symaddr(&addr_QGraphicsScene_addSimpleText2, "QGraphicsScene_addSimpleText2")
	symaddr(&addr_QGraphicsScene_addWidget2, "QGraphicsScene_addWidget2")
	reg(&pg_QGraphicsScene_addEllipse5, "QGraphicsScene_addEllipse5")
	reg(&pg_QGraphicsScene_addEllipse6, "QGraphicsScene_addEllipse6")
	reg(&pg_QGraphicsScene_addLine4, "QGraphicsScene_addLine4")
	reg(&pg_QGraphicsScene_addRect5, "QGraphicsScene_addRect5")
	reg(&pg_QGraphicsScene_addRect6, "QGraphicsScene_addRect6")
	symaddr(&addr_QGraphicsScene_setFocusItem2, "QGraphicsScene_setFocusItem2")
	symaddr(&addr_QGraphicsScene_setFocus2, "QGraphicsScene_setFocus2")
	reg(&pg_QGraphicsScene_invalidate3, "QGraphicsScene_invalidate3")
	symaddr(&addr_QGraphicsScene_update3, "QGraphicsScene_update3")
	symaddr(&addr_QGraphicsScene_invalidate4, "QGraphicsScene_invalidate4")
	symaddr(&addr_QGraphicsScene_invalidate5, "QGraphicsScene_invalidate5")
	symaddr(&addr_QGraphicsScene_connect3, "QGraphicsScene_connect3")
	symaddr(&addr_QGraphicsScene_connect4, "QGraphicsScene_connect4")
	symaddr(&addr_QGraphicsScene_inputMethodQuery_default, "QGraphicsScene_inputMethodQuery_default")
	symaddr(&addr_QGraphicsScene_metaObject_default, "QGraphicsScene_metaObject_default")
	symaddr(&addr_QGraphicsScene_event_default, "QGraphicsScene_event_default")
	symaddr(&addr_QGraphicsScene_eventFilter_default, "QGraphicsScene_eventFilter_default")
	symaddr(&addr_QGraphicsScene_contextMenuEvent_default, "QGraphicsScene_contextMenuEvent_default")
	symaddr(&addr_QGraphicsScene_dragEnterEvent_default, "QGraphicsScene_dragEnterEvent_default")
	symaddr(&addr_QGraphicsScene_dragMoveEvent_default, "QGraphicsScene_dragMoveEvent_default")
	symaddr(&addr_QGraphicsScene_dragLeaveEvent_default, "QGraphicsScene_dragLeaveEvent_default")
	symaddr(&addr_QGraphicsScene_dropEvent_default, "QGraphicsScene_dropEvent_default")
	symaddr(&addr_QGraphicsScene_focusInEvent_default, "QGraphicsScene_focusInEvent_default")
	symaddr(&addr_QGraphicsScene_focusOutEvent_default, "QGraphicsScene_focusOutEvent_default")
	symaddr(&addr_QGraphicsScene_helpEvent_default, "QGraphicsScene_helpEvent_default")
	symaddr(&addr_QGraphicsScene_keyPressEvent_default, "QGraphicsScene_keyPressEvent_default")
	symaddr(&addr_QGraphicsScene_keyReleaseEvent_default, "QGraphicsScene_keyReleaseEvent_default")
	symaddr(&addr_QGraphicsScene_mousePressEvent_default, "QGraphicsScene_mousePressEvent_default")
	symaddr(&addr_QGraphicsScene_mouseMoveEvent_default, "QGraphicsScene_mouseMoveEvent_default")
	symaddr(&addr_QGraphicsScene_mouseReleaseEvent_default, "QGraphicsScene_mouseReleaseEvent_default")
	symaddr(&addr_QGraphicsScene_mouseDoubleClickEvent_default, "QGraphicsScene_mouseDoubleClickEvent_default")
	symaddr(&addr_QGraphicsScene_wheelEvent_default, "QGraphicsScene_wheelEvent_default")
	symaddr(&addr_QGraphicsScene_inputMethodEvent_default, "QGraphicsScene_inputMethodEvent_default")
	symaddr(&addr_QGraphicsScene_drawBackground_default, "QGraphicsScene_drawBackground_default")
	symaddr(&addr_QGraphicsScene_drawForeground_default, "QGraphicsScene_drawForeground_default")
	symaddr(&addr_QGraphicsScene_focusNextPrevChild_default, "QGraphicsScene_focusNextPrevChild_default")
	symaddr(&addr_QGraphicsScene_timerEvent_default, "QGraphicsScene_timerEvent_default")
	symaddr(&addr_QGraphicsScene_childEvent_default, "QGraphicsScene_childEvent_default")
	symaddr(&addr_QGraphicsScene_customEvent_default, "QGraphicsScene_customEvent_default")
	symaddr(&addr_QGraphicsScene_connectNotify_default, "QGraphicsScene_connectNotify_default")
	symaddr(&addr_QGraphicsScene_disconnectNotify_default, "QGraphicsScene_disconnectNotify_default")
	reg(&pg_QGraphicsScene_connect_changed, "QGraphicsScene_connect_changed")
	reg(&pg_QGraphicsScene_connect_sceneRectChanged, "QGraphicsScene_connect_sceneRectChanged")
	reg(&pg_QGraphicsScene_connect_selectionChanged, "QGraphicsScene_connect_selectionChanged")
	reg(&pg_QGraphicsScene_connect_focusItemChanged, "QGraphicsScene_connect_focusItemChanged")
	reg(&pg_QGraphicsScene_connect_destroyed, "QGraphicsScene_connect_destroyed")
	reg(&pg_QGraphicsScene_connect_destroyed2, "QGraphicsScene_connect_destroyed2")
	reg(&pg_QGraphicsScene_connect_objectNameChanged, "QGraphicsScene_connect_objectNameChanged")
	reg(&pg_QGraphicsScene__disconnect, "QGraphicsScene__disconnect")
	reg(&pg_QGraphicsScene__set_override, "QGraphicsScene__set_override")
	reg(&pg_QGraphicsScene_delete, "QGraphicsScene_delete")
}

type QGraphicsScene struct {
	h uintptr
	*QObject
}

func bindQGraphicsScene(h uintptr) *QGraphicsScene {
	if h == 0 {
		return nil
	}
	return &QGraphicsScene{h: h, QObject: bindQObject(pg_QGraphicsScene_base(h))}
}

func newownedQGraphicsScene(h uintptr) *QGraphicsScene { return bindQGraphicsScene(h) }

func (this *QGraphicsScene) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QGraphicsScene) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQGraphicsScene(h unsafe.Pointer) *QGraphicsScene { return bindQGraphicsScene(uintptr(h)) }
func UnsafeNewQGraphicsSceneOwned(h unsafe.Pointer) *QGraphicsScene { return newownedQGraphicsScene(uintptr(h)) }

func (this *QGraphicsScene) UnsafeRelease() {
	this.h = 0
	if this.QObject != nil {
		this.QObject.UnsafeRelease()
	}
}

func NewQGraphicsScene() *QGraphicsScene {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QGraphicsScene_new, uintptr(__goid))
	__o := newownedQGraphicsScene(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGraphicsScene_virtIdx, QGraphicsScene_enable)
	return __o
}

func NewQGraphicsScene2(sceneRect *QRectF) *QGraphicsScene {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QGraphicsScene_new2, uintptr(__goid), uintptr(sceneRect.UnsafePointer()))
	__o := newownedQGraphicsScene(__r)
	runtime.KeepAlive(sceneRect)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGraphicsScene_virtIdx, QGraphicsScene_enable)
	return __o
}

func NewQGraphicsScene3(x float64, y float64, width float64, height float64) *QGraphicsScene {
	__goid := nextVirtID()
	__o := newownedQGraphicsScene(pg_QGraphicsScene_new3(__goid, x, y, width, height))
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGraphicsScene_virtIdx, QGraphicsScene_enable)
	return __o
}

func NewQGraphicsScene4(parent *QObject) *QGraphicsScene {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QGraphicsScene_new4, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQGraphicsScene(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGraphicsScene_virtIdx, QGraphicsScene_enable)
	return __o
}

func NewQGraphicsScene5(sceneRect *QRectF, parent *QObject) *QGraphicsScene {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QGraphicsScene_new5, uintptr(__goid), uintptr(sceneRect.UnsafePointer()), uintptr(parent.UnsafePointer()))
	__o := newownedQGraphicsScene(__r)
	runtime.KeepAlive(sceneRect)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGraphicsScene_virtIdx, QGraphicsScene_enable)
	return __o
}

func NewQGraphicsScene6(x float64, y float64, width float64, height float64, parent *QObject) *QGraphicsScene {
	__goid := nextVirtID()
	__o := newownedQGraphicsScene(pg_QGraphicsScene_new6(__goid, x, y, width, height, uintptr(parent.UnsafePointer())))
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGraphicsScene_virtIdx, QGraphicsScene_enable)
	return __o
}

func QGraphicsScene_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGraphicsScene_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QGraphicsScene) SceneRect() *QRectF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsScene_sceneRect, this.h)
	runtime.KeepAlive(this)
	return newownedQRectF(__r)
}

func (this *QGraphicsScene) Width() float64 {
	__r := pg_QGraphicsScene_width(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QGraphicsScene) Height() float64 {
	__r := pg_QGraphicsScene_height(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QGraphicsScene) SetSceneRect(rect *QRectF) {
	purego.SyscallN(addr_QGraphicsScene_setSceneRect, this.h, uintptr(rect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
}

func (this *QGraphicsScene) SetSceneRect2(x float64, y float64, w float64, h float64) {
	pg_QGraphicsScene_setSceneRect2(this.h, x, y, w, h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsScene) Render(painter *QPainter) {
	purego.SyscallN(addr_QGraphicsScene_render, this.h, uintptr(painter.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
}

func (this *QGraphicsScene) ItemIndexMethod() QGraphicsScene__ItemIndexMethod {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsScene_itemIndexMethod, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QGraphicsScene__ItemIndexMethod(__r)
}

func (this *QGraphicsScene) SetItemIndexMethod(method QGraphicsScene__ItemIndexMethod) {
	purego.SyscallN(addr_QGraphicsScene_setItemIndexMethod, this.h, uintptr(method))
	runtime.KeepAlive(this)
}

func (this *QGraphicsScene) BspTreeDepth() int {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsScene_bspTreeDepth, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QGraphicsScene) SetBspTreeDepth(depth int) {
	purego.SyscallN(addr_QGraphicsScene_setBspTreeDepth, this.h, uintptr(depth))
	runtime.KeepAlive(this)
}

func (this *QGraphicsScene) ItemsBoundingRect() *QRectF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsScene_itemsBoundingRect, this.h)
	runtime.KeepAlive(this)
	return newownedQRectF(__r)
}

func (this *QGraphicsScene) Items() []*QGraphicsItem {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGraphicsScene_items, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QGraphicsItem, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = bindQGraphicsItem(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QGraphicsScene) ItemsWithPos(pos *QPointF) []*QGraphicsItem {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGraphicsScene_items2, this.h, uintptr(pos.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pos)
	__out := make([]*QGraphicsItem, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = bindQGraphicsItem(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QGraphicsScene) ItemsWithRect(rect *QRectF) []*QGraphicsItem {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGraphicsScene_items3, this.h, uintptr(rect.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
	__out := make([]*QGraphicsItem, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = bindQGraphicsItem(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QGraphicsScene) ItemsWithPolygon(polygon *QPolygonF) []*QGraphicsItem {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGraphicsScene_items4, this.h, uintptr(polygon.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(polygon)
	__out := make([]*QGraphicsItem, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = bindQGraphicsItem(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QGraphicsScene) ItemsWithPath(path *QPainterPath) []*QGraphicsItem {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGraphicsScene_items5, this.h, uintptr(path.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(path)
	__out := make([]*QGraphicsItem, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = bindQGraphicsItem(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QGraphicsScene) Items2(x float64, y float64, w float64, h float64, mode ItemSelectionMode, order SortOrder) []*QGraphicsItem {
	var __n int64
	__r := pg_QGraphicsScene_items6(this.h, x, y, w, h, int32(mode), int32(order), &__n)
	runtime.KeepAlive(this)
	__out := make([]*QGraphicsItem, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = bindQGraphicsItem(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QGraphicsScene) CollidingItems(item *QGraphicsItem) []*QGraphicsItem {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGraphicsScene_collidingItems, this.h, uintptr(item.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
	__out := make([]*QGraphicsItem, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = bindQGraphicsItem(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QGraphicsScene) ItemAt(pos *QPointF, deviceTransform *QTransform) *QGraphicsItem {
	__r, _, _ := purego.SyscallN(addr_QGraphicsScene_itemAt, this.h, uintptr(pos.UnsafePointer()), uintptr(deviceTransform.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pos)
	runtime.KeepAlive(deviceTransform)
	return bindQGraphicsItem(__r)
}

func (this *QGraphicsScene) ItemAt2(x float64, y float64, deviceTransform *QTransform) *QGraphicsItem {
	__r := pg_QGraphicsScene_itemAt2(this.h, x, y, uintptr(deviceTransform.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(deviceTransform)
	return bindQGraphicsItem(__r)
}

func (this *QGraphicsScene) SelectedItems() []*QGraphicsItem {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGraphicsScene_selectedItems, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QGraphicsItem, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = bindQGraphicsItem(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QGraphicsScene) SelectionArea() *QPainterPath {
	__r, _, _ := purego.SyscallN(addr_QGraphicsScene_selectionArea, this.h)
	runtime.KeepAlive(this)
	return newownedQPainterPath(__r)
}

func (this *QGraphicsScene) SetSelectionArea(path *QPainterPath, deviceTransform *QTransform) {
	purego.SyscallN(addr_QGraphicsScene_setSelectionArea, this.h, uintptr(path.UnsafePointer()), uintptr(deviceTransform.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(path)
	runtime.KeepAlive(deviceTransform)
}

func (this *QGraphicsScene) SetSelectionAreaWithPath(path *QPainterPath) {
	purego.SyscallN(addr_QGraphicsScene_setSelectionArea2, this.h, uintptr(path.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(path)
}

func (this *QGraphicsScene) CreateItemGroup(items []*QGraphicsItem) *QGraphicsItemGroup {
	items_arr := make([]uintptr, len(items))
	for __i, __e := range items {
		items_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var items_p uintptr
	if len(items_arr) > 0 {
		items_p = uintptr(unsafe.Pointer(&items_arr[0]))
	}
	__r, _, _ := purego.SyscallN(addr_QGraphicsScene_createItemGroup, this.h, items_p, uintptr(len(items)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(items)
	runtime.KeepAlive(items_arr)
	return bindQGraphicsItemGroup(__r)
}

func (this *QGraphicsScene) DestroyItemGroup(group *QGraphicsItemGroup) {
	purego.SyscallN(addr_QGraphicsScene_destroyItemGroup, this.h, uintptr(group.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(group)
}

func (this *QGraphicsScene) AddItem(item *QGraphicsItem) {
	purego.SyscallN(addr_QGraphicsScene_addItem, this.h, uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QGraphicsScene) AddEllipse(rect *QRectF) *QGraphicsEllipseItem {
	__r, _, _ := purego.SyscallN(addr_QGraphicsScene_addEllipse, this.h, uintptr(rect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
	return bindQGraphicsEllipseItem(__r)
}

func (this *QGraphicsScene) AddLine(line *QLineF) *QGraphicsLineItem {
	__r, _, _ := purego.SyscallN(addr_QGraphicsScene_addLine, this.h, uintptr(line.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(line)
	return bindQGraphicsLineItem(__r)
}

func (this *QGraphicsScene) AddPath(path *QPainterPath) *QGraphicsPathItem {
	__r, _, _ := purego.SyscallN(addr_QGraphicsScene_addPath, this.h, uintptr(path.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(path)
	return bindQGraphicsPathItem(__r)
}

func (this *QGraphicsScene) AddPixmap(pixmap *QPixmap) *QGraphicsPixmapItem {
	__r, _, _ := purego.SyscallN(addr_QGraphicsScene_addPixmap, this.h, uintptr(pixmap.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pixmap)
	return bindQGraphicsPixmapItem(__r)
}

func (this *QGraphicsScene) AddPolygon(polygon *QPolygonF) *QGraphicsPolygonItem {
	__r, _, _ := purego.SyscallN(addr_QGraphicsScene_addPolygon, this.h, uintptr(polygon.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(polygon)
	return bindQGraphicsPolygonItem(__r)
}

func (this *QGraphicsScene) AddRect(rect *QRectF) *QGraphicsRectItem {
	__r, _, _ := purego.SyscallN(addr_QGraphicsScene_addRect, this.h, uintptr(rect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
	return bindQGraphicsRectItem(__r)
}

func (this *QGraphicsScene) AddText(text string) *QGraphicsTextItem {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QGraphicsScene_addText, this.h, text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
	return bindQGraphicsTextItem(__r)
}

func (this *QGraphicsScene) AddSimpleText(text string) *QGraphicsSimpleTextItem {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QGraphicsScene_addSimpleText, this.h, text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
	return bindQGraphicsSimpleTextItem(__r)
}

func (this *QGraphicsScene) AddWidget(widget *QWidget) *QGraphicsProxyWidget {
	__r, _, _ := purego.SyscallN(addr_QGraphicsScene_addWidget, this.h, uintptr(widget.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(widget)
	return bindQGraphicsProxyWidget(__r)
}

func (this *QGraphicsScene) AddEllipse2(x float64, y float64, w float64, h float64) *QGraphicsEllipseItem {
	__r := pg_QGraphicsScene_addEllipse2(this.h, x, y, w, h)
	runtime.KeepAlive(this)
	return bindQGraphicsEllipseItem(__r)
}

func (this *QGraphicsScene) AddLine2(x1 float64, y1 float64, x2 float64, y2 float64) *QGraphicsLineItem {
	__r := pg_QGraphicsScene_addLine2(this.h, x1, y1, x2, y2)
	runtime.KeepAlive(this)
	return bindQGraphicsLineItem(__r)
}

func (this *QGraphicsScene) AddRect2(x float64, y float64, w float64, h float64) *QGraphicsRectItem {
	__r := pg_QGraphicsScene_addRect2(this.h, x, y, w, h)
	runtime.KeepAlive(this)
	return bindQGraphicsRectItem(__r)
}

func (this *QGraphicsScene) RemoveItem(item *QGraphicsItem) {
	purego.SyscallN(addr_QGraphicsScene_removeItem, this.h, uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QGraphicsScene) FocusItem() *QGraphicsItem {
	__r, _, _ := purego.SyscallN(addr_QGraphicsScene_focusItem, this.h)
	runtime.KeepAlive(this)
	return bindQGraphicsItem(__r)
}

func (this *QGraphicsScene) SetFocusItem(item *QGraphicsItem) {
	purego.SyscallN(addr_QGraphicsScene_setFocusItem, this.h, uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QGraphicsScene) HasFocus() bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsScene_hasFocus, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGraphicsScene) SetFocus() {
	purego.SyscallN(addr_QGraphicsScene_setFocus, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsScene) ClearFocus() {
	purego.SyscallN(addr_QGraphicsScene_clearFocus, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsScene) SetStickyFocus(enabled bool) {
	purego.SyscallN(addr_QGraphicsScene_setStickyFocus, this.h, uintptr(b2i(enabled)))
	runtime.KeepAlive(this)
}

func (this *QGraphicsScene) StickyFocus() bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsScene_stickyFocus, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGraphicsScene) MouseGrabberItem() *QGraphicsItem {
	__r, _, _ := purego.SyscallN(addr_QGraphicsScene_mouseGrabberItem, this.h)
	runtime.KeepAlive(this)
	return bindQGraphicsItem(__r)
}

func (this *QGraphicsScene) BackgroundBrush() *QBrush {
	__r, _, _ := purego.SyscallN(addr_QGraphicsScene_backgroundBrush, this.h)
	runtime.KeepAlive(this)
	return newownedQBrush(__r)
}

func (this *QGraphicsScene) SetBackgroundBrush(brush *QBrush) {
	purego.SyscallN(addr_QGraphicsScene_setBackgroundBrush, this.h, uintptr(brush.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(brush)
}

func (this *QGraphicsScene) ForegroundBrush() *QBrush {
	__r, _, _ := purego.SyscallN(addr_QGraphicsScene_foregroundBrush, this.h)
	runtime.KeepAlive(this)
	return newownedQBrush(__r)
}

func (this *QGraphicsScene) SetForegroundBrush(brush *QBrush) {
	purego.SyscallN(addr_QGraphicsScene_setForegroundBrush, this.h, uintptr(brush.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(brush)
}

func (this *QGraphicsScene) InputMethodQuery(query InputMethodQuery) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QGraphicsScene_inputMethodQuery, this.h, uintptr(query))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QGraphicsScene) Views() []*QGraphicsView {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGraphicsScene_views, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QGraphicsView, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = bindQGraphicsView(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QGraphicsScene) Update(x float64, y float64, w float64, h float64) {
	pg_QGraphicsScene_update(this.h, x, y, w, h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsScene) Invalidate(x float64, y float64, w float64, h float64) {
	pg_QGraphicsScene_invalidate(this.h, x, y, w, h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsScene) Style() *QStyle {
	__r, _, _ := purego.SyscallN(addr_QGraphicsScene_style, this.h)
	runtime.KeepAlive(this)
	return bindQStyle(__r)
}

func (this *QGraphicsScene) SetStyle(style *QStyle) {
	purego.SyscallN(addr_QGraphicsScene_setStyle, this.h, uintptr(style.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(style)
}

func (this *QGraphicsScene) Font() *QFont {
	__r, _, _ := purego.SyscallN(addr_QGraphicsScene_font, this.h)
	runtime.KeepAlive(this)
	return newownedQFont(__r)
}

func (this *QGraphicsScene) SetFont(font *QFont) {
	purego.SyscallN(addr_QGraphicsScene_setFont, this.h, uintptr(font.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(font)
}

func (this *QGraphicsScene) Palette() *QPalette {
	__r, _, _ := purego.SyscallN(addr_QGraphicsScene_palette, this.h)
	runtime.KeepAlive(this)
	return newownedQPalette(__r)
}

func (this *QGraphicsScene) SetPalette(palette *QPalette) {
	purego.SyscallN(addr_QGraphicsScene_setPalette, this.h, uintptr(palette.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(palette)
}

func (this *QGraphicsScene) IsActive() bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsScene_isActive, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGraphicsScene) ActivePanel() *QGraphicsItem {
	__r, _, _ := purego.SyscallN(addr_QGraphicsScene_activePanel, this.h)
	runtime.KeepAlive(this)
	return bindQGraphicsItem(__r)
}

func (this *QGraphicsScene) SetActivePanel(item *QGraphicsItem) {
	purego.SyscallN(addr_QGraphicsScene_setActivePanel, this.h, uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QGraphicsScene) ActiveWindow() *QGraphicsWidget {
	__r, _, _ := purego.SyscallN(addr_QGraphicsScene_activeWindow, this.h)
	runtime.KeepAlive(this)
	return bindQGraphicsWidget(__r)
}

func (this *QGraphicsScene) SetActiveWindow(widget *QGraphicsWidget) {
	purego.SyscallN(addr_QGraphicsScene_setActiveWindow, this.h, uintptr(widget.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(widget)
}

func (this *QGraphicsScene) SendEvent(item *QGraphicsItem, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsScene_sendEvent, this.h, uintptr(item.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QGraphicsScene) MinimumRenderSize() float64 {
	__r := pg_QGraphicsScene_minimumRenderSize(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QGraphicsScene) SetMinimumRenderSize(minSize float64) {
	pg_QGraphicsScene_setMinimumRenderSize(this.h, minSize)
	runtime.KeepAlive(this)
}

func (this *QGraphicsScene) FocusOnTouch() bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsScene_focusOnTouch, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGraphicsScene) SetFocusOnTouch(enabled bool) {
	purego.SyscallN(addr_QGraphicsScene_setFocusOnTouch, this.h, uintptr(b2i(enabled)))
	runtime.KeepAlive(this)
}

func (this *QGraphicsScene) Update2() {
	purego.SyscallN(addr_QGraphicsScene_update2, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsScene) Invalidate2() {
	purego.SyscallN(addr_QGraphicsScene_invalidate2, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsScene) Advance() {
	purego.SyscallN(addr_QGraphicsScene_advance, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsScene) ClearSelection() {
	purego.SyscallN(addr_QGraphicsScene_clearSelection, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsScene) Clear() {
	purego.SyscallN(addr_QGraphicsScene_clear, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsScene) Changed(region []*QRectF) {
	region_arr := make([]uintptr, len(region))
	for __i, __e := range region {
		region_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var region_p uintptr
	if len(region_arr) > 0 {
		region_p = uintptr(unsafe.Pointer(&region_arr[0]))
	}
	purego.SyscallN(addr_QGraphicsScene_changed, this.h, region_p, uintptr(len(region)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(region)
	runtime.KeepAlive(region_arr)
}

func (this *QGraphicsScene) SceneRectChanged(rect *QRectF) {
	purego.SyscallN(addr_QGraphicsScene_sceneRectChanged, this.h, uintptr(rect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
}

func (this *QGraphicsScene) SelectionChanged() {
	purego.SyscallN(addr_QGraphicsScene_selectionChanged, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsScene) FocusItemChanged(newFocus *QGraphicsItem, oldFocus *QGraphicsItem, reason FocusReason) {
	purego.SyscallN(addr_QGraphicsScene_focusItemChanged, this.h, uintptr(newFocus.UnsafePointer()), uintptr(oldFocus.UnsafePointer()), uintptr(reason))
	runtime.KeepAlive(this)
	runtime.KeepAlive(newFocus)
	runtime.KeepAlive(oldFocus)
}

func QGraphicsScene_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QGraphicsScene_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QGraphicsScene_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QGraphicsScene_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QGraphicsScene) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QGraphicsScene_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QGraphicsScene) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsScene_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QGraphicsScene) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QGraphicsScene_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QGraphicsScene) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsScene_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QGraphicsScene_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGraphicsScene_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QGraphicsScene_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGraphicsScene_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func (this *QGraphicsScene) Render2(painter *QPainter, target *QRectF) {
	purego.SyscallN(addr_QGraphicsScene_render2, this.h, uintptr(painter.UnsafePointer()), uintptr(target.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
	runtime.KeepAlive(target)
}

func (this *QGraphicsScene) Render3(painter *QPainter, target *QRectF, source *QRectF) {
	purego.SyscallN(addr_QGraphicsScene_render3, this.h, uintptr(painter.UnsafePointer()), uintptr(target.UnsafePointer()), uintptr(source.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
	runtime.KeepAlive(target)
	runtime.KeepAlive(source)
}

func (this *QGraphicsScene) Render4(painter *QPainter, target *QRectF, source *QRectF, aspectRatioMode AspectRatioMode) {
	purego.SyscallN(addr_QGraphicsScene_render4, this.h, uintptr(painter.UnsafePointer()), uintptr(target.UnsafePointer()), uintptr(source.UnsafePointer()), uintptr(aspectRatioMode))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
	runtime.KeepAlive(target)
	runtime.KeepAlive(source)
}

func (this *QGraphicsScene) ItemsWithOrder(order SortOrder) []*QGraphicsItem {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGraphicsScene_items7, this.h, uintptr(order), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QGraphicsItem, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = bindQGraphicsItem(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QGraphicsScene) Items3(pos *QPointF, mode ItemSelectionMode) []*QGraphicsItem {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGraphicsScene_items8, this.h, uintptr(pos.UnsafePointer()), uintptr(mode), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pos)
	__out := make([]*QGraphicsItem, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = bindQGraphicsItem(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QGraphicsScene) Items4(pos *QPointF, mode ItemSelectionMode, order SortOrder) []*QGraphicsItem {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGraphicsScene_items9, this.h, uintptr(pos.UnsafePointer()), uintptr(mode), uintptr(order), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pos)
	__out := make([]*QGraphicsItem, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = bindQGraphicsItem(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QGraphicsScene) Items5(pos *QPointF, mode ItemSelectionMode, order SortOrder, deviceTransform *QTransform) []*QGraphicsItem {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGraphicsScene_items10, this.h, uintptr(pos.UnsafePointer()), uintptr(mode), uintptr(order), uintptr(deviceTransform.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pos)
	runtime.KeepAlive(deviceTransform)
	__out := make([]*QGraphicsItem, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = bindQGraphicsItem(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QGraphicsScene) Items6(rect *QRectF, mode ItemSelectionMode) []*QGraphicsItem {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGraphicsScene_items11, this.h, uintptr(rect.UnsafePointer()), uintptr(mode), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
	__out := make([]*QGraphicsItem, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = bindQGraphicsItem(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QGraphicsScene) Items7(rect *QRectF, mode ItemSelectionMode, order SortOrder) []*QGraphicsItem {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGraphicsScene_items12, this.h, uintptr(rect.UnsafePointer()), uintptr(mode), uintptr(order), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
	__out := make([]*QGraphicsItem, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = bindQGraphicsItem(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QGraphicsScene) Items8(rect *QRectF, mode ItemSelectionMode, order SortOrder, deviceTransform *QTransform) []*QGraphicsItem {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGraphicsScene_items13, this.h, uintptr(rect.UnsafePointer()), uintptr(mode), uintptr(order), uintptr(deviceTransform.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
	runtime.KeepAlive(deviceTransform)
	__out := make([]*QGraphicsItem, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = bindQGraphicsItem(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QGraphicsScene) Items9(polygon *QPolygonF, mode ItemSelectionMode) []*QGraphicsItem {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGraphicsScene_items14, this.h, uintptr(polygon.UnsafePointer()), uintptr(mode), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(polygon)
	__out := make([]*QGraphicsItem, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = bindQGraphicsItem(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QGraphicsScene) Items10(polygon *QPolygonF, mode ItemSelectionMode, order SortOrder) []*QGraphicsItem {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGraphicsScene_items15, this.h, uintptr(polygon.UnsafePointer()), uintptr(mode), uintptr(order), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(polygon)
	__out := make([]*QGraphicsItem, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = bindQGraphicsItem(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QGraphicsScene) Items11(polygon *QPolygonF, mode ItemSelectionMode, order SortOrder, deviceTransform *QTransform) []*QGraphicsItem {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGraphicsScene_items16, this.h, uintptr(polygon.UnsafePointer()), uintptr(mode), uintptr(order), uintptr(deviceTransform.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(polygon)
	runtime.KeepAlive(deviceTransform)
	__out := make([]*QGraphicsItem, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = bindQGraphicsItem(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QGraphicsScene) Items12(path *QPainterPath, mode ItemSelectionMode) []*QGraphicsItem {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGraphicsScene_items17, this.h, uintptr(path.UnsafePointer()), uintptr(mode), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(path)
	__out := make([]*QGraphicsItem, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = bindQGraphicsItem(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QGraphicsScene) Items13(path *QPainterPath, mode ItemSelectionMode, order SortOrder) []*QGraphicsItem {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGraphicsScene_items18, this.h, uintptr(path.UnsafePointer()), uintptr(mode), uintptr(order), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(path)
	__out := make([]*QGraphicsItem, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = bindQGraphicsItem(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QGraphicsScene) Items14(path *QPainterPath, mode ItemSelectionMode, order SortOrder, deviceTransform *QTransform) []*QGraphicsItem {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGraphicsScene_items19, this.h, uintptr(path.UnsafePointer()), uintptr(mode), uintptr(order), uintptr(deviceTransform.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(path)
	runtime.KeepAlive(deviceTransform)
	__out := make([]*QGraphicsItem, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = bindQGraphicsItem(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QGraphicsScene) Items15(x float64, y float64, w float64, h float64, mode ItemSelectionMode, order SortOrder, deviceTransform *QTransform) []*QGraphicsItem {
	var __n int64
	__r := pg_QGraphicsScene_items20(this.h, x, y, w, h, int32(mode), int32(order), uintptr(deviceTransform.UnsafePointer()), &__n)
	runtime.KeepAlive(this)
	runtime.KeepAlive(deviceTransform)
	__out := make([]*QGraphicsItem, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = bindQGraphicsItem(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QGraphicsScene) CollidingItems2(item *QGraphicsItem, mode ItemSelectionMode) []*QGraphicsItem {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGraphicsScene_collidingItems2, this.h, uintptr(item.UnsafePointer()), uintptr(mode), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
	__out := make([]*QGraphicsItem, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = bindQGraphicsItem(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QGraphicsScene) SetSelectionArea2(path *QPainterPath, selectionOperation ItemSelectionOperation) {
	purego.SyscallN(addr_QGraphicsScene_setSelectionArea3, this.h, uintptr(path.UnsafePointer()), uintptr(selectionOperation))
	runtime.KeepAlive(this)
	runtime.KeepAlive(path)
}

func (this *QGraphicsScene) SetSelectionArea3(path *QPainterPath, selectionOperation ItemSelectionOperation, mode ItemSelectionMode) {
	purego.SyscallN(addr_QGraphicsScene_setSelectionArea4, this.h, uintptr(path.UnsafePointer()), uintptr(selectionOperation), uintptr(mode))
	runtime.KeepAlive(this)
	runtime.KeepAlive(path)
}

func (this *QGraphicsScene) SetSelectionArea4(path *QPainterPath, selectionOperation ItemSelectionOperation, mode ItemSelectionMode, deviceTransform *QTransform) {
	purego.SyscallN(addr_QGraphicsScene_setSelectionArea5, this.h, uintptr(path.UnsafePointer()), uintptr(selectionOperation), uintptr(mode), uintptr(deviceTransform.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(path)
	runtime.KeepAlive(deviceTransform)
}

func (this *QGraphicsScene) AddEllipse3(rect *QRectF, pen *QPen) *QGraphicsEllipseItem {
	__r, _, _ := purego.SyscallN(addr_QGraphicsScene_addEllipse3, this.h, uintptr(rect.UnsafePointer()), uintptr(pen.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
	runtime.KeepAlive(pen)
	return bindQGraphicsEllipseItem(__r)
}

func (this *QGraphicsScene) AddEllipse4(rect *QRectF, pen *QPen, brush *QBrush) *QGraphicsEllipseItem {
	__r, _, _ := purego.SyscallN(addr_QGraphicsScene_addEllipse4, this.h, uintptr(rect.UnsafePointer()), uintptr(pen.UnsafePointer()), uintptr(brush.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
	runtime.KeepAlive(pen)
	runtime.KeepAlive(brush)
	return bindQGraphicsEllipseItem(__r)
}

func (this *QGraphicsScene) AddLine3(line *QLineF, pen *QPen) *QGraphicsLineItem {
	__r, _, _ := purego.SyscallN(addr_QGraphicsScene_addLine3, this.h, uintptr(line.UnsafePointer()), uintptr(pen.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(line)
	runtime.KeepAlive(pen)
	return bindQGraphicsLineItem(__r)
}

func (this *QGraphicsScene) AddPath2(path *QPainterPath, pen *QPen) *QGraphicsPathItem {
	__r, _, _ := purego.SyscallN(addr_QGraphicsScene_addPath2, this.h, uintptr(path.UnsafePointer()), uintptr(pen.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(path)
	runtime.KeepAlive(pen)
	return bindQGraphicsPathItem(__r)
}

func (this *QGraphicsScene) AddPath3(path *QPainterPath, pen *QPen, brush *QBrush) *QGraphicsPathItem {
	__r, _, _ := purego.SyscallN(addr_QGraphicsScene_addPath3, this.h, uintptr(path.UnsafePointer()), uintptr(pen.UnsafePointer()), uintptr(brush.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(path)
	runtime.KeepAlive(pen)
	runtime.KeepAlive(brush)
	return bindQGraphicsPathItem(__r)
}

func (this *QGraphicsScene) AddPolygon2(polygon *QPolygonF, pen *QPen) *QGraphicsPolygonItem {
	__r, _, _ := purego.SyscallN(addr_QGraphicsScene_addPolygon2, this.h, uintptr(polygon.UnsafePointer()), uintptr(pen.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(polygon)
	runtime.KeepAlive(pen)
	return bindQGraphicsPolygonItem(__r)
}

func (this *QGraphicsScene) AddPolygon3(polygon *QPolygonF, pen *QPen, brush *QBrush) *QGraphicsPolygonItem {
	__r, _, _ := purego.SyscallN(addr_QGraphicsScene_addPolygon3, this.h, uintptr(polygon.UnsafePointer()), uintptr(pen.UnsafePointer()), uintptr(brush.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(polygon)
	runtime.KeepAlive(pen)
	runtime.KeepAlive(brush)
	return bindQGraphicsPolygonItem(__r)
}

func (this *QGraphicsScene) AddRect3(rect *QRectF, pen *QPen) *QGraphicsRectItem {
	__r, _, _ := purego.SyscallN(addr_QGraphicsScene_addRect3, this.h, uintptr(rect.UnsafePointer()), uintptr(pen.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
	runtime.KeepAlive(pen)
	return bindQGraphicsRectItem(__r)
}

func (this *QGraphicsScene) AddRect4(rect *QRectF, pen *QPen, brush *QBrush) *QGraphicsRectItem {
	__r, _, _ := purego.SyscallN(addr_QGraphicsScene_addRect4, this.h, uintptr(rect.UnsafePointer()), uintptr(pen.UnsafePointer()), uintptr(brush.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
	runtime.KeepAlive(pen)
	runtime.KeepAlive(brush)
	return bindQGraphicsRectItem(__r)
}

func (this *QGraphicsScene) AddText2(text string, font *QFont) *QGraphicsTextItem {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QGraphicsScene_addText2, this.h, text_p, uintptr(len(text)), uintptr(font.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
	runtime.KeepAlive(font)
	return bindQGraphicsTextItem(__r)
}

func (this *QGraphicsScene) AddSimpleText2(text string, font *QFont) *QGraphicsSimpleTextItem {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QGraphicsScene_addSimpleText2, this.h, text_p, uintptr(len(text)), uintptr(font.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
	runtime.KeepAlive(font)
	return bindQGraphicsSimpleTextItem(__r)
}

func (this *QGraphicsScene) AddWidget2(widget *QWidget, wFlags WindowType) *QGraphicsProxyWidget {
	__r, _, _ := purego.SyscallN(addr_QGraphicsScene_addWidget2, this.h, uintptr(widget.UnsafePointer()), uintptr(wFlags))
	runtime.KeepAlive(this)
	runtime.KeepAlive(widget)
	return bindQGraphicsProxyWidget(__r)
}

func (this *QGraphicsScene) AddEllipse5(x float64, y float64, w float64, h float64, pen *QPen) *QGraphicsEllipseItem {
	__r := pg_QGraphicsScene_addEllipse5(this.h, x, y, w, h, uintptr(pen.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pen)
	return bindQGraphicsEllipseItem(__r)
}

func (this *QGraphicsScene) AddEllipse6(x float64, y float64, w float64, h float64, pen *QPen, brush *QBrush) *QGraphicsEllipseItem {
	__r := pg_QGraphicsScene_addEllipse6(this.h, x, y, w, h, uintptr(pen.UnsafePointer()), uintptr(brush.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pen)
	runtime.KeepAlive(brush)
	return bindQGraphicsEllipseItem(__r)
}

func (this *QGraphicsScene) AddLine4(x1 float64, y1 float64, x2 float64, y2 float64, pen *QPen) *QGraphicsLineItem {
	__r := pg_QGraphicsScene_addLine4(this.h, x1, y1, x2, y2, uintptr(pen.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pen)
	return bindQGraphicsLineItem(__r)
}

func (this *QGraphicsScene) AddRect5(x float64, y float64, w float64, h float64, pen *QPen) *QGraphicsRectItem {
	__r := pg_QGraphicsScene_addRect5(this.h, x, y, w, h, uintptr(pen.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pen)
	return bindQGraphicsRectItem(__r)
}

func (this *QGraphicsScene) AddRect6(x float64, y float64, w float64, h float64, pen *QPen, brush *QBrush) *QGraphicsRectItem {
	__r := pg_QGraphicsScene_addRect6(this.h, x, y, w, h, uintptr(pen.UnsafePointer()), uintptr(brush.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pen)
	runtime.KeepAlive(brush)
	return bindQGraphicsRectItem(__r)
}

func (this *QGraphicsScene) SetFocusItem2(item *QGraphicsItem, focusReason FocusReason) {
	purego.SyscallN(addr_QGraphicsScene_setFocusItem2, this.h, uintptr(item.UnsafePointer()), uintptr(focusReason))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QGraphicsScene) SetFocusWithFocusReason(focusReason FocusReason) {
	purego.SyscallN(addr_QGraphicsScene_setFocus2, this.h, uintptr(focusReason))
	runtime.KeepAlive(this)
}

func (this *QGraphicsScene) Invalidate3(x float64, y float64, w float64, h float64, layers QGraphicsScene__SceneLayer) {
	pg_QGraphicsScene_invalidate3(this.h, x, y, w, h, int32(layers))
	runtime.KeepAlive(this)
}

func (this *QGraphicsScene) UpdateWithRect(rect *QRectF) {
	purego.SyscallN(addr_QGraphicsScene_update3, this.h, uintptr(rect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
}

func (this *QGraphicsScene) InvalidateWithRect(rect *QRectF) {
	purego.SyscallN(addr_QGraphicsScene_invalidate4, this.h, uintptr(rect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
}

func (this *QGraphicsScene) Invalidate4(rect *QRectF, layers QGraphicsScene__SceneLayer) {
	purego.SyscallN(addr_QGraphicsScene_invalidate5, this.h, uintptr(rect.UnsafePointer()), uintptr(layers))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
}

func QGraphicsScene_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QGraphicsScene_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QGraphicsScene_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QGraphicsScene_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QGraphicsScene) InputMethodQueryDefault(query InputMethodQuery) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QGraphicsScene_inputMethodQuery_default, this.h, uintptr(query))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QGraphicsScene) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QGraphicsScene_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QGraphicsScene) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsScene_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QGraphicsScene) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsScene_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QGraphicsScene) ContextMenuEventDefault(event *QGraphicsSceneContextMenuEvent) {
	purego.SyscallN(addr_QGraphicsScene_contextMenuEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsScene) DragEnterEventDefault(event *QGraphicsSceneDragDropEvent) {
	purego.SyscallN(addr_QGraphicsScene_dragEnterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsScene) DragMoveEventDefault(event *QGraphicsSceneDragDropEvent) {
	purego.SyscallN(addr_QGraphicsScene_dragMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsScene) DragLeaveEventDefault(event *QGraphicsSceneDragDropEvent) {
	purego.SyscallN(addr_QGraphicsScene_dragLeaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsScene) DropEventDefault(event *QGraphicsSceneDragDropEvent) {
	purego.SyscallN(addr_QGraphicsScene_dropEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsScene) FocusInEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QGraphicsScene_focusInEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsScene) FocusOutEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QGraphicsScene_focusOutEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsScene) HelpEventDefault(event *QGraphicsSceneHelpEvent) {
	purego.SyscallN(addr_QGraphicsScene_helpEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsScene) KeyPressEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QGraphicsScene_keyPressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsScene) KeyReleaseEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QGraphicsScene_keyReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsScene) MousePressEventDefault(event *QGraphicsSceneMouseEvent) {
	purego.SyscallN(addr_QGraphicsScene_mousePressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsScene) MouseMoveEventDefault(event *QGraphicsSceneMouseEvent) {
	purego.SyscallN(addr_QGraphicsScene_mouseMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsScene) MouseReleaseEventDefault(event *QGraphicsSceneMouseEvent) {
	purego.SyscallN(addr_QGraphicsScene_mouseReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsScene) MouseDoubleClickEventDefault(event *QGraphicsSceneMouseEvent) {
	purego.SyscallN(addr_QGraphicsScene_mouseDoubleClickEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsScene) WheelEventDefault(event *QGraphicsSceneWheelEvent) {
	purego.SyscallN(addr_QGraphicsScene_wheelEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsScene) InputMethodEventDefault(event *QInputMethodEvent) {
	purego.SyscallN(addr_QGraphicsScene_inputMethodEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsScene) DrawBackgroundDefault(painter *QPainter, rect *QRectF) {
	purego.SyscallN(addr_QGraphicsScene_drawBackground_default, this.h, uintptr(painter.UnsafePointer()), uintptr(rect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
	runtime.KeepAlive(rect)
}

func (this *QGraphicsScene) DrawForegroundDefault(painter *QPainter, rect *QRectF) {
	purego.SyscallN(addr_QGraphicsScene_drawForeground_default, this.h, uintptr(painter.UnsafePointer()), uintptr(rect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
	runtime.KeepAlive(rect)
}

func (this *QGraphicsScene) FocusNextPrevChildDefault(next bool) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsScene_focusNextPrevChild_default, this.h, uintptr(b2i(next)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGraphicsScene) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QGraphicsScene_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsScene) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QGraphicsScene_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsScene) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QGraphicsScene_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsScene) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QGraphicsScene_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QGraphicsScene) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QGraphicsScene_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QGraphicsScene) OnChanged(fn func([]*QRectF)) int64 {
	id := registerCallback(func(region uintptr, region_len int64) { fn(func() []*QRectF { if region == 0 || region_len == 0 { return nil }; __sl := unsafe.Slice((*unsafe.Pointer)(unsafe.Pointer(region)), int(region_len)); __o := make([]*QRectF, int(region_len)); for __i := range __o { __o[__i] = UnsafeNewQRectF(__sl[__i]) }; return __o }()) })
	pg_QGraphicsScene_connect_changed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsScene) OnSceneRectChanged(fn func(*QRectF)) int64 {
	id := registerCallback(func(rect uintptr) { fn(UnsafeNewQRectF(unsafe.Pointer(rect))) })
	pg_QGraphicsScene_connect_sceneRectChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsScene) OnSelectionChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QGraphicsScene_connect_selectionChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsScene) OnFocusItemChanged(fn func(*QGraphicsItem, *QGraphicsItem, FocusReason)) int64 {
	id := registerCallback(func(newFocus uintptr, oldFocus uintptr, reason int32) { fn(UnsafeNewQGraphicsItem(unsafe.Pointer(newFocus)), UnsafeNewQGraphicsItem(unsafe.Pointer(oldFocus)), FocusReason(reason)) })
	pg_QGraphicsScene_connect_focusItemChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsScene) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QGraphicsScene_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsScene) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QGraphicsScene_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsScene) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QGraphicsScene_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsScene) Disconnect(handle int64) {
	pg_QGraphicsScene__disconnect(handle)
	unregisterCallback(handle)
}

func QGraphicsScene_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case 7345853097340709638:
		return 1
	case 1237236456312824991:
		return 2
	case 5217623514630858297:
		return 3
	case -6698043222227006015:
		return 4
	case 206122715919917589:
		return 5
	case -1781545137352745978:
		return 6
	case -3139904158373725006:
		return 7
	case 9107058682754591716:
		return 8
	case 975043767033552552:
		return 9
	case -385596010576048993:
		return 10
	case 1696901855134526294:
		return 11
	case 1023791977679793185:
		return 12
	case 772598475950050299:
		return 13
	case -332090898078543319:
		return 14
	case -4255494004099806581:
		return 15
	case -4415920049297795213:
		return 16
	case -4029073905697023887:
		return 17
	case 1919988340934511316:
		return 18
	case 5438819661942843222:
		return 19
	case -3076319991076358961:
		return 20
	case 8267095530709235866:
		return 21
	case 5898328517505784191:
		return 22
	case -5409927581812374498:
		return 23
	case -1185315367526857833:
		return 24
	case 3769868229123868474:
		return 25
	case 7379029982170503008:
		return 26
	case -568369772982215962:
		return 27
	}
	return -1
}

func QGraphicsScene_enable(h uintptr, idx int32) { pg_QGraphicsScene__set_override(h, idx) }

func (this *QGraphicsScene) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QGraphicsScene) OnInputMethodQuery(fn func(InputMethodQuery) *QVariant) {
	registerVirt(this.UnsafeGoID(), 7345853097340709638, func(query int32) uintptr { return uintptr(fn(InputMethodQuery(query)).UnsafePointer()) }) // slot "inputMethodQuery"
}

func (this *QGraphicsScene) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QGraphicsScene) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QGraphicsScene) OnContextMenuEvent(fn func(*QGraphicsSceneContextMenuEvent)) {
	registerVirt(this.UnsafeGoID(), -6698043222227006015, func(event uintptr) { fn(UnsafeNewQGraphicsSceneContextMenuEvent(unsafe.Pointer(event))) }) // slot "contextMenuEvent"
}

func (this *QGraphicsScene) OnDragEnterEvent(fn func(*QGraphicsSceneDragDropEvent)) {
	registerVirt(this.UnsafeGoID(), 206122715919917589, func(event uintptr) { fn(UnsafeNewQGraphicsSceneDragDropEvent(unsafe.Pointer(event))) }) // slot "dragEnterEvent"
}

func (this *QGraphicsScene) OnDragMoveEvent(fn func(*QGraphicsSceneDragDropEvent)) {
	registerVirt(this.UnsafeGoID(), -1781545137352745978, func(event uintptr) { fn(UnsafeNewQGraphicsSceneDragDropEvent(unsafe.Pointer(event))) }) // slot "dragMoveEvent"
}

func (this *QGraphicsScene) OnDragLeaveEvent(fn func(*QGraphicsSceneDragDropEvent)) {
	registerVirt(this.UnsafeGoID(), -3139904158373725006, func(event uintptr) { fn(UnsafeNewQGraphicsSceneDragDropEvent(unsafe.Pointer(event))) }) // slot "dragLeaveEvent"
}

func (this *QGraphicsScene) OnDropEvent(fn func(*QGraphicsSceneDragDropEvent)) {
	registerVirt(this.UnsafeGoID(), 9107058682754591716, func(event uintptr) { fn(UnsafeNewQGraphicsSceneDragDropEvent(unsafe.Pointer(event))) }) // slot "dropEvent"
}

func (this *QGraphicsScene) OnFocusInEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), 975043767033552552, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusInEvent"
}

func (this *QGraphicsScene) OnFocusOutEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), -385596010576048993, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusOutEvent"
}

func (this *QGraphicsScene) OnHelpEvent(fn func(*QGraphicsSceneHelpEvent)) {
	registerVirt(this.UnsafeGoID(), 1696901855134526294, func(event uintptr) { fn(UnsafeNewQGraphicsSceneHelpEvent(unsafe.Pointer(event))) }) // slot "helpEvent"
}

func (this *QGraphicsScene) OnKeyPressEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 1023791977679793185, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyPressEvent"
}

func (this *QGraphicsScene) OnKeyReleaseEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 772598475950050299, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyReleaseEvent"
}

func (this *QGraphicsScene) OnMousePressEvent(fn func(*QGraphicsSceneMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -332090898078543319, func(event uintptr) { fn(UnsafeNewQGraphicsSceneMouseEvent(unsafe.Pointer(event))) }) // slot "mousePressEvent"
}

func (this *QGraphicsScene) OnMouseMoveEvent(fn func(*QGraphicsSceneMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4255494004099806581, func(event uintptr) { fn(UnsafeNewQGraphicsSceneMouseEvent(unsafe.Pointer(event))) }) // slot "mouseMoveEvent"
}

func (this *QGraphicsScene) OnMouseReleaseEvent(fn func(*QGraphicsSceneMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4415920049297795213, func(event uintptr) { fn(UnsafeNewQGraphicsSceneMouseEvent(unsafe.Pointer(event))) }) // slot "mouseReleaseEvent"
}

func (this *QGraphicsScene) OnMouseDoubleClickEvent(fn func(*QGraphicsSceneMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4029073905697023887, func(event uintptr) { fn(UnsafeNewQGraphicsSceneMouseEvent(unsafe.Pointer(event))) }) // slot "mouseDoubleClickEvent"
}

func (this *QGraphicsScene) OnWheelEvent(fn func(*QGraphicsSceneWheelEvent)) {
	registerVirt(this.UnsafeGoID(), 1919988340934511316, func(event uintptr) { fn(UnsafeNewQGraphicsSceneWheelEvent(unsafe.Pointer(event))) }) // slot "wheelEvent"
}

func (this *QGraphicsScene) OnInputMethodEvent(fn func(*QInputMethodEvent)) {
	registerVirt(this.UnsafeGoID(), 5438819661942843222, func(event uintptr) { fn(UnsafeNewQInputMethodEvent(unsafe.Pointer(event))) }) // slot "inputMethodEvent"
}

func (this *QGraphicsScene) OnDrawBackground(fn func(*QPainter, *QRectF)) {
	registerVirt(this.UnsafeGoID(), -3076319991076358961, func(painter uintptr, rect uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter)), UnsafeNewQRectF(unsafe.Pointer(rect))) }) // slot "drawBackground"
}

func (this *QGraphicsScene) OnDrawForeground(fn func(*QPainter, *QRectF)) {
	registerVirt(this.UnsafeGoID(), 8267095530709235866, func(painter uintptr, rect uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter)), UnsafeNewQRectF(unsafe.Pointer(rect))) }) // slot "drawForeground"
}

func (this *QGraphicsScene) OnFocusNextPrevChild(fn func(bool) bool) {
	registerVirt(this.UnsafeGoID(), 5898328517505784191, func(next int32) int32 { return b2i(fn(next != 0)) }) // slot "focusNextPrevChild"
}

func (this *QGraphicsScene) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QGraphicsScene) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QGraphicsScene) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QGraphicsScene) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QGraphicsScene) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QGraphicsScene) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QGraphicsScene_delete(this.h)
	}
	this.UnsafeRelease()
}
