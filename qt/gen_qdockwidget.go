package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QDockWidget_chain uintptr
	addr_QDockWidget_new uintptr
	addr_QDockWidget_new2 uintptr
	addr_QDockWidget_new3 uintptr
	addr_QDockWidget_new4 uintptr
	addr_QDockWidget_new5 uintptr
	addr_QDockWidget_new6 uintptr
	addr_QDockWidget_tr uintptr
	addr_QDockWidget_widget uintptr
	addr_QDockWidget_setWidget uintptr
	addr_QDockWidget_setFeatures uintptr
	addr_QDockWidget_features uintptr
	addr_QDockWidget_setFloating uintptr
	addr_QDockWidget_isFloating uintptr
	addr_QDockWidget_setAllowedAreas uintptr
	addr_QDockWidget_allowedAreas uintptr
	addr_QDockWidget_setTitleBarWidget uintptr
	addr_QDockWidget_titleBarWidget uintptr
	addr_QDockWidget_setDockLocation uintptr
	addr_QDockWidget_dockLocation uintptr
	addr_QDockWidget_isAreaAllowed uintptr
	addr_QDockWidget_toggleViewAction uintptr
	addr_QDockWidget_featuresChanged uintptr
	addr_QDockWidget_topLevelChanged uintptr
	addr_QDockWidget_allowedAreasChanged uintptr
	addr_QDockWidget_visibilityChanged uintptr
	addr_QDockWidget_dockLocationChanged uintptr
	addr_QDockWidget_setTabOrder uintptr
	addr_QDockWidget_mouseGrabber uintptr
	addr_QDockWidget_keyboardGrabber uintptr
	addr_QDockWidget_find uintptr
	addr_QDockWidget_createWindowContainer uintptr
	addr_QDockWidget_updateMicroFocus uintptr
	addr_QDockWidget_create uintptr
	addr_QDockWidget_destroy uintptr
	addr_QDockWidget_focusNextChild uintptr
	addr_QDockWidget_focusPreviousChild uintptr
	addr_QDockWidget_connect uintptr
	addr_QDockWidget_connect2 uintptr
	addr_QDockWidget_sender uintptr
	addr_QDockWidget_senderSignalIndex uintptr
	addr_QDockWidget_receivers uintptr
	addr_QDockWidget_isSignalConnected uintptr
	pg_QDockWidget_devicePixelRatioFScale func() float64
	pg_QDockWidget_encodeMetricF func(int32, float64) int32
	pg_QDockWidget_getDecodedMetricF func(uintptr, int32, int32) float64
	addr_QDockWidget_tr2 uintptr
	addr_QDockWidget_tr3 uintptr
	addr_QDockWidget_createWindowContainer2 uintptr
	addr_QDockWidget_createWindowContainer3 uintptr
	addr_QDockWidget_updateMicroFocus2 uintptr
	addr_QDockWidget_create2 uintptr
	addr_QDockWidget_create3 uintptr
	addr_QDockWidget_create4 uintptr
	addr_QDockWidget_destroy2 uintptr
	addr_QDockWidget_destroy3 uintptr
	addr_QDockWidget_connect3 uintptr
	addr_QDockWidget_connect4 uintptr
	addr_QDockWidget_metaObject_default uintptr
	addr_QDockWidget_changeEvent_default uintptr
	addr_QDockWidget_closeEvent_default uintptr
	addr_QDockWidget_paintEvent_default uintptr
	addr_QDockWidget_event_default uintptr
	addr_QDockWidget_initStyleOption_default uintptr
	addr_QDockWidget_devType_default uintptr
	addr_QDockWidget_setVisible_default uintptr
	addr_QDockWidget_sizeHint_default uintptr
	addr_QDockWidget_minimumSizeHint_default uintptr
	addr_QDockWidget_heightForWidth_default uintptr
	addr_QDockWidget_hasHeightForWidth_default uintptr
	addr_QDockWidget_paintEngine_default uintptr
	addr_QDockWidget_mousePressEvent_default uintptr
	addr_QDockWidget_mouseReleaseEvent_default uintptr
	addr_QDockWidget_mouseDoubleClickEvent_default uintptr
	addr_QDockWidget_mouseMoveEvent_default uintptr
	addr_QDockWidget_wheelEvent_default uintptr
	addr_QDockWidget_keyPressEvent_default uintptr
	addr_QDockWidget_keyReleaseEvent_default uintptr
	addr_QDockWidget_focusInEvent_default uintptr
	addr_QDockWidget_focusOutEvent_default uintptr
	addr_QDockWidget_enterEvent_default uintptr
	addr_QDockWidget_leaveEvent_default uintptr
	addr_QDockWidget_moveEvent_default uintptr
	addr_QDockWidget_resizeEvent_default uintptr
	addr_QDockWidget_contextMenuEvent_default uintptr
	addr_QDockWidget_tabletEvent_default uintptr
	addr_QDockWidget_actionEvent_default uintptr
	addr_QDockWidget_dragEnterEvent_default uintptr
	addr_QDockWidget_dragMoveEvent_default uintptr
	addr_QDockWidget_dragLeaveEvent_default uintptr
	addr_QDockWidget_dropEvent_default uintptr
	addr_QDockWidget_showEvent_default uintptr
	addr_QDockWidget_hideEvent_default uintptr
	addr_QDockWidget_metric_default uintptr
	addr_QDockWidget_initPainter_default uintptr
	addr_QDockWidget_redirected_default uintptr
	addr_QDockWidget_sharedPainter_default uintptr
	addr_QDockWidget_inputMethodEvent_default uintptr
	addr_QDockWidget_inputMethodQuery_default uintptr
	addr_QDockWidget_focusNextPrevChild_default uintptr
	addr_QDockWidget_eventFilter_default uintptr
	addr_QDockWidget_timerEvent_default uintptr
	addr_QDockWidget_childEvent_default uintptr
	addr_QDockWidget_customEvent_default uintptr
	addr_QDockWidget_connectNotify_default uintptr
	addr_QDockWidget_disconnectNotify_default uintptr
	pg_QDockWidget_connect_featuresChanged func(uintptr, int64)
	pg_QDockWidget_connect_topLevelChanged func(uintptr, int64)
	pg_QDockWidget_connect_allowedAreasChanged func(uintptr, int64)
	pg_QDockWidget_connect_visibilityChanged func(uintptr, int64)
	pg_QDockWidget_connect_dockLocationChanged func(uintptr, int64)
	pg_QDockWidget_connect_windowTitleChanged func(uintptr, int64)
	pg_QDockWidget_connect_windowIconChanged func(uintptr, int64)
	pg_QDockWidget_connect_windowIconTextChanged func(uintptr, int64)
	pg_QDockWidget_connect_customContextMenuRequested func(uintptr, int64)
	pg_QDockWidget_connect_destroyed func(uintptr, int64)
	pg_QDockWidget_connect_destroyed2 func(uintptr, int64)
	pg_QDockWidget_connect_objectNameChanged func(uintptr, int64)
	pg_QDockWidget__disconnect func(int64)
	pg_QDockWidget__set_override func(uintptr, int32)
	pg_QDockWidget_delete func(uintptr)
)

