package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QRhiWidget_chain uintptr
	addr_QRhiWidget_new uintptr
	addr_QRhiWidget_new2 uintptr
	addr_QRhiWidget_new3 uintptr
	addr_QRhiWidget_tr uintptr
	addr_QRhiWidget_api uintptr
	addr_QRhiWidget_setApi uintptr
	addr_QRhiWidget_isDebugLayerEnabled uintptr
	addr_QRhiWidget_setDebugLayerEnabled uintptr
	addr_QRhiWidget_sampleCount uintptr
	addr_QRhiWidget_setSampleCount uintptr
	addr_QRhiWidget_colorBufferFormat uintptr
	addr_QRhiWidget_setColorBufferFormat uintptr
	addr_QRhiWidget_fixedColorBufferSize uintptr
	addr_QRhiWidget_setFixedColorBufferSize uintptr
	addr_QRhiWidget_setFixedColorBufferSize2 uintptr
	addr_QRhiWidget_isMirrorVerticallyEnabled uintptr
	addr_QRhiWidget_setMirrorVertically uintptr
	addr_QRhiWidget_grabFramebuffer uintptr
	addr_QRhiWidget_isAutoRenderTargetEnabled uintptr
	addr_QRhiWidget_setAutoRenderTarget uintptr
	addr_QRhiWidget_frameSubmitted uintptr
	addr_QRhiWidget_renderFailed uintptr
	addr_QRhiWidget_sampleCountChanged uintptr
	addr_QRhiWidget_colorBufferFormatChanged uintptr
	addr_QRhiWidget_fixedColorBufferSizeChanged uintptr
	addr_QRhiWidget_mirrorVerticallyChanged uintptr
	addr_QRhiWidget_setTabOrder uintptr
	addr_QRhiWidget_mouseGrabber uintptr
	addr_QRhiWidget_keyboardGrabber uintptr
	addr_QRhiWidget_find uintptr
	addr_QRhiWidget_createWindowContainer uintptr
	addr_QRhiWidget_updateMicroFocus uintptr
	addr_QRhiWidget_create uintptr
	addr_QRhiWidget_destroy uintptr
	addr_QRhiWidget_focusNextChild uintptr
	addr_QRhiWidget_focusPreviousChild uintptr
	addr_QRhiWidget_connect uintptr
	addr_QRhiWidget_connect2 uintptr
	addr_QRhiWidget_sender uintptr
	addr_QRhiWidget_senderSignalIndex uintptr
	addr_QRhiWidget_receivers uintptr
	addr_QRhiWidget_isSignalConnected uintptr
	pg_QRhiWidget_devicePixelRatioFScale func() float64
	pg_QRhiWidget_encodeMetricF func(int32, float64) int32
	pg_QRhiWidget_getDecodedMetricF func(uintptr, int32, int32) float64
	addr_QRhiWidget_tr2 uintptr
	addr_QRhiWidget_tr3 uintptr
	addr_QRhiWidget_createWindowContainer2 uintptr
	addr_QRhiWidget_createWindowContainer3 uintptr
	addr_QRhiWidget_updateMicroFocus2 uintptr
	addr_QRhiWidget_create2 uintptr
	addr_QRhiWidget_create3 uintptr
	addr_QRhiWidget_create4 uintptr
	addr_QRhiWidget_destroy2 uintptr
	addr_QRhiWidget_destroy3 uintptr
	addr_QRhiWidget_connect3 uintptr
	addr_QRhiWidget_connect4 uintptr
	addr_QRhiWidget_metaObject_default uintptr
	addr_QRhiWidget_releaseResources_default uintptr
	addr_QRhiWidget_resizeEvent_default uintptr
	addr_QRhiWidget_paintEvent_default uintptr
	addr_QRhiWidget_event_default uintptr
	addr_QRhiWidget_devType_default uintptr
	addr_QRhiWidget_setVisible_default uintptr
	addr_QRhiWidget_sizeHint_default uintptr
	addr_QRhiWidget_minimumSizeHint_default uintptr
	addr_QRhiWidget_heightForWidth_default uintptr
	addr_QRhiWidget_hasHeightForWidth_default uintptr
	addr_QRhiWidget_paintEngine_default uintptr
	addr_QRhiWidget_mousePressEvent_default uintptr
	addr_QRhiWidget_mouseReleaseEvent_default uintptr
	addr_QRhiWidget_mouseDoubleClickEvent_default uintptr
	addr_QRhiWidget_mouseMoveEvent_default uintptr
	addr_QRhiWidget_wheelEvent_default uintptr
	addr_QRhiWidget_keyPressEvent_default uintptr
	addr_QRhiWidget_keyReleaseEvent_default uintptr
	addr_QRhiWidget_focusInEvent_default uintptr
	addr_QRhiWidget_focusOutEvent_default uintptr
	addr_QRhiWidget_enterEvent_default uintptr
	addr_QRhiWidget_leaveEvent_default uintptr
	addr_QRhiWidget_moveEvent_default uintptr
	addr_QRhiWidget_closeEvent_default uintptr
	addr_QRhiWidget_contextMenuEvent_default uintptr
	addr_QRhiWidget_tabletEvent_default uintptr
	addr_QRhiWidget_actionEvent_default uintptr
	addr_QRhiWidget_dragEnterEvent_default uintptr
	addr_QRhiWidget_dragMoveEvent_default uintptr
	addr_QRhiWidget_dragLeaveEvent_default uintptr
	addr_QRhiWidget_dropEvent_default uintptr
	addr_QRhiWidget_showEvent_default uintptr
	addr_QRhiWidget_hideEvent_default uintptr
	addr_QRhiWidget_changeEvent_default uintptr
	addr_QRhiWidget_metric_default uintptr
	addr_QRhiWidget_initPainter_default uintptr
	addr_QRhiWidget_redirected_default uintptr
	addr_QRhiWidget_sharedPainter_default uintptr
	addr_QRhiWidget_inputMethodEvent_default uintptr
	addr_QRhiWidget_inputMethodQuery_default uintptr
	addr_QRhiWidget_focusNextPrevChild_default uintptr
	addr_QRhiWidget_eventFilter_default uintptr
	addr_QRhiWidget_timerEvent_default uintptr
	addr_QRhiWidget_childEvent_default uintptr
	addr_QRhiWidget_customEvent_default uintptr
	addr_QRhiWidget_connectNotify_default uintptr
	addr_QRhiWidget_disconnectNotify_default uintptr
	pg_QRhiWidget_connect_frameSubmitted func(uintptr, int64)
	pg_QRhiWidget_connect_renderFailed func(uintptr, int64)
	pg_QRhiWidget_connect_sampleCountChanged func(uintptr, int64)
	pg_QRhiWidget_connect_colorBufferFormatChanged func(uintptr, int64)
	pg_QRhiWidget_connect_fixedColorBufferSizeChanged func(uintptr, int64)
	pg_QRhiWidget_connect_mirrorVerticallyChanged func(uintptr, int64)
	pg_QRhiWidget_connect_windowTitleChanged func(uintptr, int64)
	pg_QRhiWidget_connect_windowIconChanged func(uintptr, int64)
	pg_QRhiWidget_connect_windowIconTextChanged func(uintptr, int64)
	pg_QRhiWidget_connect_customContextMenuRequested func(uintptr, int64)
	pg_QRhiWidget_connect_destroyed func(uintptr, int64)
	pg_QRhiWidget_connect_destroyed2 func(uintptr, int64)
	pg_QRhiWidget_connect_objectNameChanged func(uintptr, int64)
	pg_QRhiWidget__disconnect func(int64)
	pg_QRhiWidget__set_override func(uintptr, int32)
	pg_QRhiWidget_delete func(uintptr)
)

