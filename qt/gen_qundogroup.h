#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QUndoGroup_base(void* self);
void QUndoGroup__set_override(void* self, int idx);
void* QUndoGroup_fromQObject(void* o);
void* QUndoGroup_new(long long __goid);
void* QUndoGroup_new2(long long __goid, void* parent);
char* QUndoGroup_tr(const char* s, long long* __out_len);
void QUndoGroup_addStack(void* self, void* stack);
void QUndoGroup_removeStack(void* self, void* stack);
void** QUndoGroup_stacks(void* self, long long* __out_len);
void* QUndoGroup_activeStack(void* self);
void* QUndoGroup_createUndoAction(void* self, void* parent);
void* QUndoGroup_createRedoAction(void* self, void* parent);
int QUndoGroup_canUndo(void* self);
int QUndoGroup_canRedo(void* self);
char* QUndoGroup_undoText(void* self, long long* __out_len);
char* QUndoGroup_redoText(void* self, long long* __out_len);
int QUndoGroup_isClean(void* self);
void QUndoGroup_undo(void* self);
void QUndoGroup_redo(void* self);
void QUndoGroup_setActiveStack(void* self, void* stack);
void QUndoGroup_activeStackChanged(void* self, void* stack);
void QUndoGroup_indexChanged(void* self, int idx);
void QUndoGroup_cleanChanged(void* self, int clean);
void QUndoGroup_canUndoChanged(void* self, int canUndo);
void QUndoGroup_canRedoChanged(void* self, int canRedo);
void QUndoGroup_undoTextChanged(void* self, const char* undoText, long long undoText_len);
void QUndoGroup_redoTextChanged(void* self, const char* redoText, long long redoText_len);
void* QUndoGroup_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QUndoGroup_connect2(void* sender, void* signal, void* receiver, void* method);
void* QUndoGroup_sender(void* self);
int QUndoGroup_senderSignalIndex(void* self);
int QUndoGroup_receivers(void* self, const char* signal);
int QUndoGroup_isSignalConnected(void* self, void* signal);
char* QUndoGroup_tr2(const char* s, const char* c, long long* __out_len);
char* QUndoGroup_tr3(const char* s, const char* c, int n, long long* __out_len);
void* QUndoGroup_createUndoAction2(void* self, void* parent, const char* prefix, long long prefix_len);
void* QUndoGroup_createRedoAction2(void* self, void* parent, const char* prefix, long long prefix_len);
void* QUndoGroup_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QUndoGroup_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void* QUndoGroup_metaObject_default(void* self);
int QUndoGroup_event_default(void* self, void* event);
int QUndoGroup_eventFilter_default(void* self, void* watched, void* event);
void QUndoGroup_timerEvent_default(void* self, void* event);
void QUndoGroup_childEvent_default(void* self, void* event);
void QUndoGroup_customEvent_default(void* self, void* event);
void QUndoGroup_connectNotify_default(void* self, void* signal);
void QUndoGroup_disconnectNotify_default(void* self, void* signal);
void QUndoGroup_connect_activeStackChanged(void* self, long long id);
void QUndoGroup_connect_indexChanged(void* self, long long id);
void QUndoGroup_connect_cleanChanged(void* self, long long id);
void QUndoGroup_connect_canUndoChanged(void* self, long long id);
void QUndoGroup_connect_canRedoChanged(void* self, long long id);
void QUndoGroup_connect_undoTextChanged(void* self, long long id);
void QUndoGroup_connect_redoTextChanged(void* self, long long id);
void QUndoGroup_connect_destroyed(void* self, long long id);
void QUndoGroup_connect_destroyed2(void* self, long long id);
void QUndoGroup_connect_objectNameChanged(void* self, long long id);
void QUndoGroup__disconnect(long long id);
void* QUndoGroup__Go_new(long long id, void* parent);
void QUndoGroup_delete(void* self);

#ifdef __cplusplus
}
#endif