func init() {
	symaddr(&addr_QDockWidget_chain, "QDockWidget_chain")
	symaddr(&addr_QDockWidget_new, "QDockWidget_new")
	symaddr(&addr_QDockWidget_new2, "QDockWidget_new2")
	symaddr(&addr_QDockWidget_new3, "QDockWidget_new3")
	symaddr(&addr_QDockWidget_new4, "QDockWidget_new4")
	symaddr(&addr_QDockWidget_new5, "QDockWidget_new5")
	symaddr(&addr_QDockWidget_new6, "QDockWidget_new6")
	symaddr(&addr_QDockWidget_tr, "QDockWidget_tr")
	symaddr(&addr_QDockWidget_widget, "QDockWidget_widget")
	symaddr(&addr_QDockWidget_setWidget, "QDockWidget_setWidget")
	symaddr(&addr_QDockWidget_setFeatures, "QDockWidget_setFeatures")
	symaddr(&addr_QDockWidget_features, "QDockWidget_features")
	symaddr(&addr_QDockWidget_setFloating, "QDockWidget_setFloating")
	symaddr(&addr_QDockWidget_isFloating, "QDockWidget_isFloating")
	symaddr(&addr_QDockWidget_setAllowedAreas, "QDockWidget_setAllowedAreas")
	symaddr(&addr_QDockWidget_allowedAreas, "QDockWidget_allowedAreas")
	symaddr(&addr_QDockWidget_setTitleBarWidget, "QDockWidget_setTitleBarWidget")
	symaddr(&addr_QDockWidget_titleBarWidget, "QDockWidget_titleBarWidget")
	symaddr(&addr_QDockWidget_setDockLocation, "QDockWidget_setDockLocation")
	symaddr(&addr_QDockWidget_dockLocation, "QDockWidget_dockLocation")
	symaddr(&addr_QDockWidget_isAreaAllowed, "QDockWidget_isAreaAllowed")
	symaddr(&addr_QDockWidget_toggleViewAction, "QDockWidget_toggleViewAction")
	symaddr(&addr_QDockWidget_featuresChanged, "QDockWidget_featuresChanged")
	symaddr(&addr_QDockWidget_topLevelChanged, "QDockWidget_topLevelChanged")
	symaddr(&addr_QDockWidget_allowedAreasChanged, "QDockWidget_allowedAreasChanged")
	symaddr(&addr_QDockWidget_visibilityChanged, "QDockWidget_visibilityChanged")
	symaddr(&addr_QDockWidget_dockLocationChanged, "QDockWidget_dockLocationChanged")
	symaddr(&addr_QDockWidget_setTabOrder, "QDockWidget_setTabOrder")
	symaddr(&addr_QDockWidget_mouseGrabber, "QDockWidget_mouseGrabber")
	symaddr(&addr_QDockWidget_keyboardGrabber, "QDockWidget_keyboardGrabber")
	symaddr(&addr_QDockWidget_find, "QDockWidget_find")
	symaddr(&addr_QDockWidget_createWindowContainer, "QDockWidget_createWindowContainer")
	symaddr(&addr_QDockWidget_updateMicroFocus, "QDockWidget_updateMicroFocus")
	symaddr(&addr_QDockWidget_create, "QDockWidget_create")
	symaddr(&addr_QDockWidget_destroy, "QDockWidget_destroy")
	symaddr(&addr_QDockWidget_focusNextChild, "QDockWidget_focusNextChild")
	symaddr(&addr_QDockWidget_focusPreviousChild, "QDockWidget_focusPreviousChild")
	symaddr(&addr_QDockWidget_connect, "QDockWidget_connect")
	symaddr(&addr_QDockWidget_connect2, "QDockWidget_connect2")
	symaddr(&addr_QDockWidget_sender, "QDockWidget_sender")
	symaddr(&addr_QDockWidget_senderSignalIndex, "QDockWidget_senderSignalIndex")
	symaddr(&addr_QDockWidget_receivers, "QDockWidget_receivers")
	symaddr(&addr_QDockWidget_isSignalConnected, "QDockWidget_isSignalConnected")
	reg(&pg_QDockWidget_devicePixelRatioFScale, "QDockWidget_devicePixelRatioFScale")
	reg(&pg_QDockWidget_encodeMetricF, "QDockWidget_encodeMetricF")
	reg(&pg_QDockWidget_getDecodedMetricF, "QDockWidget_getDecodedMetricF")
	symaddr(&addr_QDockWidget_tr2, "QDockWidget_tr2")
	symaddr(&addr_QDockWidget_tr3, "QDockWidget_tr3")
	symaddr(&addr_QDockWidget_createWindowContainer2, "QDockWidget_createWindowContainer2")
	symaddr(&addr_QDockWidget_createWindowContainer3, "QDockWidget_createWindowContainer3")
	symaddr(&addr_QDockWidget_updateMicroFocus2, "QDockWidget_updateMicroFocus2")
	symaddr(&addr_QDockWidget_create2, "QDockWidget_create2")
	symaddr(&addr_QDockWidget_create3, "QDockWidget_create3")
	symaddr(&addr_QDockWidget_create4, "QDockWidget_create4")
	symaddr(&addr_QDockWidget_destroy2, "QDockWidget_destroy2")
	symaddr(&addr_QDockWidget_destroy3, "QDockWidget_destroy3")
	symaddr(&addr_QDockWidget_connect3, "QDockWidget_connect3")
	symaddr(&addr_QDockWidget_connect4, "QDockWidget_connect4")
	symaddr(&addr_QDockWidget_metaObject_default, "QDockWidget_metaObject_default")
	symaddr(&addr_QDockWidget_changeEvent_default, "QDockWidget_changeEvent_default")
	symaddr(&addr_QDockWidget_closeEvent_default, "QDockWidget_closeEvent_default")
	symaddr(&addr_QDockWidget_paintEvent_default, "QDockWidget_paintEvent_default")
	symaddr(&addr_QDockWidget_event_default, "QDockWidget_event_default")
	symaddr(&addr_QDockWidget_initStyleOption_default, "QDockWidget_initStyleOption_default")
	symaddr(&addr_QDockWidget_devType_default, "QDockWidget_devType_default")
	symaddr(&addr_QDockWidget_setVisible_default, "QDockWidget_setVisible_default")
	symaddr(&addr_QDockWidget_sizeHint_default, "QDockWidget_sizeHint_default")
	symaddr(&addr_QDockWidget_minimumSizeHint_default, "QDockWidget_minimumSizeHint_default")
	symaddr(&addr_QDockWidget_heightForWidth_default, "QDockWidget_heightForWidth_default")
	symaddr(&addr_QDockWidget_hasHeightForWidth_default, "QDockWidget_hasHeightForWidth_default")
	symaddr(&addr_QDockWidget_paintEngine_default, "QDockWidget_paintEngine_default")
	symaddr(&addr_QDockWidget_mousePressEvent_default, "QDockWidget_mousePressEvent_default")
	symaddr(&addr_QDockWidget_mouseReleaseEvent_default, "QDockWidget_mouseReleaseEvent_default")
	symaddr(&addr_QDockWidget_mouseDoubleClickEvent_default, "QDockWidget_mouseDoubleClickEvent_default")
	symaddr(&addr_QDockWidget_mouseMoveEvent_default, "QDockWidget_mouseMoveEvent_default")
	symaddr(&addr_QDockWidget_wheelEvent_default, "QDockWidget_wheelEvent_default")
	symaddr(&addr_QDockWidget_keyPressEvent_default, "QDockWidget_keyPressEvent_default")
	symaddr(&addr_QDockWidget_keyReleaseEvent_default, "QDockWidget_keyReleaseEvent_default")
	symaddr(&addr_QDockWidget_focusInEvent_default, "QDockWidget_focusInEvent_default")
	symaddr(&addr_QDockWidget_focusOutEvent_default, "QDockWidget_focusOutEvent_default")
	symaddr(&addr_QDockWidget_enterEvent_default, "QDockWidget_enterEvent_default")
	symaddr(&addr_QDockWidget_leaveEvent_default, "QDockWidget_leaveEvent_default")
	symaddr(&addr_QDockWidget_moveEvent_default, "QDockWidget_moveEvent_default")
	symaddr(&addr_QDockWidget_resizeEvent_default, "QDockWidget_resizeEvent_default")
	symaddr(&addr_QDockWidget_contextMenuEvent_default, "QDockWidget_contextMenuEvent_default")
	symaddr(&addr_QDockWidget_tabletEvent_default, "QDockWidget_tabletEvent_default")
	symaddr(&addr_QDockWidget_actionEvent_default, "QDockWidget_actionEvent_default")
	symaddr(&addr_QDockWidget_dragEnterEvent_default, "QDockWidget_dragEnterEvent_default")
	symaddr(&addr_QDockWidget_dragMoveEvent_default, "QDockWidget_dragMoveEvent_default")
	symaddr(&addr_QDockWidget_dragLeaveEvent_default, "QDockWidget_dragLeaveEvent_default")
	symaddr(&addr_QDockWidget_dropEvent_default, "QDockWidget_dropEvent_default")
	symaddr(&addr_QDockWidget_showEvent_default, "QDockWidget_showEvent_default")
	symaddr(&addr_QDockWidget_hideEvent_default, "QDockWidget_hideEvent_default")
	symaddr(&addr_QDockWidget_metric_default, "QDockWidget_metric_default")
	symaddr(&addr_QDockWidget_initPainter_default, "QDockWidget_initPainter_default")
	symaddr(&addr_QDockWidget_redirected_default, "QDockWidget_redirected_default")
	symaddr(&addr_QDockWidget_sharedPainter_default, "QDockWidget_sharedPainter_default")
	symaddr(&addr_QDockWidget_inputMethodEvent_default, "QDockWidget_inputMethodEvent_default")
	symaddr(&addr_QDockWidget_inputMethodQuery_default, "QDockWidget_inputMethodQuery_default")
	symaddr(&addr_QDockWidget_focusNextPrevChild_default, "QDockWidget_focusNextPrevChild_default")
	symaddr(&addr_QDockWidget_eventFilter_default, "QDockWidget_eventFilter_default")
	symaddr(&addr_QDockWidget_timerEvent_default, "QDockWidget_timerEvent_default")
	symaddr(&addr_QDockWidget_childEvent_default, "QDockWidget_childEvent_default")
	symaddr(&addr_QDockWidget_customEvent_default, "QDockWidget_customEvent_default")
	symaddr(&addr_QDockWidget_connectNotify_default, "QDockWidget_connectNotify_default")
	symaddr(&addr_QDockWidget_disconnectNotify_default, "QDockWidget_disconnectNotify_default")
	reg(&pg_QDockWidget_connect_featuresChanged, "QDockWidget_connect_featuresChanged")
	reg(&pg_QDockWidget_connect_topLevelChanged, "QDockWidget_connect_topLevelChanged")
	reg(&pg_QDockWidget_connect_allowedAreasChanged, "QDockWidget_connect_allowedAreasChanged")
	reg(&pg_QDockWidget_connect_visibilityChanged, "QDockWidget_connect_visibilityChanged")
	reg(&pg_QDockWidget_connect_dockLocationChanged, "QDockWidget_connect_dockLocationChanged")
	reg(&pg_QDockWidget_connect_windowTitleChanged, "QDockWidget_connect_windowTitleChanged")
	reg(&pg_QDockWidget_connect_windowIconChanged, "QDockWidget_connect_windowIconChanged")
	reg(&pg_QDockWidget_connect_windowIconTextChanged, "QDockWidget_connect_windowIconTextChanged")
	reg(&pg_QDockWidget_connect_customContextMenuRequested, "QDockWidget_connect_customContextMenuRequested")
	reg(&pg_QDockWidget_connect_destroyed, "QDockWidget_connect_destroyed")
	reg(&pg_QDockWidget_connect_destroyed2, "QDockWidget_connect_destroyed2")
	reg(&pg_QDockWidget_connect_objectNameChanged, "QDockWidget_connect_objectNameChanged")
	reg(&pg_QDockWidget__disconnect, "QDockWidget__disconnect")
	reg(&pg_QDockWidget__set_override, "QDockWidget__set_override")
	reg(&pg_QDockWidget_delete, "QDockWidget_delete")
}

