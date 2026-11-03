package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QKeySequenceEdit_chain uintptr
	addr_QKeySequenceEdit_new uintptr
	addr_QKeySequenceEdit_new2 uintptr
	addr_QKeySequenceEdit_new3 uintptr
	addr_QKeySequenceEdit_new4 uintptr
	addr_QKeySequenceEdit_tr uintptr
	addr_QKeySequenceEdit_keySequence uintptr
	addr_QKeySequenceEdit_maximumSequenceLength uintptr
	addr_QKeySequenceEdit_setClearButtonEnabled uintptr
	addr_QKeySequenceEdit_isClearButtonEnabled uintptr
	addr_QKeySequenceEdit_setFinishingKeyCombinations uintptr
	addr_QKeySequenceEdit_finishingKeyCombinations uintptr
	addr_QKeySequenceEdit_setKeySequence uintptr
	addr_QKeySequenceEdit_clear uintptr
	addr_QKeySequenceEdit_setMaximumSequenceLength uintptr
	addr_QKeySequenceEdit_editingFinished uintptr
	addr_QKeySequenceEdit_keySequenceChanged uintptr
	addr_QKeySequenceEdit_setTabOrder uintptr
	addr_QKeySequenceEdit_mouseGrabber uintptr
	addr_QKeySequenceEdit_keyboardGrabber uintptr
	addr_QKeySequenceEdit_find uintptr
	addr_QKeySequenceEdit_createWindowContainer uintptr
	addr_QKeySequenceEdit_updateMicroFocus uintptr
	addr_QKeySequenceEdit_create uintptr
	addr_QKeySequenceEdit_destroy uintptr
	addr_QKeySequenceEdit_focusNextChild uintptr
	addr_QKeySequenceEdit_focusPreviousChild uintptr
	addr_QKeySequenceEdit_connect uintptr
	addr_QKeySequenceEdit_connect2 uintptr
	addr_QKeySequenceEdit_sender uintptr
	addr_QKeySequenceEdit_senderSignalIndex uintptr
	addr_QKeySequenceEdit_receivers uintptr
	addr_QKeySequenceEdit_isSignalConnected uintptr
	pg_QKeySequenceEdit_devicePixelRatioFScale func() float64
	pg_QKeySequenceEdit_encodeMetricF func(int32, float64) int32
	pg_QKeySequenceEdit_getDecodedMetricF func(uintptr, int32, int32) float64
	addr_QKeySequenceEdit_tr2 uintptr
	addr_QKeySequenceEdit_tr3 uintptr
	addr_QKeySequenceEdit_createWindowContainer2 uintptr
	addr_QKeySequenceEdit_createWindowContainer3 uintptr
	addr_QKeySequenceEdit_updateMicroFocus2 uintptr
	addr_QKeySequenceEdit_create2 uintptr
	addr_QKeySequenceEdit_create3 uintptr
	addr_QKeySequenceEdit_create4 uintptr
	addr_QKeySequenceEdit_destroy2 uintptr
	addr_QKeySequenceEdit_destroy3 uintptr
	addr_QKeySequenceEdit_connect3 uintptr
	addr_QKeySequenceEdit_connect4 uintptr
	addr_QKeySequenceEdit_metaObject_default uintptr
	addr_QKeySequenceEdit_event_default uintptr
	addr_QKeySequenceEdit_keyPressEvent_default uintptr
	addr_QKeySequenceEdit_keyReleaseEvent_default uintptr
	addr_QKeySequenceEdit_timerEvent_default uintptr
	addr_QKeySequenceEdit_focusOutEvent_default uintptr
	addr_QKeySequenceEdit_devType_default uintptr
	addr_QKeySequenceEdit_setVisible_default uintptr
	addr_QKeySequenceEdit_sizeHint_default uintptr
	addr_QKeySequenceEdit_minimumSizeHint_default uintptr
	addr_QKeySequenceEdit_heightForWidth_default uintptr
	addr_QKeySequenceEdit_hasHeightForWidth_default uintptr
	addr_QKeySequenceEdit_paintEngine_default uintptr
	addr_QKeySequenceEdit_mousePressEvent_default uintptr
	addr_QKeySequenceEdit_mouseReleaseEvent_default uintptr
	addr_QKeySequenceEdit_mouseDoubleClickEvent_default uintptr
	addr_QKeySequenceEdit_mouseMoveEvent_default uintptr
	addr_QKeySequenceEdit_wheelEvent_default uintptr
	addr_QKeySequenceEdit_focusInEvent_default uintptr
	addr_QKeySequenceEdit_enterEvent_default uintptr
	addr_QKeySequenceEdit_leaveEvent_default uintptr
	addr_QKeySequenceEdit_paintEvent_default uintptr
	addr_QKeySequenceEdit_moveEvent_default uintptr
	addr_QKeySequenceEdit_resizeEvent_default uintptr
	addr_QKeySequenceEdit_closeEvent_default uintptr
	addr_QKeySequenceEdit_contextMenuEvent_default uintptr
	addr_QKeySequenceEdit_tabletEvent_default uintptr
	addr_QKeySequenceEdit_actionEvent_default uintptr
	addr_QKeySequenceEdit_dragEnterEvent_default uintptr
	addr_QKeySequenceEdit_dragMoveEvent_default uintptr
	addr_QKeySequenceEdit_dragLeaveEvent_default uintptr
	addr_QKeySequenceEdit_dropEvent_default uintptr
	addr_QKeySequenceEdit_showEvent_default uintptr
	addr_QKeySequenceEdit_hideEvent_default uintptr
	addr_QKeySequenceEdit_changeEvent_default uintptr
	addr_QKeySequenceEdit_metric_default uintptr
	addr_QKeySequenceEdit_initPainter_default uintptr
	addr_QKeySequenceEdit_redirected_default uintptr
	addr_QKeySequenceEdit_sharedPainter_default uintptr
	addr_QKeySequenceEdit_inputMethodEvent_default uintptr
	addr_QKeySequenceEdit_inputMethodQuery_default uintptr
	addr_QKeySequenceEdit_focusNextPrevChild_default uintptr
	addr_QKeySequenceEdit_eventFilter_default uintptr
	addr_QKeySequenceEdit_childEvent_default uintptr
	addr_QKeySequenceEdit_customEvent_default uintptr
	addr_QKeySequenceEdit_connectNotify_default uintptr
	addr_QKeySequenceEdit_disconnectNotify_default uintptr
	pg_QKeySequenceEdit_connect_editingFinished func(uintptr, int64)
	pg_QKeySequenceEdit_connect_keySequenceChanged func(uintptr, int64)
	pg_QKeySequenceEdit_connect_windowTitleChanged func(uintptr, int64)
	pg_QKeySequenceEdit_connect_windowIconChanged func(uintptr, int64)
	pg_QKeySequenceEdit_connect_windowIconTextChanged func(uintptr, int64)
	pg_QKeySequenceEdit_connect_customContextMenuRequested func(uintptr, int64)
	pg_QKeySequenceEdit_connect_destroyed func(uintptr, int64)
	pg_QKeySequenceEdit_connect_destroyed2 func(uintptr, int64)
	pg_QKeySequenceEdit_connect_objectNameChanged func(uintptr, int64)
	pg_QKeySequenceEdit__disconnect func(int64)
	pg_QKeySequenceEdit__set_override func(uintptr, int32)
	pg_QKeySequenceEdit_delete func(uintptr)
)

