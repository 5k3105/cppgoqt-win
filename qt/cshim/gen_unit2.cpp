#include <QtCore/qlogging.h>
#include <QtCore/qloggingcategory.h>
#include <QtCore/qdebug.h>
#include "gen_qmessagelogger.h"

extern "C" {

void* QMessageLogger_new() {
	return new QMessageLogger();
}
void* QMessageLogger_new2(const char* file, int line, const char* function) {
	return new QMessageLogger(file, line, function);
}
void* QMessageLogger_new3(const char* file, int line, const char* function, const char* category) {
	return new QMessageLogger(file, line, function, category);
}
void QMessageLogger_debug(void* self, const char* msg) {
	((QMessageLogger*)self)->debug(msg);
}
void QMessageLogger_noDebug(void* self, const char* arg1) {
	((QMessageLogger*)self)->noDebug(arg1);
}
void QMessageLogger_info(void* self, const char* msg) {
	((QMessageLogger*)self)->info(msg);
}
void QMessageLogger_warning(void* self, const char* msg) {
	((QMessageLogger*)self)->warning(msg);
}
void QMessageLogger_critical(void* self, const char* msg) {
	((QMessageLogger*)self)->critical(msg);
}
void QMessageLogger_fatal(void* self, const char* msg) {
	((QMessageLogger*)self)->fatal(msg);
}
void QMessageLogger_debug2(void* self, void* cat, const char* msg) {
	((QMessageLogger*)self)->debug(*(QLoggingCategory*)cat, msg);
}
void QMessageLogger_info2(void* self, void* cat, const char* msg) {
	((QMessageLogger*)self)->info(*(QLoggingCategory*)cat, msg);
}
void QMessageLogger_warning2(void* self, void* cat, const char* msg) {
	((QMessageLogger*)self)->warning(*(QLoggingCategory*)cat, msg);
}
void QMessageLogger_critical2(void* self, void* cat, const char* msg) {
	((QMessageLogger*)self)->critical(*(QLoggingCategory*)cat, msg);
}
void QMessageLogger_fatal2(void* self, void* cat, const char* msg) {
	((QMessageLogger*)self)->fatal(*(QLoggingCategory*)cat, msg);
}
void* QMessageLogger_debug3(void* self) {
	auto&& __r = ((QMessageLogger*)self)->debug();
	return new QDebug(__r);
}
void* QMessageLogger_debug4(void* self, void* cat) {
	auto&& __r = ((QMessageLogger*)self)->debug(*(QLoggingCategory*)cat);
	return new QDebug(__r);
}
void* QMessageLogger_info3(void* self) {
	auto&& __r = ((QMessageLogger*)self)->info();
	return new QDebug(__r);
}
void* QMessageLogger_info4(void* self, void* cat) {
	auto&& __r = ((QMessageLogger*)self)->info(*(QLoggingCategory*)cat);
	return new QDebug(__r);
}
void* QMessageLogger_warning3(void* self) {
	auto&& __r = ((QMessageLogger*)self)->warning();
	return new QDebug(__r);
}
void* QMessageLogger_warning4(void* self, void* cat) {
	auto&& __r = ((QMessageLogger*)self)->warning(*(QLoggingCategory*)cat);
	return new QDebug(__r);
}
void* QMessageLogger_critical3(void* self) {
	auto&& __r = ((QMessageLogger*)self)->critical();
	return new QDebug(__r);
}
void* QMessageLogger_critical4(void* self, void* cat) {
	auto&& __r = ((QMessageLogger*)self)->critical(*(QLoggingCategory*)cat);
	return new QDebug(__r);
}
void* QMessageLogger_fatal3(void* self) {
	auto&& __r = ((QMessageLogger*)self)->fatal();
	return new QDebug(__r);
}
void* QMessageLogger_fatal4(void* self, void* cat) {
	auto&& __r = ((QMessageLogger*)self)->fatal(*(QLoggingCategory*)cat);
	return new QDebug(__r);
}
void* QMessageLogger_noDebug2(void* self) {
	auto&& __r = ((QMessageLogger*)self)->noDebug();
	return new QNoDebug(__r);
}
void QMessageLogger_delete(void* self) { delete (QMessageLogger*)self; }

}

#include <QtCore/qlatin1stringview.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qbytearrayview.h>
#include <QtCore/qstring.h>
#include <QtCore/qchar.h>
#include <QtCore/qstringview.h>
#include <QtCore/qnamespace.h>
#include "gen_qlatin1string.h"

extern "C" {

void* QLatin1String_new() {
	return new QLatin1String();
}
void* QLatin1String_new2(const char* s) {
	return new QLatin1String(s);
}
void* QLatin1String_new3(const char* f, const char* l) {
	return new QLatin1String(f, l);
}
void* QLatin1String_new4(const char* s, long long sz) {
	return new QLatin1String(s, sz);
}
void* QLatin1String_new5(void* s) {
	return new QLatin1String(*(QByteArray*)s);
}
void* QLatin1String_new6(void* s) {
	return new QLatin1String(*(QByteArrayView*)s);
}
void* QLatin1String_new7(void* arg1) {
	return new QLatin1String(*(QLatin1String*)arg1);
}
char* QLatin1String_toString(void* self, long long* __out_len) {
	auto&& __r = ((QLatin1String*)self)->toString();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QLatin1String_toUtf8(void* self) {
	auto&& __r = ((QLatin1String*)self)->toUtf8();
	return new QByteArray(__r);
}
const char* QLatin1String_latin1(void* self) {
	auto&& __r = ((QLatin1String*)self)->latin1();
	return __r;
}
long long QLatin1String_size(void* self) {
	auto&& __r = ((QLatin1String*)self)->size();
	return __r;
}
const char* QLatin1String_data(void* self) {
	auto&& __r = ((QLatin1String*)self)->data();
	return __r;
}
const char* QLatin1String_constData(void* self) {
	auto&& __r = ((QLatin1String*)self)->constData();
	return __r;
}
const char* QLatin1String_constBegin(void* self) {
	auto&& __r = ((QLatin1String*)self)->constBegin();
	return __r;
}
const char* QLatin1String_constEnd(void* self) {
	auto&& __r = ((QLatin1String*)self)->constEnd();
	return __r;
}
void* QLatin1String_first(void* self) {
	auto&& __r = ((QLatin1String*)self)->first();
	return new QLatin1Char(__r);
}
void* QLatin1String_last(void* self) {
	auto&& __r = ((QLatin1String*)self)->last();
	return new QLatin1Char(__r);
}
long long QLatin1String_length(void* self) {
	auto&& __r = ((QLatin1String*)self)->length();
	return __r;
}
int QLatin1String_isNull(void* self) {
	auto&& __r = ((QLatin1String*)self)->isNull();
	return __r;
}
int QLatin1String_isEmpty(void* self) {
	auto&& __r = ((QLatin1String*)self)->isEmpty();
	return __r;
}
int QLatin1String_empty(void* self) {
	auto&& __r = ((QLatin1String*)self)->empty();
	return __r;
}
void* QLatin1String_at(void* self, long long i) {
	auto&& __r = ((QLatin1String*)self)->at(i);
	return new QLatin1Char(__r);
}
void* QLatin1String_OperatorSubscript(void* self, long long i) {
	auto&& __r = ((QLatin1String*)self)->operator[](i);
	return new QLatin1Char(__r);
}
void* QLatin1String_front(void* self) {
	auto&& __r = ((QLatin1String*)self)->front();
	return new QLatin1Char(__r);
}
void* QLatin1String_back(void* self) {
	auto&& __r = ((QLatin1String*)self)->back();
	return new QLatin1Char(__r);
}
int QLatin1String_compare(void* self, void* other) {
	auto&& __r = ((QLatin1String*)self)->compare(*(QStringView*)other);
	return __r;
}
int QLatin1String_compare2(void* self, void* other) {
	auto&& __r = ((QLatin1String*)self)->compare(*(QLatin1String*)other);
	return __r;
}
int QLatin1String_compare3(void* self, void* c) {
	auto&& __r = ((QLatin1String*)self)->compare(*(QChar*)c);
	return __r;
}
int QLatin1String_compare4(void* self, void* c, int cs) {
	auto&& __r = ((QLatin1String*)self)->compare(*(QChar*)c, (Qt::CaseSensitivity)(cs));
	return __r;
}
int QLatin1String_startsWith(void* self, void* s) {
	auto&& __r = ((QLatin1String*)self)->startsWith(*(QStringView*)s);
	return __r;
}
int QLatin1String_startsWith2(void* self, void* s) {
	auto&& __r = ((QLatin1String*)self)->startsWith(*(QLatin1String*)s);
	return __r;
}
int QLatin1String_startsWith3(void* self, void* c) {
	auto&& __r = ((QLatin1String*)self)->startsWith(*(QChar*)c);
	return __r;
}
int QLatin1String_startsWith4(void* self, void* c, int cs) {
	auto&& __r = ((QLatin1String*)self)->startsWith(*(QChar*)c, (Qt::CaseSensitivity)(cs));
	return __r;
}
int QLatin1String_endsWith(void* self, void* s) {
	auto&& __r = ((QLatin1String*)self)->endsWith(*(QStringView*)s);
	return __r;
}
int QLatin1String_endsWith2(void* self, void* s) {
	auto&& __r = ((QLatin1String*)self)->endsWith(*(QLatin1String*)s);
	return __r;
}
int QLatin1String_endsWith3(void* self, void* c) {
	auto&& __r = ((QLatin1String*)self)->endsWith(*(QChar*)c);
	return __r;
}
int QLatin1String_endsWith4(void* self, void* c, int cs) {
	auto&& __r = ((QLatin1String*)self)->endsWith(*(QChar*)c, (Qt::CaseSensitivity)(cs));
	return __r;
}
long long QLatin1String_indexOf(void* self, void* s) {
	auto&& __r = ((QLatin1String*)self)->indexOf(*(QStringView*)s);
	return __r;
}
long long QLatin1String_indexOf2(void* self, void* s) {
	auto&& __r = ((QLatin1String*)self)->indexOf(*(QLatin1String*)s);
	return __r;
}
long long QLatin1String_indexOf3(void* self, void* c) {
	auto&& __r = ((QLatin1String*)self)->indexOf(*(QChar*)c);
	return __r;
}
long long QLatin1String_indexOf4(void* self, void* c, long long from, int cs) {
	auto&& __r = ((QLatin1String*)self)->indexOf(*(QChar*)c, from, (Qt::CaseSensitivity)(cs));
	return __r;
}
int QLatin1String_contains(void* self, void* s) {
	auto&& __r = ((QLatin1String*)self)->contains(*(QStringView*)s);
	return __r;
}
int QLatin1String_contains2(void* self, void* s) {
	auto&& __r = ((QLatin1String*)self)->contains(*(QLatin1String*)s);
	return __r;
}
int QLatin1String_contains3(void* self, void* c) {
	auto&& __r = ((QLatin1String*)self)->contains(*(QChar*)c);
	return __r;
}
long long QLatin1String_lastIndexOf(void* self, void* s) {
	auto&& __r = ((QLatin1String*)self)->lastIndexOf(*(QStringView*)s);
	return __r;
}
long long QLatin1String_lastIndexOf2(void* self, void* s, long long from) {
	auto&& __r = ((QLatin1String*)self)->lastIndexOf(*(QStringView*)s, from);
	return __r;
}
long long QLatin1String_lastIndexOf3(void* self, void* s) {
	auto&& __r = ((QLatin1String*)self)->lastIndexOf(*(QLatin1String*)s);
	return __r;
}
long long QLatin1String_lastIndexOf4(void* self, void* s, long long from) {
	auto&& __r = ((QLatin1String*)self)->lastIndexOf(*(QLatin1String*)s, from);
	return __r;
}
long long QLatin1String_lastIndexOf5(void* self, void* c) {
	auto&& __r = ((QLatin1String*)self)->lastIndexOf(*(QChar*)c);
	return __r;
}
long long QLatin1String_lastIndexOf6(void* self, void* c, int cs) {
	auto&& __r = ((QLatin1String*)self)->lastIndexOf(*(QChar*)c, (Qt::CaseSensitivity)(cs));
	return __r;
}
long long QLatin1String_lastIndexOf7(void* self, void* c, long long from) {
	auto&& __r = ((QLatin1String*)self)->lastIndexOf(*(QChar*)c, from);
	return __r;
}
long long QLatin1String_lastIndexOf8(void* self, void* c, long long from, int cs) {
	auto&& __r = ((QLatin1String*)self)->lastIndexOf(*(QChar*)c, from, (Qt::CaseSensitivity)(cs));
	return __r;
}
long long QLatin1String_count(void* self, void* str) {
	auto&& __r = ((QLatin1String*)self)->count(*(QStringView*)str);
	return __r;
}
long long QLatin1String_count2(void* self, void* str) {
	auto&& __r = ((QLatin1String*)self)->count(*(QLatin1String*)str);
	return __r;
}
long long QLatin1String_count3(void* self, void* ch) {
	auto&& __r = ((QLatin1String*)self)->count(*(QChar*)ch);
	return __r;
}
short QLatin1String_toShort(void* self) {
	auto&& __r = ((QLatin1String*)self)->toShort();
	return __r;
}
unsigned short QLatin1String_toUShort(void* self) {
	auto&& __r = ((QLatin1String*)self)->toUShort();
	return __r;
}
int QLatin1String_toInt(void* self) {
	auto&& __r = ((QLatin1String*)self)->toInt();
	return __r;
}
unsigned int QLatin1String_toUInt(void* self) {
	auto&& __r = ((QLatin1String*)self)->toUInt();
	return __r;
}
long QLatin1String_toLong(void* self) {
	auto&& __r = ((QLatin1String*)self)->toLong();
	return __r;
}
unsigned long QLatin1String_toULong(void* self) {
	auto&& __r = ((QLatin1String*)self)->toULong();
	return __r;
}
long long QLatin1String_toLongLong(void* self) {
	auto&& __r = ((QLatin1String*)self)->toLongLong();
	return __r;
}
unsigned long long QLatin1String_toULongLong(void* self) {
	auto&& __r = ((QLatin1String*)self)->toULongLong();
	return __r;
}
float QLatin1String_toFloat(void* self) {
	auto&& __r = ((QLatin1String*)self)->toFloat();
	return __r;
}
double QLatin1String_toDouble(void* self) {
	auto&& __r = ((QLatin1String*)self)->toDouble();
	return __r;
}
const char* QLatin1String_begin(void* self) {
	auto&& __r = ((QLatin1String*)self)->begin();
	return __r;
}
const char* QLatin1String_cbegin(void* self) {
	auto&& __r = ((QLatin1String*)self)->cbegin();
	return __r;
}
const char* QLatin1String_end(void* self) {
	auto&& __r = ((QLatin1String*)self)->end();
	return __r;
}
const char* QLatin1String_cend(void* self) {
	auto&& __r = ((QLatin1String*)self)->cend();
	return __r;
}
long long QLatin1String_max_size(void* self) {
	auto&& __r = ((QLatin1String*)self)->max_size();
	return __r;
}
long long QLatin1String_maxSize() {
	auto&& __r = QLatin1String::maxSize();
	return __r;
}
void* QLatin1String_mid(void* self, long long pos) {
	auto&& __r = ((QLatin1String*)self)->mid(pos);
	return new QLatin1String(__r);
}
void* QLatin1String_left(void* self, long long n) {
	auto&& __r = ((QLatin1String*)self)->left(n);
	return new QLatin1String(__r);
}
void* QLatin1String_right(void* self, long long n) {
	auto&& __r = ((QLatin1String*)self)->right(n);
	return new QLatin1String(__r);
}
void* QLatin1String_sliced(void* self, long long pos) {
	auto&& __r = ((QLatin1String*)self)->sliced(pos);
	return new QLatin1String(__r);
}
void* QLatin1String_sliced2(void* self, long long pos, long long n) {
	auto&& __r = ((QLatin1String*)self)->sliced(pos, n);
	return new QLatin1String(__r);
}
void* QLatin1String_first2(void* self, long long n) {
	auto&& __r = ((QLatin1String*)self)->first(n);
	return new QLatin1String(__r);
}
void* QLatin1String_last2(void* self, long long n) {
	auto&& __r = ((QLatin1String*)self)->last(n);
	return new QLatin1String(__r);
}
void* QLatin1String_chopped(void* self, long long n) {
	auto&& __r = ((QLatin1String*)self)->chopped(n);
	return new QLatin1String(__r);
}
void QLatin1String_chop(void* self, long long n) {
	((QLatin1String*)self)->chop(n);
}
void QLatin1String_truncate(void* self, long long n) {
	((QLatin1String*)self)->truncate(n);
}
void* QLatin1String_trimmed(void* self) {
	auto&& __r = ((QLatin1String*)self)->trimmed();
	return new QLatin1String(__r);
}
void* QLatin1String_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QLatin1String*)self)->operator=(*(QLatin1String*)arg1);
	return (void*)&(__r);
}
int QLatin1String_compare5(void* self, void* other, int cs) {
	auto&& __r = ((QLatin1String*)self)->compare(*(QStringView*)other, (Qt::CaseSensitivity)(cs));
	return __r;
}
int QLatin1String_compare6(void* self, void* other, int cs) {
	auto&& __r = ((QLatin1String*)self)->compare(*(QLatin1String*)other, (Qt::CaseSensitivity)(cs));
	return __r;
}
int QLatin1String_startsWith5(void* self, void* s, int cs) {
	auto&& __r = ((QLatin1String*)self)->startsWith(*(QStringView*)s, (Qt::CaseSensitivity)(cs));
	return __r;
}
int QLatin1String_startsWith6(void* self, void* s, int cs) {
	auto&& __r = ((QLatin1String*)self)->startsWith(*(QLatin1String*)s, (Qt::CaseSensitivity)(cs));
	return __r;
}
int QLatin1String_endsWith5(void* self, void* s, int cs) {
	auto&& __r = ((QLatin1String*)self)->endsWith(*(QStringView*)s, (Qt::CaseSensitivity)(cs));
	return __r;
}
int QLatin1String_endsWith6(void* self, void* s, int cs) {
	auto&& __r = ((QLatin1String*)self)->endsWith(*(QLatin1String*)s, (Qt::CaseSensitivity)(cs));
	return __r;
}
long long QLatin1String_indexOf5(void* self, void* s, long long from) {
	auto&& __r = ((QLatin1String*)self)->indexOf(*(QStringView*)s, from);
	return __r;
}
long long QLatin1String_indexOf6(void* self, void* s, long long from, int cs) {
	auto&& __r = ((QLatin1String*)self)->indexOf(*(QStringView*)s, from, (Qt::CaseSensitivity)(cs));
	return __r;
}
long long QLatin1String_indexOf7(void* self, void* s, long long from) {
	auto&& __r = ((QLatin1String*)self)->indexOf(*(QLatin1String*)s, from);
	return __r;
}
long long QLatin1String_indexOf8(void* self, void* s, long long from, int cs) {
	auto&& __r = ((QLatin1String*)self)->indexOf(*(QLatin1String*)s, from, (Qt::CaseSensitivity)(cs));
	return __r;
}
long long QLatin1String_indexOf9(void* self, void* c, long long from) {
	auto&& __r = ((QLatin1String*)self)->indexOf(*(QChar*)c, from);
	return __r;
}
int QLatin1String_contains4(void* self, void* s, int cs) {
	auto&& __r = ((QLatin1String*)self)->contains(*(QStringView*)s, (Qt::CaseSensitivity)(cs));
	return __r;
}
int QLatin1String_contains5(void* self, void* s, int cs) {
	auto&& __r = ((QLatin1String*)self)->contains(*(QLatin1String*)s, (Qt::CaseSensitivity)(cs));
	return __r;
}
int QLatin1String_contains6(void* self, void* c, int cs) {
	auto&& __r = ((QLatin1String*)self)->contains(*(QChar*)c, (Qt::CaseSensitivity)(cs));
	return __r;
}
long long QLatin1String_lastIndexOf9(void* self, void* s, int cs) {
	auto&& __r = ((QLatin1String*)self)->lastIndexOf(*(QStringView*)s, (Qt::CaseSensitivity)(cs));
	return __r;
}
long long QLatin1String_lastIndexOf10(void* self, void* s, long long from, int cs) {
	auto&& __r = ((QLatin1String*)self)->lastIndexOf(*(QStringView*)s, from, (Qt::CaseSensitivity)(cs));
	return __r;
}
long long QLatin1String_lastIndexOf11(void* self, void* s, int cs) {
	auto&& __r = ((QLatin1String*)self)->lastIndexOf(*(QLatin1String*)s, (Qt::CaseSensitivity)(cs));
	return __r;
}
long long QLatin1String_lastIndexOf12(void* self, void* s, long long from, int cs) {
	auto&& __r = ((QLatin1String*)self)->lastIndexOf(*(QLatin1String*)s, from, (Qt::CaseSensitivity)(cs));
	return __r;
}
long long QLatin1String_count4(void* self, void* str, int cs) {
	auto&& __r = ((QLatin1String*)self)->count(*(QStringView*)str, (Qt::CaseSensitivity)(cs));
	return __r;
}
long long QLatin1String_count5(void* self, void* str, int cs) {
	auto&& __r = ((QLatin1String*)self)->count(*(QLatin1String*)str, (Qt::CaseSensitivity)(cs));
	return __r;
}
long long QLatin1String_count6(void* self, void* ch, int cs) {
	auto&& __r = ((QLatin1String*)self)->count(*(QChar*)ch, (Qt::CaseSensitivity)(cs));
	return __r;
}
short QLatin1String_toShort2(void* self, int* ok) {
	bool ok_local{};
	auto&& __r = ((QLatin1String*)self)->toShort(&ok_local);
	*ok = ok_local;
	return __r;
}
short QLatin1String_toShort3(void* self, int* ok, int base) {
	bool ok_local{};
	auto&& __r = ((QLatin1String*)self)->toShort(&ok_local, base);
	*ok = ok_local;
	return __r;
}
unsigned short QLatin1String_toUShort2(void* self, int* ok) {
	bool ok_local{};
	auto&& __r = ((QLatin1String*)self)->toUShort(&ok_local);
	*ok = ok_local;
	return __r;
}
unsigned short QLatin1String_toUShort3(void* self, int* ok, int base) {
	bool ok_local{};
	auto&& __r = ((QLatin1String*)self)->toUShort(&ok_local, base);
	*ok = ok_local;
	return __r;
}
int QLatin1String_toInt2(void* self, int* ok) {
	bool ok_local{};
	auto&& __r = ((QLatin1String*)self)->toInt(&ok_local);
	*ok = ok_local;
	return __r;
}
int QLatin1String_toInt3(void* self, int* ok, int base) {
	bool ok_local{};
	auto&& __r = ((QLatin1String*)self)->toInt(&ok_local, base);
	*ok = ok_local;
	return __r;
}
unsigned int QLatin1String_toUInt2(void* self, int* ok) {
	bool ok_local{};
	auto&& __r = ((QLatin1String*)self)->toUInt(&ok_local);
	*ok = ok_local;
	return __r;
}
unsigned int QLatin1String_toUInt3(void* self, int* ok, int base) {
	bool ok_local{};
	auto&& __r = ((QLatin1String*)self)->toUInt(&ok_local, base);
	*ok = ok_local;
	return __r;
}
long QLatin1String_toLong2(void* self, int* ok) {
	bool ok_local{};
	auto&& __r = ((QLatin1String*)self)->toLong(&ok_local);
	*ok = ok_local;
	return __r;
}
long QLatin1String_toLong3(void* self, int* ok, int base) {
	bool ok_local{};
	auto&& __r = ((QLatin1String*)self)->toLong(&ok_local, base);
	*ok = ok_local;
	return __r;
}
unsigned long QLatin1String_toULong2(void* self, int* ok) {
	bool ok_local{};
	auto&& __r = ((QLatin1String*)self)->toULong(&ok_local);
	*ok = ok_local;
	return __r;
}
unsigned long QLatin1String_toULong3(void* self, int* ok, int base) {
	bool ok_local{};
	auto&& __r = ((QLatin1String*)self)->toULong(&ok_local, base);
	*ok = ok_local;
	return __r;
}
long long QLatin1String_toLongLong2(void* self, int* ok) {
	bool ok_local{};
	auto&& __r = ((QLatin1String*)self)->toLongLong(&ok_local);
	*ok = ok_local;
	return __r;
}
long long QLatin1String_toLongLong3(void* self, int* ok, int base) {
	bool ok_local{};
	auto&& __r = ((QLatin1String*)self)->toLongLong(&ok_local, base);
	*ok = ok_local;
	return __r;
}
unsigned long long QLatin1String_toULongLong2(void* self, int* ok) {
	bool ok_local{};
	auto&& __r = ((QLatin1String*)self)->toULongLong(&ok_local);
	*ok = ok_local;
	return __r;
}
unsigned long long QLatin1String_toULongLong3(void* self, int* ok, int base) {
	bool ok_local{};
	auto&& __r = ((QLatin1String*)self)->toULongLong(&ok_local, base);
	*ok = ok_local;
	return __r;
}
float QLatin1String_toFloat2(void* self, int* ok) {
	bool ok_local{};
	auto&& __r = ((QLatin1String*)self)->toFloat(&ok_local);
	*ok = ok_local;
	return __r;
}
double QLatin1String_toDouble2(void* self, int* ok) {
	bool ok_local{};
	auto&& __r = ((QLatin1String*)self)->toDouble(&ok_local);
	*ok = ok_local;
	return __r;
}
void* QLatin1String_mid2(void* self, long long pos, long long n) {
	auto&& __r = ((QLatin1String*)self)->mid(pos, n);
	return new QLatin1String(__r);
}
void QLatin1String_delete(void* self) { delete (QLatin1String*)self; }

}

#include <QtCore/qabstractanimation.h>
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
#include "gen_qanimationdriver.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QAnimationDriver__Go : public QAnimationDriver {
public:
	long long __goId;
	unsigned char __goOn[12] = {};
	template<typename... A> QAnimationDriver__Go(long long id, A&&... a) : QAnimationDriver(std::forward<A>(a)...), __goId(id) {}
	using QAnimationDriver::advanceAnimation;
	using QAnimationDriver::sender;
	using QAnimationDriver::senderSignalIndex;
	using QAnimationDriver::receivers;
	using QAnimationDriver::isSignalConnected;
	void cppgoBase_start() { QAnimationDriver::start(); }
	void cppgoBase_stop() { QAnimationDriver::stop(); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QAnimationDriver::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QAnimationDriver::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QAnimationDriver::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QAnimationDriver::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QAnimationDriver::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QAnimationDriver::metaObject();
	}
	void advance() override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, -4830836341480073475LL, &__h); if (__h) return; }
		QAnimationDriver::advance();
	}
	qint64 elapsed() const override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_longlong) { long long __r = ((long long(*)(long long, long long, int*))cppgo_virt_longlong)(__goId, -1391384709503831749LL, &__h); if (__h) return __r; }
		return QAnimationDriver::elapsed();
	}
	void start() override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, -1270692749394500321LL, &__h); if (__h) return; }
		QAnimationDriver::start();
	}
	void stop() override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, -5827849124968219451LL, &__h); if (__h) return; }
		QAnimationDriver::stop();
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QAnimationDriver::event(event);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QAnimationDriver::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QAnimationDriver::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QAnimationDriver::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QAnimationDriver::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QAnimationDriver::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QAnimationDriver::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QAnimationDriver;
static std::mutex __connsMu_QAnimationDriver;

extern "C" {

void* QAnimationDriver_base(void* self) { return (void*)static_cast<QObject*>((QAnimationDriver*)self); }
void QAnimationDriver__set_override(void* self, int idx) { ((QAnimationDriver__Go*)self)->__goOn[idx] = 1; }
void* QAnimationDriver_fromQObject(void* o) { return (void*)qobject_cast<QAnimationDriver*>((QObject*)o); }
void* QAnimationDriver_new(long long __goid) {
	return new QAnimationDriver__Go(__goid);
}
void* QAnimationDriver_new2(long long __goid, void* parent) {
	return new QAnimationDriver__Go(__goid, (QObject*)parent);
}
char* QAnimationDriver_tr(const char* s, long long* __out_len) {
	auto&& __r = QAnimationDriver::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QAnimationDriver_advance(void* self) {
	((QAnimationDriver*)self)->advance();
}
void QAnimationDriver_install(void* self) {
	((QAnimationDriver*)self)->install();
}
void QAnimationDriver_uninstall(void* self) {
	((QAnimationDriver*)self)->uninstall();
}
int QAnimationDriver_isRunning(void* self) {
	auto&& __r = ((QAnimationDriver*)self)->isRunning();
	return __r;
}
long long QAnimationDriver_elapsed(void* self) {
	auto&& __r = ((QAnimationDriver*)self)->elapsed();
	return __r;
}
void QAnimationDriver_started(void* self) {
	((QAnimationDriver*)self)->started();
}
void QAnimationDriver_stopped(void* self) {
	((QAnimationDriver*)self)->stopped();
}
void QAnimationDriver_advanceAnimation(void* self) {
	((QAnimationDriver__Go*)self)->advanceAnimation();
}
void* QAnimationDriver_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QAnimationDriver::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QAnimationDriver_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QAnimationDriver::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QAnimationDriver_sender(void* self) {
	auto&& __r = ((QAnimationDriver__Go*)self)->sender();
	return (void*)(__r);
}
int QAnimationDriver_senderSignalIndex(void* self) {
	auto&& __r = ((QAnimationDriver__Go*)self)->senderSignalIndex();
	return __r;
}
int QAnimationDriver_receivers(void* self, const char* signal) {
	auto&& __r = ((QAnimationDriver__Go*)self)->receivers(signal);
	return __r;
}
int QAnimationDriver_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QAnimationDriver__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QAnimationDriver_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QAnimationDriver::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QAnimationDriver_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QAnimationDriver::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QAnimationDriver_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QAnimationDriver::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QAnimationDriver_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QAnimationDriver::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void QAnimationDriver_advance_default(void* self) {
	((QAnimationDriver*)self)->QAnimationDriver::advance();
}
long long QAnimationDriver_elapsed_default(void* self) {
	auto&& __r = ((QAnimationDriver*)self)->QAnimationDriver::elapsed();
	return __r;
}
void* QAnimationDriver_metaObject_default(void* self) {
	auto&& __r = ((QAnimationDriver*)self)->QAnimationDriver::metaObject();
	return (void*)(__r);
}
void QAnimationDriver_start_default(void* self) {
	((QAnimationDriver__Go*)self)->cppgoBase_start();
}
void QAnimationDriver_stop_default(void* self) {
	((QAnimationDriver__Go*)self)->cppgoBase_stop();
}
int QAnimationDriver_event_default(void* self, void* event) {
	auto&& __r = ((QAnimationDriver*)self)->QAnimationDriver::event((QEvent*)event);
	return __r;
}
int QAnimationDriver_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QAnimationDriver*)self)->QAnimationDriver::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QAnimationDriver_timerEvent_default(void* self, void* event) {
	((QAnimationDriver__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QAnimationDriver_childEvent_default(void* self, void* event) {
	((QAnimationDriver__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QAnimationDriver_customEvent_default(void* self, void* event) {
	((QAnimationDriver__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QAnimationDriver_connectNotify_default(void* self, void* signal) {
	((QAnimationDriver__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QAnimationDriver_disconnectNotify_default(void* self, void* signal) {
	((QAnimationDriver__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QAnimationDriver_connect_started(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAnimationDriver*)self, &QAnimationDriver::started, (QAnimationDriver*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAnimationDriver);
	__conns_QAnimationDriver[__id] = __c;
}
void QAnimationDriver_connect_stopped(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAnimationDriver*)self, &QAnimationDriver::stopped, (QAnimationDriver*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAnimationDriver);
	__conns_QAnimationDriver[__id] = __c;
}
void QAnimationDriver_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAnimationDriver*)self, &QAnimationDriver::destroyed, (QAnimationDriver*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAnimationDriver);
	__conns_QAnimationDriver[__id] = __c;
}
void QAnimationDriver_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAnimationDriver*)self, &QAnimationDriver::destroyed, (QAnimationDriver*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAnimationDriver);
	__conns_QAnimationDriver[__id] = __c;
}
void QAnimationDriver_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAnimationDriver*)self, &QAnimationDriver::objectNameChanged, (QAnimationDriver*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAnimationDriver);
	__conns_QAnimationDriver[__id] = __c;
}
void QAnimationDriver__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QAnimationDriver);
		auto __it = __conns_QAnimationDriver.find(id);
		if (__it == __conns_QAnimationDriver.end()) return;
		__c = __it->second;
		__conns_QAnimationDriver.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QAnimationDriver__Go_new(long long id, void* parent) { return new QAnimationDriver__Go(id, (QObject*)parent); }
void QAnimationDriver_delete(void* self) { delete (QAnimationDriver*)self; }

}

#include <QtCore/qabstractitemmodel.h>
#include <QtCore/qobject.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qvariant.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qmimedata.h>
#include <QtCore/qsize.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qdatastream.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qabstractitemmodel.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QAbstractItemModel__Go : public QAbstractItemModel {
public:
	long long __goId;
	unsigned char __goOn[42] = {};
	template<typename... A> QAbstractItemModel__Go(long long id, A&&... a) : QAbstractItemModel(std::forward<A>(a)...), __goId(id) {}
	using QAbstractItemModel::createIndex;
	using QAbstractItemModel::encodeData;
	using QAbstractItemModel::decodeData;
	using QAbstractItemModel::beginInsertRows;
	using QAbstractItemModel::endInsertRows;
	using QAbstractItemModel::beginRemoveRows;
	using QAbstractItemModel::endRemoveRows;
	using QAbstractItemModel::beginMoveRows;
	using QAbstractItemModel::endMoveRows;
	using QAbstractItemModel::beginInsertColumns;
	using QAbstractItemModel::endInsertColumns;
	using QAbstractItemModel::beginRemoveColumns;
	using QAbstractItemModel::endRemoveColumns;
	using QAbstractItemModel::beginMoveColumns;
	using QAbstractItemModel::endMoveColumns;
	using QAbstractItemModel::beginResetModel;
	using QAbstractItemModel::endResetModel;
	using QAbstractItemModel::changePersistentIndex;
	using QAbstractItemModel::changePersistentIndexList;
	using QAbstractItemModel::persistentIndexList;
	using QAbstractItemModel::sender;
	using QAbstractItemModel::senderSignalIndex;
	using QAbstractItemModel::receivers;
	using QAbstractItemModel::isSignalConnected;
	void cppgoBase_resetInternalData() { QAbstractItemModel::resetInternalData(); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QAbstractItemModel::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QAbstractItemModel::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QAbstractItemModel::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QAbstractItemModel::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QAbstractItemModel::disconnectNotify(__a0); }
	QModelIndex index(int row, int column, const QModelIndex& parent) const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp_int_int_voidp) { void* __r = ((void*(*)(long long, long long, int, int, void*, int*))cppgo_virt_voidp_int_int_voidp)(__goId, -8948777187306027381LL, row, column, (void*)&parent, &__h); if (__h) return *(QModelIndex*)__r; }
		return QModelIndex{};
	}
	QModelIndex parent(const QModelIndex& child) const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 7371339085820260189LL, (void*)&child, &__h); if (__h) return *(QModelIndex*)__r; }
		return QModelIndex{};
	}
	int rowCount(const QModelIndex& parent) const override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1960960171956884010LL, (void*)&parent, &__h); if (__h) return __r; }
		return int{};
	}
	int columnCount(const QModelIndex& parent) const override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, -8256963024669584834LL, (void*)&parent, &__h); if (__h) return __r; }
		return int{};
	}
	QVariant data(const QModelIndex& index, int role) const override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_voidp_voidp_int) { void* __r = ((void*(*)(long long, long long, void*, int, int*))cppgo_virt_voidp_voidp_int)(__goId, -8837375942098466299LL, (void*)&index, role, &__h); if (__h) return *(QVariant*)__r; }
		return QVariant{};
	}
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QAbstractItemModel::metaObject();
	}
	QModelIndex sibling(int row, int column, const QModelIndex& idx) const override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_voidp_int_int_voidp) { void* __r = ((void*(*)(long long, long long, int, int, void*, int*))cppgo_virt_voidp_int_int_voidp)(__goId, 397329346813782521LL, row, column, (void*)&idx, &__h); if (__h) return *(QModelIndex*)__r; }
		return QAbstractItemModel::sibling(row, column, idx);
	}
	bool hasChildren(const QModelIndex& parent) const override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 3369452259144514380LL, (void*)&parent, &__h); if (__h) return __r; }
		return QAbstractItemModel::hasChildren(parent);
	}
	bool setData(const QModelIndex& index, const QVariant& value, int role) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_int_voidp_voidp_int) { int __r = ((int(*)(long long, long long, void*, void*, int, int*))cppgo_virt_int_voidp_voidp_int)(__goId, 35191454575202219LL, (void*)&index, (void*)&value, role, &__h); if (__h) return __r; }
		return QAbstractItemModel::setData(index, value, role);
	}
	QVariant headerData(int section, Qt::Orientation orientation, int role) const override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_voidp_int_int_int) { void* __r = ((void*(*)(long long, long long, int, int, int, int*))cppgo_virt_voidp_int_int_int)(__goId, -3297140021896693688LL, section, (int)(orientation), role, &__h); if (__h) return *(QVariant*)__r; }
		return QAbstractItemModel::headerData(section, orientation, role);
	}
	bool setHeaderData(int section, Qt::Orientation orientation, const QVariant& value, int role) override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_int_int_int_voidp_int) { int __r = ((int(*)(long long, long long, int, int, void*, int, int*))cppgo_virt_int_int_int_voidp_int)(__goId, 4852200306590054850LL, section, (int)(orientation), (void*)&value, role, &__h); if (__h) return __r; }
		return QAbstractItemModel::setHeaderData(section, orientation, value, role);
	}
	bool clearItemData(const QModelIndex& index) override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 8914190500520542409LL, (void*)&index, &__h); if (__h) return __r; }
		return QAbstractItemModel::clearItemData(index);
	}
	QList<QString> mimeTypes() const override {
		int __h = 0;
		if (__goOn[12] && cppgo_virtp_voidp) {
			long long __len = 0;
			void* __a = ((void*(*)(long long, long long, long long*, int*))cppgo_virtp_voidp)(__goId, 5464467312394730418LL, &__len, &__h);
			if (__h) {
				QList<QString> __r; __r.reserve(__len);
				char** __aa = (char**)__a;
				for (long long __i = 0; __i < __len; __i++) { __r.append(QString::fromUtf8(__aa[__i])); free(__aa[__i]); }
				if (__a) free(__a);
				return __r;
			}
		}
		return QAbstractItemModel::mimeTypes();
	}
	QMimeData* mimeData(const QList<QModelIndex>& indexes) const override {
		int __h = 0;
		std::vector<void*> indexes_v; indexes_v.reserve(indexes.size());
		for (long long __i = 0; __i < indexes.size(); __i++) indexes_v.push_back((void*)&indexes[__i]);
		if (__goOn[13] && cppgo_virt_voidp_voidpp_longlong) { void* __r = ((void*(*)(long long, long long, void**, long long, int*))cppgo_virt_voidp_voidpp_longlong)(__goId, 5175297091024974749LL, indexes_v.data(), (long long)indexes.size(), &__h); if (__h) return (QMimeData*)__r; }
		return QAbstractItemModel::mimeData(indexes);
	}
	bool canDropMimeData(const QMimeData* data, Qt::DropAction action, int row, int column, const QModelIndex& parent) const override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_int_voidp_int_int_int_voidp) { int __r = ((int(*)(long long, long long, void*, int, int, int, void*, int*))cppgo_virt_int_voidp_int_int_int_voidp)(__goId, 731539533644753052LL, (void*)data, (int)(action), row, column, (void*)&parent, &__h); if (__h) return __r; }
		return QAbstractItemModel::canDropMimeData(data, action, row, column, parent);
	}
	bool dropMimeData(const QMimeData* data, Qt::DropAction action, int row, int column, const QModelIndex& parent) override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_int_voidp_int_int_int_voidp) { int __r = ((int(*)(long long, long long, void*, int, int, int, void*, int*))cppgo_virt_int_voidp_int_int_int_voidp)(__goId, -9049994684305682212LL, (void*)data, (int)(action), row, column, (void*)&parent, &__h); if (__h) return __r; }
		return QAbstractItemModel::dropMimeData(data, action, row, column, parent);
	}
	QFlags<Qt::DropAction> supportedDropActions() const override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -7784642176186821205LL, &__h); if (__h) return QFlags<Qt::DropAction>(__r); }
		return QAbstractItemModel::supportedDropActions();
	}
	QFlags<Qt::DropAction> supportedDragActions() const override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, 3421601386088570762LL, &__h); if (__h) return QFlags<Qt::DropAction>(__r); }
		return QAbstractItemModel::supportedDragActions();
	}
	bool insertRows(int row, int count, const QModelIndex& parent) override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_int_int_int_voidp) { int __r = ((int(*)(long long, long long, int, int, void*, int*))cppgo_virt_int_int_int_voidp)(__goId, -2775729825962395603LL, row, count, (void*)&parent, &__h); if (__h) return __r; }
		return QAbstractItemModel::insertRows(row, count, parent);
	}
	bool insertColumns(int column, int count, const QModelIndex& parent) override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_int_int_int_voidp) { int __r = ((int(*)(long long, long long, int, int, void*, int*))cppgo_virt_int_int_int_voidp)(__goId, -8973293846230792237LL, column, count, (void*)&parent, &__h); if (__h) return __r; }
		return QAbstractItemModel::insertColumns(column, count, parent);
	}
	bool removeRows(int row, int count, const QModelIndex& parent) override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_int_int_int_voidp) { int __r = ((int(*)(long long, long long, int, int, void*, int*))cppgo_virt_int_int_int_voidp)(__goId, 5104072549182222656LL, row, count, (void*)&parent, &__h); if (__h) return __r; }
		return QAbstractItemModel::removeRows(row, count, parent);
	}
	bool removeColumns(int column, int count, const QModelIndex& parent) override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_int_int_int_voidp) { int __r = ((int(*)(long long, long long, int, int, void*, int*))cppgo_virt_int_int_int_voidp)(__goId, 2078548737056540420LL, column, count, (void*)&parent, &__h); if (__h) return __r; }
		return QAbstractItemModel::removeColumns(column, count, parent);
	}
	bool moveRows(const QModelIndex& sourceParent, int sourceRow, int count, const QModelIndex& destinationParent, int destinationChild) override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_int_voidp_int_int_voidp_int) { int __r = ((int(*)(long long, long long, void*, int, int, void*, int, int*))cppgo_virt_int_voidp_int_int_voidp_int)(__goId, -6074819356947405415LL, (void*)&sourceParent, sourceRow, count, (void*)&destinationParent, destinationChild, &__h); if (__h) return __r; }
		return QAbstractItemModel::moveRows(sourceParent, sourceRow, count, destinationParent, destinationChild);
	}
	bool moveColumns(const QModelIndex& sourceParent, int sourceColumn, int count, const QModelIndex& destinationParent, int destinationChild) override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_int_voidp_int_int_voidp_int) { int __r = ((int(*)(long long, long long, void*, int, int, void*, int, int*))cppgo_virt_int_voidp_int_int_voidp_int)(__goId, -6321007591388309905LL, (void*)&sourceParent, sourceColumn, count, (void*)&destinationParent, destinationChild, &__h); if (__h) return __r; }
		return QAbstractItemModel::moveColumns(sourceParent, sourceColumn, count, destinationParent, destinationChild);
	}
	void fetchMore(const QModelIndex& parent) override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3553294593611130806LL, (void*)&parent, &__h); if (__h) return; }
		QAbstractItemModel::fetchMore(parent);
	}
	bool canFetchMore(const QModelIndex& parent) const override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 8639262066473964150LL, (void*)&parent, &__h); if (__h) return __r; }
		return QAbstractItemModel::canFetchMore(parent);
	}
	QFlags<Qt::ItemFlag> flags(const QModelIndex& index) const override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1703382833545304812LL, (void*)&index, &__h); if (__h) return QFlags<Qt::ItemFlag>(__r); }
		return QAbstractItemModel::flags(index);
	}
	void sort(int column, Qt::SortOrder order) override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_void_int_int) { ((void(*)(long long, long long, int, int, int*))cppgo_virt_void_int_int)(__goId, 6912441694453382161LL, column, (int)(order), &__h); if (__h) return; }
		QAbstractItemModel::sort(column, order);
	}
	QModelIndex buddy(const QModelIndex& index) const override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 6112860112248786499LL, (void*)&index, &__h); if (__h) return *(QModelIndex*)__r; }
		return QAbstractItemModel::buddy(index);
	}
	QList<QModelIndex> match(const QModelIndex& start, int role, const QVariant& value, int hits, QFlags<Qt::MatchFlag> flags) const override {
		int __h = 0;
		if (__goOn[29] && cppgo_virtp_voidp_voidp_int_voidp_int_int) {
			long long __len = 0;
			void* __a = ((void*(*)(long long, long long, void*, int, void*, int, int, long long*, int*))cppgo_virtp_voidp_voidp_int_voidp_int_int)(__goId, -4341501217981328394LL, (void*)&start, role, (void*)&value, hits, (int)(flags), &__len, &__h);
			if (__h) {
				QList<QModelIndex> __r; __r.reserve(__len);
				void** __aa = (void**)__a;
				for (long long __i = 0; __i < __len; __i++) __r.append(*(QModelIndex*)__aa[__i]);
				if (__a) free(__a);
				return __r;
			}
		}
		return QAbstractItemModel::match(start, role, value, hits, flags);
	}
	QSize span(const QModelIndex& index) const override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, -8395342916587040543LL, (void*)&index, &__h); if (__h) return *(QSize*)__r; }
		return QAbstractItemModel::span(index);
	}
	void multiData(const QModelIndex& index, QModelRoleDataSpan roleDataSpan) const override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_void_voidp_voidp) { ((void(*)(long long, long long, void*, void*, int*))cppgo_virt_void_voidp_voidp)(__goId, 731311060944096642LL, (void*)&index, (void*)&roleDataSpan, &__h); if (__h) return; }
		QAbstractItemModel::multiData(index, roleDataSpan);
	}
	bool submit() override {
		int __h = 0;
		if (__goOn[32] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -3788331537442685179LL, &__h); if (__h) return __r; }
		return QAbstractItemModel::submit();
	}
	void revert() override {
		int __h = 0;
		if (__goOn[33] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, -1522951607536117149LL, &__h); if (__h) return; }
		QAbstractItemModel::revert();
	}
	void resetInternalData() override {
		int __h = 0;
		if (__goOn[34] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, -4510675770784051183LL, &__h); if (__h) return; }
		QAbstractItemModel::resetInternalData();
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[35] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QAbstractItemModel::event(event);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[36] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QAbstractItemModel::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[37] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QAbstractItemModel::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[38] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QAbstractItemModel::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[39] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QAbstractItemModel::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[40] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QAbstractItemModel::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[41] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QAbstractItemModel::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QAbstractItemModel;
static std::mutex __connsMu_QAbstractItemModel;

extern "C" {

void* QAbstractItemModel_base(void* self) { return (void*)static_cast<QObject*>((QAbstractItemModel*)self); }
void QAbstractItemModel__set_override(void* self, int idx) { ((QAbstractItemModel__Go*)self)->__goOn[idx] = 1; }
void* QAbstractItemModel_fromQObject(void* o) { return (void*)qobject_cast<QAbstractItemModel*>((QObject*)o); }
char* QAbstractItemModel_tr(const char* s, long long* __out_len) {
	auto&& __r = QAbstractItemModel::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QAbstractItemModel_hasIndex(void* self, int row, int column) {
	auto&& __r = ((QAbstractItemModel*)self)->hasIndex(row, column);
	return __r;
}
void* QAbstractItemModel_index(void* self, int row, int column, void* parent) {
	auto&& __r = ((QAbstractItemModel*)self)->index(row, column, *(QModelIndex*)parent);
	return new QModelIndex(__r);
}
void* QAbstractItemModel_parent(void* self, void* child) {
	auto&& __r = ((QAbstractItemModel*)self)->parent(*(QModelIndex*)child);
	return new QModelIndex(__r);
}
void* QAbstractItemModel_sibling(void* self, int row, int column, void* idx) {
	auto&& __r = ((QAbstractItemModel*)self)->sibling(row, column, *(QModelIndex*)idx);
	return new QModelIndex(__r);
}
int QAbstractItemModel_rowCount(void* self, void* parent) {
	auto&& __r = ((QAbstractItemModel*)self)->rowCount(*(QModelIndex*)parent);
	return __r;
}
int QAbstractItemModel_columnCount(void* self, void* parent) {
	auto&& __r = ((QAbstractItemModel*)self)->columnCount(*(QModelIndex*)parent);
	return __r;
}
int QAbstractItemModel_hasChildren(void* self, void* parent) {
	auto&& __r = ((QAbstractItemModel*)self)->hasChildren(*(QModelIndex*)parent);
	return __r;
}
void* QAbstractItemModel_data(void* self, void* index, int role) {
	auto&& __r = ((QAbstractItemModel*)self)->data(*(QModelIndex*)index, role);
	return new QVariant(__r);
}
int QAbstractItemModel_setData(void* self, void* index, void* value, int role) {
	auto&& __r = ((QAbstractItemModel*)self)->setData(*(QModelIndex*)index, *(QVariant*)value, role);
	return __r;
}
void* QAbstractItemModel_headerData(void* self, int section, int orientation, int role) {
	auto&& __r = ((QAbstractItemModel*)self)->headerData(section, (Qt::Orientation)(orientation), role);
	return new QVariant(__r);
}
int QAbstractItemModel_setHeaderData(void* self, int section, int orientation, void* value, int role) {
	auto&& __r = ((QAbstractItemModel*)self)->setHeaderData(section, (Qt::Orientation)(orientation), *(QVariant*)value, role);
	return __r;
}
int* QAbstractItemModel_itemData(void* self, void* index, void** __out_vals, long long* __out_len) {
	auto&& __r = ((QAbstractItemModel*)self)->itemData(*(QModelIndex*)index);
	*__out_len = __r.size();
	int* __keys = (int*)malloc(sizeof(int) * (__r.size() > 0 ? __r.size() : 1));
	void** __vals = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto __it = __r.cbegin(); __it != __r.cend(); ++__it) { __keys[__i] = __it.key(); __vals[__i] = new QVariant(__it.value()); __i++; }
	*__out_vals = (void*)__vals;
	return __keys;
}
int QAbstractItemModel_setItemData(void* self, void* index, int* roles_keys, void* roles_vals, long long roles_len) {
	QMap<int, QVariant> roles_map;
	void** roles_vv = (void**)roles_vals;
	for (long long __i = 0; __i < roles_len; __i++) roles_map.insert(roles_keys[__i], *(QVariant*)roles_vv[__i]);
	auto&& __r = ((QAbstractItemModel*)self)->setItemData(*(QModelIndex*)index, roles_map);
	return __r;
}
int QAbstractItemModel_clearItemData(void* self, void* index) {
	auto&& __r = ((QAbstractItemModel*)self)->clearItemData(*(QModelIndex*)index);
	return __r;
}
char** QAbstractItemModel_mimeTypes(void* self, long long* __out_len) {
	auto&& __r = ((QAbstractItemModel*)self)->mimeTypes();
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
void* QAbstractItemModel_mimeData(void* self, void** indexes, long long indexes_len) {
	QList<QModelIndex> indexes_list; indexes_list.reserve(indexes_len);
	for (long long __i = 0; __i < indexes_len; __i++) indexes_list.append(*(QModelIndex*)indexes[__i]);
	auto&& __r = ((QAbstractItemModel*)self)->mimeData(indexes_list);
	return (void*)(__r);
}
int QAbstractItemModel_canDropMimeData(void* self, void* data, int action, int row, int column, void* parent) {
	auto&& __r = ((QAbstractItemModel*)self)->canDropMimeData((QMimeData*)data, (Qt::DropAction)(action), row, column, *(QModelIndex*)parent);
	return __r;
}
int QAbstractItemModel_dropMimeData(void* self, void* data, int action, int row, int column, void* parent) {
	auto&& __r = ((QAbstractItemModel*)self)->dropMimeData((QMimeData*)data, (Qt::DropAction)(action), row, column, *(QModelIndex*)parent);
	return __r;
}
int QAbstractItemModel_supportedDropActions(void* self) {
	auto&& __r = ((QAbstractItemModel*)self)->supportedDropActions();
	return (int)(__r);
}
int QAbstractItemModel_supportedDragActions(void* self) {
	auto&& __r = ((QAbstractItemModel*)self)->supportedDragActions();
	return (int)(__r);
}
int QAbstractItemModel_insertRows(void* self, int row, int count, void* parent) {
	auto&& __r = ((QAbstractItemModel*)self)->insertRows(row, count, *(QModelIndex*)parent);
	return __r;
}
int QAbstractItemModel_insertColumns(void* self, int column, int count, void* parent) {
	auto&& __r = ((QAbstractItemModel*)self)->insertColumns(column, count, *(QModelIndex*)parent);
	return __r;
}
int QAbstractItemModel_removeRows(void* self, int row, int count, void* parent) {
	auto&& __r = ((QAbstractItemModel*)self)->removeRows(row, count, *(QModelIndex*)parent);
	return __r;
}
int QAbstractItemModel_removeColumns(void* self, int column, int count, void* parent) {
	auto&& __r = ((QAbstractItemModel*)self)->removeColumns(column, count, *(QModelIndex*)parent);
	return __r;
}
int QAbstractItemModel_moveRows(void* self, void* sourceParent, int sourceRow, int count, void* destinationParent, int destinationChild) {
	auto&& __r = ((QAbstractItemModel*)self)->moveRows(*(QModelIndex*)sourceParent, sourceRow, count, *(QModelIndex*)destinationParent, destinationChild);
	return __r;
}
int QAbstractItemModel_moveColumns(void* self, void* sourceParent, int sourceColumn, int count, void* destinationParent, int destinationChild) {
	auto&& __r = ((QAbstractItemModel*)self)->moveColumns(*(QModelIndex*)sourceParent, sourceColumn, count, *(QModelIndex*)destinationParent, destinationChild);
	return __r;
}
int QAbstractItemModel_insertRow(void* self, int row) {
	auto&& __r = ((QAbstractItemModel*)self)->insertRow(row);
	return __r;
}
int QAbstractItemModel_insertColumn(void* self, int column) {
	auto&& __r = ((QAbstractItemModel*)self)->insertColumn(column);
	return __r;
}
int QAbstractItemModel_removeRow(void* self, int row) {
	auto&& __r = ((QAbstractItemModel*)self)->removeRow(row);
	return __r;
}
int QAbstractItemModel_removeColumn(void* self, int column) {
	auto&& __r = ((QAbstractItemModel*)self)->removeColumn(column);
	return __r;
}
int QAbstractItemModel_moveRow(void* self, void* sourceParent, int sourceRow, void* destinationParent, int destinationChild) {
	auto&& __r = ((QAbstractItemModel*)self)->moveRow(*(QModelIndex*)sourceParent, sourceRow, *(QModelIndex*)destinationParent, destinationChild);
	return __r;
}
int QAbstractItemModel_moveColumn(void* self, void* sourceParent, int sourceColumn, void* destinationParent, int destinationChild) {
	auto&& __r = ((QAbstractItemModel*)self)->moveColumn(*(QModelIndex*)sourceParent, sourceColumn, *(QModelIndex*)destinationParent, destinationChild);
	return __r;
}
void QAbstractItemModel_fetchMore(void* self, void* parent) {
	((QAbstractItemModel*)self)->fetchMore(*(QModelIndex*)parent);
}
int QAbstractItemModel_canFetchMore(void* self, void* parent) {
	auto&& __r = ((QAbstractItemModel*)self)->canFetchMore(*(QModelIndex*)parent);
	return __r;
}
int QAbstractItemModel_flags(void* self, void* index) {
	auto&& __r = ((QAbstractItemModel*)self)->flags(*(QModelIndex*)index);
	return (int)(__r);
}
void QAbstractItemModel_sort(void* self, int column, int order) {
	((QAbstractItemModel*)self)->sort(column, (Qt::SortOrder)(order));
}
void* QAbstractItemModel_buddy(void* self, void* index) {
	auto&& __r = ((QAbstractItemModel*)self)->buddy(*(QModelIndex*)index);
	return new QModelIndex(__r);
}
void** QAbstractItemModel_match(void* self, void* start, int role, void* value, int hits, int flags, long long* __out_len) {
	auto&& __r = ((QAbstractItemModel*)self)->match(*(QModelIndex*)start, role, *(QVariant*)value, hits, QFlags<Qt::MatchFlag>(flags));
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QModelIndex(__e);
	return __arr;
}
void* QAbstractItemModel_span(void* self, void* index) {
	auto&& __r = ((QAbstractItemModel*)self)->span(*(QModelIndex*)index);
	return new QSize(__r);
}
int* QAbstractItemModel_roleNames(void* self, void** __out_vals, long long* __out_len) {
	auto&& __r = ((QAbstractItemModel*)self)->roleNames();
	*__out_len = __r.size();
	int* __keys = (int*)malloc(sizeof(int) * (__r.size() > 0 ? __r.size() : 1));
	void** __vals = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto __it = __r.cbegin(); __it != __r.cend(); ++__it) { __keys[__i] = __it.key(); __vals[__i] = new QByteArray(__it.value()); __i++; }
	*__out_vals = (void*)__vals;
	return __keys;
}
int QAbstractItemModel_checkIndex(void* self, void* index) {
	auto&& __r = ((QAbstractItemModel*)self)->checkIndex(*(QModelIndex*)index);
	return __r;
}
void QAbstractItemModel_multiData(void* self, void* index, void* roleDataSpan) {
	((QAbstractItemModel*)self)->multiData(*(QModelIndex*)index, *(QModelRoleDataSpan*)roleDataSpan);
}
void QAbstractItemModel_dataChanged(void* self, void* topLeft, void* bottomRight) {
	((QAbstractItemModel*)self)->dataChanged(*(QModelIndex*)topLeft, *(QModelIndex*)bottomRight);
}
void QAbstractItemModel_headerDataChanged(void* self, int orientation, int first, int last) {
	((QAbstractItemModel*)self)->headerDataChanged((Qt::Orientation)(orientation), first, last);
}
void QAbstractItemModel_layoutChanged(void* self) {
	((QAbstractItemModel*)self)->layoutChanged();
}
void QAbstractItemModel_layoutAboutToBeChanged(void* self) {
	((QAbstractItemModel*)self)->layoutAboutToBeChanged();
}
int QAbstractItemModel_submit(void* self) {
	auto&& __r = ((QAbstractItemModel*)self)->submit();
	return __r;
}
void QAbstractItemModel_revert(void* self) {
	((QAbstractItemModel*)self)->revert();
}
void* QAbstractItemModel_createIndex(void* self, int row, int column) {
	auto&& __r = ((QAbstractItemModel__Go*)self)->createIndex(row, column);
	return new QModelIndex(__r);
}
void* QAbstractItemModel_createIndex2(void* self, int row, int column, unsigned long long id) {
	auto&& __r = ((QAbstractItemModel__Go*)self)->createIndex(row, column, id);
	return new QModelIndex(__r);
}
void QAbstractItemModel_encodeData(void* self, void** indexes, long long indexes_len, void* stream) {
	QList<QModelIndex> indexes_list; indexes_list.reserve(indexes_len);
	for (long long __i = 0; __i < indexes_len; __i++) indexes_list.append(*(QModelIndex*)indexes[__i]);
	((QAbstractItemModel__Go*)self)->encodeData(indexes_list, *(QDataStream*)stream);
}
int QAbstractItemModel_decodeData(void* self, int row, int column, void* parent, void* stream) {
	auto&& __r = ((QAbstractItemModel__Go*)self)->decodeData(row, column, *(QModelIndex*)parent, *(QDataStream*)stream);
	return __r;
}
void QAbstractItemModel_beginInsertRows(void* self, void* parent, int first, int last) {
	((QAbstractItemModel__Go*)self)->beginInsertRows(*(QModelIndex*)parent, first, last);
}
void QAbstractItemModel_endInsertRows(void* self) {
	((QAbstractItemModel__Go*)self)->endInsertRows();
}
void QAbstractItemModel_beginRemoveRows(void* self, void* parent, int first, int last) {
	((QAbstractItemModel__Go*)self)->beginRemoveRows(*(QModelIndex*)parent, first, last);
}
void QAbstractItemModel_endRemoveRows(void* self) {
	((QAbstractItemModel__Go*)self)->endRemoveRows();
}
int QAbstractItemModel_beginMoveRows(void* self, void* sourceParent, int sourceFirst, int sourceLast, void* destinationParent, int destinationRow) {
	auto&& __r = ((QAbstractItemModel__Go*)self)->beginMoveRows(*(QModelIndex*)sourceParent, sourceFirst, sourceLast, *(QModelIndex*)destinationParent, destinationRow);
	return __r;
}
void QAbstractItemModel_endMoveRows(void* self) {
	((QAbstractItemModel__Go*)self)->endMoveRows();
}
void QAbstractItemModel_beginInsertColumns(void* self, void* parent, int first, int last) {
	((QAbstractItemModel__Go*)self)->beginInsertColumns(*(QModelIndex*)parent, first, last);
}
void QAbstractItemModel_endInsertColumns(void* self) {
	((QAbstractItemModel__Go*)self)->endInsertColumns();
}
void QAbstractItemModel_beginRemoveColumns(void* self, void* parent, int first, int last) {
	((QAbstractItemModel__Go*)self)->beginRemoveColumns(*(QModelIndex*)parent, first, last);
}
void QAbstractItemModel_endRemoveColumns(void* self) {
	((QAbstractItemModel__Go*)self)->endRemoveColumns();
}
int QAbstractItemModel_beginMoveColumns(void* self, void* sourceParent, int sourceFirst, int sourceLast, void* destinationParent, int destinationColumn) {
	auto&& __r = ((QAbstractItemModel__Go*)self)->beginMoveColumns(*(QModelIndex*)sourceParent, sourceFirst, sourceLast, *(QModelIndex*)destinationParent, destinationColumn);
	return __r;
}
void QAbstractItemModel_endMoveColumns(void* self) {
	((QAbstractItemModel__Go*)self)->endMoveColumns();
}
void QAbstractItemModel_beginResetModel(void* self) {
	((QAbstractItemModel__Go*)self)->beginResetModel();
}
void QAbstractItemModel_endResetModel(void* self) {
	((QAbstractItemModel__Go*)self)->endResetModel();
}
void QAbstractItemModel_changePersistentIndex(void* self, void* from, void* to) {
	((QAbstractItemModel__Go*)self)->changePersistentIndex(*(QModelIndex*)from, *(QModelIndex*)to);
}
void QAbstractItemModel_changePersistentIndexList(void* self, void** from, long long from_len, void** to, long long to_len) {
	QList<QModelIndex> from_list; from_list.reserve(from_len);
	for (long long __i = 0; __i < from_len; __i++) from_list.append(*(QModelIndex*)from[__i]);
	QList<QModelIndex> to_list; to_list.reserve(to_len);
	for (long long __i = 0; __i < to_len; __i++) to_list.append(*(QModelIndex*)to[__i]);
	((QAbstractItemModel__Go*)self)->changePersistentIndexList(from_list, to_list);
}
void** QAbstractItemModel_persistentIndexList(void* self, long long* __out_len) {
	auto&& __r = ((QAbstractItemModel__Go*)self)->persistentIndexList();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QModelIndex(__e);
	return __arr;
}
void* QAbstractItemModel_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QAbstractItemModel::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QAbstractItemModel_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QAbstractItemModel::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QAbstractItemModel_sender(void* self) {
	auto&& __r = ((QAbstractItemModel__Go*)self)->sender();
	return (void*)(__r);
}
int QAbstractItemModel_senderSignalIndex(void* self) {
	auto&& __r = ((QAbstractItemModel__Go*)self)->senderSignalIndex();
	return __r;
}
int QAbstractItemModel_receivers(void* self, const char* signal) {
	auto&& __r = ((QAbstractItemModel__Go*)self)->receivers(signal);
	return __r;
}
int QAbstractItemModel_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QAbstractItemModel__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QAbstractItemModel_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QAbstractItemModel::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QAbstractItemModel_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QAbstractItemModel::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QAbstractItemModel_hasIndex2(void* self, int row, int column, void* parent) {
	auto&& __r = ((QAbstractItemModel*)self)->hasIndex(row, column, *(QModelIndex*)parent);
	return __r;
}
int QAbstractItemModel_insertRow2(void* self, int row, void* parent) {
	auto&& __r = ((QAbstractItemModel*)self)->insertRow(row, *(QModelIndex*)parent);
	return __r;
}
int QAbstractItemModel_insertColumn2(void* self, int column, void* parent) {
	auto&& __r = ((QAbstractItemModel*)self)->insertColumn(column, *(QModelIndex*)parent);
	return __r;
}
int QAbstractItemModel_removeRow2(void* self, int row, void* parent) {
	auto&& __r = ((QAbstractItemModel*)self)->removeRow(row, *(QModelIndex*)parent);
	return __r;
}
int QAbstractItemModel_removeColumn2(void* self, int column, void* parent) {
	auto&& __r = ((QAbstractItemModel*)self)->removeColumn(column, *(QModelIndex*)parent);
	return __r;
}
int QAbstractItemModel_checkIndex2(void* self, void* index, int options) {
	auto&& __r = ((QAbstractItemModel*)self)->checkIndex(*(QModelIndex*)index, QFlags<QAbstractItemModel::CheckIndexOption>(options));
	return __r;
}
void QAbstractItemModel_dataChanged2(void* self, void* topLeft, void* bottomRight, int* roles, long long roles_len) {
	QList<int> roles_list; roles_list.reserve(roles_len);
	for (long long __i = 0; __i < roles_len; __i++) roles_list.append(roles[__i]);
	((QAbstractItemModel*)self)->dataChanged(*(QModelIndex*)topLeft, *(QModelIndex*)bottomRight, roles_list);
}
void QAbstractItemModel_layoutChanged2(void* self, void** parents, long long parents_len) {
	QList<QPersistentModelIndex> parents_list; parents_list.reserve(parents_len);
	for (long long __i = 0; __i < parents_len; __i++) parents_list.append(*(QPersistentModelIndex*)parents[__i]);
	((QAbstractItemModel*)self)->layoutChanged(parents_list);
}
void QAbstractItemModel_layoutChanged3(void* self, void** parents, long long parents_len, int hint) {
	QList<QPersistentModelIndex> parents_list; parents_list.reserve(parents_len);
	for (long long __i = 0; __i < parents_len; __i++) parents_list.append(*(QPersistentModelIndex*)parents[__i]);
	((QAbstractItemModel*)self)->layoutChanged(parents_list, (QAbstractItemModel::LayoutChangeHint)(hint));
}
void QAbstractItemModel_layoutAboutToBeChanged2(void* self, void** parents, long long parents_len) {
	QList<QPersistentModelIndex> parents_list; parents_list.reserve(parents_len);
	for (long long __i = 0; __i < parents_len; __i++) parents_list.append(*(QPersistentModelIndex*)parents[__i]);
	((QAbstractItemModel*)self)->layoutAboutToBeChanged(parents_list);
}
void QAbstractItemModel_layoutAboutToBeChanged3(void* self, void** parents, long long parents_len, int hint) {
	QList<QPersistentModelIndex> parents_list; parents_list.reserve(parents_len);
	for (long long __i = 0; __i < parents_len; __i++) parents_list.append(*(QPersistentModelIndex*)parents[__i]);
	((QAbstractItemModel*)self)->layoutAboutToBeChanged(parents_list, (QAbstractItemModel::LayoutChangeHint)(hint));
}
void* QAbstractItemModel_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QAbstractItemModel::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QAbstractItemModel_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QAbstractItemModel::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QAbstractItemModel_sibling_default(void* self, int row, int column, void* idx) {
	auto&& __r = ((QAbstractItemModel*)self)->QAbstractItemModel::sibling(row, column, *(QModelIndex*)idx);
	return new QModelIndex(__r);
}
int QAbstractItemModel_hasChildren_default(void* self, void* parent) {
	auto&& __r = ((QAbstractItemModel*)self)->QAbstractItemModel::hasChildren(*(QModelIndex*)parent);
	return __r;
}
int QAbstractItemModel_setData_default(void* self, void* index, void* value, int role) {
	auto&& __r = ((QAbstractItemModel*)self)->QAbstractItemModel::setData(*(QModelIndex*)index, *(QVariant*)value, role);
	return __r;
}
void* QAbstractItemModel_headerData_default(void* self, int section, int orientation, int role) {
	auto&& __r = ((QAbstractItemModel*)self)->QAbstractItemModel::headerData(section, (Qt::Orientation)(orientation), role);
	return new QVariant(__r);
}
int QAbstractItemModel_setHeaderData_default(void* self, int section, int orientation, void* value, int role) {
	auto&& __r = ((QAbstractItemModel*)self)->QAbstractItemModel::setHeaderData(section, (Qt::Orientation)(orientation), *(QVariant*)value, role);
	return __r;
}
int QAbstractItemModel_clearItemData_default(void* self, void* index) {
	auto&& __r = ((QAbstractItemModel*)self)->QAbstractItemModel::clearItemData(*(QModelIndex*)index);
	return __r;
}
char** QAbstractItemModel_mimeTypes_default(void* self, long long* __out_len) {
	auto&& __r = ((QAbstractItemModel*)self)->QAbstractItemModel::mimeTypes();
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
void* QAbstractItemModel_mimeData_default(void* self, void** indexes, long long indexes_len) {
	QList<QModelIndex> indexes_list; indexes_list.reserve(indexes_len);
	for (long long __i = 0; __i < indexes_len; __i++) indexes_list.append(*(QModelIndex*)indexes[__i]);
	auto&& __r = ((QAbstractItemModel*)self)->QAbstractItemModel::mimeData(indexes_list);
	return (void*)(__r);
}
int QAbstractItemModel_canDropMimeData_default(void* self, void* data, int action, int row, int column, void* parent) {
	auto&& __r = ((QAbstractItemModel*)self)->QAbstractItemModel::canDropMimeData((QMimeData*)data, (Qt::DropAction)(action), row, column, *(QModelIndex*)parent);
	return __r;
}
int QAbstractItemModel_dropMimeData_default(void* self, void* data, int action, int row, int column, void* parent) {
	auto&& __r = ((QAbstractItemModel*)self)->QAbstractItemModel::dropMimeData((QMimeData*)data, (Qt::DropAction)(action), row, column, *(QModelIndex*)parent);
	return __r;
}
int QAbstractItemModel_supportedDropActions_default(void* self) {
	auto&& __r = ((QAbstractItemModel*)self)->QAbstractItemModel::supportedDropActions();
	return (int)(__r);
}
int QAbstractItemModel_supportedDragActions_default(void* self) {
	auto&& __r = ((QAbstractItemModel*)self)->QAbstractItemModel::supportedDragActions();
	return (int)(__r);
}
int QAbstractItemModel_insertRows_default(void* self, int row, int count, void* parent) {
	auto&& __r = ((QAbstractItemModel*)self)->QAbstractItemModel::insertRows(row, count, *(QModelIndex*)parent);
	return __r;
}
int QAbstractItemModel_insertColumns_default(void* self, int column, int count, void* parent) {
	auto&& __r = ((QAbstractItemModel*)self)->QAbstractItemModel::insertColumns(column, count, *(QModelIndex*)parent);
	return __r;
}
int QAbstractItemModel_removeRows_default(void* self, int row, int count, void* parent) {
	auto&& __r = ((QAbstractItemModel*)self)->QAbstractItemModel::removeRows(row, count, *(QModelIndex*)parent);
	return __r;
}
int QAbstractItemModel_removeColumns_default(void* self, int column, int count, void* parent) {
	auto&& __r = ((QAbstractItemModel*)self)->QAbstractItemModel::removeColumns(column, count, *(QModelIndex*)parent);
	return __r;
}
int QAbstractItemModel_moveRows_default(void* self, void* sourceParent, int sourceRow, int count, void* destinationParent, int destinationChild) {
	auto&& __r = ((QAbstractItemModel*)self)->QAbstractItemModel::moveRows(*(QModelIndex*)sourceParent, sourceRow, count, *(QModelIndex*)destinationParent, destinationChild);
	return __r;
}
int QAbstractItemModel_moveColumns_default(void* self, void* sourceParent, int sourceColumn, int count, void* destinationParent, int destinationChild) {
	auto&& __r = ((QAbstractItemModel*)self)->QAbstractItemModel::moveColumns(*(QModelIndex*)sourceParent, sourceColumn, count, *(QModelIndex*)destinationParent, destinationChild);
	return __r;
}
void QAbstractItemModel_fetchMore_default(void* self, void* parent) {
	((QAbstractItemModel*)self)->QAbstractItemModel::fetchMore(*(QModelIndex*)parent);
}
int QAbstractItemModel_canFetchMore_default(void* self, void* parent) {
	auto&& __r = ((QAbstractItemModel*)self)->QAbstractItemModel::canFetchMore(*(QModelIndex*)parent);
	return __r;
}
int QAbstractItemModel_flags_default(void* self, void* index) {
	auto&& __r = ((QAbstractItemModel*)self)->QAbstractItemModel::flags(*(QModelIndex*)index);
	return (int)(__r);
}
void QAbstractItemModel_sort_default(void* self, int column, int order) {
	((QAbstractItemModel*)self)->QAbstractItemModel::sort(column, (Qt::SortOrder)(order));
}
void* QAbstractItemModel_buddy_default(void* self, void* index) {
	auto&& __r = ((QAbstractItemModel*)self)->QAbstractItemModel::buddy(*(QModelIndex*)index);
	return new QModelIndex(__r);
}
void** QAbstractItemModel_match_default(void* self, void* start, int role, void* value, int hits, int flags, long long* __out_len) {
	auto&& __r = ((QAbstractItemModel*)self)->QAbstractItemModel::match(*(QModelIndex*)start, role, *(QVariant*)value, hits, QFlags<Qt::MatchFlag>(flags));
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QModelIndex(__e);
	return __arr;
}
void* QAbstractItemModel_span_default(void* self, void* index) {
	auto&& __r = ((QAbstractItemModel*)self)->QAbstractItemModel::span(*(QModelIndex*)index);
	return new QSize(__r);
}
void QAbstractItemModel_multiData_default(void* self, void* index, void* roleDataSpan) {
	((QAbstractItemModel*)self)->QAbstractItemModel::multiData(*(QModelIndex*)index, *(QModelRoleDataSpan*)roleDataSpan);
}
int QAbstractItemModel_submit_default(void* self) {
	auto&& __r = ((QAbstractItemModel*)self)->QAbstractItemModel::submit();
	return __r;
}
void QAbstractItemModel_revert_default(void* self) {
	((QAbstractItemModel*)self)->QAbstractItemModel::revert();
}
void* QAbstractItemModel_metaObject_default(void* self) {
	auto&& __r = ((QAbstractItemModel*)self)->QAbstractItemModel::metaObject();
	return (void*)(__r);
}
void QAbstractItemModel_resetInternalData_default(void* self) {
	((QAbstractItemModel__Go*)self)->cppgoBase_resetInternalData();
}
int QAbstractItemModel_event_default(void* self, void* event) {
	auto&& __r = ((QAbstractItemModel*)self)->QAbstractItemModel::event((QEvent*)event);
	return __r;
}
int QAbstractItemModel_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QAbstractItemModel*)self)->QAbstractItemModel::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QAbstractItemModel_timerEvent_default(void* self, void* event) {
	((QAbstractItemModel__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QAbstractItemModel_childEvent_default(void* self, void* event) {
	((QAbstractItemModel__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QAbstractItemModel_customEvent_default(void* self, void* event) {
	((QAbstractItemModel__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QAbstractItemModel_connectNotify_default(void* self, void* signal) {
	((QAbstractItemModel__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QAbstractItemModel_disconnectNotify_default(void* self, void* signal) {
	((QAbstractItemModel__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QAbstractItemModel_connect_headerDataChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractItemModel*)self, &QAbstractItemModel::headerDataChanged, (QAbstractItemModel*)self, [__id](Qt::Orientation orientation, int first, int last){
		if (cppgo_sig_void_int_int_int) ((void(*)(long long, int, int, int))cppgo_sig_void_int_int_int)(__id, (int)(orientation), first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractItemModel);
	__conns_QAbstractItemModel[__id] = __c;
}
void QAbstractItemModel_connect_layoutChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractItemModel*)self, &QAbstractItemModel::layoutChanged, (QAbstractItemModel*)self, [__id](const QList<QPersistentModelIndex>& parents, QAbstractItemModel::LayoutChangeHint hint){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractItemModel);
	__conns_QAbstractItemModel[__id] = __c;
}
void QAbstractItemModel_connect_layoutChanged2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractItemModel*)self, &QAbstractItemModel::layoutChanged, (QAbstractItemModel*)self, [__id](const QList<QPersistentModelIndex>& parents, QAbstractItemModel::LayoutChangeHint hint){
		std::vector<void*> parents_v; parents_v.reserve(parents.size());
		for (long long __i = 0; __i < parents.size(); __i++) parents_v.push_back((void*)&parents[__i]);
		if (cppgo_sig_void_voidpp_longlong_int) ((void(*)(long long, void**, long long, int))cppgo_sig_void_voidpp_longlong_int)(__id, parents_v.data(), (long long)parents.size(), (int)(hint));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractItemModel);
	__conns_QAbstractItemModel[__id] = __c;
}
void QAbstractItemModel_connect_layoutAboutToBeChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractItemModel*)self, &QAbstractItemModel::layoutAboutToBeChanged, (QAbstractItemModel*)self, [__id](const QList<QPersistentModelIndex>& parents, QAbstractItemModel::LayoutChangeHint hint){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractItemModel);
	__conns_QAbstractItemModel[__id] = __c;
}
void QAbstractItemModel_connect_layoutAboutToBeChanged2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractItemModel*)self, &QAbstractItemModel::layoutAboutToBeChanged, (QAbstractItemModel*)self, [__id](const QList<QPersistentModelIndex>& parents, QAbstractItemModel::LayoutChangeHint hint){
		std::vector<void*> parents_v; parents_v.reserve(parents.size());
		for (long long __i = 0; __i < parents.size(); __i++) parents_v.push_back((void*)&parents[__i]);
		if (cppgo_sig_void_voidpp_longlong_int) ((void(*)(long long, void**, long long, int))cppgo_sig_void_voidpp_longlong_int)(__id, parents_v.data(), (long long)parents.size(), (int)(hint));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractItemModel);
	__conns_QAbstractItemModel[__id] = __c;
}
void QAbstractItemModel_connect_rowsAboutToBeInserted(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractItemModel*)self, &QAbstractItemModel::rowsAboutToBeInserted, (QAbstractItemModel*)self, [__id](const QModelIndex& parent, int first, int last){
		if (cppgo_sig_void_voidp_int_int) ((void(*)(long long, void*, int, int))cppgo_sig_void_voidp_int_int)(__id, (void*)&parent, first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractItemModel);
	__conns_QAbstractItemModel[__id] = __c;
}
void QAbstractItemModel_connect_rowsInserted(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractItemModel*)self, &QAbstractItemModel::rowsInserted, (QAbstractItemModel*)self, [__id](const QModelIndex& parent, int first, int last){
		if (cppgo_sig_void_voidp_int_int) ((void(*)(long long, void*, int, int))cppgo_sig_void_voidp_int_int)(__id, (void*)&parent, first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractItemModel);
	__conns_QAbstractItemModel[__id] = __c;
}
void QAbstractItemModel_connect_rowsAboutToBeRemoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractItemModel*)self, &QAbstractItemModel::rowsAboutToBeRemoved, (QAbstractItemModel*)self, [__id](const QModelIndex& parent, int first, int last){
		if (cppgo_sig_void_voidp_int_int) ((void(*)(long long, void*, int, int))cppgo_sig_void_voidp_int_int)(__id, (void*)&parent, first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractItemModel);
	__conns_QAbstractItemModel[__id] = __c;
}
void QAbstractItemModel_connect_rowsRemoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractItemModel*)self, &QAbstractItemModel::rowsRemoved, (QAbstractItemModel*)self, [__id](const QModelIndex& parent, int first, int last){
		if (cppgo_sig_void_voidp_int_int) ((void(*)(long long, void*, int, int))cppgo_sig_void_voidp_int_int)(__id, (void*)&parent, first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractItemModel);
	__conns_QAbstractItemModel[__id] = __c;
}
void QAbstractItemModel_connect_columnsAboutToBeInserted(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractItemModel*)self, &QAbstractItemModel::columnsAboutToBeInserted, (QAbstractItemModel*)self, [__id](const QModelIndex& parent, int first, int last){
		if (cppgo_sig_void_voidp_int_int) ((void(*)(long long, void*, int, int))cppgo_sig_void_voidp_int_int)(__id, (void*)&parent, first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractItemModel);
	__conns_QAbstractItemModel[__id] = __c;
}
void QAbstractItemModel_connect_columnsInserted(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractItemModel*)self, &QAbstractItemModel::columnsInserted, (QAbstractItemModel*)self, [__id](const QModelIndex& parent, int first, int last){
		if (cppgo_sig_void_voidp_int_int) ((void(*)(long long, void*, int, int))cppgo_sig_void_voidp_int_int)(__id, (void*)&parent, first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractItemModel);
	__conns_QAbstractItemModel[__id] = __c;
}
void QAbstractItemModel_connect_columnsAboutToBeRemoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractItemModel*)self, &QAbstractItemModel::columnsAboutToBeRemoved, (QAbstractItemModel*)self, [__id](const QModelIndex& parent, int first, int last){
		if (cppgo_sig_void_voidp_int_int) ((void(*)(long long, void*, int, int))cppgo_sig_void_voidp_int_int)(__id, (void*)&parent, first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractItemModel);
	__conns_QAbstractItemModel[__id] = __c;
}
void QAbstractItemModel_connect_columnsRemoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractItemModel*)self, &QAbstractItemModel::columnsRemoved, (QAbstractItemModel*)self, [__id](const QModelIndex& parent, int first, int last){
		if (cppgo_sig_void_voidp_int_int) ((void(*)(long long, void*, int, int))cppgo_sig_void_voidp_int_int)(__id, (void*)&parent, first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractItemModel);
	__conns_QAbstractItemModel[__id] = __c;
}
void QAbstractItemModel_connect_modelAboutToBeReset(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractItemModel*)self, &QAbstractItemModel::modelAboutToBeReset, (QAbstractItemModel*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractItemModel);
	__conns_QAbstractItemModel[__id] = __c;
}
void QAbstractItemModel_connect_modelReset(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractItemModel*)self, &QAbstractItemModel::modelReset, (QAbstractItemModel*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractItemModel);
	__conns_QAbstractItemModel[__id] = __c;
}
void QAbstractItemModel_connect_rowsAboutToBeMoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractItemModel*)self, &QAbstractItemModel::rowsAboutToBeMoved, (QAbstractItemModel*)self, [__id](const QModelIndex& sourceParent, int sourceStart, int sourceEnd, const QModelIndex& destinationParent, int destinationRow){
		if (cppgo_sig_void_voidp_int_int_voidp_int) ((void(*)(long long, void*, int, int, void*, int))cppgo_sig_void_voidp_int_int_voidp_int)(__id, (void*)&sourceParent, sourceStart, sourceEnd, (void*)&destinationParent, destinationRow);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractItemModel);
	__conns_QAbstractItemModel[__id] = __c;
}
void QAbstractItemModel_connect_rowsMoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractItemModel*)self, &QAbstractItemModel::rowsMoved, (QAbstractItemModel*)self, [__id](const QModelIndex& sourceParent, int sourceStart, int sourceEnd, const QModelIndex& destinationParent, int destinationRow){
		if (cppgo_sig_void_voidp_int_int_voidp_int) ((void(*)(long long, void*, int, int, void*, int))cppgo_sig_void_voidp_int_int_voidp_int)(__id, (void*)&sourceParent, sourceStart, sourceEnd, (void*)&destinationParent, destinationRow);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractItemModel);
	__conns_QAbstractItemModel[__id] = __c;
}
void QAbstractItemModel_connect_columnsAboutToBeMoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractItemModel*)self, &QAbstractItemModel::columnsAboutToBeMoved, (QAbstractItemModel*)self, [__id](const QModelIndex& sourceParent, int sourceStart, int sourceEnd, const QModelIndex& destinationParent, int destinationColumn){
		if (cppgo_sig_void_voidp_int_int_voidp_int) ((void(*)(long long, void*, int, int, void*, int))cppgo_sig_void_voidp_int_int_voidp_int)(__id, (void*)&sourceParent, sourceStart, sourceEnd, (void*)&destinationParent, destinationColumn);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractItemModel);
	__conns_QAbstractItemModel[__id] = __c;
}
void QAbstractItemModel_connect_columnsMoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractItemModel*)self, &QAbstractItemModel::columnsMoved, (QAbstractItemModel*)self, [__id](const QModelIndex& sourceParent, int sourceStart, int sourceEnd, const QModelIndex& destinationParent, int destinationColumn){
		if (cppgo_sig_void_voidp_int_int_voidp_int) ((void(*)(long long, void*, int, int, void*, int))cppgo_sig_void_voidp_int_int_voidp_int)(__id, (void*)&sourceParent, sourceStart, sourceEnd, (void*)&destinationParent, destinationColumn);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractItemModel);
	__conns_QAbstractItemModel[__id] = __c;
}
void QAbstractItemModel_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractItemModel*)self, &QAbstractItemModel::destroyed, (QAbstractItemModel*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractItemModel);
	__conns_QAbstractItemModel[__id] = __c;
}
void QAbstractItemModel_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractItemModel*)self, &QAbstractItemModel::destroyed, (QAbstractItemModel*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractItemModel);
	__conns_QAbstractItemModel[__id] = __c;
}
void QAbstractItemModel_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractItemModel*)self, &QAbstractItemModel::objectNameChanged, (QAbstractItemModel*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractItemModel);
	__conns_QAbstractItemModel[__id] = __c;
}
void QAbstractItemModel__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QAbstractItemModel);
		auto __it = __conns_QAbstractItemModel.find(id);
		if (__it == __conns_QAbstractItemModel.end()) return;
		__c = __it->second;
		__conns_QAbstractItemModel.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QAbstractItemModel__Go_new(long long id, void* parent) { return new QAbstractItemModel__Go(id, (QObject*)parent); }
void QAbstractItemModel_delete(void* self) { delete (QAbstractItemModel*)self; }

}

#include <QtCore/qcalendar.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qstring.h>
#include <QtCore/qdatetime.h>
#include <QtCore/qlocale.h>
#include <QtCore/qstringview.h>
#include "gen_qcalendar.h"

extern "C" {

void* QCalendar_new() {
	return new QCalendar();
}
void* QCalendar_new2(int system) {
	return new QCalendar((QCalendar::System)(system));
}
void* QCalendar_new3(const char* name, long long name_len) {
	QAnyStringView name_qstr = QAnyStringView(name, name_len);
	return new QCalendar(name_qstr);
}
void* QCalendar_new4(void* id) {
	return new QCalendar(*(QCalendar::SystemId*)id);
}
int QCalendar_isValid(void* self) {
	auto&& __r = ((QCalendar*)self)->isValid();
	return __r;
}
int QCalendar_daysInMonth(void* self, int month) {
	auto&& __r = ((QCalendar*)self)->daysInMonth(month);
	return __r;
}
int QCalendar_daysInYear(void* self, int year) {
	auto&& __r = ((QCalendar*)self)->daysInYear(year);
	return __r;
}
int QCalendar_monthsInYear(void* self, int year) {
	auto&& __r = ((QCalendar*)self)->monthsInYear(year);
	return __r;
}
int QCalendar_isDateValid(void* self, int year, int month, int day) {
	auto&& __r = ((QCalendar*)self)->isDateValid(year, month, day);
	return __r;
}
int QCalendar_isLeapYear(void* self, int year) {
	auto&& __r = ((QCalendar*)self)->isLeapYear(year);
	return __r;
}
int QCalendar_isGregorian(void* self) {
	auto&& __r = ((QCalendar*)self)->isGregorian();
	return __r;
}
int QCalendar_isLunar(void* self) {
	auto&& __r = ((QCalendar*)self)->isLunar();
	return __r;
}
int QCalendar_isLuniSolar(void* self) {
	auto&& __r = ((QCalendar*)self)->isLuniSolar();
	return __r;
}
int QCalendar_isSolar(void* self) {
	auto&& __r = ((QCalendar*)self)->isSolar();
	return __r;
}
int QCalendar_isProleptic(void* self) {
	auto&& __r = ((QCalendar*)self)->isProleptic();
	return __r;
}
int QCalendar_hasYearZero(void* self) {
	auto&& __r = ((QCalendar*)self)->hasYearZero();
	return __r;
}
int QCalendar_maximumDaysInMonth(void* self) {
	auto&& __r = ((QCalendar*)self)->maximumDaysInMonth();
	return __r;
}
int QCalendar_minimumDaysInMonth(void* self) {
	auto&& __r = ((QCalendar*)self)->minimumDaysInMonth();
	return __r;
}
int QCalendar_maximumMonthsInYear(void* self) {
	auto&& __r = ((QCalendar*)self)->maximumMonthsInYear();
	return __r;
}
char* QCalendar_name(void* self, long long* __out_len) {
	auto&& __r = ((QCalendar*)self)->name();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QCalendar_dateFromParts(void* self, int year, int month, int day) {
	auto&& __r = ((QCalendar*)self)->dateFromParts(year, month, day);
	return new QDate(__r);
}
void* QCalendar_dateFromParts2(void* self, void* parts) {
	auto&& __r = ((QCalendar*)self)->dateFromParts(*(QCalendar::YearMonthDay*)parts);
	return new QDate(__r);
}
void* QCalendar_matchCenturyToWeekday(void* self, void* parts, int dow) {
	auto&& __r = ((QCalendar*)self)->matchCenturyToWeekday(*(QCalendar::YearMonthDay*)parts, dow);
	return new QDate(__r);
}
void* QCalendar_partsFromDate(void* self, void* date) {
	auto&& __r = ((QCalendar*)self)->partsFromDate(*(QDate*)date);
	return new QCalendar::YearMonthDay(__r);
}
int QCalendar_dayOfWeek(void* self, void* date) {
	auto&& __r = ((QCalendar*)self)->dayOfWeek(*(QDate*)date);
	return __r;
}
char* QCalendar_monthName(void* self, void* locale, int month, long long* __out_len) {
	auto&& __r = ((QCalendar*)self)->monthName(*(QLocale*)locale, month);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QCalendar_standaloneMonthName(void* self, void* locale, int month, long long* __out_len) {
	auto&& __r = ((QCalendar*)self)->standaloneMonthName(*(QLocale*)locale, month);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QCalendar_weekDayName(void* self, void* locale, int day, long long* __out_len) {
	auto&& __r = ((QCalendar*)self)->weekDayName(*(QLocale*)locale, day);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QCalendar_standaloneWeekDayName(void* self, void* locale, int day, long long* __out_len) {
	auto&& __r = ((QCalendar*)self)->standaloneWeekDayName(*(QLocale*)locale, day);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QCalendar_dateTimeToString(void* self, void* format, void* datetime, void* dateOnly, void* timeOnly, void* locale, long long* __out_len) {
	auto&& __r = ((QCalendar*)self)->dateTimeToString(*(QStringView*)format, *(QDateTime*)datetime, *(QDate*)dateOnly, *(QTime*)timeOnly, *(QLocale*)locale);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char** QCalendar_availableCalendars(long long* __out_len) {
	auto&& __r = QCalendar::availableCalendars();
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
int QCalendar_daysInMonth2(void* self, int month, int year) {
	auto&& __r = ((QCalendar*)self)->daysInMonth(month, year);
	return __r;
}
char* QCalendar_monthName2(void* self, void* locale, int month, int year, long long* __out_len) {
	auto&& __r = ((QCalendar*)self)->monthName(*(QLocale*)locale, month, year);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QCalendar_monthName3(void* self, void* locale, int month, int year, int format, long long* __out_len) {
	auto&& __r = ((QCalendar*)self)->monthName(*(QLocale*)locale, month, year, (QLocale::FormatType)(format));
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QCalendar_standaloneMonthName2(void* self, void* locale, int month, int year, long long* __out_len) {
	auto&& __r = ((QCalendar*)self)->standaloneMonthName(*(QLocale*)locale, month, year);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QCalendar_standaloneMonthName3(void* self, void* locale, int month, int year, int format, long long* __out_len) {
	auto&& __r = ((QCalendar*)self)->standaloneMonthName(*(QLocale*)locale, month, year, (QLocale::FormatType)(format));
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QCalendar_weekDayName2(void* self, void* locale, int day, int format, long long* __out_len) {
	auto&& __r = ((QCalendar*)self)->weekDayName(*(QLocale*)locale, day, (QLocale::FormatType)(format));
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QCalendar_standaloneWeekDayName2(void* self, void* locale, int day, int format, long long* __out_len) {
	auto&& __r = ((QCalendar*)self)->standaloneWeekDayName(*(QLocale*)locale, day, (QLocale::FormatType)(format));
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QCalendar_delete(void* self) { delete (QCalendar*)self; }

}

#include <QtCore/qproperty.h>
#include "gen_qscopedpropertyupdategroup.h"

extern "C" {

void* QScopedPropertyUpdateGroup_new() {
	return new QScopedPropertyUpdateGroup();
}
void QScopedPropertyUpdateGroup_delete(void* self) { delete (QScopedPropertyUpdateGroup*)self; }

}

#include <QtCore/qfileinfo.h>
#include <QtCore/qstring.h>
#include <QtCore/qfiledevice.h>
#include <QtCore/qdir.h>
#include <QtCore/qdatetime.h>
#include <QtCore/qtimezone.h>
#include "gen_qfileinfo.h"

extern "C" {

void* QFileInfo_new() {
	return new QFileInfo();
}
void* QFileInfo_new2(const char* file, long long file_len) {
	QString file_qstr = QString::fromUtf8(file, file_len);
	return new QFileInfo(file_qstr);
}
void* QFileInfo_new3(void* file) {
	return new QFileInfo(*(QFileDevice*)file);
}
void* QFileInfo_new4(void* dir, const char* file, long long file_len) {
	QString file_qstr = QString::fromUtf8(file, file_len);
	return new QFileInfo(*(QDir*)dir, file_qstr);
}
void* QFileInfo_new5(void* fileinfo) {
	return new QFileInfo(*(QFileInfo*)fileinfo);
}
void* QFileInfo_OperatorAssign(void* self, void* fileinfo) {
	auto&& __r = ((QFileInfo*)self)->operator=(*(QFileInfo*)fileinfo);
	return (void*)&(__r);
}
void QFileInfo_swap(void* self, void* other) {
	((QFileInfo*)self)->swap(*(QFileInfo*)other);
}
void QFileInfo_setFile(void* self, const char* file, long long file_len) {
	QString file_qstr = QString::fromUtf8(file, file_len);
	((QFileInfo*)self)->setFile(file_qstr);
}
void QFileInfo_setFile2(void* self, void* file) {
	((QFileInfo*)self)->setFile(*(QFileDevice*)file);
}
void QFileInfo_setFile3(void* self, void* dir, const char* file, long long file_len) {
	QString file_qstr = QString::fromUtf8(file, file_len);
	((QFileInfo*)self)->setFile(*(QDir*)dir, file_qstr);
}
int QFileInfo_exists(void* self) {
	auto&& __r = ((QFileInfo*)self)->exists();
	return __r;
}
int QFileInfo_exists2(const char* file, long long file_len) {
	QString file_qstr = QString::fromUtf8(file, file_len);
	auto&& __r = QFileInfo::exists(file_qstr);
	return __r;
}
void QFileInfo_refresh(void* self) {
	((QFileInfo*)self)->refresh();
}
char* QFileInfo_filePath(void* self, long long* __out_len) {
	auto&& __r = ((QFileInfo*)self)->filePath();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QFileInfo_absoluteFilePath(void* self, long long* __out_len) {
	auto&& __r = ((QFileInfo*)self)->absoluteFilePath();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QFileInfo_canonicalFilePath(void* self, long long* __out_len) {
	auto&& __r = ((QFileInfo*)self)->canonicalFilePath();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QFileInfo_fileName(void* self, long long* __out_len) {
	auto&& __r = ((QFileInfo*)self)->fileName();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QFileInfo_baseName(void* self, long long* __out_len) {
	auto&& __r = ((QFileInfo*)self)->baseName();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QFileInfo_completeBaseName(void* self, long long* __out_len) {
	auto&& __r = ((QFileInfo*)self)->completeBaseName();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QFileInfo_suffix(void* self, long long* __out_len) {
	auto&& __r = ((QFileInfo*)self)->suffix();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QFileInfo_bundleName(void* self, long long* __out_len) {
	auto&& __r = ((QFileInfo*)self)->bundleName();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QFileInfo_completeSuffix(void* self, long long* __out_len) {
	auto&& __r = ((QFileInfo*)self)->completeSuffix();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QFileInfo_path(void* self, long long* __out_len) {
	auto&& __r = ((QFileInfo*)self)->path();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QFileInfo_absolutePath(void* self, long long* __out_len) {
	auto&& __r = ((QFileInfo*)self)->absolutePath();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QFileInfo_canonicalPath(void* self, long long* __out_len) {
	auto&& __r = ((QFileInfo*)self)->canonicalPath();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QFileInfo_dir(void* self) {
	auto&& __r = ((QFileInfo*)self)->dir();
	return new QDir(__r);
}
void* QFileInfo_absoluteDir(void* self) {
	auto&& __r = ((QFileInfo*)self)->absoluteDir();
	return new QDir(__r);
}
int QFileInfo_isReadable(void* self) {
	auto&& __r = ((QFileInfo*)self)->isReadable();
	return __r;
}
int QFileInfo_isWritable(void* self) {
	auto&& __r = ((QFileInfo*)self)->isWritable();
	return __r;
}
int QFileInfo_isExecutable(void* self) {
	auto&& __r = ((QFileInfo*)self)->isExecutable();
	return __r;
}
int QFileInfo_isHidden(void* self) {
	auto&& __r = ((QFileInfo*)self)->isHidden();
	return __r;
}
int QFileInfo_isNativePath(void* self) {
	auto&& __r = ((QFileInfo*)self)->isNativePath();
	return __r;
}
int QFileInfo_isRelative(void* self) {
	auto&& __r = ((QFileInfo*)self)->isRelative();
	return __r;
}
int QFileInfo_isAbsolute(void* self) {
	auto&& __r = ((QFileInfo*)self)->isAbsolute();
	return __r;
}
int QFileInfo_makeAbsolute(void* self) {
	auto&& __r = ((QFileInfo*)self)->makeAbsolute();
	return __r;
}
int QFileInfo_isFile(void* self) {
	auto&& __r = ((QFileInfo*)self)->isFile();
	return __r;
}
int QFileInfo_isDir(void* self) {
	auto&& __r = ((QFileInfo*)self)->isDir();
	return __r;
}
int QFileInfo_isSymLink(void* self) {
	auto&& __r = ((QFileInfo*)self)->isSymLink();
	return __r;
}
int QFileInfo_isSymbolicLink(void* self) {
	auto&& __r = ((QFileInfo*)self)->isSymbolicLink();
	return __r;
}
int QFileInfo_isOther(void* self) {
	auto&& __r = ((QFileInfo*)self)->isOther();
	return __r;
}
int QFileInfo_isShortcut(void* self) {
	auto&& __r = ((QFileInfo*)self)->isShortcut();
	return __r;
}
int QFileInfo_isAlias(void* self) {
	auto&& __r = ((QFileInfo*)self)->isAlias();
	return __r;
}
int QFileInfo_isJunction(void* self) {
	auto&& __r = ((QFileInfo*)self)->isJunction();
	return __r;
}
int QFileInfo_isRoot(void* self) {
	auto&& __r = ((QFileInfo*)self)->isRoot();
	return __r;
}
int QFileInfo_isBundle(void* self) {
	auto&& __r = ((QFileInfo*)self)->isBundle();
	return __r;
}
char* QFileInfo_symLinkTarget(void* self, long long* __out_len) {
	auto&& __r = ((QFileInfo*)self)->symLinkTarget();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QFileInfo_readSymLink(void* self, long long* __out_len) {
	auto&& __r = ((QFileInfo*)self)->readSymLink();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QFileInfo_junctionTarget(void* self, long long* __out_len) {
	auto&& __r = ((QFileInfo*)self)->junctionTarget();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QFileInfo_owner(void* self, long long* __out_len) {
	auto&& __r = ((QFileInfo*)self)->owner();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
unsigned int QFileInfo_ownerId(void* self) {
	auto&& __r = ((QFileInfo*)self)->ownerId();
	return __r;
}
char* QFileInfo_group(void* self, long long* __out_len) {
	auto&& __r = ((QFileInfo*)self)->group();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
unsigned int QFileInfo_groupId(void* self) {
	auto&& __r = ((QFileInfo*)self)->groupId();
	return __r;
}
int QFileInfo_permission(void* self, int permissions) {
	auto&& __r = ((QFileInfo*)self)->permission(QFlags<QFileDevice::Permission>(permissions));
	return __r;
}
int QFileInfo_permissions(void* self) {
	auto&& __r = ((QFileInfo*)self)->permissions();
	return (int)(__r);
}
long long QFileInfo_size(void* self) {
	auto&& __r = ((QFileInfo*)self)->size();
	return __r;
}
void* QFileInfo_birthTime(void* self) {
	auto&& __r = ((QFileInfo*)self)->birthTime();
	return new QDateTime(__r);
}
void* QFileInfo_metadataChangeTime(void* self) {
	auto&& __r = ((QFileInfo*)self)->metadataChangeTime();
	return new QDateTime(__r);
}
void* QFileInfo_lastModified(void* self) {
	auto&& __r = ((QFileInfo*)self)->lastModified();
	return new QDateTime(__r);
}
void* QFileInfo_lastRead(void* self) {
	auto&& __r = ((QFileInfo*)self)->lastRead();
	return new QDateTime(__r);
}
void* QFileInfo_fileTime(void* self, int time) {
	auto&& __r = ((QFileInfo*)self)->fileTime((QFileDevice::FileTime)(time));
	return new QDateTime(__r);
}
void* QFileInfo_birthTime2(void* self, void* tz) {
	auto&& __r = ((QFileInfo*)self)->birthTime(*(QTimeZone*)tz);
	return new QDateTime(__r);
}
void* QFileInfo_metadataChangeTime2(void* self, void* tz) {
	auto&& __r = ((QFileInfo*)self)->metadataChangeTime(*(QTimeZone*)tz);
	return new QDateTime(__r);
}
void* QFileInfo_lastModified2(void* self, void* tz) {
	auto&& __r = ((QFileInfo*)self)->lastModified(*(QTimeZone*)tz);
	return new QDateTime(__r);
}
void* QFileInfo_lastRead2(void* self, void* tz) {
	auto&& __r = ((QFileInfo*)self)->lastRead(*(QTimeZone*)tz);
	return new QDateTime(__r);
}
void* QFileInfo_fileTime2(void* self, int time, void* tz) {
	auto&& __r = ((QFileInfo*)self)->fileTime((QFileDevice::FileTime)(time), *(QTimeZone*)tz);
	return new QDateTime(__r);
}
int QFileInfo_caching(void* self) {
	auto&& __r = ((QFileInfo*)self)->caching();
	return __r;
}
void QFileInfo_setCaching(void* self, int on) {
	((QFileInfo*)self)->setCaching(on);
}
void QFileInfo_stat(void* self) {
	((QFileInfo*)self)->stat();
}
int QFileInfo_OperatorEqual(void* self, void* rhs) {
	auto&& __r = (*(QFileInfo*)self) == (*(QFileInfo*)rhs);
	return __r;
}
int QFileInfo_OperatorNotEqual(void* self, void* rhs) {
	auto&& __r = (*(QFileInfo*)self) != (*(QFileInfo*)rhs);
	return __r;
}
void QFileInfo_delete(void* self) { delete (QFileInfo*)self; }

}

#include <QtCore/qjsonarray.h>
#include <QtCore/qjsonvalue.h>
#include <QtCore/qstring.h>
#include <QtCore/qvariant.h>
#include "gen_qjsonarray.h"

extern "C" {

void* QJsonArray_new() {
	return new QJsonArray();
}
void* QJsonArray_new2(void* other) {
	return new QJsonArray(*(QJsonArray*)other);
}
void* QJsonArray_OperatorAssign(void* self, void* other) {
	auto&& __r = ((QJsonArray*)self)->operator=(*(QJsonArray*)other);
	return (void*)&(__r);
}
void* QJsonArray_fromStringList(char** list, long long list_len) {
	QList<QString> list_list; list_list.reserve(list_len);
	for (long long __i = 0; __i < list_len; __i++) list_list.append(QString::fromUtf8(list[__i]));
	auto&& __r = QJsonArray::fromStringList(list_list);
	return new QJsonArray(__r);
}
void* QJsonArray_fromVariantList(void** list, long long list_len) {
	QList<QVariant> list_list; list_list.reserve(list_len);
	for (long long __i = 0; __i < list_len; __i++) list_list.append(*(QVariant*)list[__i]);
	auto&& __r = QJsonArray::fromVariantList(list_list);
	return new QJsonArray(__r);
}
void** QJsonArray_toVariantList(void* self, long long* __out_len) {
	auto&& __r = ((QJsonArray*)self)->toVariantList();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QVariant(__e);
	return __arr;
}
long long QJsonArray_size(void* self) {
	auto&& __r = ((QJsonArray*)self)->size();
	return __r;
}
long long QJsonArray_count(void* self) {
	auto&& __r = ((QJsonArray*)self)->count();
	return __r;
}
int QJsonArray_isEmpty(void* self) {
	auto&& __r = ((QJsonArray*)self)->isEmpty();
	return __r;
}
void* QJsonArray_at(void* self, long long i) {
	auto&& __r = ((QJsonArray*)self)->at(i);
	return new QJsonValue(__r);
}
void* QJsonArray_first(void* self) {
	auto&& __r = ((QJsonArray*)self)->first();
	return new QJsonValue(__r);
}
void* QJsonArray_last(void* self) {
	auto&& __r = ((QJsonArray*)self)->last();
	return new QJsonValue(__r);
}
void QJsonArray_prepend(void* self, void* value) {
	((QJsonArray*)self)->prepend(*(QJsonValue*)value);
}
void QJsonArray_append(void* self, void* value) {
	((QJsonArray*)self)->append(*(QJsonValue*)value);
}
void QJsonArray_removeAt(void* self, long long i) {
	((QJsonArray*)self)->removeAt(i);
}
void* QJsonArray_takeAt(void* self, long long i) {
	auto&& __r = ((QJsonArray*)self)->takeAt(i);
	return new QJsonValue(__r);
}
void QJsonArray_removeFirst(void* self) {
	((QJsonArray*)self)->removeFirst();
}
void QJsonArray_removeLast(void* self) {
	((QJsonArray*)self)->removeLast();
}
void QJsonArray_insert(void* self, long long i, void* value) {
	((QJsonArray*)self)->insert(i, *(QJsonValue*)value);
}
void QJsonArray_replace(void* self, long long i, void* value) {
	((QJsonArray*)self)->replace(i, *(QJsonValue*)value);
}
int QJsonArray_contains(void* self, void* element) {
	auto&& __r = ((QJsonArray*)self)->contains(*(QJsonValue*)element);
	return __r;
}
void* QJsonArray_OperatorSubscript(void* self, long long i) {
	auto&& __r = ((QJsonArray*)self)->operator[](i);
	return new QJsonValueRef(__r);
}
void* QJsonArray_OperatorSubscript2(void* self, long long i) {
	auto&& __r = ((QJsonArray*)self)->operator[](i);
	return new QJsonValue(__r);
}
void QJsonArray_swap(void* self, void* other) {
	((QJsonArray*)self)->swap(*(QJsonArray*)other);
}
void* QJsonArray_OperatorPlus(void* self, void* v) {
	auto&& __r = ((QJsonArray*)self)->operator+(*(QJsonValue*)v);
	return new QJsonArray(__r);
}
void* QJsonArray_OperatorPlusAssign(void* self, void* v) {
	auto&& __r = ((QJsonArray*)self)->operator+=(*(QJsonValue*)v);
	return (void*)&(__r);
}
void* QJsonArray_OperatorShiftLeft(void* self, void* v) {
	auto&& __r = ((QJsonArray*)self)->operator<<(*(QJsonValue*)v);
	return (void*)&(__r);
}
void QJsonArray_push_back(void* self, void* t) {
	((QJsonArray*)self)->push_back(*(QJsonValue*)t);
}
void QJsonArray_push_front(void* self, void* t) {
	((QJsonArray*)self)->push_front(*(QJsonValue*)t);
}
void QJsonArray_pop_front(void* self) {
	((QJsonArray*)self)->pop_front();
}
void QJsonArray_pop_back(void* self) {
	((QJsonArray*)self)->pop_back();
}
int QJsonArray_empty(void* self) {
	auto&& __r = ((QJsonArray*)self)->empty();
	return __r;
}
int QJsonArray_OperatorEqual(void* self, void* rhs) {
	auto&& __r = (*(QJsonArray*)self) == (*(QJsonArray*)rhs);
	return __r;
}
int QJsonArray_OperatorNotEqual(void* self, void* rhs) {
	auto&& __r = (*(QJsonArray*)self) != (*(QJsonArray*)rhs);
	return __r;
}
int QJsonArray_OperatorEqual2(void* self, void* rhs) {
	auto&& __r = (*(QJsonArray*)self) == (*(QJsonValue*)rhs);
	return __r;
}
int QJsonArray_OperatorNotEqual2(void* self, void* rhs) {
	auto&& __r = (*(QJsonArray*)self) != (*(QJsonValue*)rhs);
	return __r;
}
void QJsonArray_delete(void* self) { delete (QJsonArray*)self; }

}

#include <QtCore/qoperatingsystemversion.h>
#include <QtCore/qstring.h>
#include <QtCore/qversionnumber.h>
#include "gen_qoperatingsystemversionbase.h"

extern "C" {

void* QOperatingSystemVersionBase_new(int osType, int vmajor) {
	return new QOperatingSystemVersionBase((QOperatingSystemVersionBase::OSType)(osType), vmajor);
}
void* QOperatingSystemVersionBase_new2(void* arg1) {
	return new QOperatingSystemVersionBase(*(QOperatingSystemVersionBase*)arg1);
}
void* QOperatingSystemVersionBase_new3(int osType, int vmajor, int vminor) {
	return new QOperatingSystemVersionBase((QOperatingSystemVersionBase::OSType)(osType), vmajor, vminor);
}
void* QOperatingSystemVersionBase_new4(int osType, int vmajor, int vminor, int vmicro) {
	return new QOperatingSystemVersionBase((QOperatingSystemVersionBase::OSType)(osType), vmajor, vminor, vmicro);
}
void* QOperatingSystemVersionBase_current() {
	auto&& __r = QOperatingSystemVersionBase::current();
	return new QOperatingSystemVersionBase(__r);
}
char* QOperatingSystemVersionBase_name(void* osversion, long long* __out_len) {
	auto&& __r = QOperatingSystemVersionBase::name(*(QOperatingSystemVersionBase*)osversion);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QOperatingSystemVersionBase_currentType() {
	auto&& __r = QOperatingSystemVersionBase::currentType();
	return (int)(__r);
}
void* QOperatingSystemVersionBase_version(void* self) {
	auto&& __r = ((QOperatingSystemVersionBase*)self)->version();
	return new QVersionNumber(__r);
}
int QOperatingSystemVersionBase_majorVersion(void* self) {
	auto&& __r = ((QOperatingSystemVersionBase*)self)->majorVersion();
	return __r;
}
int QOperatingSystemVersionBase_minorVersion(void* self) {
	auto&& __r = ((QOperatingSystemVersionBase*)self)->minorVersion();
	return __r;
}
int QOperatingSystemVersionBase_microVersion(void* self) {
	auto&& __r = ((QOperatingSystemVersionBase*)self)->microVersion();
	return __r;
}
int QOperatingSystemVersionBase_segmentCount(void* self) {
	auto&& __r = ((QOperatingSystemVersionBase*)self)->segmentCount();
	return __r;
}
int QOperatingSystemVersionBase_type(void* self) {
	auto&& __r = ((QOperatingSystemVersionBase*)self)->type();
	return (int)(__r);
}
char* QOperatingSystemVersionBase_name2(void* self, long long* __out_len) {
	auto&& __r = ((QOperatingSystemVersionBase*)self)->name();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QOperatingSystemVersionBase_OperatorGreater(void* self, void* rhs) {
	auto&& __r = (*(QOperatingSystemVersionBase*)self) > (*(QOperatingSystemVersionBase*)rhs);
	return __r;
}
int QOperatingSystemVersionBase_OperatorGreaterOrEqual(void* self, void* rhs) {
	auto&& __r = (*(QOperatingSystemVersionBase*)self) >= (*(QOperatingSystemVersionBase*)rhs);
	return __r;
}
int QOperatingSystemVersionBase_OperatorLesser(void* self, void* rhs) {
	auto&& __r = (*(QOperatingSystemVersionBase*)self) < (*(QOperatingSystemVersionBase*)rhs);
	return __r;
}
int QOperatingSystemVersionBase_OperatorLesserOrEqual(void* self, void* rhs) {
	auto&& __r = (*(QOperatingSystemVersionBase*)self) <= (*(QOperatingSystemVersionBase*)rhs);
	return __r;
}
void QOperatingSystemVersionBase_delete(void* self) { delete (QOperatingSystemVersionBase*)self; }

}

#include <QtCore/qrandom.h>
#include "gen_qrandomgenerator64.h"

extern "C" {

void* QRandomGenerator64_base(void* self) { return (void*)static_cast<QRandomGenerator*>((QRandomGenerator64*)self); }
void* QRandomGenerator64_new() {
	return new QRandomGenerator64();
}
void* QRandomGenerator64_new2(void* other) {
	return new QRandomGenerator64(*(QRandomGenerator*)other);
}
void* QRandomGenerator64_new3(void* arg1) {
	return new QRandomGenerator64(*(QRandomGenerator64*)arg1);
}
void* QRandomGenerator64_new4(unsigned int seedValue) {
	return new QRandomGenerator64(seedValue);
}
unsigned long long QRandomGenerator64_generate(void* self) {
	auto&& __r = ((QRandomGenerator64*)self)->generate();
	return __r;
}
unsigned long long QRandomGenerator64_OperatorCall(void* self) {
	auto&& __r = ((QRandomGenerator64*)self)->operator()();
	return __r;
}
void QRandomGenerator64_discard(void* self, unsigned long long z) {
	((QRandomGenerator64*)self)->discard(z);
}
unsigned long long QRandomGenerator64_min() {
	auto&& __r = QRandomGenerator64::min();
	return __r;
}
unsigned long long QRandomGenerator64_max() {
	auto&& __r = QRandomGenerator64::max();
	return __r;
}
void* QRandomGenerator64_system() {
	auto&& __r = QRandomGenerator64::system();
	return (void*)(__r);
}
void* QRandomGenerator64_global() {
	auto&& __r = QRandomGenerator64::global();
	return (void*)(__r);
}
void* QRandomGenerator64_securelySeeded() {
	auto&& __r = QRandomGenerator64::securelySeeded();
	return new QRandomGenerator64(__r);
}
void* QRandomGenerator64_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QRandomGenerator64*)self)->operator=(*(QRandomGenerator64*)arg1);
	return (void*)&(__r);
}
void QRandomGenerator64_delete(void* self) { delete (QRandomGenerator64*)self; }

}

#include <QtCore/qsortfilterproxymodel.h>
#include <QtCore/qabstractproxymodel.h>
#include <QtCore/qobject.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qabstractitemmodel.h>
#include <QtCore/qitemselectionmodel.h>
#include <QtCore/qregularexpression.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qvariant.h>
#include <QtCore/qmimedata.h>
#include <QtCore/qsize.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qdatastream.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qsortfilterproxymodel.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QSortFilterProxyModel__Go : public QSortFilterProxyModel {
public:
	long long __goId;
	unsigned char __goOn[50] = {};
	template<typename... A> QSortFilterProxyModel__Go(long long id, A&&... a) : QSortFilterProxyModel(std::forward<A>(a)...), __goId(id) {}
	using QSortFilterProxyModel::beginFilterChange;
	using QSortFilterProxyModel::endFilterChange;
	using QSortFilterProxyModel::invalidateFilter;
	using QSortFilterProxyModel::invalidateRowsFilter;
	using QSortFilterProxyModel::invalidateColumnsFilter;
	using QSortFilterProxyModel::createIndex;
	using QSortFilterProxyModel::encodeData;
	using QSortFilterProxyModel::decodeData;
	using QSortFilterProxyModel::beginInsertRows;
	using QSortFilterProxyModel::endInsertRows;
	using QSortFilterProxyModel::beginRemoveRows;
	using QSortFilterProxyModel::endRemoveRows;
	using QSortFilterProxyModel::beginMoveRows;
	using QSortFilterProxyModel::endMoveRows;
	using QSortFilterProxyModel::beginInsertColumns;
	using QSortFilterProxyModel::endInsertColumns;
	using QSortFilterProxyModel::beginRemoveColumns;
	using QSortFilterProxyModel::endRemoveColumns;
	using QSortFilterProxyModel::beginMoveColumns;
	using QSortFilterProxyModel::endMoveColumns;
	using QSortFilterProxyModel::beginResetModel;
	using QSortFilterProxyModel::endResetModel;
	using QSortFilterProxyModel::changePersistentIndex;
	using QSortFilterProxyModel::changePersistentIndexList;
	using QSortFilterProxyModel::persistentIndexList;
	using QSortFilterProxyModel::sender;
	using QSortFilterProxyModel::senderSignalIndex;
	using QSortFilterProxyModel::receivers;
	using QSortFilterProxyModel::isSignalConnected;
	bool cppgoBase_filterAcceptsRow(int __a0, const QModelIndex& __a1) { return QSortFilterProxyModel::filterAcceptsRow(__a0, __a1); }
	bool cppgoBase_filterAcceptsColumn(int __a0, const QModelIndex& __a1) { return QSortFilterProxyModel::filterAcceptsColumn(__a0, __a1); }
	bool cppgoBase_lessThan(const QModelIndex& __a0, const QModelIndex& __a1) { return QSortFilterProxyModel::lessThan(__a0, __a1); }
	void cppgoBase_resetInternalData() { QSortFilterProxyModel::resetInternalData(); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QSortFilterProxyModel::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QSortFilterProxyModel::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QSortFilterProxyModel::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QSortFilterProxyModel::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QSortFilterProxyModel::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QSortFilterProxyModel::metaObject();
	}
	void setSourceModel(QAbstractItemModel* sourceModel) override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7818871925547456643LL, (void*)sourceModel, &__h); if (__h) return; }
		QSortFilterProxyModel::setSourceModel(sourceModel);
	}
	QModelIndex mapToSource(const QModelIndex& proxyIndex) const override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, -2333321292016234891LL, (void*)&proxyIndex, &__h); if (__h) return *(QModelIndex*)__r; }
		return QSortFilterProxyModel::mapToSource(proxyIndex);
	}
	QModelIndex mapFromSource(const QModelIndex& sourceIndex) const override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, -1717009778236344012LL, (void*)&sourceIndex, &__h); if (__h) return *(QModelIndex*)__r; }
		return QSortFilterProxyModel::mapFromSource(sourceIndex);
	}
	QItemSelection mapSelectionToSource(const QItemSelection& proxySelection) const override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 2788628866280405837LL, (void*)&proxySelection, &__h); if (__h) return *(QItemSelection*)__r; }
		return QSortFilterProxyModel::mapSelectionToSource(proxySelection);
	}
	QItemSelection mapSelectionFromSource(const QItemSelection& sourceSelection) const override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 9067522407071207004LL, (void*)&sourceSelection, &__h); if (__h) return *(QItemSelection*)__r; }
		return QSortFilterProxyModel::mapSelectionFromSource(sourceSelection);
	}
	bool filterAcceptsRow(int source_row, const QModelIndex& source_parent) const override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_int_int_voidp) { int __r = ((int(*)(long long, long long, int, void*, int*))cppgo_virt_int_int_voidp)(__goId, -8878517896795061028LL, source_row, (void*)&source_parent, &__h); if (__h) return __r; }
		return QSortFilterProxyModel::filterAcceptsRow(source_row, source_parent);
	}
	bool filterAcceptsColumn(int source_column, const QModelIndex& source_parent) const override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_int_int_voidp) { int __r = ((int(*)(long long, long long, int, void*, int*))cppgo_virt_int_int_voidp)(__goId, 436376032338170754LL, source_column, (void*)&source_parent, &__h); if (__h) return __r; }
		return QSortFilterProxyModel::filterAcceptsColumn(source_column, source_parent);
	}
	bool lessThan(const QModelIndex& source_left, const QModelIndex& source_right) const override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, -5863021011917066845LL, (void*)&source_left, (void*)&source_right, &__h); if (__h) return __r; }
		return QSortFilterProxyModel::lessThan(source_left, source_right);
	}
	QModelIndex index(int row, int column, const QModelIndex& parent) const override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_voidp_int_int_voidp) { void* __r = ((void*(*)(long long, long long, int, int, void*, int*))cppgo_virt_voidp_int_int_voidp)(__goId, -8948777187306027381LL, row, column, (void*)&parent, &__h); if (__h) return *(QModelIndex*)__r; }
		return QSortFilterProxyModel::index(row, column, parent);
	}
	QModelIndex parent(const QModelIndex& child) const override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 7371339085820260189LL, (void*)&child, &__h); if (__h) return *(QModelIndex*)__r; }
		return QSortFilterProxyModel::parent(child);
	}
	QModelIndex sibling(int row, int column, const QModelIndex& idx) const override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_voidp_int_int_voidp) { void* __r = ((void*(*)(long long, long long, int, int, void*, int*))cppgo_virt_voidp_int_int_voidp)(__goId, 397329346813782521LL, row, column, (void*)&idx, &__h); if (__h) return *(QModelIndex*)__r; }
		return QSortFilterProxyModel::sibling(row, column, idx);
	}
	int rowCount(const QModelIndex& parent) const override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1960960171956884010LL, (void*)&parent, &__h); if (__h) return __r; }
		return QSortFilterProxyModel::rowCount(parent);
	}
	int columnCount(const QModelIndex& parent) const override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, -8256963024669584834LL, (void*)&parent, &__h); if (__h) return __r; }
		return QSortFilterProxyModel::columnCount(parent);
	}
	bool hasChildren(const QModelIndex& parent) const override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 3369452259144514380LL, (void*)&parent, &__h); if (__h) return __r; }
		return QSortFilterProxyModel::hasChildren(parent);
	}
	QVariant data(const QModelIndex& index, int role) const override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_voidp_voidp_int) { void* __r = ((void*(*)(long long, long long, void*, int, int*))cppgo_virt_voidp_voidp_int)(__goId, -8837375942098466299LL, (void*)&index, role, &__h); if (__h) return *(QVariant*)__r; }
		return QSortFilterProxyModel::data(index, role);
	}
	bool setData(const QModelIndex& index, const QVariant& value, int role) override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_int_voidp_voidp_int) { int __r = ((int(*)(long long, long long, void*, void*, int, int*))cppgo_virt_int_voidp_voidp_int)(__goId, 35191454575202219LL, (void*)&index, (void*)&value, role, &__h); if (__h) return __r; }
		return QSortFilterProxyModel::setData(index, value, role);
	}
	QVariant headerData(int section, Qt::Orientation orientation, int role) const override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_voidp_int_int_int) { void* __r = ((void*(*)(long long, long long, int, int, int, int*))cppgo_virt_voidp_int_int_int)(__goId, -3297140021896693688LL, section, (int)(orientation), role, &__h); if (__h) return *(QVariant*)__r; }
		return QSortFilterProxyModel::headerData(section, orientation, role);
	}
	bool setHeaderData(int section, Qt::Orientation orientation, const QVariant& value, int role) override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_int_int_int_voidp_int) { int __r = ((int(*)(long long, long long, int, int, void*, int, int*))cppgo_virt_int_int_int_voidp_int)(__goId, 4852200306590054850LL, section, (int)(orientation), (void*)&value, role, &__h); if (__h) return __r; }
		return QSortFilterProxyModel::setHeaderData(section, orientation, value, role);
	}
	QMimeData* mimeData(const QList<QModelIndex>& indexes) const override {
		int __h = 0;
		std::vector<void*> indexes_v; indexes_v.reserve(indexes.size());
		for (long long __i = 0; __i < indexes.size(); __i++) indexes_v.push_back((void*)&indexes[__i]);
		if (__goOn[19] && cppgo_virt_voidp_voidpp_longlong) { void* __r = ((void*(*)(long long, long long, void**, long long, int*))cppgo_virt_voidp_voidpp_longlong)(__goId, 5175297091024974749LL, indexes_v.data(), (long long)indexes.size(), &__h); if (__h) return (QMimeData*)__r; }
		return QSortFilterProxyModel::mimeData(indexes);
	}
	bool dropMimeData(const QMimeData* data, Qt::DropAction action, int row, int column, const QModelIndex& parent) override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_int_voidp_int_int_int_voidp) { int __r = ((int(*)(long long, long long, void*, int, int, int, void*, int*))cppgo_virt_int_voidp_int_int_int_voidp)(__goId, -9049994684305682212LL, (void*)data, (int)(action), row, column, (void*)&parent, &__h); if (__h) return __r; }
		return QSortFilterProxyModel::dropMimeData(data, action, row, column, parent);
	}
	bool insertRows(int row, int count, const QModelIndex& parent) override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_int_int_int_voidp) { int __r = ((int(*)(long long, long long, int, int, void*, int*))cppgo_virt_int_int_int_voidp)(__goId, -2775729825962395603LL, row, count, (void*)&parent, &__h); if (__h) return __r; }
		return QSortFilterProxyModel::insertRows(row, count, parent);
	}
	bool insertColumns(int column, int count, const QModelIndex& parent) override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_int_int_int_voidp) { int __r = ((int(*)(long long, long long, int, int, void*, int*))cppgo_virt_int_int_int_voidp)(__goId, -8973293846230792237LL, column, count, (void*)&parent, &__h); if (__h) return __r; }
		return QSortFilterProxyModel::insertColumns(column, count, parent);
	}
	bool removeRows(int row, int count, const QModelIndex& parent) override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_int_int_int_voidp) { int __r = ((int(*)(long long, long long, int, int, void*, int*))cppgo_virt_int_int_int_voidp)(__goId, 5104072549182222656LL, row, count, (void*)&parent, &__h); if (__h) return __r; }
		return QSortFilterProxyModel::removeRows(row, count, parent);
	}
	bool removeColumns(int column, int count, const QModelIndex& parent) override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_int_int_int_voidp) { int __r = ((int(*)(long long, long long, int, int, void*, int*))cppgo_virt_int_int_int_voidp)(__goId, 2078548737056540420LL, column, count, (void*)&parent, &__h); if (__h) return __r; }
		return QSortFilterProxyModel::removeColumns(column, count, parent);
	}
	void fetchMore(const QModelIndex& parent) override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3553294593611130806LL, (void*)&parent, &__h); if (__h) return; }
		QSortFilterProxyModel::fetchMore(parent);
	}
	bool canFetchMore(const QModelIndex& parent) const override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 8639262066473964150LL, (void*)&parent, &__h); if (__h) return __r; }
		return QSortFilterProxyModel::canFetchMore(parent);
	}
	QFlags<Qt::ItemFlag> flags(const QModelIndex& index) const override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1703382833545304812LL, (void*)&index, &__h); if (__h) return QFlags<Qt::ItemFlag>(__r); }
		return QSortFilterProxyModel::flags(index);
	}
	QModelIndex buddy(const QModelIndex& index) const override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 6112860112248786499LL, (void*)&index, &__h); if (__h) return *(QModelIndex*)__r; }
		return QSortFilterProxyModel::buddy(index);
	}
	QList<QModelIndex> match(const QModelIndex& start, int role, const QVariant& value, int hits, QFlags<Qt::MatchFlag> flags) const override {
		int __h = 0;
		if (__goOn[29] && cppgo_virtp_voidp_voidp_int_voidp_int_int) {
			long long __len = 0;
			void* __a = ((void*(*)(long long, long long, void*, int, void*, int, int, long long*, int*))cppgo_virtp_voidp_voidp_int_voidp_int_int)(__goId, -4341501217981328394LL, (void*)&start, role, (void*)&value, hits, (int)(flags), &__len, &__h);
			if (__h) {
				QList<QModelIndex> __r; __r.reserve(__len);
				void** __aa = (void**)__a;
				for (long long __i = 0; __i < __len; __i++) __r.append(*(QModelIndex*)__aa[__i]);
				if (__a) free(__a);
				return __r;
			}
		}
		return QSortFilterProxyModel::match(start, role, value, hits, flags);
	}
	QSize span(const QModelIndex& index) const override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, -8395342916587040543LL, (void*)&index, &__h); if (__h) return *(QSize*)__r; }
		return QSortFilterProxyModel::span(index);
	}
	void sort(int column, Qt::SortOrder order) override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_void_int_int) { ((void(*)(long long, long long, int, int, int*))cppgo_virt_void_int_int)(__goId, 6912441694453382161LL, column, (int)(order), &__h); if (__h) return; }
		QSortFilterProxyModel::sort(column, order);
	}
	QList<QString> mimeTypes() const override {
		int __h = 0;
		if (__goOn[32] && cppgo_virtp_voidp) {
			long long __len = 0;
			void* __a = ((void*(*)(long long, long long, long long*, int*))cppgo_virtp_voidp)(__goId, 5464467312394730418LL, &__len, &__h);
			if (__h) {
				QList<QString> __r; __r.reserve(__len);
				char** __aa = (char**)__a;
				for (long long __i = 0; __i < __len; __i++) { __r.append(QString::fromUtf8(__aa[__i])); free(__aa[__i]); }
				if (__a) free(__a);
				return __r;
			}
		}
		return QSortFilterProxyModel::mimeTypes();
	}
	QFlags<Qt::DropAction> supportedDropActions() const override {
		int __h = 0;
		if (__goOn[33] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -7784642176186821205LL, &__h); if (__h) return QFlags<Qt::DropAction>(__r); }
		return QSortFilterProxyModel::supportedDropActions();
	}
	bool submit() override {
		int __h = 0;
		if (__goOn[34] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -3788331537442685179LL, &__h); if (__h) return __r; }
		return QSortFilterProxyModel::submit();
	}
	void revert() override {
		int __h = 0;
		if (__goOn[35] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, -1522951607536117149LL, &__h); if (__h) return; }
		QSortFilterProxyModel::revert();
	}
	bool clearItemData(const QModelIndex& index) override {
		int __h = 0;
		if (__goOn[36] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 8914190500520542409LL, (void*)&index, &__h); if (__h) return __r; }
		return QSortFilterProxyModel::clearItemData(index);
	}
	bool canDropMimeData(const QMimeData* data, Qt::DropAction action, int row, int column, const QModelIndex& parent) const override {
		int __h = 0;
		if (__goOn[37] && cppgo_virt_int_voidp_int_int_int_voidp) { int __r = ((int(*)(long long, long long, void*, int, int, int, void*, int*))cppgo_virt_int_voidp_int_int_int_voidp)(__goId, 731539533644753052LL, (void*)data, (int)(action), row, column, (void*)&parent, &__h); if (__h) return __r; }
		return QSortFilterProxyModel::canDropMimeData(data, action, row, column, parent);
	}
	QFlags<Qt::DropAction> supportedDragActions() const override {
		int __h = 0;
		if (__goOn[38] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, 3421601386088570762LL, &__h); if (__h) return QFlags<Qt::DropAction>(__r); }
		return QSortFilterProxyModel::supportedDragActions();
	}
	bool moveRows(const QModelIndex& sourceParent, int sourceRow, int count, const QModelIndex& destinationParent, int destinationChild) override {
		int __h = 0;
		if (__goOn[39] && cppgo_virt_int_voidp_int_int_voidp_int) { int __r = ((int(*)(long long, long long, void*, int, int, void*, int, int*))cppgo_virt_int_voidp_int_int_voidp_int)(__goId, -6074819356947405415LL, (void*)&sourceParent, sourceRow, count, (void*)&destinationParent, destinationChild, &__h); if (__h) return __r; }
		return QSortFilterProxyModel::moveRows(sourceParent, sourceRow, count, destinationParent, destinationChild);
	}
	bool moveColumns(const QModelIndex& sourceParent, int sourceColumn, int count, const QModelIndex& destinationParent, int destinationChild) override {
		int __h = 0;
		if (__goOn[40] && cppgo_virt_int_voidp_int_int_voidp_int) { int __r = ((int(*)(long long, long long, void*, int, int, void*, int, int*))cppgo_virt_int_voidp_int_int_voidp_int)(__goId, -6321007591388309905LL, (void*)&sourceParent, sourceColumn, count, (void*)&destinationParent, destinationChild, &__h); if (__h) return __r; }
		return QSortFilterProxyModel::moveColumns(sourceParent, sourceColumn, count, destinationParent, destinationChild);
	}
	void multiData(const QModelIndex& index, QModelRoleDataSpan roleDataSpan) const override {
		int __h = 0;
		if (__goOn[41] && cppgo_virt_void_voidp_voidp) { ((void(*)(long long, long long, void*, void*, int*))cppgo_virt_void_voidp_voidp)(__goId, 731311060944096642LL, (void*)&index, (void*)&roleDataSpan, &__h); if (__h) return; }
		QSortFilterProxyModel::multiData(index, roleDataSpan);
	}
	void resetInternalData() override {
		int __h = 0;
		if (__goOn[42] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, -4510675770784051183LL, &__h); if (__h) return; }
		QSortFilterProxyModel::resetInternalData();
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[43] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QSortFilterProxyModel::event(event);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[44] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QSortFilterProxyModel::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[45] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QSortFilterProxyModel::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[46] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QSortFilterProxyModel::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[47] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QSortFilterProxyModel::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[48] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QSortFilterProxyModel::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[49] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QSortFilterProxyModel::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QSortFilterProxyModel;
static std::mutex __connsMu_QSortFilterProxyModel;

extern "C" {

void* QSortFilterProxyModel_base(void* self) { return (void*)static_cast<QAbstractProxyModel*>((QSortFilterProxyModel*)self); }
void QSortFilterProxyModel_chain(void* self, void** out) {
	QSortFilterProxyModel* __p = (QSortFilterProxyModel*)self;
	QAbstractProxyModel* __b0 = static_cast<QAbstractProxyModel*>(__p); out[0] = (void*)__b0;
	QAbstractItemModel* __b1 = static_cast<QAbstractItemModel*>(__b0); out[1] = (void*)__b1;
	QObject* __b2 = static_cast<QObject*>(__b1); out[2] = (void*)__b2;
}
void QSortFilterProxyModel__set_override(void* self, int idx) { ((QSortFilterProxyModel__Go*)self)->__goOn[idx] = 1; }
void* QSortFilterProxyModel_fromQObject(void* o) { return (void*)qobject_cast<QSortFilterProxyModel*>((QObject*)o); }
void* QSortFilterProxyModel_new(long long __goid) {
	return new QSortFilterProxyModel__Go(__goid);
}
void* QSortFilterProxyModel_new2(long long __goid, void* parent) {
	return new QSortFilterProxyModel__Go(__goid, (QObject*)parent);
}
char* QSortFilterProxyModel_tr(const char* s, long long* __out_len) {
	auto&& __r = QSortFilterProxyModel::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QSortFilterProxyModel_filterRegularExpression(void* self) {
	auto&& __r = ((QSortFilterProxyModel*)self)->filterRegularExpression();
	return new QRegularExpression(__r);
}
int QSortFilterProxyModel_filterKeyColumn(void* self) {
	auto&& __r = ((QSortFilterProxyModel*)self)->filterKeyColumn();
	return __r;
}
void QSortFilterProxyModel_setFilterKeyColumn(void* self, int column) {
	((QSortFilterProxyModel*)self)->setFilterKeyColumn(column);
}
int QSortFilterProxyModel_filterCaseSensitivity(void* self) {
	auto&& __r = ((QSortFilterProxyModel*)self)->filterCaseSensitivity();
	return (int)(__r);
}
void QSortFilterProxyModel_setFilterCaseSensitivity(void* self, int cs) {
	((QSortFilterProxyModel*)self)->setFilterCaseSensitivity((Qt::CaseSensitivity)(cs));
}
int QSortFilterProxyModel_sortCaseSensitivity(void* self) {
	auto&& __r = ((QSortFilterProxyModel*)self)->sortCaseSensitivity();
	return (int)(__r);
}
void QSortFilterProxyModel_setSortCaseSensitivity(void* self, int cs) {
	((QSortFilterProxyModel*)self)->setSortCaseSensitivity((Qt::CaseSensitivity)(cs));
}
int QSortFilterProxyModel_isSortLocaleAware(void* self) {
	auto&& __r = ((QSortFilterProxyModel*)self)->isSortLocaleAware();
	return __r;
}
void QSortFilterProxyModel_setSortLocaleAware(void* self, int on) {
	((QSortFilterProxyModel*)self)->setSortLocaleAware(on);
}
int QSortFilterProxyModel_sortColumn(void* self) {
	auto&& __r = ((QSortFilterProxyModel*)self)->sortColumn();
	return __r;
}
int QSortFilterProxyModel_sortOrder(void* self) {
	auto&& __r = ((QSortFilterProxyModel*)self)->sortOrder();
	return (int)(__r);
}
int QSortFilterProxyModel_dynamicSortFilter(void* self) {
	auto&& __r = ((QSortFilterProxyModel*)self)->dynamicSortFilter();
	return __r;
}
void QSortFilterProxyModel_setDynamicSortFilter(void* self, int enable) {
	((QSortFilterProxyModel*)self)->setDynamicSortFilter(enable);
}
int QSortFilterProxyModel_sortRole(void* self) {
	auto&& __r = ((QSortFilterProxyModel*)self)->sortRole();
	return __r;
}
void QSortFilterProxyModel_setSortRole(void* self, int role) {
	((QSortFilterProxyModel*)self)->setSortRole(role);
}
int QSortFilterProxyModel_filterRole(void* self) {
	auto&& __r = ((QSortFilterProxyModel*)self)->filterRole();
	return __r;
}
void QSortFilterProxyModel_setFilterRole(void* self, int role) {
	((QSortFilterProxyModel*)self)->setFilterRole(role);
}
int QSortFilterProxyModel_isRecursiveFilteringEnabled(void* self) {
	auto&& __r = ((QSortFilterProxyModel*)self)->isRecursiveFilteringEnabled();
	return __r;
}
void QSortFilterProxyModel_setRecursiveFilteringEnabled(void* self, int recursive) {
	((QSortFilterProxyModel*)self)->setRecursiveFilteringEnabled(recursive);
}
int QSortFilterProxyModel_autoAcceptChildRows(void* self) {
	auto&& __r = ((QSortFilterProxyModel*)self)->autoAcceptChildRows();
	return __r;
}
void QSortFilterProxyModel_setAutoAcceptChildRows(void* self, int accept) {
	((QSortFilterProxyModel*)self)->setAutoAcceptChildRows(accept);
}
void QSortFilterProxyModel_setFilterRegularExpression(void* self, const char* pattern, long long pattern_len) {
	QString pattern_qstr = QString::fromUtf8(pattern, pattern_len);
	((QSortFilterProxyModel*)self)->setFilterRegularExpression(pattern_qstr);
}
void QSortFilterProxyModel_setFilterRegularExpression2(void* self, void* regularExpression) {
	((QSortFilterProxyModel*)self)->setFilterRegularExpression(*(QRegularExpression*)regularExpression);
}
void QSortFilterProxyModel_setFilterWildcard(void* self, const char* pattern, long long pattern_len) {
	QString pattern_qstr = QString::fromUtf8(pattern, pattern_len);
	((QSortFilterProxyModel*)self)->setFilterWildcard(pattern_qstr);
}
void QSortFilterProxyModel_setFilterFixedString(void* self, const char* pattern, long long pattern_len) {
	QString pattern_qstr = QString::fromUtf8(pattern, pattern_len);
	((QSortFilterProxyModel*)self)->setFilterFixedString(pattern_qstr);
}
void QSortFilterProxyModel_invalidate(void* self) {
	((QSortFilterProxyModel*)self)->invalidate();
}
void QSortFilterProxyModel_beginFilterChange(void* self) {
	((QSortFilterProxyModel__Go*)self)->beginFilterChange();
}
void QSortFilterProxyModel_endFilterChange(void* self) {
	((QSortFilterProxyModel__Go*)self)->endFilterChange();
}
void QSortFilterProxyModel_invalidateFilter(void* self) {
	((QSortFilterProxyModel__Go*)self)->invalidateFilter();
}
void QSortFilterProxyModel_invalidateRowsFilter(void* self) {
	((QSortFilterProxyModel__Go*)self)->invalidateRowsFilter();
}
void QSortFilterProxyModel_invalidateColumnsFilter(void* self) {
	((QSortFilterProxyModel__Go*)self)->invalidateColumnsFilter();
}
void QSortFilterProxyModel_dynamicSortFilterChanged(void* self, int dynamicSortFilter) {
	((QSortFilterProxyModel*)self)->dynamicSortFilterChanged(dynamicSortFilter);
}
void QSortFilterProxyModel_filterCaseSensitivityChanged(void* self, int filterCaseSensitivity) {
	((QSortFilterProxyModel*)self)->filterCaseSensitivityChanged((Qt::CaseSensitivity)(filterCaseSensitivity));
}
void QSortFilterProxyModel_sortCaseSensitivityChanged(void* self, int sortCaseSensitivity) {
	((QSortFilterProxyModel*)self)->sortCaseSensitivityChanged((Qt::CaseSensitivity)(sortCaseSensitivity));
}
void QSortFilterProxyModel_sortLocaleAwareChanged(void* self, int sortLocaleAware) {
	((QSortFilterProxyModel*)self)->sortLocaleAwareChanged(sortLocaleAware);
}
void QSortFilterProxyModel_sortRoleChanged(void* self, int sortRole) {
	((QSortFilterProxyModel*)self)->sortRoleChanged(sortRole);
}
void QSortFilterProxyModel_filterRoleChanged(void* self, int filterRole) {
	((QSortFilterProxyModel*)self)->filterRoleChanged(filterRole);
}
void QSortFilterProxyModel_recursiveFilteringEnabledChanged(void* self, int recursiveFilteringEnabled) {
	((QSortFilterProxyModel*)self)->recursiveFilteringEnabledChanged(recursiveFilteringEnabled);
}
void QSortFilterProxyModel_autoAcceptChildRowsChanged(void* self, int autoAcceptChildRows) {
	((QSortFilterProxyModel*)self)->autoAcceptChildRowsChanged(autoAcceptChildRows);
}
void* QSortFilterProxyModel_createIndex(void* self, int row, int column) {
	auto&& __r = ((QSortFilterProxyModel__Go*)self)->createIndex(row, column);
	return new QModelIndex(__r);
}
void* QSortFilterProxyModel_createIndex2(void* self, int row, int column, unsigned long long id) {
	auto&& __r = ((QSortFilterProxyModel__Go*)self)->createIndex(row, column, id);
	return new QModelIndex(__r);
}
void QSortFilterProxyModel_encodeData(void* self, void** indexes, long long indexes_len, void* stream) {
	QList<QModelIndex> indexes_list; indexes_list.reserve(indexes_len);
	for (long long __i = 0; __i < indexes_len; __i++) indexes_list.append(*(QModelIndex*)indexes[__i]);
	((QSortFilterProxyModel__Go*)self)->encodeData(indexes_list, *(QDataStream*)stream);
}
int QSortFilterProxyModel_decodeData(void* self, int row, int column, void* parent, void* stream) {
	auto&& __r = ((QSortFilterProxyModel__Go*)self)->decodeData(row, column, *(QModelIndex*)parent, *(QDataStream*)stream);
	return __r;
}
void QSortFilterProxyModel_beginInsertRows(void* self, void* parent, int first, int last) {
	((QSortFilterProxyModel__Go*)self)->beginInsertRows(*(QModelIndex*)parent, first, last);
}
void QSortFilterProxyModel_endInsertRows(void* self) {
	((QSortFilterProxyModel__Go*)self)->endInsertRows();
}
void QSortFilterProxyModel_beginRemoveRows(void* self, void* parent, int first, int last) {
	((QSortFilterProxyModel__Go*)self)->beginRemoveRows(*(QModelIndex*)parent, first, last);
}
void QSortFilterProxyModel_endRemoveRows(void* self) {
	((QSortFilterProxyModel__Go*)self)->endRemoveRows();
}
int QSortFilterProxyModel_beginMoveRows(void* self, void* sourceParent, int sourceFirst, int sourceLast, void* destinationParent, int destinationRow) {
	auto&& __r = ((QSortFilterProxyModel__Go*)self)->beginMoveRows(*(QModelIndex*)sourceParent, sourceFirst, sourceLast, *(QModelIndex*)destinationParent, destinationRow);
	return __r;
}
void QSortFilterProxyModel_endMoveRows(void* self) {
	((QSortFilterProxyModel__Go*)self)->endMoveRows();
}
void QSortFilterProxyModel_beginInsertColumns(void* self, void* parent, int first, int last) {
	((QSortFilterProxyModel__Go*)self)->beginInsertColumns(*(QModelIndex*)parent, first, last);
}
void QSortFilterProxyModel_endInsertColumns(void* self) {
	((QSortFilterProxyModel__Go*)self)->endInsertColumns();
}
void QSortFilterProxyModel_beginRemoveColumns(void* self, void* parent, int first, int last) {
	((QSortFilterProxyModel__Go*)self)->beginRemoveColumns(*(QModelIndex*)parent, first, last);
}
void QSortFilterProxyModel_endRemoveColumns(void* self) {
	((QSortFilterProxyModel__Go*)self)->endRemoveColumns();
}
int QSortFilterProxyModel_beginMoveColumns(void* self, void* sourceParent, int sourceFirst, int sourceLast, void* destinationParent, int destinationColumn) {
	auto&& __r = ((QSortFilterProxyModel__Go*)self)->beginMoveColumns(*(QModelIndex*)sourceParent, sourceFirst, sourceLast, *(QModelIndex*)destinationParent, destinationColumn);
	return __r;
}
void QSortFilterProxyModel_endMoveColumns(void* self) {
	((QSortFilterProxyModel__Go*)self)->endMoveColumns();
}
void QSortFilterProxyModel_beginResetModel(void* self) {
	((QSortFilterProxyModel__Go*)self)->beginResetModel();
}
void QSortFilterProxyModel_endResetModel(void* self) {
	((QSortFilterProxyModel__Go*)self)->endResetModel();
}
void QSortFilterProxyModel_changePersistentIndex(void* self, void* from, void* to) {
	((QSortFilterProxyModel__Go*)self)->changePersistentIndex(*(QModelIndex*)from, *(QModelIndex*)to);
}
void QSortFilterProxyModel_changePersistentIndexList(void* self, void** from, long long from_len, void** to, long long to_len) {
	QList<QModelIndex> from_list; from_list.reserve(from_len);
	for (long long __i = 0; __i < from_len; __i++) from_list.append(*(QModelIndex*)from[__i]);
	QList<QModelIndex> to_list; to_list.reserve(to_len);
	for (long long __i = 0; __i < to_len; __i++) to_list.append(*(QModelIndex*)to[__i]);
	((QSortFilterProxyModel__Go*)self)->changePersistentIndexList(from_list, to_list);
}
void** QSortFilterProxyModel_persistentIndexList(void* self, long long* __out_len) {
	auto&& __r = ((QSortFilterProxyModel__Go*)self)->persistentIndexList();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QModelIndex(__e);
	return __arr;
}
void* QSortFilterProxyModel_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QSortFilterProxyModel::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QSortFilterProxyModel_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QSortFilterProxyModel::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QSortFilterProxyModel_sender(void* self) {
	auto&& __r = ((QSortFilterProxyModel__Go*)self)->sender();
	return (void*)(__r);
}
int QSortFilterProxyModel_senderSignalIndex(void* self) {
	auto&& __r = ((QSortFilterProxyModel__Go*)self)->senderSignalIndex();
	return __r;
}
int QSortFilterProxyModel_receivers(void* self, const char* signal) {
	auto&& __r = ((QSortFilterProxyModel__Go*)self)->receivers(signal);
	return __r;
}
int QSortFilterProxyModel_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QSortFilterProxyModel__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QSortFilterProxyModel_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QSortFilterProxyModel::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QSortFilterProxyModel_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QSortFilterProxyModel::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QSortFilterProxyModel_endFilterChange2(void* self, int directions) {
	((QSortFilterProxyModel__Go*)self)->endFilterChange(QFlags<QSortFilterProxyModel::Direction>(directions));
}
void* QSortFilterProxyModel_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QSortFilterProxyModel::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QSortFilterProxyModel_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QSortFilterProxyModel::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QSortFilterProxyModel_metaObject_default(void* self) {
	auto&& __r = ((QSortFilterProxyModel*)self)->QSortFilterProxyModel::metaObject();
	return (void*)(__r);
}
void QSortFilterProxyModel_setSourceModel_default(void* self, void* sourceModel) {
	((QSortFilterProxyModel*)self)->QSortFilterProxyModel::setSourceModel((QAbstractItemModel*)sourceModel);
}
void* QSortFilterProxyModel_mapToSource_default(void* self, void* proxyIndex) {
	auto&& __r = ((QSortFilterProxyModel*)self)->QSortFilterProxyModel::mapToSource(*(QModelIndex*)proxyIndex);
	return new QModelIndex(__r);
}
void* QSortFilterProxyModel_mapFromSource_default(void* self, void* sourceIndex) {
	auto&& __r = ((QSortFilterProxyModel*)self)->QSortFilterProxyModel::mapFromSource(*(QModelIndex*)sourceIndex);
	return new QModelIndex(__r);
}
void* QSortFilterProxyModel_mapSelectionToSource_default(void* self, void* proxySelection) {
	auto&& __r = ((QSortFilterProxyModel*)self)->QSortFilterProxyModel::mapSelectionToSource(*(QItemSelection*)proxySelection);
	return new QItemSelection(__r);
}
void* QSortFilterProxyModel_mapSelectionFromSource_default(void* self, void* sourceSelection) {
	auto&& __r = ((QSortFilterProxyModel*)self)->QSortFilterProxyModel::mapSelectionFromSource(*(QItemSelection*)sourceSelection);
	return new QItemSelection(__r);
}
int QSortFilterProxyModel_filterAcceptsRow_default(void* self, int source_row, void* source_parent) {
	auto&& __r = ((QSortFilterProxyModel__Go*)self)->cppgoBase_filterAcceptsRow(source_row, *(QModelIndex*)source_parent);
	return __r;
}
int QSortFilterProxyModel_filterAcceptsColumn_default(void* self, int source_column, void* source_parent) {
	auto&& __r = ((QSortFilterProxyModel__Go*)self)->cppgoBase_filterAcceptsColumn(source_column, *(QModelIndex*)source_parent);
	return __r;
}
int QSortFilterProxyModel_lessThan_default(void* self, void* source_left, void* source_right) {
	auto&& __r = ((QSortFilterProxyModel__Go*)self)->cppgoBase_lessThan(*(QModelIndex*)source_left, *(QModelIndex*)source_right);
	return __r;
}
void* QSortFilterProxyModel_index_default(void* self, int row, int column, void* parent) {
	auto&& __r = ((QSortFilterProxyModel*)self)->QSortFilterProxyModel::index(row, column, *(QModelIndex*)parent);
	return new QModelIndex(__r);
}
void* QSortFilterProxyModel_parent_default(void* self, void* child) {
	auto&& __r = ((QSortFilterProxyModel*)self)->QSortFilterProxyModel::parent(*(QModelIndex*)child);
	return new QModelIndex(__r);
}
void* QSortFilterProxyModel_sibling_default(void* self, int row, int column, void* idx) {
	auto&& __r = ((QSortFilterProxyModel*)self)->QSortFilterProxyModel::sibling(row, column, *(QModelIndex*)idx);
	return new QModelIndex(__r);
}
int QSortFilterProxyModel_rowCount_default(void* self, void* parent) {
	auto&& __r = ((QSortFilterProxyModel*)self)->QSortFilterProxyModel::rowCount(*(QModelIndex*)parent);
	return __r;
}
int QSortFilterProxyModel_columnCount_default(void* self, void* parent) {
	auto&& __r = ((QSortFilterProxyModel*)self)->QSortFilterProxyModel::columnCount(*(QModelIndex*)parent);
	return __r;
}
int QSortFilterProxyModel_hasChildren_default(void* self, void* parent) {
	auto&& __r = ((QSortFilterProxyModel*)self)->QSortFilterProxyModel::hasChildren(*(QModelIndex*)parent);
	return __r;
}
void* QSortFilterProxyModel_data_default(void* self, void* index, int role) {
	auto&& __r = ((QSortFilterProxyModel*)self)->QSortFilterProxyModel::data(*(QModelIndex*)index, role);
	return new QVariant(__r);
}
int QSortFilterProxyModel_setData_default(void* self, void* index, void* value, int role) {
	auto&& __r = ((QSortFilterProxyModel*)self)->QSortFilterProxyModel::setData(*(QModelIndex*)index, *(QVariant*)value, role);
	return __r;
}
void* QSortFilterProxyModel_headerData_default(void* self, int section, int orientation, int role) {
	auto&& __r = ((QSortFilterProxyModel*)self)->QSortFilterProxyModel::headerData(section, (Qt::Orientation)(orientation), role);
	return new QVariant(__r);
}
int QSortFilterProxyModel_setHeaderData_default(void* self, int section, int orientation, void* value, int role) {
	auto&& __r = ((QSortFilterProxyModel*)self)->QSortFilterProxyModel::setHeaderData(section, (Qt::Orientation)(orientation), *(QVariant*)value, role);
	return __r;
}
void* QSortFilterProxyModel_mimeData_default(void* self, void** indexes, long long indexes_len) {
	QList<QModelIndex> indexes_list; indexes_list.reserve(indexes_len);
	for (long long __i = 0; __i < indexes_len; __i++) indexes_list.append(*(QModelIndex*)indexes[__i]);
	auto&& __r = ((QSortFilterProxyModel*)self)->QSortFilterProxyModel::mimeData(indexes_list);
	return (void*)(__r);
}
int QSortFilterProxyModel_dropMimeData_default(void* self, void* data, int action, int row, int column, void* parent) {
	auto&& __r = ((QSortFilterProxyModel*)self)->QSortFilterProxyModel::dropMimeData((QMimeData*)data, (Qt::DropAction)(action), row, column, *(QModelIndex*)parent);
	return __r;
}
int QSortFilterProxyModel_insertRows_default(void* self, int row, int count, void* parent) {
	auto&& __r = ((QSortFilterProxyModel*)self)->QSortFilterProxyModel::insertRows(row, count, *(QModelIndex*)parent);
	return __r;
}
int QSortFilterProxyModel_insertColumns_default(void* self, int column, int count, void* parent) {
	auto&& __r = ((QSortFilterProxyModel*)self)->QSortFilterProxyModel::insertColumns(column, count, *(QModelIndex*)parent);
	return __r;
}
int QSortFilterProxyModel_removeRows_default(void* self, int row, int count, void* parent) {
	auto&& __r = ((QSortFilterProxyModel*)self)->QSortFilterProxyModel::removeRows(row, count, *(QModelIndex*)parent);
	return __r;
}
int QSortFilterProxyModel_removeColumns_default(void* self, int column, int count, void* parent) {
	auto&& __r = ((QSortFilterProxyModel*)self)->QSortFilterProxyModel::removeColumns(column, count, *(QModelIndex*)parent);
	return __r;
}
void QSortFilterProxyModel_fetchMore_default(void* self, void* parent) {
	((QSortFilterProxyModel*)self)->QSortFilterProxyModel::fetchMore(*(QModelIndex*)parent);
}
int QSortFilterProxyModel_canFetchMore_default(void* self, void* parent) {
	auto&& __r = ((QSortFilterProxyModel*)self)->QSortFilterProxyModel::canFetchMore(*(QModelIndex*)parent);
	return __r;
}
int QSortFilterProxyModel_flags_default(void* self, void* index) {
	auto&& __r = ((QSortFilterProxyModel*)self)->QSortFilterProxyModel::flags(*(QModelIndex*)index);
	return (int)(__r);
}
void* QSortFilterProxyModel_buddy_default(void* self, void* index) {
	auto&& __r = ((QSortFilterProxyModel*)self)->QSortFilterProxyModel::buddy(*(QModelIndex*)index);
	return new QModelIndex(__r);
}
void** QSortFilterProxyModel_match_default(void* self, void* start, int role, void* value, int hits, int flags, long long* __out_len) {
	auto&& __r = ((QSortFilterProxyModel*)self)->QSortFilterProxyModel::match(*(QModelIndex*)start, role, *(QVariant*)value, hits, QFlags<Qt::MatchFlag>(flags));
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QModelIndex(__e);
	return __arr;
}
void* QSortFilterProxyModel_span_default(void* self, void* index) {
	auto&& __r = ((QSortFilterProxyModel*)self)->QSortFilterProxyModel::span(*(QModelIndex*)index);
	return new QSize(__r);
}
void QSortFilterProxyModel_sort_default(void* self, int column, int order) {
	((QSortFilterProxyModel*)self)->QSortFilterProxyModel::sort(column, (Qt::SortOrder)(order));
}
char** QSortFilterProxyModel_mimeTypes_default(void* self, long long* __out_len) {
	auto&& __r = ((QSortFilterProxyModel*)self)->QSortFilterProxyModel::mimeTypes();
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
int QSortFilterProxyModel_supportedDropActions_default(void* self) {
	auto&& __r = ((QSortFilterProxyModel*)self)->QSortFilterProxyModel::supportedDropActions();
	return (int)(__r);
}
int QSortFilterProxyModel_submit_default(void* self) {
	auto&& __r = ((QSortFilterProxyModel*)self)->QSortFilterProxyModel::submit();
	return __r;
}
void QSortFilterProxyModel_revert_default(void* self) {
	((QSortFilterProxyModel*)self)->QSortFilterProxyModel::revert();
}
int QSortFilterProxyModel_clearItemData_default(void* self, void* index) {
	auto&& __r = ((QSortFilterProxyModel*)self)->QSortFilterProxyModel::clearItemData(*(QModelIndex*)index);
	return __r;
}
int QSortFilterProxyModel_canDropMimeData_default(void* self, void* data, int action, int row, int column, void* parent) {
	auto&& __r = ((QSortFilterProxyModel*)self)->QSortFilterProxyModel::canDropMimeData((QMimeData*)data, (Qt::DropAction)(action), row, column, *(QModelIndex*)parent);
	return __r;
}
int QSortFilterProxyModel_supportedDragActions_default(void* self) {
	auto&& __r = ((QSortFilterProxyModel*)self)->QSortFilterProxyModel::supportedDragActions();
	return (int)(__r);
}
int QSortFilterProxyModel_moveRows_default(void* self, void* sourceParent, int sourceRow, int count, void* destinationParent, int destinationChild) {
	auto&& __r = ((QSortFilterProxyModel*)self)->QSortFilterProxyModel::moveRows(*(QModelIndex*)sourceParent, sourceRow, count, *(QModelIndex*)destinationParent, destinationChild);
	return __r;
}
int QSortFilterProxyModel_moveColumns_default(void* self, void* sourceParent, int sourceColumn, int count, void* destinationParent, int destinationChild) {
	auto&& __r = ((QSortFilterProxyModel*)self)->QSortFilterProxyModel::moveColumns(*(QModelIndex*)sourceParent, sourceColumn, count, *(QModelIndex*)destinationParent, destinationChild);
	return __r;
}
void QSortFilterProxyModel_multiData_default(void* self, void* index, void* roleDataSpan) {
	((QSortFilterProxyModel*)self)->QSortFilterProxyModel::multiData(*(QModelIndex*)index, *(QModelRoleDataSpan*)roleDataSpan);
}
void QSortFilterProxyModel_resetInternalData_default(void* self) {
	((QSortFilterProxyModel__Go*)self)->cppgoBase_resetInternalData();
}
int QSortFilterProxyModel_event_default(void* self, void* event) {
	auto&& __r = ((QSortFilterProxyModel*)self)->QSortFilterProxyModel::event((QEvent*)event);
	return __r;
}
int QSortFilterProxyModel_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QSortFilterProxyModel*)self)->QSortFilterProxyModel::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QSortFilterProxyModel_timerEvent_default(void* self, void* event) {
	((QSortFilterProxyModel__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QSortFilterProxyModel_childEvent_default(void* self, void* event) {
	((QSortFilterProxyModel__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QSortFilterProxyModel_customEvent_default(void* self, void* event) {
	((QSortFilterProxyModel__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QSortFilterProxyModel_connectNotify_default(void* self, void* signal) {
	((QSortFilterProxyModel__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QSortFilterProxyModel_disconnectNotify_default(void* self, void* signal) {
	((QSortFilterProxyModel__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QSortFilterProxyModel_connect_dynamicSortFilterChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSortFilterProxyModel*)self, &QSortFilterProxyModel::dynamicSortFilterChanged, (QSortFilterProxyModel*)self, [__id](bool dynamicSortFilter){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, dynamicSortFilter);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSortFilterProxyModel);
	__conns_QSortFilterProxyModel[__id] = __c;
}
void QSortFilterProxyModel_connect_filterCaseSensitivityChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSortFilterProxyModel*)self, &QSortFilterProxyModel::filterCaseSensitivityChanged, (QSortFilterProxyModel*)self, [__id](Qt::CaseSensitivity filterCaseSensitivity){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, (int)(filterCaseSensitivity));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSortFilterProxyModel);
	__conns_QSortFilterProxyModel[__id] = __c;
}
void QSortFilterProxyModel_connect_sortCaseSensitivityChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSortFilterProxyModel*)self, &QSortFilterProxyModel::sortCaseSensitivityChanged, (QSortFilterProxyModel*)self, [__id](Qt::CaseSensitivity sortCaseSensitivity){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, (int)(sortCaseSensitivity));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSortFilterProxyModel);
	__conns_QSortFilterProxyModel[__id] = __c;
}
void QSortFilterProxyModel_connect_sortLocaleAwareChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSortFilterProxyModel*)self, &QSortFilterProxyModel::sortLocaleAwareChanged, (QSortFilterProxyModel*)self, [__id](bool sortLocaleAware){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, sortLocaleAware);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSortFilterProxyModel);
	__conns_QSortFilterProxyModel[__id] = __c;
}
void QSortFilterProxyModel_connect_sortRoleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSortFilterProxyModel*)self, &QSortFilterProxyModel::sortRoleChanged, (QSortFilterProxyModel*)self, [__id](int sortRole){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, sortRole);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSortFilterProxyModel);
	__conns_QSortFilterProxyModel[__id] = __c;
}
void QSortFilterProxyModel_connect_filterRoleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSortFilterProxyModel*)self, &QSortFilterProxyModel::filterRoleChanged, (QSortFilterProxyModel*)self, [__id](int filterRole){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, filterRole);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSortFilterProxyModel);
	__conns_QSortFilterProxyModel[__id] = __c;
}
void QSortFilterProxyModel_connect_recursiveFilteringEnabledChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSortFilterProxyModel*)self, &QSortFilterProxyModel::recursiveFilteringEnabledChanged, (QSortFilterProxyModel*)self, [__id](bool recursiveFilteringEnabled){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, recursiveFilteringEnabled);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSortFilterProxyModel);
	__conns_QSortFilterProxyModel[__id] = __c;
}
void QSortFilterProxyModel_connect_autoAcceptChildRowsChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSortFilterProxyModel*)self, &QSortFilterProxyModel::autoAcceptChildRowsChanged, (QSortFilterProxyModel*)self, [__id](bool autoAcceptChildRows){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, autoAcceptChildRows);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSortFilterProxyModel);
	__conns_QSortFilterProxyModel[__id] = __c;
}
void QSortFilterProxyModel_connect_sourceModelChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSortFilterProxyModel*)self, &QSortFilterProxyModel::sourceModelChanged, (QSortFilterProxyModel*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSortFilterProxyModel);
	__conns_QSortFilterProxyModel[__id] = __c;
}
void QSortFilterProxyModel_connect_headerDataChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSortFilterProxyModel*)self, &QSortFilterProxyModel::headerDataChanged, (QSortFilterProxyModel*)self, [__id](Qt::Orientation orientation, int first, int last){
		if (cppgo_sig_void_int_int_int) ((void(*)(long long, int, int, int))cppgo_sig_void_int_int_int)(__id, (int)(orientation), first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSortFilterProxyModel);
	__conns_QSortFilterProxyModel[__id] = __c;
}
void QSortFilterProxyModel_connect_layoutChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSortFilterProxyModel*)self, &QSortFilterProxyModel::layoutChanged, (QSortFilterProxyModel*)self, [__id](const QList<QPersistentModelIndex>& parents, QAbstractItemModel::LayoutChangeHint hint){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSortFilterProxyModel);
	__conns_QSortFilterProxyModel[__id] = __c;
}
void QSortFilterProxyModel_connect_layoutChanged2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSortFilterProxyModel*)self, &QSortFilterProxyModel::layoutChanged, (QSortFilterProxyModel*)self, [__id](const QList<QPersistentModelIndex>& parents, QAbstractItemModel::LayoutChangeHint hint){
		std::vector<void*> parents_v; parents_v.reserve(parents.size());
		for (long long __i = 0; __i < parents.size(); __i++) parents_v.push_back((void*)&parents[__i]);
		if (cppgo_sig_void_voidpp_longlong_int) ((void(*)(long long, void**, long long, int))cppgo_sig_void_voidpp_longlong_int)(__id, parents_v.data(), (long long)parents.size(), (int)(hint));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSortFilterProxyModel);
	__conns_QSortFilterProxyModel[__id] = __c;
}
void QSortFilterProxyModel_connect_layoutAboutToBeChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSortFilterProxyModel*)self, &QSortFilterProxyModel::layoutAboutToBeChanged, (QSortFilterProxyModel*)self, [__id](const QList<QPersistentModelIndex>& parents, QAbstractItemModel::LayoutChangeHint hint){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSortFilterProxyModel);
	__conns_QSortFilterProxyModel[__id] = __c;
}
void QSortFilterProxyModel_connect_layoutAboutToBeChanged2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSortFilterProxyModel*)self, &QSortFilterProxyModel::layoutAboutToBeChanged, (QSortFilterProxyModel*)self, [__id](const QList<QPersistentModelIndex>& parents, QAbstractItemModel::LayoutChangeHint hint){
		std::vector<void*> parents_v; parents_v.reserve(parents.size());
		for (long long __i = 0; __i < parents.size(); __i++) parents_v.push_back((void*)&parents[__i]);
		if (cppgo_sig_void_voidpp_longlong_int) ((void(*)(long long, void**, long long, int))cppgo_sig_void_voidpp_longlong_int)(__id, parents_v.data(), (long long)parents.size(), (int)(hint));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSortFilterProxyModel);
	__conns_QSortFilterProxyModel[__id] = __c;
}
void QSortFilterProxyModel_connect_rowsAboutToBeInserted(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSortFilterProxyModel*)self, &QSortFilterProxyModel::rowsAboutToBeInserted, (QSortFilterProxyModel*)self, [__id](const QModelIndex& parent, int first, int last){
		if (cppgo_sig_void_voidp_int_int) ((void(*)(long long, void*, int, int))cppgo_sig_void_voidp_int_int)(__id, (void*)&parent, first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSortFilterProxyModel);
	__conns_QSortFilterProxyModel[__id] = __c;
}
void QSortFilterProxyModel_connect_rowsInserted(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSortFilterProxyModel*)self, &QSortFilterProxyModel::rowsInserted, (QSortFilterProxyModel*)self, [__id](const QModelIndex& parent, int first, int last){
		if (cppgo_sig_void_voidp_int_int) ((void(*)(long long, void*, int, int))cppgo_sig_void_voidp_int_int)(__id, (void*)&parent, first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSortFilterProxyModel);
	__conns_QSortFilterProxyModel[__id] = __c;
}
void QSortFilterProxyModel_connect_rowsAboutToBeRemoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSortFilterProxyModel*)self, &QSortFilterProxyModel::rowsAboutToBeRemoved, (QSortFilterProxyModel*)self, [__id](const QModelIndex& parent, int first, int last){
		if (cppgo_sig_void_voidp_int_int) ((void(*)(long long, void*, int, int))cppgo_sig_void_voidp_int_int)(__id, (void*)&parent, first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSortFilterProxyModel);
	__conns_QSortFilterProxyModel[__id] = __c;
}
void QSortFilterProxyModel_connect_rowsRemoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSortFilterProxyModel*)self, &QSortFilterProxyModel::rowsRemoved, (QSortFilterProxyModel*)self, [__id](const QModelIndex& parent, int first, int last){
		if (cppgo_sig_void_voidp_int_int) ((void(*)(long long, void*, int, int))cppgo_sig_void_voidp_int_int)(__id, (void*)&parent, first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSortFilterProxyModel);
	__conns_QSortFilterProxyModel[__id] = __c;
}
void QSortFilterProxyModel_connect_columnsAboutToBeInserted(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSortFilterProxyModel*)self, &QSortFilterProxyModel::columnsAboutToBeInserted, (QSortFilterProxyModel*)self, [__id](const QModelIndex& parent, int first, int last){
		if (cppgo_sig_void_voidp_int_int) ((void(*)(long long, void*, int, int))cppgo_sig_void_voidp_int_int)(__id, (void*)&parent, first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSortFilterProxyModel);
	__conns_QSortFilterProxyModel[__id] = __c;
}
void QSortFilterProxyModel_connect_columnsInserted(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSortFilterProxyModel*)self, &QSortFilterProxyModel::columnsInserted, (QSortFilterProxyModel*)self, [__id](const QModelIndex& parent, int first, int last){
		if (cppgo_sig_void_voidp_int_int) ((void(*)(long long, void*, int, int))cppgo_sig_void_voidp_int_int)(__id, (void*)&parent, first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSortFilterProxyModel);
	__conns_QSortFilterProxyModel[__id] = __c;
}
void QSortFilterProxyModel_connect_columnsAboutToBeRemoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSortFilterProxyModel*)self, &QSortFilterProxyModel::columnsAboutToBeRemoved, (QSortFilterProxyModel*)self, [__id](const QModelIndex& parent, int first, int last){
		if (cppgo_sig_void_voidp_int_int) ((void(*)(long long, void*, int, int))cppgo_sig_void_voidp_int_int)(__id, (void*)&parent, first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSortFilterProxyModel);
	__conns_QSortFilterProxyModel[__id] = __c;
}
void QSortFilterProxyModel_connect_columnsRemoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSortFilterProxyModel*)self, &QSortFilterProxyModel::columnsRemoved, (QSortFilterProxyModel*)self, [__id](const QModelIndex& parent, int first, int last){
		if (cppgo_sig_void_voidp_int_int) ((void(*)(long long, void*, int, int))cppgo_sig_void_voidp_int_int)(__id, (void*)&parent, first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSortFilterProxyModel);
	__conns_QSortFilterProxyModel[__id] = __c;
}
void QSortFilterProxyModel_connect_modelAboutToBeReset(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSortFilterProxyModel*)self, &QSortFilterProxyModel::modelAboutToBeReset, (QSortFilterProxyModel*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSortFilterProxyModel);
	__conns_QSortFilterProxyModel[__id] = __c;
}
void QSortFilterProxyModel_connect_modelReset(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSortFilterProxyModel*)self, &QSortFilterProxyModel::modelReset, (QSortFilterProxyModel*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSortFilterProxyModel);
	__conns_QSortFilterProxyModel[__id] = __c;
}
void QSortFilterProxyModel_connect_rowsAboutToBeMoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSortFilterProxyModel*)self, &QSortFilterProxyModel::rowsAboutToBeMoved, (QSortFilterProxyModel*)self, [__id](const QModelIndex& sourceParent, int sourceStart, int sourceEnd, const QModelIndex& destinationParent, int destinationRow){
		if (cppgo_sig_void_voidp_int_int_voidp_int) ((void(*)(long long, void*, int, int, void*, int))cppgo_sig_void_voidp_int_int_voidp_int)(__id, (void*)&sourceParent, sourceStart, sourceEnd, (void*)&destinationParent, destinationRow);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSortFilterProxyModel);
	__conns_QSortFilterProxyModel[__id] = __c;
}
void QSortFilterProxyModel_connect_rowsMoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSortFilterProxyModel*)self, &QSortFilterProxyModel::rowsMoved, (QSortFilterProxyModel*)self, [__id](const QModelIndex& sourceParent, int sourceStart, int sourceEnd, const QModelIndex& destinationParent, int destinationRow){
		if (cppgo_sig_void_voidp_int_int_voidp_int) ((void(*)(long long, void*, int, int, void*, int))cppgo_sig_void_voidp_int_int_voidp_int)(__id, (void*)&sourceParent, sourceStart, sourceEnd, (void*)&destinationParent, destinationRow);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSortFilterProxyModel);
	__conns_QSortFilterProxyModel[__id] = __c;
}
void QSortFilterProxyModel_connect_columnsAboutToBeMoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSortFilterProxyModel*)self, &QSortFilterProxyModel::columnsAboutToBeMoved, (QSortFilterProxyModel*)self, [__id](const QModelIndex& sourceParent, int sourceStart, int sourceEnd, const QModelIndex& destinationParent, int destinationColumn){
		if (cppgo_sig_void_voidp_int_int_voidp_int) ((void(*)(long long, void*, int, int, void*, int))cppgo_sig_void_voidp_int_int_voidp_int)(__id, (void*)&sourceParent, sourceStart, sourceEnd, (void*)&destinationParent, destinationColumn);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSortFilterProxyModel);
	__conns_QSortFilterProxyModel[__id] = __c;
}
void QSortFilterProxyModel_connect_columnsMoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSortFilterProxyModel*)self, &QSortFilterProxyModel::columnsMoved, (QSortFilterProxyModel*)self, [__id](const QModelIndex& sourceParent, int sourceStart, int sourceEnd, const QModelIndex& destinationParent, int destinationColumn){
		if (cppgo_sig_void_voidp_int_int_voidp_int) ((void(*)(long long, void*, int, int, void*, int))cppgo_sig_void_voidp_int_int_voidp_int)(__id, (void*)&sourceParent, sourceStart, sourceEnd, (void*)&destinationParent, destinationColumn);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSortFilterProxyModel);
	__conns_QSortFilterProxyModel[__id] = __c;
}
void QSortFilterProxyModel_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSortFilterProxyModel*)self, &QSortFilterProxyModel::destroyed, (QSortFilterProxyModel*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSortFilterProxyModel);
	__conns_QSortFilterProxyModel[__id] = __c;
}
void QSortFilterProxyModel_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSortFilterProxyModel*)self, &QSortFilterProxyModel::destroyed, (QSortFilterProxyModel*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSortFilterProxyModel);
	__conns_QSortFilterProxyModel[__id] = __c;
}
void QSortFilterProxyModel_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSortFilterProxyModel*)self, &QSortFilterProxyModel::objectNameChanged, (QSortFilterProxyModel*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSortFilterProxyModel);
	__conns_QSortFilterProxyModel[__id] = __c;
}
void QSortFilterProxyModel__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QSortFilterProxyModel);
		auto __it = __conns_QSortFilterProxyModel.find(id);
		if (__it == __conns_QSortFilterProxyModel.end()) return;
		__c = __it->second;
		__conns_QSortFilterProxyModel.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QSortFilterProxyModel__Go_new(long long id, void* parent) { return new QSortFilterProxyModel__Go(id, (QObject*)parent); }
void QSortFilterProxyModel_delete(void* self) { delete (QSortFilterProxyModel*)self; }

}

#include <QtGui/qrgba64.h>
#include "gen_qrgba64.h"

extern "C" {

void* QRgba64_new() {
	return new QRgba64();
}
void* QRgba64_new2(void* arg1) {
	return new QRgba64(*(QRgba64*)arg1);
}
void* QRgba64_fromRgba64(unsigned long long c) {
	auto&& __r = QRgba64::fromRgba64(c);
	return new QRgba64(__r);
}
void* QRgba64_fromRgba642(unsigned short red, unsigned short green, unsigned short blue, unsigned short alpha) {
	auto&& __r = QRgba64::fromRgba64(red, green, blue, alpha);
	return new QRgba64(__r);
}
void* QRgba64_fromRgba(unsigned char red, unsigned char green, unsigned char blue, unsigned char alpha) {
	auto&& __r = QRgba64::fromRgba(red, green, blue, alpha);
	return new QRgba64(__r);
}
void* QRgba64_fromArgb32(unsigned int rgb) {
	auto&& __r = QRgba64::fromArgb32(rgb);
	return new QRgba64(__r);
}
int QRgba64_isOpaque(void* self) {
	auto&& __r = ((QRgba64*)self)->isOpaque();
	return __r;
}
int QRgba64_isTransparent(void* self) {
	auto&& __r = ((QRgba64*)self)->isTransparent();
	return __r;
}
unsigned short QRgba64_red(void* self) {
	auto&& __r = ((QRgba64*)self)->red();
	return __r;
}
unsigned short QRgba64_green(void* self) {
	auto&& __r = ((QRgba64*)self)->green();
	return __r;
}
unsigned short QRgba64_blue(void* self) {
	auto&& __r = ((QRgba64*)self)->blue();
	return __r;
}
unsigned short QRgba64_alpha(void* self) {
	auto&& __r = ((QRgba64*)self)->alpha();
	return __r;
}
void QRgba64_setRed(void* self, unsigned short _red) {
	((QRgba64*)self)->setRed(_red);
}
void QRgba64_setGreen(void* self, unsigned short _green) {
	((QRgba64*)self)->setGreen(_green);
}
void QRgba64_setBlue(void* self, unsigned short _blue) {
	((QRgba64*)self)->setBlue(_blue);
}
void QRgba64_setAlpha(void* self, unsigned short _alpha) {
	((QRgba64*)self)->setAlpha(_alpha);
}
unsigned char QRgba64_red8(void* self) {
	auto&& __r = ((QRgba64*)self)->red8();
	return __r;
}
unsigned char QRgba64_green8(void* self) {
	auto&& __r = ((QRgba64*)self)->green8();
	return __r;
}
unsigned char QRgba64_blue8(void* self) {
	auto&& __r = ((QRgba64*)self)->blue8();
	return __r;
}
unsigned char QRgba64_alpha8(void* self) {
	auto&& __r = ((QRgba64*)self)->alpha8();
	return __r;
}
unsigned int QRgba64_toArgb32(void* self) {
	auto&& __r = ((QRgba64*)self)->toArgb32();
	return __r;
}
unsigned short QRgba64_toRgb16(void* self) {
	auto&& __r = ((QRgba64*)self)->toRgb16();
	return __r;
}
void* QRgba64_premultiplied(void* self) {
	auto&& __r = ((QRgba64*)self)->premultiplied();
	return new QRgba64(__r);
}
void* QRgba64_unpremultiplied(void* self) {
	auto&& __r = ((QRgba64*)self)->unpremultiplied();
	return new QRgba64(__r);
}
unsigned long long QRgba64_ToQuint64(void* self) {
	auto&& __r = static_cast<unsigned long long>(*(QRgba64*)self);
	return __r;
}
void* QRgba64_OperatorAssign(void* self, unsigned long long _rgba) {
	auto&& __r = ((QRgba64*)self)->operator=(_rgba);
	return (void*)&(__r);
}
void QRgba64_delete(void* self) { delete (QRgba64*)self; }

}

#include <QtGui/qevent.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qpoint.h>
#include <QtCore/qobject.h>
#include <QtGui/qpointingdevice.h>
#include <QtGui/qeventpoint.h>
#include <QtGui/qinputdevice.h>
#include <QtCore/qcoreevent.h>
#include "gen_qsinglepointevent.h"

extern "C" {

void* QSinglePointEvent_base(void* self) { return (void*)static_cast<QPointerEvent*>((QSinglePointEvent*)self); }
void QSinglePointEvent_chain(void* self, void** out) {
	QSinglePointEvent* __p = (QSinglePointEvent*)self;
	QPointerEvent* __b0 = static_cast<QPointerEvent*>(__p); out[0] = (void*)__b0;
	QInputEvent* __b1 = static_cast<QInputEvent*>(__b0); out[1] = (void*)__b1;
	QEvent* __b2 = static_cast<QEvent*>(__b1); out[2] = (void*)__b2;
}
int QSinglePointEvent_button(void* self) {
	auto&& __r = ((QSinglePointEvent*)self)->button();
	return (int)(__r);
}
int QSinglePointEvent_buttons(void* self) {
	auto&& __r = ((QSinglePointEvent*)self)->buttons();
	return (int)(__r);
}
void* QSinglePointEvent_position(void* self) {
	auto&& __r = ((QSinglePointEvent*)self)->position();
	return new QPointF(__r);
}
void* QSinglePointEvent_scenePosition(void* self) {
	auto&& __r = ((QSinglePointEvent*)self)->scenePosition();
	return new QPointF(__r);
}
void* QSinglePointEvent_globalPosition(void* self) {
	auto&& __r = ((QSinglePointEvent*)self)->globalPosition();
	return new QPointF(__r);
}
void* QSinglePointEvent_exclusivePointGrabber(void* self) {
	auto&& __r = ((QSinglePointEvent*)self)->exclusivePointGrabber();
	return (void*)(__r);
}
void QSinglePointEvent_setExclusivePointGrabber(void* self, void* exclusiveGrabber) {
	((QSinglePointEvent*)self)->setExclusivePointGrabber((QObject*)exclusiveGrabber);
}
int QSinglePointEvent_registerEventType() {
	auto&& __r = QSinglePointEvent::registerEventType();
	return __r;
}
int QSinglePointEvent_registerEventType2(int hint) {
	auto&& __r = QSinglePointEvent::registerEventType(hint);
	return __r;
}
void QSinglePointEvent_delete(void* self) { delete (QSinglePointEvent*)self; }

}

#include <QtGui/qevent.h>
#include <QtCore/qpoint.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qmimedata.h>
#include <QtCore/qrect.h>
#include <QtCore/qobject.h>
#include <QtCore/qcoreevent.h>
#include "gen_qdragenterevent.h"

extern "C" {

void* QDragEnterEvent_base(void* self) { return (void*)static_cast<QDragMoveEvent*>((QDragEnterEvent*)self); }
void QDragEnterEvent_chain(void* self, void** out) {
	QDragEnterEvent* __p = (QDragEnterEvent*)self;
	QDragMoveEvent* __b0 = static_cast<QDragMoveEvent*>(__p); out[0] = (void*)__b0;
	QDropEvent* __b1 = static_cast<QDropEvent*>(__b0); out[1] = (void*)__b1;
	QEvent* __b2 = static_cast<QEvent*>(__b1); out[2] = (void*)__b2;
}
void* QDragEnterEvent_new(void* pos, int actions, void* data, int buttons, int modifiers) {
	return new QDragEnterEvent(*(QPoint*)pos, QFlags<Qt::DropAction>(actions), (QMimeData*)data, QFlags<Qt::MouseButton>(buttons), QFlags<Qt::KeyboardModifier>(modifiers));
}
int QDragEnterEvent_registerEventType() {
	auto&& __r = QDragEnterEvent::registerEventType();
	return __r;
}
int QDragEnterEvent_registerEventType2(int hint) {
	auto&& __r = QDragEnterEvent::registerEventType(hint);
	return __r;
}
void QDragEnterEvent_delete(void* self) { delete (QDragEnterEvent*)self; }

}

#include <QtGui/qbrush.h>
#include <QtCore/qatomic.h>
#include <QtCore/qnamespace.h>
#include <QtGui/qcolor.h>
#include <QtGui/qtransform.h>
#include "gen_qbrushdata.h"

extern "C" {

void* QBrushData_ref_get(void* self) {
	auto __r = ((QBrushData*)self)->ref;
	return new QAtomicInt(__r);
}
void QBrushData_ref_set(void* self, void* ref) {
	((QBrushData*)self)->ref = *(QAtomicInt*)ref;
}
int QBrushData_style_get(void* self) {
	auto __r = ((QBrushData*)self)->style;
	return (int)(__r);
}
void QBrushData_style_set(void* self, int style) {
	((QBrushData*)self)->style = (Qt::BrushStyle)(style);
}
void* QBrushData_color_get(void* self) {
	auto __r = ((QBrushData*)self)->color;
	return new QColor(__r);
}
void QBrushData_color_set(void* self, void* color) {
	((QBrushData*)self)->color = *(QColor*)color;
}
void* QBrushData_transform_get(void* self) {
	auto __r = ((QBrushData*)self)->transform;
	return new QTransform(__r);
}
void QBrushData_transform_set(void* self, void* transform) {
	((QBrushData*)self)->transform = *(QTransform*)transform;
}

}

#include <QtGui/qabstracttextdocumentlayout.h>
#include <QtCore/qobject.h>
#include <QtGui/qtextdocument.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtGui/qpainter.h>
#include <QtCore/qpoint.h>
#include <QtCore/qnamespace.h>
#include <QtGui/qtextformat.h>
#include <QtGui/qtextobject.h>
#include <QtCore/qsize.h>
#include <QtCore/qrect.h>
#include <QtGui/qpaintdevice.h>
#include <QtGui/qtextlayout.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qvariant.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qabstracttextdocumentlayout.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

static std::unordered_map<long long, QMetaObject::Connection> __conns_QAbstractTextDocumentLayout;
static std::mutex __connsMu_QAbstractTextDocumentLayout;

extern "C" {

void* QAbstractTextDocumentLayout_base(void* self) { return (void*)static_cast<QObject*>((QAbstractTextDocumentLayout*)self); }
void* QAbstractTextDocumentLayout_fromQObject(void* o) { return (void*)qobject_cast<QAbstractTextDocumentLayout*>((QObject*)o); }
char* QAbstractTextDocumentLayout_tr(const char* s, long long* __out_len) {
	auto&& __r = QAbstractTextDocumentLayout::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QAbstractTextDocumentLayout_draw(void* self, void* painter, void* context) {
	((QAbstractTextDocumentLayout*)self)->draw((QPainter*)painter, *(QAbstractTextDocumentLayout::PaintContext*)context);
}
int QAbstractTextDocumentLayout_hitTest(void* self, void* point, int accuracy) {
	auto&& __r = ((QAbstractTextDocumentLayout*)self)->hitTest(*(QPointF*)point, (Qt::HitTestAccuracy)(accuracy));
	return __r;
}
char* QAbstractTextDocumentLayout_anchorAt(void* self, void* pos, long long* __out_len) {
	auto&& __r = ((QAbstractTextDocumentLayout*)self)->anchorAt(*(QPointF*)pos);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QAbstractTextDocumentLayout_imageAt(void* self, void* pos, long long* __out_len) {
	auto&& __r = ((QAbstractTextDocumentLayout*)self)->imageAt(*(QPointF*)pos);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QAbstractTextDocumentLayout_formatAt(void* self, void* pos) {
	auto&& __r = ((QAbstractTextDocumentLayout*)self)->formatAt(*(QPointF*)pos);
	return new QTextFormat(__r);
}
void* QAbstractTextDocumentLayout_blockWithMarkerAt(void* self, void* pos) {
	auto&& __r = ((QAbstractTextDocumentLayout*)self)->blockWithMarkerAt(*(QPointF*)pos);
	return new QTextBlock(__r);
}
int QAbstractTextDocumentLayout_pageCount(void* self) {
	auto&& __r = ((QAbstractTextDocumentLayout*)self)->pageCount();
	return __r;
}
void* QAbstractTextDocumentLayout_documentSize(void* self) {
	auto&& __r = ((QAbstractTextDocumentLayout*)self)->documentSize();
	return new QSizeF(__r);
}
void* QAbstractTextDocumentLayout_frameBoundingRect(void* self, void* frame) {
	auto&& __r = ((QAbstractTextDocumentLayout*)self)->frameBoundingRect((QTextFrame*)frame);
	return new QRectF(__r);
}
void* QAbstractTextDocumentLayout_blockBoundingRect(void* self, void* block) {
	auto&& __r = ((QAbstractTextDocumentLayout*)self)->blockBoundingRect(*(QTextBlock*)block);
	return new QRectF(__r);
}
void QAbstractTextDocumentLayout_setPaintDevice(void* self, void* device) {
	((QAbstractTextDocumentLayout*)self)->setPaintDevice((QPaintDevice*)device);
}
void* QAbstractTextDocumentLayout_paintDevice(void* self) {
	auto&& __r = ((QAbstractTextDocumentLayout*)self)->paintDevice();
	return (void*)(__r);
}
void* QAbstractTextDocumentLayout_document(void* self) {
	auto&& __r = ((QAbstractTextDocumentLayout*)self)->document();
	return (void*)(__r);
}
void QAbstractTextDocumentLayout_registerHandler(void* self, int objectType, void* component) {
	((QAbstractTextDocumentLayout*)self)->registerHandler(objectType, (QObject*)component);
}
void QAbstractTextDocumentLayout_unregisterHandler(void* self, int objectType) {
	((QAbstractTextDocumentLayout*)self)->unregisterHandler(objectType);
}
void* QAbstractTextDocumentLayout_handlerForObject(void* self, int objectType) {
	auto&& __r = ((QAbstractTextDocumentLayout*)self)->handlerForObject(objectType);
	return (void*)(__r);
}
void QAbstractTextDocumentLayout_update(void* self) {
	((QAbstractTextDocumentLayout*)self)->update();
}
void QAbstractTextDocumentLayout_updateBlock(void* self, void* block) {
	((QAbstractTextDocumentLayout*)self)->updateBlock(*(QTextBlock*)block);
}
void QAbstractTextDocumentLayout_documentSizeChanged(void* self, void* newSize) {
	((QAbstractTextDocumentLayout*)self)->documentSizeChanged(*(QSizeF*)newSize);
}
void QAbstractTextDocumentLayout_pageCountChanged(void* self, int newPages) {
	((QAbstractTextDocumentLayout*)self)->pageCountChanged(newPages);
}
void* QAbstractTextDocumentLayout_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QAbstractTextDocumentLayout::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QAbstractTextDocumentLayout_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QAbstractTextDocumentLayout::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
char* QAbstractTextDocumentLayout_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QAbstractTextDocumentLayout::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QAbstractTextDocumentLayout_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QAbstractTextDocumentLayout::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QAbstractTextDocumentLayout_unregisterHandler2(void* self, int objectType, void* component) {
	((QAbstractTextDocumentLayout*)self)->unregisterHandler(objectType, (QObject*)component);
}
void QAbstractTextDocumentLayout_update2(void* self, void* arg1) {
	((QAbstractTextDocumentLayout*)self)->update(*(QRectF*)arg1);
}
void* QAbstractTextDocumentLayout_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QAbstractTextDocumentLayout::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QAbstractTextDocumentLayout_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QAbstractTextDocumentLayout::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void QAbstractTextDocumentLayout_connect_update(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractTextDocumentLayout*)self, &QAbstractTextDocumentLayout::update, (QAbstractTextDocumentLayout*)self, [__id](const QRectF& arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractTextDocumentLayout);
	__conns_QAbstractTextDocumentLayout[__id] = __c;
}
void QAbstractTextDocumentLayout_connect_update2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractTextDocumentLayout*)self, &QAbstractTextDocumentLayout::update, (QAbstractTextDocumentLayout*)self, [__id](const QRectF& arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractTextDocumentLayout);
	__conns_QAbstractTextDocumentLayout[__id] = __c;
}
void QAbstractTextDocumentLayout_connect_updateBlock(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractTextDocumentLayout*)self, &QAbstractTextDocumentLayout::updateBlock, (QAbstractTextDocumentLayout*)self, [__id](const QTextBlock& block){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&block);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractTextDocumentLayout);
	__conns_QAbstractTextDocumentLayout[__id] = __c;
}
void QAbstractTextDocumentLayout_connect_documentSizeChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractTextDocumentLayout*)self, &QAbstractTextDocumentLayout::documentSizeChanged, (QAbstractTextDocumentLayout*)self, [__id](const QSizeF& newSize){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&newSize);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractTextDocumentLayout);
	__conns_QAbstractTextDocumentLayout[__id] = __c;
}
void QAbstractTextDocumentLayout_connect_pageCountChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractTextDocumentLayout*)self, &QAbstractTextDocumentLayout::pageCountChanged, (QAbstractTextDocumentLayout*)self, [__id](int newPages){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, newPages);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractTextDocumentLayout);
	__conns_QAbstractTextDocumentLayout[__id] = __c;
}
void QAbstractTextDocumentLayout_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractTextDocumentLayout*)self, &QAbstractTextDocumentLayout::destroyed, (QAbstractTextDocumentLayout*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractTextDocumentLayout);
	__conns_QAbstractTextDocumentLayout[__id] = __c;
}
void QAbstractTextDocumentLayout_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractTextDocumentLayout*)self, &QAbstractTextDocumentLayout::destroyed, (QAbstractTextDocumentLayout*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractTextDocumentLayout);
	__conns_QAbstractTextDocumentLayout[__id] = __c;
}
void QAbstractTextDocumentLayout_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractTextDocumentLayout*)self, &QAbstractTextDocumentLayout::objectNameChanged, (QAbstractTextDocumentLayout*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractTextDocumentLayout);
	__conns_QAbstractTextDocumentLayout[__id] = __c;
}
void QAbstractTextDocumentLayout__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QAbstractTextDocumentLayout);
		auto __it = __conns_QAbstractTextDocumentLayout.find(id);
		if (__it == __conns_QAbstractTextDocumentLayout.end()) return;
		__c = __it->second;
		__conns_QAbstractTextDocumentLayout.erase(__it);
	}
	QObject::disconnect(__c);
}
void QAbstractTextDocumentLayout_delete(void* self) { delete (QAbstractTextDocumentLayout*)self; }

}

#include <QtGui/qaccessible.h>
#include <QtCore/qobject.h>
#include <QtCore/qstring.h>
#include <QtGui/qaccessible_base.h>
#include "gen_qaccessibletextremoveevent.h"

extern "C" {

void* QAccessibleTextRemoveEvent_base(void* self) { return (void*)static_cast<QAccessibleTextCursorEvent*>((QAccessibleTextRemoveEvent*)self); }
void QAccessibleTextRemoveEvent_chain(void* self, void** out) {
	QAccessibleTextRemoveEvent* __p = (QAccessibleTextRemoveEvent*)self;
	QAccessibleTextCursorEvent* __b0 = static_cast<QAccessibleTextCursorEvent*>(__p); out[0] = (void*)__b0;
	QAccessibleEvent* __b1 = static_cast<QAccessibleEvent*>(__b0); out[1] = (void*)__b1;
}
void* QAccessibleTextRemoveEvent_new(void* obj, int position, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	return new QAccessibleTextRemoveEvent((QObject*)obj, position, text_qstr);
}
void* QAccessibleTextRemoveEvent_new2(void* iface, int position, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	return new QAccessibleTextRemoveEvent((QAccessibleInterface*)iface, position, text_qstr);
}
char* QAccessibleTextRemoveEvent_textRemoved(void* self, long long* __out_len) {
	auto&& __r = ((QAccessibleTextRemoveEvent*)self)->textRemoved();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QAccessibleTextRemoveEvent_changePosition(void* self) {
	auto&& __r = ((QAccessibleTextRemoveEvent*)self)->changePosition();
	return __r;
}
void QAccessibleTextRemoveEvent_delete(void* self) { delete (QAccessibleTextRemoveEvent*)self; }

}

#include <QtGui/qfilesystemmodel.h>
#include <QtCore/qabstractitemmodel.h>
#include <QtCore/qobject.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qvariant.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qmimedata.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qdir.h>
#include <QtGui/qabstractfileiconprovider.h>
#include <QtCore/qdatetime.h>
#include <QtCore/qtimezone.h>
#include <QtGui/qicon.h>
#include <QtCore/qfiledevice.h>
#include <QtCore/qfileinfo.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qsize.h>
#include <QtCore/qdatastream.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qfilesystemmodel.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QFileSystemModel__Go : public QFileSystemModel {
public:
	long long __goId;
	unsigned char __goOn[42] = {};
	template<typename... A> QFileSystemModel__Go(long long id, A&&... a) : QFileSystemModel(std::forward<A>(a)...), __goId(id) {}
	using QFileSystemModel::createIndex;
	using QFileSystemModel::encodeData;
	using QFileSystemModel::decodeData;
	using QFileSystemModel::beginInsertRows;
	using QFileSystemModel::endInsertRows;
	using QFileSystemModel::beginRemoveRows;
	using QFileSystemModel::endRemoveRows;
	using QFileSystemModel::beginMoveRows;
	using QFileSystemModel::endMoveRows;
	using QFileSystemModel::beginInsertColumns;
	using QFileSystemModel::endInsertColumns;
	using QFileSystemModel::beginRemoveColumns;
	using QFileSystemModel::endRemoveColumns;
	using QFileSystemModel::beginMoveColumns;
	using QFileSystemModel::endMoveColumns;
	using QFileSystemModel::beginResetModel;
	using QFileSystemModel::endResetModel;
	using QFileSystemModel::changePersistentIndex;
	using QFileSystemModel::changePersistentIndexList;
	using QFileSystemModel::persistentIndexList;
	using QFileSystemModel::sender;
	using QFileSystemModel::senderSignalIndex;
	using QFileSystemModel::receivers;
	using QFileSystemModel::isSignalConnected;
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QFileSystemModel::timerEvent(__a0); }
	bool cppgoBase_event(QEvent* __a0) { return QFileSystemModel::event(__a0); }
	void cppgoBase_resetInternalData() { QFileSystemModel::resetInternalData(); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QFileSystemModel::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QFileSystemModel::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QFileSystemModel::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QFileSystemModel::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QFileSystemModel::metaObject();
	}
	QModelIndex index(int row, int column, const QModelIndex& parent) const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_voidp_int_int_voidp) { void* __r = ((void*(*)(long long, long long, int, int, void*, int*))cppgo_virt_voidp_int_int_voidp)(__goId, -8948777187306027381LL, row, column, (void*)&parent, &__h); if (__h) return *(QModelIndex*)__r; }
		return QFileSystemModel::index(row, column, parent);
	}
	QModelIndex parent(const QModelIndex& child) const override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 7371339085820260189LL, (void*)&child, &__h); if (__h) return *(QModelIndex*)__r; }
		return QFileSystemModel::parent(child);
	}
	QModelIndex sibling(int row, int column, const QModelIndex& idx) const override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_voidp_int_int_voidp) { void* __r = ((void*(*)(long long, long long, int, int, void*, int*))cppgo_virt_voidp_int_int_voidp)(__goId, 397329346813782521LL, row, column, (void*)&idx, &__h); if (__h) return *(QModelIndex*)__r; }
		return QFileSystemModel::sibling(row, column, idx);
	}
	bool hasChildren(const QModelIndex& parent) const override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 3369452259144514380LL, (void*)&parent, &__h); if (__h) return __r; }
		return QFileSystemModel::hasChildren(parent);
	}
	bool canFetchMore(const QModelIndex& parent) const override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 8639262066473964150LL, (void*)&parent, &__h); if (__h) return __r; }
		return QFileSystemModel::canFetchMore(parent);
	}
	void fetchMore(const QModelIndex& parent) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3553294593611130806LL, (void*)&parent, &__h); if (__h) return; }
		QFileSystemModel::fetchMore(parent);
	}
	int rowCount(const QModelIndex& parent) const override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1960960171956884010LL, (void*)&parent, &__h); if (__h) return __r; }
		return QFileSystemModel::rowCount(parent);
	}
	int columnCount(const QModelIndex& parent) const override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, -8256963024669584834LL, (void*)&parent, &__h); if (__h) return __r; }
		return QFileSystemModel::columnCount(parent);
	}
	QVariant data(const QModelIndex& index, int role) const override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_voidp_voidp_int) { void* __r = ((void*(*)(long long, long long, void*, int, int*))cppgo_virt_voidp_voidp_int)(__goId, -8837375942098466299LL, (void*)&index, role, &__h); if (__h) return *(QVariant*)__r; }
		return QFileSystemModel::data(index, role);
	}
	bool setData(const QModelIndex& index, const QVariant& value, int role) override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_int_voidp_voidp_int) { int __r = ((int(*)(long long, long long, void*, void*, int, int*))cppgo_virt_int_voidp_voidp_int)(__goId, 35191454575202219LL, (void*)&index, (void*)&value, role, &__h); if (__h) return __r; }
		return QFileSystemModel::setData(index, value, role);
	}
	QVariant headerData(int section, Qt::Orientation orientation, int role) const override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_voidp_int_int_int) { void* __r = ((void*(*)(long long, long long, int, int, int, int*))cppgo_virt_voidp_int_int_int)(__goId, -3297140021896693688LL, section, (int)(orientation), role, &__h); if (__h) return *(QVariant*)__r; }
		return QFileSystemModel::headerData(section, orientation, role);
	}
	QFlags<Qt::ItemFlag> flags(const QModelIndex& index) const override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1703382833545304812LL, (void*)&index, &__h); if (__h) return QFlags<Qt::ItemFlag>(__r); }
		return QFileSystemModel::flags(index);
	}
	void sort(int column, Qt::SortOrder order) override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_void_int_int) { ((void(*)(long long, long long, int, int, int*))cppgo_virt_void_int_int)(__goId, 6912441694453382161LL, column, (int)(order), &__h); if (__h) return; }
		QFileSystemModel::sort(column, order);
	}
	QList<QString> mimeTypes() const override {
		int __h = 0;
		if (__goOn[14] && cppgo_virtp_voidp) {
			long long __len = 0;
			void* __a = ((void*(*)(long long, long long, long long*, int*))cppgo_virtp_voidp)(__goId, 5464467312394730418LL, &__len, &__h);
			if (__h) {
				QList<QString> __r; __r.reserve(__len);
				char** __aa = (char**)__a;
				for (long long __i = 0; __i < __len; __i++) { __r.append(QString::fromUtf8(__aa[__i])); free(__aa[__i]); }
				if (__a) free(__a);
				return __r;
			}
		}
		return QFileSystemModel::mimeTypes();
	}
	QMimeData* mimeData(const QList<QModelIndex>& indexes) const override {
		int __h = 0;
		std::vector<void*> indexes_v; indexes_v.reserve(indexes.size());
		for (long long __i = 0; __i < indexes.size(); __i++) indexes_v.push_back((void*)&indexes[__i]);
		if (__goOn[15] && cppgo_virt_voidp_voidpp_longlong) { void* __r = ((void*(*)(long long, long long, void**, long long, int*))cppgo_virt_voidp_voidpp_longlong)(__goId, 5175297091024974749LL, indexes_v.data(), (long long)indexes.size(), &__h); if (__h) return (QMimeData*)__r; }
		return QFileSystemModel::mimeData(indexes);
	}
	bool dropMimeData(const QMimeData* data, Qt::DropAction action, int row, int column, const QModelIndex& parent) override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_int_voidp_int_int_int_voidp) { int __r = ((int(*)(long long, long long, void*, int, int, int, void*, int*))cppgo_virt_int_voidp_int_int_int_voidp)(__goId, -9049994684305682212LL, (void*)data, (int)(action), row, column, (void*)&parent, &__h); if (__h) return __r; }
		return QFileSystemModel::dropMimeData(data, action, row, column, parent);
	}
	QFlags<Qt::DropAction> supportedDropActions() const override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -7784642176186821205LL, &__h); if (__h) return QFlags<Qt::DropAction>(__r); }
		return QFileSystemModel::supportedDropActions();
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QFileSystemModel::timerEvent(event);
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QFileSystemModel::event(event);
	}
	bool setHeaderData(int section, Qt::Orientation orientation, const QVariant& value, int role) override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_int_int_int_voidp_int) { int __r = ((int(*)(long long, long long, int, int, void*, int, int*))cppgo_virt_int_int_int_voidp_int)(__goId, 4852200306590054850LL, section, (int)(orientation), (void*)&value, role, &__h); if (__h) return __r; }
		return QFileSystemModel::setHeaderData(section, orientation, value, role);
	}
	bool clearItemData(const QModelIndex& index) override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 8914190500520542409LL, (void*)&index, &__h); if (__h) return __r; }
		return QFileSystemModel::clearItemData(index);
	}
	bool canDropMimeData(const QMimeData* data, Qt::DropAction action, int row, int column, const QModelIndex& parent) const override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_int_voidp_int_int_int_voidp) { int __r = ((int(*)(long long, long long, void*, int, int, int, void*, int*))cppgo_virt_int_voidp_int_int_int_voidp)(__goId, 731539533644753052LL, (void*)data, (int)(action), row, column, (void*)&parent, &__h); if (__h) return __r; }
		return QFileSystemModel::canDropMimeData(data, action, row, column, parent);
	}
	QFlags<Qt::DropAction> supportedDragActions() const override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, 3421601386088570762LL, &__h); if (__h) return QFlags<Qt::DropAction>(__r); }
		return QFileSystemModel::supportedDragActions();
	}
	bool insertRows(int row, int count, const QModelIndex& parent) override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_int_int_int_voidp) { int __r = ((int(*)(long long, long long, int, int, void*, int*))cppgo_virt_int_int_int_voidp)(__goId, -2775729825962395603LL, row, count, (void*)&parent, &__h); if (__h) return __r; }
		return QFileSystemModel::insertRows(row, count, parent);
	}
	bool insertColumns(int column, int count, const QModelIndex& parent) override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_int_int_int_voidp) { int __r = ((int(*)(long long, long long, int, int, void*, int*))cppgo_virt_int_int_int_voidp)(__goId, -8973293846230792237LL, column, count, (void*)&parent, &__h); if (__h) return __r; }
		return QFileSystemModel::insertColumns(column, count, parent);
	}
	bool removeRows(int row, int count, const QModelIndex& parent) override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_int_int_int_voidp) { int __r = ((int(*)(long long, long long, int, int, void*, int*))cppgo_virt_int_int_int_voidp)(__goId, 5104072549182222656LL, row, count, (void*)&parent, &__h); if (__h) return __r; }
		return QFileSystemModel::removeRows(row, count, parent);
	}
	bool removeColumns(int column, int count, const QModelIndex& parent) override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_int_int_int_voidp) { int __r = ((int(*)(long long, long long, int, int, void*, int*))cppgo_virt_int_int_int_voidp)(__goId, 2078548737056540420LL, column, count, (void*)&parent, &__h); if (__h) return __r; }
		return QFileSystemModel::removeColumns(column, count, parent);
	}
	bool moveRows(const QModelIndex& sourceParent, int sourceRow, int count, const QModelIndex& destinationParent, int destinationChild) override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_int_voidp_int_int_voidp_int) { int __r = ((int(*)(long long, long long, void*, int, int, void*, int, int*))cppgo_virt_int_voidp_int_int_voidp_int)(__goId, -6074819356947405415LL, (void*)&sourceParent, sourceRow, count, (void*)&destinationParent, destinationChild, &__h); if (__h) return __r; }
		return QFileSystemModel::moveRows(sourceParent, sourceRow, count, destinationParent, destinationChild);
	}
	bool moveColumns(const QModelIndex& sourceParent, int sourceColumn, int count, const QModelIndex& destinationParent, int destinationChild) override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_int_voidp_int_int_voidp_int) { int __r = ((int(*)(long long, long long, void*, int, int, void*, int, int*))cppgo_virt_int_voidp_int_int_voidp_int)(__goId, -6321007591388309905LL, (void*)&sourceParent, sourceColumn, count, (void*)&destinationParent, destinationChild, &__h); if (__h) return __r; }
		return QFileSystemModel::moveColumns(sourceParent, sourceColumn, count, destinationParent, destinationChild);
	}
	QModelIndex buddy(const QModelIndex& index) const override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 6112860112248786499LL, (void*)&index, &__h); if (__h) return *(QModelIndex*)__r; }
		return QFileSystemModel::buddy(index);
	}
	QList<QModelIndex> match(const QModelIndex& start, int role, const QVariant& value, int hits, QFlags<Qt::MatchFlag> flags) const override {
		int __h = 0;
		if (__goOn[31] && cppgo_virtp_voidp_voidp_int_voidp_int_int) {
			long long __len = 0;
			void* __a = ((void*(*)(long long, long long, void*, int, void*, int, int, long long*, int*))cppgo_virtp_voidp_voidp_int_voidp_int_int)(__goId, -4341501217981328394LL, (void*)&start, role, (void*)&value, hits, (int)(flags), &__len, &__h);
			if (__h) {
				QList<QModelIndex> __r; __r.reserve(__len);
				void** __aa = (void**)__a;
				for (long long __i = 0; __i < __len; __i++) __r.append(*(QModelIndex*)__aa[__i]);
				if (__a) free(__a);
				return __r;
			}
		}
		return QFileSystemModel::match(start, role, value, hits, flags);
	}
	QSize span(const QModelIndex& index) const override {
		int __h = 0;
		if (__goOn[32] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, -8395342916587040543LL, (void*)&index, &__h); if (__h) return *(QSize*)__r; }
		return QFileSystemModel::span(index);
	}
	void multiData(const QModelIndex& index, QModelRoleDataSpan roleDataSpan) const override {
		int __h = 0;
		if (__goOn[33] && cppgo_virt_void_voidp_voidp) { ((void(*)(long long, long long, void*, void*, int*))cppgo_virt_void_voidp_voidp)(__goId, 731311060944096642LL, (void*)&index, (void*)&roleDataSpan, &__h); if (__h) return; }
		QFileSystemModel::multiData(index, roleDataSpan);
	}
	bool submit() override {
		int __h = 0;
		if (__goOn[34] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -3788331537442685179LL, &__h); if (__h) return __r; }
		return QFileSystemModel::submit();
	}
	void revert() override {
		int __h = 0;
		if (__goOn[35] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, -1522951607536117149LL, &__h); if (__h) return; }
		QFileSystemModel::revert();
	}
	void resetInternalData() override {
		int __h = 0;
		if (__goOn[36] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, -4510675770784051183LL, &__h); if (__h) return; }
		QFileSystemModel::resetInternalData();
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[37] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QFileSystemModel::eventFilter(watched, event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[38] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QFileSystemModel::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[39] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QFileSystemModel::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[40] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QFileSystemModel::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[41] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QFileSystemModel::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QFileSystemModel;
static std::mutex __connsMu_QFileSystemModel;

extern "C" {

void* QFileSystemModel_base(void* self) { return (void*)static_cast<QAbstractItemModel*>((QFileSystemModel*)self); }
void QFileSystemModel_chain(void* self, void** out) {
	QFileSystemModel* __p = (QFileSystemModel*)self;
	QAbstractItemModel* __b0 = static_cast<QAbstractItemModel*>(__p); out[0] = (void*)__b0;
	QObject* __b1 = static_cast<QObject*>(__b0); out[1] = (void*)__b1;
}
void QFileSystemModel__set_override(void* self, int idx) { ((QFileSystemModel__Go*)self)->__goOn[idx] = 1; }
void* QFileSystemModel_fromQObject(void* o) { return (void*)qobject_cast<QFileSystemModel*>((QObject*)o); }
void* QFileSystemModel_new(long long __goid) {
	return new QFileSystemModel__Go(__goid);
}
void* QFileSystemModel_new2(long long __goid, void* parent) {
	return new QFileSystemModel__Go(__goid, (QObject*)parent);
}
char* QFileSystemModel_tr(const char* s, long long* __out_len) {
	auto&& __r = QFileSystemModel::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QFileSystemModel_rootPathChanged(void* self, const char* newPath, long long newPath_len) {
	QString newPath_qstr = QString::fromUtf8(newPath, newPath_len);
	((QFileSystemModel*)self)->rootPathChanged(newPath_qstr);
}
void QFileSystemModel_fileRenamed(void* self, const char* path, long long path_len, const char* oldName, long long oldName_len, const char* newName, long long newName_len) {
	QString path_qstr = QString::fromUtf8(path, path_len);
	QString oldName_qstr = QString::fromUtf8(oldName, oldName_len);
	QString newName_qstr = QString::fromUtf8(newName, newName_len);
	((QFileSystemModel*)self)->fileRenamed(path_qstr, oldName_qstr, newName_qstr);
}
void QFileSystemModel_directoryLoaded(void* self, const char* path, long long path_len) {
	QString path_qstr = QString::fromUtf8(path, path_len);
	((QFileSystemModel*)self)->directoryLoaded(path_qstr);
}
void* QFileSystemModel_index(void* self, int row, int column, void* parent) {
	auto&& __r = ((QFileSystemModel*)self)->index(row, column, *(QModelIndex*)parent);
	return new QModelIndex(__r);
}
void* QFileSystemModel_index2(void* self, const char* path, long long path_len) {
	QString path_qstr = QString::fromUtf8(path, path_len);
	auto&& __r = ((QFileSystemModel*)self)->index(path_qstr);
	return new QModelIndex(__r);
}
void* QFileSystemModel_myComputer(void* self) {
	auto&& __r = ((QFileSystemModel*)self)->myComputer();
	return new QVariant(__r);
}
void* QFileSystemModel_setRootPath(void* self, const char* path, long long path_len) {
	QString path_qstr = QString::fromUtf8(path, path_len);
	auto&& __r = ((QFileSystemModel*)self)->setRootPath(path_qstr);
	return new QModelIndex(__r);
}
char* QFileSystemModel_rootPath(void* self, long long* __out_len) {
	auto&& __r = ((QFileSystemModel*)self)->rootPath();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QFileSystemModel_rootDirectory(void* self) {
	auto&& __r = ((QFileSystemModel*)self)->rootDirectory();
	return new QDir(__r);
}
void QFileSystemModel_setIconProvider(void* self, void* provider) {
	((QFileSystemModel*)self)->setIconProvider((QAbstractFileIconProvider*)provider);
}
void* QFileSystemModel_iconProvider(void* self) {
	auto&& __r = ((QFileSystemModel*)self)->iconProvider();
	return (void*)(__r);
}
void QFileSystemModel_setFilter(void* self, int filters) {
	((QFileSystemModel*)self)->setFilter(QFlags<QDir::Filter>(filters));
}
int QFileSystemModel_filter(void* self) {
	auto&& __r = ((QFileSystemModel*)self)->filter();
	return (int)(__r);
}
void QFileSystemModel_setResolveSymlinks(void* self, int enable) {
	((QFileSystemModel*)self)->setResolveSymlinks(enable);
}
int QFileSystemModel_resolveSymlinks(void* self) {
	auto&& __r = ((QFileSystemModel*)self)->resolveSymlinks();
	return __r;
}
void QFileSystemModel_setReadOnly(void* self, int enable) {
	((QFileSystemModel*)self)->setReadOnly(enable);
}
int QFileSystemModel_isReadOnly(void* self) {
	auto&& __r = ((QFileSystemModel*)self)->isReadOnly();
	return __r;
}
void QFileSystemModel_setNameFilterDisables(void* self, int enable) {
	((QFileSystemModel*)self)->setNameFilterDisables(enable);
}
int QFileSystemModel_nameFilterDisables(void* self) {
	auto&& __r = ((QFileSystemModel*)self)->nameFilterDisables();
	return __r;
}
void QFileSystemModel_setNameFilters(void* self, char** filters, long long filters_len) {
	QList<QString> filters_list; filters_list.reserve(filters_len);
	for (long long __i = 0; __i < filters_len; __i++) filters_list.append(QString::fromUtf8(filters[__i]));
	((QFileSystemModel*)self)->setNameFilters(filters_list);
}
char** QFileSystemModel_nameFilters(void* self, long long* __out_len) {
	auto&& __r = ((QFileSystemModel*)self)->nameFilters();
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
void QFileSystemModel_setOption(void* self, int option) {
	((QFileSystemModel*)self)->setOption((QFileSystemModel::Option)(option));
}
int QFileSystemModel_testOption(void* self, int option) {
	auto&& __r = ((QFileSystemModel*)self)->testOption((QFileSystemModel::Option)(option));
	return __r;
}
void QFileSystemModel_setOptions(void* self, int options) {
	((QFileSystemModel*)self)->setOptions(QFlags<QFileSystemModel::Option>(options));
}
int QFileSystemModel_options(void* self) {
	auto&& __r = ((QFileSystemModel*)self)->options();
	return (int)(__r);
}
char* QFileSystemModel_filePath(void* self, void* index, long long* __out_len) {
	auto&& __r = ((QFileSystemModel*)self)->filePath(*(QModelIndex*)index);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QFileSystemModel_isDir(void* self, void* index) {
	auto&& __r = ((QFileSystemModel*)self)->isDir(*(QModelIndex*)index);
	return __r;
}
long long QFileSystemModel_size(void* self, void* index) {
	auto&& __r = ((QFileSystemModel*)self)->size(*(QModelIndex*)index);
	return __r;
}
char* QFileSystemModel_type(void* self, void* index, long long* __out_len) {
	auto&& __r = ((QFileSystemModel*)self)->type(*(QModelIndex*)index);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QFileSystemModel_lastModified(void* self, void* index) {
	auto&& __r = ((QFileSystemModel*)self)->lastModified(*(QModelIndex*)index);
	return new QDateTime(__r);
}
void* QFileSystemModel_lastModified2(void* self, void* index, void* tz) {
	auto&& __r = ((QFileSystemModel*)self)->lastModified(*(QModelIndex*)index, *(QTimeZone*)tz);
	return new QDateTime(__r);
}
void* QFileSystemModel_mkdir(void* self, void* parent, const char* name, long long name_len) {
	QString name_qstr = QString::fromUtf8(name, name_len);
	auto&& __r = ((QFileSystemModel*)self)->mkdir(*(QModelIndex*)parent, name_qstr);
	return new QModelIndex(__r);
}
int QFileSystemModel_rmdir(void* self, void* index) {
	auto&& __r = ((QFileSystemModel*)self)->rmdir(*(QModelIndex*)index);
	return __r;
}
char* QFileSystemModel_fileName(void* self, void* index, long long* __out_len) {
	auto&& __r = ((QFileSystemModel*)self)->fileName(*(QModelIndex*)index);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QFileSystemModel_fileIcon(void* self, void* index) {
	auto&& __r = ((QFileSystemModel*)self)->fileIcon(*(QModelIndex*)index);
	return new QIcon(__r);
}
int QFileSystemModel_permissions(void* self, void* index) {
	auto&& __r = ((QFileSystemModel*)self)->permissions(*(QModelIndex*)index);
	return (int)(__r);
}
void* QFileSystemModel_fileInfo(void* self, void* index) {
	auto&& __r = ((QFileSystemModel*)self)->fileInfo(*(QModelIndex*)index);
	return new QFileInfo(__r);
}
int QFileSystemModel_remove(void* self, void* index) {
	auto&& __r = ((QFileSystemModel*)self)->remove(*(QModelIndex*)index);
	return __r;
}
void* QFileSystemModel_createIndex(void* self, int row, int column) {
	auto&& __r = ((QFileSystemModel__Go*)self)->createIndex(row, column);
	return new QModelIndex(__r);
}
void* QFileSystemModel_createIndex2(void* self, int row, int column, unsigned long long id) {
	auto&& __r = ((QFileSystemModel__Go*)self)->createIndex(row, column, id);
	return new QModelIndex(__r);
}
void QFileSystemModel_encodeData(void* self, void** indexes, long long indexes_len, void* stream) {
	QList<QModelIndex> indexes_list; indexes_list.reserve(indexes_len);
	for (long long __i = 0; __i < indexes_len; __i++) indexes_list.append(*(QModelIndex*)indexes[__i]);
	((QFileSystemModel__Go*)self)->encodeData(indexes_list, *(QDataStream*)stream);
}
int QFileSystemModel_decodeData(void* self, int row, int column, void* parent, void* stream) {
	auto&& __r = ((QFileSystemModel__Go*)self)->decodeData(row, column, *(QModelIndex*)parent, *(QDataStream*)stream);
	return __r;
}
void QFileSystemModel_beginInsertRows(void* self, void* parent, int first, int last) {
	((QFileSystemModel__Go*)self)->beginInsertRows(*(QModelIndex*)parent, first, last);
}
void QFileSystemModel_endInsertRows(void* self) {
	((QFileSystemModel__Go*)self)->endInsertRows();
}
void QFileSystemModel_beginRemoveRows(void* self, void* parent, int first, int last) {
	((QFileSystemModel__Go*)self)->beginRemoveRows(*(QModelIndex*)parent, first, last);
}
void QFileSystemModel_endRemoveRows(void* self) {
	((QFileSystemModel__Go*)self)->endRemoveRows();
}
int QFileSystemModel_beginMoveRows(void* self, void* sourceParent, int sourceFirst, int sourceLast, void* destinationParent, int destinationRow) {
	auto&& __r = ((QFileSystemModel__Go*)self)->beginMoveRows(*(QModelIndex*)sourceParent, sourceFirst, sourceLast, *(QModelIndex*)destinationParent, destinationRow);
	return __r;
}
void QFileSystemModel_endMoveRows(void* self) {
	((QFileSystemModel__Go*)self)->endMoveRows();
}
void QFileSystemModel_beginInsertColumns(void* self, void* parent, int first, int last) {
	((QFileSystemModel__Go*)self)->beginInsertColumns(*(QModelIndex*)parent, first, last);
}
void QFileSystemModel_endInsertColumns(void* self) {
	((QFileSystemModel__Go*)self)->endInsertColumns();
}
void QFileSystemModel_beginRemoveColumns(void* self, void* parent, int first, int last) {
	((QFileSystemModel__Go*)self)->beginRemoveColumns(*(QModelIndex*)parent, first, last);
}
void QFileSystemModel_endRemoveColumns(void* self) {
	((QFileSystemModel__Go*)self)->endRemoveColumns();
}
int QFileSystemModel_beginMoveColumns(void* self, void* sourceParent, int sourceFirst, int sourceLast, void* destinationParent, int destinationColumn) {
	auto&& __r = ((QFileSystemModel__Go*)self)->beginMoveColumns(*(QModelIndex*)sourceParent, sourceFirst, sourceLast, *(QModelIndex*)destinationParent, destinationColumn);
	return __r;
}
void QFileSystemModel_endMoveColumns(void* self) {
	((QFileSystemModel__Go*)self)->endMoveColumns();
}
void QFileSystemModel_beginResetModel(void* self) {
	((QFileSystemModel__Go*)self)->beginResetModel();
}
void QFileSystemModel_endResetModel(void* self) {
	((QFileSystemModel__Go*)self)->endResetModel();
}
void QFileSystemModel_changePersistentIndex(void* self, void* from, void* to) {
	((QFileSystemModel__Go*)self)->changePersistentIndex(*(QModelIndex*)from, *(QModelIndex*)to);
}
void QFileSystemModel_changePersistentIndexList(void* self, void** from, long long from_len, void** to, long long to_len) {
	QList<QModelIndex> from_list; from_list.reserve(from_len);
	for (long long __i = 0; __i < from_len; __i++) from_list.append(*(QModelIndex*)from[__i]);
	QList<QModelIndex> to_list; to_list.reserve(to_len);
	for (long long __i = 0; __i < to_len; __i++) to_list.append(*(QModelIndex*)to[__i]);
	((QFileSystemModel__Go*)self)->changePersistentIndexList(from_list, to_list);
}
void** QFileSystemModel_persistentIndexList(void* self, long long* __out_len) {
	auto&& __r = ((QFileSystemModel__Go*)self)->persistentIndexList();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QModelIndex(__e);
	return __arr;
}
void* QFileSystemModel_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QFileSystemModel::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QFileSystemModel_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QFileSystemModel::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QFileSystemModel_sender(void* self) {
	auto&& __r = ((QFileSystemModel__Go*)self)->sender();
	return (void*)(__r);
}
int QFileSystemModel_senderSignalIndex(void* self) {
	auto&& __r = ((QFileSystemModel__Go*)self)->senderSignalIndex();
	return __r;
}
int QFileSystemModel_receivers(void* self, const char* signal) {
	auto&& __r = ((QFileSystemModel__Go*)self)->receivers(signal);
	return __r;
}
int QFileSystemModel_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QFileSystemModel__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QFileSystemModel_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QFileSystemModel::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QFileSystemModel_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QFileSystemModel::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QFileSystemModel_index3(void* self, const char* path, long long path_len, int column) {
	QString path_qstr = QString::fromUtf8(path, path_len);
	auto&& __r = ((QFileSystemModel*)self)->index(path_qstr, column);
	return new QModelIndex(__r);
}
void* QFileSystemModel_myComputer2(void* self, int role) {
	auto&& __r = ((QFileSystemModel*)self)->myComputer(role);
	return new QVariant(__r);
}
void QFileSystemModel_setOption2(void* self, int option, int on) {
	((QFileSystemModel*)self)->setOption((QFileSystemModel::Option)(option), on);
}
void* QFileSystemModel_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QFileSystemModel::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QFileSystemModel_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QFileSystemModel::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QFileSystemModel_index_default(void* self, int row, int column, void* parent) {
	auto&& __r = ((QFileSystemModel*)self)->QFileSystemModel::index(row, column, *(QModelIndex*)parent);
	return new QModelIndex(__r);
}
void* QFileSystemModel_index_default2(void* self, const char* path, long long path_len) {
	QString path_qstr = QString::fromUtf8(path, path_len);
	auto&& __r = ((QFileSystemModel*)self)->QFileSystemModel::index(path_qstr);
	return new QModelIndex(__r);
}
void* QFileSystemModel_index_default3(void* self, const char* path, long long path_len, int column) {
	QString path_qstr = QString::fromUtf8(path, path_len);
	auto&& __r = ((QFileSystemModel*)self)->QFileSystemModel::index(path_qstr, column);
	return new QModelIndex(__r);
}
void* QFileSystemModel_metaObject_default(void* self) {
	auto&& __r = ((QFileSystemModel*)self)->QFileSystemModel::metaObject();
	return (void*)(__r);
}
void* QFileSystemModel_parent_default(void* self, void* child) {
	auto&& __r = ((QFileSystemModel*)self)->QFileSystemModel::parent(*(QModelIndex*)child);
	return new QModelIndex(__r);
}
void* QFileSystemModel_sibling_default(void* self, int row, int column, void* idx) {
	auto&& __r = ((QFileSystemModel*)self)->QFileSystemModel::sibling(row, column, *(QModelIndex*)idx);
	return new QModelIndex(__r);
}
int QFileSystemModel_hasChildren_default(void* self, void* parent) {
	auto&& __r = ((QFileSystemModel*)self)->QFileSystemModel::hasChildren(*(QModelIndex*)parent);
	return __r;
}
int QFileSystemModel_canFetchMore_default(void* self, void* parent) {
	auto&& __r = ((QFileSystemModel*)self)->QFileSystemModel::canFetchMore(*(QModelIndex*)parent);
	return __r;
}
void QFileSystemModel_fetchMore_default(void* self, void* parent) {
	((QFileSystemModel*)self)->QFileSystemModel::fetchMore(*(QModelIndex*)parent);
}
int QFileSystemModel_rowCount_default(void* self, void* parent) {
	auto&& __r = ((QFileSystemModel*)self)->QFileSystemModel::rowCount(*(QModelIndex*)parent);
	return __r;
}
int QFileSystemModel_columnCount_default(void* self, void* parent) {
	auto&& __r = ((QFileSystemModel*)self)->QFileSystemModel::columnCount(*(QModelIndex*)parent);
	return __r;
}
void* QFileSystemModel_data_default(void* self, void* index, int role) {
	auto&& __r = ((QFileSystemModel*)self)->QFileSystemModel::data(*(QModelIndex*)index, role);
	return new QVariant(__r);
}
int QFileSystemModel_setData_default(void* self, void* index, void* value, int role) {
	auto&& __r = ((QFileSystemModel*)self)->QFileSystemModel::setData(*(QModelIndex*)index, *(QVariant*)value, role);
	return __r;
}
void* QFileSystemModel_headerData_default(void* self, int section, int orientation, int role) {
	auto&& __r = ((QFileSystemModel*)self)->QFileSystemModel::headerData(section, (Qt::Orientation)(orientation), role);
	return new QVariant(__r);
}
int QFileSystemModel_flags_default(void* self, void* index) {
	auto&& __r = ((QFileSystemModel*)self)->QFileSystemModel::flags(*(QModelIndex*)index);
	return (int)(__r);
}
void QFileSystemModel_sort_default(void* self, int column, int order) {
	((QFileSystemModel*)self)->QFileSystemModel::sort(column, (Qt::SortOrder)(order));
}
char** QFileSystemModel_mimeTypes_default(void* self, long long* __out_len) {
	auto&& __r = ((QFileSystemModel*)self)->QFileSystemModel::mimeTypes();
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
void* QFileSystemModel_mimeData_default(void* self, void** indexes, long long indexes_len) {
	QList<QModelIndex> indexes_list; indexes_list.reserve(indexes_len);
	for (long long __i = 0; __i < indexes_len; __i++) indexes_list.append(*(QModelIndex*)indexes[__i]);
	auto&& __r = ((QFileSystemModel*)self)->QFileSystemModel::mimeData(indexes_list);
	return (void*)(__r);
}
int QFileSystemModel_dropMimeData_default(void* self, void* data, int action, int row, int column, void* parent) {
	auto&& __r = ((QFileSystemModel*)self)->QFileSystemModel::dropMimeData((QMimeData*)data, (Qt::DropAction)(action), row, column, *(QModelIndex*)parent);
	return __r;
}
int QFileSystemModel_supportedDropActions_default(void* self) {
	auto&& __r = ((QFileSystemModel*)self)->QFileSystemModel::supportedDropActions();
	return (int)(__r);
}
void QFileSystemModel_timerEvent_default(void* self, void* event) {
	((QFileSystemModel__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
int QFileSystemModel_event_default(void* self, void* event) {
	auto&& __r = ((QFileSystemModel__Go*)self)->cppgoBase_event((QEvent*)event);
	return __r;
}
int QFileSystemModel_setHeaderData_default(void* self, int section, int orientation, void* value, int role) {
	auto&& __r = ((QFileSystemModel*)self)->QFileSystemModel::setHeaderData(section, (Qt::Orientation)(orientation), *(QVariant*)value, role);
	return __r;
}
int QFileSystemModel_clearItemData_default(void* self, void* index) {
	auto&& __r = ((QFileSystemModel*)self)->QFileSystemModel::clearItemData(*(QModelIndex*)index);
	return __r;
}
int QFileSystemModel_canDropMimeData_default(void* self, void* data, int action, int row, int column, void* parent) {
	auto&& __r = ((QFileSystemModel*)self)->QFileSystemModel::canDropMimeData((QMimeData*)data, (Qt::DropAction)(action), row, column, *(QModelIndex*)parent);
	return __r;
}
int QFileSystemModel_supportedDragActions_default(void* self) {
	auto&& __r = ((QFileSystemModel*)self)->QFileSystemModel::supportedDragActions();
	return (int)(__r);
}
int QFileSystemModel_insertRows_default(void* self, int row, int count, void* parent) {
	auto&& __r = ((QFileSystemModel*)self)->QFileSystemModel::insertRows(row, count, *(QModelIndex*)parent);
	return __r;
}
int QFileSystemModel_insertColumns_default(void* self, int column, int count, void* parent) {
	auto&& __r = ((QFileSystemModel*)self)->QFileSystemModel::insertColumns(column, count, *(QModelIndex*)parent);
	return __r;
}
int QFileSystemModel_removeRows_default(void* self, int row, int count, void* parent) {
	auto&& __r = ((QFileSystemModel*)self)->QFileSystemModel::removeRows(row, count, *(QModelIndex*)parent);
	return __r;
}
int QFileSystemModel_removeColumns_default(void* self, int column, int count, void* parent) {
	auto&& __r = ((QFileSystemModel*)self)->QFileSystemModel::removeColumns(column, count, *(QModelIndex*)parent);
	return __r;
}
int QFileSystemModel_moveRows_default(void* self, void* sourceParent, int sourceRow, int count, void* destinationParent, int destinationChild) {
	auto&& __r = ((QFileSystemModel*)self)->QFileSystemModel::moveRows(*(QModelIndex*)sourceParent, sourceRow, count, *(QModelIndex*)destinationParent, destinationChild);
	return __r;
}
int QFileSystemModel_moveColumns_default(void* self, void* sourceParent, int sourceColumn, int count, void* destinationParent, int destinationChild) {
	auto&& __r = ((QFileSystemModel*)self)->QFileSystemModel::moveColumns(*(QModelIndex*)sourceParent, sourceColumn, count, *(QModelIndex*)destinationParent, destinationChild);
	return __r;
}
void* QFileSystemModel_buddy_default(void* self, void* index) {
	auto&& __r = ((QFileSystemModel*)self)->QFileSystemModel::buddy(*(QModelIndex*)index);
	return new QModelIndex(__r);
}
void** QFileSystemModel_match_default(void* self, void* start, int role, void* value, int hits, int flags, long long* __out_len) {
	auto&& __r = ((QFileSystemModel*)self)->QFileSystemModel::match(*(QModelIndex*)start, role, *(QVariant*)value, hits, QFlags<Qt::MatchFlag>(flags));
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QModelIndex(__e);
	return __arr;
}
void* QFileSystemModel_span_default(void* self, void* index) {
	auto&& __r = ((QFileSystemModel*)self)->QFileSystemModel::span(*(QModelIndex*)index);
	return new QSize(__r);
}
void QFileSystemModel_multiData_default(void* self, void* index, void* roleDataSpan) {
	((QFileSystemModel*)self)->QFileSystemModel::multiData(*(QModelIndex*)index, *(QModelRoleDataSpan*)roleDataSpan);
}
int QFileSystemModel_submit_default(void* self) {
	auto&& __r = ((QFileSystemModel*)self)->QFileSystemModel::submit();
	return __r;
}
void QFileSystemModel_revert_default(void* self) {
	((QFileSystemModel*)self)->QFileSystemModel::revert();
}
void QFileSystemModel_resetInternalData_default(void* self) {
	((QFileSystemModel__Go*)self)->cppgoBase_resetInternalData();
}
int QFileSystemModel_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QFileSystemModel*)self)->QFileSystemModel::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QFileSystemModel_childEvent_default(void* self, void* event) {
	((QFileSystemModel__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QFileSystemModel_customEvent_default(void* self, void* event) {
	((QFileSystemModel__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QFileSystemModel_connectNotify_default(void* self, void* signal) {
	((QFileSystemModel__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QFileSystemModel_disconnectNotify_default(void* self, void* signal) {
	((QFileSystemModel__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QFileSystemModel_connect_rootPathChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFileSystemModel*)self, &QFileSystemModel::rootPathChanged, (QFileSystemModel*)self, [__id](const QString& newPath){
		QByteArray newPath_ba = newPath.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)newPath_ba.constData(), newPath_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFileSystemModel);
	__conns_QFileSystemModel[__id] = __c;
}
void QFileSystemModel_connect_fileRenamed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFileSystemModel*)self, &QFileSystemModel::fileRenamed, (QFileSystemModel*)self, [__id](const QString& path, const QString& oldName, const QString& newName){
		QByteArray path_ba = path.toUtf8();
		QByteArray oldName_ba = oldName.toUtf8();
		QByteArray newName_ba = newName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong_constcharp_longlong_constcharp_longlong) ((void(*)(long long, const char*, long long, const char*, long long, const char*, long long))cppgo_sig_void_constcharp_longlong_constcharp_longlong_constcharp_longlong)(__id, (char*)path_ba.constData(), path_ba.size(), (char*)oldName_ba.constData(), oldName_ba.size(), (char*)newName_ba.constData(), newName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFileSystemModel);
	__conns_QFileSystemModel[__id] = __c;
}
void QFileSystemModel_connect_directoryLoaded(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFileSystemModel*)self, &QFileSystemModel::directoryLoaded, (QFileSystemModel*)self, [__id](const QString& path){
		QByteArray path_ba = path.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)path_ba.constData(), path_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFileSystemModel);
	__conns_QFileSystemModel[__id] = __c;
}
void QFileSystemModel_connect_headerDataChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFileSystemModel*)self, &QFileSystemModel::headerDataChanged, (QFileSystemModel*)self, [__id](Qt::Orientation orientation, int first, int last){
		if (cppgo_sig_void_int_int_int) ((void(*)(long long, int, int, int))cppgo_sig_void_int_int_int)(__id, (int)(orientation), first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFileSystemModel);
	__conns_QFileSystemModel[__id] = __c;
}
void QFileSystemModel_connect_layoutChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFileSystemModel*)self, &QFileSystemModel::layoutChanged, (QFileSystemModel*)self, [__id](const QList<QPersistentModelIndex>& parents, QAbstractItemModel::LayoutChangeHint hint){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFileSystemModel);
	__conns_QFileSystemModel[__id] = __c;
}
void QFileSystemModel_connect_layoutChanged2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFileSystemModel*)self, &QFileSystemModel::layoutChanged, (QFileSystemModel*)self, [__id](const QList<QPersistentModelIndex>& parents, QAbstractItemModel::LayoutChangeHint hint){
		std::vector<void*> parents_v; parents_v.reserve(parents.size());
		for (long long __i = 0; __i < parents.size(); __i++) parents_v.push_back((void*)&parents[__i]);
		if (cppgo_sig_void_voidpp_longlong_int) ((void(*)(long long, void**, long long, int))cppgo_sig_void_voidpp_longlong_int)(__id, parents_v.data(), (long long)parents.size(), (int)(hint));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFileSystemModel);
	__conns_QFileSystemModel[__id] = __c;
}
void QFileSystemModel_connect_layoutAboutToBeChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFileSystemModel*)self, &QFileSystemModel::layoutAboutToBeChanged, (QFileSystemModel*)self, [__id](const QList<QPersistentModelIndex>& parents, QAbstractItemModel::LayoutChangeHint hint){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFileSystemModel);
	__conns_QFileSystemModel[__id] = __c;
}
void QFileSystemModel_connect_layoutAboutToBeChanged2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFileSystemModel*)self, &QFileSystemModel::layoutAboutToBeChanged, (QFileSystemModel*)self, [__id](const QList<QPersistentModelIndex>& parents, QAbstractItemModel::LayoutChangeHint hint){
		std::vector<void*> parents_v; parents_v.reserve(parents.size());
		for (long long __i = 0; __i < parents.size(); __i++) parents_v.push_back((void*)&parents[__i]);
		if (cppgo_sig_void_voidpp_longlong_int) ((void(*)(long long, void**, long long, int))cppgo_sig_void_voidpp_longlong_int)(__id, parents_v.data(), (long long)parents.size(), (int)(hint));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFileSystemModel);
	__conns_QFileSystemModel[__id] = __c;
}
void QFileSystemModel_connect_rowsAboutToBeInserted(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFileSystemModel*)self, &QFileSystemModel::rowsAboutToBeInserted, (QFileSystemModel*)self, [__id](const QModelIndex& parent, int first, int last){
		if (cppgo_sig_void_voidp_int_int) ((void(*)(long long, void*, int, int))cppgo_sig_void_voidp_int_int)(__id, (void*)&parent, first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFileSystemModel);
	__conns_QFileSystemModel[__id] = __c;
}
void QFileSystemModel_connect_rowsInserted(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFileSystemModel*)self, &QFileSystemModel::rowsInserted, (QFileSystemModel*)self, [__id](const QModelIndex& parent, int first, int last){
		if (cppgo_sig_void_voidp_int_int) ((void(*)(long long, void*, int, int))cppgo_sig_void_voidp_int_int)(__id, (void*)&parent, first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFileSystemModel);
	__conns_QFileSystemModel[__id] = __c;
}
void QFileSystemModel_connect_rowsAboutToBeRemoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFileSystemModel*)self, &QFileSystemModel::rowsAboutToBeRemoved, (QFileSystemModel*)self, [__id](const QModelIndex& parent, int first, int last){
		if (cppgo_sig_void_voidp_int_int) ((void(*)(long long, void*, int, int))cppgo_sig_void_voidp_int_int)(__id, (void*)&parent, first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFileSystemModel);
	__conns_QFileSystemModel[__id] = __c;
}
void QFileSystemModel_connect_rowsRemoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFileSystemModel*)self, &QFileSystemModel::rowsRemoved, (QFileSystemModel*)self, [__id](const QModelIndex& parent, int first, int last){
		if (cppgo_sig_void_voidp_int_int) ((void(*)(long long, void*, int, int))cppgo_sig_void_voidp_int_int)(__id, (void*)&parent, first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFileSystemModel);
	__conns_QFileSystemModel[__id] = __c;
}
void QFileSystemModel_connect_columnsAboutToBeInserted(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFileSystemModel*)self, &QFileSystemModel::columnsAboutToBeInserted, (QFileSystemModel*)self, [__id](const QModelIndex& parent, int first, int last){
		if (cppgo_sig_void_voidp_int_int) ((void(*)(long long, void*, int, int))cppgo_sig_void_voidp_int_int)(__id, (void*)&parent, first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFileSystemModel);
	__conns_QFileSystemModel[__id] = __c;
}
void QFileSystemModel_connect_columnsInserted(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFileSystemModel*)self, &QFileSystemModel::columnsInserted, (QFileSystemModel*)self, [__id](const QModelIndex& parent, int first, int last){
		if (cppgo_sig_void_voidp_int_int) ((void(*)(long long, void*, int, int))cppgo_sig_void_voidp_int_int)(__id, (void*)&parent, first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFileSystemModel);
	__conns_QFileSystemModel[__id] = __c;
}
void QFileSystemModel_connect_columnsAboutToBeRemoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFileSystemModel*)self, &QFileSystemModel::columnsAboutToBeRemoved, (QFileSystemModel*)self, [__id](const QModelIndex& parent, int first, int last){
		if (cppgo_sig_void_voidp_int_int) ((void(*)(long long, void*, int, int))cppgo_sig_void_voidp_int_int)(__id, (void*)&parent, first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFileSystemModel);
	__conns_QFileSystemModel[__id] = __c;
}
void QFileSystemModel_connect_columnsRemoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFileSystemModel*)self, &QFileSystemModel::columnsRemoved, (QFileSystemModel*)self, [__id](const QModelIndex& parent, int first, int last){
		if (cppgo_sig_void_voidp_int_int) ((void(*)(long long, void*, int, int))cppgo_sig_void_voidp_int_int)(__id, (void*)&parent, first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFileSystemModel);
	__conns_QFileSystemModel[__id] = __c;
}
void QFileSystemModel_connect_modelAboutToBeReset(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFileSystemModel*)self, &QFileSystemModel::modelAboutToBeReset, (QFileSystemModel*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFileSystemModel);
	__conns_QFileSystemModel[__id] = __c;
}
void QFileSystemModel_connect_modelReset(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFileSystemModel*)self, &QFileSystemModel::modelReset, (QFileSystemModel*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFileSystemModel);
	__conns_QFileSystemModel[__id] = __c;
}
void QFileSystemModel_connect_rowsAboutToBeMoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFileSystemModel*)self, &QFileSystemModel::rowsAboutToBeMoved, (QFileSystemModel*)self, [__id](const QModelIndex& sourceParent, int sourceStart, int sourceEnd, const QModelIndex& destinationParent, int destinationRow){
		if (cppgo_sig_void_voidp_int_int_voidp_int) ((void(*)(long long, void*, int, int, void*, int))cppgo_sig_void_voidp_int_int_voidp_int)(__id, (void*)&sourceParent, sourceStart, sourceEnd, (void*)&destinationParent, destinationRow);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFileSystemModel);
	__conns_QFileSystemModel[__id] = __c;
}
void QFileSystemModel_connect_rowsMoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFileSystemModel*)self, &QFileSystemModel::rowsMoved, (QFileSystemModel*)self, [__id](const QModelIndex& sourceParent, int sourceStart, int sourceEnd, const QModelIndex& destinationParent, int destinationRow){
		if (cppgo_sig_void_voidp_int_int_voidp_int) ((void(*)(long long, void*, int, int, void*, int))cppgo_sig_void_voidp_int_int_voidp_int)(__id, (void*)&sourceParent, sourceStart, sourceEnd, (void*)&destinationParent, destinationRow);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFileSystemModel);
	__conns_QFileSystemModel[__id] = __c;
}
void QFileSystemModel_connect_columnsAboutToBeMoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFileSystemModel*)self, &QFileSystemModel::columnsAboutToBeMoved, (QFileSystemModel*)self, [__id](const QModelIndex& sourceParent, int sourceStart, int sourceEnd, const QModelIndex& destinationParent, int destinationColumn){
		if (cppgo_sig_void_voidp_int_int_voidp_int) ((void(*)(long long, void*, int, int, void*, int))cppgo_sig_void_voidp_int_int_voidp_int)(__id, (void*)&sourceParent, sourceStart, sourceEnd, (void*)&destinationParent, destinationColumn);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFileSystemModel);
	__conns_QFileSystemModel[__id] = __c;
}
void QFileSystemModel_connect_columnsMoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFileSystemModel*)self, &QFileSystemModel::columnsMoved, (QFileSystemModel*)self, [__id](const QModelIndex& sourceParent, int sourceStart, int sourceEnd, const QModelIndex& destinationParent, int destinationColumn){
		if (cppgo_sig_void_voidp_int_int_voidp_int) ((void(*)(long long, void*, int, int, void*, int))cppgo_sig_void_voidp_int_int_voidp_int)(__id, (void*)&sourceParent, sourceStart, sourceEnd, (void*)&destinationParent, destinationColumn);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFileSystemModel);
	__conns_QFileSystemModel[__id] = __c;
}
void QFileSystemModel_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFileSystemModel*)self, &QFileSystemModel::destroyed, (QFileSystemModel*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFileSystemModel);
	__conns_QFileSystemModel[__id] = __c;
}
void QFileSystemModel_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFileSystemModel*)self, &QFileSystemModel::destroyed, (QFileSystemModel*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFileSystemModel);
	__conns_QFileSystemModel[__id] = __c;
}
void QFileSystemModel_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFileSystemModel*)self, &QFileSystemModel::objectNameChanged, (QFileSystemModel*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFileSystemModel);
	__conns_QFileSystemModel[__id] = __c;
}
void QFileSystemModel__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QFileSystemModel);
		auto __it = __conns_QFileSystemModel.find(id);
		if (__it == __conns_QFileSystemModel.end()) return;
		__c = __it->second;
		__conns_QFileSystemModel.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QFileSystemModel__Go_new(long long id, void* parent) { return new QFileSystemModel__Go(id, (QObject*)parent); }
void QFileSystemModel_delete(void* self) { delete (QFileSystemModel*)self; }

}

#include <QtGui/qpageranges.h>
#include "gen_qpageranges__range.h"

extern "C" {

void* QPageRanges__Range_new() {
	return new QPageRanges::Range();
}
void* QPageRanges__Range_new2(void* arg1) {
	return new QPageRanges::Range(*(QPageRanges::Range*)arg1);
}
int QPageRanges__Range_contains(void* self, int pageNumber) {
	auto&& __r = ((QPageRanges::Range*)self)->contains(pageNumber);
	return __r;
}
int QPageRanges__Range_OperatorEqual(void* self, void* rhs) {
	auto&& __r = (*(QPageRanges::Range*)self) == (*(QPageRanges::Range*)rhs);
	return __r;
}
int QPageRanges__Range_OperatorNotEqual(void* self, void* rhs) {
	auto&& __r = (*(QPageRanges::Range*)self) != (*(QPageRanges::Range*)rhs);
	return __r;
}
int QPageRanges__Range_OperatorLesser(void* self, void* rhs) {
	auto&& __r = (*(QPageRanges::Range*)self) < (*(QPageRanges::Range*)rhs);
	return __r;
}
int QPageRanges__Range_from_get(void* self) {
	auto __r = ((QPageRanges::Range*)self)->from;
	return __r;
}
void QPageRanges__Range_from_set(void* self, int from) {
	((QPageRanges::Range*)self)->from = from;
}
int QPageRanges__Range_to_get(void* self) {
	auto __r = ((QPageRanges::Range*)self)->to;
	return __r;
}
void QPageRanges__Range_to_set(void* self, int to) {
	((QPageRanges::Range*)self)->to = to;
}
void QPageRanges__Range_delete(void* self) { delete (QPageRanges::Range*)self; }

}

#include <QtGui/qstandarditemmodel.h>
#include <QtCore/qabstractitemmodel.h>
#include <QtCore/qobject.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qvariant.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qmimedata.h>
#include <QtCore/qsize.h>
#include <QtCore/qdatastream.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qstandarditemmodel.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QStandardItemModel__Go : public QStandardItemModel {
public:
	long long __goId;
	unsigned char __goOn[42] = {};
	template<typename... A> QStandardItemModel__Go(long long id, A&&... a) : QStandardItemModel(std::forward<A>(a)...), __goId(id) {}
	using QStandardItemModel::createIndex;
	using QStandardItemModel::encodeData;
	using QStandardItemModel::decodeData;
	using QStandardItemModel::beginInsertRows;
	using QStandardItemModel::endInsertRows;
	using QStandardItemModel::beginRemoveRows;
	using QStandardItemModel::endRemoveRows;
	using QStandardItemModel::beginMoveRows;
	using QStandardItemModel::endMoveRows;
	using QStandardItemModel::beginInsertColumns;
	using QStandardItemModel::endInsertColumns;
	using QStandardItemModel::beginRemoveColumns;
	using QStandardItemModel::endRemoveColumns;
	using QStandardItemModel::beginMoveColumns;
	using QStandardItemModel::endMoveColumns;
	using QStandardItemModel::beginResetModel;
	using QStandardItemModel::endResetModel;
	using QStandardItemModel::changePersistentIndex;
	using QStandardItemModel::changePersistentIndexList;
	using QStandardItemModel::persistentIndexList;
	using QStandardItemModel::sender;
	using QStandardItemModel::senderSignalIndex;
	using QStandardItemModel::receivers;
	using QStandardItemModel::isSignalConnected;
	void cppgoBase_resetInternalData() { QStandardItemModel::resetInternalData(); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QStandardItemModel::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QStandardItemModel::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QStandardItemModel::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QStandardItemModel::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QStandardItemModel::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QStandardItemModel::metaObject();
	}
	QModelIndex index(int row, int column, const QModelIndex& parent) const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_voidp_int_int_voidp) { void* __r = ((void*(*)(long long, long long, int, int, void*, int*))cppgo_virt_voidp_int_int_voidp)(__goId, -8948777187306027381LL, row, column, (void*)&parent, &__h); if (__h) return *(QModelIndex*)__r; }
		return QStandardItemModel::index(row, column, parent);
	}
	QModelIndex parent(const QModelIndex& child) const override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 7371339085820260189LL, (void*)&child, &__h); if (__h) return *(QModelIndex*)__r; }
		return QStandardItemModel::parent(child);
	}
	int rowCount(const QModelIndex& parent) const override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1960960171956884010LL, (void*)&parent, &__h); if (__h) return __r; }
		return QStandardItemModel::rowCount(parent);
	}
	int columnCount(const QModelIndex& parent) const override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, -8256963024669584834LL, (void*)&parent, &__h); if (__h) return __r; }
		return QStandardItemModel::columnCount(parent);
	}
	bool hasChildren(const QModelIndex& parent) const override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 3369452259144514380LL, (void*)&parent, &__h); if (__h) return __r; }
		return QStandardItemModel::hasChildren(parent);
	}
	QVariant data(const QModelIndex& index, int role) const override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_voidp_voidp_int) { void* __r = ((void*(*)(long long, long long, void*, int, int*))cppgo_virt_voidp_voidp_int)(__goId, -8837375942098466299LL, (void*)&index, role, &__h); if (__h) return *(QVariant*)__r; }
		return QStandardItemModel::data(index, role);
	}
	void multiData(const QModelIndex& index, QModelRoleDataSpan roleDataSpan) const override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp_voidp) { ((void(*)(long long, long long, void*, void*, int*))cppgo_virt_void_voidp_voidp)(__goId, 731311060944096642LL, (void*)&index, (void*)&roleDataSpan, &__h); if (__h) return; }
		QStandardItemModel::multiData(index, roleDataSpan);
	}
	bool setData(const QModelIndex& index, const QVariant& value, int role) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_int_voidp_voidp_int) { int __r = ((int(*)(long long, long long, void*, void*, int, int*))cppgo_virt_int_voidp_voidp_int)(__goId, 35191454575202219LL, (void*)&index, (void*)&value, role, &__h); if (__h) return __r; }
		return QStandardItemModel::setData(index, value, role);
	}
	bool clearItemData(const QModelIndex& index) override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 8914190500520542409LL, (void*)&index, &__h); if (__h) return __r; }
		return QStandardItemModel::clearItemData(index);
	}
	QVariant headerData(int section, Qt::Orientation orientation, int role) const override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_voidp_int_int_int) { void* __r = ((void*(*)(long long, long long, int, int, int, int*))cppgo_virt_voidp_int_int_int)(__goId, -3297140021896693688LL, section, (int)(orientation), role, &__h); if (__h) return *(QVariant*)__r; }
		return QStandardItemModel::headerData(section, orientation, role);
	}
	bool setHeaderData(int section, Qt::Orientation orientation, const QVariant& value, int role) override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_int_int_int_voidp_int) { int __r = ((int(*)(long long, long long, int, int, void*, int, int*))cppgo_virt_int_int_int_voidp_int)(__goId, 4852200306590054850LL, section, (int)(orientation), (void*)&value, role, &__h); if (__h) return __r; }
		return QStandardItemModel::setHeaderData(section, orientation, value, role);
	}
	bool insertRows(int row, int count, const QModelIndex& parent) override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_int_int_int_voidp) { int __r = ((int(*)(long long, long long, int, int, void*, int*))cppgo_virt_int_int_int_voidp)(__goId, -2775729825962395603LL, row, count, (void*)&parent, &__h); if (__h) return __r; }
		return QStandardItemModel::insertRows(row, count, parent);
	}
	bool insertColumns(int column, int count, const QModelIndex& parent) override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_int_int_int_voidp) { int __r = ((int(*)(long long, long long, int, int, void*, int*))cppgo_virt_int_int_int_voidp)(__goId, -8973293846230792237LL, column, count, (void*)&parent, &__h); if (__h) return __r; }
		return QStandardItemModel::insertColumns(column, count, parent);
	}
	bool removeRows(int row, int count, const QModelIndex& parent) override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_int_int_int_voidp) { int __r = ((int(*)(long long, long long, int, int, void*, int*))cppgo_virt_int_int_int_voidp)(__goId, 5104072549182222656LL, row, count, (void*)&parent, &__h); if (__h) return __r; }
		return QStandardItemModel::removeRows(row, count, parent);
	}
	bool removeColumns(int column, int count, const QModelIndex& parent) override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_int_int_int_voidp) { int __r = ((int(*)(long long, long long, int, int, void*, int*))cppgo_virt_int_int_int_voidp)(__goId, 2078548737056540420LL, column, count, (void*)&parent, &__h); if (__h) return __r; }
		return QStandardItemModel::removeColumns(column, count, parent);
	}
	QFlags<Qt::ItemFlag> flags(const QModelIndex& index) const override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1703382833545304812LL, (void*)&index, &__h); if (__h) return QFlags<Qt::ItemFlag>(__r); }
		return QStandardItemModel::flags(index);
	}
	QFlags<Qt::DropAction> supportedDropActions() const override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -7784642176186821205LL, &__h); if (__h) return QFlags<Qt::DropAction>(__r); }
		return QStandardItemModel::supportedDropActions();
	}
	void sort(int column, Qt::SortOrder order) override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_void_int_int) { ((void(*)(long long, long long, int, int, int*))cppgo_virt_void_int_int)(__goId, 6912441694453382161LL, column, (int)(order), &__h); if (__h) return; }
		QStandardItemModel::sort(column, order);
	}
	QList<QString> mimeTypes() const override {
		int __h = 0;
		if (__goOn[19] && cppgo_virtp_voidp) {
			long long __len = 0;
			void* __a = ((void*(*)(long long, long long, long long*, int*))cppgo_virtp_voidp)(__goId, 5464467312394730418LL, &__len, &__h);
			if (__h) {
				QList<QString> __r; __r.reserve(__len);
				char** __aa = (char**)__a;
				for (long long __i = 0; __i < __len; __i++) { __r.append(QString::fromUtf8(__aa[__i])); free(__aa[__i]); }
				if (__a) free(__a);
				return __r;
			}
		}
		return QStandardItemModel::mimeTypes();
	}
	QMimeData* mimeData(const QList<QModelIndex>& indexes) const override {
		int __h = 0;
		std::vector<void*> indexes_v; indexes_v.reserve(indexes.size());
		for (long long __i = 0; __i < indexes.size(); __i++) indexes_v.push_back((void*)&indexes[__i]);
		if (__goOn[20] && cppgo_virt_voidp_voidpp_longlong) { void* __r = ((void*(*)(long long, long long, void**, long long, int*))cppgo_virt_voidp_voidpp_longlong)(__goId, 5175297091024974749LL, indexes_v.data(), (long long)indexes.size(), &__h); if (__h) return (QMimeData*)__r; }
		return QStandardItemModel::mimeData(indexes);
	}
	bool dropMimeData(const QMimeData* data, Qt::DropAction action, int row, int column, const QModelIndex& parent) override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_int_voidp_int_int_int_voidp) { int __r = ((int(*)(long long, long long, void*, int, int, int, void*, int*))cppgo_virt_int_voidp_int_int_int_voidp)(__goId, -9049994684305682212LL, (void*)data, (int)(action), row, column, (void*)&parent, &__h); if (__h) return __r; }
		return QStandardItemModel::dropMimeData(data, action, row, column, parent);
	}
	QModelIndex sibling(int row, int column, const QModelIndex& idx) const override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_voidp_int_int_voidp) { void* __r = ((void*(*)(long long, long long, int, int, void*, int*))cppgo_virt_voidp_int_int_voidp)(__goId, 397329346813782521LL, row, column, (void*)&idx, &__h); if (__h) return *(QModelIndex*)__r; }
		return QStandardItemModel::sibling(row, column, idx);
	}
	bool canDropMimeData(const QMimeData* data, Qt::DropAction action, int row, int column, const QModelIndex& parent) const override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_int_voidp_int_int_int_voidp) { int __r = ((int(*)(long long, long long, void*, int, int, int, void*, int*))cppgo_virt_int_voidp_int_int_int_voidp)(__goId, 731539533644753052LL, (void*)data, (int)(action), row, column, (void*)&parent, &__h); if (__h) return __r; }
		return QStandardItemModel::canDropMimeData(data, action, row, column, parent);
	}
	QFlags<Qt::DropAction> supportedDragActions() const override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, 3421601386088570762LL, &__h); if (__h) return QFlags<Qt::DropAction>(__r); }
		return QStandardItemModel::supportedDragActions();
	}
	bool moveRows(const QModelIndex& sourceParent, int sourceRow, int count, const QModelIndex& destinationParent, int destinationChild) override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_int_voidp_int_int_voidp_int) { int __r = ((int(*)(long long, long long, void*, int, int, void*, int, int*))cppgo_virt_int_voidp_int_int_voidp_int)(__goId, -6074819356947405415LL, (void*)&sourceParent, sourceRow, count, (void*)&destinationParent, destinationChild, &__h); if (__h) return __r; }
		return QStandardItemModel::moveRows(sourceParent, sourceRow, count, destinationParent, destinationChild);
	}
	bool moveColumns(const QModelIndex& sourceParent, int sourceColumn, int count, const QModelIndex& destinationParent, int destinationChild) override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_int_voidp_int_int_voidp_int) { int __r = ((int(*)(long long, long long, void*, int, int, void*, int, int*))cppgo_virt_int_voidp_int_int_voidp_int)(__goId, -6321007591388309905LL, (void*)&sourceParent, sourceColumn, count, (void*)&destinationParent, destinationChild, &__h); if (__h) return __r; }
		return QStandardItemModel::moveColumns(sourceParent, sourceColumn, count, destinationParent, destinationChild);
	}
	void fetchMore(const QModelIndex& parent) override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3553294593611130806LL, (void*)&parent, &__h); if (__h) return; }
		QStandardItemModel::fetchMore(parent);
	}
	bool canFetchMore(const QModelIndex& parent) const override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 8639262066473964150LL, (void*)&parent, &__h); if (__h) return __r; }
		return QStandardItemModel::canFetchMore(parent);
	}
	QModelIndex buddy(const QModelIndex& index) const override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 6112860112248786499LL, (void*)&index, &__h); if (__h) return *(QModelIndex*)__r; }
		return QStandardItemModel::buddy(index);
	}
	QList<QModelIndex> match(const QModelIndex& start, int role, const QVariant& value, int hits, QFlags<Qt::MatchFlag> flags) const override {
		int __h = 0;
		if (__goOn[30] && cppgo_virtp_voidp_voidp_int_voidp_int_int) {
			long long __len = 0;
			void* __a = ((void*(*)(long long, long long, void*, int, void*, int, int, long long*, int*))cppgo_virtp_voidp_voidp_int_voidp_int_int)(__goId, -4341501217981328394LL, (void*)&start, role, (void*)&value, hits, (int)(flags), &__len, &__h);
			if (__h) {
				QList<QModelIndex> __r; __r.reserve(__len);
				void** __aa = (void**)__a;
				for (long long __i = 0; __i < __len; __i++) __r.append(*(QModelIndex*)__aa[__i]);
				if (__a) free(__a);
				return __r;
			}
		}
		return QStandardItemModel::match(start, role, value, hits, flags);
	}
	QSize span(const QModelIndex& index) const override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, -8395342916587040543LL, (void*)&index, &__h); if (__h) return *(QSize*)__r; }
		return QStandardItemModel::span(index);
	}
	bool submit() override {
		int __h = 0;
		if (__goOn[32] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -3788331537442685179LL, &__h); if (__h) return __r; }
		return QStandardItemModel::submit();
	}
	void revert() override {
		int __h = 0;
		if (__goOn[33] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, -1522951607536117149LL, &__h); if (__h) return; }
		QStandardItemModel::revert();
	}
	void resetInternalData() override {
		int __h = 0;
		if (__goOn[34] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, -4510675770784051183LL, &__h); if (__h) return; }
		QStandardItemModel::resetInternalData();
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[35] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QStandardItemModel::event(event);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[36] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QStandardItemModel::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[37] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QStandardItemModel::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[38] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QStandardItemModel::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[39] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QStandardItemModel::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[40] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QStandardItemModel::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[41] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QStandardItemModel::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QStandardItemModel;
static std::mutex __connsMu_QStandardItemModel;

extern "C" {

void* QStandardItemModel_base(void* self) { return (void*)static_cast<QAbstractItemModel*>((QStandardItemModel*)self); }
void QStandardItemModel_chain(void* self, void** out) {
	QStandardItemModel* __p = (QStandardItemModel*)self;
	QAbstractItemModel* __b0 = static_cast<QAbstractItemModel*>(__p); out[0] = (void*)__b0;
	QObject* __b1 = static_cast<QObject*>(__b0); out[1] = (void*)__b1;
}
void QStandardItemModel__set_override(void* self, int idx) { ((QStandardItemModel__Go*)self)->__goOn[idx] = 1; }
void* QStandardItemModel_fromQObject(void* o) { return (void*)qobject_cast<QStandardItemModel*>((QObject*)o); }
void* QStandardItemModel_new(long long __goid) {
	return new QStandardItemModel__Go(__goid);
}
void* QStandardItemModel_new2(long long __goid, int rows, int columns) {
	return new QStandardItemModel__Go(__goid, rows, columns);
}
void* QStandardItemModel_new3(long long __goid, void* parent) {
	return new QStandardItemModel__Go(__goid, (QObject*)parent);
}
void* QStandardItemModel_new4(long long __goid, int rows, int columns, void* parent) {
	return new QStandardItemModel__Go(__goid, rows, columns, (QObject*)parent);
}
char* QStandardItemModel_tr(const char* s, long long* __out_len) {
	auto&& __r = QStandardItemModel::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QStandardItemModel_setItemRoleNames(void* self, int* roleNames_keys, void* roleNames_vals, long long roleNames_len) {
	QHash<int, QByteArray> roleNames_map;
	void** roleNames_vv = (void**)roleNames_vals;
	for (long long __i = 0; __i < roleNames_len; __i++) roleNames_map.insert(roleNames_keys[__i], *(QByteArray*)roleNames_vv[__i]);
	((QStandardItemModel*)self)->setItemRoleNames(roleNames_map);
}
void QStandardItemModel_clear(void* self) {
	((QStandardItemModel*)self)->clear();
}
void* QStandardItemModel_itemFromIndex(void* self, void* index) {
	auto&& __r = ((QStandardItemModel*)self)->itemFromIndex(*(QModelIndex*)index);
	return (void*)(__r);
}
void* QStandardItemModel_indexFromItem(void* self, void* item) {
	auto&& __r = ((QStandardItemModel*)self)->indexFromItem((QStandardItem*)item);
	return new QModelIndex(__r);
}
void* QStandardItemModel_item(void* self, int row) {
	auto&& __r = ((QStandardItemModel*)self)->item(row);
	return (void*)(__r);
}
void QStandardItemModel_setItem(void* self, int row, int column, void* item) {
	((QStandardItemModel*)self)->setItem(row, column, (QStandardItem*)item);
}
void QStandardItemModel_setItem2(void* self, int row, void* item) {
	((QStandardItemModel*)self)->setItem(row, (QStandardItem*)item);
}
void* QStandardItemModel_invisibleRootItem(void* self) {
	auto&& __r = ((QStandardItemModel*)self)->invisibleRootItem();
	return (void*)(__r);
}
void* QStandardItemModel_horizontalHeaderItem(void* self, int column) {
	auto&& __r = ((QStandardItemModel*)self)->horizontalHeaderItem(column);
	return (void*)(__r);
}
void QStandardItemModel_setHorizontalHeaderItem(void* self, int column, void* item) {
	((QStandardItemModel*)self)->setHorizontalHeaderItem(column, (QStandardItem*)item);
}
void* QStandardItemModel_verticalHeaderItem(void* self, int row) {
	auto&& __r = ((QStandardItemModel*)self)->verticalHeaderItem(row);
	return (void*)(__r);
}
void QStandardItemModel_setVerticalHeaderItem(void* self, int row, void* item) {
	((QStandardItemModel*)self)->setVerticalHeaderItem(row, (QStandardItem*)item);
}
void QStandardItemModel_setHorizontalHeaderLabels(void* self, char** labels, long long labels_len) {
	QList<QString> labels_list; labels_list.reserve(labels_len);
	for (long long __i = 0; __i < labels_len; __i++) labels_list.append(QString::fromUtf8(labels[__i]));
	((QStandardItemModel*)self)->setHorizontalHeaderLabels(labels_list);
}
void QStandardItemModel_setVerticalHeaderLabels(void* self, char** labels, long long labels_len) {
	QList<QString> labels_list; labels_list.reserve(labels_len);
	for (long long __i = 0; __i < labels_len; __i++) labels_list.append(QString::fromUtf8(labels[__i]));
	((QStandardItemModel*)self)->setVerticalHeaderLabels(labels_list);
}
void QStandardItemModel_setRowCount(void* self, int rows) {
	((QStandardItemModel*)self)->setRowCount(rows);
}
void QStandardItemModel_setColumnCount(void* self, int columns) {
	((QStandardItemModel*)self)->setColumnCount(columns);
}
void QStandardItemModel_appendRow(void* self, void** items, long long items_len) {
	QList<QStandardItem*> items_list; items_list.reserve(items_len);
	for (long long __i = 0; __i < items_len; __i++) items_list.append((QStandardItem*)items[__i]);
	((QStandardItemModel*)self)->appendRow(items_list);
}
void QStandardItemModel_appendColumn(void* self, void** items, long long items_len) {
	QList<QStandardItem*> items_list; items_list.reserve(items_len);
	for (long long __i = 0; __i < items_len; __i++) items_list.append((QStandardItem*)items[__i]);
	((QStandardItemModel*)self)->appendColumn(items_list);
}
void QStandardItemModel_appendRow2(void* self, void* item) {
	((QStandardItemModel*)self)->appendRow((QStandardItem*)item);
}
void QStandardItemModel_insertRow(void* self, int row, void** items, long long items_len) {
	QList<QStandardItem*> items_list; items_list.reserve(items_len);
	for (long long __i = 0; __i < items_len; __i++) items_list.append((QStandardItem*)items[__i]);
	((QStandardItemModel*)self)->insertRow(row, items_list);
}
void QStandardItemModel_insertColumn(void* self, int column, void** items, long long items_len) {
	QList<QStandardItem*> items_list; items_list.reserve(items_len);
	for (long long __i = 0; __i < items_len; __i++) items_list.append((QStandardItem*)items[__i]);
	((QStandardItemModel*)self)->insertColumn(column, items_list);
}
void QStandardItemModel_insertRow2(void* self, int row, void* item) {
	((QStandardItemModel*)self)->insertRow(row, (QStandardItem*)item);
}
int QStandardItemModel_insertRow3(void* self, int row) {
	auto&& __r = ((QStandardItemModel*)self)->insertRow(row);
	return __r;
}
int QStandardItemModel_insertColumn2(void* self, int column) {
	auto&& __r = ((QStandardItemModel*)self)->insertColumn(column);
	return __r;
}
void* QStandardItemModel_takeItem(void* self, int row) {
	auto&& __r = ((QStandardItemModel*)self)->takeItem(row);
	return (void*)(__r);
}
void** QStandardItemModel_takeRow(void* self, int row, long long* __out_len) {
	auto&& __r = ((QStandardItemModel*)self)->takeRow(row);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
void** QStandardItemModel_takeColumn(void* self, int column, long long* __out_len) {
	auto&& __r = ((QStandardItemModel*)self)->takeColumn(column);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
void* QStandardItemModel_takeHorizontalHeaderItem(void* self, int column) {
	auto&& __r = ((QStandardItemModel*)self)->takeHorizontalHeaderItem(column);
	return (void*)(__r);
}
void* QStandardItemModel_takeVerticalHeaderItem(void* self, int row) {
	auto&& __r = ((QStandardItemModel*)self)->takeVerticalHeaderItem(row);
	return (void*)(__r);
}
void* QStandardItemModel_itemPrototype(void* self) {
	auto&& __r = ((QStandardItemModel*)self)->itemPrototype();
	return (void*)(__r);
}
void QStandardItemModel_setItemPrototype(void* self, void* item) {
	((QStandardItemModel*)self)->setItemPrototype((QStandardItem*)item);
}
void** QStandardItemModel_findItems(void* self, const char* text, long long text_len, long long* __out_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = ((QStandardItemModel*)self)->findItems(text_qstr);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
int QStandardItemModel_sortRole(void* self) {
	auto&& __r = ((QStandardItemModel*)self)->sortRole();
	return __r;
}
void QStandardItemModel_setSortRole(void* self, int role) {
	((QStandardItemModel*)self)->setSortRole(role);
}
void QStandardItemModel_itemChanged(void* self, void* item) {
	((QStandardItemModel*)self)->itemChanged((QStandardItem*)item);
}
void* QStandardItemModel_createIndex(void* self, int row, int column) {
	auto&& __r = ((QStandardItemModel__Go*)self)->createIndex(row, column);
	return new QModelIndex(__r);
}
void* QStandardItemModel_createIndex2(void* self, int row, int column, unsigned long long id) {
	auto&& __r = ((QStandardItemModel__Go*)self)->createIndex(row, column, id);
	return new QModelIndex(__r);
}
void QStandardItemModel_encodeData(void* self, void** indexes, long long indexes_len, void* stream) {
	QList<QModelIndex> indexes_list; indexes_list.reserve(indexes_len);
	for (long long __i = 0; __i < indexes_len; __i++) indexes_list.append(*(QModelIndex*)indexes[__i]);
	((QStandardItemModel__Go*)self)->encodeData(indexes_list, *(QDataStream*)stream);
}
int QStandardItemModel_decodeData(void* self, int row, int column, void* parent, void* stream) {
	auto&& __r = ((QStandardItemModel__Go*)self)->decodeData(row, column, *(QModelIndex*)parent, *(QDataStream*)stream);
	return __r;
}
void QStandardItemModel_beginInsertRows(void* self, void* parent, int first, int last) {
	((QStandardItemModel__Go*)self)->beginInsertRows(*(QModelIndex*)parent, first, last);
}
void QStandardItemModel_endInsertRows(void* self) {
	((QStandardItemModel__Go*)self)->endInsertRows();
}
void QStandardItemModel_beginRemoveRows(void* self, void* parent, int first, int last) {
	((QStandardItemModel__Go*)self)->beginRemoveRows(*(QModelIndex*)parent, first, last);
}
void QStandardItemModel_endRemoveRows(void* self) {
	((QStandardItemModel__Go*)self)->endRemoveRows();
}
int QStandardItemModel_beginMoveRows(void* self, void* sourceParent, int sourceFirst, int sourceLast, void* destinationParent, int destinationRow) {
	auto&& __r = ((QStandardItemModel__Go*)self)->beginMoveRows(*(QModelIndex*)sourceParent, sourceFirst, sourceLast, *(QModelIndex*)destinationParent, destinationRow);
	return __r;
}
void QStandardItemModel_endMoveRows(void* self) {
	((QStandardItemModel__Go*)self)->endMoveRows();
}
void QStandardItemModel_beginInsertColumns(void* self, void* parent, int first, int last) {
	((QStandardItemModel__Go*)self)->beginInsertColumns(*(QModelIndex*)parent, first, last);
}
void QStandardItemModel_endInsertColumns(void* self) {
	((QStandardItemModel__Go*)self)->endInsertColumns();
}
void QStandardItemModel_beginRemoveColumns(void* self, void* parent, int first, int last) {
	((QStandardItemModel__Go*)self)->beginRemoveColumns(*(QModelIndex*)parent, first, last);
}
void QStandardItemModel_endRemoveColumns(void* self) {
	((QStandardItemModel__Go*)self)->endRemoveColumns();
}
int QStandardItemModel_beginMoveColumns(void* self, void* sourceParent, int sourceFirst, int sourceLast, void* destinationParent, int destinationColumn) {
	auto&& __r = ((QStandardItemModel__Go*)self)->beginMoveColumns(*(QModelIndex*)sourceParent, sourceFirst, sourceLast, *(QModelIndex*)destinationParent, destinationColumn);
	return __r;
}
void QStandardItemModel_endMoveColumns(void* self) {
	((QStandardItemModel__Go*)self)->endMoveColumns();
}
void QStandardItemModel_beginResetModel(void* self) {
	((QStandardItemModel__Go*)self)->beginResetModel();
}
void QStandardItemModel_endResetModel(void* self) {
	((QStandardItemModel__Go*)self)->endResetModel();
}
void QStandardItemModel_changePersistentIndex(void* self, void* from, void* to) {
	((QStandardItemModel__Go*)self)->changePersistentIndex(*(QModelIndex*)from, *(QModelIndex*)to);
}
void QStandardItemModel_changePersistentIndexList(void* self, void** from, long long from_len, void** to, long long to_len) {
	QList<QModelIndex> from_list; from_list.reserve(from_len);
	for (long long __i = 0; __i < from_len; __i++) from_list.append(*(QModelIndex*)from[__i]);
	QList<QModelIndex> to_list; to_list.reserve(to_len);
	for (long long __i = 0; __i < to_len; __i++) to_list.append(*(QModelIndex*)to[__i]);
	((QStandardItemModel__Go*)self)->changePersistentIndexList(from_list, to_list);
}
void** QStandardItemModel_persistentIndexList(void* self, long long* __out_len) {
	auto&& __r = ((QStandardItemModel__Go*)self)->persistentIndexList();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QModelIndex(__e);
	return __arr;
}
void* QStandardItemModel_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QStandardItemModel::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QStandardItemModel_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QStandardItemModel::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QStandardItemModel_sender(void* self) {
	auto&& __r = ((QStandardItemModel__Go*)self)->sender();
	return (void*)(__r);
}
int QStandardItemModel_senderSignalIndex(void* self) {
	auto&& __r = ((QStandardItemModel__Go*)self)->senderSignalIndex();
	return __r;
}
int QStandardItemModel_receivers(void* self, const char* signal) {
	auto&& __r = ((QStandardItemModel__Go*)self)->receivers(signal);
	return __r;
}
int QStandardItemModel_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QStandardItemModel__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QStandardItemModel_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QStandardItemModel::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QStandardItemModel_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QStandardItemModel::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QStandardItemModel_item2(void* self, int row, int column) {
	auto&& __r = ((QStandardItemModel*)self)->item(row, column);
	return (void*)(__r);
}
int QStandardItemModel_insertRow4(void* self, int row, void* parent) {
	auto&& __r = ((QStandardItemModel*)self)->insertRow(row, *(QModelIndex*)parent);
	return __r;
}
int QStandardItemModel_insertColumn3(void* self, int column, void* parent) {
	auto&& __r = ((QStandardItemModel*)self)->insertColumn(column, *(QModelIndex*)parent);
	return __r;
}
void* QStandardItemModel_takeItem2(void* self, int row, int column) {
	auto&& __r = ((QStandardItemModel*)self)->takeItem(row, column);
	return (void*)(__r);
}
void** QStandardItemModel_findItems2(void* self, const char* text, long long text_len, int flags, long long* __out_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = ((QStandardItemModel*)self)->findItems(text_qstr, QFlags<Qt::MatchFlag>(flags));
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
void** QStandardItemModel_findItems3(void* self, const char* text, long long text_len, int flags, int column, long long* __out_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = ((QStandardItemModel*)self)->findItems(text_qstr, QFlags<Qt::MatchFlag>(flags), column);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
void* QStandardItemModel_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QStandardItemModel::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QStandardItemModel_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QStandardItemModel::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QStandardItemModel_metaObject_default(void* self) {
	auto&& __r = ((QStandardItemModel*)self)->QStandardItemModel::metaObject();
	return (void*)(__r);
}
void* QStandardItemModel_index_default(void* self, int row, int column, void* parent) {
	auto&& __r = ((QStandardItemModel*)self)->QStandardItemModel::index(row, column, *(QModelIndex*)parent);
	return new QModelIndex(__r);
}
void* QStandardItemModel_parent_default(void* self, void* child) {
	auto&& __r = ((QStandardItemModel*)self)->QStandardItemModel::parent(*(QModelIndex*)child);
	return new QModelIndex(__r);
}
int QStandardItemModel_rowCount_default(void* self, void* parent) {
	auto&& __r = ((QStandardItemModel*)self)->QStandardItemModel::rowCount(*(QModelIndex*)parent);
	return __r;
}
int QStandardItemModel_columnCount_default(void* self, void* parent) {
	auto&& __r = ((QStandardItemModel*)self)->QStandardItemModel::columnCount(*(QModelIndex*)parent);
	return __r;
}
int QStandardItemModel_hasChildren_default(void* self, void* parent) {
	auto&& __r = ((QStandardItemModel*)self)->QStandardItemModel::hasChildren(*(QModelIndex*)parent);
	return __r;
}
void* QStandardItemModel_data_default(void* self, void* index, int role) {
	auto&& __r = ((QStandardItemModel*)self)->QStandardItemModel::data(*(QModelIndex*)index, role);
	return new QVariant(__r);
}
void QStandardItemModel_multiData_default(void* self, void* index, void* roleDataSpan) {
	((QStandardItemModel*)self)->QStandardItemModel::multiData(*(QModelIndex*)index, *(QModelRoleDataSpan*)roleDataSpan);
}
int QStandardItemModel_setData_default(void* self, void* index, void* value, int role) {
	auto&& __r = ((QStandardItemModel*)self)->QStandardItemModel::setData(*(QModelIndex*)index, *(QVariant*)value, role);
	return __r;
}
int QStandardItemModel_clearItemData_default(void* self, void* index) {
	auto&& __r = ((QStandardItemModel*)self)->QStandardItemModel::clearItemData(*(QModelIndex*)index);
	return __r;
}
void* QStandardItemModel_headerData_default(void* self, int section, int orientation, int role) {
	auto&& __r = ((QStandardItemModel*)self)->QStandardItemModel::headerData(section, (Qt::Orientation)(orientation), role);
	return new QVariant(__r);
}
int QStandardItemModel_setHeaderData_default(void* self, int section, int orientation, void* value, int role) {
	auto&& __r = ((QStandardItemModel*)self)->QStandardItemModel::setHeaderData(section, (Qt::Orientation)(orientation), *(QVariant*)value, role);
	return __r;
}
int QStandardItemModel_insertRows_default(void* self, int row, int count, void* parent) {
	auto&& __r = ((QStandardItemModel*)self)->QStandardItemModel::insertRows(row, count, *(QModelIndex*)parent);
	return __r;
}
int QStandardItemModel_insertColumns_default(void* self, int column, int count, void* parent) {
	auto&& __r = ((QStandardItemModel*)self)->QStandardItemModel::insertColumns(column, count, *(QModelIndex*)parent);
	return __r;
}
int QStandardItemModel_removeRows_default(void* self, int row, int count, void* parent) {
	auto&& __r = ((QStandardItemModel*)self)->QStandardItemModel::removeRows(row, count, *(QModelIndex*)parent);
	return __r;
}
int QStandardItemModel_removeColumns_default(void* self, int column, int count, void* parent) {
	auto&& __r = ((QStandardItemModel*)self)->QStandardItemModel::removeColumns(column, count, *(QModelIndex*)parent);
	return __r;
}
int QStandardItemModel_flags_default(void* self, void* index) {
	auto&& __r = ((QStandardItemModel*)self)->QStandardItemModel::flags(*(QModelIndex*)index);
	return (int)(__r);
}
int QStandardItemModel_supportedDropActions_default(void* self) {
	auto&& __r = ((QStandardItemModel*)self)->QStandardItemModel::supportedDropActions();
	return (int)(__r);
}
void QStandardItemModel_sort_default(void* self, int column, int order) {
	((QStandardItemModel*)self)->QStandardItemModel::sort(column, (Qt::SortOrder)(order));
}
char** QStandardItemModel_mimeTypes_default(void* self, long long* __out_len) {
	auto&& __r = ((QStandardItemModel*)self)->QStandardItemModel::mimeTypes();
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
void* QStandardItemModel_mimeData_default(void* self, void** indexes, long long indexes_len) {
	QList<QModelIndex> indexes_list; indexes_list.reserve(indexes_len);
	for (long long __i = 0; __i < indexes_len; __i++) indexes_list.append(*(QModelIndex*)indexes[__i]);
	auto&& __r = ((QStandardItemModel*)self)->QStandardItemModel::mimeData(indexes_list);
	return (void*)(__r);
}
int QStandardItemModel_dropMimeData_default(void* self, void* data, int action, int row, int column, void* parent) {
	auto&& __r = ((QStandardItemModel*)self)->QStandardItemModel::dropMimeData((QMimeData*)data, (Qt::DropAction)(action), row, column, *(QModelIndex*)parent);
	return __r;
}
void* QStandardItemModel_sibling_default(void* self, int row, int column, void* idx) {
	auto&& __r = ((QStandardItemModel*)self)->QStandardItemModel::sibling(row, column, *(QModelIndex*)idx);
	return new QModelIndex(__r);
}
int QStandardItemModel_canDropMimeData_default(void* self, void* data, int action, int row, int column, void* parent) {
	auto&& __r = ((QStandardItemModel*)self)->QStandardItemModel::canDropMimeData((QMimeData*)data, (Qt::DropAction)(action), row, column, *(QModelIndex*)parent);
	return __r;
}
int QStandardItemModel_supportedDragActions_default(void* self) {
	auto&& __r = ((QStandardItemModel*)self)->QStandardItemModel::supportedDragActions();
	return (int)(__r);
}
int QStandardItemModel_moveRows_default(void* self, void* sourceParent, int sourceRow, int count, void* destinationParent, int destinationChild) {
	auto&& __r = ((QStandardItemModel*)self)->QStandardItemModel::moveRows(*(QModelIndex*)sourceParent, sourceRow, count, *(QModelIndex*)destinationParent, destinationChild);
	return __r;
}
int QStandardItemModel_moveColumns_default(void* self, void* sourceParent, int sourceColumn, int count, void* destinationParent, int destinationChild) {
	auto&& __r = ((QStandardItemModel*)self)->QStandardItemModel::moveColumns(*(QModelIndex*)sourceParent, sourceColumn, count, *(QModelIndex*)destinationParent, destinationChild);
	return __r;
}
void QStandardItemModel_fetchMore_default(void* self, void* parent) {
	((QStandardItemModel*)self)->QStandardItemModel::fetchMore(*(QModelIndex*)parent);
}
int QStandardItemModel_canFetchMore_default(void* self, void* parent) {
	auto&& __r = ((QStandardItemModel*)self)->QStandardItemModel::canFetchMore(*(QModelIndex*)parent);
	return __r;
}
void* QStandardItemModel_buddy_default(void* self, void* index) {
	auto&& __r = ((QStandardItemModel*)self)->QStandardItemModel::buddy(*(QModelIndex*)index);
	return new QModelIndex(__r);
}
void** QStandardItemModel_match_default(void* self, void* start, int role, void* value, int hits, int flags, long long* __out_len) {
	auto&& __r = ((QStandardItemModel*)self)->QStandardItemModel::match(*(QModelIndex*)start, role, *(QVariant*)value, hits, QFlags<Qt::MatchFlag>(flags));
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QModelIndex(__e);
	return __arr;
}
void* QStandardItemModel_span_default(void* self, void* index) {
	auto&& __r = ((QStandardItemModel*)self)->QStandardItemModel::span(*(QModelIndex*)index);
	return new QSize(__r);
}
int QStandardItemModel_submit_default(void* self) {
	auto&& __r = ((QStandardItemModel*)self)->QStandardItemModel::submit();
	return __r;
}
void QStandardItemModel_revert_default(void* self) {
	((QStandardItemModel*)self)->QStandardItemModel::revert();
}
void QStandardItemModel_resetInternalData_default(void* self) {
	((QStandardItemModel__Go*)self)->cppgoBase_resetInternalData();
}
int QStandardItemModel_event_default(void* self, void* event) {
	auto&& __r = ((QStandardItemModel*)self)->QStandardItemModel::event((QEvent*)event);
	return __r;
}
int QStandardItemModel_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QStandardItemModel*)self)->QStandardItemModel::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QStandardItemModel_timerEvent_default(void* self, void* event) {
	((QStandardItemModel__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QStandardItemModel_childEvent_default(void* self, void* event) {
	((QStandardItemModel__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QStandardItemModel_customEvent_default(void* self, void* event) {
	((QStandardItemModel__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QStandardItemModel_connectNotify_default(void* self, void* signal) {
	((QStandardItemModel__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QStandardItemModel_disconnectNotify_default(void* self, void* signal) {
	((QStandardItemModel__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QStandardItemModel_connect_itemChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStandardItemModel*)self, &QStandardItemModel::itemChanged, (QStandardItemModel*)self, [__id](QStandardItem* item){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)item);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStandardItemModel);
	__conns_QStandardItemModel[__id] = __c;
}
void QStandardItemModel_connect_headerDataChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStandardItemModel*)self, &QStandardItemModel::headerDataChanged, (QStandardItemModel*)self, [__id](Qt::Orientation orientation, int first, int last){
		if (cppgo_sig_void_int_int_int) ((void(*)(long long, int, int, int))cppgo_sig_void_int_int_int)(__id, (int)(orientation), first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStandardItemModel);
	__conns_QStandardItemModel[__id] = __c;
}
void QStandardItemModel_connect_layoutChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStandardItemModel*)self, &QStandardItemModel::layoutChanged, (QStandardItemModel*)self, [__id](const QList<QPersistentModelIndex>& parents, QAbstractItemModel::LayoutChangeHint hint){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStandardItemModel);
	__conns_QStandardItemModel[__id] = __c;
}
void QStandardItemModel_connect_layoutChanged2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStandardItemModel*)self, &QStandardItemModel::layoutChanged, (QStandardItemModel*)self, [__id](const QList<QPersistentModelIndex>& parents, QAbstractItemModel::LayoutChangeHint hint){
		std::vector<void*> parents_v; parents_v.reserve(parents.size());
		for (long long __i = 0; __i < parents.size(); __i++) parents_v.push_back((void*)&parents[__i]);
		if (cppgo_sig_void_voidpp_longlong_int) ((void(*)(long long, void**, long long, int))cppgo_sig_void_voidpp_longlong_int)(__id, parents_v.data(), (long long)parents.size(), (int)(hint));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStandardItemModel);
	__conns_QStandardItemModel[__id] = __c;
}
void QStandardItemModel_connect_layoutAboutToBeChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStandardItemModel*)self, &QStandardItemModel::layoutAboutToBeChanged, (QStandardItemModel*)self, [__id](const QList<QPersistentModelIndex>& parents, QAbstractItemModel::LayoutChangeHint hint){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStandardItemModel);
	__conns_QStandardItemModel[__id] = __c;
}
void QStandardItemModel_connect_layoutAboutToBeChanged2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStandardItemModel*)self, &QStandardItemModel::layoutAboutToBeChanged, (QStandardItemModel*)self, [__id](const QList<QPersistentModelIndex>& parents, QAbstractItemModel::LayoutChangeHint hint){
		std::vector<void*> parents_v; parents_v.reserve(parents.size());
		for (long long __i = 0; __i < parents.size(); __i++) parents_v.push_back((void*)&parents[__i]);
		if (cppgo_sig_void_voidpp_longlong_int) ((void(*)(long long, void**, long long, int))cppgo_sig_void_voidpp_longlong_int)(__id, parents_v.data(), (long long)parents.size(), (int)(hint));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStandardItemModel);
	__conns_QStandardItemModel[__id] = __c;
}
void QStandardItemModel_connect_rowsAboutToBeInserted(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStandardItemModel*)self, &QStandardItemModel::rowsAboutToBeInserted, (QStandardItemModel*)self, [__id](const QModelIndex& parent, int first, int last){
		if (cppgo_sig_void_voidp_int_int) ((void(*)(long long, void*, int, int))cppgo_sig_void_voidp_int_int)(__id, (void*)&parent, first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStandardItemModel);
	__conns_QStandardItemModel[__id] = __c;
}
void QStandardItemModel_connect_rowsInserted(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStandardItemModel*)self, &QStandardItemModel::rowsInserted, (QStandardItemModel*)self, [__id](const QModelIndex& parent, int first, int last){
		if (cppgo_sig_void_voidp_int_int) ((void(*)(long long, void*, int, int))cppgo_sig_void_voidp_int_int)(__id, (void*)&parent, first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStandardItemModel);
	__conns_QStandardItemModel[__id] = __c;
}
void QStandardItemModel_connect_rowsAboutToBeRemoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStandardItemModel*)self, &QStandardItemModel::rowsAboutToBeRemoved, (QStandardItemModel*)self, [__id](const QModelIndex& parent, int first, int last){
		if (cppgo_sig_void_voidp_int_int) ((void(*)(long long, void*, int, int))cppgo_sig_void_voidp_int_int)(__id, (void*)&parent, first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStandardItemModel);
	__conns_QStandardItemModel[__id] = __c;
}
void QStandardItemModel_connect_rowsRemoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStandardItemModel*)self, &QStandardItemModel::rowsRemoved, (QStandardItemModel*)self, [__id](const QModelIndex& parent, int first, int last){
		if (cppgo_sig_void_voidp_int_int) ((void(*)(long long, void*, int, int))cppgo_sig_void_voidp_int_int)(__id, (void*)&parent, first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStandardItemModel);
	__conns_QStandardItemModel[__id] = __c;
}
void QStandardItemModel_connect_columnsAboutToBeInserted(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStandardItemModel*)self, &QStandardItemModel::columnsAboutToBeInserted, (QStandardItemModel*)self, [__id](const QModelIndex& parent, int first, int last){
		if (cppgo_sig_void_voidp_int_int) ((void(*)(long long, void*, int, int))cppgo_sig_void_voidp_int_int)(__id, (void*)&parent, first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStandardItemModel);
	__conns_QStandardItemModel[__id] = __c;
}
void QStandardItemModel_connect_columnsInserted(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStandardItemModel*)self, &QStandardItemModel::columnsInserted, (QStandardItemModel*)self, [__id](const QModelIndex& parent, int first, int last){
		if (cppgo_sig_void_voidp_int_int) ((void(*)(long long, void*, int, int))cppgo_sig_void_voidp_int_int)(__id, (void*)&parent, first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStandardItemModel);
	__conns_QStandardItemModel[__id] = __c;
}
void QStandardItemModel_connect_columnsAboutToBeRemoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStandardItemModel*)self, &QStandardItemModel::columnsAboutToBeRemoved, (QStandardItemModel*)self, [__id](const QModelIndex& parent, int first, int last){
		if (cppgo_sig_void_voidp_int_int) ((void(*)(long long, void*, int, int))cppgo_sig_void_voidp_int_int)(__id, (void*)&parent, first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStandardItemModel);
	__conns_QStandardItemModel[__id] = __c;
}
void QStandardItemModel_connect_columnsRemoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStandardItemModel*)self, &QStandardItemModel::columnsRemoved, (QStandardItemModel*)self, [__id](const QModelIndex& parent, int first, int last){
		if (cppgo_sig_void_voidp_int_int) ((void(*)(long long, void*, int, int))cppgo_sig_void_voidp_int_int)(__id, (void*)&parent, first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStandardItemModel);
	__conns_QStandardItemModel[__id] = __c;
}
void QStandardItemModel_connect_modelAboutToBeReset(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStandardItemModel*)self, &QStandardItemModel::modelAboutToBeReset, (QStandardItemModel*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStandardItemModel);
	__conns_QStandardItemModel[__id] = __c;
}
void QStandardItemModel_connect_modelReset(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStandardItemModel*)self, &QStandardItemModel::modelReset, (QStandardItemModel*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStandardItemModel);
	__conns_QStandardItemModel[__id] = __c;
}
void QStandardItemModel_connect_rowsAboutToBeMoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStandardItemModel*)self, &QStandardItemModel::rowsAboutToBeMoved, (QStandardItemModel*)self, [__id](const QModelIndex& sourceParent, int sourceStart, int sourceEnd, const QModelIndex& destinationParent, int destinationRow){
		if (cppgo_sig_void_voidp_int_int_voidp_int) ((void(*)(long long, void*, int, int, void*, int))cppgo_sig_void_voidp_int_int_voidp_int)(__id, (void*)&sourceParent, sourceStart, sourceEnd, (void*)&destinationParent, destinationRow);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStandardItemModel);
	__conns_QStandardItemModel[__id] = __c;
}
void QStandardItemModel_connect_rowsMoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStandardItemModel*)self, &QStandardItemModel::rowsMoved, (QStandardItemModel*)self, [__id](const QModelIndex& sourceParent, int sourceStart, int sourceEnd, const QModelIndex& destinationParent, int destinationRow){
		if (cppgo_sig_void_voidp_int_int_voidp_int) ((void(*)(long long, void*, int, int, void*, int))cppgo_sig_void_voidp_int_int_voidp_int)(__id, (void*)&sourceParent, sourceStart, sourceEnd, (void*)&destinationParent, destinationRow);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStandardItemModel);
	__conns_QStandardItemModel[__id] = __c;
}
void QStandardItemModel_connect_columnsAboutToBeMoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStandardItemModel*)self, &QStandardItemModel::columnsAboutToBeMoved, (QStandardItemModel*)self, [__id](const QModelIndex& sourceParent, int sourceStart, int sourceEnd, const QModelIndex& destinationParent, int destinationColumn){
		if (cppgo_sig_void_voidp_int_int_voidp_int) ((void(*)(long long, void*, int, int, void*, int))cppgo_sig_void_voidp_int_int_voidp_int)(__id, (void*)&sourceParent, sourceStart, sourceEnd, (void*)&destinationParent, destinationColumn);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStandardItemModel);
	__conns_QStandardItemModel[__id] = __c;
}
void QStandardItemModel_connect_columnsMoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStandardItemModel*)self, &QStandardItemModel::columnsMoved, (QStandardItemModel*)self, [__id](const QModelIndex& sourceParent, int sourceStart, int sourceEnd, const QModelIndex& destinationParent, int destinationColumn){
		if (cppgo_sig_void_voidp_int_int_voidp_int) ((void(*)(long long, void*, int, int, void*, int))cppgo_sig_void_voidp_int_int_voidp_int)(__id, (void*)&sourceParent, sourceStart, sourceEnd, (void*)&destinationParent, destinationColumn);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStandardItemModel);
	__conns_QStandardItemModel[__id] = __c;
}
void QStandardItemModel_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStandardItemModel*)self, &QStandardItemModel::destroyed, (QStandardItemModel*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStandardItemModel);
	__conns_QStandardItemModel[__id] = __c;
}
void QStandardItemModel_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStandardItemModel*)self, &QStandardItemModel::destroyed, (QStandardItemModel*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStandardItemModel);
	__conns_QStandardItemModel[__id] = __c;
}
void QStandardItemModel_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStandardItemModel*)self, &QStandardItemModel::objectNameChanged, (QStandardItemModel*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStandardItemModel);
	__conns_QStandardItemModel[__id] = __c;
}
void QStandardItemModel__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QStandardItemModel);
		auto __it = __conns_QStandardItemModel.find(id);
		if (__it == __conns_QStandardItemModel.end()) return;
		__c = __it->second;
		__conns_QStandardItemModel.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QStandardItemModel__Go_new(long long id, void* parent) { return new QStandardItemModel__Go(id, (QObject*)parent); }
void QStandardItemModel_delete(void* self) { delete (QStandardItemModel*)self; }

}

#include <QtWidgets/qsizepolicy.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qvariant.h>
#include "gen_qsizepolicy.h"

extern "C" {

void* QSizePolicy_new() {
	return new QSizePolicy();
}
void* QSizePolicy_new2(int horizontal, int vertical) {
	return new QSizePolicy((QSizePolicy::Policy)(horizontal), (QSizePolicy::Policy)(vertical));
}
void* QSizePolicy_new3(void* arg1) {
	return new QSizePolicy(*(QSizePolicy*)arg1);
}
void* QSizePolicy_new4(int horizontal, int vertical, int type_) {
	return new QSizePolicy((QSizePolicy::Policy)(horizontal), (QSizePolicy::Policy)(vertical), (QSizePolicy::ControlType)(type_));
}
int QSizePolicy_horizontalPolicy(void* self) {
	auto&& __r = ((QSizePolicy*)self)->horizontalPolicy();
	return (int)(__r);
}
int QSizePolicy_verticalPolicy(void* self) {
	auto&& __r = ((QSizePolicy*)self)->verticalPolicy();
	return (int)(__r);
}
int QSizePolicy_controlType(void* self) {
	auto&& __r = ((QSizePolicy*)self)->controlType();
	return (int)(__r);
}
void QSizePolicy_setHorizontalPolicy(void* self, int d) {
	((QSizePolicy*)self)->setHorizontalPolicy((QSizePolicy::Policy)(d));
}
void QSizePolicy_setVerticalPolicy(void* self, int d) {
	((QSizePolicy*)self)->setVerticalPolicy((QSizePolicy::Policy)(d));
}
void QSizePolicy_setControlType(void* self, int type_) {
	((QSizePolicy*)self)->setControlType((QSizePolicy::ControlType)(type_));
}
int QSizePolicy_expandingDirections(void* self) {
	auto&& __r = ((QSizePolicy*)self)->expandingDirections();
	return (int)(__r);
}
void QSizePolicy_setHeightForWidth(void* self, int b) {
	((QSizePolicy*)self)->setHeightForWidth(b);
}
int QSizePolicy_hasHeightForWidth(void* self) {
	auto&& __r = ((QSizePolicy*)self)->hasHeightForWidth();
	return __r;
}
void QSizePolicy_setWidthForHeight(void* self, int b) {
	((QSizePolicy*)self)->setWidthForHeight(b);
}
int QSizePolicy_hasWidthForHeight(void* self) {
	auto&& __r = ((QSizePolicy*)self)->hasWidthForHeight();
	return __r;
}
int QSizePolicy_OperatorEqual(void* self, void* s) {
	auto&& __r = ((QSizePolicy*)self)->operator==(*(QSizePolicy*)s);
	return __r;
}
int QSizePolicy_OperatorNotEqual(void* self, void* s) {
	auto&& __r = ((QSizePolicy*)self)->operator!=(*(QSizePolicy*)s);
	return __r;
}
void* QSizePolicy_ToQVariant(void* self) {
	auto&& __r = static_cast<QVariant>(*(QSizePolicy*)self);
	return new QVariant(__r);
}
int QSizePolicy_horizontalStretch(void* self) {
	auto&& __r = ((QSizePolicy*)self)->horizontalStretch();
	return __r;
}
int QSizePolicy_verticalStretch(void* self) {
	auto&& __r = ((QSizePolicy*)self)->verticalStretch();
	return __r;
}
void QSizePolicy_setHorizontalStretch(void* self, int stretchFactor) {
	((QSizePolicy*)self)->setHorizontalStretch(stretchFactor);
}
void QSizePolicy_setVerticalStretch(void* self, int stretchFactor) {
	((QSizePolicy*)self)->setVerticalStretch(stretchFactor);
}
int QSizePolicy_retainSizeWhenHidden(void* self) {
	auto&& __r = ((QSizePolicy*)self)->retainSizeWhenHidden();
	return __r;
}
void QSizePolicy_setRetainSizeWhenHidden(void* self, int retainSize) {
	((QSizePolicy*)self)->setRetainSizeWhenHidden(retainSize);
}
void QSizePolicy_transpose(void* self) {
	((QSizePolicy*)self)->transpose();
}
void* QSizePolicy_transposed(void* self) {
	auto&& __r = ((QSizePolicy*)self)->transposed();
	return new QSizePolicy(__r);
}
void* QSizePolicy_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QSizePolicy*)self)->operator=(*(QSizePolicy*)arg1);
	return (void*)&(__r);
}
void* QSizePolicy_toQVariant(void* self) {
	auto&& __r = QVariant::fromValue(*(QSizePolicy*)self);
	return new QVariant(__r);
}
void QSizePolicy_delete(void* self) { delete (QSizePolicy*)self; }

}

#include <QtWidgets/qstyleoption.h>
#include <QtWidgets/qwidget.h>
#include "gen_qstyleoptionmenuitemv2.h"

extern "C" {

void* QStyleOptionMenuItemV2_base(void* self) { return (void*)static_cast<QStyleOptionMenuItem*>((QStyleOptionMenuItemV2*)self); }
void QStyleOptionMenuItemV2_chain(void* self, void** out) {
	QStyleOptionMenuItemV2* __p = (QStyleOptionMenuItemV2*)self;
	QStyleOptionMenuItem* __b0 = static_cast<QStyleOptionMenuItem*>(__p); out[0] = (void*)__b0;
	QStyleOption* __b1 = static_cast<QStyleOption*>(__b0); out[1] = (void*)__b1;
}
void* QStyleOptionMenuItemV2_new() {
	return new QStyleOptionMenuItemV2();
}
void* QStyleOptionMenuItemV2_new2(void* other) {
	return new QStyleOptionMenuItemV2(*(QStyleOptionMenuItemV2*)other);
}
void* QStyleOptionMenuItemV2_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QStyleOptionMenuItemV2*)self)->operator=(*(QStyleOptionMenuItemV2*)arg1);
	return (void*)&(__r);
}
int QStyleOptionMenuItemV2_mouseDown_get(void* self) {
	auto __r = ((QStyleOptionMenuItemV2*)self)->mouseDown;
	return __r;
}
void QStyleOptionMenuItemV2_mouseDown_set(void* self, int mouseDown) {
	((QStyleOptionMenuItemV2*)self)->mouseDown = mouseDown;
}
int QStyleOptionMenuItemV2_unused_get(void* self) {
	auto __r = ((QStyleOptionMenuItemV2*)self)->unused;
	return __r;
}
void QStyleOptionMenuItemV2_unused_set(void* self, int unused) {
	((QStyleOptionMenuItemV2*)self)->unused = unused;
}
void QStyleOptionMenuItemV2_delete(void* self) { delete (QStyleOptionMenuItemV2*)self; }

}

#include <QtWidgets/qlayoutitem.h>
#include <QtWidgets/qsizepolicy.h>
#include <QtCore/qsize.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qrect.h>
#include <QtWidgets/qwidget.h>
#include <QtWidgets/qlayout.h>
#include "gen_qspaceritem.h"

extern "C" {

void* QSpacerItem_base(void* self) { return (void*)static_cast<QLayoutItem*>((QSpacerItem*)self); }
void* QSpacerItem_new(int w, int h) {
	return new QSpacerItem(w, h);
}
void* QSpacerItem_new2(void* arg1) {
	return new QSpacerItem(*(QSpacerItem*)arg1);
}
void* QSpacerItem_new3(int w, int h, int hData) {
	return new QSpacerItem(w, h, (QSizePolicy::Policy)(hData));
}
void* QSpacerItem_new4(int w, int h, int hData, int vData) {
	return new QSpacerItem(w, h, (QSizePolicy::Policy)(hData), (QSizePolicy::Policy)(vData));
}
void QSpacerItem_changeSize(void* self, int w, int h) {
	((QSpacerItem*)self)->changeSize(w, h);
}
void* QSpacerItem_sizePolicy(void* self) {
	auto&& __r = ((QSpacerItem*)self)->sizePolicy();
	return new QSizePolicy(__r);
}
void* QSpacerItem_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QSpacerItem*)self)->operator=(*(QSpacerItem*)arg1);
	return (void*)&(__r);
}
void QSpacerItem_changeSize2(void* self, int w, int h, int hData) {
	((QSpacerItem*)self)->changeSize(w, h, (QSizePolicy::Policy)(hData));
}
void QSpacerItem_changeSize3(void* self, int w, int h, int hData, int vData) {
	((QSpacerItem*)self)->changeSize(w, h, (QSizePolicy::Policy)(hData), (QSizePolicy::Policy)(vData));
}
void QSpacerItem_delete(void* self) { delete (QSpacerItem*)self; }

}

#include <QtWidgets/qdial.h>
#include <QtWidgets/qabstractslider.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qsize.h>
#include <QtCore/qcoreevent.h>
#include <QtGui/qevent.h>
#include <QtWidgets/qstyleoption.h>
#include <QtCore/qnamespace.h>
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
#include <QtCore/qvariant.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qdial.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QDial__Go : public QDial {
public:
	long long __goId;
	unsigned char __goOn[50] = {};
	template<typename... A> QDial__Go(long long id, A&&... a) : QDial(std::forward<A>(a)...), __goId(id) {}
	using QDial::setRepeatAction;
	using QDial::repeatAction;
	using QDial::updateMicroFocus;
	using QDial::create;
	using QDial::destroy;
	using QDial::focusNextChild;
	using QDial::focusPreviousChild;
	using QDial::sender;
	using QDial::senderSignalIndex;
	using QDial::receivers;
	using QDial::isSignalConnected;
	using QDial::getDecodedMetricF;
	bool cppgoBase_event(QEvent* __a0) { return QDial::event(__a0); }
	void cppgoBase_resizeEvent(QResizeEvent* __a0) { QDial::resizeEvent(__a0); }
	void cppgoBase_paintEvent(QPaintEvent* __a0) { QDial::paintEvent(__a0); }
	void cppgoBase_mousePressEvent(QMouseEvent* __a0) { QDial::mousePressEvent(__a0); }
	void cppgoBase_mouseReleaseEvent(QMouseEvent* __a0) { QDial::mouseReleaseEvent(__a0); }
	void cppgoBase_mouseMoveEvent(QMouseEvent* __a0) { QDial::mouseMoveEvent(__a0); }
	void cppgoBase_sliderChange(int __a0) { QDial::sliderChange((QAbstractSlider::SliderChange)__a0); }
	void cppgoBase_initStyleOption(QStyleOptionSlider* __a0) { QDial::initStyleOption(__a0); }
	void cppgoBase_keyPressEvent(QKeyEvent* __a0) { QDial::keyPressEvent(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QDial::timerEvent(__a0); }
	void cppgoBase_wheelEvent(QWheelEvent* __a0) { QDial::wheelEvent(__a0); }
	void cppgoBase_changeEvent(QEvent* __a0) { QDial::changeEvent(__a0); }
	void cppgoBase_mouseDoubleClickEvent(QMouseEvent* __a0) { QDial::mouseDoubleClickEvent(__a0); }
	void cppgoBase_keyReleaseEvent(QKeyEvent* __a0) { QDial::keyReleaseEvent(__a0); }
	void cppgoBase_focusInEvent(QFocusEvent* __a0) { QDial::focusInEvent(__a0); }
	void cppgoBase_focusOutEvent(QFocusEvent* __a0) { QDial::focusOutEvent(__a0); }
	void cppgoBase_enterEvent(QEnterEvent* __a0) { QDial::enterEvent(__a0); }
	void cppgoBase_leaveEvent(QEvent* __a0) { QDial::leaveEvent(__a0); }
	void cppgoBase_moveEvent(QMoveEvent* __a0) { QDial::moveEvent(__a0); }
	void cppgoBase_closeEvent(QCloseEvent* __a0) { QDial::closeEvent(__a0); }
	void cppgoBase_contextMenuEvent(QContextMenuEvent* __a0) { QDial::contextMenuEvent(__a0); }
	void cppgoBase_tabletEvent(QTabletEvent* __a0) { QDial::tabletEvent(__a0); }
	void cppgoBase_actionEvent(QActionEvent* __a0) { QDial::actionEvent(__a0); }
	void cppgoBase_dragEnterEvent(QDragEnterEvent* __a0) { QDial::dragEnterEvent(__a0); }
	void cppgoBase_dragMoveEvent(QDragMoveEvent* __a0) { QDial::dragMoveEvent(__a0); }
	void cppgoBase_dragLeaveEvent(QDragLeaveEvent* __a0) { QDial::dragLeaveEvent(__a0); }
	void cppgoBase_dropEvent(QDropEvent* __a0) { QDial::dropEvent(__a0); }
	void cppgoBase_showEvent(QShowEvent* __a0) { QDial::showEvent(__a0); }
	void cppgoBase_hideEvent(QHideEvent* __a0) { QDial::hideEvent(__a0); }
	int cppgoBase_metric(int __a0) { return QDial::metric((QPaintDevice::PaintDeviceMetric)__a0); }
	void cppgoBase_initPainter(QPainter* __a0) { QDial::initPainter(__a0); }
	QPaintDevice* cppgoBase_redirected(QPoint* __a0) { return QDial::redirected(__a0); }
	QPainter* cppgoBase_sharedPainter() { return QDial::sharedPainter(); }
	void cppgoBase_inputMethodEvent(QInputMethodEvent* __a0) { QDial::inputMethodEvent(__a0); }
	bool cppgoBase_focusNextPrevChild(bool __a0) { return QDial::focusNextPrevChild(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QDial::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QDial::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QDial::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QDial::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QDial::metaObject();
	}
	QSize sizeHint() const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 729690261457684101LL, &__h); if (__h) return *(QSize*)__r; }
		return QDial::sizeHint();
	}
	QSize minimumSizeHint() const override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 7049264599837532491LL, &__h); if (__h) return *(QSize*)__r; }
		return QDial::minimumSizeHint();
	}
	bool event(QEvent* e) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)e, &__h); if (__h) return __r; }
		return QDial::event(e);
	}
	void resizeEvent(QResizeEvent* re) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4302917285655014661LL, (void*)re, &__h); if (__h) return; }
		QDial::resizeEvent(re);
	}
	void paintEvent(QPaintEvent* pe) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5290442163167458465LL, (void*)pe, &__h); if (__h) return; }
		QDial::paintEvent(pe);
	}
	void mousePressEvent(QMouseEvent* me) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -332090898078543319LL, (void*)me, &__h); if (__h) return; }
		QDial::mousePressEvent(me);
	}
	void mouseReleaseEvent(QMouseEvent* me) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4415920049297795213LL, (void*)me, &__h); if (__h) return; }
		QDial::mouseReleaseEvent(me);
	}
	void mouseMoveEvent(QMouseEvent* me) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4255494004099806581LL, (void*)me, &__h); if (__h) return; }
		QDial::mouseMoveEvent(me);
	}
	void sliderChange(QAbstractSlider::SliderChange change) override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 4658628703300230988LL, (int)(change), &__h); if (__h) return; }
		QDial::sliderChange(change);
	}
	void initStyleOption(QStyleOptionSlider* option) const override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -488693769771308531LL, (void*)option, &__h); if (__h) return; }
		QDial::initStyleOption(option);
	}
	void keyPressEvent(QKeyEvent* ev) override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1023791977679793185LL, (void*)ev, &__h); if (__h) return; }
		QDial::keyPressEvent(ev);
	}
	void timerEvent(QTimerEvent* arg1) override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)arg1, &__h); if (__h) return; }
		QDial::timerEvent(arg1);
	}
	void wheelEvent(QWheelEvent* e) override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1919988340934511316LL, (void*)e, &__h); if (__h) return; }
		QDial::wheelEvent(e);
	}
	void changeEvent(QEvent* e) override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4959376063826165943LL, (void*)e, &__h); if (__h) return; }
		QDial::changeEvent(e);
	}
	int devType() const override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8659679893868524044LL, &__h); if (__h) return __r; }
		return QDial::devType();
	}
	void setVisible(bool visible) override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 2110182719375245575LL, visible, &__h); if (__h) return; }
		QDial::setVisible(visible);
	}
	int heightForWidth(int arg1) const override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 4311534060495737339LL, arg1, &__h); if (__h) return __r; }
		return QDial::heightForWidth(arg1);
	}
	bool hasHeightForWidth() const override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8266762397995423915LL, &__h); if (__h) return __r; }
		return QDial::hasHeightForWidth();
	}
	QPaintEngine* paintEngine() const override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 6398007404756346381LL, &__h); if (__h) return (QPaintEngine*)__r; }
		return QDial::paintEngine();
	}
	void mouseDoubleClickEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4029073905697023887LL, (void*)event, &__h); if (__h) return; }
		QDial::mouseDoubleClickEvent(event);
	}
	void keyReleaseEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 772598475950050299LL, (void*)event, &__h); if (__h) return; }
		QDial::keyReleaseEvent(event);
	}
	void focusInEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 975043767033552552LL, (void*)event, &__h); if (__h) return; }
		QDial::focusInEvent(event);
	}
	void focusOutEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -385596010576048993LL, (void*)event, &__h); if (__h) return; }
		QDial::focusOutEvent(event);
	}
	void enterEvent(QEnterEvent* event) override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6316270346668237895LL, (void*)event, &__h); if (__h) return; }
		QDial::enterEvent(event);
	}
	void leaveEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -551070009623711640LL, (void*)event, &__h); if (__h) return; }
		QDial::leaveEvent(event);
	}
	void moveEvent(QMoveEvent* event) override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -9161949150489525396LL, (void*)event, &__h); if (__h) return; }
		QDial::moveEvent(event);
	}
	void closeEvent(QCloseEvent* event) override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1989937050019464783LL, (void*)event, &__h); if (__h) return; }
		QDial::closeEvent(event);
	}
	void contextMenuEvent(QContextMenuEvent* event) override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6698043222227006015LL, (void*)event, &__h); if (__h) return; }
		QDial::contextMenuEvent(event);
	}
	void tabletEvent(QTabletEvent* event) override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -212068866860701141LL, (void*)event, &__h); if (__h) return; }
		QDial::tabletEvent(event);
	}
	void actionEvent(QActionEvent* event) override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1697444140088555581LL, (void*)event, &__h); if (__h) return; }
		QDial::actionEvent(event);
	}
	void dragEnterEvent(QDragEnterEvent* event) override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 206122715919917589LL, (void*)event, &__h); if (__h) return; }
		QDial::dragEnterEvent(event);
	}
	void dragMoveEvent(QDragMoveEvent* event) override {
		int __h = 0;
		if (__goOn[32] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1781545137352745978LL, (void*)event, &__h); if (__h) return; }
		QDial::dragMoveEvent(event);
	}
	void dragLeaveEvent(QDragLeaveEvent* event) override {
		int __h = 0;
		if (__goOn[33] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -3139904158373725006LL, (void*)event, &__h); if (__h) return; }
		QDial::dragLeaveEvent(event);
	}
	void dropEvent(QDropEvent* event) override {
		int __h = 0;
		if (__goOn[34] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 9107058682754591716LL, (void*)event, &__h); if (__h) return; }
		QDial::dropEvent(event);
	}
	void showEvent(QShowEvent* event) override {
		int __h = 0;
		if (__goOn[35] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6825369174751916LL, (void*)event, &__h); if (__h) return; }
		QDial::showEvent(event);
	}
	void hideEvent(QHideEvent* event) override {
		int __h = 0;
		if (__goOn[36] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -199718707876719129LL, (void*)event, &__h); if (__h) return; }
		QDial::hideEvent(event);
	}
	bool nativeEvent(const QByteArray& eventType, void* message, qintptr* result) override {
		int __h = 0;
		if (__goOn[37] && cppgo_virt_int_voidp_voidp_longlongp) { int __r = ((int(*)(long long, long long, void*, void*, long long*, int*))cppgo_virt_int_voidp_voidp_longlongp)(__goId, -2039166586246443706LL, (void*)&eventType, (void*)message, (long long*)result, &__h); if (__h) return __r; }
		return QDial::nativeEvent(eventType, message, result);
	}
	int metric(QPaintDevice::PaintDeviceMetric arg1) const override {
		int __h = 0;
		if (__goOn[38] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 427629649267311777LL, (int)(arg1), &__h); if (__h) return __r; }
		return QDial::metric(arg1);
	}
	void initPainter(QPainter* painter) const override {
		int __h = 0;
		if (__goOn[39] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6477500296980050210LL, (void*)painter, &__h); if (__h) return; }
		QDial::initPainter(painter);
	}
	QPaintDevice* redirected(QPoint* offset) const override {
		int __h = 0;
		if (__goOn[40] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 4127240633310655048LL, (void*)offset, &__h); if (__h) return (QPaintDevice*)__r; }
		return QDial::redirected(offset);
	}
	QPainter* sharedPainter() const override {
		int __h = 0;
		if (__goOn[41] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 5792771328741774991LL, &__h); if (__h) return (QPainter*)__r; }
		return QDial::sharedPainter();
	}
	void inputMethodEvent(QInputMethodEvent* arg1) override {
		int __h = 0;
		if (__goOn[42] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5438819661942843222LL, (void*)arg1, &__h); if (__h) return; }
		QDial::inputMethodEvent(arg1);
	}
	QVariant inputMethodQuery(Qt::InputMethodQuery arg1) const override {
		int __h = 0;
		if (__goOn[43] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7345853097340709638LL, (int)(arg1), &__h); if (__h) return *(QVariant*)__r; }
		return QDial::inputMethodQuery(arg1);
	}
	bool focusNextPrevChild(bool next) override {
		int __h = 0;
		if (__goOn[44] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 5898328517505784191LL, next, &__h); if (__h) return __r; }
		return QDial::focusNextPrevChild(next);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[45] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QDial::eventFilter(watched, event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[46] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QDial::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[47] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QDial::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[48] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QDial::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[49] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QDial::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QDial;
static std::mutex __connsMu_QDial;

extern "C" {

void* QDial_base(void* self) { return (void*)static_cast<QAbstractSlider*>((QDial*)self); }
void QDial_chain(void* self, void** out) {
	QDial* __p = (QDial*)self;
	QAbstractSlider* __b0 = static_cast<QAbstractSlider*>(__p); out[0] = (void*)__b0;
	QWidget* __b1 = static_cast<QWidget*>(__b0); out[1] = (void*)__b1;
	QObject* __b2 = static_cast<QObject*>(__b1); out[2] = (void*)__b2;
}
void QDial__set_override(void* self, int idx) { ((QDial__Go*)self)->__goOn[idx] = 1; }
void* QDial_fromQObject(void* o) { return (void*)qobject_cast<QDial*>((QObject*)o); }
void* QDial_new(long long __goid, void* parent) {
	return new QDial__Go(__goid, (QWidget*)parent);
}
void* QDial_new2(long long __goid) {
	return new QDial__Go(__goid);
}
char* QDial_tr(const char* s, long long* __out_len) {
	auto&& __r = QDial::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QDial_wrapping(void* self) {
	auto&& __r = ((QDial*)self)->wrapping();
	return __r;
}
int QDial_notchSize(void* self) {
	auto&& __r = ((QDial*)self)->notchSize();
	return __r;
}
void QDial_setNotchTarget(void* self, double target) {
	((QDial*)self)->setNotchTarget(target);
}
double QDial_notchTarget(void* self) {
	auto&& __r = ((QDial*)self)->notchTarget();
	return __r;
}
int QDial_notchesVisible(void* self) {
	auto&& __r = ((QDial*)self)->notchesVisible();
	return __r;
}
void QDial_setNotchesVisible(void* self, int visible) {
	((QDial*)self)->setNotchesVisible(visible);
}
void QDial_setWrapping(void* self, int on) {
	((QDial*)self)->setWrapping(on);
}
void QDial_setRepeatAction(void* self, int action) {
	((QDial__Go*)self)->setRepeatAction((QAbstractSlider::SliderAction)(action));
}
int QDial_repeatAction(void* self) {
	auto&& __r = ((QDial__Go*)self)->repeatAction();
	return (int)(__r);
}
void QDial_setTabOrder(void* arg1, void* arg2) {
	QDial::setTabOrder((QWidget*)arg1, (QWidget*)arg2);
}
void* QDial_mouseGrabber() {
	auto&& __r = QDial::mouseGrabber();
	return (void*)(__r);
}
void* QDial_keyboardGrabber() {
	auto&& __r = QDial::keyboardGrabber();
	return (void*)(__r);
}
void* QDial_find(unsigned long long arg1) {
	auto&& __r = QDial::find(arg1);
	return (void*)(__r);
}
void* QDial_createWindowContainer(void* window) {
	auto&& __r = QDial::createWindowContainer((QWindow*)window);
	return (void*)(__r);
}
void QDial_updateMicroFocus(void* self) {
	((QDial__Go*)self)->updateMicroFocus();
}
void QDial_create(void* self) {
	((QDial__Go*)self)->create();
}
void QDial_destroy(void* self) {
	((QDial__Go*)self)->destroy();
}
int QDial_focusNextChild(void* self) {
	auto&& __r = ((QDial__Go*)self)->focusNextChild();
	return __r;
}
int QDial_focusPreviousChild(void* self) {
	auto&& __r = ((QDial__Go*)self)->focusPreviousChild();
	return __r;
}
void* QDial_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QDial::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QDial_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QDial::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QDial_sender(void* self) {
	auto&& __r = ((QDial__Go*)self)->sender();
	return (void*)(__r);
}
int QDial_senderSignalIndex(void* self) {
	auto&& __r = ((QDial__Go*)self)->senderSignalIndex();
	return __r;
}
int QDial_receivers(void* self, const char* signal) {
	auto&& __r = ((QDial__Go*)self)->receivers(signal);
	return __r;
}
int QDial_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QDial__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
double QDial_devicePixelRatioFScale() {
	auto&& __r = QDial::devicePixelRatioFScale();
	return __r;
}
int QDial_encodeMetricF(int metric, double value) {
	auto&& __r = QDial::encodeMetricF((QPaintDevice::PaintDeviceMetric)(metric), value);
	return __r;
}
double QDial_getDecodedMetricF(void* self, int metricA, int metricB) {
	auto&& __r = ((QDial__Go*)self)->getDecodedMetricF((QPaintDevice::PaintDeviceMetric)(metricA), (QPaintDevice::PaintDeviceMetric)(metricB));
	return __r;
}
char* QDial_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QDial::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QDial_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QDial::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QDial_setRepeatAction2(void* self, int action, int thresholdTime) {
	((QDial__Go*)self)->setRepeatAction((QAbstractSlider::SliderAction)(action), thresholdTime);
}
void QDial_setRepeatAction3(void* self, int action, int thresholdTime, int repeatTime) {
	((QDial__Go*)self)->setRepeatAction((QAbstractSlider::SliderAction)(action), thresholdTime, repeatTime);
}
void* QDial_createWindowContainer2(void* window, void* parent) {
	auto&& __r = QDial::createWindowContainer((QWindow*)window, (QWidget*)parent);
	return (void*)(__r);
}
void* QDial_createWindowContainer3(void* window, void* parent, int flags) {
	auto&& __r = QDial::createWindowContainer((QWindow*)window, (QWidget*)parent, QFlags<Qt::WindowType>(flags));
	return (void*)(__r);
}
void QDial_updateMicroFocus2(void* self, int query) {
	((QDial__Go*)self)->updateMicroFocus((Qt::InputMethodQuery)(query));
}
void QDial_create2(void* self, unsigned long long arg1) {
	((QDial__Go*)self)->create(arg1);
}
void QDial_create3(void* self, unsigned long long arg1, int initializeWindow) {
	((QDial__Go*)self)->create(arg1, initializeWindow);
}
void QDial_create4(void* self, unsigned long long arg1, int initializeWindow, int destroyOldWindow) {
	((QDial__Go*)self)->create(arg1, initializeWindow, destroyOldWindow);
}
void QDial_destroy2(void* self, int destroyWindow) {
	((QDial__Go*)self)->destroy(destroyWindow);
}
void QDial_destroy3(void* self, int destroyWindow, int destroySubWindows) {
	((QDial__Go*)self)->destroy(destroyWindow, destroySubWindows);
}
void* QDial_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QDial::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QDial_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QDial::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QDial_metaObject_default(void* self) {
	auto&& __r = ((QDial*)self)->QDial::metaObject();
	return (void*)(__r);
}
void* QDial_sizeHint_default(void* self) {
	auto&& __r = ((QDial*)self)->QDial::sizeHint();
	return new QSize(__r);
}
void* QDial_minimumSizeHint_default(void* self) {
	auto&& __r = ((QDial*)self)->QDial::minimumSizeHint();
	return new QSize(__r);
}
int QDial_event_default(void* self, void* e) {
	auto&& __r = ((QDial__Go*)self)->cppgoBase_event((QEvent*)e);
	return __r;
}
void QDial_resizeEvent_default(void* self, void* re) {
	((QDial__Go*)self)->cppgoBase_resizeEvent((QResizeEvent*)re);
}
void QDial_paintEvent_default(void* self, void* pe) {
	((QDial__Go*)self)->cppgoBase_paintEvent((QPaintEvent*)pe);
}
void QDial_mousePressEvent_default(void* self, void* me) {
	((QDial__Go*)self)->cppgoBase_mousePressEvent((QMouseEvent*)me);
}
void QDial_mouseReleaseEvent_default(void* self, void* me) {
	((QDial__Go*)self)->cppgoBase_mouseReleaseEvent((QMouseEvent*)me);
}
void QDial_mouseMoveEvent_default(void* self, void* me) {
	((QDial__Go*)self)->cppgoBase_mouseMoveEvent((QMouseEvent*)me);
}
void QDial_sliderChange_default(void* self, int change) {
	((QDial__Go*)self)->cppgoBase_sliderChange(change);
}
void QDial_initStyleOption_default(void* self, void* option) {
	((QDial__Go*)self)->cppgoBase_initStyleOption((QStyleOptionSlider*)option);
}
void QDial_keyPressEvent_default(void* self, void* ev) {
	((QDial__Go*)self)->cppgoBase_keyPressEvent((QKeyEvent*)ev);
}
void QDial_timerEvent_default(void* self, void* arg1) {
	((QDial__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)arg1);
}
void QDial_wheelEvent_default(void* self, void* e) {
	((QDial__Go*)self)->cppgoBase_wheelEvent((QWheelEvent*)e);
}
void QDial_changeEvent_default(void* self, void* e) {
	((QDial__Go*)self)->cppgoBase_changeEvent((QEvent*)e);
}
int QDial_devType_default(void* self) {
	auto&& __r = ((QDial*)self)->QDial::devType();
	return __r;
}
void QDial_setVisible_default(void* self, int visible) {
	((QDial*)self)->QDial::setVisible(visible);
}
int QDial_heightForWidth_default(void* self, int arg1) {
	auto&& __r = ((QDial*)self)->QDial::heightForWidth(arg1);
	return __r;
}
int QDial_hasHeightForWidth_default(void* self) {
	auto&& __r = ((QDial*)self)->QDial::hasHeightForWidth();
	return __r;
}
void* QDial_paintEngine_default(void* self) {
	auto&& __r = ((QDial*)self)->QDial::paintEngine();
	return (void*)(__r);
}
void QDial_mouseDoubleClickEvent_default(void* self, void* event) {
	((QDial__Go*)self)->cppgoBase_mouseDoubleClickEvent((QMouseEvent*)event);
}
void QDial_keyReleaseEvent_default(void* self, void* event) {
	((QDial__Go*)self)->cppgoBase_keyReleaseEvent((QKeyEvent*)event);
}
void QDial_focusInEvent_default(void* self, void* event) {
	((QDial__Go*)self)->cppgoBase_focusInEvent((QFocusEvent*)event);
}
void QDial_focusOutEvent_default(void* self, void* event) {
	((QDial__Go*)self)->cppgoBase_focusOutEvent((QFocusEvent*)event);
}
void QDial_enterEvent_default(void* self, void* event) {
	((QDial__Go*)self)->cppgoBase_enterEvent((QEnterEvent*)event);
}
void QDial_leaveEvent_default(void* self, void* event) {
	((QDial__Go*)self)->cppgoBase_leaveEvent((QEvent*)event);
}
void QDial_moveEvent_default(void* self, void* event) {
	((QDial__Go*)self)->cppgoBase_moveEvent((QMoveEvent*)event);
}
void QDial_closeEvent_default(void* self, void* event) {
	((QDial__Go*)self)->cppgoBase_closeEvent((QCloseEvent*)event);
}
void QDial_contextMenuEvent_default(void* self, void* event) {
	((QDial__Go*)self)->cppgoBase_contextMenuEvent((QContextMenuEvent*)event);
}
void QDial_tabletEvent_default(void* self, void* event) {
	((QDial__Go*)self)->cppgoBase_tabletEvent((QTabletEvent*)event);
}
void QDial_actionEvent_default(void* self, void* event) {
	((QDial__Go*)self)->cppgoBase_actionEvent((QActionEvent*)event);
}
void QDial_dragEnterEvent_default(void* self, void* event) {
	((QDial__Go*)self)->cppgoBase_dragEnterEvent((QDragEnterEvent*)event);
}
void QDial_dragMoveEvent_default(void* self, void* event) {
	((QDial__Go*)self)->cppgoBase_dragMoveEvent((QDragMoveEvent*)event);
}
void QDial_dragLeaveEvent_default(void* self, void* event) {
	((QDial__Go*)self)->cppgoBase_dragLeaveEvent((QDragLeaveEvent*)event);
}
void QDial_dropEvent_default(void* self, void* event) {
	((QDial__Go*)self)->cppgoBase_dropEvent((QDropEvent*)event);
}
void QDial_showEvent_default(void* self, void* event) {
	((QDial__Go*)self)->cppgoBase_showEvent((QShowEvent*)event);
}
void QDial_hideEvent_default(void* self, void* event) {
	((QDial__Go*)self)->cppgoBase_hideEvent((QHideEvent*)event);
}
int QDial_metric_default(void* self, int arg1) {
	auto&& __r = ((QDial__Go*)self)->cppgoBase_metric(arg1);
	return __r;
}
void QDial_initPainter_default(void* self, void* painter) {
	((QDial__Go*)self)->cppgoBase_initPainter((QPainter*)painter);
}
void* QDial_redirected_default(void* self, void* offset) {
	auto&& __r = ((QDial__Go*)self)->cppgoBase_redirected((QPoint*)offset);
	return (void*)(__r);
}
void* QDial_sharedPainter_default(void* self) {
	auto&& __r = ((QDial__Go*)self)->cppgoBase_sharedPainter();
	return (void*)(__r);
}
void QDial_inputMethodEvent_default(void* self, void* arg1) {
	((QDial__Go*)self)->cppgoBase_inputMethodEvent((QInputMethodEvent*)arg1);
}
void* QDial_inputMethodQuery_default(void* self, int arg1) {
	auto&& __r = ((QDial*)self)->QDial::inputMethodQuery((Qt::InputMethodQuery)(arg1));
	return new QVariant(__r);
}
int QDial_focusNextPrevChild_default(void* self, int next) {
	auto&& __r = ((QDial__Go*)self)->cppgoBase_focusNextPrevChild(next);
	return __r;
}
int QDial_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QDial*)self)->QDial::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QDial_childEvent_default(void* self, void* event) {
	((QDial__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QDial_customEvent_default(void* self, void* event) {
	((QDial__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QDial_connectNotify_default(void* self, void* signal) {
	((QDial__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QDial_disconnectNotify_default(void* self, void* signal) {
	((QDial__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QDial_connect_valueChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDial*)self, &QDial::valueChanged, (QDial*)self, [__id](int value){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, value);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDial);
	__conns_QDial[__id] = __c;
}
void QDial_connect_sliderPressed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDial*)self, &QDial::sliderPressed, (QDial*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDial);
	__conns_QDial[__id] = __c;
}
void QDial_connect_sliderMoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDial*)self, &QDial::sliderMoved, (QDial*)self, [__id](int position){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, position);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDial);
	__conns_QDial[__id] = __c;
}
void QDial_connect_sliderReleased(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDial*)self, &QDial::sliderReleased, (QDial*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDial);
	__conns_QDial[__id] = __c;
}
void QDial_connect_rangeChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDial*)self, &QDial::rangeChanged, (QDial*)self, [__id](int min, int max){
		if (cppgo_sig_void_int_int) ((void(*)(long long, int, int))cppgo_sig_void_int_int)(__id, min, max);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDial);
	__conns_QDial[__id] = __c;
}
void QDial_connect_actionTriggered(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDial*)self, &QDial::actionTriggered, (QDial*)self, [__id](int action){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, action);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDial);
	__conns_QDial[__id] = __c;
}
void QDial_connect_windowTitleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDial*)self, &QDial::windowTitleChanged, (QDial*)self, [__id](const QString& title){
		QByteArray title_ba = title.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)title_ba.constData(), title_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDial);
	__conns_QDial[__id] = __c;
}
void QDial_connect_windowIconChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDial*)self, &QDial::windowIconChanged, (QDial*)self, [__id](const QIcon& icon){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&icon);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDial);
	__conns_QDial[__id] = __c;
}
void QDial_connect_windowIconTextChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDial*)self, &QDial::windowIconTextChanged, (QDial*)self, [__id](const QString& iconText){
		QByteArray iconText_ba = iconText.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)iconText_ba.constData(), iconText_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDial);
	__conns_QDial[__id] = __c;
}
void QDial_connect_customContextMenuRequested(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDial*)self, &QDial::customContextMenuRequested, (QDial*)self, [__id](const QPoint& pos){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&pos);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDial);
	__conns_QDial[__id] = __c;
}
void QDial_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDial*)self, &QDial::destroyed, (QDial*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDial);
	__conns_QDial[__id] = __c;
}
void QDial_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDial*)self, &QDial::destroyed, (QDial*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDial);
	__conns_QDial[__id] = __c;
}
void QDial_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDial*)self, &QDial::objectNameChanged, (QDial*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDial);
	__conns_QDial[__id] = __c;
}
void QDial__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QDial);
		auto __it = __conns_QDial.find(id);
		if (__it == __conns_QDial.end()) return;
		__c = __it->second;
		__conns_QDial.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QDial__Go_new(long long id) { return new QDial__Go(id); }
void QDial_delete(void* self) { delete (QDial*)self; }

}

#include <QtWidgets/qgraphicsitem.h>
#include <QtCore/qrect.h>
#include <QtGui/qpainterpath.h>
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
#include "gen_qgraphicsrectitem.h"
#include "gen_purego_cb.h"

class QGraphicsRectItem__Go : public QGraphicsRectItem {
public:
	long long __goId;
	unsigned char __goOn[35] = {};
	template<typename... A> QGraphicsRectItem__Go(long long id, A&&... a) : QGraphicsRectItem(std::forward<A>(a)...), __goId(id) {}
	using QGraphicsRectItem::updateMicroFocus;
	using QGraphicsRectItem::addToIndex;
	using QGraphicsRectItem::removeFromIndex;
	using QGraphicsRectItem::prepareGeometryChange;
	bool cppgoBase_supportsExtension(int __a0) { return QGraphicsRectItem::supportsExtension((QGraphicsItem::Extension)__a0); }
	void cppgoBase_setExtension(int __a0, const QVariant& __a1) { QGraphicsRectItem::setExtension((QGraphicsItem::Extension)__a0, __a1); }
	QVariant cppgoBase_extension(const QVariant& __a0) { return QGraphicsRectItem::extension(__a0); }
	bool cppgoBase_sceneEventFilter(QGraphicsItem* __a0, QEvent* __a1) { return QGraphicsRectItem::sceneEventFilter(__a0, __a1); }
	bool cppgoBase_sceneEvent(QEvent* __a0) { return QGraphicsRectItem::sceneEvent(__a0); }
	void cppgoBase_contextMenuEvent(QGraphicsSceneContextMenuEvent* __a0) { QGraphicsRectItem::contextMenuEvent(__a0); }
	void cppgoBase_dragEnterEvent(QGraphicsSceneDragDropEvent* __a0) { QGraphicsRectItem::dragEnterEvent(__a0); }
	void cppgoBase_dragLeaveEvent(QGraphicsSceneDragDropEvent* __a0) { QGraphicsRectItem::dragLeaveEvent(__a0); }
	void cppgoBase_dragMoveEvent(QGraphicsSceneDragDropEvent* __a0) { QGraphicsRectItem::dragMoveEvent(__a0); }
	void cppgoBase_dropEvent(QGraphicsSceneDragDropEvent* __a0) { QGraphicsRectItem::dropEvent(__a0); }
	void cppgoBase_focusInEvent(QFocusEvent* __a0) { QGraphicsRectItem::focusInEvent(__a0); }
	void cppgoBase_focusOutEvent(QFocusEvent* __a0) { QGraphicsRectItem::focusOutEvent(__a0); }
	void cppgoBase_hoverEnterEvent(QGraphicsSceneHoverEvent* __a0) { QGraphicsRectItem::hoverEnterEvent(__a0); }
	void cppgoBase_hoverMoveEvent(QGraphicsSceneHoverEvent* __a0) { QGraphicsRectItem::hoverMoveEvent(__a0); }
	void cppgoBase_hoverLeaveEvent(QGraphicsSceneHoverEvent* __a0) { QGraphicsRectItem::hoverLeaveEvent(__a0); }
	void cppgoBase_keyPressEvent(QKeyEvent* __a0) { QGraphicsRectItem::keyPressEvent(__a0); }
	void cppgoBase_keyReleaseEvent(QKeyEvent* __a0) { QGraphicsRectItem::keyReleaseEvent(__a0); }
	void cppgoBase_mousePressEvent(QGraphicsSceneMouseEvent* __a0) { QGraphicsRectItem::mousePressEvent(__a0); }
	void cppgoBase_mouseMoveEvent(QGraphicsSceneMouseEvent* __a0) { QGraphicsRectItem::mouseMoveEvent(__a0); }
	void cppgoBase_mouseReleaseEvent(QGraphicsSceneMouseEvent* __a0) { QGraphicsRectItem::mouseReleaseEvent(__a0); }
	void cppgoBase_mouseDoubleClickEvent(QGraphicsSceneMouseEvent* __a0) { QGraphicsRectItem::mouseDoubleClickEvent(__a0); }
	void cppgoBase_wheelEvent(QGraphicsSceneWheelEvent* __a0) { QGraphicsRectItem::wheelEvent(__a0); }
	void cppgoBase_inputMethodEvent(QInputMethodEvent* __a0) { QGraphicsRectItem::inputMethodEvent(__a0); }
	QVariant cppgoBase_inputMethodQuery(int __a0) { return QGraphicsRectItem::inputMethodQuery((Qt::InputMethodQuery)__a0); }
	QVariant cppgoBase_itemChange(int __a0, const QVariant& __a1) { return QGraphicsRectItem::itemChange((QGraphicsItem::GraphicsItemChange)__a0, __a1); }
	QRectF boundingRect() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 2510008782785038831LL, &__h); if (__h) return *(QRectF*)__r; }
		return QGraphicsRectItem::boundingRect();
	}
	QPainterPath shape() const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 7938019069247444294LL, &__h); if (__h) return *(QPainterPath*)__r; }
		return QGraphicsRectItem::shape();
	}
	bool contains(const QPointF& point) const override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 5645366511944179768LL, (void*)&point, &__h); if (__h) return __r; }
		return QGraphicsRectItem::contains(point);
	}
	void paint(QPainter* painter, const QStyleOptionGraphicsItem* option, QWidget* widget) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_voidp_voidp_voidp) { ((void(*)(long long, long long, void*, void*, void*, int*))cppgo_virt_void_voidp_voidp_voidp)(__goId, 6083584261053556819LL, (void*)painter, (void*)option, (void*)widget, &__h); if (__h) return; }
		QGraphicsRectItem::paint(painter, option, widget);
	}
	bool isObscuredBy(const QGraphicsItem* item) const override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, -7215030271914903671LL, (void*)item, &__h); if (__h) return __r; }
		return QGraphicsRectItem::isObscuredBy(item);
	}
	QPainterPath opaqueArea() const override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 399370970023989435LL, &__h); if (__h) return *(QPainterPath*)__r; }
		return QGraphicsRectItem::opaqueArea();
	}
	int type() const override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -6371403872082420691LL, &__h); if (__h) return __r; }
		return QGraphicsRectItem::type();
	}
	bool supportsExtension(QGraphicsItem::Extension extension) const override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, -5729166443651306418LL, (int)(extension), &__h); if (__h) return __r; }
		return QGraphicsRectItem::supportsExtension(extension);
	}
	void setExtension(QGraphicsItem::Extension extension, const QVariant& variant) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void_int_voidp) { ((void(*)(long long, long long, int, void*, int*))cppgo_virt_void_int_voidp)(__goId, -6730181576106171580LL, (int)(extension), (void*)&variant, &__h); if (__h) return; }
		QGraphicsRectItem::setExtension(extension, variant);
	}
	QVariant extension(const QVariant& variant) const override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 2270616358835251226LL, (void*)&variant, &__h); if (__h) return *(QVariant*)__r; }
		return QGraphicsRectItem::extension(variant);
	}
	void advance(int phase) override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, -4830836341480073475LL, phase, &__h); if (__h) return; }
		QGraphicsRectItem::advance(phase);
	}
	bool collidesWithItem(const QGraphicsItem* other, Qt::ItemSelectionMode mode) const override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_int_voidp_int) { int __r = ((int(*)(long long, long long, void*, int, int*))cppgo_virt_int_voidp_int)(__goId, 6469736184362593173LL, (void*)other, (int)(mode), &__h); if (__h) return __r; }
		return QGraphicsRectItem::collidesWithItem(other, mode);
	}
	bool collidesWithPath(const QPainterPath& path, Qt::ItemSelectionMode mode) const override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_int_voidp_int) { int __r = ((int(*)(long long, long long, void*, int, int*))cppgo_virt_int_voidp_int)(__goId, 6636575182718753853LL, (void*)&path, (int)(mode), &__h); if (__h) return __r; }
		return QGraphicsRectItem::collidesWithPath(path, mode);
	}
	bool sceneEventFilter(QGraphicsItem* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, -4584203218241215725LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QGraphicsRectItem::sceneEventFilter(watched, event);
	}
	bool sceneEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, -2846371507838587071LL, (void*)event, &__h); if (__h) return __r; }
		return QGraphicsRectItem::sceneEvent(event);
	}
	void contextMenuEvent(QGraphicsSceneContextMenuEvent* event) override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6698043222227006015LL, (void*)event, &__h); if (__h) return; }
		QGraphicsRectItem::contextMenuEvent(event);
	}
	void dragEnterEvent(QGraphicsSceneDragDropEvent* event) override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 206122715919917589LL, (void*)event, &__h); if (__h) return; }
		QGraphicsRectItem::dragEnterEvent(event);
	}
	void dragLeaveEvent(QGraphicsSceneDragDropEvent* event) override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -3139904158373725006LL, (void*)event, &__h); if (__h) return; }
		QGraphicsRectItem::dragLeaveEvent(event);
	}
	void dragMoveEvent(QGraphicsSceneDragDropEvent* event) override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1781545137352745978LL, (void*)event, &__h); if (__h) return; }
		QGraphicsRectItem::dragMoveEvent(event);
	}
	void dropEvent(QGraphicsSceneDragDropEvent* event) override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 9107058682754591716LL, (void*)event, &__h); if (__h) return; }
		QGraphicsRectItem::dropEvent(event);
	}
	void focusInEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 975043767033552552LL, (void*)event, &__h); if (__h) return; }
		QGraphicsRectItem::focusInEvent(event);
	}
	void focusOutEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -385596010576048993LL, (void*)event, &__h); if (__h) return; }
		QGraphicsRectItem::focusOutEvent(event);
	}
	void hoverEnterEvent(QGraphicsSceneHoverEvent* event) override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6932941943906681679LL, (void*)event, &__h); if (__h) return; }
		QGraphicsRectItem::hoverEnterEvent(event);
	}
	void hoverMoveEvent(QGraphicsSceneHoverEvent* event) override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -2745191522144316396LL, (void*)event, &__h); if (__h) return; }
		QGraphicsRectItem::hoverMoveEvent(event);
	}
	void hoverLeaveEvent(QGraphicsSceneHoverEvent* event) override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7944408482203603984LL, (void*)event, &__h); if (__h) return; }
		QGraphicsRectItem::hoverLeaveEvent(event);
	}
	void keyPressEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1023791977679793185LL, (void*)event, &__h); if (__h) return; }
		QGraphicsRectItem::keyPressEvent(event);
	}
	void keyReleaseEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 772598475950050299LL, (void*)event, &__h); if (__h) return; }
		QGraphicsRectItem::keyReleaseEvent(event);
	}
	void mousePressEvent(QGraphicsSceneMouseEvent* event) override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -332090898078543319LL, (void*)event, &__h); if (__h) return; }
		QGraphicsRectItem::mousePressEvent(event);
	}
	void mouseMoveEvent(QGraphicsSceneMouseEvent* event) override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4255494004099806581LL, (void*)event, &__h); if (__h) return; }
		QGraphicsRectItem::mouseMoveEvent(event);
	}
	void mouseReleaseEvent(QGraphicsSceneMouseEvent* event) override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4415920049297795213LL, (void*)event, &__h); if (__h) return; }
		QGraphicsRectItem::mouseReleaseEvent(event);
	}
	void mouseDoubleClickEvent(QGraphicsSceneMouseEvent* event) override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4029073905697023887LL, (void*)event, &__h); if (__h) return; }
		QGraphicsRectItem::mouseDoubleClickEvent(event);
	}
	void wheelEvent(QGraphicsSceneWheelEvent* event) override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1919988340934511316LL, (void*)event, &__h); if (__h) return; }
		QGraphicsRectItem::wheelEvent(event);
	}
	void inputMethodEvent(QInputMethodEvent* event) override {
		int __h = 0;
		if (__goOn[32] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5438819661942843222LL, (void*)event, &__h); if (__h) return; }
		QGraphicsRectItem::inputMethodEvent(event);
	}
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const override {
		int __h = 0;
		if (__goOn[33] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7345853097340709638LL, (int)(query), &__h); if (__h) return *(QVariant*)__r; }
		return QGraphicsRectItem::inputMethodQuery(query);
	}
	QVariant itemChange(QGraphicsItem::GraphicsItemChange change, const QVariant& value) override {
		int __h = 0;
		if (__goOn[34] && cppgo_virt_voidp_int_voidp) { void* __r = ((void*(*)(long long, long long, int, void*, int*))cppgo_virt_voidp_int_voidp)(__goId, -1888572168794878830LL, (int)(change), (void*)&value, &__h); if (__h) return *(QVariant*)__r; }
		return QGraphicsRectItem::itemChange(change, value);
	}
};

extern "C" {

void* QGraphicsRectItem_base(void* self) { return (void*)static_cast<QAbstractGraphicsShapeItem*>((QGraphicsRectItem*)self); }
void QGraphicsRectItem_chain(void* self, void** out) {
	QGraphicsRectItem* __p = (QGraphicsRectItem*)self;
	QAbstractGraphicsShapeItem* __b0 = static_cast<QAbstractGraphicsShapeItem*>(__p); out[0] = (void*)__b0;
	QGraphicsItem* __b1 = static_cast<QGraphicsItem*>(__b0); out[1] = (void*)__b1;
}
void QGraphicsRectItem__set_override(void* self, int idx) { ((QGraphicsRectItem__Go*)self)->__goOn[idx] = 1; }
void* QGraphicsRectItem_new(long long __goid) {
	return new QGraphicsRectItem__Go(__goid);
}
void* QGraphicsRectItem_new2(long long __goid, void* rect) {
	return new QGraphicsRectItem__Go(__goid, *(QRectF*)rect);
}
void* QGraphicsRectItem_new3(long long __goid, double x, double y, double w, double h) {
	return new QGraphicsRectItem__Go(__goid, x, y, w, h);
}
void* QGraphicsRectItem_new4(long long __goid, void* parent) {
	return new QGraphicsRectItem__Go(__goid, (QGraphicsItem*)parent);
}
void* QGraphicsRectItem_new5(long long __goid, void* rect, void* parent) {
	return new QGraphicsRectItem__Go(__goid, *(QRectF*)rect, (QGraphicsItem*)parent);
}
void* QGraphicsRectItem_new6(long long __goid, double x, double y, double w, double h, void* parent) {
	return new QGraphicsRectItem__Go(__goid, x, y, w, h, (QGraphicsItem*)parent);
}
void* QGraphicsRectItem_rect(void* self) {
	auto&& __r = ((QGraphicsRectItem*)self)->rect();
	return new QRectF(__r);
}
void QGraphicsRectItem_setRect(void* self, void* rect) {
	((QGraphicsRectItem*)self)->setRect(*(QRectF*)rect);
}
void QGraphicsRectItem_setRect2(void* self, double x, double y, double w, double h) {
	((QGraphicsRectItem*)self)->setRect(x, y, w, h);
}
void QGraphicsRectItem_updateMicroFocus(void* self) {
	((QGraphicsRectItem__Go*)self)->updateMicroFocus();
}
void QGraphicsRectItem_addToIndex(void* self) {
	((QGraphicsRectItem__Go*)self)->addToIndex();
}
void QGraphicsRectItem_removeFromIndex(void* self) {
	((QGraphicsRectItem__Go*)self)->removeFromIndex();
}
void QGraphicsRectItem_prepareGeometryChange(void* self) {
	((QGraphicsRectItem__Go*)self)->prepareGeometryChange();
}
void* QGraphicsRectItem_boundingRect_default(void* self) {
	auto&& __r = ((QGraphicsRectItem*)self)->QGraphicsRectItem::boundingRect();
	return new QRectF(__r);
}
void* QGraphicsRectItem_shape_default(void* self) {
	auto&& __r = ((QGraphicsRectItem*)self)->QGraphicsRectItem::shape();
	return new QPainterPath(__r);
}
int QGraphicsRectItem_contains_default(void* self, void* point) {
	auto&& __r = ((QGraphicsRectItem*)self)->QGraphicsRectItem::contains(*(QPointF*)point);
	return __r;
}
void QGraphicsRectItem_paint_default(void* self, void* painter, void* option, void* widget) {
	((QGraphicsRectItem*)self)->QGraphicsRectItem::paint((QPainter*)painter, (QStyleOptionGraphicsItem*)option, (QWidget*)widget);
}
int QGraphicsRectItem_isObscuredBy_default(void* self, void* item) {
	auto&& __r = ((QGraphicsRectItem*)self)->QGraphicsRectItem::isObscuredBy((QGraphicsItem*)item);
	return __r;
}
void* QGraphicsRectItem_opaqueArea_default(void* self) {
	auto&& __r = ((QGraphicsRectItem*)self)->QGraphicsRectItem::opaqueArea();
	return new QPainterPath(__r);
}
int QGraphicsRectItem_type_default(void* self) {
	auto&& __r = ((QGraphicsRectItem*)self)->QGraphicsRectItem::type();
	return __r;
}
int QGraphicsRectItem_supportsExtension_default(void* self, int extension) {
	auto&& __r = ((QGraphicsRectItem__Go*)self)->cppgoBase_supportsExtension(extension);
	return __r;
}
void QGraphicsRectItem_setExtension_default(void* self, int extension, void* variant) {
	((QGraphicsRectItem__Go*)self)->cppgoBase_setExtension(extension, *(QVariant*)variant);
}
void* QGraphicsRectItem_extension_default(void* self, void* variant) {
	auto&& __r = ((QGraphicsRectItem__Go*)self)->cppgoBase_extension(*(QVariant*)variant);
	return new QVariant(__r);
}
void QGraphicsRectItem_advance_default(void* self, int phase) {
	((QGraphicsRectItem*)self)->QGraphicsRectItem::advance(phase);
}
int QGraphicsRectItem_collidesWithItem_default(void* self, void* other, int mode) {
	auto&& __r = ((QGraphicsRectItem*)self)->QGraphicsRectItem::collidesWithItem((QGraphicsItem*)other, (Qt::ItemSelectionMode)(mode));
	return __r;
}
int QGraphicsRectItem_collidesWithPath_default(void* self, void* path, int mode) {
	auto&& __r = ((QGraphicsRectItem*)self)->QGraphicsRectItem::collidesWithPath(*(QPainterPath*)path, (Qt::ItemSelectionMode)(mode));
	return __r;
}
int QGraphicsRectItem_sceneEventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QGraphicsRectItem__Go*)self)->cppgoBase_sceneEventFilter((QGraphicsItem*)watched, (QEvent*)event);
	return __r;
}
int QGraphicsRectItem_sceneEvent_default(void* self, void* event) {
	auto&& __r = ((QGraphicsRectItem__Go*)self)->cppgoBase_sceneEvent((QEvent*)event);
	return __r;
}
void QGraphicsRectItem_contextMenuEvent_default(void* self, void* event) {
	((QGraphicsRectItem__Go*)self)->cppgoBase_contextMenuEvent((QGraphicsSceneContextMenuEvent*)event);
}
void QGraphicsRectItem_dragEnterEvent_default(void* self, void* event) {
	((QGraphicsRectItem__Go*)self)->cppgoBase_dragEnterEvent((QGraphicsSceneDragDropEvent*)event);
}
void QGraphicsRectItem_dragLeaveEvent_default(void* self, void* event) {
	((QGraphicsRectItem__Go*)self)->cppgoBase_dragLeaveEvent((QGraphicsSceneDragDropEvent*)event);
}
void QGraphicsRectItem_dragMoveEvent_default(void* self, void* event) {
	((QGraphicsRectItem__Go*)self)->cppgoBase_dragMoveEvent((QGraphicsSceneDragDropEvent*)event);
}
void QGraphicsRectItem_dropEvent_default(void* self, void* event) {
	((QGraphicsRectItem__Go*)self)->cppgoBase_dropEvent((QGraphicsSceneDragDropEvent*)event);
}
void QGraphicsRectItem_focusInEvent_default(void* self, void* event) {
	((QGraphicsRectItem__Go*)self)->cppgoBase_focusInEvent((QFocusEvent*)event);
}
void QGraphicsRectItem_focusOutEvent_default(void* self, void* event) {
	((QGraphicsRectItem__Go*)self)->cppgoBase_focusOutEvent((QFocusEvent*)event);
}
void QGraphicsRectItem_hoverEnterEvent_default(void* self, void* event) {
	((QGraphicsRectItem__Go*)self)->cppgoBase_hoverEnterEvent((QGraphicsSceneHoverEvent*)event);
}
void QGraphicsRectItem_hoverMoveEvent_default(void* self, void* event) {
	((QGraphicsRectItem__Go*)self)->cppgoBase_hoverMoveEvent((QGraphicsSceneHoverEvent*)event);
}
void QGraphicsRectItem_hoverLeaveEvent_default(void* self, void* event) {
	((QGraphicsRectItem__Go*)self)->cppgoBase_hoverLeaveEvent((QGraphicsSceneHoverEvent*)event);
}
void QGraphicsRectItem_keyPressEvent_default(void* self, void* event) {
	((QGraphicsRectItem__Go*)self)->cppgoBase_keyPressEvent((QKeyEvent*)event);
}
void QGraphicsRectItem_keyReleaseEvent_default(void* self, void* event) {
	((QGraphicsRectItem__Go*)self)->cppgoBase_keyReleaseEvent((QKeyEvent*)event);
}
void QGraphicsRectItem_mousePressEvent_default(void* self, void* event) {
	((QGraphicsRectItem__Go*)self)->cppgoBase_mousePressEvent((QGraphicsSceneMouseEvent*)event);
}
void QGraphicsRectItem_mouseMoveEvent_default(void* self, void* event) {
	((QGraphicsRectItem__Go*)self)->cppgoBase_mouseMoveEvent((QGraphicsSceneMouseEvent*)event);
}
void QGraphicsRectItem_mouseReleaseEvent_default(void* self, void* event) {
	((QGraphicsRectItem__Go*)self)->cppgoBase_mouseReleaseEvent((QGraphicsSceneMouseEvent*)event);
}
void QGraphicsRectItem_mouseDoubleClickEvent_default(void* self, void* event) {
	((QGraphicsRectItem__Go*)self)->cppgoBase_mouseDoubleClickEvent((QGraphicsSceneMouseEvent*)event);
}
void QGraphicsRectItem_wheelEvent_default(void* self, void* event) {
	((QGraphicsRectItem__Go*)self)->cppgoBase_wheelEvent((QGraphicsSceneWheelEvent*)event);
}
void QGraphicsRectItem_inputMethodEvent_default(void* self, void* event) {
	((QGraphicsRectItem__Go*)self)->cppgoBase_inputMethodEvent((QInputMethodEvent*)event);
}
void* QGraphicsRectItem_inputMethodQuery_default(void* self, int query) {
	auto&& __r = ((QGraphicsRectItem__Go*)self)->cppgoBase_inputMethodQuery(query);
	return new QVariant(__r);
}
void* QGraphicsRectItem_itemChange_default(void* self, int change, void* value) {
	auto&& __r = ((QGraphicsRectItem__Go*)self)->cppgoBase_itemChange(change, *(QVariant*)value);
	return new QVariant(__r);
}
void* QGraphicsRectItem__Go_new(long long id) { return new QGraphicsRectItem__Go(id); }
void QGraphicsRectItem_delete(void* self) { delete (QGraphicsRectItem*)self; }

}

#include <QtWidgets/qgraphicssceneevent.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qpoint.h>
#include <QtCore/qnamespace.h>
#include <QtWidgets/qwidget.h>
#include "gen_qgraphicsscenemouseevent.h"
#include "gen_purego_cb.h"

class QGraphicsSceneMouseEvent__Go : public QGraphicsSceneMouseEvent {
public:
	long long __goId;
	unsigned char __goOn[2] = {};
	template<typename... A> QGraphicsSceneMouseEvent__Go(long long id, A&&... a) : QGraphicsSceneMouseEvent(std::forward<A>(a)...), __goId(id) {}
	void setAccepted(bool accepted) override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 6367896784077321574LL, accepted, &__h); if (__h) return; }
		QGraphicsSceneMouseEvent::setAccepted(accepted);
	}
	QEvent* clone() const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, -7032813281048616478LL, &__h); if (__h) return (QEvent*)__r; }
		return QGraphicsSceneMouseEvent::clone();
	}
};

extern "C" {

void* QGraphicsSceneMouseEvent_base(void* self) { return (void*)static_cast<QGraphicsSceneEvent*>((QGraphicsSceneMouseEvent*)self); }
void QGraphicsSceneMouseEvent_chain(void* self, void** out) {
	QGraphicsSceneMouseEvent* __p = (QGraphicsSceneMouseEvent*)self;
	QGraphicsSceneEvent* __b0 = static_cast<QGraphicsSceneEvent*>(__p); out[0] = (void*)__b0;
	QEvent* __b1 = static_cast<QEvent*>(__b0); out[1] = (void*)__b1;
}
void QGraphicsSceneMouseEvent__set_override(void* self, int idx) { ((QGraphicsSceneMouseEvent__Go*)self)->__goOn[idx] = 1; }
void* QGraphicsSceneMouseEvent_new(long long __goid) {
	return new QGraphicsSceneMouseEvent__Go(__goid);
}
void* QGraphicsSceneMouseEvent_new2(long long __goid, int type_) {
	return new QGraphicsSceneMouseEvent__Go(__goid, (QEvent::Type)(type_));
}
void* QGraphicsSceneMouseEvent_pos(void* self) {
	auto&& __r = ((QGraphicsSceneMouseEvent*)self)->pos();
	return new QPointF(__r);
}
void QGraphicsSceneMouseEvent_setPos(void* self, void* pos) {
	((QGraphicsSceneMouseEvent*)self)->setPos(*(QPointF*)pos);
}
void* QGraphicsSceneMouseEvent_scenePos(void* self) {
	auto&& __r = ((QGraphicsSceneMouseEvent*)self)->scenePos();
	return new QPointF(__r);
}
void QGraphicsSceneMouseEvent_setScenePos(void* self, void* pos) {
	((QGraphicsSceneMouseEvent*)self)->setScenePos(*(QPointF*)pos);
}
void* QGraphicsSceneMouseEvent_screenPos(void* self) {
	auto&& __r = ((QGraphicsSceneMouseEvent*)self)->screenPos();
	return new QPoint(__r);
}
void QGraphicsSceneMouseEvent_setScreenPos(void* self, void* pos) {
	((QGraphicsSceneMouseEvent*)self)->setScreenPos(*(QPoint*)pos);
}
void* QGraphicsSceneMouseEvent_buttonDownPos(void* self, int button) {
	auto&& __r = ((QGraphicsSceneMouseEvent*)self)->buttonDownPos((Qt::MouseButton)(button));
	return new QPointF(__r);
}
void QGraphicsSceneMouseEvent_setButtonDownPos(void* self, int button, void* pos) {
	((QGraphicsSceneMouseEvent*)self)->setButtonDownPos((Qt::MouseButton)(button), *(QPointF*)pos);
}
void* QGraphicsSceneMouseEvent_buttonDownScenePos(void* self, int button) {
	auto&& __r = ((QGraphicsSceneMouseEvent*)self)->buttonDownScenePos((Qt::MouseButton)(button));
	return new QPointF(__r);
}
void QGraphicsSceneMouseEvent_setButtonDownScenePos(void* self, int button, void* pos) {
	((QGraphicsSceneMouseEvent*)self)->setButtonDownScenePos((Qt::MouseButton)(button), *(QPointF*)pos);
}
void* QGraphicsSceneMouseEvent_buttonDownScreenPos(void* self, int button) {
	auto&& __r = ((QGraphicsSceneMouseEvent*)self)->buttonDownScreenPos((Qt::MouseButton)(button));
	return new QPoint(__r);
}
void QGraphicsSceneMouseEvent_setButtonDownScreenPos(void* self, int button, void* pos) {
	((QGraphicsSceneMouseEvent*)self)->setButtonDownScreenPos((Qt::MouseButton)(button), *(QPoint*)pos);
}
void* QGraphicsSceneMouseEvent_lastPos(void* self) {
	auto&& __r = ((QGraphicsSceneMouseEvent*)self)->lastPos();
	return new QPointF(__r);
}
void QGraphicsSceneMouseEvent_setLastPos(void* self, void* pos) {
	((QGraphicsSceneMouseEvent*)self)->setLastPos(*(QPointF*)pos);
}
void* QGraphicsSceneMouseEvent_lastScenePos(void* self) {
	auto&& __r = ((QGraphicsSceneMouseEvent*)self)->lastScenePos();
	return new QPointF(__r);
}
void QGraphicsSceneMouseEvent_setLastScenePos(void* self, void* pos) {
	((QGraphicsSceneMouseEvent*)self)->setLastScenePos(*(QPointF*)pos);
}
void* QGraphicsSceneMouseEvent_lastScreenPos(void* self) {
	auto&& __r = ((QGraphicsSceneMouseEvent*)self)->lastScreenPos();
	return new QPoint(__r);
}
void QGraphicsSceneMouseEvent_setLastScreenPos(void* self, void* pos) {
	((QGraphicsSceneMouseEvent*)self)->setLastScreenPos(*(QPoint*)pos);
}
int QGraphicsSceneMouseEvent_buttons(void* self) {
	auto&& __r = ((QGraphicsSceneMouseEvent*)self)->buttons();
	return (int)(__r);
}
void QGraphicsSceneMouseEvent_setButtons(void* self, int buttons) {
	((QGraphicsSceneMouseEvent*)self)->setButtons(QFlags<Qt::MouseButton>(buttons));
}
int QGraphicsSceneMouseEvent_button(void* self) {
	auto&& __r = ((QGraphicsSceneMouseEvent*)self)->button();
	return (int)(__r);
}
void QGraphicsSceneMouseEvent_setButton(void* self, int button) {
	((QGraphicsSceneMouseEvent*)self)->setButton((Qt::MouseButton)(button));
}
int QGraphicsSceneMouseEvent_modifiers(void* self) {
	auto&& __r = ((QGraphicsSceneMouseEvent*)self)->modifiers();
	return (int)(__r);
}
void QGraphicsSceneMouseEvent_setModifiers(void* self, int modifiers) {
	((QGraphicsSceneMouseEvent*)self)->setModifiers(QFlags<Qt::KeyboardModifier>(modifiers));
}
int QGraphicsSceneMouseEvent_source(void* self) {
	auto&& __r = ((QGraphicsSceneMouseEvent*)self)->source();
	return (int)(__r);
}
void QGraphicsSceneMouseEvent_setSource(void* self, int source) {
	((QGraphicsSceneMouseEvent*)self)->setSource((Qt::MouseEventSource)(source));
}
int QGraphicsSceneMouseEvent_flags(void* self) {
	auto&& __r = ((QGraphicsSceneMouseEvent*)self)->flags();
	return (int)(__r);
}
void QGraphicsSceneMouseEvent_setFlags(void* self, int arg1) {
	((QGraphicsSceneMouseEvent*)self)->setFlags(QFlags<Qt::MouseEventFlag>(arg1));
}
int QGraphicsSceneMouseEvent_registerEventType() {
	auto&& __r = QGraphicsSceneMouseEvent::registerEventType();
	return __r;
}
int QGraphicsSceneMouseEvent_registerEventType2(int hint) {
	auto&& __r = QGraphicsSceneMouseEvent::registerEventType(hint);
	return __r;
}
void QGraphicsSceneMouseEvent_setAccepted_default(void* self, int accepted) {
	((QGraphicsSceneMouseEvent*)self)->QGraphicsSceneMouseEvent::setAccepted(accepted);
}
void* QGraphicsSceneMouseEvent_clone_default(void* self) {
	auto&& __r = ((QGraphicsSceneMouseEvent*)self)->QGraphicsSceneMouseEvent::clone();
	return (void*)(__r);
}
void* QGraphicsSceneMouseEvent__Go_new(long long id) { return new QGraphicsSceneMouseEvent__Go(id); }
void QGraphicsSceneMouseEvent_delete(void* self) { delete (QGraphicsSceneMouseEvent*)self; }

}

#include <QtWidgets/qlistwidget.h>
#include <QtCore/qstring.h>
#include <QtGui/qicon.h>
#include <QtCore/qnamespace.h>
#include <QtGui/qfont.h>
#include <QtGui/qbrush.h>
#include <QtCore/qsize.h>
#include <QtCore/qvariant.h>
#include <QtCore/qdatastream.h>
#include "gen_qlistwidgetitem.h"
#include "gen_purego_cb.h"

class QListWidgetItem__Go : public QListWidgetItem {
public:
	long long __goId;
	unsigned char __goOn[5] = {};
	template<typename... A> QListWidgetItem__Go(long long id, A&&... a) : QListWidgetItem(std::forward<A>(a)...), __goId(id) {}
	QListWidgetItem* clone() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, -7032813281048616478LL, &__h); if (__h) return (QListWidgetItem*)__r; }
		return QListWidgetItem::clone();
	}
	QVariant data(int role) const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, -8837375942098466299LL, role, &__h); if (__h) return *(QVariant*)__r; }
		return QListWidgetItem::data(role);
	}
	void setData(int role, const QVariant& value) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_void_int_voidp) { ((void(*)(long long, long long, int, void*, int*))cppgo_virt_void_int_voidp)(__goId, 35191454575202219LL, role, (void*)&value, &__h); if (__h) return; }
		QListWidgetItem::setData(role, value);
	}
	void read(QDataStream& in) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5541207787354052101LL, (void*)&in, &__h); if (__h) return; }
		QListWidgetItem::read(in);
	}
	void write(QDataStream& out) const override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5099743077423337732LL, (void*)&out, &__h); if (__h) return; }
		QListWidgetItem::write(out);
	}
};

extern "C" {

void QListWidgetItem__set_override(void* self, int idx) { ((QListWidgetItem__Go*)self)->__goOn[idx] = 1; }
void* QListWidgetItem_new(long long __goid) {
	return new QListWidgetItem__Go(__goid);
}
void* QListWidgetItem_new2(long long __goid, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	return new QListWidgetItem__Go(__goid, text_qstr);
}
void* QListWidgetItem_new3(long long __goid, void* icon, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	return new QListWidgetItem__Go(__goid, *(QIcon*)icon, text_qstr);
}
void* QListWidgetItem_new4(long long __goid, void* other) {
	return new QListWidgetItem__Go(__goid, *(QListWidgetItem*)other);
}
void* QListWidgetItem_new5(long long __goid, void* listview) {
	return new QListWidgetItem__Go(__goid, (QListWidget*)listview);
}
void* QListWidgetItem_new6(long long __goid, void* listview, int type_) {
	return new QListWidgetItem__Go(__goid, (QListWidget*)listview, type_);
}
void* QListWidgetItem_new7(long long __goid, const char* text, long long text_len, void* listview) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	return new QListWidgetItem__Go(__goid, text_qstr, (QListWidget*)listview);
}
void* QListWidgetItem_new8(long long __goid, const char* text, long long text_len, void* listview, int type_) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	return new QListWidgetItem__Go(__goid, text_qstr, (QListWidget*)listview, type_);
}
void* QListWidgetItem_new9(long long __goid, void* icon, const char* text, long long text_len, void* listview) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	return new QListWidgetItem__Go(__goid, *(QIcon*)icon, text_qstr, (QListWidget*)listview);
}
void* QListWidgetItem_new10(long long __goid, void* icon, const char* text, long long text_len, void* listview, int type_) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	return new QListWidgetItem__Go(__goid, *(QIcon*)icon, text_qstr, (QListWidget*)listview, type_);
}
void* QListWidgetItem_clone(void* self) {
	auto&& __r = ((QListWidgetItem*)self)->clone();
	return (void*)(__r);
}
void* QListWidgetItem_listWidget(void* self) {
	auto&& __r = ((QListWidgetItem*)self)->listWidget();
	return (void*)(__r);
}
void QListWidgetItem_setSelected(void* self, int select_) {
	((QListWidgetItem*)self)->setSelected(select_);
}
int QListWidgetItem_isSelected(void* self) {
	auto&& __r = ((QListWidgetItem*)self)->isSelected();
	return __r;
}
void QListWidgetItem_setHidden(void* self, int hide) {
	((QListWidgetItem*)self)->setHidden(hide);
}
int QListWidgetItem_isHidden(void* self) {
	auto&& __r = ((QListWidgetItem*)self)->isHidden();
	return __r;
}
int QListWidgetItem_flags(void* self) {
	auto&& __r = ((QListWidgetItem*)self)->flags();
	return (int)(__r);
}
void QListWidgetItem_setFlags(void* self, int flags) {
	((QListWidgetItem*)self)->setFlags(QFlags<Qt::ItemFlag>(flags));
}
char* QListWidgetItem_text(void* self, long long* __out_len) {
	auto&& __r = ((QListWidgetItem*)self)->text();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QListWidgetItem_setText(void* self, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	((QListWidgetItem*)self)->setText(text_qstr);
}
void* QListWidgetItem_icon(void* self) {
	auto&& __r = ((QListWidgetItem*)self)->icon();
	return new QIcon(__r);
}
void QListWidgetItem_setIcon(void* self, void* icon) {
	((QListWidgetItem*)self)->setIcon(*(QIcon*)icon);
}
char* QListWidgetItem_statusTip(void* self, long long* __out_len) {
	auto&& __r = ((QListWidgetItem*)self)->statusTip();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QListWidgetItem_setStatusTip(void* self, const char* statusTip, long long statusTip_len) {
	QString statusTip_qstr = QString::fromUtf8(statusTip, statusTip_len);
	((QListWidgetItem*)self)->setStatusTip(statusTip_qstr);
}
char* QListWidgetItem_toolTip(void* self, long long* __out_len) {
	auto&& __r = ((QListWidgetItem*)self)->toolTip();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QListWidgetItem_setToolTip(void* self, const char* toolTip, long long toolTip_len) {
	QString toolTip_qstr = QString::fromUtf8(toolTip, toolTip_len);
	((QListWidgetItem*)self)->setToolTip(toolTip_qstr);
}
char* QListWidgetItem_whatsThis(void* self, long long* __out_len) {
	auto&& __r = ((QListWidgetItem*)self)->whatsThis();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QListWidgetItem_setWhatsThis(void* self, const char* whatsThis, long long whatsThis_len) {
	QString whatsThis_qstr = QString::fromUtf8(whatsThis, whatsThis_len);
	((QListWidgetItem*)self)->setWhatsThis(whatsThis_qstr);
}
void* QListWidgetItem_font(void* self) {
	auto&& __r = ((QListWidgetItem*)self)->font();
	return new QFont(__r);
}
void QListWidgetItem_setFont(void* self, void* font) {
	((QListWidgetItem*)self)->setFont(*(QFont*)font);
}
int QListWidgetItem_textAlignment(void* self) {
	auto&& __r = ((QListWidgetItem*)self)->textAlignment();
	return __r;
}
void QListWidgetItem_setTextAlignment(void* self, int alignment) {
	((QListWidgetItem*)self)->setTextAlignment(alignment);
}
void QListWidgetItem_setTextAlignment2(void* self, int alignment) {
	((QListWidgetItem*)self)->setTextAlignment((Qt::AlignmentFlag)(alignment));
}
void QListWidgetItem_setTextAlignment3(void* self, int alignment) {
	((QListWidgetItem*)self)->setTextAlignment(QFlags<Qt::AlignmentFlag>(alignment));
}
void* QListWidgetItem_background(void* self) {
	auto&& __r = ((QListWidgetItem*)self)->background();
	return new QBrush(__r);
}
void QListWidgetItem_setBackground(void* self, void* brush) {
	((QListWidgetItem*)self)->setBackground(*(QBrush*)brush);
}
void* QListWidgetItem_foreground(void* self) {
	auto&& __r = ((QListWidgetItem*)self)->foreground();
	return new QBrush(__r);
}
void QListWidgetItem_setForeground(void* self, void* brush) {
	((QListWidgetItem*)self)->setForeground(*(QBrush*)brush);
}
int QListWidgetItem_checkState(void* self) {
	auto&& __r = ((QListWidgetItem*)self)->checkState();
	return (int)(__r);
}
void QListWidgetItem_setCheckState(void* self, int state) {
	((QListWidgetItem*)self)->setCheckState((Qt::CheckState)(state));
}
void* QListWidgetItem_sizeHint(void* self) {
	auto&& __r = ((QListWidgetItem*)self)->sizeHint();
	return new QSize(__r);
}
void QListWidgetItem_setSizeHint(void* self, void* size) {
	((QListWidgetItem*)self)->setSizeHint(*(QSize*)size);
}
void* QListWidgetItem_data(void* self, int role) {
	auto&& __r = ((QListWidgetItem*)self)->data(role);
	return new QVariant(__r);
}
void QListWidgetItem_setData(void* self, int role, void* value) {
	((QListWidgetItem*)self)->setData(role, *(QVariant*)value);
}
int QListWidgetItem_OperatorLesser(void* self, void* other) {
	auto&& __r = ((QListWidgetItem*)self)->operator<(*(QListWidgetItem*)other);
	return __r;
}
void QListWidgetItem_read(void* self, void* in) {
	((QListWidgetItem*)self)->read(*(QDataStream*)in);
}
void QListWidgetItem_write(void* self, void* out) {
	((QListWidgetItem*)self)->write(*(QDataStream*)out);
}
void* QListWidgetItem_OperatorAssign(void* self, void* other) {
	auto&& __r = ((QListWidgetItem*)self)->operator=(*(QListWidgetItem*)other);
	return (void*)&(__r);
}
int QListWidgetItem_type(void* self) {
	auto&& __r = ((QListWidgetItem*)self)->type();
	return __r;
}
void* QListWidgetItem_clone_default(void* self) {
	auto&& __r = ((QListWidgetItem*)self)->QListWidgetItem::clone();
	return (void*)(__r);
}
void* QListWidgetItem_data_default(void* self, int role) {
	auto&& __r = ((QListWidgetItem*)self)->QListWidgetItem::data(role);
	return new QVariant(__r);
}
void QListWidgetItem_setData_default(void* self, int role, void* value) {
	((QListWidgetItem*)self)->QListWidgetItem::setData(role, *(QVariant*)value);
}
void QListWidgetItem_read_default(void* self, void* in) {
	((QListWidgetItem*)self)->QListWidgetItem::read(*(QDataStream*)in);
}
void QListWidgetItem_write_default(void* self, void* out) {
	((QListWidgetItem*)self)->QListWidgetItem::write(*(QDataStream*)out);
}
void* QListWidgetItem__Go_new(long long id) { return new QListWidgetItem__Go(id); }
void QListWidgetItem_delete(void* self) { delete (QListWidgetItem*)self; }

}

#include <QtWidgets/qsplitter.h>
#include <QtWidgets/qframe.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qsize.h>
#include <QtCore/qbytearray.h>
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
#include <QtGui/qpixmap.h>
#include <QtWidgets/qgraphicseffect.h>
#include <QtGui/qicon.h>
#include <QtCore/qlocale.h>
#include <QtGui/qkeysequence.h>
#include <QtWidgets/qgraphicsproxywidget.h>
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
#include "gen_qsplitter.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QSplitter__Go : public QSplitter {
public:
	long long __goId;
	unsigned char __goOn[50] = {};
	template<typename... A> QSplitter__Go(long long id, A&&... a) : QSplitter(std::forward<A>(a)...), __goId(id) {}
	using QSplitter::moveSplitter;
	using QSplitter::setRubberBand;
	using QSplitter::closestLegalPosition;
	using QSplitter::drawFrame;
	using QSplitter::updateMicroFocus;
	using QSplitter::create;
	using QSplitter::destroy;
	using QSplitter::focusNextChild;
	using QSplitter::focusPreviousChild;
	using QSplitter::sender;
	using QSplitter::senderSignalIndex;
	using QSplitter::receivers;
	using QSplitter::isSignalConnected;
	using QSplitter::getDecodedMetricF;
	QSplitterHandle* cppgoBase_createHandle() { return QSplitter::createHandle(); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QSplitter::childEvent(__a0); }
	bool cppgoBase_event(QEvent* __a0) { return QSplitter::event(__a0); }
	void cppgoBase_resizeEvent(QResizeEvent* __a0) { QSplitter::resizeEvent(__a0); }
	void cppgoBase_changeEvent(QEvent* __a0) { QSplitter::changeEvent(__a0); }
	void cppgoBase_paintEvent(QPaintEvent* __a0) { QSplitter::paintEvent(__a0); }
	void cppgoBase_initStyleOption(QStyleOptionFrame* __a0) { QSplitter::initStyleOption(__a0); }
	void cppgoBase_mousePressEvent(QMouseEvent* __a0) { QSplitter::mousePressEvent(__a0); }
	void cppgoBase_mouseReleaseEvent(QMouseEvent* __a0) { QSplitter::mouseReleaseEvent(__a0); }
	void cppgoBase_mouseDoubleClickEvent(QMouseEvent* __a0) { QSplitter::mouseDoubleClickEvent(__a0); }
	void cppgoBase_mouseMoveEvent(QMouseEvent* __a0) { QSplitter::mouseMoveEvent(__a0); }
	void cppgoBase_wheelEvent(QWheelEvent* __a0) { QSplitter::wheelEvent(__a0); }
	void cppgoBase_keyPressEvent(QKeyEvent* __a0) { QSplitter::keyPressEvent(__a0); }
	void cppgoBase_keyReleaseEvent(QKeyEvent* __a0) { QSplitter::keyReleaseEvent(__a0); }
	void cppgoBase_focusInEvent(QFocusEvent* __a0) { QSplitter::focusInEvent(__a0); }
	void cppgoBase_focusOutEvent(QFocusEvent* __a0) { QSplitter::focusOutEvent(__a0); }
	void cppgoBase_enterEvent(QEnterEvent* __a0) { QSplitter::enterEvent(__a0); }
	void cppgoBase_leaveEvent(QEvent* __a0) { QSplitter::leaveEvent(__a0); }
	void cppgoBase_moveEvent(QMoveEvent* __a0) { QSplitter::moveEvent(__a0); }
	void cppgoBase_closeEvent(QCloseEvent* __a0) { QSplitter::closeEvent(__a0); }
	void cppgoBase_contextMenuEvent(QContextMenuEvent* __a0) { QSplitter::contextMenuEvent(__a0); }
	void cppgoBase_tabletEvent(QTabletEvent* __a0) { QSplitter::tabletEvent(__a0); }
	void cppgoBase_actionEvent(QActionEvent* __a0) { QSplitter::actionEvent(__a0); }
	void cppgoBase_dragEnterEvent(QDragEnterEvent* __a0) { QSplitter::dragEnterEvent(__a0); }
	void cppgoBase_dragMoveEvent(QDragMoveEvent* __a0) { QSplitter::dragMoveEvent(__a0); }
	void cppgoBase_dragLeaveEvent(QDragLeaveEvent* __a0) { QSplitter::dragLeaveEvent(__a0); }
	void cppgoBase_dropEvent(QDropEvent* __a0) { QSplitter::dropEvent(__a0); }
	void cppgoBase_showEvent(QShowEvent* __a0) { QSplitter::showEvent(__a0); }
	void cppgoBase_hideEvent(QHideEvent* __a0) { QSplitter::hideEvent(__a0); }
	int cppgoBase_metric(int __a0) { return QSplitter::metric((QPaintDevice::PaintDeviceMetric)__a0); }
	void cppgoBase_initPainter(QPainter* __a0) { QSplitter::initPainter(__a0); }
	QPaintDevice* cppgoBase_redirected(QPoint* __a0) { return QSplitter::redirected(__a0); }
	QPainter* cppgoBase_sharedPainter() { return QSplitter::sharedPainter(); }
	void cppgoBase_inputMethodEvent(QInputMethodEvent* __a0) { QSplitter::inputMethodEvent(__a0); }
	bool cppgoBase_focusNextPrevChild(bool __a0) { return QSplitter::focusNextPrevChild(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QSplitter::timerEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QSplitter::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QSplitter::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QSplitter::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QSplitter::metaObject();
	}
	QSize sizeHint() const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 729690261457684101LL, &__h); if (__h) return *(QSize*)__r; }
		return QSplitter::sizeHint();
	}
	QSize minimumSizeHint() const override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 7049264599837532491LL, &__h); if (__h) return *(QSize*)__r; }
		return QSplitter::minimumSizeHint();
	}
	QSplitterHandle* createHandle() override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 7399745205369955835LL, &__h); if (__h) return (QSplitterHandle*)__r; }
		return QSplitter::createHandle();
	}
	void childEvent(QChildEvent* arg1) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)arg1, &__h); if (__h) return; }
		QSplitter::childEvent(arg1);
	}
	bool event(QEvent* arg1) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)arg1, &__h); if (__h) return __r; }
		return QSplitter::event(arg1);
	}
	void resizeEvent(QResizeEvent* arg1) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4302917285655014661LL, (void*)arg1, &__h); if (__h) return; }
		QSplitter::resizeEvent(arg1);
	}
	void changeEvent(QEvent* arg1) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4959376063826165943LL, (void*)arg1, &__h); if (__h) return; }
		QSplitter::changeEvent(arg1);
	}
	void paintEvent(QPaintEvent* arg1) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5290442163167458465LL, (void*)arg1, &__h); if (__h) return; }
		QSplitter::paintEvent(arg1);
	}
	void initStyleOption(QStyleOptionFrame* option) const override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -488693769771308531LL, (void*)option, &__h); if (__h) return; }
		QSplitter::initStyleOption(option);
	}
	int devType() const override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8659679893868524044LL, &__h); if (__h) return __r; }
		return QSplitter::devType();
	}
	void setVisible(bool visible) override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 2110182719375245575LL, visible, &__h); if (__h) return; }
		QSplitter::setVisible(visible);
	}
	int heightForWidth(int arg1) const override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 4311534060495737339LL, arg1, &__h); if (__h) return __r; }
		return QSplitter::heightForWidth(arg1);
	}
	bool hasHeightForWidth() const override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8266762397995423915LL, &__h); if (__h) return __r; }
		return QSplitter::hasHeightForWidth();
	}
	QPaintEngine* paintEngine() const override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 6398007404756346381LL, &__h); if (__h) return (QPaintEngine*)__r; }
		return QSplitter::paintEngine();
	}
	void mousePressEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -332090898078543319LL, (void*)event, &__h); if (__h) return; }
		QSplitter::mousePressEvent(event);
	}
	void mouseReleaseEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4415920049297795213LL, (void*)event, &__h); if (__h) return; }
		QSplitter::mouseReleaseEvent(event);
	}
	void mouseDoubleClickEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4029073905697023887LL, (void*)event, &__h); if (__h) return; }
		QSplitter::mouseDoubleClickEvent(event);
	}
	void mouseMoveEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4255494004099806581LL, (void*)event, &__h); if (__h) return; }
		QSplitter::mouseMoveEvent(event);
	}
	void wheelEvent(QWheelEvent* event) override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1919988340934511316LL, (void*)event, &__h); if (__h) return; }
		QSplitter::wheelEvent(event);
	}
	void keyPressEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1023791977679793185LL, (void*)event, &__h); if (__h) return; }
		QSplitter::keyPressEvent(event);
	}
	void keyReleaseEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 772598475950050299LL, (void*)event, &__h); if (__h) return; }
		QSplitter::keyReleaseEvent(event);
	}
	void focusInEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 975043767033552552LL, (void*)event, &__h); if (__h) return; }
		QSplitter::focusInEvent(event);
	}
	void focusOutEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -385596010576048993LL, (void*)event, &__h); if (__h) return; }
		QSplitter::focusOutEvent(event);
	}
	void enterEvent(QEnterEvent* event) override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6316270346668237895LL, (void*)event, &__h); if (__h) return; }
		QSplitter::enterEvent(event);
	}
	void leaveEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -551070009623711640LL, (void*)event, &__h); if (__h) return; }
		QSplitter::leaveEvent(event);
	}
	void moveEvent(QMoveEvent* event) override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -9161949150489525396LL, (void*)event, &__h); if (__h) return; }
		QSplitter::moveEvent(event);
	}
	void closeEvent(QCloseEvent* event) override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1989937050019464783LL, (void*)event, &__h); if (__h) return; }
		QSplitter::closeEvent(event);
	}
	void contextMenuEvent(QContextMenuEvent* event) override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6698043222227006015LL, (void*)event, &__h); if (__h) return; }
		QSplitter::contextMenuEvent(event);
	}
	void tabletEvent(QTabletEvent* event) override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -212068866860701141LL, (void*)event, &__h); if (__h) return; }
		QSplitter::tabletEvent(event);
	}
	void actionEvent(QActionEvent* event) override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1697444140088555581LL, (void*)event, &__h); if (__h) return; }
		QSplitter::actionEvent(event);
	}
	void dragEnterEvent(QDragEnterEvent* event) override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 206122715919917589LL, (void*)event, &__h); if (__h) return; }
		QSplitter::dragEnterEvent(event);
	}
	void dragMoveEvent(QDragMoveEvent* event) override {
		int __h = 0;
		if (__goOn[32] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1781545137352745978LL, (void*)event, &__h); if (__h) return; }
		QSplitter::dragMoveEvent(event);
	}
	void dragLeaveEvent(QDragLeaveEvent* event) override {
		int __h = 0;
		if (__goOn[33] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -3139904158373725006LL, (void*)event, &__h); if (__h) return; }
		QSplitter::dragLeaveEvent(event);
	}
	void dropEvent(QDropEvent* event) override {
		int __h = 0;
		if (__goOn[34] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 9107058682754591716LL, (void*)event, &__h); if (__h) return; }
		QSplitter::dropEvent(event);
	}
	void showEvent(QShowEvent* event) override {
		int __h = 0;
		if (__goOn[35] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6825369174751916LL, (void*)event, &__h); if (__h) return; }
		QSplitter::showEvent(event);
	}
	void hideEvent(QHideEvent* event) override {
		int __h = 0;
		if (__goOn[36] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -199718707876719129LL, (void*)event, &__h); if (__h) return; }
		QSplitter::hideEvent(event);
	}
	bool nativeEvent(const QByteArray& eventType, void* message, qintptr* result) override {
		int __h = 0;
		if (__goOn[37] && cppgo_virt_int_voidp_voidp_longlongp) { int __r = ((int(*)(long long, long long, void*, void*, long long*, int*))cppgo_virt_int_voidp_voidp_longlongp)(__goId, -2039166586246443706LL, (void*)&eventType, (void*)message, (long long*)result, &__h); if (__h) return __r; }
		return QSplitter::nativeEvent(eventType, message, result);
	}
	int metric(QPaintDevice::PaintDeviceMetric arg1) const override {
		int __h = 0;
		if (__goOn[38] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 427629649267311777LL, (int)(arg1), &__h); if (__h) return __r; }
		return QSplitter::metric(arg1);
	}
	void initPainter(QPainter* painter) const override {
		int __h = 0;
		if (__goOn[39] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6477500296980050210LL, (void*)painter, &__h); if (__h) return; }
		QSplitter::initPainter(painter);
	}
	QPaintDevice* redirected(QPoint* offset) const override {
		int __h = 0;
		if (__goOn[40] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 4127240633310655048LL, (void*)offset, &__h); if (__h) return (QPaintDevice*)__r; }
		return QSplitter::redirected(offset);
	}
	QPainter* sharedPainter() const override {
		int __h = 0;
		if (__goOn[41] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 5792771328741774991LL, &__h); if (__h) return (QPainter*)__r; }
		return QSplitter::sharedPainter();
	}
	void inputMethodEvent(QInputMethodEvent* arg1) override {
		int __h = 0;
		if (__goOn[42] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5438819661942843222LL, (void*)arg1, &__h); if (__h) return; }
		QSplitter::inputMethodEvent(arg1);
	}
	QVariant inputMethodQuery(Qt::InputMethodQuery arg1) const override {
		int __h = 0;
		if (__goOn[43] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7345853097340709638LL, (int)(arg1), &__h); if (__h) return *(QVariant*)__r; }
		return QSplitter::inputMethodQuery(arg1);
	}
	bool focusNextPrevChild(bool next) override {
		int __h = 0;
		if (__goOn[44] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 5898328517505784191LL, next, &__h); if (__h) return __r; }
		return QSplitter::focusNextPrevChild(next);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[45] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QSplitter::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[46] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QSplitter::timerEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[47] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QSplitter::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[48] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QSplitter::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[49] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QSplitter::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QSplitter;
static std::mutex __connsMu_QSplitter;

extern "C" {

void* QSplitter_base(void* self) { return (void*)static_cast<QFrame*>((QSplitter*)self); }
void QSplitter_chain(void* self, void** out) {
	QSplitter* __p = (QSplitter*)self;
	QFrame* __b0 = static_cast<QFrame*>(__p); out[0] = (void*)__b0;
	QWidget* __b1 = static_cast<QWidget*>(__b0); out[1] = (void*)__b1;
	QObject* __b2 = static_cast<QObject*>(__b1); out[2] = (void*)__b2;
}
void QSplitter__set_override(void* self, int idx) { ((QSplitter__Go*)self)->__goOn[idx] = 1; }
void* QSplitter_fromQObject(void* o) { return (void*)qobject_cast<QSplitter*>((QObject*)o); }
void* QSplitter_new(long long __goid, void* parent) {
	return new QSplitter__Go(__goid, (QWidget*)parent);
}
void* QSplitter_new2(long long __goid) {
	return new QSplitter__Go(__goid);
}
void* QSplitter_new3(long long __goid, int arg1) {
	return new QSplitter__Go(__goid, (Qt::Orientation)(arg1));
}
void* QSplitter_new4(long long __goid, int arg1, void* parent) {
	return new QSplitter__Go(__goid, (Qt::Orientation)(arg1), (QWidget*)parent);
}
char* QSplitter_tr(const char* s, long long* __out_len) {
	auto&& __r = QSplitter::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QSplitter_addWidget(void* self, void* widget) {
	((QSplitter*)self)->addWidget((QWidget*)widget);
}
void QSplitter_insertWidget(void* self, int index, void* widget) {
	((QSplitter*)self)->insertWidget(index, (QWidget*)widget);
}
void* QSplitter_replaceWidget(void* self, int index, void* widget) {
	auto&& __r = ((QSplitter*)self)->replaceWidget(index, (QWidget*)widget);
	return (void*)(__r);
}
void QSplitter_setOrientation(void* self, int arg1) {
	((QSplitter*)self)->setOrientation((Qt::Orientation)(arg1));
}
int QSplitter_orientation(void* self) {
	auto&& __r = ((QSplitter*)self)->orientation();
	return (int)(__r);
}
void QSplitter_setChildrenCollapsible(void* self, int arg1) {
	((QSplitter*)self)->setChildrenCollapsible(arg1);
}
int QSplitter_childrenCollapsible(void* self) {
	auto&& __r = ((QSplitter*)self)->childrenCollapsible();
	return __r;
}
void QSplitter_setCollapsible(void* self, int index, int arg2) {
	((QSplitter*)self)->setCollapsible(index, arg2);
}
int QSplitter_isCollapsible(void* self, int index) {
	auto&& __r = ((QSplitter*)self)->isCollapsible(index);
	return __r;
}
void QSplitter_setOpaqueResize(void* self) {
	((QSplitter*)self)->setOpaqueResize();
}
int QSplitter_opaqueResize(void* self) {
	auto&& __r = ((QSplitter*)self)->opaqueResize();
	return __r;
}
void QSplitter_refresh(void* self) {
	((QSplitter*)self)->refresh();
}
int* QSplitter_sizes(void* self, long long* __out_len) {
	auto&& __r = ((QSplitter*)self)->sizes();
	*__out_len = __r.size();
	int* __arr = (int*)malloc(sizeof(int) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = __e;
	return __arr;
}
void QSplitter_setSizes(void* self, int* list, long long list_len) {
	QList<int> list_list; list_list.reserve(list_len);
	for (long long __i = 0; __i < list_len; __i++) list_list.append(list[__i]);
	((QSplitter*)self)->setSizes(list_list);
}
void* QSplitter_saveState(void* self) {
	auto&& __r = ((QSplitter*)self)->saveState();
	return new QByteArray(__r);
}
int QSplitter_restoreState(void* self, void* state) {
	auto&& __r = ((QSplitter*)self)->restoreState(*(QByteArray*)state);
	return __r;
}
int QSplitter_handleWidth(void* self) {
	auto&& __r = ((QSplitter*)self)->handleWidth();
	return __r;
}
void QSplitter_setHandleWidth(void* self, int arg1) {
	((QSplitter*)self)->setHandleWidth(arg1);
}
int QSplitter_indexOf(void* self, void* w) {
	auto&& __r = ((QSplitter*)self)->indexOf((QWidget*)w);
	return __r;
}
void* QSplitter_widget(void* self, int index) {
	auto&& __r = ((QSplitter*)self)->widget(index);
	return (void*)(__r);
}
int QSplitter_count(void* self) {
	auto&& __r = ((QSplitter*)self)->count();
	return __r;
}
void QSplitter_getRange(void* self, int index, int* arg2, int* arg3) {
	int arg2_local{};
	int arg3_local{};
	((QSplitter*)self)->getRange(index, &arg2_local, &arg3_local);
	*arg2 = arg2_local;
	*arg3 = arg3_local;
}
void* QSplitter_handle(void* self, int index) {
	auto&& __r = ((QSplitter*)self)->handle(index);
	return (void*)(__r);
}
void QSplitter_setStretchFactor(void* self, int index, int stretch) {
	((QSplitter*)self)->setStretchFactor(index, stretch);
}
void QSplitter_splitterMoved(void* self, int pos, int index) {
	((QSplitter*)self)->splitterMoved(pos, index);
}
void QSplitter_moveSplitter(void* self, int pos, int index) {
	((QSplitter__Go*)self)->moveSplitter(pos, index);
}
void QSplitter_setRubberBand(void* self, int position) {
	((QSplitter__Go*)self)->setRubberBand(position);
}
int QSplitter_closestLegalPosition(void* self, int arg1, int arg2) {
	auto&& __r = ((QSplitter__Go*)self)->closestLegalPosition(arg1, arg2);
	return __r;
}
void QSplitter_drawFrame(void* self, void* arg1) {
	((QSplitter__Go*)self)->drawFrame((QPainter*)arg1);
}
void QSplitter_setTabOrder(void* arg1, void* arg2) {
	QSplitter::setTabOrder((QWidget*)arg1, (QWidget*)arg2);
}
void* QSplitter_mouseGrabber() {
	auto&& __r = QSplitter::mouseGrabber();
	return (void*)(__r);
}
void* QSplitter_keyboardGrabber() {
	auto&& __r = QSplitter::keyboardGrabber();
	return (void*)(__r);
}
void* QSplitter_find(unsigned long long arg1) {
	auto&& __r = QSplitter::find(arg1);
	return (void*)(__r);
}
void* QSplitter_createWindowContainer(void* window) {
	auto&& __r = QSplitter::createWindowContainer((QWindow*)window);
	return (void*)(__r);
}
void QSplitter_updateMicroFocus(void* self) {
	((QSplitter__Go*)self)->updateMicroFocus();
}
void QSplitter_create(void* self) {
	((QSplitter__Go*)self)->create();
}
void QSplitter_destroy(void* self) {
	((QSplitter__Go*)self)->destroy();
}
int QSplitter_focusNextChild(void* self) {
	auto&& __r = ((QSplitter__Go*)self)->focusNextChild();
	return __r;
}
int QSplitter_focusPreviousChild(void* self) {
	auto&& __r = ((QSplitter__Go*)self)->focusPreviousChild();
	return __r;
}
void* QSplitter_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QSplitter::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QSplitter_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QSplitter::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QSplitter_sender(void* self) {
	auto&& __r = ((QSplitter__Go*)self)->sender();
	return (void*)(__r);
}
int QSplitter_senderSignalIndex(void* self) {
	auto&& __r = ((QSplitter__Go*)self)->senderSignalIndex();
	return __r;
}
int QSplitter_receivers(void* self, const char* signal) {
	auto&& __r = ((QSplitter__Go*)self)->receivers(signal);
	return __r;
}
int QSplitter_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QSplitter__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
double QSplitter_devicePixelRatioFScale() {
	auto&& __r = QSplitter::devicePixelRatioFScale();
	return __r;
}
int QSplitter_encodeMetricF(int metric, double value) {
	auto&& __r = QSplitter::encodeMetricF((QPaintDevice::PaintDeviceMetric)(metric), value);
	return __r;
}
double QSplitter_getDecodedMetricF(void* self, int metricA, int metricB) {
	auto&& __r = ((QSplitter__Go*)self)->getDecodedMetricF((QPaintDevice::PaintDeviceMetric)(metricA), (QPaintDevice::PaintDeviceMetric)(metricB));
	return __r;
}
char* QSplitter_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QSplitter::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QSplitter_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QSplitter::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QSplitter_setOpaqueResize2(void* self, int opaque) {
	((QSplitter*)self)->setOpaqueResize(opaque);
}
void* QSplitter_createWindowContainer2(void* window, void* parent) {
	auto&& __r = QSplitter::createWindowContainer((QWindow*)window, (QWidget*)parent);
	return (void*)(__r);
}
void* QSplitter_createWindowContainer3(void* window, void* parent, int flags) {
	auto&& __r = QSplitter::createWindowContainer((QWindow*)window, (QWidget*)parent, QFlags<Qt::WindowType>(flags));
	return (void*)(__r);
}
void QSplitter_updateMicroFocus2(void* self, int query) {
	((QSplitter__Go*)self)->updateMicroFocus((Qt::InputMethodQuery)(query));
}
void QSplitter_create2(void* self, unsigned long long arg1) {
	((QSplitter__Go*)self)->create(arg1);
}
void QSplitter_create3(void* self, unsigned long long arg1, int initializeWindow) {
	((QSplitter__Go*)self)->create(arg1, initializeWindow);
}
void QSplitter_create4(void* self, unsigned long long arg1, int initializeWindow, int destroyOldWindow) {
	((QSplitter__Go*)self)->create(arg1, initializeWindow, destroyOldWindow);
}
void QSplitter_destroy2(void* self, int destroyWindow) {
	((QSplitter__Go*)self)->destroy(destroyWindow);
}
void QSplitter_destroy3(void* self, int destroyWindow, int destroySubWindows) {
	((QSplitter__Go*)self)->destroy(destroyWindow, destroySubWindows);
}
void* QSplitter_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QSplitter::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QSplitter_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QSplitter::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QSplitter_metaObject_default(void* self) {
	auto&& __r = ((QSplitter*)self)->QSplitter::metaObject();
	return (void*)(__r);
}
void* QSplitter_sizeHint_default(void* self) {
	auto&& __r = ((QSplitter*)self)->QSplitter::sizeHint();
	return new QSize(__r);
}
void* QSplitter_minimumSizeHint_default(void* self) {
	auto&& __r = ((QSplitter*)self)->QSplitter::minimumSizeHint();
	return new QSize(__r);
}
void* QSplitter_createHandle_default(void* self) {
	auto&& __r = ((QSplitter__Go*)self)->cppgoBase_createHandle();
	return (void*)(__r);
}
void QSplitter_childEvent_default(void* self, void* arg1) {
	((QSplitter__Go*)self)->cppgoBase_childEvent((QChildEvent*)arg1);
}
int QSplitter_event_default(void* self, void* arg1) {
	auto&& __r = ((QSplitter__Go*)self)->cppgoBase_event((QEvent*)arg1);
	return __r;
}
void QSplitter_resizeEvent_default(void* self, void* arg1) {
	((QSplitter__Go*)self)->cppgoBase_resizeEvent((QResizeEvent*)arg1);
}
void QSplitter_changeEvent_default(void* self, void* arg1) {
	((QSplitter__Go*)self)->cppgoBase_changeEvent((QEvent*)arg1);
}
void QSplitter_paintEvent_default(void* self, void* arg1) {
	((QSplitter__Go*)self)->cppgoBase_paintEvent((QPaintEvent*)arg1);
}
void QSplitter_initStyleOption_default(void* self, void* option) {
	((QSplitter__Go*)self)->cppgoBase_initStyleOption((QStyleOptionFrame*)option);
}
int QSplitter_devType_default(void* self) {
	auto&& __r = ((QSplitter*)self)->QSplitter::devType();
	return __r;
}
void QSplitter_setVisible_default(void* self, int visible) {
	((QSplitter*)self)->QSplitter::setVisible(visible);
}
int QSplitter_heightForWidth_default(void* self, int arg1) {
	auto&& __r = ((QSplitter*)self)->QSplitter::heightForWidth(arg1);
	return __r;
}
int QSplitter_hasHeightForWidth_default(void* self) {
	auto&& __r = ((QSplitter*)self)->QSplitter::hasHeightForWidth();
	return __r;
}
void* QSplitter_paintEngine_default(void* self) {
	auto&& __r = ((QSplitter*)self)->QSplitter::paintEngine();
	return (void*)(__r);
}
void QSplitter_mousePressEvent_default(void* self, void* event) {
	((QSplitter__Go*)self)->cppgoBase_mousePressEvent((QMouseEvent*)event);
}
void QSplitter_mouseReleaseEvent_default(void* self, void* event) {
	((QSplitter__Go*)self)->cppgoBase_mouseReleaseEvent((QMouseEvent*)event);
}
void QSplitter_mouseDoubleClickEvent_default(void* self, void* event) {
	((QSplitter__Go*)self)->cppgoBase_mouseDoubleClickEvent((QMouseEvent*)event);
}
void QSplitter_mouseMoveEvent_default(void* self, void* event) {
	((QSplitter__Go*)self)->cppgoBase_mouseMoveEvent((QMouseEvent*)event);
}
void QSplitter_wheelEvent_default(void* self, void* event) {
	((QSplitter__Go*)self)->cppgoBase_wheelEvent((QWheelEvent*)event);
}
void QSplitter_keyPressEvent_default(void* self, void* event) {
	((QSplitter__Go*)self)->cppgoBase_keyPressEvent((QKeyEvent*)event);
}
void QSplitter_keyReleaseEvent_default(void* self, void* event) {
	((QSplitter__Go*)self)->cppgoBase_keyReleaseEvent((QKeyEvent*)event);
}
void QSplitter_focusInEvent_default(void* self, void* event) {
	((QSplitter__Go*)self)->cppgoBase_focusInEvent((QFocusEvent*)event);
}
void QSplitter_focusOutEvent_default(void* self, void* event) {
	((QSplitter__Go*)self)->cppgoBase_focusOutEvent((QFocusEvent*)event);
}
void QSplitter_enterEvent_default(void* self, void* event) {
	((QSplitter__Go*)self)->cppgoBase_enterEvent((QEnterEvent*)event);
}
void QSplitter_leaveEvent_default(void* self, void* event) {
	((QSplitter__Go*)self)->cppgoBase_leaveEvent((QEvent*)event);
}
void QSplitter_moveEvent_default(void* self, void* event) {
	((QSplitter__Go*)self)->cppgoBase_moveEvent((QMoveEvent*)event);
}
void QSplitter_closeEvent_default(void* self, void* event) {
	((QSplitter__Go*)self)->cppgoBase_closeEvent((QCloseEvent*)event);
}
void QSplitter_contextMenuEvent_default(void* self, void* event) {
	((QSplitter__Go*)self)->cppgoBase_contextMenuEvent((QContextMenuEvent*)event);
}
void QSplitter_tabletEvent_default(void* self, void* event) {
	((QSplitter__Go*)self)->cppgoBase_tabletEvent((QTabletEvent*)event);
}
void QSplitter_actionEvent_default(void* self, void* event) {
	((QSplitter__Go*)self)->cppgoBase_actionEvent((QActionEvent*)event);
}
void QSplitter_dragEnterEvent_default(void* self, void* event) {
	((QSplitter__Go*)self)->cppgoBase_dragEnterEvent((QDragEnterEvent*)event);
}
void QSplitter_dragMoveEvent_default(void* self, void* event) {
	((QSplitter__Go*)self)->cppgoBase_dragMoveEvent((QDragMoveEvent*)event);
}
void QSplitter_dragLeaveEvent_default(void* self, void* event) {
	((QSplitter__Go*)self)->cppgoBase_dragLeaveEvent((QDragLeaveEvent*)event);
}
void QSplitter_dropEvent_default(void* self, void* event) {
	((QSplitter__Go*)self)->cppgoBase_dropEvent((QDropEvent*)event);
}
void QSplitter_showEvent_default(void* self, void* event) {
	((QSplitter__Go*)self)->cppgoBase_showEvent((QShowEvent*)event);
}
void QSplitter_hideEvent_default(void* self, void* event) {
	((QSplitter__Go*)self)->cppgoBase_hideEvent((QHideEvent*)event);
}
int QSplitter_metric_default(void* self, int arg1) {
	auto&& __r = ((QSplitter__Go*)self)->cppgoBase_metric(arg1);
	return __r;
}
void QSplitter_initPainter_default(void* self, void* painter) {
	((QSplitter__Go*)self)->cppgoBase_initPainter((QPainter*)painter);
}
void* QSplitter_redirected_default(void* self, void* offset) {
	auto&& __r = ((QSplitter__Go*)self)->cppgoBase_redirected((QPoint*)offset);
	return (void*)(__r);
}
void* QSplitter_sharedPainter_default(void* self) {
	auto&& __r = ((QSplitter__Go*)self)->cppgoBase_sharedPainter();
	return (void*)(__r);
}
void QSplitter_inputMethodEvent_default(void* self, void* arg1) {
	((QSplitter__Go*)self)->cppgoBase_inputMethodEvent((QInputMethodEvent*)arg1);
}
void* QSplitter_inputMethodQuery_default(void* self, int arg1) {
	auto&& __r = ((QSplitter*)self)->QSplitter::inputMethodQuery((Qt::InputMethodQuery)(arg1));
	return new QVariant(__r);
}
int QSplitter_focusNextPrevChild_default(void* self, int next) {
	auto&& __r = ((QSplitter__Go*)self)->cppgoBase_focusNextPrevChild(next);
	return __r;
}
int QSplitter_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QSplitter*)self)->QSplitter::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QSplitter_timerEvent_default(void* self, void* event) {
	((QSplitter__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QSplitter_customEvent_default(void* self, void* event) {
	((QSplitter__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QSplitter_connectNotify_default(void* self, void* signal) {
	((QSplitter__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QSplitter_disconnectNotify_default(void* self, void* signal) {
	((QSplitter__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QSplitter_connect_splitterMoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSplitter*)self, &QSplitter::splitterMoved, (QSplitter*)self, [__id](int pos, int index){
		if (cppgo_sig_void_int_int) ((void(*)(long long, int, int))cppgo_sig_void_int_int)(__id, pos, index);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSplitter);
	__conns_QSplitter[__id] = __c;
}
void QSplitter_connect_windowTitleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSplitter*)self, &QSplitter::windowTitleChanged, (QSplitter*)self, [__id](const QString& title){
		QByteArray title_ba = title.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)title_ba.constData(), title_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSplitter);
	__conns_QSplitter[__id] = __c;
}
void QSplitter_connect_windowIconChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSplitter*)self, &QSplitter::windowIconChanged, (QSplitter*)self, [__id](const QIcon& icon){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&icon);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSplitter);
	__conns_QSplitter[__id] = __c;
}
void QSplitter_connect_windowIconTextChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSplitter*)self, &QSplitter::windowIconTextChanged, (QSplitter*)self, [__id](const QString& iconText){
		QByteArray iconText_ba = iconText.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)iconText_ba.constData(), iconText_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSplitter);
	__conns_QSplitter[__id] = __c;
}
void QSplitter_connect_customContextMenuRequested(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSplitter*)self, &QSplitter::customContextMenuRequested, (QSplitter*)self, [__id](const QPoint& pos){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&pos);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSplitter);
	__conns_QSplitter[__id] = __c;
}
void QSplitter_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSplitter*)self, &QSplitter::destroyed, (QSplitter*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSplitter);
	__conns_QSplitter[__id] = __c;
}
void QSplitter_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSplitter*)self, &QSplitter::destroyed, (QSplitter*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSplitter);
	__conns_QSplitter[__id] = __c;
}
void QSplitter_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSplitter*)self, &QSplitter::objectNameChanged, (QSplitter*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSplitter);
	__conns_QSplitter[__id] = __c;
}
void QSplitter__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QSplitter);
		auto __it = __conns_QSplitter.find(id);
		if (__it == __conns_QSplitter.end()) return;
		__c = __it->second;
		__conns_QSplitter.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QSplitter__Go_new(long long id) { return new QSplitter__Go(id); }
void QSplitter_delete(void* self) { delete (QSplitter*)self; }

}

#include <QtWidgets/qwhatsthis.h>
#include <QtCore/qpoint.h>
#include <QtCore/qstring.h>
#include <QtWidgets/qwidget.h>
#include <QtGui/qaction.h>
#include <QtCore/qobject.h>
#include "gen_qwhatsthis.h"

extern "C" {

void QWhatsThis_enterWhatsThisMode() {
	QWhatsThis::enterWhatsThisMode();
}
int QWhatsThis_inWhatsThisMode() {
	auto&& __r = QWhatsThis::inWhatsThisMode();
	return __r;
}
void QWhatsThis_leaveWhatsThisMode() {
	QWhatsThis::leaveWhatsThisMode();
}
void QWhatsThis_showText(void* pos, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	QWhatsThis::showText(*(QPoint*)pos, text_qstr);
}
void QWhatsThis_hideText() {
	QWhatsThis::hideText();
}
void* QWhatsThis_createAction() {
	auto&& __r = QWhatsThis::createAction();
	return (void*)(__r);
}
void QWhatsThis_showText2(void* pos, const char* text, long long text_len, void* w) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	QWhatsThis::showText(*(QPoint*)pos, text_qstr, (QWidget*)w);
}
void* QWhatsThis_createAction2(void* parent) {
	auto&& __r = QWhatsThis::createAction((QObject*)parent);
	return (void*)(__r);
}
void QWhatsThis_delete(void* self) { delete (QWhatsThis*)self; }

}

