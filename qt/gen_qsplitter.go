package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QSplitter_chain uintptr
	addr_QSplitter_new uintptr
	addr_QSplitter_new2 uintptr
	addr_QSplitter_new3 uintptr
	addr_QSplitter_new4 uintptr
	addr_QSplitter_tr uintptr
	addr_QSplitter_addWidget uintptr
	addr_QSplitter_insertWidget uintptr
	addr_QSplitter_replaceWidget uintptr
	addr_QSplitter_setOrientation uintptr
	addr_QSplitter_orientation uintptr
	addr_QSplitter_setChildrenCollapsible uintptr
	addr_QSplitter_childrenCollapsible uintptr
	addr_QSplitter_setCollapsible uintptr
	addr_QSplitter_isCollapsible uintptr
	addr_QSplitter_setOpaqueResize uintptr
	addr_QSplitter_opaqueResize uintptr
	addr_QSplitter_refresh uintptr
	addr_QSplitter_sizes uintptr
	addr_QSplitter_setSizes uintptr
	addr_QSplitter_saveState uintptr
	addr_QSplitter_restoreState uintptr
	addr_QSplitter_handleWidth uintptr
	addr_QSplitter_setHandleWidth uintptr
	addr_QSplitter_indexOf uintptr
	addr_QSplitter_widget uintptr
	addr_QSplitter_count uintptr
	addr_QSplitter_getRange uintptr
	addr_QSplitter_handle uintptr
	addr_QSplitter_setStretchFactor uintptr
	addr_QSplitter_splitterMoved uintptr
	addr_QSplitter_moveSplitter uintptr
	addr_QSplitter_setRubberBand uintptr
	addr_QSplitter_closestLegalPosition uintptr
	addr_QSplitter_drawFrame uintptr
	addr_QSplitter_setTabOrder uintptr
	addr_QSplitter_mouseGrabber uintptr
	addr_QSplitter_keyboardGrabber uintptr
	addr_QSplitter_find uintptr
	addr_QSplitter_createWindowContainer uintptr
	addr_QSplitter_updateMicroFocus uintptr
	addr_QSplitter_create uintptr
	addr_QSplitter_destroy uintptr
	addr_QSplitter_focusNextChild uintptr
	addr_QSplitter_focusPreviousChild uintptr
	addr_QSplitter_connect uintptr
	addr_QSplitter_connect2 uintptr
	addr_QSplitter_sender uintptr
	addr_QSplitter_senderSignalIndex uintptr
	addr_QSplitter_receivers uintptr
	addr_QSplitter_isSignalConnected uintptr
	pg_QSplitter_devicePixelRatioFScale func() float64
	pg_QSplitter_encodeMetricF func(int32, float64) int32
	pg_QSplitter_getDecodedMetricF func(uintptr, int32, int32) float64
	addr_QSplitter_tr2 uintptr
	addr_QSplitter_tr3 uintptr
	addr_QSplitter_setOpaqueResize2 uintptr
	addr_QSplitter_createWindowContainer2 uintptr
	addr_QSplitter_createWindowContainer3 uintptr
	addr_QSplitter_updateMicroFocus2 uintptr
	addr_QSplitter_create2 uintptr
	addr_QSplitter_create3 uintptr
	addr_QSplitter_create4 uintptr
	addr_QSplitter_destroy2 uintptr
	addr_QSplitter_destroy3 uintptr
	addr_QSplitter_connect3 uintptr
	addr_QSplitter_connect4 uintptr
	addr_QSplitter_metaObject_default uintptr
	addr_QSplitter_sizeHint_default uintptr
	addr_QSplitter_minimumSizeHint_default uintptr
	addr_QSplitter_createHandle_default uintptr
	addr_QSplitter_childEvent_default uintptr
	addr_QSplitter_event_default uintptr
	addr_QSplitter_resizeEvent_default uintptr
	addr_QSplitter_changeEvent_default uintptr
	addr_QSplitter_paintEvent_default uintptr
	addr_QSplitter_initStyleOption_default uintptr
	addr_QSplitter_devType_default uintptr
	addr_QSplitter_setVisible_default uintptr
	addr_QSplitter_heightForWidth_default uintptr
	addr_QSplitter_hasHeightForWidth_default uintptr
	addr_QSplitter_paintEngine_default uintptr
	addr_QSplitter_mousePressEvent_default uintptr
	addr_QSplitter_mouseReleaseEvent_default uintptr
	addr_QSplitter_mouseDoubleClickEvent_default uintptr
	addr_QSplitter_mouseMoveEvent_default uintptr
	addr_QSplitter_wheelEvent_default uintptr
	addr_QSplitter_keyPressEvent_default uintptr
	addr_QSplitter_keyReleaseEvent_default uintptr
	addr_QSplitter_focusInEvent_default uintptr
	addr_QSplitter_focusOutEvent_default uintptr
	addr_QSplitter_enterEvent_default uintptr
	addr_QSplitter_leaveEvent_default uintptr
	addr_QSplitter_moveEvent_default uintptr
	addr_QSplitter_closeEvent_default uintptr
	addr_QSplitter_contextMenuEvent_default uintptr
	addr_QSplitter_tabletEvent_default uintptr
	addr_QSplitter_actionEvent_default uintptr
	addr_QSplitter_dragEnterEvent_default uintptr
	addr_QSplitter_dragMoveEvent_default uintptr
	addr_QSplitter_dragLeaveEvent_default uintptr
	addr_QSplitter_dropEvent_default uintptr
	addr_QSplitter_showEvent_default uintptr
	addr_QSplitter_hideEvent_default uintptr
	addr_QSplitter_metric_default uintptr
	addr_QSplitter_initPainter_default uintptr
	addr_QSplitter_redirected_default uintptr
	addr_QSplitter_sharedPainter_default uintptr
	addr_QSplitter_inputMethodEvent_default uintptr
	addr_QSplitter_inputMethodQuery_default uintptr
	addr_QSplitter_focusNextPrevChild_default uintptr
	addr_QSplitter_eventFilter_default uintptr
	addr_QSplitter_timerEvent_default uintptr
	addr_QSplitter_customEvent_default uintptr
	addr_QSplitter_connectNotify_default uintptr
	addr_QSplitter_disconnectNotify_default uintptr
	pg_QSplitter_connect_splitterMoved func(uintptr, int64)
	pg_QSplitter_connect_windowTitleChanged func(uintptr, int64)
	pg_QSplitter_connect_windowIconChanged func(uintptr, int64)
	pg_QSplitter_connect_windowIconTextChanged func(uintptr, int64)
	pg_QSplitter_connect_customContextMenuRequested func(uintptr, int64)
	pg_QSplitter_connect_destroyed func(uintptr, int64)
	pg_QSplitter_connect_destroyed2 func(uintptr, int64)
	pg_QSplitter_connect_objectNameChanged func(uintptr, int64)
	pg_QSplitter__disconnect func(int64)
	pg_QSplitter__set_override func(uintptr, int32)
	pg_QSplitter_delete func(uintptr)
)

