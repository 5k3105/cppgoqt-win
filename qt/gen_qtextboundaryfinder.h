#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QTextBoundaryFinder_new();
void* QTextBoundaryFinder_new2(void* other);
void* QTextBoundaryFinder_new3(int type_, const char* string_, long long string__len);
void* QTextBoundaryFinder_new4(int type_, void* chars, long long length);
void* QTextBoundaryFinder_new5(int type_, void* str);
void* QTextBoundaryFinder_new6(int type_, void* chars, long long length, char* buffer);
void* QTextBoundaryFinder_new7(int type_, void* chars, long long length, char* buffer, long long bufferSize);
void* QTextBoundaryFinder_new8(int type_, void* str, char* buffer);
void* QTextBoundaryFinder_new9(int type_, void* str, char* buffer, long long bufferSize);
void* QTextBoundaryFinder_OperatorAssign(void* self, void* other);
void QTextBoundaryFinder_swap(void* self, void* other);
int QTextBoundaryFinder_isValid(void* self);
int QTextBoundaryFinder_type(void* self);
char* QTextBoundaryFinder_string(void* self, long long* __out_len);
void QTextBoundaryFinder_toStart(void* self);
void QTextBoundaryFinder_toEnd(void* self);
long long QTextBoundaryFinder_position(void* self);
void QTextBoundaryFinder_setPosition(void* self, long long position);
long long QTextBoundaryFinder_toNextBoundary(void* self);
long long QTextBoundaryFinder_toPreviousBoundary(void* self);
int QTextBoundaryFinder_isAtBoundary(void* self);
int QTextBoundaryFinder_boundaryReasons(void* self);
void QTextBoundaryFinder_delete(void* self);

#ifdef __cplusplus
}
#endif
