#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QAbstractListModel_base(void* self);
void QAbstractListModel_chain(void* self, void** out);
void QAbstractListModel__set_override(void* self, int idx);
void* QAbstractListModel_fromQObject(void* o);
char* QAbstractListModel_tr(const char* s, long long* __out_len);
void* QAbstractListModel_createIndex(void* self, int row, int column);
void* QAbstractListModel_createIndex2(void* self, int row, int column, unsigned long long id);
void QAbstractListModel_encodeData(void* self, void** indexes, long long indexes_len, void* stream);
int QAbstractListModel_decodeData(void* self, int row, int column, void* parent, void* stream);
void QAbstractListModel_beginInsertRows(void* self, void* parent, int first, int last);
void QAbstractListModel_endInsertRows(void* self);
void QAbstractListModel_beginRemoveRows(void* self, void* parent, int first, int last);
void QAbstractListModel_endRemoveRows(void* self);
int QAbstractListModel_beginMoveRows(void* self, void* sourceParent, int sourceFirst, int sourceLast, void* destinationParent, int destinationRow);
void QAbstractListModel_endMoveRows(void* self);
void QAbstractListModel_beginInsertColumns(void* self, void* parent, int first, int last);
void QAbstractListModel_endInsertColumns(void* self);
void QAbstractListModel_beginRemoveColumns(void* self, void* parent, int first, int last);
void QAbstractListModel_endRemoveColumns(void* self);
int QAbstractListModel_beginMoveColumns(void* self, void* sourceParent, int sourceFirst, int sourceLast, void* destinationParent, int destinationColumn);
void QAbstractListModel_endMoveColumns(void* self);
void QAbstractListModel_beginResetModel(void* self);
void QAbstractListModel_endResetModel(void* self);
void QAbstractListModel_changePersistentIndex(void* self, void* from, void* to);
void QAbstractListModel_changePersistentIndexList(void* self, void** from, long long from_len, void** to, long long to_len);
void** QAbstractListModel_persistentIndexList(void* self, long long* __out_len);
void* QAbstractListModel_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QAbstractListModel_connect2(void* sender, void* signal, void* receiver, void* method);
void* QAbstractListModel_sender(void* self);
int QAbstractListModel_senderSignalIndex(void* self);
int QAbstractListModel_receivers(void* self, const char* signal);
int QAbstractListModel_isSignalConnected(void* self, void* signal);
char* QAbstractListModel_tr2(const char* s, const char* c, long long* __out_len);
char* QAbstractListModel_tr3(const char* s, const char* c, int n, long long* __out_len);
void* QAbstractListModel_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QAbstractListModel_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void* QAbstractListModel_metaObject_default(void* self);
void* QAbstractListModel_index_default(void* self, int row, int column, void* parent);
void* QAbstractListModel_sibling_default(void* self, int row, int column, void* idx);
int QAbstractListModel_dropMimeData_default(void* self, void* data, int action, int row, int column, void* parent);
int QAbstractListModel_flags_default(void* self, void* index);
int QAbstractListModel_setData_default(void* self, void* index, void* value, int role);
void* QAbstractListModel_headerData_default(void* self, int section, int orientation, int role);
int QAbstractListModel_setHeaderData_default(void* self, int section, int orientation, void* value, int role);
int QAbstractListModel_clearItemData_default(void* self, void* index);
char** QAbstractListModel_mimeTypes_default(void* self, long long* __out_len);
void* QAbstractListModel_mimeData_default(void* self, void** indexes, long long indexes_len);
int QAbstractListModel_canDropMimeData_default(void* self, void* data, int action, int row, int column, void* parent);
int QAbstractListModel_supportedDropActions_default(void* self);
int QAbstractListModel_supportedDragActions_default(void* self);
int QAbstractListModel_insertRows_default(void* self, int row, int count, void* parent);
int QAbstractListModel_insertColumns_default(void* self, int column, int count, void* parent);
int QAbstractListModel_removeRows_default(void* self, int row, int count, void* parent);
int QAbstractListModel_removeColumns_default(void* self, int column, int count, void* parent);
int QAbstractListModel_moveRows_default(void* self, void* sourceParent, int sourceRow, int count, void* destinationParent, int destinationChild);
int QAbstractListModel_moveColumns_default(void* self, void* sourceParent, int sourceColumn, int count, void* destinationParent, int destinationChild);
void QAbstractListModel_fetchMore_default(void* self, void* parent);
int QAbstractListModel_canFetchMore_default(void* self, void* parent);
void QAbstractListModel_sort_default(void* self, int column, int order);
void* QAbstractListModel_buddy_default(void* self, void* index);
void** QAbstractListModel_match_default(void* self, void* start, int role, void* value, int hits, int flags, long long* __out_len);
void* QAbstractListModel_span_default(void* self, void* index);
void QAbstractListModel_multiData_default(void* self, void* index, void* roleDataSpan);
int QAbstractListModel_submit_default(void* self);
void QAbstractListModel_revert_default(void* self);
void QAbstractListModel_resetInternalData_default(void* self);
int QAbstractListModel_event_default(void* self, void* event);
int QAbstractListModel_eventFilter_default(void* self, void* watched, void* event);
void QAbstractListModel_timerEvent_default(void* self, void* event);
void QAbstractListModel_childEvent_default(void* self, void* event);
void QAbstractListModel_customEvent_default(void* self, void* event);
void QAbstractListModel_connectNotify_default(void* self, void* signal);
void QAbstractListModel_disconnectNotify_default(void* self, void* signal);
void QAbstractListModel_connect_headerDataChanged(void* self, long long id);
void QAbstractListModel_connect_layoutChanged(void* self, long long id);
void QAbstractListModel_connect_layoutChanged2(void* self, long long id);
void QAbstractListModel_connect_layoutAboutToBeChanged(void* self, long long id);
void QAbstractListModel_connect_layoutAboutToBeChanged2(void* self, long long id);
void QAbstractListModel_connect_rowsAboutToBeInserted(void* self, long long id);
void QAbstractListModel_connect_rowsInserted(void* self, long long id);
void QAbstractListModel_connect_rowsAboutToBeRemoved(void* self, long long id);
void QAbstractListModel_connect_rowsRemoved(void* self, long long id);
void QAbstractListModel_connect_columnsAboutToBeInserted(void* self, long long id);
void QAbstractListModel_connect_columnsInserted(void* self, long long id);
void QAbstractListModel_connect_columnsAboutToBeRemoved(void* self, long long id);
void QAbstractListModel_connect_columnsRemoved(void* self, long long id);
void QAbstractListModel_connect_modelAboutToBeReset(void* self, long long id);
void QAbstractListModel_connect_modelReset(void* self, long long id);
void QAbstractListModel_connect_rowsAboutToBeMoved(void* self, long long id);
void QAbstractListModel_connect_rowsMoved(void* self, long long id);
void QAbstractListModel_connect_columnsAboutToBeMoved(void* self, long long id);
void QAbstractListModel_connect_columnsMoved(void* self, long long id);
void QAbstractListModel_connect_destroyed(void* self, long long id);
void QAbstractListModel_connect_destroyed2(void* self, long long id);
void QAbstractListModel_connect_objectNameChanged(void* self, long long id);
void QAbstractListModel__disconnect(long long id);
void* QAbstractListModel__Go_new(long long id, void* parent);
void QAbstractListModel_delete(void* self);

#ifdef __cplusplus
}
#endif
