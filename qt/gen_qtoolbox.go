package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QToolBox_chain uintptr
	addr_QToolBox_new uintptr
	addr_QToolBox_new2 uintptr
	addr_QToolBox_new3 uintptr
	addr_QToolBox_tr uintptr
	addr_QToolBox_addItem uintptr
	addr_QToolBox_addItem2 uintptr
	addr_QToolBox_insertItem uintptr
	addr_QToolBox_insertItem2 uintptr
	addr_QToolBox_removeItem uintptr
	addr_QToolBox_setItemEnabled uintptr
	addr_QToolBox_isItemEnabled uintptr
	addr_QToolBox_setItemText uintptr
	addr_QToolBox_itemText uintptr
	addr_QToolBox_setItemIcon uintptr
	addr_QToolBox_itemIcon uintptr
	addr_QToolBox_setItemToolTip uintptr
	addr_QToolBox_itemToolTip uintptr
	addr_QToolBox_currentIndex uintptr
	addr_QToolBox_currentWidget uintptr
	addr_QToolBox_widget uintptr
	addr_QToolBox_indexOf uintptr
	addr_QToolBox_count uintptr
	addr_QToolBox_setCurrentIndex uintptr
	addr_QToolBox_setCurrentWidget uintptr
	addr_QToolBox_currentChanged uintptr
	addr_QToolBox_drawFrame uintptr
	addr_QToolBox_setTabOrder uintptr
	addr_QToolBox_mouseGrabber uintptr
	addr_QToolBox_keyboardGrabber uintptr
	addr_QToolBox_find uintptr
	addr_QToolBox_createWindowContainer uintptr
	addr_QToolBox_updateMicroFocus uintptr
	addr_QToolBox_create uintptr
	addr_QToolBox_destroy uintptr
	addr_QToolBox_focusNextChild uintptr
	addr_QToolBox_focusPreviousChild uintptr
	addr_QToolBox_connect uintptr
	addr_QToolBox_connect2 uintptr
	addr_QToolBox_sender uintptr
	addr_QToolBox_senderSignalIndex uintptr
	addr_QToolBox_receivers uintptr
	addr_QToolBox_isSignalConnected uintptr
	pg_QToolBox_devicePixelRatioFScale func() float64
	pg_QToolBox_encodeMetricF func(int32, float64) int32
	pg_QToolBox_getDecodedMetricF func(uintptr, int32, int32) float64
	addr_QToolBox_tr2 uintptr
	addr_QToolBox_tr3 uintptr
	addr_QToolBox_createWindowContainer2 uintptr
	addr_QToolBox_createWindowContainer3 uintptr
	addr_QToolBox_updateMicroFocus2 uintptr
	addr_QToolBox_create2 uintptr
	addr_QToolBox_create3 uintptr
	addr_QToolBox_create4 uintptr
	addr_QToolBox_destroy2 uintptr
	addr_QToolBox_destroy3 uintptr
	addr_QToolBox_connect3 uintptr
	addr_QToolBox_connect4 uintptr
	addr_QToolBox_metaObject_default uintptr
	addr_QToolBox_event_default uintptr
	addr_QToolBox_itemInserted_default uintptr
	addr_QToolBox_itemRemoved_default uintptr
	addr_QToolBox_showEvent_default uintptr
	addr_QToolBox_changeEvent_default uintptr
	addr_QToolBox_sizeHint_default uintptr
	addr_QToolBox_paintEvent_default uintptr
	addr_QToolBox_initStyleOption_default uintptr
	addr_QToolBox_devType_default uintptr
	addr_QToolBox_setVisible_default uintptr
	addr_QToolBox_minimumSizeHint_default uintptr
	addr_QToolBox_heightForWidth_default uintptr
	addr_QToolBox_hasHeightForWidth_default uintptr
	addr_QToolBox_paintEngine_default uintptr
	addr_QToolBox_mousePressEvent_default uintptr
	addr_QToolBox_mouseReleaseEvent_default uintptr
	addr_QToolBox_mouseDoubleClickEvent_default uintptr
	addr_QToolBox_mouseMoveEvent_default uintptr
	addr_QToolBox_wheelEvent_default uintptr
	addr_QToolBox_keyPressEvent_default uintptr
	addr_QToolBox_keyReleaseEvent_default uintptr
	addr_QToolBox_focusInEvent_default uintptr
	addr_QToolBox_focusOutEvent_default uintptr
	addr_QToolBox_enterEvent_default uintptr
	addr_QToolBox_leaveEvent_default uintptr
	addr_QToolBox_moveEvent_default uintptr
	addr_QToolBox_resizeEvent_default uintptr
	addr_QToolBox_closeEvent_default uintptr
	addr_QToolBox_contextMenuEvent_default uintptr
	addr_QToolBox_tabletEvent_default uintptr
	addr_QToolBox_actionEvent_default uintptr
	addr_QToolBox_dragEnterEvent_default uintptr
	addr_QToolBox_dragMoveEvent_default uintptr
	addr_QToolBox_dragLeaveEvent_default uintptr
	addr_QToolBox_dropEvent_default uintptr
	addr_QToolBox_hideEvent_default uintptr
	addr_QToolBox_metric_default uintptr
	addr_QToolBox_initPainter_default uintptr
	addr_QToolBox_redirected_default uintptr
	addr_QToolBox_sharedPainter_default uintptr
	addr_QToolBox_inputMethodEvent_default uintptr
	addr_QToolBox_inputMethodQuery_default uintptr
	addr_QToolBox_focusNextPrevChild_default uintptr
	addr_QToolBox_eventFilter_default uintptr
	addr_QToolBox_timerEvent_default uintptr
	addr_QToolBox_childEvent_default uintptr
	addr_QToolBox_customEvent_default uintptr
	addr_QToolBox_connectNotify_default uintptr
	addr_QToolBox_disconnectNotify_default uintptr
	pg_QToolBox_connect_currentChanged func(uintptr, int64)
	pg_QToolBox_connect_windowTitleChanged func(uintptr, int64)
	pg_QToolBox_connect_windowIconChanged func(uintptr, int64)
	pg_QToolBox_connect_windowIconTextChanged func(uintptr, int64)
	pg_QToolBox_connect_customContextMenuRequested func(uintptr, int64)
	pg_QToolBox_connect_destroyed func(uintptr, int64)
	pg_QToolBox_connect_destroyed2 func(uintptr, int64)
	pg_QToolBox_connect_objectNameChanged func(uintptr, int64)
	pg_QToolBox__disconnect func(int64)
	pg_QToolBox__set_override func(uintptr, int32)
	pg_QToolBox_delete func(uintptr)
)

