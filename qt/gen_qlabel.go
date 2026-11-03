package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QLabel_chain uintptr
	addr_QLabel_new uintptr
	addr_QLabel_new2 uintptr
	addr_QLabel_new3 uintptr
	addr_QLabel_new4 uintptr
	addr_QLabel_new5 uintptr
	addr_QLabel_new6 uintptr
	addr_QLabel_tr uintptr
	addr_QLabel_text uintptr
	addr_QLabel_pixmap uintptr
	addr_QLabel_pixmap2 uintptr
	addr_QLabel_picture uintptr
	addr_QLabel_picture2 uintptr
	addr_QLabel_movie uintptr
	addr_QLabel_textFormat uintptr
	addr_QLabel_setTextFormat uintptr
	addr_QLabel_alignment uintptr
	addr_QLabel_setAlignment uintptr
	addr_QLabel_setWordWrap uintptr
	addr_QLabel_wordWrap uintptr
	addr_QLabel_indent uintptr
	addr_QLabel_setIndent uintptr
	addr_QLabel_margin uintptr
	addr_QLabel_setMargin uintptr
	addr_QLabel_hasScaledContents uintptr
	addr_QLabel_setScaledContents uintptr
	addr_QLabel_setBuddy uintptr
	addr_QLabel_buddy uintptr
	addr_QLabel_openExternalLinks uintptr
	addr_QLabel_setOpenExternalLinks uintptr
	addr_QLabel_setTextInteractionFlags uintptr
	addr_QLabel_textInteractionFlags uintptr
	addr_QLabel_setSelection uintptr
	addr_QLabel_hasSelectedText uintptr
	addr_QLabel_selectedText uintptr
	addr_QLabel_selectionStart uintptr
	addr_QLabel_setText uintptr
	addr_QLabel_setPixmap uintptr
	addr_QLabel_setPicture uintptr
	addr_QLabel_setMovie uintptr
	addr_QLabel_setNum uintptr
	pg_QLabel_setNum2 func(uintptr, float64)
	addr_QLabel_clear uintptr
	addr_QLabel_linkActivated uintptr
	addr_QLabel_linkHovered uintptr
	addr_QLabel_drawFrame uintptr
	addr_QLabel_setTabOrder uintptr
	addr_QLabel_mouseGrabber uintptr
	addr_QLabel_keyboardGrabber uintptr
	addr_QLabel_find uintptr
	addr_QLabel_createWindowContainer uintptr
	addr_QLabel_updateMicroFocus uintptr
	addr_QLabel_create uintptr
	addr_QLabel_destroy uintptr
	addr_QLabel_focusNextChild uintptr
	addr_QLabel_focusPreviousChild uintptr
	addr_QLabel_connect uintptr
	addr_QLabel_connect2 uintptr
	addr_QLabel_sender uintptr
	addr_QLabel_senderSignalIndex uintptr
	addr_QLabel_receivers uintptr
	addr_QLabel_isSignalConnected uintptr
	pg_QLabel_devicePixelRatioFScale func() float64
	pg_QLabel_encodeMetricF func(int32, float64) int32
	pg_QLabel_getDecodedMetricF func(uintptr, int32, int32) float64
	addr_QLabel_tr2 uintptr
	addr_QLabel_tr3 uintptr
	addr_QLabel_createWindowContainer2 uintptr
	addr_QLabel_createWindowContainer3 uintptr
	addr_QLabel_updateMicroFocus2 uintptr
	addr_QLabel_create2 uintptr
	addr_QLabel_create3 uintptr
	addr_QLabel_create4 uintptr
	addr_QLabel_destroy2 uintptr
	addr_QLabel_destroy3 uintptr
	addr_QLabel_connect3 uintptr
	addr_QLabel_connect4 uintptr
	addr_QLabel_metaObject_default uintptr
	addr_QLabel_sizeHint_default uintptr
	addr_QLabel_minimumSizeHint_default uintptr
	addr_QLabel_heightForWidth_default uintptr
	addr_QLabel_event_default uintptr
	addr_QLabel_keyPressEvent_default uintptr
	addr_QLabel_paintEvent_default uintptr
	addr_QLabel_changeEvent_default uintptr
	addr_QLabel_mousePressEvent_default uintptr
	addr_QLabel_mouseMoveEvent_default uintptr
	addr_QLabel_mouseReleaseEvent_default uintptr
	addr_QLabel_contextMenuEvent_default uintptr
	addr_QLabel_focusInEvent_default uintptr
	addr_QLabel_focusOutEvent_default uintptr
	addr_QLabel_focusNextPrevChild_default uintptr
	addr_QLabel_initStyleOption_default uintptr
	addr_QLabel_devType_default uintptr
	addr_QLabel_setVisible_default uintptr
	addr_QLabel_hasHeightForWidth_default uintptr
	addr_QLabel_paintEngine_default uintptr
	addr_QLabel_mouseDoubleClickEvent_default uintptr
	addr_QLabel_wheelEvent_default uintptr
	addr_QLabel_keyReleaseEvent_default uintptr
	addr_QLabel_enterEvent_default uintptr
	addr_QLabel_leaveEvent_default uintptr
	addr_QLabel_moveEvent_default uintptr
	addr_QLabel_resizeEvent_default uintptr
	addr_QLabel_closeEvent_default uintptr
	addr_QLabel_tabletEvent_default uintptr
	addr_QLabel_actionEvent_default uintptr
	addr_QLabel_dragEnterEvent_default uintptr
	addr_QLabel_dragMoveEvent_default uintptr
	addr_QLabel_dragLeaveEvent_default uintptr
	addr_QLabel_dropEvent_default uintptr
	addr_QLabel_showEvent_default uintptr
	addr_QLabel_hideEvent_default uintptr
	addr_QLabel_metric_default uintptr
	addr_QLabel_initPainter_default uintptr
	addr_QLabel_redirected_default uintptr
	addr_QLabel_sharedPainter_default uintptr
	addr_QLabel_inputMethodEvent_default uintptr
	addr_QLabel_inputMethodQuery_default uintptr
	addr_QLabel_eventFilter_default uintptr
	addr_QLabel_timerEvent_default uintptr
	addr_QLabel_childEvent_default uintptr
	addr_QLabel_customEvent_default uintptr
	addr_QLabel_connectNotify_default uintptr
	addr_QLabel_disconnectNotify_default uintptr
	pg_QLabel_connect_linkActivated func(uintptr, int64)
	pg_QLabel_connect_linkHovered func(uintptr, int64)
	pg_QLabel_connect_windowTitleChanged func(uintptr, int64)
	pg_QLabel_connect_windowIconChanged func(uintptr, int64)
	pg_QLabel_connect_windowIconTextChanged func(uintptr, int64)
	pg_QLabel_connect_customContextMenuRequested func(uintptr, int64)
	pg_QLabel_connect_destroyed func(uintptr, int64)
	pg_QLabel_connect_destroyed2 func(uintptr, int64)
	pg_QLabel_connect_objectNameChanged func(uintptr, int64)
	pg_QLabel__disconnect func(int64)
	pg_QLabel__set_override func(uintptr, int32)
	pg_QLabel_delete func(uintptr)
)

