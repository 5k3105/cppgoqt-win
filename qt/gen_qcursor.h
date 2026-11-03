#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QCursor_new();
void* QCursor_new2(int shape);
void* QCursor_new3(void* bitmap, void* mask);
void* QCursor_new4(void* pixmap);
void* QCursor_new5(void* cursor);
void* QCursor_new6(void* bitmap, void* mask, int hotX);
void* QCursor_new7(void* bitmap, void* mask, int hotX, int hotY);
void* QCursor_new8(void* pixmap, int hotX);
void* QCursor_new9(void* pixmap, int hotX, int hotY);
void* QCursor_OperatorAssign(void* self, void* cursor);
void QCursor_swap(void* self, void* other);
void* QCursor_ToQVariant(void* self);
int QCursor_shape(void* self);
void QCursor_setShape(void* self, int newShape);
void* QCursor_bitmap(void* self, int arg1);
void* QCursor_mask(void* self, int arg1);
void* QCursor_bitmap2(void* self);
void* QCursor_mask2(void* self);
void* QCursor_pixmap(void* self);
void* QCursor_hotSpot(void* self);
void* QCursor_pos();
void* QCursor_pos2(void* screen);
void QCursor_setPos(int x, int y);
void QCursor_setPos2(void* screen, int x, int y);
void QCursor_setPos3(void* p);
void QCursor_setPos4(void* screen, void* p);
int QCursor_OperatorEqual(void* self, void* rhs);
int QCursor_OperatorNotEqual(void* self, void* rhs);
void* QCursor_toQVariant(void* self);
void QCursor_delete(void* self);

#ifdef __cplusplus
}
#endif
