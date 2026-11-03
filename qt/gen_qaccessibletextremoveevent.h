#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QAccessibleTextRemoveEvent_base(void* self);
void QAccessibleTextRemoveEvent_chain(void* self, void** out);
void* QAccessibleTextRemoveEvent_new(void* obj, int position, const char* text, long long text_len);
void* QAccessibleTextRemoveEvent_new2(void* iface, int position, const char* text, long long text_len);
char* QAccessibleTextRemoveEvent_textRemoved(void* self, long long* __out_len);
int QAccessibleTextRemoveEvent_changePosition(void* self);
void QAccessibleTextRemoveEvent_delete(void* self);

#ifdef __cplusplus
}
#endif