func init() {
	symaddr(&addr_QKeySequenceEdit_chain, "QKeySequenceEdit_chain")
	symaddr(&addr_QKeySequenceEdit_new, "QKeySequenceEdit_new")
	symaddr(&addr_QKeySequenceEdit_new2, "QKeySequenceEdit_new2")
	symaddr(&addr_QKeySequenceEdit_new3, "QKeySequenceEdit_new3")
	symaddr(&addr_QKeySequenceEdit_new4, "QKeySequenceEdit_new4")
	symaddr(&addr_QKeySequenceEdit_tr, "QKeySequenceEdit_tr")
	symaddr(&addr_QKeySequenceEdit_keySequence, "QKeySequenceEdit_keySequence")
	symaddr(&addr_QKeySequenceEdit_maximumSequenceLength, "QKeySequenceEdit_maximumSequenceLength")
	symaddr(&addr_QKeySequenceEdit_setClearButtonEnabled, "QKeySequenceEdit_setClearButtonEnabled")
	symaddr(&addr_QKeySequenceEdit_isClearButtonEnabled, "QKeySequenceEdit_isClearButtonEnabled")
	symaddr(&addr_QKeySequenceEdit_setFinishingKeyCombinations, "QKeySequenceEdit_setFinishingKeyCombinations")
	symaddr(&addr_QKeySequenceEdit_finishingKeyCombinations, "QKeySequenceEdit_finishingKeyCombinations")
	symaddr(&addr_QKeySequenceEdit_setKeySequence, "QKeySequenceEdit_setKeySequence")
	symaddr(&addr_QKeySequenceEdit_clear, "QKeySequenceEdit_clear")
	symaddr(&addr_QKeySequenceEdit_setMaximumSequenceLength, "QKeySequenceEdit_setMaximumSequenceLength")
	symaddr(&addr_QKeySequenceEdit_editingFinished, "QKeySequenceEdit_editingFinished")
	symaddr(&addr_QKeySequenceEdit_keySequenceChanged, "QKeySequenceEdit_keySequenceChanged")
	symaddr(&addr_QKeySequenceEdit_setTabOrder, "QKeySequenceEdit_setTabOrder")
	symaddr(&addr_QKeySequenceEdit_mouseGrabber, "QKeySequenceEdit_mouseGrabber")
	symaddr(&addr_QKeySequenceEdit_keyboardGrabber, "QKeySequenceEdit_keyboardGrabber")
	symaddr(&addr_QKeySequenceEdit_find, "QKeySequenceEdit_find")
	symaddr(&addr_QKeySequenceEdit_createWindowContainer, "QKeySequenceEdit_createWindowContainer")
	symaddr(&addr_QKeySequenceEdit_updateMicroFocus, "QKeySequenceEdit_updateMicroFocus")
	symaddr(&addr_QKeySequenceEdit_create, "QKeySequenceEdit_create")
	symaddr(&addr_QKeySequenceEdit_destroy, "QKeySequenceEdit_destroy")
	symaddr(&addr_QKeySequenceEdit_focusNextChild, "QKeySequenceEdit_focusNextChild")
	symaddr(&addr_QKeySequenceEdit_focusPreviousChild, "QKeySequenceEdit_focusPreviousChild")
	symaddr(&addr_QKeySequenceEdit_connect, "QKeySequenceEdit_connect")
	symaddr(&addr_QKeySequenceEdit_connect2, "QKeySequenceEdit_connect2")
	symaddr(&addr_QKeySequenceEdit_sender, "QKeySequenceEdit_sender")
	symaddr(&addr_QKeySequenceEdit_senderSignalIndex, "QKeySequenceEdit_senderSignalIndex")
	symaddr(&addr_QKeySequenceEdit_receivers, "QKeySequenceEdit_receivers")
	symaddr(&addr_QKeySequenceEdit_isSignalConnected, "QKeySequenceEdit_isSignalConnected")
	reg(&pg_QKeySequenceEdit_devicePixelRatioFScale, "QKeySequenceEdit_devicePixelRatioFScale")
	reg(&pg_QKeySequenceEdit_encodeMetricF, "QKeySequenceEdit_encodeMetricF")
	reg(&pg_QKeySequenceEdit_getDecodedMetricF, "QKeySequenceEdit_getDecodedMetricF")
	symaddr(&addr_QKeySequenceEdit_tr2, "QKeySequenceEdit_tr2")
	symaddr(&addr_QKeySequenceEdit_tr3, "QKeySequenceEdit_tr3")
	symaddr(&addr_QKeySequenceEdit_createWindowContainer2, "QKeySequenceEdit_createWindowContainer2")
	symaddr(&addr_QKeySequenceEdit_createWindowContainer3, "QKeySequenceEdit_createWindowContainer3")
	symaddr(&addr_QKeySequenceEdit_updateMicroFocus2, "QKeySequenceEdit_updateMicroFocus2")
	symaddr(&addr_QKeySequenceEdit_create2, "QKeySequenceEdit_create2")
	symaddr(&addr_QKeySequenceEdit_create3, "QKeySequenceEdit_create3")
	symaddr(&addr_QKeySequenceEdit_create4, "QKeySequenceEdit_create4")
	symaddr(&addr_QKeySequenceEdit_destroy2, "QKeySequenceEdit_destroy2")
	symaddr(&addr_QKeySequenceEdit_destroy3, "QKeySequenceEdit_destroy3")
	symaddr(&addr_QKeySequenceEdit_connect3, "QKeySequenceEdit_connect3")
	symaddr(&addr_QKeySequenceEdit_connect4, "QKeySequenceEdit_connect4")
	symaddr(&addr_QKeySequenceEdit_metaObject_default, "QKeySequenceEdit_metaObject_default")
	symaddr(&addr_QKeySequenceEdit_event_default, "QKeySequenceEdit_event_default")
	symaddr(&addr_QKeySequenceEdit_keyPressEvent_default, "QKeySequenceEdit_keyPressEvent_default")
	symaddr(&addr_QKeySequenceEdit_keyReleaseEvent_default, "QKeySequenceEdit_keyReleaseEvent_default")
	symaddr(&addr_QKeySequenceEdit_timerEvent_default, "QKeySequenceEdit_timerEvent_default")
	symaddr(&addr_QKeySequenceEdit_focusOutEvent_default, "QKeySequenceEdit_focusOutEvent_default")
	symaddr(&addr_QKeySequenceEdit_devType_default, "QKeySequenceEdit_devType_default")
	symaddr(&addr_QKeySequenceEdit_setVisible_default, "QKeySequenceEdit_setVisible_default")
	symaddr(&addr_QKeySequenceEdit_sizeHint_default, "QKeySequenceEdit_sizeHint_default")
	symaddr(&addr_QKeySequenceEdit_minimumSizeHint_default, "QKeySequenceEdit_minimumSizeHint_default")
	symaddr(&addr_QKeySequenceEdit_heightForWidth_default, "QKeySequenceEdit_heightForWidth_default")
	symaddr(&addr_QKeySequenceEdit_hasHeightForWidth_default, "QKeySequenceEdit_hasHeightForWidth_default")
	symaddr(&addr_QKeySequenceEdit_paintEngine_default, "QKeySequenceEdit_paintEngine_default")
	symaddr(&addr_QKeySequenceEdit_mousePressEvent_default, "QKeySequenceEdit_mousePressEvent_default")
	symaddr(&addr_QKeySequenceEdit_mouseReleaseEvent_default, "QKeySequenceEdit_mouseReleaseEvent_default")
	symaddr(&addr_QKeySequenceEdit_mouseDoubleClickEvent_default, "QKeySequenceEdit_mouseDoubleClickEvent_default")
	symaddr(&addr_QKeySequenceEdit_mouseMoveEvent_default, "QKeySequenceEdit_mouseMoveEvent_default")
	symaddr(&addr_QKeySequenceEdit_wheelEvent_default, "QKeySequenceEdit_wheelEvent_default")
	symaddr(&addr_QKeySequenceEdit_focusInEvent_default, "QKeySequenceEdit_focusInEvent_default")
	symaddr(&addr_QKeySequenceEdit_enterEvent_default, "QKeySequenceEdit_enterEvent_default")
	symaddr(&addr_QKeySequenceEdit_leaveEvent_default, "QKeySequenceEdit_leaveEvent_default")
	symaddr(&addr_QKeySequenceEdit_paintEvent_default, "QKeySequenceEdit_paintEvent_default")
	symaddr(&addr_QKeySequenceEdit_moveEvent_default, "QKeySequenceEdit_moveEvent_default")
	symaddr(&addr_QKeySequenceEdit_resizeEvent_default, "QKeySequenceEdit_resizeEvent_default")
	symaddr(&addr_QKeySequenceEdit_closeEvent_default, "QKeySequenceEdit_closeEvent_default")
	symaddr(&addr_QKeySequenceEdit_contextMenuEvent_default, "QKeySequenceEdit_contextMenuEvent_default")
	symaddr(&addr_QKeySequenceEdit_tabletEvent_default, "QKeySequenceEdit_tabletEvent_default")
	symaddr(&addr_QKeySequenceEdit_actionEvent_default, "QKeySequenceEdit_actionEvent_default")
	symaddr(&addr_QKeySequenceEdit_dragEnterEvent_default, "QKeySequenceEdit_dragEnterEvent_default")
	symaddr(&addr_QKeySequenceEdit_dragMoveEvent_default, "QKeySequenceEdit_dragMoveEvent_default")
	symaddr(&addr_QKeySequenceEdit_dragLeaveEvent_default, "QKeySequenceEdit_dragLeaveEvent_default")
	symaddr(&addr_QKeySequenceEdit_dropEvent_default, "QKeySequenceEdit_dropEvent_default")
	symaddr(&addr_QKeySequenceEdit_showEvent_default, "QKeySequenceEdit_showEvent_default")
	symaddr(&addr_QKeySequenceEdit_hideEvent_default, "QKeySequenceEdit_hideEvent_default")
	symaddr(&addr_QKeySequenceEdit_changeEvent_default, "QKeySequenceEdit_changeEvent_default")
	symaddr(&addr_QKeySequenceEdit_metric_default, "QKeySequenceEdit_metric_default")
	symaddr(&addr_QKeySequenceEdit_initPainter_default, "QKeySequenceEdit_initPainter_default")
	symaddr(&addr_QKeySequenceEdit_redirected_default, "QKeySequenceEdit_redirected_default")
	symaddr(&addr_QKeySequenceEdit_sharedPainter_default, "QKeySequenceEdit_sharedPainter_default")
	symaddr(&addr_QKeySequenceEdit_inputMethodEvent_default, "QKeySequenceEdit_inputMethodEvent_default")
	symaddr(&addr_QKeySequenceEdit_inputMethodQuery_default, "QKeySequenceEdit_inputMethodQuery_default")
	symaddr(&addr_QKeySequenceEdit_focusNextPrevChild_default, "QKeySequenceEdit_focusNextPrevChild_default")
	symaddr(&addr_QKeySequenceEdit_eventFilter_default, "QKeySequenceEdit_eventFilter_default")
	symaddr(&addr_QKeySequenceEdit_childEvent_default, "QKeySequenceEdit_childEvent_default")
	symaddr(&addr_QKeySequenceEdit_customEvent_default, "QKeySequenceEdit_customEvent_default")
	symaddr(&addr_QKeySequenceEdit_connectNotify_default, "QKeySequenceEdit_connectNotify_default")
	symaddr(&addr_QKeySequenceEdit_disconnectNotify_default, "QKeySequenceEdit_disconnectNotify_default")
	reg(&pg_QKeySequenceEdit_connect_editingFinished, "QKeySequenceEdit_connect_editingFinished")
	reg(&pg_QKeySequenceEdit_connect_keySequenceChanged, "QKeySequenceEdit_connect_keySequenceChanged")
	reg(&pg_QKeySequenceEdit_connect_windowTitleChanged, "QKeySequenceEdit_connect_windowTitleChanged")
	reg(&pg_QKeySequenceEdit_connect_windowIconChanged, "QKeySequenceEdit_connect_windowIconChanged")
	reg(&pg_QKeySequenceEdit_connect_windowIconTextChanged, "QKeySequenceEdit_connect_windowIconTextChanged")
	reg(&pg_QKeySequenceEdit_connect_customContextMenuRequested, "QKeySequenceEdit_connect_customContextMenuRequested")
	reg(&pg_QKeySequenceEdit_connect_destroyed, "QKeySequenceEdit_connect_destroyed")
	reg(&pg_QKeySequenceEdit_connect_destroyed2, "QKeySequenceEdit_connect_destroyed2")
	reg(&pg_QKeySequenceEdit_connect_objectNameChanged, "QKeySequenceEdit_connect_objectNameChanged")
	reg(&pg_QKeySequenceEdit__disconnect, "QKeySequenceEdit__disconnect")
	reg(&pg_QKeySequenceEdit__set_override, "QKeySequenceEdit__set_override")
	reg(&pg_QKeySequenceEdit_delete, "QKeySequenceEdit_delete")
}

