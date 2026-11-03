package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QAbstractTableModel_chain uintptr
	addr_QAbstractTableModel_tr uintptr
	addr_QAbstractTableModel_createIndex uintptr
	addr_QAbstractTableModel_createIndex2 uintptr
	addr_QAbstractTableModel_encodeData uintptr
	addr_QAbstractTableModel_decodeData uintptr
	addr_QAbstractTableModel_beginInsertRows uintptr
	addr_QAbstractTableModel_endInsertRows uintptr
	addr_QAbstractTableModel_beginRemoveRows uintptr
	addr_QAbstractTableModel_endRemoveRows uintptr
	addr_QAbstractTableModel_beginMoveRows uintptr
	addr_QAbstractTableModel_endMoveRows uintptr
	addr_QAbstractTableModel_beginInsertColumns uintptr
	addr_QAbstractTableModel_endInsertColumns uintptr
	addr_QAbstractTableModel_beginRemoveColumns uintptr
	addr_QAbstractTableModel_endRemoveColumns uintptr
	addr_QAbstractTableModel_beginMoveColumns uintptr
	addr_QAbstractTableModel_endMoveColumns uintptr
	addr_QAbstractTableModel_beginResetModel uintptr
	addr_QAbstractTableModel_endResetModel uintptr
	addr_QAbstractTableModel_changePersistentIndex uintptr
	addr_QAbstractTableModel_changePersistentIndexList uintptr
	addr_QAbstractTableModel_persistentIndexList uintptr
	addr_QAbstractTableModel_connect uintptr
	addr_QAbstractTableModel_connect2 uintptr
	addr_QAbstractTableModel_sender uintptr
	addr_QAbstractTableModel_senderSignalIndex uintptr
	addr_QAbstractTableModel_receivers uintptr
	addr_QAbstractTableModel_isSignalConnected uintptr
	addr_QAbstractTableModel_tr2 uintptr
	addr_QAbstractTableModel_tr3 uintptr
	addr_QAbstractTableModel_connect3 uintptr
	addr_QAbstractTableModel_connect4 uintptr
	addr_QAbstractTableModel_metaObject_default uintptr
	addr_QAbstractTableModel_index_default uintptr
	addr_QAbstractTableModel_sibling_default uintptr
	addr_QAbstractTableModel_dropMimeData_default uintptr
	addr_QAbstractTableModel_flags_default uintptr
	addr_QAbstractTableModel_setData_default uintptr
	addr_QAbstractTableModel_headerData_default uintptr
	addr_QAbstractTableModel_setHeaderData_default uintptr
	addr_QAbstractTableModel_clearItemData_default uintptr
	addr_QAbstractTableModel_mimeTypes_default uintptr
	addr_QAbstractTableModel_mimeData_default uintptr
	addr_QAbstractTableModel_canDropMimeData_default uintptr
	addr_QAbstractTableModel_supportedDropActions_default uintptr
	addr_QAbstractTableModel_supportedDragActions_default uintptr
	addr_QAbstractTableModel_insertRows_default uintptr
	addr_QAbstractTableModel_insertColumns_default uintptr
	addr_QAbstractTableModel_removeRows_default uintptr
	addr_QAbstractTableModel_removeColumns_default uintptr
	addr_QAbstractTableModel_moveRows_default uintptr
	addr_QAbstractTableModel_moveColumns_default uintptr
	addr_QAbstractTableModel_fetchMore_default uintptr
	addr_QAbstractTableModel_canFetchMore_default uintptr
	addr_QAbstractTableModel_sort_default uintptr
	addr_QAbstractTableModel_buddy_default uintptr
	addr_QAbstractTableModel_match_default uintptr
	addr_QAbstractTableModel_span_default uintptr
	addr_QAbstractTableModel_multiData_default uintptr
	addr_QAbstractTableModel_submit_default uintptr
	addr_QAbstractTableModel_revert_default uintptr
	addr_QAbstractTableModel_resetInternalData_default uintptr
	addr_QAbstractTableModel_event_default uintptr
	addr_QAbstractTableModel_eventFilter_default uintptr
	addr_QAbstractTableModel_timerEvent_default uintptr
	addr_QAbstractTableModel_childEvent_default uintptr
	addr_QAbstractTableModel_customEvent_default uintptr
	addr_QAbstractTableModel_connectNotify_default uintptr
	addr_QAbstractTableModel_disconnectNotify_default uintptr
	pg_QAbstractTableModel_connect_headerDataChanged func(uintptr, int64)
	pg_QAbstractTableModel_connect_layoutChanged func(uintptr, int64)
	pg_QAbstractTableModel_connect_layoutChanged2 func(uintptr, int64)
	pg_QAbstractTableModel_connect_layoutAboutToBeChanged func(uintptr, int64)
	pg_QAbstractTableModel_connect_layoutAboutToBeChanged2 func(uintptr, int64)
	pg_QAbstractTableModel_connect_rowsAboutToBeInserted func(uintptr, int64)
	pg_QAbstractTableModel_connect_rowsInserted func(uintptr, int64)
	pg_QAbstractTableModel_connect_rowsAboutToBeRemoved func(uintptr, int64)
	pg_QAbstractTableModel_connect_rowsRemoved func(uintptr, int64)
	pg_QAbstractTableModel_connect_columnsAboutToBeInserted func(uintptr, int64)
	pg_QAbstractTableModel_connect_columnsInserted func(uintptr, int64)
	pg_QAbstractTableModel_connect_columnsAboutToBeRemoved func(uintptr, int64)
	pg_QAbstractTableModel_connect_columnsRemoved func(uintptr, int64)
	pg_QAbstractTableModel_connect_modelAboutToBeReset func(uintptr, int64)
	pg_QAbstractTableModel_connect_modelReset func(uintptr, int64)
	pg_QAbstractTableModel_connect_rowsAboutToBeMoved func(uintptr, int64)
	pg_QAbstractTableModel_connect_rowsMoved func(uintptr, int64)
	pg_QAbstractTableModel_connect_columnsAboutToBeMoved func(uintptr, int64)
	pg_QAbstractTableModel_connect_columnsMoved func(uintptr, int64)
	pg_QAbstractTableModel_connect_destroyed func(uintptr, int64)
	pg_QAbstractTableModel_connect_destroyed2 func(uintptr, int64)
	pg_QAbstractTableModel_connect_objectNameChanged func(uintptr, int64)
	pg_QAbstractTableModel__disconnect func(int64)
	pg_QAbstractTableModel__set_override func(uintptr, int32)
	pg_QAbstractTableModel__Go_new func(int64, uintptr) uintptr
	pg_QAbstractTableModel_delete func(uintptr)
)

