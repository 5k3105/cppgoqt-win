package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

// purego: unsupported shapes skipped (map returns, >15-arg / >9-arg-callback, exotic): SetItemData

var (
	pg_QAbstractItemModel_base func(uintptr) uintptr
	addr_QAbstractItemModel_tr uintptr
	addr_QAbstractItemModel_hasIndex uintptr
	addr_QAbstractItemModel_index uintptr
	addr_QAbstractItemModel_parent uintptr
	addr_QAbstractItemModel_sibling uintptr
	addr_QAbstractItemModel_rowCount uintptr
	addr_QAbstractItemModel_columnCount uintptr
	addr_QAbstractItemModel_hasChildren uintptr
	addr_QAbstractItemModel_data uintptr
	addr_QAbstractItemModel_setData uintptr
	addr_QAbstractItemModel_headerData uintptr
	addr_QAbstractItemModel_setHeaderData uintptr
	addr_QAbstractItemModel_itemData uintptr
	addr_QAbstractItemModel_clearItemData uintptr
	addr_QAbstractItemModel_mimeTypes uintptr
	addr_QAbstractItemModel_mimeData uintptr
	addr_QAbstractItemModel_canDropMimeData uintptr
	addr_QAbstractItemModel_dropMimeData uintptr
	addr_QAbstractItemModel_supportedDropActions uintptr
	addr_QAbstractItemModel_supportedDragActions uintptr
	addr_QAbstractItemModel_insertRows uintptr
	addr_QAbstractItemModel_insertColumns uintptr
	addr_QAbstractItemModel_removeRows uintptr
	addr_QAbstractItemModel_removeColumns uintptr
	addr_QAbstractItemModel_moveRows uintptr
	addr_QAbstractItemModel_moveColumns uintptr
	addr_QAbstractItemModel_insertRow uintptr
	addr_QAbstractItemModel_insertColumn uintptr
	addr_QAbstractItemModel_removeRow uintptr
	addr_QAbstractItemModel_removeColumn uintptr
	addr_QAbstractItemModel_moveRow uintptr
	addr_QAbstractItemModel_moveColumn uintptr
	addr_QAbstractItemModel_fetchMore uintptr
	addr_QAbstractItemModel_canFetchMore uintptr
	addr_QAbstractItemModel_flags uintptr
	addr_QAbstractItemModel_sort uintptr
	addr_QAbstractItemModel_buddy uintptr
	addr_QAbstractItemModel_match uintptr
	addr_QAbstractItemModel_span uintptr
	addr_QAbstractItemModel_roleNames uintptr
	addr_QAbstractItemModel_checkIndex uintptr
	addr_QAbstractItemModel_multiData uintptr
	addr_QAbstractItemModel_dataChanged uintptr
	addr_QAbstractItemModel_headerDataChanged uintptr
	addr_QAbstractItemModel_layoutChanged uintptr
	addr_QAbstractItemModel_layoutAboutToBeChanged uintptr
	addr_QAbstractItemModel_submit uintptr
	addr_QAbstractItemModel_revert uintptr
	addr_QAbstractItemModel_createIndex uintptr
	addr_QAbstractItemModel_createIndex2 uintptr
	addr_QAbstractItemModel_encodeData uintptr
	addr_QAbstractItemModel_decodeData uintptr
	addr_QAbstractItemModel_beginInsertRows uintptr
	addr_QAbstractItemModel_endInsertRows uintptr
	addr_QAbstractItemModel_beginRemoveRows uintptr
	addr_QAbstractItemModel_endRemoveRows uintptr
	addr_QAbstractItemModel_beginMoveRows uintptr
	addr_QAbstractItemModel_endMoveRows uintptr
	addr_QAbstractItemModel_beginInsertColumns uintptr
	addr_QAbstractItemModel_endInsertColumns uintptr
	addr_QAbstractItemModel_beginRemoveColumns uintptr
	addr_QAbstractItemModel_endRemoveColumns uintptr
	addr_QAbstractItemModel_beginMoveColumns uintptr
	addr_QAbstractItemModel_endMoveColumns uintptr
	addr_QAbstractItemModel_beginResetModel uintptr
	addr_QAbstractItemModel_endResetModel uintptr
	addr_QAbstractItemModel_changePersistentIndex uintptr
	addr_QAbstractItemModel_changePersistentIndexList uintptr
	addr_QAbstractItemModel_persistentIndexList uintptr
	addr_QAbstractItemModel_connect uintptr
	addr_QAbstractItemModel_connect2 uintptr
	addr_QAbstractItemModel_sender uintptr
	addr_QAbstractItemModel_senderSignalIndex uintptr
	addr_QAbstractItemModel_receivers uintptr
	addr_QAbstractItemModel_isSignalConnected uintptr
	addr_QAbstractItemModel_tr2 uintptr
	addr_QAbstractItemModel_tr3 uintptr
	addr_QAbstractItemModel_hasIndex2 uintptr
	addr_QAbstractItemModel_insertRow2 uintptr
	addr_QAbstractItemModel_insertColumn2 uintptr
	addr_QAbstractItemModel_removeRow2 uintptr
	addr_QAbstractItemModel_removeColumn2 uintptr
	addr_QAbstractItemModel_checkIndex2 uintptr
	addr_QAbstractItemModel_dataChanged2 uintptr
	addr_QAbstractItemModel_layoutChanged2 uintptr
	addr_QAbstractItemModel_layoutChanged3 uintptr
	addr_QAbstractItemModel_layoutAboutToBeChanged2 uintptr
	addr_QAbstractItemModel_layoutAboutToBeChanged3 uintptr
	addr_QAbstractItemModel_connect3 uintptr
	addr_QAbstractItemModel_connect4 uintptr
	addr_QAbstractItemModel_sibling_default uintptr
	addr_QAbstractItemModel_hasChildren_default uintptr
	addr_QAbstractItemModel_setData_default uintptr
	addr_QAbstractItemModel_headerData_default uintptr
	addr_QAbstractItemModel_setHeaderData_default uintptr
	addr_QAbstractItemModel_clearItemData_default uintptr
	addr_QAbstractItemModel_mimeTypes_default uintptr
	addr_QAbstractItemModel_mimeData_default uintptr
	addr_QAbstractItemModel_canDropMimeData_default uintptr
	addr_QAbstractItemModel_dropMimeData_default uintptr
	addr_QAbstractItemModel_supportedDropActions_default uintptr
	addr_QAbstractItemModel_supportedDragActions_default uintptr
	addr_QAbstractItemModel_insertRows_default uintptr
	addr_QAbstractItemModel_insertColumns_default uintptr
	addr_QAbstractItemModel_removeRows_default uintptr
	addr_QAbstractItemModel_removeColumns_default uintptr
	addr_QAbstractItemModel_moveRows_default uintptr
	addr_QAbstractItemModel_moveColumns_default uintptr
	addr_QAbstractItemModel_fetchMore_default uintptr
	addr_QAbstractItemModel_canFetchMore_default uintptr
	addr_QAbstractItemModel_flags_default uintptr
	addr_QAbstractItemModel_sort_default uintptr
	addr_QAbstractItemModel_buddy_default uintptr
	addr_QAbstractItemModel_match_default uintptr
	addr_QAbstractItemModel_span_default uintptr
	addr_QAbstractItemModel_multiData_default uintptr
	addr_QAbstractItemModel_submit_default uintptr
	addr_QAbstractItemModel_revert_default uintptr
	addr_QAbstractItemModel_metaObject_default uintptr
	addr_QAbstractItemModel_resetInternalData_default uintptr
	addr_QAbstractItemModel_event_default uintptr
	addr_QAbstractItemModel_eventFilter_default uintptr
	addr_QAbstractItemModel_timerEvent_default uintptr
	addr_QAbstractItemModel_childEvent_default uintptr
	addr_QAbstractItemModel_customEvent_default uintptr
	addr_QAbstractItemModel_connectNotify_default uintptr
	addr_QAbstractItemModel_disconnectNotify_default uintptr
	pg_QAbstractItemModel_connect_headerDataChanged func(uintptr, int64)
	pg_QAbstractItemModel_connect_layoutChanged func(uintptr, int64)
	pg_QAbstractItemModel_connect_layoutChanged2 func(uintptr, int64)
	pg_QAbstractItemModel_connect_layoutAboutToBeChanged func(uintptr, int64)
	pg_QAbstractItemModel_connect_layoutAboutToBeChanged2 func(uintptr, int64)
	pg_QAbstractItemModel_connect_rowsAboutToBeInserted func(uintptr, int64)
	pg_QAbstractItemModel_connect_rowsInserted func(uintptr, int64)
	pg_QAbstractItemModel_connect_rowsAboutToBeRemoved func(uintptr, int64)
	pg_QAbstractItemModel_connect_rowsRemoved func(uintptr, int64)
	pg_QAbstractItemModel_connect_columnsAboutToBeInserted func(uintptr, int64)
	pg_QAbstractItemModel_connect_columnsInserted func(uintptr, int64)
	pg_QAbstractItemModel_connect_columnsAboutToBeRemoved func(uintptr, int64)
	pg_QAbstractItemModel_connect_columnsRemoved func(uintptr, int64)
	pg_QAbstractItemModel_connect_modelAboutToBeReset func(uintptr, int64)
	pg_QAbstractItemModel_connect_modelReset func(uintptr, int64)
	pg_QAbstractItemModel_connect_rowsAboutToBeMoved func(uintptr, int64)
	pg_QAbstractItemModel_connect_rowsMoved func(uintptr, int64)
	pg_QAbstractItemModel_connect_columnsAboutToBeMoved func(uintptr, int64)
	pg_QAbstractItemModel_connect_columnsMoved func(uintptr, int64)
	pg_QAbstractItemModel_connect_destroyed func(uintptr, int64)
	pg_QAbstractItemModel_connect_destroyed2 func(uintptr, int64)
	pg_QAbstractItemModel_connect_objectNameChanged func(uintptr, int64)
	pg_QAbstractItemModel__disconnect func(int64)
	pg_QAbstractItemModel__set_override func(uintptr, int32)
	pg_QAbstractItemModel__Go_new func(int64, uintptr) uintptr
	pg_QAbstractItemModel_delete func(uintptr)
)

