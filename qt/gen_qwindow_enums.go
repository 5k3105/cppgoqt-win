package qt

type QWindow__Visibility int64

const (
	QWindow__Hidden QWindow__Visibility = 0
	QWindow__AutomaticVisibility QWindow__Visibility = 1
	QWindow__Windowed QWindow__Visibility = 2
	QWindow__Minimized QWindow__Visibility = 3
	QWindow__Maximized QWindow__Visibility = 4
	QWindow__FullScreen QWindow__Visibility = 5
)

type QWindow__AncestorMode int64

const (
	QWindow__ExcludeTransients QWindow__AncestorMode = 0
	QWindow__IncludeTransients QWindow__AncestorMode = 1
)