func init() {
	symaddr(&addr_QAbstractTableModel_chain, "QAbstractTableModel_chain")
	symaddr(&addr_QAbstractTableModel_tr, "QAbstractTableModel_tr")
	symaddr(&addr_QAbstractTableModel_createIndex, "QAbstractTableModel_createIndex")
	symaddr(&addr_QAbstractTableModel_createIndex2, "QAbstractTableModel_createIndex2")
	symaddr(&addr_QAbstractTableModel_encodeData, "QAbstractTableModel_encodeData")
	symaddr(&addr_QAbstractTableModel_decodeData, "QAbstractTableModel_decodeData")
	symaddr(&addr_QAbstractTableModel_beginInsertRows, "QAbstractTableModel_beginInsertRows")
	symaddr(&addr_QAbstractTableModel_endInsertRows, "QAbstractTableModel_endInsertRows")
	symaddr(&addr_QAbstractTableModel_beginRemoveRows, "QAbstractTableModel_beginRemoveRows")
	symaddr(&addr_QAbstractTableModel_endRemoveRows, "QAbstractTableModel_endRemoveRows")
	symaddr(&addr_QAbstractTableModel_beginMoveRows, "QAbstractTableModel_beginMoveRows")
	symaddr(&addr_QAbstractTableModel_endMoveRows, "QAbstractTableModel_endMoveRows")
	symaddr(&addr_QAbstractTableModel_beginInsertColumns, "QAbstractTableModel_beginInsertColumns")
	symaddr(&addr_QAbstractTableModel_endInsertColumns, "QAbstractTableModel_endInsertColumns")
	symaddr(&addr_QAbstractTableModel_beginRemoveColumns, "QAbstractTableModel_beginRemoveColumns")
	symaddr(&addr_QAbstractTableModel_endRemoveColumns, "QAbstractTableModel_endRemoveColumns")
	symaddr(&addr_QAbstractTableModel_beginMoveColumns, "QAbstractTableModel_beginMoveColumns")
	symaddr(&addr_QAbstractTableModel_endMoveColumns, "QAbstractTableModel_endMoveColumns")
	symaddr(&addr_QAbstractTableModel_beginResetModel, "QAbstractTableModel_beginResetModel")
	symaddr(&addr_QAbstractTableModel_endResetModel, "QAbstractTableModel_endResetModel")
	symaddr(&addr_QAbstractTableModel_changePersistentIndex, "QAbstractTableModel_changePersistentIndex")
	symaddr(&addr_QAbstractTableModel_changePersistentIndexList, "QAbstractTableModel_changePersistentIndexList")
	symaddr(&addr_QAbstractTableModel_persistentIndexList, "QAbstractTableModel_persistentIndexList")
	symaddr(&addr_QAbstractTableModel_connect, "QAbstractTableModel_connect")
	symaddr(&addr_QAbstractTableModel_connect2, "QAbstractTableModel_connect2")
	symaddr(&addr_QAbstractTableModel_sender, "QAbstractTableModel_sender")
	symaddr(&addr_QAbstractTableModel_senderSignalIndex, "QAbstractTableModel_senderSignalIndex")
	symaddr(&addr_QAbstractTableModel_receivers, "QAbstractTableModel_receivers")
	symaddr(&addr_QAbstractTableModel_isSignalConnected, "QAbstractTableModel_isSignalConnected")
	symaddr(&addr_QAbstractTableModel_tr2, "QAbstractTableModel_tr2")
	symaddr(&addr_QAbstractTableModel_tr3, "QAbstractTableModel_tr3")
	symaddr(&addr_QAbstractTableModel_connect3, "QAbstractTableModel_connect3")
	symaddr(&addr_QAbstractTableModel_connect4, "QAbstractTableModel_connect4")
	symaddr(&addr_QAbstractTableModel_metaObject_default, "QAbstractTableModel_metaObject_default")
	symaddr(&addr_QAbstractTableModel_index_default, "QAbstractTableModel_index_default")
	symaddr(&addr_QAbstractTableModel_sibling_default, "QAbstractTableModel_sibling_default")
	symaddr(&addr_QAbstractTableModel_dropMimeData_default, "QAbstractTableModel_dropMimeData_default")
	symaddr(&addr_QAbstractTableModel_flags_default, "QAbstractTableModel_flags_default")
	symaddr(&addr_QAbstractTableModel_setData_default, "QAbstractTableModel_setData_default")
	symaddr(&addr_QAbstractTableModel_headerData_default, "QAbstractTableModel_headerData_default")
	symaddr(&addr_QAbstractTableModel_setHeaderData_default, "QAbstractTableModel_setHeaderData_default")
	symaddr(&addr_QAbstractTableModel_clearItemData_default, "QAbstractTableModel_clearItemData_default")
	symaddr(&addr_QAbstractTableModel_mimeTypes_default, "QAbstractTableModel_mimeTypes_default")
	symaddr(&addr_QAbstractTableModel_mimeData_default, "QAbstractTableModel_mimeData_default")
	symaddr(&addr_QAbstractTableModel_canDropMimeData_default, "QAbstractTableModel_canDropMimeData_default")
	symaddr(&addr_QAbstractTableModel_supportedDropActions_default, "QAbstractTableModel_supportedDropActions_default")
	symaddr(&addr_QAbstractTableModel_supportedDragActions_default, "QAbstractTableModel_supportedDragActions_default")
	symaddr(&addr_QAbstractTableModel_insertRows_default, "QAbstractTableModel_insertRows_default")
	symaddr(&addr_QAbstractTableModel_insertColumns_default, "QAbstractTableModel_insertColumns_default")
	symaddr(&addr_QAbstractTableModel_removeRows_default, "QAbstractTableModel_removeRows_default")
	symaddr(&addr_QAbstractTableModel_removeColumns_default, "QAbstractTableModel_removeColumns_default")
	symaddr(&addr_QAbstractTableModel_moveRows_default, "QAbstractTableModel_moveRows_default")
	symaddr(&addr_QAbstractTableModel_moveColumns_default, "QAbstractTableModel_moveColumns_default")
	symaddr(&addr_QAbstractTableModel_fetchMore_default, "QAbstractTableModel_fetchMore_default")
	symaddr(&addr_QAbstractTableModel_canFetchMore_default, "QAbstractTableModel_canFetchMore_default")
	symaddr(&addr_QAbstractTableModel_sort_default, "QAbstractTableModel_sort_default")
	symaddr(&addr_QAbstractTableModel_buddy_default, "QAbstractTableModel_buddy_default")
	symaddr(&addr_QAbstractTableModel_match_default, "QAbstractTableModel_match_default")
	symaddr(&addr_QAbstractTableModel_span_default, "QAbstractTableModel_span_default")
	symaddr(&addr_QAbstractTableModel_multiData_default, "QAbstractTableModel_multiData_default")
	symaddr(&addr_QAbstractTableModel_submit_default, "QAbstractTableModel_submit_default")
	symaddr(&addr_QAbstractTableModel_revert_default, "QAbstractTableModel_revert_default")
	symaddr(&addr_QAbstractTableModel_resetInternalData_default, "QAbstractTableModel_resetInternalData_default")
	symaddr(&addr_QAbstractTableModel_event_default, "QAbstractTableModel_event_default")
	symaddr(&addr_QAbstractTableModel_eventFilter_default, "QAbstractTableModel_eventFilter_default")
	symaddr(&addr_QAbstractTableModel_timerEvent_default, "QAbstractTableModel_timerEvent_default")
	symaddr(&addr_QAbstractTableModel_childEvent_default, "QAbstractTableModel_childEvent_default")
	symaddr(&addr_QAbstractTableModel_customEvent_default, "QAbstractTableModel_customEvent_default")
	symaddr(&addr_QAbstractTableModel_connectNotify_default, "QAbstractTableModel_connectNotify_default")
	symaddr(&addr_QAbstractTableModel_disconnectNotify_default, "QAbstractTableModel_disconnectNotify_default")
	reg(&pg_QAbstractTableModel_connect_headerDataChanged, "QAbstractTableModel_connect_headerDataChanged")
	reg(&pg_QAbstractTableModel_connect_layoutChanged, "QAbstractTableModel_connect_layoutChanged")
	reg(&pg_QAbstractTableModel_connect_layoutChanged2, "QAbstractTableModel_connect_layoutChanged2")
	reg(&pg_QAbstractTableModel_connect_layoutAboutToBeChanged, "QAbstractTableModel_connect_layoutAboutToBeChanged")
	reg(&pg_QAbstractTableModel_connect_layoutAboutToBeChanged2, "QAbstractTableModel_connect_layoutAboutToBeChanged2")
	reg(&pg_QAbstractTableModel_connect_rowsAboutToBeInserted, "QAbstractTableModel_connect_rowsAboutToBeInserted")
	reg(&pg_QAbstractTableModel_connect_rowsInserted, "QAbstractTableModel_connect_rowsInserted")
	reg(&pg_QAbstractTableModel_connect_rowsAboutToBeRemoved, "QAbstractTableModel_connect_rowsAboutToBeRemoved")
	reg(&pg_QAbstractTableModel_connect_rowsRemoved, "QAbstractTableModel_connect_rowsRemoved")
	reg(&pg_QAbstractTableModel_connect_columnsAboutToBeInserted, "QAbstractTableModel_connect_columnsAboutToBeInserted")
	reg(&pg_QAbstractTableModel_connect_columnsInserted, "QAbstractTableModel_connect_columnsInserted")
	reg(&pg_QAbstractTableModel_connect_columnsAboutToBeRemoved, "QAbstractTableModel_connect_columnsAboutToBeRemoved")
	reg(&pg_QAbstractTableModel_connect_columnsRemoved, "QAbstractTableModel_connect_columnsRemoved")
	reg(&pg_QAbstractTableModel_connect_modelAboutToBeReset, "QAbstractTableModel_connect_modelAboutToBeReset")
	reg(&pg_QAbstractTableModel_connect_modelReset, "QAbstractTableModel_connect_modelReset")
	reg(&pg_QAbstractTableModel_connect_rowsAboutToBeMoved, "QAbstractTableModel_connect_rowsAboutToBeMoved")
	reg(&pg_QAbstractTableModel_connect_rowsMoved, "QAbstractTableModel_connect_rowsMoved")
	reg(&pg_QAbstractTableModel_connect_columnsAboutToBeMoved, "QAbstractTableModel_connect_columnsAboutToBeMoved")
	reg(&pg_QAbstractTableModel_connect_columnsMoved, "QAbstractTableModel_connect_columnsMoved")
	reg(&pg_QAbstractTableModel_connect_destroyed, "QAbstractTableModel_connect_destroyed")
	reg(&pg_QAbstractTableModel_connect_destroyed2, "QAbstractTableModel_connect_destroyed2")
	reg(&pg_QAbstractTableModel_connect_objectNameChanged, "QAbstractTableModel_connect_objectNameChanged")
	reg(&pg_QAbstractTableModel__disconnect, "QAbstractTableModel__disconnect")
	reg(&pg_QAbstractTableModel__set_override, "QAbstractTableModel__set_override")
	reg(&pg_QAbstractTableModel__Go_new, "QAbstractTableModel__Go_new")
	reg(&pg_QAbstractTableModel_delete, "QAbstractTableModel_delete")
}