type QDockWidget struct {
	h uintptr
	*QWidget
}

func bindQDockWidget(h uintptr) *QDockWidget {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QDockWidget_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QDockWidget{h: h, QWidget: &QWidget{h: __c[0], QObject: &QObject{h: __c[1]}}}
}

func newownedQDockWidget(h uintptr) *QDockWidget { return bindQDockWidget(h) }

func (this *QDockWidget) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QDockWidget) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQDockWidget(h unsafe.Pointer) *QDockWidget { return bindQDockWidget(uintptr(h)) }
func UnsafeNewQDockWidgetOwned(h unsafe.Pointer) *QDockWidget { return newownedQDockWidget(uintptr(h)) }

func (this *QDockWidget) UnsafeRelease() {
	this.h = 0
	if this.QWidget != nil {
		this.QWidget.UnsafeRelease()
	}
}

func NewQDockWidget(parent *QWidget) *QDockWidget {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QDockWidget_new, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQDockWidget(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QDockWidget_virtIdx, QDockWidget_enable)
	return __o
}

func NewQDockWidget2(title string) *QDockWidget {
	__goid := nextVirtID()
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	__r, _, _ := purego.SyscallN(addr_QDockWidget_new2, uintptr(__goid), title_p, uintptr(len(title)))
	__o := newownedQDockWidget(__r)
	runtime.KeepAlive(title)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QDockWidget_virtIdx, QDockWidget_enable)
	return __o
}

