package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QToolBar_chain uintptr
	addr_QToolBar_new uintptr
	addr_QToolBar_new2 uintptr
	addr_QToolBar_new3 uintptr
	addr_QToolBar_new4 uintptr
	addr_QToolBar_tr uintptr
	addr_QToolBar_setMovable uintptr
	addr_QToolBar_isMovable uintptr
	addr_QToolBar_setAllowedAreas uintptr
	addr_QToolBar_allowedAreas uintptr
	addr_QToolBar_isAreaAllowed uintptr
	addr_QToolBar_setOrientation uintptr
	addr_QToolBar_orientation uintptr
	addr_QToolBar_clear uintptr
	addr_QToolBar_addSeparator uintptr
	addr_QToolBar_insertSeparator uintptr
	addr_QToolBar_addWidget uintptr
	addr_QToolBar_insertWidget uintptr
	addr_QToolBar_actionGeometry uintptr
	addr_QToolBar_actionAt uintptr
	addr_QToolBar_actionAt2 uintptr
	addr_QToolBar_toggleViewAction uintptr
	addr_QToolBar_iconSize uintptr
	addr_QToolBar_toolButtonStyle uintptr
	addr_QToolBar_widgetForAction uintptr
	addr_QToolBar_isFloatable uintptr
	addr_QToolBar_setFloatable uintptr
	addr_QToolBar_isFloating uintptr
	addr_QToolBar_setIconSize uintptr
	addr_QToolBar_setToolButtonStyle uintptr
	addr_QToolBar_actionTriggered uintptr
	addr_QToolBar_movableChanged uintptr
	addr_QToolBar_allowedAreasChanged uintptr
	addr_QToolBar_orientationChanged uintptr
	addr_QToolBar_iconSizeChanged uintptr
	addr_QToolBar_toolButtonStyleChanged uintptr
	addr_QToolBar_topLevelChanged uintptr
	addr_QToolBar_visibilityChanged uintptr
	addr_QToolBar_setTabOrder uintptr
	addr_QToolBar_mouseGrabber uintptr
	addr_QToolBar_keyboardGrabber uintptr
	addr_QToolBar_find uintptr
	addr_QToolBar_createWindowContainer uintptr
	addr_QToolBar_updateMicroFocus uintptr
	addr_QToolBar_create uintptr
	addr_QToolBar_destroy uintptr
	addr_QToolBar_focusNextChild uintptr
	addr_QToolBar_focusPreviousChild uintptr
	addr_QToolBar_connect uintptr
	addr_QToolBar_connect2 uintptr
	addr_QToolBar_sender uintptr
	addr_QToolBar_senderSignalIndex uintptr
	addr_QToolBar_receivers uintptr
	addr_QToolBar_isSignalConnected uintptr
	pg_QToolBar_devicePixelRatioFScale func() float64
	pg_QToolBar_encodeMetricF func(int32, float64) int32
	pg_QToolBar_getDecodedMetricF func(uintptr, int32, int32) float64
	addr_QToolBar_tr2 uintptr
	addr_QToolBar_tr3 uintptr
	addr_QToolBar_createWindowContainer2 uintptr
	addr_QToolBar_createWindowContainer3 uintptr
	addr_QToolBar_updateMicroFocus2 uintptr
	addr_QToolBar_create2 uintptr
	addr_QToolBar_create3 uintptr
	addr_QToolBar_create4 uintptr
	addr_QToolBar_destroy2 uintptr
	addr_QToolBar_destroy3 uintptr
	addr_QToolBar_connect3 uintptr
	addr_QToolBar_connect4 uintptr
	addr_QToolBar_metaObject_default uintptr
	addr_QToolBar_actionEvent_default uintptr
	addr_QToolBar_changeEvent_default uintptr
	addr_QToolBar_paintEvent_default uintptr
	addr_QToolBar_event_default uintptr
	addr_QToolBar_initStyleOption_default uintptr
	addr_QToolBar_devType_default uintptr
	addr_QToolBar_setVisible_default uintptr
	addr_QToolBar_sizeHint_default uintptr
	addr_QToolBar_minimumSizeHint_default uintptr
	addr_QToolBar_heightForWidth_default uintptr
	addr_QToolBar_hasHeightForWidth_default uintptr
	addr_QToolBar_paintEngine_default uintptr
	addr_QToolBar_mousePressEvent_default uintptr
	addr_QToolBar_mouseReleaseEvent_default uintptr
	addr_QToolBar_mouseDoubleClickEvent_default uintptr
	addr_QToolBar_mouseMoveEvent_default uintptr
	addr_QToolBar_wheelEvent_default uintptr
	addr_QToolBar_keyPressEvent_default uintptr
	addr_QToolBar_keyReleaseEvent_default uintptr
	addr_QToolBar_focusInEvent_default uintptr
	addr_QToolBar_focusOutEvent_default uintptr
	addr_QToolBar_enterEvent_default uintptr
	addr_QToolBar_leaveEvent_default uintptr
	addr_QToolBar_moveEvent_default uintptr
	addr_QToolBar_resizeEvent_default uintptr
	addr_QToolBar_closeEvent_default uintptr
	addr_QToolBar_contextMenuEvent_default uintptr
	addr_QToolBar_tabletEvent_default uintptr
	addr_QToolBar_dragEnterEvent_default uintptr
	addr_QToolBar_dragMoveEvent_default uintptr
	addr_QToolBar_dragLeaveEvent_default uintptr
	addr_QToolBar_dropEvent_default uintptr
	addr_QToolBar_showEvent_default uintptr
	addr_QToolBar_hideEvent_default uintptr
	addr_QToolBar_metric_default uintptr
	addr_QToolBar_initPainter_default uintptr
	addr_QToolBar_redirected_default uintptr
	addr_QToolBar_sharedPainter_default uintptr
	addr_QToolBar_inputMethodEvent_default uintptr
	addr_QToolBar_inputMethodQuery_default uintptr
	addr_QToolBar_focusNextPrevChild_default uintptr
	addr_QToolBar_eventFilter_default uintptr
	addr_QToolBar_timerEvent_default uintptr
	addr_QToolBar_childEvent_default uintptr
	addr_QToolBar_customEvent_default uintptr
	addr_QToolBar_connectNotify_default uintptr
	addr_QToolBar_disconnectNotify_default uintptr
	pg_QToolBar_connect_actionTriggered func(uintptr, int64)
	pg_QToolBar_connect_movableChanged func(uintptr, int64)
	pg_QToolBar_connect_allowedAreasChanged func(uintptr, int64)
	pg_QToolBar_connect_orientationChanged func(uintptr, int64)
	pg_QToolBar_connect_iconSizeChanged func(uintptr, int64)
	pg_QToolBar_connect_toolButtonStyleChanged func(uintptr, int64)
	pg_QToolBar_connect_topLevelChanged func(uintptr, int64)
	pg_QToolBar_connect_visibilityChanged func(uintptr, int64)
	pg_QToolBar_connect_windowTitleChanged func(uintptr, int64)
	pg_QToolBar_connect_windowIconChanged func(uintptr, int64)
	pg_QToolBar_connect_windowIconTextChanged func(uintptr, int64)
	pg_QToolBar_connect_customContextMenuRequested func(uintptr, int64)
	pg_QToolBar_connect_destroyed func(uintptr, int64)
	pg_QToolBar_connect_destroyed2 func(uintptr, int64)
	pg_QToolBar_connect_objectNameChanged func(uintptr, int64)
	pg_QToolBar__disconnect func(int64)
	pg_QToolBar__set_override func(uintptr, int32)
	pg_QToolBar_delete func(uintptr)
)

