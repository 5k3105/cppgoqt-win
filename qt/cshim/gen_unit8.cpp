#include <QtCore/qcompare.h>
#include "gen_qt__strong_ordering.h"

extern "C" {

void* Qt__strong_ordering_new(void* arg1) {
	return new Qt::strong_ordering(*(Qt::strong_ordering*)arg1);
}
void* Qt__strong_ordering_ToQtPartialOrdering(void* self) {
	auto&& __r = static_cast<Qt::partial_ordering>(*(Qt::strong_ordering*)self);
	return new Qt::partial_ordering(__r);
}
void* Qt__strong_ordering_ToQtWeakOrdering(void* self) {
	auto&& __r = static_cast<Qt::weak_ordering>(*(Qt::strong_ordering*)self);
	return new Qt::weak_ordering(__r);
}
int Qt__strong_ordering_OperatorEqual(void* self, void* rhs) {
	auto&& __r = (*(Qt::strong_ordering*)self) == (*(Qt::strong_ordering*)rhs);
	return __r;
}
int Qt__strong_ordering_OperatorNotEqual(void* self, void* rhs) {
	auto&& __r = (*(Qt::strong_ordering*)self) != (*(Qt::strong_ordering*)rhs);
	return __r;
}
int Qt__strong_ordering_OperatorEqual2(void* self, void* rhs) {
	auto&& __r = (*(Qt::strong_ordering*)self) == (*(Qt::partial_ordering*)rhs);
	return __r;
}
int Qt__strong_ordering_OperatorNotEqual2(void* self, void* rhs) {
	auto&& __r = (*(Qt::strong_ordering*)self) != (*(Qt::partial_ordering*)rhs);
	return __r;
}
int Qt__strong_ordering_OperatorEqual3(void* self, void* rhs) {
	auto&& __r = (*(Qt::strong_ordering*)self) == (*(Qt::weak_ordering*)rhs);
	return __r;
}
int Qt__strong_ordering_OperatorNotEqual3(void* self, void* rhs) {
	auto&& __r = (*(Qt::strong_ordering*)self) != (*(Qt::weak_ordering*)rhs);
	return __r;
}
void Qt__strong_ordering_delete(void* self) { delete (Qt::strong_ordering*)self; }

}

#include <QtCore/qstringconverter_base.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qbytearrayview.h>
#include <QtCore/qstring.h>
#include "gen_qstringconverter.h"

extern "C" {

void* QStringConverter_base(void* self) { return (void*)static_cast<QStringConverterBase*>((QStringConverter*)self); }
int QStringConverter_isValid(void* self) {
	auto&& __r = ((QStringConverter*)self)->isValid();
	return __r;
}
void QStringConverter_resetState(void* self) {
	((QStringConverter*)self)->resetState();
}
int QStringConverter_hasError(void* self) {
	auto&& __r = ((QStringConverter*)self)->hasError();
	return __r;
}
const char* QStringConverter_name(void* self) {
	auto&& __r = ((QStringConverter*)self)->name();
	return __r;
}
int QStringConverter_encodingForName(const char* name, long long name_len, int* __ok) {
	QAnyStringView name_qstr = QAnyStringView(name, name_len);
	auto&& __r = QStringConverter::encodingForName(name_qstr);
	*__ok = __r.has_value() ? 1 : 0;
	return __r.has_value() ? (int)(*__r) : (int)0;
}
const char* QStringConverter_nameForEncoding(int e) {
	auto&& __r = QStringConverter::nameForEncoding((QStringConverter::Encoding)(e));
	return __r;
}
int QStringConverter_encodingForData(void* data, int* __ok) {
	auto&& __r = QStringConverter::encodingForData(*(QByteArrayView*)data);
	*__ok = __r.has_value() ? 1 : 0;
	return __r.has_value() ? (int)(*__r) : (int)0;
}
int QStringConverter_encodingForHtml(void* data, int* __ok) {
	auto&& __r = QStringConverter::encodingForHtml(*(QByteArrayView*)data);
	*__ok = __r.has_value() ? 1 : 0;
	return __r.has_value() ? (int)(*__r) : (int)0;
}
char** QStringConverter_availableCodecs(long long* __out_len) {
	auto&& __r = QStringConverter::availableCodecs();
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
int QStringConverter_encodingForData2(void* data, unsigned short expectedFirstCharacter, int* __ok) {
	auto&& __r = QStringConverter::encodingForData(*(QByteArrayView*)data, expectedFirstCharacter);
	*__ok = __r.has_value() ? 1 : 0;
	return __r.has_value() ? (int)(*__r) : (int)0;
}

}

#include <QtCore/qabstracteventdispatcher.h>
#include <QtCore/qnamespace.h>
#include "gen_qabstracteventdispatcher__timerinfov2.h"

extern "C" {

void* QAbstractEventDispatcher__TimerInfoV2_new(void* arg1) {
	return new QAbstractEventDispatcher::TimerInfoV2(*(QAbstractEventDispatcher::TimerInfoV2*)arg1);
}
void* QAbstractEventDispatcher__TimerInfoV2_new2() {
	return new QAbstractEventDispatcher::TimerInfoV2();
}
void* QAbstractEventDispatcher__TimerInfoV2_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QAbstractEventDispatcher::TimerInfoV2*)self)->operator=(*(QAbstractEventDispatcher::TimerInfoV2*)arg1);
	return (void*)&(__r);
}
int QAbstractEventDispatcher__TimerInfoV2_timerId_get(void* self) {
	auto __r = ((QAbstractEventDispatcher::TimerInfoV2*)self)->timerId;
	return (int)(__r);
}
void QAbstractEventDispatcher__TimerInfoV2_timerId_set(void* self, int timerId) {
	((QAbstractEventDispatcher::TimerInfoV2*)self)->timerId = (Qt::TimerId)(timerId);
}
int QAbstractEventDispatcher__TimerInfoV2_timerType_get(void* self) {
	auto __r = ((QAbstractEventDispatcher::TimerInfoV2*)self)->timerType;
	return (int)(__r);
}
void QAbstractEventDispatcher__TimerInfoV2_timerType_set(void* self, int timerType) {
	((QAbstractEventDispatcher::TimerInfoV2*)self)->timerType = (Qt::TimerType)(timerType);
}
void QAbstractEventDispatcher__TimerInfoV2_delete(void* self) { delete (QAbstractEventDispatcher::TimerInfoV2*)self; }

}

#include <QtCore/qmutex.h>
#include "gen_qbasicmutex.h"

extern "C" {

void* QBasicMutex_new() {
	return new QBasicMutex();
}
void QBasicMutex_lock(void* self) {
	((QBasicMutex*)self)->lock();
}
void QBasicMutex_unlock(void* self) {
	((QBasicMutex*)self)->unlock();
}
int QBasicMutex_tryLock(void* self) {
	auto&& __r = ((QBasicMutex*)self)->tryLock();
	return __r;
}
int QBasicMutex_try_lock(void* self) {
	auto&& __r = ((QBasicMutex*)self)->try_lock();
	return __r;
}
void QBasicMutex_delete(void* self) { delete (QBasicMutex*)self; }

}

#include <QtCore/qdatetime.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qtimezone.h>
#include <QtCore/qstring.h>
#include <QtCore/qcalendar.h>
#include <QtCore/qstringview.h>
#include "gen_qdatetime.h"

extern "C" {

void* QDateTime_new() {
	return new QDateTime();
}
void* QDateTime_new2(void* date, void* time, int spec) {
	return new QDateTime(*(QDate*)date, *(QTime*)time, (Qt::TimeSpec)(spec));
}
void* QDateTime_new3(void* date, void* time, void* timeZone) {
	return new QDateTime(*(QDate*)date, *(QTime*)time, *(QTimeZone*)timeZone);
}
void* QDateTime_new4(void* date, void* time) {
	return new QDateTime(*(QDate*)date, *(QTime*)time);
}
void* QDateTime_new5(void* other) {
	return new QDateTime(*(QDateTime*)other);
}
void* QDateTime_new6(void* date, void* time, int spec, int offsetSeconds) {
	return new QDateTime(*(QDate*)date, *(QTime*)time, (Qt::TimeSpec)(spec), offsetSeconds);
}
void* QDateTime_new7(void* date, void* time, void* timeZone, int resolve) {
	return new QDateTime(*(QDate*)date, *(QTime*)time, *(QTimeZone*)timeZone, (QDateTime::TransitionResolution)(resolve));
}
void* QDateTime_new8(void* date, void* time, int resolve) {
	return new QDateTime(*(QDate*)date, *(QTime*)time, (QDateTime::TransitionResolution)(resolve));
}
void* QDateTime_OperatorAssign(void* self, void* other) {
	auto&& __r = ((QDateTime*)self)->operator=(*(QDateTime*)other);
	return (void*)&(__r);
}
void QDateTime_swap(void* self, void* other) {
	((QDateTime*)self)->swap(*(QDateTime*)other);
}
int QDateTime_isNull(void* self) {
	auto&& __r = ((QDateTime*)self)->isNull();
	return __r;
}
int QDateTime_isValid(void* self) {
	auto&& __r = ((QDateTime*)self)->isValid();
	return __r;
}
void* QDateTime_date(void* self) {
	auto&& __r = ((QDateTime*)self)->date();
	return new QDate(__r);
}
void* QDateTime_time(void* self) {
	auto&& __r = ((QDateTime*)self)->time();
	return new QTime(__r);
}
int QDateTime_timeSpec(void* self) {
	auto&& __r = ((QDateTime*)self)->timeSpec();
	return (int)(__r);
}
int QDateTime_offsetFromUtc(void* self) {
	auto&& __r = ((QDateTime*)self)->offsetFromUtc();
	return __r;
}
void* QDateTime_timeRepresentation(void* self) {
	auto&& __r = ((QDateTime*)self)->timeRepresentation();
	return new QTimeZone(__r);
}
void* QDateTime_timeZone(void* self) {
	auto&& __r = ((QDateTime*)self)->timeZone();
	return new QTimeZone(__r);
}
char* QDateTime_timeZoneAbbreviation(void* self, long long* __out_len) {
	auto&& __r = ((QDateTime*)self)->timeZoneAbbreviation();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QDateTime_isDaylightTime(void* self) {
	auto&& __r = ((QDateTime*)self)->isDaylightTime();
	return __r;
}
long long QDateTime_toMSecsSinceEpoch(void* self) {
	auto&& __r = ((QDateTime*)self)->toMSecsSinceEpoch();
	return __r;
}
long long QDateTime_toSecsSinceEpoch(void* self) {
	auto&& __r = ((QDateTime*)self)->toSecsSinceEpoch();
	return __r;
}
void QDateTime_setDate(void* self, void* date) {
	((QDateTime*)self)->setDate(*(QDate*)date);
}
void QDateTime_setTime(void* self, void* time) {
	((QDateTime*)self)->setTime(*(QTime*)time);
}
void QDateTime_setTimeSpec(void* self, int spec) {
	((QDateTime*)self)->setTimeSpec((Qt::TimeSpec)(spec));
}
void QDateTime_setOffsetFromUtc(void* self, int offsetSeconds) {
	((QDateTime*)self)->setOffsetFromUtc(offsetSeconds);
}
void QDateTime_setTimeZone(void* self, void* toZone) {
	((QDateTime*)self)->setTimeZone(*(QTimeZone*)toZone);
}
void QDateTime_setMSecsSinceEpoch(void* self, long long msecs) {
	((QDateTime*)self)->setMSecsSinceEpoch(msecs);
}
void QDateTime_setSecsSinceEpoch(void* self, long long secs) {
	((QDateTime*)self)->setSecsSinceEpoch(secs);
}
char* QDateTime_toString(void* self, long long* __out_len) {
	auto&& __r = ((QDateTime*)self)->toString();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QDateTime_toString2(void* self, const char* format, long long format_len, long long* __out_len) {
	QString format_qstr = QString::fromUtf8(format, format_len);
	auto&& __r = ((QDateTime*)self)->toString(format_qstr);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QDateTime_toString3(void* self, const char* format, long long format_len, void* cal, long long* __out_len) {
	QString format_qstr = QString::fromUtf8(format, format_len);
	auto&& __r = ((QDateTime*)self)->toString(format_qstr, *(QCalendar*)cal);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QDateTime_toString4(void* self, void* format, long long* __out_len) {
	auto&& __r = ((QDateTime*)self)->toString(*(QStringView*)format);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QDateTime_toString5(void* self, void* format, void* cal, long long* __out_len) {
	auto&& __r = ((QDateTime*)self)->toString(*(QStringView*)format, *(QCalendar*)cal);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QDateTime_addDays(void* self, long long days) {
	auto&& __r = ((QDateTime*)self)->addDays(days);
	return new QDateTime(__r);
}
void* QDateTime_addMonths(void* self, int months) {
	auto&& __r = ((QDateTime*)self)->addMonths(months);
	return new QDateTime(__r);
}
void* QDateTime_addYears(void* self, int years) {
	auto&& __r = ((QDateTime*)self)->addYears(years);
	return new QDateTime(__r);
}
void* QDateTime_addSecs(void* self, long long secs) {
	auto&& __r = ((QDateTime*)self)->addSecs(secs);
	return new QDateTime(__r);
}
void* QDateTime_addMSecs(void* self, long long msecs) {
	auto&& __r = ((QDateTime*)self)->addMSecs(msecs);
	return new QDateTime(__r);
}
void* QDateTime_addDuration(void* self, long long msecs) {
	auto&& __r = ((QDateTime*)self)->addDuration(std::chrono::milliseconds((long long)msecs));
	return new QDateTime(__r);
}
void* QDateTime_toTimeSpec(void* self, int spec) {
	auto&& __r = ((QDateTime*)self)->toTimeSpec((Qt::TimeSpec)(spec));
	return new QDateTime(__r);
}
void* QDateTime_toLocalTime(void* self) {
	auto&& __r = ((QDateTime*)self)->toLocalTime();
	return new QDateTime(__r);
}
void* QDateTime_toUTC(void* self) {
	auto&& __r = ((QDateTime*)self)->toUTC();
	return new QDateTime(__r);
}
void* QDateTime_toOffsetFromUtc(void* self, int offsetSeconds) {
	auto&& __r = ((QDateTime*)self)->toOffsetFromUtc(offsetSeconds);
	return new QDateTime(__r);
}
void* QDateTime_toTimeZone(void* self, void* toZone) {
	auto&& __r = ((QDateTime*)self)->toTimeZone(*(QTimeZone*)toZone);
	return new QDateTime(__r);
}
long long QDateTime_daysTo(void* self, void* arg1) {
	auto&& __r = ((QDateTime*)self)->daysTo(*(QDateTime*)arg1);
	return __r;
}
long long QDateTime_secsTo(void* self, void* arg1) {
	auto&& __r = ((QDateTime*)self)->secsTo(*(QDateTime*)arg1);
	return __r;
}
long long QDateTime_msecsTo(void* self, void* arg1) {
	auto&& __r = ((QDateTime*)self)->msecsTo(*(QDateTime*)arg1);
	return __r;
}
void* QDateTime_currentDateTime(void* zone) {
	auto&& __r = QDateTime::currentDateTime(*(QTimeZone*)zone);
	return new QDateTime(__r);
}
void* QDateTime_currentDateTime2() {
	auto&& __r = QDateTime::currentDateTime();
	return new QDateTime(__r);
}
void* QDateTime_currentDateTimeUtc() {
	auto&& __r = QDateTime::currentDateTimeUtc();
	return new QDateTime(__r);
}
void* QDateTime_fromString(void* string_) {
	auto&& __r = QDateTime::fromString(*(QStringView*)string_);
	return new QDateTime(__r);
}
void* QDateTime_fromString2(const char* string_, long long string__len) {
	QString string__qstr = QString::fromUtf8(string_, string__len);
	auto&& __r = QDateTime::fromString(string__qstr);
	return new QDateTime(__r);
}
void* QDateTime_fromString3(void* string_, void* format, void* cal) {
	auto&& __r = QDateTime::fromString(*(QStringView*)string_, *(QStringView*)format, *(QCalendar*)cal);
	return new QDateTime(__r);
}
void* QDateTime_fromString4(const char* string_, long long string__len, void* format, void* cal) {
	QString string__qstr = QString::fromUtf8(string_, string__len);
	auto&& __r = QDateTime::fromString(string__qstr, *(QStringView*)format, *(QCalendar*)cal);
	return new QDateTime(__r);
}
void* QDateTime_fromString5(const char* string_, long long string__len, const char* format, long long format_len, void* cal) {
	QString string__qstr = QString::fromUtf8(string_, string__len);
	QString format_qstr = QString::fromUtf8(format, format_len);
	auto&& __r = QDateTime::fromString(string__qstr, format_qstr, *(QCalendar*)cal);
	return new QDateTime(__r);
}
void* QDateTime_fromString6(void* string_, void* format) {
	auto&& __r = QDateTime::fromString(*(QStringView*)string_, *(QStringView*)format);
	return new QDateTime(__r);
}
void* QDateTime_fromString7(void* string_, void* format, int baseYear, void* cal) {
	auto&& __r = QDateTime::fromString(*(QStringView*)string_, *(QStringView*)format, baseYear, *(QCalendar*)cal);
	return new QDateTime(__r);
}
void* QDateTime_fromString8(const char* string_, long long string__len, void* format) {
	QString string__qstr = QString::fromUtf8(string_, string__len);
	auto&& __r = QDateTime::fromString(string__qstr, *(QStringView*)format);
	return new QDateTime(__r);
}
void* QDateTime_fromString9(const char* string_, long long string__len, void* format, int baseYear, void* cal) {
	QString string__qstr = QString::fromUtf8(string_, string__len);
	auto&& __r = QDateTime::fromString(string__qstr, *(QStringView*)format, baseYear, *(QCalendar*)cal);
	return new QDateTime(__r);
}
void* QDateTime_fromString10(const char* string_, long long string__len, const char* format, long long format_len) {
	QString string__qstr = QString::fromUtf8(string_, string__len);
	QString format_qstr = QString::fromUtf8(format, format_len);
	auto&& __r = QDateTime::fromString(string__qstr, format_qstr);
	return new QDateTime(__r);
}
void* QDateTime_fromString11(const char* string_, long long string__len, const char* format, long long format_len, int baseYear, void* cal) {
	QString string__qstr = QString::fromUtf8(string_, string__len);
	QString format_qstr = QString::fromUtf8(format, format_len);
	auto&& __r = QDateTime::fromString(string__qstr, format_qstr, baseYear, *(QCalendar*)cal);
	return new QDateTime(__r);
}
void* QDateTime_fromMSecsSinceEpoch(long long msecs, int spec) {
	auto&& __r = QDateTime::fromMSecsSinceEpoch(msecs, (Qt::TimeSpec)(spec));
	return new QDateTime(__r);
}
void* QDateTime_fromSecsSinceEpoch(long long secs, int spec) {
	auto&& __r = QDateTime::fromSecsSinceEpoch(secs, (Qt::TimeSpec)(spec));
	return new QDateTime(__r);
}
void* QDateTime_fromMSecsSinceEpoch2(long long msecs, void* timeZone) {
	auto&& __r = QDateTime::fromMSecsSinceEpoch(msecs, *(QTimeZone*)timeZone);
	return new QDateTime(__r);
}
void* QDateTime_fromSecsSinceEpoch2(long long secs, void* timeZone) {
	auto&& __r = QDateTime::fromSecsSinceEpoch(secs, *(QTimeZone*)timeZone);
	return new QDateTime(__r);
}
void* QDateTime_fromMSecsSinceEpoch3(long long msecs) {
	auto&& __r = QDateTime::fromMSecsSinceEpoch(msecs);
	return new QDateTime(__r);
}
void* QDateTime_fromSecsSinceEpoch3(long long secs) {
	auto&& __r = QDateTime::fromSecsSinceEpoch(secs);
	return new QDateTime(__r);
}
long long QDateTime_currentMSecsSinceEpoch() {
	auto&& __r = QDateTime::currentMSecsSinceEpoch();
	return __r;
}
long long QDateTime_currentSecsSinceEpoch() {
	auto&& __r = QDateTime::currentSecsSinceEpoch();
	return __r;
}
void* QDateTime_OperatorPlusAssign(void* self, long long duration) {
	auto&& __r = ((QDateTime*)self)->operator+=(std::chrono::milliseconds((long long)duration));
	return (void*)&(__r);
}
void* QDateTime_OperatorMinusAssign(void* self, long long duration) {
	auto&& __r = ((QDateTime*)self)->operator-=(std::chrono::milliseconds((long long)duration));
	return (void*)&(__r);
}
void QDateTime_setDate2(void* self, void* date, int resolve) {
	((QDateTime*)self)->setDate(*(QDate*)date, (QDateTime::TransitionResolution)(resolve));
}
void QDateTime_setTime2(void* self, void* time, int resolve) {
	((QDateTime*)self)->setTime(*(QTime*)time, (QDateTime::TransitionResolution)(resolve));
}
void QDateTime_setTimeZone2(void* self, void* toZone, int resolve) {
	((QDateTime*)self)->setTimeZone(*(QTimeZone*)toZone, (QDateTime::TransitionResolution)(resolve));
}
char* QDateTime_toString6(void* self, int format, long long* __out_len) {
	auto&& __r = ((QDateTime*)self)->toString((Qt::DateFormat)(format));
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QDateTime_fromString12(void* string_, int format) {
	auto&& __r = QDateTime::fromString(*(QStringView*)string_, (Qt::DateFormat)(format));
	return new QDateTime(__r);
}
void* QDateTime_fromString13(const char* string_, long long string__len, int format) {
	QString string__qstr = QString::fromUtf8(string_, string__len);
	auto&& __r = QDateTime::fromString(string__qstr, (Qt::DateFormat)(format));
	return new QDateTime(__r);
}
void* QDateTime_fromString14(void* string_, void* format, int baseYear) {
	auto&& __r = QDateTime::fromString(*(QStringView*)string_, *(QStringView*)format, baseYear);
	return new QDateTime(__r);
}
void* QDateTime_fromString15(const char* string_, long long string__len, void* format, int baseYear) {
	QString string__qstr = QString::fromUtf8(string_, string__len);
	auto&& __r = QDateTime::fromString(string__qstr, *(QStringView*)format, baseYear);
	return new QDateTime(__r);
}
void* QDateTime_fromString16(const char* string_, long long string__len, const char* format, long long format_len, int baseYear) {
	QString string__qstr = QString::fromUtf8(string_, string__len);
	QString format_qstr = QString::fromUtf8(format, format_len);
	auto&& __r = QDateTime::fromString(string__qstr, format_qstr, baseYear);
	return new QDateTime(__r);
}
void* QDateTime_fromMSecsSinceEpoch4(long long msecs, int spec, int offsetFromUtc) {
	auto&& __r = QDateTime::fromMSecsSinceEpoch(msecs, (Qt::TimeSpec)(spec), offsetFromUtc);
	return new QDateTime(__r);
}
void* QDateTime_fromSecsSinceEpoch4(long long secs, int spec, int offsetFromUtc) {
	auto&& __r = QDateTime::fromSecsSinceEpoch(secs, (Qt::TimeSpec)(spec), offsetFromUtc);
	return new QDateTime(__r);
}
long long QDateTime_OperatorMinus(void* self, void* rhs) {
	auto&& __r = (*(QDateTime*)self) - (*(QDateTime*)rhs);
	return (long long)(__r).count();
}
void* QDateTime_OperatorPlus(void* self, long long duration) {
	auto&& __r = (*(QDateTime*)self) + (std::chrono::milliseconds((long long)duration));
	return new QDateTime(__r);
}
void* QDateTime_OperatorMinus2(void* self, long long duration) {
	auto&& __r = (*(QDateTime*)self) - (std::chrono::milliseconds((long long)duration));
	return new QDateTime(__r);
}
int QDateTime_OperatorEqual(void* self, void* rhs) {
	auto&& __r = (*(QDateTime*)self) == (*(QDateTime*)rhs);
	return __r;
}
int QDateTime_OperatorNotEqual(void* self, void* rhs) {
	auto&& __r = (*(QDateTime*)self) != (*(QDateTime*)rhs);
	return __r;
}
int QDateTime_OperatorLesser(void* self, void* rhs) {
	auto&& __r = (*(QDateTime*)self) < (*(QDateTime*)rhs);
	return __r;
}
int QDateTime_OperatorGreater(void* self, void* rhs) {
	auto&& __r = (*(QDateTime*)self) > (*(QDateTime*)rhs);
	return __r;
}
int QDateTime_OperatorLesserOrEqual(void* self, void* rhs) {
	auto&& __r = (*(QDateTime*)self) <= (*(QDateTime*)rhs);
	return __r;
}
int QDateTime_OperatorGreaterOrEqual(void* self, void* rhs) {
	auto&& __r = (*(QDateTime*)self) >= (*(QDateTime*)rhs);
	return __r;
}
void QDateTime_delete(void* self) { delete (QDateTime*)self; }

}

#include <QtCore/qproperty.h>
#include <QtCore/qpropertyprivate.h>
#include "gen_qpropertynotifier.h"

extern "C" {

void* QPropertyNotifier_base(void* self) { return (void*)static_cast<QPropertyObserver*>((QPropertyNotifier*)self); }
void QPropertyNotifier_chain(void* self, void** out) {
	QPropertyNotifier* __p = (QPropertyNotifier*)self;
	QPropertyObserver* __b0 = static_cast<QPropertyObserver*>(__p); out[0] = (void*)__b0;
	QPropertyObserverBase* __b1 = static_cast<QPropertyObserverBase*>(__b0); out[1] = (void*)__b1;
}
void* QPropertyNotifier_new() {
	return new QPropertyNotifier();
}
void QPropertyNotifier_delete(void* self) { delete (QPropertyNotifier*)self; }

}

#include <QtCore/qexception.h>
#include "gen_qunhandledexception.h"

extern "C" {

void* QUnhandledException_base(void* self) { return (void*)static_cast<QException*>((QUnhandledException*)self); }
void* QUnhandledException_new() {
	return new QUnhandledException();
}
void* QUnhandledException_new2(void* other) {
	return new QUnhandledException(*(QUnhandledException*)other);
}
void QUnhandledException_swap(void* self, void* other) {
	((QUnhandledException*)self)->swap(*(QUnhandledException*)other);
}
void* QUnhandledException_OperatorAssign(void* self, void* other) {
	auto&& __r = ((QUnhandledException*)self)->operator=(*(QUnhandledException*)other);
	return (void*)&(__r);
}
void QUnhandledException_delete(void* self) { delete (QUnhandledException*)self; }

}

#include <QtCore/qlibraryinfo.h>
#include <QtCore/qversionnumber.h>
#include <QtCore/qstring.h>
#include "gen_qlibraryinfo.h"

extern "C" {

const char* QLibraryInfo_build() {
	auto&& __r = QLibraryInfo::build();
	return __r;
}
int QLibraryInfo_isDebugBuild() {
	auto&& __r = QLibraryInfo::isDebugBuild();
	return __r;
}
int QLibraryInfo_isSharedBuild() {
	auto&& __r = QLibraryInfo::isSharedBuild();
	return __r;
}
void* QLibraryInfo_version() {
	auto&& __r = QLibraryInfo::version();
	return new QVersionNumber(__r);
}
char* QLibraryInfo_path(int p, long long* __out_len) {
	auto&& __r = QLibraryInfo::path((QLibraryInfo::LibraryPath)(p));
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char** QLibraryInfo_paths(int p, long long* __out_len) {
	auto&& __r = QLibraryInfo::paths((QLibraryInfo::LibraryPath)(p));
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
char* QLibraryInfo_location(int location, long long* __out_len) {
	auto&& __r = QLibraryInfo::location((QLibraryInfo::LibraryPath)(location));
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char** QLibraryInfo_platformPluginArguments(const char* platformName, long long platformName_len, long long* __out_len) {
	QString platformName_qstr = QString::fromUtf8(platformName, platformName_len);
	auto&& __r = QLibraryInfo::platformPluginArguments(platformName_qstr);
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
void QLibraryInfo_delete(void* self) { delete (QLibraryInfo*)self; }

}

#include <QtCore/qpermissions.h>
#include "gen_qlocationpermission.h"

extern "C" {

void* QLocationPermission_new() {
	return new QLocationPermission();
}
void* QLocationPermission_new2(void* other) {
	return new QLocationPermission(*(QLocationPermission*)other);
}
void QLocationPermission_setAccuracy(void* self, int accuracy) {
	((QLocationPermission*)self)->setAccuracy((QLocationPermission::Accuracy)(accuracy));
}
int QLocationPermission_accuracy(void* self) {
	auto&& __r = ((QLocationPermission*)self)->accuracy();
	return (int)(__r);
}
void QLocationPermission_setAvailability(void* self, int availability) {
	((QLocationPermission*)self)->setAvailability((QLocationPermission::Availability)(availability));
}
int QLocationPermission_availability(void* self) {
	auto&& __r = ((QLocationPermission*)self)->availability();
	return (int)(__r);
}
void* QLocationPermission_OperatorAssign(void* self, void* other) {
	auto&& __r = ((QLocationPermission*)self)->operator=(*(QLocationPermission*)other);
	return (void*)&(__r);
}
void QLocationPermission_swap(void* self, void* other) {
	((QLocationPermission*)self)->swap(*(QLocationPermission*)other);
}
void QLocationPermission_delete(void* self) { delete (QLocationPermission*)self; }

}

#include <QtCore/qsize.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qmargins.h>
#include "gen_qsize.h"

extern "C" {

void* QSize_new() {
	return new QSize();
}
void* QSize_new2(int w, int h) {
	return new QSize(w, h);
}
void* QSize_new3(void* arg1) {
	return new QSize(*(QSize*)arg1);
}
int QSize_isNull(void* self) {
	auto&& __r = ((QSize*)self)->isNull();
	return __r;
}
int QSize_isEmpty(void* self) {
	auto&& __r = ((QSize*)self)->isEmpty();
	return __r;
}
int QSize_isValid(void* self) {
	auto&& __r = ((QSize*)self)->isValid();
	return __r;
}
int QSize_width(void* self) {
	auto&& __r = ((QSize*)self)->width();
	return __r;
}
int QSize_height(void* self) {
	auto&& __r = ((QSize*)self)->height();
	return __r;
}
void QSize_setWidth(void* self, int w) {
	((QSize*)self)->setWidth(w);
}
void QSize_setHeight(void* self, int h) {
	((QSize*)self)->setHeight(h);
}
void QSize_transpose(void* self) {
	((QSize*)self)->transpose();
}
void* QSize_transposed(void* self) {
	auto&& __r = ((QSize*)self)->transposed();
	return new QSize(__r);
}
void QSize_scale(void* self, int w, int h, int mode) {
	((QSize*)self)->scale(w, h, (Qt::AspectRatioMode)(mode));
}
void QSize_scale2(void* self, void* s, int mode) {
	((QSize*)self)->scale(*(QSize*)s, (Qt::AspectRatioMode)(mode));
}
void* QSize_scaled(void* self, int w, int h, int mode) {
	auto&& __r = ((QSize*)self)->scaled(w, h, (Qt::AspectRatioMode)(mode));
	return new QSize(__r);
}
void* QSize_scaled2(void* self, void* s, int mode) {
	auto&& __r = ((QSize*)self)->scaled(*(QSize*)s, (Qt::AspectRatioMode)(mode));
	return new QSize(__r);
}
void* QSize_expandedTo(void* self, void* arg1) {
	auto&& __r = ((QSize*)self)->expandedTo(*(QSize*)arg1);
	return new QSize(__r);
}
void* QSize_boundedTo(void* self, void* arg1) {
	auto&& __r = ((QSize*)self)->boundedTo(*(QSize*)arg1);
	return new QSize(__r);
}
void* QSize_grownBy(void* self, void* m) {
	auto&& __r = ((QSize*)self)->grownBy(*(QMargins*)m);
	return new QSize(__r);
}
void* QSize_shrunkBy(void* self, void* m) {
	auto&& __r = ((QSize*)self)->shrunkBy(*(QMargins*)m);
	return new QSize(__r);
}
void* QSize_OperatorPlusAssign(void* self, void* arg1) {
	auto&& __r = ((QSize*)self)->operator+=(*(QSize*)arg1);
	return (void*)&(__r);
}
void* QSize_OperatorMinusAssign(void* self, void* arg1) {
	auto&& __r = ((QSize*)self)->operator-=(*(QSize*)arg1);
	return (void*)&(__r);
}
void* QSize_OperatorMultiplyAssign(void* self, double c) {
	auto&& __r = ((QSize*)self)->operator*=(c);
	return (void*)&(__r);
}
void* QSize_OperatorDivideAssign(void* self, double c) {
	auto&& __r = ((QSize*)self)->operator/=(c);
	return (void*)&(__r);
}
void* QSize_toSizeF(void* self) {
	auto&& __r = ((QSize*)self)->toSizeF();
	return new QSizeF(__r);
}
void* QSize_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QSize*)self)->operator=(*(QSize*)arg1);
	return (void*)&(__r);
}
int QSize_OperatorEqual(void* self, void* rhs) {
	auto&& __r = (*(QSize*)self) == (*(QSize*)rhs);
	return __r;
}
int QSize_OperatorNotEqual(void* self, void* rhs) {
	auto&& __r = (*(QSize*)self) != (*(QSize*)rhs);
	return __r;
}
void* QSize_OperatorPlus(void* self, void* s2) {
	auto&& __r = (*(QSize*)self) + (*(QSize*)s2);
	return new QSize(__r);
}
void* QSize_OperatorMinus(void* self, void* s2) {
	auto&& __r = (*(QSize*)self) - (*(QSize*)s2);
	return new QSize(__r);
}
void* QSize_OperatorMultiply(void* self, double c) {
	auto&& __r = (*(QSize*)self) * (c);
	return new QSize(__r);
}
void* QSize_OperatorDivide(void* self, double c) {
	auto&& __r = (*(QSize*)self) / (c);
	return new QSize(__r);
}
int QSize_OperatorEqual2(void* self, void* rhs) {
	auto&& __r = (*(QSize*)self) == (*(QSizeF*)rhs);
	return __r;
}
int QSize_OperatorNotEqual2(void* self, void* rhs) {
	auto&& __r = (*(QSize*)self) != (*(QSizeF*)rhs);
	return __r;
}
void QSize_delete(void* self) { delete (QSize*)self; }

}

#include <QtCore/qtemporaryfile.h>
#include <QtCore/qfile.h>
#include <QtCore/qstring.h>
#include <QtCore/qobject.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qiodevicebase.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qfiledevice.h>
#include <QtCore/qdatetime.h>
#include <QtCore/qbytearrayview.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qvariant.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qtemporaryfile.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QTemporaryFile__Go : public QTemporaryFile {
public:
	long long __goId;
	unsigned char __goOn[26] = {};
	template<typename... A> QTemporaryFile__Go(long long id, A&&... a) : QTemporaryFile(std::forward<A>(a)...), __goId(id) {}
	using QTemporaryFile::setOpenMode;
	using QTemporaryFile::setErrorString;
	using QTemporaryFile::sender;
	using QTemporaryFile::senderSignalIndex;
	using QTemporaryFile::receivers;
	using QTemporaryFile::isSignalConnected;
	qint64 cppgoBase_skipData(qint64 __a0) { return QTemporaryFile::skipData(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QTemporaryFile::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QTemporaryFile::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QTemporaryFile::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QTemporaryFile::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QTemporaryFile::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QTemporaryFile::metaObject();
	}
	QString fileName() const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virtp_charp) {
			long long __len = 0;
			char* __s = ((char*(*)(long long, long long, long long*, int*))cppgo_virtp_charp)(__goId, 8848495542063941256LL, &__len, &__h);
			if (__h) { QString __r = QString::fromUtf8(__s, __len); free(__s); return __r; }
		}
		return QTemporaryFile::fileName();
	}
	bool open(QFlags<QIODeviceBase::OpenModeFlag> flags) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, -554057428128861975LL, (int)(flags), &__h); if (__h) return __r; }
		return QTemporaryFile::open(flags);
	}
	qint64 size() const override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_longlong) { long long __r = ((long long(*)(long long, long long, int*))cppgo_virt_longlong)(__goId, 5614464919154503228LL, &__h); if (__h) return __r; }
		return QTemporaryFile::size();
	}
	bool resize(qint64 sz) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_int_longlong) { int __r = ((int(*)(long long, long long, long long, int*))cppgo_virt_int_longlong)(__goId, -4890198374319873567LL, sz, &__h); if (__h) return __r; }
		return QTemporaryFile::resize(sz);
	}
	QFlags<QFileDevice::Permission> permissions() const override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -7517927236947282013LL, &__h); if (__h) return QFlags<QFileDevice::Permission>(__r); }
		return QTemporaryFile::permissions();
	}
	bool setPermissions(QFlags<QFileDevice::Permission> permissionSpec) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 420407675554349589LL, (int)(permissionSpec), &__h); if (__h) return __r; }
		return QTemporaryFile::setPermissions(permissionSpec);
	}
	void close() override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, -7006921981232612925LL, &__h); if (__h) return; }
		QTemporaryFile::close();
	}
	bool isSequential() const override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, 8888098400087344862LL, &__h); if (__h) return __r; }
		return QTemporaryFile::isSequential();
	}
	qint64 pos() const override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_longlong) { long long __r = ((long long(*)(long long, long long, int*))cppgo_virt_longlong)(__goId, 8618582476283917417LL, &__h); if (__h) return __r; }
		return QTemporaryFile::pos();
	}
	bool seek(qint64 offset) override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_int_longlong) { int __r = ((int(*)(long long, long long, long long, int*))cppgo_virt_int_longlong)(__goId, 3241140482992819217LL, offset, &__h); if (__h) return __r; }
		return QTemporaryFile::seek(offset);
	}
	bool atEnd() const override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -1487528724786208339LL, &__h); if (__h) return __r; }
		return QTemporaryFile::atEnd();
	}
	bool reset() override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, 1086335023529244512LL, &__h); if (__h) return __r; }
		return QTemporaryFile::reset();
	}
	qint64 bytesAvailable() const override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_longlong) { long long __r = ((long long(*)(long long, long long, int*))cppgo_virt_longlong)(__goId, 4814449630054007943LL, &__h); if (__h) return __r; }
		return QTemporaryFile::bytesAvailable();
	}
	qint64 bytesToWrite() const override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_longlong) { long long __r = ((long long(*)(long long, long long, int*))cppgo_virt_longlong)(__goId, 1024747515116322948LL, &__h); if (__h) return __r; }
		return QTemporaryFile::bytesToWrite();
	}
	bool canReadLine() const override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -1250682508954504761LL, &__h); if (__h) return __r; }
		return QTemporaryFile::canReadLine();
	}
	bool waitForReadyRead(int msecs) override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 8356976193102722774LL, msecs, &__h); if (__h) return __r; }
		return QTemporaryFile::waitForReadyRead(msecs);
	}
	bool waitForBytesWritten(int msecs) override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 7726833620054767889LL, msecs, &__h); if (__h) return __r; }
		return QTemporaryFile::waitForBytesWritten(msecs);
	}
	qint64 skipData(qint64 maxSize) override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_longlong_longlong) { long long __r = ((long long(*)(long long, long long, long long, int*))cppgo_virt_longlong_longlong)(__goId, 8178738247667652082LL, maxSize, &__h); if (__h) return __r; }
		return QTemporaryFile::skipData(maxSize);
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QTemporaryFile::event(event);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QTemporaryFile::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QTemporaryFile::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QTemporaryFile::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QTemporaryFile::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QTemporaryFile::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QTemporaryFile::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QTemporaryFile;
static std::mutex __connsMu_QTemporaryFile;

