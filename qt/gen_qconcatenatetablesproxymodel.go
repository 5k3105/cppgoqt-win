package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QConcatenateTablesProxyModel_chain uintptr
	addr_QConcatenateTablesProxyModel_new uintptr
	addr_QConcatenateTablesProxyModel_new2 uintptr
	addr_QConcatenateTablesProxyModel_tr uintptr
	addr_QConcatenateTablesProxyModel_sourceModels uintptr
	addr_QConcatenateTablesProxyModel_addSourceModel uintptr
	addr_QConcatenateTablesProxyModel_removeSourceModel uintptr
	addr_QConcatenateTablesProxyModel_mapFromSource uintptr
	addr_QConcatenateTablesProxyModel_mapToSource uintptr
	addr_QConcatenateTablesProxyModel_createIndex uintptr
	addr_QConcatenateTablesProxyModel_createIndex2 uintptr
	addr_QConcatenateTablesProxyModel_encodeData uintptr
	addr_QConcatenateTablesProxyModel_decodeData uintptr
	addr_QConcatenateTablesProxyModel_beginInsertRows uintptr
	addr_QConcatenateTablesProxyModel_endInsertRows uintptr
	addr_QConcatenateTablesProxyModel_beginRemoveRows uintptr
	addr_QConcatenateTablesProxyModel_endRemoveRows uintptr
	addr_QConcatenateTablesProxyModel_beginMoveRows uintptr
	addr_QConcatenateTablesProxyModel_endMoveRows uintptr
	addr_QConcatenateTablesProxyModel_beginInsertColumns uintptr
	addr_QConcatenateTablesProxyModel_endInsertColumns uintptr
	addr_QConcatenateTablesProxyModel_beginRemoveColumns uintptr
	addr_QConcatenateTablesProxyModel_endRemoveColumns uintptr
	addr_QConcatenateTablesProxyModel_beginMoveColumns uintptr
	addr_QConcatenateTablesProxyModel_endMoveColumns uintptr
	addr_QConcatenateTablesProxyModel_beginResetModel uintptr
	addr_QConcatenateTablesProxyModel_endResetModel uintptr
	addr_QConcatenateTablesProxyModel_changePersistentIndex uintptr
	addr_QConcatenateTablesProxyModel_changePersistentIndexList uintptr
	addr_QConcatenateTablesProxyModel_persistentIndexList uintptr
	addr_QConcatenateTablesProxyModel_connect uintptr
	addr_QConcatenateTablesProxyModel_connect2 uintptr
	addr_QConcatenateTablesProxyModel_sender uintptr
	addr_QConcatenateTablesProxyModel_senderSignalIndex uintptr
	addr_QConcatenateTablesProxyModel_receivers uintptr
	addr_QConcatenateTablesProxyModel_isSignalConnected uintptr
	addr_QConcatenateTablesProxyModel_tr2 uintptr
	addr_QConcatenateTablesProxyModel_tr3 uintptr
	addr_QConcatenateTablesProxyModel_connect3 uintptr
	addr_QConcatenateTablesProxyModel_connect4 uintptr
	addr_QConcatenateTablesProxyModel_metaObject_default uintptr
	addr_QConcatenateTablesProxyModel_data_default uintptr
	addr_QConcatenateTablesProxyModel_setData_default uintptr
	addr_QConcatenateTablesProxyModel_flags_default uintptr
	addr_QConcatenateTablesProxyModel_index_default uintptr
	addr_QConcatenateTablesProxyModel_parent_default uintptr
	addr_QConcatenateTablesProxyModel_rowCount_default uintptr
	addr_QConcatenateTablesProxyModel_headerData_default uintptr
	addr_QConcatenateTablesProxyModel_columnCount_default uintptr
	addr_QConcatenateTablesProxyModel_mimeTypes_default uintptr
	addr_QConcatenateTablesProxyModel_mimeData_default uintptr
	addr_QConcatenateTablesProxyModel_canDropMimeData_default uintptr
	addr_QConcatenateTablesProxyModel_dropMimeData_default uintptr
	addr_QConcatenateTablesProxyModel_span_default uintptr
	addr_QConcatenateTablesProxyModel_sibling_default uintptr
	addr_QConcatenateTablesProxyModel_hasChildren_default uintptr
	addr_QConcatenateTablesProxyModel_setHeaderData_default uintptr
	addr_QConcatenateTablesProxyModel_clearItemData_default uintptr
	addr_QConcatenateTablesProxyModel_supportedDropActions_default uintptr
	addr_QConcatenateTablesProxyModel_supportedDragActions_default uintptr
	addr_QConcatenateTablesProxyModel_insertRows_default uintptr
	addr_QConcatenateTablesProxyModel_insertColumns_default uintptr
	addr_QConcatenateTablesProxyModel_removeRows_default uintptr
	addr_QConcatenateTablesProxyModel_removeColumns_default uintptr
	addr_QConcatenateTablesProxyModel_moveRows_default uintptr
	addr_QConcatenateTablesProxyModel_moveColumns_default uintptr
	addr_QConcatenateTablesProxyModel_fetchMore_default uintptr
	addr_QConcatenateTablesProxyModel_canFetchMore_default uintptr
	addr_QConcatenateTablesProxyModel_sort_default uintptr
	addr_QConcatenateTablesProxyModel_buddy_default uintptr
	addr_QConcatenateTablesProxyModel_match_default uintptr
	addr_QConcatenateTablesProxyModel_multiData_default uintptr
	addr_QConcatenateTablesProxyModel_submit_default uintptr
	addr_QConcatenateTablesProxyModel_revert_default uintptr
	addr_QConcatenateTablesProxyModel_resetInternalData_default uintptr
	addr_QConcatenateTablesProxyModel_event_default uintptr
	addr_QConcatenateTablesProxyModel_eventFilter_default uintptr
	addr_QConcatenateTablesProxyModel_timerEvent_default uintptr
	addr_QConcatenateTablesProxyModel_childEvent_default uintptr
	addr_QConcatenateTablesProxyModel_customEvent_default uintptr
	addr_QConcatenateTablesProxyModel_connectNotify_default uintptr
	addr_QConcatenateTablesProxyModel_disconnectNotify_default uintptr
	pg_QConcatenateTablesProxyModel_connect_headerDataChanged func(uintptr, int64)
	pg_QConcatenateTablesProxyModel_connect_layoutChanged func(uintptr, int64)
	pg_QConcatenateTablesProxyModel_connect_layoutChanged2 func(uintptr, int64)
	pg_QConcatenateTablesProxyModel_connect_layoutAboutToBeChanged func(uintptr, int64)
	pg_QConcatenateTablesProxyModel_connect_layoutAboutToBeChanged2 func(uintptr, int64)
	pg_QConcatenateTablesProxyModel_connect_rowsAboutToBeInserted func(uintptr, int64)
	pg_QConcatenateTablesProxyModel_connect_rowsInserted func(uintptr, int64)
	pg_QConcatenateTablesProxyModel_connect_rowsAboutToBeRemoved func(uintptr, int64)
	pg_QConcatenateTablesProxyModel_connect_rowsRemoved func(uintptr, int64)
	pg_QConcatenateTablesProxyModel_connect_columnsAboutToBeInserted func(uintptr, int64)
	pg_QConcatenateTablesProxyModel_connect_columnsInserted func(uintptr, int64)
	pg_QConcatenateTablesProxyModel_connect_columnsAboutToBeRemoved func(uintptr, int64)
	pg_QConcatenateTablesProxyModel_connect_columnsRemoved func(uintptr, int64)
	pg_QConcatenateTablesProxyModel_connect_modelAboutToBeReset func(uintptr, int64)
	pg_QConcatenateTablesProxyModel_connect_modelReset func(uintptr, int64)
	pg_QConcatenateTablesProxyModel_connect_rowsAboutToBeMoved func(uintptr, int64)
	pg_QConcatenateTablesProxyModel_connect_rowsMoved func(uintptr, int64)
	pg_QConcatenateTablesProxyModel_connect_columnsAboutToBeMoved func(uintptr, int64)
	pg_QConcatenateTablesProxyModel_connect_columnsMoved func(uintptr, int64)
	pg_QConcatenateTablesProxyModel_connect_destroyed func(uintptr, int64)
	pg_QConcatenateTablesProxyModel_connect_destroyed2 func(uintptr, int64)
	pg_QConcatenateTablesProxyModel_connect_objectNameChanged func(uintptr, int64)
	pg_QConcatenateTablesProxyModel__disconnect func(int64)
	pg_QConcatenateTablesProxyModel__set_override func(uintptr, int32)
	pg_QConcatenateTablesProxyModel_delete func(uintptr)
)

