package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QSizeGrip_chain uintptr
	addr_QSizeGrip_new uintptr
	addr_QSizeGrip_tr uintptr
	addr_QSizeGrip_setTabOrder uintptr
	addr_QSizeGrip_mouseGrabber uintptr
	addr_QSizeGrip_keyboardGrabber uintptr
	addr_QSizeGrip_find uintptr
	addr_QSizeGrip_createWindowContainer uintptr
	addr_QSizeGrip_connect uintptr
	addr_QSizeGrip_connect2 uintptr
	pg_QSizeGrip_devicePixelRatioFScale func() float64
	pg_QSizeGrip_encodeMetricF func(int32, float64) int32
	addr_QSizeGrip_tr2 uintptr
	addr_QSizeGrip_tr3 uintptr
	addr_QSizeGrip_createWindowContainer2 uintptr
	addr_QSizeGrip_createWindowContainer3 uintptr
	addr_QSizeGrip_connect3 uintptr
	addr_QSizeGrip_connect4 uintptr
	pg_QSizeGrip_connect_windowTitleChanged func(uintptr, int64)
	pg_QSizeGrip_connect_windowIconChanged func(uintptr, int64)
	pg_QSizeGrip_connect_windowIconTextChanged func(uintptr, int64)
	pg_QSizeGrip_connect_customContextMenuRequested func(uintptr, int64)
	pg_QSizeGrip_connect_destroyed func(uintptr, int64)
	pg_QSizeGrip_connect_destroyed2 func(uintptr, int64)
	pg_QSizeGrip_connect_objectNameChanged func(uintptr, int64)
	pg_QSizeGrip__disconnect func(int64)
	pg_QSizeGrip_delete func(uintptr)
)

func init() {
	symaddr(&addr_QSizeGrip_chain, "QSizeGrip_chain")
	symaddr(&addr_QSizeGrip_new, "QSizeGrip_new")
	symaddr(&addr_QSizeGrip_tr, "QSizeGrip_tr")
	symaddr(&addr_QSizeGrip_setTabOrder, "QSizeGrip_setTabOrder")
	symaddr(&addr_QSizeGrip_mouseGrabber, "QSizeGrip_mouseGrabber")
	symaddr(&addr_QSizeGrip_keyboardGrabber, "QSizeGrip_keyboardGrabber")
	symaddr(&addr_QSizeGrip_find, "QSizeGrip_find")
	symaddr(&addr_QSizeGrip_createWindowContainer, "QSizeGrip_createWindowContainer")
	symaddr(&addr_QSizeGrip_connect, "QSizeGrip_connect")
	symaddr(&addr_QSizeGrip_connect2, "QSizeGrip_connect2")
	reg(&pg_QSizeGrip_devicePixelRatioFScale, "QSizeGrip_devicePixelRatioFScale")
	reg(&pg_QSizeGrip_encodeMetricF, "QSizeGrip_encodeMetricF")
	symaddr(&addr_QSizeGrip_tr2, "QSizeGrip_tr2")
	symaddr(&addr_QSizeGrip_tr3, "QSizeGrip_tr3")
	symaddr(&addr_QSizeGrip_createWindowContainer2, "QSizeGrip_createWindowContainer2")
	symaddr(&addr_QSizeGrip_createWindowContainer3, "QSizeGrip_createWindowContainer3")
	symaddr(&addr_QSizeGrip_connect3, "QSizeGrip_connect3")
	symaddr(&addr_QSizeGrip_connect4, "QSizeGrip_connect4")
	reg(&pg_QSizeGrip_connect_windowTitleChanged, "QSizeGrip_connect_windowTitleChanged")
	reg(&pg_QSizeGrip_connect_windowIconChanged, "QSizeGrip_connect_windowIconChanged")
	reg(&pg_QSizeGrip_connect_windowIconTextChanged, "QSizeGrip_connect_windowIconTextChanged")
	reg(&pg_QSizeGrip_connect_customContextMenuRequested, "QSizeGrip_connect_customContextMenuRequested")
	reg(&pg_QSizeGrip_connect_destroyed, "QSizeGrip_connect_destroyed")
	reg(&pg_QSizeGrip_connect_destroyed2, "QSizeGrip_connect_destroyed2")
	reg(&pg_QSizeGrip_connect_objectNameChanged, "QSizeGrip_connect_objectNameChanged")
	reg(&pg_QSizeGrip__disconnect, "QSizeGrip__disconnect")
	reg(&pg_QSizeGrip_delete, "QSizeGrip_delete")
}