func init() {
	symaddr(&addr_QLabel_chain, "QLabel_chain")
	symaddr(&addr_QLabel_new, "QLabel_new")
	symaddr(&addr_QLabel_new2, "QLabel_new2")
	symaddr(&addr_QLabel_new3, "QLabel_new3")
	symaddr(&addr_QLabel_new4, "QLabel_new4")
	symaddr(&addr_QLabel_new5, "QLabel_new5")
	symaddr(&addr_QLabel_new6, "QLabel_new6")
	symaddr(&addr_QLabel_tr, "QLabel_tr")
	symaddr(&addr_QLabel_text, "QLabel_text")
	symaddr(&addr_QLabel_pixmap, "QLabel_pixmap")
	symaddr(&addr_QLabel_pixmap2, "QLabel_pixmap2")
	symaddr(&addr_QLabel_picture, "QLabel_picture")
	symaddr(&addr_QLabel_picture2, "QLabel_picture2")
	symaddr(&addr_QLabel_movie, "QLabel_movie")
	symaddr(&addr_QLabel_textFormat, "QLabel_textFormat")
	symaddr(&addr_QLabel_setTextFormat, "QLabel_setTextFormat")
	symaddr(&addr_QLabel_alignment, "QLabel_alignment")
	symaddr(&addr_QLabel_setAlignment, "QLabel_setAlignment")
	symaddr(&addr_QLabel_setWordWrap, "QLabel_setWordWrap")
	symaddr(&addr_QLabel_wordWrap, "QLabel_wordWrap")
	symaddr(&addr_QLabel_indent, "QLabel_indent")
	symaddr(&addr_QLabel_setIndent, "QLabel_setIndent")
	symaddr(&addr_QLabel_margin, "QLabel_margin")
	symaddr(&addr_QLabel_setMargin, "QLabel_setMargin")
	symaddr(&addr_QLabel_hasScaledContents, "QLabel_hasScaledContents")
	symaddr(&addr_QLabel_setScaledContents, "QLabel_setScaledContents")
	symaddr(&addr_QLabel_setBuddy, "QLabel_setBuddy")
	symaddr(&addr_QLabel_buddy, "QLabel_buddy")
	symaddr(&addr_QLabel_openExternalLinks, "QLabel_openExternalLinks")
	symaddr(&addr_QLabel_setOpenExternalLinks, "QLabel_setOpenExternalLinks")
	symaddr(&addr_QLabel_setTextInteractionFlags, "QLabel_setTextInteractionFlags")
	symaddr(&addr_QLabel_textInteractionFlags, "QLabel_textInteractionFlags")
	symaddr(&addr_QLabel_setSelection, "QLabel_setSelection")
	symaddr(&addr_QLabel_hasSelectedText, "QLabel_hasSelectedText")
	symaddr(&addr_QLabel_selectedText, "QLabel_selectedText")
	symaddr(&addr_QLabel_selectionStart, "QLabel_selectionStart")
	symaddr(&addr_QLabel_setText, "QLabel_setText")
	symaddr(&addr_QLabel_setPixmap, "QLabel_setPixmap")
	symaddr(&addr_QLabel_setPicture, "QLabel_setPicture")
	symaddr(&addr_QLabel_setMovie, "QLabel_setMovie")
	symaddr(&addr_QLabel_setNum, "QLabel_setNum")
	reg(&pg_QLabel_setNum2, "QLabel_setNum2")
	symaddr(&addr_QLabel_clear, "QLabel_clear")
	symaddr(&addr_QLabel_linkActivated, "QLabel_linkActivated")
	symaddr(&addr_QLabel_linkHovered, "QLabel_linkHovered")
	symaddr(&addr_QLabel_drawFrame, "QLabel_drawFrame")
	symaddr(&addr_QLabel_setTabOrder, "QLabel_setTabOrder")
	symaddr(&addr_QLabel_mouseGrabber, "QLabel_mouseGrabber")
	symaddr(&addr_QLabel_keyboardGrabber, "QLabel_keyboardGrabber")
	symaddr(&addr_QLabel_find, "QLabel_find")
	symaddr(&addr_QLabel_createWindowContainer, "QLabel_createWindowContainer")
	symaddr(&addr_QLabel_updateMicroFocus, "QLabel_updateMicroFocus")
	symaddr(&addr_QLabel_create, "QLabel_create")
	symaddr(&addr_QLabel_destroy, "QLabel_destroy")
	symaddr(&addr_QLabel_focusNextChild, "QLabel_focusNextChild")
	symaddr(&addr_QLabel_focusPreviousChild, "QLabel_focusPreviousChild")
	symaddr(&addr_QLabel_connect, "QLabel_connect")
	symaddr(&addr_QLabel_connect2, "QLabel_connect2")
	symaddr(&addr_QLabel_sender, "QLabel_sender")
	symaddr(&addr_QLabel_senderSignalIndex, "QLabel_senderSignalIndex")
	symaddr(&addr_QLabel_receivers, "QLabel_receivers")
	symaddr(&addr_QLabel_isSignalConnected, "QLabel_isSignalConnected")
	reg(&pg_QLabel_devicePixelRatioFScale, "QLabel_devicePixelRatioFScale")
	reg(&pg_QLabel_encodeMetricF, "QLabel_encodeMetricF")
	reg(&pg_QLabel_getDecodedMetricF, "QLabel_getDecodedMetricF")
	symaddr(&addr_QLabel_tr2, "QLabel_tr2")
	symaddr(&addr_QLabel_tr3, "QLabel_tr3")
	symaddr(&addr_QLabel_createWindowContainer2, "QLabel_createWindowContainer2")
	symaddr(&addr_QLabel_createWindowContainer3, "QLabel_createWindowContainer3")
	symaddr(&addr_QLabel_updateMicroFocus2, "QLabel_updateMicroFocus2")
	symaddr(&addr_QLabel_create2, "QLabel_create2")
	symaddr(&addr_QLabel_create3, "QLabel_create3")
	symaddr(&addr_QLabel_create4, "QLabel_create4")
	symaddr(&addr_QLabel_destroy2, "QLabel_destroy2")
	symaddr(&addr_QLabel_destroy3, "QLabel_destroy3")
	symaddr(&addr_QLabel_connect3, "QLabel_connect3")
	symaddr(&addr_QLabel_connect4, "QLabel_connect4")
	symaddr(&addr_QLabel_metaObject_default, "QLabel_metaObject_default")
	symaddr(&addr_QLabel_sizeHint_default, "QLabel_sizeHint_default")
	symaddr(&addr_QLabel_minimumSizeHint_default, "QLabel_minimumSizeHint_default")
	symaddr(&addr_QLabel_heightForWidth_default, "QLabel_heightForWidth_default")
	symaddr(&addr_QLabel_event_default, "QLabel_event_default")
	symaddr(&addr_QLabel_keyPressEvent_default, "QLabel_keyPressEvent_default")
	symaddr(&addr_QLabel_paintEvent_default, "QLabel_paintEvent_default")
	symaddr(&addr_QLabel_changeEvent_default, "QLabel_changeEvent_default")
	symaddr(&addr_QLabel_mousePressEvent_default, "QLabel_mousePressEvent_default")
	symaddr(&addr_QLabel_mouseMoveEvent_default, "QLabel_mouseMoveEvent_default")
	symaddr(&addr_QLabel_mouseReleaseEvent_default, "QLabel_mouseReleaseEvent_default")
	symaddr(&addr_QLabel_contextMenuEvent_default, "QLabel_contextMenuEvent_default")
	symaddr(&addr_QLabel_focusInEvent_default, "QLabel_focusInEvent_default")
	symaddr(&addr_QLabel_focusOutEvent_default, "QLabel_focusOutEvent_default")
	symaddr(&addr_QLabel_focusNextPrevChild_default, "QLabel_focusNextPrevChild_default")
	symaddr(&addr_QLabel_initStyleOption_default, "QLabel_initStyleOption_default")
	symaddr(&addr_QLabel_devType_default, "QLabel_devType_default")
	symaddr(&addr_QLabel_setVisible_default, "QLabel_setVisible_default")
	symaddr(&addr_QLabel_hasHeightForWidth_default, "QLabel_hasHeightForWidth_default")
	symaddr(&addr_QLabel_paintEngine_default, "QLabel_paintEngine_default")
	symaddr(&addr_QLabel_mouseDoubleClickEvent_default, "QLabel_mouseDoubleClickEvent_default")
	symaddr(&addr_QLabel_wheelEvent_default, "QLabel_wheelEvent_default")
	symaddr(&addr_QLabel_keyReleaseEvent_default, "QLabel_keyReleaseEvent_default")
	symaddr(&addr_QLabel_enterEvent_default, "QLabel_enterEvent_default")
	symaddr(&addr_QLabel_leaveEvent_default, "QLabel_leaveEvent_default")
	symaddr(&addr_QLabel_moveEvent_default, "QLabel_moveEvent_default")
	symaddr(&addr_QLabel_resizeEvent_default, "QLabel_resizeEvent_default")
	symaddr(&addr_QLabel_closeEvent_default, "QLabel_closeEvent_default")
	symaddr(&addr_QLabel_tabletEvent_default, "QLabel_tabletEvent_default")
	symaddr(&addr_QLabel_actionEvent_default, "QLabel_actionEvent_default")
	symaddr(&addr_QLabel_dragEnterEvent_default, "QLabel_dragEnterEvent_default")
	symaddr(&addr_QLabel_dragMoveEvent_default, "QLabel_dragMoveEvent_default")
	symaddr(&addr_QLabel_dragLeaveEvent_default, "QLabel_dragLeaveEvent_default")
	symaddr(&addr_QLabel_dropEvent_default, "QLabel_dropEvent_default")
	symaddr(&addr_QLabel_showEvent_default, "QLabel_showEvent_default")
	symaddr(&addr_QLabel_hideEvent_default, "QLabel_hideEvent_default")
	symaddr(&addr_QLabel_metric_default, "QLabel_metric_default")
	symaddr(&addr_QLabel_initPainter_default, "QLabel_initPainter_default")
	symaddr(&addr_QLabel_redirected_default, "QLabel_redirected_default")
	symaddr(&addr_QLabel_sharedPainter_default, "QLabel_sharedPainter_default")
	symaddr(&addr_QLabel_inputMethodEvent_default, "QLabel_inputMethodEvent_default")
	symaddr(&addr_QLabel_inputMethodQuery_default, "QLabel_inputMethodQuery_default")
	symaddr(&addr_QLabel_eventFilter_default, "QLabel_eventFilter_default")
	symaddr(&addr_QLabel_timerEvent_default, "QLabel_timerEvent_default")
	symaddr(&addr_QLabel_childEvent_default, "QLabel_childEvent_default")
	symaddr(&addr_QLabel_customEvent_default, "QLabel_customEvent_default")
	symaddr(&addr_QLabel_connectNotify_default, "QLabel_connectNotify_default")
	symaddr(&addr_QLabel_disconnectNotify_default, "QLabel_disconnectNotify_default")
	reg(&pg_QLabel_connect_linkActivated, "QLabel_connect_linkActivated")
	reg(&pg_QLabel_connect_linkHovered, "QLabel_connect_linkHovered")
	reg(&pg_QLabel_connect_windowTitleChanged, "QLabel_connect_windowTitleChanged")
	reg(&pg_QLabel_connect_windowIconChanged, "QLabel_connect_windowIconChanged")
	reg(&pg_QLabel_connect_windowIconTextChanged, "QLabel_connect_windowIconTextChanged")
	reg(&pg_QLabel_connect_customContextMenuRequested, "QLabel_connect_customContextMenuRequested")
	reg(&pg_QLabel_connect_destroyed, "QLabel_connect_destroyed")
	reg(&pg_QLabel_connect_destroyed2, "QLabel_connect_destroyed2")
	reg(&pg_QLabel_connect_objectNameChanged, "QLabel_connect_objectNameChanged")
	reg(&pg_QLabel__disconnect, "QLabel__disconnect")
	reg(&pg_QLabel__set_override, "QLabel__set_override")
	reg(&pg_QLabel_delete, "QLabel_delete")
}