func init() {
	reg(&pg_QAbstractItemModel_base, "QAbstractItemModel_base")
	symaddr(&addr_QAbstractItemModel_tr, "QAbstractItemModel_tr")
	symaddr(&addr_QAbstractItemModel_hasIndex, "QAbstractItemModel_hasIndex")
	symaddr(&addr_QAbstractItemModel_index, "QAbstractItemModel_index")
	symaddr(&addr_QAbstractItemModel_parent, "QAbstractItemModel_parent")
	symaddr(&addr_QAbstractItemModel_sibling, "QAbstractItemModel_sibling")
	symaddr(&addr_QAbstractItemModel_rowCount, "QAbstractItemModel_rowCount")
	symaddr(&addr_QAbstractItemModel_columnCount, "QAbstractItemModel_columnCount")
	symaddr(&addr_QAbstractItemModel_hasChildren, "QAbstractItemModel_hasChildren")
	symaddr(&addr_QAbstractItemModel_data, "QAbstractItemModel_data")
	symaddr(&addr_QAbstractItemModel_setData, "QAbstractItemModel_setData")
	symaddr(&addr_QAbstractItemModel_headerData, "QAbstractItemModel_headerData")
	symaddr(&addr_QAbstractItemModel_setHeaderData, "QAbstractItemModel_setHeaderData")
	symaddr(&addr_QAbstractItemModel_itemData, "QAbstractItemModel_itemData")
	symaddr(&addr_QAbstractItemModel_clearItemData, "QAbstractItemModel_clearItemData")
	symaddr(&addr_QAbstractItemModel_mimeTypes, "QAbstractItemModel_mimeTypes")
	symaddr(&addr_QAbstractItemModel_mimeData, "QAbstractItemModel_mimeData")
	symaddr(&addr_QAbstractItemModel_canDropMimeData, "QAbstractItemModel_canDropMimeData")
	symaddr(&addr_QAbstractItemModel_dropMimeData, "QAbstractItemModel_dropMimeData")
	symaddr(&addr_QAbstractItemModel_supportedDropActions, "QAbstractItemModel_supportedDropActions")
	symaddr(&addr_QAbstractItemModel_supportedDragActions, "QAbstractItemModel_supportedDragActions")
	symaddr(&addr_QAbstractItemModel_insertRows, "QAbstractItemModel_insertRows")
	symaddr(&addr_QAbstractItemModel_insertColumns, "QAbstractItemModel_insertColumns")
	symaddr(&addr_QAbstractItemModel_removeRows, "QAbstractItemModel_removeRows")
	symaddr(&addr_QAbstractItemModel_removeColumns, "QAbstractItemModel_removeColumns")
	symaddr(&addr_QAbstractItemModel_moveRows, "QAbstractItemModel_moveRows")
	symaddr(&addr_QAbstractItemModel_moveColumns, "QAbstractItemModel_moveColumns")
	symaddr(&addr_QAbstractItemModel_insertRow, "QAbstractItemModel_insertRow")
	symaddr(&addr_QAbstractItemModel_insertColumn, "QAbstractItemModel_insertColumn")
	symaddr(&addr_QAbstractItemModel_removeRow, "QAbstractItemModel_removeRow")
	symaddr(&addr_QAbstractItemModel_removeColumn, "QAbstractItemModel_removeColumn")
	symaddr(&addr_QAbstractItemModel_moveRow, "QAbstractItemModel_moveRow")
	symaddr(&addr_QAbstractItemModel_moveColumn, "QAbstractItemModel_moveColumn")
	symaddr(&addr_QAbstractItemModel_fetchMore, "QAbstractItemModel_fetchMore")
	symaddr(&addr_QAbstractItemModel_canFetchMore, "QAbstractItemModel_canFetchMore")
	symaddr(&addr_QAbstractItemModel_flags, "QAbstractItemModel_flags")
	symaddr(&addr_QAbstractItemModel_sort, "QAbstractItemModel_sort")
	symaddr(&addr_QAbstractItemModel_buddy, "QAbstractItemModel_buddy")
	symaddr(&addr_QAbstractItemModel_match, "QAbstractItemModel_match")
	symaddr(&addr_QAbstractItemModel_span, "QAbstractItemModel_span")
	symaddr(&addr_QAbstractItemModel_roleNames, "QAbstractItemModel_roleNames")
	symaddr(&addr_QAbstractItemModel_checkIndex, "QAbstractItemModel_checkIndex")
	symaddr(&addr_QAbstractItemModel_multiData, "QAbstractItemModel_multiData")
	symaddr(&addr_QAbstractItemModel_dataChanged, "QAbstractItemModel_dataChanged")
	symaddr(&addr_QAbstractItemModel_headerDataChanged, "QAbstractItemModel_headerDataChanged")
	symaddr(&addr_QAbstractItemModel_layoutChanged, "QAbstractItemModel_layoutChanged")
	symaddr(&addr_QAbstractItemModel_layoutAboutToBeChanged, "QAbstractItemModel_layoutAboutToBeChanged")
	symaddr(&addr_QAbstractItemModel_submit, "QAbstractItemModel_submit")
	symaddr(&addr_QAbstractItemModel_revert, "QAbstractItemModel_revert")
	symaddr(&addr_QAbstractItemModel_createIndex, "QAbstractItemModel_createIndex")
	symaddr(&addr_QAbstractItemModel_createIndex2, "QAbstractItemModel_createIndex2")
	symaddr(&addr_QAbstractItemModel_encodeData, "QAbstractItemModel_encodeData")
	symaddr(&addr_QAbstractItemModel_decodeData, "QAbstractItemModel_decodeData")
	symaddr(&addr_QAbstractItemModel_beginInsertRows, "QAbstractItemModel_beginInsertRows")
	symaddr(&addr_QAbstractItemModel_endInsertRows, "QAbstractItemModel_endInsertRows")
	symaddr(&addr_QAbstractItemModel_beginRemoveRows, "QAbstractItemModel_beginRemoveRows")
	symaddr(&addr_QAbstractItemModel_endRemoveRows, "QAbstractItemModel_endRemoveRows")
	symaddr(&addr_QAbstractItemModel_beginMoveRows, "QAbstractItemModel_beginMoveRows")
	symaddr(&addr_QAbstractItemModel_endMoveRows, "QAbstractItemModel_endMoveRows")
	symaddr(&addr_QAbstractItemModel_beginInsertColumns, "QAbstractItemModel_beginInsertColumns")
	symaddr(&addr_QAbstractItemModel_endInsertColumns, "QAbstractItemModel_endInsertColumns")
	symaddr(&addr_QAbstractItemModel_beginRemoveColumns, "QAbstractItemModel_beginRemoveColumns")
	symaddr(&addr_QAbstractItemModel_endRemoveColumns, "QAbstractItemModel_endRemoveColumns")
	symaddr(&addr_QAbstractItemModel_beginMoveColumns, "QAbstractItemModel_beginMoveColumns")
	symaddr(&addr_QAbstractItemModel_endMoveColumns, "QAbstractItemModel_endMoveColumns")
	symaddr(&addr_QAbstractItemModel_beginResetModel, "QAbstractItemModel_beginResetModel")
	symaddr(&addr_QAbstractItemModel_endResetModel, "QAbstractItemModel_endResetModel")
	symaddr(&addr_QAbstractItemModel_changePersistentIndex, "QAbstractItemModel_changePersistentIndex")
	symaddr(&addr_QAbstractItemModel_changePersistentIndexList, "QAbstractItemModel_changePersistentIndexList")
	symaddr(&addr_QAbstractItemModel_persistentIndexList, "QAbstractItemModel_persistentIndexList")
	symaddr(&addr_QAbstractItemModel_connect, "QAbstractItemModel_connect")
	symaddr(&addr_QAbstractItemModel_connect2, "QAbstractItemModel_connect2")
	symaddr(&addr_QAbstractItemModel_sender, "QAbstractItemModel_sender")
	symaddr(&addr_QAbstractItemModel_senderSignalIndex, "QAbstractItemModel_senderSignalIndex")
	symaddr(&addr_QAbstractItemModel_receivers, "QAbstractItemModel_receivers")
	symaddr(&addr_QAbstractItemModel_isSignalConnected, "QAbstractItemModel_isSignalConnected")
	symaddr(&addr_QAbstractItemModel_tr2, "QAbstractItemModel_tr2")
	symaddr(&addr_QAbstractItemModel_tr3, "QAbstractItemModel_tr3")
	symaddr(&addr_QAbstractItemModel_hasIndex2, "QAbstractItemModel_hasIndex2")
	symaddr(&addr_QAbstractItemModel_insertRow2, "QAbstractItemModel_insertRow2")
	symaddr(&addr_QAbstractItemModel_insertColumn2, "QAbstractItemModel_insertColumn2")
	symaddr(&addr_QAbstractItemModel_removeRow2, "QAbstractItemModel_removeRow2")
	symaddr(&addr_QAbstractItemModel_removeColumn2, "QAbstractItemModel_removeColumn2")
	symaddr(&addr_QAbstractItemModel_checkIndex2, "QAbstractItemModel_checkIndex2")
	symaddr(&addr_QAbstractItemModel_dataChanged2, "QAbstractItemModel_dataChanged2")
	symaddr(&addr_QAbstractItemModel_layoutChanged2, "QAbstractItemModel_layoutChanged2")
	symaddr(&addr_QAbstractItemModel_layoutChanged3, "QAbstractItemModel_layoutChanged3")
	symaddr(&addr_QAbstractItemModel_layoutAboutToBeChanged2, "QAbstractItemModel_layoutAboutToBeChanged2")
	symaddr(&addr_QAbstractItemModel_layoutAboutToBeChanged3, "QAbstractItemModel_layoutAboutToBeChanged3")
	symaddr(&addr_QAbstractItemModel_connect3, "QAbstractItemModel_connect3")
	symaddr(&addr_QAbstractItemModel_connect4, "QAbstractItemModel_connect4")
	symaddr(&addr_QAbstractItemModel_sibling_default, "QAbstractItemModel_sibling_default")
	symaddr(&addr_QAbstractItemModel_hasChildren_default, "QAbstractItemModel_hasChildren_default")
	symaddr(&addr_QAbstractItemModel_setData_default, "QAbstractItemModel_setData_default")
	symaddr(&addr_QAbstractItemModel_headerData_default, "QAbstractItemModel_headerData_default")
	symaddr(&addr_QAbstractItemModel_setHeaderData_default, "QAbstractItemModel_setHeaderData_default")
	symaddr(&addr_QAbstractItemModel_clearItemData_default, "QAbstractItemModel_clearItemData_default")
	symaddr(&addr_QAbstractItemModel_mimeTypes_default, "QAbstractItemModel_mimeTypes_default")
	symaddr(&addr_QAbstractItemModel_mimeData_default, "QAbstractItemModel_mimeData_default")
	symaddr(&addr_QAbstractItemModel_canDropMimeData_default, "QAbstractItemModel_canDropMimeData_default")
	symaddr(&addr_QAbstractItemModel_dropMimeData_default, "QAbstractItemModel_dropMimeData_default")
	symaddr(&addr_QAbstractItemModel_supportedDropActions_default, "QAbstractItemModel_supportedDropActions_default")
	symaddr(&addr_QAbstractItemModel_supportedDragActions_default, "QAbstractItemModel_supportedDragActions_default")
	symaddr(&addr_QAbstractItemModel_insertRows_default, "QAbstractItemModel_insertRows_default")
	symaddr(&addr_QAbstractItemModel_insertColumns_default, "QAbstractItemModel_insertColumns_default")
	symaddr(&addr_QAbstractItemModel_removeRows_default, "QAbstractItemModel_removeRows_default")
	symaddr(&addr_QAbstractItemModel_removeColumns_default, "QAbstractItemModel_removeColumns_default")
	symaddr(&addr_QAbstractItemModel_moveRows_default, "QAbstractItemModel_moveRows_default")
	symaddr(&addr_QAbstractItemModel_moveColumns_default, "QAbstractItemModel_moveColumns_default")
	symaddr(&addr_QAbstractItemModel_fetchMore_default, "QAbstractItemModel_fetchMore_default")
	symaddr(&addr_QAbstractItemModel_canFetchMore_default, "QAbstractItemModel_canFetchMore_default")
	symaddr(&addr_QAbstractItemModel_flags_default, "QAbstractItemModel_flags_default")
	symaddr(&addr_QAbstractItemModel_sort_default, "QAbstractItemModel_sort_default")
	symaddr(&addr_QAbstractItemModel_buddy_default, "QAbstractItemModel_buddy_default")
	symaddr(&addr_QAbstractItemModel_match_default, "QAbstractItemModel_match_default")
	symaddr(&addr_QAbstractItemModel_span_default, "QAbstractItemModel_span_default")
	symaddr(&addr_QAbstractItemModel_multiData_default, "QAbstractItemModel_multiData_default")
	symaddr(&addr_QAbstractItemModel_submit_default, "QAbstractItemModel_submit_default")
	symaddr(&addr_QAbstractItemModel_revert_default, "QAbstractItemModel_revert_default")
	symaddr(&addr_QAbstractItemModel_metaObject_default, "QAbstractItemModel_metaObject_default")
	symaddr(&addr_QAbstractItemModel_resetInternalData_default, "QAbstractItemModel_resetInternalData_default")
	symaddr(&addr_QAbstractItemModel_event_default, "QAbstractItemModel_event_default")
	symaddr(&addr_QAbstractItemModel_eventFilter_default, "QAbstractItemModel_eventFilter_default")
	symaddr(&addr_QAbstractItemModel_timerEvent_default, "QAbstractItemModel_timerEvent_default")
	symaddr(&addr_QAbstractItemModel_childEvent_default, "QAbstractItemModel_childEvent_default")
	symaddr(&addr_QAbstractItemModel_customEvent_default, "QAbstractItemModel_customEvent_default")
	symaddr(&addr_QAbstractItemModel_connectNotify_default, "QAbstractItemModel_connectNotify_default")
	symaddr(&addr_QAbstractItemModel_disconnectNotify_default, "QAbstractItemModel_disconnectNotify_default")
	reg(&pg_QAbstractItemModel_connect_headerDataChanged, "QAbstractItemModel_connect_headerDataChanged")
	reg(&pg_QAbstractItemModel_connect_layoutChanged, "QAbstractItemModel_connect_layoutChanged")
	reg(&pg_QAbstractItemModel_connect_layoutChanged2, "QAbstractItemModel_connect_layoutChanged2")
	reg(&pg_QAbstractItemModel_connect_layoutAboutToBeChanged, "QAbstractItemModel_connect_layoutAboutToBeChanged")
	reg(&pg_QAbstractItemModel_connect_layoutAboutToBeChanged2, "QAbstractItemModel_connect_layoutAboutToBeChanged2")
	reg(&pg_QAbstractItemModel_connect_rowsAboutToBeInserted, "QAbstractItemModel_connect_rowsAboutToBeInserted")
	reg(&pg_QAbstractItemModel_connect_rowsInserted, "QAbstractItemModel_connect_rowsInserted")
	reg(&pg_QAbstractItemModel_connect_rowsAboutToBeRemoved, "QAbstractItemModel_connect_rowsAboutToBeRemoved")
	reg(&pg_QAbstractItemModel_connect_rowsRemoved, "QAbstractItemModel_connect_rowsRemoved")
	reg(&pg_QAbstractItemModel_connect_columnsAboutToBeInserted, "QAbstractItemModel_connect_columnsAboutToBeInserted")
	reg(&pg_QAbstractItemModel_connect_columnsInserted, "QAbstractItemModel_connect_columnsInserted")
	reg(&pg_QAbstractItemModel_connect_columnsAboutToBeRemoved, "QAbstractItemModel_connect_columnsAboutToBeRemoved")
	reg(&pg_QAbstractItemModel_connect_columnsRemoved, "QAbstractItemModel_connect_columnsRemoved")
	reg(&pg_QAbstractItemModel_connect_modelAboutToBeReset, "QAbstractItemModel_connect_modelAboutToBeReset")
	reg(&pg_QAbstractItemModel_connect_modelReset, "QAbstractItemModel_connect_modelReset")
	reg(&pg_QAbstractItemModel_connect_rowsAboutToBeMoved, "QAbstractItemModel_connect_rowsAboutToBeMoved")
	reg(&pg_QAbstractItemModel_connect_rowsMoved, "QAbstractItemModel_connect_rowsMoved")
	reg(&pg_QAbstractItemModel_connect_columnsAboutToBeMoved, "QAbstractItemModel_connect_columnsAboutToBeMoved")
	reg(&pg_QAbstractItemModel_connect_columnsMoved, "QAbstractItemModel_connect_columnsMoved")
	reg(&pg_QAbstractItemModel_connect_destroyed, "QAbstractItemModel_connect_destroyed")
	reg(&pg_QAbstractItemModel_connect_destroyed2, "QAbstractItemModel_connect_destroyed2")
	reg(&pg_QAbstractItemModel_connect_objectNameChanged, "QAbstractItemModel_connect_objectNameChanged")
	reg(&pg_QAbstractItemModel__disconnect, "QAbstractItemModel__disconnect")
	reg(&pg_QAbstractItemModel__set_override, "QAbstractItemModel__set_override")
	reg(&pg_QAbstractItemModel__Go_new, "QAbstractItemModel__Go_new")
	reg(&pg_QAbstractItemModel_delete, "QAbstractItemModel_delete")
}

