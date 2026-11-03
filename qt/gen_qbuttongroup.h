#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QButtonGroup_base(void* self);
void QButtonGroup__set_override(void* self, int idx);
void* QButtonGroup_fromQObject(void* o);
void* QButtonGroup_new(long long __goid);
void* QButtonGroup_new2(long long __goid, void* parent);
char* QButtonGroup_tr(const char* s, long long* __out_len);
void QButtonGroup_setExclusive(void* self, int arg1);
int QButtonGroup_exclusive(void* self);
void QButtonGroup_addButton(void* self, void* arg1);
void QButtonGroup_removeButton(void* self, void* arg1);
void** QButtonGroup_buttons(void* self, long long* __out_len);
void* QButtonGroup_checkedButton(void* self);
void* QButtonGroup_button(void* self, int id);
void QButtonGroup_setId(void* self, void* button, int id);
int QButtonGroup_id(void* self, void* button);
int QButtonGroup_checkedId(void* self);
void QButtonGroup_buttonClicked(void* self, void* arg1);
void QButtonGroup_buttonPressed(void* self, void* arg1);
void QButtonGroup_buttonReleased(void* self, void* arg1);
void QButtonGroup_buttonToggled(void* self, void* arg1, int arg2);
void QButtonGroup_idClicked(void* self, int arg1);
void QButtonGroup_idPressed(void* self, int arg1);
void QButtonGroup_idReleased(void* self, int arg1);
void QButtonGroup_idToggled(void* self, int arg1, int arg2);
void* QButtonGroup_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QButtonGroup_connect2(void* sender, void* signal, void* receiver, void* method);
void* QButtonGroup_sender(void* self);
int QButtonGroup_senderSignalIndex(void* self);
int QButtonGroup_receivers(void* self, const char* signal);
int QButtonGroup_isSignalConnected(void* self, void* signal);
char* QButtonGroup_tr2(const char* s, const char* c, long long* __out_len);
char* QButtonGroup_tr3(const char* s, const char* c, int n, long long* __out_len);
void QButtonGroup_addButton2(void* self, void* arg1, int id);
void* QButtonGroup_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QButtonGroup_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void* QButtonGroup_metaObject_default(void* self);
int QButtonGroup_event_default(void* self, void* event);
int QButtonGroup_eventFilter_default(void* self, void* watched, void* event);
void QButtonGroup_timerEvent_default(void* self, void* event);
void QButtonGroup_childEvent_default(void* self, void* event);
void QButtonGroup_customEvent_default(void* self, void* event);
void QButtonGroup_connectNotify_default(void* self, void* signal);
void QButtonGroup_disconnectNotify_default(void* self, void* signal);
void QButtonGroup_connect_buttonClicked(void* self, long long id);
void QButtonGroup_connect_buttonPressed(void* self, long long id);
void QButtonGroup_connect_buttonReleased(void* self, long long id);
void QButtonGroup_connect_buttonToggled(void* self, long long id);
void QButtonGroup_connect_idClicked(void* self, long long id);
void QButtonGroup_connect_idPressed(void* self, long long id);
void QButtonGroup_connect_idReleased(void* self, long long id);
void QButtonGroup_connect_idToggled(void* self, long long id);
void QButtonGroup_connect_destroyed(void* self, long long id);
void QButtonGroup_connect_destroyed2(void* self, long long id);
void QButtonGroup_connect_objectNameChanged(void* self, long long id);
void QButtonGroup__disconnect(long long id);
void* QButtonGroup__Go_new(long long id, void* parent);
void QButtonGroup_delete(void* self);

#ifdef __cplusplus
}
#endif