extern "C" {

void* QTemporaryFile_base(void* self) { return (void*)static_cast<QFile*>((QTemporaryFile*)self); }
void QTemporaryFile_chain(void* self, void** out) {
	QTemporaryFile* __p = (QTemporaryFile*)self;
	QFile* __b0 = static_cast<QFile*>(__p); out[0] = (void*)__b0;
	QFileDevice* __b1 = static_cast<QFileDevice*>(__b0); out[1] = (void*)__b1;
	QIODevice* __b2 = static_cast<QIODevice*>(__b1); out[2] = (void*)__b2;
	QObject* __b3 = static_cast<QObject*>(__b2); out[3] = (void*)__b3;
}
void QTemporaryFile__set_override(void* self, int idx) { ((QTemporaryFile__Go*)self)->__goOn[idx] = 1; }
void* QTemporaryFile_fromQObject(void* o) { return (void*)qobject_cast<QTemporaryFile*>((QObject*)o); }
void* QTemporaryFile_new(long long __goid) {
	return new QTemporaryFile__Go(__goid);
}
void* QTemporaryFile_new2(long long __goid, const char* templateName, long long templateName_len) {
	QString templateName_qstr = QString::fromUtf8(templateName, templateName_len);
	return new QTemporaryFile__Go(__goid, templateName_qstr);
}
void* QTemporaryFile_new3(long long __goid, void* parent) {
	return new QTemporaryFile__Go(__goid, (QObject*)parent);
}
void* QTemporaryFile_new4(long long __goid, const char* templateName, long long templateName_len, void* parent) {
	QString templateName_qstr = QString::fromUtf8(templateName, templateName_len);
	return new QTemporaryFile__Go(__goid, templateName_qstr, (QObject*)parent);
}
char* QTemporaryFile_tr(const char* s, long long* __out_len) {
	auto&& __r = QTemporaryFile::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QTemporaryFile_autoRemove(void* self) {
	auto&& __r = ((QTemporaryFile*)self)->autoRemove();
	return __r;
}
void QTemporaryFile_setAutoRemove(void* self, int b) {
	((QTemporaryFile*)self)->setAutoRemove(b);
}
int QTemporaryFile_open(void* self) {
	auto&& __r = ((QTemporaryFile*)self)->open();
	return __r;
}
char* QTemporaryFile_fileTemplate(void* self, long long* __out_len) {
	auto&& __r = ((QTemporaryFile*)self)->fileTemplate();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QTemporaryFile_setFileTemplate(void* self, const char* name, long long name_len) {
	QString name_qstr = QString::fromUtf8(name, name_len);
	((QTemporaryFile*)self)->setFileTemplate(name_qstr);
}
int QTemporaryFile_rename(void* self, const char* newName, long long newName_len) {
	QString newName_qstr = QString::fromUtf8(newName, newName_len);
	auto&& __r = ((QTemporaryFile*)self)->rename(newName_qstr);
	return __r;
}
int QTemporaryFile_renameOverwrite(void* self, const char* newName, long long newName_len) {
	QString newName_qstr = QString::fromUtf8(newName, newName_len);
	auto&& __r = ((QTemporaryFile*)self)->renameOverwrite(newName_qstr);
	return __r;
}
void* QTemporaryFile_createNativeFile(const char* fileName, long long fileName_len) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	auto&& __r = QTemporaryFile::createNativeFile(fileName_qstr);
	return (void*)(__r);
}
void* QTemporaryFile_createNativeFile2(void* file) {
	auto&& __r = QTemporaryFile::createNativeFile(*(QFile*)file);
	return (void*)(__r);
}
void* QTemporaryFile_encodeName(const char* fileName, long long fileName_len) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	auto&& __r = QTemporaryFile::encodeName(fileName_qstr);
	return new QByteArray(__r);
}
char* QTemporaryFile_decodeName(void* localFileName, long long* __out_len) {
	auto&& __r = QTemporaryFile::decodeName(*(QByteArray*)localFileName);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QTemporaryFile_decodeName2(const char* localFileName, long long* __out_len) {
	auto&& __r = QTemporaryFile::decodeName(localFileName);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QTemporaryFile_exists(const char* fileName, long long fileName_len) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	auto&& __r = QTemporaryFile::exists(fileName_qstr);
	return __r;
}
char* QTemporaryFile_symLinkTarget(const char* fileName, long long fileName_len, long long* __out_len) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	auto&& __r = QTemporaryFile::symLinkTarget(fileName_qstr);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QTemporaryFile_remove(const char* fileName, long long fileName_len) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	auto&& __r = QTemporaryFile::remove(fileName_qstr);
	return __r;
}
int QTemporaryFile_supportsMoveToTrash() {
	auto&& __r = QTemporaryFile::supportsMoveToTrash();
	return __r;
}
int QTemporaryFile_moveToTrash(const char* fileName, long long fileName_len) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	auto&& __r = QTemporaryFile::moveToTrash(fileName_qstr);
	return __r;
}
int QTemporaryFile_link(const char* fileName, long long fileName_len, const char* newName, long long newName_len) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	QString newName_qstr = QString::fromUtf8(newName, newName_len);
	auto&& __r = QTemporaryFile::link(fileName_qstr, newName_qstr);
	return __r;
}
int QTemporaryFile_copy(const char* fileName, long long fileName_len, const char* newName, long long newName_len) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	QString newName_qstr = QString::fromUtf8(newName, newName_len);
	auto&& __r = QTemporaryFile::copy(fileName_qstr, newName_qstr);
	return __r;
}
int QTemporaryFile_resize(const char* filename, long long filename_len, long long sz) {
	QString filename_qstr = QString::fromUtf8(filename, filename_len);
	auto&& __r = QTemporaryFile::resize(filename_qstr, sz);
	return __r;
}
int QTemporaryFile_permissions(const char* filename, long long filename_len) {
	QString filename_qstr = QString::fromUtf8(filename, filename_len);
	auto&& __r = QTemporaryFile::permissions(filename_qstr);
	return (int)(__r);
}
int QTemporaryFile_setPermissions(const char* filename, long long filename_len, int permissionSpec) {
	QString filename_qstr = QString::fromUtf8(filename, filename_len);
	auto&& __r = QTemporaryFile::setPermissions(filename_qstr, QFlags<QFileDevice::Permission>(permissionSpec));
	return __r;
}
void QTemporaryFile_setOpenMode(void* self, int openMode) {
	((QTemporaryFile__Go*)self)->setOpenMode(QFlags<QIODeviceBase::OpenModeFlag>(openMode));
}
void QTemporaryFile_setErrorString(void* self, const char* errorString, long long errorString_len) {
	QString errorString_qstr = QString::fromUtf8(errorString, errorString_len);
	((QTemporaryFile__Go*)self)->setErrorString(errorString_qstr);
}
void* QTemporaryFile_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QTemporaryFile::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QTemporaryFile_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QTemporaryFile::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QTemporaryFile_sender(void* self) {
	auto&& __r = ((QTemporaryFile__Go*)self)->sender();
	return (void*)(__r);
}
int QTemporaryFile_senderSignalIndex(void* self) {
	auto&& __r = ((QTemporaryFile__Go*)self)->senderSignalIndex();
	return __r;
}
int QTemporaryFile_receivers(void* self, const char* signal) {
	auto&& __r = ((QTemporaryFile__Go*)self)->receivers(signal);
	return __r;
}
int QTemporaryFile_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QTemporaryFile__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QTemporaryFile_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QTemporaryFile::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QTemporaryFile_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QTemporaryFile::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QTemporaryFile_moveToTrash2(const char* fileName, long long fileName_len, void* pathInTrash) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	auto&& __r = QTemporaryFile::moveToTrash(fileName_qstr, (QString*)pathInTrash);
	return __r;
}
void* QTemporaryFile_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QTemporaryFile::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QTemporaryFile_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QTemporaryFile::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
int QTemporaryFile_open_default(void* self) {
	auto&& __r = ((QTemporaryFile*)self)->QTemporaryFile::open();
	return __r;
}
void* QTemporaryFile_metaObject_default(void* self) {
	auto&& __r = ((QTemporaryFile*)self)->QTemporaryFile::metaObject();
	return (void*)(__r);
}
char* QTemporaryFile_fileName_default(void* self, long long* __out_len) {
	auto&& __r = ((QTemporaryFile*)self)->QTemporaryFile::fileName();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
long long QTemporaryFile_size_default(void* self) {
	auto&& __r = ((QTemporaryFile*)self)->QTemporaryFile::size();
	return __r;
}
int QTemporaryFile_resize_default(void* self, long long sz) {
	auto&& __r = ((QTemporaryFile*)self)->QTemporaryFile::resize(sz);
	return __r;
}
int QTemporaryFile_permissions_default(void* self) {
	auto&& __r = ((QTemporaryFile*)self)->QTemporaryFile::permissions();
	return (int)(__r);
}
int QTemporaryFile_setPermissions_default(void* self, int permissionSpec) {
	auto&& __r = ((QTemporaryFile*)self)->QTemporaryFile::setPermissions(QFlags<QFileDevice::Permission>(permissionSpec));
	return __r;
}
void QTemporaryFile_close_default(void* self) {
	((QTemporaryFile*)self)->QTemporaryFile::close();
}
int QTemporaryFile_isSequential_default(void* self) {
	auto&& __r = ((QTemporaryFile*)self)->QTemporaryFile::isSequential();
	return __r;
}
long long QTemporaryFile_pos_default(void* self) {
	auto&& __r = ((QTemporaryFile*)self)->QTemporaryFile::pos();
	return __r;
}
int QTemporaryFile_seek_default(void* self, long long offset) {
	auto&& __r = ((QTemporaryFile*)self)->QTemporaryFile::seek(offset);
	return __r;
}
int QTemporaryFile_atEnd_default(void* self) {
	auto&& __r = ((QTemporaryFile*)self)->QTemporaryFile::atEnd();
	return __r;
}
int QTemporaryFile_reset_default(void* self) {
	auto&& __r = ((QTemporaryFile*)self)->QTemporaryFile::reset();
	return __r;
}
long long QTemporaryFile_bytesAvailable_default(void* self) {
	auto&& __r = ((QTemporaryFile*)self)->QTemporaryFile::bytesAvailable();
	return __r;
}
long long QTemporaryFile_bytesToWrite_default(void* self) {
	auto&& __r = ((QTemporaryFile*)self)->QTemporaryFile::bytesToWrite();
	return __r;
}
int QTemporaryFile_canReadLine_default(void* self) {
	auto&& __r = ((QTemporaryFile*)self)->QTemporaryFile::canReadLine();
	return __r;
}
int QTemporaryFile_waitForReadyRead_default(void* self, int msecs) {
	auto&& __r = ((QTemporaryFile*)self)->QTemporaryFile::waitForReadyRead(msecs);
	return __r;
}
int QTemporaryFile_waitForBytesWritten_default(void* self, int msecs) {
	auto&& __r = ((QTemporaryFile*)self)->QTemporaryFile::waitForBytesWritten(msecs);
	return __r;
}
long long QTemporaryFile_skipData_default(void* self, long long maxSize) {
	auto&& __r = ((QTemporaryFile__Go*)self)->cppgoBase_skipData(maxSize);
	return __r;
}
int QTemporaryFile_event_default(void* self, void* event) {
	auto&& __r = ((QTemporaryFile*)self)->QTemporaryFile::event((QEvent*)event);
	return __r;
}
int QTemporaryFile_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QTemporaryFile*)self)->QTemporaryFile::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QTemporaryFile_timerEvent_default(void* self, void* event) {
	((QTemporaryFile__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QTemporaryFile_childEvent_default(void* self, void* event) {
	((QTemporaryFile__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QTemporaryFile_customEvent_default(void* self, void* event) {
	((QTemporaryFile__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QTemporaryFile_connectNotify_default(void* self, void* signal) {
	((QTemporaryFile__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QTemporaryFile_disconnectNotify_default(void* self, void* signal) {
	((QTemporaryFile__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QTemporaryFile_connect_readyRead(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTemporaryFile*)self, &QTemporaryFile::readyRead, (QTemporaryFile*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTemporaryFile);
	__conns_QTemporaryFile[__id] = __c;
}
void QTemporaryFile_connect_channelReadyRead(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTemporaryFile*)self, &QTemporaryFile::channelReadyRead, (QTemporaryFile*)self, [__id](int channel){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, channel);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTemporaryFile);
	__conns_QTemporaryFile[__id] = __c;
}
void QTemporaryFile_connect_bytesWritten(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTemporaryFile*)self, &QTemporaryFile::bytesWritten, (QTemporaryFile*)self, [__id](qint64 bytes){
		if (cppgo_sig_void_longlong) ((void(*)(long long, long long))cppgo_sig_void_longlong)(__id, bytes);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTemporaryFile);
	__conns_QTemporaryFile[__id] = __c;
}
void QTemporaryFile_connect_channelBytesWritten(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTemporaryFile*)self, &QTemporaryFile::channelBytesWritten, (QTemporaryFile*)self, [__id](int channel, qint64 bytes){
		if (cppgo_sig_void_int_longlong) ((void(*)(long long, int, long long))cppgo_sig_void_int_longlong)(__id, channel, bytes);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTemporaryFile);
	__conns_QTemporaryFile[__id] = __c;
}
void QTemporaryFile_connect_aboutToClose(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTemporaryFile*)self, &QTemporaryFile::aboutToClose, (QTemporaryFile*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTemporaryFile);
	__conns_QTemporaryFile[__id] = __c;
}
void QTemporaryFile_connect_readChannelFinished(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTemporaryFile*)self, &QTemporaryFile::readChannelFinished, (QTemporaryFile*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTemporaryFile);
	__conns_QTemporaryFile[__id] = __c;
}
void QTemporaryFile_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTemporaryFile*)self, &QTemporaryFile::destroyed, (QTemporaryFile*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTemporaryFile);
	__conns_QTemporaryFile[__id] = __c;
}
void QTemporaryFile_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTemporaryFile*)self, &QTemporaryFile::destroyed, (QTemporaryFile*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTemporaryFile);
	__conns_QTemporaryFile[__id] = __c;
}
void QTemporaryFile_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTemporaryFile*)self, &QTemporaryFile::objectNameChanged, (QTemporaryFile*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTemporaryFile);
	__conns_QTemporaryFile[__id] = __c;
}
void QTemporaryFile__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QTemporaryFile);
		auto __it = __conns_QTemporaryFile.find(id);
		if (__it == __conns_QTemporaryFile.end()) return;
		__c = __it->second;
		__conns_QTemporaryFile.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QTemporaryFile__Go_new(long long id, void* parent) { return new QTemporaryFile__Go(id, (QObject*)parent); }
void QTemporaryFile_delete(void* self) { delete (QTemporaryFile*)self; }

}

#include <QtGui/qtransform.h>
#include <QtCore/qnamespace.h>
#include <QtGui/qpolygon.h>
#include <QtCore/qvariant.h>
#include <QtCore/qpoint.h>
#include <QtCore/qline.h>
#include <QtGui/qregion.h>
#include <QtGui/qpainterpath.h>
#include <QtCore/qrect.h>
#include "gen_qtransform.h"

