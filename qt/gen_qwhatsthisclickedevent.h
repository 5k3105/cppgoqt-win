#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QWhatsThisClickedEvent_base(void* self);
void* QWhatsThisClickedEvent_new(const char* href, long long href_len);
char* QWhatsThisClickedEvent_href(void* self, long long* __out_len);
int QWhatsThisClickedEvent_registerEventType();
int QWhatsThisClickedEvent_registerEventType2(int hint);
void QWhatsThisClickedEvent_delete(void* self);

#ifdef __cplusplus
}
#endif
