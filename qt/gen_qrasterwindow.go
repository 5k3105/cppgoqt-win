package qt

import (
	"unsafe"
	"math"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QRasterWindow_chain uintptr
	addr_QRasterWindow_new uintptr
	addr_QRasterWindow_new2 uintptr
	addr_QRasterWindow_tr uintptr
	addr_QRasterWindow_fromWinId uintptr
	addr_QRasterWindow_connect uintptr
	addr_QRasterWindow_connect2 uintptr
	addr_QRasterWindow_sender uintptr
	addr_QRasterWindow_senderSignalIndex uintptr
	addr_QRasterWindow_receivers uintptr
	addr_QRasterWindow_isSignalConnected uintptr
	pg_QRasterWindow_devicePixelRatioFScale func() float64
	pg_QRasterWindow_encodeMetricF func(int32, float64) int32
	pg_QRasterWindow_getDecodedMetricF func(uintptr, int32, int32) float64
	addr_QRasterWindow_tr2 uintptr
	addr_QRasterWindow_tr3 uintptr
	addr_QRasterWindow_connect3 uintptr
	addr_QRasterWindow_connect4 uintptr
	addr_QRasterWindow_metaObject_default uintptr
	addr_QRasterWindow_metric_default uintptr
	addr_QRasterWindow_redirected_default uintptr
	addr_QRasterWindow_resizeEvent_default uintptr
	addr_QRasterWindow_exposeEvent_default uintptr
	addr_QRasterWindow_paintEvent_default uintptr
	addr_QRasterWindow_event_default uintptr
	addr_QRasterWindow_surfaceType_default uintptr
	addr_QRasterWindow_format_default uintptr
	addr_QRasterWindow_size_default uintptr
	addr_QRasterWindow_accessibleRoot_default uintptr
	addr_QRasterWindow_focusObject_default uintptr
	addr_QRasterWindow_moveEvent_default uintptr
	addr_QRasterWindow_focusInEvent_default uintptr
	addr_QRasterWindow_focusOutEvent_default uintptr
	addr_QRasterWindow_showEvent_default uintptr
	addr_QRasterWindow_hideEvent_default uintptr
	addr_QRasterWindow_closeEvent_default uintptr
	addr_QRasterWindow_keyPressEvent_default uintptr
	addr_QRasterWindow_keyReleaseEvent_default uintptr
	addr_QRasterWindow_mousePressEvent_default uintptr
	addr_QRasterWindow_mouseReleaseEvent_default uintptr
	addr_QRasterWindow_mouseDoubleClickEvent_default uintptr
	addr_QRasterWindow_mouseMoveEvent_default uintptr
	addr_QRasterWindow_wheelEvent_default uintptr
	addr_QRasterWindow_touchEvent_default uintptr
	addr_QRasterWindow_tabletEvent_default uintptr
	addr_QRasterWindow_eventFilter_default uintptr
	addr_QRasterWindow_timerEvent_default uintptr
	addr_QRasterWindow_childEvent_default uintptr
	addr_QRasterWindow_customEvent_default uintptr
	addr_QRasterWindow_connectNotify_default uintptr
	addr_QRasterWindow_disconnectNotify_default uintptr
	addr_QRasterWindow_devType_default uintptr
	addr_QRasterWindow_initPainter_default uintptr
	addr_QRasterWindow_sharedPainter_default uintptr
	pg_QRasterWindow_connect_screenChanged func(uintptr, int64)
	pg_QRasterWindow_connect_modalityChanged func(uintptr, int64)
	pg_QRasterWindow_connect_flagsChanged func(uintptr, int64)
	pg_QRasterWindow_connect_windowStateChanged func(uintptr, int64)
	pg_QRasterWindow_connect_windowTitleChanged func(uintptr, int64)
	pg_QRasterWindow_connect_xChanged func(uintptr, int64)
	pg_QRasterWindow_connect_yChanged func(uintptr, int64)
	pg_QRasterWindow_connect_widthChanged func(uintptr, int64)
	pg_QRasterWindow_connect_heightChanged func(uintptr, int64)
	pg_QRasterWindow_connect_minimumWidthChanged func(uintptr, int64)
	pg_QRasterWindow_connect_minimumHeightChanged func(uintptr, int64)
	pg_QRasterWindow_connect_maximumWidthChanged func(uintptr, int64)
	pg_QRasterWindow_connect_maximumHeightChanged func(uintptr, int64)
	pg_QRasterWindow_connect_safeAreaMarginsChanged func(uintptr, int64)
	pg_QRasterWindow_connect_visibleChanged func(uintptr, int64)
	pg_QRasterWindow_connect_visibilityChanged func(uintptr, int64)
	pg_QRasterWindow_connect_activeChanged func(uintptr, int64)
	pg_QRasterWindow_connect_contentOrientationChanged func(uintptr, int64)
	pg_QRasterWindow_connect_focusObjectChanged func(uintptr, int64)
	pg_QRasterWindow_connect_opacityChanged func(uintptr, int64)
	pg_QRasterWindow_connect_transientParentChanged func(uintptr, int64)
	pg_QRasterWindow_connect_destroyed func(uintptr, int64)
	pg_QRasterWindow_connect_destroyed2 func(uintptr, int64)
	pg_QRasterWindow_connect_objectNameChanged func(uintptr, int64)
	pg_QRasterWindow__disconnect func(int64)
	pg_QRasterWindow__set_override func(uintptr, int32)
	pg_QRasterWindow_delete func(uintptr)
)

