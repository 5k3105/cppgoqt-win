package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QAbstractScrollArea_chain uintptr
	addr_QAbstractScrollArea_new uintptr
	addr_QAbstractScrollArea_new2 uintptr
	addr_QAbstractScrollArea_tr uintptr
	addr_QAbstractScrollArea_verticalScrollBarPolicy uintptr
	addr_QAbstractScrollArea_setVerticalScrollBarPolicy uintptr
	addr_QAbstractScrollArea_verticalScrollBar uintptr
	addr_QAbstractScrollArea_setVerticalScrollBar uintptr
	addr_QAbstractScrollArea_horizontalScrollBarPolicy uintptr
	addr_QAbstractScrollArea_setHorizontalScrollBarPolicy uintptr
	addr_QAbstractScrollArea_horizontalScrollBar uintptr
	addr_QAbstractScrollArea_setHorizontalScrollBar uintptr
	addr_QAbstractScrollArea_cornerWidget uintptr
	addr_QAbstractScrollArea_setCornerWidget uintptr
	addr_QAbstractScrollArea_addScrollBarWidget uintptr
	addr_QAbstractScrollArea_scrollBarWidgets uintptr
	addr_QAbstractScrollArea_viewport uintptr
	addr_QAbstractScrollArea_setViewport uintptr
	addr_QAbstractScrollArea_maximumViewportSize uintptr
	addr_QAbstractScrollArea_setupViewport uintptr
	addr_QAbstractScrollArea_sizeAdjustPolicy uintptr
	addr_QAbstractScrollArea_setSizeAdjustPolicy uintptr
	addr_QAbstractScrollArea_setViewportMargins uintptr
	addr_QAbstractScrollArea_setViewportMargins2 uintptr
	addr_QAbstractScrollArea_viewportMargins uintptr
	addr_QAbstractScrollArea_drawFrame uintptr
	addr_QAbstractScrollArea_setTabOrder uintptr
	addr_QAbstractScrollArea_mouseGrabber uintptr
	addr_QAbstractScrollArea_keyboardGrabber uintptr
	addr_QAbstractScrollArea_find uintptr
	addr_QAbstractScrollArea_createWindowContainer uintptr
	addr_QAbstractScrollArea_updateMicroFocus uintptr
	addr_QAbstractScrollArea_create uintptr
	addr_QAbstractScrollArea_destroy uintptr
	addr_QAbstractScrollArea_focusNextChild uintptr
	addr_QAbstractScrollArea_focusPreviousChild uintptr
	addr_QAbstractScrollArea_connect uintptr
	addr_QAbstractScrollArea_connect2 uintptr
	addr_QAbstractScrollArea_sender uintptr
	addr_QAbstractScrollArea_senderSignalIndex uintptr
	addr_QAbstractScrollArea_receivers uintptr
	addr_QAbstractScrollArea_isSignalConnected uintptr
	pg_QAbstractScrollArea_devicePixelRatioFScale func() float64
	pg_QAbstractScrollArea_encodeMetricF func(int32, float64) int32
	pg_QAbstractScrollArea_getDecodedMetricF func(uintptr, int32, int32) float64
	addr_QAbstractScrollArea_tr2 uintptr
	addr_QAbstractScrollArea_tr3 uintptr
	addr_QAbstractScrollArea_createWindowContainer2 uintptr
	addr_QAbstractScrollArea_createWindowContainer3 uintptr
	addr_QAbstractScrollArea_updateMicroFocus2 uintptr
	addr_QAbstractScrollArea_create2 uintptr
	addr_QAbstractScrollArea_create3 uintptr
	addr_QAbstractScrollArea_create4 uintptr
	addr_QAbstractScrollArea_destroy2 uintptr
	addr_QAbstractScrollArea_destroy3 uintptr
	addr_QAbstractScrollArea_connect3 uintptr
	addr_QAbstractScrollArea_connect4 uintptr
	addr_QAbstractScrollArea_setupViewport_default uintptr
	addr_QAbstractScrollArea_metaObject_default uintptr
	addr_QAbstractScrollArea_minimumSizeHint_default uintptr
	addr_QAbstractScrollArea_sizeHint_default uintptr
	addr_QAbstractScrollArea_eventFilter_default uintptr
	addr_QAbstractScrollArea_event_default uintptr
	addr_QAbstractScrollArea_viewportEvent_default uintptr
	addr_QAbstractScrollArea_resizeEvent_default uintptr
	addr_QAbstractScrollArea_paintEvent_default uintptr
	addr_QAbstractScrollArea_mousePressEvent_default uintptr
	addr_QAbstractScrollArea_mouseReleaseEvent_default uintptr
	addr_QAbstractScrollArea_mouseDoubleClickEvent_default uintptr
	addr_QAbstractScrollArea_mouseMoveEvent_default uintptr
	addr_QAbstractScrollArea_wheelEvent_default uintptr
	addr_QAbstractScrollArea_contextMenuEvent_default uintptr
	addr_QAbstractScrollArea_dragEnterEvent_default uintptr
	addr_QAbstractScrollArea_dragMoveEvent_default uintptr
	addr_QAbstractScrollArea_dragLeaveEvent_default uintptr
	addr_QAbstractScrollArea_dropEvent_default uintptr
	addr_QAbstractScrollArea_keyPressEvent_default uintptr
	addr_QAbstractScrollArea_scrollContentsBy_default uintptr
	addr_QAbstractScrollArea_viewportSizeHint_default uintptr
	addr_QAbstractScrollArea_changeEvent_default uintptr
	addr_QAbstractScrollArea_initStyleOption_default uintptr
	addr_QAbstractScrollArea_devType_default uintptr
	addr_QAbstractScrollArea_setVisible_default uintptr
	addr_QAbstractScrollArea_heightForWidth_default uintptr
	addr_QAbstractScrollArea_hasHeightForWidth_default uintptr
	addr_QAbstractScrollArea_paintEngine_default uintptr
	addr_QAbstractScrollArea_keyReleaseEvent_default uintptr
	addr_QAbstractScrollArea_focusInEvent_default uintptr
	addr_QAbstractScrollArea_focusOutEvent_default uintptr
	addr_QAbstractScrollArea_enterEvent_default uintptr
	addr_QAbstractScrollArea_leaveEvent_default uintptr
	addr_QAbstractScrollArea_moveEvent_default uintptr
	addr_QAbstractScrollArea_closeEvent_default uintptr
	addr_QAbstractScrollArea_tabletEvent_default uintptr
	addr_QAbstractScrollArea_actionEvent_default uintptr
	addr_QAbstractScrollArea_showEvent_default uintptr
	addr_QAbstractScrollArea_hideEvent_default uintptr
	addr_QAbstractScrollArea_metric_default uintptr
	addr_QAbstractScrollArea_initPainter_default uintptr
	addr_QAbstractScrollArea_redirected_default uintptr
	addr_QAbstractScrollArea_sharedPainter_default uintptr
	addr_QAbstractScrollArea_inputMethodEvent_default uintptr
	addr_QAbstractScrollArea_inputMethodQuery_default uintptr
	addr_QAbstractScrollArea_focusNextPrevChild_default uintptr
	addr_QAbstractScrollArea_timerEvent_default uintptr
	addr_QAbstractScrollArea_childEvent_default uintptr
	addr_QAbstractScrollArea_customEvent_default uintptr
	addr_QAbstractScrollArea_connectNotify_default uintptr
	addr_QAbstractScrollArea_disconnectNotify_default uintptr
	pg_QAbstractScrollArea_connect_windowTitleChanged func(uintptr, int64)
	pg_QAbstractScrollArea_connect_windowIconChanged func(uintptr, int64)
	pg_QAbstractScrollArea_connect_windowIconTextChanged func(uintptr, int64)
	pg_QAbstractScrollArea_connect_customContextMenuRequested func(uintptr, int64)
	pg_QAbstractScrollArea_connect_destroyed func(uintptr, int64)
	pg_QAbstractScrollArea_connect_destroyed2 func(uintptr, int64)
	pg_QAbstractScrollArea_connect_objectNameChanged func(uintptr, int64)
	pg_QAbstractScrollArea__disconnect func(int64)
	pg_QAbstractScrollArea__set_override func(uintptr, int32)
	pg_QAbstractScrollArea_delete func(uintptr)
)

