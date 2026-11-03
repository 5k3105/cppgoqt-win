package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QMovie_base func(uintptr) uintptr
	addr_QMovie_new uintptr
	addr_QMovie_new2 uintptr
	addr_QMovie_new3 uintptr
	addr_QMovie_new4 uintptr
	addr_QMovie_new5 uintptr
	addr_QMovie_new6 uintptr
	addr_QMovie_new7 uintptr
	addr_QMovie_new8 uintptr
	addr_QMovie_tr uintptr
	addr_QMovie_supportedFormats uintptr
	addr_QMovie_setDevice uintptr
	addr_QMovie_device uintptr
	addr_QMovie_setFileName uintptr
	addr_QMovie_fileName uintptr
	addr_QMovie_setFormat uintptr
	addr_QMovie_format uintptr
	addr_QMovie_setBackgroundColor uintptr
	addr_QMovie_backgroundColor uintptr
	addr_QMovie_state uintptr
	addr_QMovie_frameRect uintptr
	addr_QMovie_currentImage uintptr
	addr_QMovie_currentPixmap uintptr
	addr_QMovie_isValid uintptr
	addr_QMovie_lastError uintptr
	addr_QMovie_lastErrorString uintptr
	addr_QMovie_jumpToFrame uintptr
	addr_QMovie_loopCount uintptr
	addr_QMovie_frameCount uintptr
	addr_QMovie_nextFrameDelay uintptr
	addr_QMovie_currentFrameNumber uintptr
	addr_QMovie_speed uintptr
	addr_QMovie_scaledSize uintptr
	addr_QMovie_setScaledSize uintptr
	addr_QMovie_cacheMode uintptr
	addr_QMovie_setCacheMode uintptr
	addr_QMovie_started uintptr
	addr_QMovie_resized uintptr
	addr_QMovie_updated uintptr
	addr_QMovie_stateChanged uintptr
	addr_QMovie_error uintptr
	addr_QMovie_finished uintptr
	addr_QMovie_frameChanged uintptr
	addr_QMovie_start uintptr
	addr_QMovie_jumpToNextFrame uintptr
	addr_QMovie_setPaused uintptr
	addr_QMovie_stop uintptr
	addr_QMovie_setSpeed uintptr
	addr_QMovie_connect uintptr
	addr_QMovie_connect2 uintptr
	addr_QMovie_sender uintptr
	addr_QMovie_senderSignalIndex uintptr
	addr_QMovie_receivers uintptr
	addr_QMovie_isSignalConnected uintptr
	addr_QMovie_tr2 uintptr
	addr_QMovie_tr3 uintptr
	addr_QMovie_connect3 uintptr
	addr_QMovie_connect4 uintptr
	addr_QMovie_metaObject_default uintptr
	addr_QMovie_event_default uintptr
	addr_QMovie_eventFilter_default uintptr
	addr_QMovie_timerEvent_default uintptr
	addr_QMovie_childEvent_default uintptr
	addr_QMovie_customEvent_default uintptr
	addr_QMovie_connectNotify_default uintptr
	addr_QMovie_disconnectNotify_default uintptr
	pg_QMovie_connect_started func(uintptr, int64)
	pg_QMovie_connect_resized func(uintptr, int64)
	pg_QMovie_connect_updated func(uintptr, int64)
	pg_QMovie_connect_stateChanged func(uintptr, int64)
	pg_QMovie_connect_error func(uintptr, int64)
	pg_QMovie_connect_finished func(uintptr, int64)
	pg_QMovie_connect_frameChanged func(uintptr, int64)
	pg_QMovie_connect_destroyed func(uintptr, int64)
	pg_QMovie_connect_destroyed2 func(uintptr, int64)
	pg_QMovie_connect_objectNameChanged func(uintptr, int64)
	pg_QMovie__disconnect func(int64)
	pg_QMovie__set_override func(uintptr, int32)
	pg_QMovie_delete func(uintptr)
)

