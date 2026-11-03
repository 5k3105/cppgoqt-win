#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QStyledItemDelegate_base(void* self);
void QStyledItemDelegate_chain(void* self, void** out);
void QStyledItemDelegate__set_override(void* self, int idx);
void* QStyledItemDelegate_fromQObject(void* o);
void* QStyledItemDelegate_new(long long __goid);
void* QStyledItemDelegate_new2(long long __goid, void* parent);
char* QStyledItemDelegate_tr(const char* s, long long* __out_len);
void* QStyledItemDelegate_itemEditorFactory(void* self);
void QStyledItemDelegate_setItemEditorFactory(void* self, void* factory);
char* QStyledItemDelegate_displayText(void* self, void* value, void* locale, long long* __out_len);
void* QStyledItemDelegate_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QStyledItemDelegate_connect2(void* sender, void* signal, void* receiver, void* method);
void* QStyledItemDelegate_sender(void* self);
int QStyledItemDelegate_senderSignalIndex(void* self);
int QStyledItemDelegate_receivers(void* self, const char* signal);
int QStyledItemDelegate_isSignalConnected(void* self, void* signal);
char* QStyledItemDelegate_tr2(const char* s, const char* c, long long* __out_len);
char* QStyledItemDelegate_tr3(const char* s, const char* c, int n, long long* __out_len);
void* QStyledItemDelegate_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QStyledItemDelegate_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
char* QStyledItemDelegate_displayText_default(void* self, void* value, void* locale, long long* __out_len);
void* QStyledItemDelegate_metaObject_default(void* self);
void QStyledItemDelegate_paint_default(void* self, void* painter, void* option, void* index);
void* QStyledItemDelegate_sizeHint_default(void* self, void* option, void* index);
void* QStyledItemDelegate_createEditor_default(void* self, void* parent, void* option, void* index);
void QStyledItemDelegate_setEditorData_default(void* self, void* editor, void* index);
void QStyledItemDelegate_setModelData_default(void* self, void* editor, void* model, void* index);
void QStyledItemDelegate_updateEditorGeometry_default(void* self, void* editor, void* option, void* index);
void QStyledItemDelegate_initStyleOption_default(void* self, void* option, void* index);
int QStyledItemDelegate_eventFilter_default(void* self, void* object, void* event);
int QStyledItemDelegate_editorEvent_default(void* self, void* event, void* model, void* option, void* index);
void QStyledItemDelegate_destroyEditor_default(void* self, void* editor, void* index);
int QStyledItemDelegate_helpEvent_default(void* self, void* event, void* view, void* option, void* index);
int* QStyledItemDelegate_paintingRoles_default(void* self, long long* __out_len);
int QStyledItemDelegate_event_default(void* self, void* event);
void QStyledItemDelegate_timerEvent_default(void* self, void* event);
void QStyledItemDelegate_childEvent_default(void* self, void* event);
void QStyledItemDelegate_customEvent_default(void* self, void* event);
void QStyledItemDelegate_connectNotify_default(void* self, void* signal);
void QStyledItemDelegate_disconnectNotify_default(void* self, void* signal);
void QStyledItemDelegate_connect_commitData(void* self, long long id);
void QStyledItemDelegate_connect_closeEditor(void* self, long long id);
void QStyledItemDelegate_connect_closeEditor2(void* self, long long id);
void QStyledItemDelegate_connect_sizeHintChanged(void* self, long long id);
void QStyledItemDelegate_connect_destroyed(void* self, long long id);
void QStyledItemDelegate_connect_destroyed2(void* self, long long id);
void QStyledItemDelegate_connect_objectNameChanged(void* self, long long id);
void QStyledItemDelegate__disconnect(long long id);
void* QStyledItemDelegate__Go_new(long long id, void* parent);
void QStyledItemDelegate_delete(void* self);

#ifdef __cplusplus
}
#endif
