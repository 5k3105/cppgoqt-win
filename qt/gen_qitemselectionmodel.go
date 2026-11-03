package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QItemSelectionModel_base func(uintptr) uintptr
	addr_QItemSelectionModel_new uintptr
	addr_QItemSelectionModel_new2 uintptr
	addr_QItemSelectionModel_new3 uintptr
	addr_QItemSelectionModel_tr uintptr
	addr_QItemSelectionModel_currentIndex uintptr
	addr_QItemSelectionModel_isSelected uintptr
	addr_QItemSelectionModel_isRowSelected uintptr
	addr_QItemSelectionModel_isColumnSelected uintptr
	addr_QItemSelectionModel_rowIntersectsSelection uintptr
	addr_QItemSelectionModel_columnIntersectsSelection uintptr
	addr_QItemSelectionModel_hasSelection uintptr
	addr_QItemSelectionModel_selectedIndexes uintptr
	addr_QItemSelectionModel_selectedRows uintptr
	addr_QItemSelectionModel_selectedColumns uintptr
	addr_QItemSelectionModel_selection uintptr
	addr_QItemSelectionModel_model uintptr
	addr_QItemSelectionModel_model2 uintptr
	addr_QItemSelectionModel_setModel uintptr
	addr_QItemSelectionModel_setCurrentIndex uintptr
	addr_QItemSelectionModel_select uintptr
	addr_QItemSelectionModel_select2 uintptr
	addr_QItemSelectionModel_clear uintptr
	addr_QItemSelectionModel_reset uintptr
	addr_QItemSelectionModel_clearSelection uintptr
	addr_QItemSelectionModel_clearCurrentIndex uintptr
	addr_QItemSelectionModel_selectionChanged uintptr
	addr_QItemSelectionModel_currentChanged uintptr
	addr_QItemSelectionModel_currentRowChanged uintptr
	addr_QItemSelectionModel_currentColumnChanged uintptr
	addr_QItemSelectionModel_modelChanged uintptr
	addr_QItemSelectionModel_emitSelectionChanged uintptr
	addr_QItemSelectionModel_connect uintptr
	addr_QItemSelectionModel_connect2 uintptr
	addr_QItemSelectionModel_sender uintptr
	addr_QItemSelectionModel_senderSignalIndex uintptr
	addr_QItemSelectionModel_receivers uintptr
	addr_QItemSelectionModel_isSignalConnected uintptr
	addr_QItemSelectionModel_tr2 uintptr
	addr_QItemSelectionModel_tr3 uintptr
	addr_QItemSelectionModel_isRowSelected2 uintptr
	addr_QItemSelectionModel_isColumnSelected2 uintptr
	addr_QItemSelectionModel_rowIntersectsSelection2 uintptr
	addr_QItemSelectionModel_columnIntersectsSelection2 uintptr
	addr_QItemSelectionModel_selectedRows2 uintptr
	addr_QItemSelectionModel_selectedColumns2 uintptr
	addr_QItemSelectionModel_connect3 uintptr
	addr_QItemSelectionModel_connect4 uintptr
	addr_QItemSelectionModel_setCurrentIndex_default uintptr
	addr_QItemSelectionModel_select_default uintptr
	addr_QItemSelectionModel_select_default2 uintptr
	addr_QItemSelectionModel_clear_default uintptr
	addr_QItemSelectionModel_reset_default uintptr
	addr_QItemSelectionModel_clearCurrentIndex_default uintptr
	addr_QItemSelectionModel_metaObject_default uintptr
	addr_QItemSelectionModel_event_default uintptr
	addr_QItemSelectionModel_eventFilter_default uintptr
	addr_QItemSelectionModel_timerEvent_default uintptr
	addr_QItemSelectionModel_childEvent_default uintptr
	addr_QItemSelectionModel_customEvent_default uintptr
	addr_QItemSelectionModel_connectNotify_default uintptr
	addr_QItemSelectionModel_disconnectNotify_default uintptr
	pg_QItemSelectionModel_connect_selectionChanged func(uintptr, int64)
	pg_QItemSelectionModel_connect_currentChanged func(uintptr, int64)
	pg_QItemSelectionModel_connect_currentRowChanged func(uintptr, int64)
	pg_QItemSelectionModel_connect_currentColumnChanged func(uintptr, int64)
	pg_QItemSelectionModel_connect_modelChanged func(uintptr, int64)
	pg_QItemSelectionModel_connect_destroyed func(uintptr, int64)
	pg_QItemSelectionModel_connect_destroyed2 func(uintptr, int64)
	pg_QItemSelectionModel_connect_objectNameChanged func(uintptr, int64)
	pg_QItemSelectionModel__disconnect func(int64)
	pg_QItemSelectionModel__set_override func(uintptr, int32)
	pg_QItemSelectionModel_delete func(uintptr)
)