func init() {
	reg(&pg_QMovie_base, "QMovie_base")
	symaddr(&addr_QMovie_new, "QMovie_new")
	symaddr(&addr_QMovie_new2, "QMovie_new2")
	symaddr(&addr_QMovie_new3, "QMovie_new3")
	symaddr(&addr_QMovie_new4, "QMovie_new4")
	symaddr(&addr_QMovie_new5, "QMovie_new5")
	symaddr(&addr_QMovie_new6, "QMovie_new6")
	symaddr(&addr_QMovie_new7, "QMovie_new7")
	symaddr(&addr_QMovie_new8, "QMovie_new8")
	symaddr(&addr_QMovie_tr, "QMovie_tr")
	symaddr(&addr_QMovie_supportedFormats, "QMovie_supportedFormats")
	symaddr(&addr_QMovie_setDevice, "QMovie_setDevice")
	symaddr(&addr_QMovie_device, "QMovie_device")
	symaddr(&addr_QMovie_setFileName, "QMovie_setFileName")
	symaddr(&addr_QMovie_fileName, "QMovie_fileName")
	symaddr(&addr_QMovie_setFormat, "QMovie_setFormat")
	symaddr(&addr_QMovie_format, "QMovie_format")
	symaddr(&addr_QMovie_setBackgroundColor, "QMovie_setBackgroundColor")
	symaddr(&addr_QMovie_backgroundColor, "QMovie_backgroundColor")
	symaddr(&addr_QMovie_state, "QMovie_state")
	symaddr(&addr_QMovie_frameRect, "QMovie_frameRect")
	symaddr(&addr_QMovie_currentImage, "QMovie_currentImage")
	symaddr(&addr_QMovie_currentPixmap, "QMovie_currentPixmap")
	symaddr(&addr_QMovie_isValid, "QMovie_isValid")
	symaddr(&addr_QMovie_lastError, "QMovie_lastError")
	symaddr(&addr_QMovie_lastErrorString, "QMovie_lastErrorString")
	symaddr(&addr_QMovie_jumpToFrame, "QMovie_jumpToFrame")
	symaddr(&addr_QMovie_loopCount, "QMovie_loopCount")
	symaddr(&addr_QMovie_frameCount, "QMovie_frameCount")
	symaddr(&addr_QMovie_nextFrameDelay, "QMovie_nextFrameDelay")
	symaddr(&addr_QMovie_currentFrameNumber, "QMovie_currentFrameNumber")
	symaddr(&addr_QMovie_speed, "QMovie_speed")
	symaddr(&addr_QMovie_scaledSize, "QMovie_scaledSize")
	symaddr(&addr_QMovie_setScaledSize, "QMovie_setScaledSize")
	symaddr(&addr_QMovie_cacheMode, "QMovie_cacheMode")
	symaddr(&addr_QMovie_setCacheMode, "QMovie_setCacheMode")
	symaddr(&addr_QMovie_started, "QMovie_started")
	symaddr(&addr_QMovie_resized, "QMovie_resized")
	symaddr(&addr_QMovie_updated, "QMovie_updated")
	symaddr(&addr_QMovie_stateChanged, "QMovie_stateChanged")
	symaddr(&addr_QMovie_error, "QMovie_error")
	symaddr(&addr_QMovie_finished, "QMovie_finished")
	symaddr(&addr_QMovie_frameChanged, "QMovie_frameChanged")
	symaddr(&addr_QMovie_start, "QMovie_start")
	symaddr(&addr_QMovie_jumpToNextFrame, "QMovie_jumpToNextFrame")
	symaddr(&addr_QMovie_setPaused, "QMovie_setPaused")
	symaddr(&addr_QMovie_stop, "QMovie_stop")
	symaddr(&addr_QMovie_setSpeed, "QMovie_setSpeed")
	symaddr(&addr_QMovie_connect, "QMovie_connect")
	symaddr(&addr_QMovie_connect2, "QMovie_connect2")
	symaddr(&addr_QMovie_sender, "QMovie_sender")
	symaddr(&addr_QMovie_senderSignalIndex, "QMovie_senderSignalIndex")
	symaddr(&addr_QMovie_receivers, "QMovie_receivers")
	symaddr(&addr_QMovie_isSignalConnected, "QMovie_isSignalConnected")
	symaddr(&addr_QMovie_tr2, "QMovie_tr2")
	symaddr(&addr_QMovie_tr3, "QMovie_tr3")
	symaddr(&addr_QMovie_connect3, "QMovie_connect3")
	symaddr(&addr_QMovie_connect4, "QMovie_connect4")
	symaddr(&addr_QMovie_metaObject_default, "QMovie_metaObject_default")
	symaddr(&addr_QMovie_event_default, "QMovie_event_default")
	symaddr(&addr_QMovie_eventFilter_default, "QMovie_eventFilter_default")
	symaddr(&addr_QMovie_timerEvent_default, "QMovie_timerEvent_default")
	symaddr(&addr_QMovie_childEvent_default, "QMovie_childEvent_default")
	symaddr(&addr_QMovie_customEvent_default, "QMovie_customEvent_default")
	symaddr(&addr_QMovie_connectNotify_default, "QMovie_connectNotify_default")
	symaddr(&addr_QMovie_disconnectNotify_default, "QMovie_disconnectNotify_default")
	reg(&pg_QMovie_connect_started, "QMovie_connect_started")
	reg(&pg_QMovie_connect_resized, "QMovie_connect_resized")
	reg(&pg_QMovie_connect_updated, "QMovie_connect_updated")
	reg(&pg_QMovie_connect_stateChanged, "QMovie_connect_stateChanged")
	reg(&pg_QMovie_connect_error, "QMovie_connect_error")
	reg(&pg_QMovie_connect_finished, "QMovie_connect_finished")
	reg(&pg_QMovie_connect_frameChanged, "QMovie_connect_frameChanged")
	reg(&pg_QMovie_connect_destroyed, "QMovie_connect_destroyed")
	reg(&pg_QMovie_connect_destroyed2, "QMovie_connect_destroyed2")
	reg(&pg_QMovie_connect_objectNameChanged, "QMovie_connect_objectNameChanged")
	reg(&pg_QMovie__disconnect, "QMovie__disconnect")
	reg(&pg_QMovie__set_override, "QMovie__set_override")
	reg(&pg_QMovie_delete, "QMovie_delete")
}

