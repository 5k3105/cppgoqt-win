#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QIntValidator_base(void* self);
void QIntValidator_chain(void* self, void** out);
void QIntValidator__set_override(void* self, int idx);
void* QIntValidator_fromQObject(void* o);
void* QIntValidator_new(long long __goid);
void* QIntValidator_new2(long long __goid, int bottom, int top);
void* QIntValidator_new3(long long __goid, void* parent);
void* QIntValidator_new4(long long __goid, int bottom, int top, void* parent);
char* QIntValidator_tr(const char* s, long long* __out_len);
void QIntValidator_setBottom(void* self, int arg1);
void QIntValidator_setTop(void* self, int arg1);
void QIntValidator_setRange(void* self, int bottom, int top);
int QIntValidator_bottom(void* self);
int QIntValidator_top(void* self);
void QIntValidator_bottomChanged(void* self, int bottom);
void QIntValidator_topChanged(void* self, int top);
void* QIntValidator_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QIntValidator_connect2(void* sender, void* signal, void* receiver, void* method);
void* QIntValidator_sender(void* self);
int QIntValidator_senderSignalIndex(void* self);
int QIntValidator_receivers(void* self, const char* signal);
int QIntValidator_isSignalConnected(void* self, void* signal);
char* QIntValidator_tr2(const char* s, const char* c, long long* __out_len);
char* QIntValidator_tr3(const char* s, const char* c, int n, long long* __out_len);
void* QIntValidator_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QIntValidator_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void* QIntValidator_metaObject_default(void* self);
int QIntValidator_validate_default(void* self, void* arg1, int* arg2);
void QIntValidator_fixup_default(void* self, void* input);
int QIntValidator_event_default(void* self, void* event);
int QIntValidator_eventFilter_default(void* self, void* watched, void* event);
void QIntValidator_timerEvent_default(void* self, void* event);
void QIntValidator_childEvent_default(void* self, void* event);
void QIntValidator_customEvent_default(void* self, void* event);
void QIntValidator_connectNotify_default(void* self, void* signal);
void QIntValidator_disconnectNotify_default(void* self, void* signal);
void QIntValidator_connect_bottomChanged(void* self, long long id);
void QIntValidator_connect_topChanged(void* self, long long id);
void QIntValidator_connect_changed(void* self, long long id);
void QIntValidator_connect_destroyed(void* self, long long id);
void QIntValidator_connect_destroyed2(void* self, long long id);
void QIntValidator_connect_objectNameChanged(void* self, long long id);
void QIntValidator__disconnect(long long id);
void* QIntValidator__Go_new(long long id, void* parent);
void QIntValidator_delete(void* self);

#ifdef __cplusplus
}
#endif
