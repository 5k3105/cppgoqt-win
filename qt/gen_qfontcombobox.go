package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QFontComboBox_chain uintptr
	addr_QFontComboBox_new uintptr
	addr_QFontComboBox_new2 uintptr
	addr_QFontComboBox_tr uintptr
	addr_QFontComboBox_setWritingSystem uintptr
	addr_QFontComboBox_writingSystem uintptr
	addr_QFontComboBox_setFontFilters uintptr
	addr_QFontComboBox_fontFilters uintptr
	addr_QFontComboBox_currentFont uintptr
	addr_QFontComboBox_setSampleTextForSystem uintptr
	addr_QFontComboBox_sampleTextForSystem uintptr
	addr_QFontComboBox_setSampleTextForFont uintptr
	addr_QFontComboBox_sampleTextForFont uintptr
	addr_QFontComboBox_setDisplayFont uintptr
	addr_QFontComboBox_displayFont uintptr
	addr_QFontComboBox_setCurrentFont uintptr
	addr_QFontComboBox_currentFontChanged uintptr
	addr_QFontComboBox_setTabOrder uintptr
	addr_QFontComboBox_mouseGrabber uintptr
	addr_QFontComboBox_keyboardGrabber uintptr
	addr_QFontComboBox_find uintptr
	addr_QFontComboBox_createWindowContainer uintptr
	addr_QFontComboBox_updateMicroFocus uintptr
	addr_QFontComboBox_create uintptr
	addr_QFontComboBox_destroy uintptr
	addr_QFontComboBox_focusNextChild uintptr
	addr_QFontComboBox_focusPreviousChild uintptr
	addr_QFontComboBox_connect uintptr
	addr_QFontComboBox_connect2 uintptr
	addr_QFontComboBox_sender uintptr
	addr_QFontComboBox_senderSignalIndex uintptr
	addr_QFontComboBox_receivers uintptr
	addr_QFontComboBox_isSignalConnected uintptr
	pg_QFontComboBox_devicePixelRatioFScale func() float64
	pg_QFontComboBox_encodeMetricF func(int32, float64) int32
	pg_QFontComboBox_getDecodedMetricF func(uintptr, int32, int32) float64
	addr_QFontComboBox_tr2 uintptr
	addr_QFontComboBox_tr3 uintptr
	addr_QFontComboBox_createWindowContainer2 uintptr
	addr_QFontComboBox_createWindowContainer3 uintptr
	addr_QFontComboBox_updateMicroFocus2 uintptr
	addr_QFontComboBox_create2 uintptr
	addr_QFontComboBox_create3 uintptr
	addr_QFontComboBox_create4 uintptr
	addr_QFontComboBox_destroy2 uintptr
	addr_QFontComboBox_destroy3 uintptr
	addr_QFontComboBox_connect3 uintptr
	addr_QFontComboBox_connect4 uintptr
	addr_QFontComboBox_metaObject_default uintptr
	addr_QFontComboBox_sizeHint_default uintptr
	addr_QFontComboBox_event_default uintptr
	addr_QFontComboBox_setModel_default uintptr
	addr_QFontComboBox_minimumSizeHint_default uintptr
	addr_QFontComboBox_showPopup_default uintptr
	addr_QFontComboBox_hidePopup_default uintptr
	addr_QFontComboBox_inputMethodQuery_default uintptr
	addr_QFontComboBox_focusInEvent_default uintptr
	addr_QFontComboBox_focusOutEvent_default uintptr
	addr_QFontComboBox_changeEvent_default uintptr
	addr_QFontComboBox_resizeEvent_default uintptr
	addr_QFontComboBox_paintEvent_default uintptr
	addr_QFontComboBox_showEvent_default uintptr
	addr_QFontComboBox_hideEvent_default uintptr
	addr_QFontComboBox_mousePressEvent_default uintptr
	addr_QFontComboBox_mouseReleaseEvent_default uintptr
	addr_QFontComboBox_keyPressEvent_default uintptr
	addr_QFontComboBox_keyReleaseEvent_default uintptr
	addr_QFontComboBox_wheelEvent_default uintptr
	addr_QFontComboBox_contextMenuEvent_default uintptr
	addr_QFontComboBox_inputMethodEvent_default uintptr
	addr_QFontComboBox_initStyleOption_default uintptr
	addr_QFontComboBox_devType_default uintptr
	addr_QFontComboBox_setVisible_default uintptr
	addr_QFontComboBox_heightForWidth_default uintptr
	addr_QFontComboBox_hasHeightForWidth_default uintptr
	addr_QFontComboBox_paintEngine_default uintptr
	addr_QFontComboBox_mouseDoubleClickEvent_default uintptr
	addr_QFontComboBox_mouseMoveEvent_default uintptr
	addr_QFontComboBox_enterEvent_default uintptr
	addr_QFontComboBox_leaveEvent_default uintptr
	addr_QFontComboBox_moveEvent_default uintptr
	addr_QFontComboBox_closeEvent_default uintptr
	addr_QFontComboBox_tabletEvent_default uintptr
	addr_QFontComboBox_actionEvent_default uintptr
	addr_QFontComboBox_dragEnterEvent_default uintptr
	addr_QFontComboBox_dragMoveEvent_default uintptr
	addr_QFontComboBox_dragLeaveEvent_default uintptr
	addr_QFontComboBox_dropEvent_default uintptr
	addr_QFontComboBox_metric_default uintptr
	addr_QFontComboBox_initPainter_default uintptr
	addr_QFontComboBox_redirected_default uintptr
	addr_QFontComboBox_sharedPainter_default uintptr
	addr_QFontComboBox_focusNextPrevChild_default uintptr
	addr_QFontComboBox_eventFilter_default uintptr
	addr_QFontComboBox_timerEvent_default uintptr
	addr_QFontComboBox_childEvent_default uintptr
	addr_QFontComboBox_customEvent_default uintptr
	addr_QFontComboBox_connectNotify_default uintptr
	addr_QFontComboBox_disconnectNotify_default uintptr
	pg_QFontComboBox_connect_currentFontChanged func(uintptr, int64)
	pg_QFontComboBox_connect_editTextChanged func(uintptr, int64)
	pg_QFontComboBox_connect_activated func(uintptr, int64)
	pg_QFontComboBox_connect_textActivated func(uintptr, int64)
	pg_QFontComboBox_connect_highlighted func(uintptr, int64)
	pg_QFontComboBox_connect_textHighlighted func(uintptr, int64)
	pg_QFontComboBox_connect_currentIndexChanged func(uintptr, int64)
	pg_QFontComboBox_connect_currentTextChanged func(uintptr, int64)
	pg_QFontComboBox_connect_windowTitleChanged func(uintptr, int64)
	pg_QFontComboBox_connect_windowIconChanged func(uintptr, int64)
	pg_QFontComboBox_connect_windowIconTextChanged func(uintptr, int64)
	pg_QFontComboBox_connect_customContextMenuRequested func(uintptr, int64)
	pg_QFontComboBox_connect_destroyed func(uintptr, int64)
	pg_QFontComboBox_connect_destroyed2 func(uintptr, int64)
	pg_QFontComboBox_connect_objectNameChanged func(uintptr, int64)
	pg_QFontComboBox__disconnect func(int64)
	pg_QFontComboBox__set_override func(uintptr, int32)
	pg_QFontComboBox_delete func(uintptr)
)

