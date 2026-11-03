#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QRegularExpressionMatch_new();
void* QRegularExpressionMatch_new2(void* match);
void* QRegularExpressionMatch_OperatorAssign(void* self, void* match);
void QRegularExpressionMatch_swap(void* self, void* other);
void* QRegularExpressionMatch_regularExpression(void* self);
int QRegularExpressionMatch_matchType(void* self);
int QRegularExpressionMatch_matchOptions(void* self);
int QRegularExpressionMatch_hasMatch(void* self);
int QRegularExpressionMatch_hasPartialMatch(void* self);
int QRegularExpressionMatch_isValid(void* self);
int QRegularExpressionMatch_lastCapturedIndex(void* self);
int QRegularExpressionMatch_hasCaptured(void* self, const char* name, long long name_len);
int QRegularExpressionMatch_hasCaptured2(void* self, int nth);
char* QRegularExpressionMatch_captured(void* self, long long* __out_len);
void* QRegularExpressionMatch_capturedView(void* self);
char* QRegularExpressionMatch_captured2(void* self, const char* name, long long name_len, long long* __out_len);
void* QRegularExpressionMatch_capturedView2(void* self, const char* name, long long name_len);
char** QRegularExpressionMatch_capturedTexts(void* self, long long* __out_len);
long long QRegularExpressionMatch_capturedStart(void* self);
long long QRegularExpressionMatch_capturedLength(void* self);
long long QRegularExpressionMatch_capturedEnd(void* self);
long long QRegularExpressionMatch_capturedStart2(void* self, const char* name, long long name_len);
long long QRegularExpressionMatch_capturedLength2(void* self, const char* name, long long name_len);
long long QRegularExpressionMatch_capturedEnd2(void* self, const char* name, long long name_len);
char* QRegularExpressionMatch_captured3(void* self, int nth, long long* __out_len);
void* QRegularExpressionMatch_capturedView3(void* self, int nth);
long long QRegularExpressionMatch_capturedStart3(void* self, int nth);
long long QRegularExpressionMatch_capturedLength3(void* self, int nth);
long long QRegularExpressionMatch_capturedEnd3(void* self, int nth);
void QRegularExpressionMatch_delete(void* self);

#ifdef __cplusplus
}
#endif
