package qt

type QGraphicsEffect__ChangeFlag int64

const (
	QGraphicsEffect__SourceAttached QGraphicsEffect__ChangeFlag = 1
	QGraphicsEffect__SourceDetached QGraphicsEffect__ChangeFlag = 2
	QGraphicsEffect__SourceBoundingRectChanged QGraphicsEffect__ChangeFlag = 4
	QGraphicsEffect__SourceInvalidated QGraphicsEffect__ChangeFlag = 8
)

type QGraphicsEffect__PixmapPadMode int64

const (
	QGraphicsEffect__NoPad QGraphicsEffect__PixmapPadMode = 0
	QGraphicsEffect__PadToTransparentBorder QGraphicsEffect__PixmapPadMode = 1
	QGraphicsEffect__PadToEffectiveBoundingRect QGraphicsEffect__PixmapPadMode = 2
)

