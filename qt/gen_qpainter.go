package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QPainter_new uintptr
	addr_QPainter_new2 uintptr
	addr_QPainter_device uintptr
	addr_QPainter_begin uintptr
	addr_QPainter_end uintptr
	addr_QPainter_isActive uintptr
	addr_QPainter_setCompositionMode uintptr
	addr_QPainter_compositionMode uintptr
	addr_QPainter_font uintptr
	addr_QPainter_setFont uintptr
	addr_QPainter_fontMetrics uintptr
	addr_QPainter_fontInfo uintptr
	addr_QPainter_setPen uintptr
	addr_QPainter_setPen2 uintptr
	addr_QPainter_setPen3 uintptr
	addr_QPainter_pen uintptr
	addr_QPainter_setBrush uintptr
	addr_QPainter_setBrush2 uintptr
	addr_QPainter_setBrush3 uintptr
	addr_QPainter_setBrush4 uintptr
	addr_QPainter_brush uintptr
	addr_QPainter_setBackgroundMode uintptr
	addr_QPainter_backgroundMode uintptr
	addr_QPainter_brushOrigin uintptr
	addr_QPainter_brushOriginF uintptr
	addr_QPainter_setBrushOrigin uintptr
	addr_QPainter_setBrushOrigin2 uintptr
	addr_QPainter_setBrushOrigin3 uintptr
	addr_QPainter_setBackground uintptr
	addr_QPainter_background uintptr
	pg_QPainter_opacity func(uintptr) float64
	pg_QPainter_setOpacity func(uintptr, float64)
	addr_QPainter_clipRegion uintptr
	addr_QPainter_clipPath uintptr
	addr_QPainter_setClipRect uintptr
	addr_QPainter_setClipRect2 uintptr
	addr_QPainter_setClipRect3 uintptr
	addr_QPainter_setClipRegion uintptr
	addr_QPainter_setClipPath uintptr
	addr_QPainter_setClipping uintptr
	addr_QPainter_hasClipping uintptr
	addr_QPainter_clipBoundingRect uintptr
	addr_QPainter_save uintptr
	addr_QPainter_restore uintptr
	addr_QPainter_setTransform uintptr
	addr_QPainter_transform uintptr
	addr_QPainter_deviceTransform uintptr
	addr_QPainter_resetTransform uintptr
	addr_QPainter_setWorldTransform uintptr
	addr_QPainter_worldTransform uintptr
	addr_QPainter_combinedTransform uintptr
	addr_QPainter_setWorldMatrixEnabled uintptr
	addr_QPainter_worldMatrixEnabled uintptr
	pg_QPainter_scale func(uintptr, float64, float64)
	pg_QPainter_shear func(uintptr, float64, float64)
	pg_QPainter_rotate func(uintptr, float64)
	addr_QPainter_translate uintptr
	addr_QPainter_translate2 uintptr
	pg_QPainter_translate3 func(uintptr, float64, float64)
	addr_QPainter_window uintptr
	addr_QPainter_setWindow uintptr
	addr_QPainter_setWindow2 uintptr
	addr_QPainter_viewport uintptr
	addr_QPainter_setViewport uintptr
	addr_QPainter_setViewport2 uintptr
	addr_QPainter_setViewTransformEnabled uintptr
	addr_QPainter_viewTransformEnabled uintptr
	addr_QPainter_strokePath uintptr
	addr_QPainter_fillPath uintptr
	addr_QPainter_drawPath uintptr
	addr_QPainter_drawPoint uintptr
	addr_QPainter_drawPoint2 uintptr
	addr_QPainter_drawPoint3 uintptr
	addr_QPainter_drawPoints uintptr
	addr_QPainter_drawPoints2 uintptr
	addr_QPainter_drawPoints3 uintptr
	addr_QPainter_drawPoints4 uintptr
	addr_QPainter_drawLine uintptr
	addr_QPainter_drawLine2 uintptr
	addr_QPainter_drawLine3 uintptr
	addr_QPainter_drawLine4 uintptr
	addr_QPainter_drawLine5 uintptr
	addr_QPainter_drawLines uintptr
	addr_QPainter_drawLines2 uintptr
	addr_QPainter_drawLines3 uintptr
	addr_QPainter_drawLines4 uintptr
	addr_QPainter_drawLines5 uintptr
	addr_QPainter_drawLines6 uintptr
	addr_QPainter_drawLines7 uintptr
	addr_QPainter_drawLines8 uintptr
	addr_QPainter_drawRect uintptr
	addr_QPainter_drawRect2 uintptr
	addr_QPainter_drawRect3 uintptr
	addr_QPainter_drawRects uintptr
	addr_QPainter_drawRects2 uintptr
	addr_QPainter_drawRects3 uintptr
	addr_QPainter_drawRects4 uintptr
	addr_QPainter_drawEllipse uintptr
	addr_QPainter_drawEllipse2 uintptr
	addr_QPainter_drawEllipse3 uintptr
	pg_QPainter_drawEllipse4 func(uintptr, uintptr, float64, float64)
	addr_QPainter_drawEllipse5 uintptr
	addr_QPainter_drawPolyline uintptr
	addr_QPainter_drawPolyline2 uintptr
	addr_QPainter_drawPolyline3 uintptr
	addr_QPainter_drawPolyline4 uintptr
	addr_QPainter_drawPolygon uintptr
	addr_QPainter_drawPolygon2 uintptr
	addr_QPainter_drawPolygon3 uintptr
	addr_QPainter_drawPolygon4 uintptr
	addr_QPainter_drawConvexPolygon uintptr
	addr_QPainter_drawConvexPolygon2 uintptr
	addr_QPainter_drawConvexPolygon3 uintptr
	addr_QPainter_drawConvexPolygon4 uintptr
	addr_QPainter_drawArc uintptr
	addr_QPainter_drawArc2 uintptr
	addr_QPainter_drawArc3 uintptr
	addr_QPainter_drawPie uintptr
	addr_QPainter_drawPie2 uintptr
	addr_QPainter_drawPie3 uintptr
	addr_QPainter_drawChord uintptr
	addr_QPainter_drawChord2 uintptr
	addr_QPainter_drawChord3 uintptr
	pg_QPainter_drawRoundedRect func(uintptr, uintptr, float64, float64)
	pg_QPainter_drawRoundedRect2 func(uintptr, int32, int32, int32, int32, float64, float64)
	pg_QPainter_drawRoundedRect3 func(uintptr, uintptr, float64, float64)
	addr_QPainter_drawTiledPixmap uintptr
	addr_QPainter_drawTiledPixmap2 uintptr
	addr_QPainter_drawTiledPixmap3 uintptr
	addr_QPainter_drawPicture uintptr
	addr_QPainter_drawPicture2 uintptr
	addr_QPainter_drawPicture3 uintptr
	addr_QPainter_drawPixmap uintptr
	addr_QPainter_drawPixmap2 uintptr
	addr_QPainter_drawPixmap3 uintptr
	addr_QPainter_drawPixmap4 uintptr
	addr_QPainter_drawPixmap5 uintptr
	addr_QPainter_drawPixmap6 uintptr
	addr_QPainter_drawPixmap7 uintptr
	addr_QPainter_drawPixmap8 uintptr
	addr_QPainter_drawPixmap9 uintptr
	addr_QPainter_drawPixmap10 uintptr
	addr_QPainter_drawPixmap11 uintptr
	addr_QPainter_drawPixmapFragments uintptr
	addr_QPainter_drawImage uintptr
	addr_QPainter_drawImage2 uintptr
	addr_QPainter_drawImage3 uintptr
	addr_QPainter_drawImage4 uintptr
	addr_QPainter_drawImage5 uintptr
	addr_QPainter_drawImage6 uintptr
	addr_QPainter_drawImage7 uintptr
	addr_QPainter_drawImage8 uintptr
	addr_QPainter_drawImage9 uintptr
	addr_QPainter_setLayoutDirection uintptr
	addr_QPainter_layoutDirection uintptr
	addr_QPainter_drawGlyphRun uintptr
	addr_QPainter_drawStaticText uintptr
	addr_QPainter_drawStaticText2 uintptr
	addr_QPainter_drawStaticText3 uintptr
	addr_QPainter_drawText uintptr
	addr_QPainter_drawText2 uintptr
	addr_QPainter_drawText3 uintptr
	addr_QPainter_drawText4 uintptr
	addr_QPainter_drawText5 uintptr
	addr_QPainter_drawText6 uintptr
	addr_QPainter_drawText7 uintptr
	addr_QPainter_drawText8 uintptr
	addr_QPainter_boundingRect uintptr
	addr_QPainter_boundingRect2 uintptr
	addr_QPainter_boundingRect3 uintptr
	addr_QPainter_boundingRect4 uintptr
	addr_QPainter_drawTextItem uintptr
	addr_QPainter_drawTextItem2 uintptr
	addr_QPainter_drawTextItem3 uintptr
	addr_QPainter_fillRect uintptr
	addr_QPainter_fillRect2 uintptr
	addr_QPainter_fillRect3 uintptr
	addr_QPainter_fillRect4 uintptr
	addr_QPainter_fillRect5 uintptr
	addr_QPainter_fillRect6 uintptr
	addr_QPainter_fillRect7 uintptr
	addr_QPainter_fillRect8 uintptr
	addr_QPainter_fillRect9 uintptr
	addr_QPainter_fillRect10 uintptr
	addr_QPainter_fillRect11 uintptr
	addr_QPainter_fillRect12 uintptr
	addr_QPainter_fillRect13 uintptr
	addr_QPainter_fillRect14 uintptr
	addr_QPainter_fillRect15 uintptr
	addr_QPainter_eraseRect uintptr
	addr_QPainter_eraseRect2 uintptr
	addr_QPainter_eraseRect3 uintptr
	addr_QPainter_setRenderHint uintptr
	addr_QPainter_setRenderHints uintptr
	addr_QPainter_renderHints uintptr
	addr_QPainter_testRenderHint uintptr
	addr_QPainter_paintEngine uintptr
	addr_QPainter_beginNativePainting uintptr
	addr_QPainter_endNativePainting uintptr
	addr_QPainter_setClipRect4 uintptr
	addr_QPainter_setClipRect5 uintptr
	addr_QPainter_setClipRect6 uintptr
	addr_QPainter_setClipRegion2 uintptr
	addr_QPainter_setClipPath2 uintptr
	addr_QPainter_setTransform2 uintptr
	addr_QPainter_setWorldTransform2 uintptr
	addr_QPainter_drawPolygon5 uintptr
	addr_QPainter_drawPolygon6 uintptr
	addr_QPainter_drawPolygon7 uintptr
	addr_QPainter_drawPolygon8 uintptr
	pg_QPainter_drawRoundedRect4 func(uintptr, uintptr, float64, float64, int32)
	pg_QPainter_drawRoundedRect5 func(uintptr, int32, int32, int32, int32, float64, float64, int32)
	pg_QPainter_drawRoundedRect6 func(uintptr, uintptr, float64, float64, int32)
	addr_QPainter_drawTiledPixmap4 uintptr
	addr_QPainter_drawTiledPixmap5 uintptr
	addr_QPainter_drawTiledPixmap6 uintptr
	addr_QPainter_drawTiledPixmap7 uintptr
	addr_QPainter_drawPixmapFragments2 uintptr
	addr_QPainter_drawImage10 uintptr
	addr_QPainter_drawImage11 uintptr
	addr_QPainter_drawImage12 uintptr
	addr_QPainter_drawImage13 uintptr
	addr_QPainter_drawImage14 uintptr
	addr_QPainter_drawImage15 uintptr
	addr_QPainter_drawImage16 uintptr
	addr_QPainter_drawImage17 uintptr
	addr_QPainter_drawImage18 uintptr
	addr_QPainter_drawText9 uintptr
	addr_QPainter_drawText10 uintptr
	addr_QPainter_drawText11 uintptr
	addr_QPainter_drawText12 uintptr
	addr_QPainter_boundingRect5 uintptr
	addr_QPainter_setRenderHint2 uintptr
	addr_QPainter_setRenderHints2 uintptr
	pg_QPainter_delete func(uintptr)
)