func init() {
	reg(&pg_QItemSelectionModel_base, "QItemSelectionModel_base")
	symaddr(&addr_QItemSelectionModel_new, "QItemSelectionModel_new")
	symaddr(&addr_QItemSelectionModel_new2, "QItemSelectionModel_new2")
	symaddr(&addr_QItemSelectionModel_new3, "QItemSelectionModel_new3")
	symaddr(&addr_QItemSelectionModel_tr, "QItemSelectionModel_tr")
	symaddr(&addr_QItemSelectionModel_currentIndex, "QItemSelectionModel_currentIndex")
	symaddr(&addr_QItemSelectionModel_isSelected, "QItemSelectionModel_isSelected")
	symaddr(&addr_QItemSelectionModel_isRowSelected, "QItemSelectionModel_isRowSelected")
	symaddr(&addr_QItemSelectionModel_isColumnSelected, "QItemSelectionModel_isColumnSelected")
	symaddr(&addr_QItemSelectionModel_rowIntersectsSelection, "QItemSelectionModel_rowIntersectsSelection")
	symaddr(&addr_QItemSelectionModel_columnIntersectsSelection, "QItemSelectionModel_columnIntersectsSelection")
	symaddr(&addr_QItemSelectionModel_hasSelection, "QItemSelectionModel_hasSelection")
	symaddr(&addr_QItemSelectionModel_selectedIndexes, "QItemSelectionModel_selectedIndexes")
	symaddr(&addr_QItemSelectionModel_selectedRows, "QItemSelectionModel_selectedRows")
	symaddr(&addr_QItemSelectionModel_selectedColumns, "QItemSelectionModel_selectedColumns")
	symaddr(&addr_QItemSelectionModel_selection, "QItemSelectionModel_selection")
	symaddr(&addr_QItemSelectionModel_model, "QItemSelectionModel_model")
	symaddr(&addr_QItemSelectionModel_model2, "QItemSelectionModel_model2")
	symaddr(&addr_QItemSelectionModel_setModel, "QItemSelectionModel_setModel")
	symaddr(&addr_QItemSelectionModel_setCurrentIndex, "QItemSelectionModel_setCurrentIndex")
	symaddr(&addr_QItemSelectionModel_select, "QItemSelectionModel_select")
	symaddr(&addr_QItemSelectionModel_select2, "QItemSelectionModel_select2")
	symaddr(&addr_QItemSelectionModel_clear, "QItemSelectionModel_clear")
	symaddr(&addr_QItemSelectionModel_reset, "QItemSelectionModel_reset")
	symaddr(&addr_QItemSelectionModel_clearSelection, "QItemSelectionModel_clearSelection")
	symaddr(&addr_QItemSelectionModel_clearCurrentIndex, "QItemSelectionModel_clearCurrentIndex")
	symaddr(&addr_QItemSelectionModel_selectionChanged, "QItemSelectionModel_selectionChanged")
	symaddr(&addr_QItemSelectionModel_currentChanged, "QItemSelectionModel_currentChanged")
	symaddr(&addr_QItemSelectionModel_currentRowChanged, "QItemSelectionModel_currentRowChanged")
	symaddr(&addr_QItemSelectionModel_currentColumnChanged, "QItemSelectionModel_currentColumnChanged")
	symaddr(&addr_QItemSelectionModel_modelChanged, "QItemSelectionModel_modelChanged")
	symaddr(&addr_QItemSelectionModel_emitSelectionChanged, "QItemSelectionModel_emitSelectionChanged")
	symaddr(&addr_QItemSelectionModel_connect, "QItemSelectionModel_connect")
	symaddr(&addr_QItemSelectionModel_connect2, "QItemSelectionModel_connect2")
	symaddr(&addr_QItemSelectionModel_sender, "QItemSelectionModel_sender")
	symaddr(&addr_QItemSelectionModel_senderSignalIndex, "QItemSelectionModel_senderSignalIndex")
	symaddr(&addr_QItemSelectionModel_receivers, "QItemSelectionModel_receivers")
	symaddr(&addr_QItemSelectionModel_isSignalConnected, "QItemSelectionModel_isSignalConnected")
	symaddr(&addr_QItemSelectionModel_tr2, "QItemSelectionModel_tr2")
	symaddr(&addr_QItemSelectionModel_tr3, "QItemSelectionModel_tr3")
	symaddr(&addr_QItemSelectionModel_isRowSelected2, "QItemSelectionModel_isRowSelected2")
	symaddr(&addr_QItemSelectionModel_isColumnSelected2, "QItemSelectionModel_isColumnSelected2")
	symaddr(&addr_QItemSelectionModel_rowIntersectsSelection2, "QItemSelectionModel_rowIntersectsSelection2")
	symaddr(&addr_QItemSelectionModel_columnIntersectsSelection2, "QItemSelectionModel_columnIntersectsSelection2")
	symaddr(&addr_QItemSelectionModel_selectedRows2, "QItemSelectionModel_selectedRows2")
	symaddr(&addr_QItemSelectionModel_selectedColumns2, "QItemSelectionModel_selectedColumns2")
	symaddr(&addr_QItemSelectionModel_connect3, "QItemSelectionModel_connect3")
	symaddr(&addr_QItemSelectionModel_connect4, "QItemSelectionModel_connect4")
	symaddr(&addr_QItemSelectionModel_setCurrentIndex_default, "QItemSelectionModel_setCurrentIndex_default")
	symaddr(&addr_QItemSelectionModel_select_default, "QItemSelectionModel_select_default")
	symaddr(&addr_QItemSelectionModel_select_default2, "QItemSelectionModel_select_default2")
	symaddr(&addr_QItemSelectionModel_clear_default, "QItemSelectionModel_clear_default")
	symaddr(&addr_QItemSelectionModel_reset_default, "QItemSelectionModel_reset_default")
	symaddr(&addr_QItemSelectionModel_clearCurrentIndex_default, "QItemSelectionModel_clearCurrentIndex_default")
	symaddr(&addr_QItemSelectionModel_metaObject_default, "QItemSelectionModel_metaObject_default")
	symaddr(&addr_QItemSelectionModel_event_default, "QItemSelectionModel_event_default")
	symaddr(&addr_QItemSelectionModel_eventFilter_default, "QItemSelectionModel_eventFilter_default")
	symaddr(&addr_QItemSelectionModel_timerEvent_default, "QItemSelectionModel_timerEvent_default")
	symaddr(&addr_QItemSelectionModel_childEvent_default, "QItemSelectionModel_childEvent_default")
	symaddr(&addr_QItemSelectionModel_customEvent_default, "QItemSelectionModel_customEvent_default")
	symaddr(&addr_QItemSelectionModel_connectNotify_default, "QItemSelectionModel_connectNotify_default")
	symaddr(&addr_QItemSelectionModel_disconnectNotify_default, "QItemSelectionModel_disconnectNotify_default")
	reg(&pg_QItemSelectionModel_connect_selectionChanged, "QItemSelectionModel_connect_selectionChanged")
	reg(&pg_QItemSelectionModel_connect_currentChanged, "QItemSelectionModel_connect_currentChanged")
	reg(&pg_QItemSelectionModel_connect_currentRowChanged, "QItemSelectionModel_connect_currentRowChanged")
	reg(&pg_QItemSelectionModel_connect_currentColumnChanged, "QItemSelectionModel_connect_currentColumnChanged")
	reg(&pg_QItemSelectionModel_connect_modelChanged, "QItemSelectionModel_connect_modelChanged")
	reg(&pg_QItemSelectionModel_connect_destroyed, "QItemSelectionModel_connect_destroyed")
	reg(&pg_QItemSelectionModel_connect_destroyed2, "QItemSelectionModel_connect_destroyed2")
	reg(&pg_QItemSelectionModel_connect_objectNameChanged, "QItemSelectionModel_connect_objectNameChanged")
	reg(&pg_QItemSelectionModel__disconnect, "QItemSelectionModel__disconnect")
	reg(&pg_QItemSelectionModel__set_override, "QItemSelectionModel__set_override")
	reg(&pg_QItemSelectionModel_delete, "QItemSelectionModel_delete")
}