type QAbstractTableModel struct {
	h uintptr
	*QAbstractItemModel
}

func bindQAbstractTableModel(h uintptr) *QAbstractTableModel {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QAbstractTableModel_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QAbstractTableModel{h: h, QAbstractItemModel: &QAbstractItemModel{h: __c[0], QObject: &QObject{h: __c[1]}}}
}

func newownedQAbstractTableModel(h uintptr) *QAbstractTableModel { return bindQAbstractTableModel(h) }

func (this *QAbstractTableModel) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QAbstractTableModel) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQAbstractTableModel(h unsafe.Pointer) *QAbstractTableModel { return bindQAbstractTableModel(uintptr(h)) }
func UnsafeNewQAbstractTableModelOwned(h unsafe.Pointer) *QAbstractTableModel { return newownedQAbstractTableModel(uintptr(h)) }

func (this *QAbstractTableModel) UnsafeRelease() {
	this.h = 0
	if this.QAbstractItemModel != nil {
		this.QAbstractItemModel.UnsafeRelease()
	}
}

func QAbstractTableModel_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAbstractTableModel_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QAbstractTableModel) CreateIndex(row int, column int) *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QAbstractTableModel_createIndex, this.h, uintptr(row), uintptr(column))
	runtime.KeepAlive(this)
	return newownedQModelIndex(__r)
}

func (this *QAbstractTableModel) CreateIndex2(row int, column int, id uintptr) *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QAbstractTableModel_createIndex2, this.h, uintptr(row), uintptr(column), uintptr(id))
	runtime.KeepAlive(this)
	return newownedQModelIndex(__r)
}