func init() {
	symaddr(&addr_QRhiWidget_chain, "QRhiWidget_chain")
	symaddr(&addr_QRhiWidget_new, "QRhiWidget_new")
	symaddr(&addr_QRhiWidget_new2, "QRhiWidget_new2")
	symaddr(&addr_QRhiWidget_new3, "QRhiWidget_new3")
	symaddr(&addr_QRhiWidget_tr, "QRhiWidget_tr")
	symaddr(&addr_QRhiWidget_api, "QRhiWidget_api")
	symaddr(&addr_QRhiWidget_setApi, "QRhiWidget_setApi")
	symaddr(&addr_QRhiWidget_isDebugLayerEnabled, "QRhiWidget_isDebugLayerEnabled")
	symaddr(&addr_QRhiWidget_setDebugLayerEnabled, "QRhiWidget_setDebugLayerEnabled")
	symaddr(&addr_QRhiWidget_sampleCount, "QRhiWidget_sampleCount")
	symaddr(&addr_QRhiWidget_setSampleCount, "QRhiWidget_setSampleCount")
	symaddr(&addr_QRhiWidget_colorBufferFormat, "QRhiWidget_colorBufferFormat")
	symaddr(&addr_QRhiWidget_setColorBufferFormat, "QRhiWidget_setColorBufferFormat")
	symaddr(&addr_QRhiWidget_fixedColorBufferSize, "QRhiWidget_fixedColorBufferSize")
	symaddr(&addr_QRhiWidget_setFixedColorBufferSize, "QRhiWidget_setFixedColorBufferSize")
	symaddr(&addr_QRhiWidget_setFixedColorBufferSize2, "QRhiWidget_setFixedColorBufferSize2")
	symaddr(&addr_QRhiWidget_isMirrorVerticallyEnabled, "QRhiWidget_isMirrorVerticallyEnabled")
	symaddr(&addr_QRhiWidget_setMirrorVertically, "QRhiWidget_setMirrorVertically")
	symaddr(&addr_QRhiWidget_grabFramebuffer, "QRhiWidget_grabFramebuffer")
	symaddr(&addr_QRhiWidget_isAutoRenderTargetEnabled, "QRhiWidget_isAutoRenderTargetEnabled")
	symaddr(&addr_QRhiWidget_setAutoRenderTarget, "QRhiWidget_setAutoRenderTarget")
	symaddr(&addr_QRhiWidget_frameSubmitted, "QRhiWidget_frameSubmitted")
	symaddr(&addr_QRhiWidget_renderFailed, "QRhiWidget_renderFailed")
	symaddr(&addr_QRhiWidget_sampleCountChanged, "QRhiWidget_sampleCountChanged")
	symaddr(&addr_QRhiWidget_colorBufferFormatChanged, "QRhiWidget_colorBufferFormatChanged")
	symaddr(&addr_QRhiWidget_fixedColorBufferSizeChanged, "QRhiWidget_fixedColorBufferSizeChanged")
	symaddr(&addr_QRhiWidget_mirrorVerticallyChanged, "QRhiWidget_mirrorVerticallyChanged")
	symaddr(&addr_QRhiWidget_setTabOrder, "QRhiWidget_setTabOrder")
	symaddr(&addr_QRhiWidget_mouseGrabber, "QRhiWidget_mouseGrabber")
	symaddr(&addr_QRhiWidget_keyboardGrabber, "QRhiWidget_keyboardGrabber")
	symaddr(&addr_QRhiWidget_find, "QRhiWidget_find")
	symaddr(&addr_QRhiWidget_createWindowContainer, "QRhiWidget_createWindowContainer")
	symaddr(&addr_QRhiWidget_updateMicroFocus, "QRhiWidget_updateMicroFocus")
	symaddr(&addr_QRhiWidget_create, "QRhiWidget_create")
	symaddr(&addr_QRhiWidget_destroy, "QRhiWidget_destroy")
	symaddr(&addr_QRhiWidget_focusNextChild, "QRhiWidget_focusNextChild")
	symaddr(&addr_QRhiWidget_focusPreviousChild, "QRhiWidget_focusPreviousChild")
	symaddr(&addr_QRhiWidget_connect, "QRhiWidget_connect")
	symaddr(&addr_QRhiWidget_connect2, "QRhiWidget_connect2")
	symaddr(&addr_QRhiWidget_sender, "QRhiWidget_sender")
	symaddr(&addr_QRhiWidget_senderSignalIndex, "QRhiWidget_senderSignalIndex")
	symaddr(&addr_QRhiWidget_receivers, "QRhiWidget_receivers")
	symaddr(&addr_QRhiWidget_isSignalConnected, "QRhiWidget_isSignalConnected")
	reg(&pg_QRhiWidget_devicePixelRatioFScale, "QRhiWidget_devicePixelRatioFScale")
	reg(&pg_QRhiWidget_encodeMetricF, "QRhiWidget_encodeMetricF")
	reg(&pg_QRhiWidget_getDecodedMetricF, "QRhiWidget_getDecodedMetricF")
	symaddr(&addr_QRhiWidget_tr2, "QRhiWidget_tr2")
	symaddr(&addr_QRhiWidget_tr3, "QRhiWidget_tr3")
	symaddr(&addr_QRhiWidget_createWindowContainer2, "QRhiWidget_createWindowContainer2")
	symaddr(&addr_QRhiWidget_createWindowContainer3, "QRhiWidget_createWindowContainer3")
	symaddr(&addr_QRhiWidget_updateMicroFocus2, "QRhiWidget_updateMicroFocus2")
	symaddr(&addr_QRhiWidget_create2, "QRhiWidget_create2")
	symaddr(&addr_QRhiWidget_create3, "QRhiWidget_create3")
	symaddr(&addr_QRhiWidget_create4, "QRhiWidget_create4")
	symaddr(&addr_QRhiWidget_destroy2, "QRhiWidget_destroy2")
	symaddr(&addr_QRhiWidget_destroy3, "QRhiWidget_destroy3")
	symaddr(&addr_QRhiWidget_connect3, "QRhiWidget_connect3")
	symaddr(&addr_QRhiWidget_connect4, "QRhiWidget_connect4")
	symaddr(&addr_QRhiWidget_metaObject_default, "QRhiWidget_metaObject_default")
	symaddr(&addr_QRhiWidget_releaseResources_default, "QRhiWidget_releaseResources_default")
	symaddr(&addr_QRhiWidget_resizeEvent_default, "QRhiWidget_resizeEvent_default")
	symaddr(&addr_QRhiWidget_paintEvent_default, "QRhiWidget_paintEvent_default")
	symaddr(&addr_QRhiWidget_event_default, "QRhiWidget_event_default")
	symaddr(&addr_QRhiWidget_devType_default, "QRhiWidget_devType_default")
	symaddr(&addr_QRhiWidget_setVisible_default, "QRhiWidget_setVisible_default")
	symaddr(&addr_QRhiWidget_sizeHint_default, "QRhiWidget_sizeHint_default")
	symaddr(&addr_QRhiWidget_minimumSizeHint_default, "QRhiWidget_minimumSizeHint_default")
	symaddr(&addr_QRhiWidget_heightForWidth_default, "QRhiWidget_heightForWidth_default")
	symaddr(&addr_QRhiWidget_hasHeightForWidth_default, "QRhiWidget_hasHeightForWidth_default")
	symaddr(&addr_QRhiWidget_paintEngine_default, "QRhiWidget_paintEngine_default")
	symaddr(&addr_QRhiWidget_mousePressEvent_default, "QRhiWidget_mousePressEvent_default")
	symaddr(&addr_QRhiWidget_mouseReleaseEvent_default, "QRhiWidget_mouseReleaseEvent_default")
	symaddr(&addr_QRhiWidget_mouseDoubleClickEvent_default, "QRhiWidget_mouseDoubleClickEvent_default")
	symaddr(&addr_QRhiWidget_mouseMoveEvent_default, "QRhiWidget_mouseMoveEvent_default")
	symaddr(&addr_QRhiWidget_wheelEvent_default, "QRhiWidget_wheelEvent_default")
	symaddr(&addr_QRhiWidget_keyPressEvent_default, "QRhiWidget_keyPressEvent_default")
	symaddr(&addr_QRhiWidget_keyReleaseEvent_default, "QRhiWidget_keyReleaseEvent_default")
	symaddr(&addr_QRhiWidget_focusInEvent_default, "QRhiWidget_focusInEvent_default")
	symaddr(&addr_QRhiWidget_focusOutEvent_default, "QRhiWidget_focusOutEvent_default")
	symaddr(&addr_QRhiWidget_enterEvent_default, "QRhiWidget_enterEvent_default")
	symaddr(&addr_QRhiWidget_leaveEvent_default, "QRhiWidget_leaveEvent_default")
	symaddr(&addr_QRhiWidget_moveEvent_default, "QRhiWidget_moveEvent_default")
	symaddr(&addr_QRhiWidget_closeEvent_default, "QRhiWidget_closeEvent_default")
	symaddr(&addr_QRhiWidget_contextMenuEvent_default, "QRhiWidget_contextMenuEvent_default")
	symaddr(&addr_QRhiWidget_tabletEvent_default, "QRhiWidget_tabletEvent_default")
	symaddr(&addr_QRhiWidget_actionEvent_default, "QRhiWidget_actionEvent_default")
	symaddr(&addr_QRhiWidget_dragEnterEvent_default, "QRhiWidget_dragEnterEvent_default")
	symaddr(&addr_QRhiWidget_dragMoveEvent_default, "QRhiWidget_dragMoveEvent_default")
	symaddr(&addr_QRhiWidget_dragLeaveEvent_default, "QRhiWidget_dragLeaveEvent_default")
	symaddr(&addr_QRhiWidget_dropEvent_default, "QRhiWidget_dropEvent_default")
	symaddr(&addr_QRhiWidget_showEvent_default, "QRhiWidget_showEvent_default")
	symaddr(&addr_QRhiWidget_hideEvent_default, "QRhiWidget_hideEvent_default")
	symaddr(&addr_QRhiWidget_changeEvent_default, "QRhiWidget_changeEvent_default")
	symaddr(&addr_QRhiWidget_metric_default, "QRhiWidget_metric_default")
	symaddr(&addr_QRhiWidget_initPainter_default, "QRhiWidget_initPainter_default")
	symaddr(&addr_QRhiWidget_redirected_default, "QRhiWidget_redirected_default")
	symaddr(&addr_QRhiWidget_sharedPainter_default, "QRhiWidget_sharedPainter_default")
	symaddr(&addr_QRhiWidget_inputMethodEvent_default, "QRhiWidget_inputMethodEvent_default")
	symaddr(&addr_QRhiWidget_inputMethodQuery_default, "QRhiWidget_inputMethodQuery_default")
	symaddr(&addr_QRhiWidget_focusNextPrevChild_default, "QRhiWidget_focusNextPrevChild_default")
	symaddr(&addr_QRhiWidget_eventFilter_default, "QRhiWidget_eventFilter_default")
	symaddr(&addr_QRhiWidget_timerEvent_default, "QRhiWidget_timerEvent_default")
	symaddr(&addr_QRhiWidget_childEvent_default, "QRhiWidget_childEvent_default")
	symaddr(&addr_QRhiWidget_customEvent_default, "QRhiWidget_customEvent_default")
	symaddr(&addr_QRhiWidget_connectNotify_default, "QRhiWidget_connectNotify_default")
	symaddr(&addr_QRhiWidget_disconnectNotify_default, "QRhiWidget_disconnectNotify_default")
	reg(&pg_QRhiWidget_connect_frameSubmitted, "QRhiWidget_connect_frameSubmitted")
	reg(&pg_QRhiWidget_connect_renderFailed, "QRhiWidget_connect_renderFailed")
	reg(&pg_QRhiWidget_connect_sampleCountChanged, "QRhiWidget_connect_sampleCountChanged")
	reg(&pg_QRhiWidget_connect_colorBufferFormatChanged, "QRhiWidget_connect_colorBufferFormatChanged")
	reg(&pg_QRhiWidget_connect_fixedColorBufferSizeChanged, "QRhiWidget_connect_fixedColorBufferSizeChanged")
	reg(&pg_QRhiWidget_connect_mirrorVerticallyChanged, "QRhiWidget_connect_mirrorVerticallyChanged")
	reg(&pg_QRhiWidget_connect_windowTitleChanged, "QRhiWidget_connect_windowTitleChanged")
	reg(&pg_QRhiWidget_connect_windowIconChanged, "QRhiWidget_connect_windowIconChanged")
	reg(&pg_QRhiWidget_connect_windowIconTextChanged, "QRhiWidget_connect_windowIconTextChanged")
	reg(&pg_QRhiWidget_connect_customContextMenuRequested, "QRhiWidget_connect_customContextMenuRequested")
	reg(&pg_QRhiWidget_connect_destroyed, "QRhiWidget_connect_destroyed")
	reg(&pg_QRhiWidget_connect_destroyed2, "QRhiWidget_connect_destroyed2")
	reg(&pg_QRhiWidget_connect_objectNameChanged, "QRhiWidget_connect_objectNameChanged")
	reg(&pg_QRhiWidget__disconnect, "QRhiWidget__disconnect")
	reg(&pg_QRhiWidget__set_override, "QRhiWidget__set_override")
	reg(&pg_QRhiWidget_delete, "QRhiWidget_delete")
}

