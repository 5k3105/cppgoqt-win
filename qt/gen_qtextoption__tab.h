#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QTextOption__Tab_new();
void* QTextOption__Tab_new2(double pos, int tabType);
void* QTextOption__Tab_new3(double pos, int tabType, void* delim);
int QTextOption__Tab_OperatorEqual(void* self, void* other);
int QTextOption__Tab_OperatorNotEqual(void* self, void* other);
double QTextOption__Tab_position_get(void* self);
void QTextOption__Tab_position_set(void* self, double position);
int QTextOption__Tab_type_get(void* self);
void QTextOption__Tab_type_set(void* self, int type_);
void* QTextOption__Tab_delimiter_get(void* self);
void QTextOption__Tab_delimiter_set(void* self, void* delimiter);
void QTextOption__Tab_delete(void* self);

#ifdef __cplusplus
}
#endif
