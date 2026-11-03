#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QGraphicsItemGroup_base(void* self);
void QGraphicsItemGroup__set_override(void* self, int idx);
void* QGraphicsItemGroup_new(long long __goid);
void* QGraphicsItemGroup_new2(long long __goid, void* parent);
void QGraphicsItemGroup_addToGroup(void* self, void* item);
void QGraphicsItemGroup_removeFromGroup(void* self, void* item);
void QGraphicsItemGroup_updateMicroFocus(void* self);
void QGraphicsItemGroup_addToIndex(void* self);
void QGraphicsItemGroup_removeFromIndex(void* self);
void QGraphicsItemGroup_prepareGeometryChange(void* self);
void* QGraphicsItemGroup_boundingRect_default(void* self);
void QGraphicsItemGroup_paint_default(void* self, void* painter, void* option, void* widget);
int QGraphicsItemGroup_isObscuredBy_default(void* self, void* item);
void* QGraphicsItemGroup_opaqueArea_default(void* self);
int QGraphicsItemGroup_type_default(void* self);
void QGraphicsItemGroup_advance_default(void* self, int phase);
void* QGraphicsItemGroup_shape_default(void* self);
int QGraphicsItemGroup_contains_default(void* self, void* point);
int QGraphicsItemGroup_collidesWithItem_default(void* self, void* other, int mode);
int QGraphicsItemGroup_collidesWithPath_default(void* self, void* path, int mode);
int QGraphicsItemGroup_sceneEventFilter_default(void* self, void* watched, void* event);
int QGraphicsItemGroup_sceneEvent_default(void* self, void* event);
void QGraphicsItemGroup_contextMenuEvent_default(void* self, void* event);
void QGraphicsItemGroup_dragEnterEvent_default(void* self, void* event);
void QGraphicsItemGroup_dragLeaveEvent_default(void* self, void* event);
void QGraphicsItemGroup_dragMoveEvent_default(void* self, void* event);
void QGraphicsItemGroup_dropEvent_default(void* self, void* event);
void QGraphicsItemGroup_focusInEvent_default(void* self, void* event);
void QGraphicsItemGroup_focusOutEvent_default(void* self, void* event);
void QGraphicsItemGroup_hoverEnterEvent_default(void* self, void* event);
void QGraphicsItemGroup_hoverMoveEvent_default(void* self, void* event);
void QGraphicsItemGroup_hoverLeaveEvent_default(void* self, void* event);
void QGraphicsItemGroup_keyPressEvent_default(void* self, void* event);
void QGraphicsItemGroup_keyReleaseEvent_default(void* self, void* event);
void QGraphicsItemGroup_mousePressEvent_default(void* self, void* event);
void QGraphicsItemGroup_mouseMoveEvent_default(void* self, void* event);
void QGraphicsItemGroup_mouseReleaseEvent_default(void* self, void* event);
void QGraphicsItemGroup_mouseDoubleClickEvent_default(void* self, void* event);
void QGraphicsItemGroup_wheelEvent_default(void* self, void* event);
void QGraphicsItemGroup_inputMethodEvent_default(void* self, void* event);
void* QGraphicsItemGroup_inputMethodQuery_default(void* self, int query);
void* QGraphicsItemGroup_itemChange_default(void* self, int change, void* value);
int QGraphicsItemGroup_supportsExtension_default(void* self, int extension);
void QGraphicsItemGroup_setExtension_default(void* self, int extension, void* variant);
void* QGraphicsItemGroup_extension_default(void* self, void* variant);
void* QGraphicsItemGroup__Go_new(long long id);
void QGraphicsItemGroup_delete(void* self);

#ifdef __cplusplus
}
#endif
