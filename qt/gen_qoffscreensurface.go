package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QOffscreenSurface_base func(uintptr) uintptr
	addr_QOffscreenSurface_new uintptr
	addr_QOffscreenSurface_new2 uintptr
	addr_QOffscreenSurface_new3 uintptr
	addr_QOffscreenSurface_tr uintptr
	addr_QOffscreenSurface_surfaceType uintptr
	addr_QOffscreenSurface_create uintptr
	addr_QOffscreenSurface_destroy uintptr
	addr_QOffscreenSurface_isValid uintptr
	addr_QOffscreenSurface_setFormat uintptr
	addr_QOffscreenSurface_format uintptr
	addr_QOffscreenSurface_requestedFormat uintptr
	addr_QOffscreenSurface_size uintptr
	addr_QOffscreenSurface_screen uintptr
	addr_QOffscreenSurface_setScreen uintptr
	addr_QOffscreenSurface_screenChanged uintptr
	addr_QOffscreenSurface_connect uintptr
	addr_QOffscreenSurface_connect2 uintptr
	addr_QOffscreenSurface_sender uintptr
	addr_QOffscreenSurface_senderSignalIndex uintptr
	addr_QOffscreenSurface_receivers uintptr
	addr_QOffscreenSurface_isSignalConnected uintptr
	addr_QOffscreenSurface_surfaceClass uintptr
	addr_QOffscreenSurface_supportsOpenGL uintptr
	addr_QOffscreenSurface_tr2 uintptr
	addr_QOffscreenSurface_tr3 uintptr
	addr_QOffscreenSurface_connect3 uintptr
	addr_QOffscreenSurface_connect4 uintptr
	addr_QOffscreenSurface_asQSurface uintptr
	addr_QOffscreenSurface_surfaceType_default uintptr
	addr_QOffscreenSurface_format_default uintptr
	addr_QOffscreenSurface_size_default uintptr
	addr_QOffscreenSurface_metaObject_default uintptr
	addr_QOffscreenSurface_event_default uintptr
	addr_QOffscreenSurface_eventFilter_default uintptr
	addr_QOffscreenSurface_timerEvent_default uintptr
	addr_QOffscreenSurface_childEvent_default uintptr
	addr_QOffscreenSurface_customEvent_default uintptr
	addr_QOffscreenSurface_connectNotify_default uintptr
	addr_QOffscreenSurface_disconnectNotify_default uintptr
	pg_QOffscreenSurface_connect_screenChanged func(uintptr, int64)
	pg_QOffscreenSurface_connect_destroyed func(uintptr, int64)
	pg_QOffscreenSurface_connect_destroyed2 func(uintptr, int64)
	pg_QOffscreenSurface_connect_objectNameChanged func(uintptr, int64)
	pg_QOffscreenSurface__disconnect func(int64)
	pg_QOffscreenSurface__set_override func(uintptr, int32)
	pg_QOffscreenSurface_delete func(uintptr)
)