func init() {
	symaddr(&addr_QSplitter_chain, "QSplitter_chain")
	symaddr(&addr_QSplitter_new, "QSplitter_new")
	symaddr(&addr_QSplitter_new2, "QSplitter_new2")
	symaddr(&addr_QSplitter_new3, "QSplitter_new3")
	symaddr(&addr_QSplitter_new4, "QSplitter_new4")
	symaddr(&addr_QSplitter_tr, "QSplitter_tr")
	symaddr(&addr_QSplitter_addWidget, "QSplitter_addWidget")
	symaddr(&addr_QSplitter_insertWidget, "QSplitter_insertWidget")
	symaddr(&addr_QSplitter_replaceWidget, "QSplitter_replaceWidget")
	symaddr(&addr_QSplitter_setOrientation, "QSplitter_setOrientation")
	symaddr(&addr_QSplitter_orientation, "QSplitter_orientation")
	symaddr(&addr_QSplitter_setChildrenCollapsible, "QSplitter_setChildrenCollapsible")
	symaddr(&addr_QSplitter_childrenCollapsible, "QSplitter_childrenCollapsible")
	symaddr(&addr_QSplitter_setCollapsible, "QSplitter_setCollapsible")
	symaddr(&addr_QSplitter_isCollapsible, "QSplitter_isCollapsible")
	symaddr(&addr_QSplitter_setOpaqueResize, "QSplitter_setOpaqueResize")
	symaddr(&addr_QSplitter_opaqueResize, "QSplitter_opaqueResize")
	symaddr(&addr_QSplitter_refresh, "QSplitter_refresh")
	symaddr(&addr_QSplitter_sizes, "QSplitter_sizes")
	symaddr(&addr_QSplitter_setSizes, "QSplitter_setSizes")
	symaddr(&addr_QSplitter_saveState, "QSplitter_saveState")
	symaddr(&addr_QSplitter_restoreState, "QSplitter_restoreState")
	symaddr(&addr_QSplitter_handleWidth, "QSplitter_handleWidth")
	symaddr(&addr_QSplitter_setHandleWidth, "QSplitter_setHandleWidth")
	symaddr(&addr_QSplitter_indexOf, "QSplitter_indexOf")
	symaddr(&addr_QSplitter_widget, "QSplitter_widget")
	symaddr(&addr_QSplitter_count, "QSplitter_count")
	symaddr(&addr_QSplitter_getRange, "QSplitter_getRange")
	symaddr(&addr_QSplitter_handle, "QSplitter_handle")
	symaddr(&addr_QSplitter_setStretchFactor, "QSplitter_setStretchFactor")
	symaddr(&addr_QSplitter_splitterMoved, "QSplitter_splitterMoved")
	symaddr(&addr_QSplitter_moveSplitter, "QSplitter_moveSplitter")
	symaddr(&addr_QSplitter_setRubberBand, "QSplitter_setRubberBand")
	symaddr(&addr_QSplitter_closestLegalPosition, "QSplitter_closestLegalPosition")
	symaddr(&addr_QSplitter_drawFrame, "QSplitter_drawFrame")
	symaddr(&addr_QSplitter_setTabOrder, "QSplitter_setTabOrder")
	symaddr(&addr_QSplitter_mouseGrabber, "QSplitter_mouseGrabber")
	symaddr(&addr_QSplitter_keyboardGrabber, "QSplitter_keyboardGrabber")
	symaddr(&addr_QSplitter_find, "QSplitter_find")
	symaddr(&addr_QSplitter_createWindowContainer, "QSplitter_createWindowContainer")
	symaddr(&addr_QSplitter_updateMicroFocus, "QSplitter_updateMicroFocus")
	symaddr(&addr_QSplitter_create, "QSplitter_create")
	symaddr(&addr_QSplitter_destroy, "QSplitter_destroy")
	symaddr(&addr_QSplitter_focusNextChild, "QSplitter_focusNextChild")
	symaddr(&addr_QSplitter_focusPreviousChild, "QSplitter_focusPreviousChild")
	symaddr(&addr_QSplitter_connect, "QSplitter_connect")
	symaddr(&addr_QSplitter_connect2, "QSplitter_connect2")
	symaddr(&addr_QSplitter_sender, "QSplitter_sender")
	symaddr(&addr_QSplitter_senderSignalIndex, "QSplitter_senderSignalIndex")
	symaddr(&addr_QSplitter_receivers, "QSplitter_receivers")
	symaddr(&addr_QSplitter_isSignalConnected, "QSplitter_isSignalConnected")
	reg(&pg_QSplitter_devicePixelRatioFScale, "QSplitter_devicePixelRatioFScale")
	reg(&pg_QSplitter_encodeMetricF, "QSplitter_encodeMetricF")
	reg(&pg_QSplitter_getDecodedMetricF, "QSplitter_getDecodedMetricF")
	symaddr(&addr_QSplitter_tr2, "QSplitter_tr2")
	symaddr(&addr_QSplitter_tr3, "QSplitter_tr3")
	symaddr(&addr_QSplitter_setOpaqueResize2, "QSplitter_setOpaqueResize2")
	symaddr(&addr_QSplitter_createWindowContainer2, "QSplitter_createWindowContainer2")
	symaddr(&addr_QSplitter_createWindowContainer3, "QSplitter_createWindowContainer3")
	symaddr(&addr_QSplitter_updateMicroFocus2, "QSplitter_updateMicroFocus2")
	symaddr(&addr_QSplitter_create2, "QSplitter_create2")
	symaddr(&addr_QSplitter_create3, "QSplitter_create3")
	symaddr(&addr_QSplitter_create4, "QSplitter_create4")
	symaddr(&addr_QSplitter_destroy2, "QSplitter_destroy2")
	symaddr(&addr_QSplitter_destroy3, "QSplitter_destroy3")
	symaddr(&addr_QSplitter_connect3, "QSplitter_connect3")
	symaddr(&addr_QSplitter_connect4, "QSplitter_connect4")
	symaddr(&addr_QSplitter_metaObject_default, "QSplitter_metaObject_default")
	symaddr(&addr_QSplitter_sizeHint_default, "QSplitter_sizeHint_default")
	symaddr(&addr_QSplitter_minimumSizeHint_default, "QSplitter_minimumSizeHint_default")
	symaddr(&addr_QSplitter_createHandle_default, "QSplitter_createHandle_default")
	symaddr(&addr_QSplitter_childEvent_default, "QSplitter_childEvent_default")
	symaddr(&addr_QSplitter_event_default, "QSplitter_event_default")
	symaddr(&addr_QSplitter_resizeEvent_default, "QSplitter_resizeEvent_default")
	symaddr(&addr_QSplitter_changeEvent_default, "QSplitter_changeEvent_default")
	symaddr(&addr_QSplitter_paintEvent_default, "QSplitter_paintEvent_default")
	symaddr(&addr_QSplitter_initStyleOption_default, "QSplitter_initStyleOption_default")
	symaddr(&addr_QSplitter_devType_default, "QSplitter_devType_default")
	symaddr(&addr_QSplitter_setVisible_default, "QSplitter_setVisible_default")
	symaddr(&addr_QSplitter_heightForWidth_default, "QSplitter_heightForWidth_default")
	symaddr(&addr_QSplitter_hasHeightForWidth_default, "QSplitter_hasHeightForWidth_default")
	symaddr(&addr_QSplitter_paintEngine_default, "QSplitter_paintEngine_default")
	symaddr(&addr_QSplitter_mousePressEvent_default, "QSplitter_mousePressEvent_default")
	symaddr(&addr_QSplitter_mouseReleaseEvent_default, "QSplitter_mouseReleaseEvent_default")
	symaddr(&addr_QSplitter_mouseDoubleClickEvent_default, "QSplitter_mouseDoubleClickEvent_default")
	symaddr(&addr_QSplitter_mouseMoveEvent_default, "QSplitter_mouseMoveEvent_default")
	symaddr(&addr_QSplitter_wheelEvent_default, "QSplitter_wheelEvent_default")
	symaddr(&addr_QSplitter_keyPressEvent_default, "QSplitter_keyPressEvent_default")
	symaddr(&addr_QSplitter_keyReleaseEvent_default, "QSplitter_keyReleaseEvent_default")
	symaddr(&addr_QSplitter_focusInEvent_default, "QSplitter_focusInEvent_default")
	symaddr(&addr_QSplitter_focusOutEvent_default, "QSplitter_focusOutEvent_default")
	symaddr(&addr_QSplitter_enterEvent_default, "QSplitter_enterEvent_default")
	symaddr(&addr_QSplitter_leaveEvent_default, "QSplitter_leaveEvent_default")
	symaddr(&addr_QSplitter_moveEvent_default, "QSplitter_moveEvent_default")
	symaddr(&addr_QSplitter_closeEvent_default, "QSplitter_closeEvent_default")
	symaddr(&addr_QSplitter_contextMenuEvent_default, "QSplitter_contextMenuEvent_default")
	symaddr(&addr_QSplitter_tabletEvent_default, "QSplitter_tabletEvent_default")
	symaddr(&addr_QSplitter_actionEvent_default, "QSplitter_actionEvent_default")
	symaddr(&addr_QSplitter_dragEnterEvent_default, "QSplitter_dragEnterEvent_default")
	symaddr(&addr_QSplitter_dragMoveEvent_default, "QSplitter_dragMoveEvent_default")
	symaddr(&addr_QSplitter_dragLeaveEvent_default, "QSplitter_dragLeaveEvent_default")
	symaddr(&addr_QSplitter_dropEvent_default, "QSplitter_dropEvent_default")
	symaddr(&addr_QSplitter_showEvent_default, "QSplitter_showEvent_default")
	symaddr(&addr_QSplitter_hideEvent_default, "QSplitter_hideEvent_default")
	symaddr(&addr_QSplitter_metric_default, "QSplitter_metric_default")
	symaddr(&addr_QSplitter_initPainter_default, "QSplitter_initPainter_default")
	symaddr(&addr_QSplitter_redirected_default, "QSplitter_redirected_default")
	symaddr(&addr_QSplitter_sharedPainter_default, "QSplitter_sharedPainter_default")
	symaddr(&addr_QSplitter_inputMethodEvent_default, "QSplitter_inputMethodEvent_default")
	symaddr(&addr_QSplitter_inputMethodQuery_default, "QSplitter_inputMethodQuery_default")
	symaddr(&addr_QSplitter_focusNextPrevChild_default, "QSplitter_focusNextPrevChild_default")
	symaddr(&addr_QSplitter_eventFilter_default, "QSplitter_eventFilter_default")
	symaddr(&addr_QSplitter_timerEvent_default, "QSplitter_timerEvent_default")
	symaddr(&addr_QSplitter_customEvent_default, "QSplitter_customEvent_default")
	symaddr(&addr_QSplitter_connectNotify_default, "QSplitter_connectNotify_default")
	symaddr(&addr_QSplitter_disconnectNotify_default, "QSplitter_disconnectNotify_default")
	reg(&pg_QSplitter_connect_splitterMoved, "QSplitter_connect_splitterMoved")
	reg(&pg_QSplitter_connect_windowTitleChanged, "QSplitter_connect_windowTitleChanged")
	reg(&pg_QSplitter_connect_windowIconChanged, "QSplitter_connect_windowIconChanged")
	reg(&pg_QSplitter_connect_windowIconTextChanged, "QSplitter_connect_windowIconTextChanged")
	reg(&pg_QSplitter_connect_customContextMenuRequested, "QSplitter_connect_customContextMenuRequested")
	reg(&pg_QSplitter_connect_destroyed, "QSplitter_connect_destroyed")
	reg(&pg_QSplitter_connect_destroyed2, "QSplitter_connect_destroyed2")
	reg(&pg_QSplitter_connect_objectNameChanged, "QSplitter_connect_objectNameChanged")
	reg(&pg_QSplitter__disconnect, "QSplitter__disconnect")
	reg(&pg_QSplitter__set_override, "QSplitter__set_override")
	reg(&pg_QSplitter_delete, "QSplitter_delete")
}