func init() {
	symaddr(&addr_QRasterWindow_chain, "QRasterWindow_chain")
	symaddr(&addr_QRasterWindow_new, "QRasterWindow_new")
	symaddr(&addr_QRasterWindow_new2, "QRasterWindow_new2")
	symaddr(&addr_QRasterWindow_tr, "QRasterWindow_tr")
	symaddr(&addr_QRasterWindow_fromWinId, "QRasterWindow_fromWinId")
	symaddr(&addr_QRasterWindow_connect, "QRasterWindow_connect")
	symaddr(&addr_QRasterWindow_connect2, "QRasterWindow_connect2")
	symaddr(&addr_QRasterWindow_sender, "QRasterWindow_sender")
	symaddr(&addr_QRasterWindow_senderSignalIndex, "QRasterWindow_senderSignalIndex")
	symaddr(&addr_QRasterWindow_receivers, "QRasterWindow_receivers")
	symaddr(&addr_QRasterWindow_isSignalConnected, "QRasterWindow_isSignalConnected")
	reg(&pg_QRasterWindow_devicePixelRatioFScale, "QRasterWindow_devicePixelRatioFScale")
	reg(&pg_QRasterWindow_encodeMetricF, "QRasterWindow_encodeMetricF")
	reg(&pg_QRasterWindow_getDecodedMetricF, "QRasterWindow_getDecodedMetricF")
	symaddr(&addr_QRasterWindow_tr2, "QRasterWindow_tr2")
	symaddr(&addr_QRasterWindow_tr3, "QRasterWindow_tr3")
	symaddr(&addr_QRasterWindow_connect3, "QRasterWindow_connect3")
	symaddr(&addr_QRasterWindow_connect4, "QRasterWindow_connect4")
	symaddr(&addr_QRasterWindow_metaObject_default, "QRasterWindow_metaObject_default")
	symaddr(&addr_QRasterWindow_metric_default, "QRasterWindow_metric_default")
	symaddr(&addr_QRasterWindow_redirected_default, "QRasterWindow_redirected_default")
	symaddr(&addr_QRasterWindow_resizeEvent_default, "QRasterWindow_resizeEvent_default")
	symaddr(&addr_QRasterWindow_exposeEvent_default, "QRasterWindow_exposeEvent_default")
	symaddr(&addr_QRasterWindow_paintEvent_default, "QRasterWindow_paintEvent_default")
	symaddr(&addr_QRasterWindow_event_default, "QRasterWindow_event_default")
	symaddr(&addr_QRasterWindow_surfaceType_default, "QRasterWindow_surfaceType_default")
	symaddr(&addr_QRasterWindow_format_default, "QRasterWindow_format_default")
	symaddr(&addr_QRasterWindow_size_default, "QRasterWindow_size_default")
	symaddr(&addr_QRasterWindow_accessibleRoot_default, "QRasterWindow_accessibleRoot_default")
	symaddr(&addr_QRasterWindow_focusObject_default, "QRasterWindow_focusObject_default")
	symaddr(&addr_QRasterWindow_moveEvent_default, "QRasterWindow_moveEvent_default")
	symaddr(&addr_QRasterWindow_focusInEvent_default, "QRasterWindow_focusInEvent_default")
	symaddr(&addr_QRasterWindow_focusOutEvent_default, "QRasterWindow_focusOutEvent_default")
	symaddr(&addr_QRasterWindow_showEvent_default, "QRasterWindow_showEvent_default")
	symaddr(&addr_QRasterWindow_hideEvent_default, "QRasterWindow_hideEvent_default")
	symaddr(&addr_QRasterWindow_closeEvent_default, "QRasterWindow_closeEvent_default")
	symaddr(&addr_QRasterWindow_keyPressEvent_default, "QRasterWindow_keyPressEvent_default")
	symaddr(&addr_QRasterWindow_keyReleaseEvent_default, "QRasterWindow_keyReleaseEvent_default")
	symaddr(&addr_QRasterWindow_mousePressEvent_default, "QRasterWindow_mousePressEvent_default")
	symaddr(&addr_QRasterWindow_mouseReleaseEvent_default, "QRasterWindow_mouseReleaseEvent_default")
	symaddr(&addr_QRasterWindow_mouseDoubleClickEvent_default, "QRasterWindow_mouseDoubleClickEvent_default")
	symaddr(&addr_QRasterWindow_mouseMoveEvent_default, "QRasterWindow_mouseMoveEvent_default")
	symaddr(&addr_QRasterWindow_wheelEvent_default, "QRasterWindow_wheelEvent_default")
	symaddr(&addr_QRasterWindow_touchEvent_default, "QRasterWindow_touchEvent_default")
	symaddr(&addr_QRasterWindow_tabletEvent_default, "QRasterWindow_tabletEvent_default")
	symaddr(&addr_QRasterWindow_eventFilter_default, "QRasterWindow_eventFilter_default")
	symaddr(&addr_QRasterWindow_timerEvent_default, "QRasterWindow_timerEvent_default")
	symaddr(&addr_QRasterWindow_childEvent_default, "QRasterWindow_childEvent_default")
	symaddr(&addr_QRasterWindow_customEvent_default, "QRasterWindow_customEvent_default")
	symaddr(&addr_QRasterWindow_connectNotify_default, "QRasterWindow_connectNotify_default")
	symaddr(&addr_QRasterWindow_disconnectNotify_default, "QRasterWindow_disconnectNotify_default")
	symaddr(&addr_QRasterWindow_devType_default, "QRasterWindow_devType_default")
	symaddr(&addr_QRasterWindow_initPainter_default, "QRasterWindow_initPainter_default")
	symaddr(&addr_QRasterWindow_sharedPainter_default, "QRasterWindow_sharedPainter_default")
	reg(&pg_QRasterWindow_connect_screenChanged, "QRasterWindow_connect_screenChanged")
	reg(&pg_QRasterWindow_connect_modalityChanged, "QRasterWindow_connect_modalityChanged")
	reg(&pg_QRasterWindow_connect_flagsChanged, "QRasterWindow_connect_flagsChanged")
	reg(&pg_QRasterWindow_connect_windowStateChanged, "QRasterWindow_connect_windowStateChanged")
	reg(&pg_QRasterWindow_connect_windowTitleChanged, "QRasterWindow_connect_windowTitleChanged")
	reg(&pg_QRasterWindow_connect_xChanged, "QRasterWindow_connect_xChanged")
	reg(&pg_QRasterWindow_connect_yChanged, "QRasterWindow_connect_yChanged")
	reg(&pg_QRasterWindow_connect_widthChanged, "QRasterWindow_connect_widthChanged")
	reg(&pg_QRasterWindow_connect_heightChanged, "QRasterWindow_connect_heightChanged")
	reg(&pg_QRasterWindow_connect_minimumWidthChanged, "QRasterWindow_connect_minimumWidthChanged")
	reg(&pg_QRasterWindow_connect_minimumHeightChanged, "QRasterWindow_connect_minimumHeightChanged")
	reg(&pg_QRasterWindow_connect_maximumWidthChanged, "QRasterWindow_connect_maximumWidthChanged")
	reg(&pg_QRasterWindow_connect_maximumHeightChanged, "QRasterWindow_connect_maximumHeightChanged")
	reg(&pg_QRasterWindow_connect_safeAreaMarginsChanged, "QRasterWindow_connect_safeAreaMarginsChanged")
	reg(&pg_QRasterWindow_connect_visibleChanged, "QRasterWindow_connect_visibleChanged")
	reg(&pg_QRasterWindow_connect_visibilityChanged, "QRasterWindow_connect_visibilityChanged")
	reg(&pg_QRasterWindow_connect_activeChanged, "QRasterWindow_connect_activeChanged")
	reg(&pg_QRasterWindow_connect_contentOrientationChanged, "QRasterWindow_connect_contentOrientationChanged")
	reg(&pg_QRasterWindow_connect_focusObjectChanged, "QRasterWindow_connect_focusObjectChanged")
	reg(&pg_QRasterWindow_connect_opacityChanged, "QRasterWindow_connect_opacityChanged")
	reg(&pg_QRasterWindow_connect_transientParentChanged, "QRasterWindow_connect_transientParentChanged")
	reg(&pg_QRasterWindow_connect_destroyed, "QRasterWindow_connect_destroyed")
	reg(&pg_QRasterWindow_connect_destroyed2, "QRasterWindow_connect_destroyed2")
	reg(&pg_QRasterWindow_connect_objectNameChanged, "QRasterWindow_connect_objectNameChanged")
	reg(&pg_QRasterWindow__disconnect, "QRasterWindow__disconnect")
	reg(&pg_QRasterWindow__set_override, "QRasterWindow__set_override")
	reg(&pg_QRasterWindow_delete, "QRasterWindow_delete")
}

