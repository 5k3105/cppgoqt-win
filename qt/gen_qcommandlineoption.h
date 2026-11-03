#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QCommandLineOption_new(const char* name, long long name_len);
void* QCommandLineOption_new2(char** names, long long names_len);
void* QCommandLineOption_new3(const char* name, long long name_len, const char* description, long long description_len);
void* QCommandLineOption_new4(char** names, long long names_len, const char* description, long long description_len);
void* QCommandLineOption_new5(void* other);
void* QCommandLineOption_new6(const char* name, long long name_len, const char* description, long long description_len, const char* valueName, long long valueName_len);
void* QCommandLineOption_new7(const char* name, long long name_len, const char* description, long long description_len, const char* valueName, long long valueName_len, const char* defaultValue, long long defaultValue_len);
void* QCommandLineOption_new8(char** names, long long names_len, const char* description, long long description_len, const char* valueName, long long valueName_len);
void* QCommandLineOption_new9(char** names, long long names_len, const char* description, long long description_len, const char* valueName, long long valueName_len, const char* defaultValue, long long defaultValue_len);
void* QCommandLineOption_OperatorAssign(void* self, void* other);
void QCommandLineOption_swap(void* self, void* other);
char** QCommandLineOption_names(void* self, long long* __out_len);
void QCommandLineOption_setValueName(void* self, const char* name, long long name_len);
char* QCommandLineOption_valueName(void* self, long long* __out_len);
void QCommandLineOption_setDescription(void* self, const char* description, long long description_len);
char* QCommandLineOption_description(void* self, long long* __out_len);
void QCommandLineOption_setDefaultValue(void* self, const char* defaultValue, long long defaultValue_len);
void QCommandLineOption_setDefaultValues(void* self, char** defaultValues, long long defaultValues_len);
char** QCommandLineOption_defaultValues(void* self, long long* __out_len);
int QCommandLineOption_flags(void* self);
void QCommandLineOption_setFlags(void* self, int aflags);
void QCommandLineOption_delete(void* self);

#ifdef __cplusplus
}
#endif