extern "C" {

void* QTransform_new(int arg1) {
	return new QTransform((Qt::Initialization)(arg1));
}
void* QTransform_new2() {
	return new QTransform();
}
void* QTransform_new3(double h11, double h12, double h13, double h21, double h22, double h23, double h31, double h32, double h33) {
	return new QTransform(h11, h12, h13, h21, h22, h23, h31, h32, h33);
}
void* QTransform_new4(double h11, double h12, double h21, double h22, double dx, double dy) {
	return new QTransform(h11, h12, h21, h22, dx, dy);
}
void* QTransform_new5(void* other) {
	return new QTransform(*(QTransform*)other);
}
void* QTransform_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QTransform*)self)->operator=(*(QTransform*)arg1);
	return (void*)&(__r);
}
int QTransform_isAffine(void* self) {
	auto&& __r = ((QTransform*)self)->isAffine();
	return __r;
}
int QTransform_isIdentity(void* self) {
	auto&& __r = ((QTransform*)self)->isIdentity();
	return __r;
}
int QTransform_isInvertible(void* self) {
	auto&& __r = ((QTransform*)self)->isInvertible();
	return __r;
}
int QTransform_isScaling(void* self) {
	auto&& __r = ((QTransform*)self)->isScaling();
	return __r;
}
int QTransform_isRotating(void* self) {
	auto&& __r = ((QTransform*)self)->isRotating();
	return __r;
}
int QTransform_isTranslating(void* self) {
	auto&& __r = ((QTransform*)self)->isTranslating();
	return __r;
}
int QTransform_type(void* self) {
	auto&& __r = ((QTransform*)self)->type();
	return (int)(__r);
}
double QTransform_determinant(void* self) {
	auto&& __r = ((QTransform*)self)->determinant();
	return __r;
}
double QTransform_m11(void* self) {
	auto&& __r = ((QTransform*)self)->m11();
	return __r;
}
double QTransform_m12(void* self) {
	auto&& __r = ((QTransform*)self)->m12();
	return __r;
}
double QTransform_m13(void* self) {
	auto&& __r = ((QTransform*)self)->m13();
	return __r;
}
double QTransform_m21(void* self) {
	auto&& __r = ((QTransform*)self)->m21();
	return __r;
}
double QTransform_m22(void* self) {
	auto&& __r = ((QTransform*)self)->m22();
	return __r;
}
double QTransform_m23(void* self) {
	auto&& __r = ((QTransform*)self)->m23();
	return __r;
}
double QTransform_m31(void* self) {
	auto&& __r = ((QTransform*)self)->m31();
	return __r;
}
double QTransform_m32(void* self) {
	auto&& __r = ((QTransform*)self)->m32();
	return __r;
}
double QTransform_m33(void* self) {
	auto&& __r = ((QTransform*)self)->m33();
	return __r;
}
double QTransform_dx(void* self) {
	auto&& __r = ((QTransform*)self)->dx();
	return __r;
}
double QTransform_dy(void* self) {
	auto&& __r = ((QTransform*)self)->dy();
	return __r;
}
void QTransform_setMatrix(void* self, double m11, double m12, double m13, double m21, double m22, double m23, double m31, double m32, double m33) {
	((QTransform*)self)->setMatrix(m11, m12, m13, m21, m22, m23, m31, m32, m33);
}
void* QTransform_inverted(void* self) {
	auto&& __r = ((QTransform*)self)->inverted();
	return new QTransform(__r);
}
void* QTransform_adjoint(void* self) {
	auto&& __r = ((QTransform*)self)->adjoint();
	return new QTransform(__r);
}
void* QTransform_transposed(void* self) {
	auto&& __r = ((QTransform*)self)->transposed();
	return new QTransform(__r);
}
void* QTransform_translate(void* self, double dx, double dy) {
	auto&& __r = ((QTransform*)self)->translate(dx, dy);
	return (void*)&(__r);
}
void* QTransform_scale(void* self, double sx, double sy) {
	auto&& __r = ((QTransform*)self)->scale(sx, sy);
	return (void*)&(__r);
}
void* QTransform_shear(void* self, double sh, double sv) {
	auto&& __r = ((QTransform*)self)->shear(sh, sv);
	return (void*)&(__r);
}
void* QTransform_rotate(void* self, double a, int axis, double distanceToPlane) {
	auto&& __r = ((QTransform*)self)->rotate(a, (Qt::Axis)(axis), distanceToPlane);
	return (void*)&(__r);
}
void* QTransform_rotate2(void* self, double a) {
	auto&& __r = ((QTransform*)self)->rotate(a);
	return (void*)&(__r);
}
void* QTransform_rotateRadians(void* self, double a, int axis, double distanceToPlane) {
	auto&& __r = ((QTransform*)self)->rotateRadians(a, (Qt::Axis)(axis), distanceToPlane);
	return (void*)&(__r);
}
void* QTransform_rotateRadians2(void* self, double a) {
	auto&& __r = ((QTransform*)self)->rotateRadians(a);
	return (void*)&(__r);
}
int QTransform_squareToQuad(void* square, void* result) {
	auto&& __r = QTransform::squareToQuad(*(QPolygonF*)square, *(QTransform*)result);
	return __r;
}
int QTransform_quadToSquare(void* quad, void* result) {
	auto&& __r = QTransform::quadToSquare(*(QPolygonF*)quad, *(QTransform*)result);
	return __r;
}
int QTransform_quadToQuad(void* one, void* two, void* result) {
	auto&& __r = QTransform::quadToQuad(*(QPolygonF*)one, *(QPolygonF*)two, *(QTransform*)result);
	return __r;
}
int QTransform_OperatorEqual(void* self, void* arg1) {
	auto&& __r = ((QTransform*)self)->operator==(*(QTransform*)arg1);
	return __r;
}
int QTransform_OperatorNotEqual(void* self, void* arg1) {
	auto&& __r = ((QTransform*)self)->operator!=(*(QTransform*)arg1);
	return __r;
}
void* QTransform_OperatorMultiplyAssign(void* self, void* arg1) {
	auto&& __r = ((QTransform*)self)->operator*=(*(QTransform*)arg1);
	return (void*)&(__r);
}
void* QTransform_OperatorMultiply(void* self, void* o) {
	auto&& __r = ((QTransform*)self)->operator*(*(QTransform*)o);
	return new QTransform(__r);
}
void* QTransform_ToQVariant(void* self) {
	auto&& __r = static_cast<QVariant>(*(QTransform*)self);
	return new QVariant(__r);
}
void QTransform_reset(void* self) {
	((QTransform*)self)->reset();
}
void* QTransform_map(void* self, void* p) {
	auto&& __r = ((QTransform*)self)->map(*(QPoint*)p);
	return new QPoint(__r);
}
void* QTransform_map2(void* self, void* p) {
	auto&& __r = ((QTransform*)self)->map(*(QPointF*)p);
	return new QPointF(__r);
}
void* QTransform_map3(void* self, void* l) {
	auto&& __r = ((QTransform*)self)->map(*(QLine*)l);
	return new QLine(__r);
}
void* QTransform_map4(void* self, void* l) {
	auto&& __r = ((QTransform*)self)->map(*(QLineF*)l);
	return new QLineF(__r);
}
void* QTransform_map5(void* self, void* a) {
	auto&& __r = ((QTransform*)self)->map(*(QPolygonF*)a);
	return new QPolygonF(__r);
}
void* QTransform_map6(void* self, void* a) {
	auto&& __r = ((QTransform*)self)->map(*(QPolygon*)a);
	return new QPolygon(__r);
}
void* QTransform_map7(void* self, void* r) {
	auto&& __r = ((QTransform*)self)->map(*(QRegion*)r);
	return new QRegion(__r);
}
void* QTransform_map8(void* self, void* p) {
	auto&& __r = ((QTransform*)self)->map(*(QPainterPath*)p);
	return new QPainterPath(__r);
}
void* QTransform_mapToPolygon(void* self, void* r) {
	auto&& __r = ((QTransform*)self)->mapToPolygon(*(QRect*)r);
	return new QPolygon(__r);
}
void* QTransform_mapRect(void* self, void* arg1) {
	auto&& __r = ((QTransform*)self)->mapRect(*(QRect*)arg1);
	return new QRect(__r);
}
void* QTransform_mapRect2(void* self, void* arg1) {
	auto&& __r = ((QTransform*)self)->mapRect(*(QRectF*)arg1);
	return new QRectF(__r);
}
void QTransform_map9(void* self, int x, int y, int* tx, int* ty) {
	int tx_local{};
	int ty_local{};
	((QTransform*)self)->map(x, y, &tx_local, &ty_local);
	*tx = tx_local;
	*ty = ty_local;
}
void QTransform_map10(void* self, double x, double y, double* tx, double* ty) {
	qreal tx_local{};
	qreal ty_local{};
	((QTransform*)self)->map(x, y, &tx_local, &ty_local);
	*tx = tx_local;
	*ty = ty_local;
}
void* QTransform_OperatorMultiplyAssign2(void* self, double div) {
	auto&& __r = ((QTransform*)self)->operator*=(div);
	return (void*)&(__r);
}
void* QTransform_OperatorDivideAssign(void* self, double div) {
	auto&& __r = ((QTransform*)self)->operator/=(div);
	return (void*)&(__r);
}
void* QTransform_OperatorPlusAssign(void* self, double div) {
	auto&& __r = ((QTransform*)self)->operator+=(div);
	return (void*)&(__r);
}
void* QTransform_OperatorMinusAssign(void* self, double div) {
	auto&& __r = ((QTransform*)self)->operator-=(div);
	return (void*)&(__r);
}
void* QTransform_fromTranslate(double dx, double dy) {
	auto&& __r = QTransform::fromTranslate(dx, dy);
	return new QTransform(__r);
}
void* QTransform_fromScale(double dx, double dy) {
	auto&& __r = QTransform::fromScale(dx, dy);
	return new QTransform(__r);
}
void* QTransform_inverted2(void* self, int* invertible) {
	bool invertible_local{};
	auto&& __r = ((QTransform*)self)->inverted(&invertible_local);
	*invertible = invertible_local;
	return new QTransform(__r);
}
void* QTransform_rotate3(void* self, double a, int axis) {
	auto&& __r = ((QTransform*)self)->rotate(a, (Qt::Axis)(axis));
	return (void*)&(__r);
}
void* QTransform_rotateRadians3(void* self, double a, int axis) {
	auto&& __r = ((QTransform*)self)->rotateRadians(a, (Qt::Axis)(axis));
	return (void*)&(__r);
}
void* QTransform_OperatorMultiply2(void* self, double n) {
	auto&& __r = (*(QTransform*)self) * (n);
	return new QTransform(__r);
}
void* QTransform_OperatorDivide(void* self, double n) {
	auto&& __r = (*(QTransform*)self) / (n);
	return new QTransform(__r);
}
void* QTransform_OperatorPlus(void* self, double n) {
	auto&& __r = (*(QTransform*)self) + (n);
	return new QTransform(__r);
}
void* QTransform_OperatorMinus(void* self, double n) {
	auto&& __r = (*(QTransform*)self) - (n);
	return new QTransform(__r);
}
void* QTransform_toQVariant(void* self) {
	auto&& __r = QVariant::fromValue(*(QTransform*)self);
	return new QVariant(__r);
}
void QTransform_delete(void* self) { delete (QTransform*)self; }

}

#include <QtGui/qevent.h>
#include <QtCore/qnamespace.h>
#include <QtGui/qpointingdevice.h>
#include <QtCore/qpoint.h>
#include <QtCore/qobject.h>
#include <QtGui/qeventpoint.h>
#include <QtGui/qinputdevice.h>
#include <QtCore/qcoreevent.h>
#include "gen_qnativegestureevent.h"

extern "C" {

void* QNativeGestureEvent_base(void* self) { return (void*)static_cast<QSinglePointEvent*>((QNativeGestureEvent*)self); }
void QNativeGestureEvent_chain(void* self, void** out) {
	QNativeGestureEvent* __p = (QNativeGestureEvent*)self;
	QSinglePointEvent* __b0 = static_cast<QSinglePointEvent*>(__p); out[0] = (void*)__b0;
	QPointerEvent* __b1 = static_cast<QPointerEvent*>(__b0); out[1] = (void*)__b1;
	QInputEvent* __b2 = static_cast<QInputEvent*>(__b1); out[2] = (void*)__b2;
	QEvent* __b3 = static_cast<QEvent*>(__b2); out[3] = (void*)__b3;
}
void* QNativeGestureEvent_new(int type_, void* dev, void* localPos, void* scenePos, void* globalPos, double value, unsigned long long sequenceId, unsigned long long intArgument) {
	return new QNativeGestureEvent((Qt::NativeGestureType)(type_), (QPointingDevice*)dev, *(QPointF*)localPos, *(QPointF*)scenePos, *(QPointF*)globalPos, value, sequenceId, intArgument);
}
void* QNativeGestureEvent_new2(int type_, void* dev, int fingerCount, void* localPos, void* scenePos, void* globalPos, double value, void* delta) {
	return new QNativeGestureEvent((Qt::NativeGestureType)(type_), (QPointingDevice*)dev, fingerCount, *(QPointF*)localPos, *(QPointF*)scenePos, *(QPointF*)globalPos, value, *(QPointF*)delta);
}
void* QNativeGestureEvent_new3(int type_, void* dev, int fingerCount, void* localPos, void* scenePos, void* globalPos, double value, void* delta, unsigned long long sequenceId) {
	return new QNativeGestureEvent((Qt::NativeGestureType)(type_), (QPointingDevice*)dev, fingerCount, *(QPointF*)localPos, *(QPointF*)scenePos, *(QPointF*)globalPos, value, *(QPointF*)delta, sequenceId);
}
int QNativeGestureEvent_gestureType(void* self) {
	auto&& __r = ((QNativeGestureEvent*)self)->gestureType();
	return (int)(__r);
}
int QNativeGestureEvent_fingerCount(void* self) {
	auto&& __r = ((QNativeGestureEvent*)self)->fingerCount();
	return __r;
}
double QNativeGestureEvent_value(void* self) {
	auto&& __r = ((QNativeGestureEvent*)self)->value();
	return __r;
}
void* QNativeGestureEvent_delta(void* self) {
	auto&& __r = ((QNativeGestureEvent*)self)->delta();
	return new QPointF(__r);
}
void* QNativeGestureEvent_pos(void* self) {
	auto&& __r = ((QNativeGestureEvent*)self)->pos();
	return new QPoint(__r);
}
void* QNativeGestureEvent_globalPos(void* self) {
	auto&& __r = ((QNativeGestureEvent*)self)->globalPos();
	return new QPoint(__r);
}
void* QNativeGestureEvent_localPos(void* self) {
	auto&& __r = ((QNativeGestureEvent*)self)->localPos();
	return new QPointF(__r);
}
void* QNativeGestureEvent_windowPos(void* self) {
	auto&& __r = ((QNativeGestureEvent*)self)->windowPos();
	return new QPointF(__r);
}
void* QNativeGestureEvent_screenPos(void* self) {
	auto&& __r = ((QNativeGestureEvent*)self)->screenPos();
	return new QPointF(__r);
}
int QNativeGestureEvent_registerEventType() {
	auto&& __r = QNativeGestureEvent::registerEventType();
	return __r;
}
int QNativeGestureEvent_registerEventType2(int hint) {
	auto&& __r = QNativeGestureEvent::registerEventType(hint);
	return __r;
}
void QNativeGestureEvent_delete(void* self) { delete (QNativeGestureEvent*)self; }

}

#include <QtGui/qevent.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qstring.h>
#include <QtCore/qurl.h>
#include <QtCore/qfile.h>
#include <QtCore/qiodevicebase.h>
#include "gen_qfileopenevent.h"

extern "C" {

void* QFileOpenEvent_base(void* self) { return (void*)static_cast<QEvent*>((QFileOpenEvent*)self); }
void* QFileOpenEvent_new(const char* file, long long file_len) {
	QString file_qstr = QString::fromUtf8(file, file_len);
	return new QFileOpenEvent(file_qstr);
}
void* QFileOpenEvent_new2(void* url) {
	return new QFileOpenEvent(*(QUrl*)url);
}
char* QFileOpenEvent_file(void* self, long long* __out_len) {
	auto&& __r = ((QFileOpenEvent*)self)->file();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QFileOpenEvent_url(void* self) {
	auto&& __r = ((QFileOpenEvent*)self)->url();
	return new QUrl(__r);
}
int QFileOpenEvent_openFile(void* self, void* file, int flags) {
	auto&& __r = ((QFileOpenEvent*)self)->openFile(*(QFile*)file, QFlags<QIODeviceBase::OpenModeFlag>(flags));
	return __r;
}
int QFileOpenEvent_registerEventType() {
	auto&& __r = QFileOpenEvent::registerEventType();
	return __r;
}
int QFileOpenEvent_registerEventType2(int hint) {
	auto&& __r = QFileOpenEvent::registerEventType(hint);
	return __r;
}
void QFileOpenEvent_delete(void* self) { delete (QFileOpenEvent*)self; }

}

#include <QtGui/qpen.h>
#include <QtCore/qnamespace.h>
#include <QtGui/qcolor.h>
#include <QtGui/qbrush.h>
#include <QtCore/qvariant.h>
#include "gen_qpen.h"

extern "C" {

void* QPen_new() {
	return new QPen();
}
void* QPen_new2(int arg1) {
	return new QPen((Qt::PenStyle)(arg1));
}
void* QPen_new3(void* color) {
	return new QPen(*(QColor*)color);
}
void* QPen_new4(void* brush, double width) {
	return new QPen(*(QBrush*)brush, width);
}
void* QPen_new5(void* pen) {
	return new QPen(*(QPen*)pen);
}
void* QPen_new6(void* brush, double width, int s) {
	return new QPen(*(QBrush*)brush, width, (Qt::PenStyle)(s));
}
void* QPen_new7(void* brush, double width, int s, int c) {
	return new QPen(*(QBrush*)brush, width, (Qt::PenStyle)(s), (Qt::PenCapStyle)(c));
}
void* QPen_new8(void* brush, double width, int s, int c, int j) {
	return new QPen(*(QBrush*)brush, width, (Qt::PenStyle)(s), (Qt::PenCapStyle)(c), (Qt::PenJoinStyle)(j));
}
void* QPen_OperatorAssign(void* self, void* pen) {
	auto&& __r = ((QPen*)self)->operator=(*(QPen*)pen);
	return (void*)&(__r);
}
void QPen_swap(void* self, void* other) {
	((QPen*)self)->swap(*(QPen*)other);
}
void* QPen_OperatorAssign2(void* self, void* color) {
	auto&& __r = ((QPen*)self)->operator=(*(QColor*)color);
	return (void*)&(__r);
}
void* QPen_OperatorAssign3(void* self, int style) {
	auto&& __r = ((QPen*)self)->operator=((Qt::PenStyle)(style));
	return (void*)&(__r);
}
int QPen_style(void* self) {
	auto&& __r = ((QPen*)self)->style();
	return (int)(__r);
}
void QPen_setStyle(void* self, int arg1) {
	((QPen*)self)->setStyle((Qt::PenStyle)(arg1));
}
double* QPen_dashPattern(void* self, long long* __out_len) {
	auto&& __r = ((QPen*)self)->dashPattern();
	*__out_len = __r.size();
	double* __arr = (double*)malloc(sizeof(double) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = __e;
	return __arr;
}
void QPen_setDashPattern(void* self, double* pattern, long long pattern_len) {
	QList<double> pattern_list; pattern_list.reserve(pattern_len);
	for (long long __i = 0; __i < pattern_len; __i++) pattern_list.append(pattern[__i]);
	((QPen*)self)->setDashPattern(pattern_list);
}
double QPen_dashOffset(void* self) {
	auto&& __r = ((QPen*)self)->dashOffset();
	return __r;
}
void QPen_setDashOffset(void* self, double doffset) {
	((QPen*)self)->setDashOffset(doffset);
}
double QPen_miterLimit(void* self) {
	auto&& __r = ((QPen*)self)->miterLimit();
	return __r;
}
void QPen_setMiterLimit(void* self, double limit) {
	((QPen*)self)->setMiterLimit(limit);
}
double QPen_widthF(void* self) {
	auto&& __r = ((QPen*)self)->widthF();
	return __r;
}
void QPen_setWidthF(void* self, double width) {
	((QPen*)self)->setWidthF(width);
}
int QPen_width(void* self) {
	auto&& __r = ((QPen*)self)->width();
	return __r;
}
void QPen_setWidth(void* self, int width) {
	((QPen*)self)->setWidth(width);
}
void* QPen_color(void* self) {
	auto&& __r = ((QPen*)self)->color();
	return new QColor(__r);
}
void QPen_setColor(void* self, void* color) {
	((QPen*)self)->setColor(*(QColor*)color);
}
void* QPen_brush(void* self) {
	auto&& __r = ((QPen*)self)->brush();
	return new QBrush(__r);
}
void QPen_setBrush(void* self, void* brush) {
	((QPen*)self)->setBrush(*(QBrush*)brush);
}
int QPen_isSolid(void* self) {
	auto&& __r = ((QPen*)self)->isSolid();
	return __r;
}
int QPen_capStyle(void* self) {
	auto&& __r = ((QPen*)self)->capStyle();
	return (int)(__r);
}
void QPen_setCapStyle(void* self, int pcs) {
	((QPen*)self)->setCapStyle((Qt::PenCapStyle)(pcs));
}
int QPen_joinStyle(void* self) {
	auto&& __r = ((QPen*)self)->joinStyle();
	return (int)(__r);
}
void QPen_setJoinStyle(void* self, int pcs) {
	((QPen*)self)->setJoinStyle((Qt::PenJoinStyle)(pcs));
}
int QPen_isCosmetic(void* self) {
	auto&& __r = ((QPen*)self)->isCosmetic();
	return __r;
}
void QPen_setCosmetic(void* self, int cosmetic) {
	((QPen*)self)->setCosmetic(cosmetic);
}
int QPen_OperatorEqual(void* self, void* p) {
	auto&& __r = ((QPen*)self)->operator==(*(QPen*)p);
	return __r;
}
int QPen_OperatorNotEqual(void* self, void* p) {
	auto&& __r = ((QPen*)self)->operator!=(*(QPen*)p);
	return __r;
}
void* QPen_ToQVariant(void* self) {
	auto&& __r = static_cast<QVariant>(*(QPen*)self);
	return new QVariant(__r);
}
int QPen_isDetached(void* self) {
	auto&& __r = ((QPen*)self)->isDetached();
	return __r;
}
int QPen_OperatorEqual2(void* self, void* rhs) {
	auto&& __r = (*(QPen*)self) == (*(QColor*)rhs);
	return __r;
}
int QPen_OperatorNotEqual2(void* self, void* rhs) {
	auto&& __r = (*(QPen*)self) != (*(QColor*)rhs);
	return __r;
}
void* QPen_toQVariant(void* self) {
	auto&& __r = QVariant::fromValue(*(QPen*)self);
	return new QVariant(__r);
}
void QPen_delete(void* self) { delete (QPen*)self; }

}

#include <QtGui/qaccessible_base.h>
#include "gen_qaccessible__state.h"

extern "C" {

void* QAccessible__State_new() {
	return new QAccessible::State();
}
int QAccessible__State_OperatorEqual(void* self, void* second) {
	auto&& __r = (*(QAccessible::State*)self) == (*(QAccessible::State*)second);
	return __r;
}
unsigned long long QAccessible__State_disabled_get(void* self) {
	auto __r = ((QAccessible::State*)self)->disabled;
	return __r;
}
void QAccessible__State_disabled_set(void* self, unsigned long long disabled) {
	((QAccessible::State*)self)->disabled = disabled;
}
unsigned long long QAccessible__State_selected_get(void* self) {
	auto __r = ((QAccessible::State*)self)->selected;
	return __r;
}
void QAccessible__State_selected_set(void* self, unsigned long long selected) {
	((QAccessible::State*)self)->selected = selected;
}
unsigned long long QAccessible__State_focusable_get(void* self) {
	auto __r = ((QAccessible::State*)self)->focusable;
	return __r;
}
void QAccessible__State_focusable_set(void* self, unsigned long long focusable) {
	((QAccessible::State*)self)->focusable = focusable;
}
unsigned long long QAccessible__State_focused_get(void* self) {
	auto __r = ((QAccessible::State*)self)->focused;
	return __r;
}
void QAccessible__State_focused_set(void* self, unsigned long long focused) {
	((QAccessible::State*)self)->focused = focused;
}
unsigned long long QAccessible__State_pressed_get(void* self) {
	auto __r = ((QAccessible::State*)self)->pressed;
	return __r;
}
void QAccessible__State_pressed_set(void* self, unsigned long long pressed) {
	((QAccessible::State*)self)->pressed = pressed;
}
unsigned long long QAccessible__State_checkable_get(void* self) {
	auto __r = ((QAccessible::State*)self)->checkable;
	return __r;
}
void QAccessible__State_checkable_set(void* self, unsigned long long checkable) {
	((QAccessible::State*)self)->checkable = checkable;
}
unsigned long long QAccessible__State_checked_get(void* self) {
	auto __r = ((QAccessible::State*)self)->checked;
	return __r;
}
void QAccessible__State_checked_set(void* self, unsigned long long checked) {
	((QAccessible::State*)self)->checked = checked;
}
unsigned long long QAccessible__State_checkStateMixed_get(void* self) {
	auto __r = ((QAccessible::State*)self)->checkStateMixed;
	return __r;
}
void QAccessible__State_checkStateMixed_set(void* self, unsigned long long checkStateMixed) {
	((QAccessible::State*)self)->checkStateMixed = checkStateMixed;
}
unsigned long long QAccessible__State_readOnly_get(void* self) {
	auto __r = ((QAccessible::State*)self)->readOnly;
	return __r;
}
void QAccessible__State_readOnly_set(void* self, unsigned long long readOnly) {
	((QAccessible::State*)self)->readOnly = readOnly;
}
unsigned long long QAccessible__State_hotTracked_get(void* self) {
	auto __r = ((QAccessible::State*)self)->hotTracked;
	return __r;
}
void QAccessible__State_hotTracked_set(void* self, unsigned long long hotTracked) {
	((QAccessible::State*)self)->hotTracked = hotTracked;
}
unsigned long long QAccessible__State_defaultButton_get(void* self) {
	auto __r = ((QAccessible::State*)self)->defaultButton;
	return __r;
}
void QAccessible__State_defaultButton_set(void* self, unsigned long long defaultButton) {
	((QAccessible::State*)self)->defaultButton = defaultButton;
}
unsigned long long QAccessible__State_expanded_get(void* self) {
	auto __r = ((QAccessible::State*)self)->expanded;
	return __r;
}
void QAccessible__State_expanded_set(void* self, unsigned long long expanded) {
	((QAccessible::State*)self)->expanded = expanded;
}
unsigned long long QAccessible__State_collapsed_get(void* self) {
	auto __r = ((QAccessible::State*)self)->collapsed;
	return __r;
}
void QAccessible__State_collapsed_set(void* self, unsigned long long collapsed) {
	((QAccessible::State*)self)->collapsed = collapsed;
}
unsigned long long QAccessible__State_busy_get(void* self) {
	auto __r = ((QAccessible::State*)self)->busy;
	return __r;
}
void QAccessible__State_busy_set(void* self, unsigned long long busy) {
	((QAccessible::State*)self)->busy = busy;
}
unsigned long long QAccessible__State_expandable_get(void* self) {
	auto __r = ((QAccessible::State*)self)->expandable;
	return __r;
}
void QAccessible__State_expandable_set(void* self, unsigned long long expandable) {
	((QAccessible::State*)self)->expandable = expandable;
}
unsigned long long QAccessible__State_marqueed_get(void* self) {
	auto __r = ((QAccessible::State*)self)->marqueed;
	return __r;
}
void QAccessible__State_marqueed_set(void* self, unsigned long long marqueed) {
	((QAccessible::State*)self)->marqueed = marqueed;
}
unsigned long long QAccessible__State_animated_get(void* self) {
	auto __r = ((QAccessible::State*)self)->animated;
	return __r;
}
void QAccessible__State_animated_set(void* self, unsigned long long animated) {
	((QAccessible::State*)self)->animated = animated;
}
unsigned long long QAccessible__State_invisible_get(void* self) {
	auto __r = ((QAccessible::State*)self)->invisible;
	return __r;
}
void QAccessible__State_invisible_set(void* self, unsigned long long invisible) {
	((QAccessible::State*)self)->invisible = invisible;
}
unsigned long long QAccessible__State_offscreen_get(void* self) {
	auto __r = ((QAccessible::State*)self)->offscreen;
	return __r;
}
void QAccessible__State_offscreen_set(void* self, unsigned long long offscreen) {
	((QAccessible::State*)self)->offscreen = offscreen;
}
unsigned long long QAccessible__State_sizeable_get(void* self) {
	auto __r = ((QAccessible::State*)self)->sizeable;
	return __r;
}
void QAccessible__State_sizeable_set(void* self, unsigned long long sizeable) {
	((QAccessible::State*)self)->sizeable = sizeable;
}
unsigned long long QAccessible__State_movable_get(void* self) {
	auto __r = ((QAccessible::State*)self)->movable;
	return __r;
}
void QAccessible__State_movable_set(void* self, unsigned long long movable) {
	((QAccessible::State*)self)->movable = movable;
}
unsigned long long QAccessible__State_selfVoicing_get(void* self) {
	auto __r = ((QAccessible::State*)self)->selfVoicing;
	return __r;
}
void QAccessible__State_selfVoicing_set(void* self, unsigned long long selfVoicing) {
	((QAccessible::State*)self)->selfVoicing = selfVoicing;
}
unsigned long long QAccessible__State_selectable_get(void* self) {
	auto __r = ((QAccessible::State*)self)->selectable;
	return __r;
}
void QAccessible__State_selectable_set(void* self, unsigned long long selectable) {
	((QAccessible::State*)self)->selectable = selectable;
}
unsigned long long QAccessible__State_linked_get(void* self) {
	auto __r = ((QAccessible::State*)self)->linked;
	return __r;
}
void QAccessible__State_linked_set(void* self, unsigned long long linked) {
	((QAccessible::State*)self)->linked = linked;
}
unsigned long long QAccessible__State_traversed_get(void* self) {
	auto __r = ((QAccessible::State*)self)->traversed;
	return __r;
}
void QAccessible__State_traversed_set(void* self, unsigned long long traversed) {
	((QAccessible::State*)self)->traversed = traversed;
}
unsigned long long QAccessible__State_multiSelectable_get(void* self) {
	auto __r = ((QAccessible::State*)self)->multiSelectable;
	return __r;
}
void QAccessible__State_multiSelectable_set(void* self, unsigned long long multiSelectable) {
	((QAccessible::State*)self)->multiSelectable = multiSelectable;
}
unsigned long long QAccessible__State_extSelectable_get(void* self) {
	auto __r = ((QAccessible::State*)self)->extSelectable;
	return __r;
}
void QAccessible__State_extSelectable_set(void* self, unsigned long long extSelectable) {
	((QAccessible::State*)self)->extSelectable = extSelectable;
}
unsigned long long QAccessible__State_passwordEdit_get(void* self) {
	auto __r = ((QAccessible::State*)self)->passwordEdit;
	return __r;
}
void QAccessible__State_passwordEdit_set(void* self, unsigned long long passwordEdit) {
	((QAccessible::State*)self)->passwordEdit = passwordEdit;
}
unsigned long long QAccessible__State_hasPopup_get(void* self) {
	auto __r = ((QAccessible::State*)self)->hasPopup;
	return __r;
}
void QAccessible__State_hasPopup_set(void* self, unsigned long long hasPopup) {
	((QAccessible::State*)self)->hasPopup = hasPopup;
}
unsigned long long QAccessible__State_modal_get(void* self) {
	auto __r = ((QAccessible::State*)self)->modal;
	return __r;
}
void QAccessible__State_modal_set(void* self, unsigned long long modal) {
	((QAccessible::State*)self)->modal = modal;
}
unsigned long long QAccessible__State_active_get(void* self) {
	auto __r = ((QAccessible::State*)self)->active;
	return __r;
}
void QAccessible__State_active_set(void* self, unsigned long long active) {
	((QAccessible::State*)self)->active = active;
}
unsigned long long QAccessible__State_invalid_get(void* self) {
	auto __r = ((QAccessible::State*)self)->invalid;
	return __r;
}
void QAccessible__State_invalid_set(void* self, unsigned long long invalid) {
	((QAccessible::State*)self)->invalid = invalid;
}
unsigned long long QAccessible__State_editable_get(void* self) {
	auto __r = ((QAccessible::State*)self)->editable;
	return __r;
}
void QAccessible__State_editable_set(void* self, unsigned long long editable) {
	((QAccessible::State*)self)->editable = editable;
}
unsigned long long QAccessible__State_multiLine_get(void* self) {
	auto __r = ((QAccessible::State*)self)->multiLine;
	return __r;
}
void QAccessible__State_multiLine_set(void* self, unsigned long long multiLine) {
	((QAccessible::State*)self)->multiLine = multiLine;
}
unsigned long long QAccessible__State_selectableText_get(void* self) {
	auto __r = ((QAccessible::State*)self)->selectableText;
	return __r;
}
void QAccessible__State_selectableText_set(void* self, unsigned long long selectableText) {
	((QAccessible::State*)self)->selectableText = selectableText;
}
unsigned long long QAccessible__State_supportsAutoCompletion_get(void* self) {
	auto __r = ((QAccessible::State*)self)->supportsAutoCompletion;
	return __r;
}
void QAccessible__State_supportsAutoCompletion_set(void* self, unsigned long long supportsAutoCompletion) {
	((QAccessible::State*)self)->supportsAutoCompletion = supportsAutoCompletion;
}
unsigned long long QAccessible__State_searchEdit_get(void* self) {
	auto __r = ((QAccessible::State*)self)->searchEdit;
	return __r;
}
void QAccessible__State_searchEdit_set(void* self, unsigned long long searchEdit) {
	((QAccessible::State*)self)->searchEdit = searchEdit;
}
unsigned long long QAccessible__State_qt_reserved_get(void* self) {
	auto __r = ((QAccessible::State*)self)->qt_reserved;
	return __r;
}
void QAccessible__State_qt_reserved_set(void* self, unsigned long long qt_reserved) {
	((QAccessible::State*)self)->qt_reserved = qt_reserved;
}
void QAccessible__State_delete(void* self) { delete (QAccessible::State*)self; }

}

#include <QtGui/qaccessiblebridge.h>
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
#include "gen_qaccessiblebridgeplugin.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QAccessibleBridgePlugin__Go : public QAccessibleBridgePlugin {
public:
	long long __goId;
	unsigned char __goOn[9] = {};
	template<typename... A> QAccessibleBridgePlugin__Go(long long id, A&&... a) : QAccessibleBridgePlugin(std::forward<A>(a)...), __goId(id) {}
	using QAccessibleBridgePlugin::sender;
	using QAccessibleBridgePlugin::senderSignalIndex;
	using QAccessibleBridgePlugin::receivers;
	using QAccessibleBridgePlugin::isSignalConnected;
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QAccessibleBridgePlugin::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QAccessibleBridgePlugin::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QAccessibleBridgePlugin::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QAccessibleBridgePlugin::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QAccessibleBridgePlugin::disconnectNotify(__a0); }
	QAccessibleBridge* create(const QString& key) override {
		int __h = 0;
		QByteArray key_ba = key.toUtf8();
		if (__goOn[0] && cppgo_virt_voidp_constcharp_longlong) { void* __r = ((void*(*)(long long, long long, const char*, long long, int*))cppgo_virt_voidp_constcharp_longlong)(__goId, -626457835795418115LL, (char*)key_ba.constData(), key_ba.size(), &__h); if (__h) return (QAccessibleBridge*)__r; }
		return nullptr;
	}
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QAccessibleBridgePlugin::metaObject();
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QAccessibleBridgePlugin::event(event);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QAccessibleBridgePlugin::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QAccessibleBridgePlugin::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QAccessibleBridgePlugin::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QAccessibleBridgePlugin::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QAccessibleBridgePlugin::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QAccessibleBridgePlugin::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QAccessibleBridgePlugin;
static std::mutex __connsMu_QAccessibleBridgePlugin;

extern "C" {

void* QAccessibleBridgePlugin_base(void* self) { return (void*)static_cast<QObject*>((QAccessibleBridgePlugin*)self); }
void QAccessibleBridgePlugin__set_override(void* self, int idx) { ((QAccessibleBridgePlugin__Go*)self)->__goOn[idx] = 1; }
void* QAccessibleBridgePlugin_fromQObject(void* o) { return (void*)qobject_cast<QAccessibleBridgePlugin*>((QObject*)o); }
char* QAccessibleBridgePlugin_tr(const char* s, long long* __out_len) {
	auto&& __r = QAccessibleBridgePlugin::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QAccessibleBridgePlugin_create(void* self, const char* key, long long key_len) {
	QString key_qstr = QString::fromUtf8(key, key_len);
	auto&& __r = ((QAccessibleBridgePlugin*)self)->create(key_qstr);
	return (void*)(__r);
}
void* QAccessibleBridgePlugin_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QAccessibleBridgePlugin::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QAccessibleBridgePlugin_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QAccessibleBridgePlugin::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QAccessibleBridgePlugin_sender(void* self) {
	auto&& __r = ((QAccessibleBridgePlugin__Go*)self)->sender();
	return (void*)(__r);
}
int QAccessibleBridgePlugin_senderSignalIndex(void* self) {
	auto&& __r = ((QAccessibleBridgePlugin__Go*)self)->senderSignalIndex();
	return __r;
}
int QAccessibleBridgePlugin_receivers(void* self, const char* signal) {
	auto&& __r = ((QAccessibleBridgePlugin__Go*)self)->receivers(signal);
	return __r;
}
int QAccessibleBridgePlugin_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QAccessibleBridgePlugin__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QAccessibleBridgePlugin_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QAccessibleBridgePlugin::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QAccessibleBridgePlugin_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QAccessibleBridgePlugin::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QAccessibleBridgePlugin_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QAccessibleBridgePlugin::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QAccessibleBridgePlugin_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QAccessibleBridgePlugin::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QAccessibleBridgePlugin_metaObject_default(void* self) {
	auto&& __r = ((QAccessibleBridgePlugin*)self)->QAccessibleBridgePlugin::metaObject();
	return (void*)(__r);
}
int QAccessibleBridgePlugin_event_default(void* self, void* event) {
	auto&& __r = ((QAccessibleBridgePlugin*)self)->QAccessibleBridgePlugin::event((QEvent*)event);
	return __r;
}
int QAccessibleBridgePlugin_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QAccessibleBridgePlugin*)self)->QAccessibleBridgePlugin::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QAccessibleBridgePlugin_timerEvent_default(void* self, void* event) {
	((QAccessibleBridgePlugin__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QAccessibleBridgePlugin_childEvent_default(void* self, void* event) {
	((QAccessibleBridgePlugin__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QAccessibleBridgePlugin_customEvent_default(void* self, void* event) {
	((QAccessibleBridgePlugin__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QAccessibleBridgePlugin_connectNotify_default(void* self, void* signal) {
	((QAccessibleBridgePlugin__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QAccessibleBridgePlugin_disconnectNotify_default(void* self, void* signal) {
	((QAccessibleBridgePlugin__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QAccessibleBridgePlugin_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAccessibleBridgePlugin*)self, &QAccessibleBridgePlugin::destroyed, (QAccessibleBridgePlugin*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAccessibleBridgePlugin);
	__conns_QAccessibleBridgePlugin[__id] = __c;
}
void QAccessibleBridgePlugin_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAccessibleBridgePlugin*)self, &QAccessibleBridgePlugin::destroyed, (QAccessibleBridgePlugin*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAccessibleBridgePlugin);
	__conns_QAccessibleBridgePlugin[__id] = __c;
}
void QAccessibleBridgePlugin_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAccessibleBridgePlugin*)self, &QAccessibleBridgePlugin::objectNameChanged, (QAccessibleBridgePlugin*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAccessibleBridgePlugin);
	__conns_QAccessibleBridgePlugin[__id] = __c;
}
void QAccessibleBridgePlugin__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QAccessibleBridgePlugin);
		auto __it = __conns_QAccessibleBridgePlugin.find(id);
		if (__it == __conns_QAccessibleBridgePlugin.end()) return;
		__c = __it->second;
		__conns_QAccessibleBridgePlugin.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QAccessibleBridgePlugin__Go_new(long long id, void* parent) { return new QAccessibleBridgePlugin__Go(id, (QObject*)parent); }
void QAccessibleBridgePlugin_delete(void* self) { delete (QAccessibleBridgePlugin*)self; }

}

#include <QtGui/qgenericpluginfactory.h>
#include <QtCore/qstring.h>
#include <QtCore/qobject.h>
#include "gen_qgenericpluginfactory.h"

extern "C" {

char** QGenericPluginFactory_keys(long long* __out_len) {
	auto&& __r = QGenericPluginFactory::keys();
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
void* QGenericPluginFactory_create(const char* arg1, long long arg1_len, const char* arg2, long long arg2_len) {
	QString arg1_qstr = QString::fromUtf8(arg1, arg1_len);
	QString arg2_qstr = QString::fromUtf8(arg2, arg2_len);
	auto&& __r = QGenericPluginFactory::create(arg1_qstr, arg2_qstr);
	return (void*)(__r);
}
void QGenericPluginFactory_delete(void* self) { delete (QGenericPluginFactory*)self; }

}

#include <QtGui/qpaintengine.h>
#include <QtGui/qpaintdevice.h>
#include <QtCore/qrect.h>
#include <QtCore/qline.h>
#include <QtGui/qpainterpath.h>
#include <QtCore/qpoint.h>
#include <QtGui/qpixmap.h>
#include <QtGui/qimage.h>
#include <QtCore/qnamespace.h>
#include <QtGui/qregion.h>
#include <QtGui/qpainter.h>
#include <QtCore/qsize.h>
#include "gen_qpaintengine.h"
#include "gen_purego_cb.h"

class QPaintEngine__Go : public QPaintEngine {
public:
	long long __goId;
	unsigned char __goOn[22] = {};
	template<typename... A> QPaintEngine__Go(long long id, A&&... a) : QPaintEngine(std::forward<A>(a)...), __goId(id) {}
	bool begin(QPaintDevice* pdev) override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, -7312564458357607458LL, (void*)pdev, &__h); if (__h) return __r; }
		return bool{};
	}
	bool end() override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -4400013430294884342LL, &__h); if (__h) return __r; }
		return bool{};
	}
	void updateState(const QPaintEngineState& state) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7692482089134926569LL, (void*)&state, &__h); if (__h) return; }
		return;
	}
	void drawPixmap(const QRectF& r, const QPixmap& pm, const QRectF& sr) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_voidp_voidp_voidp) { ((void(*)(long long, long long, void*, void*, void*, int*))cppgo_virt_void_voidp_voidp_voidp)(__goId, -6262278833611863026LL, (void*)&r, (void*)&pm, (void*)&sr, &__h); if (__h) return; }
		return;
	}
	QPaintEngine::Type type() const override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -6371403872082420691LL, &__h); if (__h) return (QPaintEngine::Type)(__r); }
		return QPaintEngine::Type{};
	}
	void drawRects(const QRect* rects, int rectCount) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp_int) { ((void(*)(long long, long long, void*, int, int*))cppgo_virt_void_voidp_int)(__goId, 6141870017934704802LL, (void*)rects, rectCount, &__h); if (__h) return; }
		QPaintEngine::drawRects(rects, rectCount);
	}
	void drawRects(const QRectF* rects, int rectCount) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp_int) { ((void(*)(long long, long long, void*, int, int*))cppgo_virt_void_voidp_int)(__goId, 8374868496907056304LL, (void*)rects, rectCount, &__h); if (__h) return; }
		QPaintEngine::drawRects(rects, rectCount);
	}
	void drawLines(const QLine* lines, int lineCount) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp_int) { ((void(*)(long long, long long, void*, int, int*))cppgo_virt_void_voidp_int)(__goId, 3290090450242779774LL, (void*)lines, lineCount, &__h); if (__h) return; }
		QPaintEngine::drawLines(lines, lineCount);
	}
	void drawLines(const QLineF* lines, int lineCount) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void_voidp_int) { ((void(*)(long long, long long, void*, int, int*))cppgo_virt_void_voidp_int)(__goId, 3888932263865552676LL, (void*)lines, lineCount, &__h); if (__h) return; }
		QPaintEngine::drawLines(lines, lineCount);
	}
	void drawEllipse(const QRectF& r) override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4168977152563157681LL, (void*)&r, &__h); if (__h) return; }
		QPaintEngine::drawEllipse(r);
	}
	void drawEllipse(const QRect& r) override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 8175511122464863079LL, (void*)&r, &__h); if (__h) return; }
		QPaintEngine::drawEllipse(r);
	}
	void drawPath(const QPainterPath& path) override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 2689733668603817204LL, (void*)&path, &__h); if (__h) return; }
		QPaintEngine::drawPath(path);
	}
	void drawPoints(const QPointF* points, int pointCount) override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_void_voidp_int) { ((void(*)(long long, long long, void*, int, int*))cppgo_virt_void_voidp_int)(__goId, 6207679886378890104LL, (void*)points, pointCount, &__h); if (__h) return; }
		QPaintEngine::drawPoints(points, pointCount);
	}
	void drawPoints(const QPoint* points, int pointCount) override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_void_voidp_int) { ((void(*)(long long, long long, void*, int, int*))cppgo_virt_void_voidp_int)(__goId, 6211818076911756990LL, (void*)points, pointCount, &__h); if (__h) return; }
		QPaintEngine::drawPoints(points, pointCount);
	}
	void drawPolygon(const QPointF* points, int pointCount, QPaintEngine::PolygonDrawMode mode) override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_void_voidp_int_int) { ((void(*)(long long, long long, void*, int, int, int*))cppgo_virt_void_voidp_int_int)(__goId, 8440328321123084003LL, (void*)points, pointCount, (int)(mode), &__h); if (__h) return; }
		QPaintEngine::drawPolygon(points, pointCount, mode);
	}
	void drawPolygon(const QPoint* points, int pointCount, QPaintEngine::PolygonDrawMode mode) override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_void_voidp_int_int) { ((void(*)(long long, long long, void*, int, int, int*))cppgo_virt_void_voidp_int_int)(__goId, 4524770546924598563LL, (void*)points, pointCount, (int)(mode), &__h); if (__h) return; }
		QPaintEngine::drawPolygon(points, pointCount, mode);
	}
	void drawTextItem(const QPointF& p, const QTextItem& textItem) override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_void_voidp_voidp) { ((void(*)(long long, long long, void*, void*, int*))cppgo_virt_void_voidp_voidp)(__goId, -6193357004012257525LL, (void*)&p, (void*)&textItem, &__h); if (__h) return; }
		QPaintEngine::drawTextItem(p, textItem);
	}
	void drawTiledPixmap(const QRectF& r, const QPixmap& pixmap, const QPointF& s) override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_void_voidp_voidp_voidp) { ((void(*)(long long, long long, void*, void*, void*, int*))cppgo_virt_void_voidp_voidp_voidp)(__goId, 8178762518988445540LL, (void*)&r, (void*)&pixmap, (void*)&s, &__h); if (__h) return; }
		QPaintEngine::drawTiledPixmap(r, pixmap, s);
	}
	void drawImage(const QRectF& r, const QImage& pm, const QRectF& sr, QFlags<Qt::ImageConversionFlag> flags) override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_void_voidp_voidp_voidp_int) { ((void(*)(long long, long long, void*, void*, void*, int, int*))cppgo_virt_void_voidp_voidp_voidp_int)(__goId, 5783836824819670344LL, (void*)&r, (void*)&pm, (void*)&sr, (int)(flags), &__h); if (__h) return; }
		QPaintEngine::drawImage(r, pm, sr, flags);
	}
	QPoint coordinateOffset() const override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, -4718660678191360312LL, &__h); if (__h) return *(QPoint*)__r; }
		return QPaintEngine::coordinateOffset();
	}
	QPixmap createPixmap(QSize size) override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 5331809331223620608LL, (void*)&size, &__h); if (__h) return *(QPixmap*)__r; }
		return QPaintEngine::createPixmap(size);
	}
	QPixmap createPixmapFromImage(QImage image, QFlags<Qt::ImageConversionFlag> flags) override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_voidp_voidp_int) { void* __r = ((void*(*)(long long, long long, void*, int, int*))cppgo_virt_voidp_voidp_int)(__goId, -6968771525665389363LL, (void*)&image, (int)(flags), &__h); if (__h) return *(QPixmap*)__r; }
		return QPaintEngine::createPixmapFromImage(image, flags);
	}
};