func init() {
	symaddr(&addr_QPainter_new, "QPainter_new")
	symaddr(&addr_QPainter_new2, "QPainter_new2")
	symaddr(&addr_QPainter_device, "QPainter_device")
	symaddr(&addr_QPainter_begin, "QPainter_begin")
	symaddr(&addr_QPainter_end, "QPainter_end")
	symaddr(&addr_QPainter_isActive, "QPainter_isActive")
	symaddr(&addr_QPainter_setCompositionMode, "QPainter_setCompositionMode")
	symaddr(&addr_QPainter_compositionMode, "QPainter_compositionMode")
	symaddr(&addr_QPainter_font, "QPainter_font")
	symaddr(&addr_QPainter_setFont, "QPainter_setFont")
	symaddr(&addr_QPainter_fontMetrics, "QPainter_fontMetrics")
	symaddr(&addr_QPainter_fontInfo, "QPainter_fontInfo")
	symaddr(&addr_QPainter_setPen, "QPainter_setPen")
	symaddr(&addr_QPainter_setPen2, "QPainter_setPen2")
	symaddr(&addr_QPainter_setPen3, "QPainter_setPen3")
	symaddr(&addr_QPainter_pen, "QPainter_pen")
	symaddr(&addr_QPainter_setBrush, "QPainter_setBrush")
	symaddr(&addr_QPainter_setBrush2, "QPainter_setBrush2")
	symaddr(&addr_QPainter_setBrush3, "QPainter_setBrush3")
	symaddr(&addr_QPainter_setBrush4, "QPainter_setBrush4")
	symaddr(&addr_QPainter_brush, "QPainter_brush")
	symaddr(&addr_QPainter_setBackgroundMode, "QPainter_setBackgroundMode")
	symaddr(&addr_QPainter_backgroundMode, "QPainter_backgroundMode")
	symaddr(&addr_QPainter_brushOrigin, "QPainter_brushOrigin")
	symaddr(&addr_QPainter_brushOriginF, "QPainter_brushOriginF")
	symaddr(&addr_QPainter_setBrushOrigin, "QPainter_setBrushOrigin")
	symaddr(&addr_QPainter_setBrushOrigin2, "QPainter_setBrushOrigin2")
	symaddr(&addr_QPainter_setBrushOrigin3, "QPainter_setBrushOrigin3")
	symaddr(&addr_QPainter_setBackground, "QPainter_setBackground")
	symaddr(&addr_QPainter_background, "QPainter_background")
	reg(&pg_QPainter_opacity, "QPainter_opacity")
	reg(&pg_QPainter_setOpacity, "QPainter_setOpacity")
	symaddr(&addr_QPainter_clipRegion, "QPainter_clipRegion")
	symaddr(&addr_QPainter_clipPath, "QPainter_clipPath")
	symaddr(&addr_QPainter_setClipRect, "QPainter_setClipRect")
	symaddr(&addr_QPainter_setClipRect2, "QPainter_setClipRect2")
	symaddr(&addr_QPainter_setClipRect3, "QPainter_setClipRect3")
	symaddr(&addr_QPainter_setClipRegion, "QPainter_setClipRegion")
	symaddr(&addr_QPainter_setClipPath, "QPainter_setClipPath")
	symaddr(&addr_QPainter_setClipping, "QPainter_setClipping")
	symaddr(&addr_QPainter_hasClipping, "QPainter_hasClipping")
	symaddr(&addr_QPainter_clipBoundingRect, "QPainter_clipBoundingRect")
	symaddr(&addr_QPainter_save, "QPainter_save")
	symaddr(&addr_QPainter_restore, "QPainter_restore")
	symaddr(&addr_QPainter_setTransform, "QPainter_setTransform")
	symaddr(&addr_QPainter_transform, "QPainter_transform")
	symaddr(&addr_QPainter_deviceTransform, "QPainter_deviceTransform")
	symaddr(&addr_QPainter_resetTransform, "QPainter_resetTransform")
	symaddr(&addr_QPainter_setWorldTransform, "QPainter_setWorldTransform")
	symaddr(&addr_QPainter_worldTransform, "QPainter_worldTransform")
	symaddr(&addr_QPainter_combinedTransform, "QPainter_combinedTransform")
	symaddr(&addr_QPainter_setWorldMatrixEnabled, "QPainter_setWorldMatrixEnabled")
	symaddr(&addr_QPainter_worldMatrixEnabled, "QPainter_worldMatrixEnabled")
	reg(&pg_QPainter_scale, "QPainter_scale")
	reg(&pg_QPainter_shear, "QPainter_shear")
	reg(&pg_QPainter_rotate, "QPainter_rotate")
	symaddr(&addr_QPainter_translate, "QPainter_translate")
	symaddr(&addr_QPainter_translate2, "QPainter_translate2")
	reg(&pg_QPainter_translate3, "QPainter_translate3")
	symaddr(&addr_QPainter_window, "QPainter_window")
	symaddr(&addr_QPainter_setWindow, "QPainter_setWindow")
	symaddr(&addr_QPainter_setWindow2, "QPainter_setWindow2")
	symaddr(&addr_QPainter_viewport, "QPainter_viewport")
	symaddr(&addr_QPainter_setViewport, "QPainter_setViewport")
	symaddr(&addr_QPainter_setViewport2, "QPainter_setViewport2")
	symaddr(&addr_QPainter_setViewTransformEnabled, "QPainter_setViewTransformEnabled")
	symaddr(&addr_QPainter_viewTransformEnabled, "QPainter_viewTransformEnabled")
	symaddr(&addr_QPainter_strokePath, "QPainter_strokePath")
	symaddr(&addr_QPainter_fillPath, "QPainter_fillPath")
	symaddr(&addr_QPainter_drawPath, "QPainter_drawPath")
	symaddr(&addr_QPainter_drawPoint, "QPainter_drawPoint")
	symaddr(&addr_QPainter_drawPoint2, "QPainter_drawPoint2")
	symaddr(&addr_QPainter_drawPoint3, "QPainter_drawPoint3")
	symaddr(&addr_QPainter_drawPoints, "QPainter_drawPoints")
	symaddr(&addr_QPainter_drawPoints2, "QPainter_drawPoints2")
	symaddr(&addr_QPainter_drawPoints3, "QPainter_drawPoints3")
	symaddr(&addr_QPainter_drawPoints4, "QPainter_drawPoints4")
	symaddr(&addr_QPainter_drawLine, "QPainter_drawLine")
	symaddr(&addr_QPainter_drawLine2, "QPainter_drawLine2")
	symaddr(&addr_QPainter_drawLine3, "QPainter_drawLine3")
	symaddr(&addr_QPainter_drawLine4, "QPainter_drawLine4")
	symaddr(&addr_QPainter_drawLine5, "QPainter_drawLine5")
	symaddr(&addr_QPainter_drawLines, "QPainter_drawLines")
	symaddr(&addr_QPainter_drawLines2, "QPainter_drawLines2")
	symaddr(&addr_QPainter_drawLines3, "QPainter_drawLines3")
	symaddr(&addr_QPainter_drawLines4, "QPainter_drawLines4")
	symaddr(&addr_QPainter_drawLines5, "QPainter_drawLines5")
	symaddr(&addr_QPainter_drawLines6, "QPainter_drawLines6")
	symaddr(&addr_QPainter_drawLines7, "QPainter_drawLines7")
	symaddr(&addr_QPainter_drawLines8, "QPainter_drawLines8")
	symaddr(&addr_QPainter_drawRect, "QPainter_drawRect")
	symaddr(&addr_QPainter_drawRect2, "QPainter_drawRect2")
	symaddr(&addr_QPainter_drawRect3, "QPainter_drawRect3")
	symaddr(&addr_QPainter_drawRects, "QPainter_drawRects")
	symaddr(&addr_QPainter_drawRects2, "QPainter_drawRects2")
	symaddr(&addr_QPainter_drawRects3, "QPainter_drawRects3")
	symaddr(&addr_QPainter_drawRects4, "QPainter_drawRects4")
	symaddr(&addr_QPainter_drawEllipse, "QPainter_drawEllipse")
	symaddr(&addr_QPainter_drawEllipse2, "QPainter_drawEllipse2")
	symaddr(&addr_QPainter_drawEllipse3, "QPainter_drawEllipse3")
	reg(&pg_QPainter_drawEllipse4, "QPainter_drawEllipse4")
	symaddr(&addr_QPainter_drawEllipse5, "QPainter_drawEllipse5")
	symaddr(&addr_QPainter_drawPolyline, "QPainter_drawPolyline")
	symaddr(&addr_QPainter_drawPolyline2, "QPainter_drawPolyline2")
	symaddr(&addr_QPainter_drawPolyline3, "QPainter_drawPolyline3")
	symaddr(&addr_QPainter_drawPolyline4, "QPainter_drawPolyline4")
	symaddr(&addr_QPainter_drawPolygon, "QPainter_drawPolygon")
	symaddr(&addr_QPainter_drawPolygon2, "QPainter_drawPolygon2")
	symaddr(&addr_QPainter_drawPolygon3, "QPainter_drawPolygon3")
	symaddr(&addr_QPainter_drawPolygon4, "QPainter_drawPolygon4")
	symaddr(&addr_QPainter_drawConvexPolygon, "QPainter_drawConvexPolygon")
	symaddr(&addr_QPainter_drawConvexPolygon2, "QPainter_drawConvexPolygon2")
	symaddr(&addr_QPainter_drawConvexPolygon3, "QPainter_drawConvexPolygon3")
	symaddr(&addr_QPainter_drawConvexPolygon4, "QPainter_drawConvexPolygon4")
	symaddr(&addr_QPainter_drawArc, "QPainter_drawArc")
	symaddr(&addr_QPainter_drawArc2, "QPainter_drawArc2")
	symaddr(&addr_QPainter_drawArc3, "QPainter_drawArc3")
	symaddr(&addr_QPainter_drawPie, "QPainter_drawPie")
	symaddr(&addr_QPainter_drawPie2, "QPainter_drawPie2")
	symaddr(&addr_QPainter_drawPie3, "QPainter_drawPie3")
	symaddr(&addr_QPainter_drawChord, "QPainter_drawChord")
	symaddr(&addr_QPainter_drawChord2, "QPainter_drawChord2")
	symaddr(&addr_QPainter_drawChord3, "QPainter_drawChord3")
	reg(&pg_QPainter_drawRoundedRect, "QPainter_drawRoundedRect")
	reg(&pg_QPainter_drawRoundedRect2, "QPainter_drawRoundedRect2")
	reg(&pg_QPainter_drawRoundedRect3, "QPainter_drawRoundedRect3")
	symaddr(&addr_QPainter_drawTiledPixmap, "QPainter_drawTiledPixmap")
	symaddr(&addr_QPainter_drawTiledPixmap2, "QPainter_drawTiledPixmap2")
	symaddr(&addr_QPainter_drawTiledPixmap3, "QPainter_drawTiledPixmap3")
	symaddr(&addr_QPainter_drawPicture, "QPainter_drawPicture")
	symaddr(&addr_QPainter_drawPicture2, "QPainter_drawPicture2")
	symaddr(&addr_QPainter_drawPicture3, "QPainter_drawPicture3")
	symaddr(&addr_QPainter_drawPixmap, "QPainter_drawPixmap")
	symaddr(&addr_QPainter_drawPixmap2, "QPainter_drawPixmap2")
	symaddr(&addr_QPainter_drawPixmap3, "QPainter_drawPixmap3")
	symaddr(&addr_QPainter_drawPixmap4, "QPainter_drawPixmap4")
	symaddr(&addr_QPainter_drawPixmap5, "QPainter_drawPixmap5")
	symaddr(&addr_QPainter_drawPixmap6, "QPainter_drawPixmap6")
	symaddr(&addr_QPainter_drawPixmap7, "QPainter_drawPixmap7")
	symaddr(&addr_QPainter_drawPixmap8, "QPainter_drawPixmap8")
	symaddr(&addr_QPainter_drawPixmap9, "QPainter_drawPixmap9")
	symaddr(&addr_QPainter_drawPixmap10, "QPainter_drawPixmap10")
	symaddr(&addr_QPainter_drawPixmap11, "QPainter_drawPixmap11")
	symaddr(&addr_QPainter_drawPixmapFragments, "QPainter_drawPixmapFragments")
	symaddr(&addr_QPainter_drawImage, "QPainter_drawImage")
	symaddr(&addr_QPainter_drawImage2, "QPainter_drawImage2")
	symaddr(&addr_QPainter_drawImage3, "QPainter_drawImage3")
	symaddr(&addr_QPainter_drawImage4, "QPainter_drawImage4")
	symaddr(&addr_QPainter_drawImage5, "QPainter_drawImage5")
	symaddr(&addr_QPainter_drawImage6, "QPainter_drawImage6")
	symaddr(&addr_QPainter_drawImage7, "QPainter_drawImage7")
	symaddr(&addr_QPainter_drawImage8, "QPainter_drawImage8")
	symaddr(&addr_QPainter_drawImage9, "QPainter_drawImage9")
	symaddr(&addr_QPainter_setLayoutDirection, "QPainter_setLayoutDirection")
	symaddr(&addr_QPainter_layoutDirection, "QPainter_layoutDirection")
	symaddr(&addr_QPainter_drawGlyphRun, "QPainter_drawGlyphRun")
	symaddr(&addr_QPainter_drawStaticText, "QPainter_drawStaticText")
	symaddr(&addr_QPainter_drawStaticText2, "QPainter_drawStaticText2")
	symaddr(&addr_QPainter_drawStaticText3, "QPainter_drawStaticText3")
	symaddr(&addr_QPainter_drawText, "QPainter_drawText")
	symaddr(&addr_QPainter_drawText2, "QPainter_drawText2")
	symaddr(&addr_QPainter_drawText3, "QPainter_drawText3")
	symaddr(&addr_QPainter_drawText4, "QPainter_drawText4")
	symaddr(&addr_QPainter_drawText5, "QPainter_drawText5")
	symaddr(&addr_QPainter_drawText6, "QPainter_drawText6")
	symaddr(&addr_QPainter_drawText7, "QPainter_drawText7")
	symaddr(&addr_QPainter_drawText8, "QPainter_drawText8")
	symaddr(&addr_QPainter_boundingRect, "QPainter_boundingRect")
	symaddr(&addr_QPainter_boundingRect2, "QPainter_boundingRect2")
	symaddr(&addr_QPainter_boundingRect3, "QPainter_boundingRect3")
	symaddr(&addr_QPainter_boundingRect4, "QPainter_boundingRect4")
	symaddr(&addr_QPainter_drawTextItem, "QPainter_drawTextItem")
	symaddr(&addr_QPainter_drawTextItem2, "QPainter_drawTextItem2")
	symaddr(&addr_QPainter_drawTextItem3, "QPainter_drawTextItem3")
	symaddr(&addr_QPainter_fillRect, "QPainter_fillRect")
	symaddr(&addr_QPainter_fillRect2, "QPainter_fillRect2")
	symaddr(&addr_QPainter_fillRect3, "QPainter_fillRect3")
	symaddr(&addr_QPainter_fillRect4, "QPainter_fillRect4")
	symaddr(&addr_QPainter_fillRect5, "QPainter_fillRect5")
	symaddr(&addr_QPainter_fillRect6, "QPainter_fillRect6")
	symaddr(&addr_QPainter_fillRect7, "QPainter_fillRect7")
	symaddr(&addr_QPainter_fillRect8, "QPainter_fillRect8")
	symaddr(&addr_QPainter_fillRect9, "QPainter_fillRect9")
	symaddr(&addr_QPainter_fillRect10, "QPainter_fillRect10")
	symaddr(&addr_QPainter_fillRect11, "QPainter_fillRect11")
	symaddr(&addr_QPainter_fillRect12, "QPainter_fillRect12")
	symaddr(&addr_QPainter_fillRect13, "QPainter_fillRect13")
	symaddr(&addr_QPainter_fillRect14, "QPainter_fillRect14")
	symaddr(&addr_QPainter_fillRect15, "QPainter_fillRect15")
	symaddr(&addr_QPainter_eraseRect, "QPainter_eraseRect")
	symaddr(&addr_QPainter_eraseRect2, "QPainter_eraseRect2")
	symaddr(&addr_QPainter_eraseRect3, "QPainter_eraseRect3")
	symaddr(&addr_QPainter_setRenderHint, "QPainter_setRenderHint")
	symaddr(&addr_QPainter_setRenderHints, "QPainter_setRenderHints")
	symaddr(&addr_QPainter_renderHints, "QPainter_renderHints")
	symaddr(&addr_QPainter_testRenderHint, "QPainter_testRenderHint")
	symaddr(&addr_QPainter_paintEngine, "QPainter_paintEngine")
	symaddr(&addr_QPainter_beginNativePainting, "QPainter_beginNativePainting")
	symaddr(&addr_QPainter_endNativePainting, "QPainter_endNativePainting")
	symaddr(&addr_QPainter_setClipRect4, "QPainter_setClipRect4")
	symaddr(&addr_QPainter_setClipRect5, "QPainter_setClipRect5")
	symaddr(&addr_QPainter_setClipRect6, "QPainter_setClipRect6")
	symaddr(&addr_QPainter_setClipRegion2, "QPainter_setClipRegion2")
	symaddr(&addr_QPainter_setClipPath2, "QPainter_setClipPath2")
	symaddr(&addr_QPainter_setTransform2, "QPainter_setTransform2")
	symaddr(&addr_QPainter_setWorldTransform2, "QPainter_setWorldTransform2")
	symaddr(&addr_QPainter_drawPolygon5, "QPainter_drawPolygon5")
	symaddr(&addr_QPainter_drawPolygon6, "QPainter_drawPolygon6")
	symaddr(&addr_QPainter_drawPolygon7, "QPainter_drawPolygon7")
	symaddr(&addr_QPainter_drawPolygon8, "QPainter_drawPolygon8")
	reg(&pg_QPainter_drawRoundedRect4, "QPainter_drawRoundedRect4")
	reg(&pg_QPainter_drawRoundedRect5, "QPainter_drawRoundedRect5")
	reg(&pg_QPainter_drawRoundedRect6, "QPainter_drawRoundedRect6")
	symaddr(&addr_QPainter_drawTiledPixmap4, "QPainter_drawTiledPixmap4")
	symaddr(&addr_QPainter_drawTiledPixmap5, "QPainter_drawTiledPixmap5")
	symaddr(&addr_QPainter_drawTiledPixmap6, "QPainter_drawTiledPixmap6")
	symaddr(&addr_QPainter_drawTiledPixmap7, "QPainter_drawTiledPixmap7")
	symaddr(&addr_QPainter_drawPixmapFragments2, "QPainter_drawPixmapFragments2")
	symaddr(&addr_QPainter_drawImage10, "QPainter_drawImage10")
	symaddr(&addr_QPainter_drawImage11, "QPainter_drawImage11")
	symaddr(&addr_QPainter_drawImage12, "QPainter_drawImage12")
	symaddr(&addr_QPainter_drawImage13, "QPainter_drawImage13")
	symaddr(&addr_QPainter_drawImage14, "QPainter_drawImage14")
	symaddr(&addr_QPainter_drawImage15, "QPainter_drawImage15")
	symaddr(&addr_QPainter_drawImage16, "QPainter_drawImage16")
	symaddr(&addr_QPainter_drawImage17, "QPainter_drawImage17")
	symaddr(&addr_QPainter_drawImage18, "QPainter_drawImage18")
	symaddr(&addr_QPainter_drawText9, "QPainter_drawText9")
	symaddr(&addr_QPainter_drawText10, "QPainter_drawText10")
	symaddr(&addr_QPainter_drawText11, "QPainter_drawText11")
	symaddr(&addr_QPainter_drawText12, "QPainter_drawText12")
	symaddr(&addr_QPainter_boundingRect5, "QPainter_boundingRect5")
	symaddr(&addr_QPainter_setRenderHint2, "QPainter_setRenderHint2")
	symaddr(&addr_QPainter_setRenderHints2, "QPainter_setRenderHints2")
	reg(&pg_QPainter_delete, "QPainter_delete")
}