func init() {
	symaddr(&addr_QToolBar_chain, "QToolBar_chain")
	symaddr(&addr_QToolBar_new, "QToolBar_new")
	symaddr(&addr_QToolBar_new2, "QToolBar_new2")
	symaddr(&addr_QToolBar_new3, "QToolBar_new3")
	symaddr(&addr_QToolBar_new4, "QToolBar_new4")
	symaddr(&addr_QToolBar_tr, "QToolBar_tr")
	symaddr(&addr_QToolBar_setMovable, "QToolBar_setMovable")
	symaddr(&addr_QToolBar_isMovable, "QToolBar_isMovable")
	symaddr(&addr_QToolBar_setAllowedAreas, "QToolBar_setAllowedAreas")
	symaddr(&addr_QToolBar_allowedAreas, "QToolBar_allowedAreas")
	symaddr(&addr_QToolBar_isAreaAllowed, "QToolBar_isAreaAllowed")
	symaddr(&addr_QToolBar_setOrientation, "QToolBar_setOrientation")
	symaddr(&addr_QToolBar_orientation, "QToolBar_orientation")
	symaddr(&addr_QToolBar_clear, "QToolBar_clear")
	symaddr(&addr_QToolBar_addSeparator, "QToolBar_addSeparator")
	symaddr(&addr_QToolBar_insertSeparator, "QToolBar_insertSeparator")
	symaddr(&addr_QToolBar_addWidget, "QToolBar_addWidget")
	symaddr(&addr_QToolBar_insertWidget, "QToolBar_insertWidget")
	symaddr(&addr_QToolBar_actionGeometry, "QToolBar_actionGeometry")
	symaddr(&addr_QToolBar_actionAt, "QToolBar_actionAt")
	symaddr(&addr_QToolBar_actionAt2, "QToolBar_actionAt2")
	symaddr(&addr_QToolBar_toggleViewAction, "QToolBar_toggleViewAction")
	symaddr(&addr_QToolBar_iconSize, "QToolBar_iconSize")
	symaddr(&addr_QToolBar_toolButtonStyle, "QToolBar_toolButtonStyle")
	symaddr(&addr_QToolBar_widgetForAction, "QToolBar_widgetForAction")
	symaddr(&addr_QToolBar_isFloatable, "QToolBar_isFloatable")
	symaddr(&addr_QToolBar_setFloatable, "QToolBar_setFloatable")
	symaddr(&addr_QToolBar_isFloating, "QToolBar_isFloating")
	symaddr(&addr_QToolBar_setIconSize, "QToolBar_setIconSize")
	symaddr(&addr_QToolBar_setToolButtonStyle, "QToolBar_setToolButtonStyle")
	symaddr(&addr_QToolBar_actionTriggered, "QToolBar_actionTriggered")
	symaddr(&addr_QToolBar_movableChanged, "QToolBar_movableChanged")
	symaddr(&addr_QToolBar_allowedAreasChanged, "QToolBar_allowedAreasChanged")
	symaddr(&addr_QToolBar_orientationChanged, "QToolBar_orientationChanged")
	symaddr(&addr_QToolBar_iconSizeChanged, "QToolBar_iconSizeChanged")
	symaddr(&addr_QToolBar_toolButtonStyleChanged, "QToolBar_toolButtonStyleChanged")
	symaddr(&addr_QToolBar_topLevelChanged, "QToolBar_topLevelChanged")
	symaddr(&addr_QToolBar_visibilityChanged, "QToolBar_visibilityChanged")
	symaddr(&addr_QToolBar_setTabOrder, "QToolBar_setTabOrder")
	symaddr(&addr_QToolBar_mouseGrabber, "QToolBar_mouseGrabber")
	symaddr(&addr_QToolBar_keyboardGrabber, "QToolBar_keyboardGrabber")
	symaddr(&addr_QToolBar_find, "QToolBar_find")
	symaddr(&addr_QToolBar_createWindowContainer, "QToolBar_createWindowContainer")
	symaddr(&addr_QToolBar_updateMicroFocus, "QToolBar_updateMicroFocus")
	symaddr(&addr_QToolBar_create, "QToolBar_create")
	symaddr(&addr_QToolBar_destroy, "QToolBar_destroy")
	symaddr(&addr_QToolBar_focusNextChild, "QToolBar_focusNextChild")
	symaddr(&addr_QToolBar_focusPreviousChild, "QToolBar_focusPreviousChild")
	symaddr(&addr_QToolBar_connect, "QToolBar_connect")
	symaddr(&addr_QToolBar_connect2, "QToolBar_connect2")
	symaddr(&addr_QToolBar_sender, "QToolBar_sender")
	symaddr(&addr_QToolBar_senderSignalIndex, "QToolBar_senderSignalIndex")
	symaddr(&addr_QToolBar_receivers, "QToolBar_receivers")
	symaddr(&addr_QToolBar_isSignalConnected, "QToolBar_isSignalConnected")
	reg(&pg_QToolBar_devicePixelRatioFScale, "QToolBar_devicePixelRatioFScale")
	reg(&pg_QToolBar_encodeMetricF, "QToolBar_encodeMetricF")
	reg(&pg_QToolBar_getDecodedMetricF, "QToolBar_getDecodedMetricF")
	symaddr(&addr_QToolBar_tr2, "QToolBar_tr2")
	symaddr(&addr_QToolBar_tr3, "QToolBar_tr3")
	symaddr(&addr_QToolBar_createWindowContainer2, "QToolBar_createWindowContainer2")
	symaddr(&addr_QToolBar_createWindowContainer3, "QToolBar_createWindowContainer3")
	symaddr(&addr_QToolBar_updateMicroFocus2, "QToolBar_updateMicroFocus2")
	symaddr(&addr_QToolBar_create2, "QToolBar_create2")
	symaddr(&addr_QToolBar_create3, "QToolBar_create3")
	symaddr(&addr_QToolBar_create4, "QToolBar_create4")
	symaddr(&addr_QToolBar_destroy2, "QToolBar_destroy2")
	symaddr(&addr_QToolBar_destroy3, "QToolBar_destroy3")
	symaddr(&addr_QToolBar_connect3, "QToolBar_connect3")
	symaddr(&addr_QToolBar_connect4, "QToolBar_connect4")
	symaddr(&addr_QToolBar_metaObject_default, "QToolBar_metaObject_default")
	symaddr(&addr_QToolBar_actionEvent_default, "QToolBar_actionEvent_default")
	symaddr(&addr_QToolBar_changeEvent_default, "QToolBar_changeEvent_default")
	symaddr(&addr_QToolBar_paintEvent_default, "QToolBar_paintEvent_default")
	symaddr(&addr_QToolBar_event_default, "QToolBar_event_default")
	symaddr(&addr_QToolBar_initStyleOption_default, "QToolBar_initStyleOption_default")
	symaddr(&addr_QToolBar_devType_default, "QToolBar_devType_default")
	symaddr(&addr_QToolBar_setVisible_default, "QToolBar_setVisible_default")
	symaddr(&addr_QToolBar_sizeHint_default, "QToolBar_sizeHint_default")
	symaddr(&addr_QToolBar_minimumSizeHint_default, "QToolBar_minimumSizeHint_default")
	symaddr(&addr_QToolBar_heightForWidth_default, "QToolBar_heightForWidth_default")
	symaddr(&addr_QToolBar_hasHeightForWidth_default, "QToolBar_hasHeightForWidth_default")
	symaddr(&addr_QToolBar_paintEngine_default, "QToolBar_paintEngine_default")
	symaddr(&addr_QToolBar_mousePressEvent_default, "QToolBar_mousePressEvent_default")
	symaddr(&addr_QToolBar_mouseReleaseEvent_default, "QToolBar_mouseReleaseEvent_default")
	symaddr(&addr_QToolBar_mouseDoubleClickEvent_default, "QToolBar_mouseDoubleClickEvent_default")
	symaddr(&addr_QToolBar_mouseMoveEvent_default, "QToolBar_mouseMoveEvent_default")
	symaddr(&addr_QToolBar_wheelEvent_default, "QToolBar_wheelEvent_default")
	symaddr(&addr_QToolBar_keyPressEvent_default, "QToolBar_keyPressEvent_default")
	symaddr(&addr_QToolBar_keyReleaseEvent_default, "QToolBar_keyReleaseEvent_default")
	symaddr(&addr_QToolBar_focusInEvent_default, "QToolBar_focusInEvent_default")
	symaddr(&addr_QToolBar_focusOutEvent_default, "QToolBar_focusOutEvent_default")
	symaddr(&addr_QToolBar_enterEvent_default, "QToolBar_enterEvent_default")
	symaddr(&addr_QToolBar_leaveEvent_default, "QToolBar_leaveEvent_default")
	symaddr(&addr_QToolBar_moveEvent_default, "QToolBar_moveEvent_default")
	symaddr(&addr_QToolBar_resizeEvent_default, "QToolBar_resizeEvent_default")
	symaddr(&addr_QToolBar_closeEvent_default, "QToolBar_closeEvent_default")
	symaddr(&addr_QToolBar_contextMenuEvent_default, "QToolBar_contextMenuEvent_default")
	symaddr(&addr_QToolBar_tabletEvent_default, "QToolBar_tabletEvent_default")
	symaddr(&addr_QToolBar_dragEnterEvent_default, "QToolBar_dragEnterEvent_default")
	symaddr(&addr_QToolBar_dragMoveEvent_default, "QToolBar_dragMoveEvent_default")
	symaddr(&addr_QToolBar_dragLeaveEvent_default, "QToolBar_dragLeaveEvent_default")
	symaddr(&addr_QToolBar_dropEvent_default, "QToolBar_dropEvent_default")
	symaddr(&addr_QToolBar_showEvent_default, "QToolBar_showEvent_default")
	symaddr(&addr_QToolBar_hideEvent_default, "QToolBar_hideEvent_default")
	symaddr(&addr_QToolBar_metric_default, "QToolBar_metric_default")
	symaddr(&addr_QToolBar_initPainter_default, "QToolBar_initPainter_default")
	symaddr(&addr_QToolBar_redirected_default, "QToolBar_redirected_default")
	symaddr(&addr_QToolBar_sharedPainter_default, "QToolBar_sharedPainter_default")
	symaddr(&addr_QToolBar_inputMethodEvent_default, "QToolBar_inputMethodEvent_default")
	symaddr(&addr_QToolBar_inputMethodQuery_default, "QToolBar_inputMethodQuery_default")
	symaddr(&addr_QToolBar_focusNextPrevChild_default, "QToolBar_focusNextPrevChild_default")
	symaddr(&addr_QToolBar_eventFilter_default, "QToolBar_eventFilter_default")
	symaddr(&addr_QToolBar_timerEvent_default, "QToolBar_timerEvent_default")
	symaddr(&addr_QToolBar_childEvent_default, "QToolBar_childEvent_default")
	symaddr(&addr_QToolBar_customEvent_default, "QToolBar_customEvent_default")
	symaddr(&addr_QToolBar_connectNotify_default, "QToolBar_connectNotify_default")
	symaddr(&addr_QToolBar_disconnectNotify_default, "QToolBar_disconnectNotify_default")
	reg(&pg_QToolBar_connect_actionTriggered, "QToolBar_connect_actionTriggered")
	reg(&pg_QToolBar_connect_movableChanged, "QToolBar_connect_movableChanged")
	reg(&pg_QToolBar_connect_allowedAreasChanged, "QToolBar_connect_allowedAreasChanged")
	reg(&pg_QToolBar_connect_orientationChanged, "QToolBar_connect_orientationChanged")
	reg(&pg_QToolBar_connect_iconSizeChanged, "QToolBar_connect_iconSizeChanged")
	reg(&pg_QToolBar_connect_toolButtonStyleChanged, "QToolBar_connect_toolButtonStyleChanged")
	reg(&pg_QToolBar_connect_topLevelChanged, "QToolBar_connect_topLevelChanged")
	reg(&pg_QToolBar_connect_visibilityChanged, "QToolBar_connect_visibilityChanged")
	reg(&pg_QToolBar_connect_windowTitleChanged, "QToolBar_connect_windowTitleChanged")
	reg(&pg_QToolBar_connect_windowIconChanged, "QToolBar_connect_windowIconChanged")
	reg(&pg_QToolBar_connect_windowIconTextChanged, "QToolBar_connect_windowIconTextChanged")
	reg(&pg_QToolBar_connect_customContextMenuRequested, "QToolBar_connect_customContextMenuRequested")
	reg(&pg_QToolBar_connect_destroyed, "QToolBar_connect_destroyed")
	reg(&pg_QToolBar_connect_destroyed2, "QToolBar_connect_destroyed2")
	reg(&pg_QToolBar_connect_objectNameChanged, "QToolBar_connect_objectNameChanged")
	reg(&pg_QToolBar__disconnect, "QToolBar__disconnect")
	reg(&pg_QToolBar__set_override, "QToolBar__set_override")
	reg(&pg_QToolBar_delete, "QToolBar_delete")
}