type QSplitter struct {
	h uintptr
	*QFrame
}

func bindQSplitter(h uintptr) *QSplitter {
	if h == 0 {
		return nil
	}
	var __c [3]uintptr
	purego.SyscallN(addr_QSplitter_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QSplitter{h: h, QFrame: &QFrame{h: __c[0], QWidget: &QWidget{h: __c[1], QObject: &QObject{h: __c[2]}}}}
}

func newownedQSplitter(h uintptr) *QSplitter { return bindQSplitter(h) }

func (this *QSplitter) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QSplitter) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQSplitter(h unsafe.Pointer) *QSplitter { return bindQSplitter(uintptr(h)) }
func UnsafeNewQSplitterOwned(h unsafe.Pointer) *QSplitter { return newownedQSplitter(uintptr(h)) }

func (this *QSplitter) UnsafeRelease() {
	this.h = 0
	if this.QFrame != nil {
		this.QFrame.UnsafeRelease()
	}
}

func NewQSplitter(parent *QWidget) *QSplitter {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QSplitter_new, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQSplitter(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QSplitter_virtIdx, QSplitter_enable)
	return __o
}

func NewQSplitter2() *QSplitter {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QSplitter_new2, uintptr(__goid))
	__o := newownedQSplitter(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QSplitter_virtIdx, QSplitter_enable)
	return __o
}

