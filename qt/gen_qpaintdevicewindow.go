package qt

import (
	"unsafe"
	"math"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QPaintDeviceWindow_chain uintptr
	addr_QPaintDeviceWindow_tr uintptr
	addr_QPaintDeviceWindow_update uintptr
	addr_QPaintDeviceWindow_update2 uintptr
	addr_QPaintDeviceWindow_update3 uintptr
	addr_QPaintDeviceWindow_fromWinId uintptr
	addr_QPaintDeviceWindow_connect uintptr
	addr_QPaintDeviceWindow_connect2 uintptr
	addr_QPaintDeviceWindow_devType uintptr
	addr_QPaintDeviceWindow_paintingActive uintptr
	addr_QPaintDeviceWindow_widthMM uintptr
	addr_QPaintDeviceWindow_heightMM uintptr
	addr_QPaintDeviceWindow_logicalDpiX uintptr
	addr_QPaintDeviceWindow_logicalDpiY uintptr
	addr_QPaintDeviceWindow_physicalDpiX uintptr
	addr_QPaintDeviceWindow_physicalDpiY uintptr
	pg_QPaintDeviceWindow_devicePixelRatioF func(uintptr) float64
	addr_QPaintDeviceWindow_colorCount uintptr
	addr_QPaintDeviceWindow_depth uintptr
	pg_QPaintDeviceWindow_devicePixelRatioFScale func() float64
	pg_QPaintDeviceWindow_encodeMetricF func(int32, float64) int32
	addr_QPaintDeviceWindow_tr2 uintptr
	addr_QPaintDeviceWindow_tr3 uintptr
	addr_QPaintDeviceWindow_connect3 uintptr
	addr_QPaintDeviceWindow_connect4 uintptr
	addr_QPaintDeviceWindow_asQPaintDevice uintptr
	pg_QPaintDeviceWindow_connect_screenChanged func(uintptr, int64)
	pg_QPaintDeviceWindow_connect_modalityChanged func(uintptr, int64)
	pg_QPaintDeviceWindow_connect_flagsChanged func(uintptr, int64)
	pg_QPaintDeviceWindow_connect_windowStateChanged func(uintptr, int64)
	pg_QPaintDeviceWindow_connect_windowTitleChanged func(uintptr, int64)
	pg_QPaintDeviceWindow_connect_xChanged func(uintptr, int64)
	pg_QPaintDeviceWindow_connect_yChanged func(uintptr, int64)
	pg_QPaintDeviceWindow_connect_widthChanged func(uintptr, int64)
	pg_QPaintDeviceWindow_connect_heightChanged func(uintptr, int64)
	pg_QPaintDeviceWindow_connect_minimumWidthChanged func(uintptr, int64)
	pg_QPaintDeviceWindow_connect_minimumHeightChanged func(uintptr, int64)
	pg_QPaintDeviceWindow_connect_maximumWidthChanged func(uintptr, int64)
	pg_QPaintDeviceWindow_connect_maximumHeightChanged func(uintptr, int64)
	pg_QPaintDeviceWindow_connect_safeAreaMarginsChanged func(uintptr, int64)
	pg_QPaintDeviceWindow_connect_visibleChanged func(uintptr, int64)
	pg_QPaintDeviceWindow_connect_visibilityChanged func(uintptr, int64)
	pg_QPaintDeviceWindow_connect_activeChanged func(uintptr, int64)
	pg_QPaintDeviceWindow_connect_contentOrientationChanged func(uintptr, int64)
	pg_QPaintDeviceWindow_connect_focusObjectChanged func(uintptr, int64)
	pg_QPaintDeviceWindow_connect_opacityChanged func(uintptr, int64)
	pg_QPaintDeviceWindow_connect_transientParentChanged func(uintptr, int64)
	pg_QPaintDeviceWindow_connect_destroyed func(uintptr, int64)
	pg_QPaintDeviceWindow_connect_destroyed2 func(uintptr, int64)
	pg_QPaintDeviceWindow_connect_objectNameChanged func(uintptr, int64)
	pg_QPaintDeviceWindow__disconnect func(int64)
	pg_QPaintDeviceWindow_delete func(uintptr)
)