extern "C" {

void QPaintEngine__set_override(void* self, int idx) { ((QPaintEngine__Go*)self)->__goOn[idx] = 1; }
int QPaintEngine_isActive(void* self) {
	auto&& __r = ((QPaintEngine*)self)->isActive();
	return __r;
}
void QPaintEngine_setActive(void* self, int newState) {
	((QPaintEngine*)self)->setActive(newState);
}
int QPaintEngine_begin(void* self, void* pdev) {
	auto&& __r = ((QPaintEngine*)self)->begin((QPaintDevice*)pdev);
	return __r;
}
int QPaintEngine_end(void* self) {
	auto&& __r = ((QPaintEngine*)self)->end();
	return __r;
}
void QPaintEngine_updateState(void* self, void* state) {
	((QPaintEngine*)self)->updateState(*(QPaintEngineState*)state);
}
void QPaintEngine_drawRects(void* self, void* rects, int rectCount) {
	((QPaintEngine*)self)->drawRects((QRect*)rects, rectCount);
}
void QPaintEngine_drawRects2(void* self, void* rects, int rectCount) {
	((QPaintEngine*)self)->drawRects((QRectF*)rects, rectCount);
}
void QPaintEngine_drawLines(void* self, void* lines, int lineCount) {
	((QPaintEngine*)self)->drawLines((QLine*)lines, lineCount);
}
void QPaintEngine_drawLines2(void* self, void* lines, int lineCount) {
	((QPaintEngine*)self)->drawLines((QLineF*)lines, lineCount);
}
void QPaintEngine_drawEllipse(void* self, void* r) {
	((QPaintEngine*)self)->drawEllipse(*(QRectF*)r);
}
void QPaintEngine_drawEllipse2(void* self, void* r) {
	((QPaintEngine*)self)->drawEllipse(*(QRect*)r);
}
void QPaintEngine_drawPath(void* self, void* path) {
	((QPaintEngine*)self)->drawPath(*(QPainterPath*)path);
}
void QPaintEngine_drawPoints(void* self, void* points, int pointCount) {
	((QPaintEngine*)self)->drawPoints((QPointF*)points, pointCount);
}
void QPaintEngine_drawPoints2(void* self, void* points, int pointCount) {
	((QPaintEngine*)self)->drawPoints((QPoint*)points, pointCount);
}
void QPaintEngine_drawPolygon(void* self, void* points, int pointCount, int mode) {
	((QPaintEngine*)self)->drawPolygon((QPointF*)points, pointCount, (QPaintEngine::PolygonDrawMode)(mode));
}
void QPaintEngine_drawPolygon2(void* self, void* points, int pointCount, int mode) {
	((QPaintEngine*)self)->drawPolygon((QPoint*)points, pointCount, (QPaintEngine::PolygonDrawMode)(mode));
}
void QPaintEngine_drawPixmap(void* self, void* r, void* pm, void* sr) {
	((QPaintEngine*)self)->drawPixmap(*(QRectF*)r, *(QPixmap*)pm, *(QRectF*)sr);
}
void QPaintEngine_drawTextItem(void* self, void* p, void* textItem) {
	((QPaintEngine*)self)->drawTextItem(*(QPointF*)p, *(QTextItem*)textItem);
}
void QPaintEngine_drawTiledPixmap(void* self, void* r, void* pixmap, void* s) {
	((QPaintEngine*)self)->drawTiledPixmap(*(QRectF*)r, *(QPixmap*)pixmap, *(QPointF*)s);
}
void QPaintEngine_drawImage(void* self, void* r, void* pm, void* sr, int flags) {
	((QPaintEngine*)self)->drawImage(*(QRectF*)r, *(QImage*)pm, *(QRectF*)sr, QFlags<Qt::ImageConversionFlag>(flags));
}
void QPaintEngine_setPaintDevice(void* self, void* device) {
	((QPaintEngine*)self)->setPaintDevice((QPaintDevice*)device);
}
void* QPaintEngine_paintDevice(void* self) {
	auto&& __r = ((QPaintEngine*)self)->paintDevice();
	return (void*)(__r);
}
void QPaintEngine_setSystemClip(void* self, void* baseClip) {
	((QPaintEngine*)self)->setSystemClip(*(QRegion*)baseClip);
}
void* QPaintEngine_systemClip(void* self) {
	auto&& __r = ((QPaintEngine*)self)->systemClip();
	return new QRegion(__r);
}
void QPaintEngine_setSystemRect(void* self, void* rect) {
	((QPaintEngine*)self)->setSystemRect(*(QRect*)rect);
}
void* QPaintEngine_systemRect(void* self) {
	auto&& __r = ((QPaintEngine*)self)->systemRect();
	return new QRect(__r);
}
void* QPaintEngine_coordinateOffset(void* self) {
	auto&& __r = ((QPaintEngine*)self)->coordinateOffset();
	return new QPoint(__r);
}
int QPaintEngine_type(void* self) {
	auto&& __r = ((QPaintEngine*)self)->type();
	return (int)(__r);
}
void QPaintEngine_fix_neg_rect(void* self, int* x, int* y, int* w, int* h) {
	int x_local{};
	int y_local{};
	int w_local{};
	int h_local{};
	((QPaintEngine*)self)->fix_neg_rect(&x_local, &y_local, &w_local, &h_local);
	*x = x_local;
	*y = y_local;
	*w = w_local;
	*h = h_local;
}
int QPaintEngine_testDirty(void* self, int df) {
	auto&& __r = ((QPaintEngine*)self)->testDirty(QFlags<QPaintEngine::DirtyFlag>(df));
	return __r;
}
void QPaintEngine_setDirty(void* self, int df) {
	((QPaintEngine*)self)->setDirty(QFlags<QPaintEngine::DirtyFlag>(df));
}
void QPaintEngine_clearDirty(void* self, int df) {
	((QPaintEngine*)self)->clearDirty(QFlags<QPaintEngine::DirtyFlag>(df));
}
int QPaintEngine_hasFeature(void* self, int feature) {
	auto&& __r = ((QPaintEngine*)self)->hasFeature(QFlags<QPaintEngine::PaintEngineFeature>(feature));
	return __r;
}
void* QPaintEngine_painter(void* self) {
	auto&& __r = ((QPaintEngine*)self)->painter();
	return (void*)(__r);
}
void QPaintEngine_syncState(void* self) {
	((QPaintEngine*)self)->syncState();
}
int QPaintEngine_isExtended(void* self) {
	auto&& __r = ((QPaintEngine*)self)->isExtended();
	return __r;
}
void* QPaintEngine_createPixmap(void* self, void* size) {
	auto&& __r = ((QPaintEngine*)self)->createPixmap(*(QSize*)size);
	return new QPixmap(__r);
}
void* QPaintEngine_createPixmapFromImage(void* self, void* image, int flags) {
	auto&& __r = ((QPaintEngine*)self)->createPixmapFromImage(*(QImage*)image, QFlags<Qt::ImageConversionFlag>(flags));
	return new QPixmap(__r);
}
void QPaintEngine_drawRects_default(void* self, void* rects, int rectCount) {
	((QPaintEngine*)self)->QPaintEngine::drawRects((QRect*)rects, rectCount);
}
void QPaintEngine_drawRects_default2(void* self, void* rects, int rectCount) {
	((QPaintEngine*)self)->QPaintEngine::drawRects((QRectF*)rects, rectCount);
}
void QPaintEngine_drawLines_default(void* self, void* lines, int lineCount) {
	((QPaintEngine*)self)->QPaintEngine::drawLines((QLine*)lines, lineCount);
}
void QPaintEngine_drawLines_default2(void* self, void* lines, int lineCount) {
	((QPaintEngine*)self)->QPaintEngine::drawLines((QLineF*)lines, lineCount);
}
void QPaintEngine_drawEllipse_default(void* self, void* r) {
	((QPaintEngine*)self)->QPaintEngine::drawEllipse(*(QRectF*)r);
}
void QPaintEngine_drawEllipse_default2(void* self, void* r) {
	((QPaintEngine*)self)->QPaintEngine::drawEllipse(*(QRect*)r);
}
void QPaintEngine_drawPath_default(void* self, void* path) {
	((QPaintEngine*)self)->QPaintEngine::drawPath(*(QPainterPath*)path);
}
void QPaintEngine_drawPoints_default(void* self, void* points, int pointCount) {
	((QPaintEngine*)self)->QPaintEngine::drawPoints((QPointF*)points, pointCount);
}
void QPaintEngine_drawPoints_default2(void* self, void* points, int pointCount) {
	((QPaintEngine*)self)->QPaintEngine::drawPoints((QPoint*)points, pointCount);
}
void QPaintEngine_drawPolygon_default(void* self, void* points, int pointCount, int mode) {
	((QPaintEngine*)self)->QPaintEngine::drawPolygon((QPointF*)points, pointCount, (QPaintEngine::PolygonDrawMode)(mode));
}
void QPaintEngine_drawPolygon_default2(void* self, void* points, int pointCount, int mode) {
	((QPaintEngine*)self)->QPaintEngine::drawPolygon((QPoint*)points, pointCount, (QPaintEngine::PolygonDrawMode)(mode));
}
void QPaintEngine_drawTextItem_default(void* self, void* p, void* textItem) {
	((QPaintEngine*)self)->QPaintEngine::drawTextItem(*(QPointF*)p, *(QTextItem*)textItem);
}
void QPaintEngine_drawTiledPixmap_default(void* self, void* r, void* pixmap, void* s) {
	((QPaintEngine*)self)->QPaintEngine::drawTiledPixmap(*(QRectF*)r, *(QPixmap*)pixmap, *(QPointF*)s);
}
void QPaintEngine_drawImage_default(void* self, void* r, void* pm, void* sr, int flags) {
	((QPaintEngine*)self)->QPaintEngine::drawImage(*(QRectF*)r, *(QImage*)pm, *(QRectF*)sr, QFlags<Qt::ImageConversionFlag>(flags));
}
void* QPaintEngine_coordinateOffset_default(void* self) {
	auto&& __r = ((QPaintEngine*)self)->QPaintEngine::coordinateOffset();
	return new QPoint(__r);
}
void* QPaintEngine_createPixmap_default(void* self, void* size) {
	auto&& __r = ((QPaintEngine*)self)->QPaintEngine::createPixmap(*(QSize*)size);
	return new QPixmap(__r);
}
void* QPaintEngine_createPixmapFromImage_default(void* self, void* image, int flags) {
	auto&& __r = ((QPaintEngine*)self)->QPaintEngine::createPixmapFromImage(*(QImage*)image, QFlags<Qt::ImageConversionFlag>(flags));
	return new QPixmap(__r);
}
void* QPaintEngine__Go_new(long long id) { return new QPaintEngine__Go(id); }
void QPaintEngine_delete(void* self) { delete (QPaintEngine*)self; }

}

#include <QtGui/qtextobject.h>
#include <QtGui/qtextdocument.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtGui/qtextformat.h>
#include <QtGui/qtextcursor.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qobject.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qvariant.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qtextframe.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

static std::unordered_map<long long, QMetaObject::Connection> __conns_QTextFrame;
static std::mutex __connsMu_QTextFrame;

extern "C" {

void* QTextFrame_base(void* self) { return (void*)static_cast<QTextObject*>((QTextFrame*)self); }
void QTextFrame_chain(void* self, void** out) {
	QTextFrame* __p = (QTextFrame*)self;
	QTextObject* __b0 = static_cast<QTextObject*>(__p); out[0] = (void*)__b0;
	QObject* __b1 = static_cast<QObject*>(__b0); out[1] = (void*)__b1;
}
void* QTextFrame_fromQObject(void* o) { return (void*)qobject_cast<QTextFrame*>((QObject*)o); }
void* QTextFrame_new(void* doc) {
	return new QTextFrame((QTextDocument*)doc);
}
char* QTextFrame_tr(const char* s, long long* __out_len) {
	auto&& __r = QTextFrame::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QTextFrame_setFrameFormat(void* self, void* format) {
	((QTextFrame*)self)->setFrameFormat(*(QTextFrameFormat*)format);
}
void* QTextFrame_frameFormat(void* self) {
	auto&& __r = ((QTextFrame*)self)->frameFormat();
	return new QTextFrameFormat(__r);
}
void* QTextFrame_firstCursorPosition(void* self) {
	auto&& __r = ((QTextFrame*)self)->firstCursorPosition();
	return new QTextCursor(__r);
}
void* QTextFrame_lastCursorPosition(void* self) {
	auto&& __r = ((QTextFrame*)self)->lastCursorPosition();
	return new QTextCursor(__r);
}
int QTextFrame_firstPosition(void* self) {
	auto&& __r = ((QTextFrame*)self)->firstPosition();
	return __r;
}
int QTextFrame_lastPosition(void* self) {
	auto&& __r = ((QTextFrame*)self)->lastPosition();
	return __r;
}
void* QTextFrame_layoutData(void* self) {
	auto&& __r = ((QTextFrame*)self)->layoutData();
	return (void*)(__r);
}
void QTextFrame_setLayoutData(void* self, void* data) {
	((QTextFrame*)self)->setLayoutData((QTextFrameLayoutData*)data);
}
void** QTextFrame_childFrames(void* self, long long* __out_len) {
	auto&& __r = ((QTextFrame*)self)->childFrames();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
void* QTextFrame_parentFrame(void* self) {
	auto&& __r = ((QTextFrame*)self)->parentFrame();
	return (void*)(__r);
}
void* QTextFrame_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QTextFrame::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QTextFrame_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QTextFrame::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
char* QTextFrame_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QTextFrame::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QTextFrame_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QTextFrame::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QTextFrame_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QTextFrame::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QTextFrame_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QTextFrame::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void QTextFrame_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTextFrame*)self, &QTextFrame::destroyed, (QTextFrame*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTextFrame);
	__conns_QTextFrame[__id] = __c;
}
void QTextFrame_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTextFrame*)self, &QTextFrame::destroyed, (QTextFrame*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTextFrame);
	__conns_QTextFrame[__id] = __c;
}
void QTextFrame_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTextFrame*)self, &QTextFrame::objectNameChanged, (QTextFrame*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTextFrame);
	__conns_QTextFrame[__id] = __c;
}
void QTextFrame__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QTextFrame);
		auto __it = __conns_QTextFrame.find(id);
		if (__it == __conns_QTextFrame.end()) return;
		__c = __it->second;
		__conns_QTextFrame.erase(__it);
	}
	QObject::disconnect(__c);
}
void QTextFrame_delete(void* self) { delete (QTextFrame*)self; }

}

