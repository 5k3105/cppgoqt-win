package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QFileSystemModel_chain uintptr
	addr_QFileSystemModel_new uintptr
	addr_QFileSystemModel_new2 uintptr
	addr_QFileSystemModel_tr uintptr
	addr_QFileSystemModel_rootPathChanged uintptr
	addr_QFileSystemModel_fileRenamed uintptr
	addr_QFileSystemModel_directoryLoaded uintptr
	addr_QFileSystemModel_index uintptr
	addr_QFileSystemModel_index2 uintptr
	addr_QFileSystemModel_myComputer uintptr
	addr_QFileSystemModel_setRootPath uintptr
	addr_QFileSystemModel_rootPath uintptr
	addr_QFileSystemModel_rootDirectory uintptr
	addr_QFileSystemModel_setIconProvider uintptr
	addr_QFileSystemModel_iconProvider uintptr
	addr_QFileSystemModel_setFilter uintptr
	addr_QFileSystemModel_filter uintptr
	addr_QFileSystemModel_setResolveSymlinks uintptr
	addr_QFileSystemModel_resolveSymlinks uintptr
	addr_QFileSystemModel_setReadOnly uintptr
	addr_QFileSystemModel_isReadOnly uintptr
	addr_QFileSystemModel_setNameFilterDisables uintptr
	addr_QFileSystemModel_nameFilterDisables uintptr
	addr_QFileSystemModel_setNameFilters uintptr
	addr_QFileSystemModel_nameFilters uintptr
	addr_QFileSystemModel_setOption uintptr
	addr_QFileSystemModel_testOption uintptr
	addr_QFileSystemModel_setOptions uintptr
	addr_QFileSystemModel_options uintptr
	addr_QFileSystemModel_filePath uintptr
	addr_QFileSystemModel_isDir uintptr
	addr_QFileSystemModel_size uintptr
	addr_QFileSystemModel_type uintptr
	addr_QFileSystemModel_lastModified uintptr
	addr_QFileSystemModel_lastModified2 uintptr
	addr_QFileSystemModel_mkdir uintptr
	addr_QFileSystemModel_rmdir uintptr
	addr_QFileSystemModel_fileName uintptr
	addr_QFileSystemModel_fileIcon uintptr
	addr_QFileSystemModel_permissions uintptr
	addr_QFileSystemModel_fileInfo uintptr
	addr_QFileSystemModel_remove uintptr
	addr_QFileSystemModel_createIndex uintptr
	addr_QFileSystemModel_createIndex2 uintptr
	addr_QFileSystemModel_encodeData uintptr
	addr_QFileSystemModel_decodeData uintptr
	addr_QFileSystemModel_beginInsertRows uintptr
	addr_QFileSystemModel_endInsertRows uintptr
	addr_QFileSystemModel_beginRemoveRows uintptr
	addr_QFileSystemModel_endRemoveRows uintptr
	addr_QFileSystemModel_beginMoveRows uintptr
	addr_QFileSystemModel_endMoveRows uintptr
	addr_QFileSystemModel_beginInsertColumns uintptr
	addr_QFileSystemModel_endInsertColumns uintptr
	addr_QFileSystemModel_beginRemoveColumns uintptr
	addr_QFileSystemModel_endRemoveColumns uintptr
	addr_QFileSystemModel_beginMoveColumns uintptr
	addr_QFileSystemModel_endMoveColumns uintptr
	addr_QFileSystemModel_beginResetModel uintptr
	addr_QFileSystemModel_endResetModel uintptr
	addr_QFileSystemModel_changePersistentIndex uintptr
	addr_QFileSystemModel_changePersistentIndexList uintptr
	addr_QFileSystemModel_persistentIndexList uintptr
	addr_QFileSystemModel_connect uintptr
	addr_QFileSystemModel_connect2 uintptr
	addr_QFileSystemModel_sender uintptr
	addr_QFileSystemModel_senderSignalIndex uintptr
	addr_QFileSystemModel_receivers uintptr
	addr_QFileSystemModel_isSignalConnected uintptr
	addr_QFileSystemModel_tr2 uintptr
	addr_QFileSystemModel_tr3 uintptr
	addr_QFileSystemModel_index3 uintptr
	addr_QFileSystemModel_myComputer2 uintptr
	addr_QFileSystemModel_setOption2 uintptr
	addr_QFileSystemModel_connect3 uintptr
	addr_QFileSystemModel_connect4 uintptr
	addr_QFileSystemModel_index_default uintptr
	addr_QFileSystemModel_index_default2 uintptr
	addr_QFileSystemModel_index_default3 uintptr
	addr_QFileSystemModel_metaObject_default uintptr
	addr_QFileSystemModel_parent_default uintptr
	addr_QFileSystemModel_sibling_default uintptr
	addr_QFileSystemModel_hasChildren_default uintptr
	addr_QFileSystemModel_canFetchMore_default uintptr
	addr_QFileSystemModel_fetchMore_default uintptr
	addr_QFileSystemModel_rowCount_default uintptr
	addr_QFileSystemModel_columnCount_default uintptr
	addr_QFileSystemModel_data_default uintptr
	addr_QFileSystemModel_setData_default uintptr
	addr_QFileSystemModel_headerData_default uintptr
	addr_QFileSystemModel_flags_default uintptr
	addr_QFileSystemModel_sort_default uintptr
	addr_QFileSystemModel_mimeTypes_default uintptr
	addr_QFileSystemModel_mimeData_default uintptr
	addr_QFileSystemModel_dropMimeData_default uintptr
	addr_QFileSystemModel_supportedDropActions_default uintptr
	addr_QFileSystemModel_timerEvent_default uintptr
	addr_QFileSystemModel_event_default uintptr
	addr_QFileSystemModel_setHeaderData_default uintptr
	addr_QFileSystemModel_clearItemData_default uintptr
	addr_QFileSystemModel_canDropMimeData_default uintptr
	addr_QFileSystemModel_supportedDragActions_default uintptr
	addr_QFileSystemModel_insertRows_default uintptr
	addr_QFileSystemModel_insertColumns_default uintptr
	addr_QFileSystemModel_removeRows_default uintptr
	addr_QFileSystemModel_removeColumns_default uintptr
	addr_QFileSystemModel_moveRows_default uintptr
	addr_QFileSystemModel_moveColumns_default uintptr
	addr_QFileSystemModel_buddy_default uintptr
	addr_QFileSystemModel_match_default uintptr
	addr_QFileSystemModel_span_default uintptr
	addr_QFileSystemModel_multiData_default uintptr
	addr_QFileSystemModel_submit_default uintptr
	addr_QFileSystemModel_revert_default uintptr
	addr_QFileSystemModel_resetInternalData_default uintptr
	addr_QFileSystemModel_eventFilter_default uintptr
	addr_QFileSystemModel_childEvent_default uintptr
	addr_QFileSystemModel_customEvent_default uintptr
	addr_QFileSystemModel_connectNotify_default uintptr
	addr_QFileSystemModel_disconnectNotify_default uintptr
	pg_QFileSystemModel_connect_rootPathChanged func(uintptr, int64)
	pg_QFileSystemModel_connect_fileRenamed func(uintptr, int64)
	pg_QFileSystemModel_connect_directoryLoaded func(uintptr, int64)
	pg_QFileSystemModel_connect_headerDataChanged func(uintptr, int64)
	pg_QFileSystemModel_connect_layoutChanged func(uintptr, int64)
	pg_QFileSystemModel_connect_layoutChanged2 func(uintptr, int64)
	pg_QFileSystemModel_connect_layoutAboutToBeChanged func(uintptr, int64)
	pg_QFileSystemModel_connect_layoutAboutToBeChanged2 func(uintptr, int64)
	pg_QFileSystemModel_connect_rowsAboutToBeInserted func(uintptr, int64)
	pg_QFileSystemModel_connect_rowsInserted func(uintptr, int64)
	pg_QFileSystemModel_connect_rowsAboutToBeRemoved func(uintptr, int64)
	pg_QFileSystemModel_connect_rowsRemoved func(uintptr, int64)
	pg_QFileSystemModel_connect_columnsAboutToBeInserted func(uintptr, int64)
	pg_QFileSystemModel_connect_columnsInserted func(uintptr, int64)
	pg_QFileSystemModel_connect_columnsAboutToBeRemoved func(uintptr, int64)
	pg_QFileSystemModel_connect_columnsRemoved func(uintptr, int64)
	pg_QFileSystemModel_connect_modelAboutToBeReset func(uintptr, int64)
	pg_QFileSystemModel_connect_modelReset func(uintptr, int64)
	pg_QFileSystemModel_connect_rowsAboutToBeMoved func(uintptr, int64)
	pg_QFileSystemModel_connect_rowsMoved func(uintptr, int64)
	pg_QFileSystemModel_connect_columnsAboutToBeMoved func(uintptr, int64)
	pg_QFileSystemModel_connect_columnsMoved func(uintptr, int64)
	pg_QFileSystemModel_connect_destroyed func(uintptr, int64)
	pg_QFileSystemModel_connect_destroyed2 func(uintptr, int64)
	pg_QFileSystemModel_connect_objectNameChanged func(uintptr, int64)
	pg_QFileSystemModel__disconnect func(int64)
	pg_QFileSystemModel__set_override func(uintptr, int32)
	pg_QFileSystemModel_delete func(uintptr)
)

