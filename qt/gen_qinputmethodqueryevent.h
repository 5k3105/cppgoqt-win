#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QInputMethodQueryEvent_base(void* self);
void* QInputMethodQueryEvent_new(int queries);
int QInputMethodQueryEvent_queries(void* self);
void QInputMethodQueryEvent_setValue(void* self, int query, void* value);
void* QInputMethodQueryEvent_value(void* self, int query);
int QInputMethodQueryEvent_registerEventType();
int QInputMethodQueryEvent_registerEventType2(int hint);
void QInputMethodQueryEvent_delete(void* self);

#ifdef __cplusplus
}
#endif
