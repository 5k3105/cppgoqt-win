package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QTransposeProxyModel_chain uintptr
	addr_QTransposeProxyModel_new uintptr
	addr_QTransposeProxyModel_new2 uintptr
	addr_QTransposeProxyModel_tr uintptr
	addr_QTransposeProxyModel_createIndex uintptr
	addr_QTransposeProxyModel_createIndex2 uintptr
	addr_QTransposeProxyModel_encodeData uintptr
	addr_QTransposeProxyModel_decodeData uintptr
	addr_QTransposeProxyModel_beginInsertRows uintptr
	addr_QTransposeProxyModel_endInsertRows uintptr
	addr_QTransposeProxyModel_beginRemoveRows uintptr
	addr_QTransposeProxyModel_endRemoveRows uintptr
	addr_QTransposeProxyModel_beginMoveRows uintptr
	addr_QTransposeProxyModel_endMoveRows uintptr
	addr_QTransposeProxyModel_beginInsertColumns uintptr
	addr_QTransposeProxyModel_endInsertColumns uintptr
	addr_QTransposeProxyModel_beginRemoveColumns uintptr
	addr_QTransposeProxyModel_endRemoveColumns uintptr
	addr_QTransposeProxyModel_beginMoveColumns uintptr
	addr_QTransposeProxyModel_endMoveColumns uintptr
	addr_QTransposeProxyModel_beginResetModel uintptr
	addr_QTransposeProxyModel_endResetModel uintptr
	addr_QTransposeProxyModel_changePersistentIndex uintptr
	addr_QTransposeProxyModel_changePersistentIndexList uintptr
	addr_QTransposeProxyModel_persistentIndexList uintptr
	addr_QTransposeProxyModel_connect uintptr
	addr_QTransposeProxyModel_connect2 uintptr
	addr_QTransposeProxyModel_sender uintptr
	addr_QTransposeProxyModel_senderSignalIndex uintptr
	addr_QTransposeProxyModel_receivers uintptr
	addr_QTransposeProxyModel_isSignalConnected uintptr
	addr_QTransposeProxyModel_tr2 uintptr
	addr_QTransposeProxyModel_tr3 uintptr
	addr_QTransposeProxyModel_connect3 uintptr
	addr_QTransposeProxyModel_connect4 uintptr
	addr_QTransposeProxyModel_metaObject_default uintptr
	addr_QTransposeProxyModel_setSourceModel_default uintptr
	addr_QTransposeProxyModel_rowCount_default uintptr
	addr_QTransposeProxyModel_columnCount_default uintptr
	addr_QTransposeProxyModel_headerData_default uintptr
	addr_QTransposeProxyModel_setHeaderData_default uintptr
	addr_QTransposeProxyModel_span_default uintptr
	addr_QTransposeProxyModel_mapFromSource_default uintptr
	addr_QTransposeProxyModel_mapToSource_default uintptr
	addr_QTransposeProxyModel_parent_default uintptr
	addr_QTransposeProxyModel_index_default uintptr
	addr_QTransposeProxyModel_insertRows_default uintptr
	addr_QTransposeProxyModel_removeRows_default uintptr
	addr_QTransposeProxyModel_moveRows_default uintptr
	addr_QTransposeProxyModel_insertColumns_default uintptr
	addr_QTransposeProxyModel_removeColumns_default uintptr
	addr_QTransposeProxyModel_moveColumns_default uintptr
	addr_QTransposeProxyModel_sort_default uintptr
	addr_QTransposeProxyModel_mapSelectionToSource_default uintptr
	addr_QTransposeProxyModel_mapSelectionFromSource_default uintptr
	addr_QTransposeProxyModel_submit_default uintptr
	addr_QTransposeProxyModel_revert_default uintptr
	addr_QTransposeProxyModel_data_default uintptr
	addr_QTransposeProxyModel_flags_default uintptr
	addr_QTransposeProxyModel_setData_default uintptr
	addr_QTransposeProxyModel_clearItemData_default uintptr
	addr_QTransposeProxyModel_buddy_default uintptr
	addr_QTransposeProxyModel_canFetchMore_default uintptr
	addr_QTransposeProxyModel_fetchMore_default uintptr
	addr_QTransposeProxyModel_hasChildren_default uintptr
	addr_QTransposeProxyModel_sibling_default uintptr
	addr_QTransposeProxyModel_mimeData_default uintptr
	addr_QTransposeProxyModel_canDropMimeData_default uintptr
	addr_QTransposeProxyModel_dropMimeData_default uintptr
	addr_QTransposeProxyModel_mimeTypes_default uintptr
	addr_QTransposeProxyModel_supportedDragActions_default uintptr
	addr_QTransposeProxyModel_supportedDropActions_default uintptr
	addr_QTransposeProxyModel_match_default uintptr
	addr_QTransposeProxyModel_multiData_default uintptr
	addr_QTransposeProxyModel_resetInternalData_default uintptr
	addr_QTransposeProxyModel_event_default uintptr
	addr_QTransposeProxyModel_eventFilter_default uintptr
	addr_QTransposeProxyModel_timerEvent_default uintptr
	addr_QTransposeProxyModel_childEvent_default uintptr
	addr_QTransposeProxyModel_customEvent_default uintptr
	addr_QTransposeProxyModel_connectNotify_default uintptr
	addr_QTransposeProxyModel_disconnectNotify_default uintptr
	pg_QTransposeProxyModel_connect_sourceModelChanged func(uintptr, int64)
	pg_QTransposeProxyModel_connect_headerDataChanged func(uintptr, int64)
	pg_QTransposeProxyModel_connect_layoutChanged func(uintptr, int64)
	pg_QTransposeProxyModel_connect_layoutChanged2 func(uintptr, int64)
	pg_QTransposeProxyModel_connect_layoutAboutToBeChanged func(uintptr, int64)
	pg_QTransposeProxyModel_connect_layoutAboutToBeChanged2 func(uintptr, int64)
	pg_QTransposeProxyModel_connect_rowsAboutToBeInserted func(uintptr, int64)
	pg_QTransposeProxyModel_connect_rowsInserted func(uintptr, int64)
	pg_QTransposeProxyModel_connect_rowsAboutToBeRemoved func(uintptr, int64)
	pg_QTransposeProxyModel_connect_rowsRemoved func(uintptr, int64)
	pg_QTransposeProxyModel_connect_columnsAboutToBeInserted func(uintptr, int64)
	pg_QTransposeProxyModel_connect_columnsInserted func(uintptr, int64)
	pg_QTransposeProxyModel_connect_columnsAboutToBeRemoved func(uintptr, int64)
	pg_QTransposeProxyModel_connect_columnsRemoved func(uintptr, int64)
	pg_QTransposeProxyModel_connect_modelAboutToBeReset func(uintptr, int64)
	pg_QTransposeProxyModel_connect_modelReset func(uintptr, int64)
	pg_QTransposeProxyModel_connect_rowsAboutToBeMoved func(uintptr, int64)
	pg_QTransposeProxyModel_connect_rowsMoved func(uintptr, int64)
	pg_QTransposeProxyModel_connect_columnsAboutToBeMoved func(uintptr, int64)
	pg_QTransposeProxyModel_connect_columnsMoved func(uintptr, int64)
	pg_QTransposeProxyModel_connect_destroyed func(uintptr, int64)
	pg_QTransposeProxyModel_connect_destroyed2 func(uintptr, int64)
	pg_QTransposeProxyModel_connect_objectNameChanged func(uintptr, int64)
	pg_QTransposeProxyModel__disconnect func(int64)
	pg_QTransposeProxyModel__set_override func(uintptr, int32)
	pg_QTransposeProxyModel_delete func(uintptr)
)

