package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QStringListModel_chain uintptr
	addr_QStringListModel_new uintptr
	addr_QStringListModel_new2 uintptr
	addr_QStringListModel_new3 uintptr
	addr_QStringListModel_new4 uintptr
	addr_QStringListModel_tr uintptr
	addr_QStringListModel_stringList uintptr
	addr_QStringListModel_setStringList uintptr
	addr_QStringListModel_createIndex uintptr
	addr_QStringListModel_createIndex2 uintptr
	addr_QStringListModel_encodeData uintptr
	addr_QStringListModel_decodeData uintptr
	addr_QStringListModel_beginInsertRows uintptr
	addr_QStringListModel_endInsertRows uintptr
	addr_QStringListModel_beginRemoveRows uintptr
	addr_QStringListModel_endRemoveRows uintptr
	addr_QStringListModel_beginMoveRows uintptr
	addr_QStringListModel_endMoveRows uintptr
	addr_QStringListModel_beginInsertColumns uintptr
	addr_QStringListModel_endInsertColumns uintptr
	addr_QStringListModel_beginRemoveColumns uintptr
	addr_QStringListModel_endRemoveColumns uintptr
	addr_QStringListModel_beginMoveColumns uintptr
	addr_QStringListModel_endMoveColumns uintptr
	addr_QStringListModel_beginResetModel uintptr
	addr_QStringListModel_endResetModel uintptr
	addr_QStringListModel_changePersistentIndex uintptr
	addr_QStringListModel_changePersistentIndexList uintptr
	addr_QStringListModel_persistentIndexList uintptr
	addr_QStringListModel_connect uintptr
	addr_QStringListModel_connect2 uintptr
	addr_QStringListModel_sender uintptr
	addr_QStringListModel_senderSignalIndex uintptr
	addr_QStringListModel_receivers uintptr
	addr_QStringListModel_isSignalConnected uintptr
	addr_QStringListModel_tr2 uintptr
	addr_QStringListModel_tr3 uintptr
	addr_QStringListModel_connect3 uintptr
	addr_QStringListModel_connect4 uintptr
	addr_QStringListModel_metaObject_default uintptr
	addr_QStringListModel_rowCount_default uintptr
	addr_QStringListModel_sibling_default uintptr
	addr_QStringListModel_data_default uintptr
	addr_QStringListModel_setData_default uintptr
	addr_QStringListModel_clearItemData_default uintptr
	addr_QStringListModel_flags_default uintptr
	addr_QStringListModel_insertRows_default uintptr
	addr_QStringListModel_removeRows_default uintptr
	addr_QStringListModel_moveRows_default uintptr
	addr_QStringListModel_sort_default uintptr
	addr_QStringListModel_supportedDropActions_default uintptr
	addr_QStringListModel_index_default uintptr
	addr_QStringListModel_dropMimeData_default uintptr
	addr_QStringListModel_headerData_default uintptr
	addr_QStringListModel_setHeaderData_default uintptr
	addr_QStringListModel_mimeTypes_default uintptr
	addr_QStringListModel_mimeData_default uintptr
	addr_QStringListModel_canDropMimeData_default uintptr
	addr_QStringListModel_supportedDragActions_default uintptr
	addr_QStringListModel_insertColumns_default uintptr
	addr_QStringListModel_removeColumns_default uintptr
	addr_QStringListModel_moveColumns_default uintptr
	addr_QStringListModel_fetchMore_default uintptr
	addr_QStringListModel_canFetchMore_default uintptr
	addr_QStringListModel_buddy_default uintptr
	addr_QStringListModel_match_default uintptr
	addr_QStringListModel_span_default uintptr
	addr_QStringListModel_multiData_default uintptr
	addr_QStringListModel_submit_default uintptr
	addr_QStringListModel_revert_default uintptr
	addr_QStringListModel_resetInternalData_default uintptr
	addr_QStringListModel_event_default uintptr
	addr_QStringListModel_eventFilter_default uintptr
	addr_QStringListModel_timerEvent_default uintptr
	addr_QStringListModel_childEvent_default uintptr
	addr_QStringListModel_customEvent_default uintptr
	addr_QStringListModel_connectNotify_default uintptr
	addr_QStringListModel_disconnectNotify_default uintptr
	pg_QStringListModel_connect_headerDataChanged func(uintptr, int64)
	pg_QStringListModel_connect_layoutChanged func(uintptr, int64)
	pg_QStringListModel_connect_layoutChanged2 func(uintptr, int64)
	pg_QStringListModel_connect_layoutAboutToBeChanged func(uintptr, int64)
	pg_QStringListModel_connect_layoutAboutToBeChanged2 func(uintptr, int64)
	pg_QStringListModel_connect_rowsAboutToBeInserted func(uintptr, int64)
	pg_QStringListModel_connect_rowsInserted func(uintptr, int64)
	pg_QStringListModel_connect_rowsAboutToBeRemoved func(uintptr, int64)
	pg_QStringListModel_connect_rowsRemoved func(uintptr, int64)
	pg_QStringListModel_connect_columnsAboutToBeInserted func(uintptr, int64)
	pg_QStringListModel_connect_columnsInserted func(uintptr, int64)
	pg_QStringListModel_connect_columnsAboutToBeRemoved func(uintptr, int64)
	pg_QStringListModel_connect_columnsRemoved func(uintptr, int64)
	pg_QStringListModel_connect_modelAboutToBeReset func(uintptr, int64)
	pg_QStringListModel_connect_modelReset func(uintptr, int64)
	pg_QStringListModel_connect_rowsAboutToBeMoved func(uintptr, int64)
	pg_QStringListModel_connect_rowsMoved func(uintptr, int64)
	pg_QStringListModel_connect_columnsAboutToBeMoved func(uintptr, int64)
	pg_QStringListModel_connect_columnsMoved func(uintptr, int64)
	pg_QStringListModel_connect_destroyed func(uintptr, int64)
	pg_QStringListModel_connect_destroyed2 func(uintptr, int64)
	pg_QStringListModel_connect_objectNameChanged func(uintptr, int64)
	pg_QStringListModel__disconnect func(int64)
	pg_QStringListModel__set_override func(uintptr, int32)
	pg_QStringListModel_delete func(uintptr)
)