func init() {
	symaddr(&addr_QToolBox_chain, "QToolBox_chain")
	symaddr(&addr_QToolBox_new, "QToolBox_new")
	symaddr(&addr_QToolBox_new2, "QToolBox_new2")
	symaddr(&addr_QToolBox_new3, "QToolBox_new3")
	symaddr(&addr_QToolBox_tr, "QToolBox_tr")
	symaddr(&addr_QToolBox_addItem, "QToolBox_addItem")
	symaddr(&addr_QToolBox_addItem2, "QToolBox_addItem2")
	symaddr(&addr_QToolBox_insertItem, "QToolBox_insertItem")
	symaddr(&addr_QToolBox_insertItem2, "QToolBox_insertItem2")
	symaddr(&addr_QToolBox_removeItem, "QToolBox_removeItem")
	symaddr(&addr_QToolBox_setItemEnabled, "QToolBox_setItemEnabled")
	symaddr(&addr_QToolBox_isItemEnabled, "QToolBox_isItemEnabled")
	symaddr(&addr_QToolBox_setItemText, "QToolBox_setItemText")
	symaddr(&addr_QToolBox_itemText, "QToolBox_itemText")
	symaddr(&addr_QToolBox_setItemIcon, "QToolBox_setItemIcon")
	symaddr(&addr_QToolBox_itemIcon, "QToolBox_itemIcon")
	symaddr(&addr_QToolBox_setItemToolTip, "QToolBox_setItemToolTip")
	symaddr(&addr_QToolBox_itemToolTip, "QToolBox_itemToolTip")
	symaddr(&addr_QToolBox_currentIndex, "QToolBox_currentIndex")
	symaddr(&addr_QToolBox_currentWidget, "QToolBox_currentWidget")
	symaddr(&addr_QToolBox_widget, "QToolBox_widget")
	symaddr(&addr_QToolBox_indexOf, "QToolBox_indexOf")
	symaddr(&addr_QToolBox_count, "QToolBox_count")
	symaddr(&addr_QToolBox_setCurrentIndex, "QToolBox_setCurrentIndex")
	symaddr(&addr_QToolBox_setCurrentWidget, "QToolBox_setCurrentWidget")
	symaddr(&addr_QToolBox_currentChanged, "QToolBox_currentChanged")
	symaddr(&addr_QToolBox_drawFrame, "QToolBox_drawFrame")
	symaddr(&addr_QToolBox_setTabOrder, "QToolBox_setTabOrder")
	symaddr(&addr_QToolBox_mouseGrabber, "QToolBox_mouseGrabber")
	symaddr(&addr_QToolBox_keyboardGrabber, "QToolBox_keyboardGrabber")
	symaddr(&addr_QToolBox_find, "QToolBox_find")
	symaddr(&addr_QToolBox_createWindowContainer, "QToolBox_createWindowContainer")
	symaddr(&addr_QToolBox_updateMicroFocus, "QToolBox_updateMicroFocus")
	symaddr(&addr_QToolBox_create, "QToolBox_create")
	symaddr(&addr_QToolBox_destroy, "QToolBox_destroy")
	symaddr(&addr_QToolBox_focusNextChild, "QToolBox_focusNextChild")
	symaddr(&addr_QToolBox_focusPreviousChild, "QToolBox_focusPreviousChild")
	symaddr(&addr_QToolBox_connect, "QToolBox_connect")
	symaddr(&addr_QToolBox_connect2, "QToolBox_connect2")
	symaddr(&addr_QToolBox_sender, "QToolBox_sender")
	symaddr(&addr_QToolBox_senderSignalIndex, "QToolBox_senderSignalIndex")
	symaddr(&addr_QToolBox_receivers, "QToolBox_receivers")
	symaddr(&addr_QToolBox_isSignalConnected, "QToolBox_isSignalConnected")
	reg(&pg_QToolBox_devicePixelRatioFScale, "QToolBox_devicePixelRatioFScale")
	reg(&pg_QToolBox_encodeMetricF, "QToolBox_encodeMetricF")
	reg(&pg_QToolBox_getDecodedMetricF, "QToolBox_getDecodedMetricF")
	symaddr(&addr_QToolBox_tr2, "QToolBox_tr2")
	symaddr(&addr_QToolBox_tr3, "QToolBox_tr3")
	symaddr(&addr_QToolBox_createWindowContainer2, "QToolBox_createWindowContainer2")
	symaddr(&addr_QToolBox_createWindowContainer3, "QToolBox_createWindowContainer3")
	symaddr(&addr_QToolBox_updateMicroFocus2, "QToolBox_updateMicroFocus2")
	symaddr(&addr_QToolBox_create2, "QToolBox_create2")
	symaddr(&addr_QToolBox_create3, "QToolBox_create3")
	symaddr(&addr_QToolBox_create4, "QToolBox_create4")
	symaddr(&addr_QToolBox_destroy2, "QToolBox_destroy2")
	symaddr(&addr_QToolBox_destroy3, "QToolBox_destroy3")
	symaddr(&addr_QToolBox_connect3, "QToolBox_connect3")
	symaddr(&addr_QToolBox_connect4, "QToolBox_connect4")
	symaddr(&addr_QToolBox_metaObject_default, "QToolBox_metaObject_default")
	symaddr(&addr_QToolBox_event_default, "QToolBox_event_default")
	symaddr(&addr_QToolBox_itemInserted_default, "QToolBox_itemInserted_default")
	symaddr(&addr_QToolBox_itemRemoved_default, "QToolBox_itemRemoved_default")
	symaddr(&addr_QToolBox_showEvent_default, "QToolBox_showEvent_default")
	symaddr(&addr_QToolBox_changeEvent_default, "QToolBox_changeEvent_default")
	symaddr(&addr_QToolBox_sizeHint_default, "QToolBox_sizeHint_default")
	symaddr(&addr_QToolBox_paintEvent_default, "QToolBox_paintEvent_default")
	symaddr(&addr_QToolBox_initStyleOption_default, "QToolBox_initStyleOption_default")
	symaddr(&addr_QToolBox_devType_default, "QToolBox_devType_default")
	symaddr(&addr_QToolBox_setVisible_default, "QToolBox_setVisible_default")
	symaddr(&addr_QToolBox_minimumSizeHint_default, "QToolBox_minimumSizeHint_default")
	symaddr(&addr_QToolBox_heightForWidth_default, "QToolBox_heightForWidth_default")
	symaddr(&addr_QToolBox_hasHeightForWidth_default, "QToolBox_hasHeightForWidth_default")
	symaddr(&addr_QToolBox_paintEngine_default, "QToolBox_paintEngine_default")
	symaddr(&addr_QToolBox_mousePressEvent_default, "QToolBox_mousePressEvent_default")
	symaddr(&addr_QToolBox_mouseReleaseEvent_default, "QToolBox_mouseReleaseEvent_default")
	symaddr(&addr_QToolBox_mouseDoubleClickEvent_default, "QToolBox_mouseDoubleClickEvent_default")
	symaddr(&addr_QToolBox_mouseMoveEvent_default, "QToolBox_mouseMoveEvent_default")
	symaddr(&addr_QToolBox_wheelEvent_default, "QToolBox_wheelEvent_default")
	symaddr(&addr_QToolBox_keyPressEvent_default, "QToolBox_keyPressEvent_default")
	symaddr(&addr_QToolBox_keyReleaseEvent_default, "QToolBox_keyReleaseEvent_default")
	symaddr(&addr_QToolBox_focusInEvent_default, "QToolBox_focusInEvent_default")
	symaddr(&addr_QToolBox_focusOutEvent_default, "QToolBox_focusOutEvent_default")
	symaddr(&addr_QToolBox_enterEvent_default, "QToolBox_enterEvent_default")
	symaddr(&addr_QToolBox_leaveEvent_default, "QToolBox_leaveEvent_default")
	symaddr(&addr_QToolBox_moveEvent_default, "QToolBox_moveEvent_default")
	symaddr(&addr_QToolBox_resizeEvent_default, "QToolBox_resizeEvent_default")
	symaddr(&addr_QToolBox_closeEvent_default, "QToolBox_closeEvent_default")
	symaddr(&addr_QToolBox_contextMenuEvent_default, "QToolBox_contextMenuEvent_default")
	symaddr(&addr_QToolBox_tabletEvent_default, "QToolBox_tabletEvent_default")
	symaddr(&addr_QToolBox_actionEvent_default, "QToolBox_actionEvent_default")
	symaddr(&addr_QToolBox_dragEnterEvent_default, "QToolBox_dragEnterEvent_default")
	symaddr(&addr_QToolBox_dragMoveEvent_default, "QToolBox_dragMoveEvent_default")
	symaddr(&addr_QToolBox_dragLeaveEvent_default, "QToolBox_dragLeaveEvent_default")
	symaddr(&addr_QToolBox_dropEvent_default, "QToolBox_dropEvent_default")
	symaddr(&addr_QToolBox_hideEvent_default, "QToolBox_hideEvent_default")
	symaddr(&addr_QToolBox_metric_default, "QToolBox_metric_default")
	symaddr(&addr_QToolBox_initPainter_default, "QToolBox_initPainter_default")
	symaddr(&addr_QToolBox_redirected_default, "QToolBox_redirected_default")
	symaddr(&addr_QToolBox_sharedPainter_default, "QToolBox_sharedPainter_default")
	symaddr(&addr_QToolBox_inputMethodEvent_default, "QToolBox_inputMethodEvent_default")
	symaddr(&addr_QToolBox_inputMethodQuery_default, "QToolBox_inputMethodQuery_default")
	symaddr(&addr_QToolBox_focusNextPrevChild_default, "QToolBox_focusNextPrevChild_default")
	symaddr(&addr_QToolBox_eventFilter_default, "QToolBox_eventFilter_default")
	symaddr(&addr_QToolBox_timerEvent_default, "QToolBox_timerEvent_default")
	symaddr(&addr_QToolBox_childEvent_default, "QToolBox_childEvent_default")
	symaddr(&addr_QToolBox_customEvent_default, "QToolBox_customEvent_default")
	symaddr(&addr_QToolBox_connectNotify_default, "QToolBox_connectNotify_default")
	symaddr(&addr_QToolBox_disconnectNotify_default, "QToolBox_disconnectNotify_default")
	reg(&pg_QToolBox_connect_currentChanged, "QToolBox_connect_currentChanged")
	reg(&pg_QToolBox_connect_windowTitleChanged, "QToolBox_connect_windowTitleChanged")
	reg(&pg_QToolBox_connect_windowIconChanged, "QToolBox_connect_windowIconChanged")
	reg(&pg_QToolBox_connect_windowIconTextChanged, "QToolBox_connect_windowIconTextChanged")
	reg(&pg_QToolBox_connect_customContextMenuRequested, "QToolBox_connect_customContextMenuRequested")
	reg(&pg_QToolBox_connect_destroyed, "QToolBox_connect_destroyed")
	reg(&pg_QToolBox_connect_destroyed2, "QToolBox_connect_destroyed2")
	reg(&pg_QToolBox_connect_objectNameChanged, "QToolBox_connect_objectNameChanged")
	reg(&pg_QToolBox__disconnect, "QToolBox__disconnect")
	reg(&pg_QToolBox__set_override, "QToolBox__set_override")
	reg(&pg_QToolBox_delete, "QToolBox_delete")
}

