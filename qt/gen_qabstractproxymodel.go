package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QAbstractProxyModel_chain uintptr
	addr_QAbstractProxyModel_tr uintptr
	addr_QAbstractProxyModel_setSourceModel uintptr
	addr_QAbstractProxyModel_sourceModel uintptr
	addr_QAbstractProxyModel_mapToSource uintptr
	addr_QAbstractProxyModel_mapFromSource uintptr
	addr_QAbstractProxyModel_mapSelectionToSource uintptr
	addr_QAbstractProxyModel_mapSelectionFromSource uintptr
	addr_QAbstractProxyModel_createIndex uintptr
	addr_QAbstractProxyModel_createIndex2 uintptr
	addr_QAbstractProxyModel_encodeData uintptr
	addr_QAbstractProxyModel_decodeData uintptr
	addr_QAbstractProxyModel_beginInsertRows uintptr
	addr_QAbstractProxyModel_endInsertRows uintptr
	addr_QAbstractProxyModel_beginRemoveRows uintptr
	addr_QAbstractProxyModel_endRemoveRows uintptr
	addr_QAbstractProxyModel_beginMoveRows uintptr
	addr_QAbstractProxyModel_endMoveRows uintptr
	addr_QAbstractProxyModel_beginInsertColumns uintptr
	addr_QAbstractProxyModel_endInsertColumns uintptr
	addr_QAbstractProxyModel_beginRemoveColumns uintptr
	addr_QAbstractProxyModel_endRemoveColumns uintptr
	addr_QAbstractProxyModel_beginMoveColumns uintptr
	addr_QAbstractProxyModel_endMoveColumns uintptr
	addr_QAbstractProxyModel_beginResetModel uintptr
	addr_QAbstractProxyModel_endResetModel uintptr
	addr_QAbstractProxyModel_changePersistentIndex uintptr
	addr_QAbstractProxyModel_changePersistentIndexList uintptr
	addr_QAbstractProxyModel_persistentIndexList uintptr
	addr_QAbstractProxyModel_connect uintptr
	addr_QAbstractProxyModel_connect2 uintptr
	addr_QAbstractProxyModel_sender uintptr
	addr_QAbstractProxyModel_senderSignalIndex uintptr
	addr_QAbstractProxyModel_receivers uintptr
	addr_QAbstractProxyModel_isSignalConnected uintptr
	addr_QAbstractProxyModel_tr2 uintptr
	addr_QAbstractProxyModel_tr3 uintptr
	addr_QAbstractProxyModel_connect3 uintptr
	addr_QAbstractProxyModel_connect4 uintptr
	addr_QAbstractProxyModel_setSourceModel_default uintptr
	addr_QAbstractProxyModel_mapSelectionToSource_default uintptr
	addr_QAbstractProxyModel_mapSelectionFromSource_default uintptr
	addr_QAbstractProxyModel_metaObject_default uintptr
	addr_QAbstractProxyModel_submit_default uintptr
	addr_QAbstractProxyModel_revert_default uintptr
	addr_QAbstractProxyModel_data_default uintptr
	addr_QAbstractProxyModel_headerData_default uintptr
	addr_QAbstractProxyModel_flags_default uintptr
	addr_QAbstractProxyModel_setData_default uintptr
	addr_QAbstractProxyModel_setHeaderData_default uintptr
	addr_QAbstractProxyModel_clearItemData_default uintptr
	addr_QAbstractProxyModel_buddy_default uintptr
	addr_QAbstractProxyModel_canFetchMore_default uintptr
	addr_QAbstractProxyModel_fetchMore_default uintptr
	addr_QAbstractProxyModel_sort_default uintptr
	addr_QAbstractProxyModel_span_default uintptr
	addr_QAbstractProxyModel_hasChildren_default uintptr
	addr_QAbstractProxyModel_sibling_default uintptr
	addr_QAbstractProxyModel_mimeData_default uintptr
	addr_QAbstractProxyModel_canDropMimeData_default uintptr
	addr_QAbstractProxyModel_dropMimeData_default uintptr
	addr_QAbstractProxyModel_mimeTypes_default uintptr
	addr_QAbstractProxyModel_supportedDragActions_default uintptr
	addr_QAbstractProxyModel_supportedDropActions_default uintptr
	addr_QAbstractProxyModel_insertRows_default uintptr
	addr_QAbstractProxyModel_insertColumns_default uintptr
	addr_QAbstractProxyModel_removeRows_default uintptr
	addr_QAbstractProxyModel_removeColumns_default uintptr
	addr_QAbstractProxyModel_moveRows_default uintptr
	addr_QAbstractProxyModel_moveColumns_default uintptr
	addr_QAbstractProxyModel_match_default uintptr
	addr_QAbstractProxyModel_multiData_default uintptr
	addr_QAbstractProxyModel_resetInternalData_default uintptr
	addr_QAbstractProxyModel_event_default uintptr
	addr_QAbstractProxyModel_eventFilter_default uintptr
	addr_QAbstractProxyModel_timerEvent_default uintptr
	addr_QAbstractProxyModel_childEvent_default uintptr
	addr_QAbstractProxyModel_customEvent_default uintptr
	addr_QAbstractProxyModel_connectNotify_default uintptr
	addr_QAbstractProxyModel_disconnectNotify_default uintptr
	pg_QAbstractProxyModel_connect_sourceModelChanged func(uintptr, int64)
	pg_QAbstractProxyModel_connect_headerDataChanged func(uintptr, int64)
	pg_QAbstractProxyModel_connect_layoutChanged func(uintptr, int64)
	pg_QAbstractProxyModel_connect_layoutChanged2 func(uintptr, int64)
	pg_QAbstractProxyModel_connect_layoutAboutToBeChanged func(uintptr, int64)
	pg_QAbstractProxyModel_connect_layoutAboutToBeChanged2 func(uintptr, int64)
	pg_QAbstractProxyModel_connect_rowsAboutToBeInserted func(uintptr, int64)
	pg_QAbstractProxyModel_connect_rowsInserted func(uintptr, int64)
	pg_QAbstractProxyModel_connect_rowsAboutToBeRemoved func(uintptr, int64)
	pg_QAbstractProxyModel_connect_rowsRemoved func(uintptr, int64)
	pg_QAbstractProxyModel_connect_columnsAboutToBeInserted func(uintptr, int64)
	pg_QAbstractProxyModel_connect_columnsInserted func(uintptr, int64)
	pg_QAbstractProxyModel_connect_columnsAboutToBeRemoved func(uintptr, int64)
	pg_QAbstractProxyModel_connect_columnsRemoved func(uintptr, int64)
	pg_QAbstractProxyModel_connect_modelAboutToBeReset func(uintptr, int64)
	pg_QAbstractProxyModel_connect_modelReset func(uintptr, int64)
	pg_QAbstractProxyModel_connect_rowsAboutToBeMoved func(uintptr, int64)
	pg_QAbstractProxyModel_connect_rowsMoved func(uintptr, int64)
	pg_QAbstractProxyModel_connect_columnsAboutToBeMoved func(uintptr, int64)
	pg_QAbstractProxyModel_connect_columnsMoved func(uintptr, int64)
	pg_QAbstractProxyModel_connect_destroyed func(uintptr, int64)
	pg_QAbstractProxyModel_connect_destroyed2 func(uintptr, int64)
	pg_QAbstractProxyModel_connect_objectNameChanged func(uintptr, int64)
	pg_QAbstractProxyModel__disconnect func(int64)
	pg_QAbstractProxyModel__set_override func(uintptr, int32)
	pg_QAbstractProxyModel__Go_new func(int64, uintptr) uintptr
	pg_QAbstractProxyModel_delete func(uintptr)
)