func init() {
	symaddr(&addr_QPaintDeviceWindow_chain, "QPaintDeviceWindow_chain")
	symaddr(&addr_QPaintDeviceWindow_tr, "QPaintDeviceWindow_tr")
	symaddr(&addr_QPaintDeviceWindow_update, "QPaintDeviceWindow_update")
	symaddr(&addr_QPaintDeviceWindow_update2, "QPaintDeviceWindow_update2")
	symaddr(&addr_QPaintDeviceWindow_update3, "QPaintDeviceWindow_update3")
	symaddr(&addr_QPaintDeviceWindow_fromWinId, "QPaintDeviceWindow_fromWinId")
	symaddr(&addr_QPaintDeviceWindow_connect, "QPaintDeviceWindow_connect")
	symaddr(&addr_QPaintDeviceWindow_connect2, "QPaintDeviceWindow_connect2")
	symaddr(&addr_QPaintDeviceWindow_devType, "QPaintDeviceWindow_devType")
	symaddr(&addr_QPaintDeviceWindow_paintingActive, "QPaintDeviceWindow_paintingActive")
	symaddr(&addr_QPaintDeviceWindow_widthMM, "QPaintDeviceWindow_widthMM")
	symaddr(&addr_QPaintDeviceWindow_heightMM, "QPaintDeviceWindow_heightMM")
	symaddr(&addr_QPaintDeviceWindow_logicalDpiX, "QPaintDeviceWindow_logicalDpiX")
	symaddr(&addr_QPaintDeviceWindow_logicalDpiY, "QPaintDeviceWindow_logicalDpiY")
	symaddr(&addr_QPaintDeviceWindow_physicalDpiX, "QPaintDeviceWindow_physicalDpiX")
	symaddr(&addr_QPaintDeviceWindow_physicalDpiY, "QPaintDeviceWindow_physicalDpiY")
	reg(&pg_QPaintDeviceWindow_devicePixelRatioF, "QPaintDeviceWindow_devicePixelRatioF")
	symaddr(&addr_QPaintDeviceWindow_colorCount, "QPaintDeviceWindow_colorCount")
	symaddr(&addr_QPaintDeviceWindow_depth, "QPaintDeviceWindow_depth")
	reg(&pg_QPaintDeviceWindow_devicePixelRatioFScale, "QPaintDeviceWindow_devicePixelRatioFScale")
	reg(&pg_QPaintDeviceWindow_encodeMetricF, "QPaintDeviceWindow_encodeMetricF")
	symaddr(&addr_QPaintDeviceWindow_tr2, "QPaintDeviceWindow_tr2")
	symaddr(&addr_QPaintDeviceWindow_tr3, "QPaintDeviceWindow_tr3")
	symaddr(&addr_QPaintDeviceWindow_connect3, "QPaintDeviceWindow_connect3")
	symaddr(&addr_QPaintDeviceWindow_connect4, "QPaintDeviceWindow_connect4")
	symaddr(&addr_QPaintDeviceWindow_asQPaintDevice, "QPaintDeviceWindow_asQPaintDevice")
	reg(&pg_QPaintDeviceWindow_connect_screenChanged, "QPaintDeviceWindow_connect_screenChanged")
	reg(&pg_QPaintDeviceWindow_connect_modalityChanged, "QPaintDeviceWindow_connect_modalityChanged")
	reg(&pg_QPaintDeviceWindow_connect_flagsChanged, "QPaintDeviceWindow_connect_flagsChanged")
	reg(&pg_QPaintDeviceWindow_connect_windowStateChanged, "QPaintDeviceWindow_connect_windowStateChanged")
	reg(&pg_QPaintDeviceWindow_connect_windowTitleChanged, "QPaintDeviceWindow_connect_windowTitleChanged")
	reg(&pg_QPaintDeviceWindow_connect_xChanged, "QPaintDeviceWindow_connect_xChanged")
	reg(&pg_QPaintDeviceWindow_connect_yChanged, "QPaintDeviceWindow_connect_yChanged")
	reg(&pg_QPaintDeviceWindow_connect_widthChanged, "QPaintDeviceWindow_connect_widthChanged")
	reg(&pg_QPaintDeviceWindow_connect_heightChanged, "QPaintDeviceWindow_connect_heightChanged")
	reg(&pg_QPaintDeviceWindow_connect_minimumWidthChanged, "QPaintDeviceWindow_connect_minimumWidthChanged")
	reg(&pg_QPaintDeviceWindow_connect_minimumHeightChanged, "QPaintDeviceWindow_connect_minimumHeightChanged")
	reg(&pg_QPaintDeviceWindow_connect_maximumWidthChanged, "QPaintDeviceWindow_connect_maximumWidthChanged")
	reg(&pg_QPaintDeviceWindow_connect_maximumHeightChanged, "QPaintDeviceWindow_connect_maximumHeightChanged")
	reg(&pg_QPaintDeviceWindow_connect_safeAreaMarginsChanged, "QPaintDeviceWindow_connect_safeAreaMarginsChanged")
	reg(&pg_QPaintDeviceWindow_connect_visibleChanged, "QPaintDeviceWindow_connect_visibleChanged")
	reg(&pg_QPaintDeviceWindow_connect_visibilityChanged, "QPaintDeviceWindow_connect_visibilityChanged")
	reg(&pg_QPaintDeviceWindow_connect_activeChanged, "QPaintDeviceWindow_connect_activeChanged")
	reg(&pg_QPaintDeviceWindow_connect_contentOrientationChanged, "QPaintDeviceWindow_connect_contentOrientationChanged")
	reg(&pg_QPaintDeviceWindow_connect_focusObjectChanged, "QPaintDeviceWindow_connect_focusObjectChanged")
	reg(&pg_QPaintDeviceWindow_connect_opacityChanged, "QPaintDeviceWindow_connect_opacityChanged")
	reg(&pg_QPaintDeviceWindow_connect_transientParentChanged, "QPaintDeviceWindow_connect_transientParentChanged")
	reg(&pg_QPaintDeviceWindow_connect_destroyed, "QPaintDeviceWindow_connect_destroyed")
	reg(&pg_QPaintDeviceWindow_connect_destroyed2, "QPaintDeviceWindow_connect_destroyed2")
	reg(&pg_QPaintDeviceWindow_connect_objectNameChanged, "QPaintDeviceWindow_connect_objectNameChanged")
	reg(&pg_QPaintDeviceWindow__disconnect, "QPaintDeviceWindow__disconnect")
	reg(&pg_QPaintDeviceWindow_delete, "QPaintDeviceWindow_delete")
}