func NewQSplitter3(arg1 Orientation) *QSplitter {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QSplitter_new3, uintptr(__goid), uintptr(arg1))
	__o := newownedQSplitter(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QSplitter_virtIdx, QSplitter_enable)
	return __o
}

func NewQSplitter4(arg1 Orientation, parent *QWidget) *QSplitter {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QSplitter_new4, uintptr(__goid), uintptr(arg1), uintptr(parent.UnsafePointer()))
	__o := newownedQSplitter(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QSplitter_virtIdx, QSplitter_enable)
	return __o
}

func QSplitter_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QSplitter_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QSplitter) AddWidget(widget *QWidget) {
	purego.SyscallN(addr_QSplitter_addWidget, this.h, uintptr(widget.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(widget)
}

func (this *QSplitter) InsertWidget(index int, widget *QWidget) {
	purego.SyscallN(addr_QSplitter_insertWidget, this.h, uintptr(index), uintptr(widget.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(widget)
}

func (this *QSplitter) ReplaceWidget(index int, widget *QWidget) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QSplitter_replaceWidget, this.h, uintptr(index), uintptr(widget.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(widget)
	return bindQWidget(__r)
}

func (this *QSplitter) SetOrientation(arg1 Orientation) {
	purego.SyscallN(addr_QSplitter_setOrientation, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QSplitter) Orientation() Orientation {
	__r0, _, _ := purego.SyscallN(addr_QSplitter_orientation, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return Orientation(__r)
}

func (this *QSplitter) SetChildrenCollapsible(arg1 bool) {
	purego.SyscallN(addr_QSplitter_setChildrenCollapsible, this.h, uintptr(b2i(arg1)))
	runtime.KeepAlive(this)
}

func (this *QSplitter) ChildrenCollapsible() bool {
	__r0, _, _ := purego.SyscallN(addr_QSplitter_childrenCollapsible, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QSplitter) SetCollapsible(index int, arg2 bool) {
	purego.SyscallN(addr_QSplitter_setCollapsible, this.h, uintptr(index), uintptr(b2i(arg2)))
	runtime.KeepAlive(this)
}

func (this *QSplitter) IsCollapsible(index int) bool {
	__r0, _, _ := purego.SyscallN(addr_QSplitter_isCollapsible, this.h, uintptr(index))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QSplitter) SetOpaqueResize() {
	purego.SyscallN(addr_QSplitter_setOpaqueResize, this.h)
	runtime.KeepAlive(this)
}

func (this *QSplitter) OpaqueResize() bool {
	__r0, _, _ := purego.SyscallN(addr_QSplitter_opaqueResize, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QSplitter) Refresh() {
	purego.SyscallN(addr_QSplitter_refresh, this.h)
	runtime.KeepAlive(this)
}

func (this *QSplitter) Sizes() []int {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QSplitter_sizes, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]int, __n)
	__sl := unsafe.Slice((*int32)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = int(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QSplitter) SetSizes(list []int) {
	list_arr := make([]int32, len(list))
	for __i, __e := range list {
		list_arr[__i] = int32(__e)
	}
	var list_p uintptr
	if len(list_arr) > 0 {
		list_p = uintptr(unsafe.Pointer(&list_arr[0]))
	}
	purego.SyscallN(addr_QSplitter_setSizes, this.h, list_p, uintptr(len(list)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(list_arr)
}

func (this *QSplitter) SaveState() *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QSplitter_saveState, this.h)
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QSplitter) RestoreState(state *QByteArray) bool {
	__r0, _, _ := purego.SyscallN(addr_QSplitter_restoreState, this.h, uintptr(state.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(state)
	return __r != 0
}

func (this *QSplitter) HandleWidth() int {
	__r0, _, _ := purego.SyscallN(addr_QSplitter_handleWidth, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QSplitter) SetHandleWidth(arg1 int) {
	purego.SyscallN(addr_QSplitter_setHandleWidth, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QSplitter) IndexOf(w *QWidget) int {
	__r0, _, _ := purego.SyscallN(addr_QSplitter_indexOf, this.h, uintptr(w.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(w)
	return int(__r)
}

func (this *QSplitter) Widget(index int) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QSplitter_widget, this.h, uintptr(index))
	runtime.KeepAlive(this)
	return bindQWidget(__r)
}

func (this *QSplitter) Count() int {
	__r0, _, _ := purego.SyscallN(addr_QSplitter_count, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QSplitter) GetRange(index int) (int, int) {
	var arg2_out int32
	var arg3_out int32
	purego.SyscallN(addr_QSplitter_getRange, this.h, uintptr(index), uintptr(unsafe.Pointer(&arg2_out)), uintptr(unsafe.Pointer(&arg3_out)))
	runtime.KeepAlive(this)
	return int(arg2_out), int(arg3_out)
}

func (this *QSplitter) Handle(index int) *QSplitterHandle {
	__r, _, _ := purego.SyscallN(addr_QSplitter_handle, this.h, uintptr(index))
	runtime.KeepAlive(this)
	return bindQSplitterHandle(__r)
}

func (this *QSplitter) SetStretchFactor(index int, stretch int) {
	purego.SyscallN(addr_QSplitter_setStretchFactor, this.h, uintptr(index), uintptr(stretch))
	runtime.KeepAlive(this)
}

func (this *QSplitter) SplitterMoved(pos int, index int) {
	purego.SyscallN(addr_QSplitter_splitterMoved, this.h, uintptr(pos), uintptr(index))
	runtime.KeepAlive(this)
}

func (this *QSplitter) MoveSplitter(pos int, index int) {
	purego.SyscallN(addr_QSplitter_moveSplitter, this.h, uintptr(pos), uintptr(index))
	runtime.KeepAlive(this)
}

func (this *QSplitter) SetRubberBand(position int) {
	purego.SyscallN(addr_QSplitter_setRubberBand, this.h, uintptr(position))
	runtime.KeepAlive(this)
}

func (this *QSplitter) ClosestLegalPosition(arg1 int, arg2 int) int {
	__r0, _, _ := purego.SyscallN(addr_QSplitter_closestLegalPosition, this.h, uintptr(arg1), uintptr(arg2))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QSplitter) DrawFrame(arg1 *QPainter) {
	purego.SyscallN(addr_QSplitter_drawFrame, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func QSplitter_SetTabOrder(arg1 *QWidget, arg2 *QWidget) {
	purego.SyscallN(addr_QSplitter_setTabOrder, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()))
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
}

func QSplitter_MouseGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QSplitter_mouseGrabber)
	return bindQWidget(__r)
}

func QSplitter_KeyboardGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QSplitter_keyboardGrabber)
	return bindQWidget(__r)
}

func QSplitter_Find(arg1 uint64) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QSplitter_find, uintptr(arg1))
	return bindQWidget(__r)
}

func QSplitter_CreateWindowContainer(window *QWindow) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QSplitter_createWindowContainer, uintptr(window.UnsafePointer()))
	runtime.KeepAlive(window)
	return bindQWidget(__r)
}

func (this *QSplitter) UpdateMicroFocus() {
	purego.SyscallN(addr_QSplitter_updateMicroFocus, this.h)
	runtime.KeepAlive(this)
}

func (this *QSplitter) Create() {
	purego.SyscallN(addr_QSplitter_create, this.h)
	runtime.KeepAlive(this)
}

func (this *QSplitter) Destroy() {
	purego.SyscallN(addr_QSplitter_destroy, this.h)
	runtime.KeepAlive(this)
}

func (this *QSplitter) FocusNextChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QSplitter_focusNextChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QSplitter) FocusPreviousChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QSplitter_focusPreviousChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QSplitter_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QSplitter_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QSplitter_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QSplitter_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QSplitter) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QSplitter_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QSplitter) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QSplitter_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QSplitter) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QSplitter_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QSplitter) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QSplitter_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QSplitter_DevicePixelRatioFScale() float64 {
	__r := pg_QSplitter_devicePixelRatioFScale()
	return __r
}

func QSplitter_EncodeMetricF(metric QPaintDevice__PaintDeviceMetric, value float64) int {
	__r := pg_QSplitter_encodeMetricF(int32(metric), value)
	return int(__r)
}

func (this *QSplitter) GetDecodedMetricF(metricA QPaintDevice__PaintDeviceMetric, metricB QPaintDevice__PaintDeviceMetric) float64 {
	__r := pg_QSplitter_getDecodedMetricF(this.h, int32(metricA), int32(metricB))
	runtime.KeepAlive(this)
	return __r
}

func QSplitter_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QSplitter_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QSplitter_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QSplitter_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func (this *QSplitter) SetOpaqueResizeWithOpaque(opaque bool) {
	purego.SyscallN(addr_QSplitter_setOpaqueResize2, this.h, uintptr(b2i(opaque)))
	runtime.KeepAlive(this)
}

func QSplitter_CreateWindowContainer2(window *QWindow, parent *QWidget) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QSplitter_createWindowContainer2, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func QSplitter_CreateWindowContainer3(window *QWindow, parent *QWidget, flags WindowType) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QSplitter_createWindowContainer3, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()), uintptr(flags))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func (this *QSplitter) UpdateMicroFocusWithQuery(query InputMethodQuery) {
	purego.SyscallN(addr_QSplitter_updateMicroFocus2, this.h, uintptr(query))
	runtime.KeepAlive(this)
}