func init() {
	symaddr(&addr_QAbstractProxyModel_chain, "QAbstractProxyModel_chain")
	symaddr(&addr_QAbstractProxyModel_tr, "QAbstractProxyModel_tr")
	symaddr(&addr_QAbstractProxyModel_setSourceModel, "QAbstractProxyModel_setSourceModel")
	symaddr(&addr_QAbstractProxyModel_sourceModel, "QAbstractProxyModel_sourceModel")
	symaddr(&addr_QAbstractProxyModel_mapToSource, "QAbstractProxyModel_mapToSource")
	symaddr(&addr_QAbstractProxyModel_mapFromSource, "QAbstractProxyModel_mapFromSource")
	symaddr(&addr_QAbstractProxyModel_mapSelectionToSource, "QAbstractProxyModel_mapSelectionToSource")
	symaddr(&addr_QAbstractProxyModel_mapSelectionFromSource, "QAbstractProxyModel_mapSelectionFromSource")
	symaddr(&addr_QAbstractProxyModel_createIndex, "QAbstractProxyModel_createIndex")
	symaddr(&addr_QAbstractProxyModel_createIndex2, "QAbstractProxyModel_createIndex2")
	symaddr(&addr_QAbstractProxyModel_encodeData, "QAbstractProxyModel_encodeData")
	symaddr(&addr_QAbstractProxyModel_decodeData, "QAbstractProxyModel_decodeData")
	symaddr(&addr_QAbstractProxyModel_beginInsertRows, "QAbstractProxyModel_beginInsertRows")
	symaddr(&addr_QAbstractProxyModel_endInsertRows, "QAbstractProxyModel_endInsertRows")
	symaddr(&addr_QAbstractProxyModel_beginRemoveRows, "QAbstractProxyModel_beginRemoveRows")
	symaddr(&addr_QAbstractProxyModel_endRemoveRows, "QAbstractProxyModel_endRemoveRows")
	symaddr(&addr_QAbstractProxyModel_beginMoveRows, "QAbstractProxyModel_beginMoveRows")
	symaddr(&addr_QAbstractProxyModel_endMoveRows, "QAbstractProxyModel_endMoveRows")
	symaddr(&addr_QAbstractProxyModel_beginInsertColumns, "QAbstractProxyModel_beginInsertColumns")
	symaddr(&addr_QAbstractProxyModel_endInsertColumns, "QAbstractProxyModel_endInsertColumns")
	symaddr(&addr_QAbstractProxyModel_beginRemoveColumns, "QAbstractProxyModel_beginRemoveColumns")
	symaddr(&addr_QAbstractProxyModel_endRemoveColumns, "QAbstractProxyModel_endRemoveColumns")
	symaddr(&addr_QAbstractProxyModel_beginMoveColumns, "QAbstractProxyModel_beginMoveColumns")
	symaddr(&addr_QAbstractProxyModel_endMoveColumns, "QAbstractProxyModel_endMoveColumns")
	symaddr(&addr_QAbstractProxyModel_beginResetModel, "QAbstractProxyModel_beginResetModel")
	symaddr(&addr_QAbstractProxyModel_endResetModel, "QAbstractProxyModel_endResetModel")
	symaddr(&addr_QAbstractProxyModel_changePersistentIndex, "QAbstractProxyModel_changePersistentIndex")
	symaddr(&addr_QAbstractProxyModel_changePersistentIndexList, "QAbstractProxyModel_changePersistentIndexList")
	symaddr(&addr_QAbstractProxyModel_persistentIndexList, "QAbstractProxyModel_persistentIndexList")
	symaddr(&addr_QAbstractProxyModel_connect, "QAbstractProxyModel_connect")
	symaddr(&addr_QAbstractProxyModel_connect2, "QAbstractProxyModel_connect2")
	symaddr(&addr_QAbstractProxyModel_sender, "QAbstractProxyModel_sender")
	symaddr(&addr_QAbstractProxyModel_senderSignalIndex, "QAbstractProxyModel_senderSignalIndex")
	symaddr(&addr_QAbstractProxyModel_receivers, "QAbstractProxyModel_receivers")
	symaddr(&addr_QAbstractProxyModel_isSignalConnected, "QAbstractProxyModel_isSignalConnected")
	symaddr(&addr_QAbstractProxyModel_tr2, "QAbstractProxyModel_tr2")
	symaddr(&addr_QAbstractProxyModel_tr3, "QAbstractProxyModel_tr3")
	symaddr(&addr_QAbstractProxyModel_connect3, "QAbstractProxyModel_connect3")
	symaddr(&addr_QAbstractProxyModel_connect4, "QAbstractProxyModel_connect4")
	symaddr(&addr_QAbstractProxyModel_setSourceModel_default, "QAbstractProxyModel_setSourceModel_default")
	symaddr(&addr_QAbstractProxyModel_mapSelectionToSource_default, "QAbstractProxyModel_mapSelectionToSource_default")
	symaddr(&addr_QAbstractProxyModel_mapSelectionFromSource_default, "QAbstractProxyModel_mapSelectionFromSource_default")
	symaddr(&addr_QAbstractProxyModel_metaObject_default, "QAbstractProxyModel_metaObject_default")
	symaddr(&addr_QAbstractProxyModel_submit_default, "QAbstractProxyModel_submit_default")
	symaddr(&addr_QAbstractProxyModel_revert_default, "QAbstractProxyModel_revert_default")
	symaddr(&addr_QAbstractProxyModel_data_default, "QAbstractProxyModel_data_default")
	symaddr(&addr_QAbstractProxyModel_headerData_default, "QAbstractProxyModel_headerData_default")
	symaddr(&addr_QAbstractProxyModel_flags_default, "QAbstractProxyModel_flags_default")
	symaddr(&addr_QAbstractProxyModel_setData_default, "QAbstractProxyModel_setData_default")
	symaddr(&addr_QAbstractProxyModel_setHeaderData_default, "QAbstractProxyModel_setHeaderData_default")
	symaddr(&addr_QAbstractProxyModel_clearItemData_default, "QAbstractProxyModel_clearItemData_default")
	symaddr(&addr_QAbstractProxyModel_buddy_default, "QAbstractProxyModel_buddy_default")
	symaddr(&addr_QAbstractProxyModel_canFetchMore_default, "QAbstractProxyModel_canFetchMore_default")
	symaddr(&addr_QAbstractProxyModel_fetchMore_default, "QAbstractProxyModel_fetchMore_default")
	symaddr(&addr_QAbstractProxyModel_sort_default, "QAbstractProxyModel_sort_default")
	symaddr(&addr_QAbstractProxyModel_span_default, "QAbstractProxyModel_span_default")
	symaddr(&addr_QAbstractProxyModel_hasChildren_default, "QAbstractProxyModel_hasChildren_default")
	symaddr(&addr_QAbstractProxyModel_sibling_default, "QAbstractProxyModel_sibling_default")
	symaddr(&addr_QAbstractProxyModel_mimeData_default, "QAbstractProxyModel_mimeData_default")
	symaddr(&addr_QAbstractProxyModel_canDropMimeData_default, "QAbstractProxyModel_canDropMimeData_default")
	symaddr(&addr_QAbstractProxyModel_dropMimeData_default, "QAbstractProxyModel_dropMimeData_default")
	symaddr(&addr_QAbstractProxyModel_mimeTypes_default, "QAbstractProxyModel_mimeTypes_default")
	symaddr(&addr_QAbstractProxyModel_supportedDragActions_default, "QAbstractProxyModel_supportedDragActions_default")
	symaddr(&addr_QAbstractProxyModel_supportedDropActions_default, "QAbstractProxyModel_supportedDropActions_default")
	symaddr(&addr_QAbstractProxyModel_insertRows_default, "QAbstractProxyModel_insertRows_default")
	symaddr(&addr_QAbstractProxyModel_insertColumns_default, "QAbstractProxyModel_insertColumns_default")
	symaddr(&addr_QAbstractProxyModel_removeRows_default, "QAbstractProxyModel_removeRows_default")
	symaddr(&addr_QAbstractProxyModel_removeColumns_default, "QAbstractProxyModel_removeColumns_default")
	symaddr(&addr_QAbstractProxyModel_moveRows_default, "QAbstractProxyModel_moveRows_default")
	symaddr(&addr_QAbstractProxyModel_moveColumns_default, "QAbstractProxyModel_moveColumns_default")
	symaddr(&addr_QAbstractProxyModel_match_default, "QAbstractProxyModel_match_default")
	symaddr(&addr_QAbstractProxyModel_multiData_default, "QAbstractProxyModel_multiData_default")
	symaddr(&addr_QAbstractProxyModel_resetInternalData_default, "QAbstractProxyModel_resetInternalData_default")
	symaddr(&addr_QAbstractProxyModel_event_default, "QAbstractProxyModel_event_default")
	symaddr(&addr_QAbstractProxyModel_eventFilter_default, "QAbstractProxyModel_eventFilter_default")
	symaddr(&addr_QAbstractProxyModel_timerEvent_default, "QAbstractProxyModel_timerEvent_default")
	symaddr(&addr_QAbstractProxyModel_childEvent_default, "QAbstractProxyModel_childEvent_default")
	symaddr(&addr_QAbstractProxyModel_customEvent_default, "QAbstractProxyModel_customEvent_default")
	symaddr(&addr_QAbstractProxyModel_connectNotify_default, "QAbstractProxyModel_connectNotify_default")
	symaddr(&addr_QAbstractProxyModel_disconnectNotify_default, "QAbstractProxyModel_disconnectNotify_default")
	reg(&pg_QAbstractProxyModel_connect_sourceModelChanged, "QAbstractProxyModel_connect_sourceModelChanged")
	reg(&pg_QAbstractProxyModel_connect_headerDataChanged, "QAbstractProxyModel_connect_headerDataChanged")
	reg(&pg_QAbstractProxyModel_connect_layoutChanged, "QAbstractProxyModel_connect_layoutChanged")
	reg(&pg_QAbstractProxyModel_connect_layoutChanged2, "QAbstractProxyModel_connect_layoutChanged2")
	reg(&pg_QAbstractProxyModel_connect_layoutAboutToBeChanged, "QAbstractProxyModel_connect_layoutAboutToBeChanged")
	reg(&pg_QAbstractProxyModel_connect_layoutAboutToBeChanged2, "QAbstractProxyModel_connect_layoutAboutToBeChanged2")
	reg(&pg_QAbstractProxyModel_connect_rowsAboutToBeInserted, "QAbstractProxyModel_connect_rowsAboutToBeInserted")
	reg(&pg_QAbstractProxyModel_connect_rowsInserted, "QAbstractProxyModel_connect_rowsInserted")
	reg(&pg_QAbstractProxyModel_connect_rowsAboutToBeRemoved, "QAbstractProxyModel_connect_rowsAboutToBeRemoved")
	reg(&pg_QAbstractProxyModel_connect_rowsRemoved, "QAbstractProxyModel_connect_rowsRemoved")
	reg(&pg_QAbstractProxyModel_connect_columnsAboutToBeInserted, "QAbstractProxyModel_connect_columnsAboutToBeInserted")
	reg(&pg_QAbstractProxyModel_connect_columnsInserted, "QAbstractProxyModel_connect_columnsInserted")
	reg(&pg_QAbstractProxyModel_connect_columnsAboutToBeRemoved, "QAbstractProxyModel_connect_columnsAboutToBeRemoved")
	reg(&pg_QAbstractProxyModel_connect_columnsRemoved, "QAbstractProxyModel_connect_columnsRemoved")
	reg(&pg_QAbstractProxyModel_connect_modelAboutToBeReset, "QAbstractProxyModel_connect_modelAboutToBeReset")
	reg(&pg_QAbstractProxyModel_connect_modelReset, "QAbstractProxyModel_connect_modelReset")
	reg(&pg_QAbstractProxyModel_connect_rowsAboutToBeMoved, "QAbstractProxyModel_connect_rowsAboutToBeMoved")
	reg(&pg_QAbstractProxyModel_connect_rowsMoved, "QAbstractProxyModel_connect_rowsMoved")
	reg(&pg_QAbstractProxyModel_connect_columnsAboutToBeMoved, "QAbstractProxyModel_connect_columnsAboutToBeMoved")
	reg(&pg_QAbstractProxyModel_connect_columnsMoved, "QAbstractProxyModel_connect_columnsMoved")
	reg(&pg_QAbstractProxyModel_connect_destroyed, "QAbstractProxyModel_connect_destroyed")
	reg(&pg_QAbstractProxyModel_connect_destroyed2, "QAbstractProxyModel_connect_destroyed2")
	reg(&pg_QAbstractProxyModel_connect_objectNameChanged, "QAbstractProxyModel_connect_objectNameChanged")
	reg(&pg_QAbstractProxyModel__disconnect, "QAbstractProxyModel__disconnect")
	reg(&pg_QAbstractProxyModel__set_override, "QAbstractProxyModel__set_override")
	reg(&pg_QAbstractProxyModel__Go_new, "QAbstractProxyModel__Go_new")
	reg(&pg_QAbstractProxyModel_delete, "QAbstractProxyModel_delete")
}