func init() {
	symaddr(&addr_QAbstractScrollArea_chain, "QAbstractScrollArea_chain")
	symaddr(&addr_QAbstractScrollArea_new, "QAbstractScrollArea_new")
	symaddr(&addr_QAbstractScrollArea_new2, "QAbstractScrollArea_new2")
	symaddr(&addr_QAbstractScrollArea_tr, "QAbstractScrollArea_tr")
	symaddr(&addr_QAbstractScrollArea_verticalScrollBarPolicy, "QAbstractScrollArea_verticalScrollBarPolicy")
	symaddr(&addr_QAbstractScrollArea_setVerticalScrollBarPolicy, "QAbstractScrollArea_setVerticalScrollBarPolicy")
	symaddr(&addr_QAbstractScrollArea_verticalScrollBar, "QAbstractScrollArea_verticalScrollBar")
	symaddr(&addr_QAbstractScrollArea_setVerticalScrollBar, "QAbstractScrollArea_setVerticalScrollBar")
	symaddr(&addr_QAbstractScrollArea_horizontalScrollBarPolicy, "QAbstractScrollArea_horizontalScrollBarPolicy")
	symaddr(&addr_QAbstractScrollArea_setHorizontalScrollBarPolicy, "QAbstractScrollArea_setHorizontalScrollBarPolicy")
	symaddr(&addr_QAbstractScrollArea_horizontalScrollBar, "QAbstractScrollArea_horizontalScrollBar")
	symaddr(&addr_QAbstractScrollArea_setHorizontalScrollBar, "QAbstractScrollArea_setHorizontalScrollBar")
	symaddr(&addr_QAbstractScrollArea_cornerWidget, "QAbstractScrollArea_cornerWidget")
	symaddr(&addr_QAbstractScrollArea_setCornerWidget, "QAbstractScrollArea_setCornerWidget")
	symaddr(&addr_QAbstractScrollArea_addScrollBarWidget, "QAbstractScrollArea_addScrollBarWidget")
	symaddr(&addr_QAbstractScrollArea_scrollBarWidgets, "QAbstractScrollArea_scrollBarWidgets")
	symaddr(&addr_QAbstractScrollArea_viewport, "QAbstractScrollArea_viewport")
	symaddr(&addr_QAbstractScrollArea_setViewport, "QAbstractScrollArea_setViewport")
	symaddr(&addr_QAbstractScrollArea_maximumViewportSize, "QAbstractScrollArea_maximumViewportSize")
	symaddr(&addr_QAbstractScrollArea_setupViewport, "QAbstractScrollArea_setupViewport")
	symaddr(&addr_QAbstractScrollArea_sizeAdjustPolicy, "QAbstractScrollArea_sizeAdjustPolicy")
	symaddr(&addr_QAbstractScrollArea_setSizeAdjustPolicy, "QAbstractScrollArea_setSizeAdjustPolicy")
	symaddr(&addr_QAbstractScrollArea_setViewportMargins, "QAbstractScrollArea_setViewportMargins")
	symaddr(&addr_QAbstractScrollArea_setViewportMargins2, "QAbstractScrollArea_setViewportMargins2")
	symaddr(&addr_QAbstractScrollArea_viewportMargins, "QAbstractScrollArea_viewportMargins")
	symaddr(&addr_QAbstractScrollArea_drawFrame, "QAbstractScrollArea_drawFrame")
	symaddr(&addr_QAbstractScrollArea_setTabOrder, "QAbstractScrollArea_setTabOrder")
	symaddr(&addr_QAbstractScrollArea_mouseGrabber, "QAbstractScrollArea_mouseGrabber")
	symaddr(&addr_QAbstractScrollArea_keyboardGrabber, "QAbstractScrollArea_keyboardGrabber")
	symaddr(&addr_QAbstractScrollArea_find, "QAbstractScrollArea_find")
	symaddr(&addr_QAbstractScrollArea_createWindowContainer, "QAbstractScrollArea_createWindowContainer")
	symaddr(&addr_QAbstractScrollArea_updateMicroFocus, "QAbstractScrollArea_updateMicroFocus")
	symaddr(&addr_QAbstractScrollArea_create, "QAbstractScrollArea_create")
	symaddr(&addr_QAbstractScrollArea_destroy, "QAbstractScrollArea_destroy")
	symaddr(&addr_QAbstractScrollArea_focusNextChild, "QAbstractScrollArea_focusNextChild")
	symaddr(&addr_QAbstractScrollArea_focusPreviousChild, "QAbstractScrollArea_focusPreviousChild")
	symaddr(&addr_QAbstractScrollArea_connect, "QAbstractScrollArea_connect")
	symaddr(&addr_QAbstractScrollArea_connect2, "QAbstractScrollArea_connect2")
	symaddr(&addr_QAbstractScrollArea_sender, "QAbstractScrollArea_sender")
	symaddr(&addr_QAbstractScrollArea_senderSignalIndex, "QAbstractScrollArea_senderSignalIndex")
	symaddr(&addr_QAbstractScrollArea_receivers, "QAbstractScrollArea_receivers")
	symaddr(&addr_QAbstractScrollArea_isSignalConnected, "QAbstractScrollArea_isSignalConnected")
	reg(&pg_QAbstractScrollArea_devicePixelRatioFScale, "QAbstractScrollArea_devicePixelRatioFScale")
	reg(&pg_QAbstractScrollArea_encodeMetricF, "QAbstractScrollArea_encodeMetricF")
	reg(&pg_QAbstractScrollArea_getDecodedMetricF, "QAbstractScrollArea_getDecodedMetricF")
	symaddr(&addr_QAbstractScrollArea_tr2, "QAbstractScrollArea_tr2")
	symaddr(&addr_QAbstractScrollArea_tr3, "QAbstractScrollArea_tr3")
	symaddr(&addr_QAbstractScrollArea_createWindowContainer2, "QAbstractScrollArea_createWindowContainer2")
	symaddr(&addr_QAbstractScrollArea_createWindowContainer3, "QAbstractScrollArea_createWindowContainer3")
	symaddr(&addr_QAbstractScrollArea_updateMicroFocus2, "QAbstractScrollArea_updateMicroFocus2")
	symaddr(&addr_QAbstractScrollArea_create2, "QAbstractScrollArea_create2")
	symaddr(&addr_QAbstractScrollArea_create3, "QAbstractScrollArea_create3")
	symaddr(&addr_QAbstractScrollArea_create4, "QAbstractScrollArea_create4")
	symaddr(&addr_QAbstractScrollArea_destroy2, "QAbstractScrollArea_destroy2")
	symaddr(&addr_QAbstractScrollArea_destroy3, "QAbstractScrollArea_destroy3")
	symaddr(&addr_QAbstractScrollArea_connect3, "QAbstractScrollArea_connect3")
	symaddr(&addr_QAbstractScrollArea_connect4, "QAbstractScrollArea_connect4")
	symaddr(&addr_QAbstractScrollArea_setupViewport_default, "QAbstractScrollArea_setupViewport_default")
	symaddr(&addr_QAbstractScrollArea_metaObject_default, "QAbstractScrollArea_metaObject_default")
	symaddr(&addr_QAbstractScrollArea_minimumSizeHint_default, "QAbstractScrollArea_minimumSizeHint_default")
	symaddr(&addr_QAbstractScrollArea_sizeHint_default, "QAbstractScrollArea_sizeHint_default")
	symaddr(&addr_QAbstractScrollArea_eventFilter_default, "QAbstractScrollArea_eventFilter_default")
	symaddr(&addr_QAbstractScrollArea_event_default, "QAbstractScrollArea_event_default")
	symaddr(&addr_QAbstractScrollArea_viewportEvent_default, "QAbstractScrollArea_viewportEvent_default")
	symaddr(&addr_QAbstractScrollArea_resizeEvent_default, "QAbstractScrollArea_resizeEvent_default")
	symaddr(&addr_QAbstractScrollArea_paintEvent_default, "QAbstractScrollArea_paintEvent_default")
	symaddr(&addr_QAbstractScrollArea_mousePressEvent_default, "QAbstractScrollArea_mousePressEvent_default")
	symaddr(&addr_QAbstractScrollArea_mouseReleaseEvent_default, "QAbstractScrollArea_mouseReleaseEvent_default")
	symaddr(&addr_QAbstractScrollArea_mouseDoubleClickEvent_default, "QAbstractScrollArea_mouseDoubleClickEvent_default")
	symaddr(&addr_QAbstractScrollArea_mouseMoveEvent_default, "QAbstractScrollArea_mouseMoveEvent_default")
	symaddr(&addr_QAbstractScrollArea_wheelEvent_default, "QAbstractScrollArea_wheelEvent_default")
	symaddr(&addr_QAbstractScrollArea_contextMenuEvent_default, "QAbstractScrollArea_contextMenuEvent_default")
	symaddr(&addr_QAbstractScrollArea_dragEnterEvent_default, "QAbstractScrollArea_dragEnterEvent_default")
	symaddr(&addr_QAbstractScrollArea_dragMoveEvent_default, "QAbstractScrollArea_dragMoveEvent_default")
	symaddr(&addr_QAbstractScrollArea_dragLeaveEvent_default, "QAbstractScrollArea_dragLeaveEvent_default")
	symaddr(&addr_QAbstractScrollArea_dropEvent_default, "QAbstractScrollArea_dropEvent_default")
	symaddr(&addr_QAbstractScrollArea_keyPressEvent_default, "QAbstractScrollArea_keyPressEvent_default")
	symaddr(&addr_QAbstractScrollArea_scrollContentsBy_default, "QAbstractScrollArea_scrollContentsBy_default")
	symaddr(&addr_QAbstractScrollArea_viewportSizeHint_default, "QAbstractScrollArea_viewportSizeHint_default")
	symaddr(&addr_QAbstractScrollArea_changeEvent_default, "QAbstractScrollArea_changeEvent_default")
	symaddr(&addr_QAbstractScrollArea_initStyleOption_default, "QAbstractScrollArea_initStyleOption_default")
	symaddr(&addr_QAbstractScrollArea_devType_default, "QAbstractScrollArea_devType_default")
	symaddr(&addr_QAbstractScrollArea_setVisible_default, "QAbstractScrollArea_setVisible_default")
	symaddr(&addr_QAbstractScrollArea_heightForWidth_default, "QAbstractScrollArea_heightForWidth_default")
	symaddr(&addr_QAbstractScrollArea_hasHeightForWidth_default, "QAbstractScrollArea_hasHeightForWidth_default")
	symaddr(&addr_QAbstractScrollArea_paintEngine_default, "QAbstractScrollArea_paintEngine_default")
	symaddr(&addr_QAbstractScrollArea_keyReleaseEvent_default, "QAbstractScrollArea_keyReleaseEvent_default")
	symaddr(&addr_QAbstractScrollArea_focusInEvent_default, "QAbstractScrollArea_focusInEvent_default")
	symaddr(&addr_QAbstractScrollArea_focusOutEvent_default, "QAbstractScrollArea_focusOutEvent_default")
	symaddr(&addr_QAbstractScrollArea_enterEvent_default, "QAbstractScrollArea_enterEvent_default")
	symaddr(&addr_QAbstractScrollArea_leaveEvent_default, "QAbstractScrollArea_leaveEvent_default")
	symaddr(&addr_QAbstractScrollArea_moveEvent_default, "QAbstractScrollArea_moveEvent_default")
	symaddr(&addr_QAbstractScrollArea_closeEvent_default, "QAbstractScrollArea_closeEvent_default")
	symaddr(&addr_QAbstractScrollArea_tabletEvent_default, "QAbstractScrollArea_tabletEvent_default")
	symaddr(&addr_QAbstractScrollArea_actionEvent_default, "QAbstractScrollArea_actionEvent_default")
	symaddr(&addr_QAbstractScrollArea_showEvent_default, "QAbstractScrollArea_showEvent_default")
	symaddr(&addr_QAbstractScrollArea_hideEvent_default, "QAbstractScrollArea_hideEvent_default")
	symaddr(&addr_QAbstractScrollArea_metric_default, "QAbstractScrollArea_metric_default")
	symaddr(&addr_QAbstractScrollArea_initPainter_default, "QAbstractScrollArea_initPainter_default")
	symaddr(&addr_QAbstractScrollArea_redirected_default, "QAbstractScrollArea_redirected_default")
	symaddr(&addr_QAbstractScrollArea_sharedPainter_default, "QAbstractScrollArea_sharedPainter_default")
	symaddr(&addr_QAbstractScrollArea_inputMethodEvent_default, "QAbstractScrollArea_inputMethodEvent_default")
	symaddr(&addr_QAbstractScrollArea_inputMethodQuery_default, "QAbstractScrollArea_inputMethodQuery_default")
	symaddr(&addr_QAbstractScrollArea_focusNextPrevChild_default, "QAbstractScrollArea_focusNextPrevChild_default")
	symaddr(&addr_QAbstractScrollArea_timerEvent_default, "QAbstractScrollArea_timerEvent_default")
	symaddr(&addr_QAbstractScrollArea_childEvent_default, "QAbstractScrollArea_childEvent_default")
	symaddr(&addr_QAbstractScrollArea_customEvent_default, "QAbstractScrollArea_customEvent_default")
	symaddr(&addr_QAbstractScrollArea_connectNotify_default, "QAbstractScrollArea_connectNotify_default")
	symaddr(&addr_QAbstractScrollArea_disconnectNotify_default, "QAbstractScrollArea_disconnectNotify_default")
	reg(&pg_QAbstractScrollArea_connect_windowTitleChanged, "QAbstractScrollArea_connect_windowTitleChanged")
	reg(&pg_QAbstractScrollArea_connect_windowIconChanged, "QAbstractScrollArea_connect_windowIconChanged")
	reg(&pg_QAbstractScrollArea_connect_windowIconTextChanged, "QAbstractScrollArea_connect_windowIconTextChanged")
	reg(&pg_QAbstractScrollArea_connect_customContextMenuRequested, "QAbstractScrollArea_connect_customContextMenuRequested")
	reg(&pg_QAbstractScrollArea_connect_destroyed, "QAbstractScrollArea_connect_destroyed")
	reg(&pg_QAbstractScrollArea_connect_destroyed2, "QAbstractScrollArea_connect_destroyed2")
	reg(&pg_QAbstractScrollArea_connect_objectNameChanged, "QAbstractScrollArea_connect_objectNameChanged")
	reg(&pg_QAbstractScrollArea__disconnect, "QAbstractScrollArea__disconnect")
	reg(&pg_QAbstractScrollArea__set_override, "QAbstractScrollArea__set_override")
	reg(&pg_QAbstractScrollArea_delete, "QAbstractScrollArea_delete")
}