type QToolBox struct {
	h uintptr
	*QFrame
}

func bindQToolBox(h uintptr) *QToolBox {
	if h == 0 {
		return nil
	}
	var __c [3]uintptr
	purego.SyscallN(addr_QToolBox_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QToolBox{h: h, QFrame: &QFrame{h: __c[0], QWidget: &QWidget{h: __c[1], QObject: &QObject{h: __c[2]}}}}
}

func newownedQToolBox(h uintptr) *QToolBox { return bindQToolBox(h) }

func (this *QToolBox) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QToolBox) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQToolBox(h unsafe.Pointer) *QToolBox { return bindQToolBox(uintptr(h)) }
func UnsafeNewQToolBoxOwned(h unsafe.Pointer) *QToolBox { return newownedQToolBox(uintptr(h)) }

func (this *QToolBox) UnsafeRelease() {
	this.h = 0
	if this.QFrame != nil {
		this.QFrame.UnsafeRelease()
	}
}

func NewQToolBox(parent *QWidget) *QToolBox {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QToolBox_new, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQToolBox(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QToolBox_virtIdx, QToolBox_enable)
	return __o
}

func NewQToolBox2() *QToolBox {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QToolBox_new2, uintptr(__goid))
	__o := newownedQToolBox(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QToolBox_virtIdx, QToolBox_enable)
	return __o
}

