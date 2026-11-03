package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QSplitterHandle_chain uintptr
	addr_QSplitterHandle_new uintptr
	addr_QSplitterHandle_tr uintptr
	addr_QSplitterHandle_setOrientation uintptr
	addr_QSplitterHandle_orientation uintptr
	addr_QSplitterHandle_opaqueResize uintptr
	addr_QSplitterHandle_splitter uintptr
	addr_QSplitterHandle_setTabOrder uintptr
	addr_QSplitterHandle_mouseGrabber uintptr
	addr_QSplitterHandle_keyboardGrabber uintptr
	addr_QSplitterHandle_find uintptr
	addr_QSplitterHandle_createWindowContainer uintptr
	addr_QSplitterHandle_connect uintptr
	addr_QSplitterHandle_connect2 uintptr
	pg_QSplitterHandle_devicePixelRatioFScale func() float64
	pg_QSplitterHandle_encodeMetricF func(int32, float64) int32
	addr_QSplitterHandle_tr2 uintptr
	addr_QSplitterHandle_tr3 uintptr
	addr_QSplitterHandle_createWindowContainer2 uintptr
	addr_QSplitterHandle_createWindowContainer3 uintptr
	addr_QSplitterHandle_connect3 uintptr
	addr_QSplitterHandle_connect4 uintptr
	pg_QSplitterHandle_connect_windowTitleChanged func(uintptr, int64)
	pg_QSplitterHandle_connect_windowIconChanged func(uintptr, int64)
	pg_QSplitterHandle_connect_windowIconTextChanged func(uintptr, int64)
	pg_QSplitterHandle_connect_customContextMenuRequested func(uintptr, int64)
	pg_QSplitterHandle_connect_destroyed func(uintptr, int64)
	pg_QSplitterHandle_connect_destroyed2 func(uintptr, int64)
	pg_QSplitterHandle_connect_objectNameChanged func(uintptr, int64)
	pg_QSplitterHandle__disconnect func(int64)
	pg_QSplitterHandle_delete func(uintptr)
)

func init() {
	symaddr(&addr_QSplitterHandle_chain, "QSplitterHandle_chain")
	symaddr(&addr_QSplitterHandle_new, "QSplitterHandle_new")
	symaddr(&addr_QSplitterHandle_tr, "QSplitterHandle_tr")
	symaddr(&addr_QSplitterHandle_setOrientation, "QSplitterHandle_setOrientation")
	symaddr(&addr_QSplitterHandle_orientation, "QSplitterHandle_orientation")
	symaddr(&addr_QSplitterHandle_opaqueResize, "QSplitterHandle_opaqueResize")
	symaddr(&addr_QSplitterHandle_splitter, "QSplitterHandle_splitter")
	symaddr(&addr_QSplitterHandle_setTabOrder, "QSplitterHandle_setTabOrder")
	symaddr(&addr_QSplitterHandle_mouseGrabber, "QSplitterHandle_mouseGrabber")
	symaddr(&addr_QSplitterHandle_keyboardGrabber, "QSplitterHandle_keyboardGrabber")
	symaddr(&addr_QSplitterHandle_find, "QSplitterHandle_find")
	symaddr(&addr_QSplitterHandle_createWindowContainer, "QSplitterHandle_createWindowContainer")
	symaddr(&addr_QSplitterHandle_connect, "QSplitterHandle_connect")
	symaddr(&addr_QSplitterHandle_connect2, "QSplitterHandle_connect2")
	reg(&pg_QSplitterHandle_devicePixelRatioFScale, "QSplitterHandle_devicePixelRatioFScale")
	reg(&pg_QSplitterHandle_encodeMetricF, "QSplitterHandle_encodeMetricF")
	symaddr(&addr_QSplitterHandle_tr2, "QSplitterHandle_tr2")
	symaddr(&addr_QSplitterHandle_tr3, "QSplitterHandle_tr3")
	symaddr(&addr_QSplitterHandle_createWindowContainer2, "QSplitterHandle_createWindowContainer2")
	symaddr(&addr_QSplitterHandle_createWindowContainer3, "QSplitterHandle_createWindowContainer3")
	symaddr(&addr_QSplitterHandle_connect3, "QSplitterHandle_connect3")
	symaddr(&addr_QSplitterHandle_connect4, "QSplitterHandle_connect4")
	reg(&pg_QSplitterHandle_connect_windowTitleChanged, "QSplitterHandle_connect_windowTitleChanged")
	reg(&pg_QSplitterHandle_connect_windowIconChanged, "QSplitterHandle_connect_windowIconChanged")
	reg(&pg_QSplitterHandle_connect_windowIconTextChanged, "QSplitterHandle_connect_windowIconTextChanged")
	reg(&pg_QSplitterHandle_connect_customContextMenuRequested, "QSplitterHandle_connect_customContextMenuRequested")
	reg(&pg_QSplitterHandle_connect_destroyed, "QSplitterHandle_connect_destroyed")
	reg(&pg_QSplitterHandle_connect_destroyed2, "QSplitterHandle_connect_destroyed2")
	reg(&pg_QSplitterHandle_connect_objectNameChanged, "QSplitterHandle_connect_objectNameChanged")
	reg(&pg_QSplitterHandle__disconnect, "QSplitterHandle__disconnect")
	reg(&pg_QSplitterHandle_delete, "QSplitterHandle_delete")
}

type QSplitterHandle struct {
	h uintptr
	*QWidget
}