func init() {
	symaddr(&addr_QTransposeProxyModel_chain, "QTransposeProxyModel_chain")
	symaddr(&addr_QTransposeProxyModel_new, "QTransposeProxyModel_new")
	symaddr(&addr_QTransposeProxyModel_new2, "QTransposeProxyModel_new2")
	symaddr(&addr_QTransposeProxyModel_tr, "QTransposeProxyModel_tr")
	symaddr(&addr_QTransposeProxyModel_createIndex, "QTransposeProxyModel_createIndex")
	symaddr(&addr_QTransposeProxyModel_createIndex2, "QTransposeProxyModel_createIndex2")
	symaddr(&addr_QTransposeProxyModel_encodeData, "QTransposeProxyModel_encodeData")
	symaddr(&addr_QTransposeProxyModel_decodeData, "QTransposeProxyModel_decodeData")
	symaddr(&addr_QTransposeProxyModel_beginInsertRows, "QTransposeProxyModel_beginInsertRows")
	symaddr(&addr_QTransposeProxyModel_endInsertRows, "QTransposeProxyModel_endInsertRows")
	symaddr(&addr_QTransposeProxyModel_beginRemoveRows, "QTransposeProxyModel_beginRemoveRows")
	symaddr(&addr_QTransposeProxyModel_endRemoveRows, "QTransposeProxyModel_endRemoveRows")
	symaddr(&addr_QTransposeProxyModel_beginMoveRows, "QTransposeProxyModel_beginMoveRows")
	symaddr(&addr_QTransposeProxyModel_endMoveRows, "QTransposeProxyModel_endMoveRows")
	symaddr(&addr_QTransposeProxyModel_beginInsertColumns, "QTransposeProxyModel_beginInsertColumns")
	symaddr(&addr_QTransposeProxyModel_endInsertColumns, "QTransposeProxyModel_endInsertColumns")
	symaddr(&addr_QTransposeProxyModel_beginRemoveColumns, "QTransposeProxyModel_beginRemoveColumns")
	symaddr(&addr_QTransposeProxyModel_endRemoveColumns, "QTransposeProxyModel_endRemoveColumns")
	symaddr(&addr_QTransposeProxyModel_beginMoveColumns, "QTransposeProxyModel_beginMoveColumns")
	symaddr(&addr_QTransposeProxyModel_endMoveColumns, "QTransposeProxyModel_endMoveColumns")
	symaddr(&addr_QTransposeProxyModel_beginResetModel, "QTransposeProxyModel_beginResetModel")
	symaddr(&addr_QTransposeProxyModel_endResetModel, "QTransposeProxyModel_endResetModel")
	symaddr(&addr_QTransposeProxyModel_changePersistentIndex, "QTransposeProxyModel_changePersistentIndex")
	symaddr(&addr_QTransposeProxyModel_changePersistentIndexList, "QTransposeProxyModel_changePersistentIndexList")
	symaddr(&addr_QTransposeProxyModel_persistentIndexList, "QTransposeProxyModel_persistentIndexList")
	symaddr(&addr_QTransposeProxyModel_connect, "QTransposeProxyModel_connect")
	symaddr(&addr_QTransposeProxyModel_connect2, "QTransposeProxyModel_connect2")
	symaddr(&addr_QTransposeProxyModel_sender, "QTransposeProxyModel_sender")
	symaddr(&addr_QTransposeProxyModel_senderSignalIndex, "QTransposeProxyModel_senderSignalIndex")
	symaddr(&addr_QTransposeProxyModel_receivers, "QTransposeProxyModel_receivers")
	symaddr(&addr_QTransposeProxyModel_isSignalConnected, "QTransposeProxyModel_isSignalConnected")
	symaddr(&addr_QTransposeProxyModel_tr2, "QTransposeProxyModel_tr2")
	symaddr(&addr_QTransposeProxyModel_tr3, "QTransposeProxyModel_tr3")
	symaddr(&addr_QTransposeProxyModel_connect3, "QTransposeProxyModel_connect3")
	symaddr(&addr_QTransposeProxyModel_connect4, "QTransposeProxyModel_connect4")
	symaddr(&addr_QTransposeProxyModel_metaObject_default, "QTransposeProxyModel_metaObject_default")
	symaddr(&addr_QTransposeProxyModel_setSourceModel_default, "QTransposeProxyModel_setSourceModel_default")
	symaddr(&addr_QTransposeProxyModel_rowCount_default, "QTransposeProxyModel_rowCount_default")
	symaddr(&addr_QTransposeProxyModel_columnCount_default, "QTransposeProxyModel_columnCount_default")
	symaddr(&addr_QTransposeProxyModel_headerData_default, "QTransposeProxyModel_headerData_default")
	symaddr(&addr_QTransposeProxyModel_setHeaderData_default, "QTransposeProxyModel_setHeaderData_default")
	symaddr(&addr_QTransposeProxyModel_span_default, "QTransposeProxyModel_span_default")
	symaddr(&addr_QTransposeProxyModel_mapFromSource_default, "QTransposeProxyModel_mapFromSource_default")
	symaddr(&addr_QTransposeProxyModel_mapToSource_default, "QTransposeProxyModel_mapToSource_default")
	symaddr(&addr_QTransposeProxyModel_parent_default, "QTransposeProxyModel_parent_default")
	symaddr(&addr_QTransposeProxyModel_index_default, "QTransposeProxyModel_index_default")
	symaddr(&addr_QTransposeProxyModel_insertRows_default, "QTransposeProxyModel_insertRows_default")
	symaddr(&addr_QTransposeProxyModel_removeRows_default, "QTransposeProxyModel_removeRows_default")
	symaddr(&addr_QTransposeProxyModel_moveRows_default, "QTransposeProxyModel_moveRows_default")
	symaddr(&addr_QTransposeProxyModel_insertColumns_default, "QTransposeProxyModel_insertColumns_default")
	symaddr(&addr_QTransposeProxyModel_removeColumns_default, "QTransposeProxyModel_removeColumns_default")
	symaddr(&addr_QTransposeProxyModel_moveColumns_default, "QTransposeProxyModel_moveColumns_default")
	symaddr(&addr_QTransposeProxyModel_sort_default, "QTransposeProxyModel_sort_default")
	symaddr(&addr_QTransposeProxyModel_mapSelectionToSource_default, "QTransposeProxyModel_mapSelectionToSource_default")
	symaddr(&addr_QTransposeProxyModel_mapSelectionFromSource_default, "QTransposeProxyModel_mapSelectionFromSource_default")
	symaddr(&addr_QTransposeProxyModel_submit_default, "QTransposeProxyModel_submit_default")
	symaddr(&addr_QTransposeProxyModel_revert_default, "QTransposeProxyModel_revert_default")
	symaddr(&addr_QTransposeProxyModel_data_default, "QTransposeProxyModel_data_default")
	symaddr(&addr_QTransposeProxyModel_flags_default, "QTransposeProxyModel_flags_default")
	symaddr(&addr_QTransposeProxyModel_setData_default, "QTransposeProxyModel_setData_default")
	symaddr(&addr_QTransposeProxyModel_clearItemData_default, "QTransposeProxyModel_clearItemData_default")
	symaddr(&addr_QTransposeProxyModel_buddy_default, "QTransposeProxyModel_buddy_default")
	symaddr(&addr_QTransposeProxyModel_canFetchMore_default, "QTransposeProxyModel_canFetchMore_default")
	symaddr(&addr_QTransposeProxyModel_fetchMore_default, "QTransposeProxyModel_fetchMore_default")
	symaddr(&addr_QTransposeProxyModel_hasChildren_default, "QTransposeProxyModel_hasChildren_default")
	symaddr(&addr_QTransposeProxyModel_sibling_default, "QTransposeProxyModel_sibling_default")
	symaddr(&addr_QTransposeProxyModel_mimeData_default, "QTransposeProxyModel_mimeData_default")
	symaddr(&addr_QTransposeProxyModel_canDropMimeData_default, "QTransposeProxyModel_canDropMimeData_default")
	symaddr(&addr_QTransposeProxyModel_dropMimeData_default, "QTransposeProxyModel_dropMimeData_default")
	symaddr(&addr_QTransposeProxyModel_mimeTypes_default, "QTransposeProxyModel_mimeTypes_default")
	symaddr(&addr_QTransposeProxyModel_supportedDragActions_default, "QTransposeProxyModel_supportedDragActions_default")
	symaddr(&addr_QTransposeProxyModel_supportedDropActions_default, "QTransposeProxyModel_supportedDropActions_default")
	symaddr(&addr_QTransposeProxyModel_match_default, "QTransposeProxyModel_match_default")
	symaddr(&addr_QTransposeProxyModel_multiData_default, "QTransposeProxyModel_multiData_default")
	symaddr(&addr_QTransposeProxyModel_resetInternalData_default, "QTransposeProxyModel_resetInternalData_default")
	symaddr(&addr_QTransposeProxyModel_event_default, "QTransposeProxyModel_event_default")
	symaddr(&addr_QTransposeProxyModel_eventFilter_default, "QTransposeProxyModel_eventFilter_default")
	symaddr(&addr_QTransposeProxyModel_timerEvent_default, "QTransposeProxyModel_timerEvent_default")
	symaddr(&addr_QTransposeProxyModel_childEvent_default, "QTransposeProxyModel_childEvent_default")
	symaddr(&addr_QTransposeProxyModel_customEvent_default, "QTransposeProxyModel_customEvent_default")
	symaddr(&addr_QTransposeProxyModel_connectNotify_default, "QTransposeProxyModel_connectNotify_default")
	symaddr(&addr_QTransposeProxyModel_disconnectNotify_default, "QTransposeProxyModel_disconnectNotify_default")
	reg(&pg_QTransposeProxyModel_connect_sourceModelChanged, "QTransposeProxyModel_connect_sourceModelChanged")
	reg(&pg_QTransposeProxyModel_connect_headerDataChanged, "QTransposeProxyModel_connect_headerDataChanged")
	reg(&pg_QTransposeProxyModel_connect_layoutChanged, "QTransposeProxyModel_connect_layoutChanged")
	reg(&pg_QTransposeProxyModel_connect_layoutChanged2, "QTransposeProxyModel_connect_layoutChanged2")
	reg(&pg_QTransposeProxyModel_connect_layoutAboutToBeChanged, "QTransposeProxyModel_connect_layoutAboutToBeChanged")
	reg(&pg_QTransposeProxyModel_connect_layoutAboutToBeChanged2, "QTransposeProxyModel_connect_layoutAboutToBeChanged2")
	reg(&pg_QTransposeProxyModel_connect_rowsAboutToBeInserted, "QTransposeProxyModel_connect_rowsAboutToBeInserted")
	reg(&pg_QTransposeProxyModel_connect_rowsInserted, "QTransposeProxyModel_connect_rowsInserted")
	reg(&pg_QTransposeProxyModel_connect_rowsAboutToBeRemoved, "QTransposeProxyModel_connect_rowsAboutToBeRemoved")
	reg(&pg_QTransposeProxyModel_connect_rowsRemoved, "QTransposeProxyModel_connect_rowsRemoved")
	reg(&pg_QTransposeProxyModel_connect_columnsAboutToBeInserted, "QTransposeProxyModel_connect_columnsAboutToBeInserted")
	reg(&pg_QTransposeProxyModel_connect_columnsInserted, "QTransposeProxyModel_connect_columnsInserted")
	reg(&pg_QTransposeProxyModel_connect_columnsAboutToBeRemoved, "QTransposeProxyModel_connect_columnsAboutToBeRemoved")
	reg(&pg_QTransposeProxyModel_connect_columnsRemoved, "QTransposeProxyModel_connect_columnsRemoved")
	reg(&pg_QTransposeProxyModel_connect_modelAboutToBeReset, "QTransposeProxyModel_connect_modelAboutToBeReset")
	reg(&pg_QTransposeProxyModel_connect_modelReset, "QTransposeProxyModel_connect_modelReset")
	reg(&pg_QTransposeProxyModel_connect_rowsAboutToBeMoved, "QTransposeProxyModel_connect_rowsAboutToBeMoved")
	reg(&pg_QTransposeProxyModel_connect_rowsMoved, "QTransposeProxyModel_connect_rowsMoved")
	reg(&pg_QTransposeProxyModel_connect_columnsAboutToBeMoved, "QTransposeProxyModel_connect_columnsAboutToBeMoved")
	reg(&pg_QTransposeProxyModel_connect_columnsMoved, "QTransposeProxyModel_connect_columnsMoved")
	reg(&pg_QTransposeProxyModel_connect_destroyed, "QTransposeProxyModel_connect_destroyed")
	reg(&pg_QTransposeProxyModel_connect_destroyed2, "QTransposeProxyModel_connect_destroyed2")
	reg(&pg_QTransposeProxyModel_connect_objectNameChanged, "QTransposeProxyModel_connect_objectNameChanged")
	reg(&pg_QTransposeProxyModel__disconnect, "QTransposeProxyModel__disconnect")
	reg(&pg_QTransposeProxyModel__set_override, "QTransposeProxyModel__set_override")
	reg(&pg_QTransposeProxyModel_delete, "QTransposeProxyModel_delete")
}