func (this *QAbstractTableModel) EncodeData(indexes []*QModelIndex, stream *QDataStream) {
	indexes_arr := make([]uintptr, len(indexes))
	for __i, __e := range indexes {
		indexes_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var indexes_p uintptr
	if len(indexes_arr) > 0 {
		indexes_p = uintptr(unsafe.Pointer(&indexes_arr[0]))
	}
	purego.SyscallN(addr_QAbstractTableModel_encodeData, this.h, indexes_p, uintptr(len(indexes)), uintptr(stream.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(indexes)
	runtime.KeepAlive(indexes_arr)
	runtime.KeepAlive(stream)
}

func (this *QAbstractTableModel) DecodeData(row int, column int, parent *QModelIndex, stream *QDataStream) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractTableModel_decodeData, this.h, uintptr(row), uintptr(column), uintptr(parent.UnsafePointer()), uintptr(stream.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	runtime.KeepAlive(stream)
	return __r != 0
}

func (this *QAbstractTableModel) BeginInsertRows(parent *QModelIndex, first int, last int) {
	purego.SyscallN(addr_QAbstractTableModel_beginInsertRows, this.h, uintptr(parent.UnsafePointer()), uintptr(first), uintptr(last))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
}

func (this *QAbstractTableModel) EndInsertRows() {
	purego.SyscallN(addr_QAbstractTableModel_endInsertRows, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractTableModel) BeginRemoveRows(parent *QModelIndex, first int, last int) {
	purego.SyscallN(addr_QAbstractTableModel_beginRemoveRows, this.h, uintptr(parent.UnsafePointer()), uintptr(first), uintptr(last))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
}

func (this *QAbstractTableModel) EndRemoveRows() {
	purego.SyscallN(addr_QAbstractTableModel_endRemoveRows, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractTableModel) BeginMoveRows(sourceParent *QModelIndex, sourceFirst int, sourceLast int, destinationParent *QModelIndex, destinationRow int) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractTableModel_beginMoveRows, this.h, uintptr(sourceParent.UnsafePointer()), uintptr(sourceFirst), uintptr(sourceLast), uintptr(destinationParent.UnsafePointer()), uintptr(destinationRow))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sourceParent)
	runtime.KeepAlive(destinationParent)
	return __r != 0
}

func (this *QAbstractTableModel) EndMoveRows() {
	purego.SyscallN(addr_QAbstractTableModel_endMoveRows, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractTableModel) BeginInsertColumns(parent *QModelIndex, first int, last int) {
	purego.SyscallN(addr_QAbstractTableModel_beginInsertColumns, this.h, uintptr(parent.UnsafePointer()), uintptr(first), uintptr(last))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
}

func (this *QAbstractTableModel) EndInsertColumns() {
	purego.SyscallN(addr_QAbstractTableModel_endInsertColumns, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractTableModel) BeginRemoveColumns(parent *QModelIndex, first int, last int) {
	purego.SyscallN(addr_QAbstractTableModel_beginRemoveColumns, this.h, uintptr(parent.UnsafePointer()), uintptr(first), uintptr(last))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
}

func (this *QAbstractTableModel) EndRemoveColumns() {
	purego.SyscallN(addr_QAbstractTableModel_endRemoveColumns, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractTableModel) BeginMoveColumns(sourceParent *QModelIndex, sourceFirst int, sourceLast int, destinationParent *QModelIndex, destinationColumn int) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractTableModel_beginMoveColumns, this.h, uintptr(sourceParent.UnsafePointer()), uintptr(sourceFirst), uintptr(sourceLast), uintptr(destinationParent.UnsafePointer()), uintptr(destinationColumn))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sourceParent)
	runtime.KeepAlive(destinationParent)
	return __r != 0
}

func (this *QAbstractTableModel) EndMoveColumns() {
	purego.SyscallN(addr_QAbstractTableModel_endMoveColumns, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractTableModel) BeginResetModel() {
	purego.SyscallN(addr_QAbstractTableModel_beginResetModel, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractTableModel) EndResetModel() {
	purego.SyscallN(addr_QAbstractTableModel_endResetModel, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractTableModel) ChangePersistentIndex(from *QModelIndex, to *QModelIndex) {
	purego.SyscallN(addr_QAbstractTableModel_changePersistentIndex, this.h, uintptr(from.UnsafePointer()), uintptr(to.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(from)
	runtime.KeepAlive(to)
}

func (this *QAbstractTableModel) ChangePersistentIndexList(from []*QModelIndex, to []*QModelIndex) {
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
	purego.SyscallN(addr_QAbstractTableModel_changePersistentIndexList, this.h, from_p, uintptr(len(from)), to_p, uintptr(len(to)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(from)
	runtime.KeepAlive(from_arr)
	runtime.KeepAlive(to)
	runtime.KeepAlive(to_arr)
}

func (this *QAbstractTableModel) PersistentIndexList() []*QModelIndex {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAbstractTableModel_persistentIndexList, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QModelIndex, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQModelIndex(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QAbstractTableModel_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QAbstractTableModel_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QAbstractTableModel_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QAbstractTableModel_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QAbstractTableModel) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QAbstractTableModel_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QAbstractTableModel) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QAbstractTableModel_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAbstractTableModel) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QAbstractTableModel_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QAbstractTableModel) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractTableModel_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QAbstractTableModel_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAbstractTableModel_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QAbstractTableModel_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAbstractTableModel_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QAbstractTableModel_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QAbstractTableModel_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QAbstractTableModel_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QAbstractTableModel_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QAbstractTableModel) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QAbstractTableModel_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QAbstractTableModel) IndexDefault(row int, column int, parent *QModelIndex) *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QAbstractTableModel_index_default, this.h, uintptr(row), uintptr(column), uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return newownedQModelIndex(__r)
}

func (this *QAbstractTableModel) SiblingDefault(row int, column int, idx *QModelIndex) *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QAbstractTableModel_sibling_default, this.h, uintptr(row), uintptr(column), uintptr(idx.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(idx)
	return newownedQModelIndex(__r)
}

func (this *QAbstractTableModel) DropMimeDataDefault(data *QMimeData, action DropAction, row int, column int, parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractTableModel_dropMimeData_default, this.h, uintptr(data.UnsafePointer()), uintptr(action), uintptr(row), uintptr(column), uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(data)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QAbstractTableModel) FlagsDefault(index *QModelIndex) ItemFlag {
	__r0, _, _ := purego.SyscallN(addr_QAbstractTableModel_flags_default, this.h, uintptr(index.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return ItemFlag(__r)
}

func (this *QAbstractTableModel) SetDataDefault(index *QModelIndex, value *QVariant, role int) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractTableModel_setData_default, this.h, uintptr(index.UnsafePointer()), uintptr(value.UnsafePointer()), uintptr(role))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	runtime.KeepAlive(value)
	return __r != 0
}

func (this *QAbstractTableModel) HeaderDataDefault(section int, orientation Orientation, role int) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QAbstractTableModel_headerData_default, this.h, uintptr(section), uintptr(orientation), uintptr(role))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QAbstractTableModel) SetHeaderDataDefault(section int, orientation Orientation, value *QVariant, role int) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractTableModel_setHeaderData_default, this.h, uintptr(section), uintptr(orientation), uintptr(value.UnsafePointer()), uintptr(role))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(value)
	return __r != 0
}

