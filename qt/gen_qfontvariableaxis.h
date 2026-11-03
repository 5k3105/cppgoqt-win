#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QFontVariableAxis_new();
void* QFontVariableAxis_new2(void* axis);
void QFontVariableAxis_swap(void* self, void* other);
void* QFontVariableAxis_OperatorAssign(void* self, void* axis);
void* QFontVariableAxis_tag(void* self);
void QFontVariableAxis_setTag(void* self, void* tag);
char* QFontVariableAxis_name(void* self, long long* __out_len);
void QFontVariableAxis_setName(void* self, const char* name, long long name_len);
double QFontVariableAxis_minimumValue(void* self);
void QFontVariableAxis_setMinimumValue(void* self, double minimumValue);
double QFontVariableAxis_maximumValue(void* self);
void QFontVariableAxis_setMaximumValue(void* self, double maximumValue);
double QFontVariableAxis_defaultValue(void* self);
void QFontVariableAxis_setDefaultValue(void* self, double defaultValue);
void QFontVariableAxis_delete(void* self);

#ifdef __cplusplus
}
#endif
