package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QOpenGLContext_base func(uintptr) uintptr
	addr_QOpenGLContext_new uintptr
	addr_QOpenGLContext_new2 uintptr
	addr_QOpenGLContext_tr uintptr
	addr_QOpenGLContext_setFormat uintptr
	addr_QOpenGLContext_setShareContext uintptr
	addr_QOpenGLContext_setScreen uintptr
	addr_QOpenGLContext_create uintptr
	addr_QOpenGLContext_isValid uintptr
	addr_QOpenGLContext_format uintptr
	addr_QOpenGLContext_shareContext uintptr
	addr_QOpenGLContext_shareGroup uintptr
	addr_QOpenGLContext_screen uintptr
	addr_QOpenGLContext_defaultFramebufferObject uintptr
	addr_QOpenGLContext_makeCurrent uintptr
	addr_QOpenGLContext_doneCurrent uintptr
	addr_QOpenGLContext_swapBuffers uintptr
	addr_QOpenGLContext_surface uintptr
	addr_QOpenGLContext_currentContext uintptr
	addr_QOpenGLContext_areSharing uintptr
	addr_QOpenGLContext_functions uintptr
	addr_QOpenGLContext_extraFunctions uintptr
	addr_QOpenGLContext_extensions uintptr
	addr_QOpenGLContext_hasExtension uintptr
	addr_QOpenGLContext_openGLModuleType uintptr
	addr_QOpenGLContext_isOpenGLES uintptr
	addr_QOpenGLContext_supportsThreadedOpenGL uintptr
	addr_QOpenGLContext_globalShareContext uintptr
	addr_QOpenGLContext_aboutToBeDestroyed uintptr
	addr_QOpenGLContext_connect uintptr
	addr_QOpenGLContext_connect2 uintptr
	addr_QOpenGLContext_sender uintptr
	addr_QOpenGLContext_senderSignalIndex uintptr
	addr_QOpenGLContext_receivers uintptr
	addr_QOpenGLContext_isSignalConnected uintptr
	addr_QOpenGLContext_tr2 uintptr
	addr_QOpenGLContext_tr3 uintptr
	addr_QOpenGLContext_connect3 uintptr
	addr_QOpenGLContext_connect4 uintptr
	addr_QOpenGLContext_metaObject_default uintptr
	addr_QOpenGLContext_event_default uintptr
	addr_QOpenGLContext_eventFilter_default uintptr
	addr_QOpenGLContext_timerEvent_default uintptr
	addr_QOpenGLContext_childEvent_default uintptr
	addr_QOpenGLContext_customEvent_default uintptr
	addr_QOpenGLContext_connectNotify_default uintptr
	addr_QOpenGLContext_disconnectNotify_default uintptr
	pg_QOpenGLContext_connect_aboutToBeDestroyed func(uintptr, int64)
	pg_QOpenGLContext_connect_destroyed func(uintptr, int64)
	pg_QOpenGLContext_connect_destroyed2 func(uintptr, int64)
	pg_QOpenGLContext_connect_objectNameChanged func(uintptr, int64)
	pg_QOpenGLContext__disconnect func(int64)
	pg_QOpenGLContext__set_override func(uintptr, int32)
	pg_QOpenGLContext_delete func(uintptr)
)

