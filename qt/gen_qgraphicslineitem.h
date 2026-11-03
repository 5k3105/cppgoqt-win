#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QGraphicsLineItem_base(void* self);
void QGraphicsLineItem__set_override(void* self, int idx);
void* QGraphicsLineItem_new(long long __goid);
void* QGraphicsLineItem_new2(long long __goid, void* line);
void* QGraphicsLineItem_new3(long long __goid, double x1, double y1, double x2, double y2);
void* QGraphicsLineItem_new4(long long __goid, void* parent);
void* QGraphicsLineItem_new5(long long __goid, void* line, void* parent);
void* QGraphicsLineItem_new6(long long __goid, double x1, double y1, double x2, double y2, void* parent);
void* QGraphicsLineItem_pen(void* self);
void QGraphicsLineItem_setPen(void* self, void* pen);
void* QGraphicsLineItem_line(void* self);
void QGraphicsLineItem_setLine(void* self, void* line);
void QGraphicsLineItem_setLine2(void* self, double x1, double y1, double x2, double y2);
void QGraphicsLineItem_updateMicroFocus(void* self);
void QGraphicsLineItem_addToIndex(void* self);
void QGraphicsLineItem_removeFromIndex(void* self);
void QGraphicsLineItem_prepareGeometryChange(void* self);
void* QGraphicsLineItem_boundingRect_default(void* self);
void* QGraphicsLineItem_shape_default(void* self);
int QGraphicsLineItem_contains_default(void* self, void* point);
void QGraphicsLineItem_paint_default(void* self, void* painter, void* option, void* widget);
int QGraphicsLineItem_isObscuredBy_default(void* self, void* item);
void* QGraphicsLineItem_opaqueArea_default(void* self);
int QGraphicsLineItem_type_default(void* self);
int QGraphicsLineItem_supportsExtension_default(void* self, int extension);
void QGraphicsLineItem_setExtension_default(void* self, int extension, void* variant);
void* QGraphicsLineItem_extension_default(void* self, void* variant);
void QGraphicsLineItem_advance_default(void* self, int phase);
int QGraphicsLineItem_collidesWithItem_default(void* self, void* other, int mode);
int QGraphicsLineItem_collidesWithPath_default(void* self, void* path, int mode);
int QGraphicsLineItem_sceneEventFilter_default(void* self, void* watched, void* event);
int QGraphicsLineItem_sceneEvent_default(void* self, void* event);
void QGraphicsLineItem_contextMenuEvent_default(void* self, void* event);
void QGraphicsLineItem_dragEnterEvent_default(void* self, void* event);
void QGraphicsLineItem_dragLeaveEvent_default(void* self, void* event);
void QGraphicsLineItem_dragMoveEvent_default(void* self, void* event);
void QGraphicsLineItem_dropEvent_default(void* self, void* event);
void QGraphicsLineItem_focusInEvent_default(void* self, void* event);
void QGraphicsLineItem_focusOutEvent_default(void* self, void* event);
void QGraphicsLineItem_hoverEnterEvent_default(void* self, void* event);
void QGraphicsLineItem_hoverMoveEvent_default(void* self, void* event);
void QGraphicsLineItem_hoverLeaveEvent_default(void* self, void* event);
void QGraphicsLineItem_keyPressEvent_default(void* self, void* event);
void QGraphicsLineItem_keyReleaseEvent_default(void* self, void* event);
void QGraphicsLineItem_mousePressEvent_default(void* self, void* event);
void QGraphicsLineItem_mouseMoveEvent_default(void* self, void* event);
void QGraphicsLineItem_mouseReleaseEvent_default(void* self, void* event);
void QGraphicsLineItem_mouseDoubleClickEvent_default(void* self, void* event);
void QGraphicsLineItem_wheelEvent_default(void* self, void* event);
void QGraphicsLineItem_inputMethodEvent_default(void* self, void* event);
void* QGraphicsLineItem_inputMethodQuery_default(void* self, int query);
void* QGraphicsLineItem_itemChange_default(void* self, int change, void* value);
void* QGraphicsLineItem__Go_new(long long id);
void QGraphicsLineItem_delete(void* self);

#ifdef __cplusplus
}
#endif