func init() {
	symaddr(&addr_QStringListModel_chain, "QStringListModel_chain")
	symaddr(&addr_QStringListModel_new, "QStringListModel_new")
	symaddr(&addr_QStringListModel_new2, "QStringListModel_new2")
	symaddr(&addr_QStringListModel_new3, "QStringListModel_new3")
	symaddr(&addr_QStringListModel_new4, "QStringListModel_new4")
	symaddr(&addr_QStringListModel_tr, "QStringListModel_tr")
	symaddr(&addr_QStringListModel_stringList, "QStringListModel_stringList")
	symaddr(&addr_QStringListModel_setStringList, "QStringListModel_setStringList")
	symaddr(&addr_QStringListModel_createIndex, "QStringListModel_createIndex")
	symaddr(&addr_QStringListModel_createIndex2, "QStringListModel_createIndex2")
	symaddr(&addr_QStringListModel_encodeData, "QStringListModel_encodeData")
	symaddr(&addr_QStringListModel_decodeData, "QStringListModel_decodeData")
	symaddr(&addr_QStringListModel_beginInsertRows, "QStringListModel_beginInsertRows")
	symaddr(&addr_QStringListModel_endInsertRows, "QStringListModel_endInsertRows")
	symaddr(&addr_QStringListModel_beginRemoveRows, "QStringListModel_beginRemoveRows")
	symaddr(&addr_QStringListModel_endRemoveRows, "QStringListModel_endRemoveRows")
	symaddr(&addr_QStringListModel_beginMoveRows, "QStringListModel_beginMoveRows")
	symaddr(&addr_QStringListModel_endMoveRows, "QStringListModel_endMoveRows")
	symaddr(&addr_QStringListModel_beginInsertColumns, "QStringListModel_beginInsertColumns")
	symaddr(&addr_QStringListModel_endInsertColumns, "QStringListModel_endInsertColumns")
	symaddr(&addr_QStringListModel_beginRemoveColumns, "QStringListModel_beginRemoveColumns")
	symaddr(&addr_QStringListModel_endRemoveColumns, "QStringListModel_endRemoveColumns")
	symaddr(&addr_QStringListModel_beginMoveColumns, "QStringListModel_beginMoveColumns")
	symaddr(&addr_QStringListModel_endMoveColumns, "QStringListModel_endMoveColumns")
	symaddr(&addr_QStringListModel_beginResetModel, "QStringListModel_beginResetModel")
	symaddr(&addr_QStringListModel_endResetModel, "QStringListModel_endResetModel")
	symaddr(&addr_QStringListModel_changePersistentIndex, "QStringListModel_changePersistentIndex")
	symaddr(&addr_QStringListModel_changePersistentIndexList, "QStringListModel_changePersistentIndexList")
	symaddr(&addr_QStringListModel_persistentIndexList, "QStringListModel_persistentIndexList")
	symaddr(&addr_QStringListModel_connect, "QStringListModel_connect")
	symaddr(&addr_QStringListModel_connect2, "QStringListModel_connect2")
	symaddr(&addr_QStringListModel_sender, "QStringListModel_sender")
	symaddr(&addr_QStringListModel_senderSignalIndex, "QStringListModel_senderSignalIndex")
	symaddr(&addr_QStringListModel_receivers, "QStringListModel_receivers")
	symaddr(&addr_QStringListModel_isSignalConnected, "QStringListModel_isSignalConnected")
	symaddr(&addr_QStringListModel_tr2, "QStringListModel_tr2")
	symaddr(&addr_QStringListModel_tr3, "QStringListModel_tr3")
	symaddr(&addr_QStringListModel_connect3, "QStringListModel_connect3")
	symaddr(&addr_QStringListModel_connect4, "QStringListModel_connect4")
	symaddr(&addr_QStringListModel_metaObject_default, "QStringListModel_metaObject_default")
	symaddr(&addr_QStringListModel_rowCount_default, "QStringListModel_rowCount_default")
	symaddr(&addr_QStringListModel_sibling_default, "QStringListModel_sibling_default")
	symaddr(&addr_QStringListModel_data_default, "QStringListModel_data_default")
	symaddr(&addr_QStringListModel_setData_default, "QStringListModel_setData_default")
	symaddr(&addr_QStringListModel_clearItemData_default, "QStringListModel_clearItemData_default")
	symaddr(&addr_QStringListModel_flags_default, "QStringListModel_flags_default")
	symaddr(&addr_QStringListModel_insertRows_default, "QStringListModel_insertRows_default")
	symaddr(&addr_QStringListModel_removeRows_default, "QStringListModel_removeRows_default")
	symaddr(&addr_QStringListModel_moveRows_default, "QStringListModel_moveRows_default")
	symaddr(&addr_QStringListModel_sort_default, "QStringListModel_sort_default")
	symaddr(&addr_QStringListModel_supportedDropActions_default, "QStringListModel_supportedDropActions_default")
	symaddr(&addr_QStringListModel_index_default, "QStringListModel_index_default")
	symaddr(&addr_QStringListModel_dropMimeData_default, "QStringListModel_dropMimeData_default")
	symaddr(&addr_QStringListModel_headerData_default, "QStringListModel_headerData_default")
	symaddr(&addr_QStringListModel_setHeaderData_default, "QStringListModel_setHeaderData_default")
	symaddr(&addr_QStringListModel_mimeTypes_default, "QStringListModel_mimeTypes_default")
	symaddr(&addr_QStringListModel_mimeData_default, "QStringListModel_mimeData_default")
	symaddr(&addr_QStringListModel_canDropMimeData_default, "QStringListModel_canDropMimeData_default")
	symaddr(&addr_QStringListModel_supportedDragActions_default, "QStringListModel_supportedDragActions_default")
	symaddr(&addr_QStringListModel_insertColumns_default, "QStringListModel_insertColumns_default")
	symaddr(&addr_QStringListModel_removeColumns_default, "QStringListModel_removeColumns_default")
	symaddr(&addr_QStringListModel_moveColumns_default, "QStringListModel_moveColumns_default")
	symaddr(&addr_QStringListModel_fetchMore_default, "QStringListModel_fetchMore_default")
	symaddr(&addr_QStringListModel_canFetchMore_default, "QStringListModel_canFetchMore_default")
	symaddr(&addr_QStringListModel_buddy_default, "QStringListModel_buddy_default")
	symaddr(&addr_QStringListModel_match_default, "QStringListModel_match_default")
	symaddr(&addr_QStringListModel_span_default, "QStringListModel_span_default")
	symaddr(&addr_QStringListModel_multiData_default, "QStringListModel_multiData_default")
	symaddr(&addr_QStringListModel_submit_default, "QStringListModel_submit_default")
	symaddr(&addr_QStringListModel_revert_default, "QStringListModel_revert_default")
	symaddr(&addr_QStringListModel_resetInternalData_default, "QStringListModel_resetInternalData_default")
	symaddr(&addr_QStringListModel_event_default, "QStringListModel_event_default")
	symaddr(&addr_QStringListModel_eventFilter_default, "QStringListModel_eventFilter_default")
	symaddr(&addr_QStringListModel_timerEvent_default, "QStringListModel_timerEvent_default")
	symaddr(&addr_QStringListModel_childEvent_default, "QStringListModel_childEvent_default")
	symaddr(&addr_QStringListModel_customEvent_default, "QStringListModel_customEvent_default")
	symaddr(&addr_QStringListModel_connectNotify_default, "QStringListModel_connectNotify_default")
	symaddr(&addr_QStringListModel_disconnectNotify_default, "QStringListModel_disconnectNotify_default")
	reg(&pg_QStringListModel_connect_headerDataChanged, "QStringListModel_connect_headerDataChanged")
	reg(&pg_QStringListModel_connect_layoutChanged, "QStringListModel_connect_layoutChanged")
	reg(&pg_QStringListModel_connect_layoutChanged2, "QStringListModel_connect_layoutChanged2")
	reg(&pg_QStringListModel_connect_layoutAboutToBeChanged, "QStringListModel_connect_layoutAboutToBeChanged")
	reg(&pg_QStringListModel_connect_layoutAboutToBeChanged2, "QStringListModel_connect_layoutAboutToBeChanged2")
	reg(&pg_QStringListModel_connect_rowsAboutToBeInserted, "QStringListModel_connect_rowsAboutToBeInserted")
	reg(&pg_QStringListModel_connect_rowsInserted, "QStringListModel_connect_rowsInserted")
	reg(&pg_QStringListModel_connect_rowsAboutToBeRemoved, "QStringListModel_connect_rowsAboutToBeRemoved")
	reg(&pg_QStringListModel_connect_rowsRemoved, "QStringListModel_connect_rowsRemoved")
	reg(&pg_QStringListModel_connect_columnsAboutToBeInserted, "QStringListModel_connect_columnsAboutToBeInserted")
	reg(&pg_QStringListModel_connect_columnsInserted, "QStringListModel_connect_columnsInserted")
	reg(&pg_QStringListModel_connect_columnsAboutToBeRemoved, "QStringListModel_connect_columnsAboutToBeRemoved")
	reg(&pg_QStringListModel_connect_columnsRemoved, "QStringListModel_connect_columnsRemoved")
	reg(&pg_QStringListModel_connect_modelAboutToBeReset, "QStringListModel_connect_modelAboutToBeReset")
	reg(&pg_QStringListModel_connect_modelReset, "QStringListModel_connect_modelReset")
	reg(&pg_QStringListModel_connect_rowsAboutToBeMoved, "QStringListModel_connect_rowsAboutToBeMoved")
	reg(&pg_QStringListModel_connect_rowsMoved, "QStringListModel_connect_rowsMoved")
	reg(&pg_QStringListModel_connect_columnsAboutToBeMoved, "QStringListModel_connect_columnsAboutToBeMoved")
	reg(&pg_QStringListModel_connect_columnsMoved, "QStringListModel_connect_columnsMoved")
	reg(&pg_QStringListModel_connect_destroyed, "QStringListModel_connect_destroyed")
	reg(&pg_QStringListModel_connect_destroyed2, "QStringListModel_connect_destroyed2")
	reg(&pg_QStringListModel_connect_objectNameChanged, "QStringListModel_connect_objectNameChanged")
	reg(&pg_QStringListModel__disconnect, "QStringListModel__disconnect")
	reg(&pg_QStringListModel__set_override, "QStringListModel__set_override")
	reg(&pg_QStringListModel_delete, "QStringListModel_delete")
}