type QMovie struct {
	h uintptr
	*QObject
}

func bindQMovie(h uintptr) *QMovie {
	if h == 0 {
		return nil
	}
	return &QMovie{h: h, QObject: bindQObject(pg_QMovie_base(h))}
}

func newownedQMovie(h uintptr) *QMovie { return bindQMovie(h) }

func (this *QMovie) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QMovie) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQMovie(h unsafe.Pointer) *QMovie { return bindQMovie(uintptr(h)) }
func UnsafeNewQMovieOwned(h unsafe.Pointer) *QMovie { return newownedQMovie(uintptr(h)) }

func (this *QMovie) UnsafeRelease() {
	this.h = 0
	if this.QObject != nil {
		this.QObject.UnsafeRelease()
	}
}

func NewQMovie() *QMovie {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QMovie_new, uintptr(__goid))
	__o := newownedQMovie(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QMovie_virtIdx, QMovie_enable)
	return __o
}

func NewQMovie2(device *QIODevice) *QMovie {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QMovie_new2, uintptr(__goid), uintptr(device.UnsafePointer()))
	__o := newownedQMovie(__r)
	runtime.KeepAlive(device)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QMovie_virtIdx, QMovie_enable)
	return __o
}

func NewQMovie3(fileName string) *QMovie {
	__goid := nextVirtID()
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	__r, _, _ := purego.SyscallN(addr_QMovie_new3, uintptr(__goid), fileName_p, uintptr(len(fileName)))
	__o := newownedQMovie(__r)
	runtime.KeepAlive(fileName)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QMovie_virtIdx, QMovie_enable)
	return __o
}

func NewQMovie4(parent *QObject) *QMovie {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QMovie_new4, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQMovie(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QMovie_virtIdx, QMovie_enable)
	return __o
}

func NewQMovie5(device *QIODevice, format *QByteArray) *QMovie {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QMovie_new5, uintptr(__goid), uintptr(device.UnsafePointer()), uintptr(format.UnsafePointer()))
	__o := newownedQMovie(__r)
	runtime.KeepAlive(device)
	runtime.KeepAlive(format)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QMovie_virtIdx, QMovie_enable)
	return __o
}

func NewQMovie6(device *QIODevice, format *QByteArray, parent *QObject) *QMovie {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QMovie_new6, uintptr(__goid), uintptr(device.UnsafePointer()), uintptr(format.UnsafePointer()), uintptr(parent.UnsafePointer()))
	__o := newownedQMovie(__r)
	runtime.KeepAlive(device)
	runtime.KeepAlive(format)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QMovie_virtIdx, QMovie_enable)
	return __o
}

func NewQMovie7(fileName string, format *QByteArray) *QMovie {
	__goid := nextVirtID()
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	__r, _, _ := purego.SyscallN(addr_QMovie_new7, uintptr(__goid), fileName_p, uintptr(len(fileName)), uintptr(format.UnsafePointer()))
	__o := newownedQMovie(__r)
	runtime.KeepAlive(fileName)
	runtime.KeepAlive(format)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QMovie_virtIdx, QMovie_enable)
	return __o
}