func init() {
	reg(&pg_QOffscreenSurface_base, "QOffscreenSurface_base")
	symaddr(&addr_QOffscreenSurface_new, "QOffscreenSurface_new")
	symaddr(&addr_QOffscreenSurface_new2, "QOffscreenSurface_new2")
	symaddr(&addr_QOffscreenSurface_new3, "QOffscreenSurface_new3")
	symaddr(&addr_QOffscreenSurface_tr, "QOffscreenSurface_tr")
	symaddr(&addr_QOffscreenSurface_surfaceType, "QOffscreenSurface_surfaceType")
	symaddr(&addr_QOffscreenSurface_create, "QOffscreenSurface_create")
	symaddr(&addr_QOffscreenSurface_destroy, "QOffscreenSurface_destroy")
	symaddr(&addr_QOffscreenSurface_isValid, "QOffscreenSurface_isValid")
	symaddr(&addr_QOffscreenSurface_setFormat, "QOffscreenSurface_setFormat")
	symaddr(&addr_QOffscreenSurface_format, "QOffscreenSurface_format")
	symaddr(&addr_QOffscreenSurface_requestedFormat, "QOffscreenSurface_requestedFormat")
	symaddr(&addr_QOffscreenSurface_size, "QOffscreenSurface_size")
	symaddr(&addr_QOffscreenSurface_screen, "QOffscreenSurface_screen")
	symaddr(&addr_QOffscreenSurface_setScreen, "QOffscreenSurface_setScreen")
	symaddr(&addr_QOffscreenSurface_screenChanged, "QOffscreenSurface_screenChanged")
	symaddr(&addr_QOffscreenSurface_connect, "QOffscreenSurface_connect")
	symaddr(&addr_QOffscreenSurface_connect2, "QOffscreenSurface_connect2")
	symaddr(&addr_QOffscreenSurface_sender, "QOffscreenSurface_sender")
	symaddr(&addr_QOffscreenSurface_senderSignalIndex, "QOffscreenSurface_senderSignalIndex")
	symaddr(&addr_QOffscreenSurface_receivers, "QOffscreenSurface_receivers")
	symaddr(&addr_QOffscreenSurface_isSignalConnected, "QOffscreenSurface_isSignalConnected")
	symaddr(&addr_QOffscreenSurface_surfaceClass, "QOffscreenSurface_surfaceClass")
	symaddr(&addr_QOffscreenSurface_supportsOpenGL, "QOffscreenSurface_supportsOpenGL")
	symaddr(&addr_QOffscreenSurface_tr2, "QOffscreenSurface_tr2")
	symaddr(&addr_QOffscreenSurface_tr3, "QOffscreenSurface_tr3")
	symaddr(&addr_QOffscreenSurface_connect3, "QOffscreenSurface_connect3")
	symaddr(&addr_QOffscreenSurface_connect4, "QOffscreenSurface_connect4")
	symaddr(&addr_QOffscreenSurface_asQSurface, "QOffscreenSurface_asQSurface")
	symaddr(&addr_QOffscreenSurface_surfaceType_default, "QOffscreenSurface_surfaceType_default")
	symaddr(&addr_QOffscreenSurface_format_default, "QOffscreenSurface_format_default")
	symaddr(&addr_QOffscreenSurface_size_default, "QOffscreenSurface_size_default")
	symaddr(&addr_QOffscreenSurface_metaObject_default, "QOffscreenSurface_metaObject_default")
	symaddr(&addr_QOffscreenSurface_event_default, "QOffscreenSurface_event_default")
	symaddr(&addr_QOffscreenSurface_eventFilter_default, "QOffscreenSurface_eventFilter_default")
	symaddr(&addr_QOffscreenSurface_timerEvent_default, "QOffscreenSurface_timerEvent_default")
	symaddr(&addr_QOffscreenSurface_childEvent_default, "QOffscreenSurface_childEvent_default")
	symaddr(&addr_QOffscreenSurface_customEvent_default, "QOffscreenSurface_customEvent_default")
	symaddr(&addr_QOffscreenSurface_connectNotify_default, "QOffscreenSurface_connectNotify_default")
	symaddr(&addr_QOffscreenSurface_disconnectNotify_default, "QOffscreenSurface_disconnectNotify_default")
	reg(&pg_QOffscreenSurface_connect_screenChanged, "QOffscreenSurface_connect_screenChanged")
	reg(&pg_QOffscreenSurface_connect_destroyed, "QOffscreenSurface_connect_destroyed")
	reg(&pg_QOffscreenSurface_connect_destroyed2, "QOffscreenSurface_connect_destroyed2")
	reg(&pg_QOffscreenSurface_connect_objectNameChanged, "QOffscreenSurface_connect_objectNameChanged")
	reg(&pg_QOffscreenSurface__disconnect, "QOffscreenSurface__disconnect")
	reg(&pg_QOffscreenSurface__set_override, "QOffscreenSurface__set_override")
	reg(&pg_QOffscreenSurface_delete, "QOffscreenSurface_delete")
}

type QOffscreenSurface struct {
	h uintptr
	*QObject
}

func bindQOffscreenSurface(h uintptr) *QOffscreenSurface {
	if h == 0 {
		return nil
	}
	return &QOffscreenSurface{h: h, QObject: bindQObject(pg_QOffscreenSurface_base(h))}
}

