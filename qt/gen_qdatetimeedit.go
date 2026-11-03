package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QDateTimeEdit_chain uintptr
	addr_QDateTimeEdit_new uintptr
	addr_QDateTimeEdit_new2 uintptr
	addr_QDateTimeEdit_new3 uintptr
	addr_QDateTimeEdit_new4 uintptr
	addr_QDateTimeEdit_new5 uintptr
	addr_QDateTimeEdit_new6 uintptr
	addr_QDateTimeEdit_new7 uintptr
	addr_QDateTimeEdit_new8 uintptr
	addr_QDateTimeEdit_tr uintptr
	addr_QDateTimeEdit_dateTime uintptr
	addr_QDateTimeEdit_date uintptr
	addr_QDateTimeEdit_time uintptr
	addr_QDateTimeEdit_calendar uintptr
	addr_QDateTimeEdit_setCalendar uintptr
	addr_QDateTimeEdit_minimumDateTime uintptr
	addr_QDateTimeEdit_clearMinimumDateTime uintptr
	addr_QDateTimeEdit_setMinimumDateTime uintptr
	addr_QDateTimeEdit_maximumDateTime uintptr
	addr_QDateTimeEdit_clearMaximumDateTime uintptr
	addr_QDateTimeEdit_setMaximumDateTime uintptr
	addr_QDateTimeEdit_setDateTimeRange uintptr
	addr_QDateTimeEdit_minimumDate uintptr
	addr_QDateTimeEdit_setMinimumDate uintptr
	addr_QDateTimeEdit_clearMinimumDate uintptr
	addr_QDateTimeEdit_maximumDate uintptr
	addr_QDateTimeEdit_setMaximumDate uintptr
	addr_QDateTimeEdit_clearMaximumDate uintptr
	addr_QDateTimeEdit_setDateRange uintptr
	addr_QDateTimeEdit_minimumTime uintptr
	addr_QDateTimeEdit_setMinimumTime uintptr
	addr_QDateTimeEdit_clearMinimumTime uintptr
	addr_QDateTimeEdit_maximumTime uintptr
	addr_QDateTimeEdit_setMaximumTime uintptr
	addr_QDateTimeEdit_clearMaximumTime uintptr
	addr_QDateTimeEdit_setTimeRange uintptr
	addr_QDateTimeEdit_displayedSections uintptr
	addr_QDateTimeEdit_currentSection uintptr
	addr_QDateTimeEdit_sectionAt uintptr
	addr_QDateTimeEdit_setCurrentSection uintptr
	addr_QDateTimeEdit_currentSectionIndex uintptr
	addr_QDateTimeEdit_setCurrentSectionIndex uintptr
	addr_QDateTimeEdit_calendarWidget uintptr
	addr_QDateTimeEdit_setCalendarWidget uintptr
	addr_QDateTimeEdit_sectionCount uintptr
	addr_QDateTimeEdit_setSelectedSection uintptr
	addr_QDateTimeEdit_sectionText uintptr
	addr_QDateTimeEdit_displayFormat uintptr
	addr_QDateTimeEdit_setDisplayFormat uintptr
	addr_QDateTimeEdit_calendarPopup uintptr
	addr_QDateTimeEdit_setCalendarPopup uintptr
	addr_QDateTimeEdit_timeSpec uintptr
	addr_QDateTimeEdit_setTimeSpec uintptr
	addr_QDateTimeEdit_timeZone uintptr
	addr_QDateTimeEdit_setTimeZone uintptr
	addr_QDateTimeEdit_dateTimeChanged uintptr
	addr_QDateTimeEdit_timeChanged uintptr
	addr_QDateTimeEdit_dateChanged uintptr
	addr_QDateTimeEdit_setDateTime uintptr
	addr_QDateTimeEdit_setDate uintptr
	addr_QDateTimeEdit_setTime uintptr
	addr_QDateTimeEdit_lineEdit uintptr
	addr_QDateTimeEdit_setLineEdit uintptr
	addr_QDateTimeEdit_setTabOrder uintptr
	addr_QDateTimeEdit_mouseGrabber uintptr
	addr_QDateTimeEdit_keyboardGrabber uintptr
	addr_QDateTimeEdit_find uintptr
	addr_QDateTimeEdit_createWindowContainer uintptr
	addr_QDateTimeEdit_updateMicroFocus uintptr
	addr_QDateTimeEdit_create uintptr
	addr_QDateTimeEdit_destroy uintptr
	addr_QDateTimeEdit_focusNextChild uintptr
	addr_QDateTimeEdit_focusPreviousChild uintptr
	addr_QDateTimeEdit_connect uintptr
	addr_QDateTimeEdit_connect2 uintptr
	addr_QDateTimeEdit_sender uintptr
	addr_QDateTimeEdit_senderSignalIndex uintptr
	addr_QDateTimeEdit_receivers uintptr
	addr_QDateTimeEdit_isSignalConnected uintptr
	pg_QDateTimeEdit_devicePixelRatioFScale func() float64
	pg_QDateTimeEdit_encodeMetricF func(int32, float64) int32
	pg_QDateTimeEdit_getDecodedMetricF func(uintptr, int32, int32) float64
	addr_QDateTimeEdit_tr2 uintptr
	addr_QDateTimeEdit_tr3 uintptr
	addr_QDateTimeEdit_createWindowContainer2 uintptr
	addr_QDateTimeEdit_createWindowContainer3 uintptr
	addr_QDateTimeEdit_updateMicroFocus2 uintptr
	addr_QDateTimeEdit_create2 uintptr
	addr_QDateTimeEdit_create3 uintptr
	addr_QDateTimeEdit_create4 uintptr
	addr_QDateTimeEdit_destroy2 uintptr
	addr_QDateTimeEdit_destroy3 uintptr
	addr_QDateTimeEdit_connect3 uintptr
	addr_QDateTimeEdit_connect4 uintptr
	addr_QDateTimeEdit_metaObject_default uintptr
	addr_QDateTimeEdit_sizeHint_default uintptr
	addr_QDateTimeEdit_clear_default uintptr
	addr_QDateTimeEdit_stepBy_default uintptr
	addr_QDateTimeEdit_event_default uintptr
	addr_QDateTimeEdit_keyPressEvent_default uintptr
	addr_QDateTimeEdit_wheelEvent_default uintptr
	addr_QDateTimeEdit_focusInEvent_default uintptr
	addr_QDateTimeEdit_focusNextPrevChild_default uintptr
	addr_QDateTimeEdit_validate_default uintptr
	addr_QDateTimeEdit_fixup_default uintptr
	addr_QDateTimeEdit_dateTimeFromText_default uintptr
	addr_QDateTimeEdit_textFromDateTime_default uintptr
	addr_QDateTimeEdit_stepEnabled_default uintptr
	addr_QDateTimeEdit_mousePressEvent_default uintptr
	addr_QDateTimeEdit_paintEvent_default uintptr
	addr_QDateTimeEdit_initStyleOption_default uintptr
	addr_QDateTimeEdit_minimumSizeHint_default uintptr
	addr_QDateTimeEdit_inputMethodQuery_default uintptr
	addr_QDateTimeEdit_resizeEvent_default uintptr
	addr_QDateTimeEdit_keyReleaseEvent_default uintptr
	addr_QDateTimeEdit_focusOutEvent_default uintptr
	addr_QDateTimeEdit_contextMenuEvent_default uintptr
	addr_QDateTimeEdit_changeEvent_default uintptr
	addr_QDateTimeEdit_closeEvent_default uintptr
	addr_QDateTimeEdit_hideEvent_default uintptr
	addr_QDateTimeEdit_mouseReleaseEvent_default uintptr
	addr_QDateTimeEdit_mouseMoveEvent_default uintptr
	addr_QDateTimeEdit_timerEvent_default uintptr
	addr_QDateTimeEdit_showEvent_default uintptr
	addr_QDateTimeEdit_devType_default uintptr
	addr_QDateTimeEdit_setVisible_default uintptr
	addr_QDateTimeEdit_heightForWidth_default uintptr
	addr_QDateTimeEdit_hasHeightForWidth_default uintptr
	addr_QDateTimeEdit_paintEngine_default uintptr
	addr_QDateTimeEdit_mouseDoubleClickEvent_default uintptr
	addr_QDateTimeEdit_enterEvent_default uintptr
	addr_QDateTimeEdit_leaveEvent_default uintptr
	addr_QDateTimeEdit_moveEvent_default uintptr
	addr_QDateTimeEdit_tabletEvent_default uintptr
	addr_QDateTimeEdit_actionEvent_default uintptr
	addr_QDateTimeEdit_dragEnterEvent_default uintptr
	addr_QDateTimeEdit_dragMoveEvent_default uintptr
	addr_QDateTimeEdit_dragLeaveEvent_default uintptr
	addr_QDateTimeEdit_dropEvent_default uintptr
	addr_QDateTimeEdit_metric_default uintptr
	addr_QDateTimeEdit_initPainter_default uintptr
	addr_QDateTimeEdit_redirected_default uintptr
	addr_QDateTimeEdit_sharedPainter_default uintptr
	addr_QDateTimeEdit_inputMethodEvent_default uintptr
	addr_QDateTimeEdit_eventFilter_default uintptr
	addr_QDateTimeEdit_childEvent_default uintptr
	addr_QDateTimeEdit_customEvent_default uintptr
	addr_QDateTimeEdit_connectNotify_default uintptr
	addr_QDateTimeEdit_disconnectNotify_default uintptr
	pg_QDateTimeEdit_connect_dateTimeChanged func(uintptr, int64)
	pg_QDateTimeEdit_connect_timeChanged func(uintptr, int64)
	pg_QDateTimeEdit_connect_dateChanged func(uintptr, int64)
	pg_QDateTimeEdit_connect_editingFinished func(uintptr, int64)
	pg_QDateTimeEdit_connect_returnPressed func(uintptr, int64)
	pg_QDateTimeEdit_connect_windowTitleChanged func(uintptr, int64)
	pg_QDateTimeEdit_connect_windowIconChanged func(uintptr, int64)
	pg_QDateTimeEdit_connect_windowIconTextChanged func(uintptr, int64)
	pg_QDateTimeEdit_connect_customContextMenuRequested func(uintptr, int64)
	pg_QDateTimeEdit_connect_destroyed func(uintptr, int64)
	pg_QDateTimeEdit_connect_destroyed2 func(uintptr, int64)
	pg_QDateTimeEdit_connect_objectNameChanged func(uintptr, int64)
	pg_QDateTimeEdit__disconnect func(int64)
	pg_QDateTimeEdit__set_override func(uintptr, int32)
	pg_QDateTimeEdit_delete func(uintptr)
)

