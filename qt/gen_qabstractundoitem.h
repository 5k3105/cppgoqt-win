#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void QAbstractUndoItem__set_override(void* self, int idx);
void QAbstractUndoItem_undo(void* self);
void QAbstractUndoItem_redo(void* self);
void* QAbstractUndoItem_OperatorAssign(void* self, void* arg1);
void* QAbstractUndoItem__Go_new(long long id);
void QAbstractUndoItem_delete(void* self);

#ifdef __cplusplus
}
#endif
