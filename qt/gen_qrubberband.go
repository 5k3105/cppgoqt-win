package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QRubberBand_chain uintptr
	addr_QRubberBand_new uintptr
	addr_QRubberBand_new2 uintptr
	addr_QRubberBand_tr uintptr
	addr_QRubberBand_shape uintptr
	addr_QRubberBand_setGeometry uintptr
	addr_QRubberBand_setGeometry2 uintptr
	addr_QRubberBand_move uintptr
	addr_QRubberBand_move2 uintptr
	addr_QRubberBand_resize uintptr
	addr_QRubberBand_resize2 uintptr
	addr_QRubberBand_setTabOrder uintptr
	addr_QRubberBand_mouseGrabber uintptr
	addr_QRubberBand_keyboardGrabber uintptr
	addr_QRubberBand_find uintptr
	addr_QRubberBand_createWindowContainer uintptr
	addr_QRubberBand_connect uintptr
	addr_QRubberBand_connect2 uintptr
	pg_QRubberBand_devicePixelRatioFScale func() float64
	pg_QRubberBand_encodeMetricF func(int32, float64) int32
	addr_QRubberBand_tr2 uintptr
	addr_QRubberBand_tr3 uintptr
	addr_QRubberBand_createWindowContainer2 uintptr
	addr_QRubberBand_createWindowContainer3 uintptr
	addr_QRubberBand_connect3 uintptr
	addr_QRubberBand_connect4 uintptr
	pg_QRubberBand_connect_windowTitleChanged func(uintptr, int64)
	pg_QRubberBand_connect_windowIconChanged func(uintptr, int64)
	pg_QRubberBand_connect_windowIconTextChanged func(uintptr, int64)
	pg_QRubberBand_connect_customContextMenuRequested func(uintptr, int64)
	pg_QRubberBand_connect_destroyed func(uintptr, int64)
	pg_QRubberBand_connect_destroyed2 func(uintptr, int64)
	pg_QRubberBand_connect_objectNameChanged func(uintptr, int64)
	pg_QRubberBand__disconnect func(int64)
	pg_QRubberBand_delete func(uintptr)
)

func init() {
	symaddr(&addr_QRubberBand_chain, "QRubberBand_chain")
	symaddr(&addr_QRubberBand_new, "QRubberBand_new")
	symaddr(&addr_QRubberBand_new2, "QRubberBand_new2")
	symaddr(&addr_QRubberBand_tr, "QRubberBand_tr")
	symaddr(&addr_QRubberBand_shape, "QRubberBand_shape")
	symaddr(&addr_QRubberBand_setGeometry, "QRubberBand_setGeometry")
	symaddr(&addr_QRubberBand_setGeometry2, "QRubberBand_setGeometry2")
	symaddr(&addr_QRubberBand_move, "QRubberBand_move")
	symaddr(&addr_QRubberBand_move2, "QRubberBand_move2")
	symaddr(&addr_QRubberBand_resize, "QRubberBand_resize")
	symaddr(&addr_QRubberBand_resize2, "QRubberBand_resize2")
	symaddr(&addr_QRubberBand_setTabOrder, "QRubberBand_setTabOrder")
	symaddr(&addr_QRubberBand_mouseGrabber, "QRubberBand_mouseGrabber")
	symaddr(&addr_QRubberBand_keyboardGrabber, "QRubberBand_keyboardGrabber")
	symaddr(&addr_QRubberBand_find, "QRubberBand_find")
	symaddr(&addr_QRubberBand_createWindowContainer, "QRubberBand_createWindowContainer")
	symaddr(&addr_QRubberBand_connect, "QRubberBand_connect")
	symaddr(&addr_QRubberBand_connect2, "QRubberBand_connect2")
	reg(&pg_QRubberBand_devicePixelRatioFScale, "QRubberBand_devicePixelRatioFScale")
	reg(&pg_QRubberBand_encodeMetricF, "QRubberBand_encodeMetricF")
	symaddr(&addr_QRubberBand_tr2, "QRubberBand_tr2")
	symaddr(&addr_QRubberBand_tr3, "QRubberBand_tr3")
	symaddr(&addr_QRubberBand_createWindowContainer2, "QRubberBand_createWindowContainer2")
	symaddr(&addr_QRubberBand_createWindowContainer3, "QRubberBand_createWindowContainer3")
	symaddr(&addr_QRubberBand_connect3, "QRubberBand_connect3")
	symaddr(&addr_QRubberBand_connect4, "QRubberBand_connect4")
	reg(&pg_QRubberBand_connect_windowTitleChanged, "QRubberBand_connect_windowTitleChanged")
	reg(&pg_QRubberBand_connect_windowIconChanged, "QRubberBand_connect_windowIconChanged")
	reg(&pg_QRubberBand_connect_windowIconTextChanged, "QRubberBand_connect_windowIconTextChanged")
	reg(&pg_QRubberBand_connect_customContextMenuRequested, "QRubberBand_connect_customContextMenuRequested")
	reg(&pg_QRubberBand_connect_destroyed, "QRubberBand_connect_destroyed")
	reg(&pg_QRubberBand_connect_destroyed2, "QRubberBand_connect_destroyed2")
	reg(&pg_QRubberBand_connect_objectNameChanged, "QRubberBand_connect_objectNameChanged")
	reg(&pg_QRubberBand__disconnect, "QRubberBand__disconnect")
	reg(&pg_QRubberBand_delete, "QRubberBand_delete")
}

type QRubberBand struct {
	h uintptr
	*QWidget
}

func bindQRubberBand(h uintptr) *QRubberBand {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QRubberBand_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QRubberBand{h: h, QWidget: &QWidget{h: __c[0], QObject: &QObject{h: __c[1]}}}
}