type QTransposeProxyModel struct {
	h uintptr
	*QAbstractProxyModel
}

func bindQTransposeProxyModel(h uintptr) *QTransposeProxyModel {
	if h == 0 {
		return nil
	}
	var __c [3]uintptr
	purego.SyscallN(addr_QTransposeProxyModel_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QTransposeProxyModel{h: h, QAbstractProxyModel: &QAbstractProxyModel{h: __c[0], QAbstractItemModel: &QAbstractItemModel{h: __c[1], QObject: &QObject{h: __c[2]}}}}
}

func newownedQTransposeProxyModel(h uintptr) *QTransposeProxyModel { return bindQTransposeProxyModel(h) }

func (this *QTransposeProxyModel) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QTransposeProxyModel) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQTransposeProxyModel(h unsafe.Pointer) *QTransposeProxyModel { return bindQTransposeProxyModel(uintptr(h)) }
func UnsafeNewQTransposeProxyModelOwned(h unsafe.Pointer) *QTransposeProxyModel { return newownedQTransposeProxyModel(uintptr(h)) }

func (this *QTransposeProxyModel) UnsafeRelease() {
	this.h = 0
	if this.QAbstractProxyModel != nil {
		this.QAbstractProxyModel.UnsafeRelease()
	}
}

func NewQTransposeProxyModel() *QTransposeProxyModel {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QTransposeProxyModel_new, uintptr(__goid))
	__o := newownedQTransposeProxyModel(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QTransposeProxyModel_virtIdx, QTransposeProxyModel_enable)
	return __o
}

func NewQTransposeProxyModel2(parent *QObject) *QTransposeProxyModel {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QTransposeProxyModel_new2, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQTransposeProxyModel(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QTransposeProxyModel_virtIdx, QTransposeProxyModel_enable)
	return __o
}

func QTransposeProxyModel_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTransposeProxyModel_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QTransposeProxyModel) CreateIndex(row int, column int) *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QTransposeProxyModel_createIndex, this.h, uintptr(row), uintptr(column))
	runtime.KeepAlive(this)
	return newownedQModelIndex(__r)
}

func (this *QTransposeProxyModel) CreateIndex2(row int, column int, id uintptr) *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QTransposeProxyModel_createIndex2, this.h, uintptr(row), uintptr(column), uintptr(id))
	runtime.KeepAlive(this)
	return newownedQModelIndex(__r)
}

