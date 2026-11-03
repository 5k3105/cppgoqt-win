package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QHeaderView_chain uintptr
	addr_QHeaderView_new uintptr
	addr_QHeaderView_new2 uintptr
	addr_QHeaderView_tr uintptr
	addr_QHeaderView_orientation uintptr
	addr_QHeaderView_offset uintptr
	addr_QHeaderView_length uintptr
	addr_QHeaderView_sectionSizeHint uintptr
	addr_QHeaderView_visualIndexAt uintptr
	addr_QHeaderView_logicalIndexAt uintptr
	addr_QHeaderView_logicalIndexAt2 uintptr
	addr_QHeaderView_logicalIndexAt3 uintptr
	addr_QHeaderView_sectionSize uintptr
	addr_QHeaderView_sectionPosition uintptr
	addr_QHeaderView_sectionViewportPosition uintptr
	addr_QHeaderView_moveSection uintptr
	addr_QHeaderView_swapSections uintptr
	addr_QHeaderView_resizeSection uintptr
	addr_QHeaderView_resizeSections uintptr
	addr_QHeaderView_isSectionHidden uintptr
	addr_QHeaderView_setSectionHidden uintptr
	addr_QHeaderView_hiddenSectionCount uintptr
	addr_QHeaderView_hideSection uintptr
	addr_QHeaderView_showSection uintptr
	addr_QHeaderView_count uintptr
	addr_QHeaderView_visualIndex uintptr
	addr_QHeaderView_logicalIndex uintptr
	addr_QHeaderView_setSectionsMovable uintptr
	addr_QHeaderView_sectionsMovable uintptr
	addr_QHeaderView_setFirstSectionMovable uintptr
	addr_QHeaderView_isFirstSectionMovable uintptr
	addr_QHeaderView_setSectionsClickable uintptr
	addr_QHeaderView_sectionsClickable uintptr
	addr_QHeaderView_setHighlightSections uintptr
	addr_QHeaderView_highlightSections uintptr
	addr_QHeaderView_sectionResizeMode uintptr
	addr_QHeaderView_setSectionResizeMode uintptr
	addr_QHeaderView_setSectionResizeMode2 uintptr
	addr_QHeaderView_setResizeContentsPrecision uintptr
	addr_QHeaderView_resizeContentsPrecision uintptr
	addr_QHeaderView_stretchSectionCount uintptr
	addr_QHeaderView_setSortIndicatorShown uintptr
	addr_QHeaderView_isSortIndicatorShown uintptr
	addr_QHeaderView_setSortIndicator uintptr
	addr_QHeaderView_sortIndicatorSection uintptr
	addr_QHeaderView_sortIndicatorOrder uintptr
	addr_QHeaderView_setSortIndicatorClearable uintptr
	addr_QHeaderView_isSortIndicatorClearable uintptr
	addr_QHeaderView_stretchLastSection uintptr
	addr_QHeaderView_setStretchLastSection uintptr
	addr_QHeaderView_cascadingSectionResizes uintptr
	addr_QHeaderView_setCascadingSectionResizes uintptr
	addr_QHeaderView_defaultSectionSize uintptr
	addr_QHeaderView_setDefaultSectionSize uintptr
	addr_QHeaderView_resetDefaultSectionSize uintptr
	addr_QHeaderView_minimumSectionSize uintptr
	addr_QHeaderView_setMinimumSectionSize uintptr
	addr_QHeaderView_maximumSectionSize uintptr
	addr_QHeaderView_setMaximumSectionSize uintptr
	addr_QHeaderView_defaultAlignment uintptr
	addr_QHeaderView_setDefaultAlignment uintptr
	addr_QHeaderView_sectionsMoved uintptr
	addr_QHeaderView_sectionsHidden uintptr
	addr_QHeaderView_saveState uintptr
	addr_QHeaderView_restoreState uintptr
	addr_QHeaderView_setOffset uintptr
	addr_QHeaderView_setOffsetToSectionPosition uintptr
	addr_QHeaderView_setOffsetToLastSection uintptr
	addr_QHeaderView_headerDataChanged uintptr
	addr_QHeaderView_sectionMoved uintptr
	addr_QHeaderView_sectionResized uintptr
	addr_QHeaderView_sectionPressed uintptr
	addr_QHeaderView_sectionClicked uintptr
	addr_QHeaderView_sectionEntered uintptr
	addr_QHeaderView_sectionDoubleClicked uintptr
	addr_QHeaderView_sectionCountChanged uintptr
	addr_QHeaderView_sectionHandleDoubleClicked uintptr
	addr_QHeaderView_geometriesChanged uintptr
	addr_QHeaderView_sortIndicatorChanged uintptr
	addr_QHeaderView_sortIndicatorClearableChanged uintptr
	addr_QHeaderView_setTabOrder uintptr
	addr_QHeaderView_mouseGrabber uintptr
	addr_QHeaderView_keyboardGrabber uintptr
	addr_QHeaderView_find uintptr
	addr_QHeaderView_createWindowContainer uintptr
	addr_QHeaderView_connect uintptr
	addr_QHeaderView_connect2 uintptr
	pg_QHeaderView_devicePixelRatioFScale func() float64
	pg_QHeaderView_encodeMetricF func(int32, float64) int32
	addr_QHeaderView_tr2 uintptr
	addr_QHeaderView_tr3 uintptr
	addr_QHeaderView_createWindowContainer2 uintptr
	addr_QHeaderView_createWindowContainer3 uintptr
	addr_QHeaderView_connect3 uintptr
	addr_QHeaderView_connect4 uintptr
	pg_QHeaderView_connect_sectionMoved func(uintptr, int64)
	pg_QHeaderView_connect_sectionResized func(uintptr, int64)
	pg_QHeaderView_connect_sectionPressed func(uintptr, int64)
	pg_QHeaderView_connect_sectionClicked func(uintptr, int64)
	pg_QHeaderView_connect_sectionEntered func(uintptr, int64)
	pg_QHeaderView_connect_sectionDoubleClicked func(uintptr, int64)
	pg_QHeaderView_connect_sectionCountChanged func(uintptr, int64)
	pg_QHeaderView_connect_sectionHandleDoubleClicked func(uintptr, int64)
	pg_QHeaderView_connect_geometriesChanged func(uintptr, int64)
	pg_QHeaderView_connect_sortIndicatorChanged func(uintptr, int64)
	pg_QHeaderView_connect_sortIndicatorClearableChanged func(uintptr, int64)
	pg_QHeaderView_connect_pressed func(uintptr, int64)
	pg_QHeaderView_connect_clicked func(uintptr, int64)
	pg_QHeaderView_connect_doubleClicked func(uintptr, int64)
	pg_QHeaderView_connect_activated func(uintptr, int64)
	pg_QHeaderView_connect_entered func(uintptr, int64)
	pg_QHeaderView_connect_viewportEntered func(uintptr, int64)
	pg_QHeaderView_connect_iconSizeChanged func(uintptr, int64)
	pg_QHeaderView_connect_windowTitleChanged func(uintptr, int64)
	pg_QHeaderView_connect_windowIconChanged func(uintptr, int64)
	pg_QHeaderView_connect_windowIconTextChanged func(uintptr, int64)
	pg_QHeaderView_connect_customContextMenuRequested func(uintptr, int64)
	pg_QHeaderView_connect_destroyed func(uintptr, int64)
	pg_QHeaderView_connect_destroyed2 func(uintptr, int64)
	pg_QHeaderView_connect_objectNameChanged func(uintptr, int64)
	pg_QHeaderView__disconnect func(int64)
	pg_QHeaderView_delete func(uintptr)
)