type QItemSelectionModel struct {
	h uintptr
	*QObject
}

func bindQItemSelectionModel(h uintptr) *QItemSelectionModel {
	if h == 0 {
		return nil
	}
	return &QItemSelectionModel{h: h, QObject: bindQObject(pg_QItemSelectionModel_base(h))}
}

func newownedQItemSelectionModel(h uintptr) *QItemSelectionModel { return bindQItemSelectionModel(h) }

func (this *QItemSelectionModel) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QItemSelectionModel) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQItemSelectionModel(h unsafe.Pointer) *QItemSelectionModel { return bindQItemSelectionModel(uintptr(h)) }
func UnsafeNewQItemSelectionModelOwned(h unsafe.Pointer) *QItemSelectionModel { return newownedQItemSelectionModel(uintptr(h)) }

func (this *QItemSelectionModel) UnsafeRelease() {
	this.h = 0
	if this.QObject != nil {
		this.QObject.UnsafeRelease()
	}
}

func NewQItemSelectionModel() *QItemSelectionModel {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QItemSelectionModel_new, uintptr(__goid))
	__o := newownedQItemSelectionModel(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QItemSelectionModel_virtIdx, QItemSelectionModel_enable)
	return __o
}

func NewQItemSelectionModel2(model *QAbstractItemModel, parent *QObject) *QItemSelectionModel {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QItemSelectionModel_new2, uintptr(__goid), uintptr(model.UnsafePointer()), uintptr(parent.UnsafePointer()))
	__o := newownedQItemSelectionModel(__r)
	runtime.KeepAlive(model)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QItemSelectionModel_virtIdx, QItemSelectionModel_enable)
	return __o
}