func (this *QTransposeProxyModel) EncodeData(indexes []*QModelIndex, stream *QDataStream) {
	indexes_arr := make([]uintptr, len(indexes))
	for __i, __e := range indexes {
		indexes_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var indexes_p uintptr
	if len(indexes_arr) > 0 {
		indexes_p = uintptr(unsafe.Pointer(&indexes_arr[0]))
	}
	purego.SyscallN(addr_QTransposeProxyModel_encodeData, this.h, indexes_p, uintptr(len(indexes)), uintptr(stream.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(indexes)
	runtime.KeepAlive(indexes_arr)
	runtime.KeepAlive(stream)
}

func (this *QTransposeProxyModel) DecodeData(row int, column int, parent *QModelIndex, stream *QDataStream) bool {
	__r0, _, _ := purego.SyscallN(addr_QTransposeProxyModel_decodeData, this.h, uintptr(row), uintptr(column), uintptr(parent.UnsafePointer()), uintptr(stream.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	runtime.KeepAlive(stream)
	return __r != 0
}

func (this *QTransposeProxyModel) BeginInsertRows(parent *QModelIndex, first int, last int) {
	purego.SyscallN(addr_QTransposeProxyModel_beginInsertRows, this.h, uintptr(parent.UnsafePointer()), uintptr(first), uintptr(last))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
}

func (this *QTransposeProxyModel) EndInsertRows() {
	purego.SyscallN(addr_QTransposeProxyModel_endInsertRows, this.h)
	runtime.KeepAlive(this)
}

func (this *QTransposeProxyModel) BeginRemoveRows(parent *QModelIndex, first int, last int) {
	purego.SyscallN(addr_QTransposeProxyModel_beginRemoveRows, this.h, uintptr(parent.UnsafePointer()), uintptr(first), uintptr(last))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
}

func (this *QTransposeProxyModel) EndRemoveRows() {
	purego.SyscallN(addr_QTransposeProxyModel_endRemoveRows, this.h)
	runtime.KeepAlive(this)
}

func (this *QTransposeProxyModel) BeginMoveRows(sourceParent *QModelIndex, sourceFirst int, sourceLast int, destinationParent *QModelIndex, destinationRow int) bool {
	__r0, _, _ := purego.SyscallN(addr_QTransposeProxyModel_beginMoveRows, this.h, uintptr(sourceParent.UnsafePointer()), uintptr(sourceFirst), uintptr(sourceLast), uintptr(destinationParent.UnsafePointer()), uintptr(destinationRow))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sourceParent)
	runtime.KeepAlive(destinationParent)
	return __r != 0
}

func (this *QTransposeProxyModel) EndMoveRows() {
	purego.SyscallN(addr_QTransposeProxyModel_endMoveRows, this.h)
	runtime.KeepAlive(this)
}

func (this *QTransposeProxyModel) BeginInsertColumns(parent *QModelIndex, first int, last int) {
	purego.SyscallN(addr_QTransposeProxyModel_beginInsertColumns, this.h, uintptr(parent.UnsafePointer()), uintptr(first), uintptr(last))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
}

func (this *QTransposeProxyModel) EndInsertColumns() {
	purego.SyscallN(addr_QTransposeProxyModel_endInsertColumns, this.h)
	runtime.KeepAlive(this)
}

func (this *QTransposeProxyModel) BeginRemoveColumns(parent *QModelIndex, first int, last int) {
	purego.SyscallN(addr_QTransposeProxyModel_beginRemoveColumns, this.h, uintptr(parent.UnsafePointer()), uintptr(first), uintptr(last))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
}

func (this *QTransposeProxyModel) EndRemoveColumns() {
	purego.SyscallN(addr_QTransposeProxyModel_endRemoveColumns, this.h)
	runtime.KeepAlive(this)
}

func (this *QTransposeProxyModel) BeginMoveColumns(sourceParent *QModelIndex, sourceFirst int, sourceLast int, destinationParent *QModelIndex, destinationColumn int) bool {
	__r0, _, _ := purego.SyscallN(addr_QTransposeProxyModel_beginMoveColumns, this.h, uintptr(sourceParent.UnsafePointer()), uintptr(sourceFirst), uintptr(sourceLast), uintptr(destinationParent.UnsafePointer()), uintptr(destinationColumn))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sourceParent)
	runtime.KeepAlive(destinationParent)
	return __r != 0
}

func (this *QTransposeProxyModel) EndMoveColumns() {
	purego.SyscallN(addr_QTransposeProxyModel_endMoveColumns, this.h)
	runtime.KeepAlive(this)
}

func (this *QTransposeProxyModel) BeginResetModel() {
	purego.SyscallN(addr_QTransposeProxyModel_beginResetModel, this.h)
	runtime.KeepAlive(this)
}

func (this *QTransposeProxyModel) EndResetModel() {
	purego.SyscallN(addr_QTransposeProxyModel_endResetModel, this.h)
	runtime.KeepAlive(this)
}

func (this *QTransposeProxyModel) ChangePersistentIndex(from *QModelIndex, to *QModelIndex) {
	purego.SyscallN(addr_QTransposeProxyModel_changePersistentIndex, this.h, uintptr(from.UnsafePointer()), uintptr(to.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(from)
	runtime.KeepAlive(to)
}

func (this *QTransposeProxyModel) ChangePersistentIndexList(from []*QModelIndex, to []*QModelIndex) {
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
	purego.SyscallN(addr_QTransposeProxyModel_changePersistentIndexList, this.h, from_p, uintptr(len(from)), to_p, uintptr(len(to)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(from)
	runtime.KeepAlive(from_arr)
	runtime.KeepAlive(to)
	runtime.KeepAlive(to_arr)
}

func (this *QTransposeProxyModel) PersistentIndexList() []*QModelIndex {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTransposeProxyModel_persistentIndexList, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QModelIndex, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQModelIndex(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QTransposeProxyModel_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QTransposeProxyModel_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QTransposeProxyModel_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QTransposeProxyModel_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QTransposeProxyModel) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QTransposeProxyModel_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QTransposeProxyModel) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QTransposeProxyModel_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTransposeProxyModel) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QTransposeProxyModel_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QTransposeProxyModel) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QTransposeProxyModel_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QTransposeProxyModel_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTransposeProxyModel_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QTransposeProxyModel_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTransposeProxyModel_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QTransposeProxyModel_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QTransposeProxyModel_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QTransposeProxyModel_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QTransposeProxyModel_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QTransposeProxyModel) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QTransposeProxyModel_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QTransposeProxyModel) SetSourceModelDefault(newSourceModel *QAbstractItemModel) {
	purego.SyscallN(addr_QTransposeProxyModel_setSourceModel_default, this.h, uintptr(newSourceModel.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(newSourceModel)
}

func (this *QTransposeProxyModel) RowCountDefault(parent *QModelIndex) int {
	__r0, _, _ := purego.SyscallN(addr_QTransposeProxyModel_rowCount_default, this.h, uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return int(__r)
}

func (this *QTransposeProxyModel) ColumnCountDefault(parent *QModelIndex) int {
	__r0, _, _ := purego.SyscallN(addr_QTransposeProxyModel_columnCount_default, this.h, uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return int(__r)
}

func (this *QTransposeProxyModel) HeaderDataDefault(section int, orientation Orientation, role int) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QTransposeProxyModel_headerData_default, this.h, uintptr(section), uintptr(orientation), uintptr(role))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QTransposeProxyModel) SetHeaderDataDefault(section int, orientation Orientation, value *QVariant, role int) bool {
	__r0, _, _ := purego.SyscallN(addr_QTransposeProxyModel_setHeaderData_default, this.h, uintptr(section), uintptr(orientation), uintptr(value.UnsafePointer()), uintptr(role))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(value)
	return __r != 0
}

func (this *QTransposeProxyModel) SpanDefault(index *QModelIndex) *QSize {
	__r, _, _ := purego.SyscallN(addr_QTransposeProxyModel_span_default, this.h, uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return newownedQSize(__r)
}

func (this *QTransposeProxyModel) MapFromSourceDefault(sourceIndex *QModelIndex) *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QTransposeProxyModel_mapFromSource_default, this.h, uintptr(sourceIndex.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sourceIndex)
	return newownedQModelIndex(__r)
}

func (this *QTransposeProxyModel) MapToSourceDefault(proxyIndex *QModelIndex) *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QTransposeProxyModel_mapToSource_default, this.h, uintptr(proxyIndex.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(proxyIndex)
	return newownedQModelIndex(__r)
}

func (this *QTransposeProxyModel) ParentDefault(index *QModelIndex) *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QTransposeProxyModel_parent_default, this.h, uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return newownedQModelIndex(__r)
}

func (this *QTransposeProxyModel) IndexDefault(row int, column int, parent *QModelIndex) *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QTransposeProxyModel_index_default, this.h, uintptr(row), uintptr(column), uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return newownedQModelIndex(__r)
}

func (this *QTransposeProxyModel) InsertRowsDefault(row int, count int, parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QTransposeProxyModel_insertRows_default, this.h, uintptr(row), uintptr(count), uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QTransposeProxyModel) RemoveRowsDefault(row int, count int, parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QTransposeProxyModel_removeRows_default, this.h, uintptr(row), uintptr(count), uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QTransposeProxyModel) MoveRowsDefault(sourceParent *QModelIndex, sourceRow int, count int, destinationParent *QModelIndex, destinationChild int) bool {
	__r0, _, _ := purego.SyscallN(addr_QTransposeProxyModel_moveRows_default, this.h, uintptr(sourceParent.UnsafePointer()), uintptr(sourceRow), uintptr(count), uintptr(destinationParent.UnsafePointer()), uintptr(destinationChild))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sourceParent)
	runtime.KeepAlive(destinationParent)
	return __r != 0
}

func (this *QTransposeProxyModel) InsertColumnsDefault(column int, count int, parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QTransposeProxyModel_insertColumns_default, this.h, uintptr(column), uintptr(count), uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QTransposeProxyModel) RemoveColumnsDefault(column int, count int, parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QTransposeProxyModel_removeColumns_default, this.h, uintptr(column), uintptr(count), uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QTransposeProxyModel) MoveColumnsDefault(sourceParent *QModelIndex, sourceColumn int, count int, destinationParent *QModelIndex, destinationChild int) bool {
	__r0, _, _ := purego.SyscallN(addr_QTransposeProxyModel_moveColumns_default, this.h, uintptr(sourceParent.UnsafePointer()), uintptr(sourceColumn), uintptr(count), uintptr(destinationParent.UnsafePointer()), uintptr(destinationChild))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sourceParent)
	runtime.KeepAlive(destinationParent)
	return __r != 0
}

func (this *QTransposeProxyModel) SortDefault(column int, order SortOrder) {
	purego.SyscallN(addr_QTransposeProxyModel_sort_default, this.h, uintptr(column), uintptr(order))
	runtime.KeepAlive(this)
}

func (this *QTransposeProxyModel) MapSelectionToSourceDefault(selection *QItemSelection) *QItemSelection {
	__r, _, _ := purego.SyscallN(addr_QTransposeProxyModel_mapSelectionToSource_default, this.h, uintptr(selection.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(selection)
	return newownedQItemSelection(__r)
}

func (this *QTransposeProxyModel) MapSelectionFromSourceDefault(selection *QItemSelection) *QItemSelection {
	__r, _, _ := purego.SyscallN(addr_QTransposeProxyModel_mapSelectionFromSource_default, this.h, uintptr(selection.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(selection)
	return newownedQItemSelection(__r)
}

func (this *QTransposeProxyModel) SubmitDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QTransposeProxyModel_submit_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTransposeProxyModel) RevertDefault() {
	purego.SyscallN(addr_QTransposeProxyModel_revert_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QTransposeProxyModel) DataDefault(proxyIndex *QModelIndex, role int) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QTransposeProxyModel_data_default, this.h, uintptr(proxyIndex.UnsafePointer()), uintptr(role))
	runtime.KeepAlive(this)
	runtime.KeepAlive(proxyIndex)
	return newownedQVariant(__r)
}

func (this *QTransposeProxyModel) FlagsDefault(index *QModelIndex) ItemFlag {
	__r0, _, _ := purego.SyscallN(addr_QTransposeProxyModel_flags_default, this.h, uintptr(index.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return ItemFlag(__r)
}

func (this *QTransposeProxyModel) SetDataDefault(index *QModelIndex, value *QVariant, role int) bool {
	__r0, _, _ := purego.SyscallN(addr_QTransposeProxyModel_setData_default, this.h, uintptr(index.UnsafePointer()), uintptr(value.UnsafePointer()), uintptr(role))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	runtime.KeepAlive(value)
	return __r != 0
}

func (this *QTransposeProxyModel) ClearItemDataDefault(index *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QTransposeProxyModel_clearItemData_default, this.h, uintptr(index.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return __r != 0
}

func (this *QTransposeProxyModel) BuddyDefault(index *QModelIndex) *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QTransposeProxyModel_buddy_default, this.h, uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return newownedQModelIndex(__r)
}

func (this *QTransposeProxyModel) CanFetchMoreDefault(parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QTransposeProxyModel_canFetchMore_default, this.h, uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QTransposeProxyModel) FetchMoreDefault(parent *QModelIndex) {
	purego.SyscallN(addr_QTransposeProxyModel_fetchMore_default, this.h, uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
}

func (this *QTransposeProxyModel) HasChildrenDefault(parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QTransposeProxyModel_hasChildren_default, this.h, uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QTransposeProxyModel) SiblingDefault(row int, column int, idx *QModelIndex) *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QTransposeProxyModel_sibling_default, this.h, uintptr(row), uintptr(column), uintptr(idx.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(idx)
	return newownedQModelIndex(__r)
}

func (this *QTransposeProxyModel) MimeDataDefault(indexes []*QModelIndex) *QMimeData {
	indexes_arr := make([]uintptr, len(indexes))
	for __i, __e := range indexes {
		indexes_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var indexes_p uintptr
	if len(indexes_arr) > 0 {
		indexes_p = uintptr(unsafe.Pointer(&indexes_arr[0]))
	}
	__r, _, _ := purego.SyscallN(addr_QTransposeProxyModel_mimeData_default, this.h, indexes_p, uintptr(len(indexes)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(indexes)
	runtime.KeepAlive(indexes_arr)
	return bindQMimeData(__r)
}

func (this *QTransposeProxyModel) CanDropMimeDataDefault(data *QMimeData, action DropAction, row int, column int, parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QTransposeProxyModel_canDropMimeData_default, this.h, uintptr(data.UnsafePointer()), uintptr(action), uintptr(row), uintptr(column), uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(data)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QTransposeProxyModel) DropMimeDataDefault(data *QMimeData, action DropAction, row int, column int, parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QTransposeProxyModel_dropMimeData_default, this.h, uintptr(data.UnsafePointer()), uintptr(action), uintptr(row), uintptr(column), uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(data)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QTransposeProxyModel) MimeTypesDefault() []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTransposeProxyModel_mimeTypes_default, this.h, uintptr(unsafe.Pointer(&__n)))
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

func (this *QTransposeProxyModel) SupportedDragActionsDefault() DropAction {
	__r0, _, _ := purego.SyscallN(addr_QTransposeProxyModel_supportedDragActions_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return DropAction(__r)
}

func (this *QTransposeProxyModel) SupportedDropActionsDefault() DropAction {
	__r0, _, _ := purego.SyscallN(addr_QTransposeProxyModel_supportedDropActions_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return DropAction(__r)
}

func (this *QTransposeProxyModel) MatchDefault(start *QModelIndex, role int, value *QVariant, hits int, flags MatchFlag) []*QModelIndex {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTransposeProxyModel_match_default, this.h, uintptr(start.UnsafePointer()), uintptr(role), uintptr(value.UnsafePointer()), uintptr(hits), uintptr(flags), uintptr(unsafe.Pointer(&__n)))
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

func (this *QTransposeProxyModel) MultiDataDefault(index *QModelIndex, roleDataSpan *QModelRoleDataSpan) {
	purego.SyscallN(addr_QTransposeProxyModel_multiData_default, this.h, uintptr(index.UnsafePointer()), uintptr(roleDataSpan.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	runtime.KeepAlive(roleDataSpan)
}

func (this *QTransposeProxyModel) ResetInternalDataDefault() {
	purego.SyscallN(addr_QTransposeProxyModel_resetInternalData_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QTransposeProxyModel) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QTransposeProxyModel_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QTransposeProxyModel) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QTransposeProxyModel_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QTransposeProxyModel) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QTransposeProxyModel_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTransposeProxyModel) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QTransposeProxyModel_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTransposeProxyModel) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QTransposeProxyModel_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTransposeProxyModel) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QTransposeProxyModel_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QTransposeProxyModel) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QTransposeProxyModel_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QTransposeProxyModel) OnSourceModelChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QTransposeProxyModel_connect_sourceModelChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTransposeProxyModel) OnHeaderDataChanged(fn func(Orientation, int, int)) int64 {
	id := registerCallback(func(orientation int32, first int32, last int32) { fn(Orientation(orientation), int(first), int(last)) })
	pg_QTransposeProxyModel_connect_headerDataChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTransposeProxyModel) OnLayoutChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QTransposeProxyModel_connect_layoutChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTransposeProxyModel) OnLayoutChanged2(fn func([]*QPersistentModelIndex, QAbstractItemModel__LayoutChangeHint)) int64 {
	id := registerCallback(func(parents uintptr, parents_len int64, hint int32) { fn(func() []*QPersistentModelIndex { if parents == 0 || parents_len == 0 { return nil }; __sl := unsafe.Slice((*unsafe.Pointer)(unsafe.Pointer(parents)), int(parents_len)); __o := make([]*QPersistentModelIndex, int(parents_len)); for __i := range __o { __o[__i] = UnsafeNewQPersistentModelIndex(__sl[__i]) }; return __o }(), QAbstractItemModel__LayoutChangeHint(hint)) })
	pg_QTransposeProxyModel_connect_layoutChanged2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTransposeProxyModel) OnLayoutAboutToBeChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QTransposeProxyModel_connect_layoutAboutToBeChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTransposeProxyModel) OnLayoutAboutToBeChanged2(fn func([]*QPersistentModelIndex, QAbstractItemModel__LayoutChangeHint)) int64 {
	id := registerCallback(func(parents uintptr, parents_len int64, hint int32) { fn(func() []*QPersistentModelIndex { if parents == 0 || parents_len == 0 { return nil }; __sl := unsafe.Slice((*unsafe.Pointer)(unsafe.Pointer(parents)), int(parents_len)); __o := make([]*QPersistentModelIndex, int(parents_len)); for __i := range __o { __o[__i] = UnsafeNewQPersistentModelIndex(__sl[__i]) }; return __o }(), QAbstractItemModel__LayoutChangeHint(hint)) })
	pg_QTransposeProxyModel_connect_layoutAboutToBeChanged2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTransposeProxyModel) OnRowsAboutToBeInserted(fn func(*QModelIndex, int, int)) int64 {
	id := registerCallback(func(parent uintptr, first int32, last int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(first), int(last)) })
	pg_QTransposeProxyModel_connect_rowsAboutToBeInserted(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTransposeProxyModel) OnRowsInserted(fn func(*QModelIndex, int, int)) int64 {
	id := registerCallback(func(parent uintptr, first int32, last int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(first), int(last)) })
	pg_QTransposeProxyModel_connect_rowsInserted(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTransposeProxyModel) OnRowsAboutToBeRemoved(fn func(*QModelIndex, int, int)) int64 {
	id := registerCallback(func(parent uintptr, first int32, last int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(first), int(last)) })
	pg_QTransposeProxyModel_connect_rowsAboutToBeRemoved(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTransposeProxyModel) OnRowsRemoved(fn func(*QModelIndex, int, int)) int64 {
	id := registerCallback(func(parent uintptr, first int32, last int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(first), int(last)) })
	pg_QTransposeProxyModel_connect_rowsRemoved(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTransposeProxyModel) OnColumnsAboutToBeInserted(fn func(*QModelIndex, int, int)) int64 {
	id := registerCallback(func(parent uintptr, first int32, last int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(first), int(last)) })
	pg_QTransposeProxyModel_connect_columnsAboutToBeInserted(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTransposeProxyModel) OnColumnsInserted(fn func(*QModelIndex, int, int)) int64 {
	id := registerCallback(func(parent uintptr, first int32, last int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(first), int(last)) })
	pg_QTransposeProxyModel_connect_columnsInserted(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTransposeProxyModel) OnColumnsAboutToBeRemoved(fn func(*QModelIndex, int, int)) int64 {
	id := registerCallback(func(parent uintptr, first int32, last int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(first), int(last)) })
	pg_QTransposeProxyModel_connect_columnsAboutToBeRemoved(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTransposeProxyModel) OnColumnsRemoved(fn func(*QModelIndex, int, int)) int64 {
	id := registerCallback(func(parent uintptr, first int32, last int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(first), int(last)) })
	pg_QTransposeProxyModel_connect_columnsRemoved(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTransposeProxyModel) OnModelAboutToBeReset(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QTransposeProxyModel_connect_modelAboutToBeReset(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTransposeProxyModel) OnModelReset(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QTransposeProxyModel_connect_modelReset(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTransposeProxyModel) OnRowsAboutToBeMoved(fn func(*QModelIndex, int, int, *QModelIndex, int)) int64 {
	id := registerCallback(func(sourceParent uintptr, sourceStart int32, sourceEnd int32, destinationParent uintptr, destinationRow int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(sourceParent)), int(sourceStart), int(sourceEnd), UnsafeNewQModelIndex(unsafe.Pointer(destinationParent)), int(destinationRow)) })
	pg_QTransposeProxyModel_connect_rowsAboutToBeMoved(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTransposeProxyModel) OnRowsMoved(fn func(*QModelIndex, int, int, *QModelIndex, int)) int64 {
	id := registerCallback(func(sourceParent uintptr, sourceStart int32, sourceEnd int32, destinationParent uintptr, destinationRow int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(sourceParent)), int(sourceStart), int(sourceEnd), UnsafeNewQModelIndex(unsafe.Pointer(destinationParent)), int(destinationRow)) })
	pg_QTransposeProxyModel_connect_rowsMoved(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTransposeProxyModel) OnColumnsAboutToBeMoved(fn func(*QModelIndex, int, int, *QModelIndex, int)) int64 {
	id := registerCallback(func(sourceParent uintptr, sourceStart int32, sourceEnd int32, destinationParent uintptr, destinationColumn int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(sourceParent)), int(sourceStart), int(sourceEnd), UnsafeNewQModelIndex(unsafe.Pointer(destinationParent)), int(destinationColumn)) })
	pg_QTransposeProxyModel_connect_columnsAboutToBeMoved(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTransposeProxyModel) OnColumnsMoved(fn func(*QModelIndex, int, int, *QModelIndex, int)) int64 {
	id := registerCallback(func(sourceParent uintptr, sourceStart int32, sourceEnd int32, destinationParent uintptr, destinationColumn int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(sourceParent)), int(sourceStart), int(sourceEnd), UnsafeNewQModelIndex(unsafe.Pointer(destinationParent)), int(destinationColumn)) })
	pg_QTransposeProxyModel_connect_columnsMoved(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTransposeProxyModel) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QTransposeProxyModel_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTransposeProxyModel) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QTransposeProxyModel_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTransposeProxyModel) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QTransposeProxyModel_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTransposeProxyModel) Disconnect(handle int64) {
	pg_QTransposeProxyModel__disconnect(handle)
	unregisterCallback(handle)
}

