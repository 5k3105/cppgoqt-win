package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QGraphicsView_chain uintptr
	addr_QGraphicsView_new uintptr
	addr_QGraphicsView_new2 uintptr
	addr_QGraphicsView_new3 uintptr
	addr_QGraphicsView_new4 uintptr
	addr_QGraphicsView_tr uintptr
	addr_QGraphicsView_renderHints uintptr
	addr_QGraphicsView_setRenderHint uintptr
	addr_QGraphicsView_setRenderHints uintptr
	addr_QGraphicsView_alignment uintptr
	addr_QGraphicsView_setAlignment uintptr
	addr_QGraphicsView_transformationAnchor uintptr
	addr_QGraphicsView_setTransformationAnchor uintptr
	addr_QGraphicsView_resizeAnchor uintptr
	addr_QGraphicsView_setResizeAnchor uintptr
	addr_QGraphicsView_viewportUpdateMode uintptr
	addr_QGraphicsView_setViewportUpdateMode uintptr
	addr_QGraphicsView_optimizationFlags uintptr
	addr_QGraphicsView_setOptimizationFlag uintptr
	addr_QGraphicsView_setOptimizationFlags uintptr
	addr_QGraphicsView_dragMode uintptr
	addr_QGraphicsView_setDragMode uintptr
	addr_QGraphicsView_rubberBandSelectionMode uintptr
	addr_QGraphicsView_setRubberBandSelectionMode uintptr
	addr_QGraphicsView_rubberBandRect uintptr
	addr_QGraphicsView_cacheMode uintptr
	addr_QGraphicsView_setCacheMode uintptr
	addr_QGraphicsView_resetCachedContent uintptr
	addr_QGraphicsView_isInteractive uintptr
	addr_QGraphicsView_setInteractive uintptr
	addr_QGraphicsView_scene uintptr
	addr_QGraphicsView_setScene uintptr
	addr_QGraphicsView_sceneRect uintptr
	addr_QGraphicsView_setSceneRect uintptr
	pg_QGraphicsView_setSceneRect2 func(uintptr, float64, float64, float64, float64)
	addr_QGraphicsView_transform uintptr
	addr_QGraphicsView_viewportTransform uintptr
	addr_QGraphicsView_isTransformed uintptr
	addr_QGraphicsView_setTransform uintptr
	addr_QGraphicsView_resetTransform uintptr
	pg_QGraphicsView_rotate func(uintptr, float64)
	pg_QGraphicsView_scale func(uintptr, float64, float64)
	pg_QGraphicsView_shear func(uintptr, float64, float64)
	pg_QGraphicsView_translate func(uintptr, float64, float64)
	addr_QGraphicsView_centerOn uintptr
	pg_QGraphicsView_centerOn2 func(uintptr, float64, float64)
	addr_QGraphicsView_centerOn3 uintptr
	addr_QGraphicsView_ensureVisible uintptr
	pg_QGraphicsView_ensureVisible2 func(uintptr, float64, float64, float64, float64)
	addr_QGraphicsView_ensureVisible3 uintptr
	addr_QGraphicsView_fitInView uintptr
	pg_QGraphicsView_fitInView2 func(uintptr, float64, float64, float64, float64)
	addr_QGraphicsView_fitInView3 uintptr
	addr_QGraphicsView_render uintptr
	addr_QGraphicsView_items uintptr
	addr_QGraphicsView_items2 uintptr
	addr_QGraphicsView_items3 uintptr
	addr_QGraphicsView_items4 uintptr
	addr_QGraphicsView_items5 uintptr
	addr_QGraphicsView_items6 uintptr
	addr_QGraphicsView_items7 uintptr
	addr_QGraphicsView_itemAt uintptr
	addr_QGraphicsView_itemAt2 uintptr
	addr_QGraphicsView_mapToScene uintptr
	addr_QGraphicsView_mapToScene2 uintptr
	addr_QGraphicsView_mapToScene3 uintptr
	addr_QGraphicsView_mapToScene4 uintptr
	addr_QGraphicsView_mapFromScene uintptr
	addr_QGraphicsView_mapFromScene2 uintptr
	addr_QGraphicsView_mapFromScene3 uintptr
	addr_QGraphicsView_mapFromScene4 uintptr
	addr_QGraphicsView_mapToScene5 uintptr
	addr_QGraphicsView_mapToScene6 uintptr
	pg_QGraphicsView_mapFromScene5 func(uintptr, float64, float64) uintptr
	pg_QGraphicsView_mapFromScene6 func(uintptr, float64, float64, float64, float64) uintptr
	addr_QGraphicsView_backgroundBrush uintptr
	addr_QGraphicsView_setBackgroundBrush uintptr
	addr_QGraphicsView_foregroundBrush uintptr
	addr_QGraphicsView_setForegroundBrush uintptr
	addr_QGraphicsView_updateScene uintptr
	addr_QGraphicsView_invalidateScene uintptr
	addr_QGraphicsView_updateSceneRect uintptr
	addr_QGraphicsView_rubberBandChanged uintptr
	addr_QGraphicsView_setViewportMargins uintptr
	addr_QGraphicsView_setViewportMargins2 uintptr
	addr_QGraphicsView_viewportMargins uintptr
	addr_QGraphicsView_drawFrame uintptr
	addr_QGraphicsView_setTabOrder uintptr
	addr_QGraphicsView_mouseGrabber uintptr
	addr_QGraphicsView_keyboardGrabber uintptr
	addr_QGraphicsView_find uintptr
	addr_QGraphicsView_createWindowContainer uintptr
	addr_QGraphicsView_updateMicroFocus uintptr
	addr_QGraphicsView_create uintptr
	addr_QGraphicsView_destroy uintptr
	addr_QGraphicsView_focusNextChild uintptr
	addr_QGraphicsView_focusPreviousChild uintptr
	addr_QGraphicsView_connect uintptr
	addr_QGraphicsView_connect2 uintptr
	addr_QGraphicsView_sender uintptr
	addr_QGraphicsView_senderSignalIndex uintptr
	addr_QGraphicsView_receivers uintptr
	addr_QGraphicsView_isSignalConnected uintptr
	pg_QGraphicsView_devicePixelRatioFScale func() float64
	pg_QGraphicsView_encodeMetricF func(int32, float64) int32
	pg_QGraphicsView_getDecodedMetricF func(uintptr, int32, int32) float64
	addr_QGraphicsView_tr2 uintptr
	addr_QGraphicsView_tr3 uintptr
	addr_QGraphicsView_setRenderHint2 uintptr
	addr_QGraphicsView_setOptimizationFlag2 uintptr
	addr_QGraphicsView_setTransform2 uintptr
	addr_QGraphicsView_ensureVisible4 uintptr
	addr_QGraphicsView_ensureVisible5 uintptr
	pg_QGraphicsView_ensureVisible6 func(uintptr, float64, float64, float64, float64, int32)
	pg_QGraphicsView_ensureVisible7 func(uintptr, float64, float64, float64, float64, int32, int32)
	addr_QGraphicsView_ensureVisible8 uintptr
	addr_QGraphicsView_ensureVisible9 uintptr
	addr_QGraphicsView_fitInView4 uintptr
	pg_QGraphicsView_fitInView5 func(uintptr, float64, float64, float64, float64, int32)
	addr_QGraphicsView_fitInView6 uintptr
	addr_QGraphicsView_render2 uintptr
	addr_QGraphicsView_render3 uintptr
	addr_QGraphicsView_render4 uintptr
	addr_QGraphicsView_items8 uintptr
	addr_QGraphicsView_items9 uintptr
	addr_QGraphicsView_items10 uintptr
	addr_QGraphicsView_items11 uintptr
	addr_QGraphicsView_invalidateScene2 uintptr
	addr_QGraphicsView_invalidateScene3 uintptr
	addr_QGraphicsView_createWindowContainer2 uintptr
	addr_QGraphicsView_createWindowContainer3 uintptr
	addr_QGraphicsView_updateMicroFocus2 uintptr
	addr_QGraphicsView_create2 uintptr
	addr_QGraphicsView_create3 uintptr
	addr_QGraphicsView_create4 uintptr
	addr_QGraphicsView_destroy2 uintptr
	addr_QGraphicsView_destroy3 uintptr
	addr_QGraphicsView_connect3 uintptr
	addr_QGraphicsView_connect4 uintptr
	addr_QGraphicsView_metaObject_default uintptr
	addr_QGraphicsView_sizeHint_default uintptr
	addr_QGraphicsView_inputMethodQuery_default uintptr
	addr_QGraphicsView_setupViewport_default uintptr
	addr_QGraphicsView_event_default uintptr
	addr_QGraphicsView_viewportEvent_default uintptr
	addr_QGraphicsView_contextMenuEvent_default uintptr
	addr_QGraphicsView_dragEnterEvent_default uintptr
	addr_QGraphicsView_dragLeaveEvent_default uintptr
	addr_QGraphicsView_dragMoveEvent_default uintptr
	addr_QGraphicsView_dropEvent_default uintptr
	addr_QGraphicsView_focusInEvent_default uintptr
	addr_QGraphicsView_focusNextPrevChild_default uintptr
	addr_QGraphicsView_focusOutEvent_default uintptr
	addr_QGraphicsView_keyPressEvent_default uintptr
	addr_QGraphicsView_keyReleaseEvent_default uintptr
	addr_QGraphicsView_mouseDoubleClickEvent_default uintptr
	addr_QGraphicsView_mousePressEvent_default uintptr
	addr_QGraphicsView_mouseMoveEvent_default uintptr
	addr_QGraphicsView_mouseReleaseEvent_default uintptr
	addr_QGraphicsView_wheelEvent_default uintptr
	addr_QGraphicsView_paintEvent_default uintptr
	addr_QGraphicsView_resizeEvent_default uintptr
	addr_QGraphicsView_scrollContentsBy_default uintptr
	addr_QGraphicsView_showEvent_default uintptr
	addr_QGraphicsView_inputMethodEvent_default uintptr
	addr_QGraphicsView_drawBackground_default uintptr
	addr_QGraphicsView_drawForeground_default uintptr
	addr_QGraphicsView_minimumSizeHint_default uintptr
	addr_QGraphicsView_eventFilter_default uintptr
	addr_QGraphicsView_viewportSizeHint_default uintptr
	addr_QGraphicsView_changeEvent_default uintptr
	addr_QGraphicsView_initStyleOption_default uintptr
	addr_QGraphicsView_devType_default uintptr
	addr_QGraphicsView_setVisible_default uintptr
	addr_QGraphicsView_heightForWidth_default uintptr
	addr_QGraphicsView_hasHeightForWidth_default uintptr
	addr_QGraphicsView_paintEngine_default uintptr
	addr_QGraphicsView_enterEvent_default uintptr
	addr_QGraphicsView_leaveEvent_default uintptr
	addr_QGraphicsView_moveEvent_default uintptr
	addr_QGraphicsView_closeEvent_default uintptr
	addr_QGraphicsView_tabletEvent_default uintptr
	addr_QGraphicsView_actionEvent_default uintptr
	addr_QGraphicsView_hideEvent_default uintptr
	addr_QGraphicsView_metric_default uintptr
	addr_QGraphicsView_initPainter_default uintptr
	addr_QGraphicsView_redirected_default uintptr
	addr_QGraphicsView_sharedPainter_default uintptr
	addr_QGraphicsView_timerEvent_default uintptr
	addr_QGraphicsView_childEvent_default uintptr
	addr_QGraphicsView_customEvent_default uintptr
	addr_QGraphicsView_connectNotify_default uintptr
	addr_QGraphicsView_disconnectNotify_default uintptr
	pg_QGraphicsView_connect_rubberBandChanged func(uintptr, int64)
	pg_QGraphicsView_connect_windowTitleChanged func(uintptr, int64)
	pg_QGraphicsView_connect_windowIconChanged func(uintptr, int64)
	pg_QGraphicsView_connect_windowIconTextChanged func(uintptr, int64)
	pg_QGraphicsView_connect_customContextMenuRequested func(uintptr, int64)
	pg_QGraphicsView_connect_destroyed func(uintptr, int64)
	pg_QGraphicsView_connect_destroyed2 func(uintptr, int64)
	pg_QGraphicsView_connect_objectNameChanged func(uintptr, int64)
	pg_QGraphicsView__disconnect func(int64)
	pg_QGraphicsView__set_override func(uintptr, int32)
	pg_QGraphicsView_delete func(uintptr)
)

