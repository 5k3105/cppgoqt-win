#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QTextFragment_new();
void* QTextFragment_new2(void* o);
void* QTextFragment_OperatorAssign(void* self, void* o);
int QTextFragment_isValid(void* self);
int QTextFragment_OperatorEqual(void* self, void* o);
int QTextFragment_OperatorNotEqual(void* self, void* o);
int QTextFragment_OperatorLesser(void* self, void* o);
int QTextFragment_position(void* self);
int QTextFragment_length(void* self);
int QTextFragment_contains(void* self, int position);
void* QTextFragment_charFormat(void* self);
int QTextFragment_charFormatIndex(void* self);
char* QTextFragment_text(void* self, long long* __out_len);
void** QTextFragment_glyphRuns(void* self, long long* __out_len);
void** QTextFragment_glyphRuns2(void* self, int from, long long* __out_len);
void** QTextFragment_glyphRuns3(void* self, int from, int length, long long* __out_len);
void QTextFragment_delete(void* self);

#ifdef __cplusplus
}
#endif