type QAbstractScrollArea struct {
	h uintptr
	*QFrame
}

func bindQAbstractScrollArea(h uintptr) *QAbstractScrollArea {
	if h == 0 {
		return nil
	}
	var __c [3]uintptr
	purego.SyscallN(addr_QAbstractScrollArea_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QAbstractScrollArea{h: h, QFrame: &QFrame{h: __c[0], QWidget: &QWidget{h: __c[1], QObject: &QObject{h: __c[2]}}}}
}

func newownedQAbstractScrollArea(h uintptr) *QAbstractScrollArea { return bindQAbstractScrollArea(h) }

func (this *QAbstractScrollArea) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QAbstractScrollArea) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQAbstractScrollArea(h unsafe.Pointer) *QAbstractScrollArea { return bindQAbstractScrollArea(uintptr(h)) }
func UnsafeNewQAbstractScrollAreaOwned(h unsafe.Pointer) *QAbstractScrollArea { return newownedQAbstractScrollArea(uintptr(h)) }

func (this *QAbstractScrollArea) UnsafeRelease() {
	this.h = 0
	if this.QFrame != nil {
		this.QFrame.UnsafeRelease()
	}
}

func NewQAbstractScrollArea(parent *QWidget) *QAbstractScrollArea {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QAbstractScrollArea_new, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQAbstractScrollArea(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QAbstractScrollArea_virtIdx, QAbstractScrollArea_enable)
	return __o
}