type QRhiWidget struct {
	h uintptr
	*QWidget
}

func bindQRhiWidget(h uintptr) *QRhiWidget {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QRhiWidget_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QRhiWidget{h: h, QWidget: &QWidget{h: __c[0], QObject: &QObject{h: __c[1]}}}
}

func newownedQRhiWidget(h uintptr) *QRhiWidget { return bindQRhiWidget(h) }

func (this *QRhiWidget) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QRhiWidget) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQRhiWidget(h unsafe.Pointer) *QRhiWidget { return bindQRhiWidget(uintptr(h)) }
func UnsafeNewQRhiWidgetOwned(h unsafe.Pointer) *QRhiWidget { return newownedQRhiWidget(uintptr(h)) }

func (this *QRhiWidget) UnsafeRelease() {
	this.h = 0
	if this.QWidget != nil {
		this.QWidget.UnsafeRelease()
	}
}

func NewQRhiWidget(parent *QWidget) *QRhiWidget {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QRhiWidget_new, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQRhiWidget(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QRhiWidget_virtIdx, QRhiWidget_enable)
	return __o
}

func NewQRhiWidget2() *QRhiWidget {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QRhiWidget_new2, uintptr(__goid))
	__o := newownedQRhiWidget(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QRhiWidget_virtIdx, QRhiWidget_enable)
	return __o
}