func init() {
	symaddr(&addr_QHeaderView_chain, "QHeaderView_chain")
	symaddr(&addr_QHeaderView_new, "QHeaderView_new")
	symaddr(&addr_QHeaderView_new2, "QHeaderView_new2")
	symaddr(&addr_QHeaderView_tr, "QHeaderView_tr")
	symaddr(&addr_QHeaderView_orientation, "QHeaderView_orientation")
	symaddr(&addr_QHeaderView_offset, "QHeaderView_offset")
	symaddr(&addr_QHeaderView_length, "QHeaderView_length")
	symaddr(&addr_QHeaderView_sectionSizeHint, "QHeaderView_sectionSizeHint")
	symaddr(&addr_QHeaderView_visualIndexAt, "QHeaderView_visualIndexAt")
	symaddr(&addr_QHeaderView_logicalIndexAt, "QHeaderView_logicalIndexAt")
	symaddr(&addr_QHeaderView_logicalIndexAt2, "QHeaderView_logicalIndexAt2")
	symaddr(&addr_QHeaderView_logicalIndexAt3, "QHeaderView_logicalIndexAt3")
	symaddr(&addr_QHeaderView_sectionSize, "QHeaderView_sectionSize")
	symaddr(&addr_QHeaderView_sectionPosition, "QHeaderView_sectionPosition")
	symaddr(&addr_QHeaderView_sectionViewportPosition, "QHeaderView_sectionViewportPosition")
	symaddr(&addr_QHeaderView_moveSection, "QHeaderView_moveSection")
	symaddr(&addr_QHeaderView_swapSections, "QHeaderView_swapSections")
	symaddr(&addr_QHeaderView_resizeSection, "QHeaderView_resizeSection")
	symaddr(&addr_QHeaderView_resizeSections, "QHeaderView_resizeSections")
	symaddr(&addr_QHeaderView_isSectionHidden, "QHeaderView_isSectionHidden")
	symaddr(&addr_QHeaderView_setSectionHidden, "QHeaderView_setSectionHidden")
	symaddr(&addr_QHeaderView_hiddenSectionCount, "QHeaderView_hiddenSectionCount")
	symaddr(&addr_QHeaderView_hideSection, "QHeaderView_hideSection")
	symaddr(&addr_QHeaderView_showSection, "QHeaderView_showSection")
	symaddr(&addr_QHeaderView_count, "QHeaderView_count")
	symaddr(&addr_QHeaderView_visualIndex, "QHeaderView_visualIndex")
	symaddr(&addr_QHeaderView_logicalIndex, "QHeaderView_logicalIndex")
	symaddr(&addr_QHeaderView_setSectionsMovable, "QHeaderView_setSectionsMovable")
	symaddr(&addr_QHeaderView_sectionsMovable, "QHeaderView_sectionsMovable")
	symaddr(&addr_QHeaderView_setFirstSectionMovable, "QHeaderView_setFirstSectionMovable")
	symaddr(&addr_QHeaderView_isFirstSectionMovable, "QHeaderView_isFirstSectionMovable")
	symaddr(&addr_QHeaderView_setSectionsClickable, "QHeaderView_setSectionsClickable")
	symaddr(&addr_QHeaderView_sectionsClickable, "QHeaderView_sectionsClickable")
	symaddr(&addr_QHeaderView_setHighlightSections, "QHeaderView_setHighlightSections")
	symaddr(&addr_QHeaderView_highlightSections, "QHeaderView_highlightSections")
	symaddr(&addr_QHeaderView_sectionResizeMode, "QHeaderView_sectionResizeMode")
	symaddr(&addr_QHeaderView_setSectionResizeMode, "QHeaderView_setSectionResizeMode")
	symaddr(&addr_QHeaderView_setSectionResizeMode2, "QHeaderView_setSectionResizeMode2")
	symaddr(&addr_QHeaderView_setResizeContentsPrecision, "QHeaderView_setResizeContentsPrecision")
	symaddr(&addr_QHeaderView_resizeContentsPrecision, "QHeaderView_resizeContentsPrecision")
	symaddr(&addr_QHeaderView_stretchSectionCount, "QHeaderView_stretchSectionCount")
	symaddr(&addr_QHeaderView_setSortIndicatorShown, "QHeaderView_setSortIndicatorShown")
	symaddr(&addr_QHeaderView_isSortIndicatorShown, "QHeaderView_isSortIndicatorShown")
	symaddr(&addr_QHeaderView_setSortIndicator, "QHeaderView_setSortIndicator")
	symaddr(&addr_QHeaderView_sortIndicatorSection, "QHeaderView_sortIndicatorSection")
	symaddr(&addr_QHeaderView_sortIndicatorOrder, "QHeaderView_sortIndicatorOrder")
	symaddr(&addr_QHeaderView_setSortIndicatorClearable, "QHeaderView_setSortIndicatorClearable")
	symaddr(&addr_QHeaderView_isSortIndicatorClearable, "QHeaderView_isSortIndicatorClearable")
	symaddr(&addr_QHeaderView_stretchLastSection, "QHeaderView_stretchLastSection")
	symaddr(&addr_QHeaderView_setStretchLastSection, "QHeaderView_setStretchLastSection")
	symaddr(&addr_QHeaderView_cascadingSectionResizes, "QHeaderView_cascadingSectionResizes")
	symaddr(&addr_QHeaderView_setCascadingSectionResizes, "QHeaderView_setCascadingSectionResizes")
	symaddr(&addr_QHeaderView_defaultSectionSize, "QHeaderView_defaultSectionSize")
	symaddr(&addr_QHeaderView_setDefaultSectionSize, "QHeaderView_setDefaultSectionSize")
	symaddr(&addr_QHeaderView_resetDefaultSectionSize, "QHeaderView_resetDefaultSectionSize")
	symaddr(&addr_QHeaderView_minimumSectionSize, "QHeaderView_minimumSectionSize")
	symaddr(&addr_QHeaderView_setMinimumSectionSize, "QHeaderView_setMinimumSectionSize")
	symaddr(&addr_QHeaderView_maximumSectionSize, "QHeaderView_maximumSectionSize")
	symaddr(&addr_QHeaderView_setMaximumSectionSize, "QHeaderView_setMaximumSectionSize")
	symaddr(&addr_QHeaderView_defaultAlignment, "QHeaderView_defaultAlignment")
	symaddr(&addr_QHeaderView_setDefaultAlignment, "QHeaderView_setDefaultAlignment")
	symaddr(&addr_QHeaderView_sectionsMoved, "QHeaderView_sectionsMoved")
	symaddr(&addr_QHeaderView_sectionsHidden, "QHeaderView_sectionsHidden")
	symaddr(&addr_QHeaderView_saveState, "QHeaderView_saveState")
	symaddr(&addr_QHeaderView_restoreState, "QHeaderView_restoreState")
	symaddr(&addr_QHeaderView_setOffset, "QHeaderView_setOffset")
	symaddr(&addr_QHeaderView_setOffsetToSectionPosition, "QHeaderView_setOffsetToSectionPosition")
	symaddr(&addr_QHeaderView_setOffsetToLastSection, "QHeaderView_setOffsetToLastSection")
	symaddr(&addr_QHeaderView_headerDataChanged, "QHeaderView_headerDataChanged")
	symaddr(&addr_QHeaderView_sectionMoved, "QHeaderView_sectionMoved")
	symaddr(&addr_QHeaderView_sectionResized, "QHeaderView_sectionResized")
	symaddr(&addr_QHeaderView_sectionPressed, "QHeaderView_sectionPressed")
	symaddr(&addr_QHeaderView_sectionClicked, "QHeaderView_sectionClicked")
	symaddr(&addr_QHeaderView_sectionEntered, "QHeaderView_sectionEntered")
	symaddr(&addr_QHeaderView_sectionDoubleClicked, "QHeaderView_sectionDoubleClicked")
	symaddr(&addr_QHeaderView_sectionCountChanged, "QHeaderView_sectionCountChanged")
	symaddr(&addr_QHeaderView_sectionHandleDoubleClicked, "QHeaderView_sectionHandleDoubleClicked")
	symaddr(&addr_QHeaderView_geometriesChanged, "QHeaderView_geometriesChanged")
	symaddr(&addr_QHeaderView_sortIndicatorChanged, "QHeaderView_sortIndicatorChanged")
	symaddr(&addr_QHeaderView_sortIndicatorClearableChanged, "QHeaderView_sortIndicatorClearableChanged")
	symaddr(&addr_QHeaderView_setTabOrder, "QHeaderView_setTabOrder")
	symaddr(&addr_QHeaderView_mouseGrabber, "QHeaderView_mouseGrabber")
	symaddr(&addr_QHeaderView_keyboardGrabber, "QHeaderView_keyboardGrabber")
	symaddr(&addr_QHeaderView_find, "QHeaderView_find")
	symaddr(&addr_QHeaderView_createWindowContainer, "QHeaderView_createWindowContainer")
	symaddr(&addr_QHeaderView_connect, "QHeaderView_connect")
	symaddr(&addr_QHeaderView_connect2, "QHeaderView_connect2")
	reg(&pg_QHeaderView_devicePixelRatioFScale, "QHeaderView_devicePixelRatioFScale")
	reg(&pg_QHeaderView_encodeMetricF, "QHeaderView_encodeMetricF")
	symaddr(&addr_QHeaderView_tr2, "QHeaderView_tr2")
	symaddr(&addr_QHeaderView_tr3, "QHeaderView_tr3")
	symaddr(&addr_QHeaderView_createWindowContainer2, "QHeaderView_createWindowContainer2")
	symaddr(&addr_QHeaderView_createWindowContainer3, "QHeaderView_createWindowContainer3")
	symaddr(&addr_QHeaderView_connect3, "QHeaderView_connect3")
	symaddr(&addr_QHeaderView_connect4, "QHeaderView_connect4")
	reg(&pg_QHeaderView_connect_sectionMoved, "QHeaderView_connect_sectionMoved")
	reg(&pg_QHeaderView_connect_sectionResized, "QHeaderView_connect_sectionResized")
	reg(&pg_QHeaderView_connect_sectionPressed, "QHeaderView_connect_sectionPressed")
	reg(&pg_QHeaderView_connect_sectionClicked, "QHeaderView_connect_sectionClicked")
	reg(&pg_QHeaderView_connect_sectionEntered, "QHeaderView_connect_sectionEntered")
	reg(&pg_QHeaderView_connect_sectionDoubleClicked, "QHeaderView_connect_sectionDoubleClicked")
	reg(&pg_QHeaderView_connect_sectionCountChanged, "QHeaderView_connect_sectionCountChanged")
	reg(&pg_QHeaderView_connect_sectionHandleDoubleClicked, "QHeaderView_connect_sectionHandleDoubleClicked")
	reg(&pg_QHeaderView_connect_geometriesChanged, "QHeaderView_connect_geometriesChanged")
	reg(&pg_QHeaderView_connect_sortIndicatorChanged, "QHeaderView_connect_sortIndicatorChanged")
	reg(&pg_QHeaderView_connect_sortIndicatorClearableChanged, "QHeaderView_connect_sortIndicatorClearableChanged")
	reg(&pg_QHeaderView_connect_pressed, "QHeaderView_connect_pressed")
	reg(&pg_QHeaderView_connect_clicked, "QHeaderView_connect_clicked")
	reg(&pg_QHeaderView_connect_doubleClicked, "QHeaderView_connect_doubleClicked")
	reg(&pg_QHeaderView_connect_activated, "QHeaderView_connect_activated")
	reg(&pg_QHeaderView_connect_entered, "QHeaderView_connect_entered")
	reg(&pg_QHeaderView_connect_viewportEntered, "QHeaderView_connect_viewportEntered")
	reg(&pg_QHeaderView_connect_iconSizeChanged, "QHeaderView_connect_iconSizeChanged")
	reg(&pg_QHeaderView_connect_windowTitleChanged, "QHeaderView_connect_windowTitleChanged")
	reg(&pg_QHeaderView_connect_windowIconChanged, "QHeaderView_connect_windowIconChanged")
	reg(&pg_QHeaderView_connect_windowIconTextChanged, "QHeaderView_connect_windowIconTextChanged")
	reg(&pg_QHeaderView_connect_customContextMenuRequested, "QHeaderView_connect_customContextMenuRequested")
	reg(&pg_QHeaderView_connect_destroyed, "QHeaderView_connect_destroyed")
	reg(&pg_QHeaderView_connect_destroyed2, "QHeaderView_connect_destroyed2")
	reg(&pg_QHeaderView_connect_objectNameChanged, "QHeaderView_connect_objectNameChanged")
	reg(&pg_QHeaderView__disconnect, "QHeaderView__disconnect")
	reg(&pg_QHeaderView_delete, "QHeaderView_delete")
}