type QPaintDeviceWindow struct {
	h uintptr
	*QWindow
}

func bindQPaintDeviceWindow(h uintptr) *QPaintDeviceWindow {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QPaintDeviceWindow_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QPaintDeviceWindow{h: h, QWindow: &QWindow{h: __c[0], QObject: &QObject{h: __c[1]}}}
}

func newownedQPaintDeviceWindow(h uintptr) *QPaintDeviceWindow { return bindQPaintDeviceWindow(h) }

func (this *QPaintDeviceWindow) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QPaintDeviceWindow) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQPaintDeviceWindow(h unsafe.Pointer) *QPaintDeviceWindow { return bindQPaintDeviceWindow(uintptr(h)) }
func UnsafeNewQPaintDeviceWindowOwned(h unsafe.Pointer) *QPaintDeviceWindow { return newownedQPaintDeviceWindow(uintptr(h)) }

func (this *QPaintDeviceWindow) UnsafeRelease() {
	this.h = 0
	if this.QWindow != nil {
		this.QWindow.UnsafeRelease()
	}
}

func QPaintDeviceWindow_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QPaintDeviceWindow_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QPaintDeviceWindow) Update(rect *QRect) {
	purego.SyscallN(addr_QPaintDeviceWindow_update, this.h, uintptr(rect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
}

func (this *QPaintDeviceWindow) UpdateWithRegion(region *QRegion) {
	purego.SyscallN(addr_QPaintDeviceWindow_update2, this.h, uintptr(region.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(region)
}

func (this *QPaintDeviceWindow) Update2() {
	purego.SyscallN(addr_QPaintDeviceWindow_update3, this.h)
	runtime.KeepAlive(this)
}

func QPaintDeviceWindow_FromWinId(id uint64) *QWindow {
	__r, _, _ := purego.SyscallN(addr_QPaintDeviceWindow_fromWinId, uintptr(id))
	return bindQWindow(__r)
}

func QPaintDeviceWindow_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QPaintDeviceWindow_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QPaintDeviceWindow_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QPaintDeviceWindow_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QPaintDeviceWindow) DevType() int {
	__r0, _, _ := purego.SyscallN(addr_QPaintDeviceWindow_devType, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QPaintDeviceWindow) PaintingActive() bool {
	__r0, _, _ := purego.SyscallN(addr_QPaintDeviceWindow_paintingActive, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QPaintDeviceWindow) WidthMM() int {
	__r0, _, _ := purego.SyscallN(addr_QPaintDeviceWindow_widthMM, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QPaintDeviceWindow) HeightMM() int {
	__r0, _, _ := purego.SyscallN(addr_QPaintDeviceWindow_heightMM, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QPaintDeviceWindow) LogicalDpiX() int {
	__r0, _, _ := purego.SyscallN(addr_QPaintDeviceWindow_logicalDpiX, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QPaintDeviceWindow) LogicalDpiY() int {
	__r0, _, _ := purego.SyscallN(addr_QPaintDeviceWindow_logicalDpiY, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QPaintDeviceWindow) PhysicalDpiX() int {
	__r0, _, _ := purego.SyscallN(addr_QPaintDeviceWindow_physicalDpiX, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QPaintDeviceWindow) PhysicalDpiY() int {
	__r0, _, _ := purego.SyscallN(addr_QPaintDeviceWindow_physicalDpiY, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QPaintDeviceWindow) DevicePixelRatioF() float64 {
	__r := pg_QPaintDeviceWindow_devicePixelRatioF(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QPaintDeviceWindow) ColorCount() int {
	__r0, _, _ := purego.SyscallN(addr_QPaintDeviceWindow_colorCount, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QPaintDeviceWindow) Depth() int {
	__r0, _, _ := purego.SyscallN(addr_QPaintDeviceWindow_depth, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func QPaintDeviceWindow_DevicePixelRatioFScale() float64 {
	__r := pg_QPaintDeviceWindow_devicePixelRatioFScale()
	return __r
}

func QPaintDeviceWindow_EncodeMetricF(metric QPaintDevice__PaintDeviceMetric, value float64) int {
	__r := pg_QPaintDeviceWindow_encodeMetricF(int32(metric), value)
	return int(__r)
}

func QPaintDeviceWindow_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QPaintDeviceWindow_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QPaintDeviceWindow_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QPaintDeviceWindow_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QPaintDeviceWindow_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QPaintDeviceWindow_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QPaintDeviceWindow_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QPaintDeviceWindow_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QPaintDeviceWindow) AsQPaintDevice() *QPaintDevice {
	__r, _, _ := purego.SyscallN(addr_QPaintDeviceWindow_asQPaintDevice, this.h)
	runtime.KeepAlive(this)
	return bindQPaintDevice(__r)
}

func (this *QPaintDeviceWindow) OnScreenChanged(fn func(*QScreen)) int64 {
	id := registerCallback(func(screen uintptr) { fn(UnsafeNewQScreen(unsafe.Pointer(screen))) })
	pg_QPaintDeviceWindow_connect_screenChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QPaintDeviceWindow) OnModalityChanged(fn func(WindowModality)) int64 {
	id := registerCallback(func(modality int32) { fn(WindowModality(modality)) })
	pg_QPaintDeviceWindow_connect_modalityChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QPaintDeviceWindow) OnFlagsChanged(fn func(WindowType)) int64 {
	id := registerCallback(func(flags int32) { fn(WindowType(flags)) })
	pg_QPaintDeviceWindow_connect_flagsChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QPaintDeviceWindow) OnWindowStateChanged(fn func(WindowState)) int64 {
	id := registerCallback(func(windowState int32) { fn(WindowState(windowState)) })
	pg_QPaintDeviceWindow_connect_windowStateChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QPaintDeviceWindow) OnWindowTitleChanged(fn func(string)) int64 {
	id := registerCallback(func(title uintptr, title_len int64) { fn(pgStrNoFree(title, title_len)) })
	pg_QPaintDeviceWindow_connect_windowTitleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QPaintDeviceWindow) OnXChanged(fn func(int)) int64 {
	id := registerCallback(func(arg int32) { fn(int(arg)) })
	pg_QPaintDeviceWindow_connect_xChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QPaintDeviceWindow) OnYChanged(fn func(int)) int64 {
	id := registerCallback(func(arg int32) { fn(int(arg)) })
	pg_QPaintDeviceWindow_connect_yChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QPaintDeviceWindow) OnWidthChanged(fn func(int)) int64 {
	id := registerCallback(func(arg int32) { fn(int(arg)) })
	pg_QPaintDeviceWindow_connect_widthChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QPaintDeviceWindow) OnHeightChanged(fn func(int)) int64 {
	id := registerCallback(func(arg int32) { fn(int(arg)) })
	pg_QPaintDeviceWindow_connect_heightChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QPaintDeviceWindow) OnMinimumWidthChanged(fn func(int)) int64 {
	id := registerCallback(func(arg int32) { fn(int(arg)) })
	pg_QPaintDeviceWindow_connect_minimumWidthChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QPaintDeviceWindow) OnMinimumHeightChanged(fn func(int)) int64 {
	id := registerCallback(func(arg int32) { fn(int(arg)) })
	pg_QPaintDeviceWindow_connect_minimumHeightChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QPaintDeviceWindow) OnMaximumWidthChanged(fn func(int)) int64 {
	id := registerCallback(func(arg int32) { fn(int(arg)) })
	pg_QPaintDeviceWindow_connect_maximumWidthChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QPaintDeviceWindow) OnMaximumHeightChanged(fn func(int)) int64 {
	id := registerCallback(func(arg int32) { fn(int(arg)) })
	pg_QPaintDeviceWindow_connect_maximumHeightChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QPaintDeviceWindow) OnSafeAreaMarginsChanged(fn func(*QMargins)) int64 {
	id := registerCallback(func(arg uintptr) { fn(UnsafeNewQMargins(unsafe.Pointer(arg))) })
	pg_QPaintDeviceWindow_connect_safeAreaMarginsChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QPaintDeviceWindow) OnVisibleChanged(fn func(bool)) int64 {
	id := registerCallback(func(arg int32) { fn(arg != 0) })
	pg_QPaintDeviceWindow_connect_visibleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QPaintDeviceWindow) OnVisibilityChanged(fn func(QWindow__Visibility)) int64 {
	id := registerCallback(func(visibility int32) { fn(QWindow__Visibility(visibility)) })
	pg_QPaintDeviceWindow_connect_visibilityChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QPaintDeviceWindow) OnActiveChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QPaintDeviceWindow_connect_activeChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QPaintDeviceWindow) OnContentOrientationChanged(fn func(ScreenOrientation)) int64 {
	id := registerCallback(func(orientation int32) { fn(ScreenOrientation(orientation)) })
	pg_QPaintDeviceWindow_connect_contentOrientationChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QPaintDeviceWindow) OnFocusObjectChanged(fn func(*QObject)) int64 {
	id := registerCallback(func(object uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(object))) })
	pg_QPaintDeviceWindow_connect_focusObjectChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QPaintDeviceWindow) OnOpacityChanged(fn func(float64)) int64 {
	id := registerCallback(func(opacity uintptr) { fn(math.Float64frombits(uint64(opacity))) })
	pg_QPaintDeviceWindow_connect_opacityChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QPaintDeviceWindow) OnTransientParentChanged(fn func(*QWindow)) int64 {
	id := registerCallback(func(transientParent uintptr) { fn(UnsafeNewQWindow(unsafe.Pointer(transientParent))) })
	pg_QPaintDeviceWindow_connect_transientParentChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QPaintDeviceWindow) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QPaintDeviceWindow_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QPaintDeviceWindow) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QPaintDeviceWindow_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QPaintDeviceWindow) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QPaintDeviceWindow_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QPaintDeviceWindow) Disconnect(handle int64) {
	pg_QPaintDeviceWindow__disconnect(handle)
	unregisterCallback(handle)
}

func (this *QPaintDeviceWindow) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QPaintDeviceWindow_delete(this.h)
	}
	this.UnsafeRelease()
}
