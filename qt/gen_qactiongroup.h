#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QActionGroup_base(void* self);
void QActionGroup__set_override(void* self, int idx);
void* QActionGroup_fromQObject(void* o);
void* QActionGroup_new(long long __goid, void* parent);
char* QActionGroup_tr(const char* s, long long* __out_len);
void* QActionGroup_addAction(void* self, void* a);
void* QActionGroup_addAction2(void* self, const char* text, long long text_len);
void* QActionGroup_addAction3(void* self, void* icon, const char* text, long long text_len);
void QActionGroup_removeAction(void* self, void* a);
void** QActionGroup_actions(void* self, long long* __out_len);
void* QActionGroup_checkedAction(void* self);
int QActionGroup_isExclusive(void* self);
int QActionGroup_isEnabled(void* self);
int QActionGroup_isVisible(void* self);
int QActionGroup_exclusionPolicy(void* self);
void QActionGroup_setEnabled(void* self, int arg1);
void QActionGroup_setDisabled(void* self, int b);
void QActionGroup_setVisible(void* self, int arg1);
void QActionGroup_setExclusive(void* self, int arg1);
void QActionGroup_setExclusionPolicy(void* self, int policy);
void QActionGroup_triggered(void* self, void* arg1);
void QActionGroup_hovered(void* self, void* arg1);
void* QActionGroup_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QActionGroup_connect2(void* sender, void* signal, void* receiver, void* method);
void* QActionGroup_sender(void* self);
int QActionGroup_senderSignalIndex(void* self);
int QActionGroup_receivers(void* self, const char* signal);
int QActionGroup_isSignalConnected(void* self, void* signal);
char* QActionGroup_tr2(const char* s, const char* c, long long* __out_len);
char* QActionGroup_tr3(const char* s, const char* c, int n, long long* __out_len);
void* QActionGroup_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QActionGroup_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void* QActionGroup_metaObject_default(void* self);
int QActionGroup_event_default(void* self, void* event);
int QActionGroup_eventFilter_default(void* self, void* watched, void* event);
void QActionGroup_timerEvent_default(void* self, void* event);
void QActionGroup_childEvent_default(void* self, void* event);
void QActionGroup_customEvent_default(void* self, void* event);
void QActionGroup_connectNotify_default(void* self, void* signal);
void QActionGroup_disconnectNotify_default(void* self, void* signal);
void QActionGroup_connect_triggered(void* self, long long id);
void QActionGroup_connect_hovered(void* self, long long id);
void QActionGroup_connect_destroyed(void* self, long long id);
void QActionGroup_connect_destroyed2(void* self, long long id);
void QActionGroup_connect_objectNameChanged(void* self, long long id);
void QActionGroup__disconnect(long long id);
void* QActionGroup__Go_new(long long id, void* parent);
void QActionGroup_delete(void* self);

#ifdef __cplusplus
}
#endif