func (this *QSplitter) CreateWith(arg1 uint64) {
	purego.SyscallN(addr_QSplitter_create2, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QSplitter) Create2(arg1 uint64, initializeWindow bool) {
	purego.SyscallN(addr_QSplitter_create3, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)))
	runtime.KeepAlive(this)
}

func (this *QSplitter) Create3(arg1 uint64, initializeWindow bool, destroyOldWindow bool) {
	purego.SyscallN(addr_QSplitter_create4, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)), uintptr(b2i(destroyOldWindow)))
	runtime.KeepAlive(this)
}

func (this *QSplitter) DestroyWithDestroyWindow(destroyWindow bool) {
	purego.SyscallN(addr_QSplitter_destroy2, this.h, uintptr(b2i(destroyWindow)))
	runtime.KeepAlive(this)
}

func (this *QSplitter) Destroy2(destroyWindow bool, destroySubWindows bool) {
	purego.SyscallN(addr_QSplitter_destroy3, this.h, uintptr(b2i(destroyWindow)), uintptr(b2i(destroySubWindows)))
	runtime.KeepAlive(this)
}

func QSplitter_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QSplitter_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QSplitter_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QSplitter_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QSplitter) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QSplitter_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QSplitter) SizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QSplitter_sizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QSplitter) MinimumSizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QSplitter_minimumSizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QSplitter) CreateHandleDefault() *QSplitterHandle {
	__r, _, _ := purego.SyscallN(addr_QSplitter_createHandle_default, this.h)
	runtime.KeepAlive(this)
	return bindQSplitterHandle(__r)
}

