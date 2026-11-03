#include <QtCore/qbytearray.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qbytearrayview.h>
#include <QtCore/qchar.h>
#include <QtCore/qstringview.h>
#include <QtCore/qlatin1stringview.h>
#include <QtCore/qstring.h>
#include "gen_qbytearray.h"

extern "C" {

void* QByteArray_new() {
	return new QByteArray();
}
void* QByteArray_new2(const char* arg1) {
	return new QByteArray(arg1);
}
void* QByteArray_new3(long long size, char c) {
	return new QByteArray(size, c);
}
void* QByteArray_new4(long long size, int arg2) {
	return new QByteArray(size, (Qt::Initialization)(arg2));
}
void* QByteArray_new5(void* v) {
	return new QByteArray(*(QByteArrayView*)v);
}
void* QByteArray_new6(void* arg1) {
	return new QByteArray(*(QByteArray*)arg1);
}
void* QByteArray_new7(const char* arg1, long long size) {
	return new QByteArray(arg1, size);
}
void* QByteArray_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QByteArray*)self)->operator=(*(QByteArray*)arg1);
	return (void*)&(__r);
}
void* QByteArray_OperatorAssign2(void* self, const char* str) {
	auto&& __r = ((QByteArray*)self)->operator=(str);
	return (void*)&(__r);
}
void QByteArray_swap(void* self, void* other) {
	((QByteArray*)self)->swap(*(QByteArray*)other);
}
int QByteArray_isEmpty(void* self) {
	auto&& __r = ((QByteArray*)self)->isEmpty();
	return __r;
}
void QByteArray_resize(void* self, long long size) {
	((QByteArray*)self)->resize(size);
}
void QByteArray_resize2(void* self, long long size, char c) {
	((QByteArray*)self)->resize(size, c);
}
void QByteArray_resizeForOverwrite(void* self, long long size) {
	((QByteArray*)self)->resizeForOverwrite(size);
}
void* QByteArray_fill(void* self, char c) {
	auto&& __r = ((QByteArray*)self)->fill(c);
	return (void*)&(__r);
}
long long QByteArray_capacity(void* self) {
	auto&& __r = ((QByteArray*)self)->capacity();
	return __r;
}
void QByteArray_reserve(void* self, long long size) {
	((QByteArray*)self)->reserve(size);
}
void QByteArray_squeeze(void* self) {
	((QByteArray*)self)->squeeze();
}
const char* QByteArray_data(void* self) {
	auto&& __r = ((QByteArray*)self)->data();
	return __r;
}
const char* QByteArray_constData(void* self) {
	auto&& __r = ((QByteArray*)self)->constData();
	return __r;
}
void QByteArray_detach(void* self) {
	((QByteArray*)self)->detach();
}
int QByteArray_isDetached(void* self) {
	auto&& __r = ((QByteArray*)self)->isDetached();
	return __r;
}
int QByteArray_isSharedWith(void* self, void* other) {
	auto&& __r = ((QByteArray*)self)->isSharedWith(*(QByteArray*)other);
	return __r;
}
void QByteArray_clear(void* self) {
	((QByteArray*)self)->clear();
}
char QByteArray_at(void* self, long long i) {
	auto&& __r = ((QByteArray*)self)->at(i);
	return __r;
}
char QByteArray_OperatorSubscript(void* self, long long i) {
	auto&& __r = ((QByteArray*)self)->operator[](i);
	return __r;
}
char QByteArray_front(void* self) {
	auto&& __r = ((QByteArray*)self)->front();
	return __r;
}
char QByteArray_back(void* self) {
	auto&& __r = ((QByteArray*)self)->back();
	return __r;
}
long long QByteArray_indexOf(void* self, char c) {
	auto&& __r = ((QByteArray*)self)->indexOf(c);
	return __r;
}
long long QByteArray_indexOf2(void* self, void* bv) {
	auto&& __r = ((QByteArray*)self)->indexOf(*(QByteArrayView*)bv);
	return __r;
}
long long QByteArray_lastIndexOf(void* self, char c) {
	auto&& __r = ((QByteArray*)self)->lastIndexOf(c);
	return __r;
}
long long QByteArray_lastIndexOf2(void* self, void* bv) {
	auto&& __r = ((QByteArray*)self)->lastIndexOf(*(QByteArrayView*)bv);
	return __r;
}
long long QByteArray_lastIndexOf3(void* self, void* bv, long long from) {
	auto&& __r = ((QByteArray*)self)->lastIndexOf(*(QByteArrayView*)bv, from);
	return __r;
}
int QByteArray_contains(void* self, char c) {
	auto&& __r = ((QByteArray*)self)->contains(c);
	return __r;
}
int QByteArray_contains2(void* self, void* bv) {
	auto&& __r = ((QByteArray*)self)->contains(*(QByteArrayView*)bv);
	return __r;
}
long long QByteArray_count(void* self, char c) {
	auto&& __r = ((QByteArray*)self)->count(c);
	return __r;
}
long long QByteArray_count2(void* self, void* bv) {
	auto&& __r = ((QByteArray*)self)->count(*(QByteArrayView*)bv);
	return __r;
}
int QByteArray_compare(void* self, void* a) {
	auto&& __r = ((QByteArray*)self)->compare(*(QByteArrayView*)a);
	return __r;
}
void* QByteArray_left(void* self, long long n) {
	auto&& __r = ((QByteArray*)self)->left(n);
	return new QByteArray(__r);
}
void* QByteArray_left2(void* self, long long n) {
	auto&& __r = ((QByteArray*)self)->left(n);
	return new QByteArray(__r);
}
void* QByteArray_right(void* self, long long n) {
	auto&& __r = ((QByteArray*)self)->right(n);
	return new QByteArray(__r);
}
void* QByteArray_right2(void* self, long long n) {
	auto&& __r = ((QByteArray*)self)->right(n);
	return new QByteArray(__r);
}
void* QByteArray_mid(void* self, long long index) {
	auto&& __r = ((QByteArray*)self)->mid(index);
	return new QByteArray(__r);
}
void* QByteArray_mid2(void* self, long long index) {
	auto&& __r = ((QByteArray*)self)->mid(index);
	return new QByteArray(__r);
}
void* QByteArray_first(void* self, long long n) {
	auto&& __r = ((QByteArray*)self)->first(n);
	return new QByteArray(__r);
}
void* QByteArray_last(void* self, long long n) {
	auto&& __r = ((QByteArray*)self)->last(n);
	return new QByteArray(__r);
}
void* QByteArray_sliced(void* self, long long pos) {
	auto&& __r = ((QByteArray*)self)->sliced(pos);
	return new QByteArray(__r);
}
void* QByteArray_sliced2(void* self, long long pos, long long n) {
	auto&& __r = ((QByteArray*)self)->sliced(pos, n);
	return new QByteArray(__r);
}
void* QByteArray_chopped(void* self, long long len_) {
	auto&& __r = ((QByteArray*)self)->chopped(len_);
	return new QByteArray(__r);
}
void* QByteArray_first2(void* self, long long n) {
	auto&& __r = ((QByteArray*)self)->first(n);
	return new QByteArray(__r);
}
void* QByteArray_last2(void* self, long long n) {
	auto&& __r = ((QByteArray*)self)->last(n);
	return new QByteArray(__r);
}
void* QByteArray_sliced3(void* self, long long pos) {
	auto&& __r = ((QByteArray*)self)->sliced(pos);
	return new QByteArray(__r);
}
void* QByteArray_sliced4(void* self, long long pos, long long n) {
	auto&& __r = ((QByteArray*)self)->sliced(pos, n);
	return new QByteArray(__r);
}
void* QByteArray_chopped2(void* self, long long len_) {
	auto&& __r = ((QByteArray*)self)->chopped(len_);
	return new QByteArray(__r);
}
int QByteArray_startsWith(void* self, void* bv) {
	auto&& __r = ((QByteArray*)self)->startsWith(*(QByteArrayView*)bv);
	return __r;
}
int QByteArray_startsWith2(void* self, char c) {
	auto&& __r = ((QByteArray*)self)->startsWith(c);
	return __r;
}
int QByteArray_endsWith(void* self, char c) {
	auto&& __r = ((QByteArray*)self)->endsWith(c);
	return __r;
}
int QByteArray_endsWith2(void* self, void* bv) {
	auto&& __r = ((QByteArray*)self)->endsWith(*(QByteArrayView*)bv);
	return __r;
}
int QByteArray_isUpper(void* self) {
	auto&& __r = ((QByteArray*)self)->isUpper();
	return __r;
}
int QByteArray_isLower(void* self) {
	auto&& __r = ((QByteArray*)self)->isLower();
	return __r;
}
int QByteArray_isValidUtf8(void* self) {
	auto&& __r = ((QByteArray*)self)->isValidUtf8();
	return __r;
}
void QByteArray_truncate(void* self, long long pos) {
	((QByteArray*)self)->truncate(pos);
}
void QByteArray_chop(void* self, long long n) {
	((QByteArray*)self)->chop(n);
}
void* QByteArray_slice(void* self, long long pos) {
	auto&& __r = ((QByteArray*)self)->slice(pos);
	return (void*)&(__r);
}
void* QByteArray_slice2(void* self, long long pos, long long n) {
	auto&& __r = ((QByteArray*)self)->slice(pos, n);
	return (void*)&(__r);
}
void* QByteArray_toLower(void* self) {
	auto&& __r = ((QByteArray*)self)->toLower();
	return new QByteArray(__r);
}
void* QByteArray_toLower2(void* self) {
	auto&& __r = ((QByteArray*)self)->toLower();
	return new QByteArray(__r);
}
void* QByteArray_toUpper(void* self) {
	auto&& __r = ((QByteArray*)self)->toUpper();
	return new QByteArray(__r);
}
void* QByteArray_toUpper2(void* self) {
	auto&& __r = ((QByteArray*)self)->toUpper();
	return new QByteArray(__r);
}
void* QByteArray_trimmed(void* self) {
	auto&& __r = ((QByteArray*)self)->trimmed();
	return new QByteArray(__r);
}
void* QByteArray_trimmed2(void* self) {
	auto&& __r = ((QByteArray*)self)->trimmed();
	return new QByteArray(__r);
}
void* QByteArray_simplified(void* self) {
	auto&& __r = ((QByteArray*)self)->simplified();
	return new QByteArray(__r);
}
void* QByteArray_simplified2(void* self) {
	auto&& __r = ((QByteArray*)self)->simplified();
	return new QByteArray(__r);
}
void* QByteArray_leftJustified(void* self, long long width) {
	auto&& __r = ((QByteArray*)self)->leftJustified(width);
	return new QByteArray(__r);
}
void* QByteArray_rightJustified(void* self, long long width) {
	auto&& __r = ((QByteArray*)self)->rightJustified(width);
	return new QByteArray(__r);
}
void* QByteArray_prepend(void* self, char c) {
	auto&& __r = ((QByteArray*)self)->prepend(c);
	return (void*)&(__r);
}
void* QByteArray_prepend2(void* self, long long count, char c) {
	auto&& __r = ((QByteArray*)self)->prepend(count, c);
	return (void*)&(__r);
}
void* QByteArray_prepend3(void* self, const char* s) {
	auto&& __r = ((QByteArray*)self)->prepend(s);
	return (void*)&(__r);
}
void* QByteArray_prepend4(void* self, const char* s, long long len_) {
	auto&& __r = ((QByteArray*)self)->prepend(s, len_);
	return (void*)&(__r);
}
void* QByteArray_prepend5(void* self, void* a) {
	auto&& __r = ((QByteArray*)self)->prepend(*(QByteArray*)a);
	return (void*)&(__r);
}
void* QByteArray_prepend6(void* self, void* a) {
	auto&& __r = ((QByteArray*)self)->prepend(*(QByteArrayView*)a);
	return (void*)&(__r);
}
void* QByteArray_append(void* self, char c) {
	auto&& __r = ((QByteArray*)self)->append(c);
	return (void*)&(__r);
}
void* QByteArray_append2(void* self, long long count, char c) {
	auto&& __r = ((QByteArray*)self)->append(count, c);
	return (void*)&(__r);
}
void* QByteArray_append3(void* self, const char* s) {
	auto&& __r = ((QByteArray*)self)->append(s);
	return (void*)&(__r);
}
void* QByteArray_append4(void* self, const char* s, long long len_) {
	auto&& __r = ((QByteArray*)self)->append(s, len_);
	return (void*)&(__r);
}
void* QByteArray_append5(void* self, void* a) {
	auto&& __r = ((QByteArray*)self)->append(*(QByteArray*)a);
	return (void*)&(__r);
}
void* QByteArray_append6(void* self, void* a) {
	auto&& __r = ((QByteArray*)self)->append(*(QByteArrayView*)a);
	return (void*)&(__r);
}
void* QByteArray_assign(void* self, void* v) {
	auto&& __r = ((QByteArray*)self)->assign(*(QByteArrayView*)v);
	return (void*)&(__r);
}
void* QByteArray_assign2(void* self, long long n, char c) {
	auto&& __r = ((QByteArray*)self)->assign(n, c);
	return (void*)&(__r);
}
void* QByteArray_insert(void* self, long long i, void* data) {
	auto&& __r = ((QByteArray*)self)->insert(i, *(QByteArrayView*)data);
	return (void*)&(__r);
}
void* QByteArray_insert2(void* self, long long i, const char* s) {
	auto&& __r = ((QByteArray*)self)->insert(i, s);
	return (void*)&(__r);
}
void* QByteArray_insert3(void* self, long long i, void* data) {
	auto&& __r = ((QByteArray*)self)->insert(i, *(QByteArray*)data);
	return (void*)&(__r);
}
void* QByteArray_insert4(void* self, long long i, long long count, char c) {
	auto&& __r = ((QByteArray*)self)->insert(i, count, c);
	return (void*)&(__r);
}
void* QByteArray_insert5(void* self, long long i, char c) {
	auto&& __r = ((QByteArray*)self)->insert(i, c);
	return (void*)&(__r);
}
void* QByteArray_insert6(void* self, long long i, const char* s, long long len_) {
	auto&& __r = ((QByteArray*)self)->insert(i, s, len_);
	return (void*)&(__r);
}
void* QByteArray_remove(void* self, long long index, long long len_) {
	auto&& __r = ((QByteArray*)self)->remove(index, len_);
	return (void*)&(__r);
}
void* QByteArray_removeAt(void* self, long long pos) {
	auto&& __r = ((QByteArray*)self)->removeAt(pos);
	return (void*)&(__r);
}
void* QByteArray_removeFirst(void* self) {
	auto&& __r = ((QByteArray*)self)->removeFirst();
	return (void*)&(__r);
}
void* QByteArray_removeLast(void* self) {
	auto&& __r = ((QByteArray*)self)->removeLast();
	return (void*)&(__r);
}
void* QByteArray_replace(void* self, long long index, long long len_, const char* s, long long alen) {
	auto&& __r = ((QByteArray*)self)->replace(index, len_, s, alen);
	return (void*)&(__r);
}
void* QByteArray_replace2(void* self, long long index, long long len_, void* s) {
	auto&& __r = ((QByteArray*)self)->replace(index, len_, *(QByteArrayView*)s);
	return (void*)&(__r);
}
void* QByteArray_replace3(void* self, char before, void* after) {
	auto&& __r = ((QByteArray*)self)->replace(before, *(QByteArrayView*)after);
	return (void*)&(__r);
}
void* QByteArray_replace4(void* self, const char* before, long long bsize, const char* after, long long asize) {
	auto&& __r = ((QByteArray*)self)->replace(before, bsize, after, asize);
	return (void*)&(__r);
}
void* QByteArray_replace5(void* self, void* before, void* after) {
	auto&& __r = ((QByteArray*)self)->replace(*(QByteArrayView*)before, *(QByteArrayView*)after);
	return (void*)&(__r);
}
void* QByteArray_replace6(void* self, char before, char after) {
	auto&& __r = ((QByteArray*)self)->replace(before, after);
	return (void*)&(__r);
}
void* QByteArray_OperatorPlusAssign(void* self, char c) {
	auto&& __r = ((QByteArray*)self)->operator+=(c);
	return (void*)&(__r);
}
void* QByteArray_OperatorPlusAssign2(void* self, const char* s) {
	auto&& __r = ((QByteArray*)self)->operator+=(s);
	return (void*)&(__r);
}
void* QByteArray_OperatorPlusAssign3(void* self, void* a) {
	auto&& __r = ((QByteArray*)self)->operator+=(*(QByteArray*)a);
	return (void*)&(__r);
}
void* QByteArray_OperatorPlusAssign4(void* self, void* a) {
	auto&& __r = ((QByteArray*)self)->operator+=(*(QByteArrayView*)a);
	return (void*)&(__r);
}
void** QByteArray_split(void* self, char sep, long long* __out_len) {
	auto&& __r = ((QByteArray*)self)->split(sep);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QByteArray(__e);
	return __arr;
}
void* QByteArray_repeated(void* self, long long times) {
	auto&& __r = ((QByteArray*)self)->repeated(times);
	return new QByteArray(__r);
}
short QByteArray_toShort(void* self) {
	auto&& __r = ((QByteArray*)self)->toShort();
	return __r;
}
unsigned short QByteArray_toUShort(void* self) {
	auto&& __r = ((QByteArray*)self)->toUShort();
	return __r;
}
int QByteArray_toInt(void* self) {
	auto&& __r = ((QByteArray*)self)->toInt();
	return __r;
}
unsigned int QByteArray_toUInt(void* self) {
	auto&& __r = ((QByteArray*)self)->toUInt();
	return __r;
}
long QByteArray_toLong(void* self) {
	auto&& __r = ((QByteArray*)self)->toLong();
	return __r;
}
unsigned long QByteArray_toULong(void* self) {
	auto&& __r = ((QByteArray*)self)->toULong();
	return __r;
}
long long QByteArray_toLongLong(void* self) {
	auto&& __r = ((QByteArray*)self)->toLongLong();
	return __r;
}
unsigned long long QByteArray_toULongLong(void* self) {
	auto&& __r = ((QByteArray*)self)->toULongLong();
	return __r;
}
float QByteArray_toFloat(void* self) {
	auto&& __r = ((QByteArray*)self)->toFloat();
	return __r;
}
double QByteArray_toDouble(void* self) {
	auto&& __r = ((QByteArray*)self)->toDouble();
	return __r;
}
void* QByteArray_toBase64(void* self) {
	auto&& __r = ((QByteArray*)self)->toBase64();
	return new QByteArray(__r);
}
void* QByteArray_toHex(void* self) {
	auto&& __r = ((QByteArray*)self)->toHex();
	return new QByteArray(__r);
}
void* QByteArray_toPercentEncoding(void* self) {
	auto&& __r = ((QByteArray*)self)->toPercentEncoding();
	return new QByteArray(__r);
}
void* QByteArray_percentDecoded(void* self) {
	auto&& __r = ((QByteArray*)self)->percentDecoded();
	return new QByteArray(__r);
}
void* QByteArray_percentDecoded2(void* self) {
	auto&& __r = ((QByteArray*)self)->percentDecoded();
	return new QByteArray(__r);
}
void* QByteArray_setNum(void* self, short arg1) {
	auto&& __r = ((QByteArray*)self)->setNum(arg1);
	return (void*)&(__r);
}
void* QByteArray_setNum2(void* self, unsigned short arg1) {
	auto&& __r = ((QByteArray*)self)->setNum(arg1);
	return (void*)&(__r);
}
void* QByteArray_setNum3(void* self, int arg1) {
	auto&& __r = ((QByteArray*)self)->setNum(arg1);
	return (void*)&(__r);
}
void* QByteArray_setNum4(void* self, unsigned int arg1) {
	auto&& __r = ((QByteArray*)self)->setNum(arg1);
	return (void*)&(__r);
}
void* QByteArray_setNum5(void* self, long arg1) {
	auto&& __r = ((QByteArray*)self)->setNum(arg1);
	return (void*)&(__r);
}
void* QByteArray_setNum6(void* self, unsigned long arg1) {
	auto&& __r = ((QByteArray*)self)->setNum(arg1);
	return (void*)&(__r);
}
void* QByteArray_setNum7(void* self, long long arg1) {
	auto&& __r = ((QByteArray*)self)->setNum(arg1);
	return (void*)&(__r);
}
void* QByteArray_setNum8(void* self, unsigned long long arg1) {
	auto&& __r = ((QByteArray*)self)->setNum(arg1);
	return (void*)&(__r);
}
void* QByteArray_setNum9(void* self, float arg1) {
	auto&& __r = ((QByteArray*)self)->setNum(arg1);
	return (void*)&(__r);
}
void* QByteArray_setNum10(void* self, double arg1) {
	auto&& __r = ((QByteArray*)self)->setNum(arg1);
	return (void*)&(__r);
}
void* QByteArray_setRawData(void* self, const char* a, long long n) {
	auto&& __r = ((QByteArray*)self)->setRawData(a, n);
	return (void*)&(__r);
}
void* QByteArray_number(int arg1) {
	auto&& __r = QByteArray::number(arg1);
	return new QByteArray(__r);
}
void* QByteArray_number2(unsigned int arg1) {
	auto&& __r = QByteArray::number(arg1);
	return new QByteArray(__r);
}
void* QByteArray_number3(long arg1) {
	auto&& __r = QByteArray::number(arg1);
	return new QByteArray(__r);
}
void* QByteArray_number4(unsigned long arg1) {
	auto&& __r = QByteArray::number(arg1);
	return new QByteArray(__r);
}
void* QByteArray_number5(long long arg1) {
	auto&& __r = QByteArray::number(arg1);
	return new QByteArray(__r);
}
void* QByteArray_number6(unsigned long long arg1) {
	auto&& __r = QByteArray::number(arg1);
	return new QByteArray(__r);
}
void* QByteArray_number7(double arg1) {
	auto&& __r = QByteArray::number(arg1);
	return new QByteArray(__r);
}
void* QByteArray_fromRawData(const char* data, long long size) {
	auto&& __r = QByteArray::fromRawData(data, size);
	return new QByteArray(__r);
}
void* QByteArray_fromBase64Encoding(void* base64) {
	auto&& __r = QByteArray::fromBase64Encoding(*(QByteArray*)base64);
	return new QByteArray::FromBase64Result(__r);
}
void* QByteArray_fromBase64(void* base64) {
	auto&& __r = QByteArray::fromBase64(*(QByteArray*)base64);
	return new QByteArray(__r);
}
void* QByteArray_fromHex(void* hexEncoded) {
	auto&& __r = QByteArray::fromHex(*(QByteArray*)hexEncoded);
	return new QByteArray(__r);
}
void* QByteArray_fromPercentEncoding(void* pctEncoded) {
	auto&& __r = QByteArray::fromPercentEncoding(*(QByteArray*)pctEncoded);
	return new QByteArray(__r);
}
const char* QByteArray_begin(void* self) {
	auto&& __r = ((QByteArray*)self)->begin();
	return __r;
}
const char* QByteArray_cbegin(void* self) {
	auto&& __r = ((QByteArray*)self)->cbegin();
	return __r;
}
const char* QByteArray_constBegin(void* self) {
	auto&& __r = ((QByteArray*)self)->constBegin();
	return __r;
}
const char* QByteArray_end(void* self) {
	auto&& __r = ((QByteArray*)self)->end();
	return __r;
}
const char* QByteArray_cend(void* self) {
	auto&& __r = ((QByteArray*)self)->cend();
	return __r;
}
const char* QByteArray_constEnd(void* self) {
	auto&& __r = ((QByteArray*)self)->constEnd();
	return __r;
}
void QByteArray_push_back(void* self, char c) {
	((QByteArray*)self)->push_back(c);
}
void QByteArray_push_back2(void* self, const char* s) {
	((QByteArray*)self)->push_back(s);
}
void QByteArray_push_back3(void* self, void* a) {
	((QByteArray*)self)->push_back(*(QByteArray*)a);
}
void QByteArray_push_back4(void* self, void* a) {
	((QByteArray*)self)->push_back(*(QByteArrayView*)a);
}
void QByteArray_push_front(void* self, char c) {
	((QByteArray*)self)->push_front(c);
}
void QByteArray_push_front2(void* self, const char* c) {
	((QByteArray*)self)->push_front(c);
}
void QByteArray_push_front3(void* self, void* a) {
	((QByteArray*)self)->push_front(*(QByteArray*)a);
}
void QByteArray_push_front4(void* self, void* a) {
	((QByteArray*)self)->push_front(*(QByteArrayView*)a);
}
void QByteArray_shrink_to_fit(void* self) {
	((QByteArray*)self)->shrink_to_fit();
}
long long QByteArray_max_size(void* self) {
	auto&& __r = ((QByteArray*)self)->max_size();
	return __r;
}
long long QByteArray_maxSize() {
	auto&& __r = QByteArray::maxSize();
	return __r;
}
long long QByteArray_size(void* self) {
	auto&& __r = ((QByteArray*)self)->size();
	return __r;
}
long long QByteArray_count3(void* self) {
	auto&& __r = ((QByteArray*)self)->count();
	return __r;
}
long long QByteArray_length(void* self) {
	auto&& __r = ((QByteArray*)self)->length();
	return __r;
}
int QByteArray_isNull(void* self) {
	auto&& __r = ((QByteArray*)self)->isNull();
	return __r;
}
void* QByteArray_nullTerminated(void* self) {
	auto&& __r = ((QByteArray*)self)->nullTerminated();
	return new QByteArray(__r);
}
void* QByteArray_nullTerminated2(void* self) {
	auto&& __r = ((QByteArray*)self)->nullTerminated();
	return new QByteArray(__r);
}
void* QByteArray_nullTerminate(void* self) {
	auto&& __r = ((QByteArray*)self)->nullTerminate();
	return (void*)&(__r);
}
void* QByteArray_fill2(void* self, char c, long long size) {
	auto&& __r = ((QByteArray*)self)->fill(c, size);
	return (void*)&(__r);
}
long long QByteArray_indexOf3(void* self, char c, long long from) {
	auto&& __r = ((QByteArray*)self)->indexOf(c, from);
	return __r;
}
long long QByteArray_indexOf4(void* self, void* bv, long long from) {
	auto&& __r = ((QByteArray*)self)->indexOf(*(QByteArrayView*)bv, from);
	return __r;
}
long long QByteArray_lastIndexOf4(void* self, char c, long long from) {
	auto&& __r = ((QByteArray*)self)->lastIndexOf(c, from);
	return __r;
}
int QByteArray_compare2(void* self, void* a, int cs) {
	auto&& __r = ((QByteArray*)self)->compare(*(QByteArrayView*)a, (Qt::CaseSensitivity)(cs));
	return __r;
}
void* QByteArray_mid3(void* self, long long index, long long len_) {
	auto&& __r = ((QByteArray*)self)->mid(index, len_);
	return new QByteArray(__r);
}
void* QByteArray_mid4(void* self, long long index, long long len_) {
	auto&& __r = ((QByteArray*)self)->mid(index, len_);
	return new QByteArray(__r);
}
void* QByteArray_leftJustified2(void* self, long long width, char fill) {
	auto&& __r = ((QByteArray*)self)->leftJustified(width, fill);
	return new QByteArray(__r);
}
void* QByteArray_leftJustified3(void* self, long long width, char fill, int truncate) {
	auto&& __r = ((QByteArray*)self)->leftJustified(width, fill, truncate);
	return new QByteArray(__r);
}
void* QByteArray_rightJustified2(void* self, long long width, char fill) {
	auto&& __r = ((QByteArray*)self)->rightJustified(width, fill);
	return new QByteArray(__r);
}
void* QByteArray_rightJustified3(void* self, long long width, char fill, int truncate) {
	auto&& __r = ((QByteArray*)self)->rightJustified(width, fill, truncate);
	return new QByteArray(__r);
}
short QByteArray_toShort2(void* self, int* ok) {
	bool ok_local{};
	auto&& __r = ((QByteArray*)self)->toShort(&ok_local);
	*ok = ok_local;
	return __r;
}
short QByteArray_toShort3(void* self, int* ok, int base) {
	bool ok_local{};
	auto&& __r = ((QByteArray*)self)->toShort(&ok_local, base);
	*ok = ok_local;
	return __r;
}
unsigned short QByteArray_toUShort2(void* self, int* ok) {
	bool ok_local{};
	auto&& __r = ((QByteArray*)self)->toUShort(&ok_local);
	*ok = ok_local;
	return __r;
}
unsigned short QByteArray_toUShort3(void* self, int* ok, int base) {
	bool ok_local{};
	auto&& __r = ((QByteArray*)self)->toUShort(&ok_local, base);
	*ok = ok_local;
	return __r;
}
int QByteArray_toInt2(void* self, int* ok) {
	bool ok_local{};
	auto&& __r = ((QByteArray*)self)->toInt(&ok_local);
	*ok = ok_local;
	return __r;
}
int QByteArray_toInt3(void* self, int* ok, int base) {
	bool ok_local{};
	auto&& __r = ((QByteArray*)self)->toInt(&ok_local, base);
	*ok = ok_local;
	return __r;
}
unsigned int QByteArray_toUInt2(void* self, int* ok) {
	bool ok_local{};
	auto&& __r = ((QByteArray*)self)->toUInt(&ok_local);
	*ok = ok_local;
	return __r;
}
unsigned int QByteArray_toUInt3(void* self, int* ok, int base) {
	bool ok_local{};
	auto&& __r = ((QByteArray*)self)->toUInt(&ok_local, base);
	*ok = ok_local;
	return __r;
}
long QByteArray_toLong2(void* self, int* ok) {
	bool ok_local{};
	auto&& __r = ((QByteArray*)self)->toLong(&ok_local);
	*ok = ok_local;
	return __r;
}
long QByteArray_toLong3(void* self, int* ok, int base) {
	bool ok_local{};
	auto&& __r = ((QByteArray*)self)->toLong(&ok_local, base);
	*ok = ok_local;
	return __r;
}
unsigned long QByteArray_toULong2(void* self, int* ok) {
	bool ok_local{};
	auto&& __r = ((QByteArray*)self)->toULong(&ok_local);
	*ok = ok_local;
	return __r;
}
unsigned long QByteArray_toULong3(void* self, int* ok, int base) {
	bool ok_local{};
	auto&& __r = ((QByteArray*)self)->toULong(&ok_local, base);
	*ok = ok_local;
	return __r;
}
long long QByteArray_toLongLong2(void* self, int* ok) {
	bool ok_local{};
	auto&& __r = ((QByteArray*)self)->toLongLong(&ok_local);
	*ok = ok_local;
	return __r;
}
long long QByteArray_toLongLong3(void* self, int* ok, int base) {
	bool ok_local{};
	auto&& __r = ((QByteArray*)self)->toLongLong(&ok_local, base);
	*ok = ok_local;
	return __r;
}
unsigned long long QByteArray_toULongLong2(void* self, int* ok) {
	bool ok_local{};
	auto&& __r = ((QByteArray*)self)->toULongLong(&ok_local);
	*ok = ok_local;
	return __r;
}
unsigned long long QByteArray_toULongLong3(void* self, int* ok, int base) {
	bool ok_local{};
	auto&& __r = ((QByteArray*)self)->toULongLong(&ok_local, base);
	*ok = ok_local;
	return __r;
}
float QByteArray_toFloat2(void* self, int* ok) {
	bool ok_local{};
	auto&& __r = ((QByteArray*)self)->toFloat(&ok_local);
	*ok = ok_local;
	return __r;
}
double QByteArray_toDouble2(void* self, int* ok) {
	bool ok_local{};
	auto&& __r = ((QByteArray*)self)->toDouble(&ok_local);
	*ok = ok_local;
	return __r;
}
void* QByteArray_toBase642(void* self, int options) {
	auto&& __r = ((QByteArray*)self)->toBase64(QFlags<QByteArray::Base64Option>(options));
	return new QByteArray(__r);
}
void* QByteArray_toHex2(void* self, char separator) {
	auto&& __r = ((QByteArray*)self)->toHex(separator);
	return new QByteArray(__r);
}
void* QByteArray_toPercentEncoding2(void* self, void* exclude) {
	auto&& __r = ((QByteArray*)self)->toPercentEncoding(*(QByteArray*)exclude);
	return new QByteArray(__r);
}
void* QByteArray_toPercentEncoding3(void* self, void* exclude, void* include) {
	auto&& __r = ((QByteArray*)self)->toPercentEncoding(*(QByteArray*)exclude, *(QByteArray*)include);
	return new QByteArray(__r);
}
void* QByteArray_toPercentEncoding4(void* self, void* exclude, void* include, char percent) {
	auto&& __r = ((QByteArray*)self)->toPercentEncoding(*(QByteArray*)exclude, *(QByteArray*)include, percent);
	return new QByteArray(__r);
}
void* QByteArray_percentDecoded3(void* self, char percent) {
	auto&& __r = ((QByteArray*)self)->percentDecoded(percent);
	return new QByteArray(__r);
}
void* QByteArray_percentDecoded4(void* self, char percent) {
	auto&& __r = ((QByteArray*)self)->percentDecoded(percent);
	return new QByteArray(__r);
}
void* QByteArray_setNum11(void* self, short arg1, int base) {
	auto&& __r = ((QByteArray*)self)->setNum(arg1, base);
	return (void*)&(__r);
}
void* QByteArray_setNum12(void* self, unsigned short arg1, int base) {
	auto&& __r = ((QByteArray*)self)->setNum(arg1, base);
	return (void*)&(__r);
}
void* QByteArray_setNum13(void* self, int arg1, int base) {
	auto&& __r = ((QByteArray*)self)->setNum(arg1, base);
	return (void*)&(__r);
}
void* QByteArray_setNum14(void* self, unsigned int arg1, int base) {
	auto&& __r = ((QByteArray*)self)->setNum(arg1, base);
	return (void*)&(__r);
}
void* QByteArray_setNum15(void* self, long arg1, int base) {
	auto&& __r = ((QByteArray*)self)->setNum(arg1, base);
	return (void*)&(__r);
}
void* QByteArray_setNum16(void* self, unsigned long arg1, int base) {
	auto&& __r = ((QByteArray*)self)->setNum(arg1, base);
	return (void*)&(__r);
}
void* QByteArray_setNum17(void* self, long long arg1, int base) {
	auto&& __r = ((QByteArray*)self)->setNum(arg1, base);
	return (void*)&(__r);
}
void* QByteArray_setNum18(void* self, unsigned long long arg1, int base) {
	auto&& __r = ((QByteArray*)self)->setNum(arg1, base);
	return (void*)&(__r);
}
void* QByteArray_setNum19(void* self, float arg1, char format) {
	auto&& __r = ((QByteArray*)self)->setNum(arg1, format);
	return (void*)&(__r);
}
void* QByteArray_setNum20(void* self, float arg1, char format, int precision) {
	auto&& __r = ((QByteArray*)self)->setNum(arg1, format, precision);
	return (void*)&(__r);
}
void* QByteArray_setNum21(void* self, double arg1, char format) {
	auto&& __r = ((QByteArray*)self)->setNum(arg1, format);
	return (void*)&(__r);
}
void* QByteArray_setNum22(void* self, double arg1, char format, int precision) {
	auto&& __r = ((QByteArray*)self)->setNum(arg1, format, precision);
	return (void*)&(__r);
}
void* QByteArray_number8(int arg1, int base) {
	auto&& __r = QByteArray::number(arg1, base);
	return new QByteArray(__r);
}
void* QByteArray_number9(unsigned int arg1, int base) {
	auto&& __r = QByteArray::number(arg1, base);
	return new QByteArray(__r);
}
void* QByteArray_number10(long arg1, int base) {
	auto&& __r = QByteArray::number(arg1, base);
	return new QByteArray(__r);
}
void* QByteArray_number11(unsigned long arg1, int base) {
	auto&& __r = QByteArray::number(arg1, base);
	return new QByteArray(__r);
}
void* QByteArray_number12(long long arg1, int base) {
	auto&& __r = QByteArray::number(arg1, base);
	return new QByteArray(__r);
}
void* QByteArray_number13(unsigned long long arg1, int base) {
	auto&& __r = QByteArray::number(arg1, base);
	return new QByteArray(__r);
}
void* QByteArray_number14(double arg1, char format) {
	auto&& __r = QByteArray::number(arg1, format);
	return new QByteArray(__r);
}
void* QByteArray_number15(double arg1, char format, int precision) {
	auto&& __r = QByteArray::number(arg1, format, precision);
	return new QByteArray(__r);
}
void* QByteArray_fromBase64Encoding2(void* base64, int options) {
	auto&& __r = QByteArray::fromBase64Encoding(*(QByteArray*)base64, QFlags<QByteArray::Base64Option>(options));
	return new QByteArray::FromBase64Result(__r);
}
void* QByteArray_fromBase642(void* base64, int options) {
	auto&& __r = QByteArray::fromBase64(*(QByteArray*)base64, QFlags<QByteArray::Base64Option>(options));
	return new QByteArray(__r);
}
void* QByteArray_fromPercentEncoding2(void* pctEncoded, char percent) {
	auto&& __r = QByteArray::fromPercentEncoding(*(QByteArray*)pctEncoded, percent);
	return new QByteArray(__r);
}
int QByteArray_OperatorEqual(void* self, void* rhs) {
	auto&& __r = (*(QByteArray*)self) == (*(QByteArray*)rhs);
	return __r;
}
int QByteArray_OperatorNotEqual(void* self, void* rhs) {
	auto&& __r = (*(QByteArray*)self) != (*(QByteArray*)rhs);
	return __r;
}
int QByteArray_OperatorLesser(void* self, void* rhs) {
	auto&& __r = (*(QByteArray*)self) < (*(QByteArray*)rhs);
	return __r;
}
int QByteArray_OperatorGreater(void* self, void* rhs) {
	auto&& __r = (*(QByteArray*)self) > (*(QByteArray*)rhs);
	return __r;
}
int QByteArray_OperatorLesserOrEqual(void* self, void* rhs) {
	auto&& __r = (*(QByteArray*)self) <= (*(QByteArray*)rhs);
	return __r;
}
int QByteArray_OperatorGreaterOrEqual(void* self, void* rhs) {
	auto&& __r = (*(QByteArray*)self) >= (*(QByteArray*)rhs);
	return __r;
}
int QByteArray_OperatorEqual2(void* self, void* rhs) {
	auto&& __r = (*(QByteArray*)self) == (*(QChar*)rhs);
	return __r;
}
int QByteArray_OperatorNotEqual2(void* self, void* rhs) {
	auto&& __r = (*(QByteArray*)self) != (*(QChar*)rhs);
	return __r;
}
int QByteArray_OperatorLesser2(void* self, void* rhs) {
	auto&& __r = (*(QByteArray*)self) < (*(QChar*)rhs);
	return __r;
}
int QByteArray_OperatorGreater2(void* self, void* rhs) {
	auto&& __r = (*(QByteArray*)self) > (*(QChar*)rhs);
	return __r;
}
int QByteArray_OperatorLesserOrEqual2(void* self, void* rhs) {
	auto&& __r = (*(QByteArray*)self) <= (*(QChar*)rhs);
	return __r;
}
int QByteArray_OperatorGreaterOrEqual2(void* self, void* rhs) {
	auto&& __r = (*(QByteArray*)self) >= (*(QChar*)rhs);
	return __r;
}
void* QByteArray_OperatorPlus(void* self, void* a2) {
	auto&& __r = (*(QByteArray*)self) + (*(QByteArray*)a2);
	return new QByteArray(__r);
}
void* QByteArray_OperatorPlus2(void* self, const char* a2) {
	auto&& __r = (*(QByteArray*)self) + (a2);
	return new QByteArray(__r);
}
void* QByteArray_OperatorPlus3(void* self, char a2) {
	auto&& __r = (*(QByteArray*)self) + (a2);
	return new QByteArray(__r);
}
int QByteArray_OperatorEqual3(void* self, void* rhs) {
	auto&& __r = (*(QByteArray*)self) == (*(QStringView*)rhs);
	return __r;
}
int QByteArray_OperatorNotEqual3(void* self, void* rhs) {
	auto&& __r = (*(QByteArray*)self) != (*(QStringView*)rhs);
	return __r;
}
int QByteArray_OperatorLesser3(void* self, void* rhs) {
	auto&& __r = (*(QByteArray*)self) < (*(QStringView*)rhs);
	return __r;
}
int QByteArray_OperatorGreater3(void* self, void* rhs) {
	auto&& __r = (*(QByteArray*)self) > (*(QStringView*)rhs);
	return __r;
}
int QByteArray_OperatorLesserOrEqual3(void* self, void* rhs) {
	auto&& __r = (*(QByteArray*)self) <= (*(QStringView*)rhs);
	return __r;
}
int QByteArray_OperatorGreaterOrEqual3(void* self, void* rhs) {
	auto&& __r = (*(QByteArray*)self) >= (*(QStringView*)rhs);
	return __r;
}
int QByteArray_OperatorEqual4(void* self, const char* rhs, long long rhs_len) {
	QString rhs_qstr = QString::fromUtf8(rhs, rhs_len);
	auto&& __r = (*(QByteArray*)self) == (rhs_qstr);
	return __r;
}
int QByteArray_OperatorNotEqual4(void* self, const char* rhs, long long rhs_len) {
	QString rhs_qstr = QString::fromUtf8(rhs, rhs_len);
	auto&& __r = (*(QByteArray*)self) != (rhs_qstr);
	return __r;
}
int QByteArray_OperatorLesser4(void* self, const char* rhs, long long rhs_len) {
	QString rhs_qstr = QString::fromUtf8(rhs, rhs_len);
	auto&& __r = (*(QByteArray*)self) < (rhs_qstr);
	return __r;
}
int QByteArray_OperatorGreater4(void* self, const char* rhs, long long rhs_len) {
	QString rhs_qstr = QString::fromUtf8(rhs, rhs_len);
	auto&& __r = (*(QByteArray*)self) > (rhs_qstr);
	return __r;
}
int QByteArray_OperatorLesserOrEqual4(void* self, const char* rhs, long long rhs_len) {
	QString rhs_qstr = QString::fromUtf8(rhs, rhs_len);
	auto&& __r = (*(QByteArray*)self) <= (rhs_qstr);
	return __r;
}
int QByteArray_OperatorGreaterOrEqual4(void* self, const char* rhs, long long rhs_len) {
	QString rhs_qstr = QString::fromUtf8(rhs, rhs_len);
	auto&& __r = (*(QByteArray*)self) >= (rhs_qstr);
	return __r;
}
char* QByteArray_OperatorPlus4(void* self, const char* s, long long s_len, long long* __out_len) {
	QString s_qstr = QString::fromUtf8(s, s_len);
	auto&& __r = (*(QByteArray*)self) + (s_qstr);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QByteArray_delete(void* self) { delete (QByteArray*)self; }

}

#include <QtCore/qobject.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qvariant.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qobject.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QObject__Go : public QObject {
public:
	long long __goId;
	unsigned char __goOn[8] = {};
	template<typename... A> QObject__Go(long long id, A&&... a) : QObject(std::forward<A>(a)...), __goId(id) {}
	using QObject::sender;
	using QObject::senderSignalIndex;
	using QObject::receivers;
	using QObject::isSignalConnected;
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QObject::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QObject::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QObject::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QObject::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QObject::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QObject::metaObject();
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QObject::event(event);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QObject::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QObject::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QObject::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QObject::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QObject::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QObject::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QObject;
static std::mutex __connsMu_QObject;

extern "C" {

void QObject__set_override(void* self, int idx) { ((QObject__Go*)self)->__goOn[idx] = 1; }
void* QObject_new(long long __goid) {
	return new QObject__Go(__goid);
}
void* QObject_new2(long long __goid, void* parent) {
	return new QObject__Go(__goid, (QObject*)parent);
}
void* QObject_metaObject(void* self) {
	auto&& __r = ((QObject*)self)->metaObject();
	return (void*)(__r);
}
char* QObject_tr(const char* s, long long* __out_len) {
	auto&& __r = QObject::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QObject_event(void* self, void* event) {
	auto&& __r = ((QObject*)self)->event((QEvent*)event);
	return __r;
}
int QObject_eventFilter(void* self, void* watched, void* event) {
	auto&& __r = ((QObject*)self)->eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
char* QObject_objectName(void* self, long long* __out_len) {
	auto&& __r = ((QObject*)self)->objectName();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QObject_setObjectName(void* self, const char* name, long long name_len) {
	QAnyStringView name_qstr = QAnyStringView(name, name_len);
	((QObject*)self)->setObjectName(name_qstr);
}
int QObject_isWidgetType(void* self) {
	auto&& __r = ((QObject*)self)->isWidgetType();
	return __r;
}
int QObject_isWindowType(void* self) {
	auto&& __r = ((QObject*)self)->isWindowType();
	return __r;
}
int QObject_isQuickItemType(void* self) {
	auto&& __r = ((QObject*)self)->isQuickItemType();
	return __r;
}
int QObject_isQmlExposed(void* self) {
	auto&& __r = ((QObject*)self)->isQmlExposed();
	return __r;
}
int QObject_signalsBlocked(void* self) {
	auto&& __r = ((QObject*)self)->signalsBlocked();
	return __r;
}
int QObject_blockSignals(void* self, int b) {
	auto&& __r = ((QObject*)self)->blockSignals(b);
	return __r;
}
void* QObject_thread(void* self) {
	auto&& __r = ((QObject*)self)->thread();
	return (void*)(__r);
}
int QObject_moveToThread(void* self, void* thread) {
	auto&& __r = ((QObject*)self)->moveToThread((QThread*)thread);
	return __r;
}
int QObject_startTimer(void* self, int interval) {
	auto&& __r = ((QObject*)self)->startTimer(interval);
	return __r;
}
int QObject_startTimer2(void* self, long long time) {
	auto&& __r = ((QObject*)self)->startTimer(std::chrono::nanoseconds((long long)time));
	return __r;
}
void QObject_killTimer(void* self, int id) {
	((QObject*)self)->killTimer(id);
}
void QObject_killTimer2(void* self, int id) {
	((QObject*)self)->killTimer((Qt::TimerId)(id));
}
void** QObject_children(void* self, long long* __out_len) {
	auto&& __r = ((QObject*)self)->children();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
void QObject_setParent(void* self, void* parent) {
	((QObject*)self)->setParent((QObject*)parent);
}
void QObject_installEventFilter(void* self, void* filterObj) {
	((QObject*)self)->installEventFilter((QObject*)filterObj);
}
void QObject_removeEventFilter(void* self, void* obj) {
	((QObject*)self)->removeEventFilter((QObject*)obj);
}
void* QObject_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QObject::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QObject_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QObject::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QObject_connect3(void* self, void* sender, const char* signal, const char* member) {
	auto&& __r = ((QObject*)self)->connect((QObject*)sender, signal, member);
	return new QMetaObject::Connection(__r);
}
void QObject_dumpObjectTree(void* self) {
	((QObject*)self)->dumpObjectTree();
}
void QObject_dumpObjectInfo(void* self) {
	((QObject*)self)->dumpObjectInfo();
}
int QObject_setProperty(void* self, const char* name, void* value) {
	auto&& __r = ((QObject*)self)->setProperty(name, *(QVariant*)value);
	return __r;
}
void* QObject_property(void* self, const char* name) {
	auto&& __r = ((QObject*)self)->property(name);
	return new QVariant(__r);
}
void** QObject_dynamicPropertyNames(void* self, long long* __out_len) {
	auto&& __r = ((QObject*)self)->dynamicPropertyNames();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QByteArray(__e);
	return __arr;
}
void* QObject_bindingStorage(void* self) {
	auto&& __r = ((QObject*)self)->bindingStorage();
	return (void*)(__r);
}
void* QObject_bindingStorage2(void* self) {
	auto&& __r = ((QObject*)self)->bindingStorage();
	return (void*)(__r);
}
void QObject_destroyed(void* self) {
	((QObject*)self)->destroyed();
}
void* QObject_parent(void* self) {
	auto&& __r = ((QObject*)self)->parent();
	return (void*)(__r);
}
int QObject_inherits(void* self, const char* classname) {
	auto&& __r = ((QObject*)self)->inherits(classname);
	return __r;
}
void QObject_deleteLater(void* self) {
	((QObject*)self)->deleteLater();
}
void* QObject_sender(void* self) {
	auto&& __r = ((QObject__Go*)self)->sender();
	return (void*)(__r);
}
int QObject_senderSignalIndex(void* self) {
	auto&& __r = ((QObject__Go*)self)->senderSignalIndex();
	return __r;
}
int QObject_receivers(void* self, const char* signal) {
	auto&& __r = ((QObject__Go*)self)->receivers(signal);
	return __r;
}
int QObject_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QObject__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QObject_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QObject::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QObject_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QObject::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QObject_startTimer3(void* self, int interval, int timerType) {
	auto&& __r = ((QObject*)self)->startTimer(interval, (Qt::TimerType)(timerType));
	return __r;
}
int QObject_startTimer4(void* self, long long time, int timerType) {
	auto&& __r = ((QObject*)self)->startTimer(std::chrono::nanoseconds((long long)time), (Qt::TimerType)(timerType));
	return __r;
}
void* QObject_connect4(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QObject::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QObject_connect5(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QObject::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QObject_connect6(void* self, void* sender, const char* signal, const char* member, int type_) {
	auto&& __r = ((QObject*)self)->connect((QObject*)sender, signal, member, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void QObject_destroyed2(void* self, void* arg1) {
	((QObject*)self)->destroyed((QObject*)arg1);
}
void* QObject_metaObject_default(void* self) {
	auto&& __r = ((QObject*)self)->QObject::metaObject();
	return (void*)(__r);
}
int QObject_event_default(void* self, void* event) {
	auto&& __r = ((QObject*)self)->QObject::event((QEvent*)event);
	return __r;
}
int QObject_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QObject*)self)->QObject::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QObject_timerEvent_default(void* self, void* event) {
	((QObject__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QObject_childEvent_default(void* self, void* event) {
	((QObject__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QObject_customEvent_default(void* self, void* event) {
	((QObject__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QObject_connectNotify_default(void* self, void* signal) {
	((QObject__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QObject_disconnectNotify_default(void* self, void* signal) {
	((QObject__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QObject_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QObject*)self, &QObject::destroyed, (QObject*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QObject);
	__conns_QObject[__id] = __c;
}
void QObject_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QObject*)self, &QObject::destroyed, (QObject*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QObject);
	__conns_QObject[__id] = __c;
}
void QObject_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QObject*)self, &QObject::objectNameChanged, (QObject*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QObject);
	__conns_QObject[__id] = __c;
}
void QObject__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QObject);
		auto __it = __conns_QObject.find(id);
		if (__it == __conns_QObject.end()) return;
		__c = __it->second;
		__conns_QObject.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QObject__Go_new(long long id, void* parent) { return new QObject__Go(id, (QObject*)parent); }
void QObject_delete(void* self) { delete (QObject*)self; }

}

#include <QtCore/qabstractitemmodel.h>
#include <QtCore/qvariant.h>
#include "gen_qmodelroledataspan.h"

extern "C" {

void* QModelRoleDataSpan_new() {
	return new QModelRoleDataSpan();
}
void* QModelRoleDataSpan_new2(void* modelRoleData) {
	return new QModelRoleDataSpan(*(QModelRoleData*)modelRoleData);
}
void* QModelRoleDataSpan_new3(void* modelRoleData, long long len_) {
	return new QModelRoleDataSpan((QModelRoleData*)modelRoleData, len_);
}
void* QModelRoleDataSpan_new4(void* arg1) {
	return new QModelRoleDataSpan(*(QModelRoleDataSpan*)arg1);
}
long long QModelRoleDataSpan_size(void* self) {
	auto&& __r = ((QModelRoleDataSpan*)self)->size();
	return __r;
}
long long QModelRoleDataSpan_length(void* self) {
	auto&& __r = ((QModelRoleDataSpan*)self)->length();
	return __r;
}
void* QModelRoleDataSpan_data(void* self) {
	auto&& __r = ((QModelRoleDataSpan*)self)->data();
	return (void*)(__r);
}
void* QModelRoleDataSpan_begin(void* self) {
	auto&& __r = ((QModelRoleDataSpan*)self)->begin();
	return (void*)(__r);
}
void* QModelRoleDataSpan_end(void* self) {
	auto&& __r = ((QModelRoleDataSpan*)self)->end();
	return (void*)(__r);
}
void* QModelRoleDataSpan_OperatorSubscript(void* self, long long index) {
	auto&& __r = ((QModelRoleDataSpan*)self)->operator[](index);
	return (void*)&(__r);
}
void* QModelRoleDataSpan_dataForRole(void* self, int role) {
	auto&& __r = ((QModelRoleDataSpan*)self)->dataForRole(role);
	return (void*)(__r);
}
void QModelRoleDataSpan_delete(void* self) { delete (QModelRoleDataSpan*)self; }

}

#include <QtCore/qbuffer.h>
#include <QtCore/qiodevice.h>
#include <QtCore/qobject.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qiodevicebase.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qbytearrayview.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qvariant.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qbuffer.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QBuffer__Go : public QBuffer {
public:
	long long __goId;
	unsigned char __goOn[22] = {};
	template<typename... A> QBuffer__Go(long long id, A&&... a) : QBuffer(std::forward<A>(a)...), __goId(id) {}
	using QBuffer::setOpenMode;
	using QBuffer::setErrorString;
	using QBuffer::sender;
	using QBuffer::senderSignalIndex;
	using QBuffer::receivers;
	using QBuffer::isSignalConnected;
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QBuffer::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QBuffer::disconnectNotify(__a0); }
	qint64 cppgoBase_skipData(qint64 __a0) { return QBuffer::skipData(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QBuffer::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QBuffer::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QBuffer::customEvent(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QBuffer::metaObject();
	}
	bool open(QFlags<QIODeviceBase::OpenModeFlag> openMode) override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, -554057428128861975LL, (int)(openMode), &__h); if (__h) return __r; }
		return QBuffer::open(openMode);
	}
	void close() override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, -7006921981232612925LL, &__h); if (__h) return; }
		QBuffer::close();
	}
	qint64 size() const override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_longlong) { long long __r = ((long long(*)(long long, long long, int*))cppgo_virt_longlong)(__goId, 5614464919154503228LL, &__h); if (__h) return __r; }
		return QBuffer::size();
	}
	qint64 pos() const override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_longlong) { long long __r = ((long long(*)(long long, long long, int*))cppgo_virt_longlong)(__goId, 8618582476283917417LL, &__h); if (__h) return __r; }
		return QBuffer::pos();
	}
	bool seek(qint64 off) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_int_longlong) { int __r = ((int(*)(long long, long long, long long, int*))cppgo_virt_int_longlong)(__goId, 3241140482992819217LL, off, &__h); if (__h) return __r; }
		return QBuffer::seek(off);
	}
	bool atEnd() const override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -1487528724786208339LL, &__h); if (__h) return __r; }
		return QBuffer::atEnd();
	}
	bool canReadLine() const override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -1250682508954504761LL, &__h); if (__h) return __r; }
		return QBuffer::canReadLine();
	}
	void connectNotify(const QMetaMethod& arg1) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&arg1, &__h); if (__h) return; }
		QBuffer::connectNotify(arg1);
	}
	void disconnectNotify(const QMetaMethod& arg1) override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&arg1, &__h); if (__h) return; }
		QBuffer::disconnectNotify(arg1);
	}
	bool isSequential() const override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, 8888098400087344862LL, &__h); if (__h) return __r; }
		return QBuffer::isSequential();
	}
	bool reset() override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, 1086335023529244512LL, &__h); if (__h) return __r; }
		return QBuffer::reset();
	}
	qint64 bytesAvailable() const override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_longlong) { long long __r = ((long long(*)(long long, long long, int*))cppgo_virt_longlong)(__goId, 4814449630054007943LL, &__h); if (__h) return __r; }
		return QBuffer::bytesAvailable();
	}
	qint64 bytesToWrite() const override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_longlong) { long long __r = ((long long(*)(long long, long long, int*))cppgo_virt_longlong)(__goId, 1024747515116322948LL, &__h); if (__h) return __r; }
		return QBuffer::bytesToWrite();
	}
	bool waitForReadyRead(int msecs) override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 8356976193102722774LL, msecs, &__h); if (__h) return __r; }
		return QBuffer::waitForReadyRead(msecs);
	}
	bool waitForBytesWritten(int msecs) override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 7726833620054767889LL, msecs, &__h); if (__h) return __r; }
		return QBuffer::waitForBytesWritten(msecs);
	}
	qint64 skipData(qint64 maxSize) override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_longlong_longlong) { long long __r = ((long long(*)(long long, long long, long long, int*))cppgo_virt_longlong_longlong)(__goId, 8178738247667652082LL, maxSize, &__h); if (__h) return __r; }
		return QBuffer::skipData(maxSize);
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QBuffer::event(event);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QBuffer::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QBuffer::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QBuffer::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QBuffer::customEvent(event);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QBuffer;
static std::mutex __connsMu_QBuffer;

extern "C" {

void* QBuffer_base(void* self) { return (void*)static_cast<QIODevice*>((QBuffer*)self); }
void QBuffer_chain(void* self, void** out) {
	QBuffer* __p = (QBuffer*)self;
	QIODevice* __b0 = static_cast<QIODevice*>(__p); out[0] = (void*)__b0;
	QObject* __b1 = static_cast<QObject*>(__b0); out[1] = (void*)__b1;
}
void QBuffer__set_override(void* self, int idx) { ((QBuffer__Go*)self)->__goOn[idx] = 1; }
void* QBuffer_fromQObject(void* o) { return (void*)qobject_cast<QBuffer*>((QObject*)o); }
void* QBuffer_new(long long __goid) {
	return new QBuffer__Go(__goid);
}
void* QBuffer_new2(long long __goid, void* buf) {
	return new QBuffer__Go(__goid, (QByteArray*)buf);
}
void* QBuffer_new3(long long __goid, void* parent) {
	return new QBuffer__Go(__goid, (QObject*)parent);
}
void* QBuffer_new4(long long __goid, void* buf, void* parent) {
	return new QBuffer__Go(__goid, (QByteArray*)buf, (QObject*)parent);
}
char* QBuffer_tr(const char* s, long long* __out_len) {
	auto&& __r = QBuffer::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QBuffer_buffer(void* self) {
	auto&& __r = ((QBuffer*)self)->buffer();
	return (void*)&(__r);
}
void* QBuffer_buffer2(void* self) {
	auto&& __r = ((QBuffer*)self)->buffer();
	return (void*)&(__r);
}
void QBuffer_setBuffer(void* self, void* a) {
	((QBuffer*)self)->setBuffer((QByteArray*)a);
}
void QBuffer_setData(void* self, void* data) {
	((QBuffer*)self)->setData(*(QByteArray*)data);
}
void QBuffer_setData2(void* self, const char* data, long long len_) {
	((QBuffer*)self)->setData(data, len_);
}
void* QBuffer_data(void* self) {
	auto&& __r = ((QBuffer*)self)->data();
	return (void*)&(__r);
}
void QBuffer_setOpenMode(void* self, int openMode) {
	((QBuffer__Go*)self)->setOpenMode(QFlags<QIODeviceBase::OpenModeFlag>(openMode));
}
void QBuffer_setErrorString(void* self, const char* errorString, long long errorString_len) {
	QString errorString_qstr = QString::fromUtf8(errorString, errorString_len);
	((QBuffer__Go*)self)->setErrorString(errorString_qstr);
}
void* QBuffer_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QBuffer::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QBuffer_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QBuffer::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QBuffer_sender(void* self) {
	auto&& __r = ((QBuffer__Go*)self)->sender();
	return (void*)(__r);
}
int QBuffer_senderSignalIndex(void* self) {
	auto&& __r = ((QBuffer__Go*)self)->senderSignalIndex();
	return __r;
}
int QBuffer_receivers(void* self, const char* signal) {
	auto&& __r = ((QBuffer__Go*)self)->receivers(signal);
	return __r;
}
int QBuffer_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QBuffer__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QBuffer_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QBuffer::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QBuffer_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QBuffer::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QBuffer_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QBuffer::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QBuffer_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QBuffer::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QBuffer_metaObject_default(void* self) {
	auto&& __r = ((QBuffer*)self)->QBuffer::metaObject();
	return (void*)(__r);
}
int QBuffer_open_default(void* self, int openMode) {
	auto&& __r = ((QBuffer*)self)->QBuffer::open(QFlags<QIODeviceBase::OpenModeFlag>(openMode));
	return __r;
}
void QBuffer_close_default(void* self) {
	((QBuffer*)self)->QBuffer::close();
}
long long QBuffer_size_default(void* self) {
	auto&& __r = ((QBuffer*)self)->QBuffer::size();
	return __r;
}
long long QBuffer_pos_default(void* self) {
	auto&& __r = ((QBuffer*)self)->QBuffer::pos();
	return __r;
}
int QBuffer_seek_default(void* self, long long off) {
	auto&& __r = ((QBuffer*)self)->QBuffer::seek(off);
	return __r;
}
int QBuffer_atEnd_default(void* self) {
	auto&& __r = ((QBuffer*)self)->QBuffer::atEnd();
	return __r;
}
int QBuffer_canReadLine_default(void* self) {
	auto&& __r = ((QBuffer*)self)->QBuffer::canReadLine();
	return __r;
}
void QBuffer_connectNotify_default(void* self, void* arg1) {
	((QBuffer__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)arg1);
}
void QBuffer_disconnectNotify_default(void* self, void* arg1) {
	((QBuffer__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)arg1);
}
int QBuffer_isSequential_default(void* self) {
	auto&& __r = ((QBuffer*)self)->QBuffer::isSequential();
	return __r;
}
int QBuffer_reset_default(void* self) {
	auto&& __r = ((QBuffer*)self)->QBuffer::reset();
	return __r;
}
long long QBuffer_bytesAvailable_default(void* self) {
	auto&& __r = ((QBuffer*)self)->QBuffer::bytesAvailable();
	return __r;
}
long long QBuffer_bytesToWrite_default(void* self) {
	auto&& __r = ((QBuffer*)self)->QBuffer::bytesToWrite();
	return __r;
}
int QBuffer_waitForReadyRead_default(void* self, int msecs) {
	auto&& __r = ((QBuffer*)self)->QBuffer::waitForReadyRead(msecs);
	return __r;
}
int QBuffer_waitForBytesWritten_default(void* self, int msecs) {
	auto&& __r = ((QBuffer*)self)->QBuffer::waitForBytesWritten(msecs);
	return __r;
}
long long QBuffer_skipData_default(void* self, long long maxSize) {
	auto&& __r = ((QBuffer__Go*)self)->cppgoBase_skipData(maxSize);
	return __r;
}
int QBuffer_event_default(void* self, void* event) {
	auto&& __r = ((QBuffer*)self)->QBuffer::event((QEvent*)event);
	return __r;
}
int QBuffer_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QBuffer*)self)->QBuffer::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QBuffer_timerEvent_default(void* self, void* event) {
	((QBuffer__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QBuffer_childEvent_default(void* self, void* event) {
	((QBuffer__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QBuffer_customEvent_default(void* self, void* event) {
	((QBuffer__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QBuffer_connect_readyRead(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QBuffer*)self, &QBuffer::readyRead, (QBuffer*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QBuffer);
	__conns_QBuffer[__id] = __c;
}
void QBuffer_connect_channelReadyRead(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QBuffer*)self, &QBuffer::channelReadyRead, (QBuffer*)self, [__id](int channel){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, channel);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QBuffer);
	__conns_QBuffer[__id] = __c;
}
void QBuffer_connect_bytesWritten(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QBuffer*)self, &QBuffer::bytesWritten, (QBuffer*)self, [__id](qint64 bytes){
		if (cppgo_sig_void_longlong) ((void(*)(long long, long long))cppgo_sig_void_longlong)(__id, bytes);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QBuffer);
	__conns_QBuffer[__id] = __c;
}
void QBuffer_connect_channelBytesWritten(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QBuffer*)self, &QBuffer::channelBytesWritten, (QBuffer*)self, [__id](int channel, qint64 bytes){
		if (cppgo_sig_void_int_longlong) ((void(*)(long long, int, long long))cppgo_sig_void_int_longlong)(__id, channel, bytes);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QBuffer);
	__conns_QBuffer[__id] = __c;
}
void QBuffer_connect_aboutToClose(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QBuffer*)self, &QBuffer::aboutToClose, (QBuffer*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QBuffer);
	__conns_QBuffer[__id] = __c;
}
void QBuffer_connect_readChannelFinished(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QBuffer*)self, &QBuffer::readChannelFinished, (QBuffer*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QBuffer);
	__conns_QBuffer[__id] = __c;
}
void QBuffer_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QBuffer*)self, &QBuffer::destroyed, (QBuffer*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QBuffer);
	__conns_QBuffer[__id] = __c;
}
void QBuffer_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QBuffer*)self, &QBuffer::destroyed, (QBuffer*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QBuffer);
	__conns_QBuffer[__id] = __c;
}
void QBuffer_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QBuffer*)self, &QBuffer::objectNameChanged, (QBuffer*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QBuffer);
	__conns_QBuffer[__id] = __c;
}
void QBuffer__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QBuffer);
		auto __it = __conns_QBuffer.find(id);
		if (__it == __conns_QBuffer.end()) return;
		__c = __it->second;
		__conns_QBuffer.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QBuffer__Go_new(long long id, void* parent) { return new QBuffer__Go(id, (QObject*)parent); }
void QBuffer_delete(void* self) { delete (QBuffer*)self; }

}

#include <QtCore/qcborstreamwriter.h>
#include <QtCore/qiodevice.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qcborcommon.h>
#include <QtCore/qbytearrayview.h>
#include <QtCore/qlatin1stringview.h>
#include <QtCore/qstringview.h>
#include <QtCore/qfloat16.h>
#include "gen_qcborstreamwriter.h"

extern "C" {

void* QCborStreamWriter_new(void* device) {
	return new QCborStreamWriter((QIODevice*)device);
}
void* QCborStreamWriter_new2(void* data) {
	return new QCborStreamWriter((QByteArray*)data);
}
void QCborStreamWriter_setDevice(void* self, void* device) {
	((QCborStreamWriter*)self)->setDevice((QIODevice*)device);
}
void* QCborStreamWriter_device(void* self) {
	auto&& __r = ((QCborStreamWriter*)self)->device();
	return (void*)(__r);
}
void QCborStreamWriter_append(void* self, unsigned long long u) {
	((QCborStreamWriter*)self)->append(u);
}
void QCborStreamWriter_append2(void* self, long long i) {
	((QCborStreamWriter*)self)->append(i);
}
void QCborStreamWriter_append3(void* self, int n) {
	((QCborStreamWriter*)self)->append((QCborNegativeInteger)(n));
}
void QCborStreamWriter_append4(void* self, void* ba) {
	((QCborStreamWriter*)self)->append(*(QByteArrayView*)ba);
}
void QCborStreamWriter_append5(void* self, void* str) {
	((QCborStreamWriter*)self)->append(*(QLatin1String*)str);
}
void QCborStreamWriter_append6(void* self, void* str) {
	((QCborStreamWriter*)self)->append(*(QStringView*)str);
}
void QCborStreamWriter_append7(void* self, int tag) {
	((QCborStreamWriter*)self)->append((QCborTag)(tag));
}
void QCborStreamWriter_append8(void* self, int tag) {
	((QCborStreamWriter*)self)->append((QCborKnownTags)(tag));
}
void QCborStreamWriter_append9(void* self, int st) {
	((QCborStreamWriter*)self)->append((QCborSimpleType)(st));
}
void QCborStreamWriter_append10(void* self, void* f) {
	((QCborStreamWriter*)self)->append(*(qfloat16*)f);
}
void QCborStreamWriter_append11(void* self, float f) {
	((QCborStreamWriter*)self)->append(f);
}
void QCborStreamWriter_append12(void* self, double d) {
	((QCborStreamWriter*)self)->append(d);
}
void QCborStreamWriter_appendByteString(void* self, const char* data, long long len_) {
	((QCborStreamWriter*)self)->appendByteString(data, len_);
}
void QCborStreamWriter_appendTextString(void* self, const char* utf8, long long len_) {
	((QCborStreamWriter*)self)->appendTextString(utf8, len_);
}
void QCborStreamWriter_append13(void* self, int b) {
	((QCborStreamWriter*)self)->append(b);
}
void QCborStreamWriter_appendNull(void* self) {
	((QCborStreamWriter*)self)->appendNull();
}
void QCborStreamWriter_appendUndefined(void* self) {
	((QCborStreamWriter*)self)->appendUndefined();
}
void QCborStreamWriter_append14(void* self, int i) {
	((QCborStreamWriter*)self)->append(i);
}
void QCborStreamWriter_append15(void* self, unsigned int u) {
	((QCborStreamWriter*)self)->append(u);
}
void QCborStreamWriter_append16(void* self, const char* str) {
	((QCborStreamWriter*)self)->append(str);
}
void QCborStreamWriter_startArray(void* self) {
	((QCborStreamWriter*)self)->startArray();
}
void QCborStreamWriter_startArray2(void* self, unsigned long long count) {
	((QCborStreamWriter*)self)->startArray(count);
}
int QCborStreamWriter_endArray(void* self) {
	auto&& __r = ((QCborStreamWriter*)self)->endArray();
	return __r;
}
void QCborStreamWriter_startMap(void* self) {
	((QCborStreamWriter*)self)->startMap();
}
void QCborStreamWriter_startMap2(void* self, unsigned long long count) {
	((QCborStreamWriter*)self)->startMap(count);
}
int QCborStreamWriter_endMap(void* self) {
	auto&& __r = ((QCborStreamWriter*)self)->endMap();
	return __r;
}
void QCborStreamWriter_append17(void* self, const char* str, long long size) {
	((QCborStreamWriter*)self)->append(str, size);
}
void QCborStreamWriter_delete(void* self) { delete (QCborStreamWriter*)self; }

}

#include <QtCore/qfile.h>
#include <QtCore/qfiledevice.h>
#include <QtCore/qstring.h>
#include <QtCore/qobject.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qiodevicebase.h>
#include <QtCore/qdatetime.h>
#include <QtCore/qbytearrayview.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qvariant.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qfile.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QFile__Go : public QFile {
public:
	long long __goId;
	unsigned char __goOn[26] = {};
	template<typename... A> QFile__Go(long long id, A&&... a) : QFile(std::forward<A>(a)...), __goId(id) {}
	using QFile::setOpenMode;
	using QFile::setErrorString;
	using QFile::sender;
	using QFile::senderSignalIndex;
	using QFile::receivers;
	using QFile::isSignalConnected;
	qint64 cppgoBase_skipData(qint64 __a0) { return QFile::skipData(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QFile::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QFile::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QFile::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QFile::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QFile::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QFile::metaObject();
	}
	QString fileName() const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virtp_charp) {
			long long __len = 0;
			char* __s = ((char*(*)(long long, long long, long long*, int*))cppgo_virtp_charp)(__goId, 8848495542063941256LL, &__len, &__h);
			if (__h) { QString __r = QString::fromUtf8(__s, __len); free(__s); return __r; }
		}
		return QFile::fileName();
	}
	bool open(QFlags<QIODeviceBase::OpenModeFlag> flags) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, -554057428128861975LL, (int)(flags), &__h); if (__h) return __r; }
		return QFile::open(flags);
	}
	qint64 size() const override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_longlong) { long long __r = ((long long(*)(long long, long long, int*))cppgo_virt_longlong)(__goId, 5614464919154503228LL, &__h); if (__h) return __r; }
		return QFile::size();
	}
	bool resize(qint64 sz) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_int_longlong) { int __r = ((int(*)(long long, long long, long long, int*))cppgo_virt_int_longlong)(__goId, -4890198374319873567LL, sz, &__h); if (__h) return __r; }
		return QFile::resize(sz);
	}
	QFlags<QFileDevice::Permission> permissions() const override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -7517927236947282013LL, &__h); if (__h) return QFlags<QFileDevice::Permission>(__r); }
		return QFile::permissions();
	}
	bool setPermissions(QFlags<QFileDevice::Permission> permissionSpec) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 420407675554349589LL, (int)(permissionSpec), &__h); if (__h) return __r; }
		return QFile::setPermissions(permissionSpec);
	}
	void close() override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, -7006921981232612925LL, &__h); if (__h) return; }
		QFile::close();
	}
	bool isSequential() const override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, 8888098400087344862LL, &__h); if (__h) return __r; }
		return QFile::isSequential();
	}
	qint64 pos() const override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_longlong) { long long __r = ((long long(*)(long long, long long, int*))cppgo_virt_longlong)(__goId, 8618582476283917417LL, &__h); if (__h) return __r; }
		return QFile::pos();
	}
	bool seek(qint64 offset) override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_int_longlong) { int __r = ((int(*)(long long, long long, long long, int*))cppgo_virt_int_longlong)(__goId, 3241140482992819217LL, offset, &__h); if (__h) return __r; }
		return QFile::seek(offset);
	}
	bool atEnd() const override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -1487528724786208339LL, &__h); if (__h) return __r; }
		return QFile::atEnd();
	}
	bool reset() override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, 1086335023529244512LL, &__h); if (__h) return __r; }
		return QFile::reset();
	}
	qint64 bytesAvailable() const override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_longlong) { long long __r = ((long long(*)(long long, long long, int*))cppgo_virt_longlong)(__goId, 4814449630054007943LL, &__h); if (__h) return __r; }
		return QFile::bytesAvailable();
	}
	qint64 bytesToWrite() const override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_longlong) { long long __r = ((long long(*)(long long, long long, int*))cppgo_virt_longlong)(__goId, 1024747515116322948LL, &__h); if (__h) return __r; }
		return QFile::bytesToWrite();
	}
	bool canReadLine() const override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -1250682508954504761LL, &__h); if (__h) return __r; }
		return QFile::canReadLine();
	}
	bool waitForReadyRead(int msecs) override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 8356976193102722774LL, msecs, &__h); if (__h) return __r; }
		return QFile::waitForReadyRead(msecs);
	}
	bool waitForBytesWritten(int msecs) override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 7726833620054767889LL, msecs, &__h); if (__h) return __r; }
		return QFile::waitForBytesWritten(msecs);
	}
	qint64 skipData(qint64 maxSize) override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_longlong_longlong) { long long __r = ((long long(*)(long long, long long, long long, int*))cppgo_virt_longlong_longlong)(__goId, 8178738247667652082LL, maxSize, &__h); if (__h) return __r; }
		return QFile::skipData(maxSize);
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QFile::event(event);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QFile::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QFile::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QFile::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QFile::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QFile::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QFile::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QFile;
static std::mutex __connsMu_QFile;

extern "C" {

void* QFile_base(void* self) { return (void*)static_cast<QFileDevice*>((QFile*)self); }
void QFile_chain(void* self, void** out) {
	QFile* __p = (QFile*)self;
	QFileDevice* __b0 = static_cast<QFileDevice*>(__p); out[0] = (void*)__b0;
	QIODevice* __b1 = static_cast<QIODevice*>(__b0); out[1] = (void*)__b1;
	QObject* __b2 = static_cast<QObject*>(__b1); out[2] = (void*)__b2;
}
void QFile__set_override(void* self, int idx) { ((QFile__Go*)self)->__goOn[idx] = 1; }
void* QFile_fromQObject(void* o) { return (void*)qobject_cast<QFile*>((QObject*)o); }
void* QFile_new(long long __goid) {
	return new QFile__Go(__goid);
}
void* QFile_new2(long long __goid, const char* name, long long name_len) {
	QString name_qstr = QString::fromUtf8(name, name_len);
	return new QFile__Go(__goid, name_qstr);
}
void* QFile_new3(long long __goid, void* parent) {
	return new QFile__Go(__goid, (QObject*)parent);
}
void* QFile_new4(long long __goid, const char* name, long long name_len, void* parent) {
	QString name_qstr = QString::fromUtf8(name, name_len);
	return new QFile__Go(__goid, name_qstr, (QObject*)parent);
}
char* QFile_tr(const char* s, long long* __out_len) {
	auto&& __r = QFile::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QFile_setFileName(void* self, const char* name, long long name_len) {
	QString name_qstr = QString::fromUtf8(name, name_len);
	((QFile*)self)->setFileName(name_qstr);
}
void* QFile_encodeName(const char* fileName, long long fileName_len) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	auto&& __r = QFile::encodeName(fileName_qstr);
	return new QByteArray(__r);
}
char* QFile_decodeName(void* localFileName, long long* __out_len) {
	auto&& __r = QFile::decodeName(*(QByteArray*)localFileName);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QFile_decodeName2(const char* localFileName, long long* __out_len) {
	auto&& __r = QFile::decodeName(localFileName);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QFile_exists(void* self) {
	auto&& __r = ((QFile*)self)->exists();
	return __r;
}
int QFile_exists2(const char* fileName, long long fileName_len) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	auto&& __r = QFile::exists(fileName_qstr);
	return __r;
}
char* QFile_symLinkTarget(void* self, long long* __out_len) {
	auto&& __r = ((QFile*)self)->symLinkTarget();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QFile_symLinkTarget2(const char* fileName, long long fileName_len, long long* __out_len) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	auto&& __r = QFile::symLinkTarget(fileName_qstr);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QFile_remove(void* self) {
	auto&& __r = ((QFile*)self)->remove();
	return __r;
}
int QFile_remove2(const char* fileName, long long fileName_len) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	auto&& __r = QFile::remove(fileName_qstr);
	return __r;
}
int QFile_supportsMoveToTrash() {
	auto&& __r = QFile::supportsMoveToTrash();
	return __r;
}
int QFile_moveToTrash(void* self) {
	auto&& __r = ((QFile*)self)->moveToTrash();
	return __r;
}
int QFile_moveToTrash2(const char* fileName, long long fileName_len) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	auto&& __r = QFile::moveToTrash(fileName_qstr);
	return __r;
}
int QFile_rename(void* self, const char* newName, long long newName_len) {
	QString newName_qstr = QString::fromUtf8(newName, newName_len);
	auto&& __r = ((QFile*)self)->rename(newName_qstr);
	return __r;
}
int QFile_rename2(const char* oldName, long long oldName_len, const char* newName, long long newName_len) {
	QString oldName_qstr = QString::fromUtf8(oldName, oldName_len);
	QString newName_qstr = QString::fromUtf8(newName, newName_len);
	auto&& __r = QFile::rename(oldName_qstr, newName_qstr);
	return __r;
}
int QFile_link(void* self, const char* newName, long long newName_len) {
	QString newName_qstr = QString::fromUtf8(newName, newName_len);
	auto&& __r = ((QFile*)self)->link(newName_qstr);
	return __r;
}
int QFile_link2(const char* fileName, long long fileName_len, const char* newName, long long newName_len) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	QString newName_qstr = QString::fromUtf8(newName, newName_len);
	auto&& __r = QFile::link(fileName_qstr, newName_qstr);
	return __r;
}
int QFile_copy(void* self, const char* newName, long long newName_len) {
	QString newName_qstr = QString::fromUtf8(newName, newName_len);
	auto&& __r = ((QFile*)self)->copy(newName_qstr);
	return __r;
}
int QFile_copy2(const char* fileName, long long fileName_len, const char* newName, long long newName_len) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	QString newName_qstr = QString::fromUtf8(newName, newName_len);
	auto&& __r = QFile::copy(fileName_qstr, newName_qstr);
	return __r;
}
int QFile_open(void* self, int flags) {
	auto&& __r = ((QFile*)self)->open(QFlags<QIODeviceBase::OpenModeFlag>(flags));
	return __r;
}
int QFile_open2(void* self, int flags, int permissions) {
	auto&& __r = ((QFile*)self)->open(QFlags<QIODeviceBase::OpenModeFlag>(flags), QFlags<QFileDevice::Permission>(permissions));
	return __r;
}
int QFile_open3(void* self, int fd, int ioFlags) {
	auto&& __r = ((QFile*)self)->open(fd, QFlags<QIODeviceBase::OpenModeFlag>(ioFlags));
	return __r;
}
int QFile_resize(const char* filename, long long filename_len, long long sz) {
	QString filename_qstr = QString::fromUtf8(filename, filename_len);
	auto&& __r = QFile::resize(filename_qstr, sz);
	return __r;
}
int QFile_permissions(const char* filename, long long filename_len) {
	QString filename_qstr = QString::fromUtf8(filename, filename_len);
	auto&& __r = QFile::permissions(filename_qstr);
	return (int)(__r);
}
int QFile_setPermissions(const char* filename, long long filename_len, int permissionSpec) {
	QString filename_qstr = QString::fromUtf8(filename, filename_len);
	auto&& __r = QFile::setPermissions(filename_qstr, QFlags<QFileDevice::Permission>(permissionSpec));
	return __r;
}
void QFile_setOpenMode(void* self, int openMode) {
	((QFile__Go*)self)->setOpenMode(QFlags<QIODeviceBase::OpenModeFlag>(openMode));
}
void QFile_setErrorString(void* self, const char* errorString, long long errorString_len) {
	QString errorString_qstr = QString::fromUtf8(errorString, errorString_len);
	((QFile__Go*)self)->setErrorString(errorString_qstr);
}
void* QFile_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QFile::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QFile_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QFile::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QFile_sender(void* self) {
	auto&& __r = ((QFile__Go*)self)->sender();
	return (void*)(__r);
}
int QFile_senderSignalIndex(void* self) {
	auto&& __r = ((QFile__Go*)self)->senderSignalIndex();
	return __r;
}
int QFile_receivers(void* self, const char* signal) {
	auto&& __r = ((QFile__Go*)self)->receivers(signal);
	return __r;
}
int QFile_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QFile__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QFile_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QFile::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QFile_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QFile::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QFile_moveToTrash3(const char* fileName, long long fileName_len, void* pathInTrash) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	auto&& __r = QFile::moveToTrash(fileName_qstr, (QString*)pathInTrash);
	return __r;
}
int QFile_open4(void* self, int fd, int ioFlags, int handleFlags) {
	auto&& __r = ((QFile*)self)->open(fd, QFlags<QIODeviceBase::OpenModeFlag>(ioFlags), QFlags<QFileDevice::FileHandleFlag>(handleFlags));
	return __r;
}
void* QFile_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QFile::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QFile_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QFile::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
int QFile_open_default(void* self, int flags) {
	auto&& __r = ((QFile*)self)->QFile::open(QFlags<QIODeviceBase::OpenModeFlag>(flags));
	return __r;
}
int QFile_open_default2(void* self, int flags, int permissions) {
	auto&& __r = ((QFile*)self)->QFile::open(QFlags<QIODeviceBase::OpenModeFlag>(flags), QFlags<QFileDevice::Permission>(permissions));
	return __r;
}
int QFile_open_default3(void* self, int fd, int ioFlags) {
	auto&& __r = ((QFile*)self)->QFile::open(fd, QFlags<QIODeviceBase::OpenModeFlag>(ioFlags));
	return __r;
}
int QFile_open_default4(void* self, int fd, int ioFlags, int handleFlags) {
	auto&& __r = ((QFile*)self)->QFile::open(fd, QFlags<QIODeviceBase::OpenModeFlag>(ioFlags), QFlags<QFileDevice::FileHandleFlag>(handleFlags));
	return __r;
}
void* QFile_metaObject_default(void* self) {
	auto&& __r = ((QFile*)self)->QFile::metaObject();
	return (void*)(__r);
}
char* QFile_fileName_default(void* self, long long* __out_len) {
	auto&& __r = ((QFile*)self)->QFile::fileName();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
long long QFile_size_default(void* self) {
	auto&& __r = ((QFile*)self)->QFile::size();
	return __r;
}
int QFile_resize_default(void* self, long long sz) {
	auto&& __r = ((QFile*)self)->QFile::resize(sz);
	return __r;
}
int QFile_permissions_default(void* self) {
	auto&& __r = ((QFile*)self)->QFile::permissions();
	return (int)(__r);
}
int QFile_setPermissions_default(void* self, int permissionSpec) {
	auto&& __r = ((QFile*)self)->QFile::setPermissions(QFlags<QFileDevice::Permission>(permissionSpec));
	return __r;
}
void QFile_close_default(void* self) {
	((QFile*)self)->QFile::close();
}
int QFile_isSequential_default(void* self) {
	auto&& __r = ((QFile*)self)->QFile::isSequential();
	return __r;
}
long long QFile_pos_default(void* self) {
	auto&& __r = ((QFile*)self)->QFile::pos();
	return __r;
}
int QFile_seek_default(void* self, long long offset) {
	auto&& __r = ((QFile*)self)->QFile::seek(offset);
	return __r;
}
int QFile_atEnd_default(void* self) {
	auto&& __r = ((QFile*)self)->QFile::atEnd();
	return __r;
}
int QFile_reset_default(void* self) {
	auto&& __r = ((QFile*)self)->QFile::reset();
	return __r;
}
long long QFile_bytesAvailable_default(void* self) {
	auto&& __r = ((QFile*)self)->QFile::bytesAvailable();
	return __r;
}
long long QFile_bytesToWrite_default(void* self) {
	auto&& __r = ((QFile*)self)->QFile::bytesToWrite();
	return __r;
}
int QFile_canReadLine_default(void* self) {
	auto&& __r = ((QFile*)self)->QFile::canReadLine();
	return __r;
}
int QFile_waitForReadyRead_default(void* self, int msecs) {
	auto&& __r = ((QFile*)self)->QFile::waitForReadyRead(msecs);
	return __r;
}
int QFile_waitForBytesWritten_default(void* self, int msecs) {
	auto&& __r = ((QFile*)self)->QFile::waitForBytesWritten(msecs);
	return __r;
}
long long QFile_skipData_default(void* self, long long maxSize) {
	auto&& __r = ((QFile__Go*)self)->cppgoBase_skipData(maxSize);
	return __r;
}
int QFile_event_default(void* self, void* event) {
	auto&& __r = ((QFile*)self)->QFile::event((QEvent*)event);
	return __r;
}
int QFile_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QFile*)self)->QFile::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QFile_timerEvent_default(void* self, void* event) {
	((QFile__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QFile_childEvent_default(void* self, void* event) {
	((QFile__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QFile_customEvent_default(void* self, void* event) {
	((QFile__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QFile_connectNotify_default(void* self, void* signal) {
	((QFile__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QFile_disconnectNotify_default(void* self, void* signal) {
	((QFile__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QFile_connect_readyRead(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFile*)self, &QFile::readyRead, (QFile*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFile);
	__conns_QFile[__id] = __c;
}
void QFile_connect_channelReadyRead(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFile*)self, &QFile::channelReadyRead, (QFile*)self, [__id](int channel){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, channel);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFile);
	__conns_QFile[__id] = __c;
}
void QFile_connect_bytesWritten(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFile*)self, &QFile::bytesWritten, (QFile*)self, [__id](qint64 bytes){
		if (cppgo_sig_void_longlong) ((void(*)(long long, long long))cppgo_sig_void_longlong)(__id, bytes);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFile);
	__conns_QFile[__id] = __c;
}
void QFile_connect_channelBytesWritten(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFile*)self, &QFile::channelBytesWritten, (QFile*)self, [__id](int channel, qint64 bytes){
		if (cppgo_sig_void_int_longlong) ((void(*)(long long, int, long long))cppgo_sig_void_int_longlong)(__id, channel, bytes);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFile);
	__conns_QFile[__id] = __c;
}
void QFile_connect_aboutToClose(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFile*)self, &QFile::aboutToClose, (QFile*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFile);
	__conns_QFile[__id] = __c;
}
void QFile_connect_readChannelFinished(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFile*)self, &QFile::readChannelFinished, (QFile*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFile);
	__conns_QFile[__id] = __c;
}
void QFile_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFile*)self, &QFile::destroyed, (QFile*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFile);
	__conns_QFile[__id] = __c;
}
void QFile_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFile*)self, &QFile::destroyed, (QFile*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFile);
	__conns_QFile[__id] = __c;
}
void QFile_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFile*)self, &QFile::objectNameChanged, (QFile*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFile);
	__conns_QFile[__id] = __c;
}
void QFile__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QFile);
		auto __it = __conns_QFile.find(id);
		if (__it == __conns_QFile.end()) return;
		__c = __it->second;
		__conns_QFile.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QFile__Go_new(long long id, void* parent) { return new QFile__Go(id, (QObject*)parent); }
void QFile_delete(void* self) { delete (QFile*)self; }

}

#include <QtCore/qjsonvalue.h>
#include <QtCore/qstring.h>
#include <QtCore/qlatin1stringview.h>
#include <QtCore/qjsonarray.h>
#include <QtCore/qjsonobject.h>
#include <QtCore/qvariant.h>
#include <QtCore/qbytearrayview.h>
#include <QtCore/qjsonparseerror.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qjsondocument.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qstringview.h>
#include "gen_qjsonvalue.h"

extern "C" {

void* QJsonValue_new() {
	return new QJsonValue();
}
void* QJsonValue_new2(int b) {
	return new QJsonValue(b);
}
void* QJsonValue_new3(double n) {
	return new QJsonValue(n);
}
void* QJsonValue_new4(int n) {
	return new QJsonValue(n);
}
void* QJsonValue_new5(long long v) {
	return new QJsonValue(v);
}
void* QJsonValue_new6(const char* s, long long s_len) {
	QString s_qstr = QString::fromUtf8(s, s_len);
	return new QJsonValue(s_qstr);
}
void* QJsonValue_new7(void* s) {
	return new QJsonValue(*(QLatin1String*)s);
}
void* QJsonValue_new8(const char* s) {
	return new QJsonValue(s);
}
void* QJsonValue_new9(void* a) {
	return new QJsonValue(*(QJsonArray*)a);
}
void* QJsonValue_new10(void* o) {
	return new QJsonValue(*(QJsonObject*)o);
}
void* QJsonValue_new11(void* other) {
	return new QJsonValue(*(QJsonValue*)other);
}
void* QJsonValue_new12(int arg1) {
	return new QJsonValue((QJsonValue::Type)(arg1));
}
void* QJsonValue_OperatorAssign(void* self, void* other) {
	auto&& __r = ((QJsonValue*)self)->operator=(*(QJsonValue*)other);
	return (void*)&(__r);
}
void QJsonValue_swap(void* self, void* other) {
	((QJsonValue*)self)->swap(*(QJsonValue*)other);
}
void* QJsonValue_fromVariant(void* variant) {
	auto&& __r = QJsonValue::fromVariant(*(QVariant*)variant);
	return new QJsonValue(__r);
}
void* QJsonValue_toVariant(void* self) {
	auto&& __r = ((QJsonValue*)self)->toVariant();
	return new QVariant(__r);
}
void* QJsonValue_fromJson(void* json) {
	auto&& __r = QJsonValue::fromJson(*(QByteArrayView*)json);
	return new QJsonValue(__r);
}
void* QJsonValue_toJson(void* self) {
	auto&& __r = ((QJsonValue*)self)->toJson();
	return new QByteArray(__r);
}
int QJsonValue_type(void* self) {
	auto&& __r = ((QJsonValue*)self)->type();
	return (int)(__r);
}
int QJsonValue_isNull(void* self) {
	auto&& __r = ((QJsonValue*)self)->isNull();
	return __r;
}
int QJsonValue_isBool(void* self) {
	auto&& __r = ((QJsonValue*)self)->isBool();
	return __r;
}
int QJsonValue_isDouble(void* self) {
	auto&& __r = ((QJsonValue*)self)->isDouble();
	return __r;
}
int QJsonValue_isString(void* self) {
	auto&& __r = ((QJsonValue*)self)->isString();
	return __r;
}
int QJsonValue_isArray(void* self) {
	auto&& __r = ((QJsonValue*)self)->isArray();
	return __r;
}
int QJsonValue_isObject(void* self) {
	auto&& __r = ((QJsonValue*)self)->isObject();
	return __r;
}
int QJsonValue_isUndefined(void* self) {
	auto&& __r = ((QJsonValue*)self)->isUndefined();
	return __r;
}
int QJsonValue_toBool(void* self) {
	auto&& __r = ((QJsonValue*)self)->toBool();
	return __r;
}
int QJsonValue_toInt(void* self) {
	auto&& __r = ((QJsonValue*)self)->toInt();
	return __r;
}
long long QJsonValue_toInteger(void* self) {
	auto&& __r = ((QJsonValue*)self)->toInteger();
	return __r;
}
double QJsonValue_toDouble(void* self) {
	auto&& __r = ((QJsonValue*)self)->toDouble();
	return __r;
}
char* QJsonValue_toString(void* self, long long* __out_len) {
	auto&& __r = ((QJsonValue*)self)->toString();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QJsonValue_toString2(void* self, const char* defaultValue, long long defaultValue_len, long long* __out_len) {
	QString defaultValue_qstr = QString::fromUtf8(defaultValue, defaultValue_len);
	auto&& __r = ((QJsonValue*)self)->toString(defaultValue_qstr);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QJsonValue_toStringView(void* self) {
	auto&& __r = ((QJsonValue*)self)->toStringView();
	return new QAnyStringView(__r);
}
void* QJsonValue_toArray(void* self) {
	auto&& __r = ((QJsonValue*)self)->toArray();
	return new QJsonArray(__r);
}
void* QJsonValue_toArray2(void* self, void* defaultValue) {
	auto&& __r = ((QJsonValue*)self)->toArray(*(QJsonArray*)defaultValue);
	return new QJsonArray(__r);
}
void* QJsonValue_toObject(void* self) {
	auto&& __r = ((QJsonValue*)self)->toObject();
	return new QJsonObject(__r);
}
void* QJsonValue_toObject2(void* self, void* defaultValue) {
	auto&& __r = ((QJsonValue*)self)->toObject(*(QJsonObject*)defaultValue);
	return new QJsonObject(__r);
}
void* QJsonValue_OperatorSubscript(void* self, const char* key, long long key_len) {
	QString key_qstr = QString::fromUtf8(key, key_len);
	auto&& __r = ((QJsonValue*)self)->operator[](key_qstr);
	return new QJsonValue(__r);
}
void* QJsonValue_OperatorSubscript2(void* self, void* key) {
	auto&& __r = ((QJsonValue*)self)->operator[](*(QStringView*)key);
	return new QJsonValue(__r);
}
void* QJsonValue_OperatorSubscript3(void* self, void* key) {
	auto&& __r = ((QJsonValue*)self)->operator[](*(QLatin1String*)key);
	return new QJsonValue(__r);
}
void* QJsonValue_OperatorSubscript4(void* self, long long i) {
	auto&& __r = ((QJsonValue*)self)->operator[](i);
	return new QJsonValue(__r);
}
void* QJsonValue_fromJson2(void* json, void* error_) {
	auto&& __r = QJsonValue::fromJson(*(QByteArrayView*)json, (QJsonParseError*)error_);
	return new QJsonValue(__r);
}
void* QJsonValue_toJson2(void* self, int format) {
	auto&& __r = ((QJsonValue*)self)->toJson((QJsonDocument::JsonFormat)(format));
	return new QByteArray(__r);
}
int QJsonValue_toBool2(void* self, int defaultValue) {
	auto&& __r = ((QJsonValue*)self)->toBool(defaultValue);
	return __r;
}
int QJsonValue_toInt2(void* self, int defaultValue) {
	auto&& __r = ((QJsonValue*)self)->toInt(defaultValue);
	return __r;
}
long long QJsonValue_toInteger2(void* self, long long defaultValue) {
	auto&& __r = ((QJsonValue*)self)->toInteger(defaultValue);
	return __r;
}
double QJsonValue_toDouble2(void* self, double defaultValue) {
	auto&& __r = ((QJsonValue*)self)->toDouble(defaultValue);
	return __r;
}
void* QJsonValue_toStringView2(void* self, const char* defaultValue, long long defaultValue_len) {
	QAnyStringView defaultValue_qstr = QAnyStringView(defaultValue, defaultValue_len);
	auto&& __r = ((QJsonValue*)self)->toStringView(defaultValue_qstr);
	return new QAnyStringView(__r);
}
int QJsonValue_OperatorEqual(void* self, void* rhs) {
	auto&& __r = (*(QJsonValue*)self) == (*(QJsonValue*)rhs);
	return __r;
}
int QJsonValue_OperatorNotEqual(void* self, void* rhs) {
	auto&& __r = (*(QJsonValue*)self) != (*(QJsonValue*)rhs);
	return __r;
}
int QJsonValue_OperatorEqual2(void* self, void* rhs) {
	auto&& __r = (*(QJsonValue*)self) == (*(QJsonValueConstRef*)rhs);
	return __r;
}
int QJsonValue_OperatorNotEqual2(void* self, void* rhs) {
	auto&& __r = (*(QJsonValue*)self) != (*(QJsonValueConstRef*)rhs);
	return __r;
}
int QJsonValue_OperatorEqual3(void* self, void* rhs) {
	auto&& __r = (*(QJsonValue*)self) == (*(QJsonArray*)rhs);
	return __r;
}
int QJsonValue_OperatorNotEqual3(void* self, void* rhs) {
	auto&& __r = (*(QJsonValue*)self) != (*(QJsonArray*)rhs);
	return __r;
}
int QJsonValue_OperatorEqual4(void* self, void* rhs) {
	auto&& __r = (*(QJsonValue*)self) == (*(QJsonObject*)rhs);
	return __r;
}
int QJsonValue_OperatorNotEqual4(void* self, void* rhs) {
	auto&& __r = (*(QJsonValue*)self) != (*(QJsonObject*)rhs);
	return __r;
}
void QJsonValue_delete(void* self) { delete (QJsonValue*)self; }

}

#include <QtCore/qmimetype.h>
#include <QtCore/qstring.h>
#include "gen_qmimetype.h"

extern "C" {

void* QMimeType_new() {
	return new QMimeType();
}
void* QMimeType_new2(void* other) {
	return new QMimeType(*(QMimeType*)other);
}
void* QMimeType_OperatorAssign(void* self, void* other) {
	auto&& __r = ((QMimeType*)self)->operator=(*(QMimeType*)other);
	return (void*)&(__r);
}
void QMimeType_swap(void* self, void* other) {
	((QMimeType*)self)->swap(*(QMimeType*)other);
}
int QMimeType_isValid(void* self) {
	auto&& __r = ((QMimeType*)self)->isValid();
	return __r;
}
int QMimeType_isDefault(void* self) {
	auto&& __r = ((QMimeType*)self)->isDefault();
	return __r;
}
char* QMimeType_name(void* self, long long* __out_len) {
	auto&& __r = ((QMimeType*)self)->name();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QMimeType_comment(void* self, long long* __out_len) {
	auto&& __r = ((QMimeType*)self)->comment();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QMimeType_genericIconName(void* self, long long* __out_len) {
	auto&& __r = ((QMimeType*)self)->genericIconName();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QMimeType_iconName(void* self, long long* __out_len) {
	auto&& __r = ((QMimeType*)self)->iconName();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char** QMimeType_globPatterns(void* self, long long* __out_len) {
	auto&& __r = ((QMimeType*)self)->globPatterns();
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
char** QMimeType_parentMimeTypes(void* self, long long* __out_len) {
	auto&& __r = ((QMimeType*)self)->parentMimeTypes();
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
char** QMimeType_allAncestors(void* self, long long* __out_len) {
	auto&& __r = ((QMimeType*)self)->allAncestors();
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
char** QMimeType_aliases(void* self, long long* __out_len) {
	auto&& __r = ((QMimeType*)self)->aliases();
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
char** QMimeType_suffixes(void* self, long long* __out_len) {
	auto&& __r = ((QMimeType*)self)->suffixes();
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
char* QMimeType_preferredSuffix(void* self, long long* __out_len) {
	auto&& __r = ((QMimeType*)self)->preferredSuffix();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QMimeType_inherits(void* self, const char* mimeTypeName, long long mimeTypeName_len) {
	QString mimeTypeName_qstr = QString::fromUtf8(mimeTypeName, mimeTypeName_len);
	auto&& __r = ((QMimeType*)self)->inherits(mimeTypeName_qstr);
	return __r;
}
char* QMimeType_filterString(void* self, long long* __out_len) {
	auto&& __r = ((QMimeType*)self)->filterString();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QMimeType_OperatorEqual(void* self, void* rhs) {
	auto&& __r = (*(QMimeType*)self) == (*(QMimeType*)rhs);
	return __r;
}
int QMimeType_OperatorNotEqual(void* self, void* rhs) {
	auto&& __r = (*(QMimeType*)self) != (*(QMimeType*)rhs);
	return __r;
}
void QMimeType_delete(void* self) { delete (QMimeType*)self; }

}

#include <QtCore/qvariantanimation.h>
#include <QtCore/qabstractanimation.h>
#include <QtCore/qobject.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qvariant.h>
#include <QtCore/qeasingcurve.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qanimationgroup.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qvariantanimation.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QVariantAnimation__Go : public QVariantAnimation {
public:
	long long __goId;
	unsigned char __goOn[14] = {};
	template<typename... A> QVariantAnimation__Go(long long id, A&&... a) : QVariantAnimation(std::forward<A>(a)...), __goId(id) {}
	using QVariantAnimation::sender;
	using QVariantAnimation::senderSignalIndex;
	using QVariantAnimation::receivers;
	using QVariantAnimation::isSignalConnected;
	bool cppgoBase_event(QEvent* __a0) { return QVariantAnimation::event(__a0); }
	void cppgoBase_updateCurrentTime(int __a0) { QVariantAnimation::updateCurrentTime(__a0); }
	void cppgoBase_updateState(int __a0, int __a1) { QVariantAnimation::updateState((QAbstractAnimation::State)__a0, (QAbstractAnimation::State)__a1); }
	void cppgoBase_updateCurrentValue(const QVariant& __a0) { QVariantAnimation::updateCurrentValue(__a0); }
	QVariant cppgoBase_interpolated(const QVariant& __a0, const QVariant& __a1, qreal __a2) { return QVariantAnimation::interpolated(__a0, __a1, __a2); }
	void cppgoBase_updateDirection(int __a0) { QVariantAnimation::updateDirection((QAbstractAnimation::Direction)__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QVariantAnimation::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QVariantAnimation::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QVariantAnimation::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QVariantAnimation::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QVariantAnimation::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QVariantAnimation::metaObject();
	}
	int duration() const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8434675112194400115LL, &__h); if (__h) return __r; }
		return QVariantAnimation::duration();
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QVariantAnimation::event(event);
	}
	void updateCurrentTime(int arg1) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, -1342418961870150968LL, arg1, &__h); if (__h) return; }
		QVariantAnimation::updateCurrentTime(arg1);
	}
	void updateState(QAbstractAnimation::State newState, QAbstractAnimation::State oldState) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_int_int) { ((void(*)(long long, long long, int, int, int*))cppgo_virt_void_int_int)(__goId, 7692482089134926569LL, (int)(newState), (int)(oldState), &__h); if (__h) return; }
		QVariantAnimation::updateState(newState, oldState);
	}
	void updateCurrentValue(const QVariant& value) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4717240701590708418LL, (void*)&value, &__h); if (__h) return; }
		QVariantAnimation::updateCurrentValue(value);
	}
	QVariant interpolated(const QVariant& from, const QVariant& to, qreal progress) const override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_voidp_voidp_voidp_unsignedlonglong) { void* __r = ((void*(*)(long long, long long, void*, void*, unsigned long long, int*))cppgo_virt_voidp_voidp_voidp_unsignedlonglong)(__goId, 4185842742461511100LL, (void*)&from, (void*)&to, __builtin_bit_cast(unsigned long long, (double)(progress)), &__h); if (__h) return *(QVariant*)__r; }
		return QVariantAnimation::interpolated(from, to, progress);
	}
	void updateDirection(QAbstractAnimation::Direction direction) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, -3231888725407197643LL, (int)(direction), &__h); if (__h) return; }
		QVariantAnimation::updateDirection(direction);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QVariantAnimation::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QVariantAnimation::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QVariantAnimation::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QVariantAnimation::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QVariantAnimation::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QVariantAnimation::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QVariantAnimation;
static std::mutex __connsMu_QVariantAnimation;

extern "C" {

void* QVariantAnimation_base(void* self) { return (void*)static_cast<QAbstractAnimation*>((QVariantAnimation*)self); }
void QVariantAnimation_chain(void* self, void** out) {
	QVariantAnimation* __p = (QVariantAnimation*)self;
	QAbstractAnimation* __b0 = static_cast<QAbstractAnimation*>(__p); out[0] = (void*)__b0;
	QObject* __b1 = static_cast<QObject*>(__b0); out[1] = (void*)__b1;
}
void QVariantAnimation__set_override(void* self, int idx) { ((QVariantAnimation__Go*)self)->__goOn[idx] = 1; }
void* QVariantAnimation_fromQObject(void* o) { return (void*)qobject_cast<QVariantAnimation*>((QObject*)o); }
void* QVariantAnimation_new(long long __goid) {
	return new QVariantAnimation__Go(__goid);
}
void* QVariantAnimation_new2(long long __goid, void* parent) {
	return new QVariantAnimation__Go(__goid, (QObject*)parent);
}
char* QVariantAnimation_tr(const char* s, long long* __out_len) {
	auto&& __r = QVariantAnimation::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QVariantAnimation_startValue(void* self) {
	auto&& __r = ((QVariantAnimation*)self)->startValue();
	return new QVariant(__r);
}
void QVariantAnimation_setStartValue(void* self, void* value) {
	((QVariantAnimation*)self)->setStartValue(*(QVariant*)value);
}
void* QVariantAnimation_endValue(void* self) {
	auto&& __r = ((QVariantAnimation*)self)->endValue();
	return new QVariant(__r);
}
void QVariantAnimation_setEndValue(void* self, void* value) {
	((QVariantAnimation*)self)->setEndValue(*(QVariant*)value);
}
void* QVariantAnimation_keyValueAt(void* self, double step) {
	auto&& __r = ((QVariantAnimation*)self)->keyValueAt(step);
	return new QVariant(__r);
}
void QVariantAnimation_setKeyValueAt(void* self, double step, void* value) {
	((QVariantAnimation*)self)->setKeyValueAt(step, *(QVariant*)value);
}
void* QVariantAnimation_currentValue(void* self) {
	auto&& __r = ((QVariantAnimation*)self)->currentValue();
	return new QVariant(__r);
}
void QVariantAnimation_setDuration(void* self, int msecs) {
	((QVariantAnimation*)self)->setDuration(msecs);
}
void* QVariantAnimation_easingCurve(void* self) {
	auto&& __r = ((QVariantAnimation*)self)->easingCurve();
	return new QEasingCurve(__r);
}
void QVariantAnimation_setEasingCurve(void* self, void* easing) {
	((QVariantAnimation*)self)->setEasingCurve(*(QEasingCurve*)easing);
}
void QVariantAnimation_valueChanged(void* self, void* value) {
	((QVariantAnimation*)self)->valueChanged(*(QVariant*)value);
}
void* QVariantAnimation_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QVariantAnimation::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QVariantAnimation_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QVariantAnimation::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QVariantAnimation_sender(void* self) {
	auto&& __r = ((QVariantAnimation__Go*)self)->sender();
	return (void*)(__r);
}
int QVariantAnimation_senderSignalIndex(void* self) {
	auto&& __r = ((QVariantAnimation__Go*)self)->senderSignalIndex();
	return __r;
}
int QVariantAnimation_receivers(void* self, const char* signal) {
	auto&& __r = ((QVariantAnimation__Go*)self)->receivers(signal);
	return __r;
}
int QVariantAnimation_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QVariantAnimation__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QVariantAnimation_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QVariantAnimation::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QVariantAnimation_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QVariantAnimation::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QVariantAnimation_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QVariantAnimation::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QVariantAnimation_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QVariantAnimation::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QVariantAnimation_metaObject_default(void* self) {
	auto&& __r = ((QVariantAnimation*)self)->QVariantAnimation::metaObject();
	return (void*)(__r);
}
int QVariantAnimation_duration_default(void* self) {
	auto&& __r = ((QVariantAnimation*)self)->QVariantAnimation::duration();
	return __r;
}
int QVariantAnimation_event_default(void* self, void* event) {
	auto&& __r = ((QVariantAnimation__Go*)self)->cppgoBase_event((QEvent*)event);
	return __r;
}
void QVariantAnimation_updateCurrentTime_default(void* self, int arg1) {
	((QVariantAnimation__Go*)self)->cppgoBase_updateCurrentTime(arg1);
}
void QVariantAnimation_updateState_default(void* self, int newState, int oldState) {
	((QVariantAnimation__Go*)self)->cppgoBase_updateState(newState, oldState);
}
void QVariantAnimation_updateCurrentValue_default(void* self, void* value) {
	((QVariantAnimation__Go*)self)->cppgoBase_updateCurrentValue(*(QVariant*)value);
}
void* QVariantAnimation_interpolated_default(void* self, void* from, void* to, double progress) {
	auto&& __r = ((QVariantAnimation__Go*)self)->cppgoBase_interpolated(*(QVariant*)from, *(QVariant*)to, progress);
	return new QVariant(__r);
}
void QVariantAnimation_updateDirection_default(void* self, int direction) {
	((QVariantAnimation__Go*)self)->cppgoBase_updateDirection(direction);
}
int QVariantAnimation_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QVariantAnimation*)self)->QVariantAnimation::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QVariantAnimation_timerEvent_default(void* self, void* event) {
	((QVariantAnimation__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QVariantAnimation_childEvent_default(void* self, void* event) {
	((QVariantAnimation__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QVariantAnimation_customEvent_default(void* self, void* event) {
	((QVariantAnimation__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QVariantAnimation_connectNotify_default(void* self, void* signal) {
	((QVariantAnimation__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QVariantAnimation_disconnectNotify_default(void* self, void* signal) {
	((QVariantAnimation__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QVariantAnimation_connect_valueChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QVariantAnimation*)self, &QVariantAnimation::valueChanged, (QVariantAnimation*)self, [__id](const QVariant& value){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&value);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QVariantAnimation);
	__conns_QVariantAnimation[__id] = __c;
}
void QVariantAnimation_connect_finished(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QVariantAnimation*)self, &QVariantAnimation::finished, (QVariantAnimation*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QVariantAnimation);
	__conns_QVariantAnimation[__id] = __c;
}
void QVariantAnimation_connect_stateChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QVariantAnimation*)self, &QVariantAnimation::stateChanged, (QVariantAnimation*)self, [__id](QAbstractAnimation::State newState, QAbstractAnimation::State oldState){
		if (cppgo_sig_void_int_int) ((void(*)(long long, int, int))cppgo_sig_void_int_int)(__id, (int)(newState), (int)(oldState));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QVariantAnimation);
	__conns_QVariantAnimation[__id] = __c;
}
void QVariantAnimation_connect_currentLoopChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QVariantAnimation*)self, &QVariantAnimation::currentLoopChanged, (QVariantAnimation*)self, [__id](int currentLoop){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, currentLoop);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QVariantAnimation);
	__conns_QVariantAnimation[__id] = __c;
}
void QVariantAnimation_connect_directionChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QVariantAnimation*)self, &QVariantAnimation::directionChanged, (QVariantAnimation*)self, [__id](QAbstractAnimation::Direction arg1){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, (int)(arg1));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QVariantAnimation);
	__conns_QVariantAnimation[__id] = __c;
}
void QVariantAnimation_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QVariantAnimation*)self, &QVariantAnimation::destroyed, (QVariantAnimation*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QVariantAnimation);
	__conns_QVariantAnimation[__id] = __c;
}
void QVariantAnimation_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QVariantAnimation*)self, &QVariantAnimation::destroyed, (QVariantAnimation*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QVariantAnimation);
	__conns_QVariantAnimation[__id] = __c;
}
void QVariantAnimation_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QVariantAnimation*)self, &QVariantAnimation::objectNameChanged, (QVariantAnimation*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QVariantAnimation);
	__conns_QVariantAnimation[__id] = __c;
}
void QVariantAnimation__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QVariantAnimation);
		auto __it = __conns_QVariantAnimation.find(id);
		if (__it == __conns_QVariantAnimation.end()) return;
		__c = __it->second;
		__conns_QVariantAnimation.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QVariantAnimation__Go_new(long long id, void* parent) { return new QVariantAnimation__Go(id, (QObject*)parent); }
void QVariantAnimation_delete(void* self) { delete (QVariantAnimation*)self; }

}

#include <QtCore/qsignalmapper.h>
#include <QtCore/qobject.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qvariant.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qsignalmapper.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QSignalMapper__Go : public QSignalMapper {
public:
	long long __goId;
	unsigned char __goOn[8] = {};
	template<typename... A> QSignalMapper__Go(long long id, A&&... a) : QSignalMapper(std::forward<A>(a)...), __goId(id) {}
	using QSignalMapper::sender;
	using QSignalMapper::senderSignalIndex;
	using QSignalMapper::receivers;
	using QSignalMapper::isSignalConnected;
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QSignalMapper::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QSignalMapper::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QSignalMapper::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QSignalMapper::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QSignalMapper::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QSignalMapper::metaObject();
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QSignalMapper::event(event);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QSignalMapper::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QSignalMapper::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QSignalMapper::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QSignalMapper::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QSignalMapper::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QSignalMapper::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QSignalMapper;
static std::mutex __connsMu_QSignalMapper;

extern "C" {

void* QSignalMapper_base(void* self) { return (void*)static_cast<QObject*>((QSignalMapper*)self); }
void QSignalMapper__set_override(void* self, int idx) { ((QSignalMapper__Go*)self)->__goOn[idx] = 1; }
void* QSignalMapper_fromQObject(void* o) { return (void*)qobject_cast<QSignalMapper*>((QObject*)o); }
void* QSignalMapper_new(long long __goid) {
	return new QSignalMapper__Go(__goid);
}
void* QSignalMapper_new2(long long __goid, void* parent) {
	return new QSignalMapper__Go(__goid, (QObject*)parent);
}
char* QSignalMapper_tr(const char* s, long long* __out_len) {
	auto&& __r = QSignalMapper::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QSignalMapper_setMapping(void* self, void* sender, int id) {
	((QSignalMapper*)self)->setMapping((QObject*)sender, id);
}
void QSignalMapper_setMapping2(void* self, void* sender, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	((QSignalMapper*)self)->setMapping((QObject*)sender, text_qstr);
}
void QSignalMapper_setMapping3(void* self, void* sender, void* object) {
	((QSignalMapper*)self)->setMapping((QObject*)sender, (QObject*)object);
}
void QSignalMapper_removeMappings(void* self, void* sender) {
	((QSignalMapper*)self)->removeMappings((QObject*)sender);
}
void* QSignalMapper_mapping(void* self, int id) {
	auto&& __r = ((QSignalMapper*)self)->mapping(id);
	return (void*)(__r);
}
void* QSignalMapper_mapping2(void* self, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = ((QSignalMapper*)self)->mapping(text_qstr);
	return (void*)(__r);
}
void* QSignalMapper_mapping3(void* self, void* object) {
	auto&& __r = ((QSignalMapper*)self)->mapping((QObject*)object);
	return (void*)(__r);
}
void QSignalMapper_mappedInt(void* self, int arg1) {
	((QSignalMapper*)self)->mappedInt(arg1);
}
void QSignalMapper_mappedString(void* self, const char* arg1, long long arg1_len) {
	QString arg1_qstr = QString::fromUtf8(arg1, arg1_len);
	((QSignalMapper*)self)->mappedString(arg1_qstr);
}
void QSignalMapper_mappedObject(void* self, void* arg1) {
	((QSignalMapper*)self)->mappedObject((QObject*)arg1);
}
void QSignalMapper_map(void* self) {
	((QSignalMapper*)self)->map();
}
void QSignalMapper_map2(void* self, void* sender) {
	((QSignalMapper*)self)->map((QObject*)sender);
}
void* QSignalMapper_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QSignalMapper::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QSignalMapper_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QSignalMapper::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QSignalMapper_sender(void* self) {
	auto&& __r = ((QSignalMapper__Go*)self)->sender();
	return (void*)(__r);
}
int QSignalMapper_senderSignalIndex(void* self) {
	auto&& __r = ((QSignalMapper__Go*)self)->senderSignalIndex();
	return __r;
}
int QSignalMapper_receivers(void* self, const char* signal) {
	auto&& __r = ((QSignalMapper__Go*)self)->receivers(signal);
	return __r;
}
int QSignalMapper_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QSignalMapper__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QSignalMapper_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QSignalMapper::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QSignalMapper_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QSignalMapper::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QSignalMapper_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QSignalMapper::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QSignalMapper_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QSignalMapper::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QSignalMapper_metaObject_default(void* self) {
	auto&& __r = ((QSignalMapper*)self)->QSignalMapper::metaObject();
	return (void*)(__r);
}
int QSignalMapper_event_default(void* self, void* event) {
	auto&& __r = ((QSignalMapper*)self)->QSignalMapper::event((QEvent*)event);
	return __r;
}
int QSignalMapper_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QSignalMapper*)self)->QSignalMapper::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QSignalMapper_timerEvent_default(void* self, void* event) {
	((QSignalMapper__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QSignalMapper_childEvent_default(void* self, void* event) {
	((QSignalMapper__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QSignalMapper_customEvent_default(void* self, void* event) {
	((QSignalMapper__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QSignalMapper_connectNotify_default(void* self, void* signal) {
	((QSignalMapper__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QSignalMapper_disconnectNotify_default(void* self, void* signal) {
	((QSignalMapper__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QSignalMapper_connect_mappedInt(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSignalMapper*)self, &QSignalMapper::mappedInt, (QSignalMapper*)self, [__id](int arg1){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSignalMapper);
	__conns_QSignalMapper[__id] = __c;
}
void QSignalMapper_connect_mappedString(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSignalMapper*)self, &QSignalMapper::mappedString, (QSignalMapper*)self, [__id](const QString& arg1){
		QByteArray arg1_ba = arg1.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)arg1_ba.constData(), arg1_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSignalMapper);
	__conns_QSignalMapper[__id] = __c;
}
void QSignalMapper_connect_mappedObject(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSignalMapper*)self, &QSignalMapper::mappedObject, (QSignalMapper*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSignalMapper);
	__conns_QSignalMapper[__id] = __c;
}
void QSignalMapper_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSignalMapper*)self, &QSignalMapper::destroyed, (QSignalMapper*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSignalMapper);
	__conns_QSignalMapper[__id] = __c;
}
void QSignalMapper_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSignalMapper*)self, &QSignalMapper::destroyed, (QSignalMapper*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSignalMapper);
	__conns_QSignalMapper[__id] = __c;
}
void QSignalMapper_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSignalMapper*)self, &QSignalMapper::objectNameChanged, (QSignalMapper*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSignalMapper);
	__conns_QSignalMapper[__id] = __c;
}
void QSignalMapper__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QSignalMapper);
		auto __it = __conns_QSignalMapper.find(id);
		if (__it == __conns_QSignalMapper.end()) return;
		__c = __it->second;
		__conns_QSignalMapper.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QSignalMapper__Go_new(long long id, void* parent) { return new QSignalMapper__Go(id, (QObject*)parent); }
void QSignalMapper_delete(void* self) { delete (QSignalMapper*)self; }

}

#include <QtCore/qxmlstream.h>
#include <QtCore/qiodevice.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qstring.h>
#include <QtCore/qstringview.h>
#include "gen_qxmlstreamreader.h"

extern "C" {

void* QXmlStreamReader_new() {
	return new QXmlStreamReader();
}
void* QXmlStreamReader_new2(void* device) {
	return new QXmlStreamReader((QIODevice*)device);
}
void* QXmlStreamReader_new3(const char* data, long long data_len) {
	QAnyStringView data_qstr = QAnyStringView(data, data_len);
	return new QXmlStreamReader(data_qstr);
}
void QXmlStreamReader_setDevice(void* self, void* device) {
	((QXmlStreamReader*)self)->setDevice((QIODevice*)device);
}
void* QXmlStreamReader_device(void* self) {
	auto&& __r = ((QXmlStreamReader*)self)->device();
	return (void*)(__r);
}
void QXmlStreamReader_addData(void* self, const char* data, long long data_len) {
	QAnyStringView data_qstr = QAnyStringView(data, data_len);
	((QXmlStreamReader*)self)->addData(data_qstr);
}
void QXmlStreamReader_clear(void* self) {
	((QXmlStreamReader*)self)->clear();
}
int QXmlStreamReader_atEnd(void* self) {
	auto&& __r = ((QXmlStreamReader*)self)->atEnd();
	return __r;
}
int QXmlStreamReader_readNext(void* self) {
	auto&& __r = ((QXmlStreamReader*)self)->readNext();
	return (int)(__r);
}
int QXmlStreamReader_readNextStartElement(void* self) {
	auto&& __r = ((QXmlStreamReader*)self)->readNextStartElement();
	return __r;
}
void QXmlStreamReader_skipCurrentElement(void* self) {
	((QXmlStreamReader*)self)->skipCurrentElement();
}
char* QXmlStreamReader_readRawInnerData(void* self, long long* __out_len) {
	auto&& __r = ((QXmlStreamReader*)self)->readRawInnerData();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QXmlStreamReader_tokenType(void* self) {
	auto&& __r = ((QXmlStreamReader*)self)->tokenType();
	return (int)(__r);
}
char* QXmlStreamReader_tokenString(void* self, long long* __out_len) {
	auto&& __r = ((QXmlStreamReader*)self)->tokenString();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QXmlStreamReader_setNamespaceProcessing(void* self, int arg1) {
	((QXmlStreamReader*)self)->setNamespaceProcessing(arg1);
}
int QXmlStreamReader_namespaceProcessing(void* self) {
	auto&& __r = ((QXmlStreamReader*)self)->namespaceProcessing();
	return __r;
}
int QXmlStreamReader_isStartDocument(void* self) {
	auto&& __r = ((QXmlStreamReader*)self)->isStartDocument();
	return __r;
}
int QXmlStreamReader_isEndDocument(void* self) {
	auto&& __r = ((QXmlStreamReader*)self)->isEndDocument();
	return __r;
}
int QXmlStreamReader_isStartElement(void* self) {
	auto&& __r = ((QXmlStreamReader*)self)->isStartElement();
	return __r;
}
int QXmlStreamReader_isEndElement(void* self) {
	auto&& __r = ((QXmlStreamReader*)self)->isEndElement();
	return __r;
}
int QXmlStreamReader_isCharacters(void* self) {
	auto&& __r = ((QXmlStreamReader*)self)->isCharacters();
	return __r;
}
int QXmlStreamReader_isWhitespace(void* self) {
	auto&& __r = ((QXmlStreamReader*)self)->isWhitespace();
	return __r;
}
int QXmlStreamReader_isCDATA(void* self) {
	auto&& __r = ((QXmlStreamReader*)self)->isCDATA();
	return __r;
}
int QXmlStreamReader_isComment(void* self) {
	auto&& __r = ((QXmlStreamReader*)self)->isComment();
	return __r;
}
int QXmlStreamReader_isDTD(void* self) {
	auto&& __r = ((QXmlStreamReader*)self)->isDTD();
	return __r;
}
int QXmlStreamReader_isEntityReference(void* self) {
	auto&& __r = ((QXmlStreamReader*)self)->isEntityReference();
	return __r;
}
int QXmlStreamReader_isProcessingInstruction(void* self) {
	auto&& __r = ((QXmlStreamReader*)self)->isProcessingInstruction();
	return __r;
}
int QXmlStreamReader_isStandaloneDocument(void* self) {
	auto&& __r = ((QXmlStreamReader*)self)->isStandaloneDocument();
	return __r;
}
int QXmlStreamReader_hasStandaloneDeclaration(void* self) {
	auto&& __r = ((QXmlStreamReader*)self)->hasStandaloneDeclaration();
	return __r;
}
void* QXmlStreamReader_documentVersion(void* self) {
	auto&& __r = ((QXmlStreamReader*)self)->documentVersion();
	return new QStringView(__r);
}
void* QXmlStreamReader_documentEncoding(void* self) {
	auto&& __r = ((QXmlStreamReader*)self)->documentEncoding();
	return new QStringView(__r);
}
long long QXmlStreamReader_lineNumber(void* self) {
	auto&& __r = ((QXmlStreamReader*)self)->lineNumber();
	return __r;
}
long long QXmlStreamReader_columnNumber(void* self) {
	auto&& __r = ((QXmlStreamReader*)self)->columnNumber();
	return __r;
}
long long QXmlStreamReader_characterOffset(void* self) {
	auto&& __r = ((QXmlStreamReader*)self)->characterOffset();
	return __r;
}
void* QXmlStreamReader_attributes(void* self) {
	auto&& __r = ((QXmlStreamReader*)self)->attributes();
	return new QXmlStreamAttributes(__r);
}
char* QXmlStreamReader_readElementText(void* self, long long* __out_len) {
	auto&& __r = ((QXmlStreamReader*)self)->readElementText();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QXmlStreamReader_name(void* self) {
	auto&& __r = ((QXmlStreamReader*)self)->name();
	return new QStringView(__r);
}
void* QXmlStreamReader_namespaceUri(void* self) {
	auto&& __r = ((QXmlStreamReader*)self)->namespaceUri();
	return new QStringView(__r);
}
void* QXmlStreamReader_qualifiedName(void* self) {
	auto&& __r = ((QXmlStreamReader*)self)->qualifiedName();
	return new QStringView(__r);
}
void* QXmlStreamReader_prefix(void* self) {
	auto&& __r = ((QXmlStreamReader*)self)->prefix();
	return new QStringView(__r);
}
void* QXmlStreamReader_processingInstructionTarget(void* self) {
	auto&& __r = ((QXmlStreamReader*)self)->processingInstructionTarget();
	return new QStringView(__r);
}
void* QXmlStreamReader_processingInstructionData(void* self) {
	auto&& __r = ((QXmlStreamReader*)self)->processingInstructionData();
	return new QStringView(__r);
}
void* QXmlStreamReader_text(void* self) {
	auto&& __r = ((QXmlStreamReader*)self)->text();
	return new QStringView(__r);
}
void** QXmlStreamReader_namespaceDeclarations(void* self, long long* __out_len) {
	auto&& __r = ((QXmlStreamReader*)self)->namespaceDeclarations();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QXmlStreamNamespaceDeclaration(__e);
	return __arr;
}
void QXmlStreamReader_addExtraNamespaceDeclaration(void* self, void* extraNamespaceDeclaraction) {
	((QXmlStreamReader*)self)->addExtraNamespaceDeclaration(*(QXmlStreamNamespaceDeclaration*)extraNamespaceDeclaraction);
}
void QXmlStreamReader_addExtraNamespaceDeclarations(void* self, void** extraNamespaceDeclaractions, long long extraNamespaceDeclaractions_len) {
	QList<QXmlStreamNamespaceDeclaration> extraNamespaceDeclaractions_list; extraNamespaceDeclaractions_list.reserve(extraNamespaceDeclaractions_len);
	for (long long __i = 0; __i < extraNamespaceDeclaractions_len; __i++) extraNamespaceDeclaractions_list.append(*(QXmlStreamNamespaceDeclaration*)extraNamespaceDeclaractions[__i]);
	((QXmlStreamReader*)self)->addExtraNamespaceDeclarations(extraNamespaceDeclaractions_list);
}
void** QXmlStreamReader_notationDeclarations(void* self, long long* __out_len) {
	auto&& __r = ((QXmlStreamReader*)self)->notationDeclarations();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QXmlStreamNotationDeclaration(__e);
	return __arr;
}
void** QXmlStreamReader_entityDeclarations(void* self, long long* __out_len) {
	auto&& __r = ((QXmlStreamReader*)self)->entityDeclarations();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QXmlStreamEntityDeclaration(__e);
	return __arr;
}
void* QXmlStreamReader_dtdName(void* self) {
	auto&& __r = ((QXmlStreamReader*)self)->dtdName();
	return new QStringView(__r);
}
void* QXmlStreamReader_dtdPublicId(void* self) {
	auto&& __r = ((QXmlStreamReader*)self)->dtdPublicId();
	return new QStringView(__r);
}
void* QXmlStreamReader_dtdSystemId(void* self) {
	auto&& __r = ((QXmlStreamReader*)self)->dtdSystemId();
	return new QStringView(__r);
}
int QXmlStreamReader_entityExpansionLimit(void* self) {
	auto&& __r = ((QXmlStreamReader*)self)->entityExpansionLimit();
	return __r;
}
void QXmlStreamReader_setEntityExpansionLimit(void* self, int limit) {
	((QXmlStreamReader*)self)->setEntityExpansionLimit(limit);
}
void QXmlStreamReader_raiseError(void* self) {
	((QXmlStreamReader*)self)->raiseError();
}
char* QXmlStreamReader_errorString(void* self, long long* __out_len) {
	auto&& __r = ((QXmlStreamReader*)self)->errorString();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QXmlStreamReader_error(void* self) {
	auto&& __r = ((QXmlStreamReader*)self)->error();
	return (int)(__r);
}
int QXmlStreamReader_hasError(void* self) {
	auto&& __r = ((QXmlStreamReader*)self)->hasError();
	return __r;
}
void QXmlStreamReader_setEntityResolver(void* self, void* resolver) {
	((QXmlStreamReader*)self)->setEntityResolver((QXmlStreamEntityResolver*)resolver);
}
void* QXmlStreamReader_entityResolver(void* self) {
	auto&& __r = ((QXmlStreamReader*)self)->entityResolver();
	return (void*)(__r);
}
char* QXmlStreamReader_readElementText2(void* self, int behaviour, long long* __out_len) {
	auto&& __r = ((QXmlStreamReader*)self)->readElementText((QXmlStreamReader::ReadElementTextBehaviour)(behaviour));
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QXmlStreamReader_raiseError2(void* self, const char* message, long long message_len) {
	QString message_qstr = QString::fromUtf8(message, message_len);
	((QXmlStreamReader*)self)->raiseError(message_qstr);
}
void QXmlStreamReader_delete(void* self) { delete (QXmlStreamReader*)self; }

}

#include <QtGui/qkeysequence.h>
#include <QtCore/qstring.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qvariant.h>
#include "gen_qkeysequence.h"

extern "C" {

void* QKeySequence_new() {
	return new QKeySequence();
}
void* QKeySequence_new2(const char* key, long long key_len) {
	QString key_qstr = QString::fromUtf8(key, key_len);
	return new QKeySequence(key_qstr);
}
void* QKeySequence_new3(int k1) {
	return new QKeySequence(k1);
}
void* QKeySequence_new4(void* k1) {
	return new QKeySequence(*(QKeyCombination*)k1);
}
void* QKeySequence_new5(void* ks) {
	return new QKeySequence(*(QKeySequence*)ks);
}
void* QKeySequence_new6(int key) {
	return new QKeySequence((QKeySequence::StandardKey)(key));
}
void* QKeySequence_new7(const char* key, long long key_len, int format) {
	QString key_qstr = QString::fromUtf8(key, key_len);
	return new QKeySequence(key_qstr, (QKeySequence::SequenceFormat)(format));
}
void* QKeySequence_new8(int k1, int k2) {
	return new QKeySequence(k1, k2);
}
void* QKeySequence_new9(int k1, int k2, int k3) {
	return new QKeySequence(k1, k2, k3);
}
void* QKeySequence_new10(int k1, int k2, int k3, int k4) {
	return new QKeySequence(k1, k2, k3, k4);
}
void* QKeySequence_new11(void* k1, void* k2) {
	return new QKeySequence(*(QKeyCombination*)k1, *(QKeyCombination*)k2);
}
void* QKeySequence_new12(void* k1, void* k2, void* k3) {
	return new QKeySequence(*(QKeyCombination*)k1, *(QKeyCombination*)k2, *(QKeyCombination*)k3);
}
void* QKeySequence_new13(void* k1, void* k2, void* k3, void* k4) {
	return new QKeySequence(*(QKeyCombination*)k1, *(QKeyCombination*)k2, *(QKeyCombination*)k3, *(QKeyCombination*)k4);
}
int QKeySequence_count(void* self) {
	auto&& __r = ((QKeySequence*)self)->count();
	return __r;
}
int QKeySequence_isEmpty(void* self) {
	auto&& __r = ((QKeySequence*)self)->isEmpty();
	return __r;
}
char* QKeySequence_toString(void* self, long long* __out_len) {
	auto&& __r = ((QKeySequence*)self)->toString();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QKeySequence_fromString(const char* str, long long str_len) {
	QString str_qstr = QString::fromUtf8(str, str_len);
	auto&& __r = QKeySequence::fromString(str_qstr);
	return new QKeySequence(__r);
}
void** QKeySequence_listFromString(const char* str, long long str_len, long long* __out_len) {
	QString str_qstr = QString::fromUtf8(str, str_len);
	auto&& __r = QKeySequence::listFromString(str_qstr);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QKeySequence(__e);
	return __arr;
}
char* QKeySequence_listToString(void** list, long long list_len, long long* __out_len) {
	QList<QKeySequence> list_list; list_list.reserve(list_len);
	for (long long __i = 0; __i < list_len; __i++) list_list.append(*(QKeySequence*)list[__i]);
	auto&& __r = QKeySequence::listToString(list_list);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QKeySequence_matches(void* self, void* seq) {
	auto&& __r = ((QKeySequence*)self)->matches(*(QKeySequence*)seq);
	return (int)(__r);
}
void* QKeySequence_mnemonic(const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = QKeySequence::mnemonic(text_qstr);
	return new QKeySequence(__r);
}
void** QKeySequence_keyBindings(int key, long long* __out_len) {
	auto&& __r = QKeySequence::keyBindings((QKeySequence::StandardKey)(key));
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QKeySequence(__e);
	return __arr;
}
void* QKeySequence_ToQVariant(void* self) {
	auto&& __r = static_cast<QVariant>(*(QKeySequence*)self);
	return new QVariant(__r);
}
void* QKeySequence_OperatorSubscript(void* self, unsigned int i) {
	auto&& __r = ((QKeySequence*)self)->operator[](i);
	return new QKeyCombination(__r);
}
void* QKeySequence_OperatorAssign(void* self, void* other) {
	auto&& __r = ((QKeySequence*)self)->operator=(*(QKeySequence*)other);
	return (void*)&(__r);
}
void QKeySequence_swap(void* self, void* other) {
	((QKeySequence*)self)->swap(*(QKeySequence*)other);
}
int QKeySequence_OperatorEqual(void* self, void* other) {
	auto&& __r = ((QKeySequence*)self)->operator==(*(QKeySequence*)other);
	return __r;
}
int QKeySequence_OperatorNotEqual(void* self, void* other) {
	auto&& __r = ((QKeySequence*)self)->operator!=(*(QKeySequence*)other);
	return __r;
}
int QKeySequence_OperatorLesser(void* self, void* ks) {
	auto&& __r = ((QKeySequence*)self)->operator<(*(QKeySequence*)ks);
	return __r;
}
int QKeySequence_OperatorGreater(void* self, void* other) {
	auto&& __r = ((QKeySequence*)self)->operator>(*(QKeySequence*)other);
	return __r;
}
int QKeySequence_OperatorLesserOrEqual(void* self, void* other) {
	auto&& __r = ((QKeySequence*)self)->operator<=(*(QKeySequence*)other);
	return __r;
}
int QKeySequence_OperatorGreaterOrEqual(void* self, void* other) {
	auto&& __r = ((QKeySequence*)self)->operator>=(*(QKeySequence*)other);
	return __r;
}
int QKeySequence_isDetached(void* self) {
	auto&& __r = ((QKeySequence*)self)->isDetached();
	return __r;
}
char* QKeySequence_toString2(void* self, int format, long long* __out_len) {
	auto&& __r = ((QKeySequence*)self)->toString((QKeySequence::SequenceFormat)(format));
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QKeySequence_fromString2(const char* str, long long str_len, int format) {
	QString str_qstr = QString::fromUtf8(str, str_len);
	auto&& __r = QKeySequence::fromString(str_qstr, (QKeySequence::SequenceFormat)(format));
	return new QKeySequence(__r);
}
void** QKeySequence_listFromString2(const char* str, long long str_len, int format, long long* __out_len) {
	QString str_qstr = QString::fromUtf8(str, str_len);
	auto&& __r = QKeySequence::listFromString(str_qstr, (QKeySequence::SequenceFormat)(format));
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QKeySequence(__e);
	return __arr;
}
char* QKeySequence_listToString2(void** list, long long list_len, int format, long long* __out_len) {
	QList<QKeySequence> list_list; list_list.reserve(list_len);
	for (long long __i = 0; __i < list_len; __i++) list_list.append(*(QKeySequence*)list[__i]);
	auto&& __r = QKeySequence::listToString(list_list, (QKeySequence::SequenceFormat)(format));
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QKeySequence_toQVariant(void* self) {
	auto&& __r = QVariant::fromValue(*(QKeySequence*)self);
	return new QVariant(__r);
}
void QKeySequence_delete(void* self) { delete (QKeySequence*)self; }

}

#include <QtGui/qevent.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qvariant.h>
#include "gen_qinputmethodqueryevent.h"

extern "C" {

void* QInputMethodQueryEvent_base(void* self) { return (void*)static_cast<QEvent*>((QInputMethodQueryEvent*)self); }
void* QInputMethodQueryEvent_new(int queries) {
	return new QInputMethodQueryEvent(QFlags<Qt::InputMethodQuery>(queries));
}
int QInputMethodQueryEvent_queries(void* self) {
	auto&& __r = ((QInputMethodQueryEvent*)self)->queries();
	return (int)(__r);
}
void QInputMethodQueryEvent_setValue(void* self, int query, void* value) {
	((QInputMethodQueryEvent*)self)->setValue((Qt::InputMethodQuery)(query), *(QVariant*)value);
}
void* QInputMethodQueryEvent_value(void* self, int query) {
	auto&& __r = ((QInputMethodQueryEvent*)self)->value((Qt::InputMethodQuery)(query));
	return new QVariant(__r);
}
int QInputMethodQueryEvent_registerEventType() {
	auto&& __r = QInputMethodQueryEvent::registerEventType();
	return __r;
}
int QInputMethodQueryEvent_registerEventType2(int hint) {
	auto&& __r = QInputMethodQueryEvent::registerEventType(hint);
	return __r;
}
void QInputMethodQueryEvent_delete(void* self) { delete (QInputMethodQueryEvent*)self; }

}

#include <QtGui/qtextdocument.h>
#include "gen_qabstractundoitem.h"
#include "gen_purego_cb.h"

class QAbstractUndoItem__Go : public QAbstractUndoItem {
public:
	long long __goId;
	unsigned char __goOn[2] = {};
	template<typename... A> QAbstractUndoItem__Go(long long id, A&&... a) : QAbstractUndoItem(std::forward<A>(a)...), __goId(id) {}
	void undo() override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 7259214242580717359LL, &__h); if (__h) return; }
		return;
	}
	void redo() override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 5544063219051948393LL, &__h); if (__h) return; }
		return;
	}
};

extern "C" {

void QAbstractUndoItem__set_override(void* self, int idx) { ((QAbstractUndoItem__Go*)self)->__goOn[idx] = 1; }
void QAbstractUndoItem_undo(void* self) {
	((QAbstractUndoItem*)self)->undo();
}
void QAbstractUndoItem_redo(void* self) {
	((QAbstractUndoItem*)self)->redo();
}
void* QAbstractUndoItem_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QAbstractUndoItem*)self)->operator=(*(QAbstractUndoItem*)arg1);
	return (void*)&(__r);
}
void* QAbstractUndoItem__Go_new(long long id) { return new QAbstractUndoItem__Go(id); }
void QAbstractUndoItem_delete(void* self) { delete (QAbstractUndoItem*)self; }

}

#include <QtGui/qtextlayout.h>
#include <QtGui/qtextformat.h>
#include "gen_qtextlayout__formatrange.h"

extern "C" {

void* QTextLayout__FormatRange_new() {
	return new QTextLayout::FormatRange();
}
void* QTextLayout__FormatRange_new2(void* arg1) {
	return new QTextLayout::FormatRange(*(QTextLayout::FormatRange*)arg1);
}
int QTextLayout__FormatRange_OperatorEqual(void* self, void* rhs) {
	auto&& __r = (*(QTextLayout::FormatRange*)self) == (*(QTextLayout::FormatRange*)rhs);
	return __r;
}
int QTextLayout__FormatRange_OperatorNotEqual(void* self, void* rhs) {
	auto&& __r = (*(QTextLayout::FormatRange*)self) != (*(QTextLayout::FormatRange*)rhs);
	return __r;
}
int QTextLayout__FormatRange_start_get(void* self) {
	auto __r = ((QTextLayout::FormatRange*)self)->start;
	return __r;
}
void QTextLayout__FormatRange_start_set(void* self, int start) {
	((QTextLayout::FormatRange*)self)->start = start;
}
int QTextLayout__FormatRange_length_get(void* self) {
	auto __r = ((QTextLayout::FormatRange*)self)->length;
	return __r;
}
void QTextLayout__FormatRange_length_set(void* self, int length) {
	((QTextLayout::FormatRange*)self)->length = length;
}
void* QTextLayout__FormatRange_format_get(void* self) {
	auto __r = ((QTextLayout::FormatRange*)self)->format;
	return new QTextCharFormat(__r);
}
void QTextLayout__FormatRange_format_set(void* self, void* format) {
	((QTextLayout::FormatRange*)self)->format = *(QTextCharFormat*)format;
}
void QTextLayout__FormatRange_delete(void* self) { delete (QTextLayout::FormatRange*)self; }

}

#include <QtGui/qaccessible.h>
#include <QtCore/qobject.h>
#include <QtGui/qaccessible_base.h>
#include "gen_qaccessibletextcursorevent.h"

extern "C" {

void* QAccessibleTextCursorEvent_base(void* self) { return (void*)static_cast<QAccessibleEvent*>((QAccessibleTextCursorEvent*)self); }
void* QAccessibleTextCursorEvent_new(void* obj, int cursorPos) {
	return new QAccessibleTextCursorEvent((QObject*)obj, cursorPos);
}
void* QAccessibleTextCursorEvent_new2(void* iface, int cursorPos) {
	return new QAccessibleTextCursorEvent((QAccessibleInterface*)iface, cursorPos);
}
void QAccessibleTextCursorEvent_setCursorPosition(void* self, int position) {
	((QAccessibleTextCursorEvent*)self)->setCursorPosition(position);
}
int QAccessibleTextCursorEvent_cursorPosition(void* self) {
	auto&& __r = ((QAccessibleTextCursorEvent*)self)->cursorPosition();
	return __r;
}
void QAccessibleTextCursorEvent_delete(void* self) { delete (QAccessibleTextCursorEvent*)self; }

}

#include <QtGui/qcolorspace.h>
#include <QtCore/qpoint.h>
#include "gen_qcolorspace__primarypoints.h"

extern "C" {

void* QColorSpace__PrimaryPoints_fromPrimaries(int primaries) {
	auto&& __r = QColorSpace::PrimaryPoints::fromPrimaries((QColorSpace::Primaries)(primaries));
	return new QColorSpace::PrimaryPoints(__r);
}
int QColorSpace__PrimaryPoints_isValid(void* self) {
	auto&& __r = ((QColorSpace::PrimaryPoints*)self)->isValid();
	return __r;
}
void* QColorSpace__PrimaryPoints_whitePoint_get(void* self) {
	auto __r = ((QColorSpace::PrimaryPoints*)self)->whitePoint;
	return new QPointF(__r);
}
void QColorSpace__PrimaryPoints_whitePoint_set(void* self, void* whitePoint) {
	((QColorSpace::PrimaryPoints*)self)->whitePoint = *(QPointF*)whitePoint;
}
void* QColorSpace__PrimaryPoints_redPoint_get(void* self) {
	auto __r = ((QColorSpace::PrimaryPoints*)self)->redPoint;
	return new QPointF(__r);
}
void QColorSpace__PrimaryPoints_redPoint_set(void* self, void* redPoint) {
	((QColorSpace::PrimaryPoints*)self)->redPoint = *(QPointF*)redPoint;
}
void* QColorSpace__PrimaryPoints_greenPoint_get(void* self) {
	auto __r = ((QColorSpace::PrimaryPoints*)self)->greenPoint;
	return new QPointF(__r);
}
void QColorSpace__PrimaryPoints_greenPoint_set(void* self, void* greenPoint) {
	((QColorSpace::PrimaryPoints*)self)->greenPoint = *(QPointF*)greenPoint;
}
void* QColorSpace__PrimaryPoints_bluePoint_get(void* self) {
	auto __r = ((QColorSpace::PrimaryPoints*)self)->bluePoint;
	return new QPointF(__r);
}
void QColorSpace__PrimaryPoints_bluePoint_set(void* self, void* bluePoint) {
	((QColorSpace::PrimaryPoints*)self)->bluePoint = *(QPointF*)bluePoint;
}
void QColorSpace__PrimaryPoints_delete(void* self) { delete (QColorSpace::PrimaryPoints*)self; }

}

#include <QtGui/qpagesize.h>
#include <QtCore/qsize.h>
#include <QtCore/qstring.h>
#include <QtCore/qrect.h>
#include "gen_qpagesize.h"

extern "C" {

void* QPageSize_new() {
	return new QPageSize();
}
void* QPageSize_new2(int pageSizeId) {
	return new QPageSize((QPageSize::PageSizeId)(pageSizeId));
}
void* QPageSize_new3(void* pointSize) {
	return new QPageSize(*(QSize*)pointSize);
}
void* QPageSize_new4(void* size, int units) {
	return new QPageSize(*(QSizeF*)size, (QPageSize::Unit)(units));
}
void* QPageSize_new5(void* other) {
	return new QPageSize(*(QPageSize*)other);
}
void* QPageSize_new6(void* pointSize, const char* name, long long name_len) {
	QString name_qstr = QString::fromUtf8(name, name_len);
	return new QPageSize(*(QSize*)pointSize, name_qstr);
}
void* QPageSize_new7(void* pointSize, const char* name, long long name_len, int matchPolicy) {
	QString name_qstr = QString::fromUtf8(name, name_len);
	return new QPageSize(*(QSize*)pointSize, name_qstr, (QPageSize::SizeMatchPolicy)(matchPolicy));
}
void* QPageSize_new8(void* size, int units, const char* name, long long name_len) {
	QString name_qstr = QString::fromUtf8(name, name_len);
	return new QPageSize(*(QSizeF*)size, (QPageSize::Unit)(units), name_qstr);
}
void* QPageSize_new9(void* size, int units, const char* name, long long name_len, int matchPolicy) {
	QString name_qstr = QString::fromUtf8(name, name_len);
	return new QPageSize(*(QSizeF*)size, (QPageSize::Unit)(units), name_qstr, (QPageSize::SizeMatchPolicy)(matchPolicy));
}
void* QPageSize_OperatorAssign(void* self, void* other) {
	auto&& __r = ((QPageSize*)self)->operator=(*(QPageSize*)other);
	return (void*)&(__r);
}
void QPageSize_swap(void* self, void* other) {
	((QPageSize*)self)->swap(*(QPageSize*)other);
}
int QPageSize_isEquivalentTo(void* self, void* other) {
	auto&& __r = ((QPageSize*)self)->isEquivalentTo(*(QPageSize*)other);
	return __r;
}
int QPageSize_isValid(void* self) {
	auto&& __r = ((QPageSize*)self)->isValid();
	return __r;
}
char* QPageSize_key(void* self, long long* __out_len) {
	auto&& __r = ((QPageSize*)self)->key();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QPageSize_name(void* self, long long* __out_len) {
	auto&& __r = ((QPageSize*)self)->name();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QPageSize_id(void* self) {
	auto&& __r = ((QPageSize*)self)->id();
	return (int)(__r);
}
int QPageSize_windowsId(void* self) {
	auto&& __r = ((QPageSize*)self)->windowsId();
	return __r;
}
void* QPageSize_definitionSize(void* self) {
	auto&& __r = ((QPageSize*)self)->definitionSize();
	return new QSizeF(__r);
}
int QPageSize_definitionUnits(void* self) {
	auto&& __r = ((QPageSize*)self)->definitionUnits();
	return (int)(__r);
}
void* QPageSize_size(void* self, int units) {
	auto&& __r = ((QPageSize*)self)->size((QPageSize::Unit)(units));
	return new QSizeF(__r);
}
void* QPageSize_sizePoints(void* self) {
	auto&& __r = ((QPageSize*)self)->sizePoints();
	return new QSize(__r);
}
void* QPageSize_sizePixels(void* self, int resolution) {
	auto&& __r = ((QPageSize*)self)->sizePixels(resolution);
	return new QSize(__r);
}
void* QPageSize_rect(void* self, int units) {
	auto&& __r = ((QPageSize*)self)->rect((QPageSize::Unit)(units));
	return new QRectF(__r);
}
void* QPageSize_rectPoints(void* self) {
	auto&& __r = ((QPageSize*)self)->rectPoints();
	return new QRect(__r);
}
void* QPageSize_rectPixels(void* self, int resolution) {
	auto&& __r = ((QPageSize*)self)->rectPixels(resolution);
	return new QRect(__r);
}
char* QPageSize_key2(int pageSizeId, long long* __out_len) {
	auto&& __r = QPageSize::key((QPageSize::PageSizeId)(pageSizeId));
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QPageSize_name2(int pageSizeId, long long* __out_len) {
	auto&& __r = QPageSize::name((QPageSize::PageSizeId)(pageSizeId));
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QPageSize_id2(void* pointSize) {
	auto&& __r = QPageSize::id(*(QSize*)pointSize);
	return (int)(__r);
}
int QPageSize_id3(void* size, int units) {
	auto&& __r = QPageSize::id(*(QSizeF*)size, (QPageSize::Unit)(units));
	return (int)(__r);
}
int QPageSize_id4(int windowsId) {
	auto&& __r = QPageSize::id(windowsId);
	return (int)(__r);
}
int QPageSize_windowsId2(int pageSizeId) {
	auto&& __r = QPageSize::windowsId((QPageSize::PageSizeId)(pageSizeId));
	return __r;
}
void* QPageSize_definitionSize2(int pageSizeId) {
	auto&& __r = QPageSize::definitionSize((QPageSize::PageSizeId)(pageSizeId));
	return new QSizeF(__r);
}
int QPageSize_definitionUnits2(int pageSizeId) {
	auto&& __r = QPageSize::definitionUnits((QPageSize::PageSizeId)(pageSizeId));
	return (int)(__r);
}
void* QPageSize_size2(int pageSizeId, int units) {
	auto&& __r = QPageSize::size((QPageSize::PageSizeId)(pageSizeId), (QPageSize::Unit)(units));
	return new QSizeF(__r);
}
void* QPageSize_sizePoints2(int pageSizeId) {
	auto&& __r = QPageSize::sizePoints((QPageSize::PageSizeId)(pageSizeId));
	return new QSize(__r);
}
void* QPageSize_sizePixels2(int pageSizeId, int resolution) {
	auto&& __r = QPageSize::sizePixels((QPageSize::PageSizeId)(pageSizeId), resolution);
	return new QSize(__r);
}
int QPageSize_id5(void* pointSize, int matchPolicy) {
	auto&& __r = QPageSize::id(*(QSize*)pointSize, (QPageSize::SizeMatchPolicy)(matchPolicy));
	return (int)(__r);
}
int QPageSize_id6(void* size, int units, int matchPolicy) {
	auto&& __r = QPageSize::id(*(QSizeF*)size, (QPageSize::Unit)(units), (QPageSize::SizeMatchPolicy)(matchPolicy));
	return (int)(__r);
}
int QPageSize_OperatorEqual(void* self, void* rhs) {
	auto&& __r = (*(QPageSize*)self) == (*(QPageSize*)rhs);
	return __r;
}
int QPageSize_OperatorNotEqual(void* self, void* rhs) {
	auto&& __r = (*(QPageSize*)self) != (*(QPageSize*)rhs);
	return __r;
}
void QPageSize_delete(void* self) { delete (QPageSize*)self; }

}

#include <QtGui/qsessionmanager.h>
#include <QtCore/qobject.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qvariant.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qsessionmanager.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

static std::unordered_map<long long, QMetaObject::Connection> __conns_QSessionManager;
static std::mutex __connsMu_QSessionManager;

extern "C" {

void* QSessionManager_base(void* self) { return (void*)static_cast<QObject*>((QSessionManager*)self); }
void* QSessionManager_fromQObject(void* o) { return (void*)qobject_cast<QSessionManager*>((QObject*)o); }
char* QSessionManager_tr(const char* s, long long* __out_len) {
	auto&& __r = QSessionManager::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QSessionManager_sessionId(void* self, long long* __out_len) {
	auto&& __r = ((QSessionManager*)self)->sessionId();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QSessionManager_sessionKey(void* self, long long* __out_len) {
	auto&& __r = ((QSessionManager*)self)->sessionKey();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QSessionManager_allowsInteraction(void* self) {
	auto&& __r = ((QSessionManager*)self)->allowsInteraction();
	return __r;
}
int QSessionManager_allowsErrorInteraction(void* self) {
	auto&& __r = ((QSessionManager*)self)->allowsErrorInteraction();
	return __r;
}
void QSessionManager_release(void* self) {
	((QSessionManager*)self)->release();
}
void QSessionManager_cancel(void* self) {
	((QSessionManager*)self)->cancel();
}
void QSessionManager_setRestartHint(void* self, int arg1) {
	((QSessionManager*)self)->setRestartHint((QSessionManager::RestartHint)(arg1));
}
int QSessionManager_restartHint(void* self) {
	auto&& __r = ((QSessionManager*)self)->restartHint();
	return (int)(__r);
}
void QSessionManager_setRestartCommand(void* self, char** arg1, long long arg1_len) {
	QList<QString> arg1_list; arg1_list.reserve(arg1_len);
	for (long long __i = 0; __i < arg1_len; __i++) arg1_list.append(QString::fromUtf8(arg1[__i]));
	((QSessionManager*)self)->setRestartCommand(arg1_list);
}
char** QSessionManager_restartCommand(void* self, long long* __out_len) {
	auto&& __r = ((QSessionManager*)self)->restartCommand();
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
void QSessionManager_setDiscardCommand(void* self, char** arg1, long long arg1_len) {
	QList<QString> arg1_list; arg1_list.reserve(arg1_len);
	for (long long __i = 0; __i < arg1_len; __i++) arg1_list.append(QString::fromUtf8(arg1[__i]));
	((QSessionManager*)self)->setDiscardCommand(arg1_list);
}
char** QSessionManager_discardCommand(void* self, long long* __out_len) {
	auto&& __r = ((QSessionManager*)self)->discardCommand();
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
void QSessionManager_setManagerProperty(void* self, const char* name, long long name_len, const char* value, long long value_len) {
	QString name_qstr = QString::fromUtf8(name, name_len);
	QString value_qstr = QString::fromUtf8(value, value_len);
	((QSessionManager*)self)->setManagerProperty(name_qstr, value_qstr);
}
void QSessionManager_setManagerProperty2(void* self, const char* name, long long name_len, char** value, long long value_len) {
	QString name_qstr = QString::fromUtf8(name, name_len);
	QList<QString> value_list; value_list.reserve(value_len);
	for (long long __i = 0; __i < value_len; __i++) value_list.append(QString::fromUtf8(value[__i]));
	((QSessionManager*)self)->setManagerProperty(name_qstr, value_list);
}
int QSessionManager_isPhase2(void* self) {
	auto&& __r = ((QSessionManager*)self)->isPhase2();
	return __r;
}
void QSessionManager_requestPhase2(void* self) {
	((QSessionManager*)self)->requestPhase2();
}
void* QSessionManager_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QSessionManager::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QSessionManager_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QSessionManager::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
char* QSessionManager_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QSessionManager::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QSessionManager_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QSessionManager::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QSessionManager_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QSessionManager::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QSessionManager_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QSessionManager::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void QSessionManager_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSessionManager*)self, &QSessionManager::destroyed, (QSessionManager*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSessionManager);
	__conns_QSessionManager[__id] = __c;
}
void QSessionManager_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSessionManager*)self, &QSessionManager::destroyed, (QSessionManager*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSessionManager);
	__conns_QSessionManager[__id] = __c;
}
void QSessionManager_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSessionManager*)self, &QSessionManager::objectNameChanged, (QSessionManager*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSessionManager);
	__conns_QSessionManager[__id] = __c;
}
void QSessionManager__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QSessionManager);
		auto __it = __conns_QSessionManager.find(id);
		if (__it == __conns_QSessionManager.end()) return;
		__c = __it->second;
		__conns_QSessionManager.erase(__it);
	}
	QObject::disconnect(__c);
}

}

#include <QtGui/qvalidator.h>
#include <QtCore/qobject.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qlocale.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qvariant.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qdoublevalidator.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QDoubleValidator__Go : public QDoubleValidator {
public:
	long long __goId;
	unsigned char __goOn[10] = {};
	template<typename... A> QDoubleValidator__Go(long long id, A&&... a) : QDoubleValidator(std::forward<A>(a)...), __goId(id) {}
	using QDoubleValidator::sender;
	using QDoubleValidator::senderSignalIndex;
	using QDoubleValidator::receivers;
	using QDoubleValidator::isSignalConnected;
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QDoubleValidator::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QDoubleValidator::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QDoubleValidator::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QDoubleValidator::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QDoubleValidator::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QDoubleValidator::metaObject();
	}
	QValidator::State validate(QString& arg1, int& arg2) const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_int_voidp_intp) { int __r = ((int(*)(long long, long long, void*, int*, int*))cppgo_virt_int_voidp_intp)(__goId, 6784179835071232555LL, (void*)&arg1, (int*)&arg2, &__h); if (__h) return (QValidator::State)(__r); }
		return QDoubleValidator::validate(arg1, arg2);
	}
	void fixup(QString& input) const override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4855206404746966715LL, (void*)&input, &__h); if (__h) return; }
		QDoubleValidator::fixup(input);
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QDoubleValidator::event(event);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QDoubleValidator::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QDoubleValidator::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QDoubleValidator::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QDoubleValidator::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QDoubleValidator::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QDoubleValidator::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QDoubleValidator;
static std::mutex __connsMu_QDoubleValidator;

extern "C" {

void* QDoubleValidator_base(void* self) { return (void*)static_cast<QValidator*>((QDoubleValidator*)self); }
void QDoubleValidator_chain(void* self, void** out) {
	QDoubleValidator* __p = (QDoubleValidator*)self;
	QValidator* __b0 = static_cast<QValidator*>(__p); out[0] = (void*)__b0;
	QObject* __b1 = static_cast<QObject*>(__b0); out[1] = (void*)__b1;
}
void QDoubleValidator__set_override(void* self, int idx) { ((QDoubleValidator__Go*)self)->__goOn[idx] = 1; }
void* QDoubleValidator_fromQObject(void* o) { return (void*)qobject_cast<QDoubleValidator*>((QObject*)o); }
void* QDoubleValidator_new(long long __goid) {
	return new QDoubleValidator__Go(__goid);
}
void* QDoubleValidator_new2(long long __goid, double bottom, double top, int decimals) {
	return new QDoubleValidator__Go(__goid, bottom, top, decimals);
}
void* QDoubleValidator_new3(long long __goid, void* parent) {
	return new QDoubleValidator__Go(__goid, (QObject*)parent);
}
void* QDoubleValidator_new4(long long __goid, double bottom, double top, int decimals, void* parent) {
	return new QDoubleValidator__Go(__goid, bottom, top, decimals, (QObject*)parent);
}
char* QDoubleValidator_tr(const char* s, long long* __out_len) {
	auto&& __r = QDoubleValidator::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QDoubleValidator_setRange(void* self, double bottom, double top, int decimals) {
	((QDoubleValidator*)self)->setRange(bottom, top, decimals);
}
void QDoubleValidator_setRange2(void* self, double bottom, double top) {
	((QDoubleValidator*)self)->setRange(bottom, top);
}
void QDoubleValidator_setBottom(void* self, double arg1) {
	((QDoubleValidator*)self)->setBottom(arg1);
}
void QDoubleValidator_setTop(void* self, double arg1) {
	((QDoubleValidator*)self)->setTop(arg1);
}
void QDoubleValidator_setDecimals(void* self, int arg1) {
	((QDoubleValidator*)self)->setDecimals(arg1);
}
void QDoubleValidator_setNotation(void* self, int arg1) {
	((QDoubleValidator*)self)->setNotation((QDoubleValidator::Notation)(arg1));
}
double QDoubleValidator_bottom(void* self) {
	auto&& __r = ((QDoubleValidator*)self)->bottom();
	return __r;
}
double QDoubleValidator_top(void* self) {
	auto&& __r = ((QDoubleValidator*)self)->top();
	return __r;
}
int QDoubleValidator_decimals(void* self) {
	auto&& __r = ((QDoubleValidator*)self)->decimals();
	return __r;
}
int QDoubleValidator_notation(void* self) {
	auto&& __r = ((QDoubleValidator*)self)->notation();
	return (int)(__r);
}
void QDoubleValidator_bottomChanged(void* self, double bottom) {
	((QDoubleValidator*)self)->bottomChanged(bottom);
}
void QDoubleValidator_topChanged(void* self, double top) {
	((QDoubleValidator*)self)->topChanged(top);
}
void QDoubleValidator_decimalsChanged(void* self, int decimals) {
	((QDoubleValidator*)self)->decimalsChanged(decimals);
}
void QDoubleValidator_notationChanged(void* self, int notation) {
	((QDoubleValidator*)self)->notationChanged((QDoubleValidator::Notation)(notation));
}
void* QDoubleValidator_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QDoubleValidator::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QDoubleValidator_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QDoubleValidator::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QDoubleValidator_sender(void* self) {
	auto&& __r = ((QDoubleValidator__Go*)self)->sender();
	return (void*)(__r);
}
int QDoubleValidator_senderSignalIndex(void* self) {
	auto&& __r = ((QDoubleValidator__Go*)self)->senderSignalIndex();
	return __r;
}
int QDoubleValidator_receivers(void* self, const char* signal) {
	auto&& __r = ((QDoubleValidator__Go*)self)->receivers(signal);
	return __r;
}
int QDoubleValidator_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QDoubleValidator__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QDoubleValidator_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QDoubleValidator::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QDoubleValidator_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QDoubleValidator::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QDoubleValidator_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QDoubleValidator::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QDoubleValidator_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QDoubleValidator::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QDoubleValidator_metaObject_default(void* self) {
	auto&& __r = ((QDoubleValidator*)self)->QDoubleValidator::metaObject();
	return (void*)(__r);
}
int QDoubleValidator_validate_default(void* self, void* arg1, int* arg2) {
	int arg2_ref = *arg2;
	auto&& __r = ((QDoubleValidator*)self)->QDoubleValidator::validate(*(QString*)arg1, arg2_ref);
	*arg2 = arg2_ref;
	return (int)(__r);
}
void QDoubleValidator_fixup_default(void* self, void* input) {
	((QDoubleValidator*)self)->QDoubleValidator::fixup(*(QString*)input);
}
int QDoubleValidator_event_default(void* self, void* event) {
	auto&& __r = ((QDoubleValidator*)self)->QDoubleValidator::event((QEvent*)event);
	return __r;
}
int QDoubleValidator_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QDoubleValidator*)self)->QDoubleValidator::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QDoubleValidator_timerEvent_default(void* self, void* event) {
	((QDoubleValidator__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QDoubleValidator_childEvent_default(void* self, void* event) {
	((QDoubleValidator__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QDoubleValidator_customEvent_default(void* self, void* event) {
	((QDoubleValidator__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QDoubleValidator_connectNotify_default(void* self, void* signal) {
	((QDoubleValidator__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QDoubleValidator_disconnectNotify_default(void* self, void* signal) {
	((QDoubleValidator__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QDoubleValidator_connect_bottomChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDoubleValidator*)self, &QDoubleValidator::bottomChanged, (QDoubleValidator*)self, [__id](double bottom){
		if (cppgo_sig_void_unsignedlonglong) ((void(*)(long long, unsigned long long))cppgo_sig_void_unsignedlonglong)(__id, __builtin_bit_cast(unsigned long long, (double)(bottom)));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDoubleValidator);
	__conns_QDoubleValidator[__id] = __c;
}
void QDoubleValidator_connect_topChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDoubleValidator*)self, &QDoubleValidator::topChanged, (QDoubleValidator*)self, [__id](double top){
		if (cppgo_sig_void_unsignedlonglong) ((void(*)(long long, unsigned long long))cppgo_sig_void_unsignedlonglong)(__id, __builtin_bit_cast(unsigned long long, (double)(top)));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDoubleValidator);
	__conns_QDoubleValidator[__id] = __c;
}
void QDoubleValidator_connect_decimalsChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDoubleValidator*)self, &QDoubleValidator::decimalsChanged, (QDoubleValidator*)self, [__id](int decimals){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, decimals);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDoubleValidator);
	__conns_QDoubleValidator[__id] = __c;
}
void QDoubleValidator_connect_notationChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDoubleValidator*)self, &QDoubleValidator::notationChanged, (QDoubleValidator*)self, [__id](QDoubleValidator::Notation notation){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, (int)(notation));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDoubleValidator);
	__conns_QDoubleValidator[__id] = __c;
}
void QDoubleValidator_connect_changed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDoubleValidator*)self, &QDoubleValidator::changed, (QDoubleValidator*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDoubleValidator);
	__conns_QDoubleValidator[__id] = __c;
}
void QDoubleValidator_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDoubleValidator*)self, &QDoubleValidator::destroyed, (QDoubleValidator*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDoubleValidator);
	__conns_QDoubleValidator[__id] = __c;
}
void QDoubleValidator_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDoubleValidator*)self, &QDoubleValidator::destroyed, (QDoubleValidator*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDoubleValidator);
	__conns_QDoubleValidator[__id] = __c;
}
void QDoubleValidator_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDoubleValidator*)self, &QDoubleValidator::objectNameChanged, (QDoubleValidator*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDoubleValidator);
	__conns_QDoubleValidator[__id] = __c;
}
void QDoubleValidator__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QDoubleValidator);
		auto __it = __conns_QDoubleValidator.find(id);
		if (__it == __conns_QDoubleValidator.end()) return;
		__c = __it->second;
		__conns_QDoubleValidator.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QDoubleValidator__Go_new(long long id, void* parent) { return new QDoubleValidator__Go(id, (QObject*)parent); }
void QDoubleValidator_delete(void* self) { delete (QDoubleValidator*)self; }

}

#include <QtWidgets/qstyleoption.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qnamespace.h>
#include "gen_qstyleoptiontoolbar.h"

extern "C" {

void* QStyleOptionToolBar_base(void* self) { return (void*)static_cast<QStyleOption*>((QStyleOptionToolBar*)self); }
void* QStyleOptionToolBar_new() {
	return new QStyleOptionToolBar();
}
void* QStyleOptionToolBar_new2(void* other) {
	return new QStyleOptionToolBar(*(QStyleOptionToolBar*)other);
}
void* QStyleOptionToolBar_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QStyleOptionToolBar*)self)->operator=(*(QStyleOptionToolBar*)arg1);
	return (void*)&(__r);
}
int QStyleOptionToolBar_positionOfLine_get(void* self) {
	auto __r = ((QStyleOptionToolBar*)self)->positionOfLine;
	return (int)(__r);
}
void QStyleOptionToolBar_positionOfLine_set(void* self, int positionOfLine) {
	((QStyleOptionToolBar*)self)->positionOfLine = (QStyleOptionToolBar::ToolBarPosition)(positionOfLine);
}
int QStyleOptionToolBar_positionWithinLine_get(void* self) {
	auto __r = ((QStyleOptionToolBar*)self)->positionWithinLine;
	return (int)(__r);
}
void QStyleOptionToolBar_positionWithinLine_set(void* self, int positionWithinLine) {
	((QStyleOptionToolBar*)self)->positionWithinLine = (QStyleOptionToolBar::ToolBarPosition)(positionWithinLine);
}
int QStyleOptionToolBar_toolBarArea_get(void* self) {
	auto __r = ((QStyleOptionToolBar*)self)->toolBarArea;
	return (int)(__r);
}
void QStyleOptionToolBar_toolBarArea_set(void* self, int toolBarArea) {
	((QStyleOptionToolBar*)self)->toolBarArea = (Qt::ToolBarArea)(toolBarArea);
}
int QStyleOptionToolBar_features_get(void* self) {
	auto __r = ((QStyleOptionToolBar*)self)->features;
	return (int)(__r);
}
void QStyleOptionToolBar_features_set(void* self, int features) {
	((QStyleOptionToolBar*)self)->features = QFlags<QStyleOptionToolBar::ToolBarFeature>(features);
}
int QStyleOptionToolBar_lineWidth_get(void* self) {
	auto __r = ((QStyleOptionToolBar*)self)->lineWidth;
	return __r;
}
void QStyleOptionToolBar_lineWidth_set(void* self, int lineWidth) {
	((QStyleOptionToolBar*)self)->lineWidth = lineWidth;
}
int QStyleOptionToolBar_midLineWidth_get(void* self) {
	auto __r = ((QStyleOptionToolBar*)self)->midLineWidth;
	return __r;
}
void QStyleOptionToolBar_midLineWidth_set(void* self, int midLineWidth) {
	((QStyleOptionToolBar*)self)->midLineWidth = midLineWidth;
}
void QStyleOptionToolBar_delete(void* self) { delete (QStyleOptionToolBar*)self; }

}

#include <QtWidgets/qaccessiblewidget.h>
#include <QtGui/qaccessible.h>
#include <QtWidgets/qwidget.h>
#include <QtGui/qaccessible_base.h>
#include <QtCore/qstring.h>
#include <QtCore/qvariant.h>
#include <QtGui/qwindow.h>
#include <QtCore/qrect.h>
#include <QtGui/qcolor.h>
#include <QtCore/qobject.h>
#include "gen_qaccessiblewidgetv2.h"

extern "C" {

void* QAccessibleWidgetV2_base(void* self) { return (void*)static_cast<QAccessibleWidget*>((QAccessibleWidgetV2*)self); }
void QAccessibleWidgetV2_chain(void* self, void** out) {
	QAccessibleWidgetV2* __p = (QAccessibleWidgetV2*)self;
	QAccessibleWidget* __b0 = static_cast<QAccessibleWidget*>(__p); out[0] = (void*)__b0;
	QAccessibleObject* __b1 = static_cast<QAccessibleObject*>(__b0); out[1] = (void*)__b1;
	QAccessibleInterface* __b2 = static_cast<QAccessibleInterface*>(__b1); out[2] = (void*)__b2;
}
void* QAccessibleWidgetV2_new(void* object) {
	return new QAccessibleWidgetV2((QWidget*)object);
}
void* QAccessibleWidgetV2_new2(void* object, int role, const char* name, long long name_len) {
	QString name_qstr = QString::fromUtf8(name, name_len);
	return new QAccessibleWidgetV2((QWidget*)object, (QAccessible::Role)(role), name_qstr);
}
void* QAccessibleWidgetV2_new3(void* object, int role) {
	return new QAccessibleWidgetV2((QWidget*)object, (QAccessible::Role)(role));
}
int* QAccessibleWidgetV2_attributeKeys(void* self, long long* __out_len) {
	auto&& __r = ((QAccessibleWidgetV2*)self)->attributeKeys();
	*__out_len = __r.size();
	int* __arr = (int*)malloc(sizeof(int) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (int)(__e);
	return __arr;
}
void* QAccessibleWidgetV2_attributeValue(void* self, int key) {
	auto&& __r = ((QAccessibleWidgetV2*)self)->attributeValue((QAccessible::Attribute)(key));
	return new QVariant(__r);
}
char* QAccessibleWidgetV2_tr(const char* sourceText, long long* __out_len) {
	auto&& __r = QAccessibleWidgetV2::tr(sourceText);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QAccessibleWidgetV2_pressAction(long long* __out_len) {
	auto&& __r = QAccessibleWidgetV2::pressAction();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QAccessibleWidgetV2_increaseAction(long long* __out_len) {
	auto&& __r = QAccessibleWidgetV2::increaseAction();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QAccessibleWidgetV2_decreaseAction(long long* __out_len) {
	auto&& __r = QAccessibleWidgetV2::decreaseAction();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QAccessibleWidgetV2_showMenuAction(long long* __out_len) {
	auto&& __r = QAccessibleWidgetV2::showMenuAction();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QAccessibleWidgetV2_setFocusAction(long long* __out_len) {
	auto&& __r = QAccessibleWidgetV2::setFocusAction();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QAccessibleWidgetV2_toggleAction(long long* __out_len) {
	auto&& __r = QAccessibleWidgetV2::toggleAction();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QAccessibleWidgetV2_scrollLeftAction(long long* __out_len) {
	auto&& __r = QAccessibleWidgetV2::scrollLeftAction();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QAccessibleWidgetV2_scrollRightAction(long long* __out_len) {
	auto&& __r = QAccessibleWidgetV2::scrollRightAction();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QAccessibleWidgetV2_scrollUpAction(long long* __out_len) {
	auto&& __r = QAccessibleWidgetV2::scrollUpAction();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QAccessibleWidgetV2_scrollDownAction(long long* __out_len) {
	auto&& __r = QAccessibleWidgetV2::scrollDownAction();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QAccessibleWidgetV2_nextPageAction(long long* __out_len) {
	auto&& __r = QAccessibleWidgetV2::nextPageAction();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QAccessibleWidgetV2_previousPageAction(long long* __out_len) {
	auto&& __r = QAccessibleWidgetV2::previousPageAction();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QAccessibleWidgetV2_tr2(const char* sourceText, const char* disambiguation, long long* __out_len) {
	auto&& __r = QAccessibleWidgetV2::tr(sourceText, disambiguation);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QAccessibleWidgetV2_tr3(const char* sourceText, const char* disambiguation, int n, long long* __out_len) {
	auto&& __r = QAccessibleWidgetV2::tr(sourceText, disambiguation, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QAccessibleWidgetV2_asQAccessibleAttributesInterface(void* self) {
	auto&& __r = static_cast<QAccessibleAttributesInterface*>((QAccessibleWidgetV2*)self);
	return (void*)(__r);
}

}

#include <QtWidgets/qdatetimeedit.h>
#include <QtWidgets/qabstractspinbox.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qdatetime.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qcalendar.h>
#include <QtWidgets/qcalendarwidget.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qtimezone.h>
#include <QtCore/qsize.h>
#include <QtCore/qcoreevent.h>
#include <QtGui/qevent.h>
#include <QtGui/qvalidator.h>
#include <QtWidgets/qstyleoption.h>
#include <QtCore/qvariant.h>
#include <QtWidgets/qlineedit.h>
#include <QtWidgets/qstyle.h>
#include <QtCore/qrect.h>
#include <QtCore/qpoint.h>
#include <QtGui/qregion.h>
#include <QtGui/qpalette.h>
#include <QtGui/qfont.h>
#include <QtGui/qfontmetrics.h>
#include <QtGui/qfontinfo.h>
#include <QtGui/qcursor.h>
#include <QtGui/qbitmap.h>
#include <QtGui/qpaintdevice.h>
#include <QtGui/qpainter.h>
#include <QtGui/qpixmap.h>
#include <QtWidgets/qgraphicseffect.h>
#include <QtGui/qicon.h>
#include <QtCore/qlocale.h>
#include <QtGui/qkeysequence.h>
#include <QtWidgets/qgraphicsproxywidget.h>
#include <QtCore/qbytearray.h>
#include <QtWidgets/qsizepolicy.h>
#include <QtCore/qmargins.h>
#include <QtWidgets/qlayout.h>
#include <QtGui/qaction.h>
#include <QtCore/qobject.h>
#include <QtGui/qpaintengine.h>
#include <QtGui/qbackingstore.h>
#include <QtGui/qwindow.h>
#include <QtGui/qscreen.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qdatetimeedit.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QDateTimeEdit__Go : public QDateTimeEdit {
public:
	long long __goId;
	unsigned char __goOn[56] = {};
	template<typename... A> QDateTimeEdit__Go(long long id, A&&... a) : QDateTimeEdit(std::forward<A>(a)...), __goId(id) {}
	using QDateTimeEdit::lineEdit;
	using QDateTimeEdit::setLineEdit;
	using QDateTimeEdit::updateMicroFocus;
	using QDateTimeEdit::create;
	using QDateTimeEdit::destroy;
	using QDateTimeEdit::focusNextChild;
	using QDateTimeEdit::focusPreviousChild;
	using QDateTimeEdit::sender;
	using QDateTimeEdit::senderSignalIndex;
	using QDateTimeEdit::receivers;
	using QDateTimeEdit::isSignalConnected;
	using QDateTimeEdit::getDecodedMetricF;
	void cppgoBase_keyPressEvent(QKeyEvent* __a0) { QDateTimeEdit::keyPressEvent(__a0); }
	void cppgoBase_wheelEvent(QWheelEvent* __a0) { QDateTimeEdit::wheelEvent(__a0); }
	void cppgoBase_focusInEvent(QFocusEvent* __a0) { QDateTimeEdit::focusInEvent(__a0); }
	bool cppgoBase_focusNextPrevChild(bool __a0) { return QDateTimeEdit::focusNextPrevChild(__a0); }
	QValidator::State cppgoBase_validate(QString& __a0, int& __a1) { return QDateTimeEdit::validate(__a0, __a1); }
	void cppgoBase_fixup(QString& __a0) { QDateTimeEdit::fixup(__a0); }
	QDateTime cppgoBase_dateTimeFromText(const QString& __a0) { return QDateTimeEdit::dateTimeFromText(__a0); }
	QString cppgoBase_textFromDateTime(const QDateTime& __a0) { return QDateTimeEdit::textFromDateTime(__a0); }
	QFlags<QAbstractSpinBox::StepEnabledFlag> cppgoBase_stepEnabled() { return QDateTimeEdit::stepEnabled(); }
	void cppgoBase_mousePressEvent(QMouseEvent* __a0) { QDateTimeEdit::mousePressEvent(__a0); }
	void cppgoBase_paintEvent(QPaintEvent* __a0) { QDateTimeEdit::paintEvent(__a0); }
	void cppgoBase_initStyleOption(QStyleOptionSpinBox* __a0) { QDateTimeEdit::initStyleOption(__a0); }
	void cppgoBase_resizeEvent(QResizeEvent* __a0) { QDateTimeEdit::resizeEvent(__a0); }
	void cppgoBase_keyReleaseEvent(QKeyEvent* __a0) { QDateTimeEdit::keyReleaseEvent(__a0); }
	void cppgoBase_focusOutEvent(QFocusEvent* __a0) { QDateTimeEdit::focusOutEvent(__a0); }
	void cppgoBase_contextMenuEvent(QContextMenuEvent* __a0) { QDateTimeEdit::contextMenuEvent(__a0); }
	void cppgoBase_changeEvent(QEvent* __a0) { QDateTimeEdit::changeEvent(__a0); }
	void cppgoBase_closeEvent(QCloseEvent* __a0) { QDateTimeEdit::closeEvent(__a0); }
	void cppgoBase_hideEvent(QHideEvent* __a0) { QDateTimeEdit::hideEvent(__a0); }
	void cppgoBase_mouseReleaseEvent(QMouseEvent* __a0) { QDateTimeEdit::mouseReleaseEvent(__a0); }
	void cppgoBase_mouseMoveEvent(QMouseEvent* __a0) { QDateTimeEdit::mouseMoveEvent(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QDateTimeEdit::timerEvent(__a0); }
	void cppgoBase_showEvent(QShowEvent* __a0) { QDateTimeEdit::showEvent(__a0); }
	void cppgoBase_mouseDoubleClickEvent(QMouseEvent* __a0) { QDateTimeEdit::mouseDoubleClickEvent(__a0); }
	void cppgoBase_enterEvent(QEnterEvent* __a0) { QDateTimeEdit::enterEvent(__a0); }
	void cppgoBase_leaveEvent(QEvent* __a0) { QDateTimeEdit::leaveEvent(__a0); }
	void cppgoBase_moveEvent(QMoveEvent* __a0) { QDateTimeEdit::moveEvent(__a0); }
	void cppgoBase_tabletEvent(QTabletEvent* __a0) { QDateTimeEdit::tabletEvent(__a0); }
	void cppgoBase_actionEvent(QActionEvent* __a0) { QDateTimeEdit::actionEvent(__a0); }
	void cppgoBase_dragEnterEvent(QDragEnterEvent* __a0) { QDateTimeEdit::dragEnterEvent(__a0); }
	void cppgoBase_dragMoveEvent(QDragMoveEvent* __a0) { QDateTimeEdit::dragMoveEvent(__a0); }
	void cppgoBase_dragLeaveEvent(QDragLeaveEvent* __a0) { QDateTimeEdit::dragLeaveEvent(__a0); }
	void cppgoBase_dropEvent(QDropEvent* __a0) { QDateTimeEdit::dropEvent(__a0); }
	int cppgoBase_metric(int __a0) { return QDateTimeEdit::metric((QPaintDevice::PaintDeviceMetric)__a0); }
	void cppgoBase_initPainter(QPainter* __a0) { QDateTimeEdit::initPainter(__a0); }
	QPaintDevice* cppgoBase_redirected(QPoint* __a0) { return QDateTimeEdit::redirected(__a0); }
	QPainter* cppgoBase_sharedPainter() { return QDateTimeEdit::sharedPainter(); }
	void cppgoBase_inputMethodEvent(QInputMethodEvent* __a0) { QDateTimeEdit::inputMethodEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QDateTimeEdit::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QDateTimeEdit::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QDateTimeEdit::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QDateTimeEdit::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QDateTimeEdit::metaObject();
	}
	QSize sizeHint() const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 729690261457684101LL, &__h); if (__h) return *(QSize*)__r; }
		return QDateTimeEdit::sizeHint();
	}
	void clear() override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, -778567940572639070LL, &__h); if (__h) return; }
		QDateTimeEdit::clear();
	}
	void stepBy(int steps) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 7085718643589954634LL, steps, &__h); if (__h) return; }
		QDateTimeEdit::stepBy(steps);
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QDateTimeEdit::event(event);
	}
	void keyPressEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1023791977679793185LL, (void*)event, &__h); if (__h) return; }
		QDateTimeEdit::keyPressEvent(event);
	}
	void wheelEvent(QWheelEvent* event) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1919988340934511316LL, (void*)event, &__h); if (__h) return; }
		QDateTimeEdit::wheelEvent(event);
	}
	void focusInEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 975043767033552552LL, (void*)event, &__h); if (__h) return; }
		QDateTimeEdit::focusInEvent(event);
	}
	bool focusNextPrevChild(bool next) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 5898328517505784191LL, next, &__h); if (__h) return __r; }
		return QDateTimeEdit::focusNextPrevChild(next);
	}
	QValidator::State validate(QString& input, int& pos) const override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_int_voidp_intp) { int __r = ((int(*)(long long, long long, void*, int*, int*))cppgo_virt_int_voidp_intp)(__goId, 6784179835071232555LL, (void*)&input, (int*)&pos, &__h); if (__h) return (QValidator::State)(__r); }
		return QDateTimeEdit::validate(input, pos);
	}
	void fixup(QString& input) const override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4855206404746966715LL, (void*)&input, &__h); if (__h) return; }
		QDateTimeEdit::fixup(input);
	}
	QDateTime dateTimeFromText(const QString& text) const override {
		int __h = 0;
		QByteArray text_ba = text.toUtf8();
		if (__goOn[11] && cppgo_virt_voidp_constcharp_longlong) { void* __r = ((void*(*)(long long, long long, const char*, long long, int*))cppgo_virt_voidp_constcharp_longlong)(__goId, -50441425836614729LL, (char*)text_ba.constData(), text_ba.size(), &__h); if (__h) return *(QDateTime*)__r; }
		return QDateTimeEdit::dateTimeFromText(text);
	}
	QString textFromDateTime(const QDateTime& dt) const override {
		int __h = 0;
		if (__goOn[12] && cppgo_virtp_charp_voidp) {
			long long __len = 0;
			char* __s = ((char*(*)(long long, long long, void*, long long*, int*))cppgo_virtp_charp_voidp)(__goId, -3778148144965928937LL, (void*)&dt, &__len, &__h);
			if (__h) { QString __r = QString::fromUtf8(__s, __len); free(__s); return __r; }
		}
		return QDateTimeEdit::textFromDateTime(dt);
	}
	QFlags<QAbstractSpinBox::StepEnabledFlag> stepEnabled() const override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, 5257289574516412244LL, &__h); if (__h) return QFlags<QAbstractSpinBox::StepEnabledFlag>(__r); }
		return QDateTimeEdit::stepEnabled();
	}
	void mousePressEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -332090898078543319LL, (void*)event, &__h); if (__h) return; }
		QDateTimeEdit::mousePressEvent(event);
	}
	void paintEvent(QPaintEvent* event) override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5290442163167458465LL, (void*)event, &__h); if (__h) return; }
		QDateTimeEdit::paintEvent(event);
	}
	void initStyleOption(QStyleOptionSpinBox* option) const override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -488693769771308531LL, (void*)option, &__h); if (__h) return; }
		QDateTimeEdit::initStyleOption(option);
	}
	QSize minimumSizeHint() const override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 7049264599837532491LL, &__h); if (__h) return *(QSize*)__r; }
		return QDateTimeEdit::minimumSizeHint();
	}
	QVariant inputMethodQuery(Qt::InputMethodQuery arg1) const override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7345853097340709638LL, (int)(arg1), &__h); if (__h) return *(QVariant*)__r; }
		return QDateTimeEdit::inputMethodQuery(arg1);
	}
	void resizeEvent(QResizeEvent* event) override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4302917285655014661LL, (void*)event, &__h); if (__h) return; }
		QDateTimeEdit::resizeEvent(event);
	}
	void keyReleaseEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 772598475950050299LL, (void*)event, &__h); if (__h) return; }
		QDateTimeEdit::keyReleaseEvent(event);
	}
	void focusOutEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -385596010576048993LL, (void*)event, &__h); if (__h) return; }
		QDateTimeEdit::focusOutEvent(event);
	}
	void contextMenuEvent(QContextMenuEvent* event) override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6698043222227006015LL, (void*)event, &__h); if (__h) return; }
		QDateTimeEdit::contextMenuEvent(event);
	}
	void changeEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4959376063826165943LL, (void*)event, &__h); if (__h) return; }
		QDateTimeEdit::changeEvent(event);
	}
	void closeEvent(QCloseEvent* event) override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1989937050019464783LL, (void*)event, &__h); if (__h) return; }
		QDateTimeEdit::closeEvent(event);
	}
	void hideEvent(QHideEvent* event) override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -199718707876719129LL, (void*)event, &__h); if (__h) return; }
		QDateTimeEdit::hideEvent(event);
	}
	void mouseReleaseEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4415920049297795213LL, (void*)event, &__h); if (__h) return; }
		QDateTimeEdit::mouseReleaseEvent(event);
	}
	void mouseMoveEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4255494004099806581LL, (void*)event, &__h); if (__h) return; }
		QDateTimeEdit::mouseMoveEvent(event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QDateTimeEdit::timerEvent(event);
	}
	void showEvent(QShowEvent* event) override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6825369174751916LL, (void*)event, &__h); if (__h) return; }
		QDateTimeEdit::showEvent(event);
	}
	int devType() const override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8659679893868524044LL, &__h); if (__h) return __r; }
		return QDateTimeEdit::devType();
	}
	void setVisible(bool visible) override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 2110182719375245575LL, visible, &__h); if (__h) return; }
		QDateTimeEdit::setVisible(visible);
	}
	int heightForWidth(int arg1) const override {
		int __h = 0;
		if (__goOn[32] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 4311534060495737339LL, arg1, &__h); if (__h) return __r; }
		return QDateTimeEdit::heightForWidth(arg1);
	}
	bool hasHeightForWidth() const override {
		int __h = 0;
		if (__goOn[33] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8266762397995423915LL, &__h); if (__h) return __r; }
		return QDateTimeEdit::hasHeightForWidth();
	}
	QPaintEngine* paintEngine() const override {
		int __h = 0;
		if (__goOn[34] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 6398007404756346381LL, &__h); if (__h) return (QPaintEngine*)__r; }
		return QDateTimeEdit::paintEngine();
	}
	void mouseDoubleClickEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[35] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4029073905697023887LL, (void*)event, &__h); if (__h) return; }
		QDateTimeEdit::mouseDoubleClickEvent(event);
	}
	void enterEvent(QEnterEvent* event) override {
		int __h = 0;
		if (__goOn[36] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6316270346668237895LL, (void*)event, &__h); if (__h) return; }
		QDateTimeEdit::enterEvent(event);
	}
	void leaveEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[37] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -551070009623711640LL, (void*)event, &__h); if (__h) return; }
		QDateTimeEdit::leaveEvent(event);
	}
	void moveEvent(QMoveEvent* event) override {
		int __h = 0;
		if (__goOn[38] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -9161949150489525396LL, (void*)event, &__h); if (__h) return; }
		QDateTimeEdit::moveEvent(event);
	}
	void tabletEvent(QTabletEvent* event) override {
		int __h = 0;
		if (__goOn[39] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -212068866860701141LL, (void*)event, &__h); if (__h) return; }
		QDateTimeEdit::tabletEvent(event);
	}
	void actionEvent(QActionEvent* event) override {
		int __h = 0;
		if (__goOn[40] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1697444140088555581LL, (void*)event, &__h); if (__h) return; }
		QDateTimeEdit::actionEvent(event);
	}
	void dragEnterEvent(QDragEnterEvent* event) override {
		int __h = 0;
		if (__goOn[41] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 206122715919917589LL, (void*)event, &__h); if (__h) return; }
		QDateTimeEdit::dragEnterEvent(event);
	}
	void dragMoveEvent(QDragMoveEvent* event) override {
		int __h = 0;
		if (__goOn[42] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1781545137352745978LL, (void*)event, &__h); if (__h) return; }
		QDateTimeEdit::dragMoveEvent(event);
	}
	void dragLeaveEvent(QDragLeaveEvent* event) override {
		int __h = 0;
		if (__goOn[43] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -3139904158373725006LL, (void*)event, &__h); if (__h) return; }
		QDateTimeEdit::dragLeaveEvent(event);
	}
	void dropEvent(QDropEvent* event) override {
		int __h = 0;
		if (__goOn[44] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 9107058682754591716LL, (void*)event, &__h); if (__h) return; }
		QDateTimeEdit::dropEvent(event);
	}
	bool nativeEvent(const QByteArray& eventType, void* message, qintptr* result) override {
		int __h = 0;
		if (__goOn[45] && cppgo_virt_int_voidp_voidp_longlongp) { int __r = ((int(*)(long long, long long, void*, void*, long long*, int*))cppgo_virt_int_voidp_voidp_longlongp)(__goId, -2039166586246443706LL, (void*)&eventType, (void*)message, (long long*)result, &__h); if (__h) return __r; }
		return QDateTimeEdit::nativeEvent(eventType, message, result);
	}
	int metric(QPaintDevice::PaintDeviceMetric arg1) const override {
		int __h = 0;
		if (__goOn[46] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 427629649267311777LL, (int)(arg1), &__h); if (__h) return __r; }
		return QDateTimeEdit::metric(arg1);
	}
	void initPainter(QPainter* painter) const override {
		int __h = 0;
		if (__goOn[47] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6477500296980050210LL, (void*)painter, &__h); if (__h) return; }
		QDateTimeEdit::initPainter(painter);
	}
	QPaintDevice* redirected(QPoint* offset) const override {
		int __h = 0;
		if (__goOn[48] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 4127240633310655048LL, (void*)offset, &__h); if (__h) return (QPaintDevice*)__r; }
		return QDateTimeEdit::redirected(offset);
	}
	QPainter* sharedPainter() const override {
		int __h = 0;
		if (__goOn[49] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 5792771328741774991LL, &__h); if (__h) return (QPainter*)__r; }
		return QDateTimeEdit::sharedPainter();
	}
	void inputMethodEvent(QInputMethodEvent* arg1) override {
		int __h = 0;
		if (__goOn[50] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5438819661942843222LL, (void*)arg1, &__h); if (__h) return; }
		QDateTimeEdit::inputMethodEvent(arg1);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[51] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QDateTimeEdit::eventFilter(watched, event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[52] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QDateTimeEdit::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[53] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QDateTimeEdit::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[54] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QDateTimeEdit::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[55] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QDateTimeEdit::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QDateTimeEdit;
static std::mutex __connsMu_QDateTimeEdit;

extern "C" {

void* QDateTimeEdit_base(void* self) { return (void*)static_cast<QAbstractSpinBox*>((QDateTimeEdit*)self); }
void QDateTimeEdit_chain(void* self, void** out) {
	QDateTimeEdit* __p = (QDateTimeEdit*)self;
	QAbstractSpinBox* __b0 = static_cast<QAbstractSpinBox*>(__p); out[0] = (void*)__b0;
	QWidget* __b1 = static_cast<QWidget*>(__b0); out[1] = (void*)__b1;
	QObject* __b2 = static_cast<QObject*>(__b1); out[2] = (void*)__b2;
}
void QDateTimeEdit__set_override(void* self, int idx) { ((QDateTimeEdit__Go*)self)->__goOn[idx] = 1; }
void* QDateTimeEdit_fromQObject(void* o) { return (void*)qobject_cast<QDateTimeEdit*>((QObject*)o); }
void* QDateTimeEdit_new(long long __goid, void* parent) {
	return new QDateTimeEdit__Go(__goid, (QWidget*)parent);
}
void* QDateTimeEdit_new2(long long __goid) {
	return new QDateTimeEdit__Go(__goid);
}
void* QDateTimeEdit_new3(long long __goid, void* dt) {
	return new QDateTimeEdit__Go(__goid, *(QDateTime*)dt);
}
void* QDateTimeEdit_new4(long long __goid, void* d) {
	return new QDateTimeEdit__Go(__goid, *(QDate*)d);
}
void* QDateTimeEdit_new5(long long __goid, void* t) {
	return new QDateTimeEdit__Go(__goid, *(QTime*)t);
}
void* QDateTimeEdit_new6(long long __goid, void* dt, void* parent) {
	return new QDateTimeEdit__Go(__goid, *(QDateTime*)dt, (QWidget*)parent);
}
void* QDateTimeEdit_new7(long long __goid, void* d, void* parent) {
	return new QDateTimeEdit__Go(__goid, *(QDate*)d, (QWidget*)parent);
}
void* QDateTimeEdit_new8(long long __goid, void* t, void* parent) {
	return new QDateTimeEdit__Go(__goid, *(QTime*)t, (QWidget*)parent);
}
char* QDateTimeEdit_tr(const char* s, long long* __out_len) {
	auto&& __r = QDateTimeEdit::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QDateTimeEdit_dateTime(void* self) {
	auto&& __r = ((QDateTimeEdit*)self)->dateTime();
	return new QDateTime(__r);
}
void* QDateTimeEdit_date(void* self) {
	auto&& __r = ((QDateTimeEdit*)self)->date();
	return new QDate(__r);
}
void* QDateTimeEdit_time(void* self) {
	auto&& __r = ((QDateTimeEdit*)self)->time();
	return new QTime(__r);
}
void* QDateTimeEdit_calendar(void* self) {
	auto&& __r = ((QDateTimeEdit*)self)->calendar();
	return new QCalendar(__r);
}
void QDateTimeEdit_setCalendar(void* self, void* calendar) {
	((QDateTimeEdit*)self)->setCalendar(*(QCalendar*)calendar);
}
void* QDateTimeEdit_minimumDateTime(void* self) {
	auto&& __r = ((QDateTimeEdit*)self)->minimumDateTime();
	return new QDateTime(__r);
}
void QDateTimeEdit_clearMinimumDateTime(void* self) {
	((QDateTimeEdit*)self)->clearMinimumDateTime();
}
void QDateTimeEdit_setMinimumDateTime(void* self, void* dt) {
	((QDateTimeEdit*)self)->setMinimumDateTime(*(QDateTime*)dt);
}
void* QDateTimeEdit_maximumDateTime(void* self) {
	auto&& __r = ((QDateTimeEdit*)self)->maximumDateTime();
	return new QDateTime(__r);
}
void QDateTimeEdit_clearMaximumDateTime(void* self) {
	((QDateTimeEdit*)self)->clearMaximumDateTime();
}
void QDateTimeEdit_setMaximumDateTime(void* self, void* dt) {
	((QDateTimeEdit*)self)->setMaximumDateTime(*(QDateTime*)dt);
}
void QDateTimeEdit_setDateTimeRange(void* self, void* min, void* max) {
	((QDateTimeEdit*)self)->setDateTimeRange(*(QDateTime*)min, *(QDateTime*)max);
}
void* QDateTimeEdit_minimumDate(void* self) {
	auto&& __r = ((QDateTimeEdit*)self)->minimumDate();
	return new QDate(__r);
}
void QDateTimeEdit_setMinimumDate(void* self, void* min) {
	((QDateTimeEdit*)self)->setMinimumDate(*(QDate*)min);
}
void QDateTimeEdit_clearMinimumDate(void* self) {
	((QDateTimeEdit*)self)->clearMinimumDate();
}
void* QDateTimeEdit_maximumDate(void* self) {
	auto&& __r = ((QDateTimeEdit*)self)->maximumDate();
	return new QDate(__r);
}
void QDateTimeEdit_setMaximumDate(void* self, void* max) {
	((QDateTimeEdit*)self)->setMaximumDate(*(QDate*)max);
}
void QDateTimeEdit_clearMaximumDate(void* self) {
	((QDateTimeEdit*)self)->clearMaximumDate();
}
void QDateTimeEdit_setDateRange(void* self, void* min, void* max) {
	((QDateTimeEdit*)self)->setDateRange(*(QDate*)min, *(QDate*)max);
}
void* QDateTimeEdit_minimumTime(void* self) {
	auto&& __r = ((QDateTimeEdit*)self)->minimumTime();
	return new QTime(__r);
}
void QDateTimeEdit_setMinimumTime(void* self, void* min) {
	((QDateTimeEdit*)self)->setMinimumTime(*(QTime*)min);
}
void QDateTimeEdit_clearMinimumTime(void* self) {
	((QDateTimeEdit*)self)->clearMinimumTime();
}
void* QDateTimeEdit_maximumTime(void* self) {
	auto&& __r = ((QDateTimeEdit*)self)->maximumTime();
	return new QTime(__r);
}
void QDateTimeEdit_setMaximumTime(void* self, void* max) {
	((QDateTimeEdit*)self)->setMaximumTime(*(QTime*)max);
}
void QDateTimeEdit_clearMaximumTime(void* self) {
	((QDateTimeEdit*)self)->clearMaximumTime();
}
void QDateTimeEdit_setTimeRange(void* self, void* min, void* max) {
	((QDateTimeEdit*)self)->setTimeRange(*(QTime*)min, *(QTime*)max);
}
int QDateTimeEdit_displayedSections(void* self) {
	auto&& __r = ((QDateTimeEdit*)self)->displayedSections();
	return (int)(__r);
}
int QDateTimeEdit_currentSection(void* self) {
	auto&& __r = ((QDateTimeEdit*)self)->currentSection();
	return (int)(__r);
}
int QDateTimeEdit_sectionAt(void* self, int index) {
	auto&& __r = ((QDateTimeEdit*)self)->sectionAt(index);
	return (int)(__r);
}
void QDateTimeEdit_setCurrentSection(void* self, int section) {
	((QDateTimeEdit*)self)->setCurrentSection((QDateTimeEdit::Section)(section));
}
int QDateTimeEdit_currentSectionIndex(void* self) {
	auto&& __r = ((QDateTimeEdit*)self)->currentSectionIndex();
	return __r;
}
void QDateTimeEdit_setCurrentSectionIndex(void* self, int index) {
	((QDateTimeEdit*)self)->setCurrentSectionIndex(index);
}
void* QDateTimeEdit_calendarWidget(void* self) {
	auto&& __r = ((QDateTimeEdit*)self)->calendarWidget();
	return (void*)(__r);
}
void QDateTimeEdit_setCalendarWidget(void* self, void* calendarWidget) {
	((QDateTimeEdit*)self)->setCalendarWidget((QCalendarWidget*)calendarWidget);
}
int QDateTimeEdit_sectionCount(void* self) {
	auto&& __r = ((QDateTimeEdit*)self)->sectionCount();
	return __r;
}
void QDateTimeEdit_setSelectedSection(void* self, int section) {
	((QDateTimeEdit*)self)->setSelectedSection((QDateTimeEdit::Section)(section));
}
char* QDateTimeEdit_sectionText(void* self, int section, long long* __out_len) {
	auto&& __r = ((QDateTimeEdit*)self)->sectionText((QDateTimeEdit::Section)(section));
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QDateTimeEdit_displayFormat(void* self, long long* __out_len) {
	auto&& __r = ((QDateTimeEdit*)self)->displayFormat();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QDateTimeEdit_setDisplayFormat(void* self, const char* format, long long format_len) {
	QString format_qstr = QString::fromUtf8(format, format_len);
	((QDateTimeEdit*)self)->setDisplayFormat(format_qstr);
}
int QDateTimeEdit_calendarPopup(void* self) {
	auto&& __r = ((QDateTimeEdit*)self)->calendarPopup();
	return __r;
}
void QDateTimeEdit_setCalendarPopup(void* self, int enable) {
	((QDateTimeEdit*)self)->setCalendarPopup(enable);
}
int QDateTimeEdit_timeSpec(void* self) {
	auto&& __r = ((QDateTimeEdit*)self)->timeSpec();
	return (int)(__r);
}
void QDateTimeEdit_setTimeSpec(void* self, int spec) {
	((QDateTimeEdit*)self)->setTimeSpec((Qt::TimeSpec)(spec));
}
void* QDateTimeEdit_timeZone(void* self) {
	auto&& __r = ((QDateTimeEdit*)self)->timeZone();
	return new QTimeZone(__r);
}
void QDateTimeEdit_setTimeZone(void* self, void* zone) {
	((QDateTimeEdit*)self)->setTimeZone(*(QTimeZone*)zone);
}
void QDateTimeEdit_dateTimeChanged(void* self, void* dateTime) {
	((QDateTimeEdit*)self)->dateTimeChanged(*(QDateTime*)dateTime);
}
void QDateTimeEdit_timeChanged(void* self, void* time) {
	((QDateTimeEdit*)self)->timeChanged(*(QTime*)time);
}
void QDateTimeEdit_dateChanged(void* self, void* date) {
	((QDateTimeEdit*)self)->dateChanged(*(QDate*)date);
}
void QDateTimeEdit_setDateTime(void* self, void* dateTime) {
	((QDateTimeEdit*)self)->setDateTime(*(QDateTime*)dateTime);
}
void QDateTimeEdit_setDate(void* self, void* date) {
	((QDateTimeEdit*)self)->setDate(*(QDate*)date);
}
void QDateTimeEdit_setTime(void* self, void* time) {
	((QDateTimeEdit*)self)->setTime(*(QTime*)time);
}
void* QDateTimeEdit_lineEdit(void* self) {
	auto&& __r = ((QDateTimeEdit__Go*)self)->lineEdit();
	return (void*)(__r);
}
void QDateTimeEdit_setLineEdit(void* self, void* edit) {
	((QDateTimeEdit__Go*)self)->setLineEdit((QLineEdit*)edit);
}
void QDateTimeEdit_setTabOrder(void* arg1, void* arg2) {
	QDateTimeEdit::setTabOrder((QWidget*)arg1, (QWidget*)arg2);
}
void* QDateTimeEdit_mouseGrabber() {
	auto&& __r = QDateTimeEdit::mouseGrabber();
	return (void*)(__r);
}
void* QDateTimeEdit_keyboardGrabber() {
	auto&& __r = QDateTimeEdit::keyboardGrabber();
	return (void*)(__r);
}
void* QDateTimeEdit_find(unsigned long long arg1) {
	auto&& __r = QDateTimeEdit::find(arg1);
	return (void*)(__r);
}
void* QDateTimeEdit_createWindowContainer(void* window) {
	auto&& __r = QDateTimeEdit::createWindowContainer((QWindow*)window);
	return (void*)(__r);
}
void QDateTimeEdit_updateMicroFocus(void* self) {
	((QDateTimeEdit__Go*)self)->updateMicroFocus();
}
void QDateTimeEdit_create(void* self) {
	((QDateTimeEdit__Go*)self)->create();
}
void QDateTimeEdit_destroy(void* self) {
	((QDateTimeEdit__Go*)self)->destroy();
}
int QDateTimeEdit_focusNextChild(void* self) {
	auto&& __r = ((QDateTimeEdit__Go*)self)->focusNextChild();
	return __r;
}
int QDateTimeEdit_focusPreviousChild(void* self) {
	auto&& __r = ((QDateTimeEdit__Go*)self)->focusPreviousChild();
	return __r;
}
void* QDateTimeEdit_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QDateTimeEdit::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QDateTimeEdit_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QDateTimeEdit::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QDateTimeEdit_sender(void* self) {
	auto&& __r = ((QDateTimeEdit__Go*)self)->sender();
	return (void*)(__r);
}
int QDateTimeEdit_senderSignalIndex(void* self) {
	auto&& __r = ((QDateTimeEdit__Go*)self)->senderSignalIndex();
	return __r;
}
int QDateTimeEdit_receivers(void* self, const char* signal) {
	auto&& __r = ((QDateTimeEdit__Go*)self)->receivers(signal);
	return __r;
}
int QDateTimeEdit_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QDateTimeEdit__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
double QDateTimeEdit_devicePixelRatioFScale() {
	auto&& __r = QDateTimeEdit::devicePixelRatioFScale();
	return __r;
}
int QDateTimeEdit_encodeMetricF(int metric, double value) {
	auto&& __r = QDateTimeEdit::encodeMetricF((QPaintDevice::PaintDeviceMetric)(metric), value);
	return __r;
}
double QDateTimeEdit_getDecodedMetricF(void* self, int metricA, int metricB) {
	auto&& __r = ((QDateTimeEdit__Go*)self)->getDecodedMetricF((QPaintDevice::PaintDeviceMetric)(metricA), (QPaintDevice::PaintDeviceMetric)(metricB));
	return __r;
}
char* QDateTimeEdit_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QDateTimeEdit::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QDateTimeEdit_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QDateTimeEdit::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QDateTimeEdit_createWindowContainer2(void* window, void* parent) {
	auto&& __r = QDateTimeEdit::createWindowContainer((QWindow*)window, (QWidget*)parent);
	return (void*)(__r);
}
void* QDateTimeEdit_createWindowContainer3(void* window, void* parent, int flags) {
	auto&& __r = QDateTimeEdit::createWindowContainer((QWindow*)window, (QWidget*)parent, QFlags<Qt::WindowType>(flags));
	return (void*)(__r);
}
void QDateTimeEdit_updateMicroFocus2(void* self, int query) {
	((QDateTimeEdit__Go*)self)->updateMicroFocus((Qt::InputMethodQuery)(query));
}
void QDateTimeEdit_create2(void* self, unsigned long long arg1) {
	((QDateTimeEdit__Go*)self)->create(arg1);
}
void QDateTimeEdit_create3(void* self, unsigned long long arg1, int initializeWindow) {
	((QDateTimeEdit__Go*)self)->create(arg1, initializeWindow);
}
void QDateTimeEdit_create4(void* self, unsigned long long arg1, int initializeWindow, int destroyOldWindow) {
	((QDateTimeEdit__Go*)self)->create(arg1, initializeWindow, destroyOldWindow);
}
void QDateTimeEdit_destroy2(void* self, int destroyWindow) {
	((QDateTimeEdit__Go*)self)->destroy(destroyWindow);
}
void QDateTimeEdit_destroy3(void* self, int destroyWindow, int destroySubWindows) {
	((QDateTimeEdit__Go*)self)->destroy(destroyWindow, destroySubWindows);
}
void* QDateTimeEdit_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QDateTimeEdit::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QDateTimeEdit_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QDateTimeEdit::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QDateTimeEdit_metaObject_default(void* self) {
	auto&& __r = ((QDateTimeEdit*)self)->QDateTimeEdit::metaObject();
	return (void*)(__r);
}
void* QDateTimeEdit_sizeHint_default(void* self) {
	auto&& __r = ((QDateTimeEdit*)self)->QDateTimeEdit::sizeHint();
	return new QSize(__r);
}
void QDateTimeEdit_clear_default(void* self) {
	((QDateTimeEdit*)self)->QDateTimeEdit::clear();
}
void QDateTimeEdit_stepBy_default(void* self, int steps) {
	((QDateTimeEdit*)self)->QDateTimeEdit::stepBy(steps);
}
int QDateTimeEdit_event_default(void* self, void* event) {
	auto&& __r = ((QDateTimeEdit*)self)->QDateTimeEdit::event((QEvent*)event);
	return __r;
}
void QDateTimeEdit_keyPressEvent_default(void* self, void* event) {
	((QDateTimeEdit__Go*)self)->cppgoBase_keyPressEvent((QKeyEvent*)event);
}
void QDateTimeEdit_wheelEvent_default(void* self, void* event) {
	((QDateTimeEdit__Go*)self)->cppgoBase_wheelEvent((QWheelEvent*)event);
}
void QDateTimeEdit_focusInEvent_default(void* self, void* event) {
	((QDateTimeEdit__Go*)self)->cppgoBase_focusInEvent((QFocusEvent*)event);
}
int QDateTimeEdit_focusNextPrevChild_default(void* self, int next) {
	auto&& __r = ((QDateTimeEdit__Go*)self)->cppgoBase_focusNextPrevChild(next);
	return __r;
}
int QDateTimeEdit_validate_default(void* self, void* input, int* pos) {
	int pos_ref = *pos;
	auto&& __r = ((QDateTimeEdit__Go*)self)->cppgoBase_validate(*(QString*)input, pos_ref);
	*pos = pos_ref;
	return (int)(__r);
}
void QDateTimeEdit_fixup_default(void* self, void* input) {
	((QDateTimeEdit__Go*)self)->cppgoBase_fixup(*(QString*)input);
}
void* QDateTimeEdit_dateTimeFromText_default(void* self, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = ((QDateTimeEdit__Go*)self)->cppgoBase_dateTimeFromText(text_qstr);
	return new QDateTime(__r);
}
char* QDateTimeEdit_textFromDateTime_default(void* self, void* dt, long long* __out_len) {
	auto&& __r = ((QDateTimeEdit__Go*)self)->cppgoBase_textFromDateTime(*(QDateTime*)dt);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QDateTimeEdit_stepEnabled_default(void* self) {
	auto&& __r = ((QDateTimeEdit__Go*)self)->cppgoBase_stepEnabled();
	return (int)(__r);
}
void QDateTimeEdit_mousePressEvent_default(void* self, void* event) {
	((QDateTimeEdit__Go*)self)->cppgoBase_mousePressEvent((QMouseEvent*)event);
}
void QDateTimeEdit_paintEvent_default(void* self, void* event) {
	((QDateTimeEdit__Go*)self)->cppgoBase_paintEvent((QPaintEvent*)event);
}
void QDateTimeEdit_initStyleOption_default(void* self, void* option) {
	((QDateTimeEdit__Go*)self)->cppgoBase_initStyleOption((QStyleOptionSpinBox*)option);
}
void* QDateTimeEdit_minimumSizeHint_default(void* self) {
	auto&& __r = ((QDateTimeEdit*)self)->QDateTimeEdit::minimumSizeHint();
	return new QSize(__r);
}
void* QDateTimeEdit_inputMethodQuery_default(void* self, int arg1) {
	auto&& __r = ((QDateTimeEdit*)self)->QDateTimeEdit::inputMethodQuery((Qt::InputMethodQuery)(arg1));
	return new QVariant(__r);
}
void QDateTimeEdit_resizeEvent_default(void* self, void* event) {
	((QDateTimeEdit__Go*)self)->cppgoBase_resizeEvent((QResizeEvent*)event);
}
void QDateTimeEdit_keyReleaseEvent_default(void* self, void* event) {
	((QDateTimeEdit__Go*)self)->cppgoBase_keyReleaseEvent((QKeyEvent*)event);
}
void QDateTimeEdit_focusOutEvent_default(void* self, void* event) {
	((QDateTimeEdit__Go*)self)->cppgoBase_focusOutEvent((QFocusEvent*)event);
}
void QDateTimeEdit_contextMenuEvent_default(void* self, void* event) {
	((QDateTimeEdit__Go*)self)->cppgoBase_contextMenuEvent((QContextMenuEvent*)event);
}
void QDateTimeEdit_changeEvent_default(void* self, void* event) {
	((QDateTimeEdit__Go*)self)->cppgoBase_changeEvent((QEvent*)event);
}
void QDateTimeEdit_closeEvent_default(void* self, void* event) {
	((QDateTimeEdit__Go*)self)->cppgoBase_closeEvent((QCloseEvent*)event);
}
void QDateTimeEdit_hideEvent_default(void* self, void* event) {
	((QDateTimeEdit__Go*)self)->cppgoBase_hideEvent((QHideEvent*)event);
}
void QDateTimeEdit_mouseReleaseEvent_default(void* self, void* event) {
	((QDateTimeEdit__Go*)self)->cppgoBase_mouseReleaseEvent((QMouseEvent*)event);
}
void QDateTimeEdit_mouseMoveEvent_default(void* self, void* event) {
	((QDateTimeEdit__Go*)self)->cppgoBase_mouseMoveEvent((QMouseEvent*)event);
}
void QDateTimeEdit_timerEvent_default(void* self, void* event) {
	((QDateTimeEdit__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QDateTimeEdit_showEvent_default(void* self, void* event) {
	((QDateTimeEdit__Go*)self)->cppgoBase_showEvent((QShowEvent*)event);
}
int QDateTimeEdit_devType_default(void* self) {
	auto&& __r = ((QDateTimeEdit*)self)->QDateTimeEdit::devType();
	return __r;
}
void QDateTimeEdit_setVisible_default(void* self, int visible) {
	((QDateTimeEdit*)self)->QDateTimeEdit::setVisible(visible);
}
int QDateTimeEdit_heightForWidth_default(void* self, int arg1) {
	auto&& __r = ((QDateTimeEdit*)self)->QDateTimeEdit::heightForWidth(arg1);
	return __r;
}
int QDateTimeEdit_hasHeightForWidth_default(void* self) {
	auto&& __r = ((QDateTimeEdit*)self)->QDateTimeEdit::hasHeightForWidth();
	return __r;
}
void* QDateTimeEdit_paintEngine_default(void* self) {
	auto&& __r = ((QDateTimeEdit*)self)->QDateTimeEdit::paintEngine();
	return (void*)(__r);
}
void QDateTimeEdit_mouseDoubleClickEvent_default(void* self, void* event) {
	((QDateTimeEdit__Go*)self)->cppgoBase_mouseDoubleClickEvent((QMouseEvent*)event);
}
void QDateTimeEdit_enterEvent_default(void* self, void* event) {
	((QDateTimeEdit__Go*)self)->cppgoBase_enterEvent((QEnterEvent*)event);
}
void QDateTimeEdit_leaveEvent_default(void* self, void* event) {
	((QDateTimeEdit__Go*)self)->cppgoBase_leaveEvent((QEvent*)event);
}
void QDateTimeEdit_moveEvent_default(void* self, void* event) {
	((QDateTimeEdit__Go*)self)->cppgoBase_moveEvent((QMoveEvent*)event);
}
void QDateTimeEdit_tabletEvent_default(void* self, void* event) {
	((QDateTimeEdit__Go*)self)->cppgoBase_tabletEvent((QTabletEvent*)event);
}
void QDateTimeEdit_actionEvent_default(void* self, void* event) {
	((QDateTimeEdit__Go*)self)->cppgoBase_actionEvent((QActionEvent*)event);
}
void QDateTimeEdit_dragEnterEvent_default(void* self, void* event) {
	((QDateTimeEdit__Go*)self)->cppgoBase_dragEnterEvent((QDragEnterEvent*)event);
}
void QDateTimeEdit_dragMoveEvent_default(void* self, void* event) {
	((QDateTimeEdit__Go*)self)->cppgoBase_dragMoveEvent((QDragMoveEvent*)event);
}
void QDateTimeEdit_dragLeaveEvent_default(void* self, void* event) {
	((QDateTimeEdit__Go*)self)->cppgoBase_dragLeaveEvent((QDragLeaveEvent*)event);
}
void QDateTimeEdit_dropEvent_default(void* self, void* event) {
	((QDateTimeEdit__Go*)self)->cppgoBase_dropEvent((QDropEvent*)event);
}
int QDateTimeEdit_metric_default(void* self, int arg1) {
	auto&& __r = ((QDateTimeEdit__Go*)self)->cppgoBase_metric(arg1);
	return __r;
}
void QDateTimeEdit_initPainter_default(void* self, void* painter) {
	((QDateTimeEdit__Go*)self)->cppgoBase_initPainter((QPainter*)painter);
}
void* QDateTimeEdit_redirected_default(void* self, void* offset) {
	auto&& __r = ((QDateTimeEdit__Go*)self)->cppgoBase_redirected((QPoint*)offset);
	return (void*)(__r);
}
void* QDateTimeEdit_sharedPainter_default(void* self) {
	auto&& __r = ((QDateTimeEdit__Go*)self)->cppgoBase_sharedPainter();
	return (void*)(__r);
}
void QDateTimeEdit_inputMethodEvent_default(void* self, void* arg1) {
	((QDateTimeEdit__Go*)self)->cppgoBase_inputMethodEvent((QInputMethodEvent*)arg1);
}
int QDateTimeEdit_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QDateTimeEdit*)self)->QDateTimeEdit::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QDateTimeEdit_childEvent_default(void* self, void* event) {
	((QDateTimeEdit__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QDateTimeEdit_customEvent_default(void* self, void* event) {
	((QDateTimeEdit__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QDateTimeEdit_connectNotify_default(void* self, void* signal) {
	((QDateTimeEdit__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QDateTimeEdit_disconnectNotify_default(void* self, void* signal) {
	((QDateTimeEdit__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QDateTimeEdit_connect_dateTimeChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDateTimeEdit*)self, &QDateTimeEdit::dateTimeChanged, (QDateTimeEdit*)self, [__id](const QDateTime& dateTime){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&dateTime);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDateTimeEdit);
	__conns_QDateTimeEdit[__id] = __c;
}
void QDateTimeEdit_connect_timeChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDateTimeEdit*)self, &QDateTimeEdit::timeChanged, (QDateTimeEdit*)self, [__id](QTime time){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&time);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDateTimeEdit);
	__conns_QDateTimeEdit[__id] = __c;
}
void QDateTimeEdit_connect_dateChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDateTimeEdit*)self, &QDateTimeEdit::dateChanged, (QDateTimeEdit*)self, [__id](QDate date){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&date);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDateTimeEdit);
	__conns_QDateTimeEdit[__id] = __c;
}
void QDateTimeEdit_connect_editingFinished(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDateTimeEdit*)self, &QDateTimeEdit::editingFinished, (QDateTimeEdit*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDateTimeEdit);
	__conns_QDateTimeEdit[__id] = __c;
}
void QDateTimeEdit_connect_returnPressed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDateTimeEdit*)self, &QDateTimeEdit::returnPressed, (QDateTimeEdit*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDateTimeEdit);
	__conns_QDateTimeEdit[__id] = __c;
}
void QDateTimeEdit_connect_windowTitleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDateTimeEdit*)self, &QDateTimeEdit::windowTitleChanged, (QDateTimeEdit*)self, [__id](const QString& title){
		QByteArray title_ba = title.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)title_ba.constData(), title_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDateTimeEdit);
	__conns_QDateTimeEdit[__id] = __c;
}
void QDateTimeEdit_connect_windowIconChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDateTimeEdit*)self, &QDateTimeEdit::windowIconChanged, (QDateTimeEdit*)self, [__id](const QIcon& icon){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&icon);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDateTimeEdit);
	__conns_QDateTimeEdit[__id] = __c;
}
void QDateTimeEdit_connect_windowIconTextChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDateTimeEdit*)self, &QDateTimeEdit::windowIconTextChanged, (QDateTimeEdit*)self, [__id](const QString& iconText){
		QByteArray iconText_ba = iconText.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)iconText_ba.constData(), iconText_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDateTimeEdit);
	__conns_QDateTimeEdit[__id] = __c;
}
void QDateTimeEdit_connect_customContextMenuRequested(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDateTimeEdit*)self, &QDateTimeEdit::customContextMenuRequested, (QDateTimeEdit*)self, [__id](const QPoint& pos){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&pos);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDateTimeEdit);
	__conns_QDateTimeEdit[__id] = __c;
}
void QDateTimeEdit_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDateTimeEdit*)self, &QDateTimeEdit::destroyed, (QDateTimeEdit*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDateTimeEdit);
	__conns_QDateTimeEdit[__id] = __c;
}
void QDateTimeEdit_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDateTimeEdit*)self, &QDateTimeEdit::destroyed, (QDateTimeEdit*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDateTimeEdit);
	__conns_QDateTimeEdit[__id] = __c;
}
void QDateTimeEdit_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDateTimeEdit*)self, &QDateTimeEdit::objectNameChanged, (QDateTimeEdit*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDateTimeEdit);
	__conns_QDateTimeEdit[__id] = __c;
}
void QDateTimeEdit__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QDateTimeEdit);
		auto __it = __conns_QDateTimeEdit.find(id);
		if (__it == __conns_QDateTimeEdit.end()) return;
		__c = __it->second;
		__conns_QDateTimeEdit.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QDateTimeEdit__Go_new(long long id) { return new QDateTimeEdit__Go(id); }
void QDateTimeEdit_delete(void* self) { delete (QDateTimeEdit*)self; }

}

#include <QtWidgets/qgraphicsitem.h>
#include <QtCore/qobject.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qvariant.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qbindingstorage.h>
#include <QtWidgets/qgraphicsscene.h>
#include <QtWidgets/qgraphicswidget.h>
#include <QtCore/qsize.h>
#include <QtGui/qcursor.h>
#include <QtWidgets/qgraphicseffect.h>
#include <QtCore/qpoint.h>
#include <QtCore/qrect.h>
#include <QtGui/qtransform.h>
#include <QtWidgets/qgraphicstransform.h>
#include <QtGui/qpainterpath.h>
#include <QtGui/qregion.h>
#include <QtGui/qpainter.h>
#include <QtWidgets/qstyleoption.h>
#include <QtWidgets/qwidget.h>
#include <QtGui/qpolygon.h>
#include <QtWidgets/qgraphicssceneevent.h>
#include <QtGui/qevent.h>
#include "gen_qgraphicsobject.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QGraphicsObject__Go : public QGraphicsObject {
public:
	long long __goId;
	unsigned char __goOn[43] = {};
	template<typename... A> QGraphicsObject__Go(long long id, A&&... a) : QGraphicsObject(std::forward<A>(a)...), __goId(id) {}
	using QGraphicsObject::updateMicroFocus;
	using QGraphicsObject::sender;
	using QGraphicsObject::senderSignalIndex;
	using QGraphicsObject::receivers;
	using QGraphicsObject::isSignalConnected;
	using QGraphicsObject::addToIndex;
	using QGraphicsObject::removeFromIndex;
	using QGraphicsObject::prepareGeometryChange;
	bool cppgoBase_event(QEvent* __a0) { return QGraphicsObject::event(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QGraphicsObject::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QGraphicsObject::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QGraphicsObject::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QGraphicsObject::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QGraphicsObject::disconnectNotify(__a0); }
	bool cppgoBase_sceneEventFilter(QGraphicsItem* __a0, QEvent* __a1) { return QGraphicsObject::sceneEventFilter(__a0, __a1); }
	bool cppgoBase_sceneEvent(QEvent* __a0) { return QGraphicsObject::sceneEvent(__a0); }
	void cppgoBase_contextMenuEvent(QGraphicsSceneContextMenuEvent* __a0) { QGraphicsObject::contextMenuEvent(__a0); }
	void cppgoBase_dragEnterEvent(QGraphicsSceneDragDropEvent* __a0) { QGraphicsObject::dragEnterEvent(__a0); }
	void cppgoBase_dragLeaveEvent(QGraphicsSceneDragDropEvent* __a0) { QGraphicsObject::dragLeaveEvent(__a0); }
	void cppgoBase_dragMoveEvent(QGraphicsSceneDragDropEvent* __a0) { QGraphicsObject::dragMoveEvent(__a0); }
	void cppgoBase_dropEvent(QGraphicsSceneDragDropEvent* __a0) { QGraphicsObject::dropEvent(__a0); }
	void cppgoBase_focusInEvent(QFocusEvent* __a0) { QGraphicsObject::focusInEvent(__a0); }
	void cppgoBase_focusOutEvent(QFocusEvent* __a0) { QGraphicsObject::focusOutEvent(__a0); }
	void cppgoBase_hoverEnterEvent(QGraphicsSceneHoverEvent* __a0) { QGraphicsObject::hoverEnterEvent(__a0); }
	void cppgoBase_hoverMoveEvent(QGraphicsSceneHoverEvent* __a0) { QGraphicsObject::hoverMoveEvent(__a0); }
	void cppgoBase_hoverLeaveEvent(QGraphicsSceneHoverEvent* __a0) { QGraphicsObject::hoverLeaveEvent(__a0); }
	void cppgoBase_keyPressEvent(QKeyEvent* __a0) { QGraphicsObject::keyPressEvent(__a0); }
	void cppgoBase_keyReleaseEvent(QKeyEvent* __a0) { QGraphicsObject::keyReleaseEvent(__a0); }
	void cppgoBase_mousePressEvent(QGraphicsSceneMouseEvent* __a0) { QGraphicsObject::mousePressEvent(__a0); }
	void cppgoBase_mouseMoveEvent(QGraphicsSceneMouseEvent* __a0) { QGraphicsObject::mouseMoveEvent(__a0); }
	void cppgoBase_mouseReleaseEvent(QGraphicsSceneMouseEvent* __a0) { QGraphicsObject::mouseReleaseEvent(__a0); }
	void cppgoBase_mouseDoubleClickEvent(QGraphicsSceneMouseEvent* __a0) { QGraphicsObject::mouseDoubleClickEvent(__a0); }
	void cppgoBase_wheelEvent(QGraphicsSceneWheelEvent* __a0) { QGraphicsObject::wheelEvent(__a0); }
	void cppgoBase_inputMethodEvent(QInputMethodEvent* __a0) { QGraphicsObject::inputMethodEvent(__a0); }
	QVariant cppgoBase_inputMethodQuery(int __a0) { return QGraphicsObject::inputMethodQuery((Qt::InputMethodQuery)__a0); }
	QVariant cppgoBase_itemChange(int __a0, const QVariant& __a1) { return QGraphicsObject::itemChange((QGraphicsItem::GraphicsItemChange)__a0, __a1); }
	bool cppgoBase_supportsExtension(int __a0) { return QGraphicsObject::supportsExtension((QGraphicsItem::Extension)__a0); }
	void cppgoBase_setExtension(int __a0, const QVariant& __a1) { QGraphicsObject::setExtension((QGraphicsItem::Extension)__a0, __a1); }
	QVariant cppgoBase_extension(const QVariant& __a0) { return QGraphicsObject::extension(__a0); }
	QRectF boundingRect() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 2510008782785038831LL, &__h); if (__h) return *(QRectF*)__r; }
		return QRectF{};
	}
	void paint(QPainter* painter, const QStyleOptionGraphicsItem* option, QWidget* widget) override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_void_voidp_voidp_voidp) { ((void(*)(long long, long long, void*, void*, void*, int*))cppgo_virt_void_voidp_voidp_voidp)(__goId, 6083584261053556819LL, (void*)painter, (void*)option, (void*)widget, &__h); if (__h) return; }
		return;
	}
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QGraphicsObject::metaObject();
	}
	bool event(QEvent* ev) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)ev, &__h); if (__h) return __r; }
		return QGraphicsObject::event(ev);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QGraphicsObject::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QGraphicsObject::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QGraphicsObject::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QGraphicsObject::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QGraphicsObject::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QGraphicsObject::disconnectNotify(signal);
	}
	void advance(int phase) override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, -4830836341480073475LL, phase, &__h); if (__h) return; }
		QGraphicsObject::advance(phase);
	}
	QPainterPath shape() const override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 7938019069247444294LL, &__h); if (__h) return *(QPainterPath*)__r; }
		return QGraphicsObject::shape();
	}
	bool contains(const QPointF& point) const override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 5645366511944179768LL, (void*)&point, &__h); if (__h) return __r; }
		return QGraphicsObject::contains(point);
	}
	bool collidesWithItem(const QGraphicsItem* other, Qt::ItemSelectionMode mode) const override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_int_voidp_int) { int __r = ((int(*)(long long, long long, void*, int, int*))cppgo_virt_int_voidp_int)(__goId, 6469736184362593173LL, (void*)other, (int)(mode), &__h); if (__h) return __r; }
		return QGraphicsObject::collidesWithItem(other, mode);
	}
	bool collidesWithPath(const QPainterPath& path, Qt::ItemSelectionMode mode) const override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_int_voidp_int) { int __r = ((int(*)(long long, long long, void*, int, int*))cppgo_virt_int_voidp_int)(__goId, 6636575182718753853LL, (void*)&path, (int)(mode), &__h); if (__h) return __r; }
		return QGraphicsObject::collidesWithPath(path, mode);
	}
	bool isObscuredBy(const QGraphicsItem* item) const override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, -7215030271914903671LL, (void*)item, &__h); if (__h) return __r; }
		return QGraphicsObject::isObscuredBy(item);
	}
	QPainterPath opaqueArea() const override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 399370970023989435LL, &__h); if (__h) return *(QPainterPath*)__r; }
		return QGraphicsObject::opaqueArea();
	}
	int type() const override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -6371403872082420691LL, &__h); if (__h) return __r; }
		return QGraphicsObject::type();
	}
	bool sceneEventFilter(QGraphicsItem* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, -4584203218241215725LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QGraphicsObject::sceneEventFilter(watched, event);
	}
	bool sceneEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, -2846371507838587071LL, (void*)event, &__h); if (__h) return __r; }
		return QGraphicsObject::sceneEvent(event);
	}
	void contextMenuEvent(QGraphicsSceneContextMenuEvent* event) override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6698043222227006015LL, (void*)event, &__h); if (__h) return; }
		QGraphicsObject::contextMenuEvent(event);
	}
	void dragEnterEvent(QGraphicsSceneDragDropEvent* event) override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 206122715919917589LL, (void*)event, &__h); if (__h) return; }
		QGraphicsObject::dragEnterEvent(event);
	}
	void dragLeaveEvent(QGraphicsSceneDragDropEvent* event) override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -3139904158373725006LL, (void*)event, &__h); if (__h) return; }
		QGraphicsObject::dragLeaveEvent(event);
	}
	void dragMoveEvent(QGraphicsSceneDragDropEvent* event) override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1781545137352745978LL, (void*)event, &__h); if (__h) return; }
		QGraphicsObject::dragMoveEvent(event);
	}
	void dropEvent(QGraphicsSceneDragDropEvent* event) override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 9107058682754591716LL, (void*)event, &__h); if (__h) return; }
		QGraphicsObject::dropEvent(event);
	}
	void focusInEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 975043767033552552LL, (void*)event, &__h); if (__h) return; }
		QGraphicsObject::focusInEvent(event);
	}
	void focusOutEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -385596010576048993LL, (void*)event, &__h); if (__h) return; }
		QGraphicsObject::focusOutEvent(event);
	}
	void hoverEnterEvent(QGraphicsSceneHoverEvent* event) override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6932941943906681679LL, (void*)event, &__h); if (__h) return; }
		QGraphicsObject::hoverEnterEvent(event);
	}
	void hoverMoveEvent(QGraphicsSceneHoverEvent* event) override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -2745191522144316396LL, (void*)event, &__h); if (__h) return; }
		QGraphicsObject::hoverMoveEvent(event);
	}
	void hoverLeaveEvent(QGraphicsSceneHoverEvent* event) override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7944408482203603984LL, (void*)event, &__h); if (__h) return; }
		QGraphicsObject::hoverLeaveEvent(event);
	}
	void keyPressEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1023791977679793185LL, (void*)event, &__h); if (__h) return; }
		QGraphicsObject::keyPressEvent(event);
	}
	void keyReleaseEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 772598475950050299LL, (void*)event, &__h); if (__h) return; }
		QGraphicsObject::keyReleaseEvent(event);
	}
	void mousePressEvent(QGraphicsSceneMouseEvent* event) override {
		int __h = 0;
		if (__goOn[32] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -332090898078543319LL, (void*)event, &__h); if (__h) return; }
		QGraphicsObject::mousePressEvent(event);
	}
	void mouseMoveEvent(QGraphicsSceneMouseEvent* event) override {
		int __h = 0;
		if (__goOn[33] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4255494004099806581LL, (void*)event, &__h); if (__h) return; }
		QGraphicsObject::mouseMoveEvent(event);
	}
	void mouseReleaseEvent(QGraphicsSceneMouseEvent* event) override {
		int __h = 0;
		if (__goOn[34] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4415920049297795213LL, (void*)event, &__h); if (__h) return; }
		QGraphicsObject::mouseReleaseEvent(event);
	}
	void mouseDoubleClickEvent(QGraphicsSceneMouseEvent* event) override {
		int __h = 0;
		if (__goOn[35] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4029073905697023887LL, (void*)event, &__h); if (__h) return; }
		QGraphicsObject::mouseDoubleClickEvent(event);
	}
	void wheelEvent(QGraphicsSceneWheelEvent* event) override {
		int __h = 0;
		if (__goOn[36] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1919988340934511316LL, (void*)event, &__h); if (__h) return; }
		QGraphicsObject::wheelEvent(event);
	}
	void inputMethodEvent(QInputMethodEvent* event) override {
		int __h = 0;
		if (__goOn[37] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5438819661942843222LL, (void*)event, &__h); if (__h) return; }
		QGraphicsObject::inputMethodEvent(event);
	}
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const override {
		int __h = 0;
		if (__goOn[38] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7345853097340709638LL, (int)(query), &__h); if (__h) return *(QVariant*)__r; }
		return QGraphicsObject::inputMethodQuery(query);
	}
	QVariant itemChange(QGraphicsItem::GraphicsItemChange change, const QVariant& value) override {
		int __h = 0;
		if (__goOn[39] && cppgo_virt_voidp_int_voidp) { void* __r = ((void*(*)(long long, long long, int, void*, int*))cppgo_virt_voidp_int_voidp)(__goId, -1888572168794878830LL, (int)(change), (void*)&value, &__h); if (__h) return *(QVariant*)__r; }
		return QGraphicsObject::itemChange(change, value);
	}
	bool supportsExtension(QGraphicsItem::Extension extension) const override {
		int __h = 0;
		if (__goOn[40] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, -5729166443651306418LL, (int)(extension), &__h); if (__h) return __r; }
		return QGraphicsObject::supportsExtension(extension);
	}
	void setExtension(QGraphicsItem::Extension extension, const QVariant& variant) override {
		int __h = 0;
		if (__goOn[41] && cppgo_virt_void_int_voidp) { ((void(*)(long long, long long, int, void*, int*))cppgo_virt_void_int_voidp)(__goId, -6730181576106171580LL, (int)(extension), (void*)&variant, &__h); if (__h) return; }
		QGraphicsObject::setExtension(extension, variant);
	}
	QVariant extension(const QVariant& variant) const override {
		int __h = 0;
		if (__goOn[42] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 2270616358835251226LL, (void*)&variant, &__h); if (__h) return *(QVariant*)__r; }
		return QGraphicsObject::extension(variant);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QGraphicsObject;
static std::mutex __connsMu_QGraphicsObject;

extern "C" {

void* QGraphicsObject_base(void* self) { return (void*)static_cast<QObject*>((QGraphicsObject*)self); }
void QGraphicsObject__set_override(void* self, int idx) { ((QGraphicsObject__Go*)self)->__goOn[idx] = 1; }
void* QGraphicsObject_fromQObject(void* o) { return (void*)qobject_cast<QGraphicsObject*>((QObject*)o); }
char* QGraphicsObject_tr(const char* s, long long* __out_len) {
	auto&& __r = QGraphicsObject::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QGraphicsObject_grabGesture(void* self, int type_) {
	((QGraphicsObject*)self)->grabGesture((Qt::GestureType)(type_));
}
void QGraphicsObject_ungrabGesture(void* self, int type_) {
	((QGraphicsObject*)self)->ungrabGesture((Qt::GestureType)(type_));
}
void QGraphicsObject_updateMicroFocus(void* self) {
	((QGraphicsObject__Go*)self)->updateMicroFocus();
}
void QGraphicsObject_parentChanged(void* self) {
	((QGraphicsObject*)self)->parentChanged();
}
void QGraphicsObject_opacityChanged(void* self) {
	((QGraphicsObject*)self)->opacityChanged();
}
void QGraphicsObject_visibleChanged(void* self) {
	((QGraphicsObject*)self)->visibleChanged();
}
void QGraphicsObject_enabledChanged(void* self) {
	((QGraphicsObject*)self)->enabledChanged();
}
void QGraphicsObject_xChanged(void* self) {
	((QGraphicsObject*)self)->xChanged();
}
void QGraphicsObject_yChanged(void* self) {
	((QGraphicsObject*)self)->yChanged();
}
void QGraphicsObject_zChanged(void* self) {
	((QGraphicsObject*)self)->zChanged();
}
void QGraphicsObject_rotationChanged(void* self) {
	((QGraphicsObject*)self)->rotationChanged();
}
void QGraphicsObject_scaleChanged(void* self) {
	((QGraphicsObject*)self)->scaleChanged();
}
void QGraphicsObject_childrenChanged(void* self) {
	((QGraphicsObject*)self)->childrenChanged();
}
void QGraphicsObject_widthChanged(void* self) {
	((QGraphicsObject*)self)->widthChanged();
}
void QGraphicsObject_heightChanged(void* self) {
	((QGraphicsObject*)self)->heightChanged();
}
void* QGraphicsObject_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QGraphicsObject::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QGraphicsObject_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QGraphicsObject::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QGraphicsObject_sender(void* self) {
	auto&& __r = ((QGraphicsObject__Go*)self)->sender();
	return (void*)(__r);
}
int QGraphicsObject_senderSignalIndex(void* self) {
	auto&& __r = ((QGraphicsObject__Go*)self)->senderSignalIndex();
	return __r;
}
int QGraphicsObject_receivers(void* self, const char* signal) {
	auto&& __r = ((QGraphicsObject__Go*)self)->receivers(signal);
	return __r;
}
int QGraphicsObject_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QGraphicsObject__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
void* QGraphicsObject_scene(void* self) {
	auto&& __r = ((QGraphicsObject*)self)->scene();
	return (void*)(__r);
}
void* QGraphicsObject_parentItem(void* self) {
	auto&& __r = ((QGraphicsObject*)self)->parentItem();
	return (void*)(__r);
}
void* QGraphicsObject_topLevelItem(void* self) {
	auto&& __r = ((QGraphicsObject*)self)->topLevelItem();
	return (void*)(__r);
}
void* QGraphicsObject_parentObject(void* self) {
	auto&& __r = ((QGraphicsObject*)self)->parentObject();
	return (void*)(__r);
}
void* QGraphicsObject_parentWidget(void* self) {
	auto&& __r = ((QGraphicsObject*)self)->parentWidget();
	return (void*)(__r);
}
void* QGraphicsObject_topLevelWidget(void* self) {
	auto&& __r = ((QGraphicsObject*)self)->topLevelWidget();
	return (void*)(__r);
}
void* QGraphicsObject_window(void* self) {
	auto&& __r = ((QGraphicsObject*)self)->window();
	return (void*)(__r);
}
void* QGraphicsObject_panel(void* self) {
	auto&& __r = ((QGraphicsObject*)self)->panel();
	return (void*)(__r);
}
void QGraphicsObject_setParentItem(void* self, void* parent) {
	((QGraphicsObject*)self)->setParentItem((QGraphicsItem*)parent);
}
void** QGraphicsObject_childItems(void* self, long long* __out_len) {
	auto&& __r = ((QGraphicsObject*)self)->childItems();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
int QGraphicsObject_isWidget(void* self) {
	auto&& __r = ((QGraphicsObject*)self)->isWidget();
	return __r;
}
int QGraphicsObject_isWindow(void* self) {
	auto&& __r = ((QGraphicsObject*)self)->isWindow();
	return __r;
}
int QGraphicsObject_isPanel(void* self) {
	auto&& __r = ((QGraphicsObject*)self)->isPanel();
	return __r;
}
void* QGraphicsObject_toGraphicsObject(void* self) {
	auto&& __r = ((QGraphicsObject*)self)->toGraphicsObject();
	return (void*)(__r);
}
void* QGraphicsObject_toGraphicsObject2(void* self) {
	auto&& __r = ((QGraphicsObject*)self)->toGraphicsObject();
	return (void*)(__r);
}
void* QGraphicsObject_group(void* self) {
	auto&& __r = ((QGraphicsObject*)self)->group();
	return (void*)(__r);
}
void QGraphicsObject_setGroup(void* self, void* group) {
	((QGraphicsObject*)self)->setGroup((QGraphicsItemGroup*)group);
}
int QGraphicsObject_flags(void* self) {
	auto&& __r = ((QGraphicsObject*)self)->flags();
	return (int)(__r);
}
void QGraphicsObject_setFlag(void* self, int flag) {
	((QGraphicsObject*)self)->setFlag((QGraphicsItem::GraphicsItemFlag)(flag));
}
void QGraphicsObject_setFlags(void* self, int flags) {
	((QGraphicsObject*)self)->setFlags(QFlags<QGraphicsItem::GraphicsItemFlag>(flags));
}
int QGraphicsObject_cacheMode(void* self) {
	auto&& __r = ((QGraphicsObject*)self)->cacheMode();
	return (int)(__r);
}
void QGraphicsObject_setCacheMode(void* self, int mode) {
	((QGraphicsObject*)self)->setCacheMode((QGraphicsItem::CacheMode)(mode));
}
int QGraphicsObject_panelModality(void* self) {
	auto&& __r = ((QGraphicsObject*)self)->panelModality();
	return (int)(__r);
}
void QGraphicsObject_setPanelModality(void* self, int panelModality) {
	((QGraphicsObject*)self)->setPanelModality((QGraphicsItem::PanelModality)(panelModality));
}
int QGraphicsObject_isBlockedByModalPanel(void* self) {
	auto&& __r = ((QGraphicsObject*)self)->isBlockedByModalPanel();
	return __r;
}
char* QGraphicsObject_toolTip(void* self, long long* __out_len) {
	auto&& __r = ((QGraphicsObject*)self)->toolTip();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QGraphicsObject_setToolTip(void* self, const char* toolTip, long long toolTip_len) {
	QString toolTip_qstr = QString::fromUtf8(toolTip, toolTip_len);
	((QGraphicsObject*)self)->setToolTip(toolTip_qstr);
}
void* QGraphicsObject_cursor(void* self) {
	auto&& __r = ((QGraphicsObject*)self)->cursor();
	return new QCursor(__r);
}
void QGraphicsObject_setCursor(void* self, void* cursor) {
	((QGraphicsObject*)self)->setCursor(*(QCursor*)cursor);
}
int QGraphicsObject_hasCursor(void* self) {
	auto&& __r = ((QGraphicsObject*)self)->hasCursor();
	return __r;
}
void QGraphicsObject_unsetCursor(void* self) {
	((QGraphicsObject*)self)->unsetCursor();
}
int QGraphicsObject_isVisible(void* self) {
	auto&& __r = ((QGraphicsObject*)self)->isVisible();
	return __r;
}
int QGraphicsObject_isVisibleTo(void* self, void* parent) {
	auto&& __r = ((QGraphicsObject*)self)->isVisibleTo((QGraphicsItem*)parent);
	return __r;
}
void QGraphicsObject_setVisible(void* self, int visible) {
	((QGraphicsObject*)self)->setVisible(visible);
}
void QGraphicsObject_hide(void* self) {
	((QGraphicsObject*)self)->hide();
}
void QGraphicsObject_show(void* self) {
	((QGraphicsObject*)self)->show();
}
int QGraphicsObject_isEnabled(void* self) {
	auto&& __r = ((QGraphicsObject*)self)->isEnabled();
	return __r;
}
void QGraphicsObject_setEnabled(void* self, int enabled) {
	((QGraphicsObject*)self)->setEnabled(enabled);
}
int QGraphicsObject_isSelected(void* self) {
	auto&& __r = ((QGraphicsObject*)self)->isSelected();
	return __r;
}
void QGraphicsObject_setSelected(void* self, int selected) {
	((QGraphicsObject*)self)->setSelected(selected);
}
int QGraphicsObject_acceptDrops(void* self) {
	auto&& __r = ((QGraphicsObject*)self)->acceptDrops();
	return __r;
}
void QGraphicsObject_setAcceptDrops(void* self, int on) {
	((QGraphicsObject*)self)->setAcceptDrops(on);
}
double QGraphicsObject_opacity(void* self) {
	auto&& __r = ((QGraphicsObject*)self)->opacity();
	return __r;
}
double QGraphicsObject_effectiveOpacity(void* self) {
	auto&& __r = ((QGraphicsObject*)self)->effectiveOpacity();
	return __r;
}
void QGraphicsObject_setOpacity(void* self, double opacity) {
	((QGraphicsObject*)self)->setOpacity(opacity);
}
void* QGraphicsObject_graphicsEffect(void* self) {
	auto&& __r = ((QGraphicsObject*)self)->graphicsEffect();
	return (void*)(__r);
}
void QGraphicsObject_setGraphicsEffect(void* self, void* effect) {
	((QGraphicsObject*)self)->setGraphicsEffect((QGraphicsEffect*)effect);
}
int QGraphicsObject_acceptedMouseButtons(void* self) {
	auto&& __r = ((QGraphicsObject*)self)->acceptedMouseButtons();
	return (int)(__r);
}
void QGraphicsObject_setAcceptedMouseButtons(void* self, int buttons) {
	((QGraphicsObject*)self)->setAcceptedMouseButtons(QFlags<Qt::MouseButton>(buttons));
}
int QGraphicsObject_acceptHoverEvents(void* self) {
	auto&& __r = ((QGraphicsObject*)self)->acceptHoverEvents();
	return __r;
}
void QGraphicsObject_setAcceptHoverEvents(void* self, int enabled) {
	((QGraphicsObject*)self)->setAcceptHoverEvents(enabled);
}
int QGraphicsObject_acceptTouchEvents(void* self) {
	auto&& __r = ((QGraphicsObject*)self)->acceptTouchEvents();
	return __r;
}
void QGraphicsObject_setAcceptTouchEvents(void* self, int enabled) {
	((QGraphicsObject*)self)->setAcceptTouchEvents(enabled);
}
int QGraphicsObject_filtersChildEvents(void* self) {
	auto&& __r = ((QGraphicsObject*)self)->filtersChildEvents();
	return __r;
}
void QGraphicsObject_setFiltersChildEvents(void* self, int enabled) {
	((QGraphicsObject*)self)->setFiltersChildEvents(enabled);
}
int QGraphicsObject_handlesChildEvents(void* self) {
	auto&& __r = ((QGraphicsObject*)self)->handlesChildEvents();
	return __r;
}
void QGraphicsObject_setHandlesChildEvents(void* self, int enabled) {
	((QGraphicsObject*)self)->setHandlesChildEvents(enabled);
}
int QGraphicsObject_isActive(void* self) {
	auto&& __r = ((QGraphicsObject*)self)->isActive();
	return __r;
}
void QGraphicsObject_setActive(void* self, int active) {
	((QGraphicsObject*)self)->setActive(active);
}
int QGraphicsObject_hasFocus(void* self) {
	auto&& __r = ((QGraphicsObject*)self)->hasFocus();
	return __r;
}
void QGraphicsObject_setFocus(void* self) {
	((QGraphicsObject*)self)->setFocus();
}
void QGraphicsObject_clearFocus(void* self) {
	((QGraphicsObject*)self)->clearFocus();
}
void* QGraphicsObject_focusProxy(void* self) {
	auto&& __r = ((QGraphicsObject*)self)->focusProxy();
	return (void*)(__r);
}
void QGraphicsObject_setFocusProxy(void* self, void* item) {
	((QGraphicsObject*)self)->setFocusProxy((QGraphicsItem*)item);
}
void* QGraphicsObject_focusItem(void* self) {
	auto&& __r = ((QGraphicsObject*)self)->focusItem();
	return (void*)(__r);
}
void* QGraphicsObject_focusScopeItem(void* self) {
	auto&& __r = ((QGraphicsObject*)self)->focusScopeItem();
	return (void*)(__r);
}
void QGraphicsObject_grabMouse(void* self) {
	((QGraphicsObject*)self)->grabMouse();
}
void QGraphicsObject_ungrabMouse(void* self) {
	((QGraphicsObject*)self)->ungrabMouse();
}
void QGraphicsObject_grabKeyboard(void* self) {
	((QGraphicsObject*)self)->grabKeyboard();
}
void QGraphicsObject_ungrabKeyboard(void* self) {
	((QGraphicsObject*)self)->ungrabKeyboard();
}
void* QGraphicsObject_pos(void* self) {
	auto&& __r = ((QGraphicsObject*)self)->pos();
	return new QPointF(__r);
}
double QGraphicsObject_x(void* self) {
	auto&& __r = ((QGraphicsObject*)self)->x();
	return __r;
}
void QGraphicsObject_setX(void* self, double x) {
	((QGraphicsObject*)self)->setX(x);
}
double QGraphicsObject_y(void* self) {
	auto&& __r = ((QGraphicsObject*)self)->y();
	return __r;
}
void QGraphicsObject_setY(void* self, double y) {
	((QGraphicsObject*)self)->setY(y);
}
void* QGraphicsObject_scenePos(void* self) {
	auto&& __r = ((QGraphicsObject*)self)->scenePos();
	return new QPointF(__r);
}
void QGraphicsObject_setPos(void* self, void* pos) {
	((QGraphicsObject*)self)->setPos(*(QPointF*)pos);
}
void QGraphicsObject_setPos2(void* self, double x, double y) {
	((QGraphicsObject*)self)->setPos(x, y);
}
void QGraphicsObject_moveBy(void* self, double dx, double dy) {
	((QGraphicsObject*)self)->moveBy(dx, dy);
}
void QGraphicsObject_ensureVisible(void* self) {
	((QGraphicsObject*)self)->ensureVisible();
}
void QGraphicsObject_ensureVisible2(void* self, double x, double y, double w, double h) {
	((QGraphicsObject*)self)->ensureVisible(x, y, w, h);
}
void* QGraphicsObject_transform(void* self) {
	auto&& __r = ((QGraphicsObject*)self)->transform();
	return new QTransform(__r);
}
void* QGraphicsObject_sceneTransform(void* self) {
	auto&& __r = ((QGraphicsObject*)self)->sceneTransform();
	return new QTransform(__r);
}
void* QGraphicsObject_deviceTransform(void* self, void* viewportTransform) {
	auto&& __r = ((QGraphicsObject*)self)->deviceTransform(*(QTransform*)viewportTransform);
	return new QTransform(__r);
}
void* QGraphicsObject_itemTransform(void* self, void* other) {
	auto&& __r = ((QGraphicsObject*)self)->itemTransform((QGraphicsItem*)other);
	return new QTransform(__r);
}
void QGraphicsObject_setTransform(void* self, void* matrix) {
	((QGraphicsObject*)self)->setTransform(*(QTransform*)matrix);
}
void QGraphicsObject_resetTransform(void* self) {
	((QGraphicsObject*)self)->resetTransform();
}
void QGraphicsObject_setRotation(void* self, double angle) {
	((QGraphicsObject*)self)->setRotation(angle);
}
double QGraphicsObject_rotation(void* self) {
	auto&& __r = ((QGraphicsObject*)self)->rotation();
	return __r;
}
void QGraphicsObject_setScale(void* self, double scale) {
	((QGraphicsObject*)self)->setScale(scale);
}
double QGraphicsObject_scale(void* self) {
	auto&& __r = ((QGraphicsObject*)self)->scale();
	return __r;
}
void** QGraphicsObject_transformations(void* self, long long* __out_len) {
	auto&& __r = ((QGraphicsObject*)self)->transformations();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
void QGraphicsObject_setTransformations(void* self, void** transformations, long long transformations_len) {
	QList<QGraphicsTransform*> transformations_list; transformations_list.reserve(transformations_len);
	for (long long __i = 0; __i < transformations_len; __i++) transformations_list.append((QGraphicsTransform*)transformations[__i]);
	((QGraphicsObject*)self)->setTransformations(transformations_list);
}
void* QGraphicsObject_transformOriginPoint(void* self) {
	auto&& __r = ((QGraphicsObject*)self)->transformOriginPoint();
	return new QPointF(__r);
}
void QGraphicsObject_setTransformOriginPoint(void* self, void* origin) {
	((QGraphicsObject*)self)->setTransformOriginPoint(*(QPointF*)origin);
}
void QGraphicsObject_setTransformOriginPoint2(void* self, double ax, double ay) {
	((QGraphicsObject*)self)->setTransformOriginPoint(ax, ay);
}
void QGraphicsObject_advance(void* self, int phase) {
	((QGraphicsObject*)self)->advance(phase);
}
double QGraphicsObject_zValue(void* self) {
	auto&& __r = ((QGraphicsObject*)self)->zValue();
	return __r;
}
void QGraphicsObject_setZValue(void* self, double z) {
	((QGraphicsObject*)self)->setZValue(z);
}
void QGraphicsObject_stackBefore(void* self, void* sibling) {
	((QGraphicsObject*)self)->stackBefore((QGraphicsItem*)sibling);
}
void* QGraphicsObject_boundingRect(void* self) {
	auto&& __r = ((QGraphicsObject*)self)->boundingRect();
	return new QRectF(__r);
}
void* QGraphicsObject_childrenBoundingRect(void* self) {
	auto&& __r = ((QGraphicsObject*)self)->childrenBoundingRect();
	return new QRectF(__r);
}
void* QGraphicsObject_sceneBoundingRect(void* self) {
	auto&& __r = ((QGraphicsObject*)self)->sceneBoundingRect();
	return new QRectF(__r);
}
void* QGraphicsObject_shape(void* self) {
	auto&& __r = ((QGraphicsObject*)self)->shape();
	return new QPainterPath(__r);
}
int QGraphicsObject_isClipped(void* self) {
	auto&& __r = ((QGraphicsObject*)self)->isClipped();
	return __r;
}
void* QGraphicsObject_clipPath(void* self) {
	auto&& __r = ((QGraphicsObject*)self)->clipPath();
	return new QPainterPath(__r);
}
int QGraphicsObject_contains(void* self, void* point) {
	auto&& __r = ((QGraphicsObject*)self)->contains(*(QPointF*)point);
	return __r;
}
int QGraphicsObject_collidesWithItem(void* self, void* other, int mode) {
	auto&& __r = ((QGraphicsObject*)self)->collidesWithItem((QGraphicsItem*)other, (Qt::ItemSelectionMode)(mode));
	return __r;
}
int QGraphicsObject_collidesWithPath(void* self, void* path, int mode) {
	auto&& __r = ((QGraphicsObject*)self)->collidesWithPath(*(QPainterPath*)path, (Qt::ItemSelectionMode)(mode));
	return __r;
}
void** QGraphicsObject_collidingItems(void* self, long long* __out_len) {
	auto&& __r = ((QGraphicsObject*)self)->collidingItems();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
int QGraphicsObject_isObscured(void* self) {
	auto&& __r = ((QGraphicsObject*)self)->isObscured();
	return __r;
}
int QGraphicsObject_isObscured2(void* self, double x, double y, double w, double h) {
	auto&& __r = ((QGraphicsObject*)self)->isObscured(x, y, w, h);
	return __r;
}
int QGraphicsObject_isObscuredBy(void* self, void* item) {
	auto&& __r = ((QGraphicsObject*)self)->isObscuredBy((QGraphicsItem*)item);
	return __r;
}
void* QGraphicsObject_opaqueArea(void* self) {
	auto&& __r = ((QGraphicsObject*)self)->opaqueArea();
	return new QPainterPath(__r);
}
void* QGraphicsObject_boundingRegion(void* self, void* itemToDeviceTransform) {
	auto&& __r = ((QGraphicsObject*)self)->boundingRegion(*(QTransform*)itemToDeviceTransform);
	return new QRegion(__r);
}
double QGraphicsObject_boundingRegionGranularity(void* self) {
	auto&& __r = ((QGraphicsObject*)self)->boundingRegionGranularity();
	return __r;
}
void QGraphicsObject_setBoundingRegionGranularity(void* self, double granularity) {
	((QGraphicsObject*)self)->setBoundingRegionGranularity(granularity);
}
void QGraphicsObject_paint(void* self, void* painter, void* option, void* widget) {
	((QGraphicsObject*)self)->paint((QPainter*)painter, (QStyleOptionGraphicsItem*)option, (QWidget*)widget);
}
void QGraphicsObject_update(void* self) {
	((QGraphicsObject*)self)->update();
}
void QGraphicsObject_update2(void* self, double x, double y, double width, double height) {
	((QGraphicsObject*)self)->update(x, y, width, height);
}
void QGraphicsObject_scroll(void* self, double dx, double dy) {
	((QGraphicsObject*)self)->scroll(dx, dy);
}
void* QGraphicsObject_mapToItem(void* self, void* item, void* point) {
	auto&& __r = ((QGraphicsObject*)self)->mapToItem((QGraphicsItem*)item, *(QPointF*)point);
	return new QPointF(__r);
}
void* QGraphicsObject_mapToParent(void* self, void* point) {
	auto&& __r = ((QGraphicsObject*)self)->mapToParent(*(QPointF*)point);
	return new QPointF(__r);
}
void* QGraphicsObject_mapToScene(void* self, void* point) {
	auto&& __r = ((QGraphicsObject*)self)->mapToScene(*(QPointF*)point);
	return new QPointF(__r);
}
void* QGraphicsObject_mapToItem2(void* self, void* item, void* rect) {
	auto&& __r = ((QGraphicsObject*)self)->mapToItem((QGraphicsItem*)item, *(QRectF*)rect);
	return new QPolygonF(__r);
}
void* QGraphicsObject_mapToParent2(void* self, void* rect) {
	auto&& __r = ((QGraphicsObject*)self)->mapToParent(*(QRectF*)rect);
	return new QPolygonF(__r);
}
void* QGraphicsObject_mapToScene2(void* self, void* rect) {
	auto&& __r = ((QGraphicsObject*)self)->mapToScene(*(QRectF*)rect);
	return new QPolygonF(__r);
}
void* QGraphicsObject_mapRectToItem(void* self, void* item, void* rect) {
	auto&& __r = ((QGraphicsObject*)self)->mapRectToItem((QGraphicsItem*)item, *(QRectF*)rect);
	return new QRectF(__r);
}
void* QGraphicsObject_mapRectToParent(void* self, void* rect) {
	auto&& __r = ((QGraphicsObject*)self)->mapRectToParent(*(QRectF*)rect);
	return new QRectF(__r);
}
void* QGraphicsObject_mapRectToScene(void* self, void* rect) {
	auto&& __r = ((QGraphicsObject*)self)->mapRectToScene(*(QRectF*)rect);
	return new QRectF(__r);
}
void* QGraphicsObject_mapToItem3(void* self, void* item, void* polygon) {
	auto&& __r = ((QGraphicsObject*)self)->mapToItem((QGraphicsItem*)item, *(QPolygonF*)polygon);
	return new QPolygonF(__r);
}
void* QGraphicsObject_mapToParent3(void* self, void* polygon) {
	auto&& __r = ((QGraphicsObject*)self)->mapToParent(*(QPolygonF*)polygon);
	return new QPolygonF(__r);
}
void* QGraphicsObject_mapToScene3(void* self, void* polygon) {
	auto&& __r = ((QGraphicsObject*)self)->mapToScene(*(QPolygonF*)polygon);
	return new QPolygonF(__r);
}
void* QGraphicsObject_mapToItem4(void* self, void* item, void* path) {
	auto&& __r = ((QGraphicsObject*)self)->mapToItem((QGraphicsItem*)item, *(QPainterPath*)path);
	return new QPainterPath(__r);
}
void* QGraphicsObject_mapToParent4(void* self, void* path) {
	auto&& __r = ((QGraphicsObject*)self)->mapToParent(*(QPainterPath*)path);
	return new QPainterPath(__r);
}
void* QGraphicsObject_mapToScene4(void* self, void* path) {
	auto&& __r = ((QGraphicsObject*)self)->mapToScene(*(QPainterPath*)path);
	return new QPainterPath(__r);
}
void* QGraphicsObject_mapFromItem(void* self, void* item, void* point) {
	auto&& __r = ((QGraphicsObject*)self)->mapFromItem((QGraphicsItem*)item, *(QPointF*)point);
	return new QPointF(__r);
}
void* QGraphicsObject_mapFromParent(void* self, void* point) {
	auto&& __r = ((QGraphicsObject*)self)->mapFromParent(*(QPointF*)point);
	return new QPointF(__r);
}
void* QGraphicsObject_mapFromScene(void* self, void* point) {
	auto&& __r = ((QGraphicsObject*)self)->mapFromScene(*(QPointF*)point);
	return new QPointF(__r);
}
void* QGraphicsObject_mapFromItem2(void* self, void* item, void* rect) {
	auto&& __r = ((QGraphicsObject*)self)->mapFromItem((QGraphicsItem*)item, *(QRectF*)rect);
	return new QPolygonF(__r);
}
void* QGraphicsObject_mapFromParent2(void* self, void* rect) {
	auto&& __r = ((QGraphicsObject*)self)->mapFromParent(*(QRectF*)rect);
	return new QPolygonF(__r);
}
void* QGraphicsObject_mapFromScene2(void* self, void* rect) {
	auto&& __r = ((QGraphicsObject*)self)->mapFromScene(*(QRectF*)rect);
	return new QPolygonF(__r);
}
void* QGraphicsObject_mapRectFromItem(void* self, void* item, void* rect) {
	auto&& __r = ((QGraphicsObject*)self)->mapRectFromItem((QGraphicsItem*)item, *(QRectF*)rect);
	return new QRectF(__r);
}
void* QGraphicsObject_mapRectFromParent(void* self, void* rect) {
	auto&& __r = ((QGraphicsObject*)self)->mapRectFromParent(*(QRectF*)rect);
	return new QRectF(__r);
}
void* QGraphicsObject_mapRectFromScene(void* self, void* rect) {
	auto&& __r = ((QGraphicsObject*)self)->mapRectFromScene(*(QRectF*)rect);
	return new QRectF(__r);
}
void* QGraphicsObject_mapFromItem3(void* self, void* item, void* polygon) {
	auto&& __r = ((QGraphicsObject*)self)->mapFromItem((QGraphicsItem*)item, *(QPolygonF*)polygon);
	return new QPolygonF(__r);
}
void* QGraphicsObject_mapFromParent3(void* self, void* polygon) {
	auto&& __r = ((QGraphicsObject*)self)->mapFromParent(*(QPolygonF*)polygon);
	return new QPolygonF(__r);
}
void* QGraphicsObject_mapFromScene3(void* self, void* polygon) {
	auto&& __r = ((QGraphicsObject*)self)->mapFromScene(*(QPolygonF*)polygon);
	return new QPolygonF(__r);
}
void* QGraphicsObject_mapFromItem4(void* self, void* item, void* path) {
	auto&& __r = ((QGraphicsObject*)self)->mapFromItem((QGraphicsItem*)item, *(QPainterPath*)path);
	return new QPainterPath(__r);
}
void* QGraphicsObject_mapFromParent4(void* self, void* path) {
	auto&& __r = ((QGraphicsObject*)self)->mapFromParent(*(QPainterPath*)path);
	return new QPainterPath(__r);
}
void* QGraphicsObject_mapFromScene4(void* self, void* path) {
	auto&& __r = ((QGraphicsObject*)self)->mapFromScene(*(QPainterPath*)path);
	return new QPainterPath(__r);
}
void* QGraphicsObject_mapToItem5(void* self, void* item, double x, double y) {
	auto&& __r = ((QGraphicsObject*)self)->mapToItem((QGraphicsItem*)item, x, y);
	return new QPointF(__r);
}
void* QGraphicsObject_mapToParent5(void* self, double x, double y) {
	auto&& __r = ((QGraphicsObject*)self)->mapToParent(x, y);
	return new QPointF(__r);
}
void* QGraphicsObject_mapToScene5(void* self, double x, double y) {
	auto&& __r = ((QGraphicsObject*)self)->mapToScene(x, y);
	return new QPointF(__r);
}
void* QGraphicsObject_mapToItem6(void* self, void* item, double x, double y, double w, double h) {
	auto&& __r = ((QGraphicsObject*)self)->mapToItem((QGraphicsItem*)item, x, y, w, h);
	return new QPolygonF(__r);
}
void* QGraphicsObject_mapToParent6(void* self, double x, double y, double w, double h) {
	auto&& __r = ((QGraphicsObject*)self)->mapToParent(x, y, w, h);
	return new QPolygonF(__r);
}
void* QGraphicsObject_mapToScene6(void* self, double x, double y, double w, double h) {
	auto&& __r = ((QGraphicsObject*)self)->mapToScene(x, y, w, h);
	return new QPolygonF(__r);
}
void* QGraphicsObject_mapRectToItem2(void* self, void* item, double x, double y, double w, double h) {
	auto&& __r = ((QGraphicsObject*)self)->mapRectToItem((QGraphicsItem*)item, x, y, w, h);
	return new QRectF(__r);
}
void* QGraphicsObject_mapRectToParent2(void* self, double x, double y, double w, double h) {
	auto&& __r = ((QGraphicsObject*)self)->mapRectToParent(x, y, w, h);
	return new QRectF(__r);
}
void* QGraphicsObject_mapRectToScene2(void* self, double x, double y, double w, double h) {
	auto&& __r = ((QGraphicsObject*)self)->mapRectToScene(x, y, w, h);
	return new QRectF(__r);
}
void* QGraphicsObject_mapFromItem5(void* self, void* item, double x, double y) {
	auto&& __r = ((QGraphicsObject*)self)->mapFromItem((QGraphicsItem*)item, x, y);
	return new QPointF(__r);
}
void* QGraphicsObject_mapFromParent5(void* self, double x, double y) {
	auto&& __r = ((QGraphicsObject*)self)->mapFromParent(x, y);
	return new QPointF(__r);
}
void* QGraphicsObject_mapFromScene5(void* self, double x, double y) {
	auto&& __r = ((QGraphicsObject*)self)->mapFromScene(x, y);
	return new QPointF(__r);
}
void* QGraphicsObject_mapFromItem6(void* self, void* item, double x, double y, double w, double h) {
	auto&& __r = ((QGraphicsObject*)self)->mapFromItem((QGraphicsItem*)item, x, y, w, h);
	return new QPolygonF(__r);
}
void* QGraphicsObject_mapFromParent6(void* self, double x, double y, double w, double h) {
	auto&& __r = ((QGraphicsObject*)self)->mapFromParent(x, y, w, h);
	return new QPolygonF(__r);
}
void* QGraphicsObject_mapFromScene6(void* self, double x, double y, double w, double h) {
	auto&& __r = ((QGraphicsObject*)self)->mapFromScene(x, y, w, h);
	return new QPolygonF(__r);
}
void* QGraphicsObject_mapRectFromItem2(void* self, void* item, double x, double y, double w, double h) {
	auto&& __r = ((QGraphicsObject*)self)->mapRectFromItem((QGraphicsItem*)item, x, y, w, h);
	return new QRectF(__r);
}
void* QGraphicsObject_mapRectFromParent2(void* self, double x, double y, double w, double h) {
	auto&& __r = ((QGraphicsObject*)self)->mapRectFromParent(x, y, w, h);
	return new QRectF(__r);
}
void* QGraphicsObject_mapRectFromScene2(void* self, double x, double y, double w, double h) {
	auto&& __r = ((QGraphicsObject*)self)->mapRectFromScene(x, y, w, h);
	return new QRectF(__r);
}
int QGraphicsObject_isAncestorOf(void* self, void* child) {
	auto&& __r = ((QGraphicsObject*)self)->isAncestorOf((QGraphicsItem*)child);
	return __r;
}
void* QGraphicsObject_commonAncestorItem(void* self, void* other) {
	auto&& __r = ((QGraphicsObject*)self)->commonAncestorItem((QGraphicsItem*)other);
	return (void*)(__r);
}
int QGraphicsObject_isUnderMouse(void* self) {
	auto&& __r = ((QGraphicsObject*)self)->isUnderMouse();
	return __r;
}
void* QGraphicsObject_data(void* self, int key) {
	auto&& __r = ((QGraphicsObject*)self)->data(key);
	return new QVariant(__r);
}
void QGraphicsObject_setData(void* self, int key, void* value) {
	((QGraphicsObject*)self)->setData(key, *(QVariant*)value);
}
int QGraphicsObject_inputMethodHints(void* self) {
	auto&& __r = ((QGraphicsObject*)self)->inputMethodHints();
	return (int)(__r);
}
void QGraphicsObject_setInputMethodHints(void* self, int hints) {
	((QGraphicsObject*)self)->setInputMethodHints(QFlags<Qt::InputMethodHint>(hints));
}
int QGraphicsObject_type(void* self) {
	auto&& __r = ((QGraphicsObject*)self)->type();
	return __r;
}
void QGraphicsObject_installSceneEventFilter(void* self, void* filterItem) {
	((QGraphicsObject*)self)->installSceneEventFilter((QGraphicsItem*)filterItem);
}
void QGraphicsObject_removeSceneEventFilter(void* self, void* filterItem) {
	((QGraphicsObject*)self)->removeSceneEventFilter((QGraphicsItem*)filterItem);
}
void QGraphicsObject_addToIndex(void* self) {
	((QGraphicsObject__Go*)self)->addToIndex();
}
void QGraphicsObject_removeFromIndex(void* self) {
	((QGraphicsObject__Go*)self)->removeFromIndex();
}
void QGraphicsObject_prepareGeometryChange(void* self) {
	((QGraphicsObject__Go*)self)->prepareGeometryChange();
}
char* QGraphicsObject_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QGraphicsObject::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QGraphicsObject_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QGraphicsObject::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QGraphicsObject_grabGesture2(void* self, int type_, int flags) {
	((QGraphicsObject*)self)->grabGesture((Qt::GestureType)(type_), QFlags<Qt::GestureFlag>(flags));
}
void* QGraphicsObject_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QGraphicsObject::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QGraphicsObject_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QGraphicsObject::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void QGraphicsObject_setFlag2(void* self, int flag, int enabled) {
	((QGraphicsObject*)self)->setFlag((QGraphicsItem::GraphicsItemFlag)(flag), enabled);
}
void QGraphicsObject_setCacheMode2(void* self, int mode, void* cacheSize) {
	((QGraphicsObject*)self)->setCacheMode((QGraphicsItem::CacheMode)(mode), *(QSize*)cacheSize);
}
void QGraphicsObject_setFocus2(void* self, int focusReason) {
	((QGraphicsObject*)self)->setFocus((Qt::FocusReason)(focusReason));
}
void QGraphicsObject_ensureVisible3(void* self, void* rect) {
	((QGraphicsObject*)self)->ensureVisible(*(QRectF*)rect);
}
void QGraphicsObject_ensureVisible4(void* self, void* rect, int xmargin) {
	((QGraphicsObject*)self)->ensureVisible(*(QRectF*)rect, xmargin);
}
void QGraphicsObject_ensureVisible5(void* self, void* rect, int xmargin, int ymargin) {
	((QGraphicsObject*)self)->ensureVisible(*(QRectF*)rect, xmargin, ymargin);
}
void QGraphicsObject_ensureVisible6(void* self, double x, double y, double w, double h, int xmargin) {
	((QGraphicsObject*)self)->ensureVisible(x, y, w, h, xmargin);
}
void QGraphicsObject_ensureVisible7(void* self, double x, double y, double w, double h, int xmargin, int ymargin) {
	((QGraphicsObject*)self)->ensureVisible(x, y, w, h, xmargin, ymargin);
}
void* QGraphicsObject_itemTransform2(void* self, void* other, int* ok) {
	bool ok_local{};
	auto&& __r = ((QGraphicsObject*)self)->itemTransform((QGraphicsItem*)other, &ok_local);
	*ok = ok_local;
	return new QTransform(__r);
}
void QGraphicsObject_setTransform2(void* self, void* matrix, int combine) {
	((QGraphicsObject*)self)->setTransform(*(QTransform*)matrix, combine);
}
void** QGraphicsObject_collidingItems2(void* self, int mode, long long* __out_len) {
	auto&& __r = ((QGraphicsObject*)self)->collidingItems((Qt::ItemSelectionMode)(mode));
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
int QGraphicsObject_isObscured3(void* self, void* rect) {
	auto&& __r = ((QGraphicsObject*)self)->isObscured(*(QRectF*)rect);
	return __r;
}
void QGraphicsObject_update3(void* self, void* rect) {
	((QGraphicsObject*)self)->update(*(QRectF*)rect);
}
void QGraphicsObject_scroll2(void* self, double dx, double dy, void* rect) {
	((QGraphicsObject*)self)->scroll(dx, dy, *(QRectF*)rect);
}
void* QGraphicsObject_asQGraphicsItem(void* self) {
	auto&& __r = static_cast<QGraphicsItem*>((QGraphicsObject*)self);
	return (void*)(__r);
}
void QGraphicsObject_advance_default(void* self, int phase) {
	((QGraphicsObject*)self)->QGraphicsObject::advance(phase);
}
void* QGraphicsObject_shape_default(void* self) {
	auto&& __r = ((QGraphicsObject*)self)->QGraphicsObject::shape();
	return new QPainterPath(__r);
}
int QGraphicsObject_contains_default(void* self, void* point) {
	auto&& __r = ((QGraphicsObject*)self)->QGraphicsObject::contains(*(QPointF*)point);
	return __r;
}
int QGraphicsObject_collidesWithItem_default(void* self, void* other, int mode) {
	auto&& __r = ((QGraphicsObject*)self)->QGraphicsObject::collidesWithItem((QGraphicsItem*)other, (Qt::ItemSelectionMode)(mode));
	return __r;
}
int QGraphicsObject_collidesWithPath_default(void* self, void* path, int mode) {
	auto&& __r = ((QGraphicsObject*)self)->QGraphicsObject::collidesWithPath(*(QPainterPath*)path, (Qt::ItemSelectionMode)(mode));
	return __r;
}
int QGraphicsObject_isObscuredBy_default(void* self, void* item) {
	auto&& __r = ((QGraphicsObject*)self)->QGraphicsObject::isObscuredBy((QGraphicsItem*)item);
	return __r;
}
void* QGraphicsObject_opaqueArea_default(void* self) {
	auto&& __r = ((QGraphicsObject*)self)->QGraphicsObject::opaqueArea();
	return new QPainterPath(__r);
}
int QGraphicsObject_type_default(void* self) {
	auto&& __r = ((QGraphicsObject*)self)->QGraphicsObject::type();
	return __r;
}
void* QGraphicsObject_metaObject_default(void* self) {
	auto&& __r = ((QGraphicsObject*)self)->QGraphicsObject::metaObject();
	return (void*)(__r);
}
int QGraphicsObject_event_default(void* self, void* ev) {
	auto&& __r = ((QGraphicsObject__Go*)self)->cppgoBase_event((QEvent*)ev);
	return __r;
}
int QGraphicsObject_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QGraphicsObject*)self)->QGraphicsObject::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QGraphicsObject_timerEvent_default(void* self, void* event) {
	((QGraphicsObject__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QGraphicsObject_childEvent_default(void* self, void* event) {
	((QGraphicsObject__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QGraphicsObject_customEvent_default(void* self, void* event) {
	((QGraphicsObject__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QGraphicsObject_connectNotify_default(void* self, void* signal) {
	((QGraphicsObject__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QGraphicsObject_disconnectNotify_default(void* self, void* signal) {
	((QGraphicsObject__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
int QGraphicsObject_sceneEventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QGraphicsObject__Go*)self)->cppgoBase_sceneEventFilter((QGraphicsItem*)watched, (QEvent*)event);
	return __r;
}
int QGraphicsObject_sceneEvent_default(void* self, void* event) {
	auto&& __r = ((QGraphicsObject__Go*)self)->cppgoBase_sceneEvent((QEvent*)event);
	return __r;
}
void QGraphicsObject_contextMenuEvent_default(void* self, void* event) {
	((QGraphicsObject__Go*)self)->cppgoBase_contextMenuEvent((QGraphicsSceneContextMenuEvent*)event);
}
void QGraphicsObject_dragEnterEvent_default(void* self, void* event) {
	((QGraphicsObject__Go*)self)->cppgoBase_dragEnterEvent((QGraphicsSceneDragDropEvent*)event);
}
void QGraphicsObject_dragLeaveEvent_default(void* self, void* event) {
	((QGraphicsObject__Go*)self)->cppgoBase_dragLeaveEvent((QGraphicsSceneDragDropEvent*)event);
}
void QGraphicsObject_dragMoveEvent_default(void* self, void* event) {
	((QGraphicsObject__Go*)self)->cppgoBase_dragMoveEvent((QGraphicsSceneDragDropEvent*)event);
}
void QGraphicsObject_dropEvent_default(void* self, void* event) {
	((QGraphicsObject__Go*)self)->cppgoBase_dropEvent((QGraphicsSceneDragDropEvent*)event);
}
void QGraphicsObject_focusInEvent_default(void* self, void* event) {
	((QGraphicsObject__Go*)self)->cppgoBase_focusInEvent((QFocusEvent*)event);
}
void QGraphicsObject_focusOutEvent_default(void* self, void* event) {
	((QGraphicsObject__Go*)self)->cppgoBase_focusOutEvent((QFocusEvent*)event);
}
void QGraphicsObject_hoverEnterEvent_default(void* self, void* event) {
	((QGraphicsObject__Go*)self)->cppgoBase_hoverEnterEvent((QGraphicsSceneHoverEvent*)event);
}
void QGraphicsObject_hoverMoveEvent_default(void* self, void* event) {
	((QGraphicsObject__Go*)self)->cppgoBase_hoverMoveEvent((QGraphicsSceneHoverEvent*)event);
}
void QGraphicsObject_hoverLeaveEvent_default(void* self, void* event) {
	((QGraphicsObject__Go*)self)->cppgoBase_hoverLeaveEvent((QGraphicsSceneHoverEvent*)event);
}
void QGraphicsObject_keyPressEvent_default(void* self, void* event) {
	((QGraphicsObject__Go*)self)->cppgoBase_keyPressEvent((QKeyEvent*)event);
}
void QGraphicsObject_keyReleaseEvent_default(void* self, void* event) {
	((QGraphicsObject__Go*)self)->cppgoBase_keyReleaseEvent((QKeyEvent*)event);
}
void QGraphicsObject_mousePressEvent_default(void* self, void* event) {
	((QGraphicsObject__Go*)self)->cppgoBase_mousePressEvent((QGraphicsSceneMouseEvent*)event);
}
void QGraphicsObject_mouseMoveEvent_default(void* self, void* event) {
	((QGraphicsObject__Go*)self)->cppgoBase_mouseMoveEvent((QGraphicsSceneMouseEvent*)event);
}
void QGraphicsObject_mouseReleaseEvent_default(void* self, void* event) {
	((QGraphicsObject__Go*)self)->cppgoBase_mouseReleaseEvent((QGraphicsSceneMouseEvent*)event);
}
void QGraphicsObject_mouseDoubleClickEvent_default(void* self, void* event) {
	((QGraphicsObject__Go*)self)->cppgoBase_mouseDoubleClickEvent((QGraphicsSceneMouseEvent*)event);
}
void QGraphicsObject_wheelEvent_default(void* self, void* event) {
	((QGraphicsObject__Go*)self)->cppgoBase_wheelEvent((QGraphicsSceneWheelEvent*)event);
}
void QGraphicsObject_inputMethodEvent_default(void* self, void* event) {
	((QGraphicsObject__Go*)self)->cppgoBase_inputMethodEvent((QInputMethodEvent*)event);
}
void* QGraphicsObject_inputMethodQuery_default(void* self, int query) {
	auto&& __r = ((QGraphicsObject__Go*)self)->cppgoBase_inputMethodQuery(query);
	return new QVariant(__r);
}
void* QGraphicsObject_itemChange_default(void* self, int change, void* value) {
	auto&& __r = ((QGraphicsObject__Go*)self)->cppgoBase_itemChange(change, *(QVariant*)value);
	return new QVariant(__r);
}
int QGraphicsObject_supportsExtension_default(void* self, int extension) {
	auto&& __r = ((QGraphicsObject__Go*)self)->cppgoBase_supportsExtension(extension);
	return __r;
}
void QGraphicsObject_setExtension_default(void* self, int extension, void* variant) {
	((QGraphicsObject__Go*)self)->cppgoBase_setExtension(extension, *(QVariant*)variant);
}
void* QGraphicsObject_extension_default(void* self, void* variant) {
	auto&& __r = ((QGraphicsObject__Go*)self)->cppgoBase_extension(*(QVariant*)variant);
	return new QVariant(__r);
}
void QGraphicsObject_connect_parentChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsObject*)self, &QGraphicsObject::parentChanged, (QGraphicsObject*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsObject);
	__conns_QGraphicsObject[__id] = __c;
}
void QGraphicsObject_connect_opacityChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsObject*)self, &QGraphicsObject::opacityChanged, (QGraphicsObject*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsObject);
	__conns_QGraphicsObject[__id] = __c;
}
void QGraphicsObject_connect_visibleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsObject*)self, &QGraphicsObject::visibleChanged, (QGraphicsObject*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsObject);
	__conns_QGraphicsObject[__id] = __c;
}
void QGraphicsObject_connect_enabledChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsObject*)self, &QGraphicsObject::enabledChanged, (QGraphicsObject*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsObject);
	__conns_QGraphicsObject[__id] = __c;
}
void QGraphicsObject_connect_xChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsObject*)self, &QGraphicsObject::xChanged, (QGraphicsObject*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsObject);
	__conns_QGraphicsObject[__id] = __c;
}
void QGraphicsObject_connect_yChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsObject*)self, &QGraphicsObject::yChanged, (QGraphicsObject*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsObject);
	__conns_QGraphicsObject[__id] = __c;
}
void QGraphicsObject_connect_zChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsObject*)self, &QGraphicsObject::zChanged, (QGraphicsObject*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsObject);
	__conns_QGraphicsObject[__id] = __c;
}
void QGraphicsObject_connect_rotationChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsObject*)self, &QGraphicsObject::rotationChanged, (QGraphicsObject*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsObject);
	__conns_QGraphicsObject[__id] = __c;
}
void QGraphicsObject_connect_scaleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsObject*)self, &QGraphicsObject::scaleChanged, (QGraphicsObject*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsObject);
	__conns_QGraphicsObject[__id] = __c;
}
void QGraphicsObject_connect_childrenChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsObject*)self, &QGraphicsObject::childrenChanged, (QGraphicsObject*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsObject);
	__conns_QGraphicsObject[__id] = __c;
}
void QGraphicsObject_connect_widthChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsObject*)self, &QGraphicsObject::widthChanged, (QGraphicsObject*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsObject);
	__conns_QGraphicsObject[__id] = __c;
}
void QGraphicsObject_connect_heightChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsObject*)self, &QGraphicsObject::heightChanged, (QGraphicsObject*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsObject);
	__conns_QGraphicsObject[__id] = __c;
}
void QGraphicsObject_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsObject*)self, &QGraphicsObject::destroyed, (QGraphicsObject*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsObject);
	__conns_QGraphicsObject[__id] = __c;
}
void QGraphicsObject_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsObject*)self, &QGraphicsObject::destroyed, (QGraphicsObject*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsObject);
	__conns_QGraphicsObject[__id] = __c;
}
void QGraphicsObject_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsObject*)self, &QGraphicsObject::objectNameChanged, (QGraphicsObject*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsObject);
	__conns_QGraphicsObject[__id] = __c;
}
void QGraphicsObject__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsObject);
		auto __it = __conns_QGraphicsObject.find(id);
		if (__it == __conns_QGraphicsObject.end()) return;
		__c = __it->second;
		__conns_QGraphicsObject.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QGraphicsObject__Go_new(long long id) { return new QGraphicsObject__Go(id); }
void QGraphicsObject_delete(void* self) { delete (QGraphicsObject*)self; }

}

#include <QtWidgets/qgraphicsproxywidget.h>
#include <QtWidgets/qgraphicswidget.h>
#include <QtWidgets/qgraphicsitem.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qrect.h>
#include <QtGui/qpainter.h>
#include <QtWidgets/qstyleoption.h>
#include <QtCore/qvariant.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qobject.h>
#include <QtGui/qevent.h>
#include <QtWidgets/qgraphicssceneevent.h>
#include <QtCore/qsize.h>
#include <QtWidgets/qgraphicslayout.h>
#include <QtWidgets/qstyle.h>
#include <QtGui/qfont.h>
#include <QtGui/qpalette.h>
#include <QtCore/qmargins.h>
#include <QtGui/qkeysequence.h>
#include <QtGui/qaction.h>
#include <QtGui/qpainterpath.h>
#include <QtCore/qpoint.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qbindingstorage.h>
#include <QtWidgets/qgraphicsscene.h>
#include <QtGui/qcursor.h>
#include <QtWidgets/qgraphicseffect.h>
#include <QtGui/qtransform.h>
#include <QtWidgets/qgraphicstransform.h>
#include <QtGui/qregion.h>
#include <QtGui/qpolygon.h>
#include <QtWidgets/qsizepolicy.h>
#include <QtWidgets/qgraphicslayoutitem.h>
#include "gen_qgraphicsproxywidget.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QGraphicsProxyWidget__Go : public QGraphicsProxyWidget {
public:
	long long __goId;
	unsigned char __goOn[65] = {};
	template<typename... A> QGraphicsProxyWidget__Go(long long id, A&&... a) : QGraphicsProxyWidget(std::forward<A>(a)...), __goId(id) {}
	using QGraphicsProxyWidget::newProxyWidget;
	using QGraphicsProxyWidget::updateMicroFocus;
	using QGraphicsProxyWidget::sender;
	using QGraphicsProxyWidget::senderSignalIndex;
	using QGraphicsProxyWidget::receivers;
	using QGraphicsProxyWidget::isSignalConnected;
	using QGraphicsProxyWidget::addToIndex;
	using QGraphicsProxyWidget::removeFromIndex;
	using QGraphicsProxyWidget::prepareGeometryChange;
	using QGraphicsProxyWidget::setGraphicsItem;
	using QGraphicsProxyWidget::setOwnedByLayout;
	QVariant cppgoBase_itemChange(int __a0, const QVariant& __a1) { return QGraphicsProxyWidget::itemChange((QGraphicsItem::GraphicsItemChange)__a0, __a1); }
	bool cppgoBase_event(QEvent* __a0) { return QGraphicsProxyWidget::event(__a0); }
	bool cppgoBase_eventFilter(QObject* __a0, QEvent* __a1) { return QGraphicsProxyWidget::eventFilter(__a0, __a1); }
	void cppgoBase_showEvent(QShowEvent* __a0) { QGraphicsProxyWidget::showEvent(__a0); }
	void cppgoBase_hideEvent(QHideEvent* __a0) { QGraphicsProxyWidget::hideEvent(__a0); }
	void cppgoBase_contextMenuEvent(QGraphicsSceneContextMenuEvent* __a0) { QGraphicsProxyWidget::contextMenuEvent(__a0); }
	void cppgoBase_dragEnterEvent(QGraphicsSceneDragDropEvent* __a0) { QGraphicsProxyWidget::dragEnterEvent(__a0); }
	void cppgoBase_dragLeaveEvent(QGraphicsSceneDragDropEvent* __a0) { QGraphicsProxyWidget::dragLeaveEvent(__a0); }
	void cppgoBase_dragMoveEvent(QGraphicsSceneDragDropEvent* __a0) { QGraphicsProxyWidget::dragMoveEvent(__a0); }
	void cppgoBase_dropEvent(QGraphicsSceneDragDropEvent* __a0) { QGraphicsProxyWidget::dropEvent(__a0); }
	void cppgoBase_hoverEnterEvent(QGraphicsSceneHoverEvent* __a0) { QGraphicsProxyWidget::hoverEnterEvent(__a0); }
	void cppgoBase_hoverLeaveEvent(QGraphicsSceneHoverEvent* __a0) { QGraphicsProxyWidget::hoverLeaveEvent(__a0); }
	void cppgoBase_hoverMoveEvent(QGraphicsSceneHoverEvent* __a0) { QGraphicsProxyWidget::hoverMoveEvent(__a0); }
	void cppgoBase_grabMouseEvent(QEvent* __a0) { QGraphicsProxyWidget::grabMouseEvent(__a0); }
	void cppgoBase_ungrabMouseEvent(QEvent* __a0) { QGraphicsProxyWidget::ungrabMouseEvent(__a0); }
	void cppgoBase_mouseMoveEvent(QGraphicsSceneMouseEvent* __a0) { QGraphicsProxyWidget::mouseMoveEvent(__a0); }
	void cppgoBase_mousePressEvent(QGraphicsSceneMouseEvent* __a0) { QGraphicsProxyWidget::mousePressEvent(__a0); }
	void cppgoBase_mouseReleaseEvent(QGraphicsSceneMouseEvent* __a0) { QGraphicsProxyWidget::mouseReleaseEvent(__a0); }
	void cppgoBase_mouseDoubleClickEvent(QGraphicsSceneMouseEvent* __a0) { QGraphicsProxyWidget::mouseDoubleClickEvent(__a0); }
	void cppgoBase_wheelEvent(QGraphicsSceneWheelEvent* __a0) { QGraphicsProxyWidget::wheelEvent(__a0); }
	void cppgoBase_keyPressEvent(QKeyEvent* __a0) { QGraphicsProxyWidget::keyPressEvent(__a0); }
	void cppgoBase_keyReleaseEvent(QKeyEvent* __a0) { QGraphicsProxyWidget::keyReleaseEvent(__a0); }
	void cppgoBase_focusInEvent(QFocusEvent* __a0) { QGraphicsProxyWidget::focusInEvent(__a0); }
	void cppgoBase_focusOutEvent(QFocusEvent* __a0) { QGraphicsProxyWidget::focusOutEvent(__a0); }
	bool cppgoBase_focusNextPrevChild(bool __a0) { return QGraphicsProxyWidget::focusNextPrevChild(__a0); }
	QVariant cppgoBase_inputMethodQuery(int __a0) { return QGraphicsProxyWidget::inputMethodQuery((Qt::InputMethodQuery)__a0); }
	void cppgoBase_inputMethodEvent(QInputMethodEvent* __a0) { QGraphicsProxyWidget::inputMethodEvent(__a0); }
	QSizeF cppgoBase_sizeHint(int __a0, const QSizeF& __a1) { return QGraphicsProxyWidget::sizeHint((Qt::SizeHint)__a0, __a1); }
	void cppgoBase_resizeEvent(QGraphicsSceneResizeEvent* __a0) { QGraphicsProxyWidget::resizeEvent(__a0); }
	void cppgoBase_initStyleOption(QStyleOption* __a0) { QGraphicsProxyWidget::initStyleOption(__a0); }
	void cppgoBase_updateGeometry() { QGraphicsProxyWidget::updateGeometry(); }
	QVariant cppgoBase_propertyChange(const QString& __a0, const QVariant& __a1) { return QGraphicsProxyWidget::propertyChange(__a0, __a1); }
	bool cppgoBase_sceneEvent(QEvent* __a0) { return QGraphicsProxyWidget::sceneEvent(__a0); }
	bool cppgoBase_windowFrameEvent(QEvent* __a0) { return QGraphicsProxyWidget::windowFrameEvent(__a0); }
	Qt::WindowFrameSection cppgoBase_windowFrameSectionAt(const QPointF& __a0) { return QGraphicsProxyWidget::windowFrameSectionAt(__a0); }
	void cppgoBase_changeEvent(QEvent* __a0) { QGraphicsProxyWidget::changeEvent(__a0); }
	void cppgoBase_closeEvent(QCloseEvent* __a0) { QGraphicsProxyWidget::closeEvent(__a0); }
	void cppgoBase_moveEvent(QGraphicsSceneMoveEvent* __a0) { QGraphicsProxyWidget::moveEvent(__a0); }
	void cppgoBase_polishEvent() { QGraphicsProxyWidget::polishEvent(); }
	void cppgoBase_grabKeyboardEvent(QEvent* __a0) { QGraphicsProxyWidget::grabKeyboardEvent(__a0); }
	void cppgoBase_ungrabKeyboardEvent(QEvent* __a0) { QGraphicsProxyWidget::ungrabKeyboardEvent(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QGraphicsProxyWidget::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QGraphicsProxyWidget::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QGraphicsProxyWidget::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QGraphicsProxyWidget::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QGraphicsProxyWidget::disconnectNotify(__a0); }
	bool cppgoBase_sceneEventFilter(QGraphicsItem* __a0, QEvent* __a1) { return QGraphicsProxyWidget::sceneEventFilter(__a0, __a1); }
	bool cppgoBase_supportsExtension(int __a0) { return QGraphicsProxyWidget::supportsExtension((QGraphicsItem::Extension)__a0); }
	void cppgoBase_setExtension(int __a0, const QVariant& __a1) { QGraphicsProxyWidget::setExtension((QGraphicsItem::Extension)__a0, __a1); }
	QVariant cppgoBase_extension(const QVariant& __a0) { return QGraphicsProxyWidget::extension(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QGraphicsProxyWidget::metaObject();
	}
	void setGeometry(const QRectF& rect) override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5684820848726969227LL, (void*)&rect, &__h); if (__h) return; }
		QGraphicsProxyWidget::setGeometry(rect);
	}
	void paint(QPainter* painter, const QStyleOptionGraphicsItem* option, QWidget* widget) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_void_voidp_voidp_voidp) { ((void(*)(long long, long long, void*, void*, void*, int*))cppgo_virt_void_voidp_voidp_voidp)(__goId, 6083584261053556819LL, (void*)painter, (void*)option, (void*)widget, &__h); if (__h) return; }
		QGraphicsProxyWidget::paint(painter, option, widget);
	}
	int type() const override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -6371403872082420691LL, &__h); if (__h) return __r; }
		return QGraphicsProxyWidget::type();
	}
	QVariant itemChange(QGraphicsItem::GraphicsItemChange change, const QVariant& value) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_voidp_int_voidp) { void* __r = ((void*(*)(long long, long long, int, void*, int*))cppgo_virt_voidp_int_voidp)(__goId, -1888572168794878830LL, (int)(change), (void*)&value, &__h); if (__h) return *(QVariant*)__r; }
		return QGraphicsProxyWidget::itemChange(change, value);
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QGraphicsProxyWidget::event(event);
	}
	bool eventFilter(QObject* object, QEvent* event) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)object, (void*)event, &__h); if (__h) return __r; }
		return QGraphicsProxyWidget::eventFilter(object, event);
	}
	void showEvent(QShowEvent* event) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6825369174751916LL, (void*)event, &__h); if (__h) return; }
		QGraphicsProxyWidget::showEvent(event);
	}
	void hideEvent(QHideEvent* event) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -199718707876719129LL, (void*)event, &__h); if (__h) return; }
		QGraphicsProxyWidget::hideEvent(event);
	}
	void contextMenuEvent(QGraphicsSceneContextMenuEvent* event) override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6698043222227006015LL, (void*)event, &__h); if (__h) return; }
		QGraphicsProxyWidget::contextMenuEvent(event);
	}
	void dragEnterEvent(QGraphicsSceneDragDropEvent* event) override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 206122715919917589LL, (void*)event, &__h); if (__h) return; }
		QGraphicsProxyWidget::dragEnterEvent(event);
	}
	void dragLeaveEvent(QGraphicsSceneDragDropEvent* event) override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -3139904158373725006LL, (void*)event, &__h); if (__h) return; }
		QGraphicsProxyWidget::dragLeaveEvent(event);
	}
	void dragMoveEvent(QGraphicsSceneDragDropEvent* event) override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1781545137352745978LL, (void*)event, &__h); if (__h) return; }
		QGraphicsProxyWidget::dragMoveEvent(event);
	}
	void dropEvent(QGraphicsSceneDragDropEvent* event) override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 9107058682754591716LL, (void*)event, &__h); if (__h) return; }
		QGraphicsProxyWidget::dropEvent(event);
	}
	void hoverEnterEvent(QGraphicsSceneHoverEvent* event) override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6932941943906681679LL, (void*)event, &__h); if (__h) return; }
		QGraphicsProxyWidget::hoverEnterEvent(event);
	}
	void hoverLeaveEvent(QGraphicsSceneHoverEvent* event) override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7944408482203603984LL, (void*)event, &__h); if (__h) return; }
		QGraphicsProxyWidget::hoverLeaveEvent(event);
	}
	void hoverMoveEvent(QGraphicsSceneHoverEvent* event) override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -2745191522144316396LL, (void*)event, &__h); if (__h) return; }
		QGraphicsProxyWidget::hoverMoveEvent(event);
	}
	void grabMouseEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7928460550911621078LL, (void*)event, &__h); if (__h) return; }
		QGraphicsProxyWidget::grabMouseEvent(event);
	}
	void ungrabMouseEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5662911227992581871LL, (void*)event, &__h); if (__h) return; }
		QGraphicsProxyWidget::ungrabMouseEvent(event);
	}
	void mouseMoveEvent(QGraphicsSceneMouseEvent* event) override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4255494004099806581LL, (void*)event, &__h); if (__h) return; }
		QGraphicsProxyWidget::mouseMoveEvent(event);
	}
	void mousePressEvent(QGraphicsSceneMouseEvent* event) override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -332090898078543319LL, (void*)event, &__h); if (__h) return; }
		QGraphicsProxyWidget::mousePressEvent(event);
	}
	void mouseReleaseEvent(QGraphicsSceneMouseEvent* event) override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4415920049297795213LL, (void*)event, &__h); if (__h) return; }
		QGraphicsProxyWidget::mouseReleaseEvent(event);
	}
	void mouseDoubleClickEvent(QGraphicsSceneMouseEvent* event) override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4029073905697023887LL, (void*)event, &__h); if (__h) return; }
		QGraphicsProxyWidget::mouseDoubleClickEvent(event);
	}
	void wheelEvent(QGraphicsSceneWheelEvent* event) override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1919988340934511316LL, (void*)event, &__h); if (__h) return; }
		QGraphicsProxyWidget::wheelEvent(event);
	}
	void keyPressEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1023791977679793185LL, (void*)event, &__h); if (__h) return; }
		QGraphicsProxyWidget::keyPressEvent(event);
	}
	void keyReleaseEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 772598475950050299LL, (void*)event, &__h); if (__h) return; }
		QGraphicsProxyWidget::keyReleaseEvent(event);
	}
	void focusInEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 975043767033552552LL, (void*)event, &__h); if (__h) return; }
		QGraphicsProxyWidget::focusInEvent(event);
	}
	void focusOutEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -385596010576048993LL, (void*)event, &__h); if (__h) return; }
		QGraphicsProxyWidget::focusOutEvent(event);
	}
	bool focusNextPrevChild(bool next) override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 5898328517505784191LL, next, &__h); if (__h) return __r; }
		return QGraphicsProxyWidget::focusNextPrevChild(next);
	}
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7345853097340709638LL, (int)(query), &__h); if (__h) return *(QVariant*)__r; }
		return QGraphicsProxyWidget::inputMethodQuery(query);
	}
	void inputMethodEvent(QInputMethodEvent* event) override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5438819661942843222LL, (void*)event, &__h); if (__h) return; }
		QGraphicsProxyWidget::inputMethodEvent(event);
	}
	QSizeF sizeHint(Qt::SizeHint which, const QSizeF& constraint) const override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_voidp_int_voidp) { void* __r = ((void*(*)(long long, long long, int, void*, int*))cppgo_virt_voidp_int_voidp)(__goId, 729690261457684101LL, (int)(which), (void*)&constraint, &__h); if (__h) return *(QSizeF*)__r; }
		return QGraphicsProxyWidget::sizeHint(which, constraint);
	}
	void resizeEvent(QGraphicsSceneResizeEvent* event) override {
		int __h = 0;
		if (__goOn[32] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4302917285655014661LL, (void*)event, &__h); if (__h) return; }
		QGraphicsProxyWidget::resizeEvent(event);
	}
	void getContentsMargins(qreal* left, qreal* top, qreal* right, qreal* bottom) const override {
		int __h = 0;
		if (__goOn[33] && cppgo_virt_void_doublep_doublep_doublep_doublep) { ((void(*)(long long, long long, double*, double*, double*, double*, int*))cppgo_virt_void_doublep_doublep_doublep_doublep)(__goId, 164760286202646972LL, (double*)left, (double*)top, (double*)right, (double*)bottom, &__h); if (__h) return; }
		QGraphicsProxyWidget::getContentsMargins(left, top, right, bottom);
	}
	void paintWindowFrame(QPainter* painter, const QStyleOptionGraphicsItem* option, QWidget* widget) override {
		int __h = 0;
		if (__goOn[34] && cppgo_virt_void_voidp_voidp_voidp) { ((void(*)(long long, long long, void*, void*, void*, int*))cppgo_virt_void_voidp_voidp_voidp)(__goId, -1640122132753285616LL, (void*)painter, (void*)option, (void*)widget, &__h); if (__h) return; }
		QGraphicsProxyWidget::paintWindowFrame(painter, option, widget);
	}
	QRectF boundingRect() const override {
		int __h = 0;
		if (__goOn[35] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 2510008782785038831LL, &__h); if (__h) return *(QRectF*)__r; }
		return QGraphicsProxyWidget::boundingRect();
	}
	QPainterPath shape() const override {
		int __h = 0;
		if (__goOn[36] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 7938019069247444294LL, &__h); if (__h) return *(QPainterPath*)__r; }
		return QGraphicsProxyWidget::shape();
	}
	void initStyleOption(QStyleOption* option) const override {
		int __h = 0;
		if (__goOn[37] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -488693769771308531LL, (void*)option, &__h); if (__h) return; }
		QGraphicsProxyWidget::initStyleOption(option);
	}
	void updateGeometry() override {
		int __h = 0;
		if (__goOn[38] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 8082455443689927668LL, &__h); if (__h) return; }
		QGraphicsProxyWidget::updateGeometry();
	}
	QVariant propertyChange(const QString& propertyName, const QVariant& value) override {
		int __h = 0;
		QByteArray propertyName_ba = propertyName.toUtf8();
		if (__goOn[39] && cppgo_virt_voidp_constcharp_longlong_voidp) { void* __r = ((void*(*)(long long, long long, const char*, long long, void*, int*))cppgo_virt_voidp_constcharp_longlong_voidp)(__goId, 5289089528078387296LL, (char*)propertyName_ba.constData(), propertyName_ba.size(), (void*)&value, &__h); if (__h) return *(QVariant*)__r; }
		return QGraphicsProxyWidget::propertyChange(propertyName, value);
	}
	bool sceneEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[40] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, -2846371507838587071LL, (void*)event, &__h); if (__h) return __r; }
		return QGraphicsProxyWidget::sceneEvent(event);
	}
	bool windowFrameEvent(QEvent* e) override {
		int __h = 0;
		if (__goOn[41] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, -7362375632910814166LL, (void*)e, &__h); if (__h) return __r; }
		return QGraphicsProxyWidget::windowFrameEvent(e);
	}
	Qt::WindowFrameSection windowFrameSectionAt(const QPointF& pos) const override {
		int __h = 0;
		if (__goOn[42] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 8834866722574844732LL, (void*)&pos, &__h); if (__h) return (Qt::WindowFrameSection)(__r); }
		return QGraphicsProxyWidget::windowFrameSectionAt(pos);
	}
	void changeEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[43] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4959376063826165943LL, (void*)event, &__h); if (__h) return; }
		QGraphicsProxyWidget::changeEvent(event);
	}
	void closeEvent(QCloseEvent* event) override {
		int __h = 0;
		if (__goOn[44] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1989937050019464783LL, (void*)event, &__h); if (__h) return; }
		QGraphicsProxyWidget::closeEvent(event);
	}
	void moveEvent(QGraphicsSceneMoveEvent* event) override {
		int __h = 0;
		if (__goOn[45] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -9161949150489525396LL, (void*)event, &__h); if (__h) return; }
		QGraphicsProxyWidget::moveEvent(event);
	}
	void polishEvent() override {
		int __h = 0;
		if (__goOn[46] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, -214704830098031580LL, &__h); if (__h) return; }
		QGraphicsProxyWidget::polishEvent();
	}
	void grabKeyboardEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[47] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6423297148758141226LL, (void*)event, &__h); if (__h) return; }
		QGraphicsProxyWidget::grabKeyboardEvent(event);
	}
	void ungrabKeyboardEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[48] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 2912496686457383679LL, (void*)event, &__h); if (__h) return; }
		QGraphicsProxyWidget::ungrabKeyboardEvent(event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[49] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QGraphicsProxyWidget::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[50] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QGraphicsProxyWidget::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[51] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QGraphicsProxyWidget::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[52] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QGraphicsProxyWidget::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[53] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QGraphicsProxyWidget::disconnectNotify(signal);
	}
	void advance(int phase) override {
		int __h = 0;
		if (__goOn[54] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, -4830836341480073475LL, phase, &__h); if (__h) return; }
		QGraphicsProxyWidget::advance(phase);
	}
	bool contains(const QPointF& point) const override {
		int __h = 0;
		if (__goOn[55] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 5645366511944179768LL, (void*)&point, &__h); if (__h) return __r; }
		return QGraphicsProxyWidget::contains(point);
	}
	bool collidesWithItem(const QGraphicsItem* other, Qt::ItemSelectionMode mode) const override {
		int __h = 0;
		if (__goOn[56] && cppgo_virt_int_voidp_int) { int __r = ((int(*)(long long, long long, void*, int, int*))cppgo_virt_int_voidp_int)(__goId, 6469736184362593173LL, (void*)other, (int)(mode), &__h); if (__h) return __r; }
		return QGraphicsProxyWidget::collidesWithItem(other, mode);
	}
	bool collidesWithPath(const QPainterPath& path, Qt::ItemSelectionMode mode) const override {
		int __h = 0;
		if (__goOn[57] && cppgo_virt_int_voidp_int) { int __r = ((int(*)(long long, long long, void*, int, int*))cppgo_virt_int_voidp_int)(__goId, 6636575182718753853LL, (void*)&path, (int)(mode), &__h); if (__h) return __r; }
		return QGraphicsProxyWidget::collidesWithPath(path, mode);
	}
	bool isObscuredBy(const QGraphicsItem* item) const override {
		int __h = 0;
		if (__goOn[58] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, -7215030271914903671LL, (void*)item, &__h); if (__h) return __r; }
		return QGraphicsProxyWidget::isObscuredBy(item);
	}
	QPainterPath opaqueArea() const override {
		int __h = 0;
		if (__goOn[59] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 399370970023989435LL, &__h); if (__h) return *(QPainterPath*)__r; }
		return QGraphicsProxyWidget::opaqueArea();
	}
	bool sceneEventFilter(QGraphicsItem* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[60] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, -4584203218241215725LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QGraphicsProxyWidget::sceneEventFilter(watched, event);
	}
	bool supportsExtension(QGraphicsItem::Extension extension) const override {
		int __h = 0;
		if (__goOn[61] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, -5729166443651306418LL, (int)(extension), &__h); if (__h) return __r; }
		return QGraphicsProxyWidget::supportsExtension(extension);
	}
	void setExtension(QGraphicsItem::Extension extension, const QVariant& variant) override {
		int __h = 0;
		if (__goOn[62] && cppgo_virt_void_int_voidp) { ((void(*)(long long, long long, int, void*, int*))cppgo_virt_void_int_voidp)(__goId, -6730181576106171580LL, (int)(extension), (void*)&variant, &__h); if (__h) return; }
		QGraphicsProxyWidget::setExtension(extension, variant);
	}
	QVariant extension(const QVariant& variant) const override {
		int __h = 0;
		if (__goOn[63] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 2270616358835251226LL, (void*)&variant, &__h); if (__h) return *(QVariant*)__r; }
		return QGraphicsProxyWidget::extension(variant);
	}
	bool isEmpty() const override {
		int __h = 0;
		if (__goOn[64] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -626822430251098498LL, &__h); if (__h) return __r; }
		return QGraphicsProxyWidget::isEmpty();
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QGraphicsProxyWidget;
static std::mutex __connsMu_QGraphicsProxyWidget;

extern "C" {

void* QGraphicsProxyWidget_base(void* self) { return (void*)static_cast<QGraphicsWidget*>((QGraphicsProxyWidget*)self); }
void QGraphicsProxyWidget_chain(void* self, void** out) {
	QGraphicsProxyWidget* __p = (QGraphicsProxyWidget*)self;
	QGraphicsWidget* __b0 = static_cast<QGraphicsWidget*>(__p); out[0] = (void*)__b0;
	QGraphicsObject* __b1 = static_cast<QGraphicsObject*>(__b0); out[1] = (void*)__b1;
	QObject* __b2 = static_cast<QObject*>(__b1); out[2] = (void*)__b2;
}
void QGraphicsProxyWidget__set_override(void* self, int idx) { ((QGraphicsProxyWidget__Go*)self)->__goOn[idx] = 1; }
void* QGraphicsProxyWidget_fromQObject(void* o) { return (void*)qobject_cast<QGraphicsProxyWidget*>((QObject*)o); }
void* QGraphicsProxyWidget_new(long long __goid) {
	return new QGraphicsProxyWidget__Go(__goid);
}
void* QGraphicsProxyWidget_new2(long long __goid, void* parent) {
	return new QGraphicsProxyWidget__Go(__goid, (QGraphicsItem*)parent);
}
void* QGraphicsProxyWidget_new3(long long __goid, void* parent, int wFlags) {
	return new QGraphicsProxyWidget__Go(__goid, (QGraphicsItem*)parent, QFlags<Qt::WindowType>(wFlags));
}
char* QGraphicsProxyWidget_tr(const char* s, long long* __out_len) {
	auto&& __r = QGraphicsProxyWidget::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QGraphicsProxyWidget_setWidget(void* self, void* widget) {
	((QGraphicsProxyWidget*)self)->setWidget((QWidget*)widget);
}
void* QGraphicsProxyWidget_widget(void* self) {
	auto&& __r = ((QGraphicsProxyWidget*)self)->widget();
	return (void*)(__r);
}
void* QGraphicsProxyWidget_subWidgetRect(void* self, void* widget) {
	auto&& __r = ((QGraphicsProxyWidget*)self)->subWidgetRect((QWidget*)widget);
	return new QRectF(__r);
}
void* QGraphicsProxyWidget_createProxyForChildWidget(void* self, void* child) {
	auto&& __r = ((QGraphicsProxyWidget*)self)->createProxyForChildWidget((QWidget*)child);
	return (void*)(__r);
}
void* QGraphicsProxyWidget_newProxyWidget(void* self, void* arg1) {
	auto&& __r = ((QGraphicsProxyWidget__Go*)self)->newProxyWidget((QWidget*)arg1);
	return (void*)(__r);
}
void QGraphicsProxyWidget_setTabOrder(void* first, void* second) {
	QGraphicsProxyWidget::setTabOrder((QGraphicsWidget*)first, (QGraphicsWidget*)second);
}
void QGraphicsProxyWidget_updateMicroFocus(void* self) {
	((QGraphicsProxyWidget__Go*)self)->updateMicroFocus();
}
void* QGraphicsProxyWidget_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QGraphicsProxyWidget::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QGraphicsProxyWidget_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QGraphicsProxyWidget::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QGraphicsProxyWidget_sender(void* self) {
	auto&& __r = ((QGraphicsProxyWidget__Go*)self)->sender();
	return (void*)(__r);
}
int QGraphicsProxyWidget_senderSignalIndex(void* self) {
	auto&& __r = ((QGraphicsProxyWidget__Go*)self)->senderSignalIndex();
	return __r;
}
int QGraphicsProxyWidget_receivers(void* self, const char* signal) {
	auto&& __r = ((QGraphicsProxyWidget__Go*)self)->receivers(signal);
	return __r;
}
int QGraphicsProxyWidget_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QGraphicsProxyWidget__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
void QGraphicsProxyWidget_addToIndex(void* self) {
	((QGraphicsProxyWidget__Go*)self)->addToIndex();
}
void QGraphicsProxyWidget_removeFromIndex(void* self) {
	((QGraphicsProxyWidget__Go*)self)->removeFromIndex();
}
void QGraphicsProxyWidget_prepareGeometryChange(void* self) {
	((QGraphicsProxyWidget__Go*)self)->prepareGeometryChange();
}
void QGraphicsProxyWidget_setGraphicsItem(void* self, void* item) {
	((QGraphicsProxyWidget__Go*)self)->setGraphicsItem((QGraphicsItem*)item);
}
void QGraphicsProxyWidget_setOwnedByLayout(void* self, int ownedByLayout) {
	((QGraphicsProxyWidget__Go*)self)->setOwnedByLayout(ownedByLayout);
}
char* QGraphicsProxyWidget_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QGraphicsProxyWidget::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QGraphicsProxyWidget_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QGraphicsProxyWidget::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QGraphicsProxyWidget_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QGraphicsProxyWidget::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QGraphicsProxyWidget_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QGraphicsProxyWidget::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QGraphicsProxyWidget_metaObject_default(void* self) {
	auto&& __r = ((QGraphicsProxyWidget*)self)->QGraphicsProxyWidget::metaObject();
	return (void*)(__r);
}
void QGraphicsProxyWidget_setGeometry_default(void* self, void* rect) {
	((QGraphicsProxyWidget*)self)->QGraphicsProxyWidget::setGeometry(*(QRectF*)rect);
}
void QGraphicsProxyWidget_paint_default(void* self, void* painter, void* option, void* widget) {
	((QGraphicsProxyWidget*)self)->QGraphicsProxyWidget::paint((QPainter*)painter, (QStyleOptionGraphicsItem*)option, (QWidget*)widget);
}
int QGraphicsProxyWidget_type_default(void* self) {
	auto&& __r = ((QGraphicsProxyWidget*)self)->QGraphicsProxyWidget::type();
	return __r;
}
void* QGraphicsProxyWidget_itemChange_default(void* self, int change, void* value) {
	auto&& __r = ((QGraphicsProxyWidget__Go*)self)->cppgoBase_itemChange(change, *(QVariant*)value);
	return new QVariant(__r);
}
int QGraphicsProxyWidget_event_default(void* self, void* event) {
	auto&& __r = ((QGraphicsProxyWidget__Go*)self)->cppgoBase_event((QEvent*)event);
	return __r;
}
int QGraphicsProxyWidget_eventFilter_default(void* self, void* object, void* event) {
	auto&& __r = ((QGraphicsProxyWidget__Go*)self)->cppgoBase_eventFilter((QObject*)object, (QEvent*)event);
	return __r;
}
void QGraphicsProxyWidget_showEvent_default(void* self, void* event) {
	((QGraphicsProxyWidget__Go*)self)->cppgoBase_showEvent((QShowEvent*)event);
}
void QGraphicsProxyWidget_hideEvent_default(void* self, void* event) {
	((QGraphicsProxyWidget__Go*)self)->cppgoBase_hideEvent((QHideEvent*)event);
}
void QGraphicsProxyWidget_contextMenuEvent_default(void* self, void* event) {
	((QGraphicsProxyWidget__Go*)self)->cppgoBase_contextMenuEvent((QGraphicsSceneContextMenuEvent*)event);
}
void QGraphicsProxyWidget_dragEnterEvent_default(void* self, void* event) {
	((QGraphicsProxyWidget__Go*)self)->cppgoBase_dragEnterEvent((QGraphicsSceneDragDropEvent*)event);
}
void QGraphicsProxyWidget_dragLeaveEvent_default(void* self, void* event) {
	((QGraphicsProxyWidget__Go*)self)->cppgoBase_dragLeaveEvent((QGraphicsSceneDragDropEvent*)event);
}
void QGraphicsProxyWidget_dragMoveEvent_default(void* self, void* event) {
	((QGraphicsProxyWidget__Go*)self)->cppgoBase_dragMoveEvent((QGraphicsSceneDragDropEvent*)event);
}
void QGraphicsProxyWidget_dropEvent_default(void* self, void* event) {
	((QGraphicsProxyWidget__Go*)self)->cppgoBase_dropEvent((QGraphicsSceneDragDropEvent*)event);
}
void QGraphicsProxyWidget_hoverEnterEvent_default(void* self, void* event) {
	((QGraphicsProxyWidget__Go*)self)->cppgoBase_hoverEnterEvent((QGraphicsSceneHoverEvent*)event);
}
void QGraphicsProxyWidget_hoverLeaveEvent_default(void* self, void* event) {
	((QGraphicsProxyWidget__Go*)self)->cppgoBase_hoverLeaveEvent((QGraphicsSceneHoverEvent*)event);
}
void QGraphicsProxyWidget_hoverMoveEvent_default(void* self, void* event) {
	((QGraphicsProxyWidget__Go*)self)->cppgoBase_hoverMoveEvent((QGraphicsSceneHoverEvent*)event);
}
void QGraphicsProxyWidget_grabMouseEvent_default(void* self, void* event) {
	((QGraphicsProxyWidget__Go*)self)->cppgoBase_grabMouseEvent((QEvent*)event);
}
void QGraphicsProxyWidget_ungrabMouseEvent_default(void* self, void* event) {
	((QGraphicsProxyWidget__Go*)self)->cppgoBase_ungrabMouseEvent((QEvent*)event);
}
void QGraphicsProxyWidget_mouseMoveEvent_default(void* self, void* event) {
	((QGraphicsProxyWidget__Go*)self)->cppgoBase_mouseMoveEvent((QGraphicsSceneMouseEvent*)event);
}
void QGraphicsProxyWidget_mousePressEvent_default(void* self, void* event) {
	((QGraphicsProxyWidget__Go*)self)->cppgoBase_mousePressEvent((QGraphicsSceneMouseEvent*)event);
}
void QGraphicsProxyWidget_mouseReleaseEvent_default(void* self, void* event) {
	((QGraphicsProxyWidget__Go*)self)->cppgoBase_mouseReleaseEvent((QGraphicsSceneMouseEvent*)event);
}
void QGraphicsProxyWidget_mouseDoubleClickEvent_default(void* self, void* event) {
	((QGraphicsProxyWidget__Go*)self)->cppgoBase_mouseDoubleClickEvent((QGraphicsSceneMouseEvent*)event);
}
void QGraphicsProxyWidget_wheelEvent_default(void* self, void* event) {
	((QGraphicsProxyWidget__Go*)self)->cppgoBase_wheelEvent((QGraphicsSceneWheelEvent*)event);
}
void QGraphicsProxyWidget_keyPressEvent_default(void* self, void* event) {
	((QGraphicsProxyWidget__Go*)self)->cppgoBase_keyPressEvent((QKeyEvent*)event);
}
void QGraphicsProxyWidget_keyReleaseEvent_default(void* self, void* event) {
	((QGraphicsProxyWidget__Go*)self)->cppgoBase_keyReleaseEvent((QKeyEvent*)event);
}
void QGraphicsProxyWidget_focusInEvent_default(void* self, void* event) {
	((QGraphicsProxyWidget__Go*)self)->cppgoBase_focusInEvent((QFocusEvent*)event);
}
void QGraphicsProxyWidget_focusOutEvent_default(void* self, void* event) {
	((QGraphicsProxyWidget__Go*)self)->cppgoBase_focusOutEvent((QFocusEvent*)event);
}
int QGraphicsProxyWidget_focusNextPrevChild_default(void* self, int next) {
	auto&& __r = ((QGraphicsProxyWidget__Go*)self)->cppgoBase_focusNextPrevChild(next);
	return __r;
}
void* QGraphicsProxyWidget_inputMethodQuery_default(void* self, int query) {
	auto&& __r = ((QGraphicsProxyWidget__Go*)self)->cppgoBase_inputMethodQuery(query);
	return new QVariant(__r);
}
void QGraphicsProxyWidget_inputMethodEvent_default(void* self, void* event) {
	((QGraphicsProxyWidget__Go*)self)->cppgoBase_inputMethodEvent((QInputMethodEvent*)event);
}
void* QGraphicsProxyWidget_sizeHint_default(void* self, int which, void* constraint) {
	auto&& __r = ((QGraphicsProxyWidget__Go*)self)->cppgoBase_sizeHint(which, *(QSizeF*)constraint);
	return new QSizeF(__r);
}
void QGraphicsProxyWidget_resizeEvent_default(void* self, void* event) {
	((QGraphicsProxyWidget__Go*)self)->cppgoBase_resizeEvent((QGraphicsSceneResizeEvent*)event);
}
void QGraphicsProxyWidget_getContentsMargins_default(void* self, double* left, double* top, double* right, double* bottom) {
	qreal left_local{};
	qreal top_local{};
	qreal right_local{};
	qreal bottom_local{};
	((QGraphicsProxyWidget*)self)->QGraphicsProxyWidget::getContentsMargins(&left_local, &top_local, &right_local, &bottom_local);
	*left = left_local;
	*top = top_local;
	*right = right_local;
	*bottom = bottom_local;
}
void QGraphicsProxyWidget_paintWindowFrame_default(void* self, void* painter, void* option, void* widget) {
	((QGraphicsProxyWidget*)self)->QGraphicsProxyWidget::paintWindowFrame((QPainter*)painter, (QStyleOptionGraphicsItem*)option, (QWidget*)widget);
}
void* QGraphicsProxyWidget_boundingRect_default(void* self) {
	auto&& __r = ((QGraphicsProxyWidget*)self)->QGraphicsProxyWidget::boundingRect();
	return new QRectF(__r);
}
void* QGraphicsProxyWidget_shape_default(void* self) {
	auto&& __r = ((QGraphicsProxyWidget*)self)->QGraphicsProxyWidget::shape();
	return new QPainterPath(__r);
}
void QGraphicsProxyWidget_initStyleOption_default(void* self, void* option) {
	((QGraphicsProxyWidget__Go*)self)->cppgoBase_initStyleOption((QStyleOption*)option);
}
void QGraphicsProxyWidget_updateGeometry_default(void* self) {
	((QGraphicsProxyWidget__Go*)self)->cppgoBase_updateGeometry();
}
void* QGraphicsProxyWidget_propertyChange_default(void* self, const char* propertyName, long long propertyName_len, void* value) {
	QString propertyName_qstr = QString::fromUtf8(propertyName, propertyName_len);
	auto&& __r = ((QGraphicsProxyWidget__Go*)self)->cppgoBase_propertyChange(propertyName_qstr, *(QVariant*)value);
	return new QVariant(__r);
}
int QGraphicsProxyWidget_sceneEvent_default(void* self, void* event) {
	auto&& __r = ((QGraphicsProxyWidget__Go*)self)->cppgoBase_sceneEvent((QEvent*)event);
	return __r;
}
int QGraphicsProxyWidget_windowFrameEvent_default(void* self, void* e) {
	auto&& __r = ((QGraphicsProxyWidget__Go*)self)->cppgoBase_windowFrameEvent((QEvent*)e);
	return __r;
}
int QGraphicsProxyWidget_windowFrameSectionAt_default(void* self, void* pos) {
	auto&& __r = ((QGraphicsProxyWidget__Go*)self)->cppgoBase_windowFrameSectionAt(*(QPointF*)pos);
	return (int)(__r);
}
void QGraphicsProxyWidget_changeEvent_default(void* self, void* event) {
	((QGraphicsProxyWidget__Go*)self)->cppgoBase_changeEvent((QEvent*)event);
}
void QGraphicsProxyWidget_closeEvent_default(void* self, void* event) {
	((QGraphicsProxyWidget__Go*)self)->cppgoBase_closeEvent((QCloseEvent*)event);
}
void QGraphicsProxyWidget_moveEvent_default(void* self, void* event) {
	((QGraphicsProxyWidget__Go*)self)->cppgoBase_moveEvent((QGraphicsSceneMoveEvent*)event);
}
void QGraphicsProxyWidget_polishEvent_default(void* self) {
	((QGraphicsProxyWidget__Go*)self)->cppgoBase_polishEvent();
}
void QGraphicsProxyWidget_grabKeyboardEvent_default(void* self, void* event) {
	((QGraphicsProxyWidget__Go*)self)->cppgoBase_grabKeyboardEvent((QEvent*)event);
}
void QGraphicsProxyWidget_ungrabKeyboardEvent_default(void* self, void* event) {
	((QGraphicsProxyWidget__Go*)self)->cppgoBase_ungrabKeyboardEvent((QEvent*)event);
}
void QGraphicsProxyWidget_timerEvent_default(void* self, void* event) {
	((QGraphicsProxyWidget__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QGraphicsProxyWidget_childEvent_default(void* self, void* event) {
	((QGraphicsProxyWidget__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QGraphicsProxyWidget_customEvent_default(void* self, void* event) {
	((QGraphicsProxyWidget__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QGraphicsProxyWidget_connectNotify_default(void* self, void* signal) {
	((QGraphicsProxyWidget__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QGraphicsProxyWidget_disconnectNotify_default(void* self, void* signal) {
	((QGraphicsProxyWidget__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QGraphicsProxyWidget_advance_default(void* self, int phase) {
	((QGraphicsProxyWidget*)self)->QGraphicsProxyWidget::advance(phase);
}
int QGraphicsProxyWidget_contains_default(void* self, void* point) {
	auto&& __r = ((QGraphicsProxyWidget*)self)->QGraphicsProxyWidget::contains(*(QPointF*)point);
	return __r;
}
int QGraphicsProxyWidget_collidesWithItem_default(void* self, void* other, int mode) {
	auto&& __r = ((QGraphicsProxyWidget*)self)->QGraphicsProxyWidget::collidesWithItem((QGraphicsItem*)other, (Qt::ItemSelectionMode)(mode));
	return __r;
}
int QGraphicsProxyWidget_collidesWithPath_default(void* self, void* path, int mode) {
	auto&& __r = ((QGraphicsProxyWidget*)self)->QGraphicsProxyWidget::collidesWithPath(*(QPainterPath*)path, (Qt::ItemSelectionMode)(mode));
	return __r;
}
int QGraphicsProxyWidget_isObscuredBy_default(void* self, void* item) {
	auto&& __r = ((QGraphicsProxyWidget*)self)->QGraphicsProxyWidget::isObscuredBy((QGraphicsItem*)item);
	return __r;
}
void* QGraphicsProxyWidget_opaqueArea_default(void* self) {
	auto&& __r = ((QGraphicsProxyWidget*)self)->QGraphicsProxyWidget::opaqueArea();
	return new QPainterPath(__r);
}
int QGraphicsProxyWidget_sceneEventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QGraphicsProxyWidget__Go*)self)->cppgoBase_sceneEventFilter((QGraphicsItem*)watched, (QEvent*)event);
	return __r;
}
int QGraphicsProxyWidget_supportsExtension_default(void* self, int extension) {
	auto&& __r = ((QGraphicsProxyWidget__Go*)self)->cppgoBase_supportsExtension(extension);
	return __r;
}
void QGraphicsProxyWidget_setExtension_default(void* self, int extension, void* variant) {
	((QGraphicsProxyWidget__Go*)self)->cppgoBase_setExtension(extension, *(QVariant*)variant);
}
void* QGraphicsProxyWidget_extension_default(void* self, void* variant) {
	auto&& __r = ((QGraphicsProxyWidget__Go*)self)->cppgoBase_extension(*(QVariant*)variant);
	return new QVariant(__r);
}
int QGraphicsProxyWidget_isEmpty_default(void* self) {
	auto&& __r = ((QGraphicsProxyWidget*)self)->QGraphicsProxyWidget::isEmpty();
	return __r;
}
void QGraphicsProxyWidget_connect_geometryChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsProxyWidget*)self, &QGraphicsProxyWidget::geometryChanged, (QGraphicsProxyWidget*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsProxyWidget);
	__conns_QGraphicsProxyWidget[__id] = __c;
}
void QGraphicsProxyWidget_connect_layoutChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsProxyWidget*)self, &QGraphicsProxyWidget::layoutChanged, (QGraphicsProxyWidget*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsProxyWidget);
	__conns_QGraphicsProxyWidget[__id] = __c;
}
void QGraphicsProxyWidget_connect_parentChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsProxyWidget*)self, &QGraphicsProxyWidget::parentChanged, (QGraphicsProxyWidget*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsProxyWidget);
	__conns_QGraphicsProxyWidget[__id] = __c;
}
void QGraphicsProxyWidget_connect_opacityChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsProxyWidget*)self, &QGraphicsProxyWidget::opacityChanged, (QGraphicsProxyWidget*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsProxyWidget);
	__conns_QGraphicsProxyWidget[__id] = __c;
}
void QGraphicsProxyWidget_connect_visibleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsProxyWidget*)self, &QGraphicsProxyWidget::visibleChanged, (QGraphicsProxyWidget*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsProxyWidget);
	__conns_QGraphicsProxyWidget[__id] = __c;
}
void QGraphicsProxyWidget_connect_enabledChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsProxyWidget*)self, &QGraphicsProxyWidget::enabledChanged, (QGraphicsProxyWidget*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsProxyWidget);
	__conns_QGraphicsProxyWidget[__id] = __c;
}
void QGraphicsProxyWidget_connect_xChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsProxyWidget*)self, &QGraphicsProxyWidget::xChanged, (QGraphicsProxyWidget*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsProxyWidget);
	__conns_QGraphicsProxyWidget[__id] = __c;
}
void QGraphicsProxyWidget_connect_yChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsProxyWidget*)self, &QGraphicsProxyWidget::yChanged, (QGraphicsProxyWidget*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsProxyWidget);
	__conns_QGraphicsProxyWidget[__id] = __c;
}
void QGraphicsProxyWidget_connect_zChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsProxyWidget*)self, &QGraphicsProxyWidget::zChanged, (QGraphicsProxyWidget*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsProxyWidget);
	__conns_QGraphicsProxyWidget[__id] = __c;
}
void QGraphicsProxyWidget_connect_rotationChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsProxyWidget*)self, &QGraphicsProxyWidget::rotationChanged, (QGraphicsProxyWidget*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsProxyWidget);
	__conns_QGraphicsProxyWidget[__id] = __c;
}
void QGraphicsProxyWidget_connect_scaleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsProxyWidget*)self, &QGraphicsProxyWidget::scaleChanged, (QGraphicsProxyWidget*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsProxyWidget);
	__conns_QGraphicsProxyWidget[__id] = __c;
}
void QGraphicsProxyWidget_connect_childrenChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsProxyWidget*)self, &QGraphicsProxyWidget::childrenChanged, (QGraphicsProxyWidget*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsProxyWidget);
	__conns_QGraphicsProxyWidget[__id] = __c;
}
void QGraphicsProxyWidget_connect_widthChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsProxyWidget*)self, &QGraphicsProxyWidget::widthChanged, (QGraphicsProxyWidget*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsProxyWidget);
	__conns_QGraphicsProxyWidget[__id] = __c;
}
void QGraphicsProxyWidget_connect_heightChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsProxyWidget*)self, &QGraphicsProxyWidget::heightChanged, (QGraphicsProxyWidget*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsProxyWidget);
	__conns_QGraphicsProxyWidget[__id] = __c;
}
void QGraphicsProxyWidget_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsProxyWidget*)self, &QGraphicsProxyWidget::destroyed, (QGraphicsProxyWidget*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsProxyWidget);
	__conns_QGraphicsProxyWidget[__id] = __c;
}
void QGraphicsProxyWidget_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsProxyWidget*)self, &QGraphicsProxyWidget::destroyed, (QGraphicsProxyWidget*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsProxyWidget);
	__conns_QGraphicsProxyWidget[__id] = __c;
}
void QGraphicsProxyWidget_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsProxyWidget*)self, &QGraphicsProxyWidget::objectNameChanged, (QGraphicsProxyWidget*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsProxyWidget);
	__conns_QGraphicsProxyWidget[__id] = __c;
}
void QGraphicsProxyWidget__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsProxyWidget);
		auto __it = __conns_QGraphicsProxyWidget.find(id);
		if (__it == __conns_QGraphicsProxyWidget.end()) return;
		__c = __it->second;
		__conns_QGraphicsProxyWidget.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QGraphicsProxyWidget__Go_new(long long id) { return new QGraphicsProxyWidget__Go(id); }
void QGraphicsProxyWidget_delete(void* self) { delete (QGraphicsProxyWidget*)self; }

}

#include <QtWidgets/qlabel.h>
#include <QtWidgets/qframe.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qstring.h>
#include <QtCore/qobjectdefs.h>
#include <QtGui/qpixmap.h>
#include <QtGui/qpicture.h>
#include <QtGui/qmovie.h>
#include <QtCore/qsize.h>
#include <QtCore/qcoreevent.h>
#include <QtGui/qevent.h>
#include <QtCore/qrect.h>
#include <QtGui/qpainter.h>
#include <QtWidgets/qstyleoption.h>
#include <QtWidgets/qstyle.h>
#include <QtCore/qpoint.h>
#include <QtGui/qregion.h>
#include <QtGui/qpalette.h>
#include <QtGui/qfont.h>
#include <QtGui/qfontmetrics.h>
#include <QtGui/qfontinfo.h>
#include <QtGui/qcursor.h>
#include <QtGui/qbitmap.h>
#include <QtGui/qpaintdevice.h>
#include <QtWidgets/qgraphicseffect.h>
#include <QtGui/qicon.h>
#include <QtCore/qlocale.h>
#include <QtGui/qkeysequence.h>
#include <QtWidgets/qgraphicsproxywidget.h>
#include <QtCore/qbytearray.h>
#include <QtWidgets/qsizepolicy.h>
#include <QtCore/qmargins.h>
#include <QtWidgets/qlayout.h>
#include <QtGui/qaction.h>
#include <QtCore/qobject.h>
#include <QtGui/qpaintengine.h>
#include <QtGui/qbackingstore.h>
#include <QtGui/qwindow.h>
#include <QtGui/qscreen.h>
#include <QtCore/qvariant.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qlabel.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QLabel__Go : public QLabel {
public:
	long long __goId;
	unsigned char __goOn[49] = {};
	template<typename... A> QLabel__Go(long long id, A&&... a) : QLabel(std::forward<A>(a)...), __goId(id) {}
	using QLabel::drawFrame;
	using QLabel::updateMicroFocus;
	using QLabel::create;
	using QLabel::destroy;
	using QLabel::focusNextChild;
	using QLabel::focusPreviousChild;
	using QLabel::sender;
	using QLabel::senderSignalIndex;
	using QLabel::receivers;
	using QLabel::isSignalConnected;
	using QLabel::getDecodedMetricF;
	bool cppgoBase_event(QEvent* __a0) { return QLabel::event(__a0); }
	void cppgoBase_keyPressEvent(QKeyEvent* __a0) { QLabel::keyPressEvent(__a0); }
	void cppgoBase_paintEvent(QPaintEvent* __a0) { QLabel::paintEvent(__a0); }
	void cppgoBase_changeEvent(QEvent* __a0) { QLabel::changeEvent(__a0); }
	void cppgoBase_mousePressEvent(QMouseEvent* __a0) { QLabel::mousePressEvent(__a0); }
	void cppgoBase_mouseMoveEvent(QMouseEvent* __a0) { QLabel::mouseMoveEvent(__a0); }
	void cppgoBase_mouseReleaseEvent(QMouseEvent* __a0) { QLabel::mouseReleaseEvent(__a0); }
	void cppgoBase_contextMenuEvent(QContextMenuEvent* __a0) { QLabel::contextMenuEvent(__a0); }
	void cppgoBase_focusInEvent(QFocusEvent* __a0) { QLabel::focusInEvent(__a0); }
	void cppgoBase_focusOutEvent(QFocusEvent* __a0) { QLabel::focusOutEvent(__a0); }
	bool cppgoBase_focusNextPrevChild(bool __a0) { return QLabel::focusNextPrevChild(__a0); }
	void cppgoBase_initStyleOption(QStyleOptionFrame* __a0) { QLabel::initStyleOption(__a0); }
	void cppgoBase_mouseDoubleClickEvent(QMouseEvent* __a0) { QLabel::mouseDoubleClickEvent(__a0); }
	void cppgoBase_wheelEvent(QWheelEvent* __a0) { QLabel::wheelEvent(__a0); }
	void cppgoBase_keyReleaseEvent(QKeyEvent* __a0) { QLabel::keyReleaseEvent(__a0); }
	void cppgoBase_enterEvent(QEnterEvent* __a0) { QLabel::enterEvent(__a0); }
	void cppgoBase_leaveEvent(QEvent* __a0) { QLabel::leaveEvent(__a0); }
	void cppgoBase_moveEvent(QMoveEvent* __a0) { QLabel::moveEvent(__a0); }
	void cppgoBase_resizeEvent(QResizeEvent* __a0) { QLabel::resizeEvent(__a0); }
	void cppgoBase_closeEvent(QCloseEvent* __a0) { QLabel::closeEvent(__a0); }
	void cppgoBase_tabletEvent(QTabletEvent* __a0) { QLabel::tabletEvent(__a0); }
	void cppgoBase_actionEvent(QActionEvent* __a0) { QLabel::actionEvent(__a0); }
	void cppgoBase_dragEnterEvent(QDragEnterEvent* __a0) { QLabel::dragEnterEvent(__a0); }
	void cppgoBase_dragMoveEvent(QDragMoveEvent* __a0) { QLabel::dragMoveEvent(__a0); }
	void cppgoBase_dragLeaveEvent(QDragLeaveEvent* __a0) { QLabel::dragLeaveEvent(__a0); }
	void cppgoBase_dropEvent(QDropEvent* __a0) { QLabel::dropEvent(__a0); }
	void cppgoBase_showEvent(QShowEvent* __a0) { QLabel::showEvent(__a0); }
	void cppgoBase_hideEvent(QHideEvent* __a0) { QLabel::hideEvent(__a0); }
	int cppgoBase_metric(int __a0) { return QLabel::metric((QPaintDevice::PaintDeviceMetric)__a0); }
	void cppgoBase_initPainter(QPainter* __a0) { QLabel::initPainter(__a0); }
	QPaintDevice* cppgoBase_redirected(QPoint* __a0) { return QLabel::redirected(__a0); }
	QPainter* cppgoBase_sharedPainter() { return QLabel::sharedPainter(); }
	void cppgoBase_inputMethodEvent(QInputMethodEvent* __a0) { QLabel::inputMethodEvent(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QLabel::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QLabel::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QLabel::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QLabel::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QLabel::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QLabel::metaObject();
	}
	QSize sizeHint() const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 729690261457684101LL, &__h); if (__h) return *(QSize*)__r; }
		return QLabel::sizeHint();
	}
	QSize minimumSizeHint() const override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 7049264599837532491LL, &__h); if (__h) return *(QSize*)__r; }
		return QLabel::minimumSizeHint();
	}
	int heightForWidth(int arg1) const override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 4311534060495737339LL, arg1, &__h); if (__h) return __r; }
		return QLabel::heightForWidth(arg1);
	}
	bool event(QEvent* e) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)e, &__h); if (__h) return __r; }
		return QLabel::event(e);
	}
	void keyPressEvent(QKeyEvent* ev) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1023791977679793185LL, (void*)ev, &__h); if (__h) return; }
		QLabel::keyPressEvent(ev);
	}
	void paintEvent(QPaintEvent* arg1) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5290442163167458465LL, (void*)arg1, &__h); if (__h) return; }
		QLabel::paintEvent(arg1);
	}
	void changeEvent(QEvent* arg1) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4959376063826165943LL, (void*)arg1, &__h); if (__h) return; }
		QLabel::changeEvent(arg1);
	}
	void mousePressEvent(QMouseEvent* ev) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -332090898078543319LL, (void*)ev, &__h); if (__h) return; }
		QLabel::mousePressEvent(ev);
	}
	void mouseMoveEvent(QMouseEvent* ev) override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4255494004099806581LL, (void*)ev, &__h); if (__h) return; }
		QLabel::mouseMoveEvent(ev);
	}
	void mouseReleaseEvent(QMouseEvent* ev) override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4415920049297795213LL, (void*)ev, &__h); if (__h) return; }
		QLabel::mouseReleaseEvent(ev);
	}
	void contextMenuEvent(QContextMenuEvent* ev) override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6698043222227006015LL, (void*)ev, &__h); if (__h) return; }
		QLabel::contextMenuEvent(ev);
	}
	void focusInEvent(QFocusEvent* ev) override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 975043767033552552LL, (void*)ev, &__h); if (__h) return; }
		QLabel::focusInEvent(ev);
	}
	void focusOutEvent(QFocusEvent* ev) override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -385596010576048993LL, (void*)ev, &__h); if (__h) return; }
		QLabel::focusOutEvent(ev);
	}
	bool focusNextPrevChild(bool next) override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 5898328517505784191LL, next, &__h); if (__h) return __r; }
		return QLabel::focusNextPrevChild(next);
	}
	void initStyleOption(QStyleOptionFrame* option) const override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -488693769771308531LL, (void*)option, &__h); if (__h) return; }
		QLabel::initStyleOption(option);
	}
	int devType() const override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8659679893868524044LL, &__h); if (__h) return __r; }
		return QLabel::devType();
	}
	void setVisible(bool visible) override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 2110182719375245575LL, visible, &__h); if (__h) return; }
		QLabel::setVisible(visible);
	}
	bool hasHeightForWidth() const override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8266762397995423915LL, &__h); if (__h) return __r; }
		return QLabel::hasHeightForWidth();
	}
	QPaintEngine* paintEngine() const override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 6398007404756346381LL, &__h); if (__h) return (QPaintEngine*)__r; }
		return QLabel::paintEngine();
	}
	void mouseDoubleClickEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4029073905697023887LL, (void*)event, &__h); if (__h) return; }
		QLabel::mouseDoubleClickEvent(event);
	}
	void wheelEvent(QWheelEvent* event) override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1919988340934511316LL, (void*)event, &__h); if (__h) return; }
		QLabel::wheelEvent(event);
	}
	void keyReleaseEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 772598475950050299LL, (void*)event, &__h); if (__h) return; }
		QLabel::keyReleaseEvent(event);
	}
	void enterEvent(QEnterEvent* event) override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6316270346668237895LL, (void*)event, &__h); if (__h) return; }
		QLabel::enterEvent(event);
	}
	void leaveEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -551070009623711640LL, (void*)event, &__h); if (__h) return; }
		QLabel::leaveEvent(event);
	}
	void moveEvent(QMoveEvent* event) override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -9161949150489525396LL, (void*)event, &__h); if (__h) return; }
		QLabel::moveEvent(event);
	}
	void resizeEvent(QResizeEvent* event) override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4302917285655014661LL, (void*)event, &__h); if (__h) return; }
		QLabel::resizeEvent(event);
	}
	void closeEvent(QCloseEvent* event) override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1989937050019464783LL, (void*)event, &__h); if (__h) return; }
		QLabel::closeEvent(event);
	}
	void tabletEvent(QTabletEvent* event) override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -212068866860701141LL, (void*)event, &__h); if (__h) return; }
		QLabel::tabletEvent(event);
	}
	void actionEvent(QActionEvent* event) override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1697444140088555581LL, (void*)event, &__h); if (__h) return; }
		QLabel::actionEvent(event);
	}
	void dragEnterEvent(QDragEnterEvent* event) override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 206122715919917589LL, (void*)event, &__h); if (__h) return; }
		QLabel::dragEnterEvent(event);
	}
	void dragMoveEvent(QDragMoveEvent* event) override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1781545137352745978LL, (void*)event, &__h); if (__h) return; }
		QLabel::dragMoveEvent(event);
	}
	void dragLeaveEvent(QDragLeaveEvent* event) override {
		int __h = 0;
		if (__goOn[32] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -3139904158373725006LL, (void*)event, &__h); if (__h) return; }
		QLabel::dragLeaveEvent(event);
	}
	void dropEvent(QDropEvent* event) override {
		int __h = 0;
		if (__goOn[33] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 9107058682754591716LL, (void*)event, &__h); if (__h) return; }
		QLabel::dropEvent(event);
	}
	void showEvent(QShowEvent* event) override {
		int __h = 0;
		if (__goOn[34] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6825369174751916LL, (void*)event, &__h); if (__h) return; }
		QLabel::showEvent(event);
	}
	void hideEvent(QHideEvent* event) override {
		int __h = 0;
		if (__goOn[35] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -199718707876719129LL, (void*)event, &__h); if (__h) return; }
		QLabel::hideEvent(event);
	}
	bool nativeEvent(const QByteArray& eventType, void* message, qintptr* result) override {
		int __h = 0;
		if (__goOn[36] && cppgo_virt_int_voidp_voidp_longlongp) { int __r = ((int(*)(long long, long long, void*, void*, long long*, int*))cppgo_virt_int_voidp_voidp_longlongp)(__goId, -2039166586246443706LL, (void*)&eventType, (void*)message, (long long*)result, &__h); if (__h) return __r; }
		return QLabel::nativeEvent(eventType, message, result);
	}
	int metric(QPaintDevice::PaintDeviceMetric arg1) const override {
		int __h = 0;
		if (__goOn[37] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 427629649267311777LL, (int)(arg1), &__h); if (__h) return __r; }
		return QLabel::metric(arg1);
	}
	void initPainter(QPainter* painter) const override {
		int __h = 0;
		if (__goOn[38] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6477500296980050210LL, (void*)painter, &__h); if (__h) return; }
		QLabel::initPainter(painter);
	}
	QPaintDevice* redirected(QPoint* offset) const override {
		int __h = 0;
		if (__goOn[39] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 4127240633310655048LL, (void*)offset, &__h); if (__h) return (QPaintDevice*)__r; }
		return QLabel::redirected(offset);
	}
	QPainter* sharedPainter() const override {
		int __h = 0;
		if (__goOn[40] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 5792771328741774991LL, &__h); if (__h) return (QPainter*)__r; }
		return QLabel::sharedPainter();
	}
	void inputMethodEvent(QInputMethodEvent* arg1) override {
		int __h = 0;
		if (__goOn[41] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5438819661942843222LL, (void*)arg1, &__h); if (__h) return; }
		QLabel::inputMethodEvent(arg1);
	}
	QVariant inputMethodQuery(Qt::InputMethodQuery arg1) const override {
		int __h = 0;
		if (__goOn[42] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7345853097340709638LL, (int)(arg1), &__h); if (__h) return *(QVariant*)__r; }
		return QLabel::inputMethodQuery(arg1);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[43] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QLabel::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[44] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QLabel::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[45] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QLabel::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[46] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QLabel::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[47] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QLabel::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[48] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QLabel::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QLabel;
static std::mutex __connsMu_QLabel;

extern "C" {

void* QLabel_base(void* self) { return (void*)static_cast<QFrame*>((QLabel*)self); }
void QLabel_chain(void* self, void** out) {
	QLabel* __p = (QLabel*)self;
	QFrame* __b0 = static_cast<QFrame*>(__p); out[0] = (void*)__b0;
	QWidget* __b1 = static_cast<QWidget*>(__b0); out[1] = (void*)__b1;
	QObject* __b2 = static_cast<QObject*>(__b1); out[2] = (void*)__b2;
}
void QLabel__set_override(void* self, int idx) { ((QLabel__Go*)self)->__goOn[idx] = 1; }
void* QLabel_fromQObject(void* o) { return (void*)qobject_cast<QLabel*>((QObject*)o); }
void* QLabel_new(long long __goid, void* parent) {
	return new QLabel__Go(__goid, (QWidget*)parent);
}
void* QLabel_new2(long long __goid) {
	return new QLabel__Go(__goid);
}
void* QLabel_new3(long long __goid, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	return new QLabel__Go(__goid, text_qstr);
}
void* QLabel_new4(long long __goid, void* parent, int f) {
	return new QLabel__Go(__goid, (QWidget*)parent, QFlags<Qt::WindowType>(f));
}
void* QLabel_new5(long long __goid, const char* text, long long text_len, void* parent) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	return new QLabel__Go(__goid, text_qstr, (QWidget*)parent);
}
void* QLabel_new6(long long __goid, const char* text, long long text_len, void* parent, int f) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	return new QLabel__Go(__goid, text_qstr, (QWidget*)parent, QFlags<Qt::WindowType>(f));
}
char* QLabel_tr(const char* s, long long* __out_len) {
	auto&& __r = QLabel::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLabel_text(void* self, long long* __out_len) {
	auto&& __r = ((QLabel*)self)->text();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QLabel_pixmap(void* self, int arg1) {
	auto&& __r = ((QLabel*)self)->pixmap((Qt::ReturnByValueConstant)(arg1));
	return new QPixmap(__r);
}
void* QLabel_pixmap2(void* self) {
	auto&& __r = ((QLabel*)self)->pixmap();
	return new QPixmap(__r);
}
void* QLabel_picture(void* self, int arg1) {
	auto&& __r = ((QLabel*)self)->picture((Qt::ReturnByValueConstant)(arg1));
	return new QPicture(__r);
}
void* QLabel_picture2(void* self) {
	auto&& __r = ((QLabel*)self)->picture();
	return new QPicture(__r);
}
void* QLabel_movie(void* self) {
	auto&& __r = ((QLabel*)self)->movie();
	return (void*)(__r);
}
int QLabel_textFormat(void* self) {
	auto&& __r = ((QLabel*)self)->textFormat();
	return (int)(__r);
}
void QLabel_setTextFormat(void* self, int arg1) {
	((QLabel*)self)->setTextFormat((Qt::TextFormat)(arg1));
}
int QLabel_alignment(void* self) {
	auto&& __r = ((QLabel*)self)->alignment();
	return (int)(__r);
}
void QLabel_setAlignment(void* self, int arg1) {
	((QLabel*)self)->setAlignment(QFlags<Qt::AlignmentFlag>(arg1));
}
void QLabel_setWordWrap(void* self, int on) {
	((QLabel*)self)->setWordWrap(on);
}
int QLabel_wordWrap(void* self) {
	auto&& __r = ((QLabel*)self)->wordWrap();
	return __r;
}
int QLabel_indent(void* self) {
	auto&& __r = ((QLabel*)self)->indent();
	return __r;
}
void QLabel_setIndent(void* self, int arg1) {
	((QLabel*)self)->setIndent(arg1);
}
int QLabel_margin(void* self) {
	auto&& __r = ((QLabel*)self)->margin();
	return __r;
}
void QLabel_setMargin(void* self, int arg1) {
	((QLabel*)self)->setMargin(arg1);
}
int QLabel_hasScaledContents(void* self) {
	auto&& __r = ((QLabel*)self)->hasScaledContents();
	return __r;
}
void QLabel_setScaledContents(void* self, int arg1) {
	((QLabel*)self)->setScaledContents(arg1);
}
void QLabel_setBuddy(void* self, void* arg1) {
	((QLabel*)self)->setBuddy((QWidget*)arg1);
}
void* QLabel_buddy(void* self) {
	auto&& __r = ((QLabel*)self)->buddy();
	return (void*)(__r);
}
int QLabel_openExternalLinks(void* self) {
	auto&& __r = ((QLabel*)self)->openExternalLinks();
	return __r;
}
void QLabel_setOpenExternalLinks(void* self, int open) {
	((QLabel*)self)->setOpenExternalLinks(open);
}
void QLabel_setTextInteractionFlags(void* self, int flags) {
	((QLabel*)self)->setTextInteractionFlags(QFlags<Qt::TextInteractionFlag>(flags));
}
int QLabel_textInteractionFlags(void* self) {
	auto&& __r = ((QLabel*)self)->textInteractionFlags();
	return (int)(__r);
}
void QLabel_setSelection(void* self, int arg1, int arg2) {
	((QLabel*)self)->setSelection(arg1, arg2);
}
int QLabel_hasSelectedText(void* self) {
	auto&& __r = ((QLabel*)self)->hasSelectedText();
	return __r;
}
char* QLabel_selectedText(void* self, long long* __out_len) {
	auto&& __r = ((QLabel*)self)->selectedText();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QLabel_selectionStart(void* self) {
	auto&& __r = ((QLabel*)self)->selectionStart();
	return __r;
}
void QLabel_setText(void* self, const char* arg1, long long arg1_len) {
	QString arg1_qstr = QString::fromUtf8(arg1, arg1_len);
	((QLabel*)self)->setText(arg1_qstr);
}
void QLabel_setPixmap(void* self, void* arg1) {
	((QLabel*)self)->setPixmap(*(QPixmap*)arg1);
}
void QLabel_setPicture(void* self, void* arg1) {
	((QLabel*)self)->setPicture(*(QPicture*)arg1);
}
void QLabel_setMovie(void* self, void* movie) {
	((QLabel*)self)->setMovie((QMovie*)movie);
}
void QLabel_setNum(void* self, int arg1) {
	((QLabel*)self)->setNum(arg1);
}
void QLabel_setNum2(void* self, double arg1) {
	((QLabel*)self)->setNum(arg1);
}
void QLabel_clear(void* self) {
	((QLabel*)self)->clear();
}
void QLabel_linkActivated(void* self, const char* link, long long link_len) {
	QString link_qstr = QString::fromUtf8(link, link_len);
	((QLabel*)self)->linkActivated(link_qstr);
}
void QLabel_linkHovered(void* self, const char* link, long long link_len) {
	QString link_qstr = QString::fromUtf8(link, link_len);
	((QLabel*)self)->linkHovered(link_qstr);
}
void QLabel_drawFrame(void* self, void* arg1) {
	((QLabel__Go*)self)->drawFrame((QPainter*)arg1);
}
void QLabel_setTabOrder(void* arg1, void* arg2) {
	QLabel::setTabOrder((QWidget*)arg1, (QWidget*)arg2);
}
void* QLabel_mouseGrabber() {
	auto&& __r = QLabel::mouseGrabber();
	return (void*)(__r);
}
void* QLabel_keyboardGrabber() {
	auto&& __r = QLabel::keyboardGrabber();
	return (void*)(__r);
}
void* QLabel_find(unsigned long long arg1) {
	auto&& __r = QLabel::find(arg1);
	return (void*)(__r);
}
void* QLabel_createWindowContainer(void* window) {
	auto&& __r = QLabel::createWindowContainer((QWindow*)window);
	return (void*)(__r);
}
void QLabel_updateMicroFocus(void* self) {
	((QLabel__Go*)self)->updateMicroFocus();
}
void QLabel_create(void* self) {
	((QLabel__Go*)self)->create();
}
void QLabel_destroy(void* self) {
	((QLabel__Go*)self)->destroy();
}
int QLabel_focusNextChild(void* self) {
	auto&& __r = ((QLabel__Go*)self)->focusNextChild();
	return __r;
}
int QLabel_focusPreviousChild(void* self) {
	auto&& __r = ((QLabel__Go*)self)->focusPreviousChild();
	return __r;
}
void* QLabel_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QLabel::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QLabel_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QLabel::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QLabel_sender(void* self) {
	auto&& __r = ((QLabel__Go*)self)->sender();
	return (void*)(__r);
}
int QLabel_senderSignalIndex(void* self) {
	auto&& __r = ((QLabel__Go*)self)->senderSignalIndex();
	return __r;
}
int QLabel_receivers(void* self, const char* signal) {
	auto&& __r = ((QLabel__Go*)self)->receivers(signal);
	return __r;
}
int QLabel_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QLabel__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
double QLabel_devicePixelRatioFScale() {
	auto&& __r = QLabel::devicePixelRatioFScale();
	return __r;
}
int QLabel_encodeMetricF(int metric, double value) {
	auto&& __r = QLabel::encodeMetricF((QPaintDevice::PaintDeviceMetric)(metric), value);
	return __r;
}
double QLabel_getDecodedMetricF(void* self, int metricA, int metricB) {
	auto&& __r = ((QLabel__Go*)self)->getDecodedMetricF((QPaintDevice::PaintDeviceMetric)(metricA), (QPaintDevice::PaintDeviceMetric)(metricB));
	return __r;
}
char* QLabel_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QLabel::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLabel_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QLabel::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QLabel_createWindowContainer2(void* window, void* parent) {
	auto&& __r = QLabel::createWindowContainer((QWindow*)window, (QWidget*)parent);
	return (void*)(__r);
}
void* QLabel_createWindowContainer3(void* window, void* parent, int flags) {
	auto&& __r = QLabel::createWindowContainer((QWindow*)window, (QWidget*)parent, QFlags<Qt::WindowType>(flags));
	return (void*)(__r);
}
void QLabel_updateMicroFocus2(void* self, int query) {
	((QLabel__Go*)self)->updateMicroFocus((Qt::InputMethodQuery)(query));
}
void QLabel_create2(void* self, unsigned long long arg1) {
	((QLabel__Go*)self)->create(arg1);
}
void QLabel_create3(void* self, unsigned long long arg1, int initializeWindow) {
	((QLabel__Go*)self)->create(arg1, initializeWindow);
}
void QLabel_create4(void* self, unsigned long long arg1, int initializeWindow, int destroyOldWindow) {
	((QLabel__Go*)self)->create(arg1, initializeWindow, destroyOldWindow);
}
void QLabel_destroy2(void* self, int destroyWindow) {
	((QLabel__Go*)self)->destroy(destroyWindow);
}
void QLabel_destroy3(void* self, int destroyWindow, int destroySubWindows) {
	((QLabel__Go*)self)->destroy(destroyWindow, destroySubWindows);
}
void* QLabel_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QLabel::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QLabel_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QLabel::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QLabel_metaObject_default(void* self) {
	auto&& __r = ((QLabel*)self)->QLabel::metaObject();
	return (void*)(__r);
}
void* QLabel_sizeHint_default(void* self) {
	auto&& __r = ((QLabel*)self)->QLabel::sizeHint();
	return new QSize(__r);
}
void* QLabel_minimumSizeHint_default(void* self) {
	auto&& __r = ((QLabel*)self)->QLabel::minimumSizeHint();
	return new QSize(__r);
}
int QLabel_heightForWidth_default(void* self, int arg1) {
	auto&& __r = ((QLabel*)self)->QLabel::heightForWidth(arg1);
	return __r;
}
int QLabel_event_default(void* self, void* e) {
	auto&& __r = ((QLabel__Go*)self)->cppgoBase_event((QEvent*)e);
	return __r;
}
void QLabel_keyPressEvent_default(void* self, void* ev) {
	((QLabel__Go*)self)->cppgoBase_keyPressEvent((QKeyEvent*)ev);
}
void QLabel_paintEvent_default(void* self, void* arg1) {
	((QLabel__Go*)self)->cppgoBase_paintEvent((QPaintEvent*)arg1);
}
void QLabel_changeEvent_default(void* self, void* arg1) {
	((QLabel__Go*)self)->cppgoBase_changeEvent((QEvent*)arg1);
}
void QLabel_mousePressEvent_default(void* self, void* ev) {
	((QLabel__Go*)self)->cppgoBase_mousePressEvent((QMouseEvent*)ev);
}
void QLabel_mouseMoveEvent_default(void* self, void* ev) {
	((QLabel__Go*)self)->cppgoBase_mouseMoveEvent((QMouseEvent*)ev);
}
void QLabel_mouseReleaseEvent_default(void* self, void* ev) {
	((QLabel__Go*)self)->cppgoBase_mouseReleaseEvent((QMouseEvent*)ev);
}
void QLabel_contextMenuEvent_default(void* self, void* ev) {
	((QLabel__Go*)self)->cppgoBase_contextMenuEvent((QContextMenuEvent*)ev);
}
void QLabel_focusInEvent_default(void* self, void* ev) {
	((QLabel__Go*)self)->cppgoBase_focusInEvent((QFocusEvent*)ev);
}
void QLabel_focusOutEvent_default(void* self, void* ev) {
	((QLabel__Go*)self)->cppgoBase_focusOutEvent((QFocusEvent*)ev);
}
int QLabel_focusNextPrevChild_default(void* self, int next) {
	auto&& __r = ((QLabel__Go*)self)->cppgoBase_focusNextPrevChild(next);
	return __r;
}
void QLabel_initStyleOption_default(void* self, void* option) {
	((QLabel__Go*)self)->cppgoBase_initStyleOption((QStyleOptionFrame*)option);
}
int QLabel_devType_default(void* self) {
	auto&& __r = ((QLabel*)self)->QLabel::devType();
	return __r;
}
void QLabel_setVisible_default(void* self, int visible) {
	((QLabel*)self)->QLabel::setVisible(visible);
}
int QLabel_hasHeightForWidth_default(void* self) {
	auto&& __r = ((QLabel*)self)->QLabel::hasHeightForWidth();
	return __r;
}
void* QLabel_paintEngine_default(void* self) {
	auto&& __r = ((QLabel*)self)->QLabel::paintEngine();
	return (void*)(__r);
}
void QLabel_mouseDoubleClickEvent_default(void* self, void* event) {
	((QLabel__Go*)self)->cppgoBase_mouseDoubleClickEvent((QMouseEvent*)event);
}
void QLabel_wheelEvent_default(void* self, void* event) {
	((QLabel__Go*)self)->cppgoBase_wheelEvent((QWheelEvent*)event);
}
void QLabel_keyReleaseEvent_default(void* self, void* event) {
	((QLabel__Go*)self)->cppgoBase_keyReleaseEvent((QKeyEvent*)event);
}
void QLabel_enterEvent_default(void* self, void* event) {
	((QLabel__Go*)self)->cppgoBase_enterEvent((QEnterEvent*)event);
}
void QLabel_leaveEvent_default(void* self, void* event) {
	((QLabel__Go*)self)->cppgoBase_leaveEvent((QEvent*)event);
}
void QLabel_moveEvent_default(void* self, void* event) {
	((QLabel__Go*)self)->cppgoBase_moveEvent((QMoveEvent*)event);
}
void QLabel_resizeEvent_default(void* self, void* event) {
	((QLabel__Go*)self)->cppgoBase_resizeEvent((QResizeEvent*)event);
}
void QLabel_closeEvent_default(void* self, void* event) {
	((QLabel__Go*)self)->cppgoBase_closeEvent((QCloseEvent*)event);
}
void QLabel_tabletEvent_default(void* self, void* event) {
	((QLabel__Go*)self)->cppgoBase_tabletEvent((QTabletEvent*)event);
}
void QLabel_actionEvent_default(void* self, void* event) {
	((QLabel__Go*)self)->cppgoBase_actionEvent((QActionEvent*)event);
}
void QLabel_dragEnterEvent_default(void* self, void* event) {
	((QLabel__Go*)self)->cppgoBase_dragEnterEvent((QDragEnterEvent*)event);
}
void QLabel_dragMoveEvent_default(void* self, void* event) {
	((QLabel__Go*)self)->cppgoBase_dragMoveEvent((QDragMoveEvent*)event);
}
void QLabel_dragLeaveEvent_default(void* self, void* event) {
	((QLabel__Go*)self)->cppgoBase_dragLeaveEvent((QDragLeaveEvent*)event);
}
void QLabel_dropEvent_default(void* self, void* event) {
	((QLabel__Go*)self)->cppgoBase_dropEvent((QDropEvent*)event);
}
void QLabel_showEvent_default(void* self, void* event) {
	((QLabel__Go*)self)->cppgoBase_showEvent((QShowEvent*)event);
}
void QLabel_hideEvent_default(void* self, void* event) {
	((QLabel__Go*)self)->cppgoBase_hideEvent((QHideEvent*)event);
}
int QLabel_metric_default(void* self, int arg1) {
	auto&& __r = ((QLabel__Go*)self)->cppgoBase_metric(arg1);
	return __r;
}
void QLabel_initPainter_default(void* self, void* painter) {
	((QLabel__Go*)self)->cppgoBase_initPainter((QPainter*)painter);
}
void* QLabel_redirected_default(void* self, void* offset) {
	auto&& __r = ((QLabel__Go*)self)->cppgoBase_redirected((QPoint*)offset);
	return (void*)(__r);
}
void* QLabel_sharedPainter_default(void* self) {
	auto&& __r = ((QLabel__Go*)self)->cppgoBase_sharedPainter();
	return (void*)(__r);
}
void QLabel_inputMethodEvent_default(void* self, void* arg1) {
	((QLabel__Go*)self)->cppgoBase_inputMethodEvent((QInputMethodEvent*)arg1);
}
void* QLabel_inputMethodQuery_default(void* self, int arg1) {
	auto&& __r = ((QLabel*)self)->QLabel::inputMethodQuery((Qt::InputMethodQuery)(arg1));
	return new QVariant(__r);
}
int QLabel_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QLabel*)self)->QLabel::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QLabel_timerEvent_default(void* self, void* event) {
	((QLabel__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QLabel_childEvent_default(void* self, void* event) {
	((QLabel__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QLabel_customEvent_default(void* self, void* event) {
	((QLabel__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QLabel_connectNotify_default(void* self, void* signal) {
	((QLabel__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QLabel_disconnectNotify_default(void* self, void* signal) {
	((QLabel__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QLabel_connect_linkActivated(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QLabel*)self, &QLabel::linkActivated, (QLabel*)self, [__id](const QString& link){
		QByteArray link_ba = link.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)link_ba.constData(), link_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QLabel);
	__conns_QLabel[__id] = __c;
}
void QLabel_connect_linkHovered(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QLabel*)self, &QLabel::linkHovered, (QLabel*)self, [__id](const QString& link){
		QByteArray link_ba = link.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)link_ba.constData(), link_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QLabel);
	__conns_QLabel[__id] = __c;
}
void QLabel_connect_windowTitleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QLabel*)self, &QLabel::windowTitleChanged, (QLabel*)self, [__id](const QString& title){
		QByteArray title_ba = title.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)title_ba.constData(), title_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QLabel);
	__conns_QLabel[__id] = __c;
}
void QLabel_connect_windowIconChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QLabel*)self, &QLabel::windowIconChanged, (QLabel*)self, [__id](const QIcon& icon){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&icon);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QLabel);
	__conns_QLabel[__id] = __c;
}
void QLabel_connect_windowIconTextChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QLabel*)self, &QLabel::windowIconTextChanged, (QLabel*)self, [__id](const QString& iconText){
		QByteArray iconText_ba = iconText.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)iconText_ba.constData(), iconText_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QLabel);
	__conns_QLabel[__id] = __c;
}
void QLabel_connect_customContextMenuRequested(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QLabel*)self, &QLabel::customContextMenuRequested, (QLabel*)self, [__id](const QPoint& pos){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&pos);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QLabel);
	__conns_QLabel[__id] = __c;
}
void QLabel_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QLabel*)self, &QLabel::destroyed, (QLabel*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QLabel);
	__conns_QLabel[__id] = __c;
}
void QLabel_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QLabel*)self, &QLabel::destroyed, (QLabel*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QLabel);
	__conns_QLabel[__id] = __c;
}
void QLabel_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QLabel*)self, &QLabel::objectNameChanged, (QLabel*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QLabel);
	__conns_QLabel[__id] = __c;
}
void QLabel__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QLabel);
		auto __it = __conns_QLabel.find(id);
		if (__it == __conns_QLabel.end()) return;
		__c = __it->second;
		__conns_QLabel.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QLabel__Go_new(long long id) { return new QLabel__Go(id); }
void QLabel_delete(void* self) { delete (QLabel*)self; }

}

#include <QtWidgets/qspinbox.h>
#include <QtWidgets/qabstractspinbox.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qcoreevent.h>
#include <QtGui/qvalidator.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qsize.h>
#include <QtCore/qvariant.h>
#include <QtGui/qevent.h>
#include <QtWidgets/qstyleoption.h>
#include <QtWidgets/qlineedit.h>
#include <QtWidgets/qstyle.h>
#include <QtCore/qrect.h>
#include <QtCore/qpoint.h>
#include <QtGui/qregion.h>
#include <QtGui/qpalette.h>
#include <QtGui/qfont.h>
#include <QtGui/qfontmetrics.h>
#include <QtGui/qfontinfo.h>
#include <QtGui/qcursor.h>
#include <QtGui/qbitmap.h>
#include <QtGui/qpaintdevice.h>
#include <QtGui/qpainter.h>
#include <QtGui/qpixmap.h>
#include <QtWidgets/qgraphicseffect.h>
#include <QtGui/qicon.h>
#include <QtCore/qlocale.h>
#include <QtGui/qkeysequence.h>
#include <QtWidgets/qgraphicsproxywidget.h>
#include <QtCore/qbytearray.h>
#include <QtWidgets/qsizepolicy.h>
#include <QtCore/qmargins.h>
#include <QtWidgets/qlayout.h>
#include <QtGui/qaction.h>
#include <QtCore/qobject.h>
#include <QtGui/qpaintengine.h>
#include <QtGui/qbackingstore.h>
#include <QtGui/qwindow.h>
#include <QtGui/qscreen.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qspinbox.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QSpinBox__Go : public QSpinBox {
public:
	long long __goId;
	unsigned char __goOn[56] = {};
	template<typename... A> QSpinBox__Go(long long id, A&&... a) : QSpinBox(std::forward<A>(a)...), __goId(id) {}
	using QSpinBox::lineEdit;
	using QSpinBox::setLineEdit;
	using QSpinBox::updateMicroFocus;
	using QSpinBox::create;
	using QSpinBox::destroy;
	using QSpinBox::focusNextChild;
	using QSpinBox::focusPreviousChild;
	using QSpinBox::sender;
	using QSpinBox::senderSignalIndex;
	using QSpinBox::receivers;
	using QSpinBox::isSignalConnected;
	using QSpinBox::getDecodedMetricF;
	bool cppgoBase_event(QEvent* __a0) { return QSpinBox::event(__a0); }
	QValidator::State cppgoBase_validate(QString& __a0, int& __a1) { return QSpinBox::validate(__a0, __a1); }
	int cppgoBase_valueFromText(const QString& __a0) { return QSpinBox::valueFromText(__a0); }
	QString cppgoBase_textFromValue(int __a0) { return QSpinBox::textFromValue(__a0); }
	void cppgoBase_fixup(QString& __a0) { QSpinBox::fixup(__a0); }
	void cppgoBase_resizeEvent(QResizeEvent* __a0) { QSpinBox::resizeEvent(__a0); }
	void cppgoBase_keyPressEvent(QKeyEvent* __a0) { QSpinBox::keyPressEvent(__a0); }
	void cppgoBase_keyReleaseEvent(QKeyEvent* __a0) { QSpinBox::keyReleaseEvent(__a0); }
	void cppgoBase_wheelEvent(QWheelEvent* __a0) { QSpinBox::wheelEvent(__a0); }
	void cppgoBase_focusInEvent(QFocusEvent* __a0) { QSpinBox::focusInEvent(__a0); }
	void cppgoBase_focusOutEvent(QFocusEvent* __a0) { QSpinBox::focusOutEvent(__a0); }
	void cppgoBase_contextMenuEvent(QContextMenuEvent* __a0) { QSpinBox::contextMenuEvent(__a0); }
	void cppgoBase_changeEvent(QEvent* __a0) { QSpinBox::changeEvent(__a0); }
	void cppgoBase_closeEvent(QCloseEvent* __a0) { QSpinBox::closeEvent(__a0); }
	void cppgoBase_hideEvent(QHideEvent* __a0) { QSpinBox::hideEvent(__a0); }
	void cppgoBase_mousePressEvent(QMouseEvent* __a0) { QSpinBox::mousePressEvent(__a0); }
	void cppgoBase_mouseReleaseEvent(QMouseEvent* __a0) { QSpinBox::mouseReleaseEvent(__a0); }
	void cppgoBase_mouseMoveEvent(QMouseEvent* __a0) { QSpinBox::mouseMoveEvent(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QSpinBox::timerEvent(__a0); }
	void cppgoBase_paintEvent(QPaintEvent* __a0) { QSpinBox::paintEvent(__a0); }
	void cppgoBase_showEvent(QShowEvent* __a0) { QSpinBox::showEvent(__a0); }
	void cppgoBase_initStyleOption(QStyleOptionSpinBox* __a0) { QSpinBox::initStyleOption(__a0); }
	QFlags<QAbstractSpinBox::StepEnabledFlag> cppgoBase_stepEnabled() { return QSpinBox::stepEnabled(); }
	void cppgoBase_mouseDoubleClickEvent(QMouseEvent* __a0) { QSpinBox::mouseDoubleClickEvent(__a0); }
	void cppgoBase_enterEvent(QEnterEvent* __a0) { QSpinBox::enterEvent(__a0); }
	void cppgoBase_leaveEvent(QEvent* __a0) { QSpinBox::leaveEvent(__a0); }
	void cppgoBase_moveEvent(QMoveEvent* __a0) { QSpinBox::moveEvent(__a0); }
	void cppgoBase_tabletEvent(QTabletEvent* __a0) { QSpinBox::tabletEvent(__a0); }
	void cppgoBase_actionEvent(QActionEvent* __a0) { QSpinBox::actionEvent(__a0); }
	void cppgoBase_dragEnterEvent(QDragEnterEvent* __a0) { QSpinBox::dragEnterEvent(__a0); }
	void cppgoBase_dragMoveEvent(QDragMoveEvent* __a0) { QSpinBox::dragMoveEvent(__a0); }
	void cppgoBase_dragLeaveEvent(QDragLeaveEvent* __a0) { QSpinBox::dragLeaveEvent(__a0); }
	void cppgoBase_dropEvent(QDropEvent* __a0) { QSpinBox::dropEvent(__a0); }
	int cppgoBase_metric(int __a0) { return QSpinBox::metric((QPaintDevice::PaintDeviceMetric)__a0); }
	void cppgoBase_initPainter(QPainter* __a0) { QSpinBox::initPainter(__a0); }
	QPaintDevice* cppgoBase_redirected(QPoint* __a0) { return QSpinBox::redirected(__a0); }
	QPainter* cppgoBase_sharedPainter() { return QSpinBox::sharedPainter(); }
	void cppgoBase_inputMethodEvent(QInputMethodEvent* __a0) { QSpinBox::inputMethodEvent(__a0); }
	bool cppgoBase_focusNextPrevChild(bool __a0) { return QSpinBox::focusNextPrevChild(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QSpinBox::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QSpinBox::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QSpinBox::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QSpinBox::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QSpinBox::metaObject();
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QSpinBox::event(event);
	}
	QValidator::State validate(QString& input, int& pos) const override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_int_voidp_intp) { int __r = ((int(*)(long long, long long, void*, int*, int*))cppgo_virt_int_voidp_intp)(__goId, 6784179835071232555LL, (void*)&input, (int*)&pos, &__h); if (__h) return (QValidator::State)(__r); }
		return QSpinBox::validate(input, pos);
	}
	int valueFromText(const QString& text) const override {
		int __h = 0;
		QByteArray text_ba = text.toUtf8();
		if (__goOn[3] && cppgo_virt_int_constcharp_longlong) { int __r = ((int(*)(long long, long long, const char*, long long, int*))cppgo_virt_int_constcharp_longlong)(__goId, 5050482105304000557LL, (char*)text_ba.constData(), text_ba.size(), &__h); if (__h) return __r; }
		return QSpinBox::valueFromText(text);
	}
	QString textFromValue(int val) const override {
		int __h = 0;
		if (__goOn[4] && cppgo_virtp_charp_int) {
			long long __len = 0;
			char* __s = ((char*(*)(long long, long long, int, long long*, int*))cppgo_virtp_charp_int)(__goId, -9083693956578760701LL, val, &__len, &__h);
			if (__h) { QString __r = QString::fromUtf8(__s, __len); free(__s); return __r; }
		}
		return QSpinBox::textFromValue(val);
	}
	void fixup(QString& str) const override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4855206404746966715LL, (void*)&str, &__h); if (__h) return; }
		QSpinBox::fixup(str);
	}
	QSize sizeHint() const override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 729690261457684101LL, &__h); if (__h) return *(QSize*)__r; }
		return QSpinBox::sizeHint();
	}
	QSize minimumSizeHint() const override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 7049264599837532491LL, &__h); if (__h) return *(QSize*)__r; }
		return QSpinBox::minimumSizeHint();
	}
	QVariant inputMethodQuery(Qt::InputMethodQuery arg1) const override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7345853097340709638LL, (int)(arg1), &__h); if (__h) return *(QVariant*)__r; }
		return QSpinBox::inputMethodQuery(arg1);
	}
	void stepBy(int steps) override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 7085718643589954634LL, steps, &__h); if (__h) return; }
		QSpinBox::stepBy(steps);
	}
	void clear() override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, -778567940572639070LL, &__h); if (__h) return; }
		QSpinBox::clear();
	}
	void resizeEvent(QResizeEvent* event) override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4302917285655014661LL, (void*)event, &__h); if (__h) return; }
		QSpinBox::resizeEvent(event);
	}
	void keyPressEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1023791977679793185LL, (void*)event, &__h); if (__h) return; }
		QSpinBox::keyPressEvent(event);
	}
	void keyReleaseEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 772598475950050299LL, (void*)event, &__h); if (__h) return; }
		QSpinBox::keyReleaseEvent(event);
	}
	void wheelEvent(QWheelEvent* event) override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1919988340934511316LL, (void*)event, &__h); if (__h) return; }
		QSpinBox::wheelEvent(event);
	}
	void focusInEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 975043767033552552LL, (void*)event, &__h); if (__h) return; }
		QSpinBox::focusInEvent(event);
	}
	void focusOutEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -385596010576048993LL, (void*)event, &__h); if (__h) return; }
		QSpinBox::focusOutEvent(event);
	}
	void contextMenuEvent(QContextMenuEvent* event) override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6698043222227006015LL, (void*)event, &__h); if (__h) return; }
		QSpinBox::contextMenuEvent(event);
	}
	void changeEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4959376063826165943LL, (void*)event, &__h); if (__h) return; }
		QSpinBox::changeEvent(event);
	}
	void closeEvent(QCloseEvent* event) override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1989937050019464783LL, (void*)event, &__h); if (__h) return; }
		QSpinBox::closeEvent(event);
	}
	void hideEvent(QHideEvent* event) override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -199718707876719129LL, (void*)event, &__h); if (__h) return; }
		QSpinBox::hideEvent(event);
	}
	void mousePressEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -332090898078543319LL, (void*)event, &__h); if (__h) return; }
		QSpinBox::mousePressEvent(event);
	}
	void mouseReleaseEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4415920049297795213LL, (void*)event, &__h); if (__h) return; }
		QSpinBox::mouseReleaseEvent(event);
	}
	void mouseMoveEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4255494004099806581LL, (void*)event, &__h); if (__h) return; }
		QSpinBox::mouseMoveEvent(event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QSpinBox::timerEvent(event);
	}
	void paintEvent(QPaintEvent* event) override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5290442163167458465LL, (void*)event, &__h); if (__h) return; }
		QSpinBox::paintEvent(event);
	}
	void showEvent(QShowEvent* event) override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6825369174751916LL, (void*)event, &__h); if (__h) return; }
		QSpinBox::showEvent(event);
	}
	void initStyleOption(QStyleOptionSpinBox* option) const override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -488693769771308531LL, (void*)option, &__h); if (__h) return; }
		QSpinBox::initStyleOption(option);
	}
	QFlags<QAbstractSpinBox::StepEnabledFlag> stepEnabled() const override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, 5257289574516412244LL, &__h); if (__h) return QFlags<QAbstractSpinBox::StepEnabledFlag>(__r); }
		return QSpinBox::stepEnabled();
	}
	int devType() const override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8659679893868524044LL, &__h); if (__h) return __r; }
		return QSpinBox::devType();
	}
	void setVisible(bool visible) override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 2110182719375245575LL, visible, &__h); if (__h) return; }
		QSpinBox::setVisible(visible);
	}
	int heightForWidth(int arg1) const override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 4311534060495737339LL, arg1, &__h); if (__h) return __r; }
		return QSpinBox::heightForWidth(arg1);
	}
	bool hasHeightForWidth() const override {
		int __h = 0;
		if (__goOn[32] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8266762397995423915LL, &__h); if (__h) return __r; }
		return QSpinBox::hasHeightForWidth();
	}
	QPaintEngine* paintEngine() const override {
		int __h = 0;
		if (__goOn[33] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 6398007404756346381LL, &__h); if (__h) return (QPaintEngine*)__r; }
		return QSpinBox::paintEngine();
	}
	void mouseDoubleClickEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[34] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4029073905697023887LL, (void*)event, &__h); if (__h) return; }
		QSpinBox::mouseDoubleClickEvent(event);
	}
	void enterEvent(QEnterEvent* event) override {
		int __h = 0;
		if (__goOn[35] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6316270346668237895LL, (void*)event, &__h); if (__h) return; }
		QSpinBox::enterEvent(event);
	}
	void leaveEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[36] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -551070009623711640LL, (void*)event, &__h); if (__h) return; }
		QSpinBox::leaveEvent(event);
	}
	void moveEvent(QMoveEvent* event) override {
		int __h = 0;
		if (__goOn[37] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -9161949150489525396LL, (void*)event, &__h); if (__h) return; }
		QSpinBox::moveEvent(event);
	}
	void tabletEvent(QTabletEvent* event) override {
		int __h = 0;
		if (__goOn[38] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -212068866860701141LL, (void*)event, &__h); if (__h) return; }
		QSpinBox::tabletEvent(event);
	}
	void actionEvent(QActionEvent* event) override {
		int __h = 0;
		if (__goOn[39] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1697444140088555581LL, (void*)event, &__h); if (__h) return; }
		QSpinBox::actionEvent(event);
	}
	void dragEnterEvent(QDragEnterEvent* event) override {
		int __h = 0;
		if (__goOn[40] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 206122715919917589LL, (void*)event, &__h); if (__h) return; }
		QSpinBox::dragEnterEvent(event);
	}
	void dragMoveEvent(QDragMoveEvent* event) override {
		int __h = 0;
		if (__goOn[41] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1781545137352745978LL, (void*)event, &__h); if (__h) return; }
		QSpinBox::dragMoveEvent(event);
	}
	void dragLeaveEvent(QDragLeaveEvent* event) override {
		int __h = 0;
		if (__goOn[42] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -3139904158373725006LL, (void*)event, &__h); if (__h) return; }
		QSpinBox::dragLeaveEvent(event);
	}
	void dropEvent(QDropEvent* event) override {
		int __h = 0;
		if (__goOn[43] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 9107058682754591716LL, (void*)event, &__h); if (__h) return; }
		QSpinBox::dropEvent(event);
	}
	bool nativeEvent(const QByteArray& eventType, void* message, qintptr* result) override {
		int __h = 0;
		if (__goOn[44] && cppgo_virt_int_voidp_voidp_longlongp) { int __r = ((int(*)(long long, long long, void*, void*, long long*, int*))cppgo_virt_int_voidp_voidp_longlongp)(__goId, -2039166586246443706LL, (void*)&eventType, (void*)message, (long long*)result, &__h); if (__h) return __r; }
		return QSpinBox::nativeEvent(eventType, message, result);
	}
	int metric(QPaintDevice::PaintDeviceMetric arg1) const override {
		int __h = 0;
		if (__goOn[45] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 427629649267311777LL, (int)(arg1), &__h); if (__h) return __r; }
		return QSpinBox::metric(arg1);
	}
	void initPainter(QPainter* painter) const override {
		int __h = 0;
		if (__goOn[46] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6477500296980050210LL, (void*)painter, &__h); if (__h) return; }
		QSpinBox::initPainter(painter);
	}
	QPaintDevice* redirected(QPoint* offset) const override {
		int __h = 0;
		if (__goOn[47] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 4127240633310655048LL, (void*)offset, &__h); if (__h) return (QPaintDevice*)__r; }
		return QSpinBox::redirected(offset);
	}
	QPainter* sharedPainter() const override {
		int __h = 0;
		if (__goOn[48] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 5792771328741774991LL, &__h); if (__h) return (QPainter*)__r; }
		return QSpinBox::sharedPainter();
	}
	void inputMethodEvent(QInputMethodEvent* arg1) override {
		int __h = 0;
		if (__goOn[49] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5438819661942843222LL, (void*)arg1, &__h); if (__h) return; }
		QSpinBox::inputMethodEvent(arg1);
	}
	bool focusNextPrevChild(bool next) override {
		int __h = 0;
		if (__goOn[50] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 5898328517505784191LL, next, &__h); if (__h) return __r; }
		return QSpinBox::focusNextPrevChild(next);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[51] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QSpinBox::eventFilter(watched, event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[52] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QSpinBox::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[53] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QSpinBox::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[54] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QSpinBox::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[55] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QSpinBox::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QSpinBox;
static std::mutex __connsMu_QSpinBox;

extern "C" {

void* QSpinBox_base(void* self) { return (void*)static_cast<QAbstractSpinBox*>((QSpinBox*)self); }
void QSpinBox_chain(void* self, void** out) {
	QSpinBox* __p = (QSpinBox*)self;
	QAbstractSpinBox* __b0 = static_cast<QAbstractSpinBox*>(__p); out[0] = (void*)__b0;
	QWidget* __b1 = static_cast<QWidget*>(__b0); out[1] = (void*)__b1;
	QObject* __b2 = static_cast<QObject*>(__b1); out[2] = (void*)__b2;
}
void QSpinBox__set_override(void* self, int idx) { ((QSpinBox__Go*)self)->__goOn[idx] = 1; }
void* QSpinBox_fromQObject(void* o) { return (void*)qobject_cast<QSpinBox*>((QObject*)o); }
void* QSpinBox_new(long long __goid, void* parent) {
	return new QSpinBox__Go(__goid, (QWidget*)parent);
}
void* QSpinBox_new2(long long __goid) {
	return new QSpinBox__Go(__goid);
}
char* QSpinBox_tr(const char* s, long long* __out_len) {
	auto&& __r = QSpinBox::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QSpinBox_value(void* self) {
	auto&& __r = ((QSpinBox*)self)->value();
	return __r;
}
char* QSpinBox_prefix(void* self, long long* __out_len) {
	auto&& __r = ((QSpinBox*)self)->prefix();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QSpinBox_setPrefix(void* self, const char* prefix, long long prefix_len) {
	QString prefix_qstr = QString::fromUtf8(prefix, prefix_len);
	((QSpinBox*)self)->setPrefix(prefix_qstr);
}
char* QSpinBox_suffix(void* self, long long* __out_len) {
	auto&& __r = ((QSpinBox*)self)->suffix();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QSpinBox_setSuffix(void* self, const char* suffix, long long suffix_len) {
	QString suffix_qstr = QString::fromUtf8(suffix, suffix_len);
	((QSpinBox*)self)->setSuffix(suffix_qstr);
}
char* QSpinBox_cleanText(void* self, long long* __out_len) {
	auto&& __r = ((QSpinBox*)self)->cleanText();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QSpinBox_singleStep(void* self) {
	auto&& __r = ((QSpinBox*)self)->singleStep();
	return __r;
}
void QSpinBox_setSingleStep(void* self, int val) {
	((QSpinBox*)self)->setSingleStep(val);
}
int QSpinBox_minimum(void* self) {
	auto&& __r = ((QSpinBox*)self)->minimum();
	return __r;
}
void QSpinBox_setMinimum(void* self, int min) {
	((QSpinBox*)self)->setMinimum(min);
}
int QSpinBox_maximum(void* self) {
	auto&& __r = ((QSpinBox*)self)->maximum();
	return __r;
}
void QSpinBox_setMaximum(void* self, int max) {
	((QSpinBox*)self)->setMaximum(max);
}
void QSpinBox_setRange(void* self, int min, int max) {
	((QSpinBox*)self)->setRange(min, max);
}
int QSpinBox_stepType(void* self) {
	auto&& __r = ((QSpinBox*)self)->stepType();
	return (int)(__r);
}
void QSpinBox_setStepType(void* self, int stepType) {
	((QSpinBox*)self)->setStepType((QAbstractSpinBox::StepType)(stepType));
}
int QSpinBox_displayIntegerBase(void* self) {
	auto&& __r = ((QSpinBox*)self)->displayIntegerBase();
	return __r;
}
void QSpinBox_setDisplayIntegerBase(void* self, int base) {
	((QSpinBox*)self)->setDisplayIntegerBase(base);
}
void QSpinBox_setValue(void* self, int val) {
	((QSpinBox*)self)->setValue(val);
}
void QSpinBox_valueChanged(void* self, int arg1) {
	((QSpinBox*)self)->valueChanged(arg1);
}
void QSpinBox_textChanged(void* self, const char* arg1, long long arg1_len) {
	QString arg1_qstr = QString::fromUtf8(arg1, arg1_len);
	((QSpinBox*)self)->textChanged(arg1_qstr);
}
void* QSpinBox_lineEdit(void* self) {
	auto&& __r = ((QSpinBox__Go*)self)->lineEdit();
	return (void*)(__r);
}
void QSpinBox_setLineEdit(void* self, void* edit) {
	((QSpinBox__Go*)self)->setLineEdit((QLineEdit*)edit);
}
void QSpinBox_setTabOrder(void* arg1, void* arg2) {
	QSpinBox::setTabOrder((QWidget*)arg1, (QWidget*)arg2);
}
void* QSpinBox_mouseGrabber() {
	auto&& __r = QSpinBox::mouseGrabber();
	return (void*)(__r);
}
void* QSpinBox_keyboardGrabber() {
	auto&& __r = QSpinBox::keyboardGrabber();
	return (void*)(__r);
}
void* QSpinBox_find(unsigned long long arg1) {
	auto&& __r = QSpinBox::find(arg1);
	return (void*)(__r);
}
void* QSpinBox_createWindowContainer(void* window) {
	auto&& __r = QSpinBox::createWindowContainer((QWindow*)window);
	return (void*)(__r);
}
void QSpinBox_updateMicroFocus(void* self) {
	((QSpinBox__Go*)self)->updateMicroFocus();
}
void QSpinBox_create(void* self) {
	((QSpinBox__Go*)self)->create();
}
void QSpinBox_destroy(void* self) {
	((QSpinBox__Go*)self)->destroy();
}
int QSpinBox_focusNextChild(void* self) {
	auto&& __r = ((QSpinBox__Go*)self)->focusNextChild();
	return __r;
}
int QSpinBox_focusPreviousChild(void* self) {
	auto&& __r = ((QSpinBox__Go*)self)->focusPreviousChild();
	return __r;
}
void* QSpinBox_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QSpinBox::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QSpinBox_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QSpinBox::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QSpinBox_sender(void* self) {
	auto&& __r = ((QSpinBox__Go*)self)->sender();
	return (void*)(__r);
}
int QSpinBox_senderSignalIndex(void* self) {
	auto&& __r = ((QSpinBox__Go*)self)->senderSignalIndex();
	return __r;
}
int QSpinBox_receivers(void* self, const char* signal) {
	auto&& __r = ((QSpinBox__Go*)self)->receivers(signal);
	return __r;
}
int QSpinBox_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QSpinBox__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
double QSpinBox_devicePixelRatioFScale() {
	auto&& __r = QSpinBox::devicePixelRatioFScale();
	return __r;
}
int QSpinBox_encodeMetricF(int metric, double value) {
	auto&& __r = QSpinBox::encodeMetricF((QPaintDevice::PaintDeviceMetric)(metric), value);
	return __r;
}
double QSpinBox_getDecodedMetricF(void* self, int metricA, int metricB) {
	auto&& __r = ((QSpinBox__Go*)self)->getDecodedMetricF((QPaintDevice::PaintDeviceMetric)(metricA), (QPaintDevice::PaintDeviceMetric)(metricB));
	return __r;
}
char* QSpinBox_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QSpinBox::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QSpinBox_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QSpinBox::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QSpinBox_createWindowContainer2(void* window, void* parent) {
	auto&& __r = QSpinBox::createWindowContainer((QWindow*)window, (QWidget*)parent);
	return (void*)(__r);
}
void* QSpinBox_createWindowContainer3(void* window, void* parent, int flags) {
	auto&& __r = QSpinBox::createWindowContainer((QWindow*)window, (QWidget*)parent, QFlags<Qt::WindowType>(flags));
	return (void*)(__r);
}
void QSpinBox_updateMicroFocus2(void* self, int query) {
	((QSpinBox__Go*)self)->updateMicroFocus((Qt::InputMethodQuery)(query));
}
void QSpinBox_create2(void* self, unsigned long long arg1) {
	((QSpinBox__Go*)self)->create(arg1);
}
void QSpinBox_create3(void* self, unsigned long long arg1, int initializeWindow) {
	((QSpinBox__Go*)self)->create(arg1, initializeWindow);
}
void QSpinBox_create4(void* self, unsigned long long arg1, int initializeWindow, int destroyOldWindow) {
	((QSpinBox__Go*)self)->create(arg1, initializeWindow, destroyOldWindow);
}
void QSpinBox_destroy2(void* self, int destroyWindow) {
	((QSpinBox__Go*)self)->destroy(destroyWindow);
}
void QSpinBox_destroy3(void* self, int destroyWindow, int destroySubWindows) {
	((QSpinBox__Go*)self)->destroy(destroyWindow, destroySubWindows);
}
void* QSpinBox_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QSpinBox::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QSpinBox_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QSpinBox::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QSpinBox_metaObject_default(void* self) {
	auto&& __r = ((QSpinBox*)self)->QSpinBox::metaObject();
	return (void*)(__r);
}
int QSpinBox_event_default(void* self, void* event) {
	auto&& __r = ((QSpinBox__Go*)self)->cppgoBase_event((QEvent*)event);
	return __r;
}
int QSpinBox_validate_default(void* self, void* input, int* pos) {
	int pos_ref = *pos;
	auto&& __r = ((QSpinBox__Go*)self)->cppgoBase_validate(*(QString*)input, pos_ref);
	*pos = pos_ref;
	return (int)(__r);
}
int QSpinBox_valueFromText_default(void* self, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = ((QSpinBox__Go*)self)->cppgoBase_valueFromText(text_qstr);
	return __r;
}
char* QSpinBox_textFromValue_default(void* self, int val, long long* __out_len) {
	auto&& __r = ((QSpinBox__Go*)self)->cppgoBase_textFromValue(val);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QSpinBox_fixup_default(void* self, void* str) {
	((QSpinBox__Go*)self)->cppgoBase_fixup(*(QString*)str);
}
void* QSpinBox_sizeHint_default(void* self) {
	auto&& __r = ((QSpinBox*)self)->QSpinBox::sizeHint();
	return new QSize(__r);
}
void* QSpinBox_minimumSizeHint_default(void* self) {
	auto&& __r = ((QSpinBox*)self)->QSpinBox::minimumSizeHint();
	return new QSize(__r);
}
void* QSpinBox_inputMethodQuery_default(void* self, int arg1) {
	auto&& __r = ((QSpinBox*)self)->QSpinBox::inputMethodQuery((Qt::InputMethodQuery)(arg1));
	return new QVariant(__r);
}
void QSpinBox_stepBy_default(void* self, int steps) {
	((QSpinBox*)self)->QSpinBox::stepBy(steps);
}
void QSpinBox_clear_default(void* self) {
	((QSpinBox*)self)->QSpinBox::clear();
}
void QSpinBox_resizeEvent_default(void* self, void* event) {
	((QSpinBox__Go*)self)->cppgoBase_resizeEvent((QResizeEvent*)event);
}
void QSpinBox_keyPressEvent_default(void* self, void* event) {
	((QSpinBox__Go*)self)->cppgoBase_keyPressEvent((QKeyEvent*)event);
}
void QSpinBox_keyReleaseEvent_default(void* self, void* event) {
	((QSpinBox__Go*)self)->cppgoBase_keyReleaseEvent((QKeyEvent*)event);
}
void QSpinBox_wheelEvent_default(void* self, void* event) {
	((QSpinBox__Go*)self)->cppgoBase_wheelEvent((QWheelEvent*)event);
}
void QSpinBox_focusInEvent_default(void* self, void* event) {
	((QSpinBox__Go*)self)->cppgoBase_focusInEvent((QFocusEvent*)event);
}
void QSpinBox_focusOutEvent_default(void* self, void* event) {
	((QSpinBox__Go*)self)->cppgoBase_focusOutEvent((QFocusEvent*)event);
}
void QSpinBox_contextMenuEvent_default(void* self, void* event) {
	((QSpinBox__Go*)self)->cppgoBase_contextMenuEvent((QContextMenuEvent*)event);
}
void QSpinBox_changeEvent_default(void* self, void* event) {
	((QSpinBox__Go*)self)->cppgoBase_changeEvent((QEvent*)event);
}
void QSpinBox_closeEvent_default(void* self, void* event) {
	((QSpinBox__Go*)self)->cppgoBase_closeEvent((QCloseEvent*)event);
}
void QSpinBox_hideEvent_default(void* self, void* event) {
	((QSpinBox__Go*)self)->cppgoBase_hideEvent((QHideEvent*)event);
}
void QSpinBox_mousePressEvent_default(void* self, void* event) {
	((QSpinBox__Go*)self)->cppgoBase_mousePressEvent((QMouseEvent*)event);
}
void QSpinBox_mouseReleaseEvent_default(void* self, void* event) {
	((QSpinBox__Go*)self)->cppgoBase_mouseReleaseEvent((QMouseEvent*)event);
}
void QSpinBox_mouseMoveEvent_default(void* self, void* event) {
	((QSpinBox__Go*)self)->cppgoBase_mouseMoveEvent((QMouseEvent*)event);
}
void QSpinBox_timerEvent_default(void* self, void* event) {
	((QSpinBox__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QSpinBox_paintEvent_default(void* self, void* event) {
	((QSpinBox__Go*)self)->cppgoBase_paintEvent((QPaintEvent*)event);
}
void QSpinBox_showEvent_default(void* self, void* event) {
	((QSpinBox__Go*)self)->cppgoBase_showEvent((QShowEvent*)event);
}
void QSpinBox_initStyleOption_default(void* self, void* option) {
	((QSpinBox__Go*)self)->cppgoBase_initStyleOption((QStyleOptionSpinBox*)option);
}
int QSpinBox_stepEnabled_default(void* self) {
	auto&& __r = ((QSpinBox__Go*)self)->cppgoBase_stepEnabled();
	return (int)(__r);
}
int QSpinBox_devType_default(void* self) {
	auto&& __r = ((QSpinBox*)self)->QSpinBox::devType();
	return __r;
}
void QSpinBox_setVisible_default(void* self, int visible) {
	((QSpinBox*)self)->QSpinBox::setVisible(visible);
}
int QSpinBox_heightForWidth_default(void* self, int arg1) {
	auto&& __r = ((QSpinBox*)self)->QSpinBox::heightForWidth(arg1);
	return __r;
}
int QSpinBox_hasHeightForWidth_default(void* self) {
	auto&& __r = ((QSpinBox*)self)->QSpinBox::hasHeightForWidth();
	return __r;
}
void* QSpinBox_paintEngine_default(void* self) {
	auto&& __r = ((QSpinBox*)self)->QSpinBox::paintEngine();
	return (void*)(__r);
}
void QSpinBox_mouseDoubleClickEvent_default(void* self, void* event) {
	((QSpinBox__Go*)self)->cppgoBase_mouseDoubleClickEvent((QMouseEvent*)event);
}
void QSpinBox_enterEvent_default(void* self, void* event) {
	((QSpinBox__Go*)self)->cppgoBase_enterEvent((QEnterEvent*)event);
}
void QSpinBox_leaveEvent_default(void* self, void* event) {
	((QSpinBox__Go*)self)->cppgoBase_leaveEvent((QEvent*)event);
}
void QSpinBox_moveEvent_default(void* self, void* event) {
	((QSpinBox__Go*)self)->cppgoBase_moveEvent((QMoveEvent*)event);
}
void QSpinBox_tabletEvent_default(void* self, void* event) {
	((QSpinBox__Go*)self)->cppgoBase_tabletEvent((QTabletEvent*)event);
}
void QSpinBox_actionEvent_default(void* self, void* event) {
	((QSpinBox__Go*)self)->cppgoBase_actionEvent((QActionEvent*)event);
}
void QSpinBox_dragEnterEvent_default(void* self, void* event) {
	((QSpinBox__Go*)self)->cppgoBase_dragEnterEvent((QDragEnterEvent*)event);
}
void QSpinBox_dragMoveEvent_default(void* self, void* event) {
	((QSpinBox__Go*)self)->cppgoBase_dragMoveEvent((QDragMoveEvent*)event);
}
void QSpinBox_dragLeaveEvent_default(void* self, void* event) {
	((QSpinBox__Go*)self)->cppgoBase_dragLeaveEvent((QDragLeaveEvent*)event);
}
void QSpinBox_dropEvent_default(void* self, void* event) {
	((QSpinBox__Go*)self)->cppgoBase_dropEvent((QDropEvent*)event);
}
int QSpinBox_metric_default(void* self, int arg1) {
	auto&& __r = ((QSpinBox__Go*)self)->cppgoBase_metric(arg1);
	return __r;
}
void QSpinBox_initPainter_default(void* self, void* painter) {
	((QSpinBox__Go*)self)->cppgoBase_initPainter((QPainter*)painter);
}
void* QSpinBox_redirected_default(void* self, void* offset) {
	auto&& __r = ((QSpinBox__Go*)self)->cppgoBase_redirected((QPoint*)offset);
	return (void*)(__r);
}
void* QSpinBox_sharedPainter_default(void* self) {
	auto&& __r = ((QSpinBox__Go*)self)->cppgoBase_sharedPainter();
	return (void*)(__r);
}
void QSpinBox_inputMethodEvent_default(void* self, void* arg1) {
	((QSpinBox__Go*)self)->cppgoBase_inputMethodEvent((QInputMethodEvent*)arg1);
}
int QSpinBox_focusNextPrevChild_default(void* self, int next) {
	auto&& __r = ((QSpinBox__Go*)self)->cppgoBase_focusNextPrevChild(next);
	return __r;
}
int QSpinBox_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QSpinBox*)self)->QSpinBox::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QSpinBox_childEvent_default(void* self, void* event) {
	((QSpinBox__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QSpinBox_customEvent_default(void* self, void* event) {
	((QSpinBox__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QSpinBox_connectNotify_default(void* self, void* signal) {
	((QSpinBox__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QSpinBox_disconnectNotify_default(void* self, void* signal) {
	((QSpinBox__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QSpinBox_connect_valueChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSpinBox*)self, &QSpinBox::valueChanged, (QSpinBox*)self, [__id](int arg1){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSpinBox);
	__conns_QSpinBox[__id] = __c;
}
void QSpinBox_connect_textChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSpinBox*)self, &QSpinBox::textChanged, (QSpinBox*)self, [__id](const QString& arg1){
		QByteArray arg1_ba = arg1.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)arg1_ba.constData(), arg1_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSpinBox);
	__conns_QSpinBox[__id] = __c;
}
void QSpinBox_connect_editingFinished(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSpinBox*)self, &QSpinBox::editingFinished, (QSpinBox*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSpinBox);
	__conns_QSpinBox[__id] = __c;
}
void QSpinBox_connect_returnPressed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSpinBox*)self, &QSpinBox::returnPressed, (QSpinBox*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSpinBox);
	__conns_QSpinBox[__id] = __c;
}
void QSpinBox_connect_windowTitleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSpinBox*)self, &QSpinBox::windowTitleChanged, (QSpinBox*)self, [__id](const QString& title){
		QByteArray title_ba = title.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)title_ba.constData(), title_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSpinBox);
	__conns_QSpinBox[__id] = __c;
}
void QSpinBox_connect_windowIconChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSpinBox*)self, &QSpinBox::windowIconChanged, (QSpinBox*)self, [__id](const QIcon& icon){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&icon);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSpinBox);
	__conns_QSpinBox[__id] = __c;
}
void QSpinBox_connect_windowIconTextChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSpinBox*)self, &QSpinBox::windowIconTextChanged, (QSpinBox*)self, [__id](const QString& iconText){
		QByteArray iconText_ba = iconText.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)iconText_ba.constData(), iconText_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSpinBox);
	__conns_QSpinBox[__id] = __c;
}
void QSpinBox_connect_customContextMenuRequested(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSpinBox*)self, &QSpinBox::customContextMenuRequested, (QSpinBox*)self, [__id](const QPoint& pos){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&pos);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSpinBox);
	__conns_QSpinBox[__id] = __c;
}
void QSpinBox_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSpinBox*)self, &QSpinBox::destroyed, (QSpinBox*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSpinBox);
	__conns_QSpinBox[__id] = __c;
}
void QSpinBox_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSpinBox*)self, &QSpinBox::destroyed, (QSpinBox*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSpinBox);
	__conns_QSpinBox[__id] = __c;
}
void QSpinBox_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSpinBox*)self, &QSpinBox::objectNameChanged, (QSpinBox*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSpinBox);
	__conns_QSpinBox[__id] = __c;
}
void QSpinBox__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QSpinBox);
		auto __it = __conns_QSpinBox.find(id);
		if (__it == __conns_QSpinBox.end()) return;
		__c = __it->second;
		__conns_QSpinBox.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QSpinBox__Go_new(long long id) { return new QSpinBox__Go(id); }
void QSpinBox_delete(void* self) { delete (QSpinBox*)self; }

}

#include <QtWidgets/qtreewidget.h>
#include <QtCore/qstring.h>
#include <QtCore/qnamespace.h>
#include <QtGui/qicon.h>
#include <QtGui/qfont.h>
#include <QtGui/qbrush.h>
#include <QtCore/qsize.h>
#include <QtCore/qvariant.h>
#include <QtCore/qdatastream.h>
#include "gen_qtreewidgetitem.h"
#include "gen_purego_cb.h"

class QTreeWidgetItem__Go : public QTreeWidgetItem {
public:
	long long __goId;
	unsigned char __goOn[5] = {};
	template<typename... A> QTreeWidgetItem__Go(long long id, A&&... a) : QTreeWidgetItem(std::forward<A>(a)...), __goId(id) {}
	using QTreeWidgetItem::emitDataChanged;
	QTreeWidgetItem* clone() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, -7032813281048616478LL, &__h); if (__h) return (QTreeWidgetItem*)__r; }
		return QTreeWidgetItem::clone();
	}
	QVariant data(int column, int role) const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_voidp_int_int) { void* __r = ((void*(*)(long long, long long, int, int, int*))cppgo_virt_voidp_int_int)(__goId, -8837375942098466299LL, column, role, &__h); if (__h) return *(QVariant*)__r; }
		return QTreeWidgetItem::data(column, role);
	}
	void setData(int column, int role, const QVariant& value) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_void_int_int_voidp) { ((void(*)(long long, long long, int, int, void*, int*))cppgo_virt_void_int_int_voidp)(__goId, 35191454575202219LL, column, role, (void*)&value, &__h); if (__h) return; }
		QTreeWidgetItem::setData(column, role, value);
	}
	void read(QDataStream& in) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5541207787354052101LL, (void*)&in, &__h); if (__h) return; }
		QTreeWidgetItem::read(in);
	}
	void write(QDataStream& out) const override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5099743077423337732LL, (void*)&out, &__h); if (__h) return; }
		QTreeWidgetItem::write(out);
	}
};

extern "C" {

void QTreeWidgetItem__set_override(void* self, int idx) { ((QTreeWidgetItem__Go*)self)->__goOn[idx] = 1; }
void* QTreeWidgetItem_new(long long __goid) {
	return new QTreeWidgetItem__Go(__goid);
}
void* QTreeWidgetItem_new2(long long __goid, char** strings, long long strings_len) {
	QList<QString> strings_list; strings_list.reserve(strings_len);
	for (long long __i = 0; __i < strings_len; __i++) strings_list.append(QString::fromUtf8(strings[__i]));
	return new QTreeWidgetItem__Go(__goid, strings_list);
}
void* QTreeWidgetItem_new3(long long __goid, void* treeview) {
	return new QTreeWidgetItem__Go(__goid, (QTreeWidget*)treeview);
}
void* QTreeWidgetItem_new4(long long __goid, void* treeview, char** strings, long long strings_len) {
	QList<QString> strings_list; strings_list.reserve(strings_len);
	for (long long __i = 0; __i < strings_len; __i++) strings_list.append(QString::fromUtf8(strings[__i]));
	return new QTreeWidgetItem__Go(__goid, (QTreeWidget*)treeview, strings_list);
}
void* QTreeWidgetItem_new5(long long __goid, void* treeview, void* after) {
	return new QTreeWidgetItem__Go(__goid, (QTreeWidget*)treeview, (QTreeWidgetItem*)after);
}
void* QTreeWidgetItem_new6(long long __goid, void* parent) {
	return new QTreeWidgetItem__Go(__goid, (QTreeWidgetItem*)parent);
}
void* QTreeWidgetItem_new7(long long __goid, void* parent, char** strings, long long strings_len) {
	QList<QString> strings_list; strings_list.reserve(strings_len);
	for (long long __i = 0; __i < strings_len; __i++) strings_list.append(QString::fromUtf8(strings[__i]));
	return new QTreeWidgetItem__Go(__goid, (QTreeWidgetItem*)parent, strings_list);
}
void* QTreeWidgetItem_new8(long long __goid, void* parent, void* after) {
	return new QTreeWidgetItem__Go(__goid, (QTreeWidgetItem*)parent, (QTreeWidgetItem*)after);
}
void* QTreeWidgetItem_new9(long long __goid, void* other) {
	return new QTreeWidgetItem__Go(__goid, *(QTreeWidgetItem*)other);
}
void* QTreeWidgetItem_new10(long long __goid, int type_) {
	return new QTreeWidgetItem__Go(__goid, type_);
}
void* QTreeWidgetItem_new11(long long __goid, char** strings, long long strings_len, int type_) {
	QList<QString> strings_list; strings_list.reserve(strings_len);
	for (long long __i = 0; __i < strings_len; __i++) strings_list.append(QString::fromUtf8(strings[__i]));
	return new QTreeWidgetItem__Go(__goid, strings_list, type_);
}
void* QTreeWidgetItem_new12(long long __goid, void* treeview, int type_) {
	return new QTreeWidgetItem__Go(__goid, (QTreeWidget*)treeview, type_);
}
void* QTreeWidgetItem_new13(long long __goid, void* treeview, char** strings, long long strings_len, int type_) {
	QList<QString> strings_list; strings_list.reserve(strings_len);
	for (long long __i = 0; __i < strings_len; __i++) strings_list.append(QString::fromUtf8(strings[__i]));
	return new QTreeWidgetItem__Go(__goid, (QTreeWidget*)treeview, strings_list, type_);
}
void* QTreeWidgetItem_new14(long long __goid, void* treeview, void* after, int type_) {
	return new QTreeWidgetItem__Go(__goid, (QTreeWidget*)treeview, (QTreeWidgetItem*)after, type_);
}
void* QTreeWidgetItem_new15(long long __goid, void* parent, int type_) {
	return new QTreeWidgetItem__Go(__goid, (QTreeWidgetItem*)parent, type_);
}
void* QTreeWidgetItem_new16(long long __goid, void* parent, char** strings, long long strings_len, int type_) {
	QList<QString> strings_list; strings_list.reserve(strings_len);
	for (long long __i = 0; __i < strings_len; __i++) strings_list.append(QString::fromUtf8(strings[__i]));
	return new QTreeWidgetItem__Go(__goid, (QTreeWidgetItem*)parent, strings_list, type_);
}
void* QTreeWidgetItem_new17(long long __goid, void* parent, void* after, int type_) {
	return new QTreeWidgetItem__Go(__goid, (QTreeWidgetItem*)parent, (QTreeWidgetItem*)after, type_);
}
void* QTreeWidgetItem_clone(void* self) {
	auto&& __r = ((QTreeWidgetItem*)self)->clone();
	return (void*)(__r);
}
void* QTreeWidgetItem_treeWidget(void* self) {
	auto&& __r = ((QTreeWidgetItem*)self)->treeWidget();
	return (void*)(__r);
}
void QTreeWidgetItem_setSelected(void* self, int select_) {
	((QTreeWidgetItem*)self)->setSelected(select_);
}
int QTreeWidgetItem_isSelected(void* self) {
	auto&& __r = ((QTreeWidgetItem*)self)->isSelected();
	return __r;
}
void QTreeWidgetItem_setHidden(void* self, int hide) {
	((QTreeWidgetItem*)self)->setHidden(hide);
}
int QTreeWidgetItem_isHidden(void* self) {
	auto&& __r = ((QTreeWidgetItem*)self)->isHidden();
	return __r;
}
void QTreeWidgetItem_setExpanded(void* self, int expand) {
	((QTreeWidgetItem*)self)->setExpanded(expand);
}
int QTreeWidgetItem_isExpanded(void* self) {
	auto&& __r = ((QTreeWidgetItem*)self)->isExpanded();
	return __r;
}
void QTreeWidgetItem_setFirstColumnSpanned(void* self, int span) {
	((QTreeWidgetItem*)self)->setFirstColumnSpanned(span);
}
int QTreeWidgetItem_isFirstColumnSpanned(void* self) {
	auto&& __r = ((QTreeWidgetItem*)self)->isFirstColumnSpanned();
	return __r;
}
void QTreeWidgetItem_setDisabled(void* self, int disabled) {
	((QTreeWidgetItem*)self)->setDisabled(disabled);
}
int QTreeWidgetItem_isDisabled(void* self) {
	auto&& __r = ((QTreeWidgetItem*)self)->isDisabled();
	return __r;
}
void QTreeWidgetItem_setChildIndicatorPolicy(void* self, int policy) {
	((QTreeWidgetItem*)self)->setChildIndicatorPolicy((QTreeWidgetItem::ChildIndicatorPolicy)(policy));
}
int QTreeWidgetItem_childIndicatorPolicy(void* self) {
	auto&& __r = ((QTreeWidgetItem*)self)->childIndicatorPolicy();
	return (int)(__r);
}
int QTreeWidgetItem_flags(void* self) {
	auto&& __r = ((QTreeWidgetItem*)self)->flags();
	return (int)(__r);
}
void QTreeWidgetItem_setFlags(void* self, int flags) {
	((QTreeWidgetItem*)self)->setFlags(QFlags<Qt::ItemFlag>(flags));
}
char* QTreeWidgetItem_text(void* self, int column, long long* __out_len) {
	auto&& __r = ((QTreeWidgetItem*)self)->text(column);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QTreeWidgetItem_setText(void* self, int column, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	((QTreeWidgetItem*)self)->setText(column, text_qstr);
}
void* QTreeWidgetItem_icon(void* self, int column) {
	auto&& __r = ((QTreeWidgetItem*)self)->icon(column);
	return new QIcon(__r);
}
void QTreeWidgetItem_setIcon(void* self, int column, void* icon) {
	((QTreeWidgetItem*)self)->setIcon(column, *(QIcon*)icon);
}
char* QTreeWidgetItem_statusTip(void* self, int column, long long* __out_len) {
	auto&& __r = ((QTreeWidgetItem*)self)->statusTip(column);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QTreeWidgetItem_setStatusTip(void* self, int column, const char* statusTip, long long statusTip_len) {
	QString statusTip_qstr = QString::fromUtf8(statusTip, statusTip_len);
	((QTreeWidgetItem*)self)->setStatusTip(column, statusTip_qstr);
}
char* QTreeWidgetItem_toolTip(void* self, int column, long long* __out_len) {
	auto&& __r = ((QTreeWidgetItem*)self)->toolTip(column);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QTreeWidgetItem_setToolTip(void* self, int column, const char* toolTip, long long toolTip_len) {
	QString toolTip_qstr = QString::fromUtf8(toolTip, toolTip_len);
	((QTreeWidgetItem*)self)->setToolTip(column, toolTip_qstr);
}
char* QTreeWidgetItem_whatsThis(void* self, int column, long long* __out_len) {
	auto&& __r = ((QTreeWidgetItem*)self)->whatsThis(column);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QTreeWidgetItem_setWhatsThis(void* self, int column, const char* whatsThis, long long whatsThis_len) {
	QString whatsThis_qstr = QString::fromUtf8(whatsThis, whatsThis_len);
	((QTreeWidgetItem*)self)->setWhatsThis(column, whatsThis_qstr);
}
void* QTreeWidgetItem_font(void* self, int column) {
	auto&& __r = ((QTreeWidgetItem*)self)->font(column);
	return new QFont(__r);
}
void QTreeWidgetItem_setFont(void* self, int column, void* font) {
	((QTreeWidgetItem*)self)->setFont(column, *(QFont*)font);
}
int QTreeWidgetItem_textAlignment(void* self, int column) {
	auto&& __r = ((QTreeWidgetItem*)self)->textAlignment(column);
	return __r;
}
void QTreeWidgetItem_setTextAlignment(void* self, int column, int alignment) {
	((QTreeWidgetItem*)self)->setTextAlignment(column, alignment);
}
void QTreeWidgetItem_setTextAlignment2(void* self, int column, int alignment) {
	((QTreeWidgetItem*)self)->setTextAlignment(column, (Qt::AlignmentFlag)(alignment));
}
void QTreeWidgetItem_setTextAlignment3(void* self, int column, int alignment) {
	((QTreeWidgetItem*)self)->setTextAlignment(column, QFlags<Qt::AlignmentFlag>(alignment));
}
void* QTreeWidgetItem_background(void* self, int column) {
	auto&& __r = ((QTreeWidgetItem*)self)->background(column);
	return new QBrush(__r);
}
void QTreeWidgetItem_setBackground(void* self, int column, void* brush) {
	((QTreeWidgetItem*)self)->setBackground(column, *(QBrush*)brush);
}
void* QTreeWidgetItem_foreground(void* self, int column) {
	auto&& __r = ((QTreeWidgetItem*)self)->foreground(column);
	return new QBrush(__r);
}
void QTreeWidgetItem_setForeground(void* self, int column, void* brush) {
	((QTreeWidgetItem*)self)->setForeground(column, *(QBrush*)brush);
}
int QTreeWidgetItem_checkState(void* self, int column) {
	auto&& __r = ((QTreeWidgetItem*)self)->checkState(column);
	return (int)(__r);
}
void QTreeWidgetItem_setCheckState(void* self, int column, int state) {
	((QTreeWidgetItem*)self)->setCheckState(column, (Qt::CheckState)(state));
}
void* QTreeWidgetItem_sizeHint(void* self, int column) {
	auto&& __r = ((QTreeWidgetItem*)self)->sizeHint(column);
	return new QSize(__r);
}
void QTreeWidgetItem_setSizeHint(void* self, int column, void* size) {
	((QTreeWidgetItem*)self)->setSizeHint(column, *(QSize*)size);
}
void* QTreeWidgetItem_data(void* self, int column, int role) {
	auto&& __r = ((QTreeWidgetItem*)self)->data(column, role);
	return new QVariant(__r);
}
void QTreeWidgetItem_setData(void* self, int column, int role, void* value) {
	((QTreeWidgetItem*)self)->setData(column, role, *(QVariant*)value);
}
int QTreeWidgetItem_OperatorLesser(void* self, void* other) {
	auto&& __r = ((QTreeWidgetItem*)self)->operator<(*(QTreeWidgetItem*)other);
	return __r;
}
void QTreeWidgetItem_read(void* self, void* in) {
	((QTreeWidgetItem*)self)->read(*(QDataStream*)in);
}
void QTreeWidgetItem_write(void* self, void* out) {
	((QTreeWidgetItem*)self)->write(*(QDataStream*)out);
}
void* QTreeWidgetItem_OperatorAssign(void* self, void* other) {
	auto&& __r = ((QTreeWidgetItem*)self)->operator=(*(QTreeWidgetItem*)other);
	return (void*)&(__r);
}
void* QTreeWidgetItem_parent(void* self) {
	auto&& __r = ((QTreeWidgetItem*)self)->parent();
	return (void*)(__r);
}
void* QTreeWidgetItem_child(void* self, int index) {
	auto&& __r = ((QTreeWidgetItem*)self)->child(index);
	return (void*)(__r);
}
int QTreeWidgetItem_childCount(void* self) {
	auto&& __r = ((QTreeWidgetItem*)self)->childCount();
	return __r;
}
int QTreeWidgetItem_columnCount(void* self) {
	auto&& __r = ((QTreeWidgetItem*)self)->columnCount();
	return __r;
}
int QTreeWidgetItem_indexOfChild(void* self, void* child) {
	auto&& __r = ((QTreeWidgetItem*)self)->indexOfChild((QTreeWidgetItem*)child);
	return __r;
}
void QTreeWidgetItem_addChild(void* self, void* child) {
	((QTreeWidgetItem*)self)->addChild((QTreeWidgetItem*)child);
}
void QTreeWidgetItem_insertChild(void* self, int index, void* child) {
	((QTreeWidgetItem*)self)->insertChild(index, (QTreeWidgetItem*)child);
}
void QTreeWidgetItem_removeChild(void* self, void* child) {
	((QTreeWidgetItem*)self)->removeChild((QTreeWidgetItem*)child);
}
void* QTreeWidgetItem_takeChild(void* self, int index) {
	auto&& __r = ((QTreeWidgetItem*)self)->takeChild(index);
	return (void*)(__r);
}
void QTreeWidgetItem_addChildren(void* self, void** children, long long children_len) {
	QList<QTreeWidgetItem*> children_list; children_list.reserve(children_len);
	for (long long __i = 0; __i < children_len; __i++) children_list.append((QTreeWidgetItem*)children[__i]);
	((QTreeWidgetItem*)self)->addChildren(children_list);
}
void QTreeWidgetItem_insertChildren(void* self, int index, void** children, long long children_len) {
	QList<QTreeWidgetItem*> children_list; children_list.reserve(children_len);
	for (long long __i = 0; __i < children_len; __i++) children_list.append((QTreeWidgetItem*)children[__i]);
	((QTreeWidgetItem*)self)->insertChildren(index, children_list);
}
void** QTreeWidgetItem_takeChildren(void* self, long long* __out_len) {
	auto&& __r = ((QTreeWidgetItem*)self)->takeChildren();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
int QTreeWidgetItem_type(void* self) {
	auto&& __r = ((QTreeWidgetItem*)self)->type();
	return __r;
}
void QTreeWidgetItem_sortChildren(void* self, int column, int order) {
	((QTreeWidgetItem*)self)->sortChildren(column, (Qt::SortOrder)(order));
}
void QTreeWidgetItem_emitDataChanged(void* self) {
	((QTreeWidgetItem__Go*)self)->emitDataChanged();
}
void* QTreeWidgetItem_clone_default(void* self) {
	auto&& __r = ((QTreeWidgetItem*)self)->QTreeWidgetItem::clone();
	return (void*)(__r);
}
void* QTreeWidgetItem_data_default(void* self, int column, int role) {
	auto&& __r = ((QTreeWidgetItem*)self)->QTreeWidgetItem::data(column, role);
	return new QVariant(__r);
}
void QTreeWidgetItem_setData_default(void* self, int column, int role, void* value) {
	((QTreeWidgetItem*)self)->QTreeWidgetItem::setData(column, role, *(QVariant*)value);
}
void QTreeWidgetItem_read_default(void* self, void* in) {
	((QTreeWidgetItem*)self)->QTreeWidgetItem::read(*(QDataStream*)in);
}
void QTreeWidgetItem_write_default(void* self, void* out) {
	((QTreeWidgetItem*)self)->QTreeWidgetItem::write(*(QDataStream*)out);
}
void* QTreeWidgetItem__Go_new(long long id) { return new QTreeWidgetItem__Go(id); }
void QTreeWidgetItem_delete(void* self) { delete (QTreeWidgetItem*)self; }

}

