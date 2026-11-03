#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QGraphicsPolygonItem_base(void* self);
void QGraphicsPolygonItem_chain(void* self, void** out);
void QGraphicsPolygonItem__set_override(void* self, int idx);
void* QGraphicsPolygonItem_new(long long __goid);
void* QGraphicsPolygonItem_new2(long long __goid, void* polygon);
void* QGraphicsPolygonItem_new3(long long __goid, void* parent);
void* QGraphicsPolygonItem_new4(long long __goid, void* polygon, void* parent);
void* QGraphicsPolygonItem_polygon(void* self);
void QGraphicsPolygonItem_setPolygon(void* self, void* polygon);
int QGraphicsPolygonItem_fillRule(void* self);
void QGraphicsPolygonItem_setFillRule(void* self, int rule);
void QGraphicsPolygonItem_updateMicroFocus(void* self);
void QGraphicsPolygonItem_addToIndex(void* self);
void QGraphicsPolygonItem_removeFromIndex(void* self);
void QGraphicsPolygonItem_prepareGeometryChange(void* self);
void* QGraphicsPolygonItem_boundingRect_default(void* self);
void* QGraphicsPolygonItem_shape_default(void* self);
int QGraphicsPolygonItem_contains_default(void* self, void* point);
void QGraphicsPolygonItem_paint_default(void* self, void* painter, void* option, void* widget);
int QGraphicsPolygonItem_isObscuredBy_default(void* self, void* item);
void* QGraphicsPolygonItem_opaqueArea_default(void* self);
int QGraphicsPolygonItem_type_default(void* self);
int QGraphicsPolygonItem_supportsExtension_default(void* self, int extension);
void QGraphicsPolygonItem_setExtension_default(void* self, int extension, void* variant);
void* QGraphicsPolygonItem_extension_default(void* self, void* variant);
void QGraphicsPolygonItem_advance_default(void* self, int phase);
int QGraphicsPolygonItem_collidesWithItem_default(void* self, void* other, int mode);
int QGraphicsPolygonItem_collidesWithPath_default(void* self, void* path, int mode);
int QGraphicsPolygonItem_sceneEventFilter_default(void* self, void* watched, void* event);
int QGraphicsPolygonItem_sceneEvent_default(void* self, void* event);
void QGraphicsPolygonItem_contextMenuEvent_default(void* self, void* event);
void QGraphicsPolygonItem_dragEnterEvent_default(void* self, void* event);
void QGraphicsPolygonItem_dragLeaveEvent_default(void* self, void* event);
void QGraphicsPolygonItem_dragMoveEvent_default(void* self, void* event);
void QGraphicsPolygonItem_dropEvent_default(void* self, void* event);
void QGraphicsPolygonItem_focusInEvent_default(void* self, void* event);
void QGraphicsPolygonItem_focusOutEvent_default(void* self, void* event);
void QGraphicsPolygonItem_hoverEnterEvent_default(void* self, void* event);
void QGraphicsPolygonItem_hoverMoveEvent_default(void* self, void* event);
void QGraphicsPolygonItem_hoverLeaveEvent_default(void* self, void* event);
void QGraphicsPolygonItem_keyPressEvent_default(void* self, void* event);
void QGraphicsPolygonItem_keyReleaseEvent_default(void* self, void* event);
void QGraphicsPolygonItem_mousePressEvent_default(void* self, void* event);
void QGraphicsPolygonItem_mouseMoveEvent_default(void* self, void* event);
void QGraphicsPolygonItem_mouseReleaseEvent_default(void* self, void* event);
void QGraphicsPolygonItem_mouseDoubleClickEvent_default(void* self, void* event);
void QGraphicsPolygonItem_wheelEvent_default(void* self, void* event);
void QGraphicsPolygonItem_inputMethodEvent_default(void* self, void* event);
void* QGraphicsPolygonItem_inputMethodQuery_default(void* self, int query);
void* QGraphicsPolygonItem_itemChange_default(void* self, int change, void* value);
void* QGraphicsPolygonItem__Go_new(long long id);
void QGraphicsPolygonItem_delete(void* self);

#ifdef __cplusplus
}
#endif
