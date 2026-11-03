package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

// purego: unsupported shapes skipped (map returns, >15-arg / >9-arg-callback, exotic): SetItemRoleNames

var (
	addr_QStandardItemModel_chain uintptr
	addr_QStandardItemModel_new uintptr
	addr_QStandardItemModel_new2 uintptr
	addr_QStandardItemModel_new3 uintptr
	addr_QStandardItemModel_new4 uintptr
	addr_QStandardItemModel_tr uintptr
	addr_QStandardItemModel_clear uintptr
	addr_QStandardItemModel_itemFromIndex uintptr
	addr_QStandardItemModel_indexFromItem uintptr
	addr_QStandardItemModel_item uintptr
	addr_QStandardItemModel_setItem uintptr
	addr_QStandardItemModel_setItem2 uintptr
	addr_QStandardItemModel_invisibleRootItem uintptr
	addr_QStandardItemModel_horizontalHeaderItem uintptr
	addr_QStandardItemModel_setHorizontalHeaderItem uintptr
	addr_QStandardItemModel_verticalHeaderItem uintptr
	addr_QStandardItemModel_setVerticalHeaderItem uintptr
	addr_QStandardItemModel_setHorizontalHeaderLabels uintptr
	addr_QStandardItemModel_setVerticalHeaderLabels uintptr
	addr_QStandardItemModel_setRowCount uintptr
	addr_QStandardItemModel_setColumnCount uintptr
	addr_QStandardItemModel_appendRow uintptr
	addr_QStandardItemModel_appendColumn uintptr
	addr_QStandardItemModel_appendRow2 uintptr
	addr_QStandardItemModel_insertRow uintptr
	addr_QStandardItemModel_insertColumn uintptr
	addr_QStandardItemModel_insertRow2 uintptr
	addr_QStandardItemModel_insertRow3 uintptr
	addr_QStandardItemModel_insertColumn2 uintptr
	addr_QStandardItemModel_takeItem uintptr
	addr_QStandardItemModel_takeRow uintptr
	addr_QStandardItemModel_takeColumn uintptr
	addr_QStandardItemModel_takeHorizontalHeaderItem uintptr
	addr_QStandardItemModel_takeVerticalHeaderItem uintptr
	addr_QStandardItemModel_itemPrototype uintptr
	addr_QStandardItemModel_setItemPrototype uintptr
	addr_QStandardItemModel_findItems uintptr
	addr_QStandardItemModel_sortRole uintptr
	addr_QStandardItemModel_setSortRole uintptr
	addr_QStandardItemModel_itemChanged uintptr
	addr_QStandardItemModel_createIndex uintptr
	addr_QStandardItemModel_createIndex2 uintptr
	addr_QStandardItemModel_encodeData uintptr
	addr_QStandardItemModel_decodeData uintptr
	addr_QStandardItemModel_beginInsertRows uintptr
	addr_QStandardItemModel_endInsertRows uintptr
	addr_QStandardItemModel_beginRemoveRows uintptr
	addr_QStandardItemModel_endRemoveRows uintptr
	addr_QStandardItemModel_beginMoveRows uintptr
	addr_QStandardItemModel_endMoveRows uintptr
	addr_QStandardItemModel_beginInsertColumns uintptr
	addr_QStandardItemModel_endInsertColumns uintptr
	addr_QStandardItemModel_beginRemoveColumns uintptr
	addr_QStandardItemModel_endRemoveColumns uintptr
	addr_QStandardItemModel_beginMoveColumns uintptr
	addr_QStandardItemModel_endMoveColumns uintptr
	addr_QStandardItemModel_beginResetModel uintptr
	addr_QStandardItemModel_endResetModel uintptr
	addr_QStandardItemModel_changePersistentIndex uintptr
	addr_QStandardItemModel_changePersistentIndexList uintptr
	addr_QStandardItemModel_persistentIndexList uintptr
	addr_QStandardItemModel_connect uintptr
	addr_QStandardItemModel_connect2 uintptr
	addr_QStandardItemModel_sender uintptr
	addr_QStandardItemModel_senderSignalIndex uintptr
	addr_QStandardItemModel_receivers uintptr
	addr_QStandardItemModel_isSignalConnected uintptr
	addr_QStandardItemModel_tr2 uintptr
	addr_QStandardItemModel_tr3 uintptr
	addr_QStandardItemModel_item2 uintptr
	addr_QStandardItemModel_insertRow4 uintptr
	addr_QStandardItemModel_insertColumn3 uintptr
	addr_QStandardItemModel_takeItem2 uintptr
	addr_QStandardItemModel_findItems2 uintptr
	addr_QStandardItemModel_findItems3 uintptr
	addr_QStandardItemModel_connect3 uintptr
	addr_QStandardItemModel_connect4 uintptr
	addr_QStandardItemModel_metaObject_default uintptr
	addr_QStandardItemModel_index_default uintptr
	addr_QStandardItemModel_parent_default uintptr
	addr_QStandardItemModel_rowCount_default uintptr
	addr_QStandardItemModel_columnCount_default uintptr
	addr_QStandardItemModel_hasChildren_default uintptr
	addr_QStandardItemModel_data_default uintptr
	addr_QStandardItemModel_multiData_default uintptr
	addr_QStandardItemModel_setData_default uintptr
	addr_QStandardItemModel_clearItemData_default uintptr
	addr_QStandardItemModel_headerData_default uintptr
	addr_QStandardItemModel_setHeaderData_default uintptr
	addr_QStandardItemModel_insertRows_default uintptr
	addr_QStandardItemModel_insertColumns_default uintptr
	addr_QStandardItemModel_removeRows_default uintptr
	addr_QStandardItemModel_removeColumns_default uintptr
	addr_QStandardItemModel_flags_default uintptr
	addr_QStandardItemModel_supportedDropActions_default uintptr
	addr_QStandardItemModel_sort_default uintptr
	addr_QStandardItemModel_mimeTypes_default uintptr
	addr_QStandardItemModel_mimeData_default uintptr
	addr_QStandardItemModel_dropMimeData_default uintptr
	addr_QStandardItemModel_sibling_default uintptr
	addr_QStandardItemModel_canDropMimeData_default uintptr
	addr_QStandardItemModel_supportedDragActions_default uintptr
	addr_QStandardItemModel_moveRows_default uintptr
	addr_QStandardItemModel_moveColumns_default uintptr
	addr_QStandardItemModel_fetchMore_default uintptr
	addr_QStandardItemModel_canFetchMore_default uintptr
	addr_QStandardItemModel_buddy_default uintptr
	addr_QStandardItemModel_match_default uintptr
	addr_QStandardItemModel_span_default uintptr
	addr_QStandardItemModel_submit_default uintptr
	addr_QStandardItemModel_revert_default uintptr
	addr_QStandardItemModel_resetInternalData_default uintptr
	addr_QStandardItemModel_event_default uintptr
	addr_QStandardItemModel_eventFilter_default uintptr
	addr_QStandardItemModel_timerEvent_default uintptr
	addr_QStandardItemModel_childEvent_default uintptr
	addr_QStandardItemModel_customEvent_default uintptr
	addr_QStandardItemModel_connectNotify_default uintptr
	addr_QStandardItemModel_disconnectNotify_default uintptr
	pg_QStandardItemModel_connect_itemChanged func(uintptr, int64)
	pg_QStandardItemModel_connect_headerDataChanged func(uintptr, int64)
	pg_QStandardItemModel_connect_layoutChanged func(uintptr, int64)
	pg_QStandardItemModel_connect_layoutChanged2 func(uintptr, int64)
	pg_QStandardItemModel_connect_layoutAboutToBeChanged func(uintptr, int64)
	pg_QStandardItemModel_connect_layoutAboutToBeChanged2 func(uintptr, int64)
	pg_QStandardItemModel_connect_rowsAboutToBeInserted func(uintptr, int64)
	pg_QStandardItemModel_connect_rowsInserted func(uintptr, int64)
	pg_QStandardItemModel_connect_rowsAboutToBeRemoved func(uintptr, int64)
	pg_QStandardItemModel_connect_rowsRemoved func(uintptr, int64)
	pg_QStandardItemModel_connect_columnsAboutToBeInserted func(uintptr, int64)
	pg_QStandardItemModel_connect_columnsInserted func(uintptr, int64)
	pg_QStandardItemModel_connect_columnsAboutToBeRemoved func(uintptr, int64)
	pg_QStandardItemModel_connect_columnsRemoved func(uintptr, int64)
	pg_QStandardItemModel_connect_modelAboutToBeReset func(uintptr, int64)
	pg_QStandardItemModel_connect_modelReset func(uintptr, int64)
	pg_QStandardItemModel_connect_rowsAboutToBeMoved func(uintptr, int64)
	pg_QStandardItemModel_connect_rowsMoved func(uintptr, int64)
	pg_QStandardItemModel_connect_columnsAboutToBeMoved func(uintptr, int64)
	pg_QStandardItemModel_connect_columnsMoved func(uintptr, int64)
	pg_QStandardItemModel_connect_destroyed func(uintptr, int64)
	pg_QStandardItemModel_connect_destroyed2 func(uintptr, int64)
	pg_QStandardItemModel_connect_objectNameChanged func(uintptr, int64)
	pg_QStandardItemModel__disconnect func(int64)
	pg_QStandardItemModel__set_override func(uintptr, int32)
	pg_QStandardItemModel_delete func(uintptr)
)

