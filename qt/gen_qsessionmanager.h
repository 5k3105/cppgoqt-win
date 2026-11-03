#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QSessionManager_base(void* self);
void* QSessionManager_fromQObject(void* o);
char* QSessionManager_tr(const char* s, long long* __out_len);
char* QSessionManager_sessionId(void* self, long long* __out_len);
char* QSessionManager_sessionKey(void* self, long long* __out_len);
int QSessionManager_allowsInteraction(void* self);
int QSessionManager_allowsErrorInteraction(void* self);
void QSessionManager_release(void* self);
void QSessionManager_cancel(void* self);
void QSessionManager_setRestartHint(void* self, int arg1);
int QSessionManager_restartHint(void* self);
void QSessionManager_setRestartCommand(void* self, char** arg1, long long arg1_len);
char** QSessionManager_restartCommand(void* self, long long* __out_len);
void QSessionManager_setDiscardCommand(void* self, char** arg1, long long arg1_len);
char** QSessionManager_discardCommand(void* self, long long* __out_len);
void QSessionManager_setManagerProperty(void* self, const char* name, long long name_len, const char* value, long long value_len);
void QSessionManager_setManagerProperty2(void* self, const char* name, long long name_len, char** value, long long value_len);
int QSessionManager_isPhase2(void* self);
void QSessionManager_requestPhase2(void* self);
void* QSessionManager_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QSessionManager_connect2(void* sender, void* signal, void* receiver, void* method);
char* QSessionManager_tr2(const char* s, const char* c, long long* __out_len);
char* QSessionManager_tr3(const char* s, const char* c, int n, long long* __out_len);
void* QSessionManager_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QSessionManager_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void QSessionManager_connect_destroyed(void* self, long long id);
void QSessionManager_connect_destroyed2(void* self, long long id);
void QSessionManager_connect_objectNameChanged(void* self, long long id);
void QSessionManager__disconnect(long long id);
void QSessionManager_delete(void* self);

#ifdef __cplusplus
}
#endif