type QKeySequenceEdit struct {
	h uintptr
	*QWidget
}

func bindQKeySequenceEdit(h uintptr) *QKeySequenceEdit {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QKeySequenceEdit_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QKeySequenceEdit{h: h, QWidget: &QWidget{h: __c[0], QObject: &QObject{h: __c[1]}}}
}

func newownedQKeySequenceEdit(h uintptr) *QKeySequenceEdit { return bindQKeySequenceEdit(h) }

func (this *QKeySequenceEdit) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QKeySequenceEdit) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQKeySequenceEdit(h unsafe.Pointer) *QKeySequenceEdit { return bindQKeySequenceEdit(uintptr(h)) }
func UnsafeNewQKeySequenceEditOwned(h unsafe.Pointer) *QKeySequenceEdit { return newownedQKeySequenceEdit(uintptr(h)) }

func (this *QKeySequenceEdit) UnsafeRelease() {
	this.h = 0
	if this.QWidget != nil {
		this.QWidget.UnsafeRelease()
	}
}

func NewQKeySequenceEdit(parent *QWidget) *QKeySequenceEdit {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QKeySequenceEdit_new, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQKeySequenceEdit(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QKeySequenceEdit_virtIdx, QKeySequenceEdit_enable)
	return __o
}

func NewQKeySequenceEdit2() *QKeySequenceEdit {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QKeySequenceEdit_new2, uintptr(__goid))
	__o := newownedQKeySequenceEdit(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QKeySequenceEdit_virtIdx, QKeySequenceEdit_enable)
	return __o
}