func newownedQOffscreenSurface(h uintptr) *QOffscreenSurface { return bindQOffscreenSurface(h) }

func (this *QOffscreenSurface) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QOffscreenSurface) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQOffscreenSurface(h unsafe.Pointer) *QOffscreenSurface { return bindQOffscreenSurface(uintptr(h)) }
func UnsafeNewQOffscreenSurfaceOwned(h unsafe.Pointer) *QOffscreenSurface { return newownedQOffscreenSurface(uintptr(h)) }

func (this *QOffscreenSurface) UnsafeRelease() {
	this.h = 0
	if this.QObject != nil {
		this.QObject.UnsafeRelease()
	}
}

func NewQOffscreenSurface() *QOffscreenSurface {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QOffscreenSurface_new, uintptr(__goid))
	__o := newownedQOffscreenSurface(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QOffscreenSurface_virtIdx, QOffscreenSurface_enable)
	return __o
}

func NewQOffscreenSurface2(screen *QScreen) *QOffscreenSurface {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QOffscreenSurface_new2, uintptr(__goid), uintptr(screen.UnsafePointer()))
	__o := newownedQOffscreenSurface(__r)
	runtime.KeepAlive(screen)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QOffscreenSurface_virtIdx, QOffscreenSurface_enable)
	return __o
}

func NewQOffscreenSurface3(screen *QScreen, parent *QObject) *QOffscreenSurface {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QOffscreenSurface_new3, uintptr(__goid), uintptr(screen.UnsafePointer()), uintptr(parent.UnsafePointer()))
	__o := newownedQOffscreenSurface(__r)
	runtime.KeepAlive(screen)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QOffscreenSurface_virtIdx, QOffscreenSurface_enable)
	return __o
}

