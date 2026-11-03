package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QSortFilterProxyModel_chain uintptr
	addr_QSortFilterProxyModel_new uintptr
	addr_QSortFilterProxyModel_new2 uintptr
	addr_QSortFilterProxyModel_tr uintptr
	addr_QSortFilterProxyModel_filterRegularExpression uintptr
	addr_QSortFilterProxyModel_filterKeyColumn uintptr
	addr_QSortFilterProxyModel_setFilterKeyColumn uintptr
	addr_QSortFilterProxyModel_filterCaseSensitivity uintptr
	addr_QSortFilterProxyModel_setFilterCaseSensitivity uintptr
	addr_QSortFilterProxyModel_sortCaseSensitivity uintptr
	addr_QSortFilterProxyModel_setSortCaseSensitivity uintptr
	addr_QSortFilterProxyModel_isSortLocaleAware uintptr
	addr_QSortFilterProxyModel_setSortLocaleAware uintptr
	addr_QSortFilterProxyModel_sortColumn uintptr
	addr_QSortFilterProxyModel_sortOrder uintptr
	addr_QSortFilterProxyModel_dynamicSortFilter uintptr
	addr_QSortFilterProxyModel_setDynamicSortFilter uintptr
	addr_QSortFilterProxyModel_sortRole uintptr
	addr_QSortFilterProxyModel_setSortRole uintptr
	addr_QSortFilterProxyModel_filterRole uintptr
	addr_QSortFilterProxyModel_setFilterRole uintptr
	addr_QSortFilterProxyModel_isRecursiveFilteringEnabled uintptr
	addr_QSortFilterProxyModel_setRecursiveFilteringEnabled uintptr
	addr_QSortFilterProxyModel_autoAcceptChildRows uintptr
	addr_QSortFilterProxyModel_setAutoAcceptChildRows uintptr
	addr_QSortFilterProxyModel_setFilterRegularExpression uintptr
	addr_QSortFilterProxyModel_setFilterRegularExpression2 uintptr
	addr_QSortFilterProxyModel_setFilterWildcard uintptr
	addr_QSortFilterProxyModel_setFilterFixedString uintptr
	addr_QSortFilterProxyModel_invalidate uintptr
	addr_QSortFilterProxyModel_beginFilterChange uintptr
	addr_QSortFilterProxyModel_endFilterChange uintptr
	addr_QSortFilterProxyModel_invalidateFilter uintptr
	addr_QSortFilterProxyModel_invalidateRowsFilter uintptr
	addr_QSortFilterProxyModel_invalidateColumnsFilter uintptr
	addr_QSortFilterProxyModel_dynamicSortFilterChanged uintptr
	addr_QSortFilterProxyModel_filterCaseSensitivityChanged uintptr
	addr_QSortFilterProxyModel_sortCaseSensitivityChanged uintptr
	addr_QSortFilterProxyModel_sortLocaleAwareChanged uintptr
	addr_QSortFilterProxyModel_sortRoleChanged uintptr
	addr_QSortFilterProxyModel_filterRoleChanged uintptr
	addr_QSortFilterProxyModel_recursiveFilteringEnabledChanged uintptr
	addr_QSortFilterProxyModel_autoAcceptChildRowsChanged uintptr
	addr_QSortFilterProxyModel_createIndex uintptr
	addr_QSortFilterProxyModel_createIndex2 uintptr
	addr_QSortFilterProxyModel_encodeData uintptr
	addr_QSortFilterProxyModel_decodeData uintptr
	addr_QSortFilterProxyModel_beginInsertRows uintptr
	addr_QSortFilterProxyModel_endInsertRows uintptr
	addr_QSortFilterProxyModel_beginRemoveRows uintptr
	addr_QSortFilterProxyModel_endRemoveRows uintptr
	addr_QSortFilterProxyModel_beginMoveRows uintptr
	addr_QSortFilterProxyModel_endMoveRows uintptr
	addr_QSortFilterProxyModel_beginInsertColumns uintptr
	addr_QSortFilterProxyModel_endInsertColumns uintptr
	addr_QSortFilterProxyModel_beginRemoveColumns uintptr
	addr_QSortFilterProxyModel_endRemoveColumns uintptr
	addr_QSortFilterProxyModel_beginMoveColumns uintptr
	addr_QSortFilterProxyModel_endMoveColumns uintptr
	addr_QSortFilterProxyModel_beginResetModel uintptr
	addr_QSortFilterProxyModel_endResetModel uintptr
	addr_QSortFilterProxyModel_changePersistentIndex uintptr
	addr_QSortFilterProxyModel_changePersistentIndexList uintptr
	addr_QSortFilterProxyModel_persistentIndexList uintptr
	addr_QSortFilterProxyModel_connect uintptr
	addr_QSortFilterProxyModel_connect2 uintptr
	addr_QSortFilterProxyModel_sender uintptr
	addr_QSortFilterProxyModel_senderSignalIndex uintptr
	addr_QSortFilterProxyModel_receivers uintptr
	addr_QSortFilterProxyModel_isSignalConnected uintptr
	addr_QSortFilterProxyModel_tr2 uintptr
	addr_QSortFilterProxyModel_tr3 uintptr
	addr_QSortFilterProxyModel_endFilterChange2 uintptr
	addr_QSortFilterProxyModel_connect3 uintptr
	addr_QSortFilterProxyModel_connect4 uintptr
	addr_QSortFilterProxyModel_metaObject_default uintptr
	addr_QSortFilterProxyModel_setSourceModel_default uintptr
	addr_QSortFilterProxyModel_mapToSource_default uintptr
	addr_QSortFilterProxyModel_mapFromSource_default uintptr
	addr_QSortFilterProxyModel_mapSelectionToSource_default uintptr
	addr_QSortFilterProxyModel_mapSelectionFromSource_default uintptr
	addr_QSortFilterProxyModel_filterAcceptsRow_default uintptr
	addr_QSortFilterProxyModel_filterAcceptsColumn_default uintptr
	addr_QSortFilterProxyModel_lessThan_default uintptr
	addr_QSortFilterProxyModel_index_default uintptr
	addr_QSortFilterProxyModel_parent_default uintptr
	addr_QSortFilterProxyModel_sibling_default uintptr
	addr_QSortFilterProxyModel_rowCount_default uintptr
	addr_QSortFilterProxyModel_columnCount_default uintptr
	addr_QSortFilterProxyModel_hasChildren_default uintptr
	addr_QSortFilterProxyModel_data_default uintptr
	addr_QSortFilterProxyModel_setData_default uintptr
	addr_QSortFilterProxyModel_headerData_default uintptr
	addr_QSortFilterProxyModel_setHeaderData_default uintptr
	addr_QSortFilterProxyModel_mimeData_default uintptr
	addr_QSortFilterProxyModel_dropMimeData_default uintptr
	addr_QSortFilterProxyModel_insertRows_default uintptr
	addr_QSortFilterProxyModel_insertColumns_default uintptr
	addr_QSortFilterProxyModel_removeRows_default uintptr
	addr_QSortFilterProxyModel_removeColumns_default uintptr
	addr_QSortFilterProxyModel_fetchMore_default uintptr
	addr_QSortFilterProxyModel_canFetchMore_default uintptr
	addr_QSortFilterProxyModel_flags_default uintptr
	addr_QSortFilterProxyModel_buddy_default uintptr
	addr_QSortFilterProxyModel_match_default uintptr
	addr_QSortFilterProxyModel_span_default uintptr
	addr_QSortFilterProxyModel_sort_default uintptr
	addr_QSortFilterProxyModel_mimeTypes_default uintptr
	addr_QSortFilterProxyModel_supportedDropActions_default uintptr
	addr_QSortFilterProxyModel_submit_default uintptr
	addr_QSortFilterProxyModel_revert_default uintptr
	addr_QSortFilterProxyModel_clearItemData_default uintptr
	addr_QSortFilterProxyModel_canDropMimeData_default uintptr
	addr_QSortFilterProxyModel_supportedDragActions_default uintptr
	addr_QSortFilterProxyModel_moveRows_default uintptr
	addr_QSortFilterProxyModel_moveColumns_default uintptr
	addr_QSortFilterProxyModel_multiData_default uintptr
	addr_QSortFilterProxyModel_resetInternalData_default uintptr
	addr_QSortFilterProxyModel_event_default uintptr
	addr_QSortFilterProxyModel_eventFilter_default uintptr
	addr_QSortFilterProxyModel_timerEvent_default uintptr
	addr_QSortFilterProxyModel_childEvent_default uintptr
	addr_QSortFilterProxyModel_customEvent_default uintptr
	addr_QSortFilterProxyModel_connectNotify_default uintptr
	addr_QSortFilterProxyModel_disconnectNotify_default uintptr
	pg_QSortFilterProxyModel_connect_dynamicSortFilterChanged func(uintptr, int64)
	pg_QSortFilterProxyModel_connect_filterCaseSensitivityChanged func(uintptr, int64)
	pg_QSortFilterProxyModel_connect_sortCaseSensitivityChanged func(uintptr, int64)
	pg_QSortFilterProxyModel_connect_sortLocaleAwareChanged func(uintptr, int64)
	pg_QSortFilterProxyModel_connect_sortRoleChanged func(uintptr, int64)
	pg_QSortFilterProxyModel_connect_filterRoleChanged func(uintptr, int64)
	pg_QSortFilterProxyModel_connect_recursiveFilteringEnabledChanged func(uintptr, int64)
	pg_QSortFilterProxyModel_connect_autoAcceptChildRowsChanged func(uintptr, int64)
	pg_QSortFilterProxyModel_connect_sourceModelChanged func(uintptr, int64)
	pg_QSortFilterProxyModel_connect_headerDataChanged func(uintptr, int64)
	pg_QSortFilterProxyModel_connect_layoutChanged func(uintptr, int64)
	pg_QSortFilterProxyModel_connect_layoutChanged2 func(uintptr, int64)
	pg_QSortFilterProxyModel_connect_layoutAboutToBeChanged func(uintptr, int64)
	pg_QSortFilterProxyModel_connect_layoutAboutToBeChanged2 func(uintptr, int64)
	pg_QSortFilterProxyModel_connect_rowsAboutToBeInserted func(uintptr, int64)
	pg_QSortFilterProxyModel_connect_rowsInserted func(uintptr, int64)
	pg_QSortFilterProxyModel_connect_rowsAboutToBeRemoved func(uintptr, int64)
	pg_QSortFilterProxyModel_connect_rowsRemoved func(uintptr, int64)
	pg_QSortFilterProxyModel_connect_columnsAboutToBeInserted func(uintptr, int64)
	pg_QSortFilterProxyModel_connect_columnsInserted func(uintptr, int64)
	pg_QSortFilterProxyModel_connect_columnsAboutToBeRemoved func(uintptr, int64)
	pg_QSortFilterProxyModel_connect_columnsRemoved func(uintptr, int64)
	pg_QSortFilterProxyModel_connect_modelAboutToBeReset func(uintptr, int64)
	pg_QSortFilterProxyModel_connect_modelReset func(uintptr, int64)
	pg_QSortFilterProxyModel_connect_rowsAboutToBeMoved func(uintptr, int64)
	pg_QSortFilterProxyModel_connect_rowsMoved func(uintptr, int64)
	pg_QSortFilterProxyModel_connect_columnsAboutToBeMoved func(uintptr, int64)
	pg_QSortFilterProxyModel_connect_columnsMoved func(uintptr, int64)
	pg_QSortFilterProxyModel_connect_destroyed func(uintptr, int64)
	pg_QSortFilterProxyModel_connect_destroyed2 func(uintptr, int64)
	pg_QSortFilterProxyModel_connect_objectNameChanged func(uintptr, int64)
	pg_QSortFilterProxyModel__disconnect func(int64)
	pg_QSortFilterProxyModel__set_override func(uintptr, int32)
	pg_QSortFilterProxyModel_delete func(uintptr)
)