func (this *QSplitter) ChildEventDefault(arg1 *QChildEvent) {
	purego.SyscallN(addr_QSplitter_childEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QSplitter) EventDefault(arg1 *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QSplitter_event_default, this.h, uintptr(arg1.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return __r != 0
}

func (this *QSplitter) ResizeEventDefault(arg1 *QResizeEvent) {
	purego.SyscallN(addr_QSplitter_resizeEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QSplitter) ChangeEventDefault(arg1 *QEvent) {
	purego.SyscallN(addr_QSplitter_changeEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QSplitter) PaintEventDefault(arg1 *QPaintEvent) {
	purego.SyscallN(addr_QSplitter_paintEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QSplitter) InitStyleOptionDefault(option *QStyleOptionFrame) {
	purego.SyscallN(addr_QSplitter_initStyleOption_default, this.h, uintptr(option.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(option)
}

func (this *QSplitter) DevTypeDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QSplitter_devType_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QSplitter) SetVisibleDefault(visible bool) {
	purego.SyscallN(addr_QSplitter_setVisible_default, this.h, uintptr(b2i(visible)))
	runtime.KeepAlive(this)
}

func (this *QSplitter) HeightForWidthDefault(arg1 int) int {
	__r0, _, _ := purego.SyscallN(addr_QSplitter_heightForWidth_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QSplitter) HasHeightForWidthDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QSplitter_hasHeightForWidth_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QSplitter) PaintEngineDefault() *QPaintEngine {
	__r, _, _ := purego.SyscallN(addr_QSplitter_paintEngine_default, this.h)
	runtime.KeepAlive(this)
	return bindQPaintEngine(__r)
}

func (this *QSplitter) MousePressEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QSplitter_mousePressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSplitter) MouseReleaseEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QSplitter_mouseReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSplitter) MouseDoubleClickEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QSplitter_mouseDoubleClickEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSplitter) MouseMoveEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QSplitter_mouseMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSplitter) WheelEventDefault(event *QWheelEvent) {
	purego.SyscallN(addr_QSplitter_wheelEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSplitter) KeyPressEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QSplitter_keyPressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSplitter) KeyReleaseEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QSplitter_keyReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSplitter) FocusInEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QSplitter_focusInEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSplitter) FocusOutEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QSplitter_focusOutEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSplitter) EnterEventDefault(event *QEnterEvent) {
	purego.SyscallN(addr_QSplitter_enterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSplitter) LeaveEventDefault(event *QEvent) {
	purego.SyscallN(addr_QSplitter_leaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSplitter) MoveEventDefault(event *QMoveEvent) {
	purego.SyscallN(addr_QSplitter_moveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSplitter) CloseEventDefault(event *QCloseEvent) {
	purego.SyscallN(addr_QSplitter_closeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSplitter) ContextMenuEventDefault(event *QContextMenuEvent) {
	purego.SyscallN(addr_QSplitter_contextMenuEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSplitter) TabletEventDefault(event *QTabletEvent) {
	purego.SyscallN(addr_QSplitter_tabletEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSplitter) ActionEventDefault(event *QActionEvent) {
	purego.SyscallN(addr_QSplitter_actionEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSplitter) DragEnterEventDefault(event *QDragEnterEvent) {
	purego.SyscallN(addr_QSplitter_dragEnterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSplitter) DragMoveEventDefault(event *QDragMoveEvent) {
	purego.SyscallN(addr_QSplitter_dragMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSplitter) DragLeaveEventDefault(event *QDragLeaveEvent) {
	purego.SyscallN(addr_QSplitter_dragLeaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSplitter) DropEventDefault(event *QDropEvent) {
	purego.SyscallN(addr_QSplitter_dropEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSplitter) ShowEventDefault(event *QShowEvent) {
	purego.SyscallN(addr_QSplitter_showEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSplitter) HideEventDefault(event *QHideEvent) {
	purego.SyscallN(addr_QSplitter_hideEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSplitter) MetricDefault(arg1 QPaintDevice__PaintDeviceMetric) int {
	__r0, _, _ := purego.SyscallN(addr_QSplitter_metric_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QSplitter) InitPainterDefault(painter *QPainter) {
	purego.SyscallN(addr_QSplitter_initPainter_default, this.h, uintptr(painter.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
}

func (this *QSplitter) RedirectedDefault(offset *QPoint) *QPaintDevice {
	__r, _, _ := purego.SyscallN(addr_QSplitter_redirected_default, this.h, uintptr(offset.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(offset)
	return bindQPaintDevice(__r)
}

func (this *QSplitter) SharedPainterDefault() *QPainter {
	__r, _, _ := purego.SyscallN(addr_QSplitter_sharedPainter_default, this.h)
	runtime.KeepAlive(this)
	return bindQPainter(__r)
}

func (this *QSplitter) InputMethodEventDefault(arg1 *QInputMethodEvent) {
	purego.SyscallN(addr_QSplitter_inputMethodEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QSplitter) InputMethodQueryDefault(arg1 InputMethodQuery) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QSplitter_inputMethodQuery_default, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QSplitter) FocusNextPrevChildDefault(next bool) bool {
	__r0, _, _ := purego.SyscallN(addr_QSplitter_focusNextPrevChild_default, this.h, uintptr(b2i(next)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QSplitter) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QSplitter_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QSplitter) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QSplitter_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSplitter) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QSplitter_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSplitter) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QSplitter_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QSplitter) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QSplitter_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QSplitter) OnSplitterMoved(fn func(int, int)) int64 {
	id := registerCallback(func(pos int32, index int32) { fn(int(pos), int(index)) })
	pg_QSplitter_connect_splitterMoved(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSplitter) OnWindowTitleChanged(fn func(string)) int64 {
	id := registerCallback(func(title uintptr, title_len int64) { fn(pgStrNoFree(title, title_len)) })
	pg_QSplitter_connect_windowTitleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSplitter) OnWindowIconChanged(fn func(*QIcon)) int64 {
	id := registerCallback(func(icon uintptr) { fn(UnsafeNewQIcon(unsafe.Pointer(icon))) })
	pg_QSplitter_connect_windowIconChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSplitter) OnWindowIconTextChanged(fn func(string)) int64 {
	id := registerCallback(func(iconText uintptr, iconText_len int64) { fn(pgStrNoFree(iconText, iconText_len)) })
	pg_QSplitter_connect_windowIconTextChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSplitter) OnCustomContextMenuRequested(fn func(*QPoint)) int64 {
	id := registerCallback(func(pos uintptr) { fn(UnsafeNewQPoint(unsafe.Pointer(pos))) })
	pg_QSplitter_connect_customContextMenuRequested(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSplitter) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QSplitter_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSplitter) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QSplitter_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSplitter) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QSplitter_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSplitter) Disconnect(handle int64) {
	pg_QSplitter__disconnect(handle)
	unregisterCallback(handle)
}

func QSplitter_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case 729690261457684101:
		return 1
	case 7049264599837532491:
		return 2
	case 7399745205369955835:
		return 3
	case -1185315367526857833:
		return 4
	case 1237236456312824991:
		return 5
	case -4302917285655014661:
		return 6
	case 4959376063826165943:
		return 7
	case 5290442163167458465:
		return 8
	case -488693769771308531:
		return 9
	case -8659679893868524044:
		return 10
	case 2110182719375245575:
		return 11
	case 4311534060495737339:
		return 12
	case -8266762397995423915:
		return 13
	case 6398007404756346381:
		return 14
	case -332090898078543319:
		return 15
	case -4415920049297795213:
		return 16
	case -4029073905697023887:
		return 17
	case -4255494004099806581:
		return 18
	case 1919988340934511316:
		return 19
	case 1023791977679793185:
		return 20
	case 772598475950050299:
		return 21
	case 975043767033552552:
		return 22
	case -385596010576048993:
		return 23
	case 6316270346668237895:
		return 24
	case -551070009623711640:
		return 25
	case -9161949150489525396:
		return 26
	case -1989937050019464783:
		return 27
	case -6698043222227006015:
		return 28
	case -212068866860701141:
		return 29
	case 1697444140088555581:
		return 30
	case 206122715919917589:
		return 31
	case -1781545137352745978:
		return 32
	case -3139904158373725006:
		return 33
	case 9107058682754591716:
		return 34
	case -6825369174751916:
		return 35
	case -199718707876719129:
		return 36
	case -2039166586246443706:
		return 37
	case 427629649267311777:
		return 38
	case 6477500296980050210:
		return 39
	case 4127240633310655048:
		return 40
	case 5792771328741774991:
		return 41
	case 5438819661942843222:
		return 42
	case 7345853097340709638:
		return 43
	case 5898328517505784191:
		return 44
	case 5217623514630858297:
		return 45
	case -5409927581812374498:
		return 46
	case 3769868229123868474:
		return 47
	case 7379029982170503008:
		return 48
	case -568369772982215962:
		return 49
	}
	return -1
}

func QSplitter_enable(h uintptr, idx int32) { pg_QSplitter__set_override(h, idx) }

func (this *QSplitter) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QSplitter) OnSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 729690261457684101, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sizeHint"
}

func (this *QSplitter) OnMinimumSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 7049264599837532491, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "minimumSizeHint"
}

func (this *QSplitter) OnCreateHandle(fn func() *QSplitterHandle) {
	registerVirt(this.UnsafeGoID(), 7399745205369955835, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "createHandle"
}

func (this *QSplitter) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(arg1 uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(arg1))) }) // slot "childEvent"
}

func (this *QSplitter) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(arg1 uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(arg1)))) }) // slot "event"
}

func (this *QSplitter) OnResizeEvent(fn func(*QResizeEvent)) {
	registerVirt(this.UnsafeGoID(), -4302917285655014661, func(arg1 uintptr) { fn(UnsafeNewQResizeEvent(unsafe.Pointer(arg1))) }) // slot "resizeEvent"
}

func (this *QSplitter) OnChangeEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 4959376063826165943, func(arg1 uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(arg1))) }) // slot "changeEvent"
}

