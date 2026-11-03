#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QAccessibleAnnouncementEvent_base(void* self);
void* QAccessibleAnnouncementEvent_new(void* object, const char* message, long long message_len);
void* QAccessibleAnnouncementEvent_new2(void* iface, const char* message, long long message_len);
char* QAccessibleAnnouncementEvent_message(void* self, long long* __out_len);
int QAccessibleAnnouncementEvent_politeness(void* self);
void QAccessibleAnnouncementEvent_setPoliteness(void* self, int politeness);
void QAccessibleAnnouncementEvent_delete(void* self);

#ifdef __cplusplus
}
#endif