type QLabel struct {
	h uintptr
	*QFrame
}

func bindQLabel(h uintptr) *QLabel {
	if h == 0 {
		return nil
	}
	var __c [3]uintptr
	purego.SyscallN(addr_QLabel_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QLabel{h: h, QFrame: &QFrame{h: __c[0], QWidget: &QWidget{h: __c[1], QObject: &QObject{h: __c[2]}}}}
}

func newownedQLabel(h uintptr) *QLabel { return bindQLabel(h) }

func (this *QLabel) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QLabel) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQLabel(h unsafe.Pointer) *QLabel { return bindQLabel(uintptr(h)) }
func UnsafeNewQLabelOwned(h unsafe.Pointer) *QLabel { return newownedQLabel(uintptr(h)) }

func (this *QLabel) UnsafeRelease() {
	this.h = 0
	if this.QFrame != nil {
		this.QFrame.UnsafeRelease()
	}
}

func NewQLabel(parent *QWidget) *QLabel {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QLabel_new, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQLabel(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QLabel_virtIdx, QLabel_enable)
	return __o
}

func NewQLabel2() *QLabel {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QLabel_new2, uintptr(__goid))
	__o := newownedQLabel(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QLabel_virtIdx, QLabel_enable)
	return __o
}

func NewQLabel3(text string) *QLabel {
	__goid := nextVirtID()
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QLabel_new3, uintptr(__goid), text_p, uintptr(len(text)))
	__o := newownedQLabel(__r)
	runtime.KeepAlive(text)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QLabel_virtIdx, QLabel_enable)
	return __o
}

