#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QStatusTipEvent_base(void* self);
void* QStatusTipEvent_new(const char* tip, long long tip_len);
char* QStatusTipEvent_tip(void* self, long long* __out_len);
int QStatusTipEvent_registerEventType();
int QStatusTipEvent_registerEventType2(int hint);
void QStatusTipEvent_delete(void* self);

#ifdef __cplusplus
}
#endif
