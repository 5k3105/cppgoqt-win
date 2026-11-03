#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QGraphicsRectItem_base(void* self);
void QGraphicsRectItem_chain(void* self, void** out);
void QGraphicsRectItem__set_override(void* self, int idx);
void* QGraphicsRectItem_new(long long __goid);
void* QGraphicsRectItem_new2(long long __goid, void* rect);
void* QGraphicsRectItem_new3(long long __goid, double x, double y, double w, double h);
void* QGraphicsRectItem_new4(long long __goid, void* parent);
void* QGraphicsRectItem_new5(long long __goid, void* rect, void* parent);
void* QGraphicsRectItem_new6(long long __goid, double x, double y, double w, double h, void* parent);
void* QGraphicsRectItem_rect(void* self);
void QGraphicsRectItem_setRect(void* self, void* rect);
void QGraphicsRectItem_setRect2(void* self, double x, double y, double w, double h);
void QGraphicsRectItem_updateMicroFocus(void* self);
void QGraphicsRectItem_addToIndex(void* self);
void QGraphicsRectItem_removeFromIndex(void* self);
void QGraphicsRectItem_prepareGeometryChange(void* self);
void* QGraphicsRectItem_boundingRect_default(void* self);
void* QGraphicsRectItem_shape_default(void* self);
int QGraphicsRectItem_contains_default(void* self, void* point);
void QGraphicsRectItem_paint_default(void* self, void* painter, void* option, void* widget);
int QGraphicsRectItem_isObscuredBy_default(void* self, void* item);
void* QGraphicsRectItem_opaqueArea_default(void* self);
int QGraphicsRectItem_type_default(void* self);
int QGraphicsRectItem_supportsExtension_default(void* self, int extension);
void QGraphicsRectItem_setExtension_default(void* self, int extension, void* variant);
void* QGraphicsRectItem_extension_default(void* self, void* variant);
void QGraphicsRectItem_advance_default(void* self, int phase);
int QGraphicsRectItem_collidesWithItem_default(void* self, void* other, int mode);
int QGraphicsRectItem_collidesWithPath_default(void* self, void* path, int mode);
int QGraphicsRectItem_sceneEventFilter_default(void* self, void* watched, void* event);
int QGraphicsRectItem_sceneEvent_default(void* self, void* event);
void QGraphicsRectItem_contextMenuEvent_default(void* self, void* event);
void QGraphicsRectItem_dragEnterEvent_default(void* self, void* event);
void QGraphicsRectItem_dragLeaveEvent_default(void* self, void* event);
void QGraphicsRectItem_dragMoveEvent_default(void* self, void* event);
void QGraphicsRectItem_dropEvent_default(void* self, void* event);
void QGraphicsRectItem_focusInEvent_default(void* self, void* event);
void QGraphicsRectItem_focusOutEvent_default(void* self, void* event);
void QGraphicsRectItem_hoverEnterEvent_default(void* self, void* event);
void QGraphicsRectItem_hoverMoveEvent_default(void* self, void* event);
void QGraphicsRectItem_hoverLeaveEvent_default(void* self, void* event);
void QGraphicsRectItem_keyPressEvent_default(void* self, void* event);
void QGraphicsRectItem_keyReleaseEvent_default(void* self, void* event);
void QGraphicsRectItem_mousePressEvent_default(void* self, void* event);
void QGraphicsRectItem_mouseMoveEvent_default(void* self, void* event);
void QGraphicsRectItem_mouseReleaseEvent_default(void* self, void* event);
void QGraphicsRectItem_mouseDoubleClickEvent_default(void* self, void* event);
void QGraphicsRectItem_wheelEvent_default(void* self, void* event);
void QGraphicsRectItem_inputMethodEvent_default(void* self, void* event);
void* QGraphicsRectItem_inputMethodQuery_default(void* self, int query);
void* QGraphicsRectItem_itemChange_default(void* self, int change, void* value);
void* QGraphicsRectItem__Go_new(long long id);
void QGraphicsRectItem_delete(void* self);

#ifdef __cplusplus
}
#endif