type QStringListModel struct {
	h uintptr
	*QAbstractListModel
}

func bindQStringListModel(h uintptr) *QStringListModel {
	if h == 0 {
		return nil
	}
	var __c [3]uintptr
	purego.SyscallN(addr_QStringListModel_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QStringListModel{h: h, QAbstractListModel: &QAbstractListModel{h: __c[0], QAbstractItemModel: &QAbstractItemModel{h: __c[1], QObject: &QObject{h: __c[2]}}}}
}

func newownedQStringListModel(h uintptr) *QStringListModel { return bindQStringListModel(h) }

func (this *QStringListModel) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QStringListModel) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQStringListModel(h unsafe.Pointer) *QStringListModel { return bindQStringListModel(uintptr(h)) }
func UnsafeNewQStringListModelOwned(h unsafe.Pointer) *QStringListModel { return newownedQStringListModel(uintptr(h)) }

func (this *QStringListModel) UnsafeRelease() {
	this.h = 0
	if this.QAbstractListModel != nil {
		this.QAbstractListModel.UnsafeRelease()
	}
}

func NewQStringListModel() *QStringListModel {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QStringListModel_new, uintptr(__goid))
	__o := newownedQStringListModel(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QStringListModel_virtIdx, QStringListModel_enable)
	return __o
}

func NewQStringListModel2(strings []string) *QStringListModel {
	__goid := nextVirtID()
	strings_arr := make([]uintptr, len(strings))
	for __i, __e := range strings {
		strings_arr[__i] = pgCString(__e)
	}
	var strings_p uintptr
	if len(strings_arr) > 0 {
		strings_p = uintptr(unsafe.Pointer(&strings_arr[0]))
	}
	__r, _, _ := purego.SyscallN(addr_QStringListModel_new2, uintptr(__goid), strings_p, uintptr(len(strings)))
	__o := newownedQStringListModel(__r)
	for __i := range strings_arr {
		pgFree(strings_arr[__i])
	}
	runtime.KeepAlive(strings_arr)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QStringListModel_virtIdx, QStringListModel_enable)
	return __o
}

func NewQStringListModel3(parent *QObject) *QStringListModel {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QStringListModel_new3, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQStringListModel(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QStringListModel_virtIdx, QStringListModel_enable)
	return __o
}

func NewQStringListModel4(strings []string, parent *QObject) *QStringListModel {
	__goid := nextVirtID()
	strings_arr := make([]uintptr, len(strings))
	for __i, __e := range strings {
		strings_arr[__i] = pgCString(__e)
	}
	var strings_p uintptr
	if len(strings_arr) > 0 {
		strings_p = uintptr(unsafe.Pointer(&strings_arr[0]))
	}
	__r, _, _ := purego.SyscallN(addr_QStringListModel_new4, uintptr(__goid), strings_p, uintptr(len(strings)), uintptr(parent.UnsafePointer()))
	__o := newownedQStringListModel(__r)
	for __i := range strings_arr {
		pgFree(strings_arr[__i])
	}
	runtime.KeepAlive(strings_arr)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QStringListModel_virtIdx, QStringListModel_enable)
	return __o
}

func QStringListModel_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStringListModel_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QStringListModel) StringList() []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStringListModel_stringList, this.h, uintptr(unsafe.Pointer(&__n)))
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

