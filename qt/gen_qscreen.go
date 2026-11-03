package qt

import (
	"unsafe"
	"math"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QScreen_base func(uintptr) uintptr
	addr_QScreen_tr uintptr
	addr_QScreen_name uintptr
	addr_QScreen_manufacturer uintptr
	addr_QScreen_model uintptr
	addr_QScreen_serialNumber uintptr
	addr_QScreen_depth uintptr
	addr_QScreen_size uintptr
	addr_QScreen_geometry uintptr
	addr_QScreen_physicalSize uintptr
	pg_QScreen_physicalDotsPerInchX func(uintptr) float64
	pg_QScreen_physicalDotsPerInchY func(uintptr) float64
	pg_QScreen_physicalDotsPerInch func(uintptr) float64
	pg_QScreen_logicalDotsPerInchX func(uintptr) float64
	pg_QScreen_logicalDotsPerInchY func(uintptr) float64
	pg_QScreen_logicalDotsPerInch func(uintptr) float64
	pg_QScreen_devicePixelRatio func(uintptr) float64
	addr_QScreen_availableSize uintptr
	addr_QScreen_availableGeometry uintptr
	addr_QScreen_virtualSiblings uintptr
	addr_QScreen_virtualSiblingAt uintptr
	addr_QScreen_virtualSize uintptr
	addr_QScreen_virtualGeometry uintptr
	addr_QScreen_availableVirtualSize uintptr
	addr_QScreen_availableVirtualGeometry uintptr
	addr_QScreen_primaryOrientation uintptr
	addr_QScreen_orientation uintptr
	addr_QScreen_nativeOrientation uintptr
	addr_QScreen_angleBetween uintptr
	addr_QScreen_transformBetween uintptr
	addr_QScreen_mapBetween uintptr
	addr_QScreen_isPortrait uintptr
	addr_QScreen_isLandscape uintptr
	addr_QScreen_grabWindow uintptr
	pg_QScreen_refreshRate func(uintptr) float64
	addr_QScreen_geometryChanged uintptr
	addr_QScreen_availableGeometryChanged uintptr
	addr_QScreen_physicalSizeChanged uintptr
	pg_QScreen_physicalDotsPerInchChanged func(uintptr, float64)
	pg_QScreen_logicalDotsPerInchChanged func(uintptr, float64)
	addr_QScreen_virtualGeometryChanged uintptr
	addr_QScreen_primaryOrientationChanged uintptr
	addr_QScreen_orientationChanged uintptr
	pg_QScreen_refreshRateChanged func(uintptr, float64)
	addr_QScreen_connect uintptr
	addr_QScreen_connect2 uintptr
	addr_QScreen_tr2 uintptr
	addr_QScreen_tr3 uintptr
	addr_QScreen_grabWindow2 uintptr
	addr_QScreen_grabWindow3 uintptr
	addr_QScreen_grabWindow4 uintptr
	addr_QScreen_grabWindow5 uintptr
	addr_QScreen_grabWindow6 uintptr
	addr_QScreen_connect3 uintptr
	addr_QScreen_connect4 uintptr
	pg_QScreen_connect_geometryChanged func(uintptr, int64)
	pg_QScreen_connect_availableGeometryChanged func(uintptr, int64)
	pg_QScreen_connect_physicalSizeChanged func(uintptr, int64)
	pg_QScreen_connect_physicalDotsPerInchChanged func(uintptr, int64)
	pg_QScreen_connect_logicalDotsPerInchChanged func(uintptr, int64)
	pg_QScreen_connect_virtualGeometryChanged func(uintptr, int64)
	pg_QScreen_connect_primaryOrientationChanged func(uintptr, int64)
	pg_QScreen_connect_orientationChanged func(uintptr, int64)
	pg_QScreen_connect_refreshRateChanged func(uintptr, int64)
	pg_QScreen_connect_destroyed func(uintptr, int64)
	pg_QScreen_connect_destroyed2 func(uintptr, int64)
	pg_QScreen_connect_objectNameChanged func(uintptr, int64)
	pg_QScreen__disconnect func(int64)
	pg_QScreen_delete func(uintptr)
)