func init() {
	symaddr(&addr_QSortFilterProxyModel_chain, "QSortFilterProxyModel_chain")
	symaddr(&addr_QSortFilterProxyModel_new, "QSortFilterProxyModel_new")
	symaddr(&addr_QSortFilterProxyModel_new2, "QSortFilterProxyModel_new2")
	symaddr(&addr_QSortFilterProxyModel_tr, "QSortFilterProxyModel_tr")
	symaddr(&addr_QSortFilterProxyModel_filterRegularExpression, "QSortFilterProxyModel_filterRegularExpression")
	symaddr(&addr_QSortFilterProxyModel_filterKeyColumn, "QSortFilterProxyModel_filterKeyColumn")
	symaddr(&addr_QSortFilterProxyModel_setFilterKeyColumn, "QSortFilterProxyModel_setFilterKeyColumn")
	symaddr(&addr_QSortFilterProxyModel_filterCaseSensitivity, "QSortFilterProxyModel_filterCaseSensitivity")
	symaddr(&addr_QSortFilterProxyModel_setFilterCaseSensitivity, "QSortFilterProxyModel_setFilterCaseSensitivity")
	symaddr(&addr_QSortFilterProxyModel_sortCaseSensitivity, "QSortFilterProxyModel_sortCaseSensitivity")
	symaddr(&addr_QSortFilterProxyModel_setSortCaseSensitivity, "QSortFilterProxyModel_setSortCaseSensitivity")
	symaddr(&addr_QSortFilterProxyModel_isSortLocaleAware, "QSortFilterProxyModel_isSortLocaleAware")
	symaddr(&addr_QSortFilterProxyModel_setSortLocaleAware, "QSortFilterProxyModel_setSortLocaleAware")
	symaddr(&addr_QSortFilterProxyModel_sortColumn, "QSortFilterProxyModel_sortColumn")
	symaddr(&addr_QSortFilterProxyModel_sortOrder, "QSortFilterProxyModel_sortOrder")
	symaddr(&addr_QSortFilterProxyModel_dynamicSortFilter, "QSortFilterProxyModel_dynamicSortFilter")
	symaddr(&addr_QSortFilterProxyModel_setDynamicSortFilter, "QSortFilterProxyModel_setDynamicSortFilter")
	symaddr(&addr_QSortFilterProxyModel_sortRole, "QSortFilterProxyModel_sortRole")
	symaddr(&addr_QSortFilterProxyModel_setSortRole, "QSortFilterProxyModel_setSortRole")
	symaddr(&addr_QSortFilterProxyModel_filterRole, "QSortFilterProxyModel_filterRole")
	symaddr(&addr_QSortFilterProxyModel_setFilterRole, "QSortFilterProxyModel_setFilterRole")
	symaddr(&addr_QSortFilterProxyModel_isRecursiveFilteringEnabled, "QSortFilterProxyModel_isRecursiveFilteringEnabled")
	symaddr(&addr_QSortFilterProxyModel_setRecursiveFilteringEnabled, "QSortFilterProxyModel_setRecursiveFilteringEnabled")
	symaddr(&addr_QSortFilterProxyModel_autoAcceptChildRows, "QSortFilterProxyModel_autoAcceptChildRows")
	symaddr(&addr_QSortFilterProxyModel_setAutoAcceptChildRows, "QSortFilterProxyModel_setAutoAcceptChildRows")
	symaddr(&addr_QSortFilterProxyModel_setFilterRegularExpression, "QSortFilterProxyModel_setFilterRegularExpression")
	symaddr(&addr_QSortFilterProxyModel_setFilterRegularExpression2, "QSortFilterProxyModel_setFilterRegularExpression2")
	symaddr(&addr_QSortFilterProxyModel_setFilterWildcard, "QSortFilterProxyModel_setFilterWildcard")
	symaddr(&addr_QSortFilterProxyModel_setFilterFixedString, "QSortFilterProxyModel_setFilterFixedString")
	symaddr(&addr_QSortFilterProxyModel_invalidate, "QSortFilterProxyModel_invalidate")
	symaddr(&addr_QSortFilterProxyModel_beginFilterChange, "QSortFilterProxyModel_beginFilterChange")
	symaddr(&addr_QSortFilterProxyModel_endFilterChange, "QSortFilterProxyModel_endFilterChange")
	symaddr(&addr_QSortFilterProxyModel_invalidateFilter, "QSortFilterProxyModel_invalidateFilter")
	symaddr(&addr_QSortFilterProxyModel_invalidateRowsFilter, "QSortFilterProxyModel_invalidateRowsFilter")
	symaddr(&addr_QSortFilterProxyModel_invalidateColumnsFilter, "QSortFilterProxyModel_invalidateColumnsFilter")
	symaddr(&addr_QSortFilterProxyModel_dynamicSortFilterChanged, "QSortFilterProxyModel_dynamicSortFilterChanged")
	symaddr(&addr_QSortFilterProxyModel_filterCaseSensitivityChanged, "QSortFilterProxyModel_filterCaseSensitivityChanged")
	symaddr(&addr_QSortFilterProxyModel_sortCaseSensitivityChanged, "QSortFilterProxyModel_sortCaseSensitivityChanged")
	symaddr(&addr_QSortFilterProxyModel_sortLocaleAwareChanged, "QSortFilterProxyModel_sortLocaleAwareChanged")
	symaddr(&addr_QSortFilterProxyModel_sortRoleChanged, "QSortFilterProxyModel_sortRoleChanged")
	symaddr(&addr_QSortFilterProxyModel_filterRoleChanged, "QSortFilterProxyModel_filterRoleChanged")
	symaddr(&addr_QSortFilterProxyModel_recursiveFilteringEnabledChanged, "QSortFilterProxyModel_recursiveFilteringEnabledChanged")
	symaddr(&addr_QSortFilterProxyModel_autoAcceptChildRowsChanged, "QSortFilterProxyModel_autoAcceptChildRowsChanged")
	symaddr(&addr_QSortFilterProxyModel_createIndex, "QSortFilterProxyModel_createIndex")
	symaddr(&addr_QSortFilterProxyModel_createIndex2, "QSortFilterProxyModel_createIndex2")
	symaddr(&addr_QSortFilterProxyModel_encodeData, "QSortFilterProxyModel_encodeData")
	symaddr(&addr_QSortFilterProxyModel_decodeData, "QSortFilterProxyModel_decodeData")
	symaddr(&addr_QSortFilterProxyModel_beginInsertRows, "QSortFilterProxyModel_beginInsertRows")
	symaddr(&addr_QSortFilterProxyModel_endInsertRows, "QSortFilterProxyModel_endInsertRows")
	symaddr(&addr_QSortFilterProxyModel_beginRemoveRows, "QSortFilterProxyModel_beginRemoveRows")
	symaddr(&addr_QSortFilterProxyModel_endRemoveRows, "QSortFilterProxyModel_endRemoveRows")
	symaddr(&addr_QSortFilterProxyModel_beginMoveRows, "QSortFilterProxyModel_beginMoveRows")
	symaddr(&addr_QSortFilterProxyModel_endMoveRows, "QSortFilterProxyModel_endMoveRows")
	symaddr(&addr_QSortFilterProxyModel_beginInsertColumns, "QSortFilterProxyModel_beginInsertColumns")
	symaddr(&addr_QSortFilterProxyModel_endInsertColumns, "QSortFilterProxyModel_endInsertColumns")
	symaddr(&addr_QSortFilterProxyModel_beginRemoveColumns, "QSortFilterProxyModel_beginRemoveColumns")
	symaddr(&addr_QSortFilterProxyModel_endRemoveColumns, "QSortFilterProxyModel_endRemoveColumns")
	symaddr(&addr_QSortFilterProxyModel_beginMoveColumns, "QSortFilterProxyModel_beginMoveColumns")
	symaddr(&addr_QSortFilterProxyModel_endMoveColumns, "QSortFilterProxyModel_endMoveColumns")
	symaddr(&addr_QSortFilterProxyModel_beginResetModel, "QSortFilterProxyModel_beginResetModel")
	symaddr(&addr_QSortFilterProxyModel_endResetModel, "QSortFilterProxyModel_endResetModel")
	symaddr(&addr_QSortFilterProxyModel_changePersistentIndex, "QSortFilterProxyModel_changePersistentIndex")
	symaddr(&addr_QSortFilterProxyModel_changePersistentIndexList, "QSortFilterProxyModel_changePersistentIndexList")
	symaddr(&addr_QSortFilterProxyModel_persistentIndexList, "QSortFilterProxyModel_persistentIndexList")
	symaddr(&addr_QSortFilterProxyModel_connect, "QSortFilterProxyModel_connect")
	symaddr(&addr_QSortFilterProxyModel_connect2, "QSortFilterProxyModel_connect2")
	symaddr(&addr_QSortFilterProxyModel_sender, "QSortFilterProxyModel_sender")
	symaddr(&addr_QSortFilterProxyModel_senderSignalIndex, "QSortFilterProxyModel_senderSignalIndex")
	symaddr(&addr_QSortFilterProxyModel_receivers, "QSortFilterProxyModel_receivers")
	symaddr(&addr_QSortFilterProxyModel_isSignalConnected, "QSortFilterProxyModel_isSignalConnected")
	symaddr(&addr_QSortFilterProxyModel_tr2, "QSortFilterProxyModel_tr2")
	symaddr(&addr_QSortFilterProxyModel_tr3, "QSortFilterProxyModel_tr3")
	symaddr(&addr_QSortFilterProxyModel_endFilterChange2, "QSortFilterProxyModel_endFilterChange2")
	symaddr(&addr_QSortFilterProxyModel_connect3, "QSortFilterProxyModel_connect3")
	symaddr(&addr_QSortFilterProxyModel_connect4, "QSortFilterProxyModel_connect4")
	symaddr(&addr_QSortFilterProxyModel_metaObject_default, "QSortFilterProxyModel_metaObject_default")
	symaddr(&addr_QSortFilterProxyModel_setSourceModel_default, "QSortFilterProxyModel_setSourceModel_default")
	symaddr(&addr_QSortFilterProxyModel_mapToSource_default, "QSortFilterProxyModel_mapToSource_default")
	symaddr(&addr_QSortFilterProxyModel_mapFromSource_default, "QSortFilterProxyModel_mapFromSource_default")
	symaddr(&addr_QSortFilterProxyModel_mapSelectionToSource_default, "QSortFilterProxyModel_mapSelectionToSource_default")
	symaddr(&addr_QSortFilterProxyModel_mapSelectionFromSource_default, "QSortFilterProxyModel_mapSelectionFromSource_default")
	symaddr(&addr_QSortFilterProxyModel_filterAcceptsRow_default, "QSortFilterProxyModel_filterAcceptsRow_default")
	symaddr(&addr_QSortFilterProxyModel_filterAcceptsColumn_default, "QSortFilterProxyModel_filterAcceptsColumn_default")
	symaddr(&addr_QSortFilterProxyModel_lessThan_default, "QSortFilterProxyModel_lessThan_default")
	symaddr(&addr_QSortFilterProxyModel_index_default, "QSortFilterProxyModel_index_default")
	symaddr(&addr_QSortFilterProxyModel_parent_default, "QSortFilterProxyModel_parent_default")
	symaddr(&addr_QSortFilterProxyModel_sibling_default, "QSortFilterProxyModel_sibling_default")
	symaddr(&addr_QSortFilterProxyModel_rowCount_default, "QSortFilterProxyModel_rowCount_default")
	symaddr(&addr_QSortFilterProxyModel_columnCount_default, "QSortFilterProxyModel_columnCount_default")
	symaddr(&addr_QSortFilterProxyModel_hasChildren_default, "QSortFilterProxyModel_hasChildren_default")
	symaddr(&addr_QSortFilterProxyModel_data_default, "QSortFilterProxyModel_data_default")
	symaddr(&addr_QSortFilterProxyModel_setData_default, "QSortFilterProxyModel_setData_default")
	symaddr(&addr_QSortFilterProxyModel_headerData_default, "QSortFilterProxyModel_headerData_default")
	symaddr(&addr_QSortFilterProxyModel_setHeaderData_default, "QSortFilterProxyModel_setHeaderData_default")
	symaddr(&addr_QSortFilterProxyModel_mimeData_default, "QSortFilterProxyModel_mimeData_default")
	symaddr(&addr_QSortFilterProxyModel_dropMimeData_default, "QSortFilterProxyModel_dropMimeData_default")
	symaddr(&addr_QSortFilterProxyModel_insertRows_default, "QSortFilterProxyModel_insertRows_default")
	symaddr(&addr_QSortFilterProxyModel_insertColumns_default, "QSortFilterProxyModel_insertColumns_default")
	symaddr(&addr_QSortFilterProxyModel_removeRows_default, "QSortFilterProxyModel_removeRows_default")
	symaddr(&addr_QSortFilterProxyModel_removeColumns_default, "QSortFilterProxyModel_removeColumns_default")
	symaddr(&addr_QSortFilterProxyModel_fetchMore_default, "QSortFilterProxyModel_fetchMore_default")
	symaddr(&addr_QSortFilterProxyModel_canFetchMore_default, "QSortFilterProxyModel_canFetchMore_default")
	symaddr(&addr_QSortFilterProxyModel_flags_default, "QSortFilterProxyModel_flags_default")
	symaddr(&addr_QSortFilterProxyModel_buddy_default, "QSortFilterProxyModel_buddy_default")
	symaddr(&addr_QSortFilterProxyModel_match_default, "QSortFilterProxyModel_match_default")
	symaddr(&addr_QSortFilterProxyModel_span_default, "QSortFilterProxyModel_span_default")
	symaddr(&addr_QSortFilterProxyModel_sort_default, "QSortFilterProxyModel_sort_default")
	symaddr(&addr_QSortFilterProxyModel_mimeTypes_default, "QSortFilterProxyModel_mimeTypes_default")
	symaddr(&addr_QSortFilterProxyModel_supportedDropActions_default, "QSortFilterProxyModel_supportedDropActions_default")
	symaddr(&addr_QSortFilterProxyModel_submit_default, "QSortFilterProxyModel_submit_default")
	symaddr(&addr_QSortFilterProxyModel_revert_default, "QSortFilterProxyModel_revert_default")
	symaddr(&addr_QSortFilterProxyModel_clearItemData_default, "QSortFilterProxyModel_clearItemData_default")
	symaddr(&addr_QSortFilterProxyModel_canDropMimeData_default, "QSortFilterProxyModel_canDropMimeData_default")
	symaddr(&addr_QSortFilterProxyModel_supportedDragActions_default, "QSortFilterProxyModel_supportedDragActions_default")
	symaddr(&addr_QSortFilterProxyModel_moveRows_default, "QSortFilterProxyModel_moveRows_default")
	symaddr(&addr_QSortFilterProxyModel_moveColumns_default, "QSortFilterProxyModel_moveColumns_default")
	symaddr(&addr_QSortFilterProxyModel_multiData_default, "QSortFilterProxyModel_multiData_default")
	symaddr(&addr_QSortFilterProxyModel_resetInternalData_default, "QSortFilterProxyModel_resetInternalData_default")
	symaddr(&addr_QSortFilterProxyModel_event_default, "QSortFilterProxyModel_event_default")
	symaddr(&addr_QSortFilterProxyModel_eventFilter_default, "QSortFilterProxyModel_eventFilter_default")
	symaddr(&addr_QSortFilterProxyModel_timerEvent_default, "QSortFilterProxyModel_timerEvent_default")
	symaddr(&addr_QSortFilterProxyModel_childEvent_default, "QSortFilterProxyModel_childEvent_default")
	symaddr(&addr_QSortFilterProxyModel_customEvent_default, "QSortFilterProxyModel_customEvent_default")
	symaddr(&addr_QSortFilterProxyModel_connectNotify_default, "QSortFilterProxyModel_connectNotify_default")
	symaddr(&addr_QSortFilterProxyModel_disconnectNotify_default, "QSortFilterProxyModel_disconnectNotify_default")
	reg(&pg_QSortFilterProxyModel_connect_dynamicSortFilterChanged, "QSortFilterProxyModel_connect_dynamicSortFilterChanged")
	reg(&pg_QSortFilterProxyModel_connect_filterCaseSensitivityChanged, "QSortFilterProxyModel_connect_filterCaseSensitivityChanged")
	reg(&pg_QSortFilterProxyModel_connect_sortCaseSensitivityChanged, "QSortFilterProxyModel_connect_sortCaseSensitivityChanged")
	reg(&pg_QSortFilterProxyModel_connect_sortLocaleAwareChanged, "QSortFilterProxyModel_connect_sortLocaleAwareChanged")
	reg(&pg_QSortFilterProxyModel_connect_sortRoleChanged, "QSortFilterProxyModel_connect_sortRoleChanged")
	reg(&pg_QSortFilterProxyModel_connect_filterRoleChanged, "QSortFilterProxyModel_connect_filterRoleChanged")
	reg(&pg_QSortFilterProxyModel_connect_recursiveFilteringEnabledChanged, "QSortFilterProxyModel_connect_recursiveFilteringEnabledChanged")
	reg(&pg_QSortFilterProxyModel_connect_autoAcceptChildRowsChanged, "QSortFilterProxyModel_connect_autoAcceptChildRowsChanged")
	reg(&pg_QSortFilterProxyModel_connect_sourceModelChanged, "QSortFilterProxyModel_connect_sourceModelChanged")
	reg(&pg_QSortFilterProxyModel_connect_headerDataChanged, "QSortFilterProxyModel_connect_headerDataChanged")
	reg(&pg_QSortFilterProxyModel_connect_layoutChanged, "QSortFilterProxyModel_connect_layoutChanged")
	reg(&pg_QSortFilterProxyModel_connect_layoutChanged2, "QSortFilterProxyModel_connect_layoutChanged2")
	reg(&pg_QSortFilterProxyModel_connect_layoutAboutToBeChanged, "QSortFilterProxyModel_connect_layoutAboutToBeChanged")
	reg(&pg_QSortFilterProxyModel_connect_layoutAboutToBeChanged2, "QSortFilterProxyModel_connect_layoutAboutToBeChanged2")
	reg(&pg_QSortFilterProxyModel_connect_rowsAboutToBeInserted, "QSortFilterProxyModel_connect_rowsAboutToBeInserted")
	reg(&pg_QSortFilterProxyModel_connect_rowsInserted, "QSortFilterProxyModel_connect_rowsInserted")
	reg(&pg_QSortFilterProxyModel_connect_rowsAboutToBeRemoved, "QSortFilterProxyModel_connect_rowsAboutToBeRemoved")
	reg(&pg_QSortFilterProxyModel_connect_rowsRemoved, "QSortFilterProxyModel_connect_rowsRemoved")
	reg(&pg_QSortFilterProxyModel_connect_columnsAboutToBeInserted, "QSortFilterProxyModel_connect_columnsAboutToBeInserted")
	reg(&pg_QSortFilterProxyModel_connect_columnsInserted, "QSortFilterProxyModel_connect_columnsInserted")
	reg(&pg_QSortFilterProxyModel_connect_columnsAboutToBeRemoved, "QSortFilterProxyModel_connect_columnsAboutToBeRemoved")
	reg(&pg_QSortFilterProxyModel_connect_columnsRemoved, "QSortFilterProxyModel_connect_columnsRemoved")
	reg(&pg_QSortFilterProxyModel_connect_modelAboutToBeReset, "QSortFilterProxyModel_connect_modelAboutToBeReset")
	reg(&pg_QSortFilterProxyModel_connect_modelReset, "QSortFilterProxyModel_connect_modelReset")
	reg(&pg_QSortFilterProxyModel_connect_rowsAboutToBeMoved, "QSortFilterProxyModel_connect_rowsAboutToBeMoved")
	reg(&pg_QSortFilterProxyModel_connect_rowsMoved, "QSortFilterProxyModel_connect_rowsMoved")
	reg(&pg_QSortFilterProxyModel_connect_columnsAboutToBeMoved, "QSortFilterProxyModel_connect_columnsAboutToBeMoved")
	reg(&pg_QSortFilterProxyModel_connect_columnsMoved, "QSortFilterProxyModel_connect_columnsMoved")
	reg(&pg_QSortFilterProxyModel_connect_destroyed, "QSortFilterProxyModel_connect_destroyed")
	reg(&pg_QSortFilterProxyModel_connect_destroyed2, "QSortFilterProxyModel_connect_destroyed2")
	reg(&pg_QSortFilterProxyModel_connect_objectNameChanged, "QSortFilterProxyModel_connect_objectNameChanged")
	reg(&pg_QSortFilterProxyModel__disconnect, "QSortFilterProxyModel__disconnect")
	reg(&pg_QSortFilterProxyModel__set_override, "QSortFilterProxyModel__set_override")
	reg(&pg_QSortFilterProxyModel_delete, "QSortFilterProxyModel_delete")
}