func init() {
	reg(&pg_QOpenGLContext_base, "QOpenGLContext_base")
	symaddr(&addr_QOpenGLContext_new, "QOpenGLContext_new")
	symaddr(&addr_QOpenGLContext_new2, "QOpenGLContext_new2")
	symaddr(&addr_QOpenGLContext_tr, "QOpenGLContext_tr")
	symaddr(&addr_QOpenGLContext_setFormat, "QOpenGLContext_setFormat")
	symaddr(&addr_QOpenGLContext_setShareContext, "QOpenGLContext_setShareContext")
	symaddr(&addr_QOpenGLContext_setScreen, "QOpenGLContext_setScreen")
	symaddr(&addr_QOpenGLContext_create, "QOpenGLContext_create")
	symaddr(&addr_QOpenGLContext_isValid, "QOpenGLContext_isValid")
	symaddr(&addr_QOpenGLContext_format, "QOpenGLContext_format")
	symaddr(&addr_QOpenGLContext_shareContext, "QOpenGLContext_shareContext")
	symaddr(&addr_QOpenGLContext_shareGroup, "QOpenGLContext_shareGroup")
	symaddr(&addr_QOpenGLContext_screen, "QOpenGLContext_screen")
	symaddr(&addr_QOpenGLContext_defaultFramebufferObject, "QOpenGLContext_defaultFramebufferObject")
	symaddr(&addr_QOpenGLContext_makeCurrent, "QOpenGLContext_makeCurrent")
	symaddr(&addr_QOpenGLContext_doneCurrent, "QOpenGLContext_doneCurrent")
	symaddr(&addr_QOpenGLContext_swapBuffers, "QOpenGLContext_swapBuffers")
	symaddr(&addr_QOpenGLContext_surface, "QOpenGLContext_surface")
	symaddr(&addr_QOpenGLContext_currentContext, "QOpenGLContext_currentContext")
	symaddr(&addr_QOpenGLContext_areSharing, "QOpenGLContext_areSharing")
	symaddr(&addr_QOpenGLContext_functions, "QOpenGLContext_functions")
	symaddr(&addr_QOpenGLContext_extraFunctions, "QOpenGLContext_extraFunctions")
	symaddr(&addr_QOpenGLContext_extensions, "QOpenGLContext_extensions")
	symaddr(&addr_QOpenGLContext_hasExtension, "QOpenGLContext_hasExtension")
	symaddr(&addr_QOpenGLContext_openGLModuleType, "QOpenGLContext_openGLModuleType")
	symaddr(&addr_QOpenGLContext_isOpenGLES, "QOpenGLContext_isOpenGLES")
	symaddr(&addr_QOpenGLContext_supportsThreadedOpenGL, "QOpenGLContext_supportsThreadedOpenGL")
	symaddr(&addr_QOpenGLContext_globalShareContext, "QOpenGLContext_globalShareContext")
	symaddr(&addr_QOpenGLContext_aboutToBeDestroyed, "QOpenGLContext_aboutToBeDestroyed")
	symaddr(&addr_QOpenGLContext_connect, "QOpenGLContext_connect")
	symaddr(&addr_QOpenGLContext_connect2, "QOpenGLContext_connect2")
	symaddr(&addr_QOpenGLContext_sender, "QOpenGLContext_sender")
	symaddr(&addr_QOpenGLContext_senderSignalIndex, "QOpenGLContext_senderSignalIndex")
	symaddr(&addr_QOpenGLContext_receivers, "QOpenGLContext_receivers")
	symaddr(&addr_QOpenGLContext_isSignalConnected, "QOpenGLContext_isSignalConnected")
	symaddr(&addr_QOpenGLContext_tr2, "QOpenGLContext_tr2")
	symaddr(&addr_QOpenGLContext_tr3, "QOpenGLContext_tr3")
	symaddr(&addr_QOpenGLContext_connect3, "QOpenGLContext_connect3")
	symaddr(&addr_QOpenGLContext_connect4, "QOpenGLContext_connect4")
	symaddr(&addr_QOpenGLContext_metaObject_default, "QOpenGLContext_metaObject_default")
	symaddr(&addr_QOpenGLContext_event_default, "QOpenGLContext_event_default")
	symaddr(&addr_QOpenGLContext_eventFilter_default, "QOpenGLContext_eventFilter_default")
	symaddr(&addr_QOpenGLContext_timerEvent_default, "QOpenGLContext_timerEvent_default")
	symaddr(&addr_QOpenGLContext_childEvent_default, "QOpenGLContext_childEvent_default")
	symaddr(&addr_QOpenGLContext_customEvent_default, "QOpenGLContext_customEvent_default")
	symaddr(&addr_QOpenGLContext_connectNotify_default, "QOpenGLContext_connectNotify_default")
	symaddr(&addr_QOpenGLContext_disconnectNotify_default, "QOpenGLContext_disconnectNotify_default")
	reg(&pg_QOpenGLContext_connect_aboutToBeDestroyed, "QOpenGLContext_connect_aboutToBeDestroyed")
	reg(&pg_QOpenGLContext_connect_destroyed, "QOpenGLContext_connect_destroyed")
	reg(&pg_QOpenGLContext_connect_destroyed2, "QOpenGLContext_connect_destroyed2")
	reg(&pg_QOpenGLContext_connect_objectNameChanged, "QOpenGLContext_connect_objectNameChanged")
	reg(&pg_QOpenGLContext__disconnect, "QOpenGLContext__disconnect")
	reg(&pg_QOpenGLContext__set_override, "QOpenGLContext__set_override")
	reg(&pg_QOpenGLContext_delete, "QOpenGLContext_delete")
}

type QOpenGLContext struct {
	h uintptr
	*QObject
}