type QAbstractItemModel struct {
	h uintptr
	*QObject
}

func bindQAbstractItemModel(h uintptr) *QAbstractItemModel {
	if h == 0 {
		return nil
	}
	return &QAbstractItemModel{h: h, QObject: bindQObject(pg_QAbstractItemModel_base(h))}
}

func newownedQAbstractItemModel(h uintptr) *QAbstractItemModel { return bindQAbstractItemModel(h) }

func (this *QAbstractItemModel) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QAbstractItemModel) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQAbstractItemModel(h unsafe.Pointer) *QAbstractItemModel { return bindQAbstractItemModel(uintptr(h)) }
func UnsafeNewQAbstractItemModelOwned(h unsafe.Pointer) *QAbstractItemModel { return newownedQAbstractItemModel(uintptr(h)) }

func (this *QAbstractItemModel) UnsafeRelease() {
	this.h = 0
	if this.QObject != nil {
		this.QObject.UnsafeRelease()
	}
}

func QAbstractItemModel_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAbstractItemModel_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QAbstractItemModel) HasIndex(row int, column int) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemModel_hasIndex, this.h, uintptr(row), uintptr(column))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QAbstractItemModel) Index(row int, column int, parent *QModelIndex) *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QAbstractItemModel_index, this.h, uintptr(row), uintptr(column), uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return newownedQModelIndex(__r)
}