func NewQKeySequenceEdit3(keySequence *QKeySequence) *QKeySequenceEdit {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QKeySequenceEdit_new3, uintptr(__goid), uintptr(keySequence.UnsafePointer()))
	__o := newownedQKeySequenceEdit(__r)
	runtime.KeepAlive(keySequence)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QKeySequenceEdit_virtIdx, QKeySequenceEdit_enable)
	return __o
}

func NewQKeySequenceEdit4(keySequence *QKeySequence, parent *QWidget) *QKeySequenceEdit {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QKeySequenceEdit_new4, uintptr(__goid), uintptr(keySequence.UnsafePointer()), uintptr(parent.UnsafePointer()))
	__o := newownedQKeySequenceEdit(__r)
	runtime.KeepAlive(keySequence)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QKeySequenceEdit_virtIdx, QKeySequenceEdit_enable)
	return __o
}

func QKeySequenceEdit_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QKeySequenceEdit_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QKeySequenceEdit) KeySequence() *QKeySequence {
	__r, _, _ := purego.SyscallN(addr_QKeySequenceEdit_keySequence, this.h)
	runtime.KeepAlive(this)
	return newownedQKeySequence(__r)
}

func (this *QKeySequenceEdit) MaximumSequenceLength() int64 {
	__r0, _, _ := purego.SyscallN(addr_QKeySequenceEdit_maximumSequenceLength, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QKeySequenceEdit) SetClearButtonEnabled(enable bool) {
	purego.SyscallN(addr_QKeySequenceEdit_setClearButtonEnabled, this.h, uintptr(b2i(enable)))
	runtime.KeepAlive(this)
}

func (this *QKeySequenceEdit) IsClearButtonEnabled() bool {
	__r0, _, _ := purego.SyscallN(addr_QKeySequenceEdit_isClearButtonEnabled, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QKeySequenceEdit) SetFinishingKeyCombinations(finishingKeyCombinations []*QKeyCombination) {
	finishingKeyCombinations_arr := make([]uintptr, len(finishingKeyCombinations))
	for __i, __e := range finishingKeyCombinations {
		finishingKeyCombinations_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var finishingKeyCombinations_p uintptr
	if len(finishingKeyCombinations_arr) > 0 {
		finishingKeyCombinations_p = uintptr(unsafe.Pointer(&finishingKeyCombinations_arr[0]))
	}
	purego.SyscallN(addr_QKeySequenceEdit_setFinishingKeyCombinations, this.h, finishingKeyCombinations_p, uintptr(len(finishingKeyCombinations)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(finishingKeyCombinations)
	runtime.KeepAlive(finishingKeyCombinations_arr)
}

func (this *QKeySequenceEdit) FinishingKeyCombinations() []*QKeyCombination {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QKeySequenceEdit_finishingKeyCombinations, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QKeyCombination, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQKeyCombination(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QKeySequenceEdit) SetKeySequence(keySequence *QKeySequence) {
	purego.SyscallN(addr_QKeySequenceEdit_setKeySequence, this.h, uintptr(keySequence.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(keySequence)
}

func (this *QKeySequenceEdit) Clear() {
	purego.SyscallN(addr_QKeySequenceEdit_clear, this.h)
	runtime.KeepAlive(this)
}

func (this *QKeySequenceEdit) SetMaximumSequenceLength(count int64) {
	purego.SyscallN(addr_QKeySequenceEdit_setMaximumSequenceLength, this.h, uintptr(count))
	runtime.KeepAlive(this)
}

func (this *QKeySequenceEdit) EditingFinished() {
	purego.SyscallN(addr_QKeySequenceEdit_editingFinished, this.h)
	runtime.KeepAlive(this)
}

func (this *QKeySequenceEdit) KeySequenceChanged(keySequence *QKeySequence) {
	purego.SyscallN(addr_QKeySequenceEdit_keySequenceChanged, this.h, uintptr(keySequence.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(keySequence)
}

func QKeySequenceEdit_SetTabOrder(arg1 *QWidget, arg2 *QWidget) {
	purego.SyscallN(addr_QKeySequenceEdit_setTabOrder, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()))
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
}

func QKeySequenceEdit_MouseGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QKeySequenceEdit_mouseGrabber)
	return bindQWidget(__r)
}

func QKeySequenceEdit_KeyboardGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QKeySequenceEdit_keyboardGrabber)
	return bindQWidget(__r)
}

func QKeySequenceEdit_Find(arg1 uint64) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QKeySequenceEdit_find, uintptr(arg1))
	return bindQWidget(__r)
}

func QKeySequenceEdit_CreateWindowContainer(window *QWindow) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QKeySequenceEdit_createWindowContainer, uintptr(window.UnsafePointer()))
	runtime.KeepAlive(window)
	return bindQWidget(__r)
}

func (this *QKeySequenceEdit) UpdateMicroFocus() {
	purego.SyscallN(addr_QKeySequenceEdit_updateMicroFocus, this.h)
	runtime.KeepAlive(this)
}

func (this *QKeySequenceEdit) Create() {
	purego.SyscallN(addr_QKeySequenceEdit_create, this.h)
	runtime.KeepAlive(this)
}

func (this *QKeySequenceEdit) Destroy() {
	purego.SyscallN(addr_QKeySequenceEdit_destroy, this.h)
	runtime.KeepAlive(this)
}

func (this *QKeySequenceEdit) FocusNextChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QKeySequenceEdit_focusNextChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QKeySequenceEdit) FocusPreviousChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QKeySequenceEdit_focusPreviousChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QKeySequenceEdit_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QKeySequenceEdit_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QKeySequenceEdit_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QKeySequenceEdit_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QKeySequenceEdit) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QKeySequenceEdit_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QKeySequenceEdit) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QKeySequenceEdit_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QKeySequenceEdit) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QKeySequenceEdit_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QKeySequenceEdit) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QKeySequenceEdit_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QKeySequenceEdit_DevicePixelRatioFScale() float64 {
	__r := pg_QKeySequenceEdit_devicePixelRatioFScale()
	return __r
}