func init() {
	symaddr(&addr_QDateTimeEdit_chain, "QDateTimeEdit_chain")
	symaddr(&addr_QDateTimeEdit_new, "QDateTimeEdit_new")
	symaddr(&addr_QDateTimeEdit_new2, "QDateTimeEdit_new2")
	symaddr(&addr_QDateTimeEdit_new3, "QDateTimeEdit_new3")
	symaddr(&addr_QDateTimeEdit_new4, "QDateTimeEdit_new4")
	symaddr(&addr_QDateTimeEdit_new5, "QDateTimeEdit_new5")
	symaddr(&addr_QDateTimeEdit_new6, "QDateTimeEdit_new6")
	symaddr(&addr_QDateTimeEdit_new7, "QDateTimeEdit_new7")
	symaddr(&addr_QDateTimeEdit_new8, "QDateTimeEdit_new8")
	symaddr(&addr_QDateTimeEdit_tr, "QDateTimeEdit_tr")
	symaddr(&addr_QDateTimeEdit_dateTime, "QDateTimeEdit_dateTime")
	symaddr(&addr_QDateTimeEdit_date, "QDateTimeEdit_date")
	symaddr(&addr_QDateTimeEdit_time, "QDateTimeEdit_time")
	symaddr(&addr_QDateTimeEdit_calendar, "QDateTimeEdit_calendar")
	symaddr(&addr_QDateTimeEdit_setCalendar, "QDateTimeEdit_setCalendar")
	symaddr(&addr_QDateTimeEdit_minimumDateTime, "QDateTimeEdit_minimumDateTime")
	symaddr(&addr_QDateTimeEdit_clearMinimumDateTime, "QDateTimeEdit_clearMinimumDateTime")
	symaddr(&addr_QDateTimeEdit_setMinimumDateTime, "QDateTimeEdit_setMinimumDateTime")
	symaddr(&addr_QDateTimeEdit_maximumDateTime, "QDateTimeEdit_maximumDateTime")
	symaddr(&addr_QDateTimeEdit_clearMaximumDateTime, "QDateTimeEdit_clearMaximumDateTime")
	symaddr(&addr_QDateTimeEdit_setMaximumDateTime, "QDateTimeEdit_setMaximumDateTime")
	symaddr(&addr_QDateTimeEdit_setDateTimeRange, "QDateTimeEdit_setDateTimeRange")
	symaddr(&addr_QDateTimeEdit_minimumDate, "QDateTimeEdit_minimumDate")
	symaddr(&addr_QDateTimeEdit_setMinimumDate, "QDateTimeEdit_setMinimumDate")
	symaddr(&addr_QDateTimeEdit_clearMinimumDate, "QDateTimeEdit_clearMinimumDate")
	symaddr(&addr_QDateTimeEdit_maximumDate, "QDateTimeEdit_maximumDate")
	symaddr(&addr_QDateTimeEdit_setMaximumDate, "QDateTimeEdit_setMaximumDate")
	symaddr(&addr_QDateTimeEdit_clearMaximumDate, "QDateTimeEdit_clearMaximumDate")
	symaddr(&addr_QDateTimeEdit_setDateRange, "QDateTimeEdit_setDateRange")
	symaddr(&addr_QDateTimeEdit_minimumTime, "QDateTimeEdit_minimumTime")
	symaddr(&addr_QDateTimeEdit_setMinimumTime, "QDateTimeEdit_setMinimumTime")
	symaddr(&addr_QDateTimeEdit_clearMinimumTime, "QDateTimeEdit_clearMinimumTime")
	symaddr(&addr_QDateTimeEdit_maximumTime, "QDateTimeEdit_maximumTime")
	symaddr(&addr_QDateTimeEdit_setMaximumTime, "QDateTimeEdit_setMaximumTime")
	symaddr(&addr_QDateTimeEdit_clearMaximumTime, "QDateTimeEdit_clearMaximumTime")
	symaddr(&addr_QDateTimeEdit_setTimeRange, "QDateTimeEdit_setTimeRange")
	symaddr(&addr_QDateTimeEdit_displayedSections, "QDateTimeEdit_displayedSections")
	symaddr(&addr_QDateTimeEdit_currentSection, "QDateTimeEdit_currentSection")
	symaddr(&addr_QDateTimeEdit_sectionAt, "QDateTimeEdit_sectionAt")
	symaddr(&addr_QDateTimeEdit_setCurrentSection, "QDateTimeEdit_setCurrentSection")
	symaddr(&addr_QDateTimeEdit_currentSectionIndex, "QDateTimeEdit_currentSectionIndex")
	symaddr(&addr_QDateTimeEdit_setCurrentSectionIndex, "QDateTimeEdit_setCurrentSectionIndex")
	symaddr(&addr_QDateTimeEdit_calendarWidget, "QDateTimeEdit_calendarWidget")
	symaddr(&addr_QDateTimeEdit_setCalendarWidget, "QDateTimeEdit_setCalendarWidget")
	symaddr(&addr_QDateTimeEdit_sectionCount, "QDateTimeEdit_sectionCount")
	symaddr(&addr_QDateTimeEdit_setSelectedSection, "QDateTimeEdit_setSelectedSection")
	symaddr(&addr_QDateTimeEdit_sectionText, "QDateTimeEdit_sectionText")
	symaddr(&addr_QDateTimeEdit_displayFormat, "QDateTimeEdit_displayFormat")
	symaddr(&addr_QDateTimeEdit_setDisplayFormat, "QDateTimeEdit_setDisplayFormat")
	symaddr(&addr_QDateTimeEdit_calendarPopup, "QDateTimeEdit_calendarPopup")
	symaddr(&addr_QDateTimeEdit_setCalendarPopup, "QDateTimeEdit_setCalendarPopup")
	symaddr(&addr_QDateTimeEdit_timeSpec, "QDateTimeEdit_timeSpec")
	symaddr(&addr_QDateTimeEdit_setTimeSpec, "QDateTimeEdit_setTimeSpec")
	symaddr(&addr_QDateTimeEdit_timeZone, "QDateTimeEdit_timeZone")
	symaddr(&addr_QDateTimeEdit_setTimeZone, "QDateTimeEdit_setTimeZone")
	symaddr(&addr_QDateTimeEdit_dateTimeChanged, "QDateTimeEdit_dateTimeChanged")
	symaddr(&addr_QDateTimeEdit_timeChanged, "QDateTimeEdit_timeChanged")
	symaddr(&addr_QDateTimeEdit_dateChanged, "QDateTimeEdit_dateChanged")
	symaddr(&addr_QDateTimeEdit_setDateTime, "QDateTimeEdit_setDateTime")
	symaddr(&addr_QDateTimeEdit_setDate, "QDateTimeEdit_setDate")
	symaddr(&addr_QDateTimeEdit_setTime, "QDateTimeEdit_setTime")
	symaddr(&addr_QDateTimeEdit_lineEdit, "QDateTimeEdit_lineEdit")
	symaddr(&addr_QDateTimeEdit_setLineEdit, "QDateTimeEdit_setLineEdit")
	symaddr(&addr_QDateTimeEdit_setTabOrder, "QDateTimeEdit_setTabOrder")
	symaddr(&addr_QDateTimeEdit_mouseGrabber, "QDateTimeEdit_mouseGrabber")
	symaddr(&addr_QDateTimeEdit_keyboardGrabber, "QDateTimeEdit_keyboardGrabber")
	symaddr(&addr_QDateTimeEdit_find, "QDateTimeEdit_find")
	symaddr(&addr_QDateTimeEdit_createWindowContainer, "QDateTimeEdit_createWindowContainer")
	symaddr(&addr_QDateTimeEdit_updateMicroFocus, "QDateTimeEdit_updateMicroFocus")
	symaddr(&addr_QDateTimeEdit_create, "QDateTimeEdit_create")
	symaddr(&addr_QDateTimeEdit_destroy, "QDateTimeEdit_destroy")
	symaddr(&addr_QDateTimeEdit_focusNextChild, "QDateTimeEdit_focusNextChild")
	symaddr(&addr_QDateTimeEdit_focusPreviousChild, "QDateTimeEdit_focusPreviousChild")
	symaddr(&addr_QDateTimeEdit_connect, "QDateTimeEdit_connect")
	symaddr(&addr_QDateTimeEdit_connect2, "QDateTimeEdit_connect2")
	symaddr(&addr_QDateTimeEdit_sender, "QDateTimeEdit_sender")
	symaddr(&addr_QDateTimeEdit_senderSignalIndex, "QDateTimeEdit_senderSignalIndex")
	symaddr(&addr_QDateTimeEdit_receivers, "QDateTimeEdit_receivers")
	symaddr(&addr_QDateTimeEdit_isSignalConnected, "QDateTimeEdit_isSignalConnected")
	reg(&pg_QDateTimeEdit_devicePixelRatioFScale, "QDateTimeEdit_devicePixelRatioFScale")
	reg(&pg_QDateTimeEdit_encodeMetricF, "QDateTimeEdit_encodeMetricF")
	reg(&pg_QDateTimeEdit_getDecodedMetricF, "QDateTimeEdit_getDecodedMetricF")
	symaddr(&addr_QDateTimeEdit_tr2, "QDateTimeEdit_tr2")
	symaddr(&addr_QDateTimeEdit_tr3, "QDateTimeEdit_tr3")
	symaddr(&addr_QDateTimeEdit_createWindowContainer2, "QDateTimeEdit_createWindowContainer2")
	symaddr(&addr_QDateTimeEdit_createWindowContainer3, "QDateTimeEdit_createWindowContainer3")
	symaddr(&addr_QDateTimeEdit_updateMicroFocus2, "QDateTimeEdit_updateMicroFocus2")
	symaddr(&addr_QDateTimeEdit_create2, "QDateTimeEdit_create2")
	symaddr(&addr_QDateTimeEdit_create3, "QDateTimeEdit_create3")
	symaddr(&addr_QDateTimeEdit_create4, "QDateTimeEdit_create4")
	symaddr(&addr_QDateTimeEdit_destroy2, "QDateTimeEdit_destroy2")
	symaddr(&addr_QDateTimeEdit_destroy3, "QDateTimeEdit_destroy3")
	symaddr(&addr_QDateTimeEdit_connect3, "QDateTimeEdit_connect3")
	symaddr(&addr_QDateTimeEdit_connect4, "QDateTimeEdit_connect4")
	symaddr(&addr_QDateTimeEdit_metaObject_default, "QDateTimeEdit_metaObject_default")
	symaddr(&addr_QDateTimeEdit_sizeHint_default, "QDateTimeEdit_sizeHint_default")
	symaddr(&addr_QDateTimeEdit_clear_default, "QDateTimeEdit_clear_default")
	symaddr(&addr_QDateTimeEdit_stepBy_default, "QDateTimeEdit_stepBy_default")
	symaddr(&addr_QDateTimeEdit_event_default, "QDateTimeEdit_event_default")
	symaddr(&addr_QDateTimeEdit_keyPressEvent_default, "QDateTimeEdit_keyPressEvent_default")
	symaddr(&addr_QDateTimeEdit_wheelEvent_default, "QDateTimeEdit_wheelEvent_default")
	symaddr(&addr_QDateTimeEdit_focusInEvent_default, "QDateTimeEdit_focusInEvent_default")
	symaddr(&addr_QDateTimeEdit_focusNextPrevChild_default, "QDateTimeEdit_focusNextPrevChild_default")
	symaddr(&addr_QDateTimeEdit_validate_default, "QDateTimeEdit_validate_default")
	symaddr(&addr_QDateTimeEdit_fixup_default, "QDateTimeEdit_fixup_default")
	symaddr(&addr_QDateTimeEdit_dateTimeFromText_default, "QDateTimeEdit_dateTimeFromText_default")
	symaddr(&addr_QDateTimeEdit_textFromDateTime_default, "QDateTimeEdit_textFromDateTime_default")
	symaddr(&addr_QDateTimeEdit_stepEnabled_default, "QDateTimeEdit_stepEnabled_default")
	symaddr(&addr_QDateTimeEdit_mousePressEvent_default, "QDateTimeEdit_mousePressEvent_default")
	symaddr(&addr_QDateTimeEdit_paintEvent_default, "QDateTimeEdit_paintEvent_default")
	symaddr(&addr_QDateTimeEdit_initStyleOption_default, "QDateTimeEdit_initStyleOption_default")
	symaddr(&addr_QDateTimeEdit_minimumSizeHint_default, "QDateTimeEdit_minimumSizeHint_default")
	symaddr(&addr_QDateTimeEdit_inputMethodQuery_default, "QDateTimeEdit_inputMethodQuery_default")
	symaddr(&addr_QDateTimeEdit_resizeEvent_default, "QDateTimeEdit_resizeEvent_default")
	symaddr(&addr_QDateTimeEdit_keyReleaseEvent_default, "QDateTimeEdit_keyReleaseEvent_default")
	symaddr(&addr_QDateTimeEdit_focusOutEvent_default, "QDateTimeEdit_focusOutEvent_default")
	symaddr(&addr_QDateTimeEdit_contextMenuEvent_default, "QDateTimeEdit_contextMenuEvent_default")
	symaddr(&addr_QDateTimeEdit_changeEvent_default, "QDateTimeEdit_changeEvent_default")
	symaddr(&addr_QDateTimeEdit_closeEvent_default, "QDateTimeEdit_closeEvent_default")
	symaddr(&addr_QDateTimeEdit_hideEvent_default, "QDateTimeEdit_hideEvent_default")
	symaddr(&addr_QDateTimeEdit_mouseReleaseEvent_default, "QDateTimeEdit_mouseReleaseEvent_default")
	symaddr(&addr_QDateTimeEdit_mouseMoveEvent_default, "QDateTimeEdit_mouseMoveEvent_default")
	symaddr(&addr_QDateTimeEdit_timerEvent_default, "QDateTimeEdit_timerEvent_default")
	symaddr(&addr_QDateTimeEdit_showEvent_default, "QDateTimeEdit_showEvent_default")
	symaddr(&addr_QDateTimeEdit_devType_default, "QDateTimeEdit_devType_default")
	symaddr(&addr_QDateTimeEdit_setVisible_default, "QDateTimeEdit_setVisible_default")
	symaddr(&addr_QDateTimeEdit_heightForWidth_default, "QDateTimeEdit_heightForWidth_default")
	symaddr(&addr_QDateTimeEdit_hasHeightForWidth_default, "QDateTimeEdit_hasHeightForWidth_default")
	symaddr(&addr_QDateTimeEdit_paintEngine_default, "QDateTimeEdit_paintEngine_default")
	symaddr(&addr_QDateTimeEdit_mouseDoubleClickEvent_default, "QDateTimeEdit_mouseDoubleClickEvent_default")
	symaddr(&addr_QDateTimeEdit_enterEvent_default, "QDateTimeEdit_enterEvent_default")
	symaddr(&addr_QDateTimeEdit_leaveEvent_default, "QDateTimeEdit_leaveEvent_default")
	symaddr(&addr_QDateTimeEdit_moveEvent_default, "QDateTimeEdit_moveEvent_default")
	symaddr(&addr_QDateTimeEdit_tabletEvent_default, "QDateTimeEdit_tabletEvent_default")
	symaddr(&addr_QDateTimeEdit_actionEvent_default, "QDateTimeEdit_actionEvent_default")
	symaddr(&addr_QDateTimeEdit_dragEnterEvent_default, "QDateTimeEdit_dragEnterEvent_default")
	symaddr(&addr_QDateTimeEdit_dragMoveEvent_default, "QDateTimeEdit_dragMoveEvent_default")
	symaddr(&addr_QDateTimeEdit_dragLeaveEvent_default, "QDateTimeEdit_dragLeaveEvent_default")
	symaddr(&addr_QDateTimeEdit_dropEvent_default, "QDateTimeEdit_dropEvent_default")
	symaddr(&addr_QDateTimeEdit_metric_default, "QDateTimeEdit_metric_default")
	symaddr(&addr_QDateTimeEdit_initPainter_default, "QDateTimeEdit_initPainter_default")
	symaddr(&addr_QDateTimeEdit_redirected_default, "QDateTimeEdit_redirected_default")
	symaddr(&addr_QDateTimeEdit_sharedPainter_default, "QDateTimeEdit_sharedPainter_default")
	symaddr(&addr_QDateTimeEdit_inputMethodEvent_default, "QDateTimeEdit_inputMethodEvent_default")
	symaddr(&addr_QDateTimeEdit_eventFilter_default, "QDateTimeEdit_eventFilter_default")
	symaddr(&addr_QDateTimeEdit_childEvent_default, "QDateTimeEdit_childEvent_default")
	symaddr(&addr_QDateTimeEdit_customEvent_default, "QDateTimeEdit_customEvent_default")
	symaddr(&addr_QDateTimeEdit_connectNotify_default, "QDateTimeEdit_connectNotify_default")
	symaddr(&addr_QDateTimeEdit_disconnectNotify_default, "QDateTimeEdit_disconnectNotify_default")
	reg(&pg_QDateTimeEdit_connect_dateTimeChanged, "QDateTimeEdit_connect_dateTimeChanged")
	reg(&pg_QDateTimeEdit_connect_timeChanged, "QDateTimeEdit_connect_timeChanged")
	reg(&pg_QDateTimeEdit_connect_dateChanged, "QDateTimeEdit_connect_dateChanged")
	reg(&pg_QDateTimeEdit_connect_editingFinished, "QDateTimeEdit_connect_editingFinished")
	reg(&pg_QDateTimeEdit_connect_returnPressed, "QDateTimeEdit_connect_returnPressed")
	reg(&pg_QDateTimeEdit_connect_windowTitleChanged, "QDateTimeEdit_connect_windowTitleChanged")
	reg(&pg_QDateTimeEdit_connect_windowIconChanged, "QDateTimeEdit_connect_windowIconChanged")
	reg(&pg_QDateTimeEdit_connect_windowIconTextChanged, "QDateTimeEdit_connect_windowIconTextChanged")
	reg(&pg_QDateTimeEdit_connect_customContextMenuRequested, "QDateTimeEdit_connect_customContextMenuRequested")
	reg(&pg_QDateTimeEdit_connect_destroyed, "QDateTimeEdit_connect_destroyed")
	reg(&pg_QDateTimeEdit_connect_destroyed2, "QDateTimeEdit_connect_destroyed2")
	reg(&pg_QDateTimeEdit_connect_objectNameChanged, "QDateTimeEdit_connect_objectNameChanged")
	reg(&pg_QDateTimeEdit__disconnect, "QDateTimeEdit__disconnect")
	reg(&pg_QDateTimeEdit__set_override, "QDateTimeEdit__set_override")
	reg(&pg_QDateTimeEdit_delete, "QDateTimeEdit_delete")
}