type QAbstractProxyModel struct {
	h uintptr
	*QAbstractItemModel
}

func bindQAbstractProxyModel(h uintptr) *QAbstractProxyModel {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QAbstractProxyModel_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QAbstractProxyModel{h: h, QAbstractItemModel: &QAbstractItemModel{h: __c[0], QObject: &QObject{h: __c[1]}}}
}

func newownedQAbstractProxyModel(h uintptr) *QAbstractProxyModel { return bindQAbstractProxyModel(h) }

func (this *QAbstractProxyModel) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QAbstractProxyModel) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQAbstractProxyModel(h unsafe.Pointer) *QAbstractProxyModel { return bindQAbstractProxyModel(uintptr(h)) }
func UnsafeNewQAbstractProxyModelOwned(h unsafe.Pointer) *QAbstractProxyModel { return newownedQAbstractProxyModel(uintptr(h)) }

func (this *QAbstractProxyModel) UnsafeRelease() {
	this.h = 0
	if this.QAbstractItemModel != nil {
		this.QAbstractItemModel.UnsafeRelease()
	}
}

func QAbstractProxyModel_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAbstractProxyModel_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QAbstractProxyModel) SetSourceModel(sourceModel *QAbstractItemModel) {
	purego.SyscallN(addr_QAbstractProxyModel_setSourceModel, this.h, uintptr(sourceModel.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sourceModel)
}

func (this *QAbstractProxyModel) SourceModel() *QAbstractItemModel {
	__r, _, _ := purego.SyscallN(addr_QAbstractProxyModel_sourceModel, this.h)
	runtime.KeepAlive(this)
	return bindQAbstractItemModel(__r)
}

func (this *QAbstractProxyModel) MapToSource(proxyIndex *QModelIndex) *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QAbstractProxyModel_mapToSource, this.h, uintptr(proxyIndex.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(proxyIndex)
	return newownedQModelIndex(__r)
}

func (this *QAbstractProxyModel) MapFromSource(sourceIndex *QModelIndex) *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QAbstractProxyModel_mapFromSource, this.h, uintptr(sourceIndex.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sourceIndex)
	return newownedQModelIndex(__r)
}