#include <QtWidgets/qslider.h>
#include <QtWidgets/qabstractslider.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qsize.h>
#include <QtCore/qcoreevent.h>
#include <QtGui/qevent.h>
#include <QtWidgets/qstyleoption.h>
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
#include "gen_qslider.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QSlider__Go : public QSlider {
public:
	long long __goId;
	unsigned char __goOn[50] = {};
	template<typename... A> QSlider__Go(long long id, A&&... a) : QSlider(std::forward<A>(a)...), __goId(id) {}
	using QSlider::setRepeatAction;
	using QSlider::repeatAction;
	using QSlider::updateMicroFocus;
	using QSlider::create;
	using QSlider::destroy;
	using QSlider::focusNextChild;
	using QSlider::focusPreviousChild;
	using QSlider::sender;
	using QSlider::senderSignalIndex;
	using QSlider::receivers;
	using QSlider::isSignalConnected;
	using QSlider::getDecodedMetricF;
	void cppgoBase_paintEvent(QPaintEvent* __a0) { QSlider::paintEvent(__a0); }
	void cppgoBase_mousePressEvent(QMouseEvent* __a0) { QSlider::mousePressEvent(__a0); }
	void cppgoBase_mouseReleaseEvent(QMouseEvent* __a0) { QSlider::mouseReleaseEvent(__a0); }
	void cppgoBase_mouseMoveEvent(QMouseEvent* __a0) { QSlider::mouseMoveEvent(__a0); }
	void cppgoBase_initStyleOption(QStyleOptionSlider* __a0) { QSlider::initStyleOption(__a0); }
	void cppgoBase_sliderChange(int __a0) { QSlider::sliderChange((QAbstractSlider::SliderChange)__a0); }
	void cppgoBase_keyPressEvent(QKeyEvent* __a0) { QSlider::keyPressEvent(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QSlider::timerEvent(__a0); }
	void cppgoBase_wheelEvent(QWheelEvent* __a0) { QSlider::wheelEvent(__a0); }
	void cppgoBase_changeEvent(QEvent* __a0) { QSlider::changeEvent(__a0); }
	void cppgoBase_mouseDoubleClickEvent(QMouseEvent* __a0) { QSlider::mouseDoubleClickEvent(__a0); }
	void cppgoBase_keyReleaseEvent(QKeyEvent* __a0) { QSlider::keyReleaseEvent(__a0); }
	void cppgoBase_focusInEvent(QFocusEvent* __a0) { QSlider::focusInEvent(__a0); }
	void cppgoBase_focusOutEvent(QFocusEvent* __a0) { QSlider::focusOutEvent(__a0); }
	void cppgoBase_enterEvent(QEnterEvent* __a0) { QSlider::enterEvent(__a0); }
	void cppgoBase_leaveEvent(QEvent* __a0) { QSlider::leaveEvent(__a0); }
	void cppgoBase_moveEvent(QMoveEvent* __a0) { QSlider::moveEvent(__a0); }
	void cppgoBase_resizeEvent(QResizeEvent* __a0) { QSlider::resizeEvent(__a0); }
	void cppgoBase_closeEvent(QCloseEvent* __a0) { QSlider::closeEvent(__a0); }
	void cppgoBase_contextMenuEvent(QContextMenuEvent* __a0) { QSlider::contextMenuEvent(__a0); }
	void cppgoBase_tabletEvent(QTabletEvent* __a0) { QSlider::tabletEvent(__a0); }
	void cppgoBase_actionEvent(QActionEvent* __a0) { QSlider::actionEvent(__a0); }
	void cppgoBase_dragEnterEvent(QDragEnterEvent* __a0) { QSlider::dragEnterEvent(__a0); }
	void cppgoBase_dragMoveEvent(QDragMoveEvent* __a0) { QSlider::dragMoveEvent(__a0); }
	void cppgoBase_dragLeaveEvent(QDragLeaveEvent* __a0) { QSlider::dragLeaveEvent(__a0); }
	void cppgoBase_dropEvent(QDropEvent* __a0) { QSlider::dropEvent(__a0); }
	void cppgoBase_showEvent(QShowEvent* __a0) { QSlider::showEvent(__a0); }
	void cppgoBase_hideEvent(QHideEvent* __a0) { QSlider::hideEvent(__a0); }
	int cppgoBase_metric(int __a0) { return QSlider::metric((QPaintDevice::PaintDeviceMetric)__a0); }
	void cppgoBase_initPainter(QPainter* __a0) { QSlider::initPainter(__a0); }
	QPaintDevice* cppgoBase_redirected(QPoint* __a0) { return QSlider::redirected(__a0); }
	QPainter* cppgoBase_sharedPainter() { return QSlider::sharedPainter(); }
	void cppgoBase_inputMethodEvent(QInputMethodEvent* __a0) { QSlider::inputMethodEvent(__a0); }
	bool cppgoBase_focusNextPrevChild(bool __a0) { return QSlider::focusNextPrevChild(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QSlider::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QSlider::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QSlider::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QSlider::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QSlider::metaObject();
	}
	QSize sizeHint() const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 729690261457684101LL, &__h); if (__h) return *(QSize*)__r; }
		return QSlider::sizeHint();
	}
	QSize minimumSizeHint() const override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 7049264599837532491LL, &__h); if (__h) return *(QSize*)__r; }
		return QSlider::minimumSizeHint();
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QSlider::event(event);
	}
	void paintEvent(QPaintEvent* ev) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5290442163167458465LL, (void*)ev, &__h); if (__h) return; }
		QSlider::paintEvent(ev);
	}
	void mousePressEvent(QMouseEvent* ev) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -332090898078543319LL, (void*)ev, &__h); if (__h) return; }
		QSlider::mousePressEvent(ev);
	}
	void mouseReleaseEvent(QMouseEvent* ev) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4415920049297795213LL, (void*)ev, &__h); if (__h) return; }
		QSlider::mouseReleaseEvent(ev);
	}
	void mouseMoveEvent(QMouseEvent* ev) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4255494004099806581LL, (void*)ev, &__h); if (__h) return; }
		QSlider::mouseMoveEvent(ev);
	}
	void initStyleOption(QStyleOptionSlider* option) const override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -488693769771308531LL, (void*)option, &__h); if (__h) return; }
		QSlider::initStyleOption(option);
	}
	void sliderChange(QAbstractSlider::SliderChange change) override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 4658628703300230988LL, (int)(change), &__h); if (__h) return; }
		QSlider::sliderChange(change);
	}
	void keyPressEvent(QKeyEvent* ev) override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1023791977679793185LL, (void*)ev, &__h); if (__h) return; }
		QSlider::keyPressEvent(ev);
	}
	void timerEvent(QTimerEvent* arg1) override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)arg1, &__h); if (__h) return; }
		QSlider::timerEvent(arg1);
	}
	void wheelEvent(QWheelEvent* e) override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1919988340934511316LL, (void*)e, &__h); if (__h) return; }
		QSlider::wheelEvent(e);
	}
	void changeEvent(QEvent* e) override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4959376063826165943LL, (void*)e, &__h); if (__h) return; }
		QSlider::changeEvent(e);
	}
	int devType() const override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8659679893868524044LL, &__h); if (__h) return __r; }
		return QSlider::devType();
	}
	void setVisible(bool visible) override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 2110182719375245575LL, visible, &__h); if (__h) return; }
		QSlider::setVisible(visible);
	}
	int heightForWidth(int arg1) const override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 4311534060495737339LL, arg1, &__h); if (__h) return __r; }
		return QSlider::heightForWidth(arg1);
	}
	bool hasHeightForWidth() const override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8266762397995423915LL, &__h); if (__h) return __r; }
		return QSlider::hasHeightForWidth();
	}
	QPaintEngine* paintEngine() const override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 6398007404756346381LL, &__h); if (__h) return (QPaintEngine*)__r; }
		return QSlider::paintEngine();
	}
	void mouseDoubleClickEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4029073905697023887LL, (void*)event, &__h); if (__h) return; }
		QSlider::mouseDoubleClickEvent(event);
	}
	void keyReleaseEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 772598475950050299LL, (void*)event, &__h); if (__h) return; }
		QSlider::keyReleaseEvent(event);
	}
	void focusInEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 975043767033552552LL, (void*)event, &__h); if (__h) return; }
		QSlider::focusInEvent(event);
	}
	void focusOutEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -385596010576048993LL, (void*)event, &__h); if (__h) return; }
		QSlider::focusOutEvent(event);
	}
	void enterEvent(QEnterEvent* event) override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6316270346668237895LL, (void*)event, &__h); if (__h) return; }
		QSlider::enterEvent(event);
	}
	void leaveEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -551070009623711640LL, (void*)event, &__h); if (__h) return; }
		QSlider::leaveEvent(event);
	}
	void moveEvent(QMoveEvent* event) override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -9161949150489525396LL, (void*)event, &__h); if (__h) return; }
		QSlider::moveEvent(event);
	}
	void resizeEvent(QResizeEvent* event) override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4302917285655014661LL, (void*)event, &__h); if (__h) return; }
		QSlider::resizeEvent(event);
	}
	void closeEvent(QCloseEvent* event) override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1989937050019464783LL, (void*)event, &__h); if (__h) return; }
		QSlider::closeEvent(event);
	}
	void contextMenuEvent(QContextMenuEvent* event) override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6698043222227006015LL, (void*)event, &__h); if (__h) return; }
		QSlider::contextMenuEvent(event);
	}
	void tabletEvent(QTabletEvent* event) override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -212068866860701141LL, (void*)event, &__h); if (__h) return; }
		QSlider::tabletEvent(event);
	}
	void actionEvent(QActionEvent* event) override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1697444140088555581LL, (void*)event, &__h); if (__h) return; }
		QSlider::actionEvent(event);
	}
	void dragEnterEvent(QDragEnterEvent* event) override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 206122715919917589LL, (void*)event, &__h); if (__h) return; }
		QSlider::dragEnterEvent(event);
	}
	void dragMoveEvent(QDragMoveEvent* event) override {
		int __h = 0;
		if (__goOn[32] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1781545137352745978LL, (void*)event, &__h); if (__h) return; }
		QSlider::dragMoveEvent(event);
	}
	void dragLeaveEvent(QDragLeaveEvent* event) override {
		int __h = 0;
		if (__goOn[33] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -3139904158373725006LL, (void*)event, &__h); if (__h) return; }
		QSlider::dragLeaveEvent(event);
	}
	void dropEvent(QDropEvent* event) override {
		int __h = 0;
		if (__goOn[34] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 9107058682754591716LL, (void*)event, &__h); if (__h) return; }
		QSlider::dropEvent(event);
	}
	void showEvent(QShowEvent* event) override {
		int __h = 0;
		if (__goOn[35] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6825369174751916LL, (void*)event, &__h); if (__h) return; }
		QSlider::showEvent(event);
	}
	void hideEvent(QHideEvent* event) override {
		int __h = 0;
		if (__goOn[36] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -199718707876719129LL, (void*)event, &__h); if (__h) return; }
		QSlider::hideEvent(event);
	}
	bool nativeEvent(const QByteArray& eventType, void* message, qintptr* result) override {
		int __h = 0;
		if (__goOn[37] && cppgo_virt_int_voidp_voidp_longlongp) { int __r = ((int(*)(long long, long long, void*, void*, long long*, int*))cppgo_virt_int_voidp_voidp_longlongp)(__goId, -2039166586246443706LL, (void*)&eventType, (void*)message, (long long*)result, &__h); if (__h) return __r; }
		return QSlider::nativeEvent(eventType, message, result);
	}
	int metric(QPaintDevice::PaintDeviceMetric arg1) const override {
		int __h = 0;
		if (__goOn[38] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 427629649267311777LL, (int)(arg1), &__h); if (__h) return __r; }
		return QSlider::metric(arg1);
	}
	void initPainter(QPainter* painter) const override {
		int __h = 0;
		if (__goOn[39] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6477500296980050210LL, (void*)painter, &__h); if (__h) return; }
		QSlider::initPainter(painter);
	}
	QPaintDevice* redirected(QPoint* offset) const override {
		int __h = 0;
		if (__goOn[40] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 4127240633310655048LL, (void*)offset, &__h); if (__h) return (QPaintDevice*)__r; }
		return QSlider::redirected(offset);
	}
	QPainter* sharedPainter() const override {
		int __h = 0;
		if (__goOn[41] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 5792771328741774991LL, &__h); if (__h) return (QPainter*)__r; }
		return QSlider::sharedPainter();
	}
	void inputMethodEvent(QInputMethodEvent* arg1) override {
		int __h = 0;
		if (__goOn[42] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5438819661942843222LL, (void*)arg1, &__h); if (__h) return; }
		QSlider::inputMethodEvent(arg1);
	}
	QVariant inputMethodQuery(Qt::InputMethodQuery arg1) const override {
		int __h = 0;
		if (__goOn[43] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7345853097340709638LL, (int)(arg1), &__h); if (__h) return *(QVariant*)__r; }
		return QSlider::inputMethodQuery(arg1);
	}
	bool focusNextPrevChild(bool next) override {
		int __h = 0;
		if (__goOn[44] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 5898328517505784191LL, next, &__h); if (__h) return __r; }
		return QSlider::focusNextPrevChild(next);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[45] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QSlider::eventFilter(watched, event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[46] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QSlider::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[47] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QSlider::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[48] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QSlider::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[49] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QSlider::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QSlider;
static std::mutex __connsMu_QSlider;

extern "C" {

void* QSlider_base(void* self) { return (void*)static_cast<QAbstractSlider*>((QSlider*)self); }
void QSlider_chain(void* self, void** out) {
	QSlider* __p = (QSlider*)self;
	QAbstractSlider* __b0 = static_cast<QAbstractSlider*>(__p); out[0] = (void*)__b0;
	QWidget* __b1 = static_cast<QWidget*>(__b0); out[1] = (void*)__b1;
	QObject* __b2 = static_cast<QObject*>(__b1); out[2] = (void*)__b2;
}
void QSlider__set_override(void* self, int idx) { ((QSlider__Go*)self)->__goOn[idx] = 1; }
void* QSlider_fromQObject(void* o) { return (void*)qobject_cast<QSlider*>((QObject*)o); }
void* QSlider_new(long long __goid, void* parent) {
	return new QSlider__Go(__goid, (QWidget*)parent);
}
void* QSlider_new2(long long __goid) {
	return new QSlider__Go(__goid);
}
void* QSlider_new3(long long __goid, int orientation) {
	return new QSlider__Go(__goid, (Qt::Orientation)(orientation));
}
void* QSlider_new4(long long __goid, int orientation, void* parent) {
	return new QSlider__Go(__goid, (Qt::Orientation)(orientation), (QWidget*)parent);
}
char* QSlider_tr(const char* s, long long* __out_len) {
	auto&& __r = QSlider::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QSlider_setTickPosition(void* self, int position) {
	((QSlider*)self)->setTickPosition((QSlider::TickPosition)(position));
}
int QSlider_tickPosition(void* self) {
	auto&& __r = ((QSlider*)self)->tickPosition();
	return (int)(__r);
}
void QSlider_setTickInterval(void* self, int ti) {
	((QSlider*)self)->setTickInterval(ti);
}
int QSlider_tickInterval(void* self) {
	auto&& __r = ((QSlider*)self)->tickInterval();
	return __r;
}
void QSlider_setRepeatAction(void* self, int action) {
	((QSlider__Go*)self)->setRepeatAction((QAbstractSlider::SliderAction)(action));
}
int QSlider_repeatAction(void* self) {
	auto&& __r = ((QSlider__Go*)self)->repeatAction();
	return (int)(__r);
}
void QSlider_setTabOrder(void* arg1, void* arg2) {
	QSlider::setTabOrder((QWidget*)arg1, (QWidget*)arg2);
}
void* QSlider_mouseGrabber() {
	auto&& __r = QSlider::mouseGrabber();
	return (void*)(__r);
}
void* QSlider_keyboardGrabber() {
	auto&& __r = QSlider::keyboardGrabber();
	return (void*)(__r);
}
void* QSlider_find(unsigned long long arg1) {
	auto&& __r = QSlider::find(arg1);
	return (void*)(__r);
}
void* QSlider_createWindowContainer(void* window) {
	auto&& __r = QSlider::createWindowContainer((QWindow*)window);
	return (void*)(__r);
}
void QSlider_updateMicroFocus(void* self) {
	((QSlider__Go*)self)->updateMicroFocus();
}
void QSlider_create(void* self) {
	((QSlider__Go*)self)->create();
}
void QSlider_destroy(void* self) {
	((QSlider__Go*)self)->destroy();
}
int QSlider_focusNextChild(void* self) {
	auto&& __r = ((QSlider__Go*)self)->focusNextChild();
	return __r;
}
int QSlider_focusPreviousChild(void* self) {
	auto&& __r = ((QSlider__Go*)self)->focusPreviousChild();
	return __r;
}
void* QSlider_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QSlider::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QSlider_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QSlider::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QSlider_sender(void* self) {
	auto&& __r = ((QSlider__Go*)self)->sender();
	return (void*)(__r);
}
int QSlider_senderSignalIndex(void* self) {
	auto&& __r = ((QSlider__Go*)self)->senderSignalIndex();
	return __r;
}
int QSlider_receivers(void* self, const char* signal) {
	auto&& __r = ((QSlider__Go*)self)->receivers(signal);
	return __r;
}
int QSlider_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QSlider__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
double QSlider_devicePixelRatioFScale() {
	auto&& __r = QSlider::devicePixelRatioFScale();
	return __r;
}
int QSlider_encodeMetricF(int metric, double value) {
	auto&& __r = QSlider::encodeMetricF((QPaintDevice::PaintDeviceMetric)(metric), value);
	return __r;
}
double QSlider_getDecodedMetricF(void* self, int metricA, int metricB) {
	auto&& __r = ((QSlider__Go*)self)->getDecodedMetricF((QPaintDevice::PaintDeviceMetric)(metricA), (QPaintDevice::PaintDeviceMetric)(metricB));
	return __r;
}
char* QSlider_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QSlider::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QSlider_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QSlider::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QSlider_setRepeatAction2(void* self, int action, int thresholdTime) {
	((QSlider__Go*)self)->setRepeatAction((QAbstractSlider::SliderAction)(action), thresholdTime);
}
void QSlider_setRepeatAction3(void* self, int action, int thresholdTime, int repeatTime) {
	((QSlider__Go*)self)->setRepeatAction((QAbstractSlider::SliderAction)(action), thresholdTime, repeatTime);
}
void* QSlider_createWindowContainer2(void* window, void* parent) {
	auto&& __r = QSlider::createWindowContainer((QWindow*)window, (QWidget*)parent);
	return (void*)(__r);
}
void* QSlider_createWindowContainer3(void* window, void* parent, int flags) {
	auto&& __r = QSlider::createWindowContainer((QWindow*)window, (QWidget*)parent, QFlags<Qt::WindowType>(flags));
	return (void*)(__r);
}
void QSlider_updateMicroFocus2(void* self, int query) {
	((QSlider__Go*)self)->updateMicroFocus((Qt::InputMethodQuery)(query));
}
void QSlider_create2(void* self, unsigned long long arg1) {
	((QSlider__Go*)self)->create(arg1);
}
void QSlider_create3(void* self, unsigned long long arg1, int initializeWindow) {
	((QSlider__Go*)self)->create(arg1, initializeWindow);
}
void QSlider_create4(void* self, unsigned long long arg1, int initializeWindow, int destroyOldWindow) {
	((QSlider__Go*)self)->create(arg1, initializeWindow, destroyOldWindow);
}
void QSlider_destroy2(void* self, int destroyWindow) {
	((QSlider__Go*)self)->destroy(destroyWindow);
}
void QSlider_destroy3(void* self, int destroyWindow, int destroySubWindows) {
	((QSlider__Go*)self)->destroy(destroyWindow, destroySubWindows);
}
void* QSlider_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QSlider::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QSlider_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QSlider::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QSlider_metaObject_default(void* self) {
	auto&& __r = ((QSlider*)self)->QSlider::metaObject();
	return (void*)(__r);
}
void* QSlider_sizeHint_default(void* self) {
	auto&& __r = ((QSlider*)self)->QSlider::sizeHint();
	return new QSize(__r);
}
void* QSlider_minimumSizeHint_default(void* self) {
	auto&& __r = ((QSlider*)self)->QSlider::minimumSizeHint();
	return new QSize(__r);
}
int QSlider_event_default(void* self, void* event) {
	auto&& __r = ((QSlider*)self)->QSlider::event((QEvent*)event);
	return __r;
}
void QSlider_paintEvent_default(void* self, void* ev) {
	((QSlider__Go*)self)->cppgoBase_paintEvent((QPaintEvent*)ev);
}
void QSlider_mousePressEvent_default(void* self, void* ev) {
	((QSlider__Go*)self)->cppgoBase_mousePressEvent((QMouseEvent*)ev);
}
void QSlider_mouseReleaseEvent_default(void* self, void* ev) {
	((QSlider__Go*)self)->cppgoBase_mouseReleaseEvent((QMouseEvent*)ev);
}
void QSlider_mouseMoveEvent_default(void* self, void* ev) {
	((QSlider__Go*)self)->cppgoBase_mouseMoveEvent((QMouseEvent*)ev);
}
void QSlider_initStyleOption_default(void* self, void* option) {
	((QSlider__Go*)self)->cppgoBase_initStyleOption((QStyleOptionSlider*)option);
}
void QSlider_sliderChange_default(void* self, int change) {
	((QSlider__Go*)self)->cppgoBase_sliderChange(change);
}
void QSlider_keyPressEvent_default(void* self, void* ev) {
	((QSlider__Go*)self)->cppgoBase_keyPressEvent((QKeyEvent*)ev);
}
void QSlider_timerEvent_default(void* self, void* arg1) {
	((QSlider__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)arg1);
}
void QSlider_wheelEvent_default(void* self, void* e) {
	((QSlider__Go*)self)->cppgoBase_wheelEvent((QWheelEvent*)e);
}
void QSlider_changeEvent_default(void* self, void* e) {
	((QSlider__Go*)self)->cppgoBase_changeEvent((QEvent*)e);
}
int QSlider_devType_default(void* self) {
	auto&& __r = ((QSlider*)self)->QSlider::devType();
	return __r;
}
void QSlider_setVisible_default(void* self, int visible) {
	((QSlider*)self)->QSlider::setVisible(visible);
}
int QSlider_heightForWidth_default(void* self, int arg1) {
	auto&& __r = ((QSlider*)self)->QSlider::heightForWidth(arg1);
	return __r;
}
int QSlider_hasHeightForWidth_default(void* self) {
	auto&& __r = ((QSlider*)self)->QSlider::hasHeightForWidth();
	return __r;
}
void* QSlider_paintEngine_default(void* self) {
	auto&& __r = ((QSlider*)self)->QSlider::paintEngine();
	return (void*)(__r);
}
void QSlider_mouseDoubleClickEvent_default(void* self, void* event) {
	((QSlider__Go*)self)->cppgoBase_mouseDoubleClickEvent((QMouseEvent*)event);
}
void QSlider_keyReleaseEvent_default(void* self, void* event) {
	((QSlider__Go*)self)->cppgoBase_keyReleaseEvent((QKeyEvent*)event);
}
void QSlider_focusInEvent_default(void* self, void* event) {
	((QSlider__Go*)self)->cppgoBase_focusInEvent((QFocusEvent*)event);
}
void QSlider_focusOutEvent_default(void* self, void* event) {
	((QSlider__Go*)self)->cppgoBase_focusOutEvent((QFocusEvent*)event);
}
void QSlider_enterEvent_default(void* self, void* event) {
	((QSlider__Go*)self)->cppgoBase_enterEvent((QEnterEvent*)event);
}
void QSlider_leaveEvent_default(void* self, void* event) {
	((QSlider__Go*)self)->cppgoBase_leaveEvent((QEvent*)event);
}
void QSlider_moveEvent_default(void* self, void* event) {
	((QSlider__Go*)self)->cppgoBase_moveEvent((QMoveEvent*)event);
}
void QSlider_resizeEvent_default(void* self, void* event) {
	((QSlider__Go*)self)->cppgoBase_resizeEvent((QResizeEvent*)event);
}
void QSlider_closeEvent_default(void* self, void* event) {
	((QSlider__Go*)self)->cppgoBase_closeEvent((QCloseEvent*)event);
}
void QSlider_contextMenuEvent_default(void* self, void* event) {
	((QSlider__Go*)self)->cppgoBase_contextMenuEvent((QContextMenuEvent*)event);
}
void QSlider_tabletEvent_default(void* self, void* event) {
	((QSlider__Go*)self)->cppgoBase_tabletEvent((QTabletEvent*)event);
}
void QSlider_actionEvent_default(void* self, void* event) {
	((QSlider__Go*)self)->cppgoBase_actionEvent((QActionEvent*)event);
}
void QSlider_dragEnterEvent_default(void* self, void* event) {
	((QSlider__Go*)self)->cppgoBase_dragEnterEvent((QDragEnterEvent*)event);
}
void QSlider_dragMoveEvent_default(void* self, void* event) {
	((QSlider__Go*)self)->cppgoBase_dragMoveEvent((QDragMoveEvent*)event);
}
void QSlider_dragLeaveEvent_default(void* self, void* event) {
	((QSlider__Go*)self)->cppgoBase_dragLeaveEvent((QDragLeaveEvent*)event);
}
void QSlider_dropEvent_default(void* self, void* event) {
	((QSlider__Go*)self)->cppgoBase_dropEvent((QDropEvent*)event);
}
void QSlider_showEvent_default(void* self, void* event) {
	((QSlider__Go*)self)->cppgoBase_showEvent((QShowEvent*)event);
}
void QSlider_hideEvent_default(void* self, void* event) {
	((QSlider__Go*)self)->cppgoBase_hideEvent((QHideEvent*)event);
}
int QSlider_metric_default(void* self, int arg1) {
	auto&& __r = ((QSlider__Go*)self)->cppgoBase_metric(arg1);
	return __r;
}
void QSlider_initPainter_default(void* self, void* painter) {
	((QSlider__Go*)self)->cppgoBase_initPainter((QPainter*)painter);
}
void* QSlider_redirected_default(void* self, void* offset) {
	auto&& __r = ((QSlider__Go*)self)->cppgoBase_redirected((QPoint*)offset);
	return (void*)(__r);
}
void* QSlider_sharedPainter_default(void* self) {
	auto&& __r = ((QSlider__Go*)self)->cppgoBase_sharedPainter();
	return (void*)(__r);
}
void QSlider_inputMethodEvent_default(void* self, void* arg1) {
	((QSlider__Go*)self)->cppgoBase_inputMethodEvent((QInputMethodEvent*)arg1);
}
void* QSlider_inputMethodQuery_default(void* self, int arg1) {
	auto&& __r = ((QSlider*)self)->QSlider::inputMethodQuery((Qt::InputMethodQuery)(arg1));
	return new QVariant(__r);
}
int QSlider_focusNextPrevChild_default(void* self, int next) {
	auto&& __r = ((QSlider__Go*)self)->cppgoBase_focusNextPrevChild(next);
	return __r;
}
int QSlider_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QSlider*)self)->QSlider::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QSlider_childEvent_default(void* self, void* event) {
	((QSlider__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QSlider_customEvent_default(void* self, void* event) {
	((QSlider__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QSlider_connectNotify_default(void* self, void* signal) {
	((QSlider__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QSlider_disconnectNotify_default(void* self, void* signal) {
	((QSlider__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QSlider_connect_valueChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSlider*)self, &QSlider::valueChanged, (QSlider*)self, [__id](int value){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, value);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSlider);
	__conns_QSlider[__id] = __c;
}
void QSlider_connect_sliderPressed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSlider*)self, &QSlider::sliderPressed, (QSlider*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSlider);
	__conns_QSlider[__id] = __c;
}
void QSlider_connect_sliderMoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSlider*)self, &QSlider::sliderMoved, (QSlider*)self, [__id](int position){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, position);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSlider);
	__conns_QSlider[__id] = __c;
}
void QSlider_connect_sliderReleased(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSlider*)self, &QSlider::sliderReleased, (QSlider*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSlider);
	__conns_QSlider[__id] = __c;
}
void QSlider_connect_rangeChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSlider*)self, &QSlider::rangeChanged, (QSlider*)self, [__id](int min, int max){
		if (cppgo_sig_void_int_int) ((void(*)(long long, int, int))cppgo_sig_void_int_int)(__id, min, max);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSlider);
	__conns_QSlider[__id] = __c;
}
void QSlider_connect_actionTriggered(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSlider*)self, &QSlider::actionTriggered, (QSlider*)self, [__id](int action){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, action);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSlider);
	__conns_QSlider[__id] = __c;
}
void QSlider_connect_windowTitleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSlider*)self, &QSlider::windowTitleChanged, (QSlider*)self, [__id](const QString& title){
		QByteArray title_ba = title.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)title_ba.constData(), title_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSlider);
	__conns_QSlider[__id] = __c;
}
void QSlider_connect_windowIconChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSlider*)self, &QSlider::windowIconChanged, (QSlider*)self, [__id](const QIcon& icon){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&icon);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSlider);
	__conns_QSlider[__id] = __c;
}
void QSlider_connect_windowIconTextChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSlider*)self, &QSlider::windowIconTextChanged, (QSlider*)self, [__id](const QString& iconText){
		QByteArray iconText_ba = iconText.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)iconText_ba.constData(), iconText_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSlider);
	__conns_QSlider[__id] = __c;
}
void QSlider_connect_customContextMenuRequested(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSlider*)self, &QSlider::customContextMenuRequested, (QSlider*)self, [__id](const QPoint& pos){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&pos);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSlider);
	__conns_QSlider[__id] = __c;
}
void QSlider_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSlider*)self, &QSlider::destroyed, (QSlider*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSlider);
	__conns_QSlider[__id] = __c;
}
void QSlider_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSlider*)self, &QSlider::destroyed, (QSlider*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSlider);
	__conns_QSlider[__id] = __c;
}
void QSlider_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSlider*)self, &QSlider::objectNameChanged, (QSlider*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSlider);
	__conns_QSlider[__id] = __c;
}
void QSlider__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QSlider);
		auto __it = __conns_QSlider.find(id);
		if (__it == __conns_QSlider.end()) return;
		__c = __it->second;
		__conns_QSlider.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QSlider__Go_new(long long id) { return new QSlider__Go(id); }