type QSizeGrip struct {
	h uintptr
	*QWidget
}

func bindQSizeGrip(h uintptr) *QSizeGrip {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QSizeGrip_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QSizeGrip{h: h, QWidget: &QWidget{h: __c[0], QObject: &QObject{h: __c[1]}}}
}

func newownedQSizeGrip(h uintptr) *QSizeGrip { return bindQSizeGrip(h) }

func (this *QSizeGrip) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QSizeGrip) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQSizeGrip(h unsafe.Pointer) *QSizeGrip { return bindQSizeGrip(uintptr(h)) }
func UnsafeNewQSizeGripOwned(h unsafe.Pointer) *QSizeGrip { return newownedQSizeGrip(uintptr(h)) }

func (this *QSizeGrip) UnsafeRelease() {
	this.h = 0
	if this.QWidget != nil {
		this.QWidget.UnsafeRelease()
	}
}

func NewQSizeGrip(parent *QWidget) *QSizeGrip {
	__r, _, _ := purego.SyscallN(addr_QSizeGrip_new, uintptr(parent.UnsafePointer()))
	__o := newownedQSizeGrip(__r)
	runtime.KeepAlive(parent)
	return __o
}

func QSizeGrip_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QSizeGrip_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func QSizeGrip_SetTabOrder(arg1 *QWidget, arg2 *QWidget) {
	purego.SyscallN(addr_QSizeGrip_setTabOrder, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()))
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
}

func QSizeGrip_MouseGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QSizeGrip_mouseGrabber)
	return bindQWidget(__r)
}

func QSizeGrip_KeyboardGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QSizeGrip_keyboardGrabber)
	return bindQWidget(__r)
}

func QSizeGrip_Find(arg1 uint64) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QSizeGrip_find, uintptr(arg1))
	return bindQWidget(__r)
}

func QSizeGrip_CreateWindowContainer(window *QWindow) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QSizeGrip_createWindowContainer, uintptr(window.UnsafePointer()))
	runtime.KeepAlive(window)
	return bindQWidget(__r)
}

func QSizeGrip_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QSizeGrip_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QSizeGrip_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QSizeGrip_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func QSizeGrip_DevicePixelRatioFScale() float64 {
	__r := pg_QSizeGrip_devicePixelRatioFScale()
	return __r
}

func QSizeGrip_EncodeMetricF(metric QPaintDevice__PaintDeviceMetric, value float64) int {
	__r := pg_QSizeGrip_encodeMetricF(int32(metric), value)
	return int(__r)
}

func QSizeGrip_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QSizeGrip_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QSizeGrip_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QSizeGrip_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QSizeGrip_CreateWindowContainer2(window *QWindow, parent *QWidget) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QSizeGrip_createWindowContainer2, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func QSizeGrip_CreateWindowContainer3(window *QWindow, parent *QWidget, flags WindowType) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QSizeGrip_createWindowContainer3, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()), uintptr(flags))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func QSizeGrip_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QSizeGrip_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QSizeGrip_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QSizeGrip_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QSizeGrip) OnWindowTitleChanged(fn func(string)) int64 {
	id := registerCallback(func(title uintptr, title_len int64) { fn(pgStrNoFree(title, title_len)) })
	pg_QSizeGrip_connect_windowTitleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSizeGrip) OnWindowIconChanged(fn func(*QIcon)) int64 {
	id := registerCallback(func(icon uintptr) { fn(UnsafeNewQIcon(unsafe.Pointer(icon))) })
	pg_QSizeGrip_connect_windowIconChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSizeGrip) OnWindowIconTextChanged(fn func(string)) int64 {
	id := registerCallback(func(iconText uintptr, iconText_len int64) { fn(pgStrNoFree(iconText, iconText_len)) })
	pg_QSizeGrip_connect_windowIconTextChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSizeGrip) OnCustomContextMenuRequested(fn func(*QPoint)) int64 {
	id := registerCallback(func(pos uintptr) { fn(UnsafeNewQPoint(unsafe.Pointer(pos))) })
	pg_QSizeGrip_connect_customContextMenuRequested(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSizeGrip) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QSizeGrip_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSizeGrip) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QSizeGrip_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSizeGrip) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QSizeGrip_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSizeGrip) Disconnect(handle int64) {
	pg_QSizeGrip__disconnect(handle)
	unregisterCallback(handle)
}

func (this *QSizeGrip) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QSizeGrip_delete(this.h)
	}
	this.UnsafeRelease()
}