func init() {
	reg(&pg_QScreen_base, "QScreen_base")
	symaddr(&addr_QScreen_tr, "QScreen_tr")
	symaddr(&addr_QScreen_name, "QScreen_name")
	symaddr(&addr_QScreen_manufacturer, "QScreen_manufacturer")
	symaddr(&addr_QScreen_model, "QScreen_model")
	symaddr(&addr_QScreen_serialNumber, "QScreen_serialNumber")
	symaddr(&addr_QScreen_depth, "QScreen_depth")
	symaddr(&addr_QScreen_size, "QScreen_size")
	symaddr(&addr_QScreen_geometry, "QScreen_geometry")
	symaddr(&addr_QScreen_physicalSize, "QScreen_physicalSize")
	reg(&pg_QScreen_physicalDotsPerInchX, "QScreen_physicalDotsPerInchX")
	reg(&pg_QScreen_physicalDotsPerInchY, "QScreen_physicalDotsPerInchY")
	reg(&pg_QScreen_physicalDotsPerInch, "QScreen_physicalDotsPerInch")
	reg(&pg_QScreen_logicalDotsPerInchX, "QScreen_logicalDotsPerInchX")
	reg(&pg_QScreen_logicalDotsPerInchY, "QScreen_logicalDotsPerInchY")
	reg(&pg_QScreen_logicalDotsPerInch, "QScreen_logicalDotsPerInch")
	reg(&pg_QScreen_devicePixelRatio, "QScreen_devicePixelRatio")
	symaddr(&addr_QScreen_availableSize, "QScreen_availableSize")
	symaddr(&addr_QScreen_availableGeometry, "QScreen_availableGeometry")
	symaddr(&addr_QScreen_virtualSiblings, "QScreen_virtualSiblings")
	symaddr(&addr_QScreen_virtualSiblingAt, "QScreen_virtualSiblingAt")
	symaddr(&addr_QScreen_virtualSize, "QScreen_virtualSize")
	symaddr(&addr_QScreen_virtualGeometry, "QScreen_virtualGeometry")
	symaddr(&addr_QScreen_availableVirtualSize, "QScreen_availableVirtualSize")
	symaddr(&addr_QScreen_availableVirtualGeometry, "QScreen_availableVirtualGeometry")
	symaddr(&addr_QScreen_primaryOrientation, "QScreen_primaryOrientation")
	symaddr(&addr_QScreen_orientation, "QScreen_orientation")
	symaddr(&addr_QScreen_nativeOrientation, "QScreen_nativeOrientation")
	symaddr(&addr_QScreen_angleBetween, "QScreen_angleBetween")
	symaddr(&addr_QScreen_transformBetween, "QScreen_transformBetween")
	symaddr(&addr_QScreen_mapBetween, "QScreen_mapBetween")
	symaddr(&addr_QScreen_isPortrait, "QScreen_isPortrait")
	symaddr(&addr_QScreen_isLandscape, "QScreen_isLandscape")
	symaddr(&addr_QScreen_grabWindow, "QScreen_grabWindow")
	reg(&pg_QScreen_refreshRate, "QScreen_refreshRate")
	symaddr(&addr_QScreen_geometryChanged, "QScreen_geometryChanged")
	symaddr(&addr_QScreen_availableGeometryChanged, "QScreen_availableGeometryChanged")
	symaddr(&addr_QScreen_physicalSizeChanged, "QScreen_physicalSizeChanged")
	reg(&pg_QScreen_physicalDotsPerInchChanged, "QScreen_physicalDotsPerInchChanged")
	reg(&pg_QScreen_logicalDotsPerInchChanged, "QScreen_logicalDotsPerInchChanged")
	symaddr(&addr_QScreen_virtualGeometryChanged, "QScreen_virtualGeometryChanged")
	symaddr(&addr_QScreen_primaryOrientationChanged, "QScreen_primaryOrientationChanged")
	symaddr(&addr_QScreen_orientationChanged, "QScreen_orientationChanged")
	reg(&pg_QScreen_refreshRateChanged, "QScreen_refreshRateChanged")
	symaddr(&addr_QScreen_connect, "QScreen_connect")
	symaddr(&addr_QScreen_connect2, "QScreen_connect2")
	symaddr(&addr_QScreen_tr2, "QScreen_tr2")
	symaddr(&addr_QScreen_tr3, "QScreen_tr3")
	symaddr(&addr_QScreen_grabWindow2, "QScreen_grabWindow2")
	symaddr(&addr_QScreen_grabWindow3, "QScreen_grabWindow3")
	symaddr(&addr_QScreen_grabWindow4, "QScreen_grabWindow4")
	symaddr(&addr_QScreen_grabWindow5, "QScreen_grabWindow5")
	symaddr(&addr_QScreen_grabWindow6, "QScreen_grabWindow6")
	symaddr(&addr_QScreen_connect3, "QScreen_connect3")
	symaddr(&addr_QScreen_connect4, "QScreen_connect4")
	reg(&pg_QScreen_connect_geometryChanged, "QScreen_connect_geometryChanged")
	reg(&pg_QScreen_connect_availableGeometryChanged, "QScreen_connect_availableGeometryChanged")
	reg(&pg_QScreen_connect_physicalSizeChanged, "QScreen_connect_physicalSizeChanged")
	reg(&pg_QScreen_connect_physicalDotsPerInchChanged, "QScreen_connect_physicalDotsPerInchChanged")
	reg(&pg_QScreen_connect_logicalDotsPerInchChanged, "QScreen_connect_logicalDotsPerInchChanged")
	reg(&pg_QScreen_connect_virtualGeometryChanged, "QScreen_connect_virtualGeometryChanged")
	reg(&pg_QScreen_connect_primaryOrientationChanged, "QScreen_connect_primaryOrientationChanged")
	reg(&pg_QScreen_connect_orientationChanged, "QScreen_connect_orientationChanged")
	reg(&pg_QScreen_connect_refreshRateChanged, "QScreen_connect_refreshRateChanged")
	reg(&pg_QScreen_connect_destroyed, "QScreen_connect_destroyed")
	reg(&pg_QScreen_connect_destroyed2, "QScreen_connect_destroyed2")
	reg(&pg_QScreen_connect_objectNameChanged, "QScreen_connect_objectNameChanged")
	reg(&pg_QScreen__disconnect, "QScreen__disconnect")
	reg(&pg_QScreen_delete, "QScreen_delete")
}