func (this *QSplitter) OnPaintEvent(fn func(*QPaintEvent)) {
	registerVirt(this.UnsafeGoID(), 5290442163167458465, func(arg1 uintptr) { fn(UnsafeNewQPaintEvent(unsafe.Pointer(arg1))) }) // slot "paintEvent"
}

func (this *QSplitter) OnInitStyleOption(fn func(*QStyleOptionFrame)) {
	registerVirt(this.UnsafeGoID(), -488693769771308531, func(option uintptr) { fn(UnsafeNewQStyleOptionFrame(unsafe.Pointer(option))) }) // slot "initStyleOption"
}

func (this *QSplitter) OnDevType(fn func() int) {
	registerVirt(this.UnsafeGoID(), -8659679893868524044, func() int32 { return int32(fn()) }) // slot "devType"
}

func (this *QSplitter) OnSetVisible(fn func(bool)) {
	registerVirt(this.UnsafeGoID(), 2110182719375245575, func(visible int32) { fn(visible != 0) }) // slot "setVisible"
}

func (this *QSplitter) OnHeightForWidth(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), 4311534060495737339, func(arg1 int32) int32 { return int32(fn(int(arg1))) }) // slot "heightForWidth"
}

func (this *QSplitter) OnHasHeightForWidth(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -8266762397995423915, func() int32 { return b2i(fn()) }) // slot "hasHeightForWidth"
}

func (this *QSplitter) OnPaintEngine(fn func() *QPaintEngine) {
	registerVirt(this.UnsafeGoID(), 6398007404756346381, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "paintEngine"
}

func (this *QSplitter) OnMousePressEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -332090898078543319, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mousePressEvent"
}

func (this *QSplitter) OnMouseReleaseEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4415920049297795213, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseReleaseEvent"
}

func (this *QSplitter) OnMouseDoubleClickEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4029073905697023887, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseDoubleClickEvent"
}

func (this *QSplitter) OnMouseMoveEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4255494004099806581, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseMoveEvent"
}

func (this *QSplitter) OnWheelEvent(fn func(*QWheelEvent)) {
	registerVirt(this.UnsafeGoID(), 1919988340934511316, func(event uintptr) { fn(UnsafeNewQWheelEvent(unsafe.Pointer(event))) }) // slot "wheelEvent"
}

func (this *QSplitter) OnKeyPressEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 1023791977679793185, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyPressEvent"
}

func (this *QSplitter) OnKeyReleaseEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 772598475950050299, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyReleaseEvent"
}

