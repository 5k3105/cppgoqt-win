#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QAbstractGraphicsShapeItem_base(void* self);
void QAbstractGraphicsShapeItem__set_override(void* self, int idx);
void* QAbstractGraphicsShapeItem_pen(void* self);
void QAbstractGraphicsShapeItem_setPen(void* self, void* pen);
void* QAbstractGraphicsShapeItem_brush(void* self);
void QAbstractGraphicsShapeItem_setBrush(void* self, void* brush);
void QAbstractGraphicsShapeItem_updateMicroFocus(void* self);
void QAbstractGraphicsShapeItem_addToIndex(void* self);
void QAbstractGraphicsShapeItem_removeFromIndex(void* self);
void QAbstractGraphicsShapeItem_prepareGeometryChange(void* self);
int QAbstractGraphicsShapeItem_isObscuredBy_default(void* self, void* item);
void* QAbstractGraphicsShapeItem_opaqueArea_default(void* self);
void QAbstractGraphicsShapeItem_advance_default(void* self, int phase);
void* QAbstractGraphicsShapeItem_shape_default(void* self);
int QAbstractGraphicsShapeItem_contains_default(void* self, void* point);
int QAbstractGraphicsShapeItem_collidesWithItem_default(void* self, void* other, int mode);
int QAbstractGraphicsShapeItem_collidesWithPath_default(void* self, void* path, int mode);
int QAbstractGraphicsShapeItem_type_default(void* self);
int QAbstractGraphicsShapeItem_sceneEventFilter_default(void* self, void* watched, void* event);
int QAbstractGraphicsShapeItem_sceneEvent_default(void* self, void* event);
void QAbstractGraphicsShapeItem_contextMenuEvent_default(void* self, void* event);
void QAbstractGraphicsShapeItem_dragEnterEvent_default(void* self, void* event);
void QAbstractGraphicsShapeItem_dragLeaveEvent_default(void* self, void* event);
void QAbstractGraphicsShapeItem_dragMoveEvent_default(void* self, void* event);
void QAbstractGraphicsShapeItem_dropEvent_default(void* self, void* event);
void QAbstractGraphicsShapeItem_focusInEvent_default(void* self, void* event);
void QAbstractGraphicsShapeItem_focusOutEvent_default(void* self, void* event);
void QAbstractGraphicsShapeItem_hoverEnterEvent_default(void* self, void* event);
void QAbstractGraphicsShapeItem_hoverMoveEvent_default(void* self, void* event);
void QAbstractGraphicsShapeItem_hoverLeaveEvent_default(void* self, void* event);
void QAbstractGraphicsShapeItem_keyPressEvent_default(void* self, void* event);
void QAbstractGraphicsShapeItem_keyReleaseEvent_default(void* self, void* event);
void QAbstractGraphicsShapeItem_mousePressEvent_default(void* self, void* event);
void QAbstractGraphicsShapeItem_mouseMoveEvent_default(void* self, void* event);
void QAbstractGraphicsShapeItem_mouseReleaseEvent_default(void* self, void* event);
void QAbstractGraphicsShapeItem_mouseDoubleClickEvent_default(void* self, void* event);
void QAbstractGraphicsShapeItem_wheelEvent_default(void* self, void* event);
void QAbstractGraphicsShapeItem_inputMethodEvent_default(void* self, void* event);
void* QAbstractGraphicsShapeItem_inputMethodQuery_default(void* self, int query);
void* QAbstractGraphicsShapeItem_itemChange_default(void* self, int change, void* value);
int QAbstractGraphicsShapeItem_supportsExtension_default(void* self, int extension);
void QAbstractGraphicsShapeItem_setExtension_default(void* self, int extension, void* variant);
void* QAbstractGraphicsShapeItem_extension_default(void* self, void* variant);
void* QAbstractGraphicsShapeItem__Go_new(long long id);
void QAbstractGraphicsShapeItem_delete(void* self);

#ifdef __cplusplus
}
#endif