func NewQAbstractScrollArea2() *QAbstractScrollArea {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QAbstractScrollArea_new2, uintptr(__goid))
	__o := newownedQAbstractScrollArea(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QAbstractScrollArea_virtIdx, QAbstractScrollArea_enable)
	return __o
}

func QAbstractScrollArea_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAbstractScrollArea_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QAbstractScrollArea) VerticalScrollBarPolicy() ScrollBarPolicy {
	__r0, _, _ := purego.SyscallN(addr_QAbstractScrollArea_verticalScrollBarPolicy, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return ScrollBarPolicy(__r)
}

func (this *QAbstractScrollArea) SetVerticalScrollBarPolicy(arg1 ScrollBarPolicy) {
	purego.SyscallN(addr_QAbstractScrollArea_setVerticalScrollBarPolicy, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QAbstractScrollArea) VerticalScrollBar() *QScrollBar {
	__r, _, _ := purego.SyscallN(addr_QAbstractScrollArea_verticalScrollBar, this.h)
	runtime.KeepAlive(this)
	return bindQScrollBar(__r)
}

func (this *QAbstractScrollArea) SetVerticalScrollBar(scrollbar *QScrollBar) {
	purego.SyscallN(addr_QAbstractScrollArea_setVerticalScrollBar, this.h, uintptr(scrollbar.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(scrollbar)
}

func (this *QAbstractScrollArea) HorizontalScrollBarPolicy() ScrollBarPolicy {
	__r0, _, _ := purego.SyscallN(addr_QAbstractScrollArea_horizontalScrollBarPolicy, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return ScrollBarPolicy(__r)
}

func (this *QAbstractScrollArea) SetHorizontalScrollBarPolicy(arg1 ScrollBarPolicy) {
	purego.SyscallN(addr_QAbstractScrollArea_setHorizontalScrollBarPolicy, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QAbstractScrollArea) HorizontalScrollBar() *QScrollBar {
	__r, _, _ := purego.SyscallN(addr_QAbstractScrollArea_horizontalScrollBar, this.h)
	runtime.KeepAlive(this)
	return bindQScrollBar(__r)
}

func (this *QAbstractScrollArea) SetHorizontalScrollBar(scrollbar *QScrollBar) {
	purego.SyscallN(addr_QAbstractScrollArea_setHorizontalScrollBar, this.h, uintptr(scrollbar.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(scrollbar)
}

func (this *QAbstractScrollArea) CornerWidget() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QAbstractScrollArea_cornerWidget, this.h)
	runtime.KeepAlive(this)
	return bindQWidget(__r)
}

func (this *QAbstractScrollArea) SetCornerWidget(widget *QWidget) {
	purego.SyscallN(addr_QAbstractScrollArea_setCornerWidget, this.h, uintptr(widget.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(widget)
}

func (this *QAbstractScrollArea) AddScrollBarWidget(widget *QWidget, alignment AlignmentFlag) {
	purego.SyscallN(addr_QAbstractScrollArea_addScrollBarWidget, this.h, uintptr(widget.UnsafePointer()), uintptr(alignment))
	runtime.KeepAlive(this)
	runtime.KeepAlive(widget)
}

func (this *QAbstractScrollArea) ScrollBarWidgets(alignment AlignmentFlag) []*QWidget {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAbstractScrollArea_scrollBarWidgets, this.h, uintptr(alignment), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QWidget, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = bindQWidget(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QAbstractScrollArea) Viewport() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QAbstractScrollArea_viewport, this.h)
	runtime.KeepAlive(this)
	return bindQWidget(__r)
}

func (this *QAbstractScrollArea) SetViewport(widget *QWidget) {
	purego.SyscallN(addr_QAbstractScrollArea_setViewport, this.h, uintptr(widget.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(widget)
}

func (this *QAbstractScrollArea) MaximumViewportSize() *QSize {
	__r, _, _ := purego.SyscallN(addr_QAbstractScrollArea_maximumViewportSize, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QAbstractScrollArea) SetupViewport(viewport *QWidget) {
	purego.SyscallN(addr_QAbstractScrollArea_setupViewport, this.h, uintptr(viewport.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(viewport)
}

func (this *QAbstractScrollArea) SizeAdjustPolicy() QAbstractScrollArea__SizeAdjustPolicy {
	__r0, _, _ := purego.SyscallN(addr_QAbstractScrollArea_sizeAdjustPolicy, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QAbstractScrollArea__SizeAdjustPolicy(__r)
}

func (this *QAbstractScrollArea) SetSizeAdjustPolicy(policy QAbstractScrollArea__SizeAdjustPolicy) {
	purego.SyscallN(addr_QAbstractScrollArea_setSizeAdjustPolicy, this.h, uintptr(policy))
	runtime.KeepAlive(this)
}

func (this *QAbstractScrollArea) SetViewportMargins(left int, top int, right int, bottom int) {
	purego.SyscallN(addr_QAbstractScrollArea_setViewportMargins, this.h, uintptr(left), uintptr(top), uintptr(right), uintptr(bottom))
	runtime.KeepAlive(this)
}

func (this *QAbstractScrollArea) SetViewportMarginsWithMargins(margins *QMargins) {
	purego.SyscallN(addr_QAbstractScrollArea_setViewportMargins2, this.h, uintptr(margins.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(margins)
}

func (this *QAbstractScrollArea) ViewportMargins() *QMargins {
	__r, _, _ := purego.SyscallN(addr_QAbstractScrollArea_viewportMargins, this.h)
	runtime.KeepAlive(this)
	return newownedQMargins(__r)
}

func (this *QAbstractScrollArea) DrawFrame(arg1 *QPainter) {
	purego.SyscallN(addr_QAbstractScrollArea_drawFrame, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func QAbstractScrollArea_SetTabOrder(arg1 *QWidget, arg2 *QWidget) {
	purego.SyscallN(addr_QAbstractScrollArea_setTabOrder, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()))
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
}

func QAbstractScrollArea_MouseGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QAbstractScrollArea_mouseGrabber)
	return bindQWidget(__r)
}

func QAbstractScrollArea_KeyboardGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QAbstractScrollArea_keyboardGrabber)
	return bindQWidget(__r)
}

func QAbstractScrollArea_Find(arg1 uint64) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QAbstractScrollArea_find, uintptr(arg1))
	return bindQWidget(__r)
}

func QAbstractScrollArea_CreateWindowContainer(window *QWindow) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QAbstractScrollArea_createWindowContainer, uintptr(window.UnsafePointer()))
	runtime.KeepAlive(window)
	return bindQWidget(__r)
}

func (this *QAbstractScrollArea) UpdateMicroFocus() {
	purego.SyscallN(addr_QAbstractScrollArea_updateMicroFocus, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractScrollArea) Create() {
	purego.SyscallN(addr_QAbstractScrollArea_create, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractScrollArea) Destroy() {
	purego.SyscallN(addr_QAbstractScrollArea_destroy, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractScrollArea) FocusNextChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractScrollArea_focusNextChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QAbstractScrollArea) FocusPreviousChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractScrollArea_focusPreviousChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QAbstractScrollArea_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QAbstractScrollArea_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QAbstractScrollArea_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QAbstractScrollArea_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QAbstractScrollArea) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QAbstractScrollArea_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QAbstractScrollArea) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QAbstractScrollArea_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAbstractScrollArea) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QAbstractScrollArea_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QAbstractScrollArea) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractScrollArea_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QAbstractScrollArea_DevicePixelRatioFScale() float64 {
	__r := pg_QAbstractScrollArea_devicePixelRatioFScale()
	return __r
}

func QAbstractScrollArea_EncodeMetricF(metric QPaintDevice__PaintDeviceMetric, value float64) int {
	__r := pg_QAbstractScrollArea_encodeMetricF(int32(metric), value)
	return int(__r)
}

func (this *QAbstractScrollArea) GetDecodedMetricF(metricA QPaintDevice__PaintDeviceMetric, metricB QPaintDevice__PaintDeviceMetric) float64 {
	__r := pg_QAbstractScrollArea_getDecodedMetricF(this.h, int32(metricA), int32(metricB))
	runtime.KeepAlive(this)
	return __r
}

func QAbstractScrollArea_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAbstractScrollArea_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QAbstractScrollArea_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAbstractScrollArea_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QAbstractScrollArea_CreateWindowContainer2(window *QWindow, parent *QWidget) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QAbstractScrollArea_createWindowContainer2, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func QAbstractScrollArea_CreateWindowContainer3(window *QWindow, parent *QWidget, flags WindowType) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QAbstractScrollArea_createWindowContainer3, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()), uintptr(flags))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func (this *QAbstractScrollArea) UpdateMicroFocusWithQuery(query InputMethodQuery) {
	purego.SyscallN(addr_QAbstractScrollArea_updateMicroFocus2, this.h, uintptr(query))
	runtime.KeepAlive(this)
}

func (this *QAbstractScrollArea) CreateWith(arg1 uint64) {
	purego.SyscallN(addr_QAbstractScrollArea_create2, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QAbstractScrollArea) Create2(arg1 uint64, initializeWindow bool) {
	purego.SyscallN(addr_QAbstractScrollArea_create3, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)))
	runtime.KeepAlive(this)
}

func (this *QAbstractScrollArea) Create3(arg1 uint64, initializeWindow bool, destroyOldWindow bool) {
	purego.SyscallN(addr_QAbstractScrollArea_create4, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)), uintptr(b2i(destroyOldWindow)))
	runtime.KeepAlive(this)
}

func (this *QAbstractScrollArea) DestroyWithDestroyWindow(destroyWindow bool) {
	purego.SyscallN(addr_QAbstractScrollArea_destroy2, this.h, uintptr(b2i(destroyWindow)))
	runtime.KeepAlive(this)
}

func (this *QAbstractScrollArea) Destroy2(destroyWindow bool, destroySubWindows bool) {
	purego.SyscallN(addr_QAbstractScrollArea_destroy3, this.h, uintptr(b2i(destroyWindow)), uintptr(b2i(destroySubWindows)))
	runtime.KeepAlive(this)
}

func QAbstractScrollArea_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QAbstractScrollArea_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QAbstractScrollArea_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QAbstractScrollArea_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QAbstractScrollArea) SetupViewportDefault(viewport *QWidget) {
	purego.SyscallN(addr_QAbstractScrollArea_setupViewport_default, this.h, uintptr(viewport.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(viewport)
}

func (this *QAbstractScrollArea) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QAbstractScrollArea_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QAbstractScrollArea) MinimumSizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QAbstractScrollArea_minimumSizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QAbstractScrollArea) SizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QAbstractScrollArea_sizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QAbstractScrollArea) EventFilterDefault(arg1 *QObject, arg2 *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractScrollArea_eventFilter_default, this.h, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
	return __r != 0
}

