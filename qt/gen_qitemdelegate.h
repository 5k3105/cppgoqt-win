#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QItemDelegate_base(void* self);
void QItemDelegate_chain(void* self, void** out);
void QItemDelegate__set_override(void* self, int idx);
void* QItemDelegate_fromQObject(void* o);
void* QItemDelegate_new(long long __goid);
void* QItemDelegate_new2(long long __goid, void* parent);
char* QItemDelegate_tr(const char* s, long long* __out_len);
int QItemDelegate_hasClipping(void* self);
void QItemDelegate_setClipping(void* self, int clip);
void* QItemDelegate_itemEditorFactory(void* self);
void QItemDelegate_setItemEditorFactory(void* self, void* factory);
void QItemDelegate_drawBackground(void* self, void* painter, void* option, void* index);
void QItemDelegate_doLayout(void* self, void* option, void* checkRect, void* iconRect, void* textRect, int hint);
void* QItemDelegate_rect(void* self, void* option, void* index, int role);
void* QItemDelegate_setOptions(void* self, void* index, void* option);
void* QItemDelegate_decoration(void* self, void* option, void* variant);
void* QItemDelegate_doCheck(void* self, void* option, void* bounding, void* variant);
void* QItemDelegate_textRectangle(void* self, void* painter, void* rect, void* font, const char* text, long long text_len);
void* QItemDelegate_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QItemDelegate_connect2(void* sender, void* signal, void* receiver, void* method);
void* QItemDelegate_sender(void* self);
int QItemDelegate_senderSignalIndex(void* self);
int QItemDelegate_receivers(void* self, const char* signal);
int QItemDelegate_isSignalConnected(void* self, void* signal);
char* QItemDelegate_tr2(const char* s, const char* c, long long* __out_len);
char* QItemDelegate_tr3(const char* s, const char* c, int n, long long* __out_len);
void* QItemDelegate_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QItemDelegate_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void* QItemDelegate_metaObject_default(void* self);
void QItemDelegate_paint_default(void* self, void* painter, void* option, void* index);
void* QItemDelegate_sizeHint_default(void* self, void* option, void* index);
void* QItemDelegate_createEditor_default(void* self, void* parent, void* option, void* index);
void QItemDelegate_setEditorData_default(void* self, void* editor, void* index);
void QItemDelegate_setModelData_default(void* self, void* editor, void* model, void* index);
void QItemDelegate_updateEditorGeometry_default(void* self, void* editor, void* option, void* index);
void QItemDelegate_drawDisplay_default(void* self, void* painter, void* option, void* rect, const char* text, long long text_len);
void QItemDelegate_drawDecoration_default(void* self, void* painter, void* option, void* rect, void* pixmap);
void QItemDelegate_drawFocus_default(void* self, void* painter, void* option, void* rect);
void QItemDelegate_drawCheck_default(void* self, void* painter, void* option, void* rect, int state);
int QItemDelegate_eventFilter_default(void* self, void* object, void* event);
int QItemDelegate_editorEvent_default(void* self, void* event, void* model, void* option, void* index);
void QItemDelegate_destroyEditor_default(void* self, void* editor, void* index);
int QItemDelegate_helpEvent_default(void* self, void* event, void* view, void* option, void* index);
int* QItemDelegate_paintingRoles_default(void* self, long long* __out_len);
int QItemDelegate_event_default(void* self, void* event);
void QItemDelegate_timerEvent_default(void* self, void* event);
void QItemDelegate_childEvent_default(void* self, void* event);
void QItemDelegate_customEvent_default(void* self, void* event);
void QItemDelegate_connectNotify_default(void* self, void* signal);
void QItemDelegate_disconnectNotify_default(void* self, void* signal);
void QItemDelegate_connect_commitData(void* self, long long id);
void QItemDelegate_connect_closeEditor(void* self, long long id);
void QItemDelegate_connect_closeEditor2(void* self, long long id);
void QItemDelegate_connect_sizeHintChanged(void* self, long long id);
void QItemDelegate_connect_destroyed(void* self, long long id);
void QItemDelegate_connect_destroyed2(void* self, long long id);
void QItemDelegate_connect_objectNameChanged(void* self, long long id);
void QItemDelegate__disconnect(long long id);
void* QItemDelegate__Go_new(long long id, void* parent);
void QItemDelegate_delete(void* self);

#ifdef __cplusplus
}
#endif