type QPainter struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QPainter) UnsafeGoID() int64     { return this.vid }
func (this *QPainter) UnsafeSetGoID(id int64) { this.vid = id }

func bindQPainter(h uintptr) *QPainter {
	if h == 0 {
		return nil
	}
	return &QPainter{h: h}
}

func newownedQPainter(h uintptr) *QPainter {
	o := bindQPainter(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QPainter).Delete)
	}
	return o
}

func (this *QPainter) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QPainter) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQPainter(h unsafe.Pointer) *QPainter { return bindQPainter(uintptr(h)) }
func UnsafeNewQPainterOwned(h unsafe.Pointer) *QPainter { return newownedQPainter(uintptr(h)) }

func (this *QPainter) UnsafeRelease() { this.h = 0 }

func NewQPainter() *QPainter {
	__r, _, _ := purego.SyscallN(addr_QPainter_new)
	__o := newownedQPainter(__r)
	return __o
}

func NewQPainter2(arg1 *QPaintDevice) *QPainter {
	__r, _, _ := purego.SyscallN(addr_QPainter_new2, uintptr(arg1.UnsafePointer()))
	__o := newownedQPainter(__r)
	runtime.KeepAlive(arg1)
	return __o
}

func (this *QPainter) Device() *QPaintDevice {
	__r, _, _ := purego.SyscallN(addr_QPainter_device, this.h)
	runtime.KeepAlive(this)
	return bindQPaintDevice(__r)
}