type QToolBar struct {
	h uintptr
	*QWidget
}

func bindQToolBar(h uintptr) *QToolBar {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QToolBar_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QToolBar{h: h, QWidget: &QWidget{h: __c[0], QObject: &QObject{h: __c[1]}}}
}

func newownedQToolBar(h uintptr) *QToolBar { return bindQToolBar(h) }

func (this *QToolBar) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QToolBar) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQToolBar(h unsafe.Pointer) *QToolBar { return bindQToolBar(uintptr(h)) }
func UnsafeNewQToolBarOwned(h unsafe.Pointer) *QToolBar { return newownedQToolBar(uintptr(h)) }

func (this *QToolBar) UnsafeRelease() {
	this.h = 0
	if this.QWidget != nil {
		this.QWidget.UnsafeRelease()
	}
}

func NewQToolBar(parent *QWidget) *QToolBar {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QToolBar_new, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQToolBar(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QToolBar_virtIdx, QToolBar_enable)
	return __o
}

func NewQToolBar2(title string) *QToolBar {
	__goid := nextVirtID()
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	__r, _, _ := purego.SyscallN(addr_QToolBar_new2, uintptr(__goid), title_p, uintptr(len(title)))
	__o := newownedQToolBar(__r)
	runtime.KeepAlive(title)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QToolBar_virtIdx, QToolBar_enable)
	return __o
}

