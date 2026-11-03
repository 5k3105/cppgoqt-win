#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QLoggingCategory_new(const char* category);
void* QLoggingCategory_new2(const char* category, int severityLevel);
int QLoggingCategory_isEnabled(void* self, int type_);
void QLoggingCategory_setEnabled(void* self, int type_, int enable);
int QLoggingCategory_isDebugEnabled(void* self);
int QLoggingCategory_isInfoEnabled(void* self);
int QLoggingCategory_isWarningEnabled(void* self);
int QLoggingCategory_isCriticalEnabled(void* self);
const char* QLoggingCategory_categoryName(void* self);
void* QLoggingCategory_OperatorCall(void* self);
void* QLoggingCategory_OperatorCall2(void* self);
void* QLoggingCategory_defaultCategory();
void QLoggingCategory_setFilterRules(const char* rules, long long rules_len);
void QLoggingCategory_delete(void* self);

#ifdef __cplusplus
}
#endif