func init() {
	symaddr(&addr_QGraphicsView_chain, "QGraphicsView_chain")
	symaddr(&addr_QGraphicsView_new, "QGraphicsView_new")
	symaddr(&addr_QGraphicsView_new2, "QGraphicsView_new2")
	symaddr(&addr_QGraphicsView_new3, "QGraphicsView_new3")
	symaddr(&addr_QGraphicsView_new4, "QGraphicsView_new4")
	symaddr(&addr_QGraphicsView_tr, "QGraphicsView_tr")
	symaddr(&addr_QGraphicsView_renderHints, "QGraphicsView_renderHints")
	symaddr(&addr_QGraphicsView_setRenderHint, "QGraphicsView_setRenderHint")
	symaddr(&addr_QGraphicsView_setRenderHints, "QGraphicsView_setRenderHints")
	symaddr(&addr_QGraphicsView_alignment, "QGraphicsView_alignment")
	symaddr(&addr_QGraphicsView_setAlignment, "QGraphicsView_setAlignment")
	symaddr(&addr_QGraphicsView_transformationAnchor, "QGraphicsView_transformationAnchor")
	symaddr(&addr_QGraphicsView_setTransformationAnchor, "QGraphicsView_setTransformationAnchor")
	symaddr(&addr_QGraphicsView_resizeAnchor, "QGraphicsView_resizeAnchor")
	symaddr(&addr_QGraphicsView_setResizeAnchor, "QGraphicsView_setResizeAnchor")
	symaddr(&addr_QGraphicsView_viewportUpdateMode, "QGraphicsView_viewportUpdateMode")
	symaddr(&addr_QGraphicsView_setViewportUpdateMode, "QGraphicsView_setViewportUpdateMode")
	symaddr(&addr_QGraphicsView_optimizationFlags, "QGraphicsView_optimizationFlags")
	symaddr(&addr_QGraphicsView_setOptimizationFlag, "QGraphicsView_setOptimizationFlag")
	symaddr(&addr_QGraphicsView_setOptimizationFlags, "QGraphicsView_setOptimizationFlags")
	symaddr(&addr_QGraphicsView_dragMode, "QGraphicsView_dragMode")
	symaddr(&addr_QGraphicsView_setDragMode, "QGraphicsView_setDragMode")
	symaddr(&addr_QGraphicsView_rubberBandSelectionMode, "QGraphicsView_rubberBandSelectionMode")
	symaddr(&addr_QGraphicsView_setRubberBandSelectionMode, "QGraphicsView_setRubberBandSelectionMode")
	symaddr(&addr_QGraphicsView_rubberBandRect, "QGraphicsView_rubberBandRect")
	symaddr(&addr_QGraphicsView_cacheMode, "QGraphicsView_cacheMode")
	symaddr(&addr_QGraphicsView_setCacheMode, "QGraphicsView_setCacheMode")
	symaddr(&addr_QGraphicsView_resetCachedContent, "QGraphicsView_resetCachedContent")
	symaddr(&addr_QGraphicsView_isInteractive, "QGraphicsView_isInteractive")
	symaddr(&addr_QGraphicsView_setInteractive, "QGraphicsView_setInteractive")
	symaddr(&addr_QGraphicsView_scene, "QGraphicsView_scene")
	symaddr(&addr_QGraphicsView_setScene, "QGraphicsView_setScene")
	symaddr(&addr_QGraphicsView_sceneRect, "QGraphicsView_sceneRect")
	symaddr(&addr_QGraphicsView_setSceneRect, "QGraphicsView_setSceneRect")
	reg(&pg_QGraphicsView_setSceneRect2, "QGraphicsView_setSceneRect2")
	symaddr(&addr_QGraphicsView_transform, "QGraphicsView_transform")
	symaddr(&addr_QGraphicsView_viewportTransform, "QGraphicsView_viewportTransform")
	symaddr(&addr_QGraphicsView_isTransformed, "QGraphicsView_isTransformed")
	symaddr(&addr_QGraphicsView_setTransform, "QGraphicsView_setTransform")
	symaddr(&addr_QGraphicsView_resetTransform, "QGraphicsView_resetTransform")
	reg(&pg_QGraphicsView_rotate, "QGraphicsView_rotate")
	reg(&pg_QGraphicsView_scale, "QGraphicsView_scale")
	reg(&pg_QGraphicsView_shear, "QGraphicsView_shear")
	reg(&pg_QGraphicsView_translate, "QGraphicsView_translate")
	symaddr(&addr_QGraphicsView_centerOn, "QGraphicsView_centerOn")
	reg(&pg_QGraphicsView_centerOn2, "QGraphicsView_centerOn2")
	symaddr(&addr_QGraphicsView_centerOn3, "QGraphicsView_centerOn3")
	symaddr(&addr_QGraphicsView_ensureVisible, "QGraphicsView_ensureVisible")
	reg(&pg_QGraphicsView_ensureVisible2, "QGraphicsView_ensureVisible2")
	symaddr(&addr_QGraphicsView_ensureVisible3, "QGraphicsView_ensureVisible3")
	symaddr(&addr_QGraphicsView_fitInView, "QGraphicsView_fitInView")
	reg(&pg_QGraphicsView_fitInView2, "QGraphicsView_fitInView2")
	symaddr(&addr_QGraphicsView_fitInView3, "QGraphicsView_fitInView3")
	symaddr(&addr_QGraphicsView_render, "QGraphicsView_render")
	symaddr(&addr_QGraphicsView_items, "QGraphicsView_items")
	symaddr(&addr_QGraphicsView_items2, "QGraphicsView_items2")
	symaddr(&addr_QGraphicsView_items3, "QGraphicsView_items3")
	symaddr(&addr_QGraphicsView_items4, "QGraphicsView_items4")
	symaddr(&addr_QGraphicsView_items5, "QGraphicsView_items5")
	symaddr(&addr_QGraphicsView_items6, "QGraphicsView_items6")
	symaddr(&addr_QGraphicsView_items7, "QGraphicsView_items7")
	symaddr(&addr_QGraphicsView_itemAt, "QGraphicsView_itemAt")
	symaddr(&addr_QGraphicsView_itemAt2, "QGraphicsView_itemAt2")
	symaddr(&addr_QGraphicsView_mapToScene, "QGraphicsView_mapToScene")
	symaddr(&addr_QGraphicsView_mapToScene2, "QGraphicsView_mapToScene2")
	symaddr(&addr_QGraphicsView_mapToScene3, "QGraphicsView_mapToScene3")
	symaddr(&addr_QGraphicsView_mapToScene4, "QGraphicsView_mapToScene4")
	symaddr(&addr_QGraphicsView_mapFromScene, "QGraphicsView_mapFromScene")
	symaddr(&addr_QGraphicsView_mapFromScene2, "QGraphicsView_mapFromScene2")
	symaddr(&addr_QGraphicsView_mapFromScene3, "QGraphicsView_mapFromScene3")
	symaddr(&addr_QGraphicsView_mapFromScene4, "QGraphicsView_mapFromScene4")
	symaddr(&addr_QGraphicsView_mapToScene5, "QGraphicsView_mapToScene5")
	symaddr(&addr_QGraphicsView_mapToScene6, "QGraphicsView_mapToScene6")
	reg(&pg_QGraphicsView_mapFromScene5, "QGraphicsView_mapFromScene5")
	reg(&pg_QGraphicsView_mapFromScene6, "QGraphicsView_mapFromScene6")
	symaddr(&addr_QGraphicsView_backgroundBrush, "QGraphicsView_backgroundBrush")
	symaddr(&addr_QGraphicsView_setBackgroundBrush, "QGraphicsView_setBackgroundBrush")
	symaddr(&addr_QGraphicsView_foregroundBrush, "QGraphicsView_foregroundBrush")
	symaddr(&addr_QGraphicsView_setForegroundBrush, "QGraphicsView_setForegroundBrush")
	symaddr(&addr_QGraphicsView_updateScene, "QGraphicsView_updateScene")
	symaddr(&addr_QGraphicsView_invalidateScene, "QGraphicsView_invalidateScene")
	symaddr(&addr_QGraphicsView_updateSceneRect, "QGraphicsView_updateSceneRect")
	symaddr(&addr_QGraphicsView_rubberBandChanged, "QGraphicsView_rubberBandChanged")
	symaddr(&addr_QGraphicsView_setViewportMargins, "QGraphicsView_setViewportMargins")
	symaddr(&addr_QGraphicsView_setViewportMargins2, "QGraphicsView_setViewportMargins2")
	symaddr(&addr_QGraphicsView_viewportMargins, "QGraphicsView_viewportMargins")
	symaddr(&addr_QGraphicsView_drawFrame, "QGraphicsView_drawFrame")
	symaddr(&addr_QGraphicsView_setTabOrder, "QGraphicsView_setTabOrder")
	symaddr(&addr_QGraphicsView_mouseGrabber, "QGraphicsView_mouseGrabber")
	symaddr(&addr_QGraphicsView_keyboardGrabber, "QGraphicsView_keyboardGrabber")
	symaddr(&addr_QGraphicsView_find, "QGraphicsView_find")
	symaddr(&addr_QGraphicsView_createWindowContainer, "QGraphicsView_createWindowContainer")
	symaddr(&addr_QGraphicsView_updateMicroFocus, "QGraphicsView_updateMicroFocus")
	symaddr(&addr_QGraphicsView_create, "QGraphicsView_create")
	symaddr(&addr_QGraphicsView_destroy, "QGraphicsView_destroy")
	symaddr(&addr_QGraphicsView_focusNextChild, "QGraphicsView_focusNextChild")
	symaddr(&addr_QGraphicsView_focusPreviousChild, "QGraphicsView_focusPreviousChild")
	symaddr(&addr_QGraphicsView_connect, "QGraphicsView_connect")
	symaddr(&addr_QGraphicsView_connect2, "QGraphicsView_connect2")
	symaddr(&addr_QGraphicsView_sender, "QGraphicsView_sender")
	symaddr(&addr_QGraphicsView_senderSignalIndex, "QGraphicsView_senderSignalIndex")
	symaddr(&addr_QGraphicsView_receivers, "QGraphicsView_receivers")
	symaddr(&addr_QGraphicsView_isSignalConnected, "QGraphicsView_isSignalConnected")
	reg(&pg_QGraphicsView_devicePixelRatioFScale, "QGraphicsView_devicePixelRatioFScale")
	reg(&pg_QGraphicsView_encodeMetricF, "QGraphicsView_encodeMetricF")
	reg(&pg_QGraphicsView_getDecodedMetricF, "QGraphicsView_getDecodedMetricF")
	symaddr(&addr_QGraphicsView_tr2, "QGraphicsView_tr2")
	symaddr(&addr_QGraphicsView_tr3, "QGraphicsView_tr3")
	symaddr(&addr_QGraphicsView_setRenderHint2, "QGraphicsView_setRenderHint2")
	symaddr(&addr_QGraphicsView_setOptimizationFlag2, "QGraphicsView_setOptimizationFlag2")
	symaddr(&addr_QGraphicsView_setTransform2, "QGraphicsView_setTransform2")
	symaddr(&addr_QGraphicsView_ensureVisible4, "QGraphicsView_ensureVisible4")
	symaddr(&addr_QGraphicsView_ensureVisible5, "QGraphicsView_ensureVisible5")
	reg(&pg_QGraphicsView_ensureVisible6, "QGraphicsView_ensureVisible6")
	reg(&pg_QGraphicsView_ensureVisible7, "QGraphicsView_ensureVisible7")
	symaddr(&addr_QGraphicsView_ensureVisible8, "QGraphicsView_ensureVisible8")
	symaddr(&addr_QGraphicsView_ensureVisible9, "QGraphicsView_ensureVisible9")
	symaddr(&addr_QGraphicsView_fitInView4, "QGraphicsView_fitInView4")
	reg(&pg_QGraphicsView_fitInView5, "QGraphicsView_fitInView5")
	symaddr(&addr_QGraphicsView_fitInView6, "QGraphicsView_fitInView6")
	symaddr(&addr_QGraphicsView_render2, "QGraphicsView_render2")
	symaddr(&addr_QGraphicsView_render3, "QGraphicsView_render3")
	symaddr(&addr_QGraphicsView_render4, "QGraphicsView_render4")
	symaddr(&addr_QGraphicsView_items8, "QGraphicsView_items8")
	symaddr(&addr_QGraphicsView_items9, "QGraphicsView_items9")
	symaddr(&addr_QGraphicsView_items10, "QGraphicsView_items10")
	symaddr(&addr_QGraphicsView_items11, "QGraphicsView_items11")
	symaddr(&addr_QGraphicsView_invalidateScene2, "QGraphicsView_invalidateScene2")
	symaddr(&addr_QGraphicsView_invalidateScene3, "QGraphicsView_invalidateScene3")
	symaddr(&addr_QGraphicsView_createWindowContainer2, "QGraphicsView_createWindowContainer2")
	symaddr(&addr_QGraphicsView_createWindowContainer3, "QGraphicsView_createWindowContainer3")
	symaddr(&addr_QGraphicsView_updateMicroFocus2, "QGraphicsView_updateMicroFocus2")
	symaddr(&addr_QGraphicsView_create2, "QGraphicsView_create2")
	symaddr(&addr_QGraphicsView_create3, "QGraphicsView_create3")
	symaddr(&addr_QGraphicsView_create4, "QGraphicsView_create4")
	symaddr(&addr_QGraphicsView_destroy2, "QGraphicsView_destroy2")
	symaddr(&addr_QGraphicsView_destroy3, "QGraphicsView_destroy3")
	symaddr(&addr_QGraphicsView_connect3, "QGraphicsView_connect3")
	symaddr(&addr_QGraphicsView_connect4, "QGraphicsView_connect4")
	symaddr(&addr_QGraphicsView_metaObject_default, "QGraphicsView_metaObject_default")
	symaddr(&addr_QGraphicsView_sizeHint_default, "QGraphicsView_sizeHint_default")
	symaddr(&addr_QGraphicsView_inputMethodQuery_default, "QGraphicsView_inputMethodQuery_default")
	symaddr(&addr_QGraphicsView_setupViewport_default, "QGraphicsView_setupViewport_default")
	symaddr(&addr_QGraphicsView_event_default, "QGraphicsView_event_default")
	symaddr(&addr_QGraphicsView_viewportEvent_default, "QGraphicsView_viewportEvent_default")
	symaddr(&addr_QGraphicsView_contextMenuEvent_default, "QGraphicsView_contextMenuEvent_default")
	symaddr(&addr_QGraphicsView_dragEnterEvent_default, "QGraphicsView_dragEnterEvent_default")
	symaddr(&addr_QGraphicsView_dragLeaveEvent_default, "QGraphicsView_dragLeaveEvent_default")
	symaddr(&addr_QGraphicsView_dragMoveEvent_default, "QGraphicsView_dragMoveEvent_default")
	symaddr(&addr_QGraphicsView_dropEvent_default, "QGraphicsView_dropEvent_default")
	symaddr(&addr_QGraphicsView_focusInEvent_default, "QGraphicsView_focusInEvent_default")
	symaddr(&addr_QGraphicsView_focusNextPrevChild_default, "QGraphicsView_focusNextPrevChild_default")
	symaddr(&addr_QGraphicsView_focusOutEvent_default, "QGraphicsView_focusOutEvent_default")
	symaddr(&addr_QGraphicsView_keyPressEvent_default, "QGraphicsView_keyPressEvent_default")
	symaddr(&addr_QGraphicsView_keyReleaseEvent_default, "QGraphicsView_keyReleaseEvent_default")
	symaddr(&addr_QGraphicsView_mouseDoubleClickEvent_default, "QGraphicsView_mouseDoubleClickEvent_default")
	symaddr(&addr_QGraphicsView_mousePressEvent_default, "QGraphicsView_mousePressEvent_default")
	symaddr(&addr_QGraphicsView_mouseMoveEvent_default, "QGraphicsView_mouseMoveEvent_default")
	symaddr(&addr_QGraphicsView_mouseReleaseEvent_default, "QGraphicsView_mouseReleaseEvent_default")
	symaddr(&addr_QGraphicsView_wheelEvent_default, "QGraphicsView_wheelEvent_default")
	symaddr(&addr_QGraphicsView_paintEvent_default, "QGraphicsView_paintEvent_default")
	symaddr(&addr_QGraphicsView_resizeEvent_default, "QGraphicsView_resizeEvent_default")
	symaddr(&addr_QGraphicsView_scrollContentsBy_default, "QGraphicsView_scrollContentsBy_default")
	symaddr(&addr_QGraphicsView_showEvent_default, "QGraphicsView_showEvent_default")
	symaddr(&addr_QGraphicsView_inputMethodEvent_default, "QGraphicsView_inputMethodEvent_default")
	symaddr(&addr_QGraphicsView_drawBackground_default, "QGraphicsView_drawBackground_default")
	symaddr(&addr_QGraphicsView_drawForeground_default, "QGraphicsView_drawForeground_default")
	symaddr(&addr_QGraphicsView_minimumSizeHint_default, "QGraphicsView_minimumSizeHint_default")
	symaddr(&addr_QGraphicsView_eventFilter_default, "QGraphicsView_eventFilter_default")
	symaddr(&addr_QGraphicsView_viewportSizeHint_default, "QGraphicsView_viewportSizeHint_default")
	symaddr(&addr_QGraphicsView_changeEvent_default, "QGraphicsView_changeEvent_default")
	symaddr(&addr_QGraphicsView_initStyleOption_default, "QGraphicsView_initStyleOption_default")
	symaddr(&addr_QGraphicsView_devType_default, "QGraphicsView_devType_default")
	symaddr(&addr_QGraphicsView_setVisible_default, "QGraphicsView_setVisible_default")
	symaddr(&addr_QGraphicsView_heightForWidth_default, "QGraphicsView_heightForWidth_default")
	symaddr(&addr_QGraphicsView_hasHeightForWidth_default, "QGraphicsView_hasHeightForWidth_default")
	symaddr(&addr_QGraphicsView_paintEngine_default, "QGraphicsView_paintEngine_default")
	symaddr(&addr_QGraphicsView_enterEvent_default, "QGraphicsView_enterEvent_default")
	symaddr(&addr_QGraphicsView_leaveEvent_default, "QGraphicsView_leaveEvent_default")
	symaddr(&addr_QGraphicsView_moveEvent_default, "QGraphicsView_moveEvent_default")
	symaddr(&addr_QGraphicsView_closeEvent_default, "QGraphicsView_closeEvent_default")
	symaddr(&addr_QGraphicsView_tabletEvent_default, "QGraphicsView_tabletEvent_default")
	symaddr(&addr_QGraphicsView_actionEvent_default, "QGraphicsView_actionEvent_default")
	symaddr(&addr_QGraphicsView_hideEvent_default, "QGraphicsView_hideEvent_default")
	symaddr(&addr_QGraphicsView_metric_default, "QGraphicsView_metric_default")
	symaddr(&addr_QGraphicsView_initPainter_default, "QGraphicsView_initPainter_default")
	symaddr(&addr_QGraphicsView_redirected_default, "QGraphicsView_redirected_default")
	symaddr(&addr_QGraphicsView_sharedPainter_default, "QGraphicsView_sharedPainter_default")
	symaddr(&addr_QGraphicsView_timerEvent_default, "QGraphicsView_timerEvent_default")
	symaddr(&addr_QGraphicsView_childEvent_default, "QGraphicsView_childEvent_default")
	symaddr(&addr_QGraphicsView_customEvent_default, "QGraphicsView_customEvent_default")
	symaddr(&addr_QGraphicsView_connectNotify_default, "QGraphicsView_connectNotify_default")
	symaddr(&addr_QGraphicsView_disconnectNotify_default, "QGraphicsView_disconnectNotify_default")
	reg(&pg_QGraphicsView_connect_rubberBandChanged, "QGraphicsView_connect_rubberBandChanged")
	reg(&pg_QGraphicsView_connect_windowTitleChanged, "QGraphicsView_connect_windowTitleChanged")
	reg(&pg_QGraphicsView_connect_windowIconChanged, "QGraphicsView_connect_windowIconChanged")
	reg(&pg_QGraphicsView_connect_windowIconTextChanged, "QGraphicsView_connect_windowIconTextChanged")
	reg(&pg_QGraphicsView_connect_customContextMenuRequested, "QGraphicsView_connect_customContextMenuRequested")
	reg(&pg_QGraphicsView_connect_destroyed, "QGraphicsView_connect_destroyed")
	reg(&pg_QGraphicsView_connect_destroyed2, "QGraphicsView_connect_destroyed2")
	reg(&pg_QGraphicsView_connect_objectNameChanged, "QGraphicsView_connect_objectNameChanged")
	reg(&pg_QGraphicsView__disconnect, "QGraphicsView__disconnect")
	reg(&pg_QGraphicsView__set_override, "QGraphicsView__set_override")
	reg(&pg_QGraphicsView_delete, "QGraphicsView_delete")
}