func NewQToolBar3() *QToolBar {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QToolBar_new3, uintptr(__goid))
	__o := newownedQToolBar(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QToolBar_virtIdx, QToolBar_enable)
	return __o
}

func NewQToolBar4(title string, parent *QWidget) *QToolBar {
	__goid := nextVirtID()
	var title_p uintptr
	if len(title) > 0 {
		title_p = uintptr(unsafe.Pointer(unsafe.StringData(title)))
	}
	__r, _, _ := purego.SyscallN(addr_QToolBar_new4, uintptr(__goid), title_p, uintptr(len(title)), uintptr(parent.UnsafePointer()))
	__o := newownedQToolBar(__r)
	runtime.KeepAlive(title)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QToolBar_virtIdx, QToolBar_enable)
	return __o
}

func QToolBar_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QToolBar_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QToolBar) SetMovable(movable bool) {
	purego.SyscallN(addr_QToolBar_setMovable, this.h, uintptr(b2i(movable)))
	runtime.KeepAlive(this)
}

func (this *QToolBar) IsMovable() bool {
	__r0, _, _ := purego.SyscallN(addr_QToolBar_isMovable, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QToolBar) SetAllowedAreas(areas ToolBarArea) {
	purego.SyscallN(addr_QToolBar_setAllowedAreas, this.h, uintptr(areas))
	runtime.KeepAlive(this)
}

func (this *QToolBar) AllowedAreas() ToolBarArea {
	__r0, _, _ := purego.SyscallN(addr_QToolBar_allowedAreas, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return ToolBarArea(__r)
}

func (this *QToolBar) IsAreaAllowed(area ToolBarArea) bool {
	__r0, _, _ := purego.SyscallN(addr_QToolBar_isAreaAllowed, this.h, uintptr(area))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QToolBar) SetOrientation(orientation Orientation) {
	purego.SyscallN(addr_QToolBar_setOrientation, this.h, uintptr(orientation))
	runtime.KeepAlive(this)
}

func (this *QToolBar) Orientation() Orientation {
	__r0, _, _ := purego.SyscallN(addr_QToolBar_orientation, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return Orientation(__r)
}

func (this *QToolBar) Clear() {
	purego.SyscallN(addr_QToolBar_clear, this.h)
	runtime.KeepAlive(this)
}

func (this *QToolBar) AddSeparator() *QAction {
	__r, _, _ := purego.SyscallN(addr_QToolBar_addSeparator, this.h)
	runtime.KeepAlive(this)
	return bindQAction(__r)
}

func (this *QToolBar) InsertSeparator(before *QAction) *QAction {
	__r, _, _ := purego.SyscallN(addr_QToolBar_insertSeparator, this.h, uintptr(before.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(before)
	return bindQAction(__r)
}

func (this *QToolBar) AddWidget(widget *QWidget) *QAction {
	__r, _, _ := purego.SyscallN(addr_QToolBar_addWidget, this.h, uintptr(widget.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(widget)
	return bindQAction(__r)
}

func (this *QToolBar) InsertWidget(before *QAction, widget *QWidget) *QAction {
	__r, _, _ := purego.SyscallN(addr_QToolBar_insertWidget, this.h, uintptr(before.UnsafePointer()), uintptr(widget.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(before)
	runtime.KeepAlive(widget)
	return bindQAction(__r)
}

func (this *QToolBar) ActionGeometry(action *QAction) *QRect {
	__r, _, _ := purego.SyscallN(addr_QToolBar_actionGeometry, this.h, uintptr(action.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(action)
	return newownedQRect(__r)
}

func (this *QToolBar) ActionAt(p *QPoint) *QAction {
	__r, _, _ := purego.SyscallN(addr_QToolBar_actionAt, this.h, uintptr(p.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(p)
	return bindQAction(__r)
}

func (this *QToolBar) ActionAt2(x int, y int) *QAction {
	__r, _, _ := purego.SyscallN(addr_QToolBar_actionAt2, this.h, uintptr(x), uintptr(y))
	runtime.KeepAlive(this)
	return bindQAction(__r)
}

func (this *QToolBar) ToggleViewAction() *QAction {
	__r, _, _ := purego.SyscallN(addr_QToolBar_toggleViewAction, this.h)
	runtime.KeepAlive(this)
	return bindQAction(__r)
}

func (this *QToolBar) IconSize() *QSize {
	__r, _, _ := purego.SyscallN(addr_QToolBar_iconSize, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QToolBar) ToolButtonStyle() ToolButtonStyle {
	__r0, _, _ := purego.SyscallN(addr_QToolBar_toolButtonStyle, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return ToolButtonStyle(__r)
}

func (this *QToolBar) WidgetForAction(action *QAction) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QToolBar_widgetForAction, this.h, uintptr(action.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(action)
	return bindQWidget(__r)
}

func (this *QToolBar) IsFloatable() bool {
	__r0, _, _ := purego.SyscallN(addr_QToolBar_isFloatable, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QToolBar) SetFloatable(floatable bool) {
	purego.SyscallN(addr_QToolBar_setFloatable, this.h, uintptr(b2i(floatable)))
	runtime.KeepAlive(this)
}

func (this *QToolBar) IsFloating() bool {
	__r0, _, _ := purego.SyscallN(addr_QToolBar_isFloating, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QToolBar) SetIconSize(iconSize *QSize) {
	purego.SyscallN(addr_QToolBar_setIconSize, this.h, uintptr(iconSize.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(iconSize)
}

func (this *QToolBar) SetToolButtonStyle(toolButtonStyle ToolButtonStyle) {
	purego.SyscallN(addr_QToolBar_setToolButtonStyle, this.h, uintptr(toolButtonStyle))
	runtime.KeepAlive(this)
}

func (this *QToolBar) ActionTriggered(action *QAction) {
	purego.SyscallN(addr_QToolBar_actionTriggered, this.h, uintptr(action.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(action)
}

func (this *QToolBar) MovableChanged(movable bool) {
	purego.SyscallN(addr_QToolBar_movableChanged, this.h, uintptr(b2i(movable)))
	runtime.KeepAlive(this)
}

func (this *QToolBar) AllowedAreasChanged(allowedAreas ToolBarArea) {
	purego.SyscallN(addr_QToolBar_allowedAreasChanged, this.h, uintptr(allowedAreas))
	runtime.KeepAlive(this)
}

func (this *QToolBar) OrientationChanged(orientation Orientation) {
	purego.SyscallN(addr_QToolBar_orientationChanged, this.h, uintptr(orientation))
	runtime.KeepAlive(this)
}

func (this *QToolBar) IconSizeChanged(iconSize *QSize) {
	purego.SyscallN(addr_QToolBar_iconSizeChanged, this.h, uintptr(iconSize.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(iconSize)
}

func (this *QToolBar) ToolButtonStyleChanged(toolButtonStyle ToolButtonStyle) {
	purego.SyscallN(addr_QToolBar_toolButtonStyleChanged, this.h, uintptr(toolButtonStyle))
	runtime.KeepAlive(this)
}

func (this *QToolBar) TopLevelChanged(topLevel bool) {
	purego.SyscallN(addr_QToolBar_topLevelChanged, this.h, uintptr(b2i(topLevel)))
	runtime.KeepAlive(this)
}

func (this *QToolBar) VisibilityChanged(visible bool) {
	purego.SyscallN(addr_QToolBar_visibilityChanged, this.h, uintptr(b2i(visible)))
	runtime.KeepAlive(this)
}

func QToolBar_SetTabOrder(arg1 *QWidget, arg2 *QWidget) {
	purego.SyscallN(addr_QToolBar_setTabOrder, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()))
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
}

func QToolBar_MouseGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QToolBar_mouseGrabber)
	return bindQWidget(__r)
}

func QToolBar_KeyboardGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QToolBar_keyboardGrabber)
	return bindQWidget(__r)
}

func QToolBar_Find(arg1 uint64) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QToolBar_find, uintptr(arg1))
	return bindQWidget(__r)
}

func QToolBar_CreateWindowContainer(window *QWindow) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QToolBar_createWindowContainer, uintptr(window.UnsafePointer()))
	runtime.KeepAlive(window)
	return bindQWidget(__r)
}

func (this *QToolBar) UpdateMicroFocus() {
	purego.SyscallN(addr_QToolBar_updateMicroFocus, this.h)
	runtime.KeepAlive(this)
}

func (this *QToolBar) Create() {
	purego.SyscallN(addr_QToolBar_create, this.h)
	runtime.KeepAlive(this)
}

func (this *QToolBar) Destroy() {
	purego.SyscallN(addr_QToolBar_destroy, this.h)
	runtime.KeepAlive(this)
}

func (this *QToolBar) FocusNextChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QToolBar_focusNextChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QToolBar) FocusPreviousChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QToolBar_focusPreviousChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QToolBar_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QToolBar_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QToolBar_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QToolBar_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QToolBar) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QToolBar_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QToolBar) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QToolBar_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QToolBar) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QToolBar_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QToolBar) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QToolBar_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QToolBar_DevicePixelRatioFScale() float64 {
	__r := pg_QToolBar_devicePixelRatioFScale()
	return __r
}

func QToolBar_EncodeMetricF(metric QPaintDevice__PaintDeviceMetric, value float64) int {
	__r := pg_QToolBar_encodeMetricF(int32(metric), value)
	return int(__r)
}

func (this *QToolBar) GetDecodedMetricF(metricA QPaintDevice__PaintDeviceMetric, metricB QPaintDevice__PaintDeviceMetric) float64 {
	__r := pg_QToolBar_getDecodedMetricF(this.h, int32(metricA), int32(metricB))
	runtime.KeepAlive(this)
	return __r
}

func QToolBar_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QToolBar_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QToolBar_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QToolBar_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QToolBar_CreateWindowContainer2(window *QWindow, parent *QWidget) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QToolBar_createWindowContainer2, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func QToolBar_CreateWindowContainer3(window *QWindow, parent *QWidget, flags WindowType) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QToolBar_createWindowContainer3, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()), uintptr(flags))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func (this *QToolBar) UpdateMicroFocusWithQuery(query InputMethodQuery) {
	purego.SyscallN(addr_QToolBar_updateMicroFocus2, this.h, uintptr(query))
	runtime.KeepAlive(this)
}

func (this *QToolBar) CreateWith(arg1 uint64) {
	purego.SyscallN(addr_QToolBar_create2, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QToolBar) Create2(arg1 uint64, initializeWindow bool) {
	purego.SyscallN(addr_QToolBar_create3, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)))
	runtime.KeepAlive(this)
}

func (this *QToolBar) Create3(arg1 uint64, initializeWindow bool, destroyOldWindow bool) {
	purego.SyscallN(addr_QToolBar_create4, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)), uintptr(b2i(destroyOldWindow)))
	runtime.KeepAlive(this)
}

func (this *QToolBar) DestroyWithDestroyWindow(destroyWindow bool) {
	purego.SyscallN(addr_QToolBar_destroy2, this.h, uintptr(b2i(destroyWindow)))
	runtime.KeepAlive(this)
}

func (this *QToolBar) Destroy2(destroyWindow bool, destroySubWindows bool) {
	purego.SyscallN(addr_QToolBar_destroy3, this.h, uintptr(b2i(destroyWindow)), uintptr(b2i(destroySubWindows)))
	runtime.KeepAlive(this)
}

func QToolBar_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QToolBar_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QToolBar_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QToolBar_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QToolBar) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QToolBar_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QToolBar) ActionEventDefault(event *QActionEvent) {
	purego.SyscallN(addr_QToolBar_actionEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QToolBar) ChangeEventDefault(event *QEvent) {
	purego.SyscallN(addr_QToolBar_changeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QToolBar) PaintEventDefault(event *QPaintEvent) {
	purego.SyscallN(addr_QToolBar_paintEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QToolBar) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QToolBar_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QToolBar) InitStyleOptionDefault(option *QStyleOptionToolBar) {
	purego.SyscallN(addr_QToolBar_initStyleOption_default, this.h, uintptr(option.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(option)
}

func (this *QToolBar) DevTypeDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QToolBar_devType_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QToolBar) SetVisibleDefault(visible bool) {
	purego.SyscallN(addr_QToolBar_setVisible_default, this.h, uintptr(b2i(visible)))
	runtime.KeepAlive(this)
}

func (this *QToolBar) SizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QToolBar_sizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QToolBar) MinimumSizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QToolBar_minimumSizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QToolBar) HeightForWidthDefault(arg1 int) int {
	__r0, _, _ := purego.SyscallN(addr_QToolBar_heightForWidth_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QToolBar) HasHeightForWidthDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QToolBar_hasHeightForWidth_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QToolBar) PaintEngineDefault() *QPaintEngine {
	__r, _, _ := purego.SyscallN(addr_QToolBar_paintEngine_default, this.h)
	runtime.KeepAlive(this)
	return bindQPaintEngine(__r)
}

func (this *QToolBar) MousePressEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QToolBar_mousePressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QToolBar) MouseReleaseEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QToolBar_mouseReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QToolBar) MouseDoubleClickEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QToolBar_mouseDoubleClickEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QToolBar) MouseMoveEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QToolBar_mouseMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QToolBar) WheelEventDefault(event *QWheelEvent) {
	purego.SyscallN(addr_QToolBar_wheelEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QToolBar) KeyPressEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QToolBar_keyPressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QToolBar) KeyReleaseEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QToolBar_keyReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QToolBar) FocusInEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QToolBar_focusInEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QToolBar) FocusOutEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QToolBar_focusOutEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QToolBar) EnterEventDefault(event *QEnterEvent) {
	purego.SyscallN(addr_QToolBar_enterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QToolBar) LeaveEventDefault(event *QEvent) {
	purego.SyscallN(addr_QToolBar_leaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QToolBar) MoveEventDefault(event *QMoveEvent) {
	purego.SyscallN(addr_QToolBar_moveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QToolBar) ResizeEventDefault(event *QResizeEvent) {
	purego.SyscallN(addr_QToolBar_resizeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QToolBar) CloseEventDefault(event *QCloseEvent) {
	purego.SyscallN(addr_QToolBar_closeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QToolBar) ContextMenuEventDefault(event *QContextMenuEvent) {
	purego.SyscallN(addr_QToolBar_contextMenuEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QToolBar) TabletEventDefault(event *QTabletEvent) {
	purego.SyscallN(addr_QToolBar_tabletEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QToolBar) DragEnterEventDefault(event *QDragEnterEvent) {
	purego.SyscallN(addr_QToolBar_dragEnterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QToolBar) DragMoveEventDefault(event *QDragMoveEvent) {
	purego.SyscallN(addr_QToolBar_dragMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QToolBar) DragLeaveEventDefault(event *QDragLeaveEvent) {
	purego.SyscallN(addr_QToolBar_dragLeaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QToolBar) DropEventDefault(event *QDropEvent) {
	purego.SyscallN(addr_QToolBar_dropEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QToolBar) ShowEventDefault(event *QShowEvent) {
	purego.SyscallN(addr_QToolBar_showEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QToolBar) HideEventDefault(event *QHideEvent) {
	purego.SyscallN(addr_QToolBar_hideEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QToolBar) MetricDefault(arg1 QPaintDevice__PaintDeviceMetric) int {
	__r0, _, _ := purego.SyscallN(addr_QToolBar_metric_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QToolBar) InitPainterDefault(painter *QPainter) {
	purego.SyscallN(addr_QToolBar_initPainter_default, this.h, uintptr(painter.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
}

func (this *QToolBar) RedirectedDefault(offset *QPoint) *QPaintDevice {
	__r, _, _ := purego.SyscallN(addr_QToolBar_redirected_default, this.h, uintptr(offset.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(offset)
	return bindQPaintDevice(__r)
}

func (this *QToolBar) SharedPainterDefault() *QPainter {
	__r, _, _ := purego.SyscallN(addr_QToolBar_sharedPainter_default, this.h)
	runtime.KeepAlive(this)
	return bindQPainter(__r)
}

func (this *QToolBar) InputMethodEventDefault(arg1 *QInputMethodEvent) {
	purego.SyscallN(addr_QToolBar_inputMethodEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QToolBar) InputMethodQueryDefault(arg1 InputMethodQuery) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QToolBar_inputMethodQuery_default, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QToolBar) FocusNextPrevChildDefault(next bool) bool {
	__r0, _, _ := purego.SyscallN(addr_QToolBar_focusNextPrevChild_default, this.h, uintptr(b2i(next)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QToolBar) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QToolBar_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QToolBar) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QToolBar_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QToolBar) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QToolBar_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QToolBar) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QToolBar_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QToolBar) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QToolBar_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QToolBar) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QToolBar_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QToolBar) OnActionTriggered(fn func(*QAction)) int64 {
	id := registerCallback(func(action uintptr) { fn(UnsafeNewQAction(unsafe.Pointer(action))) })
	pg_QToolBar_connect_actionTriggered(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QToolBar) OnMovableChanged(fn func(bool)) int64 {
	id := registerCallback(func(movable int32) { fn(movable != 0) })
	pg_QToolBar_connect_movableChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QToolBar) OnAllowedAreasChanged(fn func(ToolBarArea)) int64 {
	id := registerCallback(func(allowedAreas int32) { fn(ToolBarArea(allowedAreas)) })
	pg_QToolBar_connect_allowedAreasChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QToolBar) OnOrientationChanged(fn func(Orientation)) int64 {
	id := registerCallback(func(orientation int32) { fn(Orientation(orientation)) })
	pg_QToolBar_connect_orientationChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QToolBar) OnIconSizeChanged(fn func(*QSize)) int64 {
	id := registerCallback(func(iconSize uintptr) { fn(UnsafeNewQSize(unsafe.Pointer(iconSize))) })
	pg_QToolBar_connect_iconSizeChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QToolBar) OnToolButtonStyleChanged(fn func(ToolButtonStyle)) int64 {
	id := registerCallback(func(toolButtonStyle int32) { fn(ToolButtonStyle(toolButtonStyle)) })
	pg_QToolBar_connect_toolButtonStyleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QToolBar) OnTopLevelChanged(fn func(bool)) int64 {
	id := registerCallback(func(topLevel int32) { fn(topLevel != 0) })
	pg_QToolBar_connect_topLevelChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QToolBar) OnVisibilityChanged(fn func(bool)) int64 {
	id := registerCallback(func(visible int32) { fn(visible != 0) })
	pg_QToolBar_connect_visibilityChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QToolBar) OnWindowTitleChanged(fn func(string)) int64 {
	id := registerCallback(func(title uintptr, title_len int64) { fn(pgStrNoFree(title, title_len)) })
	pg_QToolBar_connect_windowTitleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QToolBar) OnWindowIconChanged(fn func(*QIcon)) int64 {
	id := registerCallback(func(icon uintptr) { fn(UnsafeNewQIcon(unsafe.Pointer(icon))) })
	pg_QToolBar_connect_windowIconChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QToolBar) OnWindowIconTextChanged(fn func(string)) int64 {
	id := registerCallback(func(iconText uintptr, iconText_len int64) { fn(pgStrNoFree(iconText, iconText_len)) })
	pg_QToolBar_connect_windowIconTextChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QToolBar) OnCustomContextMenuRequested(fn func(*QPoint)) int64 {
	id := registerCallback(func(pos uintptr) { fn(UnsafeNewQPoint(unsafe.Pointer(pos))) })
	pg_QToolBar_connect_customContextMenuRequested(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QToolBar) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QToolBar_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QToolBar) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QToolBar_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QToolBar) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QToolBar_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QToolBar) Disconnect(handle int64) {
	pg_QToolBar__disconnect(handle)
	unregisterCallback(handle)
}

func QToolBar_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case 1697444140088555581:
		return 1
	case 4959376063826165943:
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
	case -1989937050019464783:
		return 26
	case -6698043222227006015:
		return 27
	case -212068866860701141:
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

func QToolBar_enable(h uintptr, idx int32) { pg_QToolBar__set_override(h, idx) }

func (this *QToolBar) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QToolBar) OnActionEvent(fn func(*QActionEvent)) {
	registerVirt(this.UnsafeGoID(), 1697444140088555581, func(event uintptr) { fn(UnsafeNewQActionEvent(unsafe.Pointer(event))) }) // slot "actionEvent"
}

func (this *QToolBar) OnChangeEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 4959376063826165943, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "changeEvent"
}

func (this *QToolBar) OnPaintEvent(fn func(*QPaintEvent)) {
	registerVirt(this.UnsafeGoID(), 5290442163167458465, func(event uintptr) { fn(UnsafeNewQPaintEvent(unsafe.Pointer(event))) }) // slot "paintEvent"
}

func (this *QToolBar) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QToolBar) OnInitStyleOption(fn func(*QStyleOptionToolBar)) {
	registerVirt(this.UnsafeGoID(), -488693769771308531, func(option uintptr) { fn(UnsafeNewQStyleOptionToolBar(unsafe.Pointer(option))) }) // slot "initStyleOption"
}

func (this *QToolBar) OnDevType(fn func() int) {
	registerVirt(this.UnsafeGoID(), -8659679893868524044, func() int32 { return int32(fn()) }) // slot "devType"
}

func (this *QToolBar) OnSetVisible(fn func(bool)) {
	registerVirt(this.UnsafeGoID(), 2110182719375245575, func(visible int32) { fn(visible != 0) }) // slot "setVisible"
}

func (this *QToolBar) OnSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 729690261457684101, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sizeHint"
}

func (this *QToolBar) OnMinimumSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 7049264599837532491, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "minimumSizeHint"
}

func (this *QToolBar) OnHeightForWidth(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), 4311534060495737339, func(arg1 int32) int32 { return int32(fn(int(arg1))) }) // slot "heightForWidth"
}

func (this *QToolBar) OnHasHeightForWidth(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -8266762397995423915, func() int32 { return b2i(fn()) }) // slot "hasHeightForWidth"
}

func (this *QToolBar) OnPaintEngine(fn func() *QPaintEngine) {
	registerVirt(this.UnsafeGoID(), 6398007404756346381, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "paintEngine"
}

func (this *QToolBar) OnMousePressEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -332090898078543319, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mousePressEvent"
}

func (this *QToolBar) OnMouseReleaseEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4415920049297795213, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseReleaseEvent"
}

func (this *QToolBar) OnMouseDoubleClickEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4029073905697023887, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseDoubleClickEvent"
}

func (this *QToolBar) OnMouseMoveEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4255494004099806581, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseMoveEvent"
}

func (this *QToolBar) OnWheelEvent(fn func(*QWheelEvent)) {
	registerVirt(this.UnsafeGoID(), 1919988340934511316, func(event uintptr) { fn(UnsafeNewQWheelEvent(unsafe.Pointer(event))) }) // slot "wheelEvent"
}

func (this *QToolBar) OnKeyPressEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 1023791977679793185, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyPressEvent"
}

func (this *QToolBar) OnKeyReleaseEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 772598475950050299, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyReleaseEvent"
}