func newownedQRubberBand(h uintptr) *QRubberBand { return bindQRubberBand(h) }

func (this *QRubberBand) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QRubberBand) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQRubberBand(h unsafe.Pointer) *QRubberBand { return bindQRubberBand(uintptr(h)) }
func UnsafeNewQRubberBandOwned(h unsafe.Pointer) *QRubberBand { return newownedQRubberBand(uintptr(h)) }

func (this *QRubberBand) UnsafeRelease() {
	this.h = 0
	if this.QWidget != nil {
		this.QWidget.UnsafeRelease()
	}
}

func NewQRubberBand(arg1 QRubberBand__Shape) *QRubberBand {
	__r, _, _ := purego.SyscallN(addr_QRubberBand_new, uintptr(arg1))
	__o := newownedQRubberBand(__r)
	return __o
}

func NewQRubberBand2(arg1 QRubberBand__Shape, arg2 *QWidget) *QRubberBand {
	__r, _, _ := purego.SyscallN(addr_QRubberBand_new2, uintptr(arg1), uintptr(arg2.UnsafePointer()))
	__o := newownedQRubberBand(__r)
	runtime.KeepAlive(arg2)
	return __o
}

func QRubberBand_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QRubberBand_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QRubberBand) Shape() QRubberBand__Shape {
	__r0, _, _ := purego.SyscallN(addr_QRubberBand_shape, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QRubberBand__Shape(__r)
}

func (this *QRubberBand) SetGeometry(r *QRect) {
	purego.SyscallN(addr_QRubberBand_setGeometry, this.h, uintptr(r.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(r)
}

func (this *QRubberBand) SetGeometry2(x int, y int, w int, h int) {
	purego.SyscallN(addr_QRubberBand_setGeometry2, this.h, uintptr(x), uintptr(y), uintptr(w), uintptr(h))
	runtime.KeepAlive(this)
}

func (this *QRubberBand) Move(x int, y int) {
	purego.SyscallN(addr_QRubberBand_move, this.h, uintptr(x), uintptr(y))
	runtime.KeepAlive(this)
}

func (this *QRubberBand) MoveWithQPoint(p *QPoint) {
	purego.SyscallN(addr_QRubberBand_move2, this.h, uintptr(p.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
}

func (this *QRubberBand) Resize(w int, h int) {
	purego.SyscallN(addr_QRubberBand_resize, this.h, uintptr(w), uintptr(h))
	runtime.KeepAlive(this)
}

func (this *QRubberBand) ResizeWithQSize(s *QSize) {
	purego.SyscallN(addr_QRubberBand_resize2, this.h, uintptr(s.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(s)
}

func QRubberBand_SetTabOrder(arg1 *QWidget, arg2 *QWidget) {
	purego.SyscallN(addr_QRubberBand_setTabOrder, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()))
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
}

func QRubberBand_MouseGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QRubberBand_mouseGrabber)
	return bindQWidget(__r)
}

func QRubberBand_KeyboardGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QRubberBand_keyboardGrabber)
	return bindQWidget(__r)
}

func QRubberBand_Find(arg1 uint64) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QRubberBand_find, uintptr(arg1))
	return bindQWidget(__r)
}

func QRubberBand_CreateWindowContainer(window *QWindow) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QRubberBand_createWindowContainer, uintptr(window.UnsafePointer()))
	runtime.KeepAlive(window)
	return bindQWidget(__r)
}

func QRubberBand_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QRubberBand_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QRubberBand_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QRubberBand_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func QRubberBand_DevicePixelRatioFScale() float64 {
	__r := pg_QRubberBand_devicePixelRatioFScale()
	return __r
}

func QRubberBand_EncodeMetricF(metric QPaintDevice__PaintDeviceMetric, value float64) int {
	__r := pg_QRubberBand_encodeMetricF(int32(metric), value)
	return int(__r)
}

func QRubberBand_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QRubberBand_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QRubberBand_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QRubberBand_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QRubberBand_CreateWindowContainer2(window *QWindow, parent *QWidget) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QRubberBand_createWindowContainer2, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func QRubberBand_CreateWindowContainer3(window *QWindow, parent *QWidget, flags WindowType) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QRubberBand_createWindowContainer3, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()), uintptr(flags))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func QRubberBand_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QRubberBand_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QRubberBand_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QRubberBand_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QRubberBand) OnWindowTitleChanged(fn func(string)) int64 {
	id := registerCallback(func(title uintptr, title_len int64) { fn(pgStrNoFree(title, title_len)) })
	pg_QRubberBand_connect_windowTitleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QRubberBand) OnWindowIconChanged(fn func(*QIcon)) int64 {
	id := registerCallback(func(icon uintptr) { fn(UnsafeNewQIcon(unsafe.Pointer(icon))) })
	pg_QRubberBand_connect_windowIconChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QRubberBand) OnWindowIconTextChanged(fn func(string)) int64 {
	id := registerCallback(func(iconText uintptr, iconText_len int64) { fn(pgStrNoFree(iconText, iconText_len)) })
	pg_QRubberBand_connect_windowIconTextChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QRubberBand) OnCustomContextMenuRequested(fn func(*QPoint)) int64 {
	id := registerCallback(func(pos uintptr) { fn(UnsafeNewQPoint(unsafe.Pointer(pos))) })
	pg_QRubberBand_connect_customContextMenuRequested(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QRubberBand) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QRubberBand_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QRubberBand) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QRubberBand_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QRubberBand) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QRubberBand_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QRubberBand) Disconnect(handle int64) {
	pg_QRubberBand__disconnect(handle)
	unregisterCallback(handle)
}

func (this *QRubberBand) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QRubberBand_delete(this.h)
	}
	this.UnsafeRelease()
}