func NewQToolBox3(parent *QWidget, f WindowType) *QToolBox {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QToolBox_new3, uintptr(__goid), uintptr(parent.UnsafePointer()), uintptr(f))
	__o := newownedQToolBox(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QToolBox_virtIdx, QToolBox_enable)
	return __o
}

func QToolBox_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QToolBox_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QToolBox) AddItem(widget *QWidget, text string) int {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r0, _, _ := purego.SyscallN(addr_QToolBox_addItem, this.h, uintptr(widget.UnsafePointer()), text_p, uintptr(len(text)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(widget)
	runtime.KeepAlive(text)
	return int(__r)
}

func (this *QToolBox) AddItem2(widget *QWidget, icon *QIcon, text string) int {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r0, _, _ := purego.SyscallN(addr_QToolBox_addItem2, this.h, uintptr(widget.UnsafePointer()), uintptr(icon.UnsafePointer()), text_p, uintptr(len(text)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(widget)
	runtime.KeepAlive(icon)
	runtime.KeepAlive(text)
	return int(__r)
}

func (this *QToolBox) InsertItem(index int, widget *QWidget, text string) int {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r0, _, _ := purego.SyscallN(addr_QToolBox_insertItem, this.h, uintptr(index), uintptr(widget.UnsafePointer()), text_p, uintptr(len(text)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(widget)
	runtime.KeepAlive(text)
	return int(__r)
}

func (this *QToolBox) InsertItem2(index int, widget *QWidget, icon *QIcon, text string) int {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r0, _, _ := purego.SyscallN(addr_QToolBox_insertItem2, this.h, uintptr(index), uintptr(widget.UnsafePointer()), uintptr(icon.UnsafePointer()), text_p, uintptr(len(text)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(widget)
	runtime.KeepAlive(icon)
	runtime.KeepAlive(text)
	return int(__r)
}

func (this *QToolBox) RemoveItem(index int) {
	purego.SyscallN(addr_QToolBox_removeItem, this.h, uintptr(index))
	runtime.KeepAlive(this)
}

func (this *QToolBox) SetItemEnabled(index int, enabled bool) {
	purego.SyscallN(addr_QToolBox_setItemEnabled, this.h, uintptr(index), uintptr(b2i(enabled)))
	runtime.KeepAlive(this)
}

func (this *QToolBox) IsItemEnabled(index int) bool {
	__r0, _, _ := purego.SyscallN(addr_QToolBox_isItemEnabled, this.h, uintptr(index))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QToolBox) SetItemText(index int, text string) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QToolBox_setItemText, this.h, uintptr(index), text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
}

func (this *QToolBox) ItemText(index int) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QToolBox_itemText, this.h, uintptr(index), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QToolBox) SetItemIcon(index int, icon *QIcon) {
	purego.SyscallN(addr_QToolBox_setItemIcon, this.h, uintptr(index), uintptr(icon.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(icon)
}

func (this *QToolBox) ItemIcon(index int) *QIcon {
	__r, _, _ := purego.SyscallN(addr_QToolBox_itemIcon, this.h, uintptr(index))
	runtime.KeepAlive(this)
	return newownedQIcon(__r)
}

func (this *QToolBox) SetItemToolTip(index int, toolTip string) {
	var toolTip_p uintptr
	if len(toolTip) > 0 {
		toolTip_p = uintptr(unsafe.Pointer(unsafe.StringData(toolTip)))
	}
	purego.SyscallN(addr_QToolBox_setItemToolTip, this.h, uintptr(index), toolTip_p, uintptr(len(toolTip)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(toolTip)
}

func (this *QToolBox) ItemToolTip(index int) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QToolBox_itemToolTip, this.h, uintptr(index), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QToolBox) CurrentIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QToolBox_currentIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QToolBox) CurrentWidget() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QToolBox_currentWidget, this.h)
	runtime.KeepAlive(this)
	return bindQWidget(__r)
}

func (this *QToolBox) Widget(index int) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QToolBox_widget, this.h, uintptr(index))
	runtime.KeepAlive(this)
	return bindQWidget(__r)
}

func (this *QToolBox) IndexOf(widget *QWidget) int {
	__r0, _, _ := purego.SyscallN(addr_QToolBox_indexOf, this.h, uintptr(widget.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(widget)
	return int(__r)
}

func (this *QToolBox) Count() int {
	__r0, _, _ := purego.SyscallN(addr_QToolBox_count, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QToolBox) SetCurrentIndex(index int) {
	purego.SyscallN(addr_QToolBox_setCurrentIndex, this.h, uintptr(index))
	runtime.KeepAlive(this)
}

func (this *QToolBox) SetCurrentWidget(widget *QWidget) {
	purego.SyscallN(addr_QToolBox_setCurrentWidget, this.h, uintptr(widget.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(widget)
}

func (this *QToolBox) CurrentChanged(index int) {
	purego.SyscallN(addr_QToolBox_currentChanged, this.h, uintptr(index))
	runtime.KeepAlive(this)
}

func (this *QToolBox) DrawFrame(arg1 *QPainter) {
	purego.SyscallN(addr_QToolBox_drawFrame, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func QToolBox_SetTabOrder(arg1 *QWidget, arg2 *QWidget) {
	purego.SyscallN(addr_QToolBox_setTabOrder, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()))
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
}

func QToolBox_MouseGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QToolBox_mouseGrabber)
	return bindQWidget(__r)
}

func QToolBox_KeyboardGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QToolBox_keyboardGrabber)
	return bindQWidget(__r)
}

func QToolBox_Find(arg1 uint64) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QToolBox_find, uintptr(arg1))
	return bindQWidget(__r)
}

func QToolBox_CreateWindowContainer(window *QWindow) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QToolBox_createWindowContainer, uintptr(window.UnsafePointer()))
	runtime.KeepAlive(window)
	return bindQWidget(__r)
}

func (this *QToolBox) UpdateMicroFocus() {
	purego.SyscallN(addr_QToolBox_updateMicroFocus, this.h)
	runtime.KeepAlive(this)
}

func (this *QToolBox) Create() {
	purego.SyscallN(addr_QToolBox_create, this.h)
	runtime.KeepAlive(this)
}

func (this *QToolBox) Destroy() {
	purego.SyscallN(addr_QToolBox_destroy, this.h)
	runtime.KeepAlive(this)
}

func (this *QToolBox) FocusNextChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QToolBox_focusNextChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QToolBox) FocusPreviousChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QToolBox_focusPreviousChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QToolBox_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QToolBox_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QToolBox_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QToolBox_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QToolBox) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QToolBox_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QToolBox) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QToolBox_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QToolBox) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QToolBox_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QToolBox) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QToolBox_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QToolBox_DevicePixelRatioFScale() float64 {
	__r := pg_QToolBox_devicePixelRatioFScale()
	return __r
}