func QTransposeProxyModel_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case 7818871925547456643:
		return 1
	case 1960960171956884010:
		return 2
	case -8256963024669584834:
		return 3
	case -3297140021896693688:
		return 4
	case 4852200306590054850:
		return 5
	case -8395342916587040543:
		return 6
	case -1717009778236344012:
		return 7
	case -2333321292016234891:
		return 8
	case 7371339085820260189:
		return 9
	case -8948777187306027381:
		return 10
	case -2775729825962395603:
		return 11
	case 5104072549182222656:
		return 12
	case -6074819356947405415:
		return 13
	case -8973293846230792237:
		return 14
	case 2078548737056540420:
		return 15
	case -6321007591388309905:
		return 16
	case 6912441694453382161:
		return 17
	case 2788628866280405837:
		return 18
	case 9067522407071207004:
		return 19
	case -3788331537442685179:
		return 20
	case -1522951607536117149:
		return 21
	case -8837375942098466299:
		return 22
	case 1703382833545304812:
		return 23
	case 35191454575202219:
		return 24
	case 8914190500520542409:
		return 25
	case 6112860112248786499:
		return 26
	case 8639262066473964150:
		return 27
	case 3553294593611130806:
		return 28
	case 3369452259144514380:
		return 29
	case 397329346813782521:
		return 30
	case 5175297091024974749:
		return 31
	case 731539533644753052:
		return 32
	case -9049994684305682212:
		return 33
	case 5464467312394730418:
		return 34
	case 3421601386088570762:
		return 35
	case -7784642176186821205:
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