func init() {
	symaddr(&addr_QConcatenateTablesProxyModel_chain, "QConcatenateTablesProxyModel_chain")
	symaddr(&addr_QConcatenateTablesProxyModel_new, "QConcatenateTablesProxyModel_new")
	symaddr(&addr_QConcatenateTablesProxyModel_new2, "QConcatenateTablesProxyModel_new2")
	symaddr(&addr_QConcatenateTablesProxyModel_tr, "QConcatenateTablesProxyModel_tr")
	symaddr(&addr_QConcatenateTablesProxyModel_sourceModels, "QConcatenateTablesProxyModel_sourceModels")
	symaddr(&addr_QConcatenateTablesProxyModel_addSourceModel, "QConcatenateTablesProxyModel_addSourceModel")
	symaddr(&addr_QConcatenateTablesProxyModel_removeSourceModel, "QConcatenateTablesProxyModel_removeSourceModel")
	symaddr(&addr_QConcatenateTablesProxyModel_mapFromSource, "QConcatenateTablesProxyModel_mapFromSource")
	symaddr(&addr_QConcatenateTablesProxyModel_mapToSource, "QConcatenateTablesProxyModel_mapToSource")
	symaddr(&addr_QConcatenateTablesProxyModel_createIndex, "QConcatenateTablesProxyModel_createIndex")
	symaddr(&addr_QConcatenateTablesProxyModel_createIndex2, "QConcatenateTablesProxyModel_createIndex2")
	symaddr(&addr_QConcatenateTablesProxyModel_encodeData, "QConcatenateTablesProxyModel_encodeData")
	symaddr(&addr_QConcatenateTablesProxyModel_decodeData, "QConcatenateTablesProxyModel_decodeData")
	symaddr(&addr_QConcatenateTablesProxyModel_beginInsertRows, "QConcatenateTablesProxyModel_beginInsertRows")
	symaddr(&addr_QConcatenateTablesProxyModel_endInsertRows, "QConcatenateTablesProxyModel_endInsertRows")
	symaddr(&addr_QConcatenateTablesProxyModel_beginRemoveRows, "QConcatenateTablesProxyModel_beginRemoveRows")
	symaddr(&addr_QConcatenateTablesProxyModel_endRemoveRows, "QConcatenateTablesProxyModel_endRemoveRows")
	symaddr(&addr_QConcatenateTablesProxyModel_beginMoveRows, "QConcatenateTablesProxyModel_beginMoveRows")
	symaddr(&addr_QConcatenateTablesProxyModel_endMoveRows, "QConcatenateTablesProxyModel_endMoveRows")
	symaddr(&addr_QConcatenateTablesProxyModel_beginInsertColumns, "QConcatenateTablesProxyModel_beginInsertColumns")
	symaddr(&addr_QConcatenateTablesProxyModel_endInsertColumns, "QConcatenateTablesProxyModel_endInsertColumns")
	symaddr(&addr_QConcatenateTablesProxyModel_beginRemoveColumns, "QConcatenateTablesProxyModel_beginRemoveColumns")
	symaddr(&addr_QConcatenateTablesProxyModel_endRemoveColumns, "QConcatenateTablesProxyModel_endRemoveColumns")
	symaddr(&addr_QConcatenateTablesProxyModel_beginMoveColumns, "QConcatenateTablesProxyModel_beginMoveColumns")
	symaddr(&addr_QConcatenateTablesProxyModel_endMoveColumns, "QConcatenateTablesProxyModel_endMoveColumns")
	symaddr(&addr_QConcatenateTablesProxyModel_beginResetModel, "QConcatenateTablesProxyModel_beginResetModel")
	symaddr(&addr_QConcatenateTablesProxyModel_endResetModel, "QConcatenateTablesProxyModel_endResetModel")
	symaddr(&addr_QConcatenateTablesProxyModel_changePersistentIndex, "QConcatenateTablesProxyModel_changePersistentIndex")
	symaddr(&addr_QConcatenateTablesProxyModel_changePersistentIndexList, "QConcatenateTablesProxyModel_changePersistentIndexList")
	symaddr(&addr_QConcatenateTablesProxyModel_persistentIndexList, "QConcatenateTablesProxyModel_persistentIndexList")
	symaddr(&addr_QConcatenateTablesProxyModel_connect, "QConcatenateTablesProxyModel_connect")
	symaddr(&addr_QConcatenateTablesProxyModel_connect2, "QConcatenateTablesProxyModel_connect2")
	symaddr(&addr_QConcatenateTablesProxyModel_sender, "QConcatenateTablesProxyModel_sender")
	symaddr(&addr_QConcatenateTablesProxyModel_senderSignalIndex, "QConcatenateTablesProxyModel_senderSignalIndex")
	symaddr(&addr_QConcatenateTablesProxyModel_receivers, "QConcatenateTablesProxyModel_receivers")
	symaddr(&addr_QConcatenateTablesProxyModel_isSignalConnected, "QConcatenateTablesProxyModel_isSignalConnected")
	symaddr(&addr_QConcatenateTablesProxyModel_tr2, "QConcatenateTablesProxyModel_tr2")
	symaddr(&addr_QConcatenateTablesProxyModel_tr3, "QConcatenateTablesProxyModel_tr3")
	symaddr(&addr_QConcatenateTablesProxyModel_connect3, "QConcatenateTablesProxyModel_connect3")
	symaddr(&addr_QConcatenateTablesProxyModel_connect4, "QConcatenateTablesProxyModel_connect4")
	symaddr(&addr_QConcatenateTablesProxyModel_metaObject_default, "QConcatenateTablesProxyModel_metaObject_default")
	symaddr(&addr_QConcatenateTablesProxyModel_data_default, "QConcatenateTablesProxyModel_data_default")
	symaddr(&addr_QConcatenateTablesProxyModel_setData_default, "QConcatenateTablesProxyModel_setData_default")
	symaddr(&addr_QConcatenateTablesProxyModel_flags_default, "QConcatenateTablesProxyModel_flags_default")
	symaddr(&addr_QConcatenateTablesProxyModel_index_default, "QConcatenateTablesProxyModel_index_default")
	symaddr(&addr_QConcatenateTablesProxyModel_parent_default, "QConcatenateTablesProxyModel_parent_default")
	symaddr(&addr_QConcatenateTablesProxyModel_rowCount_default, "QConcatenateTablesProxyModel_rowCount_default")
	symaddr(&addr_QConcatenateTablesProxyModel_headerData_default, "QConcatenateTablesProxyModel_headerData_default")
	symaddr(&addr_QConcatenateTablesProxyModel_columnCount_default, "QConcatenateTablesProxyModel_columnCount_default")
	symaddr(&addr_QConcatenateTablesProxyModel_mimeTypes_default, "QConcatenateTablesProxyModel_mimeTypes_default")
	symaddr(&addr_QConcatenateTablesProxyModel_mimeData_default, "QConcatenateTablesProxyModel_mimeData_default")
	symaddr(&addr_QConcatenateTablesProxyModel_canDropMimeData_default, "QConcatenateTablesProxyModel_canDropMimeData_default")
	symaddr(&addr_QConcatenateTablesProxyModel_dropMimeData_default, "QConcatenateTablesProxyModel_dropMimeData_default")
	symaddr(&addr_QConcatenateTablesProxyModel_span_default, "QConcatenateTablesProxyModel_span_default")
	symaddr(&addr_QConcatenateTablesProxyModel_sibling_default, "QConcatenateTablesProxyModel_sibling_default")
	symaddr(&addr_QConcatenateTablesProxyModel_hasChildren_default, "QConcatenateTablesProxyModel_hasChildren_default")
	symaddr(&addr_QConcatenateTablesProxyModel_setHeaderData_default, "QConcatenateTablesProxyModel_setHeaderData_default")
	symaddr(&addr_QConcatenateTablesProxyModel_clearItemData_default, "QConcatenateTablesProxyModel_clearItemData_default")
	symaddr(&addr_QConcatenateTablesProxyModel_supportedDropActions_default, "QConcatenateTablesProxyModel_supportedDropActions_default")
	symaddr(&addr_QConcatenateTablesProxyModel_supportedDragActions_default, "QConcatenateTablesProxyModel_supportedDragActions_default")
	symaddr(&addr_QConcatenateTablesProxyModel_insertRows_default, "QConcatenateTablesProxyModel_insertRows_default")
	symaddr(&addr_QConcatenateTablesProxyModel_insertColumns_default, "QConcatenateTablesProxyModel_insertColumns_default")
	symaddr(&addr_QConcatenateTablesProxyModel_removeRows_default, "QConcatenateTablesProxyModel_removeRows_default")
	symaddr(&addr_QConcatenateTablesProxyModel_removeColumns_default, "QConcatenateTablesProxyModel_removeColumns_default")
	symaddr(&addr_QConcatenateTablesProxyModel_moveRows_default, "QConcatenateTablesProxyModel_moveRows_default")
	symaddr(&addr_QConcatenateTablesProxyModel_moveColumns_default, "QConcatenateTablesProxyModel_moveColumns_default")
	symaddr(&addr_QConcatenateTablesProxyModel_fetchMore_default, "QConcatenateTablesProxyModel_fetchMore_default")
	symaddr(&addr_QConcatenateTablesProxyModel_canFetchMore_default, "QConcatenateTablesProxyModel_canFetchMore_default")
	symaddr(&addr_QConcatenateTablesProxyModel_sort_default, "QConcatenateTablesProxyModel_sort_default")
	symaddr(&addr_QConcatenateTablesProxyModel_buddy_default, "QConcatenateTablesProxyModel_buddy_default")
	symaddr(&addr_QConcatenateTablesProxyModel_match_default, "QConcatenateTablesProxyModel_match_default")
	symaddr(&addr_QConcatenateTablesProxyModel_multiData_default, "QConcatenateTablesProxyModel_multiData_default")
	symaddr(&addr_QConcatenateTablesProxyModel_submit_default, "QConcatenateTablesProxyModel_submit_default")
	symaddr(&addr_QConcatenateTablesProxyModel_revert_default, "QConcatenateTablesProxyModel_revert_default")
	symaddr(&addr_QConcatenateTablesProxyModel_resetInternalData_default, "QConcatenateTablesProxyModel_resetInternalData_default")
	symaddr(&addr_QConcatenateTablesProxyModel_event_default, "QConcatenateTablesProxyModel_event_default")
	symaddr(&addr_QConcatenateTablesProxyModel_eventFilter_default, "QConcatenateTablesProxyModel_eventFilter_default")
	symaddr(&addr_QConcatenateTablesProxyModel_timerEvent_default, "QConcatenateTablesProxyModel_timerEvent_default")
	symaddr(&addr_QConcatenateTablesProxyModel_childEvent_default, "QConcatenateTablesProxyModel_childEvent_default")
	symaddr(&addr_QConcatenateTablesProxyModel_customEvent_default, "QConcatenateTablesProxyModel_customEvent_default")
	symaddr(&addr_QConcatenateTablesProxyModel_connectNotify_default, "QConcatenateTablesProxyModel_connectNotify_default")
	symaddr(&addr_QConcatenateTablesProxyModel_disconnectNotify_default, "QConcatenateTablesProxyModel_disconnectNotify_default")
	reg(&pg_QConcatenateTablesProxyModel_connect_headerDataChanged, "QConcatenateTablesProxyModel_connect_headerDataChanged")
	reg(&pg_QConcatenateTablesProxyModel_connect_layoutChanged, "QConcatenateTablesProxyModel_connect_layoutChanged")
	reg(&pg_QConcatenateTablesProxyModel_connect_layoutChanged2, "QConcatenateTablesProxyModel_connect_layoutChanged2")
	reg(&pg_QConcatenateTablesProxyModel_connect_layoutAboutToBeChanged, "QConcatenateTablesProxyModel_connect_layoutAboutToBeChanged")
	reg(&pg_QConcatenateTablesProxyModel_connect_layoutAboutToBeChanged2, "QConcatenateTablesProxyModel_connect_layoutAboutToBeChanged2")
	reg(&pg_QConcatenateTablesProxyModel_connect_rowsAboutToBeInserted, "QConcatenateTablesProxyModel_connect_rowsAboutToBeInserted")
	reg(&pg_QConcatenateTablesProxyModel_connect_rowsInserted, "QConcatenateTablesProxyModel_connect_rowsInserted")
	reg(&pg_QConcatenateTablesProxyModel_connect_rowsAboutToBeRemoved, "QConcatenateTablesProxyModel_connect_rowsAboutToBeRemoved")
	reg(&pg_QConcatenateTablesProxyModel_connect_rowsRemoved, "QConcatenateTablesProxyModel_connect_rowsRemoved")
	reg(&pg_QConcatenateTablesProxyModel_connect_columnsAboutToBeInserted, "QConcatenateTablesProxyModel_connect_columnsAboutToBeInserted")
	reg(&pg_QConcatenateTablesProxyModel_connect_columnsInserted, "QConcatenateTablesProxyModel_connect_columnsInserted")
	reg(&pg_QConcatenateTablesProxyModel_connect_columnsAboutToBeRemoved, "QConcatenateTablesProxyModel_connect_columnsAboutToBeRemoved")
	reg(&pg_QConcatenateTablesProxyModel_connect_columnsRemoved, "QConcatenateTablesProxyModel_connect_columnsRemoved")
	reg(&pg_QConcatenateTablesProxyModel_connect_modelAboutToBeReset, "QConcatenateTablesProxyModel_connect_modelAboutToBeReset")
	reg(&pg_QConcatenateTablesProxyModel_connect_modelReset, "QConcatenateTablesProxyModel_connect_modelReset")
	reg(&pg_QConcatenateTablesProxyModel_connect_rowsAboutToBeMoved, "QConcatenateTablesProxyModel_connect_rowsAboutToBeMoved")
	reg(&pg_QConcatenateTablesProxyModel_connect_rowsMoved, "QConcatenateTablesProxyModel_connect_rowsMoved")
	reg(&pg_QConcatenateTablesProxyModel_connect_columnsAboutToBeMoved, "QConcatenateTablesProxyModel_connect_columnsAboutToBeMoved")
	reg(&pg_QConcatenateTablesProxyModel_connect_columnsMoved, "QConcatenateTablesProxyModel_connect_columnsMoved")
	reg(&pg_QConcatenateTablesProxyModel_connect_destroyed, "QConcatenateTablesProxyModel_connect_destroyed")
	reg(&pg_QConcatenateTablesProxyModel_connect_destroyed2, "QConcatenateTablesProxyModel_connect_destroyed2")
	reg(&pg_QConcatenateTablesProxyModel_connect_objectNameChanged, "QConcatenateTablesProxyModel_connect_objectNameChanged")
	reg(&pg_QConcatenateTablesProxyModel__disconnect, "QConcatenateTablesProxyModel__disconnect")
	reg(&pg_QConcatenateTablesProxyModel__set_override, "QConcatenateTablesProxyModel__set_override")
	reg(&pg_QConcatenateTablesProxyModel_delete, "QConcatenateTablesProxyModel_delete")
}