func (this *QStringListModel) SetStringList(strings []string) {
	strings_arr := make([]uintptr, len(strings))
	for __i, __e := range strings {
		strings_arr[__i] = pgCString(__e)
	}
	var strings_p uintptr
	if len(strings_arr) > 0 {
		strings_p = uintptr(unsafe.Pointer(&strings_arr[0]))
	}
	purego.SyscallN(addr_QStringListModel_setStringList, this.h, strings_p, uintptr(len(strings)))
	runtime.KeepAlive(this)
	for __i := range strings_arr {
		pgFree(strings_arr[__i])
	}
	runtime.KeepAlive(strings_arr)
}

func (this *QStringListModel) CreateIndex(row int, column int) *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QStringListModel_createIndex, this.h, uintptr(row), uintptr(column))
	runtime.KeepAlive(this)
	return newownedQModelIndex(__r)
}

func (this *QStringListModel) CreateIndex2(row int, column int, id uintptr) *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QStringListModel_createIndex2, this.h, uintptr(row), uintptr(column), uintptr(id))
	runtime.KeepAlive(this)
	return newownedQModelIndex(__r)
}

func (this *QStringListModel) EncodeData(indexes []*QModelIndex, stream *QDataStream) {
	indexes_arr := make([]uintptr, len(indexes))
	for __i, __e := range indexes {
		indexes_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var indexes_p uintptr
	if len(indexes_arr) > 0 {
		indexes_p = uintptr(unsafe.Pointer(&indexes_arr[0]))
	}
	purego.SyscallN(addr_QStringListModel_encodeData, this.h, indexes_p, uintptr(len(indexes)), uintptr(stream.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(indexes)
	runtime.KeepAlive(indexes_arr)
	runtime.KeepAlive(stream)
}

func (this *QStringListModel) DecodeData(row int, column int, parent *QModelIndex, stream *QDataStream) bool {
	__r0, _, _ := purego.SyscallN(addr_QStringListModel_decodeData, this.h, uintptr(row), uintptr(column), uintptr(parent.UnsafePointer()), uintptr(stream.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	runtime.KeepAlive(stream)
	return __r != 0
}

func (this *QStringListModel) BeginInsertRows(parent *QModelIndex, first int, last int) {
	purego.SyscallN(addr_QStringListModel_beginInsertRows, this.h, uintptr(parent.UnsafePointer()), uintptr(first), uintptr(last))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
}

func (this *QStringListModel) EndInsertRows() {
	purego.SyscallN(addr_QStringListModel_endInsertRows, this.h)
	runtime.KeepAlive(this)
}

func (this *QStringListModel) BeginRemoveRows(parent *QModelIndex, first int, last int) {
	purego.SyscallN(addr_QStringListModel_beginRemoveRows, this.h, uintptr(parent.UnsafePointer()), uintptr(first), uintptr(last))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
}

func (this *QStringListModel) EndRemoveRows() {
	purego.SyscallN(addr_QStringListModel_endRemoveRows, this.h)
	runtime.KeepAlive(this)
}

func (this *QStringListModel) BeginMoveRows(sourceParent *QModelIndex, sourceFirst int, sourceLast int, destinationParent *QModelIndex, destinationRow int) bool {
	__r0, _, _ := purego.SyscallN(addr_QStringListModel_beginMoveRows, this.h, uintptr(sourceParent.UnsafePointer()), uintptr(sourceFirst), uintptr(sourceLast), uintptr(destinationParent.UnsafePointer()), uintptr(destinationRow))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sourceParent)
	runtime.KeepAlive(destinationParent)
	return __r != 0
}

func (this *QStringListModel) EndMoveRows() {
	purego.SyscallN(addr_QStringListModel_endMoveRows, this.h)
	runtime.KeepAlive(this)
}

func (this *QStringListModel) BeginInsertColumns(parent *QModelIndex, first int, last int) {
	purego.SyscallN(addr_QStringListModel_beginInsertColumns, this.h, uintptr(parent.UnsafePointer()), uintptr(first), uintptr(last))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
}

func (this *QStringListModel) EndInsertColumns() {
	purego.SyscallN(addr_QStringListModel_endInsertColumns, this.h)
	runtime.KeepAlive(this)
}

func (this *QStringListModel) BeginRemoveColumns(parent *QModelIndex, first int, last int) {
	purego.SyscallN(addr_QStringListModel_beginRemoveColumns, this.h, uintptr(parent.UnsafePointer()), uintptr(first), uintptr(last))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
}

func (this *QStringListModel) EndRemoveColumns() {
	purego.SyscallN(addr_QStringListModel_endRemoveColumns, this.h)
	runtime.KeepAlive(this)
}

func (this *QStringListModel) BeginMoveColumns(sourceParent *QModelIndex, sourceFirst int, sourceLast int, destinationParent *QModelIndex, destinationColumn int) bool {
	__r0, _, _ := purego.SyscallN(addr_QStringListModel_beginMoveColumns, this.h, uintptr(sourceParent.UnsafePointer()), uintptr(sourceFirst), uintptr(sourceLast), uintptr(destinationParent.UnsafePointer()), uintptr(destinationColumn))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sourceParent)
	runtime.KeepAlive(destinationParent)
	return __r != 0
}

func (this *QStringListModel) EndMoveColumns() {
	purego.SyscallN(addr_QStringListModel_endMoveColumns, this.h)
	runtime.KeepAlive(this)
}

func (this *QStringListModel) BeginResetModel() {
	purego.SyscallN(addr_QStringListModel_beginResetModel, this.h)
	runtime.KeepAlive(this)
}

func (this *QStringListModel) EndResetModel() {
	purego.SyscallN(addr_QStringListModel_endResetModel, this.h)
	runtime.KeepAlive(this)
}

func (this *QStringListModel) ChangePersistentIndex(from *QModelIndex, to *QModelIndex) {
	purego.SyscallN(addr_QStringListModel_changePersistentIndex, this.h, uintptr(from.UnsafePointer()), uintptr(to.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(from)
	runtime.KeepAlive(to)
}

func (this *QStringListModel) ChangePersistentIndexList(from []*QModelIndex, to []*QModelIndex) {
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
	purego.SyscallN(addr_QStringListModel_changePersistentIndexList, this.h, from_p, uintptr(len(from)), to_p, uintptr(len(to)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(from)
	runtime.KeepAlive(from_arr)
	runtime.KeepAlive(to)
	runtime.KeepAlive(to_arr)
}

func (this *QStringListModel) PersistentIndexList() []*QModelIndex {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStringListModel_persistentIndexList, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QModelIndex, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQModelIndex(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QStringListModel_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QStringListModel_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QStringListModel_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QStringListModel_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QStringListModel) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QStringListModel_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QStringListModel) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QStringListModel_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QStringListModel) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QStringListModel_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QStringListModel) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QStringListModel_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QStringListModel_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStringListModel_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QStringListModel_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStringListModel_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QStringListModel_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QStringListModel_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QStringListModel_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QStringListModel_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QStringListModel) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QStringListModel_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QStringListModel) RowCountDefault(parent *QModelIndex) int {
	__r0, _, _ := purego.SyscallN(addr_QStringListModel_rowCount_default, this.h, uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return int(__r)
}

func (this *QStringListModel) SiblingDefault(row int, column int, idx *QModelIndex) *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QStringListModel_sibling_default, this.h, uintptr(row), uintptr(column), uintptr(idx.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(idx)
	return newownedQModelIndex(__r)
}

func (this *QStringListModel) DataDefault(index *QModelIndex, role int) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QStringListModel_data_default, this.h, uintptr(index.UnsafePointer()), uintptr(role))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return newownedQVariant(__r)
}

func (this *QStringListModel) SetDataDefault(index *QModelIndex, value *QVariant, role int) bool {
	__r0, _, _ := purego.SyscallN(addr_QStringListModel_setData_default, this.h, uintptr(index.UnsafePointer()), uintptr(value.UnsafePointer()), uintptr(role))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	runtime.KeepAlive(value)
	return __r != 0
}

func (this *QStringListModel) ClearItemDataDefault(index *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QStringListModel_clearItemData_default, this.h, uintptr(index.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return __r != 0
}

func (this *QStringListModel) FlagsDefault(index *QModelIndex) ItemFlag {
	__r0, _, _ := purego.SyscallN(addr_QStringListModel_flags_default, this.h, uintptr(index.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return ItemFlag(__r)
}

func (this *QStringListModel) InsertRowsDefault(row int, count int, parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QStringListModel_insertRows_default, this.h, uintptr(row), uintptr(count), uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QStringListModel) RemoveRowsDefault(row int, count int, parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QStringListModel_removeRows_default, this.h, uintptr(row), uintptr(count), uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QStringListModel) MoveRowsDefault(sourceParent *QModelIndex, sourceRow int, count int, destinationParent *QModelIndex, destinationChild int) bool {
	__r0, _, _ := purego.SyscallN(addr_QStringListModel_moveRows_default, this.h, uintptr(sourceParent.UnsafePointer()), uintptr(sourceRow), uintptr(count), uintptr(destinationParent.UnsafePointer()), uintptr(destinationChild))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sourceParent)
	runtime.KeepAlive(destinationParent)
	return __r != 0
}

func (this *QStringListModel) SortDefault(column int, order SortOrder) {
	purego.SyscallN(addr_QStringListModel_sort_default, this.h, uintptr(column), uintptr(order))
	runtime.KeepAlive(this)
}

func (this *QStringListModel) SupportedDropActionsDefault() DropAction {
	__r0, _, _ := purego.SyscallN(addr_QStringListModel_supportedDropActions_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return DropAction(__r)
}

func (this *QStringListModel) IndexDefault(row int, column int, parent *QModelIndex) *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QStringListModel_index_default, this.h, uintptr(row), uintptr(column), uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return newownedQModelIndex(__r)
}

func (this *QStringListModel) DropMimeDataDefault(data *QMimeData, action DropAction, row int, column int, parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QStringListModel_dropMimeData_default, this.h, uintptr(data.UnsafePointer()), uintptr(action), uintptr(row), uintptr(column), uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(data)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QStringListModel) HeaderDataDefault(section int, orientation Orientation, role int) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QStringListModel_headerData_default, this.h, uintptr(section), uintptr(orientation), uintptr(role))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QStringListModel) SetHeaderDataDefault(section int, orientation Orientation, value *QVariant, role int) bool {
	__r0, _, _ := purego.SyscallN(addr_QStringListModel_setHeaderData_default, this.h, uintptr(section), uintptr(orientation), uintptr(value.UnsafePointer()), uintptr(role))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(value)
	return __r != 0
}

func (this *QStringListModel) MimeTypesDefault() []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStringListModel_mimeTypes_default, this.h, uintptr(unsafe.Pointer(&__n)))
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

func (this *QStringListModel) MimeDataDefault(indexes []*QModelIndex) *QMimeData {
	indexes_arr := make([]uintptr, len(indexes))
	for __i, __e := range indexes {
		indexes_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var indexes_p uintptr
	if len(indexes_arr) > 0 {
		indexes_p = uintptr(unsafe.Pointer(&indexes_arr[0]))
	}
	__r, _, _ := purego.SyscallN(addr_QStringListModel_mimeData_default, this.h, indexes_p, uintptr(len(indexes)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(indexes)
	runtime.KeepAlive(indexes_arr)
	return bindQMimeData(__r)
}

func (this *QStringListModel) CanDropMimeDataDefault(data *QMimeData, action DropAction, row int, column int, parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QStringListModel_canDropMimeData_default, this.h, uintptr(data.UnsafePointer()), uintptr(action), uintptr(row), uintptr(column), uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(data)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QStringListModel) SupportedDragActionsDefault() DropAction {
	__r0, _, _ := purego.SyscallN(addr_QStringListModel_supportedDragActions_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return DropAction(__r)
}

func (this *QStringListModel) InsertColumnsDefault(column int, count int, parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QStringListModel_insertColumns_default, this.h, uintptr(column), uintptr(count), uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QStringListModel) RemoveColumnsDefault(column int, count int, parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QStringListModel_removeColumns_default, this.h, uintptr(column), uintptr(count), uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QStringListModel) MoveColumnsDefault(sourceParent *QModelIndex, sourceColumn int, count int, destinationParent *QModelIndex, destinationChild int) bool {
	__r0, _, _ := purego.SyscallN(addr_QStringListModel_moveColumns_default, this.h, uintptr(sourceParent.UnsafePointer()), uintptr(sourceColumn), uintptr(count), uintptr(destinationParent.UnsafePointer()), uintptr(destinationChild))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sourceParent)
	runtime.KeepAlive(destinationParent)
	return __r != 0
}

func (this *QStringListModel) FetchMoreDefault(parent *QModelIndex) {
	purego.SyscallN(addr_QStringListModel_fetchMore_default, this.h, uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
}

func (this *QStringListModel) CanFetchMoreDefault(parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QStringListModel_canFetchMore_default, this.h, uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QStringListModel) BuddyDefault(index *QModelIndex) *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QStringListModel_buddy_default, this.h, uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return newownedQModelIndex(__r)
}

func (this *QStringListModel) MatchDefault(start *QModelIndex, role int, value *QVariant, hits int, flags MatchFlag) []*QModelIndex {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QStringListModel_match_default, this.h, uintptr(start.UnsafePointer()), uintptr(role), uintptr(value.UnsafePointer()), uintptr(hits), uintptr(flags), uintptr(unsafe.Pointer(&__n)))
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

func (this *QStringListModel) SpanDefault(index *QModelIndex) *QSize {
	__r, _, _ := purego.SyscallN(addr_QStringListModel_span_default, this.h, uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return newownedQSize(__r)
}

func (this *QStringListModel) MultiDataDefault(index *QModelIndex, roleDataSpan *QModelRoleDataSpan) {
	purego.SyscallN(addr_QStringListModel_multiData_default, this.h, uintptr(index.UnsafePointer()), uintptr(roleDataSpan.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	runtime.KeepAlive(roleDataSpan)
}

func (this *QStringListModel) SubmitDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QStringListModel_submit_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QStringListModel) RevertDefault() {
	purego.SyscallN(addr_QStringListModel_revert_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QStringListModel) ResetInternalDataDefault() {
	purego.SyscallN(addr_QStringListModel_resetInternalData_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QStringListModel) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QStringListModel_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QStringListModel) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QStringListModel_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QStringListModel) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QStringListModel_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QStringListModel) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QStringListModel_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QStringListModel) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QStringListModel_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QStringListModel) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QStringListModel_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QStringListModel) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QStringListModel_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QStringListModel) OnHeaderDataChanged(fn func(Orientation, int, int)) int64 {
	id := registerCallback(func(orientation int32, first int32, last int32) { fn(Orientation(orientation), int(first), int(last)) })
	pg_QStringListModel_connect_headerDataChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStringListModel) OnLayoutChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QStringListModel_connect_layoutChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStringListModel) OnLayoutChanged2(fn func([]*QPersistentModelIndex, QAbstractItemModel__LayoutChangeHint)) int64 {
	id := registerCallback(func(parents uintptr, parents_len int64, hint int32) { fn(func() []*QPersistentModelIndex { if parents == 0 || parents_len == 0 { return nil }; __sl := unsafe.Slice((*unsafe.Pointer)(unsafe.Pointer(parents)), int(parents_len)); __o := make([]*QPersistentModelIndex, int(parents_len)); for __i := range __o { __o[__i] = UnsafeNewQPersistentModelIndex(__sl[__i]) }; return __o }(), QAbstractItemModel__LayoutChangeHint(hint)) })
	pg_QStringListModel_connect_layoutChanged2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStringListModel) OnLayoutAboutToBeChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QStringListModel_connect_layoutAboutToBeChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStringListModel) OnLayoutAboutToBeChanged2(fn func([]*QPersistentModelIndex, QAbstractItemModel__LayoutChangeHint)) int64 {
	id := registerCallback(func(parents uintptr, parents_len int64, hint int32) { fn(func() []*QPersistentModelIndex { if parents == 0 || parents_len == 0 { return nil }; __sl := unsafe.Slice((*unsafe.Pointer)(unsafe.Pointer(parents)), int(parents_len)); __o := make([]*QPersistentModelIndex, int(parents_len)); for __i := range __o { __o[__i] = UnsafeNewQPersistentModelIndex(__sl[__i]) }; return __o }(), QAbstractItemModel__LayoutChangeHint(hint)) })
	pg_QStringListModel_connect_layoutAboutToBeChanged2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStringListModel) OnRowsAboutToBeInserted(fn func(*QModelIndex, int, int)) int64 {
	id := registerCallback(func(parent uintptr, first int32, last int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(first), int(last)) })
	pg_QStringListModel_connect_rowsAboutToBeInserted(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStringListModel) OnRowsInserted(fn func(*QModelIndex, int, int)) int64 {
	id := registerCallback(func(parent uintptr, first int32, last int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(first), int(last)) })
	pg_QStringListModel_connect_rowsInserted(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStringListModel) OnRowsAboutToBeRemoved(fn func(*QModelIndex, int, int)) int64 {
	id := registerCallback(func(parent uintptr, first int32, last int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(first), int(last)) })
	pg_QStringListModel_connect_rowsAboutToBeRemoved(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStringListModel) OnRowsRemoved(fn func(*QModelIndex, int, int)) int64 {
	id := registerCallback(func(parent uintptr, first int32, last int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(first), int(last)) })
	pg_QStringListModel_connect_rowsRemoved(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStringListModel) OnColumnsAboutToBeInserted(fn func(*QModelIndex, int, int)) int64 {
	id := registerCallback(func(parent uintptr, first int32, last int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(first), int(last)) })
	pg_QStringListModel_connect_columnsAboutToBeInserted(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStringListModel) OnColumnsInserted(fn func(*QModelIndex, int, int)) int64 {
	id := registerCallback(func(parent uintptr, first int32, last int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(first), int(last)) })
	pg_QStringListModel_connect_columnsInserted(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStringListModel) OnColumnsAboutToBeRemoved(fn func(*QModelIndex, int, int)) int64 {
	id := registerCallback(func(parent uintptr, first int32, last int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(first), int(last)) })
	pg_QStringListModel_connect_columnsAboutToBeRemoved(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStringListModel) OnColumnsRemoved(fn func(*QModelIndex, int, int)) int64 {
	id := registerCallback(func(parent uintptr, first int32, last int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(first), int(last)) })
	pg_QStringListModel_connect_columnsRemoved(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStringListModel) OnModelAboutToBeReset(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QStringListModel_connect_modelAboutToBeReset(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStringListModel) OnModelReset(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QStringListModel_connect_modelReset(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStringListModel) OnRowsAboutToBeMoved(fn func(*QModelIndex, int, int, *QModelIndex, int)) int64 {
	id := registerCallback(func(sourceParent uintptr, sourceStart int32, sourceEnd int32, destinationParent uintptr, destinationRow int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(sourceParent)), int(sourceStart), int(sourceEnd), UnsafeNewQModelIndex(unsafe.Pointer(destinationParent)), int(destinationRow)) })
	pg_QStringListModel_connect_rowsAboutToBeMoved(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStringListModel) OnRowsMoved(fn func(*QModelIndex, int, int, *QModelIndex, int)) int64 {
	id := registerCallback(func(sourceParent uintptr, sourceStart int32, sourceEnd int32, destinationParent uintptr, destinationRow int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(sourceParent)), int(sourceStart), int(sourceEnd), UnsafeNewQModelIndex(unsafe.Pointer(destinationParent)), int(destinationRow)) })
	pg_QStringListModel_connect_rowsMoved(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStringListModel) OnColumnsAboutToBeMoved(fn func(*QModelIndex, int, int, *QModelIndex, int)) int64 {
	id := registerCallback(func(sourceParent uintptr, sourceStart int32, sourceEnd int32, destinationParent uintptr, destinationColumn int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(sourceParent)), int(sourceStart), int(sourceEnd), UnsafeNewQModelIndex(unsafe.Pointer(destinationParent)), int(destinationColumn)) })
	pg_QStringListModel_connect_columnsAboutToBeMoved(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStringListModel) OnColumnsMoved(fn func(*QModelIndex, int, int, *QModelIndex, int)) int64 {
	id := registerCallback(func(sourceParent uintptr, sourceStart int32, sourceEnd int32, destinationParent uintptr, destinationColumn int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(sourceParent)), int(sourceStart), int(sourceEnd), UnsafeNewQModelIndex(unsafe.Pointer(destinationParent)), int(destinationColumn)) })
	pg_QStringListModel_connect_columnsMoved(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStringListModel) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QStringListModel_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStringListModel) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QStringListModel_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStringListModel) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QStringListModel_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QStringListModel) Disconnect(handle int64) {
	pg_QStringListModel__disconnect(handle)
	unregisterCallback(handle)
}

