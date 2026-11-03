#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QGraphicsPixmapItem_base(void* self);
void QGraphicsPixmapItem__set_override(void* self, int idx);
void* QGraphicsPixmapItem_new(long long __goid);
void* QGraphicsPixmapItem_new2(long long __goid, void* pixmap);
void* QGraphicsPixmapItem_new3(long long __goid, void* parent);
void* QGraphicsPixmapItem_new4(long long __goid, void* pixmap, void* parent);
void* QGraphicsPixmapItem_pixmap(void* self);
void QGraphicsPixmapItem_setPixmap(void* self, void* pixmap);
int QGraphicsPixmapItem_transformationMode(void* self);
void QGraphicsPixmapItem_setTransformationMode(void* self, int mode);
void* QGraphicsPixmapItem_offset(void* self);
void QGraphicsPixmapItem_setOffset(void* self, void* offset);
void QGraphicsPixmapItem_setOffset2(void* self, double x, double y);
int QGraphicsPixmapItem_shapeMode(void* self);
void QGraphicsPixmapItem_setShapeMode(void* self, int mode);
void QGraphicsPixmapItem_updateMicroFocus(void* self);
void QGraphicsPixmapItem_addToIndex(void* self);
void QGraphicsPixmapItem_removeFromIndex(void* self);
void QGraphicsPixmapItem_prepareGeometryChange(void* self);
void* QGraphicsPixmapItem_boundingRect_default(void* self);
void* QGraphicsPixmapItem_shape_default(void* self);
int QGraphicsPixmapItem_contains_default(void* self, void* point);
void QGraphicsPixmapItem_paint_default(void* self, void* painter, void* option, void* widget);
int QGraphicsPixmapItem_isObscuredBy_default(void* self, void* item);
void* QGraphicsPixmapItem_opaqueArea_default(void* self);
int QGraphicsPixmapItem_type_default(void* self);
int QGraphicsPixmapItem_supportsExtension_default(void* self, int extension);
void QGraphicsPixmapItem_setExtension_default(void* self, int extension, void* variant);
void* QGraphicsPixmapItem_extension_default(void* self, void* variant);
void QGraphicsPixmapItem_advance_default(void* self, int phase);
int QGraphicsPixmapItem_collidesWithItem_default(void* self, void* other, int mode);
int QGraphicsPixmapItem_collidesWithPath_default(void* self, void* path, int mode);
int QGraphicsPixmapItem_sceneEventFilter_default(void* self, void* watched, void* event);
int QGraphicsPixmapItem_sceneEvent_default(void* self, void* event);
void QGraphicsPixmapItem_contextMenuEvent_default(void* self, void* event);
void QGraphicsPixmapItem_dragEnterEvent_default(void* self, void* event);
void QGraphicsPixmapItem_dragLeaveEvent_default(void* self, void* event);
void QGraphicsPixmapItem_dragMoveEvent_default(void* self, void* event);
void QGraphicsPixmapItem_dropEvent_default(void* self, void* event);
void QGraphicsPixmapItem_focusInEvent_default(void* self, void* event);
void QGraphicsPixmapItem_focusOutEvent_default(void* self, void* event);
void QGraphicsPixmapItem_hoverEnterEvent_default(void* self, void* event);
void QGraphicsPixmapItem_hoverMoveEvent_default(void* self, void* event);
void QGraphicsPixmapItem_hoverLeaveEvent_default(void* self, void* event);
void QGraphicsPixmapItem_keyPressEvent_default(void* self, void* event);
void QGraphicsPixmapItem_keyReleaseEvent_default(void* self, void* event);
void QGraphicsPixmapItem_mousePressEvent_default(void* self, void* event);
void QGraphicsPixmapItem_mouseMoveEvent_default(void* self, void* event);
void QGraphicsPixmapItem_mouseReleaseEvent_default(void* self, void* event);
void QGraphicsPixmapItem_mouseDoubleClickEvent_default(void* self, void* event);
void QGraphicsPixmapItem_wheelEvent_default(void* self, void* event);
void QGraphicsPixmapItem_inputMethodEvent_default(void* self, void* event);
void* QGraphicsPixmapItem_inputMethodQuery_default(void* self, int query);
void* QGraphicsPixmapItem_itemChange_default(void* self, int change, void* value);
void* QGraphicsPixmapItem__Go_new(long long id);
void QGraphicsPixmapItem_delete(void* self);

#ifdef __cplusplus
}
#endif