type QRasterWindow struct {
	h uintptr
	*QPaintDeviceWindow
}

func bindQRasterWindow(h uintptr) *QRasterWindow {
	if h == 0 {
		return nil
	}
	var __c [3]uintptr
	purego.SyscallN(addr_QRasterWindow_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QRasterWindow{h: h, QPaintDeviceWindow: &QPaintDeviceWindow{h: __c[0], QWindow: &QWindow{h: __c[1], QObject: &QObject{h: __c[2]}}}}
}

func newownedQRasterWindow(h uintptr) *QRasterWindow { return bindQRasterWindow(h) }

func (this *QRasterWindow) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QRasterWindow) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQRasterWindow(h unsafe.Pointer) *QRasterWindow { return bindQRasterWindow(uintptr(h)) }
func UnsafeNewQRasterWindowOwned(h unsafe.Pointer) *QRasterWindow { return newownedQRasterWindow(uintptr(h)) }

func (this *QRasterWindow) UnsafeRelease() {
	this.h = 0
	if this.QPaintDeviceWindow != nil {
		this.QPaintDeviceWindow.UnsafeRelease()
	}
}

func NewQRasterWindow() *QRasterWindow {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QRasterWindow_new, uintptr(__goid))
	__o := newownedQRasterWindow(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QRasterWindow_virtIdx, QRasterWindow_enable)
	return __o
}