func init() {
	symaddr(&addr_QFontComboBox_chain, "QFontComboBox_chain")
	symaddr(&addr_QFontComboBox_new, "QFontComboBox_new")
	symaddr(&addr_QFontComboBox_new2, "QFontComboBox_new2")
	symaddr(&addr_QFontComboBox_tr, "QFontComboBox_tr")
	symaddr(&addr_QFontComboBox_setWritingSystem, "QFontComboBox_setWritingSystem")
	symaddr(&addr_QFontComboBox_writingSystem, "QFontComboBox_writingSystem")
	symaddr(&addr_QFontComboBox_setFontFilters, "QFontComboBox_setFontFilters")
	symaddr(&addr_QFontComboBox_fontFilters, "QFontComboBox_fontFilters")
	symaddr(&addr_QFontComboBox_currentFont, "QFontComboBox_currentFont")
	symaddr(&addr_QFontComboBox_setSampleTextForSystem, "QFontComboBox_setSampleTextForSystem")
	symaddr(&addr_QFontComboBox_sampleTextForSystem, "QFontComboBox_sampleTextForSystem")
	symaddr(&addr_QFontComboBox_setSampleTextForFont, "QFontComboBox_setSampleTextForFont")
	symaddr(&addr_QFontComboBox_sampleTextForFont, "QFontComboBox_sampleTextForFont")
	symaddr(&addr_QFontComboBox_setDisplayFont, "QFontComboBox_setDisplayFont")
	symaddr(&addr_QFontComboBox_displayFont, "QFontComboBox_displayFont")
	symaddr(&addr_QFontComboBox_setCurrentFont, "QFontComboBox_setCurrentFont")
	symaddr(&addr_QFontComboBox_currentFontChanged, "QFontComboBox_currentFontChanged")
	symaddr(&addr_QFontComboBox_setTabOrder, "QFontComboBox_setTabOrder")
	symaddr(&addr_QFontComboBox_mouseGrabber, "QFontComboBox_mouseGrabber")
	symaddr(&addr_QFontComboBox_keyboardGrabber, "QFontComboBox_keyboardGrabber")
	symaddr(&addr_QFontComboBox_find, "QFontComboBox_find")
	symaddr(&addr_QFontComboBox_createWindowContainer, "QFontComboBox_createWindowContainer")
	symaddr(&addr_QFontComboBox_updateMicroFocus, "QFontComboBox_updateMicroFocus")
	symaddr(&addr_QFontComboBox_create, "QFontComboBox_create")
	symaddr(&addr_QFontComboBox_destroy, "QFontComboBox_destroy")
	symaddr(&addr_QFontComboBox_focusNextChild, "QFontComboBox_focusNextChild")
	symaddr(&addr_QFontComboBox_focusPreviousChild, "QFontComboBox_focusPreviousChild")
	symaddr(&addr_QFontComboBox_connect, "QFontComboBox_connect")
	symaddr(&addr_QFontComboBox_connect2, "QFontComboBox_connect2")
	symaddr(&addr_QFontComboBox_sender, "QFontComboBox_sender")
	symaddr(&addr_QFontComboBox_senderSignalIndex, "QFontComboBox_senderSignalIndex")
	symaddr(&addr_QFontComboBox_receivers, "QFontComboBox_receivers")
	symaddr(&addr_QFontComboBox_isSignalConnected, "QFontComboBox_isSignalConnected")
	reg(&pg_QFontComboBox_devicePixelRatioFScale, "QFontComboBox_devicePixelRatioFScale")
	reg(&pg_QFontComboBox_encodeMetricF, "QFontComboBox_encodeMetricF")
	reg(&pg_QFontComboBox_getDecodedMetricF, "QFontComboBox_getDecodedMetricF")
	symaddr(&addr_QFontComboBox_tr2, "QFontComboBox_tr2")
	symaddr(&addr_QFontComboBox_tr3, "QFontComboBox_tr3")
	symaddr(&addr_QFontComboBox_createWindowContainer2, "QFontComboBox_createWindowContainer2")
	symaddr(&addr_QFontComboBox_createWindowContainer3, "QFontComboBox_createWindowContainer3")
	symaddr(&addr_QFontComboBox_updateMicroFocus2, "QFontComboBox_updateMicroFocus2")
	symaddr(&addr_QFontComboBox_create2, "QFontComboBox_create2")
	symaddr(&addr_QFontComboBox_create3, "QFontComboBox_create3")
	symaddr(&addr_QFontComboBox_create4, "QFontComboBox_create4")
	symaddr(&addr_QFontComboBox_destroy2, "QFontComboBox_destroy2")
	symaddr(&addr_QFontComboBox_destroy3, "QFontComboBox_destroy3")
	symaddr(&addr_QFontComboBox_connect3, "QFontComboBox_connect3")
	symaddr(&addr_QFontComboBox_connect4, "QFontComboBox_connect4")
	symaddr(&addr_QFontComboBox_metaObject_default, "QFontComboBox_metaObject_default")
	symaddr(&addr_QFontComboBox_sizeHint_default, "QFontComboBox_sizeHint_default")
	symaddr(&addr_QFontComboBox_event_default, "QFontComboBox_event_default")
	symaddr(&addr_QFontComboBox_setModel_default, "QFontComboBox_setModel_default")
	symaddr(&addr_QFontComboBox_minimumSizeHint_default, "QFontComboBox_minimumSizeHint_default")
	symaddr(&addr_QFontComboBox_showPopup_default, "QFontComboBox_showPopup_default")
	symaddr(&addr_QFontComboBox_hidePopup_default, "QFontComboBox_hidePopup_default")
	symaddr(&addr_QFontComboBox_inputMethodQuery_default, "QFontComboBox_inputMethodQuery_default")
	symaddr(&addr_QFontComboBox_focusInEvent_default, "QFontComboBox_focusInEvent_default")
	symaddr(&addr_QFontComboBox_focusOutEvent_default, "QFontComboBox_focusOutEvent_default")
	symaddr(&addr_QFontComboBox_changeEvent_default, "QFontComboBox_changeEvent_default")
	symaddr(&addr_QFontComboBox_resizeEvent_default, "QFontComboBox_resizeEvent_default")
	symaddr(&addr_QFontComboBox_paintEvent_default, "QFontComboBox_paintEvent_default")
	symaddr(&addr_QFontComboBox_showEvent_default, "QFontComboBox_showEvent_default")
	symaddr(&addr_QFontComboBox_hideEvent_default, "QFontComboBox_hideEvent_default")
	symaddr(&addr_QFontComboBox_mousePressEvent_default, "QFontComboBox_mousePressEvent_default")
	symaddr(&addr_QFontComboBox_mouseReleaseEvent_default, "QFontComboBox_mouseReleaseEvent_default")
	symaddr(&addr_QFontComboBox_keyPressEvent_default, "QFontComboBox_keyPressEvent_default")
	symaddr(&addr_QFontComboBox_keyReleaseEvent_default, "QFontComboBox_keyReleaseEvent_default")
	symaddr(&addr_QFontComboBox_wheelEvent_default, "QFontComboBox_wheelEvent_default")
	symaddr(&addr_QFontComboBox_contextMenuEvent_default, "QFontComboBox_contextMenuEvent_default")
	symaddr(&addr_QFontComboBox_inputMethodEvent_default, "QFontComboBox_inputMethodEvent_default")
	symaddr(&addr_QFontComboBox_initStyleOption_default, "QFontComboBox_initStyleOption_default")
	symaddr(&addr_QFontComboBox_devType_default, "QFontComboBox_devType_default")
	symaddr(&addr_QFontComboBox_setVisible_default, "QFontComboBox_setVisible_default")
	symaddr(&addr_QFontComboBox_heightForWidth_default, "QFontComboBox_heightForWidth_default")
	symaddr(&addr_QFontComboBox_hasHeightForWidth_default, "QFontComboBox_hasHeightForWidth_default")
	symaddr(&addr_QFontComboBox_paintEngine_default, "QFontComboBox_paintEngine_default")
	symaddr(&addr_QFontComboBox_mouseDoubleClickEvent_default, "QFontComboBox_mouseDoubleClickEvent_default")
	symaddr(&addr_QFontComboBox_mouseMoveEvent_default, "QFontComboBox_mouseMoveEvent_default")
	symaddr(&addr_QFontComboBox_enterEvent_default, "QFontComboBox_enterEvent_default")
	symaddr(&addr_QFontComboBox_leaveEvent_default, "QFontComboBox_leaveEvent_default")
	symaddr(&addr_QFontComboBox_moveEvent_default, "QFontComboBox_moveEvent_default")
	symaddr(&addr_QFontComboBox_closeEvent_default, "QFontComboBox_closeEvent_default")
	symaddr(&addr_QFontComboBox_tabletEvent_default, "QFontComboBox_tabletEvent_default")
	symaddr(&addr_QFontComboBox_actionEvent_default, "QFontComboBox_actionEvent_default")
	symaddr(&addr_QFontComboBox_dragEnterEvent_default, "QFontComboBox_dragEnterEvent_default")
	symaddr(&addr_QFontComboBox_dragMoveEvent_default, "QFontComboBox_dragMoveEvent_default")
	symaddr(&addr_QFontComboBox_dragLeaveEvent_default, "QFontComboBox_dragLeaveEvent_default")
	symaddr(&addr_QFontComboBox_dropEvent_default, "QFontComboBox_dropEvent_default")
	symaddr(&addr_QFontComboBox_metric_default, "QFontComboBox_metric_default")
	symaddr(&addr_QFontComboBox_initPainter_default, "QFontComboBox_initPainter_default")
	symaddr(&addr_QFontComboBox_redirected_default, "QFontComboBox_redirected_default")
	symaddr(&addr_QFontComboBox_sharedPainter_default, "QFontComboBox_sharedPainter_default")
	symaddr(&addr_QFontComboBox_focusNextPrevChild_default, "QFontComboBox_focusNextPrevChild_default")
	symaddr(&addr_QFontComboBox_eventFilter_default, "QFontComboBox_eventFilter_default")
	symaddr(&addr_QFontComboBox_timerEvent_default, "QFontComboBox_timerEvent_default")
	symaddr(&addr_QFontComboBox_childEvent_default, "QFontComboBox_childEvent_default")
	symaddr(&addr_QFontComboBox_customEvent_default, "QFontComboBox_customEvent_default")
	symaddr(&addr_QFontComboBox_connectNotify_default, "QFontComboBox_connectNotify_default")
	symaddr(&addr_QFontComboBox_disconnectNotify_default, "QFontComboBox_disconnectNotify_default")
	reg(&pg_QFontComboBox_connect_currentFontChanged, "QFontComboBox_connect_currentFontChanged")
	reg(&pg_QFontComboBox_connect_editTextChanged, "QFontComboBox_connect_editTextChanged")
	reg(&pg_QFontComboBox_connect_activated, "QFontComboBox_connect_activated")
	reg(&pg_QFontComboBox_connect_textActivated, "QFontComboBox_connect_textActivated")
	reg(&pg_QFontComboBox_connect_highlighted, "QFontComboBox_connect_highlighted")
	reg(&pg_QFontComboBox_connect_textHighlighted, "QFontComboBox_connect_textHighlighted")
	reg(&pg_QFontComboBox_connect_currentIndexChanged, "QFontComboBox_connect_currentIndexChanged")
	reg(&pg_QFontComboBox_connect_currentTextChanged, "QFontComboBox_connect_currentTextChanged")
	reg(&pg_QFontComboBox_connect_windowTitleChanged, "QFontComboBox_connect_windowTitleChanged")
	reg(&pg_QFontComboBox_connect_windowIconChanged, "QFontComboBox_connect_windowIconChanged")
	reg(&pg_QFontComboBox_connect_windowIconTextChanged, "QFontComboBox_connect_windowIconTextChanged")
	reg(&pg_QFontComboBox_connect_customContextMenuRequested, "QFontComboBox_connect_customContextMenuRequested")
	reg(&pg_QFontComboBox_connect_destroyed, "QFontComboBox_connect_destroyed")
	reg(&pg_QFontComboBox_connect_destroyed2, "QFontComboBox_connect_destroyed2")
	reg(&pg_QFontComboBox_connect_objectNameChanged, "QFontComboBox_connect_objectNameChanged")
	reg(&pg_QFontComboBox__disconnect, "QFontComboBox__disconnect")
	reg(&pg_QFontComboBox__set_override, "QFontComboBox__set_override")
	reg(&pg_QFontComboBox_delete, "QFontComboBox_delete")
}

