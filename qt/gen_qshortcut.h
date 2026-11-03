#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QShortcut_base(void* self);
void QShortcut__set_override(void* self, int idx);
void* QShortcut_fromQObject(void* o);
void* QShortcut_new(long long __goid, void* parent);
void* QShortcut_new2(long long __goid, void* key, void* parent);
void* QShortcut_new3(long long __goid, int key, void* parent);
void* QShortcut_new4(long long __goid, void* key, void* parent, const char* member);
void* QShortcut_new5(long long __goid, void* key, void* parent, const char* member, const char* ambiguousMember);
void* QShortcut_new6(long long __goid, void* key, void* parent, const char* member, const char* ambiguousMember, int context);
void* QShortcut_new7(long long __goid, int key, void* parent, const char* member);
void* QShortcut_new8(long long __goid, int key, void* parent, const char* member, const char* ambiguousMember);
void* QShortcut_new9(long long __goid, int key, void* parent, const char* member, const char* ambiguousMember, int context);
char* QShortcut_tr(const char* s, long long* __out_len);
void QShortcut_setKey(void* self, void* key);
void* QShortcut_key(void* self);
void QShortcut_setKeys(void* self, int key);
void QShortcut_setKeys2(void* self, void** keys, long long keys_len);
void** QShortcut_keys(void* self, long long* __out_len);
void QShortcut_setEnabled(void* self, int enable);
int QShortcut_isEnabled(void* self);
void QShortcut_setContext(void* self, int context);
int QShortcut_context(void* self);
void QShortcut_setAutoRepeat(void* self, int on);
int QShortcut_autoRepeat(void* self);
int QShortcut_id(void* self);
void QShortcut_setWhatsThis(void* self, const char* text, long long text_len);
char* QShortcut_whatsThis(void* self, long long* __out_len);
void QShortcut_activated(void* self);
void QShortcut_activatedAmbiguously(void* self);
void* QShortcut_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QShortcut_connect2(void* sender, void* signal, void* receiver, void* method);
void* QShortcut_sender(void* self);
int QShortcut_senderSignalIndex(void* self);
int QShortcut_receivers(void* self, const char* signal);
int QShortcut_isSignalConnected(void* self, void* signal);
char* QShortcut_tr2(const char* s, const char* c, long long* __out_len);
char* QShortcut_tr3(const char* s, const char* c, int n, long long* __out_len);
void* QShortcut_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QShortcut_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void* QShortcut_metaObject_default(void* self);
int QShortcut_event_default(void* self, void* e);
int QShortcut_eventFilter_default(void* self, void* watched, void* event);
void QShortcut_timerEvent_default(void* self, void* event);
void QShortcut_childEvent_default(void* self, void* event);
void QShortcut_customEvent_default(void* self, void* event);
void QShortcut_connectNotify_default(void* self, void* signal);
void QShortcut_disconnectNotify_default(void* self, void* signal);
void QShortcut_connect_activated(void* self, long long id);
void QShortcut_connect_activatedAmbiguously(void* self, long long id);
void QShortcut_connect_destroyed(void* self, long long id);
void QShortcut_connect_destroyed2(void* self, long long id);
void QShortcut_connect_objectNameChanged(void* self, long long id);
void QShortcut__disconnect(long long id);
void* QShortcut__Go_new(long long id, void* parent);
void QShortcut_delete(void* self);

#ifdef __cplusplus
}
#endif