func bindQOpenGLContext(h uintptr) *QOpenGLContext {
	if h == 0 {
		return nil
	}
	return &QOpenGLContext{h: h, QObject: bindQObject(pg_QOpenGLContext_base(h))}
}

func newownedQOpenGLContext(h uintptr) *QOpenGLContext { return bindQOpenGLContext(h) }

func (this *QOpenGLContext) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QOpenGLContext) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQOpenGLContext(h unsafe.Pointer) *QOpenGLContext { return bindQOpenGLContext(uintptr(h)) }
func UnsafeNewQOpenGLContextOwned(h unsafe.Pointer) *QOpenGLContext { return newownedQOpenGLContext(uintptr(h)) }

func (this *QOpenGLContext) UnsafeRelease() {
	this.h = 0
	if this.QObject != nil {
		this.QObject.UnsafeRelease()
	}
}

func NewQOpenGLContext() *QOpenGLContext {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QOpenGLContext_new, uintptr(__goid))
	__o := newownedQOpenGLContext(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QOpenGLContext_virtIdx, QOpenGLContext_enable)
	return __o
}

func NewQOpenGLContext2(parent *QObject) *QOpenGLContext {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QOpenGLContext_new2, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQOpenGLContext(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QOpenGLContext_virtIdx, QOpenGLContext_enable)
	return __o
}

