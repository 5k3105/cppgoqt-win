#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QItemSelectionModel_base(void* self);
void QItemSelectionModel__set_override(void* self, int idx);
void* QItemSelectionModel_fromQObject(void* o);
void* QItemSelectionModel_new(long long __goid);
void* QItemSelectionModel_new2(long long __goid, void* model, void* parent);
void* QItemSelectionModel_new3(long long __goid, void* model);
char* QItemSelectionModel_tr(const char* s, long long* __out_len);
void* QItemSelectionModel_currentIndex(void* self);
int QItemSelectionModel_isSelected(void* self, void* index);
int QItemSelectionModel_isRowSelected(void* self, int row);
int QItemSelectionModel_isColumnSelected(void* self, int column);
int QItemSelectionModel_rowIntersectsSelection(void* self, int row);
int QItemSelectionModel_columnIntersectsSelection(void* self, int column);
int QItemSelectionModel_hasSelection(void* self);
void** QItemSelectionModel_selectedIndexes(void* self, long long* __out_len);
void** QItemSelectionModel_selectedRows(void* self, long long* __out_len);
void** QItemSelectionModel_selectedColumns(void* self, long long* __out_len);
void* QItemSelectionModel_selection(void* self);
void* QItemSelectionModel_model(void* self);
void* QItemSelectionModel_model2(void* self);
void QItemSelectionModel_setModel(void* self, void* model);
void QItemSelectionModel_setCurrentIndex(void* self, void* index, int command);
void QItemSelectionModel_select(void* self, void* index, int command);
void QItemSelectionModel_select2(void* self, void* selection, int command);
void QItemSelectionModel_clear(void* self);
void QItemSelectionModel_reset(void* self);
void QItemSelectionModel_clearSelection(void* self);
void QItemSelectionModel_clearCurrentIndex(void* self);
void QItemSelectionModel_selectionChanged(void* self, void* selected, void* deselected);
void QItemSelectionModel_currentChanged(void* self, void* current, void* previous);
void QItemSelectionModel_currentRowChanged(void* self, void* current, void* previous);
void QItemSelectionModel_currentColumnChanged(void* self, void* current, void* previous);
void QItemSelectionModel_modelChanged(void* self, void* model);
void QItemSelectionModel_emitSelectionChanged(void* self, void* newSelection, void* oldSelection);
void* QItemSelectionModel_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QItemSelectionModel_connect2(void* sender, void* signal, void* receiver, void* method);
void* QItemSelectionModel_sender(void* self);
int QItemSelectionModel_senderSignalIndex(void* self);
int QItemSelectionModel_receivers(void* self, const char* signal);
int QItemSelectionModel_isSignalConnected(void* self, void* signal);
char* QItemSelectionModel_tr2(const char* s, const char* c, long long* __out_len);
char* QItemSelectionModel_tr3(const char* s, const char* c, int n, long long* __out_len);
int QItemSelectionModel_isRowSelected2(void* self, int row, void* parent);
int QItemSelectionModel_isColumnSelected2(void* self, int column, void* parent);
int QItemSelectionModel_rowIntersectsSelection2(void* self, int row, void* parent);
int QItemSelectionModel_columnIntersectsSelection2(void* self, int column, void* parent);
void** QItemSelectionModel_selectedRows2(void* self, int column, long long* __out_len);
void** QItemSelectionModel_selectedColumns2(void* self, int row, long long* __out_len);
void* QItemSelectionModel_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QItemSelectionModel_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void QItemSelectionModel_setCurrentIndex_default(void* self, void* index, int command);
void QItemSelectionModel_select_default(void* self, void* index, int command);
void QItemSelectionModel_select_default2(void* self, void* selection, int command);
void QItemSelectionModel_clear_default(void* self);
void QItemSelectionModel_reset_default(void* self);
void QItemSelectionModel_clearCurrentIndex_default(void* self);
void* QItemSelectionModel_metaObject_default(void* self);
int QItemSelectionModel_event_default(void* self, void* event);
int QItemSelectionModel_eventFilter_default(void* self, void* watched, void* event);
void QItemSelectionModel_timerEvent_default(void* self, void* event);
void QItemSelectionModel_childEvent_default(void* self, void* event);
void QItemSelectionModel_customEvent_default(void* self, void* event);
void QItemSelectionModel_connectNotify_default(void* self, void* signal);
void QItemSelectionModel_disconnectNotify_default(void* self, void* signal);
void QItemSelectionModel_connect_selectionChanged(void* self, long long id);
void QItemSelectionModel_connect_currentChanged(void* self, long long id);
void QItemSelectionModel_connect_currentRowChanged(void* self, long long id);
void QItemSelectionModel_connect_currentColumnChanged(void* self, long long id);
void QItemSelectionModel_connect_modelChanged(void* self, long long id);
void QItemSelectionModel_connect_destroyed(void* self, long long id);
void QItemSelectionModel_connect_destroyed2(void* self, long long id);
void QItemSelectionModel_connect_objectNameChanged(void* self, long long id);
void QItemSelectionModel__disconnect(long long id);
void* QItemSelectionModel__Go_new(long long id);
void QItemSelectionModel_delete(void* self);

#ifdef __cplusplus
}
#endif