type QDateTimeEdit struct {
	h uintptr
	*QAbstractSpinBox
}

func bindQDateTimeEdit(h uintptr) *QDateTimeEdit {
	if h == 0 {
		return nil
	}
	var __c [3]uintptr
	purego.SyscallN(addr_QDateTimeEdit_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QDateTimeEdit{h: h, QAbstractSpinBox: &QAbstractSpinBox{h: __c[0], QWidget: &QWidget{h: __c[1], QObject: &QObject{h: __c[2]}}}}
}

func newownedQDateTimeEdit(h uintptr) *QDateTimeEdit { return bindQDateTimeEdit(h) }

func (this *QDateTimeEdit) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QDateTimeEdit) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQDateTimeEdit(h unsafe.Pointer) *QDateTimeEdit { return bindQDateTimeEdit(uintptr(h)) }
func UnsafeNewQDateTimeEditOwned(h unsafe.Pointer) *QDateTimeEdit { return newownedQDateTimeEdit(uintptr(h)) }

func (this *QDateTimeEdit) UnsafeRelease() {
	this.h = 0
	if this.QAbstractSpinBox != nil {
		this.QAbstractSpinBox.UnsafeRelease()
	}
}

func NewQDateTimeEdit(parent *QWidget) *QDateTimeEdit {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QDateTimeEdit_new, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQDateTimeEdit(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QDateTimeEdit_virtIdx, QDateTimeEdit_enable)
	return __o
}