func NewQLabel4(parent *QWidget, f WindowType) *QLabel {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QLabel_new4, uintptr(__goid), uintptr(parent.UnsafePointer()), uintptr(f))
	__o := newownedQLabel(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QLabel_virtIdx, QLabel_enable)
	return __o
}

func NewQLabel5(text string, parent *QWidget) *QLabel {
	__goid := nextVirtID()
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QLabel_new5, uintptr(__goid), text_p, uintptr(len(text)), uintptr(parent.UnsafePointer()))
	__o := newownedQLabel(__r)
	runtime.KeepAlive(text)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QLabel_virtIdx, QLabel_enable)
	return __o
}

func NewQLabel6(text string, parent *QWidget, f WindowType) *QLabel {
	__goid := nextVirtID()
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QLabel_new6, uintptr(__goid), text_p, uintptr(len(text)), uintptr(parent.UnsafePointer()), uintptr(f))
	__o := newownedQLabel(__r)
	runtime.KeepAlive(text)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QLabel_virtIdx, QLabel_enable)
	return __o
}

func QLabel_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLabel_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QLabel) Text() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLabel_text, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QLabel) Pixmap(arg1 ReturnByValueConstant) *QPixmap {
	__r, _, _ := purego.SyscallN(addr_QLabel_pixmap, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return newownedQPixmap(__r)
}

func (this *QLabel) Pixmap2() *QPixmap {
	__r, _, _ := purego.SyscallN(addr_QLabel_pixmap2, this.h)
	runtime.KeepAlive(this)
	return newownedQPixmap(__r)
}

func (this *QLabel) Picture(arg1 ReturnByValueConstant) *QPicture {
	__r, _, _ := purego.SyscallN(addr_QLabel_picture, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return newownedQPicture(__r)
}

func (this *QLabel) Picture2() *QPicture {
	__r, _, _ := purego.SyscallN(addr_QLabel_picture2, this.h)
	runtime.KeepAlive(this)
	return newownedQPicture(__r)
}

func (this *QLabel) Movie() *QMovie {
	__r, _, _ := purego.SyscallN(addr_QLabel_movie, this.h)
	runtime.KeepAlive(this)
	return bindQMovie(__r)
}

func (this *QLabel) TextFormat() TextFormat {
	__r0, _, _ := purego.SyscallN(addr_QLabel_textFormat, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return TextFormat(__r)
}

func (this *QLabel) SetTextFormat(arg1 TextFormat) {
	purego.SyscallN(addr_QLabel_setTextFormat, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QLabel) Alignment() AlignmentFlag {
	__r0, _, _ := purego.SyscallN(addr_QLabel_alignment, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return AlignmentFlag(__r)
}

func (this *QLabel) SetAlignment(arg1 AlignmentFlag) {
	purego.SyscallN(addr_QLabel_setAlignment, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QLabel) SetWordWrap(on bool) {
	purego.SyscallN(addr_QLabel_setWordWrap, this.h, uintptr(b2i(on)))
	runtime.KeepAlive(this)
}

func (this *QLabel) WordWrap() bool {
	__r0, _, _ := purego.SyscallN(addr_QLabel_wordWrap, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QLabel) Indent() int {
	__r0, _, _ := purego.SyscallN(addr_QLabel_indent, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QLabel) SetIndent(arg1 int) {
	purego.SyscallN(addr_QLabel_setIndent, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QLabel) Margin() int {
	__r0, _, _ := purego.SyscallN(addr_QLabel_margin, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QLabel) SetMargin(arg1 int) {
	purego.SyscallN(addr_QLabel_setMargin, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QLabel) HasScaledContents() bool {
	__r0, _, _ := purego.SyscallN(addr_QLabel_hasScaledContents, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QLabel) SetScaledContents(arg1 bool) {
	purego.SyscallN(addr_QLabel_setScaledContents, this.h, uintptr(b2i(arg1)))
	runtime.KeepAlive(this)
}

func (this *QLabel) SetBuddy(arg1 *QWidget) {
	purego.SyscallN(addr_QLabel_setBuddy, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QLabel) Buddy() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QLabel_buddy, this.h)
	runtime.KeepAlive(this)
	return bindQWidget(__r)
}

func (this *QLabel) OpenExternalLinks() bool {
	__r0, _, _ := purego.SyscallN(addr_QLabel_openExternalLinks, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QLabel) SetOpenExternalLinks(open bool) {
	purego.SyscallN(addr_QLabel_setOpenExternalLinks, this.h, uintptr(b2i(open)))
	runtime.KeepAlive(this)
}

func (this *QLabel) SetTextInteractionFlags(flags TextInteractionFlag) {
	purego.SyscallN(addr_QLabel_setTextInteractionFlags, this.h, uintptr(flags))
	runtime.KeepAlive(this)
}

func (this *QLabel) TextInteractionFlags() TextInteractionFlag {
	__r0, _, _ := purego.SyscallN(addr_QLabel_textInteractionFlags, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return TextInteractionFlag(__r)
}

func (this *QLabel) SetSelection(arg1 int, arg2 int) {
	purego.SyscallN(addr_QLabel_setSelection, this.h, uintptr(arg1), uintptr(arg2))
	runtime.KeepAlive(this)
}

func (this *QLabel) HasSelectedText() bool {
	__r0, _, _ := purego.SyscallN(addr_QLabel_hasSelectedText, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QLabel) SelectedText() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLabel_selectedText, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QLabel) SelectionStart() int {
	__r0, _, _ := purego.SyscallN(addr_QLabel_selectionStart, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QLabel) SetText(arg1 string) {
	var arg1_p uintptr
	if len(arg1) > 0 {
		arg1_p = uintptr(unsafe.Pointer(unsafe.StringData(arg1)))
	}
	purego.SyscallN(addr_QLabel_setText, this.h, arg1_p, uintptr(len(arg1)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QLabel) SetPixmap(arg1 *QPixmap) {
	purego.SyscallN(addr_QLabel_setPixmap, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QLabel) SetPicture(arg1 *QPicture) {
	purego.SyscallN(addr_QLabel_setPicture, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QLabel) SetMovie(movie *QMovie) {
	purego.SyscallN(addr_QLabel_setMovie, this.h, uintptr(movie.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(movie)
}

func (this *QLabel) SetNum(arg1 int) {
	purego.SyscallN(addr_QLabel_setNum, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QLabel) SetNumWithDouble(arg1 float64) {
	pg_QLabel_setNum2(this.h, arg1)
	runtime.KeepAlive(this)
}

func (this *QLabel) Clear() {
	purego.SyscallN(addr_QLabel_clear, this.h)
	runtime.KeepAlive(this)
}

func (this *QLabel) LinkActivated(link string) {
	var link_p uintptr
	if len(link) > 0 {
		link_p = uintptr(unsafe.Pointer(unsafe.StringData(link)))
	}
	purego.SyscallN(addr_QLabel_linkActivated, this.h, link_p, uintptr(len(link)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(link)
}

func (this *QLabel) LinkHovered(link string) {
	var link_p uintptr
	if len(link) > 0 {
		link_p = uintptr(unsafe.Pointer(unsafe.StringData(link)))
	}
	purego.SyscallN(addr_QLabel_linkHovered, this.h, link_p, uintptr(len(link)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(link)
}

func (this *QLabel) DrawFrame(arg1 *QPainter) {
	purego.SyscallN(addr_QLabel_drawFrame, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func QLabel_SetTabOrder(arg1 *QWidget, arg2 *QWidget) {
	purego.SyscallN(addr_QLabel_setTabOrder, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()))
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
}

func QLabel_MouseGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QLabel_mouseGrabber)
	return bindQWidget(__r)
}

func QLabel_KeyboardGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QLabel_keyboardGrabber)
	return bindQWidget(__r)
}

func QLabel_Find(arg1 uint64) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QLabel_find, uintptr(arg1))
	return bindQWidget(__r)
}

func QLabel_CreateWindowContainer(window *QWindow) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QLabel_createWindowContainer, uintptr(window.UnsafePointer()))
	runtime.KeepAlive(window)
	return bindQWidget(__r)
}

func (this *QLabel) UpdateMicroFocus() {
	purego.SyscallN(addr_QLabel_updateMicroFocus, this.h)
	runtime.KeepAlive(this)
}

func (this *QLabel) Create() {
	purego.SyscallN(addr_QLabel_create, this.h)
	runtime.KeepAlive(this)
}

func (this *QLabel) Destroy() {
	purego.SyscallN(addr_QLabel_destroy, this.h)
	runtime.KeepAlive(this)
}

func (this *QLabel) FocusNextChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QLabel_focusNextChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QLabel) FocusPreviousChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QLabel_focusPreviousChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QLabel_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QLabel_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QLabel_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QLabel_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QLabel) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QLabel_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QLabel) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QLabel_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QLabel) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QLabel_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QLabel) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QLabel_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QLabel_DevicePixelRatioFScale() float64 {
	__r := pg_QLabel_devicePixelRatioFScale()
	return __r
}

func QLabel_EncodeMetricF(metric QPaintDevice__PaintDeviceMetric, value float64) int {
	__r := pg_QLabel_encodeMetricF(int32(metric), value)
	return int(__r)
}

func (this *QLabel) GetDecodedMetricF(metricA QPaintDevice__PaintDeviceMetric, metricB QPaintDevice__PaintDeviceMetric) float64 {
	__r := pg_QLabel_getDecodedMetricF(this.h, int32(metricA), int32(metricB))
	runtime.KeepAlive(this)
	return __r
}

func QLabel_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLabel_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QLabel_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QLabel_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QLabel_CreateWindowContainer2(window *QWindow, parent *QWidget) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QLabel_createWindowContainer2, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func QLabel_CreateWindowContainer3(window *QWindow, parent *QWidget, flags WindowType) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QLabel_createWindowContainer3, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()), uintptr(flags))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func (this *QLabel) UpdateMicroFocusWithQuery(query InputMethodQuery) {
	purego.SyscallN(addr_QLabel_updateMicroFocus2, this.h, uintptr(query))
	runtime.KeepAlive(this)
}

func (this *QLabel) CreateWith(arg1 uint64) {
	purego.SyscallN(addr_QLabel_create2, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QLabel) Create2(arg1 uint64, initializeWindow bool) {
	purego.SyscallN(addr_QLabel_create3, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)))
	runtime.KeepAlive(this)
}

func (this *QLabel) Create3(arg1 uint64, initializeWindow bool, destroyOldWindow bool) {
	purego.SyscallN(addr_QLabel_create4, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)), uintptr(b2i(destroyOldWindow)))
	runtime.KeepAlive(this)
}

func (this *QLabel) DestroyWithDestroyWindow(destroyWindow bool) {
	purego.SyscallN(addr_QLabel_destroy2, this.h, uintptr(b2i(destroyWindow)))
	runtime.KeepAlive(this)
}

func (this *QLabel) Destroy2(destroyWindow bool, destroySubWindows bool) {
	purego.SyscallN(addr_QLabel_destroy3, this.h, uintptr(b2i(destroyWindow)), uintptr(b2i(destroySubWindows)))
	runtime.KeepAlive(this)
}

func QLabel_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QLabel_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QLabel_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QLabel_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QLabel) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QLabel_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QLabel) SizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QLabel_sizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QLabel) MinimumSizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QLabel_minimumSizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QLabel) HeightForWidthDefault(arg1 int) int {
	__r0, _, _ := purego.SyscallN(addr_QLabel_heightForWidth_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QLabel) EventDefault(e *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QLabel_event_default, this.h, uintptr(e.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(e)
	return __r != 0
}

func (this *QLabel) KeyPressEventDefault(ev *QKeyEvent) {
	purego.SyscallN(addr_QLabel_keyPressEvent_default, this.h, uintptr(ev.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(ev)
}

func (this *QLabel) PaintEventDefault(arg1 *QPaintEvent) {
	purego.SyscallN(addr_QLabel_paintEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QLabel) ChangeEventDefault(arg1 *QEvent) {
	purego.SyscallN(addr_QLabel_changeEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QLabel) MousePressEventDefault(ev *QMouseEvent) {
	purego.SyscallN(addr_QLabel_mousePressEvent_default, this.h, uintptr(ev.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(ev)
}

func (this *QLabel) MouseMoveEventDefault(ev *QMouseEvent) {
	purego.SyscallN(addr_QLabel_mouseMoveEvent_default, this.h, uintptr(ev.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(ev)
}

func (this *QLabel) MouseReleaseEventDefault(ev *QMouseEvent) {
	purego.SyscallN(addr_QLabel_mouseReleaseEvent_default, this.h, uintptr(ev.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(ev)
}

func (this *QLabel) ContextMenuEventDefault(ev *QContextMenuEvent) {
	purego.SyscallN(addr_QLabel_contextMenuEvent_default, this.h, uintptr(ev.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(ev)
}

func (this *QLabel) FocusInEventDefault(ev *QFocusEvent) {
	purego.SyscallN(addr_QLabel_focusInEvent_default, this.h, uintptr(ev.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(ev)
}

func (this *QLabel) FocusOutEventDefault(ev *QFocusEvent) {
	purego.SyscallN(addr_QLabel_focusOutEvent_default, this.h, uintptr(ev.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(ev)
}

func (this *QLabel) FocusNextPrevChildDefault(next bool) bool {
	__r0, _, _ := purego.SyscallN(addr_QLabel_focusNextPrevChild_default, this.h, uintptr(b2i(next)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QLabel) InitStyleOptionDefault(option *QStyleOptionFrame) {
	purego.SyscallN(addr_QLabel_initStyleOption_default, this.h, uintptr(option.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(option)
}

func (this *QLabel) DevTypeDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QLabel_devType_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QLabel) SetVisibleDefault(visible bool) {
	purego.SyscallN(addr_QLabel_setVisible_default, this.h, uintptr(b2i(visible)))
	runtime.KeepAlive(this)
}

func (this *QLabel) HasHeightForWidthDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QLabel_hasHeightForWidth_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QLabel) PaintEngineDefault() *QPaintEngine {
	__r, _, _ := purego.SyscallN(addr_QLabel_paintEngine_default, this.h)
	runtime.KeepAlive(this)
	return bindQPaintEngine(__r)
}

func (this *QLabel) MouseDoubleClickEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QLabel_mouseDoubleClickEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QLabel) WheelEventDefault(event *QWheelEvent) {
	purego.SyscallN(addr_QLabel_wheelEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QLabel) KeyReleaseEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QLabel_keyReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QLabel) EnterEventDefault(event *QEnterEvent) {
	purego.SyscallN(addr_QLabel_enterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QLabel) LeaveEventDefault(event *QEvent) {
	purego.SyscallN(addr_QLabel_leaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QLabel) MoveEventDefault(event *QMoveEvent) {
	purego.SyscallN(addr_QLabel_moveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QLabel) ResizeEventDefault(event *QResizeEvent) {
	purego.SyscallN(addr_QLabel_resizeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QLabel) CloseEventDefault(event *QCloseEvent) {
	purego.SyscallN(addr_QLabel_closeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QLabel) TabletEventDefault(event *QTabletEvent) {
	purego.SyscallN(addr_QLabel_tabletEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QLabel) ActionEventDefault(event *QActionEvent) {
	purego.SyscallN(addr_QLabel_actionEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QLabel) DragEnterEventDefault(event *QDragEnterEvent) {
	purego.SyscallN(addr_QLabel_dragEnterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QLabel) DragMoveEventDefault(event *QDragMoveEvent) {
	purego.SyscallN(addr_QLabel_dragMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QLabel) DragLeaveEventDefault(event *QDragLeaveEvent) {
	purego.SyscallN(addr_QLabel_dragLeaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QLabel) DropEventDefault(event *QDropEvent) {
	purego.SyscallN(addr_QLabel_dropEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QLabel) ShowEventDefault(event *QShowEvent) {
	purego.SyscallN(addr_QLabel_showEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QLabel) HideEventDefault(event *QHideEvent) {
	purego.SyscallN(addr_QLabel_hideEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QLabel) MetricDefault(arg1 QPaintDevice__PaintDeviceMetric) int {
	__r0, _, _ := purego.SyscallN(addr_QLabel_metric_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QLabel) InitPainterDefault(painter *QPainter) {
	purego.SyscallN(addr_QLabel_initPainter_default, this.h, uintptr(painter.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
}

func (this *QLabel) RedirectedDefault(offset *QPoint) *QPaintDevice {
	__r, _, _ := purego.SyscallN(addr_QLabel_redirected_default, this.h, uintptr(offset.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(offset)
	return bindQPaintDevice(__r)
}

func (this *QLabel) SharedPainterDefault() *QPainter {
	__r, _, _ := purego.SyscallN(addr_QLabel_sharedPainter_default, this.h)
	runtime.KeepAlive(this)
	return bindQPainter(__r)
}

func (this *QLabel) InputMethodEventDefault(arg1 *QInputMethodEvent) {
	purego.SyscallN(addr_QLabel_inputMethodEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QLabel) InputMethodQueryDefault(arg1 InputMethodQuery) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QLabel_inputMethodQuery_default, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QLabel) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QLabel_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QLabel) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QLabel_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QLabel) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QLabel_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QLabel) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QLabel_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QLabel) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QLabel_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QLabel) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QLabel_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QLabel) OnLinkActivated(fn func(string)) int64 {
	id := registerCallback(func(link uintptr, link_len int64) { fn(pgStrNoFree(link, link_len)) })
	pg_QLabel_connect_linkActivated(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QLabel) OnLinkHovered(fn func(string)) int64 {
	id := registerCallback(func(link uintptr, link_len int64) { fn(pgStrNoFree(link, link_len)) })
	pg_QLabel_connect_linkHovered(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QLabel) OnWindowTitleChanged(fn func(string)) int64 {
	id := registerCallback(func(title uintptr, title_len int64) { fn(pgStrNoFree(title, title_len)) })
	pg_QLabel_connect_windowTitleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QLabel) OnWindowIconChanged(fn func(*QIcon)) int64 {
	id := registerCallback(func(icon uintptr) { fn(UnsafeNewQIcon(unsafe.Pointer(icon))) })
	pg_QLabel_connect_windowIconChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QLabel) OnWindowIconTextChanged(fn func(string)) int64 {
	id := registerCallback(func(iconText uintptr, iconText_len int64) { fn(pgStrNoFree(iconText, iconText_len)) })
	pg_QLabel_connect_windowIconTextChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QLabel) OnCustomContextMenuRequested(fn func(*QPoint)) int64 {
	id := registerCallback(func(pos uintptr) { fn(UnsafeNewQPoint(unsafe.Pointer(pos))) })
	pg_QLabel_connect_customContextMenuRequested(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QLabel) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QLabel_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QLabel) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QLabel_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QLabel) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QLabel_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QLabel) Disconnect(handle int64) {
	pg_QLabel__disconnect(handle)
	unregisterCallback(handle)
}

func QLabel_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case 729690261457684101:
		return 1
	case 7049264599837532491:
		return 2
	case 4311534060495737339:
		return 3
	case 1237236456312824991:
		return 4
	case 1023791977679793185:
		return 5
	case 5290442163167458465:
		return 6
	case 4959376063826165943:
		return 7
	case -332090898078543319:
		return 8
	case -4255494004099806581:
		return 9
	case -4415920049297795213:
		return 10
	case -6698043222227006015:
		return 11
	case 975043767033552552:
		return 12
	case -385596010576048993:
		return 13
	case 5898328517505784191:
		return 14
	case -488693769771308531:
		return 15
	case -8659679893868524044:
		return 16
	case 2110182719375245575:
		return 17
	case -8266762397995423915:
		return 18
	case 6398007404756346381:
		return 19
	case -4029073905697023887:
		return 20
	case 1919988340934511316:
		return 21
	case 772598475950050299:
		return 22
	case 6316270346668237895:
		return 23
	case -551070009623711640:
		return 24
	case -9161949150489525396:
		return 25
	case -4302917285655014661:
		return 26
	case -1989937050019464783:
		return 27
	case -212068866860701141:
		return 28
	case 1697444140088555581:
		return 29
	case 206122715919917589:
		return 30
	case -1781545137352745978:
		return 31
	case -3139904158373725006:
		return 32
	case 9107058682754591716:
		return 33
	case -6825369174751916:
		return 34
	case -199718707876719129:
		return 35
	case -2039166586246443706:
		return 36
	case 427629649267311777:
		return 37
	case 6477500296980050210:
		return 38
	case 4127240633310655048:
		return 39
	case 5792771328741774991:
		return 40
	case 5438819661942843222:
		return 41
	case 7345853097340709638:
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

func QLabel_enable(h uintptr, idx int32) { pg_QLabel__set_override(h, idx) }

func (this *QLabel) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QLabel) OnSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 729690261457684101, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sizeHint"
}

func (this *QLabel) OnMinimumSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 7049264599837532491, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "minimumSizeHint"
}

func (this *QLabel) OnHeightForWidth(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), 4311534060495737339, func(arg1 int32) int32 { return int32(fn(int(arg1))) }) // slot "heightForWidth"
}

func (this *QLabel) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(e uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(e)))) }) // slot "event"
}

func (this *QLabel) OnKeyPressEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 1023791977679793185, func(ev uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(ev))) }) // slot "keyPressEvent"
}

func (this *QLabel) OnPaintEvent(fn func(*QPaintEvent)) {
	registerVirt(this.UnsafeGoID(), 5290442163167458465, func(arg1 uintptr) { fn(UnsafeNewQPaintEvent(unsafe.Pointer(arg1))) }) // slot "paintEvent"
}

func (this *QLabel) OnChangeEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 4959376063826165943, func(arg1 uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(arg1))) }) // slot "changeEvent"
}

func (this *QLabel) OnMousePressEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -332090898078543319, func(ev uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(ev))) }) // slot "mousePressEvent"
}

func (this *QLabel) OnMouseMoveEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4255494004099806581, func(ev uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(ev))) }) // slot "mouseMoveEvent"
}

func (this *QLabel) OnMouseReleaseEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4415920049297795213, func(ev uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(ev))) }) // slot "mouseReleaseEvent"
}

func (this *QLabel) OnContextMenuEvent(fn func(*QContextMenuEvent)) {
	registerVirt(this.UnsafeGoID(), -6698043222227006015, func(ev uintptr) { fn(UnsafeNewQContextMenuEvent(unsafe.Pointer(ev))) }) // slot "contextMenuEvent"
}

func (this *QLabel) OnFocusInEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), 975043767033552552, func(ev uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(ev))) }) // slot "focusInEvent"
}

