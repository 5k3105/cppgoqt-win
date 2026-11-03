package qt

type QSizePolicy__PolicyFlag int64

const (
	QSizePolicy__GrowFlag QSizePolicy__PolicyFlag = 1
	QSizePolicy__ExpandFlag QSizePolicy__PolicyFlag = 2
	QSizePolicy__ShrinkFlag QSizePolicy__PolicyFlag = 4
	QSizePolicy__IgnoreFlag QSizePolicy__PolicyFlag = 8
)

type QSizePolicy__Policy int64

const (
	QSizePolicy__Fixed QSizePolicy__Policy = 0
	QSizePolicy__Minimum QSizePolicy__Policy = 1
	QSizePolicy__Maximum QSizePolicy__Policy = 4
	QSizePolicy__Preferred QSizePolicy__Policy = 5
	QSizePolicy__MinimumExpanding QSizePolicy__Policy = 3
	QSizePolicy__Expanding QSizePolicy__Policy = 7
	QSizePolicy__Ignored QSizePolicy__Policy = 13
)

type QSizePolicy__ControlType int64

const (
	QSizePolicy__DefaultType QSizePolicy__ControlType = 1
	QSizePolicy__ButtonBox QSizePolicy__ControlType = 2
	QSizePolicy__CheckBox QSizePolicy__ControlType = 4
	QSizePolicy__ComboBox QSizePolicy__ControlType = 8
	QSizePolicy__Frame QSizePolicy__ControlType = 16
	QSizePolicy__GroupBox QSizePolicy__ControlType = 32
	QSizePolicy__Label QSizePolicy__ControlType = 64
	QSizePolicy__Line QSizePolicy__ControlType = 128
	QSizePolicy__LineEdit QSizePolicy__ControlType = 256
	QSizePolicy__PushButton QSizePolicy__ControlType = 512
	QSizePolicy__RadioButton QSizePolicy__ControlType = 1024
	QSizePolicy__Slider QSizePolicy__ControlType = 2048
	QSizePolicy__SpinBox QSizePolicy__ControlType = 4096
	QSizePolicy__TabWidget QSizePolicy__ControlType = 8192
	QSizePolicy__ToolButton QSizePolicy__ControlType = 16384
)

