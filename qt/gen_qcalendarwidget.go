package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QCalendarWidget_chain uintptr
	addr_QCalendarWidget_new uintptr
	addr_QCalendarWidget_new2 uintptr
	addr_QCalendarWidget_tr uintptr
	addr_QCalendarWidget_selectedDate uintptr
	addr_QCalendarWidget_yearShown uintptr
	addr_QCalendarWidget_monthShown uintptr
	addr_QCalendarWidget_minimumDate uintptr
	addr_QCalendarWidget_setMinimumDate uintptr
	addr_QCalendarWidget_clearMinimumDate uintptr
	addr_QCalendarWidget_maximumDate uintptr
	addr_QCalendarWidget_setMaximumDate uintptr
	addr_QCalendarWidget_clearMaximumDate uintptr
	addr_QCalendarWidget_firstDayOfWeek uintptr
	addr_QCalendarWidget_setFirstDayOfWeek uintptr
	addr_QCalendarWidget_isNavigationBarVisible uintptr
	addr_QCalendarWidget_isGridVisible uintptr
	addr_QCalendarWidget_calendar uintptr
	addr_QCalendarWidget_setCalendar uintptr
	addr_QCalendarWidget_selectionMode uintptr
	addr_QCalendarWidget_setSelectionMode uintptr
	addr_QCalendarWidget_horizontalHeaderFormat uintptr
	addr_QCalendarWidget_setHorizontalHeaderFormat uintptr
	addr_QCalendarWidget_verticalHeaderFormat uintptr
	addr_QCalendarWidget_setVerticalHeaderFormat uintptr
	addr_QCalendarWidget_headerTextFormat uintptr
	addr_QCalendarWidget_setHeaderTextFormat uintptr
	addr_QCalendarWidget_weekdayTextFormat uintptr
	addr_QCalendarWidget_setWeekdayTextFormat uintptr
	addr_QCalendarWidget_dateTextFormat uintptr
	addr_QCalendarWidget_setDateTextFormat uintptr
	addr_QCalendarWidget_isDateEditEnabled uintptr
	addr_QCalendarWidget_setDateEditEnabled uintptr
	addr_QCalendarWidget_dateEditAcceptDelay uintptr
	addr_QCalendarWidget_setDateEditAcceptDelay uintptr
	addr_QCalendarWidget_updateCell uintptr
	addr_QCalendarWidget_updateCells uintptr
	addr_QCalendarWidget_setSelectedDate uintptr
	addr_QCalendarWidget_setDateRange uintptr
	addr_QCalendarWidget_setCurrentPage uintptr
	addr_QCalendarWidget_setGridVisible uintptr
	addr_QCalendarWidget_setNavigationBarVisible uintptr
	addr_QCalendarWidget_showNextMonth uintptr
	addr_QCalendarWidget_showPreviousMonth uintptr
	addr_QCalendarWidget_showNextYear uintptr
	addr_QCalendarWidget_showPreviousYear uintptr
	addr_QCalendarWidget_showSelectedDate uintptr
	addr_QCalendarWidget_showToday uintptr
	addr_QCalendarWidget_selectionChanged uintptr
	addr_QCalendarWidget_clicked uintptr
	addr_QCalendarWidget_activated uintptr
	addr_QCalendarWidget_currentPageChanged uintptr
	addr_QCalendarWidget_setTabOrder uintptr
	addr_QCalendarWidget_mouseGrabber uintptr
	addr_QCalendarWidget_keyboardGrabber uintptr
	addr_QCalendarWidget_find uintptr
	addr_QCalendarWidget_createWindowContainer uintptr
	addr_QCalendarWidget_updateMicroFocus uintptr
	addr_QCalendarWidget_create uintptr
	addr_QCalendarWidget_destroy uintptr
	addr_QCalendarWidget_focusNextChild uintptr
	addr_QCalendarWidget_focusPreviousChild uintptr
	addr_QCalendarWidget_connect uintptr
	addr_QCalendarWidget_connect2 uintptr
	addr_QCalendarWidget_sender uintptr
	addr_QCalendarWidget_senderSignalIndex uintptr
	addr_QCalendarWidget_receivers uintptr
	addr_QCalendarWidget_isSignalConnected uintptr
	pg_QCalendarWidget_devicePixelRatioFScale func() float64
	pg_QCalendarWidget_encodeMetricF func(int32, float64) int32
	pg_QCalendarWidget_getDecodedMetricF func(uintptr, int32, int32) float64
	addr_QCalendarWidget_tr2 uintptr
	addr_QCalendarWidget_tr3 uintptr
	addr_QCalendarWidget_createWindowContainer2 uintptr
	addr_QCalendarWidget_createWindowContainer3 uintptr
	addr_QCalendarWidget_updateMicroFocus2 uintptr
	addr_QCalendarWidget_create2 uintptr
	addr_QCalendarWidget_create3 uintptr
	addr_QCalendarWidget_create4 uintptr
	addr_QCalendarWidget_destroy2 uintptr
	addr_QCalendarWidget_destroy3 uintptr
	addr_QCalendarWidget_connect3 uintptr
	addr_QCalendarWidget_connect4 uintptr
	addr_QCalendarWidget_metaObject_default uintptr
	addr_QCalendarWidget_sizeHint_default uintptr
	addr_QCalendarWidget_minimumSizeHint_default uintptr
	addr_QCalendarWidget_event_default uintptr
	addr_QCalendarWidget_eventFilter_default uintptr
	addr_QCalendarWidget_mousePressEvent_default uintptr
	addr_QCalendarWidget_resizeEvent_default uintptr
	addr_QCalendarWidget_keyPressEvent_default uintptr
	addr_QCalendarWidget_paintCell_default uintptr
	addr_QCalendarWidget_devType_default uintptr
	addr_QCalendarWidget_setVisible_default uintptr
	addr_QCalendarWidget_heightForWidth_default uintptr
	addr_QCalendarWidget_hasHeightForWidth_default uintptr
	addr_QCalendarWidget_paintEngine_default uintptr
	addr_QCalendarWidget_mouseReleaseEvent_default uintptr
	addr_QCalendarWidget_mouseDoubleClickEvent_default uintptr
	addr_QCalendarWidget_mouseMoveEvent_default uintptr
	addr_QCalendarWidget_wheelEvent_default uintptr
	addr_QCalendarWidget_keyReleaseEvent_default uintptr
	addr_QCalendarWidget_focusInEvent_default uintptr
	addr_QCalendarWidget_focusOutEvent_default uintptr
	addr_QCalendarWidget_enterEvent_default uintptr
	addr_QCalendarWidget_leaveEvent_default uintptr
	addr_QCalendarWidget_paintEvent_default uintptr
	addr_QCalendarWidget_moveEvent_default uintptr
	addr_QCalendarWidget_closeEvent_default uintptr
	addr_QCalendarWidget_contextMenuEvent_default uintptr
	addr_QCalendarWidget_tabletEvent_default uintptr
	addr_QCalendarWidget_actionEvent_default uintptr
	addr_QCalendarWidget_dragEnterEvent_default uintptr
	addr_QCalendarWidget_dragMoveEvent_default uintptr
	addr_QCalendarWidget_dragLeaveEvent_default uintptr
	addr_QCalendarWidget_dropEvent_default uintptr
	addr_QCalendarWidget_showEvent_default uintptr
	addr_QCalendarWidget_hideEvent_default uintptr
	addr_QCalendarWidget_changeEvent_default uintptr
	addr_QCalendarWidget_metric_default uintptr
	addr_QCalendarWidget_initPainter_default uintptr
	addr_QCalendarWidget_redirected_default uintptr
	addr_QCalendarWidget_sharedPainter_default uintptr
	addr_QCalendarWidget_inputMethodEvent_default uintptr
	addr_QCalendarWidget_inputMethodQuery_default uintptr
	addr_QCalendarWidget_focusNextPrevChild_default uintptr
	addr_QCalendarWidget_timerEvent_default uintptr
	addr_QCalendarWidget_childEvent_default uintptr
	addr_QCalendarWidget_customEvent_default uintptr
	addr_QCalendarWidget_connectNotify_default uintptr
	addr_QCalendarWidget_disconnectNotify_default uintptr
	pg_QCalendarWidget_connect_selectionChanged func(uintptr, int64)
	pg_QCalendarWidget_connect_clicked func(uintptr, int64)
	pg_QCalendarWidget_connect_activated func(uintptr, int64)
	pg_QCalendarWidget_connect_currentPageChanged func(uintptr, int64)
	pg_QCalendarWidget_connect_windowTitleChanged func(uintptr, int64)
	pg_QCalendarWidget_connect_windowIconChanged func(uintptr, int64)
	pg_QCalendarWidget_connect_windowIconTextChanged func(uintptr, int64)
	pg_QCalendarWidget_connect_customContextMenuRequested func(uintptr, int64)
	pg_QCalendarWidget_connect_destroyed func(uintptr, int64)
	pg_QCalendarWidget_connect_destroyed2 func(uintptr, int64)
	pg_QCalendarWidget_connect_objectNameChanged func(uintptr, int64)
	pg_QCalendarWidget__disconnect func(int64)
	pg_QCalendarWidget__set_override func(uintptr, int32)
	pg_QCalendarWidget_delete func(uintptr)
)