func NewQMovie8(fileName string, format *QByteArray, parent *QObject) *QMovie {
	__goid := nextVirtID()
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	__r, _, _ := purego.SyscallN(addr_QMovie_new8, uintptr(__goid), fileName_p, uintptr(len(fileName)), uintptr(format.UnsafePointer()), uintptr(parent.UnsafePointer()))
	__o := newownedQMovie(__r)
	runtime.KeepAlive(fileName)
	runtime.KeepAlive(format)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QMovie_virtIdx, QMovie_enable)
	return __o
}

func QMovie_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QMovie_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func QMovie_SupportedFormats() []*QByteArray {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QMovie_supportedFormats, uintptr(unsafe.Pointer(&__n)))
	__out := make([]*QByteArray, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQByteArray(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QMovie) SetDevice(device *QIODevice) {
	purego.SyscallN(addr_QMovie_setDevice, this.h, uintptr(device.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(device)
}

func (this *QMovie) Device() *QIODevice {
	__r, _, _ := purego.SyscallN(addr_QMovie_device, this.h)
	runtime.KeepAlive(this)
	return bindQIODevice(__r)
}

func (this *QMovie) SetFileName(fileName string) {
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	purego.SyscallN(addr_QMovie_setFileName, this.h, fileName_p, uintptr(len(fileName)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(fileName)
}

func (this *QMovie) FileName() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QMovie_fileName, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QMovie) SetFormat(format *QByteArray) {
	purego.SyscallN(addr_QMovie_setFormat, this.h, uintptr(format.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(format)
}

func (this *QMovie) Format() *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QMovie_format, this.h)
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QMovie) SetBackgroundColor(color *QColor) {
	purego.SyscallN(addr_QMovie_setBackgroundColor, this.h, uintptr(color.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(color)
}

func (this *QMovie) BackgroundColor() *QColor {
	__r, _, _ := purego.SyscallN(addr_QMovie_backgroundColor, this.h)
	runtime.KeepAlive(this)
	return newownedQColor(__r)
}

func (this *QMovie) State() QMovie__MovieState {
	__r0, _, _ := purego.SyscallN(addr_QMovie_state, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QMovie__MovieState(__r)
}

func (this *QMovie) FrameRect() *QRect {
	__r, _, _ := purego.SyscallN(addr_QMovie_frameRect, this.h)
	runtime.KeepAlive(this)
	return newownedQRect(__r)
}

func (this *QMovie) CurrentImage() *QImage {
	__r, _, _ := purego.SyscallN(addr_QMovie_currentImage, this.h)
	runtime.KeepAlive(this)
	return newownedQImage(__r)
}

func (this *QMovie) CurrentPixmap() *QPixmap {
	__r, _, _ := purego.SyscallN(addr_QMovie_currentPixmap, this.h)
	runtime.KeepAlive(this)
	return newownedQPixmap(__r)
}

func (this *QMovie) IsValid() bool {
	__r0, _, _ := purego.SyscallN(addr_QMovie_isValid, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMovie) LastError() QImageReader__ImageReaderError {
	__r0, _, _ := purego.SyscallN(addr_QMovie_lastError, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QImageReader__ImageReaderError(__r)
}

func (this *QMovie) LastErrorString() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QMovie_lastErrorString, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QMovie) JumpToFrame(frameNumber int) bool {
	__r0, _, _ := purego.SyscallN(addr_QMovie_jumpToFrame, this.h, uintptr(frameNumber))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMovie) LoopCount() int {
	__r0, _, _ := purego.SyscallN(addr_QMovie_loopCount, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QMovie) FrameCount() int {
	__r0, _, _ := purego.SyscallN(addr_QMovie_frameCount, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QMovie) NextFrameDelay() int {
	__r0, _, _ := purego.SyscallN(addr_QMovie_nextFrameDelay, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QMovie) CurrentFrameNumber() int {
	__r0, _, _ := purego.SyscallN(addr_QMovie_currentFrameNumber, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QMovie) Speed() int {
	__r0, _, _ := purego.SyscallN(addr_QMovie_speed, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QMovie) ScaledSize() *QSize {
	__r, _, _ := purego.SyscallN(addr_QMovie_scaledSize, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QMovie) SetScaledSize(size *QSize) {
	purego.SyscallN(addr_QMovie_setScaledSize, this.h, uintptr(size.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(size)
}

func (this *QMovie) CacheMode() QMovie__CacheMode {
	__r0, _, _ := purego.SyscallN(addr_QMovie_cacheMode, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QMovie__CacheMode(__r)
}

func (this *QMovie) SetCacheMode(mode QMovie__CacheMode) {
	purego.SyscallN(addr_QMovie_setCacheMode, this.h, uintptr(mode))
	runtime.KeepAlive(this)
}

func (this *QMovie) Started() {
	purego.SyscallN(addr_QMovie_started, this.h)
	runtime.KeepAlive(this)
}

func (this *QMovie) Resized(size *QSize) {
	purego.SyscallN(addr_QMovie_resized, this.h, uintptr(size.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(size)
}

func (this *QMovie) Updated(rect *QRect) {
	purego.SyscallN(addr_QMovie_updated, this.h, uintptr(rect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
}

func (this *QMovie) StateChanged(state QMovie__MovieState) {
	purego.SyscallN(addr_QMovie_stateChanged, this.h, uintptr(state))
	runtime.KeepAlive(this)
}

func (this *QMovie) Error(error_ QImageReader__ImageReaderError) {
	purego.SyscallN(addr_QMovie_error, this.h, uintptr(error_))
	runtime.KeepAlive(this)
}

func (this *QMovie) Finished() {
	purego.SyscallN(addr_QMovie_finished, this.h)
	runtime.KeepAlive(this)
}

func (this *QMovie) FrameChanged(frameNumber int) {
	purego.SyscallN(addr_QMovie_frameChanged, this.h, uintptr(frameNumber))
	runtime.KeepAlive(this)
}

func (this *QMovie) Start() {
	purego.SyscallN(addr_QMovie_start, this.h)
	runtime.KeepAlive(this)
}

func (this *QMovie) JumpToNextFrame() bool {
	__r0, _, _ := purego.SyscallN(addr_QMovie_jumpToNextFrame, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMovie) SetPaused(paused bool) {
	purego.SyscallN(addr_QMovie_setPaused, this.h, uintptr(b2i(paused)))
	runtime.KeepAlive(this)
}

func (this *QMovie) Stop() {
	purego.SyscallN(addr_QMovie_stop, this.h)
	runtime.KeepAlive(this)
}

func (this *QMovie) SetSpeed(percentSpeed int) {
	purego.SyscallN(addr_QMovie_setSpeed, this.h, uintptr(percentSpeed))
	runtime.KeepAlive(this)
}

func QMovie_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QMovie_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QMovie_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QMovie_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QMovie) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QMovie_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QMovie) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QMovie_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QMovie) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QMovie_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QMovie) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QMovie_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QMovie_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QMovie_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QMovie_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QMovie_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QMovie_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QMovie_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QMovie_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QMovie_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QMovie) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QMovie_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QMovie) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QMovie_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QMovie) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QMovie_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QMovie) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QMovie_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMovie) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QMovie_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMovie) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QMovie_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMovie) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QMovie_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QMovie) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QMovie_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QMovie) OnStarted(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QMovie_connect_started(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QMovie) OnResized(fn func(*QSize)) int64 {
	id := registerCallback(func(size uintptr) { fn(UnsafeNewQSize(unsafe.Pointer(size))) })
	pg_QMovie_connect_resized(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QMovie) OnUpdated(fn func(*QRect)) int64 {
	id := registerCallback(func(rect uintptr) { fn(UnsafeNewQRect(unsafe.Pointer(rect))) })
	pg_QMovie_connect_updated(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QMovie) OnStateChanged(fn func(QMovie__MovieState)) int64 {
	id := registerCallback(func(state int32) { fn(QMovie__MovieState(state)) })
	pg_QMovie_connect_stateChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QMovie) OnError(fn func(QImageReader__ImageReaderError)) int64 {
	id := registerCallback(func(error_ int32) { fn(QImageReader__ImageReaderError(error_)) })
	pg_QMovie_connect_error(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QMovie) OnFinished(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QMovie_connect_finished(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QMovie) OnFrameChanged(fn func(int)) int64 {
	id := registerCallback(func(frameNumber int32) { fn(int(frameNumber)) })
	pg_QMovie_connect_frameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QMovie) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QMovie_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QMovie) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QMovie_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QMovie) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QMovie_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QMovie) Disconnect(handle int64) {
	pg_QMovie__disconnect(handle)
	unregisterCallback(handle)
}

func QMovie_virtIdx(slot int64) int32 {
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

func QMovie_enable(h uintptr, idx int32) { pg_QMovie__set_override(h, idx) }

func (this *QMovie) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QMovie) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QMovie) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QMovie) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QMovie) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QMovie) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QMovie) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QMovie) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QMovie) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QMovie_delete(this.h)
	}
	this.UnsafeRelease()
}
