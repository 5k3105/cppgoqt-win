#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QGraphicsGridLayout_base(void* self);
void QGraphicsGridLayout_chain(void* self, void** out);
void QGraphicsGridLayout__set_override(void* self, int idx);
void* QGraphicsGridLayout_new(long long __goid);
void* QGraphicsGridLayout_new2(long long __goid, void* parent);
void QGraphicsGridLayout_addItem(void* self, void* item, int row, int column, int rowSpan, int columnSpan);
void QGraphicsGridLayout_addItem2(void* self, void* item, int row, int column);
void QGraphicsGridLayout_setHorizontalSpacing(void* self, double spacing);
double QGraphicsGridLayout_horizontalSpacing(void* self);
void QGraphicsGridLayout_setVerticalSpacing(void* self, double spacing);
double QGraphicsGridLayout_verticalSpacing(void* self);
void QGraphicsGridLayout_setSpacing(void* self, double spacing);
void QGraphicsGridLayout_setRowSpacing(void* self, int row, double spacing);
double QGraphicsGridLayout_rowSpacing(void* self, int row);
void QGraphicsGridLayout_setColumnSpacing(void* self, int column, double spacing);
double QGraphicsGridLayout_columnSpacing(void* self, int column);
void QGraphicsGridLayout_setRowStretchFactor(void* self, int row, int stretch);
int QGraphicsGridLayout_rowStretchFactor(void* self, int row);
void QGraphicsGridLayout_setColumnStretchFactor(void* self, int column, int stretch);
int QGraphicsGridLayout_columnStretchFactor(void* self, int column);
void QGraphicsGridLayout_setRowMinimumHeight(void* self, int row, double height);
double QGraphicsGridLayout_rowMinimumHeight(void* self, int row);
void QGraphicsGridLayout_setRowPreferredHeight(void* self, int row, double height);
double QGraphicsGridLayout_rowPreferredHeight(void* self, int row);
void QGraphicsGridLayout_setRowMaximumHeight(void* self, int row, double height);
double QGraphicsGridLayout_rowMaximumHeight(void* self, int row);
void QGraphicsGridLayout_setRowFixedHeight(void* self, int row, double height);
void QGraphicsGridLayout_setColumnMinimumWidth(void* self, int column, double width);
double QGraphicsGridLayout_columnMinimumWidth(void* self, int column);
void QGraphicsGridLayout_setColumnPreferredWidth(void* self, int column, double width);
double QGraphicsGridLayout_columnPreferredWidth(void* self, int column);
void QGraphicsGridLayout_setColumnMaximumWidth(void* self, int column, double width);
double QGraphicsGridLayout_columnMaximumWidth(void* self, int column);
void QGraphicsGridLayout_setColumnFixedWidth(void* self, int column, double width);
void QGraphicsGridLayout_setRowAlignment(void* self, int row, int alignment);
int QGraphicsGridLayout_rowAlignment(void* self, int row);
void QGraphicsGridLayout_setColumnAlignment(void* self, int column, int alignment);
int QGraphicsGridLayout_columnAlignment(void* self, int column);
void QGraphicsGridLayout_setAlignment(void* self, void* item, int alignment);
int QGraphicsGridLayout_alignment(void* self, void* item);
int QGraphicsGridLayout_rowCount(void* self);
int QGraphicsGridLayout_columnCount(void* self);
void* QGraphicsGridLayout_itemAt(void* self, int row, int column);
void* QGraphicsGridLayout_itemAt2(void* self, int index);
void QGraphicsGridLayout_removeItem(void* self, void* item);
void QGraphicsGridLayout_setInstantInvalidatePropagation(int enable);
int QGraphicsGridLayout_instantInvalidatePropagation();
void QGraphicsGridLayout_addChildLayoutItem(void* self, void* layoutItem);
void QGraphicsGridLayout_setGraphicsItem(void* self, void* item);
void QGraphicsGridLayout_setOwnedByLayout(void* self, int ownedByLayout);
void QGraphicsGridLayout_addItem3(void* self, void* item, int row, int column, int rowSpan, int columnSpan, int alignment);
void QGraphicsGridLayout_addItem4(void* self, void* item, int row, int column, int alignment);
void* QGraphicsGridLayout_itemAt_default(void* self, int row, int column);
void* QGraphicsGridLayout_itemAt_default2(void* self, int index);
int QGraphicsGridLayout_count_default(void* self);
void QGraphicsGridLayout_removeAt_default(void* self, int index);
void QGraphicsGridLayout_invalidate_default(void* self);
void QGraphicsGridLayout_setGeometry_default(void* self, void* rect);
void* QGraphicsGridLayout_sizeHint_default(void* self, int which, void* constraint);
void QGraphicsGridLayout_getContentsMargins_default(void* self, double* left, double* top, double* right, double* bottom);
void QGraphicsGridLayout_updateGeometry_default(void* self);
void QGraphicsGridLayout_widgetEvent_default(void* self, void* e);
int QGraphicsGridLayout_isEmpty_default(void* self);
void* QGraphicsGridLayout__Go_new(long long id);
void QGraphicsGridLayout_delete(void* self);

#ifdef __cplusplus
}
#endif
