#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void QTextObjectInterface__set_override(void* self, int idx);
void* QTextObjectInterface_intrinsicSize(void* self, void* doc, int posInDocument, void* format);
void QTextObjectInterface_drawObject(void* self, void* painter, void* rect, void* doc, int posInDocument, void* format);
void* QTextObjectInterface_OperatorAssign(void* self, void* arg1);
void* QTextObjectInterface__Go_new(long long id);
void QTextObjectInterface_delete(void* self);

#ifdef __cplusplus
}
#endif