func (this *QAbstractProxyModel) MapSelectionToSource(selection *QItemSelection) *QItemSelection {
	__r, _, _ := purego.SyscallN(addr_QAbstractProxyModel_mapSelectionToSource, this.h, uintptr(selection.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(selection)
	return newownedQItemSelection(__r)
}

func (this *QAbstractProxyModel) MapSelectionFromSource(selection *QItemSelection) *QItemSelection {
	__r, _, _ := purego.SyscallN(addr_QAbstractProxyModel_mapSelectionFromSource, this.h, uintptr(selection.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(selection)
	return newownedQItemSelection(__r)
}

func (this *QAbstractProxyModel) CreateIndex(row int, column int) *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QAbstractProxyModel_createIndex, this.h, uintptr(row), uintptr(column))
	runtime.KeepAlive(this)
	return newownedQModelIndex(__r)
}

func (this *QAbstractProxyModel) CreateIndex2(row int, column int, id uintptr) *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QAbstractProxyModel_createIndex2, this.h, uintptr(row), uintptr(column), uintptr(id))
	runtime.KeepAlive(this)
	return newownedQModelIndex(__r)
}

func (this *QAbstractProxyModel) EncodeData(indexes []*QModelIndex, stream *QDataStream) {
	indexes_arr := make([]uintptr, len(indexes))
	for __i, __e := range indexes {
		indexes_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var indexes_p uintptr
	if len(indexes_arr) > 0 {
		indexes_p = uintptr(unsafe.Pointer(&indexes_arr[0]))
	}
	purego.SyscallN(addr_QAbstractProxyModel_encodeData, this.h, indexes_p, uintptr(len(indexes)), uintptr(stream.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(indexes)
	runtime.KeepAlive(indexes_arr)
	runtime.KeepAlive(stream)
}

func (this *QAbstractProxyModel) DecodeData(row int, column int, parent *QModelIndex, stream *QDataStream) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractProxyModel_decodeData, this.h, uintptr(row), uintptr(column), uintptr(parent.UnsafePointer()), uintptr(stream.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	runtime.KeepAlive(stream)
	return __r != 0
}

func (this *QAbstractProxyModel) BeginInsertRows(parent *QModelIndex, first int, last int) {
	purego.SyscallN(addr_QAbstractProxyModel_beginInsertRows, this.h, uintptr(parent.UnsafePointer()), uintptr(first), uintptr(last))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
}

func (this *QAbstractProxyModel) EndInsertRows() {
	purego.SyscallN(addr_QAbstractProxyModel_endInsertRows, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractProxyModel) BeginRemoveRows(parent *QModelIndex, first int, last int) {
	purego.SyscallN(addr_QAbstractProxyModel_beginRemoveRows, this.h, uintptr(parent.UnsafePointer()), uintptr(first), uintptr(last))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
}

func (this *QAbstractProxyModel) EndRemoveRows() {
	purego.SyscallN(addr_QAbstractProxyModel_endRemoveRows, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractProxyModel) BeginMoveRows(sourceParent *QModelIndex, sourceFirst int, sourceLast int, destinationParent *QModelIndex, destinationRow int) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractProxyModel_beginMoveRows, this.h, uintptr(sourceParent.UnsafePointer()), uintptr(sourceFirst), uintptr(sourceLast), uintptr(destinationParent.UnsafePointer()), uintptr(destinationRow))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sourceParent)
	runtime.KeepAlive(destinationParent)
	return __r != 0
}

func (this *QAbstractProxyModel) EndMoveRows() {
	purego.SyscallN(addr_QAbstractProxyModel_endMoveRows, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractProxyModel) BeginInsertColumns(parent *QModelIndex, first int, last int) {
	purego.SyscallN(addr_QAbstractProxyModel_beginInsertColumns, this.h, uintptr(parent.UnsafePointer()), uintptr(first), uintptr(last))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
}

func (this *QAbstractProxyModel) EndInsertColumns() {
	purego.SyscallN(addr_QAbstractProxyModel_endInsertColumns, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractProxyModel) BeginRemoveColumns(parent *QModelIndex, first int, last int) {
	purego.SyscallN(addr_QAbstractProxyModel_beginRemoveColumns, this.h, uintptr(parent.UnsafePointer()), uintptr(first), uintptr(last))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
}

func (this *QAbstractProxyModel) EndRemoveColumns() {
	purego.SyscallN(addr_QAbstractProxyModel_endRemoveColumns, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractProxyModel) BeginMoveColumns(sourceParent *QModelIndex, sourceFirst int, sourceLast int, destinationParent *QModelIndex, destinationColumn int) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractProxyModel_beginMoveColumns, this.h, uintptr(sourceParent.UnsafePointer()), uintptr(sourceFirst), uintptr(sourceLast), uintptr(destinationParent.UnsafePointer()), uintptr(destinationColumn))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sourceParent)
	runtime.KeepAlive(destinationParent)
	return __r != 0
}

func (this *QAbstractProxyModel) EndMoveColumns() {
	purego.SyscallN(addr_QAbstractProxyModel_endMoveColumns, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractProxyModel) BeginResetModel() {
	purego.SyscallN(addr_QAbstractProxyModel_beginResetModel, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractProxyModel) EndResetModel() {
	purego.SyscallN(addr_QAbstractProxyModel_endResetModel, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractProxyModel) ChangePersistentIndex(from *QModelIndex, to *QModelIndex) {
	purego.SyscallN(addr_QAbstractProxyModel_changePersistentIndex, this.h, uintptr(from.UnsafePointer()), uintptr(to.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(from)
	runtime.KeepAlive(to)
}

func (this *QAbstractProxyModel) ChangePersistentIndexList(from []*QModelIndex, to []*QModelIndex) {
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
	purego.SyscallN(addr_QAbstractProxyModel_changePersistentIndexList, this.h, from_p, uintptr(len(from)), to_p, uintptr(len(to)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(from)
	runtime.KeepAlive(from_arr)
	runtime.KeepAlive(to)
	runtime.KeepAlive(to_arr)
}

func (this *QAbstractProxyModel) PersistentIndexList() []*QModelIndex {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAbstractProxyModel_persistentIndexList, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QModelIndex, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQModelIndex(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QAbstractProxyModel_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QAbstractProxyModel_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QAbstractProxyModel_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QAbstractProxyModel_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QAbstractProxyModel) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QAbstractProxyModel_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QAbstractProxyModel) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QAbstractProxyModel_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAbstractProxyModel) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QAbstractProxyModel_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QAbstractProxyModel) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractProxyModel_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QAbstractProxyModel_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAbstractProxyModel_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QAbstractProxyModel_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAbstractProxyModel_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QAbstractProxyModel_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QAbstractProxyModel_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QAbstractProxyModel_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QAbstractProxyModel_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QAbstractProxyModel) SetSourceModelDefault(sourceModel *QAbstractItemModel) {
	purego.SyscallN(addr_QAbstractProxyModel_setSourceModel_default, this.h, uintptr(sourceModel.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sourceModel)
}

func (this *QAbstractProxyModel) MapSelectionToSourceDefault(selection *QItemSelection) *QItemSelection {
	__r, _, _ := purego.SyscallN(addr_QAbstractProxyModel_mapSelectionToSource_default, this.h, uintptr(selection.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(selection)
	return newownedQItemSelection(__r)
}

func (this *QAbstractProxyModel) MapSelectionFromSourceDefault(selection *QItemSelection) *QItemSelection {
	__r, _, _ := purego.SyscallN(addr_QAbstractProxyModel_mapSelectionFromSource_default, this.h, uintptr(selection.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(selection)
	return newownedQItemSelection(__r)
}

func (this *QAbstractProxyModel) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QAbstractProxyModel_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QAbstractProxyModel) SubmitDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractProxyModel_submit_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QAbstractProxyModel) RevertDefault() {
	purego.SyscallN(addr_QAbstractProxyModel_revert_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractProxyModel) DataDefault(proxyIndex *QModelIndex, role int) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QAbstractProxyModel_data_default, this.h, uintptr(proxyIndex.UnsafePointer()), uintptr(role))
	runtime.KeepAlive(this)
	runtime.KeepAlive(proxyIndex)
	return newownedQVariant(__r)
}

func (this *QAbstractProxyModel) HeaderDataDefault(section int, orientation Orientation, role int) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QAbstractProxyModel_headerData_default, this.h, uintptr(section), uintptr(orientation), uintptr(role))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QAbstractProxyModel) FlagsDefault(index *QModelIndex) ItemFlag {
	__r0, _, _ := purego.SyscallN(addr_QAbstractProxyModel_flags_default, this.h, uintptr(index.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return ItemFlag(__r)
}

func (this *QAbstractProxyModel) SetDataDefault(index *QModelIndex, value *QVariant, role int) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractProxyModel_setData_default, this.h, uintptr(index.UnsafePointer()), uintptr(value.UnsafePointer()), uintptr(role))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	runtime.KeepAlive(value)
	return __r != 0
}

func (this *QAbstractProxyModel) SetHeaderDataDefault(section int, orientation Orientation, value *QVariant, role int) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractProxyModel_setHeaderData_default, this.h, uintptr(section), uintptr(orientation), uintptr(value.UnsafePointer()), uintptr(role))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(value)
	return __r != 0
}

func (this *QAbstractProxyModel) ClearItemDataDefault(index *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractProxyModel_clearItemData_default, this.h, uintptr(index.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return __r != 0
}

func (this *QAbstractProxyModel) BuddyDefault(index *QModelIndex) *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QAbstractProxyModel_buddy_default, this.h, uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return newownedQModelIndex(__r)
}

func (this *QAbstractProxyModel) CanFetchMoreDefault(parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractProxyModel_canFetchMore_default, this.h, uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QAbstractProxyModel) FetchMoreDefault(parent *QModelIndex) {
	purego.SyscallN(addr_QAbstractProxyModel_fetchMore_default, this.h, uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
}

func (this *QAbstractProxyModel) SortDefault(column int, order SortOrder) {
	purego.SyscallN(addr_QAbstractProxyModel_sort_default, this.h, uintptr(column), uintptr(order))
	runtime.KeepAlive(this)
}

func (this *QAbstractProxyModel) SpanDefault(index *QModelIndex) *QSize {
	__r, _, _ := purego.SyscallN(addr_QAbstractProxyModel_span_default, this.h, uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return newownedQSize(__r)
}

func (this *QAbstractProxyModel) HasChildrenDefault(parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractProxyModel_hasChildren_default, this.h, uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QAbstractProxyModel) SiblingDefault(row int, column int, idx *QModelIndex) *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QAbstractProxyModel_sibling_default, this.h, uintptr(row), uintptr(column), uintptr(idx.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(idx)
	return newownedQModelIndex(__r)
}

func (this *QAbstractProxyModel) MimeDataDefault(indexes []*QModelIndex) *QMimeData {
	indexes_arr := make([]uintptr, len(indexes))
	for __i, __e := range indexes {
		indexes_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var indexes_p uintptr
	if len(indexes_arr) > 0 {
		indexes_p = uintptr(unsafe.Pointer(&indexes_arr[0]))
	}
	__r, _, _ := purego.SyscallN(addr_QAbstractProxyModel_mimeData_default, this.h, indexes_p, uintptr(len(indexes)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(indexes)
	runtime.KeepAlive(indexes_arr)
	return bindQMimeData(__r)
}

func (this *QAbstractProxyModel) CanDropMimeDataDefault(data *QMimeData, action DropAction, row int, column int, parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractProxyModel_canDropMimeData_default, this.h, uintptr(data.UnsafePointer()), uintptr(action), uintptr(row), uintptr(column), uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(data)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QAbstractProxyModel) DropMimeDataDefault(data *QMimeData, action DropAction, row int, column int, parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractProxyModel_dropMimeData_default, this.h, uintptr(data.UnsafePointer()), uintptr(action), uintptr(row), uintptr(column), uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(data)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QAbstractProxyModel) MimeTypesDefault() []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAbstractProxyModel_mimeTypes_default, this.h, uintptr(unsafe.Pointer(&__n)))
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

func (this *QAbstractProxyModel) SupportedDragActionsDefault() DropAction {
	__r0, _, _ := purego.SyscallN(addr_QAbstractProxyModel_supportedDragActions_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return DropAction(__r)
}

func (this *QAbstractProxyModel) SupportedDropActionsDefault() DropAction {
	__r0, _, _ := purego.SyscallN(addr_QAbstractProxyModel_supportedDropActions_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return DropAction(__r)
}

func (this *QAbstractProxyModel) InsertRowsDefault(row int, count int, parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractProxyModel_insertRows_default, this.h, uintptr(row), uintptr(count), uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QAbstractProxyModel) InsertColumnsDefault(column int, count int, parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractProxyModel_insertColumns_default, this.h, uintptr(column), uintptr(count), uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QAbstractProxyModel) RemoveRowsDefault(row int, count int, parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractProxyModel_removeRows_default, this.h, uintptr(row), uintptr(count), uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QAbstractProxyModel) RemoveColumnsDefault(column int, count int, parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractProxyModel_removeColumns_default, this.h, uintptr(column), uintptr(count), uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QAbstractProxyModel) MoveRowsDefault(sourceParent *QModelIndex, sourceRow int, count int, destinationParent *QModelIndex, destinationChild int) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractProxyModel_moveRows_default, this.h, uintptr(sourceParent.UnsafePointer()), uintptr(sourceRow), uintptr(count), uintptr(destinationParent.UnsafePointer()), uintptr(destinationChild))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sourceParent)
	runtime.KeepAlive(destinationParent)
	return __r != 0
}

func (this *QAbstractProxyModel) MoveColumnsDefault(sourceParent *QModelIndex, sourceColumn int, count int, destinationParent *QModelIndex, destinationChild int) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractProxyModel_moveColumns_default, this.h, uintptr(sourceParent.UnsafePointer()), uintptr(sourceColumn), uintptr(count), uintptr(destinationParent.UnsafePointer()), uintptr(destinationChild))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sourceParent)
	runtime.KeepAlive(destinationParent)
	return __r != 0
}

func (this *QAbstractProxyModel) MatchDefault(start *QModelIndex, role int, value *QVariant, hits int, flags MatchFlag) []*QModelIndex {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAbstractProxyModel_match_default, this.h, uintptr(start.UnsafePointer()), uintptr(role), uintptr(value.UnsafePointer()), uintptr(hits), uintptr(flags), uintptr(unsafe.Pointer(&__n)))
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

func (this *QAbstractProxyModel) MultiDataDefault(index *QModelIndex, roleDataSpan *QModelRoleDataSpan) {
	purego.SyscallN(addr_QAbstractProxyModel_multiData_default, this.h, uintptr(index.UnsafePointer()), uintptr(roleDataSpan.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	runtime.KeepAlive(roleDataSpan)
}

func (this *QAbstractProxyModel) ResetInternalDataDefault() {
	purego.SyscallN(addr_QAbstractProxyModel_resetInternalData_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractProxyModel) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractProxyModel_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QAbstractProxyModel) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractProxyModel_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QAbstractProxyModel) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QAbstractProxyModel_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractProxyModel) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QAbstractProxyModel_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractProxyModel) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QAbstractProxyModel_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractProxyModel) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QAbstractProxyModel_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QAbstractProxyModel) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QAbstractProxyModel_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QAbstractProxyModel) OnSourceModelChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QAbstractProxyModel_connect_sourceModelChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractProxyModel) OnHeaderDataChanged(fn func(Orientation, int, int)) int64 {
	id := registerCallback(func(orientation int32, first int32, last int32) { fn(Orientation(orientation), int(first), int(last)) })
	pg_QAbstractProxyModel_connect_headerDataChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractProxyModel) OnLayoutChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QAbstractProxyModel_connect_layoutChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractProxyModel) OnLayoutChanged2(fn func([]*QPersistentModelIndex, QAbstractItemModel__LayoutChangeHint)) int64 {
	id := registerCallback(func(parents uintptr, parents_len int64, hint int32) { fn(func() []*QPersistentModelIndex { if parents == 0 || parents_len == 0 { return nil }; __sl := unsafe.Slice((*unsafe.Pointer)(unsafe.Pointer(parents)), int(parents_len)); __o := make([]*QPersistentModelIndex, int(parents_len)); for __i := range __o { __o[__i] = UnsafeNewQPersistentModelIndex(__sl[__i]) }; return __o }(), QAbstractItemModel__LayoutChangeHint(hint)) })
	pg_QAbstractProxyModel_connect_layoutChanged2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractProxyModel) OnLayoutAboutToBeChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QAbstractProxyModel_connect_layoutAboutToBeChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractProxyModel) OnLayoutAboutToBeChanged2(fn func([]*QPersistentModelIndex, QAbstractItemModel__LayoutChangeHint)) int64 {
	id := registerCallback(func(parents uintptr, parents_len int64, hint int32) { fn(func() []*QPersistentModelIndex { if parents == 0 || parents_len == 0 { return nil }; __sl := unsafe.Slice((*unsafe.Pointer)(unsafe.Pointer(parents)), int(parents_len)); __o := make([]*QPersistentModelIndex, int(parents_len)); for __i := range __o { __o[__i] = UnsafeNewQPersistentModelIndex(__sl[__i]) }; return __o }(), QAbstractItemModel__LayoutChangeHint(hint)) })
	pg_QAbstractProxyModel_connect_layoutAboutToBeChanged2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractProxyModel) OnRowsAboutToBeInserted(fn func(*QModelIndex, int, int)) int64 {
	id := registerCallback(func(parent uintptr, first int32, last int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(first), int(last)) })
	pg_QAbstractProxyModel_connect_rowsAboutToBeInserted(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractProxyModel) OnRowsInserted(fn func(*QModelIndex, int, int)) int64 {
	id := registerCallback(func(parent uintptr, first int32, last int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(first), int(last)) })
	pg_QAbstractProxyModel_connect_rowsInserted(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractProxyModel) OnRowsAboutToBeRemoved(fn func(*QModelIndex, int, int)) int64 {
	id := registerCallback(func(parent uintptr, first int32, last int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(first), int(last)) })
	pg_QAbstractProxyModel_connect_rowsAboutToBeRemoved(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractProxyModel) OnRowsRemoved(fn func(*QModelIndex, int, int)) int64 {
	id := registerCallback(func(parent uintptr, first int32, last int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(first), int(last)) })
	pg_QAbstractProxyModel_connect_rowsRemoved(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractProxyModel) OnColumnsAboutToBeInserted(fn func(*QModelIndex, int, int)) int64 {
	id := registerCallback(func(parent uintptr, first int32, last int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(first), int(last)) })
	pg_QAbstractProxyModel_connect_columnsAboutToBeInserted(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractProxyModel) OnColumnsInserted(fn func(*QModelIndex, int, int)) int64 {
	id := registerCallback(func(parent uintptr, first int32, last int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(first), int(last)) })
	pg_QAbstractProxyModel_connect_columnsInserted(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractProxyModel) OnColumnsAboutToBeRemoved(fn func(*QModelIndex, int, int)) int64 {
	id := registerCallback(func(parent uintptr, first int32, last int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(first), int(last)) })
	pg_QAbstractProxyModel_connect_columnsAboutToBeRemoved(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractProxyModel) OnColumnsRemoved(fn func(*QModelIndex, int, int)) int64 {
	id := registerCallback(func(parent uintptr, first int32, last int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(first), int(last)) })
	pg_QAbstractProxyModel_connect_columnsRemoved(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractProxyModel) OnModelAboutToBeReset(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QAbstractProxyModel_connect_modelAboutToBeReset(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractProxyModel) OnModelReset(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QAbstractProxyModel_connect_modelReset(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractProxyModel) OnRowsAboutToBeMoved(fn func(*QModelIndex, int, int, *QModelIndex, int)) int64 {
	id := registerCallback(func(sourceParent uintptr, sourceStart int32, sourceEnd int32, destinationParent uintptr, destinationRow int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(sourceParent)), int(sourceStart), int(sourceEnd), UnsafeNewQModelIndex(unsafe.Pointer(destinationParent)), int(destinationRow)) })
	pg_QAbstractProxyModel_connect_rowsAboutToBeMoved(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractProxyModel) OnRowsMoved(fn func(*QModelIndex, int, int, *QModelIndex, int)) int64 {
	id := registerCallback(func(sourceParent uintptr, sourceStart int32, sourceEnd int32, destinationParent uintptr, destinationRow int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(sourceParent)), int(sourceStart), int(sourceEnd), UnsafeNewQModelIndex(unsafe.Pointer(destinationParent)), int(destinationRow)) })
	pg_QAbstractProxyModel_connect_rowsMoved(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractProxyModel) OnColumnsAboutToBeMoved(fn func(*QModelIndex, int, int, *QModelIndex, int)) int64 {
	id := registerCallback(func(sourceParent uintptr, sourceStart int32, sourceEnd int32, destinationParent uintptr, destinationColumn int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(sourceParent)), int(sourceStart), int(sourceEnd), UnsafeNewQModelIndex(unsafe.Pointer(destinationParent)), int(destinationColumn)) })
	pg_QAbstractProxyModel_connect_columnsAboutToBeMoved(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractProxyModel) OnColumnsMoved(fn func(*QModelIndex, int, int, *QModelIndex, int)) int64 {
	id := registerCallback(func(sourceParent uintptr, sourceStart int32, sourceEnd int32, destinationParent uintptr, destinationColumn int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(sourceParent)), int(sourceStart), int(sourceEnd), UnsafeNewQModelIndex(unsafe.Pointer(destinationParent)), int(destinationColumn)) })
	pg_QAbstractProxyModel_connect_columnsMoved(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractProxyModel) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QAbstractProxyModel_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractProxyModel) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QAbstractProxyModel_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractProxyModel) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QAbstractProxyModel_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractProxyModel) Disconnect(handle int64) {
	pg_QAbstractProxyModel__disconnect(handle)
	unregisterCallback(handle)
}

