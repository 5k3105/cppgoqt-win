#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QRegularExpressionValidator_base(void* self);
void QRegularExpressionValidator_chain(void* self, void** out);
void QRegularExpressionValidator__set_override(void* self, int idx);
void* QRegularExpressionValidator_fromQObject(void* o);
void* QRegularExpressionValidator_new(long long __goid);
void* QRegularExpressionValidator_new2(long long __goid, void* re);
void* QRegularExpressionValidator_new3(long long __goid, void* parent);
void* QRegularExpressionValidator_new4(long long __goid, void* re, void* parent);
char* QRegularExpressionValidator_tr(const char* s, long long* __out_len);
void* QRegularExpressionValidator_regularExpression(void* self);
void QRegularExpressionValidator_setRegularExpression(void* self, void* re);
void QRegularExpressionValidator_regularExpressionChanged(void* self, void* re);
void* QRegularExpressionValidator_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QRegularExpressionValidator_connect2(void* sender, void* signal, void* receiver, void* method);
void* QRegularExpressionValidator_sender(void* self);
int QRegularExpressionValidator_senderSignalIndex(void* self);
int QRegularExpressionValidator_receivers(void* self, const char* signal);
int QRegularExpressionValidator_isSignalConnected(void* self, void* signal);
char* QRegularExpressionValidator_tr2(const char* s, const char* c, long long* __out_len);
char* QRegularExpressionValidator_tr3(const char* s, const char* c, int n, long long* __out_len);
void* QRegularExpressionValidator_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QRegularExpressionValidator_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void* QRegularExpressionValidator_metaObject_default(void* self);
int QRegularExpressionValidator_validate_default(void* self, void* input, int* pos);
void QRegularExpressionValidator_fixup_default(void* self, void* arg1);
int QRegularExpressionValidator_event_default(void* self, void* event);
int QRegularExpressionValidator_eventFilter_default(void* self, void* watched, void* event);
void QRegularExpressionValidator_timerEvent_default(void* self, void* event);
void QRegularExpressionValidator_childEvent_default(void* self, void* event);
void QRegularExpressionValidator_customEvent_default(void* self, void* event);
void QRegularExpressionValidator_connectNotify_default(void* self, void* signal);
void QRegularExpressionValidator_disconnectNotify_default(void* self, void* signal);
void QRegularExpressionValidator_connect_regularExpressionChanged(void* self, long long id);
void QRegularExpressionValidator_connect_changed(void* self, long long id);
void QRegularExpressionValidator_connect_destroyed(void* self, long long id);
void QRegularExpressionValidator_connect_destroyed2(void* self, long long id);
void QRegularExpressionValidator_connect_objectNameChanged(void* self, long long id);
void QRegularExpressionValidator__disconnect(long long id);
void* QRegularExpressionValidator__Go_new(long long id, void* parent);
void QRegularExpressionValidator_delete(void* self);

#ifdef __cplusplus
}
#endif