func QStringListModel_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case 1960960171956884010:
		return 1
	case 397329346813782521:
		return 2
	case -8837375942098466299:
		return 3
	case 35191454575202219:
		return 4
	case 8914190500520542409:
		return 5
	case 1703382833545304812:
		return 6
	case -2775729825962395603:
		return 7
	case 5104072549182222656:
		return 8
	case -6074819356947405415:
		return 9
	case 6912441694453382161:
		return 10
	case -7784642176186821205:
		return 11
	case -8948777187306027381:
		return 12
	case -9049994684305682212:
		return 13
	case -3297140021896693688:
		return 14
	case 4852200306590054850:
		return 15
	case 5464467312394730418:
		return 16
	case 5175297091024974749:
		return 17
	case 731539533644753052:
		return 18
	case 3421601386088570762:
		return 19
	case -8973293846230792237:
		return 20
	case 2078548737056540420:
		return 21
	case -6321007591388309905:
		return 22
	case 3553294593611130806:
		return 23
	case 8639262066473964150:
		return 24
	case 6112860112248786499:
		return 25
	case -4341501217981328394:
		return 26
	case -8395342916587040543:
		return 27
	case 731311060944096642:
		return 28
	case -3788331537442685179:
		return 29
	case -1522951607536117149:
		return 30
	case -4510675770784051183:
		return 31
	case 1237236456312824991:
		return 32
	case 5217623514630858297:
		return 33
	case -5409927581812374498:
		return 34
	case -1185315367526857833:
		return 35
	case 3769868229123868474:
		return 36
	case 7379029982170503008:
		return 37
	case -568369772982215962:
		return 38
	}
	return -1
}

