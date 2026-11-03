#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QAccessibleTextInsertEvent_base(void* self);
void QAccessibleTextInsertEvent_chain(void* self, void** out);
void* QAccessibleTextInsertEvent_new(void* obj, int position, const char* text, long long text_len);
void* QAccessibleTextInsertEvent_new2(void* iface, int position, const char* text, long long text_len);
char* QAccessibleTextInsertEvent_textInserted(void* self, long long* __out_len);
int QAccessibleTextInsertEvent_changePosition(void* self);
void QAccessibleTextInsertEvent_delete(void* self);

#ifdef __cplusplus
}
#endif