func (this *QToolBar) OnFocusInEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), 975043767033552552, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusInEvent"
}

func (this *QToolBar) OnFocusOutEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), -385596010576048993, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusOutEvent"
}

func (this *QToolBar) OnEnterEvent(fn func(*QEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 6316270346668237895, func(event uintptr) { fn(UnsafeNewQEnterEvent(unsafe.Pointer(event))) }) // slot "enterEvent"
}

func (this *QToolBar) OnLeaveEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), -551070009623711640, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "leaveEvent"
}

func (this *QToolBar) OnMoveEvent(fn func(*QMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -9161949150489525396, func(event uintptr) { fn(UnsafeNewQMoveEvent(unsafe.Pointer(event))) }) // slot "moveEvent"
}

func (this *QToolBar) OnResizeEvent(fn func(*QResizeEvent)) {
	registerVirt(this.UnsafeGoID(), -4302917285655014661, func(event uintptr) { fn(UnsafeNewQResizeEvent(unsafe.Pointer(event))) }) // slot "resizeEvent"
}

func (this *QToolBar) OnCloseEvent(fn func(*QCloseEvent)) {
	registerVirt(this.UnsafeGoID(), -1989937050019464783, func(event uintptr) { fn(UnsafeNewQCloseEvent(unsafe.Pointer(event))) }) // slot "closeEvent"
}