func NewQRasterWindow2(parent *QWindow) *QRasterWindow {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QRasterWindow_new2, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQRasterWindow(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QRasterWindow_virtIdx, QRasterWindow_enable)
	return __o
}

func QRasterWindow_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QRasterWindow_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func QRasterWindow_FromWinId(id uint64) *QWindow {
	__r, _, _ := purego.SyscallN(addr_QRasterWindow_fromWinId, uintptr(id))
	return bindQWindow(__r)
}

func QRasterWindow_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QRasterWindow_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QRasterWindow_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QRasterWindow_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QRasterWindow) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QRasterWindow_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QRasterWindow) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QRasterWindow_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QRasterWindow) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QRasterWindow_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QRasterWindow) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QRasterWindow_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QRasterWindow_DevicePixelRatioFScale() float64 {
	__r := pg_QRasterWindow_devicePixelRatioFScale()
	return __r
}

func QRasterWindow_EncodeMetricF(metric QPaintDevice__PaintDeviceMetric, value float64) int {
	__r := pg_QRasterWindow_encodeMetricF(int32(metric), value)
	return int(__r)
}

func (this *QRasterWindow) GetDecodedMetricF(metricA QPaintDevice__PaintDeviceMetric, metricB QPaintDevice__PaintDeviceMetric) float64 {
	__r := pg_QRasterWindow_getDecodedMetricF(this.h, int32(metricA), int32(metricB))
	runtime.KeepAlive(this)
	return __r
}