func QTransposeProxyModel_enable(h uintptr, idx int32) { pg_QTransposeProxyModel__set_override(h, idx) }

func (this *QTransposeProxyModel) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QTransposeProxyModel) OnSetSourceModel(fn func(*QAbstractItemModel)) {
	registerVirt(this.UnsafeGoID(), 7818871925547456643, func(newSourceModel uintptr) { fn(UnsafeNewQAbstractItemModel(unsafe.Pointer(newSourceModel))) }) // slot "setSourceModel"
}

func (this *QTransposeProxyModel) OnRowCount(fn func(*QModelIndex) int) {
	registerVirt(this.UnsafeGoID(), 1960960171956884010, func(parent uintptr) int32 { return int32(fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)))) }) // slot "rowCount"
}

func (this *QTransposeProxyModel) OnColumnCount(fn func(*QModelIndex) int) {
	registerVirt(this.UnsafeGoID(), -8256963024669584834, func(parent uintptr) int32 { return int32(fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)))) }) // slot "columnCount"
}

func (this *QTransposeProxyModel) OnHeaderData(fn func(int, Orientation, int) *QVariant) {
	registerVirt(this.UnsafeGoID(), -3297140021896693688, func(section int32, orientation int32, role int32) uintptr { return uintptr(fn(int(section), Orientation(orientation), int(role)).UnsafePointer()) }) // slot "headerData"
}