func NewQRhiWidget3(parent *QWidget, f WindowType) *QRhiWidget {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QRhiWidget_new3, uintptr(__goid), uintptr(parent.UnsafePointer()), uintptr(f))
	__o := newownedQRhiWidget(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QRhiWidget_virtIdx, QRhiWidget_enable)
	return __o
}

func QRhiWidget_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QRhiWidget_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QRhiWidget) Api() QRhiWidget__Api {
	__r0, _, _ := purego.SyscallN(addr_QRhiWidget_api, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QRhiWidget__Api(__r)
}

func (this *QRhiWidget) SetApi(api QRhiWidget__Api) {
	purego.SyscallN(addr_QRhiWidget_setApi, this.h, uintptr(api))
	runtime.KeepAlive(this)
}

func (this *QRhiWidget) IsDebugLayerEnabled() bool {
	__r0, _, _ := purego.SyscallN(addr_QRhiWidget_isDebugLayerEnabled, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QRhiWidget) SetDebugLayerEnabled(enable bool) {
	purego.SyscallN(addr_QRhiWidget_setDebugLayerEnabled, this.h, uintptr(b2i(enable)))
	runtime.KeepAlive(this)
}

func (this *QRhiWidget) SampleCount() int {
	__r0, _, _ := purego.SyscallN(addr_QRhiWidget_sampleCount, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QRhiWidget) SetSampleCount(samples int) {
	purego.SyscallN(addr_QRhiWidget_setSampleCount, this.h, uintptr(samples))
	runtime.KeepAlive(this)
}

func (this *QRhiWidget) ColorBufferFormat() QRhiWidget__TextureFormat {
	__r0, _, _ := purego.SyscallN(addr_QRhiWidget_colorBufferFormat, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QRhiWidget__TextureFormat(__r)
}

func (this *QRhiWidget) SetColorBufferFormat(format QRhiWidget__TextureFormat) {
	purego.SyscallN(addr_QRhiWidget_setColorBufferFormat, this.h, uintptr(format))
	runtime.KeepAlive(this)
}

func (this *QRhiWidget) FixedColorBufferSize() *QSize {
	__r, _, _ := purego.SyscallN(addr_QRhiWidget_fixedColorBufferSize, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QRhiWidget) SetFixedColorBufferSize(pixelSize *QSize) {
	purego.SyscallN(addr_QRhiWidget_setFixedColorBufferSize, this.h, uintptr(pixelSize.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pixelSize)
}

func (this *QRhiWidget) SetFixedColorBufferSize2(w int, h int) {
	purego.SyscallN(addr_QRhiWidget_setFixedColorBufferSize2, this.h, uintptr(w), uintptr(h))
	runtime.KeepAlive(this)
}

func (this *QRhiWidget) IsMirrorVerticallyEnabled() bool {
	__r0, _, _ := purego.SyscallN(addr_QRhiWidget_isMirrorVerticallyEnabled, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QRhiWidget) SetMirrorVertically(enabled bool) {
	purego.SyscallN(addr_QRhiWidget_setMirrorVertically, this.h, uintptr(b2i(enabled)))
	runtime.KeepAlive(this)
}

func (this *QRhiWidget) GrabFramebuffer() *QImage {
	__r, _, _ := purego.SyscallN(addr_QRhiWidget_grabFramebuffer, this.h)
	runtime.KeepAlive(this)
	return newownedQImage(__r)
}

func (this *QRhiWidget) IsAutoRenderTargetEnabled() bool {
	__r0, _, _ := purego.SyscallN(addr_QRhiWidget_isAutoRenderTargetEnabled, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QRhiWidget) SetAutoRenderTarget(enabled bool) {
	purego.SyscallN(addr_QRhiWidget_setAutoRenderTarget, this.h, uintptr(b2i(enabled)))
	runtime.KeepAlive(this)
}

func (this *QRhiWidget) FrameSubmitted() {
	purego.SyscallN(addr_QRhiWidget_frameSubmitted, this.h)
	runtime.KeepAlive(this)
}

func (this *QRhiWidget) RenderFailed() {
	purego.SyscallN(addr_QRhiWidget_renderFailed, this.h)
	runtime.KeepAlive(this)
}

func (this *QRhiWidget) SampleCountChanged(samples int) {
	purego.SyscallN(addr_QRhiWidget_sampleCountChanged, this.h, uintptr(samples))
	runtime.KeepAlive(this)
}

func (this *QRhiWidget) ColorBufferFormatChanged(format QRhiWidget__TextureFormat) {
	purego.SyscallN(addr_QRhiWidget_colorBufferFormatChanged, this.h, uintptr(format))
	runtime.KeepAlive(this)
}

func (this *QRhiWidget) FixedColorBufferSizeChanged(pixelSize *QSize) {
	purego.SyscallN(addr_QRhiWidget_fixedColorBufferSizeChanged, this.h, uintptr(pixelSize.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pixelSize)
}

func (this *QRhiWidget) MirrorVerticallyChanged(enabled bool) {
	purego.SyscallN(addr_QRhiWidget_mirrorVerticallyChanged, this.h, uintptr(b2i(enabled)))
	runtime.KeepAlive(this)
}

func QRhiWidget_SetTabOrder(arg1 *QWidget, arg2 *QWidget) {
	purego.SyscallN(addr_QRhiWidget_setTabOrder, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()))
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
}

func QRhiWidget_MouseGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QRhiWidget_mouseGrabber)
	return bindQWidget(__r)
}

func QRhiWidget_KeyboardGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QRhiWidget_keyboardGrabber)
	return bindQWidget(__r)
}

func QRhiWidget_Find(arg1 uint64) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QRhiWidget_find, uintptr(arg1))
	return bindQWidget(__r)
}

func QRhiWidget_CreateWindowContainer(window *QWindow) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QRhiWidget_createWindowContainer, uintptr(window.UnsafePointer()))
	runtime.KeepAlive(window)
	return bindQWidget(__r)
}