func init() {
	symaddr(&addr_QStandardItemModel_chain, "QStandardItemModel_chain")
	symaddr(&addr_QStandardItemModel_new, "QStandardItemModel_new")
	symaddr(&addr_QStandardItemModel_new2, "QStandardItemModel_new2")
	symaddr(&addr_QStandardItemModel_new3, "QStandardItemModel_new3")
	symaddr(&addr_QStandardItemModel_new4, "QStandardItemModel_new4")
	symaddr(&addr_QStandardItemModel_tr, "QStandardItemModel_tr")
	symaddr(&addr_QStandardItemModel_clear, "QStandardItemModel_clear")
	symaddr(&addr_QStandardItemModel_itemFromIndex, "QStandardItemModel_itemFromIndex")
	symaddr(&addr_QStandardItemModel_indexFromItem, "QStandardItemModel_indexFromItem")
	symaddr(&addr_QStandardItemModel_item, "QStandardItemModel_item")
	symaddr(&addr_QStandardItemModel_setItem, "QStandardItemModel_setItem")
	symaddr(&addr_QStandardItemModel_setItem2, "QStandardItemModel_setItem2")
	symaddr(&addr_QStandardItemModel_invisibleRootItem, "QStandardItemModel_invisibleRootItem")
	symaddr(&addr_QStandardItemModel_horizontalHeaderItem, "QStandardItemModel_horizontalHeaderItem")
	symaddr(&addr_QStandardItemModel_setHorizontalHeaderItem, "QStandardItemModel_setHorizontalHeaderItem")
	symaddr(&addr_QStandardItemModel_verticalHeaderItem, "QStandardItemModel_verticalHeaderItem")
	symaddr(&addr_QStandardItemModel_setVerticalHeaderItem, "QStandardItemModel_setVerticalHeaderItem")
	symaddr(&addr_QStandardItemModel_setHorizontalHeaderLabels, "QStandardItemModel_setHorizontalHeaderLabels")
	symaddr(&addr_QStandardItemModel_setVerticalHeaderLabels, "QStandardItemModel_setVerticalHeaderLabels")
	symaddr(&addr_QStandardItemModel_setRowCount, "QStandardItemModel_setRowCount")
	symaddr(&addr_QStandardItemModel_setColumnCount, "QStandardItemModel_setColumnCount")
	symaddr(&addr_QStandardItemModel_appendRow, "QStandardItemModel_appendRow")
	symaddr(&addr_QStandardItemModel_appendColumn, "QStandardItemModel_appendColumn")
	symaddr(&addr_QStandardItemModel_appendRow2, "QStandardItemModel_appendRow2")
	symaddr(&addr_QStandardItemModel_insertRow, "QStandardItemModel_insertRow")
	symaddr(&addr_QStandardItemModel_insertColumn, "QStandardItemModel_insertColumn")
	symaddr(&addr_QStandardItemModel_insertRow2, "QStandardItemModel_insertRow2")
	symaddr(&addr_QStandardItemModel_insertRow3, "QStandardItemModel_insertRow3")
	symaddr(&addr_QStandardItemModel_insertColumn2, "QStandardItemModel_insertColumn2")
	symaddr(&addr_QStandardItemModel_takeItem, "QStandardItemModel_takeItem")
	symaddr(&addr_QStandardItemModel_takeRow, "QStandardItemModel_takeRow")
	symaddr(&addr_QStandardItemModel_takeColumn, "QStandardItemModel_takeColumn")
	symaddr(&addr_QStandardItemModel_takeHorizontalHeaderItem, "QStandardItemModel_takeHorizontalHeaderItem")
	symaddr(&addr_QStandardItemModel_takeVerticalHeaderItem, "QStandardItemModel_takeVerticalHeaderItem")
	symaddr(&addr_QStandardItemModel_itemPrototype, "QStandardItemModel_itemPrototype")
	symaddr(&addr_QStandardItemModel_setItemPrototype, "QStandardItemModel_setItemPrototype")
	symaddr(&addr_QStandardItemModel_findItems, "QStandardItemModel_findItems")
	symaddr(&addr_QStandardItemModel_sortRole, "QStandardItemModel_sortRole")
	symaddr(&addr_QStandardItemModel_setSortRole, "QStandardItemModel_setSortRole")
	symaddr(&addr_QStandardItemModel_itemChanged, "QStandardItemModel_itemChanged")
	symaddr(&addr_QStandardItemModel_createIndex, "QStandardItemModel_createIndex")
	symaddr(&addr_QStandardItemModel_createIndex2, "QStandardItemModel_createIndex2")
	symaddr(&addr_QStandardItemModel_encodeData, "QStandardItemModel_encodeData")
	symaddr(&addr_QStandardItemModel_decodeData, "QStandardItemModel_decodeData")
	symaddr(&addr_QStandardItemModel_beginInsertRows, "QStandardItemModel_beginInsertRows")
	symaddr(&addr_QStandardItemModel_endInsertRows, "QStandardItemModel_endInsertRows")
	symaddr(&addr_QStandardItemModel_beginRemoveRows, "QStandardItemModel_beginRemoveRows")
	symaddr(&addr_QStandardItemModel_endRemoveRows, "QStandardItemModel_endRemoveRows")
	symaddr(&addr_QStandardItemModel_beginMoveRows, "QStandardItemModel_beginMoveRows")
	symaddr(&addr_QStandardItemModel_endMoveRows, "QStandardItemModel_endMoveRows")
	symaddr(&addr_QStandardItemModel_beginInsertColumns, "QStandardItemModel_beginInsertColumns")
	symaddr(&addr_QStandardItemModel_endInsertColumns, "QStandardItemModel_endInsertColumns")
	symaddr(&addr_QStandardItemModel_beginRemoveColumns, "QStandardItemModel_beginRemoveColumns")
	symaddr(&addr_QStandardItemModel_endRemoveColumns, "QStandardItemModel_endRemoveColumns")
	symaddr(&addr_QStandardItemModel_beginMoveColumns, "QStandardItemModel_beginMoveColumns")
	symaddr(&addr_QStandardItemModel_endMoveColumns, "QStandardItemModel_endMoveColumns")
	symaddr(&addr_QStandardItemModel_beginResetModel, "QStandardItemModel_beginResetModel")
	symaddr(&addr_QStandardItemModel_endResetModel, "QStandardItemModel_endResetModel")
	symaddr(&addr_QStandardItemModel_changePersistentIndex, "QStandardItemModel_changePersistentIndex")
	symaddr(&addr_QStandardItemModel_changePersistentIndexList, "QStandardItemModel_changePersistentIndexList")
	symaddr(&addr_QStandardItemModel_persistentIndexList, "QStandardItemModel_persistentIndexList")
	symaddr(&addr_QStandardItemModel_connect, "QStandardItemModel_connect")
	symaddr(&addr_QStandardItemModel_connect2, "QStandardItemModel_connect2")
	symaddr(&addr_QStandardItemModel_sender, "QStandardItemModel_sender")
	symaddr(&addr_QStandardItemModel_senderSignalIndex, "QStandardItemModel_senderSignalIndex")
	symaddr(&addr_QStandardItemModel_receivers, "QStandardItemModel_receivers")
	symaddr(&addr_QStandardItemModel_isSignalConnected, "QStandardItemModel_isSignalConnected")
	symaddr(&addr_QStandardItemModel_tr2, "QStandardItemModel_tr2")
	symaddr(&addr_QStandardItemModel_tr3, "QStandardItemModel_tr3")
	symaddr(&addr_QStandardItemModel_item2, "QStandardItemModel_item2")
	symaddr(&addr_QStandardItemModel_insertRow4, "QStandardItemModel_insertRow4")
	symaddr(&addr_QStandardItemModel_insertColumn3, "QStandardItemModel_insertColumn3")
	symaddr(&addr_QStandardItemModel_takeItem2, "QStandardItemModel_takeItem2")
	symaddr(&addr_QStandardItemModel_findItems2, "QStandardItemModel_findItems2")
	symaddr(&addr_QStandardItemModel_findItems3, "QStandardItemModel_findItems3")
	symaddr(&addr_QStandardItemModel_connect3, "QStandardItemModel_connect3")
	symaddr(&addr_QStandardItemModel_connect4, "QStandardItemModel_connect4")
	symaddr(&addr_QStandardItemModel_metaObject_default, "QStandardItemModel_metaObject_default")
	symaddr(&addr_QStandardItemModel_index_default, "QStandardItemModel_index_default")
	symaddr(&addr_QStandardItemModel_parent_default, "QStandardItemModel_parent_default")
	symaddr(&addr_QStandardItemModel_rowCount_default, "QStandardItemModel_rowCount_default")
	symaddr(&addr_QStandardItemModel_columnCount_default, "QStandardItemModel_columnCount_default")
	symaddr(&addr_QStandardItemModel_hasChildren_default, "QStandardItemModel_hasChildren_default")
	symaddr(&addr_QStandardItemModel_data_default, "QStandardItemModel_data_default")
	symaddr(&addr_QStandardItemModel_multiData_default, "QStandardItemModel_multiData_default")
	symaddr(&addr_QStandardItemModel_setData_default, "QStandardItemModel_setData_default")
	symaddr(&addr_QStandardItemModel_clearItemData_default, "QStandardItemModel_clearItemData_default")
	symaddr(&addr_QStandardItemModel_headerData_default, "QStandardItemModel_headerData_default")
	symaddr(&addr_QStandardItemModel_setHeaderData_default, "QStandardItemModel_setHeaderData_default")
	symaddr(&addr_QStandardItemModel_insertRows_default, "QStandardItemModel_insertRows_default")
	symaddr(&addr_QStandardItemModel_insertColumns_default, "QStandardItemModel_insertColumns_default")
	symaddr(&addr_QStandardItemModel_removeRows_default, "QStandardItemModel_removeRows_default")
	symaddr(&addr_QStandardItemModel_removeColumns_default, "QStandardItemModel_removeColumns_default")
	symaddr(&addr_QStandardItemModel_flags_default, "QStandardItemModel_flags_default")
	symaddr(&addr_QStandardItemModel_supportedDropActions_default, "QStandardItemModel_supportedDropActions_default")
	symaddr(&addr_QStandardItemModel_sort_default, "QStandardItemModel_sort_default")
	symaddr(&addr_QStandardItemModel_mimeTypes_default, "QStandardItemModel_mimeTypes_default")
	symaddr(&addr_QStandardItemModel_mimeData_default, "QStandardItemModel_mimeData_default")
	symaddr(&addr_QStandardItemModel_dropMimeData_default, "QStandardItemModel_dropMimeData_default")
	symaddr(&addr_QStandardItemModel_sibling_default, "QStandardItemModel_sibling_default")
	symaddr(&addr_QStandardItemModel_canDropMimeData_default, "QStandardItemModel_canDropMimeData_default")
	symaddr(&addr_QStandardItemModel_supportedDragActions_default, "QStandardItemModel_supportedDragActions_default")
	symaddr(&addr_QStandardItemModel_moveRows_default, "QStandardItemModel_moveRows_default")
	symaddr(&addr_QStandardItemModel_moveColumns_default, "QStandardItemModel_moveColumns_default")
	symaddr(&addr_QStandardItemModel_fetchMore_default, "QStandardItemModel_fetchMore_default")
	symaddr(&addr_QStandardItemModel_canFetchMore_default, "QStandardItemModel_canFetchMore_default")
	symaddr(&addr_QStandardItemModel_buddy_default, "QStandardItemModel_buddy_default")
	symaddr(&addr_QStandardItemModel_match_default, "QStandardItemModel_match_default")
	symaddr(&addr_QStandardItemModel_span_default, "QStandardItemModel_span_default")
	symaddr(&addr_QStandardItemModel_submit_default, "QStandardItemModel_submit_default")
	symaddr(&addr_QStandardItemModel_revert_default, "QStandardItemModel_revert_default")
	symaddr(&addr_QStandardItemModel_resetInternalData_default, "QStandardItemModel_resetInternalData_default")
	symaddr(&addr_QStandardItemModel_event_default, "QStandardItemModel_event_default")
	symaddr(&addr_QStandardItemModel_eventFilter_default, "QStandardItemModel_eventFilter_default")
	symaddr(&addr_QStandardItemModel_timerEvent_default, "QStandardItemModel_timerEvent_default")
	symaddr(&addr_QStandardItemModel_childEvent_default, "QStandardItemModel_childEvent_default")
	symaddr(&addr_QStandardItemModel_customEvent_default, "QStandardItemModel_customEvent_default")
	symaddr(&addr_QStandardItemModel_connectNotify_default, "QStandardItemModel_connectNotify_default")
	symaddr(&addr_QStandardItemModel_disconnectNotify_default, "QStandardItemModel_disconnectNotify_default")
	reg(&pg_QStandardItemModel_connect_itemChanged, "QStandardItemModel_connect_itemChanged")
	reg(&pg_QStandardItemModel_connect_headerDataChanged, "QStandardItemModel_connect_headerDataChanged")
	reg(&pg_QStandardItemModel_connect_layoutChanged, "QStandardItemModel_connect_layoutChanged")
	reg(&pg_QStandardItemModel_connect_layoutChanged2, "QStandardItemModel_connect_layoutChanged2")
	reg(&pg_QStandardItemModel_connect_layoutAboutToBeChanged, "QStandardItemModel_connect_layoutAboutToBeChanged")
	reg(&pg_QStandardItemModel_connect_layoutAboutToBeChanged2, "QStandardItemModel_connect_layoutAboutToBeChanged2")
	reg(&pg_QStandardItemModel_connect_rowsAboutToBeInserted, "QStandardItemModel_connect_rowsAboutToBeInserted")
	reg(&pg_QStandardItemModel_connect_rowsInserted, "QStandardItemModel_connect_rowsInserted")
	reg(&pg_QStandardItemModel_connect_rowsAboutToBeRemoved, "QStandardItemModel_connect_rowsAboutToBeRemoved")
	reg(&pg_QStandardItemModel_connect_rowsRemoved, "QStandardItemModel_connect_rowsRemoved")
	reg(&pg_QStandardItemModel_connect_columnsAboutToBeInserted, "QStandardItemModel_connect_columnsAboutToBeInserted")
	reg(&pg_QStandardItemModel_connect_columnsInserted, "QStandardItemModel_connect_columnsInserted")
	reg(&pg_QStandardItemModel_connect_columnsAboutToBeRemoved, "QStandardItemModel_connect_columnsAboutToBeRemoved")
	reg(&pg_QStandardItemModel_connect_columnsRemoved, "QStandardItemModel_connect_columnsRemoved")
	reg(&pg_QStandardItemModel_connect_modelAboutToBeReset, "QStandardItemModel_connect_modelAboutToBeReset")
	reg(&pg_QStandardItemModel_connect_modelReset, "QStandardItemModel_connect_modelReset")
	reg(&pg_QStandardItemModel_connect_rowsAboutToBeMoved, "QStandardItemModel_connect_rowsAboutToBeMoved")
	reg(&pg_QStandardItemModel_connect_rowsMoved, "QStandardItemModel_connect_rowsMoved")
	reg(&pg_QStandardItemModel_connect_columnsAboutToBeMoved, "QStandardItemModel_connect_columnsAboutToBeMoved")
	reg(&pg_QStandardItemModel_connect_columnsMoved, "QStandardItemModel_connect_columnsMoved")
	reg(&pg_QStandardItemModel_connect_destroyed, "QStandardItemModel_connect_destroyed")
	reg(&pg_QStandardItemModel_connect_destroyed2, "QStandardItemModel_connect_destroyed2")
	reg(&pg_QStandardItemModel_connect_objectNameChanged, "QStandardItemModel_connect_objectNameChanged")
	reg(&pg_QStandardItemModel__disconnect, "QStandardItemModel__disconnect")
	reg(&pg_QStandardItemModel__set_override, "QStandardItemModel__set_override")
	reg(&pg_QStandardItemModel_delete, "QStandardItemModel_delete")
}

