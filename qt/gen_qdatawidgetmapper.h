#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QDataWidgetMapper_base(void* self);
void QDataWidgetMapper__set_override(void* self, int idx);
void* QDataWidgetMapper_fromQObject(void* o);
void* QDataWidgetMapper_new(long long __goid);
void* QDataWidgetMapper_new2(long long __goid, void* parent);
char* QDataWidgetMapper_tr(const char* s, long long* __out_len);
void QDataWidgetMapper_setModel(void* self, void* model);
void* QDataWidgetMapper_model(void* self);
void QDataWidgetMapper_setItemDelegate(void* self, void* delegate);
void* QDataWidgetMapper_itemDelegate(void* self);
void QDataWidgetMapper_setRootIndex(void* self, void* index);
void* QDataWidgetMapper_rootIndex(void* self);
void QDataWidgetMapper_setOrientation(void* self, int aOrientation);
int QDataWidgetMapper_orientation(void* self);
void QDataWidgetMapper_setSubmitPolicy(void* self, int policy);
int QDataWidgetMapper_submitPolicy(void* self);
void QDataWidgetMapper_addMapping(void* self, void* widget, int section);
void QDataWidgetMapper_addMapping2(void* self, void* widget, int section, void* propertyName);
void QDataWidgetMapper_removeMapping(void* self, void* widget);
int QDataWidgetMapper_mappedSection(void* self, void* widget);
void* QDataWidgetMapper_mappedPropertyName(void* self, void* widget);
void* QDataWidgetMapper_mappedWidgetAt(void* self, int section);
void QDataWidgetMapper_clearMapping(void* self);
int QDataWidgetMapper_currentIndex(void* self);
void QDataWidgetMapper_revert(void* self);
int QDataWidgetMapper_submit(void* self);
void QDataWidgetMapper_toFirst(void* self);
void QDataWidgetMapper_toLast(void* self);
void QDataWidgetMapper_toNext(void* self);
void QDataWidgetMapper_toPrevious(void* self);
void QDataWidgetMapper_setCurrentIndex(void* self, int index);
void QDataWidgetMapper_setCurrentModelIndex(void* self, void* index);
void QDataWidgetMapper_currentIndexChanged(void* self, int index);
void* QDataWidgetMapper_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QDataWidgetMapper_connect2(void* sender, void* signal, void* receiver, void* method);
void* QDataWidgetMapper_sender(void* self);
int QDataWidgetMapper_senderSignalIndex(void* self);
int QDataWidgetMapper_receivers(void* self, const char* signal);
int QDataWidgetMapper_isSignalConnected(void* self, void* signal);
char* QDataWidgetMapper_tr2(const char* s, const char* c, long long* __out_len);
char* QDataWidgetMapper_tr3(const char* s, const char* c, int n, long long* __out_len);
void* QDataWidgetMapper_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QDataWidgetMapper_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void QDataWidgetMapper_setCurrentIndex_default(void* self, int index);
void* QDataWidgetMapper_metaObject_default(void* self);
int QDataWidgetMapper_event_default(void* self, void* event);
int QDataWidgetMapper_eventFilter_default(void* self, void* watched, void* event);
void QDataWidgetMapper_timerEvent_default(void* self, void* event);
void QDataWidgetMapper_childEvent_default(void* self, void* event);
void QDataWidgetMapper_customEvent_default(void* self, void* event);
void QDataWidgetMapper_connectNotify_default(void* self, void* signal);
void QDataWidgetMapper_disconnectNotify_default(void* self, void* signal);
void QDataWidgetMapper_connect_currentIndexChanged(void* self, long long id);
void QDataWidgetMapper_connect_destroyed(void* self, long long id);
void QDataWidgetMapper_connect_destroyed2(void* self, long long id);
void QDataWidgetMapper_connect_objectNameChanged(void* self, long long id);
void QDataWidgetMapper__disconnect(long long id);
void* QDataWidgetMapper__Go_new(long long id, void* parent);
void QDataWidgetMapper_delete(void* self);

#ifdef __cplusplus
}
#endif