func NewQDateTimeEdit2() *QDateTimeEdit {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QDateTimeEdit_new2, uintptr(__goid))
	__o := newownedQDateTimeEdit(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QDateTimeEdit_virtIdx, QDateTimeEdit_enable)
	return __o
}

func NewQDateTimeEdit3(dt *QDateTime) *QDateTimeEdit {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QDateTimeEdit_new3, uintptr(__goid), uintptr(dt.UnsafePointer()))
	__o := newownedQDateTimeEdit(__r)
	runtime.KeepAlive(dt)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QDateTimeEdit_virtIdx, QDateTimeEdit_enable)
	return __o
}

func NewQDateTimeEdit4(d *QDate) *QDateTimeEdit {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QDateTimeEdit_new4, uintptr(__goid), uintptr(d.UnsafePointer()))
	__o := newownedQDateTimeEdit(__r)
	runtime.KeepAlive(d)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QDateTimeEdit_virtIdx, QDateTimeEdit_enable)
	return __o
}

func NewQDateTimeEdit5(t *QTime) *QDateTimeEdit {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QDateTimeEdit_new5, uintptr(__goid), uintptr(t.UnsafePointer()))
	__o := newownedQDateTimeEdit(__r)
	runtime.KeepAlive(t)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QDateTimeEdit_virtIdx, QDateTimeEdit_enable)
	return __o
}

func NewQDateTimeEdit6(dt *QDateTime, parent *QWidget) *QDateTimeEdit {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QDateTimeEdit_new6, uintptr(__goid), uintptr(dt.UnsafePointer()), uintptr(parent.UnsafePointer()))
	__o := newownedQDateTimeEdit(__r)
	runtime.KeepAlive(dt)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QDateTimeEdit_virtIdx, QDateTimeEdit_enable)
	return __o
}

func NewQDateTimeEdit7(d *QDate, parent *QWidget) *QDateTimeEdit {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QDateTimeEdit_new7, uintptr(__goid), uintptr(d.UnsafePointer()), uintptr(parent.UnsafePointer()))
	__o := newownedQDateTimeEdit(__r)
	runtime.KeepAlive(d)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QDateTimeEdit_virtIdx, QDateTimeEdit_enable)
	return __o
}

func NewQDateTimeEdit8(t *QTime, parent *QWidget) *QDateTimeEdit {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QDateTimeEdit_new8, uintptr(__goid), uintptr(t.UnsafePointer()), uintptr(parent.UnsafePointer()))
	__o := newownedQDateTimeEdit(__r)
	runtime.KeepAlive(t)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QDateTimeEdit_virtIdx, QDateTimeEdit_enable)
	return __o
}

func QDateTimeEdit_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDateTimeEdit_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QDateTimeEdit) DateTime() *QDateTime {
	__r, _, _ := purego.SyscallN(addr_QDateTimeEdit_dateTime, this.h)
	runtime.KeepAlive(this)
	return newownedQDateTime(__r)
}

func (this *QDateTimeEdit) Date() *QDate {
	__r, _, _ := purego.SyscallN(addr_QDateTimeEdit_date, this.h)
	runtime.KeepAlive(this)
	return newownedQDate(__r)
}

func (this *QDateTimeEdit) Time() *QTime {
	__r, _, _ := purego.SyscallN(addr_QDateTimeEdit_time, this.h)
	runtime.KeepAlive(this)
	return newownedQTime(__r)
}

func (this *QDateTimeEdit) Calendar() *QCalendar {
	__r, _, _ := purego.SyscallN(addr_QDateTimeEdit_calendar, this.h)
	runtime.KeepAlive(this)
	return newownedQCalendar(__r)
}