func QRasterWindow_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QRasterWindow_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QRasterWindow_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QRasterWindow_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QRasterWindow_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QRasterWindow_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QRasterWindow_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QRasterWindow_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QRasterWindow) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QRasterWindow_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QRasterWindow) MetricDefault(metric QPaintDevice__PaintDeviceMetric) int {
	__r0, _, _ := purego.SyscallN(addr_QRasterWindow_metric_default, this.h, uintptr(metric))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QRasterWindow) RedirectedDefault(arg1 *QPoint) *QPaintDevice {
	__r, _, _ := purego.SyscallN(addr_QRasterWindow_redirected_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return bindQPaintDevice(__r)
}

func (this *QRasterWindow) ResizeEventDefault(event *QResizeEvent) {
	purego.SyscallN(addr_QRasterWindow_resizeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QRasterWindow) ExposeEventDefault(arg1 *QExposeEvent) {
	purego.SyscallN(addr_QRasterWindow_exposeEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QRasterWindow) PaintEventDefault(event *QPaintEvent) {
	purego.SyscallN(addr_QRasterWindow_paintEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QRasterWindow) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QRasterWindow_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QRasterWindow) SurfaceTypeDefault() QSurface__SurfaceType {
	__r0, _, _ := purego.SyscallN(addr_QRasterWindow_surfaceType_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QSurface__SurfaceType(__r)
}

func (this *QRasterWindow) FormatDefault() *QSurfaceFormat {
	__r, _, _ := purego.SyscallN(addr_QRasterWindow_format_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSurfaceFormat(__r)
}

func (this *QRasterWindow) SizeDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QRasterWindow_size_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QRasterWindow) AccessibleRootDefault() *QAccessibleInterface {
	__r, _, _ := purego.SyscallN(addr_QRasterWindow_accessibleRoot_default, this.h)
	runtime.KeepAlive(this)
	return bindQAccessibleInterface(__r)
}

func (this *QRasterWindow) FocusObjectDefault() *QObject {
	__r, _, _ := purego.SyscallN(addr_QRasterWindow_focusObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QRasterWindow) MoveEventDefault(arg1 *QMoveEvent) {
	purego.SyscallN(addr_QRasterWindow_moveEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QRasterWindow) FocusInEventDefault(arg1 *QFocusEvent) {
	purego.SyscallN(addr_QRasterWindow_focusInEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QRasterWindow) FocusOutEventDefault(arg1 *QFocusEvent) {
	purego.SyscallN(addr_QRasterWindow_focusOutEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QRasterWindow) ShowEventDefault(arg1 *QShowEvent) {
	purego.SyscallN(addr_QRasterWindow_showEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QRasterWindow) HideEventDefault(arg1 *QHideEvent) {
	purego.SyscallN(addr_QRasterWindow_hideEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QRasterWindow) CloseEventDefault(arg1 *QCloseEvent) {
	purego.SyscallN(addr_QRasterWindow_closeEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QRasterWindow) KeyPressEventDefault(arg1 *QKeyEvent) {
	purego.SyscallN(addr_QRasterWindow_keyPressEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QRasterWindow) KeyReleaseEventDefault(arg1 *QKeyEvent) {
	purego.SyscallN(addr_QRasterWindow_keyReleaseEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QRasterWindow) MousePressEventDefault(arg1 *QMouseEvent) {
	purego.SyscallN(addr_QRasterWindow_mousePressEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QRasterWindow) MouseReleaseEventDefault(arg1 *QMouseEvent) {
	purego.SyscallN(addr_QRasterWindow_mouseReleaseEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QRasterWindow) MouseDoubleClickEventDefault(arg1 *QMouseEvent) {
	purego.SyscallN(addr_QRasterWindow_mouseDoubleClickEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QRasterWindow) MouseMoveEventDefault(arg1 *QMouseEvent) {
	purego.SyscallN(addr_QRasterWindow_mouseMoveEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QRasterWindow) WheelEventDefault(arg1 *QWheelEvent) {
	purego.SyscallN(addr_QRasterWindow_wheelEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QRasterWindow) TouchEventDefault(arg1 *QTouchEvent) {
	purego.SyscallN(addr_QRasterWindow_touchEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QRasterWindow) TabletEventDefault(arg1 *QTabletEvent) {
	purego.SyscallN(addr_QRasterWindow_tabletEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QRasterWindow) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QRasterWindow_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QRasterWindow) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QRasterWindow_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QRasterWindow) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QRasterWindow_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QRasterWindow) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QRasterWindow_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QRasterWindow) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QRasterWindow_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QRasterWindow) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QRasterWindow_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QRasterWindow) DevTypeDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QRasterWindow_devType_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QRasterWindow) InitPainterDefault(painter *QPainter) {
	purego.SyscallN(addr_QRasterWindow_initPainter_default, this.h, uintptr(painter.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
}

func (this *QRasterWindow) SharedPainterDefault() *QPainter {
	__r, _, _ := purego.SyscallN(addr_QRasterWindow_sharedPainter_default, this.h)
	runtime.KeepAlive(this)
	return bindQPainter(__r)
}

func (this *QRasterWindow) OnScreenChanged(fn func(*QScreen)) int64 {
	id := registerCallback(func(screen uintptr) { fn(UnsafeNewQScreen(unsafe.Pointer(screen))) })
	pg_QRasterWindow_connect_screenChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QRasterWindow) OnModalityChanged(fn func(WindowModality)) int64 {
	id := registerCallback(func(modality int32) { fn(WindowModality(modality)) })
	pg_QRasterWindow_connect_modalityChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QRasterWindow) OnFlagsChanged(fn func(WindowType)) int64 {
	id := registerCallback(func(flags int32) { fn(WindowType(flags)) })
	pg_QRasterWindow_connect_flagsChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QRasterWindow) OnWindowStateChanged(fn func(WindowState)) int64 {
	id := registerCallback(func(windowState int32) { fn(WindowState(windowState)) })
	pg_QRasterWindow_connect_windowStateChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QRasterWindow) OnWindowTitleChanged(fn func(string)) int64 {
	id := registerCallback(func(title uintptr, title_len int64) { fn(pgStrNoFree(title, title_len)) })
	pg_QRasterWindow_connect_windowTitleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QRasterWindow) OnXChanged(fn func(int)) int64 {
	id := registerCallback(func(arg int32) { fn(int(arg)) })
	pg_QRasterWindow_connect_xChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QRasterWindow) OnYChanged(fn func(int)) int64 {
	id := registerCallback(func(arg int32) { fn(int(arg)) })
	pg_QRasterWindow_connect_yChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QRasterWindow) OnWidthChanged(fn func(int)) int64 {
	id := registerCallback(func(arg int32) { fn(int(arg)) })
	pg_QRasterWindow_connect_widthChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QRasterWindow) OnHeightChanged(fn func(int)) int64 {
	id := registerCallback(func(arg int32) { fn(int(arg)) })
	pg_QRasterWindow_connect_heightChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QRasterWindow) OnMinimumWidthChanged(fn func(int)) int64 {
	id := registerCallback(func(arg int32) { fn(int(arg)) })
	pg_QRasterWindow_connect_minimumWidthChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QRasterWindow) OnMinimumHeightChanged(fn func(int)) int64 {
	id := registerCallback(func(arg int32) { fn(int(arg)) })
	pg_QRasterWindow_connect_minimumHeightChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QRasterWindow) OnMaximumWidthChanged(fn func(int)) int64 {
	id := registerCallback(func(arg int32) { fn(int(arg)) })
	pg_QRasterWindow_connect_maximumWidthChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QRasterWindow) OnMaximumHeightChanged(fn func(int)) int64 {
	id := registerCallback(func(arg int32) { fn(int(arg)) })
	pg_QRasterWindow_connect_maximumHeightChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QRasterWindow) OnSafeAreaMarginsChanged(fn func(*QMargins)) int64 {
	id := registerCallback(func(arg uintptr) { fn(UnsafeNewQMargins(unsafe.Pointer(arg))) })
	pg_QRasterWindow_connect_safeAreaMarginsChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QRasterWindow) OnVisibleChanged(fn func(bool)) int64 {
	id := registerCallback(func(arg int32) { fn(arg != 0) })
	pg_QRasterWindow_connect_visibleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QRasterWindow) OnVisibilityChanged(fn func(QWindow__Visibility)) int64 {
	id := registerCallback(func(visibility int32) { fn(QWindow__Visibility(visibility)) })
	pg_QRasterWindow_connect_visibilityChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QRasterWindow) OnActiveChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QRasterWindow_connect_activeChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QRasterWindow) OnContentOrientationChanged(fn func(ScreenOrientation)) int64 {
	id := registerCallback(func(orientation int32) { fn(ScreenOrientation(orientation)) })
	pg_QRasterWindow_connect_contentOrientationChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QRasterWindow) OnFocusObjectChanged(fn func(*QObject)) int64 {
	id := registerCallback(func(object uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(object))) })
	pg_QRasterWindow_connect_focusObjectChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QRasterWindow) OnOpacityChanged(fn func(float64)) int64 {
	id := registerCallback(func(opacity uintptr) { fn(math.Float64frombits(uint64(opacity))) })
	pg_QRasterWindow_connect_opacityChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QRasterWindow) OnTransientParentChanged(fn func(*QWindow)) int64 {
	id := registerCallback(func(transientParent uintptr) { fn(UnsafeNewQWindow(unsafe.Pointer(transientParent))) })
	pg_QRasterWindow_connect_transientParentChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QRasterWindow) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QRasterWindow_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QRasterWindow) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QRasterWindow_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QRasterWindow) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QRasterWindow_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QRasterWindow) Disconnect(handle int64) {
	pg_QRasterWindow__disconnect(handle)
	unregisterCallback(handle)
}

func QRasterWindow_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case 427629649267311777:
		return 1
	case 4127240633310655048:
		return 2
	case -4302917285655014661:
		return 3
	case 3239118852587300383:
		return 4
	case 5290442163167458465:
		return 5
	case 1237236456312824991:
		return 6
	case 7339588379958472532:
		return 7
	case 7267508590720539762:
		return 8
	case 5614464919154503228:
		return 9
	case -2255663269888600317:
		return 10
	case 8913927178009815944:
		return 11
	case -9161949150489525396:
		return 12
	case 975043767033552552:
		return 13
	case -385596010576048993:
		return 14
	case -6825369174751916:
		return 15
	case -199718707876719129:
		return 16
	case -1989937050019464783:
		return 17
	case 1023791977679793185:
		return 18
	case 772598475950050299:
		return 19
	case -332090898078543319:
		return 20
	case -4415920049297795213:
		return 21
	case -4029073905697023887:
		return 22
	case -4255494004099806581:
		return 23
	case 1919988340934511316:
		return 24
	case -2204038732330958816:
		return 25
	case -212068866860701141:
		return 26
	case -2039166586246443706:
		return 27
	case 5217623514630858297:
		return 28
	case -5409927581812374498:
		return 29
	case -1185315367526857833:
		return 30
	case 3769868229123868474:
		return 31
	case 7379029982170503008:
		return 32
	case -568369772982215962:
		return 33
	case -8659679893868524044:
		return 34
	case 6477500296980050210:
		return 35
	case 5792771328741774991:
		return 36
	}
	return -1
}

func QRasterWindow_enable(h uintptr, idx int32) { pg_QRasterWindow__set_override(h, idx) }

func (this *QRasterWindow) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QRasterWindow) OnMetric(fn func(QPaintDevice__PaintDeviceMetric) int) {
	registerVirt(this.UnsafeGoID(), 427629649267311777, func(metric int32) int32 { return int32(fn(QPaintDevice__PaintDeviceMetric(metric))) }) // slot "metric"
}

func (this *QRasterWindow) OnRedirected(fn func(*QPoint) *QPaintDevice) {
	registerVirt(this.UnsafeGoID(), 4127240633310655048, func(arg1 uintptr) uintptr { return uintptr(fn(UnsafeNewQPoint(unsafe.Pointer(arg1))).UnsafePointer()) }) // slot "redirected"
}

func (this *QRasterWindow) OnResizeEvent(fn func(*QResizeEvent)) {
	registerVirt(this.UnsafeGoID(), -4302917285655014661, func(event uintptr) { fn(UnsafeNewQResizeEvent(unsafe.Pointer(event))) }) // slot "resizeEvent"
}

func (this *QRasterWindow) OnExposeEvent(fn func(*QExposeEvent)) {
	registerVirt(this.UnsafeGoID(), 3239118852587300383, func(arg1 uintptr) { fn(UnsafeNewQExposeEvent(unsafe.Pointer(arg1))) }) // slot "exposeEvent"
}

func (this *QRasterWindow) OnPaintEvent(fn func(*QPaintEvent)) {
	registerVirt(this.UnsafeGoID(), 5290442163167458465, func(event uintptr) { fn(UnsafeNewQPaintEvent(unsafe.Pointer(event))) }) // slot "paintEvent"
}

func (this *QRasterWindow) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QRasterWindow) OnSurfaceType(fn func() QSurface__SurfaceType) {
	registerVirt(this.UnsafeGoID(), 7339588379958472532, func() int32 { return int32(fn()) }) // slot "surfaceType"
}

func (this *QRasterWindow) OnFormat(fn func() *QSurfaceFormat) {
	registerVirt(this.UnsafeGoID(), 7267508590720539762, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "format"
}

func (this *QRasterWindow) OnSize(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 5614464919154503228, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "size"
}

func (this *QRasterWindow) OnAccessibleRoot(fn func() *QAccessibleInterface) {
	registerVirt(this.UnsafeGoID(), -2255663269888600317, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "accessibleRoot"
}

func (this *QRasterWindow) OnFocusObject(fn func() *QObject) {
	registerVirt(this.UnsafeGoID(), 8913927178009815944, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "focusObject"
}

func (this *QRasterWindow) OnMoveEvent(fn func(*QMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -9161949150489525396, func(arg1 uintptr) { fn(UnsafeNewQMoveEvent(unsafe.Pointer(arg1))) }) // slot "moveEvent"
}

func (this *QRasterWindow) OnFocusInEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), 975043767033552552, func(arg1 uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(arg1))) }) // slot "focusInEvent"
}

func (this *QRasterWindow) OnFocusOutEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), -385596010576048993, func(arg1 uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(arg1))) }) // slot "focusOutEvent"
}