func (this *QAbstractTableModel) ClearItemDataDefault(index *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractTableModel_clearItemData_default, this.h, uintptr(index.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return __r != 0
}

func (this *QAbstractTableModel) MimeTypesDefault() []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAbstractTableModel_mimeTypes_default, this.h, uintptr(unsafe.Pointer(&__n)))
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

func (this *QAbstractTableModel) MimeDataDefault(indexes []*QModelIndex) *QMimeData {
	indexes_arr := make([]uintptr, len(indexes))
	for __i, __e := range indexes {
		indexes_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var indexes_p uintptr
	if len(indexes_arr) > 0 {
		indexes_p = uintptr(unsafe.Pointer(&indexes_arr[0]))
	}
	__r, _, _ := purego.SyscallN(addr_QAbstractTableModel_mimeData_default, this.h, indexes_p, uintptr(len(indexes)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(indexes)
	runtime.KeepAlive(indexes_arr)
	return bindQMimeData(__r)
}

func (this *QAbstractTableModel) CanDropMimeDataDefault(data *QMimeData, action DropAction, row int, column int, parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractTableModel_canDropMimeData_default, this.h, uintptr(data.UnsafePointer()), uintptr(action), uintptr(row), uintptr(column), uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(data)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QAbstractTableModel) SupportedDropActionsDefault() DropAction {
	__r0, _, _ := purego.SyscallN(addr_QAbstractTableModel_supportedDropActions_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return DropAction(__r)
}

func (this *QAbstractTableModel) SupportedDragActionsDefault() DropAction {
	__r0, _, _ := purego.SyscallN(addr_QAbstractTableModel_supportedDragActions_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return DropAction(__r)
}

func (this *QAbstractTableModel) InsertRowsDefault(row int, count int, parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractTableModel_insertRows_default, this.h, uintptr(row), uintptr(count), uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QAbstractTableModel) InsertColumnsDefault(column int, count int, parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractTableModel_insertColumns_default, this.h, uintptr(column), uintptr(count), uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QAbstractTableModel) RemoveRowsDefault(row int, count int, parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractTableModel_removeRows_default, this.h, uintptr(row), uintptr(count), uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QAbstractTableModel) RemoveColumnsDefault(column int, count int, parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractTableModel_removeColumns_default, this.h, uintptr(column), uintptr(count), uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QAbstractTableModel) MoveRowsDefault(sourceParent *QModelIndex, sourceRow int, count int, destinationParent *QModelIndex, destinationChild int) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractTableModel_moveRows_default, this.h, uintptr(sourceParent.UnsafePointer()), uintptr(sourceRow), uintptr(count), uintptr(destinationParent.UnsafePointer()), uintptr(destinationChild))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sourceParent)
	runtime.KeepAlive(destinationParent)
	return __r != 0
}

func (this *QAbstractTableModel) MoveColumnsDefault(sourceParent *QModelIndex, sourceColumn int, count int, destinationParent *QModelIndex, destinationChild int) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractTableModel_moveColumns_default, this.h, uintptr(sourceParent.UnsafePointer()), uintptr(sourceColumn), uintptr(count), uintptr(destinationParent.UnsafePointer()), uintptr(destinationChild))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sourceParent)
	runtime.KeepAlive(destinationParent)
	return __r != 0
}

func (this *QAbstractTableModel) FetchMoreDefault(parent *QModelIndex) {
	purego.SyscallN(addr_QAbstractTableModel_fetchMore_default, this.h, uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
}

func (this *QAbstractTableModel) CanFetchMoreDefault(parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractTableModel_canFetchMore_default, this.h, uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QAbstractTableModel) SortDefault(column int, order SortOrder) {
	purego.SyscallN(addr_QAbstractTableModel_sort_default, this.h, uintptr(column), uintptr(order))
	runtime.KeepAlive(this)
}

func (this *QAbstractTableModel) BuddyDefault(index *QModelIndex) *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QAbstractTableModel_buddy_default, this.h, uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return newownedQModelIndex(__r)
}

func (this *QAbstractTableModel) MatchDefault(start *QModelIndex, role int, value *QVariant, hits int, flags MatchFlag) []*QModelIndex {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAbstractTableModel_match_default, this.h, uintptr(start.UnsafePointer()), uintptr(role), uintptr(value.UnsafePointer()), uintptr(hits), uintptr(flags), uintptr(unsafe.Pointer(&__n)))
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