func init() {
	symaddr(&addr_QFileSystemModel_chain, "QFileSystemModel_chain")
	symaddr(&addr_QFileSystemModel_new, "QFileSystemModel_new")
	symaddr(&addr_QFileSystemModel_new2, "QFileSystemModel_new2")
	symaddr(&addr_QFileSystemModel_tr, "QFileSystemModel_tr")
	symaddr(&addr_QFileSystemModel_rootPathChanged, "QFileSystemModel_rootPathChanged")
	symaddr(&addr_QFileSystemModel_fileRenamed, "QFileSystemModel_fileRenamed")
	symaddr(&addr_QFileSystemModel_directoryLoaded, "QFileSystemModel_directoryLoaded")
	symaddr(&addr_QFileSystemModel_index, "QFileSystemModel_index")
	symaddr(&addr_QFileSystemModel_index2, "QFileSystemModel_index2")
	symaddr(&addr_QFileSystemModel_myComputer, "QFileSystemModel_myComputer")
	symaddr(&addr_QFileSystemModel_setRootPath, "QFileSystemModel_setRootPath")
	symaddr(&addr_QFileSystemModel_rootPath, "QFileSystemModel_rootPath")
	symaddr(&addr_QFileSystemModel_rootDirectory, "QFileSystemModel_rootDirectory")
	symaddr(&addr_QFileSystemModel_setIconProvider, "QFileSystemModel_setIconProvider")
	symaddr(&addr_QFileSystemModel_iconProvider, "QFileSystemModel_iconProvider")
	symaddr(&addr_QFileSystemModel_setFilter, "QFileSystemModel_setFilter")
	symaddr(&addr_QFileSystemModel_filter, "QFileSystemModel_filter")
	symaddr(&addr_QFileSystemModel_setResolveSymlinks, "QFileSystemModel_setResolveSymlinks")
	symaddr(&addr_QFileSystemModel_resolveSymlinks, "QFileSystemModel_resolveSymlinks")
	symaddr(&addr_QFileSystemModel_setReadOnly, "QFileSystemModel_setReadOnly")
	symaddr(&addr_QFileSystemModel_isReadOnly, "QFileSystemModel_isReadOnly")
	symaddr(&addr_QFileSystemModel_setNameFilterDisables, "QFileSystemModel_setNameFilterDisables")
	symaddr(&addr_QFileSystemModel_nameFilterDisables, "QFileSystemModel_nameFilterDisables")
	symaddr(&addr_QFileSystemModel_setNameFilters, "QFileSystemModel_setNameFilters")
	symaddr(&addr_QFileSystemModel_nameFilters, "QFileSystemModel_nameFilters")
	symaddr(&addr_QFileSystemModel_setOption, "QFileSystemModel_setOption")
	symaddr(&addr_QFileSystemModel_testOption, "QFileSystemModel_testOption")
	symaddr(&addr_QFileSystemModel_setOptions, "QFileSystemModel_setOptions")
	symaddr(&addr_QFileSystemModel_options, "QFileSystemModel_options")
	symaddr(&addr_QFileSystemModel_filePath, "QFileSystemModel_filePath")
	symaddr(&addr_QFileSystemModel_isDir, "QFileSystemModel_isDir")
	symaddr(&addr_QFileSystemModel_size, "QFileSystemModel_size")
	symaddr(&addr_QFileSystemModel_type, "QFileSystemModel_type")
	symaddr(&addr_QFileSystemModel_lastModified, "QFileSystemModel_lastModified")
	symaddr(&addr_QFileSystemModel_lastModified2, "QFileSystemModel_lastModified2")
	symaddr(&addr_QFileSystemModel_mkdir, "QFileSystemModel_mkdir")
	symaddr(&addr_QFileSystemModel_rmdir, "QFileSystemModel_rmdir")
	symaddr(&addr_QFileSystemModel_fileName, "QFileSystemModel_fileName")
	symaddr(&addr_QFileSystemModel_fileIcon, "QFileSystemModel_fileIcon")
	symaddr(&addr_QFileSystemModel_permissions, "QFileSystemModel_permissions")
	symaddr(&addr_QFileSystemModel_fileInfo, "QFileSystemModel_fileInfo")
	symaddr(&addr_QFileSystemModel_remove, "QFileSystemModel_remove")
	symaddr(&addr_QFileSystemModel_createIndex, "QFileSystemModel_createIndex")
	symaddr(&addr_QFileSystemModel_createIndex2, "QFileSystemModel_createIndex2")
	symaddr(&addr_QFileSystemModel_encodeData, "QFileSystemModel_encodeData")
	symaddr(&addr_QFileSystemModel_decodeData, "QFileSystemModel_decodeData")
	symaddr(&addr_QFileSystemModel_beginInsertRows, "QFileSystemModel_beginInsertRows")
	symaddr(&addr_QFileSystemModel_endInsertRows, "QFileSystemModel_endInsertRows")
	symaddr(&addr_QFileSystemModel_beginRemoveRows, "QFileSystemModel_beginRemoveRows")
	symaddr(&addr_QFileSystemModel_endRemoveRows, "QFileSystemModel_endRemoveRows")
	symaddr(&addr_QFileSystemModel_beginMoveRows, "QFileSystemModel_beginMoveRows")
	symaddr(&addr_QFileSystemModel_endMoveRows, "QFileSystemModel_endMoveRows")
	symaddr(&addr_QFileSystemModel_beginInsertColumns, "QFileSystemModel_beginInsertColumns")
	symaddr(&addr_QFileSystemModel_endInsertColumns, "QFileSystemModel_endInsertColumns")
	symaddr(&addr_QFileSystemModel_beginRemoveColumns, "QFileSystemModel_beginRemoveColumns")
	symaddr(&addr_QFileSystemModel_endRemoveColumns, "QFileSystemModel_endRemoveColumns")
	symaddr(&addr_QFileSystemModel_beginMoveColumns, "QFileSystemModel_beginMoveColumns")
	symaddr(&addr_QFileSystemModel_endMoveColumns, "QFileSystemModel_endMoveColumns")
	symaddr(&addr_QFileSystemModel_beginResetModel, "QFileSystemModel_beginResetModel")
	symaddr(&addr_QFileSystemModel_endResetModel, "QFileSystemModel_endResetModel")
	symaddr(&addr_QFileSystemModel_changePersistentIndex, "QFileSystemModel_changePersistentIndex")
	symaddr(&addr_QFileSystemModel_changePersistentIndexList, "QFileSystemModel_changePersistentIndexList")
	symaddr(&addr_QFileSystemModel_persistentIndexList, "QFileSystemModel_persistentIndexList")
	symaddr(&addr_QFileSystemModel_connect, "QFileSystemModel_connect")
	symaddr(&addr_QFileSystemModel_connect2, "QFileSystemModel_connect2")
	symaddr(&addr_QFileSystemModel_sender, "QFileSystemModel_sender")
	symaddr(&addr_QFileSystemModel_senderSignalIndex, "QFileSystemModel_senderSignalIndex")
	symaddr(&addr_QFileSystemModel_receivers, "QFileSystemModel_receivers")
	symaddr(&addr_QFileSystemModel_isSignalConnected, "QFileSystemModel_isSignalConnected")
	symaddr(&addr_QFileSystemModel_tr2, "QFileSystemModel_tr2")
	symaddr(&addr_QFileSystemModel_tr3, "QFileSystemModel_tr3")
	symaddr(&addr_QFileSystemModel_index3, "QFileSystemModel_index3")
	symaddr(&addr_QFileSystemModel_myComputer2, "QFileSystemModel_myComputer2")
	symaddr(&addr_QFileSystemModel_setOption2, "QFileSystemModel_setOption2")
	symaddr(&addr_QFileSystemModel_connect3, "QFileSystemModel_connect3")
	symaddr(&addr_QFileSystemModel_connect4, "QFileSystemModel_connect4")
	symaddr(&addr_QFileSystemModel_index_default, "QFileSystemModel_index_default")
	symaddr(&addr_QFileSystemModel_index_default2, "QFileSystemModel_index_default2")
	symaddr(&addr_QFileSystemModel_index_default3, "QFileSystemModel_index_default3")
	symaddr(&addr_QFileSystemModel_metaObject_default, "QFileSystemModel_metaObject_default")
	symaddr(&addr_QFileSystemModel_parent_default, "QFileSystemModel_parent_default")
	symaddr(&addr_QFileSystemModel_sibling_default, "QFileSystemModel_sibling_default")
	symaddr(&addr_QFileSystemModel_hasChildren_default, "QFileSystemModel_hasChildren_default")
	symaddr(&addr_QFileSystemModel_canFetchMore_default, "QFileSystemModel_canFetchMore_default")
	symaddr(&addr_QFileSystemModel_fetchMore_default, "QFileSystemModel_fetchMore_default")
	symaddr(&addr_QFileSystemModel_rowCount_default, "QFileSystemModel_rowCount_default")
	symaddr(&addr_QFileSystemModel_columnCount_default, "QFileSystemModel_columnCount_default")
	symaddr(&addr_QFileSystemModel_data_default, "QFileSystemModel_data_default")
	symaddr(&addr_QFileSystemModel_setData_default, "QFileSystemModel_setData_default")
	symaddr(&addr_QFileSystemModel_headerData_default, "QFileSystemModel_headerData_default")
	symaddr(&addr_QFileSystemModel_flags_default, "QFileSystemModel_flags_default")
	symaddr(&addr_QFileSystemModel_sort_default, "QFileSystemModel_sort_default")
	symaddr(&addr_QFileSystemModel_mimeTypes_default, "QFileSystemModel_mimeTypes_default")
	symaddr(&addr_QFileSystemModel_mimeData_default, "QFileSystemModel_mimeData_default")
	symaddr(&addr_QFileSystemModel_dropMimeData_default, "QFileSystemModel_dropMimeData_default")
	symaddr(&addr_QFileSystemModel_supportedDropActions_default, "QFileSystemModel_supportedDropActions_default")
	symaddr(&addr_QFileSystemModel_timerEvent_default, "QFileSystemModel_timerEvent_default")
	symaddr(&addr_QFileSystemModel_event_default, "QFileSystemModel_event_default")
	symaddr(&addr_QFileSystemModel_setHeaderData_default, "QFileSystemModel_setHeaderData_default")
	symaddr(&addr_QFileSystemModel_clearItemData_default, "QFileSystemModel_clearItemData_default")
	symaddr(&addr_QFileSystemModel_canDropMimeData_default, "QFileSystemModel_canDropMimeData_default")
	symaddr(&addr_QFileSystemModel_supportedDragActions_default, "QFileSystemModel_supportedDragActions_default")
	symaddr(&addr_QFileSystemModel_insertRows_default, "QFileSystemModel_insertRows_default")
	symaddr(&addr_QFileSystemModel_insertColumns_default, "QFileSystemModel_insertColumns_default")
	symaddr(&addr_QFileSystemModel_removeRows_default, "QFileSystemModel_removeRows_default")
	symaddr(&addr_QFileSystemModel_removeColumns_default, "QFileSystemModel_removeColumns_default")
	symaddr(&addr_QFileSystemModel_moveRows_default, "QFileSystemModel_moveRows_default")
	symaddr(&addr_QFileSystemModel_moveColumns_default, "QFileSystemModel_moveColumns_default")
	symaddr(&addr_QFileSystemModel_buddy_default, "QFileSystemModel_buddy_default")
	symaddr(&addr_QFileSystemModel_match_default, "QFileSystemModel_match_default")
	symaddr(&addr_QFileSystemModel_span_default, "QFileSystemModel_span_default")
	symaddr(&addr_QFileSystemModel_multiData_default, "QFileSystemModel_multiData_default")
	symaddr(&addr_QFileSystemModel_submit_default, "QFileSystemModel_submit_default")
	symaddr(&addr_QFileSystemModel_revert_default, "QFileSystemModel_revert_default")
	symaddr(&addr_QFileSystemModel_resetInternalData_default, "QFileSystemModel_resetInternalData_default")
	symaddr(&addr_QFileSystemModel_eventFilter_default, "QFileSystemModel_eventFilter_default")
	symaddr(&addr_QFileSystemModel_childEvent_default, "QFileSystemModel_childEvent_default")
	symaddr(&addr_QFileSystemModel_customEvent_default, "QFileSystemModel_customEvent_default")
	symaddr(&addr_QFileSystemModel_connectNotify_default, "QFileSystemModel_connectNotify_default")
	symaddr(&addr_QFileSystemModel_disconnectNotify_default, "QFileSystemModel_disconnectNotify_default")
	reg(&pg_QFileSystemModel_connect_rootPathChanged, "QFileSystemModel_connect_rootPathChanged")
	reg(&pg_QFileSystemModel_connect_fileRenamed, "QFileSystemModel_connect_fileRenamed")
	reg(&pg_QFileSystemModel_connect_directoryLoaded, "QFileSystemModel_connect_directoryLoaded")
	reg(&pg_QFileSystemModel_connect_headerDataChanged, "QFileSystemModel_connect_headerDataChanged")
	reg(&pg_QFileSystemModel_connect_layoutChanged, "QFileSystemModel_connect_layoutChanged")
	reg(&pg_QFileSystemModel_connect_layoutChanged2, "QFileSystemModel_connect_layoutChanged2")
	reg(&pg_QFileSystemModel_connect_layoutAboutToBeChanged, "QFileSystemModel_connect_layoutAboutToBeChanged")
	reg(&pg_QFileSystemModel_connect_layoutAboutToBeChanged2, "QFileSystemModel_connect_layoutAboutToBeChanged2")
	reg(&pg_QFileSystemModel_connect_rowsAboutToBeInserted, "QFileSystemModel_connect_rowsAboutToBeInserted")
	reg(&pg_QFileSystemModel_connect_rowsInserted, "QFileSystemModel_connect_rowsInserted")
	reg(&pg_QFileSystemModel_connect_rowsAboutToBeRemoved, "QFileSystemModel_connect_rowsAboutToBeRemoved")
	reg(&pg_QFileSystemModel_connect_rowsRemoved, "QFileSystemModel_connect_rowsRemoved")
	reg(&pg_QFileSystemModel_connect_columnsAboutToBeInserted, "QFileSystemModel_connect_columnsAboutToBeInserted")
	reg(&pg_QFileSystemModel_connect_columnsInserted, "QFileSystemModel_connect_columnsInserted")
	reg(&pg_QFileSystemModel_connect_columnsAboutToBeRemoved, "QFileSystemModel_connect_columnsAboutToBeRemoved")
	reg(&pg_QFileSystemModel_connect_columnsRemoved, "QFileSystemModel_connect_columnsRemoved")
	reg(&pg_QFileSystemModel_connect_modelAboutToBeReset, "QFileSystemModel_connect_modelAboutToBeReset")
	reg(&pg_QFileSystemModel_connect_modelReset, "QFileSystemModel_connect_modelReset")
	reg(&pg_QFileSystemModel_connect_rowsAboutToBeMoved, "QFileSystemModel_connect_rowsAboutToBeMoved")
	reg(&pg_QFileSystemModel_connect_rowsMoved, "QFileSystemModel_connect_rowsMoved")
	reg(&pg_QFileSystemModel_connect_columnsAboutToBeMoved, "QFileSystemModel_connect_columnsAboutToBeMoved")
	reg(&pg_QFileSystemModel_connect_columnsMoved, "QFileSystemModel_connect_columnsMoved")
	reg(&pg_QFileSystemModel_connect_destroyed, "QFileSystemModel_connect_destroyed")
	reg(&pg_QFileSystemModel_connect_destroyed2, "QFileSystemModel_connect_destroyed2")
	reg(&pg_QFileSystemModel_connect_objectNameChanged, "QFileSystemModel_connect_objectNameChanged")
	reg(&pg_QFileSystemModel__disconnect, "QFileSystemModel__disconnect")
	reg(&pg_QFileSystemModel__set_override, "QFileSystemModel__set_override")
	reg(&pg_QFileSystemModel_delete, "QFileSystemModel_delete")
}