func (this *QLabel) OnFocusOutEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), -385596010576048993, func(ev uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(ev))) }) // slot "focusOutEvent"
}

func (this *QLabel) OnFocusNextPrevChild(fn func(bool) bool) {
	registerVirt(this.UnsafeGoID(), 5898328517505784191, func(next int32) int32 { return b2i(fn(next != 0)) }) // slot "focusNextPrevChild"
}

func (this *QLabel) OnInitStyleOption(fn func(*QStyleOptionFrame)) {
	registerVirt(this.UnsafeGoID(), -488693769771308531, func(option uintptr) { fn(UnsafeNewQStyleOptionFrame(unsafe.Pointer(option))) }) // slot "initStyleOption"
}

func (this *QLabel) OnDevType(fn func() int) {
	registerVirt(this.UnsafeGoID(), -8659679893868524044, func() int32 { return int32(fn()) }) // slot "devType"
}

func (this *QLabel) OnSetVisible(fn func(bool)) {
	registerVirt(this.UnsafeGoID(), 2110182719375245575, func(visible int32) { fn(visible != 0) }) // slot "setVisible"
}

func (this *QLabel) OnHasHeightForWidth(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -8266762397995423915, func() int32 { return b2i(fn()) }) // slot "hasHeightForWidth"
}

