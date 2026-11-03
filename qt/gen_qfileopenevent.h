#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QFileOpenEvent_base(void* self);
void* QFileOpenEvent_new(const char* file, long long file_len);
void* QFileOpenEvent_new2(void* url);
char* QFileOpenEvent_file(void* self, long long* __out_len);
void* QFileOpenEvent_url(void* self);
int QFileOpenEvent_openFile(void* self, void* file, int flags);
int QFileOpenEvent_registerEventType();
int QFileOpenEvent_registerEventType2(int hint);
void QFileOpenEvent_delete(void* self);

#ifdef __cplusplus
}
#endif