type QFontComboBox struct {
	h uintptr
	*QComboBox
}

func bindQFontComboBox(h uintptr) *QFontComboBox {
	if h == 0 {
		return nil
	}
	var __c [3]uintptr
	purego.SyscallN(addr_QFontComboBox_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QFontComboBox{h: h, QComboBox: &QComboBox{h: __c[0], QWidget: &QWidget{h: __c[1], QObject: &QObject{h: __c[2]}}}}
}

func newownedQFontComboBox(h uintptr) *QFontComboBox { return bindQFontComboBox(h) }

func (this *QFontComboBox) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QFontComboBox) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQFontComboBox(h unsafe.Pointer) *QFontComboBox { return bindQFontComboBox(uintptr(h)) }
func UnsafeNewQFontComboBoxOwned(h unsafe.Pointer) *QFontComboBox { return newownedQFontComboBox(uintptr(h)) }

func (this *QFontComboBox) UnsafeRelease() {
	this.h = 0
	if this.QComboBox != nil {
		this.QComboBox.UnsafeRelease()
	}
}

func NewQFontComboBox(parent *QWidget) *QFontComboBox {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QFontComboBox_new, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQFontComboBox(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QFontComboBox_virtIdx, QFontComboBox_enable)
	return __o
}

func NewQFontComboBox2() *QFontComboBox {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QFontComboBox_new2, uintptr(__goid))
	__o := newownedQFontComboBox(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QFontComboBox_virtIdx, QFontComboBox_enable)
	return __o
}

