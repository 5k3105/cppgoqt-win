#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QAbstractTableModel_base(void* self);
void QAbstractTableModel_chain(void* self, void** out);
void QAbstractTableModel__set_override(void* self, int idx);
void* QAbstractTableModel_fromQObject(void* o);
char* QAbstractTableModel_tr(const char* s, long long* __out_len);
void* QAbstractTableModel_createIndex(void* self, int row, int column);
void* QAbstractTableModel_createIndex2(void* self, int row, int column, unsigned long long id);
void QAbstractTableModel_encodeData(void* self, void** indexes, long long indexes_len, void* stream);
int QAbstractTableModel_decodeData(void* self, int row, int column, void* parent, void* stream);
void QAbstractTableModel_beginInsertRows(void* self, void* parent, int first, int last);
void QAbstractTableModel_endInsertRows(void* self);
void QAbstractTableModel_beginRemoveRows(void* self, void* parent, int first, int last);
void QAbstractTableModel_endRemoveRows(void* self);
int QAbstractTableModel_beginMoveRows(void* self, void* sourceParent, int sourceFirst, int sourceLast, void* destinationParent, int destinationRow);
void QAbstractTableModel_endMoveRows(void* self);
void QAbstractTableModel_beginInsertColumns(void* self, void* parent, int first, int last);
void QAbstractTableModel_endInsertColumns(void* self);
void QAbstractTableModel_beginRemoveColumns(void* self, void* parent, int first, int last);
void QAbstractTableModel_endRemoveColumns(void* self);
int QAbstractTableModel_beginMoveColumns(void* self, void* sourceParent, int sourceFirst, int sourceLast, void* destinationParent, int destinationColumn);
void QAbstractTableModel_endMoveColumns(void* self);
void QAbstractTableModel_beginResetModel(void* self);
void QAbstractTableModel_endResetModel(void* self);
void QAbstractTableModel_changePersistentIndex(void* self, void* from, void* to);
void QAbstractTableModel_changePersistentIndexList(void* self, void** from, long long from_len, void** to, long long to_len);
void** QAbstractTableModel_persistentIndexList(void* self, long long* __out_len);
void* QAbstractTableModel_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QAbstractTableModel_connect2(void* sender, void* signal, void* receiver, void* method);
void* QAbstractTableModel_sender(void* self);
int QAbstractTableModel_senderSignalIndex(void* self);
int QAbstractTableModel_receivers(void* self, const char* signal);
int QAbstractTableModel_isSignalConnected(void* self, void* signal);
char* QAbstractTableModel_tr2(const char* s, const char* c, long long* __out_len);
char* QAbstractTableModel_tr3(const char* s, const char* c, int n, long long* __out_len);
void* QAbstractTableModel_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QAbstractTableModel_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void* QAbstractTableModel_metaObject_default(void* self);
void* QAbstractTableModel_index_default(void* self, int row, int column, void* parent);
void* QAbstractTableModel_sibling_default(void* self, int row, int column, void* idx);
int QAbstractTableModel_dropMimeData_default(void* self, void* data, int action, int row, int column, void* parent);
int QAbstractTableModel_flags_default(void* self, void* index);
int QAbstractTableModel_setData_default(void* self, void* index, void* value, int role);
void* QAbstractTableModel_headerData_default(void* self, int section, int orientation, int role);
int QAbstractTableModel_setHeaderData_default(void* self, int section, int orientation, void* value, int role);
int QAbstractTableModel_clearItemData_default(void* self, void* index);
char** QAbstractTableModel_mimeTypes_default(void* self, long long* __out_len);
void* QAbstractTableModel_mimeData_default(void* self, void** indexes, long long indexes_len);
int QAbstractTableModel_canDropMimeData_default(void* self, void* data, int action, int row, int column, void* parent);
int QAbstractTableModel_supportedDropActions_default(void* self);
int QAbstractTableModel_supportedDragActions_default(void* self);
int QAbstractTableModel_insertRows_default(void* self, int row, int count, void* parent);
int QAbstractTableModel_insertColumns_default(void* self, int column, int count, void* parent);
int QAbstractTableModel_removeRows_default(void* self, int row, int count, void* parent);
int QAbstractTableModel_removeColumns_default(void* self, int column, int count, void* parent);
int QAbstractTableModel_moveRows_default(void* self, void* sourceParent, int sourceRow, int count, void* destinationParent, int destinationChild);
int QAbstractTableModel_moveColumns_default(void* self, void* sourceParent, int sourceColumn, int count, void* destinationParent, int destinationChild);
void QAbstractTableModel_fetchMore_default(void* self, void* parent);
int QAbstractTableModel_canFetchMore_default(void* self, void* parent);
void QAbstractTableModel_sort_default(void* self, int column, int order);
void* QAbstractTableModel_buddy_default(void* self, void* index);
void** QAbstractTableModel_match_default(void* self, void* start, int role, void* value, int hits, int flags, long long* __out_len);
void* QAbstractTableModel_span_default(void* self, void* index);
void QAbstractTableModel_multiData_default(void* self, void* index, void* roleDataSpan);
int QAbstractTableModel_submit_default(void* self);
void QAbstractTableModel_revert_default(void* self);
void QAbstractTableModel_resetInternalData_default(void* self);
int QAbstractTableModel_event_default(void* self, void* event);
int QAbstractTableModel_eventFilter_default(void* self, void* watched, void* event);
void QAbstractTableModel_timerEvent_default(void* self, void* event);
void QAbstractTableModel_childEvent_default(void* self, void* event);
void QAbstractTableModel_customEvent_default(void* self, void* event);
void QAbstractTableModel_connectNotify_default(void* self, void* signal);
void QAbstractTableModel_disconnectNotify_default(void* self, void* signal);
void QAbstractTableModel_connect_headerDataChanged(void* self, long long id);
void QAbstractTableModel_connect_layoutChanged(void* self, long long id);
void QAbstractTableModel_connect_layoutChanged2(void* self, long long id);
void QAbstractTableModel_connect_layoutAboutToBeChanged(void* self, long long id);
void QAbstractTableModel_connect_layoutAboutToBeChanged2(void* self, long long id);
void QAbstractTableModel_connect_rowsAboutToBeInserted(void* self, long long id);
void QAbstractTableModel_connect_rowsInserted(void* self, long long id);
void QAbstractTableModel_connect_rowsAboutToBeRemoved(void* self, long long id);
void QAbstractTableModel_connect_rowsRemoved(void* self, long long id);
void QAbstractTableModel_connect_columnsAboutToBeInserted(void* self, long long id);
void QAbstractTableModel_connect_columnsInserted(void* self, long long id);
void QAbstractTableModel_connect_columnsAboutToBeRemoved(void* self, long long id);
void QAbstractTableModel_connect_columnsRemoved(void* self, long long id);
void QAbstractTableModel_connect_modelAboutToBeReset(void* self, long long id);
void QAbstractTableModel_connect_modelReset(void* self, long long id);
void QAbstractTableModel_connect_rowsAboutToBeMoved(void* self, long long id);
void QAbstractTableModel_connect_rowsMoved(void* self, long long id);
void QAbstractTableModel_connect_columnsAboutToBeMoved(void* self, long long id);
void QAbstractTableModel_connect_columnsMoved(void* self, long long id);
void QAbstractTableModel_connect_destroyed(void* self, long long id);
void QAbstractTableModel_connect_destroyed2(void* self, long long id);
void QAbstractTableModel_connect_objectNameChanged(void* self, long long id);
void QAbstractTableModel__disconnect(long long id);
void* QAbstractTableModel__Go_new(long long id, void* parent);
void QAbstractTableModel_delete(void* self);

#ifdef __cplusplus
}
#endif
