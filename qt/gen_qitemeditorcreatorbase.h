#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void QItemEditorCreatorBase__set_override(void* self, int idx);
void* QItemEditorCreatorBase_createWidget(void* self, void* parent);
void* QItemEditorCreatorBase_valuePropertyName(void* self);
void* QItemEditorCreatorBase_OperatorAssign(void* self, void* arg1);
void* QItemEditorCreatorBase__Go_new(long long id);
void QItemEditorCreatorBase_delete(void* self);

#ifdef __cplusplus
}
#endif