type QSortFilterProxyModel struct {
	h uintptr
	*QAbstractProxyModel
}

func bindQSortFilterProxyModel(h uintptr) *QSortFilterProxyModel {
	if h == 0 {
		return nil
	}
	var __c [3]uintptr
	purego.SyscallN(addr_QSortFilterProxyModel_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QSortFilterProxyModel{h: h, QAbstractProxyModel: &QAbstractProxyModel{h: __c[0], QAbstractItemModel: &QAbstractItemModel{h: __c[1], QObject: &QObject{h: __c[2]}}}}
}

func newownedQSortFilterProxyModel(h uintptr) *QSortFilterProxyModel { return bindQSortFilterProxyModel(h) }

func (this *QSortFilterProxyModel) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QSortFilterProxyModel) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQSortFilterProxyModel(h unsafe.Pointer) *QSortFilterProxyModel { return bindQSortFilterProxyModel(uintptr(h)) }
func UnsafeNewQSortFilterProxyModelOwned(h unsafe.Pointer) *QSortFilterProxyModel { return newownedQSortFilterProxyModel(uintptr(h)) }

func (this *QSortFilterProxyModel) UnsafeRelease() {
	this.h = 0
	if this.QAbstractProxyModel != nil {
		this.QAbstractProxyModel.UnsafeRelease()
	}
}