func (this *QAbstractItemModel) Parent(child *QModelIndex) *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QAbstractItemModel_parent, this.h, uintptr(child.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(child)
	return newownedQModelIndex(__r)
}

func (this *QAbstractItemModel) Sibling(row int, column int, idx *QModelIndex) *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QAbstractItemModel_sibling, this.h, uintptr(row), uintptr(column), uintptr(idx.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(idx)
	return newownedQModelIndex(__r)
}

func (this *QAbstractItemModel) RowCount(parent *QModelIndex) int {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemModel_rowCount, this.h, uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return int(__r)
}

func (this *QAbstractItemModel) ColumnCount(parent *QModelIndex) int {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemModel_columnCount, this.h, uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return int(__r)
}

func (this *QAbstractItemModel) HasChildren(parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemModel_hasChildren, this.h, uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QAbstractItemModel) Data(index *QModelIndex, role int) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QAbstractItemModel_data, this.h, uintptr(index.UnsafePointer()), uintptr(role))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return newownedQVariant(__r)
}

func (this *QAbstractItemModel) SetData(index *QModelIndex, value *QVariant, role int) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemModel_setData, this.h, uintptr(index.UnsafePointer()), uintptr(value.UnsafePointer()), uintptr(role))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	runtime.KeepAlive(value)
	return __r != 0
}

func (this *QAbstractItemModel) HeaderData(section int, orientation Orientation, role int) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QAbstractItemModel_headerData, this.h, uintptr(section), uintptr(orientation), uintptr(role))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QAbstractItemModel) SetHeaderData(section int, orientation Orientation, value *QVariant, role int) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemModel_setHeaderData, this.h, uintptr(section), uintptr(orientation), uintptr(value.UnsafePointer()), uintptr(role))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(value)
	return __r != 0
}

func (this *QAbstractItemModel) ItemData(index *QModelIndex) map[int]*QVariant {
	var __vp uintptr
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAbstractItemModel_itemData, this.h, uintptr(index.UnsafePointer()), uintptr(unsafe.Pointer(&__vp)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	__ks := unsafe.Slice((*int32)(unsafe.Pointer(__r)), int(__n))
	__vs := unsafe.Slice((*uintptr)(unsafe.Pointer(__vp)), int(__n))
	__out := make(map[int]*QVariant, int(__n))
	for __i := 0; __i < int(__n); __i++ {
		__out[int(__ks[__i])] = newownedQVariant(__vs[__i])
	}
	pgFree(__r)
	pgFree(__vp)
	return __out
}

func (this *QAbstractItemModel) ClearItemData(index *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemModel_clearItemData, this.h, uintptr(index.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return __r != 0
}

func (this *QAbstractItemModel) MimeTypes() []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAbstractItemModel_mimeTypes, this.h, uintptr(unsafe.Pointer(&__n)))
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

func (this *QAbstractItemModel) MimeData(indexes []*QModelIndex) *QMimeData {
	indexes_arr := make([]uintptr, len(indexes))
	for __i, __e := range indexes {
		indexes_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var indexes_p uintptr
	if len(indexes_arr) > 0 {
		indexes_p = uintptr(unsafe.Pointer(&indexes_arr[0]))
	}
	__r, _, _ := purego.SyscallN(addr_QAbstractItemModel_mimeData, this.h, indexes_p, uintptr(len(indexes)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(indexes)
	runtime.KeepAlive(indexes_arr)
	return bindQMimeData(__r)
}

func (this *QAbstractItemModel) CanDropMimeData(data *QMimeData, action DropAction, row int, column int, parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemModel_canDropMimeData, this.h, uintptr(data.UnsafePointer()), uintptr(action), uintptr(row), uintptr(column), uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(data)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QAbstractItemModel) DropMimeData(data *QMimeData, action DropAction, row int, column int, parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemModel_dropMimeData, this.h, uintptr(data.UnsafePointer()), uintptr(action), uintptr(row), uintptr(column), uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(data)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QAbstractItemModel) SupportedDropActions() DropAction {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemModel_supportedDropActions, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return DropAction(__r)
}

func (this *QAbstractItemModel) SupportedDragActions() DropAction {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemModel_supportedDragActions, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return DropAction(__r)
}

func (this *QAbstractItemModel) InsertRows(row int, count int, parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemModel_insertRows, this.h, uintptr(row), uintptr(count), uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QAbstractItemModel) InsertColumns(column int, count int, parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemModel_insertColumns, this.h, uintptr(column), uintptr(count), uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QAbstractItemModel) RemoveRows(row int, count int, parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemModel_removeRows, this.h, uintptr(row), uintptr(count), uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QAbstractItemModel) RemoveColumns(column int, count int, parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemModel_removeColumns, this.h, uintptr(column), uintptr(count), uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QAbstractItemModel) MoveRows(sourceParent *QModelIndex, sourceRow int, count int, destinationParent *QModelIndex, destinationChild int) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemModel_moveRows, this.h, uintptr(sourceParent.UnsafePointer()), uintptr(sourceRow), uintptr(count), uintptr(destinationParent.UnsafePointer()), uintptr(destinationChild))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sourceParent)
	runtime.KeepAlive(destinationParent)
	return __r != 0
}

func (this *QAbstractItemModel) MoveColumns(sourceParent *QModelIndex, sourceColumn int, count int, destinationParent *QModelIndex, destinationChild int) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemModel_moveColumns, this.h, uintptr(sourceParent.UnsafePointer()), uintptr(sourceColumn), uintptr(count), uintptr(destinationParent.UnsafePointer()), uintptr(destinationChild))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sourceParent)
	runtime.KeepAlive(destinationParent)
	return __r != 0
}

func (this *QAbstractItemModel) InsertRow(row int) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemModel_insertRow, this.h, uintptr(row))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QAbstractItemModel) InsertColumn(column int) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemModel_insertColumn, this.h, uintptr(column))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QAbstractItemModel) RemoveRow(row int) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemModel_removeRow, this.h, uintptr(row))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QAbstractItemModel) RemoveColumn(column int) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemModel_removeColumn, this.h, uintptr(column))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QAbstractItemModel) MoveRow(sourceParent *QModelIndex, sourceRow int, destinationParent *QModelIndex, destinationChild int) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemModel_moveRow, this.h, uintptr(sourceParent.UnsafePointer()), uintptr(sourceRow), uintptr(destinationParent.UnsafePointer()), uintptr(destinationChild))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sourceParent)
	runtime.KeepAlive(destinationParent)
	return __r != 0
}

func (this *QAbstractItemModel) MoveColumn(sourceParent *QModelIndex, sourceColumn int, destinationParent *QModelIndex, destinationChild int) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemModel_moveColumn, this.h, uintptr(sourceParent.UnsafePointer()), uintptr(sourceColumn), uintptr(destinationParent.UnsafePointer()), uintptr(destinationChild))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sourceParent)
	runtime.KeepAlive(destinationParent)
	return __r != 0
}

