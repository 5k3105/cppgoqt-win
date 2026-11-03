#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QSignalMapper_base(void* self);
void QSignalMapper__set_override(void* self, int idx);
void* QSignalMapper_fromQObject(void* o);
void* QSignalMapper_new(long long __goid);
void* QSignalMapper_new2(long long __goid, void* parent);
char* QSignalMapper_tr(const char* s, long long* __out_len);
void QSignalMapper_setMapping(void* self, void* sender, int id);
void QSignalMapper_setMapping2(void* self, void* sender, const char* text, long long text_len);
void QSignalMapper_setMapping3(void* self, void* sender, void* object);
void QSignalMapper_removeMappings(void* self, void* sender);
void* QSignalMapper_mapping(void* self, int id);
void* QSignalMapper_mapping2(void* self, const char* text, long long text_len);
void* QSignalMapper_mapping3(void* self, void* object);
void QSignalMapper_mappedInt(void* self, int arg1);
void QSignalMapper_mappedString(void* self, const char* arg1, long long arg1_len);
void QSignalMapper_mappedObject(void* self, void* arg1);
void QSignalMapper_map(void* self);
void QSignalMapper_map2(void* self, void* sender);
void* QSignalMapper_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QSignalMapper_connect2(void* sender, void* signal, void* receiver, void* method);
void* QSignalMapper_sender(void* self);
int QSignalMapper_senderSignalIndex(void* self);
int QSignalMapper_receivers(void* self, const char* signal);
int QSignalMapper_isSignalConnected(void* self, void* signal);
char* QSignalMapper_tr2(const char* s, const char* c, long long* __out_len);
char* QSignalMapper_tr3(const char* s, const char* c, int n, long long* __out_len);
void* QSignalMapper_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QSignalMapper_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void* QSignalMapper_metaObject_default(void* self);
int QSignalMapper_event_default(void* self, void* event);
int QSignalMapper_eventFilter_default(void* self, void* watched, void* event);
void QSignalMapper_timerEvent_default(void* self, void* event);
void QSignalMapper_childEvent_default(void* self, void* event);
void QSignalMapper_customEvent_default(void* self, void* event);
void QSignalMapper_connectNotify_default(void* self, void* signal);
void QSignalMapper_disconnectNotify_default(void* self, void* signal);
void QSignalMapper_connect_mappedInt(void* self, long long id);
void QSignalMapper_connect_mappedString(void* self, long long id);
void QSignalMapper_connect_mappedObject(void* self, long long id);
void QSignalMapper_connect_destroyed(void* self, long long id);
void QSignalMapper_connect_destroyed2(void* self, long long id);
void QSignalMapper_connect_objectNameChanged(void* self, long long id);
void QSignalMapper__disconnect(long long id);
void* QSignalMapper__Go_new(long long id, void* parent);
void QSignalMapper_delete(void* self);

#ifdef __cplusplus
}
#endif