func (this *QRhiWidget) UpdateMicroFocus() {
	purego.SyscallN(addr_QRhiWidget_updateMicroFocus, this.h)
	runtime.KeepAlive(this)
}

func (this *QRhiWidget) Create() {
	purego.SyscallN(addr_QRhiWidget_create, this.h)
	runtime.KeepAlive(this)
}

func (this *QRhiWidget) Destroy() {
	purego.SyscallN(addr_QRhiWidget_destroy, this.h)
	runtime.KeepAlive(this)
}

func (this *QRhiWidget) FocusNextChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QRhiWidget_focusNextChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QRhiWidget) FocusPreviousChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QRhiWidget_focusPreviousChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QRhiWidget_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QRhiWidget_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QRhiWidget_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QRhiWidget_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QRhiWidget) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QRhiWidget_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QRhiWidget) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QRhiWidget_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QRhiWidget) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QRhiWidget_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QRhiWidget) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QRhiWidget_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QRhiWidget_DevicePixelRatioFScale() float64 {
	__r := pg_QRhiWidget_devicePixelRatioFScale()
	return __r
}

func QRhiWidget_EncodeMetricF(metric QPaintDevice__PaintDeviceMetric, value float64) int {
	__r := pg_QRhiWidget_encodeMetricF(int32(metric), value)
	return int(__r)
}

func (this *QRhiWidget) GetDecodedMetricF(metricA QPaintDevice__PaintDeviceMetric, metricB QPaintDevice__PaintDeviceMetric) float64 {
	__r := pg_QRhiWidget_getDecodedMetricF(this.h, int32(metricA), int32(metricB))
	runtime.KeepAlive(this)
	return __r
}

func QRhiWidget_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QRhiWidget_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QRhiWidget_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QRhiWidget_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QRhiWidget_CreateWindowContainer2(window *QWindow, parent *QWidget) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QRhiWidget_createWindowContainer2, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func QRhiWidget_CreateWindowContainer3(window *QWindow, parent *QWidget, flags WindowType) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QRhiWidget_createWindowContainer3, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()), uintptr(flags))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func (this *QRhiWidget) UpdateMicroFocusWithQuery(query InputMethodQuery) {
	purego.SyscallN(addr_QRhiWidget_updateMicroFocus2, this.h, uintptr(query))
	runtime.KeepAlive(this)
}

func (this *QRhiWidget) CreateWith(arg1 uint64) {
	purego.SyscallN(addr_QRhiWidget_create2, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QRhiWidget) Create2(arg1 uint64, initializeWindow bool) {
	purego.SyscallN(addr_QRhiWidget_create3, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)))
	runtime.KeepAlive(this)
}

func (this *QRhiWidget) Create3(arg1 uint64, initializeWindow bool, destroyOldWindow bool) {
	purego.SyscallN(addr_QRhiWidget_create4, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)), uintptr(b2i(destroyOldWindow)))
	runtime.KeepAlive(this)
}

func (this *QRhiWidget) DestroyWithDestroyWindow(destroyWindow bool) {
	purego.SyscallN(addr_QRhiWidget_destroy2, this.h, uintptr(b2i(destroyWindow)))
	runtime.KeepAlive(this)
}

func (this *QRhiWidget) Destroy2(destroyWindow bool, destroySubWindows bool) {
	purego.SyscallN(addr_QRhiWidget_destroy3, this.h, uintptr(b2i(destroyWindow)), uintptr(b2i(destroySubWindows)))
	runtime.KeepAlive(this)
}

