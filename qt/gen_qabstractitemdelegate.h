#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QAbstractItemDelegate_base(void* self);
void QAbstractItemDelegate__set_override(void* self, int idx);
void* QAbstractItemDelegate_fromQObject(void* o);
char* QAbstractItemDelegate_tr(const char* s, long long* __out_len);
void QAbstractItemDelegate_paint(void* self, void* painter, void* option, void* index);
void* QAbstractItemDelegate_sizeHint(void* self, void* option, void* index);
void* QAbstractItemDelegate_createEditor(void* self, void* parent, void* option, void* index);
void QAbstractItemDelegate_destroyEditor(void* self, void* editor, void* index);
void QAbstractItemDelegate_setEditorData(void* self, void* editor, void* index);
void QAbstractItemDelegate_setModelData(void* self, void* editor, void* model, void* index);
void QAbstractItemDelegate_updateEditorGeometry(void* self, void* editor, void* option, void* index);
int QAbstractItemDelegate_editorEvent(void* self, void* event, void* model, void* option, void* index);
int QAbstractItemDelegate_handleEditorEvent(void* self, void* object, void* event);
int QAbstractItemDelegate_helpEvent(void* self, void* event, void* view, void* option, void* index);
int* QAbstractItemDelegate_paintingRoles(void* self, long long* __out_len);
void QAbstractItemDelegate_commitData(void* self, void* editor);
void QAbstractItemDelegate_closeEditor(void* self, void* editor);
void QAbstractItemDelegate_sizeHintChanged(void* self, void* arg1);
void* QAbstractItemDelegate_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QAbstractItemDelegate_connect2(void* sender, void* signal, void* receiver, void* method);
void* QAbstractItemDelegate_sender(void* self);
int QAbstractItemDelegate_senderSignalIndex(void* self);
int QAbstractItemDelegate_receivers(void* self, const char* signal);
int QAbstractItemDelegate_isSignalConnected(void* self, void* signal);
char* QAbstractItemDelegate_tr2(const char* s, const char* c, long long* __out_len);
char* QAbstractItemDelegate_tr3(const char* s, const char* c, int n, long long* __out_len);
void QAbstractItemDelegate_closeEditor2(void* self, void* editor, int hint);
void* QAbstractItemDelegate_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QAbstractItemDelegate_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void* QAbstractItemDelegate_createEditor_default(void* self, void* parent, void* option, void* index);
void QAbstractItemDelegate_destroyEditor_default(void* self, void* editor, void* index);
void QAbstractItemDelegate_setEditorData_default(void* self, void* editor, void* index);
void QAbstractItemDelegate_setModelData_default(void* self, void* editor, void* model, void* index);
void QAbstractItemDelegate_updateEditorGeometry_default(void* self, void* editor, void* option, void* index);
int QAbstractItemDelegate_editorEvent_default(void* self, void* event, void* model, void* option, void* index);
int QAbstractItemDelegate_helpEvent_default(void* self, void* event, void* view, void* option, void* index);
int* QAbstractItemDelegate_paintingRoles_default(void* self, long long* __out_len);
void* QAbstractItemDelegate_metaObject_default(void* self);
int QAbstractItemDelegate_event_default(void* self, void* event);
int QAbstractItemDelegate_eventFilter_default(void* self, void* watched, void* event);
void QAbstractItemDelegate_timerEvent_default(void* self, void* event);
void QAbstractItemDelegate_childEvent_default(void* self, void* event);
void QAbstractItemDelegate_customEvent_default(void* self, void* event);
void QAbstractItemDelegate_connectNotify_default(void* self, void* signal);
void QAbstractItemDelegate_disconnectNotify_default(void* self, void* signal);
void QAbstractItemDelegate_connect_commitData(void* self, long long id);
void QAbstractItemDelegate_connect_closeEditor(void* self, long long id);
void QAbstractItemDelegate_connect_closeEditor2(void* self, long long id);
void QAbstractItemDelegate_connect_sizeHintChanged(void* self, long long id);
void QAbstractItemDelegate_connect_destroyed(void* self, long long id);
void QAbstractItemDelegate_connect_destroyed2(void* self, long long id);
void QAbstractItemDelegate_connect_objectNameChanged(void* self, long long id);
void QAbstractItemDelegate__disconnect(long long id);
void* QAbstractItemDelegate__Go_new(long long id, void* parent);
void QAbstractItemDelegate_delete(void* self);

#ifdef __cplusplus
}
#endif