func (this *QAbstractScrollArea) EventDefault(arg1 *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractScrollArea_event_default, this.h, uintptr(arg1.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return __r != 0
}

func (this *QAbstractScrollArea) ViewportEventDefault(arg1 *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractScrollArea_viewportEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return __r != 0
}

func (this *QAbstractScrollArea) ResizeEventDefault(arg1 *QResizeEvent) {
	purego.SyscallN(addr_QAbstractScrollArea_resizeEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QAbstractScrollArea) PaintEventDefault(arg1 *QPaintEvent) {
	purego.SyscallN(addr_QAbstractScrollArea_paintEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QAbstractScrollArea) MousePressEventDefault(arg1 *QMouseEvent) {
	purego.SyscallN(addr_QAbstractScrollArea_mousePressEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QAbstractScrollArea) MouseReleaseEventDefault(arg1 *QMouseEvent) {
	purego.SyscallN(addr_QAbstractScrollArea_mouseReleaseEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QAbstractScrollArea) MouseDoubleClickEventDefault(arg1 *QMouseEvent) {
	purego.SyscallN(addr_QAbstractScrollArea_mouseDoubleClickEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QAbstractScrollArea) MouseMoveEventDefault(arg1 *QMouseEvent) {
	purego.SyscallN(addr_QAbstractScrollArea_mouseMoveEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QAbstractScrollArea) WheelEventDefault(arg1 *QWheelEvent) {
	purego.SyscallN(addr_QAbstractScrollArea_wheelEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QAbstractScrollArea) ContextMenuEventDefault(arg1 *QContextMenuEvent) {
	purego.SyscallN(addr_QAbstractScrollArea_contextMenuEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QAbstractScrollArea) DragEnterEventDefault(arg1 *QDragEnterEvent) {
	purego.SyscallN(addr_QAbstractScrollArea_dragEnterEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QAbstractScrollArea) DragMoveEventDefault(arg1 *QDragMoveEvent) {
	purego.SyscallN(addr_QAbstractScrollArea_dragMoveEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QAbstractScrollArea) DragLeaveEventDefault(arg1 *QDragLeaveEvent) {
	purego.SyscallN(addr_QAbstractScrollArea_dragLeaveEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QAbstractScrollArea) DropEventDefault(arg1 *QDropEvent) {
	purego.SyscallN(addr_QAbstractScrollArea_dropEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QAbstractScrollArea) KeyPressEventDefault(arg1 *QKeyEvent) {
	purego.SyscallN(addr_QAbstractScrollArea_keyPressEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QAbstractScrollArea) ScrollContentsByDefault(dx int, dy int) {
	purego.SyscallN(addr_QAbstractScrollArea_scrollContentsBy_default, this.h, uintptr(dx), uintptr(dy))
	runtime.KeepAlive(this)
}

func (this *QAbstractScrollArea) ViewportSizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QAbstractScrollArea_viewportSizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QAbstractScrollArea) ChangeEventDefault(arg1 *QEvent) {
	purego.SyscallN(addr_QAbstractScrollArea_changeEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QAbstractScrollArea) InitStyleOptionDefault(option *QStyleOptionFrame) {
	purego.SyscallN(addr_QAbstractScrollArea_initStyleOption_default, this.h, uintptr(option.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(option)
}

func (this *QAbstractScrollArea) DevTypeDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QAbstractScrollArea_devType_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAbstractScrollArea) SetVisibleDefault(visible bool) {
	purego.SyscallN(addr_QAbstractScrollArea_setVisible_default, this.h, uintptr(b2i(visible)))
	runtime.KeepAlive(this)
}

func (this *QAbstractScrollArea) HeightForWidthDefault(arg1 int) int {
	__r0, _, _ := purego.SyscallN(addr_QAbstractScrollArea_heightForWidth_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAbstractScrollArea) HasHeightForWidthDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractScrollArea_hasHeightForWidth_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QAbstractScrollArea) PaintEngineDefault() *QPaintEngine {
	__r, _, _ := purego.SyscallN(addr_QAbstractScrollArea_paintEngine_default, this.h)
	runtime.KeepAlive(this)
	return bindQPaintEngine(__r)
}

func (this *QAbstractScrollArea) KeyReleaseEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QAbstractScrollArea_keyReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractScrollArea) FocusInEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QAbstractScrollArea_focusInEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractScrollArea) FocusOutEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QAbstractScrollArea_focusOutEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractScrollArea) EnterEventDefault(event *QEnterEvent) {
	purego.SyscallN(addr_QAbstractScrollArea_enterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractScrollArea) LeaveEventDefault(event *QEvent) {
	purego.SyscallN(addr_QAbstractScrollArea_leaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractScrollArea) MoveEventDefault(event *QMoveEvent) {
	purego.SyscallN(addr_QAbstractScrollArea_moveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractScrollArea) CloseEventDefault(event *QCloseEvent) {
	purego.SyscallN(addr_QAbstractScrollArea_closeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractScrollArea) TabletEventDefault(event *QTabletEvent) {
	purego.SyscallN(addr_QAbstractScrollArea_tabletEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractScrollArea) ActionEventDefault(event *QActionEvent) {
	purego.SyscallN(addr_QAbstractScrollArea_actionEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractScrollArea) ShowEventDefault(event *QShowEvent) {
	purego.SyscallN(addr_QAbstractScrollArea_showEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractScrollArea) HideEventDefault(event *QHideEvent) {
	purego.SyscallN(addr_QAbstractScrollArea_hideEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractScrollArea) MetricDefault(arg1 QPaintDevice__PaintDeviceMetric) int {
	__r0, _, _ := purego.SyscallN(addr_QAbstractScrollArea_metric_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAbstractScrollArea) InitPainterDefault(painter *QPainter) {
	purego.SyscallN(addr_QAbstractScrollArea_initPainter_default, this.h, uintptr(painter.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
}

func (this *QAbstractScrollArea) RedirectedDefault(offset *QPoint) *QPaintDevice {
	__r, _, _ := purego.SyscallN(addr_QAbstractScrollArea_redirected_default, this.h, uintptr(offset.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(offset)
	return bindQPaintDevice(__r)
}

func (this *QAbstractScrollArea) SharedPainterDefault() *QPainter {
	__r, _, _ := purego.SyscallN(addr_QAbstractScrollArea_sharedPainter_default, this.h)
	runtime.KeepAlive(this)
	return bindQPainter(__r)
}

func (this *QAbstractScrollArea) InputMethodEventDefault(arg1 *QInputMethodEvent) {
	purego.SyscallN(addr_QAbstractScrollArea_inputMethodEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QAbstractScrollArea) InputMethodQueryDefault(arg1 InputMethodQuery) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QAbstractScrollArea_inputMethodQuery_default, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QAbstractScrollArea) FocusNextPrevChildDefault(next bool) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractScrollArea_focusNextPrevChild_default, this.h, uintptr(b2i(next)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QAbstractScrollArea) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QAbstractScrollArea_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractScrollArea) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QAbstractScrollArea_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractScrollArea) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QAbstractScrollArea_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractScrollArea) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QAbstractScrollArea_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QAbstractScrollArea) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QAbstractScrollArea_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QAbstractScrollArea) OnWindowTitleChanged(fn func(string)) int64 {
	id := registerCallback(func(title uintptr, title_len int64) { fn(pgStrNoFree(title, title_len)) })
	pg_QAbstractScrollArea_connect_windowTitleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractScrollArea) OnWindowIconChanged(fn func(*QIcon)) int64 {
	id := registerCallback(func(icon uintptr) { fn(UnsafeNewQIcon(unsafe.Pointer(icon))) })
	pg_QAbstractScrollArea_connect_windowIconChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractScrollArea) OnWindowIconTextChanged(fn func(string)) int64 {
	id := registerCallback(func(iconText uintptr, iconText_len int64) { fn(pgStrNoFree(iconText, iconText_len)) })
	pg_QAbstractScrollArea_connect_windowIconTextChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractScrollArea) OnCustomContextMenuRequested(fn func(*QPoint)) int64 {
	id := registerCallback(func(pos uintptr) { fn(UnsafeNewQPoint(unsafe.Pointer(pos))) })
	pg_QAbstractScrollArea_connect_customContextMenuRequested(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractScrollArea) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QAbstractScrollArea_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractScrollArea) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QAbstractScrollArea_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractScrollArea) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QAbstractScrollArea_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractScrollArea) Disconnect(handle int64) {
	pg_QAbstractScrollArea__disconnect(handle)
	unregisterCallback(handle)
}

func QAbstractScrollArea_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case 7049264599837532491:
		return 1
	case 729690261457684101:
		return 2
	case 4531996381736730688:
		return 3
	case 5217623514630858297:
		return 4
	case 1237236456312824991:
		return 5
	case -1971465653061861519:
		return 6
	case -4302917285655014661:
		return 7
	case 5290442163167458465:
		return 8
	case -332090898078543319:
		return 9
	case -4415920049297795213:
		return 10
	case -4029073905697023887:
		return 11
	case -4255494004099806581:
		return 12
	case 1919988340934511316:
		return 13
	case -6698043222227006015:
		return 14
	case 206122715919917589:
		return 15
	case -1781545137352745978:
		return 16
	case -3139904158373725006:
		return 17
	case 9107058682754591716:
		return 18
	case 1023791977679793185:
		return 19
	case -6118001150126564745:
		return 20
	case 503579067252048239:
		return 21
	case 4959376063826165943:
		return 22
	case -488693769771308531:
		return 23
	case -8659679893868524044:
		return 24
	case 2110182719375245575:
		return 25
	case 4311534060495737339:
		return 26
	case -8266762397995423915:
		return 27
	case 6398007404756346381:
		return 28
	case 772598475950050299:
		return 29
	case 975043767033552552:
		return 30
	case -385596010576048993:
		return 31
	case 6316270346668237895:
		return 32
	case -551070009623711640:
		return 33
	case -9161949150489525396:
		return 34
	case -1989937050019464783:
		return 35
	case -212068866860701141:
		return 36
	case 1697444140088555581:
		return 37
	case -6825369174751916:
		return 38
	case -199718707876719129:
		return 39
	case -2039166586246443706:
		return 40
	case 427629649267311777:
		return 41
	case 6477500296980050210:
		return 42
	case 4127240633310655048:
		return 43
	case 5792771328741774991:
		return 44
	case 5438819661942843222:
		return 45
	case 7345853097340709638:
		return 46
	case 5898328517505784191:
		return 47
	case -5409927581812374498:
		return 48
	case -1185315367526857833:
		return 49
	case 3769868229123868474:
		return 50
	case 7379029982170503008:
		return 51
	case -568369772982215962:
		return 52
	}
	return -1
}

func QAbstractScrollArea_enable(h uintptr, idx int32) { pg_QAbstractScrollArea__set_override(h, idx) }

func (this *QAbstractScrollArea) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QAbstractScrollArea) OnMinimumSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 7049264599837532491, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "minimumSizeHint"
}

func (this *QAbstractScrollArea) OnSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 729690261457684101, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sizeHint"
}

func (this *QAbstractScrollArea) OnSetupViewport(fn func(*QWidget)) {
	registerVirt(this.UnsafeGoID(), 4531996381736730688, func(viewport uintptr) { fn(UnsafeNewQWidget(unsafe.Pointer(viewport))) }) // slot "setupViewport"
}

func (this *QAbstractScrollArea) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(arg1 uintptr, arg2 uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(arg1)), UnsafeNewQEvent(unsafe.Pointer(arg2)))) }) // slot "eventFilter"
}

func (this *QAbstractScrollArea) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(arg1 uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(arg1)))) }) // slot "event"
}

func (this *QAbstractScrollArea) OnViewportEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), -1971465653061861519, func(arg1 uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(arg1)))) }) // slot "viewportEvent"
}

