#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QSystemTrayIcon_base(void* self);
void QSystemTrayIcon__set_override(void* self, int idx);
void* QSystemTrayIcon_fromQObject(void* o);
void* QSystemTrayIcon_new(long long __goid);
void* QSystemTrayIcon_new2(long long __goid, void* icon);
void* QSystemTrayIcon_new3(long long __goid, void* parent);
void* QSystemTrayIcon_new4(long long __goid, void* icon, void* parent);
char* QSystemTrayIcon_tr(const char* s, long long* __out_len);
void QSystemTrayIcon_setContextMenu(void* self, void* menu);
void* QSystemTrayIcon_contextMenu(void* self);
void* QSystemTrayIcon_icon(void* self);
void QSystemTrayIcon_setIcon(void* self, void* icon);
char* QSystemTrayIcon_toolTip(void* self, long long* __out_len);
void QSystemTrayIcon_setToolTip(void* self, const char* tip, long long tip_len);
int QSystemTrayIcon_isSystemTrayAvailable();
int QSystemTrayIcon_supportsMessages();
void* QSystemTrayIcon_geometry(void* self);
int QSystemTrayIcon_isVisible(void* self);
void QSystemTrayIcon_setVisible(void* self, int visible);
void QSystemTrayIcon_show(void* self);
void QSystemTrayIcon_hide(void* self);
void QSystemTrayIcon_showMessage(void* self, const char* title, long long title_len, const char* msg, long long msg_len, void* icon);
void QSystemTrayIcon_showMessage2(void* self, const char* title, long long title_len, const char* msg, long long msg_len);
void QSystemTrayIcon_activated(void* self, int reason);
void QSystemTrayIcon_messageClicked(void* self);
void* QSystemTrayIcon_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QSystemTrayIcon_connect2(void* sender, void* signal, void* receiver, void* method);
void* QSystemTrayIcon_sender(void* self);
int QSystemTrayIcon_senderSignalIndex(void* self);
int QSystemTrayIcon_receivers(void* self, const char* signal);
int QSystemTrayIcon_isSignalConnected(void* self, void* signal);
char* QSystemTrayIcon_tr2(const char* s, const char* c, long long* __out_len);
char* QSystemTrayIcon_tr3(const char* s, const char* c, int n, long long* __out_len);
void QSystemTrayIcon_showMessage3(void* self, const char* title, long long title_len, const char* msg, long long msg_len, void* icon, int msecs);
void QSystemTrayIcon_showMessage4(void* self, const char* title, long long title_len, const char* msg, long long msg_len, int icon);
void QSystemTrayIcon_showMessage5(void* self, const char* title, long long title_len, const char* msg, long long msg_len, int icon, int msecs);
void* QSystemTrayIcon_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QSystemTrayIcon_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void* QSystemTrayIcon_metaObject_default(void* self);
int QSystemTrayIcon_event_default(void* self, void* event);
int QSystemTrayIcon_eventFilter_default(void* self, void* watched, void* event);
void QSystemTrayIcon_timerEvent_default(void* self, void* event);
void QSystemTrayIcon_childEvent_default(void* self, void* event);
void QSystemTrayIcon_customEvent_default(void* self, void* event);
void QSystemTrayIcon_connectNotify_default(void* self, void* signal);
void QSystemTrayIcon_disconnectNotify_default(void* self, void* signal);
void QSystemTrayIcon_connect_activated(void* self, long long id);
void QSystemTrayIcon_connect_messageClicked(void* self, long long id);
void QSystemTrayIcon_connect_destroyed(void* self, long long id);
void QSystemTrayIcon_connect_destroyed2(void* self, long long id);
void QSystemTrayIcon_connect_objectNameChanged(void* self, long long id);
void QSystemTrayIcon__disconnect(long long id);
void* QSystemTrayIcon__Go_new(long long id, void* parent);
void QSystemTrayIcon_delete(void* self);

#ifdef __cplusplus
}
#endif