func NewQSortFilterProxyModel() *QSortFilterProxyModel {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QSortFilterProxyModel_new, uintptr(__goid))
	__o := newownedQSortFilterProxyModel(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QSortFilterProxyModel_virtIdx, QSortFilterProxyModel_enable)
	return __o
}

func NewQSortFilterProxyModel2(parent *QObject) *QSortFilterProxyModel {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QSortFilterProxyModel_new2, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQSortFilterProxyModel(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QSortFilterProxyModel_virtIdx, QSortFilterProxyModel_enable)
	return __o
}

func QSortFilterProxyModel_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QSortFilterProxyModel_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QSortFilterProxyModel) FilterRegularExpression() *QRegularExpression {
	__r, _, _ := purego.SyscallN(addr_QSortFilterProxyModel_filterRegularExpression, this.h)
	runtime.KeepAlive(this)
	return newownedQRegularExpression(__r)
}

func (this *QSortFilterProxyModel) FilterKeyColumn() int {
	__r0, _, _ := purego.SyscallN(addr_QSortFilterProxyModel_filterKeyColumn, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QSortFilterProxyModel) SetFilterKeyColumn(column int) {
	purego.SyscallN(addr_QSortFilterProxyModel_setFilterKeyColumn, this.h, uintptr(column))
	runtime.KeepAlive(this)
}

func (this *QSortFilterProxyModel) FilterCaseSensitivity() CaseSensitivity {
	__r0, _, _ := purego.SyscallN(addr_QSortFilterProxyModel_filterCaseSensitivity, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return CaseSensitivity(__r)
}

func (this *QSortFilterProxyModel) SetFilterCaseSensitivity(cs CaseSensitivity) {
	purego.SyscallN(addr_QSortFilterProxyModel_setFilterCaseSensitivity, this.h, uintptr(cs))
	runtime.KeepAlive(this)
}

func (this *QSortFilterProxyModel) SortCaseSensitivity() CaseSensitivity {
	__r0, _, _ := purego.SyscallN(addr_QSortFilterProxyModel_sortCaseSensitivity, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return CaseSensitivity(__r)
}

func (this *QSortFilterProxyModel) SetSortCaseSensitivity(cs CaseSensitivity) {
	purego.SyscallN(addr_QSortFilterProxyModel_setSortCaseSensitivity, this.h, uintptr(cs))
	runtime.KeepAlive(this)
}

func (this *QSortFilterProxyModel) IsSortLocaleAware() bool {
	__r0, _, _ := purego.SyscallN(addr_QSortFilterProxyModel_isSortLocaleAware, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QSortFilterProxyModel) SetSortLocaleAware(on bool) {
	purego.SyscallN(addr_QSortFilterProxyModel_setSortLocaleAware, this.h, uintptr(b2i(on)))
	runtime.KeepAlive(this)
}

func (this *QSortFilterProxyModel) SortColumn() int {
	__r0, _, _ := purego.SyscallN(addr_QSortFilterProxyModel_sortColumn, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QSortFilterProxyModel) SortOrder() SortOrder {
	__r0, _, _ := purego.SyscallN(addr_QSortFilterProxyModel_sortOrder, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return SortOrder(__r)
}

func (this *QSortFilterProxyModel) DynamicSortFilter() bool {
	__r0, _, _ := purego.SyscallN(addr_QSortFilterProxyModel_dynamicSortFilter, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QSortFilterProxyModel) SetDynamicSortFilter(enable bool) {
	purego.SyscallN(addr_QSortFilterProxyModel_setDynamicSortFilter, this.h, uintptr(b2i(enable)))
	runtime.KeepAlive(this)
}

func (this *QSortFilterProxyModel) SortRole() int {
	__r0, _, _ := purego.SyscallN(addr_QSortFilterProxyModel_sortRole, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QSortFilterProxyModel) SetSortRole(role int) {
	purego.SyscallN(addr_QSortFilterProxyModel_setSortRole, this.h, uintptr(role))
	runtime.KeepAlive(this)
}

func (this *QSortFilterProxyModel) FilterRole() int {
	__r0, _, _ := purego.SyscallN(addr_QSortFilterProxyModel_filterRole, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QSortFilterProxyModel) SetFilterRole(role int) {
	purego.SyscallN(addr_QSortFilterProxyModel_setFilterRole, this.h, uintptr(role))
	runtime.KeepAlive(this)
}

func (this *QSortFilterProxyModel) IsRecursiveFilteringEnabled() bool {
	__r0, _, _ := purego.SyscallN(addr_QSortFilterProxyModel_isRecursiveFilteringEnabled, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QSortFilterProxyModel) SetRecursiveFilteringEnabled(recursive bool) {
	purego.SyscallN(addr_QSortFilterProxyModel_setRecursiveFilteringEnabled, this.h, uintptr(b2i(recursive)))
	runtime.KeepAlive(this)
}

func (this *QSortFilterProxyModel) AutoAcceptChildRows() bool {
	__r0, _, _ := purego.SyscallN(addr_QSortFilterProxyModel_autoAcceptChildRows, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QSortFilterProxyModel) SetAutoAcceptChildRows(accept bool) {
	purego.SyscallN(addr_QSortFilterProxyModel_setAutoAcceptChildRows, this.h, uintptr(b2i(accept)))
	runtime.KeepAlive(this)
}

func (this *QSortFilterProxyModel) SetFilterRegularExpression(pattern string) {
	var pattern_p uintptr
	if len(pattern) > 0 {
		pattern_p = uintptr(unsafe.Pointer(unsafe.StringData(pattern)))
	}
	purego.SyscallN(addr_QSortFilterProxyModel_setFilterRegularExpression, this.h, pattern_p, uintptr(len(pattern)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pattern)
}

func (this *QSortFilterProxyModel) SetFilterRegularExpressionWithRegularExpression(regularExpression *QRegularExpression) {
	purego.SyscallN(addr_QSortFilterProxyModel_setFilterRegularExpression2, this.h, uintptr(regularExpression.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(regularExpression)
}

func (this *QSortFilterProxyModel) SetFilterWildcard(pattern string) {
	var pattern_p uintptr
	if len(pattern) > 0 {
		pattern_p = uintptr(unsafe.Pointer(unsafe.StringData(pattern)))
	}
	purego.SyscallN(addr_QSortFilterProxyModel_setFilterWildcard, this.h, pattern_p, uintptr(len(pattern)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pattern)
}

func (this *QSortFilterProxyModel) SetFilterFixedString(pattern string) {
	var pattern_p uintptr
	if len(pattern) > 0 {
		pattern_p = uintptr(unsafe.Pointer(unsafe.StringData(pattern)))
	}
	purego.SyscallN(addr_QSortFilterProxyModel_setFilterFixedString, this.h, pattern_p, uintptr(len(pattern)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pattern)
}

func (this *QSortFilterProxyModel) Invalidate() {
	purego.SyscallN(addr_QSortFilterProxyModel_invalidate, this.h)
	runtime.KeepAlive(this)
}

func (this *QSortFilterProxyModel) BeginFilterChange() {
	purego.SyscallN(addr_QSortFilterProxyModel_beginFilterChange, this.h)
	runtime.KeepAlive(this)
}

func (this *QSortFilterProxyModel) EndFilterChange() {
	purego.SyscallN(addr_QSortFilterProxyModel_endFilterChange, this.h)
	runtime.KeepAlive(this)
}

func (this *QSortFilterProxyModel) InvalidateFilter() {
	purego.SyscallN(addr_QSortFilterProxyModel_invalidateFilter, this.h)
	runtime.KeepAlive(this)
}

func (this *QSortFilterProxyModel) InvalidateRowsFilter() {
	purego.SyscallN(addr_QSortFilterProxyModel_invalidateRowsFilter, this.h)
	runtime.KeepAlive(this)
}

func (this *QSortFilterProxyModel) InvalidateColumnsFilter() {
	purego.SyscallN(addr_QSortFilterProxyModel_invalidateColumnsFilter, this.h)
	runtime.KeepAlive(this)
}

func (this *QSortFilterProxyModel) DynamicSortFilterChanged(dynamicSortFilter bool) {
	purego.SyscallN(addr_QSortFilterProxyModel_dynamicSortFilterChanged, this.h, uintptr(b2i(dynamicSortFilter)))
	runtime.KeepAlive(this)
}

func (this *QSortFilterProxyModel) FilterCaseSensitivityChanged(filterCaseSensitivity CaseSensitivity) {
	purego.SyscallN(addr_QSortFilterProxyModel_filterCaseSensitivityChanged, this.h, uintptr(filterCaseSensitivity))
	runtime.KeepAlive(this)
}

func (this *QSortFilterProxyModel) SortCaseSensitivityChanged(sortCaseSensitivity CaseSensitivity) {
	purego.SyscallN(addr_QSortFilterProxyModel_sortCaseSensitivityChanged, this.h, uintptr(sortCaseSensitivity))
	runtime.KeepAlive(this)
}

func (this *QSortFilterProxyModel) SortLocaleAwareChanged(sortLocaleAware bool) {
	purego.SyscallN(addr_QSortFilterProxyModel_sortLocaleAwareChanged, this.h, uintptr(b2i(sortLocaleAware)))
	runtime.KeepAlive(this)
}

func (this *QSortFilterProxyModel) SortRoleChanged(sortRole int) {
	purego.SyscallN(addr_QSortFilterProxyModel_sortRoleChanged, this.h, uintptr(sortRole))
	runtime.KeepAlive(this)
}

func (this *QSortFilterProxyModel) FilterRoleChanged(filterRole int) {
	purego.SyscallN(addr_QSortFilterProxyModel_filterRoleChanged, this.h, uintptr(filterRole))
	runtime.KeepAlive(this)
}

func (this *QSortFilterProxyModel) RecursiveFilteringEnabledChanged(recursiveFilteringEnabled bool) {
	purego.SyscallN(addr_QSortFilterProxyModel_recursiveFilteringEnabledChanged, this.h, uintptr(b2i(recursiveFilteringEnabled)))
	runtime.KeepAlive(this)
}

func (this *QSortFilterProxyModel) AutoAcceptChildRowsChanged(autoAcceptChildRows bool) {
	purego.SyscallN(addr_QSortFilterProxyModel_autoAcceptChildRowsChanged, this.h, uintptr(b2i(autoAcceptChildRows)))
	runtime.KeepAlive(this)
}

func (this *QSortFilterProxyModel) CreateIndex(row int, column int) *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QSortFilterProxyModel_createIndex, this.h, uintptr(row), uintptr(column))
	runtime.KeepAlive(this)
	return newownedQModelIndex(__r)
}

func (this *QSortFilterProxyModel) CreateIndex2(row int, column int, id uintptr) *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QSortFilterProxyModel_createIndex2, this.h, uintptr(row), uintptr(column), uintptr(id))
	runtime.KeepAlive(this)
	return newownedQModelIndex(__r)
}

func (this *QSortFilterProxyModel) EncodeData(indexes []*QModelIndex, stream *QDataStream) {
	indexes_arr := make([]uintptr, len(indexes))
	for __i, __e := range indexes {
		indexes_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var indexes_p uintptr
	if len(indexes_arr) > 0 {
		indexes_p = uintptr(unsafe.Pointer(&indexes_arr[0]))
	}
	purego.SyscallN(addr_QSortFilterProxyModel_encodeData, this.h, indexes_p, uintptr(len(indexes)), uintptr(stream.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(indexes)
	runtime.KeepAlive(indexes_arr)
	runtime.KeepAlive(stream)
}

func (this *QSortFilterProxyModel) DecodeData(row int, column int, parent *QModelIndex, stream *QDataStream) bool {
	__r0, _, _ := purego.SyscallN(addr_QSortFilterProxyModel_decodeData, this.h, uintptr(row), uintptr(column), uintptr(parent.UnsafePointer()), uintptr(stream.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	runtime.KeepAlive(stream)
	return __r != 0
}

func (this *QSortFilterProxyModel) BeginInsertRows(parent *QModelIndex, first int, last int) {
	purego.SyscallN(addr_QSortFilterProxyModel_beginInsertRows, this.h, uintptr(parent.UnsafePointer()), uintptr(first), uintptr(last))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
}

func (this *QSortFilterProxyModel) EndInsertRows() {
	purego.SyscallN(addr_QSortFilterProxyModel_endInsertRows, this.h)
	runtime.KeepAlive(this)
}

func (this *QSortFilterProxyModel) BeginRemoveRows(parent *QModelIndex, first int, last int) {
	purego.SyscallN(addr_QSortFilterProxyModel_beginRemoveRows, this.h, uintptr(parent.UnsafePointer()), uintptr(first), uintptr(last))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
}

func (this *QSortFilterProxyModel) EndRemoveRows() {
	purego.SyscallN(addr_QSortFilterProxyModel_endRemoveRows, this.h)
	runtime.KeepAlive(this)
}

func (this *QSortFilterProxyModel) BeginMoveRows(sourceParent *QModelIndex, sourceFirst int, sourceLast int, destinationParent *QModelIndex, destinationRow int) bool {
	__r0, _, _ := purego.SyscallN(addr_QSortFilterProxyModel_beginMoveRows, this.h, uintptr(sourceParent.UnsafePointer()), uintptr(sourceFirst), uintptr(sourceLast), uintptr(destinationParent.UnsafePointer()), uintptr(destinationRow))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sourceParent)
	runtime.KeepAlive(destinationParent)
	return __r != 0
}

func (this *QSortFilterProxyModel) EndMoveRows() {
	purego.SyscallN(addr_QSortFilterProxyModel_endMoveRows, this.h)
	runtime.KeepAlive(this)
}

func (this *QSortFilterProxyModel) BeginInsertColumns(parent *QModelIndex, first int, last int) {
	purego.SyscallN(addr_QSortFilterProxyModel_beginInsertColumns, this.h, uintptr(parent.UnsafePointer()), uintptr(first), uintptr(last))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
}

func (this *QSortFilterProxyModel) EndInsertColumns() {
	purego.SyscallN(addr_QSortFilterProxyModel_endInsertColumns, this.h)
	runtime.KeepAlive(this)
}

func (this *QSortFilterProxyModel) BeginRemoveColumns(parent *QModelIndex, first int, last int) {
	purego.SyscallN(addr_QSortFilterProxyModel_beginRemoveColumns, this.h, uintptr(parent.UnsafePointer()), uintptr(first), uintptr(last))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
}

func (this *QSortFilterProxyModel) EndRemoveColumns() {
	purego.SyscallN(addr_QSortFilterProxyModel_endRemoveColumns, this.h)
	runtime.KeepAlive(this)
}

func (this *QSortFilterProxyModel) BeginMoveColumns(sourceParent *QModelIndex, sourceFirst int, sourceLast int, destinationParent *QModelIndex, destinationColumn int) bool {
	__r0, _, _ := purego.SyscallN(addr_QSortFilterProxyModel_beginMoveColumns, this.h, uintptr(sourceParent.UnsafePointer()), uintptr(sourceFirst), uintptr(sourceLast), uintptr(destinationParent.UnsafePointer()), uintptr(destinationColumn))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sourceParent)
	runtime.KeepAlive(destinationParent)
	return __r != 0
}

func (this *QSortFilterProxyModel) EndMoveColumns() {
	purego.SyscallN(addr_QSortFilterProxyModel_endMoveColumns, this.h)
	runtime.KeepAlive(this)
}

func (this *QSortFilterProxyModel) BeginResetModel() {
	purego.SyscallN(addr_QSortFilterProxyModel_beginResetModel, this.h)
	runtime.KeepAlive(this)
}

func (this *QSortFilterProxyModel) EndResetModel() {
	purego.SyscallN(addr_QSortFilterProxyModel_endResetModel, this.h)
	runtime.KeepAlive(this)
}

func (this *QSortFilterProxyModel) ChangePersistentIndex(from *QModelIndex, to *QModelIndex) {
	purego.SyscallN(addr_QSortFilterProxyModel_changePersistentIndex, this.h, uintptr(from.UnsafePointer()), uintptr(to.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(from)
	runtime.KeepAlive(to)
}

func (this *QSortFilterProxyModel) ChangePersistentIndexList(from []*QModelIndex, to []*QModelIndex) {
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
	purego.SyscallN(addr_QSortFilterProxyModel_changePersistentIndexList, this.h, from_p, uintptr(len(from)), to_p, uintptr(len(to)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(from)
	runtime.KeepAlive(from_arr)
	runtime.KeepAlive(to)
	runtime.KeepAlive(to_arr)
}

func (this *QSortFilterProxyModel) PersistentIndexList() []*QModelIndex {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QSortFilterProxyModel_persistentIndexList, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QModelIndex, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQModelIndex(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QSortFilterProxyModel_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QSortFilterProxyModel_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QSortFilterProxyModel_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QSortFilterProxyModel_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QSortFilterProxyModel) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QSortFilterProxyModel_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QSortFilterProxyModel) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QSortFilterProxyModel_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QSortFilterProxyModel) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QSortFilterProxyModel_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QSortFilterProxyModel) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QSortFilterProxyModel_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QSortFilterProxyModel_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QSortFilterProxyModel_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QSortFilterProxyModel_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QSortFilterProxyModel_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func (this *QSortFilterProxyModel) EndFilterChangeWithDirections(directions QSortFilterProxyModel__Direction) {
	purego.SyscallN(addr_QSortFilterProxyModel_endFilterChange2, this.h, uintptr(directions))
	runtime.KeepAlive(this)
}

func QSortFilterProxyModel_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QSortFilterProxyModel_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QSortFilterProxyModel_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QSortFilterProxyModel_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QSortFilterProxyModel) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QSortFilterProxyModel_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QSortFilterProxyModel) SetSourceModelDefault(sourceModel *QAbstractItemModel) {
	purego.SyscallN(addr_QSortFilterProxyModel_setSourceModel_default, this.h, uintptr(sourceModel.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sourceModel)
}

func (this *QSortFilterProxyModel) MapToSourceDefault(proxyIndex *QModelIndex) *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QSortFilterProxyModel_mapToSource_default, this.h, uintptr(proxyIndex.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(proxyIndex)
	return newownedQModelIndex(__r)
}

func (this *QSortFilterProxyModel) MapFromSourceDefault(sourceIndex *QModelIndex) *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QSortFilterProxyModel_mapFromSource_default, this.h, uintptr(sourceIndex.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sourceIndex)
	return newownedQModelIndex(__r)
}

func (this *QSortFilterProxyModel) MapSelectionToSourceDefault(proxySelection *QItemSelection) *QItemSelection {
	__r, _, _ := purego.SyscallN(addr_QSortFilterProxyModel_mapSelectionToSource_default, this.h, uintptr(proxySelection.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(proxySelection)
	return newownedQItemSelection(__r)
}

func (this *QSortFilterProxyModel) MapSelectionFromSourceDefault(sourceSelection *QItemSelection) *QItemSelection {
	__r, _, _ := purego.SyscallN(addr_QSortFilterProxyModel_mapSelectionFromSource_default, this.h, uintptr(sourceSelection.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sourceSelection)
	return newownedQItemSelection(__r)
}

func (this *QSortFilterProxyModel) FilterAcceptsRowDefault(source_row int, source_parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QSortFilterProxyModel_filterAcceptsRow_default, this.h, uintptr(source_row), uintptr(source_parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(source_parent)
	return __r != 0
}

func (this *QSortFilterProxyModel) FilterAcceptsColumnDefault(source_column int, source_parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QSortFilterProxyModel_filterAcceptsColumn_default, this.h, uintptr(source_column), uintptr(source_parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(source_parent)
	return __r != 0
}

func (this *QSortFilterProxyModel) LessThanDefault(source_left *QModelIndex, source_right *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QSortFilterProxyModel_lessThan_default, this.h, uintptr(source_left.UnsafePointer()), uintptr(source_right.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(source_left)
	runtime.KeepAlive(source_right)
	return __r != 0
}

func (this *QSortFilterProxyModel) IndexDefault(row int, column int, parent *QModelIndex) *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QSortFilterProxyModel_index_default, this.h, uintptr(row), uintptr(column), uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return newownedQModelIndex(__r)
}

func (this *QSortFilterProxyModel) ParentDefault(child *QModelIndex) *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QSortFilterProxyModel_parent_default, this.h, uintptr(child.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(child)
	return newownedQModelIndex(__r)
}

func (this *QSortFilterProxyModel) SiblingDefault(row int, column int, idx *QModelIndex) *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QSortFilterProxyModel_sibling_default, this.h, uintptr(row), uintptr(column), uintptr(idx.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(idx)
	return newownedQModelIndex(__r)
}

func (this *QSortFilterProxyModel) RowCountDefault(parent *QModelIndex) int {
	__r0, _, _ := purego.SyscallN(addr_QSortFilterProxyModel_rowCount_default, this.h, uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return int(__r)
}

func (this *QSortFilterProxyModel) ColumnCountDefault(parent *QModelIndex) int {
	__r0, _, _ := purego.SyscallN(addr_QSortFilterProxyModel_columnCount_default, this.h, uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return int(__r)
}

func (this *QSortFilterProxyModel) HasChildrenDefault(parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QSortFilterProxyModel_hasChildren_default, this.h, uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QSortFilterProxyModel) DataDefault(index *QModelIndex, role int) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QSortFilterProxyModel_data_default, this.h, uintptr(index.UnsafePointer()), uintptr(role))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return newownedQVariant(__r)
}

func (this *QSortFilterProxyModel) SetDataDefault(index *QModelIndex, value *QVariant, role int) bool {
	__r0, _, _ := purego.SyscallN(addr_QSortFilterProxyModel_setData_default, this.h, uintptr(index.UnsafePointer()), uintptr(value.UnsafePointer()), uintptr(role))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	runtime.KeepAlive(value)
	return __r != 0
}

func (this *QSortFilterProxyModel) HeaderDataDefault(section int, orientation Orientation, role int) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QSortFilterProxyModel_headerData_default, this.h, uintptr(section), uintptr(orientation), uintptr(role))
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QSortFilterProxyModel) SetHeaderDataDefault(section int, orientation Orientation, value *QVariant, role int) bool {
	__r0, _, _ := purego.SyscallN(addr_QSortFilterProxyModel_setHeaderData_default, this.h, uintptr(section), uintptr(orientation), uintptr(value.UnsafePointer()), uintptr(role))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(value)
	return __r != 0
}

func (this *QSortFilterProxyModel) MimeDataDefault(indexes []*QModelIndex) *QMimeData {
	indexes_arr := make([]uintptr, len(indexes))
	for __i, __e := range indexes {
		indexes_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var indexes_p uintptr
	if len(indexes_arr) > 0 {
		indexes_p = uintptr(unsafe.Pointer(&indexes_arr[0]))
	}
	__r, _, _ := purego.SyscallN(addr_QSortFilterProxyModel_mimeData_default, this.h, indexes_p, uintptr(len(indexes)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(indexes)
	runtime.KeepAlive(indexes_arr)
	return bindQMimeData(__r)
}

func (this *QSortFilterProxyModel) DropMimeDataDefault(data *QMimeData, action DropAction, row int, column int, parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QSortFilterProxyModel_dropMimeData_default, this.h, uintptr(data.UnsafePointer()), uintptr(action), uintptr(row), uintptr(column), uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(data)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QSortFilterProxyModel) InsertRowsDefault(row int, count int, parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QSortFilterProxyModel_insertRows_default, this.h, uintptr(row), uintptr(count), uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QSortFilterProxyModel) InsertColumnsDefault(column int, count int, parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QSortFilterProxyModel_insertColumns_default, this.h, uintptr(column), uintptr(count), uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QSortFilterProxyModel) RemoveRowsDefault(row int, count int, parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QSortFilterProxyModel_removeRows_default, this.h, uintptr(row), uintptr(count), uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QSortFilterProxyModel) RemoveColumnsDefault(column int, count int, parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QSortFilterProxyModel_removeColumns_default, this.h, uintptr(column), uintptr(count), uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QSortFilterProxyModel) FetchMoreDefault(parent *QModelIndex) {
	purego.SyscallN(addr_QSortFilterProxyModel_fetchMore_default, this.h, uintptr(parent.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
}

func (this *QSortFilterProxyModel) CanFetchMoreDefault(parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QSortFilterProxyModel_canFetchMore_default, this.h, uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QSortFilterProxyModel) FlagsDefault(index *QModelIndex) ItemFlag {
	__r0, _, _ := purego.SyscallN(addr_QSortFilterProxyModel_flags_default, this.h, uintptr(index.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return ItemFlag(__r)
}

func (this *QSortFilterProxyModel) BuddyDefault(index *QModelIndex) *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QSortFilterProxyModel_buddy_default, this.h, uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return newownedQModelIndex(__r)
}

func (this *QSortFilterProxyModel) MatchDefault(start *QModelIndex, role int, value *QVariant, hits int, flags MatchFlag) []*QModelIndex {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QSortFilterProxyModel_match_default, this.h, uintptr(start.UnsafePointer()), uintptr(role), uintptr(value.UnsafePointer()), uintptr(hits), uintptr(flags), uintptr(unsafe.Pointer(&__n)))
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

func (this *QSortFilterProxyModel) SpanDefault(index *QModelIndex) *QSize {
	__r, _, _ := purego.SyscallN(addr_QSortFilterProxyModel_span_default, this.h, uintptr(index.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return newownedQSize(__r)
}

func (this *QSortFilterProxyModel) SortDefault(column int, order SortOrder) {
	purego.SyscallN(addr_QSortFilterProxyModel_sort_default, this.h, uintptr(column), uintptr(order))
	runtime.KeepAlive(this)
}

func (this *QSortFilterProxyModel) MimeTypesDefault() []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QSortFilterProxyModel_mimeTypes_default, this.h, uintptr(unsafe.Pointer(&__n)))
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

func (this *QSortFilterProxyModel) SupportedDropActionsDefault() DropAction {
	__r0, _, _ := purego.SyscallN(addr_QSortFilterProxyModel_supportedDropActions_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return DropAction(__r)
}

func (this *QSortFilterProxyModel) SubmitDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QSortFilterProxyModel_submit_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QSortFilterProxyModel) RevertDefault() {
	purego.SyscallN(addr_QSortFilterProxyModel_revert_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QSortFilterProxyModel) ClearItemDataDefault(index *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QSortFilterProxyModel_clearItemData_default, this.h, uintptr(index.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	return __r != 0
}

func (this *QSortFilterProxyModel) CanDropMimeDataDefault(data *QMimeData, action DropAction, row int, column int, parent *QModelIndex) bool {
	__r0, _, _ := purego.SyscallN(addr_QSortFilterProxyModel_canDropMimeData_default, this.h, uintptr(data.UnsafePointer()), uintptr(action), uintptr(row), uintptr(column), uintptr(parent.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(data)
	runtime.KeepAlive(parent)
	return __r != 0
}

func (this *QSortFilterProxyModel) SupportedDragActionsDefault() DropAction {
	__r0, _, _ := purego.SyscallN(addr_QSortFilterProxyModel_supportedDragActions_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return DropAction(__r)
}

func (this *QSortFilterProxyModel) MoveRowsDefault(sourceParent *QModelIndex, sourceRow int, count int, destinationParent *QModelIndex, destinationChild int) bool {
	__r0, _, _ := purego.SyscallN(addr_QSortFilterProxyModel_moveRows_default, this.h, uintptr(sourceParent.UnsafePointer()), uintptr(sourceRow), uintptr(count), uintptr(destinationParent.UnsafePointer()), uintptr(destinationChild))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sourceParent)
	runtime.KeepAlive(destinationParent)
	return __r != 0
}

func (this *QSortFilterProxyModel) MoveColumnsDefault(sourceParent *QModelIndex, sourceColumn int, count int, destinationParent *QModelIndex, destinationChild int) bool {
	__r0, _, _ := purego.SyscallN(addr_QSortFilterProxyModel_moveColumns_default, this.h, uintptr(sourceParent.UnsafePointer()), uintptr(sourceColumn), uintptr(count), uintptr(destinationParent.UnsafePointer()), uintptr(destinationChild))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(sourceParent)
	runtime.KeepAlive(destinationParent)
	return __r != 0
}

func (this *QSortFilterProxyModel) MultiDataDefault(index *QModelIndex, roleDataSpan *QModelRoleDataSpan) {
	purego.SyscallN(addr_QSortFilterProxyModel_multiData_default, this.h, uintptr(index.UnsafePointer()), uintptr(roleDataSpan.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(index)
	runtime.KeepAlive(roleDataSpan)
}

func (this *QSortFilterProxyModel) ResetInternalDataDefault() {
	purego.SyscallN(addr_QSortFilterProxyModel_resetInternalData_default, this.h)
	runtime.KeepAlive(this)
}

func (this *QSortFilterProxyModel) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QSortFilterProxyModel_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QSortFilterProxyModel) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QSortFilterProxyModel_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QSortFilterProxyModel) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QSortFilterProxyModel_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSortFilterProxyModel) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QSortFilterProxyModel_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSortFilterProxyModel) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QSortFilterProxyModel_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSortFilterProxyModel) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QSortFilterProxyModel_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QSortFilterProxyModel) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QSortFilterProxyModel_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QSortFilterProxyModel) OnDynamicSortFilterChanged(fn func(bool)) int64 {
	id := registerCallback(func(dynamicSortFilter int32) { fn(dynamicSortFilter != 0) })
	pg_QSortFilterProxyModel_connect_dynamicSortFilterChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSortFilterProxyModel) OnFilterCaseSensitivityChanged(fn func(CaseSensitivity)) int64 {
	id := registerCallback(func(filterCaseSensitivity int32) { fn(CaseSensitivity(filterCaseSensitivity)) })
	pg_QSortFilterProxyModel_connect_filterCaseSensitivityChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSortFilterProxyModel) OnSortCaseSensitivityChanged(fn func(CaseSensitivity)) int64 {
	id := registerCallback(func(sortCaseSensitivity int32) { fn(CaseSensitivity(sortCaseSensitivity)) })
	pg_QSortFilterProxyModel_connect_sortCaseSensitivityChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSortFilterProxyModel) OnSortLocaleAwareChanged(fn func(bool)) int64 {
	id := registerCallback(func(sortLocaleAware int32) { fn(sortLocaleAware != 0) })
	pg_QSortFilterProxyModel_connect_sortLocaleAwareChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSortFilterProxyModel) OnSortRoleChanged(fn func(int)) int64 {
	id := registerCallback(func(sortRole int32) { fn(int(sortRole)) })
	pg_QSortFilterProxyModel_connect_sortRoleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSortFilterProxyModel) OnFilterRoleChanged(fn func(int)) int64 {
	id := registerCallback(func(filterRole int32) { fn(int(filterRole)) })
	pg_QSortFilterProxyModel_connect_filterRoleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSortFilterProxyModel) OnRecursiveFilteringEnabledChanged(fn func(bool)) int64 {
	id := registerCallback(func(recursiveFilteringEnabled int32) { fn(recursiveFilteringEnabled != 0) })
	pg_QSortFilterProxyModel_connect_recursiveFilteringEnabledChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSortFilterProxyModel) OnAutoAcceptChildRowsChanged(fn func(bool)) int64 {
	id := registerCallback(func(autoAcceptChildRows int32) { fn(autoAcceptChildRows != 0) })
	pg_QSortFilterProxyModel_connect_autoAcceptChildRowsChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSortFilterProxyModel) OnSourceModelChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QSortFilterProxyModel_connect_sourceModelChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSortFilterProxyModel) OnHeaderDataChanged(fn func(Orientation, int, int)) int64 {
	id := registerCallback(func(orientation int32, first int32, last int32) { fn(Orientation(orientation), int(first), int(last)) })
	pg_QSortFilterProxyModel_connect_headerDataChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSortFilterProxyModel) OnLayoutChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QSortFilterProxyModel_connect_layoutChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSortFilterProxyModel) OnLayoutChanged2(fn func([]*QPersistentModelIndex, QAbstractItemModel__LayoutChangeHint)) int64 {
	id := registerCallback(func(parents uintptr, parents_len int64, hint int32) { fn(func() []*QPersistentModelIndex { if parents == 0 || parents_len == 0 { return nil }; __sl := unsafe.Slice((*unsafe.Pointer)(unsafe.Pointer(parents)), int(parents_len)); __o := make([]*QPersistentModelIndex, int(parents_len)); for __i := range __o { __o[__i] = UnsafeNewQPersistentModelIndex(__sl[__i]) }; return __o }(), QAbstractItemModel__LayoutChangeHint(hint)) })
	pg_QSortFilterProxyModel_connect_layoutChanged2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSortFilterProxyModel) OnLayoutAboutToBeChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QSortFilterProxyModel_connect_layoutAboutToBeChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSortFilterProxyModel) OnLayoutAboutToBeChanged2(fn func([]*QPersistentModelIndex, QAbstractItemModel__LayoutChangeHint)) int64 {
	id := registerCallback(func(parents uintptr, parents_len int64, hint int32) { fn(func() []*QPersistentModelIndex { if parents == 0 || parents_len == 0 { return nil }; __sl := unsafe.Slice((*unsafe.Pointer)(unsafe.Pointer(parents)), int(parents_len)); __o := make([]*QPersistentModelIndex, int(parents_len)); for __i := range __o { __o[__i] = UnsafeNewQPersistentModelIndex(__sl[__i]) }; return __o }(), QAbstractItemModel__LayoutChangeHint(hint)) })
	pg_QSortFilterProxyModel_connect_layoutAboutToBeChanged2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSortFilterProxyModel) OnRowsAboutToBeInserted(fn func(*QModelIndex, int, int)) int64 {
	id := registerCallback(func(parent uintptr, first int32, last int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(first), int(last)) })
	pg_QSortFilterProxyModel_connect_rowsAboutToBeInserted(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSortFilterProxyModel) OnRowsInserted(fn func(*QModelIndex, int, int)) int64 {
	id := registerCallback(func(parent uintptr, first int32, last int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(first), int(last)) })
	pg_QSortFilterProxyModel_connect_rowsInserted(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSortFilterProxyModel) OnRowsAboutToBeRemoved(fn func(*QModelIndex, int, int)) int64 {
	id := registerCallback(func(parent uintptr, first int32, last int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(first), int(last)) })
	pg_QSortFilterProxyModel_connect_rowsAboutToBeRemoved(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSortFilterProxyModel) OnRowsRemoved(fn func(*QModelIndex, int, int)) int64 {
	id := registerCallback(func(parent uintptr, first int32, last int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(first), int(last)) })
	pg_QSortFilterProxyModel_connect_rowsRemoved(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSortFilterProxyModel) OnColumnsAboutToBeInserted(fn func(*QModelIndex, int, int)) int64 {
	id := registerCallback(func(parent uintptr, first int32, last int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(first), int(last)) })
	pg_QSortFilterProxyModel_connect_columnsAboutToBeInserted(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSortFilterProxyModel) OnColumnsInserted(fn func(*QModelIndex, int, int)) int64 {
	id := registerCallback(func(parent uintptr, first int32, last int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(first), int(last)) })
	pg_QSortFilterProxyModel_connect_columnsInserted(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSortFilterProxyModel) OnColumnsAboutToBeRemoved(fn func(*QModelIndex, int, int)) int64 {
	id := registerCallback(func(parent uintptr, first int32, last int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(first), int(last)) })
	pg_QSortFilterProxyModel_connect_columnsAboutToBeRemoved(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSortFilterProxyModel) OnColumnsRemoved(fn func(*QModelIndex, int, int)) int64 {
	id := registerCallback(func(parent uintptr, first int32, last int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)), int(first), int(last)) })
	pg_QSortFilterProxyModel_connect_columnsRemoved(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSortFilterProxyModel) OnModelAboutToBeReset(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QSortFilterProxyModel_connect_modelAboutToBeReset(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSortFilterProxyModel) OnModelReset(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QSortFilterProxyModel_connect_modelReset(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSortFilterProxyModel) OnRowsAboutToBeMoved(fn func(*QModelIndex, int, int, *QModelIndex, int)) int64 {
	id := registerCallback(func(sourceParent uintptr, sourceStart int32, sourceEnd int32, destinationParent uintptr, destinationRow int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(sourceParent)), int(sourceStart), int(sourceEnd), UnsafeNewQModelIndex(unsafe.Pointer(destinationParent)), int(destinationRow)) })
	pg_QSortFilterProxyModel_connect_rowsAboutToBeMoved(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSortFilterProxyModel) OnRowsMoved(fn func(*QModelIndex, int, int, *QModelIndex, int)) int64 {
	id := registerCallback(func(sourceParent uintptr, sourceStart int32, sourceEnd int32, destinationParent uintptr, destinationRow int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(sourceParent)), int(sourceStart), int(sourceEnd), UnsafeNewQModelIndex(unsafe.Pointer(destinationParent)), int(destinationRow)) })
	pg_QSortFilterProxyModel_connect_rowsMoved(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSortFilterProxyModel) OnColumnsAboutToBeMoved(fn func(*QModelIndex, int, int, *QModelIndex, int)) int64 {
	id := registerCallback(func(sourceParent uintptr, sourceStart int32, sourceEnd int32, destinationParent uintptr, destinationColumn int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(sourceParent)), int(sourceStart), int(sourceEnd), UnsafeNewQModelIndex(unsafe.Pointer(destinationParent)), int(destinationColumn)) })
	pg_QSortFilterProxyModel_connect_columnsAboutToBeMoved(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSortFilterProxyModel) OnColumnsMoved(fn func(*QModelIndex, int, int, *QModelIndex, int)) int64 {
	id := registerCallback(func(sourceParent uintptr, sourceStart int32, sourceEnd int32, destinationParent uintptr, destinationColumn int32) { fn(UnsafeNewQModelIndex(unsafe.Pointer(sourceParent)), int(sourceStart), int(sourceEnd), UnsafeNewQModelIndex(unsafe.Pointer(destinationParent)), int(destinationColumn)) })
	pg_QSortFilterProxyModel_connect_columnsMoved(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSortFilterProxyModel) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QSortFilterProxyModel_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSortFilterProxyModel) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QSortFilterProxyModel_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSortFilterProxyModel) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QSortFilterProxyModel_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSortFilterProxyModel) Disconnect(handle int64) {
	pg_QSortFilterProxyModel__disconnect(handle)
	unregisterCallback(handle)
}

func QSortFilterProxyModel_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case 7818871925547456643:
		return 1
	case -2333321292016234891:
		return 2
	case -1717009778236344012:
		return 3
	case 2788628866280405837:
		return 4
	case 9067522407071207004:
		return 5
	case -8878517896795061028:
		return 6
	case 436376032338170754:
		return 7
	case -5863021011917066845:
		return 8
	case -8948777187306027381:
		return 9
	case 7371339085820260189:
		return 10
	case 397329346813782521:
		return 11
	case 1960960171956884010:
		return 12
	case -8256963024669584834:
		return 13
	case 3369452259144514380:
		return 14
	case -8837375942098466299:
		return 15
	case 35191454575202219:
		return 16
	case -3297140021896693688:
		return 17
	case 4852200306590054850:
		return 18
	case 5175297091024974749:
		return 19
	case -9049994684305682212:
		return 20
	case -2775729825962395603:
		return 21
	case -8973293846230792237:
		return 22
	case 5104072549182222656:
		return 23
	case 2078548737056540420:
		return 24
	case 3553294593611130806:
		return 25
	case 8639262066473964150:
		return 26
	case 1703382833545304812:
		return 27
	case 6112860112248786499:
		return 28
	case -4341501217981328394:
		return 29
	case -8395342916587040543:
		return 30
	case 6912441694453382161:
		return 31
	case 5464467312394730418:
		return 32
	case -7784642176186821205:
		return 33
	case -3788331537442685179:
		return 34
	case -1522951607536117149:
		return 35
	case 8914190500520542409:
		return 36
	case 731539533644753052:
		return 37
	case 3421601386088570762:
		return 38
	case -6074819356947405415:
		return 39
	case -6321007591388309905:
		return 40
	case 731311060944096642:
		return 41
	case -4510675770784051183:
		return 42
	case 1237236456312824991:
		return 43
	case 5217623514630858297:
		return 44
	case -5409927581812374498:
		return 45
	case -1185315367526857833:
		return 46
	case 3769868229123868474:
		return 47
	case 7379029982170503008:
		return 48
	case -568369772982215962:
		return 49
	}
	return -1
}

func QSortFilterProxyModel_enable(h uintptr, idx int32) { pg_QSortFilterProxyModel__set_override(h, idx) }

func (this *QSortFilterProxyModel) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QSortFilterProxyModel) OnSetSourceModel(fn func(*QAbstractItemModel)) {
	registerVirt(this.UnsafeGoID(), 7818871925547456643, func(sourceModel uintptr) { fn(UnsafeNewQAbstractItemModel(unsafe.Pointer(sourceModel))) }) // slot "setSourceModel"
}

func (this *QSortFilterProxyModel) OnMapToSource(fn func(*QModelIndex) *QModelIndex) {
	registerVirt(this.UnsafeGoID(), -2333321292016234891, func(proxyIndex uintptr) uintptr { return uintptr(fn(UnsafeNewQModelIndex(unsafe.Pointer(proxyIndex))).UnsafePointer()) }) // slot "mapToSource"
}

func (this *QSortFilterProxyModel) OnMapFromSource(fn func(*QModelIndex) *QModelIndex) {
	registerVirt(this.UnsafeGoID(), -1717009778236344012, func(sourceIndex uintptr) uintptr { return uintptr(fn(UnsafeNewQModelIndex(unsafe.Pointer(sourceIndex))).UnsafePointer()) }) // slot "mapFromSource"
}

func (this *QSortFilterProxyModel) OnMapSelectionToSource(fn func(*QItemSelection) *QItemSelection) {
	registerVirt(this.UnsafeGoID(), 2788628866280405837, func(proxySelection uintptr) uintptr { return uintptr(fn(UnsafeNewQItemSelection(unsafe.Pointer(proxySelection))).UnsafePointer()) }) // slot "mapSelectionToSource"
}

func (this *QSortFilterProxyModel) OnMapSelectionFromSource(fn func(*QItemSelection) *QItemSelection) {
	registerVirt(this.UnsafeGoID(), 9067522407071207004, func(sourceSelection uintptr) uintptr { return uintptr(fn(UnsafeNewQItemSelection(unsafe.Pointer(sourceSelection))).UnsafePointer()) }) // slot "mapSelectionFromSource"
}

func (this *QSortFilterProxyModel) OnFilterAcceptsRow(fn func(int, *QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), -8878517896795061028, func(source_row int32, source_parent uintptr) int32 { return b2i(fn(int(source_row), UnsafeNewQModelIndex(unsafe.Pointer(source_parent)))) }) // slot "filterAcceptsRow"
}

func (this *QSortFilterProxyModel) OnFilterAcceptsColumn(fn func(int, *QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), 436376032338170754, func(source_column int32, source_parent uintptr) int32 { return b2i(fn(int(source_column), UnsafeNewQModelIndex(unsafe.Pointer(source_parent)))) }) // slot "filterAcceptsColumn"
}

func (this *QSortFilterProxyModel) OnLessThan(fn func(*QModelIndex, *QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), -5863021011917066845, func(source_left uintptr, source_right uintptr) int32 { return b2i(fn(UnsafeNewQModelIndex(unsafe.Pointer(source_left)), UnsafeNewQModelIndex(unsafe.Pointer(source_right)))) }) // slot "lessThan"
}

func (this *QSortFilterProxyModel) OnIndex(fn func(int, int, *QModelIndex) *QModelIndex) {
	registerVirt(this.UnsafeGoID(), -8948777187306027381, func(row int32, column int32, parent uintptr) uintptr { return uintptr(fn(int(row), int(column), UnsafeNewQModelIndex(unsafe.Pointer(parent))).UnsafePointer()) }) // slot "index"
}

func (this *QSortFilterProxyModel) OnParent(fn func(*QModelIndex) *QModelIndex) {
	registerVirt(this.UnsafeGoID(), 7371339085820260189, func(child uintptr) uintptr { return uintptr(fn(UnsafeNewQModelIndex(unsafe.Pointer(child))).UnsafePointer()) }) // slot "parent"
}

func (this *QSortFilterProxyModel) OnSibling(fn func(int, int, *QModelIndex) *QModelIndex) {
	registerVirt(this.UnsafeGoID(), 397329346813782521, func(row int32, column int32, idx uintptr) uintptr { return uintptr(fn(int(row), int(column), UnsafeNewQModelIndex(unsafe.Pointer(idx))).UnsafePointer()) }) // slot "sibling"
}

func (this *QSortFilterProxyModel) OnRowCount(fn func(*QModelIndex) int) {
	registerVirt(this.UnsafeGoID(), 1960960171956884010, func(parent uintptr) int32 { return int32(fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)))) }) // slot "rowCount"
}

func (this *QSortFilterProxyModel) OnColumnCount(fn func(*QModelIndex) int) {
	registerVirt(this.UnsafeGoID(), -8256963024669584834, func(parent uintptr) int32 { return int32(fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)))) }) // slot "columnCount"
}

func (this *QSortFilterProxyModel) OnHasChildren(fn func(*QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), 3369452259144514380, func(parent uintptr) int32 { return b2i(fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)))) }) // slot "hasChildren"
}

