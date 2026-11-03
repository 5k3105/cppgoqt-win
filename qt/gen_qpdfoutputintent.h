#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QPdfOutputIntent_new();
void* QPdfOutputIntent_new2(void* other);
void* QPdfOutputIntent_OperatorAssign(void* self, void* other);
void QPdfOutputIntent_swap(void* self, void* other);
char* QPdfOutputIntent_outputConditionIdentifier(void* self, long long* __out_len);
void QPdfOutputIntent_setOutputConditionIdentifier(void* self, const char* identifier, long long identifier_len);
char* QPdfOutputIntent_outputCondition(void* self, long long* __out_len);
void QPdfOutputIntent_setOutputCondition(void* self, const char* condition, long long condition_len);
void* QPdfOutputIntent_registryName(void* self);
void QPdfOutputIntent_setRegistryName(void* self, void* name);
void* QPdfOutputIntent_outputProfile(void* self);
void QPdfOutputIntent_setOutputProfile(void* self, void* profile);
void QPdfOutputIntent_delete(void* self);

#ifdef __cplusplus
}
#endif