func (this *QAbstractItemModel) FetchMore(parent *QModelIndex) {
	purego.SyscallN(addr_QAbstractItemModel_fetchMore, this.h, uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
}

func (this *QAbstractItemModel) CanFetchMore(parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemModel_canFetchMore, this.h, uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QAbstractItemModel) Flags(index *QModelIndex) ItemFlag {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemModel_flags, this.h, uintptr(index.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return ItemFlag(__r)
}

func (this *QAbstractItemModel) Sort(column int, order SortOrder) {
	purego.SyscallN(addr_QAbstractItemModel_sort, this.h, uintptr(column), uintptr(order))
	runtime.KeepAlive(this)
}

func (this *QAbstractItemModel) Buddy(index *QModelIndex) *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QAbstractItemModel_buddy, this.h, uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return newownedQModelIndex(__r)
}

func (this *QAbstractItemModel) Match(start *QModelIndex, role int, value *QVariant, hits int, flags MatchFlag) []*QModelIndex {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAbstractItemModel_match, this.h, uintptr(start.UnsafePointer()), uintptr(role), uintptr(value.UnsafePointer()), uintptr(hits), uintptr(flags), uintptr(unsafe.Pointer(&__n)))
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

func (this *QAbstractItemModel) Span(index *QModelIndex) *QSize {
	__r, _, _ := purego.SyscallN(addr_QAbstractItemModel_span, this.h, uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return newownedQSize(__r)
}

func (this *QAbstractItemModel) RoleNames() map[int]*QByteArray {
	var __vp uintptr
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAbstractItemModel_roleNames, this.h, uintptr(unsafe.Pointer(&__vp)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__ks := unsafe.Slice((*int32)(unsafe.Pointer(__r)), int(__n))
	__vs := unsafe.Slice((*uintptr)(unsafe.Pointer(__vp)), int(__n))
	__out := make(map[int]*QByteArray, int(__n))
	for __i := 0; __i < int(__n); __i++ {
		__out[int(__ks[__i])] = newownedQByteArray(__vs[__i])
	}
	pgFree(__r)
	pgFree(__vp)
	return __out
}

func (this *QAbstractItemModel) CheckIndex(index *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemModel_checkIndex, this.h, uintptr(index.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return __r != 0
}

func (this *QAbstractItemModel) MultiData(index *QModelIndex, roleDataSpan *QModelRoleDataSpan) {
	purego.SyscallN(addr_QAbstractItemModel_multiData, this.h, uintptr(index.UnsafePointer()), uintptr(roleDataSpan.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	runtime.KeepAlive(roleDataSpan)
}

func (this *QAbstractItemModel) DataChanged(topLeft *QModelIndex, bottomRight *QModelIndex) {
	purego.SyscallN(addr_QAbstractItemModel_dataChanged, this.h, uintptr(topLeft.UnsafePointer()), uintptr(bottomRight.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(topLeft)
	runtime.KeepAlive(bottomRight)
}

func (this *QAbstractItemModel) HeaderDataChanged(orientation Orientation, first int, last int) {
	purego.SyscallN(addr_QAbstractItemModel_headerDataChanged, this.h, uintptr(orientation), uintptr(first), uintptr(last))
	runtime.KeepAlive(this)
}

func (this *QAbstractItemModel) LayoutChanged() {
	purego.SyscallN(addr_QAbstractItemModel_layoutChanged, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractItemModel) LayoutAboutToBeChanged() {
	purego.SyscallN(addr_QAbstractItemModel_layoutAboutToBeChanged, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractItemModel) Submit() bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemModel_submit, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QAbstractItemModel) Revert() {
	purego.SyscallN(addr_QAbstractItemModel_revert, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractItemModel) CreateIndex(row int, column int) *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QAbstractItemModel_createIndex, this.h, uintptr(row), uintptr(column))
	runtime.KeepAlive(this)
	return newownedQModelIndex(__r)
}

func (this *QAbstractItemModel) CreateIndex2(row int, column int, id uintptr) *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QAbstractItemModel_createIndex2, this.h, uintptr(row), uintptr(column), uintptr(id))
	runtime.KeepAlive(this)
	return newownedQModelIndex(__r)
}

func (this *QAbstractItemModel) EncodeData(indexes []*QModelIndex, stream *QDataStream) {
	indexes_arr := make([]uintptr, len(indexes))
	for __i, __e := range indexes {
		indexes_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var indexes_p uintptr
	if len(indexes_arr) > 0 {
		indexes_p = uintptr(unsafe.Pointer(&indexes_arr[0]))
	}
	purego.SyscallN(addr_QAbstractItemModel_encodeData, this.h, indexes_p, uintptr(len(indexes)), uintptr(stream.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(indexes)
	runtime.KeepAlive(indexes_arr)
	runtime.KeepAlive(stream)
}

func (this *QAbstractItemModel) DecodeData(row int, column int, parent *QModelIndex, stream *QDataStream) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemModel_decodeData, this.h, uintptr(row), uintptr(column), uintptr(parent.UnsafePointer()), uintptr(stream.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	runtime.KeepAlive(stream)
	return __r != 0
}

func (this *QAbstractItemModel) BeginInsertRows(parent *QModelIndex, first int, last int) {
	purego.SyscallN(addr_QAbstractItemModel_beginInsertRows, this.h, uintptr(parent.UnsafePointer()), uintptr(first), uintptr(last))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
}

func (this *QAbstractItemModel) EndInsertRows() {
	purego.SyscallN(addr_QAbstractItemModel_endInsertRows, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractItemModel) BeginRemoveRows(parent *QModelIndex, first int, last int) {
	purego.SyscallN(addr_QAbstractItemModel_beginRemoveRows, this.h, uintptr(parent.UnsafePointer()), uintptr(first), uintptr(last))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
}

func (this *QAbstractItemModel) EndRemoveRows() {
	purego.SyscallN(addr_QAbstractItemModel_endRemoveRows, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractItemModel) BeginMoveRows(sourceParent *QModelIndex, sourceFirst int, sourceLast int, destinationParent *QModelIndex, destinationRow int) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemModel_beginMoveRows, this.h, uintptr(sourceParent.UnsafePointer()), uintptr(sourceFirst), uintptr(sourceLast), uintptr(destinationParent.UnsafePointer()), uintptr(destinationRow))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sourceParent)
	runtime.KeepAlive(destinationParent)
	return __r != 0
}

func (this *QAbstractItemModel) EndMoveRows() {
	purego.SyscallN(addr_QAbstractItemModel_endMoveRows, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractItemModel) BeginInsertColumns(parent *QModelIndex, first int, last int) {
	purego.SyscallN(addr_QAbstractItemModel_beginInsertColumns, this.h, uintptr(parent.UnsafePointer()), uintptr(first), uintptr(last))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
}

func (this *QAbstractItemModel) EndInsertColumns() {
	purego.SyscallN(addr_QAbstractItemModel_endInsertColumns, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractItemModel) BeginRemoveColumns(parent *QModelIndex, first int, last int) {
	purego.SyscallN(addr_QAbstractItemModel_beginRemoveColumns, this.h, uintptr(parent.UnsafePointer()), uintptr(first), uintptr(last))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
}

func (this *QAbstractItemModel) EndRemoveColumns() {
	purego.SyscallN(addr_QAbstractItemModel_endRemoveColumns, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractItemModel) BeginMoveColumns(sourceParent *QModelIndex, sourceFirst int, sourceLast int, destinationParent *QModelIndex, destinationColumn int) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemModel_beginMoveColumns, this.h, uintptr(sourceParent.UnsafePointer()), uintptr(sourceFirst), uintptr(sourceLast), uintptr(destinationParent.UnsafePointer()), uintptr(destinationColumn))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sourceParent)
	runtime.KeepAlive(destinationParent)
	return __r != 0
}

func (this *QAbstractItemModel) EndMoveColumns() {
	purego.SyscallN(addr_QAbstractItemModel_endMoveColumns, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractItemModel) BeginResetModel() {
	purego.SyscallN(addr_QAbstractItemModel_beginResetModel, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractItemModel) EndResetModel() {
	purego.SyscallN(addr_QAbstractItemModel_endResetModel, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractItemModel) ChangePersistentIndex(from *QModelIndex, to *QModelIndex) {
	purego.SyscallN(addr_QAbstractItemModel_changePersistentIndex, this.h, uintptr(from.UnsafePointer()), uintptr(to.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(from)
	runtime.KeepAlive(to)
}

func (this *QAbstractItemModel) ChangePersistentIndexList(from []*QModelIndex, to []*QModelIndex) {
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
	purego.SyscallN(addr_QAbstractItemModel_changePersistentIndexList, this.h, from_p, uintptr(len(from)), to_p, uintptr(len(to)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(from)
	runtime.KeepAlive(from_arr)
	runtime.KeepAlive(to)
	runtime.KeepAlive(to_arr)
}

func (this *QAbstractItemModel) PersistentIndexList() []*QModelIndex {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAbstractItemModel_persistentIndexList, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QModelIndex, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQModelIndex(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QAbstractItemModel_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QAbstractItemModel_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QAbstractItemModel_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QAbstractItemModel_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QAbstractItemModel) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QAbstractItemModel_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QAbstractItemModel) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemModel_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAbstractItemModel) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemModel_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QAbstractItemModel) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemModel_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QAbstractItemModel_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAbstractItemModel_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QAbstractItemModel_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAbstractItemModel_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func (this *QAbstractItemModel) HasIndex2(row int, column int, parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemModel_hasIndex2, this.h, uintptr(row), uintptr(column), uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QAbstractItemModel) InsertRow2(row int, parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemModel_insertRow2, this.h, uintptr(row), uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QAbstractItemModel) InsertColumn2(column int, parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemModel_insertColumn2, this.h, uintptr(column), uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QAbstractItemModel) RemoveRow2(row int, parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemModel_removeRow2, this.h, uintptr(row), uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QAbstractItemModel) RemoveColumn2(column int, parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemModel_removeColumn2, this.h, uintptr(column), uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QAbstractItemModel) CheckIndex2(index *QModelIndex, options QAbstractItemModel__CheckIndexOption) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemModel_checkIndex2, this.h, uintptr(index.UnsafePointer()), uintptr(options))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return __r != 0
}

func (this *QAbstractItemModel) DataChanged2(topLeft *QModelIndex, bottomRight *QModelIndex, roles []int) {
	roles_arr := make([]int32, len(roles))
	for __i, __e := range roles {
		roles_arr[__i] = int32(__e)
	}
	var roles_p uintptr
	if len(roles_arr) > 0 {
		roles_p = uintptr(unsafe.Pointer(&roles_arr[0]))
	}
	purego.SyscallN(addr_QAbstractItemModel_dataChanged2, this.h, uintptr(topLeft.UnsafePointer()), uintptr(bottomRight.UnsafePointer()), roles_p, uintptr(len(roles)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(topLeft)
	runtime.KeepAlive(bottomRight)
	runtime.KeepAlive(roles_arr)
}

func (this *QAbstractItemModel) LayoutChangedWithParents(parents []*QPersistentModelIndex) {
	parents_arr := make([]uintptr, len(parents))
	for __i, __e := range parents {
		parents_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var parents_p uintptr
	if len(parents_arr) > 0 {
		parents_p = uintptr(unsafe.Pointer(&parents_arr[0]))
	}
	purego.SyscallN(addr_QAbstractItemModel_layoutChanged2, this.h, parents_p, uintptr(len(parents)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parents)
	runtime.KeepAlive(parents_arr)
}

func (this *QAbstractItemModel) LayoutChanged2(parents []*QPersistentModelIndex, hint QAbstractItemModel__LayoutChangeHint) {
	parents_arr := make([]uintptr, len(parents))
	for __i, __e := range parents {
		parents_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var parents_p uintptr
	if len(parents_arr) > 0 {
		parents_p = uintptr(unsafe.Pointer(&parents_arr[0]))
	}
	purego.SyscallN(addr_QAbstractItemModel_layoutChanged3, this.h, parents_p, uintptr(len(parents)), uintptr(hint))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parents)
	runtime.KeepAlive(parents_arr)
}

func (this *QAbstractItemModel) LayoutAboutToBeChangedWithParents(parents []*QPersistentModelIndex) {
	parents_arr := make([]uintptr, len(parents))
	for __i, __e := range parents {
		parents_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var parents_p uintptr
	if len(parents_arr) > 0 {
		parents_p = uintptr(unsafe.Pointer(&parents_arr[0]))
	}
	purego.SyscallN(addr_QAbstractItemModel_layoutAboutToBeChanged2, this.h, parents_p, uintptr(len(parents)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parents)
	runtime.KeepAlive(parents_arr)
}

func (this *QAbstractItemModel) LayoutAboutToBeChanged2(parents []*QPersistentModelIndex, hint QAbstractItemModel__LayoutChangeHint) {
	parents_arr := make([]uintptr, len(parents))
	for __i, __e := range parents {
		parents_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var parents_p uintptr
	if len(parents_arr) > 0 {
		parents_p = uintptr(unsafe.Pointer(&parents_arr[0]))
	}
	purego.SyscallN(addr_QAbstractItemModel_layoutAboutToBeChanged3, this.h, parents_p, uintptr(len(parents)), uintptr(hint))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parents)
	runtime.KeepAlive(parents_arr)
}

func QAbstractItemModel_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QAbstractItemModel_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QAbstractItemModel_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QAbstractItemModel_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QAbstractItemModel) SiblingDefault(row int, column int, idx *QModelIndex) *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QAbstractItemModel_sibling_default, this.h, uintptr(row), uintptr(column), uintptr(idx.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(idx)
	return newownedQModelIndex(__r)
}

func (this *QAbstractItemModel) HasChildrenDefault(parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemModel_hasChildren_default, this.h, uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QAbstractItemModel) SetDataDefault(index *QModelIndex, value *QVariant, role int) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemModel_setData_default, this.h, uintptr(index.UnsafePointer()), uintptr(value.UnsafePointer()), uintptr(role))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	runtime.KeepAlive(value)
	return __r != 0
}

func (this *QAbstractItemModel) HeaderDataDefault(section int, orientation Orientation, role int) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QAbstractItemModel_headerData_default, this.h, uintptr(section), uintptr(orientation), uintptr(role))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QAbstractItemModel) SetHeaderDataDefault(section int, orientation Orientation, value *QVariant, role int) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemModel_setHeaderData_default, this.h, uintptr(section), uintptr(orientation), uintptr(value.UnsafePointer()), uintptr(role))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(value)
	return __r != 0
}

func (this *QAbstractItemModel) ClearItemDataDefault(index *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemModel_clearItemData_default, this.h, uintptr(index.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return __r != 0
}

func (this *QAbstractItemModel) MimeTypesDefault() []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAbstractItemModel_mimeTypes_default, this.h, uintptr(unsafe.Pointer(&__n)))
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

func (this *QAbstractItemModel) MimeDataDefault(indexes []*QModelIndex) *QMimeData {
	indexes_arr := make([]uintptr, len(indexes))
	for __i, __e := range indexes {
		indexes_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var indexes_p uintptr
	if len(indexes_arr) > 0 {
		indexes_p = uintptr(unsafe.Pointer(&indexes_arr[0]))
	}
	__r, _, _ := purego.SyscallN(addr_QAbstractItemModel_mimeData_default, this.h, indexes_p, uintptr(len(indexes)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(indexes)
	runtime.KeepAlive(indexes_arr)
	return bindQMimeData(__r)
}

func (this *QAbstractItemModel) CanDropMimeDataDefault(data *QMimeData, action DropAction, row int, column int, parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemModel_canDropMimeData_default, this.h, uintptr(data.UnsafePointer()), uintptr(action), uintptr(row), uintptr(column), uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(data)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QAbstractItemModel) DropMimeDataDefault(data *QMimeData, action DropAction, row int, column int, parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemModel_dropMimeData_default, this.h, uintptr(data.UnsafePointer()), uintptr(action), uintptr(row), uintptr(column), uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(data)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QAbstractItemModel) SupportedDropActionsDefault() DropAction {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemModel_supportedDropActions_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return DropAction(__r)
}

func (this *QAbstractItemModel) SupportedDragActionsDefault() DropAction {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemModel_supportedDragActions_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return DropAction(__r)
}

func (this *QAbstractItemModel) InsertRowsDefault(row int, count int, parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemModel_insertRows_default, this.h, uintptr(row), uintptr(count), uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QAbstractItemModel) InsertColumnsDefault(column int, count int, parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemModel_insertColumns_default, this.h, uintptr(column), uintptr(count), uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QAbstractItemModel) RemoveRowsDefault(row int, count int, parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemModel_removeRows_default, this.h, uintptr(row), uintptr(count), uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QAbstractItemModel) RemoveColumnsDefault(column int, count int, parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemModel_removeColumns_default, this.h, uintptr(column), uintptr(count), uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QAbstractItemModel) MoveRowsDefault(sourceParent *QModelIndex, sourceRow int, count int, destinationParent *QModelIndex, destinationChild int) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemModel_moveRows_default, this.h, uintptr(sourceParent.UnsafePointer()), uintptr(sourceRow), uintptr(count), uintptr(destinationParent.UnsafePointer()), uintptr(destinationChild))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sourceParent)
	runtime.KeepAlive(destinationParent)
	return __r != 0
}

func (this *QAbstractItemModel) MoveColumnsDefault(sourceParent *QModelIndex, sourceColumn int, count int, destinationParent *QModelIndex, destinationChild int) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemModel_moveColumns_default, this.h, uintptr(sourceParent.UnsafePointer()), uintptr(sourceColumn), uintptr(count), uintptr(destinationParent.UnsafePointer()), uintptr(destinationChild))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sourceParent)
	runtime.KeepAlive(destinationParent)
	return __r != 0
}

func (this *QAbstractItemModel) FetchMoreDefault(parent *QModelIndex) {
	purego.SyscallN(addr_QAbstractItemModel_fetchMore_default, this.h, uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
}

func (this *QAbstractItemModel) CanFetchMoreDefault(parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemModel_canFetchMore_default, this.h, uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QAbstractItemModel) FlagsDefault(index *QModelIndex) ItemFlag {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemModel_flags_default, this.h, uintptr(index.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return ItemFlag(__r)
}

func (this *QAbstractItemModel) SortDefault(column int, order SortOrder) {
	purego.SyscallN(addr_QAbstractItemModel_sort_default, this.h, uintptr(column), uintptr(order))
	runtime.KeepAlive(this)
}

func (this *QAbstractItemModel) BuddyDefault(index *QModelIndex) *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QAbstractItemModel_buddy_default, this.h, uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return newownedQModelIndex(__r)
}

func (this *QAbstractItemModel) MatchDefault(start *QModelIndex, role int, value *QVariant, hits int, flags MatchFlag) []*QModelIndex {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAbstractItemModel_match_default, this.h, uintptr(start.UnsafePointer()), uintptr(role), uintptr(value.UnsafePointer()), uintptr(hits), uintptr(flags), uintptr(unsafe.Pointer(&__n)))
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

func (this *QAbstractItemModel) SpanDefault(index *QModelIndex) *QSize {
	__r, _, _ := purego.SyscallN(addr_QAbstractItemModel_span_default, this.h, uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return newownedQSize(__r)
}

func (this *QAbstractItemModel) MultiDataDefault(index *QModelIndex, roleDataSpan *QModelRoleDataSpan) {
	purego.SyscallN(addr_QAbstractItemModel_multiData_default, this.h, uintptr(index.UnsafePointer()), uintptr(roleDataSpan.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	runtime.KeepAlive(roleDataSpan)
}

func (this *QAbstractItemModel) SubmitDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemModel_submit_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QAbstractItemModel) RevertDefault() {
	purego.SyscallN(addr_QAbstractItemModel_revert_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractItemModel) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QAbstractItemModel_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QAbstractItemModel) ResetInternalDataDefault() {
	purego.SyscallN(addr_QAbstractItemModel_resetInternalData_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractItemModel) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemModel_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QAbstractItemModel) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractItemModel_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QAbstractItemModel) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QAbstractItemModel_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractItemModel) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QAbstractItemModel_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractItemModel) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QAbstractItemModel_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractItemModel) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QAbstractItemModel_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QAbstractItemModel) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QAbstractItemModel_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QAbstractItemModel) OnHeaderDataChanged(fn func(Orientation, int, int)) int64 {
	id := registerCallback(func(orientation int32, first int32, last int32) { fn(Orientation(orientation), int(first), int(last)) })
	pg_QAbstractItemModel_connect_headerDataChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractItemModel) OnLayoutChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QAbstractItemModel_connect_layoutChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractItemModel) OnLayoutChanged2(fn func([]*QPersistentModelIndex, QAbstractItemModel__LayoutChangeHint)) int64 {
	id := registerCallback(func(parents uintptr, parents_len int64, hint int32) { fn(func() []*QPersistentModelIndex { if parents == 0 || parents_len == 0 { return nil }; __sl := unsafe.Slice((*unsafe.Pointer)(unsafe.Pointer(parents)), int(parents_len)); __o := make([]*QPersistentModelIndex, int(parents_len)); for __i := range __o { __o[__i] = UnsafeNewQPersistentModelIndex(__sl[__i]) }; return __o }(), QAbstractItemModel__LayoutChangeHint(hint)) })
	pg_QAbstractItemModel_connect_layoutChanged2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractItemModel) OnLayoutAboutToBeChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QAbstractItemModel_connect_layoutAboutToBeChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractItemModel) OnLayoutAboutToBeChanged2(fn func([]*QPersistentModelIndex, QAbstractItemModel__LayoutChangeHint)) int64 {
	id := registerCallback(func(parents uintptr, parents_len int64, hint int32) { fn(func() []*QPersistentModelIndex { if parents == 0 || parents_len == 0 { return nil }; __sl := unsafe.Slice((*unsafe.Pointer)(unsafe.Pointer(parents)), int(parents_len)); __o := make([]*QPersistentModelIndex, int(parents_len)); for __i := range __o { __o[__i] = UnsafeNewQPersistentModelIndex(__sl[__i]) }; return __o }(), QAbstractItemModel__LayoutChangeHint(hint)) })
	pg_QAbstractItemModel_connect_layoutAboutToBeChanged2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractItemModel) OnRowsAboutToBeInserted(fn func(*QModelIndex, int, int)) int64 {
	id := registerCallback(func(parent uintptr, first int32, last int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(first), int(last)) })
	pg_QAbstractItemModel_connect_rowsAboutToBeInserted(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractItemModel) OnRowsInserted(fn func(*QModelIndex, int, int)) int64 {
	id := registerCallback(func(parent uintptr, first int32, last int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(first), int(last)) })
	pg_QAbstractItemModel_connect_rowsInserted(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractItemModel) OnRowsAboutToBeRemoved(fn func(*QModelIndex, int, int)) int64 {
	id := registerCallback(func(parent uintptr, first int32, last int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(first), int(last)) })
	pg_QAbstractItemModel_connect_rowsAboutToBeRemoved(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractItemModel) OnRowsRemoved(fn func(*QModelIndex, int, int)) int64 {
	id := registerCallback(func(parent uintptr, first int32, last int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(first), int(last)) })
	pg_QAbstractItemModel_connect_rowsRemoved(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractItemModel) OnColumnsAboutToBeInserted(fn func(*QModelIndex, int, int)) int64 {
	id := registerCallback(func(parent uintptr, first int32, last int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(first), int(last)) })
	pg_QAbstractItemModel_connect_columnsAboutToBeInserted(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractItemModel) OnColumnsInserted(fn func(*QModelIndex, int, int)) int64 {
	id := registerCallback(func(parent uintptr, first int32, last int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(first), int(last)) })
	pg_QAbstractItemModel_connect_columnsInserted(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractItemModel) OnColumnsAboutToBeRemoved(fn func(*QModelIndex, int, int)) int64 {
	id := registerCallback(func(parent uintptr, first int32, last int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(first), int(last)) })
	pg_QAbstractItemModel_connect_columnsAboutToBeRemoved(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractItemModel) OnColumnsRemoved(fn func(*QModelIndex, int, int)) int64 {
	id := registerCallback(func(parent uintptr, first int32, last int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(first), int(last)) })
	pg_QAbstractItemModel_connect_columnsRemoved(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractItemModel) OnModelAboutToBeReset(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QAbstractItemModel_connect_modelAboutToBeReset(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractItemModel) OnModelReset(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QAbstractItemModel_connect_modelReset(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractItemModel) OnRowsAboutToBeMoved(fn func(*QModelIndex, int, int, *QModelIndex, int)) int64 {
	id := registerCallback(func(sourceParent uintptr, sourceStart int32, sourceEnd int32, destinationParent uintptr, destinationRow int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(sourceParent)), int(sourceStart), int(sourceEnd), UnsafeNewQModelIndex(unsafe.Pointer(destinationParent)), int(destinationRow)) })
	pg_QAbstractItemModel_connect_rowsAboutToBeMoved(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractItemModel) OnRowsMoved(fn func(*QModelIndex, int, int, *QModelIndex, int)) int64 {
	id := registerCallback(func(sourceParent uintptr, sourceStart int32, sourceEnd int32, destinationParent uintptr, destinationRow int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(sourceParent)), int(sourceStart), int(sourceEnd), UnsafeNewQModelIndex(unsafe.Pointer(destinationParent)), int(destinationRow)) })
	pg_QAbstractItemModel_connect_rowsMoved(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractItemModel) OnColumnsAboutToBeMoved(fn func(*QModelIndex, int, int, *QModelIndex, int)) int64 {
	id := registerCallback(func(sourceParent uintptr, sourceStart int32, sourceEnd int32, destinationParent uintptr, destinationColumn int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(sourceParent)), int(sourceStart), int(sourceEnd), UnsafeNewQModelIndex(unsafe.Pointer(destinationParent)), int(destinationColumn)) })
	pg_QAbstractItemModel_connect_columnsAboutToBeMoved(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractItemModel) OnColumnsMoved(fn func(*QModelIndex, int, int, *QModelIndex, int)) int64 {
	id := registerCallback(func(sourceParent uintptr, sourceStart int32, sourceEnd int32, destinationParent uintptr, destinationColumn int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(sourceParent)), int(sourceStart), int(sourceEnd), UnsafeNewQModelIndex(unsafe.Pointer(destinationParent)), int(destinationColumn)) })
	pg_QAbstractItemModel_connect_columnsMoved(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractItemModel) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QAbstractItemModel_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractItemModel) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QAbstractItemModel_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractItemModel) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QAbstractItemModel_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractItemModel) Disconnect(handle int64) {
	pg_QAbstractItemModel__disconnect(handle)
	unregisterCallback(handle)
}