func (this *QAbstractScrollArea) OnResizeEvent(fn func(*QResizeEvent)) {
	registerVirt(this.UnsafeGoID(), -4302917285655014661, func(arg1 uintptr) { fn(UnsafeNewQResizeEvent(unsafe.Pointer(arg1))) }) // slot "resizeEvent"
}

func (this *QAbstractScrollArea) OnPaintEvent(fn func(*QPaintEvent)) {
	registerVirt(this.UnsafeGoID(), 5290442163167458465, func(arg1 uintptr) { fn(UnsafeNewQPaintEvent(unsafe.Pointer(arg1))) }) // slot "paintEvent"
}

func (this *QAbstractScrollArea) OnMousePressEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -332090898078543319, func(arg1 uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(arg1))) }) // slot "mousePressEvent"
}

func (this *QAbstractScrollArea) OnMouseReleaseEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4415920049297795213, func(arg1 uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(arg1))) }) // slot "mouseReleaseEvent"
}

func (this *QAbstractScrollArea) OnMouseDoubleClickEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4029073905697023887, func(arg1 uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(arg1))) }) // slot "mouseDoubleClickEvent"
}

func (this *QAbstractScrollArea) OnMouseMoveEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4255494004099806581, func(arg1 uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(arg1))) }) // slot "mouseMoveEvent"
}