func init() {
	symaddr(&addr_QCalendarWidget_chain, "QCalendarWidget_chain")
	symaddr(&addr_QCalendarWidget_new, "QCalendarWidget_new")
	symaddr(&addr_QCalendarWidget_new2, "QCalendarWidget_new2")
	symaddr(&addr_QCalendarWidget_tr, "QCalendarWidget_tr")
	symaddr(&addr_QCalendarWidget_selectedDate, "QCalendarWidget_selectedDate")
	symaddr(&addr_QCalendarWidget_yearShown, "QCalendarWidget_yearShown")
	symaddr(&addr_QCalendarWidget_monthShown, "QCalendarWidget_monthShown")
	symaddr(&addr_QCalendarWidget_minimumDate, "QCalendarWidget_minimumDate")
	symaddr(&addr_QCalendarWidget_setMinimumDate, "QCalendarWidget_setMinimumDate")
	symaddr(&addr_QCalendarWidget_clearMinimumDate, "QCalendarWidget_clearMinimumDate")
	symaddr(&addr_QCalendarWidget_maximumDate, "QCalendarWidget_maximumDate")
	symaddr(&addr_QCalendarWidget_setMaximumDate, "QCalendarWidget_setMaximumDate")
	symaddr(&addr_QCalendarWidget_clearMaximumDate, "QCalendarWidget_clearMaximumDate")
	symaddr(&addr_QCalendarWidget_firstDayOfWeek, "QCalendarWidget_firstDayOfWeek")
	symaddr(&addr_QCalendarWidget_setFirstDayOfWeek, "QCalendarWidget_setFirstDayOfWeek")
	symaddr(&addr_QCalendarWidget_isNavigationBarVisible, "QCalendarWidget_isNavigationBarVisible")
	symaddr(&addr_QCalendarWidget_isGridVisible, "QCalendarWidget_isGridVisible")
	symaddr(&addr_QCalendarWidget_calendar, "QCalendarWidget_calendar")
	symaddr(&addr_QCalendarWidget_setCalendar, "QCalendarWidget_setCalendar")
	symaddr(&addr_QCalendarWidget_selectionMode, "QCalendarWidget_selectionMode")
	symaddr(&addr_QCalendarWidget_setSelectionMode, "QCalendarWidget_setSelectionMode")
	symaddr(&addr_QCalendarWidget_horizontalHeaderFormat, "QCalendarWidget_horizontalHeaderFormat")
	symaddr(&addr_QCalendarWidget_setHorizontalHeaderFormat, "QCalendarWidget_setHorizontalHeaderFormat")
	symaddr(&addr_QCalendarWidget_verticalHeaderFormat, "QCalendarWidget_verticalHeaderFormat")
	symaddr(&addr_QCalendarWidget_setVerticalHeaderFormat, "QCalendarWidget_setVerticalHeaderFormat")
	symaddr(&addr_QCalendarWidget_headerTextFormat, "QCalendarWidget_headerTextFormat")
	symaddr(&addr_QCalendarWidget_setHeaderTextFormat, "QCalendarWidget_setHeaderTextFormat")
	symaddr(&addr_QCalendarWidget_weekdayTextFormat, "QCalendarWidget_weekdayTextFormat")
	symaddr(&addr_QCalendarWidget_setWeekdayTextFormat, "QCalendarWidget_setWeekdayTextFormat")
	symaddr(&addr_QCalendarWidget_dateTextFormat, "QCalendarWidget_dateTextFormat")
	symaddr(&addr_QCalendarWidget_setDateTextFormat, "QCalendarWidget_setDateTextFormat")
	symaddr(&addr_QCalendarWidget_isDateEditEnabled, "QCalendarWidget_isDateEditEnabled")
	symaddr(&addr_QCalendarWidget_setDateEditEnabled, "QCalendarWidget_setDateEditEnabled")
	symaddr(&addr_QCalendarWidget_dateEditAcceptDelay, "QCalendarWidget_dateEditAcceptDelay")
	symaddr(&addr_QCalendarWidget_setDateEditAcceptDelay, "QCalendarWidget_setDateEditAcceptDelay")
	symaddr(&addr_QCalendarWidget_updateCell, "QCalendarWidget_updateCell")
	symaddr(&addr_QCalendarWidget_updateCells, "QCalendarWidget_updateCells")
	symaddr(&addr_QCalendarWidget_setSelectedDate, "QCalendarWidget_setSelectedDate")
	symaddr(&addr_QCalendarWidget_setDateRange, "QCalendarWidget_setDateRange")
	symaddr(&addr_QCalendarWidget_setCurrentPage, "QCalendarWidget_setCurrentPage")
	symaddr(&addr_QCalendarWidget_setGridVisible, "QCalendarWidget_setGridVisible")
	symaddr(&addr_QCalendarWidget_setNavigationBarVisible, "QCalendarWidget_setNavigationBarVisible")
	symaddr(&addr_QCalendarWidget_showNextMonth, "QCalendarWidget_showNextMonth")
	symaddr(&addr_QCalendarWidget_showPreviousMonth, "QCalendarWidget_showPreviousMonth")
	symaddr(&addr_QCalendarWidget_showNextYear, "QCalendarWidget_showNextYear")
	symaddr(&addr_QCalendarWidget_showPreviousYear, "QCalendarWidget_showPreviousYear")
	symaddr(&addr_QCalendarWidget_showSelectedDate, "QCalendarWidget_showSelectedDate")
	symaddr(&addr_QCalendarWidget_showToday, "QCalendarWidget_showToday")
	symaddr(&addr_QCalendarWidget_selectionChanged, "QCalendarWidget_selectionChanged")
	symaddr(&addr_QCalendarWidget_clicked, "QCalendarWidget_clicked")
	symaddr(&addr_QCalendarWidget_activated, "QCalendarWidget_activated")
	symaddr(&addr_QCalendarWidget_currentPageChanged, "QCalendarWidget_currentPageChanged")
	symaddr(&addr_QCalendarWidget_setTabOrder, "QCalendarWidget_setTabOrder")
	symaddr(&addr_QCalendarWidget_mouseGrabber, "QCalendarWidget_mouseGrabber")
	symaddr(&addr_QCalendarWidget_keyboardGrabber, "QCalendarWidget_keyboardGrabber")
	symaddr(&addr_QCalendarWidget_find, "QCalendarWidget_find")
	symaddr(&addr_QCalendarWidget_createWindowContainer, "QCalendarWidget_createWindowContainer")
	symaddr(&addr_QCalendarWidget_updateMicroFocus, "QCalendarWidget_updateMicroFocus")
	symaddr(&addr_QCalendarWidget_create, "QCalendarWidget_create")
	symaddr(&addr_QCalendarWidget_destroy, "QCalendarWidget_destroy")
	symaddr(&addr_QCalendarWidget_focusNextChild, "QCalendarWidget_focusNextChild")
	symaddr(&addr_QCalendarWidget_focusPreviousChild, "QCalendarWidget_focusPreviousChild")
	symaddr(&addr_QCalendarWidget_connect, "QCalendarWidget_connect")
	symaddr(&addr_QCalendarWidget_connect2, "QCalendarWidget_connect2")
	symaddr(&addr_QCalendarWidget_sender, "QCalendarWidget_sender")
	symaddr(&addr_QCalendarWidget_senderSignalIndex, "QCalendarWidget_senderSignalIndex")
	symaddr(&addr_QCalendarWidget_receivers, "QCalendarWidget_receivers")
	symaddr(&addr_QCalendarWidget_isSignalConnected, "QCalendarWidget_isSignalConnected")
	reg(&pg_QCalendarWidget_devicePixelRatioFScale, "QCalendarWidget_devicePixelRatioFScale")
	reg(&pg_QCalendarWidget_encodeMetricF, "QCalendarWidget_encodeMetricF")
	reg(&pg_QCalendarWidget_getDecodedMetricF, "QCalendarWidget_getDecodedMetricF")
	symaddr(&addr_QCalendarWidget_tr2, "QCalendarWidget_tr2")
	symaddr(&addr_QCalendarWidget_tr3, "QCalendarWidget_tr3")
	symaddr(&addr_QCalendarWidget_createWindowContainer2, "QCalendarWidget_createWindowContainer2")
	symaddr(&addr_QCalendarWidget_createWindowContainer3, "QCalendarWidget_createWindowContainer3")
	symaddr(&addr_QCalendarWidget_updateMicroFocus2, "QCalendarWidget_updateMicroFocus2")
	symaddr(&addr_QCalendarWidget_create2, "QCalendarWidget_create2")
	symaddr(&addr_QCalendarWidget_create3, "QCalendarWidget_create3")
	symaddr(&addr_QCalendarWidget_create4, "QCalendarWidget_create4")
	symaddr(&addr_QCalendarWidget_destroy2, "QCalendarWidget_destroy2")
	symaddr(&addr_QCalendarWidget_destroy3, "QCalendarWidget_destroy3")
	symaddr(&addr_QCalendarWidget_connect3, "QCalendarWidget_connect3")
	symaddr(&addr_QCalendarWidget_connect4, "QCalendarWidget_connect4")
	symaddr(&addr_QCalendarWidget_metaObject_default, "QCalendarWidget_metaObject_default")
	symaddr(&addr_QCalendarWidget_sizeHint_default, "QCalendarWidget_sizeHint_default")
	symaddr(&addr_QCalendarWidget_minimumSizeHint_default, "QCalendarWidget_minimumSizeHint_default")
	symaddr(&addr_QCalendarWidget_event_default, "QCalendarWidget_event_default")
	symaddr(&addr_QCalendarWidget_eventFilter_default, "QCalendarWidget_eventFilter_default")
	symaddr(&addr_QCalendarWidget_mousePressEvent_default, "QCalendarWidget_mousePressEvent_default")
	symaddr(&addr_QCalendarWidget_resizeEvent_default, "QCalendarWidget_resizeEvent_default")
	symaddr(&addr_QCalendarWidget_keyPressEvent_default, "QCalendarWidget_keyPressEvent_default")
	symaddr(&addr_QCalendarWidget_paintCell_default, "QCalendarWidget_paintCell_default")
	symaddr(&addr_QCalendarWidget_devType_default, "QCalendarWidget_devType_default")
	symaddr(&addr_QCalendarWidget_setVisible_default, "QCalendarWidget_setVisible_default")
	symaddr(&addr_QCalendarWidget_heightForWidth_default, "QCalendarWidget_heightForWidth_default")
	symaddr(&addr_QCalendarWidget_hasHeightForWidth_default, "QCalendarWidget_hasHeightForWidth_default")
	symaddr(&addr_QCalendarWidget_paintEngine_default, "QCalendarWidget_paintEngine_default")
	symaddr(&addr_QCalendarWidget_mouseReleaseEvent_default, "QCalendarWidget_mouseReleaseEvent_default")
	symaddr(&addr_QCalendarWidget_mouseDoubleClickEvent_default, "QCalendarWidget_mouseDoubleClickEvent_default")
	symaddr(&addr_QCalendarWidget_mouseMoveEvent_default, "QCalendarWidget_mouseMoveEvent_default")
	symaddr(&addr_QCalendarWidget_wheelEvent_default, "QCalendarWidget_wheelEvent_default")
	symaddr(&addr_QCalendarWidget_keyReleaseEvent_default, "QCalendarWidget_keyReleaseEvent_default")
	symaddr(&addr_QCalendarWidget_focusInEvent_default, "QCalendarWidget_focusInEvent_default")
	symaddr(&addr_QCalendarWidget_focusOutEvent_default, "QCalendarWidget_focusOutEvent_default")
	symaddr(&addr_QCalendarWidget_enterEvent_default, "QCalendarWidget_enterEvent_default")
	symaddr(&addr_QCalendarWidget_leaveEvent_default, "QCalendarWidget_leaveEvent_default")
	symaddr(&addr_QCalendarWidget_paintEvent_default, "QCalendarWidget_paintEvent_default")
	symaddr(&addr_QCalendarWidget_moveEvent_default, "QCalendarWidget_moveEvent_default")
	symaddr(&addr_QCalendarWidget_closeEvent_default, "QCalendarWidget_closeEvent_default")
	symaddr(&addr_QCalendarWidget_contextMenuEvent_default, "QCalendarWidget_contextMenuEvent_default")
	symaddr(&addr_QCalendarWidget_tabletEvent_default, "QCalendarWidget_tabletEvent_default")
	symaddr(&addr_QCalendarWidget_actionEvent_default, "QCalendarWidget_actionEvent_default")
	symaddr(&addr_QCalendarWidget_dragEnterEvent_default, "QCalendarWidget_dragEnterEvent_default")
	symaddr(&addr_QCalendarWidget_dragMoveEvent_default, "QCalendarWidget_dragMoveEvent_default")
	symaddr(&addr_QCalendarWidget_dragLeaveEvent_default, "QCalendarWidget_dragLeaveEvent_default")
	symaddr(&addr_QCalendarWidget_dropEvent_default, "QCalendarWidget_dropEvent_default")
	symaddr(&addr_QCalendarWidget_showEvent_default, "QCalendarWidget_showEvent_default")
	symaddr(&addr_QCalendarWidget_hideEvent_default, "QCalendarWidget_hideEvent_default")
	symaddr(&addr_QCalendarWidget_changeEvent_default, "QCalendarWidget_changeEvent_default")
	symaddr(&addr_QCalendarWidget_metric_default, "QCalendarWidget_metric_default")
	symaddr(&addr_QCalendarWidget_initPainter_default, "QCalendarWidget_initPainter_default")
	symaddr(&addr_QCalendarWidget_redirected_default, "QCalendarWidget_redirected_default")
	symaddr(&addr_QCalendarWidget_sharedPainter_default, "QCalendarWidget_sharedPainter_default")
	symaddr(&addr_QCalendarWidget_inputMethodEvent_default, "QCalendarWidget_inputMethodEvent_default")
	symaddr(&addr_QCalendarWidget_inputMethodQuery_default, "QCalendarWidget_inputMethodQuery_default")
	symaddr(&addr_QCalendarWidget_focusNextPrevChild_default, "QCalendarWidget_focusNextPrevChild_default")
	symaddr(&addr_QCalendarWidget_timerEvent_default, "QCalendarWidget_timerEvent_default")
	symaddr(&addr_QCalendarWidget_childEvent_default, "QCalendarWidget_childEvent_default")
	symaddr(&addr_QCalendarWidget_customEvent_default, "QCalendarWidget_customEvent_default")
	symaddr(&addr_QCalendarWidget_connectNotify_default, "QCalendarWidget_connectNotify_default")
	symaddr(&addr_QCalendarWidget_disconnectNotify_default, "QCalendarWidget_disconnectNotify_default")
	reg(&pg_QCalendarWidget_connect_selectionChanged, "QCalendarWidget_connect_selectionChanged")
	reg(&pg_QCalendarWidget_connect_clicked, "QCalendarWidget_connect_clicked")
	reg(&pg_QCalendarWidget_connect_activated, "QCalendarWidget_connect_activated")
	reg(&pg_QCalendarWidget_connect_currentPageChanged, "QCalendarWidget_connect_currentPageChanged")
	reg(&pg_QCalendarWidget_connect_windowTitleChanged, "QCalendarWidget_connect_windowTitleChanged")
	reg(&pg_QCalendarWidget_connect_windowIconChanged, "QCalendarWidget_connect_windowIconChanged")
	reg(&pg_QCalendarWidget_connect_windowIconTextChanged, "QCalendarWidget_connect_windowIconTextChanged")
	reg(&pg_QCalendarWidget_connect_customContextMenuRequested, "QCalendarWidget_connect_customContextMenuRequested")
	reg(&pg_QCalendarWidget_connect_destroyed, "QCalendarWidget_connect_destroyed")
	reg(&pg_QCalendarWidget_connect_destroyed2, "QCalendarWidget_connect_destroyed2")
	reg(&pg_QCalendarWidget_connect_objectNameChanged, "QCalendarWidget_connect_objectNameChanged")
	reg(&pg_QCalendarWidget__disconnect, "QCalendarWidget__disconnect")
	reg(&pg_QCalendarWidget__set_override, "QCalendarWidget__set_override")
	reg(&pg_QCalendarWidget_delete, "QCalendarWidget_delete")
}