func QFontComboBox_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFontComboBox_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QFontComboBox) SetWritingSystem(arg1 QFontDatabase__WritingSystem) {
	purego.SyscallN(addr_QFontComboBox_setWritingSystem, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QFontComboBox) WritingSystem() QFontDatabase__WritingSystem {
	__r0, _, _ := purego.SyscallN(addr_QFontComboBox_writingSystem, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QFontDatabase__WritingSystem(__r)
}

func (this *QFontComboBox) SetFontFilters(filters QFontComboBox__FontFilter) {
	purego.SyscallN(addr_QFontComboBox_setFontFilters, this.h, uintptr(filters))
	runtime.KeepAlive(this)
}

func (this *QFontComboBox) FontFilters() QFontComboBox__FontFilter {
	__r0, _, _ := purego.SyscallN(addr_QFontComboBox_fontFilters, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QFontComboBox__FontFilter(__r)
}

func (this *QFontComboBox) CurrentFont() *QFont {
	__r, _, _ := purego.SyscallN(addr_QFontComboBox_currentFont, this.h)
	runtime.KeepAlive(this)
	return newownedQFont(__r)
}

func (this *QFontComboBox) SetSampleTextForSystem(writingSystem QFontDatabase__WritingSystem, sampleText string) {
	var sampleText_p uintptr
	if len(sampleText) > 0 {
		sampleText_p = uintptr(unsafe.Pointer(unsafe.StringData(sampleText)))
	}
	purego.SyscallN(addr_QFontComboBox_setSampleTextForSystem, this.h, uintptr(writingSystem), sampleText_p, uintptr(len(sampleText)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sampleText)
}

func (this *QFontComboBox) SampleTextForSystem(writingSystem QFontDatabase__WritingSystem) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFontComboBox_sampleTextForSystem, this.h, uintptr(writingSystem), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QFontComboBox) SetSampleTextForFont(fontFamily string, sampleText string) {
	var fontFamily_p uintptr
	if len(fontFamily) > 0 {
		fontFamily_p = uintptr(unsafe.Pointer(unsafe.StringData(fontFamily)))
	}
	var sampleText_p uintptr
	if len(sampleText) > 0 {
		sampleText_p = uintptr(unsafe.Pointer(unsafe.StringData(sampleText)))
	}
	purego.SyscallN(addr_QFontComboBox_setSampleTextForFont, this.h, fontFamily_p, uintptr(len(fontFamily)), sampleText_p, uintptr(len(sampleText)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(fontFamily)
	runtime.KeepAlive(sampleText)
}

func (this *QFontComboBox) SampleTextForFont(fontFamily string) string {
	var fontFamily_p uintptr
	if len(fontFamily) > 0 {
		fontFamily_p = uintptr(unsafe.Pointer(unsafe.StringData(fontFamily)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFontComboBox_sampleTextForFont, this.h, fontFamily_p, uintptr(len(fontFamily)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(fontFamily)
	return pgStr(__r, __n)
}

func (this *QFontComboBox) SetDisplayFont(fontFamily string, font *QFont) {
	var fontFamily_p uintptr
	if len(fontFamily) > 0 {
		fontFamily_p = uintptr(unsafe.Pointer(unsafe.StringData(fontFamily)))
	}
	purego.SyscallN(addr_QFontComboBox_setDisplayFont, this.h, fontFamily_p, uintptr(len(fontFamily)), uintptr(font.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(fontFamily)
	runtime.KeepAlive(font)
}

func (this *QFontComboBox) DisplayFont(fontFamily string) (*QFont, bool) {
	var fontFamily_p uintptr
	if len(fontFamily) > 0 {
		fontFamily_p = uintptr(unsafe.Pointer(unsafe.StringData(fontFamily)))
	}
	var __b int32
	__r, _, _ := purego.SyscallN(addr_QFontComboBox_displayFont, this.h, fontFamily_p, uintptr(len(fontFamily)), uintptr(unsafe.Pointer(&__b)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(fontFamily)
	return newownedQFont(__r), __b != 0
}

func (this *QFontComboBox) SetCurrentFont(f *QFont) {
	purego.SyscallN(addr_QFontComboBox_setCurrentFont, this.h, uintptr(f.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(f)
}

func (this *QFontComboBox) CurrentFontChanged(f *QFont) {
	purego.SyscallN(addr_QFontComboBox_currentFontChanged, this.h, uintptr(f.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(f)
}

func QFontComboBox_SetTabOrder(arg1 *QWidget, arg2 *QWidget) {
	purego.SyscallN(addr_QFontComboBox_setTabOrder, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()))
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
}

func QFontComboBox_MouseGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QFontComboBox_mouseGrabber)
	return bindQWidget(__r)
}

func QFontComboBox_KeyboardGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QFontComboBox_keyboardGrabber)
	return bindQWidget(__r)
}

func QFontComboBox_Find(arg1 uint64) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QFontComboBox_find, uintptr(arg1))
	return bindQWidget(__r)
}

func QFontComboBox_CreateWindowContainer(window *QWindow) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QFontComboBox_createWindowContainer, uintptr(window.UnsafePointer()))
	runtime.KeepAlive(window)
	return bindQWidget(__r)
}

func (this *QFontComboBox) UpdateMicroFocus() {
	purego.SyscallN(addr_QFontComboBox_updateMicroFocus, this.h)
	runtime.KeepAlive(this)
}

func (this *QFontComboBox) Create() {
	purego.SyscallN(addr_QFontComboBox_create, this.h)
	runtime.KeepAlive(this)
}

func (this *QFontComboBox) Destroy() {
	purego.SyscallN(addr_QFontComboBox_destroy, this.h)
	runtime.KeepAlive(this)
}

func (this *QFontComboBox) FocusNextChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QFontComboBox_focusNextChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFontComboBox) FocusPreviousChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QFontComboBox_focusPreviousChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QFontComboBox_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QFontComboBox_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QFontComboBox_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QFontComboBox_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QFontComboBox) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QFontComboBox_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QFontComboBox) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QFontComboBox_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QFontComboBox) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QFontComboBox_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QFontComboBox) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QFontComboBox_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QFontComboBox_DevicePixelRatioFScale() float64 {
	__r := pg_QFontComboBox_devicePixelRatioFScale()
	return __r
}

func QFontComboBox_EncodeMetricF(metric QPaintDevice__PaintDeviceMetric, value float64) int {
	__r := pg_QFontComboBox_encodeMetricF(int32(metric), value)
	return int(__r)
}

func (this *QFontComboBox) GetDecodedMetricF(metricA QPaintDevice__PaintDeviceMetric, metricB QPaintDevice__PaintDeviceMetric) float64 {
	__r := pg_QFontComboBox_getDecodedMetricF(this.h, int32(metricA), int32(metricB))
	runtime.KeepAlive(this)
	return __r
}

func QFontComboBox_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFontComboBox_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QFontComboBox_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFontComboBox_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QFontComboBox_CreateWindowContainer2(window *QWindow, parent *QWidget) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QFontComboBox_createWindowContainer2, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func QFontComboBox_CreateWindowContainer3(window *QWindow, parent *QWidget, flags WindowType) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QFontComboBox_createWindowContainer3, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()), uintptr(flags))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func (this *QFontComboBox) UpdateMicroFocusWithQuery(query InputMethodQuery) {
	purego.SyscallN(addr_QFontComboBox_updateMicroFocus2, this.h, uintptr(query))
	runtime.KeepAlive(this)
}

func (this *QFontComboBox) CreateWith(arg1 uint64) {
	purego.SyscallN(addr_QFontComboBox_create2, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QFontComboBox) Create2(arg1 uint64, initializeWindow bool) {
	purego.SyscallN(addr_QFontComboBox_create3, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)))
	runtime.KeepAlive(this)
}