void QSlider_delete(void* self) { delete (QSlider*)self; }

}

#include <QtWidgets/qstyleoption.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qnamespace.h>
#include <QtWidgets/qslider.h>
#include "gen_qstyleoptionslider.h"

extern "C" {

void* QStyleOptionSlider_base(void* self) { return (void*)static_cast<QStyleOptionComplex*>((QStyleOptionSlider*)self); }
void QStyleOptionSlider_chain(void* self, void** out) {
	QStyleOptionSlider* __p = (QStyleOptionSlider*)self;
	QStyleOptionComplex* __b0 = static_cast<QStyleOptionComplex*>(__p); out[0] = (void*)__b0;
	QStyleOption* __b1 = static_cast<QStyleOption*>(__b0); out[1] = (void*)__b1;
}
void* QStyleOptionSlider_new() {
	return new QStyleOptionSlider();
}
void* QStyleOptionSlider_new2(void* other) {
	return new QStyleOptionSlider(*(QStyleOptionSlider*)other);
}
void* QStyleOptionSlider_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QStyleOptionSlider*)self)->operator=(*(QStyleOptionSlider*)arg1);
	return (void*)&(__r);
}
int QStyleOptionSlider_orientation_get(void* self) {
	auto __r = ((QStyleOptionSlider*)self)->orientation;
	return (int)(__r);
}
void QStyleOptionSlider_orientation_set(void* self, int orientation) {
	((QStyleOptionSlider*)self)->orientation = (Qt::Orientation)(orientation);
}
int QStyleOptionSlider_minimum_get(void* self) {
	auto __r = ((QStyleOptionSlider*)self)->minimum;
	return __r;
}
void QStyleOptionSlider_minimum_set(void* self, int minimum) {
	((QStyleOptionSlider*)self)->minimum = minimum;
}
int QStyleOptionSlider_maximum_get(void* self) {
	auto __r = ((QStyleOptionSlider*)self)->maximum;
	return __r;
}
void QStyleOptionSlider_maximum_set(void* self, int maximum) {
	((QStyleOptionSlider*)self)->maximum = maximum;
}
int QStyleOptionSlider_tickPosition_get(void* self) {
	auto __r = ((QStyleOptionSlider*)self)->tickPosition;
	return (int)(__r);
}
void QStyleOptionSlider_tickPosition_set(void* self, int tickPosition) {
	((QStyleOptionSlider*)self)->tickPosition = (QSlider::TickPosition)(tickPosition);
}
int QStyleOptionSlider_tickInterval_get(void* self) {
	auto __r = ((QStyleOptionSlider*)self)->tickInterval;
	return __r;
}
void QStyleOptionSlider_tickInterval_set(void* self, int tickInterval) {
	((QStyleOptionSlider*)self)->tickInterval = tickInterval;
}
int QStyleOptionSlider_upsideDown_get(void* self) {
	auto __r = ((QStyleOptionSlider*)self)->upsideDown;
	return __r;
}
void QStyleOptionSlider_upsideDown_set(void* self, int upsideDown) {
	((QStyleOptionSlider*)self)->upsideDown = upsideDown;
}
int QStyleOptionSlider_sliderPosition_get(void* self) {
	auto __r = ((QStyleOptionSlider*)self)->sliderPosition;
	return __r;
}
void QStyleOptionSlider_sliderPosition_set(void* self, int sliderPosition) {
	((QStyleOptionSlider*)self)->sliderPosition = sliderPosition;
}
int QStyleOptionSlider_sliderValue_get(void* self) {
	auto __r = ((QStyleOptionSlider*)self)->sliderValue;
	return __r;
}
void QStyleOptionSlider_sliderValue_set(void* self, int sliderValue) {
	((QStyleOptionSlider*)self)->sliderValue = sliderValue;
}
int QStyleOptionSlider_singleStep_get(void* self) {
	auto __r = ((QStyleOptionSlider*)self)->singleStep;
	return __r;
}
void QStyleOptionSlider_singleStep_set(void* self, int singleStep) {
	((QStyleOptionSlider*)self)->singleStep = singleStep;
}
int QStyleOptionSlider_pageStep_get(void* self) {
	auto __r = ((QStyleOptionSlider*)self)->pageStep;
	return __r;
}
void QStyleOptionSlider_pageStep_set(void* self, int pageStep) {
	((QStyleOptionSlider*)self)->pageStep = pageStep;
}
double QStyleOptionSlider_notchTarget_get(void* self) {
	auto __r = ((QStyleOptionSlider*)self)->notchTarget;
	return __r;
}
void QStyleOptionSlider_notchTarget_set(void* self, double notchTarget) {
	((QStyleOptionSlider*)self)->notchTarget = notchTarget;
}
int QStyleOptionSlider_dialWrapping_get(void* self) {
	auto __r = ((QStyleOptionSlider*)self)->dialWrapping;
	return __r;
}
void QStyleOptionSlider_dialWrapping_set(void* self, int dialWrapping) {
	((QStyleOptionSlider*)self)->dialWrapping = dialWrapping;
}
int QStyleOptionSlider_keyboardModifiers_get(void* self) {
	auto __r = ((QStyleOptionSlider*)self)->keyboardModifiers;
	return (int)(__r);
}
void QStyleOptionSlider_keyboardModifiers_set(void* self, int keyboardModifiers) {
	((QStyleOptionSlider*)self)->keyboardModifiers = QFlags<Qt::KeyboardModifier>(keyboardModifiers);
}
void QStyleOptionSlider_delete(void* self) { delete (QStyleOptionSlider*)self; }

}

#include <QtWidgets/qboxlayout.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtWidgets/qlayoutitem.h>
#include <QtCore/qnamespace.h>
#include <QtWidgets/qlayout.h>
#include <QtCore/qsize.h>
#include <QtCore/qrect.h>
#include <QtCore/qmargins.h>
#include <QtWidgets/qsizepolicy.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qobject.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qvariant.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qhboxlayout.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QHBoxLayout__Go : public QHBoxLayout {
public:
	long long __goId;
	unsigned char __goOn[32] = {};
	template<typename... A> QHBoxLayout__Go(long long id, A&&... a) : QHBoxLayout(std::forward<A>(a)...), __goId(id) {}
	using QHBoxLayout::widgetEvent;
	using QHBoxLayout::addChildLayout;
	using QHBoxLayout::addChildWidget;
	using QHBoxLayout::adoptLayout;
	using QHBoxLayout::alignmentRect;
	using QHBoxLayout::sender;
	using QHBoxLayout::senderSignalIndex;
	using QHBoxLayout::receivers;
	using QHBoxLayout::isSignalConnected;
	void cppgoBase_childEvent(QChildEvent* __a0) { QHBoxLayout::childEvent(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QHBoxLayout::timerEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QHBoxLayout::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QHBoxLayout::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QHBoxLayout::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QHBoxLayout::metaObject();
	}
	void addItem(QLayoutItem* arg1) override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6884465960736393867LL, (void*)arg1, &__h); if (__h) return; }
		QHBoxLayout::addItem(arg1);
	}
	int spacing() const override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -4587105490415799682LL, &__h); if (__h) return __r; }
		return QHBoxLayout::spacing();
	}
	void setSpacing(int spacing) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 4039479028406881496LL, spacing, &__h); if (__h) return; }
		QHBoxLayout::setSpacing(spacing);
	}
	QSize sizeHint() const override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 729690261457684101LL, &__h); if (__h) return *(QSize*)__r; }
		return QHBoxLayout::sizeHint();
	}
	QSize minimumSize() const override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, -7246748916541192690LL, &__h); if (__h) return *(QSize*)__r; }
		return QHBoxLayout::minimumSize();
	}
	QSize maximumSize() const override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 4158099606088301136LL, &__h); if (__h) return *(QSize*)__r; }
		return QHBoxLayout::maximumSize();
	}
	bool hasHeightForWidth() const override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8266762397995423915LL, &__h); if (__h) return __r; }
		return QHBoxLayout::hasHeightForWidth();
	}
	int heightForWidth(int arg1) const override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 4311534060495737339LL, arg1, &__h); if (__h) return __r; }
		return QHBoxLayout::heightForWidth(arg1);
	}
	int minimumHeightForWidth(int arg1) const override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 5037837229528203129LL, arg1, &__h); if (__h) return __r; }
		return QHBoxLayout::minimumHeightForWidth(arg1);
	}
	QFlags<Qt::Orientation> expandingDirections() const override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, 415665469015095569LL, &__h); if (__h) return QFlags<Qt::Orientation>(__r); }
		return QHBoxLayout::expandingDirections();
	}
	void invalidate() override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 6991341317990632368LL, &__h); if (__h) return; }
		QHBoxLayout::invalidate();
	}
	QLayoutItem* itemAt(int arg1) const override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7508512214013391107LL, arg1, &__h); if (__h) return (QLayoutItem*)__r; }
		return QHBoxLayout::itemAt(arg1);
	}
	QLayoutItem* takeAt(int arg1) override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7549889108533486643LL, arg1, &__h); if (__h) return (QLayoutItem*)__r; }
		return QHBoxLayout::takeAt(arg1);
	}
	int count() const override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -5627843058667511180LL, &__h); if (__h) return __r; }
		return QHBoxLayout::count();
	}
	void setGeometry(const QRect& arg1) override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5684820848726969227LL, (void*)&arg1, &__h); if (__h) return; }
		QHBoxLayout::setGeometry(arg1);
	}
	QRect geometry() const override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1050490530495513297LL, &__h); if (__h) return *(QRect*)__r; }
		return QHBoxLayout::geometry();
	}
	int indexOf(const QWidget* arg1) const override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, -7258473138030329570LL, (void*)arg1, &__h); if (__h) return __r; }
		return QHBoxLayout::indexOf(arg1);
	}
	int indexOf(const QLayoutItem* arg1) const override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, -3276435765994963516LL, (void*)arg1, &__h); if (__h) return __r; }
		return QHBoxLayout::indexOf(arg1);
	}
	bool isEmpty() const override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -626822430251098498LL, &__h); if (__h) return __r; }
		return QHBoxLayout::isEmpty();
	}
	QFlags<QSizePolicy::ControlType> controlTypes() const override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, 3278947275714896887LL, &__h); if (__h) return QFlags<QSizePolicy::ControlType>(__r); }
		return QHBoxLayout::controlTypes();
	}
	QLayoutItem* replaceWidget(QWidget* from, QWidget* to, QFlags<Qt::FindChildOption> options) override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_voidp_voidp_voidp_int) { void* __r = ((void*(*)(long long, long long, void*, void*, int, int*))cppgo_virt_voidp_voidp_voidp_int)(__goId, -712862742159708061LL, (void*)from, (void*)to, (int)(options), &__h); if (__h) return (QLayoutItem*)__r; }
		return QHBoxLayout::replaceWidget(from, to, options);
	}
	QLayout* layout() override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, -7183390618627992737LL, &__h); if (__h) return (QLayout*)__r; }
		return QHBoxLayout::layout();
	}
	void childEvent(QChildEvent* e) override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)e, &__h); if (__h) return; }
		QHBoxLayout::childEvent(e);
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QHBoxLayout::event(event);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QHBoxLayout::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QHBoxLayout::timerEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QHBoxLayout::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QHBoxLayout::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QHBoxLayout::disconnectNotify(signal);
	}
	QWidget* widget() const override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, -4575937485689502171LL, &__h); if (__h) return (QWidget*)__r; }
		return QHBoxLayout::widget();
	}
	QSpacerItem* spacerItem() override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 2272240669359776278LL, &__h); if (__h) return (QSpacerItem*)__r; }
		return QHBoxLayout::spacerItem();
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QHBoxLayout;
static std::mutex __connsMu_QHBoxLayout;

extern "C" {

void* QHBoxLayout_base(void* self) { return (void*)static_cast<QBoxLayout*>((QHBoxLayout*)self); }
void QHBoxLayout_chain(void* self, void** out) {
	QHBoxLayout* __p = (QHBoxLayout*)self;
	QBoxLayout* __b0 = static_cast<QBoxLayout*>(__p); out[0] = (void*)__b0;
	QLayout* __b1 = static_cast<QLayout*>(__b0); out[1] = (void*)__b1;
	QObject* __b2 = static_cast<QObject*>(__b1); out[2] = (void*)__b2;
}
void QHBoxLayout__set_override(void* self, int idx) { ((QHBoxLayout__Go*)self)->__goOn[idx] = 1; }
void* QHBoxLayout_fromQObject(void* o) { return (void*)qobject_cast<QHBoxLayout*>((QObject*)o); }
void* QHBoxLayout_new(long long __goid, void* parent) {
	return new QHBoxLayout__Go(__goid, (QWidget*)parent);
}
void* QHBoxLayout_new2(long long __goid) {
	return new QHBoxLayout__Go(__goid);
}
char* QHBoxLayout_tr(const char* s, long long* __out_len) {
	auto&& __r = QHBoxLayout::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QHBoxLayout_closestAcceptableSize(void* w, void* s) {
	auto&& __r = QHBoxLayout::closestAcceptableSize((QWidget*)w, *(QSize*)s);
	return new QSize(__r);
}
void QHBoxLayout_widgetEvent(void* self, void* arg1) {
	((QHBoxLayout__Go*)self)->widgetEvent((QEvent*)arg1);
}
void QHBoxLayout_addChildLayout(void* self, void* l) {
	((QHBoxLayout__Go*)self)->addChildLayout((QLayout*)l);
}
void QHBoxLayout_addChildWidget(void* self, void* w) {
	((QHBoxLayout__Go*)self)->addChildWidget((QWidget*)w);
}
int QHBoxLayout_adoptLayout(void* self, void* layout) {
	auto&& __r = ((QHBoxLayout__Go*)self)->adoptLayout((QLayout*)layout);
	return __r;
}
void* QHBoxLayout_alignmentRect(void* self, void* arg1) {
	auto&& __r = ((QHBoxLayout__Go*)self)->alignmentRect(*(QRect*)arg1);
	return new QRect(__r);
}
void* QHBoxLayout_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QHBoxLayout::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QHBoxLayout_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QHBoxLayout::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QHBoxLayout_sender(void* self) {
	auto&& __r = ((QHBoxLayout__Go*)self)->sender();
	return (void*)(__r);
}
int QHBoxLayout_senderSignalIndex(void* self) {
	auto&& __r = ((QHBoxLayout__Go*)self)->senderSignalIndex();
	return __r;
}
int QHBoxLayout_receivers(void* self, const char* signal) {
	auto&& __r = ((QHBoxLayout__Go*)self)->receivers(signal);
	return __r;
}
int QHBoxLayout_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QHBoxLayout__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QHBoxLayout_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QHBoxLayout::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QHBoxLayout_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QHBoxLayout::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QHBoxLayout_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QHBoxLayout::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QHBoxLayout_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QHBoxLayout::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QHBoxLayout_metaObject_default(void* self) {
	auto&& __r = ((QHBoxLayout*)self)->QHBoxLayout::metaObject();
	return (void*)(__r);
}
void QHBoxLayout_addItem_default(void* self, void* arg1) {
	((QHBoxLayout*)self)->QHBoxLayout::addItem((QLayoutItem*)arg1);
}
int QHBoxLayout_spacing_default(void* self) {
	auto&& __r = ((QHBoxLayout*)self)->QHBoxLayout::spacing();
	return __r;
}
void QHBoxLayout_setSpacing_default(void* self, int spacing) {
	((QHBoxLayout*)self)->QHBoxLayout::setSpacing(spacing);
}
void* QHBoxLayout_sizeHint_default(void* self) {
	auto&& __r = ((QHBoxLayout*)self)->QHBoxLayout::sizeHint();
	return new QSize(__r);
}
void* QHBoxLayout_minimumSize_default(void* self) {
	auto&& __r = ((QHBoxLayout*)self)->QHBoxLayout::minimumSize();
	return new QSize(__r);
}
void* QHBoxLayout_maximumSize_default(void* self) {
	auto&& __r = ((QHBoxLayout*)self)->QHBoxLayout::maximumSize();
	return new QSize(__r);
}
int QHBoxLayout_hasHeightForWidth_default(void* self) {
	auto&& __r = ((QHBoxLayout*)self)->QHBoxLayout::hasHeightForWidth();
	return __r;
}
int QHBoxLayout_heightForWidth_default(void* self, int arg1) {
	auto&& __r = ((QHBoxLayout*)self)->QHBoxLayout::heightForWidth(arg1);
	return __r;
}
int QHBoxLayout_minimumHeightForWidth_default(void* self, int arg1) {
	auto&& __r = ((QHBoxLayout*)self)->QHBoxLayout::minimumHeightForWidth(arg1);
	return __r;
}
int QHBoxLayout_expandingDirections_default(void* self) {
	auto&& __r = ((QHBoxLayout*)self)->QHBoxLayout::expandingDirections();
	return (int)(__r);
}
void QHBoxLayout_invalidate_default(void* self) {
	((QHBoxLayout*)self)->QHBoxLayout::invalidate();
}
void* QHBoxLayout_itemAt_default(void* self, int arg1) {
	auto&& __r = ((QHBoxLayout*)self)->QHBoxLayout::itemAt(arg1);
	return (void*)(__r);
}
void* QHBoxLayout_takeAt_default(void* self, int arg1) {
	auto&& __r = ((QHBoxLayout*)self)->QHBoxLayout::takeAt(arg1);
	return (void*)(__r);
}
int QHBoxLayout_count_default(void* self) {
	auto&& __r = ((QHBoxLayout*)self)->QHBoxLayout::count();
	return __r;
}
void QHBoxLayout_setGeometry_default(void* self, void* arg1) {
	((QHBoxLayout*)self)->QHBoxLayout::setGeometry(*(QRect*)arg1);
}
void* QHBoxLayout_geometry_default(void* self) {
	auto&& __r = ((QHBoxLayout*)self)->QHBoxLayout::geometry();
	return new QRect(__r);
}
int QHBoxLayout_indexOf_default(void* self, void* arg1) {
	auto&& __r = ((QHBoxLayout*)self)->QHBoxLayout::indexOf((QWidget*)arg1);
	return __r;
}
int QHBoxLayout_isEmpty_default(void* self) {
	auto&& __r = ((QHBoxLayout*)self)->QHBoxLayout::isEmpty();
	return __r;
}
int QHBoxLayout_controlTypes_default(void* self) {
	auto&& __r = ((QHBoxLayout*)self)->QHBoxLayout::controlTypes();
	return (int)(__r);
}
void* QHBoxLayout_replaceWidget_default(void* self, void* from, void* to, int options) {
	auto&& __r = ((QHBoxLayout*)self)->QHBoxLayout::replaceWidget((QWidget*)from, (QWidget*)to, QFlags<Qt::FindChildOption>(options));
	return (void*)(__r);
}
void* QHBoxLayout_layout_default(void* self) {
	auto&& __r = ((QHBoxLayout*)self)->QHBoxLayout::layout();
	return (void*)(__r);
}
void QHBoxLayout_childEvent_default(void* self, void* e) {
	((QHBoxLayout__Go*)self)->cppgoBase_childEvent((QChildEvent*)e);
}
int QHBoxLayout_event_default(void* self, void* event) {
	auto&& __r = ((QHBoxLayout*)self)->QHBoxLayout::event((QEvent*)event);
	return __r;
}
int QHBoxLayout_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QHBoxLayout*)self)->QHBoxLayout::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QHBoxLayout_timerEvent_default(void* self, void* event) {
	((QHBoxLayout__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QHBoxLayout_customEvent_default(void* self, void* event) {
	((QHBoxLayout__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QHBoxLayout_connectNotify_default(void* self, void* signal) {
	((QHBoxLayout__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QHBoxLayout_disconnectNotify_default(void* self, void* signal) {
	((QHBoxLayout__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void* QHBoxLayout_widget_default(void* self) {
	auto&& __r = ((QHBoxLayout*)self)->QHBoxLayout::widget();
	return (void*)(__r);
}
void* QHBoxLayout_spacerItem_default(void* self) {
	auto&& __r = ((QHBoxLayout*)self)->QHBoxLayout::spacerItem();
	return (void*)(__r);
}
void QHBoxLayout_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QHBoxLayout*)self, &QHBoxLayout::destroyed, (QHBoxLayout*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QHBoxLayout);
	__conns_QHBoxLayout[__id] = __c;
}
void QHBoxLayout_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QHBoxLayout*)self, &QHBoxLayout::destroyed, (QHBoxLayout*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QHBoxLayout);
	__conns_QHBoxLayout[__id] = __c;
}
void QHBoxLayout_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QHBoxLayout*)self, &QHBoxLayout::objectNameChanged, (QHBoxLayout*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QHBoxLayout);
	__conns_QHBoxLayout[__id] = __c;
}
void QHBoxLayout__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QHBoxLayout);
		auto __it = __conns_QHBoxLayout.find(id);
		if (__it == __conns_QHBoxLayout.end()) return;
		__c = __it->second;
		__conns_QHBoxLayout.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QHBoxLayout__Go_new(long long id) { return new QHBoxLayout__Go(id); }
void QHBoxLayout_delete(void* self) { delete (QHBoxLayout*)self; }

}

#include <QtWidgets/qfileiconprovider.h>
#include <QtGui/qabstractfileiconprovider.h>
#include <QtGui/qicon.h>
#include <QtCore/qfileinfo.h>
#include <QtCore/qstring.h>
#include "gen_qfileiconprovider.h"
#include "gen_purego_cb.h"

class QFileIconProvider__Go : public QFileIconProvider {
public:
	long long __goId;
	unsigned char __goOn[5] = {};
	template<typename... A> QFileIconProvider__Go(long long id, A&&... a) : QFileIconProvider(std::forward<A>(a)...), __goId(id) {}
	QIcon icon(QAbstractFileIconProvider::IconType type_) const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, -2594270449535773968LL, (int)(type_), &__h); if (__h) return *(QIcon*)__r; }
		return QFileIconProvider::icon(type_);
	}
	QIcon icon(const QFileInfo& info) const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, -909105390709809242LL, (void*)&info, &__h); if (__h) return *(QIcon*)__r; }
		return QFileIconProvider::icon(info);
	}
	QString type(const QFileInfo& arg1) const override {
		int __h = 0;
		if (__goOn[2] && cppgo_virtp_charp_voidp) {
			long long __len = 0;
			char* __s = ((char*(*)(long long, long long, void*, long long*, int*))cppgo_virtp_charp_voidp)(__goId, -6371403872082420691LL, (void*)&arg1, &__len, &__h);
			if (__h) { QString __r = QString::fromUtf8(__s, __len); free(__s); return __r; }
		}
		return QFileIconProvider::type(arg1);
	}
	void setOptions(QFlags<QAbstractFileIconProvider::Option> arg1) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 5783776084322042327LL, (int)(arg1), &__h); if (__h) return; }
		QFileIconProvider::setOptions(arg1);
	}
	QFlags<QAbstractFileIconProvider::Option> options() const override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -3871419584579008699LL, &__h); if (__h) return QFlags<QAbstractFileIconProvider::Option>(__r); }
		return QFileIconProvider::options();
	}
};