func (this *QRasterWindow) OnShowEvent(fn func(*QShowEvent)) {
	registerVirt(this.UnsafeGoID(), -6825369174751916, func(arg1 uintptr) { fn(UnsafeNewQShowEvent(unsafe.Pointer(arg1))) }) // slot "showEvent"
}

func (this *QRasterWindow) OnHideEvent(fn func(*QHideEvent)) {
	registerVirt(this.UnsafeGoID(), -199718707876719129, func(arg1 uintptr) { fn(UnsafeNewQHideEvent(unsafe.Pointer(arg1))) }) // slot "hideEvent"
}

func (this *QRasterWindow) OnCloseEvent(fn func(*QCloseEvent)) {
	registerVirt(this.UnsafeGoID(), -1989937050019464783, func(arg1 uintptr) { fn(UnsafeNewQCloseEvent(unsafe.Pointer(arg1))) }) // slot "closeEvent"
}

func (this *QRasterWindow) OnKeyPressEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 1023791977679793185, func(arg1 uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(arg1))) }) // slot "keyPressEvent"
}

func (this *QRasterWindow) OnKeyReleaseEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 772598475950050299, func(arg1 uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(arg1))) }) // slot "keyReleaseEvent"
}

func (this *QRasterWindow) OnMousePressEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -332090898078543319, func(arg1 uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(arg1))) }) // slot "mousePressEvent"
}