type QConcatenateTablesProxyModel struct {
	h uintptr
	*QAbstractItemModel
}

func bindQConcatenateTablesProxyModel(h uintptr) *QConcatenateTablesProxyModel {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QConcatenateTablesProxyModel_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QConcatenateTablesProxyModel{h: h, QAbstractItemModel: &QAbstractItemModel{h: __c[0], QObject: &QObject{h: __c[1]}}}
}

func newownedQConcatenateTablesProxyModel(h uintptr) *QConcatenateTablesProxyModel { return bindQConcatenateTablesProxyModel(h) }

func (this *QConcatenateTablesProxyModel) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QConcatenateTablesProxyModel) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQConcatenateTablesProxyModel(h unsafe.Pointer) *QConcatenateTablesProxyModel { return bindQConcatenateTablesProxyModel(uintptr(h)) }
func UnsafeNewQConcatenateTablesProxyModelOwned(h unsafe.Pointer) *QConcatenateTablesProxyModel { return newownedQConcatenateTablesProxyModel(uintptr(h)) }

func (this *QConcatenateTablesProxyModel) UnsafeRelease() {
	this.h = 0
	if this.QAbstractItemModel != nil {
		this.QAbstractItemModel.UnsafeRelease()
	}
}

func NewQConcatenateTablesProxyModel() *QConcatenateTablesProxyModel {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QConcatenateTablesProxyModel_new, uintptr(__goid))
	__o := newownedQConcatenateTablesProxyModel(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QConcatenateTablesProxyModel_virtIdx, QConcatenateTablesProxyModel_enable)
	return __o
}