func (this *QFontComboBox) Create3(arg1 uint64, initializeWindow bool, destroyOldWindow bool) {
	purego.SyscallN(addr_QFontComboBox_create4, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)), uintptr(b2i(destroyOldWindow)))
	runtime.KeepAlive(this)
}

func (this *QFontComboBox) DestroyWithDestroyWindow(destroyWindow bool) {
	purego.SyscallN(addr_QFontComboBox_destroy2, this.h, uintptr(b2i(destroyWindow)))
	runtime.KeepAlive(this)
}

func (this *QFontComboBox) Destroy2(destroyWindow bool, destroySubWindows bool) {
	purego.SyscallN(addr_QFontComboBox_destroy3, this.h, uintptr(b2i(destroyWindow)), uintptr(b2i(destroySubWindows)))
	runtime.KeepAlive(this)
}

func QFontComboBox_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QFontComboBox_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QFontComboBox_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QFontComboBox_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QFontComboBox) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QFontComboBox_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QFontComboBox) SizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QFontComboBox_sizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QFontComboBox) EventDefault(e *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QFontComboBox_event_default, this.h, uintptr(e.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
	return __r != 0
}

func (this *QFontComboBox) SetModelDefault(model *QAbstractItemModel) {
	purego.SyscallN(addr_QFontComboBox_setModel_default, this.h, uintptr(model.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(model)
}

func (this *QFontComboBox) MinimumSizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QFontComboBox_minimumSizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QFontComboBox) ShowPopupDefault() {
	purego.SyscallN(addr_QFontComboBox_showPopup_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QFontComboBox) HidePopupDefault() {
	purego.SyscallN(addr_QFontComboBox_hidePopup_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QFontComboBox) InputMethodQueryDefault(arg1 InputMethodQuery) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QFontComboBox_inputMethodQuery_default, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QFontComboBox) FocusInEventDefault(e *QFocusEvent) {
	purego.SyscallN(addr_QFontComboBox_focusInEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QFontComboBox) FocusOutEventDefault(e *QFocusEvent) {
	purego.SyscallN(addr_QFontComboBox_focusOutEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QFontComboBox) ChangeEventDefault(e *QEvent) {
	purego.SyscallN(addr_QFontComboBox_changeEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QFontComboBox) ResizeEventDefault(e *QResizeEvent) {
	purego.SyscallN(addr_QFontComboBox_resizeEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QFontComboBox) PaintEventDefault(e *QPaintEvent) {
	purego.SyscallN(addr_QFontComboBox_paintEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QFontComboBox) ShowEventDefault(e *QShowEvent) {
	purego.SyscallN(addr_QFontComboBox_showEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QFontComboBox) HideEventDefault(e *QHideEvent) {
	purego.SyscallN(addr_QFontComboBox_hideEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QFontComboBox) MousePressEventDefault(e *QMouseEvent) {
	purego.SyscallN(addr_QFontComboBox_mousePressEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QFontComboBox) MouseReleaseEventDefault(e *QMouseEvent) {
	purego.SyscallN(addr_QFontComboBox_mouseReleaseEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QFontComboBox) KeyPressEventDefault(e *QKeyEvent) {
	purego.SyscallN(addr_QFontComboBox_keyPressEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QFontComboBox) KeyReleaseEventDefault(e *QKeyEvent) {
	purego.SyscallN(addr_QFontComboBox_keyReleaseEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QFontComboBox) WheelEventDefault(e *QWheelEvent) {
	purego.SyscallN(addr_QFontComboBox_wheelEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QFontComboBox) ContextMenuEventDefault(e *QContextMenuEvent) {
	purego.SyscallN(addr_QFontComboBox_contextMenuEvent_default, this.h, uintptr(e.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
}

func (this *QFontComboBox) InputMethodEventDefault(arg1 *QInputMethodEvent) {
	purego.SyscallN(addr_QFontComboBox_inputMethodEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QFontComboBox) InitStyleOptionDefault(option *QStyleOptionComboBox) {
	purego.SyscallN(addr_QFontComboBox_initStyleOption_default, this.h, uintptr(option.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(option)
}

func (this *QFontComboBox) DevTypeDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QFontComboBox_devType_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QFontComboBox) SetVisibleDefault(visible bool) {
	purego.SyscallN(addr_QFontComboBox_setVisible_default, this.h, uintptr(b2i(visible)))
	runtime.KeepAlive(this)
}

func (this *QFontComboBox) HeightForWidthDefault(arg1 int) int {
	__r0, _, _ := purego.SyscallN(addr_QFontComboBox_heightForWidth_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QFontComboBox) HasHeightForWidthDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QFontComboBox_hasHeightForWidth_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFontComboBox) PaintEngineDefault() *QPaintEngine {
	__r, _, _ := purego.SyscallN(addr_QFontComboBox_paintEngine_default, this.h)
	runtime.KeepAlive(this)
	return bindQPaintEngine(__r)
}

func (this *QFontComboBox) MouseDoubleClickEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QFontComboBox_mouseDoubleClickEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFontComboBox) MouseMoveEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QFontComboBox_mouseMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFontComboBox) EnterEventDefault(event *QEnterEvent) {
	purego.SyscallN(addr_QFontComboBox_enterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFontComboBox) LeaveEventDefault(event *QEvent) {
	purego.SyscallN(addr_QFontComboBox_leaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFontComboBox) MoveEventDefault(event *QMoveEvent) {
	purego.SyscallN(addr_QFontComboBox_moveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFontComboBox) CloseEventDefault(event *QCloseEvent) {
	purego.SyscallN(addr_QFontComboBox_closeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFontComboBox) TabletEventDefault(event *QTabletEvent) {
	purego.SyscallN(addr_QFontComboBox_tabletEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFontComboBox) ActionEventDefault(event *QActionEvent) {
	purego.SyscallN(addr_QFontComboBox_actionEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFontComboBox) DragEnterEventDefault(event *QDragEnterEvent) {
	purego.SyscallN(addr_QFontComboBox_dragEnterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFontComboBox) DragMoveEventDefault(event *QDragMoveEvent) {
	purego.SyscallN(addr_QFontComboBox_dragMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFontComboBox) DragLeaveEventDefault(event *QDragLeaveEvent) {
	purego.SyscallN(addr_QFontComboBox_dragLeaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFontComboBox) DropEventDefault(event *QDropEvent) {
	purego.SyscallN(addr_QFontComboBox_dropEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFontComboBox) MetricDefault(arg1 QPaintDevice__PaintDeviceMetric) int {
	__r0, _, _ := purego.SyscallN(addr_QFontComboBox_metric_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QFontComboBox) InitPainterDefault(painter *QPainter) {
	purego.SyscallN(addr_QFontComboBox_initPainter_default, this.h, uintptr(painter.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
}

func (this *QFontComboBox) RedirectedDefault(offset *QPoint) *QPaintDevice {
	__r, _, _ := purego.SyscallN(addr_QFontComboBox_redirected_default, this.h, uintptr(offset.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(offset)
	return bindQPaintDevice(__r)
}

func (this *QFontComboBox) SharedPainterDefault() *QPainter {
	__r, _, _ := purego.SyscallN(addr_QFontComboBox_sharedPainter_default, this.h)
	runtime.KeepAlive(this)
	return bindQPainter(__r)
}

func (this *QFontComboBox) FocusNextPrevChildDefault(next bool) bool {
	__r0, _, _ := purego.SyscallN(addr_QFontComboBox_focusNextPrevChild_default, this.h, uintptr(b2i(next)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFontComboBox) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QFontComboBox_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QFontComboBox) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QFontComboBox_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFontComboBox) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QFontComboBox_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFontComboBox) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QFontComboBox_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFontComboBox) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QFontComboBox_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QFontComboBox) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QFontComboBox_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QFontComboBox) OnCurrentFontChanged(fn func(*QFont)) int64 {
	id := registerCallback(func(f uintptr) { fn(UnsafeNewQFont(unsafe.Pointer(f))) })
	pg_QFontComboBox_connect_currentFontChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFontComboBox) OnEditTextChanged(fn func(string)) int64 {
	id := registerCallback(func(arg1 uintptr, arg1_len int64) { fn(pgStrNoFree(arg1, arg1_len)) })
	pg_QFontComboBox_connect_editTextChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFontComboBox) OnActivated(fn func(int)) int64 {
	id := registerCallback(func(index int32) { fn(int(index)) })
	pg_QFontComboBox_connect_activated(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFontComboBox) OnTextActivated(fn func(string)) int64 {
	id := registerCallback(func(arg1 uintptr, arg1_len int64) { fn(pgStrNoFree(arg1, arg1_len)) })
	pg_QFontComboBox_connect_textActivated(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFontComboBox) OnHighlighted(fn func(int)) int64 {
	id := registerCallback(func(index int32) { fn(int(index)) })
	pg_QFontComboBox_connect_highlighted(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFontComboBox) OnTextHighlighted(fn func(string)) int64 {
	id := registerCallback(func(arg1 uintptr, arg1_len int64) { fn(pgStrNoFree(arg1, arg1_len)) })
	pg_QFontComboBox_connect_textHighlighted(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFontComboBox) OnCurrentIndexChanged(fn func(int)) int64 {
	id := registerCallback(func(index int32) { fn(int(index)) })
	pg_QFontComboBox_connect_currentIndexChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFontComboBox) OnCurrentTextChanged(fn func(string)) int64 {
	id := registerCallback(func(arg1 uintptr, arg1_len int64) { fn(pgStrNoFree(arg1, arg1_len)) })
	pg_QFontComboBox_connect_currentTextChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFontComboBox) OnWindowTitleChanged(fn func(string)) int64 {
	id := registerCallback(func(title uintptr, title_len int64) { fn(pgStrNoFree(title, title_len)) })
	pg_QFontComboBox_connect_windowTitleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFontComboBox) OnWindowIconChanged(fn func(*QIcon)) int64 {
	id := registerCallback(func(icon uintptr) { fn(UnsafeNewQIcon(unsafe.Pointer(icon))) })
	pg_QFontComboBox_connect_windowIconChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFontComboBox) OnWindowIconTextChanged(fn func(string)) int64 {
	id := registerCallback(func(iconText uintptr, iconText_len int64) { fn(pgStrNoFree(iconText, iconText_len)) })
	pg_QFontComboBox_connect_windowIconTextChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFontComboBox) OnCustomContextMenuRequested(fn func(*QPoint)) int64 {
	id := registerCallback(func(pos uintptr) { fn(UnsafeNewQPoint(unsafe.Pointer(pos))) })
	pg_QFontComboBox_connect_customContextMenuRequested(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFontComboBox) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QFontComboBox_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFontComboBox) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QFontComboBox_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFontComboBox) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QFontComboBox_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFontComboBox) Disconnect(handle int64) {
	pg_QFontComboBox__disconnect(handle)
	unregisterCallback(handle)
}

func QFontComboBox_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case 729690261457684101:
		return 1
	case 1237236456312824991:
		return 2
	case 4283922507359547360:
		return 3
	case 7049264599837532491:
		return 4
	case -977583612921719482:
		return 5
	case 3308350117881191909:
		return 6
	case 7345853097340709638:
		return 7
	case 975043767033552552:
		return 8
	case -385596010576048993:
		return 9
	case 4959376063826165943:
		return 10
	case -4302917285655014661:
		return 11
	case 5290442163167458465:
		return 12
	case -6825369174751916:
		return 13
	case -199718707876719129:
		return 14
	case -332090898078543319:
		return 15
	case -4415920049297795213:
		return 16
	case 1023791977679793185:
		return 17
	case 772598475950050299:
		return 18
	case 1919988340934511316:
		return 19
	case -6698043222227006015:
		return 20
	case 5438819661942843222:
		return 21
	case -488693769771308531:
		return 22
	case -8659679893868524044:
		return 23
	case 2110182719375245575:
		return 24
	case 4311534060495737339:
		return 25
	case -8266762397995423915:
		return 26
	case 6398007404756346381:
		return 27
	case -4029073905697023887:
		return 28
	case -4255494004099806581:
		return 29
	case 6316270346668237895:
		return 30
	case -551070009623711640:
		return 31
	case -9161949150489525396:
		return 32
	case -1989937050019464783:
		return 33
	case -212068866860701141:
		return 34
	case 1697444140088555581:
		return 35
	case 206122715919917589:
		return 36
	case -1781545137352745978:
		return 37
	case -3139904158373725006:
		return 38
	case 9107058682754591716:
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
	case 5898328517505784191:
		return 45
	case 5217623514630858297:
		return 46
	case -5409927581812374498:
		return 47
	case -1185315367526857833:
		return 48
	case 3769868229123868474:
		return 49
	case 7379029982170503008:
		return 50
	case -568369772982215962:
		return 51
	}
	return -1
}

func QFontComboBox_enable(h uintptr, idx int32) { pg_QFontComboBox__set_override(h, idx) }

func (this *QFontComboBox) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QFontComboBox) OnSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 729690261457684101, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sizeHint"
}

func (this *QFontComboBox) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(e uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(e)))) }) // slot "event"
}

func (this *QFontComboBox) OnSetModel(fn func(*QAbstractItemModel)) {
	registerVirt(this.UnsafeGoID(), 4283922507359547360, func(model uintptr) { fn(UnsafeNewQAbstractItemModel(unsafe.Pointer(model))) }) // slot "setModel"
}

func (this *QFontComboBox) OnMinimumSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 7049264599837532491, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "minimumSizeHint"
}

func (this *QFontComboBox) OnShowPopup(fn func()) {
	registerVirt(this.UnsafeGoID(), -977583612921719482, func() { fn() }) // slot "showPopup"
}

func (this *QFontComboBox) OnHidePopup(fn func()) {
	registerVirt(this.UnsafeGoID(), 3308350117881191909, func() { fn() }) // slot "hidePopup"
}

func (this *QFontComboBox) OnInputMethodQuery(fn func(InputMethodQuery) *QVariant) {
	registerVirt(this.UnsafeGoID(), 7345853097340709638, func(arg1 int32) uintptr { return uintptr(fn(InputMethodQuery(arg1)).UnsafePointer()) }) // slot "inputMethodQuery"
}

func (this *QFontComboBox) OnFocusInEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), 975043767033552552, func(e uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(e))) }) // slot "focusInEvent"
}

func (this *QFontComboBox) OnFocusOutEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), -385596010576048993, func(e uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(e))) }) // slot "focusOutEvent"
}

func (this *QFontComboBox) OnChangeEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 4959376063826165943, func(e uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(e))) }) // slot "changeEvent"
}

func (this *QFontComboBox) OnResizeEvent(fn func(*QResizeEvent)) {
	registerVirt(this.UnsafeGoID(), -4302917285655014661, func(e uintptr) { fn(UnsafeNewQResizeEvent(unsafe.Pointer(e))) }) // slot "resizeEvent"
}

func (this *QFontComboBox) OnPaintEvent(fn func(*QPaintEvent)) {
	registerVirt(this.UnsafeGoID(), 5290442163167458465, func(e uintptr) { fn(UnsafeNewQPaintEvent(unsafe.Pointer(e))) }) // slot "paintEvent"
}

func (this *QFontComboBox) OnShowEvent(fn func(*QShowEvent)) {
	registerVirt(this.UnsafeGoID(), -6825369174751916, func(e uintptr) { fn(UnsafeNewQShowEvent(unsafe.Pointer(e))) }) // slot "showEvent"
}

func (this *QFontComboBox) OnHideEvent(fn func(*QHideEvent)) {
	registerVirt(this.UnsafeGoID(), -199718707876719129, func(e uintptr) { fn(UnsafeNewQHideEvent(unsafe.Pointer(e))) }) // slot "hideEvent"
}

func (this *QFontComboBox) OnMousePressEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -332090898078543319, func(e uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(e))) }) // slot "mousePressEvent"
}

func (this *QFontComboBox) OnMouseReleaseEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4415920049297795213, func(e uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(e))) }) // slot "mouseReleaseEvent"
}

