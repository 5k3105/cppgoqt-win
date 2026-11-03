#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QVBoxLayout_base(void* self);
void QVBoxLayout_chain(void* self, void** out);
void QVBoxLayout__set_override(void* self, int idx);
void* QVBoxLayout_fromQObject(void* o);
void* QVBoxLayout_new(long long __goid, void* parent);
void* QVBoxLayout_new2(long long __goid);
char* QVBoxLayout_tr(const char* s, long long* __out_len);
void* QVBoxLayout_closestAcceptableSize(void* w, void* s);
void QVBoxLayout_widgetEvent(void* self, void* arg1);
void QVBoxLayout_addChildLayout(void* self, void* l);
void QVBoxLayout_addChildWidget(void* self, void* w);
int QVBoxLayout_adoptLayout(void* self, void* layout);
void* QVBoxLayout_alignmentRect(void* self, void* arg1);
void* QVBoxLayout_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QVBoxLayout_connect2(void* sender, void* signal, void* receiver, void* method);
void* QVBoxLayout_sender(void* self);
int QVBoxLayout_senderSignalIndex(void* self);
int QVBoxLayout_receivers(void* self, const char* signal);
int QVBoxLayout_isSignalConnected(void* self, void* signal);
char* QVBoxLayout_tr2(const char* s, const char* c, long long* __out_len);
char* QVBoxLayout_tr3(const char* s, const char* c, int n, long long* __out_len);
void* QVBoxLayout_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QVBoxLayout_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void* QVBoxLayout_metaObject_default(void* self);
void QVBoxLayout_addItem_default(void* self, void* arg1);
int QVBoxLayout_spacing_default(void* self);
void QVBoxLayout_setSpacing_default(void* self, int spacing);
void* QVBoxLayout_sizeHint_default(void* self);
void* QVBoxLayout_minimumSize_default(void* self);
void* QVBoxLayout_maximumSize_default(void* self);
int QVBoxLayout_hasHeightForWidth_default(void* self);
int QVBoxLayout_heightForWidth_default(void* self, int arg1);
int QVBoxLayout_minimumHeightForWidth_default(void* self, int arg1);
int QVBoxLayout_expandingDirections_default(void* self);
void QVBoxLayout_invalidate_default(void* self);
void* QVBoxLayout_itemAt_default(void* self, int arg1);
void* QVBoxLayout_takeAt_default(void* self, int arg1);
int QVBoxLayout_count_default(void* self);
void QVBoxLayout_setGeometry_default(void* self, void* arg1);
void* QVBoxLayout_geometry_default(void* self);
int QVBoxLayout_indexOf_default(void* self, void* arg1);
int QVBoxLayout_isEmpty_default(void* self);
int QVBoxLayout_controlTypes_default(void* self);
void* QVBoxLayout_replaceWidget_default(void* self, void* from, void* to, int options);
void* QVBoxLayout_layout_default(void* self);
void QVBoxLayout_childEvent_default(void* self, void* e);
int QVBoxLayout_event_default(void* self, void* event);
int QVBoxLayout_eventFilter_default(void* self, void* watched, void* event);
void QVBoxLayout_timerEvent_default(void* self, void* event);
void QVBoxLayout_customEvent_default(void* self, void* event);
void QVBoxLayout_connectNotify_default(void* self, void* signal);
void QVBoxLayout_disconnectNotify_default(void* self, void* signal);
void* QVBoxLayout_widget_default(void* self);
void* QVBoxLayout_spacerItem_default(void* self);
void QVBoxLayout_connect_destroyed(void* self, long long id);
void QVBoxLayout_connect_destroyed2(void* self, long long id);
void QVBoxLayout_connect_objectNameChanged(void* self, long long id);
void QVBoxLayout__disconnect(long long id);
void* QVBoxLayout__Go_new(long long id);
void QVBoxLayout_delete(void* self);

#ifdef __cplusplus
}
#endif
