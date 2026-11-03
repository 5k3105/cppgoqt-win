#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void QUndoCommand__set_override(void* self, int idx);
void* QUndoCommand_new(long long __goid);
void* QUndoCommand_new2(long long __goid, const char* text, long long text_len);
void* QUndoCommand_new3(long long __goid, void* parent);
void* QUndoCommand_new4(long long __goid, const char* text, long long text_len, void* parent);
void QUndoCommand_undo(void* self);
void QUndoCommand_redo(void* self);
char* QUndoCommand_text(void* self, long long* __out_len);
char* QUndoCommand_actionText(void* self, long long* __out_len);
void QUndoCommand_setText(void* self, const char* text, long long text_len);
int QUndoCommand_isObsolete(void* self);
void QUndoCommand_setObsolete(void* self, int obsolete);
int QUndoCommand_id(void* self);
int QUndoCommand_mergeWith(void* self, void* other);
int QUndoCommand_childCount(void* self);
void* QUndoCommand_child(void* self, int index);
void QUndoCommand_undo_default(void* self);
void QUndoCommand_redo_default(void* self);
int QUndoCommand_id_default(void* self);
int QUndoCommand_mergeWith_default(void* self, void* other);
void* QUndoCommand__Go_new(long long id);
void QUndoCommand_delete(void* self);

#ifdef __cplusplus
}
#endif