func QAbstractItemModel_virtIdx(slot int64) int32 {
	switch slot {
	case -8948777187306027381:
		return 0
	case 7371339085820260189:
		return 1
	case 1960960171956884010:
		return 2
	case -8256963024669584834:
		return 3
	case -8837375942098466299:
		return 4
	case 1607084412591819503:
		return 5
	case 397329346813782521:
		return 6
	case 3369452259144514380:
		return 7
	case 35191454575202219:
		return 8
	case -3297140021896693688:
		return 9
	case 4852200306590054850:
		return 10
	case 8914190500520542409:
		return 11
	case 5464467312394730418:
		return 12
	case 5175297091024974749:
		return 13
	case 731539533644753052:
		return 14
	case -9049994684305682212:
		return 15
	case -7784642176186821205:
		return 16
	case 3421601386088570762:
		return 17
	case -2775729825962395603:
		return 18
	case -8973293846230792237:
		return 19
	case 5104072549182222656:
		return 20
	case 2078548737056540420:
		return 21
	case -6074819356947405415:
		return 22
	case -6321007591388309905:
		return 23
	case 3553294593611130806:
		return 24
	case 8639262066473964150:
		return 25
	case 1703382833545304812:
		return 26
	case 6912441694453382161:
		return 27
	case 6112860112248786499:
		return 28
	case -4341501217981328394:
		return 29
	case -8395342916587040543:
		return 30
	case 731311060944096642:
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

func QAbstractItemModel_enable(h uintptr, idx int32) { pg_QAbstractItemModel__set_override(h, idx) }

func NewQAbstractItemModel() *QAbstractItemModel {
	id := nextVirtID()
	o := newownedQAbstractItemModel(pg_QAbstractItemModel__Go_new(id, 0))
	o.UnsafeSetGoID(id)
	registerInstance(id, o.h, QAbstractItemModel_virtIdx, QAbstractItemModel_enable)
	return o
}

func NewQAbstractItemModel2(parent *QObject) *QAbstractItemModel {
	id := nextVirtID()
	o := newownedQAbstractItemModel(pg_QAbstractItemModel__Go_new(id, uintptr(parent.UnsafePointer())))
	o.UnsafeSetGoID(id)
	registerInstance(id, o.h, QAbstractItemModel_virtIdx, QAbstractItemModel_enable)
	return o
}

func (this *QAbstractItemModel) OnIndex(fn func(int, int, *QModelIndex) *QModelIndex) {
	registerVirt(this.UnsafeGoID(), -8948777187306027381, func(row int32, column int32, parent uintptr) uintptr { return uintptr(fn(int(row), int(column), UnsafeNewQModelIndex(unsafe.Pointer(parent))).UnsafePointer()) }) // slot "index"
}

func (this *QAbstractItemModel) OnParent(fn func(*QModelIndex) *QModelIndex) {
	registerVirt(this.UnsafeGoID(), 7371339085820260189, func(child uintptr) uintptr { return uintptr(fn(UnsafeNewQModelIndex(unsafe.Pointer(child))).UnsafePointer()) }) // slot "parent"
}

func (this *QAbstractItemModel) OnRowCount(fn func(*QModelIndex) int) {
	registerVirt(this.UnsafeGoID(), 1960960171956884010, func(parent uintptr) int32 { return int32(fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)))) }) // slot "rowCount"
}