type QFileSystemModel struct {
	h uintptr
	*QAbstractItemModel
}

func bindQFileSystemModel(h uintptr) *QFileSystemModel {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QFileSystemModel_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QFileSystemModel{h: h, QAbstractItemModel: &QAbstractItemModel{h: __c[0], QObject: &QObject{h: __c[1]}}}
}

func newownedQFileSystemModel(h uintptr) *QFileSystemModel { return bindQFileSystemModel(h) }

func (this *QFileSystemModel) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QFileSystemModel) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQFileSystemModel(h unsafe.Pointer) *QFileSystemModel { return bindQFileSystemModel(uintptr(h)) }
func UnsafeNewQFileSystemModelOwned(h unsafe.Pointer) *QFileSystemModel { return newownedQFileSystemModel(uintptr(h)) }

func (this *QFileSystemModel) UnsafeRelease() {
	this.h = 0
	if this.QAbstractItemModel != nil {
		this.QAbstractItemModel.UnsafeRelease()
	}
}

func NewQFileSystemModel() *QFileSystemModel {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QFileSystemModel_new, uintptr(__goid))
	__o := newownedQFileSystemModel(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QFileSystemModel_virtIdx, QFileSystemModel_enable)
	return __o
}

func NewQFileSystemModel2(parent *QObject) *QFileSystemModel {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QFileSystemModel_new2, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQFileSystemModel(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QFileSystemModel_virtIdx, QFileSystemModel_enable)
	return __o
}