func QRhiWidget_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QRhiWidget_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QRhiWidget_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QRhiWidget_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QRhiWidget) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QRhiWidget_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QRhiWidget) ReleaseResourcesDefault() {
	purego.SyscallN(addr_QRhiWidget_releaseResources_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QRhiWidget) ResizeEventDefault(e *QResizeEvent) {
	purego.SyscallN(addr_QRhiWidget_resizeEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QRhiWidget) PaintEventDefault(e *QPaintEvent) {
	purego.SyscallN(addr_QRhiWidget_paintEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QRhiWidget) EventDefault(e *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QRhiWidget_event_default, this.h, uintptr(e.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
	return __r != 0
}

func (this *QRhiWidget) DevTypeDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QRhiWidget_devType_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QRhiWidget) SetVisibleDefault(visible bool) {
	purego.SyscallN(addr_QRhiWidget_setVisible_default, this.h, uintptr(b2i(visible)))
	runtime.KeepAlive(this)
}

func (this *QRhiWidget) SizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QRhiWidget_sizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QRhiWidget) MinimumSizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QRhiWidget_minimumSizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QRhiWidget) HeightForWidthDefault(arg1 int) int {
	__r0, _, _ := purego.SyscallN(addr_QRhiWidget_heightForWidth_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QRhiWidget) HasHeightForWidthDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QRhiWidget_hasHeightForWidth_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QRhiWidget) PaintEngineDefault() *QPaintEngine {
	__r, _, _ := purego.SyscallN(addr_QRhiWidget_paintEngine_default, this.h)
	runtime.KeepAlive(this)
	return bindQPaintEngine(__r)
}

func (this *QRhiWidget) MousePressEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QRhiWidget_mousePressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QRhiWidget) MouseReleaseEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QRhiWidget_mouseReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QRhiWidget) MouseDoubleClickEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QRhiWidget_mouseDoubleClickEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QRhiWidget) MouseMoveEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QRhiWidget_mouseMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QRhiWidget) WheelEventDefault(event *QWheelEvent) {
	purego.SyscallN(addr_QRhiWidget_wheelEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QRhiWidget) KeyPressEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QRhiWidget_keyPressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QRhiWidget) KeyReleaseEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QRhiWidget_keyReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QRhiWidget) FocusInEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QRhiWidget_focusInEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QRhiWidget) FocusOutEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QRhiWidget_focusOutEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QRhiWidget) EnterEventDefault(event *QEnterEvent) {
	purego.SyscallN(addr_QRhiWidget_enterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QRhiWidget) LeaveEventDefault(event *QEvent) {
	purego.SyscallN(addr_QRhiWidget_leaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QRhiWidget) MoveEventDefault(event *QMoveEvent) {
	purego.SyscallN(addr_QRhiWidget_moveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QRhiWidget) CloseEventDefault(event *QCloseEvent) {
	purego.SyscallN(addr_QRhiWidget_closeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QRhiWidget) ContextMenuEventDefault(event *QContextMenuEvent) {
	purego.SyscallN(addr_QRhiWidget_contextMenuEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QRhiWidget) TabletEventDefault(event *QTabletEvent) {
	purego.SyscallN(addr_QRhiWidget_tabletEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QRhiWidget) ActionEventDefault(event *QActionEvent) {
	purego.SyscallN(addr_QRhiWidget_actionEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QRhiWidget) DragEnterEventDefault(event *QDragEnterEvent) {
	purego.SyscallN(addr_QRhiWidget_dragEnterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QRhiWidget) DragMoveEventDefault(event *QDragMoveEvent) {
	purego.SyscallN(addr_QRhiWidget_dragMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QRhiWidget) DragLeaveEventDefault(event *QDragLeaveEvent) {
	purego.SyscallN(addr_QRhiWidget_dragLeaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QRhiWidget) DropEventDefault(event *QDropEvent) {
	purego.SyscallN(addr_QRhiWidget_dropEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QRhiWidget) ShowEventDefault(event *QShowEvent) {
	purego.SyscallN(addr_QRhiWidget_showEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QRhiWidget) HideEventDefault(event *QHideEvent) {
	purego.SyscallN(addr_QRhiWidget_hideEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QRhiWidget) ChangeEventDefault(arg1 *QEvent) {
	purego.SyscallN(addr_QRhiWidget_changeEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QRhiWidget) MetricDefault(arg1 QPaintDevice__PaintDeviceMetric) int {
	__r0, _, _ := purego.SyscallN(addr_QRhiWidget_metric_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QRhiWidget) InitPainterDefault(painter *QPainter) {
	purego.SyscallN(addr_QRhiWidget_initPainter_default, this.h, uintptr(painter.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
}

func (this *QRhiWidget) RedirectedDefault(offset *QPoint) *QPaintDevice {
	__r, _, _ := purego.SyscallN(addr_QRhiWidget_redirected_default, this.h, uintptr(offset.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(offset)
	return bindQPaintDevice(__r)
}

func (this *QRhiWidget) SharedPainterDefault() *QPainter {
	__r, _, _ := purego.SyscallN(addr_QRhiWidget_sharedPainter_default, this.h)
	runtime.KeepAlive(this)
	return bindQPainter(__r)
}

func (this *QRhiWidget) InputMethodEventDefault(arg1 *QInputMethodEvent) {
	purego.SyscallN(addr_QRhiWidget_inputMethodEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QRhiWidget) InputMethodQueryDefault(arg1 InputMethodQuery) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QRhiWidget_inputMethodQuery_default, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QRhiWidget) FocusNextPrevChildDefault(next bool) bool {
	__r0, _, _ := purego.SyscallN(addr_QRhiWidget_focusNextPrevChild_default, this.h, uintptr(b2i(next)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QRhiWidget) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QRhiWidget_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QRhiWidget) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QRhiWidget_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QRhiWidget) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QRhiWidget_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QRhiWidget) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QRhiWidget_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QRhiWidget) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QRhiWidget_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QRhiWidget) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QRhiWidget_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QRhiWidget) OnFrameSubmitted(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QRhiWidget_connect_frameSubmitted(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QRhiWidget) OnRenderFailed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QRhiWidget_connect_renderFailed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QRhiWidget) OnSampleCountChanged(fn func(int)) int64 {
	id := registerCallback(func(samples int32) { fn(int(samples)) })
	pg_QRhiWidget_connect_sampleCountChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QRhiWidget) OnColorBufferFormatChanged(fn func(QRhiWidget__TextureFormat)) int64 {
	id := registerCallback(func(format int32) { fn(QRhiWidget__TextureFormat(format)) })
	pg_QRhiWidget_connect_colorBufferFormatChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QRhiWidget) OnFixedColorBufferSizeChanged(fn func(*QSize)) int64 {
	id := registerCallback(func(pixelSize uintptr) { fn(UnsafeNewQSize(unsafe.Pointer(pixelSize))) })
	pg_QRhiWidget_connect_fixedColorBufferSizeChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QRhiWidget) OnMirrorVerticallyChanged(fn func(bool)) int64 {
	id := registerCallback(func(enabled int32) { fn(enabled != 0) })
	pg_QRhiWidget_connect_mirrorVerticallyChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QRhiWidget) OnWindowTitleChanged(fn func(string)) int64 {
	id := registerCallback(func(title uintptr, title_len int64) { fn(pgStrNoFree(title, title_len)) })
	pg_QRhiWidget_connect_windowTitleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QRhiWidget) OnWindowIconChanged(fn func(*QIcon)) int64 {
	id := registerCallback(func(icon uintptr) { fn(UnsafeNewQIcon(unsafe.Pointer(icon))) })
	pg_QRhiWidget_connect_windowIconChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QRhiWidget) OnWindowIconTextChanged(fn func(string)) int64 {
	id := registerCallback(func(iconText uintptr, iconText_len int64) { fn(pgStrNoFree(iconText, iconText_len)) })
	pg_QRhiWidget_connect_windowIconTextChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QRhiWidget) OnCustomContextMenuRequested(fn func(*QPoint)) int64 {
	id := registerCallback(func(pos uintptr) { fn(UnsafeNewQPoint(unsafe.Pointer(pos))) })
	pg_QRhiWidget_connect_customContextMenuRequested(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QRhiWidget) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QRhiWidget_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QRhiWidget) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QRhiWidget_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QRhiWidget) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QRhiWidget_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QRhiWidget) Disconnect(handle int64) {
	pg_QRhiWidget__disconnect(handle)
	unregisterCallback(handle)
}

func QRhiWidget_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case -434767277154741265:
		return 1
	case -4302917285655014661:
		return 2
	case 5290442163167458465:
		return 3
	case 1237236456312824991:
		return 4
	case -8659679893868524044:
		return 5
	case 2110182719375245575:
		return 6
	case 729690261457684101:
		return 7
	case 7049264599837532491:
		return 8
	case 4311534060495737339:
		return 9
	case -8266762397995423915:
		return 10
	case 6398007404756346381:
		return 11
	case -332090898078543319:
		return 12
	case -4415920049297795213:
		return 13
	case -4029073905697023887:
		return 14
	case -4255494004099806581:
		return 15
	case 1919988340934511316:
		return 16
	case 1023791977679793185:
		return 17
	case 772598475950050299:
		return 18
	case 975043767033552552:
		return 19
	case -385596010576048993:
		return 20
	case 6316270346668237895:
		return 21
	case -551070009623711640:
		return 22
	case -9161949150489525396:
		return 23
	case -1989937050019464783:
		return 24
	case -6698043222227006015:
		return 25
	case -212068866860701141:
		return 26
	case 1697444140088555581:
		return 27
	case 206122715919917589:
		return 28
	case -1781545137352745978:
		return 29
	case -3139904158373725006:
		return 30
	case 9107058682754591716:
		return 31
	case -6825369174751916:
		return 32
	case -199718707876719129:
		return 33
	case -2039166586246443706:
		return 34
	case 4959376063826165943:
		return 35
	case 427629649267311777:
		return 36
	case 6477500296980050210:
		return 37
	case 4127240633310655048:
		return 38
	case 5792771328741774991:
		return 39
	case 5438819661942843222:
		return 40
	case 7345853097340709638:
		return 41
	case 5898328517505784191:
		return 42
	case 5217623514630858297:
		return 43
	case -5409927581812374498:
		return 44
	case -1185315367526857833:
		return 45
	case 3769868229123868474:
		return 46
	case 7379029982170503008:
		return 47
	case -568369772982215962:
		return 48
	}
	return -1
}

func QRhiWidget_enable(h uintptr, idx int32) { pg_QRhiWidget__set_override(h, idx) }

func (this *QRhiWidget) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QRhiWidget) OnReleaseResources(fn func()) {
	registerVirt(this.UnsafeGoID(), -434767277154741265, func() { fn() }) // slot "releaseResources"
}

func (this *QRhiWidget) OnResizeEvent(fn func(*QResizeEvent)) {
	registerVirt(this.UnsafeGoID(), -4302917285655014661, func(e uintptr) { fn(UnsafeNewQResizeEvent(unsafe.Pointer(e))) }) // slot "resizeEvent"
}

func (this *QRhiWidget) OnPaintEvent(fn func(*QPaintEvent)) {
	registerVirt(this.UnsafeGoID(), 5290442163167458465, func(e uintptr) { fn(UnsafeNewQPaintEvent(unsafe.Pointer(e))) }) // slot "paintEvent"
}

func (this *QRhiWidget) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(e uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(e)))) }) // slot "event"
}

func (this *QRhiWidget) OnDevType(fn func() int) {
	registerVirt(this.UnsafeGoID(), -8659679893868524044, func() int32 { return int32(fn()) }) // slot "devType"
}

func (this *QRhiWidget) OnSetVisible(fn func(bool)) {
	registerVirt(this.UnsafeGoID(), 2110182719375245575, func(visible int32) { fn(visible != 0) }) // slot "setVisible"
}

func (this *QRhiWidget) OnSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 729690261457684101, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sizeHint"
}

func (this *QRhiWidget) OnMinimumSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 7049264599837532491, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "minimumSizeHint"
}

func (this *QRhiWidget) OnHeightForWidth(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), 4311534060495737339, func(arg1 int32) int32 { return int32(fn(int(arg1))) }) // slot "heightForWidth"
}

func (this *QRhiWidget) OnHasHeightForWidth(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -8266762397995423915, func() int32 { return b2i(fn()) }) // slot "hasHeightForWidth"
}