extern "C" {

void* QFileIconProvider_base(void* self) { return (void*)static_cast<QAbstractFileIconProvider*>((QFileIconProvider*)self); }
void QFileIconProvider__set_override(void* self, int idx) { ((QFileIconProvider__Go*)self)->__goOn[idx] = 1; }
void* QFileIconProvider_new(long long __goid) {
	return new QFileIconProvider__Go(__goid);
}
void* QFileIconProvider_icon_default(void* self, int type_) {
	auto&& __r = ((QFileIconProvider*)self)->QFileIconProvider::icon((QAbstractFileIconProvider::IconType)(type_));
	return new QIcon(__r);
}
char* QFileIconProvider_type_default(void* self, void* arg1, long long* __out_len) {
	auto&& __r = ((QFileIconProvider*)self)->QFileIconProvider::type(*(QFileInfo*)arg1);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QFileIconProvider_setOptions_default(void* self, int arg1) {
	((QFileIconProvider*)self)->QFileIconProvider::setOptions(QFlags<QAbstractFileIconProvider::Option>(arg1));
}
int QFileIconProvider_options_default(void* self) {
	auto&& __r = ((QFileIconProvider*)self)->QFileIconProvider::options();
	return (int)(__r);
}
void* QFileIconProvider__Go_new(long long id) { return new QFileIconProvider__Go(id); }
void QFileIconProvider_delete(void* self) { delete (QFileIconProvider*)self; }

}

#include <QtWidgets/qgraphicsitem.h>
#include <QtCore/qstring.h>
#include <QtGui/qfont.h>
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
#include "gen_qgraphicssimpletextitem.h"
#include "gen_purego_cb.h"

class QGraphicsSimpleTextItem__Go : public QGraphicsSimpleTextItem {
public:
	long long __goId;
	unsigned char __goOn[35] = {};
	template<typename... A> QGraphicsSimpleTextItem__Go(long long id, A&&... a) : QGraphicsSimpleTextItem(std::forward<A>(a)...), __goId(id) {}
	using QGraphicsSimpleTextItem::updateMicroFocus;
	using QGraphicsSimpleTextItem::addToIndex;
	using QGraphicsSimpleTextItem::removeFromIndex;
	using QGraphicsSimpleTextItem::prepareGeometryChange;
	bool cppgoBase_supportsExtension(int __a0) { return QGraphicsSimpleTextItem::supportsExtension((QGraphicsItem::Extension)__a0); }
	void cppgoBase_setExtension(int __a0, const QVariant& __a1) { QGraphicsSimpleTextItem::setExtension((QGraphicsItem::Extension)__a0, __a1); }
	QVariant cppgoBase_extension(const QVariant& __a0) { return QGraphicsSimpleTextItem::extension(__a0); }
	bool cppgoBase_sceneEventFilter(QGraphicsItem* __a0, QEvent* __a1) { return QGraphicsSimpleTextItem::sceneEventFilter(__a0, __a1); }
	bool cppgoBase_sceneEvent(QEvent* __a0) { return QGraphicsSimpleTextItem::sceneEvent(__a0); }
	void cppgoBase_contextMenuEvent(QGraphicsSceneContextMenuEvent* __a0) { QGraphicsSimpleTextItem::contextMenuEvent(__a0); }
	void cppgoBase_dragEnterEvent(QGraphicsSceneDragDropEvent* __a0) { QGraphicsSimpleTextItem::dragEnterEvent(__a0); }
	void cppgoBase_dragLeaveEvent(QGraphicsSceneDragDropEvent* __a0) { QGraphicsSimpleTextItem::dragLeaveEvent(__a0); }
	void cppgoBase_dragMoveEvent(QGraphicsSceneDragDropEvent* __a0) { QGraphicsSimpleTextItem::dragMoveEvent(__a0); }
	void cppgoBase_dropEvent(QGraphicsSceneDragDropEvent* __a0) { QGraphicsSimpleTextItem::dropEvent(__a0); }
	void cppgoBase_focusInEvent(QFocusEvent* __a0) { QGraphicsSimpleTextItem::focusInEvent(__a0); }
	void cppgoBase_focusOutEvent(QFocusEvent* __a0) { QGraphicsSimpleTextItem::focusOutEvent(__a0); }
	void cppgoBase_hoverEnterEvent(QGraphicsSceneHoverEvent* __a0) { QGraphicsSimpleTextItem::hoverEnterEvent(__a0); }
	void cppgoBase_hoverMoveEvent(QGraphicsSceneHoverEvent* __a0) { QGraphicsSimpleTextItem::hoverMoveEvent(__a0); }
	void cppgoBase_hoverLeaveEvent(QGraphicsSceneHoverEvent* __a0) { QGraphicsSimpleTextItem::hoverLeaveEvent(__a0); }
	void cppgoBase_keyPressEvent(QKeyEvent* __a0) { QGraphicsSimpleTextItem::keyPressEvent(__a0); }
	void cppgoBase_keyReleaseEvent(QKeyEvent* __a0) { QGraphicsSimpleTextItem::keyReleaseEvent(__a0); }
	void cppgoBase_mousePressEvent(QGraphicsSceneMouseEvent* __a0) { QGraphicsSimpleTextItem::mousePressEvent(__a0); }
	void cppgoBase_mouseMoveEvent(QGraphicsSceneMouseEvent* __a0) { QGraphicsSimpleTextItem::mouseMoveEvent(__a0); }
	void cppgoBase_mouseReleaseEvent(QGraphicsSceneMouseEvent* __a0) { QGraphicsSimpleTextItem::mouseReleaseEvent(__a0); }
	void cppgoBase_mouseDoubleClickEvent(QGraphicsSceneMouseEvent* __a0) { QGraphicsSimpleTextItem::mouseDoubleClickEvent(__a0); }
	void cppgoBase_wheelEvent(QGraphicsSceneWheelEvent* __a0) { QGraphicsSimpleTextItem::wheelEvent(__a0); }
	void cppgoBase_inputMethodEvent(QInputMethodEvent* __a0) { QGraphicsSimpleTextItem::inputMethodEvent(__a0); }
	QVariant cppgoBase_inputMethodQuery(int __a0) { return QGraphicsSimpleTextItem::inputMethodQuery((Qt::InputMethodQuery)__a0); }
	QVariant cppgoBase_itemChange(int __a0, const QVariant& __a1) { return QGraphicsSimpleTextItem::itemChange((QGraphicsItem::GraphicsItemChange)__a0, __a1); }
	QRectF boundingRect() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 2510008782785038831LL, &__h); if (__h) return *(QRectF*)__r; }
		return QGraphicsSimpleTextItem::boundingRect();
	}
	QPainterPath shape() const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 7938019069247444294LL, &__h); if (__h) return *(QPainterPath*)__r; }
		return QGraphicsSimpleTextItem::shape();
	}
	bool contains(const QPointF& point) const override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 5645366511944179768LL, (void*)&point, &__h); if (__h) return __r; }
		return QGraphicsSimpleTextItem::contains(point);
	}
	void paint(QPainter* painter, const QStyleOptionGraphicsItem* option, QWidget* widget) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_voidp_voidp_voidp) { ((void(*)(long long, long long, void*, void*, void*, int*))cppgo_virt_void_voidp_voidp_voidp)(__goId, 6083584261053556819LL, (void*)painter, (void*)option, (void*)widget, &__h); if (__h) return; }
		QGraphicsSimpleTextItem::paint(painter, option, widget);
	}
	bool isObscuredBy(const QGraphicsItem* item) const override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, -7215030271914903671LL, (void*)item, &__h); if (__h) return __r; }
		return QGraphicsSimpleTextItem::isObscuredBy(item);
	}
	QPainterPath opaqueArea() const override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 399370970023989435LL, &__h); if (__h) return *(QPainterPath*)__r; }
		return QGraphicsSimpleTextItem::opaqueArea();
	}
	int type() const override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -6371403872082420691LL, &__h); if (__h) return __r; }
		return QGraphicsSimpleTextItem::type();
	}
	bool supportsExtension(QGraphicsItem::Extension extension) const override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, -5729166443651306418LL, (int)(extension), &__h); if (__h) return __r; }
		return QGraphicsSimpleTextItem::supportsExtension(extension);
	}
	void setExtension(QGraphicsItem::Extension extension, const QVariant& variant) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void_int_voidp) { ((void(*)(long long, long long, int, void*, int*))cppgo_virt_void_int_voidp)(__goId, -6730181576106171580LL, (int)(extension), (void*)&variant, &__h); if (__h) return; }
		QGraphicsSimpleTextItem::setExtension(extension, variant);
	}
	QVariant extension(const QVariant& variant) const override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 2270616358835251226LL, (void*)&variant, &__h); if (__h) return *(QVariant*)__r; }
		return QGraphicsSimpleTextItem::extension(variant);
	}
	void advance(int phase) override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, -4830836341480073475LL, phase, &__h); if (__h) return; }
		QGraphicsSimpleTextItem::advance(phase);
	}
	bool collidesWithItem(const QGraphicsItem* other, Qt::ItemSelectionMode mode) const override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_int_voidp_int) { int __r = ((int(*)(long long, long long, void*, int, int*))cppgo_virt_int_voidp_int)(__goId, 6469736184362593173LL, (void*)other, (int)(mode), &__h); if (__h) return __r; }
		return QGraphicsSimpleTextItem::collidesWithItem(other, mode);
	}
	bool collidesWithPath(const QPainterPath& path, Qt::ItemSelectionMode mode) const override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_int_voidp_int) { int __r = ((int(*)(long long, long long, void*, int, int*))cppgo_virt_int_voidp_int)(__goId, 6636575182718753853LL, (void*)&path, (int)(mode), &__h); if (__h) return __r; }
		return QGraphicsSimpleTextItem::collidesWithPath(path, mode);
	}
	bool sceneEventFilter(QGraphicsItem* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, -4584203218241215725LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QGraphicsSimpleTextItem::sceneEventFilter(watched, event);
	}
	bool sceneEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, -2846371507838587071LL, (void*)event, &__h); if (__h) return __r; }
		return QGraphicsSimpleTextItem::sceneEvent(event);
	}
	void contextMenuEvent(QGraphicsSceneContextMenuEvent* event) override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6698043222227006015LL, (void*)event, &__h); if (__h) return; }
		QGraphicsSimpleTextItem::contextMenuEvent(event);
	}
	void dragEnterEvent(QGraphicsSceneDragDropEvent* event) override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 206122715919917589LL, (void*)event, &__h); if (__h) return; }
		QGraphicsSimpleTextItem::dragEnterEvent(event);
	}
	void dragLeaveEvent(QGraphicsSceneDragDropEvent* event) override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -3139904158373725006LL, (void*)event, &__h); if (__h) return; }
		QGraphicsSimpleTextItem::dragLeaveEvent(event);
	}
	void dragMoveEvent(QGraphicsSceneDragDropEvent* event) override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1781545137352745978LL, (void*)event, &__h); if (__h) return; }
		QGraphicsSimpleTextItem::dragMoveEvent(event);
	}
	void dropEvent(QGraphicsSceneDragDropEvent* event) override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 9107058682754591716LL, (void*)event, &__h); if (__h) return; }
		QGraphicsSimpleTextItem::dropEvent(event);
	}
	void focusInEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 975043767033552552LL, (void*)event, &__h); if (__h) return; }
		QGraphicsSimpleTextItem::focusInEvent(event);
	}
	void focusOutEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -385596010576048993LL, (void*)event, &__h); if (__h) return; }
		QGraphicsSimpleTextItem::focusOutEvent(event);
	}
	void hoverEnterEvent(QGraphicsSceneHoverEvent* event) override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6932941943906681679LL, (void*)event, &__h); if (__h) return; }
		QGraphicsSimpleTextItem::hoverEnterEvent(event);
	}
	void hoverMoveEvent(QGraphicsSceneHoverEvent* event) override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -2745191522144316396LL, (void*)event, &__h); if (__h) return; }
		QGraphicsSimpleTextItem::hoverMoveEvent(event);
	}
	void hoverLeaveEvent(QGraphicsSceneHoverEvent* event) override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7944408482203603984LL, (void*)event, &__h); if (__h) return; }
		QGraphicsSimpleTextItem::hoverLeaveEvent(event);
	}
	void keyPressEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1023791977679793185LL, (void*)event, &__h); if (__h) return; }
		QGraphicsSimpleTextItem::keyPressEvent(event);
	}
	void keyReleaseEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 772598475950050299LL, (void*)event, &__h); if (__h) return; }
		QGraphicsSimpleTextItem::keyReleaseEvent(event);
	}
	void mousePressEvent(QGraphicsSceneMouseEvent* event) override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -332090898078543319LL, (void*)event, &__h); if (__h) return; }
		QGraphicsSimpleTextItem::mousePressEvent(event);
	}
	void mouseMoveEvent(QGraphicsSceneMouseEvent* event) override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4255494004099806581LL, (void*)event, &__h); if (__h) return; }
		QGraphicsSimpleTextItem::mouseMoveEvent(event);
	}
	void mouseReleaseEvent(QGraphicsSceneMouseEvent* event) override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4415920049297795213LL, (void*)event, &__h); if (__h) return; }
		QGraphicsSimpleTextItem::mouseReleaseEvent(event);
	}
	void mouseDoubleClickEvent(QGraphicsSceneMouseEvent* event) override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4029073905697023887LL, (void*)event, &__h); if (__h) return; }
		QGraphicsSimpleTextItem::mouseDoubleClickEvent(event);
	}
	void wheelEvent(QGraphicsSceneWheelEvent* event) override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1919988340934511316LL, (void*)event, &__h); if (__h) return; }
		QGraphicsSimpleTextItem::wheelEvent(event);
	}
	void inputMethodEvent(QInputMethodEvent* event) override {
		int __h = 0;
		if (__goOn[32] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5438819661942843222LL, (void*)event, &__h); if (__h) return; }
		QGraphicsSimpleTextItem::inputMethodEvent(event);
	}
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const override {
		int __h = 0;
		if (__goOn[33] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7345853097340709638LL, (int)(query), &__h); if (__h) return *(QVariant*)__r; }
		return QGraphicsSimpleTextItem::inputMethodQuery(query);
	}
	QVariant itemChange(QGraphicsItem::GraphicsItemChange change, const QVariant& value) override {
		int __h = 0;
		if (__goOn[34] && cppgo_virt_voidp_int_voidp) { void* __r = ((void*(*)(long long, long long, int, void*, int*))cppgo_virt_voidp_int_voidp)(__goId, -1888572168794878830LL, (int)(change), (void*)&value, &__h); if (__h) return *(QVariant*)__r; }
		return QGraphicsSimpleTextItem::itemChange(change, value);
	}
};

extern "C" {

void* QGraphicsSimpleTextItem_base(void* self) { return (void*)static_cast<QAbstractGraphicsShapeItem*>((QGraphicsSimpleTextItem*)self); }
void QGraphicsSimpleTextItem_chain(void* self, void** out) {
	QGraphicsSimpleTextItem* __p = (QGraphicsSimpleTextItem*)self;
	QAbstractGraphicsShapeItem* __b0 = static_cast<QAbstractGraphicsShapeItem*>(__p); out[0] = (void*)__b0;
	QGraphicsItem* __b1 = static_cast<QGraphicsItem*>(__b0); out[1] = (void*)__b1;
}
void QGraphicsSimpleTextItem__set_override(void* self, int idx) { ((QGraphicsSimpleTextItem__Go*)self)->__goOn[idx] = 1; }
void* QGraphicsSimpleTextItem_new(long long __goid) {
	return new QGraphicsSimpleTextItem__Go(__goid);
}
void* QGraphicsSimpleTextItem_new2(long long __goid, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	return new QGraphicsSimpleTextItem__Go(__goid, text_qstr);
}
void* QGraphicsSimpleTextItem_new3(long long __goid, void* parent) {
	return new QGraphicsSimpleTextItem__Go(__goid, (QGraphicsItem*)parent);
}
void* QGraphicsSimpleTextItem_new4(long long __goid, const char* text, long long text_len, void* parent) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	return new QGraphicsSimpleTextItem__Go(__goid, text_qstr, (QGraphicsItem*)parent);
}
void QGraphicsSimpleTextItem_setText(void* self, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	((QGraphicsSimpleTextItem*)self)->setText(text_qstr);
}
char* QGraphicsSimpleTextItem_text(void* self, long long* __out_len) {
	auto&& __r = ((QGraphicsSimpleTextItem*)self)->text();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QGraphicsSimpleTextItem_setFont(void* self, void* font) {
	((QGraphicsSimpleTextItem*)self)->setFont(*(QFont*)font);
}
void* QGraphicsSimpleTextItem_font(void* self) {
	auto&& __r = ((QGraphicsSimpleTextItem*)self)->font();
	return new QFont(__r);
}
void QGraphicsSimpleTextItem_updateMicroFocus(void* self) {
	((QGraphicsSimpleTextItem__Go*)self)->updateMicroFocus();
}
void QGraphicsSimpleTextItem_addToIndex(void* self) {
	((QGraphicsSimpleTextItem__Go*)self)->addToIndex();
}
void QGraphicsSimpleTextItem_removeFromIndex(void* self) {
	((QGraphicsSimpleTextItem__Go*)self)->removeFromIndex();
}
void QGraphicsSimpleTextItem_prepareGeometryChange(void* self) {
	((QGraphicsSimpleTextItem__Go*)self)->prepareGeometryChange();
}
void* QGraphicsSimpleTextItem_boundingRect_default(void* self) {
	auto&& __r = ((QGraphicsSimpleTextItem*)self)->QGraphicsSimpleTextItem::boundingRect();
	return new QRectF(__r);
}
void* QGraphicsSimpleTextItem_shape_default(void* self) {
	auto&& __r = ((QGraphicsSimpleTextItem*)self)->QGraphicsSimpleTextItem::shape();
	return new QPainterPath(__r);
}
int QGraphicsSimpleTextItem_contains_default(void* self, void* point) {
	auto&& __r = ((QGraphicsSimpleTextItem*)self)->QGraphicsSimpleTextItem::contains(*(QPointF*)point);
	return __r;
}
void QGraphicsSimpleTextItem_paint_default(void* self, void* painter, void* option, void* widget) {
	((QGraphicsSimpleTextItem*)self)->QGraphicsSimpleTextItem::paint((QPainter*)painter, (QStyleOptionGraphicsItem*)option, (QWidget*)widget);
}
int QGraphicsSimpleTextItem_isObscuredBy_default(void* self, void* item) {
	auto&& __r = ((QGraphicsSimpleTextItem*)self)->QGraphicsSimpleTextItem::isObscuredBy((QGraphicsItem*)item);
	return __r;
}
void* QGraphicsSimpleTextItem_opaqueArea_default(void* self) {
	auto&& __r = ((QGraphicsSimpleTextItem*)self)->QGraphicsSimpleTextItem::opaqueArea();
	return new QPainterPath(__r);
}
int QGraphicsSimpleTextItem_type_default(void* self) {
	auto&& __r = ((QGraphicsSimpleTextItem*)self)->QGraphicsSimpleTextItem::type();
	return __r;
}
int QGraphicsSimpleTextItem_supportsExtension_default(void* self, int extension) {
	auto&& __r = ((QGraphicsSimpleTextItem__Go*)self)->cppgoBase_supportsExtension(extension);
	return __r;
}
void QGraphicsSimpleTextItem_setExtension_default(void* self, int extension, void* variant) {
	((QGraphicsSimpleTextItem__Go*)self)->cppgoBase_setExtension(extension, *(QVariant*)variant);
}
void* QGraphicsSimpleTextItem_extension_default(void* self, void* variant) {
	auto&& __r = ((QGraphicsSimpleTextItem__Go*)self)->cppgoBase_extension(*(QVariant*)variant);
	return new QVariant(__r);
}
void QGraphicsSimpleTextItem_advance_default(void* self, int phase) {
	((QGraphicsSimpleTextItem*)self)->QGraphicsSimpleTextItem::advance(phase);
}
int QGraphicsSimpleTextItem_collidesWithItem_default(void* self, void* other, int mode) {
	auto&& __r = ((QGraphicsSimpleTextItem*)self)->QGraphicsSimpleTextItem::collidesWithItem((QGraphicsItem*)other, (Qt::ItemSelectionMode)(mode));
	return __r;
}
int QGraphicsSimpleTextItem_collidesWithPath_default(void* self, void* path, int mode) {
	auto&& __r = ((QGraphicsSimpleTextItem*)self)->QGraphicsSimpleTextItem::collidesWithPath(*(QPainterPath*)path, (Qt::ItemSelectionMode)(mode));
	return __r;
}
int QGraphicsSimpleTextItem_sceneEventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QGraphicsSimpleTextItem__Go*)self)->cppgoBase_sceneEventFilter((QGraphicsItem*)watched, (QEvent*)event);
	return __r;
}
int QGraphicsSimpleTextItem_sceneEvent_default(void* self, void* event) {
	auto&& __r = ((QGraphicsSimpleTextItem__Go*)self)->cppgoBase_sceneEvent((QEvent*)event);
	return __r;
}
void QGraphicsSimpleTextItem_contextMenuEvent_default(void* self, void* event) {
	((QGraphicsSimpleTextItem__Go*)self)->cppgoBase_contextMenuEvent((QGraphicsSceneContextMenuEvent*)event);
}
void QGraphicsSimpleTextItem_dragEnterEvent_default(void* self, void* event) {
	((QGraphicsSimpleTextItem__Go*)self)->cppgoBase_dragEnterEvent((QGraphicsSceneDragDropEvent*)event);
}
void QGraphicsSimpleTextItem_dragLeaveEvent_default(void* self, void* event) {
	((QGraphicsSimpleTextItem__Go*)self)->cppgoBase_dragLeaveEvent((QGraphicsSceneDragDropEvent*)event);
}
void QGraphicsSimpleTextItem_dragMoveEvent_default(void* self, void* event) {
	((QGraphicsSimpleTextItem__Go*)self)->cppgoBase_dragMoveEvent((QGraphicsSceneDragDropEvent*)event);
}
void QGraphicsSimpleTextItem_dropEvent_default(void* self, void* event) {
	((QGraphicsSimpleTextItem__Go*)self)->cppgoBase_dropEvent((QGraphicsSceneDragDropEvent*)event);
}
void QGraphicsSimpleTextItem_focusInEvent_default(void* self, void* event) {
	((QGraphicsSimpleTextItem__Go*)self)->cppgoBase_focusInEvent((QFocusEvent*)event);
}
void QGraphicsSimpleTextItem_focusOutEvent_default(void* self, void* event) {
	((QGraphicsSimpleTextItem__Go*)self)->cppgoBase_focusOutEvent((QFocusEvent*)event);
}
void QGraphicsSimpleTextItem_hoverEnterEvent_default(void* self, void* event) {
	((QGraphicsSimpleTextItem__Go*)self)->cppgoBase_hoverEnterEvent((QGraphicsSceneHoverEvent*)event);
}
void QGraphicsSimpleTextItem_hoverMoveEvent_default(void* self, void* event) {
	((QGraphicsSimpleTextItem__Go*)self)->cppgoBase_hoverMoveEvent((QGraphicsSceneHoverEvent*)event);
}
void QGraphicsSimpleTextItem_hoverLeaveEvent_default(void* self, void* event) {
	((QGraphicsSimpleTextItem__Go*)self)->cppgoBase_hoverLeaveEvent((QGraphicsSceneHoverEvent*)event);
}
void QGraphicsSimpleTextItem_keyPressEvent_default(void* self, void* event) {
	((QGraphicsSimpleTextItem__Go*)self)->cppgoBase_keyPressEvent((QKeyEvent*)event);
}
void QGraphicsSimpleTextItem_keyReleaseEvent_default(void* self, void* event) {
	((QGraphicsSimpleTextItem__Go*)self)->cppgoBase_keyReleaseEvent((QKeyEvent*)event);
}
void QGraphicsSimpleTextItem_mousePressEvent_default(void* self, void* event) {
	((QGraphicsSimpleTextItem__Go*)self)->cppgoBase_mousePressEvent((QGraphicsSceneMouseEvent*)event);
}
void QGraphicsSimpleTextItem_mouseMoveEvent_default(void* self, void* event) {
	((QGraphicsSimpleTextItem__Go*)self)->cppgoBase_mouseMoveEvent((QGraphicsSceneMouseEvent*)event);
}
void QGraphicsSimpleTextItem_mouseReleaseEvent_default(void* self, void* event) {
	((QGraphicsSimpleTextItem__Go*)self)->cppgoBase_mouseReleaseEvent((QGraphicsSceneMouseEvent*)event);
}
void QGraphicsSimpleTextItem_mouseDoubleClickEvent_default(void* self, void* event) {
	((QGraphicsSimpleTextItem__Go*)self)->cppgoBase_mouseDoubleClickEvent((QGraphicsSceneMouseEvent*)event);
}
void QGraphicsSimpleTextItem_wheelEvent_default(void* self, void* event) {
	((QGraphicsSimpleTextItem__Go*)self)->cppgoBase_wheelEvent((QGraphicsSceneWheelEvent*)event);
}
void QGraphicsSimpleTextItem_inputMethodEvent_default(void* self, void* event) {
	((QGraphicsSimpleTextItem__Go*)self)->cppgoBase_inputMethodEvent((QInputMethodEvent*)event);
}
void* QGraphicsSimpleTextItem_inputMethodQuery_default(void* self, int query) {
	auto&& __r = ((QGraphicsSimpleTextItem__Go*)self)->cppgoBase_inputMethodQuery(query);
	return new QVariant(__r);
}
void* QGraphicsSimpleTextItem_itemChange_default(void* self, int change, void* value) {
	auto&& __r = ((QGraphicsSimpleTextItem__Go*)self)->cppgoBase_itemChange(change, *(QVariant*)value);
	return new QVariant(__r);
}
void* QGraphicsSimpleTextItem__Go_new(long long id) { return new QGraphicsSimpleTextItem__Go(id); }
void QGraphicsSimpleTextItem_delete(void* self) { delete (QGraphicsSimpleTextItem*)self; }

}

#include <QtWidgets/qgraphicssceneevent.h>
#include <QtCore/qsize.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qcoreevent.h>
#include "gen_qgraphicssceneresizeevent.h"
#include "gen_purego_cb.h"

class QGraphicsSceneResizeEvent__Go : public QGraphicsSceneResizeEvent {
public:
	long long __goId;
	unsigned char __goOn[2] = {};
	template<typename... A> QGraphicsSceneResizeEvent__Go(long long id, A&&... a) : QGraphicsSceneResizeEvent(std::forward<A>(a)...), __goId(id) {}
	void setAccepted(bool accepted) override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 6367896784077321574LL, accepted, &__h); if (__h) return; }
		QGraphicsSceneResizeEvent::setAccepted(accepted);
	}
	QEvent* clone() const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, -7032813281048616478LL, &__h); if (__h) return (QEvent*)__r; }
		return QGraphicsSceneResizeEvent::clone();
	}
};

extern "C" {

void* QGraphicsSceneResizeEvent_base(void* self) { return (void*)static_cast<QGraphicsSceneEvent*>((QGraphicsSceneResizeEvent*)self); }
void QGraphicsSceneResizeEvent_chain(void* self, void** out) {
	QGraphicsSceneResizeEvent* __p = (QGraphicsSceneResizeEvent*)self;
	QGraphicsSceneEvent* __b0 = static_cast<QGraphicsSceneEvent*>(__p); out[0] = (void*)__b0;
	QEvent* __b1 = static_cast<QEvent*>(__b0); out[1] = (void*)__b1;
}
void QGraphicsSceneResizeEvent__set_override(void* self, int idx) { ((QGraphicsSceneResizeEvent__Go*)self)->__goOn[idx] = 1; }
void* QGraphicsSceneResizeEvent_new(long long __goid) {
	return new QGraphicsSceneResizeEvent__Go(__goid);
}
void* QGraphicsSceneResizeEvent_oldSize(void* self) {
	auto&& __r = ((QGraphicsSceneResizeEvent*)self)->oldSize();
	return new QSizeF(__r);
}
void QGraphicsSceneResizeEvent_setOldSize(void* self, void* size) {
	((QGraphicsSceneResizeEvent*)self)->setOldSize(*(QSizeF*)size);
}
void* QGraphicsSceneResizeEvent_newSize(void* self) {
	auto&& __r = ((QGraphicsSceneResizeEvent*)self)->newSize();
	return new QSizeF(__r);
}
void QGraphicsSceneResizeEvent_setNewSize(void* self, void* size) {
	((QGraphicsSceneResizeEvent*)self)->setNewSize(*(QSizeF*)size);
}
int QGraphicsSceneResizeEvent_registerEventType() {
	auto&& __r = QGraphicsSceneResizeEvent::registerEventType();
	return __r;
}
int QGraphicsSceneResizeEvent_registerEventType2(int hint) {
	auto&& __r = QGraphicsSceneResizeEvent::registerEventType(hint);
	return __r;
}
void QGraphicsSceneResizeEvent_setAccepted_default(void* self, int accepted) {
	((QGraphicsSceneResizeEvent*)self)->QGraphicsSceneResizeEvent::setAccepted(accepted);
}
void* QGraphicsSceneResizeEvent_clone_default(void* self) {
	auto&& __r = ((QGraphicsSceneResizeEvent*)self)->QGraphicsSceneResizeEvent::clone();
	return (void*)(__r);
}
void* QGraphicsSceneResizeEvent__Go_new(long long id) { return new QGraphicsSceneResizeEvent__Go(id); }
void QGraphicsSceneResizeEvent_delete(void* self) { delete (QGraphicsSceneResizeEvent*)self; }

}

