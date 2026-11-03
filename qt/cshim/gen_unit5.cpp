#include <QtCore/qatomic.h>
#include "gen_qatomicint.h"

extern "C" {

void* QAtomicInt_new() {
	return new QAtomicInt();
}
void* QAtomicInt_new2(void* arg1) {
	return new QAtomicInt(*(QAtomicInt*)arg1);
}
void* QAtomicInt_new3(int value) {
	return new QAtomicInt(value);
}
void* QAtomicInt_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QAtomicInt*)self)->operator=(*(QAtomicInt*)arg1);
	return (void*)&(__r);
}
int QAtomicInt_loadRelaxed(void* self) {
	auto&& __r = ((QAtomicInt*)self)->loadRelaxed();
	return __r;
}
void QAtomicInt_storeRelaxed(void* self, int newValue) {
	((QAtomicInt*)self)->storeRelaxed(newValue);
}
int QAtomicInt_loadAcquire(void* self) {
	auto&& __r = ((QAtomicInt*)self)->loadAcquire();
	return __r;
}
void QAtomicInt_storeRelease(void* self, int newValue) {
	((QAtomicInt*)self)->storeRelease(newValue);
}
int QAtomicInt_ToInt(void* self) {
	auto&& __r = static_cast<int>(*(QAtomicInt*)self);
	return __r;
}
int QAtomicInt_isReferenceCountingNative() {
	auto&& __r = QAtomicInt::isReferenceCountingNative();
	return __r;
}
int QAtomicInt_isReferenceCountingWaitFree() {
	auto&& __r = QAtomicInt::isReferenceCountingWaitFree();
	return __r;
}
int QAtomicInt_ref(void* self) {
	auto&& __r = ((QAtomicInt*)self)->ref();
	return __r;
}
void QAtomicInt_refRelaxed(void* self) {
	((QAtomicInt*)self)->refRelaxed();
}
int QAtomicInt_deref(void* self) {
	auto&& __r = ((QAtomicInt*)self)->deref();
	return __r;
}
int QAtomicInt_isTestAndSetNative() {
	auto&& __r = QAtomicInt::isTestAndSetNative();
	return __r;
}
int QAtomicInt_isTestAndSetWaitFree() {
	auto&& __r = QAtomicInt::isTestAndSetWaitFree();
	return __r;
}
int QAtomicInt_testAndSetRelaxed(void* self, int expectedValue, int newValue) {
	auto&& __r = ((QAtomicInt*)self)->testAndSetRelaxed(expectedValue, newValue);
	return __r;
}
int QAtomicInt_testAndSetAcquire(void* self, int expectedValue, int newValue) {
	auto&& __r = ((QAtomicInt*)self)->testAndSetAcquire(expectedValue, newValue);
	return __r;
}
int QAtomicInt_testAndSetRelease(void* self, int expectedValue, int newValue) {
	auto&& __r = ((QAtomicInt*)self)->testAndSetRelease(expectedValue, newValue);
	return __r;
}
int QAtomicInt_testAndSetOrdered(void* self, int expectedValue, int newValue) {
	auto&& __r = ((QAtomicInt*)self)->testAndSetOrdered(expectedValue, newValue);
	return __r;
}
int QAtomicInt_testAndSetRelaxed2(void* self, int expectedValue, int newValue, int* currentValue) {
	int currentValue_ref = *currentValue;
	auto&& __r = ((QAtomicInt*)self)->testAndSetRelaxed(expectedValue, newValue, currentValue_ref);
	*currentValue = currentValue_ref;
	return __r;
}
int QAtomicInt_testAndSetAcquire2(void* self, int expectedValue, int newValue, int* currentValue) {
	int currentValue_ref = *currentValue;
	auto&& __r = ((QAtomicInt*)self)->testAndSetAcquire(expectedValue, newValue, currentValue_ref);
	*currentValue = currentValue_ref;
	return __r;
}
int QAtomicInt_testAndSetRelease2(void* self, int expectedValue, int newValue, int* currentValue) {
	int currentValue_ref = *currentValue;
	auto&& __r = ((QAtomicInt*)self)->testAndSetRelease(expectedValue, newValue, currentValue_ref);
	*currentValue = currentValue_ref;
	return __r;
}
int QAtomicInt_testAndSetOrdered2(void* self, int expectedValue, int newValue, int* currentValue) {
	int currentValue_ref = *currentValue;
	auto&& __r = ((QAtomicInt*)self)->testAndSetOrdered(expectedValue, newValue, currentValue_ref);
	*currentValue = currentValue_ref;
	return __r;
}
int QAtomicInt_isFetchAndStoreNative() {
	auto&& __r = QAtomicInt::isFetchAndStoreNative();
	return __r;
}
int QAtomicInt_isFetchAndStoreWaitFree() {
	auto&& __r = QAtomicInt::isFetchAndStoreWaitFree();
	return __r;
}
int QAtomicInt_fetchAndStoreRelaxed(void* self, int newValue) {
	auto&& __r = ((QAtomicInt*)self)->fetchAndStoreRelaxed(newValue);
	return __r;
}
int QAtomicInt_fetchAndStoreAcquire(void* self, int newValue) {
	auto&& __r = ((QAtomicInt*)self)->fetchAndStoreAcquire(newValue);
	return __r;
}
int QAtomicInt_fetchAndStoreRelease(void* self, int newValue) {
	auto&& __r = ((QAtomicInt*)self)->fetchAndStoreRelease(newValue);
	return __r;
}
int QAtomicInt_fetchAndStoreOrdered(void* self, int newValue) {
	auto&& __r = ((QAtomicInt*)self)->fetchAndStoreOrdered(newValue);
	return __r;
}
int QAtomicInt_isFetchAndAddNative() {
	auto&& __r = QAtomicInt::isFetchAndAddNative();
	return __r;
}
int QAtomicInt_isFetchAndAddWaitFree() {
	auto&& __r = QAtomicInt::isFetchAndAddWaitFree();
	return __r;
}
int QAtomicInt_fetchAndAddRelaxed(void* self, int valueToAdd) {
	auto&& __r = ((QAtomicInt*)self)->fetchAndAddRelaxed(valueToAdd);
	return __r;
}
int QAtomicInt_fetchAndAddAcquire(void* self, int valueToAdd) {
	auto&& __r = ((QAtomicInt*)self)->fetchAndAddAcquire(valueToAdd);
	return __r;
}
int QAtomicInt_fetchAndAddRelease(void* self, int valueToAdd) {
	auto&& __r = ((QAtomicInt*)self)->fetchAndAddRelease(valueToAdd);
	return __r;
}
int QAtomicInt_fetchAndAddOrdered(void* self, int valueToAdd) {
	auto&& __r = ((QAtomicInt*)self)->fetchAndAddOrdered(valueToAdd);
	return __r;
}
int QAtomicInt_fetchAndSubRelaxed(void* self, int valueToAdd) {
	auto&& __r = ((QAtomicInt*)self)->fetchAndSubRelaxed(valueToAdd);
	return __r;
}
int QAtomicInt_fetchAndSubAcquire(void* self, int valueToAdd) {
	auto&& __r = ((QAtomicInt*)self)->fetchAndSubAcquire(valueToAdd);
	return __r;
}
int QAtomicInt_fetchAndSubRelease(void* self, int valueToAdd) {
	auto&& __r = ((QAtomicInt*)self)->fetchAndSubRelease(valueToAdd);
	return __r;
}
int QAtomicInt_fetchAndSubOrdered(void* self, int valueToAdd) {
	auto&& __r = ((QAtomicInt*)self)->fetchAndSubOrdered(valueToAdd);
	return __r;
}
int QAtomicInt_fetchAndAndRelaxed(void* self, int valueToAdd) {
	auto&& __r = ((QAtomicInt*)self)->fetchAndAndRelaxed(valueToAdd);
	return __r;
}
int QAtomicInt_fetchAndAndAcquire(void* self, int valueToAdd) {
	auto&& __r = ((QAtomicInt*)self)->fetchAndAndAcquire(valueToAdd);
	return __r;
}
int QAtomicInt_fetchAndAndRelease(void* self, int valueToAdd) {
	auto&& __r = ((QAtomicInt*)self)->fetchAndAndRelease(valueToAdd);
	return __r;
}
int QAtomicInt_fetchAndAndOrdered(void* self, int valueToAdd) {
	auto&& __r = ((QAtomicInt*)self)->fetchAndAndOrdered(valueToAdd);
	return __r;
}
int QAtomicInt_fetchAndOrRelaxed(void* self, int valueToAdd) {
	auto&& __r = ((QAtomicInt*)self)->fetchAndOrRelaxed(valueToAdd);
	return __r;
}
int QAtomicInt_fetchAndOrAcquire(void* self, int valueToAdd) {
	auto&& __r = ((QAtomicInt*)self)->fetchAndOrAcquire(valueToAdd);
	return __r;
}
int QAtomicInt_fetchAndOrRelease(void* self, int valueToAdd) {
	auto&& __r = ((QAtomicInt*)self)->fetchAndOrRelease(valueToAdd);
	return __r;
}
int QAtomicInt_fetchAndOrOrdered(void* self, int valueToAdd) {
	auto&& __r = ((QAtomicInt*)self)->fetchAndOrOrdered(valueToAdd);
	return __r;
}
int QAtomicInt_fetchAndXorRelaxed(void* self, int valueToAdd) {
	auto&& __r = ((QAtomicInt*)self)->fetchAndXorRelaxed(valueToAdd);
	return __r;
}
int QAtomicInt_fetchAndXorAcquire(void* self, int valueToAdd) {
	auto&& __r = ((QAtomicInt*)self)->fetchAndXorAcquire(valueToAdd);
	return __r;
}
int QAtomicInt_fetchAndXorRelease(void* self, int valueToAdd) {
	auto&& __r = ((QAtomicInt*)self)->fetchAndXorRelease(valueToAdd);
	return __r;
}
int QAtomicInt_fetchAndXorOrdered(void* self, int valueToAdd) {
	auto&& __r = ((QAtomicInt*)self)->fetchAndXorOrdered(valueToAdd);
	return __r;
}
int QAtomicInt_OperatorPlusPlus(void* self) {
	auto&& __r = ((QAtomicInt*)self)->operator++();
	return __r;
}
int QAtomicInt_OperatorPlusPlus2(void* self, int arg1) {
	auto&& __r = ((QAtomicInt*)self)->operator++(arg1);
	return __r;
}
int QAtomicInt_OperatorMinusMinus(void* self) {
	auto&& __r = ((QAtomicInt*)self)->operator--();
	return __r;
}
int QAtomicInt_OperatorMinusMinus2(void* self, int arg1) {
	auto&& __r = ((QAtomicInt*)self)->operator--(arg1);
	return __r;
}
int QAtomicInt_OperatorPlusAssign(void* self, int v) {
	auto&& __r = ((QAtomicInt*)self)->operator+=(v);
	return __r;
}
int QAtomicInt_OperatorMinusAssign(void* self, int v) {
	auto&& __r = ((QAtomicInt*)self)->operator-=(v);
	return __r;
}
int QAtomicInt_OperatorBitwiseAndAssign(void* self, int v) {
	auto&& __r = ((QAtomicInt*)self)->operator&=(v);
	return __r;
}
int QAtomicInt_OperatorBitwiseOrAssign(void* self, int v) {
	auto&& __r = ((QAtomicInt*)self)->operator|=(v);
	return __r;
}
int QAtomicInt_OperatorBitwiseXorAssign(void* self, int v) {
	auto&& __r = ((QAtomicInt*)self)->operator^=(v);
	return __r;
}
void QAtomicInt_delete(void* self) { delete (QAtomicInt*)self; }

}

#include <QtCore/qstring.h>
#include <QtCore/qchar.h>
#include <QtCore/qlatin1stringview.h>
#include <QtCore/qstringview.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qregularexpression.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qbytearrayview.h>
#include "gen_qstring.h"

extern "C" {

void* QString_new() {
	return new QString();
}
void* QString_new2(void* unicode) {
	return new QString((QChar*)unicode);
}
void* QString_new3(void* c) {
	return new QString(*(QChar*)c);
}
void* QString_new4(long long size, void* c) {
	return new QString(size, *(QChar*)c);
}
void* QString_new5(void* latin1) {
	return new QString(*(QLatin1String*)latin1);
}
void* QString_new6(void* sv) {
	return new QString(*(QStringView*)sv);
}
void* QString_new7(const char* arg1, long long arg1_len) {
	QString arg1_qstr = QString::fromUtf8(arg1, arg1_len);
	return new QString(arg1_qstr);
}
void* QString_new8(const char* ch) {
	return new QString(ch);
}
void* QString_new9(void* a) {
	return new QString(*(QByteArray*)a);
}
void* QString_new10(long long size, int arg2) {
	return new QString(size, (Qt::Initialization)(arg2));
}
void* QString_new11(void* unicode, long long size) {
	return new QString((QChar*)unicode, size);
}
void* QString_OperatorAssign(void* self, void* c) {
	auto&& __r = ((QString*)self)->operator=(*(QChar*)c);
	return (void*)&(__r);
}
void* QString_OperatorAssign2(void* self, const char* arg1, long long arg1_len) {
	QString arg1_qstr = QString::fromUtf8(arg1, arg1_len);
	auto&& __r = ((QString*)self)->operator=(arg1_qstr);
	return (void*)&(__r);
}
void* QString_OperatorAssign3(void* self, void* latin1) {
	auto&& __r = ((QString*)self)->operator=(*(QLatin1String*)latin1);
	return (void*)&(__r);
}
void QString_swap(void* self, void* other) {
	((QString*)self)->swap(*(QString*)other);
}
long long QString_maxSize() {
	auto&& __r = QString::maxSize();
	return __r;
}
long long QString_size(void* self) {
	auto&& __r = ((QString*)self)->size();
	return __r;
}
long long QString_count(void* self) {
	auto&& __r = ((QString*)self)->count();
	return __r;
}
long long QString_length(void* self) {
	auto&& __r = ((QString*)self)->length();
	return __r;
}
int QString_isEmpty(void* self) {
	auto&& __r = ((QString*)self)->isEmpty();
	return __r;
}
void QString_resize(void* self, long long size) {
	((QString*)self)->resize(size);
}
void QString_resize2(void* self, long long size, void* fillChar) {
	((QString*)self)->resize(size, *(QChar*)fillChar);
}
void QString_resizeForOverwrite(void* self, long long size) {
	((QString*)self)->resizeForOverwrite(size);
}
void* QString_fill(void* self, void* c) {
	auto&& __r = ((QString*)self)->fill(*(QChar*)c);
	return (void*)&(__r);
}
void QString_truncate(void* self, long long pos) {
	((QString*)self)->truncate(pos);
}
void QString_chop(void* self, long long n) {
	((QString*)self)->chop(n);
}
void* QString_slice(void* self, long long pos) {
	auto&& __r = ((QString*)self)->slice(pos);
	return (void*)&(__r);
}
void* QString_slice2(void* self, long long pos, long long n) {
	auto&& __r = ((QString*)self)->slice(pos, n);
	return (void*)&(__r);
}
long long QString_capacity(void* self) {
	auto&& __r = ((QString*)self)->capacity();
	return __r;
}
void QString_reserve(void* self, long long size) {
	((QString*)self)->reserve(size);
}
void QString_squeeze(void* self) {
	((QString*)self)->squeeze();
}
void* QString_unicode(void* self) {
	auto&& __r = ((QString*)self)->unicode();
	return (void*)(__r);
}
void* QString_data(void* self) {
	auto&& __r = ((QString*)self)->data();
	return (void*)(__r);
}
void* QString_data2(void* self) {
	auto&& __r = ((QString*)self)->data();
	return (void*)(__r);
}
void* QString_constData(void* self) {
	auto&& __r = ((QString*)self)->constData();
	return (void*)(__r);
}
void QString_detach(void* self) {
	((QString*)self)->detach();
}
int QString_isDetached(void* self) {
	auto&& __r = ((QString*)self)->isDetached();
	return __r;
}
int QString_isSharedWith(void* self, const char* other, long long other_len) {
	QString other_qstr = QString::fromUtf8(other, other_len);
	auto&& __r = ((QString*)self)->isSharedWith(other_qstr);
	return __r;
}
void QString_clear(void* self) {
	((QString*)self)->clear();
}
void* QString_at(void* self, long long i) {
	auto&& __r = ((QString*)self)->at(i);
	return new QChar(__r);
}
void* QString_OperatorSubscript(void* self, long long i) {
	auto&& __r = ((QString*)self)->operator[](i);
	return new QChar(__r);
}
void* QString_OperatorSubscript2(void* self, long long i) {
	auto&& __r = ((QString*)self)->operator[](i);
	return (void*)&(__r);
}
void* QString_front(void* self) {
	auto&& __r = ((QString*)self)->front();
	return new QChar(__r);
}
void* QString_front2(void* self) {
	auto&& __r = ((QString*)self)->front();
	return (void*)&(__r);
}
void* QString_back(void* self) {
	auto&& __r = ((QString*)self)->back();
	return new QChar(__r);
}
void* QString_back2(void* self) {
	auto&& __r = ((QString*)self)->back();
	return (void*)&(__r);
}
char* QString_asprintf(const char* format, long long* __out_len) {
	auto&& __r = QString::asprintf(format);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
long long QString_indexOf(void* self, void* c) {
	auto&& __r = ((QString*)self)->indexOf(*(QChar*)c);
	return __r;
}
long long QString_indexOf2(void* self, void* s) {
	auto&& __r = ((QString*)self)->indexOf(*(QLatin1String*)s);
	return __r;
}
long long QString_indexOf3(void* self, const char* s, long long s_len) {
	QString s_qstr = QString::fromUtf8(s, s_len);
	auto&& __r = ((QString*)self)->indexOf(s_qstr);
	return __r;
}
long long QString_indexOf4(void* self, void* s) {
	auto&& __r = ((QString*)self)->indexOf(*(QStringView*)s);
	return __r;
}
long long QString_lastIndexOf(void* self, void* c) {
	auto&& __r = ((QString*)self)->lastIndexOf(*(QChar*)c);
	return __r;
}
long long QString_lastIndexOf2(void* self, void* c, long long from) {
	auto&& __r = ((QString*)self)->lastIndexOf(*(QChar*)c, from);
	return __r;
}
long long QString_lastIndexOf3(void* self, void* s) {
	auto&& __r = ((QString*)self)->lastIndexOf(*(QLatin1String*)s);
	return __r;
}
long long QString_lastIndexOf4(void* self, void* s, long long from) {
	auto&& __r = ((QString*)self)->lastIndexOf(*(QLatin1String*)s, from);
	return __r;
}
long long QString_lastIndexOf5(void* self, const char* s, long long s_len) {
	QString s_qstr = QString::fromUtf8(s, s_len);
	auto&& __r = ((QString*)self)->lastIndexOf(s_qstr);
	return __r;
}
long long QString_lastIndexOf6(void* self, const char* s, long long s_len, long long from) {
	QString s_qstr = QString::fromUtf8(s, s_len);
	auto&& __r = ((QString*)self)->lastIndexOf(s_qstr, from);
	return __r;
}
long long QString_lastIndexOf7(void* self, void* s) {
	auto&& __r = ((QString*)self)->lastIndexOf(*(QStringView*)s);
	return __r;
}
long long QString_lastIndexOf8(void* self, void* s, long long from) {
	auto&& __r = ((QString*)self)->lastIndexOf(*(QStringView*)s, from);
	return __r;
}
int QString_contains(void* self, void* c) {
	auto&& __r = ((QString*)self)->contains(*(QChar*)c);
	return __r;
}
int QString_contains2(void* self, const char* s, long long s_len) {
	QString s_qstr = QString::fromUtf8(s, s_len);
	auto&& __r = ((QString*)self)->contains(s_qstr);
	return __r;
}
int QString_contains3(void* self, void* s) {
	auto&& __r = ((QString*)self)->contains(*(QLatin1String*)s);
	return __r;
}
int QString_contains4(void* self, void* s) {
	auto&& __r = ((QString*)self)->contains(*(QStringView*)s);
	return __r;
}
long long QString_count2(void* self, void* c) {
	auto&& __r = ((QString*)self)->count(*(QChar*)c);
	return __r;
}
long long QString_count3(void* self, const char* s, long long s_len) {
	QString s_qstr = QString::fromUtf8(s, s_len);
	auto&& __r = ((QString*)self)->count(s_qstr);
	return __r;
}
long long QString_count4(void* self, void* s) {
	auto&& __r = ((QString*)self)->count(*(QStringView*)s);
	return __r;
}
long long QString_indexOf5(void* self, void* re) {
	auto&& __r = ((QString*)self)->indexOf(*(QRegularExpression*)re);
	return __r;
}
long long QString_lastIndexOf9(void* self, void* re, long long from) {
	auto&& __r = ((QString*)self)->lastIndexOf(*(QRegularExpression*)re, from);
	return __r;
}
int QString_contains5(void* self, void* re) {
	auto&& __r = ((QString*)self)->contains(*(QRegularExpression*)re);
	return __r;
}
long long QString_count5(void* self, void* re) {
	auto&& __r = ((QString*)self)->count(*(QRegularExpression*)re);
	return __r;
}
char* QString_section(void* self, void* sep, long long start, long long* __out_len) {
	auto&& __r = ((QString*)self)->section(*(QChar*)sep, start);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QString_section2(void* self, const char* in_sep, long long in_sep_len, long long start, long long* __out_len) {
	QString in_sep_qstr = QString::fromUtf8(in_sep, in_sep_len);
	auto&& __r = ((QString*)self)->section(in_sep_qstr, start);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QString_section3(void* self, void* re, long long start, long long* __out_len) {
	auto&& __r = ((QString*)self)->section(*(QRegularExpression*)re, start);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QString_left(void* self, long long n, long long* __out_len) {
	auto&& __r = ((QString*)self)->left(n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QString_left2(void* self, long long n, long long* __out_len) {
	auto&& __r = ((QString*)self)->left(n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QString_right(void* self, long long n, long long* __out_len) {
	auto&& __r = ((QString*)self)->right(n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QString_right2(void* self, long long n, long long* __out_len) {
	auto&& __r = ((QString*)self)->right(n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QString_mid(void* self, long long position, long long* __out_len) {
	auto&& __r = ((QString*)self)->mid(position);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QString_mid2(void* self, long long position, long long* __out_len) {
	auto&& __r = ((QString*)self)->mid(position);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QString_first(void* self, long long n, long long* __out_len) {
	auto&& __r = ((QString*)self)->first(n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QString_last(void* self, long long n, long long* __out_len) {
	auto&& __r = ((QString*)self)->last(n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QString_sliced(void* self, long long pos, long long* __out_len) {
	auto&& __r = ((QString*)self)->sliced(pos);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QString_sliced2(void* self, long long pos, long long n, long long* __out_len) {
	auto&& __r = ((QString*)self)->sliced(pos, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QString_chopped(void* self, long long n, long long* __out_len) {
	auto&& __r = ((QString*)self)->chopped(n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QString_first2(void* self, long long n, long long* __out_len) {
	auto&& __r = ((QString*)self)->first(n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QString_last2(void* self, long long n, long long* __out_len) {
	auto&& __r = ((QString*)self)->last(n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QString_sliced3(void* self, long long pos, long long* __out_len) {
	auto&& __r = ((QString*)self)->sliced(pos);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QString_sliced4(void* self, long long pos, long long n, long long* __out_len) {
	auto&& __r = ((QString*)self)->sliced(pos, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QString_chopped2(void* self, long long n, long long* __out_len) {
	auto&& __r = ((QString*)self)->chopped(n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QString_startsWith(void* self, const char* s, long long s_len) {
	QString s_qstr = QString::fromUtf8(s, s_len);
	auto&& __r = ((QString*)self)->startsWith(s_qstr);
	return __r;
}
int QString_startsWith2(void* self, void* s) {
	auto&& __r = ((QString*)self)->startsWith(*(QStringView*)s);
	return __r;
}
int QString_startsWith3(void* self, void* s) {
	auto&& __r = ((QString*)self)->startsWith(*(QLatin1String*)s);
	return __r;
}
int QString_startsWith4(void* self, void* c) {
	auto&& __r = ((QString*)self)->startsWith(*(QChar*)c);
	return __r;
}
int QString_endsWith(void* self, const char* s, long long s_len) {
	QString s_qstr = QString::fromUtf8(s, s_len);
	auto&& __r = ((QString*)self)->endsWith(s_qstr);
	return __r;
}
int QString_endsWith2(void* self, void* s) {
	auto&& __r = ((QString*)self)->endsWith(*(QStringView*)s);
	return __r;
}
int QString_endsWith3(void* self, void* s) {
	auto&& __r = ((QString*)self)->endsWith(*(QLatin1String*)s);
	return __r;
}
int QString_endsWith4(void* self, void* c) {
	auto&& __r = ((QString*)self)->endsWith(*(QChar*)c);
	return __r;
}
int QString_isUpper(void* self) {
	auto&& __r = ((QString*)self)->isUpper();
	return __r;
}
int QString_isLower(void* self) {
	auto&& __r = ((QString*)self)->isLower();
	return __r;
}
char* QString_leftJustified(void* self, long long width, long long* __out_len) {
	auto&& __r = ((QString*)self)->leftJustified(width);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QString_rightJustified(void* self, long long width, long long* __out_len) {
	auto&& __r = ((QString*)self)->rightJustified(width);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QString_toLower(void* self, long long* __out_len) {
	auto&& __r = ((QString*)self)->toLower();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QString_toLower2(void* self, long long* __out_len) {
	auto&& __r = ((QString*)self)->toLower();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QString_toUpper(void* self, long long* __out_len) {
	auto&& __r = ((QString*)self)->toUpper();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QString_toUpper2(void* self, long long* __out_len) {
	auto&& __r = ((QString*)self)->toUpper();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QString_toCaseFolded(void* self, long long* __out_len) {
	auto&& __r = ((QString*)self)->toCaseFolded();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QString_toCaseFolded2(void* self, long long* __out_len) {
	auto&& __r = ((QString*)self)->toCaseFolded();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QString_trimmed(void* self, long long* __out_len) {
	auto&& __r = ((QString*)self)->trimmed();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QString_trimmed2(void* self, long long* __out_len) {
	auto&& __r = ((QString*)self)->trimmed();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QString_simplified(void* self, long long* __out_len) {
	auto&& __r = ((QString*)self)->simplified();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QString_simplified2(void* self, long long* __out_len) {
	auto&& __r = ((QString*)self)->simplified();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QString_toHtmlEscaped(void* self, long long* __out_len) {
	auto&& __r = ((QString*)self)->toHtmlEscaped();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QString_insert(void* self, long long i, void* c) {
	auto&& __r = ((QString*)self)->insert(i, *(QChar*)c);
	return (void*)&(__r);
}
void* QString_insert2(void* self, long long i, void* uc, long long len_) {
	auto&& __r = ((QString*)self)->insert(i, (QChar*)uc, len_);
	return (void*)&(__r);
}
void* QString_insert3(void* self, long long i, const char* s, long long s_len) {
	QString s_qstr = QString::fromUtf8(s, s_len);
	auto&& __r = ((QString*)self)->insert(i, s_qstr);
	return (void*)&(__r);
}
void* QString_insert4(void* self, long long i, void* v) {
	auto&& __r = ((QString*)self)->insert(i, *(QStringView*)v);
	return (void*)&(__r);
}
void* QString_insert5(void* self, long long i, void* s) {
	auto&& __r = ((QString*)self)->insert(i, *(QLatin1String*)s);
	return (void*)&(__r);
}
void* QString_append(void* self, void* c) {
	auto&& __r = ((QString*)self)->append(*(QChar*)c);
	return (void*)&(__r);
}
void* QString_append2(void* self, void* uc, long long len_) {
	auto&& __r = ((QString*)self)->append((QChar*)uc, len_);
	return (void*)&(__r);
}
void* QString_append3(void* self, const char* s, long long s_len) {
	QString s_qstr = QString::fromUtf8(s, s_len);
	auto&& __r = ((QString*)self)->append(s_qstr);
	return (void*)&(__r);
}
void* QString_append4(void* self, void* v) {
	auto&& __r = ((QString*)self)->append(*(QStringView*)v);
	return (void*)&(__r);
}
void* QString_append5(void* self, void* s) {
	auto&& __r = ((QString*)self)->append(*(QLatin1String*)s);
	return (void*)&(__r);
}
void* QString_prepend(void* self, void* c) {
	auto&& __r = ((QString*)self)->prepend(*(QChar*)c);
	return (void*)&(__r);
}
void* QString_prepend2(void* self, void* uc, long long len_) {
	auto&& __r = ((QString*)self)->prepend((QChar*)uc, len_);
	return (void*)&(__r);
}
void* QString_prepend3(void* self, const char* s, long long s_len) {
	QString s_qstr = QString::fromUtf8(s, s_len);
	auto&& __r = ((QString*)self)->prepend(s_qstr);
	return (void*)&(__r);
}
void* QString_prepend4(void* self, void* v) {
	auto&& __r = ((QString*)self)->prepend(*(QStringView*)v);
	return (void*)&(__r);
}
void* QString_prepend5(void* self, void* s) {
	auto&& __r = ((QString*)self)->prepend(*(QLatin1String*)s);
	return (void*)&(__r);
}
void* QString_assign(void* self, const char* s, long long s_len) {
	QAnyStringView s_qstr = QAnyStringView(s, s_len);
	auto&& __r = ((QString*)self)->assign(s_qstr);
	return (void*)&(__r);
}
void* QString_assign2(void* self, long long n, void* c) {
	auto&& __r = ((QString*)self)->assign(n, *(QChar*)c);
	return (void*)&(__r);
}
void* QString_OperatorPlusAssign(void* self, void* c) {
	auto&& __r = ((QString*)self)->operator+=(*(QChar*)c);
	return (void*)&(__r);
}
void* QString_OperatorPlusAssign2(void* self, const char* s, long long s_len) {
	QString s_qstr = QString::fromUtf8(s, s_len);
	auto&& __r = ((QString*)self)->operator+=(s_qstr);
	return (void*)&(__r);
}
void* QString_OperatorPlusAssign3(void* self, void* v) {
	auto&& __r = ((QString*)self)->operator+=(*(QStringView*)v);
	return (void*)&(__r);
}
void* QString_OperatorPlusAssign4(void* self, void* s) {
	auto&& __r = ((QString*)self)->operator+=(*(QLatin1String*)s);
	return (void*)&(__r);
}
void* QString_remove(void* self, long long i, long long len_) {
	auto&& __r = ((QString*)self)->remove(i, len_);
	return (void*)&(__r);
}
void* QString_remove2(void* self, void* c) {
	auto&& __r = ((QString*)self)->remove(*(QChar*)c);
	return (void*)&(__r);
}
void* QString_remove3(void* self, void* s) {
	auto&& __r = ((QString*)self)->remove(*(QLatin1String*)s);
	return (void*)&(__r);
}
void* QString_remove4(void* self, const char* s, long long s_len) {
	QString s_qstr = QString::fromUtf8(s, s_len);
	auto&& __r = ((QString*)self)->remove(s_qstr);
	return (void*)&(__r);
}
void* QString_removeAt(void* self, long long pos) {
	auto&& __r = ((QString*)self)->removeAt(pos);
	return (void*)&(__r);
}
void* QString_removeFirst(void* self) {
	auto&& __r = ((QString*)self)->removeFirst();
	return (void*)&(__r);
}
void* QString_removeLast(void* self) {
	auto&& __r = ((QString*)self)->removeLast();
	return (void*)&(__r);
}
void* QString_replace(void* self, long long i, long long len_, void* after) {
	auto&& __r = ((QString*)self)->replace(i, len_, *(QChar*)after);
	return (void*)&(__r);
}
void* QString_replace2(void* self, long long i, long long len_, void* s, long long slen) {
	auto&& __r = ((QString*)self)->replace(i, len_, (QChar*)s, slen);
	return (void*)&(__r);
}
void* QString_replace3(void* self, long long i, long long len_, const char* after, long long after_len) {
	QString after_qstr = QString::fromUtf8(after, after_len);
	auto&& __r = ((QString*)self)->replace(i, len_, after_qstr);
	return (void*)&(__r);
}
void* QString_replace4(void* self, void* before, void* after) {
	auto&& __r = ((QString*)self)->replace(*(QChar*)before, *(QChar*)after);
	return (void*)&(__r);
}
void* QString_replace5(void* self, void* before, long long blen, void* after, long long alen) {
	auto&& __r = ((QString*)self)->replace((QChar*)before, blen, (QChar*)after, alen);
	return (void*)&(__r);
}
void* QString_replace6(void* self, void* before, void* after) {
	auto&& __r = ((QString*)self)->replace(*(QLatin1String*)before, *(QLatin1String*)after);
	return (void*)&(__r);
}
void* QString_replace7(void* self, void* before, const char* after, long long after_len) {
	QString after_qstr = QString::fromUtf8(after, after_len);
	auto&& __r = ((QString*)self)->replace(*(QLatin1String*)before, after_qstr);
	return (void*)&(__r);
}
void* QString_replace8(void* self, const char* before, long long before_len, void* after) {
	QString before_qstr = QString::fromUtf8(before, before_len);
	auto&& __r = ((QString*)self)->replace(before_qstr, *(QLatin1String*)after);
	return (void*)&(__r);
}
void* QString_replace9(void* self, const char* before, long long before_len, const char* after, long long after_len) {
	QString before_qstr = QString::fromUtf8(before, before_len);
	QString after_qstr = QString::fromUtf8(after, after_len);
	auto&& __r = ((QString*)self)->replace(before_qstr, after_qstr);
	return (void*)&(__r);
}
void* QString_replace10(void* self, void* c, const char* after, long long after_len) {
	QString after_qstr = QString::fromUtf8(after, after_len);
	auto&& __r = ((QString*)self)->replace(*(QChar*)c, after_qstr);
	return (void*)&(__r);
}
void* QString_replace11(void* self, void* c, void* after) {
	auto&& __r = ((QString*)self)->replace(*(QChar*)c, *(QLatin1String*)after);
	return (void*)&(__r);
}
void* QString_replace12(void* self, void* re, const char* after, long long after_len) {
	QString after_qstr = QString::fromUtf8(after, after_len);
	auto&& __r = ((QString*)self)->replace(*(QRegularExpression*)re, after_qstr);
	return (void*)&(__r);
}
void* QString_remove5(void* self, void* re) {
	auto&& __r = ((QString*)self)->remove(*(QRegularExpression*)re);
	return (void*)&(__r);
}
char** QString_split(void* self, const char* sep, long long sep_len, long long* __out_len) {
	QString sep_qstr = QString::fromUtf8(sep, sep_len);
	auto&& __r = ((QString*)self)->split(sep_qstr);
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
char** QString_split2(void* self, void* sep, long long* __out_len) {
	auto&& __r = ((QString*)self)->split(*(QChar*)sep);
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
char** QString_split3(void* self, void* sep, long long* __out_len) {
	auto&& __r = ((QString*)self)->split(*(QRegularExpression*)sep);
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
char* QString_normalized(void* self, int mode, long long* __out_len) {
	auto&& __r = ((QString*)self)->normalized((QString::NormalizationForm)(mode));
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QString_repeated(void* self, long long times, long long* __out_len) {
	auto&& __r = ((QString*)self)->repeated(times);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QString_nullTerminated(void* self, long long* __out_len) {
	auto&& __r = ((QString*)self)->nullTerminated();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QString_nullTerminated2(void* self, long long* __out_len) {
	auto&& __r = ((QString*)self)->nullTerminated();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QString_nullTerminate(void* self) {
	auto&& __r = ((QString*)self)->nullTerminate();
	return (void*)&(__r);
}
void* QString_toLatin1(void* self) {
	auto&& __r = ((QString*)self)->toLatin1();
	return new QByteArray(__r);
}
void* QString_toLatin12(void* self) {
	auto&& __r = ((QString*)self)->toLatin1();
	return new QByteArray(__r);
}
void* QString_toUtf8(void* self) {
	auto&& __r = ((QString*)self)->toUtf8();
	return new QByteArray(__r);
}
void* QString_toUtf82(void* self) {
	auto&& __r = ((QString*)self)->toUtf8();
	return new QByteArray(__r);
}
void* QString_toLocal8Bit(void* self) {
	auto&& __r = ((QString*)self)->toLocal8Bit();
	return new QByteArray(__r);
}
void* QString_toLocal8Bit2(void* self) {
	auto&& __r = ((QString*)self)->toLocal8Bit();
	return new QByteArray(__r);
}
unsigned int* QString_toUcs4(void* self, long long* __out_len) {
	auto&& __r = ((QString*)self)->toUcs4();
	*__out_len = __r.size();
	unsigned int* __arr = (unsigned int*)malloc(sizeof(unsigned int) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = __e;
	return __arr;
}
char* QString_fromLatin1(void* ba, long long* __out_len) {
	auto&& __r = QString::fromLatin1(*(QByteArrayView*)ba);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QString_fromLatin12(const char* str, long long size, long long* __out_len) {
	auto&& __r = QString::fromLatin1(str, size);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QString_fromUtf8(void* utf8, long long* __out_len) {
	auto&& __r = QString::fromUtf8(*(QByteArrayView*)utf8);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QString_fromUtf82(const char* utf8, long long size, long long* __out_len) {
	auto&& __r = QString::fromUtf8(utf8, size);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QString_fromLocal8Bit(void* ba, long long* __out_len) {
	auto&& __r = QString::fromLocal8Bit(*(QByteArrayView*)ba);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QString_fromLocal8Bit2(const char* str, long long size, long long* __out_len) {
	auto&& __r = QString::fromLocal8Bit(str, size);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QString_fromRawData(void* arg1, long long size, long long* __out_len) {
	auto&& __r = QString::fromRawData((QChar*)arg1, size);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QString_setRawData(void* self, void* unicode, long long size) {
	auto&& __r = ((QString*)self)->setRawData((QChar*)unicode, size);
	return (void*)&(__r);
}
void* QString_setUnicode(void* self, void* unicode, long long size) {
	auto&& __r = ((QString*)self)->setUnicode((QChar*)unicode, size);
	return (void*)&(__r);
}
int QString_compare(void* self, const char* s, long long s_len) {
	QString s_qstr = QString::fromUtf8(s, s_len);
	auto&& __r = ((QString*)self)->compare(s_qstr);
	return __r;
}
int QString_compare2(void* self, void* other) {
	auto&& __r = ((QString*)self)->compare(*(QLatin1String*)other);
	return __r;
}
int QString_compare3(void* self, void* s) {
	auto&& __r = ((QString*)self)->compare(*(QStringView*)s);
	return __r;
}
int QString_compare4(void* self, void* ch) {
	auto&& __r = ((QString*)self)->compare(*(QChar*)ch);
	return __r;
}
int QString_compare5(const char* s1, long long s1_len, const char* s2, long long s2_len) {
	QString s1_qstr = QString::fromUtf8(s1, s1_len);
	QString s2_qstr = QString::fromUtf8(s2, s2_len);
	auto&& __r = QString::compare(s1_qstr, s2_qstr);
	return __r;
}
int QString_compare6(const char* s1, long long s1_len, void* s2) {
	QString s1_qstr = QString::fromUtf8(s1, s1_len);
	auto&& __r = QString::compare(s1_qstr, *(QLatin1String*)s2);
	return __r;
}
int QString_compare7(void* s1, const char* s2, long long s2_len) {
	QString s2_qstr = QString::fromUtf8(s2, s2_len);
	auto&& __r = QString::compare(*(QLatin1String*)s1, s2_qstr);
	return __r;
}
int QString_compare8(const char* s1, long long s1_len, void* s2) {
	QString s1_qstr = QString::fromUtf8(s1, s1_len);
	auto&& __r = QString::compare(s1_qstr, *(QStringView*)s2);
	return __r;
}
int QString_compare9(void* s1, const char* s2, long long s2_len) {
	QString s2_qstr = QString::fromUtf8(s2, s2_len);
	auto&& __r = QString::compare(*(QStringView*)s1, s2_qstr);
	return __r;
}
int QString_localeAwareCompare(void* self, const char* s, long long s_len) {
	QString s_qstr = QString::fromUtf8(s, s_len);
	auto&& __r = ((QString*)self)->localeAwareCompare(s_qstr);
	return __r;
}
int QString_localeAwareCompare2(void* self, void* s) {
	auto&& __r = ((QString*)self)->localeAwareCompare(*(QStringView*)s);
	return __r;
}
int QString_localeAwareCompare3(const char* s1, long long s1_len, const char* s2, long long s2_len) {
	QString s1_qstr = QString::fromUtf8(s1, s1_len);
	QString s2_qstr = QString::fromUtf8(s2, s2_len);
	auto&& __r = QString::localeAwareCompare(s1_qstr, s2_qstr);
	return __r;
}
int QString_localeAwareCompare4(void* s1, void* s2) {
	auto&& __r = QString::localeAwareCompare(*(QStringView*)s1, *(QStringView*)s2);
	return __r;
}
short QString_toShort(void* self) {
	auto&& __r = ((QString*)self)->toShort();
	return __r;
}
unsigned short QString_toUShort(void* self) {
	auto&& __r = ((QString*)self)->toUShort();
	return __r;
}
int QString_toInt(void* self) {
	auto&& __r = ((QString*)self)->toInt();
	return __r;
}
unsigned int QString_toUInt(void* self) {
	auto&& __r = ((QString*)self)->toUInt();
	return __r;
}
long QString_toLong(void* self) {
	auto&& __r = ((QString*)self)->toLong();
	return __r;
}
unsigned long QString_toULong(void* self) {
	auto&& __r = ((QString*)self)->toULong();
	return __r;
}
long long QString_toLongLong(void* self) {
	auto&& __r = ((QString*)self)->toLongLong();
	return __r;
}
unsigned long long QString_toULongLong(void* self) {
	auto&& __r = ((QString*)self)->toULongLong();
	return __r;
}
float QString_toFloat(void* self) {
	auto&& __r = ((QString*)self)->toFloat();
	return __r;
}
double QString_toDouble(void* self) {
	auto&& __r = ((QString*)self)->toDouble();
	return __r;
}
void* QString_setNum(void* self, short arg1) {
	auto&& __r = ((QString*)self)->setNum(arg1);
	return (void*)&(__r);
}
void* QString_setNum2(void* self, unsigned short arg1) {
	auto&& __r = ((QString*)self)->setNum(arg1);
	return (void*)&(__r);
}
void* QString_setNum3(void* self, int arg1) {
	auto&& __r = ((QString*)self)->setNum(arg1);
	return (void*)&(__r);
}
void* QString_setNum4(void* self, unsigned int arg1) {
	auto&& __r = ((QString*)self)->setNum(arg1);
	return (void*)&(__r);
}
void* QString_setNum5(void* self, long arg1) {
	auto&& __r = ((QString*)self)->setNum(arg1);
	return (void*)&(__r);
}
void* QString_setNum6(void* self, unsigned long arg1) {
	auto&& __r = ((QString*)self)->setNum(arg1);
	return (void*)&(__r);
}
void* QString_setNum7(void* self, long long arg1) {
	auto&& __r = ((QString*)self)->setNum(arg1);
	return (void*)&(__r);
}
void* QString_setNum8(void* self, unsigned long long arg1) {
	auto&& __r = ((QString*)self)->setNum(arg1);
	return (void*)&(__r);
}
void* QString_setNum9(void* self, float arg1) {
	auto&& __r = ((QString*)self)->setNum(arg1);
	return (void*)&(__r);
}
void* QString_setNum10(void* self, double arg1) {
	auto&& __r = ((QString*)self)->setNum(arg1);
	return (void*)&(__r);
}
char* QString_number(int arg1, long long* __out_len) {
	auto&& __r = QString::number(arg1);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QString_number2(unsigned int arg1, long long* __out_len) {
	auto&& __r = QString::number(arg1);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QString_number3(long arg1, long long* __out_len) {
	auto&& __r = QString::number(arg1);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QString_number4(unsigned long arg1, long long* __out_len) {
	auto&& __r = QString::number(arg1);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QString_number5(long long arg1, long long* __out_len) {
	auto&& __r = QString::number(arg1);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QString_number6(unsigned long long arg1, long long* __out_len) {
	auto&& __r = QString::number(arg1);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QString_number7(double arg1, long long* __out_len) {
	auto&& __r = QString::number(arg1);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QString_OperatorAssign4(void* self, const char* ch) {
	auto&& __r = ((QString*)self)->operator=(ch);
	return (void*)&(__r);
}
void* QString_OperatorAssign5(void* self, void* a) {
	auto&& __r = ((QString*)self)->operator=(*(QByteArray*)a);
	return (void*)&(__r);
}
void* QString_prepend6(void* self, const char* s) {
	auto&& __r = ((QString*)self)->prepend(s);
	return (void*)&(__r);
}
void* QString_prepend7(void* self, void* s) {
	auto&& __r = ((QString*)self)->prepend(*(QByteArray*)s);
	return (void*)&(__r);
}
void* QString_append6(void* self, const char* s) {
	auto&& __r = ((QString*)self)->append(s);
	return (void*)&(__r);
}
void* QString_append7(void* self, void* s) {
	auto&& __r = ((QString*)self)->append(*(QByteArray*)s);
	return (void*)&(__r);
}
void* QString_insert6(void* self, long long i, const char* s) {
	auto&& __r = ((QString*)self)->insert(i, s);
	return (void*)&(__r);
}
void* QString_insert7(void* self, long long i, void* s) {
	auto&& __r = ((QString*)self)->insert(i, *(QByteArray*)s);
	return (void*)&(__r);
}
void* QString_OperatorPlusAssign5(void* self, const char* s) {
	auto&& __r = ((QString*)self)->operator+=(s);
	return (void*)&(__r);
}
void* QString_OperatorPlusAssign6(void* self, void* s) {
	auto&& __r = ((QString*)self)->operator+=(*(QByteArray*)s);
	return (void*)&(__r);
}
void* QString_begin(void* self) {
	auto&& __r = ((QString*)self)->begin();
	return (void*)(__r);
}
void* QString_begin2(void* self) {
	auto&& __r = ((QString*)self)->begin();
	return (void*)(__r);
}
void* QString_cbegin(void* self) {
	auto&& __r = ((QString*)self)->cbegin();
	return (void*)(__r);
}
void* QString_constBegin(void* self) {
	auto&& __r = ((QString*)self)->constBegin();
	return (void*)(__r);
}
void* QString_end(void* self) {
	auto&& __r = ((QString*)self)->end();
	return (void*)(__r);
}
void* QString_end2(void* self) {
	auto&& __r = ((QString*)self)->end();
	return (void*)(__r);
}
void* QString_cend(void* self) {
	auto&& __r = ((QString*)self)->cend();
	return (void*)(__r);
}
void* QString_constEnd(void* self) {
	auto&& __r = ((QString*)self)->constEnd();
	return (void*)(__r);
}
void QString_push_back(void* self, void* c) {
	((QString*)self)->push_back(*(QChar*)c);
}
void QString_push_back2(void* self, const char* s, long long s_len) {
	QString s_qstr = QString::fromUtf8(s, s_len);
	((QString*)self)->push_back(s_qstr);
}
void QString_push_front(void* self, void* c) {
	((QString*)self)->push_front(*(QChar*)c);
}
void QString_push_front2(void* self, const char* s, long long s_len) {
	QString s_qstr = QString::fromUtf8(s, s_len);
	((QString*)self)->push_front(s_qstr);
}
void QString_shrink_to_fit(void* self) {
	((QString*)self)->shrink_to_fit();
}
void* QString_erase(void* self, void* first, void* last) {
	auto&& __r = ((QString*)self)->erase((QChar*)first, (QChar*)last);
	return (void*)(__r);
}
void* QString_erase2(void* self, void* it) {
	auto&& __r = ((QString*)self)->erase((QChar*)it);
	return (void*)(__r);
}
long long QString_max_size(void* self) {
	auto&& __r = ((QString*)self)->max_size();
	return __r;
}
int QString_isNull(void* self) {
	auto&& __r = ((QString*)self)->isNull();
	return __r;
}
int QString_isRightToLeft(void* self) {
	auto&& __r = ((QString*)self)->isRightToLeft();
	return __r;
}
int QString_isValidUtf16(void* self) {
	auto&& __r = ((QString*)self)->isValidUtf16();
	return __r;
}
void* QString_fill2(void* self, void* c, long long size) {
	auto&& __r = ((QString*)self)->fill(*(QChar*)c, size);
	return (void*)&(__r);
}
long long QString_indexOf6(void* self, void* c, long long from) {
	auto&& __r = ((QString*)self)->indexOf(*(QChar*)c, from);
	return __r;
}
long long QString_indexOf7(void* self, void* c, long long from, int cs) {
	auto&& __r = ((QString*)self)->indexOf(*(QChar*)c, from, (Qt::CaseSensitivity)(cs));
	return __r;
}
long long QString_indexOf8(void* self, void* s, long long from) {
	auto&& __r = ((QString*)self)->indexOf(*(QLatin1String*)s, from);
	return __r;
}
long long QString_indexOf9(void* self, void* s, long long from, int cs) {
	auto&& __r = ((QString*)self)->indexOf(*(QLatin1String*)s, from, (Qt::CaseSensitivity)(cs));
	return __r;
}
long long QString_indexOf10(void* self, const char* s, long long s_len, long long from) {
	QString s_qstr = QString::fromUtf8(s, s_len);
	auto&& __r = ((QString*)self)->indexOf(s_qstr, from);
	return __r;
}
long long QString_indexOf11(void* self, const char* s, long long s_len, long long from, int cs) {
	QString s_qstr = QString::fromUtf8(s, s_len);
	auto&& __r = ((QString*)self)->indexOf(s_qstr, from, (Qt::CaseSensitivity)(cs));
	return __r;
}
long long QString_indexOf12(void* self, void* s, long long from) {
	auto&& __r = ((QString*)self)->indexOf(*(QStringView*)s, from);
	return __r;
}
long long QString_indexOf13(void* self, void* s, long long from, int cs) {
	auto&& __r = ((QString*)self)->indexOf(*(QStringView*)s, from, (Qt::CaseSensitivity)(cs));
	return __r;
}
long long QString_lastIndexOf10(void* self, void* c, int cs) {
	auto&& __r = ((QString*)self)->lastIndexOf(*(QChar*)c, (Qt::CaseSensitivity)(cs));
	return __r;
}
long long QString_lastIndexOf11(void* self, void* c, long long from, int cs) {
	auto&& __r = ((QString*)self)->lastIndexOf(*(QChar*)c, from, (Qt::CaseSensitivity)(cs));
	return __r;
}
long long QString_lastIndexOf12(void* self, void* s, int cs) {
	auto&& __r = ((QString*)self)->lastIndexOf(*(QLatin1String*)s, (Qt::CaseSensitivity)(cs));
	return __r;
}
long long QString_lastIndexOf13(void* self, void* s, long long from, int cs) {
	auto&& __r = ((QString*)self)->lastIndexOf(*(QLatin1String*)s, from, (Qt::CaseSensitivity)(cs));
	return __r;
}
long long QString_lastIndexOf14(void* self, const char* s, long long s_len, int cs) {
	QString s_qstr = QString::fromUtf8(s, s_len);
	auto&& __r = ((QString*)self)->lastIndexOf(s_qstr, (Qt::CaseSensitivity)(cs));
	return __r;
}
long long QString_lastIndexOf15(void* self, const char* s, long long s_len, long long from, int cs) {
	QString s_qstr = QString::fromUtf8(s, s_len);
	auto&& __r = ((QString*)self)->lastIndexOf(s_qstr, from, (Qt::CaseSensitivity)(cs));
	return __r;
}
long long QString_lastIndexOf16(void* self, void* s, int cs) {
	auto&& __r = ((QString*)self)->lastIndexOf(*(QStringView*)s, (Qt::CaseSensitivity)(cs));
	return __r;
}
long long QString_lastIndexOf17(void* self, void* s, long long from, int cs) {
	auto&& __r = ((QString*)self)->lastIndexOf(*(QStringView*)s, from, (Qt::CaseSensitivity)(cs));
	return __r;
}
int QString_contains6(void* self, void* c, int cs) {
	auto&& __r = ((QString*)self)->contains(*(QChar*)c, (Qt::CaseSensitivity)(cs));
	return __r;
}
int QString_contains7(void* self, const char* s, long long s_len, int cs) {
	QString s_qstr = QString::fromUtf8(s, s_len);
	auto&& __r = ((QString*)self)->contains(s_qstr, (Qt::CaseSensitivity)(cs));
	return __r;
}
int QString_contains8(void* self, void* s, int cs) {
	auto&& __r = ((QString*)self)->contains(*(QLatin1String*)s, (Qt::CaseSensitivity)(cs));
	return __r;
}
int QString_contains9(void* self, void* s, int cs) {
	auto&& __r = ((QString*)self)->contains(*(QStringView*)s, (Qt::CaseSensitivity)(cs));
	return __r;
}
long long QString_count6(void* self, void* c, int cs) {
	auto&& __r = ((QString*)self)->count(*(QChar*)c, (Qt::CaseSensitivity)(cs));
	return __r;
}
long long QString_count7(void* self, const char* s, long long s_len, int cs) {
	QString s_qstr = QString::fromUtf8(s, s_len);
	auto&& __r = ((QString*)self)->count(s_qstr, (Qt::CaseSensitivity)(cs));
	return __r;
}
long long QString_count8(void* self, void* s, int cs) {
	auto&& __r = ((QString*)self)->count(*(QStringView*)s, (Qt::CaseSensitivity)(cs));
	return __r;
}
long long QString_indexOf14(void* self, void* re, long long from) {
	auto&& __r = ((QString*)self)->indexOf(*(QRegularExpression*)re, from);
	return __r;
}
long long QString_indexOf15(void* self, void* re, long long from, void* rmatch) {
	auto&& __r = ((QString*)self)->indexOf(*(QRegularExpression*)re, from, (QRegularExpressionMatch*)rmatch);
	return __r;
}
long long QString_lastIndexOf18(void* self, void* re, long long from, void* rmatch) {
	auto&& __r = ((QString*)self)->lastIndexOf(*(QRegularExpression*)re, from, (QRegularExpressionMatch*)rmatch);
	return __r;
}
int QString_contains10(void* self, void* re, void* rmatch) {
	auto&& __r = ((QString*)self)->contains(*(QRegularExpression*)re, (QRegularExpressionMatch*)rmatch);
	return __r;
}
char* QString_section4(void* self, void* sep, long long start, long long end, long long* __out_len) {
	auto&& __r = ((QString*)self)->section(*(QChar*)sep, start, end);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QString_section5(void* self, void* sep, long long start, long long end, int flags, long long* __out_len) {
	auto&& __r = ((QString*)self)->section(*(QChar*)sep, start, end, QFlags<QString::SectionFlag>(flags));
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QString_section6(void* self, const char* in_sep, long long in_sep_len, long long start, long long end, long long* __out_len) {
	QString in_sep_qstr = QString::fromUtf8(in_sep, in_sep_len);
	auto&& __r = ((QString*)self)->section(in_sep_qstr, start, end);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QString_section7(void* self, const char* in_sep, long long in_sep_len, long long start, long long end, int flags, long long* __out_len) {
	QString in_sep_qstr = QString::fromUtf8(in_sep, in_sep_len);
	auto&& __r = ((QString*)self)->section(in_sep_qstr, start, end, QFlags<QString::SectionFlag>(flags));
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QString_section8(void* self, void* re, long long start, long long end, long long* __out_len) {
	auto&& __r = ((QString*)self)->section(*(QRegularExpression*)re, start, end);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QString_section9(void* self, void* re, long long start, long long end, int flags, long long* __out_len) {
	auto&& __r = ((QString*)self)->section(*(QRegularExpression*)re, start, end, QFlags<QString::SectionFlag>(flags));
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QString_mid3(void* self, long long position, long long n, long long* __out_len) {
	auto&& __r = ((QString*)self)->mid(position, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QString_mid4(void* self, long long position, long long n, long long* __out_len) {
	auto&& __r = ((QString*)self)->mid(position, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QString_startsWith5(void* self, const char* s, long long s_len, int cs) {
	QString s_qstr = QString::fromUtf8(s, s_len);
	auto&& __r = ((QString*)self)->startsWith(s_qstr, (Qt::CaseSensitivity)(cs));
	return __r;
}
int QString_startsWith6(void* self, void* s, int cs) {
	auto&& __r = ((QString*)self)->startsWith(*(QStringView*)s, (Qt::CaseSensitivity)(cs));
	return __r;
}
int QString_startsWith7(void* self, void* s, int cs) {
	auto&& __r = ((QString*)self)->startsWith(*(QLatin1String*)s, (Qt::CaseSensitivity)(cs));
	return __r;
}
int QString_startsWith8(void* self, void* c, int cs) {
	auto&& __r = ((QString*)self)->startsWith(*(QChar*)c, (Qt::CaseSensitivity)(cs));
	return __r;
}
int QString_endsWith5(void* self, const char* s, long long s_len, int cs) {
	QString s_qstr = QString::fromUtf8(s, s_len);
	auto&& __r = ((QString*)self)->endsWith(s_qstr, (Qt::CaseSensitivity)(cs));
	return __r;
}
int QString_endsWith6(void* self, void* s, int cs) {
	auto&& __r = ((QString*)self)->endsWith(*(QStringView*)s, (Qt::CaseSensitivity)(cs));
	return __r;
}
int QString_endsWith7(void* self, void* s, int cs) {
	auto&& __r = ((QString*)self)->endsWith(*(QLatin1String*)s, (Qt::CaseSensitivity)(cs));
	return __r;
}
int QString_endsWith8(void* self, void* c, int cs) {
	auto&& __r = ((QString*)self)->endsWith(*(QChar*)c, (Qt::CaseSensitivity)(cs));
	return __r;
}
char* QString_leftJustified2(void* self, long long width, void* fill, long long* __out_len) {
	auto&& __r = ((QString*)self)->leftJustified(width, *(QChar*)fill);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QString_leftJustified3(void* self, long long width, void* fill, int trunc, long long* __out_len) {
	auto&& __r = ((QString*)self)->leftJustified(width, *(QChar*)fill, trunc);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QString_rightJustified2(void* self, long long width, void* fill, long long* __out_len) {
	auto&& __r = ((QString*)self)->rightJustified(width, *(QChar*)fill);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QString_rightJustified3(void* self, long long width, void* fill, int trunc, long long* __out_len) {
	auto&& __r = ((QString*)self)->rightJustified(width, *(QChar*)fill, trunc);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QString_remove6(void* self, void* c, int cs) {
	auto&& __r = ((QString*)self)->remove(*(QChar*)c, (Qt::CaseSensitivity)(cs));
	return (void*)&(__r);
}
void* QString_remove7(void* self, void* s, int cs) {
	auto&& __r = ((QString*)self)->remove(*(QLatin1String*)s, (Qt::CaseSensitivity)(cs));
	return (void*)&(__r);
}
void* QString_remove8(void* self, const char* s, long long s_len, int cs) {
	QString s_qstr = QString::fromUtf8(s, s_len);
	auto&& __r = ((QString*)self)->remove(s_qstr, (Qt::CaseSensitivity)(cs));
	return (void*)&(__r);
}
void* QString_replace13(void* self, void* before, void* after, int cs) {
	auto&& __r = ((QString*)self)->replace(*(QChar*)before, *(QChar*)after, (Qt::CaseSensitivity)(cs));
	return (void*)&(__r);
}
void* QString_replace14(void* self, void* before, long long blen, void* after, long long alen, int cs) {
	auto&& __r = ((QString*)self)->replace((QChar*)before, blen, (QChar*)after, alen, (Qt::CaseSensitivity)(cs));
	return (void*)&(__r);
}
void* QString_replace15(void* self, void* before, void* after, int cs) {
	auto&& __r = ((QString*)self)->replace(*(QLatin1String*)before, *(QLatin1String*)after, (Qt::CaseSensitivity)(cs));
	return (void*)&(__r);
}
void* QString_replace16(void* self, void* before, const char* after, long long after_len, int cs) {
	QString after_qstr = QString::fromUtf8(after, after_len);
	auto&& __r = ((QString*)self)->replace(*(QLatin1String*)before, after_qstr, (Qt::CaseSensitivity)(cs));
	return (void*)&(__r);
}
void* QString_replace17(void* self, const char* before, long long before_len, void* after, int cs) {
	QString before_qstr = QString::fromUtf8(before, before_len);
	auto&& __r = ((QString*)self)->replace(before_qstr, *(QLatin1String*)after, (Qt::CaseSensitivity)(cs));
	return (void*)&(__r);
}
void* QString_replace18(void* self, const char* before, long long before_len, const char* after, long long after_len, int cs) {
	QString before_qstr = QString::fromUtf8(before, before_len);
	QString after_qstr = QString::fromUtf8(after, after_len);
	auto&& __r = ((QString*)self)->replace(before_qstr, after_qstr, (Qt::CaseSensitivity)(cs));
	return (void*)&(__r);
}
void* QString_replace19(void* self, void* c, const char* after, long long after_len, int cs) {
	QString after_qstr = QString::fromUtf8(after, after_len);
	auto&& __r = ((QString*)self)->replace(*(QChar*)c, after_qstr, (Qt::CaseSensitivity)(cs));
	return (void*)&(__r);
}
void* QString_replace20(void* self, void* c, void* after, int cs) {
	auto&& __r = ((QString*)self)->replace(*(QChar*)c, *(QLatin1String*)after, (Qt::CaseSensitivity)(cs));
	return (void*)&(__r);
}
char** QString_split4(void* self, const char* sep, long long sep_len, int behavior, long long* __out_len) {
	QString sep_qstr = QString::fromUtf8(sep, sep_len);
	auto&& __r = ((QString*)self)->split(sep_qstr, QFlags<Qt::SplitBehaviorFlags>(behavior));
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
char** QString_split5(void* self, const char* sep, long long sep_len, int behavior, int cs, long long* __out_len) {
	QString sep_qstr = QString::fromUtf8(sep, sep_len);
	auto&& __r = ((QString*)self)->split(sep_qstr, QFlags<Qt::SplitBehaviorFlags>(behavior), (Qt::CaseSensitivity)(cs));
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
char** QString_split6(void* self, void* sep, int behavior, long long* __out_len) {
	auto&& __r = ((QString*)self)->split(*(QChar*)sep, QFlags<Qt::SplitBehaviorFlags>(behavior));
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
char** QString_split7(void* self, void* sep, int behavior, int cs, long long* __out_len) {
	auto&& __r = ((QString*)self)->split(*(QChar*)sep, QFlags<Qt::SplitBehaviorFlags>(behavior), (Qt::CaseSensitivity)(cs));
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
char** QString_split8(void* self, void* sep, int behavior, long long* __out_len) {
	auto&& __r = ((QString*)self)->split(*(QRegularExpression*)sep, QFlags<Qt::SplitBehaviorFlags>(behavior));
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
char* QString_normalized2(void* self, int mode, int version, long long* __out_len) {
	auto&& __r = ((QString*)self)->normalized((QString::NormalizationForm)(mode), (QChar::UnicodeVersion)(version));
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QString_compare10(void* self, const char* s, long long s_len, int cs) {
	QString s_qstr = QString::fromUtf8(s, s_len);
	auto&& __r = ((QString*)self)->compare(s_qstr, (Qt::CaseSensitivity)(cs));
	return __r;
}
int QString_compare11(void* self, void* other, int cs) {
	auto&& __r = ((QString*)self)->compare(*(QLatin1String*)other, (Qt::CaseSensitivity)(cs));
	return __r;
}
int QString_compare12(void* self, void* s, int cs) {
	auto&& __r = ((QString*)self)->compare(*(QStringView*)s, (Qt::CaseSensitivity)(cs));
	return __r;
}
int QString_compare13(void* self, void* ch, int cs) {
	auto&& __r = ((QString*)self)->compare(*(QChar*)ch, (Qt::CaseSensitivity)(cs));
	return __r;
}
int QString_compare14(const char* s1, long long s1_len, const char* s2, long long s2_len, int cs) {
	QString s1_qstr = QString::fromUtf8(s1, s1_len);
	QString s2_qstr = QString::fromUtf8(s2, s2_len);
	auto&& __r = QString::compare(s1_qstr, s2_qstr, (Qt::CaseSensitivity)(cs));
	return __r;
}
int QString_compare15(const char* s1, long long s1_len, void* s2, int cs) {
	QString s1_qstr = QString::fromUtf8(s1, s1_len);
	auto&& __r = QString::compare(s1_qstr, *(QLatin1String*)s2, (Qt::CaseSensitivity)(cs));
	return __r;
}
int QString_compare16(void* s1, const char* s2, long long s2_len, int cs) {
	QString s2_qstr = QString::fromUtf8(s2, s2_len);
	auto&& __r = QString::compare(*(QLatin1String*)s1, s2_qstr, (Qt::CaseSensitivity)(cs));
	return __r;
}
int QString_compare17(const char* s1, long long s1_len, void* s2, int cs) {
	QString s1_qstr = QString::fromUtf8(s1, s1_len);
	auto&& __r = QString::compare(s1_qstr, *(QStringView*)s2, (Qt::CaseSensitivity)(cs));
	return __r;
}
int QString_compare18(void* s1, const char* s2, long long s2_len, int cs) {
	QString s2_qstr = QString::fromUtf8(s2, s2_len);
	auto&& __r = QString::compare(*(QStringView*)s1, s2_qstr, (Qt::CaseSensitivity)(cs));
	return __r;
}
short QString_toShort2(void* self, int* ok) {
	bool ok_local{};
	auto&& __r = ((QString*)self)->toShort(&ok_local);
	*ok = ok_local;
	return __r;
}
short QString_toShort3(void* self, int* ok, int base) {
	bool ok_local{};
	auto&& __r = ((QString*)self)->toShort(&ok_local, base);
	*ok = ok_local;
	return __r;
}
unsigned short QString_toUShort2(void* self, int* ok) {
	bool ok_local{};
	auto&& __r = ((QString*)self)->toUShort(&ok_local);
	*ok = ok_local;
	return __r;
}
unsigned short QString_toUShort3(void* self, int* ok, int base) {
	bool ok_local{};
	auto&& __r = ((QString*)self)->toUShort(&ok_local, base);
	*ok = ok_local;
	return __r;
}
int QString_toInt2(void* self, int* ok) {
	bool ok_local{};
	auto&& __r = ((QString*)self)->toInt(&ok_local);
	*ok = ok_local;
	return __r;
}
int QString_toInt3(void* self, int* ok, int base) {
	bool ok_local{};
	auto&& __r = ((QString*)self)->toInt(&ok_local, base);
	*ok = ok_local;
	return __r;
}
unsigned int QString_toUInt2(void* self, int* ok) {
	bool ok_local{};
	auto&& __r = ((QString*)self)->toUInt(&ok_local);
	*ok = ok_local;
	return __r;
}
unsigned int QString_toUInt3(void* self, int* ok, int base) {
	bool ok_local{};
	auto&& __r = ((QString*)self)->toUInt(&ok_local, base);
	*ok = ok_local;
	return __r;
}
long QString_toLong2(void* self, int* ok) {
	bool ok_local{};
	auto&& __r = ((QString*)self)->toLong(&ok_local);
	*ok = ok_local;
	return __r;
}
long QString_toLong3(void* self, int* ok, int base) {
	bool ok_local{};
	auto&& __r = ((QString*)self)->toLong(&ok_local, base);
	*ok = ok_local;
	return __r;
}
unsigned long QString_toULong2(void* self, int* ok) {
	bool ok_local{};
	auto&& __r = ((QString*)self)->toULong(&ok_local);
	*ok = ok_local;
	return __r;
}
unsigned long QString_toULong3(void* self, int* ok, int base) {
	bool ok_local{};
	auto&& __r = ((QString*)self)->toULong(&ok_local, base);
	*ok = ok_local;
	return __r;
}
long long QString_toLongLong2(void* self, int* ok) {
	bool ok_local{};
	auto&& __r = ((QString*)self)->toLongLong(&ok_local);
	*ok = ok_local;
	return __r;
}
long long QString_toLongLong3(void* self, int* ok, int base) {
	bool ok_local{};
	auto&& __r = ((QString*)self)->toLongLong(&ok_local, base);
	*ok = ok_local;
	return __r;
}
unsigned long long QString_toULongLong2(void* self, int* ok) {
	bool ok_local{};
	auto&& __r = ((QString*)self)->toULongLong(&ok_local);
	*ok = ok_local;
	return __r;
}
unsigned long long QString_toULongLong3(void* self, int* ok, int base) {
	bool ok_local{};
	auto&& __r = ((QString*)self)->toULongLong(&ok_local, base);
	*ok = ok_local;
	return __r;
}
float QString_toFloat2(void* self, int* ok) {
	bool ok_local{};
	auto&& __r = ((QString*)self)->toFloat(&ok_local);
	*ok = ok_local;
	return __r;
}
double QString_toDouble2(void* self, int* ok) {
	bool ok_local{};
	auto&& __r = ((QString*)self)->toDouble(&ok_local);
	*ok = ok_local;
	return __r;
}
void* QString_setNum11(void* self, short arg1, int base) {
	auto&& __r = ((QString*)self)->setNum(arg1, base);
	return (void*)&(__r);
}
void* QString_setNum12(void* self, unsigned short arg1, int base) {
	auto&& __r = ((QString*)self)->setNum(arg1, base);
	return (void*)&(__r);
}
void* QString_setNum13(void* self, int arg1, int base) {
	auto&& __r = ((QString*)self)->setNum(arg1, base);
	return (void*)&(__r);
}
void* QString_setNum14(void* self, unsigned int arg1, int base) {
	auto&& __r = ((QString*)self)->setNum(arg1, base);
	return (void*)&(__r);
}
void* QString_setNum15(void* self, long arg1, int base) {
	auto&& __r = ((QString*)self)->setNum(arg1, base);
	return (void*)&(__r);
}
void* QString_setNum16(void* self, unsigned long arg1, int base) {
	auto&& __r = ((QString*)self)->setNum(arg1, base);
	return (void*)&(__r);
}
void* QString_setNum17(void* self, long long arg1, int base) {
	auto&& __r = ((QString*)self)->setNum(arg1, base);
	return (void*)&(__r);
}
void* QString_setNum18(void* self, unsigned long long arg1, int base) {
	auto&& __r = ((QString*)self)->setNum(arg1, base);
	return (void*)&(__r);
}
void* QString_setNum19(void* self, float arg1, char format) {
	auto&& __r = ((QString*)self)->setNum(arg1, format);
	return (void*)&(__r);
}
void* QString_setNum20(void* self, float arg1, char format, int precision) {
	auto&& __r = ((QString*)self)->setNum(arg1, format, precision);
	return (void*)&(__r);
}
void* QString_setNum21(void* self, double arg1, char format) {
	auto&& __r = ((QString*)self)->setNum(arg1, format);
	return (void*)&(__r);
}
void* QString_setNum22(void* self, double arg1, char format, int precision) {
	auto&& __r = ((QString*)self)->setNum(arg1, format, precision);
	return (void*)&(__r);
}
char* QString_number8(int arg1, int base, long long* __out_len) {
	auto&& __r = QString::number(arg1, base);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QString_number9(unsigned int arg1, int base, long long* __out_len) {
	auto&& __r = QString::number(arg1, base);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QString_number10(long arg1, int base, long long* __out_len) {
	auto&& __r = QString::number(arg1, base);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QString_number11(unsigned long arg1, int base, long long* __out_len) {
	auto&& __r = QString::number(arg1, base);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QString_number12(long long arg1, int base, long long* __out_len) {
	auto&& __r = QString::number(arg1, base);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QString_number13(unsigned long long arg1, int base, long long* __out_len) {
	auto&& __r = QString::number(arg1, base);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QString_number14(double arg1, char format, long long* __out_len) {
	auto&& __r = QString::number(arg1, format);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QString_number15(double arg1, char format, int precision, long long* __out_len) {
	auto&& __r = QString::number(arg1, format, precision);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QString_OperatorEqual(void* self, const char* rhs, long long rhs_len) {
	QString rhs_qstr = QString::fromUtf8(rhs, rhs_len);
	auto&& __r = (*(QString*)self) == (rhs_qstr);
	return __r;
}
int QString_OperatorNotEqual(void* self, const char* rhs, long long rhs_len) {
	QString rhs_qstr = QString::fromUtf8(rhs, rhs_len);
	auto&& __r = (*(QString*)self) != (rhs_qstr);
	return __r;
}
int QString_OperatorLesser(void* self, const char* rhs, long long rhs_len) {
	QString rhs_qstr = QString::fromUtf8(rhs, rhs_len);
	auto&& __r = (*(QString*)self) < (rhs_qstr);
	return __r;
}
int QString_OperatorGreater(void* self, const char* rhs, long long rhs_len) {
	QString rhs_qstr = QString::fromUtf8(rhs, rhs_len);
	auto&& __r = (*(QString*)self) > (rhs_qstr);
	return __r;
}
int QString_OperatorLesserOrEqual(void* self, const char* rhs, long long rhs_len) {
	QString rhs_qstr = QString::fromUtf8(rhs, rhs_len);
	auto&& __r = (*(QString*)self) <= (rhs_qstr);
	return __r;
}
int QString_OperatorGreaterOrEqual(void* self, const char* rhs, long long rhs_len) {
	QString rhs_qstr = QString::fromUtf8(rhs, rhs_len);
	auto&& __r = (*(QString*)self) >= (rhs_qstr);
	return __r;
}
int QString_OperatorEqual2(void* self, void* rhs) {
	auto&& __r = (*(QString*)self) == (*(QChar*)rhs);
	return __r;
}
int QString_OperatorNotEqual2(void* self, void* rhs) {
	auto&& __r = (*(QString*)self) != (*(QChar*)rhs);
	return __r;
}
int QString_OperatorLesser2(void* self, void* rhs) {
	auto&& __r = (*(QString*)self) < (*(QChar*)rhs);
	return __r;
}
int QString_OperatorGreater2(void* self, void* rhs) {
	auto&& __r = (*(QString*)self) > (*(QChar*)rhs);
	return __r;
}
int QString_OperatorLesserOrEqual2(void* self, void* rhs) {
	auto&& __r = (*(QString*)self) <= (*(QChar*)rhs);
	return __r;
}
int QString_OperatorGreaterOrEqual2(void* self, void* rhs) {
	auto&& __r = (*(QString*)self) >= (*(QChar*)rhs);
	return __r;
}
int QString_OperatorEqual3(void* self, void* rhs) {
	auto&& __r = (*(QString*)self) == (*(QByteArrayView*)rhs);
	return __r;
}
int QString_OperatorNotEqual3(void* self, void* rhs) {
	auto&& __r = (*(QString*)self) != (*(QByteArrayView*)rhs);
	return __r;
}
int QString_OperatorLesser3(void* self, void* rhs) {
	auto&& __r = (*(QString*)self) < (*(QByteArrayView*)rhs);
	return __r;
}
int QString_OperatorGreater3(void* self, void* rhs) {
	auto&& __r = (*(QString*)self) > (*(QByteArrayView*)rhs);
	return __r;
}
int QString_OperatorLesserOrEqual3(void* self, void* rhs) {
	auto&& __r = (*(QString*)self) <= (*(QByteArrayView*)rhs);
	return __r;
}
int QString_OperatorGreaterOrEqual3(void* self, void* rhs) {
	auto&& __r = (*(QString*)self) >= (*(QByteArrayView*)rhs);
	return __r;
}
int QString_OperatorEqual4(void* self, void* rhs) {
	auto&& __r = (*(QString*)self) == (*(QByteArray*)rhs);
	return __r;
}
int QString_OperatorNotEqual4(void* self, void* rhs) {
	auto&& __r = (*(QString*)self) != (*(QByteArray*)rhs);
	return __r;
}
int QString_OperatorLesser4(void* self, void* rhs) {
	auto&& __r = (*(QString*)self) < (*(QByteArray*)rhs);
	return __r;
}
int QString_OperatorGreater4(void* self, void* rhs) {
	auto&& __r = (*(QString*)self) > (*(QByteArray*)rhs);
	return __r;
}
int QString_OperatorLesserOrEqual4(void* self, void* rhs) {
	auto&& __r = (*(QString*)self) <= (*(QByteArray*)rhs);
	return __r;
}
int QString_OperatorGreaterOrEqual4(void* self, void* rhs) {
	auto&& __r = (*(QString*)self) >= (*(QByteArray*)rhs);
	return __r;
}
char* QString_OperatorPlus(void* self, const char* s2, long long s2_len, long long* __out_len) {
	QString s2_qstr = QString::fromUtf8(s2, s2_len);
	auto&& __r = (*(QString*)self) + (s2_qstr);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QString_OperatorPlus2(void* self, void* s2, long long* __out_len) {
	auto&& __r = (*(QString*)self) + (*(QChar*)s2);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QString_OperatorPlus3(void* self, void* rhs, long long* __out_len) {
	auto&& __r = (*(QString*)self) + (*(QStringView*)rhs);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QString_OperatorPlus4(void* self, const char* s2, long long* __out_len) {
	auto&& __r = (*(QString*)self) + (s2);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QString_OperatorPlus5(void* self, void* ba, long long* __out_len) {
	auto&& __r = (*(QString*)self) + (*(QByteArray*)ba);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QString_delete(void* self) { delete (QString*)self; }

}

#include <QtCore/qeventloop.h>
#include <QtCore/qthread.h>
#include "gen_qeventlooplocker.h"

extern "C" {

void* QEventLoopLocker_new() {
	return new QEventLoopLocker();
}
void* QEventLoopLocker_new2(void* loop) {
	return new QEventLoopLocker((QEventLoop*)loop);
}
void* QEventLoopLocker_new3(void* thread) {
	return new QEventLoopLocker((QThread*)thread);
}
void QEventLoopLocker_swap(void* self, void* other) {
	((QEventLoopLocker*)self)->swap(*(QEventLoopLocker*)other);
}
void QEventLoopLocker_delete(void* self) { delete (QEventLoopLocker*)self; }

}

#include <QtCore/qabstractnativeeventfilter.h>
#include <QtCore/qbytearray.h>
#include "gen_qabstractnativeeventfilter.h"
#include "gen_purego_cb.h"

class QAbstractNativeEventFilter__Go : public QAbstractNativeEventFilter {
public:
	long long __goId;
	unsigned char __goOn[1] = {};
	template<typename... A> QAbstractNativeEventFilter__Go(long long id, A&&... a) : QAbstractNativeEventFilter(std::forward<A>(a)...), __goId(id) {}
	bool nativeEventFilter(const QByteArray& eventType, void* message, qintptr* result) override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_int_voidp_voidp_longlongp) { int __r = ((int(*)(long long, long long, void*, void*, long long*, int*))cppgo_virt_int_voidp_voidp_longlongp)(__goId, -3445417634659622720LL, (void*)&eventType, (void*)message, (long long*)result, &__h); if (__h) return __r; }
		return bool{};
	}
};

extern "C" {

void QAbstractNativeEventFilter__set_override(void* self, int idx) { ((QAbstractNativeEventFilter__Go*)self)->__goOn[idx] = 1; }
void* QAbstractNativeEventFilter__Go_new(long long id) { return new QAbstractNativeEventFilter__Go(id); }
void QAbstractNativeEventFilter_delete(void* self) { delete (QAbstractNativeEventFilter*)self; }

}

#include <QtCore/qcborcommon.h>
#include <QtCore/qstring.h>
#include "gen_qcborerror.h"

extern "C" {

int QCborError_ToQCborErrorCode(void* self) {
	auto&& __r = static_cast<QCborError::Code>(*(QCborError*)self);
	return (int)(__r);
}
char* QCborError_toString(void* self, long long* __out_len) {
	auto&& __r = ((QCborError*)self)->toString();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QCborError_c_get(void* self) {
	auto __r = ((QCborError*)self)->c;
	return (int)(__r);
}
void QCborError_c_set(void* self, int c) {
	((QCborError*)self)->c = (QCborError::Code)(c);
}
void QCborError_delete(void* self) { delete (QCborError*)self; }

}

#include <QtCore/qproperty.h>
#include <QtCore/qmetatype.h>
#include "gen_quntypedpropertybinding.h"

extern "C" {

void* QUntypedPropertyBinding_new() {
	return new QUntypedPropertyBinding();
}
void* QUntypedPropertyBinding_new2(void* other) {
	return new QUntypedPropertyBinding(*(QUntypedPropertyBinding*)other);
}
void* QUntypedPropertyBinding_OperatorAssign(void* self, void* other) {
	auto&& __r = ((QUntypedPropertyBinding*)self)->operator=(*(QUntypedPropertyBinding*)other);
	return (void*)&(__r);
}
int QUntypedPropertyBinding_isNull(void* self) {
	auto&& __r = ((QUntypedPropertyBinding*)self)->isNull();
	return __r;
}
void* QUntypedPropertyBinding_error(void* self) {
	auto&& __r = ((QUntypedPropertyBinding*)self)->error();
	return new QPropertyBindingError(__r);
}
void* QUntypedPropertyBinding_valueMetaType(void* self) {
	auto&& __r = ((QUntypedPropertyBinding*)self)->valueMetaType();
	return new QMetaType(__r);
}
void QUntypedPropertyBinding_delete(void* self) { delete (QUntypedPropertyBinding*)self; }

}

#include <QtCore/qeasingcurve.h>
#include <QtCore/qpoint.h>
#include "gen_qeasingcurve.h"

extern "C" {

void* QEasingCurve_new() {
	return new QEasingCurve();
}
void* QEasingCurve_new2(void* other) {
	return new QEasingCurve(*(QEasingCurve*)other);
}
void* QEasingCurve_new3(int type_) {
	return new QEasingCurve((QEasingCurve::Type)(type_));
}
void* QEasingCurve_OperatorAssign(void* self, void* other) {
	auto&& __r = ((QEasingCurve*)self)->operator=(*(QEasingCurve*)other);
	return (void*)&(__r);
}
void QEasingCurve_swap(void* self, void* other) {
	((QEasingCurve*)self)->swap(*(QEasingCurve*)other);
}
double QEasingCurve_amplitude(void* self) {
	auto&& __r = ((QEasingCurve*)self)->amplitude();
	return __r;
}
void QEasingCurve_setAmplitude(void* self, double amplitude) {
	((QEasingCurve*)self)->setAmplitude(amplitude);
}
double QEasingCurve_period(void* self) {
	auto&& __r = ((QEasingCurve*)self)->period();
	return __r;
}
void QEasingCurve_setPeriod(void* self, double period) {
	((QEasingCurve*)self)->setPeriod(period);
}
double QEasingCurve_overshoot(void* self) {
	auto&& __r = ((QEasingCurve*)self)->overshoot();
	return __r;
}
void QEasingCurve_setOvershoot(void* self, double overshoot) {
	((QEasingCurve*)self)->setOvershoot(overshoot);
}
void QEasingCurve_addCubicBezierSegment(void* self, void* c1, void* c2, void* endPoint) {
	((QEasingCurve*)self)->addCubicBezierSegment(*(QPointF*)c1, *(QPointF*)c2, *(QPointF*)endPoint);
}
void QEasingCurve_addTCBSegment(void* self, void* nextPoint, double t, double c, double b) {
	((QEasingCurve*)self)->addTCBSegment(*(QPointF*)nextPoint, t, c, b);
}
void** QEasingCurve_toCubicSpline(void* self, long long* __out_len) {
	auto&& __r = ((QEasingCurve*)self)->toCubicSpline();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QPointF(__e);
	return __arr;
}
int QEasingCurve_type(void* self) {
	auto&& __r = ((QEasingCurve*)self)->type();
	return (int)(__r);
}
void QEasingCurve_setType(void* self, int type_) {
	((QEasingCurve*)self)->setType((QEasingCurve::Type)(type_));
}
double QEasingCurve_valueForProgress(void* self, double progress) {
	auto&& __r = ((QEasingCurve*)self)->valueForProgress(progress);
	return __r;
}
int QEasingCurve_OperatorEqual(void* self, void* rhs) {
	auto&& __r = (*(QEasingCurve*)self) == (*(QEasingCurve*)rhs);
	return __r;
}
int QEasingCurve_OperatorNotEqual(void* self, void* rhs) {
	auto&& __r = (*(QEasingCurve*)self) != (*(QEasingCurve*)rhs);
	return __r;
}
void QEasingCurve_delete(void* self) { delete (QEasingCurve*)self; }

}

#include <QtCore/qlibrary.h>
#include <QtCore/qobject.h>
#include <QtCore/qstring.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qvariant.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qlibrary.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QLibrary__Go : public QLibrary {
public:
	long long __goId;
	unsigned char __goOn[8] = {};
	template<typename... A> QLibrary__Go(long long id, A&&... a) : QLibrary(std::forward<A>(a)...), __goId(id) {}
	using QLibrary::sender;
	using QLibrary::senderSignalIndex;
	using QLibrary::receivers;
	using QLibrary::isSignalConnected;
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QLibrary::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QLibrary::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QLibrary::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QLibrary::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QLibrary::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QLibrary::metaObject();
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QLibrary::event(event);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QLibrary::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QLibrary::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QLibrary::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QLibrary::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QLibrary::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QLibrary::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QLibrary;
static std::mutex __connsMu_QLibrary;

extern "C" {

void* QLibrary_base(void* self) { return (void*)static_cast<QObject*>((QLibrary*)self); }
void QLibrary__set_override(void* self, int idx) { ((QLibrary__Go*)self)->__goOn[idx] = 1; }
void* QLibrary_fromQObject(void* o) { return (void*)qobject_cast<QLibrary*>((QObject*)o); }
void* QLibrary_new(long long __goid) {
	return new QLibrary__Go(__goid);
}
void* QLibrary_new2(long long __goid, const char* fileName, long long fileName_len) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	return new QLibrary__Go(__goid, fileName_qstr);
}
void* QLibrary_new3(long long __goid, const char* fileName, long long fileName_len, int verNum) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	return new QLibrary__Go(__goid, fileName_qstr, verNum);
}
void* QLibrary_new4(long long __goid, const char* fileName, long long fileName_len, const char* version, long long version_len) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	QString version_qstr = QString::fromUtf8(version, version_len);
	return new QLibrary__Go(__goid, fileName_qstr, version_qstr);
}
void* QLibrary_new5(long long __goid, void* parent) {
	return new QLibrary__Go(__goid, (QObject*)parent);
}
void* QLibrary_new6(long long __goid, const char* fileName, long long fileName_len, void* parent) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	return new QLibrary__Go(__goid, fileName_qstr, (QObject*)parent);
}
void* QLibrary_new7(long long __goid, const char* fileName, long long fileName_len, int verNum, void* parent) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	return new QLibrary__Go(__goid, fileName_qstr, verNum, (QObject*)parent);
}
void* QLibrary_new8(long long __goid, const char* fileName, long long fileName_len, const char* version, long long version_len, void* parent) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	QString version_qstr = QString::fromUtf8(version, version_len);
	return new QLibrary__Go(__goid, fileName_qstr, version_qstr, (QObject*)parent);
}
char* QLibrary_tr(const char* s, long long* __out_len) {
	auto&& __r = QLibrary::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QLibrary_load(void* self) {
	auto&& __r = ((QLibrary*)self)->load();
	return __r;
}
int QLibrary_unload(void* self) {
	auto&& __r = ((QLibrary*)self)->unload();
	return __r;
}
int QLibrary_isLoaded(void* self) {
	auto&& __r = ((QLibrary*)self)->isLoaded();
	return __r;
}
int QLibrary_isLibrary(const char* fileName, long long fileName_len) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	auto&& __r = QLibrary::isLibrary(fileName_qstr);
	return __r;
}
void QLibrary_setFileName(void* self, const char* fileName, long long fileName_len) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	((QLibrary*)self)->setFileName(fileName_qstr);
}
char* QLibrary_fileName(void* self, long long* __out_len) {
	auto&& __r = ((QLibrary*)self)->fileName();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QLibrary_setFileNameAndVersion(void* self, const char* fileName, long long fileName_len, int verNum) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	((QLibrary*)self)->setFileNameAndVersion(fileName_qstr, verNum);
}
void QLibrary_setFileNameAndVersion2(void* self, const char* fileName, long long fileName_len, const char* version, long long version_len) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	QString version_qstr = QString::fromUtf8(version, version_len);
	((QLibrary*)self)->setFileNameAndVersion(fileName_qstr, version_qstr);
}
char* QLibrary_errorString(void* self, long long* __out_len) {
	auto&& __r = ((QLibrary*)self)->errorString();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QLibrary_setLoadHints(void* self, int hints) {
	((QLibrary*)self)->setLoadHints(QFlags<QLibrary::LoadHint>(hints));
}
int QLibrary_loadHints(void* self) {
	auto&& __r = ((QLibrary*)self)->loadHints();
	return (int)(__r);
}
void* QLibrary_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QLibrary::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QLibrary_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QLibrary::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QLibrary_sender(void* self) {
	auto&& __r = ((QLibrary__Go*)self)->sender();
	return (void*)(__r);
}
int QLibrary_senderSignalIndex(void* self) {
	auto&& __r = ((QLibrary__Go*)self)->senderSignalIndex();
	return __r;
}
int QLibrary_receivers(void* self, const char* signal) {
	auto&& __r = ((QLibrary__Go*)self)->receivers(signal);
	return __r;
}
int QLibrary_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QLibrary__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QLibrary_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QLibrary::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLibrary_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QLibrary::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QLibrary_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QLibrary::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QLibrary_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QLibrary::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QLibrary_metaObject_default(void* self) {
	auto&& __r = ((QLibrary*)self)->QLibrary::metaObject();
	return (void*)(__r);
}
int QLibrary_event_default(void* self, void* event) {
	auto&& __r = ((QLibrary*)self)->QLibrary::event((QEvent*)event);
	return __r;
}
int QLibrary_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QLibrary*)self)->QLibrary::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QLibrary_timerEvent_default(void* self, void* event) {
	((QLibrary__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QLibrary_childEvent_default(void* self, void* event) {
	((QLibrary__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QLibrary_customEvent_default(void* self, void* event) {
	((QLibrary__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QLibrary_connectNotify_default(void* self, void* signal) {
	((QLibrary__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QLibrary_disconnectNotify_default(void* self, void* signal) {
	((QLibrary__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QLibrary_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QLibrary*)self, &QLibrary::destroyed, (QLibrary*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QLibrary);
	__conns_QLibrary[__id] = __c;
}
void QLibrary_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QLibrary*)self, &QLibrary::destroyed, (QLibrary*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QLibrary);
	__conns_QLibrary[__id] = __c;
}
void QLibrary_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QLibrary*)self, &QLibrary::objectNameChanged, (QLibrary*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QLibrary);
	__conns_QLibrary[__id] = __c;
}
void QLibrary__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QLibrary);
		auto __it = __conns_QLibrary.find(id);
		if (__it == __conns_QLibrary.end()) return;
		__c = __it->second;
		__conns_QLibrary.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QLibrary__Go_new(long long id, void* parent) { return new QLibrary__Go(id, (QObject*)parent); }
void QLibrary_delete(void* self) { delete (QLibrary*)self; }

}

#include <QtCore/qparallelanimationgroup.h>
#include <QtCore/qanimationgroup.h>
#include <QtCore/qobject.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qabstractanimation.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qvariant.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qparallelanimationgroup.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QParallelAnimationGroup__Go : public QParallelAnimationGroup {
public:
	long long __goId;
	unsigned char __goOn[12] = {};
	template<typename... A> QParallelAnimationGroup__Go(long long id, A&&... a) : QParallelAnimationGroup(std::forward<A>(a)...), __goId(id) {}
	using QParallelAnimationGroup::sender;
	using QParallelAnimationGroup::senderSignalIndex;
	using QParallelAnimationGroup::receivers;
	using QParallelAnimationGroup::isSignalConnected;
	bool cppgoBase_event(QEvent* __a0) { return QParallelAnimationGroup::event(__a0); }
	void cppgoBase_updateCurrentTime(int __a0) { QParallelAnimationGroup::updateCurrentTime(__a0); }
	void cppgoBase_updateState(int __a0, int __a1) { QParallelAnimationGroup::updateState((QAbstractAnimation::State)__a0, (QAbstractAnimation::State)__a1); }
	void cppgoBase_updateDirection(int __a0) { QParallelAnimationGroup::updateDirection((QAbstractAnimation::Direction)__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QParallelAnimationGroup::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QParallelAnimationGroup::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QParallelAnimationGroup::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QParallelAnimationGroup::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QParallelAnimationGroup::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QParallelAnimationGroup::metaObject();
	}
	int duration() const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8434675112194400115LL, &__h); if (__h) return __r; }
		return QParallelAnimationGroup::duration();
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QParallelAnimationGroup::event(event);
	}
	void updateCurrentTime(int currentTime) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, -1342418961870150968LL, currentTime, &__h); if (__h) return; }
		QParallelAnimationGroup::updateCurrentTime(currentTime);
	}
	void updateState(QAbstractAnimation::State newState, QAbstractAnimation::State oldState) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_int_int) { ((void(*)(long long, long long, int, int, int*))cppgo_virt_void_int_int)(__goId, 7692482089134926569LL, (int)(newState), (int)(oldState), &__h); if (__h) return; }
		QParallelAnimationGroup::updateState(newState, oldState);
	}
	void updateDirection(QAbstractAnimation::Direction direction) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, -3231888725407197643LL, (int)(direction), &__h); if (__h) return; }
		QParallelAnimationGroup::updateDirection(direction);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QParallelAnimationGroup::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QParallelAnimationGroup::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QParallelAnimationGroup::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QParallelAnimationGroup::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QParallelAnimationGroup::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QParallelAnimationGroup::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QParallelAnimationGroup;
static std::mutex __connsMu_QParallelAnimationGroup;

extern "C" {

void* QParallelAnimationGroup_base(void* self) { return (void*)static_cast<QAnimationGroup*>((QParallelAnimationGroup*)self); }
void QParallelAnimationGroup_chain(void* self, void** out) {
	QParallelAnimationGroup* __p = (QParallelAnimationGroup*)self;
	QAnimationGroup* __b0 = static_cast<QAnimationGroup*>(__p); out[0] = (void*)__b0;
	QAbstractAnimation* __b1 = static_cast<QAbstractAnimation*>(__b0); out[1] = (void*)__b1;
	QObject* __b2 = static_cast<QObject*>(__b1); out[2] = (void*)__b2;
}
void QParallelAnimationGroup__set_override(void* self, int idx) { ((QParallelAnimationGroup__Go*)self)->__goOn[idx] = 1; }
void* QParallelAnimationGroup_fromQObject(void* o) { return (void*)qobject_cast<QParallelAnimationGroup*>((QObject*)o); }
void* QParallelAnimationGroup_new(long long __goid) {
	return new QParallelAnimationGroup__Go(__goid);
}
void* QParallelAnimationGroup_new2(long long __goid, void* parent) {
	return new QParallelAnimationGroup__Go(__goid, (QObject*)parent);
}
char* QParallelAnimationGroup_tr(const char* s, long long* __out_len) {
	auto&& __r = QParallelAnimationGroup::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QParallelAnimationGroup_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QParallelAnimationGroup::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QParallelAnimationGroup_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QParallelAnimationGroup::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QParallelAnimationGroup_sender(void* self) {
	auto&& __r = ((QParallelAnimationGroup__Go*)self)->sender();
	return (void*)(__r);
}
int QParallelAnimationGroup_senderSignalIndex(void* self) {
	auto&& __r = ((QParallelAnimationGroup__Go*)self)->senderSignalIndex();
	return __r;
}
int QParallelAnimationGroup_receivers(void* self, const char* signal) {
	auto&& __r = ((QParallelAnimationGroup__Go*)self)->receivers(signal);
	return __r;
}
int QParallelAnimationGroup_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QParallelAnimationGroup__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QParallelAnimationGroup_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QParallelAnimationGroup::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QParallelAnimationGroup_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QParallelAnimationGroup::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QParallelAnimationGroup_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QParallelAnimationGroup::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QParallelAnimationGroup_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QParallelAnimationGroup::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QParallelAnimationGroup_metaObject_default(void* self) {
	auto&& __r = ((QParallelAnimationGroup*)self)->QParallelAnimationGroup::metaObject();
	return (void*)(__r);
}
int QParallelAnimationGroup_duration_default(void* self) {
	auto&& __r = ((QParallelAnimationGroup*)self)->QParallelAnimationGroup::duration();
	return __r;
}
int QParallelAnimationGroup_event_default(void* self, void* event) {
	auto&& __r = ((QParallelAnimationGroup__Go*)self)->cppgoBase_event((QEvent*)event);
	return __r;
}
void QParallelAnimationGroup_updateCurrentTime_default(void* self, int currentTime) {
	((QParallelAnimationGroup__Go*)self)->cppgoBase_updateCurrentTime(currentTime);
}
void QParallelAnimationGroup_updateState_default(void* self, int newState, int oldState) {
	((QParallelAnimationGroup__Go*)self)->cppgoBase_updateState(newState, oldState);
}
void QParallelAnimationGroup_updateDirection_default(void* self, int direction) {
	((QParallelAnimationGroup__Go*)self)->cppgoBase_updateDirection(direction);
}
int QParallelAnimationGroup_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QParallelAnimationGroup*)self)->QParallelAnimationGroup::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QParallelAnimationGroup_timerEvent_default(void* self, void* event) {
	((QParallelAnimationGroup__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QParallelAnimationGroup_childEvent_default(void* self, void* event) {
	((QParallelAnimationGroup__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QParallelAnimationGroup_customEvent_default(void* self, void* event) {
	((QParallelAnimationGroup__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QParallelAnimationGroup_connectNotify_default(void* self, void* signal) {
	((QParallelAnimationGroup__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QParallelAnimationGroup_disconnectNotify_default(void* self, void* signal) {
	((QParallelAnimationGroup__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QParallelAnimationGroup_connect_finished(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QParallelAnimationGroup*)self, &QParallelAnimationGroup::finished, (QParallelAnimationGroup*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QParallelAnimationGroup);
	__conns_QParallelAnimationGroup[__id] = __c;
}
void QParallelAnimationGroup_connect_stateChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QParallelAnimationGroup*)self, &QParallelAnimationGroup::stateChanged, (QParallelAnimationGroup*)self, [__id](QAbstractAnimation::State newState, QAbstractAnimation::State oldState){
		if (cppgo_sig_void_int_int) ((void(*)(long long, int, int))cppgo_sig_void_int_int)(__id, (int)(newState), (int)(oldState));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QParallelAnimationGroup);
	__conns_QParallelAnimationGroup[__id] = __c;
}
void QParallelAnimationGroup_connect_currentLoopChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QParallelAnimationGroup*)self, &QParallelAnimationGroup::currentLoopChanged, (QParallelAnimationGroup*)self, [__id](int currentLoop){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, currentLoop);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QParallelAnimationGroup);
	__conns_QParallelAnimationGroup[__id] = __c;
}
void QParallelAnimationGroup_connect_directionChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QParallelAnimationGroup*)self, &QParallelAnimationGroup::directionChanged, (QParallelAnimationGroup*)self, [__id](QAbstractAnimation::Direction arg1){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, (int)(arg1));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QParallelAnimationGroup);
	__conns_QParallelAnimationGroup[__id] = __c;
}
void QParallelAnimationGroup_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QParallelAnimationGroup*)self, &QParallelAnimationGroup::destroyed, (QParallelAnimationGroup*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QParallelAnimationGroup);
	__conns_QParallelAnimationGroup[__id] = __c;
}
void QParallelAnimationGroup_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QParallelAnimationGroup*)self, &QParallelAnimationGroup::destroyed, (QParallelAnimationGroup*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QParallelAnimationGroup);
	__conns_QParallelAnimationGroup[__id] = __c;
}
void QParallelAnimationGroup_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QParallelAnimationGroup*)self, &QParallelAnimationGroup::objectNameChanged, (QParallelAnimationGroup*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QParallelAnimationGroup);
	__conns_QParallelAnimationGroup[__id] = __c;
}
void QParallelAnimationGroup__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QParallelAnimationGroup);
		auto __it = __conns_QParallelAnimationGroup.find(id);
		if (__it == __conns_QParallelAnimationGroup.end()) return;
		__c = __it->second;
		__conns_QParallelAnimationGroup.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QParallelAnimationGroup__Go_new(long long id, void* parent) { return new QParallelAnimationGroup__Go(id, (QObject*)parent); }
void QParallelAnimationGroup_delete(void* self) { delete (QParallelAnimationGroup*)self; }

}

#include <QtCore/qreadwritelock.h>
#include <QtCore/qdeadlinetimer.h>
#include "gen_qreadwritelock.h"

extern "C" {

void* QReadWriteLock_base(void* self) { return (void*)static_cast<QBasicReadWriteLock*>((QReadWriteLock*)self); }
void* QReadWriteLock_new() {
	return new QReadWriteLock();
}
void* QReadWriteLock_new2(int recursionMode) {
	return new QReadWriteLock((QReadWriteLock::RecursionMode)(recursionMode));
}
int QReadWriteLock_tryLockForRead(void* self, int timeout) {
	auto&& __r = ((QReadWriteLock*)self)->tryLockForRead(timeout);
	return __r;
}
int QReadWriteLock_tryLockForWrite(void* self, int timeout) {
	auto&& __r = ((QReadWriteLock*)self)->tryLockForWrite(timeout);
	return __r;
}
void QReadWriteLock_delete(void* self) { delete (QReadWriteLock*)self; }

}

#include <QtCore/qstringlistmodel.h>
#include <QtCore/qabstractitemmodel.h>
#include <QtCore/qobject.h>
#include <QtCore/qstring.h>
#include <QtCore/qobjectdefs.h>
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
#include "gen_qstringlistmodel.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QStringListModel__Go : public QStringListModel {
public:
	long long __goId;
	unsigned char __goOn[39] = {};
	template<typename... A> QStringListModel__Go(long long id, A&&... a) : QStringListModel(std::forward<A>(a)...), __goId(id) {}
	using QStringListModel::createIndex;
	using QStringListModel::encodeData;
	using QStringListModel::decodeData;
	using QStringListModel::beginInsertRows;
	using QStringListModel::endInsertRows;
	using QStringListModel::beginRemoveRows;
	using QStringListModel::endRemoveRows;
	using QStringListModel::beginMoveRows;
	using QStringListModel::endMoveRows;
	using QStringListModel::beginInsertColumns;
	using QStringListModel::endInsertColumns;
	using QStringListModel::beginRemoveColumns;
	using QStringListModel::endRemoveColumns;
	using QStringListModel::beginMoveColumns;
	using QStringListModel::endMoveColumns;
	using QStringListModel::beginResetModel;
	using QStringListModel::endResetModel;
	using QStringListModel::changePersistentIndex;
	using QStringListModel::changePersistentIndexList;
	using QStringListModel::persistentIndexList;
	using QStringListModel::sender;
	using QStringListModel::senderSignalIndex;
	using QStringListModel::receivers;
	using QStringListModel::isSignalConnected;
	void cppgoBase_resetInternalData() { QStringListModel::resetInternalData(); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QStringListModel::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QStringListModel::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QStringListModel::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QStringListModel::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QStringListModel::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QStringListModel::metaObject();
	}
	int rowCount(const QModelIndex& parent) const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1960960171956884010LL, (void*)&parent, &__h); if (__h) return __r; }
		return QStringListModel::rowCount(parent);
	}
	QModelIndex sibling(int row, int column, const QModelIndex& idx) const override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_voidp_int_int_voidp) { void* __r = ((void*(*)(long long, long long, int, int, void*, int*))cppgo_virt_voidp_int_int_voidp)(__goId, 397329346813782521LL, row, column, (void*)&idx, &__h); if (__h) return *(QModelIndex*)__r; }
		return QStringListModel::sibling(row, column, idx);
	}
	QVariant data(const QModelIndex& index, int role) const override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_voidp_voidp_int) { void* __r = ((void*(*)(long long, long long, void*, int, int*))cppgo_virt_voidp_voidp_int)(__goId, -8837375942098466299LL, (void*)&index, role, &__h); if (__h) return *(QVariant*)__r; }
		return QStringListModel::data(index, role);
	}
	bool setData(const QModelIndex& index, const QVariant& value, int role) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_int_voidp_voidp_int) { int __r = ((int(*)(long long, long long, void*, void*, int, int*))cppgo_virt_int_voidp_voidp_int)(__goId, 35191454575202219LL, (void*)&index, (void*)&value, role, &__h); if (__h) return __r; }
		return QStringListModel::setData(index, value, role);
	}
	bool clearItemData(const QModelIndex& index) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 8914190500520542409LL, (void*)&index, &__h); if (__h) return __r; }
		return QStringListModel::clearItemData(index);
	}
	QFlags<Qt::ItemFlag> flags(const QModelIndex& index) const override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1703382833545304812LL, (void*)&index, &__h); if (__h) return QFlags<Qt::ItemFlag>(__r); }
		return QStringListModel::flags(index);
	}
	bool insertRows(int row, int count, const QModelIndex& parent) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_int_int_int_voidp) { int __r = ((int(*)(long long, long long, int, int, void*, int*))cppgo_virt_int_int_int_voidp)(__goId, -2775729825962395603LL, row, count, (void*)&parent, &__h); if (__h) return __r; }
		return QStringListModel::insertRows(row, count, parent);
	}
	bool removeRows(int row, int count, const QModelIndex& parent) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_int_int_int_voidp) { int __r = ((int(*)(long long, long long, int, int, void*, int*))cppgo_virt_int_int_int_voidp)(__goId, 5104072549182222656LL, row, count, (void*)&parent, &__h); if (__h) return __r; }
		return QStringListModel::removeRows(row, count, parent);
	}
	bool moveRows(const QModelIndex& sourceParent, int sourceRow, int count, const QModelIndex& destinationParent, int destinationChild) override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_int_voidp_int_int_voidp_int) { int __r = ((int(*)(long long, long long, void*, int, int, void*, int, int*))cppgo_virt_int_voidp_int_int_voidp_int)(__goId, -6074819356947405415LL, (void*)&sourceParent, sourceRow, count, (void*)&destinationParent, destinationChild, &__h); if (__h) return __r; }
		return QStringListModel::moveRows(sourceParent, sourceRow, count, destinationParent, destinationChild);
	}
	void sort(int column, Qt::SortOrder order) override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_void_int_int) { ((void(*)(long long, long long, int, int, int*))cppgo_virt_void_int_int)(__goId, 6912441694453382161LL, column, (int)(order), &__h); if (__h) return; }
		QStringListModel::sort(column, order);
	}
	QFlags<Qt::DropAction> supportedDropActions() const override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -7784642176186821205LL, &__h); if (__h) return QFlags<Qt::DropAction>(__r); }
		return QStringListModel::supportedDropActions();
	}
	QModelIndex index(int row, int column, const QModelIndex& parent) const override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_voidp_int_int_voidp) { void* __r = ((void*(*)(long long, long long, int, int, void*, int*))cppgo_virt_voidp_int_int_voidp)(__goId, -8948777187306027381LL, row, column, (void*)&parent, &__h); if (__h) return *(QModelIndex*)__r; }
		return QStringListModel::index(row, column, parent);
	}
	bool dropMimeData(const QMimeData* data, Qt::DropAction action, int row, int column, const QModelIndex& parent) override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_int_voidp_int_int_int_voidp) { int __r = ((int(*)(long long, long long, void*, int, int, int, void*, int*))cppgo_virt_int_voidp_int_int_int_voidp)(__goId, -9049994684305682212LL, (void*)data, (int)(action), row, column, (void*)&parent, &__h); if (__h) return __r; }
		return QStringListModel::dropMimeData(data, action, row, column, parent);
	}
	QVariant headerData(int section, Qt::Orientation orientation, int role) const override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_voidp_int_int_int) { void* __r = ((void*(*)(long long, long long, int, int, int, int*))cppgo_virt_voidp_int_int_int)(__goId, -3297140021896693688LL, section, (int)(orientation), role, &__h); if (__h) return *(QVariant*)__r; }
		return QStringListModel::headerData(section, orientation, role);
	}
	bool setHeaderData(int section, Qt::Orientation orientation, const QVariant& value, int role) override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_int_int_int_voidp_int) { int __r = ((int(*)(long long, long long, int, int, void*, int, int*))cppgo_virt_int_int_int_voidp_int)(__goId, 4852200306590054850LL, section, (int)(orientation), (void*)&value, role, &__h); if (__h) return __r; }
		return QStringListModel::setHeaderData(section, orientation, value, role);
	}
	QList<QString> mimeTypes() const override {
		int __h = 0;
		if (__goOn[16] && cppgo_virtp_voidp) {
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
		return QStringListModel::mimeTypes();
	}
	QMimeData* mimeData(const QList<QModelIndex>& indexes) const override {
		int __h = 0;
		std::vector<void*> indexes_v; indexes_v.reserve(indexes.size());
		for (long long __i = 0; __i < indexes.size(); __i++) indexes_v.push_back((void*)&indexes[__i]);
		if (__goOn[17] && cppgo_virt_voidp_voidpp_longlong) { void* __r = ((void*(*)(long long, long long, void**, long long, int*))cppgo_virt_voidp_voidpp_longlong)(__goId, 5175297091024974749LL, indexes_v.data(), (long long)indexes.size(), &__h); if (__h) return (QMimeData*)__r; }
		return QStringListModel::mimeData(indexes);
	}
	bool canDropMimeData(const QMimeData* data, Qt::DropAction action, int row, int column, const QModelIndex& parent) const override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_int_voidp_int_int_int_voidp) { int __r = ((int(*)(long long, long long, void*, int, int, int, void*, int*))cppgo_virt_int_voidp_int_int_int_voidp)(__goId, 731539533644753052LL, (void*)data, (int)(action), row, column, (void*)&parent, &__h); if (__h) return __r; }
		return QStringListModel::canDropMimeData(data, action, row, column, parent);
	}
	QFlags<Qt::DropAction> supportedDragActions() const override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, 3421601386088570762LL, &__h); if (__h) return QFlags<Qt::DropAction>(__r); }
		return QStringListModel::supportedDragActions();
	}
	bool insertColumns(int column, int count, const QModelIndex& parent) override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_int_int_int_voidp) { int __r = ((int(*)(long long, long long, int, int, void*, int*))cppgo_virt_int_int_int_voidp)(__goId, -8973293846230792237LL, column, count, (void*)&parent, &__h); if (__h) return __r; }
		return QStringListModel::insertColumns(column, count, parent);
	}
	bool removeColumns(int column, int count, const QModelIndex& parent) override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_int_int_int_voidp) { int __r = ((int(*)(long long, long long, int, int, void*, int*))cppgo_virt_int_int_int_voidp)(__goId, 2078548737056540420LL, column, count, (void*)&parent, &__h); if (__h) return __r; }
		return QStringListModel::removeColumns(column, count, parent);
	}
	bool moveColumns(const QModelIndex& sourceParent, int sourceColumn, int count, const QModelIndex& destinationParent, int destinationChild) override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_int_voidp_int_int_voidp_int) { int __r = ((int(*)(long long, long long, void*, int, int, void*, int, int*))cppgo_virt_int_voidp_int_int_voidp_int)(__goId, -6321007591388309905LL, (void*)&sourceParent, sourceColumn, count, (void*)&destinationParent, destinationChild, &__h); if (__h) return __r; }
		return QStringListModel::moveColumns(sourceParent, sourceColumn, count, destinationParent, destinationChild);
	}
	void fetchMore(const QModelIndex& parent) override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3553294593611130806LL, (void*)&parent, &__h); if (__h) return; }
		QStringListModel::fetchMore(parent);
	}
	bool canFetchMore(const QModelIndex& parent) const override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 8639262066473964150LL, (void*)&parent, &__h); if (__h) return __r; }
		return QStringListModel::canFetchMore(parent);
	}
	QModelIndex buddy(const QModelIndex& index) const override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 6112860112248786499LL, (void*)&index, &__h); if (__h) return *(QModelIndex*)__r; }
		return QStringListModel::buddy(index);
	}
	QList<QModelIndex> match(const QModelIndex& start, int role, const QVariant& value, int hits, QFlags<Qt::MatchFlag> flags) const override {
		int __h = 0;
		if (__goOn[26] && cppgo_virtp_voidp_voidp_int_voidp_int_int) {
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
		return QStringListModel::match(start, role, value, hits, flags);
	}
	QSize span(const QModelIndex& index) const override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, -8395342916587040543LL, (void*)&index, &__h); if (__h) return *(QSize*)__r; }
		return QStringListModel::span(index);
	}
	void multiData(const QModelIndex& index, QModelRoleDataSpan roleDataSpan) const override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_void_voidp_voidp) { ((void(*)(long long, long long, void*, void*, int*))cppgo_virt_void_voidp_voidp)(__goId, 731311060944096642LL, (void*)&index, (void*)&roleDataSpan, &__h); if (__h) return; }
		QStringListModel::multiData(index, roleDataSpan);
	}
	bool submit() override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -3788331537442685179LL, &__h); if (__h) return __r; }
		return QStringListModel::submit();
	}
	void revert() override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, -1522951607536117149LL, &__h); if (__h) return; }
		QStringListModel::revert();
	}
	void resetInternalData() override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, -4510675770784051183LL, &__h); if (__h) return; }
		QStringListModel::resetInternalData();
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[32] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QStringListModel::event(event);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[33] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QStringListModel::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[34] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QStringListModel::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[35] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QStringListModel::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[36] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QStringListModel::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[37] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QStringListModel::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[38] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QStringListModel::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QStringListModel;
static std::mutex __connsMu_QStringListModel;

extern "C" {

void* QStringListModel_base(void* self) { return (void*)static_cast<QAbstractListModel*>((QStringListModel*)self); }
void QStringListModel_chain(void* self, void** out) {
	QStringListModel* __p = (QStringListModel*)self;
	QAbstractListModel* __b0 = static_cast<QAbstractListModel*>(__p); out[0] = (void*)__b0;
	QAbstractItemModel* __b1 = static_cast<QAbstractItemModel*>(__b0); out[1] = (void*)__b1;
	QObject* __b2 = static_cast<QObject*>(__b1); out[2] = (void*)__b2;
}
void QStringListModel__set_override(void* self, int idx) { ((QStringListModel__Go*)self)->__goOn[idx] = 1; }
void* QStringListModel_fromQObject(void* o) { return (void*)qobject_cast<QStringListModel*>((QObject*)o); }
void* QStringListModel_new(long long __goid) {
	return new QStringListModel__Go(__goid);
}
void* QStringListModel_new2(long long __goid, char** strings, long long strings_len) {
	QList<QString> strings_list; strings_list.reserve(strings_len);
	for (long long __i = 0; __i < strings_len; __i++) strings_list.append(QString::fromUtf8(strings[__i]));
	return new QStringListModel__Go(__goid, strings_list);
}
void* QStringListModel_new3(long long __goid, void* parent) {
	return new QStringListModel__Go(__goid, (QObject*)parent);
}
void* QStringListModel_new4(long long __goid, char** strings, long long strings_len, void* parent) {
	QList<QString> strings_list; strings_list.reserve(strings_len);
	for (long long __i = 0; __i < strings_len; __i++) strings_list.append(QString::fromUtf8(strings[__i]));
	return new QStringListModel__Go(__goid, strings_list, (QObject*)parent);
}
char* QStringListModel_tr(const char* s, long long* __out_len) {
	auto&& __r = QStringListModel::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char** QStringListModel_stringList(void* self, long long* __out_len) {
	auto&& __r = ((QStringListModel*)self)->stringList();
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
void QStringListModel_setStringList(void* self, char** strings, long long strings_len) {
	QList<QString> strings_list; strings_list.reserve(strings_len);
	for (long long __i = 0; __i < strings_len; __i++) strings_list.append(QString::fromUtf8(strings[__i]));
	((QStringListModel*)self)->setStringList(strings_list);
}
void* QStringListModel_createIndex(void* self, int row, int column) {
	auto&& __r = ((QStringListModel__Go*)self)->createIndex(row, column);
	return new QModelIndex(__r);
}
void* QStringListModel_createIndex2(void* self, int row, int column, unsigned long long id) {
	auto&& __r = ((QStringListModel__Go*)self)->createIndex(row, column, id);
	return new QModelIndex(__r);
}
void QStringListModel_encodeData(void* self, void** indexes, long long indexes_len, void* stream) {
	QList<QModelIndex> indexes_list; indexes_list.reserve(indexes_len);
	for (long long __i = 0; __i < indexes_len; __i++) indexes_list.append(*(QModelIndex*)indexes[__i]);
	((QStringListModel__Go*)self)->encodeData(indexes_list, *(QDataStream*)stream);
}
int QStringListModel_decodeData(void* self, int row, int column, void* parent, void* stream) {
	auto&& __r = ((QStringListModel__Go*)self)->decodeData(row, column, *(QModelIndex*)parent, *(QDataStream*)stream);
	return __r;
}
void QStringListModel_beginInsertRows(void* self, void* parent, int first, int last) {
	((QStringListModel__Go*)self)->beginInsertRows(*(QModelIndex*)parent, first, last);
}
void QStringListModel_endInsertRows(void* self) {
	((QStringListModel__Go*)self)->endInsertRows();
}
void QStringListModel_beginRemoveRows(void* self, void* parent, int first, int last) {
	((QStringListModel__Go*)self)->beginRemoveRows(*(QModelIndex*)parent, first, last);
}
void QStringListModel_endRemoveRows(void* self) {
	((QStringListModel__Go*)self)->endRemoveRows();
}
int QStringListModel_beginMoveRows(void* self, void* sourceParent, int sourceFirst, int sourceLast, void* destinationParent, int destinationRow) {
	auto&& __r = ((QStringListModel__Go*)self)->beginMoveRows(*(QModelIndex*)sourceParent, sourceFirst, sourceLast, *(QModelIndex*)destinationParent, destinationRow);
	return __r;
}
void QStringListModel_endMoveRows(void* self) {
	((QStringListModel__Go*)self)->endMoveRows();
}
void QStringListModel_beginInsertColumns(void* self, void* parent, int first, int last) {
	((QStringListModel__Go*)self)->beginInsertColumns(*(QModelIndex*)parent, first, last);
}
void QStringListModel_endInsertColumns(void* self) {
	((QStringListModel__Go*)self)->endInsertColumns();
}
void QStringListModel_beginRemoveColumns(void* self, void* parent, int first, int last) {
	((QStringListModel__Go*)self)->beginRemoveColumns(*(QModelIndex*)parent, first, last);
}
void QStringListModel_endRemoveColumns(void* self) {
	((QStringListModel__Go*)self)->endRemoveColumns();
}
int QStringListModel_beginMoveColumns(void* self, void* sourceParent, int sourceFirst, int sourceLast, void* destinationParent, int destinationColumn) {
	auto&& __r = ((QStringListModel__Go*)self)->beginMoveColumns(*(QModelIndex*)sourceParent, sourceFirst, sourceLast, *(QModelIndex*)destinationParent, destinationColumn);
	return __r;
}
void QStringListModel_endMoveColumns(void* self) {
	((QStringListModel__Go*)self)->endMoveColumns();
}
void QStringListModel_beginResetModel(void* self) {
	((QStringListModel__Go*)self)->beginResetModel();
}
void QStringListModel_endResetModel(void* self) {
	((QStringListModel__Go*)self)->endResetModel();
}
void QStringListModel_changePersistentIndex(void* self, void* from, void* to) {
	((QStringListModel__Go*)self)->changePersistentIndex(*(QModelIndex*)from, *(QModelIndex*)to);
}
void QStringListModel_changePersistentIndexList(void* self, void** from, long long from_len, void** to, long long to_len) {
	QList<QModelIndex> from_list; from_list.reserve(from_len);
	for (long long __i = 0; __i < from_len; __i++) from_list.append(*(QModelIndex*)from[__i]);
	QList<QModelIndex> to_list; to_list.reserve(to_len);
	for (long long __i = 0; __i < to_len; __i++) to_list.append(*(QModelIndex*)to[__i]);
	((QStringListModel__Go*)self)->changePersistentIndexList(from_list, to_list);
}
void** QStringListModel_persistentIndexList(void* self, long long* __out_len) {
	auto&& __r = ((QStringListModel__Go*)self)->persistentIndexList();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QModelIndex(__e);
	return __arr;
}
void* QStringListModel_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QStringListModel::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QStringListModel_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QStringListModel::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QStringListModel_sender(void* self) {
	auto&& __r = ((QStringListModel__Go*)self)->sender();
	return (void*)(__r);
}
int QStringListModel_senderSignalIndex(void* self) {
	auto&& __r = ((QStringListModel__Go*)self)->senderSignalIndex();
	return __r;
}
int QStringListModel_receivers(void* self, const char* signal) {
	auto&& __r = ((QStringListModel__Go*)self)->receivers(signal);
	return __r;
}
int QStringListModel_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QStringListModel__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QStringListModel_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QStringListModel::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QStringListModel_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QStringListModel::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QStringListModel_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QStringListModel::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QStringListModel_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QStringListModel::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QStringListModel_metaObject_default(void* self) {
	auto&& __r = ((QStringListModel*)self)->QStringListModel::metaObject();
	return (void*)(__r);
}
int QStringListModel_rowCount_default(void* self, void* parent) {
	auto&& __r = ((QStringListModel*)self)->QStringListModel::rowCount(*(QModelIndex*)parent);
	return __r;
}
void* QStringListModel_sibling_default(void* self, int row, int column, void* idx) {
	auto&& __r = ((QStringListModel*)self)->QStringListModel::sibling(row, column, *(QModelIndex*)idx);
	return new QModelIndex(__r);
}
void* QStringListModel_data_default(void* self, void* index, int role) {
	auto&& __r = ((QStringListModel*)self)->QStringListModel::data(*(QModelIndex*)index, role);
	return new QVariant(__r);
}
int QStringListModel_setData_default(void* self, void* index, void* value, int role) {
	auto&& __r = ((QStringListModel*)self)->QStringListModel::setData(*(QModelIndex*)index, *(QVariant*)value, role);
	return __r;
}
int QStringListModel_clearItemData_default(void* self, void* index) {
	auto&& __r = ((QStringListModel*)self)->QStringListModel::clearItemData(*(QModelIndex*)index);
	return __r;
}
int QStringListModel_flags_default(void* self, void* index) {
	auto&& __r = ((QStringListModel*)self)->QStringListModel::flags(*(QModelIndex*)index);
	return (int)(__r);
}
int QStringListModel_insertRows_default(void* self, int row, int count, void* parent) {
	auto&& __r = ((QStringListModel*)self)->QStringListModel::insertRows(row, count, *(QModelIndex*)parent);
	return __r;
}
int QStringListModel_removeRows_default(void* self, int row, int count, void* parent) {
	auto&& __r = ((QStringListModel*)self)->QStringListModel::removeRows(row, count, *(QModelIndex*)parent);
	return __r;
}
int QStringListModel_moveRows_default(void* self, void* sourceParent, int sourceRow, int count, void* destinationParent, int destinationChild) {
	auto&& __r = ((QStringListModel*)self)->QStringListModel::moveRows(*(QModelIndex*)sourceParent, sourceRow, count, *(QModelIndex*)destinationParent, destinationChild);
	return __r;
}
void QStringListModel_sort_default(void* self, int column, int order) {
	((QStringListModel*)self)->QStringListModel::sort(column, (Qt::SortOrder)(order));
}
int QStringListModel_supportedDropActions_default(void* self) {
	auto&& __r = ((QStringListModel*)self)->QStringListModel::supportedDropActions();
	return (int)(__r);
}
void* QStringListModel_index_default(void* self, int row, int column, void* parent) {
	auto&& __r = ((QStringListModel*)self)->QStringListModel::index(row, column, *(QModelIndex*)parent);
	return new QModelIndex(__r);
}
int QStringListModel_dropMimeData_default(void* self, void* data, int action, int row, int column, void* parent) {
	auto&& __r = ((QStringListModel*)self)->QStringListModel::dropMimeData((QMimeData*)data, (Qt::DropAction)(action), row, column, *(QModelIndex*)parent);
	return __r;
}
void* QStringListModel_headerData_default(void* self, int section, int orientation, int role) {
	auto&& __r = ((QStringListModel*)self)->QStringListModel::headerData(section, (Qt::Orientation)(orientation), role);
	return new QVariant(__r);
}
int QStringListModel_setHeaderData_default(void* self, int section, int orientation, void* value, int role) {
	auto&& __r = ((QStringListModel*)self)->QStringListModel::setHeaderData(section, (Qt::Orientation)(orientation), *(QVariant*)value, role);
	return __r;
}
char** QStringListModel_mimeTypes_default(void* self, long long* __out_len) {
	auto&& __r = ((QStringListModel*)self)->QStringListModel::mimeTypes();
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
void* QStringListModel_mimeData_default(void* self, void** indexes, long long indexes_len) {
	QList<QModelIndex> indexes_list; indexes_list.reserve(indexes_len);
	for (long long __i = 0; __i < indexes_len; __i++) indexes_list.append(*(QModelIndex*)indexes[__i]);
	auto&& __r = ((QStringListModel*)self)->QStringListModel::mimeData(indexes_list);
	return (void*)(__r);
}
int QStringListModel_canDropMimeData_default(void* self, void* data, int action, int row, int column, void* parent) {
	auto&& __r = ((QStringListModel*)self)->QStringListModel::canDropMimeData((QMimeData*)data, (Qt::DropAction)(action), row, column, *(QModelIndex*)parent);
	return __r;
}
int QStringListModel_supportedDragActions_default(void* self) {
	auto&& __r = ((QStringListModel*)self)->QStringListModel::supportedDragActions();
	return (int)(__r);
}
int QStringListModel_insertColumns_default(void* self, int column, int count, void* parent) {
	auto&& __r = ((QStringListModel*)self)->QStringListModel::insertColumns(column, count, *(QModelIndex*)parent);
	return __r;
}
int QStringListModel_removeColumns_default(void* self, int column, int count, void* parent) {
	auto&& __r = ((QStringListModel*)self)->QStringListModel::removeColumns(column, count, *(QModelIndex*)parent);
	return __r;
}
int QStringListModel_moveColumns_default(void* self, void* sourceParent, int sourceColumn, int count, void* destinationParent, int destinationChild) {
	auto&& __r = ((QStringListModel*)self)->QStringListModel::moveColumns(*(QModelIndex*)sourceParent, sourceColumn, count, *(QModelIndex*)destinationParent, destinationChild);
	return __r;
}
void QStringListModel_fetchMore_default(void* self, void* parent) {
	((QStringListModel*)self)->QStringListModel::fetchMore(*(QModelIndex*)parent);
}
int QStringListModel_canFetchMore_default(void* self, void* parent) {
	auto&& __r = ((QStringListModel*)self)->QStringListModel::canFetchMore(*(QModelIndex*)parent);
	return __r;
}
void* QStringListModel_buddy_default(void* self, void* index) {
	auto&& __r = ((QStringListModel*)self)->QStringListModel::buddy(*(QModelIndex*)index);
	return new QModelIndex(__r);
}
void** QStringListModel_match_default(void* self, void* start, int role, void* value, int hits, int flags, long long* __out_len) {
	auto&& __r = ((QStringListModel*)self)->QStringListModel::match(*(QModelIndex*)start, role, *(QVariant*)value, hits, QFlags<Qt::MatchFlag>(flags));
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QModelIndex(__e);
	return __arr;
}
void* QStringListModel_span_default(void* self, void* index) {
	auto&& __r = ((QStringListModel*)self)->QStringListModel::span(*(QModelIndex*)index);
	return new QSize(__r);
}
void QStringListModel_multiData_default(void* self, void* index, void* roleDataSpan) {
	((QStringListModel*)self)->QStringListModel::multiData(*(QModelIndex*)index, *(QModelRoleDataSpan*)roleDataSpan);
}
int QStringListModel_submit_default(void* self) {
	auto&& __r = ((QStringListModel*)self)->QStringListModel::submit();
	return __r;
}
void QStringListModel_revert_default(void* self) {
	((QStringListModel*)self)->QStringListModel::revert();
}
void QStringListModel_resetInternalData_default(void* self) {
	((QStringListModel__Go*)self)->cppgoBase_resetInternalData();
}
int QStringListModel_event_default(void* self, void* event) {
	auto&& __r = ((QStringListModel*)self)->QStringListModel::event((QEvent*)event);
	return __r;
}
int QStringListModel_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QStringListModel*)self)->QStringListModel::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QStringListModel_timerEvent_default(void* self, void* event) {
	((QStringListModel__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QStringListModel_childEvent_default(void* self, void* event) {
	((QStringListModel__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QStringListModel_customEvent_default(void* self, void* event) {
	((QStringListModel__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QStringListModel_connectNotify_default(void* self, void* signal) {
	((QStringListModel__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QStringListModel_disconnectNotify_default(void* self, void* signal) {
	((QStringListModel__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QStringListModel_connect_headerDataChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStringListModel*)self, &QStringListModel::headerDataChanged, (QStringListModel*)self, [__id](Qt::Orientation orientation, int first, int last){
		if (cppgo_sig_void_int_int_int) ((void(*)(long long, int, int, int))cppgo_sig_void_int_int_int)(__id, (int)(orientation), first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStringListModel);
	__conns_QStringListModel[__id] = __c;
}
void QStringListModel_connect_layoutChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStringListModel*)self, &QStringListModel::layoutChanged, (QStringListModel*)self, [__id](const QList<QPersistentModelIndex>& parents, QAbstractItemModel::LayoutChangeHint hint){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStringListModel);
	__conns_QStringListModel[__id] = __c;
}
void QStringListModel_connect_layoutChanged2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStringListModel*)self, &QStringListModel::layoutChanged, (QStringListModel*)self, [__id](const QList<QPersistentModelIndex>& parents, QAbstractItemModel::LayoutChangeHint hint){
		std::vector<void*> parents_v; parents_v.reserve(parents.size());
		for (long long __i = 0; __i < parents.size(); __i++) parents_v.push_back((void*)&parents[__i]);
		if (cppgo_sig_void_voidpp_longlong_int) ((void(*)(long long, void**, long long, int))cppgo_sig_void_voidpp_longlong_int)(__id, parents_v.data(), (long long)parents.size(), (int)(hint));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStringListModel);
	__conns_QStringListModel[__id] = __c;
}
void QStringListModel_connect_layoutAboutToBeChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStringListModel*)self, &QStringListModel::layoutAboutToBeChanged, (QStringListModel*)self, [__id](const QList<QPersistentModelIndex>& parents, QAbstractItemModel::LayoutChangeHint hint){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStringListModel);
	__conns_QStringListModel[__id] = __c;
}
void QStringListModel_connect_layoutAboutToBeChanged2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStringListModel*)self, &QStringListModel::layoutAboutToBeChanged, (QStringListModel*)self, [__id](const QList<QPersistentModelIndex>& parents, QAbstractItemModel::LayoutChangeHint hint){
		std::vector<void*> parents_v; parents_v.reserve(parents.size());
		for (long long __i = 0; __i < parents.size(); __i++) parents_v.push_back((void*)&parents[__i]);
		if (cppgo_sig_void_voidpp_longlong_int) ((void(*)(long long, void**, long long, int))cppgo_sig_void_voidpp_longlong_int)(__id, parents_v.data(), (long long)parents.size(), (int)(hint));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStringListModel);
	__conns_QStringListModel[__id] = __c;
}
void QStringListModel_connect_rowsAboutToBeInserted(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStringListModel*)self, &QStringListModel::rowsAboutToBeInserted, (QStringListModel*)self, [__id](const QModelIndex& parent, int first, int last){
		if (cppgo_sig_void_voidp_int_int) ((void(*)(long long, void*, int, int))cppgo_sig_void_voidp_int_int)(__id, (void*)&parent, first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStringListModel);
	__conns_QStringListModel[__id] = __c;
}
void QStringListModel_connect_rowsInserted(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStringListModel*)self, &QStringListModel::rowsInserted, (QStringListModel*)self, [__id](const QModelIndex& parent, int first, int last){
		if (cppgo_sig_void_voidp_int_int) ((void(*)(long long, void*, int, int))cppgo_sig_void_voidp_int_int)(__id, (void*)&parent, first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStringListModel);
	__conns_QStringListModel[__id] = __c;
}
void QStringListModel_connect_rowsAboutToBeRemoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStringListModel*)self, &QStringListModel::rowsAboutToBeRemoved, (QStringListModel*)self, [__id](const QModelIndex& parent, int first, int last){
		if (cppgo_sig_void_voidp_int_int) ((void(*)(long long, void*, int, int))cppgo_sig_void_voidp_int_int)(__id, (void*)&parent, first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStringListModel);
	__conns_QStringListModel[__id] = __c;
}
void QStringListModel_connect_rowsRemoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStringListModel*)self, &QStringListModel::rowsRemoved, (QStringListModel*)self, [__id](const QModelIndex& parent, int first, int last){
		if (cppgo_sig_void_voidp_int_int) ((void(*)(long long, void*, int, int))cppgo_sig_void_voidp_int_int)(__id, (void*)&parent, first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStringListModel);
	__conns_QStringListModel[__id] = __c;
}
void QStringListModel_connect_columnsAboutToBeInserted(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStringListModel*)self, &QStringListModel::columnsAboutToBeInserted, (QStringListModel*)self, [__id](const QModelIndex& parent, int first, int last){
		if (cppgo_sig_void_voidp_int_int) ((void(*)(long long, void*, int, int))cppgo_sig_void_voidp_int_int)(__id, (void*)&parent, first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStringListModel);
	__conns_QStringListModel[__id] = __c;
}
void QStringListModel_connect_columnsInserted(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStringListModel*)self, &QStringListModel::columnsInserted, (QStringListModel*)self, [__id](const QModelIndex& parent, int first, int last){
		if (cppgo_sig_void_voidp_int_int) ((void(*)(long long, void*, int, int))cppgo_sig_void_voidp_int_int)(__id, (void*)&parent, first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStringListModel);
	__conns_QStringListModel[__id] = __c;
}
void QStringListModel_connect_columnsAboutToBeRemoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStringListModel*)self, &QStringListModel::columnsAboutToBeRemoved, (QStringListModel*)self, [__id](const QModelIndex& parent, int first, int last){
		if (cppgo_sig_void_voidp_int_int) ((void(*)(long long, void*, int, int))cppgo_sig_void_voidp_int_int)(__id, (void*)&parent, first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStringListModel);
	__conns_QStringListModel[__id] = __c;
}
void QStringListModel_connect_columnsRemoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStringListModel*)self, &QStringListModel::columnsRemoved, (QStringListModel*)self, [__id](const QModelIndex& parent, int first, int last){
		if (cppgo_sig_void_voidp_int_int) ((void(*)(long long, void*, int, int))cppgo_sig_void_voidp_int_int)(__id, (void*)&parent, first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStringListModel);
	__conns_QStringListModel[__id] = __c;
}
void QStringListModel_connect_modelAboutToBeReset(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStringListModel*)self, &QStringListModel::modelAboutToBeReset, (QStringListModel*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStringListModel);
	__conns_QStringListModel[__id] = __c;
}
void QStringListModel_connect_modelReset(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStringListModel*)self, &QStringListModel::modelReset, (QStringListModel*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStringListModel);
	__conns_QStringListModel[__id] = __c;
}
void QStringListModel_connect_rowsAboutToBeMoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStringListModel*)self, &QStringListModel::rowsAboutToBeMoved, (QStringListModel*)self, [__id](const QModelIndex& sourceParent, int sourceStart, int sourceEnd, const QModelIndex& destinationParent, int destinationRow){
		if (cppgo_sig_void_voidp_int_int_voidp_int) ((void(*)(long long, void*, int, int, void*, int))cppgo_sig_void_voidp_int_int_voidp_int)(__id, (void*)&sourceParent, sourceStart, sourceEnd, (void*)&destinationParent, destinationRow);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStringListModel);
	__conns_QStringListModel[__id] = __c;
}
void QStringListModel_connect_rowsMoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStringListModel*)self, &QStringListModel::rowsMoved, (QStringListModel*)self, [__id](const QModelIndex& sourceParent, int sourceStart, int sourceEnd, const QModelIndex& destinationParent, int destinationRow){
		if (cppgo_sig_void_voidp_int_int_voidp_int) ((void(*)(long long, void*, int, int, void*, int))cppgo_sig_void_voidp_int_int_voidp_int)(__id, (void*)&sourceParent, sourceStart, sourceEnd, (void*)&destinationParent, destinationRow);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStringListModel);
	__conns_QStringListModel[__id] = __c;
}
void QStringListModel_connect_columnsAboutToBeMoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStringListModel*)self, &QStringListModel::columnsAboutToBeMoved, (QStringListModel*)self, [__id](const QModelIndex& sourceParent, int sourceStart, int sourceEnd, const QModelIndex& destinationParent, int destinationColumn){
		if (cppgo_sig_void_voidp_int_int_voidp_int) ((void(*)(long long, void*, int, int, void*, int))cppgo_sig_void_voidp_int_int_voidp_int)(__id, (void*)&sourceParent, sourceStart, sourceEnd, (void*)&destinationParent, destinationColumn);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStringListModel);
	__conns_QStringListModel[__id] = __c;
}
void QStringListModel_connect_columnsMoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStringListModel*)self, &QStringListModel::columnsMoved, (QStringListModel*)self, [__id](const QModelIndex& sourceParent, int sourceStart, int sourceEnd, const QModelIndex& destinationParent, int destinationColumn){
		if (cppgo_sig_void_voidp_int_int_voidp_int) ((void(*)(long long, void*, int, int, void*, int))cppgo_sig_void_voidp_int_int_voidp_int)(__id, (void*)&sourceParent, sourceStart, sourceEnd, (void*)&destinationParent, destinationColumn);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStringListModel);
	__conns_QStringListModel[__id] = __c;
}
void QStringListModel_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStringListModel*)self, &QStringListModel::destroyed, (QStringListModel*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStringListModel);
	__conns_QStringListModel[__id] = __c;
}
void QStringListModel_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStringListModel*)self, &QStringListModel::destroyed, (QStringListModel*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStringListModel);
	__conns_QStringListModel[__id] = __c;
}
void QStringListModel_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStringListModel*)self, &QStringListModel::objectNameChanged, (QStringListModel*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStringListModel);
	__conns_QStringListModel[__id] = __c;
}
void QStringListModel__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QStringListModel);
		auto __it = __conns_QStringListModel.find(id);
		if (__it == __conns_QStringListModel.end()) return;
		__c = __it->second;
		__conns_QStringListModel.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QStringListModel__Go_new(long long id, void* parent) { return new QStringListModel__Go(id, (QObject*)parent); }
void QStringListModel_delete(void* self) { delete (QStringListModel*)self; }

}

#include <QtGui/qpolygon.h>
#include <QtCore/qpoint.h>
#include <QtCore/qrect.h>
#include <QtCore/qvariant.h>
#include <QtCore/qnamespace.h>
#include <QtGui/qtransform.h>
#include "gen_qpolygon.h"

extern "C" {

void* QPolygon_new() {
	return new QPolygon();
}
void* QPolygon_new2(void** v, long long v_len) {
	QList<QPoint> v_list; v_list.reserve(v_len);
	for (long long __i = 0; __i < v_len; __i++) v_list.append(*(QPoint*)v[__i]);
	return new QPolygon(v_list);
}
void* QPolygon_new3(void* r) {
	return new QPolygon(*(QRect*)r);
}
void* QPolygon_new4(void* arg1) {
	return new QPolygon(*(QPolygon*)arg1);
}
void* QPolygon_new5(void* r, int closed) {
	return new QPolygon(*(QRect*)r, closed);
}
void QPolygon_swap(void* self, void* other) {
	((QPolygon*)self)->swap(*(QPolygon*)other);
}
void* QPolygon_ToQVariant(void* self) {
	auto&& __r = static_cast<QVariant>(*(QPolygon*)self);
	return new QVariant(__r);
}
void QPolygon_translate(void* self, int dx, int dy) {
	((QPolygon*)self)->translate(dx, dy);
}
void QPolygon_translate2(void* self, void* offset) {
	((QPolygon*)self)->translate(*(QPoint*)offset);
}
void* QPolygon_translated(void* self, int dx, int dy) {
	auto&& __r = ((QPolygon*)self)->translated(dx, dy);
	return new QPolygon(__r);
}
void* QPolygon_translated2(void* self, void* offset) {
	auto&& __r = ((QPolygon*)self)->translated(*(QPoint*)offset);
	return new QPolygon(__r);
}
void* QPolygon_boundingRect(void* self) {
	auto&& __r = ((QPolygon*)self)->boundingRect();
	return new QRect(__r);
}
void QPolygon_point(void* self, int i, int* x, int* y) {
	int x_local{};
	int y_local{};
	((QPolygon*)self)->point(i, &x_local, &y_local);
	*x = x_local;
	*y = y_local;
}
void* QPolygon_point2(void* self, int i) {
	auto&& __r = ((QPolygon*)self)->point(i);
	return new QPoint(__r);
}
void QPolygon_setPoint(void* self, int index, int x, int y) {
	((QPolygon*)self)->setPoint(index, x, y);
}
void QPolygon_setPoint2(void* self, int index, void* p) {
	((QPolygon*)self)->setPoint(index, *(QPoint*)p);
}
void QPolygon_setPoints(void* self, int nPoints, int firstx, int firsty) {
	((QPolygon*)self)->setPoints(nPoints, firstx, firsty);
}
void QPolygon_putPoints(void* self, int index, int nPoints, int firstx, int firsty) {
	((QPolygon*)self)->putPoints(index, nPoints, firstx, firsty);
}
void QPolygon_putPoints2(void* self, int index, int nPoints, void* from) {
	((QPolygon*)self)->putPoints(index, nPoints, *(QPolygon*)from);
}
int QPolygon_containsPoint(void* self, void* pt, int fillRule) {
	auto&& __r = ((QPolygon*)self)->containsPoint(*(QPoint*)pt, (Qt::FillRule)(fillRule));
	return __r;
}
void* QPolygon_united(void* self, void* r) {
	auto&& __r = ((QPolygon*)self)->united(*(QPolygon*)r);
	return new QPolygon(__r);
}
void* QPolygon_intersected(void* self, void* r) {
	auto&& __r = ((QPolygon*)self)->intersected(*(QPolygon*)r);
	return new QPolygon(__r);
}
void* QPolygon_subtracted(void* self, void* r) {
	auto&& __r = ((QPolygon*)self)->subtracted(*(QPolygon*)r);
	return new QPolygon(__r);
}
int QPolygon_intersects(void* self, void* r) {
	auto&& __r = ((QPolygon*)self)->intersects(*(QPolygon*)r);
	return __r;
}
void* QPolygon_toPolygonF(void* self) {
	auto&& __r = ((QPolygon*)self)->toPolygonF();
	return new QPolygonF(__r);
}
long long QPolygon_maxSize() {
	auto&& __r = QPolygon::maxSize();
	return __r;
}
long long QPolygon_size(void* self) {
	auto&& __r = ((QPolygon*)self)->size();
	return __r;
}
long long QPolygon_count(void* self) {
	auto&& __r = ((QPolygon*)self)->count();
	return __r;
}
long long QPolygon_length(void* self) {
	auto&& __r = ((QPolygon*)self)->length();
	return __r;
}
int QPolygon_isEmpty(void* self) {
	auto&& __r = ((QPolygon*)self)->isEmpty();
	return __r;
}
void QPolygon_resize(void* self, long long size) {
	((QPolygon*)self)->resize(size);
}
void QPolygon_resize2(void* self, long long size, void* c) {
	((QPolygon*)self)->resize(size, *(QPoint*)c);
}
void QPolygon_resizeForOverwrite(void* self, long long size) {
	((QPolygon*)self)->resizeForOverwrite(size);
}
long long QPolygon_capacity(void* self) {
	auto&& __r = ((QPolygon*)self)->capacity();
	return __r;
}
void QPolygon_reserve(void* self, long long size) {
	((QPolygon*)self)->reserve(size);
}
void QPolygon_squeeze(void* self) {
	((QPolygon*)self)->squeeze();
}
void QPolygon_detach(void* self) {
	((QPolygon*)self)->detach();
}
int QPolygon_isDetached(void* self) {
	auto&& __r = ((QPolygon*)self)->isDetached();
	return __r;
}
int QPolygon_isSharedWith(void* self, void** other, long long other_len) {
	QList<QPoint> other_list; other_list.reserve(other_len);
	for (long long __i = 0; __i < other_len; __i++) other_list.append(*(QPoint*)other[__i]);
	auto&& __r = ((QPolygon*)self)->isSharedWith(other_list);
	return __r;
}
void* QPolygon_data(void* self) {
	auto&& __r = ((QPolygon*)self)->data();
	return (void*)(__r);
}
void* QPolygon_data2(void* self) {
	auto&& __r = ((QPolygon*)self)->data();
	return (void*)(__r);
}
void* QPolygon_constData(void* self) {
	auto&& __r = ((QPolygon*)self)->constData();
	return (void*)(__r);
}
void QPolygon_clear(void* self) {
	((QPolygon*)self)->clear();
}
void* QPolygon_at(void* self, long long i) {
	auto&& __r = ((QPolygon*)self)->at(i);
	return (void*)&(__r);
}
void* QPolygon_OperatorSubscript(void* self, long long i) {
	auto&& __r = ((QPolygon*)self)->operator[](i);
	return (void*)&(__r);
}
void* QPolygon_OperatorSubscript2(void* self, long long i) {
	auto&& __r = ((QPolygon*)self)->operator[](i);
	return (void*)&(__r);
}
void QPolygon_append(void* self, void* t) {
	((QPolygon*)self)->append(*(QPoint*)t);
}
void QPolygon_append2(void* self, void** l, long long l_len) {
	QList<QPoint> l_list; l_list.reserve(l_len);
	for (long long __i = 0; __i < l_len; __i++) l_list.append(*(QPoint*)l[__i]);
	((QPolygon*)self)->append(l_list);
}
void QPolygon_prepend(void* self, void* t) {
	((QPolygon*)self)->prepend(*(QPoint*)t);
}
void QPolygon_replace(void* self, long long i, void* t) {
	((QPolygon*)self)->replace(i, *(QPoint*)t);
}
void QPolygon_remove(void* self, long long i) {
	((QPolygon*)self)->remove(i);
}
void QPolygon_removeFirst(void* self) {
	((QPolygon*)self)->removeFirst();
}
void QPolygon_removeLast(void* self) {
	((QPolygon*)self)->removeLast();
}
void* QPolygon_takeFirst(void* self) {
	auto&& __r = ((QPolygon*)self)->takeFirst();
	return new QPoint(__r);
}
void* QPolygon_takeLast(void* self) {
	auto&& __r = ((QPolygon*)self)->takeLast();
	return new QPoint(__r);
}
void QPolygon_removeAt(void* self, long long i) {
	((QPolygon*)self)->removeAt(i);
}
void* QPolygon_takeAt(void* self, long long i) {
	auto&& __r = ((QPolygon*)self)->takeAt(i);
	return new QPoint(__r);
}
void QPolygon_move(void* self, long long from, long long to) {
	((QPolygon*)self)->move(from, to);
}
void* QPolygon_first(void* self) {
	auto&& __r = ((QPolygon*)self)->first();
	return (void*)&(__r);
}
void* QPolygon_first2(void* self) {
	auto&& __r = ((QPolygon*)self)->first();
	return (void*)&(__r);
}
void* QPolygon_constFirst(void* self) {
	auto&& __r = ((QPolygon*)self)->constFirst();
	return (void*)&(__r);
}
void* QPolygon_last(void* self) {
	auto&& __r = ((QPolygon*)self)->last();
	return (void*)&(__r);
}
void* QPolygon_last2(void* self) {
	auto&& __r = ((QPolygon*)self)->last();
	return (void*)&(__r);
}
void* QPolygon_constLast(void* self) {
	auto&& __r = ((QPolygon*)self)->constLast();
	return (void*)&(__r);
}
int QPolygon_startsWith(void* self, void* t) {
	auto&& __r = ((QPolygon*)self)->startsWith(*(QPoint*)t);
	return __r;
}
int QPolygon_endsWith(void* self, void* t) {
	auto&& __r = ((QPolygon*)self)->endsWith(*(QPoint*)t);
	return __r;
}
void** QPolygon_mid(void* self, long long pos, long long* __out_len) {
	auto&& __r = ((QPolygon*)self)->mid(pos);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QPoint(__e);
	return __arr;
}
void** QPolygon_first3(void* self, long long n, long long* __out_len) {
	auto&& __r = ((QPolygon*)self)->first(n);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QPoint(__e);
	return __arr;
}
void** QPolygon_last3(void* self, long long n, long long* __out_len) {
	auto&& __r = ((QPolygon*)self)->last(n);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QPoint(__e);
	return __arr;
}
void** QPolygon_sliced(void* self, long long pos, long long* __out_len) {
	auto&& __r = ((QPolygon*)self)->sliced(pos);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QPoint(__e);
	return __arr;
}
void** QPolygon_sliced2(void* self, long long pos, long long n, long long* __out_len) {
	auto&& __r = ((QPolygon*)self)->sliced(pos, n);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QPoint(__e);
	return __arr;
}
void* QPolygon_value(void* self, long long i) {
	auto&& __r = ((QPolygon*)self)->value(i);
	return new QPoint(__r);
}
void* QPolygon_value2(void* self, long long i, void* defaultValue) {
	auto&& __r = ((QPolygon*)self)->value(i, *(QPoint*)defaultValue);
	return new QPoint(__r);
}
void QPolygon_swapItemsAt(void* self, long long i, long long j) {
	((QPolygon*)self)->swapItemsAt(i, j);
}
void QPolygon_push_back(void* self, void* t) {
	((QPolygon*)self)->push_back(*(QPoint*)t);
}
void QPolygon_push_front(void* self, void* t) {
	((QPolygon*)self)->push_front(*(QPoint*)t);
}
void QPolygon_pop_back(void* self) {
	((QPolygon*)self)->pop_back();
}
void QPolygon_pop_front(void* self) {
	((QPolygon*)self)->pop_front();
}
int QPolygon_empty(void* self) {
	auto&& __r = ((QPolygon*)self)->empty();
	return __r;
}
void* QPolygon_front(void* self) {
	auto&& __r = ((QPolygon*)self)->front();
	return (void*)&(__r);
}
void* QPolygon_front2(void* self) {
	auto&& __r = ((QPolygon*)self)->front();
	return (void*)&(__r);
}
void* QPolygon_back(void* self) {
	auto&& __r = ((QPolygon*)self)->back();
	return (void*)&(__r);
}
void* QPolygon_back2(void* self) {
	auto&& __r = ((QPolygon*)self)->back();
	return (void*)&(__r);
}
void QPolygon_shrink_to_fit(void* self) {
	((QPolygon*)self)->shrink_to_fit();
}
long long QPolygon_max_size(void* self) {
	auto&& __r = ((QPolygon*)self)->max_size();
	return __r;
}
void** QPolygon_OperatorPlus(void* self, void** l, long long l_len, long long* __out_len) {
	QList<QPoint> l_list; l_list.reserve(l_len);
	for (long long __i = 0; __i < l_len; __i++) l_list.append(*(QPoint*)l[__i]);
	auto&& __r = ((QPolygon*)self)->operator+(l_list);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QPoint(__e);
	return __arr;
}
void** QPolygon_OperatorPlus2(void* self, void** l, long long l_len, long long* __out_len) {
	QList<QPoint> l_list; l_list.reserve(l_len);
	for (long long __i = 0; __i < l_len; __i++) l_list.append(*(QPoint*)l[__i]);
	auto&& __r = ((QPolygon*)self)->operator+(l_list);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QPoint(__e);
	return __arr;
}
void** QPolygon_fromList(void** list, long long list_len, long long* __out_len) {
	QList<QPoint> list_list; list_list.reserve(list_len);
	for (long long __i = 0; __i < list_len; __i++) list_list.append(*(QPoint*)list[__i]);
	auto&& __r = QPolygon::fromList(list_list);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QPoint(__e);
	return __arr;
}
void** QPolygon_toList(void* self, long long* __out_len) {
	auto&& __r = ((QPolygon*)self)->toList();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QPoint(__e);
	return __arr;
}
void** QPolygon_fromVector(void** vector, long long vector_len, long long* __out_len) {
	QList<QPoint> vector_list; vector_list.reserve(vector_len);
	for (long long __i = 0; __i < vector_len; __i++) vector_list.append(*(QPoint*)vector[__i]);
	auto&& __r = QPolygon::fromVector(vector_list);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QPoint(__e);
	return __arr;
}
void** QPolygon_toVector(void* self, long long* __out_len) {
	auto&& __r = ((QPolygon*)self)->toVector();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QPoint(__e);
	return __arr;
}
void QPolygon_putPoints3(void* self, int index, int nPoints, void* from, int fromIndex) {
	((QPolygon*)self)->putPoints(index, nPoints, *(QPolygon*)from, fromIndex);
}
void QPolygon_remove2(void* self, long long i, long long n) {
	((QPolygon*)self)->remove(i, n);
}
void** QPolygon_mid2(void* self, long long pos, long long len_, long long* __out_len) {
	auto&& __r = ((QPolygon*)self)->mid(pos, len_);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QPoint(__e);
	return __arr;
}
void* QPolygon_OperatorMultiply(void* self, void* m) {
	auto&& __r = (*(QPolygon*)self) * (*(QTransform*)m);
	return new QPolygon(__r);
}
void* QPolygon_toQVariant(void* self) {
	auto&& __r = QVariant::fromValue(*(QPolygon*)self);
	return new QVariant(__r);
}
void QPolygon_delete(void* self) { delete (QPolygon*)self; }

}

#include <QtGui/qevent.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qpoint.h>
#include <QtCore/qnamespace.h>
#include <QtGui/qpointingdevice.h>
#include <QtCore/qobject.h>
#include <QtGui/qeventpoint.h>
#include <QtGui/qinputdevice.h>
#include "gen_qhoverevent.h"

extern "C" {

void* QHoverEvent_base(void* self) { return (void*)static_cast<QSinglePointEvent*>((QHoverEvent*)self); }
void QHoverEvent_chain(void* self, void** out) {
	QHoverEvent* __p = (QHoverEvent*)self;
	QSinglePointEvent* __b0 = static_cast<QSinglePointEvent*>(__p); out[0] = (void*)__b0;
	QPointerEvent* __b1 = static_cast<QPointerEvent*>(__b0); out[1] = (void*)__b1;
	QInputEvent* __b2 = static_cast<QInputEvent*>(__b1); out[2] = (void*)__b2;
	QEvent* __b3 = static_cast<QEvent*>(__b2); out[3] = (void*)__b3;
}
void* QHoverEvent_new(int type_, void* scenePos, void* globalPos, void* oldPos) {
	return new QHoverEvent((QEvent::Type)(type_), *(QPointF*)scenePos, *(QPointF*)globalPos, *(QPointF*)oldPos);
}
void* QHoverEvent_new2(int type_, void* pos, void* oldPos) {
	return new QHoverEvent((QEvent::Type)(type_), *(QPointF*)pos, *(QPointF*)oldPos);
}
void* QHoverEvent_new3(int type_, void* scenePos, void* globalPos, void* oldPos, int modifiers) {
	return new QHoverEvent((QEvent::Type)(type_), *(QPointF*)scenePos, *(QPointF*)globalPos, *(QPointF*)oldPos, QFlags<Qt::KeyboardModifier>(modifiers));
}
void* QHoverEvent_new4(int type_, void* scenePos, void* globalPos, void* oldPos, int modifiers, void* device) {
	return new QHoverEvent((QEvent::Type)(type_), *(QPointF*)scenePos, *(QPointF*)globalPos, *(QPointF*)oldPos, QFlags<Qt::KeyboardModifier>(modifiers), (QPointingDevice*)device);
}
void* QHoverEvent_new5(int type_, void* pos, void* oldPos, int modifiers) {
	return new QHoverEvent((QEvent::Type)(type_), *(QPointF*)pos, *(QPointF*)oldPos, QFlags<Qt::KeyboardModifier>(modifiers));
}
void* QHoverEvent_new6(int type_, void* pos, void* oldPos, int modifiers, void* device) {
	return new QHoverEvent((QEvent::Type)(type_), *(QPointF*)pos, *(QPointF*)oldPos, QFlags<Qt::KeyboardModifier>(modifiers), (QPointingDevice*)device);
}
void* QHoverEvent_pos(void* self) {
	auto&& __r = ((QHoverEvent*)self)->pos();
	return new QPoint(__r);
}
void* QHoverEvent_posF(void* self) {
	auto&& __r = ((QHoverEvent*)self)->posF();
	return new QPointF(__r);
}
void* QHoverEvent_oldPos(void* self) {
	auto&& __r = ((QHoverEvent*)self)->oldPos();
	return new QPoint(__r);
}
void* QHoverEvent_oldPosF(void* self) {
	auto&& __r = ((QHoverEvent*)self)->oldPosF();
	return new QPointF(__r);
}
int QHoverEvent_registerEventType() {
	auto&& __r = QHoverEvent::registerEventType();
	return __r;
}
int QHoverEvent_registerEventType2(int hint) {
	auto&& __r = QHoverEvent::registerEventType(hint);
	return __r;
}
void QHoverEvent_delete(void* self) { delete (QHoverEvent*)self; }

}

#include <QtGui/qevent.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qstring.h>
#include "gen_qstatustipevent.h"

extern "C" {

void* QStatusTipEvent_base(void* self) { return (void*)static_cast<QEvent*>((QStatusTipEvent*)self); }
void* QStatusTipEvent_new(const char* tip, long long tip_len) {
	QString tip_qstr = QString::fromUtf8(tip, tip_len);
	return new QStatusTipEvent(tip_qstr);
}
char* QStatusTipEvent_tip(void* self, long long* __out_len) {
	auto&& __r = ((QStatusTipEvent*)self)->tip();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QStatusTipEvent_registerEventType() {
	auto&& __r = QStatusTipEvent::registerEventType();
	return __r;
}
int QStatusTipEvent_registerEventType2(int hint) {
	auto&& __r = QStatusTipEvent::registerEventType(hint);
	return __r;
}
void QStatusTipEvent_delete(void* self) { delete (QStatusTipEvent*)self; }

}

#include <QtGui/qbrush.h>
#include <QtCore/qpoint.h>
#include <QtGui/qcolor.h>
#include "gen_qlineargradient.h"

extern "C" {

void* QLinearGradient_base(void* self) { return (void*)static_cast<QGradient*>((QLinearGradient*)self); }
void* QLinearGradient_new() {
	return new QLinearGradient();
}
void* QLinearGradient_new2(void* start, void* finalStop) {
	return new QLinearGradient(*(QPointF*)start, *(QPointF*)finalStop);
}
void* QLinearGradient_new3(double xStart, double yStart, double xFinalStop, double yFinalStop) {
	return new QLinearGradient(xStart, yStart, xFinalStop, yFinalStop);
}
void* QLinearGradient_new4(void* arg1) {
	return new QLinearGradient(*(QLinearGradient*)arg1);
}
void* QLinearGradient_start(void* self) {
	auto&& __r = ((QLinearGradient*)self)->start();
	return new QPointF(__r);
}
void QLinearGradient_setStart(void* self, void* start) {
	((QLinearGradient*)self)->setStart(*(QPointF*)start);
}
void QLinearGradient_setStart2(void* self, double x, double y) {
	((QLinearGradient*)self)->setStart(x, y);
}
void* QLinearGradient_finalStop(void* self) {
	auto&& __r = ((QLinearGradient*)self)->finalStop();
	return new QPointF(__r);
}
void QLinearGradient_setFinalStop(void* self, void* stop) {
	((QLinearGradient*)self)->setFinalStop(*(QPointF*)stop);
}
void QLinearGradient_setFinalStop2(void* self, double x, double y) {
	((QLinearGradient*)self)->setFinalStop(x, y);
}
void QLinearGradient_delete(void* self) { delete (QLinearGradient*)self; }

}

#include <QtGui/qabstracttextdocumentlayout.h>
#include <QtCore/qsize.h>
#include <QtGui/qtextdocument.h>
#include <QtGui/qtextformat.h>
#include <QtGui/qpainter.h>
#include <QtCore/qrect.h>
#include "gen_qtextobjectinterface.h"
#include "gen_purego_cb.h"

class QTextObjectInterface__Go : public QTextObjectInterface {
public:
	long long __goId;
	unsigned char __goOn[2] = {};
	template<typename... A> QTextObjectInterface__Go(long long id, A&&... a) : QTextObjectInterface(std::forward<A>(a)...), __goId(id) {}
	QSizeF intrinsicSize(QTextDocument* doc, int posInDocument, const QTextFormat& format) override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp_voidp_int_voidp) { void* __r = ((void*(*)(long long, long long, void*, int, void*, int*))cppgo_virt_voidp_voidp_int_voidp)(__goId, 5173984003908732465LL, (void*)doc, posInDocument, (void*)&format, &__h); if (__h) return *(QSizeF*)__r; }
		return QSizeF{};
	}
	void drawObject(QPainter* painter, const QRectF& rect, QTextDocument* doc, int posInDocument, const QTextFormat& format) override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_void_voidp_voidp_voidp_int_voidp) { ((void(*)(long long, long long, void*, void*, void*, int, void*, int*))cppgo_virt_void_voidp_voidp_voidp_int_voidp)(__goId, 8290039120595127048LL, (void*)painter, (void*)&rect, (void*)doc, posInDocument, (void*)&format, &__h); if (__h) return; }
		return;
	}
};

extern "C" {

void QTextObjectInterface__set_override(void* self, int idx) { ((QTextObjectInterface__Go*)self)->__goOn[idx] = 1; }
void* QTextObjectInterface_intrinsicSize(void* self, void* doc, int posInDocument, void* format) {
	auto&& __r = ((QTextObjectInterface*)self)->intrinsicSize((QTextDocument*)doc, posInDocument, *(QTextFormat*)format);
	return new QSizeF(__r);
}
void QTextObjectInterface_drawObject(void* self, void* painter, void* rect, void* doc, int posInDocument, void* format) {
	((QTextObjectInterface*)self)->drawObject((QPainter*)painter, *(QRectF*)rect, (QTextDocument*)doc, posInDocument, *(QTextFormat*)format);
}
void* QTextObjectInterface_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QTextObjectInterface*)self)->operator=(*(QTextObjectInterface*)arg1);
	return (void*)&(__r);
}
void* QTextObjectInterface__Go_new(long long id) { return new QTextObjectInterface__Go(id); }
void QTextObjectInterface_delete(void* self) { delete (QTextObjectInterface*)self; }

}

#include <QtGui/qaccessible.h>
#include <QtCore/qobject.h>
#include <QtGui/qaccessible_base.h>
#include "gen_qaccessibletablemodelchangeevent.h"

extern "C" {

void* QAccessibleTableModelChangeEvent_base(void* self) { return (void*)static_cast<QAccessibleEvent*>((QAccessibleTableModelChangeEvent*)self); }
void* QAccessibleTableModelChangeEvent_new(void* obj, int changeType) {
	return new QAccessibleTableModelChangeEvent((QObject*)obj, (QAccessibleTableModelChangeEvent::ModelChangeType)(changeType));
}
void* QAccessibleTableModelChangeEvent_new2(void* iface, int changeType) {
	return new QAccessibleTableModelChangeEvent((QAccessibleInterface*)iface, (QAccessibleTableModelChangeEvent::ModelChangeType)(changeType));
}
void QAccessibleTableModelChangeEvent_setModelChangeType(void* self, int changeType) {
	((QAccessibleTableModelChangeEvent*)self)->setModelChangeType((QAccessibleTableModelChangeEvent::ModelChangeType)(changeType));
}
int QAccessibleTableModelChangeEvent_modelChangeType(void* self) {
	auto&& __r = ((QAccessibleTableModelChangeEvent*)self)->modelChangeType();
	return (int)(__r);
}
void QAccessibleTableModelChangeEvent_setFirstRow(void* self, int row) {
	((QAccessibleTableModelChangeEvent*)self)->setFirstRow(row);
}
void QAccessibleTableModelChangeEvent_setFirstColumn(void* self, int col) {
	((QAccessibleTableModelChangeEvent*)self)->setFirstColumn(col);
}
void QAccessibleTableModelChangeEvent_setLastRow(void* self, int row) {
	((QAccessibleTableModelChangeEvent*)self)->setLastRow(row);
}
void QAccessibleTableModelChangeEvent_setLastColumn(void* self, int col) {
	((QAccessibleTableModelChangeEvent*)self)->setLastColumn(col);
}
int QAccessibleTableModelChangeEvent_firstRow(void* self) {
	auto&& __r = ((QAccessibleTableModelChangeEvent*)self)->firstRow();
	return __r;
}
int QAccessibleTableModelChangeEvent_firstColumn(void* self) {
	auto&& __r = ((QAccessibleTableModelChangeEvent*)self)->firstColumn();
	return __r;
}
int QAccessibleTableModelChangeEvent_lastRow(void* self) {
	auto&& __r = ((QAccessibleTableModelChangeEvent*)self)->lastRow();
	return __r;
}
int QAccessibleTableModelChangeEvent_lastColumn(void* self) {
	auto&& __r = ((QAccessibleTableModelChangeEvent*)self)->lastColumn();
	return __r;
}
void QAccessibleTableModelChangeEvent_delete(void* self) { delete (QAccessibleTableModelChangeEvent*)self; }

}

#include <QtGui/qfontmetrics.h>
#include <QtGui/qfont.h>
#include <QtGui/qpaintdevice.h>
#include <QtCore/qchar.h>
#include <QtCore/qstring.h>
#include <QtGui/qtextoption.h>
#include <QtCore/qrect.h>
#include <QtCore/qsize.h>
#include <QtCore/qnamespace.h>
#include "gen_qfontmetrics.h"

extern "C" {

void* QFontMetrics_new(void* arg1) {
	return new QFontMetrics(*(QFont*)arg1);
}
void* QFontMetrics_new2(void* font, void* pd) {
	return new QFontMetrics(*(QFont*)font, (QPaintDevice*)pd);
}
void* QFontMetrics_new3(void* arg1) {
	return new QFontMetrics(*(QFontMetrics*)arg1);
}
void* QFontMetrics_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QFontMetrics*)self)->operator=(*(QFontMetrics*)arg1);
	return (void*)&(__r);
}
void QFontMetrics_swap(void* self, void* other) {
	((QFontMetrics*)self)->swap(*(QFontMetrics*)other);
}
int QFontMetrics_ascent(void* self) {
	auto&& __r = ((QFontMetrics*)self)->ascent();
	return __r;
}
int QFontMetrics_capHeight(void* self) {
	auto&& __r = ((QFontMetrics*)self)->capHeight();
	return __r;
}
int QFontMetrics_descent(void* self) {
	auto&& __r = ((QFontMetrics*)self)->descent();
	return __r;
}
int QFontMetrics_height(void* self) {
	auto&& __r = ((QFontMetrics*)self)->height();
	return __r;
}
int QFontMetrics_leading(void* self) {
	auto&& __r = ((QFontMetrics*)self)->leading();
	return __r;
}
int QFontMetrics_lineSpacing(void* self) {
	auto&& __r = ((QFontMetrics*)self)->lineSpacing();
	return __r;
}
int QFontMetrics_minLeftBearing(void* self) {
	auto&& __r = ((QFontMetrics*)self)->minLeftBearing();
	return __r;
}
int QFontMetrics_minRightBearing(void* self) {
	auto&& __r = ((QFontMetrics*)self)->minRightBearing();
	return __r;
}
int QFontMetrics_maxWidth(void* self) {
	auto&& __r = ((QFontMetrics*)self)->maxWidth();
	return __r;
}
int QFontMetrics_xHeight(void* self) {
	auto&& __r = ((QFontMetrics*)self)->xHeight();
	return __r;
}
int QFontMetrics_averageCharWidth(void* self) {
	auto&& __r = ((QFontMetrics*)self)->averageCharWidth();
	return __r;
}
int QFontMetrics_inFont(void* self, void* arg1) {
	auto&& __r = ((QFontMetrics*)self)->inFont(*(QChar*)arg1);
	return __r;
}
int QFontMetrics_inFontUcs4(void* self, unsigned int ucs4) {
	auto&& __r = ((QFontMetrics*)self)->inFontUcs4(ucs4);
	return __r;
}
int QFontMetrics_leftBearing(void* self, void* arg1) {
	auto&& __r = ((QFontMetrics*)self)->leftBearing(*(QChar*)arg1);
	return __r;
}
int QFontMetrics_rightBearing(void* self, void* arg1) {
	auto&& __r = ((QFontMetrics*)self)->rightBearing(*(QChar*)arg1);
	return __r;
}
int QFontMetrics_horizontalAdvance(void* self, const char* arg1, long long arg1_len) {
	QString arg1_qstr = QString::fromUtf8(arg1, arg1_len);
	auto&& __r = ((QFontMetrics*)self)->horizontalAdvance(arg1_qstr);
	return __r;
}
int QFontMetrics_horizontalAdvance2(void* self, const char* arg1, long long arg1_len, void* textOption) {
	QString arg1_qstr = QString::fromUtf8(arg1, arg1_len);
	auto&& __r = ((QFontMetrics*)self)->horizontalAdvance(arg1_qstr, *(QTextOption*)textOption);
	return __r;
}
int QFontMetrics_horizontalAdvance3(void* self, void* arg1) {
	auto&& __r = ((QFontMetrics*)self)->horizontalAdvance(*(QChar*)arg1);
	return __r;
}
void* QFontMetrics_boundingRect(void* self, void* arg1) {
	auto&& __r = ((QFontMetrics*)self)->boundingRect(*(QChar*)arg1);
	return new QRect(__r);
}
void* QFontMetrics_boundingRect2(void* self, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = ((QFontMetrics*)self)->boundingRect(text_qstr);
	return new QRect(__r);
}
void* QFontMetrics_boundingRect3(void* self, const char* text, long long text_len, void* textOption) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = ((QFontMetrics*)self)->boundingRect(text_qstr, *(QTextOption*)textOption);
	return new QRect(__r);
}
void* QFontMetrics_boundingRect4(void* self, void* r, int flags, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = ((QFontMetrics*)self)->boundingRect(*(QRect*)r, flags, text_qstr);
	return new QRect(__r);
}
void* QFontMetrics_boundingRect5(void* self, int x, int y, int w, int h, int flags, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = ((QFontMetrics*)self)->boundingRect(x, y, w, h, flags, text_qstr);
	return new QRect(__r);
}
void* QFontMetrics_size(void* self, int flags, const char* str, long long str_len) {
	QString str_qstr = QString::fromUtf8(str, str_len);
	auto&& __r = ((QFontMetrics*)self)->size(flags, str_qstr);
	return new QSize(__r);
}
void* QFontMetrics_tightBoundingRect(void* self, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = ((QFontMetrics*)self)->tightBoundingRect(text_qstr);
	return new QRect(__r);
}
void* QFontMetrics_tightBoundingRect2(void* self, const char* text, long long text_len, void* textOption) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = ((QFontMetrics*)self)->tightBoundingRect(text_qstr, *(QTextOption*)textOption);
	return new QRect(__r);
}
char* QFontMetrics_elidedText(void* self, const char* text, long long text_len, int mode, int width, long long* __out_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = ((QFontMetrics*)self)->elidedText(text_qstr, (Qt::TextElideMode)(mode), width);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QFontMetrics_underlinePos(void* self) {
	auto&& __r = ((QFontMetrics*)self)->underlinePos();
	return __r;
}
int QFontMetrics_overlinePos(void* self) {
	auto&& __r = ((QFontMetrics*)self)->overlinePos();
	return __r;
}
int QFontMetrics_strikeOutPos(void* self) {
	auto&& __r = ((QFontMetrics*)self)->strikeOutPos();
	return __r;
}
int QFontMetrics_lineWidth(void* self) {
	auto&& __r = ((QFontMetrics*)self)->lineWidth();
	return __r;
}
double QFontMetrics_fontDpi(void* self) {
	auto&& __r = ((QFontMetrics*)self)->fontDpi();
	return __r;
}
int QFontMetrics_OperatorEqual(void* self, void* other) {
	auto&& __r = ((QFontMetrics*)self)->operator==(*(QFontMetrics*)other);
	return __r;
}
int QFontMetrics_OperatorNotEqual(void* self, void* other) {
	auto&& __r = ((QFontMetrics*)self)->operator!=(*(QFontMetrics*)other);
	return __r;
}
int QFontMetrics_horizontalAdvance4(void* self, const char* arg1, long long arg1_len, int len_) {
	QString arg1_qstr = QString::fromUtf8(arg1, arg1_len);
	auto&& __r = ((QFontMetrics*)self)->horizontalAdvance(arg1_qstr, len_);
	return __r;
}
void* QFontMetrics_boundingRect6(void* self, void* r, int flags, const char* text, long long text_len, int tabstops) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = ((QFontMetrics*)self)->boundingRect(*(QRect*)r, flags, text_qstr, tabstops);
	return new QRect(__r);
}
void* QFontMetrics_boundingRect7(void* self, void* r, int flags, const char* text, long long text_len, int tabstops, int* tabarray) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	int tabarray_local{};
	auto&& __r = ((QFontMetrics*)self)->boundingRect(*(QRect*)r, flags, text_qstr, tabstops, &tabarray_local);
	*tabarray = tabarray_local;
	return new QRect(__r);
}
void* QFontMetrics_boundingRect8(void* self, int x, int y, int w, int h, int flags, const char* text, long long text_len, int tabstops) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = ((QFontMetrics*)self)->boundingRect(x, y, w, h, flags, text_qstr, tabstops);
	return new QRect(__r);
}
void* QFontMetrics_boundingRect9(void* self, int x, int y, int w, int h, int flags, const char* text, long long text_len, int tabstops, int* tabarray) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	int tabarray_local{};
	auto&& __r = ((QFontMetrics*)self)->boundingRect(x, y, w, h, flags, text_qstr, tabstops, &tabarray_local);
	*tabarray = tabarray_local;
	return new QRect(__r);
}
void* QFontMetrics_size2(void* self, int flags, const char* str, long long str_len, int tabstops) {
	QString str_qstr = QString::fromUtf8(str, str_len);
	auto&& __r = ((QFontMetrics*)self)->size(flags, str_qstr, tabstops);
	return new QSize(__r);
}
void* QFontMetrics_size3(void* self, int flags, const char* str, long long str_len, int tabstops, int* tabarray) {
	QString str_qstr = QString::fromUtf8(str, str_len);
	int tabarray_local{};
	auto&& __r = ((QFontMetrics*)self)->size(flags, str_qstr, tabstops, &tabarray_local);
	*tabarray = tabarray_local;
	return new QSize(__r);
}
char* QFontMetrics_elidedText2(void* self, const char* text, long long text_len, int mode, int width, int flags, long long* __out_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = ((QFontMetrics*)self)->elidedText(text_qstr, (Qt::TextElideMode)(mode), width, flags);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QFontMetrics_delete(void* self) { delete (QFontMetrics*)self; }

}

#include <QtGui/qpainter.h>
#include <QtGui/qpaintdevice.h>
#include <QtGui/qfont.h>
#include <QtGui/qfontmetrics.h>
#include <QtGui/qfontinfo.h>
#include <QtGui/qcolor.h>
#include <QtGui/qpen.h>
#include <QtCore/qnamespace.h>
#include <QtGui/qbrush.h>
#include <QtCore/qpoint.h>
#include <QtGui/qregion.h>
#include <QtGui/qpainterpath.h>
#include <QtCore/qrect.h>
#include <QtGui/qtransform.h>
#include <QtGui/qpolygon.h>
#include <QtCore/qline.h>
#include <QtGui/qpixmap.h>
#include <QtGui/qpicture.h>
#include <QtGui/qimage.h>
#include <QtGui/qglyphrun.h>
#include <QtGui/qstatictext.h>
#include <QtCore/qstring.h>
#include <QtGui/qtextoption.h>
#include <QtGui/qpaintengine.h>
#include "gen_qpainter.h"

extern "C" {

void* QPainter_new() {
	return new QPainter();
}
void* QPainter_new2(void* arg1) {
	return new QPainter((QPaintDevice*)arg1);
}
void* QPainter_device(void* self) {
	auto&& __r = ((QPainter*)self)->device();
	return (void*)(__r);
}
int QPainter_begin(void* self, void* arg1) {
	auto&& __r = ((QPainter*)self)->begin((QPaintDevice*)arg1);
	return __r;
}
int QPainter_end(void* self) {
	auto&& __r = ((QPainter*)self)->end();
	return __r;
}
int QPainter_isActive(void* self) {
	auto&& __r = ((QPainter*)self)->isActive();
	return __r;
}
void QPainter_setCompositionMode(void* self, int mode) {
	((QPainter*)self)->setCompositionMode((QPainter::CompositionMode)(mode));
}
int QPainter_compositionMode(void* self) {
	auto&& __r = ((QPainter*)self)->compositionMode();
	return (int)(__r);
}
void* QPainter_font(void* self) {
	auto&& __r = ((QPainter*)self)->font();
	return (void*)&(__r);
}
void QPainter_setFont(void* self, void* f) {
	((QPainter*)self)->setFont(*(QFont*)f);
}
void* QPainter_fontMetrics(void* self) {
	auto&& __r = ((QPainter*)self)->fontMetrics();
	return new QFontMetrics(__r);
}
void* QPainter_fontInfo(void* self) {
	auto&& __r = ((QPainter*)self)->fontInfo();
	return new QFontInfo(__r);
}
void QPainter_setPen(void* self, void* color) {
	((QPainter*)self)->setPen(*(QColor*)color);
}
void QPainter_setPen2(void* self, void* pen) {
	((QPainter*)self)->setPen(*(QPen*)pen);
}
void QPainter_setPen3(void* self, int style) {
	((QPainter*)self)->setPen((Qt::PenStyle)(style));
}
void* QPainter_pen(void* self) {
	auto&& __r = ((QPainter*)self)->pen();
	return (void*)&(__r);
}
void QPainter_setBrush(void* self, void* brush) {
	((QPainter*)self)->setBrush(*(QBrush*)brush);
}
void QPainter_setBrush2(void* self, int style) {
	((QPainter*)self)->setBrush((Qt::BrushStyle)(style));
}
void QPainter_setBrush3(void* self, void* color) {
	((QPainter*)self)->setBrush(*(QColor*)color);
}
void QPainter_setBrush4(void* self, int color) {
	((QPainter*)self)->setBrush((Qt::GlobalColor)(color));
}
void* QPainter_brush(void* self) {
	auto&& __r = ((QPainter*)self)->brush();
	return (void*)&(__r);
}
void QPainter_setBackgroundMode(void* self, int mode) {
	((QPainter*)self)->setBackgroundMode((Qt::BGMode)(mode));
}
int QPainter_backgroundMode(void* self) {
	auto&& __r = ((QPainter*)self)->backgroundMode();
	return (int)(__r);
}
void* QPainter_brushOrigin(void* self) {
	auto&& __r = ((QPainter*)self)->brushOrigin();
	return new QPoint(__r);
}
void* QPainter_brushOriginF(void* self) {
	auto&& __r = ((QPainter*)self)->brushOriginF();
	return new QPointF(__r);
}
void QPainter_setBrushOrigin(void* self, int x, int y) {
	((QPainter*)self)->setBrushOrigin(x, y);
}
void QPainter_setBrushOrigin2(void* self, void* arg1) {
	((QPainter*)self)->setBrushOrigin(*(QPoint*)arg1);
}
void QPainter_setBrushOrigin3(void* self, void* arg1) {
	((QPainter*)self)->setBrushOrigin(*(QPointF*)arg1);
}
void QPainter_setBackground(void* self, void* bg) {
	((QPainter*)self)->setBackground(*(QBrush*)bg);
}
void* QPainter_background(void* self) {
	auto&& __r = ((QPainter*)self)->background();
	return (void*)&(__r);
}
double QPainter_opacity(void* self) {
	auto&& __r = ((QPainter*)self)->opacity();
	return __r;
}
void QPainter_setOpacity(void* self, double opacity) {
	((QPainter*)self)->setOpacity(opacity);
}
void* QPainter_clipRegion(void* self) {
	auto&& __r = ((QPainter*)self)->clipRegion();
	return new QRegion(__r);
}
void* QPainter_clipPath(void* self) {
	auto&& __r = ((QPainter*)self)->clipPath();
	return new QPainterPath(__r);
}
void QPainter_setClipRect(void* self, void* arg1) {
	((QPainter*)self)->setClipRect(*(QRectF*)arg1);
}
void QPainter_setClipRect2(void* self, void* arg1) {
	((QPainter*)self)->setClipRect(*(QRect*)arg1);
}
void QPainter_setClipRect3(void* self, int x, int y, int w, int h) {
	((QPainter*)self)->setClipRect(x, y, w, h);
}
void QPainter_setClipRegion(void* self, void* arg1) {
	((QPainter*)self)->setClipRegion(*(QRegion*)arg1);
}
void QPainter_setClipPath(void* self, void* path) {
	((QPainter*)self)->setClipPath(*(QPainterPath*)path);
}
void QPainter_setClipping(void* self, int enable) {
	((QPainter*)self)->setClipping(enable);
}
int QPainter_hasClipping(void* self) {
	auto&& __r = ((QPainter*)self)->hasClipping();
	return __r;
}
void* QPainter_clipBoundingRect(void* self) {
	auto&& __r = ((QPainter*)self)->clipBoundingRect();
	return new QRectF(__r);
}
void QPainter_save(void* self) {
	((QPainter*)self)->save();
}
void QPainter_restore(void* self) {
	((QPainter*)self)->restore();
}
void QPainter_setTransform(void* self, void* transform) {
	((QPainter*)self)->setTransform(*(QTransform*)transform);
}
void* QPainter_transform(void* self) {
	auto&& __r = ((QPainter*)self)->transform();
	return (void*)&(__r);
}
void* QPainter_deviceTransform(void* self) {
	auto&& __r = ((QPainter*)self)->deviceTransform();
	return (void*)&(__r);
}
void QPainter_resetTransform(void* self) {
	((QPainter*)self)->resetTransform();
}
void QPainter_setWorldTransform(void* self, void* matrix) {
	((QPainter*)self)->setWorldTransform(*(QTransform*)matrix);
}
void* QPainter_worldTransform(void* self) {
	auto&& __r = ((QPainter*)self)->worldTransform();
	return (void*)&(__r);
}
void* QPainter_combinedTransform(void* self) {
	auto&& __r = ((QPainter*)self)->combinedTransform();
	return new QTransform(__r);
}
void QPainter_setWorldMatrixEnabled(void* self, int enabled) {
	((QPainter*)self)->setWorldMatrixEnabled(enabled);
}
int QPainter_worldMatrixEnabled(void* self) {
	auto&& __r = ((QPainter*)self)->worldMatrixEnabled();
	return __r;
}
void QPainter_scale(void* self, double sx, double sy) {
	((QPainter*)self)->scale(sx, sy);
}
void QPainter_shear(void* self, double sh, double sv) {
	((QPainter*)self)->shear(sh, sv);
}
void QPainter_rotate(void* self, double a) {
	((QPainter*)self)->rotate(a);
}
void QPainter_translate(void* self, void* offset) {
	((QPainter*)self)->translate(*(QPointF*)offset);
}
void QPainter_translate2(void* self, void* offset) {
	((QPainter*)self)->translate(*(QPoint*)offset);
}
void QPainter_translate3(void* self, double dx, double dy) {
	((QPainter*)self)->translate(dx, dy);
}
void* QPainter_window(void* self) {
	auto&& __r = ((QPainter*)self)->window();
	return new QRect(__r);
}
void QPainter_setWindow(void* self, void* window) {
	((QPainter*)self)->setWindow(*(QRect*)window);
}
void QPainter_setWindow2(void* self, int x, int y, int w, int h) {
	((QPainter*)self)->setWindow(x, y, w, h);
}
void* QPainter_viewport(void* self) {
	auto&& __r = ((QPainter*)self)->viewport();
	return new QRect(__r);
}
void QPainter_setViewport(void* self, void* viewport) {
	((QPainter*)self)->setViewport(*(QRect*)viewport);
}
void QPainter_setViewport2(void* self, int x, int y, int w, int h) {
	((QPainter*)self)->setViewport(x, y, w, h);
}
void QPainter_setViewTransformEnabled(void* self, int enable) {
	((QPainter*)self)->setViewTransformEnabled(enable);
}
int QPainter_viewTransformEnabled(void* self) {
	auto&& __r = ((QPainter*)self)->viewTransformEnabled();
	return __r;
}
void QPainter_strokePath(void* self, void* path, void* pen) {
	((QPainter*)self)->strokePath(*(QPainterPath*)path, *(QPen*)pen);
}
void QPainter_fillPath(void* self, void* path, void* brush) {
	((QPainter*)self)->fillPath(*(QPainterPath*)path, *(QBrush*)brush);
}
void QPainter_drawPath(void* self, void* path) {
	((QPainter*)self)->drawPath(*(QPainterPath*)path);
}
void QPainter_drawPoint(void* self, void* pt) {
	((QPainter*)self)->drawPoint(*(QPointF*)pt);
}
void QPainter_drawPoint2(void* self, void* p) {
	((QPainter*)self)->drawPoint(*(QPoint*)p);
}
void QPainter_drawPoint3(void* self, int x, int y) {
	((QPainter*)self)->drawPoint(x, y);
}
void QPainter_drawPoints(void* self, void* points, int pointCount) {
	((QPainter*)self)->drawPoints((QPointF*)points, pointCount);
}
void QPainter_drawPoints2(void* self, void* points) {
	((QPainter*)self)->drawPoints(*(QPolygonF*)points);
}
void QPainter_drawPoints3(void* self, void* points, int pointCount) {
	((QPainter*)self)->drawPoints((QPoint*)points, pointCount);
}
void QPainter_drawPoints4(void* self, void* points) {
	((QPainter*)self)->drawPoints(*(QPolygon*)points);
}
void QPainter_drawLine(void* self, void* line) {
	((QPainter*)self)->drawLine(*(QLineF*)line);
}
void QPainter_drawLine2(void* self, void* line) {
	((QPainter*)self)->drawLine(*(QLine*)line);
}
void QPainter_drawLine3(void* self, int x1, int y1, int x2, int y2) {
	((QPainter*)self)->drawLine(x1, y1, x2, y2);
}
void QPainter_drawLine4(void* self, void* p1, void* p2) {
	((QPainter*)self)->drawLine(*(QPoint*)p1, *(QPoint*)p2);
}
void QPainter_drawLine5(void* self, void* p1, void* p2) {
	((QPainter*)self)->drawLine(*(QPointF*)p1, *(QPointF*)p2);
}
void QPainter_drawLines(void* self, void* lines, int lineCount) {
	((QPainter*)self)->drawLines((QLineF*)lines, lineCount);
}
void QPainter_drawLines2(void* self, void** lines, long long lines_len) {
	QList<QLineF> lines_list; lines_list.reserve(lines_len);
	for (long long __i = 0; __i < lines_len; __i++) lines_list.append(*(QLineF*)lines[__i]);
	((QPainter*)self)->drawLines(lines_list);
}
void QPainter_drawLines3(void* self, void* pointPairs, int lineCount) {
	((QPainter*)self)->drawLines((QPointF*)pointPairs, lineCount);
}
void QPainter_drawLines4(void* self, void** pointPairs, long long pointPairs_len) {
	QList<QPointF> pointPairs_list; pointPairs_list.reserve(pointPairs_len);
	for (long long __i = 0; __i < pointPairs_len; __i++) pointPairs_list.append(*(QPointF*)pointPairs[__i]);
	((QPainter*)self)->drawLines(pointPairs_list);
}
void QPainter_drawLines5(void* self, void* lines, int lineCount) {
	((QPainter*)self)->drawLines((QLine*)lines, lineCount);
}
void QPainter_drawLines6(void* self, void** lines, long long lines_len) {
	QList<QLine> lines_list; lines_list.reserve(lines_len);
	for (long long __i = 0; __i < lines_len; __i++) lines_list.append(*(QLine*)lines[__i]);
	((QPainter*)self)->drawLines(lines_list);
}
void QPainter_drawLines7(void* self, void* pointPairs, int lineCount) {
	((QPainter*)self)->drawLines((QPoint*)pointPairs, lineCount);
}
void QPainter_drawLines8(void* self, void** pointPairs, long long pointPairs_len) {
	QList<QPoint> pointPairs_list; pointPairs_list.reserve(pointPairs_len);
	for (long long __i = 0; __i < pointPairs_len; __i++) pointPairs_list.append(*(QPoint*)pointPairs[__i]);
	((QPainter*)self)->drawLines(pointPairs_list);
}
void QPainter_drawRect(void* self, void* rect) {
	((QPainter*)self)->drawRect(*(QRectF*)rect);
}
void QPainter_drawRect2(void* self, int x1, int y1, int w, int h) {
	((QPainter*)self)->drawRect(x1, y1, w, h);
}
void QPainter_drawRect3(void* self, void* rect) {
	((QPainter*)self)->drawRect(*(QRect*)rect);
}
void QPainter_drawRects(void* self, void* rects, int rectCount) {
	((QPainter*)self)->drawRects((QRectF*)rects, rectCount);
}
void QPainter_drawRects2(void* self, void** rectangles, long long rectangles_len) {
	QList<QRectF> rectangles_list; rectangles_list.reserve(rectangles_len);
	for (long long __i = 0; __i < rectangles_len; __i++) rectangles_list.append(*(QRectF*)rectangles[__i]);
	((QPainter*)self)->drawRects(rectangles_list);
}
void QPainter_drawRects3(void* self, void* rects, int rectCount) {
	((QPainter*)self)->drawRects((QRect*)rects, rectCount);
}
void QPainter_drawRects4(void* self, void** rectangles, long long rectangles_len) {
	QList<QRect> rectangles_list; rectangles_list.reserve(rectangles_len);
	for (long long __i = 0; __i < rectangles_len; __i++) rectangles_list.append(*(QRect*)rectangles[__i]);
	((QPainter*)self)->drawRects(rectangles_list);
}
void QPainter_drawEllipse(void* self, void* r) {
	((QPainter*)self)->drawEllipse(*(QRectF*)r);
}
void QPainter_drawEllipse2(void* self, void* r) {
	((QPainter*)self)->drawEllipse(*(QRect*)r);
}
void QPainter_drawEllipse3(void* self, int x, int y, int w, int h) {
	((QPainter*)self)->drawEllipse(x, y, w, h);
}
void QPainter_drawEllipse4(void* self, void* center, double rx, double ry) {
	((QPainter*)self)->drawEllipse(*(QPointF*)center, rx, ry);
}
void QPainter_drawEllipse5(void* self, void* center, int rx, int ry) {
	((QPainter*)self)->drawEllipse(*(QPoint*)center, rx, ry);
}
void QPainter_drawPolyline(void* self, void* points, int pointCount) {
	((QPainter*)self)->drawPolyline((QPointF*)points, pointCount);
}
void QPainter_drawPolyline2(void* self, void* polyline) {
	((QPainter*)self)->drawPolyline(*(QPolygonF*)polyline);
}
void QPainter_drawPolyline3(void* self, void* points, int pointCount) {
	((QPainter*)self)->drawPolyline((QPoint*)points, pointCount);
}
void QPainter_drawPolyline4(void* self, void* polygon) {
	((QPainter*)self)->drawPolyline(*(QPolygon*)polygon);
}
void QPainter_drawPolygon(void* self, void* points, int pointCount) {
	((QPainter*)self)->drawPolygon((QPointF*)points, pointCount);
}
void QPainter_drawPolygon2(void* self, void* polygon) {
	((QPainter*)self)->drawPolygon(*(QPolygonF*)polygon);
}
void QPainter_drawPolygon3(void* self, void* points, int pointCount) {
	((QPainter*)self)->drawPolygon((QPoint*)points, pointCount);
}
void QPainter_drawPolygon4(void* self, void* polygon) {
	((QPainter*)self)->drawPolygon(*(QPolygon*)polygon);
}
void QPainter_drawConvexPolygon(void* self, void* points, int pointCount) {
	((QPainter*)self)->drawConvexPolygon((QPointF*)points, pointCount);
}
void QPainter_drawConvexPolygon2(void* self, void* polygon) {
	((QPainter*)self)->drawConvexPolygon(*(QPolygonF*)polygon);
}
void QPainter_drawConvexPolygon3(void* self, void* points, int pointCount) {
	((QPainter*)self)->drawConvexPolygon((QPoint*)points, pointCount);
}
void QPainter_drawConvexPolygon4(void* self, void* polygon) {
	((QPainter*)self)->drawConvexPolygon(*(QPolygon*)polygon);
}
void QPainter_drawArc(void* self, void* rect, int a, int alen) {
	((QPainter*)self)->drawArc(*(QRectF*)rect, a, alen);
}
void QPainter_drawArc2(void* self, void* arg1, int a, int alen) {
	((QPainter*)self)->drawArc(*(QRect*)arg1, a, alen);
}
void QPainter_drawArc3(void* self, int x, int y, int w, int h, int a, int alen) {
	((QPainter*)self)->drawArc(x, y, w, h, a, alen);
}
void QPainter_drawPie(void* self, void* rect, int a, int alen) {
	((QPainter*)self)->drawPie(*(QRectF*)rect, a, alen);
}
void QPainter_drawPie2(void* self, int x, int y, int w, int h, int a, int alen) {
	((QPainter*)self)->drawPie(x, y, w, h, a, alen);
}
void QPainter_drawPie3(void* self, void* arg1, int a, int alen) {
	((QPainter*)self)->drawPie(*(QRect*)arg1, a, alen);
}
void QPainter_drawChord(void* self, void* rect, int a, int alen) {
	((QPainter*)self)->drawChord(*(QRectF*)rect, a, alen);
}
void QPainter_drawChord2(void* self, int x, int y, int w, int h, int a, int alen) {
	((QPainter*)self)->drawChord(x, y, w, h, a, alen);
}
void QPainter_drawChord3(void* self, void* arg1, int a, int alen) {
	((QPainter*)self)->drawChord(*(QRect*)arg1, a, alen);
}
void QPainter_drawRoundedRect(void* self, void* rect, double xRadius, double yRadius) {
	((QPainter*)self)->drawRoundedRect(*(QRectF*)rect, xRadius, yRadius);
}
void QPainter_drawRoundedRect2(void* self, int x, int y, int w, int h, double xRadius, double yRadius) {
	((QPainter*)self)->drawRoundedRect(x, y, w, h, xRadius, yRadius);
}
void QPainter_drawRoundedRect3(void* self, void* rect, double xRadius, double yRadius) {
	((QPainter*)self)->drawRoundedRect(*(QRect*)rect, xRadius, yRadius);
}
void QPainter_drawTiledPixmap(void* self, void* rect, void* pm) {
	((QPainter*)self)->drawTiledPixmap(*(QRectF*)rect, *(QPixmap*)pm);
}
void QPainter_drawTiledPixmap2(void* self, int x, int y, int w, int h, void* arg5) {
	((QPainter*)self)->drawTiledPixmap(x, y, w, h, *(QPixmap*)arg5);
}
void QPainter_drawTiledPixmap3(void* self, void* arg1, void* arg2) {
	((QPainter*)self)->drawTiledPixmap(*(QRect*)arg1, *(QPixmap*)arg2);
}
void QPainter_drawPicture(void* self, void* p, void* picture) {
	((QPainter*)self)->drawPicture(*(QPointF*)p, *(QPicture*)picture);
}
void QPainter_drawPicture2(void* self, int x, int y, void* picture) {
	((QPainter*)self)->drawPicture(x, y, *(QPicture*)picture);
}
void QPainter_drawPicture3(void* self, void* p, void* picture) {
	((QPainter*)self)->drawPicture(*(QPoint*)p, *(QPicture*)picture);
}
void QPainter_drawPixmap(void* self, void* targetRect, void* pixmap, void* sourceRect) {
	((QPainter*)self)->drawPixmap(*(QRectF*)targetRect, *(QPixmap*)pixmap, *(QRectF*)sourceRect);
}
void QPainter_drawPixmap2(void* self, void* targetRect, void* pixmap, void* sourceRect) {
	((QPainter*)self)->drawPixmap(*(QRect*)targetRect, *(QPixmap*)pixmap, *(QRect*)sourceRect);
}
void QPainter_drawPixmap3(void* self, int x, int y, int w, int h, void* pm, int sx, int sy, int sw, int sh) {
	((QPainter*)self)->drawPixmap(x, y, w, h, *(QPixmap*)pm, sx, sy, sw, sh);
}
void QPainter_drawPixmap4(void* self, int x, int y, void* pm, int sx, int sy, int sw, int sh) {
	((QPainter*)self)->drawPixmap(x, y, *(QPixmap*)pm, sx, sy, sw, sh);
}
void QPainter_drawPixmap5(void* self, void* p, void* pm, void* sr) {
	((QPainter*)self)->drawPixmap(*(QPointF*)p, *(QPixmap*)pm, *(QRectF*)sr);
}
void QPainter_drawPixmap6(void* self, void* p, void* pm, void* sr) {
	((QPainter*)self)->drawPixmap(*(QPoint*)p, *(QPixmap*)pm, *(QRect*)sr);
}
void QPainter_drawPixmap7(void* self, void* p, void* pm) {
	((QPainter*)self)->drawPixmap(*(QPointF*)p, *(QPixmap*)pm);
}
void QPainter_drawPixmap8(void* self, void* p, void* pm) {
	((QPainter*)self)->drawPixmap(*(QPoint*)p, *(QPixmap*)pm);
}
void QPainter_drawPixmap9(void* self, int x, int y, void* pm) {
	((QPainter*)self)->drawPixmap(x, y, *(QPixmap*)pm);
}
void QPainter_drawPixmap10(void* self, void* r, void* pm) {
	((QPainter*)self)->drawPixmap(*(QRect*)r, *(QPixmap*)pm);
}
void QPainter_drawPixmap11(void* self, int x, int y, int w, int h, void* pm) {
	((QPainter*)self)->drawPixmap(x, y, w, h, *(QPixmap*)pm);
}
void QPainter_drawPixmapFragments(void* self, void* fragments, int fragmentCount, void* pixmap) {
	((QPainter*)self)->drawPixmapFragments((QPainter::PixmapFragment*)fragments, fragmentCount, *(QPixmap*)pixmap);
}
void QPainter_drawImage(void* self, void* targetRect, void* image, void* sourceRect) {
	((QPainter*)self)->drawImage(*(QRectF*)targetRect, *(QImage*)image, *(QRectF*)sourceRect);
}
void QPainter_drawImage2(void* self, void* targetRect, void* image, void* sourceRect) {
	((QPainter*)self)->drawImage(*(QRect*)targetRect, *(QImage*)image, *(QRect*)sourceRect);
}
void QPainter_drawImage3(void* self, void* p, void* image, void* sr) {
	((QPainter*)self)->drawImage(*(QPointF*)p, *(QImage*)image, *(QRectF*)sr);
}
void QPainter_drawImage4(void* self, void* p, void* image, void* sr) {
	((QPainter*)self)->drawImage(*(QPoint*)p, *(QImage*)image, *(QRect*)sr);
}
void QPainter_drawImage5(void* self, void* r, void* image) {
	((QPainter*)self)->drawImage(*(QRectF*)r, *(QImage*)image);
}
void QPainter_drawImage6(void* self, void* r, void* image) {
	((QPainter*)self)->drawImage(*(QRect*)r, *(QImage*)image);
}
void QPainter_drawImage7(void* self, void* p, void* image) {
	((QPainter*)self)->drawImage(*(QPointF*)p, *(QImage*)image);
}
void QPainter_drawImage8(void* self, void* p, void* image) {
	((QPainter*)self)->drawImage(*(QPoint*)p, *(QImage*)image);
}
void QPainter_drawImage9(void* self, int x, int y, void* image) {
	((QPainter*)self)->drawImage(x, y, *(QImage*)image);
}
void QPainter_setLayoutDirection(void* self, int direction) {
	((QPainter*)self)->setLayoutDirection((Qt::LayoutDirection)(direction));
}
int QPainter_layoutDirection(void* self) {
	auto&& __r = ((QPainter*)self)->layoutDirection();
	return (int)(__r);
}
void QPainter_drawGlyphRun(void* self, void* position, void* glyphRun) {
	((QPainter*)self)->drawGlyphRun(*(QPointF*)position, *(QGlyphRun*)glyphRun);
}
void QPainter_drawStaticText(void* self, void* topLeftPosition, void* staticText) {
	((QPainter*)self)->drawStaticText(*(QPointF*)topLeftPosition, *(QStaticText*)staticText);
}
void QPainter_drawStaticText2(void* self, void* topLeftPosition, void* staticText) {
	((QPainter*)self)->drawStaticText(*(QPoint*)topLeftPosition, *(QStaticText*)staticText);
}
void QPainter_drawStaticText3(void* self, int left, int top, void* staticText) {
	((QPainter*)self)->drawStaticText(left, top, *(QStaticText*)staticText);
}
void QPainter_drawText(void* self, void* p, const char* s, long long s_len) {
	QString s_qstr = QString::fromUtf8(s, s_len);
	((QPainter*)self)->drawText(*(QPointF*)p, s_qstr);
}
void QPainter_drawText2(void* self, void* p, const char* s, long long s_len) {
	QString s_qstr = QString::fromUtf8(s, s_len);
	((QPainter*)self)->drawText(*(QPoint*)p, s_qstr);
}
void QPainter_drawText3(void* self, int x, int y, const char* s, long long s_len) {
	QString s_qstr = QString::fromUtf8(s, s_len);
	((QPainter*)self)->drawText(x, y, s_qstr);
}
void QPainter_drawText4(void* self, void* p, const char* str, long long str_len, int tf, int justificationPadding) {
	QString str_qstr = QString::fromUtf8(str, str_len);
	((QPainter*)self)->drawText(*(QPointF*)p, str_qstr, tf, justificationPadding);
}
void QPainter_drawText5(void* self, void* r, int flags, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	((QPainter*)self)->drawText(*(QRectF*)r, flags, text_qstr);
}
void QPainter_drawText6(void* self, void* r, int flags, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	((QPainter*)self)->drawText(*(QRect*)r, flags, text_qstr);
}
void QPainter_drawText7(void* self, int x, int y, int w, int h, int flags, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	((QPainter*)self)->drawText(x, y, w, h, flags, text_qstr);
}
void QPainter_drawText8(void* self, void* r, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	((QPainter*)self)->drawText(*(QRectF*)r, text_qstr);
}
void* QPainter_boundingRect(void* self, void* rect, int flags, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = ((QPainter*)self)->boundingRect(*(QRectF*)rect, flags, text_qstr);
	return new QRectF(__r);
}
void* QPainter_boundingRect2(void* self, void* rect, int flags, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = ((QPainter*)self)->boundingRect(*(QRect*)rect, flags, text_qstr);
	return new QRect(__r);
}
void* QPainter_boundingRect3(void* self, int x, int y, int w, int h, int flags, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = ((QPainter*)self)->boundingRect(x, y, w, h, flags, text_qstr);
	return new QRect(__r);
}
void* QPainter_boundingRect4(void* self, void* rect, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = ((QPainter*)self)->boundingRect(*(QRectF*)rect, text_qstr);
	return new QRectF(__r);
}
void QPainter_drawTextItem(void* self, void* p, void* ti) {
	((QPainter*)self)->drawTextItem(*(QPointF*)p, *(QTextItem*)ti);
}
void QPainter_drawTextItem2(void* self, int x, int y, void* ti) {
	((QPainter*)self)->drawTextItem(x, y, *(QTextItem*)ti);
}
void QPainter_drawTextItem3(void* self, void* p, void* ti) {
	((QPainter*)self)->drawTextItem(*(QPoint*)p, *(QTextItem*)ti);
}
void QPainter_fillRect(void* self, void* arg1, void* arg2) {
	((QPainter*)self)->fillRect(*(QRectF*)arg1, *(QBrush*)arg2);
}
void QPainter_fillRect2(void* self, int x, int y, int w, int h, void* arg5) {
	((QPainter*)self)->fillRect(x, y, w, h, *(QBrush*)arg5);
}
void QPainter_fillRect3(void* self, void* arg1, void* arg2) {
	((QPainter*)self)->fillRect(*(QRect*)arg1, *(QBrush*)arg2);
}
void QPainter_fillRect4(void* self, void* arg1, void* color) {
	((QPainter*)self)->fillRect(*(QRectF*)arg1, *(QColor*)color);
}
void QPainter_fillRect5(void* self, int x, int y, int w, int h, void* color) {
	((QPainter*)self)->fillRect(x, y, w, h, *(QColor*)color);
}
void QPainter_fillRect6(void* self, void* arg1, void* color) {
	((QPainter*)self)->fillRect(*(QRect*)arg1, *(QColor*)color);
}
void QPainter_fillRect7(void* self, int x, int y, int w, int h, int c) {
	((QPainter*)self)->fillRect(x, y, w, h, (Qt::GlobalColor)(c));
}
void QPainter_fillRect8(void* self, void* r, int c) {
	((QPainter*)self)->fillRect(*(QRect*)r, (Qt::GlobalColor)(c));
}
void QPainter_fillRect9(void* self, void* r, int c) {
	((QPainter*)self)->fillRect(*(QRectF*)r, (Qt::GlobalColor)(c));
}
void QPainter_fillRect10(void* self, int x, int y, int w, int h, int style) {
	((QPainter*)self)->fillRect(x, y, w, h, (Qt::BrushStyle)(style));
}
void QPainter_fillRect11(void* self, void* r, int style) {
	((QPainter*)self)->fillRect(*(QRect*)r, (Qt::BrushStyle)(style));
}
void QPainter_fillRect12(void* self, void* r, int style) {
	((QPainter*)self)->fillRect(*(QRectF*)r, (Qt::BrushStyle)(style));
}
void QPainter_fillRect13(void* self, int x, int y, int w, int h, int preset) {
	((QPainter*)self)->fillRect(x, y, w, h, (QGradient::Preset)(preset));
}
void QPainter_fillRect14(void* self, void* r, int preset) {
	((QPainter*)self)->fillRect(*(QRect*)r, (QGradient::Preset)(preset));
}
void QPainter_fillRect15(void* self, void* r, int preset) {
	((QPainter*)self)->fillRect(*(QRectF*)r, (QGradient::Preset)(preset));
}
void QPainter_eraseRect(void* self, void* arg1) {
	((QPainter*)self)->eraseRect(*(QRectF*)arg1);
}
void QPainter_eraseRect2(void* self, int x, int y, int w, int h) {
	((QPainter*)self)->eraseRect(x, y, w, h);
}
void QPainter_eraseRect3(void* self, void* arg1) {
	((QPainter*)self)->eraseRect(*(QRect*)arg1);
}
void QPainter_setRenderHint(void* self, int hint) {
	((QPainter*)self)->setRenderHint((QPainter::RenderHint)(hint));
}
void QPainter_setRenderHints(void* self, int hints) {
	((QPainter*)self)->setRenderHints(QFlags<QPainter::RenderHint>(hints));
}
int QPainter_renderHints(void* self) {
	auto&& __r = ((QPainter*)self)->renderHints();
	return (int)(__r);
}
int QPainter_testRenderHint(void* self, int hint) {
	auto&& __r = ((QPainter*)self)->testRenderHint((QPainter::RenderHint)(hint));
	return __r;
}
void* QPainter_paintEngine(void* self) {
	auto&& __r = ((QPainter*)self)->paintEngine();
	return (void*)(__r);
}
void QPainter_beginNativePainting(void* self) {
	((QPainter*)self)->beginNativePainting();
}
void QPainter_endNativePainting(void* self) {
	((QPainter*)self)->endNativePainting();
}
void QPainter_setClipRect4(void* self, void* arg1, int op) {
	((QPainter*)self)->setClipRect(*(QRectF*)arg1, (Qt::ClipOperation)(op));
}
void QPainter_setClipRect5(void* self, void* arg1, int op) {
	((QPainter*)self)->setClipRect(*(QRect*)arg1, (Qt::ClipOperation)(op));
}
void QPainter_setClipRect6(void* self, int x, int y, int w, int h, int op) {
	((QPainter*)self)->setClipRect(x, y, w, h, (Qt::ClipOperation)(op));
}
void QPainter_setClipRegion2(void* self, void* arg1, int op) {
	((QPainter*)self)->setClipRegion(*(QRegion*)arg1, (Qt::ClipOperation)(op));
}
void QPainter_setClipPath2(void* self, void* path, int op) {
	((QPainter*)self)->setClipPath(*(QPainterPath*)path, (Qt::ClipOperation)(op));
}
void QPainter_setTransform2(void* self, void* transform, int combine) {
	((QPainter*)self)->setTransform(*(QTransform*)transform, combine);
}
void QPainter_setWorldTransform2(void* self, void* matrix, int combine) {
	((QPainter*)self)->setWorldTransform(*(QTransform*)matrix, combine);
}
void QPainter_drawPolygon5(void* self, void* points, int pointCount, int fillRule) {
	((QPainter*)self)->drawPolygon((QPointF*)points, pointCount, (Qt::FillRule)(fillRule));
}
void QPainter_drawPolygon6(void* self, void* polygon, int fillRule) {
	((QPainter*)self)->drawPolygon(*(QPolygonF*)polygon, (Qt::FillRule)(fillRule));
}
void QPainter_drawPolygon7(void* self, void* points, int pointCount, int fillRule) {
	((QPainter*)self)->drawPolygon((QPoint*)points, pointCount, (Qt::FillRule)(fillRule));
}
void QPainter_drawPolygon8(void* self, void* polygon, int fillRule) {
	((QPainter*)self)->drawPolygon(*(QPolygon*)polygon, (Qt::FillRule)(fillRule));
}
void QPainter_drawRoundedRect4(void* self, void* rect, double xRadius, double yRadius, int mode) {
	((QPainter*)self)->drawRoundedRect(*(QRectF*)rect, xRadius, yRadius, (Qt::SizeMode)(mode));
}
void QPainter_drawRoundedRect5(void* self, int x, int y, int w, int h, double xRadius, double yRadius, int mode) {
	((QPainter*)self)->drawRoundedRect(x, y, w, h, xRadius, yRadius, (Qt::SizeMode)(mode));
}
void QPainter_drawRoundedRect6(void* self, void* rect, double xRadius, double yRadius, int mode) {
	((QPainter*)self)->drawRoundedRect(*(QRect*)rect, xRadius, yRadius, (Qt::SizeMode)(mode));
}
void QPainter_drawTiledPixmap4(void* self, void* rect, void* pm, void* offset) {
	((QPainter*)self)->drawTiledPixmap(*(QRectF*)rect, *(QPixmap*)pm, *(QPointF*)offset);
}
void QPainter_drawTiledPixmap5(void* self, int x, int y, int w, int h, void* arg5, int sx) {
	((QPainter*)self)->drawTiledPixmap(x, y, w, h, *(QPixmap*)arg5, sx);
}
void QPainter_drawTiledPixmap6(void* self, int x, int y, int w, int h, void* arg5, int sx, int sy) {
	((QPainter*)self)->drawTiledPixmap(x, y, w, h, *(QPixmap*)arg5, sx, sy);
}
void QPainter_drawTiledPixmap7(void* self, void* arg1, void* arg2, void* arg3) {
	((QPainter*)self)->drawTiledPixmap(*(QRect*)arg1, *(QPixmap*)arg2, *(QPoint*)arg3);
}
void QPainter_drawPixmapFragments2(void* self, void* fragments, int fragmentCount, void* pixmap, int hints) {
	((QPainter*)self)->drawPixmapFragments((QPainter::PixmapFragment*)fragments, fragmentCount, *(QPixmap*)pixmap, QFlags<QPainter::PixmapFragmentHint>(hints));
}
void QPainter_drawImage10(void* self, void* targetRect, void* image, void* sourceRect, int flags) {
	((QPainter*)self)->drawImage(*(QRectF*)targetRect, *(QImage*)image, *(QRectF*)sourceRect, QFlags<Qt::ImageConversionFlag>(flags));
}
void QPainter_drawImage11(void* self, void* targetRect, void* image, void* sourceRect, int flags) {
	((QPainter*)self)->drawImage(*(QRect*)targetRect, *(QImage*)image, *(QRect*)sourceRect, QFlags<Qt::ImageConversionFlag>(flags));
}
void QPainter_drawImage12(void* self, void* p, void* image, void* sr, int flags) {
	((QPainter*)self)->drawImage(*(QPointF*)p, *(QImage*)image, *(QRectF*)sr, QFlags<Qt::ImageConversionFlag>(flags));
}
void QPainter_drawImage13(void* self, void* p, void* image, void* sr, int flags) {
	((QPainter*)self)->drawImage(*(QPoint*)p, *(QImage*)image, *(QRect*)sr, QFlags<Qt::ImageConversionFlag>(flags));
}
void QPainter_drawImage14(void* self, int x, int y, void* image, int sx) {
	((QPainter*)self)->drawImage(x, y, *(QImage*)image, sx);
}
void QPainter_drawImage15(void* self, int x, int y, void* image, int sx, int sy) {
	((QPainter*)self)->drawImage(x, y, *(QImage*)image, sx, sy);
}
void QPainter_drawImage16(void* self, int x, int y, void* image, int sx, int sy, int sw) {
	((QPainter*)self)->drawImage(x, y, *(QImage*)image, sx, sy, sw);
}
void QPainter_drawImage17(void* self, int x, int y, void* image, int sx, int sy, int sw, int sh) {
	((QPainter*)self)->drawImage(x, y, *(QImage*)image, sx, sy, sw, sh);
}
void QPainter_drawImage18(void* self, int x, int y, void* image, int sx, int sy, int sw, int sh, int flags) {
	((QPainter*)self)->drawImage(x, y, *(QImage*)image, sx, sy, sw, sh, QFlags<Qt::ImageConversionFlag>(flags));
}
void QPainter_drawText9(void* self, void* r, int flags, const char* text, long long text_len, void* br) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	((QPainter*)self)->drawText(*(QRectF*)r, flags, text_qstr, (QRectF*)br);
}
void QPainter_drawText10(void* self, void* r, int flags, const char* text, long long text_len, void* br) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	((QPainter*)self)->drawText(*(QRect*)r, flags, text_qstr, (QRect*)br);
}
void QPainter_drawText11(void* self, int x, int y, int w, int h, int flags, const char* text, long long text_len, void* br) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	((QPainter*)self)->drawText(x, y, w, h, flags, text_qstr, (QRect*)br);
}
void QPainter_drawText12(void* self, void* r, const char* text, long long text_len, void* o) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	((QPainter*)self)->drawText(*(QRectF*)r, text_qstr, *(QTextOption*)o);
}
void* QPainter_boundingRect5(void* self, void* rect, const char* text, long long text_len, void* o) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = ((QPainter*)self)->boundingRect(*(QRectF*)rect, text_qstr, *(QTextOption*)o);
	return new QRectF(__r);
}
void QPainter_setRenderHint2(void* self, int hint, int on) {
	((QPainter*)self)->setRenderHint((QPainter::RenderHint)(hint), on);
}
void QPainter_setRenderHints2(void* self, int hints, int on) {
	((QPainter*)self)->setRenderHints(QFlags<QPainter::RenderHint>(hints), on);
}
void QPainter_delete(void* self) { delete (QPainter*)self; }

}

#include <QtGui/qtextobject.h>
#include <QtCore/qobject.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtGui/qtextformat.h>
#include <QtGui/qtextdocument.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qvariant.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qtextobject.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

static std::unordered_map<long long, QMetaObject::Connection> __conns_QTextObject;
static std::mutex __connsMu_QTextObject;

extern "C" {

void* QTextObject_base(void* self) { return (void*)static_cast<QObject*>((QTextObject*)self); }
void* QTextObject_fromQObject(void* o) { return (void*)qobject_cast<QTextObject*>((QObject*)o); }
char* QTextObject_tr(const char* s, long long* __out_len) {
	auto&& __r = QTextObject::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QTextObject_format(void* self) {
	auto&& __r = ((QTextObject*)self)->format();
	return new QTextFormat(__r);
}
int QTextObject_formatIndex(void* self) {
	auto&& __r = ((QTextObject*)self)->formatIndex();
	return __r;
}
void* QTextObject_document(void* self) {
	auto&& __r = ((QTextObject*)self)->document();
	return (void*)(__r);
}
int QTextObject_objectIndex(void* self) {
	auto&& __r = ((QTextObject*)self)->objectIndex();
	return __r;
}
void* QTextObject_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QTextObject::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QTextObject_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QTextObject::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
char* QTextObject_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QTextObject::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QTextObject_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QTextObject::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QTextObject_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QTextObject::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QTextObject_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QTextObject::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void QTextObject_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTextObject*)self, &QTextObject::destroyed, (QTextObject*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTextObject);
	__conns_QTextObject[__id] = __c;
}
void QTextObject_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTextObject*)self, &QTextObject::destroyed, (QTextObject*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTextObject);
	__conns_QTextObject[__id] = __c;
}
void QTextObject_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTextObject*)self, &QTextObject::objectNameChanged, (QTextObject*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTextObject);
	__conns_QTextObject[__id] = __c;
}
void QTextObject__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QTextObject);
		auto __it = __conns_QTextObject.find(id);
		if (__it == __conns_QTextObject.end()) return;
		__c = __it->second;
		__conns_QTextObject.erase(__it);
	}
	QObject::disconnect(__c);
}

}

#include <QtWidgets/qabstractbutton.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtGui/qicon.h>
#include <QtCore/qsize.h>
#include <QtGui/qkeysequence.h>
#include <QtWidgets/qbuttongroup.h>
#include <QtGui/qevent.h>
#include <QtCore/qpoint.h>
#include <QtCore/qcoreevent.h>
#include <QtWidgets/qstyle.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qrect.h>
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
#include <QtCore/qlocale.h>
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
#include "gen_qabstractbutton.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QAbstractButton__Go : public QAbstractButton {
public:
	long long __goId;
	unsigned char __goOn[51] = {};
	template<typename... A> QAbstractButton__Go(long long id, A&&... a) : QAbstractButton(std::forward<A>(a)...), __goId(id) {}
	using QAbstractButton::updateMicroFocus;
	using QAbstractButton::create;
	using QAbstractButton::destroy;
	using QAbstractButton::focusNextChild;
	using QAbstractButton::focusPreviousChild;
	using QAbstractButton::sender;
	using QAbstractButton::senderSignalIndex;
	using QAbstractButton::receivers;
	using QAbstractButton::isSignalConnected;
	using QAbstractButton::getDecodedMetricF;
	bool cppgoBase_hitButton(const QPoint& __a0) { return QAbstractButton::hitButton(__a0); }
	void cppgoBase_checkStateSet() { QAbstractButton::checkStateSet(); }
	void cppgoBase_nextCheckState() { QAbstractButton::nextCheckState(); }
	bool cppgoBase_event(QEvent* __a0) { return QAbstractButton::event(__a0); }
	void cppgoBase_keyPressEvent(QKeyEvent* __a0) { QAbstractButton::keyPressEvent(__a0); }
	void cppgoBase_keyReleaseEvent(QKeyEvent* __a0) { QAbstractButton::keyReleaseEvent(__a0); }
	void cppgoBase_mousePressEvent(QMouseEvent* __a0) { QAbstractButton::mousePressEvent(__a0); }
	void cppgoBase_mouseReleaseEvent(QMouseEvent* __a0) { QAbstractButton::mouseReleaseEvent(__a0); }
	void cppgoBase_mouseMoveEvent(QMouseEvent* __a0) { QAbstractButton::mouseMoveEvent(__a0); }
	void cppgoBase_focusInEvent(QFocusEvent* __a0) { QAbstractButton::focusInEvent(__a0); }
	void cppgoBase_focusOutEvent(QFocusEvent* __a0) { QAbstractButton::focusOutEvent(__a0); }
	void cppgoBase_changeEvent(QEvent* __a0) { QAbstractButton::changeEvent(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QAbstractButton::timerEvent(__a0); }
	void cppgoBase_mouseDoubleClickEvent(QMouseEvent* __a0) { QAbstractButton::mouseDoubleClickEvent(__a0); }
	void cppgoBase_wheelEvent(QWheelEvent* __a0) { QAbstractButton::wheelEvent(__a0); }
	void cppgoBase_enterEvent(QEnterEvent* __a0) { QAbstractButton::enterEvent(__a0); }
	void cppgoBase_leaveEvent(QEvent* __a0) { QAbstractButton::leaveEvent(__a0); }
	void cppgoBase_moveEvent(QMoveEvent* __a0) { QAbstractButton::moveEvent(__a0); }
	void cppgoBase_resizeEvent(QResizeEvent* __a0) { QAbstractButton::resizeEvent(__a0); }
	void cppgoBase_closeEvent(QCloseEvent* __a0) { QAbstractButton::closeEvent(__a0); }
	void cppgoBase_contextMenuEvent(QContextMenuEvent* __a0) { QAbstractButton::contextMenuEvent(__a0); }
	void cppgoBase_tabletEvent(QTabletEvent* __a0) { QAbstractButton::tabletEvent(__a0); }
	void cppgoBase_actionEvent(QActionEvent* __a0) { QAbstractButton::actionEvent(__a0); }
	void cppgoBase_dragEnterEvent(QDragEnterEvent* __a0) { QAbstractButton::dragEnterEvent(__a0); }
	void cppgoBase_dragMoveEvent(QDragMoveEvent* __a0) { QAbstractButton::dragMoveEvent(__a0); }
	void cppgoBase_dragLeaveEvent(QDragLeaveEvent* __a0) { QAbstractButton::dragLeaveEvent(__a0); }
	void cppgoBase_dropEvent(QDropEvent* __a0) { QAbstractButton::dropEvent(__a0); }
	void cppgoBase_showEvent(QShowEvent* __a0) { QAbstractButton::showEvent(__a0); }
	void cppgoBase_hideEvent(QHideEvent* __a0) { QAbstractButton::hideEvent(__a0); }
	int cppgoBase_metric(int __a0) { return QAbstractButton::metric((QPaintDevice::PaintDeviceMetric)__a0); }
	void cppgoBase_initPainter(QPainter* __a0) { QAbstractButton::initPainter(__a0); }
	QPaintDevice* cppgoBase_redirected(QPoint* __a0) { return QAbstractButton::redirected(__a0); }
	QPainter* cppgoBase_sharedPainter() { return QAbstractButton::sharedPainter(); }
	void cppgoBase_inputMethodEvent(QInputMethodEvent* __a0) { QAbstractButton::inputMethodEvent(__a0); }
	bool cppgoBase_focusNextPrevChild(bool __a0) { return QAbstractButton::focusNextPrevChild(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QAbstractButton::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QAbstractButton::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QAbstractButton::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QAbstractButton::disconnectNotify(__a0); }
	void paintEvent(QPaintEvent* e) override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5290442163167458465LL, (void*)e, &__h); if (__h) return; }
		return;
	}
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QAbstractButton::metaObject();
	}
	bool hitButton(const QPoint& pos) const override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 5736040353629178930LL, (void*)&pos, &__h); if (__h) return __r; }
		return QAbstractButton::hitButton(pos);
	}
	void checkStateSet() override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 5747806663224476200LL, &__h); if (__h) return; }
		QAbstractButton::checkStateSet();
	}
	void nextCheckState() override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 1474647914901213797LL, &__h); if (__h) return; }
		QAbstractButton::nextCheckState();
	}
	bool event(QEvent* e) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)e, &__h); if (__h) return __r; }
		return QAbstractButton::event(e);
	}
	void keyPressEvent(QKeyEvent* e) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1023791977679793185LL, (void*)e, &__h); if (__h) return; }
		QAbstractButton::keyPressEvent(e);
	}
	void keyReleaseEvent(QKeyEvent* e) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 772598475950050299LL, (void*)e, &__h); if (__h) return; }
		QAbstractButton::keyReleaseEvent(e);
	}
	void mousePressEvent(QMouseEvent* e) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -332090898078543319LL, (void*)e, &__h); if (__h) return; }
		QAbstractButton::mousePressEvent(e);
	}
	void mouseReleaseEvent(QMouseEvent* e) override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4415920049297795213LL, (void*)e, &__h); if (__h) return; }
		QAbstractButton::mouseReleaseEvent(e);
	}
	void mouseMoveEvent(QMouseEvent* e) override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4255494004099806581LL, (void*)e, &__h); if (__h) return; }
		QAbstractButton::mouseMoveEvent(e);
	}
	void focusInEvent(QFocusEvent* e) override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 975043767033552552LL, (void*)e, &__h); if (__h) return; }
		QAbstractButton::focusInEvent(e);
	}
	void focusOutEvent(QFocusEvent* e) override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -385596010576048993LL, (void*)e, &__h); if (__h) return; }
		QAbstractButton::focusOutEvent(e);
	}
	void changeEvent(QEvent* e) override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4959376063826165943LL, (void*)e, &__h); if (__h) return; }
		QAbstractButton::changeEvent(e);
	}
	void timerEvent(QTimerEvent* e) override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)e, &__h); if (__h) return; }
		QAbstractButton::timerEvent(e);
	}
	int devType() const override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8659679893868524044LL, &__h); if (__h) return __r; }
		return QAbstractButton::devType();
	}
	void setVisible(bool visible) override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 2110182719375245575LL, visible, &__h); if (__h) return; }
		QAbstractButton::setVisible(visible);
	}
	QSize sizeHint() const override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 729690261457684101LL, &__h); if (__h) return *(QSize*)__r; }
		return QAbstractButton::sizeHint();
	}
	QSize minimumSizeHint() const override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 7049264599837532491LL, &__h); if (__h) return *(QSize*)__r; }
		return QAbstractButton::minimumSizeHint();
	}
	int heightForWidth(int arg1) const override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 4311534060495737339LL, arg1, &__h); if (__h) return __r; }
		return QAbstractButton::heightForWidth(arg1);
	}
	bool hasHeightForWidth() const override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8266762397995423915LL, &__h); if (__h) return __r; }
		return QAbstractButton::hasHeightForWidth();
	}
	QPaintEngine* paintEngine() const override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 6398007404756346381LL, &__h); if (__h) return (QPaintEngine*)__r; }
		return QAbstractButton::paintEngine();
	}
	void mouseDoubleClickEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4029073905697023887LL, (void*)event, &__h); if (__h) return; }
		QAbstractButton::mouseDoubleClickEvent(event);
	}
	void wheelEvent(QWheelEvent* event) override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1919988340934511316LL, (void*)event, &__h); if (__h) return; }
		QAbstractButton::wheelEvent(event);
	}
	void enterEvent(QEnterEvent* event) override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6316270346668237895LL, (void*)event, &__h); if (__h) return; }
		QAbstractButton::enterEvent(event);
	}
	void leaveEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -551070009623711640LL, (void*)event, &__h); if (__h) return; }
		QAbstractButton::leaveEvent(event);
	}
	void moveEvent(QMoveEvent* event) override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -9161949150489525396LL, (void*)event, &__h); if (__h) return; }
		QAbstractButton::moveEvent(event);
	}
	void resizeEvent(QResizeEvent* event) override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4302917285655014661LL, (void*)event, &__h); if (__h) return; }
		QAbstractButton::resizeEvent(event);
	}
	void closeEvent(QCloseEvent* event) override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1989937050019464783LL, (void*)event, &__h); if (__h) return; }
		QAbstractButton::closeEvent(event);
	}
	void contextMenuEvent(QContextMenuEvent* event) override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6698043222227006015LL, (void*)event, &__h); if (__h) return; }
		QAbstractButton::contextMenuEvent(event);
	}
	void tabletEvent(QTabletEvent* event) override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -212068866860701141LL, (void*)event, &__h); if (__h) return; }
		QAbstractButton::tabletEvent(event);
	}
	void actionEvent(QActionEvent* event) override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1697444140088555581LL, (void*)event, &__h); if (__h) return; }
		QAbstractButton::actionEvent(event);
	}
	void dragEnterEvent(QDragEnterEvent* event) override {
		int __h = 0;
		if (__goOn[32] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 206122715919917589LL, (void*)event, &__h); if (__h) return; }
		QAbstractButton::dragEnterEvent(event);
	}
	void dragMoveEvent(QDragMoveEvent* event) override {
		int __h = 0;
		if (__goOn[33] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1781545137352745978LL, (void*)event, &__h); if (__h) return; }
		QAbstractButton::dragMoveEvent(event);
	}
	void dragLeaveEvent(QDragLeaveEvent* event) override {
		int __h = 0;
		if (__goOn[34] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -3139904158373725006LL, (void*)event, &__h); if (__h) return; }
		QAbstractButton::dragLeaveEvent(event);
	}
	void dropEvent(QDropEvent* event) override {
		int __h = 0;
		if (__goOn[35] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 9107058682754591716LL, (void*)event, &__h); if (__h) return; }
		QAbstractButton::dropEvent(event);
	}
	void showEvent(QShowEvent* event) override {
		int __h = 0;
		if (__goOn[36] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6825369174751916LL, (void*)event, &__h); if (__h) return; }
		QAbstractButton::showEvent(event);
	}
	void hideEvent(QHideEvent* event) override {
		int __h = 0;
		if (__goOn[37] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -199718707876719129LL, (void*)event, &__h); if (__h) return; }
		QAbstractButton::hideEvent(event);
	}
	bool nativeEvent(const QByteArray& eventType, void* message, qintptr* result) override {
		int __h = 0;
		if (__goOn[38] && cppgo_virt_int_voidp_voidp_longlongp) { int __r = ((int(*)(long long, long long, void*, void*, long long*, int*))cppgo_virt_int_voidp_voidp_longlongp)(__goId, -2039166586246443706LL, (void*)&eventType, (void*)message, (long long*)result, &__h); if (__h) return __r; }
		return QAbstractButton::nativeEvent(eventType, message, result);
	}
	int metric(QPaintDevice::PaintDeviceMetric arg1) const override {
		int __h = 0;
		if (__goOn[39] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 427629649267311777LL, (int)(arg1), &__h); if (__h) return __r; }
		return QAbstractButton::metric(arg1);
	}
	void initPainter(QPainter* painter) const override {
		int __h = 0;
		if (__goOn[40] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6477500296980050210LL, (void*)painter, &__h); if (__h) return; }
		QAbstractButton::initPainter(painter);
	}
	QPaintDevice* redirected(QPoint* offset) const override {
		int __h = 0;
		if (__goOn[41] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 4127240633310655048LL, (void*)offset, &__h); if (__h) return (QPaintDevice*)__r; }
		return QAbstractButton::redirected(offset);
	}
	QPainter* sharedPainter() const override {
		int __h = 0;
		if (__goOn[42] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 5792771328741774991LL, &__h); if (__h) return (QPainter*)__r; }
		return QAbstractButton::sharedPainter();
	}
	void inputMethodEvent(QInputMethodEvent* arg1) override {
		int __h = 0;
		if (__goOn[43] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5438819661942843222LL, (void*)arg1, &__h); if (__h) return; }
		QAbstractButton::inputMethodEvent(arg1);
	}
	QVariant inputMethodQuery(Qt::InputMethodQuery arg1) const override {
		int __h = 0;
		if (__goOn[44] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7345853097340709638LL, (int)(arg1), &__h); if (__h) return *(QVariant*)__r; }
		return QAbstractButton::inputMethodQuery(arg1);
	}
	bool focusNextPrevChild(bool next) override {
		int __h = 0;
		if (__goOn[45] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 5898328517505784191LL, next, &__h); if (__h) return __r; }
		return QAbstractButton::focusNextPrevChild(next);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[46] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QAbstractButton::eventFilter(watched, event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[47] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QAbstractButton::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[48] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QAbstractButton::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[49] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QAbstractButton::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[50] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QAbstractButton::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QAbstractButton;
static std::mutex __connsMu_QAbstractButton;

extern "C" {

void* QAbstractButton_base(void* self) { return (void*)static_cast<QWidget*>((QAbstractButton*)self); }
void QAbstractButton_chain(void* self, void** out) {
	QAbstractButton* __p = (QAbstractButton*)self;
	QWidget* __b0 = static_cast<QWidget*>(__p); out[0] = (void*)__b0;
	QObject* __b1 = static_cast<QObject*>(__b0); out[1] = (void*)__b1;
}
void QAbstractButton__set_override(void* self, int idx) { ((QAbstractButton__Go*)self)->__goOn[idx] = 1; }
void* QAbstractButton_fromQObject(void* o) { return (void*)qobject_cast<QAbstractButton*>((QObject*)o); }
char* QAbstractButton_tr(const char* s, long long* __out_len) {
	auto&& __r = QAbstractButton::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QAbstractButton_setText(void* self, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	((QAbstractButton*)self)->setText(text_qstr);
}
char* QAbstractButton_text(void* self, long long* __out_len) {
	auto&& __r = ((QAbstractButton*)self)->text();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QAbstractButton_setIcon(void* self, void* icon) {
	((QAbstractButton*)self)->setIcon(*(QIcon*)icon);
}
void* QAbstractButton_icon(void* self) {
	auto&& __r = ((QAbstractButton*)self)->icon();
	return new QIcon(__r);
}
void* QAbstractButton_iconSize(void* self) {
	auto&& __r = ((QAbstractButton*)self)->iconSize();
	return new QSize(__r);
}
void QAbstractButton_setShortcut(void* self, void* key) {
	((QAbstractButton*)self)->setShortcut(*(QKeySequence*)key);
}
void* QAbstractButton_shortcut(void* self) {
	auto&& __r = ((QAbstractButton*)self)->shortcut();
	return new QKeySequence(__r);
}
void QAbstractButton_setCheckable(void* self, int arg1) {
	((QAbstractButton*)self)->setCheckable(arg1);
}
int QAbstractButton_isCheckable(void* self) {
	auto&& __r = ((QAbstractButton*)self)->isCheckable();
	return __r;
}
int QAbstractButton_isChecked(void* self) {
	auto&& __r = ((QAbstractButton*)self)->isChecked();
	return __r;
}
void QAbstractButton_setDown(void* self, int arg1) {
	((QAbstractButton*)self)->setDown(arg1);
}
int QAbstractButton_isDown(void* self) {
	auto&& __r = ((QAbstractButton*)self)->isDown();
	return __r;
}
void QAbstractButton_setAutoRepeat(void* self, int arg1) {
	((QAbstractButton*)self)->setAutoRepeat(arg1);
}
int QAbstractButton_autoRepeat(void* self) {
	auto&& __r = ((QAbstractButton*)self)->autoRepeat();
	return __r;
}
void QAbstractButton_setAutoRepeatDelay(void* self, int arg1) {
	((QAbstractButton*)self)->setAutoRepeatDelay(arg1);
}
int QAbstractButton_autoRepeatDelay(void* self) {
	auto&& __r = ((QAbstractButton*)self)->autoRepeatDelay();
	return __r;
}
void QAbstractButton_setAutoRepeatInterval(void* self, int arg1) {
	((QAbstractButton*)self)->setAutoRepeatInterval(arg1);
}
int QAbstractButton_autoRepeatInterval(void* self) {
	auto&& __r = ((QAbstractButton*)self)->autoRepeatInterval();
	return __r;
}
void QAbstractButton_setAutoExclusive(void* self, int arg1) {
	((QAbstractButton*)self)->setAutoExclusive(arg1);
}
int QAbstractButton_autoExclusive(void* self) {
	auto&& __r = ((QAbstractButton*)self)->autoExclusive();
	return __r;
}
void* QAbstractButton_group(void* self) {
	auto&& __r = ((QAbstractButton*)self)->group();
	return (void*)(__r);
}
void QAbstractButton_setIconSize(void* self, void* size) {
	((QAbstractButton*)self)->setIconSize(*(QSize*)size);
}
void QAbstractButton_animateClick(void* self) {
	((QAbstractButton*)self)->animateClick();
}
void QAbstractButton_click(void* self) {
	((QAbstractButton*)self)->click();
}
void QAbstractButton_toggle(void* self) {
	((QAbstractButton*)self)->toggle();
}
void QAbstractButton_setChecked(void* self, int arg1) {
	((QAbstractButton*)self)->setChecked(arg1);
}
void QAbstractButton_pressed(void* self) {
	((QAbstractButton*)self)->pressed();
}
void QAbstractButton_released(void* self) {
	((QAbstractButton*)self)->released();
}
void QAbstractButton_clicked(void* self) {
	((QAbstractButton*)self)->clicked();
}
void QAbstractButton_toggled(void* self, int checked) {
	((QAbstractButton*)self)->toggled(checked);
}
void QAbstractButton_setTabOrder(void* arg1, void* arg2) {
	QAbstractButton::setTabOrder((QWidget*)arg1, (QWidget*)arg2);
}
void* QAbstractButton_mouseGrabber() {
	auto&& __r = QAbstractButton::mouseGrabber();
	return (void*)(__r);
}
void* QAbstractButton_keyboardGrabber() {
	auto&& __r = QAbstractButton::keyboardGrabber();
	return (void*)(__r);
}
void* QAbstractButton_find(unsigned long long arg1) {
	auto&& __r = QAbstractButton::find(arg1);
	return (void*)(__r);
}
void* QAbstractButton_createWindowContainer(void* window) {
	auto&& __r = QAbstractButton::createWindowContainer((QWindow*)window);
	return (void*)(__r);
}
void QAbstractButton_updateMicroFocus(void* self) {
	((QAbstractButton__Go*)self)->updateMicroFocus();
}
void QAbstractButton_create(void* self) {
	((QAbstractButton__Go*)self)->create();
}
void QAbstractButton_destroy(void* self) {
	((QAbstractButton__Go*)self)->destroy();
}
int QAbstractButton_focusNextChild(void* self) {
	auto&& __r = ((QAbstractButton__Go*)self)->focusNextChild();
	return __r;
}
int QAbstractButton_focusPreviousChild(void* self) {
	auto&& __r = ((QAbstractButton__Go*)self)->focusPreviousChild();
	return __r;
}
void* QAbstractButton_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QAbstractButton::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QAbstractButton_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QAbstractButton::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QAbstractButton_sender(void* self) {
	auto&& __r = ((QAbstractButton__Go*)self)->sender();
	return (void*)(__r);
}
int QAbstractButton_senderSignalIndex(void* self) {
	auto&& __r = ((QAbstractButton__Go*)self)->senderSignalIndex();
	return __r;
}
int QAbstractButton_receivers(void* self, const char* signal) {
	auto&& __r = ((QAbstractButton__Go*)self)->receivers(signal);
	return __r;
}
int QAbstractButton_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QAbstractButton__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
double QAbstractButton_devicePixelRatioFScale() {
	auto&& __r = QAbstractButton::devicePixelRatioFScale();
	return __r;
}
int QAbstractButton_encodeMetricF(int metric, double value) {
	auto&& __r = QAbstractButton::encodeMetricF((QPaintDevice::PaintDeviceMetric)(metric), value);
	return __r;
}
double QAbstractButton_getDecodedMetricF(void* self, int metricA, int metricB) {
	auto&& __r = ((QAbstractButton__Go*)self)->getDecodedMetricF((QPaintDevice::PaintDeviceMetric)(metricA), (QPaintDevice::PaintDeviceMetric)(metricB));
	return __r;
}
char* QAbstractButton_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QAbstractButton::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QAbstractButton_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QAbstractButton::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QAbstractButton_clicked2(void* self, int checked) {
	((QAbstractButton*)self)->clicked(checked);
}
void* QAbstractButton_createWindowContainer2(void* window, void* parent) {
	auto&& __r = QAbstractButton::createWindowContainer((QWindow*)window, (QWidget*)parent);
	return (void*)(__r);
}
void* QAbstractButton_createWindowContainer3(void* window, void* parent, int flags) {
	auto&& __r = QAbstractButton::createWindowContainer((QWindow*)window, (QWidget*)parent, QFlags<Qt::WindowType>(flags));
	return (void*)(__r);
}
void QAbstractButton_updateMicroFocus2(void* self, int query) {
	((QAbstractButton__Go*)self)->updateMicroFocus((Qt::InputMethodQuery)(query));
}
void QAbstractButton_create2(void* self, unsigned long long arg1) {
	((QAbstractButton__Go*)self)->create(arg1);
}
void QAbstractButton_create3(void* self, unsigned long long arg1, int initializeWindow) {
	((QAbstractButton__Go*)self)->create(arg1, initializeWindow);
}
void QAbstractButton_create4(void* self, unsigned long long arg1, int initializeWindow, int destroyOldWindow) {
	((QAbstractButton__Go*)self)->create(arg1, initializeWindow, destroyOldWindow);
}
void QAbstractButton_destroy2(void* self, int destroyWindow) {
	((QAbstractButton__Go*)self)->destroy(destroyWindow);
}
void QAbstractButton_destroy3(void* self, int destroyWindow, int destroySubWindows) {
	((QAbstractButton__Go*)self)->destroy(destroyWindow, destroySubWindows);
}
void* QAbstractButton_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QAbstractButton::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QAbstractButton_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QAbstractButton::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QAbstractButton_metaObject_default(void* self) {
	auto&& __r = ((QAbstractButton*)self)->QAbstractButton::metaObject();
	return (void*)(__r);
}
int QAbstractButton_hitButton_default(void* self, void* pos) {
	auto&& __r = ((QAbstractButton__Go*)self)->cppgoBase_hitButton(*(QPoint*)pos);
	return __r;
}
void QAbstractButton_checkStateSet_default(void* self) {
	((QAbstractButton__Go*)self)->cppgoBase_checkStateSet();
}
void QAbstractButton_nextCheckState_default(void* self) {
	((QAbstractButton__Go*)self)->cppgoBase_nextCheckState();
}
int QAbstractButton_event_default(void* self, void* e) {
	auto&& __r = ((QAbstractButton__Go*)self)->cppgoBase_event((QEvent*)e);
	return __r;
}
void QAbstractButton_keyPressEvent_default(void* self, void* e) {
	((QAbstractButton__Go*)self)->cppgoBase_keyPressEvent((QKeyEvent*)e);
}
void QAbstractButton_keyReleaseEvent_default(void* self, void* e) {
	((QAbstractButton__Go*)self)->cppgoBase_keyReleaseEvent((QKeyEvent*)e);
}
void QAbstractButton_mousePressEvent_default(void* self, void* e) {
	((QAbstractButton__Go*)self)->cppgoBase_mousePressEvent((QMouseEvent*)e);
}
void QAbstractButton_mouseReleaseEvent_default(void* self, void* e) {
	((QAbstractButton__Go*)self)->cppgoBase_mouseReleaseEvent((QMouseEvent*)e);
}
void QAbstractButton_mouseMoveEvent_default(void* self, void* e) {
	((QAbstractButton__Go*)self)->cppgoBase_mouseMoveEvent((QMouseEvent*)e);
}
void QAbstractButton_focusInEvent_default(void* self, void* e) {
	((QAbstractButton__Go*)self)->cppgoBase_focusInEvent((QFocusEvent*)e);
}
void QAbstractButton_focusOutEvent_default(void* self, void* e) {
	((QAbstractButton__Go*)self)->cppgoBase_focusOutEvent((QFocusEvent*)e);
}
void QAbstractButton_changeEvent_default(void* self, void* e) {
	((QAbstractButton__Go*)self)->cppgoBase_changeEvent((QEvent*)e);
}
void QAbstractButton_timerEvent_default(void* self, void* e) {
	((QAbstractButton__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)e);
}
int QAbstractButton_devType_default(void* self) {
	auto&& __r = ((QAbstractButton*)self)->QAbstractButton::devType();
	return __r;
}
void QAbstractButton_setVisible_default(void* self, int visible) {
	((QAbstractButton*)self)->QAbstractButton::setVisible(visible);
}
void* QAbstractButton_sizeHint_default(void* self) {
	auto&& __r = ((QAbstractButton*)self)->QAbstractButton::sizeHint();
	return new QSize(__r);
}
void* QAbstractButton_minimumSizeHint_default(void* self) {
	auto&& __r = ((QAbstractButton*)self)->QAbstractButton::minimumSizeHint();
	return new QSize(__r);
}
int QAbstractButton_heightForWidth_default(void* self, int arg1) {
	auto&& __r = ((QAbstractButton*)self)->QAbstractButton::heightForWidth(arg1);
	return __r;
}
int QAbstractButton_hasHeightForWidth_default(void* self) {
	auto&& __r = ((QAbstractButton*)self)->QAbstractButton::hasHeightForWidth();
	return __r;
}
void* QAbstractButton_paintEngine_default(void* self) {
	auto&& __r = ((QAbstractButton*)self)->QAbstractButton::paintEngine();
	return (void*)(__r);
}
void QAbstractButton_mouseDoubleClickEvent_default(void* self, void* event) {
	((QAbstractButton__Go*)self)->cppgoBase_mouseDoubleClickEvent((QMouseEvent*)event);
}
void QAbstractButton_wheelEvent_default(void* self, void* event) {
	((QAbstractButton__Go*)self)->cppgoBase_wheelEvent((QWheelEvent*)event);
}
void QAbstractButton_enterEvent_default(void* self, void* event) {
	((QAbstractButton__Go*)self)->cppgoBase_enterEvent((QEnterEvent*)event);
}
void QAbstractButton_leaveEvent_default(void* self, void* event) {
	((QAbstractButton__Go*)self)->cppgoBase_leaveEvent((QEvent*)event);
}
void QAbstractButton_moveEvent_default(void* self, void* event) {
	((QAbstractButton__Go*)self)->cppgoBase_moveEvent((QMoveEvent*)event);
}
void QAbstractButton_resizeEvent_default(void* self, void* event) {
	((QAbstractButton__Go*)self)->cppgoBase_resizeEvent((QResizeEvent*)event);
}
void QAbstractButton_closeEvent_default(void* self, void* event) {
	((QAbstractButton__Go*)self)->cppgoBase_closeEvent((QCloseEvent*)event);
}
void QAbstractButton_contextMenuEvent_default(void* self, void* event) {
	((QAbstractButton__Go*)self)->cppgoBase_contextMenuEvent((QContextMenuEvent*)event);
}
void QAbstractButton_tabletEvent_default(void* self, void* event) {
	((QAbstractButton__Go*)self)->cppgoBase_tabletEvent((QTabletEvent*)event);
}
void QAbstractButton_actionEvent_default(void* self, void* event) {
	((QAbstractButton__Go*)self)->cppgoBase_actionEvent((QActionEvent*)event);
}
void QAbstractButton_dragEnterEvent_default(void* self, void* event) {
	((QAbstractButton__Go*)self)->cppgoBase_dragEnterEvent((QDragEnterEvent*)event);
}
void QAbstractButton_dragMoveEvent_default(void* self, void* event) {
	((QAbstractButton__Go*)self)->cppgoBase_dragMoveEvent((QDragMoveEvent*)event);
}
void QAbstractButton_dragLeaveEvent_default(void* self, void* event) {
	((QAbstractButton__Go*)self)->cppgoBase_dragLeaveEvent((QDragLeaveEvent*)event);
}
void QAbstractButton_dropEvent_default(void* self, void* event) {
	((QAbstractButton__Go*)self)->cppgoBase_dropEvent((QDropEvent*)event);
}
void QAbstractButton_showEvent_default(void* self, void* event) {
	((QAbstractButton__Go*)self)->cppgoBase_showEvent((QShowEvent*)event);
}
void QAbstractButton_hideEvent_default(void* self, void* event) {
	((QAbstractButton__Go*)self)->cppgoBase_hideEvent((QHideEvent*)event);
}
int QAbstractButton_metric_default(void* self, int arg1) {
	auto&& __r = ((QAbstractButton__Go*)self)->cppgoBase_metric(arg1);
	return __r;
}
void QAbstractButton_initPainter_default(void* self, void* painter) {
	((QAbstractButton__Go*)self)->cppgoBase_initPainter((QPainter*)painter);
}
void* QAbstractButton_redirected_default(void* self, void* offset) {
	auto&& __r = ((QAbstractButton__Go*)self)->cppgoBase_redirected((QPoint*)offset);
	return (void*)(__r);
}
void* QAbstractButton_sharedPainter_default(void* self) {
	auto&& __r = ((QAbstractButton__Go*)self)->cppgoBase_sharedPainter();
	return (void*)(__r);
}
void QAbstractButton_inputMethodEvent_default(void* self, void* arg1) {
	((QAbstractButton__Go*)self)->cppgoBase_inputMethodEvent((QInputMethodEvent*)arg1);
}
void* QAbstractButton_inputMethodQuery_default(void* self, int arg1) {
	auto&& __r = ((QAbstractButton*)self)->QAbstractButton::inputMethodQuery((Qt::InputMethodQuery)(arg1));
	return new QVariant(__r);
}
int QAbstractButton_focusNextPrevChild_default(void* self, int next) {
	auto&& __r = ((QAbstractButton__Go*)self)->cppgoBase_focusNextPrevChild(next);
	return __r;
}
int QAbstractButton_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QAbstractButton*)self)->QAbstractButton::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QAbstractButton_childEvent_default(void* self, void* event) {
	((QAbstractButton__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QAbstractButton_customEvent_default(void* self, void* event) {
	((QAbstractButton__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QAbstractButton_connectNotify_default(void* self, void* signal) {
	((QAbstractButton__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QAbstractButton_disconnectNotify_default(void* self, void* signal) {
	((QAbstractButton__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QAbstractButton_connect_pressed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractButton*)self, &QAbstractButton::pressed, (QAbstractButton*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractButton);
	__conns_QAbstractButton[__id] = __c;
}
void QAbstractButton_connect_released(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractButton*)self, &QAbstractButton::released, (QAbstractButton*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractButton);
	__conns_QAbstractButton[__id] = __c;
}
void QAbstractButton_connect_clicked(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractButton*)self, &QAbstractButton::clicked, (QAbstractButton*)self, [__id](bool checked){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractButton);
	__conns_QAbstractButton[__id] = __c;
}
void QAbstractButton_connect_clicked2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractButton*)self, &QAbstractButton::clicked, (QAbstractButton*)self, [__id](bool checked){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, checked);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractButton);
	__conns_QAbstractButton[__id] = __c;
}
void QAbstractButton_connect_toggled(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractButton*)self, &QAbstractButton::toggled, (QAbstractButton*)self, [__id](bool checked){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, checked);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractButton);
	__conns_QAbstractButton[__id] = __c;
}
void QAbstractButton_connect_windowTitleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractButton*)self, &QAbstractButton::windowTitleChanged, (QAbstractButton*)self, [__id](const QString& title){
		QByteArray title_ba = title.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)title_ba.constData(), title_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractButton);
	__conns_QAbstractButton[__id] = __c;
}
void QAbstractButton_connect_windowIconChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractButton*)self, &QAbstractButton::windowIconChanged, (QAbstractButton*)self, [__id](const QIcon& icon){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&icon);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractButton);
	__conns_QAbstractButton[__id] = __c;
}
void QAbstractButton_connect_windowIconTextChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractButton*)self, &QAbstractButton::windowIconTextChanged, (QAbstractButton*)self, [__id](const QString& iconText){
		QByteArray iconText_ba = iconText.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)iconText_ba.constData(), iconText_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractButton);
	__conns_QAbstractButton[__id] = __c;
}
void QAbstractButton_connect_customContextMenuRequested(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractButton*)self, &QAbstractButton::customContextMenuRequested, (QAbstractButton*)self, [__id](const QPoint& pos){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&pos);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractButton);
	__conns_QAbstractButton[__id] = __c;
}
void QAbstractButton_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractButton*)self, &QAbstractButton::destroyed, (QAbstractButton*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractButton);
	__conns_QAbstractButton[__id] = __c;
}
void QAbstractButton_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractButton*)self, &QAbstractButton::destroyed, (QAbstractButton*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractButton);
	__conns_QAbstractButton[__id] = __c;
}
void QAbstractButton_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractButton*)self, &QAbstractButton::objectNameChanged, (QAbstractButton*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractButton);
	__conns_QAbstractButton[__id] = __c;
}
void QAbstractButton__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QAbstractButton);
		auto __it = __conns_QAbstractButton.find(id);
		if (__it == __conns_QAbstractButton.end()) return;
		__c = __it->second;
		__conns_QAbstractButton.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QAbstractButton__Go_new(long long id) { return new QAbstractButton__Go(id); }
void QAbstractButton_delete(void* self) { delete (QAbstractButton*)self; }

}

#include <QtWidgets/qstyleoption.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qstring.h>
#include <QtGui/qicon.h>
#include "gen_qstyleoptiontoolbox.h"

extern "C" {

void* QStyleOptionToolBox_base(void* self) { return (void*)static_cast<QStyleOption*>((QStyleOptionToolBox*)self); }
void* QStyleOptionToolBox_new() {
	return new QStyleOptionToolBox();
}
void* QStyleOptionToolBox_new2(void* other) {
	return new QStyleOptionToolBox(*(QStyleOptionToolBox*)other);
}
void* QStyleOptionToolBox_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QStyleOptionToolBox*)self)->operator=(*(QStyleOptionToolBox*)arg1);
	return (void*)&(__r);
}
char* QStyleOptionToolBox_text_get(void* self, long long* __out_len) {
	auto __r = ((QStyleOptionToolBox*)self)->text;
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QStyleOptionToolBox_text_set(void* self, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	((QStyleOptionToolBox*)self)->text = text_qstr;
}
void* QStyleOptionToolBox_icon_get(void* self) {
	auto __r = ((QStyleOptionToolBox*)self)->icon;
	return new QIcon(__r);
}
void QStyleOptionToolBox_icon_set(void* self, void* icon) {
	((QStyleOptionToolBox*)self)->icon = *(QIcon*)icon;
}
int QStyleOptionToolBox_position_get(void* self) {
	auto __r = ((QStyleOptionToolBox*)self)->position;
	return (int)(__r);
}
void QStyleOptionToolBox_position_set(void* self, int position) {
	((QStyleOptionToolBox*)self)->position = (QStyleOptionToolBox::TabPosition)(position);
}
int QStyleOptionToolBox_selectedPosition_get(void* self) {
	auto __r = ((QStyleOptionToolBox*)self)->selectedPosition;
	return (int)(__r);
}
void QStyleOptionToolBox_selectedPosition_set(void* self, int selectedPosition) {
	((QStyleOptionToolBox*)self)->selectedPosition = (QStyleOptionToolBox::SelectedPosition)(selectedPosition);
}
void QStyleOptionToolBox_delete(void* self) { delete (QStyleOptionToolBox*)self; }

}

#include <QtWidgets/qlayout.h>
#include <QtCore/qobject.h>
#include <QtWidgets/qlayoutitem.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qmargins.h>
#include <QtCore/qrect.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qsize.h>
#include <QtWidgets/qsizepolicy.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qvariant.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qlayout.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QLayout__Go : public QLayout {
public:
	long long __goId;
	unsigned char __goOn[32] = {};
	template<typename... A> QLayout__Go(long long id, A&&... a) : QLayout(std::forward<A>(a)...), __goId(id) {}
	using QLayout::widgetEvent;
	using QLayout::addChildLayout;
	using QLayout::addChildWidget;
	using QLayout::adoptLayout;
	using QLayout::alignmentRect;
	using QLayout::sender;
	using QLayout::senderSignalIndex;
	using QLayout::receivers;
	using QLayout::isSignalConnected;
	void cppgoBase_childEvent(QChildEvent* __a0) { QLayout::childEvent(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QLayout::timerEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QLayout::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QLayout::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QLayout::disconnectNotify(__a0); }
	QSize sizeHint() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 729690261457684101LL, &__h); if (__h) return *(QSize*)__r; }
		return QSize{};
	}
	void addItem(QLayoutItem* arg1) override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6884465960736393867LL, (void*)arg1, &__h); if (__h) return; }
		return;
	}
	QLayoutItem* itemAt(int index) const override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7508512214013391107LL, index, &__h); if (__h) return (QLayoutItem*)__r; }
		return nullptr;
	}
	QLayoutItem* takeAt(int index) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7549889108533486643LL, index, &__h); if (__h) return (QLayoutItem*)__r; }
		return nullptr;
	}
	int count() const override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -5627843058667511180LL, &__h); if (__h) return __r; }
		return int{};
	}
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QLayout::metaObject();
	}
	int spacing() const override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -4587105490415799682LL, &__h); if (__h) return __r; }
		return QLayout::spacing();
	}
	void setSpacing(int arg1) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 4039479028406881496LL, arg1, &__h); if (__h) return; }
		QLayout::setSpacing(arg1);
	}
	void invalidate() override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 6991341317990632368LL, &__h); if (__h) return; }
		QLayout::invalidate();
	}
	QRect geometry() const override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1050490530495513297LL, &__h); if (__h) return *(QRect*)__r; }
		return QLayout::geometry();
	}
	QFlags<Qt::Orientation> expandingDirections() const override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, 415665469015095569LL, &__h); if (__h) return QFlags<Qt::Orientation>(__r); }
		return QLayout::expandingDirections();
	}
	QSize minimumSize() const override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, -7246748916541192690LL, &__h); if (__h) return *(QSize*)__r; }
		return QLayout::minimumSize();
	}
	QSize maximumSize() const override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 4158099606088301136LL, &__h); if (__h) return *(QSize*)__r; }
		return QLayout::maximumSize();
	}
	void setGeometry(const QRect& arg1) override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5684820848726969227LL, (void*)&arg1, &__h); if (__h) return; }
		QLayout::setGeometry(arg1);
	}
	int indexOf(const QWidget* arg1) const override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, -7258473138030329570LL, (void*)arg1, &__h); if (__h) return __r; }
		return QLayout::indexOf(arg1);
	}
	int indexOf(const QLayoutItem* arg1) const override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, -3276435765994963516LL, (void*)arg1, &__h); if (__h) return __r; }
		return QLayout::indexOf(arg1);
	}
	bool isEmpty() const override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -626822430251098498LL, &__h); if (__h) return __r; }
		return QLayout::isEmpty();
	}
	QFlags<QSizePolicy::ControlType> controlTypes() const override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, 3278947275714896887LL, &__h); if (__h) return QFlags<QSizePolicy::ControlType>(__r); }
		return QLayout::controlTypes();
	}
	QLayoutItem* replaceWidget(QWidget* from, QWidget* to, QFlags<Qt::FindChildOption> options) override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_voidp_voidp_voidp_int) { void* __r = ((void*(*)(long long, long long, void*, void*, int, int*))cppgo_virt_voidp_voidp_voidp_int)(__goId, -712862742159708061LL, (void*)from, (void*)to, (int)(options), &__h); if (__h) return (QLayoutItem*)__r; }
		return QLayout::replaceWidget(from, to, options);
	}
	QLayout* layout() override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, -7183390618627992737LL, &__h); if (__h) return (QLayout*)__r; }
		return QLayout::layout();
	}
	void childEvent(QChildEvent* e) override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)e, &__h); if (__h) return; }
		QLayout::childEvent(e);
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QLayout::event(event);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QLayout::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QLayout::timerEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QLayout::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QLayout::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QLayout::disconnectNotify(signal);
	}
	bool hasHeightForWidth() const override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8266762397995423915LL, &__h); if (__h) return __r; }
		return QLayout::hasHeightForWidth();
	}
	int heightForWidth(int arg1) const override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 4311534060495737339LL, arg1, &__h); if (__h) return __r; }
		return QLayout::heightForWidth(arg1);
	}
	int minimumHeightForWidth(int arg1) const override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 5037837229528203129LL, arg1, &__h); if (__h) return __r; }
		return QLayout::minimumHeightForWidth(arg1);
	}
	QWidget* widget() const override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, -4575937485689502171LL, &__h); if (__h) return (QWidget*)__r; }
		return QLayout::widget();
	}
	QSpacerItem* spacerItem() override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 2272240669359776278LL, &__h); if (__h) return (QSpacerItem*)__r; }
		return QLayout::spacerItem();
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QLayout;
static std::mutex __connsMu_QLayout;

extern "C" {

void* QLayout_base(void* self) { return (void*)static_cast<QObject*>((QLayout*)self); }
void QLayout__set_override(void* self, int idx) { ((QLayout__Go*)self)->__goOn[idx] = 1; }
void* QLayout_fromQObject(void* o) { return (void*)qobject_cast<QLayout*>((QObject*)o); }
char* QLayout_tr(const char* s, long long* __out_len) {
	auto&& __r = QLayout::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QLayout_spacing(void* self) {
	auto&& __r = ((QLayout*)self)->spacing();
	return __r;
}
void QLayout_setSpacing(void* self, int arg1) {
	((QLayout*)self)->setSpacing(arg1);
}
void QLayout_setContentsMargins(void* self, int left, int top, int right, int bottom) {
	((QLayout*)self)->setContentsMargins(left, top, right, bottom);
}
void QLayout_setContentsMargins2(void* self, void* margins) {
	((QLayout*)self)->setContentsMargins(*(QMargins*)margins);
}
void QLayout_unsetContentsMargins(void* self) {
	((QLayout*)self)->unsetContentsMargins();
}
void QLayout_getContentsMargins(void* self, int* left, int* top, int* right, int* bottom) {
	int left_local{};
	int top_local{};
	int right_local{};
	int bottom_local{};
	((QLayout*)self)->getContentsMargins(&left_local, &top_local, &right_local, &bottom_local);
	*left = left_local;
	*top = top_local;
	*right = right_local;
	*bottom = bottom_local;
}
void* QLayout_contentsMargins(void* self) {
	auto&& __r = ((QLayout*)self)->contentsMargins();
	return new QMargins(__r);
}
void* QLayout_contentsRect(void* self) {
	auto&& __r = ((QLayout*)self)->contentsRect();
	return new QRect(__r);
}
int QLayout_setAlignment(void* self, void* w, int alignment) {
	auto&& __r = ((QLayout*)self)->setAlignment((QWidget*)w, QFlags<Qt::AlignmentFlag>(alignment));
	return __r;
}
int QLayout_setAlignment2(void* self, void* l, int alignment) {
	auto&& __r = ((QLayout*)self)->setAlignment((QLayout*)l, QFlags<Qt::AlignmentFlag>(alignment));
	return __r;
}
void QLayout_setSizeConstraint(void* self, int constraint) {
	((QLayout*)self)->setSizeConstraint((QLayout::SizeConstraint)(constraint));
}
int QLayout_sizeConstraint(void* self) {
	auto&& __r = ((QLayout*)self)->sizeConstraint();
	return (int)(__r);
}
void QLayout_setSizeConstraints(void* self, int horizontal, int vertical) {
	((QLayout*)self)->setSizeConstraints((QLayout::SizeConstraint)(horizontal), (QLayout::SizeConstraint)(vertical));
}
void QLayout_setHorizontalSizeConstraint(void* self, int constraint) {
	((QLayout*)self)->setHorizontalSizeConstraint((QLayout::SizeConstraint)(constraint));
}
int QLayout_horizontalSizeConstraint(void* self) {
	auto&& __r = ((QLayout*)self)->horizontalSizeConstraint();
	return (int)(__r);
}
void QLayout_setVerticalSizeConstraint(void* self, int constraint) {
	((QLayout*)self)->setVerticalSizeConstraint((QLayout::SizeConstraint)(constraint));
}
int QLayout_verticalSizeConstraint(void* self) {
	auto&& __r = ((QLayout*)self)->verticalSizeConstraint();
	return (int)(__r);
}
void QLayout_setMenuBar(void* self, void* w) {
	((QLayout*)self)->setMenuBar((QWidget*)w);
}
void* QLayout_menuBar(void* self) {
	auto&& __r = ((QLayout*)self)->menuBar();
	return (void*)(__r);
}
void* QLayout_parentWidget(void* self) {
	auto&& __r = ((QLayout*)self)->parentWidget();
	return (void*)(__r);
}
void QLayout_invalidate(void* self) {
	((QLayout*)self)->invalidate();
}
void* QLayout_geometry(void* self) {
	auto&& __r = ((QLayout*)self)->geometry();
	return new QRect(__r);
}
int QLayout_activate(void* self) {
	auto&& __r = ((QLayout*)self)->activate();
	return __r;
}
void QLayout_update(void* self) {
	((QLayout*)self)->update();
}
void QLayout_addWidget(void* self, void* w) {
	((QLayout*)self)->addWidget((QWidget*)w);
}
void QLayout_addItem(void* self, void* arg1) {
	((QLayout*)self)->addItem((QLayoutItem*)arg1);
}
void QLayout_removeWidget(void* self, void* w) {
	((QLayout*)self)->removeWidget((QWidget*)w);
}
void QLayout_removeItem(void* self, void* arg1) {
	((QLayout*)self)->removeItem((QLayoutItem*)arg1);
}
int QLayout_expandingDirections(void* self) {
	auto&& __r = ((QLayout*)self)->expandingDirections();
	return (int)(__r);
}
void* QLayout_minimumSize(void* self) {
	auto&& __r = ((QLayout*)self)->minimumSize();
	return new QSize(__r);
}
void* QLayout_maximumSize(void* self) {
	auto&& __r = ((QLayout*)self)->maximumSize();
	return new QSize(__r);
}
void QLayout_setGeometry(void* self, void* arg1) {
	((QLayout*)self)->setGeometry(*(QRect*)arg1);
}
void* QLayout_itemAt(void* self, int index) {
	auto&& __r = ((QLayout*)self)->itemAt(index);
	return (void*)(__r);
}
void* QLayout_takeAt(void* self, int index) {
	auto&& __r = ((QLayout*)self)->takeAt(index);
	return (void*)(__r);
}
int QLayout_indexOf(void* self, void* arg1) {
	auto&& __r = ((QLayout*)self)->indexOf((QWidget*)arg1);
	return __r;
}
int QLayout_indexOf2(void* self, void* arg1) {
	auto&& __r = ((QLayout*)self)->indexOf((QLayoutItem*)arg1);
	return __r;
}
int QLayout_count(void* self) {
	auto&& __r = ((QLayout*)self)->count();
	return __r;
}
int QLayout_isEmpty(void* self) {
	auto&& __r = ((QLayout*)self)->isEmpty();
	return __r;
}
int QLayout_controlTypes(void* self) {
	auto&& __r = ((QLayout*)self)->controlTypes();
	return (int)(__r);
}
void* QLayout_replaceWidget(void* self, void* from, void* to, int options) {
	auto&& __r = ((QLayout*)self)->replaceWidget((QWidget*)from, (QWidget*)to, QFlags<Qt::FindChildOption>(options));
	return (void*)(__r);
}
int QLayout_totalMinimumHeightForWidth(void* self, int w) {
	auto&& __r = ((QLayout*)self)->totalMinimumHeightForWidth(w);
	return __r;
}
int QLayout_totalHeightForWidth(void* self, int w) {
	auto&& __r = ((QLayout*)self)->totalHeightForWidth(w);
	return __r;
}
void* QLayout_totalMinimumSize(void* self) {
	auto&& __r = ((QLayout*)self)->totalMinimumSize();
	return new QSize(__r);
}
void* QLayout_totalMaximumSize(void* self) {
	auto&& __r = ((QLayout*)self)->totalMaximumSize();
	return new QSize(__r);
}
void* QLayout_totalSizeHint(void* self) {
	auto&& __r = ((QLayout*)self)->totalSizeHint();
	return new QSize(__r);
}
void* QLayout_layout(void* self) {
	auto&& __r = ((QLayout*)self)->layout();
	return (void*)(__r);
}
void QLayout_setEnabled(void* self, int arg1) {
	((QLayout*)self)->setEnabled(arg1);
}
int QLayout_isEnabled(void* self) {
	auto&& __r = ((QLayout*)self)->isEnabled();
	return __r;
}
void* QLayout_closestAcceptableSize(void* w, void* s) {
	auto&& __r = QLayout::closestAcceptableSize((QWidget*)w, *(QSize*)s);
	return new QSize(__r);
}
void QLayout_widgetEvent(void* self, void* arg1) {
	((QLayout__Go*)self)->widgetEvent((QEvent*)arg1);
}
void QLayout_addChildLayout(void* self, void* l) {
	((QLayout__Go*)self)->addChildLayout((QLayout*)l);
}
void QLayout_addChildWidget(void* self, void* w) {
	((QLayout__Go*)self)->addChildWidget((QWidget*)w);
}
int QLayout_adoptLayout(void* self, void* layout) {
	auto&& __r = ((QLayout__Go*)self)->adoptLayout((QLayout*)layout);
	return __r;
}
void* QLayout_alignmentRect(void* self, void* arg1) {
	auto&& __r = ((QLayout__Go*)self)->alignmentRect(*(QRect*)arg1);
	return new QRect(__r);
}
void* QLayout_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QLayout::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QLayout_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QLayout::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QLayout_sender(void* self) {
	auto&& __r = ((QLayout__Go*)self)->sender();
	return (void*)(__r);
}
int QLayout_senderSignalIndex(void* self) {
	auto&& __r = ((QLayout__Go*)self)->senderSignalIndex();
	return __r;
}
int QLayout_receivers(void* self, const char* signal) {
	auto&& __r = ((QLayout__Go*)self)->receivers(signal);
	return __r;
}
int QLayout_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QLayout__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
void* QLayout_sizeHint(void* self) {
	auto&& __r = ((QLayout*)self)->sizeHint();
	return new QSize(__r);
}
int QLayout_hasHeightForWidth(void* self) {
	auto&& __r = ((QLayout*)self)->hasHeightForWidth();
	return __r;
}
int QLayout_heightForWidth(void* self, int arg1) {
	auto&& __r = ((QLayout*)self)->heightForWidth(arg1);
	return __r;
}
int QLayout_minimumHeightForWidth(void* self, int arg1) {
	auto&& __r = ((QLayout*)self)->minimumHeightForWidth(arg1);
	return __r;
}
void* QLayout_widget(void* self) {
	auto&& __r = ((QLayout*)self)->widget();
	return (void*)(__r);
}
void* QLayout_spacerItem(void* self) {
	auto&& __r = ((QLayout*)self)->spacerItem();
	return (void*)(__r);
}
int QLayout_alignment(void* self) {
	auto&& __r = ((QLayout*)self)->alignment();
	return (int)(__r);
}
char* QLayout_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QLayout::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLayout_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QLayout::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QLayout_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QLayout::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QLayout_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QLayout::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QLayout_asQLayoutItem(void* self) {
	auto&& __r = static_cast<QLayoutItem*>((QLayout*)self);
	return (void*)(__r);
}
int QLayout_spacing_default(void* self) {
	auto&& __r = ((QLayout*)self)->QLayout::spacing();
	return __r;
}
void QLayout_setSpacing_default(void* self, int arg1) {
	((QLayout*)self)->QLayout::setSpacing(arg1);
}
void QLayout_invalidate_default(void* self) {
	((QLayout*)self)->QLayout::invalidate();
}
void* QLayout_geometry_default(void* self) {
	auto&& __r = ((QLayout*)self)->QLayout::geometry();
	return new QRect(__r);
}
int QLayout_expandingDirections_default(void* self) {
	auto&& __r = ((QLayout*)self)->QLayout::expandingDirections();
	return (int)(__r);
}
void* QLayout_minimumSize_default(void* self) {
	auto&& __r = ((QLayout*)self)->QLayout::minimumSize();
	return new QSize(__r);
}
void* QLayout_maximumSize_default(void* self) {
	auto&& __r = ((QLayout*)self)->QLayout::maximumSize();
	return new QSize(__r);
}
void QLayout_setGeometry_default(void* self, void* arg1) {
	((QLayout*)self)->QLayout::setGeometry(*(QRect*)arg1);
}
int QLayout_indexOf_default(void* self, void* arg1) {
	auto&& __r = ((QLayout*)self)->QLayout::indexOf((QWidget*)arg1);
	return __r;
}
int QLayout_indexOf_default2(void* self, void* arg1) {
	auto&& __r = ((QLayout*)self)->QLayout::indexOf((QLayoutItem*)arg1);
	return __r;
}
int QLayout_isEmpty_default(void* self) {
	auto&& __r = ((QLayout*)self)->QLayout::isEmpty();
	return __r;
}
int QLayout_controlTypes_default(void* self) {
	auto&& __r = ((QLayout*)self)->QLayout::controlTypes();
	return (int)(__r);
}
void* QLayout_replaceWidget_default(void* self, void* from, void* to, int options) {
	auto&& __r = ((QLayout*)self)->QLayout::replaceWidget((QWidget*)from, (QWidget*)to, QFlags<Qt::FindChildOption>(options));
	return (void*)(__r);
}
void* QLayout_layout_default(void* self) {
	auto&& __r = ((QLayout*)self)->QLayout::layout();
	return (void*)(__r);
}
int QLayout_hasHeightForWidth_default(void* self) {
	auto&& __r = ((QLayout*)self)->QLayout::hasHeightForWidth();
	return __r;
}
int QLayout_heightForWidth_default(void* self, int arg1) {
	auto&& __r = ((QLayout*)self)->QLayout::heightForWidth(arg1);
	return __r;
}
int QLayout_minimumHeightForWidth_default(void* self, int arg1) {
	auto&& __r = ((QLayout*)self)->QLayout::minimumHeightForWidth(arg1);
	return __r;
}
void* QLayout_widget_default(void* self) {
	auto&& __r = ((QLayout*)self)->QLayout::widget();
	return (void*)(__r);
}
void* QLayout_spacerItem_default(void* self) {
	auto&& __r = ((QLayout*)self)->QLayout::spacerItem();
	return (void*)(__r);
}
void* QLayout_metaObject_default(void* self) {
	auto&& __r = ((QLayout*)self)->QLayout::metaObject();
	return (void*)(__r);
}
void QLayout_childEvent_default(void* self, void* e) {
	((QLayout__Go*)self)->cppgoBase_childEvent((QChildEvent*)e);
}
int QLayout_event_default(void* self, void* event) {
	auto&& __r = ((QLayout*)self)->QLayout::event((QEvent*)event);
	return __r;
}
int QLayout_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QLayout*)self)->QLayout::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QLayout_timerEvent_default(void* self, void* event) {
	((QLayout__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QLayout_customEvent_default(void* self, void* event) {
	((QLayout__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QLayout_connectNotify_default(void* self, void* signal) {
	((QLayout__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QLayout_disconnectNotify_default(void* self, void* signal) {
	((QLayout__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QLayout_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QLayout*)self, &QLayout::destroyed, (QLayout*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QLayout);
	__conns_QLayout[__id] = __c;
}
void QLayout_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QLayout*)self, &QLayout::destroyed, (QLayout*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QLayout);
	__conns_QLayout[__id] = __c;
}
void QLayout_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QLayout*)self, &QLayout::objectNameChanged, (QLayout*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QLayout);
	__conns_QLayout[__id] = __c;
}
void QLayout__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QLayout);
		auto __it = __conns_QLayout.find(id);
		if (__it == __conns_QLayout.end()) return;
		__c = __it->second;
		__conns_QLayout.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QLayout__Go_new(long long id) { return new QLayout__Go(id); }
void QLayout_delete(void* self) { delete (QLayout*)self; }

}

#include <QtWidgets/qdrawutil.h>
#include <QtCore/qnamespace.h>
#include "gen_qtilerules.h"

extern "C" {

void* QTileRules_new(int horizontalRule, int verticalRule) {
	return new QTileRules((Qt::TileRule)(horizontalRule), (Qt::TileRule)(verticalRule));
}
void* QTileRules_new2() {
	return new QTileRules();
}
void* QTileRules_new3(void* arg1) {
	return new QTileRules(*(QTileRules*)arg1);
}
void* QTileRules_new4(int rule) {
	return new QTileRules((Qt::TileRule)(rule));
}
int QTileRules_horizontal_get(void* self) {
	auto __r = ((QTileRules*)self)->horizontal;
	return (int)(__r);
}
void QTileRules_horizontal_set(void* self, int horizontal) {
	((QTileRules*)self)->horizontal = (Qt::TileRule)(horizontal);
}
int QTileRules_vertical_get(void* self) {
	auto __r = ((QTileRules*)self)->vertical;
	return (int)(__r);
}
void QTileRules_vertical_set(void* self, int vertical) {
	((QTileRules*)self)->vertical = (Qt::TileRule)(vertical);
}
void QTileRules_delete(void* self) { delete (QTileRules*)self; }

}

#include <QtWidgets/qgraphicsitem.h>
#include <QtCore/qline.h>
#include <QtGui/qpen.h>
#include <QtCore/qrect.h>
#include <QtGui/qpainterpath.h>
#include <QtCore/qpoint.h>
#include <QtGui/qpainter.h>
#include <QtWidgets/qstyleoption.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qvariant.h>
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
#include "gen_qgraphicslineitem.h"
#include "gen_purego_cb.h"

class QGraphicsLineItem__Go : public QGraphicsLineItem {
public:
	long long __goId;
	unsigned char __goOn[35] = {};
	template<typename... A> QGraphicsLineItem__Go(long long id, A&&... a) : QGraphicsLineItem(std::forward<A>(a)...), __goId(id) {}
	using QGraphicsLineItem::updateMicroFocus;
	using QGraphicsLineItem::addToIndex;
	using QGraphicsLineItem::removeFromIndex;
	using QGraphicsLineItem::prepareGeometryChange;
	bool cppgoBase_supportsExtension(int __a0) { return QGraphicsLineItem::supportsExtension((QGraphicsItem::Extension)__a0); }
	void cppgoBase_setExtension(int __a0, const QVariant& __a1) { QGraphicsLineItem::setExtension((QGraphicsItem::Extension)__a0, __a1); }
	QVariant cppgoBase_extension(const QVariant& __a0) { return QGraphicsLineItem::extension(__a0); }
	bool cppgoBase_sceneEventFilter(QGraphicsItem* __a0, QEvent* __a1) { return QGraphicsLineItem::sceneEventFilter(__a0, __a1); }
	bool cppgoBase_sceneEvent(QEvent* __a0) { return QGraphicsLineItem::sceneEvent(__a0); }
	void cppgoBase_contextMenuEvent(QGraphicsSceneContextMenuEvent* __a0) { QGraphicsLineItem::contextMenuEvent(__a0); }
	void cppgoBase_dragEnterEvent(QGraphicsSceneDragDropEvent* __a0) { QGraphicsLineItem::dragEnterEvent(__a0); }
	void cppgoBase_dragLeaveEvent(QGraphicsSceneDragDropEvent* __a0) { QGraphicsLineItem::dragLeaveEvent(__a0); }
	void cppgoBase_dragMoveEvent(QGraphicsSceneDragDropEvent* __a0) { QGraphicsLineItem::dragMoveEvent(__a0); }
	void cppgoBase_dropEvent(QGraphicsSceneDragDropEvent* __a0) { QGraphicsLineItem::dropEvent(__a0); }
	void cppgoBase_focusInEvent(QFocusEvent* __a0) { QGraphicsLineItem::focusInEvent(__a0); }
	void cppgoBase_focusOutEvent(QFocusEvent* __a0) { QGraphicsLineItem::focusOutEvent(__a0); }
	void cppgoBase_hoverEnterEvent(QGraphicsSceneHoverEvent* __a0) { QGraphicsLineItem::hoverEnterEvent(__a0); }
	void cppgoBase_hoverMoveEvent(QGraphicsSceneHoverEvent* __a0) { QGraphicsLineItem::hoverMoveEvent(__a0); }
	void cppgoBase_hoverLeaveEvent(QGraphicsSceneHoverEvent* __a0) { QGraphicsLineItem::hoverLeaveEvent(__a0); }
	void cppgoBase_keyPressEvent(QKeyEvent* __a0) { QGraphicsLineItem::keyPressEvent(__a0); }
	void cppgoBase_keyReleaseEvent(QKeyEvent* __a0) { QGraphicsLineItem::keyReleaseEvent(__a0); }
	void cppgoBase_mousePressEvent(QGraphicsSceneMouseEvent* __a0) { QGraphicsLineItem::mousePressEvent(__a0); }
	void cppgoBase_mouseMoveEvent(QGraphicsSceneMouseEvent* __a0) { QGraphicsLineItem::mouseMoveEvent(__a0); }
	void cppgoBase_mouseReleaseEvent(QGraphicsSceneMouseEvent* __a0) { QGraphicsLineItem::mouseReleaseEvent(__a0); }
	void cppgoBase_mouseDoubleClickEvent(QGraphicsSceneMouseEvent* __a0) { QGraphicsLineItem::mouseDoubleClickEvent(__a0); }
	void cppgoBase_wheelEvent(QGraphicsSceneWheelEvent* __a0) { QGraphicsLineItem::wheelEvent(__a0); }
	void cppgoBase_inputMethodEvent(QInputMethodEvent* __a0) { QGraphicsLineItem::inputMethodEvent(__a0); }
	QVariant cppgoBase_inputMethodQuery(int __a0) { return QGraphicsLineItem::inputMethodQuery((Qt::InputMethodQuery)__a0); }
	QVariant cppgoBase_itemChange(int __a0, const QVariant& __a1) { return QGraphicsLineItem::itemChange((QGraphicsItem::GraphicsItemChange)__a0, __a1); }
	QRectF boundingRect() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 2510008782785038831LL, &__h); if (__h) return *(QRectF*)__r; }
		return QGraphicsLineItem::boundingRect();
	}
	QPainterPath shape() const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 7938019069247444294LL, &__h); if (__h) return *(QPainterPath*)__r; }
		return QGraphicsLineItem::shape();
	}
	bool contains(const QPointF& point) const override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 5645366511944179768LL, (void*)&point, &__h); if (__h) return __r; }
		return QGraphicsLineItem::contains(point);
	}
	void paint(QPainter* painter, const QStyleOptionGraphicsItem* option, QWidget* widget) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_voidp_voidp_voidp) { ((void(*)(long long, long long, void*, void*, void*, int*))cppgo_virt_void_voidp_voidp_voidp)(__goId, 6083584261053556819LL, (void*)painter, (void*)option, (void*)widget, &__h); if (__h) return; }
		QGraphicsLineItem::paint(painter, option, widget);
	}
	bool isObscuredBy(const QGraphicsItem* item) const override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, -7215030271914903671LL, (void*)item, &__h); if (__h) return __r; }
		return QGraphicsLineItem::isObscuredBy(item);
	}
	QPainterPath opaqueArea() const override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 399370970023989435LL, &__h); if (__h) return *(QPainterPath*)__r; }
		return QGraphicsLineItem::opaqueArea();
	}
	int type() const override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -6371403872082420691LL, &__h); if (__h) return __r; }
		return QGraphicsLineItem::type();
	}
	bool supportsExtension(QGraphicsItem::Extension extension) const override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, -5729166443651306418LL, (int)(extension), &__h); if (__h) return __r; }
		return QGraphicsLineItem::supportsExtension(extension);
	}
	void setExtension(QGraphicsItem::Extension extension, const QVariant& variant) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void_int_voidp) { ((void(*)(long long, long long, int, void*, int*))cppgo_virt_void_int_voidp)(__goId, -6730181576106171580LL, (int)(extension), (void*)&variant, &__h); if (__h) return; }
		QGraphicsLineItem::setExtension(extension, variant);
	}
	QVariant extension(const QVariant& variant) const override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 2270616358835251226LL, (void*)&variant, &__h); if (__h) return *(QVariant*)__r; }
		return QGraphicsLineItem::extension(variant);
	}
	void advance(int phase) override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, -4830836341480073475LL, phase, &__h); if (__h) return; }
		QGraphicsLineItem::advance(phase);
	}
	bool collidesWithItem(const QGraphicsItem* other, Qt::ItemSelectionMode mode) const override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_int_voidp_int) { int __r = ((int(*)(long long, long long, void*, int, int*))cppgo_virt_int_voidp_int)(__goId, 6469736184362593173LL, (void*)other, (int)(mode), &__h); if (__h) return __r; }
		return QGraphicsLineItem::collidesWithItem(other, mode);
	}
	bool collidesWithPath(const QPainterPath& path, Qt::ItemSelectionMode mode) const override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_int_voidp_int) { int __r = ((int(*)(long long, long long, void*, int, int*))cppgo_virt_int_voidp_int)(__goId, 6636575182718753853LL, (void*)&path, (int)(mode), &__h); if (__h) return __r; }
		return QGraphicsLineItem::collidesWithPath(path, mode);
	}
	bool sceneEventFilter(QGraphicsItem* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, -4584203218241215725LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QGraphicsLineItem::sceneEventFilter(watched, event);
	}
	bool sceneEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, -2846371507838587071LL, (void*)event, &__h); if (__h) return __r; }
		return QGraphicsLineItem::sceneEvent(event);
	}
	void contextMenuEvent(QGraphicsSceneContextMenuEvent* event) override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6698043222227006015LL, (void*)event, &__h); if (__h) return; }
		QGraphicsLineItem::contextMenuEvent(event);
	}
	void dragEnterEvent(QGraphicsSceneDragDropEvent* event) override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 206122715919917589LL, (void*)event, &__h); if (__h) return; }
		QGraphicsLineItem::dragEnterEvent(event);
	}
	void dragLeaveEvent(QGraphicsSceneDragDropEvent* event) override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -3139904158373725006LL, (void*)event, &__h); if (__h) return; }
		QGraphicsLineItem::dragLeaveEvent(event);
	}
	void dragMoveEvent(QGraphicsSceneDragDropEvent* event) override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1781545137352745978LL, (void*)event, &__h); if (__h) return; }
		QGraphicsLineItem::dragMoveEvent(event);
	}
	void dropEvent(QGraphicsSceneDragDropEvent* event) override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 9107058682754591716LL, (void*)event, &__h); if (__h) return; }
		QGraphicsLineItem::dropEvent(event);
	}
	void focusInEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 975043767033552552LL, (void*)event, &__h); if (__h) return; }
		QGraphicsLineItem::focusInEvent(event);
	}
	void focusOutEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -385596010576048993LL, (void*)event, &__h); if (__h) return; }
		QGraphicsLineItem::focusOutEvent(event);
	}
	void hoverEnterEvent(QGraphicsSceneHoverEvent* event) override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6932941943906681679LL, (void*)event, &__h); if (__h) return; }
		QGraphicsLineItem::hoverEnterEvent(event);
	}
	void hoverMoveEvent(QGraphicsSceneHoverEvent* event) override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -2745191522144316396LL, (void*)event, &__h); if (__h) return; }
		QGraphicsLineItem::hoverMoveEvent(event);
	}
	void hoverLeaveEvent(QGraphicsSceneHoverEvent* event) override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7944408482203603984LL, (void*)event, &__h); if (__h) return; }
		QGraphicsLineItem::hoverLeaveEvent(event);
	}
	void keyPressEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1023791977679793185LL, (void*)event, &__h); if (__h) return; }
		QGraphicsLineItem::keyPressEvent(event);
	}
	void keyReleaseEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 772598475950050299LL, (void*)event, &__h); if (__h) return; }
		QGraphicsLineItem::keyReleaseEvent(event);
	}
	void mousePressEvent(QGraphicsSceneMouseEvent* event) override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -332090898078543319LL, (void*)event, &__h); if (__h) return; }
		QGraphicsLineItem::mousePressEvent(event);
	}
	void mouseMoveEvent(QGraphicsSceneMouseEvent* event) override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4255494004099806581LL, (void*)event, &__h); if (__h) return; }
		QGraphicsLineItem::mouseMoveEvent(event);
	}
	void mouseReleaseEvent(QGraphicsSceneMouseEvent* event) override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4415920049297795213LL, (void*)event, &__h); if (__h) return; }
		QGraphicsLineItem::mouseReleaseEvent(event);
	}
	void mouseDoubleClickEvent(QGraphicsSceneMouseEvent* event) override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4029073905697023887LL, (void*)event, &__h); if (__h) return; }
		QGraphicsLineItem::mouseDoubleClickEvent(event);
	}
	void wheelEvent(QGraphicsSceneWheelEvent* event) override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1919988340934511316LL, (void*)event, &__h); if (__h) return; }
		QGraphicsLineItem::wheelEvent(event);
	}
	void inputMethodEvent(QInputMethodEvent* event) override {
		int __h = 0;
		if (__goOn[32] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5438819661942843222LL, (void*)event, &__h); if (__h) return; }
		QGraphicsLineItem::inputMethodEvent(event);
	}
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const override {
		int __h = 0;
		if (__goOn[33] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7345853097340709638LL, (int)(query), &__h); if (__h) return *(QVariant*)__r; }
		return QGraphicsLineItem::inputMethodQuery(query);
	}
	QVariant itemChange(QGraphicsItem::GraphicsItemChange change, const QVariant& value) override {
		int __h = 0;
		if (__goOn[34] && cppgo_virt_voidp_int_voidp) { void* __r = ((void*(*)(long long, long long, int, void*, int*))cppgo_virt_voidp_int_voidp)(__goId, -1888572168794878830LL, (int)(change), (void*)&value, &__h); if (__h) return *(QVariant*)__r; }
		return QGraphicsLineItem::itemChange(change, value);
	}
};

extern "C" {

void* QGraphicsLineItem_base(void* self) { return (void*)static_cast<QGraphicsItem*>((QGraphicsLineItem*)self); }
void QGraphicsLineItem__set_override(void* self, int idx) { ((QGraphicsLineItem__Go*)self)->__goOn[idx] = 1; }
void* QGraphicsLineItem_new(long long __goid) {
	return new QGraphicsLineItem__Go(__goid);
}
void* QGraphicsLineItem_new2(long long __goid, void* line) {
	return new QGraphicsLineItem__Go(__goid, *(QLineF*)line);
}
void* QGraphicsLineItem_new3(long long __goid, double x1, double y1, double x2, double y2) {
	return new QGraphicsLineItem__Go(__goid, x1, y1, x2, y2);
}
void* QGraphicsLineItem_new4(long long __goid, void* parent) {
	return new QGraphicsLineItem__Go(__goid, (QGraphicsItem*)parent);
}
void* QGraphicsLineItem_new5(long long __goid, void* line, void* parent) {
	return new QGraphicsLineItem__Go(__goid, *(QLineF*)line, (QGraphicsItem*)parent);
}
void* QGraphicsLineItem_new6(long long __goid, double x1, double y1, double x2, double y2, void* parent) {
	return new QGraphicsLineItem__Go(__goid, x1, y1, x2, y2, (QGraphicsItem*)parent);
}
void* QGraphicsLineItem_pen(void* self) {
	auto&& __r = ((QGraphicsLineItem*)self)->pen();
	return new QPen(__r);
}
void QGraphicsLineItem_setPen(void* self, void* pen) {
	((QGraphicsLineItem*)self)->setPen(*(QPen*)pen);
}
void* QGraphicsLineItem_line(void* self) {
	auto&& __r = ((QGraphicsLineItem*)self)->line();
	return new QLineF(__r);
}
void QGraphicsLineItem_setLine(void* self, void* line) {
	((QGraphicsLineItem*)self)->setLine(*(QLineF*)line);
}
void QGraphicsLineItem_setLine2(void* self, double x1, double y1, double x2, double y2) {
	((QGraphicsLineItem*)self)->setLine(x1, y1, x2, y2);
}
void QGraphicsLineItem_updateMicroFocus(void* self) {
	((QGraphicsLineItem__Go*)self)->updateMicroFocus();
}
void QGraphicsLineItem_addToIndex(void* self) {
	((QGraphicsLineItem__Go*)self)->addToIndex();
}
void QGraphicsLineItem_removeFromIndex(void* self) {
	((QGraphicsLineItem__Go*)self)->removeFromIndex();
}
void QGraphicsLineItem_prepareGeometryChange(void* self) {
	((QGraphicsLineItem__Go*)self)->prepareGeometryChange();
}
void* QGraphicsLineItem_boundingRect_default(void* self) {
	auto&& __r = ((QGraphicsLineItem*)self)->QGraphicsLineItem::boundingRect();
	return new QRectF(__r);
}
void* QGraphicsLineItem_shape_default(void* self) {
	auto&& __r = ((QGraphicsLineItem*)self)->QGraphicsLineItem::shape();
	return new QPainterPath(__r);
}
int QGraphicsLineItem_contains_default(void* self, void* point) {
	auto&& __r = ((QGraphicsLineItem*)self)->QGraphicsLineItem::contains(*(QPointF*)point);
	return __r;
}
void QGraphicsLineItem_paint_default(void* self, void* painter, void* option, void* widget) {
	((QGraphicsLineItem*)self)->QGraphicsLineItem::paint((QPainter*)painter, (QStyleOptionGraphicsItem*)option, (QWidget*)widget);
}
int QGraphicsLineItem_isObscuredBy_default(void* self, void* item) {
	auto&& __r = ((QGraphicsLineItem*)self)->QGraphicsLineItem::isObscuredBy((QGraphicsItem*)item);
	return __r;
}
void* QGraphicsLineItem_opaqueArea_default(void* self) {
	auto&& __r = ((QGraphicsLineItem*)self)->QGraphicsLineItem::opaqueArea();
	return new QPainterPath(__r);
}
int QGraphicsLineItem_type_default(void* self) {
	auto&& __r = ((QGraphicsLineItem*)self)->QGraphicsLineItem::type();
	return __r;
}
int QGraphicsLineItem_supportsExtension_default(void* self, int extension) {
	auto&& __r = ((QGraphicsLineItem__Go*)self)->cppgoBase_supportsExtension(extension);
	return __r;
}
void QGraphicsLineItem_setExtension_default(void* self, int extension, void* variant) {
	((QGraphicsLineItem__Go*)self)->cppgoBase_setExtension(extension, *(QVariant*)variant);
}
void* QGraphicsLineItem_extension_default(void* self, void* variant) {
	auto&& __r = ((QGraphicsLineItem__Go*)self)->cppgoBase_extension(*(QVariant*)variant);
	return new QVariant(__r);
}
void QGraphicsLineItem_advance_default(void* self, int phase) {
	((QGraphicsLineItem*)self)->QGraphicsLineItem::advance(phase);
}
int QGraphicsLineItem_collidesWithItem_default(void* self, void* other, int mode) {
	auto&& __r = ((QGraphicsLineItem*)self)->QGraphicsLineItem::collidesWithItem((QGraphicsItem*)other, (Qt::ItemSelectionMode)(mode));
	return __r;
}
int QGraphicsLineItem_collidesWithPath_default(void* self, void* path, int mode) {
	auto&& __r = ((QGraphicsLineItem*)self)->QGraphicsLineItem::collidesWithPath(*(QPainterPath*)path, (Qt::ItemSelectionMode)(mode));
	return __r;
}
int QGraphicsLineItem_sceneEventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QGraphicsLineItem__Go*)self)->cppgoBase_sceneEventFilter((QGraphicsItem*)watched, (QEvent*)event);
	return __r;
}
int QGraphicsLineItem_sceneEvent_default(void* self, void* event) {
	auto&& __r = ((QGraphicsLineItem__Go*)self)->cppgoBase_sceneEvent((QEvent*)event);
	return __r;
}
void QGraphicsLineItem_contextMenuEvent_default(void* self, void* event) {
	((QGraphicsLineItem__Go*)self)->cppgoBase_contextMenuEvent((QGraphicsSceneContextMenuEvent*)event);
}
void QGraphicsLineItem_dragEnterEvent_default(void* self, void* event) {
	((QGraphicsLineItem__Go*)self)->cppgoBase_dragEnterEvent((QGraphicsSceneDragDropEvent*)event);
}
void QGraphicsLineItem_dragLeaveEvent_default(void* self, void* event) {
	((QGraphicsLineItem__Go*)self)->cppgoBase_dragLeaveEvent((QGraphicsSceneDragDropEvent*)event);
}
void QGraphicsLineItem_dragMoveEvent_default(void* self, void* event) {
	((QGraphicsLineItem__Go*)self)->cppgoBase_dragMoveEvent((QGraphicsSceneDragDropEvent*)event);
}
void QGraphicsLineItem_dropEvent_default(void* self, void* event) {
	((QGraphicsLineItem__Go*)self)->cppgoBase_dropEvent((QGraphicsSceneDragDropEvent*)event);
}
void QGraphicsLineItem_focusInEvent_default(void* self, void* event) {
	((QGraphicsLineItem__Go*)self)->cppgoBase_focusInEvent((QFocusEvent*)event);
}
void QGraphicsLineItem_focusOutEvent_default(void* self, void* event) {
	((QGraphicsLineItem__Go*)self)->cppgoBase_focusOutEvent((QFocusEvent*)event);
}
void QGraphicsLineItem_hoverEnterEvent_default(void* self, void* event) {
	((QGraphicsLineItem__Go*)self)->cppgoBase_hoverEnterEvent((QGraphicsSceneHoverEvent*)event);
}
void QGraphicsLineItem_hoverMoveEvent_default(void* self, void* event) {
	((QGraphicsLineItem__Go*)self)->cppgoBase_hoverMoveEvent((QGraphicsSceneHoverEvent*)event);
}
void QGraphicsLineItem_hoverLeaveEvent_default(void* self, void* event) {
	((QGraphicsLineItem__Go*)self)->cppgoBase_hoverLeaveEvent((QGraphicsSceneHoverEvent*)event);
}
void QGraphicsLineItem_keyPressEvent_default(void* self, void* event) {
	((QGraphicsLineItem__Go*)self)->cppgoBase_keyPressEvent((QKeyEvent*)event);
}
void QGraphicsLineItem_keyReleaseEvent_default(void* self, void* event) {
	((QGraphicsLineItem__Go*)self)->cppgoBase_keyReleaseEvent((QKeyEvent*)event);
}
void QGraphicsLineItem_mousePressEvent_default(void* self, void* event) {
	((QGraphicsLineItem__Go*)self)->cppgoBase_mousePressEvent((QGraphicsSceneMouseEvent*)event);
}
void QGraphicsLineItem_mouseMoveEvent_default(void* self, void* event) {
	((QGraphicsLineItem__Go*)self)->cppgoBase_mouseMoveEvent((QGraphicsSceneMouseEvent*)event);
}
void QGraphicsLineItem_mouseReleaseEvent_default(void* self, void* event) {
	((QGraphicsLineItem__Go*)self)->cppgoBase_mouseReleaseEvent((QGraphicsSceneMouseEvent*)event);
}
void QGraphicsLineItem_mouseDoubleClickEvent_default(void* self, void* event) {
	((QGraphicsLineItem__Go*)self)->cppgoBase_mouseDoubleClickEvent((QGraphicsSceneMouseEvent*)event);
}
void QGraphicsLineItem_wheelEvent_default(void* self, void* event) {
	((QGraphicsLineItem__Go*)self)->cppgoBase_wheelEvent((QGraphicsSceneWheelEvent*)event);
}
void QGraphicsLineItem_inputMethodEvent_default(void* self, void* event) {
	((QGraphicsLineItem__Go*)self)->cppgoBase_inputMethodEvent((QInputMethodEvent*)event);
}
void* QGraphicsLineItem_inputMethodQuery_default(void* self, int query) {
	auto&& __r = ((QGraphicsLineItem__Go*)self)->cppgoBase_inputMethodQuery(query);
	return new QVariant(__r);
}
void* QGraphicsLineItem_itemChange_default(void* self, int change, void* value) {
	auto&& __r = ((QGraphicsLineItem__Go*)self)->cppgoBase_itemChange(change, *(QVariant*)value);
	return new QVariant(__r);
}
void* QGraphicsLineItem__Go_new(long long id) { return new QGraphicsLineItem__Go(id); }
void QGraphicsLineItem_delete(void* self) { delete (QGraphicsLineItem*)self; }

}

#include <QtWidgets/qgraphicssceneevent.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qpoint.h>
#include <QtCore/qnamespace.h>
#include <QtWidgets/qwidget.h>
#include "gen_qgraphicsscenehoverevent.h"
#include "gen_purego_cb.h"

class QGraphicsSceneHoverEvent__Go : public QGraphicsSceneHoverEvent {
public:
	long long __goId;
	unsigned char __goOn[2] = {};
	template<typename... A> QGraphicsSceneHoverEvent__Go(long long id, A&&... a) : QGraphicsSceneHoverEvent(std::forward<A>(a)...), __goId(id) {}
	void setAccepted(bool accepted) override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 6367896784077321574LL, accepted, &__h); if (__h) return; }
		QGraphicsSceneHoverEvent::setAccepted(accepted);
	}
	QEvent* clone() const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, -7032813281048616478LL, &__h); if (__h) return (QEvent*)__r; }
		return QGraphicsSceneHoverEvent::clone();
	}
};

extern "C" {

void* QGraphicsSceneHoverEvent_base(void* self) { return (void*)static_cast<QGraphicsSceneEvent*>((QGraphicsSceneHoverEvent*)self); }
void QGraphicsSceneHoverEvent_chain(void* self, void** out) {
	QGraphicsSceneHoverEvent* __p = (QGraphicsSceneHoverEvent*)self;
	QGraphicsSceneEvent* __b0 = static_cast<QGraphicsSceneEvent*>(__p); out[0] = (void*)__b0;
	QEvent* __b1 = static_cast<QEvent*>(__b0); out[1] = (void*)__b1;
}
void QGraphicsSceneHoverEvent__set_override(void* self, int idx) { ((QGraphicsSceneHoverEvent__Go*)self)->__goOn[idx] = 1; }
void* QGraphicsSceneHoverEvent_new(long long __goid) {
	return new QGraphicsSceneHoverEvent__Go(__goid);
}
void* QGraphicsSceneHoverEvent_new2(long long __goid, int type_) {
	return new QGraphicsSceneHoverEvent__Go(__goid, (QEvent::Type)(type_));
}
void* QGraphicsSceneHoverEvent_pos(void* self) {
	auto&& __r = ((QGraphicsSceneHoverEvent*)self)->pos();
	return new QPointF(__r);
}
void QGraphicsSceneHoverEvent_setPos(void* self, void* pos) {
	((QGraphicsSceneHoverEvent*)self)->setPos(*(QPointF*)pos);
}
void* QGraphicsSceneHoverEvent_scenePos(void* self) {
	auto&& __r = ((QGraphicsSceneHoverEvent*)self)->scenePos();
	return new QPointF(__r);
}
void QGraphicsSceneHoverEvent_setScenePos(void* self, void* pos) {
	((QGraphicsSceneHoverEvent*)self)->setScenePos(*(QPointF*)pos);
}
void* QGraphicsSceneHoverEvent_screenPos(void* self) {
	auto&& __r = ((QGraphicsSceneHoverEvent*)self)->screenPos();
	return new QPoint(__r);
}
void QGraphicsSceneHoverEvent_setScreenPos(void* self, void* pos) {
	((QGraphicsSceneHoverEvent*)self)->setScreenPos(*(QPoint*)pos);
}
void* QGraphicsSceneHoverEvent_lastPos(void* self) {
	auto&& __r = ((QGraphicsSceneHoverEvent*)self)->lastPos();
	return new QPointF(__r);
}
void QGraphicsSceneHoverEvent_setLastPos(void* self, void* pos) {
	((QGraphicsSceneHoverEvent*)self)->setLastPos(*(QPointF*)pos);
}
void* QGraphicsSceneHoverEvent_lastScenePos(void* self) {
	auto&& __r = ((QGraphicsSceneHoverEvent*)self)->lastScenePos();
	return new QPointF(__r);
}
void QGraphicsSceneHoverEvent_setLastScenePos(void* self, void* pos) {
	((QGraphicsSceneHoverEvent*)self)->setLastScenePos(*(QPointF*)pos);
}
void* QGraphicsSceneHoverEvent_lastScreenPos(void* self) {
	auto&& __r = ((QGraphicsSceneHoverEvent*)self)->lastScreenPos();
	return new QPoint(__r);
}
void QGraphicsSceneHoverEvent_setLastScreenPos(void* self, void* pos) {
	((QGraphicsSceneHoverEvent*)self)->setLastScreenPos(*(QPoint*)pos);
}
int QGraphicsSceneHoverEvent_modifiers(void* self) {
	auto&& __r = ((QGraphicsSceneHoverEvent*)self)->modifiers();
	return (int)(__r);
}
void QGraphicsSceneHoverEvent_setModifiers(void* self, int modifiers) {
	((QGraphicsSceneHoverEvent*)self)->setModifiers(QFlags<Qt::KeyboardModifier>(modifiers));
}
int QGraphicsSceneHoverEvent_registerEventType() {
	auto&& __r = QGraphicsSceneHoverEvent::registerEventType();
	return __r;
}
int QGraphicsSceneHoverEvent_registerEventType2(int hint) {
	auto&& __r = QGraphicsSceneHoverEvent::registerEventType(hint);
	return __r;
}
void QGraphicsSceneHoverEvent_setAccepted_default(void* self, int accepted) {
	((QGraphicsSceneHoverEvent*)self)->QGraphicsSceneHoverEvent::setAccepted(accepted);
}
void* QGraphicsSceneHoverEvent_clone_default(void* self) {
	auto&& __r = ((QGraphicsSceneHoverEvent*)self)->QGraphicsSceneHoverEvent::clone();
	return (void*)(__r);
}
void* QGraphicsSceneHoverEvent__Go_new(long long id) { return new QGraphicsSceneHoverEvent__Go(id); }
void QGraphicsSceneHoverEvent_delete(void* self) { delete (QGraphicsSceneHoverEvent*)self; }

}

#include <QtWidgets/qmdiarea.h>
#include <QtWidgets/qabstractscrollarea.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qsize.h>
#include <QtWidgets/qmdisubwindow.h>
#include <QtCore/qnamespace.h>
#include <QtGui/qbrush.h>
#include <QtWidgets/qtabwidget.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qobject.h>
#include <QtGui/qevent.h>
#include <QtWidgets/qscrollbar.h>
#include <QtCore/qmargins.h>
#include <QtWidgets/qframe.h>
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
#include <QtCore/qbytearray.h>
#include <QtWidgets/qsizepolicy.h>
#include <QtWidgets/qlayout.h>
#include <QtGui/qaction.h>
#include <QtGui/qpaintengine.h>
#include <QtGui/qbackingstore.h>
#include <QtGui/qwindow.h>
#include <QtGui/qscreen.h>
#include <QtCore/qvariant.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qmdiarea.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QMdiArea__Go : public QMdiArea {
public:
	long long __goId;
	unsigned char __goOn[53] = {};
	template<typename... A> QMdiArea__Go(long long id, A&&... a) : QMdiArea(std::forward<A>(a)...), __goId(id) {}
	using QMdiArea::setViewportMargins;
	using QMdiArea::viewportMargins;
	using QMdiArea::drawFrame;
	using QMdiArea::updateMicroFocus;
	using QMdiArea::create;
	using QMdiArea::destroy;
	using QMdiArea::focusNextChild;
	using QMdiArea::focusPreviousChild;
	using QMdiArea::sender;
	using QMdiArea::senderSignalIndex;
	using QMdiArea::receivers;
	using QMdiArea::isSignalConnected;
	using QMdiArea::getDecodedMetricF;
	void cppgoBase_setupViewport(QWidget* __a0) { QMdiArea::setupViewport(__a0); }
	bool cppgoBase_event(QEvent* __a0) { return QMdiArea::event(__a0); }
	bool cppgoBase_eventFilter(QObject* __a0, QEvent* __a1) { return QMdiArea::eventFilter(__a0, __a1); }
	void cppgoBase_paintEvent(QPaintEvent* __a0) { QMdiArea::paintEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QMdiArea::childEvent(__a0); }
	void cppgoBase_resizeEvent(QResizeEvent* __a0) { QMdiArea::resizeEvent(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QMdiArea::timerEvent(__a0); }
	void cppgoBase_showEvent(QShowEvent* __a0) { QMdiArea::showEvent(__a0); }
	bool cppgoBase_viewportEvent(QEvent* __a0) { return QMdiArea::viewportEvent(__a0); }
	void cppgoBase_scrollContentsBy(int __a0, int __a1) { QMdiArea::scrollContentsBy(__a0, __a1); }
	void cppgoBase_mousePressEvent(QMouseEvent* __a0) { QMdiArea::mousePressEvent(__a0); }
	void cppgoBase_mouseReleaseEvent(QMouseEvent* __a0) { QMdiArea::mouseReleaseEvent(__a0); }
	void cppgoBase_mouseDoubleClickEvent(QMouseEvent* __a0) { QMdiArea::mouseDoubleClickEvent(__a0); }
	void cppgoBase_mouseMoveEvent(QMouseEvent* __a0) { QMdiArea::mouseMoveEvent(__a0); }
	void cppgoBase_wheelEvent(QWheelEvent* __a0) { QMdiArea::wheelEvent(__a0); }
	void cppgoBase_contextMenuEvent(QContextMenuEvent* __a0) { QMdiArea::contextMenuEvent(__a0); }
	void cppgoBase_dragEnterEvent(QDragEnterEvent* __a0) { QMdiArea::dragEnterEvent(__a0); }
	void cppgoBase_dragMoveEvent(QDragMoveEvent* __a0) { QMdiArea::dragMoveEvent(__a0); }
	void cppgoBase_dragLeaveEvent(QDragLeaveEvent* __a0) { QMdiArea::dragLeaveEvent(__a0); }
	void cppgoBase_dropEvent(QDropEvent* __a0) { QMdiArea::dropEvent(__a0); }
	void cppgoBase_keyPressEvent(QKeyEvent* __a0) { QMdiArea::keyPressEvent(__a0); }
	QSize cppgoBase_viewportSizeHint() { return QMdiArea::viewportSizeHint(); }
	void cppgoBase_changeEvent(QEvent* __a0) { QMdiArea::changeEvent(__a0); }
	void cppgoBase_initStyleOption(QStyleOptionFrame* __a0) { QMdiArea::initStyleOption(__a0); }
	void cppgoBase_keyReleaseEvent(QKeyEvent* __a0) { QMdiArea::keyReleaseEvent(__a0); }
	void cppgoBase_focusInEvent(QFocusEvent* __a0) { QMdiArea::focusInEvent(__a0); }
	void cppgoBase_focusOutEvent(QFocusEvent* __a0) { QMdiArea::focusOutEvent(__a0); }
	void cppgoBase_enterEvent(QEnterEvent* __a0) { QMdiArea::enterEvent(__a0); }
	void cppgoBase_leaveEvent(QEvent* __a0) { QMdiArea::leaveEvent(__a0); }
	void cppgoBase_moveEvent(QMoveEvent* __a0) { QMdiArea::moveEvent(__a0); }
	void cppgoBase_closeEvent(QCloseEvent* __a0) { QMdiArea::closeEvent(__a0); }
	void cppgoBase_tabletEvent(QTabletEvent* __a0) { QMdiArea::tabletEvent(__a0); }
	void cppgoBase_actionEvent(QActionEvent* __a0) { QMdiArea::actionEvent(__a0); }
	void cppgoBase_hideEvent(QHideEvent* __a0) { QMdiArea::hideEvent(__a0); }
	int cppgoBase_metric(int __a0) { return QMdiArea::metric((QPaintDevice::PaintDeviceMetric)__a0); }
	void cppgoBase_initPainter(QPainter* __a0) { QMdiArea::initPainter(__a0); }
	QPaintDevice* cppgoBase_redirected(QPoint* __a0) { return QMdiArea::redirected(__a0); }
	QPainter* cppgoBase_sharedPainter() { return QMdiArea::sharedPainter(); }
	void cppgoBase_inputMethodEvent(QInputMethodEvent* __a0) { QMdiArea::inputMethodEvent(__a0); }
	bool cppgoBase_focusNextPrevChild(bool __a0) { return QMdiArea::focusNextPrevChild(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QMdiArea::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QMdiArea::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QMdiArea::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QMdiArea::metaObject();
	}
	QSize sizeHint() const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 729690261457684101LL, &__h); if (__h) return *(QSize*)__r; }
		return QMdiArea::sizeHint();
	}
	QSize minimumSizeHint() const override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 7049264599837532491LL, &__h); if (__h) return *(QSize*)__r; }
		return QMdiArea::minimumSizeHint();
	}
	void setupViewport(QWidget* viewport) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4531996381736730688LL, (void*)viewport, &__h); if (__h) return; }
		QMdiArea::setupViewport(viewport);
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QMdiArea::event(event);
	}
	bool eventFilter(QObject* object, QEvent* event) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)object, (void*)event, &__h); if (__h) return __r; }
		return QMdiArea::eventFilter(object, event);
	}
	void paintEvent(QPaintEvent* paintEvent) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5290442163167458465LL, (void*)paintEvent, &__h); if (__h) return; }
		QMdiArea::paintEvent(paintEvent);
	}
	void childEvent(QChildEvent* childEvent) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)childEvent, &__h); if (__h) return; }
		QMdiArea::childEvent(childEvent);
	}
	void resizeEvent(QResizeEvent* resizeEvent) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4302917285655014661LL, (void*)resizeEvent, &__h); if (__h) return; }
		QMdiArea::resizeEvent(resizeEvent);
	}
	void timerEvent(QTimerEvent* timerEvent) override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)timerEvent, &__h); if (__h) return; }
		QMdiArea::timerEvent(timerEvent);
	}
	void showEvent(QShowEvent* showEvent) override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6825369174751916LL, (void*)showEvent, &__h); if (__h) return; }
		QMdiArea::showEvent(showEvent);
	}
	bool viewportEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, -1971465653061861519LL, (void*)event, &__h); if (__h) return __r; }
		return QMdiArea::viewportEvent(event);
	}
	void scrollContentsBy(int dx, int dy) override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_void_int_int) { ((void(*)(long long, long long, int, int, int*))cppgo_virt_void_int_int)(__goId, -6118001150126564745LL, dx, dy, &__h); if (__h) return; }
		QMdiArea::scrollContentsBy(dx, dy);
	}
	void mousePressEvent(QMouseEvent* arg1) override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -332090898078543319LL, (void*)arg1, &__h); if (__h) return; }
		QMdiArea::mousePressEvent(arg1);
	}
	void mouseReleaseEvent(QMouseEvent* arg1) override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4415920049297795213LL, (void*)arg1, &__h); if (__h) return; }
		QMdiArea::mouseReleaseEvent(arg1);
	}
	void mouseDoubleClickEvent(QMouseEvent* arg1) override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4029073905697023887LL, (void*)arg1, &__h); if (__h) return; }
		QMdiArea::mouseDoubleClickEvent(arg1);
	}
	void mouseMoveEvent(QMouseEvent* arg1) override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4255494004099806581LL, (void*)arg1, &__h); if (__h) return; }
		QMdiArea::mouseMoveEvent(arg1);
	}
	void wheelEvent(QWheelEvent* arg1) override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1919988340934511316LL, (void*)arg1, &__h); if (__h) return; }
		QMdiArea::wheelEvent(arg1);
	}
	void contextMenuEvent(QContextMenuEvent* arg1) override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6698043222227006015LL, (void*)arg1, &__h); if (__h) return; }
		QMdiArea::contextMenuEvent(arg1);
	}
	void dragEnterEvent(QDragEnterEvent* arg1) override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 206122715919917589LL, (void*)arg1, &__h); if (__h) return; }
		QMdiArea::dragEnterEvent(arg1);
	}
	void dragMoveEvent(QDragMoveEvent* arg1) override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1781545137352745978LL, (void*)arg1, &__h); if (__h) return; }
		QMdiArea::dragMoveEvent(arg1);
	}
	void dragLeaveEvent(QDragLeaveEvent* arg1) override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -3139904158373725006LL, (void*)arg1, &__h); if (__h) return; }
		QMdiArea::dragLeaveEvent(arg1);
	}
	void dropEvent(QDropEvent* arg1) override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 9107058682754591716LL, (void*)arg1, &__h); if (__h) return; }
		QMdiArea::dropEvent(arg1);
	}
	void keyPressEvent(QKeyEvent* arg1) override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1023791977679793185LL, (void*)arg1, &__h); if (__h) return; }
		QMdiArea::keyPressEvent(arg1);
	}
	QSize viewportSizeHint() const override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 503579067252048239LL, &__h); if (__h) return *(QSize*)__r; }
		return QMdiArea::viewportSizeHint();
	}
	void changeEvent(QEvent* arg1) override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4959376063826165943LL, (void*)arg1, &__h); if (__h) return; }
		QMdiArea::changeEvent(arg1);
	}
	void initStyleOption(QStyleOptionFrame* option) const override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -488693769771308531LL, (void*)option, &__h); if (__h) return; }
		QMdiArea::initStyleOption(option);
	}
	int devType() const override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8659679893868524044LL, &__h); if (__h) return __r; }
		return QMdiArea::devType();
	}
	void setVisible(bool visible) override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 2110182719375245575LL, visible, &__h); if (__h) return; }
		QMdiArea::setVisible(visible);
	}
	int heightForWidth(int arg1) const override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 4311534060495737339LL, arg1, &__h); if (__h) return __r; }
		return QMdiArea::heightForWidth(arg1);
	}
	bool hasHeightForWidth() const override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8266762397995423915LL, &__h); if (__h) return __r; }
		return QMdiArea::hasHeightForWidth();
	}
	QPaintEngine* paintEngine() const override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 6398007404756346381LL, &__h); if (__h) return (QPaintEngine*)__r; }
		return QMdiArea::paintEngine();
	}
	void keyReleaseEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[32] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 772598475950050299LL, (void*)event, &__h); if (__h) return; }
		QMdiArea::keyReleaseEvent(event);
	}
	void focusInEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[33] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 975043767033552552LL, (void*)event, &__h); if (__h) return; }
		QMdiArea::focusInEvent(event);
	}
	void focusOutEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[34] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -385596010576048993LL, (void*)event, &__h); if (__h) return; }
		QMdiArea::focusOutEvent(event);
	}
	void enterEvent(QEnterEvent* event) override {
		int __h = 0;
		if (__goOn[35] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6316270346668237895LL, (void*)event, &__h); if (__h) return; }
		QMdiArea::enterEvent(event);
	}
	void leaveEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[36] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -551070009623711640LL, (void*)event, &__h); if (__h) return; }
		QMdiArea::leaveEvent(event);
	}
	void moveEvent(QMoveEvent* event) override {
		int __h = 0;
		if (__goOn[37] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -9161949150489525396LL, (void*)event, &__h); if (__h) return; }
		QMdiArea::moveEvent(event);
	}
	void closeEvent(QCloseEvent* event) override {
		int __h = 0;
		if (__goOn[38] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1989937050019464783LL, (void*)event, &__h); if (__h) return; }
		QMdiArea::closeEvent(event);
	}
	void tabletEvent(QTabletEvent* event) override {
		int __h = 0;
		if (__goOn[39] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -212068866860701141LL, (void*)event, &__h); if (__h) return; }
		QMdiArea::tabletEvent(event);
	}
	void actionEvent(QActionEvent* event) override {
		int __h = 0;
		if (__goOn[40] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1697444140088555581LL, (void*)event, &__h); if (__h) return; }
		QMdiArea::actionEvent(event);
	}
	void hideEvent(QHideEvent* event) override {
		int __h = 0;
		if (__goOn[41] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -199718707876719129LL, (void*)event, &__h); if (__h) return; }
		QMdiArea::hideEvent(event);
	}
	bool nativeEvent(const QByteArray& eventType, void* message, qintptr* result) override {
		int __h = 0;
		if (__goOn[42] && cppgo_virt_int_voidp_voidp_longlongp) { int __r = ((int(*)(long long, long long, void*, void*, long long*, int*))cppgo_virt_int_voidp_voidp_longlongp)(__goId, -2039166586246443706LL, (void*)&eventType, (void*)message, (long long*)result, &__h); if (__h) return __r; }
		return QMdiArea::nativeEvent(eventType, message, result);
	}
	int metric(QPaintDevice::PaintDeviceMetric arg1) const override {
		int __h = 0;
		if (__goOn[43] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 427629649267311777LL, (int)(arg1), &__h); if (__h) return __r; }
		return QMdiArea::metric(arg1);
	}
	void initPainter(QPainter* painter) const override {
		int __h = 0;
		if (__goOn[44] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6477500296980050210LL, (void*)painter, &__h); if (__h) return; }
		QMdiArea::initPainter(painter);
	}
	QPaintDevice* redirected(QPoint* offset) const override {
		int __h = 0;
		if (__goOn[45] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 4127240633310655048LL, (void*)offset, &__h); if (__h) return (QPaintDevice*)__r; }
		return QMdiArea::redirected(offset);
	}
	QPainter* sharedPainter() const override {
		int __h = 0;
		if (__goOn[46] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 5792771328741774991LL, &__h); if (__h) return (QPainter*)__r; }
		return QMdiArea::sharedPainter();
	}
	void inputMethodEvent(QInputMethodEvent* arg1) override {
		int __h = 0;
		if (__goOn[47] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5438819661942843222LL, (void*)arg1, &__h); if (__h) return; }
		QMdiArea::inputMethodEvent(arg1);
	}
	QVariant inputMethodQuery(Qt::InputMethodQuery arg1) const override {
		int __h = 0;
		if (__goOn[48] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7345853097340709638LL, (int)(arg1), &__h); if (__h) return *(QVariant*)__r; }
		return QMdiArea::inputMethodQuery(arg1);
	}
	bool focusNextPrevChild(bool next) override {
		int __h = 0;
		if (__goOn[49] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 5898328517505784191LL, next, &__h); if (__h) return __r; }
		return QMdiArea::focusNextPrevChild(next);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[50] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QMdiArea::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[51] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QMdiArea::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[52] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QMdiArea::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QMdiArea;
static std::mutex __connsMu_QMdiArea;

extern "C" {

void* QMdiArea_base(void* self) { return (void*)static_cast<QAbstractScrollArea*>((QMdiArea*)self); }
void QMdiArea_chain(void* self, void** out) {
	QMdiArea* __p = (QMdiArea*)self;
	QAbstractScrollArea* __b0 = static_cast<QAbstractScrollArea*>(__p); out[0] = (void*)__b0;
	QFrame* __b1 = static_cast<QFrame*>(__b0); out[1] = (void*)__b1;
	QWidget* __b2 = static_cast<QWidget*>(__b1); out[2] = (void*)__b2;
	QObject* __b3 = static_cast<QObject*>(__b2); out[3] = (void*)__b3;
}
void QMdiArea__set_override(void* self, int idx) { ((QMdiArea__Go*)self)->__goOn[idx] = 1; }
void* QMdiArea_fromQObject(void* o) { return (void*)qobject_cast<QMdiArea*>((QObject*)o); }
void* QMdiArea_new(long long __goid, void* parent) {
	return new QMdiArea__Go(__goid, (QWidget*)parent);
}
void* QMdiArea_new2(long long __goid) {
	return new QMdiArea__Go(__goid);
}
char* QMdiArea_tr(const char* s, long long* __out_len) {
	auto&& __r = QMdiArea::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QMdiArea_currentSubWindow(void* self) {
	auto&& __r = ((QMdiArea*)self)->currentSubWindow();
	return (void*)(__r);
}
void* QMdiArea_activeSubWindow(void* self) {
	auto&& __r = ((QMdiArea*)self)->activeSubWindow();
	return (void*)(__r);
}
void** QMdiArea_subWindowList(void* self, long long* __out_len) {
	auto&& __r = ((QMdiArea*)self)->subWindowList();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
void* QMdiArea_addSubWindow(void* self, void* widget) {
	auto&& __r = ((QMdiArea*)self)->addSubWindow((QWidget*)widget);
	return (void*)(__r);
}
void QMdiArea_removeSubWindow(void* self, void* widget) {
	((QMdiArea*)self)->removeSubWindow((QWidget*)widget);
}
void* QMdiArea_background(void* self) {
	auto&& __r = ((QMdiArea*)self)->background();
	return new QBrush(__r);
}
void QMdiArea_setBackground(void* self, void* background) {
	((QMdiArea*)self)->setBackground(*(QBrush*)background);
}
int QMdiArea_activationOrder(void* self) {
	auto&& __r = ((QMdiArea*)self)->activationOrder();
	return (int)(__r);
}
void QMdiArea_setActivationOrder(void* self, int order) {
	((QMdiArea*)self)->setActivationOrder((QMdiArea::WindowOrder)(order));
}
void QMdiArea_setOption(void* self, int option) {
	((QMdiArea*)self)->setOption((QMdiArea::AreaOption)(option));
}
int QMdiArea_testOption(void* self, int opton) {
	auto&& __r = ((QMdiArea*)self)->testOption((QMdiArea::AreaOption)(opton));
	return __r;
}
void QMdiArea_setViewMode(void* self, int mode) {
	((QMdiArea*)self)->setViewMode((QMdiArea::ViewMode)(mode));
}
int QMdiArea_viewMode(void* self) {
	auto&& __r = ((QMdiArea*)self)->viewMode();
	return (int)(__r);
}
int QMdiArea_documentMode(void* self) {
	auto&& __r = ((QMdiArea*)self)->documentMode();
	return __r;
}
void QMdiArea_setDocumentMode(void* self, int enabled) {
	((QMdiArea*)self)->setDocumentMode(enabled);
}
void QMdiArea_setTabsClosable(void* self, int closable) {
	((QMdiArea*)self)->setTabsClosable(closable);
}
int QMdiArea_tabsClosable(void* self) {
	auto&& __r = ((QMdiArea*)self)->tabsClosable();
	return __r;
}
void QMdiArea_setTabsMovable(void* self, int movable) {
	((QMdiArea*)self)->setTabsMovable(movable);
}
int QMdiArea_tabsMovable(void* self) {
	auto&& __r = ((QMdiArea*)self)->tabsMovable();
	return __r;
}
void QMdiArea_setTabShape(void* self, int shape) {
	((QMdiArea*)self)->setTabShape((QTabWidget::TabShape)(shape));
}
int QMdiArea_tabShape(void* self) {
	auto&& __r = ((QMdiArea*)self)->tabShape();
	return (int)(__r);
}
void QMdiArea_setTabPosition(void* self, int position) {
	((QMdiArea*)self)->setTabPosition((QTabWidget::TabPosition)(position));
}
int QMdiArea_tabPosition(void* self) {
	auto&& __r = ((QMdiArea*)self)->tabPosition();
	return (int)(__r);
}
void QMdiArea_subWindowActivated(void* self, void* arg1) {
	((QMdiArea*)self)->subWindowActivated((QMdiSubWindow*)arg1);
}
void QMdiArea_setActiveSubWindow(void* self, void* window) {
	((QMdiArea*)self)->setActiveSubWindow((QMdiSubWindow*)window);
}
void QMdiArea_tileSubWindows(void* self) {
	((QMdiArea*)self)->tileSubWindows();
}
void QMdiArea_cascadeSubWindows(void* self) {
	((QMdiArea*)self)->cascadeSubWindows();
}
void QMdiArea_closeActiveSubWindow(void* self) {
	((QMdiArea*)self)->closeActiveSubWindow();
}
void QMdiArea_closeAllSubWindows(void* self) {
	((QMdiArea*)self)->closeAllSubWindows();
}
void QMdiArea_activateNextSubWindow(void* self) {
	((QMdiArea*)self)->activateNextSubWindow();
}
void QMdiArea_activatePreviousSubWindow(void* self) {
	((QMdiArea*)self)->activatePreviousSubWindow();
}
void QMdiArea_setViewportMargins(void* self, int left, int top, int right, int bottom) {
	((QMdiArea__Go*)self)->setViewportMargins(left, top, right, bottom);
}
void QMdiArea_setViewportMargins2(void* self, void* margins) {
	((QMdiArea__Go*)self)->setViewportMargins(*(QMargins*)margins);
}
void* QMdiArea_viewportMargins(void* self) {
	auto&& __r = ((QMdiArea__Go*)self)->viewportMargins();
	return new QMargins(__r);
}
void QMdiArea_drawFrame(void* self, void* arg1) {
	((QMdiArea__Go*)self)->drawFrame((QPainter*)arg1);
}
void QMdiArea_setTabOrder(void* arg1, void* arg2) {
	QMdiArea::setTabOrder((QWidget*)arg1, (QWidget*)arg2);
}
void* QMdiArea_mouseGrabber() {
	auto&& __r = QMdiArea::mouseGrabber();
	return (void*)(__r);
}
void* QMdiArea_keyboardGrabber() {
	auto&& __r = QMdiArea::keyboardGrabber();
	return (void*)(__r);
}
void* QMdiArea_find(unsigned long long arg1) {
	auto&& __r = QMdiArea::find(arg1);
	return (void*)(__r);
}
void* QMdiArea_createWindowContainer(void* window) {
	auto&& __r = QMdiArea::createWindowContainer((QWindow*)window);
	return (void*)(__r);
}
void QMdiArea_updateMicroFocus(void* self) {
	((QMdiArea__Go*)self)->updateMicroFocus();
}
void QMdiArea_create(void* self) {
	((QMdiArea__Go*)self)->create();
}
void QMdiArea_destroy(void* self) {
	((QMdiArea__Go*)self)->destroy();
}
int QMdiArea_focusNextChild(void* self) {
	auto&& __r = ((QMdiArea__Go*)self)->focusNextChild();
	return __r;
}
int QMdiArea_focusPreviousChild(void* self) {
	auto&& __r = ((QMdiArea__Go*)self)->focusPreviousChild();
	return __r;
}
void* QMdiArea_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QMdiArea::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QMdiArea_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QMdiArea::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QMdiArea_sender(void* self) {
	auto&& __r = ((QMdiArea__Go*)self)->sender();
	return (void*)(__r);
}
int QMdiArea_senderSignalIndex(void* self) {
	auto&& __r = ((QMdiArea__Go*)self)->senderSignalIndex();
	return __r;
}
int QMdiArea_receivers(void* self, const char* signal) {
	auto&& __r = ((QMdiArea__Go*)self)->receivers(signal);
	return __r;
}
int QMdiArea_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QMdiArea__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
double QMdiArea_devicePixelRatioFScale() {
	auto&& __r = QMdiArea::devicePixelRatioFScale();
	return __r;
}
int QMdiArea_encodeMetricF(int metric, double value) {
	auto&& __r = QMdiArea::encodeMetricF((QPaintDevice::PaintDeviceMetric)(metric), value);
	return __r;
}
double QMdiArea_getDecodedMetricF(void* self, int metricA, int metricB) {
	auto&& __r = ((QMdiArea__Go*)self)->getDecodedMetricF((QPaintDevice::PaintDeviceMetric)(metricA), (QPaintDevice::PaintDeviceMetric)(metricB));
	return __r;
}
char* QMdiArea_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QMdiArea::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QMdiArea_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QMdiArea::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void** QMdiArea_subWindowList2(void* self, int order, long long* __out_len) {
	auto&& __r = ((QMdiArea*)self)->subWindowList((QMdiArea::WindowOrder)(order));
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
void* QMdiArea_addSubWindow2(void* self, void* widget, int flags) {
	auto&& __r = ((QMdiArea*)self)->addSubWindow((QWidget*)widget, QFlags<Qt::WindowType>(flags));
	return (void*)(__r);
}
void QMdiArea_setOption2(void* self, int option, int on) {
	((QMdiArea*)self)->setOption((QMdiArea::AreaOption)(option), on);
}
void* QMdiArea_createWindowContainer2(void* window, void* parent) {
	auto&& __r = QMdiArea::createWindowContainer((QWindow*)window, (QWidget*)parent);
	return (void*)(__r);
}
void* QMdiArea_createWindowContainer3(void* window, void* parent, int flags) {
	auto&& __r = QMdiArea::createWindowContainer((QWindow*)window, (QWidget*)parent, QFlags<Qt::WindowType>(flags));
	return (void*)(__r);
}
void QMdiArea_updateMicroFocus2(void* self, int query) {
	((QMdiArea__Go*)self)->updateMicroFocus((Qt::InputMethodQuery)(query));
}
void QMdiArea_create2(void* self, unsigned long long arg1) {
	((QMdiArea__Go*)self)->create(arg1);
}
void QMdiArea_create3(void* self, unsigned long long arg1, int initializeWindow) {
	((QMdiArea__Go*)self)->create(arg1, initializeWindow);
}
void QMdiArea_create4(void* self, unsigned long long arg1, int initializeWindow, int destroyOldWindow) {
	((QMdiArea__Go*)self)->create(arg1, initializeWindow, destroyOldWindow);
}
void QMdiArea_destroy2(void* self, int destroyWindow) {
	((QMdiArea__Go*)self)->destroy(destroyWindow);
}
void QMdiArea_destroy3(void* self, int destroyWindow, int destroySubWindows) {
	((QMdiArea__Go*)self)->destroy(destroyWindow, destroySubWindows);
}
void* QMdiArea_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QMdiArea::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QMdiArea_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QMdiArea::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QMdiArea_metaObject_default(void* self) {
	auto&& __r = ((QMdiArea*)self)->QMdiArea::metaObject();
	return (void*)(__r);
}
void* QMdiArea_sizeHint_default(void* self) {
	auto&& __r = ((QMdiArea*)self)->QMdiArea::sizeHint();
	return new QSize(__r);
}
void* QMdiArea_minimumSizeHint_default(void* self) {
	auto&& __r = ((QMdiArea*)self)->QMdiArea::minimumSizeHint();
	return new QSize(__r);
}
void QMdiArea_setupViewport_default(void* self, void* viewport) {
	((QMdiArea__Go*)self)->cppgoBase_setupViewport((QWidget*)viewport);
}
int QMdiArea_event_default(void* self, void* event) {
	auto&& __r = ((QMdiArea__Go*)self)->cppgoBase_event((QEvent*)event);
	return __r;
}
int QMdiArea_eventFilter_default(void* self, void* object, void* event) {
	auto&& __r = ((QMdiArea__Go*)self)->cppgoBase_eventFilter((QObject*)object, (QEvent*)event);
	return __r;
}
void QMdiArea_paintEvent_default(void* self, void* paintEvent) {
	((QMdiArea__Go*)self)->cppgoBase_paintEvent((QPaintEvent*)paintEvent);
}
void QMdiArea_childEvent_default(void* self, void* childEvent) {
	((QMdiArea__Go*)self)->cppgoBase_childEvent((QChildEvent*)childEvent);
}
void QMdiArea_resizeEvent_default(void* self, void* resizeEvent) {
	((QMdiArea__Go*)self)->cppgoBase_resizeEvent((QResizeEvent*)resizeEvent);
}
void QMdiArea_timerEvent_default(void* self, void* timerEvent) {
	((QMdiArea__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)timerEvent);
}
void QMdiArea_showEvent_default(void* self, void* showEvent) {
	((QMdiArea__Go*)self)->cppgoBase_showEvent((QShowEvent*)showEvent);
}
int QMdiArea_viewportEvent_default(void* self, void* event) {
	auto&& __r = ((QMdiArea__Go*)self)->cppgoBase_viewportEvent((QEvent*)event);
	return __r;
}
void QMdiArea_scrollContentsBy_default(void* self, int dx, int dy) {
	((QMdiArea__Go*)self)->cppgoBase_scrollContentsBy(dx, dy);
}
void QMdiArea_mousePressEvent_default(void* self, void* arg1) {
	((QMdiArea__Go*)self)->cppgoBase_mousePressEvent((QMouseEvent*)arg1);
}
void QMdiArea_mouseReleaseEvent_default(void* self, void* arg1) {
	((QMdiArea__Go*)self)->cppgoBase_mouseReleaseEvent((QMouseEvent*)arg1);
}
void QMdiArea_mouseDoubleClickEvent_default(void* self, void* arg1) {
	((QMdiArea__Go*)self)->cppgoBase_mouseDoubleClickEvent((QMouseEvent*)arg1);
}
void QMdiArea_mouseMoveEvent_default(void* self, void* arg1) {
	((QMdiArea__Go*)self)->cppgoBase_mouseMoveEvent((QMouseEvent*)arg1);
}
void QMdiArea_wheelEvent_default(void* self, void* arg1) {
	((QMdiArea__Go*)self)->cppgoBase_wheelEvent((QWheelEvent*)arg1);
}
void QMdiArea_contextMenuEvent_default(void* self, void* arg1) {
	((QMdiArea__Go*)self)->cppgoBase_contextMenuEvent((QContextMenuEvent*)arg1);
}
void QMdiArea_dragEnterEvent_default(void* self, void* arg1) {
	((QMdiArea__Go*)self)->cppgoBase_dragEnterEvent((QDragEnterEvent*)arg1);
}
void QMdiArea_dragMoveEvent_default(void* self, void* arg1) {
	((QMdiArea__Go*)self)->cppgoBase_dragMoveEvent((QDragMoveEvent*)arg1);
}
void QMdiArea_dragLeaveEvent_default(void* self, void* arg1) {
	((QMdiArea__Go*)self)->cppgoBase_dragLeaveEvent((QDragLeaveEvent*)arg1);
}
void QMdiArea_dropEvent_default(void* self, void* arg1) {
	((QMdiArea__Go*)self)->cppgoBase_dropEvent((QDropEvent*)arg1);
}
void QMdiArea_keyPressEvent_default(void* self, void* arg1) {
	((QMdiArea__Go*)self)->cppgoBase_keyPressEvent((QKeyEvent*)arg1);
}
void* QMdiArea_viewportSizeHint_default(void* self) {
	auto&& __r = ((QMdiArea__Go*)self)->cppgoBase_viewportSizeHint();
	return new QSize(__r);
}
void QMdiArea_changeEvent_default(void* self, void* arg1) {
	((QMdiArea__Go*)self)->cppgoBase_changeEvent((QEvent*)arg1);
}
void QMdiArea_initStyleOption_default(void* self, void* option) {
	((QMdiArea__Go*)self)->cppgoBase_initStyleOption((QStyleOptionFrame*)option);
}
int QMdiArea_devType_default(void* self) {
	auto&& __r = ((QMdiArea*)self)->QMdiArea::devType();
	return __r;
}
void QMdiArea_setVisible_default(void* self, int visible) {
	((QMdiArea*)self)->QMdiArea::setVisible(visible);
}
int QMdiArea_heightForWidth_default(void* self, int arg1) {
	auto&& __r = ((QMdiArea*)self)->QMdiArea::heightForWidth(arg1);
	return __r;
}
int QMdiArea_hasHeightForWidth_default(void* self) {
	auto&& __r = ((QMdiArea*)self)->QMdiArea::hasHeightForWidth();
	return __r;
}
void* QMdiArea_paintEngine_default(void* self) {
	auto&& __r = ((QMdiArea*)self)->QMdiArea::paintEngine();
	return (void*)(__r);
}
void QMdiArea_keyReleaseEvent_default(void* self, void* event) {
	((QMdiArea__Go*)self)->cppgoBase_keyReleaseEvent((QKeyEvent*)event);
}
void QMdiArea_focusInEvent_default(void* self, void* event) {
	((QMdiArea__Go*)self)->cppgoBase_focusInEvent((QFocusEvent*)event);
}
void QMdiArea_focusOutEvent_default(void* self, void* event) {
	((QMdiArea__Go*)self)->cppgoBase_focusOutEvent((QFocusEvent*)event);
}
void QMdiArea_enterEvent_default(void* self, void* event) {
	((QMdiArea__Go*)self)->cppgoBase_enterEvent((QEnterEvent*)event);
}
void QMdiArea_leaveEvent_default(void* self, void* event) {
	((QMdiArea__Go*)self)->cppgoBase_leaveEvent((QEvent*)event);
}
void QMdiArea_moveEvent_default(void* self, void* event) {
	((QMdiArea__Go*)self)->cppgoBase_moveEvent((QMoveEvent*)event);
}
void QMdiArea_closeEvent_default(void* self, void* event) {
	((QMdiArea__Go*)self)->cppgoBase_closeEvent((QCloseEvent*)event);
}
void QMdiArea_tabletEvent_default(void* self, void* event) {
	((QMdiArea__Go*)self)->cppgoBase_tabletEvent((QTabletEvent*)event);
}
void QMdiArea_actionEvent_default(void* self, void* event) {
	((QMdiArea__Go*)self)->cppgoBase_actionEvent((QActionEvent*)event);
}
void QMdiArea_hideEvent_default(void* self, void* event) {
	((QMdiArea__Go*)self)->cppgoBase_hideEvent((QHideEvent*)event);
}
int QMdiArea_metric_default(void* self, int arg1) {
	auto&& __r = ((QMdiArea__Go*)self)->cppgoBase_metric(arg1);
	return __r;
}
void QMdiArea_initPainter_default(void* self, void* painter) {
	((QMdiArea__Go*)self)->cppgoBase_initPainter((QPainter*)painter);
}
void* QMdiArea_redirected_default(void* self, void* offset) {
	auto&& __r = ((QMdiArea__Go*)self)->cppgoBase_redirected((QPoint*)offset);
	return (void*)(__r);
}
void* QMdiArea_sharedPainter_default(void* self) {
	auto&& __r = ((QMdiArea__Go*)self)->cppgoBase_sharedPainter();
	return (void*)(__r);
}
void QMdiArea_inputMethodEvent_default(void* self, void* arg1) {
	((QMdiArea__Go*)self)->cppgoBase_inputMethodEvent((QInputMethodEvent*)arg1);
}
void* QMdiArea_inputMethodQuery_default(void* self, int arg1) {
	auto&& __r = ((QMdiArea*)self)->QMdiArea::inputMethodQuery((Qt::InputMethodQuery)(arg1));
	return new QVariant(__r);
}
int QMdiArea_focusNextPrevChild_default(void* self, int next) {
	auto&& __r = ((QMdiArea__Go*)self)->cppgoBase_focusNextPrevChild(next);
	return __r;
}
void QMdiArea_customEvent_default(void* self, void* event) {
	((QMdiArea__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QMdiArea_connectNotify_default(void* self, void* signal) {
	((QMdiArea__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QMdiArea_disconnectNotify_default(void* self, void* signal) {
	((QMdiArea__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QMdiArea_connect_subWindowActivated(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QMdiArea*)self, &QMdiArea::subWindowActivated, (QMdiArea*)self, [__id](QMdiSubWindow* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QMdiArea);
	__conns_QMdiArea[__id] = __c;
}
void QMdiArea_connect_windowTitleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QMdiArea*)self, &QMdiArea::windowTitleChanged, (QMdiArea*)self, [__id](const QString& title){
		QByteArray title_ba = title.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)title_ba.constData(), title_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QMdiArea);
	__conns_QMdiArea[__id] = __c;
}
void QMdiArea_connect_windowIconChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QMdiArea*)self, &QMdiArea::windowIconChanged, (QMdiArea*)self, [__id](const QIcon& icon){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&icon);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QMdiArea);
	__conns_QMdiArea[__id] = __c;
}
void QMdiArea_connect_windowIconTextChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QMdiArea*)self, &QMdiArea::windowIconTextChanged, (QMdiArea*)self, [__id](const QString& iconText){
		QByteArray iconText_ba = iconText.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)iconText_ba.constData(), iconText_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QMdiArea);
	__conns_QMdiArea[__id] = __c;
}
void QMdiArea_connect_customContextMenuRequested(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QMdiArea*)self, &QMdiArea::customContextMenuRequested, (QMdiArea*)self, [__id](const QPoint& pos){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&pos);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QMdiArea);
	__conns_QMdiArea[__id] = __c;
}
void QMdiArea_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QMdiArea*)self, &QMdiArea::destroyed, (QMdiArea*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QMdiArea);
	__conns_QMdiArea[__id] = __c;
}
void QMdiArea_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QMdiArea*)self, &QMdiArea::destroyed, (QMdiArea*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QMdiArea);
	__conns_QMdiArea[__id] = __c;
}
void QMdiArea_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QMdiArea*)self, &QMdiArea::objectNameChanged, (QMdiArea*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QMdiArea);
	__conns_QMdiArea[__id] = __c;
}
void QMdiArea__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QMdiArea);
		auto __it = __conns_QMdiArea.find(id);
		if (__it == __conns_QMdiArea.end()) return;
		__c = __it->second;
		__conns_QMdiArea.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QMdiArea__Go_new(long long id) { return new QMdiArea__Go(id); }
void QMdiArea_delete(void* self) { delete (QMdiArea*)self; }

}

#include <QtWidgets/qstackedwidget.h>
#include <QtWidgets/qframe.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qsize.h>
#include <QtCore/qrect.h>
#include <QtGui/qevent.h>
#include <QtGui/qpainter.h>
#include <QtWidgets/qstyleoption.h>
#include <QtWidgets/qstyle.h>
#include <QtCore/qnamespace.h>
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
#include "gen_qstackedwidget.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QStackedWidget__Go : public QStackedWidget {
public:
	long long __goId;
	unsigned char __goOn[49] = {};
	template<typename... A> QStackedWidget__Go(long long id, A&&... a) : QStackedWidget(std::forward<A>(a)...), __goId(id) {}
	using QStackedWidget::drawFrame;
	using QStackedWidget::updateMicroFocus;
	using QStackedWidget::create;
	using QStackedWidget::destroy;
	using QStackedWidget::focusNextChild;
	using QStackedWidget::focusPreviousChild;
	using QStackedWidget::sender;
	using QStackedWidget::senderSignalIndex;
	using QStackedWidget::receivers;
	using QStackedWidget::isSignalConnected;
	using QStackedWidget::getDecodedMetricF;
	bool cppgoBase_event(QEvent* __a0) { return QStackedWidget::event(__a0); }
	void cppgoBase_paintEvent(QPaintEvent* __a0) { QStackedWidget::paintEvent(__a0); }
	void cppgoBase_changeEvent(QEvent* __a0) { QStackedWidget::changeEvent(__a0); }
	void cppgoBase_initStyleOption(QStyleOptionFrame* __a0) { QStackedWidget::initStyleOption(__a0); }
	void cppgoBase_mousePressEvent(QMouseEvent* __a0) { QStackedWidget::mousePressEvent(__a0); }
	void cppgoBase_mouseReleaseEvent(QMouseEvent* __a0) { QStackedWidget::mouseReleaseEvent(__a0); }
	void cppgoBase_mouseDoubleClickEvent(QMouseEvent* __a0) { QStackedWidget::mouseDoubleClickEvent(__a0); }
	void cppgoBase_mouseMoveEvent(QMouseEvent* __a0) { QStackedWidget::mouseMoveEvent(__a0); }
	void cppgoBase_wheelEvent(QWheelEvent* __a0) { QStackedWidget::wheelEvent(__a0); }
	void cppgoBase_keyPressEvent(QKeyEvent* __a0) { QStackedWidget::keyPressEvent(__a0); }
	void cppgoBase_keyReleaseEvent(QKeyEvent* __a0) { QStackedWidget::keyReleaseEvent(__a0); }
	void cppgoBase_focusInEvent(QFocusEvent* __a0) { QStackedWidget::focusInEvent(__a0); }
	void cppgoBase_focusOutEvent(QFocusEvent* __a0) { QStackedWidget::focusOutEvent(__a0); }
	void cppgoBase_enterEvent(QEnterEvent* __a0) { QStackedWidget::enterEvent(__a0); }
	void cppgoBase_leaveEvent(QEvent* __a0) { QStackedWidget::leaveEvent(__a0); }
	void cppgoBase_moveEvent(QMoveEvent* __a0) { QStackedWidget::moveEvent(__a0); }
	void cppgoBase_resizeEvent(QResizeEvent* __a0) { QStackedWidget::resizeEvent(__a0); }
	void cppgoBase_closeEvent(QCloseEvent* __a0) { QStackedWidget::closeEvent(__a0); }
	void cppgoBase_contextMenuEvent(QContextMenuEvent* __a0) { QStackedWidget::contextMenuEvent(__a0); }
	void cppgoBase_tabletEvent(QTabletEvent* __a0) { QStackedWidget::tabletEvent(__a0); }
	void cppgoBase_actionEvent(QActionEvent* __a0) { QStackedWidget::actionEvent(__a0); }
	void cppgoBase_dragEnterEvent(QDragEnterEvent* __a0) { QStackedWidget::dragEnterEvent(__a0); }
	void cppgoBase_dragMoveEvent(QDragMoveEvent* __a0) { QStackedWidget::dragMoveEvent(__a0); }
	void cppgoBase_dragLeaveEvent(QDragLeaveEvent* __a0) { QStackedWidget::dragLeaveEvent(__a0); }
	void cppgoBase_dropEvent(QDropEvent* __a0) { QStackedWidget::dropEvent(__a0); }
	void cppgoBase_showEvent(QShowEvent* __a0) { QStackedWidget::showEvent(__a0); }
	void cppgoBase_hideEvent(QHideEvent* __a0) { QStackedWidget::hideEvent(__a0); }
	int cppgoBase_metric(int __a0) { return QStackedWidget::metric((QPaintDevice::PaintDeviceMetric)__a0); }
	void cppgoBase_initPainter(QPainter* __a0) { QStackedWidget::initPainter(__a0); }
	QPaintDevice* cppgoBase_redirected(QPoint* __a0) { return QStackedWidget::redirected(__a0); }
	QPainter* cppgoBase_sharedPainter() { return QStackedWidget::sharedPainter(); }
	void cppgoBase_inputMethodEvent(QInputMethodEvent* __a0) { QStackedWidget::inputMethodEvent(__a0); }
	bool cppgoBase_focusNextPrevChild(bool __a0) { return QStackedWidget::focusNextPrevChild(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QStackedWidget::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QStackedWidget::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QStackedWidget::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QStackedWidget::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QStackedWidget::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QStackedWidget::metaObject();
	}
	bool event(QEvent* e) override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)e, &__h); if (__h) return __r; }
		return QStackedWidget::event(e);
	}
	QSize sizeHint() const override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 729690261457684101LL, &__h); if (__h) return *(QSize*)__r; }
		return QStackedWidget::sizeHint();
	}
	void paintEvent(QPaintEvent* arg1) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5290442163167458465LL, (void*)arg1, &__h); if (__h) return; }
		QStackedWidget::paintEvent(arg1);
	}
	void changeEvent(QEvent* arg1) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4959376063826165943LL, (void*)arg1, &__h); if (__h) return; }
		QStackedWidget::changeEvent(arg1);
	}
	void initStyleOption(QStyleOptionFrame* option) const override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -488693769771308531LL, (void*)option, &__h); if (__h) return; }
		QStackedWidget::initStyleOption(option);
	}
	int devType() const override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8659679893868524044LL, &__h); if (__h) return __r; }
		return QStackedWidget::devType();
	}
	void setVisible(bool visible) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 2110182719375245575LL, visible, &__h); if (__h) return; }
		QStackedWidget::setVisible(visible);
	}
	QSize minimumSizeHint() const override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 7049264599837532491LL, &__h); if (__h) return *(QSize*)__r; }
		return QStackedWidget::minimumSizeHint();
	}
	int heightForWidth(int arg1) const override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 4311534060495737339LL, arg1, &__h); if (__h) return __r; }
		return QStackedWidget::heightForWidth(arg1);
	}
	bool hasHeightForWidth() const override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8266762397995423915LL, &__h); if (__h) return __r; }
		return QStackedWidget::hasHeightForWidth();
	}
	QPaintEngine* paintEngine() const override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 6398007404756346381LL, &__h); if (__h) return (QPaintEngine*)__r; }
		return QStackedWidget::paintEngine();
	}
	void mousePressEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -332090898078543319LL, (void*)event, &__h); if (__h) return; }
		QStackedWidget::mousePressEvent(event);
	}
	void mouseReleaseEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4415920049297795213LL, (void*)event, &__h); if (__h) return; }
		QStackedWidget::mouseReleaseEvent(event);
	}
	void mouseDoubleClickEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4029073905697023887LL, (void*)event, &__h); if (__h) return; }
		QStackedWidget::mouseDoubleClickEvent(event);
	}
	void mouseMoveEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4255494004099806581LL, (void*)event, &__h); if (__h) return; }
		QStackedWidget::mouseMoveEvent(event);
	}
	void wheelEvent(QWheelEvent* event) override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1919988340934511316LL, (void*)event, &__h); if (__h) return; }
		QStackedWidget::wheelEvent(event);
	}
	void keyPressEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1023791977679793185LL, (void*)event, &__h); if (__h) return; }
		QStackedWidget::keyPressEvent(event);
	}
	void keyReleaseEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 772598475950050299LL, (void*)event, &__h); if (__h) return; }
		QStackedWidget::keyReleaseEvent(event);
	}
	void focusInEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 975043767033552552LL, (void*)event, &__h); if (__h) return; }
		QStackedWidget::focusInEvent(event);
	}
	void focusOutEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -385596010576048993LL, (void*)event, &__h); if (__h) return; }
		QStackedWidget::focusOutEvent(event);
	}
	void enterEvent(QEnterEvent* event) override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6316270346668237895LL, (void*)event, &__h); if (__h) return; }
		QStackedWidget::enterEvent(event);
	}
	void leaveEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -551070009623711640LL, (void*)event, &__h); if (__h) return; }
		QStackedWidget::leaveEvent(event);
	}
	void moveEvent(QMoveEvent* event) override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -9161949150489525396LL, (void*)event, &__h); if (__h) return; }
		QStackedWidget::moveEvent(event);
	}
	void resizeEvent(QResizeEvent* event) override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4302917285655014661LL, (void*)event, &__h); if (__h) return; }
		QStackedWidget::resizeEvent(event);
	}
	void closeEvent(QCloseEvent* event) override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1989937050019464783LL, (void*)event, &__h); if (__h) return; }
		QStackedWidget::closeEvent(event);
	}
	void contextMenuEvent(QContextMenuEvent* event) override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6698043222227006015LL, (void*)event, &__h); if (__h) return; }
		QStackedWidget::contextMenuEvent(event);
	}
	void tabletEvent(QTabletEvent* event) override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -212068866860701141LL, (void*)event, &__h); if (__h) return; }
		QStackedWidget::tabletEvent(event);
	}
	void actionEvent(QActionEvent* event) override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1697444140088555581LL, (void*)event, &__h); if (__h) return; }
		QStackedWidget::actionEvent(event);
	}
	void dragEnterEvent(QDragEnterEvent* event) override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 206122715919917589LL, (void*)event, &__h); if (__h) return; }
		QStackedWidget::dragEnterEvent(event);
	}
	void dragMoveEvent(QDragMoveEvent* event) override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1781545137352745978LL, (void*)event, &__h); if (__h) return; }
		QStackedWidget::dragMoveEvent(event);
	}
	void dragLeaveEvent(QDragLeaveEvent* event) override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -3139904158373725006LL, (void*)event, &__h); if (__h) return; }
		QStackedWidget::dragLeaveEvent(event);
	}
	void dropEvent(QDropEvent* event) override {
		int __h = 0;
		if (__goOn[32] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 9107058682754591716LL, (void*)event, &__h); if (__h) return; }
		QStackedWidget::dropEvent(event);
	}
	void showEvent(QShowEvent* event) override {
		int __h = 0;
		if (__goOn[33] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6825369174751916LL, (void*)event, &__h); if (__h) return; }
		QStackedWidget::showEvent(event);
	}
	void hideEvent(QHideEvent* event) override {
		int __h = 0;
		if (__goOn[34] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -199718707876719129LL, (void*)event, &__h); if (__h) return; }
		QStackedWidget::hideEvent(event);
	}
	bool nativeEvent(const QByteArray& eventType, void* message, qintptr* result) override {
		int __h = 0;
		if (__goOn[35] && cppgo_virt_int_voidp_voidp_longlongp) { int __r = ((int(*)(long long, long long, void*, void*, long long*, int*))cppgo_virt_int_voidp_voidp_longlongp)(__goId, -2039166586246443706LL, (void*)&eventType, (void*)message, (long long*)result, &__h); if (__h) return __r; }
		return QStackedWidget::nativeEvent(eventType, message, result);
	}
	int metric(QPaintDevice::PaintDeviceMetric arg1) const override {
		int __h = 0;
		if (__goOn[36] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 427629649267311777LL, (int)(arg1), &__h); if (__h) return __r; }
		return QStackedWidget::metric(arg1);
	}
	void initPainter(QPainter* painter) const override {
		int __h = 0;
		if (__goOn[37] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6477500296980050210LL, (void*)painter, &__h); if (__h) return; }
		QStackedWidget::initPainter(painter);
	}
	QPaintDevice* redirected(QPoint* offset) const override {
		int __h = 0;
		if (__goOn[38] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 4127240633310655048LL, (void*)offset, &__h); if (__h) return (QPaintDevice*)__r; }
		return QStackedWidget::redirected(offset);
	}
	QPainter* sharedPainter() const override {
		int __h = 0;
		if (__goOn[39] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 5792771328741774991LL, &__h); if (__h) return (QPainter*)__r; }
		return QStackedWidget::sharedPainter();
	}
	void inputMethodEvent(QInputMethodEvent* arg1) override {
		int __h = 0;
		if (__goOn[40] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5438819661942843222LL, (void*)arg1, &__h); if (__h) return; }
		QStackedWidget::inputMethodEvent(arg1);
	}
	QVariant inputMethodQuery(Qt::InputMethodQuery arg1) const override {
		int __h = 0;
		if (__goOn[41] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7345853097340709638LL, (int)(arg1), &__h); if (__h) return *(QVariant*)__r; }
		return QStackedWidget::inputMethodQuery(arg1);
	}
	bool focusNextPrevChild(bool next) override {
		int __h = 0;
		if (__goOn[42] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 5898328517505784191LL, next, &__h); if (__h) return __r; }
		return QStackedWidget::focusNextPrevChild(next);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[43] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QStackedWidget::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[44] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QStackedWidget::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[45] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QStackedWidget::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[46] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QStackedWidget::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[47] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QStackedWidget::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[48] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QStackedWidget::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QStackedWidget;
static std::mutex __connsMu_QStackedWidget;

extern "C" {

void* QStackedWidget_base(void* self) { return (void*)static_cast<QFrame*>((QStackedWidget*)self); }
void QStackedWidget_chain(void* self, void** out) {
	QStackedWidget* __p = (QStackedWidget*)self;
	QFrame* __b0 = static_cast<QFrame*>(__p); out[0] = (void*)__b0;
	QWidget* __b1 = static_cast<QWidget*>(__b0); out[1] = (void*)__b1;
	QObject* __b2 = static_cast<QObject*>(__b1); out[2] = (void*)__b2;
}
void QStackedWidget__set_override(void* self, int idx) { ((QStackedWidget__Go*)self)->__goOn[idx] = 1; }
void* QStackedWidget_fromQObject(void* o) { return (void*)qobject_cast<QStackedWidget*>((QObject*)o); }
void* QStackedWidget_new(long long __goid, void* parent) {
	return new QStackedWidget__Go(__goid, (QWidget*)parent);
}
void* QStackedWidget_new2(long long __goid) {
	return new QStackedWidget__Go(__goid);
}
char* QStackedWidget_tr(const char* s, long long* __out_len) {
	auto&& __r = QStackedWidget::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QStackedWidget_addWidget(void* self, void* w) {
	auto&& __r = ((QStackedWidget*)self)->addWidget((QWidget*)w);
	return __r;
}
int QStackedWidget_insertWidget(void* self, int index, void* w) {
	auto&& __r = ((QStackedWidget*)self)->insertWidget(index, (QWidget*)w);
	return __r;
}
void QStackedWidget_removeWidget(void* self, void* w) {
	((QStackedWidget*)self)->removeWidget((QWidget*)w);
}
void* QStackedWidget_currentWidget(void* self) {
	auto&& __r = ((QStackedWidget*)self)->currentWidget();
	return (void*)(__r);
}
int QStackedWidget_currentIndex(void* self) {
	auto&& __r = ((QStackedWidget*)self)->currentIndex();
	return __r;
}
int QStackedWidget_indexOf(void* self, void* arg1) {
	auto&& __r = ((QStackedWidget*)self)->indexOf((QWidget*)arg1);
	return __r;
}
void* QStackedWidget_widget(void* self, int arg1) {
	auto&& __r = ((QStackedWidget*)self)->widget(arg1);
	return (void*)(__r);
}
int QStackedWidget_count(void* self) {
	auto&& __r = ((QStackedWidget*)self)->count();
	return __r;
}
void QStackedWidget_setCurrentIndex(void* self, int index) {
	((QStackedWidget*)self)->setCurrentIndex(index);
}
void QStackedWidget_setCurrentWidget(void* self, void* w) {
	((QStackedWidget*)self)->setCurrentWidget((QWidget*)w);
}
void QStackedWidget_currentChanged(void* self, int arg1) {
	((QStackedWidget*)self)->currentChanged(arg1);
}
void QStackedWidget_widgetRemoved(void* self, int index) {
	((QStackedWidget*)self)->widgetRemoved(index);
}
void QStackedWidget_widgetAdded(void* self, int index) {
	((QStackedWidget*)self)->widgetAdded(index);
}
void QStackedWidget_drawFrame(void* self, void* arg1) {
	((QStackedWidget__Go*)self)->drawFrame((QPainter*)arg1);
}
void QStackedWidget_setTabOrder(void* arg1, void* arg2) {
	QStackedWidget::setTabOrder((QWidget*)arg1, (QWidget*)arg2);
}
void* QStackedWidget_mouseGrabber() {
	auto&& __r = QStackedWidget::mouseGrabber();
	return (void*)(__r);
}
void* QStackedWidget_keyboardGrabber() {
	auto&& __r = QStackedWidget::keyboardGrabber();
	return (void*)(__r);
}
void* QStackedWidget_find(unsigned long long arg1) {
	auto&& __r = QStackedWidget::find(arg1);
	return (void*)(__r);
}
void* QStackedWidget_createWindowContainer(void* window) {
	auto&& __r = QStackedWidget::createWindowContainer((QWindow*)window);
	return (void*)(__r);
}
void QStackedWidget_updateMicroFocus(void* self) {
	((QStackedWidget__Go*)self)->updateMicroFocus();
}
void QStackedWidget_create(void* self) {
	((QStackedWidget__Go*)self)->create();
}
void QStackedWidget_destroy(void* self) {
	((QStackedWidget__Go*)self)->destroy();
}
int QStackedWidget_focusNextChild(void* self) {
	auto&& __r = ((QStackedWidget__Go*)self)->focusNextChild();
	return __r;
}
int QStackedWidget_focusPreviousChild(void* self) {
	auto&& __r = ((QStackedWidget__Go*)self)->focusPreviousChild();
	return __r;
}
void* QStackedWidget_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QStackedWidget::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QStackedWidget_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QStackedWidget::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QStackedWidget_sender(void* self) {
	auto&& __r = ((QStackedWidget__Go*)self)->sender();
	return (void*)(__r);
}
int QStackedWidget_senderSignalIndex(void* self) {
	auto&& __r = ((QStackedWidget__Go*)self)->senderSignalIndex();
	return __r;
}
int QStackedWidget_receivers(void* self, const char* signal) {
	auto&& __r = ((QStackedWidget__Go*)self)->receivers(signal);
	return __r;
}
int QStackedWidget_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QStackedWidget__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
double QStackedWidget_devicePixelRatioFScale() {
	auto&& __r = QStackedWidget::devicePixelRatioFScale();
	return __r;
}
int QStackedWidget_encodeMetricF(int metric, double value) {
	auto&& __r = QStackedWidget::encodeMetricF((QPaintDevice::PaintDeviceMetric)(metric), value);
	return __r;
}
double QStackedWidget_getDecodedMetricF(void* self, int metricA, int metricB) {
	auto&& __r = ((QStackedWidget__Go*)self)->getDecodedMetricF((QPaintDevice::PaintDeviceMetric)(metricA), (QPaintDevice::PaintDeviceMetric)(metricB));
	return __r;
}
char* QStackedWidget_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QStackedWidget::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QStackedWidget_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QStackedWidget::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QStackedWidget_createWindowContainer2(void* window, void* parent) {
	auto&& __r = QStackedWidget::createWindowContainer((QWindow*)window, (QWidget*)parent);
	return (void*)(__r);
}
void* QStackedWidget_createWindowContainer3(void* window, void* parent, int flags) {
	auto&& __r = QStackedWidget::createWindowContainer((QWindow*)window, (QWidget*)parent, QFlags<Qt::WindowType>(flags));
	return (void*)(__r);
}
void QStackedWidget_updateMicroFocus2(void* self, int query) {
	((QStackedWidget__Go*)self)->updateMicroFocus((Qt::InputMethodQuery)(query));
}
void QStackedWidget_create2(void* self, unsigned long long arg1) {
	((QStackedWidget__Go*)self)->create(arg1);
}
void QStackedWidget_create3(void* self, unsigned long long arg1, int initializeWindow) {
	((QStackedWidget__Go*)self)->create(arg1, initializeWindow);
}
void QStackedWidget_create4(void* self, unsigned long long arg1, int initializeWindow, int destroyOldWindow) {
	((QStackedWidget__Go*)self)->create(arg1, initializeWindow, destroyOldWindow);
}
void QStackedWidget_destroy2(void* self, int destroyWindow) {
	((QStackedWidget__Go*)self)->destroy(destroyWindow);
}
void QStackedWidget_destroy3(void* self, int destroyWindow, int destroySubWindows) {
	((QStackedWidget__Go*)self)->destroy(destroyWindow, destroySubWindows);
}
void* QStackedWidget_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QStackedWidget::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QStackedWidget_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QStackedWidget::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QStackedWidget_metaObject_default(void* self) {
	auto&& __r = ((QStackedWidget*)self)->QStackedWidget::metaObject();
	return (void*)(__r);
}
int QStackedWidget_event_default(void* self, void* e) {
	auto&& __r = ((QStackedWidget__Go*)self)->cppgoBase_event((QEvent*)e);
	return __r;
}
void* QStackedWidget_sizeHint_default(void* self) {
	auto&& __r = ((QStackedWidget*)self)->QStackedWidget::sizeHint();
	return new QSize(__r);
}
void QStackedWidget_paintEvent_default(void* self, void* arg1) {
	((QStackedWidget__Go*)self)->cppgoBase_paintEvent((QPaintEvent*)arg1);
}
void QStackedWidget_changeEvent_default(void* self, void* arg1) {
	((QStackedWidget__Go*)self)->cppgoBase_changeEvent((QEvent*)arg1);
}
void QStackedWidget_initStyleOption_default(void* self, void* option) {
	((QStackedWidget__Go*)self)->cppgoBase_initStyleOption((QStyleOptionFrame*)option);
}
int QStackedWidget_devType_default(void* self) {
	auto&& __r = ((QStackedWidget*)self)->QStackedWidget::devType();
	return __r;
}
void QStackedWidget_setVisible_default(void* self, int visible) {
	((QStackedWidget*)self)->QStackedWidget::setVisible(visible);
}
void* QStackedWidget_minimumSizeHint_default(void* self) {
	auto&& __r = ((QStackedWidget*)self)->QStackedWidget::minimumSizeHint();
	return new QSize(__r);
}
int QStackedWidget_heightForWidth_default(void* self, int arg1) {
	auto&& __r = ((QStackedWidget*)self)->QStackedWidget::heightForWidth(arg1);
	return __r;
}
int QStackedWidget_hasHeightForWidth_default(void* self) {
	auto&& __r = ((QStackedWidget*)self)->QStackedWidget::hasHeightForWidth();
	return __r;
}
void* QStackedWidget_paintEngine_default(void* self) {
	auto&& __r = ((QStackedWidget*)self)->QStackedWidget::paintEngine();
	return (void*)(__r);
}
void QStackedWidget_mousePressEvent_default(void* self, void* event) {
	((QStackedWidget__Go*)self)->cppgoBase_mousePressEvent((QMouseEvent*)event);
}
void QStackedWidget_mouseReleaseEvent_default(void* self, void* event) {
	((QStackedWidget__Go*)self)->cppgoBase_mouseReleaseEvent((QMouseEvent*)event);
}
void QStackedWidget_mouseDoubleClickEvent_default(void* self, void* event) {
	((QStackedWidget__Go*)self)->cppgoBase_mouseDoubleClickEvent((QMouseEvent*)event);
}
void QStackedWidget_mouseMoveEvent_default(void* self, void* event) {
	((QStackedWidget__Go*)self)->cppgoBase_mouseMoveEvent((QMouseEvent*)event);
}
void QStackedWidget_wheelEvent_default(void* self, void* event) {
	((QStackedWidget__Go*)self)->cppgoBase_wheelEvent((QWheelEvent*)event);
}
void QStackedWidget_keyPressEvent_default(void* self, void* event) {
	((QStackedWidget__Go*)self)->cppgoBase_keyPressEvent((QKeyEvent*)event);
}
void QStackedWidget_keyReleaseEvent_default(void* self, void* event) {
	((QStackedWidget__Go*)self)->cppgoBase_keyReleaseEvent((QKeyEvent*)event);
}
void QStackedWidget_focusInEvent_default(void* self, void* event) {
	((QStackedWidget__Go*)self)->cppgoBase_focusInEvent((QFocusEvent*)event);
}
void QStackedWidget_focusOutEvent_default(void* self, void* event) {
	((QStackedWidget__Go*)self)->cppgoBase_focusOutEvent((QFocusEvent*)event);
}
void QStackedWidget_enterEvent_default(void* self, void* event) {
	((QStackedWidget__Go*)self)->cppgoBase_enterEvent((QEnterEvent*)event);
}
void QStackedWidget_leaveEvent_default(void* self, void* event) {
	((QStackedWidget__Go*)self)->cppgoBase_leaveEvent((QEvent*)event);
}
void QStackedWidget_moveEvent_default(void* self, void* event) {
	((QStackedWidget__Go*)self)->cppgoBase_moveEvent((QMoveEvent*)event);
}
void QStackedWidget_resizeEvent_default(void* self, void* event) {
	((QStackedWidget__Go*)self)->cppgoBase_resizeEvent((QResizeEvent*)event);
}
void QStackedWidget_closeEvent_default(void* self, void* event) {
	((QStackedWidget__Go*)self)->cppgoBase_closeEvent((QCloseEvent*)event);
}
void QStackedWidget_contextMenuEvent_default(void* self, void* event) {
	((QStackedWidget__Go*)self)->cppgoBase_contextMenuEvent((QContextMenuEvent*)event);
}
void QStackedWidget_tabletEvent_default(void* self, void* event) {
	((QStackedWidget__Go*)self)->cppgoBase_tabletEvent((QTabletEvent*)event);
}
void QStackedWidget_actionEvent_default(void* self, void* event) {
	((QStackedWidget__Go*)self)->cppgoBase_actionEvent((QActionEvent*)event);
}
void QStackedWidget_dragEnterEvent_default(void* self, void* event) {
	((QStackedWidget__Go*)self)->cppgoBase_dragEnterEvent((QDragEnterEvent*)event);
}
void QStackedWidget_dragMoveEvent_default(void* self, void* event) {
	((QStackedWidget__Go*)self)->cppgoBase_dragMoveEvent((QDragMoveEvent*)event);
}
void QStackedWidget_dragLeaveEvent_default(void* self, void* event) {
	((QStackedWidget__Go*)self)->cppgoBase_dragLeaveEvent((QDragLeaveEvent*)event);
}
void QStackedWidget_dropEvent_default(void* self, void* event) {
	((QStackedWidget__Go*)self)->cppgoBase_dropEvent((QDropEvent*)event);
}
void QStackedWidget_showEvent_default(void* self, void* event) {
	((QStackedWidget__Go*)self)->cppgoBase_showEvent((QShowEvent*)event);
}
void QStackedWidget_hideEvent_default(void* self, void* event) {
	((QStackedWidget__Go*)self)->cppgoBase_hideEvent((QHideEvent*)event);
}
int QStackedWidget_metric_default(void* self, int arg1) {
	auto&& __r = ((QStackedWidget__Go*)self)->cppgoBase_metric(arg1);
	return __r;
}
void QStackedWidget_initPainter_default(void* self, void* painter) {
	((QStackedWidget__Go*)self)->cppgoBase_initPainter((QPainter*)painter);
}
void* QStackedWidget_redirected_default(void* self, void* offset) {
	auto&& __r = ((QStackedWidget__Go*)self)->cppgoBase_redirected((QPoint*)offset);
	return (void*)(__r);
}
void* QStackedWidget_sharedPainter_default(void* self) {
	auto&& __r = ((QStackedWidget__Go*)self)->cppgoBase_sharedPainter();
	return (void*)(__r);
}
void QStackedWidget_inputMethodEvent_default(void* self, void* arg1) {
	((QStackedWidget__Go*)self)->cppgoBase_inputMethodEvent((QInputMethodEvent*)arg1);
}
void* QStackedWidget_inputMethodQuery_default(void* self, int arg1) {
	auto&& __r = ((QStackedWidget*)self)->QStackedWidget::inputMethodQuery((Qt::InputMethodQuery)(arg1));
	return new QVariant(__r);
}
int QStackedWidget_focusNextPrevChild_default(void* self, int next) {
	auto&& __r = ((QStackedWidget__Go*)self)->cppgoBase_focusNextPrevChild(next);
	return __r;
}
int QStackedWidget_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QStackedWidget*)self)->QStackedWidget::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QStackedWidget_timerEvent_default(void* self, void* event) {
	((QStackedWidget__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QStackedWidget_childEvent_default(void* self, void* event) {
	((QStackedWidget__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QStackedWidget_customEvent_default(void* self, void* event) {
	((QStackedWidget__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QStackedWidget_connectNotify_default(void* self, void* signal) {
	((QStackedWidget__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QStackedWidget_disconnectNotify_default(void* self, void* signal) {
	((QStackedWidget__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QStackedWidget_connect_currentChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStackedWidget*)self, &QStackedWidget::currentChanged, (QStackedWidget*)self, [__id](int arg1){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStackedWidget);
	__conns_QStackedWidget[__id] = __c;
}
void QStackedWidget_connect_widgetRemoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStackedWidget*)self, &QStackedWidget::widgetRemoved, (QStackedWidget*)self, [__id](int index){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, index);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStackedWidget);
	__conns_QStackedWidget[__id] = __c;
}
void QStackedWidget_connect_widgetAdded(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStackedWidget*)self, &QStackedWidget::widgetAdded, (QStackedWidget*)self, [__id](int index){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, index);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStackedWidget);
	__conns_QStackedWidget[__id] = __c;
}
void QStackedWidget_connect_windowTitleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStackedWidget*)self, &QStackedWidget::windowTitleChanged, (QStackedWidget*)self, [__id](const QString& title){
		QByteArray title_ba = title.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)title_ba.constData(), title_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStackedWidget);
	__conns_QStackedWidget[__id] = __c;
}
void QStackedWidget_connect_windowIconChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStackedWidget*)self, &QStackedWidget::windowIconChanged, (QStackedWidget*)self, [__id](const QIcon& icon){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&icon);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStackedWidget);
	__conns_QStackedWidget[__id] = __c;
}
void QStackedWidget_connect_windowIconTextChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStackedWidget*)self, &QStackedWidget::windowIconTextChanged, (QStackedWidget*)self, [__id](const QString& iconText){
		QByteArray iconText_ba = iconText.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)iconText_ba.constData(), iconText_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStackedWidget);
	__conns_QStackedWidget[__id] = __c;
}
void QStackedWidget_connect_customContextMenuRequested(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStackedWidget*)self, &QStackedWidget::customContextMenuRequested, (QStackedWidget*)self, [__id](const QPoint& pos){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&pos);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStackedWidget);
	__conns_QStackedWidget[__id] = __c;
}
void QStackedWidget_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStackedWidget*)self, &QStackedWidget::destroyed, (QStackedWidget*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStackedWidget);
	__conns_QStackedWidget[__id] = __c;
}
void QStackedWidget_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStackedWidget*)self, &QStackedWidget::destroyed, (QStackedWidget*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStackedWidget);
	__conns_QStackedWidget[__id] = __c;
}
void QStackedWidget_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStackedWidget*)self, &QStackedWidget::objectNameChanged, (QStackedWidget*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStackedWidget);
	__conns_QStackedWidget[__id] = __c;
}
void QStackedWidget__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QStackedWidget);
		auto __it = __conns_QStackedWidget.find(id);
		if (__it == __conns_QStackedWidget.end()) return;
		__c = __it->second;
		__conns_QStackedWidget.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QStackedWidget__Go_new(long long id) { return new QStackedWidget__Go(id); }
void QStackedWidget_delete(void* self) { delete (QStackedWidget*)self; }

}

#include <QtWidgets/qwizard.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtGui/qpixmap.h>
#include <QtCore/qvariant.h>
#include <QtWidgets/qstyle.h>
#include <QtCore/qnamespace.h>
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
#include <QtGui/qpainter.h>
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
#include <QtCore/qcoreevent.h>
#include <QtGui/qevent.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qwizardpage.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QWizardPage__Go : public QWizardPage {
public:
	long long __goId;
	unsigned char __goOn[53] = {};
	template<typename... A> QWizardPage__Go(long long id, A&&... a) : QWizardPage(std::forward<A>(a)...), __goId(id) {}
	using QWizardPage::setField;
	using QWizardPage::field;
	using QWizardPage::registerField;
	using QWizardPage::wizard;
	using QWizardPage::updateMicroFocus;
	using QWizardPage::create;
	using QWizardPage::destroy;
	using QWizardPage::focusNextChild;
	using QWizardPage::focusPreviousChild;
	using QWizardPage::sender;
	using QWizardPage::senderSignalIndex;
	using QWizardPage::receivers;
	using QWizardPage::isSignalConnected;
	using QWizardPage::getDecodedMetricF;
	bool cppgoBase_event(QEvent* __a0) { return QWizardPage::event(__a0); }
	void cppgoBase_mousePressEvent(QMouseEvent* __a0) { QWizardPage::mousePressEvent(__a0); }
	void cppgoBase_mouseReleaseEvent(QMouseEvent* __a0) { QWizardPage::mouseReleaseEvent(__a0); }
	void cppgoBase_mouseDoubleClickEvent(QMouseEvent* __a0) { QWizardPage::mouseDoubleClickEvent(__a0); }
	void cppgoBase_mouseMoveEvent(QMouseEvent* __a0) { QWizardPage::mouseMoveEvent(__a0); }
	void cppgoBase_wheelEvent(QWheelEvent* __a0) { QWizardPage::wheelEvent(__a0); }
	void cppgoBase_keyPressEvent(QKeyEvent* __a0) { QWizardPage::keyPressEvent(__a0); }
	void cppgoBase_keyReleaseEvent(QKeyEvent* __a0) { QWizardPage::keyReleaseEvent(__a0); }
	void cppgoBase_focusInEvent(QFocusEvent* __a0) { QWizardPage::focusInEvent(__a0); }
	void cppgoBase_focusOutEvent(QFocusEvent* __a0) { QWizardPage::focusOutEvent(__a0); }
	void cppgoBase_enterEvent(QEnterEvent* __a0) { QWizardPage::enterEvent(__a0); }
	void cppgoBase_leaveEvent(QEvent* __a0) { QWizardPage::leaveEvent(__a0); }
	void cppgoBase_paintEvent(QPaintEvent* __a0) { QWizardPage::paintEvent(__a0); }
	void cppgoBase_moveEvent(QMoveEvent* __a0) { QWizardPage::moveEvent(__a0); }
	void cppgoBase_resizeEvent(QResizeEvent* __a0) { QWizardPage::resizeEvent(__a0); }
	void cppgoBase_closeEvent(QCloseEvent* __a0) { QWizardPage::closeEvent(__a0); }
	void cppgoBase_contextMenuEvent(QContextMenuEvent* __a0) { QWizardPage::contextMenuEvent(__a0); }
	void cppgoBase_tabletEvent(QTabletEvent* __a0) { QWizardPage::tabletEvent(__a0); }
	void cppgoBase_actionEvent(QActionEvent* __a0) { QWizardPage::actionEvent(__a0); }
	void cppgoBase_dragEnterEvent(QDragEnterEvent* __a0) { QWizardPage::dragEnterEvent(__a0); }
	void cppgoBase_dragMoveEvent(QDragMoveEvent* __a0) { QWizardPage::dragMoveEvent(__a0); }
	void cppgoBase_dragLeaveEvent(QDragLeaveEvent* __a0) { QWizardPage::dragLeaveEvent(__a0); }
	void cppgoBase_dropEvent(QDropEvent* __a0) { QWizardPage::dropEvent(__a0); }
	void cppgoBase_showEvent(QShowEvent* __a0) { QWizardPage::showEvent(__a0); }
	void cppgoBase_hideEvent(QHideEvent* __a0) { QWizardPage::hideEvent(__a0); }
	void cppgoBase_changeEvent(QEvent* __a0) { QWizardPage::changeEvent(__a0); }
	int cppgoBase_metric(int __a0) { return QWizardPage::metric((QPaintDevice::PaintDeviceMetric)__a0); }
	void cppgoBase_initPainter(QPainter* __a0) { QWizardPage::initPainter(__a0); }
	QPaintDevice* cppgoBase_redirected(QPoint* __a0) { return QWizardPage::redirected(__a0); }
	QPainter* cppgoBase_sharedPainter() { return QWizardPage::sharedPainter(); }
	void cppgoBase_inputMethodEvent(QInputMethodEvent* __a0) { QWizardPage::inputMethodEvent(__a0); }
	bool cppgoBase_focusNextPrevChild(bool __a0) { return QWizardPage::focusNextPrevChild(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QWizardPage::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QWizardPage::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QWizardPage::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QWizardPage::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QWizardPage::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QWizardPage::metaObject();
	}
	void initializePage() override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 3074997362940195296LL, &__h); if (__h) return; }
		QWizardPage::initializePage();
	}
	void cleanupPage() override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 5913490871578514792LL, &__h); if (__h) return; }
		QWizardPage::cleanupPage();
	}
	bool validatePage() override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -4021270481451057600LL, &__h); if (__h) return __r; }
		return QWizardPage::validatePage();
	}
	bool isComplete() const override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, 9144100205243375250LL, &__h); if (__h) return __r; }
		return QWizardPage::isComplete();
	}
	int nextId() const override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, 5814903794109731573LL, &__h); if (__h) return __r; }
		return QWizardPage::nextId();
	}
	int devType() const override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8659679893868524044LL, &__h); if (__h) return __r; }
		return QWizardPage::devType();
	}
	void setVisible(bool visible) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 2110182719375245575LL, visible, &__h); if (__h) return; }
		QWizardPage::setVisible(visible);
	}
	QSize sizeHint() const override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 729690261457684101LL, &__h); if (__h) return *(QSize*)__r; }
		return QWizardPage::sizeHint();
	}
	QSize minimumSizeHint() const override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 7049264599837532491LL, &__h); if (__h) return *(QSize*)__r; }
		return QWizardPage::minimumSizeHint();
	}
	int heightForWidth(int arg1) const override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 4311534060495737339LL, arg1, &__h); if (__h) return __r; }
		return QWizardPage::heightForWidth(arg1);
	}
	bool hasHeightForWidth() const override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8266762397995423915LL, &__h); if (__h) return __r; }
		return QWizardPage::hasHeightForWidth();
	}
	QPaintEngine* paintEngine() const override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 6398007404756346381LL, &__h); if (__h) return (QPaintEngine*)__r; }
		return QWizardPage::paintEngine();
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QWizardPage::event(event);
	}
	void mousePressEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -332090898078543319LL, (void*)event, &__h); if (__h) return; }
		QWizardPage::mousePressEvent(event);
	}
	void mouseReleaseEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4415920049297795213LL, (void*)event, &__h); if (__h) return; }
		QWizardPage::mouseReleaseEvent(event);
	}
	void mouseDoubleClickEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4029073905697023887LL, (void*)event, &__h); if (__h) return; }
		QWizardPage::mouseDoubleClickEvent(event);
	}
	void mouseMoveEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4255494004099806581LL, (void*)event, &__h); if (__h) return; }
		QWizardPage::mouseMoveEvent(event);
	}
	void wheelEvent(QWheelEvent* event) override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1919988340934511316LL, (void*)event, &__h); if (__h) return; }
		QWizardPage::wheelEvent(event);
	}
	void keyPressEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1023791977679793185LL, (void*)event, &__h); if (__h) return; }
		QWizardPage::keyPressEvent(event);
	}
	void keyReleaseEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 772598475950050299LL, (void*)event, &__h); if (__h) return; }
		QWizardPage::keyReleaseEvent(event);
	}
	void focusInEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 975043767033552552LL, (void*)event, &__h); if (__h) return; }
		QWizardPage::focusInEvent(event);
	}
	void focusOutEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -385596010576048993LL, (void*)event, &__h); if (__h) return; }
		QWizardPage::focusOutEvent(event);
	}
	void enterEvent(QEnterEvent* event) override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6316270346668237895LL, (void*)event, &__h); if (__h) return; }
		QWizardPage::enterEvent(event);
	}
	void leaveEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -551070009623711640LL, (void*)event, &__h); if (__h) return; }
		QWizardPage::leaveEvent(event);
	}
	void paintEvent(QPaintEvent* event) override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5290442163167458465LL, (void*)event, &__h); if (__h) return; }
		QWizardPage::paintEvent(event);
	}
	void moveEvent(QMoveEvent* event) override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -9161949150489525396LL, (void*)event, &__h); if (__h) return; }
		QWizardPage::moveEvent(event);
	}
	void resizeEvent(QResizeEvent* event) override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4302917285655014661LL, (void*)event, &__h); if (__h) return; }
		QWizardPage::resizeEvent(event);
	}
	void closeEvent(QCloseEvent* event) override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1989937050019464783LL, (void*)event, &__h); if (__h) return; }
		QWizardPage::closeEvent(event);
	}
	void contextMenuEvent(QContextMenuEvent* event) override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6698043222227006015LL, (void*)event, &__h); if (__h) return; }
		QWizardPage::contextMenuEvent(event);
	}
	void tabletEvent(QTabletEvent* event) override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -212068866860701141LL, (void*)event, &__h); if (__h) return; }
		QWizardPage::tabletEvent(event);
	}
	void actionEvent(QActionEvent* event) override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1697444140088555581LL, (void*)event, &__h); if (__h) return; }
		QWizardPage::actionEvent(event);
	}
	void dragEnterEvent(QDragEnterEvent* event) override {
		int __h = 0;
		if (__goOn[32] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 206122715919917589LL, (void*)event, &__h); if (__h) return; }
		QWizardPage::dragEnterEvent(event);
	}
	void dragMoveEvent(QDragMoveEvent* event) override {
		int __h = 0;
		if (__goOn[33] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1781545137352745978LL, (void*)event, &__h); if (__h) return; }
		QWizardPage::dragMoveEvent(event);
	}
	void dragLeaveEvent(QDragLeaveEvent* event) override {
		int __h = 0;
		if (__goOn[34] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -3139904158373725006LL, (void*)event, &__h); if (__h) return; }
		QWizardPage::dragLeaveEvent(event);
	}
	void dropEvent(QDropEvent* event) override {
		int __h = 0;
		if (__goOn[35] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 9107058682754591716LL, (void*)event, &__h); if (__h) return; }
		QWizardPage::dropEvent(event);
	}
	void showEvent(QShowEvent* event) override {
		int __h = 0;
		if (__goOn[36] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6825369174751916LL, (void*)event, &__h); if (__h) return; }
		QWizardPage::showEvent(event);
	}
	void hideEvent(QHideEvent* event) override {
		int __h = 0;
		if (__goOn[37] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -199718707876719129LL, (void*)event, &__h); if (__h) return; }
		QWizardPage::hideEvent(event);
	}
	bool nativeEvent(const QByteArray& eventType, void* message, qintptr* result) override {
		int __h = 0;
		if (__goOn[38] && cppgo_virt_int_voidp_voidp_longlongp) { int __r = ((int(*)(long long, long long, void*, void*, long long*, int*))cppgo_virt_int_voidp_voidp_longlongp)(__goId, -2039166586246443706LL, (void*)&eventType, (void*)message, (long long*)result, &__h); if (__h) return __r; }
		return QWizardPage::nativeEvent(eventType, message, result);
	}
	void changeEvent(QEvent* arg1) override {
		int __h = 0;
		if (__goOn[39] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4959376063826165943LL, (void*)arg1, &__h); if (__h) return; }
		QWizardPage::changeEvent(arg1);
	}
	int metric(QPaintDevice::PaintDeviceMetric arg1) const override {
		int __h = 0;
		if (__goOn[40] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 427629649267311777LL, (int)(arg1), &__h); if (__h) return __r; }
		return QWizardPage::metric(arg1);
	}
	void initPainter(QPainter* painter) const override {
		int __h = 0;
		if (__goOn[41] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6477500296980050210LL, (void*)painter, &__h); if (__h) return; }
		QWizardPage::initPainter(painter);
	}
	QPaintDevice* redirected(QPoint* offset) const override {
		int __h = 0;
		if (__goOn[42] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 4127240633310655048LL, (void*)offset, &__h); if (__h) return (QPaintDevice*)__r; }
		return QWizardPage::redirected(offset);
	}
	QPainter* sharedPainter() const override {
		int __h = 0;
		if (__goOn[43] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 5792771328741774991LL, &__h); if (__h) return (QPainter*)__r; }
		return QWizardPage::sharedPainter();
	}
	void inputMethodEvent(QInputMethodEvent* arg1) override {
		int __h = 0;
		if (__goOn[44] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5438819661942843222LL, (void*)arg1, &__h); if (__h) return; }
		QWizardPage::inputMethodEvent(arg1);
	}
	QVariant inputMethodQuery(Qt::InputMethodQuery arg1) const override {
		int __h = 0;
		if (__goOn[45] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7345853097340709638LL, (int)(arg1), &__h); if (__h) return *(QVariant*)__r; }
		return QWizardPage::inputMethodQuery(arg1);
	}
	bool focusNextPrevChild(bool next) override {
		int __h = 0;
		if (__goOn[46] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 5898328517505784191LL, next, &__h); if (__h) return __r; }
		return QWizardPage::focusNextPrevChild(next);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[47] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QWizardPage::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[48] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QWizardPage::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[49] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QWizardPage::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[50] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QWizardPage::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[51] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QWizardPage::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[52] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QWizardPage::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QWizardPage;
static std::mutex __connsMu_QWizardPage;

extern "C" {

void* QWizardPage_base(void* self) { return (void*)static_cast<QWidget*>((QWizardPage*)self); }
void QWizardPage_chain(void* self, void** out) {
	QWizardPage* __p = (QWizardPage*)self;
	QWidget* __b0 = static_cast<QWidget*>(__p); out[0] = (void*)__b0;
	QObject* __b1 = static_cast<QObject*>(__b0); out[1] = (void*)__b1;
}
void QWizardPage__set_override(void* self, int idx) { ((QWizardPage__Go*)self)->__goOn[idx] = 1; }
void* QWizardPage_fromQObject(void* o) { return (void*)qobject_cast<QWizardPage*>((QObject*)o); }
void* QWizardPage_new(long long __goid, void* parent) {
	return new QWizardPage__Go(__goid, (QWidget*)parent);
}
void* QWizardPage_new2(long long __goid) {
	return new QWizardPage__Go(__goid);
}
char* QWizardPage_tr(const char* s, long long* __out_len) {
	auto&& __r = QWizardPage::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QWizardPage_setTitle(void* self, const char* title, long long title_len) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	((QWizardPage*)self)->setTitle(title_qstr);
}
char* QWizardPage_title(void* self, long long* __out_len) {
	auto&& __r = ((QWizardPage*)self)->title();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QWizardPage_setSubTitle(void* self, const char* subTitle, long long subTitle_len) {
	QString subTitle_qstr = QString::fromUtf8(subTitle, subTitle_len);
	((QWizardPage*)self)->setSubTitle(subTitle_qstr);
}
char* QWizardPage_subTitle(void* self, long long* __out_len) {
	auto&& __r = ((QWizardPage*)self)->subTitle();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QWizardPage_setPixmap(void* self, int which, void* pixmap) {
	((QWizardPage*)self)->setPixmap((QWizard::WizardPixmap)(which), *(QPixmap*)pixmap);
}
void* QWizardPage_pixmap(void* self, int which) {
	auto&& __r = ((QWizardPage*)self)->pixmap((QWizard::WizardPixmap)(which));
	return new QPixmap(__r);
}
void QWizardPage_setFinalPage(void* self, int finalPage) {
	((QWizardPage*)self)->setFinalPage(finalPage);
}
int QWizardPage_isFinalPage(void* self) {
	auto&& __r = ((QWizardPage*)self)->isFinalPage();
	return __r;
}
void QWizardPage_setCommitPage(void* self, int commitPage) {
	((QWizardPage*)self)->setCommitPage(commitPage);
}
int QWizardPage_isCommitPage(void* self) {
	auto&& __r = ((QWizardPage*)self)->isCommitPage();
	return __r;
}
void QWizardPage_setButtonText(void* self, int which, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	((QWizardPage*)self)->setButtonText((QWizard::WizardButton)(which), text_qstr);
}
char* QWizardPage_buttonText(void* self, int which, long long* __out_len) {
	auto&& __r = ((QWizardPage*)self)->buttonText((QWizard::WizardButton)(which));
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QWizardPage_initializePage(void* self) {
	((QWizardPage*)self)->initializePage();
}
void QWizardPage_cleanupPage(void* self) {
	((QWizardPage*)self)->cleanupPage();
}
int QWizardPage_validatePage(void* self) {
	auto&& __r = ((QWizardPage*)self)->validatePage();
	return __r;
}
int QWizardPage_isComplete(void* self) {
	auto&& __r = ((QWizardPage*)self)->isComplete();
	return __r;
}
int QWizardPage_nextId(void* self) {
	auto&& __r = ((QWizardPage*)self)->nextId();
	return __r;
}
void QWizardPage_completeChanged(void* self) {
	((QWizardPage*)self)->completeChanged();
}
void QWizardPage_setField(void* self, const char* name, long long name_len, void* value) {
	QString name_qstr = QString::fromUtf8(name, name_len);
	((QWizardPage__Go*)self)->setField(name_qstr, *(QVariant*)value);
}
void* QWizardPage_field(void* self, const char* name, long long name_len) {
	QString name_qstr = QString::fromUtf8(name, name_len);
	auto&& __r = ((QWizardPage__Go*)self)->field(name_qstr);
	return new QVariant(__r);
}
void QWizardPage_registerField(void* self, const char* name, long long name_len, void* widget) {
	QString name_qstr = QString::fromUtf8(name, name_len);
	((QWizardPage__Go*)self)->registerField(name_qstr, (QWidget*)widget);
}
void* QWizardPage_wizard(void* self) {
	auto&& __r = ((QWizardPage__Go*)self)->wizard();
	return (void*)(__r);
}
void QWizardPage_setTabOrder(void* arg1, void* arg2) {
	QWizardPage::setTabOrder((QWidget*)arg1, (QWidget*)arg2);
}
void* QWizardPage_mouseGrabber() {
	auto&& __r = QWizardPage::mouseGrabber();
	return (void*)(__r);
}
void* QWizardPage_keyboardGrabber() {
	auto&& __r = QWizardPage::keyboardGrabber();
	return (void*)(__r);
}
void* QWizardPage_find(unsigned long long arg1) {
	auto&& __r = QWizardPage::find(arg1);
	return (void*)(__r);
}
void* QWizardPage_createWindowContainer(void* window) {
	auto&& __r = QWizardPage::createWindowContainer((QWindow*)window);
	return (void*)(__r);
}
void QWizardPage_updateMicroFocus(void* self) {
	((QWizardPage__Go*)self)->updateMicroFocus();
}
void QWizardPage_create(void* self) {
	((QWizardPage__Go*)self)->create();
}
void QWizardPage_destroy(void* self) {
	((QWizardPage__Go*)self)->destroy();
}
int QWizardPage_focusNextChild(void* self) {
	auto&& __r = ((QWizardPage__Go*)self)->focusNextChild();
	return __r;
}
int QWizardPage_focusPreviousChild(void* self) {
	auto&& __r = ((QWizardPage__Go*)self)->focusPreviousChild();
	return __r;
}
void* QWizardPage_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QWizardPage::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QWizardPage_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QWizardPage::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QWizardPage_sender(void* self) {
	auto&& __r = ((QWizardPage__Go*)self)->sender();
	return (void*)(__r);
}
int QWizardPage_senderSignalIndex(void* self) {
	auto&& __r = ((QWizardPage__Go*)self)->senderSignalIndex();
	return __r;
}
int QWizardPage_receivers(void* self, const char* signal) {
	auto&& __r = ((QWizardPage__Go*)self)->receivers(signal);
	return __r;
}
int QWizardPage_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QWizardPage__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
double QWizardPage_devicePixelRatioFScale() {
	auto&& __r = QWizardPage::devicePixelRatioFScale();
	return __r;
}
int QWizardPage_encodeMetricF(int metric, double value) {
	auto&& __r = QWizardPage::encodeMetricF((QPaintDevice::PaintDeviceMetric)(metric), value);
	return __r;
}
double QWizardPage_getDecodedMetricF(void* self, int metricA, int metricB) {
	auto&& __r = ((QWizardPage__Go*)self)->getDecodedMetricF((QPaintDevice::PaintDeviceMetric)(metricA), (QPaintDevice::PaintDeviceMetric)(metricB));
	return __r;
}
char* QWizardPage_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QWizardPage::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QWizardPage_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QWizardPage::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QWizardPage_registerField2(void* self, const char* name, long long name_len, void* widget, const char* property) {
	QString name_qstr = QString::fromUtf8(name, name_len);
	((QWizardPage__Go*)self)->registerField(name_qstr, (QWidget*)widget, property);
}
void QWizardPage_registerField3(void* self, const char* name, long long name_len, void* widget, const char* property, const char* changedSignal) {
	QString name_qstr = QString::fromUtf8(name, name_len);
	((QWizardPage__Go*)self)->registerField(name_qstr, (QWidget*)widget, property, changedSignal);
}
void* QWizardPage_createWindowContainer2(void* window, void* parent) {
	auto&& __r = QWizardPage::createWindowContainer((QWindow*)window, (QWidget*)parent);
	return (void*)(__r);
}
void* QWizardPage_createWindowContainer3(void* window, void* parent, int flags) {
	auto&& __r = QWizardPage::createWindowContainer((QWindow*)window, (QWidget*)parent, QFlags<Qt::WindowType>(flags));
	return (void*)(__r);
}
void QWizardPage_updateMicroFocus2(void* self, int query) {
	((QWizardPage__Go*)self)->updateMicroFocus((Qt::InputMethodQuery)(query));
}
void QWizardPage_create2(void* self, unsigned long long arg1) {
	((QWizardPage__Go*)self)->create(arg1);
}
void QWizardPage_create3(void* self, unsigned long long arg1, int initializeWindow) {
	((QWizardPage__Go*)self)->create(arg1, initializeWindow);
}
void QWizardPage_create4(void* self, unsigned long long arg1, int initializeWindow, int destroyOldWindow) {
	((QWizardPage__Go*)self)->create(arg1, initializeWindow, destroyOldWindow);
}
void QWizardPage_destroy2(void* self, int destroyWindow) {
	((QWizardPage__Go*)self)->destroy(destroyWindow);
}
void QWizardPage_destroy3(void* self, int destroyWindow, int destroySubWindows) {
	((QWizardPage__Go*)self)->destroy(destroyWindow, destroySubWindows);
}
void* QWizardPage_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QWizardPage::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QWizardPage_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QWizardPage::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void QWizardPage_initializePage_default(void* self) {
	((QWizardPage*)self)->QWizardPage::initializePage();
}
void QWizardPage_cleanupPage_default(void* self) {
	((QWizardPage*)self)->QWizardPage::cleanupPage();
}
int QWizardPage_validatePage_default(void* self) {
	auto&& __r = ((QWizardPage*)self)->QWizardPage::validatePage();
	return __r;
}
int QWizardPage_isComplete_default(void* self) {
	auto&& __r = ((QWizardPage*)self)->QWizardPage::isComplete();
	return __r;
}
int QWizardPage_nextId_default(void* self) {
	auto&& __r = ((QWizardPage*)self)->QWizardPage::nextId();
	return __r;
}
void* QWizardPage_metaObject_default(void* self) {
	auto&& __r = ((QWizardPage*)self)->QWizardPage::metaObject();
	return (void*)(__r);
}
int QWizardPage_devType_default(void* self) {
	auto&& __r = ((QWizardPage*)self)->QWizardPage::devType();
	return __r;
}
void QWizardPage_setVisible_default(void* self, int visible) {
	((QWizardPage*)self)->QWizardPage::setVisible(visible);
}
void* QWizardPage_sizeHint_default(void* self) {
	auto&& __r = ((QWizardPage*)self)->QWizardPage::sizeHint();
	return new QSize(__r);
}
void* QWizardPage_minimumSizeHint_default(void* self) {
	auto&& __r = ((QWizardPage*)self)->QWizardPage::minimumSizeHint();
	return new QSize(__r);
}
int QWizardPage_heightForWidth_default(void* self, int arg1) {
	auto&& __r = ((QWizardPage*)self)->QWizardPage::heightForWidth(arg1);
	return __r;
}
int QWizardPage_hasHeightForWidth_default(void* self) {
	auto&& __r = ((QWizardPage*)self)->QWizardPage::hasHeightForWidth();
	return __r;
}
void* QWizardPage_paintEngine_default(void* self) {
	auto&& __r = ((QWizardPage*)self)->QWizardPage::paintEngine();
	return (void*)(__r);
}
int QWizardPage_event_default(void* self, void* event) {
	auto&& __r = ((QWizardPage__Go*)self)->cppgoBase_event((QEvent*)event);
	return __r;
}
void QWizardPage_mousePressEvent_default(void* self, void* event) {
	((QWizardPage__Go*)self)->cppgoBase_mousePressEvent((QMouseEvent*)event);
}
void QWizardPage_mouseReleaseEvent_default(void* self, void* event) {
	((QWizardPage__Go*)self)->cppgoBase_mouseReleaseEvent((QMouseEvent*)event);
}
void QWizardPage_mouseDoubleClickEvent_default(void* self, void* event) {
	((QWizardPage__Go*)self)->cppgoBase_mouseDoubleClickEvent((QMouseEvent*)event);
}
void QWizardPage_mouseMoveEvent_default(void* self, void* event) {
	((QWizardPage__Go*)self)->cppgoBase_mouseMoveEvent((QMouseEvent*)event);
}
void QWizardPage_wheelEvent_default(void* self, void* event) {
	((QWizardPage__Go*)self)->cppgoBase_wheelEvent((QWheelEvent*)event);
}
void QWizardPage_keyPressEvent_default(void* self, void* event) {
	((QWizardPage__Go*)self)->cppgoBase_keyPressEvent((QKeyEvent*)event);
}
void QWizardPage_keyReleaseEvent_default(void* self, void* event) {
	((QWizardPage__Go*)self)->cppgoBase_keyReleaseEvent((QKeyEvent*)event);
}
void QWizardPage_focusInEvent_default(void* self, void* event) {
	((QWizardPage__Go*)self)->cppgoBase_focusInEvent((QFocusEvent*)event);
}
void QWizardPage_focusOutEvent_default(void* self, void* event) {
	((QWizardPage__Go*)self)->cppgoBase_focusOutEvent((QFocusEvent*)event);
}
void QWizardPage_enterEvent_default(void* self, void* event) {
	((QWizardPage__Go*)self)->cppgoBase_enterEvent((QEnterEvent*)event);
}
void QWizardPage_leaveEvent_default(void* self, void* event) {
	((QWizardPage__Go*)self)->cppgoBase_leaveEvent((QEvent*)event);
}
void QWizardPage_paintEvent_default(void* self, void* event) {
	((QWizardPage__Go*)self)->cppgoBase_paintEvent((QPaintEvent*)event);
}
void QWizardPage_moveEvent_default(void* self, void* event) {
	((QWizardPage__Go*)self)->cppgoBase_moveEvent((QMoveEvent*)event);
}
void QWizardPage_resizeEvent_default(void* self, void* event) {
	((QWizardPage__Go*)self)->cppgoBase_resizeEvent((QResizeEvent*)event);
}
void QWizardPage_closeEvent_default(void* self, void* event) {
	((QWizardPage__Go*)self)->cppgoBase_closeEvent((QCloseEvent*)event);
}
void QWizardPage_contextMenuEvent_default(void* self, void* event) {
	((QWizardPage__Go*)self)->cppgoBase_contextMenuEvent((QContextMenuEvent*)event);
}
void QWizardPage_tabletEvent_default(void* self, void* event) {
	((QWizardPage__Go*)self)->cppgoBase_tabletEvent((QTabletEvent*)event);
}
void QWizardPage_actionEvent_default(void* self, void* event) {
	((QWizardPage__Go*)self)->cppgoBase_actionEvent((QActionEvent*)event);
}
void QWizardPage_dragEnterEvent_default(void* self, void* event) {
	((QWizardPage__Go*)self)->cppgoBase_dragEnterEvent((QDragEnterEvent*)event);
}
void QWizardPage_dragMoveEvent_default(void* self, void* event) {
	((QWizardPage__Go*)self)->cppgoBase_dragMoveEvent((QDragMoveEvent*)event);
}
void QWizardPage_dragLeaveEvent_default(void* self, void* event) {
	((QWizardPage__Go*)self)->cppgoBase_dragLeaveEvent((QDragLeaveEvent*)event);
}
void QWizardPage_dropEvent_default(void* self, void* event) {
	((QWizardPage__Go*)self)->cppgoBase_dropEvent((QDropEvent*)event);
}
void QWizardPage_showEvent_default(void* self, void* event) {
	((QWizardPage__Go*)self)->cppgoBase_showEvent((QShowEvent*)event);
}
void QWizardPage_hideEvent_default(void* self, void* event) {
	((QWizardPage__Go*)self)->cppgoBase_hideEvent((QHideEvent*)event);
}
void QWizardPage_changeEvent_default(void* self, void* arg1) {
	((QWizardPage__Go*)self)->cppgoBase_changeEvent((QEvent*)arg1);
}
int QWizardPage_metric_default(void* self, int arg1) {
	auto&& __r = ((QWizardPage__Go*)self)->cppgoBase_metric(arg1);
	return __r;
}
void QWizardPage_initPainter_default(void* self, void* painter) {
	((QWizardPage__Go*)self)->cppgoBase_initPainter((QPainter*)painter);
}
void* QWizardPage_redirected_default(void* self, void* offset) {
	auto&& __r = ((QWizardPage__Go*)self)->cppgoBase_redirected((QPoint*)offset);
	return (void*)(__r);
}
void* QWizardPage_sharedPainter_default(void* self) {
	auto&& __r = ((QWizardPage__Go*)self)->cppgoBase_sharedPainter();
	return (void*)(__r);
}
void QWizardPage_inputMethodEvent_default(void* self, void* arg1) {
	((QWizardPage__Go*)self)->cppgoBase_inputMethodEvent((QInputMethodEvent*)arg1);
}
void* QWizardPage_inputMethodQuery_default(void* self, int arg1) {
	auto&& __r = ((QWizardPage*)self)->QWizardPage::inputMethodQuery((Qt::InputMethodQuery)(arg1));
	return new QVariant(__r);
}
int QWizardPage_focusNextPrevChild_default(void* self, int next) {
	auto&& __r = ((QWizardPage__Go*)self)->cppgoBase_focusNextPrevChild(next);
	return __r;
}
int QWizardPage_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QWizardPage*)self)->QWizardPage::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QWizardPage_timerEvent_default(void* self, void* event) {
	((QWizardPage__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QWizardPage_childEvent_default(void* self, void* event) {
	((QWizardPage__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QWizardPage_customEvent_default(void* self, void* event) {
	((QWizardPage__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QWizardPage_connectNotify_default(void* self, void* signal) {
	((QWizardPage__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QWizardPage_disconnectNotify_default(void* self, void* signal) {
	((QWizardPage__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QWizardPage_connect_completeChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QWizardPage*)self, &QWizardPage::completeChanged, (QWizardPage*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QWizardPage);
	__conns_QWizardPage[__id] = __c;
}
void QWizardPage_connect_windowTitleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QWizardPage*)self, &QWizardPage::windowTitleChanged, (QWizardPage*)self, [__id](const QString& title){
		QByteArray title_ba = title.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)title_ba.constData(), title_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QWizardPage);
	__conns_QWizardPage[__id] = __c;
}
void QWizardPage_connect_windowIconChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QWizardPage*)self, &QWizardPage::windowIconChanged, (QWizardPage*)self, [__id](const QIcon& icon){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&icon);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QWizardPage);
	__conns_QWizardPage[__id] = __c;
}
void QWizardPage_connect_windowIconTextChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QWizardPage*)self, &QWizardPage::windowIconTextChanged, (QWizardPage*)self, [__id](const QString& iconText){
		QByteArray iconText_ba = iconText.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)iconText_ba.constData(), iconText_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QWizardPage);
	__conns_QWizardPage[__id] = __c;
}
void QWizardPage_connect_customContextMenuRequested(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QWizardPage*)self, &QWizardPage::customContextMenuRequested, (QWizardPage*)self, [__id](const QPoint& pos){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&pos);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QWizardPage);
	__conns_QWizardPage[__id] = __c;
}
void QWizardPage_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QWizardPage*)self, &QWizardPage::destroyed, (QWizardPage*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QWizardPage);
	__conns_QWizardPage[__id] = __c;
}
void QWizardPage_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QWizardPage*)self, &QWizardPage::destroyed, (QWizardPage*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QWizardPage);
	__conns_QWizardPage[__id] = __c;
}
void QWizardPage_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QWizardPage*)self, &QWizardPage::objectNameChanged, (QWizardPage*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QWizardPage);
	__conns_QWizardPage[__id] = __c;
}
void QWizardPage__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QWizardPage);
		auto __it = __conns_QWizardPage.find(id);
		if (__it == __conns_QWizardPage.end()) return;
		__c = __it->second;
		__conns_QWizardPage.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QWizardPage__Go_new(long long id) { return new QWizardPage__Go(id); }
void QWizardPage_delete(void* self) { delete (QWizardPage*)self; }

}