func (this *QFontComboBox) OnKeyPressEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 1023791977679793185, func(e uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(e))) }) // slot "keyPressEvent"
}

func (this *QFontComboBox) OnKeyReleaseEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 772598475950050299, func(e uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(e))) }) // slot "keyReleaseEvent"
}

func (this *QFontComboBox) OnWheelEvent(fn func(*QWheelEvent)) {
	registerVirt(this.UnsafeGoID(), 1919988340934511316, func(e uintptr) { fn(UnsafeNewQWheelEvent(unsafe.Pointer(e))) }) // slot "wheelEvent"
}

func (this *QFontComboBox) OnContextMenuEvent(fn func(*QContextMenuEvent)) {
	registerVirt(this.UnsafeGoID(), -6698043222227006015, func(e uintptr) { fn(UnsafeNewQContextMenuEvent(unsafe.Pointer(e))) }) // slot "contextMenuEvent"
}

func (this *QFontComboBox) OnInputMethodEvent(fn func(*QInputMethodEvent)) {
	registerVirt(this.UnsafeGoID(), 5438819661942843222, func(arg1 uintptr) { fn(UnsafeNewQInputMethodEvent(unsafe.Pointer(arg1))) }) // slot "inputMethodEvent"
}

func (this *QFontComboBox) OnInitStyleOption(fn func(*QStyleOptionComboBox)) {
	registerVirt(this.UnsafeGoID(), -488693769771308531, func(option uintptr) { fn(UnsafeNewQStyleOptionComboBox(unsafe.Pointer(option))) }) // slot "initStyleOption"
}