func (this *QAbstractItemModel) OnColumnCount(fn func(*QModelIndex) int) {
	registerVirt(this.UnsafeGoID(), -8256963024669584834, func(parent uintptr) int32 { return int32(fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)))) }) // slot "columnCount"
}

func (this *QAbstractItemModel) OnData(fn func(*QModelIndex, int) *QVariant) {
	registerVirt(this.UnsafeGoID(), -8837375942098466299, func(index uintptr, role int32) uintptr { return uintptr(fn(UnsafeNewQModelIndex(unsafe.Pointer(index)), int(role)).UnsafePointer()) }) // slot "data"
}

func (this *QAbstractItemModel) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QAbstractItemModel) OnSibling(fn func(int, int, *QModelIndex) *QModelIndex) {
	registerVirt(this.UnsafeGoID(), 397329346813782521, func(row int32, column int32, idx uintptr) uintptr { return uintptr(fn(int(row), int(column), UnsafeNewQModelIndex(unsafe.Pointer(idx))).UnsafePointer()) }) // slot "sibling"
}

func (this *QAbstractItemModel) OnHasChildren(fn func(*QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), 3369452259144514380, func(parent uintptr) int32 { return b2i(fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)))) }) // slot "hasChildren"
}

func (this *QAbstractItemModel) OnSetData(fn func(*QModelIndex, *QVariant, int) bool) {
	registerVirt(this.UnsafeGoID(), 35191454575202219, func(index uintptr, value uintptr, role int32) int32 { return b2i(fn(UnsafeNewQModelIndex(unsafe.Pointer(index)), UnsafeNewQVariant(unsafe.Pointer(value)), int(role))) }) // slot "setData"
}