func NewQConcatenateTablesProxyModel2(parent *QObject) *QConcatenateTablesProxyModel {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QConcatenateTablesProxyModel_new2, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQConcatenateTablesProxyModel(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QConcatenateTablesProxyModel_virtIdx, QConcatenateTablesProxyModel_enable)
	return __o
}

func QConcatenateTablesProxyModel_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QConcatenateTablesProxyModel_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QConcatenateTablesProxyModel) SourceModels() []*QAbstractItemModel {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QConcatenateTablesProxyModel_sourceModels, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QAbstractItemModel, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = bindQAbstractItemModel(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QConcatenateTablesProxyModel) AddSourceModel(sourceModel *QAbstractItemModel) {
	purego.SyscallN(addr_QConcatenateTablesProxyModel_addSourceModel, this.h, uintptr(sourceModel.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sourceModel)
}

func (this *QConcatenateTablesProxyModel) RemoveSourceModel(sourceModel *QAbstractItemModel) {
	purego.SyscallN(addr_QConcatenateTablesProxyModel_removeSourceModel, this.h, uintptr(sourceModel.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sourceModel)
}

func (this *QConcatenateTablesProxyModel) MapFromSource(sourceIndex *QModelIndex) *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QConcatenateTablesProxyModel_mapFromSource, this.h, uintptr(sourceIndex.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sourceIndex)
	return newownedQModelIndex(__r)
}

func (this *QConcatenateTablesProxyModel) MapToSource(proxyIndex *QModelIndex) *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QConcatenateTablesProxyModel_mapToSource, this.h, uintptr(proxyIndex.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(proxyIndex)
	return newownedQModelIndex(__r)
}

func (this *QConcatenateTablesProxyModel) CreateIndex(row int, column int) *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QConcatenateTablesProxyModel_createIndex, this.h, uintptr(row), uintptr(column))
	runtime.KeepAlive(this)
	return newownedQModelIndex(__r)
}

func (this *QConcatenateTablesProxyModel) CreateIndex2(row int, column int, id uintptr) *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QConcatenateTablesProxyModel_createIndex2, this.h, uintptr(row), uintptr(column), uintptr(id))
	runtime.KeepAlive(this)
	return newownedQModelIndex(__r)
}

func (this *QConcatenateTablesProxyModel) EncodeData(indexes []*QModelIndex, stream *QDataStream) {
	indexes_arr := make([]uintptr, len(indexes))
	for __i, __e := range indexes {
		indexes_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var indexes_p uintptr
	if len(indexes_arr) > 0 {
		indexes_p = uintptr(unsafe.Pointer(&indexes_arr[0]))
	}
	purego.SyscallN(addr_QConcatenateTablesProxyModel_encodeData, this.h, indexes_p, uintptr(len(indexes)), uintptr(stream.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(indexes)
	runtime.KeepAlive(indexes_arr)
	runtime.KeepAlive(stream)
}

func (this *QConcatenateTablesProxyModel) DecodeData(row int, column int, parent *QModelIndex, stream *QDataStream) bool {
	__r0, _, _ := purego.SyscallN(addr_QConcatenateTablesProxyModel_decodeData, this.h, uintptr(row), uintptr(column), uintptr(parent.UnsafePointer()), uintptr(stream.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	runtime.KeepAlive(stream)
	return __r != 0
}

func (this *QConcatenateTablesProxyModel) BeginInsertRows(parent *QModelIndex, first int, last int) {
	purego.SyscallN(addr_QConcatenateTablesProxyModel_beginInsertRows, this.h, uintptr(parent.UnsafePointer()), uintptr(first), uintptr(last))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
}

func (this *QConcatenateTablesProxyModel) EndInsertRows() {
	purego.SyscallN(addr_QConcatenateTablesProxyModel_endInsertRows, this.h)
	runtime.KeepAlive(this)
}

func (this *QConcatenateTablesProxyModel) BeginRemoveRows(parent *QModelIndex, first int, last int) {
	purego.SyscallN(addr_QConcatenateTablesProxyModel_beginRemoveRows, this.h, uintptr(parent.UnsafePointer()), uintptr(first), uintptr(last))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
}

func (this *QConcatenateTablesProxyModel) EndRemoveRows() {
	purego.SyscallN(addr_QConcatenateTablesProxyModel_endRemoveRows, this.h)
	runtime.KeepAlive(this)
}

func (this *QConcatenateTablesProxyModel) BeginMoveRows(sourceParent *QModelIndex, sourceFirst int, sourceLast int, destinationParent *QModelIndex, destinationRow int) bool {
	__r0, _, _ := purego.SyscallN(addr_QConcatenateTablesProxyModel_beginMoveRows, this.h, uintptr(sourceParent.UnsafePointer()), uintptr(sourceFirst), uintptr(sourceLast), uintptr(destinationParent.UnsafePointer()), uintptr(destinationRow))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sourceParent)
	runtime.KeepAlive(destinationParent)
	return __r != 0
}

func (this *QConcatenateTablesProxyModel) EndMoveRows() {
	purego.SyscallN(addr_QConcatenateTablesProxyModel_endMoveRows, this.h)
	runtime.KeepAlive(this)
}

func (this *QConcatenateTablesProxyModel) BeginInsertColumns(parent *QModelIndex, first int, last int) {
	purego.SyscallN(addr_QConcatenateTablesProxyModel_beginInsertColumns, this.h, uintptr(parent.UnsafePointer()), uintptr(first), uintptr(last))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
}

func (this *QConcatenateTablesProxyModel) EndInsertColumns() {
	purego.SyscallN(addr_QConcatenateTablesProxyModel_endInsertColumns, this.h)
	runtime.KeepAlive(this)
}

func (this *QConcatenateTablesProxyModel) BeginRemoveColumns(parent *QModelIndex, first int, last int) {
	purego.SyscallN(addr_QConcatenateTablesProxyModel_beginRemoveColumns, this.h, uintptr(parent.UnsafePointer()), uintptr(first), uintptr(last))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
}

func (this *QConcatenateTablesProxyModel) EndRemoveColumns() {
	purego.SyscallN(addr_QConcatenateTablesProxyModel_endRemoveColumns, this.h)
	runtime.KeepAlive(this)
}

func (this *QConcatenateTablesProxyModel) BeginMoveColumns(sourceParent *QModelIndex, sourceFirst int, sourceLast int, destinationParent *QModelIndex, destinationColumn int) bool {
	__r0, _, _ := purego.SyscallN(addr_QConcatenateTablesProxyModel_beginMoveColumns, this.h, uintptr(sourceParent.UnsafePointer()), uintptr(sourceFirst), uintptr(sourceLast), uintptr(destinationParent.UnsafePointer()), uintptr(destinationColumn))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sourceParent)
	runtime.KeepAlive(destinationParent)
	return __r != 0
}

func (this *QConcatenateTablesProxyModel) EndMoveColumns() {
	purego.SyscallN(addr_QConcatenateTablesProxyModel_endMoveColumns, this.h)
	runtime.KeepAlive(this)
}

func (this *QConcatenateTablesProxyModel) BeginResetModel() {
	purego.SyscallN(addr_QConcatenateTablesProxyModel_beginResetModel, this.h)
	runtime.KeepAlive(this)
}

func (this *QConcatenateTablesProxyModel) EndResetModel() {
	purego.SyscallN(addr_QConcatenateTablesProxyModel_endResetModel, this.h)
	runtime.KeepAlive(this)
}

func (this *QConcatenateTablesProxyModel) ChangePersistentIndex(from *QModelIndex, to *QModelIndex) {
	purego.SyscallN(addr_QConcatenateTablesProxyModel_changePersistentIndex, this.h, uintptr(from.UnsafePointer()), uintptr(to.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(from)
	runtime.KeepAlive(to)
}

func (this *QConcatenateTablesProxyModel) ChangePersistentIndexList(from []*QModelIndex, to []*QModelIndex) {
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
	purego.SyscallN(addr_QConcatenateTablesProxyModel_changePersistentIndexList, this.h, from_p, uintptr(len(from)), to_p, uintptr(len(to)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(from)
	runtime.KeepAlive(from_arr)
	runtime.KeepAlive(to)
	runtime.KeepAlive(to_arr)
}

func (this *QConcatenateTablesProxyModel) PersistentIndexList() []*QModelIndex {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QConcatenateTablesProxyModel_persistentIndexList, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QModelIndex, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQModelIndex(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QConcatenateTablesProxyModel_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QConcatenateTablesProxyModel_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QConcatenateTablesProxyModel_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QConcatenateTablesProxyModel_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QConcatenateTablesProxyModel) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QConcatenateTablesProxyModel_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QConcatenateTablesProxyModel) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QConcatenateTablesProxyModel_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QConcatenateTablesProxyModel) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QConcatenateTablesProxyModel_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QConcatenateTablesProxyModel) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QConcatenateTablesProxyModel_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QConcatenateTablesProxyModel_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QConcatenateTablesProxyModel_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QConcatenateTablesProxyModel_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QConcatenateTablesProxyModel_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QConcatenateTablesProxyModel_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QConcatenateTablesProxyModel_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QConcatenateTablesProxyModel_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QConcatenateTablesProxyModel_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QConcatenateTablesProxyModel) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QConcatenateTablesProxyModel_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QConcatenateTablesProxyModel) DataDefault(index *QModelIndex, role int) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QConcatenateTablesProxyModel_data_default, this.h, uintptr(index.UnsafePointer()), uintptr(role))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return newownedQVariant(__r)
}

func (this *QConcatenateTablesProxyModel) SetDataDefault(index *QModelIndex, value *QVariant, role int) bool {
	__r0, _, _ := purego.SyscallN(addr_QConcatenateTablesProxyModel_setData_default, this.h, uintptr(index.UnsafePointer()), uintptr(value.UnsafePointer()), uintptr(role))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	runtime.KeepAlive(value)
	return __r != 0
}

func (this *QConcatenateTablesProxyModel) FlagsDefault(index *QModelIndex) ItemFlag {
	__r0, _, _ := purego.SyscallN(addr_QConcatenateTablesProxyModel_flags_default, this.h, uintptr(index.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return ItemFlag(__r)
}

func (this *QConcatenateTablesProxyModel) IndexDefault(row int, column int, parent *QModelIndex) *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QConcatenateTablesProxyModel_index_default, this.h, uintptr(row), uintptr(column), uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return newownedQModelIndex(__r)
}

func (this *QConcatenateTablesProxyModel) ParentDefault(index *QModelIndex) *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QConcatenateTablesProxyModel_parent_default, this.h, uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return newownedQModelIndex(__r)
}

func (this *QConcatenateTablesProxyModel) RowCountDefault(parent *QModelIndex) int {
	__r0, _, _ := purego.SyscallN(addr_QConcatenateTablesProxyModel_rowCount_default, this.h, uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return int(__r)
}

func (this *QConcatenateTablesProxyModel) HeaderDataDefault(section int, orientation Orientation, role int) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QConcatenateTablesProxyModel_headerData_default, this.h, uintptr(section), uintptr(orientation), uintptr(role))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QConcatenateTablesProxyModel) ColumnCountDefault(parent *QModelIndex) int {
	__r0, _, _ := purego.SyscallN(addr_QConcatenateTablesProxyModel_columnCount_default, this.h, uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return int(__r)
}

func (this *QConcatenateTablesProxyModel) MimeTypesDefault() []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QConcatenateTablesProxyModel_mimeTypes_default, this.h, uintptr(unsafe.Pointer(&__n)))
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