func (this *QSortFilterProxyModel) OnData(fn func(*QModelIndex, int) *QVariant) {
	registerVirt(this.UnsafeGoID(), -8837375942098466299, func(index uintptr, role int32) uintptr { return uintptr(fn(UnsafeNewQModelIndex(unsafe.Pointer(index)), int(role)).UnsafePointer()) }) // slot "data"
}

func (this *QSortFilterProxyModel) OnSetData(fn func(*QModelIndex, *QVariant, int) bool) {
	registerVirt(this.UnsafeGoID(), 35191454575202219, func(index uintptr, value uintptr, role int32) int32 { return b2i(fn(UnsafeNewQModelIndex(unsafe.Pointer(index)), UnsafeNewQVariant(unsafe.Pointer(value)), int(role))) }) // slot "setData"
}

func (this *QSortFilterProxyModel) OnHeaderData(fn func(int, Orientation, int) *QVariant) {
	registerVirt(this.UnsafeGoID(), -3297140021896693688, func(section int32, orientation int32, role int32) uintptr { return uintptr(fn(int(section), Orientation(orientation), int(role)).UnsafePointer()) }) // slot "headerData"
}

func (this *QSortFilterProxyModel) OnSetHeaderData(fn func(int, Orientation, *QVariant, int) bool) {
	registerVirt(this.UnsafeGoID(), 4852200306590054850, func(section int32, orientation int32, value uintptr, role int32) int32 { return b2i(fn(int(section), Orientation(orientation), UnsafeNewQVariant(unsafe.Pointer(value)), int(role))) }) // slot "setHeaderData"
}

