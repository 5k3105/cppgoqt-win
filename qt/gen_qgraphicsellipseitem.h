#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QGraphicsEllipseItem_base(void* self);
void QGraphicsEllipseItem_chain(void* self, void** out);
void QGraphicsEllipseItem__set_override(void* self, int idx);
void* QGraphicsEllipseItem_new(long long __goid);
void* QGraphicsEllipseItem_new2(long long __goid, void* rect);
void* QGraphicsEllipseItem_new3(long long __goid, double x, double y, double w, double h);
void* QGraphicsEllipseItem_new4(long long __goid, void* parent);
void* QGraphicsEllipseItem_new5(long long __goid, void* rect, void* parent);
void* QGraphicsEllipseItem_new6(long long __goid, double x, double y, double w, double h, void* parent);
void* QGraphicsEllipseItem_rect(void* self);
void QGraphicsEllipseItem_setRect(void* self, void* rect);
void QGraphicsEllipseItem_setRect2(void* self, double x, double y, double w, double h);
int QGraphicsEllipseItem_startAngle(void* self);
void QGraphicsEllipseItem_setStartAngle(void* self, int angle);
int QGraphicsEllipseItem_spanAngle(void* self);
void QGraphicsEllipseItem_setSpanAngle(void* self, int angle);
void QGraphicsEllipseItem_updateMicroFocus(void* self);
void QGraphicsEllipseItem_addToIndex(void* self);
void QGraphicsEllipseItem_removeFromIndex(void* self);
void QGraphicsEllipseItem_prepareGeometryChange(void* self);
void* QGraphicsEllipseItem_boundingRect_default(void* self);
void* QGraphicsEllipseItem_shape_default(void* self);
int QGraphicsEllipseItem_contains_default(void* self, void* point);
void QGraphicsEllipseItem_paint_default(void* self, void* painter, void* option, void* widget);
int QGraphicsEllipseItem_isObscuredBy_default(void* self, void* item);
void* QGraphicsEllipseItem_opaqueArea_default(void* self);
int QGraphicsEllipseItem_type_default(void* self);
int QGraphicsEllipseItem_supportsExtension_default(void* self, int extension);
void QGraphicsEllipseItem_setExtension_default(void* self, int extension, void* variant);
void* QGraphicsEllipseItem_extension_default(void* self, void* variant);
void QGraphicsEllipseItem_advance_default(void* self, int phase);
int QGraphicsEllipseItem_collidesWithItem_default(void* self, void* other, int mode);
int QGraphicsEllipseItem_collidesWithPath_default(void* self, void* path, int mode);
int QGraphicsEllipseItem_sceneEventFilter_default(void* self, void* watched, void* event);
int QGraphicsEllipseItem_sceneEvent_default(void* self, void* event);
void QGraphicsEllipseItem_contextMenuEvent_default(void* self, void* event);
void QGraphicsEllipseItem_dragEnterEvent_default(void* self, void* event);
void QGraphicsEllipseItem_dragLeaveEvent_default(void* self, void* event);
void QGraphicsEllipseItem_dragMoveEvent_default(void* self, void* event);
void QGraphicsEllipseItem_dropEvent_default(void* self, void* event);
void QGraphicsEllipseItem_focusInEvent_default(void* self, void* event);
void QGraphicsEllipseItem_focusOutEvent_default(void* self, void* event);
void QGraphicsEllipseItem_hoverEnterEvent_default(void* self, void* event);
void QGraphicsEllipseItem_hoverMoveEvent_default(void* self, void* event);
void QGraphicsEllipseItem_hoverLeaveEvent_default(void* self, void* event);
void QGraphicsEllipseItem_keyPressEvent_default(void* self, void* event);
void QGraphicsEllipseItem_keyReleaseEvent_default(void* self, void* event);
void QGraphicsEllipseItem_mousePressEvent_default(void* self, void* event);
void QGraphicsEllipseItem_mouseMoveEvent_default(void* self, void* event);
void QGraphicsEllipseItem_mouseReleaseEvent_default(void* self, void* event);
void QGraphicsEllipseItem_mouseDoubleClickEvent_default(void* self, void* event);
void QGraphicsEllipseItem_wheelEvent_default(void* self, void* event);
void QGraphicsEllipseItem_inputMethodEvent_default(void* self, void* event);
void* QGraphicsEllipseItem_inputMethodQuery_default(void* self, int query);
void* QGraphicsEllipseItem_itemChange_default(void* self, int change, void* value);
void* QGraphicsEllipseItem__Go_new(long long id);
void QGraphicsEllipseItem_delete(void* self);

#ifdef __cplusplus
}
#endif