func QKeySequenceEdit_EncodeMetricF(metric QPaintDevice__PaintDeviceMetric, value float64) int {
	__r := pg_QKeySequenceEdit_encodeMetricF(int32(metric), value)
	return int(__r)
}

func (this *QKeySequenceEdit) GetDecodedMetricF(metricA QPaintDevice__PaintDeviceMetric, metricB QPaintDevice__PaintDeviceMetric) float64 {
	__r := pg_QKeySequenceEdit_getDecodedMetricF(this.h, int32(metricA), int32(metricB))
	runtime.KeepAlive(this)
	return __r
}

func QKeySequenceEdit_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QKeySequenceEdit_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QKeySequenceEdit_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QKeySequenceEdit_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QKeySequenceEdit_CreateWindowContainer2(window *QWindow, parent *QWidget) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QKeySequenceEdit_createWindowContainer2, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func QKeySequenceEdit_CreateWindowContainer3(window *QWindow, parent *QWidget, flags WindowType) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QKeySequenceEdit_createWindowContainer3, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()), uintptr(flags))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func (this *QKeySequenceEdit) UpdateMicroFocusWithQuery(query InputMethodQuery) {
	purego.SyscallN(addr_QKeySequenceEdit_updateMicroFocus2, this.h, uintptr(query))
	runtime.KeepAlive(this)
}

func (this *QKeySequenceEdit) CreateWith(arg1 uint64) {
	purego.SyscallN(addr_QKeySequenceEdit_create2, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QKeySequenceEdit) Create2(arg1 uint64, initializeWindow bool) {
	purego.SyscallN(addr_QKeySequenceEdit_create3, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)))
	runtime.KeepAlive(this)
}

func (this *QKeySequenceEdit) Create3(arg1 uint64, initializeWindow bool, destroyOldWindow bool) {
	purego.SyscallN(addr_QKeySequenceEdit_create4, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)), uintptr(b2i(destroyOldWindow)))
	runtime.KeepAlive(this)
}

func (this *QKeySequenceEdit) DestroyWithDestroyWindow(destroyWindow bool) {
	purego.SyscallN(addr_QKeySequenceEdit_destroy2, this.h, uintptr(b2i(destroyWindow)))
	runtime.KeepAlive(this)
}

func (this *QKeySequenceEdit) Destroy2(destroyWindow bool, destroySubWindows bool) {
	purego.SyscallN(addr_QKeySequenceEdit_destroy3, this.h, uintptr(b2i(destroyWindow)), uintptr(b2i(destroySubWindows)))
	runtime.KeepAlive(this)
}