type QScreen struct {
	h uintptr
	*QObject
}

func bindQScreen(h uintptr) *QScreen {
	if h == 0 {
		return nil
	}
	return &QScreen{h: h, QObject: bindQObject(pg_QScreen_base(h))}
}

func newownedQScreen(h uintptr) *QScreen { return bindQScreen(h) }

func (this *QScreen) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QScreen) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQScreen(h unsafe.Pointer) *QScreen { return bindQScreen(uintptr(h)) }
func UnsafeNewQScreenOwned(h unsafe.Pointer) *QScreen { return newownedQScreen(uintptr(h)) }

func (this *QScreen) UnsafeRelease() {
	this.h = 0
	if this.QObject != nil {
		this.QObject.UnsafeRelease()
	}
}

func QScreen_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QScreen_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QScreen) Name() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QScreen_name, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QScreen) Manufacturer() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QScreen_manufacturer, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QScreen) Model() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QScreen_model, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QScreen) SerialNumber() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QScreen_serialNumber, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QScreen) Depth() int {
	__r0, _, _ := purego.SyscallN(addr_QScreen_depth, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QScreen) Size() *QSize {
	__r, _, _ := purego.SyscallN(addr_QScreen_size, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QScreen) Geometry() *QRect {
	__r, _, _ := purego.SyscallN(addr_QScreen_geometry, this.h)
	runtime.KeepAlive(this)
	return newownedQRect(__r)
}

func (this *QScreen) PhysicalSize() *QSizeF {
	__r, _, _ := purego.SyscallN(addr_QScreen_physicalSize, this.h)
	runtime.KeepAlive(this)
	return newownedQSizeF(__r)
}

func (this *QScreen) PhysicalDotsPerInchX() float64 {
	__r := pg_QScreen_physicalDotsPerInchX(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QScreen) PhysicalDotsPerInchY() float64 {
	__r := pg_QScreen_physicalDotsPerInchY(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QScreen) PhysicalDotsPerInch() float64 {
	__r := pg_QScreen_physicalDotsPerInch(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QScreen) LogicalDotsPerInchX() float64 {
	__r := pg_QScreen_logicalDotsPerInchX(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QScreen) LogicalDotsPerInchY() float64 {
	__r := pg_QScreen_logicalDotsPerInchY(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QScreen) LogicalDotsPerInch() float64 {
	__r := pg_QScreen_logicalDotsPerInch(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QScreen) DevicePixelRatio() float64 {
	__r := pg_QScreen_devicePixelRatio(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QScreen) AvailableSize() *QSize {
	__r, _, _ := purego.SyscallN(addr_QScreen_availableSize, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QScreen) AvailableGeometry() *QRect {
	__r, _, _ := purego.SyscallN(addr_QScreen_availableGeometry, this.h)
	runtime.KeepAlive(this)
	return newownedQRect(__r)
}

func (this *QScreen) VirtualSiblings() []*QScreen {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QScreen_virtualSiblings, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QScreen, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = bindQScreen(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QScreen) VirtualSiblingAt(point *QPoint) *QScreen {
	__r, _, _ := purego.SyscallN(addr_QScreen_virtualSiblingAt, this.h, uintptr(point.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(point)
	return bindQScreen(__r)
}

func (this *QScreen) VirtualSize() *QSize {
	__r, _, _ := purego.SyscallN(addr_QScreen_virtualSize, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QScreen) VirtualGeometry() *QRect {
	__r, _, _ := purego.SyscallN(addr_QScreen_virtualGeometry, this.h)
	runtime.KeepAlive(this)
	return newownedQRect(__r)
}

func (this *QScreen) AvailableVirtualSize() *QSize {
	__r, _, _ := purego.SyscallN(addr_QScreen_availableVirtualSize, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QScreen) AvailableVirtualGeometry() *QRect {
	__r, _, _ := purego.SyscallN(addr_QScreen_availableVirtualGeometry, this.h)
	runtime.KeepAlive(this)
	return newownedQRect(__r)
}

func (this *QScreen) PrimaryOrientation() ScreenOrientation {
	__r0, _, _ := purego.SyscallN(addr_QScreen_primaryOrientation, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return ScreenOrientation(__r)
}

func (this *QScreen) Orientation() ScreenOrientation {
	__r0, _, _ := purego.SyscallN(addr_QScreen_orientation, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return ScreenOrientation(__r)
}

func (this *QScreen) NativeOrientation() ScreenOrientation {
	__r0, _, _ := purego.SyscallN(addr_QScreen_nativeOrientation, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return ScreenOrientation(__r)
}

func (this *QScreen) AngleBetween(a ScreenOrientation, b ScreenOrientation) int {
	__r0, _, _ := purego.SyscallN(addr_QScreen_angleBetween, this.h, uintptr(a), uintptr(b))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QScreen) TransformBetween(a ScreenOrientation, b ScreenOrientation, target *QRect) *QTransform {
	__r, _, _ := purego.SyscallN(addr_QScreen_transformBetween, this.h, uintptr(a), uintptr(b), uintptr(target.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(target)
	return newownedQTransform(__r)
}

func (this *QScreen) MapBetween(a ScreenOrientation, b ScreenOrientation, rect *QRect) *QRect {
	__r, _, _ := purego.SyscallN(addr_QScreen_mapBetween, this.h, uintptr(a), uintptr(b), uintptr(rect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
	return newownedQRect(__r)
}

func (this *QScreen) IsPortrait(orientation ScreenOrientation) bool {
	__r0, _, _ := purego.SyscallN(addr_QScreen_isPortrait, this.h, uintptr(orientation))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QScreen) IsLandscape(orientation ScreenOrientation) bool {
	__r0, _, _ := purego.SyscallN(addr_QScreen_isLandscape, this.h, uintptr(orientation))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QScreen) GrabWindow() *QPixmap {
	__r, _, _ := purego.SyscallN(addr_QScreen_grabWindow, this.h)
	runtime.KeepAlive(this)
	return newownedQPixmap(__r)
}

func (this *QScreen) RefreshRate() float64 {
	__r := pg_QScreen_refreshRate(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QScreen) GeometryChanged(geometry *QRect) {
	purego.SyscallN(addr_QScreen_geometryChanged, this.h, uintptr(geometry.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(geometry)
}

func (this *QScreen) AvailableGeometryChanged(geometry *QRect) {
	purego.SyscallN(addr_QScreen_availableGeometryChanged, this.h, uintptr(geometry.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(geometry)
}

func (this *QScreen) PhysicalSizeChanged(size *QSizeF) {
	purego.SyscallN(addr_QScreen_physicalSizeChanged, this.h, uintptr(size.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(size)
}

func (this *QScreen) PhysicalDotsPerInchChanged(dpi float64) {
	pg_QScreen_physicalDotsPerInchChanged(this.h, dpi)
	runtime.KeepAlive(this)
}

func (this *QScreen) LogicalDotsPerInchChanged(dpi float64) {
	pg_QScreen_logicalDotsPerInchChanged(this.h, dpi)
	runtime.KeepAlive(this)
}

func (this *QScreen) VirtualGeometryChanged(rect *QRect) {
	purego.SyscallN(addr_QScreen_virtualGeometryChanged, this.h, uintptr(rect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
}

func (this *QScreen) PrimaryOrientationChanged(orientation ScreenOrientation) {
	purego.SyscallN(addr_QScreen_primaryOrientationChanged, this.h, uintptr(orientation))
	runtime.KeepAlive(this)
}

func (this *QScreen) OrientationChanged(orientation ScreenOrientation) {
	purego.SyscallN(addr_QScreen_orientationChanged, this.h, uintptr(orientation))
	runtime.KeepAlive(this)
}

func (this *QScreen) RefreshRateChanged(refreshRate float64) {
	pg_QScreen_refreshRateChanged(this.h, refreshRate)
	runtime.KeepAlive(this)
}

func QScreen_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QScreen_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QScreen_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QScreen_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func QScreen_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QScreen_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QScreen_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QScreen_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func (this *QScreen) GrabWindowWithWindow(window uint64) *QPixmap {
	__r, _, _ := purego.SyscallN(addr_QScreen_grabWindow2, this.h, uintptr(window))
	runtime.KeepAlive(this)
	return newownedQPixmap(__r)
}

func (this *QScreen) GrabWindow2(window uint64, x int) *QPixmap {
	__r, _, _ := purego.SyscallN(addr_QScreen_grabWindow3, this.h, uintptr(window), uintptr(x))
	runtime.KeepAlive(this)
	return newownedQPixmap(__r)
}

func (this *QScreen) GrabWindow3(window uint64, x int, y int) *QPixmap {
	__r, _, _ := purego.SyscallN(addr_QScreen_grabWindow4, this.h, uintptr(window), uintptr(x), uintptr(y))
	runtime.KeepAlive(this)
	return newownedQPixmap(__r)
}

func (this *QScreen) GrabWindow4(window uint64, x int, y int, w int) *QPixmap {
	__r, _, _ := purego.SyscallN(addr_QScreen_grabWindow5, this.h, uintptr(window), uintptr(x), uintptr(y), uintptr(w))
	runtime.KeepAlive(this)
	return newownedQPixmap(__r)
}

func (this *QScreen) GrabWindow5(window uint64, x int, y int, w int, h int) *QPixmap {
	__r, _, _ := purego.SyscallN(addr_QScreen_grabWindow6, this.h, uintptr(window), uintptr(x), uintptr(y), uintptr(w), uintptr(h))
	runtime.KeepAlive(this)
	return newownedQPixmap(__r)
}

func QScreen_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QScreen_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QScreen_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QScreen_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QScreen) OnGeometryChanged(fn func(*QRect)) int64 {
	id := registerCallback(func(geometry uintptr) { fn(UnsafeNewQRect(unsafe.Pointer(geometry))) })
	pg_QScreen_connect_geometryChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QScreen) OnAvailableGeometryChanged(fn func(*QRect)) int64 {
	id := registerCallback(func(geometry uintptr) { fn(UnsafeNewQRect(unsafe.Pointer(geometry))) })
	pg_QScreen_connect_availableGeometryChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QScreen) OnPhysicalSizeChanged(fn func(*QSizeF)) int64 {
	id := registerCallback(func(size uintptr) { fn(UnsafeNewQSizeF(unsafe.Pointer(size))) })
	pg_QScreen_connect_physicalSizeChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QScreen) OnPhysicalDotsPerInchChanged(fn func(float64)) int64 {
	id := registerCallback(func(dpi uintptr) { fn(math.Float64frombits(uint64(dpi))) })
	pg_QScreen_connect_physicalDotsPerInchChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QScreen) OnLogicalDotsPerInchChanged(fn func(float64)) int64 {
	id := registerCallback(func(dpi uintptr) { fn(math.Float64frombits(uint64(dpi))) })
	pg_QScreen_connect_logicalDotsPerInchChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QScreen) OnVirtualGeometryChanged(fn func(*QRect)) int64 {
	id := registerCallback(func(rect uintptr) { fn(UnsafeNewQRect(unsafe.Pointer(rect))) })
	pg_QScreen_connect_virtualGeometryChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QScreen) OnPrimaryOrientationChanged(fn func(ScreenOrientation)) int64 {
	id := registerCallback(func(orientation int32) { fn(ScreenOrientation(orientation)) })
	pg_QScreen_connect_primaryOrientationChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QScreen) OnOrientationChanged(fn func(ScreenOrientation)) int64 {
	id := registerCallback(func(orientation int32) { fn(ScreenOrientation(orientation)) })
	pg_QScreen_connect_orientationChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QScreen) OnRefreshRateChanged(fn func(float64)) int64 {
	id := registerCallback(func(refreshRate uintptr) { fn(math.Float64frombits(uint64(refreshRate))) })
	pg_QScreen_connect_refreshRateChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QScreen) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QScreen_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QScreen) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QScreen_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QScreen) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QScreen_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QScreen) Disconnect(handle int64) {
	pg_QScreen__disconnect(handle)
	unregisterCallback(handle)
}

func (this *QScreen) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QScreen_delete(this.h)
	}
	this.UnsafeRelease()
}