func (this *QFontComboBox) OnDevType(fn func() int) {
	registerVirt(this.UnsafeGoID(), -8659679893868524044, func() int32 { return int32(fn()) }) // slot "devType"
}

func (this *QFontComboBox) OnSetVisible(fn func(bool)) {
	registerVirt(this.UnsafeGoID(), 2110182719375245575, func(visible int32) { fn(visible != 0) }) // slot "setVisible"
}

func (this *QFontComboBox) OnHeightForWidth(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), 4311534060495737339, func(arg1 int32) int32 { return int32(fn(int(arg1))) }) // slot "heightForWidth"
}

func (this *QFontComboBox) OnHasHeightForWidth(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -8266762397995423915, func() int32 { return b2i(fn()) }) // slot "hasHeightForWidth"
}

func (this *QFontComboBox) OnPaintEngine(fn func() *QPaintEngine) {
	registerVirt(this.UnsafeGoID(), 6398007404756346381, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "paintEngine"
}

func (this *QFontComboBox) OnMouseDoubleClickEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4029073905697023887, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseDoubleClickEvent"
}

func (this *QFontComboBox) OnMouseMoveEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4255494004099806581, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseMoveEvent"
}

func (this *QFontComboBox) OnEnterEvent(fn func(*QEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 6316270346668237895, func(event uintptr) { fn(UnsafeNewQEnterEvent(unsafe.Pointer(event))) }) // slot "enterEvent"
}

