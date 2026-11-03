#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QInputMethodEvent_base(void* self);
void QInputMethodEvent__set_override(void* self, int idx);
void* QInputMethodEvent_new(long long __goid);
void* QInputMethodEvent_new2(long long __goid, const char* preeditText, long long preeditText_len, void** attributes, long long attributes_len);
void QInputMethodEvent_setCommitString(void* self, const char* commitString, long long commitString_len);
void** QInputMethodEvent_attributes(void* self, long long* __out_len);
char* QInputMethodEvent_preeditString(void* self, long long* __out_len);
char* QInputMethodEvent_commitString(void* self, long long* __out_len);
int QInputMethodEvent_replacementStart(void* self);
int QInputMethodEvent_replacementLength(void* self);
int QInputMethodEvent_registerEventType();
void QInputMethodEvent_setCommitString2(void* self, const char* commitString, long long commitString_len, int replaceFrom);
void QInputMethodEvent_setCommitString3(void* self, const char* commitString, long long commitString_len, int replaceFrom, int replaceLength);
int QInputMethodEvent_registerEventType2(int hint);
void* QInputMethodEvent_clone_default(void* self);
void QInputMethodEvent_setAccepted_default(void* self, int accepted);
void* QInputMethodEvent__Go_new(long long id);
void QInputMethodEvent_delete(void* self);

#ifdef __cplusplus
}
#endif