func QAbstractProxyModel_virtIdx(slot int64) int32 {
	switch slot {
	case -8948777187306027381:
		return 0
	case 7371339085820260189:
		return 1
	case 1960960171956884010:
		return 2
	case -8256963024669584834:
		return 3
	case -2333321292016234891:
		return 4
	case -1717009778236344012:
		return 5
	case 1607084412591819503:
		return 6
	case 7818871925547456643:
		return 7
	case 2788628866280405837:
		return 8
	case 9067522407071207004:
		return 9
	case -3788331537442685179:
		return 10
	case -1522951607536117149:
		return 11
	case -8837375942098466299:
		return 12
	case -3297140021896693688:
		return 13
	case 1703382833545304812:
		return 14
	case 35191454575202219:
		return 15
	case 4852200306590054850:
		return 16
	case 8914190500520542409:
		return 17
	case 6112860112248786499:
		return 18
	case 8639262066473964150:
		return 19
	case 3553294593611130806:
		return 20
	case 6912441694453382161:
		return 21
	case -8395342916587040543:
		return 22
	case 3369452259144514380:
		return 23
	case 397329346813782521:
		return 24
	case 5175297091024974749:
		return 25
	case 731539533644753052:
		return 26
	case -9049994684305682212:
		return 27
	case 5464467312394730418:
		return 28
	case 3421601386088570762:
		return 29
	case -7784642176186821205:
		return 30
	case -2775729825962395603:
		return 31
	case -8973293846230792237:
		return 32
	case 5104072549182222656:
		return 33
	case 2078548737056540420:
		return 34
	case -6074819356947405415:
		return 35
	case -6321007591388309905:
		return 36
	case -4341501217981328394:
		return 37
	case 731311060944096642:
		return 38
	case -4510675770784051183:
		return 39
	case 1237236456312824991:
		return 40
	case 5217623514630858297:
		return 41
	case -5409927581812374498:
		return 42
	case -1185315367526857833:
		return 43
	case 3769868229123868474:
		return 44
	case 7379029982170503008:
		return 45
	case -568369772982215962:
		return 46
	}
	return -1
}