func QOffscreenSurface_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QOffscreenSurface_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QOffscreenSurface) SurfaceType() QSurface__SurfaceType {
	__r0, _, _ := purego.SyscallN(addr_QOffscreenSurface_surfaceType, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QSurface__SurfaceType(__r)
}

func (this *QOffscreenSurface) Create() {
	purego.SyscallN(addr_QOffscreenSurface_create, this.h)
	runtime.KeepAlive(this)
}

func (this *QOffscreenSurface) Destroy() {
	purego.SyscallN(addr_QOffscreenSurface_destroy, this.h)
	runtime.KeepAlive(this)
}

func (this *QOffscreenSurface) IsValid() bool {
	__r0, _, _ := purego.SyscallN(addr_QOffscreenSurface_isValid, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QOffscreenSurface) SetFormat(format *QSurfaceFormat) {
	purego.SyscallN(addr_QOffscreenSurface_setFormat, this.h, uintptr(format.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(format)
}

func (this *QOffscreenSurface) Format() *QSurfaceFormat {
	__r, _, _ := purego.SyscallN(addr_QOffscreenSurface_format, this.h)
	runtime.KeepAlive(this)
	return newownedQSurfaceFormat(__r)
}

func (this *QOffscreenSurface) RequestedFormat() *QSurfaceFormat {
	__r, _, _ := purego.SyscallN(addr_QOffscreenSurface_requestedFormat, this.h)
	runtime.KeepAlive(this)
	return newownedQSurfaceFormat(__r)
}

func (this *QOffscreenSurface) Size() *QSize {
	__r, _, _ := purego.SyscallN(addr_QOffscreenSurface_size, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QOffscreenSurface) Screen() *QScreen {
	__r, _, _ := purego.SyscallN(addr_QOffscreenSurface_screen, this.h)
	runtime.KeepAlive(this)
	return bindQScreen(__r)
}

func (this *QOffscreenSurface) SetScreen(screen *QScreen) {
	purego.SyscallN(addr_QOffscreenSurface_setScreen, this.h, uintptr(screen.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(screen)
}

func (this *QOffscreenSurface) ScreenChanged(screen *QScreen) {
	purego.SyscallN(addr_QOffscreenSurface_screenChanged, this.h, uintptr(screen.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(screen)
}

func QOffscreenSurface_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QOffscreenSurface_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QOffscreenSurface_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QOffscreenSurface_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QOffscreenSurface) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QOffscreenSurface_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QOffscreenSurface) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QOffscreenSurface_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QOffscreenSurface) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QOffscreenSurface_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QOffscreenSurface) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QOffscreenSurface_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func (this *QOffscreenSurface) SurfaceClass() QSurface__SurfaceClass {
	__r0, _, _ := purego.SyscallN(addr_QOffscreenSurface_surfaceClass, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QSurface__SurfaceClass(__r)
}

func (this *QOffscreenSurface) SupportsOpenGL() bool {
	__r0, _, _ := purego.SyscallN(addr_QOffscreenSurface_supportsOpenGL, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QOffscreenSurface_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QOffscreenSurface_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QOffscreenSurface_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QOffscreenSurface_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QOffscreenSurface_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QOffscreenSurface_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QOffscreenSurface_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QOffscreenSurface_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QOffscreenSurface) AsQSurface() *QSurface {
	__r, _, _ := purego.SyscallN(addr_QOffscreenSurface_asQSurface, this.h)
	runtime.KeepAlive(this)
	return bindQSurface(__r)
}

func (this *QOffscreenSurface) SurfaceTypeDefault() QSurface__SurfaceType {
	__r0, _, _ := purego.SyscallN(addr_QOffscreenSurface_surfaceType_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QSurface__SurfaceType(__r)
}

func (this *QOffscreenSurface) FormatDefault() *QSurfaceFormat {
	__r, _, _ := purego.SyscallN(addr_QOffscreenSurface_format_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSurfaceFormat(__r)
}

func (this *QOffscreenSurface) SizeDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QOffscreenSurface_size_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QOffscreenSurface) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QOffscreenSurface_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QOffscreenSurface) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QOffscreenSurface_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QOffscreenSurface) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QOffscreenSurface_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QOffscreenSurface) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QOffscreenSurface_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QOffscreenSurface) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QOffscreenSurface_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QOffscreenSurface) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QOffscreenSurface_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QOffscreenSurface) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QOffscreenSurface_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QOffscreenSurface) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QOffscreenSurface_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QOffscreenSurface) OnScreenChanged(fn func(*QScreen)) int64 {
	id := registerCallback(func(screen uintptr) { fn(UnsafeNewQScreen(unsafe.Pointer(screen))) })
	pg_QOffscreenSurface_connect_screenChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QOffscreenSurface) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QOffscreenSurface_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QOffscreenSurface) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QOffscreenSurface_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QOffscreenSurface) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QOffscreenSurface_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QOffscreenSurface) Disconnect(handle int64) {
	pg_QOffscreenSurface__disconnect(handle)
	unregisterCallback(handle)
}

func QOffscreenSurface_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case 7339588379958472532:
		return 1
	case 7267508590720539762:
		return 2
	case 5614464919154503228:
		return 3
	case 1237236456312824991:
		return 4
	case 5217623514630858297:
		return 5
	case -5409927581812374498:
		return 6
	case -1185315367526857833:
		return 7
	case 3769868229123868474:
		return 8
	case 7379029982170503008:
		return 9
	case -568369772982215962:
		return 10
	}
	return -1
}

func QOffscreenSurface_enable(h uintptr, idx int32) { pg_QOffscreenSurface__set_override(h, idx) }

func (this *QOffscreenSurface) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QOffscreenSurface) OnSurfaceType(fn func() QSurface__SurfaceType) {
	registerVirt(this.UnsafeGoID(), 7339588379958472532, func() int32 { return int32(fn()) }) // slot "surfaceType"
}

func (this *QOffscreenSurface) OnFormat(fn func() *QSurfaceFormat) {
	registerVirt(this.UnsafeGoID(), 7267508590720539762, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "format"
}

func (this *QOffscreenSurface) OnSize(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 5614464919154503228, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "size"
}

func (this *QOffscreenSurface) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QOffscreenSurface) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QOffscreenSurface) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QOffscreenSurface) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QOffscreenSurface) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QOffscreenSurface) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QOffscreenSurface) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QOffscreenSurface) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QOffscreenSurface_delete(this.h)
	}
	this.UnsafeRelease()
}
