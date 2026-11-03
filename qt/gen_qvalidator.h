#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QValidator_base(void* self);
void QValidator__set_override(void* self, int idx);
void* QValidator_fromQObject(void* o);
char* QValidator_tr(const char* s, long long* __out_len);
void QValidator_setLocale(void* self, void* locale);
void* QValidator_locale(void* self);
int QValidator_validate(void* self, void* arg1, int* arg2);
void QValidator_fixup(void* self, void* arg1);
void QValidator_changed(void* self);
void* QValidator_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QValidator_connect2(void* sender, void* signal, void* receiver, void* method);
void* QValidator_sender(void* self);
int QValidator_senderSignalIndex(void* self);
int QValidator_receivers(void* self, const char* signal);
int QValidator_isSignalConnected(void* self, void* signal);
char* QValidator_tr2(const char* s, const char* c, long long* __out_len);
char* QValidator_tr3(const char* s, const char* c, int n, long long* __out_len);
void* QValidator_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QValidator_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void QValidator_fixup_default(void* self, void* arg1);
void* QValidator_metaObject_default(void* self);
int QValidator_event_default(void* self, void* event);
int QValidator_eventFilter_default(void* self, void* watched, void* event);
void QValidator_timerEvent_default(void* self, void* event);
void QValidator_childEvent_default(void* self, void* event);
void QValidator_customEvent_default(void* self, void* event);
void QValidator_connectNotify_default(void* self, void* signal);
void QValidator_disconnectNotify_default(void* self, void* signal);
void QValidator_connect_changed(void* self, long long id);
void QValidator_connect_destroyed(void* self, long long id);
void QValidator_connect_destroyed2(void* self, long long id);
void QValidator_connect_objectNameChanged(void* self, long long id);
void QValidator__disconnect(long long id);
void* QValidator__Go_new(long long id, void* parent);
void QValidator_delete(void* self);

#ifdef __cplusplus
}
#endif