func QFileSystemModel_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFileSystemModel_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QFileSystemModel) RootPathChanged(newPath string) {
	var newPath_p uintptr
	if len(newPath) > 0 {
		newPath_p = uintptr(unsafe.Pointer(unsafe.StringData(newPath)))
	}
	purego.SyscallN(addr_QFileSystemModel_rootPathChanged, this.h, newPath_p, uintptr(len(newPath)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(newPath)
}

func (this *QFileSystemModel) FileRenamed(path string, oldName string, newName string) {
	var path_p uintptr
	if len(path) > 0 {
		path_p = uintptr(unsafe.Pointer(unsafe.StringData(path)))
	}
	var oldName_p uintptr
	if len(oldName) > 0 {
		oldName_p = uintptr(unsafe.Pointer(unsafe.StringData(oldName)))
	}
	var newName_p uintptr
	if len(newName) > 0 {
		newName_p = uintptr(unsafe.Pointer(unsafe.StringData(newName)))
	}
	purego.SyscallN(addr_QFileSystemModel_fileRenamed, this.h, path_p, uintptr(len(path)), oldName_p, uintptr(len(oldName)), newName_p, uintptr(len(newName)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(path)
	runtime.KeepAlive(oldName)
	runtime.KeepAlive(newName)
}

func (this *QFileSystemModel) DirectoryLoaded(path string) {
	var path_p uintptr
	if len(path) > 0 {
		path_p = uintptr(unsafe.Pointer(unsafe.StringData(path)))
	}
	purego.SyscallN(addr_QFileSystemModel_directoryLoaded, this.h, path_p, uintptr(len(path)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(path)
}

func (this *QFileSystemModel) Index(row int, column int, parent *QModelIndex) *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QFileSystemModel_index, this.h, uintptr(row), uintptr(column), uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return newownedQModelIndex(__r)
}

func (this *QFileSystemModel) IndexWithPath(path string) *QModelIndex {
	var path_p uintptr
	if len(path) > 0 {
		path_p = uintptr(unsafe.Pointer(unsafe.StringData(path)))
	}
	__r, _, _ := purego.SyscallN(addr_QFileSystemModel_index2, this.h, path_p, uintptr(len(path)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(path)
	return newownedQModelIndex(__r)
}

func (this *QFileSystemModel) MyComputer() *QVariant {
	__r, _, _ := purego.SyscallN(addr_QFileSystemModel_myComputer, this.h)
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QFileSystemModel) SetRootPath(path string) *QModelIndex {
	var path_p uintptr
	if len(path) > 0 {
		path_p = uintptr(unsafe.Pointer(unsafe.StringData(path)))
	}
	__r, _, _ := purego.SyscallN(addr_QFileSystemModel_setRootPath, this.h, path_p, uintptr(len(path)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(path)
	return newownedQModelIndex(__r)
}

func (this *QFileSystemModel) RootPath() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFileSystemModel_rootPath, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QFileSystemModel) RootDirectory() *QDir {
	__r, _, _ := purego.SyscallN(addr_QFileSystemModel_rootDirectory, this.h)
	runtime.KeepAlive(this)
	return newownedQDir(__r)
}

func (this *QFileSystemModel) SetIconProvider(provider *QAbstractFileIconProvider) {
	purego.SyscallN(addr_QFileSystemModel_setIconProvider, this.h, uintptr(provider.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(provider)
}

func (this *QFileSystemModel) IconProvider() *QAbstractFileIconProvider {
	__r, _, _ := purego.SyscallN(addr_QFileSystemModel_iconProvider, this.h)
	runtime.KeepAlive(this)
	return bindQAbstractFileIconProvider(__r)
}

func (this *QFileSystemModel) SetFilter(filters QDir__Filter) {
	purego.SyscallN(addr_QFileSystemModel_setFilter, this.h, uintptr(filters))
	runtime.KeepAlive(this)
}

func (this *QFileSystemModel) Filter() QDir__Filter {
	__r0, _, _ := purego.SyscallN(addr_QFileSystemModel_filter, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QDir__Filter(__r)
}

func (this *QFileSystemModel) SetResolveSymlinks(enable bool) {
	purego.SyscallN(addr_QFileSystemModel_setResolveSymlinks, this.h, uintptr(b2i(enable)))
	runtime.KeepAlive(this)
}

func (this *QFileSystemModel) ResolveSymlinks() bool {
	__r0, _, _ := purego.SyscallN(addr_QFileSystemModel_resolveSymlinks, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFileSystemModel) SetReadOnly(enable bool) {
	purego.SyscallN(addr_QFileSystemModel_setReadOnly, this.h, uintptr(b2i(enable)))
	runtime.KeepAlive(this)
}

func (this *QFileSystemModel) IsReadOnly() bool {
	__r0, _, _ := purego.SyscallN(addr_QFileSystemModel_isReadOnly, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFileSystemModel) SetNameFilterDisables(enable bool) {
	purego.SyscallN(addr_QFileSystemModel_setNameFilterDisables, this.h, uintptr(b2i(enable)))
	runtime.KeepAlive(this)
}

func (this *QFileSystemModel) NameFilterDisables() bool {
	__r0, _, _ := purego.SyscallN(addr_QFileSystemModel_nameFilterDisables, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFileSystemModel) SetNameFilters(filters []string) {
	filters_arr := make([]uintptr, len(filters))
	for __i, __e := range filters {
		filters_arr[__i] = pgCString(__e)
	}
	var filters_p uintptr
	if len(filters_arr) > 0 {
		filters_p = uintptr(unsafe.Pointer(&filters_arr[0]))
	}
	purego.SyscallN(addr_QFileSystemModel_setNameFilters, this.h, filters_p, uintptr(len(filters)))
	runtime.KeepAlive(this)
	for __i := range filters_arr {
		pgFree(filters_arr[__i])
	}
	runtime.KeepAlive(filters_arr)
}

func (this *QFileSystemModel) NameFilters() []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFileSystemModel_nameFilters, this.h, uintptr(unsafe.Pointer(&__n)))
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

func (this *QFileSystemModel) SetOption(option QFileSystemModel__Option) {
	purego.SyscallN(addr_QFileSystemModel_setOption, this.h, uintptr(option))
	runtime.KeepAlive(this)
}

func (this *QFileSystemModel) TestOption(option QFileSystemModel__Option) bool {
	__r0, _, _ := purego.SyscallN(addr_QFileSystemModel_testOption, this.h, uintptr(option))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFileSystemModel) SetOptions(options QFileSystemModel__Option) {
	purego.SyscallN(addr_QFileSystemModel_setOptions, this.h, uintptr(options))
	runtime.KeepAlive(this)
}

func (this *QFileSystemModel) Options() QFileSystemModel__Option {
	__r0, _, _ := purego.SyscallN(addr_QFileSystemModel_options, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QFileSystemModel__Option(__r)
}

func (this *QFileSystemModel) FilePath(index *QModelIndex) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFileSystemModel_filePath, this.h, uintptr(index.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return pgStr(__r, __n)
}

func (this *QFileSystemModel) IsDir(index *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QFileSystemModel_isDir, this.h, uintptr(index.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return __r != 0
}

func (this *QFileSystemModel) Size(index *QModelIndex) int64 {
	__r0, _, _ := purego.SyscallN(addr_QFileSystemModel_size, this.h, uintptr(index.UnsafePointer()))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return __r
}

func (this *QFileSystemModel) Type(index *QModelIndex) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFileSystemModel_type, this.h, uintptr(index.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return pgStr(__r, __n)
}

func (this *QFileSystemModel) LastModified(index *QModelIndex) *QDateTime {
	__r, _, _ := purego.SyscallN(addr_QFileSystemModel_lastModified, this.h, uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return newownedQDateTime(__r)
}

func (this *QFileSystemModel) LastModified2(index *QModelIndex, tz *QTimeZone) *QDateTime {
	__r, _, _ := purego.SyscallN(addr_QFileSystemModel_lastModified2, this.h, uintptr(index.UnsafePointer()), uintptr(tz.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	runtime.KeepAlive(tz)
	return newownedQDateTime(__r)
}

func (this *QFileSystemModel) Mkdir(parent *QModelIndex, name string) *QModelIndex {
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	__r, _, _ := purego.SyscallN(addr_QFileSystemModel_mkdir, this.h, uintptr(parent.UnsafePointer()), name_p, uintptr(len(name)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	runtime.KeepAlive(name)
	return newownedQModelIndex(__r)
}

func (this *QFileSystemModel) Rmdir(index *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QFileSystemModel_rmdir, this.h, uintptr(index.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return __r != 0
}

func (this *QFileSystemModel) FileName(index *QModelIndex) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFileSystemModel_fileName, this.h, uintptr(index.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return pgStr(__r, __n)
}

func (this *QFileSystemModel) FileIcon(index *QModelIndex) *QIcon {
	__r, _, _ := purego.SyscallN(addr_QFileSystemModel_fileIcon, this.h, uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return newownedQIcon(__r)
}

func (this *QFileSystemModel) Permissions(index *QModelIndex) QFileDevice__Permission {
	__r0, _, _ := purego.SyscallN(addr_QFileSystemModel_permissions, this.h, uintptr(index.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return QFileDevice__Permission(__r)
}

func (this *QFileSystemModel) FileInfo(index *QModelIndex) *QFileInfo {
	__r, _, _ := purego.SyscallN(addr_QFileSystemModel_fileInfo, this.h, uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return newownedQFileInfo(__r)
}

func (this *QFileSystemModel) Remove(index *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QFileSystemModel_remove, this.h, uintptr(index.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return __r != 0
}

func (this *QFileSystemModel) CreateIndex(row int, column int) *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QFileSystemModel_createIndex, this.h, uintptr(row), uintptr(column))
	runtime.KeepAlive(this)
	return newownedQModelIndex(__r)
}

func (this *QFileSystemModel) CreateIndex2(row int, column int, id uintptr) *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QFileSystemModel_createIndex2, this.h, uintptr(row), uintptr(column), uintptr(id))
	runtime.KeepAlive(this)
	return newownedQModelIndex(__r)
}

func (this *QFileSystemModel) EncodeData(indexes []*QModelIndex, stream *QDataStream) {
	indexes_arr := make([]uintptr, len(indexes))
	for __i, __e := range indexes {
		indexes_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var indexes_p uintptr
	if len(indexes_arr) > 0 {
		indexes_p = uintptr(unsafe.Pointer(&indexes_arr[0]))
	}
	purego.SyscallN(addr_QFileSystemModel_encodeData, this.h, indexes_p, uintptr(len(indexes)), uintptr(stream.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(indexes)
	runtime.KeepAlive(indexes_arr)
	runtime.KeepAlive(stream)
}

func (this *QFileSystemModel) DecodeData(row int, column int, parent *QModelIndex, stream *QDataStream) bool {
	__r0, _, _ := purego.SyscallN(addr_QFileSystemModel_decodeData, this.h, uintptr(row), uintptr(column), uintptr(parent.UnsafePointer()), uintptr(stream.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	runtime.KeepAlive(stream)
	return __r != 0
}

func (this *QFileSystemModel) BeginInsertRows(parent *QModelIndex, first int, last int) {
	purego.SyscallN(addr_QFileSystemModel_beginInsertRows, this.h, uintptr(parent.UnsafePointer()), uintptr(first), uintptr(last))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
}

func (this *QFileSystemModel) EndInsertRows() {
	purego.SyscallN(addr_QFileSystemModel_endInsertRows, this.h)
	runtime.KeepAlive(this)
}

func (this *QFileSystemModel) BeginRemoveRows(parent *QModelIndex, first int, last int) {
	purego.SyscallN(addr_QFileSystemModel_beginRemoveRows, this.h, uintptr(parent.UnsafePointer()), uintptr(first), uintptr(last))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
}

func (this *QFileSystemModel) EndRemoveRows() {
	purego.SyscallN(addr_QFileSystemModel_endRemoveRows, this.h)
	runtime.KeepAlive(this)
}

func (this *QFileSystemModel) BeginMoveRows(sourceParent *QModelIndex, sourceFirst int, sourceLast int, destinationParent *QModelIndex, destinationRow int) bool {
	__r0, _, _ := purego.SyscallN(addr_QFileSystemModel_beginMoveRows, this.h, uintptr(sourceParent.UnsafePointer()), uintptr(sourceFirst), uintptr(sourceLast), uintptr(destinationParent.UnsafePointer()), uintptr(destinationRow))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sourceParent)
	runtime.KeepAlive(destinationParent)
	return __r != 0
}

func (this *QFileSystemModel) EndMoveRows() {
	purego.SyscallN(addr_QFileSystemModel_endMoveRows, this.h)
	runtime.KeepAlive(this)
}

func (this *QFileSystemModel) BeginInsertColumns(parent *QModelIndex, first int, last int) {
	purego.SyscallN(addr_QFileSystemModel_beginInsertColumns, this.h, uintptr(parent.UnsafePointer()), uintptr(first), uintptr(last))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
}

func (this *QFileSystemModel) EndInsertColumns() {
	purego.SyscallN(addr_QFileSystemModel_endInsertColumns, this.h)
	runtime.KeepAlive(this)
}

func (this *QFileSystemModel) BeginRemoveColumns(parent *QModelIndex, first int, last int) {
	purego.SyscallN(addr_QFileSystemModel_beginRemoveColumns, this.h, uintptr(parent.UnsafePointer()), uintptr(first), uintptr(last))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
}

func (this *QFileSystemModel) EndRemoveColumns() {
	purego.SyscallN(addr_QFileSystemModel_endRemoveColumns, this.h)
	runtime.KeepAlive(this)
}

func (this *QFileSystemModel) BeginMoveColumns(sourceParent *QModelIndex, sourceFirst int, sourceLast int, destinationParent *QModelIndex, destinationColumn int) bool {
	__r0, _, _ := purego.SyscallN(addr_QFileSystemModel_beginMoveColumns, this.h, uintptr(sourceParent.UnsafePointer()), uintptr(sourceFirst), uintptr(sourceLast), uintptr(destinationParent.UnsafePointer()), uintptr(destinationColumn))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sourceParent)
	runtime.KeepAlive(destinationParent)
	return __r != 0
}

func (this *QFileSystemModel) EndMoveColumns() {
	purego.SyscallN(addr_QFileSystemModel_endMoveColumns, this.h)
	runtime.KeepAlive(this)
}

func (this *QFileSystemModel) BeginResetModel() {
	purego.SyscallN(addr_QFileSystemModel_beginResetModel, this.h)
	runtime.KeepAlive(this)
}

func (this *QFileSystemModel) EndResetModel() {
	purego.SyscallN(addr_QFileSystemModel_endResetModel, this.h)
	runtime.KeepAlive(this)
}

func (this *QFileSystemModel) ChangePersistentIndex(from *QModelIndex, to *QModelIndex) {
	purego.SyscallN(addr_QFileSystemModel_changePersistentIndex, this.h, uintptr(from.UnsafePointer()), uintptr(to.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(from)
	runtime.KeepAlive(to)
}

func (this *QFileSystemModel) ChangePersistentIndexList(from []*QModelIndex, to []*QModelIndex) {
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
	purego.SyscallN(addr_QFileSystemModel_changePersistentIndexList, this.h, from_p, uintptr(len(from)), to_p, uintptr(len(to)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(from)
	runtime.KeepAlive(from_arr)
	runtime.KeepAlive(to)
	runtime.KeepAlive(to_arr)
}

func (this *QFileSystemModel) PersistentIndexList() []*QModelIndex {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFileSystemModel_persistentIndexList, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QModelIndex, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQModelIndex(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QFileSystemModel_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QFileSystemModel_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QFileSystemModel_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QFileSystemModel_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QFileSystemModel) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QFileSystemModel_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QFileSystemModel) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QFileSystemModel_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QFileSystemModel) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QFileSystemModel_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QFileSystemModel) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QFileSystemModel_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QFileSystemModel_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFileSystemModel_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QFileSystemModel_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFileSystemModel_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func (this *QFileSystemModel) Index2(path string, column int) *QModelIndex {
	var path_p uintptr
	if len(path) > 0 {
		path_p = uintptr(unsafe.Pointer(unsafe.StringData(path)))
	}
	__r, _, _ := purego.SyscallN(addr_QFileSystemModel_index3, this.h, path_p, uintptr(len(path)), uintptr(column))
	runtime.KeepAlive(this)
	runtime.KeepAlive(path)
	return newownedQModelIndex(__r)
}

func (this *QFileSystemModel) MyComputerWithRole(role int) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QFileSystemModel_myComputer2, this.h, uintptr(role))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QFileSystemModel) SetOption2(option QFileSystemModel__Option, on bool) {
	purego.SyscallN(addr_QFileSystemModel_setOption2, this.h, uintptr(option), uintptr(b2i(on)))
	runtime.KeepAlive(this)
}

func QFileSystemModel_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QFileSystemModel_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QFileSystemModel_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QFileSystemModel_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QFileSystemModel) IndexDefault(row int, column int, parent *QModelIndex) *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QFileSystemModel_index_default, this.h, uintptr(row), uintptr(column), uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return newownedQModelIndex(__r)
}

func (this *QFileSystemModel) IndexWithPathDefault(path string) *QModelIndex {
	var path_p uintptr
	if len(path) > 0 {
		path_p = uintptr(unsafe.Pointer(unsafe.StringData(path)))
	}
	__r, _, _ := purego.SyscallN(addr_QFileSystemModel_index_default2, this.h, path_p, uintptr(len(path)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(path)
	return newownedQModelIndex(__r)
}

func (this *QFileSystemModel) Index2Default(path string, column int) *QModelIndex {
	var path_p uintptr
	if len(path) > 0 {
		path_p = uintptr(unsafe.Pointer(unsafe.StringData(path)))
	}
	__r, _, _ := purego.SyscallN(addr_QFileSystemModel_index_default3, this.h, path_p, uintptr(len(path)), uintptr(column))
	runtime.KeepAlive(this)
	runtime.KeepAlive(path)
	return newownedQModelIndex(__r)
}

func (this *QFileSystemModel) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QFileSystemModel_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QFileSystemModel) ParentDefault(child *QModelIndex) *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QFileSystemModel_parent_default, this.h, uintptr(child.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(child)
	return newownedQModelIndex(__r)
}

func (this *QFileSystemModel) SiblingDefault(row int, column int, idx *QModelIndex) *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QFileSystemModel_sibling_default, this.h, uintptr(row), uintptr(column), uintptr(idx.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(idx)
	return newownedQModelIndex(__r)
}

func (this *QFileSystemModel) HasChildrenDefault(parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QFileSystemModel_hasChildren_default, this.h, uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QFileSystemModel) CanFetchMoreDefault(parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QFileSystemModel_canFetchMore_default, this.h, uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QFileSystemModel) FetchMoreDefault(parent *QModelIndex) {
	purego.SyscallN(addr_QFileSystemModel_fetchMore_default, this.h, uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
}

func (this *QFileSystemModel) RowCountDefault(parent *QModelIndex) int {
	__r0, _, _ := purego.SyscallN(addr_QFileSystemModel_rowCount_default, this.h, uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return int(__r)
}

func (this *QFileSystemModel) ColumnCountDefault(parent *QModelIndex) int {
	__r0, _, _ := purego.SyscallN(addr_QFileSystemModel_columnCount_default, this.h, uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return int(__r)
}

func (this *QFileSystemModel) DataDefault(index *QModelIndex, role int) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QFileSystemModel_data_default, this.h, uintptr(index.UnsafePointer()), uintptr(role))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return newownedQVariant(__r)
}

func (this *QFileSystemModel) SetDataDefault(index *QModelIndex, value *QVariant, role int) bool {
	__r0, _, _ := purego.SyscallN(addr_QFileSystemModel_setData_default, this.h, uintptr(index.UnsafePointer()), uintptr(value.UnsafePointer()), uintptr(role))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	runtime.KeepAlive(value)
	return __r != 0
}

func (this *QFileSystemModel) HeaderDataDefault(section int, orientation Orientation, role int) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QFileSystemModel_headerData_default, this.h, uintptr(section), uintptr(orientation), uintptr(role))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QFileSystemModel) FlagsDefault(index *QModelIndex) ItemFlag {
	__r0, _, _ := purego.SyscallN(addr_QFileSystemModel_flags_default, this.h, uintptr(index.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return ItemFlag(__r)
}

func (this *QFileSystemModel) SortDefault(column int, order SortOrder) {
	purego.SyscallN(addr_QFileSystemModel_sort_default, this.h, uintptr(column), uintptr(order))
	runtime.KeepAlive(this)
}

func (this *QFileSystemModel) MimeTypesDefault() []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFileSystemModel_mimeTypes_default, this.h, uintptr(unsafe.Pointer(&__n)))
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

func (this *QFileSystemModel) MimeDataDefault(indexes []*QModelIndex) *QMimeData {
	indexes_arr := make([]uintptr, len(indexes))
	for __i, __e := range indexes {
		indexes_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var indexes_p uintptr
	if len(indexes_arr) > 0 {
		indexes_p = uintptr(unsafe.Pointer(&indexes_arr[0]))
	}
	__r, _, _ := purego.SyscallN(addr_QFileSystemModel_mimeData_default, this.h, indexes_p, uintptr(len(indexes)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(indexes)
	runtime.KeepAlive(indexes_arr)
	return bindQMimeData(__r)
}

func (this *QFileSystemModel) DropMimeDataDefault(data *QMimeData, action DropAction, row int, column int, parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QFileSystemModel_dropMimeData_default, this.h, uintptr(data.UnsafePointer()), uintptr(action), uintptr(row), uintptr(column), uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(data)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QFileSystemModel) SupportedDropActionsDefault() DropAction {
	__r0, _, _ := purego.SyscallN(addr_QFileSystemModel_supportedDropActions_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return DropAction(__r)
}

func (this *QFileSystemModel) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QFileSystemModel_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFileSystemModel) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QFileSystemModel_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QFileSystemModel) SetHeaderDataDefault(section int, orientation Orientation, value *QVariant, role int) bool {
	__r0, _, _ := purego.SyscallN(addr_QFileSystemModel_setHeaderData_default, this.h, uintptr(section), uintptr(orientation), uintptr(value.UnsafePointer()), uintptr(role))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(value)
	return __r != 0
}

func (this *QFileSystemModel) ClearItemDataDefault(index *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QFileSystemModel_clearItemData_default, this.h, uintptr(index.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return __r != 0
}

func (this *QFileSystemModel) CanDropMimeDataDefault(data *QMimeData, action DropAction, row int, column int, parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QFileSystemModel_canDropMimeData_default, this.h, uintptr(data.UnsafePointer()), uintptr(action), uintptr(row), uintptr(column), uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(data)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QFileSystemModel) SupportedDragActionsDefault() DropAction {
	__r0, _, _ := purego.SyscallN(addr_QFileSystemModel_supportedDragActions_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return DropAction(__r)
}

func (this *QFileSystemModel) InsertRowsDefault(row int, count int, parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QFileSystemModel_insertRows_default, this.h, uintptr(row), uintptr(count), uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QFileSystemModel) InsertColumnsDefault(column int, count int, parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QFileSystemModel_insertColumns_default, this.h, uintptr(column), uintptr(count), uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QFileSystemModel) RemoveRowsDefault(row int, count int, parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QFileSystemModel_removeRows_default, this.h, uintptr(row), uintptr(count), uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QFileSystemModel) RemoveColumnsDefault(column int, count int, parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QFileSystemModel_removeColumns_default, this.h, uintptr(column), uintptr(count), uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QFileSystemModel) MoveRowsDefault(sourceParent *QModelIndex, sourceRow int, count int, destinationParent *QModelIndex, destinationChild int) bool {
	__r0, _, _ := purego.SyscallN(addr_QFileSystemModel_moveRows_default, this.h, uintptr(sourceParent.UnsafePointer()), uintptr(sourceRow), uintptr(count), uintptr(destinationParent.UnsafePointer()), uintptr(destinationChild))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sourceParent)
	runtime.KeepAlive(destinationParent)
	return __r != 0
}

func (this *QFileSystemModel) MoveColumnsDefault(sourceParent *QModelIndex, sourceColumn int, count int, destinationParent *QModelIndex, destinationChild int) bool {
	__r0, _, _ := purego.SyscallN(addr_QFileSystemModel_moveColumns_default, this.h, uintptr(sourceParent.UnsafePointer()), uintptr(sourceColumn), uintptr(count), uintptr(destinationParent.UnsafePointer()), uintptr(destinationChild))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sourceParent)
	runtime.KeepAlive(destinationParent)
	return __r != 0
}

func (this *QFileSystemModel) BuddyDefault(index *QModelIndex) *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QFileSystemModel_buddy_default, this.h, uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return newownedQModelIndex(__r)
}

func (this *QFileSystemModel) MatchDefault(start *QModelIndex, role int, value *QVariant, hits int, flags MatchFlag) []*QModelIndex {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QFileSystemModel_match_default, this.h, uintptr(start.UnsafePointer()), uintptr(role), uintptr(value.UnsafePointer()), uintptr(hits), uintptr(flags), uintptr(unsafe.Pointer(&__n)))
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

func (this *QFileSystemModel) SpanDefault(index *QModelIndex) *QSize {
	__r, _, _ := purego.SyscallN(addr_QFileSystemModel_span_default, this.h, uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return newownedQSize(__r)
}

func (this *QFileSystemModel) MultiDataDefault(index *QModelIndex, roleDataSpan *QModelRoleDataSpan) {
	purego.SyscallN(addr_QFileSystemModel_multiData_default, this.h, uintptr(index.UnsafePointer()), uintptr(roleDataSpan.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	runtime.KeepAlive(roleDataSpan)
}

func (this *QFileSystemModel) SubmitDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QFileSystemModel_submit_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QFileSystemModel) RevertDefault() {
	purego.SyscallN(addr_QFileSystemModel_revert_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QFileSystemModel) ResetInternalDataDefault() {
	purego.SyscallN(addr_QFileSystemModel_resetInternalData_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QFileSystemModel) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QFileSystemModel_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QFileSystemModel) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QFileSystemModel_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFileSystemModel) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QFileSystemModel_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QFileSystemModel) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QFileSystemModel_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QFileSystemModel) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QFileSystemModel_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QFileSystemModel) OnRootPathChanged(fn func(string)) int64 {
	id := registerCallback(func(newPath uintptr, newPath_len int64) { fn(pgStrNoFree(newPath, newPath_len)) })
	pg_QFileSystemModel_connect_rootPathChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFileSystemModel) OnFileRenamed(fn func(string, string, string)) int64 {
	id := registerCallback(func(path uintptr, path_len int64, oldName uintptr, oldName_len int64, newName uintptr, newName_len int64) { fn(pgStrNoFree(path, path_len), pgStrNoFree(oldName, oldName_len), pgStrNoFree(newName, newName_len)) })
	pg_QFileSystemModel_connect_fileRenamed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFileSystemModel) OnDirectoryLoaded(fn func(string)) int64 {
	id := registerCallback(func(path uintptr, path_len int64) { fn(pgStrNoFree(path, path_len)) })
	pg_QFileSystemModel_connect_directoryLoaded(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFileSystemModel) OnHeaderDataChanged(fn func(Orientation, int, int)) int64 {
	id := registerCallback(func(orientation int32, first int32, last int32) { fn(Orientation(orientation), int(first), int(last)) })
	pg_QFileSystemModel_connect_headerDataChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFileSystemModel) OnLayoutChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QFileSystemModel_connect_layoutChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFileSystemModel) OnLayoutChanged2(fn func([]*QPersistentModelIndex, QAbstractItemModel__LayoutChangeHint)) int64 {
	id := registerCallback(func(parents uintptr, parents_len int64, hint int32) { fn(func() []*QPersistentModelIndex { if parents == 0 || parents_len == 0 { return nil }; __sl := unsafe.Slice((*unsafe.Pointer)(unsafe.Pointer(parents)), int(parents_len)); __o := make([]*QPersistentModelIndex, int(parents_len)); for __i := range __o { __o[__i] = UnsafeNewQPersistentModelIndex(__sl[__i]) }; return __o }(), QAbstractItemModel__LayoutChangeHint(hint)) })
	pg_QFileSystemModel_connect_layoutChanged2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFileSystemModel) OnLayoutAboutToBeChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QFileSystemModel_connect_layoutAboutToBeChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFileSystemModel) OnLayoutAboutToBeChanged2(fn func([]*QPersistentModelIndex, QAbstractItemModel__LayoutChangeHint)) int64 {
	id := registerCallback(func(parents uintptr, parents_len int64, hint int32) { fn(func() []*QPersistentModelIndex { if parents == 0 || parents_len == 0 { return nil }; __sl := unsafe.Slice((*unsafe.Pointer)(unsafe.Pointer(parents)), int(parents_len)); __o := make([]*QPersistentModelIndex, int(parents_len)); for __i := range __o { __o[__i] = UnsafeNewQPersistentModelIndex(__sl[__i]) }; return __o }(), QAbstractItemModel__LayoutChangeHint(hint)) })
	pg_QFileSystemModel_connect_layoutAboutToBeChanged2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFileSystemModel) OnRowsAboutToBeInserted(fn func(*QModelIndex, int, int)) int64 {
	id := registerCallback(func(parent uintptr, first int32, last int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(first), int(last)) })
	pg_QFileSystemModel_connect_rowsAboutToBeInserted(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFileSystemModel) OnRowsInserted(fn func(*QModelIndex, int, int)) int64 {
	id := registerCallback(func(parent uintptr, first int32, last int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(first), int(last)) })
	pg_QFileSystemModel_connect_rowsInserted(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFileSystemModel) OnRowsAboutToBeRemoved(fn func(*QModelIndex, int, int)) int64 {
	id := registerCallback(func(parent uintptr, first int32, last int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(first), int(last)) })
	pg_QFileSystemModel_connect_rowsAboutToBeRemoved(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFileSystemModel) OnRowsRemoved(fn func(*QModelIndex, int, int)) int64 {
	id := registerCallback(func(parent uintptr, first int32, last int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(first), int(last)) })
	pg_QFileSystemModel_connect_rowsRemoved(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFileSystemModel) OnColumnsAboutToBeInserted(fn func(*QModelIndex, int, int)) int64 {
	id := registerCallback(func(parent uintptr, first int32, last int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(first), int(last)) })
	pg_QFileSystemModel_connect_columnsAboutToBeInserted(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFileSystemModel) OnColumnsInserted(fn func(*QModelIndex, int, int)) int64 {
	id := registerCallback(func(parent uintptr, first int32, last int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(first), int(last)) })
	pg_QFileSystemModel_connect_columnsInserted(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFileSystemModel) OnColumnsAboutToBeRemoved(fn func(*QModelIndex, int, int)) int64 {
	id := registerCallback(func(parent uintptr, first int32, last int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(first), int(last)) })
	pg_QFileSystemModel_connect_columnsAboutToBeRemoved(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFileSystemModel) OnColumnsRemoved(fn func(*QModelIndex, int, int)) int64 {
	id := registerCallback(func(parent uintptr, first int32, last int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(first), int(last)) })
	pg_QFileSystemModel_connect_columnsRemoved(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFileSystemModel) OnModelAboutToBeReset(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QFileSystemModel_connect_modelAboutToBeReset(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFileSystemModel) OnModelReset(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QFileSystemModel_connect_modelReset(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFileSystemModel) OnRowsAboutToBeMoved(fn func(*QModelIndex, int, int, *QModelIndex, int)) int64 {
	id := registerCallback(func(sourceParent uintptr, sourceStart int32, sourceEnd int32, destinationParent uintptr, destinationRow int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(sourceParent)), int(sourceStart), int(sourceEnd), UnsafeNewQModelIndex(unsafe.Pointer(destinationParent)), int(destinationRow)) })
	pg_QFileSystemModel_connect_rowsAboutToBeMoved(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFileSystemModel) OnRowsMoved(fn func(*QModelIndex, int, int, *QModelIndex, int)) int64 {
	id := registerCallback(func(sourceParent uintptr, sourceStart int32, sourceEnd int32, destinationParent uintptr, destinationRow int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(sourceParent)), int(sourceStart), int(sourceEnd), UnsafeNewQModelIndex(unsafe.Pointer(destinationParent)), int(destinationRow)) })
	pg_QFileSystemModel_connect_rowsMoved(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFileSystemModel) OnColumnsAboutToBeMoved(fn func(*QModelIndex, int, int, *QModelIndex, int)) int64 {
	id := registerCallback(func(sourceParent uintptr, sourceStart int32, sourceEnd int32, destinationParent uintptr, destinationColumn int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(sourceParent)), int(sourceStart), int(sourceEnd), UnsafeNewQModelIndex(unsafe.Pointer(destinationParent)), int(destinationColumn)) })
	pg_QFileSystemModel_connect_columnsAboutToBeMoved(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFileSystemModel) OnColumnsMoved(fn func(*QModelIndex, int, int, *QModelIndex, int)) int64 {
	id := registerCallback(func(sourceParent uintptr, sourceStart int32, sourceEnd int32, destinationParent uintptr, destinationColumn int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(sourceParent)), int(sourceStart), int(sourceEnd), UnsafeNewQModelIndex(unsafe.Pointer(destinationParent)), int(destinationColumn)) })
	pg_QFileSystemModel_connect_columnsMoved(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFileSystemModel) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QFileSystemModel_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFileSystemModel) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QFileSystemModel_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFileSystemModel) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QFileSystemModel_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QFileSystemModel) Disconnect(handle int64) {
	pg_QFileSystemModel__disconnect(handle)
	unregisterCallback(handle)
}

func QFileSystemModel_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case -8948777187306027381:
		return 1
	case 7371339085820260189:
		return 2
	case 397329346813782521:
		return 3
	case 3369452259144514380:
		return 4
	case 8639262066473964150:
		return 5
	case 3553294593611130806:
		return 6
	case 1960960171956884010:
		return 7
	case -8256963024669584834:
		return 8
	case -8837375942098466299:
		return 9
	case 35191454575202219:
		return 10
	case -3297140021896693688:
		return 11
	case 1703382833545304812:
		return 12
	case 6912441694453382161:
		return 13
	case 5464467312394730418:
		return 14
	case 5175297091024974749:
		return 15
	case -9049994684305682212:
		return 16
	case -7784642176186821205:
		return 17
	case -5409927581812374498:
		return 18
	case 1237236456312824991:
		return 19
	case 4852200306590054850:
		return 20
	case 8914190500520542409:
		return 21
	case 731539533644753052:
		return 22
	case 3421601386088570762:
		return 23
	case -2775729825962395603:
		return 24
	case -8973293846230792237:
		return 25
	case 5104072549182222656:
		return 26
	case 2078548737056540420:
		return 27
	case -6074819356947405415:
		return 28
	case -6321007591388309905:
		return 29
	case 6112860112248786499:
		return 30
	case -4341501217981328394:
		return 31
	case -8395342916587040543:
		return 32
	case 731311060944096642:
		return 33
	case -3788331537442685179:
		return 34
	case -1522951607536117149:
		return 35
	case -4510675770784051183:
		return 36
	case 5217623514630858297:
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

func QFileSystemModel_enable(h uintptr, idx int32) { pg_QFileSystemModel__set_override(h, idx) }

func (this *QFileSystemModel) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QFileSystemModel) OnIndex(fn func(int, int, *QModelIndex) *QModelIndex) {
	registerVirt(this.UnsafeGoID(), -8948777187306027381, func(row int32, column int32, parent uintptr) uintptr { return uintptr(fn(int(row), int(column), UnsafeNewQModelIndex(unsafe.Pointer(parent))).UnsafePointer()) }) // slot "index"
}

func (this *QFileSystemModel) OnParent(fn func(*QModelIndex) *QModelIndex) {
	registerVirt(this.UnsafeGoID(), 7371339085820260189, func(child uintptr) uintptr { return uintptr(fn(UnsafeNewQModelIndex(unsafe.Pointer(child))).UnsafePointer()) }) // slot "parent"
}

func (this *QFileSystemModel) OnSibling(fn func(int, int, *QModelIndex) *QModelIndex) {
	registerVirt(this.UnsafeGoID(), 397329346813782521, func(row int32, column int32, idx uintptr) uintptr { return uintptr(fn(int(row), int(column), UnsafeNewQModelIndex(unsafe.Pointer(idx))).UnsafePointer()) }) // slot "sibling"
}

func (this *QFileSystemModel) OnHasChildren(fn func(*QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), 3369452259144514380, func(parent uintptr) int32 { return b2i(fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)))) }) // slot "hasChildren"
}

func (this *QFileSystemModel) OnCanFetchMore(fn func(*QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), 8639262066473964150, func(parent uintptr) int32 { return b2i(fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)))) }) // slot "canFetchMore"
}

func (this *QFileSystemModel) OnFetchMore(fn func(*QModelIndex)) {
	registerVirt(this.UnsafeGoID(), 3553294593611130806, func(parent uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent))) }) // slot "fetchMore"
}

func (this *QFileSystemModel) OnRowCount(fn func(*QModelIndex) int) {
	registerVirt(this.UnsafeGoID(), 1960960171956884010, func(parent uintptr) int32 { return int32(fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)))) }) // slot "rowCount"
}

func (this *QFileSystemModel) OnColumnCount(fn func(*QModelIndex) int) {
	registerVirt(this.UnsafeGoID(), -8256963024669584834, func(parent uintptr) int32 { return int32(fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)))) }) // slot "columnCount"
}