type QStandardItemModel struct {
	h uintptr
	*QAbstractItemModel
}

func bindQStandardItemModel(h uintptr) *QStandardItemModel {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QStandardItemModel_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QStandardItemModel{h: h, QAbstractItemModel: &QAbstractItemModel{h: __c[0], QObject: &QObject{h: __c[1]}}}
}

func newownedQStandardItemModel(h uintptr) *QStandardItemModel { return bindQStandardItemModel(h) }

func (this *QStandardItemModel) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QStandardItemModel) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQStandardItemModel(h unsafe.Pointer) *QStandardItemModel { return bindQStandardItemModel(uintptr(h)) }
func UnsafeNewQStandardItemModelOwned(h unsafe.Pointer) *QStandardItemModel { return newownedQStandardItemModel(uintptr(h)) }

func (this *QStandardItemModel) UnsafeRelease() {
	this.h = 0
	if this.QAbstractItemModel != nil {
		this.QAbstractItemModel.UnsafeRelease()
	}
}

func NewQStandardItemModel() *QStandardItemModel {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QStandardItemModel_new, uintptr(__goid))
	__o := newownedQStandardItemModel(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QStandardItemModel_virtIdx, QStandardItemModel_enable)
	return __o
}

func NewQStandardItemModel2(rows int, columns int) *QStandardItemModel {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QStandardItemModel_new2, uintptr(__goid), uintptr(rows), uintptr(columns))
	__o := newownedQStandardItemModel(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QStandardItemModel_virtIdx, QStandardItemModel_enable)
	return __o
}

func NewQStandardItemModel3(parent *QObject) *QStandardItemModel {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QStandardItemModel_new3, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQStandardItemModel(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QStandardItemModel_virtIdx, QStandardItemModel_enable)
	return __o
}

func NewQStandardItemModel4(rows int, columns int, parent *QObject) *QStandardItemModel {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QStandardItemModel_new4, uintptr(__goid), uintptr(rows), uintptr(columns), uintptr(parent.UnsafePointer()))
	__o := newownedQStandardItemModel(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QStandardItemModel_virtIdx, QStandardItemModel_enable)
	return __o
}

func QStandardItemModel_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStandardItemModel_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QStandardItemModel) Clear() {
	purego.SyscallN(addr_QStandardItemModel_clear, this.h)
	runtime.KeepAlive(this)
}

func (this *QStandardItemModel) ItemFromIndex(index *QModelIndex) *QStandardItem {
	__r, _, _ := purego.SyscallN(addr_QStandardItemModel_itemFromIndex, this.h, uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return bindQStandardItem(__r)
}

func (this *QStandardItemModel) IndexFromItem(item *QStandardItem) *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QStandardItemModel_indexFromItem, this.h, uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
	return newownedQModelIndex(__r)
}