func QToolBox_EncodeMetricF(metric QPaintDevice__PaintDeviceMetric, value float64) int {
	__r := pg_QToolBox_encodeMetricF(int32(metric), value)
	return int(__r)
}

func (this *QToolBox) GetDecodedMetricF(metricA QPaintDevice__PaintDeviceMetric, metricB QPaintDevice__PaintDeviceMetric) float64 {
	__r := pg_QToolBox_getDecodedMetricF(this.h, int32(metricA), int32(metricB))
	runtime.KeepAlive(this)
	return __r
}

func QToolBox_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QToolBox_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QToolBox_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QToolBox_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QToolBox_CreateWindowContainer2(window *QWindow, parent *QWidget) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QToolBox_createWindowContainer2, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func QToolBox_CreateWindowContainer3(window *QWindow, parent *QWidget, flags WindowType) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QToolBox_createWindowContainer3, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()), uintptr(flags))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func (this *QToolBox) UpdateMicroFocusWithQuery(query InputMethodQuery) {
	purego.SyscallN(addr_QToolBox_updateMicroFocus2, this.h, uintptr(query))
	runtime.KeepAlive(this)
}

func (this *QToolBox) CreateWith(arg1 uint64) {
	purego.SyscallN(addr_QToolBox_create2, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QToolBox) Create2(arg1 uint64, initializeWindow bool) {
	purego.SyscallN(addr_QToolBox_create3, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)))
	runtime.KeepAlive(this)
}

func (this *QToolBox) Create3(arg1 uint64, initializeWindow bool, destroyOldWindow bool) {
	purego.SyscallN(addr_QToolBox_create4, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)), uintptr(b2i(destroyOldWindow)))
	runtime.KeepAlive(this)
}

func (this *QToolBox) DestroyWithDestroyWindow(destroyWindow bool) {
	purego.SyscallN(addr_QToolBox_destroy2, this.h, uintptr(b2i(destroyWindow)))
	runtime.KeepAlive(this)
}

func (this *QToolBox) Destroy2(destroyWindow bool, destroySubWindows bool) {
	purego.SyscallN(addr_QToolBox_destroy3, this.h, uintptr(b2i(destroyWindow)), uintptr(b2i(destroySubWindows)))
	runtime.KeepAlive(this)
}