func (this *QFileSystemModel) OnData(fn func(*QModelIndex, int) *QVariant) {
	registerVirt(this.UnsafeGoID(), -8837375942098466299, func(index uintptr, role int32) uintptr { return uintptr(fn(UnsafeNewQModelIndex(unsafe.Pointer(index)), int(role)).UnsafePointer()) }) // slot "data"
}

func (this *QFileSystemModel) OnSetData(fn func(*QModelIndex, *QVariant, int) bool) {
	registerVirt(this.UnsafeGoID(), 35191454575202219, func(index uintptr, value uintptr, role int32) int32 { return b2i(fn(UnsafeNewQModelIndex(unsafe.Pointer(index)), UnsafeNewQVariant(unsafe.Pointer(value)), int(role))) }) // slot "setData"
}

func (this *QFileSystemModel) OnHeaderData(fn func(int, Orientation, int) *QVariant) {
	registerVirt(this.UnsafeGoID(), -3297140021896693688, func(section int32, orientation int32, role int32) uintptr { return uintptr(fn(int(section), Orientation(orientation), int(role)).UnsafePointer()) }) // slot "headerData"
}

func (this *QFileSystemModel) OnFlags(fn func(*QModelIndex) ItemFlag) {
	registerVirt(this.UnsafeGoID(), 1703382833545304812, func(index uintptr) int32 { return int32(fn(UnsafeNewQModelIndex(unsafe.Pointer(index)))) }) // slot "flags"
}

