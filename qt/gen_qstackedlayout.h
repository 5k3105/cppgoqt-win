#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QStackedLayout_base(void* self);
void QStackedLayout_chain(void* self, void** out);
void QStackedLayout__set_override(void* self, int idx);
void* QStackedLayout_fromQObject(void* o);
void* QStackedLayout_new(long long __goid, void* parent);
void* QStackedLayout_new2(long long __goid);
void* QStackedLayout_new3(long long __goid, void* parentLayout);
char* QStackedLayout_tr(const char* s, long long* __out_len);
int QStackedLayout_addWidget(void* self, void* w);
int QStackedLayout_insertWidget(void* self, int index, void* w);
void* QStackedLayout_currentWidget(void* self);
int QStackedLayout_currentIndex(void* self);
void* QStackedLayout_widget(void* self, int arg1);
int QStackedLayout_stackingMode(void* self);
void QStackedLayout_setStackingMode(void* self, int stackingMode);
void QStackedLayout_widgetRemoved(void* self, int index);
void QStackedLayout_currentChanged(void* self, int index);
void QStackedLayout_widgetAdded(void* self, int index);
void QStackedLayout_setCurrentIndex(void* self, int index);
void QStackedLayout_setCurrentWidget(void* self, void* w);
void* QStackedLayout_closestAcceptableSize(void* w, void* s);
void QStackedLayout_widgetEvent(void* self, void* arg1);
void QStackedLayout_addChildLayout(void* self, void* l);
void QStackedLayout_addChildWidget(void* self, void* w);
int QStackedLayout_adoptLayout(void* self, void* layout);
void* QStackedLayout_alignmentRect(void* self, void* arg1);
void* QStackedLayout_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QStackedLayout_connect2(void* sender, void* signal, void* receiver, void* method);
void* QStackedLayout_sender(void* self);
int QStackedLayout_senderSignalIndex(void* self);
int QStackedLayout_receivers(void* self, const char* signal);
int QStackedLayout_isSignalConnected(void* self, void* signal);
char* QStackedLayout_tr2(const char* s, const char* c, long long* __out_len);
char* QStackedLayout_tr3(const char* s, const char* c, int n, long long* __out_len);
void* QStackedLayout_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QStackedLayout_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void* QStackedLayout_metaObject_default(void* self);
int QStackedLayout_count_default(void* self);
void QStackedLayout_addItem_default(void* self, void* item);
void* QStackedLayout_sizeHint_default(void* self);
void* QStackedLayout_minimumSize_default(void* self);
void* QStackedLayout_itemAt_default(void* self, int arg1);
void* QStackedLayout_takeAt_default(void* self, int arg1);
void QStackedLayout_setGeometry_default(void* self, void* rect);
int QStackedLayout_hasHeightForWidth_default(void* self);
int QStackedLayout_heightForWidth_default(void* self, int width);
int QStackedLayout_spacing_default(void* self);
void QStackedLayout_setSpacing_default(void* self, int arg1);
void QStackedLayout_invalidate_default(void* self);
void* QStackedLayout_geometry_default(void* self);
int QStackedLayout_expandingDirections_default(void* self);
void* QStackedLayout_maximumSize_default(void* self);
int QStackedLayout_indexOf_default(void* self, void* arg1);
int QStackedLayout_isEmpty_default(void* self);
int QStackedLayout_controlTypes_default(void* self);
void* QStackedLayout_replaceWidget_default(void* self, void* from, void* to, int options);
void* QStackedLayout_layout_default(void* self);
void QStackedLayout_childEvent_default(void* self, void* e);
int QStackedLayout_event_default(void* self, void* event);
int QStackedLayout_eventFilter_default(void* self, void* watched, void* event);
void QStackedLayout_timerEvent_default(void* self, void* event);
void QStackedLayout_customEvent_default(void* self, void* event);
void QStackedLayout_connectNotify_default(void* self, void* signal);
void QStackedLayout_disconnectNotify_default(void* self, void* signal);
int QStackedLayout_minimumHeightForWidth_default(void* self, int arg1);
void* QStackedLayout_spacerItem_default(void* self);
void QStackedLayout_connect_widgetRemoved(void* self, long long id);
void QStackedLayout_connect_currentChanged(void* self, long long id);
void QStackedLayout_connect_widgetAdded(void* self, long long id);
void QStackedLayout_connect_destroyed(void* self, long long id);
void QStackedLayout_connect_destroyed2(void* self, long long id);
void QStackedLayout_connect_objectNameChanged(void* self, long long id);
void QStackedLayout__disconnect(long long id);
void* QStackedLayout__Go_new(long long id);
void QStackedLayout_delete(void* self);

#ifdef __cplusplus
}
#endif