func NewQDockWidget3() *QDockWidget {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QDockWidget_new3, uintptr(__goid))
	__o := newownedQDockWidget(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QDockWidget_virtIdx, QDockWidget_enable)
	return __o
}

func NewQDockWidget4(title string, parent *QWidget) *QDockWidget {
	__goid := nextVirtID()
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	__r, _, _ := purego.SyscallN(addr_QDockWidget_new4, uintptr(__goid), title_p, uintptr(len(title)), uintptr(parent.UnsafePointer()))
	__o := newownedQDockWidget(__r)
	runtime.KeepAlive(title)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QDockWidget_virtIdx, QDockWidget_enable)
	return __o
}

func NewQDockWidget5(title string, parent *QWidget, flags WindowType) *QDockWidget {
	__goid := nextVirtID()
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	__r, _, _ := purego.SyscallN(addr_QDockWidget_new5, uintptr(__goid), title_p, uintptr(len(title)), uintptr(parent.UnsafePointer()), uintptr(flags))
	__o := newownedQDockWidget(__r)
	runtime.KeepAlive(title)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QDockWidget_virtIdx, QDockWidget_enable)
	return __o
}

func NewQDockWidget6(parent *QWidget, flags WindowType) *QDockWidget {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QDockWidget_new6, uintptr(__goid), uintptr(parent.UnsafePointer()), uintptr(flags))
	__o := newownedQDockWidget(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QDockWidget_virtIdx, QDockWidget_enable)
	return __o
}

func QDockWidget_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDockWidget_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QDockWidget) Widget() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QDockWidget_widget, this.h)
	runtime.KeepAlive(this)
	return bindQWidget(__r)
}

func (this *QDockWidget) SetWidget(widget *QWidget) {
	purego.SyscallN(addr_QDockWidget_setWidget, this.h, uintptr(widget.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(widget)
}

func (this *QDockWidget) SetFeatures(features QDockWidget__DockWidgetFeature) {
	purego.SyscallN(addr_QDockWidget_setFeatures, this.h, uintptr(features))
	runtime.KeepAlive(this)
}

func (this *QDockWidget) Features() QDockWidget__DockWidgetFeature {
	__r0, _, _ := purego.SyscallN(addr_QDockWidget_features, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QDockWidget__DockWidgetFeature(__r)
}

func (this *QDockWidget) SetFloating(floating bool) {
	purego.SyscallN(addr_QDockWidget_setFloating, this.h, uintptr(b2i(floating)))
	runtime.KeepAlive(this)
}

func (this *QDockWidget) IsFloating() bool {
	__r0, _, _ := purego.SyscallN(addr_QDockWidget_isFloating, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QDockWidget) SetAllowedAreas(areas DockWidgetArea) {
	purego.SyscallN(addr_QDockWidget_setAllowedAreas, this.h, uintptr(areas))
	runtime.KeepAlive(this)
}

func (this *QDockWidget) AllowedAreas() DockWidgetArea {
	__r0, _, _ := purego.SyscallN(addr_QDockWidget_allowedAreas, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return DockWidgetArea(__r)
}

func (this *QDockWidget) SetTitleBarWidget(widget *QWidget) {
	purego.SyscallN(addr_QDockWidget_setTitleBarWidget, this.h, uintptr(widget.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(widget)
}

func (this *QDockWidget) TitleBarWidget() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QDockWidget_titleBarWidget, this.h)
	runtime.KeepAlive(this)
	return bindQWidget(__r)
}

func (this *QDockWidget) SetDockLocation(area DockWidgetArea) {
	purego.SyscallN(addr_QDockWidget_setDockLocation, this.h, uintptr(area))
	runtime.KeepAlive(this)
}

func (this *QDockWidget) DockLocation() DockWidgetArea {
	__r0, _, _ := purego.SyscallN(addr_QDockWidget_dockLocation, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return DockWidgetArea(__r)
}

func (this *QDockWidget) IsAreaAllowed(area DockWidgetArea) bool {
	__r0, _, _ := purego.SyscallN(addr_QDockWidget_isAreaAllowed, this.h, uintptr(area))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QDockWidget) ToggleViewAction() *QAction {
	__r, _, _ := purego.SyscallN(addr_QDockWidget_toggleViewAction, this.h)
	runtime.KeepAlive(this)
	return bindQAction(__r)
}

func (this *QDockWidget) FeaturesChanged(features QDockWidget__DockWidgetFeature) {
	purego.SyscallN(addr_QDockWidget_featuresChanged, this.h, uintptr(features))
	runtime.KeepAlive(this)
}

func (this *QDockWidget) TopLevelChanged(topLevel bool) {
	purego.SyscallN(addr_QDockWidget_topLevelChanged, this.h, uintptr(b2i(topLevel)))
	runtime.KeepAlive(this)
}

func (this *QDockWidget) AllowedAreasChanged(allowedAreas DockWidgetArea) {
	purego.SyscallN(addr_QDockWidget_allowedAreasChanged, this.h, uintptr(allowedAreas))
	runtime.KeepAlive(this)
}

func (this *QDockWidget) VisibilityChanged(visible bool) {
	purego.SyscallN(addr_QDockWidget_visibilityChanged, this.h, uintptr(b2i(visible)))
	runtime.KeepAlive(this)
}

func (this *QDockWidget) DockLocationChanged(area DockWidgetArea) {
	purego.SyscallN(addr_QDockWidget_dockLocationChanged, this.h, uintptr(area))
	runtime.KeepAlive(this)
}

func QDockWidget_SetTabOrder(arg1 *QWidget, arg2 *QWidget) {
	purego.SyscallN(addr_QDockWidget_setTabOrder, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()))
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
}

func QDockWidget_MouseGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QDockWidget_mouseGrabber)
	return bindQWidget(__r)
}

func QDockWidget_KeyboardGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QDockWidget_keyboardGrabber)
	return bindQWidget(__r)
}

func QDockWidget_Find(arg1 uint64) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QDockWidget_find, uintptr(arg1))
	return bindQWidget(__r)
}

