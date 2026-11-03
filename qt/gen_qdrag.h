#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QDrag_base(void* self);
void QDrag__set_override(void* self, int idx);
void* QDrag_fromQObject(void* o);
void* QDrag_new(long long __goid, void* dragSource);
char* QDrag_tr(const char* s, long long* __out_len);
void QDrag_setMimeData(void* self, void* data);
void* QDrag_mimeData(void* self);
void QDrag_setPixmap(void* self, void* arg1);
void* QDrag_pixmap(void* self);
void QDrag_setHotSpot(void* self, void* hotspot);
void* QDrag_hotSpot(void* self);
void* QDrag_source(void* self);
void* QDrag_target(void* self);
int QDrag_exec(void* self);
int QDrag_exec2(void* self, int supportedActions, int defaultAction);
void QDrag_setDragCursor(void* self, void* cursor, int action);
void* QDrag_dragCursor(void* self, int action);
int QDrag_supportedActions(void* self);
int QDrag_defaultAction(void* self);
void QDrag_cancel();
void QDrag_actionChanged(void* self, int action);
void QDrag_targetChanged(void* self, void* newTarget);
void* QDrag_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QDrag_connect2(void* sender, void* signal, void* receiver, void* method);
void* QDrag_sender(void* self);
int QDrag_senderSignalIndex(void* self);
int QDrag_receivers(void* self, const char* signal);
int QDrag_isSignalConnected(void* self, void* signal);
char* QDrag_tr2(const char* s, const char* c, long long* __out_len);
char* QDrag_tr3(const char* s, const char* c, int n, long long* __out_len);
int QDrag_exec3(void* self, int supportedActions);
void* QDrag_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QDrag_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void* QDrag_metaObject_default(void* self);
int QDrag_event_default(void* self, void* event);
int QDrag_eventFilter_default(void* self, void* watched, void* event);
void QDrag_timerEvent_default(void* self, void* event);
void QDrag_childEvent_default(void* self, void* event);
void QDrag_customEvent_default(void* self, void* event);
void QDrag_connectNotify_default(void* self, void* signal);
void QDrag_disconnectNotify_default(void* self, void* signal);
void QDrag_connect_actionChanged(void* self, long long id);
void QDrag_connect_targetChanged(void* self, long long id);
void QDrag_connect_destroyed(void* self, long long id);
void QDrag_connect_destroyed2(void* self, long long id);
void QDrag_connect_objectNameChanged(void* self, long long id);
void QDrag__disconnect(long long id);
void* QDrag__Go_new(long long id, void* parent);
void QDrag_delete(void* self);

#ifdef __cplusplus
}
#endif
