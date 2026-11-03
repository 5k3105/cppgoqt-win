#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

int QDesktopServices_openUrl(void* url);
void QDesktopServices_setUrlHandler(const char* scheme, long long scheme_len, void* receiver, const char* method);
void QDesktopServices_unsetUrlHandler(const char* scheme, long long scheme_len);
void QDesktopServices_delete(void* self);

#ifdef __cplusplus
}
#endif