type QHeaderView struct {
	h uintptr
	*QAbstractItemView
}

func bindQHeaderView(h uintptr) *QHeaderView {
	if h == 0 {
		return nil
	}
	var __c [5]uintptr
	purego.SyscallN(addr_QHeaderView_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QHeaderView{h: h, QAbstractItemView: &QAbstractItemView{h: __c[0], QAbstractScrollArea: &QAbstractScrollArea{h: __c[1], QFrame: &QFrame{h: __c[2], QWidget: &QWidget{h: __c[3], QObject: &QObject{h: __c[4]}}}}}}
}

func newownedQHeaderView(h uintptr) *QHeaderView { return bindQHeaderView(h) }

func (this *QHeaderView) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QHeaderView) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQHeaderView(h unsafe.Pointer) *QHeaderView { return bindQHeaderView(uintptr(h)) }
func UnsafeNewQHeaderViewOwned(h unsafe.Pointer) *QHeaderView { return newownedQHeaderView(uintptr(h)) }

func (this *QHeaderView) UnsafeRelease() {
	this.h = 0
	if this.QAbstractItemView != nil {
		this.QAbstractItemView.UnsafeRelease()
	}
}

func NewQHeaderView(orientation Orientation) *QHeaderView {
	__r, _, _ := purego.SyscallN(addr_QHeaderView_new, uintptr(orientation))
	__o := newownedQHeaderView(__r)
	return __o
}