func (this *QFontComboBox) OnLeaveEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), -551070009623711640, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "leaveEvent"
}

func (this *QFontComboBox) OnMoveEvent(fn func(*QMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -9161949150489525396, func(event uintptr) { fn(UnsafeNewQMoveEvent(unsafe.Pointer(event))) }) // slot "moveEvent"
}

func (this *QFontComboBox) OnCloseEvent(fn func(*QCloseEvent)) {
	registerVirt(this.UnsafeGoID(), -1989937050019464783, func(event uintptr) { fn(UnsafeNewQCloseEvent(unsafe.Pointer(event))) }) // slot "closeEvent"
}

func (this *QFontComboBox) OnTabletEvent(fn func(*QTabletEvent)) {
	registerVirt(this.UnsafeGoID(), -212068866860701141, func(event uintptr) { fn(UnsafeNewQTabletEvent(unsafe.Pointer(event))) }) // slot "tabletEvent"
}

func (this *QFontComboBox) OnActionEvent(fn func(*QActionEvent)) {
	registerVirt(this.UnsafeGoID(), 1697444140088555581, func(event uintptr) { fn(UnsafeNewQActionEvent(unsafe.Pointer(event))) }) // slot "actionEvent"
}

func (this *QFontComboBox) OnDragEnterEvent(fn func(*QDragEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 206122715919917589, func(event uintptr) { fn(UnsafeNewQDragEnterEvent(unsafe.Pointer(event))) }) // slot "dragEnterEvent"
}

func (this *QFontComboBox) OnDragMoveEvent(fn func(*QDragMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -1781545137352745978, func(event uintptr) { fn(UnsafeNewQDragMoveEvent(unsafe.Pointer(event))) }) // slot "dragMoveEvent"
}

func (this *QFontComboBox) OnDragLeaveEvent(fn func(*QDragLeaveEvent)) {
	registerVirt(this.UnsafeGoID(), -3139904158373725006, func(event uintptr) { fn(UnsafeNewQDragLeaveEvent(unsafe.Pointer(event))) }) // slot "dragLeaveEvent"
}

func (this *QFontComboBox) OnDropEvent(fn func(*QDropEvent)) {
	registerVirt(this.UnsafeGoID(), 9107058682754591716, func(event uintptr) { fn(UnsafeNewQDropEvent(unsafe.Pointer(event))) }) // slot "dropEvent"
}

func (this *QFontComboBox) OnNativeEvent(fn func(*QByteArray, unsafe.Pointer, *int64) bool) {
	registerVirt(this.UnsafeGoID(), -2039166586246443706, func(eventType uintptr, message uintptr, result uintptr) int32 { return b2i(fn(UnsafeNewQByteArray(unsafe.Pointer(eventType)), unsafe.Pointer(message), (*int64)(unsafe.Pointer(result)))) }) // slot "nativeEvent"
}

func (this *QFontComboBox) OnMetric(fn func(QPaintDevice__PaintDeviceMetric) int) {
	registerVirt(this.UnsafeGoID(), 427629649267311777, func(arg1 int32) int32 { return int32(fn(QPaintDevice__PaintDeviceMetric(arg1))) }) // slot "metric"
}

func (this *QFontComboBox) OnInitPainter(fn func(*QPainter)) {
	registerVirt(this.UnsafeGoID(), 6477500296980050210, func(painter uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter))) }) // slot "initPainter"
}

func (this *QFontComboBox) OnRedirected(fn func(*QPoint) *QPaintDevice) {
	registerVirt(this.UnsafeGoID(), 4127240633310655048, func(offset uintptr) uintptr { return uintptr(fn(UnsafeNewQPoint(unsafe.Pointer(offset))).UnsafePointer()) }) // slot "redirected"
}

func (this *QFontComboBox) OnSharedPainter(fn func() *QPainter) {
	registerVirt(this.UnsafeGoID(), 5792771328741774991, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sharedPainter"
}

func (this *QFontComboBox) OnFocusNextPrevChild(fn func(bool) bool) {
	registerVirt(this.UnsafeGoID(), 5898328517505784191, func(next int32) int32 { return b2i(fn(next != 0)) }) // slot "focusNextPrevChild"
}

func (this *QFontComboBox) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QFontComboBox) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QFontComboBox) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QFontComboBox) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QFontComboBox) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QFontComboBox) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QFontComboBox) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QFontComboBox_delete(this.h)
	}
	this.UnsafeRelease()
}
