#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QTextOption_new();
void* QTextOption_new2(int alignment);
void* QTextOption_new3(void* o);
void* QTextOption_OperatorAssign(void* self, void* o);
void QTextOption_setAlignment(void* self, int alignment);
int QTextOption_alignment(void* self);
void QTextOption_setTextDirection(void* self, int aDirection);
int QTextOption_textDirection(void* self);
void QTextOption_setWrapMode(void* self, int wrap);
int QTextOption_wrapMode(void* self);
void QTextOption_setFlags(void* self, int flags);
int QTextOption_flags(void* self);
void QTextOption_setTabStopDistance(void* self, double tabStopDistance);
double QTextOption_tabStopDistance(void* self);
void QTextOption_setTabArray(void* self, double* tabStops, long long tabStops_len);
double* QTextOption_tabArray(void* self, long long* __out_len);
void QTextOption_setTabs(void* self, void** tabStops, long long tabStops_len);
void** QTextOption_tabs(void* self, long long* __out_len);
void QTextOption_setUseDesignMetrics(void* self, int b);
int QTextOption_useDesignMetrics(void* self);
void QTextOption_delete(void* self);

#ifdef __cplusplus
}
#endif