func (this *QAbstractScrollArea) OnWheelEvent(fn func(*QWheelEvent)) {
	registerVirt(this.UnsafeGoID(), 1919988340934511316, func(arg1 uintptr) { fn(UnsafeNewQWheelEvent(unsafe.Pointer(arg1))) }) // slot "wheelEvent"
}

func (this *QAbstractScrollArea) OnContextMenuEvent(fn func(*QContextMenuEvent)) {
	registerVirt(this.UnsafeGoID(), -6698043222227006015, func(arg1 uintptr) { fn(UnsafeNewQContextMenuEvent(unsafe.Pointer(arg1))) }) // slot "contextMenuEvent"
}

func (this *QAbstractScrollArea) OnDragEnterEvent(fn func(*QDragEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 206122715919917589, func(arg1 uintptr) { fn(UnsafeNewQDragEnterEvent(unsafe.Pointer(arg1))) }) // slot "dragEnterEvent"
}

func (this *QAbstractScrollArea) OnDragMoveEvent(fn func(*QDragMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -1781545137352745978, func(arg1 uintptr) { fn(UnsafeNewQDragMoveEvent(unsafe.Pointer(arg1))) }) // slot "dragMoveEvent"
}

func (this *QAbstractScrollArea) OnDragLeaveEvent(fn func(*QDragLeaveEvent)) {
	registerVirt(this.UnsafeGoID(), -3139904158373725006, func(arg1 uintptr) { fn(UnsafeNewQDragLeaveEvent(unsafe.Pointer(arg1))) }) // slot "dragLeaveEvent"
}

func (this *QAbstractScrollArea) OnDropEvent(fn func(*QDropEvent)) {
	registerVirt(this.UnsafeGoID(), 9107058682754591716, func(arg1 uintptr) { fn(UnsafeNewQDropEvent(unsafe.Pointer(arg1))) }) // slot "dropEvent"
}

