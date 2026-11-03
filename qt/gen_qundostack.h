#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QUndoStack_base(void* self);
void QUndoStack__set_override(void* self, int idx);
void* QUndoStack_fromQObject(void* o);
void* QUndoStack_new(long long __goid);
void* QUndoStack_new2(long long __goid, void* parent);
char* QUndoStack_tr(const char* s, long long* __out_len);
void QUndoStack_clear(void* self);
void QUndoStack_push(void* self, void* cmd);
int QUndoStack_canUndo(void* self);
int QUndoStack_canRedo(void* self);
char* QUndoStack_undoText(void* self, long long* __out_len);
char* QUndoStack_redoText(void* self, long long* __out_len);
int QUndoStack_count(void* self);
int QUndoStack_index(void* self);
char* QUndoStack_text(void* self, int idx, long long* __out_len);
void* QUndoStack_createUndoAction(void* self, void* parent);
void* QUndoStack_createRedoAction(void* self, void* parent);
int QUndoStack_isActive(void* self);
int QUndoStack_isClean(void* self);
int QUndoStack_cleanIndex(void* self);
void QUndoStack_beginMacro(void* self, const char* text, long long text_len);
void QUndoStack_endMacro(void* self);
void QUndoStack_setUndoLimit(void* self, int limit);
int QUndoStack_undoLimit(void* self);
void* QUndoStack_command(void* self, int index);
void QUndoStack_setClean(void* self);
void QUndoStack_resetClean(void* self);
void QUndoStack_setIndex(void* self, int idx);
void QUndoStack_undo(void* self);
void QUndoStack_redo(void* self);
void QUndoStack_setActive(void* self);
void QUndoStack_indexChanged(void* self, int idx);
void QUndoStack_cleanChanged(void* self, int clean);
void QUndoStack_canUndoChanged(void* self, int canUndo);
void QUndoStack_canRedoChanged(void* self, int canRedo);
void QUndoStack_undoTextChanged(void* self, const char* undoText, long long undoText_len);
void QUndoStack_redoTextChanged(void* self, const char* redoText, long long redoText_len);
void* QUndoStack_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QUndoStack_connect2(void* sender, void* signal, void* receiver, void* method);
void* QUndoStack_sender(void* self);
int QUndoStack_senderSignalIndex(void* self);
int QUndoStack_receivers(void* self, const char* signal);
int QUndoStack_isSignalConnected(void* self, void* signal);
char* QUndoStack_tr2(const char* s, const char* c, long long* __out_len);
char* QUndoStack_tr3(const char* s, const char* c, int n, long long* __out_len);
void* QUndoStack_createUndoAction2(void* self, void* parent, const char* prefix, long long prefix_len);
void* QUndoStack_createRedoAction2(void* self, void* parent, const char* prefix, long long prefix_len);
void QUndoStack_setActive2(void* self, int active);
void* QUndoStack_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QUndoStack_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void* QUndoStack_metaObject_default(void* self);
int QUndoStack_event_default(void* self, void* event);
int QUndoStack_eventFilter_default(void* self, void* watched, void* event);
void QUndoStack_timerEvent_default(void* self, void* event);
void QUndoStack_childEvent_default(void* self, void* event);
void QUndoStack_customEvent_default(void* self, void* event);
void QUndoStack_connectNotify_default(void* self, void* signal);
void QUndoStack_disconnectNotify_default(void* self, void* signal);
void QUndoStack_connect_indexChanged(void* self, long long id);
void QUndoStack_connect_cleanChanged(void* self, long long id);
void QUndoStack_connect_canUndoChanged(void* self, long long id);
void QUndoStack_connect_canRedoChanged(void* self, long long id);
void QUndoStack_connect_undoTextChanged(void* self, long long id);
void QUndoStack_connect_redoTextChanged(void* self, long long id);
void QUndoStack_connect_destroyed(void* self, long long id);
void QUndoStack_connect_destroyed2(void* self, long long id);
void QUndoStack_connect_objectNameChanged(void* self, long long id);
void QUndoStack__disconnect(long long id);
void* QUndoStack__Go_new(long long id, void* parent);
void QUndoStack_delete(void* self);

#ifdef __cplusplus
}
#endif