func NewQItemSelectionModel3(model *QAbstractItemModel) *QItemSelectionModel {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QItemSelectionModel_new3, uintptr(__goid), uintptr(model.UnsafePointer()))
	__o := newownedQItemSelectionModel(__r)
	runtime.KeepAlive(model)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QItemSelectionModel_virtIdx, QItemSelectionModel_enable)
	return __o
}

func QItemSelectionModel_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QItemSelectionModel_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QItemSelectionModel) CurrentIndex() *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QItemSelectionModel_currentIndex, this.h)
	runtime.KeepAlive(this)
	return newownedQModelIndex(__r)
}

func (this *QItemSelectionModel) IsSelected(index *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QItemSelectionModel_isSelected, this.h, uintptr(index.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return __r != 0
}

func (this *QItemSelectionModel) IsRowSelected(row int) bool {
	__r0, _, _ := purego.SyscallN(addr_QItemSelectionModel_isRowSelected, this.h, uintptr(row))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QItemSelectionModel) IsColumnSelected(column int) bool {
	__r0, _, _ := purego.SyscallN(addr_QItemSelectionModel_isColumnSelected, this.h, uintptr(column))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QItemSelectionModel) RowIntersectsSelection(row int) bool {
	__r0, _, _ := purego.SyscallN(addr_QItemSelectionModel_rowIntersectsSelection, this.h, uintptr(row))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QItemSelectionModel) ColumnIntersectsSelection(column int) bool {
	__r0, _, _ := purego.SyscallN(addr_QItemSelectionModel_columnIntersectsSelection, this.h, uintptr(column))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QItemSelectionModel) HasSelection() bool {
	__r0, _, _ := purego.SyscallN(addr_QItemSelectionModel_hasSelection, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QItemSelectionModel) SelectedIndexes() []*QModelIndex {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QItemSelectionModel_selectedIndexes, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QModelIndex, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQModelIndex(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QItemSelectionModel) SelectedRows() []*QModelIndex {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QItemSelectionModel_selectedRows, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QModelIndex, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQModelIndex(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QItemSelectionModel) SelectedColumns() []*QModelIndex {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QItemSelectionModel_selectedColumns, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QModelIndex, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQModelIndex(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QItemSelectionModel) Selection() *QItemSelection {
	__r, _, _ := purego.SyscallN(addr_QItemSelectionModel_selection, this.h)
	runtime.KeepAlive(this)
	return newownedQItemSelection(__r)
}

func (this *QItemSelectionModel) Model() *QAbstractItemModel {
	__r, _, _ := purego.SyscallN(addr_QItemSelectionModel_model, this.h)
	runtime.KeepAlive(this)
	return bindQAbstractItemModel(__r)
}

func (this *QItemSelectionModel) Model2() *QAbstractItemModel {
	__r, _, _ := purego.SyscallN(addr_QItemSelectionModel_model2, this.h)
	runtime.KeepAlive(this)
	return bindQAbstractItemModel(__r)
}

func (this *QItemSelectionModel) SetModel(model *QAbstractItemModel) {
	purego.SyscallN(addr_QItemSelectionModel_setModel, this.h, uintptr(model.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(model)
}

func (this *QItemSelectionModel) SetCurrentIndex(index *QModelIndex, command QItemSelectionModel__SelectionFlag) {
	purego.SyscallN(addr_QItemSelectionModel_setCurrentIndex, this.h, uintptr(index.UnsafePointer()), uintptr(command))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
}

func (this *QItemSelectionModel) Select(index *QModelIndex, command QItemSelectionModel__SelectionFlag) {
	purego.SyscallN(addr_QItemSelectionModel_select, this.h, uintptr(index.UnsafePointer()), uintptr(command))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
}

func (this *QItemSelectionModel) Select2(selection *QItemSelection, command QItemSelectionModel__SelectionFlag) {
	purego.SyscallN(addr_QItemSelectionModel_select2, this.h, uintptr(selection.UnsafePointer()), uintptr(command))
	runtime.KeepAlive(this)
	runtime.KeepAlive(selection)
}

func (this *QItemSelectionModel) Clear() {
	purego.SyscallN(addr_QItemSelectionModel_clear, this.h)
	runtime.KeepAlive(this)
}

func (this *QItemSelectionModel) Reset() {
	purego.SyscallN(addr_QItemSelectionModel_reset, this.h)
	runtime.KeepAlive(this)
}

func (this *QItemSelectionModel) ClearSelection() {
	purego.SyscallN(addr_QItemSelectionModel_clearSelection, this.h)
	runtime.KeepAlive(this)
}

func (this *QItemSelectionModel) ClearCurrentIndex() {
	purego.SyscallN(addr_QItemSelectionModel_clearCurrentIndex, this.h)
	runtime.KeepAlive(this)
}

func (this *QItemSelectionModel) SelectionChanged(selected *QItemSelection, deselected *QItemSelection) {
	purego.SyscallN(addr_QItemSelectionModel_selectionChanged, this.h, uintptr(selected.UnsafePointer()), uintptr(deselected.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(selected)
	runtime.KeepAlive(deselected)
}

func (this *QItemSelectionModel) CurrentChanged(current *QModelIndex, previous *QModelIndex) {
	purego.SyscallN(addr_QItemSelectionModel_currentChanged, this.h, uintptr(current.UnsafePointer()), uintptr(previous.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(current)
	runtime.KeepAlive(previous)
}

func (this *QItemSelectionModel) CurrentRowChanged(current *QModelIndex, previous *QModelIndex) {
	purego.SyscallN(addr_QItemSelectionModel_currentRowChanged, this.h, uintptr(current.UnsafePointer()), uintptr(previous.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(current)
	runtime.KeepAlive(previous)
}

func (this *QItemSelectionModel) CurrentColumnChanged(current *QModelIndex, previous *QModelIndex) {
	purego.SyscallN(addr_QItemSelectionModel_currentColumnChanged, this.h, uintptr(current.UnsafePointer()), uintptr(previous.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(current)
	runtime.KeepAlive(previous)
}

func (this *QItemSelectionModel) ModelChanged(model *QAbstractItemModel) {
	purego.SyscallN(addr_QItemSelectionModel_modelChanged, this.h, uintptr(model.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(model)
}

func (this *QItemSelectionModel) EmitSelectionChanged(newSelection *QItemSelection, oldSelection *QItemSelection) {
	purego.SyscallN(addr_QItemSelectionModel_emitSelectionChanged, this.h, uintptr(newSelection.UnsafePointer()), uintptr(oldSelection.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(newSelection)
	runtime.KeepAlive(oldSelection)
}

func QItemSelectionModel_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QItemSelectionModel_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QItemSelectionModel_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QItemSelectionModel_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QItemSelectionModel) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QItemSelectionModel_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QItemSelectionModel) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QItemSelectionModel_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QItemSelectionModel) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QItemSelectionModel_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QItemSelectionModel) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QItemSelectionModel_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QItemSelectionModel_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QItemSelectionModel_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QItemSelectionModel_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QItemSelectionModel_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func (this *QItemSelectionModel) IsRowSelected2(row int, parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QItemSelectionModel_isRowSelected2, this.h, uintptr(row), uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QItemSelectionModel) IsColumnSelected2(column int, parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QItemSelectionModel_isColumnSelected2, this.h, uintptr(column), uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QItemSelectionModel) RowIntersectsSelection2(row int, parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QItemSelectionModel_rowIntersectsSelection2, this.h, uintptr(row), uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QItemSelectionModel) ColumnIntersectsSelection2(column int, parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QItemSelectionModel_columnIntersectsSelection2, this.h, uintptr(column), uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QItemSelectionModel) SelectedRowsWithColumn(column int) []*QModelIndex {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QItemSelectionModel_selectedRows2, this.h, uintptr(column), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QModelIndex, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQModelIndex(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QItemSelectionModel) SelectedColumnsWithRow(row int) []*QModelIndex {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QItemSelectionModel_selectedColumns2, this.h, uintptr(row), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QModelIndex, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQModelIndex(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QItemSelectionModel_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QItemSelectionModel_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QItemSelectionModel_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QItemSelectionModel_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QItemSelectionModel) SetCurrentIndexDefault(index *QModelIndex, command QItemSelectionModel__SelectionFlag) {
	purego.SyscallN(addr_QItemSelectionModel_setCurrentIndex_default, this.h, uintptr(index.UnsafePointer()), uintptr(command))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
}

func (this *QItemSelectionModel) SelectDefault(index *QModelIndex, command QItemSelectionModel__SelectionFlag) {
	purego.SyscallN(addr_QItemSelectionModel_select_default, this.h, uintptr(index.UnsafePointer()), uintptr(command))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
}

func (this *QItemSelectionModel) Select2Default(selection *QItemSelection, command QItemSelectionModel__SelectionFlag) {
	purego.SyscallN(addr_QItemSelectionModel_select_default2, this.h, uintptr(selection.UnsafePointer()), uintptr(command))
	runtime.KeepAlive(this)
	runtime.KeepAlive(selection)
}

func (this *QItemSelectionModel) ClearDefault() {
	purego.SyscallN(addr_QItemSelectionModel_clear_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QItemSelectionModel) ResetDefault() {
	purego.SyscallN(addr_QItemSelectionModel_reset_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QItemSelectionModel) ClearCurrentIndexDefault() {
	purego.SyscallN(addr_QItemSelectionModel_clearCurrentIndex_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QItemSelectionModel) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QItemSelectionModel_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QItemSelectionModel) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QItemSelectionModel_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QItemSelectionModel) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QItemSelectionModel_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QItemSelectionModel) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QItemSelectionModel_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QItemSelectionModel) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QItemSelectionModel_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QItemSelectionModel) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QItemSelectionModel_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QItemSelectionModel) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QItemSelectionModel_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QItemSelectionModel) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QItemSelectionModel_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QItemSelectionModel) OnSelectionChanged(fn func(*QItemSelection, *QItemSelection)) int64 {
	id := registerCallback(func(selected uintptr, deselected uintptr) { fn(UnsafeNewQItemSelection(unsafe.Pointer(selected)), UnsafeNewQItemSelection(unsafe.Pointer(deselected))) })
	pg_QItemSelectionModel_connect_selectionChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QItemSelectionModel) OnCurrentChanged(fn func(*QModelIndex, *QModelIndex)) int64 {
	id := registerCallback(func(current uintptr, previous uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(current)), UnsafeNewQModelIndex(unsafe.Pointer(previous))) })
	pg_QItemSelectionModel_connect_currentChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QItemSelectionModel) OnCurrentRowChanged(fn func(*QModelIndex, *QModelIndex)) int64 {
	id := registerCallback(func(current uintptr, previous uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(current)), UnsafeNewQModelIndex(unsafe.Pointer(previous))) })
	pg_QItemSelectionModel_connect_currentRowChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QItemSelectionModel) OnCurrentColumnChanged(fn func(*QModelIndex, *QModelIndex)) int64 {
	id := registerCallback(func(current uintptr, previous uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(current)), UnsafeNewQModelIndex(unsafe.Pointer(previous))) })
	pg_QItemSelectionModel_connect_currentColumnChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QItemSelectionModel) OnModelChanged(fn func(*QAbstractItemModel)) int64 {
	id := registerCallback(func(model uintptr) { fn(UnsafeNewQAbstractItemModel(unsafe.Pointer(model))) })
	pg_QItemSelectionModel_connect_modelChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QItemSelectionModel) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QItemSelectionModel_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QItemSelectionModel) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QItemSelectionModel_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QItemSelectionModel) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QItemSelectionModel_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QItemSelectionModel) Disconnect(handle int64) {
	pg_QItemSelectionModel__disconnect(handle)
	unregisterCallback(handle)
}

func QItemSelectionModel_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case -6491574503686613496:
		return 1
	case -4794627110275867635:
		return 2
	case 8951329890086256397:
		return 3
	case -778567940572639070:
		return 4
	case 1086335023529244512:
		return 5
	case 9143260233085619705:
		return 6
	case 1237236456312824991:
		return 7
	case 5217623514630858297:
		return 8
	case -5409927581812374498:
		return 9
	case -1185315367526857833:
		return 10
	case 3769868229123868474:
		return 11
	case 7379029982170503008:
		return 12
	case -568369772982215962:
		return 13
	}
	return -1
}

func QItemSelectionModel_enable(h uintptr, idx int32) { pg_QItemSelectionModel__set_override(h, idx) }

func (this *QItemSelectionModel) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QItemSelectionModel) OnSetCurrentIndex(fn func(*QModelIndex, QItemSelectionModel__SelectionFlag)) {
	registerVirt(this.UnsafeGoID(), -6491574503686613496, func(index uintptr, command int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(index)), QItemSelectionModel__SelectionFlag(command)) }) // slot "setCurrentIndex"
}

