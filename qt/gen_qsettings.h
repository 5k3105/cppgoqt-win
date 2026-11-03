#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QSettings_base(void* self);
void QSettings__set_override(void* self, int idx);
void* QSettings_fromQObject(void* o);
void* QSettings_new(long long __goid, const char* organization, long long organization_len);
void* QSettings_new2(long long __goid, int scope, const char* organization, long long organization_len);
void* QSettings_new3(long long __goid, int format, int scope, const char* organization, long long organization_len);
void* QSettings_new4(long long __goid, const char* fileName, long long fileName_len, int format);
void* QSettings_new5(long long __goid);
void* QSettings_new6(long long __goid, int scope);
void* QSettings_new7(long long __goid, const char* organization, long long organization_len, const char* application, long long application_len);
void* QSettings_new8(long long __goid, const char* organization, long long organization_len, const char* application, long long application_len, void* parent);
void* QSettings_new9(long long __goid, int scope, const char* organization, long long organization_len, const char* application, long long application_len);
void* QSettings_new10(long long __goid, int scope, const char* organization, long long organization_len, const char* application, long long application_len, void* parent);
void* QSettings_new11(long long __goid, int format, int scope, const char* organization, long long organization_len, const char* application, long long application_len);
void* QSettings_new12(long long __goid, int format, int scope, const char* organization, long long organization_len, const char* application, long long application_len, void* parent);
void* QSettings_new13(long long __goid, const char* fileName, long long fileName_len, int format, void* parent);
void* QSettings_new14(long long __goid, void* parent);
void* QSettings_new15(long long __goid, int scope, void* parent);
char* QSettings_tr(const char* s, long long* __out_len);
void QSettings_clear(void* self);
void QSettings_sync(void* self);
int QSettings_status(void* self);
int QSettings_isAtomicSyncRequired(void* self);
void QSettings_setAtomicSyncRequired(void* self, int enable);
void QSettings_beginGroup(void* self, const char* prefix, long long prefix_len);
void QSettings_endGroup(void* self);
char* QSettings_group(void* self, long long* __out_len);
int QSettings_beginReadArray(void* self, const char* prefix, long long prefix_len);
void QSettings_beginWriteArray(void* self, const char* prefix, long long prefix_len);
void QSettings_endArray(void* self);
void QSettings_setArrayIndex(void* self, int i);
char** QSettings_allKeys(void* self, long long* __out_len);
char** QSettings_childKeys(void* self, long long* __out_len);
char** QSettings_childGroups(void* self, long long* __out_len);
int QSettings_isWritable(void* self);
void QSettings_setValue(void* self, const char* key, long long key_len, void* value);
void* QSettings_value(void* self, const char* key, long long key_len, void* defaultValue);
void* QSettings_value2(void* self, const char* key, long long key_len);
void QSettings_remove(void* self, const char* key, long long key_len);
int QSettings_contains(void* self, const char* key, long long key_len);
void QSettings_setFallbacksEnabled(void* self, int b);
int QSettings_fallbacksEnabled(void* self);
char* QSettings_fileName(void* self, long long* __out_len);
int QSettings_format(void* self);
int QSettings_scope(void* self);
char* QSettings_organizationName(void* self, long long* __out_len);
char* QSettings_applicationName(void* self, long long* __out_len);
void QSettings_setDefaultFormat(int format);
int QSettings_defaultFormat();
void QSettings_setPath(int format, int scope, const char* path, long long path_len);
void* QSettings_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QSettings_connect2(void* sender, void* signal, void* receiver, void* method);
void* QSettings_sender(void* self);
int QSettings_senderSignalIndex(void* self);
int QSettings_receivers(void* self, const char* signal);
int QSettings_isSignalConnected(void* self, void* signal);
char* QSettings_tr2(const char* s, const char* c, long long* __out_len);
char* QSettings_tr3(const char* s, const char* c, int n, long long* __out_len);
void QSettings_beginWriteArray2(void* self, const char* prefix, long long prefix_len, int size);
void* QSettings_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QSettings_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void* QSettings_metaObject_default(void* self);
int QSettings_event_default(void* self, void* event);
int QSettings_eventFilter_default(void* self, void* watched, void* event);
void QSettings_timerEvent_default(void* self, void* event);
void QSettings_childEvent_default(void* self, void* event);
void QSettings_customEvent_default(void* self, void* event);
void QSettings_connectNotify_default(void* self, void* signal);
void QSettings_disconnectNotify_default(void* self, void* signal);
void QSettings_connect_destroyed(void* self, long long id);
void QSettings_connect_destroyed2(void* self, long long id);
void QSettings_connect_objectNameChanged(void* self, long long id);
void QSettings__disconnect(long long id);
void* QSettings__Go_new(long long id, void* parent);
void QSettings_delete(void* self);

#ifdef __cplusplus
}
#endif