func (this *QSortFilterProxyModel) OnMimeData(fn func([]*QModelIndex) *QMimeData) {
	registerVirt(this.UnsafeGoID(), 5175297091024974749, func(indexes uintptr, indexes_len int64) uintptr { return uintptr(fn(func() []*QModelIndex { if indexes == 0 || indexes_len == 0 { return nil }; __sl := unsafe.Slice((*unsafe.Pointer)(unsafe.Pointer(indexes)), int(indexes_len)); __o := make([]*QModelIndex, int(indexes_len)); for __i := range __o { __o[__i] = UnsafeNewQModelIndex(__sl[__i]) }; return __o }()).UnsafePointer()) }) // slot "mimeData"
}

func (this *QSortFilterProxyModel) OnDropMimeData(fn func(*QMimeData, DropAction, int, int, *QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), -9049994684305682212, func(data uintptr, action int32, row int32, column int32, parent uintptr) int32 { return b2i(fn(UnsafeNewQMimeData(unsafe.Pointer(data)), DropAction(action), int(row), int(column), UnsafeNewQModelIndex(unsafe.Pointer(parent)))) }) // slot "dropMimeData"
}

func (this *QSortFilterProxyModel) OnInsertRows(fn func(int, int, *QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), -2775729825962395603, func(row int32, count int32, parent uintptr) int32 { return b2i(fn(int(row), int(count), UnsafeNewQModelIndex(unsafe.Pointer(parent)))) }) // slot "insertRows"
}