func QKeySequenceEdit_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QKeySequenceEdit_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QKeySequenceEdit_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QKeySequenceEdit_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QKeySequenceEdit) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QKeySequenceEdit_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QKeySequenceEdit) EventDefault(arg1 *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QKeySequenceEdit_event_default, this.h, uintptr(arg1.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	return __r != 0
}

func (this *QKeySequenceEdit) KeyPressEventDefault(arg1 *QKeyEvent) {
	purego.SyscallN(addr_QKeySequenceEdit_keyPressEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QKeySequenceEdit) KeyReleaseEventDefault(arg1 *QKeyEvent) {
	purego.SyscallN(addr_QKeySequenceEdit_keyReleaseEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QKeySequenceEdit) TimerEventDefault(arg1 *QTimerEvent) {
	purego.SyscallN(addr_QKeySequenceEdit_timerEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QKeySequenceEdit) FocusOutEventDefault(arg1 *QFocusEvent) {
	purego.SyscallN(addr_QKeySequenceEdit_focusOutEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QKeySequenceEdit) DevTypeDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QKeySequenceEdit_devType_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QKeySequenceEdit) SetVisibleDefault(visible bool) {
	purego.SyscallN(addr_QKeySequenceEdit_setVisible_default, this.h, uintptr(b2i(visible)))
	runtime.KeepAlive(this)
}

func (this *QKeySequenceEdit) SizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QKeySequenceEdit_sizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QKeySequenceEdit) MinimumSizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QKeySequenceEdit_minimumSizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QKeySequenceEdit) HeightForWidthDefault(arg1 int) int {
	__r0, _, _ := purego.SyscallN(addr_QKeySequenceEdit_heightForWidth_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QKeySequenceEdit) HasHeightForWidthDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QKeySequenceEdit_hasHeightForWidth_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QKeySequenceEdit) PaintEngineDefault() *QPaintEngine {
	__r, _, _ := purego.SyscallN(addr_QKeySequenceEdit_paintEngine_default, this.h)
	runtime.KeepAlive(this)
	return bindQPaintEngine(__r)
}

func (this *QKeySequenceEdit) MousePressEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QKeySequenceEdit_mousePressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QKeySequenceEdit) MouseReleaseEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QKeySequenceEdit_mouseReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QKeySequenceEdit) MouseDoubleClickEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QKeySequenceEdit_mouseDoubleClickEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QKeySequenceEdit) MouseMoveEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QKeySequenceEdit_mouseMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QKeySequenceEdit) WheelEventDefault(event *QWheelEvent) {
	purego.SyscallN(addr_QKeySequenceEdit_wheelEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QKeySequenceEdit) FocusInEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QKeySequenceEdit_focusInEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QKeySequenceEdit) EnterEventDefault(event *QEnterEvent) {
	purego.SyscallN(addr_QKeySequenceEdit_enterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QKeySequenceEdit) LeaveEventDefault(event *QEvent) {
	purego.SyscallN(addr_QKeySequenceEdit_leaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QKeySequenceEdit) PaintEventDefault(event *QPaintEvent) {
	purego.SyscallN(addr_QKeySequenceEdit_paintEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QKeySequenceEdit) MoveEventDefault(event *QMoveEvent) {
	purego.SyscallN(addr_QKeySequenceEdit_moveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QKeySequenceEdit) ResizeEventDefault(event *QResizeEvent) {
	purego.SyscallN(addr_QKeySequenceEdit_resizeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QKeySequenceEdit) CloseEventDefault(event *QCloseEvent) {
	purego.SyscallN(addr_QKeySequenceEdit_closeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QKeySequenceEdit) ContextMenuEventDefault(event *QContextMenuEvent) {
	purego.SyscallN(addr_QKeySequenceEdit_contextMenuEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QKeySequenceEdit) TabletEventDefault(event *QTabletEvent) {
	purego.SyscallN(addr_QKeySequenceEdit_tabletEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QKeySequenceEdit) ActionEventDefault(event *QActionEvent) {
	purego.SyscallN(addr_QKeySequenceEdit_actionEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QKeySequenceEdit) DragEnterEventDefault(event *QDragEnterEvent) {
	purego.SyscallN(addr_QKeySequenceEdit_dragEnterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QKeySequenceEdit) DragMoveEventDefault(event *QDragMoveEvent) {
	purego.SyscallN(addr_QKeySequenceEdit_dragMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QKeySequenceEdit) DragLeaveEventDefault(event *QDragLeaveEvent) {
	purego.SyscallN(addr_QKeySequenceEdit_dragLeaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QKeySequenceEdit) DropEventDefault(event *QDropEvent) {
	purego.SyscallN(addr_QKeySequenceEdit_dropEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QKeySequenceEdit) ShowEventDefault(event *QShowEvent) {
	purego.SyscallN(addr_QKeySequenceEdit_showEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QKeySequenceEdit) HideEventDefault(event *QHideEvent) {
	purego.SyscallN(addr_QKeySequenceEdit_hideEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QKeySequenceEdit) ChangeEventDefault(arg1 *QEvent) {
	purego.SyscallN(addr_QKeySequenceEdit_changeEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QKeySequenceEdit) MetricDefault(arg1 QPaintDevice__PaintDeviceMetric) int {
	__r0, _, _ := purego.SyscallN(addr_QKeySequenceEdit_metric_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QKeySequenceEdit) InitPainterDefault(painter *QPainter) {
	purego.SyscallN(addr_QKeySequenceEdit_initPainter_default, this.h, uintptr(painter.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
}

func (this *QKeySequenceEdit) RedirectedDefault(offset *QPoint) *QPaintDevice {
	__r, _, _ := purego.SyscallN(addr_QKeySequenceEdit_redirected_default, this.h, uintptr(offset.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(offset)
	return bindQPaintDevice(__r)
}

func (this *QKeySequenceEdit) SharedPainterDefault() *QPainter {
	__r, _, _ := purego.SyscallN(addr_QKeySequenceEdit_sharedPainter_default, this.h)
	runtime.KeepAlive(this)
	return bindQPainter(__r)
}

func (this *QKeySequenceEdit) InputMethodEventDefault(arg1 *QInputMethodEvent) {
	purego.SyscallN(addr_QKeySequenceEdit_inputMethodEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QKeySequenceEdit) InputMethodQueryDefault(arg1 InputMethodQuery) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QKeySequenceEdit_inputMethodQuery_default, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QKeySequenceEdit) FocusNextPrevChildDefault(next bool) bool {
	__r0, _, _ := purego.SyscallN(addr_QKeySequenceEdit_focusNextPrevChild_default, this.h, uintptr(b2i(next)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QKeySequenceEdit) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QKeySequenceEdit_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QKeySequenceEdit) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QKeySequenceEdit_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QKeySequenceEdit) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QKeySequenceEdit_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QKeySequenceEdit) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QKeySequenceEdit_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QKeySequenceEdit) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QKeySequenceEdit_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QKeySequenceEdit) OnEditingFinished(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QKeySequenceEdit_connect_editingFinished(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QKeySequenceEdit) OnKeySequenceChanged(fn func(*QKeySequence)) int64 {
	id := registerCallback(func(keySequence uintptr) { fn(UnsafeNewQKeySequence(unsafe.Pointer(keySequence))) })
	pg_QKeySequenceEdit_connect_keySequenceChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QKeySequenceEdit) OnWindowTitleChanged(fn func(string)) int64 {
	id := registerCallback(func(title uintptr, title_len int64) { fn(pgStrNoFree(title, title_len)) })
	pg_QKeySequenceEdit_connect_windowTitleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QKeySequenceEdit) OnWindowIconChanged(fn func(*QIcon)) int64 {
	id := registerCallback(func(icon uintptr) { fn(UnsafeNewQIcon(unsafe.Pointer(icon))) })
	pg_QKeySequenceEdit_connect_windowIconChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QKeySequenceEdit) OnWindowIconTextChanged(fn func(string)) int64 {
	id := registerCallback(func(iconText uintptr, iconText_len int64) { fn(pgStrNoFree(iconText, iconText_len)) })
	pg_QKeySequenceEdit_connect_windowIconTextChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QKeySequenceEdit) OnCustomContextMenuRequested(fn func(*QPoint)) int64 {
	id := registerCallback(func(pos uintptr) { fn(UnsafeNewQPoint(unsafe.Pointer(pos))) })
	pg_QKeySequenceEdit_connect_customContextMenuRequested(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QKeySequenceEdit) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QKeySequenceEdit_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QKeySequenceEdit) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QKeySequenceEdit_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QKeySequenceEdit) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QKeySequenceEdit_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QKeySequenceEdit) Disconnect(handle int64) {
	pg_QKeySequenceEdit__disconnect(handle)
	unregisterCallback(handle)
}

func QKeySequenceEdit_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case 1237236456312824991:
		return 1
	case 1023791977679793185:
		return 2
	case 772598475950050299:
		return 3
	case -5409927581812374498:
		return 4
	case -385596010576048993:
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
	case 975043767033552552:
		return 18
	case 6316270346668237895:
		return 19
	case -551070009623711640:
		return 20
	case 5290442163167458465:
		return 21
	case -9161949150489525396:
		return 22
	case -4302917285655014661:
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
	case -1185315367526857833:
		return 44
	case 3769868229123868474:
		return 45
	case 7379029982170503008:
		return 46
	case -568369772982215962:
		return 47
	}
	return -1
}

func QKeySequenceEdit_enable(h uintptr, idx int32) { pg_QKeySequenceEdit__set_override(h, idx) }

func (this *QKeySequenceEdit) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QKeySequenceEdit) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(arg1 uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(arg1)))) }) // slot "event"
}

func (this *QKeySequenceEdit) OnKeyPressEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 1023791977679793185, func(arg1 uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(arg1))) }) // slot "keyPressEvent"
}

func (this *QKeySequenceEdit) OnKeyReleaseEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 772598475950050299, func(arg1 uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(arg1))) }) // slot "keyReleaseEvent"
}

func (this *QKeySequenceEdit) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(arg1 uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(arg1))) }) // slot "timerEvent"
}