type QCalendarWidget struct {
	h uintptr
	*QWidget
}

func bindQCalendarWidget(h uintptr) *QCalendarWidget {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QCalendarWidget_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QCalendarWidget{h: h, QWidget: &QWidget{h: __c[0], QObject: &QObject{h: __c[1]}}}
}

func newownedQCalendarWidget(h uintptr) *QCalendarWidget { return bindQCalendarWidget(h) }

func (this *QCalendarWidget) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QCalendarWidget) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQCalendarWidget(h unsafe.Pointer) *QCalendarWidget { return bindQCalendarWidget(uintptr(h)) }
func UnsafeNewQCalendarWidgetOwned(h unsafe.Pointer) *QCalendarWidget { return newownedQCalendarWidget(uintptr(h)) }

func (this *QCalendarWidget) UnsafeRelease() {
	this.h = 0
	if this.QWidget != nil {
		this.QWidget.UnsafeRelease()
	}
}

func NewQCalendarWidget(parent *QWidget) *QCalendarWidget {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QCalendarWidget_new, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQCalendarWidget(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QCalendarWidget_virtIdx, QCalendarWidget_enable)
	return __o
}

func NewQCalendarWidget2() *QCalendarWidget {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QCalendarWidget_new2, uintptr(__goid))
	__o := newownedQCalendarWidget(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QCalendarWidget_virtIdx, QCalendarWidget_enable)
	return __o
}