func (this *QLabel) OnPaintEngine(fn func() *QPaintEngine) {
	registerVirt(this.UnsafeGoID(), 6398007404756346381, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "paintEngine"
}

func (this *QLabel) OnMouseDoubleClickEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4029073905697023887, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseDoubleClickEvent"
}

func (this *QLabel) OnWheelEvent(fn func(*QWheelEvent)) {
	registerVirt(this.UnsafeGoID(), 1919988340934511316, func(event uintptr) { fn(UnsafeNewQWheelEvent(unsafe.Pointer(event))) }) // slot "wheelEvent"
}

func (this *QLabel) OnKeyReleaseEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 772598475950050299, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyReleaseEvent"
}

func (this *QLabel) OnEnterEvent(fn func(*QEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 6316270346668237895, func(event uintptr) { fn(UnsafeNewQEnterEvent(unsafe.Pointer(event))) }) // slot "enterEvent"
}

func (this *QLabel) OnLeaveEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), -551070009623711640, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "leaveEvent"
}

func (this *QLabel) OnMoveEvent(fn func(*QMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -9161949150489525396, func(event uintptr) { fn(UnsafeNewQMoveEvent(unsafe.Pointer(event))) }) // slot "moveEvent"
}

func (this *QLabel) OnResizeEvent(fn func(*QResizeEvent)) {
	registerVirt(this.UnsafeGoID(), -4302917285655014661, func(event uintptr) { fn(UnsafeNewQResizeEvent(unsafe.Pointer(event))) }) // slot "resizeEvent"
}