func (this *QSplitter) OnFocusInEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), 975043767033552552, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusInEvent"
}

func (this *QSplitter) OnFocusOutEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), -385596010576048993, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusOutEvent"
}

func (this *QSplitter) OnEnterEvent(fn func(*QEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 6316270346668237895, func(event uintptr) { fn(UnsafeNewQEnterEvent(unsafe.Pointer(event))) }) // slot "enterEvent"
}

func (this *QSplitter) OnLeaveEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), -551070009623711640, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "leaveEvent"
}

func (this *QSplitter) OnMoveEvent(fn func(*QMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -9161949150489525396, func(event uintptr) { fn(UnsafeNewQMoveEvent(unsafe.Pointer(event))) }) // slot "moveEvent"
}

func (this *QSplitter) OnCloseEvent(fn func(*QCloseEvent)) {
	registerVirt(this.UnsafeGoID(), -1989937050019464783, func(event uintptr) { fn(UnsafeNewQCloseEvent(unsafe.Pointer(event))) }) // slot "closeEvent"
}

func (this *QSplitter) OnContextMenuEvent(fn func(*QContextMenuEvent)) {
	registerVirt(this.UnsafeGoID(), -6698043222227006015, func(event uintptr) { fn(UnsafeNewQContextMenuEvent(unsafe.Pointer(event))) }) // slot "contextMenuEvent"
}

func (this *QSplitter) OnTabletEvent(fn func(*QTabletEvent)) {
	registerVirt(this.UnsafeGoID(), -212068866860701141, func(event uintptr) { fn(UnsafeNewQTabletEvent(unsafe.Pointer(event))) }) // slot "tabletEvent"
}

func (this *QSplitter) OnActionEvent(fn func(*QActionEvent)) {
	registerVirt(this.UnsafeGoID(), 1697444140088555581, func(event uintptr) { fn(UnsafeNewQActionEvent(unsafe.Pointer(event))) }) // slot "actionEvent"
}

func (this *QSplitter) OnDragEnterEvent(fn func(*QDragEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 206122715919917589, func(event uintptr) { fn(UnsafeNewQDragEnterEvent(unsafe.Pointer(event))) }) // slot "dragEnterEvent"
}

func (this *QSplitter) OnDragMoveEvent(fn func(*QDragMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -1781545137352745978, func(event uintptr) { fn(UnsafeNewQDragMoveEvent(unsafe.Pointer(event))) }) // slot "dragMoveEvent"
}

func (this *QSplitter) OnDragLeaveEvent(fn func(*QDragLeaveEvent)) {
	registerVirt(this.UnsafeGoID(), -3139904158373725006, func(event uintptr) { fn(UnsafeNewQDragLeaveEvent(unsafe.Pointer(event))) }) // slot "dragLeaveEvent"
}

func (this *QSplitter) OnDropEvent(fn func(*QDropEvent)) {
	registerVirt(this.UnsafeGoID(), 9107058682754591716, func(event uintptr) { fn(UnsafeNewQDropEvent(unsafe.Pointer(event))) }) // slot "dropEvent"
}

func (this *QSplitter) OnShowEvent(fn func(*QShowEvent)) {
	registerVirt(this.UnsafeGoID(), -6825369174751916, func(event uintptr) { fn(UnsafeNewQShowEvent(unsafe.Pointer(event))) }) // slot "showEvent"
}

func (this *QSplitter) OnHideEvent(fn func(*QHideEvent)) {
	registerVirt(this.UnsafeGoID(), -199718707876719129, func(event uintptr) { fn(UnsafeNewQHideEvent(unsafe.Pointer(event))) }) // slot "hideEvent"
}

func (this *QSplitter) OnNativeEvent(fn func(*QByteArray, unsafe.Pointer, *int64) bool) {
	registerVirt(this.UnsafeGoID(), -2039166586246443706, func(eventType uintptr, message uintptr, result uintptr) int32 { return b2i(fn(UnsafeNewQByteArray(unsafe.Pointer(eventType)), unsafe.Pointer(message), (*int64)(unsafe.Pointer(result)))) }) // slot "nativeEvent"
}

func (this *QSplitter) OnMetric(fn func(QPaintDevice__PaintDeviceMetric) int) {
	registerVirt(this.UnsafeGoID(), 427629649267311777, func(arg1 int32) int32 { return int32(fn(QPaintDevice__PaintDeviceMetric(arg1))) }) // slot "metric"
}

func (this *QSplitter) OnInitPainter(fn func(*QPainter)) {
	registerVirt(this.UnsafeGoID(), 6477500296980050210, func(painter uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter))) }) // slot "initPainter"
}

func (this *QSplitter) OnRedirected(fn func(*QPoint) *QPaintDevice) {
	registerVirt(this.UnsafeGoID(), 4127240633310655048, func(offset uintptr) uintptr { return uintptr(fn(UnsafeNewQPoint(unsafe.Pointer(offset))).UnsafePointer()) }) // slot "redirected"
}

func (this *QSplitter) OnSharedPainter(fn func() *QPainter) {
	registerVirt(this.UnsafeGoID(), 5792771328741774991, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sharedPainter"
}

func (this *QSplitter) OnInputMethodEvent(fn func(*QInputMethodEvent)) {
	registerVirt(this.UnsafeGoID(), 5438819661942843222, func(arg1 uintptr) { fn(UnsafeNewQInputMethodEvent(unsafe.Pointer(arg1))) }) // slot "inputMethodEvent"
}

func (this *QSplitter) OnInputMethodQuery(fn func(InputMethodQuery) *QVariant) {
	registerVirt(this.UnsafeGoID(), 7345853097340709638, func(arg1 int32) uintptr { return uintptr(fn(InputMethodQuery(arg1)).UnsafePointer()) }) // slot "inputMethodQuery"
}

func (this *QSplitter) OnFocusNextPrevChild(fn func(bool) bool) {
	registerVirt(this.UnsafeGoID(), 5898328517505784191, func(next int32) int32 { return b2i(fn(next != 0)) }) // slot "focusNextPrevChild"
}

func (this *QSplitter) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QSplitter) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QSplitter) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QSplitter) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QSplitter) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QSplitter) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QSplitter_delete(this.h)
	}
	this.UnsafeRelease()
}
