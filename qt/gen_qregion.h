#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QRegion_new();
void* QRegion_new2(int x, int y, int w, int h);
void* QRegion_new3(void* r);
void* QRegion_new4(void* pa);
void* QRegion_new5(void* region);
void* QRegion_new6(void* bitmap);
void* QRegion_new7(int x, int y, int w, int h, int t);
void* QRegion_new8(void* r, int t);
void* QRegion_new9(void* pa, int fillRule);
void* QRegion_OperatorAssign(void* self, void* arg1);
void QRegion_swap(void* self, void* other);
int QRegion_isEmpty(void* self);
int QRegion_isNull(void* self);
void* QRegion_begin(void* self);
void* QRegion_cbegin(void* self);
void* QRegion_end(void* self);
void* QRegion_cend(void* self);
int QRegion_contains(void* self, void* p);
int QRegion_contains2(void* self, void* r);
void QRegion_translate(void* self, int dx, int dy);
void QRegion_translate2(void* self, void* p);
void* QRegion_translated(void* self, int dx, int dy);
void* QRegion_translated2(void* self, void* p);
void* QRegion_united(void* self, void* r);
void* QRegion_united2(void* self, void* r);
void* QRegion_intersected(void* self, void* r);
void* QRegion_intersected2(void* self, void* r);
void* QRegion_subtracted(void* self, void* r);
void* QRegion_xored(void* self, void* r);
int QRegion_intersects(void* self, void* r);
int QRegion_intersects2(void* self, void* r);
void* QRegion_boundingRect(void* self);
void QRegion_setRects(void* self, void* rect, int num);
int QRegion_rectCount(void* self);
void* QRegion_OperatorBitwiseOr(void* self, void* r);
void* QRegion_OperatorPlus(void* self, void* r);
void* QRegion_OperatorPlus2(void* self, void* r);
void* QRegion_OperatorBitwiseAnd(void* self, void* r);
void* QRegion_OperatorBitwiseAnd2(void* self, void* r);
void* QRegion_OperatorMinus(void* self, void* r);
void* QRegion_OperatorBitwiseXor(void* self, void* r);
void* QRegion_OperatorBitwiseOrAssign(void* self, void* r);
void* QRegion_OperatorPlusAssign(void* self, void* r);
void* QRegion_OperatorPlusAssign2(void* self, void* r);
void* QRegion_OperatorBitwiseAndAssign(void* self, void* r);
void* QRegion_OperatorBitwiseAndAssign2(void* self, void* r);
void* QRegion_OperatorMinusAssign(void* self, void* r);
void* QRegion_OperatorBitwiseXorAssign(void* self, void* r);
int QRegion_OperatorEqual(void* self, void* r);
int QRegion_OperatorNotEqual(void* self, void* r);
void* QRegion_ToQVariant(void* self);
void* QRegion_OperatorMultiply(void* self, void* m);
void* QRegion_toQVariant(void* self);
void QRegion_delete(void* self);

#ifdef __cplusplus
}
#endif