func (this *QLabel) OnCloseEvent(fn func(*QCloseEvent)) {
	registerVirt(this.UnsafeGoID(), -1989937050019464783, func(event uintptr) { fn(UnsafeNewQCloseEvent(unsafe.Pointer(event))) }) // slot "closeEvent"
}

func (this *QLabel) OnTabletEvent(fn func(*QTabletEvent)) {
	registerVirt(this.UnsafeGoID(), -212068866860701141, func(event uintptr) { fn(UnsafeNewQTabletEvent(unsafe.Pointer(event))) }) // slot "tabletEvent"
}

func (this *QLabel) OnActionEvent(fn func(*QActionEvent)) {
	registerVirt(this.UnsafeGoID(), 1697444140088555581, func(event uintptr) { fn(UnsafeNewQActionEvent(unsafe.Pointer(event))) }) // slot "actionEvent"
}

func (this *QLabel) OnDragEnterEvent(fn func(*QDragEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 206122715919917589, func(event uintptr) { fn(UnsafeNewQDragEnterEvent(unsafe.Pointer(event))) }) // slot "dragEnterEvent"
}

func (this *QLabel) OnDragMoveEvent(fn func(*QDragMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -1781545137352745978, func(event uintptr) { fn(UnsafeNewQDragMoveEvent(unsafe.Pointer(event))) }) // slot "dragMoveEvent"
}

