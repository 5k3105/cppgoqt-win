#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QTextStream_base(void* self);
void* QTextStream_new();
void* QTextStream_new2(void* device);
void* QTextStream_new3(void* string_);
void* QTextStream_new4(void* array);
void* QTextStream_new5(void* array);
void* QTextStream_new6(void* string_, int openMode);
void* QTextStream_new7(void* array, int openMode);
void* QTextStream_new8(void* array, int openMode);
void QTextStream_setEncoding(void* self, int encoding);
int QTextStream_encoding(void* self);
void QTextStream_setAutoDetectUnicode(void* self, int enabled);
int QTextStream_autoDetectUnicode(void* self);
void QTextStream_setGenerateByteOrderMark(void* self, int generate);
int QTextStream_generateByteOrderMark(void* self);
void QTextStream_setLocale(void* self, void* locale);
void* QTextStream_locale(void* self);
void QTextStream_setDevice(void* self, void* device);
void* QTextStream_device(void* self);
void QTextStream_setString(void* self, void* string_);
void* QTextStream_string(void* self);
int QTextStream_status(void* self);
void QTextStream_setStatus(void* self, int status);
void QTextStream_resetStatus(void* self);
int QTextStream_atEnd(void* self);
void QTextStream_reset(void* self);
void QTextStream_flush(void* self);
int QTextStream_seek(void* self, long long pos);
long long QTextStream_pos(void* self);
void QTextStream_skipWhiteSpace(void* self);
char* QTextStream_readLine(void* self, long long* __out_len);
int QTextStream_readLineInto(void* self, void* line);
char* QTextStream_readAll(void* self, long long* __out_len);
char* QTextStream_read(void* self, long long maxlen, long long* __out_len);
void QTextStream_setFieldAlignment(void* self, int alignment);
int QTextStream_fieldAlignment(void* self);
void QTextStream_setPadChar(void* self, void* ch);
void* QTextStream_padChar(void* self);
void QTextStream_setFieldWidth(void* self, int width);
int QTextStream_fieldWidth(void* self);
void QTextStream_setNumberFlags(void* self, int flags);
int QTextStream_numberFlags(void* self);
void QTextStream_setIntegerBase(void* self, int base);
int QTextStream_integerBase(void* self);
void QTextStream_setRealNumberNotation(void* self, int notation);
int QTextStream_realNumberNotation(void* self);
void QTextStream_setRealNumberPrecision(void* self, int precision);
int QTextStream_realNumberPrecision(void* self);
void* QTextStream_OperatorShiftRight(void* self, void* ch);
void* QTextStream_OperatorShiftRight2(void* self, char* ch);
void* QTextStream_OperatorShiftRight3(void* self, unsigned short* ch);
void* QTextStream_OperatorShiftRight4(void* self, short* i);
void* QTextStream_OperatorShiftRight5(void* self, unsigned short* i);
void* QTextStream_OperatorShiftRight6(void* self, int* i);
void* QTextStream_OperatorShiftRight7(void* self, unsigned int* i);
void* QTextStream_OperatorShiftRight8(void* self, long* i);
void* QTextStream_OperatorShiftRight9(void* self, unsigned long* i);
void* QTextStream_OperatorShiftRight10(void* self, long long* i);
void* QTextStream_OperatorShiftRight11(void* self, unsigned long long* i);
void* QTextStream_OperatorShiftRight12(void* self, float* f);
void* QTextStream_OperatorShiftRight13(void* self, double* f);
void* QTextStream_OperatorShiftRight14(void* self, void* s);
void* QTextStream_OperatorShiftRight15(void* self, void* array);
void* QTextStream_OperatorShiftRight16(void* self, char* c);
void* QTextStream_OperatorShiftLeft(void* self, void* ch);
void* QTextStream_OperatorShiftLeft2(void* self, char ch);
void* QTextStream_OperatorShiftLeft3(void* self, unsigned short ch);
void* QTextStream_OperatorShiftLeft4(void* self, short i);
void* QTextStream_OperatorShiftLeft5(void* self, unsigned short i);
void* QTextStream_OperatorShiftLeft6(void* self, int i);
void* QTextStream_OperatorShiftLeft7(void* self, unsigned int i);
void* QTextStream_OperatorShiftLeft8(void* self, long i);
void* QTextStream_OperatorShiftLeft9(void* self, unsigned long i);
void* QTextStream_OperatorShiftLeft10(void* self, long long i);
void* QTextStream_OperatorShiftLeft11(void* self, unsigned long long i);
void* QTextStream_OperatorShiftLeft12(void* self, float f);
void* QTextStream_OperatorShiftLeft13(void* self, double f);
void* QTextStream_OperatorShiftLeft14(void* self, const char* s, long long s_len);
void* QTextStream_OperatorShiftLeft15(void* self, void* s);
void* QTextStream_OperatorShiftLeft16(void* self, void* s);
void* QTextStream_OperatorShiftLeft17(void* self, void* array);
void* QTextStream_OperatorShiftLeft18(void* self, const char* c);
int QTextStream_ToBool(void* self);
void QTextStream_setString2(void* self, void* string_, int openMode);
char* QTextStream_readLine2(void* self, long long maxlen, long long* __out_len);
int QTextStream_readLineInto2(void* self, void* line, long long maxlen);
void* QTextStream_OperatorShiftLeft19(void* self, void* f);
void* QTextStream_OperatorShiftRight17(void* self, void* f);
void* QTextStream_OperatorShiftLeft20(void* self, void* m);
void QTextStream_delete(void* self);

#ifdef __cplusplus
}
#endif