func QToolBox_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QToolBox_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QToolBox_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QToolBox_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QToolBox) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QToolBox_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QToolBox) EventDefault(e *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QToolBox_event_default, this.h, uintptr(e.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
	return __r != 0
}

func (this *QToolBox) ItemInsertedDefault(index int) {
	purego.SyscallN(addr_QToolBox_itemInserted_default, this.h, uintptr(index))
	runtime.KeepAlive(this)
}

func (this *QToolBox) ItemRemovedDefault(index int) {
	purego.SyscallN(addr_QToolBox_itemRemoved_default, this.h, uintptr(index))
	runtime.KeepAlive(this)
}

func (this *QToolBox) ShowEventDefault(e *QShowEvent) {
	purego.SyscallN(addr_QToolBox_showEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QToolBox) ChangeEventDefault(arg1 *QEvent) {
	purego.SyscallN(addr_QToolBox_changeEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QToolBox) SizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QToolBox_sizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QToolBox) PaintEventDefault(arg1 *QPaintEvent) {
	purego.SyscallN(addr_QToolBox_paintEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QToolBox) InitStyleOptionDefault(option *QStyleOptionFrame) {
	purego.SyscallN(addr_QToolBox_initStyleOption_default, this.h, uintptr(option.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(option)
}

func (this *QToolBox) DevTypeDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QToolBox_devType_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QToolBox) SetVisibleDefault(visible bool) {
	purego.SyscallN(addr_QToolBox_setVisible_default, this.h, uintptr(b2i(visible)))
	runtime.KeepAlive(this)
}

func (this *QToolBox) MinimumSizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QToolBox_minimumSizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QToolBox) HeightForWidthDefault(arg1 int) int {
	__r0, _, _ := purego.SyscallN(addr_QToolBox_heightForWidth_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QToolBox) HasHeightForWidthDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QToolBox_hasHeightForWidth_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QToolBox) PaintEngineDefault() *QPaintEngine {
	__r, _, _ := purego.SyscallN(addr_QToolBox_paintEngine_default, this.h)
	runtime.KeepAlive(this)
	return bindQPaintEngine(__r)
}

func (this *QToolBox) MousePressEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QToolBox_mousePressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QToolBox) MouseReleaseEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QToolBox_mouseReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QToolBox) MouseDoubleClickEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QToolBox_mouseDoubleClickEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QToolBox) MouseMoveEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QToolBox_mouseMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QToolBox) WheelEventDefault(event *QWheelEvent) {
	purego.SyscallN(addr_QToolBox_wheelEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QToolBox) KeyPressEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QToolBox_keyPressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QToolBox) KeyReleaseEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QToolBox_keyReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QToolBox) FocusInEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QToolBox_focusInEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QToolBox) FocusOutEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QToolBox_focusOutEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QToolBox) EnterEventDefault(event *QEnterEvent) {
	purego.SyscallN(addr_QToolBox_enterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QToolBox) LeaveEventDefault(event *QEvent) {
	purego.SyscallN(addr_QToolBox_leaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QToolBox) MoveEventDefault(event *QMoveEvent) {
	purego.SyscallN(addr_QToolBox_moveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QToolBox) ResizeEventDefault(event *QResizeEvent) {
	purego.SyscallN(addr_QToolBox_resizeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QToolBox) CloseEventDefault(event *QCloseEvent) {
	purego.SyscallN(addr_QToolBox_closeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QToolBox) ContextMenuEventDefault(event *QContextMenuEvent) {
	purego.SyscallN(addr_QToolBox_contextMenuEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QToolBox) TabletEventDefault(event *QTabletEvent) {
	purego.SyscallN(addr_QToolBox_tabletEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QToolBox) ActionEventDefault(event *QActionEvent) {
	purego.SyscallN(addr_QToolBox_actionEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QToolBox) DragEnterEventDefault(event *QDragEnterEvent) {
	purego.SyscallN(addr_QToolBox_dragEnterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QToolBox) DragMoveEventDefault(event *QDragMoveEvent) {
	purego.SyscallN(addr_QToolBox_dragMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QToolBox) DragLeaveEventDefault(event *QDragLeaveEvent) {
	purego.SyscallN(addr_QToolBox_dragLeaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QToolBox) DropEventDefault(event *QDropEvent) {
	purego.SyscallN(addr_QToolBox_dropEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QToolBox) HideEventDefault(event *QHideEvent) {
	purego.SyscallN(addr_QToolBox_hideEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QToolBox) MetricDefault(arg1 QPaintDevice__PaintDeviceMetric) int {
	__r0, _, _ := purego.SyscallN(addr_QToolBox_metric_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QToolBox) InitPainterDefault(painter *QPainter) {
	purego.SyscallN(addr_QToolBox_initPainter_default, this.h, uintptr(painter.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
}

func (this *QToolBox) RedirectedDefault(offset *QPoint) *QPaintDevice {
	__r, _, _ := purego.SyscallN(addr_QToolBox_redirected_default, this.h, uintptr(offset.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(offset)
	return bindQPaintDevice(__r)
}

func (this *QToolBox) SharedPainterDefault() *QPainter {
	__r, _, _ := purego.SyscallN(addr_QToolBox_sharedPainter_default, this.h)
	runtime.KeepAlive(this)
	return bindQPainter(__r)
}

func (this *QToolBox) InputMethodEventDefault(arg1 *QInputMethodEvent) {
	purego.SyscallN(addr_QToolBox_inputMethodEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QToolBox) InputMethodQueryDefault(arg1 InputMethodQuery) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QToolBox_inputMethodQuery_default, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QToolBox) FocusNextPrevChildDefault(next bool) bool {
	__r0, _, _ := purego.SyscallN(addr_QToolBox_focusNextPrevChild_default, this.h, uintptr(b2i(next)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QToolBox) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QToolBox_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QToolBox) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QToolBox_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QToolBox) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QToolBox_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QToolBox) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QToolBox_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QToolBox) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QToolBox_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QToolBox) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QToolBox_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QToolBox) OnCurrentChanged(fn func(int)) int64 {
	id := registerCallback(func(index int32) { fn(int(index)) })
	pg_QToolBox_connect_currentChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QToolBox) OnWindowTitleChanged(fn func(string)) int64 {
	id := registerCallback(func(title uintptr, title_len int64) { fn(pgStrNoFree(title, title_len)) })
	pg_QToolBox_connect_windowTitleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QToolBox) OnWindowIconChanged(fn func(*QIcon)) int64 {
	id := registerCallback(func(icon uintptr) { fn(UnsafeNewQIcon(unsafe.Pointer(icon))) })
	pg_QToolBox_connect_windowIconChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QToolBox) OnWindowIconTextChanged(fn func(string)) int64 {
	id := registerCallback(func(iconText uintptr, iconText_len int64) { fn(pgStrNoFree(iconText, iconText_len)) })
	pg_QToolBox_connect_windowIconTextChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QToolBox) OnCustomContextMenuRequested(fn func(*QPoint)) int64 {
	id := registerCallback(func(pos uintptr) { fn(UnsafeNewQPoint(unsafe.Pointer(pos))) })
	pg_QToolBox_connect_customContextMenuRequested(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QToolBox) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QToolBox_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QToolBox) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QToolBox_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QToolBox) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QToolBox_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QToolBox) Disconnect(handle int64) {
	pg_QToolBox__disconnect(handle)
	unregisterCallback(handle)
}

func QToolBox_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case 1237236456312824991:
		return 1
	case -2981295881700994698:
		return 2
	case 5624127216289758258:
		return 3
	case -6825369174751916:
		return 4
	case 4959376063826165943:
		return 5
	case 729690261457684101:
		return 6
	case 5290442163167458465:
		return 7
	case -488693769771308531:
		return 8
	case -8659679893868524044:
		return 9
	case 2110182719375245575:
		return 10
	case 7049264599837532491:
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
	case -4302917285655014661:
		return 27
	case -1989937050019464783:
		return 28
	case -6698043222227006015:
		return 29
	case -212068866860701141:
		return 30
	case 1697444140088555581:
		return 31
	case 206122715919917589:
		return 32
	case -1781545137352745978:
		return 33
	case -3139904158373725006:
		return 34
	case 9107058682754591716:
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
	case -1185315367526857833:
		return 47
	case 3769868229123868474:
		return 48
	case 7379029982170503008:
		return 49
	case -568369772982215962:
		return 50
	}
	return -1
}

func QToolBox_enable(h uintptr, idx int32) { pg_QToolBox__set_override(h, idx) }

func (this *QToolBox) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QToolBox) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(e uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(e)))) }) // slot "event"
}

func (this *QToolBox) OnItemInserted(fn func(int)) {
	registerVirt(this.UnsafeGoID(), -2981295881700994698, func(index int32) { fn(int(index)) }) // slot "itemInserted"
}

func (this *QToolBox) OnItemRemoved(fn func(int)) {
	registerVirt(this.UnsafeGoID(), 5624127216289758258, func(index int32) { fn(int(index)) }) // slot "itemRemoved"
}

func (this *QToolBox) OnShowEvent(fn func(*QShowEvent)) {
	registerVirt(this.UnsafeGoID(), -6825369174751916, func(e uintptr) { fn(UnsafeNewQShowEvent(unsafe.Pointer(e))) }) // slot "showEvent"
}

func (this *QToolBox) OnChangeEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 4959376063826165943, func(arg1 uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(arg1))) }) // slot "changeEvent"
}