func (this *QSortFilterProxyModel) OnInsertColumns(fn func(int, int, *QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), -8973293846230792237, func(column int32, count int32, parent uintptr) int32 { return b2i(fn(int(column), int(count), UnsafeNewQModelIndex(unsafe.Pointer(parent)))) }) // slot "insertColumns"
}

func (this *QSortFilterProxyModel) OnRemoveRows(fn func(int, int, *QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), 5104072549182222656, func(row int32, count int32, parent uintptr) int32 { return b2i(fn(int(row), int(count), UnsafeNewQModelIndex(unsafe.Pointer(parent)))) }) // slot "removeRows"
}

func (this *QSortFilterProxyModel) OnRemoveColumns(fn func(int, int, *QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), 2078548737056540420, func(column int32, count int32, parent uintptr) int32 { return b2i(fn(int(column), int(count), UnsafeNewQModelIndex(unsafe.Pointer(parent)))) }) // slot "removeColumns"
}

func (this *QSortFilterProxyModel) OnFetchMore(fn func(*QModelIndex)) {
	registerVirt(this.UnsafeGoID(), 3553294593611130806, func(parent uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(parent))) }) // slot "fetchMore"
}

func (this *QSortFilterProxyModel) OnCanFetchMore(fn func(*QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), 8639262066473964150, func(parent uintptr) int32 { return b2i(fn(UnsafeNewQModelIndex(unsafe.Pointer(parent)))) }) // slot "canFetchMore"
}