func (this *QItemSelectionModel) OnSelect(fn func(*QModelIndex, QItemSelectionModel__SelectionFlag)) {
	registerVirt(this.UnsafeGoID(), -4794627110275867635, func(index uintptr, command int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(index)), QItemSelectionModel__SelectionFlag(command)) }) // slot "select"
}

func (this *QItemSelectionModel) OnSelect2(fn func(*QItemSelection, QItemSelectionModel__SelectionFlag)) {
	registerVirt(this.UnsafeGoID(), 8951329890086256397, func(selection uintptr, command int32) { fn(UnsafeNewQItemSelection(unsafe.Pointer(selection)), QItemSelectionModel__SelectionFlag(command)) }) // slot "select2"
}

func (this *QItemSelectionModel) OnClear(fn func()) {
	registerVirt(this.UnsafeGoID(), -778567940572639070, func() { fn() }) // slot "clear"
}

func (this *QItemSelectionModel) OnReset(fn func()) {
	registerVirt(this.UnsafeGoID(), 1086335023529244512, func() { fn() }) // slot "reset"
}

func (this *QItemSelectionModel) OnClearCurrentIndex(fn func()) {
	registerVirt(this.UnsafeGoID(), 9143260233085619705, func() { fn() }) // slot "clearCurrentIndex"
}

func (this *QItemSelectionModel) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QItemSelectionModel) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QItemSelectionModel) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QItemSelectionModel) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QItemSelectionModel) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QItemSelectionModel) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QItemSelectionModel) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QItemSelectionModel) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QItemSelectionModel_delete(this.h)
	}
	this.UnsafeRelease()
}