func (this *QLabel) OnDragLeaveEvent(fn func(*QDragLeaveEvent)) {
	registerVirt(this.UnsafeGoID(), -3139904158373725006, func(event uintptr) { fn(UnsafeNewQDragLeaveEvent(unsafe.Pointer(event))) }) // slot "dragLeaveEvent"
}

func (this *QLabel) OnDropEvent(fn func(*QDropEvent)) {
	registerVirt(this.UnsafeGoID(), 9107058682754591716, func(event uintptr) { fn(UnsafeNewQDropEvent(unsafe.Pointer(event))) }) // slot "dropEvent"
}

func (this *QLabel) OnShowEvent(fn func(*QShowEvent)) {
	registerVirt(this.UnsafeGoID(), -6825369174751916, func(event uintptr) { fn(UnsafeNewQShowEvent(unsafe.Pointer(event))) }) // slot "showEvent"
}

func (this *QLabel) OnHideEvent(fn func(*QHideEvent)) {
	registerVirt(this.UnsafeGoID(), -199718707876719129, func(event uintptr) { fn(UnsafeNewQHideEvent(unsafe.Pointer(event))) }) // slot "hideEvent"
}

func (this *QLabel) OnNativeEvent(fn func(*QByteArray, unsafe.Pointer, *int64) bool) {
	registerVirt(this.UnsafeGoID(), -2039166586246443706, func(eventType uintptr, message uintptr, result uintptr) int32 { return b2i(fn(UnsafeNewQByteArray(unsafe.Pointer(eventType)), unsafe.Pointer(message), (*int64)(unsafe.Pointer(result)))) }) // slot "nativeEvent"
}

func (this *QLabel) OnMetric(fn func(QPaintDevice__PaintDeviceMetric) int) {
	registerVirt(this.UnsafeGoID(), 427629649267311777, func(arg1 int32) int32 { return int32(fn(QPaintDevice__PaintDeviceMetric(arg1))) }) // slot "metric"
}

func (this *QLabel) OnInitPainter(fn func(*QPainter)) {
	registerVirt(this.UnsafeGoID(), 6477500296980050210, func(painter uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter))) }) // slot "initPainter"
}

func (this *QLabel) OnRedirected(fn func(*QPoint) *QPaintDevice) {
	registerVirt(this.UnsafeGoID(), 4127240633310655048, func(offset uintptr) uintptr { return uintptr(fn(UnsafeNewQPoint(unsafe.Pointer(offset))).UnsafePointer()) }) // slot "redirected"
}

func (this *QLabel) OnSharedPainter(fn func() *QPainter) {
	registerVirt(this.UnsafeGoID(), 5792771328741774991, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sharedPainter"
}

func (this *QLabel) OnInputMethodEvent(fn func(*QInputMethodEvent)) {
	registerVirt(this.UnsafeGoID(), 5438819661942843222, func(arg1 uintptr) { fn(UnsafeNewQInputMethodEvent(unsafe.Pointer(arg1))) }) // slot "inputMethodEvent"
}

func (this *QLabel) OnInputMethodQuery(fn func(InputMethodQuery) *QVariant) {
	registerVirt(this.UnsafeGoID(), 7345853097340709638, func(arg1 int32) uintptr { return uintptr(fn(InputMethodQuery(arg1)).UnsafePointer()) }) // slot "inputMethodQuery"
}

func (this *QLabel) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QLabel) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QLabel) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QLabel) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QLabel) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QLabel) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QLabel) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QLabel_delete(this.h)
	}
	this.UnsafeRelease()
}