func (this *QAbstractTableModel) SpanDefault(index *QModelIndex) *QSize {
	__r, _, _ := purego.SyscallN(addr_QAbstractTableModel_span_default, this.h, uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return newownedQSize(__r)
}

func (this *QAbstractTableModel) MultiDataDefault(index *QModelIndex, roleDataSpan *QModelRoleDataSpan) {
	purego.SyscallN(addr_QAbstractTableModel_multiData_default, this.h, uintptr(index.UnsafePointer()), uintptr(roleDataSpan.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	runtime.KeepAlive(roleDataSpan)
}

func (this *QAbstractTableModel) SubmitDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractTableModel_submit_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QAbstractTableModel) RevertDefault() {
	purego.SyscallN(addr_QAbstractTableModel_revert_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractTableModel) ResetInternalDataDefault() {
	purego.SyscallN(addr_QAbstractTableModel_resetInternalData_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractTableModel) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractTableModel_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QAbstractTableModel) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QAbstractTableModel_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QAbstractTableModel) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QAbstractTableModel_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractTableModel) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QAbstractTableModel_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractTableModel) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QAbstractTableModel_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QAbstractTableModel) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QAbstractTableModel_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QAbstractTableModel) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QAbstractTableModel_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QAbstractTableModel) OnHeaderDataChanged(fn func(Orientation, int, int)) int64 {
	id := registerCallback(func(orientation int32, first int32, last int32) { fn(Orientation(orientation), int(first), int(last)) })
	pg_QAbstractTableModel_connect_headerDataChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractTableModel) OnLayoutChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QAbstractTableModel_connect_layoutChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractTableModel) OnLayoutChanged2(fn func([]*QPersistentModelIndex, QAbstractItemModel__LayoutChangeHint)) int64 {
	id := registerCallback(func(parents uintptr, parents_len int64, hint int32) { fn(func() []*QPersistentModelIndex { if parents == 0 || parents_len == 0 { return nil }; __sl := unsafe.Slice((*unsafe.Pointer)(unsafe.Pointer(parents)), int(parents_len)); __o := make([]*QPersistentModelIndex, int(parents_len)); for __i := range __o { __o[__i] = UnsafeNewQPersistentModelIndex(__sl[__i]) }; return __o }(), QAbstractItemModel__LayoutChangeHint(hint)) })
	pg_QAbstractTableModel_connect_layoutChanged2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractTableModel) OnLayoutAboutToBeChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QAbstractTableModel_connect_layoutAboutToBeChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractTableModel) OnLayoutAboutToBeChanged2(fn func([]*QPersistentModelIndex, QAbstractItemModel__LayoutChangeHint)) int64 {
	id := registerCallback(func(parents uintptr, parents_len int64, hint int32) { fn(func() []*QPersistentModelIndex { if parents == 0 || parents_len == 0 { return nil }; __sl := unsafe.Slice((*unsafe.Pointer)(unsafe.Pointer(parents)), int(parents_len)); __o := make([]*QPersistentModelIndex, int(parents_len)); for __i := range __o { __o[__i] = UnsafeNewQPersistentModelIndex(__sl[__i]) }; return __o }(), QAbstractItemModel__LayoutChangeHint(hint)) })
	pg_QAbstractTableModel_connect_layoutAboutToBeChanged2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractTableModel) OnRowsAboutToBeInserted(fn func(*QModelIndex, int, int)) int64 {
	id := registerCallback(func(parent uintptr, first int32, last int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(first), int(last)) })
	pg_QAbstractTableModel_connect_rowsAboutToBeInserted(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractTableModel) OnRowsInserted(fn func(*QModelIndex, int, int)) int64 {
	id := registerCallback(func(parent uintptr, first int32, last int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(first), int(last)) })
	pg_QAbstractTableModel_connect_rowsInserted(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractTableModel) OnRowsAboutToBeRemoved(fn func(*QModelIndex, int, int)) int64 {
	id := registerCallback(func(parent uintptr, first int32, last int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(first), int(last)) })
	pg_QAbstractTableModel_connect_rowsAboutToBeRemoved(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractTableModel) OnRowsRemoved(fn func(*QModelIndex, int, int)) int64 {
	id := registerCallback(func(parent uintptr, first int32, last int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(first), int(last)) })
	pg_QAbstractTableModel_connect_rowsRemoved(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractTableModel) OnColumnsAboutToBeInserted(fn func(*QModelIndex, int, int)) int64 {
	id := registerCallback(func(parent uintptr, first int32, last int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(first), int(last)) })
	pg_QAbstractTableModel_connect_columnsAboutToBeInserted(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractTableModel) OnColumnsInserted(fn func(*QModelIndex, int, int)) int64 {
	id := registerCallback(func(parent uintptr, first int32, last int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(first), int(last)) })
	pg_QAbstractTableModel_connect_columnsInserted(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractTableModel) OnColumnsAboutToBeRemoved(fn func(*QModelIndex, int, int)) int64 {
	id := registerCallback(func(parent uintptr, first int32, last int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(first), int(last)) })
	pg_QAbstractTableModel_connect_columnsAboutToBeRemoved(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractTableModel) OnColumnsRemoved(fn func(*QModelIndex, int, int)) int64 {
	id := registerCallback(func(parent uintptr, first int32, last int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(first), int(last)) })
	pg_QAbstractTableModel_connect_columnsRemoved(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractTableModel) OnModelAboutToBeReset(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QAbstractTableModel_connect_modelAboutToBeReset(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractTableModel) OnModelReset(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QAbstractTableModel_connect_modelReset(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractTableModel) OnRowsAboutToBeMoved(fn func(*QModelIndex, int, int, *QModelIndex, int)) int64 {
	id := registerCallback(func(sourceParent uintptr, sourceStart int32, sourceEnd int32, destinationParent uintptr, destinationRow int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(sourceParent)), int(sourceStart), int(sourceEnd), UnsafeNewQModelIndex(unsafe.Pointer(destinationParent)), int(destinationRow)) })
	pg_QAbstractTableModel_connect_rowsAboutToBeMoved(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractTableModel) OnRowsMoved(fn func(*QModelIndex, int, int, *QModelIndex, int)) int64 {
	id := registerCallback(func(sourceParent uintptr, sourceStart int32, sourceEnd int32, destinationParent uintptr, destinationRow int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(sourceParent)), int(sourceStart), int(sourceEnd), UnsafeNewQModelIndex(unsafe.Pointer(destinationParent)), int(destinationRow)) })
	pg_QAbstractTableModel_connect_rowsMoved(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractTableModel) OnColumnsAboutToBeMoved(fn func(*QModelIndex, int, int, *QModelIndex, int)) int64 {
	id := registerCallback(func(sourceParent uintptr, sourceStart int32, sourceEnd int32, destinationParent uintptr, destinationColumn int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(sourceParent)), int(sourceStart), int(sourceEnd), UnsafeNewQModelIndex(unsafe.Pointer(destinationParent)), int(destinationColumn)) })
	pg_QAbstractTableModel_connect_columnsAboutToBeMoved(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractTableModel) OnColumnsMoved(fn func(*QModelIndex, int, int, *QModelIndex, int)) int64 {
	id := registerCallback(func(sourceParent uintptr, sourceStart int32, sourceEnd int32, destinationParent uintptr, destinationColumn int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(sourceParent)), int(sourceStart), int(sourceEnd), UnsafeNewQModelIndex(unsafe.Pointer(destinationParent)), int(destinationColumn)) })
	pg_QAbstractTableModel_connect_columnsMoved(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractTableModel) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QAbstractTableModel_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractTableModel) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QAbstractTableModel_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractTableModel) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QAbstractTableModel_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractTableModel) Disconnect(handle int64) {
	pg_QAbstractTableModel__disconnect(handle)
	unregisterCallback(handle)
}