func (this *QStandardItemModel) Item(row int) *QStandardItem {
	__r, _, _ := purego.SyscallN(addr_QStandardItemModel_item, this.h, uintptr(row))
	runtime.KeepAlive(this)
	return bindQStandardItem(__r)
}

func (this *QStandardItemModel) SetItem(row int, column int, item *QStandardItem) {
	purego.SyscallN(addr_QStandardItemModel_setItem, this.h, uintptr(row), uintptr(column), uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QStandardItemModel) SetItem2(row int, item *QStandardItem) {
	purego.SyscallN(addr_QStandardItemModel_setItem2, this.h, uintptr(row), uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QStandardItemModel) InvisibleRootItem() *QStandardItem {
	__r, _, _ := purego.SyscallN(addr_QStandardItemModel_invisibleRootItem, this.h)
	runtime.KeepAlive(this)
	return bindQStandardItem(__r)
}

func (this *QStandardItemModel) HorizontalHeaderItem(column int) *QStandardItem {
	__r, _, _ := purego.SyscallN(addr_QStandardItemModel_horizontalHeaderItem, this.h, uintptr(column))
	runtime.KeepAlive(this)
	return bindQStandardItem(__r)
}

func (this *QStandardItemModel) SetHorizontalHeaderItem(column int, item *QStandardItem) {
	purego.SyscallN(addr_QStandardItemModel_setHorizontalHeaderItem, this.h, uintptr(column), uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QStandardItemModel) VerticalHeaderItem(row int) *QStandardItem {
	__r, _, _ := purego.SyscallN(addr_QStandardItemModel_verticalHeaderItem, this.h, uintptr(row))
	runtime.KeepAlive(this)
	return bindQStandardItem(__r)
}

func (this *QStandardItemModel) SetVerticalHeaderItem(row int, item *QStandardItem) {
	purego.SyscallN(addr_QStandardItemModel_setVerticalHeaderItem, this.h, uintptr(row), uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QStandardItemModel) SetHorizontalHeaderLabels(labels []string) {
	labels_arr := make([]uintptr, len(labels))
	for __i, __e := range labels {
		labels_arr[__i] = pgCString(__e)
	}
	var labels_p uintptr
	if len(labels_arr) > 0 {
		labels_p = uintptr(unsafe.Pointer(&labels_arr[0]))
	}
	purego.SyscallN(addr_QStandardItemModel_setHorizontalHeaderLabels, this.h, labels_p, uintptr(len(labels)))
	runtime.KeepAlive(this)
	for __i := range labels_arr {
		pgFree(labels_arr[__i])
	}
	runtime.KeepAlive(labels_arr)
}

func (this *QStandardItemModel) SetVerticalHeaderLabels(labels []string) {
	labels_arr := make([]uintptr, len(labels))
	for __i, __e := range labels {
		labels_arr[__i] = pgCString(__e)
	}
	var labels_p uintptr
	if len(labels_arr) > 0 {
		labels_p = uintptr(unsafe.Pointer(&labels_arr[0]))
	}
	purego.SyscallN(addr_QStandardItemModel_setVerticalHeaderLabels, this.h, labels_p, uintptr(len(labels)))
	runtime.KeepAlive(this)
	for __i := range labels_arr {
		pgFree(labels_arr[__i])
	}
	runtime.KeepAlive(labels_arr)
}

func (this *QStandardItemModel) SetRowCount(rows int) {
	purego.SyscallN(addr_QStandardItemModel_setRowCount, this.h, uintptr(rows))
	runtime.KeepAlive(this)
}

func (this *QStandardItemModel) SetColumnCount(columns int) {
	purego.SyscallN(addr_QStandardItemModel_setColumnCount, this.h, uintptr(columns))
	runtime.KeepAlive(this)
}

func (this *QStandardItemModel) AppendRow(items []*QStandardItem) {
	items_arr := make([]uintptr, len(items))
	for __i, __e := range items {
		items_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var items_p uintptr
	if len(items_arr) > 0 {
		items_p = uintptr(unsafe.Pointer(&items_arr[0]))
	}
	purego.SyscallN(addr_QStandardItemModel_appendRow, this.h, items_p, uintptr(len(items)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(items)
	runtime.KeepAlive(items_arr)
}

func (this *QStandardItemModel) AppendColumn(items []*QStandardItem) {
	items_arr := make([]uintptr, len(items))
	for __i, __e := range items {
		items_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var items_p uintptr
	if len(items_arr) > 0 {
		items_p = uintptr(unsafe.Pointer(&items_arr[0]))
	}
	purego.SyscallN(addr_QStandardItemModel_appendColumn, this.h, items_p, uintptr(len(items)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(items)
	runtime.KeepAlive(items_arr)
}

func (this *QStandardItemModel) AppendRowWithItem(item *QStandardItem) {
	purego.SyscallN(addr_QStandardItemModel_appendRow2, this.h, uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QStandardItemModel) InsertRow(row int, items []*QStandardItem) {
	items_arr := make([]uintptr, len(items))
	for __i, __e := range items {
		items_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var items_p uintptr
	if len(items_arr) > 0 {
		items_p = uintptr(unsafe.Pointer(&items_arr[0]))
	}
	purego.SyscallN(addr_QStandardItemModel_insertRow, this.h, uintptr(row), items_p, uintptr(len(items)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(items)
	runtime.KeepAlive(items_arr)
}

func (this *QStandardItemModel) InsertColumn(column int, items []*QStandardItem) {
	items_arr := make([]uintptr, len(items))
	for __i, __e := range items {
		items_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var items_p uintptr
	if len(items_arr) > 0 {
		items_p = uintptr(unsafe.Pointer(&items_arr[0]))
	}
	purego.SyscallN(addr_QStandardItemModel_insertColumn, this.h, uintptr(column), items_p, uintptr(len(items)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(items)
	runtime.KeepAlive(items_arr)
}

func (this *QStandardItemModel) InsertRow2(row int, item *QStandardItem) {
	purego.SyscallN(addr_QStandardItemModel_insertRow2, this.h, uintptr(row), uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QStandardItemModel) InsertRowWithRow(row int) bool {
	__r0, _, _ := purego.SyscallN(addr_QStandardItemModel_insertRow3, this.h, uintptr(row))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QStandardItemModel) InsertColumnWithColumn(column int) bool {
	__r0, _, _ := purego.SyscallN(addr_QStandardItemModel_insertColumn2, this.h, uintptr(column))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QStandardItemModel) TakeItem(row int) *QStandardItem {
	__r, _, _ := purego.SyscallN(addr_QStandardItemModel_takeItem, this.h, uintptr(row))
	runtime.KeepAlive(this)
	return bindQStandardItem(__r)
}

func (this *QStandardItemModel) TakeRow(row int) []*QStandardItem {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStandardItemModel_takeRow, this.h, uintptr(row), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QStandardItem, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = bindQStandardItem(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QStandardItemModel) TakeColumn(column int) []*QStandardItem {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStandardItemModel_takeColumn, this.h, uintptr(column), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QStandardItem, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = bindQStandardItem(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QStandardItemModel) TakeHorizontalHeaderItem(column int) *QStandardItem {
	__r, _, _ := purego.SyscallN(addr_QStandardItemModel_takeHorizontalHeaderItem, this.h, uintptr(column))
	runtime.KeepAlive(this)
	return bindQStandardItem(__r)
}

func (this *QStandardItemModel) TakeVerticalHeaderItem(row int) *QStandardItem {
	__r, _, _ := purego.SyscallN(addr_QStandardItemModel_takeVerticalHeaderItem, this.h, uintptr(row))
	runtime.KeepAlive(this)
	return bindQStandardItem(__r)
}

func (this *QStandardItemModel) ItemPrototype() *QStandardItem {
	__r, _, _ := purego.SyscallN(addr_QStandardItemModel_itemPrototype, this.h)
	runtime.KeepAlive(this)
	return bindQStandardItem(__r)
}

func (this *QStandardItemModel) SetItemPrototype(item *QStandardItem) {
	purego.SyscallN(addr_QStandardItemModel_setItemPrototype, this.h, uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QStandardItemModel) FindItems(text string) []*QStandardItem {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStandardItemModel_findItems, this.h, text_p, uintptr(len(text)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
	__out := make([]*QStandardItem, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = bindQStandardItem(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QStandardItemModel) SortRole() int {
	__r0, _, _ := purego.SyscallN(addr_QStandardItemModel_sortRole, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QStandardItemModel) SetSortRole(role int) {
	purego.SyscallN(addr_QStandardItemModel_setSortRole, this.h, uintptr(role))
	runtime.KeepAlive(this)
}

func (this *QStandardItemModel) ItemChanged(item *QStandardItem) {
	purego.SyscallN(addr_QStandardItemModel_itemChanged, this.h, uintptr(item.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(item)
}

func (this *QStandardItemModel) CreateIndex(row int, column int) *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QStandardItemModel_createIndex, this.h, uintptr(row), uintptr(column))
	runtime.KeepAlive(this)
	return newownedQModelIndex(__r)
}

func (this *QStandardItemModel) CreateIndex2(row int, column int, id uintptr) *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QStandardItemModel_createIndex2, this.h, uintptr(row), uintptr(column), uintptr(id))
	runtime.KeepAlive(this)
	return newownedQModelIndex(__r)
}

func (this *QStandardItemModel) EncodeData(indexes []*QModelIndex, stream *QDataStream) {
	indexes_arr := make([]uintptr, len(indexes))
	for __i, __e := range indexes {
		indexes_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var indexes_p uintptr
	if len(indexes_arr) > 0 {
		indexes_p = uintptr(unsafe.Pointer(&indexes_arr[0]))
	}
	purego.SyscallN(addr_QStandardItemModel_encodeData, this.h, indexes_p, uintptr(len(indexes)), uintptr(stream.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(indexes)
	runtime.KeepAlive(indexes_arr)
	runtime.KeepAlive(stream)
}

func (this *QStandardItemModel) DecodeData(row int, column int, parent *QModelIndex, stream *QDataStream) bool {
	__r0, _, _ := purego.SyscallN(addr_QStandardItemModel_decodeData, this.h, uintptr(row), uintptr(column), uintptr(parent.UnsafePointer()), uintptr(stream.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	runtime.KeepAlive(stream)
	return __r != 0
}

func (this *QStandardItemModel) BeginInsertRows(parent *QModelIndex, first int, last int) {
	purego.SyscallN(addr_QStandardItemModel_beginInsertRows, this.h, uintptr(parent.UnsafePointer()), uintptr(first), uintptr(last))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
}

func (this *QStandardItemModel) EndInsertRows() {
	purego.SyscallN(addr_QStandardItemModel_endInsertRows, this.h)
	runtime.KeepAlive(this)
}

func (this *QStandardItemModel) BeginRemoveRows(parent *QModelIndex, first int, last int) {
	purego.SyscallN(addr_QStandardItemModel_beginRemoveRows, this.h, uintptr(parent.UnsafePointer()), uintptr(first), uintptr(last))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
}

func (this *QStandardItemModel) EndRemoveRows() {
	purego.SyscallN(addr_QStandardItemModel_endRemoveRows, this.h)
	runtime.KeepAlive(this)
}

func (this *QStandardItemModel) BeginMoveRows(sourceParent *QModelIndex, sourceFirst int, sourceLast int, destinationParent *QModelIndex, destinationRow int) bool {
	__r0, _, _ := purego.SyscallN(addr_QStandardItemModel_beginMoveRows, this.h, uintptr(sourceParent.UnsafePointer()), uintptr(sourceFirst), uintptr(sourceLast), uintptr(destinationParent.UnsafePointer()), uintptr(destinationRow))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sourceParent)
	runtime.KeepAlive(destinationParent)
	return __r != 0
}

func (this *QStandardItemModel) EndMoveRows() {
	purego.SyscallN(addr_QStandardItemModel_endMoveRows, this.h)
	runtime.KeepAlive(this)
}

func (this *QStandardItemModel) BeginInsertColumns(parent *QModelIndex, first int, last int) {
	purego.SyscallN(addr_QStandardItemModel_beginInsertColumns, this.h, uintptr(parent.UnsafePointer()), uintptr(first), uintptr(last))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
}

func (this *QStandardItemModel) EndInsertColumns() {
	purego.SyscallN(addr_QStandardItemModel_endInsertColumns, this.h)
	runtime.KeepAlive(this)
}

func (this *QStandardItemModel) BeginRemoveColumns(parent *QModelIndex, first int, last int) {
	purego.SyscallN(addr_QStandardItemModel_beginRemoveColumns, this.h, uintptr(parent.UnsafePointer()), uintptr(first), uintptr(last))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
}

func (this *QStandardItemModel) EndRemoveColumns() {
	purego.SyscallN(addr_QStandardItemModel_endRemoveColumns, this.h)
	runtime.KeepAlive(this)
}

func (this *QStandardItemModel) BeginMoveColumns(sourceParent *QModelIndex, sourceFirst int, sourceLast int, destinationParent *QModelIndex, destinationColumn int) bool {
	__r0, _, _ := purego.SyscallN(addr_QStandardItemModel_beginMoveColumns, this.h, uintptr(sourceParent.UnsafePointer()), uintptr(sourceFirst), uintptr(sourceLast), uintptr(destinationParent.UnsafePointer()), uintptr(destinationColumn))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sourceParent)
	runtime.KeepAlive(destinationParent)
	return __r != 0
}

func (this *QStandardItemModel) EndMoveColumns() {
	purego.SyscallN(addr_QStandardItemModel_endMoveColumns, this.h)
	runtime.KeepAlive(this)
}

func (this *QStandardItemModel) BeginResetModel() {
	purego.SyscallN(addr_QStandardItemModel_beginResetModel, this.h)
	runtime.KeepAlive(this)
}

func (this *QStandardItemModel) EndResetModel() {
	purego.SyscallN(addr_QStandardItemModel_endResetModel, this.h)
	runtime.KeepAlive(this)
}

func (this *QStandardItemModel) ChangePersistentIndex(from *QModelIndex, to *QModelIndex) {
	purego.SyscallN(addr_QStandardItemModel_changePersistentIndex, this.h, uintptr(from.UnsafePointer()), uintptr(to.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(from)
	runtime.KeepAlive(to)
}

func (this *QStandardItemModel) ChangePersistentIndexList(from []*QModelIndex, to []*QModelIndex) {
	from_arr := make([]uintptr, len(from))
	for __i, __e := range from {
		from_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var from_p uintptr
	if len(from_arr) > 0 {
		from_p = uintptr(unsafe.Pointer(&from_arr[0]))
	}
	to_arr := make([]uintptr, len(to))
	for __i, __e := range to {
		to_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var to_p uintptr
	if len(to_arr) > 0 {
		to_p = uintptr(unsafe.Pointer(&to_arr[0]))
	}
	purego.SyscallN(addr_QStandardItemModel_changePersistentIndexList, this.h, from_p, uintptr(len(from)), to_p, uintptr(len(to)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(from)
	runtime.KeepAlive(from_arr)
	runtime.KeepAlive(to)
	runtime.KeepAlive(to_arr)
}

func (this *QStandardItemModel) PersistentIndexList() []*QModelIndex {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStandardItemModel_persistentIndexList, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QModelIndex, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQModelIndex(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QStandardItemModel_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QStandardItemModel_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QStandardItemModel_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QStandardItemModel_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QStandardItemModel) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QStandardItemModel_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QStandardItemModel) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QStandardItemModel_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QStandardItemModel) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QStandardItemModel_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QStandardItemModel) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QStandardItemModel_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QStandardItemModel_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStandardItemModel_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QStandardItemModel_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStandardItemModel_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func (this *QStandardItemModel) Item2(row int, column int) *QStandardItem {
	__r, _, _ := purego.SyscallN(addr_QStandardItemModel_item2, this.h, uintptr(row), uintptr(column))
	runtime.KeepAlive(this)
	return bindQStandardItem(__r)
}

func (this *QStandardItemModel) InsertRow3(row int, parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QStandardItemModel_insertRow4, this.h, uintptr(row), uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QStandardItemModel) InsertColumn2(column int, parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QStandardItemModel_insertColumn3, this.h, uintptr(column), uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QStandardItemModel) TakeItem2(row int, column int) *QStandardItem {
	__r, _, _ := purego.SyscallN(addr_QStandardItemModel_takeItem2, this.h, uintptr(row), uintptr(column))
	runtime.KeepAlive(this)
	return bindQStandardItem(__r)
}

func (this *QStandardItemModel) FindItems2(text string, flags MatchFlag) []*QStandardItem {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStandardItemModel_findItems2, this.h, text_p, uintptr(len(text)), uintptr(flags), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
	__out := make([]*QStandardItem, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = bindQStandardItem(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QStandardItemModel) FindItems3(text string, flags MatchFlag, column int) []*QStandardItem {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStandardItemModel_findItems3, this.h, text_p, uintptr(len(text)), uintptr(flags), uintptr(column), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
	__out := make([]*QStandardItem, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = bindQStandardItem(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QStandardItemModel_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QStandardItemModel_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QStandardItemModel_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QStandardItemModel_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QStandardItemModel) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QStandardItemModel_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QStandardItemModel) IndexDefault(row int, column int, parent *QModelIndex) *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QStandardItemModel_index_default, this.h, uintptr(row), uintptr(column), uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return newownedQModelIndex(__r)
}

func (this *QStandardItemModel) ParentDefault(child *QModelIndex) *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QStandardItemModel_parent_default, this.h, uintptr(child.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(child)
	return newownedQModelIndex(__r)
}

func (this *QStandardItemModel) RowCountDefault(parent *QModelIndex) int {
	__r0, _, _ := purego.SyscallN(addr_QStandardItemModel_rowCount_default, this.h, uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return int(__r)
}

func (this *QStandardItemModel) ColumnCountDefault(parent *QModelIndex) int {
	__r0, _, _ := purego.SyscallN(addr_QStandardItemModel_columnCount_default, this.h, uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return int(__r)
}

func (this *QStandardItemModel) HasChildrenDefault(parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QStandardItemModel_hasChildren_default, this.h, uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QStandardItemModel) DataDefault(index *QModelIndex, role int) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QStandardItemModel_data_default, this.h, uintptr(index.UnsafePointer()), uintptr(role))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return newownedQVariant(__r)
}

func (this *QStandardItemModel) MultiDataDefault(index *QModelIndex, roleDataSpan *QModelRoleDataSpan) {
	purego.SyscallN(addr_QStandardItemModel_multiData_default, this.h, uintptr(index.UnsafePointer()), uintptr(roleDataSpan.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	runtime.KeepAlive(roleDataSpan)
}

func (this *QStandardItemModel) SetDataDefault(index *QModelIndex, value *QVariant, role int) bool {
	__r0, _, _ := purego.SyscallN(addr_QStandardItemModel_setData_default, this.h, uintptr(index.UnsafePointer()), uintptr(value.UnsafePointer()), uintptr(role))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	runtime.KeepAlive(value)
	return __r != 0
}

func (this *QStandardItemModel) ClearItemDataDefault(index *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QStandardItemModel_clearItemData_default, this.h, uintptr(index.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return __r != 0
}

func (this *QStandardItemModel) HeaderDataDefault(section int, orientation Orientation, role int) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QStandardItemModel_headerData_default, this.h, uintptr(section), uintptr(orientation), uintptr(role))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QStandardItemModel) SetHeaderDataDefault(section int, orientation Orientation, value *QVariant, role int) bool {
	__r0, _, _ := purego.SyscallN(addr_QStandardItemModel_setHeaderData_default, this.h, uintptr(section), uintptr(orientation), uintptr(value.UnsafePointer()), uintptr(role))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(value)
	return __r != 0
}

func (this *QStandardItemModel) InsertRowsDefault(row int, count int, parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QStandardItemModel_insertRows_default, this.h, uintptr(row), uintptr(count), uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QStandardItemModel) InsertColumnsDefault(column int, count int, parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QStandardItemModel_insertColumns_default, this.h, uintptr(column), uintptr(count), uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QStandardItemModel) RemoveRowsDefault(row int, count int, parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QStandardItemModel_removeRows_default, this.h, uintptr(row), uintptr(count), uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QStandardItemModel) RemoveColumnsDefault(column int, count int, parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QStandardItemModel_removeColumns_default, this.h, uintptr(column), uintptr(count), uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QStandardItemModel) FlagsDefault(index *QModelIndex) ItemFlag {
	__r0, _, _ := purego.SyscallN(addr_QStandardItemModel_flags_default, this.h, uintptr(index.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return ItemFlag(__r)
}

func (this *QStandardItemModel) SupportedDropActionsDefault() DropAction {
	__r0, _, _ := purego.SyscallN(addr_QStandardItemModel_supportedDropActions_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return DropAction(__r)
}

func (this *QStandardItemModel) SortDefault(column int, order SortOrder) {
	purego.SyscallN(addr_QStandardItemModel_sort_default, this.h, uintptr(column), uintptr(order))
	runtime.KeepAlive(this)
}

func (this *QStandardItemModel) MimeTypesDefault() []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStandardItemModel_mimeTypes_default, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]string, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = pgGoStr(__sl[__i])
		pgFree(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QStandardItemModel) MimeDataDefault(indexes []*QModelIndex) *QMimeData {
	indexes_arr := make([]uintptr, len(indexes))
	for __i, __e := range indexes {
		indexes_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var indexes_p uintptr
	if len(indexes_arr) > 0 {
		indexes_p = uintptr(unsafe.Pointer(&indexes_arr[0]))
	}
	__r, _, _ := purego.SyscallN(addr_QStandardItemModel_mimeData_default, this.h, indexes_p, uintptr(len(indexes)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(indexes)
	runtime.KeepAlive(indexes_arr)
	return bindQMimeData(__r)
}

func (this *QStandardItemModel) DropMimeDataDefault(data *QMimeData, action DropAction, row int, column int, parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QStandardItemModel_dropMimeData_default, this.h, uintptr(data.UnsafePointer()), uintptr(action), uintptr(row), uintptr(column), uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(data)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QStandardItemModel) SiblingDefault(row int, column int, idx *QModelIndex) *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QStandardItemModel_sibling_default, this.h, uintptr(row), uintptr(column), uintptr(idx.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(idx)
	return newownedQModelIndex(__r)
}

func (this *QStandardItemModel) CanDropMimeDataDefault(data *QMimeData, action DropAction, row int, column int, parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QStandardItemModel_canDropMimeData_default, this.h, uintptr(data.UnsafePointer()), uintptr(action), uintptr(row), uintptr(column), uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(data)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QStandardItemModel) SupportedDragActionsDefault() DropAction {
	__r0, _, _ := purego.SyscallN(addr_QStandardItemModel_supportedDragActions_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return DropAction(__r)
}

func (this *QStandardItemModel) MoveRowsDefault(sourceParent *QModelIndex, sourceRow int, count int, destinationParent *QModelIndex, destinationChild int) bool {
	__r0, _, _ := purego.SyscallN(addr_QStandardItemModel_moveRows_default, this.h, uintptr(sourceParent.UnsafePointer()), uintptr(sourceRow), uintptr(count), uintptr(destinationParent.UnsafePointer()), uintptr(destinationChild))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sourceParent)
	runtime.KeepAlive(destinationParent)
	return __r != 0
}

func (this *QStandardItemModel) MoveColumnsDefault(sourceParent *QModelIndex, sourceColumn int, count int, destinationParent *QModelIndex, destinationChild int) bool {
	__r0, _, _ := purego.SyscallN(addr_QStandardItemModel_moveColumns_default, this.h, uintptr(sourceParent.UnsafePointer()), uintptr(sourceColumn), uintptr(count), uintptr(destinationParent.UnsafePointer()), uintptr(destinationChild))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sourceParent)
	runtime.KeepAlive(destinationParent)
	return __r != 0
}

func (this *QStandardItemModel) FetchMoreDefault(parent *QModelIndex) {
	purego.SyscallN(addr_QStandardItemModel_fetchMore_default, this.h, uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
}

func (this *QStandardItemModel) CanFetchMoreDefault(parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QStandardItemModel_canFetchMore_default, this.h, uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QStandardItemModel) BuddyDefault(index *QModelIndex) *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QStandardItemModel_buddy_default, this.h, uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return newownedQModelIndex(__r)
}

func (this *QStandardItemModel) MatchDefault(start *QModelIndex, role int, value *QVariant, hits int, flags MatchFlag) []*QModelIndex {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStandardItemModel_match_default, this.h, uintptr(start.UnsafePointer()), uintptr(role), uintptr(value.UnsafePointer()), uintptr(hits), uintptr(flags), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(start)
	runtime.KeepAlive(value)
	__out := make([]*QModelIndex, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQModelIndex(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QStandardItemModel) SpanDefault(index *QModelIndex) *QSize {
	__r, _, _ := purego.SyscallN(addr_QStandardItemModel_span_default, this.h, uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return newownedQSize(__r)
}

func (this *QStandardItemModel) SubmitDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QStandardItemModel_submit_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QStandardItemModel) RevertDefault() {
	purego.SyscallN(addr_QStandardItemModel_revert_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QStandardItemModel) ResetInternalDataDefault() {
	purego.SyscallN(addr_QStandardItemModel_resetInternalData_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QStandardItemModel) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QStandardItemModel_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QStandardItemModel) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QStandardItemModel_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QStandardItemModel) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QStandardItemModel_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QStandardItemModel) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QStandardItemModel_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QStandardItemModel) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QStandardItemModel_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QStandardItemModel) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QStandardItemModel_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QStandardItemModel) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QStandardItemModel_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QStandardItemModel) OnItemChanged(fn func(*QStandardItem)) int64 {
	id := registerCallback(func(item uintptr) { fn(UnsafeNewQStandardItem(unsafe.Pointer(item))) })
	pg_QStandardItemModel_connect_itemChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStandardItemModel) OnHeaderDataChanged(fn func(Orientation, int, int)) int64 {
	id := registerCallback(func(orientation int32, first int32, last int32) { fn(Orientation(orientation), int(first), int(last)) })
	pg_QStandardItemModel_connect_headerDataChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStandardItemModel) OnLayoutChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QStandardItemModel_connect_layoutChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStandardItemModel) OnLayoutChanged2(fn func([]*QPersistentModelIndex, QAbstractItemModel__LayoutChangeHint)) int64 {
	id := registerCallback(func(parents uintptr, parents_len int64, hint int32) { fn(func() []*QPersistentModelIndex { if parents == 0 || parents_len == 0 { return nil }; __sl := unsafe.Slice((*unsafe.Pointer)(unsafe.Pointer(parents)), int(parents_len)); __o := make([]*QPersistentModelIndex, int(parents_len)); for __i := range __o { __o[__i] = UnsafeNewQPersistentModelIndex(__sl[__i]) }; return __o }(), QAbstractItemModel__LayoutChangeHint(hint)) })
	pg_QStandardItemModel_connect_layoutChanged2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStandardItemModel) OnLayoutAboutToBeChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QStandardItemModel_connect_layoutAboutToBeChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStandardItemModel) OnLayoutAboutToBeChanged2(fn func([]*QPersistentModelIndex, QAbstractItemModel__LayoutChangeHint)) int64 {
	id := registerCallback(func(parents uintptr, parents_len int64, hint int32) { fn(func() []*QPersistentModelIndex { if parents == 0 || parents_len == 0 { return nil }; __sl := unsafe.Slice((*unsafe.Pointer)(unsafe.Pointer(parents)), int(parents_len)); __o := make([]*QPersistentModelIndex, int(parents_len)); for __i := range __o { __o[__i] = UnsafeNewQPersistentModelIndex(__sl[__i]) }; return __o }(), QAbstractItemModel__LayoutChangeHint(hint)) })
	pg_QStandardItemModel_connect_layoutAboutToBeChanged2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStandardItemModel) OnRowsAboutToBeInserted(fn func(*QModelIndex, int, int)) int64 {
	id := registerCallback(func(parent uintptr, first int32, last int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(first), int(last)) })
	pg_QStandardItemModel_connect_rowsAboutToBeInserted(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStandardItemModel) OnRowsInserted(fn func(*QModelIndex, int, int)) int64 {
	id := registerCallback(func(parent uintptr, first int32, last int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(first), int(last)) })
	pg_QStandardItemModel_connect_rowsInserted(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStandardItemModel) OnRowsAboutToBeRemoved(fn func(*QModelIndex, int, int)) int64 {
	id := registerCallback(func(parent uintptr, first int32, last int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(first), int(last)) })
	pg_QStandardItemModel_connect_rowsAboutToBeRemoved(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStandardItemModel) OnRowsRemoved(fn func(*QModelIndex, int, int)) int64 {
	id := registerCallback(func(parent uintptr, first int32, last int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(first), int(last)) })
	pg_QStandardItemModel_connect_rowsRemoved(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStandardItemModel) OnColumnsAboutToBeInserted(fn func(*QModelIndex, int, int)) int64 {
	id := registerCallback(func(parent uintptr, first int32, last int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(first), int(last)) })
	pg_QStandardItemModel_connect_columnsAboutToBeInserted(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStandardItemModel) OnColumnsInserted(fn func(*QModelIndex, int, int)) int64 {
	id := registerCallback(func(parent uintptr, first int32, last int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(first), int(last)) })
	pg_QStandardItemModel_connect_columnsInserted(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStandardItemModel) OnColumnsAboutToBeRemoved(fn func(*QModelIndex, int, int)) int64 {
	id := registerCallback(func(parent uintptr, first int32, last int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(first), int(last)) })
	pg_QStandardItemModel_connect_columnsAboutToBeRemoved(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStandardItemModel) OnColumnsRemoved(fn func(*QModelIndex, int, int)) int64 {
	id := registerCallback(func(parent uintptr, first int32, last int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(first), int(last)) })
	pg_QStandardItemModel_connect_columnsRemoved(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStandardItemModel) OnModelAboutToBeReset(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QStandardItemModel_connect_modelAboutToBeReset(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStandardItemModel) OnModelReset(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QStandardItemModel_connect_modelReset(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStandardItemModel) OnRowsAboutToBeMoved(fn func(*QModelIndex, int, int, *QModelIndex, int)) int64 {
	id := registerCallback(func(sourceParent uintptr, sourceStart int32, sourceEnd int32, destinationParent uintptr, destinationRow int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(sourceParent)), int(sourceStart), int(sourceEnd), UnsafeNewQModelIndex(unsafe.Pointer(destinationParent)), int(destinationRow)) })
	pg_QStandardItemModel_connect_rowsAboutToBeMoved(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStandardItemModel) OnRowsMoved(fn func(*QModelIndex, int, int, *QModelIndex, int)) int64 {
	id := registerCallback(func(sourceParent uintptr, sourceStart int32, sourceEnd int32, destinationParent uintptr, destinationRow int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(sourceParent)), int(sourceStart), int(sourceEnd), UnsafeNewQModelIndex(unsafe.Pointer(destinationParent)), int(destinationRow)) })
	pg_QStandardItemModel_connect_rowsMoved(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStandardItemModel) OnColumnsAboutToBeMoved(fn func(*QModelIndex, int, int, *QModelIndex, int)) int64 {
	id := registerCallback(func(sourceParent uintptr, sourceStart int32, sourceEnd int32, destinationParent uintptr, destinationColumn int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(sourceParent)), int(sourceStart), int(sourceEnd), UnsafeNewQModelIndex(unsafe.Pointer(destinationParent)), int(destinationColumn)) })
	pg_QStandardItemModel_connect_columnsAboutToBeMoved(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStandardItemModel) OnColumnsMoved(fn func(*QModelIndex, int, int, *QModelIndex, int)) int64 {
	id := registerCallback(func(sourceParent uintptr, sourceStart int32, sourceEnd int32, destinationParent uintptr, destinationColumn int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(sourceParent)), int(sourceStart), int(sourceEnd), UnsafeNewQModelIndex(unsafe.Pointer(destinationParent)), int(destinationColumn)) })
	pg_QStandardItemModel_connect_columnsMoved(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStandardItemModel) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QStandardItemModel_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStandardItemModel) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QStandardItemModel_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStandardItemModel) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QStandardItemModel_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStandardItemModel) Disconnect(handle int64) {
	pg_QStandardItemModel__disconnect(handle)
	unregisterCallback(handle)
}

func QStandardItemModel_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case -8948777187306027381:
		return 1
	case 7371339085820260189:
		return 2
	case 1960960171956884010:
		return 3
	case -8256963024669584834:
		return 4
	case 3369452259144514380:
		return 5
	case -8837375942098466299:
		return 6
	case 731311060944096642:
		return 7
	case 35191454575202219:
		return 8
	case 8914190500520542409:
		return 9
	case -3297140021896693688:
		return 10
	case 4852200306590054850:
		return 11
	case -2775729825962395603:
		return 12
	case -8973293846230792237:
		return 13
	case 5104072549182222656:
		return 14
	case 2078548737056540420:
		return 15
	case 1703382833545304812:
		return 16
	case -7784642176186821205:
		return 17
	case 6912441694453382161:
		return 18
	case 5464467312394730418:
		return 19
	case 5175297091024974749:
		return 20
	case -9049994684305682212:
		return 21
	case 397329346813782521:
		return 22
	case 731539533644753052:
		return 23
	case 3421601386088570762:
		return 24
	case -6074819356947405415:
		return 25
	case -6321007591388309905:
		return 26
	case 3553294593611130806:
		return 27
	case 8639262066473964150:
		return 28
	case 6112860112248786499:
		return 29
	case -4341501217981328394:
		return 30
	case -8395342916587040543:
		return 31
	case -3788331537442685179:
		return 32
	case -1522951607536117149:
		return 33
	case -4510675770784051183:
		return 34
	case 1237236456312824991:
		return 35
	case 5217623514630858297:
		return 36
	case -5409927581812374498:
		return 37
	case -1185315367526857833:
		return 38
	case 3769868229123868474:
		return 39
	case 7379029982170503008:
		return 40
	case -568369772982215962:
		return 41
	}
	return -1
}

func QStandardItemModel_enable(h uintptr, idx int32) { pg_QStandardItemModel__set_override(h, idx) }

func (this *QStandardItemModel) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QStandardItemModel) OnIndex(fn func(int, int, *QModelIndex) *QModelIndex) {
	registerVirt(this.UnsafeGoID(), -8948777187306027381, func(row int32, column int32, parent uintptr) uintptr { return uintptr(fn(int(row), int(column), UnsafeNewQModelIndex(unsafe.Pointer(parent))).UnsafePointer()) }) // slot "index"
}

func (this *QStandardItemModel) OnParent(fn func(*QModelIndex) *QModelIndex) {
	registerVirt(this.UnsafeGoID(), 7371339085820260189, func(child uintptr) uintptr { return uintptr(fn(UnsafeNewQModelIndex(unsafe.Pointer(child))).UnsafePointer()) }) // slot "parent"
}

func (this *QStandardItemModel) OnRowCount(fn func(*QModelIndex) int) {
	registerVirt(this.UnsafeGoID(), 1960960171956884010, func(parent uintptr) int32 { return int32(fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)))) }) // slot "rowCount"
}

func (this *QStandardItemModel) OnColumnCount(fn func(*QModelIndex) int) {
	registerVirt(this.UnsafeGoID(), -8256963024669584834, func(parent uintptr) int32 { return int32(fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)))) }) // slot "columnCount"
}

func (this *QStandardItemModel) OnHasChildren(fn func(*QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), 3369452259144514380, func(parent uintptr) int32 { return b2i(fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)))) }) // slot "hasChildren"
}

func (this *QStandardItemModel) OnData(fn func(*QModelIndex, int) *QVariant) {
	registerVirt(this.UnsafeGoID(), -8837375942098466299, func(index uintptr, role int32) uintptr { return uintptr(fn(UnsafeNewQModelIndex(unsafe.Pointer(index)), int(role)).UnsafePointer()) }) // slot "data"
}

func (this *QStandardItemModel) OnMultiData(fn func(*QModelIndex, *QModelRoleDataSpan)) {
	registerVirt(this.UnsafeGoID(), 731311060944096642, func(index uintptr, roleDataSpan uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(index)), UnsafeNewQModelRoleDataSpan(unsafe.Pointer(roleDataSpan))) }) // slot "multiData"
}

func (this *QStandardItemModel) OnSetData(fn func(*QModelIndex, *QVariant, int) bool) {
	registerVirt(this.UnsafeGoID(), 35191454575202219, func(index uintptr, value uintptr, role int32) int32 { return b2i(fn(UnsafeNewQModelIndex(unsafe.Pointer(index)), UnsafeNewQVariant(unsafe.Pointer(value)), int(role))) }) // slot "setData"
}

func (this *QStandardItemModel) OnClearItemData(fn func(*QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), 8914190500520542409, func(index uintptr) int32 { return b2i(fn(UnsafeNewQModelIndex(unsafe.Pointer(index)))) }) // slot "clearItemData"
}

func (this *QStandardItemModel) OnHeaderData(fn func(int, Orientation, int) *QVariant) {
	registerVirt(this.UnsafeGoID(), -3297140021896693688, func(section int32, orientation int32, role int32) uintptr { return uintptr(fn(int(section), Orientation(orientation), int(role)).UnsafePointer()) }) // slot "headerData"
}

func (this *QStandardItemModel) OnSetHeaderData(fn func(int, Orientation, *QVariant, int) bool) {
	registerVirt(this.UnsafeGoID(), 4852200306590054850, func(section int32, orientation int32, value uintptr, role int32) int32 { return b2i(fn(int(section), Orientation(orientation), UnsafeNewQVariant(unsafe.Pointer(value)), int(role))) }) // slot "setHeaderData"
}

func (this *QStandardItemModel) OnInsertRows(fn func(int, int, *QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), -2775729825962395603, func(row int32, count int32, parent uintptr) int32 { return b2i(fn(int(row), int(count), UnsafeNewQModelIndex(unsafe.Pointer(parent)))) }) // slot "insertRows"
}

func (this *QStandardItemModel) OnInsertColumns(fn func(int, int, *QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), -8973293846230792237, func(column int32, count int32, parent uintptr) int32 { return b2i(fn(int(column), int(count), UnsafeNewQModelIndex(unsafe.Pointer(parent)))) }) // slot "insertColumns"
}

func (this *QStandardItemModel) OnRemoveRows(fn func(int, int, *QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), 5104072549182222656, func(row int32, count int32, parent uintptr) int32 { return b2i(fn(int(row), int(count), UnsafeNewQModelIndex(unsafe.Pointer(parent)))) }) // slot "removeRows"
}

func (this *QStandardItemModel) OnRemoveColumns(fn func(int, int, *QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), 2078548737056540420, func(column int32, count int32, parent uintptr) int32 { return b2i(fn(int(column), int(count), UnsafeNewQModelIndex(unsafe.Pointer(parent)))) }) // slot "removeColumns"
}

func (this *QStandardItemModel) OnFlags(fn func(*QModelIndex) ItemFlag) {
	registerVirt(this.UnsafeGoID(), 1703382833545304812, func(index uintptr) int32 { return int32(fn(UnsafeNewQModelIndex(unsafe.Pointer(index)))) }) // slot "flags"
}

func (this *QStandardItemModel) OnSupportedDropActions(fn func() DropAction) {
	registerVirt(this.UnsafeGoID(), -7784642176186821205, func() int32 { return int32(fn()) }) // slot "supportedDropActions"
}

func (this *QStandardItemModel) OnSort(fn func(int, SortOrder)) {
	registerVirt(this.UnsafeGoID(), 6912441694453382161, func(column int32, order int32) { fn(int(column), SortOrder(order)) }) // slot "sort"
}

func (this *QStandardItemModel) OnMimeTypes(fn func() []string) {
	registerVirt(this.UnsafeGoID(), 5464467312394730418, func() (uintptr, int64) {
		__rs := fn()
		if len(__rs) == 0 {
			return 0, 0
		}
		__a := pgMalloc(uintptr(len(__rs)) * unsafe.Sizeof(uintptr(0)))
		__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__a)), len(__rs))
		for __i, __e := range __rs {
			__sl[__i] = pgCString(__e)
		}
		return __a, int64(len(__rs))
	}) // slot "mimeTypes"
}

func (this *QStandardItemModel) OnMimeData(fn func([]*QModelIndex) *QMimeData) {
	registerVirt(this.UnsafeGoID(), 5175297091024974749, func(indexes uintptr, indexes_len int64) uintptr { return uintptr(fn(func() []*QModelIndex { if indexes == 0 || indexes_len == 0 { return nil }; __sl := unsafe.Slice((*unsafe.Pointer)(unsafe.Pointer(indexes)), int(indexes_len)); __o := make([]*QModelIndex, int(indexes_len)); for __i := range __o { __o[__i] = UnsafeNewQModelIndex(__sl[__i]) }; return __o }()).UnsafePointer()) }) // slot "mimeData"
}

func (this *QStandardItemModel) OnDropMimeData(fn func(*QMimeData, DropAction, int, int, *QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), -9049994684305682212, func(data uintptr, action int32, row int32, column int32, parent uintptr) int32 { return b2i(fn(UnsafeNewQMimeData(unsafe.Pointer(data)), DropAction(action), int(row), int(column), UnsafeNewQModelIndex(unsafe.Pointer(parent)))) }) // slot "dropMimeData"
}

func (this *QStandardItemModel) OnSibling(fn func(int, int, *QModelIndex) *QModelIndex) {
	registerVirt(this.UnsafeGoID(), 397329346813782521, func(row int32, column int32, idx uintptr) uintptr { return uintptr(fn(int(row), int(column), UnsafeNewQModelIndex(unsafe.Pointer(idx))).UnsafePointer()) }) // slot "sibling"
}

func (this *QStandardItemModel) OnCanDropMimeData(fn func(*QMimeData, DropAction, int, int, *QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), 731539533644753052, func(data uintptr, action int32, row int32, column int32, parent uintptr) int32 { return b2i(fn(UnsafeNewQMimeData(unsafe.Pointer(data)), DropAction(action), int(row), int(column), UnsafeNewQModelIndex(unsafe.Pointer(parent)))) }) // slot "canDropMimeData"
}

func (this *QStandardItemModel) OnSupportedDragActions(fn func() DropAction) {
	registerVirt(this.UnsafeGoID(), 3421601386088570762, func() int32 { return int32(fn()) }) // slot "supportedDragActions"
}

func (this *QStandardItemModel) OnMoveRows(fn func(*QModelIndex, int, int, *QModelIndex, int) bool) {
	registerVirt(this.UnsafeGoID(), -6074819356947405415, func(sourceParent uintptr, sourceRow int32, count int32, destinationParent uintptr, destinationChild int32) int32 { return b2i(fn(UnsafeNewQModelIndex(unsafe.Pointer(sourceParent)), int(sourceRow), int(count), UnsafeNewQModelIndex(unsafe.Pointer(destinationParent)), int(destinationChild))) }) // slot "moveRows"
}

func (this *QStandardItemModel) OnMoveColumns(fn func(*QModelIndex, int, int, *QModelIndex, int) bool) {
	registerVirt(this.UnsafeGoID(), -6321007591388309905, func(sourceParent uintptr, sourceColumn int32, count int32, destinationParent uintptr, destinationChild int32) int32 { return b2i(fn(UnsafeNewQModelIndex(unsafe.Pointer(sourceParent)), int(sourceColumn), int(count), UnsafeNewQModelIndex(unsafe.Pointer(destinationParent)), int(destinationChild))) }) // slot "moveColumns"
}

func (this *QStandardItemModel) OnFetchMore(fn func(*QModelIndex)) {
	registerVirt(this.UnsafeGoID(), 3553294593611130806, func(parent uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent))) }) // slot "fetchMore"
}