func (this *QConcatenateTablesProxyModel) MimeDataDefault(indexes []*QModelIndex) *QMimeData {
	indexes_arr := make([]uintptr, len(indexes))
	for __i, __e := range indexes {
		indexes_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var indexes_p uintptr
	if len(indexes_arr) > 0 {
		indexes_p = uintptr(unsafe.Pointer(&indexes_arr[0]))
	}
	__r, _, _ := purego.SyscallN(addr_QConcatenateTablesProxyModel_mimeData_default, this.h, indexes_p, uintptr(len(indexes)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(indexes)
	runtime.KeepAlive(indexes_arr)
	return bindQMimeData(__r)
}

func (this *QConcatenateTablesProxyModel) CanDropMimeDataDefault(data *QMimeData, action DropAction, row int, column int, parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QConcatenateTablesProxyModel_canDropMimeData_default, this.h, uintptr(data.UnsafePointer()), uintptr(action), uintptr(row), uintptr(column), uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(data)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QConcatenateTablesProxyModel) DropMimeDataDefault(data *QMimeData, action DropAction, row int, column int, parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QConcatenateTablesProxyModel_dropMimeData_default, this.h, uintptr(data.UnsafePointer()), uintptr(action), uintptr(row), uintptr(column), uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(data)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QConcatenateTablesProxyModel) SpanDefault(index *QModelIndex) *QSize {
	__r, _, _ := purego.SyscallN(addr_QConcatenateTablesProxyModel_span_default, this.h, uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return newownedQSize(__r)
}

func (this *QConcatenateTablesProxyModel) SiblingDefault(row int, column int, idx *QModelIndex) *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QConcatenateTablesProxyModel_sibling_default, this.h, uintptr(row), uintptr(column), uintptr(idx.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(idx)
	return newownedQModelIndex(__r)
}

func (this *QConcatenateTablesProxyModel) HasChildrenDefault(parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QConcatenateTablesProxyModel_hasChildren_default, this.h, uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QConcatenateTablesProxyModel) SetHeaderDataDefault(section int, orientation Orientation, value *QVariant, role int) bool {
	__r0, _, _ := purego.SyscallN(addr_QConcatenateTablesProxyModel_setHeaderData_default, this.h, uintptr(section), uintptr(orientation), uintptr(value.UnsafePointer()), uintptr(role))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(value)
	return __r != 0
}

func (this *QConcatenateTablesProxyModel) ClearItemDataDefault(index *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QConcatenateTablesProxyModel_clearItemData_default, this.h, uintptr(index.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return __r != 0
}

func (this *QConcatenateTablesProxyModel) SupportedDropActionsDefault() DropAction {
	__r0, _, _ := purego.SyscallN(addr_QConcatenateTablesProxyModel_supportedDropActions_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return DropAction(__r)
}

func (this *QConcatenateTablesProxyModel) SupportedDragActionsDefault() DropAction {
	__r0, _, _ := purego.SyscallN(addr_QConcatenateTablesProxyModel_supportedDragActions_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return DropAction(__r)
}

func (this *QConcatenateTablesProxyModel) InsertRowsDefault(row int, count int, parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QConcatenateTablesProxyModel_insertRows_default, this.h, uintptr(row), uintptr(count), uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QConcatenateTablesProxyModel) InsertColumnsDefault(column int, count int, parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QConcatenateTablesProxyModel_insertColumns_default, this.h, uintptr(column), uintptr(count), uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QConcatenateTablesProxyModel) RemoveRowsDefault(row int, count int, parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QConcatenateTablesProxyModel_removeRows_default, this.h, uintptr(row), uintptr(count), uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QConcatenateTablesProxyModel) RemoveColumnsDefault(column int, count int, parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QConcatenateTablesProxyModel_removeColumns_default, this.h, uintptr(column), uintptr(count), uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QConcatenateTablesProxyModel) MoveRowsDefault(sourceParent *QModelIndex, sourceRow int, count int, destinationParent *QModelIndex, destinationChild int) bool {
	__r0, _, _ := purego.SyscallN(addr_QConcatenateTablesProxyModel_moveRows_default, this.h, uintptr(sourceParent.UnsafePointer()), uintptr(sourceRow), uintptr(count), uintptr(destinationParent.UnsafePointer()), uintptr(destinationChild))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sourceParent)
	runtime.KeepAlive(destinationParent)
	return __r != 0
}

func (this *QConcatenateTablesProxyModel) MoveColumnsDefault(sourceParent *QModelIndex, sourceColumn int, count int, destinationParent *QModelIndex, destinationChild int) bool {
	__r0, _, _ := purego.SyscallN(addr_QConcatenateTablesProxyModel_moveColumns_default, this.h, uintptr(sourceParent.UnsafePointer()), uintptr(sourceColumn), uintptr(count), uintptr(destinationParent.UnsafePointer()), uintptr(destinationChild))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sourceParent)
	runtime.KeepAlive(destinationParent)
	return __r != 0
}

func (this *QConcatenateTablesProxyModel) FetchMoreDefault(parent *QModelIndex) {
	purego.SyscallN(addr_QConcatenateTablesProxyModel_fetchMore_default, this.h, uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
}

func (this *QConcatenateTablesProxyModel) CanFetchMoreDefault(parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QConcatenateTablesProxyModel_canFetchMore_default, this.h, uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QConcatenateTablesProxyModel) SortDefault(column int, order SortOrder) {
	purego.SyscallN(addr_QConcatenateTablesProxyModel_sort_default, this.h, uintptr(column), uintptr(order))
	runtime.KeepAlive(this)
}

func (this *QConcatenateTablesProxyModel) BuddyDefault(index *QModelIndex) *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QConcatenateTablesProxyModel_buddy_default, this.h, uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return newownedQModelIndex(__r)
}

func (this *QConcatenateTablesProxyModel) MatchDefault(start *QModelIndex, role int, value *QVariant, hits int, flags MatchFlag) []*QModelIndex {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QConcatenateTablesProxyModel_match_default, this.h, uintptr(start.UnsafePointer()), uintptr(role), uintptr(value.UnsafePointer()), uintptr(hits), uintptr(flags), uintptr(unsafe.Pointer(&__n)))
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

func (this *QConcatenateTablesProxyModel) MultiDataDefault(index *QModelIndex, roleDataSpan *QModelRoleDataSpan) {
	purego.SyscallN(addr_QConcatenateTablesProxyModel_multiData_default, this.h, uintptr(index.UnsafePointer()), uintptr(roleDataSpan.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	runtime.KeepAlive(roleDataSpan)
}

func (this *QConcatenateTablesProxyModel) SubmitDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QConcatenateTablesProxyModel_submit_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QConcatenateTablesProxyModel) RevertDefault() {
	purego.SyscallN(addr_QConcatenateTablesProxyModel_revert_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QConcatenateTablesProxyModel) ResetInternalDataDefault() {
	purego.SyscallN(addr_QConcatenateTablesProxyModel_resetInternalData_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QConcatenateTablesProxyModel) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QConcatenateTablesProxyModel_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QConcatenateTablesProxyModel) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QConcatenateTablesProxyModel_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QConcatenateTablesProxyModel) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QConcatenateTablesProxyModel_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QConcatenateTablesProxyModel) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QConcatenateTablesProxyModel_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QConcatenateTablesProxyModel) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QConcatenateTablesProxyModel_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QConcatenateTablesProxyModel) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QConcatenateTablesProxyModel_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QConcatenateTablesProxyModel) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QConcatenateTablesProxyModel_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QConcatenateTablesProxyModel) OnHeaderDataChanged(fn func(Orientation, int, int)) int64 {
	id := registerCallback(func(orientation int32, first int32, last int32) { fn(Orientation(orientation), int(first), int(last)) })
	pg_QConcatenateTablesProxyModel_connect_headerDataChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QConcatenateTablesProxyModel) OnLayoutChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QConcatenateTablesProxyModel_connect_layoutChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QConcatenateTablesProxyModel) OnLayoutChanged2(fn func([]*QPersistentModelIndex, QAbstractItemModel__LayoutChangeHint)) int64 {
	id := registerCallback(func(parents uintptr, parents_len int64, hint int32) { fn(func() []*QPersistentModelIndex { if parents == 0 || parents_len == 0 { return nil }; __sl := unsafe.Slice((*unsafe.Pointer)(unsafe.Pointer(parents)), int(parents_len)); __o := make([]*QPersistentModelIndex, int(parents_len)); for __i := range __o { __o[__i] = UnsafeNewQPersistentModelIndex(__sl[__i]) }; return __o }(), QAbstractItemModel__LayoutChangeHint(hint)) })
	pg_QConcatenateTablesProxyModel_connect_layoutChanged2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QConcatenateTablesProxyModel) OnLayoutAboutToBeChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QConcatenateTablesProxyModel_connect_layoutAboutToBeChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QConcatenateTablesProxyModel) OnLayoutAboutToBeChanged2(fn func([]*QPersistentModelIndex, QAbstractItemModel__LayoutChangeHint)) int64 {
	id := registerCallback(func(parents uintptr, parents_len int64, hint int32) { fn(func() []*QPersistentModelIndex { if parents == 0 || parents_len == 0 { return nil }; __sl := unsafe.Slice((*unsafe.Pointer)(unsafe.Pointer(parents)), int(parents_len)); __o := make([]*QPersistentModelIndex, int(parents_len)); for __i := range __o { __o[__i] = UnsafeNewQPersistentModelIndex(__sl[__i]) }; return __o }(), QAbstractItemModel__LayoutChangeHint(hint)) })
	pg_QConcatenateTablesProxyModel_connect_layoutAboutToBeChanged2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QConcatenateTablesProxyModel) OnRowsAboutToBeInserted(fn func(*QModelIndex, int, int)) int64 {
	id := registerCallback(func(parent uintptr, first int32, last int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(first), int(last)) })
	pg_QConcatenateTablesProxyModel_connect_rowsAboutToBeInserted(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QConcatenateTablesProxyModel) OnRowsInserted(fn func(*QModelIndex, int, int)) int64 {
	id := registerCallback(func(parent uintptr, first int32, last int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(first), int(last)) })
	pg_QConcatenateTablesProxyModel_connect_rowsInserted(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QConcatenateTablesProxyModel) OnRowsAboutToBeRemoved(fn func(*QModelIndex, int, int)) int64 {
	id := registerCallback(func(parent uintptr, first int32, last int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(first), int(last)) })
	pg_QConcatenateTablesProxyModel_connect_rowsAboutToBeRemoved(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QConcatenateTablesProxyModel) OnRowsRemoved(fn func(*QModelIndex, int, int)) int64 {
	id := registerCallback(func(parent uintptr, first int32, last int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(first), int(last)) })
	pg_QConcatenateTablesProxyModel_connect_rowsRemoved(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QConcatenateTablesProxyModel) OnColumnsAboutToBeInserted(fn func(*QModelIndex, int, int)) int64 {
	id := registerCallback(func(parent uintptr, first int32, last int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(first), int(last)) })
	pg_QConcatenateTablesProxyModel_connect_columnsAboutToBeInserted(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QConcatenateTablesProxyModel) OnColumnsInserted(fn func(*QModelIndex, int, int)) int64 {
	id := registerCallback(func(parent uintptr, first int32, last int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(first), int(last)) })
	pg_QConcatenateTablesProxyModel_connect_columnsInserted(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QConcatenateTablesProxyModel) OnColumnsAboutToBeRemoved(fn func(*QModelIndex, int, int)) int64 {
	id := registerCallback(func(parent uintptr, first int32, last int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(first), int(last)) })
	pg_QConcatenateTablesProxyModel_connect_columnsAboutToBeRemoved(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QConcatenateTablesProxyModel) OnColumnsRemoved(fn func(*QModelIndex, int, int)) int64 {
	id := registerCallback(func(parent uintptr, first int32, last int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(first), int(last)) })
	pg_QConcatenateTablesProxyModel_connect_columnsRemoved(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QConcatenateTablesProxyModel) OnModelAboutToBeReset(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QConcatenateTablesProxyModel_connect_modelAboutToBeReset(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QConcatenateTablesProxyModel) OnModelReset(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QConcatenateTablesProxyModel_connect_modelReset(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QConcatenateTablesProxyModel) OnRowsAboutToBeMoved(fn func(*QModelIndex, int, int, *QModelIndex, int)) int64 {
	id := registerCallback(func(sourceParent uintptr, sourceStart int32, sourceEnd int32, destinationParent uintptr, destinationRow int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(sourceParent)), int(sourceStart), int(sourceEnd), UnsafeNewQModelIndex(unsafe.Pointer(destinationParent)), int(destinationRow)) })
	pg_QConcatenateTablesProxyModel_connect_rowsAboutToBeMoved(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QConcatenateTablesProxyModel) OnRowsMoved(fn func(*QModelIndex, int, int, *QModelIndex, int)) int64 {
	id := registerCallback(func(sourceParent uintptr, sourceStart int32, sourceEnd int32, destinationParent uintptr, destinationRow int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(sourceParent)), int(sourceStart), int(sourceEnd), UnsafeNewQModelIndex(unsafe.Pointer(destinationParent)), int(destinationRow)) })
	pg_QConcatenateTablesProxyModel_connect_rowsMoved(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QConcatenateTablesProxyModel) OnColumnsAboutToBeMoved(fn func(*QModelIndex, int, int, *QModelIndex, int)) int64 {
	id := registerCallback(func(sourceParent uintptr, sourceStart int32, sourceEnd int32, destinationParent uintptr, destinationColumn int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(sourceParent)), int(sourceStart), int(sourceEnd), UnsafeNewQModelIndex(unsafe.Pointer(destinationParent)), int(destinationColumn)) })
	pg_QConcatenateTablesProxyModel_connect_columnsAboutToBeMoved(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QConcatenateTablesProxyModel) OnColumnsMoved(fn func(*QModelIndex, int, int, *QModelIndex, int)) int64 {
	id := registerCallback(func(sourceParent uintptr, sourceStart int32, sourceEnd int32, destinationParent uintptr, destinationColumn int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(sourceParent)), int(sourceStart), int(sourceEnd), UnsafeNewQModelIndex(unsafe.Pointer(destinationParent)), int(destinationColumn)) })
	pg_QConcatenateTablesProxyModel_connect_columnsMoved(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QConcatenateTablesProxyModel) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QConcatenateTablesProxyModel_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QConcatenateTablesProxyModel) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QConcatenateTablesProxyModel_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QConcatenateTablesProxyModel) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QConcatenateTablesProxyModel_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QConcatenateTablesProxyModel) Disconnect(handle int64) {
	pg_QConcatenateTablesProxyModel__disconnect(handle)
	unregisterCallback(handle)
}