func (this *QRhiWidget) OnPaintEngine(fn func() *QPaintEngine) {
	registerVirt(this.UnsafeGoID(), 6398007404756346381, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "paintEngine"
}

func (this *QRhiWidget) OnMousePressEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -332090898078543319, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mousePressEvent"
}

func (this *QRhiWidget) OnMouseReleaseEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4415920049297795213, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseReleaseEvent"
}

func (this *QRhiWidget) OnMouseDoubleClickEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4029073905697023887, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseDoubleClickEvent"
}

func (this *QRhiWidget) OnMouseMoveEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4255494004099806581, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseMoveEvent"
}

func (this *QRhiWidget) OnWheelEvent(fn func(*QWheelEvent)) {
	registerVirt(this.UnsafeGoID(), 1919988340934511316, func(event uintptr) { fn(UnsafeNewQWheelEvent(unsafe.Pointer(event))) }) // slot "wheelEvent"
}

func (this *QRhiWidget) OnKeyPressEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 1023791977679793185, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyPressEvent"
}

func (this *QRhiWidget) OnKeyReleaseEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 772598475950050299, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyReleaseEvent"
}

func (this *QRhiWidget) OnFocusInEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), 975043767033552552, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusInEvent"
}

func (this *QRhiWidget) OnFocusOutEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), -385596010576048993, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusOutEvent"
}