func (this *QTransposeProxyModel) OnSetHeaderData(fn func(int, Orientation, *QVariant, int) bool) {
	registerVirt(this.UnsafeGoID(), 4852200306590054850, func(section int32, orientation int32, value uintptr, role int32) int32 { return b2i(fn(int(section), Orientation(orientation), UnsafeNewQVariant(unsafe.Pointer(value)), int(role))) }) // slot "setHeaderData"
}

func (this *QTransposeProxyModel) OnSpan(fn func(*QModelIndex) *QSize) {
	registerVirt(this.UnsafeGoID(), -8395342916587040543, func(index uintptr) uintptr { return uintptr(fn(UnsafeNewQModelIndex(unsafe.Pointer(index))).UnsafePointer()) }) // slot "span"
}

func (this *QTransposeProxyModel) OnMapFromSource(fn func(*QModelIndex) *QModelIndex) {
	registerVirt(this.UnsafeGoID(), -1717009778236344012, func(sourceIndex uintptr) uintptr { return uintptr(fn(UnsafeNewQModelIndex(unsafe.Pointer(sourceIndex))).UnsafePointer()) }) // slot "mapFromSource"
}

func (this *QTransposeProxyModel) OnMapToSource(fn func(*QModelIndex) *QModelIndex) {
	registerVirt(this.UnsafeGoID(), -2333321292016234891, func(proxyIndex uintptr) uintptr { return uintptr(fn(UnsafeNewQModelIndex(unsafe.Pointer(proxyIndex))).UnsafePointer()) }) // slot "mapToSource"
}

func (this *QTransposeProxyModel) OnParent(fn func(*QModelIndex) *QModelIndex) {
	registerVirt(this.UnsafeGoID(), 7371339085820260189, func(index uintptr) uintptr { return uintptr(fn(UnsafeNewQModelIndex(unsafe.Pointer(index))).UnsafePointer()) }) // slot "parent"
}

func (this *QTransposeProxyModel) OnIndex(fn func(int, int, *QModelIndex) *QModelIndex) {
	registerVirt(this.UnsafeGoID(), -8948777187306027381, func(row int32, column int32, parent uintptr) uintptr { return uintptr(fn(int(row), int(column), UnsafeNewQModelIndex(unsafe.Pointer(parent))).UnsafePointer()) }) // slot "index"
}

func (this *QTransposeProxyModel) OnInsertRows(fn func(int, int, *QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), -2775729825962395603, func(row int32, count int32, parent uintptr) int32 { return b2i(fn(int(row), int(count), UnsafeNewQModelIndex(unsafe.Pointer(parent)))) }) // slot "insertRows"
}

func (this *QTransposeProxyModel) OnRemoveRows(fn func(int, int, *QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), 5104072549182222656, func(row int32, count int32, parent uintptr) int32 { return b2i(fn(int(row), int(count), UnsafeNewQModelIndex(unsafe.Pointer(parent)))) }) // slot "removeRows"
}

func (this *QTransposeProxyModel) OnMoveRows(fn func(*QModelIndex, int, int, *QModelIndex, int) bool) {
	registerVirt(this.UnsafeGoID(), -6074819356947405415, func(sourceParent uintptr, sourceRow int32, count int32, destinationParent uintptr, destinationChild int32) int32 { return b2i(fn(UnsafeNewQModelIndex(unsafe.Pointer(sourceParent)), int(sourceRow), int(count), UnsafeNewQModelIndex(unsafe.Pointer(destinationParent)), int(destinationChild))) }) // slot "moveRows"
}

func (this *QTransposeProxyModel) OnInsertColumns(fn func(int, int, *QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), -8973293846230792237, func(column int32, count int32, parent uintptr) int32 { return b2i(fn(int(column), int(count), UnsafeNewQModelIndex(unsafe.Pointer(parent)))) }) // slot "insertColumns"
}

func (this *QTransposeProxyModel) OnRemoveColumns(fn func(int, int, *QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), 2078548737056540420, func(column int32, count int32, parent uintptr) int32 { return b2i(fn(int(column), int(count), UnsafeNewQModelIndex(unsafe.Pointer(parent)))) }) // slot "removeColumns"
}