func QDockWidget_CreateWindowContainer(window *QWindow) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QDockWidget_createWindowContainer, uintptr(window.UnsafePointer()))
	runtime.KeepAlive(window)
	return bindQWidget(__r)
}

func (this *QDockWidget) UpdateMicroFocus() {
	purego.SyscallN(addr_QDockWidget_updateMicroFocus, this.h)
	runtime.KeepAlive(this)
}

func (this *QDockWidget) Create() {
	purego.SyscallN(addr_QDockWidget_create, this.h)
	runtime.KeepAlive(this)
}

func (this *QDockWidget) Destroy() {
	purego.SyscallN(addr_QDockWidget_destroy, this.h)
	runtime.KeepAlive(this)
}

func (this *QDockWidget) FocusNextChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QDockWidget_focusNextChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QDockWidget) FocusPreviousChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QDockWidget_focusPreviousChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QDockWidget_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QDockWidget_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QDockWidget_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QDockWidget_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QDockWidget) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QDockWidget_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QDockWidget) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QDockWidget_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QDockWidget) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QDockWidget_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QDockWidget) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QDockWidget_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QDockWidget_DevicePixelRatioFScale() float64 {
	__r := pg_QDockWidget_devicePixelRatioFScale()
	return __r
}

func QDockWidget_EncodeMetricF(metric QPaintDevice__PaintDeviceMetric, value float64) int {
	__r := pg_QDockWidget_encodeMetricF(int32(metric), value)
	return int(__r)
}

func (this *QDockWidget) GetDecodedMetricF(metricA QPaintDevice__PaintDeviceMetric, metricB QPaintDevice__PaintDeviceMetric) float64 {
	__r := pg_QDockWidget_getDecodedMetricF(this.h, int32(metricA), int32(metricB))
	runtime.KeepAlive(this)
	return __r
}

func QDockWidget_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDockWidget_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QDockWidget_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDockWidget_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QDockWidget_CreateWindowContainer2(window *QWindow, parent *QWidget) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QDockWidget_createWindowContainer2, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func QDockWidget_CreateWindowContainer3(window *QWindow, parent *QWidget, flags WindowType) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QDockWidget_createWindowContainer3, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()), uintptr(flags))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func (this *QDockWidget) UpdateMicroFocusWithQuery(query InputMethodQuery) {
	purego.SyscallN(addr_QDockWidget_updateMicroFocus2, this.h, uintptr(query))
	runtime.KeepAlive(this)
}

func (this *QDockWidget) CreateWith(arg1 uint64) {
	purego.SyscallN(addr_QDockWidget_create2, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QDockWidget) Create2(arg1 uint64, initializeWindow bool) {
	purego.SyscallN(addr_QDockWidget_create3, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)))
	runtime.KeepAlive(this)
}

func (this *QDockWidget) Create3(arg1 uint64, initializeWindow bool, destroyOldWindow bool) {
	purego.SyscallN(addr_QDockWidget_create4, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)), uintptr(b2i(destroyOldWindow)))
	runtime.KeepAlive(this)
}

func (this *QDockWidget) DestroyWithDestroyWindow(destroyWindow bool) {
	purego.SyscallN(addr_QDockWidget_destroy2, this.h, uintptr(b2i(destroyWindow)))
	runtime.KeepAlive(this)
}

func (this *QDockWidget) Destroy2(destroyWindow bool, destroySubWindows bool) {
	purego.SyscallN(addr_QDockWidget_destroy3, this.h, uintptr(b2i(destroyWindow)), uintptr(b2i(destroySubWindows)))
	runtime.KeepAlive(this)
}