func (this *QRhiWidget) OnEnterEvent(fn func(*QEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 6316270346668237895, func(event uintptr) { fn(UnsafeNewQEnterEvent(unsafe.Pointer(event))) }) // slot "enterEvent"
}

func (this *QRhiWidget) OnLeaveEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), -551070009623711640, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "leaveEvent"
}

func (this *QRhiWidget) OnMoveEvent(fn func(*QMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -9161949150489525396, func(event uintptr) { fn(UnsafeNewQMoveEvent(unsafe.Pointer(event))) }) // slot "moveEvent"
}

func (this *QRhiWidget) OnCloseEvent(fn func(*QCloseEvent)) {
	registerVirt(this.UnsafeGoID(), -1989937050019464783, func(event uintptr) { fn(UnsafeNewQCloseEvent(unsafe.Pointer(event))) }) // slot "closeEvent"
}

func (this *QRhiWidget) OnContextMenuEvent(fn func(*QContextMenuEvent)) {
	registerVirt(this.UnsafeGoID(), -6698043222227006015, func(event uintptr) { fn(UnsafeNewQContextMenuEvent(unsafe.Pointer(event))) }) // slot "contextMenuEvent"
}

func (this *QRhiWidget) OnTabletEvent(fn func(*QTabletEvent)) {
	registerVirt(this.UnsafeGoID(), -212068866860701141, func(event uintptr) { fn(UnsafeNewQTabletEvent(unsafe.Pointer(event))) }) // slot "tabletEvent"
}

func (this *QRhiWidget) OnActionEvent(fn func(*QActionEvent)) {
	registerVirt(this.UnsafeGoID(), 1697444140088555581, func(event uintptr) { fn(UnsafeNewQActionEvent(unsafe.Pointer(event))) }) // slot "actionEvent"
}

func (this *QRhiWidget) OnDragEnterEvent(fn func(*QDragEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 206122715919917589, func(event uintptr) { fn(UnsafeNewQDragEnterEvent(unsafe.Pointer(event))) }) // slot "dragEnterEvent"
}

func (this *QRhiWidget) OnDragMoveEvent(fn func(*QDragMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -1781545137352745978, func(event uintptr) { fn(UnsafeNewQDragMoveEvent(unsafe.Pointer(event))) }) // slot "dragMoveEvent"
}

func (this *QRhiWidget) OnDragLeaveEvent(fn func(*QDragLeaveEvent)) {
	registerVirt(this.UnsafeGoID(), -3139904158373725006, func(event uintptr) { fn(UnsafeNewQDragLeaveEvent(unsafe.Pointer(event))) }) // slot "dragLeaveEvent"
}

func (this *QRhiWidget) OnDropEvent(fn func(*QDropEvent)) {
	registerVirt(this.UnsafeGoID(), 9107058682754591716, func(event uintptr) { fn(UnsafeNewQDropEvent(unsafe.Pointer(event))) }) // slot "dropEvent"
}

func (this *QRhiWidget) OnShowEvent(fn func(*QShowEvent)) {
	registerVirt(this.UnsafeGoID(), -6825369174751916, func(event uintptr) { fn(UnsafeNewQShowEvent(unsafe.Pointer(event))) }) // slot "showEvent"
}

func (this *QRhiWidget) OnHideEvent(fn func(*QHideEvent)) {
	registerVirt(this.UnsafeGoID(), -199718707876719129, func(event uintptr) { fn(UnsafeNewQHideEvent(unsafe.Pointer(event))) }) // slot "hideEvent"
}

func (this *QRhiWidget) OnNativeEvent(fn func(*QByteArray, unsafe.Pointer, *int64) bool) {
	registerVirt(this.UnsafeGoID(), -2039166586246443706, func(eventType uintptr, message uintptr, result uintptr) int32 { return b2i(fn(UnsafeNewQByteArray(unsafe.Pointer(eventType)), unsafe.Pointer(message), (*int64)(unsafe.Pointer(result)))) }) // slot "nativeEvent"
}

func (this *QRhiWidget) OnChangeEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 4959376063826165943, func(arg1 uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(arg1))) }) // slot "changeEvent"
}

func (this *QRhiWidget) OnMetric(fn func(QPaintDevice__PaintDeviceMetric) int) {
	registerVirt(this.UnsafeGoID(), 427629649267311777, func(arg1 int32) int32 { return int32(fn(QPaintDevice__PaintDeviceMetric(arg1))) }) // slot "metric"
}

func (this *QRhiWidget) OnInitPainter(fn func(*QPainter)) {
	registerVirt(this.UnsafeGoID(), 6477500296980050210, func(painter uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter))) }) // slot "initPainter"
}

func (this *QRhiWidget) OnRedirected(fn func(*QPoint) *QPaintDevice) {
	registerVirt(this.UnsafeGoID(), 4127240633310655048, func(offset uintptr) uintptr { return uintptr(fn(UnsafeNewQPoint(unsafe.Pointer(offset))).UnsafePointer()) }) // slot "redirected"
}

func (this *QRhiWidget) OnSharedPainter(fn func() *QPainter) {
	registerVirt(this.UnsafeGoID(), 5792771328741774991, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sharedPainter"
}

func (this *QRhiWidget) OnInputMethodEvent(fn func(*QInputMethodEvent)) {
	registerVirt(this.UnsafeGoID(), 5438819661942843222, func(arg1 uintptr) { fn(UnsafeNewQInputMethodEvent(unsafe.Pointer(arg1))) }) // slot "inputMethodEvent"
}

func (this *QRhiWidget) OnInputMethodQuery(fn func(InputMethodQuery) *QVariant) {
	registerVirt(this.UnsafeGoID(), 7345853097340709638, func(arg1 int32) uintptr { return uintptr(fn(InputMethodQuery(arg1)).UnsafePointer()) }) // slot "inputMethodQuery"
}

func (this *QRhiWidget) OnFocusNextPrevChild(fn func(bool) bool) {
	registerVirt(this.UnsafeGoID(), 5898328517505784191, func(next int32) int32 { return b2i(fn(next != 0)) }) // slot "focusNextPrevChild"
}

func (this *QRhiWidget) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QRhiWidget) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QRhiWidget) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QRhiWidget) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QRhiWidget) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QRhiWidget) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QRhiWidget) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QRhiWidget_delete(this.h)
	}
	this.UnsafeRelease()
}