func bindQSplitterHandle(h uintptr) *QSplitterHandle {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QSplitterHandle_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QSplitterHandle{h: h, QWidget: &QWidget{h: __c[0], QObject: &QObject{h: __c[1]}}}
}

func newownedQSplitterHandle(h uintptr) *QSplitterHandle { return bindQSplitterHandle(h) }

func (this *QSplitterHandle) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QSplitterHandle) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQSplitterHandle(h unsafe.Pointer) *QSplitterHandle { return bindQSplitterHandle(uintptr(h)) }
func UnsafeNewQSplitterHandleOwned(h unsafe.Pointer) *QSplitterHandle { return newownedQSplitterHandle(uintptr(h)) }

func (this *QSplitterHandle) UnsafeRelease() {
	this.h = 0
	if this.QWidget != nil {
		this.QWidget.UnsafeRelease()
	}
}

func NewQSplitterHandle(o Orientation, parent *QSplitter) *QSplitterHandle {
	__r, _, _ := purego.SyscallN(addr_QSplitterHandle_new, uintptr(o), uintptr(parent.UnsafePointer()))
	__o := newownedQSplitterHandle(__r)
	runtime.KeepAlive(parent)
	return __o
}

func QSplitterHandle_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QSplitterHandle_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QSplitterHandle) SetOrientation(o Orientation) {
	purego.SyscallN(addr_QSplitterHandle_setOrientation, this.h, uintptr(o))
	runtime.KeepAlive(this)
}

func (this *QSplitterHandle) Orientation() Orientation {
	__r0, _, _ := purego.SyscallN(addr_QSplitterHandle_orientation, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return Orientation(__r)
}

func (this *QSplitterHandle) OpaqueResize() bool {
	__r0, _, _ := purego.SyscallN(addr_QSplitterHandle_opaqueResize, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QSplitterHandle) Splitter() *QSplitter {
	__r, _, _ := purego.SyscallN(addr_QSplitterHandle_splitter, this.h)
	runtime.KeepAlive(this)
	return bindQSplitter(__r)
}

func QSplitterHandle_SetTabOrder(arg1 *QWidget, arg2 *QWidget) {
	purego.SyscallN(addr_QSplitterHandle_setTabOrder, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()))
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
}

func QSplitterHandle_MouseGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QSplitterHandle_mouseGrabber)
	return bindQWidget(__r)
}

func QSplitterHandle_KeyboardGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QSplitterHandle_keyboardGrabber)
	return bindQWidget(__r)
}

func QSplitterHandle_Find(arg1 uint64) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QSplitterHandle_find, uintptr(arg1))
	return bindQWidget(__r)
}

func QSplitterHandle_CreateWindowContainer(window *QWindow) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QSplitterHandle_createWindowContainer, uintptr(window.UnsafePointer()))
	runtime.KeepAlive(window)
	return bindQWidget(__r)
}

func QSplitterHandle_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QSplitterHandle_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QSplitterHandle_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QSplitterHandle_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func QSplitterHandle_DevicePixelRatioFScale() float64 {
	__r := pg_QSplitterHandle_devicePixelRatioFScale()
	return __r
}

func QSplitterHandle_EncodeMetricF(metric QPaintDevice__PaintDeviceMetric, value float64) int {
	__r := pg_QSplitterHandle_encodeMetricF(int32(metric), value)
	return int(__r)
}

func QSplitterHandle_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QSplitterHandle_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QSplitterHandle_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QSplitterHandle_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QSplitterHandle_CreateWindowContainer2(window *QWindow, parent *QWidget) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QSplitterHandle_createWindowContainer2, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func QSplitterHandle_CreateWindowContainer3(window *QWindow, parent *QWidget, flags WindowType) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QSplitterHandle_createWindowContainer3, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()), uintptr(flags))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func QSplitterHandle_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QSplitterHandle_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QSplitterHandle_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QSplitterHandle_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QSplitterHandle) OnWindowTitleChanged(fn func(string)) int64 {
	id := registerCallback(func(title uintptr, title_len int64) { fn(pgStrNoFree(title, title_len)) })
	pg_QSplitterHandle_connect_windowTitleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSplitterHandle) OnWindowIconChanged(fn func(*QIcon)) int64 {
	id := registerCallback(func(icon uintptr) { fn(UnsafeNewQIcon(unsafe.Pointer(icon))) })
	pg_QSplitterHandle_connect_windowIconChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSplitterHandle) OnWindowIconTextChanged(fn func(string)) int64 {
	id := registerCallback(func(iconText uintptr, iconText_len int64) { fn(pgStrNoFree(iconText, iconText_len)) })
	pg_QSplitterHandle_connect_windowIconTextChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSplitterHandle) OnCustomContextMenuRequested(fn func(*QPoint)) int64 {
	id := registerCallback(func(pos uintptr) { fn(UnsafeNewQPoint(unsafe.Pointer(pos))) })
	pg_QSplitterHandle_connect_customContextMenuRequested(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSplitterHandle) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QSplitterHandle_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSplitterHandle) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QSplitterHandle_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSplitterHandle) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QSplitterHandle_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSplitterHandle) Disconnect(handle int64) {
	pg_QSplitterHandle__disconnect(handle)
	unregisterCallback(handle)
}

func (this *QSplitterHandle) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QSplitterHandle_delete(this.h)
	}
	this.UnsafeRelease()
}