func QConcatenateTablesProxyModel_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case -8837375942098466299:
		return 1
	case 35191454575202219:
		return 2
	case 1703382833545304812:
		return 3
	case -8948777187306027381:
		return 4
	case 7371339085820260189:
		return 5
	case 1960960171956884010:
		return 6
	case -3297140021896693688:
		return 7
	case -8256963024669584834:
		return 8
	case 5464467312394730418:
		return 9
	case 5175297091024974749:
		return 10
	case 731539533644753052:
		return 11
	case -9049994684305682212:
		return 12
	case -8395342916587040543:
		return 13
	case 397329346813782521:
		return 14
	case 3369452259144514380:
		return 15
	case 4852200306590054850:
		return 16
	case 8914190500520542409:
		return 17
	case -7784642176186821205:
		return 18
	case 3421601386088570762:
		return 19
	case -2775729825962395603:
		return 20
	case -8973293846230792237:
		return 21
	case 5104072549182222656:
		return 22
	case 2078548737056540420:
		return 23
	case -6074819356947405415:
		return 24
	case -6321007591388309905:
		return 25
	case 3553294593611130806:
		return 26
	case 8639262066473964150:
		return 27
	case 6912441694453382161:
		return 28
	case 6112860112248786499:
		return 29
	case -4341501217981328394:
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