func (this *QTransposeProxyModel) OnMoveColumns(fn func(*QModelIndex, int, int, *QModelIndex, int) bool) {
	registerVirt(this.UnsafeGoID(), -6321007591388309905, func(sourceParent uintptr, sourceColumn int32, count int32, destinationParent uintptr, destinationChild int32) int32 { return b2i(fn(UnsafeNewQModelIndex(unsafe.Pointer(sourceParent)), int(sourceColumn), int(count), UnsafeNewQModelIndex(unsafe.Pointer(destinationParent)), int(destinationChild))) }) // slot "moveColumns"
}

func (this *QTransposeProxyModel) OnSort(fn func(int, SortOrder)) {
	registerVirt(this.UnsafeGoID(), 6912441694453382161, func(column int32, order int32) { fn(int(column), SortOrder(order)) }) // slot "sort"
}

func (this *QTransposeProxyModel) OnMapSelectionToSource(fn func(*QItemSelection) *QItemSelection) {
	registerVirt(this.UnsafeGoID(), 2788628866280405837, func(selection uintptr) uintptr { return uintptr(fn(UnsafeNewQItemSelection(unsafe.Pointer(selection))).UnsafePointer()) }) // slot "mapSelectionToSource"
}

func (this *QTransposeProxyModel) OnMapSelectionFromSource(fn func(*QItemSelection) *QItemSelection) {
	registerVirt(this.UnsafeGoID(), 9067522407071207004, func(selection uintptr) uintptr { return uintptr(fn(UnsafeNewQItemSelection(unsafe.Pointer(selection))).UnsafePointer()) }) // slot "mapSelectionFromSource"
}

func (this *QTransposeProxyModel) OnSubmit(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -3788331537442685179, func() int32 { return b2i(fn()) }) // slot "submit"
}

func (this *QTransposeProxyModel) OnRevert(fn func()) {
	registerVirt(this.UnsafeGoID(), -1522951607536117149, func() { fn() }) // slot "revert"
}

func (this *QTransposeProxyModel) OnData(fn func(*QModelIndex, int) *QVariant) {
	registerVirt(this.UnsafeGoID(), -8837375942098466299, func(proxyIndex uintptr, role int32) uintptr { return uintptr(fn(UnsafeNewQModelIndex(unsafe.Pointer(proxyIndex)), int(role)).UnsafePointer()) }) // slot "data"
}

func (this *QTransposeProxyModel) OnFlags(fn func(*QModelIndex) ItemFlag) {
	registerVirt(this.UnsafeGoID(), 1703382833545304812, func(index uintptr) int32 { return int32(fn(UnsafeNewQModelIndex(unsafe.Pointer(index)))) }) // slot "flags"
}

func (this *QTransposeProxyModel) OnSetData(fn func(*QModelIndex, *QVariant, int) bool) {
	registerVirt(this.UnsafeGoID(), 35191454575202219, func(index uintptr, value uintptr, role int32) int32 { return b2i(fn(UnsafeNewQModelIndex(unsafe.Pointer(index)), UnsafeNewQVariant(unsafe.Pointer(value)), int(role))) }) // slot "setData"
}

func (this *QTransposeProxyModel) OnClearItemData(fn func(*QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), 8914190500520542409, func(index uintptr) int32 { return b2i(fn(UnsafeNewQModelIndex(unsafe.Pointer(index)))) }) // slot "clearItemData"
}

func (this *QTransposeProxyModel) OnBuddy(fn func(*QModelIndex) *QModelIndex) {
	registerVirt(this.UnsafeGoID(), 6112860112248786499, func(index uintptr) uintptr { return uintptr(fn(UnsafeNewQModelIndex(unsafe.Pointer(index))).UnsafePointer()) }) // slot "buddy"
}

func (this *QTransposeProxyModel) OnCanFetchMore(fn func(*QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), 8639262066473964150, func(parent uintptr) int32 { return b2i(fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)))) }) // slot "canFetchMore"
}

func (this *QTransposeProxyModel) OnFetchMore(fn func(*QModelIndex)) {
	registerVirt(this.UnsafeGoID(), 3553294593611130806, func(parent uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent))) }) // slot "fetchMore"
}

func (this *QTransposeProxyModel) OnHasChildren(fn func(*QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), 3369452259144514380, func(parent uintptr) int32 { return b2i(fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)))) }) // slot "hasChildren"
}

func (this *QTransposeProxyModel) OnSibling(fn func(int, int, *QModelIndex) *QModelIndex) {
	registerVirt(this.UnsafeGoID(), 397329346813782521, func(row int32, column int32, idx uintptr) uintptr { return uintptr(fn(int(row), int(column), UnsafeNewQModelIndex(unsafe.Pointer(idx))).UnsafePointer()) }) // slot "sibling"
}

func (this *QTransposeProxyModel) OnMimeData(fn func([]*QModelIndex) *QMimeData) {
	registerVirt(this.UnsafeGoID(), 5175297091024974749, func(indexes uintptr, indexes_len int64) uintptr { return uintptr(fn(func() []*QModelIndex { if indexes == 0 || indexes_len == 0 { return nil }; __sl := unsafe.Slice((*unsafe.Pointer)(unsafe.Pointer(indexes)), int(indexes_len)); __o := make([]*QModelIndex, int(indexes_len)); for __i := range __o { __o[__i] = UnsafeNewQModelIndex(__sl[__i]) }; return __o }()).UnsafePointer()) }) // slot "mimeData"
}

func (this *QTransposeProxyModel) OnCanDropMimeData(fn func(*QMimeData, DropAction, int, int, *QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), 731539533644753052, func(data uintptr, action int32, row int32, column int32, parent uintptr) int32 { return b2i(fn(UnsafeNewQMimeData(unsafe.Pointer(data)), DropAction(action), int(row), int(column), UnsafeNewQModelIndex(unsafe.Pointer(parent)))) }) // slot "canDropMimeData"
}

func (this *QTransposeProxyModel) OnDropMimeData(fn func(*QMimeData, DropAction, int, int, *QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), -9049994684305682212, func(data uintptr, action int32, row int32, column int32, parent uintptr) int32 { return b2i(fn(UnsafeNewQMimeData(unsafe.Pointer(data)), DropAction(action), int(row), int(column), UnsafeNewQModelIndex(unsafe.Pointer(parent)))) }) // slot "dropMimeData"
}

func (this *QTransposeProxyModel) OnMimeTypes(fn func() []string) {
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

func (this *QTransposeProxyModel) OnSupportedDragActions(fn func() DropAction) {
	registerVirt(this.UnsafeGoID(), 3421601386088570762, func() int32 { return int32(fn()) }) // slot "supportedDragActions"
}

func (this *QTransposeProxyModel) OnSupportedDropActions(fn func() DropAction) {
	registerVirt(this.UnsafeGoID(), -7784642176186821205, func() int32 { return int32(fn()) }) // slot "supportedDropActions"
}

func (this *QTransposeProxyModel) OnMatch(fn func(*QModelIndex, int, *QVariant, int, MatchFlag) []*QModelIndex) {
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

func (this *QTransposeProxyModel) OnMultiData(fn func(*QModelIndex, *QModelRoleDataSpan)) {
	registerVirt(this.UnsafeGoID(), 731311060944096642, func(index uintptr, roleDataSpan uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(index)), UnsafeNewQModelRoleDataSpan(unsafe.Pointer(roleDataSpan))) }) // slot "multiData"
}

func (this *QTransposeProxyModel) OnResetInternalData(fn func()) {
	registerVirt(this.UnsafeGoID(), -4510675770784051183, func() { fn() }) // slot "resetInternalData"
}

func (this *QTransposeProxyModel) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QTransposeProxyModel) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QTransposeProxyModel) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QTransposeProxyModel) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QTransposeProxyModel) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QTransposeProxyModel) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QTransposeProxyModel) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QTransposeProxyModel) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QTransposeProxyModel_delete(this.h)
	}
	this.UnsafeRelease()
}