func QDockWidget_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QDockWidget_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QDockWidget_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QDockWidget_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QDockWidget) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QDockWidget_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QDockWidget) ChangeEventDefault(event *QEvent) {
	purego.SyscallN(addr_QDockWidget_changeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDockWidget) CloseEventDefault(event *QCloseEvent) {
	purego.SyscallN(addr_QDockWidget_closeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDockWidget) PaintEventDefault(event *QPaintEvent) {
	purego.SyscallN(addr_QDockWidget_paintEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDockWidget) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QDockWidget_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QDockWidget) InitStyleOptionDefault(option *QStyleOptionDockWidget) {
	purego.SyscallN(addr_QDockWidget_initStyleOption_default, this.h, uintptr(option.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(option)
}

func (this *QDockWidget) DevTypeDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QDockWidget_devType_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QDockWidget) SetVisibleDefault(visible bool) {
	purego.SyscallN(addr_QDockWidget_setVisible_default, this.h, uintptr(b2i(visible)))
	runtime.KeepAlive(this)
}

func (this *QDockWidget) SizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QDockWidget_sizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QDockWidget) MinimumSizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QDockWidget_minimumSizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QDockWidget) HeightForWidthDefault(arg1 int) int {
	__r0, _, _ := purego.SyscallN(addr_QDockWidget_heightForWidth_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QDockWidget) HasHeightForWidthDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QDockWidget_hasHeightForWidth_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QDockWidget) PaintEngineDefault() *QPaintEngine {
	__r, _, _ := purego.SyscallN(addr_QDockWidget_paintEngine_default, this.h)
	runtime.KeepAlive(this)
	return bindQPaintEngine(__r)
}

func (this *QDockWidget) MousePressEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QDockWidget_mousePressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDockWidget) MouseReleaseEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QDockWidget_mouseReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDockWidget) MouseDoubleClickEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QDockWidget_mouseDoubleClickEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDockWidget) MouseMoveEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QDockWidget_mouseMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDockWidget) WheelEventDefault(event *QWheelEvent) {
	purego.SyscallN(addr_QDockWidget_wheelEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDockWidget) KeyPressEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QDockWidget_keyPressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDockWidget) KeyReleaseEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QDockWidget_keyReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDockWidget) FocusInEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QDockWidget_focusInEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDockWidget) FocusOutEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QDockWidget_focusOutEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDockWidget) EnterEventDefault(event *QEnterEvent) {
	purego.SyscallN(addr_QDockWidget_enterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDockWidget) LeaveEventDefault(event *QEvent) {
	purego.SyscallN(addr_QDockWidget_leaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDockWidget) MoveEventDefault(event *QMoveEvent) {
	purego.SyscallN(addr_QDockWidget_moveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDockWidget) ResizeEventDefault(event *QResizeEvent) {
	purego.SyscallN(addr_QDockWidget_resizeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDockWidget) ContextMenuEventDefault(event *QContextMenuEvent) {
	purego.SyscallN(addr_QDockWidget_contextMenuEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDockWidget) TabletEventDefault(event *QTabletEvent) {
	purego.SyscallN(addr_QDockWidget_tabletEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDockWidget) ActionEventDefault(event *QActionEvent) {
	purego.SyscallN(addr_QDockWidget_actionEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDockWidget) DragEnterEventDefault(event *QDragEnterEvent) {
	purego.SyscallN(addr_QDockWidget_dragEnterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDockWidget) DragMoveEventDefault(event *QDragMoveEvent) {
	purego.SyscallN(addr_QDockWidget_dragMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDockWidget) DragLeaveEventDefault(event *QDragLeaveEvent) {
	purego.SyscallN(addr_QDockWidget_dragLeaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDockWidget) DropEventDefault(event *QDropEvent) {
	purego.SyscallN(addr_QDockWidget_dropEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDockWidget) ShowEventDefault(event *QShowEvent) {
	purego.SyscallN(addr_QDockWidget_showEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDockWidget) HideEventDefault(event *QHideEvent) {
	purego.SyscallN(addr_QDockWidget_hideEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDockWidget) MetricDefault(arg1 QPaintDevice__PaintDeviceMetric) int {
	__r0, _, _ := purego.SyscallN(addr_QDockWidget_metric_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QDockWidget) InitPainterDefault(painter *QPainter) {
	purego.SyscallN(addr_QDockWidget_initPainter_default, this.h, uintptr(painter.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
}

func (this *QDockWidget) RedirectedDefault(offset *QPoint) *QPaintDevice {
	__r, _, _ := purego.SyscallN(addr_QDockWidget_redirected_default, this.h, uintptr(offset.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(offset)
	return bindQPaintDevice(__r)
}

func (this *QDockWidget) SharedPainterDefault() *QPainter {
	__r, _, _ := purego.SyscallN(addr_QDockWidget_sharedPainter_default, this.h)
	runtime.KeepAlive(this)
	return bindQPainter(__r)
}

func (this *QDockWidget) InputMethodEventDefault(arg1 *QInputMethodEvent) {
	purego.SyscallN(addr_QDockWidget_inputMethodEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QDockWidget) InputMethodQueryDefault(arg1 InputMethodQuery) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QDockWidget_inputMethodQuery_default, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QDockWidget) FocusNextPrevChildDefault(next bool) bool {
	__r0, _, _ := purego.SyscallN(addr_QDockWidget_focusNextPrevChild_default, this.h, uintptr(b2i(next)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QDockWidget) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QDockWidget_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QDockWidget) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QDockWidget_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDockWidget) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QDockWidget_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDockWidget) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QDockWidget_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDockWidget) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QDockWidget_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QDockWidget) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QDockWidget_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QDockWidget) OnFeaturesChanged(fn func(QDockWidget__DockWidgetFeature)) int64 {
	id := registerCallback(func(features int32) { fn(QDockWidget__DockWidgetFeature(features)) })
	pg_QDockWidget_connect_featuresChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDockWidget) OnTopLevelChanged(fn func(bool)) int64 {
	id := registerCallback(func(topLevel int32) { fn(topLevel != 0) })
	pg_QDockWidget_connect_topLevelChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDockWidget) OnAllowedAreasChanged(fn func(DockWidgetArea)) int64 {
	id := registerCallback(func(allowedAreas int32) { fn(DockWidgetArea(allowedAreas)) })
	pg_QDockWidget_connect_allowedAreasChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDockWidget) OnVisibilityChanged(fn func(bool)) int64 {
	id := registerCallback(func(visible int32) { fn(visible != 0) })
	pg_QDockWidget_connect_visibilityChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDockWidget) OnDockLocationChanged(fn func(DockWidgetArea)) int64 {
	id := registerCallback(func(area int32) { fn(DockWidgetArea(area)) })
	pg_QDockWidget_connect_dockLocationChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDockWidget) OnWindowTitleChanged(fn func(string)) int64 {
	id := registerCallback(func(title uintptr, title_len int64) { fn(pgStrNoFree(title, title_len)) })
	pg_QDockWidget_connect_windowTitleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDockWidget) OnWindowIconChanged(fn func(*QIcon)) int64 {
	id := registerCallback(func(icon uintptr) { fn(UnsafeNewQIcon(unsafe.Pointer(icon))) })
	pg_QDockWidget_connect_windowIconChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDockWidget) OnWindowIconTextChanged(fn func(string)) int64 {
	id := registerCallback(func(iconText uintptr, iconText_len int64) { fn(pgStrNoFree(iconText, iconText_len)) })
	pg_QDockWidget_connect_windowIconTextChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDockWidget) OnCustomContextMenuRequested(fn func(*QPoint)) int64 {
	id := registerCallback(func(pos uintptr) { fn(UnsafeNewQPoint(unsafe.Pointer(pos))) })
	pg_QDockWidget_connect_customContextMenuRequested(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDockWidget) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QDockWidget_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDockWidget) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QDockWidget_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDockWidget) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QDockWidget_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDockWidget) Disconnect(handle int64) {
	pg_QDockWidget__disconnect(handle)
	unregisterCallback(handle)
}

func QDockWidget_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case 4959376063826165943:
		return 1
	case -1989937050019464783:
		return 2
	case 5290442163167458465:
		return 3
	case 1237236456312824991:
		return 4
	case -488693769771308531:
		return 5
	case -8659679893868524044:
		return 6
	case 2110182719375245575:
		return 7
	case 729690261457684101:
		return 8
	case 7049264599837532491:
		return 9
	case 4311534060495737339:
		return 10
	case -8266762397995423915:
		return 11
	case 6398007404756346381:
		return 12
	case -332090898078543319:
		return 13
	case -4415920049297795213:
		return 14
	case -4029073905697023887:
		return 15
	case -4255494004099806581:
		return 16
	case 1919988340934511316:
		return 17
	case 1023791977679793185:
		return 18
	case 772598475950050299:
		return 19
	case 975043767033552552:
		return 20
	case -385596010576048993:
		return 21
	case 6316270346668237895:
		return 22
	case -551070009623711640:
		return 23
	case -9161949150489525396:
		return 24
	case -4302917285655014661:
		return 25
	case -6698043222227006015:
		return 26
	case -212068866860701141:
		return 27
	case 1697444140088555581:
		return 28
	case 206122715919917589:
		return 29
	case -1781545137352745978:
		return 30
	case -3139904158373725006:
		return 31
	case 9107058682754591716:
		return 32
	case -6825369174751916:
		return 33
	case -199718707876719129:
		return 34
	case -2039166586246443706:
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

func QDockWidget_enable(h uintptr, idx int32) { pg_QDockWidget__set_override(h, idx) }

func (this *QDockWidget) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QDockWidget) OnChangeEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 4959376063826165943, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "changeEvent"
}

func (this *QDockWidget) OnCloseEvent(fn func(*QCloseEvent)) {
	registerVirt(this.UnsafeGoID(), -1989937050019464783, func(event uintptr) { fn(UnsafeNewQCloseEvent(unsafe.Pointer(event))) }) // slot "closeEvent"
}

func (this *QDockWidget) OnPaintEvent(fn func(*QPaintEvent)) {
	registerVirt(this.UnsafeGoID(), 5290442163167458465, func(event uintptr) { fn(UnsafeNewQPaintEvent(unsafe.Pointer(event))) }) // slot "paintEvent"
}

func (this *QDockWidget) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QDockWidget) OnInitStyleOption(fn func(*QStyleOptionDockWidget)) {
	registerVirt(this.UnsafeGoID(), -488693769771308531, func(option uintptr) { fn(UnsafeNewQStyleOptionDockWidget(unsafe.Pointer(option))) }) // slot "initStyleOption"
}

func (this *QDockWidget) OnDevType(fn func() int) {
	registerVirt(this.UnsafeGoID(), -8659679893868524044, func() int32 { return int32(fn()) }) // slot "devType"
}

func (this *QDockWidget) OnSetVisible(fn func(bool)) {
	registerVirt(this.UnsafeGoID(), 2110182719375245575, func(visible int32) { fn(visible != 0) }) // slot "setVisible"
}

func (this *QDockWidget) OnSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 729690261457684101, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sizeHint"
}

func (this *QDockWidget) OnMinimumSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 7049264599837532491, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "minimumSizeHint"
}

func (this *QDockWidget) OnHeightForWidth(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), 4311534060495737339, func(arg1 int32) int32 { return int32(fn(int(arg1))) }) // slot "heightForWidth"
}

func (this *QDockWidget) OnHasHeightForWidth(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -8266762397995423915, func() int32 { return b2i(fn()) }) // slot "hasHeightForWidth"
}

func (this *QDockWidget) OnPaintEngine(fn func() *QPaintEngine) {
	registerVirt(this.UnsafeGoID(), 6398007404756346381, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "paintEngine"
}

func (this *QDockWidget) OnMousePressEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -332090898078543319, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mousePressEvent"
}

func (this *QDockWidget) OnMouseReleaseEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4415920049297795213, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseReleaseEvent"
}

func (this *QDockWidget) OnMouseDoubleClickEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4029073905697023887, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseDoubleClickEvent"
}

func (this *QDockWidget) OnMouseMoveEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4255494004099806581, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseMoveEvent"
}

func (this *QDockWidget) OnWheelEvent(fn func(*QWheelEvent)) {
	registerVirt(this.UnsafeGoID(), 1919988340934511316, func(event uintptr) { fn(UnsafeNewQWheelEvent(unsafe.Pointer(event))) }) // slot "wheelEvent"
}

func (this *QDockWidget) OnKeyPressEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 1023791977679793185, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyPressEvent"
}