func (this *QRasterWindow) OnMouseReleaseEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4415920049297795213, func(arg1 uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(arg1))) }) // slot "mouseReleaseEvent"
}

func (this *QRasterWindow) OnMouseDoubleClickEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4029073905697023887, func(arg1 uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(arg1))) }) // slot "mouseDoubleClickEvent"
}

func (this *QRasterWindow) OnMouseMoveEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4255494004099806581, func(arg1 uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(arg1))) }) // slot "mouseMoveEvent"
}

func (this *QRasterWindow) OnWheelEvent(fn func(*QWheelEvent)) {
	registerVirt(this.UnsafeGoID(), 1919988340934511316, func(arg1 uintptr) { fn(UnsafeNewQWheelEvent(unsafe.Pointer(arg1))) }) // slot "wheelEvent"
}

func (this *QRasterWindow) OnTouchEvent(fn func(*QTouchEvent)) {
	registerVirt(this.UnsafeGoID(), -2204038732330958816, func(arg1 uintptr) { fn(UnsafeNewQTouchEvent(unsafe.Pointer(arg1))) }) // slot "touchEvent"
}

func (this *QRasterWindow) OnTabletEvent(fn func(*QTabletEvent)) {
	registerVirt(this.UnsafeGoID(), -212068866860701141, func(arg1 uintptr) { fn(UnsafeNewQTabletEvent(unsafe.Pointer(arg1))) }) // slot "tabletEvent"
}