#include <QtWidgets/qmenubar.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtGui/qaction.h>
#include <QtWidgets/qmenu.h>
#include <QtGui/qicon.h>
#include <QtCore/qsize.h>
#include <QtCore/qrect.h>
#include <QtCore/qpoint.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qcoreevent.h>
#include <QtGui/qevent.h>
#include <QtCore/qobject.h>
#include <QtWidgets/qstyleoption.h>
#include <QtWidgets/qstyle.h>
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
#include <QtGui/qkeysequence.h>
#include <QtWidgets/qgraphicsproxywidget.h>
#include <QtCore/qbytearray.h>
#include <QtWidgets/qsizepolicy.h>
#include <QtCore/qmargins.h>
#include <QtWidgets/qlayout.h>
#include <QtGui/qpaintengine.h>
#include <QtGui/qbackingstore.h>
#include <QtGui/qwindow.h>
#include <QtGui/qscreen.h>
#include <QtCore/qvariant.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qmenubar.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QMenuBar__Go : public QMenuBar {
public:
	long long __goId;
	unsigned char __goOn[49] = {};
	template<typename... A> QMenuBar__Go(long long id, A&&... a) : QMenuBar(std::forward<A>(a)...), __goId(id) {}
	using QMenuBar::updateMicroFocus;
	using QMenuBar::create;
	using QMenuBar::destroy;
	using QMenuBar::focusNextChild;
	using QMenuBar::focusPreviousChild;
	using QMenuBar::sender;
	using QMenuBar::senderSignalIndex;
	using QMenuBar::receivers;
	using QMenuBar::isSignalConnected;
	using QMenuBar::getDecodedMetricF;
	void cppgoBase_changeEvent(QEvent* __a0) { QMenuBar::changeEvent(__a0); }
	void cppgoBase_keyPressEvent(QKeyEvent* __a0) { QMenuBar::keyPressEvent(__a0); }
	void cppgoBase_mouseReleaseEvent(QMouseEvent* __a0) { QMenuBar::mouseReleaseEvent(__a0); }
	void cppgoBase_mousePressEvent(QMouseEvent* __a0) { QMenuBar::mousePressEvent(__a0); }
	void cppgoBase_mouseMoveEvent(QMouseEvent* __a0) { QMenuBar::mouseMoveEvent(__a0); }
	void cppgoBase_leaveEvent(QEvent* __a0) { QMenuBar::leaveEvent(__a0); }
	void cppgoBase_paintEvent(QPaintEvent* __a0) { QMenuBar::paintEvent(__a0); }
	void cppgoBase_resizeEvent(QResizeEvent* __a0) { QMenuBar::resizeEvent(__a0); }
	void cppgoBase_actionEvent(QActionEvent* __a0) { QMenuBar::actionEvent(__a0); }
	void cppgoBase_focusOutEvent(QFocusEvent* __a0) { QMenuBar::focusOutEvent(__a0); }
	void cppgoBase_focusInEvent(QFocusEvent* __a0) { QMenuBar::focusInEvent(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QMenuBar::timerEvent(__a0); }
	bool cppgoBase_eventFilter(QObject* __a0, QEvent* __a1) { return QMenuBar::eventFilter(__a0, __a1); }
	bool cppgoBase_event(QEvent* __a0) { return QMenuBar::event(__a0); }
	void cppgoBase_initStyleOption(QStyleOptionMenuItem* __a0, const QAction* __a1) { QMenuBar::initStyleOption(__a0, __a1); }
	void cppgoBase_mouseDoubleClickEvent(QMouseEvent* __a0) { QMenuBar::mouseDoubleClickEvent(__a0); }
	void cppgoBase_wheelEvent(QWheelEvent* __a0) { QMenuBar::wheelEvent(__a0); }
	void cppgoBase_keyReleaseEvent(QKeyEvent* __a0) { QMenuBar::keyReleaseEvent(__a0); }
	void cppgoBase_enterEvent(QEnterEvent* __a0) { QMenuBar::enterEvent(__a0); }
	void cppgoBase_moveEvent(QMoveEvent* __a0) { QMenuBar::moveEvent(__a0); }
	void cppgoBase_closeEvent(QCloseEvent* __a0) { QMenuBar::closeEvent(__a0); }
	void cppgoBase_contextMenuEvent(QContextMenuEvent* __a0) { QMenuBar::contextMenuEvent(__a0); }
	void cppgoBase_tabletEvent(QTabletEvent* __a0) { QMenuBar::tabletEvent(__a0); }
	void cppgoBase_dragEnterEvent(QDragEnterEvent* __a0) { QMenuBar::dragEnterEvent(__a0); }
	void cppgoBase_dragMoveEvent(QDragMoveEvent* __a0) { QMenuBar::dragMoveEvent(__a0); }
	void cppgoBase_dragLeaveEvent(QDragLeaveEvent* __a0) { QMenuBar::dragLeaveEvent(__a0); }
	void cppgoBase_dropEvent(QDropEvent* __a0) { QMenuBar::dropEvent(__a0); }
	void cppgoBase_showEvent(QShowEvent* __a0) { QMenuBar::showEvent(__a0); }
	void cppgoBase_hideEvent(QHideEvent* __a0) { QMenuBar::hideEvent(__a0); }
	int cppgoBase_metric(int __a0) { return QMenuBar::metric((QPaintDevice::PaintDeviceMetric)__a0); }
	void cppgoBase_initPainter(QPainter* __a0) { QMenuBar::initPainter(__a0); }
	QPaintDevice* cppgoBase_redirected(QPoint* __a0) { return QMenuBar::redirected(__a0); }
	QPainter* cppgoBase_sharedPainter() { return QMenuBar::sharedPainter(); }
	void cppgoBase_inputMethodEvent(QInputMethodEvent* __a0) { QMenuBar::inputMethodEvent(__a0); }
	bool cppgoBase_focusNextPrevChild(bool __a0) { return QMenuBar::focusNextPrevChild(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QMenuBar::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QMenuBar::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QMenuBar::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QMenuBar::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QMenuBar::metaObject();
	}
	QSize sizeHint() const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 729690261457684101LL, &__h); if (__h) return *(QSize*)__r; }
		return QMenuBar::sizeHint();
	}
	QSize minimumSizeHint() const override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 7049264599837532491LL, &__h); if (__h) return *(QSize*)__r; }
		return QMenuBar::minimumSizeHint();
	}
	int heightForWidth(int arg1) const override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 4311534060495737339LL, arg1, &__h); if (__h) return __r; }
		return QMenuBar::heightForWidth(arg1);
	}
	void setVisible(bool visible) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 2110182719375245575LL, visible, &__h); if (__h) return; }
		QMenuBar::setVisible(visible);
	}
	void changeEvent(QEvent* arg1) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4959376063826165943LL, (void*)arg1, &__h); if (__h) return; }
		QMenuBar::changeEvent(arg1);
	}
	void keyPressEvent(QKeyEvent* arg1) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1023791977679793185LL, (void*)arg1, &__h); if (__h) return; }
		QMenuBar::keyPressEvent(arg1);
	}
	void mouseReleaseEvent(QMouseEvent* arg1) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4415920049297795213LL, (void*)arg1, &__h); if (__h) return; }
		QMenuBar::mouseReleaseEvent(arg1);
	}
	void mousePressEvent(QMouseEvent* arg1) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -332090898078543319LL, (void*)arg1, &__h); if (__h) return; }
		QMenuBar::mousePressEvent(arg1);
	}
	void mouseMoveEvent(QMouseEvent* arg1) override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4255494004099806581LL, (void*)arg1, &__h); if (__h) return; }
		QMenuBar::mouseMoveEvent(arg1);
	}
	void leaveEvent(QEvent* arg1) override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -551070009623711640LL, (void*)arg1, &__h); if (__h) return; }
		QMenuBar::leaveEvent(arg1);
	}
	void paintEvent(QPaintEvent* arg1) override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5290442163167458465LL, (void*)arg1, &__h); if (__h) return; }
		QMenuBar::paintEvent(arg1);
	}
	void resizeEvent(QResizeEvent* arg1) override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4302917285655014661LL, (void*)arg1, &__h); if (__h) return; }
		QMenuBar::resizeEvent(arg1);
	}
	void actionEvent(QActionEvent* arg1) override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1697444140088555581LL, (void*)arg1, &__h); if (__h) return; }
		QMenuBar::actionEvent(arg1);
	}
	void focusOutEvent(QFocusEvent* arg1) override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -385596010576048993LL, (void*)arg1, &__h); if (__h) return; }
		QMenuBar::focusOutEvent(arg1);
	}
	void focusInEvent(QFocusEvent* arg1) override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 975043767033552552LL, (void*)arg1, &__h); if (__h) return; }
		QMenuBar::focusInEvent(arg1);
	}
	void timerEvent(QTimerEvent* arg1) override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)arg1, &__h); if (__h) return; }
		QMenuBar::timerEvent(arg1);
	}
	bool eventFilter(QObject* arg1, QEvent* arg2) override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)arg1, (void*)arg2, &__h); if (__h) return __r; }
		return QMenuBar::eventFilter(arg1, arg2);
	}
	bool event(QEvent* arg1) override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)arg1, &__h); if (__h) return __r; }
		return QMenuBar::event(arg1);
	}
	void initStyleOption(QStyleOptionMenuItem* option, const QAction* action) const override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_void_voidp_voidp) { ((void(*)(long long, long long, void*, void*, int*))cppgo_virt_void_voidp_voidp)(__goId, -488693769771308531LL, (void*)option, (void*)action, &__h); if (__h) return; }
		QMenuBar::initStyleOption(option, action);
	}
	int devType() const override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8659679893868524044LL, &__h); if (__h) return __r; }
		return QMenuBar::devType();
	}
	bool hasHeightForWidth() const override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8266762397995423915LL, &__h); if (__h) return __r; }
		return QMenuBar::hasHeightForWidth();
	}
	QPaintEngine* paintEngine() const override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 6398007404756346381LL, &__h); if (__h) return (QPaintEngine*)__r; }
		return QMenuBar::paintEngine();
	}
	void mouseDoubleClickEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4029073905697023887LL, (void*)event, &__h); if (__h) return; }
		QMenuBar::mouseDoubleClickEvent(event);
	}
	void wheelEvent(QWheelEvent* event) override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1919988340934511316LL, (void*)event, &__h); if (__h) return; }
		QMenuBar::wheelEvent(event);
	}
	void keyReleaseEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 772598475950050299LL, (void*)event, &__h); if (__h) return; }
		QMenuBar::keyReleaseEvent(event);
	}
	void enterEvent(QEnterEvent* event) override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6316270346668237895LL, (void*)event, &__h); if (__h) return; }
		QMenuBar::enterEvent(event);
	}
	void moveEvent(QMoveEvent* event) override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -9161949150489525396LL, (void*)event, &__h); if (__h) return; }
		QMenuBar::moveEvent(event);
	}
	void closeEvent(QCloseEvent* event) override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1989937050019464783LL, (void*)event, &__h); if (__h) return; }
		QMenuBar::closeEvent(event);
	}
	void contextMenuEvent(QContextMenuEvent* event) override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6698043222227006015LL, (void*)event, &__h); if (__h) return; }
		QMenuBar::contextMenuEvent(event);
	}
	void tabletEvent(QTabletEvent* event) override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -212068866860701141LL, (void*)event, &__h); if (__h) return; }
		QMenuBar::tabletEvent(event);
	}
	void dragEnterEvent(QDragEnterEvent* event) override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 206122715919917589LL, (void*)event, &__h); if (__h) return; }
		QMenuBar::dragEnterEvent(event);
	}
	void dragMoveEvent(QDragMoveEvent* event) override {
		int __h = 0;
		if (__goOn[32] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1781545137352745978LL, (void*)event, &__h); if (__h) return; }
		QMenuBar::dragMoveEvent(event);
	}
	void dragLeaveEvent(QDragLeaveEvent* event) override {
		int __h = 0;
		if (__goOn[33] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -3139904158373725006LL, (void*)event, &__h); if (__h) return; }
		QMenuBar::dragLeaveEvent(event);
	}
	void dropEvent(QDropEvent* event) override {
		int __h = 0;
		if (__goOn[34] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 9107058682754591716LL, (void*)event, &__h); if (__h) return; }
		QMenuBar::dropEvent(event);
	}
	void showEvent(QShowEvent* event) override {
		int __h = 0;
		if (__goOn[35] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6825369174751916LL, (void*)event, &__h); if (__h) return; }
		QMenuBar::showEvent(event);
	}
	void hideEvent(QHideEvent* event) override {
		int __h = 0;
		if (__goOn[36] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -199718707876719129LL, (void*)event, &__h); if (__h) return; }
		QMenuBar::hideEvent(event);
	}
	bool nativeEvent(const QByteArray& eventType, void* message, qintptr* result) override {
		int __h = 0;
		if (__goOn[37] && cppgo_virt_int_voidp_voidp_longlongp) { int __r = ((int(*)(long long, long long, void*, void*, long long*, int*))cppgo_virt_int_voidp_voidp_longlongp)(__goId, -2039166586246443706LL, (void*)&eventType, (void*)message, (long long*)result, &__h); if (__h) return __r; }
		return QMenuBar::nativeEvent(eventType, message, result);
	}
	int metric(QPaintDevice::PaintDeviceMetric arg1) const override {
		int __h = 0;
		if (__goOn[38] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 427629649267311777LL, (int)(arg1), &__h); if (__h) return __r; }
		return QMenuBar::metric(arg1);
	}
	void initPainter(QPainter* painter) const override {
		int __h = 0;
		if (__goOn[39] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6477500296980050210LL, (void*)painter, &__h); if (__h) return; }
		QMenuBar::initPainter(painter);
	}
	QPaintDevice* redirected(QPoint* offset) const override {
		int __h = 0;
		if (__goOn[40] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 4127240633310655048LL, (void*)offset, &__h); if (__h) return (QPaintDevice*)__r; }
		return QMenuBar::redirected(offset);
	}
	QPainter* sharedPainter() const override {
		int __h = 0;
		if (__goOn[41] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 5792771328741774991LL, &__h); if (__h) return (QPainter*)__r; }
		return QMenuBar::sharedPainter();
	}
	void inputMethodEvent(QInputMethodEvent* arg1) override {
		int __h = 0;
		if (__goOn[42] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5438819661942843222LL, (void*)arg1, &__h); if (__h) return; }
		QMenuBar::inputMethodEvent(arg1);
	}
	QVariant inputMethodQuery(Qt::InputMethodQuery arg1) const override {
		int __h = 0;
		if (__goOn[43] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7345853097340709638LL, (int)(arg1), &__h); if (__h) return *(QVariant*)__r; }
		return QMenuBar::inputMethodQuery(arg1);
	}
	bool focusNextPrevChild(bool next) override {
		int __h = 0;
		if (__goOn[44] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 5898328517505784191LL, next, &__h); if (__h) return __r; }
		return QMenuBar::focusNextPrevChild(next);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[45] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QMenuBar::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[46] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QMenuBar::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[47] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QMenuBar::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[48] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QMenuBar::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QMenuBar;
static std::mutex __connsMu_QMenuBar;

extern "C" {

void* QMenuBar_base(void* self) { return (void*)static_cast<QWidget*>((QMenuBar*)self); }
void QMenuBar_chain(void* self, void** out) {
	QMenuBar* __p = (QMenuBar*)self;
	QWidget* __b0 = static_cast<QWidget*>(__p); out[0] = (void*)__b0;
	QObject* __b1 = static_cast<QObject*>(__b0); out[1] = (void*)__b1;
}
void QMenuBar__set_override(void* self, int idx) { ((QMenuBar__Go*)self)->__goOn[idx] = 1; }
void* QMenuBar_fromQObject(void* o) { return (void*)qobject_cast<QMenuBar*>((QObject*)o); }
void* QMenuBar_new(long long __goid, void* parent) {
	return new QMenuBar__Go(__goid, (QWidget*)parent);
}
void* QMenuBar_new2(long long __goid) {
	return new QMenuBar__Go(__goid);
}
char* QMenuBar_tr(const char* s, long long* __out_len) {
	auto&& __r = QMenuBar::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QMenuBar_addMenu(void* self, void* menu) {
	auto&& __r = ((QMenuBar*)self)->addMenu((QMenu*)menu);
	return (void*)(__r);
}
void* QMenuBar_addMenu2(void* self, const char* title, long long title_len) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	auto&& __r = ((QMenuBar*)self)->addMenu(title_qstr);
	return (void*)(__r);
}
void* QMenuBar_addMenu3(void* self, void* icon, const char* title, long long title_len) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	auto&& __r = ((QMenuBar*)self)->addMenu(*(QIcon*)icon, title_qstr);
	return (void*)(__r);
}
void* QMenuBar_addSeparator(void* self) {
	auto&& __r = ((QMenuBar*)self)->addSeparator();
	return (void*)(__r);
}
void* QMenuBar_insertSeparator(void* self, void* before) {
	auto&& __r = ((QMenuBar*)self)->insertSeparator((QAction*)before);
	return (void*)(__r);
}
void* QMenuBar_insertMenu(void* self, void* before, void* menu) {
	auto&& __r = ((QMenuBar*)self)->insertMenu((QAction*)before, (QMenu*)menu);
	return (void*)(__r);
}
void QMenuBar_clear(void* self) {
	((QMenuBar*)self)->clear();
}
void* QMenuBar_activeAction(void* self) {
	auto&& __r = ((QMenuBar*)self)->activeAction();
	return (void*)(__r);
}
void QMenuBar_setActiveAction(void* self, void* action) {
	((QMenuBar*)self)->setActiveAction((QAction*)action);
}
void QMenuBar_setDefaultUp(void* self, int arg1) {
	((QMenuBar*)self)->setDefaultUp(arg1);
}
int QMenuBar_isDefaultUp(void* self) {
	auto&& __r = ((QMenuBar*)self)->isDefaultUp();
	return __r;
}
void* QMenuBar_actionGeometry(void* self, void* arg1) {
	auto&& __r = ((QMenuBar*)self)->actionGeometry((QAction*)arg1);
	return new QRect(__r);
}
void* QMenuBar_actionAt(void* self, void* arg1) {
	auto&& __r = ((QMenuBar*)self)->actionAt(*(QPoint*)arg1);
	return (void*)(__r);
}
void QMenuBar_setCornerWidget(void* self, void* w) {
	((QMenuBar*)self)->setCornerWidget((QWidget*)w);
}
void* QMenuBar_cornerWidget(void* self) {
	auto&& __r = ((QMenuBar*)self)->cornerWidget();
	return (void*)(__r);
}
int QMenuBar_isNativeMenuBar(void* self) {
	auto&& __r = ((QMenuBar*)self)->isNativeMenuBar();
	return __r;
}
void QMenuBar_setNativeMenuBar(void* self, int nativeMenuBar) {
	((QMenuBar*)self)->setNativeMenuBar(nativeMenuBar);
}
void QMenuBar_triggered(void* self, void* action) {
	((QMenuBar*)self)->triggered((QAction*)action);
}
void QMenuBar_hovered(void* self, void* action) {
	((QMenuBar*)self)->hovered((QAction*)action);
}
void QMenuBar_setTabOrder(void* arg1, void* arg2) {
	QMenuBar::setTabOrder((QWidget*)arg1, (QWidget*)arg2);
}
void* QMenuBar_mouseGrabber() {
	auto&& __r = QMenuBar::mouseGrabber();
	return (void*)(__r);
}
void* QMenuBar_keyboardGrabber() {
	auto&& __r = QMenuBar::keyboardGrabber();
	return (void*)(__r);
}
void* QMenuBar_find(unsigned long long arg1) {
	auto&& __r = QMenuBar::find(arg1);
	return (void*)(__r);
}
void* QMenuBar_createWindowContainer(void* window) {
	auto&& __r = QMenuBar::createWindowContainer((QWindow*)window);
	return (void*)(__r);
}
void QMenuBar_updateMicroFocus(void* self) {
	((QMenuBar__Go*)self)->updateMicroFocus();
}
void QMenuBar_create(void* self) {
	((QMenuBar__Go*)self)->create();
}
void QMenuBar_destroy(void* self) {
	((QMenuBar__Go*)self)->destroy();
}
int QMenuBar_focusNextChild(void* self) {
	auto&& __r = ((QMenuBar__Go*)self)->focusNextChild();
	return __r;
}
int QMenuBar_focusPreviousChild(void* self) {
	auto&& __r = ((QMenuBar__Go*)self)->focusPreviousChild();
	return __r;
}
void* QMenuBar_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QMenuBar::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QMenuBar_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QMenuBar::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QMenuBar_sender(void* self) {
	auto&& __r = ((QMenuBar__Go*)self)->sender();
	return (void*)(__r);
}
int QMenuBar_senderSignalIndex(void* self) {
	auto&& __r = ((QMenuBar__Go*)self)->senderSignalIndex();
	return __r;
}
int QMenuBar_receivers(void* self, const char* signal) {
	auto&& __r = ((QMenuBar__Go*)self)->receivers(signal);
	return __r;
}
int QMenuBar_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QMenuBar__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
double QMenuBar_devicePixelRatioFScale() {
	auto&& __r = QMenuBar::devicePixelRatioFScale();
	return __r;
}
int QMenuBar_encodeMetricF(int metric, double value) {
	auto&& __r = QMenuBar::encodeMetricF((QPaintDevice::PaintDeviceMetric)(metric), value);
	return __r;
}
double QMenuBar_getDecodedMetricF(void* self, int metricA, int metricB) {
	auto&& __r = ((QMenuBar__Go*)self)->getDecodedMetricF((QPaintDevice::PaintDeviceMetric)(metricA), (QPaintDevice::PaintDeviceMetric)(metricB));
	return __r;
}
char* QMenuBar_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QMenuBar::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QMenuBar_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QMenuBar::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QMenuBar_setCornerWidget2(void* self, void* w, int corner) {
	((QMenuBar*)self)->setCornerWidget((QWidget*)w, (Qt::Corner)(corner));
}
void* QMenuBar_cornerWidget2(void* self, int corner) {
	auto&& __r = ((QMenuBar*)self)->cornerWidget((Qt::Corner)(corner));
	return (void*)(__r);
}
void* QMenuBar_createWindowContainer2(void* window, void* parent) {
	auto&& __r = QMenuBar::createWindowContainer((QWindow*)window, (QWidget*)parent);
	return (void*)(__r);
}
void* QMenuBar_createWindowContainer3(void* window, void* parent, int flags) {
	auto&& __r = QMenuBar::createWindowContainer((QWindow*)window, (QWidget*)parent, QFlags<Qt::WindowType>(flags));
	return (void*)(__r);
}
void QMenuBar_updateMicroFocus2(void* self, int query) {
	((QMenuBar__Go*)self)->updateMicroFocus((Qt::InputMethodQuery)(query));
}
void QMenuBar_create2(void* self, unsigned long long arg1) {
	((QMenuBar__Go*)self)->create(arg1);
}
void QMenuBar_create3(void* self, unsigned long long arg1, int initializeWindow) {
	((QMenuBar__Go*)self)->create(arg1, initializeWindow);
}
void QMenuBar_create4(void* self, unsigned long long arg1, int initializeWindow, int destroyOldWindow) {
	((QMenuBar__Go*)self)->create(arg1, initializeWindow, destroyOldWindow);
}
void QMenuBar_destroy2(void* self, int destroyWindow) {
	((QMenuBar__Go*)self)->destroy(destroyWindow);
}
void QMenuBar_destroy3(void* self, int destroyWindow, int destroySubWindows) {
	((QMenuBar__Go*)self)->destroy(destroyWindow, destroySubWindows);
}
void* QMenuBar_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QMenuBar::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QMenuBar_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QMenuBar::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QMenuBar_metaObject_default(void* self) {
	auto&& __r = ((QMenuBar*)self)->QMenuBar::metaObject();
	return (void*)(__r);
}
void* QMenuBar_sizeHint_default(void* self) {
	auto&& __r = ((QMenuBar*)self)->QMenuBar::sizeHint();
	return new QSize(__r);
}
void* QMenuBar_minimumSizeHint_default(void* self) {
	auto&& __r = ((QMenuBar*)self)->QMenuBar::minimumSizeHint();
	return new QSize(__r);
}
int QMenuBar_heightForWidth_default(void* self, int arg1) {
	auto&& __r = ((QMenuBar*)self)->QMenuBar::heightForWidth(arg1);
	return __r;
}
void QMenuBar_setVisible_default(void* self, int visible) {
	((QMenuBar*)self)->QMenuBar::setVisible(visible);
}
void QMenuBar_changeEvent_default(void* self, void* arg1) {
	((QMenuBar__Go*)self)->cppgoBase_changeEvent((QEvent*)arg1);
}
void QMenuBar_keyPressEvent_default(void* self, void* arg1) {
	((QMenuBar__Go*)self)->cppgoBase_keyPressEvent((QKeyEvent*)arg1);
}
void QMenuBar_mouseReleaseEvent_default(void* self, void* arg1) {
	((QMenuBar__Go*)self)->cppgoBase_mouseReleaseEvent((QMouseEvent*)arg1);
}
void QMenuBar_mousePressEvent_default(void* self, void* arg1) {
	((QMenuBar__Go*)self)->cppgoBase_mousePressEvent((QMouseEvent*)arg1);
}
void QMenuBar_mouseMoveEvent_default(void* self, void* arg1) {
	((QMenuBar__Go*)self)->cppgoBase_mouseMoveEvent((QMouseEvent*)arg1);
}
void QMenuBar_leaveEvent_default(void* self, void* arg1) {
	((QMenuBar__Go*)self)->cppgoBase_leaveEvent((QEvent*)arg1);
}
void QMenuBar_paintEvent_default(void* self, void* arg1) {
	((QMenuBar__Go*)self)->cppgoBase_paintEvent((QPaintEvent*)arg1);
}
void QMenuBar_resizeEvent_default(void* self, void* arg1) {
	((QMenuBar__Go*)self)->cppgoBase_resizeEvent((QResizeEvent*)arg1);
}
void QMenuBar_actionEvent_default(void* self, void* arg1) {
	((QMenuBar__Go*)self)->cppgoBase_actionEvent((QActionEvent*)arg1);
}
void QMenuBar_focusOutEvent_default(void* self, void* arg1) {
	((QMenuBar__Go*)self)->cppgoBase_focusOutEvent((QFocusEvent*)arg1);
}
void QMenuBar_focusInEvent_default(void* self, void* arg1) {
	((QMenuBar__Go*)self)->cppgoBase_focusInEvent((QFocusEvent*)arg1);
}
void QMenuBar_timerEvent_default(void* self, void* arg1) {
	((QMenuBar__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)arg1);
}
int QMenuBar_eventFilter_default(void* self, void* arg1, void* arg2) {
	auto&& __r = ((QMenuBar__Go*)self)->cppgoBase_eventFilter((QObject*)arg1, (QEvent*)arg2);
	return __r;
}
int QMenuBar_event_default(void* self, void* arg1) {
	auto&& __r = ((QMenuBar__Go*)self)->cppgoBase_event((QEvent*)arg1);
	return __r;
}
void QMenuBar_initStyleOption_default(void* self, void* option, void* action) {
	((QMenuBar__Go*)self)->cppgoBase_initStyleOption((QStyleOptionMenuItem*)option, (QAction*)action);
}
int QMenuBar_devType_default(void* self) {
	auto&& __r = ((QMenuBar*)self)->QMenuBar::devType();
	return __r;
}
int QMenuBar_hasHeightForWidth_default(void* self) {
	auto&& __r = ((QMenuBar*)self)->QMenuBar::hasHeightForWidth();
	return __r;
}
void* QMenuBar_paintEngine_default(void* self) {
	auto&& __r = ((QMenuBar*)self)->QMenuBar::paintEngine();
	return (void*)(__r);
}
void QMenuBar_mouseDoubleClickEvent_default(void* self, void* event) {
	((QMenuBar__Go*)self)->cppgoBase_mouseDoubleClickEvent((QMouseEvent*)event);
}
void QMenuBar_wheelEvent_default(void* self, void* event) {
	((QMenuBar__Go*)self)->cppgoBase_wheelEvent((QWheelEvent*)event);
}
void QMenuBar_keyReleaseEvent_default(void* self, void* event) {
	((QMenuBar__Go*)self)->cppgoBase_keyReleaseEvent((QKeyEvent*)event);
}
void QMenuBar_enterEvent_default(void* self, void* event) {
	((QMenuBar__Go*)self)->cppgoBase_enterEvent((QEnterEvent*)event);
}
void QMenuBar_moveEvent_default(void* self, void* event) {
	((QMenuBar__Go*)self)->cppgoBase_moveEvent((QMoveEvent*)event);
}
void QMenuBar_closeEvent_default(void* self, void* event) {
	((QMenuBar__Go*)self)->cppgoBase_closeEvent((QCloseEvent*)event);
}
void QMenuBar_contextMenuEvent_default(void* self, void* event) {
	((QMenuBar__Go*)self)->cppgoBase_contextMenuEvent((QContextMenuEvent*)event);
}
void QMenuBar_tabletEvent_default(void* self, void* event) {
	((QMenuBar__Go*)self)->cppgoBase_tabletEvent((QTabletEvent*)event);
}
void QMenuBar_dragEnterEvent_default(void* self, void* event) {
	((QMenuBar__Go*)self)->cppgoBase_dragEnterEvent((QDragEnterEvent*)event);
}
void QMenuBar_dragMoveEvent_default(void* self, void* event) {
	((QMenuBar__Go*)self)->cppgoBase_dragMoveEvent((QDragMoveEvent*)event);
}
void QMenuBar_dragLeaveEvent_default(void* self, void* event) {
	((QMenuBar__Go*)self)->cppgoBase_dragLeaveEvent((QDragLeaveEvent*)event);
}
void QMenuBar_dropEvent_default(void* self, void* event) {
	((QMenuBar__Go*)self)->cppgoBase_dropEvent((QDropEvent*)event);
}
void QMenuBar_showEvent_default(void* self, void* event) {
	((QMenuBar__Go*)self)->cppgoBase_showEvent((QShowEvent*)event);
}
void QMenuBar_hideEvent_default(void* self, void* event) {
	((QMenuBar__Go*)self)->cppgoBase_hideEvent((QHideEvent*)event);
}
int QMenuBar_metric_default(void* self, int arg1) {
	auto&& __r = ((QMenuBar__Go*)self)->cppgoBase_metric(arg1);
	return __r;
}
void QMenuBar_initPainter_default(void* self, void* painter) {
	((QMenuBar__Go*)self)->cppgoBase_initPainter((QPainter*)painter);
}
void* QMenuBar_redirected_default(void* self, void* offset) {
	auto&& __r = ((QMenuBar__Go*)self)->cppgoBase_redirected((QPoint*)offset);
	return (void*)(__r);
}
void* QMenuBar_sharedPainter_default(void* self) {
	auto&& __r = ((QMenuBar__Go*)self)->cppgoBase_sharedPainter();
	return (void*)(__r);
}
void QMenuBar_inputMethodEvent_default(void* self, void* arg1) {
	((QMenuBar__Go*)self)->cppgoBase_inputMethodEvent((QInputMethodEvent*)arg1);
}
void* QMenuBar_inputMethodQuery_default(void* self, int arg1) {
	auto&& __r = ((QMenuBar*)self)->QMenuBar::inputMethodQuery((Qt::InputMethodQuery)(arg1));
	return new QVariant(__r);
}
int QMenuBar_focusNextPrevChild_default(void* self, int next) {
	auto&& __r = ((QMenuBar__Go*)self)->cppgoBase_focusNextPrevChild(next);
	return __r;
}
void QMenuBar_childEvent_default(void* self, void* event) {
	((QMenuBar__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QMenuBar_customEvent_default(void* self, void* event) {
	((QMenuBar__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QMenuBar_connectNotify_default(void* self, void* signal) {
	((QMenuBar__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QMenuBar_disconnectNotify_default(void* self, void* signal) {
	((QMenuBar__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QMenuBar_connect_triggered(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QMenuBar*)self, &QMenuBar::triggered, (QMenuBar*)self, [__id](QAction* action){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)action);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QMenuBar);
	__conns_QMenuBar[__id] = __c;
}
void QMenuBar_connect_hovered(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QMenuBar*)self, &QMenuBar::hovered, (QMenuBar*)self, [__id](QAction* action){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)action);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QMenuBar);
	__conns_QMenuBar[__id] = __c;
}
void QMenuBar_connect_windowTitleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QMenuBar*)self, &QMenuBar::windowTitleChanged, (QMenuBar*)self, [__id](const QString& title){
		QByteArray title_ba = title.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)title_ba.constData(), title_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QMenuBar);
	__conns_QMenuBar[__id] = __c;
}
void QMenuBar_connect_windowIconChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QMenuBar*)self, &QMenuBar::windowIconChanged, (QMenuBar*)self, [__id](const QIcon& icon){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&icon);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QMenuBar);
	__conns_QMenuBar[__id] = __c;
}
void QMenuBar_connect_windowIconTextChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QMenuBar*)self, &QMenuBar::windowIconTextChanged, (QMenuBar*)self, [__id](const QString& iconText){
		QByteArray iconText_ba = iconText.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)iconText_ba.constData(), iconText_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QMenuBar);
	__conns_QMenuBar[__id] = __c;
}
void QMenuBar_connect_customContextMenuRequested(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QMenuBar*)self, &QMenuBar::customContextMenuRequested, (QMenuBar*)self, [__id](const QPoint& pos){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&pos);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QMenuBar);
	__conns_QMenuBar[__id] = __c;
}
void QMenuBar_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QMenuBar*)self, &QMenuBar::destroyed, (QMenuBar*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QMenuBar);
	__conns_QMenuBar[__id] = __c;
}
void QMenuBar_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QMenuBar*)self, &QMenuBar::destroyed, (QMenuBar*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QMenuBar);
	__conns_QMenuBar[__id] = __c;
}
void QMenuBar_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QMenuBar*)self, &QMenuBar::objectNameChanged, (QMenuBar*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QMenuBar);
	__conns_QMenuBar[__id] = __c;
}
void QMenuBar__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QMenuBar);
		auto __it = __conns_QMenuBar.find(id);
		if (__it == __conns_QMenuBar.end()) return;
		__c = __it->second;
		__conns_QMenuBar.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QMenuBar__Go_new(long long id) { return new QMenuBar__Go(id); }
void QMenuBar_delete(void* self) { delete (QMenuBar*)self; }

}

#include <QtWidgets/qstylefactory.h>
#include <QtCore/qstring.h>
#include <QtWidgets/qstyle.h>
#include "gen_qstylefactory.h"

extern "C" {

char** QStyleFactory_keys(long long* __out_len) {
	auto&& __r = QStyleFactory::keys();
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
void* QStyleFactory_create(const char* arg1, long long arg1_len) {
	QString arg1_qstr = QString::fromUtf8(arg1, arg1_len);
	auto&& __r = QStyleFactory::create(arg1_qstr);
	return (void*)(__r);
}
void QStyleFactory_delete(void* self) { delete (QStyleFactory*)self; }

}