func (this *QDockWidget) OnKeyReleaseEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 772598475950050299, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyReleaseEvent"
}

func (this *QDockWidget) OnFocusInEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), 975043767033552552, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusInEvent"
}

func (this *QDockWidget) OnFocusOutEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), -385596010576048993, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusOutEvent"
}

func (this *QDockWidget) OnEnterEvent(fn func(*QEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 6316270346668237895, func(event uintptr) { fn(UnsafeNewQEnterEvent(unsafe.Pointer(event))) }) // slot "enterEvent"
}

func (this *QDockWidget) OnLeaveEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), -551070009623711640, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "leaveEvent"
}

func (this *QDockWidget) OnMoveEvent(fn func(*QMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -9161949150489525396, func(event uintptr) { fn(UnsafeNewQMoveEvent(unsafe.Pointer(event))) }) // slot "moveEvent"
}

func (this *QDockWidget) OnResizeEvent(fn func(*QResizeEvent)) {
	registerVirt(this.UnsafeGoID(), -4302917285655014661, func(event uintptr) { fn(UnsafeNewQResizeEvent(unsafe.Pointer(event))) }) // slot "resizeEvent"
}

func (this *QDockWidget) OnContextMenuEvent(fn func(*QContextMenuEvent)) {
	registerVirt(this.UnsafeGoID(), -6698043222227006015, func(event uintptr) { fn(UnsafeNewQContextMenuEvent(unsafe.Pointer(event))) }) // slot "contextMenuEvent"
}

