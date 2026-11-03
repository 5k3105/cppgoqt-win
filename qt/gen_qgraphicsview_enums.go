package qt

type QGraphicsView__ViewportAnchor int64

const (
	QGraphicsView__NoAnchor QGraphicsView__ViewportAnchor = 0
	QGraphicsView__AnchorViewCenter QGraphicsView__ViewportAnchor = 1
	QGraphicsView__AnchorUnderMouse QGraphicsView__ViewportAnchor = 2
)

type QGraphicsView__CacheModeFlag int64

const (
	QGraphicsView__CacheNone QGraphicsView__CacheModeFlag = 0
	QGraphicsView__CacheBackground QGraphicsView__CacheModeFlag = 1
)

type QGraphicsView__DragMode int64

const (
	QGraphicsView__NoDrag QGraphicsView__DragMode = 0
	QGraphicsView__ScrollHandDrag QGraphicsView__DragMode = 1
	QGraphicsView__RubberBandDrag QGraphicsView__DragMode = 2
)

type QGraphicsView__ViewportUpdateMode int64

const (
	QGraphicsView__FullViewportUpdate QGraphicsView__ViewportUpdateMode = 0
	QGraphicsView__MinimalViewportUpdate QGraphicsView__ViewportUpdateMode = 1
	QGraphicsView__SmartViewportUpdate QGraphicsView__ViewportUpdateMode = 2
	QGraphicsView__NoViewportUpdate QGraphicsView__ViewportUpdateMode = 3
	QGraphicsView__BoundingRectViewportUpdate QGraphicsView__ViewportUpdateMode = 4
)

type QGraphicsView__OptimizationFlag int64

const (
	QGraphicsView__DontSavePainterState QGraphicsView__OptimizationFlag = 1
	QGraphicsView__DontAdjustForAntialiasing QGraphicsView__OptimizationFlag = 2
	QGraphicsView__IndirectPainting QGraphicsView__OptimizationFlag = 4
)

