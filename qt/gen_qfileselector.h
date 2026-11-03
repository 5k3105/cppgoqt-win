#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QFileSelector_base(void* self);
void QFileSelector__set_override(void* self, int idx);
void* QFileSelector_fromQObject(void* o);
void* QFileSelector_new(long long __goid);
void* QFileSelector_new2(long long __goid, void* parent);
char* QFileSelector_tr(const char* s, long long* __out_len);
char* QFileSelector_select(void* self, const char* filePath, long long filePath_len, long long* __out_len);
void* QFileSelector_select2(void* self, void* filePath);
char** QFileSelector_extraSelectors(void* self, long long* __out_len);
void QFileSelector_setExtraSelectors(void* self, char** list, long long list_len);
char** QFileSelector_allSelectors(void* self, long long* __out_len);
void* QFileSelector_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QFileSelector_connect2(void* sender, void* signal, void* receiver, void* method);
void* QFileSelector_sender(void* self);
int QFileSelector_senderSignalIndex(void* self);
int QFileSelector_receivers(void* self, const char* signal);
int QFileSelector_isSignalConnected(void* self, void* signal);
char* QFileSelector_tr2(const char* s, const char* c, long long* __out_len);
char* QFileSelector_tr3(const char* s, const char* c, int n, long long* __out_len);
void* QFileSelector_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QFileSelector_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void* QFileSelector_metaObject_default(void* self);
int QFileSelector_event_default(void* self, void* event);
int QFileSelector_eventFilter_default(void* self, void* watched, void* event);
void QFileSelector_timerEvent_default(void* self, void* event);
void QFileSelector_childEvent_default(void* self, void* event);
void QFileSelector_customEvent_default(void* self, void* event);
void QFileSelector_connectNotify_default(void* self, void* signal);
void QFileSelector_disconnectNotify_default(void* self, void* signal);
void QFileSelector_connect_destroyed(void* self, long long id);
void QFileSelector_connect_destroyed2(void* self, long long id);
void QFileSelector_connect_objectNameChanged(void* self, long long id);
void QFileSelector__disconnect(long long id);
void* QFileSelector__Go_new(long long id, void* parent);
void QFileSelector_delete(void* self);

#ifdef __cplusplus
}
#endif