func QConcatenateTablesProxyModel_enable(h uintptr, idx int32) { pg_QConcatenateTablesProxyModel__set_override(h, idx) }

func (this *QConcatenateTablesProxyModel) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QConcatenateTablesProxyModel) OnData(fn func(*QModelIndex, int) *QVariant) {
	registerVirt(this.UnsafeGoID(), -8837375942098466299, func(index uintptr, role int32) uintptr { return uintptr(fn(UnsafeNewQModelIndex(unsafe.Pointer(index)), int(role)).UnsafePointer()) }) // slot "data"
}

func (this *QConcatenateTablesProxyModel) OnSetData(fn func(*QModelIndex, *QVariant, int) bool) {
	registerVirt(this.UnsafeGoID(), 35191454575202219, func(index uintptr, value uintptr, role int32) int32 { return b2i(fn(UnsafeNewQModelIndex(unsafe.Pointer(index)), UnsafeNewQVariant(unsafe.Pointer(value)), int(role))) }) // slot "setData"
}

func (this *QConcatenateTablesProxyModel) OnFlags(fn func(*QModelIndex) ItemFlag) {
	registerVirt(this.UnsafeGoID(), 1703382833545304812, func(index uintptr) int32 { return int32(fn(UnsafeNewQModelIndex(unsafe.Pointer(index)))) }) // slot "flags"
}

func (this *QConcatenateTablesProxyModel) OnIndex(fn func(int, int, *QModelIndex) *QModelIndex) {
	registerVirt(this.UnsafeGoID(), -8948777187306027381, func(row int32, column int32, parent uintptr) uintptr { return uintptr(fn(int(row), int(column), UnsafeNewQModelIndex(unsafe.Pointer(parent))).UnsafePointer()) }) // slot "index"
}

func (this *QConcatenateTablesProxyModel) OnParent(fn func(*QModelIndex) *QModelIndex) {
	registerVirt(this.UnsafeGoID(), 7371339085820260189, func(index uintptr) uintptr { return uintptr(fn(UnsafeNewQModelIndex(unsafe.Pointer(index))).UnsafePointer()) }) // slot "parent"
}

func (this *QConcatenateTablesProxyModel) OnRowCount(fn func(*QModelIndex) int) {
	registerVirt(this.UnsafeGoID(), 1960960171956884010, func(parent uintptr) int32 { return int32(fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)))) }) // slot "rowCount"
}

func (this *QConcatenateTablesProxyModel) OnHeaderData(fn func(int, Orientation, int) *QVariant) {
	registerVirt(this.UnsafeGoID(), -3297140021896693688, func(section int32, orientation int32, role int32) uintptr { return uintptr(fn(int(section), Orientation(orientation), int(role)).UnsafePointer()) }) // slot "headerData"
}

func (this *QConcatenateTablesProxyModel) OnColumnCount(fn func(*QModelIndex) int) {
	registerVirt(this.UnsafeGoID(), -8256963024669584834, func(parent uintptr) int32 { return int32(fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)))) }) // slot "columnCount"
}

func (this *QConcatenateTablesProxyModel) OnMimeTypes(fn func() []string) {
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

func (this *QConcatenateTablesProxyModel) OnMimeData(fn func([]*QModelIndex) *QMimeData) {
	registerVirt(this.UnsafeGoID(), 5175297091024974749, func(indexes uintptr, indexes_len int64) uintptr { return uintptr(fn(func() []*QModelIndex { if indexes == 0 || indexes_len == 0 { return nil }; __sl := unsafe.Slice((*unsafe.Pointer)(unsafe.Pointer(indexes)), int(indexes_len)); __o := make([]*QModelIndex, int(indexes_len)); for __i := range __o { __o[__i] = UnsafeNewQModelIndex(__sl[__i]) }; return __o }()).UnsafePointer()) }) // slot "mimeData"
}

func (this *QConcatenateTablesProxyModel) OnCanDropMimeData(fn func(*QMimeData, DropAction, int, int, *QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), 731539533644753052, func(data uintptr, action int32, row int32, column int32, parent uintptr) int32 { return b2i(fn(UnsafeNewQMimeData(unsafe.Pointer(data)), DropAction(action), int(row), int(column), UnsafeNewQModelIndex(unsafe.Pointer(parent)))) }) // slot "canDropMimeData"
}