func QAbstractProxyModel_enable(h uintptr, idx int32) { pg_QAbstractProxyModel__set_override(h, idx) }

func NewQAbstractProxyModel() *QAbstractProxyModel {
	id := nextVirtID()
	o := newownedQAbstractProxyModel(pg_QAbstractProxyModel__Go_new(id, 0))
	o.UnsafeSetGoID(id)
	registerInstance(id, o.h, QAbstractProxyModel_virtIdx, QAbstractProxyModel_enable)
	return o
}

func NewQAbstractProxyModel2(parent *QObject) *QAbstractProxyModel {
	id := nextVirtID()
	o := newownedQAbstractProxyModel(pg_QAbstractProxyModel__Go_new(id, uintptr(parent.UnsafePointer())))
	o.UnsafeSetGoID(id)
	registerInstance(id, o.h, QAbstractProxyModel_virtIdx, QAbstractProxyModel_enable)
	return o
}

func (this *QAbstractProxyModel) OnIndex(fn func(int, int, *QModelIndex) *QModelIndex) {
	registerVirt(this.UnsafeGoID(), -8948777187306027381, func(row int32, column int32, parent uintptr) uintptr { return uintptr(fn(int(row), int(column), UnsafeNewQModelIndex(unsafe.Pointer(parent))).UnsafePointer()) }) // slot "index"
}

