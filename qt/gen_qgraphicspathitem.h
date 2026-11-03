#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QGraphicsPathItem_base(void* self);
void QGraphicsPathItem_chain(void* self, void** out);
void QGraphicsPathItem__set_override(void* self, int idx);
void* QGraphicsPathItem_new(long long __goid);
void* QGraphicsPathItem_new2(long long __goid, void* path);
void* QGraphicsPathItem_new3(long long __goid, void* parent);
void* QGraphicsPathItem_new4(long long __goid, void* path, void* parent);
void* QGraphicsPathItem_path(void* self);
void QGraphicsPathItem_setPath(void* self, void* path);
void QGraphicsPathItem_updateMicroFocus(void* self);
void QGraphicsPathItem_addToIndex(void* self);
void QGraphicsPathItem_removeFromIndex(void* self);
void QGraphicsPathItem_prepareGeometryChange(void* self);
void* QGraphicsPathItem_boundingRect_default(void* self);
void* QGraphicsPathItem_shape_default(void* self);
int QGraphicsPathItem_contains_default(void* self, void* point);
void QGraphicsPathItem_paint_default(void* self, void* painter, void* option, void* widget);
int QGraphicsPathItem_isObscuredBy_default(void* self, void* item);
void* QGraphicsPathItem_opaqueArea_default(void* self);
int QGraphicsPathItem_type_default(void* self);
int QGraphicsPathItem_supportsExtension_default(void* self, int extension);
void QGraphicsPathItem_setExtension_default(void* self, int extension, void* variant);
void* QGraphicsPathItem_extension_default(void* self, void* variant);
void QGraphicsPathItem_advance_default(void* self, int phase);
int QGraphicsPathItem_collidesWithItem_default(void* self, void* other, int mode);
int QGraphicsPathItem_collidesWithPath_default(void* self, void* path, int mode);
int QGraphicsPathItem_sceneEventFilter_default(void* self, void* watched, void* event);
int QGraphicsPathItem_sceneEvent_default(void* self, void* event);
void QGraphicsPathItem_contextMenuEvent_default(void* self, void* event);
void QGraphicsPathItem_dragEnterEvent_default(void* self, void* event);
void QGraphicsPathItem_dragLeaveEvent_default(void* self, void* event);
void QGraphicsPathItem_dragMoveEvent_default(void* self, void* event);
void QGraphicsPathItem_dropEvent_default(void* self, void* event);
void QGraphicsPathItem_focusInEvent_default(void* self, void* event);
void QGraphicsPathItem_focusOutEvent_default(void* self, void* event);
void QGraphicsPathItem_hoverEnterEvent_default(void* self, void* event);
void QGraphicsPathItem_hoverMoveEvent_default(void* self, void* event);
void QGraphicsPathItem_hoverLeaveEvent_default(void* self, void* event);
void QGraphicsPathItem_keyPressEvent_default(void* self, void* event);
void QGraphicsPathItem_keyReleaseEvent_default(void* self, void* event);
void QGraphicsPathItem_mousePressEvent_default(void* self, void* event);
void QGraphicsPathItem_mouseMoveEvent_default(void* self, void* event);
void QGraphicsPathItem_mouseReleaseEvent_default(void* self, void* event);
void QGraphicsPathItem_mouseDoubleClickEvent_default(void* self, void* event);
void QGraphicsPathItem_wheelEvent_default(void* self, void* event);
void QGraphicsPathItem_inputMethodEvent_default(void* self, void* event);
void* QGraphicsPathItem_inputMethodQuery_default(void* self, int query);
void* QGraphicsPathItem_itemChange_default(void* self, int change, void* value);
void* QGraphicsPathItem__Go_new(long long id);
void QGraphicsPathItem_delete(void* self);

#ifdef __cplusplus
}
#endif
