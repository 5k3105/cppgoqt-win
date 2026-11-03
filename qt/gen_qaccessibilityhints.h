#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QAccessibilityHints_base(void* self);
void QAccessibilityHints__set_override(void* self, int idx);
void* QAccessibilityHints_fromQObject(void* o);
void* QAccessibilityHints_new(long long __goid);
void* QAccessibilityHints_new2(long long __goid, void* parent);
char* QAccessibilityHints_tr(const char* s, long long* __out_len);
int QAccessibilityHints_contrastPreference(void* self);
void QAccessibilityHints_contrastPreferenceChanged(void* self, int contrastPreference);
void* QAccessibilityHints_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QAccessibilityHints_connect2(void* sender, void* signal, void* receiver, void* method);
void* QAccessibilityHints_sender(void* self);
int QAccessibilityHints_senderSignalIndex(void* self);
int QAccessibilityHints_receivers(void* self, const char* signal);
int QAccessibilityHints_isSignalConnected(void* self, void* signal);
char* QAccessibilityHints_tr2(const char* s, const char* c, long long* __out_len);
char* QAccessibilityHints_tr3(const char* s, const char* c, int n, long long* __out_len);
void* QAccessibilityHints_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QAccessibilityHints_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void* QAccessibilityHints_metaObject_default(void* self);
int QAccessibilityHints_event_default(void* self, void* event);
int QAccessibilityHints_eventFilter_default(void* self, void* watched, void* event);
void QAccessibilityHints_timerEvent_default(void* self, void* event);
void QAccessibilityHints_childEvent_default(void* self, void* event);
void QAccessibilityHints_customEvent_default(void* self, void* event);
void QAccessibilityHints_connectNotify_default(void* self, void* signal);
void QAccessibilityHints_disconnectNotify_default(void* self, void* signal);
void QAccessibilityHints_connect_contrastPreferenceChanged(void* self, long long id);
void QAccessibilityHints_connect_destroyed(void* self, long long id);
void QAccessibilityHints_connect_destroyed2(void* self, long long id);
void QAccessibilityHints_connect_objectNameChanged(void* self, long long id);
void QAccessibilityHints__disconnect(long long id);
void* QAccessibilityHints__Go_new(long long id, void* parent);
void QAccessibilityHints_delete(void* self);

#ifdef __cplusplus
}
#endif