func (this *QToolBar) OnContextMenuEvent(fn func(*QContextMenuEvent)) {
	registerVirt(this.UnsafeGoID(), -6698043222227006015, func(event uintptr) { fn(UnsafeNewQContextMenuEvent(unsafe.Pointer(event))) }) // slot "contextMenuEvent"
}

func (this *QToolBar) OnTabletEvent(fn func(*QTabletEvent)) {
	registerVirt(this.UnsafeGoID(), -212068866860701141, func(event uintptr) { fn(UnsafeNewQTabletEvent(unsafe.Pointer(event))) }) // slot "tabletEvent"
}

func (this *QToolBar) OnDragEnterEvent(fn func(*QDragEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 206122715919917589, func(event uintptr) { fn(UnsafeNewQDragEnterEvent(unsafe.Pointer(event))) }) // slot "dragEnterEvent"
}

func (this *QToolBar) OnDragMoveEvent(fn func(*QDragMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -1781545137352745978, func(event uintptr) { fn(UnsafeNewQDragMoveEvent(unsafe.Pointer(event))) }) // slot "dragMoveEvent"
}

func (this *QToolBar) OnDragLeaveEvent(fn func(*QDragLeaveEvent)) {
	registerVirt(this.UnsafeGoID(), -3139904158373725006, func(event uintptr) { fn(UnsafeNewQDragLeaveEvent(unsafe.Pointer(event))) }) // slot "dragLeaveEvent"
}

