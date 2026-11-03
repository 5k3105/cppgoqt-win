#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QLatin1Char_new(char c);
void* QLatin1Char_new2(void* arg1);
char QLatin1Char_toLatin1(void* self);
unsigned short QLatin1Char_unicode(void* self);
int QLatin1Char_OperatorEqual(void* self, void* rhs);
int QLatin1Char_OperatorNotEqual(void* self, void* rhs);
int QLatin1Char_OperatorLesser(void* self, void* rhs);
int QLatin1Char_OperatorGreater(void* self, void* rhs);
int QLatin1Char_OperatorLesserOrEqual(void* self, void* rhs);
int QLatin1Char_OperatorGreaterOrEqual(void* self, void* rhs);
void QLatin1Char_delete(void* self);

#ifdef __cplusplus
}
#endif