func (this *QToolBox) OnSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 729690261457684101, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sizeHint"
}

func (this *QToolBox) OnPaintEvent(fn func(*QPaintEvent)) {
	registerVirt(this.UnsafeGoID(), 5290442163167458465, func(arg1 uintptr) { fn(UnsafeNewQPaintEvent(unsafe.Pointer(arg1))) }) // slot "paintEvent"
}

func (this *QToolBox) OnInitStyleOption(fn func(*QStyleOptionFrame)) {
	registerVirt(this.UnsafeGoID(), -488693769771308531, func(option uintptr) { fn(UnsafeNewQStyleOptionFrame(unsafe.Pointer(option))) }) // slot "initStyleOption"
}

func (this *QToolBox) OnDevType(fn func() int) {
	registerVirt(this.UnsafeGoID(), -8659679893868524044, func() int32 { return int32(fn()) }) // slot "devType"
}

func (this *QToolBox) OnSetVisible(fn func(bool)) {
	registerVirt(this.UnsafeGoID(), 2110182719375245575, func(visible int32) { fn(visible != 0) }) // slot "setVisible"
}

func (this *QToolBox) OnMinimumSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 7049264599837532491, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "minimumSizeHint"
}

func (this *QToolBox) OnHeightForWidth(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), 4311534060495737339, func(arg1 int32) int32 { return int32(fn(int(arg1))) }) // slot "heightForWidth"
}

func (this *QToolBox) OnHasHeightForWidth(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -8266762397995423915, func() int32 { return b2i(fn()) }) // slot "hasHeightForWidth"
}

func (this *QToolBox) OnPaintEngine(fn func() *QPaintEngine) {
	registerVirt(this.UnsafeGoID(), 6398007404756346381, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "paintEngine"
}

func (this *QToolBox) OnMousePressEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -332090898078543319, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mousePressEvent"
}

func (this *QToolBox) OnMouseReleaseEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4415920049297795213, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseReleaseEvent"
}

func (this *QToolBox) OnMouseDoubleClickEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4029073905697023887, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseDoubleClickEvent"
}

func (this *QToolBox) OnMouseMoveEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4255494004099806581, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseMoveEvent"
}

func (this *QToolBox) OnWheelEvent(fn func(*QWheelEvent)) {
	registerVirt(this.UnsafeGoID(), 1919988340934511316, func(event uintptr) { fn(UnsafeNewQWheelEvent(unsafe.Pointer(event))) }) // slot "wheelEvent"
}

func (this *QToolBox) OnKeyPressEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 1023791977679793185, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyPressEvent"
}

func (this *QToolBox) OnKeyReleaseEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 772598475950050299, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyReleaseEvent"
}

func (this *QToolBox) OnFocusInEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), 975043767033552552, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusInEvent"
}

func (this *QToolBox) OnFocusOutEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), -385596010576048993, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusOutEvent"
}

func (this *QToolBox) OnEnterEvent(fn func(*QEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 6316270346668237895, func(event uintptr) { fn(UnsafeNewQEnterEvent(unsafe.Pointer(event))) }) // slot "enterEvent"
}