func QStringListModel_enable(h uintptr, idx int32) { pg_QStringListModel__set_override(h, idx) }

func (this *QStringListModel) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QStringListModel) OnRowCount(fn func(*QModelIndex) int) {
	registerVirt(this.UnsafeGoID(), 1960960171956884010, func(parent uintptr) int32 { return int32(fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)))) }) // slot "rowCount"
}

func (this *QStringListModel) OnSibling(fn func(int, int, *QModelIndex) *QModelIndex) {
	registerVirt(this.UnsafeGoID(), 397329346813782521, func(row int32, column int32, idx uintptr) uintptr { return uintptr(fn(int(row), int(column), UnsafeNewQModelIndex(unsafe.Pointer(idx))).UnsafePointer()) }) // slot "sibling"
}

func (this *QStringListModel) OnData(fn func(*QModelIndex, int) *QVariant) {
	registerVirt(this.UnsafeGoID(), -8837375942098466299, func(index uintptr, role int32) uintptr { return uintptr(fn(UnsafeNewQModelIndex(unsafe.Pointer(index)), int(role)).UnsafePointer()) }) // slot "data"
}

func (this *QStringListModel) OnSetData(fn func(*QModelIndex, *QVariant, int) bool) {
	registerVirt(this.UnsafeGoID(), 35191454575202219, func(index uintptr, value uintptr, role int32) int32 { return b2i(fn(UnsafeNewQModelIndex(unsafe.Pointer(index)), UnsafeNewQVariant(unsafe.Pointer(value)), int(role))) }) // slot "setData"
}