func (this *QToolBar) OnDropEvent(fn func(*QDropEvent)) {
	registerVirt(this.UnsafeGoID(), 9107058682754591716, func(event uintptr) { fn(UnsafeNewQDropEvent(unsafe.Pointer(event))) }) // slot "dropEvent"
}

func (this *QToolBar) OnShowEvent(fn func(*QShowEvent)) {
	registerVirt(this.UnsafeGoID(), -6825369174751916, func(event uintptr) { fn(UnsafeNewQShowEvent(unsafe.Pointer(event))) }) // slot "showEvent"
}

func (this *QToolBar) OnHideEvent(fn func(*QHideEvent)) {
	registerVirt(this.UnsafeGoID(), -199718707876719129, func(event uintptr) { fn(UnsafeNewQHideEvent(unsafe.Pointer(event))) }) // slot "hideEvent"
}

func (this *QToolBar) OnNativeEvent(fn func(*QByteArray, unsafe.Pointer, *int64) bool) {
	registerVirt(this.UnsafeGoID(), -2039166586246443706, func(eventType uintptr, message uintptr, result uintptr) int32 { return b2i(fn(UnsafeNewQByteArray(unsafe.Pointer(eventType)), unsafe.Pointer(message), (*int64)(unsafe.Pointer(result)))) }) // slot "nativeEvent"
}