func (this *QKeySequenceEdit) OnFocusOutEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), -385596010576048993, func(arg1 uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(arg1))) }) // slot "focusOutEvent"
}

func (this *QKeySequenceEdit) OnDevType(fn func() int) {
	registerVirt(this.UnsafeGoID(), -8659679893868524044, func() int32 { return int32(fn()) }) // slot "devType"
}

func (this *QKeySequenceEdit) OnSetVisible(fn func(bool)) {
	registerVirt(this.UnsafeGoID(), 2110182719375245575, func(visible int32) { fn(visible != 0) }) // slot "setVisible"
}

func (this *QKeySequenceEdit) OnSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 729690261457684101, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sizeHint"
}

func (this *QKeySequenceEdit) OnMinimumSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 7049264599837532491, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "minimumSizeHint"
}

func (this *QKeySequenceEdit) OnHeightForWidth(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), 4311534060495737339, func(arg1 int32) int32 { return int32(fn(int(arg1))) }) // slot "heightForWidth"
}

func (this *QKeySequenceEdit) OnHasHeightForWidth(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -8266762397995423915, func() int32 { return b2i(fn()) }) // slot "hasHeightForWidth"
}

func (this *QKeySequenceEdit) OnPaintEngine(fn func() *QPaintEngine) {
	registerVirt(this.UnsafeGoID(), 6398007404756346381, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "paintEngine"
}

func (this *QKeySequenceEdit) OnMousePressEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -332090898078543319, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mousePressEvent"
}

func (this *QKeySequenceEdit) OnMouseReleaseEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4415920049297795213, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseReleaseEvent"
}

func (this *QKeySequenceEdit) OnMouseDoubleClickEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4029073905697023887, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseDoubleClickEvent"
}

func (this *QKeySequenceEdit) OnMouseMoveEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4255494004099806581, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseMoveEvent"
}

func (this *QKeySequenceEdit) OnWheelEvent(fn func(*QWheelEvent)) {
	registerVirt(this.UnsafeGoID(), 1919988340934511316, func(event uintptr) { fn(UnsafeNewQWheelEvent(unsafe.Pointer(event))) }) // slot "wheelEvent"
}

func (this *QKeySequenceEdit) OnFocusInEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), 975043767033552552, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusInEvent"
}

func (this *QKeySequenceEdit) OnEnterEvent(fn func(*QEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 6316270346668237895, func(event uintptr) { fn(UnsafeNewQEnterEvent(unsafe.Pointer(event))) }) // slot "enterEvent"
}

func (this *QKeySequenceEdit) OnLeaveEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), -551070009623711640, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "leaveEvent"
}

func (this *QKeySequenceEdit) OnPaintEvent(fn func(*QPaintEvent)) {
	registerVirt(this.UnsafeGoID(), 5290442163167458465, func(event uintptr) { fn(UnsafeNewQPaintEvent(unsafe.Pointer(event))) }) // slot "paintEvent"
}

func (this *QKeySequenceEdit) OnMoveEvent(fn func(*QMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -9161949150489525396, func(event uintptr) { fn(UnsafeNewQMoveEvent(unsafe.Pointer(event))) }) // slot "moveEvent"
}