func (this *QStringListModel) OnClearItemData(fn func(*QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), 8914190500520542409, func(index uintptr) int32 { return b2i(fn(UnsafeNewQModelIndex(unsafe.Pointer(index)))) }) // slot "clearItemData"
}

func (this *QStringListModel) OnFlags(fn func(*QModelIndex) ItemFlag) {
	registerVirt(this.UnsafeGoID(), 1703382833545304812, func(index uintptr) int32 { return int32(fn(UnsafeNewQModelIndex(unsafe.Pointer(index)))) }) // slot "flags"
}

func (this *QStringListModel) OnInsertRows(fn func(int, int, *QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), -2775729825962395603, func(row int32, count int32, parent uintptr) int32 { return b2i(fn(int(row), int(count), UnsafeNewQModelIndex(unsafe.Pointer(parent)))) }) // slot "insertRows"
}

func (this *QStringListModel) OnRemoveRows(fn func(int, int, *QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), 5104072549182222656, func(row int32, count int32, parent uintptr) int32 { return b2i(fn(int(row), int(count), UnsafeNewQModelIndex(unsafe.Pointer(parent)))) }) // slot "removeRows"
}

func (this *QStringListModel) OnMoveRows(fn func(*QModelIndex, int, int, *QModelIndex, int) bool) {
	registerVirt(this.UnsafeGoID(), -6074819356947405415, func(sourceParent uintptr, sourceRow int32, count int32, destinationParent uintptr, destinationChild int32) int32 { return b2i(fn(UnsafeNewQModelIndex(unsafe.Pointer(sourceParent)), int(sourceRow), int(count), UnsafeNewQModelIndex(unsafe.Pointer(destinationParent)), int(destinationChild))) }) // slot "moveRows"
}

func (this *QStringListModel) OnSort(fn func(int, SortOrder)) {
	registerVirt(this.UnsafeGoID(), 6912441694453382161, func(column int32, order int32) { fn(int(column), SortOrder(order)) }) // slot "sort"
}

func (this *QStringListModel) OnSupportedDropActions(fn func() DropAction) {
	registerVirt(this.UnsafeGoID(), -7784642176186821205, func() int32 { return int32(fn()) }) // slot "supportedDropActions"
}