type QGraphicsView struct {
	h uintptr
	*QAbstractScrollArea
}

func bindQGraphicsView(h uintptr) *QGraphicsView {
	if h == 0 {
		return nil
	}
	var __c [4]uintptr
	purego.SyscallN(addr_QGraphicsView_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QGraphicsView{h: h, QAbstractScrollArea: &QAbstractScrollArea{h: __c[0], QFrame: &QFrame{h: __c[1], QWidget: &QWidget{h: __c[2], QObject: &QObject{h: __c[3]}}}}}
}

func newownedQGraphicsView(h uintptr) *QGraphicsView { return bindQGraphicsView(h) }

func (this *QGraphicsView) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QGraphicsView) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQGraphicsView(h unsafe.Pointer) *QGraphicsView { return bindQGraphicsView(uintptr(h)) }
func UnsafeNewQGraphicsViewOwned(h unsafe.Pointer) *QGraphicsView { return newownedQGraphicsView(uintptr(h)) }

func (this *QGraphicsView) UnsafeRelease() {
	this.h = 0
	if this.QAbstractScrollArea != nil {
		this.QAbstractScrollArea.UnsafeRelease()
	}
}

func NewQGraphicsView(parent *QWidget) *QGraphicsView {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QGraphicsView_new, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQGraphicsView(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGraphicsView_virtIdx, QGraphicsView_enable)
	return __o
}