func (this *QStandardItemModel) OnCanFetchMore(fn func(*QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), 8639262066473964150, func(parent uintptr) int32 { return b2i(fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)))) }) // slot "canFetchMore"
}

func (this *QStandardItemModel) OnBuddy(fn func(*QModelIndex) *QModelIndex) {
	registerVirt(this.UnsafeGoID(), 6112860112248786499, func(index uintptr) uintptr { return uintptr(fn(UnsafeNewQModelIndex(unsafe.Pointer(index))).UnsafePointer()) }) // slot "buddy"
}

func (this *QStandardItemModel) OnMatch(fn func(*QModelIndex, int, *QVariant, int, MatchFlag) []*QModelIndex) {
	registerVirt(this.UnsafeGoID(), -4341501217981328394, func(start uintptr, role int32, value uintptr, hits int32, flags int32) (uintptr, int64) {
		__rs := fn(UnsafeNewQModelIndex(unsafe.Pointer(start)), int(role), UnsafeNewQVariant(unsafe.Pointer(value)), int(hits), MatchFlag(flags))
		if len(__rs) == 0 {
			return 0, 0
		}
		__a := pgMalloc(uintptr(len(__rs)) * unsafe.Sizeof(uintptr(0)))
		__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__a)), len(__rs))
		for __i, __e := range __rs {
			__sl[__i] = uintptr(__e.UnsafePointer())
		}
		return __a, int64(len(__rs))
	}) // slot "match"
}

func (this *QStandardItemModel) OnSpan(fn func(*QModelIndex) *QSize) {
	registerVirt(this.UnsafeGoID(), -8395342916587040543, func(index uintptr) uintptr { return uintptr(fn(UnsafeNewQModelIndex(unsafe.Pointer(index))).UnsafePointer()) }) // slot "span"
}

func (this *QStandardItemModel) OnSubmit(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -3788331537442685179, func() int32 { return b2i(fn()) }) // slot "submit"
}

func (this *QStandardItemModel) OnRevert(fn func()) {
	registerVirt(this.UnsafeGoID(), -1522951607536117149, func() { fn() }) // slot "revert"
}

func (this *QStandardItemModel) OnResetInternalData(fn func()) {
	registerVirt(this.UnsafeGoID(), -4510675770784051183, func() { fn() }) // slot "resetInternalData"
}

func (this *QStandardItemModel) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QStandardItemModel) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QStandardItemModel) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QStandardItemModel) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QStandardItemModel) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QStandardItemModel) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QStandardItemModel) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QStandardItemModel) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QStandardItemModel_delete(this.h)
	}
	this.UnsafeRelease()
}
