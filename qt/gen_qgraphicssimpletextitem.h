#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QGraphicsSimpleTextItem_base(void* self);
void QGraphicsSimpleTextItem_chain(void* self, void** out);
void QGraphicsSimpleTextItem__set_override(void* self, int idx);
void* QGraphicsSimpleTextItem_new(long long __goid);
void* QGraphicsSimpleTextItem_new2(long long __goid, const char* text, long long text_len);
void* QGraphicsSimpleTextItem_new3(long long __goid, void* parent);
void* QGraphicsSimpleTextItem_new4(long long __goid, const char* text, long long text_len, void* parent);
void QGraphicsSimpleTextItem_setText(void* self, const char* text, long long text_len);
char* QGraphicsSimpleTextItem_text(void* self, long long* __out_len);
void QGraphicsSimpleTextItem_setFont(void* self, void* font);
void* QGraphicsSimpleTextItem_font(void* self);
void QGraphicsSimpleTextItem_updateMicroFocus(void* self);
void QGraphicsSimpleTextItem_addToIndex(void* self);
void QGraphicsSimpleTextItem_removeFromIndex(void* self);
void QGraphicsSimpleTextItem_prepareGeometryChange(void* self);
void* QGraphicsSimpleTextItem_boundingRect_default(void* self);
void* QGraphicsSimpleTextItem_shape_default(void* self);
int QGraphicsSimpleTextItem_contains_default(void* self, void* point);
void QGraphicsSimpleTextItem_paint_default(void* self, void* painter, void* option, void* widget);
int QGraphicsSimpleTextItem_isObscuredBy_default(void* self, void* item);
void* QGraphicsSimpleTextItem_opaqueArea_default(void* self);
int QGraphicsSimpleTextItem_type_default(void* self);
int QGraphicsSimpleTextItem_supportsExtension_default(void* self, int extension);
void QGraphicsSimpleTextItem_setExtension_default(void* self, int extension, void* variant);
void* QGraphicsSimpleTextItem_extension_default(void* self, void* variant);
void QGraphicsSimpleTextItem_advance_default(void* self, int phase);
int QGraphicsSimpleTextItem_collidesWithItem_default(void* self, void* other, int mode);
int QGraphicsSimpleTextItem_collidesWithPath_default(void* self, void* path, int mode);
int QGraphicsSimpleTextItem_sceneEventFilter_default(void* self, void* watched, void* event);
int QGraphicsSimpleTextItem_sceneEvent_default(void* self, void* event);
void QGraphicsSimpleTextItem_contextMenuEvent_default(void* self, void* event);
void QGraphicsSimpleTextItem_dragEnterEvent_default(void* self, void* event);
void QGraphicsSimpleTextItem_dragLeaveEvent_default(void* self, void* event);
void QGraphicsSimpleTextItem_dragMoveEvent_default(void* self, void* event);
void QGraphicsSimpleTextItem_dropEvent_default(void* self, void* event);
void QGraphicsSimpleTextItem_focusInEvent_default(void* self, void* event);
void QGraphicsSimpleTextItem_focusOutEvent_default(void* self, void* event);
void QGraphicsSimpleTextItem_hoverEnterEvent_default(void* self, void* event);
void QGraphicsSimpleTextItem_hoverMoveEvent_default(void* self, void* event);
void QGraphicsSimpleTextItem_hoverLeaveEvent_default(void* self, void* event);
void QGraphicsSimpleTextItem_keyPressEvent_default(void* self, void* event);
void QGraphicsSimpleTextItem_keyReleaseEvent_default(void* self, void* event);
void QGraphicsSimpleTextItem_mousePressEvent_default(void* self, void* event);
void QGraphicsSimpleTextItem_mouseMoveEvent_default(void* self, void* event);
void QGraphicsSimpleTextItem_mouseReleaseEvent_default(void* self, void* event);
void QGraphicsSimpleTextItem_mouseDoubleClickEvent_default(void* self, void* event);
void QGraphicsSimpleTextItem_wheelEvent_default(void* self, void* event);
void QGraphicsSimpleTextItem_inputMethodEvent_default(void* self, void* event);
void* QGraphicsSimpleTextItem_inputMethodQuery_default(void* self, int query);
void* QGraphicsSimpleTextItem_itemChange_default(void* self, int change, void* value);
void* QGraphicsSimpleTextItem__Go_new(long long id);
void QGraphicsSimpleTextItem_delete(void* self);

#ifdef __cplusplus
}
#endif