func (this *QFileSystemModel) OnSort(fn func(int, SortOrder)) {
	registerVirt(this.UnsafeGoID(), 6912441694453382161, func(column int32, order int32) { fn(int(column), SortOrder(order)) }) // slot "sort"
}

func (this *QFileSystemModel) OnMimeTypes(fn func() []string) {
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

func (this *QFileSystemModel) OnMimeData(fn func([]*QModelIndex) *QMimeData) {
	registerVirt(this.UnsafeGoID(), 5175297091024974749, func(indexes uintptr, indexes_len int64) uintptr { return uintptr(fn(func() []*QModelIndex { if indexes == 0 || indexes_len == 0 { return nil }; __sl := unsafe.Slice((*unsafe.Pointer)(unsafe.Pointer(indexes)), int(indexes_len)); __o := make([]*QModelIndex, int(indexes_len)); for __i := range __o { __o[__i] = UnsafeNewQModelIndex(__sl[__i]) }; return __o }()).UnsafePointer()) }) // slot "mimeData"
}

func (this *QFileSystemModel) OnDropMimeData(fn func(*QMimeData, DropAction, int, int, *QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), -9049994684305682212, func(data uintptr, action int32, row int32, column int32, parent uintptr) int32 { return b2i(fn(UnsafeNewQMimeData(unsafe.Pointer(data)), DropAction(action), int(row), int(column), UnsafeNewQModelIndex(unsafe.Pointer(parent)))) }) // slot "dropMimeData"
}