func QAbstractTableModel_virtIdx(slot int64) int32 {
	switch slot {
	case 1960960171956884010:
		return 0
	case -8256963024669584834:
		return 1
	case -8837375942098466299:
		return 2
	case 1607084412591819503:
		return 3
	case -8948777187306027381:
		return 4
	case 397329346813782521:
		return 5
	case -9049994684305682212:
		return 6
	case 1703382833545304812:
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
	case -7784642176186821205:
		return 15
	case 3421601386088570762:
		return 16
	case -2775729825962395603:
		return 17
	case -8973293846230792237:
		return 18
	case 5104072549182222656:
		return 19
	case 2078548737056540420:
		return 20
	case -6074819356947405415:
		return 21
	case -6321007591388309905:
		return 22
	case 3553294593611130806:
		return 23
	case 8639262066473964150:
		return 24
	case 6912441694453382161:
		return 25
	case 6112860112248786499:
		return 26
	case -4341501217981328394:
		return 27
	case -8395342916587040543:
		return 28
	case 731311060944096642:
		return 29
	case -3788331537442685179:
		return 30
	case -1522951607536117149:
		return 31
	case -4510675770784051183:
		return 32
	case 1237236456312824991:
		return 33
	case 5217623514630858297:
		return 34
	case -5409927581812374498:
		return 35
	case -1185315367526857833:
		return 36
	case 3769868229123868474:
		return 37
	case 7379029982170503008:
		return 38
	case -568369772982215962:
		return 39
	}
	return -1
}

func QAbstractTableModel_enable(h uintptr, idx int32) { pg_QAbstractTableModel__set_override(h, idx) }

func NewQAbstractTableModel() *QAbstractTableModel {
	id := nextVirtID()
	o := newownedQAbstractTableModel(pg_QAbstractTableModel__Go_new(id, 0))
	o.UnsafeSetGoID(id)
	registerInstance(id, o.h, QAbstractTableModel_virtIdx, QAbstractTableModel_enable)
	return o
}

func NewQAbstractTableModel2(parent *QObject) *QAbstractTableModel {
	id := nextVirtID()
	o := newownedQAbstractTableModel(pg_QAbstractTableModel__Go_new(id, uintptr(parent.UnsafePointer())))
	o.UnsafeSetGoID(id)
	registerInstance(id, o.h, QAbstractTableModel_virtIdx, QAbstractTableModel_enable)
	return o
}

func (this *QAbstractTableModel) OnRowCount(fn func(*QModelIndex) int) {
	registerVirt(this.UnsafeGoID(), 1960960171956884010, func(parent uintptr) int32 { return int32(fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)))) }) // slot "rowCount"
}

func (this *QAbstractTableModel) OnColumnCount(fn func(*QModelIndex) int) {
	registerVirt(this.UnsafeGoID(), -8256963024669584834, func(parent uintptr) int32 { return int32(fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)))) }) // slot "columnCount"
}

func (this *QAbstractTableModel) OnData(fn func(*QModelIndex, int) *QVariant) {
	registerVirt(this.UnsafeGoID(), -8837375942098466299, func(index uintptr, role int32) uintptr { return uintptr(fn(UnsafeNewQModelIndex(unsafe.Pointer(index)), int(role)).UnsafePointer()) }) // slot "data"
}

func (this *QAbstractTableModel) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QAbstractTableModel) OnIndex(fn func(int, int, *QModelIndex) *QModelIndex) {
	registerVirt(this.UnsafeGoID(), -8948777187306027381, func(row int32, column int32, parent uintptr) uintptr { return uintptr(fn(int(row), int(column), UnsafeNewQModelIndex(unsafe.Pointer(parent))).UnsafePointer()) }) // slot "index"
}

func (this *QAbstractTableModel) OnSibling(fn func(int, int, *QModelIndex) *QModelIndex) {
	registerVirt(this.UnsafeGoID(), 397329346813782521, func(row int32, column int32, idx uintptr) uintptr { return uintptr(fn(int(row), int(column), UnsafeNewQModelIndex(unsafe.Pointer(idx))).UnsafePointer()) }) // slot "sibling"
}

func (this *QAbstractTableModel) OnDropMimeData(fn func(*QMimeData, DropAction, int, int, *QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), -9049994684305682212, func(data uintptr, action int32, row int32, column int32, parent uintptr) int32 { return b2i(fn(UnsafeNewQMimeData(unsafe.Pointer(data)), DropAction(action), int(row), int(column), UnsafeNewQModelIndex(unsafe.Pointer(parent)))) }) // slot "dropMimeData"
}

func (this *QAbstractTableModel) OnFlags(fn func(*QModelIndex) ItemFlag) {
	registerVirt(this.UnsafeGoID(), 1703382833545304812, func(index uintptr) int32 { return int32(fn(UnsafeNewQModelIndex(unsafe.Pointer(index)))) }) // slot "flags"
}

func (this *QAbstractTableModel) OnSetData(fn func(*QModelIndex, *QVariant, int) bool) {
	registerVirt(this.UnsafeGoID(), 35191454575202219, func(index uintptr, value uintptr, role int32) int32 { return b2i(fn(UnsafeNewQModelIndex(unsafe.Pointer(index)), UnsafeNewQVariant(unsafe.Pointer(value)), int(role))) }) // slot "setData"
}

func (this *QAbstractTableModel) OnHeaderData(fn func(int, Orientation, int) *QVariant) {
	registerVirt(this.UnsafeGoID(), -3297140021896693688, func(section int32, orientation int32, role int32) uintptr { return uintptr(fn(int(section), Orientation(orientation), int(role)).UnsafePointer()) }) // slot "headerData"
}

func (this *QAbstractTableModel) OnSetHeaderData(fn func(int, Orientation, *QVariant, int) bool) {
	registerVirt(this.UnsafeGoID(), 4852200306590054850, func(section int32, orientation int32, value uintptr, role int32) int32 { return b2i(fn(int(section), Orientation(orientation), UnsafeNewQVariant(unsafe.Pointer(value)), int(role))) }) // slot "setHeaderData"
}

func (this *QAbstractTableModel) OnClearItemData(fn func(*QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), 8914190500520542409, func(index uintptr) int32 { return b2i(fn(UnsafeNewQModelIndex(unsafe.Pointer(index)))) }) // slot "clearItemData"
}