func (this *QToolBar) OnMetric(fn func(QPaintDevice__PaintDeviceMetric) int) {
	registerVirt(this.UnsafeGoID(), 427629649267311777, func(arg1 int32) int32 { return int32(fn(QPaintDevice__PaintDeviceMetric(arg1))) }) // slot "metric"
}

func (this *QToolBar) OnInitPainter(fn func(*QPainter)) {
	registerVirt(this.UnsafeGoID(), 6477500296980050210, func(painter uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter))) }) // slot "initPainter"
}

func (this *QToolBar) OnRedirected(fn func(*QPoint) *QPaintDevice) {
	registerVirt(this.UnsafeGoID(), 4127240633310655048, func(offset uintptr) uintptr { return uintptr(fn(UnsafeNewQPoint(unsafe.Pointer(offset))).UnsafePointer()) }) // slot "redirected"
}

func (this *QToolBar) OnSharedPainter(fn func() *QPainter) {
	registerVirt(this.UnsafeGoID(), 5792771328741774991, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sharedPainter"
}

func (this *QToolBar) OnInputMethodEvent(fn func(*QInputMethodEvent)) {
	registerVirt(this.UnsafeGoID(), 5438819661942843222, func(arg1 uintptr) { fn(UnsafeNewQInputMethodEvent(unsafe.Pointer(arg1))) }) // slot "inputMethodEvent"
}

func (this *QToolBar) OnInputMethodQuery(fn func(InputMethodQuery) *QVariant) {
	registerVirt(this.UnsafeGoID(), 7345853097340709638, func(arg1 int32) uintptr { return uintptr(fn(InputMethodQuery(arg1)).UnsafePointer()) }) // slot "inputMethodQuery"
}

func (this *QToolBar) OnFocusNextPrevChild(fn func(bool) bool) {
	registerVirt(this.UnsafeGoID(), 5898328517505784191, func(next int32) int32 { return b2i(fn(next != 0)) }) // slot "focusNextPrevChild"
}

func (this *QToolBar) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QToolBar) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QToolBar) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QToolBar) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QToolBar) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QToolBar) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QToolBar) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QToolBar_delete(this.h)
	}
	this.UnsafeRelease()
}