func (this *QFileSystemModel) OnSupportedDropActions(fn func() DropAction) {
	registerVirt(this.UnsafeGoID(), -7784642176186821205, func() int32 { return int32(fn()) }) // slot "supportedDropActions"
}

func (this *QFileSystemModel) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QFileSystemModel) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QFileSystemModel) OnSetHeaderData(fn func(int, Orientation, *QVariant, int) bool) {
	registerVirt(this.UnsafeGoID(), 4852200306590054850, func(section int32, orientation int32, value uintptr, role int32) int32 { return b2i(fn(int(section), Orientation(orientation), UnsafeNewQVariant(unsafe.Pointer(value)), int(role))) }) // slot "setHeaderData"
}

func (this *QFileSystemModel) OnClearItemData(fn func(*QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), 8914190500520542409, func(index uintptr) int32 { return b2i(fn(UnsafeNewQModelIndex(unsafe.Pointer(index)))) }) // slot "clearItemData"
}

func (this *QFileSystemModel) OnCanDropMimeData(fn func(*QMimeData, DropAction, int, int, *QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), 731539533644753052, func(data uintptr, action int32, row int32, column int32, parent uintptr) int32 { return b2i(fn(UnsafeNewQMimeData(unsafe.Pointer(data)), DropAction(action), int(row), int(column), UnsafeNewQModelIndex(unsafe.Pointer(parent)))) }) // slot "canDropMimeData"
}