func (this *QSortFilterProxyModel) OnFlags(fn func(*QModelIndex) ItemFlag) {
	registerVirt(this.UnsafeGoID(), 1703382833545304812, func(index uintptr) int32 { return int32(fn(UnsafeNewQModelIndex(unsafe.Pointer(index)))) }) // slot "flags"
}

func (this *QSortFilterProxyModel) OnBuddy(fn func(*QModelIndex) *QModelIndex) {
	registerVirt(this.UnsafeGoID(), 6112860112248786499, func(index uintptr) uintptr { return uintptr(fn(UnsafeNewQModelIndex(unsafe.Pointer(index))).UnsafePointer()) }) // slot "buddy"
}

func (this *QSortFilterProxyModel) OnMatch(fn func(*QModelIndex, int, *QVariant, int, MatchFlag) []*QModelIndex) {
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

func (this *QSortFilterProxyModel) OnSpan(fn func(*QModelIndex) *QSize) {
	registerVirt(this.UnsafeGoID(), -8395342916587040543, func(index uintptr) uintptr { return uintptr(fn(UnsafeNewQModelIndex(unsafe.Pointer(index))).UnsafePointer()) }) // slot "span"
}

func (this *QSortFilterProxyModel) OnSort(fn func(int, SortOrder)) {
	registerVirt(this.UnsafeGoID(), 6912441694453382161, func(column int32, order int32) { fn(int(column), SortOrder(order)) }) // slot "sort"
}

func (this *QSortFilterProxyModel) OnMimeTypes(fn func() []string) {
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

func (this *QSortFilterProxyModel) OnSupportedDropActions(fn func() DropAction) {
	registerVirt(this.UnsafeGoID(), -7784642176186821205, func() int32 { return int32(fn()) }) // slot "supportedDropActions"
}

func (this *QSortFilterProxyModel) OnSubmit(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -3788331537442685179, func() int32 { return b2i(fn()) }) // slot "submit"
}

func (this *QSortFilterProxyModel) OnRevert(fn func()) {
	registerVirt(this.UnsafeGoID(), -1522951607536117149, func() { fn() }) // slot "revert"
}

func (this *QSortFilterProxyModel) OnClearItemData(fn func(*QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), 8914190500520542409, func(index uintptr) int32 { return b2i(fn(UnsafeNewQModelIndex(unsafe.Pointer(index)))) }) // slot "clearItemData"
}

func (this *QSortFilterProxyModel) OnCanDropMimeData(fn func(*QMimeData, DropAction, int, int, *QModelIndex) bool) {
	registerVirt(this.UnsafeGoID(), 731539533644753052, func(data uintptr, action int32, row int32, column int32, parent uintptr) int32 { return b2i(fn(UnsafeNewQMimeData(unsafe.Pointer(data)), DropAction(action), int(row), int(column), UnsafeNewQModelIndex(unsafe.Pointer(parent)))) }) // slot "canDropMimeData"
}

func (this *QSortFilterProxyModel) OnSupportedDragActions(fn func() DropAction) {
	registerVirt(this.UnsafeGoID(), 3421601386088570762, func() int32 { return int32(fn()) }) // slot "supportedDragActions"
}

func (this *QSortFilterProxyModel) OnMoveRows(fn func(*QModelIndex, int, int, *QModelIndex, int) bool) {
	registerVirt(this.UnsafeGoID(), -6074819356947405415, func(sourceParent uintptr, sourceRow int32, count int32, destinationParent uintptr, destinationChild int32) int32 { return b2i(fn(UnsafeNewQModelIndex(unsafe.Pointer(sourceParent)), int(sourceRow), int(count), UnsafeNewQModelIndex(unsafe.Pointer(destinationParent)), int(destinationChild))) }) // slot "moveRows"
}

func (this *QSortFilterProxyModel) OnMoveColumns(fn func(*QModelIndex, int, int, *QModelIndex, int) bool) {
	registerVirt(this.UnsafeGoID(), -6321007591388309905, func(sourceParent uintptr, sourceColumn int32, count int32, destinationParent uintptr, destinationChild int32) int32 { return b2i(fn(UnsafeNewQModelIndex(unsafe.Pointer(sourceParent)), int(sourceColumn), int(count), UnsafeNewQModelIndex(unsafe.Pointer(destinationParent)), int(destinationChild))) }) // slot "moveColumns"
}

func (this *QSortFilterProxyModel) OnMultiData(fn func(*QModelIndex, *QModelRoleDataSpan)) {
	registerVirt(this.UnsafeGoID(), 731311060944096642, func(index uintptr, roleDataSpan uintptr) { fn(UnsafeNewQModelIndex(unsafe.Pointer(index)), UnsafeNewQModelRoleDataSpan(unsafe.Pointer(roleDataSpan))) }) // slot "multiData"
}

func (this *QSortFilterProxyModel) OnResetInternalData(fn func()) {
	registerVirt(this.UnsafeGoID(), -4510675770784051183, func() { fn() }) // slot "resetInternalData"
}

func (this *QSortFilterProxyModel) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QSortFilterProxyModel) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QSortFilterProxyModel) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QSortFilterProxyModel) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QSortFilterProxyModel) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QSortFilterProxyModel) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QSortFilterProxyModel) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QSortFilterProxyModel) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QSortFilterProxyModel_delete(this.h)
	}
	this.UnsafeRelease()
}