func (this *QDateTimeEdit) SetCalendar(calendar *QCalendar) {
	purego.SyscallN(addr_QDateTimeEdit_setCalendar, this.h, uintptr(calendar.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(calendar)
}

func (this *QDateTimeEdit) MinimumDateTime() *QDateTime {
	__r, _, _ := purego.SyscallN(addr_QDateTimeEdit_minimumDateTime, this.h)
	runtime.KeepAlive(this)
	return newownedQDateTime(__r)
}

func (this *QDateTimeEdit) ClearMinimumDateTime() {
	purego.SyscallN(addr_QDateTimeEdit_clearMinimumDateTime, this.h)
	runtime.KeepAlive(this)
}

func (this *QDateTimeEdit) SetMinimumDateTime(dt *QDateTime) {
	purego.SyscallN(addr_QDateTimeEdit_setMinimumDateTime, this.h, uintptr(dt.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(dt)
}

func (this *QDateTimeEdit) MaximumDateTime() *QDateTime {
	__r, _, _ := purego.SyscallN(addr_QDateTimeEdit_maximumDateTime, this.h)
	runtime.KeepAlive(this)
	return newownedQDateTime(__r)
}

func (this *QDateTimeEdit) ClearMaximumDateTime() {
	purego.SyscallN(addr_QDateTimeEdit_clearMaximumDateTime, this.h)
	runtime.KeepAlive(this)
}

func (this *QDateTimeEdit) SetMaximumDateTime(dt *QDateTime) {
	purego.SyscallN(addr_QDateTimeEdit_setMaximumDateTime, this.h, uintptr(dt.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(dt)
}

func (this *QDateTimeEdit) SetDateTimeRange(min *QDateTime, max *QDateTime) {
	purego.SyscallN(addr_QDateTimeEdit_setDateTimeRange, this.h, uintptr(min.UnsafePointer()), uintptr(max.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(min)
	runtime.KeepAlive(max)
}

func (this *QDateTimeEdit) MinimumDate() *QDate {
	__r, _, _ := purego.SyscallN(addr_QDateTimeEdit_minimumDate, this.h)
	runtime.KeepAlive(this)
	return newownedQDate(__r)
}

func (this *QDateTimeEdit) SetMinimumDate(min *QDate) {
	purego.SyscallN(addr_QDateTimeEdit_setMinimumDate, this.h, uintptr(min.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(min)
}

func (this *QDateTimeEdit) ClearMinimumDate() {
	purego.SyscallN(addr_QDateTimeEdit_clearMinimumDate, this.h)
	runtime.KeepAlive(this)
}

func (this *QDateTimeEdit) MaximumDate() *QDate {
	__r, _, _ := purego.SyscallN(addr_QDateTimeEdit_maximumDate, this.h)
	runtime.KeepAlive(this)
	return newownedQDate(__r)
}

func (this *QDateTimeEdit) SetMaximumDate(max *QDate) {
	purego.SyscallN(addr_QDateTimeEdit_setMaximumDate, this.h, uintptr(max.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(max)
}

func (this *QDateTimeEdit) ClearMaximumDate() {
	purego.SyscallN(addr_QDateTimeEdit_clearMaximumDate, this.h)
	runtime.KeepAlive(this)
}

func (this *QDateTimeEdit) SetDateRange(min *QDate, max *QDate) {
	purego.SyscallN(addr_QDateTimeEdit_setDateRange, this.h, uintptr(min.UnsafePointer()), uintptr(max.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(min)
	runtime.KeepAlive(max)
}

func (this *QDateTimeEdit) MinimumTime() *QTime {
	__r, _, _ := purego.SyscallN(addr_QDateTimeEdit_minimumTime, this.h)
	runtime.KeepAlive(this)
	return newownedQTime(__r)
}

func (this *QDateTimeEdit) SetMinimumTime(min *QTime) {
	purego.SyscallN(addr_QDateTimeEdit_setMinimumTime, this.h, uintptr(min.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(min)
}

func (this *QDateTimeEdit) ClearMinimumTime() {
	purego.SyscallN(addr_QDateTimeEdit_clearMinimumTime, this.h)
	runtime.KeepAlive(this)
}

func (this *QDateTimeEdit) MaximumTime() *QTime {
	__r, _, _ := purego.SyscallN(addr_QDateTimeEdit_maximumTime, this.h)
	runtime.KeepAlive(this)
	return newownedQTime(__r)
}

func (this *QDateTimeEdit) SetMaximumTime(max *QTime) {
	purego.SyscallN(addr_QDateTimeEdit_setMaximumTime, this.h, uintptr(max.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(max)
}

func (this *QDateTimeEdit) ClearMaximumTime() {
	purego.SyscallN(addr_QDateTimeEdit_clearMaximumTime, this.h)
	runtime.KeepAlive(this)
}

func (this *QDateTimeEdit) SetTimeRange(min *QTime, max *QTime) {
	purego.SyscallN(addr_QDateTimeEdit_setTimeRange, this.h, uintptr(min.UnsafePointer()), uintptr(max.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(min)
	runtime.KeepAlive(max)
}

func (this *QDateTimeEdit) DisplayedSections() QDateTimeEdit__Section {
	__r0, _, _ := purego.SyscallN(addr_QDateTimeEdit_displayedSections, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QDateTimeEdit__Section(__r)
}

func (this *QDateTimeEdit) CurrentSection() QDateTimeEdit__Section {
	__r0, _, _ := purego.SyscallN(addr_QDateTimeEdit_currentSection, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QDateTimeEdit__Section(__r)
}

func (this *QDateTimeEdit) SectionAt(index int) QDateTimeEdit__Section {
	__r0, _, _ := purego.SyscallN(addr_QDateTimeEdit_sectionAt, this.h, uintptr(index))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QDateTimeEdit__Section(__r)
}

func (this *QDateTimeEdit) SetCurrentSection(section QDateTimeEdit__Section) {
	purego.SyscallN(addr_QDateTimeEdit_setCurrentSection, this.h, uintptr(section))
	runtime.KeepAlive(this)
}

func (this *QDateTimeEdit) CurrentSectionIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QDateTimeEdit_currentSectionIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QDateTimeEdit) SetCurrentSectionIndex(index int) {
	purego.SyscallN(addr_QDateTimeEdit_setCurrentSectionIndex, this.h, uintptr(index))
	runtime.KeepAlive(this)
}

func (this *QDateTimeEdit) CalendarWidget() *QCalendarWidget {
	__r, _, _ := purego.SyscallN(addr_QDateTimeEdit_calendarWidget, this.h)
	runtime.KeepAlive(this)
	return bindQCalendarWidget(__r)
}

func (this *QDateTimeEdit) SetCalendarWidget(calendarWidget *QCalendarWidget) {
	purego.SyscallN(addr_QDateTimeEdit_setCalendarWidget, this.h, uintptr(calendarWidget.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(calendarWidget)
}

func (this *QDateTimeEdit) SectionCount() int {
	__r0, _, _ := purego.SyscallN(addr_QDateTimeEdit_sectionCount, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QDateTimeEdit) SetSelectedSection(section QDateTimeEdit__Section) {
	purego.SyscallN(addr_QDateTimeEdit_setSelectedSection, this.h, uintptr(section))
	runtime.KeepAlive(this)
}

func (this *QDateTimeEdit) SectionText(section QDateTimeEdit__Section) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDateTimeEdit_sectionText, this.h, uintptr(section), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QDateTimeEdit) DisplayFormat() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDateTimeEdit_displayFormat, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QDateTimeEdit) SetDisplayFormat(format string) {
	var format_p uintptr
	if len(format) > 0 {
		format_p = uintptr(unsafe.Pointer(unsafe.StringData(format)))
	}
	purego.SyscallN(addr_QDateTimeEdit_setDisplayFormat, this.h, format_p, uintptr(len(format)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(format)
}

func (this *QDateTimeEdit) CalendarPopup() bool {
	__r0, _, _ := purego.SyscallN(addr_QDateTimeEdit_calendarPopup, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QDateTimeEdit) SetCalendarPopup(enable bool) {
	purego.SyscallN(addr_QDateTimeEdit_setCalendarPopup, this.h, uintptr(b2i(enable)))
	runtime.KeepAlive(this)
}

func (this *QDateTimeEdit) TimeSpec() TimeSpec {
	__r0, _, _ := purego.SyscallN(addr_QDateTimeEdit_timeSpec, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return TimeSpec(__r)
}

func (this *QDateTimeEdit) SetTimeSpec(spec TimeSpec) {
	purego.SyscallN(addr_QDateTimeEdit_setTimeSpec, this.h, uintptr(spec))
	runtime.KeepAlive(this)
}

func (this *QDateTimeEdit) TimeZone() *QTimeZone {
	__r, _, _ := purego.SyscallN(addr_QDateTimeEdit_timeZone, this.h)
	runtime.KeepAlive(this)
	return newownedQTimeZone(__r)
}

func (this *QDateTimeEdit) SetTimeZone(zone *QTimeZone) {
	purego.SyscallN(addr_QDateTimeEdit_setTimeZone, this.h, uintptr(zone.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(zone)
}

func (this *QDateTimeEdit) DateTimeChanged(dateTime *QDateTime) {
	purego.SyscallN(addr_QDateTimeEdit_dateTimeChanged, this.h, uintptr(dateTime.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(dateTime)
}

func (this *QDateTimeEdit) TimeChanged(time *QTime) {
	purego.SyscallN(addr_QDateTimeEdit_timeChanged, this.h, uintptr(time.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(time)
}

func (this *QDateTimeEdit) DateChanged(date *QDate) {
	purego.SyscallN(addr_QDateTimeEdit_dateChanged, this.h, uintptr(date.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(date)
}

func (this *QDateTimeEdit) SetDateTime(dateTime *QDateTime) {
	purego.SyscallN(addr_QDateTimeEdit_setDateTime, this.h, uintptr(dateTime.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(dateTime)
}

func (this *QDateTimeEdit) SetDate(date *QDate) {
	purego.SyscallN(addr_QDateTimeEdit_setDate, this.h, uintptr(date.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(date)
}

func (this *QDateTimeEdit) SetTime(time *QTime) {
	purego.SyscallN(addr_QDateTimeEdit_setTime, this.h, uintptr(time.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(time)
}

func (this *QDateTimeEdit) LineEdit() *QLineEdit {
	__r, _, _ := purego.SyscallN(addr_QDateTimeEdit_lineEdit, this.h)
	runtime.KeepAlive(this)
	return bindQLineEdit(__r)
}

func (this *QDateTimeEdit) SetLineEdit(edit *QLineEdit) {
	purego.SyscallN(addr_QDateTimeEdit_setLineEdit, this.h, uintptr(edit.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(edit)
}

func QDateTimeEdit_SetTabOrder(arg1 *QWidget, arg2 *QWidget) {
	purego.SyscallN(addr_QDateTimeEdit_setTabOrder, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()))
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
}

func QDateTimeEdit_MouseGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QDateTimeEdit_mouseGrabber)
	return bindQWidget(__r)
}

func QDateTimeEdit_KeyboardGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QDateTimeEdit_keyboardGrabber)
	return bindQWidget(__r)
}

func QDateTimeEdit_Find(arg1 uint64) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QDateTimeEdit_find, uintptr(arg1))
	return bindQWidget(__r)
}

func QDateTimeEdit_CreateWindowContainer(window *QWindow) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QDateTimeEdit_createWindowContainer, uintptr(window.UnsafePointer()))
	runtime.KeepAlive(window)
	return bindQWidget(__r)
}

func (this *QDateTimeEdit) UpdateMicroFocus() {
	purego.SyscallN(addr_QDateTimeEdit_updateMicroFocus, this.h)
	runtime.KeepAlive(this)
}

func (this *QDateTimeEdit) Create() {
	purego.SyscallN(addr_QDateTimeEdit_create, this.h)
	runtime.KeepAlive(this)
}

func (this *QDateTimeEdit) Destroy() {
	purego.SyscallN(addr_QDateTimeEdit_destroy, this.h)
	runtime.KeepAlive(this)
}

func (this *QDateTimeEdit) FocusNextChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QDateTimeEdit_focusNextChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QDateTimeEdit) FocusPreviousChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QDateTimeEdit_focusPreviousChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QDateTimeEdit_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QDateTimeEdit_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QDateTimeEdit_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QDateTimeEdit_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QDateTimeEdit) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QDateTimeEdit_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QDateTimeEdit) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QDateTimeEdit_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QDateTimeEdit) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QDateTimeEdit_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QDateTimeEdit) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QDateTimeEdit_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QDateTimeEdit_DevicePixelRatioFScale() float64 {
	__r := pg_QDateTimeEdit_devicePixelRatioFScale()
	return __r
}

func QDateTimeEdit_EncodeMetricF(metric QPaintDevice__PaintDeviceMetric, value float64) int {
	__r := pg_QDateTimeEdit_encodeMetricF(int32(metric), value)
	return int(__r)
}

func (this *QDateTimeEdit) GetDecodedMetricF(metricA QPaintDevice__PaintDeviceMetric, metricB QPaintDevice__PaintDeviceMetric) float64 {
	__r := pg_QDateTimeEdit_getDecodedMetricF(this.h, int32(metricA), int32(metricB))
	runtime.KeepAlive(this)
	return __r
}

func QDateTimeEdit_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDateTimeEdit_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QDateTimeEdit_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDateTimeEdit_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QDateTimeEdit_CreateWindowContainer2(window *QWindow, parent *QWidget) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QDateTimeEdit_createWindowContainer2, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func QDateTimeEdit_CreateWindowContainer3(window *QWindow, parent *QWidget, flags WindowType) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QDateTimeEdit_createWindowContainer3, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()), uintptr(flags))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func (this *QDateTimeEdit) UpdateMicroFocusWithQuery(query InputMethodQuery) {
	purego.SyscallN(addr_QDateTimeEdit_updateMicroFocus2, this.h, uintptr(query))
	runtime.KeepAlive(this)
}

func (this *QDateTimeEdit) CreateWith(arg1 uint64) {
	purego.SyscallN(addr_QDateTimeEdit_create2, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QDateTimeEdit) Create2(arg1 uint64, initializeWindow bool) {
	purego.SyscallN(addr_QDateTimeEdit_create3, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)))
	runtime.KeepAlive(this)
}

func (this *QDateTimeEdit) Create3(arg1 uint64, initializeWindow bool, destroyOldWindow bool) {
	purego.SyscallN(addr_QDateTimeEdit_create4, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)), uintptr(b2i(destroyOldWindow)))
	runtime.KeepAlive(this)
}

func (this *QDateTimeEdit) DestroyWithDestroyWindow(destroyWindow bool) {
	purego.SyscallN(addr_QDateTimeEdit_destroy2, this.h, uintptr(b2i(destroyWindow)))
	runtime.KeepAlive(this)
}

func (this *QDateTimeEdit) Destroy2(destroyWindow bool, destroySubWindows bool) {
	purego.SyscallN(addr_QDateTimeEdit_destroy3, this.h, uintptr(b2i(destroyWindow)), uintptr(b2i(destroySubWindows)))
	runtime.KeepAlive(this)
}

func QDateTimeEdit_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QDateTimeEdit_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QDateTimeEdit_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QDateTimeEdit_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QDateTimeEdit) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QDateTimeEdit_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QDateTimeEdit) SizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QDateTimeEdit_sizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QDateTimeEdit) ClearDefault() {
	purego.SyscallN(addr_QDateTimeEdit_clear_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QDateTimeEdit) StepByDefault(steps int) {
	purego.SyscallN(addr_QDateTimeEdit_stepBy_default, this.h, uintptr(steps))
	runtime.KeepAlive(this)
}

func (this *QDateTimeEdit) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QDateTimeEdit_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QDateTimeEdit) KeyPressEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QDateTimeEdit_keyPressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDateTimeEdit) WheelEventDefault(event *QWheelEvent) {
	purego.SyscallN(addr_QDateTimeEdit_wheelEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDateTimeEdit) FocusInEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QDateTimeEdit_focusInEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDateTimeEdit) FocusNextPrevChildDefault(next bool) bool {
	__r0, _, _ := purego.SyscallN(addr_QDateTimeEdit_focusNextPrevChild_default, this.h, uintptr(b2i(next)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QDateTimeEdit) ValidateDefault(input *QString, pos int) (QValidator__State, int) {
	pos_io := int32(pos)
	__r0, _, _ := purego.SyscallN(addr_QDateTimeEdit_validate_default, this.h, uintptr(input.UnsafePointer()), uintptr(unsafe.Pointer(&pos_io)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(input)
	return QValidator__State(__r), int(pos_io)
}

func (this *QDateTimeEdit) FixupDefault(input *QString) {
	purego.SyscallN(addr_QDateTimeEdit_fixup_default, this.h, uintptr(input.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(input)
}

func (this *QDateTimeEdit) DateTimeFromTextDefault(text string) *QDateTime {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	__r, _, _ := purego.SyscallN(addr_QDateTimeEdit_dateTimeFromText_default, this.h, text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
	return newownedQDateTime(__r)
}

func (this *QDateTimeEdit) TextFromDateTimeDefault(dt *QDateTime) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QDateTimeEdit_textFromDateTime_default, this.h, uintptr(dt.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(dt)
	return pgStr(__r, __n)
}

func (this *QDateTimeEdit) StepEnabledDefault() QAbstractSpinBox__StepEnabledFlag {
	__r0, _, _ := purego.SyscallN(addr_QDateTimeEdit_stepEnabled_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QAbstractSpinBox__StepEnabledFlag(__r)
}

func (this *QDateTimeEdit) MousePressEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QDateTimeEdit_mousePressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDateTimeEdit) PaintEventDefault(event *QPaintEvent) {
	purego.SyscallN(addr_QDateTimeEdit_paintEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDateTimeEdit) InitStyleOptionDefault(option *QStyleOptionSpinBox) {
	purego.SyscallN(addr_QDateTimeEdit_initStyleOption_default, this.h, uintptr(option.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(option)
}

func (this *QDateTimeEdit) MinimumSizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QDateTimeEdit_minimumSizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QDateTimeEdit) InputMethodQueryDefault(arg1 InputMethodQuery) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QDateTimeEdit_inputMethodQuery_default, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QDateTimeEdit) ResizeEventDefault(event *QResizeEvent) {
	purego.SyscallN(addr_QDateTimeEdit_resizeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDateTimeEdit) KeyReleaseEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QDateTimeEdit_keyReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDateTimeEdit) FocusOutEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QDateTimeEdit_focusOutEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDateTimeEdit) ContextMenuEventDefault(event *QContextMenuEvent) {
	purego.SyscallN(addr_QDateTimeEdit_contextMenuEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDateTimeEdit) ChangeEventDefault(event *QEvent) {
	purego.SyscallN(addr_QDateTimeEdit_changeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDateTimeEdit) CloseEventDefault(event *QCloseEvent) {
	purego.SyscallN(addr_QDateTimeEdit_closeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDateTimeEdit) HideEventDefault(event *QHideEvent) {
	purego.SyscallN(addr_QDateTimeEdit_hideEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDateTimeEdit) MouseReleaseEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QDateTimeEdit_mouseReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDateTimeEdit) MouseMoveEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QDateTimeEdit_mouseMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDateTimeEdit) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QDateTimeEdit_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDateTimeEdit) ShowEventDefault(event *QShowEvent) {
	purego.SyscallN(addr_QDateTimeEdit_showEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDateTimeEdit) DevTypeDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QDateTimeEdit_devType_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QDateTimeEdit) SetVisibleDefault(visible bool) {
	purego.SyscallN(addr_QDateTimeEdit_setVisible_default, this.h, uintptr(b2i(visible)))
	runtime.KeepAlive(this)
}

func (this *QDateTimeEdit) HeightForWidthDefault(arg1 int) int {
	__r0, _, _ := purego.SyscallN(addr_QDateTimeEdit_heightForWidth_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QDateTimeEdit) HasHeightForWidthDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QDateTimeEdit_hasHeightForWidth_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QDateTimeEdit) PaintEngineDefault() *QPaintEngine {
	__r, _, _ := purego.SyscallN(addr_QDateTimeEdit_paintEngine_default, this.h)
	runtime.KeepAlive(this)
	return bindQPaintEngine(__r)
}

func (this *QDateTimeEdit) MouseDoubleClickEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QDateTimeEdit_mouseDoubleClickEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDateTimeEdit) EnterEventDefault(event *QEnterEvent) {
	purego.SyscallN(addr_QDateTimeEdit_enterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDateTimeEdit) LeaveEventDefault(event *QEvent) {
	purego.SyscallN(addr_QDateTimeEdit_leaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDateTimeEdit) MoveEventDefault(event *QMoveEvent) {
	purego.SyscallN(addr_QDateTimeEdit_moveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDateTimeEdit) TabletEventDefault(event *QTabletEvent) {
	purego.SyscallN(addr_QDateTimeEdit_tabletEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDateTimeEdit) ActionEventDefault(event *QActionEvent) {
	purego.SyscallN(addr_QDateTimeEdit_actionEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDateTimeEdit) DragEnterEventDefault(event *QDragEnterEvent) {
	purego.SyscallN(addr_QDateTimeEdit_dragEnterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDateTimeEdit) DragMoveEventDefault(event *QDragMoveEvent) {
	purego.SyscallN(addr_QDateTimeEdit_dragMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDateTimeEdit) DragLeaveEventDefault(event *QDragLeaveEvent) {
	purego.SyscallN(addr_QDateTimeEdit_dragLeaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDateTimeEdit) DropEventDefault(event *QDropEvent) {
	purego.SyscallN(addr_QDateTimeEdit_dropEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDateTimeEdit) MetricDefault(arg1 QPaintDevice__PaintDeviceMetric) int {
	__r0, _, _ := purego.SyscallN(addr_QDateTimeEdit_metric_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QDateTimeEdit) InitPainterDefault(painter *QPainter) {
	purego.SyscallN(addr_QDateTimeEdit_initPainter_default, this.h, uintptr(painter.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
}

func (this *QDateTimeEdit) RedirectedDefault(offset *QPoint) *QPaintDevice {
	__r, _, _ := purego.SyscallN(addr_QDateTimeEdit_redirected_default, this.h, uintptr(offset.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(offset)
	return bindQPaintDevice(__r)
}

func (this *QDateTimeEdit) SharedPainterDefault() *QPainter {
	__r, _, _ := purego.SyscallN(addr_QDateTimeEdit_sharedPainter_default, this.h)
	runtime.KeepAlive(this)
	return bindQPainter(__r)
}

func (this *QDateTimeEdit) InputMethodEventDefault(arg1 *QInputMethodEvent) {
	purego.SyscallN(addr_QDateTimeEdit_inputMethodEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QDateTimeEdit) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QDateTimeEdit_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QDateTimeEdit) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QDateTimeEdit_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDateTimeEdit) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QDateTimeEdit_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QDateTimeEdit) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QDateTimeEdit_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QDateTimeEdit) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QDateTimeEdit_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QDateTimeEdit) OnDateTimeChanged(fn func(*QDateTime)) int64 {
	id := registerCallback(func(dateTime uintptr) { fn(UnsafeNewQDateTime(unsafe.Pointer(dateTime))) })
	pg_QDateTimeEdit_connect_dateTimeChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDateTimeEdit) OnTimeChanged(fn func(*QTime)) int64 {
	id := registerCallback(func(time uintptr) { fn(UnsafeNewQTime(unsafe.Pointer(time))) })
	pg_QDateTimeEdit_connect_timeChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDateTimeEdit) OnDateChanged(fn func(*QDate)) int64 {
	id := registerCallback(func(date uintptr) { fn(UnsafeNewQDate(unsafe.Pointer(date))) })
	pg_QDateTimeEdit_connect_dateChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDateTimeEdit) OnEditingFinished(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QDateTimeEdit_connect_editingFinished(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDateTimeEdit) OnReturnPressed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QDateTimeEdit_connect_returnPressed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDateTimeEdit) OnWindowTitleChanged(fn func(string)) int64 {
	id := registerCallback(func(title uintptr, title_len int64) { fn(pgStrNoFree(title, title_len)) })
	pg_QDateTimeEdit_connect_windowTitleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDateTimeEdit) OnWindowIconChanged(fn func(*QIcon)) int64 {
	id := registerCallback(func(icon uintptr) { fn(UnsafeNewQIcon(unsafe.Pointer(icon))) })
	pg_QDateTimeEdit_connect_windowIconChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDateTimeEdit) OnWindowIconTextChanged(fn func(string)) int64 {
	id := registerCallback(func(iconText uintptr, iconText_len int64) { fn(pgStrNoFree(iconText, iconText_len)) })
	pg_QDateTimeEdit_connect_windowIconTextChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDateTimeEdit) OnCustomContextMenuRequested(fn func(*QPoint)) int64 {
	id := registerCallback(func(pos uintptr) { fn(UnsafeNewQPoint(unsafe.Pointer(pos))) })
	pg_QDateTimeEdit_connect_customContextMenuRequested(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDateTimeEdit) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QDateTimeEdit_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDateTimeEdit) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QDateTimeEdit_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDateTimeEdit) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QDateTimeEdit_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QDateTimeEdit) Disconnect(handle int64) {
	pg_QDateTimeEdit__disconnect(handle)
	unregisterCallback(handle)
}

func QDateTimeEdit_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case 729690261457684101:
		return 1
	case -778567940572639070:
		return 2
	case 7085718643589954634:
		return 3
	case 1237236456312824991:
		return 4
	case 1023791977679793185:
		return 5
	case 1919988340934511316:
		return 6
	case 975043767033552552:
		return 7
	case 5898328517505784191:
		return 8
	case 6784179835071232555:
		return 9
	case -4855206404746966715:
		return 10
	case -50441425836614729:
		return 11
	case -3778148144965928937:
		return 12
	case 5257289574516412244:
		return 13
	case -332090898078543319:
		return 14
	case 5290442163167458465:
		return 15
	case -488693769771308531:
		return 16
	case 7049264599837532491:
		return 17
	case 7345853097340709638:
		return 18
	case -4302917285655014661:
		return 19
	case 772598475950050299:
		return 20
	case -385596010576048993:
		return 21
	case -6698043222227006015:
		return 22
	case 4959376063826165943:
		return 23
	case -1989937050019464783:
		return 24
	case -199718707876719129:
		return 25
	case -4415920049297795213:
		return 26
	case -4255494004099806581:
		return 27
	case -5409927581812374498:
		return 28
	case -6825369174751916:
		return 29
	case -8659679893868524044:
		return 30
	case 2110182719375245575:
		return 31
	case 4311534060495737339:
		return 32
	case -8266762397995423915:
		return 33
	case 6398007404756346381:
		return 34
	case -4029073905697023887:
		return 35
	case 6316270346668237895:
		return 36
	case -551070009623711640:
		return 37
	case -9161949150489525396:
		return 38
	case -212068866860701141:
		return 39
	case 1697444140088555581:
		return 40
	case 206122715919917589:
		return 41
	case -1781545137352745978:
		return 42
	case -3139904158373725006:
		return 43
	case 9107058682754591716:
		return 44
	case -2039166586246443706:
		return 45
	case 427629649267311777:
		return 46
	case 6477500296980050210:
		return 47
	case 4127240633310655048:
		return 48
	case 5792771328741774991:
		return 49
	case 5438819661942843222:
		return 50
	case 5217623514630858297:
		return 51
	case -1185315367526857833:
		return 52
	case 3769868229123868474:
		return 53
	case 7379029982170503008:
		return 54
	case -568369772982215962:
		return 55
	}
	return -1
}

func QDateTimeEdit_enable(h uintptr, idx int32) { pg_QDateTimeEdit__set_override(h, idx) }

func (this *QDateTimeEdit) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QDateTimeEdit) OnSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 729690261457684101, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sizeHint"
}

func (this *QDateTimeEdit) OnClear(fn func()) {
	registerVirt(this.UnsafeGoID(), -778567940572639070, func() { fn() }) // slot "clear"
}

func (this *QDateTimeEdit) OnStepBy(fn func(int)) {
	registerVirt(this.UnsafeGoID(), 7085718643589954634, func(steps int32) { fn(int(steps)) }) // slot "stepBy"
}

func (this *QDateTimeEdit) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QDateTimeEdit) OnKeyPressEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 1023791977679793185, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyPressEvent"
}

func (this *QDateTimeEdit) OnWheelEvent(fn func(*QWheelEvent)) {
	registerVirt(this.UnsafeGoID(), 1919988340934511316, func(event uintptr) { fn(UnsafeNewQWheelEvent(unsafe.Pointer(event))) }) // slot "wheelEvent"
}

func (this *QDateTimeEdit) OnFocusInEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), 975043767033552552, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusInEvent"
}

func (this *QDateTimeEdit) OnFocusNextPrevChild(fn func(bool) bool) {
	registerVirt(this.UnsafeGoID(), 5898328517505784191, func(next int32) int32 { return b2i(fn(next != 0)) }) // slot "focusNextPrevChild"
}

func (this *QDateTimeEdit) OnValidate(fn func(*QString, *int32) QValidator__State) {
	registerVirt(this.UnsafeGoID(), 6784179835071232555, func(input uintptr, pos uintptr) int32 { return int32(fn(UnsafeNewQString(unsafe.Pointer(input)), (*int32)(unsafe.Pointer(pos)))) }) // slot "validate"
}

func (this *QDateTimeEdit) OnFixup(fn func(*QString)) {
	registerVirt(this.UnsafeGoID(), -4855206404746966715, func(input uintptr) { fn(UnsafeNewQString(unsafe.Pointer(input))) }) // slot "fixup"
}

func (this *QDateTimeEdit) OnDateTimeFromText(fn func(string) *QDateTime) {
	registerVirt(this.UnsafeGoID(), -50441425836614729, func(text uintptr, text_len int64) uintptr { return uintptr(fn(pgStrNoFree(text, text_len)).UnsafePointer()) }) // slot "dateTimeFromText"
}

func (this *QDateTimeEdit) OnTextFromDateTime(fn func(*QDateTime) string) {
	registerVirt(this.UnsafeGoID(), -3778148144965928937, func(dt uintptr) (uintptr, int64) {
		__s := fn(UnsafeNewQDateTime(unsafe.Pointer(dt)))
		return pgCString(__s), int64(len(__s))
	}) // slot "textFromDateTime"
}

func (this *QDateTimeEdit) OnStepEnabled(fn func() QAbstractSpinBox__StepEnabledFlag) {
	registerVirt(this.UnsafeGoID(), 5257289574516412244, func() int32 { return int32(fn()) }) // slot "stepEnabled"
}

func (this *QDateTimeEdit) OnMousePressEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -332090898078543319, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mousePressEvent"
}

func (this *QDateTimeEdit) OnPaintEvent(fn func(*QPaintEvent)) {
	registerVirt(this.UnsafeGoID(), 5290442163167458465, func(event uintptr) { fn(UnsafeNewQPaintEvent(unsafe.Pointer(event))) }) // slot "paintEvent"
}

func (this *QDateTimeEdit) OnInitStyleOption(fn func(*QStyleOptionSpinBox)) {
	registerVirt(this.UnsafeGoID(), -488693769771308531, func(option uintptr) { fn(UnsafeNewQStyleOptionSpinBox(unsafe.Pointer(option))) }) // slot "initStyleOption"
}

func (this *QDateTimeEdit) OnMinimumSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 7049264599837532491, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "minimumSizeHint"
}

func (this *QDateTimeEdit) OnInputMethodQuery(fn func(InputMethodQuery) *QVariant) {
	registerVirt(this.UnsafeGoID(), 7345853097340709638, func(arg1 int32) uintptr { return uintptr(fn(InputMethodQuery(arg1)).UnsafePointer()) }) // slot "inputMethodQuery"
}

func (this *QDateTimeEdit) OnResizeEvent(fn func(*QResizeEvent)) {
	registerVirt(this.UnsafeGoID(), -4302917285655014661, func(event uintptr) { fn(UnsafeNewQResizeEvent(unsafe.Pointer(event))) }) // slot "resizeEvent"
}

func (this *QDateTimeEdit) OnKeyReleaseEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 772598475950050299, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyReleaseEvent"
}

func (this *QDateTimeEdit) OnFocusOutEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), -385596010576048993, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusOutEvent"
}

func (this *QDateTimeEdit) OnContextMenuEvent(fn func(*QContextMenuEvent)) {
	registerVirt(this.UnsafeGoID(), -6698043222227006015, func(event uintptr) { fn(UnsafeNewQContextMenuEvent(unsafe.Pointer(event))) }) // slot "contextMenuEvent"
}

func (this *QDateTimeEdit) OnChangeEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 4959376063826165943, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "changeEvent"
}

func (this *QDateTimeEdit) OnCloseEvent(fn func(*QCloseEvent)) {
	registerVirt(this.UnsafeGoID(), -1989937050019464783, func(event uintptr) { fn(UnsafeNewQCloseEvent(unsafe.Pointer(event))) }) // slot "closeEvent"
}

func (this *QDateTimeEdit) OnHideEvent(fn func(*QHideEvent)) {
	registerVirt(this.UnsafeGoID(), -199718707876719129, func(event uintptr) { fn(UnsafeNewQHideEvent(unsafe.Pointer(event))) }) // slot "hideEvent"
}

func (this *QDateTimeEdit) OnMouseReleaseEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4415920049297795213, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseReleaseEvent"
}

func (this *QDateTimeEdit) OnMouseMoveEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4255494004099806581, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseMoveEvent"
}

func (this *QDateTimeEdit) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QDateTimeEdit) OnShowEvent(fn func(*QShowEvent)) {
	registerVirt(this.UnsafeGoID(), -6825369174751916, func(event uintptr) { fn(UnsafeNewQShowEvent(unsafe.Pointer(event))) }) // slot "showEvent"
}

func (this *QDateTimeEdit) OnDevType(fn func() int) {
	registerVirt(this.UnsafeGoID(), -8659679893868524044, func() int32 { return int32(fn()) }) // slot "devType"
}

func (this *QDateTimeEdit) OnSetVisible(fn func(bool)) {
	registerVirt(this.UnsafeGoID(), 2110182719375245575, func(visible int32) { fn(visible != 0) }) // slot "setVisible"
}

func (this *QDateTimeEdit) OnHeightForWidth(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), 4311534060495737339, func(arg1 int32) int32 { return int32(fn(int(arg1))) }) // slot "heightForWidth"
}

func (this *QDateTimeEdit) OnHasHeightForWidth(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -8266762397995423915, func() int32 { return b2i(fn()) }) // slot "hasHeightForWidth"
}

func (this *QDateTimeEdit) OnPaintEngine(fn func() *QPaintEngine) {
	registerVirt(this.UnsafeGoID(), 6398007404756346381, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "paintEngine"
}

func (this *QDateTimeEdit) OnMouseDoubleClickEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4029073905697023887, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseDoubleClickEvent"
}

func (this *QDateTimeEdit) OnEnterEvent(fn func(*QEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 6316270346668237895, func(event uintptr) { fn(UnsafeNewQEnterEvent(unsafe.Pointer(event))) }) // slot "enterEvent"
}

func (this *QDateTimeEdit) OnLeaveEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), -551070009623711640, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "leaveEvent"
}

func (this *QDateTimeEdit) OnMoveEvent(fn func(*QMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -9161949150489525396, func(event uintptr) { fn(UnsafeNewQMoveEvent(unsafe.Pointer(event))) }) // slot "moveEvent"
}

func (this *QDateTimeEdit) OnTabletEvent(fn func(*QTabletEvent)) {
	registerVirt(this.UnsafeGoID(), -212068866860701141, func(event uintptr) { fn(UnsafeNewQTabletEvent(unsafe.Pointer(event))) }) // slot "tabletEvent"
}

func (this *QDateTimeEdit) OnActionEvent(fn func(*QActionEvent)) {
	registerVirt(this.UnsafeGoID(), 1697444140088555581, func(event uintptr) { fn(UnsafeNewQActionEvent(unsafe.Pointer(event))) }) // slot "actionEvent"
}

func (this *QDateTimeEdit) OnDragEnterEvent(fn func(*QDragEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 206122715919917589, func(event uintptr) { fn(UnsafeNewQDragEnterEvent(unsafe.Pointer(event))) }) // slot "dragEnterEvent"
}

func (this *QDateTimeEdit) OnDragMoveEvent(fn func(*QDragMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -1781545137352745978, func(event uintptr) { fn(UnsafeNewQDragMoveEvent(unsafe.Pointer(event))) }) // slot "dragMoveEvent"
}

func (this *QDateTimeEdit) OnDragLeaveEvent(fn func(*QDragLeaveEvent)) {
	registerVirt(this.UnsafeGoID(), -3139904158373725006, func(event uintptr) { fn(UnsafeNewQDragLeaveEvent(unsafe.Pointer(event))) }) // slot "dragLeaveEvent"
}

func (this *QDateTimeEdit) OnDropEvent(fn func(*QDropEvent)) {
	registerVirt(this.UnsafeGoID(), 9107058682754591716, func(event uintptr) { fn(UnsafeNewQDropEvent(unsafe.Pointer(event))) }) // slot "dropEvent"
}

func (this *QDateTimeEdit) OnNativeEvent(fn func(*QByteArray, unsafe.Pointer, *int64) bool) {
	registerVirt(this.UnsafeGoID(), -2039166586246443706, func(eventType uintptr, message uintptr, result uintptr) int32 { return b2i(fn(UnsafeNewQByteArray(unsafe.Pointer(eventType)), unsafe.Pointer(message), (*int64)(unsafe.Pointer(result)))) }) // slot "nativeEvent"
}

func (this *QDateTimeEdit) OnMetric(fn func(QPaintDevice__PaintDeviceMetric) int) {
	registerVirt(this.UnsafeGoID(), 427629649267311777, func(arg1 int32) int32 { return int32(fn(QPaintDevice__PaintDeviceMetric(arg1))) }) // slot "metric"
}

func (this *QDateTimeEdit) OnInitPainter(fn func(*QPainter)) {
	registerVirt(this.UnsafeGoID(), 6477500296980050210, func(painter uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter))) }) // slot "initPainter"
}

func (this *QDateTimeEdit) OnRedirected(fn func(*QPoint) *QPaintDevice) {
	registerVirt(this.UnsafeGoID(), 4127240633310655048, func(offset uintptr) uintptr { return uintptr(fn(UnsafeNewQPoint(unsafe.Pointer(offset))).UnsafePointer()) }) // slot "redirected"
}

func (this *QDateTimeEdit) OnSharedPainter(fn func() *QPainter) {
	registerVirt(this.UnsafeGoID(), 5792771328741774991, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sharedPainter"
}

func (this *QDateTimeEdit) OnInputMethodEvent(fn func(*QInputMethodEvent)) {
	registerVirt(this.UnsafeGoID(), 5438819661942843222, func(arg1 uintptr) { fn(UnsafeNewQInputMethodEvent(unsafe.Pointer(arg1))) }) // slot "inputMethodEvent"
}

func (this *QDateTimeEdit) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QDateTimeEdit) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QDateTimeEdit) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QDateTimeEdit) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QDateTimeEdit) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QDateTimeEdit) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QDateTimeEdit_delete(this.h)
	}
	this.UnsafeRelease()
}
