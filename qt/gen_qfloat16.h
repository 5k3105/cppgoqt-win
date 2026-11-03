#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* qfloat16_new();
void* qfloat16_new2(int arg1);
void* qfloat16_new3(float f);
void* qfloat16_new4(void* arg1);
float qfloat16_ToFloat(void* self);
int qfloat16_isInf(void* self);
int qfloat16_isNaN(void* self);
int qfloat16_isFinite(void* self);
int qfloat16_fpClassify(void* self);
void* qfloat16_copySign(void* self, void* sign);
void* qfloat16__limit_epsilon();
void* qfloat16__limit_min();
void* qfloat16__limit_denorm_min();
void* qfloat16__limit_max();
void* qfloat16__limit_lowest();
void* qfloat16__limit_infinity();
void* qfloat16__limit_quiet_NaN();
void* qfloat16__limit_signaling_NaN();
int qfloat16_isNormal(void* self);
void* qfloat16_OperatorMinus(void* self);
void* qfloat16_OperatorPlus(void* self, void* b);
void* qfloat16_OperatorMinus2(void* self, void* b);
void* qfloat16_OperatorMultiply(void* self, void* b);
void* qfloat16_OperatorDivide(void* self, void* b);
double qfloat16_OperatorPlus2(void* self, double rhs);
double qfloat16_OperatorMinus3(void* self, double rhs);
double qfloat16_OperatorMultiply2(void* self, double rhs);
double qfloat16_OperatorDivide2(void* self, double rhs);
void* qfloat16_OperatorPlusAssign(void* self, double rhs);
void* qfloat16_OperatorMinusAssign(void* self, double rhs);
void* qfloat16_OperatorMultiplyAssign(void* self, double rhs);
void* qfloat16_OperatorDivideAssign(void* self, double rhs);
float qfloat16_OperatorPlus3(void* self, float rhs);
float qfloat16_OperatorMinus4(void* self, float rhs);
float qfloat16_OperatorMultiply3(void* self, float rhs);
float qfloat16_OperatorDivide3(void* self, float rhs);
void* qfloat16_OperatorPlusAssign2(void* self, float rhs);
void* qfloat16_OperatorMinusAssign2(void* self, float rhs);
void* qfloat16_OperatorMultiplyAssign2(void* self, float rhs);
void* qfloat16_OperatorDivideAssign2(void* self, float rhs);
double qfloat16_OperatorPlus4(void* self, int rhs);
double qfloat16_OperatorMinus5(void* self, int rhs);
double qfloat16_OperatorMultiply4(void* self, int rhs);
double qfloat16_OperatorDivide4(void* self, int rhs);
int qfloat16_OperatorEqual(void* self, void* rhs);
int qfloat16_OperatorNotEqual(void* self, void* rhs);
int qfloat16_OperatorLesser(void* self, void* rhs);
int qfloat16_OperatorGreater(void* self, void* rhs);
int qfloat16_OperatorLesserOrEqual(void* self, void* rhs);
int qfloat16_OperatorGreaterOrEqual(void* self, void* rhs);
void qfloat16_delete(void* self);

#ifdef __cplusplus
}
#endif