func (this *QKeySequenceEdit) OnResizeEvent(fn func(*QResizeEvent)) {
	registerVirt(this.UnsafeGoID(), -4302917285655014661, func(event uintptr) { fn(UnsafeNewQResizeEvent(unsafe.Pointer(event))) }) // slot "resizeEvent"
}

func (this *QKeySequenceEdit) OnCloseEvent(fn func(*QCloseEvent)) {
	registerVirt(this.UnsafeGoID(), -1989937050019464783, func(event uintptr) { fn(UnsafeNewQCloseEvent(unsafe.Pointer(event))) }) // slot "closeEvent"
}

func (this *QKeySequenceEdit) OnContextMenuEvent(fn func(*QContextMenuEvent)) {
	registerVirt(this.UnsafeGoID(), -6698043222227006015, func(event uintptr) { fn(UnsafeNewQContextMenuEvent(unsafe.Pointer(event))) }) // slot "contextMenuEvent"
}

func (this *QKeySequenceEdit) OnTabletEvent(fn func(*QTabletEvent)) {
	registerVirt(this.UnsafeGoID(), -212068866860701141, func(event uintptr) { fn(UnsafeNewQTabletEvent(unsafe.Pointer(event))) }) // slot "tabletEvent"
}

func (this *QKeySequenceEdit) OnActionEvent(fn func(*QActionEvent)) {
	registerVirt(this.UnsafeGoID(), 1697444140088555581, func(event uintptr) { fn(UnsafeNewQActionEvent(unsafe.Pointer(event))) }) // slot "actionEvent"
}

func (this *QKeySequenceEdit) OnDragEnterEvent(fn func(*QDragEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 206122715919917589, func(event uintptr) { fn(UnsafeNewQDragEnterEvent(unsafe.Pointer(event))) }) // slot "dragEnterEvent"
}

func (this *QKeySequenceEdit) OnDragMoveEvent(fn func(*QDragMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -1781545137352745978, func(event uintptr) { fn(UnsafeNewQDragMoveEvent(unsafe.Pointer(event))) }) // slot "dragMoveEvent"
}

func (this *QKeySequenceEdit) OnDragLeaveEvent(fn func(*QDragLeaveEvent)) {
	registerVirt(this.UnsafeGoID(), -3139904158373725006, func(event uintptr) { fn(UnsafeNewQDragLeaveEvent(unsafe.Pointer(event))) }) // slot "dragLeaveEvent"
}

func (this *QKeySequenceEdit) OnDropEvent(fn func(*QDropEvent)) {
	registerVirt(this.UnsafeGoID(), 9107058682754591716, func(event uintptr) { fn(UnsafeNewQDropEvent(unsafe.Pointer(event))) }) // slot "dropEvent"
}

func (this *QKeySequenceEdit) OnShowEvent(fn func(*QShowEvent)) {
	registerVirt(this.UnsafeGoID(), -6825369174751916, func(event uintptr) { fn(UnsafeNewQShowEvent(unsafe.Pointer(event))) }) // slot "showEvent"
}

func (this *QKeySequenceEdit) OnHideEvent(fn func(*QHideEvent)) {
	registerVirt(this.UnsafeGoID(), -199718707876719129, func(event uintptr) { fn(UnsafeNewQHideEvent(unsafe.Pointer(event))) }) // slot "hideEvent"
}

func (this *QKeySequenceEdit) OnNativeEvent(fn func(*QByteArray, unsafe.Pointer, *int64) bool) {
	registerVirt(this.UnsafeGoID(), -2039166586246443706, func(eventType uintptr, message uintptr, result uintptr) int32 { return b2i(fn(UnsafeNewQByteArray(unsafe.Pointer(eventType)), unsafe.Pointer(message), (*int64)(unsafe.Pointer(result)))) }) // slot "nativeEvent"
}

func (this *QKeySequenceEdit) OnChangeEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 4959376063826165943, func(arg1 uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(arg1))) }) // slot "changeEvent"
}

func (this *QKeySequenceEdit) OnMetric(fn func(QPaintDevice__PaintDeviceMetric) int) {
	registerVirt(this.UnsafeGoID(), 427629649267311777, func(arg1 int32) int32 { return int32(fn(QPaintDevice__PaintDeviceMetric(arg1))) }) // slot "metric"
}

func (this *QKeySequenceEdit) OnInitPainter(fn func(*QPainter)) {
	registerVirt(this.UnsafeGoID(), 6477500296980050210, func(painter uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter))) }) // slot "initPainter"
}

func (this *QKeySequenceEdit) OnRedirected(fn func(*QPoint) *QPaintDevice) {
	registerVirt(this.UnsafeGoID(), 4127240633310655048, func(offset uintptr) uintptr { return uintptr(fn(UnsafeNewQPoint(unsafe.Pointer(offset))).UnsafePointer()) }) // slot "redirected"
}

func (this *QKeySequenceEdit) OnSharedPainter(fn func() *QPainter) {
	registerVirt(this.UnsafeGoID(), 5792771328741774991, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sharedPainter"
}

func (this *QKeySequenceEdit) OnInputMethodEvent(fn func(*QInputMethodEvent)) {
	registerVirt(this.UnsafeGoID(), 5438819661942843222, func(arg1 uintptr) { fn(UnsafeNewQInputMethodEvent(unsafe.Pointer(arg1))) }) // slot "inputMethodEvent"
}

func (this *QKeySequenceEdit) OnInputMethodQuery(fn func(InputMethodQuery) *QVariant) {
	registerVirt(this.UnsafeGoID(), 7345853097340709638, func(arg1 int32) uintptr { return uintptr(fn(InputMethodQuery(arg1)).UnsafePointer()) }) // slot "inputMethodQuery"
}

func (this *QKeySequenceEdit) OnFocusNextPrevChild(fn func(bool) bool) {
	registerVirt(this.UnsafeGoID(), 5898328517505784191, func(next int32) int32 { return b2i(fn(next != 0)) }) // slot "focusNextPrevChild"
}

func (this *QKeySequenceEdit) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QKeySequenceEdit) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QKeySequenceEdit) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QKeySequenceEdit) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QKeySequenceEdit) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QKeySequenceEdit) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QKeySequenceEdit_delete(this.h)
	}
	this.UnsafeRelease()
}