func (this *QFileSystemModel) OnSupportedDragActions(fn func() DropAction) {
	registerVirt(this.UnsafeGoID(), 3421601386088570762, func() int32 { return int32(fn()) }) // slot "supportedDragActions"
}

func (this *QFileSystemModel) OnInsertRows(fn func(int, int, *QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), -2775729825962395603, func(row int32, count int32, parent uintptr) int32 { return b2i(fn(int(row), int(count), UnsafeNewQModelIndex(unsafe.Pointer(parent)))) }) // slot "insertRows"
}

func (this *QFileSystemModel) OnInsertColumns(fn func(int, int, *QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), -8973293846230792237, func(column int32, count int32, parent uintptr) int32 { return b2i(fn(int(column), int(count), UnsafeNewQModelIndex(unsafe.Pointer(parent)))) }) // slot "insertColumns"
}

func (this *QFileSystemModel) OnRemoveRows(fn func(int, int, *QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), 5104072549182222656, func(row int32, count int32, parent uintptr) int32 { return b2i(fn(int(row), int(count), UnsafeNewQModelIndex(unsafe.Pointer(parent)))) }) // slot "removeRows"
}

func (this *QFileSystemModel) OnRemoveColumns(fn func(int, int, *QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), 2078548737056540420, func(column int32, count int32, parent uintptr) int32 { return b2i(fn(int(column), int(count), UnsafeNewQModelIndex(unsafe.Pointer(parent)))) }) // slot "removeColumns"
}

func (this *QFileSystemModel) OnMoveRows(fn func(*QModelIndex, int, int, *QModelIndex, int) bool) {
	registerVirt(this.UnsafeGoID(), -6074819356947405415, func(sourceParent uintptr, sourceRow int32, count int32, destinationParent uintptr, destinationChild int32) int32 { return b2i(fn(UnsafeNewQModelIndex(unsafe.Pointer(sourceParent)), int(sourceRow), int(count), UnsafeNewQModelIndex(unsafe.Pointer(destinationParent)), int(destinationChild))) }) // slot "moveRows"
}

func (this *QFileSystemModel) OnMoveColumns(fn func(*QModelIndex, int, int, *QModelIndex, int) bool) {
	registerVirt(this.UnsafeGoID(), -6321007591388309905, func(sourceParent uintptr, sourceColumn int32, count int32, destinationParent uintptr, destinationChild int32) int32 { return b2i(fn(UnsafeNewQModelIndex(unsafe.Pointer(sourceParent)), int(sourceColumn), int(count), UnsafeNewQModelIndex(unsafe.Pointer(destinationParent)), int(destinationChild))) }) // slot "moveColumns"
}

func (this *QFileSystemModel) OnBuddy(fn func(*QModelIndex) *QModelIndex) {
	registerVirt(this.UnsafeGoID(), 6112860112248786499, func(index uintptr) uintptr { return uintptr(fn(UnsafeNewQModelIndex(unsafe.Pointer(index))).UnsafePointer()) }) // slot "buddy"
}

func (this *QFileSystemModel) OnMatch(fn func(*QModelIndex, int, *QVariant, int, MatchFlag) []*QModelIndex) {
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

func (this *QFileSystemModel) OnSpan(fn func(*QModelIndex) *QSize) {
	registerVirt(this.UnsafeGoID(), -8395342916587040543, func(index uintptr) uintptr { return uintptr(fn(UnsafeNewQModelIndex(unsafe.Pointer(index))).UnsafePointer()) }) // slot "span"
}

func (this *QFileSystemModel) OnMultiData(fn func(*QModelIndex, *QModelRoleDataSpan)) {
	registerVirt(this.UnsafeGoID(), 731311060944096642, func(index uintptr, roleDataSpan uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(index)), UnsafeNewQModelRoleDataSpan(unsafe.Pointer(roleDataSpan))) }) // slot "multiData"
}

func (this *QFileSystemModel) OnSubmit(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -3788331537442685179, func() int32 { return b2i(fn()) }) // slot "submit"
}

func (this *QFileSystemModel) OnRevert(fn func()) {
	registerVirt(this.UnsafeGoID(), -1522951607536117149, func() { fn() }) // slot "revert"
}

func (this *QFileSystemModel) OnResetInternalData(fn func()) {
	registerVirt(this.UnsafeGoID(), -4510675770784051183, func() { fn() }) // slot "resetInternalData"
}

func (this *QFileSystemModel) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QFileSystemModel) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QFileSystemModel) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QFileSystemModel) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QFileSystemModel) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QFileSystemModel) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QFileSystemModel_delete(this.h)
	}
	this.UnsafeRelease()
}