func QCalendarWidget_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QCalendarWidget_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QCalendarWidget) SelectedDate() *QDate {
	__r, _, _ := purego.SyscallN(addr_QCalendarWidget_selectedDate, this.h)
	runtime.KeepAlive(this)
	return newownedQDate(__r)
}

func (this *QCalendarWidget) YearShown() int {
	__r0, _, _ := purego.SyscallN(addr_QCalendarWidget_yearShown, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QCalendarWidget) MonthShown() int {
	__r0, _, _ := purego.SyscallN(addr_QCalendarWidget_monthShown, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QCalendarWidget) MinimumDate() *QDate {
	__r, _, _ := purego.SyscallN(addr_QCalendarWidget_minimumDate, this.h)
	runtime.KeepAlive(this)
	return newownedQDate(__r)
}

func (this *QCalendarWidget) SetMinimumDate(date *QDate) {
	purego.SyscallN(addr_QCalendarWidget_setMinimumDate, this.h, uintptr(date.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(date)
}

func (this *QCalendarWidget) ClearMinimumDate() {
	purego.SyscallN(addr_QCalendarWidget_clearMinimumDate, this.h)
	runtime.KeepAlive(this)
}

func (this *QCalendarWidget) MaximumDate() *QDate {
	__r, _, _ := purego.SyscallN(addr_QCalendarWidget_maximumDate, this.h)
	runtime.KeepAlive(this)
	return newownedQDate(__r)
}

func (this *QCalendarWidget) SetMaximumDate(date *QDate) {
	purego.SyscallN(addr_QCalendarWidget_setMaximumDate, this.h, uintptr(date.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(date)
}

func (this *QCalendarWidget) ClearMaximumDate() {
	purego.SyscallN(addr_QCalendarWidget_clearMaximumDate, this.h)
	runtime.KeepAlive(this)
}

func (this *QCalendarWidget) FirstDayOfWeek() DayOfWeek {
	__r0, _, _ := purego.SyscallN(addr_QCalendarWidget_firstDayOfWeek, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return DayOfWeek(__r)
}

func (this *QCalendarWidget) SetFirstDayOfWeek(dayOfWeek DayOfWeek) {
	purego.SyscallN(addr_QCalendarWidget_setFirstDayOfWeek, this.h, uintptr(dayOfWeek))
	runtime.KeepAlive(this)
}

func (this *QCalendarWidget) IsNavigationBarVisible() bool {
	__r0, _, _ := purego.SyscallN(addr_QCalendarWidget_isNavigationBarVisible, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCalendarWidget) IsGridVisible() bool {
	__r0, _, _ := purego.SyscallN(addr_QCalendarWidget_isGridVisible, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCalendarWidget) Calendar() *QCalendar {
	__r, _, _ := purego.SyscallN(addr_QCalendarWidget_calendar, this.h)
	runtime.KeepAlive(this)
	return newownedQCalendar(__r)
}

func (this *QCalendarWidget) SetCalendar(calendar *QCalendar) {
	purego.SyscallN(addr_QCalendarWidget_setCalendar, this.h, uintptr(calendar.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(calendar)
}

func (this *QCalendarWidget) SelectionMode() QCalendarWidget__SelectionMode {
	__r0, _, _ := purego.SyscallN(addr_QCalendarWidget_selectionMode, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QCalendarWidget__SelectionMode(__r)
}

func (this *QCalendarWidget) SetSelectionMode(mode QCalendarWidget__SelectionMode) {
	purego.SyscallN(addr_QCalendarWidget_setSelectionMode, this.h, uintptr(mode))
	runtime.KeepAlive(this)
}

func (this *QCalendarWidget) HorizontalHeaderFormat() QCalendarWidget__HorizontalHeaderFormat {
	__r0, _, _ := purego.SyscallN(addr_QCalendarWidget_horizontalHeaderFormat, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QCalendarWidget__HorizontalHeaderFormat(__r)
}

func (this *QCalendarWidget) SetHorizontalHeaderFormat(format QCalendarWidget__HorizontalHeaderFormat) {
	purego.SyscallN(addr_QCalendarWidget_setHorizontalHeaderFormat, this.h, uintptr(format))
	runtime.KeepAlive(this)
}

func (this *QCalendarWidget) VerticalHeaderFormat() QCalendarWidget__VerticalHeaderFormat {
	__r0, _, _ := purego.SyscallN(addr_QCalendarWidget_verticalHeaderFormat, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QCalendarWidget__VerticalHeaderFormat(__r)
}

func (this *QCalendarWidget) SetVerticalHeaderFormat(format QCalendarWidget__VerticalHeaderFormat) {
	purego.SyscallN(addr_QCalendarWidget_setVerticalHeaderFormat, this.h, uintptr(format))
	runtime.KeepAlive(this)
}

func (this *QCalendarWidget) HeaderTextFormat() *QTextCharFormat {
	__r, _, _ := purego.SyscallN(addr_QCalendarWidget_headerTextFormat, this.h)
	runtime.KeepAlive(this)
	return newownedQTextCharFormat(__r)
}

func (this *QCalendarWidget) SetHeaderTextFormat(format *QTextCharFormat) {
	purego.SyscallN(addr_QCalendarWidget_setHeaderTextFormat, this.h, uintptr(format.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(format)
}

func (this *QCalendarWidget) WeekdayTextFormat(dayOfWeek DayOfWeek) *QTextCharFormat {
	__r, _, _ := purego.SyscallN(addr_QCalendarWidget_weekdayTextFormat, this.h, uintptr(dayOfWeek))
	runtime.KeepAlive(this)
	return newownedQTextCharFormat(__r)
}

func (this *QCalendarWidget) SetWeekdayTextFormat(dayOfWeek DayOfWeek, format *QTextCharFormat) {
	purego.SyscallN(addr_QCalendarWidget_setWeekdayTextFormat, this.h, uintptr(dayOfWeek), uintptr(format.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(format)
}

func (this *QCalendarWidget) DateTextFormat(date *QDate) *QTextCharFormat {
	__r, _, _ := purego.SyscallN(addr_QCalendarWidget_dateTextFormat, this.h, uintptr(date.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(date)
	return newownedQTextCharFormat(__r)
}

func (this *QCalendarWidget) SetDateTextFormat(date *QDate, format *QTextCharFormat) {
	purego.SyscallN(addr_QCalendarWidget_setDateTextFormat, this.h, uintptr(date.UnsafePointer()), uintptr(format.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(date)
	runtime.KeepAlive(format)
}

func (this *QCalendarWidget) IsDateEditEnabled() bool {
	__r0, _, _ := purego.SyscallN(addr_QCalendarWidget_isDateEditEnabled, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCalendarWidget) SetDateEditEnabled(enable bool) {
	purego.SyscallN(addr_QCalendarWidget_setDateEditEnabled, this.h, uintptr(b2i(enable)))
	runtime.KeepAlive(this)
}

func (this *QCalendarWidget) DateEditAcceptDelay() int {
	__r0, _, _ := purego.SyscallN(addr_QCalendarWidget_dateEditAcceptDelay, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QCalendarWidget) SetDateEditAcceptDelay(delay int) {
	purego.SyscallN(addr_QCalendarWidget_setDateEditAcceptDelay, this.h, uintptr(delay))
	runtime.KeepAlive(this)
}

func (this *QCalendarWidget) UpdateCell(date *QDate) {
	purego.SyscallN(addr_QCalendarWidget_updateCell, this.h, uintptr(date.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(date)
}

func (this *QCalendarWidget) UpdateCells() {
	purego.SyscallN(addr_QCalendarWidget_updateCells, this.h)
	runtime.KeepAlive(this)
}

func (this *QCalendarWidget) SetSelectedDate(date *QDate) {
	purego.SyscallN(addr_QCalendarWidget_setSelectedDate, this.h, uintptr(date.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(date)
}

func (this *QCalendarWidget) SetDateRange(min *QDate, max *QDate) {
	purego.SyscallN(addr_QCalendarWidget_setDateRange, this.h, uintptr(min.UnsafePointer()), uintptr(max.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(min)
	runtime.KeepAlive(max)
}

func (this *QCalendarWidget) SetCurrentPage(year int, month int) {
	purego.SyscallN(addr_QCalendarWidget_setCurrentPage, this.h, uintptr(year), uintptr(month))
	runtime.KeepAlive(this)
}

func (this *QCalendarWidget) SetGridVisible(show bool) {
	purego.SyscallN(addr_QCalendarWidget_setGridVisible, this.h, uintptr(b2i(show)))
	runtime.KeepAlive(this)
}

func (this *QCalendarWidget) SetNavigationBarVisible(visible bool) {
	purego.SyscallN(addr_QCalendarWidget_setNavigationBarVisible, this.h, uintptr(b2i(visible)))
	runtime.KeepAlive(this)
}

func (this *QCalendarWidget) ShowNextMonth() {
	purego.SyscallN(addr_QCalendarWidget_showNextMonth, this.h)
	runtime.KeepAlive(this)
}

func (this *QCalendarWidget) ShowPreviousMonth() {
	purego.SyscallN(addr_QCalendarWidget_showPreviousMonth, this.h)
	runtime.KeepAlive(this)
}

func (this *QCalendarWidget) ShowNextYear() {
	purego.SyscallN(addr_QCalendarWidget_showNextYear, this.h)
	runtime.KeepAlive(this)
}

func (this *QCalendarWidget) ShowPreviousYear() {
	purego.SyscallN(addr_QCalendarWidget_showPreviousYear, this.h)
	runtime.KeepAlive(this)
}

func (this *QCalendarWidget) ShowSelectedDate() {
	purego.SyscallN(addr_QCalendarWidget_showSelectedDate, this.h)
	runtime.KeepAlive(this)
}

func (this *QCalendarWidget) ShowToday() {
	purego.SyscallN(addr_QCalendarWidget_showToday, this.h)
	runtime.KeepAlive(this)
}

func (this *QCalendarWidget) SelectionChanged() {
	purego.SyscallN(addr_QCalendarWidget_selectionChanged, this.h)
	runtime.KeepAlive(this)
}

func (this *QCalendarWidget) Clicked(date *QDate) {
	purego.SyscallN(addr_QCalendarWidget_clicked, this.h, uintptr(date.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(date)
}

func (this *QCalendarWidget) Activated(date *QDate) {
	purego.SyscallN(addr_QCalendarWidget_activated, this.h, uintptr(date.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(date)
}

func (this *QCalendarWidget) CurrentPageChanged(year int, month int) {
	purego.SyscallN(addr_QCalendarWidget_currentPageChanged, this.h, uintptr(year), uintptr(month))
	runtime.KeepAlive(this)
}

func QCalendarWidget_SetTabOrder(arg1 *QWidget, arg2 *QWidget) {
	purego.SyscallN(addr_QCalendarWidget_setTabOrder, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()))
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
}

func QCalendarWidget_MouseGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QCalendarWidget_mouseGrabber)
	return bindQWidget(__r)
}

func QCalendarWidget_KeyboardGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QCalendarWidget_keyboardGrabber)
	return bindQWidget(__r)
}

func QCalendarWidget_Find(arg1 uint64) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QCalendarWidget_find, uintptr(arg1))
	return bindQWidget(__r)
}

func QCalendarWidget_CreateWindowContainer(window *QWindow) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QCalendarWidget_createWindowContainer, uintptr(window.UnsafePointer()))
	runtime.KeepAlive(window)
	return bindQWidget(__r)
}

func (this *QCalendarWidget) UpdateMicroFocus() {
	purego.SyscallN(addr_QCalendarWidget_updateMicroFocus, this.h)
	runtime.KeepAlive(this)
}

func (this *QCalendarWidget) Create() {
	purego.SyscallN(addr_QCalendarWidget_create, this.h)
	runtime.KeepAlive(this)
}

func (this *QCalendarWidget) Destroy() {
	purego.SyscallN(addr_QCalendarWidget_destroy, this.h)
	runtime.KeepAlive(this)
}

func (this *QCalendarWidget) FocusNextChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QCalendarWidget_focusNextChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCalendarWidget) FocusPreviousChild() bool {
	__r0, _, _ := purego.SyscallN(addr_QCalendarWidget_focusPreviousChild, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QCalendarWidget_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QCalendarWidget_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QCalendarWidget_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QCalendarWidget_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QCalendarWidget) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QCalendarWidget_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QCalendarWidget) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QCalendarWidget_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QCalendarWidget) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QCalendarWidget_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QCalendarWidget) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QCalendarWidget_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QCalendarWidget_DevicePixelRatioFScale() float64 {
	__r := pg_QCalendarWidget_devicePixelRatioFScale()
	return __r
}

func QCalendarWidget_EncodeMetricF(metric QPaintDevice__PaintDeviceMetric, value float64) int {
	__r := pg_QCalendarWidget_encodeMetricF(int32(metric), value)
	return int(__r)
}

func (this *QCalendarWidget) GetDecodedMetricF(metricA QPaintDevice__PaintDeviceMetric, metricB QPaintDevice__PaintDeviceMetric) float64 {
	__r := pg_QCalendarWidget_getDecodedMetricF(this.h, int32(metricA), int32(metricB))
	runtime.KeepAlive(this)
	return __r
}

func QCalendarWidget_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QCalendarWidget_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QCalendarWidget_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QCalendarWidget_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QCalendarWidget_CreateWindowContainer2(window *QWindow, parent *QWidget) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QCalendarWidget_createWindowContainer2, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func QCalendarWidget_CreateWindowContainer3(window *QWindow, parent *QWidget, flags WindowType) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QCalendarWidget_createWindowContainer3, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()), uintptr(flags))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func (this *QCalendarWidget) UpdateMicroFocusWithQuery(query InputMethodQuery) {
	purego.SyscallN(addr_QCalendarWidget_updateMicroFocus2, this.h, uintptr(query))
	runtime.KeepAlive(this)
}

func (this *QCalendarWidget) CreateWith(arg1 uint64) {
	purego.SyscallN(addr_QCalendarWidget_create2, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QCalendarWidget) Create2(arg1 uint64, initializeWindow bool) {
	purego.SyscallN(addr_QCalendarWidget_create3, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)))
	runtime.KeepAlive(this)
}

func (this *QCalendarWidget) Create3(arg1 uint64, initializeWindow bool, destroyOldWindow bool) {
	purego.SyscallN(addr_QCalendarWidget_create4, this.h, uintptr(arg1), uintptr(b2i(initializeWindow)), uintptr(b2i(destroyOldWindow)))
	runtime.KeepAlive(this)
}

func (this *QCalendarWidget) DestroyWithDestroyWindow(destroyWindow bool) {
	purego.SyscallN(addr_QCalendarWidget_destroy2, this.h, uintptr(b2i(destroyWindow)))
	runtime.KeepAlive(this)
}

func (this *QCalendarWidget) Destroy2(destroyWindow bool, destroySubWindows bool) {
	purego.SyscallN(addr_QCalendarWidget_destroy3, this.h, uintptr(b2i(destroyWindow)), uintptr(b2i(destroySubWindows)))
	runtime.KeepAlive(this)
}

func QCalendarWidget_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QCalendarWidget_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QCalendarWidget_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QCalendarWidget_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QCalendarWidget) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QCalendarWidget_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QCalendarWidget) SizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QCalendarWidget_sizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QCalendarWidget) MinimumSizeHintDefault() *QSize {
	__r, _, _ := purego.SyscallN(addr_QCalendarWidget_minimumSizeHint_default, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QCalendarWidget) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QCalendarWidget_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QCalendarWidget) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QCalendarWidget_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QCalendarWidget) MousePressEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QCalendarWidget_mousePressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QCalendarWidget) ResizeEventDefault(event *QResizeEvent) {
	purego.SyscallN(addr_QCalendarWidget_resizeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QCalendarWidget) KeyPressEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QCalendarWidget_keyPressEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QCalendarWidget) PaintCellDefault(painter *QPainter, rect *QRect, date *QDate) {
	purego.SyscallN(addr_QCalendarWidget_paintCell_default, this.h, uintptr(painter.UnsafePointer()), uintptr(rect.UnsafePointer()), uintptr(date.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
	runtime.KeepAlive(rect)
	runtime.KeepAlive(date)
}

func (this *QCalendarWidget) DevTypeDefault() int {
	__r0, _, _ := purego.SyscallN(addr_QCalendarWidget_devType_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QCalendarWidget) SetVisibleDefault(visible bool) {
	purego.SyscallN(addr_QCalendarWidget_setVisible_default, this.h, uintptr(b2i(visible)))
	runtime.KeepAlive(this)
}

func (this *QCalendarWidget) HeightForWidthDefault(arg1 int) int {
	__r0, _, _ := purego.SyscallN(addr_QCalendarWidget_heightForWidth_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QCalendarWidget) HasHeightForWidthDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QCalendarWidget_hasHeightForWidth_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCalendarWidget) PaintEngineDefault() *QPaintEngine {
	__r, _, _ := purego.SyscallN(addr_QCalendarWidget_paintEngine_default, this.h)
	runtime.KeepAlive(this)
	return bindQPaintEngine(__r)
}

func (this *QCalendarWidget) MouseReleaseEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QCalendarWidget_mouseReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QCalendarWidget) MouseDoubleClickEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QCalendarWidget_mouseDoubleClickEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QCalendarWidget) MouseMoveEventDefault(event *QMouseEvent) {
	purego.SyscallN(addr_QCalendarWidget_mouseMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QCalendarWidget) WheelEventDefault(event *QWheelEvent) {
	purego.SyscallN(addr_QCalendarWidget_wheelEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QCalendarWidget) KeyReleaseEventDefault(event *QKeyEvent) {
	purego.SyscallN(addr_QCalendarWidget_keyReleaseEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QCalendarWidget) FocusInEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QCalendarWidget_focusInEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QCalendarWidget) FocusOutEventDefault(event *QFocusEvent) {
	purego.SyscallN(addr_QCalendarWidget_focusOutEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QCalendarWidget) EnterEventDefault(event *QEnterEvent) {
	purego.SyscallN(addr_QCalendarWidget_enterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QCalendarWidget) LeaveEventDefault(event *QEvent) {
	purego.SyscallN(addr_QCalendarWidget_leaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QCalendarWidget) PaintEventDefault(event *QPaintEvent) {
	purego.SyscallN(addr_QCalendarWidget_paintEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QCalendarWidget) MoveEventDefault(event *QMoveEvent) {
	purego.SyscallN(addr_QCalendarWidget_moveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QCalendarWidget) CloseEventDefault(event *QCloseEvent) {
	purego.SyscallN(addr_QCalendarWidget_closeEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QCalendarWidget) ContextMenuEventDefault(event *QContextMenuEvent) {
	purego.SyscallN(addr_QCalendarWidget_contextMenuEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QCalendarWidget) TabletEventDefault(event *QTabletEvent) {
	purego.SyscallN(addr_QCalendarWidget_tabletEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QCalendarWidget) ActionEventDefault(event *QActionEvent) {
	purego.SyscallN(addr_QCalendarWidget_actionEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QCalendarWidget) DragEnterEventDefault(event *QDragEnterEvent) {
	purego.SyscallN(addr_QCalendarWidget_dragEnterEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QCalendarWidget) DragMoveEventDefault(event *QDragMoveEvent) {
	purego.SyscallN(addr_QCalendarWidget_dragMoveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QCalendarWidget) DragLeaveEventDefault(event *QDragLeaveEvent) {
	purego.SyscallN(addr_QCalendarWidget_dragLeaveEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QCalendarWidget) DropEventDefault(event *QDropEvent) {
	purego.SyscallN(addr_QCalendarWidget_dropEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QCalendarWidget) ShowEventDefault(event *QShowEvent) {
	purego.SyscallN(addr_QCalendarWidget_showEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QCalendarWidget) HideEventDefault(event *QHideEvent) {
	purego.SyscallN(addr_QCalendarWidget_hideEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QCalendarWidget) ChangeEventDefault(arg1 *QEvent) {
	purego.SyscallN(addr_QCalendarWidget_changeEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QCalendarWidget) MetricDefault(arg1 QPaintDevice__PaintDeviceMetric) int {
	__r0, _, _ := purego.SyscallN(addr_QCalendarWidget_metric_default, this.h, uintptr(arg1))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QCalendarWidget) InitPainterDefault(painter *QPainter) {
	purego.SyscallN(addr_QCalendarWidget_initPainter_default, this.h, uintptr(painter.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
}

func (this *QCalendarWidget) RedirectedDefault(offset *QPoint) *QPaintDevice {
	__r, _, _ := purego.SyscallN(addr_QCalendarWidget_redirected_default, this.h, uintptr(offset.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(offset)
	return bindQPaintDevice(__r)
}

func (this *QCalendarWidget) SharedPainterDefault() *QPainter {
	__r, _, _ := purego.SyscallN(addr_QCalendarWidget_sharedPainter_default, this.h)
	runtime.KeepAlive(this)
	return bindQPainter(__r)
}

func (this *QCalendarWidget) InputMethodEventDefault(arg1 *QInputMethodEvent) {
	purego.SyscallN(addr_QCalendarWidget_inputMethodEvent_default, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QCalendarWidget) InputMethodQueryDefault(arg1 InputMethodQuery) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QCalendarWidget_inputMethodQuery_default, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QCalendarWidget) FocusNextPrevChildDefault(next bool) bool {
	__r0, _, _ := purego.SyscallN(addr_QCalendarWidget_focusNextPrevChild_default, this.h, uintptr(b2i(next)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QCalendarWidget) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QCalendarWidget_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QCalendarWidget) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QCalendarWidget_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QCalendarWidget) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QCalendarWidget_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QCalendarWidget) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QCalendarWidget_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QCalendarWidget) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QCalendarWidget_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QCalendarWidget) OnSelectionChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QCalendarWidget_connect_selectionChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QCalendarWidget) OnClicked(fn func(*QDate)) int64 {
	id := registerCallback(func(date uintptr) { fn(UnsafeNewQDate(unsafe.Pointer(date))) })
	pg_QCalendarWidget_connect_clicked(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QCalendarWidget) OnActivated(fn func(*QDate)) int64 {
	id := registerCallback(func(date uintptr) { fn(UnsafeNewQDate(unsafe.Pointer(date))) })
	pg_QCalendarWidget_connect_activated(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QCalendarWidget) OnCurrentPageChanged(fn func(int, int)) int64 {
	id := registerCallback(func(year int32, month int32) { fn(int(year), int(month)) })
	pg_QCalendarWidget_connect_currentPageChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QCalendarWidget) OnWindowTitleChanged(fn func(string)) int64 {
	id := registerCallback(func(title uintptr, title_len int64) { fn(pgStrNoFree(title, title_len)) })
	pg_QCalendarWidget_connect_windowTitleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QCalendarWidget) OnWindowIconChanged(fn func(*QIcon)) int64 {
	id := registerCallback(func(icon uintptr) { fn(UnsafeNewQIcon(unsafe.Pointer(icon))) })
	pg_QCalendarWidget_connect_windowIconChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QCalendarWidget) OnWindowIconTextChanged(fn func(string)) int64 {
	id := registerCallback(func(iconText uintptr, iconText_len int64) { fn(pgStrNoFree(iconText, iconText_len)) })
	pg_QCalendarWidget_connect_windowIconTextChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QCalendarWidget) OnCustomContextMenuRequested(fn func(*QPoint)) int64 {
	id := registerCallback(func(pos uintptr) { fn(UnsafeNewQPoint(unsafe.Pointer(pos))) })
	pg_QCalendarWidget_connect_customContextMenuRequested(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QCalendarWidget) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QCalendarWidget_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QCalendarWidget) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QCalendarWidget_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QCalendarWidget) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QCalendarWidget_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QCalendarWidget) Disconnect(handle int64) {
	pg_QCalendarWidget__disconnect(handle)
	unregisterCallback(handle)
}

func QCalendarWidget_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case 729690261457684101:
		return 1
	case 7049264599837532491:
		return 2
	case 1237236456312824991:
		return 3
	case 5217623514630858297:
		return 4
	case -332090898078543319:
		return 5
	case -4302917285655014661:
		return 6
	case 1023791977679793185:
		return 7
	case 8082959910766380751:
		return 8
	case -8659679893868524044:
		return 9
	case 2110182719375245575:
		return 10
	case 4311534060495737339:
		return 11
	case -8266762397995423915:
		return 12
	case 6398007404756346381:
		return 13
	case -4415920049297795213:
		return 14
	case -4029073905697023887:
		return 15
	case -4255494004099806581:
		return 16
	case 1919988340934511316:
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
	case 5290442163167458465:
		return 23
	case -9161949150489525396:
		return 24
	case -1989937050019464783:
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
	case 4959376063826165943:
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
	case 5898328517505784191:
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

func QCalendarWidget_enable(h uintptr, idx int32) { pg_QCalendarWidget__set_override(h, idx) }

func (this *QCalendarWidget) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QCalendarWidget) OnSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 729690261457684101, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sizeHint"
}

func (this *QCalendarWidget) OnMinimumSizeHint(fn func() *QSize) {
	registerVirt(this.UnsafeGoID(), 7049264599837532491, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "minimumSizeHint"
}

func (this *QCalendarWidget) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QCalendarWidget) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QCalendarWidget) OnMousePressEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -332090898078543319, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mousePressEvent"
}

func (this *QCalendarWidget) OnResizeEvent(fn func(*QResizeEvent)) {
	registerVirt(this.UnsafeGoID(), -4302917285655014661, func(event uintptr) { fn(UnsafeNewQResizeEvent(unsafe.Pointer(event))) }) // slot "resizeEvent"
}

func (this *QCalendarWidget) OnKeyPressEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 1023791977679793185, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyPressEvent"
}

func (this *QCalendarWidget) OnPaintCell(fn func(*QPainter, *QRect, *QDate)) {
	registerVirt(this.UnsafeGoID(), 8082959910766380751, func(painter uintptr, rect uintptr, date uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter)), UnsafeNewQRect(unsafe.Pointer(rect)), UnsafeNewQDate(unsafe.Pointer(date))) }) // slot "paintCell"
}

func (this *QCalendarWidget) OnDevType(fn func() int) {
	registerVirt(this.UnsafeGoID(), -8659679893868524044, func() int32 { return int32(fn()) }) // slot "devType"
}

func (this *QCalendarWidget) OnSetVisible(fn func(bool)) {
	registerVirt(this.UnsafeGoID(), 2110182719375245575, func(visible int32) { fn(visible != 0) }) // slot "setVisible"
}

func (this *QCalendarWidget) OnHeightForWidth(fn func(int) int) {
	registerVirt(this.UnsafeGoID(), 4311534060495737339, func(arg1 int32) int32 { return int32(fn(int(arg1))) }) // slot "heightForWidth"
}

func (this *QCalendarWidget) OnHasHeightForWidth(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -8266762397995423915, func() int32 { return b2i(fn()) }) // slot "hasHeightForWidth"
}

func (this *QCalendarWidget) OnPaintEngine(fn func() *QPaintEngine) {
	registerVirt(this.UnsafeGoID(), 6398007404756346381, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "paintEngine"
}

func (this *QCalendarWidget) OnMouseReleaseEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4415920049297795213, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseReleaseEvent"
}

func (this *QCalendarWidget) OnMouseDoubleClickEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4029073905697023887, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseDoubleClickEvent"
}

func (this *QCalendarWidget) OnMouseMoveEvent(fn func(*QMouseEvent)) {
	registerVirt(this.UnsafeGoID(), -4255494004099806581, func(event uintptr) { fn(UnsafeNewQMouseEvent(unsafe.Pointer(event))) }) // slot "mouseMoveEvent"
}

func (this *QCalendarWidget) OnWheelEvent(fn func(*QWheelEvent)) {
	registerVirt(this.UnsafeGoID(), 1919988340934511316, func(event uintptr) { fn(UnsafeNewQWheelEvent(unsafe.Pointer(event))) }) // slot "wheelEvent"
}

func (this *QCalendarWidget) OnKeyReleaseEvent(fn func(*QKeyEvent)) {
	registerVirt(this.UnsafeGoID(), 772598475950050299, func(event uintptr) { fn(UnsafeNewQKeyEvent(unsafe.Pointer(event))) }) // slot "keyReleaseEvent"
}

func (this *QCalendarWidget) OnFocusInEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), 975043767033552552, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusInEvent"
}

func (this *QCalendarWidget) OnFocusOutEvent(fn func(*QFocusEvent)) {
	registerVirt(this.UnsafeGoID(), -385596010576048993, func(event uintptr) { fn(UnsafeNewQFocusEvent(unsafe.Pointer(event))) }) // slot "focusOutEvent"
}

func (this *QCalendarWidget) OnEnterEvent(fn func(*QEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 6316270346668237895, func(event uintptr) { fn(UnsafeNewQEnterEvent(unsafe.Pointer(event))) }) // slot "enterEvent"
}

func (this *QCalendarWidget) OnLeaveEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), -551070009623711640, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "leaveEvent"
}

func (this *QCalendarWidget) OnPaintEvent(fn func(*QPaintEvent)) {
	registerVirt(this.UnsafeGoID(), 5290442163167458465, func(event uintptr) { fn(UnsafeNewQPaintEvent(unsafe.Pointer(event))) }) // slot "paintEvent"
}

func (this *QCalendarWidget) OnMoveEvent(fn func(*QMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -9161949150489525396, func(event uintptr) { fn(UnsafeNewQMoveEvent(unsafe.Pointer(event))) }) // slot "moveEvent"
}

func (this *QCalendarWidget) OnCloseEvent(fn func(*QCloseEvent)) {
	registerVirt(this.UnsafeGoID(), -1989937050019464783, func(event uintptr) { fn(UnsafeNewQCloseEvent(unsafe.Pointer(event))) }) // slot "closeEvent"
}

func (this *QCalendarWidget) OnContextMenuEvent(fn func(*QContextMenuEvent)) {
	registerVirt(this.UnsafeGoID(), -6698043222227006015, func(event uintptr) { fn(UnsafeNewQContextMenuEvent(unsafe.Pointer(event))) }) // slot "contextMenuEvent"
}

func (this *QCalendarWidget) OnTabletEvent(fn func(*QTabletEvent)) {
	registerVirt(this.UnsafeGoID(), -212068866860701141, func(event uintptr) { fn(UnsafeNewQTabletEvent(unsafe.Pointer(event))) }) // slot "tabletEvent"
}

func (this *QCalendarWidget) OnActionEvent(fn func(*QActionEvent)) {
	registerVirt(this.UnsafeGoID(), 1697444140088555581, func(event uintptr) { fn(UnsafeNewQActionEvent(unsafe.Pointer(event))) }) // slot "actionEvent"
}

func (this *QCalendarWidget) OnDragEnterEvent(fn func(*QDragEnterEvent)) {
	registerVirt(this.UnsafeGoID(), 206122715919917589, func(event uintptr) { fn(UnsafeNewQDragEnterEvent(unsafe.Pointer(event))) }) // slot "dragEnterEvent"
}

func (this *QCalendarWidget) OnDragMoveEvent(fn func(*QDragMoveEvent)) {
	registerVirt(this.UnsafeGoID(), -1781545137352745978, func(event uintptr) { fn(UnsafeNewQDragMoveEvent(unsafe.Pointer(event))) }) // slot "dragMoveEvent"
}

func (this *QCalendarWidget) OnDragLeaveEvent(fn func(*QDragLeaveEvent)) {
	registerVirt(this.UnsafeGoID(), -3139904158373725006, func(event uintptr) { fn(UnsafeNewQDragLeaveEvent(unsafe.Pointer(event))) }) // slot "dragLeaveEvent"
}

func (this *QCalendarWidget) OnDropEvent(fn func(*QDropEvent)) {
	registerVirt(this.UnsafeGoID(), 9107058682754591716, func(event uintptr) { fn(UnsafeNewQDropEvent(unsafe.Pointer(event))) }) // slot "dropEvent"
}

func (this *QCalendarWidget) OnShowEvent(fn func(*QShowEvent)) {
	registerVirt(this.UnsafeGoID(), -6825369174751916, func(event uintptr) { fn(UnsafeNewQShowEvent(unsafe.Pointer(event))) }) // slot "showEvent"
}

func (this *QCalendarWidget) OnHideEvent(fn func(*QHideEvent)) {
	registerVirt(this.UnsafeGoID(), -199718707876719129, func(event uintptr) { fn(UnsafeNewQHideEvent(unsafe.Pointer(event))) }) // slot "hideEvent"
}

func (this *QCalendarWidget) OnNativeEvent(fn func(*QByteArray, unsafe.Pointer, *int64) bool) {
	registerVirt(this.UnsafeGoID(), -2039166586246443706, func(eventType uintptr, message uintptr, result uintptr) int32 { return b2i(fn(UnsafeNewQByteArray(unsafe.Pointer(eventType)), unsafe.Pointer(message), (*int64)(unsafe.Pointer(result)))) }) // slot "nativeEvent"
}

func (this *QCalendarWidget) OnChangeEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 4959376063826165943, func(arg1 uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(arg1))) }) // slot "changeEvent"
}

func (this *QCalendarWidget) OnMetric(fn func(QPaintDevice__PaintDeviceMetric) int) {
	registerVirt(this.UnsafeGoID(), 427629649267311777, func(arg1 int32) int32 { return int32(fn(QPaintDevice__PaintDeviceMetric(arg1))) }) // slot "metric"
}

func (this *QCalendarWidget) OnInitPainter(fn func(*QPainter)) {
	registerVirt(this.UnsafeGoID(), 6477500296980050210, func(painter uintptr) { fn(UnsafeNewQPainter(unsafe.Pointer(painter))) }) // slot "initPainter"
}

func (this *QCalendarWidget) OnRedirected(fn func(*QPoint) *QPaintDevice) {
	registerVirt(this.UnsafeGoID(), 4127240633310655048, func(offset uintptr) uintptr { return uintptr(fn(UnsafeNewQPoint(unsafe.Pointer(offset))).UnsafePointer()) }) // slot "redirected"
}

func (this *QCalendarWidget) OnSharedPainter(fn func() *QPainter) {
	registerVirt(this.UnsafeGoID(), 5792771328741774991, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "sharedPainter"
}

func (this *QCalendarWidget) OnInputMethodEvent(fn func(*QInputMethodEvent)) {
	registerVirt(this.UnsafeGoID(), 5438819661942843222, func(arg1 uintptr) { fn(UnsafeNewQInputMethodEvent(unsafe.Pointer(arg1))) }) // slot "inputMethodEvent"
}

func (this *QCalendarWidget) OnInputMethodQuery(fn func(InputMethodQuery) *QVariant) {
	registerVirt(this.UnsafeGoID(), 7345853097340709638, func(arg1 int32) uintptr { return uintptr(fn(InputMethodQuery(arg1)).UnsafePointer()) }) // slot "inputMethodQuery"
}

func (this *QCalendarWidget) OnFocusNextPrevChild(fn func(bool) bool) {
	registerVirt(this.UnsafeGoID(), 5898328517505784191, func(next int32) int32 { return b2i(fn(next != 0)) }) // slot "focusNextPrevChild"
}

func (this *QCalendarWidget) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QCalendarWidget) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QCalendarWidget) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QCalendarWidget) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QCalendarWidget) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QCalendarWidget) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QCalendarWidget_delete(this.h)
	}
	this.UnsafeRelease()
}