func (this *QConcatenateTablesProxyModel) OnDropMimeData(fn func(*QMimeData, DropAction, int, int, *QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), -9049994684305682212, func(data uintptr, action int32, row int32, column int32, parent uintptr) int32 { return b2i(fn(UnsafeNewQMimeData(unsafe.Pointer(data)), DropAction(action), int(row), int(column), UnsafeNewQModelIndex(unsafe.Pointer(parent)))) }) // slot "dropMimeData"
}

func (this *QConcatenateTablesProxyModel) OnSpan(fn func(*QModelIndex) *QSize) {
	registerVirt(this.UnsafeGoID(), -8395342916587040543, func(index uintptr) uintptr { return uintptr(fn(UnsafeNewQModelIndex(unsafe.Pointer(index))).UnsafePointer()) }) // slot "span"
}

func (this *QConcatenateTablesProxyModel) OnSibling(fn func(int, int, *QModelIndex) *QModelIndex) {
	registerVirt(this.UnsafeGoID(), 397329346813782521, func(row int32, column int32, idx uintptr) uintptr { return uintptr(fn(int(row), int(column), UnsafeNewQModelIndex(unsafe.Pointer(idx))).UnsafePointer()) }) // slot "sibling"
}

func (this *QConcatenateTablesProxyModel) OnHasChildren(fn func(*QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), 3369452259144514380, func(parent uintptr) int32 { return b2i(fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)))) }) // slot "hasChildren"
}

func (this *QConcatenateTablesProxyModel) OnSetHeaderData(fn func(int, Orientation, *QVariant, int) bool) {
	registerVirt(this.UnsafeGoID(), 4852200306590054850, func(section int32, orientation int32, value uintptr, role int32) int32 { return b2i(fn(int(section), Orientation(orientation), UnsafeNewQVariant(unsafe.Pointer(value)), int(role))) }) // slot "setHeaderData"
}

func (this *QConcatenateTablesProxyModel) OnClearItemData(fn func(*QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), 8914190500520542409, func(index uintptr) int32 { return b2i(fn(UnsafeNewQModelIndex(unsafe.Pointer(index)))) }) // slot "clearItemData"
}

func (this *QConcatenateTablesProxyModel) OnSupportedDropActions(fn func() DropAction) {
	registerVirt(this.UnsafeGoID(), -7784642176186821205, func() int32 { return int32(fn()) }) // slot "supportedDropActions"
}

func (this *QConcatenateTablesProxyModel) OnSupportedDragActions(fn func() DropAction) {
	registerVirt(this.UnsafeGoID(), 3421601386088570762, func() int32 { return int32(fn()) }) // slot "supportedDragActions"
}

func (this *QConcatenateTablesProxyModel) OnInsertRows(fn func(int, int, *QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), -2775729825962395603, func(row int32, count int32, parent uintptr) int32 { return b2i(fn(int(row), int(count), UnsafeNewQModelIndex(unsafe.Pointer(parent)))) }) // slot "insertRows"
}

func (this *QConcatenateTablesProxyModel) OnInsertColumns(fn func(int, int, *QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), -8973293846230792237, func(column int32, count int32, parent uintptr) int32 { return b2i(fn(int(column), int(count), UnsafeNewQModelIndex(unsafe.Pointer(parent)))) }) // slot "insertColumns"
}

func (this *QConcatenateTablesProxyModel) OnRemoveRows(fn func(int, int, *QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), 5104072549182222656, func(row int32, count int32, parent uintptr) int32 { return b2i(fn(int(row), int(count), UnsafeNewQModelIndex(unsafe.Pointer(parent)))) }) // slot "removeRows"
}

func (this *QConcatenateTablesProxyModel) OnRemoveColumns(fn func(int, int, *QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), 2078548737056540420, func(column int32, count int32, parent uintptr) int32 { return b2i(fn(int(column), int(count), UnsafeNewQModelIndex(unsafe.Pointer(parent)))) }) // slot "removeColumns"
}

func (this *QConcatenateTablesProxyModel) OnMoveRows(fn func(*QModelIndex, int, int, *QModelIndex, int) bool) {
	registerVirt(this.UnsafeGoID(), -6074819356947405415, func(sourceParent uintptr, sourceRow int32, count int32, destinationParent uintptr, destinationChild int32) int32 { return b2i(fn(UnsafeNewQModelIndex(unsafe.Pointer(sourceParent)), int(sourceRow), int(count), UnsafeNewQModelIndex(unsafe.Pointer(destinationParent)), int(destinationChild))) }) // slot "moveRows"
}

func (this *QConcatenateTablesProxyModel) OnMoveColumns(fn func(*QModelIndex, int, int, *QModelIndex, int) bool) {
	registerVirt(this.UnsafeGoID(), -6321007591388309905, func(sourceParent uintptr, sourceColumn int32, count int32, destinationParent uintptr, destinationChild int32) int32 { return b2i(fn(UnsafeNewQModelIndex(unsafe.Pointer(sourceParent)), int(sourceColumn), int(count), UnsafeNewQModelIndex(unsafe.Pointer(destinationParent)), int(destinationChild))) }) // slot "moveColumns"
}

func (this *QConcatenateTablesProxyModel) OnFetchMore(fn func(*QModelIndex)) {
	registerVirt(this.UnsafeGoID(), 3553294593611130806, func(parent uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent))) }) // slot "fetchMore"
}

func (this *QConcatenateTablesProxyModel) OnCanFetchMore(fn func(*QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), 8639262066473964150, func(parent uintptr) int32 { return b2i(fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)))) }) // slot "canFetchMore"
}

func (this *QConcatenateTablesProxyModel) OnSort(fn func(int, SortOrder)) {
	registerVirt(this.UnsafeGoID(), 6912441694453382161, func(column int32, order int32) { fn(int(column), SortOrder(order)) }) // slot "sort"
}

func (this *QConcatenateTablesProxyModel) OnBuddy(fn func(*QModelIndex) *QModelIndex) {
	registerVirt(this.UnsafeGoID(), 6112860112248786499, func(index uintptr) uintptr { return uintptr(fn(UnsafeNewQModelIndex(unsafe.Pointer(index))).UnsafePointer()) }) // slot "buddy"
}

func (this *QConcatenateTablesProxyModel) OnMatch(fn func(*QModelIndex, int, *QVariant, int, MatchFlag) []*QModelIndex) {
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

func (this *QConcatenateTablesProxyModel) OnMultiData(fn func(*QModelIndex, *QModelRoleDataSpan)) {
	registerVirt(this.UnsafeGoID(), 731311060944096642, func(index uintptr, roleDataSpan uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(index)), UnsafeNewQModelRoleDataSpan(unsafe.Pointer(roleDataSpan))) }) // slot "multiData"
}

func (this *QConcatenateTablesProxyModel) OnSubmit(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -3788331537442685179, func() int32 { return b2i(fn()) }) // slot "submit"
}

func (this *QConcatenateTablesProxyModel) OnRevert(fn func()) {
	registerVirt(this.UnsafeGoID(), -1522951607536117149, func() { fn() }) // slot "revert"
}

func (this *QConcatenateTablesProxyModel) OnResetInternalData(fn func()) {
	registerVirt(this.UnsafeGoID(), -4510675770784051183, func() { fn() }) // slot "resetInternalData"
}

func (this *QConcatenateTablesProxyModel) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QConcatenateTablesProxyModel) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QConcatenateTablesProxyModel) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QConcatenateTablesProxyModel) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QConcatenateTablesProxyModel) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QConcatenateTablesProxyModel) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QConcatenateTablesProxyModel) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QConcatenateTablesProxyModel) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QConcatenateTablesProxyModel_delete(this.h)
	}
	this.UnsafeRelease()
}