func NewQGraphicsView2() *QGraphicsView {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QGraphicsView_new2, uintptr(__goid))
	__o := newownedQGraphicsView(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGraphicsView_virtIdx, QGraphicsView_enable)
	return __o
}

func NewQGraphicsView3(scene *QGraphicsScene) *QGraphicsView {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QGraphicsView_new3, uintptr(__goid), uintptr(scene.UnsafePointer()))
	__o := newownedQGraphicsView(__r)
	runtime.KeepAlive(scene)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGraphicsView_virtIdx, QGraphicsView_enable)
	return __o
}

func NewQGraphicsView4(scene *QGraphicsScene, parent *QWidget) *QGraphicsView {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QGraphicsView_new4, uintptr(__goid), uintptr(scene.UnsafePointer()), uintptr(parent.UnsafePointer()))
	__o := newownedQGraphicsView(__r)
	runtime.KeepAlive(scene)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QGraphicsView_virtIdx, QGraphicsView_enable)
	return __o
}

func QGraphicsView_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGraphicsView_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QGraphicsView) RenderHints() QPainter__RenderHint {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsView_renderHints, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QPainter__RenderHint(__r)
}

func (this *QGraphicsView) SetRenderHint(hint QPainter__RenderHint) {
	purego.SyscallN(addr_QGraphicsView_setRenderHint, this.h, uintptr(hint))
	runtime.KeepAlive(this)
}

func (this *QGraphicsView) SetRenderHints(hints QPainter__RenderHint) {
	purego.SyscallN(addr_QGraphicsView_setRenderHints, this.h, uintptr(hints))
	runtime.KeepAlive(this)
}

func (this *QGraphicsView) Alignment() AlignmentFlag {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsView_alignment, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return AlignmentFlag(__r)
}

func (this *QGraphicsView) SetAlignment(alignment AlignmentFlag) {
	purego.SyscallN(addr_QGraphicsView_setAlignment, this.h, uintptr(alignment))
	runtime.KeepAlive(this)
}

func (this *QGraphicsView) TransformationAnchor() QGraphicsView__ViewportAnchor {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsView_transformationAnchor, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QGraphicsView__ViewportAnchor(__r)
}

func (this *QGraphicsView) SetTransformationAnchor(anchor QGraphicsView__ViewportAnchor) {
	purego.SyscallN(addr_QGraphicsView_setTransformationAnchor, this.h, uintptr(anchor))
	runtime.KeepAlive(this)
}

func (this *QGraphicsView) ResizeAnchor() QGraphicsView__ViewportAnchor {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsView_resizeAnchor, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QGraphicsView__ViewportAnchor(__r)
}

func (this *QGraphicsView) SetResizeAnchor(anchor QGraphicsView__ViewportAnchor) {
	purego.SyscallN(addr_QGraphicsView_setResizeAnchor, this.h, uintptr(anchor))
	runtime.KeepAlive(this)
}

func (this *QGraphicsView) ViewportUpdateMode() QGraphicsView__ViewportUpdateMode {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsView_viewportUpdateMode, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QGraphicsView__ViewportUpdateMode(__r)
}

func (this *QGraphicsView) SetViewportUpdateMode(mode QGraphicsView__ViewportUpdateMode) {
	purego.SyscallN(addr_QGraphicsView_setViewportUpdateMode, this.h, uintptr(mode))
	runtime.KeepAlive(this)
}

func (this *QGraphicsView) OptimizationFlags() QGraphicsView__OptimizationFlag {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsView_optimizationFlags, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QGraphicsView__OptimizationFlag(__r)
}

func (this *QGraphicsView) SetOptimizationFlag(flag QGraphicsView__OptimizationFlag) {
	purego.SyscallN(addr_QGraphicsView_setOptimizationFlag, this.h, uintptr(flag))
	runtime.KeepAlive(this)
}

func (this *QGraphicsView) SetOptimizationFlags(flags QGraphicsView__OptimizationFlag) {
	purego.SyscallN(addr_QGraphicsView_setOptimizationFlags, this.h, uintptr(flags))
	runtime.KeepAlive(this)
}

func (this *QGraphicsView) DragMode() QGraphicsView__DragMode {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsView_dragMode, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QGraphicsView__DragMode(__r)
}

func (this *QGraphicsView) SetDragMode(mode QGraphicsView__DragMode) {
	purego.SyscallN(addr_QGraphicsView_setDragMode, this.h, uintptr(mode))
	runtime.KeepAlive(this)
}

func (this *QGraphicsView) RubberBandSelectionMode() ItemSelectionMode {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsView_rubberBandSelectionMode, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return ItemSelectionMode(__r)
}

func (this *QGraphicsView) SetRubberBandSelectionMode(mode ItemSelectionMode) {
	purego.SyscallN(addr_QGraphicsView_setRubberBandSelectionMode, this.h, uintptr(mode))
	runtime.KeepAlive(this)
}

func (this *QGraphicsView) RubberBandRect() *QRect {
	__r, _, _ := purego.SyscallN(addr_QGraphicsView_rubberBandRect, this.h)
	runtime.KeepAlive(this)
	return newownedQRect(__r)
}

func (this *QGraphicsView) CacheMode() QGraphicsView__CacheModeFlag {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsView_cacheMode, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QGraphicsView__CacheModeFlag(__r)
}

func (this *QGraphicsView) SetCacheMode(mode QGraphicsView__CacheModeFlag) {
	purego.SyscallN(addr_QGraphicsView_setCacheMode, this.h, uintptr(mode))
	runtime.KeepAlive(this)
}

func (this *QGraphicsView) ResetCachedContent() {
	purego.SyscallN(addr_QGraphicsView_resetCachedContent, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsView) IsInteractive() bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsView_isInteractive, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGraphicsView) SetInteractive(allowed bool) {
	purego.SyscallN(addr_QGraphicsView_setInteractive, this.h, uintptr(b2i(allowed)))
	runtime.KeepAlive(this)
}

func (this *QGraphicsView) Scene() *QGraphicsScene {
	__r, _, _ := purego.SyscallN(addr_QGraphicsView_scene, this.h)
	runtime.KeepAlive(this)
	return bindQGraphicsScene(__r)
}

func (this *QGraphicsView) SetScene(scene *QGraphicsScene) {
	purego.SyscallN(addr_QGraphicsView_setScene, this.h, uintptr(scene.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(scene)
}

func (this *QGraphicsView) SceneRect() *QRectF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsView_sceneRect, this.h)
	runtime.KeepAlive(this)
	return newownedQRectF(__r)
}

