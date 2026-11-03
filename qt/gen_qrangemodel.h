#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QRangeModel_base(void* self);
void QRangeModel_chain(void* self, void** out);
void* QRangeModel_fromQObject(void* o);
char* QRangeModel_tr(const char* s, long long* __out_len);
void* QRangeModel_index(void* self, int row, int column);
void* QRangeModel_parent(void* self, void* child);
void* QRangeModel_sibling(void* self, int row, int column, void* index);
int QRangeModel_rowCount(void* self);
int QRangeModel_columnCount(void* self);
int QRangeModel_insertColumns(void* self, int column, int count);
int QRangeModel_removeColumns(void* self, int column, int count);
int QRangeModel_moveColumns(void* self, void* sourceParent, int sourceColumn, int count, void* destParent, int destColumn);
int QRangeModel_insertRows(void* self, int row, int count);
int QRangeModel_removeRows(void* self, int row, int count);
int QRangeModel_moveRows(void* self, void* sourceParent, int sourceRow, int count, void* destParent, int destRow);
void QRangeModel_setRoleNames(void* self, int* names_keys, void* names_vals, long long names_len);
void QRangeModel_resetRoleNames(void* self);
int QRangeModel_hasChildren(void* self);
int QRangeModel_autoConnectPolicy(void* self);
void QRangeModel_setAutoConnectPolicy(void* self, int policy);
void QRangeModel_roleNamesChanged(void* self);
void QRangeModel_autoConnectPolicyChanged(void* self, int policy);
void* QRangeModel_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QRangeModel_connect2(void* sender, void* signal, void* receiver, void* method);
char* QRangeModel_tr2(const char* s, const char* c, long long* __out_len);
char* QRangeModel_tr3(const char* s, const char* c, int n, long long* __out_len);
void* QRangeModel_index2(void* self, int row, int column, void* parent);
int QRangeModel_rowCount2(void* self, void* parent);
int QRangeModel_columnCount2(void* self, void* parent);
int QRangeModel_insertColumns2(void* self, int column, int count, void* parent);
int QRangeModel_removeColumns2(void* self, int column, int count, void* parent);
int QRangeModel_insertRows2(void* self, int row, int count, void* parent);
int QRangeModel_removeRows2(void* self, int row, int count, void* parent);
int QRangeModel_hasChildren2(void* self, void* parent);
void* QRangeModel_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QRangeModel_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void QRangeModel_connect_roleNamesChanged(void* self, long long id);
void QRangeModel_connect_autoConnectPolicyChanged(void* self, long long id);
void QRangeModel_connect_headerDataChanged(void* self, long long id);
void QRangeModel_connect_layoutChanged(void* self, long long id);
void QRangeModel_connect_layoutChanged2(void* self, long long id);
void QRangeModel_connect_layoutAboutToBeChanged(void* self, long long id);
void QRangeModel_connect_layoutAboutToBeChanged2(void* self, long long id);
void QRangeModel_connect_rowsAboutToBeInserted(void* self, long long id);
void QRangeModel_connect_rowsInserted(void* self, long long id);
void QRangeModel_connect_rowsAboutToBeRemoved(void* self, long long id);
void QRangeModel_connect_rowsRemoved(void* self, long long id);
void QRangeModel_connect_columnsAboutToBeInserted(void* self, long long id);
void QRangeModel_connect_columnsInserted(void* self, long long id);
void QRangeModel_connect_columnsAboutToBeRemoved(void* self, long long id);
void QRangeModel_connect_columnsRemoved(void* self, long long id);
void QRangeModel_connect_modelAboutToBeReset(void* self, long long id);
void QRangeModel_connect_modelReset(void* self, long long id);
void QRangeModel_connect_rowsAboutToBeMoved(void* self, long long id);
void QRangeModel_connect_rowsMoved(void* self, long long id);
void QRangeModel_connect_columnsAboutToBeMoved(void* self, long long id);
void QRangeModel_connect_columnsMoved(void* self, long long id);
void QRangeModel_connect_destroyed(void* self, long long id);
void QRangeModel_connect_destroyed2(void* self, long long id);
void QRangeModel_connect_objectNameChanged(void* self, long long id);
void QRangeModel__disconnect(long long id);
void QRangeModel_delete(void* self);

#ifdef __cplusplus
}
#endif