func (this *QAbstractProxyModel) OnParent(fn func(*QModelIndex) *QModelIndex) {
	registerVirt(this.UnsafeGoID(), 7371339085820260189, func(child uintptr) uintptr { return uintptr(fn(UnsafeNewQModelIndex(unsafe.Pointer(child))).UnsafePointer()) }) // slot "parent"
}

func (this *QAbstractProxyModel) OnRowCount(fn func(*QModelIndex) int) {
	registerVirt(this.UnsafeGoID(), 1960960171956884010, func(parent uintptr) int32 { return int32(fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)))) }) // slot "rowCount"
}

func (this *QAbstractProxyModel) OnColumnCount(fn func(*QModelIndex) int) {
	registerVirt(this.UnsafeGoID(), -8256963024669584834, func(parent uintptr) int32 { return int32(fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)))) }) // slot "columnCount"
}

func (this *QAbstractProxyModel) OnMapToSource(fn func(*QModelIndex) *QModelIndex) {
	registerVirt(this.UnsafeGoID(), -2333321292016234891, func(proxyIndex uintptr) uintptr { return uintptr(fn(UnsafeNewQModelIndex(unsafe.Pointer(proxyIndex))).UnsafePointer()) }) // slot "mapToSource"
}

func (this *QAbstractProxyModel) OnMapFromSource(fn func(*QModelIndex) *QModelIndex) {
	registerVirt(this.UnsafeGoID(), -1717009778236344012, func(sourceIndex uintptr) uintptr { return uintptr(fn(UnsafeNewQModelIndex(unsafe.Pointer(sourceIndex))).UnsafePointer()) }) // slot "mapFromSource"
}

func (this *QAbstractProxyModel) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QAbstractProxyModel) OnSetSourceModel(fn func(*QAbstractItemModel)) {
	registerVirt(this.UnsafeGoID(), 7818871925547456643, func(sourceModel uintptr) { fn(UnsafeNewQAbstractItemModel(unsafe.Pointer(sourceModel))) }) // slot "setSourceModel"
}

func (this *QAbstractProxyModel) OnMapSelectionToSource(fn func(*QItemSelection) *QItemSelection) {
	registerVirt(this.UnsafeGoID(), 2788628866280405837, func(selection uintptr) uintptr { return uintptr(fn(UnsafeNewQItemSelection(unsafe.Pointer(selection))).UnsafePointer()) }) // slot "mapSelectionToSource"
}

func (this *QAbstractProxyModel) OnMapSelectionFromSource(fn func(*QItemSelection) *QItemSelection) {
	registerVirt(this.UnsafeGoID(), 9067522407071207004, func(selection uintptr) uintptr { return uintptr(fn(UnsafeNewQItemSelection(unsafe.Pointer(selection))).UnsafePointer()) }) // slot "mapSelectionFromSource"
}

func (this *QAbstractProxyModel) OnSubmit(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -3788331537442685179, func() int32 { return b2i(fn()) }) // slot "submit"
}

func (this *QAbstractProxyModel) OnRevert(fn func()) {
	registerVirt(this.UnsafeGoID(), -1522951607536117149, func() { fn() }) // slot "revert"
}

func (this *QAbstractProxyModel) OnData(fn func(*QModelIndex, int) *QVariant) {
	registerVirt(this.UnsafeGoID(), -8837375942098466299, func(proxyIndex uintptr, role int32) uintptr { return uintptr(fn(UnsafeNewQModelIndex(unsafe.Pointer(proxyIndex)), int(role)).UnsafePointer()) }) // slot "data"
}

func (this *QAbstractProxyModel) OnHeaderData(fn func(int, Orientation, int) *QVariant) {
	registerVirt(this.UnsafeGoID(), -3297140021896693688, func(section int32, orientation int32, role int32) uintptr { return uintptr(fn(int(section), Orientation(orientation), int(role)).UnsafePointer()) }) // slot "headerData"
}

func (this *QAbstractProxyModel) OnFlags(fn func(*QModelIndex) ItemFlag) {
	registerVirt(this.UnsafeGoID(), 1703382833545304812, func(index uintptr) int32 { return int32(fn(UnsafeNewQModelIndex(unsafe.Pointer(index)))) }) // slot "flags"
}

func (this *QAbstractProxyModel) OnSetData(fn func(*QModelIndex, *QVariant, int) bool) {
	registerVirt(this.UnsafeGoID(), 35191454575202219, func(index uintptr, value uintptr, role int32) int32 { return b2i(fn(UnsafeNewQModelIndex(unsafe.Pointer(index)), UnsafeNewQVariant(unsafe.Pointer(value)), int(role))) }) // slot "setData"
}

func (this *QAbstractProxyModel) OnSetHeaderData(fn func(int, Orientation, *QVariant, int) bool) {
	registerVirt(this.UnsafeGoID(), 4852200306590054850, func(section int32, orientation int32, value uintptr, role int32) int32 { return b2i(fn(int(section), Orientation(orientation), UnsafeNewQVariant(unsafe.Pointer(value)), int(role))) }) // slot "setHeaderData"
}

func (this *QAbstractProxyModel) OnClearItemData(fn func(*QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), 8914190500520542409, func(index uintptr) int32 { return b2i(fn(UnsafeNewQModelIndex(unsafe.Pointer(index)))) }) // slot "clearItemData"
}

