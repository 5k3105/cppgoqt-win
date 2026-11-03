#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QAccessibleTextUpdateEvent_base(void* self);
void QAccessibleTextUpdateEvent_chain(void* self, void** out);
void* QAccessibleTextUpdateEvent_new(void* obj, int position, const char* oldText, long long oldText_len, const char* text, long long text_len);
void* QAccessibleTextUpdateEvent_new2(void* iface, int position, const char* oldText, long long oldText_len, const char* text, long long text_len);
char* QAccessibleTextUpdateEvent_textRemoved(void* self, long long* __out_len);
char* QAccessibleTextUpdateEvent_textInserted(void* self, long long* __out_len);
int QAccessibleTextUpdateEvent_changePosition(void* self);
void QAccessibleTextUpdateEvent_delete(void* self);

#ifdef __cplusplus
}
#endif
