#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void QItemEditorFactory__set_override(void* self, int idx);
void* QItemEditorFactory_new(long long __goid);
void* QItemEditorFactory_new2(long long __goid, void* arg1);
void* QItemEditorFactory_createEditor(void* self, int userType, void* parent);
void* QItemEditorFactory_valuePropertyName(void* self, int userType);
void QItemEditorFactory_registerEditor(void* self, int userType, void* creator);
void* QItemEditorFactory_defaultFactory();
void QItemEditorFactory_setDefaultFactory(void* factory);
void* QItemEditorFactory_OperatorAssign(void* self, void* arg1);
void* QItemEditorFactory_createEditor_default(void* self, int userType, void* parent);
void* QItemEditorFactory_valuePropertyName_default(void* self, int userType);
void* QItemEditorFactory__Go_new(long long id);
void QItemEditorFactory_delete(void* self);

#ifdef __cplusplus
}
#endif
