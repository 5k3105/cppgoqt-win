package qt

type QInternal__PaintDeviceFlags int64

const (
	QInternal__UnknownDevice QInternal__PaintDeviceFlags = 0
	QInternal__Widget QInternal__PaintDeviceFlags = 1
	QInternal__Pixmap QInternal__PaintDeviceFlags = 2
	QInternal__Image QInternal__PaintDeviceFlags = 3
	QInternal__Printer QInternal__PaintDeviceFlags = 4
	QInternal__Picture QInternal__PaintDeviceFlags = 5
	QInternal__Pbuffer QInternal__PaintDeviceFlags = 6
	QInternal__FramebufferObject QInternal__PaintDeviceFlags = 7
	QInternal__CustomRaster QInternal__PaintDeviceFlags = 8
	QInternal__PaintBuffer QInternal__PaintDeviceFlags = 10
	QInternal__OpenGL QInternal__PaintDeviceFlags = 11
)

type QInternal__RelayoutType int64

const (
	QInternal__RelayoutNormal QInternal__RelayoutType = 0
	QInternal__RelayoutDragging QInternal__RelayoutType = 1
	QInternal__RelayoutDropped QInternal__RelayoutType = 2
)

type QInternal__DockPosition int64

const (
	QInternal__LeftDock QInternal__DockPosition = 0
	QInternal__RightDock QInternal__DockPosition = 1
	QInternal__TopDock QInternal__DockPosition = 2
	QInternal__BottomDock QInternal__DockPosition = 3
	QInternal__DockCount QInternal__DockPosition = 4
)

type QInternal__CallMode int64

const (
	QInternal__Testing QInternal__CallMode = 0
	QInternal__Live QInternal__CallMode = 1
)

type QInternal__SaveStateRule int64

const (
	QInternal__KeepSavedState QInternal__SaveStateRule = 0
	QInternal__ClearSavedState QInternal__SaveStateRule = 1
)

type QInternal__Callback int64

const (
	QInternal__EventNotifyCallback QInternal__Callback = 0
	QInternal__LastCallback QInternal__Callback = 1
)