func (this *QAbstractTableModel) OnMimeTypes(fn func() []string) {
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

func (this *QAbstractTableModel) OnMimeData(fn func([]*QModelIndex) *QMimeData) {
	registerVirt(this.UnsafeGoID(), 5175297091024974749, func(indexes uintptr, indexes_len int64) uintptr { return uintptr(fn(func() []*QModelIndex { if indexes == 0 || indexes_len == 0 { return nil }; __sl := unsafe.Slice((*unsafe.Pointer)(unsafe.Pointer(indexes)), int(indexes_len)); __o := make([]*QModelIndex, int(indexes_len)); for __i := range __o { __o[__i] = UnsafeNewQModelIndex(__sl[__i]) }; return __o }()).UnsafePointer()) }) // slot "mimeData"
}

func (this *QAbstractTableModel) OnCanDropMimeData(fn func(*QMimeData, DropAction, int, int, *QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), 731539533644753052, func(data uintptr, action int32, row int32, column int32, parent uintptr) int32 { return b2i(fn(UnsafeNewQMimeData(unsafe.Pointer(data)), DropAction(action), int(row), int(column), UnsafeNewQModelIndex(unsafe.Pointer(parent)))) }) // slot "canDropMimeData"
}

func (this *QAbstractTableModel) OnSupportedDropActions(fn func() DropAction) {
	registerVirt(this.UnsafeGoID(), -7784642176186821205, func() int32 { return int32(fn()) }) // slot "supportedDropActions"
}

func (this *QAbstractTableModel) OnSupportedDragActions(fn func() DropAction) {
	registerVirt(this.UnsafeGoID(), 3421601386088570762, func() int32 { return int32(fn()) }) // slot "supportedDragActions"
}

func (this *QAbstractTableModel) OnInsertRows(fn func(int, int, *QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), -2775729825962395603, func(row int32, count int32, parent uintptr) int32 { return b2i(fn(int(row), int(count), UnsafeNewQModelIndex(unsafe.Pointer(parent)))) }) // slot "insertRows"
}

func (this *QAbstractTableModel) OnInsertColumns(fn func(int, int, *QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), -8973293846230792237, func(column int32, count int32, parent uintptr) int32 { return b2i(fn(int(column), int(count), UnsafeNewQModelIndex(unsafe.Pointer(parent)))) }) // slot "insertColumns"
}

func (this *QAbstractTableModel) OnRemoveRows(fn func(int, int, *QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), 5104072549182222656, func(row int32, count int32, parent uintptr) int32 { return b2i(fn(int(row), int(count), UnsafeNewQModelIndex(unsafe.Pointer(parent)))) }) // slot "removeRows"
}

func (this *QAbstractTableModel) OnRemoveColumns(fn func(int, int, *QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), 2078548737056540420, func(column int32, count int32, parent uintptr) int32 { return b2i(fn(int(column), int(count), UnsafeNewQModelIndex(unsafe.Pointer(parent)))) }) // slot "removeColumns"
}

func (this *QAbstractTableModel) OnMoveRows(fn func(*QModelIndex, int, int, *QModelIndex, int) bool) {
	registerVirt(this.UnsafeGoID(), -6074819356947405415, func(sourceParent uintptr, sourceRow int32, count int32, destinationParent uintptr, destinationChild int32) int32 { return b2i(fn(UnsafeNewQModelIndex(unsafe.Pointer(sourceParent)), int(sourceRow), int(count), UnsafeNewQModelIndex(unsafe.Pointer(destinationParent)), int(destinationChild))) }) // slot "moveRows"
}

func (this *QAbstractTableModel) OnMoveColumns(fn func(*QModelIndex, int, int, *QModelIndex, int) bool) {
	registerVirt(this.UnsafeGoID(), -6321007591388309905, func(sourceParent uintptr, sourceColumn int32, count int32, destinationParent uintptr, destinationChild int32) int32 { return b2i(fn(UnsafeNewQModelIndex(unsafe.Pointer(sourceParent)), int(sourceColumn), int(count), UnsafeNewQModelIndex(unsafe.Pointer(destinationParent)), int(destinationChild))) }) // slot "moveColumns"
}

func (this *QAbstractTableModel) OnFetchMore(fn func(*QModelIndex)) {
	registerVirt(this.UnsafeGoID(), 3553294593611130806, func(parent uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent))) }) // slot "fetchMore"
}

func (this *QAbstractTableModel) OnCanFetchMore(fn func(*QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), 8639262066473964150, func(parent uintptr) int32 { return b2i(fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)))) }) // slot "canFetchMore"
}

func (this *QAbstractTableModel) OnSort(fn func(int, SortOrder)) {
	registerVirt(this.UnsafeGoID(), 6912441694453382161, func(column int32, order int32) { fn(int(column), SortOrder(order)) }) // slot "sort"
}

func (this *QAbstractTableModel) OnBuddy(fn func(*QModelIndex) *QModelIndex) {
	registerVirt(this.UnsafeGoID(), 6112860112248786499, func(index uintptr) uintptr { return uintptr(fn(UnsafeNewQModelIndex(unsafe.Pointer(index))).UnsafePointer()) }) // slot "buddy"
}

func (this *QAbstractTableModel) OnMatch(fn func(*QModelIndex, int, *QVariant, int, MatchFlag) []*QModelIndex) {
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

func (this *QAbstractTableModel) OnSpan(fn func(*QModelIndex) *QSize) {
	registerVirt(this.UnsafeGoID(), -8395342916587040543, func(index uintptr) uintptr { return uintptr(fn(UnsafeNewQModelIndex(unsafe.Pointer(index))).UnsafePointer()) }) // slot "span"
}

func (this *QAbstractTableModel) OnMultiData(fn func(*QModelIndex, *QModelRoleDataSpan)) {
	registerVirt(this.UnsafeGoID(), 731311060944096642, func(index uintptr, roleDataSpan uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(index)), UnsafeNewQModelRoleDataSpan(unsafe.Pointer(roleDataSpan))) }) // slot "multiData"
}

func (this *QAbstractTableModel) OnSubmit(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -3788331537442685179, func() int32 { return b2i(fn()) }) // slot "submit"
}

func (this *QAbstractTableModel) OnRevert(fn func()) {
	registerVirt(this.UnsafeGoID(), -1522951607536117149, func() { fn() }) // slot "revert"
}

func (this *QAbstractTableModel) OnResetInternalData(fn func()) {
	registerVirt(this.UnsafeGoID(), -4510675770784051183, func() { fn() }) // slot "resetInternalData"
}

func (this *QAbstractTableModel) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QAbstractTableModel) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QAbstractTableModel) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QAbstractTableModel) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QAbstractTableModel) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QAbstractTableModel) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QAbstractTableModel) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QAbstractTableModel) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QAbstractTableModel_delete(this.h)
	}
	this.UnsafeRelease()
}