func (this *QDockWidget) OnTabletEvent(fn func(*QTabletEvent)) {
	registerVirt(this.UnsafeGoID(), -212068866860701141, func(event uintptr) { fn(UnsafeNewQTabletEvent(unsafe.Pointer(event))) }) // slot "tabletEvent"
}

func (this *QDockWidget) OnActionEvent(fn func(*QActionEvent)) {
	registerVirt(this.UnsafeGoID(), 1697444140088555581, func(event uintptr) { fn(UnsafeNewQActionEvent(unsafe.Pointer(event))) }) // slot "actionEvent"
}

func (this *QDockWidget) OnDragEnterEvent(fn func(*QDragEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 206122715919917589, func(event uintptr) { fn(UnsafeNewQDragEnterEvent(unsafe.Pointer(event))) }) // slot "dragEnterEvent"
}

func (this *QDockWidget) OnDragMoveEvent(fn func(*QDragMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -1781545137352745978, func(event uintptr) { fn(UnsafeNewQDragMoveEvent(unsafe.Pointer(event))) }) // slot "dragMoveEvent"
}

func (this *QDockWidget) OnDragLeaveEvent(fn func(*QDragLeaveEvent)) {
	registerVirt(this.UnsafeGoID(), -3139904158373725006, func(event uintptr) { fn(UnsafeNewQDragLeaveEvent(unsafe.Pointer(event))) }) // slot "dragLeaveEvent"
}

func (this *QDockWidget) OnDropEvent(fn func(*QDropEvent)) {
	registerVirt(this.UnsafeGoID(), 9107058682754591716, func(event uintptr) { fn(UnsafeNewQDropEvent(unsafe.Pointer(event))) }) // slot "dropEvent"
}

func (this *QDockWidget) OnShowEvent(fn func(*QShowEvent)) {
	registerVirt(this.UnsafeGoID(), -6825369174751916, func(event uintptr) { fn(UnsafeNewQShowEvent(unsafe.Pointer(event))) }) // slot "showEvent"
}

func (this *QDockWidget) OnHideEvent(fn func(*QHideEvent)) {
	registerVirt(this.UnsafeGoID(), -199718707876719129, func(event uintptr) { fn(UnsafeNewQHideEvent(unsafe.Pointer(event))) }) // slot "hideEvent"
}

func (this *QDockWidget) OnNativeEvent(fn func(*QByteArray, unsafe.Pointer, *int64) bool) {
	registerVirt(this.UnsafeGoID(), -2039166586246443706, func(eventType uintptr, message uintptr, result uintptr) int32 { return b2i(fn(UnsafeNewQByteArray(unsafe.Pointer(eventType)), unsafe.Pointer(message), (*int64)(unsafe.Pointer(result)))) }) // slot "nativeEvent"
}

func (this *QDockWidget) OnMetric(fn func(QPaintDevice__PaintDeviceMetric) int) {
	registerVirt(this.UnsafeGoID(), 427629649267311777, func(arg1 int32) int32 { return int32(fn(QPaintDevice__PaintDeviceMetric(arg1))) }) // slot "metric"
}

func (this *QDockWidget) OnInitPainter(fn func(*QPainter)) {
	registerVirt(this.UnsafeGoID(), 6477500296980050210, func(painter uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter))) }) // slot "initPainter"
}

func (this *QDockWidget) OnRedirected(fn func(*QPoint) *QPaintDevice) {
	registerVirt(this.UnsafeGoID(), 4127240633310655048, func(offset uintptr) uintptr { return uintptr(fn(UnsafeNewQPoint(unsafe.Pointer(offset))).UnsafePointer()) }) // slot "redirected"
}

func (this *QDockWidget) OnSharedPainter(fn func() *QPainter) {
	registerVirt(this.UnsafeGoID(), 5792771328741774991, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sharedPainter"
}

func (this *QDockWidget) OnInputMethodEvent(fn func(*QInputMethodEvent)) {
	registerVirt(this.UnsafeGoID(), 5438819661942843222, func(arg1 uintptr) { fn(UnsafeNewQInputMethodEvent(unsafe.Pointer(arg1))) }) // slot "inputMethodEvent"
}

func (this *QDockWidget) OnInputMethodQuery(fn func(InputMethodQuery) *QVariant) {
	registerVirt(this.UnsafeGoID(), 7345853097340709638, func(arg1 int32) uintptr { return uintptr(fn(InputMethodQuery(arg1)).UnsafePointer()) }) // slot "inputMethodQuery"
}

func (this *QDockWidget) OnFocusNextPrevChild(fn func(bool) bool) {
	registerVirt(this.UnsafeGoID(), 5898328517505784191, func(next int32) int32 { return b2i(fn(next != 0)) }) // slot "focusNextPrevChild"
}

func (this *QDockWidget) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QDockWidget) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QDockWidget) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QDockWidget) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QDockWidget) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QDockWidget) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QDockWidget) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QDockWidget_delete(this.h)
	}
	this.UnsafeRelease()
}