func QOpenGLContext_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QOpenGLContext_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QOpenGLContext) SetFormat(format *QSurfaceFormat) {
	purego.SyscallN(addr_QOpenGLContext_setFormat, this.h, uintptr(format.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(format)
}

func (this *QOpenGLContext) SetShareContext(shareContext *QOpenGLContext) {
	purego.SyscallN(addr_QOpenGLContext_setShareContext, this.h, uintptr(shareContext.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(shareContext)
}

func (this *QOpenGLContext) SetScreen(screen *QScreen) {
	purego.SyscallN(addr_QOpenGLContext_setScreen, this.h, uintptr(screen.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(screen)
}

func (this *QOpenGLContext) Create() bool {
	__r0, _, _ := purego.SyscallN(addr_QOpenGLContext_create, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QOpenGLContext) IsValid() bool {
	__r0, _, _ := purego.SyscallN(addr_QOpenGLContext_isValid, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QOpenGLContext) Format() *QSurfaceFormat {
	__r, _, _ := purego.SyscallN(addr_QOpenGLContext_format, this.h)
	runtime.KeepAlive(this)
	return newownedQSurfaceFormat(__r)
}

func (this *QOpenGLContext) ShareContext() *QOpenGLContext {
	__r, _, _ := purego.SyscallN(addr_QOpenGLContext_shareContext, this.h)
	runtime.KeepAlive(this)
	return bindQOpenGLContext(__r)
}

func (this *QOpenGLContext) ShareGroup() *QOpenGLContextGroup {
	__r, _, _ := purego.SyscallN(addr_QOpenGLContext_shareGroup, this.h)
	runtime.KeepAlive(this)
	return bindQOpenGLContextGroup(__r)
}

func (this *QOpenGLContext) Screen() *QScreen {
	__r, _, _ := purego.SyscallN(addr_QOpenGLContext_screen, this.h)
	runtime.KeepAlive(this)
	return bindQScreen(__r)
}

func (this *QOpenGLContext) DefaultFramebufferObject() uint {
	__r0, _, _ := purego.SyscallN(addr_QOpenGLContext_defaultFramebufferObject, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint(__r)
}

func (this *QOpenGLContext) MakeCurrent(surface *QSurface) bool {
	__r0, _, _ := purego.SyscallN(addr_QOpenGLContext_makeCurrent, this.h, uintptr(surface.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(surface)
	return __r != 0
}

func (this *QOpenGLContext) DoneCurrent() {
	purego.SyscallN(addr_QOpenGLContext_doneCurrent, this.h)
	runtime.KeepAlive(this)
}

func (this *QOpenGLContext) SwapBuffers(surface *QSurface) {
	purego.SyscallN(addr_QOpenGLContext_swapBuffers, this.h, uintptr(surface.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(surface)
}

func (this *QOpenGLContext) Surface() *QSurface {
	__r, _, _ := purego.SyscallN(addr_QOpenGLContext_surface, this.h)
	runtime.KeepAlive(this)
	return bindQSurface(__r)
}

func QOpenGLContext_CurrentContext() *QOpenGLContext {
	__r, _, _ := purego.SyscallN(addr_QOpenGLContext_currentContext)
	return bindQOpenGLContext(__r)
}

func QOpenGLContext_AreSharing(first *QOpenGLContext, second *QOpenGLContext) bool {
	__r0, _, _ := purego.SyscallN(addr_QOpenGLContext_areSharing, uintptr(first.UnsafePointer()), uintptr(second.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(first)
	runtime.KeepAlive(second)
	return __r != 0
}

func (this *QOpenGLContext) Functions() *QOpenGLFunctions {
	__r, _, _ := purego.SyscallN(addr_QOpenGLContext_functions, this.h)
	runtime.KeepAlive(this)
	return bindQOpenGLFunctions(__r)
}

func (this *QOpenGLContext) ExtraFunctions() *QOpenGLExtraFunctions {
	__r, _, _ := purego.SyscallN(addr_QOpenGLContext_extraFunctions, this.h)
	runtime.KeepAlive(this)
	return bindQOpenGLExtraFunctions(__r)
}

func (this *QOpenGLContext) Extensions() []*QByteArray {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QOpenGLContext_extensions, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QByteArray, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQByteArray(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QOpenGLContext) HasExtension(extension *QByteArray) bool {
	__r0, _, _ := purego.SyscallN(addr_QOpenGLContext_hasExtension, this.h, uintptr(extension.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(extension)
	return __r != 0
}

func QOpenGLContext_OpenGLModuleType() QOpenGLContext__OpenGLModuleType {
	__r0, _, _ := purego.SyscallN(addr_QOpenGLContext_openGLModuleType)
	__r := int32(uint32(__r0))
	return QOpenGLContext__OpenGLModuleType(__r)
}

func (this *QOpenGLContext) IsOpenGLES() bool {
	__r0, _, _ := purego.SyscallN(addr_QOpenGLContext_isOpenGLES, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QOpenGLContext_SupportsThreadedOpenGL() bool {
	__r0, _, _ := purego.SyscallN(addr_QOpenGLContext_supportsThreadedOpenGL)
	__r := int32(uint32(__r0))
	return __r != 0
}

func QOpenGLContext_GlobalShareContext() *QOpenGLContext {
	__r, _, _ := purego.SyscallN(addr_QOpenGLContext_globalShareContext)
	return bindQOpenGLContext(__r)
}

func (this *QOpenGLContext) AboutToBeDestroyed() {
	purego.SyscallN(addr_QOpenGLContext_aboutToBeDestroyed, this.h)
	runtime.KeepAlive(this)
}

func QOpenGLContext_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QOpenGLContext_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QOpenGLContext_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QOpenGLContext_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QOpenGLContext) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QOpenGLContext_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QOpenGLContext) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QOpenGLContext_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QOpenGLContext) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QOpenGLContext_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QOpenGLContext) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QOpenGLContext_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QOpenGLContext_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QOpenGLContext_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QOpenGLContext_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QOpenGLContext_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QOpenGLContext_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QOpenGLContext_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QOpenGLContext_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QOpenGLContext_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QOpenGLContext) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QOpenGLContext_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QOpenGLContext) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QOpenGLContext_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QOpenGLContext) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QOpenGLContext_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QOpenGLContext) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QOpenGLContext_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QOpenGLContext) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QOpenGLContext_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QOpenGLContext) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QOpenGLContext_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QOpenGLContext) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QOpenGLContext_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QOpenGLContext) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QOpenGLContext_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QOpenGLContext) OnAboutToBeDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QOpenGLContext_connect_aboutToBeDestroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QOpenGLContext) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QOpenGLContext_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QOpenGLContext) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QOpenGLContext_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QOpenGLContext) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QOpenGLContext_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QOpenGLContext) Disconnect(handle int64) {
	pg_QOpenGLContext__disconnect(handle)
	unregisterCallback(handle)
}

func QOpenGLContext_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case 1237236456312824991:
		return 1
	case 5217623514630858297:
		return 2
	case -5409927581812374498:
		return 3
	case -1185315367526857833:
		return 4
	case 3769868229123868474:
		return 5
	case 7379029982170503008:
		return 6
	case -568369772982215962:
		return 7
	}
	return -1
}

func QOpenGLContext_enable(h uintptr, idx int32) { pg_QOpenGLContext__set_override(h, idx) }

func (this *QOpenGLContext) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QOpenGLContext) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QOpenGLContext) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QOpenGLContext) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QOpenGLContext) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QOpenGLContext) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QOpenGLContext) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QOpenGLContext) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QOpenGLContext) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QOpenGLContext_delete(this.h)
	}
	this.UnsafeRelease()
}
