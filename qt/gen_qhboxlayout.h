#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QHBoxLayout_base(void* self);
void QHBoxLayout_chain(void* self, void** out);
void QHBoxLayout__set_override(void* self, int idx);
void* QHBoxLayout_fromQObject(void* o);
void* QHBoxLayout_new(long long __goid, void* parent);
void* QHBoxLayout_new2(long long __goid);
char* QHBoxLayout_tr(const char* s, long long* __out_len);
void* QHBoxLayout_closestAcceptableSize(void* w, void* s);
void QHBoxLayout_widgetEvent(void* self, void* arg1);
void QHBoxLayout_addChildLayout(void* self, void* l);
void QHBoxLayout_addChildWidget(void* self, void* w);
int QHBoxLayout_adoptLayout(void* self, void* layout);
void* QHBoxLayout_alignmentRect(void* self, void* arg1);
void* QHBoxLayout_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QHBoxLayout_connect2(void* sender, void* signal, void* receiver, void* method);
void* QHBoxLayout_sender(void* self);
int QHBoxLayout_senderSignalIndex(void* self);
int QHBoxLayout_receivers(void* self, const char* signal);
int QHBoxLayout_isSignalConnected(void* self, void* signal);
char* QHBoxLayout_tr2(const char* s, const char* c, long long* __out_len);
char* QHBoxLayout_tr3(const char* s, const char* c, int n, long long* __out_len);
void* QHBoxLayout_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QHBoxLayout_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void* QHBoxLayout_metaObject_default(void* self);
void QHBoxLayout_addItem_default(void* self, void* arg1);
int QHBoxLayout_spacing_default(void* self);
void QHBoxLayout_setSpacing_default(void* self, int spacing);
void* QHBoxLayout_sizeHint_default(void* self);
void* QHBoxLayout_minimumSize_default(void* self);
void* QHBoxLayout_maximumSize_default(void* self);
int QHBoxLayout_hasHeightForWidth_default(void* self);
int QHBoxLayout_heightForWidth_default(void* self, int arg1);
int QHBoxLayout_minimumHeightForWidth_default(void* self, int arg1);
int QHBoxLayout_expandingDirections_default(void* self);
void QHBoxLayout_invalidate_default(void* self);
void* QHBoxLayout_itemAt_default(void* self, int arg1);
void* QHBoxLayout_takeAt_default(void* self, int arg1);
int QHBoxLayout_count_default(void* self);
void QHBoxLayout_setGeometry_default(void* self, void* arg1);
void* QHBoxLayout_geometry_default(void* self);
int QHBoxLayout_indexOf_default(void* self, void* arg1);
int QHBoxLayout_isEmpty_default(void* self);
int QHBoxLayout_controlTypes_default(void* self);
void* QHBoxLayout_replaceWidget_default(void* self, void* from, void* to, int options);
void* QHBoxLayout_layout_default(void* self);
void QHBoxLayout_childEvent_default(void* self, void* e);
int QHBoxLayout_event_default(void* self, void* event);
int QHBoxLayout_eventFilter_default(void* self, void* watched, void* event);
void QHBoxLayout_timerEvent_default(void* self, void* event);
void QHBoxLayout_customEvent_default(void* self, void* event);
void QHBoxLayout_connectNotify_default(void* self, void* signal);
void QHBoxLayout_disconnectNotify_default(void* self, void* signal);
void* QHBoxLayout_widget_default(void* self);
void* QHBoxLayout_spacerItem_default(void* self);
void QHBoxLayout_connect_destroyed(void* self, long long id);
void QHBoxLayout_connect_destroyed2(void* self, long long id);
void QHBoxLayout_connect_objectNameChanged(void* self, long long id);
void QHBoxLayout__disconnect(long long id);
void* QHBoxLayout__Go_new(long long id);
void QHBoxLayout_delete(void* self);

#ifdef __cplusplus
}
#endif
