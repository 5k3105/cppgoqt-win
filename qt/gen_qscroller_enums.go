package qt

type QScroller__State int64

const (
	QScroller__Inactive QScroller__State = 0
	QScroller__Pressed QScroller__State = 1
	QScroller__Dragging QScroller__State = 2
	QScroller__Scrolling QScroller__State = 3
)

type QScroller__ScrollerGestureType int64

const (
	QScroller__TouchGesture QScroller__ScrollerGestureType = 0
	QScroller__LeftMouseButtonGesture QScroller__ScrollerGestureType = 1
	QScroller__RightMouseButtonGesture QScroller__ScrollerGestureType = 2
	QScroller__MiddleMouseButtonGesture QScroller__ScrollerGestureType = 3
)

type QScroller__Input int64

const (
	QScroller__InputPress QScroller__Input = 1
	QScroller__InputMove QScroller__Input = 2
	QScroller__InputRelease QScroller__Input = 3
)

