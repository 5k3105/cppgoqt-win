#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QDoubleValidator_base(void* self);
void QDoubleValidator_chain(void* self, void** out);
void QDoubleValidator__set_override(void* self, int idx);
void* QDoubleValidator_fromQObject(void* o);
void* QDoubleValidator_new(long long __goid);
void* QDoubleValidator_new2(long long __goid, double bottom, double top, int decimals);
void* QDoubleValidator_new3(long long __goid, void* parent);
void* QDoubleValidator_new4(long long __goid, double bottom, double top, int decimals, void* parent);
char* QDoubleValidator_tr(const char* s, long long* __out_len);
void QDoubleValidator_setRange(void* self, double bottom, double top, int decimals);
void QDoubleValidator_setRange2(void* self, double bottom, double top);
void QDoubleValidator_setBottom(void* self, double arg1);
void QDoubleValidator_setTop(void* self, double arg1);
void QDoubleValidator_setDecimals(void* self, int arg1);
void QDoubleValidator_setNotation(void* self, int arg1);
double QDoubleValidator_bottom(void* self);
double QDoubleValidator_top(void* self);
int QDoubleValidator_decimals(void* self);
int QDoubleValidator_notation(void* self);
void QDoubleValidator_bottomChanged(void* self, double bottom);
void QDoubleValidator_topChanged(void* self, double top);
void QDoubleValidator_decimalsChanged(void* self, int decimals);
void QDoubleValidator_notationChanged(void* self, int notation);
void* QDoubleValidator_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QDoubleValidator_connect2(void* sender, void* signal, void* receiver, void* method);
void* QDoubleValidator_sender(void* self);
int QDoubleValidator_senderSignalIndex(void* self);
int QDoubleValidator_receivers(void* self, const char* signal);
int QDoubleValidator_isSignalConnected(void* self, void* signal);
char* QDoubleValidator_tr2(const char* s, const char* c, long long* __out_len);
char* QDoubleValidator_tr3(const char* s, const char* c, int n, long long* __out_len);
void* QDoubleValidator_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QDoubleValidator_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void* QDoubleValidator_metaObject_default(void* self);
int QDoubleValidator_validate_default(void* self, void* arg1, int* arg2);
void QDoubleValidator_fixup_default(void* self, void* input);
int QDoubleValidator_event_default(void* self, void* event);
int QDoubleValidator_eventFilter_default(void* self, void* watched, void* event);
void QDoubleValidator_timerEvent_default(void* self, void* event);
void QDoubleValidator_childEvent_default(void* self, void* event);
void QDoubleValidator_customEvent_default(void* self, void* event);
void QDoubleValidator_connectNotify_default(void* self, void* signal);
void QDoubleValidator_disconnectNotify_default(void* self, void* signal);
void QDoubleValidator_connect_bottomChanged(void* self, long long id);
void QDoubleValidator_connect_topChanged(void* self, long long id);
void QDoubleValidator_connect_decimalsChanged(void* self, long long id);
void QDoubleValidator_connect_notationChanged(void* self, long long id);
void QDoubleValidator_connect_changed(void* self, long long id);
void QDoubleValidator_connect_destroyed(void* self, long long id);
void QDoubleValidator_connect_destroyed2(void* self, long long id);
void QDoubleValidator_connect_objectNameChanged(void* self, long long id);
void QDoubleValidator__disconnect(long long id);
void* QDoubleValidator__Go_new(long long id, void* parent);
void QDoubleValidator_delete(void* self);

#ifdef __cplusplus
}
#endif
