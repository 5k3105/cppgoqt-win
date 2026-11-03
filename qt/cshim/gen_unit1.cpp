#include <QtCore/qlogging.h>
#include "gen_qmessagelogcontext.h"

extern "C" {

void* QMessageLogContext_new() {
	return new QMessageLogContext();
}
void* QMessageLogContext_new2(const char* fileName, int lineNumber, const char* functionName, const char* categoryName) {
	return new QMessageLogContext(fileName, lineNumber, functionName, categoryName);
}
int QMessageLogContext_version_get(void* self) {
	auto __r = ((QMessageLogContext*)self)->version;
	return __r;
}
void QMessageLogContext_version_set(void* self, int version) {
	((QMessageLogContext*)self)->version = version;
}
int QMessageLogContext_line_get(void* self) {
	auto __r = ((QMessageLogContext*)self)->line;
	return __r;
}
void QMessageLogContext_line_set(void* self, int line) {
	((QMessageLogContext*)self)->line = line;
}
const char* QMessageLogContext_file_get(void* self) {
	auto __r = ((QMessageLogContext*)self)->file;
	return __r;
}
void QMessageLogContext_file_set(void* self, const char* file) {
	((QMessageLogContext*)self)->file = file;
}
const char* QMessageLogContext_function_get(void* self) {
	auto __r = ((QMessageLogContext*)self)->function;
	return __r;
}
void QMessageLogContext_function_set(void* self, const char* function) {
	((QMessageLogContext*)self)->function = function;
}
const char* QMessageLogContext_category_get(void* self) {
	auto __r = ((QMessageLogContext*)self)->category;
	return __r;
}
void QMessageLogContext_category_set(void* self, const char* category) {
	((QMessageLogContext*)self)->category = category;
}
void QMessageLogContext_delete(void* self) { delete (QMessageLogContext*)self; }

}

#include <QtCore/qstringview.h>
#include <QtCore/qstring.h>
#include <QtCore/qchar.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qlatin1stringview.h>
#include <QtCore/qregularexpression.h>
#include <QtCore/qbytearrayview.h>
#include "gen_qstringview.h"

extern "C" {

void* QStringView_new() {
	return new QStringView();
}
void* QStringView_new2(void* arg1) {
	return new QStringView(*(QStringView*)arg1);
}
char* QStringView_toString(void* self, long long* __out_len) {
	auto&& __r = ((QStringView*)self)->toString();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
long long QStringView_size(void* self) {
	auto&& __r = ((QStringView*)self)->size();
	return __r;
}
void* QStringView_data(void* self) {
	auto&& __r = ((QStringView*)self)->data();
	return (void*)(__r);
}
void* QStringView_constData(void* self) {
	auto&& __r = ((QStringView*)self)->constData();
	return (void*)(__r);
}
void* QStringView_OperatorSubscript(void* self, long long n) {
	auto&& __r = ((QStringView*)self)->operator[](n);
	return new QChar(__r);
}
void* QStringView_toLatin1(void* self) {
	auto&& __r = ((QStringView*)self)->toLatin1();
	return new QByteArray(__r);
}
void* QStringView_toUtf8(void* self) {
	auto&& __r = ((QStringView*)self)->toUtf8();
	return new QByteArray(__r);
}
void* QStringView_toLocal8Bit(void* self) {
	auto&& __r = ((QStringView*)self)->toLocal8Bit();
	return new QByteArray(__r);
}
unsigned int* QStringView_toUcs4(void* self, long long* __out_len) {
	auto&& __r = ((QStringView*)self)->toUcs4();
	*__out_len = __r.size();
	unsigned int* __arr = (unsigned int*)malloc(sizeof(unsigned int) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = __e;
	return __arr;
}
void* QStringView_at(void* self, long long n) {
	auto&& __r = ((QStringView*)self)->at(n);
	return new QChar(__r);
}
void* QStringView_mid(void* self, long long pos) {
	auto&& __r = ((QStringView*)self)->mid(pos);
	return new QStringView(__r);
}
void* QStringView_left(void* self, long long n) {
	auto&& __r = ((QStringView*)self)->left(n);
	return new QStringView(__r);
}
void* QStringView_right(void* self, long long n) {
	auto&& __r = ((QStringView*)self)->right(n);
	return new QStringView(__r);
}
void* QStringView_first(void* self, long long n) {
	auto&& __r = ((QStringView*)self)->first(n);
	return new QStringView(__r);
}
void* QStringView_last(void* self, long long n) {
	auto&& __r = ((QStringView*)self)->last(n);
	return new QStringView(__r);
}
void* QStringView_sliced(void* self, long long pos) {
	auto&& __r = ((QStringView*)self)->sliced(pos);
	return new QStringView(__r);
}
void* QStringView_sliced2(void* self, long long pos, long long n) {
	auto&& __r = ((QStringView*)self)->sliced(pos, n);
	return new QStringView(__r);
}
void* QStringView_chopped(void* self, long long n) {
	auto&& __r = ((QStringView*)self)->chopped(n);
	return new QStringView(__r);
}
void QStringView_truncate(void* self, long long n) {
	((QStringView*)self)->truncate(n);
}
void QStringView_chop(void* self, long long n) {
	((QStringView*)self)->chop(n);
}
void* QStringView_trimmed(void* self) {
	auto&& __r = ((QStringView*)self)->trimmed();
	return new QStringView(__r);
}
void* QStringView_slice(void* self, long long pos) {
	auto&& __r = ((QStringView*)self)->slice(pos);
	return (void*)&(__r);
}
void* QStringView_slice2(void* self, long long pos, long long n) {
	auto&& __r = ((QStringView*)self)->slice(pos, n);
	return (void*)&(__r);
}
int QStringView_compare(void* self, void* other) {
	auto&& __r = ((QStringView*)self)->compare(*(QStringView*)other);
	return __r;
}
int QStringView_compare2(void* self, void* other) {
	auto&& __r = ((QStringView*)self)->compare(*(QLatin1String*)other);
	return __r;
}
int QStringView_compare3(void* self, void* c) {
	auto&& __r = ((QStringView*)self)->compare(*(QChar*)c);
	return __r;
}
int QStringView_compare4(void* self, void* c, int cs) {
	auto&& __r = ((QStringView*)self)->compare(*(QChar*)c, (Qt::CaseSensitivity)(cs));
	return __r;
}
int QStringView_localeAwareCompare(void* self, void* other) {
	auto&& __r = ((QStringView*)self)->localeAwareCompare(*(QStringView*)other);
	return __r;
}
int QStringView_startsWith(void* self, void* s) {
	auto&& __r = ((QStringView*)self)->startsWith(*(QStringView*)s);
	return __r;
}
int QStringView_startsWith2(void* self, void* s) {
	auto&& __r = ((QStringView*)self)->startsWith(*(QLatin1String*)s);
	return __r;
}
int QStringView_startsWith3(void* self, void* c) {
	auto&& __r = ((QStringView*)self)->startsWith(*(QChar*)c);
	return __r;
}
int QStringView_startsWith4(void* self, void* c, int cs) {
	auto&& __r = ((QStringView*)self)->startsWith(*(QChar*)c, (Qt::CaseSensitivity)(cs));
	return __r;
}
int QStringView_endsWith(void* self, void* s) {
	auto&& __r = ((QStringView*)self)->endsWith(*(QStringView*)s);
	return __r;
}
int QStringView_endsWith2(void* self, void* s) {
	auto&& __r = ((QStringView*)self)->endsWith(*(QLatin1String*)s);
	return __r;
}
int QStringView_endsWith3(void* self, void* c) {
	auto&& __r = ((QStringView*)self)->endsWith(*(QChar*)c);
	return __r;
}
int QStringView_endsWith4(void* self, void* c, int cs) {
	auto&& __r = ((QStringView*)self)->endsWith(*(QChar*)c, (Qt::CaseSensitivity)(cs));
	return __r;
}
long long QStringView_indexOf(void* self, void* c) {
	auto&& __r = ((QStringView*)self)->indexOf(*(QChar*)c);
	return __r;
}
long long QStringView_indexOf2(void* self, void* s) {
	auto&& __r = ((QStringView*)self)->indexOf(*(QStringView*)s);
	return __r;
}
long long QStringView_indexOf3(void* self, void* s) {
	auto&& __r = ((QStringView*)self)->indexOf(*(QLatin1String*)s);
	return __r;
}
int QStringView_contains(void* self, void* c) {
	auto&& __r = ((QStringView*)self)->contains(*(QChar*)c);
	return __r;
}
int QStringView_contains2(void* self, void* s) {
	auto&& __r = ((QStringView*)self)->contains(*(QStringView*)s);
	return __r;
}
int QStringView_contains3(void* self, void* s) {
	auto&& __r = ((QStringView*)self)->contains(*(QLatin1String*)s);
	return __r;
}
long long QStringView_count(void* self, void* c) {
	auto&& __r = ((QStringView*)self)->count(*(QChar*)c);
	return __r;
}
long long QStringView_count2(void* self, void* s) {
	auto&& __r = ((QStringView*)self)->count(*(QStringView*)s);
	return __r;
}
long long QStringView_count3(void* self, void* s) {
	auto&& __r = ((QStringView*)self)->count(*(QLatin1String*)s);
	return __r;
}
long long QStringView_lastIndexOf(void* self, void* c) {
	auto&& __r = ((QStringView*)self)->lastIndexOf(*(QChar*)c);
	return __r;
}
long long QStringView_lastIndexOf2(void* self, void* c, long long from) {
	auto&& __r = ((QStringView*)self)->lastIndexOf(*(QChar*)c, from);
	return __r;
}
long long QStringView_lastIndexOf3(void* self, void* s) {
	auto&& __r = ((QStringView*)self)->lastIndexOf(*(QStringView*)s);
	return __r;
}
long long QStringView_lastIndexOf4(void* self, void* s, long long from) {
	auto&& __r = ((QStringView*)self)->lastIndexOf(*(QStringView*)s, from);
	return __r;
}
long long QStringView_lastIndexOf5(void* self, void* s) {
	auto&& __r = ((QStringView*)self)->lastIndexOf(*(QLatin1String*)s);
	return __r;
}
long long QStringView_lastIndexOf6(void* self, void* s, long long from) {
	auto&& __r = ((QStringView*)self)->lastIndexOf(*(QLatin1String*)s, from);
	return __r;
}
long long QStringView_indexOf4(void* self, void* re) {
	auto&& __r = ((QStringView*)self)->indexOf(*(QRegularExpression*)re);
	return __r;
}
long long QStringView_lastIndexOf7(void* self, void* re, long long from) {
	auto&& __r = ((QStringView*)self)->lastIndexOf(*(QRegularExpression*)re, from);
	return __r;
}
int QStringView_contains4(void* self, void* re) {
	auto&& __r = ((QStringView*)self)->contains(*(QRegularExpression*)re);
	return __r;
}
long long QStringView_count4(void* self, void* re) {
	auto&& __r = ((QStringView*)self)->count(*(QRegularExpression*)re);
	return __r;
}
int QStringView_isRightToLeft(void* self) {
	auto&& __r = ((QStringView*)self)->isRightToLeft();
	return __r;
}
int QStringView_isValidUtf16(void* self) {
	auto&& __r = ((QStringView*)self)->isValidUtf16();
	return __r;
}
int QStringView_isUpper(void* self) {
	auto&& __r = ((QStringView*)self)->isUpper();
	return __r;
}
int QStringView_isLower(void* self) {
	auto&& __r = ((QStringView*)self)->isLower();
	return __r;
}
short QStringView_toShort(void* self) {
	auto&& __r = ((QStringView*)self)->toShort();
	return __r;
}
unsigned short QStringView_toUShort(void* self) {
	auto&& __r = ((QStringView*)self)->toUShort();
	return __r;
}
int QStringView_toInt(void* self) {
	auto&& __r = ((QStringView*)self)->toInt();
	return __r;
}
unsigned int QStringView_toUInt(void* self) {
	auto&& __r = ((QStringView*)self)->toUInt();
	return __r;
}
long QStringView_toLong(void* self) {
	auto&& __r = ((QStringView*)self)->toLong();
	return __r;
}
unsigned long QStringView_toULong(void* self) {
	auto&& __r = ((QStringView*)self)->toULong();
	return __r;
}
long long QStringView_toLongLong(void* self) {
	auto&& __r = ((QStringView*)self)->toLongLong();
	return __r;
}
unsigned long long QStringView_toULongLong(void* self) {
	auto&& __r = ((QStringView*)self)->toULongLong();
	return __r;
}
float QStringView_toFloat(void* self) {
	auto&& __r = ((QStringView*)self)->toFloat();
	return __r;
}
double QStringView_toDouble(void* self) {
	auto&& __r = ((QStringView*)self)->toDouble();
	return __r;
}
void** QStringView_split(void* self, void* sep, long long* __out_len) {
	auto&& __r = ((QStringView*)self)->split(*(QStringView*)sep);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QStringView(__e);
	return __arr;
}
void** QStringView_split2(void* self, void* sep, long long* __out_len) {
	auto&& __r = ((QStringView*)self)->split(*(QChar*)sep);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QStringView(__e);
	return __arr;
}
void** QStringView_split3(void* self, void* sep, long long* __out_len) {
	auto&& __r = ((QStringView*)self)->split(*(QRegularExpression*)sep);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QStringView(__e);
	return __arr;
}
void* QStringView_begin(void* self) {
	auto&& __r = ((QStringView*)self)->begin();
	return (void*)(__r);
}
void* QStringView_end(void* self) {
	auto&& __r = ((QStringView*)self)->end();
	return (void*)(__r);
}
void* QStringView_cbegin(void* self) {
	auto&& __r = ((QStringView*)self)->cbegin();
	return (void*)(__r);
}
void* QStringView_cend(void* self) {
	auto&& __r = ((QStringView*)self)->cend();
	return (void*)(__r);
}
int QStringView_empty(void* self) {
	auto&& __r = ((QStringView*)self)->empty();
	return __r;
}
void* QStringView_front(void* self) {
	auto&& __r = ((QStringView*)self)->front();
	return new QChar(__r);
}
void* QStringView_back(void* self) {
	auto&& __r = ((QStringView*)self)->back();
	return new QChar(__r);
}
long long QStringView_max_size(void* self) {
	auto&& __r = ((QStringView*)self)->max_size();
	return __r;
}
void* QStringView_constBegin(void* self) {
	auto&& __r = ((QStringView*)self)->constBegin();
	return (void*)(__r);
}
void* QStringView_constEnd(void* self) {
	auto&& __r = ((QStringView*)self)->constEnd();
	return (void*)(__r);
}
int QStringView_isNull(void* self) {
	auto&& __r = ((QStringView*)self)->isNull();
	return __r;
}
int QStringView_isEmpty(void* self) {
	auto&& __r = ((QStringView*)self)->isEmpty();
	return __r;
}
long long QStringView_length(void* self) {
	auto&& __r = ((QStringView*)self)->length();
	return __r;
}
void* QStringView_first2(void* self) {
	auto&& __r = ((QStringView*)self)->first();
	return new QChar(__r);
}
void* QStringView_last2(void* self) {
	auto&& __r = ((QStringView*)self)->last();
	return new QChar(__r);
}
long long QStringView_maxSize() {
	auto&& __r = QStringView::maxSize();
	return __r;
}
void* QStringView_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QStringView*)self)->operator=(*(QStringView*)arg1);
	return (void*)&(__r);
}
void* QStringView_mid2(void* self, long long pos, long long n) {
	auto&& __r = ((QStringView*)self)->mid(pos, n);
	return new QStringView(__r);
}
int QStringView_compare5(void* self, void* other, int cs) {
	auto&& __r = ((QStringView*)self)->compare(*(QStringView*)other, (Qt::CaseSensitivity)(cs));
	return __r;
}
int QStringView_compare6(void* self, void* other, int cs) {
	auto&& __r = ((QStringView*)self)->compare(*(QLatin1String*)other, (Qt::CaseSensitivity)(cs));
	return __r;
}
int QStringView_startsWith5(void* self, void* s, int cs) {
	auto&& __r = ((QStringView*)self)->startsWith(*(QStringView*)s, (Qt::CaseSensitivity)(cs));
	return __r;
}
int QStringView_startsWith6(void* self, void* s, int cs) {
	auto&& __r = ((QStringView*)self)->startsWith(*(QLatin1String*)s, (Qt::CaseSensitivity)(cs));
	return __r;
}
int QStringView_endsWith5(void* self, void* s, int cs) {
	auto&& __r = ((QStringView*)self)->endsWith(*(QStringView*)s, (Qt::CaseSensitivity)(cs));
	return __r;
}
int QStringView_endsWith6(void* self, void* s, int cs) {
	auto&& __r = ((QStringView*)self)->endsWith(*(QLatin1String*)s, (Qt::CaseSensitivity)(cs));
	return __r;
}
long long QStringView_indexOf5(void* self, void* c, long long from) {
	auto&& __r = ((QStringView*)self)->indexOf(*(QChar*)c, from);
	return __r;
}
long long QStringView_indexOf6(void* self, void* c, long long from, int cs) {
	auto&& __r = ((QStringView*)self)->indexOf(*(QChar*)c, from, (Qt::CaseSensitivity)(cs));
	return __r;
}
long long QStringView_indexOf7(void* self, void* s, long long from) {
	auto&& __r = ((QStringView*)self)->indexOf(*(QStringView*)s, from);
	return __r;
}
long long QStringView_indexOf8(void* self, void* s, long long from, int cs) {
	auto&& __r = ((QStringView*)self)->indexOf(*(QStringView*)s, from, (Qt::CaseSensitivity)(cs));
	return __r;
}
long long QStringView_indexOf9(void* self, void* s, long long from) {
	auto&& __r = ((QStringView*)self)->indexOf(*(QLatin1String*)s, from);
	return __r;
}
long long QStringView_indexOf10(void* self, void* s, long long from, int cs) {
	auto&& __r = ((QStringView*)self)->indexOf(*(QLatin1String*)s, from, (Qt::CaseSensitivity)(cs));
	return __r;
}
int QStringView_contains5(void* self, void* c, int cs) {
	auto&& __r = ((QStringView*)self)->contains(*(QChar*)c, (Qt::CaseSensitivity)(cs));
	return __r;
}
int QStringView_contains6(void* self, void* s, int cs) {
	auto&& __r = ((QStringView*)self)->contains(*(QStringView*)s, (Qt::CaseSensitivity)(cs));
	return __r;
}
int QStringView_contains7(void* self, void* s, int cs) {
	auto&& __r = ((QStringView*)self)->contains(*(QLatin1String*)s, (Qt::CaseSensitivity)(cs));
	return __r;
}
long long QStringView_count5(void* self, void* c, int cs) {
	auto&& __r = ((QStringView*)self)->count(*(QChar*)c, (Qt::CaseSensitivity)(cs));
	return __r;
}
long long QStringView_count6(void* self, void* s, int cs) {
	auto&& __r = ((QStringView*)self)->count(*(QStringView*)s, (Qt::CaseSensitivity)(cs));
	return __r;
}
long long QStringView_count7(void* self, void* s, int cs) {
	auto&& __r = ((QStringView*)self)->count(*(QLatin1String*)s, (Qt::CaseSensitivity)(cs));
	return __r;
}
long long QStringView_lastIndexOf8(void* self, void* c, int cs) {
	auto&& __r = ((QStringView*)self)->lastIndexOf(*(QChar*)c, (Qt::CaseSensitivity)(cs));
	return __r;
}
long long QStringView_lastIndexOf9(void* self, void* c, long long from, int cs) {
	auto&& __r = ((QStringView*)self)->lastIndexOf(*(QChar*)c, from, (Qt::CaseSensitivity)(cs));
	return __r;
}
long long QStringView_lastIndexOf10(void* self, void* s, int cs) {
	auto&& __r = ((QStringView*)self)->lastIndexOf(*(QStringView*)s, (Qt::CaseSensitivity)(cs));
	return __r;
}
long long QStringView_lastIndexOf11(void* self, void* s, long long from, int cs) {
	auto&& __r = ((QStringView*)self)->lastIndexOf(*(QStringView*)s, from, (Qt::CaseSensitivity)(cs));
	return __r;
}
long long QStringView_lastIndexOf12(void* self, void* s, int cs) {
	auto&& __r = ((QStringView*)self)->lastIndexOf(*(QLatin1String*)s, (Qt::CaseSensitivity)(cs));
	return __r;
}
long long QStringView_lastIndexOf13(void* self, void* s, long long from, int cs) {
	auto&& __r = ((QStringView*)self)->lastIndexOf(*(QLatin1String*)s, from, (Qt::CaseSensitivity)(cs));
	return __r;
}
long long QStringView_indexOf11(void* self, void* re, long long from) {
	auto&& __r = ((QStringView*)self)->indexOf(*(QRegularExpression*)re, from);
	return __r;
}
long long QStringView_indexOf12(void* self, void* re, long long from, void* rmatch) {
	auto&& __r = ((QStringView*)self)->indexOf(*(QRegularExpression*)re, from, (QRegularExpressionMatch*)rmatch);
	return __r;
}
long long QStringView_lastIndexOf14(void* self, void* re, long long from, void* rmatch) {
	auto&& __r = ((QStringView*)self)->lastIndexOf(*(QRegularExpression*)re, from, (QRegularExpressionMatch*)rmatch);
	return __r;
}
int QStringView_contains8(void* self, void* re, void* rmatch) {
	auto&& __r = ((QStringView*)self)->contains(*(QRegularExpression*)re, (QRegularExpressionMatch*)rmatch);
	return __r;
}
short QStringView_toShort2(void* self, int* ok) {
	bool ok_local{};
	auto&& __r = ((QStringView*)self)->toShort(&ok_local);
	*ok = ok_local;
	return __r;
}
short QStringView_toShort3(void* self, int* ok, int base) {
	bool ok_local{};
	auto&& __r = ((QStringView*)self)->toShort(&ok_local, base);
	*ok = ok_local;
	return __r;
}
unsigned short QStringView_toUShort2(void* self, int* ok) {
	bool ok_local{};
	auto&& __r = ((QStringView*)self)->toUShort(&ok_local);
	*ok = ok_local;
	return __r;
}
unsigned short QStringView_toUShort3(void* self, int* ok, int base) {
	bool ok_local{};
	auto&& __r = ((QStringView*)self)->toUShort(&ok_local, base);
	*ok = ok_local;
	return __r;
}
int QStringView_toInt2(void* self, int* ok) {
	bool ok_local{};
	auto&& __r = ((QStringView*)self)->toInt(&ok_local);
	*ok = ok_local;
	return __r;
}
int QStringView_toInt3(void* self, int* ok, int base) {
	bool ok_local{};
	auto&& __r = ((QStringView*)self)->toInt(&ok_local, base);
	*ok = ok_local;
	return __r;
}
unsigned int QStringView_toUInt2(void* self, int* ok) {
	bool ok_local{};
	auto&& __r = ((QStringView*)self)->toUInt(&ok_local);
	*ok = ok_local;
	return __r;
}
unsigned int QStringView_toUInt3(void* self, int* ok, int base) {
	bool ok_local{};
	auto&& __r = ((QStringView*)self)->toUInt(&ok_local, base);
	*ok = ok_local;
	return __r;
}
long QStringView_toLong2(void* self, int* ok) {
	bool ok_local{};
	auto&& __r = ((QStringView*)self)->toLong(&ok_local);
	*ok = ok_local;
	return __r;
}
long QStringView_toLong3(void* self, int* ok, int base) {
	bool ok_local{};
	auto&& __r = ((QStringView*)self)->toLong(&ok_local, base);
	*ok = ok_local;
	return __r;
}
unsigned long QStringView_toULong2(void* self, int* ok) {
	bool ok_local{};
	auto&& __r = ((QStringView*)self)->toULong(&ok_local);
	*ok = ok_local;
	return __r;
}
unsigned long QStringView_toULong3(void* self, int* ok, int base) {
	bool ok_local{};
	auto&& __r = ((QStringView*)self)->toULong(&ok_local, base);
	*ok = ok_local;
	return __r;
}
long long QStringView_toLongLong2(void* self, int* ok) {
	bool ok_local{};
	auto&& __r = ((QStringView*)self)->toLongLong(&ok_local);
	*ok = ok_local;
	return __r;
}
long long QStringView_toLongLong3(void* self, int* ok, int base) {
	bool ok_local{};
	auto&& __r = ((QStringView*)self)->toLongLong(&ok_local, base);
	*ok = ok_local;
	return __r;
}
unsigned long long QStringView_toULongLong2(void* self, int* ok) {
	bool ok_local{};
	auto&& __r = ((QStringView*)self)->toULongLong(&ok_local);
	*ok = ok_local;
	return __r;
}
unsigned long long QStringView_toULongLong3(void* self, int* ok, int base) {
	bool ok_local{};
	auto&& __r = ((QStringView*)self)->toULongLong(&ok_local, base);
	*ok = ok_local;
	return __r;
}
float QStringView_toFloat2(void* self, int* ok) {
	bool ok_local{};
	auto&& __r = ((QStringView*)self)->toFloat(&ok_local);
	*ok = ok_local;
	return __r;
}
double QStringView_toDouble2(void* self, int* ok) {
	bool ok_local{};
	auto&& __r = ((QStringView*)self)->toDouble(&ok_local);
	*ok = ok_local;
	return __r;
}
void** QStringView_split4(void* self, void* sep, int behavior, long long* __out_len) {
	auto&& __r = ((QStringView*)self)->split(*(QStringView*)sep, QFlags<Qt::SplitBehaviorFlags>(behavior));
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QStringView(__e);
	return __arr;
}
void** QStringView_split5(void* self, void* sep, int behavior, int cs, long long* __out_len) {
	auto&& __r = ((QStringView*)self)->split(*(QStringView*)sep, QFlags<Qt::SplitBehaviorFlags>(behavior), (Qt::CaseSensitivity)(cs));
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QStringView(__e);
	return __arr;
}
void** QStringView_split6(void* self, void* sep, int behavior, long long* __out_len) {
	auto&& __r = ((QStringView*)self)->split(*(QChar*)sep, QFlags<Qt::SplitBehaviorFlags>(behavior));
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QStringView(__e);
	return __arr;
}
void** QStringView_split7(void* self, void* sep, int behavior, int cs, long long* __out_len) {
	auto&& __r = ((QStringView*)self)->split(*(QChar*)sep, QFlags<Qt::SplitBehaviorFlags>(behavior), (Qt::CaseSensitivity)(cs));
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QStringView(__e);
	return __arr;
}
void** QStringView_split8(void* self, void* sep, int behavior, long long* __out_len) {
	auto&& __r = ((QStringView*)self)->split(*(QRegularExpression*)sep, QFlags<Qt::SplitBehaviorFlags>(behavior));
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QStringView(__e);
	return __arr;
}
int QStringView_OperatorEqual(void* self, void* rhs) {
	auto&& __r = (*(QStringView*)self) == (*(QStringView*)rhs);
	return __r;
}
int QStringView_OperatorNotEqual(void* self, void* rhs) {
	auto&& __r = (*(QStringView*)self) != (*(QStringView*)rhs);
	return __r;
}
int QStringView_OperatorLesser(void* self, void* rhs) {
	auto&& __r = (*(QStringView*)self) < (*(QStringView*)rhs);
	return __r;
}
int QStringView_OperatorGreater(void* self, void* rhs) {
	auto&& __r = (*(QStringView*)self) > (*(QStringView*)rhs);
	return __r;
}
int QStringView_OperatorLesserOrEqual(void* self, void* rhs) {
	auto&& __r = (*(QStringView*)self) <= (*(QStringView*)rhs);
	return __r;
}
int QStringView_OperatorGreaterOrEqual(void* self, void* rhs) {
	auto&& __r = (*(QStringView*)self) >= (*(QStringView*)rhs);
	return __r;
}
int QStringView_OperatorEqual2(void* self, void* rhs) {
	auto&& __r = (*(QStringView*)self) == (*(QChar*)rhs);
	return __r;
}
int QStringView_OperatorNotEqual2(void* self, void* rhs) {
	auto&& __r = (*(QStringView*)self) != (*(QChar*)rhs);
	return __r;
}
int QStringView_OperatorLesser2(void* self, void* rhs) {
	auto&& __r = (*(QStringView*)self) < (*(QChar*)rhs);
	return __r;
}
int QStringView_OperatorGreater2(void* self, void* rhs) {
	auto&& __r = (*(QStringView*)self) > (*(QChar*)rhs);
	return __r;
}
int QStringView_OperatorLesserOrEqual2(void* self, void* rhs) {
	auto&& __r = (*(QStringView*)self) <= (*(QChar*)rhs);
	return __r;
}
int QStringView_OperatorGreaterOrEqual2(void* self, void* rhs) {
	auto&& __r = (*(QStringView*)self) >= (*(QChar*)rhs);
	return __r;
}
int QStringView_OperatorEqual3(void* self, void* rhs) {
	auto&& __r = (*(QStringView*)self) == (*(QByteArrayView*)rhs);
	return __r;
}
int QStringView_OperatorNotEqual3(void* self, void* rhs) {
	auto&& __r = (*(QStringView*)self) != (*(QByteArrayView*)rhs);
	return __r;
}
int QStringView_OperatorLesser3(void* self, void* rhs) {
	auto&& __r = (*(QStringView*)self) < (*(QByteArrayView*)rhs);
	return __r;
}
int QStringView_OperatorGreater3(void* self, void* rhs) {
	auto&& __r = (*(QStringView*)self) > (*(QByteArrayView*)rhs);
	return __r;
}
int QStringView_OperatorLesserOrEqual3(void* self, void* rhs) {
	auto&& __r = (*(QStringView*)self) <= (*(QByteArrayView*)rhs);
	return __r;
}
int QStringView_OperatorGreaterOrEqual3(void* self, void* rhs) {
	auto&& __r = (*(QStringView*)self) >= (*(QByteArrayView*)rhs);
	return __r;
}
int QStringView_OperatorEqual4(void* self, void* rhs) {
	auto&& __r = (*(QStringView*)self) == (*(QByteArray*)rhs);
	return __r;
}
int QStringView_OperatorNotEqual4(void* self, void* rhs) {
	auto&& __r = (*(QStringView*)self) != (*(QByteArray*)rhs);
	return __r;
}
int QStringView_OperatorLesser4(void* self, void* rhs) {
	auto&& __r = (*(QStringView*)self) < (*(QByteArray*)rhs);
	return __r;
}
int QStringView_OperatorGreater4(void* self, void* rhs) {
	auto&& __r = (*(QStringView*)self) > (*(QByteArray*)rhs);
	return __r;
}
int QStringView_OperatorLesserOrEqual4(void* self, void* rhs) {
	auto&& __r = (*(QStringView*)self) <= (*(QByteArray*)rhs);
	return __r;
}
int QStringView_OperatorGreaterOrEqual4(void* self, void* rhs) {
	auto&& __r = (*(QStringView*)self) >= (*(QByteArray*)rhs);
	return __r;
}
char* QStringView_OperatorPlus(void* self, const char* rhs, long long rhs_len, long long* __out_len) {
	QString rhs_qstr = QString::fromUtf8(rhs, rhs_len);
	auto&& __r = (*(QStringView*)self) + (rhs_qstr);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QStringView_delete(void* self) { delete (QStringView*)self; }

}

#include <QtCore/qabstractanimation.h>
#include <QtCore/qobject.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qanimationgroup.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qvariant.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qabstractanimation.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QAbstractAnimation__Go : public QAbstractAnimation {
public:
	long long __goId;
	unsigned char __goOn[12] = {};
	template<typename... A> QAbstractAnimation__Go(long long id, A&&... a) : QAbstractAnimation(std::forward<A>(a)...), __goId(id) {}
	using QAbstractAnimation::sender;
	using QAbstractAnimation::senderSignalIndex;
	using QAbstractAnimation::receivers;
	using QAbstractAnimation::isSignalConnected;
	bool cppgoBase_event(QEvent* __a0) { return QAbstractAnimation::event(__a0); }
	void cppgoBase_updateState(int __a0, int __a1) { QAbstractAnimation::updateState((QAbstractAnimation::State)__a0, (QAbstractAnimation::State)__a1); }
	void cppgoBase_updateDirection(int __a0) { QAbstractAnimation::updateDirection((QAbstractAnimation::Direction)__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QAbstractAnimation::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QAbstractAnimation::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QAbstractAnimation::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QAbstractAnimation::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QAbstractAnimation::disconnectNotify(__a0); }
	int duration() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8434675112194400115LL, &__h); if (__h) return __r; }
		return int{};
	}
	void updateCurrentTime(int currentTime) override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, -1342418961870150968LL, currentTime, &__h); if (__h) return; }
		return;
	}
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QAbstractAnimation::metaObject();
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QAbstractAnimation::event(event);
	}
	void updateState(QAbstractAnimation::State newState, QAbstractAnimation::State oldState) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_int_int) { ((void(*)(long long, long long, int, int, int*))cppgo_virt_void_int_int)(__goId, 7692482089134926569LL, (int)(newState), (int)(oldState), &__h); if (__h) return; }
		QAbstractAnimation::updateState(newState, oldState);
	}
	void updateDirection(QAbstractAnimation::Direction direction) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, -3231888725407197643LL, (int)(direction), &__h); if (__h) return; }
		QAbstractAnimation::updateDirection(direction);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QAbstractAnimation::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QAbstractAnimation::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QAbstractAnimation::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QAbstractAnimation::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QAbstractAnimation::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QAbstractAnimation::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QAbstractAnimation;
static std::mutex __connsMu_QAbstractAnimation;

extern "C" {

void* QAbstractAnimation_base(void* self) { return (void*)static_cast<QObject*>((QAbstractAnimation*)self); }
void QAbstractAnimation__set_override(void* self, int idx) { ((QAbstractAnimation__Go*)self)->__goOn[idx] = 1; }
void* QAbstractAnimation_fromQObject(void* o) { return (void*)qobject_cast<QAbstractAnimation*>((QObject*)o); }
char* QAbstractAnimation_tr(const char* s, long long* __out_len) {
	auto&& __r = QAbstractAnimation::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QAbstractAnimation_state(void* self) {
	auto&& __r = ((QAbstractAnimation*)self)->state();
	return (int)(__r);
}
void* QAbstractAnimation_group(void* self) {
	auto&& __r = ((QAbstractAnimation*)self)->group();
	return (void*)(__r);
}
int QAbstractAnimation_direction(void* self) {
	auto&& __r = ((QAbstractAnimation*)self)->direction();
	return (int)(__r);
}
void QAbstractAnimation_setDirection(void* self, int direction) {
	((QAbstractAnimation*)self)->setDirection((QAbstractAnimation::Direction)(direction));
}
int QAbstractAnimation_currentTime(void* self) {
	auto&& __r = ((QAbstractAnimation*)self)->currentTime();
	return __r;
}
int QAbstractAnimation_currentLoopTime(void* self) {
	auto&& __r = ((QAbstractAnimation*)self)->currentLoopTime();
	return __r;
}
int QAbstractAnimation_loopCount(void* self) {
	auto&& __r = ((QAbstractAnimation*)self)->loopCount();
	return __r;
}
void QAbstractAnimation_setLoopCount(void* self, int loopCount) {
	((QAbstractAnimation*)self)->setLoopCount(loopCount);
}
int QAbstractAnimation_currentLoop(void* self) {
	auto&& __r = ((QAbstractAnimation*)self)->currentLoop();
	return __r;
}
int QAbstractAnimation_duration(void* self) {
	auto&& __r = ((QAbstractAnimation*)self)->duration();
	return __r;
}
int QAbstractAnimation_totalDuration(void* self) {
	auto&& __r = ((QAbstractAnimation*)self)->totalDuration();
	return __r;
}
void QAbstractAnimation_finished(void* self) {
	((QAbstractAnimation*)self)->finished();
}
void QAbstractAnimation_stateChanged(void* self, int newState, int oldState) {
	((QAbstractAnimation*)self)->stateChanged((QAbstractAnimation::State)(newState), (QAbstractAnimation::State)(oldState));
}
void QAbstractAnimation_currentLoopChanged(void* self, int currentLoop) {
	((QAbstractAnimation*)self)->currentLoopChanged(currentLoop);
}
void QAbstractAnimation_directionChanged(void* self, int arg1) {
	((QAbstractAnimation*)self)->directionChanged((QAbstractAnimation::Direction)(arg1));
}
void QAbstractAnimation_start(void* self) {
	((QAbstractAnimation*)self)->start();
}
void QAbstractAnimation_pause(void* self) {
	((QAbstractAnimation*)self)->pause();
}
void QAbstractAnimation_resume(void* self) {
	((QAbstractAnimation*)self)->resume();
}
void QAbstractAnimation_setPaused(void* self, int arg1) {
	((QAbstractAnimation*)self)->setPaused(arg1);
}
void QAbstractAnimation_stop(void* self) {
	((QAbstractAnimation*)self)->stop();
}
void QAbstractAnimation_setCurrentTime(void* self, int msecs) {
	((QAbstractAnimation*)self)->setCurrentTime(msecs);
}
void* QAbstractAnimation_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QAbstractAnimation::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QAbstractAnimation_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QAbstractAnimation::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QAbstractAnimation_sender(void* self) {
	auto&& __r = ((QAbstractAnimation__Go*)self)->sender();
	return (void*)(__r);
}
int QAbstractAnimation_senderSignalIndex(void* self) {
	auto&& __r = ((QAbstractAnimation__Go*)self)->senderSignalIndex();
	return __r;
}
int QAbstractAnimation_receivers(void* self, const char* signal) {
	auto&& __r = ((QAbstractAnimation__Go*)self)->receivers(signal);
	return __r;
}
int QAbstractAnimation_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QAbstractAnimation__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QAbstractAnimation_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QAbstractAnimation::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QAbstractAnimation_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QAbstractAnimation::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QAbstractAnimation_start2(void* self, int policy) {
	((QAbstractAnimation*)self)->start((QAbstractAnimation::DeletionPolicy)(policy));
}
void* QAbstractAnimation_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QAbstractAnimation::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QAbstractAnimation_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QAbstractAnimation::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QAbstractAnimation_metaObject_default(void* self) {
	auto&& __r = ((QAbstractAnimation*)self)->QAbstractAnimation::metaObject();
	return (void*)(__r);
}
int QAbstractAnimation_event_default(void* self, void* event) {
	auto&& __r = ((QAbstractAnimation__Go*)self)->cppgoBase_event((QEvent*)event);
	return __r;
}
void QAbstractAnimation_updateState_default(void* self, int newState, int oldState) {
	((QAbstractAnimation__Go*)self)->cppgoBase_updateState(newState, oldState);
}
void QAbstractAnimation_updateDirection_default(void* self, int direction) {
	((QAbstractAnimation__Go*)self)->cppgoBase_updateDirection(direction);
}
int QAbstractAnimation_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QAbstractAnimation*)self)->QAbstractAnimation::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QAbstractAnimation_timerEvent_default(void* self, void* event) {
	((QAbstractAnimation__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QAbstractAnimation_childEvent_default(void* self, void* event) {
	((QAbstractAnimation__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QAbstractAnimation_customEvent_default(void* self, void* event) {
	((QAbstractAnimation__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QAbstractAnimation_connectNotify_default(void* self, void* signal) {
	((QAbstractAnimation__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QAbstractAnimation_disconnectNotify_default(void* self, void* signal) {
	((QAbstractAnimation__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QAbstractAnimation_connect_finished(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractAnimation*)self, &QAbstractAnimation::finished, (QAbstractAnimation*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractAnimation);
	__conns_QAbstractAnimation[__id] = __c;
}
void QAbstractAnimation_connect_stateChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractAnimation*)self, &QAbstractAnimation::stateChanged, (QAbstractAnimation*)self, [__id](QAbstractAnimation::State newState, QAbstractAnimation::State oldState){
		if (cppgo_sig_void_int_int) ((void(*)(long long, int, int))cppgo_sig_void_int_int)(__id, (int)(newState), (int)(oldState));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractAnimation);
	__conns_QAbstractAnimation[__id] = __c;
}
void QAbstractAnimation_connect_currentLoopChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractAnimation*)self, &QAbstractAnimation::currentLoopChanged, (QAbstractAnimation*)self, [__id](int currentLoop){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, currentLoop);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractAnimation);
	__conns_QAbstractAnimation[__id] = __c;
}
void QAbstractAnimation_connect_directionChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractAnimation*)self, &QAbstractAnimation::directionChanged, (QAbstractAnimation*)self, [__id](QAbstractAnimation::Direction arg1){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, (int)(arg1));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractAnimation);
	__conns_QAbstractAnimation[__id] = __c;
}
void QAbstractAnimation_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractAnimation*)self, &QAbstractAnimation::destroyed, (QAbstractAnimation*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractAnimation);
	__conns_QAbstractAnimation[__id] = __c;
}
void QAbstractAnimation_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractAnimation*)self, &QAbstractAnimation::destroyed, (QAbstractAnimation*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractAnimation);
	__conns_QAbstractAnimation[__id] = __c;
}
void QAbstractAnimation_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractAnimation*)self, &QAbstractAnimation::objectNameChanged, (QAbstractAnimation*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractAnimation);
	__conns_QAbstractAnimation[__id] = __c;
}
void QAbstractAnimation__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QAbstractAnimation);
		auto __it = __conns_QAbstractAnimation.find(id);
		if (__it == __conns_QAbstractAnimation.end()) return;
		__c = __it->second;
		__conns_QAbstractAnimation.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QAbstractAnimation__Go_new(long long id, void* parent) { return new QAbstractAnimation__Go(id, (QObject*)parent); }
void QAbstractAnimation_delete(void* self) { delete (QAbstractAnimation*)self; }

}

#include <QtCore/qabstractitemmodel.h>
#include <QtCore/qvariant.h>
#include <QtCore/qnamespace.h>
#include "gen_qpersistentmodelindex.h"

extern "C" {

void* QPersistentModelIndex_new() {
	return new QPersistentModelIndex();
}
void* QPersistentModelIndex_new2(void* index) {
	return new QPersistentModelIndex(*(QModelIndex*)index);
}
void* QPersistentModelIndex_new3(void* other) {
	return new QPersistentModelIndex(*(QPersistentModelIndex*)other);
}
void* QPersistentModelIndex_OperatorAssign(void* self, void* other) {
	auto&& __r = ((QPersistentModelIndex*)self)->operator=(*(QPersistentModelIndex*)other);
	return (void*)&(__r);
}
void QPersistentModelIndex_swap(void* self, void* other) {
	((QPersistentModelIndex*)self)->swap(*(QPersistentModelIndex*)other);
}
void* QPersistentModelIndex_OperatorAssign2(void* self, void* other) {
	auto&& __r = ((QPersistentModelIndex*)self)->operator=(*(QModelIndex*)other);
	return (void*)&(__r);
}
void* QPersistentModelIndex_ToQModelIndex(void* self) {
	auto&& __r = static_cast<QModelIndex>(*(QPersistentModelIndex*)self);
	return new QModelIndex(__r);
}
int QPersistentModelIndex_row(void* self) {
	auto&& __r = ((QPersistentModelIndex*)self)->row();
	return __r;
}
int QPersistentModelIndex_column(void* self) {
	auto&& __r = ((QPersistentModelIndex*)self)->column();
	return __r;
}
unsigned long long QPersistentModelIndex_internalId(void* self) {
	auto&& __r = ((QPersistentModelIndex*)self)->internalId();
	return __r;
}
void* QPersistentModelIndex_parent(void* self) {
	auto&& __r = ((QPersistentModelIndex*)self)->parent();
	return new QModelIndex(__r);
}
void* QPersistentModelIndex_sibling(void* self, int row, int column) {
	auto&& __r = ((QPersistentModelIndex*)self)->sibling(row, column);
	return new QModelIndex(__r);
}
void* QPersistentModelIndex_data(void* self) {
	auto&& __r = ((QPersistentModelIndex*)self)->data();
	return new QVariant(__r);
}
void QPersistentModelIndex_multiData(void* self, void* roleDataSpan) {
	((QPersistentModelIndex*)self)->multiData(*(QModelRoleDataSpan*)roleDataSpan);
}
int QPersistentModelIndex_flags(void* self) {
	auto&& __r = ((QPersistentModelIndex*)self)->flags();
	return (int)(__r);
}
void* QPersistentModelIndex_model(void* self) {
	auto&& __r = ((QPersistentModelIndex*)self)->model();
	return (void*)(__r);
}
int QPersistentModelIndex_isValid(void* self) {
	auto&& __r = ((QPersistentModelIndex*)self)->isValid();
	return __r;
}
void* QPersistentModelIndex_data2(void* self, int role) {
	auto&& __r = ((QPersistentModelIndex*)self)->data(role);
	return new QVariant(__r);
}
int QPersistentModelIndex_OperatorEqual(void* self, void* rhs) {
	auto&& __r = (*(QPersistentModelIndex*)self) == (*(QPersistentModelIndex*)rhs);
	return __r;
}
int QPersistentModelIndex_OperatorNotEqual(void* self, void* rhs) {
	auto&& __r = (*(QPersistentModelIndex*)self) != (*(QPersistentModelIndex*)rhs);
	return __r;
}
int QPersistentModelIndex_OperatorLesser(void* self, void* rhs) {
	auto&& __r = (*(QPersistentModelIndex*)self) < (*(QPersistentModelIndex*)rhs);
	return __r;
}
int QPersistentModelIndex_OperatorGreater(void* self, void* rhs) {
	auto&& __r = (*(QPersistentModelIndex*)self) > (*(QPersistentModelIndex*)rhs);
	return __r;
}
int QPersistentModelIndex_OperatorLesserOrEqual(void* self, void* rhs) {
	auto&& __r = (*(QPersistentModelIndex*)self) <= (*(QPersistentModelIndex*)rhs);
	return __r;
}
int QPersistentModelIndex_OperatorGreaterOrEqual(void* self, void* rhs) {
	auto&& __r = (*(QPersistentModelIndex*)self) >= (*(QPersistentModelIndex*)rhs);
	return __r;
}
int QPersistentModelIndex_OperatorEqual2(void* self, void* rhs) {
	auto&& __r = (*(QPersistentModelIndex*)self) == (*(QModelIndex*)rhs);
	return __r;
}
int QPersistentModelIndex_OperatorNotEqual2(void* self, void* rhs) {
	auto&& __r = (*(QPersistentModelIndex*)self) != (*(QModelIndex*)rhs);
	return __r;
}
int QPersistentModelIndex_OperatorLesser2(void* self, void* rhs) {
	auto&& __r = (*(QPersistentModelIndex*)self) < (*(QModelIndex*)rhs);
	return __r;
}
int QPersistentModelIndex_OperatorGreater2(void* self, void* rhs) {
	auto&& __r = (*(QPersistentModelIndex*)self) > (*(QModelIndex*)rhs);
	return __r;
}
int QPersistentModelIndex_OperatorLesserOrEqual2(void* self, void* rhs) {
	auto&& __r = (*(QPersistentModelIndex*)self) <= (*(QModelIndex*)rhs);
	return __r;
}
int QPersistentModelIndex_OperatorGreaterOrEqual2(void* self, void* rhs) {
	auto&& __r = (*(QPersistentModelIndex*)self) >= (*(QModelIndex*)rhs);
	return __r;
}
void QPersistentModelIndex_delete(void* self) { delete (QPersistentModelIndex*)self; }

}

#include <QtCore/qlocale.h>
#include <QtCore/qstring.h>
#include <QtCore/qstringview.h>
#include <QtCore/qdatetime.h>
#include <QtCore/qcalendar.h>
#include <QtCore/qnamespace.h>
#include "gen_qlocale.h"

extern "C" {

void* QLocale_new() {
	return new QLocale();
}
void* QLocale_new2(const char* name, long long name_len) {
	QString name_qstr = QString::fromUtf8(name, name_len);
	return new QLocale(name_qstr);
}
void* QLocale_new3(void* name) {
	return new QLocale(*(QStringView*)name);
}
void* QLocale_new4(int language, int territory) {
	return new QLocale((QLocale::Language)(language), (QLocale::Country)(territory));
}
void* QLocale_new5(int language) {
	return new QLocale((QLocale::Language)(language));
}
void* QLocale_new6(void* other) {
	return new QLocale(*(QLocale*)other);
}
void* QLocale_new7(int language, int script) {
	return new QLocale((QLocale::Language)(language), (QLocale::Script)(script));
}
void* QLocale_new8(int language, int script, int territory) {
	return new QLocale((QLocale::Language)(language), (QLocale::Script)(script), (QLocale::Country)(territory));
}
void* QLocale_OperatorAssign(void* self, void* other) {
	auto&& __r = ((QLocale*)self)->operator=(*(QLocale*)other);
	return (void*)&(__r);
}
void QLocale_swap(void* self, void* other) {
	((QLocale*)self)->swap(*(QLocale*)other);
}
int QLocale_language(void* self) {
	auto&& __r = ((QLocale*)self)->language();
	return (int)(__r);
}
int QLocale_script(void* self) {
	auto&& __r = ((QLocale*)self)->script();
	return (int)(__r);
}
int QLocale_territory(void* self) {
	auto&& __r = ((QLocale*)self)->territory();
	return (int)(__r);
}
int QLocale_country(void* self) {
	auto&& __r = ((QLocale*)self)->country();
	return (int)(__r);
}
char* QLocale_name(void* self, long long* __out_len) {
	auto&& __r = ((QLocale*)self)->name();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLocale_bcp47Name(void* self, long long* __out_len) {
	auto&& __r = ((QLocale*)self)->bcp47Name();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLocale_nativeLanguageName(void* self, long long* __out_len) {
	auto&& __r = ((QLocale*)self)->nativeLanguageName();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLocale_nativeTerritoryName(void* self, long long* __out_len) {
	auto&& __r = ((QLocale*)self)->nativeTerritoryName();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLocale_nativeCountryName(void* self, long long* __out_len) {
	auto&& __r = ((QLocale*)self)->nativeCountryName();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
short QLocale_toShort(void* self, const char* s, long long s_len) {
	QString s_qstr = QString::fromUtf8(s, s_len);
	auto&& __r = ((QLocale*)self)->toShort(s_qstr);
	return __r;
}
unsigned short QLocale_toUShort(void* self, const char* s, long long s_len) {
	QString s_qstr = QString::fromUtf8(s, s_len);
	auto&& __r = ((QLocale*)self)->toUShort(s_qstr);
	return __r;
}
int QLocale_toInt(void* self, const char* s, long long s_len) {
	QString s_qstr = QString::fromUtf8(s, s_len);
	auto&& __r = ((QLocale*)self)->toInt(s_qstr);
	return __r;
}
unsigned int QLocale_toUInt(void* self, const char* s, long long s_len) {
	QString s_qstr = QString::fromUtf8(s, s_len);
	auto&& __r = ((QLocale*)self)->toUInt(s_qstr);
	return __r;
}
long QLocale_toLong(void* self, const char* s, long long s_len) {
	QString s_qstr = QString::fromUtf8(s, s_len);
	auto&& __r = ((QLocale*)self)->toLong(s_qstr);
	return __r;
}
unsigned long QLocale_toULong(void* self, const char* s, long long s_len) {
	QString s_qstr = QString::fromUtf8(s, s_len);
	auto&& __r = ((QLocale*)self)->toULong(s_qstr);
	return __r;
}
long long QLocale_toLongLong(void* self, const char* s, long long s_len) {
	QString s_qstr = QString::fromUtf8(s, s_len);
	auto&& __r = ((QLocale*)self)->toLongLong(s_qstr);
	return __r;
}
unsigned long long QLocale_toULongLong(void* self, const char* s, long long s_len) {
	QString s_qstr = QString::fromUtf8(s, s_len);
	auto&& __r = ((QLocale*)self)->toULongLong(s_qstr);
	return __r;
}
float QLocale_toFloat(void* self, const char* s, long long s_len) {
	QString s_qstr = QString::fromUtf8(s, s_len);
	auto&& __r = ((QLocale*)self)->toFloat(s_qstr);
	return __r;
}
double QLocale_toDouble(void* self, const char* s, long long s_len) {
	QString s_qstr = QString::fromUtf8(s, s_len);
	auto&& __r = ((QLocale*)self)->toDouble(s_qstr);
	return __r;
}
short QLocale_toShort2(void* self, void* s) {
	auto&& __r = ((QLocale*)self)->toShort(*(QStringView*)s);
	return __r;
}
unsigned short QLocale_toUShort2(void* self, void* s) {
	auto&& __r = ((QLocale*)self)->toUShort(*(QStringView*)s);
	return __r;
}
int QLocale_toInt2(void* self, void* s) {
	auto&& __r = ((QLocale*)self)->toInt(*(QStringView*)s);
	return __r;
}
unsigned int QLocale_toUInt2(void* self, void* s) {
	auto&& __r = ((QLocale*)self)->toUInt(*(QStringView*)s);
	return __r;
}
long QLocale_toLong2(void* self, void* s) {
	auto&& __r = ((QLocale*)self)->toLong(*(QStringView*)s);
	return __r;
}
unsigned long QLocale_toULong2(void* self, void* s) {
	auto&& __r = ((QLocale*)self)->toULong(*(QStringView*)s);
	return __r;
}
long long QLocale_toLongLong2(void* self, void* s) {
	auto&& __r = ((QLocale*)self)->toLongLong(*(QStringView*)s);
	return __r;
}
unsigned long long QLocale_toULongLong2(void* self, void* s) {
	auto&& __r = ((QLocale*)self)->toULongLong(*(QStringView*)s);
	return __r;
}
float QLocale_toFloat2(void* self, void* s) {
	auto&& __r = ((QLocale*)self)->toFloat(*(QStringView*)s);
	return __r;
}
double QLocale_toDouble2(void* self, void* s) {
	auto&& __r = ((QLocale*)self)->toDouble(*(QStringView*)s);
	return __r;
}
char* QLocale_toString(void* self, long long i, long long* __out_len) {
	auto&& __r = ((QLocale*)self)->toString(i);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLocale_toString2(void* self, unsigned long long i, long long* __out_len) {
	auto&& __r = ((QLocale*)self)->toString(i);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLocale_toString3(void* self, long i, long long* __out_len) {
	auto&& __r = ((QLocale*)self)->toString(i);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLocale_toString4(void* self, unsigned long i, long long* __out_len) {
	auto&& __r = ((QLocale*)self)->toString(i);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLocale_toString5(void* self, short i, long long* __out_len) {
	auto&& __r = ((QLocale*)self)->toString(i);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLocale_toString6(void* self, unsigned short i, long long* __out_len) {
	auto&& __r = ((QLocale*)self)->toString(i);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLocale_toString7(void* self, int i, long long* __out_len) {
	auto&& __r = ((QLocale*)self)->toString(i);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLocale_toString8(void* self, unsigned int i, long long* __out_len) {
	auto&& __r = ((QLocale*)self)->toString(i);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLocale_toString9(void* self, long long number, int fieldWidth, unsigned int fillChar, long long* __out_len) {
	auto&& __r = ((QLocale*)self)->toString(number, fieldWidth, fillChar);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLocale_toString10(void* self, unsigned long long number, int fieldWidth, unsigned int fillChar, long long* __out_len) {
	auto&& __r = ((QLocale*)self)->toString(number, fieldWidth, fillChar);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLocale_toString11(void* self, int number, int fieldWidth, unsigned int fillChar, long long* __out_len) {
	auto&& __r = ((QLocale*)self)->toString(number, fieldWidth, fillChar);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLocale_toString12(void* self, unsigned int number, int fieldWidth, unsigned int fillChar, long long* __out_len) {
	auto&& __r = ((QLocale*)self)->toString(number, fieldWidth, fillChar);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLocale_toString13(void* self, long number, int fieldWidth, unsigned int fillChar, long long* __out_len) {
	auto&& __r = ((QLocale*)self)->toString(number, fieldWidth, fillChar);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLocale_toString14(void* self, unsigned long number, int fieldWidth, unsigned int fillChar, long long* __out_len) {
	auto&& __r = ((QLocale*)self)->toString(number, fieldWidth, fillChar);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLocale_toString15(void* self, short number, int fieldWidth, unsigned int fillChar, long long* __out_len) {
	auto&& __r = ((QLocale*)self)->toString(number, fieldWidth, fillChar);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLocale_toString16(void* self, unsigned short number, int fieldWidth, unsigned int fillChar, long long* __out_len) {
	auto&& __r = ((QLocale*)self)->toString(number, fieldWidth, fillChar);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLocale_toString17(void* self, double f, long long* __out_len) {
	auto&& __r = ((QLocale*)self)->toString(f);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLocale_toString18(void* self, float f, long long* __out_len) {
	auto&& __r = ((QLocale*)self)->toString(f);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLocale_toString19(void* self, void* date, const char* format, long long format_len, long long* __out_len) {
	QString format_qstr = QString::fromUtf8(format, format_len);
	auto&& __r = ((QLocale*)self)->toString(*(QDate*)date, format_qstr);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLocale_toString20(void* self, void* time, const char* format, long long format_len, long long* __out_len) {
	QString format_qstr = QString::fromUtf8(format, format_len);
	auto&& __r = ((QLocale*)self)->toString(*(QTime*)time, format_qstr);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLocale_toString21(void* self, void* dateTime, const char* format, long long format_len, long long* __out_len) {
	QString format_qstr = QString::fromUtf8(format, format_len);
	auto&& __r = ((QLocale*)self)->toString(*(QDateTime*)dateTime, format_qstr);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLocale_toString22(void* self, void* date, void* format, long long* __out_len) {
	auto&& __r = ((QLocale*)self)->toString(*(QDate*)date, *(QStringView*)format);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLocale_toString23(void* self, void* time, void* format, long long* __out_len) {
	auto&& __r = ((QLocale*)self)->toString(*(QTime*)time, *(QStringView*)format);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLocale_toString24(void* self, void* dateTime, void* format, long long* __out_len) {
	auto&& __r = ((QLocale*)self)->toString(*(QDateTime*)dateTime, *(QStringView*)format);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLocale_toString25(void* self, void* date, long long* __out_len) {
	auto&& __r = ((QLocale*)self)->toString(*(QDate*)date);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLocale_toString26(void* self, void* time, long long* __out_len) {
	auto&& __r = ((QLocale*)self)->toString(*(QTime*)time);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLocale_toString27(void* self, void* dateTime, long long* __out_len) {
	auto&& __r = ((QLocale*)self)->toString(*(QDateTime*)dateTime);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLocale_toString28(void* self, void* date, void* format, void* cal, long long* __out_len) {
	auto&& __r = ((QLocale*)self)->toString(*(QDate*)date, *(QStringView*)format, *(QCalendar*)cal);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLocale_toString29(void* self, void* date, int format, void* cal, long long* __out_len) {
	auto&& __r = ((QLocale*)self)->toString(*(QDate*)date, (QLocale::FormatType)(format), *(QCalendar*)cal);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLocale_toString30(void* self, void* dateTime, int format, void* cal, long long* __out_len) {
	auto&& __r = ((QLocale*)self)->toString(*(QDateTime*)dateTime, (QLocale::FormatType)(format), *(QCalendar*)cal);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLocale_toString31(void* self, void* dateTime, void* format, void* cal, long long* __out_len) {
	auto&& __r = ((QLocale*)self)->toString(*(QDateTime*)dateTime, *(QStringView*)format, *(QCalendar*)cal);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLocale_dateFormat(void* self, long long* __out_len) {
	auto&& __r = ((QLocale*)self)->dateFormat();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLocale_timeFormat(void* self, long long* __out_len) {
	auto&& __r = ((QLocale*)self)->timeFormat();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLocale_dateTimeFormat(void* self, long long* __out_len) {
	auto&& __r = ((QLocale*)self)->dateTimeFormat();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QLocale_toTime(void* self, const char* string_, long long string__len) {
	QString string__qstr = QString::fromUtf8(string_, string__len);
	auto&& __r = ((QLocale*)self)->toTime(string__qstr);
	return new QTime(__r);
}
void* QLocale_toTime2(void* self, const char* string_, long long string__len, const char* format, long long format_len) {
	QString string__qstr = QString::fromUtf8(string_, string__len);
	QString format_qstr = QString::fromUtf8(format, format_len);
	auto&& __r = ((QLocale*)self)->toTime(string__qstr, format_qstr);
	return new QTime(__r);
}
void* QLocale_toDate(void* self, const char* string_, long long string__len) {
	QString string__qstr = QString::fromUtf8(string_, string__len);
	auto&& __r = ((QLocale*)self)->toDate(string__qstr);
	return new QDate(__r);
}
void* QLocale_toDate2(void* self, const char* string_, long long string__len, const char* format, long long format_len) {
	QString string__qstr = QString::fromUtf8(string_, string__len);
	QString format_qstr = QString::fromUtf8(format, format_len);
	auto&& __r = ((QLocale*)self)->toDate(string__qstr, format_qstr);
	return new QDate(__r);
}
void* QLocale_toDateTime(void* self, const char* string_, long long string__len) {
	QString string__qstr = QString::fromUtf8(string_, string__len);
	auto&& __r = ((QLocale*)self)->toDateTime(string__qstr);
	return new QDateTime(__r);
}
void* QLocale_toDateTime2(void* self, const char* string_, long long string__len, const char* format, long long format_len) {
	QString string__qstr = QString::fromUtf8(string_, string__len);
	QString format_qstr = QString::fromUtf8(format, format_len);
	auto&& __r = ((QLocale*)self)->toDateTime(string__qstr, format_qstr);
	return new QDateTime(__r);
}
void* QLocale_toDate3(void* self, const char* string_, long long string__len, int format, void* cal) {
	QString string__qstr = QString::fromUtf8(string_, string__len);
	auto&& __r = ((QLocale*)self)->toDate(string__qstr, (QLocale::FormatType)(format), *(QCalendar*)cal);
	return new QDate(__r);
}
void* QLocale_toDate4(void* self, const char* string_, long long string__len, const char* format, long long format_len, void* cal) {
	QString string__qstr = QString::fromUtf8(string_, string__len);
	QString format_qstr = QString::fromUtf8(format, format_len);
	auto&& __r = ((QLocale*)self)->toDate(string__qstr, format_qstr, *(QCalendar*)cal);
	return new QDate(__r);
}
void* QLocale_toDateTime3(void* self, const char* string_, long long string__len, int format, void* cal) {
	QString string__qstr = QString::fromUtf8(string_, string__len);
	auto&& __r = ((QLocale*)self)->toDateTime(string__qstr, (QLocale::FormatType)(format), *(QCalendar*)cal);
	return new QDateTime(__r);
}
void* QLocale_toDateTime4(void* self, const char* string_, long long string__len, const char* format, long long format_len, void* cal) {
	QString string__qstr = QString::fromUtf8(string_, string__len);
	QString format_qstr = QString::fromUtf8(format, format_len);
	auto&& __r = ((QLocale*)self)->toDateTime(string__qstr, format_qstr, *(QCalendar*)cal);
	return new QDateTime(__r);
}
char* QLocale_decimalPoint(void* self, long long* __out_len) {
	auto&& __r = ((QLocale*)self)->decimalPoint();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLocale_groupSeparator(void* self, long long* __out_len) {
	auto&& __r = ((QLocale*)self)->groupSeparator();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLocale_percent(void* self, long long* __out_len) {
	auto&& __r = ((QLocale*)self)->percent();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLocale_zeroDigit(void* self, long long* __out_len) {
	auto&& __r = ((QLocale*)self)->zeroDigit();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLocale_negativeSign(void* self, long long* __out_len) {
	auto&& __r = ((QLocale*)self)->negativeSign();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLocale_positiveSign(void* self, long long* __out_len) {
	auto&& __r = ((QLocale*)self)->positiveSign();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLocale_exponential(void* self, long long* __out_len) {
	auto&& __r = ((QLocale*)self)->exponential();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLocale_monthName(void* self, int arg1, long long* __out_len) {
	auto&& __r = ((QLocale*)self)->monthName(arg1);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLocale_standaloneMonthName(void* self, int arg1, long long* __out_len) {
	auto&& __r = ((QLocale*)self)->standaloneMonthName(arg1);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLocale_dayName(void* self, int arg1, long long* __out_len) {
	auto&& __r = ((QLocale*)self)->dayName(arg1);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLocale_standaloneDayName(void* self, int arg1, long long* __out_len) {
	auto&& __r = ((QLocale*)self)->standaloneDayName(arg1);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QLocale_firstDayOfWeek(void* self) {
	auto&& __r = ((QLocale*)self)->firstDayOfWeek();
	return (int)(__r);
}
int* QLocale_weekdays(void* self, long long* __out_len) {
	auto&& __r = ((QLocale*)self)->weekdays();
	*__out_len = __r.size();
	int* __arr = (int*)malloc(sizeof(int) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (int)(__e);
	return __arr;
}
char* QLocale_amText(void* self, long long* __out_len) {
	auto&& __r = ((QLocale*)self)->amText();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLocale_pmText(void* self, long long* __out_len) {
	auto&& __r = ((QLocale*)self)->pmText();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QLocale_measurementSystem(void* self) {
	auto&& __r = ((QLocale*)self)->measurementSystem();
	return (int)(__r);
}
void* QLocale_collation(void* self) {
	auto&& __r = ((QLocale*)self)->collation();
	return new QLocale(__r);
}
int QLocale_textDirection(void* self) {
	auto&& __r = ((QLocale*)self)->textDirection();
	return (int)(__r);
}
char* QLocale_toUpper(void* self, const char* str, long long str_len, long long* __out_len) {
	QString str_qstr = QString::fromUtf8(str, str_len);
	auto&& __r = ((QLocale*)self)->toUpper(str_qstr);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLocale_toLower(void* self, const char* str, long long str_len, long long* __out_len) {
	QString str_qstr = QString::fromUtf8(str, str_len);
	auto&& __r = ((QLocale*)self)->toLower(str_qstr);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLocale_currencySymbol(void* self, long long* __out_len) {
	auto&& __r = ((QLocale*)self)->currencySymbol();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLocale_toCurrencyString(void* self, long long arg1, long long* __out_len) {
	auto&& __r = ((QLocale*)self)->toCurrencyString(arg1);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLocale_toCurrencyString2(void* self, unsigned long long arg1, long long* __out_len) {
	auto&& __r = ((QLocale*)self)->toCurrencyString(arg1);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLocale_toCurrencyString3(void* self, short i, long long* __out_len) {
	auto&& __r = ((QLocale*)self)->toCurrencyString(i);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLocale_toCurrencyString4(void* self, unsigned short i, long long* __out_len) {
	auto&& __r = ((QLocale*)self)->toCurrencyString(i);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLocale_toCurrencyString5(void* self, int i, long long* __out_len) {
	auto&& __r = ((QLocale*)self)->toCurrencyString(i);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLocale_toCurrencyString6(void* self, unsigned int i, long long* __out_len) {
	auto&& __r = ((QLocale*)self)->toCurrencyString(i);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLocale_toCurrencyString7(void* self, double arg1, long long* __out_len) {
	auto&& __r = ((QLocale*)self)->toCurrencyString(arg1);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLocale_toCurrencyString8(void* self, float i, long long* __out_len) {
	auto&& __r = ((QLocale*)self)->toCurrencyString(i);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLocale_formattedDataSize(void* self, long long bytes, long long* __out_len) {
	auto&& __r = ((QLocale*)self)->formattedDataSize(bytes);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char** QLocale_uiLanguages(void* self, long long* __out_len) {
	auto&& __r = ((QLocale*)self)->uiLanguages();
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
char* QLocale_languageToCode(int language, long long* __out_len) {
	auto&& __r = QLocale::languageToCode((QLocale::Language)(language));
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QLocale_codeToLanguage(void* languageCode) {
	auto&& __r = QLocale::codeToLanguage(*(QStringView*)languageCode);
	return (int)(__r);
}
char* QLocale_territoryToCode(int territory, long long* __out_len) {
	auto&& __r = QLocale::territoryToCode((QLocale::Country)(territory));
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QLocale_codeToTerritory(void* territoryCode) {
	auto&& __r = QLocale::codeToTerritory(*(QStringView*)territoryCode);
	return (int)(__r);
}
char* QLocale_countryToCode(int country, long long* __out_len) {
	auto&& __r = QLocale::countryToCode((QLocale::Country)(country));
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QLocale_codeToCountry(void* countryCode) {
	auto&& __r = QLocale::codeToCountry(*(QStringView*)countryCode);
	return (int)(__r);
}
char* QLocale_scriptToCode(int script, long long* __out_len) {
	auto&& __r = QLocale::scriptToCode((QLocale::Script)(script));
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QLocale_codeToScript(void* scriptCode) {
	auto&& __r = QLocale::codeToScript(*(QStringView*)scriptCode);
	return (int)(__r);
}
char* QLocale_languageToString(int language, long long* __out_len) {
	auto&& __r = QLocale::languageToString((QLocale::Language)(language));
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLocale_territoryToString(int territory, long long* __out_len) {
	auto&& __r = QLocale::territoryToString((QLocale::Country)(territory));
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLocale_countryToString(int country, long long* __out_len) {
	auto&& __r = QLocale::countryToString((QLocale::Country)(country));
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLocale_scriptToString(int script, long long* __out_len) {
	auto&& __r = QLocale::scriptToString((QLocale::Script)(script));
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QLocale_setDefault(void* locale) {
	QLocale::setDefault(*(QLocale*)locale);
}
void* QLocale_c() {
	auto&& __r = QLocale::c();
	return new QLocale(__r);
}
void* QLocale_system() {
	auto&& __r = QLocale::system();
	return new QLocale(__r);
}
void** QLocale_matchingLocales(int language, int script, int territory, long long* __out_len) {
	auto&& __r = QLocale::matchingLocales((QLocale::Language)(language), (QLocale::Script)(script), (QLocale::Country)(territory));
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QLocale(__e);
	return __arr;
}
int* QLocale_countriesForLanguage(int lang, long long* __out_len) {
	auto&& __r = QLocale::countriesForLanguage((QLocale::Language)(lang));
	*__out_len = __r.size();
	int* __arr = (int*)malloc(sizeof(int) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (int)(__e);
	return __arr;
}
void QLocale_setNumberOptions(void* self, int options) {
	((QLocale*)self)->setNumberOptions(QFlags<QLocale::NumberOption>(options));
}
int QLocale_numberOptions(void* self) {
	auto&& __r = ((QLocale*)self)->numberOptions();
	return (int)(__r);
}
char* QLocale_quoteString(void* self, const char* str, long long str_len, long long* __out_len) {
	QString str_qstr = QString::fromUtf8(str, str_len);
	auto&& __r = ((QLocale*)self)->quoteString(str_qstr);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLocale_quoteString2(void* self, void* str, long long* __out_len) {
	auto&& __r = ((QLocale*)self)->quoteString(*(QStringView*)str);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLocale_createSeparatedList(void* self, char** strl, long long strl_len, long long* __out_len) {
	QList<QString> strl_list; strl_list.reserve(strl_len);
	for (long long __i = 0; __i < strl_len; __i++) strl_list.append(QString::fromUtf8(strl[__i]));
	auto&& __r = ((QLocale*)self)->createSeparatedList(strl_list);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLocale_name2(void* self, int separator, long long* __out_len) {
	auto&& __r = ((QLocale*)self)->name((QLocale::TagSeparator)(separator));
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLocale_bcp47Name2(void* self, int separator, long long* __out_len) {
	auto&& __r = ((QLocale*)self)->bcp47Name((QLocale::TagSeparator)(separator));
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
short QLocale_toShort3(void* self, const char* s, long long s_len, int* ok) {
	QString s_qstr = QString::fromUtf8(s, s_len);
	bool ok_local{};
	auto&& __r = ((QLocale*)self)->toShort(s_qstr, &ok_local);
	*ok = ok_local;
	return __r;
}
unsigned short QLocale_toUShort3(void* self, const char* s, long long s_len, int* ok) {
	QString s_qstr = QString::fromUtf8(s, s_len);
	bool ok_local{};
	auto&& __r = ((QLocale*)self)->toUShort(s_qstr, &ok_local);
	*ok = ok_local;
	return __r;
}
int QLocale_toInt3(void* self, const char* s, long long s_len, int* ok) {
	QString s_qstr = QString::fromUtf8(s, s_len);
	bool ok_local{};
	auto&& __r = ((QLocale*)self)->toInt(s_qstr, &ok_local);
	*ok = ok_local;
	return __r;
}
unsigned int QLocale_toUInt3(void* self, const char* s, long long s_len, int* ok) {
	QString s_qstr = QString::fromUtf8(s, s_len);
	bool ok_local{};
	auto&& __r = ((QLocale*)self)->toUInt(s_qstr, &ok_local);
	*ok = ok_local;
	return __r;
}
long QLocale_toLong3(void* self, const char* s, long long s_len, int* ok) {
	QString s_qstr = QString::fromUtf8(s, s_len);
	bool ok_local{};
	auto&& __r = ((QLocale*)self)->toLong(s_qstr, &ok_local);
	*ok = ok_local;
	return __r;
}
unsigned long QLocale_toULong3(void* self, const char* s, long long s_len, int* ok) {
	QString s_qstr = QString::fromUtf8(s, s_len);
	bool ok_local{};
	auto&& __r = ((QLocale*)self)->toULong(s_qstr, &ok_local);
	*ok = ok_local;
	return __r;
}
long long QLocale_toLongLong3(void* self, const char* s, long long s_len, int* ok) {
	QString s_qstr = QString::fromUtf8(s, s_len);
	bool ok_local{};
	auto&& __r = ((QLocale*)self)->toLongLong(s_qstr, &ok_local);
	*ok = ok_local;
	return __r;
}
unsigned long long QLocale_toULongLong3(void* self, const char* s, long long s_len, int* ok) {
	QString s_qstr = QString::fromUtf8(s, s_len);
	bool ok_local{};
	auto&& __r = ((QLocale*)self)->toULongLong(s_qstr, &ok_local);
	*ok = ok_local;
	return __r;
}
float QLocale_toFloat3(void* self, const char* s, long long s_len, int* ok) {
	QString s_qstr = QString::fromUtf8(s, s_len);
	bool ok_local{};
	auto&& __r = ((QLocale*)self)->toFloat(s_qstr, &ok_local);
	*ok = ok_local;
	return __r;
}
double QLocale_toDouble3(void* self, const char* s, long long s_len, int* ok) {
	QString s_qstr = QString::fromUtf8(s, s_len);
	bool ok_local{};
	auto&& __r = ((QLocale*)self)->toDouble(s_qstr, &ok_local);
	*ok = ok_local;
	return __r;
}
short QLocale_toShort4(void* self, void* s, int* ok) {
	bool ok_local{};
	auto&& __r = ((QLocale*)self)->toShort(*(QStringView*)s, &ok_local);
	*ok = ok_local;
	return __r;
}
unsigned short QLocale_toUShort4(void* self, void* s, int* ok) {
	bool ok_local{};
	auto&& __r = ((QLocale*)self)->toUShort(*(QStringView*)s, &ok_local);
	*ok = ok_local;
	return __r;
}
int QLocale_toInt4(void* self, void* s, int* ok) {
	bool ok_local{};
	auto&& __r = ((QLocale*)self)->toInt(*(QStringView*)s, &ok_local);
	*ok = ok_local;
	return __r;
}
unsigned int QLocale_toUInt4(void* self, void* s, int* ok) {
	bool ok_local{};
	auto&& __r = ((QLocale*)self)->toUInt(*(QStringView*)s, &ok_local);
	*ok = ok_local;
	return __r;
}
long QLocale_toLong4(void* self, void* s, int* ok) {
	bool ok_local{};
	auto&& __r = ((QLocale*)self)->toLong(*(QStringView*)s, &ok_local);
	*ok = ok_local;
	return __r;
}
unsigned long QLocale_toULong4(void* self, void* s, int* ok) {
	bool ok_local{};
	auto&& __r = ((QLocale*)self)->toULong(*(QStringView*)s, &ok_local);
	*ok = ok_local;
	return __r;
}
long long QLocale_toLongLong4(void* self, void* s, int* ok) {
	bool ok_local{};
	auto&& __r = ((QLocale*)self)->toLongLong(*(QStringView*)s, &ok_local);
	*ok = ok_local;
	return __r;
}
unsigned long long QLocale_toULongLong4(void* self, void* s, int* ok) {
	bool ok_local{};
	auto&& __r = ((QLocale*)self)->toULongLong(*(QStringView*)s, &ok_local);
	*ok = ok_local;
	return __r;
}
float QLocale_toFloat4(void* self, void* s, int* ok) {
	bool ok_local{};
	auto&& __r = ((QLocale*)self)->toFloat(*(QStringView*)s, &ok_local);
	*ok = ok_local;
	return __r;
}
double QLocale_toDouble4(void* self, void* s, int* ok) {
	bool ok_local{};
	auto&& __r = ((QLocale*)self)->toDouble(*(QStringView*)s, &ok_local);
	*ok = ok_local;
	return __r;
}
char* QLocale_toString32(void* self, double f, char format, long long* __out_len) {
	auto&& __r = ((QLocale*)self)->toString(f, format);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLocale_toString33(void* self, double f, char format, int precision, long long* __out_len) {
	auto&& __r = ((QLocale*)self)->toString(f, format, precision);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLocale_toString34(void* self, float f, char format, long long* __out_len) {
	auto&& __r = ((QLocale*)self)->toString(f, format);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLocale_toString35(void* self, float f, char format, int precision, long long* __out_len) {
	auto&& __r = ((QLocale*)self)->toString(f, format, precision);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLocale_toString36(void* self, void* date, int format, long long* __out_len) {
	auto&& __r = ((QLocale*)self)->toString(*(QDate*)date, (QLocale::FormatType)(format));
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLocale_toString37(void* self, void* time, int format, long long* __out_len) {
	auto&& __r = ((QLocale*)self)->toString(*(QTime*)time, (QLocale::FormatType)(format));
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLocale_toString38(void* self, void* dateTime, int format, long long* __out_len) {
	auto&& __r = ((QLocale*)self)->toString(*(QDateTime*)dateTime, (QLocale::FormatType)(format));
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLocale_dateFormat2(void* self, int format, long long* __out_len) {
	auto&& __r = ((QLocale*)self)->dateFormat((QLocale::FormatType)(format));
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLocale_timeFormat2(void* self, int format, long long* __out_len) {
	auto&& __r = ((QLocale*)self)->timeFormat((QLocale::FormatType)(format));
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLocale_dateTimeFormat2(void* self, int format, long long* __out_len) {
	auto&& __r = ((QLocale*)self)->dateTimeFormat((QLocale::FormatType)(format));
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QLocale_toTime3(void* self, const char* string_, long long string__len, int arg2) {
	QString string__qstr = QString::fromUtf8(string_, string__len);
	auto&& __r = ((QLocale*)self)->toTime(string__qstr, (QLocale::FormatType)(arg2));
	return new QTime(__r);
}
void* QLocale_toDate5(void* self, const char* string_, long long string__len, int arg2) {
	QString string__qstr = QString::fromUtf8(string_, string__len);
	auto&& __r = ((QLocale*)self)->toDate(string__qstr, (QLocale::FormatType)(arg2));
	return new QDate(__r);
}
void* QLocale_toDate6(void* self, const char* string_, long long string__len, int arg2, int baseYear) {
	QString string__qstr = QString::fromUtf8(string_, string__len);
	auto&& __r = ((QLocale*)self)->toDate(string__qstr, (QLocale::FormatType)(arg2), baseYear);
	return new QDate(__r);
}
void* QLocale_toDate7(void* self, const char* string_, long long string__len, const char* format, long long format_len, int baseYear) {
	QString string__qstr = QString::fromUtf8(string_, string__len);
	QString format_qstr = QString::fromUtf8(format, format_len);
	auto&& __r = ((QLocale*)self)->toDate(string__qstr, format_qstr, baseYear);
	return new QDate(__r);
}
void* QLocale_toDateTime5(void* self, const char* string_, long long string__len, int format) {
	QString string__qstr = QString::fromUtf8(string_, string__len);
	auto&& __r = ((QLocale*)self)->toDateTime(string__qstr, (QLocale::FormatType)(format));
	return new QDateTime(__r);
}
void* QLocale_toDateTime6(void* self, const char* string_, long long string__len, int format, int baseYear) {
	QString string__qstr = QString::fromUtf8(string_, string__len);
	auto&& __r = ((QLocale*)self)->toDateTime(string__qstr, (QLocale::FormatType)(format), baseYear);
	return new QDateTime(__r);
}
void* QLocale_toDateTime7(void* self, const char* string_, long long string__len, const char* format, long long format_len, int baseYear) {
	QString string__qstr = QString::fromUtf8(string_, string__len);
	QString format_qstr = QString::fromUtf8(format, format_len);
	auto&& __r = ((QLocale*)self)->toDateTime(string__qstr, format_qstr, baseYear);
	return new QDateTime(__r);
}
void* QLocale_toDate8(void* self, const char* string_, long long string__len, int format, void* cal, int baseYear) {
	QString string__qstr = QString::fromUtf8(string_, string__len);
	auto&& __r = ((QLocale*)self)->toDate(string__qstr, (QLocale::FormatType)(format), *(QCalendar*)cal, baseYear);
	return new QDate(__r);
}
void* QLocale_toDate9(void* self, const char* string_, long long string__len, const char* format, long long format_len, void* cal, int baseYear) {
	QString string__qstr = QString::fromUtf8(string_, string__len);
	QString format_qstr = QString::fromUtf8(format, format_len);
	auto&& __r = ((QLocale*)self)->toDate(string__qstr, format_qstr, *(QCalendar*)cal, baseYear);
	return new QDate(__r);
}
void* QLocale_toDateTime8(void* self, const char* string_, long long string__len, int format, void* cal, int baseYear) {
	QString string__qstr = QString::fromUtf8(string_, string__len);
	auto&& __r = ((QLocale*)self)->toDateTime(string__qstr, (QLocale::FormatType)(format), *(QCalendar*)cal, baseYear);
	return new QDateTime(__r);
}
void* QLocale_toDateTime9(void* self, const char* string_, long long string__len, const char* format, long long format_len, void* cal, int baseYear) {
	QString string__qstr = QString::fromUtf8(string_, string__len);
	QString format_qstr = QString::fromUtf8(format, format_len);
	auto&& __r = ((QLocale*)self)->toDateTime(string__qstr, format_qstr, *(QCalendar*)cal, baseYear);
	return new QDateTime(__r);
}
char* QLocale_monthName2(void* self, int arg1, int format, long long* __out_len) {
	auto&& __r = ((QLocale*)self)->monthName(arg1, (QLocale::FormatType)(format));
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLocale_standaloneMonthName2(void* self, int arg1, int format, long long* __out_len) {
	auto&& __r = ((QLocale*)self)->standaloneMonthName(arg1, (QLocale::FormatType)(format));
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLocale_dayName2(void* self, int arg1, int format, long long* __out_len) {
	auto&& __r = ((QLocale*)self)->dayName(arg1, (QLocale::FormatType)(format));
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLocale_standaloneDayName2(void* self, int arg1, int format, long long* __out_len) {
	auto&& __r = ((QLocale*)self)->standaloneDayName(arg1, (QLocale::FormatType)(format));
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLocale_currencySymbol2(void* self, int arg1, long long* __out_len) {
	auto&& __r = ((QLocale*)self)->currencySymbol((QLocale::CurrencySymbolFormat)(arg1));
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLocale_toCurrencyString9(void* self, long long arg1, const char* symbol, long long symbol_len, long long* __out_len) {
	QString symbol_qstr = QString::fromUtf8(symbol, symbol_len);
	auto&& __r = ((QLocale*)self)->toCurrencyString(arg1, symbol_qstr);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLocale_toCurrencyString10(void* self, unsigned long long arg1, const char* symbol, long long symbol_len, long long* __out_len) {
	QString symbol_qstr = QString::fromUtf8(symbol, symbol_len);
	auto&& __r = ((QLocale*)self)->toCurrencyString(arg1, symbol_qstr);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLocale_toCurrencyString11(void* self, short i, const char* symbol, long long symbol_len, long long* __out_len) {
	QString symbol_qstr = QString::fromUtf8(symbol, symbol_len);
	auto&& __r = ((QLocale*)self)->toCurrencyString(i, symbol_qstr);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLocale_toCurrencyString12(void* self, unsigned short i, const char* symbol, long long symbol_len, long long* __out_len) {
	QString symbol_qstr = QString::fromUtf8(symbol, symbol_len);
	auto&& __r = ((QLocale*)self)->toCurrencyString(i, symbol_qstr);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLocale_toCurrencyString13(void* self, int i, const char* symbol, long long symbol_len, long long* __out_len) {
	QString symbol_qstr = QString::fromUtf8(symbol, symbol_len);
	auto&& __r = ((QLocale*)self)->toCurrencyString(i, symbol_qstr);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLocale_toCurrencyString14(void* self, unsigned int i, const char* symbol, long long symbol_len, long long* __out_len) {
	QString symbol_qstr = QString::fromUtf8(symbol, symbol_len);
	auto&& __r = ((QLocale*)self)->toCurrencyString(i, symbol_qstr);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLocale_toCurrencyString15(void* self, double arg1, const char* symbol, long long symbol_len, long long* __out_len) {
	QString symbol_qstr = QString::fromUtf8(symbol, symbol_len);
	auto&& __r = ((QLocale*)self)->toCurrencyString(arg1, symbol_qstr);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLocale_toCurrencyString16(void* self, double arg1, const char* symbol, long long symbol_len, int precision, long long* __out_len) {
	QString symbol_qstr = QString::fromUtf8(symbol, symbol_len);
	auto&& __r = ((QLocale*)self)->toCurrencyString(arg1, symbol_qstr, precision);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLocale_toCurrencyString17(void* self, float i, const char* symbol, long long symbol_len, long long* __out_len) {
	QString symbol_qstr = QString::fromUtf8(symbol, symbol_len);
	auto&& __r = ((QLocale*)self)->toCurrencyString(i, symbol_qstr);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLocale_toCurrencyString18(void* self, float i, const char* symbol, long long symbol_len, int precision, long long* __out_len) {
	QString symbol_qstr = QString::fromUtf8(symbol, symbol_len);
	auto&& __r = ((QLocale*)self)->toCurrencyString(i, symbol_qstr, precision);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLocale_formattedDataSize2(void* self, long long bytes, int precision, long long* __out_len) {
	auto&& __r = ((QLocale*)self)->formattedDataSize(bytes, precision);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLocale_formattedDataSize3(void* self, long long bytes, int precision, int format, long long* __out_len) {
	auto&& __r = ((QLocale*)self)->formattedDataSize(bytes, precision, QFlags<QLocale::DataSizeFormat>(format));
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char** QLocale_uiLanguages2(void* self, int separator, long long* __out_len) {
	auto&& __r = ((QLocale*)self)->uiLanguages((QLocale::TagSeparator)(separator));
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
char* QLocale_languageToCode2(int language, int codeTypes, long long* __out_len) {
	auto&& __r = QLocale::languageToCode((QLocale::Language)(language), QFlags<QLocale::LanguageCodeType>(codeTypes));
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QLocale_codeToLanguage2(void* languageCode, int codeTypes) {
	auto&& __r = QLocale::codeToLanguage(*(QStringView*)languageCode, QFlags<QLocale::LanguageCodeType>(codeTypes));
	return (int)(__r);
}
char* QLocale_quoteString3(void* self, const char* str, long long str_len, int style, long long* __out_len) {
	QString str_qstr = QString::fromUtf8(str, str_len);
	auto&& __r = ((QLocale*)self)->quoteString(str_qstr, (QLocale::QuotationStyle)(style));
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLocale_quoteString4(void* self, void* str, int style, long long* __out_len) {
	auto&& __r = ((QLocale*)self)->quoteString(*(QStringView*)str, (QLocale::QuotationStyle)(style));
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QLocale_OperatorEqual(void* self, void* rhs) {
	auto&& __r = (*(QLocale*)self) == (*(QLocale*)rhs);
	return __r;
}
int QLocale_OperatorNotEqual(void* self, void* rhs) {
	auto&& __r = (*(QLocale*)self) != (*(QLocale*)rhs);
	return __r;
}
void QLocale_delete(void* self) { delete (QLocale*)self; }

}

#include <QtCore/qpropertyprivate.h>
#include "gen_qpropertyproxybindingdata.h"

extern "C" {

unsigned long long QPropertyProxyBindingData_d_ptr_get(void* self) {
	auto __r = ((QPropertyProxyBindingData*)self)->d_ptr;
	return __r;
}
void QPropertyProxyBindingData_d_ptr_set(void* self, unsigned long long d_ptr) {
	((QPropertyProxyBindingData*)self)->d_ptr = d_ptr;
}
void* QPropertyProxyBindingData_propertyData_get(void* self) {
	auto __r = ((QPropertyProxyBindingData*)self)->propertyData;
	return (void*)(__r);
}
void QPropertyProxyBindingData_propertyData_set(void* self, void* propertyData) {
	((QPropertyProxyBindingData*)self)->propertyData = (QUntypedPropertyData*)propertyData;
}
void QPropertyProxyBindingData_delete(void* self) { delete (QPropertyProxyBindingData*)self; }

}

#include <QtCore/qtimezone.h>
#include <QtCore/qstring.h>
#include <QtCore/qdatetime.h>
#include "gen_qtimezone__offsetdata.h"

extern "C" {

void* QTimeZone__OffsetData_new(void* arg1) {
	return new QTimeZone::OffsetData(*(QTimeZone::OffsetData*)arg1);
}
void* QTimeZone__OffsetData_new2() {
	return new QTimeZone::OffsetData();
}
void* QTimeZone__OffsetData_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QTimeZone::OffsetData*)self)->operator=(*(QTimeZone::OffsetData*)arg1);
	return (void*)&(__r);
}
char* QTimeZone__OffsetData_abbreviation_get(void* self, long long* __out_len) {
	auto __r = ((QTimeZone::OffsetData*)self)->abbreviation;
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QTimeZone__OffsetData_abbreviation_set(void* self, const char* abbreviation, long long abbreviation_len) {
	QString abbreviation_qstr = QString::fromUtf8(abbreviation, abbreviation_len);
	((QTimeZone::OffsetData*)self)->abbreviation = abbreviation_qstr;
}
void* QTimeZone__OffsetData_atUtc_get(void* self) {
	auto __r = ((QTimeZone::OffsetData*)self)->atUtc;
	return new QDateTime(__r);
}
void QTimeZone__OffsetData_atUtc_set(void* self, void* atUtc) {
	((QTimeZone::OffsetData*)self)->atUtc = *(QDateTime*)atUtc;
}
int QTimeZone__OffsetData_offsetFromUtc_get(void* self) {
	auto __r = ((QTimeZone::OffsetData*)self)->offsetFromUtc;
	return __r;
}
void QTimeZone__OffsetData_offsetFromUtc_set(void* self, int offsetFromUtc) {
	((QTimeZone::OffsetData*)self)->offsetFromUtc = offsetFromUtc;
}
int QTimeZone__OffsetData_standardTimeOffset_get(void* self) {
	auto __r = ((QTimeZone::OffsetData*)self)->standardTimeOffset;
	return __r;
}
void QTimeZone__OffsetData_standardTimeOffset_set(void* self, int standardTimeOffset) {
	((QTimeZone::OffsetData*)self)->standardTimeOffset = standardTimeOffset;
}
int QTimeZone__OffsetData_daylightTimeOffset_get(void* self) {
	auto __r = ((QTimeZone::OffsetData*)self)->daylightTimeOffset;
	return __r;
}
void QTimeZone__OffsetData_daylightTimeOffset_set(void* self, int daylightTimeOffset) {
	((QTimeZone::OffsetData*)self)->daylightTimeOffset = daylightTimeOffset;
}
void QTimeZone__OffsetData_delete(void* self) { delete (QTimeZone::OffsetData*)self; }

}

#include <QtCore/qjsonvalue.h>
#include <QtCore/qjsonarray.h>
#include <QtCore/qjsonobject.h>
#include <QtCore/qvariant.h>
#include <QtCore/qstring.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qstringview.h>
#include <QtCore/qlatin1stringview.h>
#include "gen_qjsonvalueref.h"

extern "C" {

void* QJsonValueRef_base(void* self) { return (void*)static_cast<QJsonValueConstRef*>((QJsonValueRef*)self); }
void* QJsonValueRef_new(void* arg1) {
	return new QJsonValueRef(*(QJsonValueRef*)arg1);
}
void* QJsonValueRef_new2(void* array, long long idx) {
	return new QJsonValueRef((QJsonArray*)array, idx);
}
void* QJsonValueRef_new3(void* object, long long idx) {
	return new QJsonValueRef((QJsonObject*)object, idx);
}
void* QJsonValueRef_OperatorAssign(void* self, void* val) {
	auto&& __r = ((QJsonValueRef*)self)->operator=(*(QJsonValue*)val);
	return (void*)&(__r);
}
void* QJsonValueRef_OperatorAssign2(void* self, void* val) {
	auto&& __r = ((QJsonValueRef*)self)->operator=(*(QJsonValueRef*)val);
	return (void*)&(__r);
}
void* QJsonValueRef_ToQJsonValue(void* self) {
	auto&& __r = static_cast<QJsonValue>(*(QJsonValueRef*)self);
	return new QJsonValue(__r);
}
void* QJsonValueRef_toVariant(void* self) {
	auto&& __r = ((QJsonValueRef*)self)->toVariant();
	return new QVariant(__r);
}
int QJsonValueRef_type(void* self) {
	auto&& __r = ((QJsonValueRef*)self)->type();
	return (int)(__r);
}
int QJsonValueRef_isNull(void* self) {
	auto&& __r = ((QJsonValueRef*)self)->isNull();
	return __r;
}
int QJsonValueRef_isBool(void* self) {
	auto&& __r = ((QJsonValueRef*)self)->isBool();
	return __r;
}
int QJsonValueRef_isDouble(void* self) {
	auto&& __r = ((QJsonValueRef*)self)->isDouble();
	return __r;
}
int QJsonValueRef_isString(void* self) {
	auto&& __r = ((QJsonValueRef*)self)->isString();
	return __r;
}
int QJsonValueRef_isArray(void* self) {
	auto&& __r = ((QJsonValueRef*)self)->isArray();
	return __r;
}
int QJsonValueRef_isObject(void* self) {
	auto&& __r = ((QJsonValueRef*)self)->isObject();
	return __r;
}
int QJsonValueRef_isUndefined(void* self) {
	auto&& __r = ((QJsonValueRef*)self)->isUndefined();
	return __r;
}
int QJsonValueRef_toBool(void* self) {
	auto&& __r = ((QJsonValueRef*)self)->toBool();
	return __r;
}
int QJsonValueRef_toInt(void* self) {
	auto&& __r = ((QJsonValueRef*)self)->toInt();
	return __r;
}
long long QJsonValueRef_toInteger(void* self) {
	auto&& __r = ((QJsonValueRef*)self)->toInteger();
	return __r;
}
double QJsonValueRef_toDouble(void* self) {
	auto&& __r = ((QJsonValueRef*)self)->toDouble();
	return __r;
}
char* QJsonValueRef_toString(void* self, long long* __out_len) {
	auto&& __r = ((QJsonValueRef*)self)->toString();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QJsonValueRef_toStringView(void* self) {
	auto&& __r = ((QJsonValueRef*)self)->toStringView();
	return new QAnyStringView(__r);
}
void* QJsonValueRef_toArray(void* self) {
	auto&& __r = ((QJsonValueRef*)self)->toArray();
	return new QJsonArray(__r);
}
void* QJsonValueRef_toObject(void* self) {
	auto&& __r = ((QJsonValueRef*)self)->toObject();
	return new QJsonObject(__r);
}
void* QJsonValueRef_OperatorSubscript(void* self, void* key) {
	auto&& __r = ((QJsonValueRef*)self)->operator[](*(QStringView*)key);
	return new QJsonValue(__r);
}
void* QJsonValueRef_OperatorSubscript2(void* self, void* key) {
	auto&& __r = ((QJsonValueRef*)self)->operator[](*(QLatin1String*)key);
	return new QJsonValue(__r);
}
void* QJsonValueRef_OperatorSubscript3(void* self, long long i) {
	auto&& __r = ((QJsonValueRef*)self)->operator[](i);
	return new QJsonValue(__r);
}
int QJsonValueRef_toBool2(void* self, int defaultValue) {
	auto&& __r = ((QJsonValueRef*)self)->toBool(defaultValue);
	return __r;
}
int QJsonValueRef_toInt2(void* self, int defaultValue) {
	auto&& __r = ((QJsonValueRef*)self)->toInt(defaultValue);
	return __r;
}
long long QJsonValueRef_toInteger2(void* self, long long defaultValue) {
	auto&& __r = ((QJsonValueRef*)self)->toInteger(defaultValue);
	return __r;
}
double QJsonValueRef_toDouble2(void* self, double defaultValue) {
	auto&& __r = ((QJsonValueRef*)self)->toDouble(defaultValue);
	return __r;
}
char* QJsonValueRef_toString2(void* self, const char* defaultValue, long long defaultValue_len, long long* __out_len) {
	QString defaultValue_qstr = QString::fromUtf8(defaultValue, defaultValue_len);
	auto&& __r = ((QJsonValueRef*)self)->toString(defaultValue_qstr);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QJsonValueRef_toStringView2(void* self, const char* defaultValue, long long defaultValue_len) {
	QAnyStringView defaultValue_qstr = QAnyStringView(defaultValue, defaultValue_len);
	auto&& __r = ((QJsonValueRef*)self)->toStringView(defaultValue_qstr);
	return new QAnyStringView(__r);
}
int QJsonValueRef_OperatorEqual(void* self, void* rhs) {
	auto&& __r = (*(QJsonValueRef*)self) == (*(QJsonValueRef*)rhs);
	return __r;
}
int QJsonValueRef_OperatorNotEqual(void* self, void* rhs) {
	auto&& __r = (*(QJsonValueRef*)self) != (*(QJsonValueRef*)rhs);
	return __r;
}
int QJsonValueRef_OperatorEqual2(void* self, void* rhs) {
	auto&& __r = (*(QJsonValueRef*)self) == (*(QJsonValueConstRef*)rhs);
	return __r;
}
int QJsonValueRef_OperatorNotEqual2(void* self, void* rhs) {
	auto&& __r = (*(QJsonValueRef*)self) != (*(QJsonValueConstRef*)rhs);
	return __r;
}
void QJsonValueRef_delete(void* self) { delete (QJsonValueRef*)self; }

}

#include <QtCore/qobjectcleanuphandler.h>
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
#include "gen_qobjectcleanuphandler.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QObjectCleanupHandler__Go : public QObjectCleanupHandler {
public:
	long long __goId;
	unsigned char __goOn[8] = {};
	template<typename... A> QObjectCleanupHandler__Go(long long id, A&&... a) : QObjectCleanupHandler(std::forward<A>(a)...), __goId(id) {}
	using QObjectCleanupHandler::sender;
	using QObjectCleanupHandler::senderSignalIndex;
	using QObjectCleanupHandler::receivers;
	using QObjectCleanupHandler::isSignalConnected;
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QObjectCleanupHandler::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QObjectCleanupHandler::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QObjectCleanupHandler::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QObjectCleanupHandler::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QObjectCleanupHandler::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QObjectCleanupHandler::metaObject();
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QObjectCleanupHandler::event(event);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QObjectCleanupHandler::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QObjectCleanupHandler::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QObjectCleanupHandler::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QObjectCleanupHandler::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QObjectCleanupHandler::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QObjectCleanupHandler::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QObjectCleanupHandler;
static std::mutex __connsMu_QObjectCleanupHandler;

extern "C" {

void* QObjectCleanupHandler_base(void* self) { return (void*)static_cast<QObject*>((QObjectCleanupHandler*)self); }
void QObjectCleanupHandler__set_override(void* self, int idx) { ((QObjectCleanupHandler__Go*)self)->__goOn[idx] = 1; }
void* QObjectCleanupHandler_fromQObject(void* o) { return (void*)qobject_cast<QObjectCleanupHandler*>((QObject*)o); }
void* QObjectCleanupHandler_new(long long __goid) {
	return new QObjectCleanupHandler__Go(__goid);
}
char* QObjectCleanupHandler_tr(const char* s, long long* __out_len) {
	auto&& __r = QObjectCleanupHandler::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QObjectCleanupHandler_add(void* self, void* object) {
	auto&& __r = ((QObjectCleanupHandler*)self)->add((QObject*)object);
	return (void*)(__r);
}
void QObjectCleanupHandler_remove(void* self, void* object) {
	((QObjectCleanupHandler*)self)->remove((QObject*)object);
}
int QObjectCleanupHandler_isEmpty(void* self) {
	auto&& __r = ((QObjectCleanupHandler*)self)->isEmpty();
	return __r;
}
void QObjectCleanupHandler_clear(void* self) {
	((QObjectCleanupHandler*)self)->clear();
}
void* QObjectCleanupHandler_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QObjectCleanupHandler::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QObjectCleanupHandler_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QObjectCleanupHandler::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QObjectCleanupHandler_sender(void* self) {
	auto&& __r = ((QObjectCleanupHandler__Go*)self)->sender();
	return (void*)(__r);
}
int QObjectCleanupHandler_senderSignalIndex(void* self) {
	auto&& __r = ((QObjectCleanupHandler__Go*)self)->senderSignalIndex();
	return __r;
}
int QObjectCleanupHandler_receivers(void* self, const char* signal) {
	auto&& __r = ((QObjectCleanupHandler__Go*)self)->receivers(signal);
	return __r;
}
int QObjectCleanupHandler_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QObjectCleanupHandler__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QObjectCleanupHandler_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QObjectCleanupHandler::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QObjectCleanupHandler_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QObjectCleanupHandler::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QObjectCleanupHandler_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QObjectCleanupHandler::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QObjectCleanupHandler_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QObjectCleanupHandler::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QObjectCleanupHandler_metaObject_default(void* self) {
	auto&& __r = ((QObjectCleanupHandler*)self)->QObjectCleanupHandler::metaObject();
	return (void*)(__r);
}
int QObjectCleanupHandler_event_default(void* self, void* event) {
	auto&& __r = ((QObjectCleanupHandler*)self)->QObjectCleanupHandler::event((QEvent*)event);
	return __r;
}
int QObjectCleanupHandler_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QObjectCleanupHandler*)self)->QObjectCleanupHandler::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QObjectCleanupHandler_timerEvent_default(void* self, void* event) {
	((QObjectCleanupHandler__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QObjectCleanupHandler_childEvent_default(void* self, void* event) {
	((QObjectCleanupHandler__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QObjectCleanupHandler_customEvent_default(void* self, void* event) {
	((QObjectCleanupHandler__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QObjectCleanupHandler_connectNotify_default(void* self, void* signal) {
	((QObjectCleanupHandler__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QObjectCleanupHandler_disconnectNotify_default(void* self, void* signal) {
	((QObjectCleanupHandler__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QObjectCleanupHandler_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QObjectCleanupHandler*)self, &QObjectCleanupHandler::destroyed, (QObjectCleanupHandler*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QObjectCleanupHandler);
	__conns_QObjectCleanupHandler[__id] = __c;
}
void QObjectCleanupHandler_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QObjectCleanupHandler*)self, &QObjectCleanupHandler::destroyed, (QObjectCleanupHandler*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QObjectCleanupHandler);
	__conns_QObjectCleanupHandler[__id] = __c;
}
void QObjectCleanupHandler_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QObjectCleanupHandler*)self, &QObjectCleanupHandler::objectNameChanged, (QObjectCleanupHandler*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QObjectCleanupHandler);
	__conns_QObjectCleanupHandler[__id] = __c;
}
void QObjectCleanupHandler__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QObjectCleanupHandler);
		auto __it = __conns_QObjectCleanupHandler.find(id);
		if (__it == __conns_QObjectCleanupHandler.end()) return;
		__c = __it->second;
		__conns_QObjectCleanupHandler.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QObjectCleanupHandler__Go_new(long long id) { return new QObjectCleanupHandler__Go(id); }
void QObjectCleanupHandler_delete(void* self) { delete (QObjectCleanupHandler*)self; }

}

#include <QtCore/qrandom.h>
#include "gen_qrandomgenerator.h"

extern "C" {

void* QRandomGenerator_new() {
	return new QRandomGenerator();
}
void* QRandomGenerator_new2(void* other) {
	return new QRandomGenerator(*(QRandomGenerator*)other);
}
void* QRandomGenerator_new3(unsigned int seedValue) {
	return new QRandomGenerator(seedValue);
}
void* QRandomGenerator_OperatorAssign(void* self, void* other) {
	auto&& __r = ((QRandomGenerator*)self)->operator=(*(QRandomGenerator*)other);
	return (void*)&(__r);
}
unsigned int QRandomGenerator_generate(void* self) {
	auto&& __r = ((QRandomGenerator*)self)->generate();
	return __r;
}
unsigned long long QRandomGenerator_generate64(void* self) {
	auto&& __r = ((QRandomGenerator*)self)->generate64();
	return __r;
}
double QRandomGenerator_generateDouble(void* self) {
	auto&& __r = ((QRandomGenerator*)self)->generateDouble();
	return __r;
}
double QRandomGenerator_bounded(void* self, double highest) {
	auto&& __r = ((QRandomGenerator*)self)->bounded(highest);
	return __r;
}
unsigned int QRandomGenerator_bounded2(void* self, unsigned int highest) {
	auto&& __r = ((QRandomGenerator*)self)->bounded(highest);
	return __r;
}
unsigned int QRandomGenerator_bounded3(void* self, unsigned int lowest, unsigned int highest) {
	auto&& __r = ((QRandomGenerator*)self)->bounded(lowest, highest);
	return __r;
}
int QRandomGenerator_bounded4(void* self, int highest) {
	auto&& __r = ((QRandomGenerator*)self)->bounded(highest);
	return __r;
}
int QRandomGenerator_bounded5(void* self, int lowest, int highest) {
	auto&& __r = ((QRandomGenerator*)self)->bounded(lowest, highest);
	return __r;
}
unsigned long long QRandomGenerator_bounded6(void* self, unsigned long long highest) {
	auto&& __r = ((QRandomGenerator*)self)->bounded(highest);
	return __r;
}
unsigned long long QRandomGenerator_bounded7(void* self, unsigned long long lowest, unsigned long long highest) {
	auto&& __r = ((QRandomGenerator*)self)->bounded(lowest, highest);
	return __r;
}
long long QRandomGenerator_bounded8(void* self, long long highest) {
	auto&& __r = ((QRandomGenerator*)self)->bounded(highest);
	return __r;
}
long long QRandomGenerator_bounded9(void* self, long long lowest, long long highest) {
	auto&& __r = ((QRandomGenerator*)self)->bounded(lowest, highest);
	return __r;
}
long long QRandomGenerator_bounded10(void* self, int lowest, long long highest) {
	auto&& __r = ((QRandomGenerator*)self)->bounded(lowest, highest);
	return __r;
}
long long QRandomGenerator_bounded11(void* self, long long lowest, int highest) {
	auto&& __r = ((QRandomGenerator*)self)->bounded(lowest, highest);
	return __r;
}
unsigned long long QRandomGenerator_bounded12(void* self, unsigned int lowest, unsigned long long highest) {
	auto&& __r = ((QRandomGenerator*)self)->bounded(lowest, highest);
	return __r;
}
unsigned long long QRandomGenerator_bounded13(void* self, unsigned long long lowest, unsigned int highest) {
	auto&& __r = ((QRandomGenerator*)self)->bounded(lowest, highest);
	return __r;
}
void QRandomGenerator_generate2(void* self, unsigned int* begin, unsigned int* end) {
	quint32 begin_local{};
	quint32 end_local{};
	((QRandomGenerator*)self)->generate(&begin_local, &end_local);
	*begin = begin_local;
	*end = end_local;
}
unsigned int QRandomGenerator_OperatorCall(void* self) {
	auto&& __r = ((QRandomGenerator*)self)->operator()();
	return __r;
}
void QRandomGenerator_seed(void* self) {
	((QRandomGenerator*)self)->seed();
}
void QRandomGenerator_discard(void* self, unsigned long long z) {
	((QRandomGenerator*)self)->discard(z);
}
unsigned int QRandomGenerator_min() {
	auto&& __r = QRandomGenerator::min();
	return __r;
}
unsigned int QRandomGenerator_max() {
	auto&& __r = QRandomGenerator::max();
	return __r;
}
void* QRandomGenerator_system() {
	auto&& __r = QRandomGenerator::system();
	return (void*)(__r);
}
void* QRandomGenerator_global() {
	auto&& __r = QRandomGenerator::global();
	return (void*)(__r);
}
void* QRandomGenerator_securelySeeded() {
	auto&& __r = QRandomGenerator::securelySeeded();
	return new QRandomGenerator(__r);
}
void QRandomGenerator_seed2(void* self, unsigned int s) {
	((QRandomGenerator*)self)->seed(s);
}
int QRandomGenerator_OperatorEqual(void* self, void* rng2) {
	auto&& __r = (*(QRandomGenerator*)self) == (*(QRandomGenerator*)rng2);
	return __r;
}
int QRandomGenerator_OperatorNotEqual(void* self, void* rng2) {
	auto&& __r = (*(QRandomGenerator*)self) != (*(QRandomGenerator*)rng2);
	return __r;
}
void QRandomGenerator_delete(void* self) { delete (QRandomGenerator*)self; }

}

#include <QtCore/qsocketnotifier.h>
#include "gen_qsocketdescriptor.h"

extern "C" {

void* QSocketDescriptor_new() {
	return new QSocketDescriptor();
}
void* QSocketDescriptor_new2(long long desc) {
	return new QSocketDescriptor(desc);
}
long long QSocketDescriptor_ToQintptr(void* self) {
	auto&& __r = static_cast<long long>(*(QSocketDescriptor*)self);
	return __r;
}
int QSocketDescriptor_isValid(void* self) {
	auto&& __r = ((QSocketDescriptor*)self)->isValid();
	return __r;
}
int QSocketDescriptor_OperatorEqual(void* self, void* rhs) {
	auto&& __r = (*(QSocketDescriptor*)self) == (*(QSocketDescriptor*)rhs);
	return __r;
}
int QSocketDescriptor_OperatorNotEqual(void* self, void* rhs) {
	auto&& __r = (*(QSocketDescriptor*)self) != (*(QSocketDescriptor*)rhs);
	return __r;
}
void QSocketDescriptor_delete(void* self) { delete (QSocketDescriptor*)self; }

}

#include <QtGui/qpaintdevice.h>
#include <QtGui/qpaintengine.h>
#include <QtGui/qpainter.h>
#include <QtCore/qpoint.h>
#include "gen_qpaintdevice.h"

extern "C" {

int QPaintDevice_devType(void* self) {
	auto&& __r = ((QPaintDevice*)self)->devType();
	return __r;
}
int QPaintDevice_paintingActive(void* self) {
	auto&& __r = ((QPaintDevice*)self)->paintingActive();
	return __r;
}
void* QPaintDevice_paintEngine(void* self) {
	auto&& __r = ((QPaintDevice*)self)->paintEngine();
	return (void*)(__r);
}
int QPaintDevice_width(void* self) {
	auto&& __r = ((QPaintDevice*)self)->width();
	return __r;
}
int QPaintDevice_height(void* self) {
	auto&& __r = ((QPaintDevice*)self)->height();
	return __r;
}
int QPaintDevice_widthMM(void* self) {
	auto&& __r = ((QPaintDevice*)self)->widthMM();
	return __r;
}
int QPaintDevice_heightMM(void* self) {
	auto&& __r = ((QPaintDevice*)self)->heightMM();
	return __r;
}
int QPaintDevice_logicalDpiX(void* self) {
	auto&& __r = ((QPaintDevice*)self)->logicalDpiX();
	return __r;
}
int QPaintDevice_logicalDpiY(void* self) {
	auto&& __r = ((QPaintDevice*)self)->logicalDpiY();
	return __r;
}
int QPaintDevice_physicalDpiX(void* self) {
	auto&& __r = ((QPaintDevice*)self)->physicalDpiX();
	return __r;
}
int QPaintDevice_physicalDpiY(void* self) {
	auto&& __r = ((QPaintDevice*)self)->physicalDpiY();
	return __r;
}
double QPaintDevice_devicePixelRatio(void* self) {
	auto&& __r = ((QPaintDevice*)self)->devicePixelRatio();
	return __r;
}
double QPaintDevice_devicePixelRatioF(void* self) {
	auto&& __r = ((QPaintDevice*)self)->devicePixelRatioF();
	return __r;
}
int QPaintDevice_colorCount(void* self) {
	auto&& __r = ((QPaintDevice*)self)->colorCount();
	return __r;
}
int QPaintDevice_depth(void* self) {
	auto&& __r = ((QPaintDevice*)self)->depth();
	return __r;
}
double QPaintDevice_devicePixelRatioFScale() {
	auto&& __r = QPaintDevice::devicePixelRatioFScale();
	return __r;
}
int QPaintDevice_encodeMetricF(int metric, double value) {
	auto&& __r = QPaintDevice::encodeMetricF((QPaintDevice::PaintDeviceMetric)(metric), value);
	return __r;
}
void QPaintDevice_delete(void* self) { delete (QPaintDevice*)self; }

}

#include <QtGui/qevent.h>
#include <QtCore/qcoreevent.h>
#include <QtGui/qpointingdevice.h>
#include <QtCore/qnamespace.h>
#include <QtGui/qeventpoint.h>
#include <QtCore/qobject.h>
#include <QtGui/qinputdevice.h>
#include "gen_qpointerevent.h"

extern "C" {

void* QPointerEvent_base(void* self) { return (void*)static_cast<QInputEvent*>((QPointerEvent*)self); }
void QPointerEvent_chain(void* self, void** out) {
	QPointerEvent* __p = (QPointerEvent*)self;
	QInputEvent* __b0 = static_cast<QInputEvent*>(__p); out[0] = (void*)__b0;
	QEvent* __b1 = static_cast<QEvent*>(__b0); out[1] = (void*)__b1;
}
void* QPointerEvent_new(int type_, void* dev) {
	return new QPointerEvent((QEvent::Type)(type_), (QPointingDevice*)dev);
}
void* QPointerEvent_new2(int type_, void* dev, int modifiers) {
	return new QPointerEvent((QEvent::Type)(type_), (QPointingDevice*)dev, QFlags<Qt::KeyboardModifier>(modifiers));
}
void* QPointerEvent_new3(int type_, void* dev, int modifiers, void** points, long long points_len) {
	QList<QEventPoint> points_list; points_list.reserve(points_len);
	for (long long __i = 0; __i < points_len; __i++) points_list.append(*(QEventPoint*)points[__i]);
	return new QPointerEvent((QEvent::Type)(type_), (QPointingDevice*)dev, QFlags<Qt::KeyboardModifier>(modifiers), points_list);
}
void* QPointerEvent_pointingDevice(void* self) {
	auto&& __r = ((QPointerEvent*)self)->pointingDevice();
	return (void*)(__r);
}
int QPointerEvent_pointerType(void* self) {
	auto&& __r = ((QPointerEvent*)self)->pointerType();
	return (int)(__r);
}
long long QPointerEvent_pointCount(void* self) {
	auto&& __r = ((QPointerEvent*)self)->pointCount();
	return __r;
}
void* QPointerEvent_point(void* self, long long i) {
	auto&& __r = ((QPointerEvent*)self)->point(i);
	return (void*)&(__r);
}
void** QPointerEvent_points(void* self, long long* __out_len) {
	auto&& __r = ((QPointerEvent*)self)->points();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QEventPoint(__e);
	return __arr;
}
void* QPointerEvent_pointById(void* self, int id) {
	auto&& __r = ((QPointerEvent*)self)->pointById(id);
	return (void*)(__r);
}
int QPointerEvent_allPointsGrabbed(void* self) {
	auto&& __r = ((QPointerEvent*)self)->allPointsGrabbed();
	return __r;
}
int QPointerEvent_isBeginEvent(void* self) {
	auto&& __r = ((QPointerEvent*)self)->isBeginEvent();
	return __r;
}
int QPointerEvent_isUpdateEvent(void* self) {
	auto&& __r = ((QPointerEvent*)self)->isUpdateEvent();
	return __r;
}
int QPointerEvent_isEndEvent(void* self) {
	auto&& __r = ((QPointerEvent*)self)->isEndEvent();
	return __r;
}
int QPointerEvent_allPointsAccepted(void* self) {
	auto&& __r = ((QPointerEvent*)self)->allPointsAccepted();
	return __r;
}
void* QPointerEvent_exclusiveGrabber(void* self, void* point) {
	auto&& __r = ((QPointerEvent*)self)->exclusiveGrabber(*(QEventPoint*)point);
	return (void*)(__r);
}
void QPointerEvent_setExclusiveGrabber(void* self, void* point, void* exclusiveGrabber) {
	((QPointerEvent*)self)->setExclusiveGrabber(*(QEventPoint*)point, (QObject*)exclusiveGrabber);
}
void QPointerEvent_clearPassiveGrabbers(void* self, void* point) {
	((QPointerEvent*)self)->clearPassiveGrabbers(*(QEventPoint*)point);
}
int QPointerEvent_addPassiveGrabber(void* self, void* point, void* grabber) {
	auto&& __r = ((QPointerEvent*)self)->addPassiveGrabber(*(QEventPoint*)point, (QObject*)grabber);
	return __r;
}
int QPointerEvent_removePassiveGrabber(void* self, void* point, void* grabber) {
	auto&& __r = ((QPointerEvent*)self)->removePassiveGrabber(*(QEventPoint*)point, (QObject*)grabber);
	return __r;
}
int QPointerEvent_registerEventType() {
	auto&& __r = QPointerEvent::registerEventType();
	return __r;
}
int QPointerEvent_registerEventType2(int hint) {
	auto&& __r = QPointerEvent::registerEventType(hint);
	return __r;
}
void QPointerEvent_delete(void* self) { delete (QPointerEvent*)self; }

}

#include <QtGui/qevent.h>
#include <QtCore/qpoint.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qmimedata.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qrect.h>
#include <QtCore/qobject.h>
#include "gen_qdragmoveevent.h"

extern "C" {

void* QDragMoveEvent_base(void* self) { return (void*)static_cast<QDropEvent*>((QDragMoveEvent*)self); }
void QDragMoveEvent_chain(void* self, void** out) {
	QDragMoveEvent* __p = (QDragMoveEvent*)self;
	QDropEvent* __b0 = static_cast<QDropEvent*>(__p); out[0] = (void*)__b0;
	QEvent* __b1 = static_cast<QEvent*>(__b0); out[1] = (void*)__b1;
}
void* QDragMoveEvent_new(void* pos, int actions, void* data, int buttons, int modifiers) {
	return new QDragMoveEvent(*(QPoint*)pos, QFlags<Qt::DropAction>(actions), (QMimeData*)data, QFlags<Qt::MouseButton>(buttons), QFlags<Qt::KeyboardModifier>(modifiers));
}
void* QDragMoveEvent_new2(void* pos, int actions, void* data, int buttons, int modifiers, int type_) {
	return new QDragMoveEvent(*(QPoint*)pos, QFlags<Qt::DropAction>(actions), (QMimeData*)data, QFlags<Qt::MouseButton>(buttons), QFlags<Qt::KeyboardModifier>(modifiers), (QEvent::Type)(type_));
}
void* QDragMoveEvent_answerRect(void* self) {
	auto&& __r = ((QDragMoveEvent*)self)->answerRect();
	return new QRect(__r);
}
void QDragMoveEvent_accept(void* self) {
	((QDragMoveEvent*)self)->accept();
}
void QDragMoveEvent_ignore(void* self) {
	((QDragMoveEvent*)self)->ignore();
}
void QDragMoveEvent_accept2(void* self, void* r) {
	((QDragMoveEvent*)self)->accept(*(QRect*)r);
}
void QDragMoveEvent_ignore2(void* self, void* r) {
	((QDragMoveEvent*)self)->ignore(*(QRect*)r);
}
int QDragMoveEvent_registerEventType() {
	auto&& __r = QDragMoveEvent::registerEventType();
	return __r;
}
int QDragMoveEvent_registerEventType2(int hint) {
	auto&& __r = QDragMoveEvent::registerEventType(hint);
	return __r;
}
void QDragMoveEvent_delete(void* self) { delete (QDragMoveEvent*)self; }

}

#include <QtGui/qbrush.h>
#include <QtCore/qnamespace.h>
#include <QtGui/qcolor.h>
#include <QtGui/qpixmap.h>
#include <QtGui/qimage.h>
#include <QtCore/qvariant.h>
#include <QtGui/qtransform.h>
#include "gen_qbrush.h"

extern "C" {

void* QBrush_new() {
	return new QBrush();
}
void* QBrush_new2(int bs) {
	return new QBrush((Qt::BrushStyle)(bs));
}
void* QBrush_new3(void* color) {
	return new QBrush(*(QColor*)color);
}
void* QBrush_new4(int color) {
	return new QBrush((Qt::GlobalColor)(color));
}
void* QBrush_new5(void* color, void* pixmap) {
	return new QBrush(*(QColor*)color, *(QPixmap*)pixmap);
}
void* QBrush_new6(int color, void* pixmap) {
	return new QBrush((Qt::GlobalColor)(color), *(QPixmap*)pixmap);
}
void* QBrush_new7(void* pixmap) {
	return new QBrush(*(QPixmap*)pixmap);
}
void* QBrush_new8(void* image) {
	return new QBrush(*(QImage*)image);
}
void* QBrush_new9(void* brush) {
	return new QBrush(*(QBrush*)brush);
}
void* QBrush_new10(void* gradient) {
	return new QBrush(*(QGradient*)gradient);
}
void* QBrush_new11(void* color, int bs) {
	return new QBrush(*(QColor*)color, (Qt::BrushStyle)(bs));
}
void* QBrush_new12(int color, int bs) {
	return new QBrush((Qt::GlobalColor)(color), (Qt::BrushStyle)(bs));
}
void* QBrush_OperatorAssign(void* self, void* brush) {
	auto&& __r = ((QBrush*)self)->operator=(*(QBrush*)brush);
	return (void*)&(__r);
}
void QBrush_swap(void* self, void* other) {
	((QBrush*)self)->swap(*(QBrush*)other);
}
void* QBrush_OperatorAssign2(void* self, int style) {
	auto&& __r = ((QBrush*)self)->operator=((Qt::BrushStyle)(style));
	return (void*)&(__r);
}
void* QBrush_OperatorAssign3(void* self, void* color) {
	auto&& __r = ((QBrush*)self)->operator=(*(QColor*)color);
	return (void*)&(__r);
}
void* QBrush_OperatorAssign4(void* self, int color) {
	auto&& __r = ((QBrush*)self)->operator=((Qt::GlobalColor)(color));
	return (void*)&(__r);
}
void* QBrush_ToQVariant(void* self) {
	auto&& __r = static_cast<QVariant>(*(QBrush*)self);
	return new QVariant(__r);
}
int QBrush_style(void* self) {
	auto&& __r = ((QBrush*)self)->style();
	return (int)(__r);
}
void QBrush_setStyle(void* self, int arg1) {
	((QBrush*)self)->setStyle((Qt::BrushStyle)(arg1));
}
void* QBrush_transform(void* self) {
	auto&& __r = ((QBrush*)self)->transform();
	return new QTransform(__r);
}
void QBrush_setTransform(void* self, void* arg1) {
	((QBrush*)self)->setTransform(*(QTransform*)arg1);
}
void* QBrush_texture(void* self) {
	auto&& __r = ((QBrush*)self)->texture();
	return new QPixmap(__r);
}
void QBrush_setTexture(void* self, void* pixmap) {
	((QBrush*)self)->setTexture(*(QPixmap*)pixmap);
}
void* QBrush_textureImage(void* self) {
	auto&& __r = ((QBrush*)self)->textureImage();
	return new QImage(__r);
}
void QBrush_setTextureImage(void* self, void* image) {
	((QBrush*)self)->setTextureImage(*(QImage*)image);
}
void* QBrush_color(void* self) {
	auto&& __r = ((QBrush*)self)->color();
	return (void*)&(__r);
}
void QBrush_setColor(void* self, void* color) {
	((QBrush*)self)->setColor(*(QColor*)color);
}
void QBrush_setColor2(void* self, int color) {
	((QBrush*)self)->setColor((Qt::GlobalColor)(color));
}
void* QBrush_gradient(void* self) {
	auto&& __r = ((QBrush*)self)->gradient();
	return (void*)(__r);
}
int QBrush_isOpaque(void* self) {
	auto&& __r = ((QBrush*)self)->isOpaque();
	return __r;
}
int QBrush_OperatorEqual(void* self, void* b) {
	auto&& __r = ((QBrush*)self)->operator==(*(QBrush*)b);
	return __r;
}
int QBrush_OperatorNotEqual(void* self, void* b) {
	auto&& __r = ((QBrush*)self)->operator!=(*(QBrush*)b);
	return __r;
}
int QBrush_isDetached(void* self) {
	auto&& __r = ((QBrush*)self)->isDetached();
	return __r;
}
int QBrush_OperatorEqual2(void* self, void* rhs) {
	auto&& __r = (*(QBrush*)self) == (*(QColor*)rhs);
	return __r;
}
int QBrush_OperatorNotEqual2(void* self, void* rhs) {
	auto&& __r = (*(QBrush*)self) != (*(QColor*)rhs);
	return __r;
}
void* QBrush_toQVariant(void* self) {
	auto&& __r = QVariant::fromValue(*(QBrush*)self);
	return new QVariant(__r);
}
void QBrush_delete(void* self) { delete (QBrush*)self; }

}

#include <QtGui/qpalette.h>
#include <QtGui/qcolor.h>
#include <QtCore/qnamespace.h>
#include <QtGui/qbrush.h>
#include <QtCore/qvariant.h>
#include "gen_qpalette.h"

extern "C" {

void* QPalette_new() {
	return new QPalette();
}
void* QPalette_new2(void* button) {
	return new QPalette(*(QColor*)button);
}
void* QPalette_new3(int button) {
	return new QPalette((Qt::GlobalColor)(button));
}
void* QPalette_new4(void* button, void* window) {
	return new QPalette(*(QColor*)button, *(QColor*)window);
}
void* QPalette_new5(void* windowText, void* button, void* light, void* dark, void* mid, void* text, void* bright_text, void* base, void* window) {
	return new QPalette(*(QBrush*)windowText, *(QBrush*)button, *(QBrush*)light, *(QBrush*)dark, *(QBrush*)mid, *(QBrush*)text, *(QBrush*)bright_text, *(QBrush*)base, *(QBrush*)window);
}
void* QPalette_new6(void* windowText, void* window, void* light, void* dark, void* mid, void* text, void* base) {
	return new QPalette(*(QColor*)windowText, *(QColor*)window, *(QColor*)light, *(QColor*)dark, *(QColor*)mid, *(QColor*)text, *(QColor*)base);
}
void* QPalette_new7(void* palette) {
	return new QPalette(*(QPalette*)palette);
}
void* QPalette_OperatorAssign(void* self, void* palette) {
	auto&& __r = ((QPalette*)self)->operator=(*(QPalette*)palette);
	return (void*)&(__r);
}
void QPalette_swap(void* self, void* other) {
	((QPalette*)self)->swap(*(QPalette*)other);
}
void* QPalette_ToQVariant(void* self) {
	auto&& __r = static_cast<QVariant>(*(QPalette*)self);
	return new QVariant(__r);
}
int QPalette_currentColorGroup(void* self) {
	auto&& __r = ((QPalette*)self)->currentColorGroup();
	return (int)(__r);
}
void QPalette_setCurrentColorGroup(void* self, int cg) {
	((QPalette*)self)->setCurrentColorGroup((QPalette::ColorGroup)(cg));
}
void* QPalette_color(void* self, int cg, int cr) {
	auto&& __r = ((QPalette*)self)->color((QPalette::ColorGroup)(cg), (QPalette::ColorRole)(cr));
	return (void*)&(__r);
}
void* QPalette_brush(void* self, int cg, int cr) {
	auto&& __r = ((QPalette*)self)->brush((QPalette::ColorGroup)(cg), (QPalette::ColorRole)(cr));
	return (void*)&(__r);
}
void QPalette_setColor(void* self, int cg, int cr, void* color) {
	((QPalette*)self)->setColor((QPalette::ColorGroup)(cg), (QPalette::ColorRole)(cr), *(QColor*)color);
}
void QPalette_setColor2(void* self, int cr, void* color) {
	((QPalette*)self)->setColor((QPalette::ColorRole)(cr), *(QColor*)color);
}
void QPalette_setBrush(void* self, int cr, void* brush) {
	((QPalette*)self)->setBrush((QPalette::ColorRole)(cr), *(QBrush*)brush);
}
int QPalette_isBrushSet(void* self, int cg, int cr) {
	auto&& __r = ((QPalette*)self)->isBrushSet((QPalette::ColorGroup)(cg), (QPalette::ColorRole)(cr));
	return __r;
}
void QPalette_setBrush2(void* self, int cg, int cr, void* brush) {
	((QPalette*)self)->setBrush((QPalette::ColorGroup)(cg), (QPalette::ColorRole)(cr), *(QBrush*)brush);
}
void QPalette_setColorGroup(void* self, int cr, void* windowText, void* button, void* light, void* dark, void* mid, void* text, void* bright_text, void* base, void* window) {
	((QPalette*)self)->setColorGroup((QPalette::ColorGroup)(cr), *(QBrush*)windowText, *(QBrush*)button, *(QBrush*)light, *(QBrush*)dark, *(QBrush*)mid, *(QBrush*)text, *(QBrush*)bright_text, *(QBrush*)base, *(QBrush*)window);
}
int QPalette_isEqual(void* self, int cr1, int cr2) {
	auto&& __r = ((QPalette*)self)->isEqual((QPalette::ColorGroup)(cr1), (QPalette::ColorGroup)(cr2));
	return __r;
}
void* QPalette_color2(void* self, int cr) {
	auto&& __r = ((QPalette*)self)->color((QPalette::ColorRole)(cr));
	return (void*)&(__r);
}
void* QPalette_brush2(void* self, int cr) {
	auto&& __r = ((QPalette*)self)->brush((QPalette::ColorRole)(cr));
	return (void*)&(__r);
}
void* QPalette_windowText(void* self) {
	auto&& __r = ((QPalette*)self)->windowText();
	return (void*)&(__r);
}
void* QPalette_button(void* self) {
	auto&& __r = ((QPalette*)self)->button();
	return (void*)&(__r);
}
void* QPalette_light(void* self) {
	auto&& __r = ((QPalette*)self)->light();
	return (void*)&(__r);
}
void* QPalette_dark(void* self) {
	auto&& __r = ((QPalette*)self)->dark();
	return (void*)&(__r);
}
void* QPalette_mid(void* self) {
	auto&& __r = ((QPalette*)self)->mid();
	return (void*)&(__r);
}
void* QPalette_text(void* self) {
	auto&& __r = ((QPalette*)self)->text();
	return (void*)&(__r);
}
void* QPalette_base(void* self) {
	auto&& __r = ((QPalette*)self)->base();
	return (void*)&(__r);
}
void* QPalette_alternateBase(void* self) {
	auto&& __r = ((QPalette*)self)->alternateBase();
	return (void*)&(__r);
}
void* QPalette_toolTipBase(void* self) {
	auto&& __r = ((QPalette*)self)->toolTipBase();
	return (void*)&(__r);
}
void* QPalette_toolTipText(void* self) {
	auto&& __r = ((QPalette*)self)->toolTipText();
	return (void*)&(__r);
}
void* QPalette_window(void* self) {
	auto&& __r = ((QPalette*)self)->window();
	return (void*)&(__r);
}
void* QPalette_midlight(void* self) {
	auto&& __r = ((QPalette*)self)->midlight();
	return (void*)&(__r);
}
void* QPalette_brightText(void* self) {
	auto&& __r = ((QPalette*)self)->brightText();
	return (void*)&(__r);
}
void* QPalette_buttonText(void* self) {
	auto&& __r = ((QPalette*)self)->buttonText();
	return (void*)&(__r);
}
void* QPalette_shadow(void* self) {
	auto&& __r = ((QPalette*)self)->shadow();
	return (void*)&(__r);
}
void* QPalette_highlight(void* self) {
	auto&& __r = ((QPalette*)self)->highlight();
	return (void*)&(__r);
}
void* QPalette_highlightedText(void* self) {
	auto&& __r = ((QPalette*)self)->highlightedText();
	return (void*)&(__r);
}
void* QPalette_link(void* self) {
	auto&& __r = ((QPalette*)self)->link();
	return (void*)&(__r);
}
void* QPalette_linkVisited(void* self) {
	auto&& __r = ((QPalette*)self)->linkVisited();
	return (void*)&(__r);
}
void* QPalette_placeholderText(void* self) {
	auto&& __r = ((QPalette*)self)->placeholderText();
	return (void*)&(__r);
}
void* QPalette_accent(void* self) {
	auto&& __r = ((QPalette*)self)->accent();
	return (void*)&(__r);
}
int QPalette_OperatorEqual(void* self, void* p) {
	auto&& __r = ((QPalette*)self)->operator==(*(QPalette*)p);
	return __r;
}
int QPalette_OperatorNotEqual(void* self, void* p) {
	auto&& __r = ((QPalette*)self)->operator!=(*(QPalette*)p);
	return __r;
}
int QPalette_isCopyOf(void* self, void* p) {
	auto&& __r = ((QPalette*)self)->isCopyOf(*(QPalette*)p);
	return __r;
}
long long QPalette_cacheKey(void* self) {
	auto&& __r = ((QPalette*)self)->cacheKey();
	return __r;
}
void* QPalette_resolve(void* self, void* other) {
	auto&& __r = ((QPalette*)self)->resolve(*(QPalette*)other);
	return new QPalette(__r);
}
unsigned long long QPalette_resolveMask(void* self) {
	auto&& __r = ((QPalette*)self)->resolveMask();
	return __r;
}
void QPalette_setResolveMask(void* self, unsigned long long mask) {
	((QPalette*)self)->setResolveMask(mask);
}
void* QPalette_toQVariant(void* self) {
	auto&& __r = QVariant::fromValue(*(QPalette*)self);
	return new QVariant(__r);
}
void QPalette_delete(void* self) { delete (QPalette*)self; }

}

#include <QtGui/qaccessible.h>
#include <QtCore/qobject.h>
#include <QtCore/qstring.h>
#include <QtGui/qaccessible_base.h>
#include "gen_qaccessibletextinsertevent.h"

extern "C" {

void* QAccessibleTextInsertEvent_base(void* self) { return (void*)static_cast<QAccessibleTextCursorEvent*>((QAccessibleTextInsertEvent*)self); }
void QAccessibleTextInsertEvent_chain(void* self, void** out) {
	QAccessibleTextInsertEvent* __p = (QAccessibleTextInsertEvent*)self;
	QAccessibleTextCursorEvent* __b0 = static_cast<QAccessibleTextCursorEvent*>(__p); out[0] = (void*)__b0;
	QAccessibleEvent* __b1 = static_cast<QAccessibleEvent*>(__b0); out[1] = (void*)__b1;
}
void* QAccessibleTextInsertEvent_new(void* obj, int position, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	return new QAccessibleTextInsertEvent((QObject*)obj, position, text_qstr);
}
void* QAccessibleTextInsertEvent_new2(void* iface, int position, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	return new QAccessibleTextInsertEvent((QAccessibleInterface*)iface, position, text_qstr);
}
char* QAccessibleTextInsertEvent_textInserted(void* self, long long* __out_len) {
	auto&& __r = ((QAccessibleTextInsertEvent*)self)->textInserted();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QAccessibleTextInsertEvent_changePosition(void* self) {
	auto&& __r = ((QAccessibleTextInsertEvent*)self)->changePosition();
	return __r;
}
void QAccessibleTextInsertEvent_delete(void* self) { delete (QAccessibleTextInsertEvent*)self; }

}

#include <QtGui/qdrag.h>
#include <QtCore/qobject.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qmimedata.h>
#include <QtGui/qpixmap.h>
#include <QtCore/qpoint.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qvariant.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qdrag.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QDrag__Go : public QDrag {
public:
	long long __goId;
	unsigned char __goOn[8] = {};
	template<typename... A> QDrag__Go(long long id, A&&... a) : QDrag(std::forward<A>(a)...), __goId(id) {}
	using QDrag::sender;
	using QDrag::senderSignalIndex;
	using QDrag::receivers;
	using QDrag::isSignalConnected;
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QDrag::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QDrag::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QDrag::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QDrag::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QDrag::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QDrag::metaObject();
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QDrag::event(event);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QDrag::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QDrag::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QDrag::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QDrag::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QDrag::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QDrag::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QDrag;
static std::mutex __connsMu_QDrag;

extern "C" {

void* QDrag_base(void* self) { return (void*)static_cast<QObject*>((QDrag*)self); }
void QDrag__set_override(void* self, int idx) { ((QDrag__Go*)self)->__goOn[idx] = 1; }
void* QDrag_fromQObject(void* o) { return (void*)qobject_cast<QDrag*>((QObject*)o); }
void* QDrag_new(long long __goid, void* dragSource) {
	return new QDrag__Go(__goid, (QObject*)dragSource);
}
char* QDrag_tr(const char* s, long long* __out_len) {
	auto&& __r = QDrag::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QDrag_setMimeData(void* self, void* data) {
	((QDrag*)self)->setMimeData((QMimeData*)data);
}
void* QDrag_mimeData(void* self) {
	auto&& __r = ((QDrag*)self)->mimeData();
	return (void*)(__r);
}
void QDrag_setPixmap(void* self, void* arg1) {
	((QDrag*)self)->setPixmap(*(QPixmap*)arg1);
}
void* QDrag_pixmap(void* self) {
	auto&& __r = ((QDrag*)self)->pixmap();
	return new QPixmap(__r);
}
void QDrag_setHotSpot(void* self, void* hotspot) {
	((QDrag*)self)->setHotSpot(*(QPoint*)hotspot);
}
void* QDrag_hotSpot(void* self) {
	auto&& __r = ((QDrag*)self)->hotSpot();
	return new QPoint(__r);
}
void* QDrag_source(void* self) {
	auto&& __r = ((QDrag*)self)->source();
	return (void*)(__r);
}
void* QDrag_target(void* self) {
	auto&& __r = ((QDrag*)self)->target();
	return (void*)(__r);
}
int QDrag_exec(void* self) {
	auto&& __r = ((QDrag*)self)->exec();
	return (int)(__r);
}
int QDrag_exec2(void* self, int supportedActions, int defaultAction) {
	auto&& __r = ((QDrag*)self)->exec(QFlags<Qt::DropAction>(supportedActions), (Qt::DropAction)(defaultAction));
	return (int)(__r);
}
void QDrag_setDragCursor(void* self, void* cursor, int action) {
	((QDrag*)self)->setDragCursor(*(QPixmap*)cursor, (Qt::DropAction)(action));
}
void* QDrag_dragCursor(void* self, int action) {
	auto&& __r = ((QDrag*)self)->dragCursor((Qt::DropAction)(action));
	return new QPixmap(__r);
}
int QDrag_supportedActions(void* self) {
	auto&& __r = ((QDrag*)self)->supportedActions();
	return (int)(__r);
}
int QDrag_defaultAction(void* self) {
	auto&& __r = ((QDrag*)self)->defaultAction();
	return (int)(__r);
}
void QDrag_cancel() {
	QDrag::cancel();
}
void QDrag_actionChanged(void* self, int action) {
	((QDrag*)self)->actionChanged((Qt::DropAction)(action));
}
void QDrag_targetChanged(void* self, void* newTarget) {
	((QDrag*)self)->targetChanged((QObject*)newTarget);
}
void* QDrag_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QDrag::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QDrag_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QDrag::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QDrag_sender(void* self) {
	auto&& __r = ((QDrag__Go*)self)->sender();
	return (void*)(__r);
}
int QDrag_senderSignalIndex(void* self) {
	auto&& __r = ((QDrag__Go*)self)->senderSignalIndex();
	return __r;
}
int QDrag_receivers(void* self, const char* signal) {
	auto&& __r = ((QDrag__Go*)self)->receivers(signal);
	return __r;
}
int QDrag_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QDrag__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QDrag_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QDrag::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QDrag_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QDrag::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QDrag_exec3(void* self, int supportedActions) {
	auto&& __r = ((QDrag*)self)->exec(QFlags<Qt::DropAction>(supportedActions));
	return (int)(__r);
}
void* QDrag_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QDrag::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QDrag_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QDrag::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QDrag_metaObject_default(void* self) {
	auto&& __r = ((QDrag*)self)->QDrag::metaObject();
	return (void*)(__r);
}
int QDrag_event_default(void* self, void* event) {
	auto&& __r = ((QDrag*)self)->QDrag::event((QEvent*)event);
	return __r;
}
int QDrag_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QDrag*)self)->QDrag::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QDrag_timerEvent_default(void* self, void* event) {
	((QDrag__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QDrag_childEvent_default(void* self, void* event) {
	((QDrag__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QDrag_customEvent_default(void* self, void* event) {
	((QDrag__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QDrag_connectNotify_default(void* self, void* signal) {
	((QDrag__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QDrag_disconnectNotify_default(void* self, void* signal) {
	((QDrag__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QDrag_connect_actionChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDrag*)self, &QDrag::actionChanged, (QDrag*)self, [__id](Qt::DropAction action){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, (int)(action));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDrag);
	__conns_QDrag[__id] = __c;
}
void QDrag_connect_targetChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDrag*)self, &QDrag::targetChanged, (QDrag*)self, [__id](QObject* newTarget){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)newTarget);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDrag);
	__conns_QDrag[__id] = __c;
}
void QDrag_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDrag*)self, &QDrag::destroyed, (QDrag*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDrag);
	__conns_QDrag[__id] = __c;
}
void QDrag_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDrag*)self, &QDrag::destroyed, (QDrag*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDrag);
	__conns_QDrag[__id] = __c;
}
void QDrag_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDrag*)self, &QDrag::objectNameChanged, (QDrag*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDrag);
	__conns_QDrag[__id] = __c;
}
void QDrag__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QDrag);
		auto __it = __conns_QDrag.find(id);
		if (__it == __conns_QDrag.end()) return;
		__c = __it->second;
		__conns_QDrag.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QDrag__Go_new(long long id, void* parent) { return new QDrag__Go(id, (QObject*)parent); }
void QDrag_delete(void* self) { delete (QDrag*)self; }

}

#include <QtGui/qpageranges.h>
#include <QtCore/qstring.h>
#include "gen_qpageranges.h"

extern "C" {

void* QPageRanges_new() {
	return new QPageRanges();
}
void* QPageRanges_new2(void* other) {
	return new QPageRanges(*(QPageRanges*)other);
}
void* QPageRanges_OperatorAssign(void* self, void* other) {
	auto&& __r = ((QPageRanges*)self)->operator=(*(QPageRanges*)other);
	return (void*)&(__r);
}
void QPageRanges_swap(void* self, void* other) {
	((QPageRanges*)self)->swap(*(QPageRanges*)other);
}
void QPageRanges_addPage(void* self, int pageNumber) {
	((QPageRanges*)self)->addPage(pageNumber);
}
void QPageRanges_addRange(void* self, int from, int to) {
	((QPageRanges*)self)->addRange(from, to);
}
void** QPageRanges_toRangeList(void* self, long long* __out_len) {
	auto&& __r = ((QPageRanges*)self)->toRangeList();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QPageRanges::Range(__e);
	return __arr;
}
void QPageRanges_clear(void* self) {
	((QPageRanges*)self)->clear();
}
char* QPageRanges_toString(void* self, long long* __out_len) {
	auto&& __r = ((QPageRanges*)self)->toString();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QPageRanges_fromString(const char* ranges, long long ranges_len) {
	QString ranges_qstr = QString::fromUtf8(ranges, ranges_len);
	auto&& __r = QPageRanges::fromString(ranges_qstr);
	return new QPageRanges(__r);
}
int QPageRanges_contains(void* self, int pageNumber) {
	auto&& __r = ((QPageRanges*)self)->contains(pageNumber);
	return __r;
}
int QPageRanges_isEmpty(void* self) {
	auto&& __r = ((QPageRanges*)self)->isEmpty();
	return __r;
}
int QPageRanges_firstPage(void* self) {
	auto&& __r = ((QPageRanges*)self)->firstPage();
	return __r;
}
int QPageRanges_lastPage(void* self) {
	auto&& __r = ((QPageRanges*)self)->lastPage();
	return __r;
}
void QPageRanges_detach(void* self) {
	((QPageRanges*)self)->detach();
}
int QPageRanges_OperatorEqual(void* self, void* rhs) {
	auto&& __r = (*(QPageRanges*)self) == (*(QPageRanges*)rhs);
	return __r;
}
int QPageRanges_OperatorNotEqual(void* self, void* rhs) {
	auto&& __r = (*(QPageRanges*)self) != (*(QPageRanges*)rhs);
	return __r;
}
void QPageRanges_delete(void* self) { delete (QPageRanges*)self; }

}

#include <QtGui/qstandarditemmodel.h>
#include <QtCore/qstring.h>
#include <QtGui/qicon.h>
#include <QtCore/qvariant.h>
#include <QtCore/qabstractitemmodel.h>
#include <QtCore/qsize.h>
#include <QtGui/qfont.h>
#include <QtCore/qnamespace.h>
#include <QtGui/qbrush.h>
#include <QtCore/qdatastream.h>
#include "gen_qstandarditem.h"
#include "gen_purego_cb.h"

class QStandardItem__Go : public QStandardItem {
public:
	long long __goId;
	unsigned char __goOn[7] = {};
	template<typename... A> QStandardItem__Go(long long id, A&&... a) : QStandardItem(std::forward<A>(a)...), __goId(id) {}
	using QStandardItem::emitDataChanged;
	QVariant data(int role) const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, -8837375942098466299LL, role, &__h); if (__h) return *(QVariant*)__r; }
		return QStandardItem::data(role);
	}
	void multiData(QModelRoleDataSpan roleDataSpan) const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 731311060944096642LL, (void*)&roleDataSpan, &__h); if (__h) return; }
		QStandardItem::multiData(roleDataSpan);
	}
	void setData(const QVariant& value, int role) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_void_voidp_int) { ((void(*)(long long, long long, void*, int, int*))cppgo_virt_void_voidp_int)(__goId, 35191454575202219LL, (void*)&value, role, &__h); if (__h) return; }
		QStandardItem::setData(value, role);
	}
	QStandardItem* clone() const override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, -7032813281048616478LL, &__h); if (__h) return (QStandardItem*)__r; }
		return QStandardItem::clone();
	}
	int type() const override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -6371403872082420691LL, &__h); if (__h) return __r; }
		return QStandardItem::type();
	}
	void read(QDataStream& in) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5541207787354052101LL, (void*)&in, &__h); if (__h) return; }
		QStandardItem::read(in);
	}
	void write(QDataStream& out) const override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5099743077423337732LL, (void*)&out, &__h); if (__h) return; }
		QStandardItem::write(out);
	}
};

extern "C" {

void QStandardItem__set_override(void* self, int idx) { ((QStandardItem__Go*)self)->__goOn[idx] = 1; }
void* QStandardItem_new(long long __goid) {
	return new QStandardItem__Go(__goid);
}
void* QStandardItem_new2(long long __goid, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	return new QStandardItem__Go(__goid, text_qstr);
}
void* QStandardItem_new3(long long __goid, void* icon, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	return new QStandardItem__Go(__goid, *(QIcon*)icon, text_qstr);
}
void* QStandardItem_new4(long long __goid, int rows) {
	return new QStandardItem__Go(__goid, rows);
}
void* QStandardItem_new5(long long __goid, int rows, int columns) {
	return new QStandardItem__Go(__goid, rows, columns);
}
void* QStandardItem_data(void* self, int role) {
	auto&& __r = ((QStandardItem*)self)->data(role);
	return new QVariant(__r);
}
void QStandardItem_multiData(void* self, void* roleDataSpan) {
	((QStandardItem*)self)->multiData(*(QModelRoleDataSpan*)roleDataSpan);
}
void QStandardItem_setData(void* self, void* value, int role) {
	((QStandardItem*)self)->setData(*(QVariant*)value, role);
}
void QStandardItem_clearData(void* self) {
	((QStandardItem*)self)->clearData();
}
char* QStandardItem_text(void* self, long long* __out_len) {
	auto&& __r = ((QStandardItem*)self)->text();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QStandardItem_setText(void* self, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	((QStandardItem*)self)->setText(text_qstr);
}
void* QStandardItem_icon(void* self) {
	auto&& __r = ((QStandardItem*)self)->icon();
	return new QIcon(__r);
}
void QStandardItem_setIcon(void* self, void* icon) {
	((QStandardItem*)self)->setIcon(*(QIcon*)icon);
}
char* QStandardItem_toolTip(void* self, long long* __out_len) {
	auto&& __r = ((QStandardItem*)self)->toolTip();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QStandardItem_setToolTip(void* self, const char* toolTip, long long toolTip_len) {
	QString toolTip_qstr = QString::fromUtf8(toolTip, toolTip_len);
	((QStandardItem*)self)->setToolTip(toolTip_qstr);
}
char* QStandardItem_statusTip(void* self, long long* __out_len) {
	auto&& __r = ((QStandardItem*)self)->statusTip();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QStandardItem_setStatusTip(void* self, const char* statusTip, long long statusTip_len) {
	QString statusTip_qstr = QString::fromUtf8(statusTip, statusTip_len);
	((QStandardItem*)self)->setStatusTip(statusTip_qstr);
}
char* QStandardItem_whatsThis(void* self, long long* __out_len) {
	auto&& __r = ((QStandardItem*)self)->whatsThis();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QStandardItem_setWhatsThis(void* self, const char* whatsThis, long long whatsThis_len) {
	QString whatsThis_qstr = QString::fromUtf8(whatsThis, whatsThis_len);
	((QStandardItem*)self)->setWhatsThis(whatsThis_qstr);
}
void* QStandardItem_sizeHint(void* self) {
	auto&& __r = ((QStandardItem*)self)->sizeHint();
	return new QSize(__r);
}
void QStandardItem_setSizeHint(void* self, void* sizeHint) {
	((QStandardItem*)self)->setSizeHint(*(QSize*)sizeHint);
}
void* QStandardItem_font(void* self) {
	auto&& __r = ((QStandardItem*)self)->font();
	return new QFont(__r);
}
void QStandardItem_setFont(void* self, void* font) {
	((QStandardItem*)self)->setFont(*(QFont*)font);
}
int QStandardItem_textAlignment(void* self) {
	auto&& __r = ((QStandardItem*)self)->textAlignment();
	return (int)(__r);
}
void QStandardItem_setTextAlignment(void* self, int textAlignment) {
	((QStandardItem*)self)->setTextAlignment(QFlags<Qt::AlignmentFlag>(textAlignment));
}
void* QStandardItem_background(void* self) {
	auto&& __r = ((QStandardItem*)self)->background();
	return new QBrush(__r);
}
void QStandardItem_setBackground(void* self, void* brush) {
	((QStandardItem*)self)->setBackground(*(QBrush*)brush);
}
void* QStandardItem_foreground(void* self) {
	auto&& __r = ((QStandardItem*)self)->foreground();
	return new QBrush(__r);
}
void QStandardItem_setForeground(void* self, void* brush) {
	((QStandardItem*)self)->setForeground(*(QBrush*)brush);
}
int QStandardItem_checkState(void* self) {
	auto&& __r = ((QStandardItem*)self)->checkState();
	return (int)(__r);
}
void QStandardItem_setCheckState(void* self, int checkState) {
	((QStandardItem*)self)->setCheckState((Qt::CheckState)(checkState));
}
char* QStandardItem_accessibleText(void* self, long long* __out_len) {
	auto&& __r = ((QStandardItem*)self)->accessibleText();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QStandardItem_setAccessibleText(void* self, const char* accessibleText, long long accessibleText_len) {
	QString accessibleText_qstr = QString::fromUtf8(accessibleText, accessibleText_len);
	((QStandardItem*)self)->setAccessibleText(accessibleText_qstr);
}
char* QStandardItem_accessibleDescription(void* self, long long* __out_len) {
	auto&& __r = ((QStandardItem*)self)->accessibleDescription();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QStandardItem_setAccessibleDescription(void* self, const char* accessibleDescription, long long accessibleDescription_len) {
	QString accessibleDescription_qstr = QString::fromUtf8(accessibleDescription, accessibleDescription_len);
	((QStandardItem*)self)->setAccessibleDescription(accessibleDescription_qstr);
}
int QStandardItem_flags(void* self) {
	auto&& __r = ((QStandardItem*)self)->flags();
	return (int)(__r);
}
void QStandardItem_setFlags(void* self, int flags) {
	((QStandardItem*)self)->setFlags(QFlags<Qt::ItemFlag>(flags));
}
int QStandardItem_isEnabled(void* self) {
	auto&& __r = ((QStandardItem*)self)->isEnabled();
	return __r;
}
void QStandardItem_setEnabled(void* self, int enabled) {
	((QStandardItem*)self)->setEnabled(enabled);
}
int QStandardItem_isEditable(void* self) {
	auto&& __r = ((QStandardItem*)self)->isEditable();
	return __r;
}
void QStandardItem_setEditable(void* self, int editable) {
	((QStandardItem*)self)->setEditable(editable);
}
int QStandardItem_isSelectable(void* self) {
	auto&& __r = ((QStandardItem*)self)->isSelectable();
	return __r;
}
void QStandardItem_setSelectable(void* self, int selectable) {
	((QStandardItem*)self)->setSelectable(selectable);
}
int QStandardItem_isCheckable(void* self) {
	auto&& __r = ((QStandardItem*)self)->isCheckable();
	return __r;
}
void QStandardItem_setCheckable(void* self, int checkable) {
	((QStandardItem*)self)->setCheckable(checkable);
}
int QStandardItem_isAutoTristate(void* self) {
	auto&& __r = ((QStandardItem*)self)->isAutoTristate();
	return __r;
}
void QStandardItem_setAutoTristate(void* self, int tristate) {
	((QStandardItem*)self)->setAutoTristate(tristate);
}
int QStandardItem_isUserTristate(void* self) {
	auto&& __r = ((QStandardItem*)self)->isUserTristate();
	return __r;
}
void QStandardItem_setUserTristate(void* self, int tristate) {
	((QStandardItem*)self)->setUserTristate(tristate);
}
int QStandardItem_isDragEnabled(void* self) {
	auto&& __r = ((QStandardItem*)self)->isDragEnabled();
	return __r;
}
void QStandardItem_setDragEnabled(void* self, int dragEnabled) {
	((QStandardItem*)self)->setDragEnabled(dragEnabled);
}
int QStandardItem_isDropEnabled(void* self) {
	auto&& __r = ((QStandardItem*)self)->isDropEnabled();
	return __r;
}
void QStandardItem_setDropEnabled(void* self, int dropEnabled) {
	((QStandardItem*)self)->setDropEnabled(dropEnabled);
}
void* QStandardItem_parent(void* self) {
	auto&& __r = ((QStandardItem*)self)->parent();
	return (void*)(__r);
}
int QStandardItem_row(void* self) {
	auto&& __r = ((QStandardItem*)self)->row();
	return __r;
}
int QStandardItem_column(void* self) {
	auto&& __r = ((QStandardItem*)self)->column();
	return __r;
}
void* QStandardItem_index(void* self) {
	auto&& __r = ((QStandardItem*)self)->index();
	return new QModelIndex(__r);
}
void* QStandardItem_model(void* self) {
	auto&& __r = ((QStandardItem*)self)->model();
	return (void*)(__r);
}
int QStandardItem_rowCount(void* self) {
	auto&& __r = ((QStandardItem*)self)->rowCount();
	return __r;
}
void QStandardItem_setRowCount(void* self, int rows) {
	((QStandardItem*)self)->setRowCount(rows);
}
int QStandardItem_columnCount(void* self) {
	auto&& __r = ((QStandardItem*)self)->columnCount();
	return __r;
}
void QStandardItem_setColumnCount(void* self, int columns) {
	((QStandardItem*)self)->setColumnCount(columns);
}
int QStandardItem_hasChildren(void* self) {
	auto&& __r = ((QStandardItem*)self)->hasChildren();
	return __r;
}
void* QStandardItem_child(void* self, int row) {
	auto&& __r = ((QStandardItem*)self)->child(row);
	return (void*)(__r);
}
void QStandardItem_setChild(void* self, int row, int column, void* item) {
	((QStandardItem*)self)->setChild(row, column, (QStandardItem*)item);
}
void QStandardItem_setChild2(void* self, int row, void* item) {
	((QStandardItem*)self)->setChild(row, (QStandardItem*)item);
}
void QStandardItem_insertRow(void* self, int row, void** items, long long items_len) {
	QList<QStandardItem*> items_list; items_list.reserve(items_len);
	for (long long __i = 0; __i < items_len; __i++) items_list.append((QStandardItem*)items[__i]);
	((QStandardItem*)self)->insertRow(row, items_list);
}
void QStandardItem_insertColumn(void* self, int column, void** items, long long items_len) {
	QList<QStandardItem*> items_list; items_list.reserve(items_len);
	for (long long __i = 0; __i < items_len; __i++) items_list.append((QStandardItem*)items[__i]);
	((QStandardItem*)self)->insertColumn(column, items_list);
}
void QStandardItem_insertRows(void* self, int row, void** items, long long items_len) {
	QList<QStandardItem*> items_list; items_list.reserve(items_len);
	for (long long __i = 0; __i < items_len; __i++) items_list.append((QStandardItem*)items[__i]);
	((QStandardItem*)self)->insertRows(row, items_list);
}
void QStandardItem_insertRows2(void* self, int row, int count) {
	((QStandardItem*)self)->insertRows(row, count);
}
void QStandardItem_insertColumns(void* self, int column, int count) {
	((QStandardItem*)self)->insertColumns(column, count);
}
void QStandardItem_removeRow(void* self, int row) {
	((QStandardItem*)self)->removeRow(row);
}
void QStandardItem_removeColumn(void* self, int column) {
	((QStandardItem*)self)->removeColumn(column);
}
void QStandardItem_removeRows(void* self, int row, int count) {
	((QStandardItem*)self)->removeRows(row, count);
}
void QStandardItem_removeColumns(void* self, int column, int count) {
	((QStandardItem*)self)->removeColumns(column, count);
}
void QStandardItem_appendRow(void* self, void** items, long long items_len) {
	QList<QStandardItem*> items_list; items_list.reserve(items_len);
	for (long long __i = 0; __i < items_len; __i++) items_list.append((QStandardItem*)items[__i]);
	((QStandardItem*)self)->appendRow(items_list);
}
void QStandardItem_appendRows(void* self, void** items, long long items_len) {
	QList<QStandardItem*> items_list; items_list.reserve(items_len);
	for (long long __i = 0; __i < items_len; __i++) items_list.append((QStandardItem*)items[__i]);
	((QStandardItem*)self)->appendRows(items_list);
}
void QStandardItem_appendColumn(void* self, void** items, long long items_len) {
	QList<QStandardItem*> items_list; items_list.reserve(items_len);
	for (long long __i = 0; __i < items_len; __i++) items_list.append((QStandardItem*)items[__i]);
	((QStandardItem*)self)->appendColumn(items_list);
}
void QStandardItem_insertRow2(void* self, int row, void* item) {
	((QStandardItem*)self)->insertRow(row, (QStandardItem*)item);
}
void QStandardItem_appendRow2(void* self, void* item) {
	((QStandardItem*)self)->appendRow((QStandardItem*)item);
}
void* QStandardItem_takeChild(void* self, int row) {
	auto&& __r = ((QStandardItem*)self)->takeChild(row);
	return (void*)(__r);
}
void** QStandardItem_takeRow(void* self, int row, long long* __out_len) {
	auto&& __r = ((QStandardItem*)self)->takeRow(row);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
void** QStandardItem_takeColumn(void* self, int column, long long* __out_len) {
	auto&& __r = ((QStandardItem*)self)->takeColumn(column);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
void QStandardItem_sortChildren(void* self, int column) {
	((QStandardItem*)self)->sortChildren(column);
}
void* QStandardItem_clone(void* self) {
	auto&& __r = ((QStandardItem*)self)->clone();
	return (void*)(__r);
}
int QStandardItem_type(void* self) {
	auto&& __r = ((QStandardItem*)self)->type();
	return __r;
}
void QStandardItem_read(void* self, void* in) {
	((QStandardItem*)self)->read(*(QDataStream*)in);
}
void QStandardItem_write(void* self, void* out) {
	((QStandardItem*)self)->write(*(QDataStream*)out);
}
int QStandardItem_OperatorLesser(void* self, void* other) {
	auto&& __r = ((QStandardItem*)self)->operator<(*(QStandardItem*)other);
	return __r;
}
void QStandardItem_emitDataChanged(void* self) {
	((QStandardItem__Go*)self)->emitDataChanged();
}
void* QStandardItem_child2(void* self, int row, int column) {
	auto&& __r = ((QStandardItem*)self)->child(row, column);
	return (void*)(__r);
}
void* QStandardItem_takeChild2(void* self, int row, int column) {
	auto&& __r = ((QStandardItem*)self)->takeChild(row, column);
	return (void*)(__r);
}
void QStandardItem_sortChildren2(void* self, int column, int order) {
	((QStandardItem*)self)->sortChildren(column, (Qt::SortOrder)(order));
}
void* QStandardItem_data_default(void* self, int role) {
	auto&& __r = ((QStandardItem*)self)->QStandardItem::data(role);
	return new QVariant(__r);
}
void QStandardItem_multiData_default(void* self, void* roleDataSpan) {
	((QStandardItem*)self)->QStandardItem::multiData(*(QModelRoleDataSpan*)roleDataSpan);
}
void QStandardItem_setData_default(void* self, void* value, int role) {
	((QStandardItem*)self)->QStandardItem::setData(*(QVariant*)value, role);
}
void* QStandardItem_clone_default(void* self) {
	auto&& __r = ((QStandardItem*)self)->QStandardItem::clone();
	return (void*)(__r);
}
int QStandardItem_type_default(void* self) {
	auto&& __r = ((QStandardItem*)self)->QStandardItem::type();
	return __r;
}
void QStandardItem_read_default(void* self, void* in) {
	((QStandardItem*)self)->QStandardItem::read(*(QDataStream*)in);
}
void QStandardItem_write_default(void* self, void* out) {
	((QStandardItem*)self)->QStandardItem::write(*(QDataStream*)out);
}
void* QStandardItem__Go_new(long long id) { return new QStandardItem__Go(id); }
void QStandardItem_delete(void* self) { delete (QStandardItem*)self; }

}

#include <QtGui/qwindowsmimeconverter.h>
#include <QtCore/qstring.h>
#include <QtCore/qmimedata.h>
#include <QtCore/qvariant.h>
#include <QtCore/qmetatype.h>
#include "gen_qwindowsmimeconverter.h"

extern "C" {

int QWindowsMimeConverter_registerMimeType(const char* mimeType, long long mimeType_len) {
	QString mimeType_qstr = QString::fromUtf8(mimeType, mimeType_len);
	auto&& __r = QWindowsMimeConverter::registerMimeType(mimeType_qstr);
	return __r;
}
void QWindowsMimeConverter_delete(void* self) { delete (QWindowsMimeConverter*)self; }

}

#include <QtWidgets/qstyleoption.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qrect.h>
#include <QtCore/qstring.h>
#include <QtGui/qicon.h>
#include <QtGui/qfont.h>
#include "gen_qstyleoptionmenuitem.h"

extern "C" {

void* QStyleOptionMenuItem_base(void* self) { return (void*)static_cast<QStyleOption*>((QStyleOptionMenuItem*)self); }
void* QStyleOptionMenuItem_new() {
	return new QStyleOptionMenuItem();
}
void* QStyleOptionMenuItem_new2(void* other) {
	return new QStyleOptionMenuItem(*(QStyleOptionMenuItem*)other);
}
void* QStyleOptionMenuItem_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QStyleOptionMenuItem*)self)->operator=(*(QStyleOptionMenuItem*)arg1);
	return (void*)&(__r);
}
int QStyleOptionMenuItem_menuItemType_get(void* self) {
	auto __r = ((QStyleOptionMenuItem*)self)->menuItemType;
	return (int)(__r);
}
void QStyleOptionMenuItem_menuItemType_set(void* self, int menuItemType) {
	((QStyleOptionMenuItem*)self)->menuItemType = (QStyleOptionMenuItem::MenuItemType)(menuItemType);
}
int QStyleOptionMenuItem_checkType_get(void* self) {
	auto __r = ((QStyleOptionMenuItem*)self)->checkType;
	return (int)(__r);
}
void QStyleOptionMenuItem_checkType_set(void* self, int checkType) {
	((QStyleOptionMenuItem*)self)->checkType = (QStyleOptionMenuItem::CheckType)(checkType);
}
int QStyleOptionMenuItem_checked_get(void* self) {
	auto __r = ((QStyleOptionMenuItem*)self)->checked;
	return __r;
}
void QStyleOptionMenuItem_checked_set(void* self, int checked) {
	((QStyleOptionMenuItem*)self)->checked = checked;
}
int QStyleOptionMenuItem_menuHasCheckableItems_get(void* self) {
	auto __r = ((QStyleOptionMenuItem*)self)->menuHasCheckableItems;
	return __r;
}
void QStyleOptionMenuItem_menuHasCheckableItems_set(void* self, int menuHasCheckableItems) {
	((QStyleOptionMenuItem*)self)->menuHasCheckableItems = menuHasCheckableItems;
}
void* QStyleOptionMenuItem_menuRect_get(void* self) {
	auto __r = ((QStyleOptionMenuItem*)self)->menuRect;
	return new QRect(__r);
}
void QStyleOptionMenuItem_menuRect_set(void* self, void* menuRect) {
	((QStyleOptionMenuItem*)self)->menuRect = *(QRect*)menuRect;
}
char* QStyleOptionMenuItem_text_get(void* self, long long* __out_len) {
	auto __r = ((QStyleOptionMenuItem*)self)->text;
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QStyleOptionMenuItem_text_set(void* self, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	((QStyleOptionMenuItem*)self)->text = text_qstr;
}
void* QStyleOptionMenuItem_icon_get(void* self) {
	auto __r = ((QStyleOptionMenuItem*)self)->icon;
	return new QIcon(__r);
}
void QStyleOptionMenuItem_icon_set(void* self, void* icon) {
	((QStyleOptionMenuItem*)self)->icon = *(QIcon*)icon;
}
int QStyleOptionMenuItem_maxIconWidth_get(void* self) {
	auto __r = ((QStyleOptionMenuItem*)self)->maxIconWidth;
	return __r;
}
void QStyleOptionMenuItem_maxIconWidth_set(void* self, int maxIconWidth) {
	((QStyleOptionMenuItem*)self)->maxIconWidth = maxIconWidth;
}
int QStyleOptionMenuItem_reservedShortcutWidth_get(void* self) {
	auto __r = ((QStyleOptionMenuItem*)self)->reservedShortcutWidth;
	return __r;
}
void QStyleOptionMenuItem_reservedShortcutWidth_set(void* self, int reservedShortcutWidth) {
	((QStyleOptionMenuItem*)self)->reservedShortcutWidth = reservedShortcutWidth;
}
void* QStyleOptionMenuItem_font_get(void* self) {
	auto __r = ((QStyleOptionMenuItem*)self)->font;
	return new QFont(__r);
}
void QStyleOptionMenuItem_font_set(void* self, void* font) {
	((QStyleOptionMenuItem*)self)->font = *(QFont*)font;
}
void QStyleOptionMenuItem_delete(void* self) { delete (QStyleOptionMenuItem*)self; }

}

#include <QtWidgets/qlayoutitem.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qsize.h>
#include <QtCore/qrect.h>
#include <QtWidgets/qwidget.h>
#include <QtWidgets/qlayout.h>
#include <QtWidgets/qsizepolicy.h>
#include "gen_qlayoutitem.h"
#include "gen_purego_cb.h"

class QLayoutItem__Go : public QLayoutItem {
public:
	long long __goId;
	unsigned char __goOn[15] = {};
	template<typename... A> QLayoutItem__Go(long long id, A&&... a) : QLayoutItem(std::forward<A>(a)...), __goId(id) {}
	QSize sizeHint() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 729690261457684101LL, &__h); if (__h) return *(QSize*)__r; }
		return QSize{};
	}
	QSize minimumSize() const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, -7246748916541192690LL, &__h); if (__h) return *(QSize*)__r; }
		return QSize{};
	}
	QSize maximumSize() const override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 4158099606088301136LL, &__h); if (__h) return *(QSize*)__r; }
		return QSize{};
	}
	QFlags<Qt::Orientation> expandingDirections() const override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, 415665469015095569LL, &__h); if (__h) return QFlags<Qt::Orientation>(__r); }
		return QFlags<Qt::Orientation>{};
	}
	void setGeometry(const QRect& arg1) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5684820848726969227LL, (void*)&arg1, &__h); if (__h) return; }
		return;
	}
	QRect geometry() const override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1050490530495513297LL, &__h); if (__h) return *(QRect*)__r; }
		return QRect{};
	}
	bool isEmpty() const override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -626822430251098498LL, &__h); if (__h) return __r; }
		return bool{};
	}
	bool hasHeightForWidth() const override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8266762397995423915LL, &__h); if (__h) return __r; }
		return QLayoutItem::hasHeightForWidth();
	}
	int heightForWidth(int arg1) const override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 4311534060495737339LL, arg1, &__h); if (__h) return __r; }
		return QLayoutItem::heightForWidth(arg1);
	}
	int minimumHeightForWidth(int arg1) const override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 5037837229528203129LL, arg1, &__h); if (__h) return __r; }
		return QLayoutItem::minimumHeightForWidth(arg1);
	}
	void invalidate() override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 6991341317990632368LL, &__h); if (__h) return; }
		QLayoutItem::invalidate();
	}
	QWidget* widget() const override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, -4575937485689502171LL, &__h); if (__h) return (QWidget*)__r; }
		return QLayoutItem::widget();
	}
	QLayout* layout() override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, -7183390618627992737LL, &__h); if (__h) return (QLayout*)__r; }
		return QLayoutItem::layout();
	}
	QSpacerItem* spacerItem() override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 2272240669359776278LL, &__h); if (__h) return (QSpacerItem*)__r; }
		return QLayoutItem::spacerItem();
	}
	QFlags<QSizePolicy::ControlType> controlTypes() const override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, 3278947275714896887LL, &__h); if (__h) return QFlags<QSizePolicy::ControlType>(__r); }
		return QLayoutItem::controlTypes();
	}
};

extern "C" {

void QLayoutItem__set_override(void* self, int idx) { ((QLayoutItem__Go*)self)->__goOn[idx] = 1; }
void* QLayoutItem_sizeHint(void* self) {
	auto&& __r = ((QLayoutItem*)self)->sizeHint();
	return new QSize(__r);
}
void* QLayoutItem_minimumSize(void* self) {
	auto&& __r = ((QLayoutItem*)self)->minimumSize();
	return new QSize(__r);
}
void* QLayoutItem_maximumSize(void* self) {
	auto&& __r = ((QLayoutItem*)self)->maximumSize();
	return new QSize(__r);
}
int QLayoutItem_expandingDirections(void* self) {
	auto&& __r = ((QLayoutItem*)self)->expandingDirections();
	return (int)(__r);
}
void QLayoutItem_setGeometry(void* self, void* arg1) {
	((QLayoutItem*)self)->setGeometry(*(QRect*)arg1);
}
void* QLayoutItem_geometry(void* self) {
	auto&& __r = ((QLayoutItem*)self)->geometry();
	return new QRect(__r);
}
int QLayoutItem_isEmpty(void* self) {
	auto&& __r = ((QLayoutItem*)self)->isEmpty();
	return __r;
}
int QLayoutItem_hasHeightForWidth(void* self) {
	auto&& __r = ((QLayoutItem*)self)->hasHeightForWidth();
	return __r;
}
int QLayoutItem_heightForWidth(void* self, int arg1) {
	auto&& __r = ((QLayoutItem*)self)->heightForWidth(arg1);
	return __r;
}
int QLayoutItem_minimumHeightForWidth(void* self, int arg1) {
	auto&& __r = ((QLayoutItem*)self)->minimumHeightForWidth(arg1);
	return __r;
}
void QLayoutItem_invalidate(void* self) {
	((QLayoutItem*)self)->invalidate();
}
void* QLayoutItem_widget(void* self) {
	auto&& __r = ((QLayoutItem*)self)->widget();
	return (void*)(__r);
}
void* QLayoutItem_layout(void* self) {
	auto&& __r = ((QLayoutItem*)self)->layout();
	return (void*)(__r);
}
void* QLayoutItem_spacerItem(void* self) {
	auto&& __r = ((QLayoutItem*)self)->spacerItem();
	return (void*)(__r);
}
int QLayoutItem_alignment(void* self) {
	auto&& __r = ((QLayoutItem*)self)->alignment();
	return (int)(__r);
}
void QLayoutItem_setAlignment(void* self, int a) {
	((QLayoutItem*)self)->setAlignment(QFlags<Qt::AlignmentFlag>(a));
}
int QLayoutItem_controlTypes(void* self) {
	auto&& __r = ((QLayoutItem*)self)->controlTypes();
	return (int)(__r);
}
void* QLayoutItem_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QLayoutItem*)self)->operator=(*(QLayoutItem*)arg1);
	return (void*)&(__r);
}
int QLayoutItem_hasHeightForWidth_default(void* self) {
	auto&& __r = ((QLayoutItem*)self)->QLayoutItem::hasHeightForWidth();
	return __r;
}
int QLayoutItem_heightForWidth_default(void* self, int arg1) {
	auto&& __r = ((QLayoutItem*)self)->QLayoutItem::heightForWidth(arg1);
	return __r;
}
int QLayoutItem_minimumHeightForWidth_default(void* self, int arg1) {
	auto&& __r = ((QLayoutItem*)self)->QLayoutItem::minimumHeightForWidth(arg1);
	return __r;
}
void QLayoutItem_invalidate_default(void* self) {
	((QLayoutItem*)self)->QLayoutItem::invalidate();
}
void* QLayoutItem_widget_default(void* self) {
	auto&& __r = ((QLayoutItem*)self)->QLayoutItem::widget();
	return (void*)(__r);
}
void* QLayoutItem_layout_default(void* self) {
	auto&& __r = ((QLayoutItem*)self)->QLayoutItem::layout();
	return (void*)(__r);
}
void* QLayoutItem_spacerItem_default(void* self) {
	auto&& __r = ((QLayoutItem*)self)->QLayoutItem::spacerItem();
	return (void*)(__r);
}
int QLayoutItem_controlTypes_default(void* self) {
	auto&& __r = ((QLayoutItem*)self)->QLayoutItem::controlTypes();
	return (int)(__r);
}
void* QLayoutItem__Go_new(long long id) { return new QLayoutItem__Go(id); }
void QLayoutItem_delete(void* self) { delete (QLayoutItem*)self; }

}

#include <QtWidgets/qdatetimeedit.h>
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
#include <QtWidgets/qabstractspinbox.h>
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
#include "gen_qdateedit.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QDateEdit__Go : public QDateEdit {
public:
	long long __goId;
	unsigned char __goOn[56] = {};
	template<typename... A> QDateEdit__Go(long long id, A&&... a) : QDateEdit(std::forward<A>(a)...), __goId(id) {}
	using QDateEdit::lineEdit;
	using QDateEdit::setLineEdit;
	using QDateEdit::updateMicroFocus;
	using QDateEdit::create;
	using QDateEdit::destroy;
	using QDateEdit::focusNextChild;
	using QDateEdit::focusPreviousChild;
	using QDateEdit::sender;
	using QDateEdit::senderSignalIndex;
	using QDateEdit::receivers;
	using QDateEdit::isSignalConnected;
	using QDateEdit::getDecodedMetricF;
	void cppgoBase_keyPressEvent(QKeyEvent* __a0) { QDateEdit::keyPressEvent(__a0); }
	void cppgoBase_wheelEvent(QWheelEvent* __a0) { QDateEdit::wheelEvent(__a0); }
	void cppgoBase_focusInEvent(QFocusEvent* __a0) { QDateEdit::focusInEvent(__a0); }
	bool cppgoBase_focusNextPrevChild(bool __a0) { return QDateEdit::focusNextPrevChild(__a0); }
	QValidator::State cppgoBase_validate(QString& __a0, int& __a1) { return QDateEdit::validate(__a0, __a1); }
	void cppgoBase_fixup(QString& __a0) { QDateEdit::fixup(__a0); }
	QDateTime cppgoBase_dateTimeFromText(const QString& __a0) { return QDateEdit::dateTimeFromText(__a0); }
	QString cppgoBase_textFromDateTime(const QDateTime& __a0) { return QDateEdit::textFromDateTime(__a0); }
	QFlags<QAbstractSpinBox::StepEnabledFlag> cppgoBase_stepEnabled() { return QDateEdit::stepEnabled(); }
	void cppgoBase_mousePressEvent(QMouseEvent* __a0) { QDateEdit::mousePressEvent(__a0); }
	void cppgoBase_paintEvent(QPaintEvent* __a0) { QDateEdit::paintEvent(__a0); }
	void cppgoBase_initStyleOption(QStyleOptionSpinBox* __a0) { QDateEdit::initStyleOption(__a0); }
	void cppgoBase_resizeEvent(QResizeEvent* __a0) { QDateEdit::resizeEvent(__a0); }
	void cppgoBase_keyReleaseEvent(QKeyEvent* __a0) { QDateEdit::keyReleaseEvent(__a0); }
	void cppgoBase_focusOutEvent(QFocusEvent* __a0) { QDateEdit::focusOutEvent(__a0); }
	void cppgoBase_contextMenuEvent(QContextMenuEvent* __a0) { QDateEdit::contextMenuEvent(__a0); }
	void cppgoBase_changeEvent(QEvent* __a0) { QDateEdit::changeEvent(__a0); }
	void cppgoBase_closeEvent(QCloseEvent* __a0) { QDateEdit::closeEvent(__a0); }
	void cppgoBase_hideEvent(QHideEvent* __a0) { QDateEdit::hideEvent(__a0); }
	void cppgoBase_mouseReleaseEvent(QMouseEvent* __a0) { QDateEdit::mouseReleaseEvent(__a0); }
	void cppgoBase_mouseMoveEvent(QMouseEvent* __a0) { QDateEdit::mouseMoveEvent(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QDateEdit::timerEvent(__a0); }
	void cppgoBase_showEvent(QShowEvent* __a0) { QDateEdit::showEvent(__a0); }
	void cppgoBase_mouseDoubleClickEvent(QMouseEvent* __a0) { QDateEdit::mouseDoubleClickEvent(__a0); }
	void cppgoBase_enterEvent(QEnterEvent* __a0) { QDateEdit::enterEvent(__a0); }
	void cppgoBase_leaveEvent(QEvent* __a0) { QDateEdit::leaveEvent(__a0); }
	void cppgoBase_moveEvent(QMoveEvent* __a0) { QDateEdit::moveEvent(__a0); }
	void cppgoBase_tabletEvent(QTabletEvent* __a0) { QDateEdit::tabletEvent(__a0); }
	void cppgoBase_actionEvent(QActionEvent* __a0) { QDateEdit::actionEvent(__a0); }
	void cppgoBase_dragEnterEvent(QDragEnterEvent* __a0) { QDateEdit::dragEnterEvent(__a0); }
	void cppgoBase_dragMoveEvent(QDragMoveEvent* __a0) { QDateEdit::dragMoveEvent(__a0); }
	void cppgoBase_dragLeaveEvent(QDragLeaveEvent* __a0) { QDateEdit::dragLeaveEvent(__a0); }
	void cppgoBase_dropEvent(QDropEvent* __a0) { QDateEdit::dropEvent(__a0); }
	int cppgoBase_metric(int __a0) { return QDateEdit::metric((QPaintDevice::PaintDeviceMetric)__a0); }
	void cppgoBase_initPainter(QPainter* __a0) { QDateEdit::initPainter(__a0); }
	QPaintDevice* cppgoBase_redirected(QPoint* __a0) { return QDateEdit::redirected(__a0); }
	QPainter* cppgoBase_sharedPainter() { return QDateEdit::sharedPainter(); }
	void cppgoBase_inputMethodEvent(QInputMethodEvent* __a0) { QDateEdit::inputMethodEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QDateEdit::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QDateEdit::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QDateEdit::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QDateEdit::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QDateEdit::metaObject();
	}
	QSize sizeHint() const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 729690261457684101LL, &__h); if (__h) return *(QSize*)__r; }
		return QDateEdit::sizeHint();
	}
	void clear() override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, -778567940572639070LL, &__h); if (__h) return; }
		QDateEdit::clear();
	}
	void stepBy(int steps) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 7085718643589954634LL, steps, &__h); if (__h) return; }
		QDateEdit::stepBy(steps);
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QDateEdit::event(event);
	}
	void keyPressEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1023791977679793185LL, (void*)event, &__h); if (__h) return; }
		QDateEdit::keyPressEvent(event);
	}
	void wheelEvent(QWheelEvent* event) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1919988340934511316LL, (void*)event, &__h); if (__h) return; }
		QDateEdit::wheelEvent(event);
	}
	void focusInEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 975043767033552552LL, (void*)event, &__h); if (__h) return; }
		QDateEdit::focusInEvent(event);
	}
	bool focusNextPrevChild(bool next) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 5898328517505784191LL, next, &__h); if (__h) return __r; }
		return QDateEdit::focusNextPrevChild(next);
	}
	QValidator::State validate(QString& input, int& pos) const override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_int_voidp_intp) { int __r = ((int(*)(long long, long long, void*, int*, int*))cppgo_virt_int_voidp_intp)(__goId, 6784179835071232555LL, (void*)&input, (int*)&pos, &__h); if (__h) return (QValidator::State)(__r); }
		return QDateEdit::validate(input, pos);
	}
	void fixup(QString& input) const override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4855206404746966715LL, (void*)&input, &__h); if (__h) return; }
		QDateEdit::fixup(input);
	}
	QDateTime dateTimeFromText(const QString& text) const override {
		int __h = 0;
		QByteArray text_ba = text.toUtf8();
		if (__goOn[11] && cppgo_virt_voidp_constcharp_longlong) { void* __r = ((void*(*)(long long, long long, const char*, long long, int*))cppgo_virt_voidp_constcharp_longlong)(__goId, -50441425836614729LL, (char*)text_ba.constData(), text_ba.size(), &__h); if (__h) return *(QDateTime*)__r; }
		return QDateEdit::dateTimeFromText(text);
	}
	QString textFromDateTime(const QDateTime& dt) const override {
		int __h = 0;
		if (__goOn[12] && cppgo_virtp_charp_voidp) {
			long long __len = 0;
			char* __s = ((char*(*)(long long, long long, void*, long long*, int*))cppgo_virtp_charp_voidp)(__goId, -3778148144965928937LL, (void*)&dt, &__len, &__h);
			if (__h) { QString __r = QString::fromUtf8(__s, __len); free(__s); return __r; }
		}
		return QDateEdit::textFromDateTime(dt);
	}
	QFlags<QAbstractSpinBox::StepEnabledFlag> stepEnabled() const override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, 5257289574516412244LL, &__h); if (__h) return QFlags<QAbstractSpinBox::StepEnabledFlag>(__r); }
		return QDateEdit::stepEnabled();
	}
	void mousePressEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -332090898078543319LL, (void*)event, &__h); if (__h) return; }
		QDateEdit::mousePressEvent(event);
	}
	void paintEvent(QPaintEvent* event) override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5290442163167458465LL, (void*)event, &__h); if (__h) return; }
		QDateEdit::paintEvent(event);
	}
	void initStyleOption(QStyleOptionSpinBox* option) const override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -488693769771308531LL, (void*)option, &__h); if (__h) return; }
		QDateEdit::initStyleOption(option);
	}
	QSize minimumSizeHint() const override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 7049264599837532491LL, &__h); if (__h) return *(QSize*)__r; }
		return QDateEdit::minimumSizeHint();
	}
	QVariant inputMethodQuery(Qt::InputMethodQuery arg1) const override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7345853097340709638LL, (int)(arg1), &__h); if (__h) return *(QVariant*)__r; }
		return QDateEdit::inputMethodQuery(arg1);
	}
	void resizeEvent(QResizeEvent* event) override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4302917285655014661LL, (void*)event, &__h); if (__h) return; }
		QDateEdit::resizeEvent(event);
	}
	void keyReleaseEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 772598475950050299LL, (void*)event, &__h); if (__h) return; }
		QDateEdit::keyReleaseEvent(event);
	}
	void focusOutEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -385596010576048993LL, (void*)event, &__h); if (__h) return; }
		QDateEdit::focusOutEvent(event);
	}
	void contextMenuEvent(QContextMenuEvent* event) override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6698043222227006015LL, (void*)event, &__h); if (__h) return; }
		QDateEdit::contextMenuEvent(event);
	}
	void changeEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4959376063826165943LL, (void*)event, &__h); if (__h) return; }
		QDateEdit::changeEvent(event);
	}
	void closeEvent(QCloseEvent* event) override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1989937050019464783LL, (void*)event, &__h); if (__h) return; }
		QDateEdit::closeEvent(event);
	}
	void hideEvent(QHideEvent* event) override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -199718707876719129LL, (void*)event, &__h); if (__h) return; }
		QDateEdit::hideEvent(event);
	}
	void mouseReleaseEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4415920049297795213LL, (void*)event, &__h); if (__h) return; }
		QDateEdit::mouseReleaseEvent(event);
	}
	void mouseMoveEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4255494004099806581LL, (void*)event, &__h); if (__h) return; }
		QDateEdit::mouseMoveEvent(event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QDateEdit::timerEvent(event);
	}
	void showEvent(QShowEvent* event) override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6825369174751916LL, (void*)event, &__h); if (__h) return; }
		QDateEdit::showEvent(event);
	}
	int devType() const override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8659679893868524044LL, &__h); if (__h) return __r; }
		return QDateEdit::devType();
	}
	void setVisible(bool visible) override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 2110182719375245575LL, visible, &__h); if (__h) return; }
		QDateEdit::setVisible(visible);
	}
	int heightForWidth(int arg1) const override {
		int __h = 0;
		if (__goOn[32] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 4311534060495737339LL, arg1, &__h); if (__h) return __r; }
		return QDateEdit::heightForWidth(arg1);
	}
	bool hasHeightForWidth() const override {
		int __h = 0;
		if (__goOn[33] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8266762397995423915LL, &__h); if (__h) return __r; }
		return QDateEdit::hasHeightForWidth();
	}
	QPaintEngine* paintEngine() const override {
		int __h = 0;
		if (__goOn[34] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 6398007404756346381LL, &__h); if (__h) return (QPaintEngine*)__r; }
		return QDateEdit::paintEngine();
	}
	void mouseDoubleClickEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[35] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4029073905697023887LL, (void*)event, &__h); if (__h) return; }
		QDateEdit::mouseDoubleClickEvent(event);
	}
	void enterEvent(QEnterEvent* event) override {
		int __h = 0;
		if (__goOn[36] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6316270346668237895LL, (void*)event, &__h); if (__h) return; }
		QDateEdit::enterEvent(event);
	}
	void leaveEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[37] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -551070009623711640LL, (void*)event, &__h); if (__h) return; }
		QDateEdit::leaveEvent(event);
	}
	void moveEvent(QMoveEvent* event) override {
		int __h = 0;
		if (__goOn[38] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -9161949150489525396LL, (void*)event, &__h); if (__h) return; }
		QDateEdit::moveEvent(event);
	}
	void tabletEvent(QTabletEvent* event) override {
		int __h = 0;
		if (__goOn[39] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -212068866860701141LL, (void*)event, &__h); if (__h) return; }
		QDateEdit::tabletEvent(event);
	}
	void actionEvent(QActionEvent* event) override {
		int __h = 0;
		if (__goOn[40] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1697444140088555581LL, (void*)event, &__h); if (__h) return; }
		QDateEdit::actionEvent(event);
	}
	void dragEnterEvent(QDragEnterEvent* event) override {
		int __h = 0;
		if (__goOn[41] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 206122715919917589LL, (void*)event, &__h); if (__h) return; }
		QDateEdit::dragEnterEvent(event);
	}
	void dragMoveEvent(QDragMoveEvent* event) override {
		int __h = 0;
		if (__goOn[42] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1781545137352745978LL, (void*)event, &__h); if (__h) return; }
		QDateEdit::dragMoveEvent(event);
	}
	void dragLeaveEvent(QDragLeaveEvent* event) override {
		int __h = 0;
		if (__goOn[43] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -3139904158373725006LL, (void*)event, &__h); if (__h) return; }
		QDateEdit::dragLeaveEvent(event);
	}
	void dropEvent(QDropEvent* event) override {
		int __h = 0;
		if (__goOn[44] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 9107058682754591716LL, (void*)event, &__h); if (__h) return; }
		QDateEdit::dropEvent(event);
	}
	bool nativeEvent(const QByteArray& eventType, void* message, qintptr* result) override {
		int __h = 0;
		if (__goOn[45] && cppgo_virt_int_voidp_voidp_longlongp) { int __r = ((int(*)(long long, long long, void*, void*, long long*, int*))cppgo_virt_int_voidp_voidp_longlongp)(__goId, -2039166586246443706LL, (void*)&eventType, (void*)message, (long long*)result, &__h); if (__h) return __r; }
		return QDateEdit::nativeEvent(eventType, message, result);
	}
	int metric(QPaintDevice::PaintDeviceMetric arg1) const override {
		int __h = 0;
		if (__goOn[46] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 427629649267311777LL, (int)(arg1), &__h); if (__h) return __r; }
		return QDateEdit::metric(arg1);
	}
	void initPainter(QPainter* painter) const override {
		int __h = 0;
		if (__goOn[47] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6477500296980050210LL, (void*)painter, &__h); if (__h) return; }
		QDateEdit::initPainter(painter);
	}
	QPaintDevice* redirected(QPoint* offset) const override {
		int __h = 0;
		if (__goOn[48] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 4127240633310655048LL, (void*)offset, &__h); if (__h) return (QPaintDevice*)__r; }
		return QDateEdit::redirected(offset);
	}
	QPainter* sharedPainter() const override {
		int __h = 0;
		if (__goOn[49] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 5792771328741774991LL, &__h); if (__h) return (QPainter*)__r; }
		return QDateEdit::sharedPainter();
	}
	void inputMethodEvent(QInputMethodEvent* arg1) override {
		int __h = 0;
		if (__goOn[50] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5438819661942843222LL, (void*)arg1, &__h); if (__h) return; }
		QDateEdit::inputMethodEvent(arg1);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[51] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QDateEdit::eventFilter(watched, event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[52] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QDateEdit::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[53] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QDateEdit::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[54] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QDateEdit::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[55] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QDateEdit::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QDateEdit;
static std::mutex __connsMu_QDateEdit;

extern "C" {

void* QDateEdit_base(void* self) { return (void*)static_cast<QDateTimeEdit*>((QDateEdit*)self); }
void QDateEdit_chain(void* self, void** out) {
	QDateEdit* __p = (QDateEdit*)self;
	QDateTimeEdit* __b0 = static_cast<QDateTimeEdit*>(__p); out[0] = (void*)__b0;
	QAbstractSpinBox* __b1 = static_cast<QAbstractSpinBox*>(__b0); out[1] = (void*)__b1;
	QWidget* __b2 = static_cast<QWidget*>(__b1); out[2] = (void*)__b2;
	QObject* __b3 = static_cast<QObject*>(__b2); out[3] = (void*)__b3;
}
void QDateEdit__set_override(void* self, int idx) { ((QDateEdit__Go*)self)->__goOn[idx] = 1; }
void* QDateEdit_fromQObject(void* o) { return (void*)qobject_cast<QDateEdit*>((QObject*)o); }
void* QDateEdit_new(long long __goid, void* parent) {
	return new QDateEdit__Go(__goid, (QWidget*)parent);
}
void* QDateEdit_new2(long long __goid) {
	return new QDateEdit__Go(__goid);
}
void* QDateEdit_new3(long long __goid, void* date) {
	return new QDateEdit__Go(__goid, *(QDate*)date);
}
void* QDateEdit_new4(long long __goid, void* date, void* parent) {
	return new QDateEdit__Go(__goid, *(QDate*)date, (QWidget*)parent);
}
char* QDateEdit_tr(const char* s, long long* __out_len) {
	auto&& __r = QDateEdit::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QDateEdit_userDateChanged(void* self, void* date) {
	((QDateEdit*)self)->userDateChanged(*(QDate*)date);
}
void* QDateEdit_lineEdit(void* self) {
	auto&& __r = ((QDateEdit__Go*)self)->lineEdit();
	return (void*)(__r);
}
void QDateEdit_setLineEdit(void* self, void* edit) {
	((QDateEdit__Go*)self)->setLineEdit((QLineEdit*)edit);
}
void QDateEdit_setTabOrder(void* arg1, void* arg2) {
	QDateEdit::setTabOrder((QWidget*)arg1, (QWidget*)arg2);
}
void* QDateEdit_mouseGrabber() {
	auto&& __r = QDateEdit::mouseGrabber();
	return (void*)(__r);
}
void* QDateEdit_keyboardGrabber() {
	auto&& __r = QDateEdit::keyboardGrabber();
	return (void*)(__r);
}
void* QDateEdit_find(unsigned long long arg1) {
	auto&& __r = QDateEdit::find(arg1);
	return (void*)(__r);
}
void* QDateEdit_createWindowContainer(void* window) {
	auto&& __r = QDateEdit::createWindowContainer((QWindow*)window);
	return (void*)(__r);
}
void QDateEdit_updateMicroFocus(void* self) {
	((QDateEdit__Go*)self)->updateMicroFocus();
}
void QDateEdit_create(void* self) {
	((QDateEdit__Go*)self)->create();
}
void QDateEdit_destroy(void* self) {
	((QDateEdit__Go*)self)->destroy();
}
int QDateEdit_focusNextChild(void* self) {
	auto&& __r = ((QDateEdit__Go*)self)->focusNextChild();
	return __r;
}
int QDateEdit_focusPreviousChild(void* self) {
	auto&& __r = ((QDateEdit__Go*)self)->focusPreviousChild();
	return __r;
}
void* QDateEdit_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QDateEdit::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QDateEdit_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QDateEdit::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QDateEdit_sender(void* self) {
	auto&& __r = ((QDateEdit__Go*)self)->sender();
	return (void*)(__r);
}
int QDateEdit_senderSignalIndex(void* self) {
	auto&& __r = ((QDateEdit__Go*)self)->senderSignalIndex();
	return __r;
}
int QDateEdit_receivers(void* self, const char* signal) {
	auto&& __r = ((QDateEdit__Go*)self)->receivers(signal);
	return __r;
}
int QDateEdit_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QDateEdit__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
double QDateEdit_devicePixelRatioFScale() {
	auto&& __r = QDateEdit::devicePixelRatioFScale();
	return __r;
}
int QDateEdit_encodeMetricF(int metric, double value) {
	auto&& __r = QDateEdit::encodeMetricF((QPaintDevice::PaintDeviceMetric)(metric), value);
	return __r;
}
double QDateEdit_getDecodedMetricF(void* self, int metricA, int metricB) {
	auto&& __r = ((QDateEdit__Go*)self)->getDecodedMetricF((QPaintDevice::PaintDeviceMetric)(metricA), (QPaintDevice::PaintDeviceMetric)(metricB));
	return __r;
}
char* QDateEdit_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QDateEdit::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QDateEdit_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QDateEdit::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QDateEdit_createWindowContainer2(void* window, void* parent) {
	auto&& __r = QDateEdit::createWindowContainer((QWindow*)window, (QWidget*)parent);
	return (void*)(__r);
}
void* QDateEdit_createWindowContainer3(void* window, void* parent, int flags) {
	auto&& __r = QDateEdit::createWindowContainer((QWindow*)window, (QWidget*)parent, QFlags<Qt::WindowType>(flags));
	return (void*)(__r);
}
void QDateEdit_updateMicroFocus2(void* self, int query) {
	((QDateEdit__Go*)self)->updateMicroFocus((Qt::InputMethodQuery)(query));
}
void QDateEdit_create2(void* self, unsigned long long arg1) {
	((QDateEdit__Go*)self)->create(arg1);
}
void QDateEdit_create3(void* self, unsigned long long arg1, int initializeWindow) {
	((QDateEdit__Go*)self)->create(arg1, initializeWindow);
}
void QDateEdit_create4(void* self, unsigned long long arg1, int initializeWindow, int destroyOldWindow) {
	((QDateEdit__Go*)self)->create(arg1, initializeWindow, destroyOldWindow);
}
void QDateEdit_destroy2(void* self, int destroyWindow) {
	((QDateEdit__Go*)self)->destroy(destroyWindow);
}
void QDateEdit_destroy3(void* self, int destroyWindow, int destroySubWindows) {
	((QDateEdit__Go*)self)->destroy(destroyWindow, destroySubWindows);
}
void* QDateEdit_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QDateEdit::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QDateEdit_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QDateEdit::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QDateEdit_metaObject_default(void* self) {
	auto&& __r = ((QDateEdit*)self)->QDateEdit::metaObject();
	return (void*)(__r);
}
void* QDateEdit_sizeHint_default(void* self) {
	auto&& __r = ((QDateEdit*)self)->QDateEdit::sizeHint();
	return new QSize(__r);
}
void QDateEdit_clear_default(void* self) {
	((QDateEdit*)self)->QDateEdit::clear();
}
void QDateEdit_stepBy_default(void* self, int steps) {
	((QDateEdit*)self)->QDateEdit::stepBy(steps);
}
int QDateEdit_event_default(void* self, void* event) {
	auto&& __r = ((QDateEdit*)self)->QDateEdit::event((QEvent*)event);
	return __r;
}
void QDateEdit_keyPressEvent_default(void* self, void* event) {
	((QDateEdit__Go*)self)->cppgoBase_keyPressEvent((QKeyEvent*)event);
}
void QDateEdit_wheelEvent_default(void* self, void* event) {
	((QDateEdit__Go*)self)->cppgoBase_wheelEvent((QWheelEvent*)event);
}
void QDateEdit_focusInEvent_default(void* self, void* event) {
	((QDateEdit__Go*)self)->cppgoBase_focusInEvent((QFocusEvent*)event);
}
int QDateEdit_focusNextPrevChild_default(void* self, int next) {
	auto&& __r = ((QDateEdit__Go*)self)->cppgoBase_focusNextPrevChild(next);
	return __r;
}
int QDateEdit_validate_default(void* self, void* input, int* pos) {
	int pos_ref = *pos;
	auto&& __r = ((QDateEdit__Go*)self)->cppgoBase_validate(*(QString*)input, pos_ref);
	*pos = pos_ref;
	return (int)(__r);
}
void QDateEdit_fixup_default(void* self, void* input) {
	((QDateEdit__Go*)self)->cppgoBase_fixup(*(QString*)input);
}
void* QDateEdit_dateTimeFromText_default(void* self, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = ((QDateEdit__Go*)self)->cppgoBase_dateTimeFromText(text_qstr);
	return new QDateTime(__r);
}
char* QDateEdit_textFromDateTime_default(void* self, void* dt, long long* __out_len) {
	auto&& __r = ((QDateEdit__Go*)self)->cppgoBase_textFromDateTime(*(QDateTime*)dt);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QDateEdit_stepEnabled_default(void* self) {
	auto&& __r = ((QDateEdit__Go*)self)->cppgoBase_stepEnabled();
	return (int)(__r);
}
void QDateEdit_mousePressEvent_default(void* self, void* event) {
	((QDateEdit__Go*)self)->cppgoBase_mousePressEvent((QMouseEvent*)event);
}
void QDateEdit_paintEvent_default(void* self, void* event) {
	((QDateEdit__Go*)self)->cppgoBase_paintEvent((QPaintEvent*)event);
}
void QDateEdit_initStyleOption_default(void* self, void* option) {
	((QDateEdit__Go*)self)->cppgoBase_initStyleOption((QStyleOptionSpinBox*)option);
}
void* QDateEdit_minimumSizeHint_default(void* self) {
	auto&& __r = ((QDateEdit*)self)->QDateEdit::minimumSizeHint();
	return new QSize(__r);
}
void* QDateEdit_inputMethodQuery_default(void* self, int arg1) {
	auto&& __r = ((QDateEdit*)self)->QDateEdit::inputMethodQuery((Qt::InputMethodQuery)(arg1));
	return new QVariant(__r);
}
void QDateEdit_resizeEvent_default(void* self, void* event) {
	((QDateEdit__Go*)self)->cppgoBase_resizeEvent((QResizeEvent*)event);
}
void QDateEdit_keyReleaseEvent_default(void* self, void* event) {
	((QDateEdit__Go*)self)->cppgoBase_keyReleaseEvent((QKeyEvent*)event);
}
void QDateEdit_focusOutEvent_default(void* self, void* event) {
	((QDateEdit__Go*)self)->cppgoBase_focusOutEvent((QFocusEvent*)event);
}
void QDateEdit_contextMenuEvent_default(void* self, void* event) {
	((QDateEdit__Go*)self)->cppgoBase_contextMenuEvent((QContextMenuEvent*)event);
}
void QDateEdit_changeEvent_default(void* self, void* event) {
	((QDateEdit__Go*)self)->cppgoBase_changeEvent((QEvent*)event);
}
void QDateEdit_closeEvent_default(void* self, void* event) {
	((QDateEdit__Go*)self)->cppgoBase_closeEvent((QCloseEvent*)event);
}
void QDateEdit_hideEvent_default(void* self, void* event) {
	((QDateEdit__Go*)self)->cppgoBase_hideEvent((QHideEvent*)event);
}
void QDateEdit_mouseReleaseEvent_default(void* self, void* event) {
	((QDateEdit__Go*)self)->cppgoBase_mouseReleaseEvent((QMouseEvent*)event);
}
void QDateEdit_mouseMoveEvent_default(void* self, void* event) {
	((QDateEdit__Go*)self)->cppgoBase_mouseMoveEvent((QMouseEvent*)event);
}
void QDateEdit_timerEvent_default(void* self, void* event) {
	((QDateEdit__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QDateEdit_showEvent_default(void* self, void* event) {
	((QDateEdit__Go*)self)->cppgoBase_showEvent((QShowEvent*)event);
}
int QDateEdit_devType_default(void* self) {
	auto&& __r = ((QDateEdit*)self)->QDateEdit::devType();
	return __r;
}
void QDateEdit_setVisible_default(void* self, int visible) {
	((QDateEdit*)self)->QDateEdit::setVisible(visible);
}
int QDateEdit_heightForWidth_default(void* self, int arg1) {
	auto&& __r = ((QDateEdit*)self)->QDateEdit::heightForWidth(arg1);
	return __r;
}
int QDateEdit_hasHeightForWidth_default(void* self) {
	auto&& __r = ((QDateEdit*)self)->QDateEdit::hasHeightForWidth();
	return __r;
}
void* QDateEdit_paintEngine_default(void* self) {
	auto&& __r = ((QDateEdit*)self)->QDateEdit::paintEngine();
	return (void*)(__r);
}
void QDateEdit_mouseDoubleClickEvent_default(void* self, void* event) {
	((QDateEdit__Go*)self)->cppgoBase_mouseDoubleClickEvent((QMouseEvent*)event);
}
void QDateEdit_enterEvent_default(void* self, void* event) {
	((QDateEdit__Go*)self)->cppgoBase_enterEvent((QEnterEvent*)event);
}
void QDateEdit_leaveEvent_default(void* self, void* event) {
	((QDateEdit__Go*)self)->cppgoBase_leaveEvent((QEvent*)event);
}
void QDateEdit_moveEvent_default(void* self, void* event) {
	((QDateEdit__Go*)self)->cppgoBase_moveEvent((QMoveEvent*)event);
}
void QDateEdit_tabletEvent_default(void* self, void* event) {
	((QDateEdit__Go*)self)->cppgoBase_tabletEvent((QTabletEvent*)event);
}
void QDateEdit_actionEvent_default(void* self, void* event) {
	((QDateEdit__Go*)self)->cppgoBase_actionEvent((QActionEvent*)event);
}
void QDateEdit_dragEnterEvent_default(void* self, void* event) {
	((QDateEdit__Go*)self)->cppgoBase_dragEnterEvent((QDragEnterEvent*)event);
}
void QDateEdit_dragMoveEvent_default(void* self, void* event) {
	((QDateEdit__Go*)self)->cppgoBase_dragMoveEvent((QDragMoveEvent*)event);
}
void QDateEdit_dragLeaveEvent_default(void* self, void* event) {
	((QDateEdit__Go*)self)->cppgoBase_dragLeaveEvent((QDragLeaveEvent*)event);
}
void QDateEdit_dropEvent_default(void* self, void* event) {
	((QDateEdit__Go*)self)->cppgoBase_dropEvent((QDropEvent*)event);
}
int QDateEdit_metric_default(void* self, int arg1) {
	auto&& __r = ((QDateEdit__Go*)self)->cppgoBase_metric(arg1);
	return __r;
}
void QDateEdit_initPainter_default(void* self, void* painter) {
	((QDateEdit__Go*)self)->cppgoBase_initPainter((QPainter*)painter);
}
void* QDateEdit_redirected_default(void* self, void* offset) {
	auto&& __r = ((QDateEdit__Go*)self)->cppgoBase_redirected((QPoint*)offset);
	return (void*)(__r);
}
void* QDateEdit_sharedPainter_default(void* self) {
	auto&& __r = ((QDateEdit__Go*)self)->cppgoBase_sharedPainter();
	return (void*)(__r);
}
void QDateEdit_inputMethodEvent_default(void* self, void* arg1) {
	((QDateEdit__Go*)self)->cppgoBase_inputMethodEvent((QInputMethodEvent*)arg1);
}
int QDateEdit_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QDateEdit*)self)->QDateEdit::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QDateEdit_childEvent_default(void* self, void* event) {
	((QDateEdit__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QDateEdit_customEvent_default(void* self, void* event) {
	((QDateEdit__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QDateEdit_connectNotify_default(void* self, void* signal) {
	((QDateEdit__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QDateEdit_disconnectNotify_default(void* self, void* signal) {
	((QDateEdit__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QDateEdit_connect_userDateChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDateEdit*)self, &QDateEdit::userDateChanged, (QDateEdit*)self, [__id](QDate date){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&date);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDateEdit);
	__conns_QDateEdit[__id] = __c;
}
void QDateEdit_connect_dateTimeChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDateEdit*)self, &QDateEdit::dateTimeChanged, (QDateEdit*)self, [__id](const QDateTime& dateTime){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&dateTime);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDateEdit);
	__conns_QDateEdit[__id] = __c;
}
void QDateEdit_connect_timeChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDateEdit*)self, &QDateEdit::timeChanged, (QDateEdit*)self, [__id](QTime time){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&time);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDateEdit);
	__conns_QDateEdit[__id] = __c;
}
void QDateEdit_connect_dateChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDateEdit*)self, &QDateEdit::dateChanged, (QDateEdit*)self, [__id](QDate date){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&date);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDateEdit);
	__conns_QDateEdit[__id] = __c;
}
void QDateEdit_connect_editingFinished(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDateEdit*)self, &QDateEdit::editingFinished, (QDateEdit*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDateEdit);
	__conns_QDateEdit[__id] = __c;
}
void QDateEdit_connect_returnPressed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDateEdit*)self, &QDateEdit::returnPressed, (QDateEdit*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDateEdit);
	__conns_QDateEdit[__id] = __c;
}
void QDateEdit_connect_windowTitleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDateEdit*)self, &QDateEdit::windowTitleChanged, (QDateEdit*)self, [__id](const QString& title){
		QByteArray title_ba = title.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)title_ba.constData(), title_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDateEdit);
	__conns_QDateEdit[__id] = __c;
}
void QDateEdit_connect_windowIconChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDateEdit*)self, &QDateEdit::windowIconChanged, (QDateEdit*)self, [__id](const QIcon& icon){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&icon);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDateEdit);
	__conns_QDateEdit[__id] = __c;
}
void QDateEdit_connect_windowIconTextChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDateEdit*)self, &QDateEdit::windowIconTextChanged, (QDateEdit*)self, [__id](const QString& iconText){
		QByteArray iconText_ba = iconText.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)iconText_ba.constData(), iconText_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDateEdit);
	__conns_QDateEdit[__id] = __c;
}
void QDateEdit_connect_customContextMenuRequested(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDateEdit*)self, &QDateEdit::customContextMenuRequested, (QDateEdit*)self, [__id](const QPoint& pos){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&pos);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDateEdit);
	__conns_QDateEdit[__id] = __c;
}
void QDateEdit_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDateEdit*)self, &QDateEdit::destroyed, (QDateEdit*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDateEdit);
	__conns_QDateEdit[__id] = __c;
}
void QDateEdit_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDateEdit*)self, &QDateEdit::destroyed, (QDateEdit*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDateEdit);
	__conns_QDateEdit[__id] = __c;
}
void QDateEdit_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDateEdit*)self, &QDateEdit::objectNameChanged, (QDateEdit*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDateEdit);
	__conns_QDateEdit[__id] = __c;
}
void QDateEdit__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QDateEdit);
		auto __it = __conns_QDateEdit.find(id);
		if (__it == __conns_QDateEdit.end()) return;
		__c = __it->second;
		__conns_QDateEdit.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QDateEdit__Go_new(long long id) { return new QDateEdit__Go(id); }
void QDateEdit_delete(void* self) { delete (QDateEdit*)self; }

}

#include <QtWidgets/qgraphicsitem.h>
#include <QtGui/qpainterpath.h>
#include <QtCore/qrect.h>
#include <QtCore/qpoint.h>
#include <QtGui/qpainter.h>
#include <QtWidgets/qstyleoption.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qvariant.h>
#include <QtGui/qpen.h>
#include <QtGui/qbrush.h>
#include <QtWidgets/qgraphicsscene.h>
#include <QtWidgets/qgraphicswidget.h>
#include <QtCore/qsize.h>
#include <QtCore/qstring.h>
#include <QtGui/qcursor.h>
#include <QtWidgets/qgraphicseffect.h>
#include <QtCore/qnamespace.h>
#include <QtGui/qtransform.h>
#include <QtWidgets/qgraphicstransform.h>
#include <QtGui/qregion.h>
#include <QtGui/qpolygon.h>
#include <QtCore/qcoreevent.h>
#include <QtWidgets/qgraphicssceneevent.h>
#include <QtGui/qevent.h>
#include "gen_qgraphicspathitem.h"
#include "gen_purego_cb.h"

class QGraphicsPathItem__Go : public QGraphicsPathItem {
public:
	long long __goId;
	unsigned char __goOn[35] = {};
	template<typename... A> QGraphicsPathItem__Go(long long id, A&&... a) : QGraphicsPathItem(std::forward<A>(a)...), __goId(id) {}
	using QGraphicsPathItem::updateMicroFocus;
	using QGraphicsPathItem::addToIndex;
	using QGraphicsPathItem::removeFromIndex;
	using QGraphicsPathItem::prepareGeometryChange;
	bool cppgoBase_supportsExtension(int __a0) { return QGraphicsPathItem::supportsExtension((QGraphicsItem::Extension)__a0); }
	void cppgoBase_setExtension(int __a0, const QVariant& __a1) { QGraphicsPathItem::setExtension((QGraphicsItem::Extension)__a0, __a1); }
	QVariant cppgoBase_extension(const QVariant& __a0) { return QGraphicsPathItem::extension(__a0); }
	bool cppgoBase_sceneEventFilter(QGraphicsItem* __a0, QEvent* __a1) { return QGraphicsPathItem::sceneEventFilter(__a0, __a1); }
	bool cppgoBase_sceneEvent(QEvent* __a0) { return QGraphicsPathItem::sceneEvent(__a0); }
	void cppgoBase_contextMenuEvent(QGraphicsSceneContextMenuEvent* __a0) { QGraphicsPathItem::contextMenuEvent(__a0); }
	void cppgoBase_dragEnterEvent(QGraphicsSceneDragDropEvent* __a0) { QGraphicsPathItem::dragEnterEvent(__a0); }
	void cppgoBase_dragLeaveEvent(QGraphicsSceneDragDropEvent* __a0) { QGraphicsPathItem::dragLeaveEvent(__a0); }
	void cppgoBase_dragMoveEvent(QGraphicsSceneDragDropEvent* __a0) { QGraphicsPathItem::dragMoveEvent(__a0); }
	void cppgoBase_dropEvent(QGraphicsSceneDragDropEvent* __a0) { QGraphicsPathItem::dropEvent(__a0); }
	void cppgoBase_focusInEvent(QFocusEvent* __a0) { QGraphicsPathItem::focusInEvent(__a0); }
	void cppgoBase_focusOutEvent(QFocusEvent* __a0) { QGraphicsPathItem::focusOutEvent(__a0); }
	void cppgoBase_hoverEnterEvent(QGraphicsSceneHoverEvent* __a0) { QGraphicsPathItem::hoverEnterEvent(__a0); }
	void cppgoBase_hoverMoveEvent(QGraphicsSceneHoverEvent* __a0) { QGraphicsPathItem::hoverMoveEvent(__a0); }
	void cppgoBase_hoverLeaveEvent(QGraphicsSceneHoverEvent* __a0) { QGraphicsPathItem::hoverLeaveEvent(__a0); }
	void cppgoBase_keyPressEvent(QKeyEvent* __a0) { QGraphicsPathItem::keyPressEvent(__a0); }
	void cppgoBase_keyReleaseEvent(QKeyEvent* __a0) { QGraphicsPathItem::keyReleaseEvent(__a0); }
	void cppgoBase_mousePressEvent(QGraphicsSceneMouseEvent* __a0) { QGraphicsPathItem::mousePressEvent(__a0); }
	void cppgoBase_mouseMoveEvent(QGraphicsSceneMouseEvent* __a0) { QGraphicsPathItem::mouseMoveEvent(__a0); }
	void cppgoBase_mouseReleaseEvent(QGraphicsSceneMouseEvent* __a0) { QGraphicsPathItem::mouseReleaseEvent(__a0); }
	void cppgoBase_mouseDoubleClickEvent(QGraphicsSceneMouseEvent* __a0) { QGraphicsPathItem::mouseDoubleClickEvent(__a0); }
	void cppgoBase_wheelEvent(QGraphicsSceneWheelEvent* __a0) { QGraphicsPathItem::wheelEvent(__a0); }
	void cppgoBase_inputMethodEvent(QInputMethodEvent* __a0) { QGraphicsPathItem::inputMethodEvent(__a0); }
	QVariant cppgoBase_inputMethodQuery(int __a0) { return QGraphicsPathItem::inputMethodQuery((Qt::InputMethodQuery)__a0); }
	QVariant cppgoBase_itemChange(int __a0, const QVariant& __a1) { return QGraphicsPathItem::itemChange((QGraphicsItem::GraphicsItemChange)__a0, __a1); }
	QRectF boundingRect() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 2510008782785038831LL, &__h); if (__h) return *(QRectF*)__r; }
		return QGraphicsPathItem::boundingRect();
	}
	QPainterPath shape() const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 7938019069247444294LL, &__h); if (__h) return *(QPainterPath*)__r; }
		return QGraphicsPathItem::shape();
	}
	bool contains(const QPointF& point) const override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 5645366511944179768LL, (void*)&point, &__h); if (__h) return __r; }
		return QGraphicsPathItem::contains(point);
	}
	void paint(QPainter* painter, const QStyleOptionGraphicsItem* option, QWidget* widget) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_voidp_voidp_voidp) { ((void(*)(long long, long long, void*, void*, void*, int*))cppgo_virt_void_voidp_voidp_voidp)(__goId, 6083584261053556819LL, (void*)painter, (void*)option, (void*)widget, &__h); if (__h) return; }
		QGraphicsPathItem::paint(painter, option, widget);
	}
	bool isObscuredBy(const QGraphicsItem* item) const override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, -7215030271914903671LL, (void*)item, &__h); if (__h) return __r; }
		return QGraphicsPathItem::isObscuredBy(item);
	}
	QPainterPath opaqueArea() const override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 399370970023989435LL, &__h); if (__h) return *(QPainterPath*)__r; }
		return QGraphicsPathItem::opaqueArea();
	}
	int type() const override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -6371403872082420691LL, &__h); if (__h) return __r; }
		return QGraphicsPathItem::type();
	}
	bool supportsExtension(QGraphicsItem::Extension extension) const override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, -5729166443651306418LL, (int)(extension), &__h); if (__h) return __r; }
		return QGraphicsPathItem::supportsExtension(extension);
	}
	void setExtension(QGraphicsItem::Extension extension, const QVariant& variant) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void_int_voidp) { ((void(*)(long long, long long, int, void*, int*))cppgo_virt_void_int_voidp)(__goId, -6730181576106171580LL, (int)(extension), (void*)&variant, &__h); if (__h) return; }
		QGraphicsPathItem::setExtension(extension, variant);
	}
	QVariant extension(const QVariant& variant) const override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 2270616358835251226LL, (void*)&variant, &__h); if (__h) return *(QVariant*)__r; }
		return QGraphicsPathItem::extension(variant);
	}
	void advance(int phase) override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, -4830836341480073475LL, phase, &__h); if (__h) return; }
		QGraphicsPathItem::advance(phase);
	}
	bool collidesWithItem(const QGraphicsItem* other, Qt::ItemSelectionMode mode) const override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_int_voidp_int) { int __r = ((int(*)(long long, long long, void*, int, int*))cppgo_virt_int_voidp_int)(__goId, 6469736184362593173LL, (void*)other, (int)(mode), &__h); if (__h) return __r; }
		return QGraphicsPathItem::collidesWithItem(other, mode);
	}
	bool collidesWithPath(const QPainterPath& path, Qt::ItemSelectionMode mode) const override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_int_voidp_int) { int __r = ((int(*)(long long, long long, void*, int, int*))cppgo_virt_int_voidp_int)(__goId, 6636575182718753853LL, (void*)&path, (int)(mode), &__h); if (__h) return __r; }
		return QGraphicsPathItem::collidesWithPath(path, mode);
	}
	bool sceneEventFilter(QGraphicsItem* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, -4584203218241215725LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QGraphicsPathItem::sceneEventFilter(watched, event);
	}
	bool sceneEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, -2846371507838587071LL, (void*)event, &__h); if (__h) return __r; }
		return QGraphicsPathItem::sceneEvent(event);
	}
	void contextMenuEvent(QGraphicsSceneContextMenuEvent* event) override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6698043222227006015LL, (void*)event, &__h); if (__h) return; }
		QGraphicsPathItem::contextMenuEvent(event);
	}
	void dragEnterEvent(QGraphicsSceneDragDropEvent* event) override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 206122715919917589LL, (void*)event, &__h); if (__h) return; }
		QGraphicsPathItem::dragEnterEvent(event);
	}
	void dragLeaveEvent(QGraphicsSceneDragDropEvent* event) override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -3139904158373725006LL, (void*)event, &__h); if (__h) return; }
		QGraphicsPathItem::dragLeaveEvent(event);
	}
	void dragMoveEvent(QGraphicsSceneDragDropEvent* event) override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1781545137352745978LL, (void*)event, &__h); if (__h) return; }
		QGraphicsPathItem::dragMoveEvent(event);
	}
	void dropEvent(QGraphicsSceneDragDropEvent* event) override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 9107058682754591716LL, (void*)event, &__h); if (__h) return; }
		QGraphicsPathItem::dropEvent(event);
	}
	void focusInEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 975043767033552552LL, (void*)event, &__h); if (__h) return; }
		QGraphicsPathItem::focusInEvent(event);
	}
	void focusOutEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -385596010576048993LL, (void*)event, &__h); if (__h) return; }
		QGraphicsPathItem::focusOutEvent(event);
	}
	void hoverEnterEvent(QGraphicsSceneHoverEvent* event) override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6932941943906681679LL, (void*)event, &__h); if (__h) return; }
		QGraphicsPathItem::hoverEnterEvent(event);
	}
	void hoverMoveEvent(QGraphicsSceneHoverEvent* event) override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -2745191522144316396LL, (void*)event, &__h); if (__h) return; }
		QGraphicsPathItem::hoverMoveEvent(event);
	}
	void hoverLeaveEvent(QGraphicsSceneHoverEvent* event) override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7944408482203603984LL, (void*)event, &__h); if (__h) return; }
		QGraphicsPathItem::hoverLeaveEvent(event);
	}
	void keyPressEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1023791977679793185LL, (void*)event, &__h); if (__h) return; }
		QGraphicsPathItem::keyPressEvent(event);
	}
	void keyReleaseEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 772598475950050299LL, (void*)event, &__h); if (__h) return; }
		QGraphicsPathItem::keyReleaseEvent(event);
	}
	void mousePressEvent(QGraphicsSceneMouseEvent* event) override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -332090898078543319LL, (void*)event, &__h); if (__h) return; }
		QGraphicsPathItem::mousePressEvent(event);
	}
	void mouseMoveEvent(QGraphicsSceneMouseEvent* event) override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4255494004099806581LL, (void*)event, &__h); if (__h) return; }
		QGraphicsPathItem::mouseMoveEvent(event);
	}
	void mouseReleaseEvent(QGraphicsSceneMouseEvent* event) override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4415920049297795213LL, (void*)event, &__h); if (__h) return; }
		QGraphicsPathItem::mouseReleaseEvent(event);
	}
	void mouseDoubleClickEvent(QGraphicsSceneMouseEvent* event) override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4029073905697023887LL, (void*)event, &__h); if (__h) return; }
		QGraphicsPathItem::mouseDoubleClickEvent(event);
	}
	void wheelEvent(QGraphicsSceneWheelEvent* event) override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1919988340934511316LL, (void*)event, &__h); if (__h) return; }
		QGraphicsPathItem::wheelEvent(event);
	}
	void inputMethodEvent(QInputMethodEvent* event) override {
		int __h = 0;
		if (__goOn[32] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5438819661942843222LL, (void*)event, &__h); if (__h) return; }
		QGraphicsPathItem::inputMethodEvent(event);
	}
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const override {
		int __h = 0;
		if (__goOn[33] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7345853097340709638LL, (int)(query), &__h); if (__h) return *(QVariant*)__r; }
		return QGraphicsPathItem::inputMethodQuery(query);
	}
	QVariant itemChange(QGraphicsItem::GraphicsItemChange change, const QVariant& value) override {
		int __h = 0;
		if (__goOn[34] && cppgo_virt_voidp_int_voidp) { void* __r = ((void*(*)(long long, long long, int, void*, int*))cppgo_virt_voidp_int_voidp)(__goId, -1888572168794878830LL, (int)(change), (void*)&value, &__h); if (__h) return *(QVariant*)__r; }
		return QGraphicsPathItem::itemChange(change, value);
	}
};

extern "C" {

void* QGraphicsPathItem_base(void* self) { return (void*)static_cast<QAbstractGraphicsShapeItem*>((QGraphicsPathItem*)self); }
void QGraphicsPathItem_chain(void* self, void** out) {
	QGraphicsPathItem* __p = (QGraphicsPathItem*)self;
	QAbstractGraphicsShapeItem* __b0 = static_cast<QAbstractGraphicsShapeItem*>(__p); out[0] = (void*)__b0;
	QGraphicsItem* __b1 = static_cast<QGraphicsItem*>(__b0); out[1] = (void*)__b1;
}
void QGraphicsPathItem__set_override(void* self, int idx) { ((QGraphicsPathItem__Go*)self)->__goOn[idx] = 1; }
void* QGraphicsPathItem_new(long long __goid) {
	return new QGraphicsPathItem__Go(__goid);
}
void* QGraphicsPathItem_new2(long long __goid, void* path) {
	return new QGraphicsPathItem__Go(__goid, *(QPainterPath*)path);
}
void* QGraphicsPathItem_new3(long long __goid, void* parent) {
	return new QGraphicsPathItem__Go(__goid, (QGraphicsItem*)parent);
}
void* QGraphicsPathItem_new4(long long __goid, void* path, void* parent) {
	return new QGraphicsPathItem__Go(__goid, *(QPainterPath*)path, (QGraphicsItem*)parent);
}
void* QGraphicsPathItem_path(void* self) {
	auto&& __r = ((QGraphicsPathItem*)self)->path();
	return new QPainterPath(__r);
}
void QGraphicsPathItem_setPath(void* self, void* path) {
	((QGraphicsPathItem*)self)->setPath(*(QPainterPath*)path);
}
void QGraphicsPathItem_updateMicroFocus(void* self) {
	((QGraphicsPathItem__Go*)self)->updateMicroFocus();
}
void QGraphicsPathItem_addToIndex(void* self) {
	((QGraphicsPathItem__Go*)self)->addToIndex();
}
void QGraphicsPathItem_removeFromIndex(void* self) {
	((QGraphicsPathItem__Go*)self)->removeFromIndex();
}
void QGraphicsPathItem_prepareGeometryChange(void* self) {
	((QGraphicsPathItem__Go*)self)->prepareGeometryChange();
}
void* QGraphicsPathItem_boundingRect_default(void* self) {
	auto&& __r = ((QGraphicsPathItem*)self)->QGraphicsPathItem::boundingRect();
	return new QRectF(__r);
}
void* QGraphicsPathItem_shape_default(void* self) {
	auto&& __r = ((QGraphicsPathItem*)self)->QGraphicsPathItem::shape();
	return new QPainterPath(__r);
}
int QGraphicsPathItem_contains_default(void* self, void* point) {
	auto&& __r = ((QGraphicsPathItem*)self)->QGraphicsPathItem::contains(*(QPointF*)point);
	return __r;
}
void QGraphicsPathItem_paint_default(void* self, void* painter, void* option, void* widget) {
	((QGraphicsPathItem*)self)->QGraphicsPathItem::paint((QPainter*)painter, (QStyleOptionGraphicsItem*)option, (QWidget*)widget);
}
int QGraphicsPathItem_isObscuredBy_default(void* self, void* item) {
	auto&& __r = ((QGraphicsPathItem*)self)->QGraphicsPathItem::isObscuredBy((QGraphicsItem*)item);
	return __r;
}
void* QGraphicsPathItem_opaqueArea_default(void* self) {
	auto&& __r = ((QGraphicsPathItem*)self)->QGraphicsPathItem::opaqueArea();
	return new QPainterPath(__r);
}
int QGraphicsPathItem_type_default(void* self) {
	auto&& __r = ((QGraphicsPathItem*)self)->QGraphicsPathItem::type();
	return __r;
}
int QGraphicsPathItem_supportsExtension_default(void* self, int extension) {
	auto&& __r = ((QGraphicsPathItem__Go*)self)->cppgoBase_supportsExtension(extension);
	return __r;
}
void QGraphicsPathItem_setExtension_default(void* self, int extension, void* variant) {
	((QGraphicsPathItem__Go*)self)->cppgoBase_setExtension(extension, *(QVariant*)variant);
}
void* QGraphicsPathItem_extension_default(void* self, void* variant) {
	auto&& __r = ((QGraphicsPathItem__Go*)self)->cppgoBase_extension(*(QVariant*)variant);
	return new QVariant(__r);
}
void QGraphicsPathItem_advance_default(void* self, int phase) {
	((QGraphicsPathItem*)self)->QGraphicsPathItem::advance(phase);
}
int QGraphicsPathItem_collidesWithItem_default(void* self, void* other, int mode) {
	auto&& __r = ((QGraphicsPathItem*)self)->QGraphicsPathItem::collidesWithItem((QGraphicsItem*)other, (Qt::ItemSelectionMode)(mode));
	return __r;
}
int QGraphicsPathItem_collidesWithPath_default(void* self, void* path, int mode) {
	auto&& __r = ((QGraphicsPathItem*)self)->QGraphicsPathItem::collidesWithPath(*(QPainterPath*)path, (Qt::ItemSelectionMode)(mode));
	return __r;
}
int QGraphicsPathItem_sceneEventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QGraphicsPathItem__Go*)self)->cppgoBase_sceneEventFilter((QGraphicsItem*)watched, (QEvent*)event);
	return __r;
}
int QGraphicsPathItem_sceneEvent_default(void* self, void* event) {
	auto&& __r = ((QGraphicsPathItem__Go*)self)->cppgoBase_sceneEvent((QEvent*)event);
	return __r;
}
void QGraphicsPathItem_contextMenuEvent_default(void* self, void* event) {
	((QGraphicsPathItem__Go*)self)->cppgoBase_contextMenuEvent((QGraphicsSceneContextMenuEvent*)event);
}
void QGraphicsPathItem_dragEnterEvent_default(void* self, void* event) {
	((QGraphicsPathItem__Go*)self)->cppgoBase_dragEnterEvent((QGraphicsSceneDragDropEvent*)event);
}
void QGraphicsPathItem_dragLeaveEvent_default(void* self, void* event) {
	((QGraphicsPathItem__Go*)self)->cppgoBase_dragLeaveEvent((QGraphicsSceneDragDropEvent*)event);
}
void QGraphicsPathItem_dragMoveEvent_default(void* self, void* event) {
	((QGraphicsPathItem__Go*)self)->cppgoBase_dragMoveEvent((QGraphicsSceneDragDropEvent*)event);
}
void QGraphicsPathItem_dropEvent_default(void* self, void* event) {
	((QGraphicsPathItem__Go*)self)->cppgoBase_dropEvent((QGraphicsSceneDragDropEvent*)event);
}
void QGraphicsPathItem_focusInEvent_default(void* self, void* event) {
	((QGraphicsPathItem__Go*)self)->cppgoBase_focusInEvent((QFocusEvent*)event);
}
void QGraphicsPathItem_focusOutEvent_default(void* self, void* event) {
	((QGraphicsPathItem__Go*)self)->cppgoBase_focusOutEvent((QFocusEvent*)event);
}
void QGraphicsPathItem_hoverEnterEvent_default(void* self, void* event) {
	((QGraphicsPathItem__Go*)self)->cppgoBase_hoverEnterEvent((QGraphicsSceneHoverEvent*)event);
}
void QGraphicsPathItem_hoverMoveEvent_default(void* self, void* event) {
	((QGraphicsPathItem__Go*)self)->cppgoBase_hoverMoveEvent((QGraphicsSceneHoverEvent*)event);
}
void QGraphicsPathItem_hoverLeaveEvent_default(void* self, void* event) {
	((QGraphicsPathItem__Go*)self)->cppgoBase_hoverLeaveEvent((QGraphicsSceneHoverEvent*)event);
}
void QGraphicsPathItem_keyPressEvent_default(void* self, void* event) {
	((QGraphicsPathItem__Go*)self)->cppgoBase_keyPressEvent((QKeyEvent*)event);
}
void QGraphicsPathItem_keyReleaseEvent_default(void* self, void* event) {
	((QGraphicsPathItem__Go*)self)->cppgoBase_keyReleaseEvent((QKeyEvent*)event);
}
void QGraphicsPathItem_mousePressEvent_default(void* self, void* event) {
	((QGraphicsPathItem__Go*)self)->cppgoBase_mousePressEvent((QGraphicsSceneMouseEvent*)event);
}
void QGraphicsPathItem_mouseMoveEvent_default(void* self, void* event) {
	((QGraphicsPathItem__Go*)self)->cppgoBase_mouseMoveEvent((QGraphicsSceneMouseEvent*)event);
}
void QGraphicsPathItem_mouseReleaseEvent_default(void* self, void* event) {
	((QGraphicsPathItem__Go*)self)->cppgoBase_mouseReleaseEvent((QGraphicsSceneMouseEvent*)event);
}
void QGraphicsPathItem_mouseDoubleClickEvent_default(void* self, void* event) {
	((QGraphicsPathItem__Go*)self)->cppgoBase_mouseDoubleClickEvent((QGraphicsSceneMouseEvent*)event);
}
void QGraphicsPathItem_wheelEvent_default(void* self, void* event) {
	((QGraphicsPathItem__Go*)self)->cppgoBase_wheelEvent((QGraphicsSceneWheelEvent*)event);
}
void QGraphicsPathItem_inputMethodEvent_default(void* self, void* event) {
	((QGraphicsPathItem__Go*)self)->cppgoBase_inputMethodEvent((QInputMethodEvent*)event);
}
void* QGraphicsPathItem_inputMethodQuery_default(void* self, int query) {
	auto&& __r = ((QGraphicsPathItem__Go*)self)->cppgoBase_inputMethodQuery(query);
	return new QVariant(__r);
}
void* QGraphicsPathItem_itemChange_default(void* self, int change, void* value) {
	auto&& __r = ((QGraphicsPathItem__Go*)self)->cppgoBase_itemChange(change, *(QVariant*)value);
	return new QVariant(__r);
}
void* QGraphicsPathItem__Go_new(long long id) { return new QGraphicsPathItem__Go(id); }
void QGraphicsPathItem_delete(void* self) { delete (QGraphicsPathItem*)self; }

}

#include <QtWidgets/qgraphicssceneevent.h>
#include <QtCore/qcoreevent.h>
#include <QtWidgets/qwidget.h>
#include "gen_qgraphicssceneevent.h"

extern "C" {

void* QGraphicsSceneEvent_base(void* self) { return (void*)static_cast<QEvent*>((QGraphicsSceneEvent*)self); }
void* QGraphicsSceneEvent_new(int type_) {
	return new QGraphicsSceneEvent((QEvent::Type)(type_));
}
void* QGraphicsSceneEvent_widget(void* self) {
	auto&& __r = ((QGraphicsSceneEvent*)self)->widget();
	return (void*)(__r);
}
void QGraphicsSceneEvent_setWidget(void* self, void* widget) {
	((QGraphicsSceneEvent*)self)->setWidget((QWidget*)widget);
}
unsigned long long QGraphicsSceneEvent_timestamp(void* self) {
	auto&& __r = ((QGraphicsSceneEvent*)self)->timestamp();
	return __r;
}
void QGraphicsSceneEvent_setTimestamp(void* self, unsigned long long ts) {
	((QGraphicsSceneEvent*)self)->setTimestamp(ts);
}
int QGraphicsSceneEvent_registerEventType() {
	auto&& __r = QGraphicsSceneEvent::registerEventType();
	return __r;
}
int QGraphicsSceneEvent_registerEventType2(int hint) {
	auto&& __r = QGraphicsSceneEvent::registerEventType(hint);
	return __r;
}
void QGraphicsSceneEvent_delete(void* self) { delete (QGraphicsSceneEvent*)self; }

}

#include <QtWidgets/qlistview.h>
#include <QtWidgets/qabstractitemview.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qsize.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qrect.h>
#include <QtCore/qabstractitemmodel.h>
#include <QtCore/qpoint.h>
#include <QtCore/qcoreevent.h>
#include <QtGui/qevent.h>
#include <QtWidgets/qstyleoption.h>
#include <QtCore/qitemselectionmodel.h>
#include <QtGui/qregion.h>
#include <QtWidgets/qabstractitemdelegate.h>
#include <QtCore/qvariant.h>
#include <QtCore/qobject.h>
#include <QtWidgets/qscrollbar.h>
#include <QtWidgets/qabstractscrollarea.h>
#include <QtCore/qmargins.h>
#include <QtWidgets/qframe.h>
#include <QtGui/qpainter.h>
#include <QtWidgets/qstyle.h>
#include <QtGui/qpalette.h>
#include <QtGui/qfont.h>
#include <QtGui/qfontmetrics.h>
#include <QtGui/qfontinfo.h>
#include <QtGui/qcursor.h>
#include <QtGui/qbitmap.h>
#include <QtGui/qpaintdevice.h>
#include <QtGui/qpixmap.h>
#include <QtWidgets/qgraphicseffect.h>
#include <QtGui/qicon.h>
#include <QtCore/qlocale.h>
#include <QtGui/qkeysequence.h>
#include <QtWidgets/qgraphicsproxywidget.h>
#include <QtCore/qbytearray.h>
#include <QtWidgets/qsizepolicy.h>
#include <QtWidgets/qlayout.h>
#include <QtGui/qaction.h>
#include <QtGui/qpaintengine.h>
#include <QtGui/qbackingstore.h>
#include <QtGui/qwindow.h>
#include <QtGui/qscreen.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qlistview.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QListView__Go : public QListView {
public:
	long long __goId;
	unsigned char __goOn[91] = {};
	template<typename... A> QListView__Go(long long id, A&&... a) : QListView(std::forward<A>(a)...), __goId(id) {}
	using QListView::resizeContents;
	using QListView::contentsSize;
	using QListView::rectForIndex;
	using QListView::setPositionForIndex;
	using QListView::state;
	using QListView::scheduleDelayedItemsLayout;
	using QListView::executeDelayedItemsLayout;
	using QListView::setDirtyRegion;
	using QListView::scrollDirtyRegion;
	using QListView::dirtyRegionOffset;
	using QListView::startAutoScroll;
	using QListView::stopAutoScroll;
	using QListView::doAutoScroll;
	using QListView::dropIndicatorPosition;
	using QListView::setViewportMargins;
	using QListView::viewportMargins;
	using QListView::drawFrame;
	using QListView::updateMicroFocus;
	using QListView::create;
	using QListView::destroy;
	using QListView::focusNextChild;
	using QListView::focusPreviousChild;
	using QListView::sender;
	using QListView::senderSignalIndex;
	using QListView::receivers;
	using QListView::isSignalConnected;
	using QListView::getDecodedMetricF;
	bool cppgoBase_event(QEvent* __a0) { return QListView::event(__a0); }
	void cppgoBase_scrollContentsBy(int __a0, int __a1) { QListView::scrollContentsBy(__a0, __a1); }
	void cppgoBase_rowsInserted(const QModelIndex& __a0, int __a1, int __a2) { QListView::rowsInserted(__a0, __a1, __a2); }
	void cppgoBase_rowsAboutToBeRemoved(const QModelIndex& __a0, int __a1, int __a2) { QListView::rowsAboutToBeRemoved(__a0, __a1, __a2); }
	void cppgoBase_mouseMoveEvent(QMouseEvent* __a0) { QListView::mouseMoveEvent(__a0); }
	void cppgoBase_mouseReleaseEvent(QMouseEvent* __a0) { QListView::mouseReleaseEvent(__a0); }
	void cppgoBase_wheelEvent(QWheelEvent* __a0) { QListView::wheelEvent(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QListView::timerEvent(__a0); }
	void cppgoBase_resizeEvent(QResizeEvent* __a0) { QListView::resizeEvent(__a0); }
	void cppgoBase_dragMoveEvent(QDragMoveEvent* __a0) { QListView::dragMoveEvent(__a0); }
	void cppgoBase_dragLeaveEvent(QDragLeaveEvent* __a0) { QListView::dragLeaveEvent(__a0); }
	void cppgoBase_dropEvent(QDropEvent* __a0) { QListView::dropEvent(__a0); }
	void cppgoBase_startDrag(QFlags<Qt::DropAction> __a0) { QListView::startDrag(__a0); }
	void cppgoBase_initViewItemOption(QStyleOptionViewItem* __a0) { QListView::initViewItemOption(__a0); }
	void cppgoBase_paintEvent(QPaintEvent* __a0) { QListView::paintEvent(__a0); }
	int cppgoBase_horizontalOffset() { return QListView::horizontalOffset(); }
	int cppgoBase_verticalOffset() { return QListView::verticalOffset(); }
	QModelIndex cppgoBase_moveCursor(int __a0, QFlags<Qt::KeyboardModifier> __a1) { return QListView::moveCursor((QAbstractItemView::CursorAction)__a0, __a1); }
	void cppgoBase_setSelection(const QRect& __a0, QFlags<QItemSelectionModel::SelectionFlag> __a1) { QListView::setSelection(__a0, __a1); }
	QRegion cppgoBase_visualRegionForSelection(const QItemSelection& __a0) { return QListView::visualRegionForSelection(__a0); }
	QList<QModelIndex> cppgoBase_selectedIndexes() { return QListView::selectedIndexes(); }
	void cppgoBase_updateGeometries() { QListView::updateGeometries(); }
	bool cppgoBase_isIndexHidden(const QModelIndex& __a0) { return QListView::isIndexHidden(__a0); }
	void cppgoBase_selectionChanged(const QItemSelection& __a0, const QItemSelection& __a1) { QListView::selectionChanged(__a0, __a1); }
	void cppgoBase_currentChanged(const QModelIndex& __a0, const QModelIndex& __a1) { QListView::currentChanged(__a0, __a1); }
	QSize cppgoBase_viewportSizeHint() { return QListView::viewportSizeHint(); }
	void cppgoBase_updateEditorData() { QListView::updateEditorData(); }
	void cppgoBase_updateEditorGeometries() { QListView::updateEditorGeometries(); }
	void cppgoBase_verticalScrollbarAction(int __a0) { QListView::verticalScrollbarAction(__a0); }
	void cppgoBase_horizontalScrollbarAction(int __a0) { QListView::horizontalScrollbarAction(__a0); }
	void cppgoBase_verticalScrollbarValueChanged(int __a0) { QListView::verticalScrollbarValueChanged(__a0); }
	void cppgoBase_horizontalScrollbarValueChanged(int __a0) { QListView::horizontalScrollbarValueChanged(__a0); }
	void cppgoBase_closeEditor(QWidget* __a0, int __a1) { QListView::closeEditor(__a0, (QAbstractItemDelegate::EndEditHint)__a1); }
	void cppgoBase_commitData(QWidget* __a0) { QListView::commitData(__a0); }
	void cppgoBase_editorDestroyed(QObject* __a0) { QListView::editorDestroyed(__a0); }
	bool cppgoBase_edit(const QModelIndex& __a0, int __a1, QEvent* __a2) { return QListView::edit(__a0, (QAbstractItemView::EditTrigger)__a1, __a2); }
	QFlags<QItemSelectionModel::SelectionFlag> cppgoBase_selectionCommand(const QModelIndex& __a0, const QEvent* __a1) { return QListView::selectionCommand(__a0, __a1); }
	bool cppgoBase_focusNextPrevChild(bool __a0) { return QListView::focusNextPrevChild(__a0); }
	bool cppgoBase_viewportEvent(QEvent* __a0) { return QListView::viewportEvent(__a0); }
	void cppgoBase_mousePressEvent(QMouseEvent* __a0) { QListView::mousePressEvent(__a0); }
	void cppgoBase_mouseDoubleClickEvent(QMouseEvent* __a0) { QListView::mouseDoubleClickEvent(__a0); }
	void cppgoBase_dragEnterEvent(QDragEnterEvent* __a0) { QListView::dragEnterEvent(__a0); }
	void cppgoBase_focusInEvent(QFocusEvent* __a0) { QListView::focusInEvent(__a0); }
	void cppgoBase_focusOutEvent(QFocusEvent* __a0) { QListView::focusOutEvent(__a0); }
	void cppgoBase_keyPressEvent(QKeyEvent* __a0) { QListView::keyPressEvent(__a0); }
	void cppgoBase_inputMethodEvent(QInputMethodEvent* __a0) { QListView::inputMethodEvent(__a0); }
	bool cppgoBase_eventFilter(QObject* __a0, QEvent* __a1) { return QListView::eventFilter(__a0, __a1); }
	void cppgoBase_contextMenuEvent(QContextMenuEvent* __a0) { QListView::contextMenuEvent(__a0); }
	void cppgoBase_changeEvent(QEvent* __a0) { QListView::changeEvent(__a0); }
	void cppgoBase_initStyleOption(QStyleOptionFrame* __a0) { QListView::initStyleOption(__a0); }
	void cppgoBase_keyReleaseEvent(QKeyEvent* __a0) { QListView::keyReleaseEvent(__a0); }
	void cppgoBase_enterEvent(QEnterEvent* __a0) { QListView::enterEvent(__a0); }
	void cppgoBase_leaveEvent(QEvent* __a0) { QListView::leaveEvent(__a0); }
	void cppgoBase_moveEvent(QMoveEvent* __a0) { QListView::moveEvent(__a0); }
	void cppgoBase_closeEvent(QCloseEvent* __a0) { QListView::closeEvent(__a0); }
	void cppgoBase_tabletEvent(QTabletEvent* __a0) { QListView::tabletEvent(__a0); }
	void cppgoBase_actionEvent(QActionEvent* __a0) { QListView::actionEvent(__a0); }
	void cppgoBase_showEvent(QShowEvent* __a0) { QListView::showEvent(__a0); }
	void cppgoBase_hideEvent(QHideEvent* __a0) { QListView::hideEvent(__a0); }
	int cppgoBase_metric(int __a0) { return QListView::metric((QPaintDevice::PaintDeviceMetric)__a0); }
	void cppgoBase_initPainter(QPainter* __a0) { QListView::initPainter(__a0); }
	QPaintDevice* cppgoBase_redirected(QPoint* __a0) { return QListView::redirected(__a0); }
	QPainter* cppgoBase_sharedPainter() { return QListView::sharedPainter(); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QListView::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QListView::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QListView::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QListView::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QListView::metaObject();
	}
	QRect visualRect(const QModelIndex& index) const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, -1267696851865372997LL, (void*)&index, &__h); if (__h) return *(QRect*)__r; }
		return QListView::visualRect(index);
	}
	void scrollTo(const QModelIndex& index, QAbstractItemView::ScrollHint hint) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_void_voidp_int) { ((void(*)(long long, long long, void*, int, int*))cppgo_virt_void_voidp_int)(__goId, -4832340528568269947LL, (void*)&index, (int)(hint), &__h); if (__h) return; }
		QListView::scrollTo(index, hint);
	}
	QModelIndex indexAt(const QPoint& p) const override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, -7252823847285716802LL, (void*)&p, &__h); if (__h) return *(QModelIndex*)__r; }
		return QListView::indexAt(p);
	}
	void doItemsLayout() override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, -1888629738259754754LL, &__h); if (__h) return; }
		QListView::doItemsLayout();
	}
	void reset() override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 1086335023529244512LL, &__h); if (__h) return; }
		QListView::reset();
	}
	void setRootIndex(const QModelIndex& index) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1574077111192992763LL, (void*)&index, &__h); if (__h) return; }
		QListView::setRootIndex(index);
	}
	bool event(QEvent* e) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)e, &__h); if (__h) return __r; }
		return QListView::event(e);
	}
	void scrollContentsBy(int dx, int dy) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void_int_int) { ((void(*)(long long, long long, int, int, int*))cppgo_virt_void_int_int)(__goId, -6118001150126564745LL, dx, dy, &__h); if (__h) return; }
		QListView::scrollContentsBy(dx, dy);
	}
	void rowsInserted(const QModelIndex& parent, int start, int end) override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_void_voidp_int_int) { ((void(*)(long long, long long, void*, int, int, int*))cppgo_virt_void_voidp_int_int)(__goId, 4472434241583495620LL, (void*)&parent, start, end, &__h); if (__h) return; }
		QListView::rowsInserted(parent, start, end);
	}
	void rowsAboutToBeRemoved(const QModelIndex& parent, int start, int end) override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_void_voidp_int_int) { ((void(*)(long long, long long, void*, int, int, int*))cppgo_virt_void_voidp_int_int)(__goId, -8151049665126058157LL, (void*)&parent, start, end, &__h); if (__h) return; }
		QListView::rowsAboutToBeRemoved(parent, start, end);
	}
	void mouseMoveEvent(QMouseEvent* e) override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4255494004099806581LL, (void*)e, &__h); if (__h) return; }
		QListView::mouseMoveEvent(e);
	}
	void mouseReleaseEvent(QMouseEvent* e) override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4415920049297795213LL, (void*)e, &__h); if (__h) return; }
		QListView::mouseReleaseEvent(e);
	}
	void wheelEvent(QWheelEvent* e) override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1919988340934511316LL, (void*)e, &__h); if (__h) return; }
		QListView::wheelEvent(e);
	}
	void timerEvent(QTimerEvent* e) override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)e, &__h); if (__h) return; }
		QListView::timerEvent(e);
	}
	void resizeEvent(QResizeEvent* e) override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4302917285655014661LL, (void*)e, &__h); if (__h) return; }
		QListView::resizeEvent(e);
	}
	void dragMoveEvent(QDragMoveEvent* e) override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1781545137352745978LL, (void*)e, &__h); if (__h) return; }
		QListView::dragMoveEvent(e);
	}
	void dragLeaveEvent(QDragLeaveEvent* e) override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -3139904158373725006LL, (void*)e, &__h); if (__h) return; }
		QListView::dragLeaveEvent(e);
	}
	void dropEvent(QDropEvent* e) override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 9107058682754591716LL, (void*)e, &__h); if (__h) return; }
		QListView::dropEvent(e);
	}
	void startDrag(QFlags<Qt::DropAction> supportedActions) override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, -8112366604639385491LL, (int)(supportedActions), &__h); if (__h) return; }
		QListView::startDrag(supportedActions);
	}
	void initViewItemOption(QStyleOptionViewItem* option) const override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -7450787277202094556LL, (void*)option, &__h); if (__h) return; }
		QListView::initViewItemOption(option);
	}
	void paintEvent(QPaintEvent* e) override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5290442163167458465LL, (void*)e, &__h); if (__h) return; }
		QListView::paintEvent(e);
	}
	int horizontalOffset() const override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -1500310024801917802LL, &__h); if (__h) return __r; }
		return QListView::horizontalOffset();
	}
	int verticalOffset() const override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, 5218707810677921016LL, &__h); if (__h) return __r; }
		return QListView::verticalOffset();
	}
	QModelIndex moveCursor(QAbstractItemView::CursorAction cursorAction, QFlags<Qt::KeyboardModifier> modifiers) override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_voidp_int_int) { void* __r = ((void*(*)(long long, long long, int, int, int*))cppgo_virt_voidp_int_int)(__goId, -158523882194220890LL, (int)(cursorAction), (int)(modifiers), &__h); if (__h) return *(QModelIndex*)__r; }
		return QListView::moveCursor(cursorAction, modifiers);
	}
	void setSelection(const QRect& rect, QFlags<QItemSelectionModel::SelectionFlag> command) override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_void_voidp_int) { ((void(*)(long long, long long, void*, int, int*))cppgo_virt_void_voidp_int)(__goId, -1539266838375037305LL, (void*)&rect, (int)(command), &__h); if (__h) return; }
		QListView::setSelection(rect, command);
	}
	QRegion visualRegionForSelection(const QItemSelection& selection) const override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 7063795720810552074LL, (void*)&selection, &__h); if (__h) return *(QRegion*)__r; }
		return QListView::visualRegionForSelection(selection);
	}
	QList<QModelIndex> selectedIndexes() const override {
		int __h = 0;
		if (__goOn[27] && cppgo_virtp_voidp) {
			long long __len = 0;
			void* __a = ((void*(*)(long long, long long, long long*, int*))cppgo_virtp_voidp)(__goId, 740869982805721908LL, &__len, &__h);
			if (__h) {
				QList<QModelIndex> __r; __r.reserve(__len);
				void** __aa = (void**)__a;
				for (long long __i = 0; __i < __len; __i++) __r.append(*(QModelIndex*)__aa[__i]);
				if (__a) free(__a);
				return __r;
			}
		}
		return QListView::selectedIndexes();
	}
	void updateGeometries() override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 1022771470898460576LL, &__h); if (__h) return; }
		QListView::updateGeometries();
	}
	bool isIndexHidden(const QModelIndex& index) const override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 6057989667089378859LL, (void*)&index, &__h); if (__h) return __r; }
		return QListView::isIndexHidden(index);
	}
	void selectionChanged(const QItemSelection& selected, const QItemSelection& deselected) override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_void_voidp_voidp) { ((void(*)(long long, long long, void*, void*, int*))cppgo_virt_void_voidp_voidp)(__goId, -5892205462145366901LL, (void*)&selected, (void*)&deselected, &__h); if (__h) return; }
		QListView::selectionChanged(selected, deselected);
	}
	void currentChanged(const QModelIndex& current, const QModelIndex& previous) override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_void_voidp_voidp) { ((void(*)(long long, long long, void*, void*, int*))cppgo_virt_void_voidp_voidp)(__goId, -6978583972112316034LL, (void*)&current, (void*)&previous, &__h); if (__h) return; }
		QListView::currentChanged(current, previous);
	}
	QSize viewportSizeHint() const override {
		int __h = 0;
		if (__goOn[32] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 503579067252048239LL, &__h); if (__h) return *(QSize*)__r; }
		return QListView::viewportSizeHint();
	}
	void setModel(QAbstractItemModel* model) override {
		int __h = 0;
		if (__goOn[33] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4283922507359547360LL, (void*)model, &__h); if (__h) return; }
		QListView::setModel(model);
	}
	void setSelectionModel(QItemSelectionModel* selectionModel) override {
		int __h = 0;
		if (__goOn[34] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4689322347236366780LL, (void*)selectionModel, &__h); if (__h) return; }
		QListView::setSelectionModel(selectionModel);
	}
	void keyboardSearch(const QString& search) override {
		int __h = 0;
		QByteArray search_ba = search.toUtf8();
		if (__goOn[35] && cppgo_virt_void_constcharp_longlong) { ((void(*)(long long, long long, const char*, long long, int*))cppgo_virt_void_constcharp_longlong)(__goId, -5041452885022114200LL, (char*)search_ba.constData(), search_ba.size(), &__h); if (__h) return; }
		QListView::keyboardSearch(search);
	}
	int sizeHintForRow(int row) const override {
		int __h = 0;
		if (__goOn[36] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, -4084029023708556876LL, row, &__h); if (__h) return __r; }
		return QListView::sizeHintForRow(row);
	}
	int sizeHintForColumn(int column) const override {
		int __h = 0;
		if (__goOn[37] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, -3666792239577137638LL, column, &__h); if (__h) return __r; }
		return QListView::sizeHintForColumn(column);
	}
	QAbstractItemDelegate* itemDelegateForIndex(const QModelIndex& index) const override {
		int __h = 0;
		if (__goOn[38] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, -3811438228839363824LL, (void*)&index, &__h); if (__h) return (QAbstractItemDelegate*)__r; }
		return QListView::itemDelegateForIndex(index);
	}
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const override {
		int __h = 0;
		if (__goOn[39] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7345853097340709638LL, (int)(query), &__h); if (__h) return *(QVariant*)__r; }
		return QListView::inputMethodQuery(query);
	}
	void selectAll() override {
		int __h = 0;
		if (__goOn[40] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 8758113809152574300LL, &__h); if (__h) return; }
		QListView::selectAll();
	}
	void updateEditorData() override {
		int __h = 0;
		if (__goOn[41] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 9029941737538200621LL, &__h); if (__h) return; }
		QListView::updateEditorData();
	}
	void updateEditorGeometries() override {
		int __h = 0;
		if (__goOn[42] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, -8726300139121449691LL, &__h); if (__h) return; }
		QListView::updateEditorGeometries();
	}
	void verticalScrollbarAction(int action) override {
		int __h = 0;
		if (__goOn[43] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, -1155555122391250623LL, action, &__h); if (__h) return; }
		QListView::verticalScrollbarAction(action);
	}
	void horizontalScrollbarAction(int action) override {
		int __h = 0;
		if (__goOn[44] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, -785551405282638793LL, action, &__h); if (__h) return; }
		QListView::horizontalScrollbarAction(action);
	}
	void verticalScrollbarValueChanged(int value) override {
		int __h = 0;
		if (__goOn[45] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 6366953527806415524LL, value, &__h); if (__h) return; }
		QListView::verticalScrollbarValueChanged(value);
	}
	void horizontalScrollbarValueChanged(int value) override {
		int __h = 0;
		if (__goOn[46] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 1786430177580376614LL, value, &__h); if (__h) return; }
		QListView::horizontalScrollbarValueChanged(value);
	}
	void closeEditor(QWidget* editor, QAbstractItemDelegate::EndEditHint hint) override {
		int __h = 0;
		if (__goOn[47] && cppgo_virt_void_voidp_int) { ((void(*)(long long, long long, void*, int, int*))cppgo_virt_void_voidp_int)(__goId, 8648484264628816806LL, (void*)editor, (int)(hint), &__h); if (__h) return; }
		QListView::closeEditor(editor, hint);
	}
	void commitData(QWidget* editor) override {
		int __h = 0;
		if (__goOn[48] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6021383853323633948LL, (void*)editor, &__h); if (__h) return; }
		QListView::commitData(editor);
	}
	void editorDestroyed(QObject* editor) override {
		int __h = 0;
		if (__goOn[49] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6979418566323745651LL, (void*)editor, &__h); if (__h) return; }
		QListView::editorDestroyed(editor);
	}
	bool edit(const QModelIndex& index, QAbstractItemView::EditTrigger trigger, QEvent* event) override {
		int __h = 0;
		if (__goOn[50] && cppgo_virt_int_voidp_int_voidp) { int __r = ((int(*)(long long, long long, void*, int, void*, int*))cppgo_virt_int_voidp_int_voidp)(__goId, 4206689019978932209LL, (void*)&index, (int)(trigger), (void*)event, &__h); if (__h) return __r; }
		return QListView::edit(index, trigger, event);
	}
	QFlags<QItemSelectionModel::SelectionFlag> selectionCommand(const QModelIndex& index, const QEvent* event) const override {
		int __h = 0;
		if (__goOn[51] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, -2447490284072398302LL, (void*)&index, (void*)event, &__h); if (__h) return QFlags<QItemSelectionModel::SelectionFlag>(__r); }
		return QListView::selectionCommand(index, event);
	}
	bool focusNextPrevChild(bool next) override {
		int __h = 0;
		if (__goOn[52] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 5898328517505784191LL, next, &__h); if (__h) return __r; }
		return QListView::focusNextPrevChild(next);
	}
	bool viewportEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[53] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, -1971465653061861519LL, (void*)event, &__h); if (__h) return __r; }
		return QListView::viewportEvent(event);
	}
	void mousePressEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[54] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -332090898078543319LL, (void*)event, &__h); if (__h) return; }
		QListView::mousePressEvent(event);
	}
	void mouseDoubleClickEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[55] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4029073905697023887LL, (void*)event, &__h); if (__h) return; }
		QListView::mouseDoubleClickEvent(event);
	}
	void dragEnterEvent(QDragEnterEvent* event) override {
		int __h = 0;
		if (__goOn[56] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 206122715919917589LL, (void*)event, &__h); if (__h) return; }
		QListView::dragEnterEvent(event);
	}
	void focusInEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[57] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 975043767033552552LL, (void*)event, &__h); if (__h) return; }
		QListView::focusInEvent(event);
	}
	void focusOutEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[58] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -385596010576048993LL, (void*)event, &__h); if (__h) return; }
		QListView::focusOutEvent(event);
	}
	void keyPressEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[59] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1023791977679793185LL, (void*)event, &__h); if (__h) return; }
		QListView::keyPressEvent(event);
	}
	void inputMethodEvent(QInputMethodEvent* event) override {
		int __h = 0;
		if (__goOn[60] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5438819661942843222LL, (void*)event, &__h); if (__h) return; }
		QListView::inputMethodEvent(event);
	}
	bool eventFilter(QObject* object, QEvent* event) override {
		int __h = 0;
		if (__goOn[61] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)object, (void*)event, &__h); if (__h) return __r; }
		return QListView::eventFilter(object, event);
	}
	QSize minimumSizeHint() const override {
		int __h = 0;
		if (__goOn[62] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 7049264599837532491LL, &__h); if (__h) return *(QSize*)__r; }
		return QListView::minimumSizeHint();
	}
	QSize sizeHint() const override {
		int __h = 0;
		if (__goOn[63] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 729690261457684101LL, &__h); if (__h) return *(QSize*)__r; }
		return QListView::sizeHint();
	}
	void setupViewport(QWidget* viewport) override {
		int __h = 0;
		if (__goOn[64] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4531996381736730688LL, (void*)viewport, &__h); if (__h) return; }
		QListView::setupViewport(viewport);
	}
	void contextMenuEvent(QContextMenuEvent* arg1) override {
		int __h = 0;
		if (__goOn[65] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6698043222227006015LL, (void*)arg1, &__h); if (__h) return; }
		QListView::contextMenuEvent(arg1);
	}
	void changeEvent(QEvent* arg1) override {
		int __h = 0;
		if (__goOn[66] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4959376063826165943LL, (void*)arg1, &__h); if (__h) return; }
		QListView::changeEvent(arg1);
	}
	void initStyleOption(QStyleOptionFrame* option) const override {
		int __h = 0;
		if (__goOn[67] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -488693769771308531LL, (void*)option, &__h); if (__h) return; }
		QListView::initStyleOption(option);
	}
	int devType() const override {
		int __h = 0;
		if (__goOn[68] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8659679893868524044LL, &__h); if (__h) return __r; }
		return QListView::devType();
	}
	void setVisible(bool visible) override {
		int __h = 0;
		if (__goOn[69] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 2110182719375245575LL, visible, &__h); if (__h) return; }
		QListView::setVisible(visible);
	}
	int heightForWidth(int arg1) const override {
		int __h = 0;
		if (__goOn[70] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 4311534060495737339LL, arg1, &__h); if (__h) return __r; }
		return QListView::heightForWidth(arg1);
	}
	bool hasHeightForWidth() const override {
		int __h = 0;
		if (__goOn[71] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8266762397995423915LL, &__h); if (__h) return __r; }
		return QListView::hasHeightForWidth();
	}
	QPaintEngine* paintEngine() const override {
		int __h = 0;
		if (__goOn[72] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 6398007404756346381LL, &__h); if (__h) return (QPaintEngine*)__r; }
		return QListView::paintEngine();
	}
	void keyReleaseEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[73] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 772598475950050299LL, (void*)event, &__h); if (__h) return; }
		QListView::keyReleaseEvent(event);
	}
	void enterEvent(QEnterEvent* event) override {
		int __h = 0;
		if (__goOn[74] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6316270346668237895LL, (void*)event, &__h); if (__h) return; }
		QListView::enterEvent(event);
	}
	void leaveEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[75] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -551070009623711640LL, (void*)event, &__h); if (__h) return; }
		QListView::leaveEvent(event);
	}
	void moveEvent(QMoveEvent* event) override {
		int __h = 0;
		if (__goOn[76] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -9161949150489525396LL, (void*)event, &__h); if (__h) return; }
		QListView::moveEvent(event);
	}
	void closeEvent(QCloseEvent* event) override {
		int __h = 0;
		if (__goOn[77] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1989937050019464783LL, (void*)event, &__h); if (__h) return; }
		QListView::closeEvent(event);
	}
	void tabletEvent(QTabletEvent* event) override {
		int __h = 0;
		if (__goOn[78] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -212068866860701141LL, (void*)event, &__h); if (__h) return; }
		QListView::tabletEvent(event);
	}
	void actionEvent(QActionEvent* event) override {
		int __h = 0;
		if (__goOn[79] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1697444140088555581LL, (void*)event, &__h); if (__h) return; }
		QListView::actionEvent(event);
	}
	void showEvent(QShowEvent* event) override {
		int __h = 0;
		if (__goOn[80] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6825369174751916LL, (void*)event, &__h); if (__h) return; }
		QListView::showEvent(event);
	}
	void hideEvent(QHideEvent* event) override {
		int __h = 0;
		if (__goOn[81] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -199718707876719129LL, (void*)event, &__h); if (__h) return; }
		QListView::hideEvent(event);
	}
	bool nativeEvent(const QByteArray& eventType, void* message, qintptr* result) override {
		int __h = 0;
		if (__goOn[82] && cppgo_virt_int_voidp_voidp_longlongp) { int __r = ((int(*)(long long, long long, void*, void*, long long*, int*))cppgo_virt_int_voidp_voidp_longlongp)(__goId, -2039166586246443706LL, (void*)&eventType, (void*)message, (long long*)result, &__h); if (__h) return __r; }
		return QListView::nativeEvent(eventType, message, result);
	}
	int metric(QPaintDevice::PaintDeviceMetric arg1) const override {
		int __h = 0;
		if (__goOn[83] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 427629649267311777LL, (int)(arg1), &__h); if (__h) return __r; }
		return QListView::metric(arg1);
	}
	void initPainter(QPainter* painter) const override {
		int __h = 0;
		if (__goOn[84] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6477500296980050210LL, (void*)painter, &__h); if (__h) return; }
		QListView::initPainter(painter);
	}
	QPaintDevice* redirected(QPoint* offset) const override {
		int __h = 0;
		if (__goOn[85] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 4127240633310655048LL, (void*)offset, &__h); if (__h) return (QPaintDevice*)__r; }
		return QListView::redirected(offset);
	}
	QPainter* sharedPainter() const override {
		int __h = 0;
		if (__goOn[86] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 5792771328741774991LL, &__h); if (__h) return (QPainter*)__r; }
		return QListView::sharedPainter();
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[87] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QListView::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[88] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QListView::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[89] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QListView::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[90] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QListView::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QListView;
static std::mutex __connsMu_QListView;

extern "C" {

void* QListView_base(void* self) { return (void*)static_cast<QAbstractItemView*>((QListView*)self); }
void QListView_chain(void* self, void** out) {
	QListView* __p = (QListView*)self;
	QAbstractItemView* __b0 = static_cast<QAbstractItemView*>(__p); out[0] = (void*)__b0;
	QAbstractScrollArea* __b1 = static_cast<QAbstractScrollArea*>(__b0); out[1] = (void*)__b1;
	QFrame* __b2 = static_cast<QFrame*>(__b1); out[2] = (void*)__b2;
	QWidget* __b3 = static_cast<QWidget*>(__b2); out[3] = (void*)__b3;
	QObject* __b4 = static_cast<QObject*>(__b3); out[4] = (void*)__b4;
}
void QListView__set_override(void* self, int idx) { ((QListView__Go*)self)->__goOn[idx] = 1; }
void* QListView_fromQObject(void* o) { return (void*)qobject_cast<QListView*>((QObject*)o); }
void* QListView_new(long long __goid, void* parent) {
	return new QListView__Go(__goid, (QWidget*)parent);
}
void* QListView_new2(long long __goid) {
	return new QListView__Go(__goid);
}
char* QListView_tr(const char* s, long long* __out_len) {
	auto&& __r = QListView::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QListView_setMovement(void* self, int movement) {
	((QListView*)self)->setMovement((QListView::Movement)(movement));
}
int QListView_movement(void* self) {
	auto&& __r = ((QListView*)self)->movement();
	return (int)(__r);
}
void QListView_setFlow(void* self, int flow) {
	((QListView*)self)->setFlow((QListView::Flow)(flow));
}
int QListView_flow(void* self) {
	auto&& __r = ((QListView*)self)->flow();
	return (int)(__r);
}
void QListView_setWrapping(void* self, int enable) {
	((QListView*)self)->setWrapping(enable);
}
int QListView_isWrapping(void* self) {
	auto&& __r = ((QListView*)self)->isWrapping();
	return __r;
}
void QListView_setResizeMode(void* self, int mode) {
	((QListView*)self)->setResizeMode((QListView::ResizeMode)(mode));
}
int QListView_resizeMode(void* self) {
	auto&& __r = ((QListView*)self)->resizeMode();
	return (int)(__r);
}
void QListView_setLayoutMode(void* self, int mode) {
	((QListView*)self)->setLayoutMode((QListView::LayoutMode)(mode));
}
int QListView_layoutMode(void* self) {
	auto&& __r = ((QListView*)self)->layoutMode();
	return (int)(__r);
}
void QListView_setSpacing(void* self, int space) {
	((QListView*)self)->setSpacing(space);
}
int QListView_spacing(void* self) {
	auto&& __r = ((QListView*)self)->spacing();
	return __r;
}
void QListView_setBatchSize(void* self, int batchSize) {
	((QListView*)self)->setBatchSize(batchSize);
}
int QListView_batchSize(void* self) {
	auto&& __r = ((QListView*)self)->batchSize();
	return __r;
}
void QListView_setGridSize(void* self, void* size) {
	((QListView*)self)->setGridSize(*(QSize*)size);
}
void* QListView_gridSize(void* self) {
	auto&& __r = ((QListView*)self)->gridSize();
	return new QSize(__r);
}
void QListView_setViewMode(void* self, int mode) {
	((QListView*)self)->setViewMode((QListView::ViewMode)(mode));
}
int QListView_viewMode(void* self) {
	auto&& __r = ((QListView*)self)->viewMode();
	return (int)(__r);
}
void QListView_clearPropertyFlags(void* self) {
	((QListView*)self)->clearPropertyFlags();
}
int QListView_isRowHidden(void* self, int row) {
	auto&& __r = ((QListView*)self)->isRowHidden(row);
	return __r;
}
void QListView_setRowHidden(void* self, int row, int hide) {
	((QListView*)self)->setRowHidden(row, hide);
}
void QListView_setModelColumn(void* self, int column) {
	((QListView*)self)->setModelColumn(column);
}
int QListView_modelColumn(void* self) {
	auto&& __r = ((QListView*)self)->modelColumn();
	return __r;
}
void QListView_setUniformItemSizes(void* self, int enable) {
	((QListView*)self)->setUniformItemSizes(enable);
}
int QListView_uniformItemSizes(void* self) {
	auto&& __r = ((QListView*)self)->uniformItemSizes();
	return __r;
}
void QListView_setWordWrap(void* self, int on) {
	((QListView*)self)->setWordWrap(on);
}
int QListView_wordWrap(void* self) {
	auto&& __r = ((QListView*)self)->wordWrap();
	return __r;
}
void QListView_setSelectionRectVisible(void* self, int show) {
	((QListView*)self)->setSelectionRectVisible(show);
}
int QListView_isSelectionRectVisible(void* self) {
	auto&& __r = ((QListView*)self)->isSelectionRectVisible();
	return __r;
}
void QListView_setItemAlignment(void* self, int alignment) {
	((QListView*)self)->setItemAlignment(QFlags<Qt::AlignmentFlag>(alignment));
}
int QListView_itemAlignment(void* self) {
	auto&& __r = ((QListView*)self)->itemAlignment();
	return (int)(__r);
}
void QListView_indexesMoved(void* self, void** indexes, long long indexes_len) {
	QList<QModelIndex> indexes_list; indexes_list.reserve(indexes_len);
	for (long long __i = 0; __i < indexes_len; __i++) indexes_list.append(*(QModelIndex*)indexes[__i]);
	((QListView*)self)->indexesMoved(indexes_list);
}
void QListView_resizeContents(void* self, int width, int height) {
	((QListView__Go*)self)->resizeContents(width, height);
}
void* QListView_contentsSize(void* self) {
	auto&& __r = ((QListView__Go*)self)->contentsSize();
	return new QSize(__r);
}
void* QListView_rectForIndex(void* self, void* index) {
	auto&& __r = ((QListView__Go*)self)->rectForIndex(*(QModelIndex*)index);
	return new QRect(__r);
}
void QListView_setPositionForIndex(void* self, void* position, void* index) {
	((QListView__Go*)self)->setPositionForIndex(*(QPoint*)position, *(QModelIndex*)index);
}
int QListView_state(void* self) {
	auto&& __r = ((QListView__Go*)self)->state();
	return (int)(__r);
}
void QListView_scheduleDelayedItemsLayout(void* self) {
	((QListView__Go*)self)->scheduleDelayedItemsLayout();
}
void QListView_executeDelayedItemsLayout(void* self) {
	((QListView__Go*)self)->executeDelayedItemsLayout();
}
void QListView_setDirtyRegion(void* self, void* region) {
	((QListView__Go*)self)->setDirtyRegion(*(QRegion*)region);
}
void QListView_scrollDirtyRegion(void* self, int dx, int dy) {
	((QListView__Go*)self)->scrollDirtyRegion(dx, dy);
}
void* QListView_dirtyRegionOffset(void* self) {
	auto&& __r = ((QListView__Go*)self)->dirtyRegionOffset();
	return new QPoint(__r);
}
void QListView_startAutoScroll(void* self) {
	((QListView__Go*)self)->startAutoScroll();
}
void QListView_stopAutoScroll(void* self) {
	((QListView__Go*)self)->stopAutoScroll();
}
void QListView_doAutoScroll(void* self) {
	((QListView__Go*)self)->doAutoScroll();
}
int QListView_dropIndicatorPosition(void* self) {
	auto&& __r = ((QListView__Go*)self)->dropIndicatorPosition();
	return (int)(__r);
}
void QListView_setViewportMargins(void* self, int left, int top, int right, int bottom) {
	((QListView__Go*)self)->setViewportMargins(left, top, right, bottom);
}
void QListView_setViewportMargins2(void* self, void* margins) {
	((QListView__Go*)self)->setViewportMargins(*(QMargins*)margins);
}
void* QListView_viewportMargins(void* self) {
	auto&& __r = ((QListView__Go*)self)->viewportMargins();
	return new QMargins(__r);
}
void QListView_drawFrame(void* self, void* arg1) {
	((QListView__Go*)self)->drawFrame((QPainter*)arg1);
}
void QListView_setTabOrder(void* arg1, void* arg2) {
	QListView::setTabOrder((QWidget*)arg1, (QWidget*)arg2);
}
void* QListView_mouseGrabber() {
	auto&& __r = QListView::mouseGrabber();
	return (void*)(__r);
}
void* QListView_keyboardGrabber() {
	auto&& __r = QListView::keyboardGrabber();
	return (void*)(__r);
}
void* QListView_find(unsigned long long arg1) {
	auto&& __r = QListView::find(arg1);
	return (void*)(__r);
}
void* QListView_createWindowContainer(void* window) {
	auto&& __r = QListView::createWindowContainer((QWindow*)window);
	return (void*)(__r);
}
void QListView_updateMicroFocus(void* self) {
	((QListView__Go*)self)->updateMicroFocus();
}
void QListView_create(void* self) {
	((QListView__Go*)self)->create();
}
void QListView_destroy(void* self) {
	((QListView__Go*)self)->destroy();
}
int QListView_focusNextChild(void* self) {
	auto&& __r = ((QListView__Go*)self)->focusNextChild();
	return __r;
}
int QListView_focusPreviousChild(void* self) {
	auto&& __r = ((QListView__Go*)self)->focusPreviousChild();
	return __r;
}
void* QListView_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QListView::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QListView_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QListView::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QListView_sender(void* self) {
	auto&& __r = ((QListView__Go*)self)->sender();
	return (void*)(__r);
}
int QListView_senderSignalIndex(void* self) {
	auto&& __r = ((QListView__Go*)self)->senderSignalIndex();
	return __r;
}
int QListView_receivers(void* self, const char* signal) {
	auto&& __r = ((QListView__Go*)self)->receivers(signal);
	return __r;
}
int QListView_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QListView__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
double QListView_devicePixelRatioFScale() {
	auto&& __r = QListView::devicePixelRatioFScale();
	return __r;
}
int QListView_encodeMetricF(int metric, double value) {
	auto&& __r = QListView::encodeMetricF((QPaintDevice::PaintDeviceMetric)(metric), value);
	return __r;
}
double QListView_getDecodedMetricF(void* self, int metricA, int metricB) {
	auto&& __r = ((QListView__Go*)self)->getDecodedMetricF((QPaintDevice::PaintDeviceMetric)(metricA), (QPaintDevice::PaintDeviceMetric)(metricB));
	return __r;
}
char* QListView_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QListView::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QListView_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QListView::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QListView_createWindowContainer2(void* window, void* parent) {
	auto&& __r = QListView::createWindowContainer((QWindow*)window, (QWidget*)parent);
	return (void*)(__r);
}
void* QListView_createWindowContainer3(void* window, void* parent, int flags) {
	auto&& __r = QListView::createWindowContainer((QWindow*)window, (QWidget*)parent, QFlags<Qt::WindowType>(flags));
	return (void*)(__r);
}
void QListView_updateMicroFocus2(void* self, int query) {
	((QListView__Go*)self)->updateMicroFocus((Qt::InputMethodQuery)(query));
}
void QListView_create2(void* self, unsigned long long arg1) {
	((QListView__Go*)self)->create(arg1);
}
void QListView_create3(void* self, unsigned long long arg1, int initializeWindow) {
	((QListView__Go*)self)->create(arg1, initializeWindow);
}
void QListView_create4(void* self, unsigned long long arg1, int initializeWindow, int destroyOldWindow) {
	((QListView__Go*)self)->create(arg1, initializeWindow, destroyOldWindow);
}
void QListView_destroy2(void* self, int destroyWindow) {
	((QListView__Go*)self)->destroy(destroyWindow);
}
void QListView_destroy3(void* self, int destroyWindow, int destroySubWindows) {
	((QListView__Go*)self)->destroy(destroyWindow, destroySubWindows);
}
void* QListView_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QListView::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QListView_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QListView::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QListView_metaObject_default(void* self) {
	auto&& __r = ((QListView*)self)->QListView::metaObject();
	return (void*)(__r);
}
void* QListView_visualRect_default(void* self, void* index) {
	auto&& __r = ((QListView*)self)->QListView::visualRect(*(QModelIndex*)index);
	return new QRect(__r);
}
void QListView_scrollTo_default(void* self, void* index, int hint) {
	((QListView*)self)->QListView::scrollTo(*(QModelIndex*)index, (QAbstractItemView::ScrollHint)(hint));
}
void* QListView_indexAt_default(void* self, void* p) {
	auto&& __r = ((QListView*)self)->QListView::indexAt(*(QPoint*)p);
	return new QModelIndex(__r);
}
void QListView_doItemsLayout_default(void* self) {
	((QListView*)self)->QListView::doItemsLayout();
}
void QListView_reset_default(void* self) {
	((QListView*)self)->QListView::reset();
}
void QListView_setRootIndex_default(void* self, void* index) {
	((QListView*)self)->QListView::setRootIndex(*(QModelIndex*)index);
}
int QListView_event_default(void* self, void* e) {
	auto&& __r = ((QListView__Go*)self)->cppgoBase_event((QEvent*)e);
	return __r;
}
void QListView_scrollContentsBy_default(void* self, int dx, int dy) {
	((QListView__Go*)self)->cppgoBase_scrollContentsBy(dx, dy);
}
void QListView_rowsInserted_default(void* self, void* parent, int start, int end) {
	((QListView__Go*)self)->cppgoBase_rowsInserted(*(QModelIndex*)parent, start, end);
}
void QListView_rowsAboutToBeRemoved_default(void* self, void* parent, int start, int end) {
	((QListView__Go*)self)->cppgoBase_rowsAboutToBeRemoved(*(QModelIndex*)parent, start, end);
}
void QListView_mouseMoveEvent_default(void* self, void* e) {
	((QListView__Go*)self)->cppgoBase_mouseMoveEvent((QMouseEvent*)e);
}
void QListView_mouseReleaseEvent_default(void* self, void* e) {
	((QListView__Go*)self)->cppgoBase_mouseReleaseEvent((QMouseEvent*)e);
}
void QListView_wheelEvent_default(void* self, void* e) {
	((QListView__Go*)self)->cppgoBase_wheelEvent((QWheelEvent*)e);
}
void QListView_timerEvent_default(void* self, void* e) {
	((QListView__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)e);
}
void QListView_resizeEvent_default(void* self, void* e) {
	((QListView__Go*)self)->cppgoBase_resizeEvent((QResizeEvent*)e);
}
void QListView_dragMoveEvent_default(void* self, void* e) {
	((QListView__Go*)self)->cppgoBase_dragMoveEvent((QDragMoveEvent*)e);
}
void QListView_dragLeaveEvent_default(void* self, void* e) {
	((QListView__Go*)self)->cppgoBase_dragLeaveEvent((QDragLeaveEvent*)e);
}
void QListView_dropEvent_default(void* self, void* e) {
	((QListView__Go*)self)->cppgoBase_dropEvent((QDropEvent*)e);
}
void QListView_startDrag_default(void* self, int supportedActions) {
	((QListView__Go*)self)->cppgoBase_startDrag(QFlags<Qt::DropAction>(supportedActions));
}
void QListView_initViewItemOption_default(void* self, void* option) {
	((QListView__Go*)self)->cppgoBase_initViewItemOption((QStyleOptionViewItem*)option);
}
void QListView_paintEvent_default(void* self, void* e) {
	((QListView__Go*)self)->cppgoBase_paintEvent((QPaintEvent*)e);
}
int QListView_horizontalOffset_default(void* self) {
	auto&& __r = ((QListView__Go*)self)->cppgoBase_horizontalOffset();
	return __r;
}
int QListView_verticalOffset_default(void* self) {
	auto&& __r = ((QListView__Go*)self)->cppgoBase_verticalOffset();
	return __r;
}
void* QListView_moveCursor_default(void* self, int cursorAction, int modifiers) {
	auto&& __r = ((QListView__Go*)self)->cppgoBase_moveCursor(cursorAction, QFlags<Qt::KeyboardModifier>(modifiers));
	return new QModelIndex(__r);
}
void QListView_setSelection_default(void* self, void* rect, int command) {
	((QListView__Go*)self)->cppgoBase_setSelection(*(QRect*)rect, QFlags<QItemSelectionModel::SelectionFlag>(command));
}
void* QListView_visualRegionForSelection_default(void* self, void* selection) {
	auto&& __r = ((QListView__Go*)self)->cppgoBase_visualRegionForSelection(*(QItemSelection*)selection);
	return new QRegion(__r);
}
void** QListView_selectedIndexes_default(void* self, long long* __out_len) {
	auto&& __r = ((QListView__Go*)self)->cppgoBase_selectedIndexes();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QModelIndex(__e);
	return __arr;
}
void QListView_updateGeometries_default(void* self) {
	((QListView__Go*)self)->cppgoBase_updateGeometries();
}
int QListView_isIndexHidden_default(void* self, void* index) {
	auto&& __r = ((QListView__Go*)self)->cppgoBase_isIndexHidden(*(QModelIndex*)index);
	return __r;
}
void QListView_selectionChanged_default(void* self, void* selected, void* deselected) {
	((QListView__Go*)self)->cppgoBase_selectionChanged(*(QItemSelection*)selected, *(QItemSelection*)deselected);
}
void QListView_currentChanged_default(void* self, void* current, void* previous) {
	((QListView__Go*)self)->cppgoBase_currentChanged(*(QModelIndex*)current, *(QModelIndex*)previous);
}
void* QListView_viewportSizeHint_default(void* self) {
	auto&& __r = ((QListView__Go*)self)->cppgoBase_viewportSizeHint();
	return new QSize(__r);
}
void QListView_setModel_default(void* self, void* model) {
	((QListView*)self)->QListView::setModel((QAbstractItemModel*)model);
}
void QListView_setSelectionModel_default(void* self, void* selectionModel) {
	((QListView*)self)->QListView::setSelectionModel((QItemSelectionModel*)selectionModel);
}
void QListView_keyboardSearch_default(void* self, const char* search, long long search_len) {
	QString search_qstr = QString::fromUtf8(search, search_len);
	((QListView*)self)->QListView::keyboardSearch(search_qstr);
}
int QListView_sizeHintForRow_default(void* self, int row) {
	auto&& __r = ((QListView*)self)->QListView::sizeHintForRow(row);
	return __r;
}
int QListView_sizeHintForColumn_default(void* self, int column) {
	auto&& __r = ((QListView*)self)->QListView::sizeHintForColumn(column);
	return __r;
}
void* QListView_itemDelegateForIndex_default(void* self, void* index) {
	auto&& __r = ((QListView*)self)->QListView::itemDelegateForIndex(*(QModelIndex*)index);
	return (void*)(__r);
}
void* QListView_inputMethodQuery_default(void* self, int query) {
	auto&& __r = ((QListView*)self)->QListView::inputMethodQuery((Qt::InputMethodQuery)(query));
	return new QVariant(__r);
}
void QListView_selectAll_default(void* self) {
	((QListView*)self)->QListView::selectAll();
}
void QListView_updateEditorData_default(void* self) {
	((QListView__Go*)self)->cppgoBase_updateEditorData();
}
void QListView_updateEditorGeometries_default(void* self) {
	((QListView__Go*)self)->cppgoBase_updateEditorGeometries();
}
void QListView_verticalScrollbarAction_default(void* self, int action) {
	((QListView__Go*)self)->cppgoBase_verticalScrollbarAction(action);
}
void QListView_horizontalScrollbarAction_default(void* self, int action) {
	((QListView__Go*)self)->cppgoBase_horizontalScrollbarAction(action);
}
void QListView_verticalScrollbarValueChanged_default(void* self, int value) {
	((QListView__Go*)self)->cppgoBase_verticalScrollbarValueChanged(value);
}
void QListView_horizontalScrollbarValueChanged_default(void* self, int value) {
	((QListView__Go*)self)->cppgoBase_horizontalScrollbarValueChanged(value);
}
void QListView_closeEditor_default(void* self, void* editor, int hint) {
	((QListView__Go*)self)->cppgoBase_closeEditor((QWidget*)editor, hint);
}
void QListView_commitData_default(void* self, void* editor) {
	((QListView__Go*)self)->cppgoBase_commitData((QWidget*)editor);
}
void QListView_editorDestroyed_default(void* self, void* editor) {
	((QListView__Go*)self)->cppgoBase_editorDestroyed((QObject*)editor);
}
int QListView_edit_default(void* self, void* index, int trigger, void* event) {
	auto&& __r = ((QListView__Go*)self)->cppgoBase_edit(*(QModelIndex*)index, trigger, (QEvent*)event);
	return __r;
}
int QListView_selectionCommand_default(void* self, void* index, void* event) {
	auto&& __r = ((QListView__Go*)self)->cppgoBase_selectionCommand(*(QModelIndex*)index, (QEvent*)event);
	return (int)(__r);
}
int QListView_focusNextPrevChild_default(void* self, int next) {
	auto&& __r = ((QListView__Go*)self)->cppgoBase_focusNextPrevChild(next);
	return __r;
}
int QListView_viewportEvent_default(void* self, void* event) {
	auto&& __r = ((QListView__Go*)self)->cppgoBase_viewportEvent((QEvent*)event);
	return __r;
}
void QListView_mousePressEvent_default(void* self, void* event) {
	((QListView__Go*)self)->cppgoBase_mousePressEvent((QMouseEvent*)event);
}
void QListView_mouseDoubleClickEvent_default(void* self, void* event) {
	((QListView__Go*)self)->cppgoBase_mouseDoubleClickEvent((QMouseEvent*)event);
}
void QListView_dragEnterEvent_default(void* self, void* event) {
	((QListView__Go*)self)->cppgoBase_dragEnterEvent((QDragEnterEvent*)event);
}
void QListView_focusInEvent_default(void* self, void* event) {
	((QListView__Go*)self)->cppgoBase_focusInEvent((QFocusEvent*)event);
}
void QListView_focusOutEvent_default(void* self, void* event) {
	((QListView__Go*)self)->cppgoBase_focusOutEvent((QFocusEvent*)event);
}
void QListView_keyPressEvent_default(void* self, void* event) {
	((QListView__Go*)self)->cppgoBase_keyPressEvent((QKeyEvent*)event);
}
void QListView_inputMethodEvent_default(void* self, void* event) {
	((QListView__Go*)self)->cppgoBase_inputMethodEvent((QInputMethodEvent*)event);
}
int QListView_eventFilter_default(void* self, void* object, void* event) {
	auto&& __r = ((QListView__Go*)self)->cppgoBase_eventFilter((QObject*)object, (QEvent*)event);
	return __r;
}
void* QListView_minimumSizeHint_default(void* self) {
	auto&& __r = ((QListView*)self)->QListView::minimumSizeHint();
	return new QSize(__r);
}
void* QListView_sizeHint_default(void* self) {
	auto&& __r = ((QListView*)self)->QListView::sizeHint();
	return new QSize(__r);
}
void QListView_setupViewport_default(void* self, void* viewport) {
	((QListView*)self)->QListView::setupViewport((QWidget*)viewport);
}
void QListView_contextMenuEvent_default(void* self, void* arg1) {
	((QListView__Go*)self)->cppgoBase_contextMenuEvent((QContextMenuEvent*)arg1);
}
void QListView_changeEvent_default(void* self, void* arg1) {
	((QListView__Go*)self)->cppgoBase_changeEvent((QEvent*)arg1);
}
void QListView_initStyleOption_default(void* self, void* option) {
	((QListView__Go*)self)->cppgoBase_initStyleOption((QStyleOptionFrame*)option);
}
int QListView_devType_default(void* self) {
	auto&& __r = ((QListView*)self)->QListView::devType();
	return __r;
}
void QListView_setVisible_default(void* self, int visible) {
	((QListView*)self)->QListView::setVisible(visible);
}
int QListView_heightForWidth_default(void* self, int arg1) {
	auto&& __r = ((QListView*)self)->QListView::heightForWidth(arg1);
	return __r;
}
int QListView_hasHeightForWidth_default(void* self) {
	auto&& __r = ((QListView*)self)->QListView::hasHeightForWidth();
	return __r;
}
void* QListView_paintEngine_default(void* self) {
	auto&& __r = ((QListView*)self)->QListView::paintEngine();
	return (void*)(__r);
}
void QListView_keyReleaseEvent_default(void* self, void* event) {
	((QListView__Go*)self)->cppgoBase_keyReleaseEvent((QKeyEvent*)event);
}
void QListView_enterEvent_default(void* self, void* event) {
	((QListView__Go*)self)->cppgoBase_enterEvent((QEnterEvent*)event);
}
void QListView_leaveEvent_default(void* self, void* event) {
	((QListView__Go*)self)->cppgoBase_leaveEvent((QEvent*)event);
}
void QListView_moveEvent_default(void* self, void* event) {
	((QListView__Go*)self)->cppgoBase_moveEvent((QMoveEvent*)event);
}
void QListView_closeEvent_default(void* self, void* event) {
	((QListView__Go*)self)->cppgoBase_closeEvent((QCloseEvent*)event);
}
void QListView_tabletEvent_default(void* self, void* event) {
	((QListView__Go*)self)->cppgoBase_tabletEvent((QTabletEvent*)event);
}
void QListView_actionEvent_default(void* self, void* event) {
	((QListView__Go*)self)->cppgoBase_actionEvent((QActionEvent*)event);
}
void QListView_showEvent_default(void* self, void* event) {
	((QListView__Go*)self)->cppgoBase_showEvent((QShowEvent*)event);
}
void QListView_hideEvent_default(void* self, void* event) {
	((QListView__Go*)self)->cppgoBase_hideEvent((QHideEvent*)event);
}
int QListView_metric_default(void* self, int arg1) {
	auto&& __r = ((QListView__Go*)self)->cppgoBase_metric(arg1);
	return __r;
}
void QListView_initPainter_default(void* self, void* painter) {
	((QListView__Go*)self)->cppgoBase_initPainter((QPainter*)painter);
}
void* QListView_redirected_default(void* self, void* offset) {
	auto&& __r = ((QListView__Go*)self)->cppgoBase_redirected((QPoint*)offset);
	return (void*)(__r);
}
void* QListView_sharedPainter_default(void* self) {
	auto&& __r = ((QListView__Go*)self)->cppgoBase_sharedPainter();
	return (void*)(__r);
}
void QListView_childEvent_default(void* self, void* event) {
	((QListView__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QListView_customEvent_default(void* self, void* event) {
	((QListView__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QListView_connectNotify_default(void* self, void* signal) {
	((QListView__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QListView_disconnectNotify_default(void* self, void* signal) {
	((QListView__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QListView_connect_indexesMoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QListView*)self, &QListView::indexesMoved, (QListView*)self, [__id](const QList<QModelIndex>& indexes){
		std::vector<void*> indexes_v; indexes_v.reserve(indexes.size());
		for (long long __i = 0; __i < indexes.size(); __i++) indexes_v.push_back((void*)&indexes[__i]);
		if (cppgo_sig_void_voidpp_longlong) ((void(*)(long long, void**, long long))cppgo_sig_void_voidpp_longlong)(__id, indexes_v.data(), (long long)indexes.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QListView);
	__conns_QListView[__id] = __c;
}
void QListView_connect_pressed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QListView*)self, &QListView::pressed, (QListView*)self, [__id](const QModelIndex& index){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&index);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QListView);
	__conns_QListView[__id] = __c;
}
void QListView_connect_clicked(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QListView*)self, &QListView::clicked, (QListView*)self, [__id](const QModelIndex& index){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&index);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QListView);
	__conns_QListView[__id] = __c;
}
void QListView_connect_doubleClicked(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QListView*)self, &QListView::doubleClicked, (QListView*)self, [__id](const QModelIndex& index){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&index);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QListView);
	__conns_QListView[__id] = __c;
}
void QListView_connect_activated(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QListView*)self, &QListView::activated, (QListView*)self, [__id](const QModelIndex& index){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&index);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QListView);
	__conns_QListView[__id] = __c;
}
void QListView_connect_entered(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QListView*)self, &QListView::entered, (QListView*)self, [__id](const QModelIndex& index){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&index);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QListView);
	__conns_QListView[__id] = __c;
}
void QListView_connect_viewportEntered(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QListView*)self, &QListView::viewportEntered, (QListView*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QListView);
	__conns_QListView[__id] = __c;
}
void QListView_connect_iconSizeChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QListView*)self, &QListView::iconSizeChanged, (QListView*)self, [__id](const QSize& size){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&size);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QListView);
	__conns_QListView[__id] = __c;
}
void QListView_connect_windowTitleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QListView*)self, &QListView::windowTitleChanged, (QListView*)self, [__id](const QString& title){
		QByteArray title_ba = title.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)title_ba.constData(), title_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QListView);
	__conns_QListView[__id] = __c;
}
void QListView_connect_windowIconChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QListView*)self, &QListView::windowIconChanged, (QListView*)self, [__id](const QIcon& icon){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&icon);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QListView);
	__conns_QListView[__id] = __c;
}
void QListView_connect_windowIconTextChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QListView*)self, &QListView::windowIconTextChanged, (QListView*)self, [__id](const QString& iconText){
		QByteArray iconText_ba = iconText.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)iconText_ba.constData(), iconText_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QListView);
	__conns_QListView[__id] = __c;
}
void QListView_connect_customContextMenuRequested(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QListView*)self, &QListView::customContextMenuRequested, (QListView*)self, [__id](const QPoint& pos){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&pos);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QListView);
	__conns_QListView[__id] = __c;
}
void QListView_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QListView*)self, &QListView::destroyed, (QListView*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QListView);
	__conns_QListView[__id] = __c;
}
void QListView_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QListView*)self, &QListView::destroyed, (QListView*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QListView);
	__conns_QListView[__id] = __c;
}
void QListView_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QListView*)self, &QListView::objectNameChanged, (QListView*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QListView);
	__conns_QListView[__id] = __c;
}
void QListView__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QListView);
		auto __it = __conns_QListView.find(id);
		if (__it == __conns_QListView.end()) return;
		__c = __it->second;
		__conns_QListView.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QListView__Go_new(long long id) { return new QListView__Go(id); }
void QListView_delete(void* self) { delete (QListView*)self; }

}

#include <QtWidgets/qsplashscreen.h>
#include <QtWidgets/qwidget.h>
#include <QtGui/qpixmap.h>
#include <QtCore/qnamespace.h>
#include <QtGui/qscreen.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtGui/qcolor.h>
#include <QtCore/qcoreevent.h>
#include <QtGui/qpainter.h>
#include <QtGui/qevent.h>
#include <QtWidgets/qstyle.h>
#include <QtCore/qrect.h>
#include <QtCore/qpoint.h>
#include <QtCore/qsize.h>
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
#include <QtCore/qvariant.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qsplashscreen.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QSplashScreen__Go : public QSplashScreen {
public:
	long long __goId;
	unsigned char __goOn[49] = {};
	template<typename... A> QSplashScreen__Go(long long id, A&&... a) : QSplashScreen(std::forward<A>(a)...), __goId(id) {}
	using QSplashScreen::updateMicroFocus;
	using QSplashScreen::create;
	using QSplashScreen::destroy;
	using QSplashScreen::focusNextChild;
	using QSplashScreen::focusPreviousChild;
	using QSplashScreen::sender;
	using QSplashScreen::senderSignalIndex;
	using QSplashScreen::receivers;
	using QSplashScreen::isSignalConnected;
	using QSplashScreen::getDecodedMetricF;
	bool cppgoBase_event(QEvent* __a0) { return QSplashScreen::event(__a0); }
	void cppgoBase_drawContents(QPainter* __a0) { QSplashScreen::drawContents(__a0); }
	void cppgoBase_mousePressEvent(QMouseEvent* __a0) { QSplashScreen::mousePressEvent(__a0); }
	void cppgoBase_mouseReleaseEvent(QMouseEvent* __a0) { QSplashScreen::mouseReleaseEvent(__a0); }
	void cppgoBase_mouseDoubleClickEvent(QMouseEvent* __a0) { QSplashScreen::mouseDoubleClickEvent(__a0); }
	void cppgoBase_mouseMoveEvent(QMouseEvent* __a0) { QSplashScreen::mouseMoveEvent(__a0); }
	void cppgoBase_wheelEvent(QWheelEvent* __a0) { QSplashScreen::wheelEvent(__a0); }
	void cppgoBase_keyPressEvent(QKeyEvent* __a0) { QSplashScreen::keyPressEvent(__a0); }
	void cppgoBase_keyReleaseEvent(QKeyEvent* __a0) { QSplashScreen::keyReleaseEvent(__a0); }
	void cppgoBase_focusInEvent(QFocusEvent* __a0) { QSplashScreen::focusInEvent(__a0); }
	void cppgoBase_focusOutEvent(QFocusEvent* __a0) { QSplashScreen::focusOutEvent(__a0); }
	void cppgoBase_enterEvent(QEnterEvent* __a0) { QSplashScreen::enterEvent(__a0); }
	void cppgoBase_leaveEvent(QEvent* __a0) { QSplashScreen::leaveEvent(__a0); }
	void cppgoBase_paintEvent(QPaintEvent* __a0) { QSplashScreen::paintEvent(__a0); }
	void cppgoBase_moveEvent(QMoveEvent* __a0) { QSplashScreen::moveEvent(__a0); }
	void cppgoBase_resizeEvent(QResizeEvent* __a0) { QSplashScreen::resizeEvent(__a0); }
	void cppgoBase_closeEvent(QCloseEvent* __a0) { QSplashScreen::closeEvent(__a0); }
	void cppgoBase_contextMenuEvent(QContextMenuEvent* __a0) { QSplashScreen::contextMenuEvent(__a0); }
	void cppgoBase_tabletEvent(QTabletEvent* __a0) { QSplashScreen::tabletEvent(__a0); }
	void cppgoBase_actionEvent(QActionEvent* __a0) { QSplashScreen::actionEvent(__a0); }
	void cppgoBase_dragEnterEvent(QDragEnterEvent* __a0) { QSplashScreen::dragEnterEvent(__a0); }
	void cppgoBase_dragMoveEvent(QDragMoveEvent* __a0) { QSplashScreen::dragMoveEvent(__a0); }
	void cppgoBase_dragLeaveEvent(QDragLeaveEvent* __a0) { QSplashScreen::dragLeaveEvent(__a0); }
	void cppgoBase_dropEvent(QDropEvent* __a0) { QSplashScreen::dropEvent(__a0); }
	void cppgoBase_showEvent(QShowEvent* __a0) { QSplashScreen::showEvent(__a0); }
	void cppgoBase_hideEvent(QHideEvent* __a0) { QSplashScreen::hideEvent(__a0); }
	void cppgoBase_changeEvent(QEvent* __a0) { QSplashScreen::changeEvent(__a0); }
	int cppgoBase_metric(int __a0) { return QSplashScreen::metric((QPaintDevice::PaintDeviceMetric)__a0); }
	void cppgoBase_initPainter(QPainter* __a0) { QSplashScreen::initPainter(__a0); }
	QPaintDevice* cppgoBase_redirected(QPoint* __a0) { return QSplashScreen::redirected(__a0); }
	QPainter* cppgoBase_sharedPainter() { return QSplashScreen::sharedPainter(); }
	void cppgoBase_inputMethodEvent(QInputMethodEvent* __a0) { QSplashScreen::inputMethodEvent(__a0); }
	bool cppgoBase_focusNextPrevChild(bool __a0) { return QSplashScreen::focusNextPrevChild(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QSplashScreen::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QSplashScreen::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QSplashScreen::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QSplashScreen::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QSplashScreen::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QSplashScreen::metaObject();
	}
	bool event(QEvent* e) override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)e, &__h); if (__h) return __r; }
		return QSplashScreen::event(e);
	}
	void drawContents(QPainter* painter) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -2902866686677005435LL, (void*)painter, &__h); if (__h) return; }
		QSplashScreen::drawContents(painter);
	}
	void mousePressEvent(QMouseEvent* arg1) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -332090898078543319LL, (void*)arg1, &__h); if (__h) return; }
		QSplashScreen::mousePressEvent(arg1);
	}
	int devType() const override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8659679893868524044LL, &__h); if (__h) return __r; }
		return QSplashScreen::devType();
	}
	void setVisible(bool visible) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 2110182719375245575LL, visible, &__h); if (__h) return; }
		QSplashScreen::setVisible(visible);
	}
	QSize sizeHint() const override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 729690261457684101LL, &__h); if (__h) return *(QSize*)__r; }
		return QSplashScreen::sizeHint();
	}
	QSize minimumSizeHint() const override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 7049264599837532491LL, &__h); if (__h) return *(QSize*)__r; }
		return QSplashScreen::minimumSizeHint();
	}
	int heightForWidth(int arg1) const override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 4311534060495737339LL, arg1, &__h); if (__h) return __r; }
		return QSplashScreen::heightForWidth(arg1);
	}
	bool hasHeightForWidth() const override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8266762397995423915LL, &__h); if (__h) return __r; }
		return QSplashScreen::hasHeightForWidth();
	}
	QPaintEngine* paintEngine() const override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 6398007404756346381LL, &__h); if (__h) return (QPaintEngine*)__r; }
		return QSplashScreen::paintEngine();
	}
	void mouseReleaseEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4415920049297795213LL, (void*)event, &__h); if (__h) return; }
		QSplashScreen::mouseReleaseEvent(event);
	}
	void mouseDoubleClickEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4029073905697023887LL, (void*)event, &__h); if (__h) return; }
		QSplashScreen::mouseDoubleClickEvent(event);
	}
	void mouseMoveEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4255494004099806581LL, (void*)event, &__h); if (__h) return; }
		QSplashScreen::mouseMoveEvent(event);
	}
	void wheelEvent(QWheelEvent* event) override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1919988340934511316LL, (void*)event, &__h); if (__h) return; }
		QSplashScreen::wheelEvent(event);
	}
	void keyPressEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1023791977679793185LL, (void*)event, &__h); if (__h) return; }
		QSplashScreen::keyPressEvent(event);
	}
	void keyReleaseEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 772598475950050299LL, (void*)event, &__h); if (__h) return; }
		QSplashScreen::keyReleaseEvent(event);
	}
	void focusInEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 975043767033552552LL, (void*)event, &__h); if (__h) return; }
		QSplashScreen::focusInEvent(event);
	}
	void focusOutEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -385596010576048993LL, (void*)event, &__h); if (__h) return; }
		QSplashScreen::focusOutEvent(event);
	}
	void enterEvent(QEnterEvent* event) override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6316270346668237895LL, (void*)event, &__h); if (__h) return; }
		QSplashScreen::enterEvent(event);
	}
	void leaveEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -551070009623711640LL, (void*)event, &__h); if (__h) return; }
		QSplashScreen::leaveEvent(event);
	}
	void paintEvent(QPaintEvent* event) override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5290442163167458465LL, (void*)event, &__h); if (__h) return; }
		QSplashScreen::paintEvent(event);
	}
	void moveEvent(QMoveEvent* event) override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -9161949150489525396LL, (void*)event, &__h); if (__h) return; }
		QSplashScreen::moveEvent(event);
	}
	void resizeEvent(QResizeEvent* event) override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4302917285655014661LL, (void*)event, &__h); if (__h) return; }
		QSplashScreen::resizeEvent(event);
	}
	void closeEvent(QCloseEvent* event) override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1989937050019464783LL, (void*)event, &__h); if (__h) return; }
		QSplashScreen::closeEvent(event);
	}
	void contextMenuEvent(QContextMenuEvent* event) override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6698043222227006015LL, (void*)event, &__h); if (__h) return; }
		QSplashScreen::contextMenuEvent(event);
	}
	void tabletEvent(QTabletEvent* event) override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -212068866860701141LL, (void*)event, &__h); if (__h) return; }
		QSplashScreen::tabletEvent(event);
	}
	void actionEvent(QActionEvent* event) override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1697444140088555581LL, (void*)event, &__h); if (__h) return; }
		QSplashScreen::actionEvent(event);
	}
	void dragEnterEvent(QDragEnterEvent* event) override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 206122715919917589LL, (void*)event, &__h); if (__h) return; }
		QSplashScreen::dragEnterEvent(event);
	}
	void dragMoveEvent(QDragMoveEvent* event) override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1781545137352745978LL, (void*)event, &__h); if (__h) return; }
		QSplashScreen::dragMoveEvent(event);
	}
	void dragLeaveEvent(QDragLeaveEvent* event) override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -3139904158373725006LL, (void*)event, &__h); if (__h) return; }
		QSplashScreen::dragLeaveEvent(event);
	}
	void dropEvent(QDropEvent* event) override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 9107058682754591716LL, (void*)event, &__h); if (__h) return; }
		QSplashScreen::dropEvent(event);
	}
	void showEvent(QShowEvent* event) override {
		int __h = 0;
		if (__goOn[32] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6825369174751916LL, (void*)event, &__h); if (__h) return; }
		QSplashScreen::showEvent(event);
	}
	void hideEvent(QHideEvent* event) override {
		int __h = 0;
		if (__goOn[33] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -199718707876719129LL, (void*)event, &__h); if (__h) return; }
		QSplashScreen::hideEvent(event);
	}
	bool nativeEvent(const QByteArray& eventType, void* message, qintptr* result) override {
		int __h = 0;
		if (__goOn[34] && cppgo_virt_int_voidp_voidp_longlongp) { int __r = ((int(*)(long long, long long, void*, void*, long long*, int*))cppgo_virt_int_voidp_voidp_longlongp)(__goId, -2039166586246443706LL, (void*)&eventType, (void*)message, (long long*)result, &__h); if (__h) return __r; }
		return QSplashScreen::nativeEvent(eventType, message, result);
	}
	void changeEvent(QEvent* arg1) override {
		int __h = 0;
		if (__goOn[35] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4959376063826165943LL, (void*)arg1, &__h); if (__h) return; }
		QSplashScreen::changeEvent(arg1);
	}
	int metric(QPaintDevice::PaintDeviceMetric arg1) const override {
		int __h = 0;
		if (__goOn[36] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 427629649267311777LL, (int)(arg1), &__h); if (__h) return __r; }
		return QSplashScreen::metric(arg1);
	}
	void initPainter(QPainter* painter) const override {
		int __h = 0;
		if (__goOn[37] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6477500296980050210LL, (void*)painter, &__h); if (__h) return; }
		QSplashScreen::initPainter(painter);
	}
	QPaintDevice* redirected(QPoint* offset) const override {
		int __h = 0;
		if (__goOn[38] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 4127240633310655048LL, (void*)offset, &__h); if (__h) return (QPaintDevice*)__r; }
		return QSplashScreen::redirected(offset);
	}
	QPainter* sharedPainter() const override {
		int __h = 0;
		if (__goOn[39] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 5792771328741774991LL, &__h); if (__h) return (QPainter*)__r; }
		return QSplashScreen::sharedPainter();
	}
	void inputMethodEvent(QInputMethodEvent* arg1) override {
		int __h = 0;
		if (__goOn[40] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5438819661942843222LL, (void*)arg1, &__h); if (__h) return; }
		QSplashScreen::inputMethodEvent(arg1);
	}
	QVariant inputMethodQuery(Qt::InputMethodQuery arg1) const override {
		int __h = 0;
		if (__goOn[41] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7345853097340709638LL, (int)(arg1), &__h); if (__h) return *(QVariant*)__r; }
		return QSplashScreen::inputMethodQuery(arg1);
	}
	bool focusNextPrevChild(bool next) override {
		int __h = 0;
		if (__goOn[42] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 5898328517505784191LL, next, &__h); if (__h) return __r; }
		return QSplashScreen::focusNextPrevChild(next);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[43] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QSplashScreen::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[44] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QSplashScreen::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[45] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QSplashScreen::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[46] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QSplashScreen::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[47] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QSplashScreen::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[48] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QSplashScreen::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QSplashScreen;
static std::mutex __connsMu_QSplashScreen;

extern "C" {

void* QSplashScreen_base(void* self) { return (void*)static_cast<QWidget*>((QSplashScreen*)self); }
void QSplashScreen_chain(void* self, void** out) {
	QSplashScreen* __p = (QSplashScreen*)self;
	QWidget* __b0 = static_cast<QWidget*>(__p); out[0] = (void*)__b0;
	QObject* __b1 = static_cast<QObject*>(__b0); out[1] = (void*)__b1;
}
void QSplashScreen__set_override(void* self, int idx) { ((QSplashScreen__Go*)self)->__goOn[idx] = 1; }
void* QSplashScreen_fromQObject(void* o) { return (void*)qobject_cast<QSplashScreen*>((QObject*)o); }
void* QSplashScreen_new(long long __goid) {
	return new QSplashScreen__Go(__goid);
}
void* QSplashScreen_new2(long long __goid, void* screen) {
	return new QSplashScreen__Go(__goid, (QScreen*)screen);
}
void* QSplashScreen_new3(long long __goid, void* pixmap) {
	return new QSplashScreen__Go(__goid, *(QPixmap*)pixmap);
}
void* QSplashScreen_new4(long long __goid, void* pixmap, int f) {
	return new QSplashScreen__Go(__goid, *(QPixmap*)pixmap, QFlags<Qt::WindowType>(f));
}
void* QSplashScreen_new5(long long __goid, void* screen, void* pixmap) {
	return new QSplashScreen__Go(__goid, (QScreen*)screen, *(QPixmap*)pixmap);
}
void* QSplashScreen_new6(long long __goid, void* screen, void* pixmap, int f) {
	return new QSplashScreen__Go(__goid, (QScreen*)screen, *(QPixmap*)pixmap, QFlags<Qt::WindowType>(f));
}
char* QSplashScreen_tr(const char* s, long long* __out_len) {
	auto&& __r = QSplashScreen::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QSplashScreen_setPixmap(void* self, void* pixmap) {
	((QSplashScreen*)self)->setPixmap(*(QPixmap*)pixmap);
}
void* QSplashScreen_pixmap(void* self) {
	auto&& __r = ((QSplashScreen*)self)->pixmap();
	return new QPixmap(__r);
}
void QSplashScreen_finish(void* self, void* w) {
	((QSplashScreen*)self)->finish((QWidget*)w);
}
void QSplashScreen_repaint(void* self) {
	((QSplashScreen*)self)->repaint();
}
char* QSplashScreen_message(void* self, long long* __out_len) {
	auto&& __r = ((QSplashScreen*)self)->message();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QSplashScreen_showMessage(void* self, const char* message, long long message_len) {
	QString message_qstr = QString::fromUtf8(message, message_len);
	((QSplashScreen*)self)->showMessage(message_qstr);
}
void QSplashScreen_clearMessage(void* self) {
	((QSplashScreen*)self)->clearMessage();
}
void QSplashScreen_messageChanged(void* self, const char* message, long long message_len) {
	QString message_qstr = QString::fromUtf8(message, message_len);
	((QSplashScreen*)self)->messageChanged(message_qstr);
}
void QSplashScreen_setTabOrder(void* arg1, void* arg2) {
	QSplashScreen::setTabOrder((QWidget*)arg1, (QWidget*)arg2);
}
void* QSplashScreen_mouseGrabber() {
	auto&& __r = QSplashScreen::mouseGrabber();
	return (void*)(__r);
}
void* QSplashScreen_keyboardGrabber() {
	auto&& __r = QSplashScreen::keyboardGrabber();
	return (void*)(__r);
}
void* QSplashScreen_find(unsigned long long arg1) {
	auto&& __r = QSplashScreen::find(arg1);
	return (void*)(__r);
}
void* QSplashScreen_createWindowContainer(void* window) {
	auto&& __r = QSplashScreen::createWindowContainer((QWindow*)window);
	return (void*)(__r);
}
void QSplashScreen_updateMicroFocus(void* self) {
	((QSplashScreen__Go*)self)->updateMicroFocus();
}
void QSplashScreen_create(void* self) {
	((QSplashScreen__Go*)self)->create();
}
void QSplashScreen_destroy(void* self) {
	((QSplashScreen__Go*)self)->destroy();
}
int QSplashScreen_focusNextChild(void* self) {
	auto&& __r = ((QSplashScreen__Go*)self)->focusNextChild();
	return __r;
}
int QSplashScreen_focusPreviousChild(void* self) {
	auto&& __r = ((QSplashScreen__Go*)self)->focusPreviousChild();
	return __r;
}
void* QSplashScreen_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QSplashScreen::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QSplashScreen_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QSplashScreen::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QSplashScreen_sender(void* self) {
	auto&& __r = ((QSplashScreen__Go*)self)->sender();
	return (void*)(__r);
}
int QSplashScreen_senderSignalIndex(void* self) {
	auto&& __r = ((QSplashScreen__Go*)self)->senderSignalIndex();
	return __r;
}
int QSplashScreen_receivers(void* self, const char* signal) {
	auto&& __r = ((QSplashScreen__Go*)self)->receivers(signal);
	return __r;
}
int QSplashScreen_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QSplashScreen__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
double QSplashScreen_devicePixelRatioFScale() {
	auto&& __r = QSplashScreen::devicePixelRatioFScale();
	return __r;
}
int QSplashScreen_encodeMetricF(int metric, double value) {
	auto&& __r = QSplashScreen::encodeMetricF((QPaintDevice::PaintDeviceMetric)(metric), value);
	return __r;
}
double QSplashScreen_getDecodedMetricF(void* self, int metricA, int metricB) {
	auto&& __r = ((QSplashScreen__Go*)self)->getDecodedMetricF((QPaintDevice::PaintDeviceMetric)(metricA), (QPaintDevice::PaintDeviceMetric)(metricB));
	return __r;
}
char* QSplashScreen_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QSplashScreen::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QSplashScreen_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QSplashScreen::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QSplashScreen_showMessage2(void* self, const char* message, long long message_len, int alignment) {
	QString message_qstr = QString::fromUtf8(message, message_len);
	((QSplashScreen*)self)->showMessage(message_qstr, alignment);
}
void QSplashScreen_showMessage3(void* self, const char* message, long long message_len, int alignment, void* color) {
	QString message_qstr = QString::fromUtf8(message, message_len);
	((QSplashScreen*)self)->showMessage(message_qstr, alignment, *(QColor*)color);
}
void* QSplashScreen_createWindowContainer2(void* window, void* parent) {
	auto&& __r = QSplashScreen::createWindowContainer((QWindow*)window, (QWidget*)parent);
	return (void*)(__r);
}
void* QSplashScreen_createWindowContainer3(void* window, void* parent, int flags) {
	auto&& __r = QSplashScreen::createWindowContainer((QWindow*)window, (QWidget*)parent, QFlags<Qt::WindowType>(flags));
	return (void*)(__r);
}
void QSplashScreen_updateMicroFocus2(void* self, int query) {
	((QSplashScreen__Go*)self)->updateMicroFocus((Qt::InputMethodQuery)(query));
}
void QSplashScreen_create2(void* self, unsigned long long arg1) {
	((QSplashScreen__Go*)self)->create(arg1);
}
void QSplashScreen_create3(void* self, unsigned long long arg1, int initializeWindow) {
	((QSplashScreen__Go*)self)->create(arg1, initializeWindow);
}
void QSplashScreen_create4(void* self, unsigned long long arg1, int initializeWindow, int destroyOldWindow) {
	((QSplashScreen__Go*)self)->create(arg1, initializeWindow, destroyOldWindow);
}
void QSplashScreen_destroy2(void* self, int destroyWindow) {
	((QSplashScreen__Go*)self)->destroy(destroyWindow);
}
void QSplashScreen_destroy3(void* self, int destroyWindow, int destroySubWindows) {
	((QSplashScreen__Go*)self)->destroy(destroyWindow, destroySubWindows);
}
void* QSplashScreen_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QSplashScreen::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QSplashScreen_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QSplashScreen::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QSplashScreen_metaObject_default(void* self) {
	auto&& __r = ((QSplashScreen*)self)->QSplashScreen::metaObject();
	return (void*)(__r);
}
int QSplashScreen_event_default(void* self, void* e) {
	auto&& __r = ((QSplashScreen__Go*)self)->cppgoBase_event((QEvent*)e);
	return __r;
}
void QSplashScreen_drawContents_default(void* self, void* painter) {
	((QSplashScreen__Go*)self)->cppgoBase_drawContents((QPainter*)painter);
}
void QSplashScreen_mousePressEvent_default(void* self, void* arg1) {
	((QSplashScreen__Go*)self)->cppgoBase_mousePressEvent((QMouseEvent*)arg1);
}
int QSplashScreen_devType_default(void* self) {
	auto&& __r = ((QSplashScreen*)self)->QSplashScreen::devType();
	return __r;
}
void QSplashScreen_setVisible_default(void* self, int visible) {
	((QSplashScreen*)self)->QSplashScreen::setVisible(visible);
}
void* QSplashScreen_sizeHint_default(void* self) {
	auto&& __r = ((QSplashScreen*)self)->QSplashScreen::sizeHint();
	return new QSize(__r);
}
void* QSplashScreen_minimumSizeHint_default(void* self) {
	auto&& __r = ((QSplashScreen*)self)->QSplashScreen::minimumSizeHint();
	return new QSize(__r);
}
int QSplashScreen_heightForWidth_default(void* self, int arg1) {
	auto&& __r = ((QSplashScreen*)self)->QSplashScreen::heightForWidth(arg1);
	return __r;
}
int QSplashScreen_hasHeightForWidth_default(void* self) {
	auto&& __r = ((QSplashScreen*)self)->QSplashScreen::hasHeightForWidth();
	return __r;
}
void* QSplashScreen_paintEngine_default(void* self) {
	auto&& __r = ((QSplashScreen*)self)->QSplashScreen::paintEngine();
	return (void*)(__r);
}
void QSplashScreen_mouseReleaseEvent_default(void* self, void* event) {
	((QSplashScreen__Go*)self)->cppgoBase_mouseReleaseEvent((QMouseEvent*)event);
}
void QSplashScreen_mouseDoubleClickEvent_default(void* self, void* event) {
	((QSplashScreen__Go*)self)->cppgoBase_mouseDoubleClickEvent((QMouseEvent*)event);
}
void QSplashScreen_mouseMoveEvent_default(void* self, void* event) {
	((QSplashScreen__Go*)self)->cppgoBase_mouseMoveEvent((QMouseEvent*)event);
}
void QSplashScreen_wheelEvent_default(void* self, void* event) {
	((QSplashScreen__Go*)self)->cppgoBase_wheelEvent((QWheelEvent*)event);
}
void QSplashScreen_keyPressEvent_default(void* self, void* event) {
	((QSplashScreen__Go*)self)->cppgoBase_keyPressEvent((QKeyEvent*)event);
}
void QSplashScreen_keyReleaseEvent_default(void* self, void* event) {
	((QSplashScreen__Go*)self)->cppgoBase_keyReleaseEvent((QKeyEvent*)event);
}
void QSplashScreen_focusInEvent_default(void* self, void* event) {
	((QSplashScreen__Go*)self)->cppgoBase_focusInEvent((QFocusEvent*)event);
}
void QSplashScreen_focusOutEvent_default(void* self, void* event) {
	((QSplashScreen__Go*)self)->cppgoBase_focusOutEvent((QFocusEvent*)event);
}
void QSplashScreen_enterEvent_default(void* self, void* event) {
	((QSplashScreen__Go*)self)->cppgoBase_enterEvent((QEnterEvent*)event);
}
void QSplashScreen_leaveEvent_default(void* self, void* event) {
	((QSplashScreen__Go*)self)->cppgoBase_leaveEvent((QEvent*)event);
}
void QSplashScreen_paintEvent_default(void* self, void* event) {
	((QSplashScreen__Go*)self)->cppgoBase_paintEvent((QPaintEvent*)event);
}
void QSplashScreen_moveEvent_default(void* self, void* event) {
	((QSplashScreen__Go*)self)->cppgoBase_moveEvent((QMoveEvent*)event);
}
void QSplashScreen_resizeEvent_default(void* self, void* event) {
	((QSplashScreen__Go*)self)->cppgoBase_resizeEvent((QResizeEvent*)event);
}
void QSplashScreen_closeEvent_default(void* self, void* event) {
	((QSplashScreen__Go*)self)->cppgoBase_closeEvent((QCloseEvent*)event);
}
void QSplashScreen_contextMenuEvent_default(void* self, void* event) {
	((QSplashScreen__Go*)self)->cppgoBase_contextMenuEvent((QContextMenuEvent*)event);
}
void QSplashScreen_tabletEvent_default(void* self, void* event) {
	((QSplashScreen__Go*)self)->cppgoBase_tabletEvent((QTabletEvent*)event);
}
void QSplashScreen_actionEvent_default(void* self, void* event) {
	((QSplashScreen__Go*)self)->cppgoBase_actionEvent((QActionEvent*)event);
}
void QSplashScreen_dragEnterEvent_default(void* self, void* event) {
	((QSplashScreen__Go*)self)->cppgoBase_dragEnterEvent((QDragEnterEvent*)event);
}
void QSplashScreen_dragMoveEvent_default(void* self, void* event) {
	((QSplashScreen__Go*)self)->cppgoBase_dragMoveEvent((QDragMoveEvent*)event);
}
void QSplashScreen_dragLeaveEvent_default(void* self, void* event) {
	((QSplashScreen__Go*)self)->cppgoBase_dragLeaveEvent((QDragLeaveEvent*)event);
}
void QSplashScreen_dropEvent_default(void* self, void* event) {
	((QSplashScreen__Go*)self)->cppgoBase_dropEvent((QDropEvent*)event);
}
void QSplashScreen_showEvent_default(void* self, void* event) {
	((QSplashScreen__Go*)self)->cppgoBase_showEvent((QShowEvent*)event);
}
void QSplashScreen_hideEvent_default(void* self, void* event) {
	((QSplashScreen__Go*)self)->cppgoBase_hideEvent((QHideEvent*)event);
}
void QSplashScreen_changeEvent_default(void* self, void* arg1) {
	((QSplashScreen__Go*)self)->cppgoBase_changeEvent((QEvent*)arg1);
}
int QSplashScreen_metric_default(void* self, int arg1) {
	auto&& __r = ((QSplashScreen__Go*)self)->cppgoBase_metric(arg1);
	return __r;
}
void QSplashScreen_initPainter_default(void* self, void* painter) {
	((QSplashScreen__Go*)self)->cppgoBase_initPainter((QPainter*)painter);
}
void* QSplashScreen_redirected_default(void* self, void* offset) {
	auto&& __r = ((QSplashScreen__Go*)self)->cppgoBase_redirected((QPoint*)offset);
	return (void*)(__r);
}
void* QSplashScreen_sharedPainter_default(void* self) {
	auto&& __r = ((QSplashScreen__Go*)self)->cppgoBase_sharedPainter();
	return (void*)(__r);
}
void QSplashScreen_inputMethodEvent_default(void* self, void* arg1) {
	((QSplashScreen__Go*)self)->cppgoBase_inputMethodEvent((QInputMethodEvent*)arg1);
}
void* QSplashScreen_inputMethodQuery_default(void* self, int arg1) {
	auto&& __r = ((QSplashScreen*)self)->QSplashScreen::inputMethodQuery((Qt::InputMethodQuery)(arg1));
	return new QVariant(__r);
}
int QSplashScreen_focusNextPrevChild_default(void* self, int next) {
	auto&& __r = ((QSplashScreen__Go*)self)->cppgoBase_focusNextPrevChild(next);
	return __r;
}
int QSplashScreen_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QSplashScreen*)self)->QSplashScreen::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QSplashScreen_timerEvent_default(void* self, void* event) {
	((QSplashScreen__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QSplashScreen_childEvent_default(void* self, void* event) {
	((QSplashScreen__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QSplashScreen_customEvent_default(void* self, void* event) {
	((QSplashScreen__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QSplashScreen_connectNotify_default(void* self, void* signal) {
	((QSplashScreen__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QSplashScreen_disconnectNotify_default(void* self, void* signal) {
	((QSplashScreen__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QSplashScreen_connect_messageChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSplashScreen*)self, &QSplashScreen::messageChanged, (QSplashScreen*)self, [__id](const QString& message){
		QByteArray message_ba = message.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)message_ba.constData(), message_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSplashScreen);
	__conns_QSplashScreen[__id] = __c;
}
void QSplashScreen_connect_windowTitleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSplashScreen*)self, &QSplashScreen::windowTitleChanged, (QSplashScreen*)self, [__id](const QString& title){
		QByteArray title_ba = title.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)title_ba.constData(), title_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSplashScreen);
	__conns_QSplashScreen[__id] = __c;
}
void QSplashScreen_connect_windowIconChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSplashScreen*)self, &QSplashScreen::windowIconChanged, (QSplashScreen*)self, [__id](const QIcon& icon){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&icon);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSplashScreen);
	__conns_QSplashScreen[__id] = __c;
}
void QSplashScreen_connect_windowIconTextChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSplashScreen*)self, &QSplashScreen::windowIconTextChanged, (QSplashScreen*)self, [__id](const QString& iconText){
		QByteArray iconText_ba = iconText.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)iconText_ba.constData(), iconText_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSplashScreen);
	__conns_QSplashScreen[__id] = __c;
}
void QSplashScreen_connect_customContextMenuRequested(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSplashScreen*)self, &QSplashScreen::customContextMenuRequested, (QSplashScreen*)self, [__id](const QPoint& pos){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&pos);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSplashScreen);
	__conns_QSplashScreen[__id] = __c;
}
void QSplashScreen_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSplashScreen*)self, &QSplashScreen::destroyed, (QSplashScreen*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSplashScreen);
	__conns_QSplashScreen[__id] = __c;
}
void QSplashScreen_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSplashScreen*)self, &QSplashScreen::destroyed, (QSplashScreen*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSplashScreen);
	__conns_QSplashScreen[__id] = __c;
}
void QSplashScreen_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSplashScreen*)self, &QSplashScreen::objectNameChanged, (QSplashScreen*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSplashScreen);
	__conns_QSplashScreen[__id] = __c;
}
void QSplashScreen__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QSplashScreen);
		auto __it = __conns_QSplashScreen.find(id);
		if (__it == __conns_QSplashScreen.end()) return;
		__c = __it->second;
		__conns_QSplashScreen.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QSplashScreen__Go_new(long long id) { return new QSplashScreen__Go(id); }
void QSplashScreen_delete(void* self) { delete (QSplashScreen*)self; }

}

#include <QtWidgets/qundoview.h>
#include <QtWidgets/qlistview.h>
#include <QtWidgets/qwidget.h>
#include <QtGui/qundostack.h>
#include <QtGui/qundogroup.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtGui/qicon.h>
#include <QtCore/qsize.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qrect.h>
#include <QtCore/qabstractitemmodel.h>
#include <QtWidgets/qabstractitemview.h>
#include <QtCore/qpoint.h>
#include <QtCore/qcoreevent.h>
#include <QtGui/qevent.h>
#include <QtWidgets/qstyleoption.h>
#include <QtCore/qitemselectionmodel.h>
#include <QtGui/qregion.h>
#include <QtWidgets/qabstractitemdelegate.h>
#include <QtCore/qvariant.h>
#include <QtCore/qobject.h>
#include <QtWidgets/qscrollbar.h>
#include <QtWidgets/qabstractscrollarea.h>
#include <QtCore/qmargins.h>
#include <QtWidgets/qframe.h>
#include <QtGui/qpainter.h>
#include <QtWidgets/qstyle.h>
#include <QtGui/qpalette.h>
#include <QtGui/qfont.h>
#include <QtGui/qfontmetrics.h>
#include <QtGui/qfontinfo.h>
#include <QtGui/qcursor.h>
#include <QtGui/qbitmap.h>
#include <QtGui/qpaintdevice.h>
#include <QtGui/qpixmap.h>
#include <QtWidgets/qgraphicseffect.h>
#include <QtCore/qlocale.h>
#include <QtGui/qkeysequence.h>
#include <QtWidgets/qgraphicsproxywidget.h>
#include <QtCore/qbytearray.h>
#include <QtWidgets/qsizepolicy.h>
#include <QtWidgets/qlayout.h>
#include <QtGui/qaction.h>
#include <QtGui/qpaintengine.h>
#include <QtGui/qbackingstore.h>
#include <QtGui/qwindow.h>
#include <QtGui/qscreen.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qundoview.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QUndoView__Go : public QUndoView {
public:
	long long __goId;
	unsigned char __goOn[91] = {};
	template<typename... A> QUndoView__Go(long long id, A&&... a) : QUndoView(std::forward<A>(a)...), __goId(id) {}
	using QUndoView::resizeContents;
	using QUndoView::contentsSize;
	using QUndoView::rectForIndex;
	using QUndoView::setPositionForIndex;
	using QUndoView::state;
	using QUndoView::scheduleDelayedItemsLayout;
	using QUndoView::executeDelayedItemsLayout;
	using QUndoView::setDirtyRegion;
	using QUndoView::scrollDirtyRegion;
	using QUndoView::dirtyRegionOffset;
	using QUndoView::startAutoScroll;
	using QUndoView::stopAutoScroll;
	using QUndoView::doAutoScroll;
	using QUndoView::dropIndicatorPosition;
	using QUndoView::setViewportMargins;
	using QUndoView::viewportMargins;
	using QUndoView::drawFrame;
	using QUndoView::updateMicroFocus;
	using QUndoView::create;
	using QUndoView::destroy;
	using QUndoView::focusNextChild;
	using QUndoView::focusPreviousChild;
	using QUndoView::sender;
	using QUndoView::senderSignalIndex;
	using QUndoView::receivers;
	using QUndoView::isSignalConnected;
	using QUndoView::getDecodedMetricF;
	bool cppgoBase_event(QEvent* __a0) { return QUndoView::event(__a0); }
	void cppgoBase_scrollContentsBy(int __a0, int __a1) { QUndoView::scrollContentsBy(__a0, __a1); }
	void cppgoBase_rowsInserted(const QModelIndex& __a0, int __a1, int __a2) { QUndoView::rowsInserted(__a0, __a1, __a2); }
	void cppgoBase_rowsAboutToBeRemoved(const QModelIndex& __a0, int __a1, int __a2) { QUndoView::rowsAboutToBeRemoved(__a0, __a1, __a2); }
	void cppgoBase_mouseMoveEvent(QMouseEvent* __a0) { QUndoView::mouseMoveEvent(__a0); }
	void cppgoBase_mouseReleaseEvent(QMouseEvent* __a0) { QUndoView::mouseReleaseEvent(__a0); }
	void cppgoBase_wheelEvent(QWheelEvent* __a0) { QUndoView::wheelEvent(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QUndoView::timerEvent(__a0); }
	void cppgoBase_resizeEvent(QResizeEvent* __a0) { QUndoView::resizeEvent(__a0); }
	void cppgoBase_dragMoveEvent(QDragMoveEvent* __a0) { QUndoView::dragMoveEvent(__a0); }
	void cppgoBase_dragLeaveEvent(QDragLeaveEvent* __a0) { QUndoView::dragLeaveEvent(__a0); }
	void cppgoBase_dropEvent(QDropEvent* __a0) { QUndoView::dropEvent(__a0); }
	void cppgoBase_startDrag(QFlags<Qt::DropAction> __a0) { QUndoView::startDrag(__a0); }
	void cppgoBase_initViewItemOption(QStyleOptionViewItem* __a0) { QUndoView::initViewItemOption(__a0); }
	void cppgoBase_paintEvent(QPaintEvent* __a0) { QUndoView::paintEvent(__a0); }
	int cppgoBase_horizontalOffset() { return QUndoView::horizontalOffset(); }
	int cppgoBase_verticalOffset() { return QUndoView::verticalOffset(); }
	QModelIndex cppgoBase_moveCursor(int __a0, QFlags<Qt::KeyboardModifier> __a1) { return QUndoView::moveCursor((QAbstractItemView::CursorAction)__a0, __a1); }
	void cppgoBase_setSelection(const QRect& __a0, QFlags<QItemSelectionModel::SelectionFlag> __a1) { QUndoView::setSelection(__a0, __a1); }
	QRegion cppgoBase_visualRegionForSelection(const QItemSelection& __a0) { return QUndoView::visualRegionForSelection(__a0); }
	QList<QModelIndex> cppgoBase_selectedIndexes() { return QUndoView::selectedIndexes(); }
	void cppgoBase_updateGeometries() { QUndoView::updateGeometries(); }
	bool cppgoBase_isIndexHidden(const QModelIndex& __a0) { return QUndoView::isIndexHidden(__a0); }
	void cppgoBase_selectionChanged(const QItemSelection& __a0, const QItemSelection& __a1) { QUndoView::selectionChanged(__a0, __a1); }
	void cppgoBase_currentChanged(const QModelIndex& __a0, const QModelIndex& __a1) { QUndoView::currentChanged(__a0, __a1); }
	QSize cppgoBase_viewportSizeHint() { return QUndoView::viewportSizeHint(); }
	void cppgoBase_updateEditorData() { QUndoView::updateEditorData(); }
	void cppgoBase_updateEditorGeometries() { QUndoView::updateEditorGeometries(); }
	void cppgoBase_verticalScrollbarAction(int __a0) { QUndoView::verticalScrollbarAction(__a0); }
	void cppgoBase_horizontalScrollbarAction(int __a0) { QUndoView::horizontalScrollbarAction(__a0); }
	void cppgoBase_verticalScrollbarValueChanged(int __a0) { QUndoView::verticalScrollbarValueChanged(__a0); }
	void cppgoBase_horizontalScrollbarValueChanged(int __a0) { QUndoView::horizontalScrollbarValueChanged(__a0); }
	void cppgoBase_closeEditor(QWidget* __a0, int __a1) { QUndoView::closeEditor(__a0, (QAbstractItemDelegate::EndEditHint)__a1); }
	void cppgoBase_commitData(QWidget* __a0) { QUndoView::commitData(__a0); }
	void cppgoBase_editorDestroyed(QObject* __a0) { QUndoView::editorDestroyed(__a0); }
	bool cppgoBase_edit(const QModelIndex& __a0, int __a1, QEvent* __a2) { return QUndoView::edit(__a0, (QAbstractItemView::EditTrigger)__a1, __a2); }
	QFlags<QItemSelectionModel::SelectionFlag> cppgoBase_selectionCommand(const QModelIndex& __a0, const QEvent* __a1) { return QUndoView::selectionCommand(__a0, __a1); }
	bool cppgoBase_focusNextPrevChild(bool __a0) { return QUndoView::focusNextPrevChild(__a0); }
	bool cppgoBase_viewportEvent(QEvent* __a0) { return QUndoView::viewportEvent(__a0); }
	void cppgoBase_mousePressEvent(QMouseEvent* __a0) { QUndoView::mousePressEvent(__a0); }
	void cppgoBase_mouseDoubleClickEvent(QMouseEvent* __a0) { QUndoView::mouseDoubleClickEvent(__a0); }
	void cppgoBase_dragEnterEvent(QDragEnterEvent* __a0) { QUndoView::dragEnterEvent(__a0); }
	void cppgoBase_focusInEvent(QFocusEvent* __a0) { QUndoView::focusInEvent(__a0); }
	void cppgoBase_focusOutEvent(QFocusEvent* __a0) { QUndoView::focusOutEvent(__a0); }
	void cppgoBase_keyPressEvent(QKeyEvent* __a0) { QUndoView::keyPressEvent(__a0); }
	void cppgoBase_inputMethodEvent(QInputMethodEvent* __a0) { QUndoView::inputMethodEvent(__a0); }
	bool cppgoBase_eventFilter(QObject* __a0, QEvent* __a1) { return QUndoView::eventFilter(__a0, __a1); }
	void cppgoBase_contextMenuEvent(QContextMenuEvent* __a0) { QUndoView::contextMenuEvent(__a0); }
	void cppgoBase_changeEvent(QEvent* __a0) { QUndoView::changeEvent(__a0); }
	void cppgoBase_initStyleOption(QStyleOptionFrame* __a0) { QUndoView::initStyleOption(__a0); }
	void cppgoBase_keyReleaseEvent(QKeyEvent* __a0) { QUndoView::keyReleaseEvent(__a0); }
	void cppgoBase_enterEvent(QEnterEvent* __a0) { QUndoView::enterEvent(__a0); }
	void cppgoBase_leaveEvent(QEvent* __a0) { QUndoView::leaveEvent(__a0); }
	void cppgoBase_moveEvent(QMoveEvent* __a0) { QUndoView::moveEvent(__a0); }
	void cppgoBase_closeEvent(QCloseEvent* __a0) { QUndoView::closeEvent(__a0); }
	void cppgoBase_tabletEvent(QTabletEvent* __a0) { QUndoView::tabletEvent(__a0); }
	void cppgoBase_actionEvent(QActionEvent* __a0) { QUndoView::actionEvent(__a0); }
	void cppgoBase_showEvent(QShowEvent* __a0) { QUndoView::showEvent(__a0); }
	void cppgoBase_hideEvent(QHideEvent* __a0) { QUndoView::hideEvent(__a0); }
	int cppgoBase_metric(int __a0) { return QUndoView::metric((QPaintDevice::PaintDeviceMetric)__a0); }
	void cppgoBase_initPainter(QPainter* __a0) { QUndoView::initPainter(__a0); }
	QPaintDevice* cppgoBase_redirected(QPoint* __a0) { return QUndoView::redirected(__a0); }
	QPainter* cppgoBase_sharedPainter() { return QUndoView::sharedPainter(); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QUndoView::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QUndoView::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QUndoView::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QUndoView::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QUndoView::metaObject();
	}
	QRect visualRect(const QModelIndex& index) const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, -1267696851865372997LL, (void*)&index, &__h); if (__h) return *(QRect*)__r; }
		return QUndoView::visualRect(index);
	}
	void scrollTo(const QModelIndex& index, QAbstractItemView::ScrollHint hint) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_void_voidp_int) { ((void(*)(long long, long long, void*, int, int*))cppgo_virt_void_voidp_int)(__goId, -4832340528568269947LL, (void*)&index, (int)(hint), &__h); if (__h) return; }
		QUndoView::scrollTo(index, hint);
	}
	QModelIndex indexAt(const QPoint& p) const override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, -7252823847285716802LL, (void*)&p, &__h); if (__h) return *(QModelIndex*)__r; }
		return QUndoView::indexAt(p);
	}
	void doItemsLayout() override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, -1888629738259754754LL, &__h); if (__h) return; }
		QUndoView::doItemsLayout();
	}
	void reset() override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 1086335023529244512LL, &__h); if (__h) return; }
		QUndoView::reset();
	}
	void setRootIndex(const QModelIndex& index) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1574077111192992763LL, (void*)&index, &__h); if (__h) return; }
		QUndoView::setRootIndex(index);
	}
	bool event(QEvent* e) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)e, &__h); if (__h) return __r; }
		return QUndoView::event(e);
	}
	void scrollContentsBy(int dx, int dy) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void_int_int) { ((void(*)(long long, long long, int, int, int*))cppgo_virt_void_int_int)(__goId, -6118001150126564745LL, dx, dy, &__h); if (__h) return; }
		QUndoView::scrollContentsBy(dx, dy);
	}
	void rowsInserted(const QModelIndex& parent, int start, int end) override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_void_voidp_int_int) { ((void(*)(long long, long long, void*, int, int, int*))cppgo_virt_void_voidp_int_int)(__goId, 4472434241583495620LL, (void*)&parent, start, end, &__h); if (__h) return; }
		QUndoView::rowsInserted(parent, start, end);
	}
	void rowsAboutToBeRemoved(const QModelIndex& parent, int start, int end) override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_void_voidp_int_int) { ((void(*)(long long, long long, void*, int, int, int*))cppgo_virt_void_voidp_int_int)(__goId, -8151049665126058157LL, (void*)&parent, start, end, &__h); if (__h) return; }
		QUndoView::rowsAboutToBeRemoved(parent, start, end);
	}
	void mouseMoveEvent(QMouseEvent* e) override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4255494004099806581LL, (void*)e, &__h); if (__h) return; }
		QUndoView::mouseMoveEvent(e);
	}
	void mouseReleaseEvent(QMouseEvent* e) override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4415920049297795213LL, (void*)e, &__h); if (__h) return; }
		QUndoView::mouseReleaseEvent(e);
	}
	void wheelEvent(QWheelEvent* e) override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1919988340934511316LL, (void*)e, &__h); if (__h) return; }
		QUndoView::wheelEvent(e);
	}
	void timerEvent(QTimerEvent* e) override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)e, &__h); if (__h) return; }
		QUndoView::timerEvent(e);
	}
	void resizeEvent(QResizeEvent* e) override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4302917285655014661LL, (void*)e, &__h); if (__h) return; }
		QUndoView::resizeEvent(e);
	}
	void dragMoveEvent(QDragMoveEvent* e) override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1781545137352745978LL, (void*)e, &__h); if (__h) return; }
		QUndoView::dragMoveEvent(e);
	}
	void dragLeaveEvent(QDragLeaveEvent* e) override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -3139904158373725006LL, (void*)e, &__h); if (__h) return; }
		QUndoView::dragLeaveEvent(e);
	}
	void dropEvent(QDropEvent* e) override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 9107058682754591716LL, (void*)e, &__h); if (__h) return; }
		QUndoView::dropEvent(e);
	}
	void startDrag(QFlags<Qt::DropAction> supportedActions) override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, -8112366604639385491LL, (int)(supportedActions), &__h); if (__h) return; }
		QUndoView::startDrag(supportedActions);
	}
	void initViewItemOption(QStyleOptionViewItem* option) const override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -7450787277202094556LL, (void*)option, &__h); if (__h) return; }
		QUndoView::initViewItemOption(option);
	}
	void paintEvent(QPaintEvent* e) override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5290442163167458465LL, (void*)e, &__h); if (__h) return; }
		QUndoView::paintEvent(e);
	}
	int horizontalOffset() const override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -1500310024801917802LL, &__h); if (__h) return __r; }
		return QUndoView::horizontalOffset();
	}
	int verticalOffset() const override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, 5218707810677921016LL, &__h); if (__h) return __r; }
		return QUndoView::verticalOffset();
	}
	QModelIndex moveCursor(QAbstractItemView::CursorAction cursorAction, QFlags<Qt::KeyboardModifier> modifiers) override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_voidp_int_int) { void* __r = ((void*(*)(long long, long long, int, int, int*))cppgo_virt_voidp_int_int)(__goId, -158523882194220890LL, (int)(cursorAction), (int)(modifiers), &__h); if (__h) return *(QModelIndex*)__r; }
		return QUndoView::moveCursor(cursorAction, modifiers);
	}
	void setSelection(const QRect& rect, QFlags<QItemSelectionModel::SelectionFlag> command) override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_void_voidp_int) { ((void(*)(long long, long long, void*, int, int*))cppgo_virt_void_voidp_int)(__goId, -1539266838375037305LL, (void*)&rect, (int)(command), &__h); if (__h) return; }
		QUndoView::setSelection(rect, command);
	}
	QRegion visualRegionForSelection(const QItemSelection& selection) const override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 7063795720810552074LL, (void*)&selection, &__h); if (__h) return *(QRegion*)__r; }
		return QUndoView::visualRegionForSelection(selection);
	}
	QList<QModelIndex> selectedIndexes() const override {
		int __h = 0;
		if (__goOn[27] && cppgo_virtp_voidp) {
			long long __len = 0;
			void* __a = ((void*(*)(long long, long long, long long*, int*))cppgo_virtp_voidp)(__goId, 740869982805721908LL, &__len, &__h);
			if (__h) {
				QList<QModelIndex> __r; __r.reserve(__len);
				void** __aa = (void**)__a;
				for (long long __i = 0; __i < __len; __i++) __r.append(*(QModelIndex*)__aa[__i]);
				if (__a) free(__a);
				return __r;
			}
		}
		return QUndoView::selectedIndexes();
	}
	void updateGeometries() override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 1022771470898460576LL, &__h); if (__h) return; }
		QUndoView::updateGeometries();
	}
	bool isIndexHidden(const QModelIndex& index) const override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 6057989667089378859LL, (void*)&index, &__h); if (__h) return __r; }
		return QUndoView::isIndexHidden(index);
	}
	void selectionChanged(const QItemSelection& selected, const QItemSelection& deselected) override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_void_voidp_voidp) { ((void(*)(long long, long long, void*, void*, int*))cppgo_virt_void_voidp_voidp)(__goId, -5892205462145366901LL, (void*)&selected, (void*)&deselected, &__h); if (__h) return; }
		QUndoView::selectionChanged(selected, deselected);
	}
	void currentChanged(const QModelIndex& current, const QModelIndex& previous) override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_void_voidp_voidp) { ((void(*)(long long, long long, void*, void*, int*))cppgo_virt_void_voidp_voidp)(__goId, -6978583972112316034LL, (void*)&current, (void*)&previous, &__h); if (__h) return; }
		QUndoView::currentChanged(current, previous);
	}
	QSize viewportSizeHint() const override {
		int __h = 0;
		if (__goOn[32] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 503579067252048239LL, &__h); if (__h) return *(QSize*)__r; }
		return QUndoView::viewportSizeHint();
	}
	void setModel(QAbstractItemModel* model) override {
		int __h = 0;
		if (__goOn[33] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4283922507359547360LL, (void*)model, &__h); if (__h) return; }
		QUndoView::setModel(model);
	}
	void setSelectionModel(QItemSelectionModel* selectionModel) override {
		int __h = 0;
		if (__goOn[34] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4689322347236366780LL, (void*)selectionModel, &__h); if (__h) return; }
		QUndoView::setSelectionModel(selectionModel);
	}
	void keyboardSearch(const QString& search) override {
		int __h = 0;
		QByteArray search_ba = search.toUtf8();
		if (__goOn[35] && cppgo_virt_void_constcharp_longlong) { ((void(*)(long long, long long, const char*, long long, int*))cppgo_virt_void_constcharp_longlong)(__goId, -5041452885022114200LL, (char*)search_ba.constData(), search_ba.size(), &__h); if (__h) return; }
		QUndoView::keyboardSearch(search);
	}
	int sizeHintForRow(int row) const override {
		int __h = 0;
		if (__goOn[36] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, -4084029023708556876LL, row, &__h); if (__h) return __r; }
		return QUndoView::sizeHintForRow(row);
	}
	int sizeHintForColumn(int column) const override {
		int __h = 0;
		if (__goOn[37] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, -3666792239577137638LL, column, &__h); if (__h) return __r; }
		return QUndoView::sizeHintForColumn(column);
	}
	QAbstractItemDelegate* itemDelegateForIndex(const QModelIndex& index) const override {
		int __h = 0;
		if (__goOn[38] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, -3811438228839363824LL, (void*)&index, &__h); if (__h) return (QAbstractItemDelegate*)__r; }
		return QUndoView::itemDelegateForIndex(index);
	}
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const override {
		int __h = 0;
		if (__goOn[39] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7345853097340709638LL, (int)(query), &__h); if (__h) return *(QVariant*)__r; }
		return QUndoView::inputMethodQuery(query);
	}
	void selectAll() override {
		int __h = 0;
		if (__goOn[40] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 8758113809152574300LL, &__h); if (__h) return; }
		QUndoView::selectAll();
	}
	void updateEditorData() override {
		int __h = 0;
		if (__goOn[41] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 9029941737538200621LL, &__h); if (__h) return; }
		QUndoView::updateEditorData();
	}
	void updateEditorGeometries() override {
		int __h = 0;
		if (__goOn[42] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, -8726300139121449691LL, &__h); if (__h) return; }
		QUndoView::updateEditorGeometries();
	}
	void verticalScrollbarAction(int action) override {
		int __h = 0;
		if (__goOn[43] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, -1155555122391250623LL, action, &__h); if (__h) return; }
		QUndoView::verticalScrollbarAction(action);
	}
	void horizontalScrollbarAction(int action) override {
		int __h = 0;
		if (__goOn[44] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, -785551405282638793LL, action, &__h); if (__h) return; }
		QUndoView::horizontalScrollbarAction(action);
	}
	void verticalScrollbarValueChanged(int value) override {
		int __h = 0;
		if (__goOn[45] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 6366953527806415524LL, value, &__h); if (__h) return; }
		QUndoView::verticalScrollbarValueChanged(value);
	}
	void horizontalScrollbarValueChanged(int value) override {
		int __h = 0;
		if (__goOn[46] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 1786430177580376614LL, value, &__h); if (__h) return; }
		QUndoView::horizontalScrollbarValueChanged(value);
	}
	void closeEditor(QWidget* editor, QAbstractItemDelegate::EndEditHint hint) override {
		int __h = 0;
		if (__goOn[47] && cppgo_virt_void_voidp_int) { ((void(*)(long long, long long, void*, int, int*))cppgo_virt_void_voidp_int)(__goId, 8648484264628816806LL, (void*)editor, (int)(hint), &__h); if (__h) return; }
		QUndoView::closeEditor(editor, hint);
	}
	void commitData(QWidget* editor) override {
		int __h = 0;
		if (__goOn[48] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6021383853323633948LL, (void*)editor, &__h); if (__h) return; }
		QUndoView::commitData(editor);
	}
	void editorDestroyed(QObject* editor) override {
		int __h = 0;
		if (__goOn[49] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6979418566323745651LL, (void*)editor, &__h); if (__h) return; }
		QUndoView::editorDestroyed(editor);
	}
	bool edit(const QModelIndex& index, QAbstractItemView::EditTrigger trigger, QEvent* event) override {
		int __h = 0;
		if (__goOn[50] && cppgo_virt_int_voidp_int_voidp) { int __r = ((int(*)(long long, long long, void*, int, void*, int*))cppgo_virt_int_voidp_int_voidp)(__goId, 4206689019978932209LL, (void*)&index, (int)(trigger), (void*)event, &__h); if (__h) return __r; }
		return QUndoView::edit(index, trigger, event);
	}
	QFlags<QItemSelectionModel::SelectionFlag> selectionCommand(const QModelIndex& index, const QEvent* event) const override {
		int __h = 0;
		if (__goOn[51] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, -2447490284072398302LL, (void*)&index, (void*)event, &__h); if (__h) return QFlags<QItemSelectionModel::SelectionFlag>(__r); }
		return QUndoView::selectionCommand(index, event);
	}
	bool focusNextPrevChild(bool next) override {
		int __h = 0;
		if (__goOn[52] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 5898328517505784191LL, next, &__h); if (__h) return __r; }
		return QUndoView::focusNextPrevChild(next);
	}
	bool viewportEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[53] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, -1971465653061861519LL, (void*)event, &__h); if (__h) return __r; }
		return QUndoView::viewportEvent(event);
	}
	void mousePressEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[54] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -332090898078543319LL, (void*)event, &__h); if (__h) return; }
		QUndoView::mousePressEvent(event);
	}
	void mouseDoubleClickEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[55] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4029073905697023887LL, (void*)event, &__h); if (__h) return; }
		QUndoView::mouseDoubleClickEvent(event);
	}
	void dragEnterEvent(QDragEnterEvent* event) override {
		int __h = 0;
		if (__goOn[56] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 206122715919917589LL, (void*)event, &__h); if (__h) return; }
		QUndoView::dragEnterEvent(event);
	}
	void focusInEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[57] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 975043767033552552LL, (void*)event, &__h); if (__h) return; }
		QUndoView::focusInEvent(event);
	}
	void focusOutEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[58] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -385596010576048993LL, (void*)event, &__h); if (__h) return; }
		QUndoView::focusOutEvent(event);
	}
	void keyPressEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[59] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1023791977679793185LL, (void*)event, &__h); if (__h) return; }
		QUndoView::keyPressEvent(event);
	}
	void inputMethodEvent(QInputMethodEvent* event) override {
		int __h = 0;
		if (__goOn[60] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5438819661942843222LL, (void*)event, &__h); if (__h) return; }
		QUndoView::inputMethodEvent(event);
	}
	bool eventFilter(QObject* object, QEvent* event) override {
		int __h = 0;
		if (__goOn[61] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)object, (void*)event, &__h); if (__h) return __r; }
		return QUndoView::eventFilter(object, event);
	}
	QSize minimumSizeHint() const override {
		int __h = 0;
		if (__goOn[62] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 7049264599837532491LL, &__h); if (__h) return *(QSize*)__r; }
		return QUndoView::minimumSizeHint();
	}
	QSize sizeHint() const override {
		int __h = 0;
		if (__goOn[63] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 729690261457684101LL, &__h); if (__h) return *(QSize*)__r; }
		return QUndoView::sizeHint();
	}
	void setupViewport(QWidget* viewport) override {
		int __h = 0;
		if (__goOn[64] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4531996381736730688LL, (void*)viewport, &__h); if (__h) return; }
		QUndoView::setupViewport(viewport);
	}
	void contextMenuEvent(QContextMenuEvent* arg1) override {
		int __h = 0;
		if (__goOn[65] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6698043222227006015LL, (void*)arg1, &__h); if (__h) return; }
		QUndoView::contextMenuEvent(arg1);
	}
	void changeEvent(QEvent* arg1) override {
		int __h = 0;
		if (__goOn[66] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4959376063826165943LL, (void*)arg1, &__h); if (__h) return; }
		QUndoView::changeEvent(arg1);
	}
	void initStyleOption(QStyleOptionFrame* option) const override {
		int __h = 0;
		if (__goOn[67] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -488693769771308531LL, (void*)option, &__h); if (__h) return; }
		QUndoView::initStyleOption(option);
	}
	int devType() const override {
		int __h = 0;
		if (__goOn[68] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8659679893868524044LL, &__h); if (__h) return __r; }
		return QUndoView::devType();
	}
	void setVisible(bool visible) override {
		int __h = 0;
		if (__goOn[69] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 2110182719375245575LL, visible, &__h); if (__h) return; }
		QUndoView::setVisible(visible);
	}
	int heightForWidth(int arg1) const override {
		int __h = 0;
		if (__goOn[70] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 4311534060495737339LL, arg1, &__h); if (__h) return __r; }
		return QUndoView::heightForWidth(arg1);
	}
	bool hasHeightForWidth() const override {
		int __h = 0;
		if (__goOn[71] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8266762397995423915LL, &__h); if (__h) return __r; }
		return QUndoView::hasHeightForWidth();
	}
	QPaintEngine* paintEngine() const override {
		int __h = 0;
		if (__goOn[72] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 6398007404756346381LL, &__h); if (__h) return (QPaintEngine*)__r; }
		return QUndoView::paintEngine();
	}
	void keyReleaseEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[73] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 772598475950050299LL, (void*)event, &__h); if (__h) return; }
		QUndoView::keyReleaseEvent(event);
	}
	void enterEvent(QEnterEvent* event) override {
		int __h = 0;
		if (__goOn[74] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6316270346668237895LL, (void*)event, &__h); if (__h) return; }
		QUndoView::enterEvent(event);
	}
	void leaveEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[75] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -551070009623711640LL, (void*)event, &__h); if (__h) return; }
		QUndoView::leaveEvent(event);
	}
	void moveEvent(QMoveEvent* event) override {
		int __h = 0;
		if (__goOn[76] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -9161949150489525396LL, (void*)event, &__h); if (__h) return; }
		QUndoView::moveEvent(event);
	}
	void closeEvent(QCloseEvent* event) override {
		int __h = 0;
		if (__goOn[77] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1989937050019464783LL, (void*)event, &__h); if (__h) return; }
		QUndoView::closeEvent(event);
	}
	void tabletEvent(QTabletEvent* event) override {
		int __h = 0;
		if (__goOn[78] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -212068866860701141LL, (void*)event, &__h); if (__h) return; }
		QUndoView::tabletEvent(event);
	}
	void actionEvent(QActionEvent* event) override {
		int __h = 0;
		if (__goOn[79] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1697444140088555581LL, (void*)event, &__h); if (__h) return; }
		QUndoView::actionEvent(event);
	}
	void showEvent(QShowEvent* event) override {
		int __h = 0;
		if (__goOn[80] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6825369174751916LL, (void*)event, &__h); if (__h) return; }
		QUndoView::showEvent(event);
	}
	void hideEvent(QHideEvent* event) override {
		int __h = 0;
		if (__goOn[81] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -199718707876719129LL, (void*)event, &__h); if (__h) return; }
		QUndoView::hideEvent(event);
	}
	bool nativeEvent(const QByteArray& eventType, void* message, qintptr* result) override {
		int __h = 0;
		if (__goOn[82] && cppgo_virt_int_voidp_voidp_longlongp) { int __r = ((int(*)(long long, long long, void*, void*, long long*, int*))cppgo_virt_int_voidp_voidp_longlongp)(__goId, -2039166586246443706LL, (void*)&eventType, (void*)message, (long long*)result, &__h); if (__h) return __r; }
		return QUndoView::nativeEvent(eventType, message, result);
	}
	int metric(QPaintDevice::PaintDeviceMetric arg1) const override {
		int __h = 0;
		if (__goOn[83] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 427629649267311777LL, (int)(arg1), &__h); if (__h) return __r; }
		return QUndoView::metric(arg1);
	}
	void initPainter(QPainter* painter) const override {
		int __h = 0;
		if (__goOn[84] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6477500296980050210LL, (void*)painter, &__h); if (__h) return; }
		QUndoView::initPainter(painter);
	}
	QPaintDevice* redirected(QPoint* offset) const override {
		int __h = 0;
		if (__goOn[85] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 4127240633310655048LL, (void*)offset, &__h); if (__h) return (QPaintDevice*)__r; }
		return QUndoView::redirected(offset);
	}
	QPainter* sharedPainter() const override {
		int __h = 0;
		if (__goOn[86] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 5792771328741774991LL, &__h); if (__h) return (QPainter*)__r; }
		return QUndoView::sharedPainter();
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[87] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QUndoView::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[88] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QUndoView::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[89] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QUndoView::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[90] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QUndoView::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QUndoView;
static std::mutex __connsMu_QUndoView;

extern "C" {

void* QUndoView_base(void* self) { return (void*)static_cast<QListView*>((QUndoView*)self); }
void QUndoView_chain(void* self, void** out) {
	QUndoView* __p = (QUndoView*)self;
	QListView* __b0 = static_cast<QListView*>(__p); out[0] = (void*)__b0;
	QAbstractItemView* __b1 = static_cast<QAbstractItemView*>(__b0); out[1] = (void*)__b1;
	QAbstractScrollArea* __b2 = static_cast<QAbstractScrollArea*>(__b1); out[2] = (void*)__b2;
	QFrame* __b3 = static_cast<QFrame*>(__b2); out[3] = (void*)__b3;
	QWidget* __b4 = static_cast<QWidget*>(__b3); out[4] = (void*)__b4;
	QObject* __b5 = static_cast<QObject*>(__b4); out[5] = (void*)__b5;
}
void QUndoView__set_override(void* self, int idx) { ((QUndoView__Go*)self)->__goOn[idx] = 1; }
void* QUndoView_fromQObject(void* o) { return (void*)qobject_cast<QUndoView*>((QObject*)o); }
void* QUndoView_new(long long __goid, void* parent) {
	return new QUndoView__Go(__goid, (QWidget*)parent);
}
void* QUndoView_new2(long long __goid) {
	return new QUndoView__Go(__goid);
}
void* QUndoView_new3(long long __goid, void* stack) {
	return new QUndoView__Go(__goid, (QUndoStack*)stack);
}
void* QUndoView_new4(long long __goid, void* group) {
	return new QUndoView__Go(__goid, (QUndoGroup*)group);
}
void* QUndoView_new5(long long __goid, void* stack, void* parent) {
	return new QUndoView__Go(__goid, (QUndoStack*)stack, (QWidget*)parent);
}
void* QUndoView_new6(long long __goid, void* group, void* parent) {
	return new QUndoView__Go(__goid, (QUndoGroup*)group, (QWidget*)parent);
}
char* QUndoView_tr(const char* s, long long* __out_len) {
	auto&& __r = QUndoView::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QUndoView_stack(void* self) {
	auto&& __r = ((QUndoView*)self)->stack();
	return (void*)(__r);
}
void* QUndoView_group(void* self) {
	auto&& __r = ((QUndoView*)self)->group();
	return (void*)(__r);
}
void QUndoView_setEmptyLabel(void* self, const char* label, long long label_len) {
	QString label_qstr = QString::fromUtf8(label, label_len);
	((QUndoView*)self)->setEmptyLabel(label_qstr);
}
char* QUndoView_emptyLabel(void* self, long long* __out_len) {
	auto&& __r = ((QUndoView*)self)->emptyLabel();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QUndoView_setCleanIcon(void* self, void* icon) {
	((QUndoView*)self)->setCleanIcon(*(QIcon*)icon);
}
void* QUndoView_cleanIcon(void* self) {
	auto&& __r = ((QUndoView*)self)->cleanIcon();
	return new QIcon(__r);
}
void QUndoView_setStack(void* self, void* stack) {
	((QUndoView*)self)->setStack((QUndoStack*)stack);
}
void QUndoView_setGroup(void* self, void* group) {
	((QUndoView*)self)->setGroup((QUndoGroup*)group);
}
void QUndoView_resizeContents(void* self, int width, int height) {
	((QUndoView__Go*)self)->resizeContents(width, height);
}
void* QUndoView_contentsSize(void* self) {
	auto&& __r = ((QUndoView__Go*)self)->contentsSize();
	return new QSize(__r);
}
void* QUndoView_rectForIndex(void* self, void* index) {
	auto&& __r = ((QUndoView__Go*)self)->rectForIndex(*(QModelIndex*)index);
	return new QRect(__r);
}
void QUndoView_setPositionForIndex(void* self, void* position, void* index) {
	((QUndoView__Go*)self)->setPositionForIndex(*(QPoint*)position, *(QModelIndex*)index);
}
int QUndoView_state(void* self) {
	auto&& __r = ((QUndoView__Go*)self)->state();
	return (int)(__r);
}
void QUndoView_scheduleDelayedItemsLayout(void* self) {
	((QUndoView__Go*)self)->scheduleDelayedItemsLayout();
}
void QUndoView_executeDelayedItemsLayout(void* self) {
	((QUndoView__Go*)self)->executeDelayedItemsLayout();
}
void QUndoView_setDirtyRegion(void* self, void* region) {
	((QUndoView__Go*)self)->setDirtyRegion(*(QRegion*)region);
}
void QUndoView_scrollDirtyRegion(void* self, int dx, int dy) {
	((QUndoView__Go*)self)->scrollDirtyRegion(dx, dy);
}
void* QUndoView_dirtyRegionOffset(void* self) {
	auto&& __r = ((QUndoView__Go*)self)->dirtyRegionOffset();
	return new QPoint(__r);
}
void QUndoView_startAutoScroll(void* self) {
	((QUndoView__Go*)self)->startAutoScroll();
}
void QUndoView_stopAutoScroll(void* self) {
	((QUndoView__Go*)self)->stopAutoScroll();
}
void QUndoView_doAutoScroll(void* self) {
	((QUndoView__Go*)self)->doAutoScroll();
}
int QUndoView_dropIndicatorPosition(void* self) {
	auto&& __r = ((QUndoView__Go*)self)->dropIndicatorPosition();
	return (int)(__r);
}
void QUndoView_setViewportMargins(void* self, int left, int top, int right, int bottom) {
	((QUndoView__Go*)self)->setViewportMargins(left, top, right, bottom);
}
void QUndoView_setViewportMargins2(void* self, void* margins) {
	((QUndoView__Go*)self)->setViewportMargins(*(QMargins*)margins);
}
void* QUndoView_viewportMargins(void* self) {
	auto&& __r = ((QUndoView__Go*)self)->viewportMargins();
	return new QMargins(__r);
}
void QUndoView_drawFrame(void* self, void* arg1) {
	((QUndoView__Go*)self)->drawFrame((QPainter*)arg1);
}
void QUndoView_setTabOrder(void* arg1, void* arg2) {
	QUndoView::setTabOrder((QWidget*)arg1, (QWidget*)arg2);
}
void* QUndoView_mouseGrabber() {
	auto&& __r = QUndoView::mouseGrabber();
	return (void*)(__r);
}
void* QUndoView_keyboardGrabber() {
	auto&& __r = QUndoView::keyboardGrabber();
	return (void*)(__r);
}
void* QUndoView_find(unsigned long long arg1) {
	auto&& __r = QUndoView::find(arg1);
	return (void*)(__r);
}
void* QUndoView_createWindowContainer(void* window) {
	auto&& __r = QUndoView::createWindowContainer((QWindow*)window);
	return (void*)(__r);
}
void QUndoView_updateMicroFocus(void* self) {
	((QUndoView__Go*)self)->updateMicroFocus();
}
void QUndoView_create(void* self) {
	((QUndoView__Go*)self)->create();
}
void QUndoView_destroy(void* self) {
	((QUndoView__Go*)self)->destroy();
}
int QUndoView_focusNextChild(void* self) {
	auto&& __r = ((QUndoView__Go*)self)->focusNextChild();
	return __r;
}
int QUndoView_focusPreviousChild(void* self) {
	auto&& __r = ((QUndoView__Go*)self)->focusPreviousChild();
	return __r;
}
void* QUndoView_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QUndoView::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QUndoView_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QUndoView::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QUndoView_sender(void* self) {
	auto&& __r = ((QUndoView__Go*)self)->sender();
	return (void*)(__r);
}
int QUndoView_senderSignalIndex(void* self) {
	auto&& __r = ((QUndoView__Go*)self)->senderSignalIndex();
	return __r;
}
int QUndoView_receivers(void* self, const char* signal) {
	auto&& __r = ((QUndoView__Go*)self)->receivers(signal);
	return __r;
}
int QUndoView_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QUndoView__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
double QUndoView_devicePixelRatioFScale() {
	auto&& __r = QUndoView::devicePixelRatioFScale();
	return __r;
}
int QUndoView_encodeMetricF(int metric, double value) {
	auto&& __r = QUndoView::encodeMetricF((QPaintDevice::PaintDeviceMetric)(metric), value);
	return __r;
}
double QUndoView_getDecodedMetricF(void* self, int metricA, int metricB) {
	auto&& __r = ((QUndoView__Go*)self)->getDecodedMetricF((QPaintDevice::PaintDeviceMetric)(metricA), (QPaintDevice::PaintDeviceMetric)(metricB));
	return __r;
}
char* QUndoView_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QUndoView::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QUndoView_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QUndoView::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QUndoView_createWindowContainer2(void* window, void* parent) {
	auto&& __r = QUndoView::createWindowContainer((QWindow*)window, (QWidget*)parent);
	return (void*)(__r);
}
void* QUndoView_createWindowContainer3(void* window, void* parent, int flags) {
	auto&& __r = QUndoView::createWindowContainer((QWindow*)window, (QWidget*)parent, QFlags<Qt::WindowType>(flags));
	return (void*)(__r);
}
void QUndoView_updateMicroFocus2(void* self, int query) {
	((QUndoView__Go*)self)->updateMicroFocus((Qt::InputMethodQuery)(query));
}
void QUndoView_create2(void* self, unsigned long long arg1) {
	((QUndoView__Go*)self)->create(arg1);
}
void QUndoView_create3(void* self, unsigned long long arg1, int initializeWindow) {
	((QUndoView__Go*)self)->create(arg1, initializeWindow);
}
void QUndoView_create4(void* self, unsigned long long arg1, int initializeWindow, int destroyOldWindow) {
	((QUndoView__Go*)self)->create(arg1, initializeWindow, destroyOldWindow);
}
void QUndoView_destroy2(void* self, int destroyWindow) {
	((QUndoView__Go*)self)->destroy(destroyWindow);
}
void QUndoView_destroy3(void* self, int destroyWindow, int destroySubWindows) {
	((QUndoView__Go*)self)->destroy(destroyWindow, destroySubWindows);
}
void* QUndoView_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QUndoView::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QUndoView_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QUndoView::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QUndoView_metaObject_default(void* self) {
	auto&& __r = ((QUndoView*)self)->QUndoView::metaObject();
	return (void*)(__r);
}
void* QUndoView_visualRect_default(void* self, void* index) {
	auto&& __r = ((QUndoView*)self)->QUndoView::visualRect(*(QModelIndex*)index);
	return new QRect(__r);
}
void QUndoView_scrollTo_default(void* self, void* index, int hint) {
	((QUndoView*)self)->QUndoView::scrollTo(*(QModelIndex*)index, (QAbstractItemView::ScrollHint)(hint));
}
void* QUndoView_indexAt_default(void* self, void* p) {
	auto&& __r = ((QUndoView*)self)->QUndoView::indexAt(*(QPoint*)p);
	return new QModelIndex(__r);
}
void QUndoView_doItemsLayout_default(void* self) {
	((QUndoView*)self)->QUndoView::doItemsLayout();
}
void QUndoView_reset_default(void* self) {
	((QUndoView*)self)->QUndoView::reset();
}
void QUndoView_setRootIndex_default(void* self, void* index) {
	((QUndoView*)self)->QUndoView::setRootIndex(*(QModelIndex*)index);
}
int QUndoView_event_default(void* self, void* e) {
	auto&& __r = ((QUndoView__Go*)self)->cppgoBase_event((QEvent*)e);
	return __r;
}
void QUndoView_scrollContentsBy_default(void* self, int dx, int dy) {
	((QUndoView__Go*)self)->cppgoBase_scrollContentsBy(dx, dy);
}
void QUndoView_rowsInserted_default(void* self, void* parent, int start, int end) {
	((QUndoView__Go*)self)->cppgoBase_rowsInserted(*(QModelIndex*)parent, start, end);
}
void QUndoView_rowsAboutToBeRemoved_default(void* self, void* parent, int start, int end) {
	((QUndoView__Go*)self)->cppgoBase_rowsAboutToBeRemoved(*(QModelIndex*)parent, start, end);
}
void QUndoView_mouseMoveEvent_default(void* self, void* e) {
	((QUndoView__Go*)self)->cppgoBase_mouseMoveEvent((QMouseEvent*)e);
}
void QUndoView_mouseReleaseEvent_default(void* self, void* e) {
	((QUndoView__Go*)self)->cppgoBase_mouseReleaseEvent((QMouseEvent*)e);
}
void QUndoView_wheelEvent_default(void* self, void* e) {
	((QUndoView__Go*)self)->cppgoBase_wheelEvent((QWheelEvent*)e);
}
void QUndoView_timerEvent_default(void* self, void* e) {
	((QUndoView__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)e);
}
void QUndoView_resizeEvent_default(void* self, void* e) {
	((QUndoView__Go*)self)->cppgoBase_resizeEvent((QResizeEvent*)e);
}
void QUndoView_dragMoveEvent_default(void* self, void* e) {
	((QUndoView__Go*)self)->cppgoBase_dragMoveEvent((QDragMoveEvent*)e);
}
void QUndoView_dragLeaveEvent_default(void* self, void* e) {
	((QUndoView__Go*)self)->cppgoBase_dragLeaveEvent((QDragLeaveEvent*)e);
}
void QUndoView_dropEvent_default(void* self, void* e) {
	((QUndoView__Go*)self)->cppgoBase_dropEvent((QDropEvent*)e);
}
void QUndoView_startDrag_default(void* self, int supportedActions) {
	((QUndoView__Go*)self)->cppgoBase_startDrag(QFlags<Qt::DropAction>(supportedActions));
}
void QUndoView_initViewItemOption_default(void* self, void* option) {
	((QUndoView__Go*)self)->cppgoBase_initViewItemOption((QStyleOptionViewItem*)option);
}
void QUndoView_paintEvent_default(void* self, void* e) {
	((QUndoView__Go*)self)->cppgoBase_paintEvent((QPaintEvent*)e);
}
int QUndoView_horizontalOffset_default(void* self) {
	auto&& __r = ((QUndoView__Go*)self)->cppgoBase_horizontalOffset();
	return __r;
}
int QUndoView_verticalOffset_default(void* self) {
	auto&& __r = ((QUndoView__Go*)self)->cppgoBase_verticalOffset();
	return __r;
}
void* QUndoView_moveCursor_default(void* self, int cursorAction, int modifiers) {
	auto&& __r = ((QUndoView__Go*)self)->cppgoBase_moveCursor(cursorAction, QFlags<Qt::KeyboardModifier>(modifiers));
	return new QModelIndex(__r);
}
void QUndoView_setSelection_default(void* self, void* rect, int command) {
	((QUndoView__Go*)self)->cppgoBase_setSelection(*(QRect*)rect, QFlags<QItemSelectionModel::SelectionFlag>(command));
}
void* QUndoView_visualRegionForSelection_default(void* self, void* selection) {
	auto&& __r = ((QUndoView__Go*)self)->cppgoBase_visualRegionForSelection(*(QItemSelection*)selection);
	return new QRegion(__r);
}
void** QUndoView_selectedIndexes_default(void* self, long long* __out_len) {
	auto&& __r = ((QUndoView__Go*)self)->cppgoBase_selectedIndexes();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QModelIndex(__e);
	return __arr;
}
void QUndoView_updateGeometries_default(void* self) {
	((QUndoView__Go*)self)->cppgoBase_updateGeometries();
}
int QUndoView_isIndexHidden_default(void* self, void* index) {
	auto&& __r = ((QUndoView__Go*)self)->cppgoBase_isIndexHidden(*(QModelIndex*)index);
	return __r;
}
void QUndoView_selectionChanged_default(void* self, void* selected, void* deselected) {
	((QUndoView__Go*)self)->cppgoBase_selectionChanged(*(QItemSelection*)selected, *(QItemSelection*)deselected);
}
void QUndoView_currentChanged_default(void* self, void* current, void* previous) {
	((QUndoView__Go*)self)->cppgoBase_currentChanged(*(QModelIndex*)current, *(QModelIndex*)previous);
}
void* QUndoView_viewportSizeHint_default(void* self) {
	auto&& __r = ((QUndoView__Go*)self)->cppgoBase_viewportSizeHint();
	return new QSize(__r);
}
void QUndoView_setModel_default(void* self, void* model) {
	((QUndoView*)self)->QUndoView::setModel((QAbstractItemModel*)model);
}
void QUndoView_setSelectionModel_default(void* self, void* selectionModel) {
	((QUndoView*)self)->QUndoView::setSelectionModel((QItemSelectionModel*)selectionModel);
}
void QUndoView_keyboardSearch_default(void* self, const char* search, long long search_len) {
	QString search_qstr = QString::fromUtf8(search, search_len);
	((QUndoView*)self)->QUndoView::keyboardSearch(search_qstr);
}
int QUndoView_sizeHintForRow_default(void* self, int row) {
	auto&& __r = ((QUndoView*)self)->QUndoView::sizeHintForRow(row);
	return __r;
}
int QUndoView_sizeHintForColumn_default(void* self, int column) {
	auto&& __r = ((QUndoView*)self)->QUndoView::sizeHintForColumn(column);
	return __r;
}
void* QUndoView_itemDelegateForIndex_default(void* self, void* index) {
	auto&& __r = ((QUndoView*)self)->QUndoView::itemDelegateForIndex(*(QModelIndex*)index);
	return (void*)(__r);
}
void* QUndoView_inputMethodQuery_default(void* self, int query) {
	auto&& __r = ((QUndoView*)self)->QUndoView::inputMethodQuery((Qt::InputMethodQuery)(query));
	return new QVariant(__r);
}
void QUndoView_selectAll_default(void* self) {
	((QUndoView*)self)->QUndoView::selectAll();
}
void QUndoView_updateEditorData_default(void* self) {
	((QUndoView__Go*)self)->cppgoBase_updateEditorData();
}
void QUndoView_updateEditorGeometries_default(void* self) {
	((QUndoView__Go*)self)->cppgoBase_updateEditorGeometries();
}
void QUndoView_verticalScrollbarAction_default(void* self, int action) {
	((QUndoView__Go*)self)->cppgoBase_verticalScrollbarAction(action);
}
void QUndoView_horizontalScrollbarAction_default(void* self, int action) {
	((QUndoView__Go*)self)->cppgoBase_horizontalScrollbarAction(action);
}
void QUndoView_verticalScrollbarValueChanged_default(void* self, int value) {
	((QUndoView__Go*)self)->cppgoBase_verticalScrollbarValueChanged(value);
}
void QUndoView_horizontalScrollbarValueChanged_default(void* self, int value) {
	((QUndoView__Go*)self)->cppgoBase_horizontalScrollbarValueChanged(value);
}
void QUndoView_closeEditor_default(void* self, void* editor, int hint) {
	((QUndoView__Go*)self)->cppgoBase_closeEditor((QWidget*)editor, hint);
}
void QUndoView_commitData_default(void* self, void* editor) {
	((QUndoView__Go*)self)->cppgoBase_commitData((QWidget*)editor);
}
void QUndoView_editorDestroyed_default(void* self, void* editor) {
	((QUndoView__Go*)self)->cppgoBase_editorDestroyed((QObject*)editor);
}
int QUndoView_edit_default(void* self, void* index, int trigger, void* event) {
	auto&& __r = ((QUndoView__Go*)self)->cppgoBase_edit(*(QModelIndex*)index, trigger, (QEvent*)event);
	return __r;
}
int QUndoView_selectionCommand_default(void* self, void* index, void* event) {
	auto&& __r = ((QUndoView__Go*)self)->cppgoBase_selectionCommand(*(QModelIndex*)index, (QEvent*)event);
	return (int)(__r);
}
int QUndoView_focusNextPrevChild_default(void* self, int next) {
	auto&& __r = ((QUndoView__Go*)self)->cppgoBase_focusNextPrevChild(next);
	return __r;
}
int QUndoView_viewportEvent_default(void* self, void* event) {
	auto&& __r = ((QUndoView__Go*)self)->cppgoBase_viewportEvent((QEvent*)event);
	return __r;
}
void QUndoView_mousePressEvent_default(void* self, void* event) {
	((QUndoView__Go*)self)->cppgoBase_mousePressEvent((QMouseEvent*)event);
}
void QUndoView_mouseDoubleClickEvent_default(void* self, void* event) {
	((QUndoView__Go*)self)->cppgoBase_mouseDoubleClickEvent((QMouseEvent*)event);
}
void QUndoView_dragEnterEvent_default(void* self, void* event) {
	((QUndoView__Go*)self)->cppgoBase_dragEnterEvent((QDragEnterEvent*)event);
}
void QUndoView_focusInEvent_default(void* self, void* event) {
	((QUndoView__Go*)self)->cppgoBase_focusInEvent((QFocusEvent*)event);
}
void QUndoView_focusOutEvent_default(void* self, void* event) {
	((QUndoView__Go*)self)->cppgoBase_focusOutEvent((QFocusEvent*)event);
}
void QUndoView_keyPressEvent_default(void* self, void* event) {
	((QUndoView__Go*)self)->cppgoBase_keyPressEvent((QKeyEvent*)event);
}
void QUndoView_inputMethodEvent_default(void* self, void* event) {
	((QUndoView__Go*)self)->cppgoBase_inputMethodEvent((QInputMethodEvent*)event);
}
int QUndoView_eventFilter_default(void* self, void* object, void* event) {
	auto&& __r = ((QUndoView__Go*)self)->cppgoBase_eventFilter((QObject*)object, (QEvent*)event);
	return __r;
}
void* QUndoView_minimumSizeHint_default(void* self) {
	auto&& __r = ((QUndoView*)self)->QUndoView::minimumSizeHint();
	return new QSize(__r);
}
void* QUndoView_sizeHint_default(void* self) {
	auto&& __r = ((QUndoView*)self)->QUndoView::sizeHint();
	return new QSize(__r);
}
void QUndoView_setupViewport_default(void* self, void* viewport) {
	((QUndoView*)self)->QUndoView::setupViewport((QWidget*)viewport);
}
void QUndoView_contextMenuEvent_default(void* self, void* arg1) {
	((QUndoView__Go*)self)->cppgoBase_contextMenuEvent((QContextMenuEvent*)arg1);
}
void QUndoView_changeEvent_default(void* self, void* arg1) {
	((QUndoView__Go*)self)->cppgoBase_changeEvent((QEvent*)arg1);
}
void QUndoView_initStyleOption_default(void* self, void* option) {
	((QUndoView__Go*)self)->cppgoBase_initStyleOption((QStyleOptionFrame*)option);
}
int QUndoView_devType_default(void* self) {
	auto&& __r = ((QUndoView*)self)->QUndoView::devType();
	return __r;
}
void QUndoView_setVisible_default(void* self, int visible) {
	((QUndoView*)self)->QUndoView::setVisible(visible);
}
int QUndoView_heightForWidth_default(void* self, int arg1) {
	auto&& __r = ((QUndoView*)self)->QUndoView::heightForWidth(arg1);
	return __r;
}
int QUndoView_hasHeightForWidth_default(void* self) {
	auto&& __r = ((QUndoView*)self)->QUndoView::hasHeightForWidth();
	return __r;
}
void* QUndoView_paintEngine_default(void* self) {
	auto&& __r = ((QUndoView*)self)->QUndoView::paintEngine();
	return (void*)(__r);
}
void QUndoView_keyReleaseEvent_default(void* self, void* event) {
	((QUndoView__Go*)self)->cppgoBase_keyReleaseEvent((QKeyEvent*)event);
}
void QUndoView_enterEvent_default(void* self, void* event) {
	((QUndoView__Go*)self)->cppgoBase_enterEvent((QEnterEvent*)event);
}
void QUndoView_leaveEvent_default(void* self, void* event) {
	((QUndoView__Go*)self)->cppgoBase_leaveEvent((QEvent*)event);
}
void QUndoView_moveEvent_default(void* self, void* event) {
	((QUndoView__Go*)self)->cppgoBase_moveEvent((QMoveEvent*)event);
}
void QUndoView_closeEvent_default(void* self, void* event) {
	((QUndoView__Go*)self)->cppgoBase_closeEvent((QCloseEvent*)event);
}
void QUndoView_tabletEvent_default(void* self, void* event) {
	((QUndoView__Go*)self)->cppgoBase_tabletEvent((QTabletEvent*)event);
}
void QUndoView_actionEvent_default(void* self, void* event) {
	((QUndoView__Go*)self)->cppgoBase_actionEvent((QActionEvent*)event);
}
void QUndoView_showEvent_default(void* self, void* event) {
	((QUndoView__Go*)self)->cppgoBase_showEvent((QShowEvent*)event);
}
void QUndoView_hideEvent_default(void* self, void* event) {
	((QUndoView__Go*)self)->cppgoBase_hideEvent((QHideEvent*)event);
}
int QUndoView_metric_default(void* self, int arg1) {
	auto&& __r = ((QUndoView__Go*)self)->cppgoBase_metric(arg1);
	return __r;
}
void QUndoView_initPainter_default(void* self, void* painter) {
	((QUndoView__Go*)self)->cppgoBase_initPainter((QPainter*)painter);
}
void* QUndoView_redirected_default(void* self, void* offset) {
	auto&& __r = ((QUndoView__Go*)self)->cppgoBase_redirected((QPoint*)offset);
	return (void*)(__r);
}
void* QUndoView_sharedPainter_default(void* self) {
	auto&& __r = ((QUndoView__Go*)self)->cppgoBase_sharedPainter();
	return (void*)(__r);
}
void QUndoView_childEvent_default(void* self, void* event) {
	((QUndoView__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QUndoView_customEvent_default(void* self, void* event) {
	((QUndoView__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QUndoView_connectNotify_default(void* self, void* signal) {
	((QUndoView__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QUndoView_disconnectNotify_default(void* self, void* signal) {
	((QUndoView__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QUndoView_connect_indexesMoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QUndoView*)self, &QUndoView::indexesMoved, (QUndoView*)self, [__id](const QList<QModelIndex>& indexes){
		std::vector<void*> indexes_v; indexes_v.reserve(indexes.size());
		for (long long __i = 0; __i < indexes.size(); __i++) indexes_v.push_back((void*)&indexes[__i]);
		if (cppgo_sig_void_voidpp_longlong) ((void(*)(long long, void**, long long))cppgo_sig_void_voidpp_longlong)(__id, indexes_v.data(), (long long)indexes.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QUndoView);
	__conns_QUndoView[__id] = __c;
}
void QUndoView_connect_pressed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QUndoView*)self, &QUndoView::pressed, (QUndoView*)self, [__id](const QModelIndex& index){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&index);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QUndoView);
	__conns_QUndoView[__id] = __c;
}
void QUndoView_connect_clicked(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QUndoView*)self, &QUndoView::clicked, (QUndoView*)self, [__id](const QModelIndex& index){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&index);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QUndoView);
	__conns_QUndoView[__id] = __c;
}
void QUndoView_connect_doubleClicked(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QUndoView*)self, &QUndoView::doubleClicked, (QUndoView*)self, [__id](const QModelIndex& index){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&index);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QUndoView);
	__conns_QUndoView[__id] = __c;
}
void QUndoView_connect_activated(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QUndoView*)self, &QUndoView::activated, (QUndoView*)self, [__id](const QModelIndex& index){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&index);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QUndoView);
	__conns_QUndoView[__id] = __c;
}
void QUndoView_connect_entered(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QUndoView*)self, &QUndoView::entered, (QUndoView*)self, [__id](const QModelIndex& index){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&index);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QUndoView);
	__conns_QUndoView[__id] = __c;
}
void QUndoView_connect_viewportEntered(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QUndoView*)self, &QUndoView::viewportEntered, (QUndoView*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QUndoView);
	__conns_QUndoView[__id] = __c;
}
void QUndoView_connect_iconSizeChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QUndoView*)self, &QUndoView::iconSizeChanged, (QUndoView*)self, [__id](const QSize& size){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&size);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QUndoView);
	__conns_QUndoView[__id] = __c;
}
void QUndoView_connect_windowTitleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QUndoView*)self, &QUndoView::windowTitleChanged, (QUndoView*)self, [__id](const QString& title){
		QByteArray title_ba = title.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)title_ba.constData(), title_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QUndoView);
	__conns_QUndoView[__id] = __c;
}
void QUndoView_connect_windowIconChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QUndoView*)self, &QUndoView::windowIconChanged, (QUndoView*)self, [__id](const QIcon& icon){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&icon);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QUndoView);
	__conns_QUndoView[__id] = __c;
}
void QUndoView_connect_windowIconTextChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QUndoView*)self, &QUndoView::windowIconTextChanged, (QUndoView*)self, [__id](const QString& iconText){
		QByteArray iconText_ba = iconText.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)iconText_ba.constData(), iconText_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QUndoView);
	__conns_QUndoView[__id] = __c;
}
void QUndoView_connect_customContextMenuRequested(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QUndoView*)self, &QUndoView::customContextMenuRequested, (QUndoView*)self, [__id](const QPoint& pos){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&pos);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QUndoView);
	__conns_QUndoView[__id] = __c;
}
void QUndoView_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QUndoView*)self, &QUndoView::destroyed, (QUndoView*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QUndoView);
	__conns_QUndoView[__id] = __c;
}
void QUndoView_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QUndoView*)self, &QUndoView::destroyed, (QUndoView*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QUndoView);
	__conns_QUndoView[__id] = __c;
}
void QUndoView_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QUndoView*)self, &QUndoView::objectNameChanged, (QUndoView*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QUndoView);
	__conns_QUndoView[__id] = __c;
}
void QUndoView__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QUndoView);
		auto __it = __conns_QUndoView.find(id);
		if (__it == __conns_QUndoView.end()) return;
		__c = __it->second;
		__conns_QUndoView.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QUndoView__Go_new(long long id) { return new QUndoView__Go(id); }
void QUndoView_delete(void* self) { delete (QUndoView*)self; }

}