func (this *QStringListModel) OnIndex(fn func(int, int, *QModelIndex) *QModelIndex) {
	registerVirt(this.UnsafeGoID(), -8948777187306027381, func(row int32, column int32, parent uintptr) uintptr { return uintptr(fn(int(row), int(column), UnsafeNewQModelIndex(unsafe.Pointer(parent))).UnsafePointer()) }) // slot "index"
}

func (this *QStringListModel) OnDropMimeData(fn func(*QMimeData, DropAction, int, int, *QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), -9049994684305682212, func(data uintptr, action int32, row int32, column int32, parent uintptr) int32 { return b2i(fn(UnsafeNewQMimeData(unsafe.Pointer(data)), DropAction(action), int(row), int(column), UnsafeNewQModelIndex(unsafe.Pointer(parent)))) }) // slot "dropMimeData"
}

func (this *QStringListModel) OnHeaderData(fn func(int, Orientation, int) *QVariant) {
	registerVirt(this.UnsafeGoID(), -3297140021896693688, func(section int32, orientation int32, role int32) uintptr { return uintptr(fn(int(section), Orientation(orientation), int(role)).UnsafePointer()) }) // slot "headerData"
}

func (this *QStringListModel) OnSetHeaderData(fn func(int, Orientation, *QVariant, int) bool) {
	registerVirt(this.UnsafeGoID(), 4852200306590054850, func(section int32, orientation int32, value uintptr, role int32) int32 { return b2i(fn(int(section), Orientation(orientation), UnsafeNewQVariant(unsafe.Pointer(value)), int(role))) }) // slot "setHeaderData"
}

func (this *QStringListModel) OnMimeTypes(fn func() []string) {
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

func (this *QStringListModel) OnMimeData(fn func([]*QModelIndex) *QMimeData) {
	registerVirt(this.UnsafeGoID(), 5175297091024974749, func(indexes uintptr, indexes_len int64) uintptr { return uintptr(fn(func() []*QModelIndex { if indexes == 0 || indexes_len == 0 { return nil }; __sl := unsafe.Slice((*unsafe.Pointer)(unsafe.Pointer(indexes)), int(indexes_len)); __o := make([]*QModelIndex, int(indexes_len)); for __i := range __o { __o[__i] = UnsafeNewQModelIndex(__sl[__i]) }; return __o }()).UnsafePointer()) }) // slot "mimeData"
}

func (this *QStringListModel) OnCanDropMimeData(fn func(*QMimeData, DropAction, int, int, *QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), 731539533644753052, func(data uintptr, action int32, row int32, column int32, parent uintptr) int32 { return b2i(fn(UnsafeNewQMimeData(unsafe.Pointer(data)), DropAction(action), int(row), int(column), UnsafeNewQModelIndex(unsafe.Pointer(parent)))) }) // slot "canDropMimeData"
}

func (this *QStringListModel) OnSupportedDragActions(fn func() DropAction) {
	registerVirt(this.UnsafeGoID(), 3421601386088570762, func() int32 { return int32(fn()) }) // slot "supportedDragActions"
}

func (this *QStringListModel) OnInsertColumns(fn func(int, int, *QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), -8973293846230792237, func(column int32, count int32, parent uintptr) int32 { return b2i(fn(int(column), int(count), UnsafeNewQModelIndex(unsafe.Pointer(parent)))) }) // slot "insertColumns"
}

func (this *QStringListModel) OnRemoveColumns(fn func(int, int, *QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), 2078548737056540420, func(column int32, count int32, parent uintptr) int32 { return b2i(fn(int(column), int(count), UnsafeNewQModelIndex(unsafe.Pointer(parent)))) }) // slot "removeColumns"
}

func (this *QStringListModel) OnMoveColumns(fn func(*QModelIndex, int, int, *QModelIndex, int) bool) {
	registerVirt(this.UnsafeGoID(), -6321007591388309905, func(sourceParent uintptr, sourceColumn int32, count int32, destinationParent uintptr, destinationChild int32) int32 { return b2i(fn(UnsafeNewQModelIndex(unsafe.Pointer(sourceParent)), int(sourceColumn), int(count), UnsafeNewQModelIndex(unsafe.Pointer(destinationParent)), int(destinationChild))) }) // slot "moveColumns"
}

func (this *QStringListModel) OnFetchMore(fn func(*QModelIndex)) {
	registerVirt(this.UnsafeGoID(), 3553294593611130806, func(parent uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent))) }) // slot "fetchMore"
}

func (this *QStringListModel) OnCanFetchMore(fn func(*QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), 8639262066473964150, func(parent uintptr) int32 { return b2i(fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)))) }) // slot "canFetchMore"
}

func (this *QStringListModel) OnBuddy(fn func(*QModelIndex) *QModelIndex) {
	registerVirt(this.UnsafeGoID(), 6112860112248786499, func(index uintptr) uintptr { return uintptr(fn(UnsafeNewQModelIndex(unsafe.Pointer(index))).UnsafePointer()) }) // slot "buddy"
}

func (this *QStringListModel) OnMatch(fn func(*QModelIndex, int, *QVariant, int, MatchFlag) []*QModelIndex) {
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

func (this *QStringListModel) OnSpan(fn func(*QModelIndex) *QSize) {
	registerVirt(this.UnsafeGoID(), -8395342916587040543, func(index uintptr) uintptr { return uintptr(fn(UnsafeNewQModelIndex(unsafe.Pointer(index))).UnsafePointer()) }) // slot "span"
}

func (this *QStringListModel) OnMultiData(fn func(*QModelIndex, *QModelRoleDataSpan)) {
	registerVirt(this.UnsafeGoID(), 731311060944096642, func(index uintptr, roleDataSpan uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(index)), UnsafeNewQModelRoleDataSpan(unsafe.Pointer(roleDataSpan))) }) // slot "multiData"
}

func (this *QStringListModel) OnSubmit(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -3788331537442685179, func() int32 { return b2i(fn()) }) // slot "submit"
}

func (this *QStringListModel) OnRevert(fn func()) {
	registerVirt(this.UnsafeGoID(), -1522951607536117149, func() { fn() }) // slot "revert"
}

func (this *QStringListModel) OnResetInternalData(fn func()) {
	registerVirt(this.UnsafeGoID(), -4510675770784051183, func() { fn() }) // slot "resetInternalData"
}

func (this *QStringListModel) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QStringListModel) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QStringListModel) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QStringListModel) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QStringListModel) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QStringListModel) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QStringListModel) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QStringListModel) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QStringListModel_delete(this.h)
	}
	this.UnsafeRelease()
}
