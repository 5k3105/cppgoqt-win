#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QRegularExpression_new();
void* QRegularExpression_new2(const char* pattern, long long pattern_len);
void* QRegularExpression_new3(void* re);
void* QRegularExpression_new4(const char* pattern, long long pattern_len, int options);
int QRegularExpression_patternOptions(void* self);
void QRegularExpression_setPatternOptions(void* self, int options);
void* QRegularExpression_OperatorAssign(void* self, void* re);
void QRegularExpression_swap(void* self, void* other);
char* QRegularExpression_pattern(void* self, long long* __out_len);
void QRegularExpression_setPattern(void* self, const char* pattern, long long pattern_len);
int QRegularExpression_isValid(void* self);
long long QRegularExpression_patternErrorOffset(void* self);
char* QRegularExpression_errorString(void* self, long long* __out_len);
int QRegularExpression_captureCount(void* self);
char** QRegularExpression_namedCaptureGroups(void* self, long long* __out_len);
void* QRegularExpression_match(void* self, const char* subject, long long subject_len);
void* QRegularExpression_match2(void* self, void* subjectView);
void* QRegularExpression_matchView(void* self, void* subjectView);
void* QRegularExpression_globalMatch(void* self, const char* subject, long long subject_len);
void* QRegularExpression_globalMatch2(void* self, void* subjectView);
void* QRegularExpression_globalMatchView(void* self, void* subjectView);
void QRegularExpression_optimize(void* self);
char* QRegularExpression_escape(const char* str, long long str_len, long long* __out_len);
char* QRegularExpression_wildcardToRegularExpression(const char* str, long long str_len, long long* __out_len);
char* QRegularExpression_anchoredPattern(const char* expression, long long expression_len, long long* __out_len);
char* QRegularExpression_escape2(void* str, long long* __out_len);
char* QRegularExpression_wildcardToRegularExpression2(void* str, long long* __out_len);
char* QRegularExpression_anchoredPattern2(void* expression, long long* __out_len);
void* QRegularExpression_fromWildcard(void* pattern);
void* QRegularExpression_match3(void* self, const char* subject, long long subject_len, long long offset);
void* QRegularExpression_match4(void* self, const char* subject, long long subject_len, long long offset, int matchType);
void* QRegularExpression_match5(void* self, const char* subject, long long subject_len, long long offset, int matchType, int matchOptions);
void* QRegularExpression_match6(void* self, void* subjectView, long long offset);
void* QRegularExpression_match7(void* self, void* subjectView, long long offset, int matchType);
void* QRegularExpression_match8(void* self, void* subjectView, long long offset, int matchType, int matchOptions);
void* QRegularExpression_matchView2(void* self, void* subjectView, long long offset);
void* QRegularExpression_matchView3(void* self, void* subjectView, long long offset, int matchType);
void* QRegularExpression_matchView4(void* self, void* subjectView, long long offset, int matchType, int matchOptions);
void* QRegularExpression_globalMatch3(void* self, const char* subject, long long subject_len, long long offset);
void* QRegularExpression_globalMatch4(void* self, const char* subject, long long subject_len, long long offset, int matchType);
void* QRegularExpression_globalMatch5(void* self, const char* subject, long long subject_len, long long offset, int matchType, int matchOptions);
void* QRegularExpression_globalMatch6(void* self, void* subjectView, long long offset);
void* QRegularExpression_globalMatch7(void* self, void* subjectView, long long offset, int matchType);
void* QRegularExpression_globalMatch8(void* self, void* subjectView, long long offset, int matchType, int matchOptions);
void* QRegularExpression_globalMatchView2(void* self, void* subjectView, long long offset);
void* QRegularExpression_globalMatchView3(void* self, void* subjectView, long long offset, int matchType);
void* QRegularExpression_globalMatchView4(void* self, void* subjectView, long long offset, int matchType, int matchOptions);
char* QRegularExpression_wildcardToRegularExpression3(const char* str, long long str_len, int options, long long* __out_len);
char* QRegularExpression_wildcardToRegularExpression4(void* str, int options, long long* __out_len);
void* QRegularExpression_fromWildcard2(void* pattern, int cs);
void* QRegularExpression_fromWildcard3(void* pattern, int cs, int options);
int QRegularExpression_OperatorEqual(void* self, void* rhs);
int QRegularExpression_OperatorNotEqual(void* self, void* rhs);
void QRegularExpression_delete(void* self);

#ifdef __cplusplus
}
#endif
