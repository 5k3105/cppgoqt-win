#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QFileSystemWatcher_base(void* self);
void QFileSystemWatcher__set_override(void* self, int idx);
void* QFileSystemWatcher_fromQObject(void* o);
void* QFileSystemWatcher_new(long long __goid);
void* QFileSystemWatcher_new2(long long __goid, char** paths, long long paths_len);
void* QFileSystemWatcher_new3(long long __goid, void* parent);
void* QFileSystemWatcher_new4(long long __goid, char** paths, long long paths_len, void* parent);
char* QFileSystemWatcher_tr(const char* s, long long* __out_len);
int QFileSystemWatcher_addPath(void* self, const char* file, long long file_len);
char** QFileSystemWatcher_addPaths(void* self, char** files, long long files_len, long long* __out_len);
int QFileSystemWatcher_removePath(void* self, const char* file, long long file_len);
char** QFileSystemWatcher_removePaths(void* self, char** files, long long files_len, long long* __out_len);
char** QFileSystemWatcher_files(void* self, long long* __out_len);
char** QFileSystemWatcher_directories(void* self, long long* __out_len);
void* QFileSystemWatcher_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QFileSystemWatcher_connect2(void* sender, void* signal, void* receiver, void* method);
void* QFileSystemWatcher_sender(void* self);
int QFileSystemWatcher_senderSignalIndex(void* self);
int QFileSystemWatcher_receivers(void* self, const char* signal);
int QFileSystemWatcher_isSignalConnected(void* self, void* signal);
char* QFileSystemWatcher_tr2(const char* s, const char* c, long long* __out_len);
char* QFileSystemWatcher_tr3(const char* s, const char* c, int n, long long* __out_len);
void* QFileSystemWatcher_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QFileSystemWatcher_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void* QFileSystemWatcher_metaObject_default(void* self);
int QFileSystemWatcher_event_default(void* self, void* event);
int QFileSystemWatcher_eventFilter_default(void* self, void* watched, void* event);
void QFileSystemWatcher_timerEvent_default(void* self, void* event);
void QFileSystemWatcher_childEvent_default(void* self, void* event);
void QFileSystemWatcher_customEvent_default(void* self, void* event);
void QFileSystemWatcher_connectNotify_default(void* self, void* signal);
void QFileSystemWatcher_disconnectNotify_default(void* self, void* signal);
void QFileSystemWatcher_connect_fileChanged(void* self, long long id);
void QFileSystemWatcher_connect_directoryChanged(void* self, long long id);
void QFileSystemWatcher_connect_destroyed(void* self, long long id);
void QFileSystemWatcher_connect_destroyed2(void* self, long long id);
void QFileSystemWatcher_connect_objectNameChanged(void* self, long long id);
void QFileSystemWatcher__disconnect(long long id);
void* QFileSystemWatcher__Go_new(long long id, void* parent);
void QFileSystemWatcher_delete(void* self);

#ifdef __cplusplus
}
#endif