func (this *QAbstractItemModel) OnHeaderData(fn func(int, Orientation, int) *QVariant) {
	registerVirt(this.UnsafeGoID(), -3297140021896693688, func(section int32, orientation int32, role int32) uintptr { return uintptr(fn(int(section), Orientation(orientation), int(role)).UnsafePointer()) }) // slot "headerData"
}

func (this *QAbstractItemModel) OnSetHeaderData(fn func(int, Orientation, *QVariant, int) bool) {
	registerVirt(this.UnsafeGoID(), 4852200306590054850, func(section int32, orientation int32, value uintptr, role int32) int32 { return b2i(fn(int(section), Orientation(orientation), UnsafeNewQVariant(unsafe.Pointer(value)), int(role))) }) // slot "setHeaderData"
}

func (this *QAbstractItemModel) OnClearItemData(fn func(*QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), 8914190500520542409, func(index uintptr) int32 { return b2i(fn(UnsafeNewQModelIndex(unsafe.Pointer(index)))) }) // slot "clearItemData"
}

func (this *QAbstractItemModel) OnMimeTypes(fn func() []string) {
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

func (this *QAbstractItemModel) OnMimeData(fn func([]*QModelIndex) *QMimeData) {
	registerVirt(this.UnsafeGoID(), 5175297091024974749, func(indexes uintptr, indexes_len int64) uintptr { return uintptr(fn(func() []*QModelIndex { if indexes == 0 || indexes_len == 0 { return nil }; __sl := unsafe.Slice((*unsafe.Pointer)(unsafe.Pointer(indexes)), int(indexes_len)); __o := make([]*QModelIndex, int(indexes_len)); for __i := range __o { __o[__i] = UnsafeNewQModelIndex(__sl[__i]) }; return __o }()).UnsafePointer()) }) // slot "mimeData"
}

func (this *QAbstractItemModel) OnCanDropMimeData(fn func(*QMimeData, DropAction, int, int, *QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), 731539533644753052, func(data uintptr, action int32, row int32, column int32, parent uintptr) int32 { return b2i(fn(UnsafeNewQMimeData(unsafe.Pointer(data)), DropAction(action), int(row), int(column), UnsafeNewQModelIndex(unsafe.Pointer(parent)))) }) // slot "canDropMimeData"
}

func (this *QAbstractItemModel) OnDropMimeData(fn func(*QMimeData, DropAction, int, int, *QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), -9049994684305682212, func(data uintptr, action int32, row int32, column int32, parent uintptr) int32 { return b2i(fn(UnsafeNewQMimeData(unsafe.Pointer(data)), DropAction(action), int(row), int(column), UnsafeNewQModelIndex(unsafe.Pointer(parent)))) }) // slot "dropMimeData"
}

func (this *QAbstractItemModel) OnSupportedDropActions(fn func() DropAction) {
	registerVirt(this.UnsafeGoID(), -7784642176186821205, func() int32 { return int32(fn()) }) // slot "supportedDropActions"
}

func (this *QAbstractItemModel) OnSupportedDragActions(fn func() DropAction) {
	registerVirt(this.UnsafeGoID(), 3421601386088570762, func() int32 { return int32(fn()) }) // slot "supportedDragActions"
}

func (this *QAbstractItemModel) OnInsertRows(fn func(int, int, *QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), -2775729825962395603, func(row int32, count int32, parent uintptr) int32 { return b2i(fn(int(row), int(count), UnsafeNewQModelIndex(unsafe.Pointer(parent)))) }) // slot "insertRows"
}

func (this *QAbstractItemModel) OnInsertColumns(fn func(int, int, *QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), -8973293846230792237, func(column int32, count int32, parent uintptr) int32 { return b2i(fn(int(column), int(count), UnsafeNewQModelIndex(unsafe.Pointer(parent)))) }) // slot "insertColumns"
}

func (this *QAbstractItemModel) OnRemoveRows(fn func(int, int, *QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), 5104072549182222656, func(row int32, count int32, parent uintptr) int32 { return b2i(fn(int(row), int(count), UnsafeNewQModelIndex(unsafe.Pointer(parent)))) }) // slot "removeRows"
}

func (this *QAbstractItemModel) OnRemoveColumns(fn func(int, int, *QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), 2078548737056540420, func(column int32, count int32, parent uintptr) int32 { return b2i(fn(int(column), int(count), UnsafeNewQModelIndex(unsafe.Pointer(parent)))) }) // slot "removeColumns"
}

func (this *QAbstractItemModel) OnMoveRows(fn func(*QModelIndex, int, int, *QModelIndex, int) bool) {
	registerVirt(this.UnsafeGoID(), -6074819356947405415, func(sourceParent uintptr, sourceRow int32, count int32, destinationParent uintptr, destinationChild int32) int32 { return b2i(fn(UnsafeNewQModelIndex(unsafe.Pointer(sourceParent)), int(sourceRow), int(count), UnsafeNewQModelIndex(unsafe.Pointer(destinationParent)), int(destinationChild))) }) // slot "moveRows"
}

func (this *QAbstractItemModel) OnMoveColumns(fn func(*QModelIndex, int, int, *QModelIndex, int) bool) {
	registerVirt(this.UnsafeGoID(), -6321007591388309905, func(sourceParent uintptr, sourceColumn int32, count int32, destinationParent uintptr, destinationChild int32) int32 { return b2i(fn(UnsafeNewQModelIndex(unsafe.Pointer(sourceParent)), int(sourceColumn), int(count), UnsafeNewQModelIndex(unsafe.Pointer(destinationParent)), int(destinationChild))) }) // slot "moveColumns"
}

func (this *QAbstractItemModel) OnFetchMore(fn func(*QModelIndex)) {
	registerVirt(this.UnsafeGoID(), 3553294593611130806, func(parent uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent))) }) // slot "fetchMore"
}

func (this *QAbstractItemModel) OnCanFetchMore(fn func(*QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), 8639262066473964150, func(parent uintptr) int32 { return b2i(fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)))) }) // slot "canFetchMore"
}

func (this *QAbstractItemModel) OnFlags(fn func(*QModelIndex) ItemFlag) {
	registerVirt(this.UnsafeGoID(), 1703382833545304812, func(index uintptr) int32 { return int32(fn(UnsafeNewQModelIndex(unsafe.Pointer(index)))) }) // slot "flags"
}

func (this *QAbstractItemModel) OnSort(fn func(int, SortOrder)) {
	registerVirt(this.UnsafeGoID(), 6912441694453382161, func(column int32, order int32) { fn(int(column), SortOrder(order)) }) // slot "sort"
}

func (this *QAbstractItemModel) OnBuddy(fn func(*QModelIndex) *QModelIndex) {
	registerVirt(this.UnsafeGoID(), 6112860112248786499, func(index uintptr) uintptr { return uintptr(fn(UnsafeNewQModelIndex(unsafe.Pointer(index))).UnsafePointer()) }) // slot "buddy"
}

func (this *QAbstractItemModel) OnMatch(fn func(*QModelIndex, int, *QVariant, int, MatchFlag) []*QModelIndex) {
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

func (this *QAbstractItemModel) OnSpan(fn func(*QModelIndex) *QSize) {
	registerVirt(this.UnsafeGoID(), -8395342916587040543, func(index uintptr) uintptr { return uintptr(fn(UnsafeNewQModelIndex(unsafe.Pointer(index))).UnsafePointer()) }) // slot "span"
}

func (this *QAbstractItemModel) OnMultiData(fn func(*QModelIndex, *QModelRoleDataSpan)) {
	registerVirt(this.UnsafeGoID(), 731311060944096642, func(index uintptr, roleDataSpan uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(index)), UnsafeNewQModelRoleDataSpan(unsafe.Pointer(roleDataSpan))) }) // slot "multiData"
}

func (this *QAbstractItemModel) OnSubmit(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -3788331537442685179, func() int32 { return b2i(fn()) }) // slot "submit"
}

func (this *QAbstractItemModel) OnRevert(fn func()) {
	registerVirt(this.UnsafeGoID(), -1522951607536117149, func() { fn() }) // slot "revert"
}

func (this *QAbstractItemModel) OnResetInternalData(fn func()) {
	registerVirt(this.UnsafeGoID(), -4510675770784051183, func() { fn() }) // slot "resetInternalData"
}

func (this *QAbstractItemModel) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QAbstractItemModel) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QAbstractItemModel) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QAbstractItemModel) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QAbstractItemModel) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QAbstractItemModel) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QAbstractItemModel) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QAbstractItemModel) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QAbstractItemModel_delete(this.h)
	}
	this.UnsafeRelease()
}
