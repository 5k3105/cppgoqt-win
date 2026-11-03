package qt

type QPageLayout__Unit int64

const (
	QPageLayout__Millimeter QPageLayout__Unit = 0
	QPageLayout__Point QPageLayout__Unit = 1
	QPageLayout__Inch QPageLayout__Unit = 2
	QPageLayout__Pica QPageLayout__Unit = 3
	QPageLayout__Didot QPageLayout__Unit = 4
	QPageLayout__Cicero QPageLayout__Unit = 5
)

type QPageLayout__Orientation int64

const (
	QPageLayout__Portrait QPageLayout__Orientation = 0
	QPageLayout__Landscape QPageLayout__Orientation = 1
)

type QPageLayout__Mode int64

const (
	QPageLayout__StandardMode QPageLayout__Mode = 0
	QPageLayout__FullPageMode QPageLayout__Mode = 1
)

type QPageLayout__OutOfBoundsPolicy int64

const (
	QPageLayout__OutOfBoundsPolicy__Reject QPageLayout__OutOfBoundsPolicy = 0
	QPageLayout__OutOfBoundsPolicy__Clamp QPageLayout__OutOfBoundsPolicy = 1
)

