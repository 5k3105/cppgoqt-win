#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QPageLayout_new();
void* QPageLayout_new2(void* pageSize, int orientation, void* margins);
void* QPageLayout_new3(void* other);
void* QPageLayout_new4(void* pageSize, int orientation, void* margins, int units);
void* QPageLayout_new5(void* pageSize, int orientation, void* margins, int units, void* minMargins);
void* QPageLayout_OperatorAssign(void* self, void* other);
void QPageLayout_swap(void* self, void* other);
int QPageLayout_isEquivalentTo(void* self, void* other);
int QPageLayout_isValid(void* self);
void QPageLayout_setMode(void* self, int mode);
int QPageLayout_mode(void* self);
void QPageLayout_setPageSize(void* self, void* pageSize);
void* QPageLayout_pageSize(void* self);
void QPageLayout_setOrientation(void* self, int orientation);
int QPageLayout_orientation(void* self);
void QPageLayout_setUnits(void* self, int units);
int QPageLayout_units(void* self);
int QPageLayout_setMargins(void* self, void* margins);
int QPageLayout_setLeftMargin(void* self, double leftMargin);
int QPageLayout_setRightMargin(void* self, double rightMargin);
int QPageLayout_setTopMargin(void* self, double topMargin);
int QPageLayout_setBottomMargin(void* self, double bottomMargin);
void* QPageLayout_margins(void* self);
void* QPageLayout_margins2(void* self, int units);
void* QPageLayout_marginsPoints(void* self);
void* QPageLayout_marginsPixels(void* self, int resolution);
void QPageLayout_setMinimumMargins(void* self, void* minMargins);
void* QPageLayout_minimumMargins(void* self);
void* QPageLayout_maximumMargins(void* self);
void* QPageLayout_fullRect(void* self);
void* QPageLayout_fullRect2(void* self, int units);
void* QPageLayout_fullRectPoints(void* self);
void* QPageLayout_fullRectPixels(void* self, int resolution);
void* QPageLayout_paintRect(void* self);
void* QPageLayout_paintRect2(void* self, int units);
void* QPageLayout_paintRectPoints(void* self);
void* QPageLayout_paintRectPixels(void* self, int resolution);
void QPageLayout_setPageSize2(void* self, void* pageSize, void* minMargins);
int QPageLayout_setMargins2(void* self, void* margins, int outOfBoundsPolicy);
int QPageLayout_setLeftMargin2(void* self, double leftMargin, int outOfBoundsPolicy);
int QPageLayout_setRightMargin2(void* self, double rightMargin, int outOfBoundsPolicy);
int QPageLayout_setTopMargin2(void* self, double topMargin, int outOfBoundsPolicy);
int QPageLayout_setBottomMargin2(void* self, double bottomMargin, int outOfBoundsPolicy);
int QPageLayout_OperatorEqual(void* self, void* rhs);
int QPageLayout_OperatorNotEqual(void* self, void* rhs);
void QPageLayout_delete(void* self);

#ifdef __cplusplus
}
#endif