func NewQHeaderView2(orientation Orientation, parent *QWidget) *QHeaderView {
	__r, _, _ := purego.SyscallN(addr_QHeaderView_new2, uintptr(orientation), uintptr(parent.UnsafePointer()))
	__o := newownedQHeaderView(__r)
	runtime.KeepAlive(parent)
	return __o
}

func QHeaderView_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QHeaderView_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QHeaderView) Orientation() Orientation {
	__r0, _, _ := purego.SyscallN(addr_QHeaderView_orientation, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return Orientation(__r)
}

func (this *QHeaderView) Offset() int {
	__r0, _, _ := purego.SyscallN(addr_QHeaderView_offset, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QHeaderView) Length() int {
	__r0, _, _ := purego.SyscallN(addr_QHeaderView_length, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QHeaderView) SectionSizeHint(logicalIndex int) int {
	__r0, _, _ := purego.SyscallN(addr_QHeaderView_sectionSizeHint, this.h, uintptr(logicalIndex))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QHeaderView) VisualIndexAt(position int) int {
	__r0, _, _ := purego.SyscallN(addr_QHeaderView_visualIndexAt, this.h, uintptr(position))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QHeaderView) LogicalIndexAt(position int) int {
	__r0, _, _ := purego.SyscallN(addr_QHeaderView_logicalIndexAt, this.h, uintptr(position))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QHeaderView) LogicalIndexAt2(x int, y int) int {
	__r0, _, _ := purego.SyscallN(addr_QHeaderView_logicalIndexAt2, this.h, uintptr(x), uintptr(y))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QHeaderView) LogicalIndexAtWithPos(pos *QPoint) int {
	__r0, _, _ := purego.SyscallN(addr_QHeaderView_logicalIndexAt3, this.h, uintptr(pos.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pos)
	return int(__r)
}

func (this *QHeaderView) SectionSize(logicalIndex int) int {
	__r0, _, _ := purego.SyscallN(addr_QHeaderView_sectionSize, this.h, uintptr(logicalIndex))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QHeaderView) SectionPosition(logicalIndex int) int {
	__r0, _, _ := purego.SyscallN(addr_QHeaderView_sectionPosition, this.h, uintptr(logicalIndex))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QHeaderView) SectionViewportPosition(logicalIndex int) int {
	__r0, _, _ := purego.SyscallN(addr_QHeaderView_sectionViewportPosition, this.h, uintptr(logicalIndex))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QHeaderView) MoveSection(from int, to int) {
	purego.SyscallN(addr_QHeaderView_moveSection, this.h, uintptr(from), uintptr(to))
	runtime.KeepAlive(this)
}

func (this *QHeaderView) SwapSections(first int, second int) {
	purego.SyscallN(addr_QHeaderView_swapSections, this.h, uintptr(first), uintptr(second))
	runtime.KeepAlive(this)
}

func (this *QHeaderView) ResizeSection(logicalIndex int, size int) {
	purego.SyscallN(addr_QHeaderView_resizeSection, this.h, uintptr(logicalIndex), uintptr(size))
	runtime.KeepAlive(this)
}

func (this *QHeaderView) ResizeSections(mode QHeaderView__ResizeMode) {
	purego.SyscallN(addr_QHeaderView_resizeSections, this.h, uintptr(mode))
	runtime.KeepAlive(this)
}

func (this *QHeaderView) IsSectionHidden(logicalIndex int) bool {
	__r0, _, _ := purego.SyscallN(addr_QHeaderView_isSectionHidden, this.h, uintptr(logicalIndex))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QHeaderView) SetSectionHidden(logicalIndex int, hide bool) {
	purego.SyscallN(addr_QHeaderView_setSectionHidden, this.h, uintptr(logicalIndex), uintptr(b2i(hide)))
	runtime.KeepAlive(this)
}

func (this *QHeaderView) HiddenSectionCount() int {
	__r0, _, _ := purego.SyscallN(addr_QHeaderView_hiddenSectionCount, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QHeaderView) HideSection(logicalIndex int) {
	purego.SyscallN(addr_QHeaderView_hideSection, this.h, uintptr(logicalIndex))
	runtime.KeepAlive(this)
}

func (this *QHeaderView) ShowSection(logicalIndex int) {
	purego.SyscallN(addr_QHeaderView_showSection, this.h, uintptr(logicalIndex))
	runtime.KeepAlive(this)
}

func (this *QHeaderView) Count() int {
	__r0, _, _ := purego.SyscallN(addr_QHeaderView_count, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QHeaderView) VisualIndex(logicalIndex int) int {
	__r0, _, _ := purego.SyscallN(addr_QHeaderView_visualIndex, this.h, uintptr(logicalIndex))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QHeaderView) LogicalIndex(visualIndex int) int {
	__r0, _, _ := purego.SyscallN(addr_QHeaderView_logicalIndex, this.h, uintptr(visualIndex))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QHeaderView) SetSectionsMovable(movable bool) {
	purego.SyscallN(addr_QHeaderView_setSectionsMovable, this.h, uintptr(b2i(movable)))
	runtime.KeepAlive(this)
}

func (this *QHeaderView) SectionsMovable() bool {
	__r0, _, _ := purego.SyscallN(addr_QHeaderView_sectionsMovable, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QHeaderView) SetFirstSectionMovable(movable bool) {
	purego.SyscallN(addr_QHeaderView_setFirstSectionMovable, this.h, uintptr(b2i(movable)))
	runtime.KeepAlive(this)
}

func (this *QHeaderView) IsFirstSectionMovable() bool {
	__r0, _, _ := purego.SyscallN(addr_QHeaderView_isFirstSectionMovable, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QHeaderView) SetSectionsClickable(clickable bool) {
	purego.SyscallN(addr_QHeaderView_setSectionsClickable, this.h, uintptr(b2i(clickable)))
	runtime.KeepAlive(this)
}

func (this *QHeaderView) SectionsClickable() bool {
	__r0, _, _ := purego.SyscallN(addr_QHeaderView_sectionsClickable, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QHeaderView) SetHighlightSections(highlight bool) {
	purego.SyscallN(addr_QHeaderView_setHighlightSections, this.h, uintptr(b2i(highlight)))
	runtime.KeepAlive(this)
}

func (this *QHeaderView) HighlightSections() bool {
	__r0, _, _ := purego.SyscallN(addr_QHeaderView_highlightSections, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QHeaderView) SectionResizeMode(logicalIndex int) QHeaderView__ResizeMode {
	__r0, _, _ := purego.SyscallN(addr_QHeaderView_sectionResizeMode, this.h, uintptr(logicalIndex))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QHeaderView__ResizeMode(__r)
}

func (this *QHeaderView) SetSectionResizeMode(mode QHeaderView__ResizeMode) {
	purego.SyscallN(addr_QHeaderView_setSectionResizeMode, this.h, uintptr(mode))
	runtime.KeepAlive(this)
}

func (this *QHeaderView) SetSectionResizeMode2(logicalIndex int, mode QHeaderView__ResizeMode) {
	purego.SyscallN(addr_QHeaderView_setSectionResizeMode2, this.h, uintptr(logicalIndex), uintptr(mode))
	runtime.KeepAlive(this)
}

func (this *QHeaderView) SetResizeContentsPrecision(precision int) {
	purego.SyscallN(addr_QHeaderView_setResizeContentsPrecision, this.h, uintptr(precision))
	runtime.KeepAlive(this)
}

func (this *QHeaderView) ResizeContentsPrecision() int {
	__r0, _, _ := purego.SyscallN(addr_QHeaderView_resizeContentsPrecision, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QHeaderView) StretchSectionCount() int {
	__r0, _, _ := purego.SyscallN(addr_QHeaderView_stretchSectionCount, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QHeaderView) SetSortIndicatorShown(show bool) {
	purego.SyscallN(addr_QHeaderView_setSortIndicatorShown, this.h, uintptr(b2i(show)))
	runtime.KeepAlive(this)
}

func (this *QHeaderView) IsSortIndicatorShown() bool {
	__r0, _, _ := purego.SyscallN(addr_QHeaderView_isSortIndicatorShown, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QHeaderView) SetSortIndicator(logicalIndex int, order SortOrder) {
	purego.SyscallN(addr_QHeaderView_setSortIndicator, this.h, uintptr(logicalIndex), uintptr(order))
	runtime.KeepAlive(this)
}

func (this *QHeaderView) SortIndicatorSection() int {
	__r0, _, _ := purego.SyscallN(addr_QHeaderView_sortIndicatorSection, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QHeaderView) SortIndicatorOrder() SortOrder {
	__r0, _, _ := purego.SyscallN(addr_QHeaderView_sortIndicatorOrder, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return SortOrder(__r)
}

func (this *QHeaderView) SetSortIndicatorClearable(clearable bool) {
	purego.SyscallN(addr_QHeaderView_setSortIndicatorClearable, this.h, uintptr(b2i(clearable)))
	runtime.KeepAlive(this)
}

func (this *QHeaderView) IsSortIndicatorClearable() bool {
	__r0, _, _ := purego.SyscallN(addr_QHeaderView_isSortIndicatorClearable, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QHeaderView) StretchLastSection() bool {
	__r0, _, _ := purego.SyscallN(addr_QHeaderView_stretchLastSection, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QHeaderView) SetStretchLastSection(stretch bool) {
	purego.SyscallN(addr_QHeaderView_setStretchLastSection, this.h, uintptr(b2i(stretch)))
	runtime.KeepAlive(this)
}

func (this *QHeaderView) CascadingSectionResizes() bool {
	__r0, _, _ := purego.SyscallN(addr_QHeaderView_cascadingSectionResizes, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QHeaderView) SetCascadingSectionResizes(enable bool) {
	purego.SyscallN(addr_QHeaderView_setCascadingSectionResizes, this.h, uintptr(b2i(enable)))
	runtime.KeepAlive(this)
}

func (this *QHeaderView) DefaultSectionSize() int {
	__r0, _, _ := purego.SyscallN(addr_QHeaderView_defaultSectionSize, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QHeaderView) SetDefaultSectionSize(size int) {
	purego.SyscallN(addr_QHeaderView_setDefaultSectionSize, this.h, uintptr(size))
	runtime.KeepAlive(this)
}

func (this *QHeaderView) ResetDefaultSectionSize() {
	purego.SyscallN(addr_QHeaderView_resetDefaultSectionSize, this.h)
	runtime.KeepAlive(this)
}

func (this *QHeaderView) MinimumSectionSize() int {
	__r0, _, _ := purego.SyscallN(addr_QHeaderView_minimumSectionSize, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QHeaderView) SetMinimumSectionSize(size int) {
	purego.SyscallN(addr_QHeaderView_setMinimumSectionSize, this.h, uintptr(size))
	runtime.KeepAlive(this)
}

func (this *QHeaderView) MaximumSectionSize() int {
	__r0, _, _ := purego.SyscallN(addr_QHeaderView_maximumSectionSize, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QHeaderView) SetMaximumSectionSize(size int) {
	purego.SyscallN(addr_QHeaderView_setMaximumSectionSize, this.h, uintptr(size))
	runtime.KeepAlive(this)
}

func (this *QHeaderView) DefaultAlignment() AlignmentFlag {
	__r0, _, _ := purego.SyscallN(addr_QHeaderView_defaultAlignment, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return AlignmentFlag(__r)
}

func (this *QHeaderView) SetDefaultAlignment(alignment AlignmentFlag) {
	purego.SyscallN(addr_QHeaderView_setDefaultAlignment, this.h, uintptr(alignment))
	runtime.KeepAlive(this)
}

func (this *QHeaderView) SectionsMoved() bool {
	__r0, _, _ := purego.SyscallN(addr_QHeaderView_sectionsMoved, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QHeaderView) SectionsHidden() bool {
	__r0, _, _ := purego.SyscallN(addr_QHeaderView_sectionsHidden, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QHeaderView) SaveState() *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QHeaderView_saveState, this.h)
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QHeaderView) RestoreState(state *QByteArray) bool {
	__r0, _, _ := purego.SyscallN(addr_QHeaderView_restoreState, this.h, uintptr(state.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(state)
	return __r != 0
}

func (this *QHeaderView) SetOffset(offset int) {
	purego.SyscallN(addr_QHeaderView_setOffset, this.h, uintptr(offset))
	runtime.KeepAlive(this)
}

func (this *QHeaderView) SetOffsetToSectionPosition(visualIndex int) {
	purego.SyscallN(addr_QHeaderView_setOffsetToSectionPosition, this.h, uintptr(visualIndex))
	runtime.KeepAlive(this)
}

func (this *QHeaderView) SetOffsetToLastSection() {
	purego.SyscallN(addr_QHeaderView_setOffsetToLastSection, this.h)
	runtime.KeepAlive(this)
}

func (this *QHeaderView) HeaderDataChanged(orientation Orientation, logicalFirst int, logicalLast int) {
	purego.SyscallN(addr_QHeaderView_headerDataChanged, this.h, uintptr(orientation), uintptr(logicalFirst), uintptr(logicalLast))
	runtime.KeepAlive(this)
}

func (this *QHeaderView) SectionMoved(logicalIndex int, oldVisualIndex int, newVisualIndex int) {
	purego.SyscallN(addr_QHeaderView_sectionMoved, this.h, uintptr(logicalIndex), uintptr(oldVisualIndex), uintptr(newVisualIndex))
	runtime.KeepAlive(this)
}

func (this *QHeaderView) SectionResized(logicalIndex int, oldSize int, newSize int) {
	purego.SyscallN(addr_QHeaderView_sectionResized, this.h, uintptr(logicalIndex), uintptr(oldSize), uintptr(newSize))
	runtime.KeepAlive(this)
}

func (this *QHeaderView) SectionPressed(logicalIndex int) {
	purego.SyscallN(addr_QHeaderView_sectionPressed, this.h, uintptr(logicalIndex))
	runtime.KeepAlive(this)
}

func (this *QHeaderView) SectionClicked(logicalIndex int) {
	purego.SyscallN(addr_QHeaderView_sectionClicked, this.h, uintptr(logicalIndex))
	runtime.KeepAlive(this)
}

func (this *QHeaderView) SectionEntered(logicalIndex int) {
	purego.SyscallN(addr_QHeaderView_sectionEntered, this.h, uintptr(logicalIndex))
	runtime.KeepAlive(this)
}

func (this *QHeaderView) SectionDoubleClicked(logicalIndex int) {
	purego.SyscallN(addr_QHeaderView_sectionDoubleClicked, this.h, uintptr(logicalIndex))
	runtime.KeepAlive(this)
}

func (this *QHeaderView) SectionCountChanged(oldCount int, newCount int) {
	purego.SyscallN(addr_QHeaderView_sectionCountChanged, this.h, uintptr(oldCount), uintptr(newCount))
	runtime.KeepAlive(this)
}

func (this *QHeaderView) SectionHandleDoubleClicked(logicalIndex int) {
	purego.SyscallN(addr_QHeaderView_sectionHandleDoubleClicked, this.h, uintptr(logicalIndex))
	runtime.KeepAlive(this)
}

func (this *QHeaderView) GeometriesChanged() {
	purego.SyscallN(addr_QHeaderView_geometriesChanged, this.h)
	runtime.KeepAlive(this)
}

func (this *QHeaderView) SortIndicatorChanged(logicalIndex int, order SortOrder) {
	purego.SyscallN(addr_QHeaderView_sortIndicatorChanged, this.h, uintptr(logicalIndex), uintptr(order))
	runtime.KeepAlive(this)
}

func (this *QHeaderView) SortIndicatorClearableChanged(clearable bool) {
	purego.SyscallN(addr_QHeaderView_sortIndicatorClearableChanged, this.h, uintptr(b2i(clearable)))
	runtime.KeepAlive(this)
}

func QHeaderView_SetTabOrder(arg1 *QWidget, arg2 *QWidget) {
	purego.SyscallN(addr_QHeaderView_setTabOrder, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()))
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
}

func QHeaderView_MouseGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QHeaderView_mouseGrabber)
	return bindQWidget(__r)
}

func QHeaderView_KeyboardGrabber() *QWidget {
	__r, _, _ := purego.SyscallN(addr_QHeaderView_keyboardGrabber)
	return bindQWidget(__r)
}

func QHeaderView_Find(arg1 uint64) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QHeaderView_find, uintptr(arg1))
	return bindQWidget(__r)
}

func QHeaderView_CreateWindowContainer(window *QWindow) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QHeaderView_createWindowContainer, uintptr(window.UnsafePointer()))
	runtime.KeepAlive(window)
	return bindQWidget(__r)
}

func QHeaderView_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QHeaderView_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QHeaderView_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QHeaderView_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func QHeaderView_DevicePixelRatioFScale() float64 {
	__r := pg_QHeaderView_devicePixelRatioFScale()
	return __r
}

func QHeaderView_EncodeMetricF(metric QPaintDevice__PaintDeviceMetric, value float64) int {
	__r := pg_QHeaderView_encodeMetricF(int32(metric), value)
	return int(__r)
}

func QHeaderView_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QHeaderView_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QHeaderView_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QHeaderView_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QHeaderView_CreateWindowContainer2(window *QWindow, parent *QWidget) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QHeaderView_createWindowContainer2, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func QHeaderView_CreateWindowContainer3(window *QWindow, parent *QWidget, flags WindowType) *QWidget {
	__r, _, _ := purego.SyscallN(addr_QHeaderView_createWindowContainer3, uintptr(window.UnsafePointer()), uintptr(parent.UnsafePointer()), uintptr(flags))
	runtime.KeepAlive(window)
	runtime.KeepAlive(parent)
	return bindQWidget(__r)
}

func QHeaderView_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QHeaderView_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QHeaderView_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QHeaderView_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QHeaderView) OnSectionMoved(fn func(int, int, int)) int64 {
	id := registerCallback(func(logicalIndex int32, oldVisualIndex int32, newVisualIndex int32) { fn(int(logicalIndex), int(oldVisualIndex), int(newVisualIndex)) })
	pg_QHeaderView_connect_sectionMoved(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QHeaderView) OnSectionResized(fn func(int, int, int)) int64 {
	id := registerCallback(func(logicalIndex int32, oldSize int32, newSize int32) { fn(int(logicalIndex), int(oldSize), int(newSize)) })
	pg_QHeaderView_connect_sectionResized(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QHeaderView) OnSectionPressed(fn func(int)) int64 {
	id := registerCallback(func(logicalIndex int32) { fn(int(logicalIndex)) })
	pg_QHeaderView_connect_sectionPressed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QHeaderView) OnSectionClicked(fn func(int)) int64 {
	id := registerCallback(func(logicalIndex int32) { fn(int(logicalIndex)) })
	pg_QHeaderView_connect_sectionClicked(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QHeaderView) OnSectionEntered(fn func(int)) int64 {
	id := registerCallback(func(logicalIndex int32) { fn(int(logicalIndex)) })
	pg_QHeaderView_connect_sectionEntered(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QHeaderView) OnSectionDoubleClicked(fn func(int)) int64 {
	id := registerCallback(func(logicalIndex int32) { fn(int(logicalIndex)) })
	pg_QHeaderView_connect_sectionDoubleClicked(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QHeaderView) OnSectionCountChanged(fn func(int, int)) int64 {
	id := registerCallback(func(oldCount int32, newCount int32) { fn(int(oldCount), int(newCount)) })
	pg_QHeaderView_connect_sectionCountChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QHeaderView) OnSectionHandleDoubleClicked(fn func(int)) int64 {
	id := registerCallback(func(logicalIndex int32) { fn(int(logicalIndex)) })
	pg_QHeaderView_connect_sectionHandleDoubleClicked(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QHeaderView) OnGeometriesChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QHeaderView_connect_geometriesChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QHeaderView) OnSortIndicatorChanged(fn func(int, SortOrder)) int64 {
	id := registerCallback(func(logicalIndex int32, order int32) { fn(int(logicalIndex), SortOrder(order)) })
	pg_QHeaderView_connect_sortIndicatorChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QHeaderView) OnSortIndicatorClearableChanged(fn func(bool)) int64 {
	id := registerCallback(func(clearable int32) { fn(clearable != 0) })
	pg_QHeaderView_connect_sortIndicatorClearableChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QHeaderView) OnPressed(fn func(*QModelIndex)) int64 {
	id := registerCallback(func(index uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(index))) })
	pg_QHeaderView_connect_pressed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QHeaderView) OnClicked(fn func(*QModelIndex)) int64 {
	id := registerCallback(func(index uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(index))) })
	pg_QHeaderView_connect_clicked(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QHeaderView) OnDoubleClicked(fn func(*QModelIndex)) int64 {
	id := registerCallback(func(index uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(index))) })
	pg_QHeaderView_connect_doubleClicked(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QHeaderView) OnActivated(fn func(*QModelIndex)) int64 {
	id := registerCallback(func(index uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(index))) })
	pg_QHeaderView_connect_activated(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QHeaderView) OnEntered(fn func(*QModelIndex)) int64 {
	id := registerCallback(func(index uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(index))) })
	pg_QHeaderView_connect_entered(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QHeaderView) OnViewportEntered(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QHeaderView_connect_viewportEntered(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QHeaderView) OnIconSizeChanged(fn func(*QSize)) int64 {
	id := registerCallback(func(size uintptr) { fn(UnsafeNewQSize(unsafe.Pointer(size))) })
	pg_QHeaderView_connect_iconSizeChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QHeaderView) OnWindowTitleChanged(fn func(string)) int64 {
	id := registerCallback(func(title uintptr, title_len int64) { fn(pgStrNoFree(title, title_len)) })
	pg_QHeaderView_connect_windowTitleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QHeaderView) OnWindowIconChanged(fn func(*QIcon)) int64 {
	id := registerCallback(func(icon uintptr) { fn(UnsafeNewQIcon(unsafe.Pointer(icon))) })
	pg_QHeaderView_connect_windowIconChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QHeaderView) OnWindowIconTextChanged(fn func(string)) int64 {
	id := registerCallback(func(iconText uintptr, iconText_len int64) { fn(pgStrNoFree(iconText, iconText_len)) })
	pg_QHeaderView_connect_windowIconTextChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QHeaderView) OnCustomContextMenuRequested(fn func(*QPoint)) int64 {
	id := registerCallback(func(pos uintptr) { fn(UnsafeNewQPoint(unsafe.Pointer(pos))) })
	pg_QHeaderView_connect_customContextMenuRequested(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QHeaderView) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QHeaderView_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QHeaderView) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QHeaderView_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QHeaderView) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QHeaderView_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QHeaderView) Disconnect(handle int64) {
	pg_QHeaderView__disconnect(handle)
	unregisterCallback(handle)
}

func (this *QHeaderView) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QHeaderView_delete(this.h)
	}
	this.UnsafeRelease()
}