func (this *QAbstractScrollArea) OnKeyPressEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 1023791977679793185, func(arg1 uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(arg1))) }) // slot "keyPressEvent"
}

func (this *QAbstractScrollArea) OnScrollContentsBy(fn func(int, int)) {
	registerVirt(this.UnsafeGoID(), -6118001150126564745, func(dx int32, dy int32) { fn(int(dx), int(dy)) }) // slot "scrollContentsBy"
}

func (this *QAbstractScrollArea) OnViewportSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 503579067252048239, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "viewportSizeHint"
}

func (this *QAbstractScrollArea) OnChangeEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 4959376063826165943, func(arg1 uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(arg1))) }) // slot "changeEvent"
}

func (this *QAbstractScrollArea) OnInitStyleOption(fn func(*QStyleOptionFrame)) {
	registerVirt(this.UnsafeGoID(), -488693769771308531, func(option uintptr) { fn(UnsafeNewQStyleOptionFrame(unsafe.Pointer(option))) }) // slot "initStyleOption"
}

func (this *QAbstractScrollArea) OnDevType(fn func() int) {
	registerVirt(this.UnsafeGoID(), -8659679893868524044, func() int32 { return int32(fn()) }) // slot "devType"
}

func (this *QAbstractScrollArea) OnSetVisible(fn func(bool)) {
	registerVirt(this.UnsafeGoID(), 2110182719375245575, func(visible int32) { fn(visible != 0) }) // slot "setVisible"
}

func (this *QAbstractScrollArea) OnHeightForWidth(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), 4311534060495737339, func(arg1 int32) int32 { return int32(fn(int(arg1))) }) // slot "heightForWidth"
}

func (this *QAbstractScrollArea) OnHasHeightForWidth(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -8266762397995423915, func() int32 { return b2i(fn()) }) // slot "hasHeightForWidth"
}

func (this *QAbstractScrollArea) OnPaintEngine(fn func() *QPaintEngine) {
	registerVirt(this.UnsafeGoID(), 6398007404756346381, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "paintEngine"
}

func (this *QAbstractScrollArea) OnKeyReleaseEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 772598475950050299, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyReleaseEvent"
}

func (this *QAbstractScrollArea) OnFocusInEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), 975043767033552552, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusInEvent"
}

func (this *QAbstractScrollArea) OnFocusOutEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), -385596010576048993, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusOutEvent"
}

func (this *QAbstractScrollArea) OnEnterEvent(fn func(*QEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 6316270346668237895, func(event uintptr) { fn(UnsafeNewQEnterEvent(unsafe.Pointer(event))) }) // slot "enterEvent"
}

func (this *QAbstractScrollArea) OnLeaveEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), -551070009623711640, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "leaveEvent"
}

func (this *QAbstractScrollArea) OnMoveEvent(fn func(*QMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -9161949150489525396, func(event uintptr) { fn(UnsafeNewQMoveEvent(unsafe.Pointer(event))) }) // slot "moveEvent"
}

func (this *QAbstractScrollArea) OnCloseEvent(fn func(*QCloseEvent)) {
	registerVirt(this.UnsafeGoID(), -1989937050019464783, func(event uintptr) { fn(UnsafeNewQCloseEvent(unsafe.Pointer(event))) }) // slot "closeEvent"
}

func (this *QAbstractScrollArea) OnTabletEvent(fn func(*QTabletEvent)) {
	registerVirt(this.UnsafeGoID(), -212068866860701141, func(event uintptr) { fn(UnsafeNewQTabletEvent(unsafe.Pointer(event))) }) // slot "tabletEvent"
}

func (this *QAbstractScrollArea) OnActionEvent(fn func(*QActionEvent)) {
	registerVirt(this.UnsafeGoID(), 1697444140088555581, func(event uintptr) { fn(UnsafeNewQActionEvent(unsafe.Pointer(event))) }) // slot "actionEvent"
}

func (this *QAbstractScrollArea) OnShowEvent(fn func(*QShowEvent)) {
	registerVirt(this.UnsafeGoID(), -6825369174751916, func(event uintptr) { fn(UnsafeNewQShowEvent(unsafe.Pointer(event))) }) // slot "showEvent"
}

func (this *QAbstractScrollArea) OnHideEvent(fn func(*QHideEvent)) {
	registerVirt(this.UnsafeGoID(), -199718707876719129, func(event uintptr) { fn(UnsafeNewQHideEvent(unsafe.Pointer(event))) }) // slot "hideEvent"
}

func (this *QAbstractScrollArea) OnNativeEvent(fn func(*QByteArray, unsafe.Pointer, *int64) bool) {
	registerVirt(this.UnsafeGoID(), -2039166586246443706, func(eventType uintptr, message uintptr, result uintptr) int32 { return b2i(fn(UnsafeNewQByteArray(unsafe.Pointer(eventType)), unsafe.Pointer(message), (*int64)(unsafe.Pointer(result)))) }) // slot "nativeEvent"
}

func (this *QAbstractScrollArea) OnMetric(fn func(QPaintDevice__PaintDeviceMetric) int) {
	registerVirt(this.UnsafeGoID(), 427629649267311777, func(arg1 int32) int32 { return int32(fn(QPaintDevice__PaintDeviceMetric(arg1))) }) // slot "metric"
}

func (this *QAbstractScrollArea) OnInitPainter(fn func(*QPainter)) {
	registerVirt(this.UnsafeGoID(), 6477500296980050210, func(painter uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter))) }) // slot "initPainter"
}

func (this *QAbstractScrollArea) OnRedirected(fn func(*QPoint) *QPaintDevice) {
	registerVirt(this.UnsafeGoID(), 4127240633310655048, func(offset uintptr) uintptr { return uintptr(fn(UnsafeNewQPoint(unsafe.Pointer(offset))).UnsafePointer()) }) // slot "redirected"
}

func (this *QAbstractScrollArea) OnSharedPainter(fn func() *QPainter) {
	registerVirt(this.UnsafeGoID(), 5792771328741774991, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sharedPainter"
}

func (this *QAbstractScrollArea) OnInputMethodEvent(fn func(*QInputMethodEvent)) {
	registerVirt(this.UnsafeGoID(), 5438819661942843222, func(arg1 uintptr) { fn(UnsafeNewQInputMethodEvent(unsafe.Pointer(arg1))) }) // slot "inputMethodEvent"
}

func (this *QAbstractScrollArea) OnInputMethodQuery(fn func(InputMethodQuery) *QVariant) {
	registerVirt(this.UnsafeGoID(), 7345853097340709638, func(arg1 int32) uintptr { return uintptr(fn(InputMethodQuery(arg1)).UnsafePointer()) }) // slot "inputMethodQuery"
}

func (this *QAbstractScrollArea) OnFocusNextPrevChild(fn func(bool) bool) {
	registerVirt(this.UnsafeGoID(), 5898328517505784191, func(next int32) int32 { return b2i(fn(next != 0)) }) // slot "focusNextPrevChild"
}

func (this *QAbstractScrollArea) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QAbstractScrollArea) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QAbstractScrollArea) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QAbstractScrollArea) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QAbstractScrollArea) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QAbstractScrollArea) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QAbstractScrollArea_delete(this.h)
	}
	this.UnsafeRelease()
}