func (this *QAbstractProxyModel) OnBuddy(fn func(*QModelIndex) *QModelIndex) {
	registerVirt(this.UnsafeGoID(), 6112860112248786499, func(index uintptr) uintptr { return uintptr(fn(UnsafeNewQModelIndex(unsafe.Pointer(index))).UnsafePointer()) }) // slot "buddy"
}

func (this *QAbstractProxyModel) OnCanFetchMore(fn func(*QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), 8639262066473964150, func(parent uintptr) int32 { return b2i(fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)))) }) // slot "canFetchMore"
}

func (this *QAbstractProxyModel) OnFetchMore(fn func(*QModelIndex)) {
	registerVirt(this.UnsafeGoID(), 3553294593611130806, func(parent uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent))) }) // slot "fetchMore"
}

func (this *QAbstractProxyModel) OnSort(fn func(int, SortOrder)) {
	registerVirt(this.UnsafeGoID(), 6912441694453382161, func(column int32, order int32) { fn(int(column), SortOrder(order)) }) // slot "sort"
}

func (this *QAbstractProxyModel) OnSpan(fn func(*QModelIndex) *QSize) {
	registerVirt(this.UnsafeGoID(), -8395342916587040543, func(index uintptr) uintptr { return uintptr(fn(UnsafeNewQModelIndex(unsafe.Pointer(index))).UnsafePointer()) }) // slot "span"
}

func (this *QAbstractProxyModel) OnHasChildren(fn func(*QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), 3369452259144514380, func(parent uintptr) int32 { return b2i(fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)))) }) // slot "hasChildren"
}

func (this *QAbstractProxyModel) OnSibling(fn func(int, int, *QModelIndex) *QModelIndex) {
	registerVirt(this.UnsafeGoID(), 397329346813782521, func(row int32, column int32, idx uintptr) uintptr { return uintptr(fn(int(row), int(column), UnsafeNewQModelIndex(unsafe.Pointer(idx))).UnsafePointer()) }) // slot "sibling"
}

func (this *QAbstractProxyModel) OnMimeData(fn func([]*QModelIndex) *QMimeData) {
	registerVirt(this.UnsafeGoID(), 5175297091024974749, func(indexes uintptr, indexes_len int64) uintptr { return uintptr(fn(func() []*QModelIndex { if indexes == 0 || indexes_len == 0 { return nil }; __sl := unsafe.Slice((*unsafe.Pointer)(unsafe.Pointer(indexes)), int(indexes_len)); __o := make([]*QModelIndex, int(indexes_len)); for __i := range __o { __o[__i] = UnsafeNewQModelIndex(__sl[__i]) }; return __o }()).UnsafePointer()) }) // slot "mimeData"
}

func (this *QAbstractProxyModel) OnCanDropMimeData(fn func(*QMimeData, DropAction, int, int, *QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), 731539533644753052, func(data uintptr, action int32, row int32, column int32, parent uintptr) int32 { return b2i(fn(UnsafeNewQMimeData(unsafe.Pointer(data)), DropAction(action), int(row), int(column), UnsafeNewQModelIndex(unsafe.Pointer(parent)))) }) // slot "canDropMimeData"
}

func (this *QAbstractProxyModel) OnDropMimeData(fn func(*QMimeData, DropAction, int, int, *QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), -9049994684305682212, func(data uintptr, action int32, row int32, column int32, parent uintptr) int32 { return b2i(fn(UnsafeNewQMimeData(unsafe.Pointer(data)), DropAction(action), int(row), int(column), UnsafeNewQModelIndex(unsafe.Pointer(parent)))) }) // slot "dropMimeData"
}

func (this *QAbstractProxyModel) OnMimeTypes(fn func() []string) {
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

func (this *QAbstractProxyModel) OnSupportedDragActions(fn func() DropAction) {
	registerVirt(this.UnsafeGoID(), 3421601386088570762, func() int32 { return int32(fn()) }) // slot "supportedDragActions"
}

func (this *QAbstractProxyModel) OnSupportedDropActions(fn func() DropAction) {
	registerVirt(this.UnsafeGoID(), -7784642176186821205, func() int32 { return int32(fn()) }) // slot "supportedDropActions"
}

func (this *QAbstractProxyModel) OnInsertRows(fn func(int, int, *QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), -2775729825962395603, func(row int32, count int32, parent uintptr) int32 { return b2i(fn(int(row), int(count), UnsafeNewQModelIndex(unsafe.Pointer(parent)))) }) // slot "insertRows"
}

func (this *QAbstractProxyModel) OnInsertColumns(fn func(int, int, *QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), -8973293846230792237, func(column int32, count int32, parent uintptr) int32 { return b2i(fn(int(column), int(count), UnsafeNewQModelIndex(unsafe.Pointer(parent)))) }) // slot "insertColumns"
}

func (this *QAbstractProxyModel) OnRemoveRows(fn func(int, int, *QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), 5104072549182222656, func(row int32, count int32, parent uintptr) int32 { return b2i(fn(int(row), int(count), UnsafeNewQModelIndex(unsafe.Pointer(parent)))) }) // slot "removeRows"
}

func (this *QAbstractProxyModel) OnRemoveColumns(fn func(int, int, *QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), 2078548737056540420, func(column int32, count int32, parent uintptr) int32 { return b2i(fn(int(column), int(count), UnsafeNewQModelIndex(unsafe.Pointer(parent)))) }) // slot "removeColumns"
}

func (this *QAbstractProxyModel) OnMoveRows(fn func(*QModelIndex, int, int, *QModelIndex, int) bool) {
	registerVirt(this.UnsafeGoID(), -6074819356947405415, func(sourceParent uintptr, sourceRow int32, count int32, destinationParent uintptr, destinationChild int32) int32 { return b2i(fn(UnsafeNewQModelIndex(unsafe.Pointer(sourceParent)), int(sourceRow), int(count), UnsafeNewQModelIndex(unsafe.Pointer(destinationParent)), int(destinationChild))) }) // slot "moveRows"
}

func (this *QAbstractProxyModel) OnMoveColumns(fn func(*QModelIndex, int, int, *QModelIndex, int) bool) {
	registerVirt(this.UnsafeGoID(), -6321007591388309905, func(sourceParent uintptr, sourceColumn int32, count int32, destinationParent uintptr, destinationChild int32) int32 { return b2i(fn(UnsafeNewQModelIndex(unsafe.Pointer(sourceParent)), int(sourceColumn), int(count), UnsafeNewQModelIndex(unsafe.Pointer(destinationParent)), int(destinationChild))) }) // slot "moveColumns"
}

func (this *QAbstractProxyModel) OnMatch(fn func(*QModelIndex, int, *QVariant, int, MatchFlag) []*QModelIndex) {
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

func (this *QAbstractProxyModel) OnMultiData(fn func(*QModelIndex, *QModelRoleDataSpan)) {
	registerVirt(this.UnsafeGoID(), 731311060944096642, func(index uintptr, roleDataSpan uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(index)), UnsafeNewQModelRoleDataSpan(unsafe.Pointer(roleDataSpan))) }) // slot "multiData"
}

func (this *QAbstractProxyModel) OnResetInternalData(fn func()) {
	registerVirt(this.UnsafeGoID(), -4510675770784051183, func() { fn() }) // slot "resetInternalData"
}

func (this *QAbstractProxyModel) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QAbstractProxyModel) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QAbstractProxyModel) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QAbstractProxyModel) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QAbstractProxyModel) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QAbstractProxyModel) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QAbstractProxyModel) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QAbstractProxyModel) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QAbstractProxyModel_delete(this.h)
	}
	this.UnsafeRelease()
}