func (this *QToolBox) OnLeaveEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), -551070009623711640, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "leaveEvent"
}

func (this *QToolBox) OnMoveEvent(fn func(*QMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -9161949150489525396, func(event uintptr) { fn(UnsafeNewQMoveEvent(unsafe.Pointer(event))) }) // slot "moveEvent"
}

func (this *QToolBox) OnResizeEvent(fn func(*QResizeEvent)) {
	registerVirt(this.UnsafeGoID(), -4302917285655014661, func(event uintptr) { fn(UnsafeNewQResizeEvent(unsafe.Pointer(event))) }) // slot "resizeEvent"
}

func (this *QToolBox) OnCloseEvent(fn func(*QCloseEvent)) {
	registerVirt(this.UnsafeGoID(), -1989937050019464783, func(event uintptr) { fn(UnsafeNewQCloseEvent(unsafe.Pointer(event))) }) // slot "closeEvent"
}

func (this *QToolBox) OnContextMenuEvent(fn func(*QContextMenuEvent)) {
	registerVirt(this.UnsafeGoID(), -6698043222227006015, func(event uintptr) { fn(UnsafeNewQContextMenuEvent(unsafe.Pointer(event))) }) // slot "contextMenuEvent"
}

func (this *QToolBox) OnTabletEvent(fn func(*QTabletEvent)) {
	registerVirt(this.UnsafeGoID(), -212068866860701141, func(event uintptr) { fn(UnsafeNewQTabletEvent(unsafe.Pointer(event))) }) // slot "tabletEvent"
}

func (this *QToolBox) OnActionEvent(fn func(*QActionEvent)) {
	registerVirt(this.UnsafeGoID(), 1697444140088555581, func(event uintptr) { fn(UnsafeNewQActionEvent(unsafe.Pointer(event))) }) // slot "actionEvent"
}

func (this *QToolBox) OnDragEnterEvent(fn func(*QDragEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 206122715919917589, func(event uintptr) { fn(UnsafeNewQDragEnterEvent(unsafe.Pointer(event))) }) // slot "dragEnterEvent"
}

func (this *QToolBox) OnDragMoveEvent(fn func(*QDragMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -1781545137352745978, func(event uintptr) { fn(UnsafeNewQDragMoveEvent(unsafe.Pointer(event))) }) // slot "dragMoveEvent"
}

func (this *QToolBox) OnDragLeaveEvent(fn func(*QDragLeaveEvent)) {
	registerVirt(this.UnsafeGoID(), -3139904158373725006, func(event uintptr) { fn(UnsafeNewQDragLeaveEvent(unsafe.Pointer(event))) }) // slot "dragLeaveEvent"
}

func (this *QToolBox) OnDropEvent(fn func(*QDropEvent)) {
	registerVirt(this.UnsafeGoID(), 9107058682754591716, func(event uintptr) { fn(UnsafeNewQDropEvent(unsafe.Pointer(event))) }) // slot "dropEvent"
}

func (this *QToolBox) OnHideEvent(fn func(*QHideEvent)) {
	registerVirt(this.UnsafeGoID(), -199718707876719129, func(event uintptr) { fn(UnsafeNewQHideEvent(unsafe.Pointer(event))) }) // slot "hideEvent"
}

func (this *QToolBox) OnNativeEvent(fn func(*QByteArray, unsafe.Pointer, *int64) bool) {
	registerVirt(this.UnsafeGoID(), -2039166586246443706, func(eventType uintptr, message uintptr, result uintptr) int32 { return b2i(fn(UnsafeNewQByteArray(unsafe.Pointer(eventType)), unsafe.Pointer(message), (*int64)(unsafe.Pointer(result)))) }) // slot "nativeEvent"
}

func (this *QToolBox) OnMetric(fn func(QPaintDevice__PaintDeviceMetric) int) {
	registerVirt(this.UnsafeGoID(), 427629649267311777, func(arg1 int32) int32 { return int32(fn(QPaintDevice__PaintDeviceMetric(arg1))) }) // slot "metric"
}

func (this *QToolBox) OnInitPainter(fn func(*QPainter)) {
	registerVirt(this.UnsafeGoID(), 6477500296980050210, func(painter uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter))) }) // slot "initPainter"
}

func (this *QToolBox) OnRedirected(fn func(*QPoint) *QPaintDevice) {
	registerVirt(this.UnsafeGoID(), 4127240633310655048, func(offset uintptr) uintptr { return uintptr(fn(UnsafeNewQPoint(unsafe.Pointer(offset))).UnsafePointer()) }) // slot "redirected"
}

func (this *QToolBox) OnSharedPainter(fn func() *QPainter) {
	registerVirt(this.UnsafeGoID(), 5792771328741774991, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sharedPainter"
}

func (this *QToolBox) OnInputMethodEvent(fn func(*QInputMethodEvent)) {
	registerVirt(this.UnsafeGoID(), 5438819661942843222, func(arg1 uintptr) { fn(UnsafeNewQInputMethodEvent(unsafe.Pointer(arg1))) }) // slot "inputMethodEvent"
}

func (this *QToolBox) OnInputMethodQuery(fn func(InputMethodQuery) *QVariant) {
	registerVirt(this.UnsafeGoID(), 7345853097340709638, func(arg1 int32) uintptr { return uintptr(fn(InputMethodQuery(arg1)).UnsafePointer()) }) // slot "inputMethodQuery"
}

func (this *QToolBox) OnFocusNextPrevChild(fn func(bool) bool) {
	registerVirt(this.UnsafeGoID(), 5898328517505784191, func(next int32) int32 { return b2i(fn(next != 0)) }) // slot "focusNextPrevChild"
}

func (this *QToolBox) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QToolBox) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QToolBox) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QToolBox) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QToolBox) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QToolBox) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QToolBox) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QToolBox_delete(this.h)
	}
	this.UnsafeRelease()
}