func (this *QPainter) Begin(arg1 *QPaintDevice) bool {
	__r0, _, _ := purego.SyscallN(addr_QPainter_begin, this.h, uintptr(arg1.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return __r != 0
}

func (this *QPainter) End() bool {
	__r0, _, _ := purego.SyscallN(addr_QPainter_end, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QPainter) IsActive() bool {
	__r0, _, _ := purego.SyscallN(addr_QPainter_isActive, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QPainter) SetCompositionMode(mode QPainter__CompositionMode) {
	purego.SyscallN(addr_QPainter_setCompositionMode, this.h, uintptr(mode))
	runtime.KeepAlive(this)
}

func (this *QPainter) CompositionMode() QPainter__CompositionMode {
	__r0, _, _ := purego.SyscallN(addr_QPainter_compositionMode, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QPainter__CompositionMode(__r)
}

func (this *QPainter) Font() *QFont {
	__r, _, _ := purego.SyscallN(addr_QPainter_font, this.h)
	runtime.KeepAlive(this)
	return bindQFont(__r)
}

func (this *QPainter) SetFont(f *QFont) {
	purego.SyscallN(addr_QPainter_setFont, this.h, uintptr(f.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(f)
}

func (this *QPainter) FontMetrics() *QFontMetrics {
	__r, _, _ := purego.SyscallN(addr_QPainter_fontMetrics, this.h)
	runtime.KeepAlive(this)
	return newownedQFontMetrics(__r)
}

func (this *QPainter) FontInfo() *QFontInfo {
	__r, _, _ := purego.SyscallN(addr_QPainter_fontInfo, this.h)
	runtime.KeepAlive(this)
	return newownedQFontInfo(__r)
}

func (this *QPainter) SetPen(color *QColor) {
	purego.SyscallN(addr_QPainter_setPen, this.h, uintptr(color.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(color)
}

func (this *QPainter) SetPenWithPen(pen *QPen) {
	purego.SyscallN(addr_QPainter_setPen2, this.h, uintptr(pen.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pen)
}

func (this *QPainter) SetPenWithStyle(style PenStyle) {
	purego.SyscallN(addr_QPainter_setPen3, this.h, uintptr(style))
	runtime.KeepAlive(this)
}

func (this *QPainter) Pen() *QPen {
	__r, _, _ := purego.SyscallN(addr_QPainter_pen, this.h)
	runtime.KeepAlive(this)
	return bindQPen(__r)
}

func (this *QPainter) SetBrush(brush *QBrush) {
	purego.SyscallN(addr_QPainter_setBrush, this.h, uintptr(brush.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(brush)
}

func (this *QPainter) SetBrushWithStyle(style BrushStyle) {
	purego.SyscallN(addr_QPainter_setBrush2, this.h, uintptr(style))
	runtime.KeepAlive(this)
}

func (this *QPainter) SetBrushWithColor(color *QColor) {
	purego.SyscallN(addr_QPainter_setBrush3, this.h, uintptr(color.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(color)
}

func (this *QPainter) SetBrush2(color GlobalColor) {
	purego.SyscallN(addr_QPainter_setBrush4, this.h, uintptr(color))
	runtime.KeepAlive(this)
}

func (this *QPainter) Brush() *QBrush {
	__r, _, _ := purego.SyscallN(addr_QPainter_brush, this.h)
	runtime.KeepAlive(this)
	return bindQBrush(__r)
}

func (this *QPainter) SetBackgroundMode(mode BGMode) {
	purego.SyscallN(addr_QPainter_setBackgroundMode, this.h, uintptr(mode))
	runtime.KeepAlive(this)
}

func (this *QPainter) BackgroundMode() BGMode {
	__r0, _, _ := purego.SyscallN(addr_QPainter_backgroundMode, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return BGMode(__r)
}

func (this *QPainter) BrushOrigin() *QPoint {
	__r, _, _ := purego.SyscallN(addr_QPainter_brushOrigin, this.h)
	runtime.KeepAlive(this)
	return newownedQPoint(__r)
}

func (this *QPainter) BrushOriginF() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QPainter_brushOriginF, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QPainter) SetBrushOrigin(x int, y int) {
	purego.SyscallN(addr_QPainter_setBrushOrigin, this.h, uintptr(x), uintptr(y))
	runtime.KeepAlive(this)
}

func (this *QPainter) SetBrushOriginWithQPoint(arg1 *QPoint) {
	purego.SyscallN(addr_QPainter_setBrushOrigin2, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QPainter) SetBrushOriginWithQPointF(arg1 *QPointF) {
	purego.SyscallN(addr_QPainter_setBrushOrigin3, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QPainter) SetBackground(bg *QBrush) {
	purego.SyscallN(addr_QPainter_setBackground, this.h, uintptr(bg.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(bg)
}

func (this *QPainter) Background() *QBrush {
	__r, _, _ := purego.SyscallN(addr_QPainter_background, this.h)
	runtime.KeepAlive(this)
	return bindQBrush(__r)
}

func (this *QPainter) Opacity() float64 {
	__r := pg_QPainter_opacity(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QPainter) SetOpacity(opacity float64) {
	pg_QPainter_setOpacity(this.h, opacity)
	runtime.KeepAlive(this)
}

func (this *QPainter) ClipRegion() *QRegion {
	__r, _, _ := purego.SyscallN(addr_QPainter_clipRegion, this.h)
	runtime.KeepAlive(this)
	return newownedQRegion(__r)
}

func (this *QPainter) ClipPath() *QPainterPath {
	__r, _, _ := purego.SyscallN(addr_QPainter_clipPath, this.h)
	runtime.KeepAlive(this)
	return newownedQPainterPath(__r)
}

func (this *QPainter) SetClipRect(arg1 *QRectF) {
	purego.SyscallN(addr_QPainter_setClipRect, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QPainter) SetClipRectWithQRect(arg1 *QRect) {
	purego.SyscallN(addr_QPainter_setClipRect2, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QPainter) SetClipRect2(x int, y int, w int, h int) {
	purego.SyscallN(addr_QPainter_setClipRect3, this.h, uintptr(x), uintptr(y), uintptr(w), uintptr(h))
	runtime.KeepAlive(this)
}

func (this *QPainter) SetClipRegion(arg1 *QRegion) {
	purego.SyscallN(addr_QPainter_setClipRegion, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QPainter) SetClipPath(path *QPainterPath) {
	purego.SyscallN(addr_QPainter_setClipPath, this.h, uintptr(path.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(path)
}

func (this *QPainter) SetClipping(enable bool) {
	purego.SyscallN(addr_QPainter_setClipping, this.h, uintptr(b2i(enable)))
	runtime.KeepAlive(this)
}

func (this *QPainter) HasClipping() bool {
	__r0, _, _ := purego.SyscallN(addr_QPainter_hasClipping, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QPainter) ClipBoundingRect() *QRectF {
	__r, _, _ := purego.SyscallN(addr_QPainter_clipBoundingRect, this.h)
	runtime.KeepAlive(this)
	return newownedQRectF(__r)
}

func (this *QPainter) Save() {
	purego.SyscallN(addr_QPainter_save, this.h)
	runtime.KeepAlive(this)
}

func (this *QPainter) Restore() {
	purego.SyscallN(addr_QPainter_restore, this.h)
	runtime.KeepAlive(this)
}

func (this *QPainter) SetTransform(transform *QTransform) {
	purego.SyscallN(addr_QPainter_setTransform, this.h, uintptr(transform.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(transform)
}

func (this *QPainter) Transform() *QTransform {
	__r, _, _ := purego.SyscallN(addr_QPainter_transform, this.h)
	runtime.KeepAlive(this)
	return bindQTransform(__r)
}

func (this *QPainter) DeviceTransform() *QTransform {
	__r, _, _ := purego.SyscallN(addr_QPainter_deviceTransform, this.h)
	runtime.KeepAlive(this)
	return bindQTransform(__r)
}

func (this *QPainter) ResetTransform() {
	purego.SyscallN(addr_QPainter_resetTransform, this.h)
	runtime.KeepAlive(this)
}

func (this *QPainter) SetWorldTransform(matrix *QTransform) {
	purego.SyscallN(addr_QPainter_setWorldTransform, this.h, uintptr(matrix.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(matrix)
}

func (this *QPainter) WorldTransform() *QTransform {
	__r, _, _ := purego.SyscallN(addr_QPainter_worldTransform, this.h)
	runtime.KeepAlive(this)
	return bindQTransform(__r)
}

func (this *QPainter) CombinedTransform() *QTransform {
	__r, _, _ := purego.SyscallN(addr_QPainter_combinedTransform, this.h)
	runtime.KeepAlive(this)
	return newownedQTransform(__r)
}

func (this *QPainter) SetWorldMatrixEnabled(enabled bool) {
	purego.SyscallN(addr_QPainter_setWorldMatrixEnabled, this.h, uintptr(b2i(enabled)))
	runtime.KeepAlive(this)
}

func (this *QPainter) WorldMatrixEnabled() bool {
	__r0, _, _ := purego.SyscallN(addr_QPainter_worldMatrixEnabled, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QPainter) Scale(sx float64, sy float64) {
	pg_QPainter_scale(this.h, sx, sy)
	runtime.KeepAlive(this)
}

func (this *QPainter) Shear(sh float64, sv float64) {
	pg_QPainter_shear(this.h, sh, sv)
	runtime.KeepAlive(this)
}

func (this *QPainter) Rotate(a float64) {
	pg_QPainter_rotate(this.h, a)
	runtime.KeepAlive(this)
}

func (this *QPainter) Translate(offset *QPointF) {
	purego.SyscallN(addr_QPainter_translate, this.h, uintptr(offset.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(offset)
}

func (this *QPainter) TranslateWithOffset(offset *QPoint) {
	purego.SyscallN(addr_QPainter_translate2, this.h, uintptr(offset.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(offset)
}

func (this *QPainter) Translate2(dx float64, dy float64) {
	pg_QPainter_translate3(this.h, dx, dy)
	runtime.KeepAlive(this)
}

func (this *QPainter) Window() *QRect {
	__r, _, _ := purego.SyscallN(addr_QPainter_window, this.h)
	runtime.KeepAlive(this)
	return newownedQRect(__r)
}

func (this *QPainter) SetWindow(window *QRect) {
	purego.SyscallN(addr_QPainter_setWindow, this.h, uintptr(window.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(window)
}

func (this *QPainter) SetWindow2(x int, y int, w int, h int) {
	purego.SyscallN(addr_QPainter_setWindow2, this.h, uintptr(x), uintptr(y), uintptr(w), uintptr(h))
	runtime.KeepAlive(this)
}

func (this *QPainter) Viewport() *QRect {
	__r, _, _ := purego.SyscallN(addr_QPainter_viewport, this.h)
	runtime.KeepAlive(this)
	return newownedQRect(__r)
}

func (this *QPainter) SetViewport(viewport *QRect) {
	purego.SyscallN(addr_QPainter_setViewport, this.h, uintptr(viewport.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(viewport)
}

func (this *QPainter) SetViewport2(x int, y int, w int, h int) {
	purego.SyscallN(addr_QPainter_setViewport2, this.h, uintptr(x), uintptr(y), uintptr(w), uintptr(h))
	runtime.KeepAlive(this)
}

func (this *QPainter) SetViewTransformEnabled(enable bool) {
	purego.SyscallN(addr_QPainter_setViewTransformEnabled, this.h, uintptr(b2i(enable)))
	runtime.KeepAlive(this)
}

func (this *QPainter) ViewTransformEnabled() bool {
	__r0, _, _ := purego.SyscallN(addr_QPainter_viewTransformEnabled, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QPainter) StrokePath(path *QPainterPath, pen *QPen) {
	purego.SyscallN(addr_QPainter_strokePath, this.h, uintptr(path.UnsafePointer()), uintptr(pen.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(path)
	runtime.KeepAlive(pen)
}

func (this *QPainter) FillPath(path *QPainterPath, brush *QBrush) {
	purego.SyscallN(addr_QPainter_fillPath, this.h, uintptr(path.UnsafePointer()), uintptr(brush.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(path)
	runtime.KeepAlive(brush)
}

func (this *QPainter) DrawPath(path *QPainterPath) {
	purego.SyscallN(addr_QPainter_drawPath, this.h, uintptr(path.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(path)
}

func (this *QPainter) DrawPoint(pt *QPointF) {
	purego.SyscallN(addr_QPainter_drawPoint, this.h, uintptr(pt.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pt)
}

func (this *QPainter) DrawPointWithQPoint(p *QPoint) {
	purego.SyscallN(addr_QPainter_drawPoint2, this.h, uintptr(p.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
}

func (this *QPainter) DrawPoint2(x int, y int) {
	purego.SyscallN(addr_QPainter_drawPoint3, this.h, uintptr(x), uintptr(y))
	runtime.KeepAlive(this)
}

func (this *QPainter) DrawPoints(points *QPointF, pointCount int) {
	purego.SyscallN(addr_QPainter_drawPoints, this.h, uintptr(points.UnsafePointer()), uintptr(pointCount))
	runtime.KeepAlive(this)
	runtime.KeepAlive(points)
}

func (this *QPainter) DrawPointsWithPoints(points *QPolygonF) {
	purego.SyscallN(addr_QPainter_drawPoints2, this.h, uintptr(points.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(points)
}

func (this *QPainter) DrawPoints2(points *QPoint, pointCount int) {
	purego.SyscallN(addr_QPainter_drawPoints3, this.h, uintptr(points.UnsafePointer()), uintptr(pointCount))
	runtime.KeepAlive(this)
	runtime.KeepAlive(points)
}

func (this *QPainter) DrawPoints3(points *QPolygon) {
	purego.SyscallN(addr_QPainter_drawPoints4, this.h, uintptr(points.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(points)
}

func (this *QPainter) DrawLine(line *QLineF) {
	purego.SyscallN(addr_QPainter_drawLine, this.h, uintptr(line.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(line)
}

func (this *QPainter) DrawLineWithLine(line *QLine) {
	purego.SyscallN(addr_QPainter_drawLine2, this.h, uintptr(line.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(line)
}

func (this *QPainter) DrawLine2(x1 int, y1 int, x2 int, y2 int) {
	purego.SyscallN(addr_QPainter_drawLine3, this.h, uintptr(x1), uintptr(y1), uintptr(x2), uintptr(y2))
	runtime.KeepAlive(this)
}

func (this *QPainter) DrawLine3(p1 *QPoint, p2 *QPoint) {
	purego.SyscallN(addr_QPainter_drawLine4, this.h, uintptr(p1.UnsafePointer()), uintptr(p2.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p1)
	runtime.KeepAlive(p2)
}

func (this *QPainter) DrawLine4(p1 *QPointF, p2 *QPointF) {
	purego.SyscallN(addr_QPainter_drawLine5, this.h, uintptr(p1.UnsafePointer()), uintptr(p2.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p1)
	runtime.KeepAlive(p2)
}

func (this *QPainter) DrawLines(lines *QLineF, lineCount int) {
	purego.SyscallN(addr_QPainter_drawLines, this.h, uintptr(lines.UnsafePointer()), uintptr(lineCount))
	runtime.KeepAlive(this)
	runtime.KeepAlive(lines)
}

func (this *QPainter) DrawLinesWithLines(lines []*QLineF) {
	lines_arr := make([]uintptr, len(lines))
	for __i, __e := range lines {
		lines_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var lines_p uintptr
	if len(lines_arr) > 0 {
		lines_p = uintptr(unsafe.Pointer(&lines_arr[0]))
	}
	purego.SyscallN(addr_QPainter_drawLines2, this.h, lines_p, uintptr(len(lines)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(lines)
	runtime.KeepAlive(lines_arr)
}

func (this *QPainter) DrawLines2(pointPairs *QPointF, lineCount int) {
	purego.SyscallN(addr_QPainter_drawLines3, this.h, uintptr(pointPairs.UnsafePointer()), uintptr(lineCount))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pointPairs)
}

func (this *QPainter) DrawLinesWithPointPairs(pointPairs []*QPointF) {
	pointPairs_arr := make([]uintptr, len(pointPairs))
	for __i, __e := range pointPairs {
		pointPairs_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var pointPairs_p uintptr
	if len(pointPairs_arr) > 0 {
		pointPairs_p = uintptr(unsafe.Pointer(&pointPairs_arr[0]))
	}
	purego.SyscallN(addr_QPainter_drawLines4, this.h, pointPairs_p, uintptr(len(pointPairs)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pointPairs)
	runtime.KeepAlive(pointPairs_arr)
}

func (this *QPainter) DrawLines3(lines *QLine, lineCount int) {
	purego.SyscallN(addr_QPainter_drawLines5, this.h, uintptr(lines.UnsafePointer()), uintptr(lineCount))
	runtime.KeepAlive(this)
	runtime.KeepAlive(lines)
}

func (this *QPainter) DrawLines4(lines []*QLine) {
	lines_arr := make([]uintptr, len(lines))
	for __i, __e := range lines {
		lines_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var lines_p uintptr
	if len(lines_arr) > 0 {
		lines_p = uintptr(unsafe.Pointer(&lines_arr[0]))
	}
	purego.SyscallN(addr_QPainter_drawLines6, this.h, lines_p, uintptr(len(lines)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(lines)
	runtime.KeepAlive(lines_arr)
}

func (this *QPainter) DrawLines5(pointPairs *QPoint, lineCount int) {
	purego.SyscallN(addr_QPainter_drawLines7, this.h, uintptr(pointPairs.UnsafePointer()), uintptr(lineCount))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pointPairs)
}

func (this *QPainter) DrawLines6(pointPairs []*QPoint) {
	pointPairs_arr := make([]uintptr, len(pointPairs))
	for __i, __e := range pointPairs {
		pointPairs_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var pointPairs_p uintptr
	if len(pointPairs_arr) > 0 {
		pointPairs_p = uintptr(unsafe.Pointer(&pointPairs_arr[0]))
	}
	purego.SyscallN(addr_QPainter_drawLines8, this.h, pointPairs_p, uintptr(len(pointPairs)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pointPairs)
	runtime.KeepAlive(pointPairs_arr)
}

func (this *QPainter) DrawRect(rect *QRectF) {
	purego.SyscallN(addr_QPainter_drawRect, this.h, uintptr(rect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
}

func (this *QPainter) DrawRect2(x1 int, y1 int, w int, h int) {
	purego.SyscallN(addr_QPainter_drawRect2, this.h, uintptr(x1), uintptr(y1), uintptr(w), uintptr(h))
	runtime.KeepAlive(this)
}

func (this *QPainter) DrawRectWithRect(rect *QRect) {
	purego.SyscallN(addr_QPainter_drawRect3, this.h, uintptr(rect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
}

func (this *QPainter) DrawRects(rects *QRectF, rectCount int) {
	purego.SyscallN(addr_QPainter_drawRects, this.h, uintptr(rects.UnsafePointer()), uintptr(rectCount))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rects)
}

func (this *QPainter) DrawRectsWithRectangles(rectangles []*QRectF) {
	rectangles_arr := make([]uintptr, len(rectangles))
	for __i, __e := range rectangles {
		rectangles_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var rectangles_p uintptr
	if len(rectangles_arr) > 0 {
		rectangles_p = uintptr(unsafe.Pointer(&rectangles_arr[0]))
	}
	purego.SyscallN(addr_QPainter_drawRects2, this.h, rectangles_p, uintptr(len(rectangles)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rectangles)
	runtime.KeepAlive(rectangles_arr)
}

func (this *QPainter) DrawRects2(rects *QRect, rectCount int) {
	purego.SyscallN(addr_QPainter_drawRects3, this.h, uintptr(rects.UnsafePointer()), uintptr(rectCount))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rects)
}

func (this *QPainter) DrawRects3(rectangles []*QRect) {
	rectangles_arr := make([]uintptr, len(rectangles))
	for __i, __e := range rectangles {
		rectangles_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var rectangles_p uintptr
	if len(rectangles_arr) > 0 {
		rectangles_p = uintptr(unsafe.Pointer(&rectangles_arr[0]))
	}
	purego.SyscallN(addr_QPainter_drawRects4, this.h, rectangles_p, uintptr(len(rectangles)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rectangles)
	runtime.KeepAlive(rectangles_arr)
}

func (this *QPainter) DrawEllipse(r *QRectF) {
	purego.SyscallN(addr_QPainter_drawEllipse, this.h, uintptr(r.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(r)
}

func (this *QPainter) DrawEllipseWithQRect(r *QRect) {
	purego.SyscallN(addr_QPainter_drawEllipse2, this.h, uintptr(r.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(r)
}

func (this *QPainter) DrawEllipse2(x int, y int, w int, h int) {
	purego.SyscallN(addr_QPainter_drawEllipse3, this.h, uintptr(x), uintptr(y), uintptr(w), uintptr(h))
	runtime.KeepAlive(this)
}

func (this *QPainter) DrawEllipse3(center *QPointF, rx float64, ry float64) {
	pg_QPainter_drawEllipse4(this.h, uintptr(center.UnsafePointer()), rx, ry)
	runtime.KeepAlive(this)
	runtime.KeepAlive(center)
}

func (this *QPainter) DrawEllipse4(center *QPoint, rx int, ry int) {
	purego.SyscallN(addr_QPainter_drawEllipse5, this.h, uintptr(center.UnsafePointer()), uintptr(rx), uintptr(ry))
	runtime.KeepAlive(this)
	runtime.KeepAlive(center)
}

func (this *QPainter) DrawPolyline(points *QPointF, pointCount int) {
	purego.SyscallN(addr_QPainter_drawPolyline, this.h, uintptr(points.UnsafePointer()), uintptr(pointCount))
	runtime.KeepAlive(this)
	runtime.KeepAlive(points)
}

func (this *QPainter) DrawPolylineWithPolyline(polyline *QPolygonF) {
	purego.SyscallN(addr_QPainter_drawPolyline2, this.h, uintptr(polyline.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(polyline)
}

func (this *QPainter) DrawPolyline2(points *QPoint, pointCount int) {
	purego.SyscallN(addr_QPainter_drawPolyline3, this.h, uintptr(points.UnsafePointer()), uintptr(pointCount))
	runtime.KeepAlive(this)
	runtime.KeepAlive(points)
}

func (this *QPainter) DrawPolylineWithPolygon(polygon *QPolygon) {
	purego.SyscallN(addr_QPainter_drawPolyline4, this.h, uintptr(polygon.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(polygon)
}

func (this *QPainter) DrawPolygon(points *QPointF, pointCount int) {
	purego.SyscallN(addr_QPainter_drawPolygon, this.h, uintptr(points.UnsafePointer()), uintptr(pointCount))
	runtime.KeepAlive(this)
	runtime.KeepAlive(points)
}

func (this *QPainter) DrawPolygonWithPolygon(polygon *QPolygonF) {
	purego.SyscallN(addr_QPainter_drawPolygon2, this.h, uintptr(polygon.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(polygon)
}

func (this *QPainter) DrawPolygon2(points *QPoint, pointCount int) {
	purego.SyscallN(addr_QPainter_drawPolygon3, this.h, uintptr(points.UnsafePointer()), uintptr(pointCount))
	runtime.KeepAlive(this)
	runtime.KeepAlive(points)
}

func (this *QPainter) DrawPolygon3(polygon *QPolygon) {
	purego.SyscallN(addr_QPainter_drawPolygon4, this.h, uintptr(polygon.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(polygon)
}

func (this *QPainter) DrawConvexPolygon(points *QPointF, pointCount int) {
	purego.SyscallN(addr_QPainter_drawConvexPolygon, this.h, uintptr(points.UnsafePointer()), uintptr(pointCount))
	runtime.KeepAlive(this)
	runtime.KeepAlive(points)
}

func (this *QPainter) DrawConvexPolygonWithPolygon(polygon *QPolygonF) {
	purego.SyscallN(addr_QPainter_drawConvexPolygon2, this.h, uintptr(polygon.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(polygon)
}

func (this *QPainter) DrawConvexPolygon2(points *QPoint, pointCount int) {
	purego.SyscallN(addr_QPainter_drawConvexPolygon3, this.h, uintptr(points.UnsafePointer()), uintptr(pointCount))
	runtime.KeepAlive(this)
	runtime.KeepAlive(points)
}

func (this *QPainter) DrawConvexPolygon3(polygon *QPolygon) {
	purego.SyscallN(addr_QPainter_drawConvexPolygon4, this.h, uintptr(polygon.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(polygon)
}

func (this *QPainter) DrawArc(rect *QRectF, a int, alen int) {
	purego.SyscallN(addr_QPainter_drawArc, this.h, uintptr(rect.UnsafePointer()), uintptr(a), uintptr(alen))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
}

func (this *QPainter) DrawArc2(arg1 *QRect, a int, alen int) {
	purego.SyscallN(addr_QPainter_drawArc2, this.h, uintptr(arg1.UnsafePointer()), uintptr(a), uintptr(alen))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QPainter) DrawArc3(x int, y int, w int, h int, a int, alen int) {
	purego.SyscallN(addr_QPainter_drawArc3, this.h, uintptr(x), uintptr(y), uintptr(w), uintptr(h), uintptr(a), uintptr(alen))
	runtime.KeepAlive(this)
}

func (this *QPainter) DrawPie(rect *QRectF, a int, alen int) {
	purego.SyscallN(addr_QPainter_drawPie, this.h, uintptr(rect.UnsafePointer()), uintptr(a), uintptr(alen))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
}

func (this *QPainter) DrawPie2(x int, y int, w int, h int, a int, alen int) {
	purego.SyscallN(addr_QPainter_drawPie2, this.h, uintptr(x), uintptr(y), uintptr(w), uintptr(h), uintptr(a), uintptr(alen))
	runtime.KeepAlive(this)
}

func (this *QPainter) DrawPie3(arg1 *QRect, a int, alen int) {
	purego.SyscallN(addr_QPainter_drawPie3, this.h, uintptr(arg1.UnsafePointer()), uintptr(a), uintptr(alen))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QPainter) DrawChord(rect *QRectF, a int, alen int) {
	purego.SyscallN(addr_QPainter_drawChord, this.h, uintptr(rect.UnsafePointer()), uintptr(a), uintptr(alen))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
}

func (this *QPainter) DrawChord2(x int, y int, w int, h int, a int, alen int) {
	purego.SyscallN(addr_QPainter_drawChord2, this.h, uintptr(x), uintptr(y), uintptr(w), uintptr(h), uintptr(a), uintptr(alen))
	runtime.KeepAlive(this)
}

func (this *QPainter) DrawChord3(arg1 *QRect, a int, alen int) {
	purego.SyscallN(addr_QPainter_drawChord3, this.h, uintptr(arg1.UnsafePointer()), uintptr(a), uintptr(alen))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QPainter) DrawRoundedRect(rect *QRectF, xRadius float64, yRadius float64) {
	pg_QPainter_drawRoundedRect(this.h, uintptr(rect.UnsafePointer()), xRadius, yRadius)
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
}

func (this *QPainter) DrawRoundedRect2(x int, y int, w int, h int, xRadius float64, yRadius float64) {
	pg_QPainter_drawRoundedRect2(this.h, int32(x), int32(y), int32(w), int32(h), xRadius, yRadius)
	runtime.KeepAlive(this)
}

func (this *QPainter) DrawRoundedRect3(rect *QRect, xRadius float64, yRadius float64) {
	pg_QPainter_drawRoundedRect3(this.h, uintptr(rect.UnsafePointer()), xRadius, yRadius)
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
}

func (this *QPainter) DrawTiledPixmap(rect *QRectF, pm *QPixmap) {
	purego.SyscallN(addr_QPainter_drawTiledPixmap, this.h, uintptr(rect.UnsafePointer()), uintptr(pm.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
	runtime.KeepAlive(pm)
}

func (this *QPainter) DrawTiledPixmap2(x int, y int, w int, h int, arg5 *QPixmap) {
	purego.SyscallN(addr_QPainter_drawTiledPixmap2, this.h, uintptr(x), uintptr(y), uintptr(w), uintptr(h), uintptr(arg5.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg5)
}

func (this *QPainter) DrawTiledPixmap3(arg1 *QRect, arg2 *QPixmap) {
	purego.SyscallN(addr_QPainter_drawTiledPixmap3, this.h, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
}

func (this *QPainter) DrawPicture(p *QPointF, picture *QPicture) {
	purego.SyscallN(addr_QPainter_drawPicture, this.h, uintptr(p.UnsafePointer()), uintptr(picture.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
	runtime.KeepAlive(picture)
}

func (this *QPainter) DrawPicture2(x int, y int, picture *QPicture) {
	purego.SyscallN(addr_QPainter_drawPicture2, this.h, uintptr(x), uintptr(y), uintptr(picture.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(picture)
}

func (this *QPainter) DrawPicture3(p *QPoint, picture *QPicture) {
	purego.SyscallN(addr_QPainter_drawPicture3, this.h, uintptr(p.UnsafePointer()), uintptr(picture.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
	runtime.KeepAlive(picture)
}

func (this *QPainter) DrawPixmap(targetRect *QRectF, pixmap *QPixmap, sourceRect *QRectF) {
	purego.SyscallN(addr_QPainter_drawPixmap, this.h, uintptr(targetRect.UnsafePointer()), uintptr(pixmap.UnsafePointer()), uintptr(sourceRect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(targetRect)
	runtime.KeepAlive(pixmap)
	runtime.KeepAlive(sourceRect)
}

func (this *QPainter) DrawPixmap2(targetRect *QRect, pixmap *QPixmap, sourceRect *QRect) {
	purego.SyscallN(addr_QPainter_drawPixmap2, this.h, uintptr(targetRect.UnsafePointer()), uintptr(pixmap.UnsafePointer()), uintptr(sourceRect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(targetRect)
	runtime.KeepAlive(pixmap)
	runtime.KeepAlive(sourceRect)
}

func (this *QPainter) DrawPixmap3(x int, y int, w int, h int, pm *QPixmap, sx int, sy int, sw int, sh int) {
	purego.SyscallN(addr_QPainter_drawPixmap3, this.h, uintptr(x), uintptr(y), uintptr(w), uintptr(h), uintptr(pm.UnsafePointer()), uintptr(sx), uintptr(sy), uintptr(sw), uintptr(sh))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pm)
}

func (this *QPainter) DrawPixmap4(x int, y int, pm *QPixmap, sx int, sy int, sw int, sh int) {
	purego.SyscallN(addr_QPainter_drawPixmap4, this.h, uintptr(x), uintptr(y), uintptr(pm.UnsafePointer()), uintptr(sx), uintptr(sy), uintptr(sw), uintptr(sh))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pm)
}

func (this *QPainter) DrawPixmap5(p *QPointF, pm *QPixmap, sr *QRectF) {
	purego.SyscallN(addr_QPainter_drawPixmap5, this.h, uintptr(p.UnsafePointer()), uintptr(pm.UnsafePointer()), uintptr(sr.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
	runtime.KeepAlive(pm)
	runtime.KeepAlive(sr)
}

func (this *QPainter) DrawPixmap6(p *QPoint, pm *QPixmap, sr *QRect) {
	purego.SyscallN(addr_QPainter_drawPixmap6, this.h, uintptr(p.UnsafePointer()), uintptr(pm.UnsafePointer()), uintptr(sr.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
	runtime.KeepAlive(pm)
	runtime.KeepAlive(sr)
}

func (this *QPainter) DrawPixmap7(p *QPointF, pm *QPixmap) {
	purego.SyscallN(addr_QPainter_drawPixmap7, this.h, uintptr(p.UnsafePointer()), uintptr(pm.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
	runtime.KeepAlive(pm)
}

func (this *QPainter) DrawPixmap8(p *QPoint, pm *QPixmap) {
	purego.SyscallN(addr_QPainter_drawPixmap8, this.h, uintptr(p.UnsafePointer()), uintptr(pm.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
	runtime.KeepAlive(pm)
}

func (this *QPainter) DrawPixmap9(x int, y int, pm *QPixmap) {
	purego.SyscallN(addr_QPainter_drawPixmap9, this.h, uintptr(x), uintptr(y), uintptr(pm.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pm)
}

func (this *QPainter) DrawPixmap10(r *QRect, pm *QPixmap) {
	purego.SyscallN(addr_QPainter_drawPixmap10, this.h, uintptr(r.UnsafePointer()), uintptr(pm.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(r)
	runtime.KeepAlive(pm)
}

func (this *QPainter) DrawPixmap11(x int, y int, w int, h int, pm *QPixmap) {
	purego.SyscallN(addr_QPainter_drawPixmap11, this.h, uintptr(x), uintptr(y), uintptr(w), uintptr(h), uintptr(pm.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pm)
}

func (this *QPainter) DrawPixmapFragments(fragments *QPainter__PixmapFragment, fragmentCount int, pixmap *QPixmap) {
	purego.SyscallN(addr_QPainter_drawPixmapFragments, this.h, uintptr(fragments.UnsafePointer()), uintptr(fragmentCount), uintptr(pixmap.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(fragments)
	runtime.KeepAlive(pixmap)
}

func (this *QPainter) DrawImage(targetRect *QRectF, image *QImage, sourceRect *QRectF) {
	purego.SyscallN(addr_QPainter_drawImage, this.h, uintptr(targetRect.UnsafePointer()), uintptr(image.UnsafePointer()), uintptr(sourceRect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(targetRect)
	runtime.KeepAlive(image)
	runtime.KeepAlive(sourceRect)
}

func (this *QPainter) DrawImage2(targetRect *QRect, image *QImage, sourceRect *QRect) {
	purego.SyscallN(addr_QPainter_drawImage2, this.h, uintptr(targetRect.UnsafePointer()), uintptr(image.UnsafePointer()), uintptr(sourceRect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(targetRect)
	runtime.KeepAlive(image)
	runtime.KeepAlive(sourceRect)
}

func (this *QPainter) DrawImage3(p *QPointF, image *QImage, sr *QRectF) {
	purego.SyscallN(addr_QPainter_drawImage3, this.h, uintptr(p.UnsafePointer()), uintptr(image.UnsafePointer()), uintptr(sr.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
	runtime.KeepAlive(image)
	runtime.KeepAlive(sr)
}

func (this *QPainter) DrawImage4(p *QPoint, image *QImage, sr *QRect) {
	purego.SyscallN(addr_QPainter_drawImage4, this.h, uintptr(p.UnsafePointer()), uintptr(image.UnsafePointer()), uintptr(sr.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
	runtime.KeepAlive(image)
	runtime.KeepAlive(sr)
}

func (this *QPainter) DrawImage5(r *QRectF, image *QImage) {
	purego.SyscallN(addr_QPainter_drawImage5, this.h, uintptr(r.UnsafePointer()), uintptr(image.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(r)
	runtime.KeepAlive(image)
}

func (this *QPainter) DrawImage6(r *QRect, image *QImage) {
	purego.SyscallN(addr_QPainter_drawImage6, this.h, uintptr(r.UnsafePointer()), uintptr(image.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(r)
	runtime.KeepAlive(image)
}

func (this *QPainter) DrawImage7(p *QPointF, image *QImage) {
	purego.SyscallN(addr_QPainter_drawImage7, this.h, uintptr(p.UnsafePointer()), uintptr(image.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
	runtime.KeepAlive(image)
}

func (this *QPainter) DrawImage8(p *QPoint, image *QImage) {
	purego.SyscallN(addr_QPainter_drawImage8, this.h, uintptr(p.UnsafePointer()), uintptr(image.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
	runtime.KeepAlive(image)
}

func (this *QPainter) DrawImage9(x int, y int, image *QImage) {
	purego.SyscallN(addr_QPainter_drawImage9, this.h, uintptr(x), uintptr(y), uintptr(image.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(image)
}

func (this *QPainter) SetLayoutDirection(direction LayoutDirection) {
	purego.SyscallN(addr_QPainter_setLayoutDirection, this.h, uintptr(direction))
	runtime.KeepAlive(this)
}

func (this *QPainter) LayoutDirection() LayoutDirection {
	__r0, _, _ := purego.SyscallN(addr_QPainter_layoutDirection, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return LayoutDirection(__r)
}

func (this *QPainter) DrawGlyphRun(position *QPointF, glyphRun *QGlyphRun) {
	purego.SyscallN(addr_QPainter_drawGlyphRun, this.h, uintptr(position.UnsafePointer()), uintptr(glyphRun.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(position)
	runtime.KeepAlive(glyphRun)
}

func (this *QPainter) DrawStaticText(topLeftPosition *QPointF, staticText *QStaticText) {
	purego.SyscallN(addr_QPainter_drawStaticText, this.h, uintptr(topLeftPosition.UnsafePointer()), uintptr(staticText.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(topLeftPosition)
	runtime.KeepAlive(staticText)
}

func (this *QPainter) DrawStaticText2(topLeftPosition *QPoint, staticText *QStaticText) {
	purego.SyscallN(addr_QPainter_drawStaticText2, this.h, uintptr(topLeftPosition.UnsafePointer()), uintptr(staticText.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(topLeftPosition)
	runtime.KeepAlive(staticText)
}

func (this *QPainter) DrawStaticText3(left int, top int, staticText *QStaticText) {
	purego.SyscallN(addr_QPainter_drawStaticText3, this.h, uintptr(left), uintptr(top), uintptr(staticText.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(staticText)
}

func (this *QPainter) DrawText(p *QPointF, s string) {
	var s_p uintptr
	if len(s) > 0 {
		s_p = uintptr(unsafe.Pointer(unsafe.StringData(s)))
	}
	purego.SyscallN(addr_QPainter_drawText, this.h, uintptr(p.UnsafePointer()), s_p, uintptr(len(s)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
	runtime.KeepAlive(s)
}

func (this *QPainter) DrawText2(p *QPoint, s string) {
	var s_p uintptr
	if len(s) > 0 {
		s_p = uintptr(unsafe.Pointer(unsafe.StringData(s)))
	}
	purego.SyscallN(addr_QPainter_drawText2, this.h, uintptr(p.UnsafePointer()), s_p, uintptr(len(s)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
	runtime.KeepAlive(s)
}

func (this *QPainter) DrawText3(x int, y int, s string) {
	var s_p uintptr
	if len(s) > 0 {
		s_p = uintptr(unsafe.Pointer(unsafe.StringData(s)))
	}
	purego.SyscallN(addr_QPainter_drawText3, this.h, uintptr(x), uintptr(y), s_p, uintptr(len(s)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
}

func (this *QPainter) DrawText4(p *QPointF, str string, tf int, justificationPadding int) {
	var str_p uintptr
	if len(str) > 0 {
		str_p = uintptr(unsafe.Pointer(unsafe.StringData(str)))
	}
	purego.SyscallN(addr_QPainter_drawText4, this.h, uintptr(p.UnsafePointer()), str_p, uintptr(len(str)), uintptr(tf), uintptr(justificationPadding))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
	runtime.KeepAlive(str)
}

func (this *QPainter) DrawText5(r *QRectF, flags int, text string) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QPainter_drawText5, this.h, uintptr(r.UnsafePointer()), uintptr(flags), text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(r)
	runtime.KeepAlive(text)
}

func (this *QPainter) DrawText6(r *QRect, flags int, text string) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QPainter_drawText6, this.h, uintptr(r.UnsafePointer()), uintptr(flags), text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(r)
	runtime.KeepAlive(text)
}

func (this *QPainter) DrawText7(x int, y int, w int, h int, flags int, text string) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QPainter_drawText7, this.h, uintptr(x), uintptr(y), uintptr(w), uintptr(h), uintptr(flags), text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
}

func (this *QPainter) DrawText8(r *QRectF, text string) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QPainter_drawText8, this.h, uintptr(r.UnsafePointer()), text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(r)
	runtime.KeepAlive(text)
}

func (this *QPainter) BoundingRect(rect *QRectF, flags int, text string) *QRectF {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QPainter_boundingRect, this.h, uintptr(rect.UnsafePointer()), uintptr(flags), text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
	runtime.KeepAlive(text)
	return newownedQRectF(__r)
}

func (this *QPainter) BoundingRect2(rect *QRect, flags int, text string) *QRect {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QPainter_boundingRect2, this.h, uintptr(rect.UnsafePointer()), uintptr(flags), text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
	runtime.KeepAlive(text)
	return newownedQRect(__r)
}

func (this *QPainter) BoundingRect3(x int, y int, w int, h int, flags int, text string) *QRect {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QPainter_boundingRect3, this.h, uintptr(x), uintptr(y), uintptr(w), uintptr(h), uintptr(flags), text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
	return newownedQRect(__r)
}

func (this *QPainter) BoundingRect4(rect *QRectF, text string) *QRectF {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QPainter_boundingRect4, this.h, uintptr(rect.UnsafePointer()), text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
	runtime.KeepAlive(text)
	return newownedQRectF(__r)
}

func (this *QPainter) DrawTextItem(p *QPointF, ti *QTextItem) {
	purego.SyscallN(addr_QPainter_drawTextItem, this.h, uintptr(p.UnsafePointer()), uintptr(ti.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
	runtime.KeepAlive(ti)
}

func (this *QPainter) DrawTextItem2(x int, y int, ti *QTextItem) {
	purego.SyscallN(addr_QPainter_drawTextItem2, this.h, uintptr(x), uintptr(y), uintptr(ti.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(ti)
}

func (this *QPainter) DrawTextItem3(p *QPoint, ti *QTextItem) {
	purego.SyscallN(addr_QPainter_drawTextItem3, this.h, uintptr(p.UnsafePointer()), uintptr(ti.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
	runtime.KeepAlive(ti)
}

func (this *QPainter) FillRect(arg1 *QRectF, arg2 *QBrush) {
	purego.SyscallN(addr_QPainter_fillRect, this.h, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
}

func (this *QPainter) FillRect2(x int, y int, w int, h int, arg5 *QBrush) {
	purego.SyscallN(addr_QPainter_fillRect2, this.h, uintptr(x), uintptr(y), uintptr(w), uintptr(h), uintptr(arg5.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg5)
}

func (this *QPainter) FillRect3(arg1 *QRect, arg2 *QBrush) {
	purego.SyscallN(addr_QPainter_fillRect3, this.h, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
}

func (this *QPainter) FillRect4(arg1 *QRectF, color *QColor) {
	purego.SyscallN(addr_QPainter_fillRect4, this.h, uintptr(arg1.UnsafePointer()), uintptr(color.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(color)
}

func (this *QPainter) FillRect5(x int, y int, w int, h int, color *QColor) {
	purego.SyscallN(addr_QPainter_fillRect5, this.h, uintptr(x), uintptr(y), uintptr(w), uintptr(h), uintptr(color.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(color)
}

func (this *QPainter) FillRect6(arg1 *QRect, color *QColor) {
	purego.SyscallN(addr_QPainter_fillRect6, this.h, uintptr(arg1.UnsafePointer()), uintptr(color.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(color)
}

func (this *QPainter) FillRect7(x int, y int, w int, h int, c GlobalColor) {
	purego.SyscallN(addr_QPainter_fillRect7, this.h, uintptr(x), uintptr(y), uintptr(w), uintptr(h), uintptr(c))
	runtime.KeepAlive(this)
}

func (this *QPainter) FillRect8(r *QRect, c GlobalColor) {
	purego.SyscallN(addr_QPainter_fillRect8, this.h, uintptr(r.UnsafePointer()), uintptr(c))
	runtime.KeepAlive(this)
	runtime.KeepAlive(r)
}

func (this *QPainter) FillRect9(r *QRectF, c GlobalColor) {
	purego.SyscallN(addr_QPainter_fillRect9, this.h, uintptr(r.UnsafePointer()), uintptr(c))
	runtime.KeepAlive(this)
	runtime.KeepAlive(r)
}

func (this *QPainter) FillRect10(x int, y int, w int, h int, style BrushStyle) {
	purego.SyscallN(addr_QPainter_fillRect10, this.h, uintptr(x), uintptr(y), uintptr(w), uintptr(h), uintptr(style))
	runtime.KeepAlive(this)
}

func (this *QPainter) FillRect11(r *QRect, style BrushStyle) {
	purego.SyscallN(addr_QPainter_fillRect11, this.h, uintptr(r.UnsafePointer()), uintptr(style))
	runtime.KeepAlive(this)
	runtime.KeepAlive(r)
}

func (this *QPainter) FillRect12(r *QRectF, style BrushStyle) {
	purego.SyscallN(addr_QPainter_fillRect12, this.h, uintptr(r.UnsafePointer()), uintptr(style))
	runtime.KeepAlive(this)
	runtime.KeepAlive(r)
}

func (this *QPainter) FillRect13(x int, y int, w int, h int, preset QGradient__Preset) {
	purego.SyscallN(addr_QPainter_fillRect13, this.h, uintptr(x), uintptr(y), uintptr(w), uintptr(h), uintptr(preset))
	runtime.KeepAlive(this)
}

func (this *QPainter) FillRect14(r *QRect, preset QGradient__Preset) {
	purego.SyscallN(addr_QPainter_fillRect14, this.h, uintptr(r.UnsafePointer()), uintptr(preset))
	runtime.KeepAlive(this)
	runtime.KeepAlive(r)
}

func (this *QPainter) FillRect15(r *QRectF, preset QGradient__Preset) {
	purego.SyscallN(addr_QPainter_fillRect15, this.h, uintptr(r.UnsafePointer()), uintptr(preset))
	runtime.KeepAlive(this)
	runtime.KeepAlive(r)
}

func (this *QPainter) EraseRect(arg1 *QRectF) {
	purego.SyscallN(addr_QPainter_eraseRect, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QPainter) EraseRect2(x int, y int, w int, h int) {
	purego.SyscallN(addr_QPainter_eraseRect2, this.h, uintptr(x), uintptr(y), uintptr(w), uintptr(h))
	runtime.KeepAlive(this)
}

func (this *QPainter) EraseRectWithQRect(arg1 *QRect) {
	purego.SyscallN(addr_QPainter_eraseRect3, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QPainter) SetRenderHint(hint QPainter__RenderHint) {
	purego.SyscallN(addr_QPainter_setRenderHint, this.h, uintptr(hint))
	runtime.KeepAlive(this)
}

func (this *QPainter) SetRenderHints(hints QPainter__RenderHint) {
	purego.SyscallN(addr_QPainter_setRenderHints, this.h, uintptr(hints))
	runtime.KeepAlive(this)
}

func (this *QPainter) RenderHints() QPainter__RenderHint {
	__r0, _, _ := purego.SyscallN(addr_QPainter_renderHints, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QPainter__RenderHint(__r)
}

func (this *QPainter) TestRenderHint(hint QPainter__RenderHint) bool {
	__r0, _, _ := purego.SyscallN(addr_QPainter_testRenderHint, this.h, uintptr(hint))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QPainter) PaintEngine() *QPaintEngine {
	__r, _, _ := purego.SyscallN(addr_QPainter_paintEngine, this.h)
	runtime.KeepAlive(this)
	return bindQPaintEngine(__r)
}

func (this *QPainter) BeginNativePainting() {
	purego.SyscallN(addr_QPainter_beginNativePainting, this.h)
	runtime.KeepAlive(this)
}

func (this *QPainter) EndNativePainting() {
	purego.SyscallN(addr_QPainter_endNativePainting, this.h)
	runtime.KeepAlive(this)
}

func (this *QPainter) SetClipRect3(arg1 *QRectF, op ClipOperation) {
	purego.SyscallN(addr_QPainter_setClipRect4, this.h, uintptr(arg1.UnsafePointer()), uintptr(op))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QPainter) SetClipRect4(arg1 *QRect, op ClipOperation) {
	purego.SyscallN(addr_QPainter_setClipRect5, this.h, uintptr(arg1.UnsafePointer()), uintptr(op))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QPainter) SetClipRect5(x int, y int, w int, h int, op ClipOperation) {
	purego.SyscallN(addr_QPainter_setClipRect6, this.h, uintptr(x), uintptr(y), uintptr(w), uintptr(h), uintptr(op))
	runtime.KeepAlive(this)
}

func (this *QPainter) SetClipRegion2(arg1 *QRegion, op ClipOperation) {
	purego.SyscallN(addr_QPainter_setClipRegion2, this.h, uintptr(arg1.UnsafePointer()), uintptr(op))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QPainter) SetClipPath2(path *QPainterPath, op ClipOperation) {
	purego.SyscallN(addr_QPainter_setClipPath2, this.h, uintptr(path.UnsafePointer()), uintptr(op))
	runtime.KeepAlive(this)
	runtime.KeepAlive(path)
}

func (this *QPainter) SetTransform2(transform *QTransform, combine bool) {
	purego.SyscallN(addr_QPainter_setTransform2, this.h, uintptr(transform.UnsafePointer()), uintptr(b2i(combine)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(transform)
}

func (this *QPainter) SetWorldTransform2(matrix *QTransform, combine bool) {
	purego.SyscallN(addr_QPainter_setWorldTransform2, this.h, uintptr(matrix.UnsafePointer()), uintptr(b2i(combine)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(matrix)
}

func (this *QPainter) DrawPolygon4(points *QPointF, pointCount int, fillRule FillRule) {
	purego.SyscallN(addr_QPainter_drawPolygon5, this.h, uintptr(points.UnsafePointer()), uintptr(pointCount), uintptr(fillRule))
	runtime.KeepAlive(this)
	runtime.KeepAlive(points)
}

func (this *QPainter) DrawPolygon5(polygon *QPolygonF, fillRule FillRule) {
	purego.SyscallN(addr_QPainter_drawPolygon6, this.h, uintptr(polygon.UnsafePointer()), uintptr(fillRule))
	runtime.KeepAlive(this)
	runtime.KeepAlive(polygon)
}

func (this *QPainter) DrawPolygon6(points *QPoint, pointCount int, fillRule FillRule) {
	purego.SyscallN(addr_QPainter_drawPolygon7, this.h, uintptr(points.UnsafePointer()), uintptr(pointCount), uintptr(fillRule))
	runtime.KeepAlive(this)
	runtime.KeepAlive(points)
}

func (this *QPainter) DrawPolygon7(polygon *QPolygon, fillRule FillRule) {
	purego.SyscallN(addr_QPainter_drawPolygon8, this.h, uintptr(polygon.UnsafePointer()), uintptr(fillRule))
	runtime.KeepAlive(this)
	runtime.KeepAlive(polygon)
}

func (this *QPainter) DrawRoundedRect4(rect *QRectF, xRadius float64, yRadius float64, mode SizeMode) {
	pg_QPainter_drawRoundedRect4(this.h, uintptr(rect.UnsafePointer()), xRadius, yRadius, int32(mode))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
}

func (this *QPainter) DrawRoundedRect5(x int, y int, w int, h int, xRadius float64, yRadius float64, mode SizeMode) {
	pg_QPainter_drawRoundedRect5(this.h, int32(x), int32(y), int32(w), int32(h), xRadius, yRadius, int32(mode))
	runtime.KeepAlive(this)
}

func (this *QPainter) DrawRoundedRect6(rect *QRect, xRadius float64, yRadius float64, mode SizeMode) {
	pg_QPainter_drawRoundedRect6(this.h, uintptr(rect.UnsafePointer()), xRadius, yRadius, int32(mode))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
}

func (this *QPainter) DrawTiledPixmap4(rect *QRectF, pm *QPixmap, offset *QPointF) {
	purego.SyscallN(addr_QPainter_drawTiledPixmap4, this.h, uintptr(rect.UnsafePointer()), uintptr(pm.UnsafePointer()), uintptr(offset.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
	runtime.KeepAlive(pm)
	runtime.KeepAlive(offset)
}

func (this *QPainter) DrawTiledPixmap5(x int, y int, w int, h int, arg5 *QPixmap, sx int) {
	purego.SyscallN(addr_QPainter_drawTiledPixmap5, this.h, uintptr(x), uintptr(y), uintptr(w), uintptr(h), uintptr(arg5.UnsafePointer()), uintptr(sx))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg5)
}

func (this *QPainter) DrawTiledPixmap6(x int, y int, w int, h int, arg5 *QPixmap, sx int, sy int) {
	purego.SyscallN(addr_QPainter_drawTiledPixmap6, this.h, uintptr(x), uintptr(y), uintptr(w), uintptr(h), uintptr(arg5.UnsafePointer()), uintptr(sx), uintptr(sy))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg5)
}

func (this *QPainter) DrawTiledPixmap7(arg1 *QRect, arg2 *QPixmap, arg3 *QPoint) {
	purego.SyscallN(addr_QPainter_drawTiledPixmap7, this.h, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()), uintptr(arg3.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
	runtime.KeepAlive(arg3)
}

func (this *QPainter) DrawPixmapFragments2(fragments *QPainter__PixmapFragment, fragmentCount int, pixmap *QPixmap, hints QPainter__PixmapFragmentHint) {
	purego.SyscallN(addr_QPainter_drawPixmapFragments2, this.h, uintptr(fragments.UnsafePointer()), uintptr(fragmentCount), uintptr(pixmap.UnsafePointer()), uintptr(hints))
	runtime.KeepAlive(this)
	runtime.KeepAlive(fragments)
	runtime.KeepAlive(pixmap)
}

func (this *QPainter) DrawImage10(targetRect *QRectF, image *QImage, sourceRect *QRectF, flags ImageConversionFlag) {
	purego.SyscallN(addr_QPainter_drawImage10, this.h, uintptr(targetRect.UnsafePointer()), uintptr(image.UnsafePointer()), uintptr(sourceRect.UnsafePointer()), uintptr(flags))
	runtime.KeepAlive(this)
	runtime.KeepAlive(targetRect)
	runtime.KeepAlive(image)
	runtime.KeepAlive(sourceRect)
}

func (this *QPainter) DrawImage11(targetRect *QRect, image *QImage, sourceRect *QRect, flags ImageConversionFlag) {
	purego.SyscallN(addr_QPainter_drawImage11, this.h, uintptr(targetRect.UnsafePointer()), uintptr(image.UnsafePointer()), uintptr(sourceRect.UnsafePointer()), uintptr(flags))
	runtime.KeepAlive(this)
	runtime.KeepAlive(targetRect)
	runtime.KeepAlive(image)
	runtime.KeepAlive(sourceRect)
}

func (this *QPainter) DrawImage12(p *QPointF, image *QImage, sr *QRectF, flags ImageConversionFlag) {
	purego.SyscallN(addr_QPainter_drawImage12, this.h, uintptr(p.UnsafePointer()), uintptr(image.UnsafePointer()), uintptr(sr.UnsafePointer()), uintptr(flags))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
	runtime.KeepAlive(image)
	runtime.KeepAlive(sr)
}

func (this *QPainter) DrawImage13(p *QPoint, image *QImage, sr *QRect, flags ImageConversionFlag) {
	purego.SyscallN(addr_QPainter_drawImage13, this.h, uintptr(p.UnsafePointer()), uintptr(image.UnsafePointer()), uintptr(sr.UnsafePointer()), uintptr(flags))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
	runtime.KeepAlive(image)
	runtime.KeepAlive(sr)
}

func (this *QPainter) DrawImage14(x int, y int, image *QImage, sx int) {
	purego.SyscallN(addr_QPainter_drawImage14, this.h, uintptr(x), uintptr(y), uintptr(image.UnsafePointer()), uintptr(sx))
	runtime.KeepAlive(this)
	runtime.KeepAlive(image)
}

func (this *QPainter) DrawImage15(x int, y int, image *QImage, sx int, sy int) {
	purego.SyscallN(addr_QPainter_drawImage15, this.h, uintptr(x), uintptr(y), uintptr(image.UnsafePointer()), uintptr(sx), uintptr(sy))
	runtime.KeepAlive(this)
	runtime.KeepAlive(image)
}

func (this *QPainter) DrawImage16(x int, y int, image *QImage, sx int, sy int, sw int) {
	purego.SyscallN(addr_QPainter_drawImage16, this.h, uintptr(x), uintptr(y), uintptr(image.UnsafePointer()), uintptr(sx), uintptr(sy), uintptr(sw))
	runtime.KeepAlive(this)
	runtime.KeepAlive(image)
}

func (this *QPainter) DrawImage17(x int, y int, image *QImage, sx int, sy int, sw int, sh int) {
	purego.SyscallN(addr_QPainter_drawImage17, this.h, uintptr(x), uintptr(y), uintptr(image.UnsafePointer()), uintptr(sx), uintptr(sy), uintptr(sw), uintptr(sh))
	runtime.KeepAlive(this)
	runtime.KeepAlive(image)
}

func (this *QPainter) DrawImage18(x int, y int, image *QImage, sx int, sy int, sw int, sh int, flags ImageConversionFlag) {
	purego.SyscallN(addr_QPainter_drawImage18, this.h, uintptr(x), uintptr(y), uintptr(image.UnsafePointer()), uintptr(sx), uintptr(sy), uintptr(sw), uintptr(sh), uintptr(flags))
	runtime.KeepAlive(this)
	runtime.KeepAlive(image)
}

func (this *QPainter) DrawText9(r *QRectF, flags int, text string, br *QRectF) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QPainter_drawText9, this.h, uintptr(r.UnsafePointer()), uintptr(flags), text_p, uintptr(len(text)), uintptr(br.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(r)
	runtime.KeepAlive(text)
	runtime.KeepAlive(br)
}

func (this *QPainter) DrawText10(r *QRect, flags int, text string, br *QRect) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QPainter_drawText10, this.h, uintptr(r.UnsafePointer()), uintptr(flags), text_p, uintptr(len(text)), uintptr(br.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(r)
	runtime.KeepAlive(text)
	runtime.KeepAlive(br)
}

func (this *QPainter) DrawText11(x int, y int, w int, h int, flags int, text string, br *QRect) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QPainter_drawText11, this.h, uintptr(x), uintptr(y), uintptr(w), uintptr(h), uintptr(flags), text_p, uintptr(len(text)), uintptr(br.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
	runtime.KeepAlive(br)
}

func (this *QPainter) DrawText12(r *QRectF, text string, o *QTextOption) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QPainter_drawText12, this.h, uintptr(r.UnsafePointer()), text_p, uintptr(len(text)), uintptr(o.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(r)
	runtime.KeepAlive(text)
	runtime.KeepAlive(o)
}

func (this *QPainter) BoundingRect5(rect *QRectF, text string, o *QTextOption) *QRectF {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QPainter_boundingRect5, this.h, uintptr(rect.UnsafePointer()), text_p, uintptr(len(text)), uintptr(o.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
	runtime.KeepAlive(text)
	runtime.KeepAlive(o)
	return newownedQRectF(__r)
}

func (this *QPainter) SetRenderHint2(hint QPainter__RenderHint, on bool) {
	purego.SyscallN(addr_QPainter_setRenderHint2, this.h, uintptr(hint), uintptr(b2i(on)))
	runtime.KeepAlive(this)
}

func (this *QPainter) SetRenderHints2(hints QPainter__RenderHint, on bool) {
	purego.SyscallN(addr_QPainter_setRenderHints2, this.h, uintptr(hints), uintptr(b2i(on)))
	runtime.KeepAlive(this)
}

func (this *QPainter) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QPainter_delete(this.h)
	}
	this.UnsafeRelease()
}