func (this *QRasterWindow) OnNativeEvent(fn func(*QByteArray, unsafe.Pointer, *int64) bool) {
	registerVirt(this.UnsafeGoID(), -2039166586246443706, func(eventType uintptr, message uintptr, result uintptr) int32 { return b2i(fn(UnsafeNewQByteArray(unsafe.Pointer(eventType)), unsafe.Pointer(message), (*int64)(unsafe.Pointer(result)))) }) // slot "nativeEvent"
}

func (this *QRasterWindow) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QRasterWindow) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QRasterWindow) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QRasterWindow) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QRasterWindow) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QRasterWindow) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QRasterWindow) OnDevType(fn func() int) {
	registerVirt(this.UnsafeGoID(), -8659679893868524044, func() int32 { return int32(fn()) }) // slot "devType"
}

func (this *QRasterWindow) OnInitPainter(fn func(*QPainter)) {
	registerVirt(this.UnsafeGoID(), 6477500296980050210, func(painter uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter))) }) // slot "initPainter"
}

func (this *QRasterWindow) OnSharedPainter(fn func() *QPainter) {
	registerVirt(this.UnsafeGoID(), 5792771328741774991, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sharedPainter"
}

func (this *QRasterWindow) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QRasterWindow_delete(this.h)
	}
	this.UnsafeRelease()
}