func (this *QGraphicsView) SetSceneRect(rect *QRectF) {
	purego.SyscallN(addr_QGraphicsView_setSceneRect, this.h, uintptr(rect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
}

func (this *QGraphicsView) SetSceneRect2(x float64, y float64, w float64, h float64) {
	pg_QGraphicsView_setSceneRect2(this.h, x, y, w, h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsView) Transform() *QTransform {
	__r, _, _ := purego.SyscallN(addr_QGraphicsView_transform, this.h)
	runtime.KeepAlive(this)
	return newownedQTransform(__r)
}

func (this *QGraphicsView) ViewportTransform() *QTransform {
	__r, _, _ := purego.SyscallN(addr_QGraphicsView_viewportTransform, this.h)
	runtime.KeepAlive(this)
	return newownedQTransform(__r)
}

func (this *QGraphicsView) IsTransformed() bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsView_isTransformed, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGraphicsView) SetTransform(matrix *QTransform) {
	purego.SyscallN(addr_QGraphicsView_setTransform, this.h, uintptr(matrix.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(matrix)
}

func (this *QGraphicsView) ResetTransform() {
	purego.SyscallN(addr_QGraphicsView_resetTransform, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsView) Rotate(angle float64) {
	pg_QGraphicsView_rotate(this.h, angle)
	runtime.KeepAlive(this)
}

func (this *QGraphicsView) Scale(sx float64, sy float64) {
	pg_QGraphicsView_scale(this.h, sx, sy)
	runtime.KeepAlive(this)
}

func (this *QGraphicsView) Shear(sh float64, sv float64) {
	pg_QGraphicsView_shear(this.h, sh, sv)
	runtime.KeepAlive(this)
}

func (this *QGraphicsView) Translate(dx float64, dy float64) {
	pg_QGraphicsView_translate(this.h, dx, dy)
	runtime.KeepAlive(this)
}

func (this *QGraphicsView) CenterOn(pos *QPointF) {
	purego.SyscallN(addr_QGraphicsView_centerOn, this.h, uintptr(pos.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pos)
}

func (this *QGraphicsView) CenterOn2(x float64, y float64) {
	pg_QGraphicsView_centerOn2(this.h, x, y)
	runtime.KeepAlive(this)
}

func (this *QGraphicsView) CenterOnWithItem(item *QGraphicsItem) {
	purego.SyscallN(addr_QGraphicsView_centerOn3, this.h, uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QGraphicsView) EnsureVisible(rect *QRectF) {
	purego.SyscallN(addr_QGraphicsView_ensureVisible, this.h, uintptr(rect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
}

func (this *QGraphicsView) EnsureVisible2(x float64, y float64, w float64, h float64) {
	pg_QGraphicsView_ensureVisible2(this.h, x, y, w, h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsView) EnsureVisibleWithItem(item *QGraphicsItem) {
	purego.SyscallN(addr_QGraphicsView_ensureVisible3, this.h, uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QGraphicsView) FitInView(rect *QRectF) {
	purego.SyscallN(addr_QGraphicsView_fitInView, this.h, uintptr(rect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
}

func (this *QGraphicsView) FitInView2(x float64, y float64, w float64, h float64) {
	pg_QGraphicsView_fitInView2(this.h, x, y, w, h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsView) FitInViewWithItem(item *QGraphicsItem) {
	purego.SyscallN(addr_QGraphicsView_fitInView3, this.h, uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QGraphicsView) Render(painter *QPainter) {
	purego.SyscallN(addr_QGraphicsView_render, this.h, uintptr(painter.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
}

func (this *QGraphicsView) Items() []*QGraphicsItem {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGraphicsView_items, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QGraphicsItem, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = bindQGraphicsItem(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QGraphicsView) ItemsWithPos(pos *QPoint) []*QGraphicsItem {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGraphicsView_items2, this.h, uintptr(pos.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
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

func (this *QGraphicsView) Items2(x int, y int) []*QGraphicsItem {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGraphicsView_items3, this.h, uintptr(x), uintptr(y), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QGraphicsItem, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = bindQGraphicsItem(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QGraphicsView) ItemsWithRect(rect *QRect) []*QGraphicsItem {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGraphicsView_items4, this.h, uintptr(rect.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
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

func (this *QGraphicsView) Items3(x int, y int, w int, h int) []*QGraphicsItem {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGraphicsView_items5, this.h, uintptr(x), uintptr(y), uintptr(w), uintptr(h), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QGraphicsItem, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = bindQGraphicsItem(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QGraphicsView) ItemsWithPolygon(polygon *QPolygon) []*QGraphicsItem {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGraphicsView_items6, this.h, uintptr(polygon.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
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

func (this *QGraphicsView) ItemsWithPath(path *QPainterPath) []*QGraphicsItem {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGraphicsView_items7, this.h, uintptr(path.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
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

func (this *QGraphicsView) ItemAt(pos *QPoint) *QGraphicsItem {
	__r, _, _ := purego.SyscallN(addr_QGraphicsView_itemAt, this.h, uintptr(pos.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pos)
	return bindQGraphicsItem(__r)
}

func (this *QGraphicsView) ItemAt2(x int, y int) *QGraphicsItem {
	__r, _, _ := purego.SyscallN(addr_QGraphicsView_itemAt2, this.h, uintptr(x), uintptr(y))
	runtime.KeepAlive(this)
	return bindQGraphicsItem(__r)
}

func (this *QGraphicsView) MapToScene(point *QPoint) *QPointF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsView_mapToScene, this.h, uintptr(point.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(point)
	return newownedQPointF(__r)
}

func (this *QGraphicsView) MapToSceneWithRect(rect *QRect) *QPolygonF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsView_mapToScene2, this.h, uintptr(rect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
	return newownedQPolygonF(__r)
}

func (this *QGraphicsView) MapToSceneWithPolygon(polygon *QPolygon) *QPolygonF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsView_mapToScene3, this.h, uintptr(polygon.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(polygon)
	return newownedQPolygonF(__r)
}

func (this *QGraphicsView) MapToSceneWithPath(path *QPainterPath) *QPainterPath {
	__r, _, _ := purego.SyscallN(addr_QGraphicsView_mapToScene4, this.h, uintptr(path.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(path)
	return newownedQPainterPath(__r)
}

func (this *QGraphicsView) MapFromScene(point *QPointF) *QPoint {
	__r, _, _ := purego.SyscallN(addr_QGraphicsView_mapFromScene, this.h, uintptr(point.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(point)
	return newownedQPoint(__r)
}

func (this *QGraphicsView) MapFromSceneWithRect(rect *QRectF) *QPolygon {
	__r, _, _ := purego.SyscallN(addr_QGraphicsView_mapFromScene2, this.h, uintptr(rect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
	return newownedQPolygon(__r)
}

func (this *QGraphicsView) MapFromSceneWithPolygon(polygon *QPolygonF) *QPolygon {
	__r, _, _ := purego.SyscallN(addr_QGraphicsView_mapFromScene3, this.h, uintptr(polygon.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(polygon)
	return newownedQPolygon(__r)
}

func (this *QGraphicsView) MapFromSceneWithPath(path *QPainterPath) *QPainterPath {
	__r, _, _ := purego.SyscallN(addr_QGraphicsView_mapFromScene4, this.h, uintptr(path.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(path)
	return newownedQPainterPath(__r)
}

func (this *QGraphicsView) MapToScene2(x int, y int) *QPointF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsView_mapToScene5, this.h, uintptr(x), uintptr(y))
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QGraphicsView) MapToScene3(x int, y int, w int, h int) *QPolygonF {
	__r, _, _ := purego.SyscallN(addr_QGraphicsView_mapToScene6, this.h, uintptr(x), uintptr(y), uintptr(w), uintptr(h))
	runtime.KeepAlive(this)
	return newownedQPolygonF(__r)
}

func (this *QGraphicsView) MapFromScene2(x float64, y float64) *QPoint {
	__r := pg_QGraphicsView_mapFromScene5(this.h, x, y)
	runtime.KeepAlive(this)
	return newownedQPoint(__r)
}

func (this *QGraphicsView) MapFromScene3(x float64, y float64, w float64, h float64) *QPolygon {
	__r := pg_QGraphicsView_mapFromScene6(this.h, x, y, w, h)
	runtime.KeepAlive(this)
	return newownedQPolygon(__r)
}

func (this *QGraphicsView) BackgroundBrush() *QBrush {
	__r, _, _ := purego.SyscallN(addr_QGraphicsView_backgroundBrush, this.h)
	runtime.KeepAlive(this)
	return newownedQBrush(__r)
}

func (this *QGraphicsView) SetBackgroundBrush(brush *QBrush) {
	purego.SyscallN(addr_QGraphicsView_setBackgroundBrush, this.h, uintptr(brush.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(brush)
}

func (this *QGraphicsView) ForegroundBrush() *QBrush {
	__r, _, _ := purego.SyscallN(addr_QGraphicsView_foregroundBrush, this.h)
	runtime.KeepAlive(this)
	return newownedQBrush(__r)
}

func (this *QGraphicsView) SetForegroundBrush(brush *QBrush) {
	purego.SyscallN(addr_QGraphicsView_setForegroundBrush, this.h, uintptr(brush.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(brush)
}

func (this *QGraphicsView) UpdateScene(rects []*QRectF) {
	rects_arr := make([]uintptr, len(rects))
	for __i, __e := range rects {
		rects_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var rects_p uintptr
	if len(rects_arr) > 0 {
		rects_p = uintptr(unsafe.Pointer(&rects_arr[0]))
	}
	purego.SyscallN(addr_QGraphicsView_updateScene, this.h, rects_p, uintptr(len(rects)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rects)
	runtime.KeepAlive(rects_arr)
}

func (this *QGraphicsView) InvalidateScene() {
	purego.SyscallN(addr_QGraphicsView_invalidateScene, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsView) UpdateSceneRect(rect *QRectF) {
	purego.SyscallN(addr_QGraphicsView_updateSceneRect, this.h, uintptr(rect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
}

func (this *QGraphicsView) RubberBandChanged(viewportRect *QRect, fromScenePoint *QPointF, toScenePoint *QPointF) {
	purego.SyscallN(addr_QGraphicsView_rubberBandChanged, this.h, uintptr(viewportRect.UnsafePointer()), uintptr(fromScenePoint.UnsafePointer()), uintptr(toScenePoint.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(viewportRect)
	runtime.KeepAlive(fromScenePoint)
	runtime.KeepAlive(toScenePoint)
}

func (this *QGraphicsView) SetViewportMargins(left int, top int, right int, bottom int) {
	purego.SyscallN(addr_QGraphicsView_setViewportMargins, this.h, uintptr(left), uintptr(top), uintptr(right), uintptr(bottom))
	runtime.KeepAlive(this)
}

func (this *QGraphicsView) SetViewportMarginsWithMargins(margins *QMargins) {
	purego.SyscallN(addr_QGraphicsView_setViewportMargins2, this.h, uintptr(margins.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(margins)
}

func (this *QGraphicsView) ViewportMargins() *QMargins {
	__r, _, _ := purego.SyscallN(addr_QGraphicsView_viewportMargins, this.h)
	runtime.KeepAlive(this)
	return newownedQMargins(__r)
}

func (this *QGraphicsView) DrawFrame(arg1 *QPainter) {
	purego.SyscallN(addr_QGraphicsView_drawFrame, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func QGraphicsView_SetTabOrder(arg1 *QWidget, arg2 *QWidget) {
	purego.SyscallN(addr_QGraphicsView_setTabOrder, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()))
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
}

func QGraphicsView_MouseGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QGraphicsView_mouseGrabber)
	return bindQWidget(__r)
}

func QGraphicsView_KeyboardGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QGraphicsView_keyboardGrabber)
	return bindQWidget(__r)
}

func QGraphicsView_Find(arg1 uint64) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QGraphicsView_find, uintptr(arg1))
	return bindQWidget(__r)
}

func QGraphicsView_CreateWindowContainer(window *QWindow) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QGraphicsView_createWindowContainer, uintptr(window.UnsafePointer()))
	runtime.KeepAlive(window)
	return bindQWidget(__r)
}

func (this *QGraphicsView) UpdateMicroFocus() {
	purego.SyscallN(addr_QGraphicsView_updateMicroFocus, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsView) Create() {
	purego.SyscallN(addr_QGraphicsView_create, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsView) Destroy() {
	purego.SyscallN(addr_QGraphicsView_destroy, this.h)
	runtime.KeepAlive(this)
}

func (this *QGraphicsView) FocusNextChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsView_focusNextChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGraphicsView) FocusPreviousChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsView_focusPreviousChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QGraphicsView_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QGraphicsView_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QGraphicsView_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QGraphicsView_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QGraphicsView) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QGraphicsView_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QGraphicsView) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsView_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QGraphicsView) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QGraphicsView_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QGraphicsView) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsView_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QGraphicsView_DevicePixelRatioFScale() float64 {
	__r := pg_QGraphicsView_devicePixelRatioFScale()
	return __r
}

func QGraphicsView_EncodeMetricF(metric QPaintDevice__PaintDeviceMetric, value float64) int {
	__r := pg_QGraphicsView_encodeMetricF(int32(metric), value)
	return int(__r)
}

func (this *QGraphicsView) GetDecodedMetricF(metricA QPaintDevice__PaintDeviceMetric, metricB QPaintDevice__PaintDeviceMetric) float64 {
	__r := pg_QGraphicsView_getDecodedMetricF(this.h, int32(metricA), int32(metricB))
	runtime.KeepAlive(this)
	return __r
}

func QGraphicsView_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGraphicsView_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QGraphicsView_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGraphicsView_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func (this *QGraphicsView) SetRenderHint2(hint QPainter__RenderHint, enabled bool) {
	purego.SyscallN(addr_QGraphicsView_setRenderHint2, this.h, uintptr(hint), uintptr(b2i(enabled)))
	runtime.KeepAlive(this)
}

func (this *QGraphicsView) SetOptimizationFlag2(flag QGraphicsView__OptimizationFlag, enabled bool) {
	purego.SyscallN(addr_QGraphicsView_setOptimizationFlag2, this.h, uintptr(flag), uintptr(b2i(enabled)))
	runtime.KeepAlive(this)
}

func (this *QGraphicsView) SetTransform2(matrix *QTransform, combine bool) {
	purego.SyscallN(addr_QGraphicsView_setTransform2, this.h, uintptr(matrix.UnsafePointer()), uintptr(b2i(combine)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(matrix)
}

func (this *QGraphicsView) EnsureVisible3(rect *QRectF, xmargin int) {
	purego.SyscallN(addr_QGraphicsView_ensureVisible4, this.h, uintptr(rect.UnsafePointer()), uintptr(xmargin))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
}

func (this *QGraphicsView) EnsureVisible4(rect *QRectF, xmargin int, ymargin int) {
	purego.SyscallN(addr_QGraphicsView_ensureVisible5, this.h, uintptr(rect.UnsafePointer()), uintptr(xmargin), uintptr(ymargin))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
}

func (this *QGraphicsView) EnsureVisible5(x float64, y float64, w float64, h float64, xmargin int) {
	pg_QGraphicsView_ensureVisible6(this.h, x, y, w, h, int32(xmargin))
	runtime.KeepAlive(this)
}

func (this *QGraphicsView) EnsureVisible6(x float64, y float64, w float64, h float64, xmargin int, ymargin int) {
	pg_QGraphicsView_ensureVisible7(this.h, x, y, w, h, int32(xmargin), int32(ymargin))
	runtime.KeepAlive(this)
}

func (this *QGraphicsView) EnsureVisible7(item *QGraphicsItem, xmargin int) {
	purego.SyscallN(addr_QGraphicsView_ensureVisible8, this.h, uintptr(item.UnsafePointer()), uintptr(xmargin))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QGraphicsView) EnsureVisible8(item *QGraphicsItem, xmargin int, ymargin int) {
	purego.SyscallN(addr_QGraphicsView_ensureVisible9, this.h, uintptr(item.UnsafePointer()), uintptr(xmargin), uintptr(ymargin))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QGraphicsView) FitInView3(rect *QRectF, aspectRadioMode AspectRatioMode) {
	purego.SyscallN(addr_QGraphicsView_fitInView4, this.h, uintptr(rect.UnsafePointer()), uintptr(aspectRadioMode))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
}

func (this *QGraphicsView) FitInView4(x float64, y float64, w float64, h float64, aspectRadioMode AspectRatioMode) {
	pg_QGraphicsView_fitInView5(this.h, x, y, w, h, int32(aspectRadioMode))
	runtime.KeepAlive(this)
}

func (this *QGraphicsView) FitInView5(item *QGraphicsItem, aspectRadioMode AspectRatioMode) {
	purego.SyscallN(addr_QGraphicsView_fitInView6, this.h, uintptr(item.UnsafePointer()), uintptr(aspectRadioMode))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QGraphicsView) Render2(painter *QPainter, target *QRectF) {
	purego.SyscallN(addr_QGraphicsView_render2, this.h, uintptr(painter.UnsafePointer()), uintptr(target.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
	runtime.KeepAlive(target)
}

func (this *QGraphicsView) Render3(painter *QPainter, target *QRectF, source *QRect) {
	purego.SyscallN(addr_QGraphicsView_render3, this.h, uintptr(painter.UnsafePointer()), uintptr(target.UnsafePointer()), uintptr(source.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
	runtime.KeepAlive(target)
	runtime.KeepAlive(source)
}

func (this *QGraphicsView) Render4(painter *QPainter, target *QRectF, source *QRect, aspectRatioMode AspectRatioMode) {
	purego.SyscallN(addr_QGraphicsView_render4, this.h, uintptr(painter.UnsafePointer()), uintptr(target.UnsafePointer()), uintptr(source.UnsafePointer()), uintptr(aspectRatioMode))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
	runtime.KeepAlive(target)
	runtime.KeepAlive(source)
}

func (this *QGraphicsView) Items4(rect *QRect, mode ItemSelectionMode) []*QGraphicsItem {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGraphicsView_items8, this.h, uintptr(rect.UnsafePointer()), uintptr(mode), uintptr(unsafe.Pointer(&__n)))
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

func (this *QGraphicsView) Items5(x int, y int, w int, h int, mode ItemSelectionMode) []*QGraphicsItem {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGraphicsView_items9, this.h, uintptr(x), uintptr(y), uintptr(w), uintptr(h), uintptr(mode), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QGraphicsItem, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = bindQGraphicsItem(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QGraphicsView) Items6(polygon *QPolygon, mode ItemSelectionMode) []*QGraphicsItem {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGraphicsView_items10, this.h, uintptr(polygon.UnsafePointer()), uintptr(mode), uintptr(unsafe.Pointer(&__n)))
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

func (this *QGraphicsView) Items7(path *QPainterPath, mode ItemSelectionMode) []*QGraphicsItem {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QGraphicsView_items11, this.h, uintptr(path.UnsafePointer()), uintptr(mode), uintptr(unsafe.Pointer(&__n)))
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

func (this *QGraphicsView) InvalidateSceneWithRect(rect *QRectF) {
	purego.SyscallN(addr_QGraphicsView_invalidateScene2, this.h, uintptr(rect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
}

func (this *QGraphicsView) InvalidateScene2(rect *QRectF, layers QGraphicsScene__SceneLayer) {
	purego.SyscallN(addr_QGraphicsView_invalidateScene3, this.h, uintptr(rect.UnsafePointer()), uintptr(layers))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
}

func QGraphicsView_CreateWindowContainer2(window *QWindow, parent *QWidget) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QGraphicsView_createWindowContainer2, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func QGraphicsView_CreateWindowContainer3(window *QWindow, parent *QWidget, flags WindowType) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QGraphicsView_createWindowContainer3, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()), uintptr(flags))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func (this *QGraphicsView) UpdateMicroFocusWithQuery(query InputMethodQuery) {
	purego.SyscallN(addr_QGraphicsView_updateMicroFocus2, this.h, uintptr(query))
	runtime.KeepAlive(this)
}

func (this *QGraphicsView) CreateWith(arg1 uint64) {
	purego.SyscallN(addr_QGraphicsView_create2, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QGraphicsView) Create2(arg1 uint64, initializeWindow bool) {
	purego.SyscallN(addr_QGraphicsView_create3, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)))
	runtime.KeepAlive(this)
}

func (this *QGraphicsView) Create3(arg1 uint64, initializeWindow bool, destroyOldWindow bool) {
	purego.SyscallN(addr_QGraphicsView_create4, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)), uintptr(b2i(destroyOldWindow)))
	runtime.KeepAlive(this)
}

func (this *QGraphicsView) DestroyWithDestroyWindow(destroyWindow bool) {
	purego.SyscallN(addr_QGraphicsView_destroy2, this.h, uintptr(b2i(destroyWindow)))
	runtime.KeepAlive(this)
}

func (this *QGraphicsView) Destroy2(destroyWindow bool, destroySubWindows bool) {
	purego.SyscallN(addr_QGraphicsView_destroy3, this.h, uintptr(b2i(destroyWindow)), uintptr(b2i(destroySubWindows)))
	runtime.KeepAlive(this)
}

func QGraphicsView_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QGraphicsView_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QGraphicsView_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QGraphicsView_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QGraphicsView) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QGraphicsView_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QGraphicsView) SizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QGraphicsView_sizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QGraphicsView) InputMethodQueryDefault(query InputMethodQuery) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QGraphicsView_inputMethodQuery_default, this.h, uintptr(query))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QGraphicsView) SetupViewportDefault(widget *QWidget) {
	purego.SyscallN(addr_QGraphicsView_setupViewport_default, this.h, uintptr(widget.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(widget)
}

func (this *QGraphicsView) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsView_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QGraphicsView) ViewportEventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsView_viewportEvent_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QGraphicsView) ContextMenuEventDefault(event *QContextMenuEvent) {
	purego.SyscallN(addr_QGraphicsView_contextMenuEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsView) DragEnterEventDefault(event *QDragEnterEvent) {
	purego.SyscallN(addr_QGraphicsView_dragEnterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsView) DragLeaveEventDefault(event *QDragLeaveEvent) {
	purego.SyscallN(addr_QGraphicsView_dragLeaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsView) DragMoveEventDefault(event *QDragMoveEvent) {
	purego.SyscallN(addr_QGraphicsView_dragMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsView) DropEventDefault(event *QDropEvent) {
	purego.SyscallN(addr_QGraphicsView_dropEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsView) FocusInEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QGraphicsView_focusInEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsView) FocusNextPrevChildDefault(next bool) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsView_focusNextPrevChild_default, this.h, uintptr(b2i(next)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGraphicsView) FocusOutEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QGraphicsView_focusOutEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsView) KeyPressEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QGraphicsView_keyPressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsView) KeyReleaseEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QGraphicsView_keyReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsView) MouseDoubleClickEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QGraphicsView_mouseDoubleClickEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsView) MousePressEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QGraphicsView_mousePressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsView) MouseMoveEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QGraphicsView_mouseMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsView) MouseReleaseEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QGraphicsView_mouseReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsView) WheelEventDefault(event *QWheelEvent) {
	purego.SyscallN(addr_QGraphicsView_wheelEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsView) PaintEventDefault(event *QPaintEvent) {
	purego.SyscallN(addr_QGraphicsView_paintEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsView) ResizeEventDefault(event *QResizeEvent) {
	purego.SyscallN(addr_QGraphicsView_resizeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsView) ScrollContentsByDefault(dx int, dy int) {
	purego.SyscallN(addr_QGraphicsView_scrollContentsBy_default, this.h, uintptr(dx), uintptr(dy))
	runtime.KeepAlive(this)
}

func (this *QGraphicsView) ShowEventDefault(event *QShowEvent) {
	purego.SyscallN(addr_QGraphicsView_showEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsView) InputMethodEventDefault(event *QInputMethodEvent) {
	purego.SyscallN(addr_QGraphicsView_inputMethodEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsView) DrawBackgroundDefault(painter *QPainter, rect *QRectF) {
	purego.SyscallN(addr_QGraphicsView_drawBackground_default, this.h, uintptr(painter.UnsafePointer()), uintptr(rect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
	runtime.KeepAlive(rect)
}

func (this *QGraphicsView) DrawForegroundDefault(painter *QPainter, rect *QRectF) {
	purego.SyscallN(addr_QGraphicsView_drawForeground_default, this.h, uintptr(painter.UnsafePointer()), uintptr(rect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
	runtime.KeepAlive(rect)
}

func (this *QGraphicsView) MinimumSizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QGraphicsView_minimumSizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QGraphicsView) EventFilterDefault(arg1 *QObject, arg2 *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsView_eventFilter_default, this.h, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
	return __r != 0
}

func (this *QGraphicsView) ViewportSizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QGraphicsView_viewportSizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QGraphicsView) ChangeEventDefault(arg1 *QEvent) {
	purego.SyscallN(addr_QGraphicsView_changeEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QGraphicsView) InitStyleOptionDefault(option *QStyleOptionFrame) {
	purego.SyscallN(addr_QGraphicsView_initStyleOption_default, this.h, uintptr(option.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(option)
}

func (this *QGraphicsView) DevTypeDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsView_devType_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QGraphicsView) SetVisibleDefault(visible bool) {
	purego.SyscallN(addr_QGraphicsView_setVisible_default, this.h, uintptr(b2i(visible)))
	runtime.KeepAlive(this)
}

func (this *QGraphicsView) HeightForWidthDefault(arg1 int) int {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsView_heightForWidth_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QGraphicsView) HasHeightForWidthDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsView_hasHeightForWidth_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QGraphicsView) PaintEngineDefault() *QPaintEngine {
	__r, _, _ := purego.SyscallN(addr_QGraphicsView_paintEngine_default, this.h)
	runtime.KeepAlive(this)
	return bindQPaintEngine(__r)
}

func (this *QGraphicsView) EnterEventDefault(event *QEnterEvent) {
	purego.SyscallN(addr_QGraphicsView_enterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsView) LeaveEventDefault(event *QEvent) {
	purego.SyscallN(addr_QGraphicsView_leaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsView) MoveEventDefault(event *QMoveEvent) {
	purego.SyscallN(addr_QGraphicsView_moveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsView) CloseEventDefault(event *QCloseEvent) {
	purego.SyscallN(addr_QGraphicsView_closeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsView) TabletEventDefault(event *QTabletEvent) {
	purego.SyscallN(addr_QGraphicsView_tabletEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsView) ActionEventDefault(event *QActionEvent) {
	purego.SyscallN(addr_QGraphicsView_actionEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsView) HideEventDefault(event *QHideEvent) {
	purego.SyscallN(addr_QGraphicsView_hideEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsView) MetricDefault(arg1 QPaintDevice__PaintDeviceMetric) int {
	__r0, _, _ := purego.SyscallN(addr_QGraphicsView_metric_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QGraphicsView) InitPainterDefault(painter *QPainter) {
	purego.SyscallN(addr_QGraphicsView_initPainter_default, this.h, uintptr(painter.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
}

func (this *QGraphicsView) RedirectedDefault(offset *QPoint) *QPaintDevice {
	__r, _, _ := purego.SyscallN(addr_QGraphicsView_redirected_default, this.h, uintptr(offset.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(offset)
	return bindQPaintDevice(__r)
}

func (this *QGraphicsView) SharedPainterDefault() *QPainter {
	__r, _, _ := purego.SyscallN(addr_QGraphicsView_sharedPainter_default, this.h)
	runtime.KeepAlive(this)
	return bindQPainter(__r)
}

func (this *QGraphicsView) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QGraphicsView_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsView) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QGraphicsView_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsView) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QGraphicsView_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QGraphicsView) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QGraphicsView_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QGraphicsView) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QGraphicsView_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QGraphicsView) OnRubberBandChanged(fn func(*QRect, *QPointF, *QPointF)) int64 {
	id := registerCallback(func(viewportRect uintptr, fromScenePoint uintptr, toScenePoint uintptr) { fn(UnsafeNewQRect(unsafe.Pointer(viewportRect)), UnsafeNewQPointF(unsafe.Pointer(fromScenePoint)), UnsafeNewQPointF(unsafe.Pointer(toScenePoint))) })
	pg_QGraphicsView_connect_rubberBandChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsView) OnWindowTitleChanged(fn func(string)) int64 {
	id := registerCallback(func(title uintptr, title_len int64) { fn(pgStrNoFree(title, title_len)) })
	pg_QGraphicsView_connect_windowTitleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsView) OnWindowIconChanged(fn func(*QIcon)) int64 {
	id := registerCallback(func(icon uintptr) { fn(UnsafeNewQIcon(unsafe.Pointer(icon))) })
	pg_QGraphicsView_connect_windowIconChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsView) OnWindowIconTextChanged(fn func(string)) int64 {
	id := registerCallback(func(iconText uintptr, iconText_len int64) { fn(pgStrNoFree(iconText, iconText_len)) })
	pg_QGraphicsView_connect_windowIconTextChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsView) OnCustomContextMenuRequested(fn func(*QPoint)) int64 {
	id := registerCallback(func(pos uintptr) { fn(UnsafeNewQPoint(unsafe.Pointer(pos))) })
	pg_QGraphicsView_connect_customContextMenuRequested(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsView) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QGraphicsView_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsView) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QGraphicsView_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsView) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QGraphicsView_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QGraphicsView) Disconnect(handle int64) {
	pg_QGraphicsView__disconnect(handle)
	unregisterCallback(handle)
}

func QGraphicsView_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case 729690261457684101:
		return 1
	case 7345853097340709638:
		return 2
	case 4531996381736730688:
		return 3
	case 1237236456312824991:
		return 4
	case -1971465653061861519:
		return 5
	case -6698043222227006015:
		return 6
	case 206122715919917589:
		return 7
	case -3139904158373725006:
		return 8
	case -1781545137352745978:
		return 9
	case 9107058682754591716:
		return 10
	case 975043767033552552:
		return 11
	case 5898328517505784191:
		return 12
	case -385596010576048993:
		return 13
	case 1023791977679793185:
		return 14
	case 772598475950050299:
		return 15
	case -4029073905697023887:
		return 16
	case -332090898078543319:
		return 17
	case -4255494004099806581:
		return 18
	case -4415920049297795213:
		return 19
	case 1919988340934511316:
		return 20
	case 5290442163167458465:
		return 21
	case -4302917285655014661:
		return 22
	case -6118001150126564745:
		return 23
	case -6825369174751916:
		return 24
	case 5438819661942843222:
		return 25
	case -3076319991076358961:
		return 26
	case 8267095530709235866:
		return 27
	case 7049264599837532491:
		return 28
	case 5217623514630858297:
		return 29
	case 503579067252048239:
		return 30
	case 4959376063826165943:
		return 31
	case -488693769771308531:
		return 32
	case -8659679893868524044:
		return 33
	case 2110182719375245575:
		return 34
	case 4311534060495737339:
		return 35
	case -8266762397995423915:
		return 36
	case 6398007404756346381:
		return 37
	case 6316270346668237895:
		return 38
	case -551070009623711640:
		return 39
	case -9161949150489525396:
		return 40
	case -1989937050019464783:
		return 41
	case -212068866860701141:
		return 42
	case 1697444140088555581:
		return 43
	case -199718707876719129:
		return 44
	case -2039166586246443706:
		return 45
	case 427629649267311777:
		return 46
	case 6477500296980050210:
		return 47
	case 4127240633310655048:
		return 48
	case 5792771328741774991:
		return 49
	case -5409927581812374498:
		return 50
	case -1185315367526857833:
		return 51
	case 3769868229123868474:
		return 52
	case 7379029982170503008:
		return 53
	case -568369772982215962:
		return 54
	}
	return -1
}

func QGraphicsView_enable(h uintptr, idx int32) { pg_QGraphicsView__set_override(h, idx) }

func (this *QGraphicsView) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QGraphicsView) OnSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 729690261457684101, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sizeHint"
}

func (this *QGraphicsView) OnInputMethodQuery(fn func(InputMethodQuery) *QVariant) {
	registerVirt(this.UnsafeGoID(), 7345853097340709638, func(query int32) uintptr { return uintptr(fn(InputMethodQuery(query)).UnsafePointer()) }) // slot "inputMethodQuery"
}

func (this *QGraphicsView) OnSetupViewport(fn func(*QWidget)) {
	registerVirt(this.UnsafeGoID(), 4531996381736730688, func(widget uintptr) { fn(UnsafeNewQWidget(unsafe.Pointer(widget))) }) // slot "setupViewport"
}

func (this *QGraphicsView) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QGraphicsView) OnViewportEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), -1971465653061861519, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "viewportEvent"
}

func (this *QGraphicsView) OnContextMenuEvent(fn func(*QContextMenuEvent)) {
	registerVirt(this.UnsafeGoID(), -6698043222227006015, func(event uintptr) { fn(UnsafeNewQContextMenuEvent(unsafe.Pointer(event))) }) // slot "contextMenuEvent"
}

func (this *QGraphicsView) OnDragEnterEvent(fn func(*QDragEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 206122715919917589, func(event uintptr) { fn(UnsafeNewQDragEnterEvent(unsafe.Pointer(event))) }) // slot "dragEnterEvent"
}

func (this *QGraphicsView) OnDragLeaveEvent(fn func(*QDragLeaveEvent)) {
	registerVirt(this.UnsafeGoID(), -3139904158373725006, func(event uintptr) { fn(UnsafeNewQDragLeaveEvent(unsafe.Pointer(event))) }) // slot "dragLeaveEvent"
}

func (this *QGraphicsView) OnDragMoveEvent(fn func(*QDragMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -1781545137352745978, func(event uintptr) { fn(UnsafeNewQDragMoveEvent(unsafe.Pointer(event))) }) // slot "dragMoveEvent"
}

func (this *QGraphicsView) OnDropEvent(fn func(*QDropEvent)) {
	registerVirt(this.UnsafeGoID(), 9107058682754591716, func(event uintptr) { fn(UnsafeNewQDropEvent(unsafe.Pointer(event))) }) // slot "dropEvent"
}

func (this *QGraphicsView) OnFocusInEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), 975043767033552552, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusInEvent"
}

func (this *QGraphicsView) OnFocusNextPrevChild(fn func(bool) bool) {
	registerVirt(this.UnsafeGoID(), 5898328517505784191, func(next int32) int32 { return b2i(fn(next != 0)) }) // slot "focusNextPrevChild"
}

func (this *QGraphicsView) OnFocusOutEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), -385596010576048993, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusOutEvent"
}

func (this *QGraphicsView) OnKeyPressEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 1023791977679793185, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyPressEvent"
}

func (this *QGraphicsView) OnKeyReleaseEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 772598475950050299, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyReleaseEvent"
}

func (this *QGraphicsView) OnMouseDoubleClickEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4029073905697023887, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseDoubleClickEvent"
}

func (this *QGraphicsView) OnMousePressEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -332090898078543319, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mousePressEvent"
}

func (this *QGraphicsView) OnMouseMoveEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4255494004099806581, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseMoveEvent"
}

func (this *QGraphicsView) OnMouseReleaseEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4415920049297795213, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseReleaseEvent"
}

func (this *QGraphicsView) OnWheelEvent(fn func(*QWheelEvent)) {
	registerVirt(this.UnsafeGoID(), 1919988340934511316, func(event uintptr) { fn(UnsafeNewQWheelEvent(unsafe.Pointer(event))) }) // slot "wheelEvent"
}

func (this *QGraphicsView) OnPaintEvent(fn func(*QPaintEvent)) {
	registerVirt(this.UnsafeGoID(), 5290442163167458465, func(event uintptr) { fn(UnsafeNewQPaintEvent(unsafe.Pointer(event))) }) // slot "paintEvent"
}

func (this *QGraphicsView) OnResizeEvent(fn func(*QResizeEvent)) {
	registerVirt(this.UnsafeGoID(), -4302917285655014661, func(event uintptr) { fn(UnsafeNewQResizeEvent(unsafe.Pointer(event))) }) // slot "resizeEvent"
}

func (this *QGraphicsView) OnScrollContentsBy(fn func(int, int)) {
	registerVirt(this.UnsafeGoID(), -6118001150126564745, func(dx int32, dy int32) { fn(int(dx), int(dy)) }) // slot "scrollContentsBy"
}

func (this *QGraphicsView) OnShowEvent(fn func(*QShowEvent)) {
	registerVirt(this.UnsafeGoID(), -6825369174751916, func(event uintptr) { fn(UnsafeNewQShowEvent(unsafe.Pointer(event))) }) // slot "showEvent"
}

func (this *QGraphicsView) OnInputMethodEvent(fn func(*QInputMethodEvent)) {
	registerVirt(this.UnsafeGoID(), 5438819661942843222, func(event uintptr) { fn(UnsafeNewQInputMethodEvent(unsafe.Pointer(event))) }) // slot "inputMethodEvent"
}

func (this *QGraphicsView) OnDrawBackground(fn func(*QPainter, *QRectF)) {
	registerVirt(this.UnsafeGoID(), -3076319991076358961, func(painter uintptr, rect uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter)), UnsafeNewQRectF(unsafe.Pointer(rect))) }) // slot "drawBackground"
}

func (this *QGraphicsView) OnDrawForeground(fn func(*QPainter, *QRectF)) {
	registerVirt(this.UnsafeGoID(), 8267095530709235866, func(painter uintptr, rect uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter)), UnsafeNewQRectF(unsafe.Pointer(rect))) }) // slot "drawForeground"
}

func (this *QGraphicsView) OnMinimumSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 7049264599837532491, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "minimumSizeHint"
}

func (this *QGraphicsView) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(arg1 uintptr, arg2 uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(arg1)), UnsafeNewQEvent(unsafe.Pointer(arg2)))) }) // slot "eventFilter"
}

func (this *QGraphicsView) OnViewportSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 503579067252048239, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "viewportSizeHint"
}

func (this *QGraphicsView) OnChangeEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 4959376063826165943, func(arg1 uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(arg1))) }) // slot "changeEvent"
}

func (this *QGraphicsView) OnInitStyleOption(fn func(*QStyleOptionFrame)) {
	registerVirt(this.UnsafeGoID(), -488693769771308531, func(option uintptr) { fn(UnsafeNewQStyleOptionFrame(unsafe.Pointer(option))) }) // slot "initStyleOption"
}

func (this *QGraphicsView) OnDevType(fn func() int) {
	registerVirt(this.UnsafeGoID(), -8659679893868524044, func() int32 { return int32(fn()) }) // slot "devType"
}

func (this *QGraphicsView) OnSetVisible(fn func(bool)) {
	registerVirt(this.UnsafeGoID(), 2110182719375245575, func(visible int32) { fn(visible != 0) }) // slot "setVisible"
}

func (this *QGraphicsView) OnHeightForWidth(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), 4311534060495737339, func(arg1 int32) int32 { return int32(fn(int(arg1))) }) // slot "heightForWidth"
}

func (this *QGraphicsView) OnHasHeightForWidth(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -8266762397995423915, func() int32 { return b2i(fn()) }) // slot "hasHeightForWidth"
}

func (this *QGraphicsView) OnPaintEngine(fn func() *QPaintEngine) {
	registerVirt(this.UnsafeGoID(), 6398007404756346381, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "paintEngine"
}

func (this *QGraphicsView) OnEnterEvent(fn func(*QEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 6316270346668237895, func(event uintptr) { fn(UnsafeNewQEnterEvent(unsafe.Pointer(event))) }) // slot "enterEvent"
}

func (this *QGraphicsView) OnLeaveEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), -551070009623711640, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "leaveEvent"
}

func (this *QGraphicsView) OnMoveEvent(fn func(*QMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -9161949150489525396, func(event uintptr) { fn(UnsafeNewQMoveEvent(unsafe.Pointer(event))) }) // slot "moveEvent"
}

func (this *QGraphicsView) OnCloseEvent(fn func(*QCloseEvent)) {
	registerVirt(this.UnsafeGoID(), -1989937050019464783, func(event uintptr) { fn(UnsafeNewQCloseEvent(unsafe.Pointer(event))) }) // slot "closeEvent"
}

func (this *QGraphicsView) OnTabletEvent(fn func(*QTabletEvent)) {
	registerVirt(this.UnsafeGoID(), -212068866860701141, func(event uintptr) { fn(UnsafeNewQTabletEvent(unsafe.Pointer(event))) }) // slot "tabletEvent"
}

func (this *QGraphicsView) OnActionEvent(fn func(*QActionEvent)) {
	registerVirt(this.UnsafeGoID(), 1697444140088555581, func(event uintptr) { fn(UnsafeNewQActionEvent(unsafe.Pointer(event))) }) // slot "actionEvent"
}

func (this *QGraphicsView) OnHideEvent(fn func(*QHideEvent)) {
	registerVirt(this.UnsafeGoID(), -199718707876719129, func(event uintptr) { fn(UnsafeNewQHideEvent(unsafe.Pointer(event))) }) // slot "hideEvent"
}

func (this *QGraphicsView) OnNativeEvent(fn func(*QByteArray, unsafe.Pointer, *int64) bool) {
	registerVirt(this.UnsafeGoID(), -2039166586246443706, func(eventType uintptr, message uintptr, result uintptr) int32 { return b2i(fn(UnsafeNewQByteArray(unsafe.Pointer(eventType)), unsafe.Pointer(message), (*int64)(unsafe.Pointer(result)))) }) // slot "nativeEvent"
}

func (this *QGraphicsView) OnMetric(fn func(QPaintDevice__PaintDeviceMetric) int) {
	registerVirt(this.UnsafeGoID(), 427629649267311777, func(arg1 int32) int32 { return int32(fn(QPaintDevice__PaintDeviceMetric(arg1))) }) // slot "metric"
}

func (this *QGraphicsView) OnInitPainter(fn func(*QPainter)) {
	registerVirt(this.UnsafeGoID(), 6477500296980050210, func(painter uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter))) }) // slot "initPainter"
}

func (this *QGraphicsView) OnRedirected(fn func(*QPoint) *QPaintDevice) {
	registerVirt(this.UnsafeGoID(), 4127240633310655048, func(offset uintptr) uintptr { return uintptr(fn(UnsafeNewQPoint(unsafe.Pointer(offset))).UnsafePointer()) }) // slot "redirected"
}

func (this *QGraphicsView) OnSharedPainter(fn func() *QPainter) {
	registerVirt(this.UnsafeGoID(), 5792771328741774991, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sharedPainter"
}

func (this *QGraphicsView) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QGraphicsView) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QGraphicsView) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QGraphicsView) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QGraphicsView) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QGraphicsView) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QGraphicsView_delete(this.h)
	}
	this.UnsafeRelease()
}
