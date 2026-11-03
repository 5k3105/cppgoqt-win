#include <QtCore/qsysinfo.h>
#include <QtCore/qstring.h>
#include <QtCore/qbytearray.h>
#include "gen_qsysinfo.h"

extern "C" {

char* QSysInfo_buildCpuArchitecture(long long* __out_len) {
	auto&& __r = QSysInfo::buildCpuArchitecture();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QSysInfo_currentCpuArchitecture(long long* __out_len) {
	auto&& __r = QSysInfo::currentCpuArchitecture();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QSysInfo_buildAbi(long long* __out_len) {
	auto&& __r = QSysInfo::buildAbi();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QSysInfo_kernelType(long long* __out_len) {
	auto&& __r = QSysInfo::kernelType();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QSysInfo_kernelVersion(long long* __out_len) {
	auto&& __r = QSysInfo::kernelVersion();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QSysInfo_productType(long long* __out_len) {
	auto&& __r = QSysInfo::productType();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QSysInfo_productVersion(long long* __out_len) {
	auto&& __r = QSysInfo::productVersion();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QSysInfo_prettyProductName(long long* __out_len) {
	auto&& __r = QSysInfo::prettyProductName();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QSysInfo_machineHostName(long long* __out_len) {
	auto&& __r = QSysInfo::machineHostName();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QSysInfo_machineUniqueId() {
	auto&& __r = QSysInfo::machineUniqueId();
	return new QByteArray(__r);
}
void* QSysInfo_bootUniqueId() {
	auto&& __r = QSysInfo::bootUniqueId();
	return new QByteArray(__r);
}
void QSysInfo_delete(void* self) { delete (QSysInfo*)self; }

}

#include <QtCore/qbytearray.h>
#include "gen_qbytearray__frombase64result.h"

extern "C" {

void* QByteArray__FromBase64Result_new(void* arg1) {
	return new QByteArray::FromBase64Result(*(QByteArray::FromBase64Result*)arg1);
}
void* QByteArray__FromBase64Result_new2() {
	return new QByteArray::FromBase64Result();
}
void QByteArray__FromBase64Result_swap(void* self, void* other) {
	((QByteArray::FromBase64Result*)self)->swap(*(QByteArray::FromBase64Result*)other);
}
int QByteArray__FromBase64Result_ToBool(void* self) {
	auto&& __r = static_cast<bool>(*(QByteArray::FromBase64Result*)self);
	return __r;
}
void* QByteArray__FromBase64Result_OperatorMultiply(void* self) {
	auto&& __r = ((QByteArray::FromBase64Result*)self)->operator*();
	return (void*)&(__r);
}
void* QByteArray__FromBase64Result_OperatorMultiply2(void* self) {
	auto&& __r = ((QByteArray::FromBase64Result*)self)->operator*();
	return (void*)&(__r);
}
void* QByteArray__FromBase64Result_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QByteArray::FromBase64Result*)self)->operator=(*(QByteArray::FromBase64Result*)arg1);
	return (void*)&(__r);
}
int QByteArray__FromBase64Result_OperatorEqual(void* self, void* rhs) {
	auto&& __r = (*(QByteArray::FromBase64Result*)self) == (*(QByteArray::FromBase64Result*)rhs);
	return __r;
}
int QByteArray__FromBase64Result_OperatorNotEqual(void* self, void* rhs) {
	auto&& __r = (*(QByteArray::FromBase64Result*)self) != (*(QByteArray::FromBase64Result*)rhs);
	return __r;
}
void* QByteArray__FromBase64Result_decoded_get(void* self) {
	auto __r = ((QByteArray::FromBase64Result*)self)->decoded;
	return new QByteArray(__r);
}
void QByteArray__FromBase64Result_decoded_set(void* self, void* decoded) {
	((QByteArray::FromBase64Result*)self)->decoded = *(QByteArray*)decoded;
}
int QByteArray__FromBase64Result_decodingStatus_get(void* self) {
	auto __r = ((QByteArray::FromBase64Result*)self)->decodingStatus;
	return (int)(__r);
}
void QByteArray__FromBase64Result_decodingStatus_set(void* self, int decodingStatus) {
	((QByteArray::FromBase64Result*)self)->decodingStatus = (QByteArray::Base64DecodingStatus)(decodingStatus);
}
void QByteArray__FromBase64Result_delete(void* self) { delete (QByteArray::FromBase64Result*)self; }

}

#include <QtCore/qobject.h>
#include "gen_qsignalblocker.h"

extern "C" {

void* QSignalBlocker_new(void* o) {
	return new QSignalBlocker((QObject*)o);
}
void* QSignalBlocker_new2(void* o) {
	return new QSignalBlocker(*(QObject*)o);
}
void QSignalBlocker_reblock(void* self) {
	((QSignalBlocker*)self)->reblock();
}
void QSignalBlocker_unblock(void* self) {
	((QSignalBlocker*)self)->unblock();
}
void QSignalBlocker_dismiss(void* self) {
	((QSignalBlocker*)self)->dismiss();
}
void QSignalBlocker_delete(void* self) { delete (QSignalBlocker*)self; }

}

#include <QtCore/qabstractitemmodel.h>
#include <QtCore/qvariant.h>
#include <QtCore/qnamespace.h>
#include "gen_qmodelindex.h"

extern "C" {

void* QModelIndex_new() {
	return new QModelIndex();
}
void* QModelIndex_new2(void* arg1) {
	return new QModelIndex(*(QModelIndex*)arg1);
}
int QModelIndex_row(void* self) {
	auto&& __r = ((QModelIndex*)self)->row();
	return __r;
}
int QModelIndex_column(void* self) {
	auto&& __r = ((QModelIndex*)self)->column();
	return __r;
}
unsigned long long QModelIndex_internalId(void* self) {
	auto&& __r = ((QModelIndex*)self)->internalId();
	return __r;
}
void* QModelIndex_parent(void* self) {
	auto&& __r = ((QModelIndex*)self)->parent();
	return new QModelIndex(__r);
}
void* QModelIndex_sibling(void* self, int row, int column) {
	auto&& __r = ((QModelIndex*)self)->sibling(row, column);
	return new QModelIndex(__r);
}
void* QModelIndex_siblingAtColumn(void* self, int column) {
	auto&& __r = ((QModelIndex*)self)->siblingAtColumn(column);
	return new QModelIndex(__r);
}
void* QModelIndex_siblingAtRow(void* self, int row) {
	auto&& __r = ((QModelIndex*)self)->siblingAtRow(row);
	return new QModelIndex(__r);
}
void* QModelIndex_data(void* self) {
	auto&& __r = ((QModelIndex*)self)->data();
	return new QVariant(__r);
}
void QModelIndex_multiData(void* self, void* roleDataSpan) {
	((QModelIndex*)self)->multiData(*(QModelRoleDataSpan*)roleDataSpan);
}
int QModelIndex_flags(void* self) {
	auto&& __r = ((QModelIndex*)self)->flags();
	return (int)(__r);
}
void* QModelIndex_model(void* self) {
	auto&& __r = ((QModelIndex*)self)->model();
	return (void*)(__r);
}
int QModelIndex_isValid(void* self) {
	auto&& __r = ((QModelIndex*)self)->isValid();
	return __r;
}
void* QModelIndex_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QModelIndex*)self)->operator=(*(QModelIndex*)arg1);
	return (void*)&(__r);
}
void* QModelIndex_data2(void* self, int role) {
	auto&& __r = ((QModelIndex*)self)->data(role);
	return new QVariant(__r);
}
int QModelIndex_OperatorEqual(void* self, void* rhs) {
	auto&& __r = (*(QModelIndex*)self) == (*(QModelIndex*)rhs);
	return __r;
}
int QModelIndex_OperatorNotEqual(void* self, void* rhs) {
	auto&& __r = (*(QModelIndex*)self) != (*(QModelIndex*)rhs);
	return __r;
}
int QModelIndex_OperatorLesser(void* self, void* rhs) {
	auto&& __r = (*(QModelIndex*)self) < (*(QModelIndex*)rhs);
	return __r;
}
int QModelIndex_OperatorGreater(void* self, void* rhs) {
	auto&& __r = (*(QModelIndex*)self) > (*(QModelIndex*)rhs);
	return __r;
}
int QModelIndex_OperatorLesserOrEqual(void* self, void* rhs) {
	auto&& __r = (*(QModelIndex*)self) <= (*(QModelIndex*)rhs);
	return __r;
}
int QModelIndex_OperatorGreaterOrEqual(void* self, void* rhs) {
	auto&& __r = (*(QModelIndex*)self) >= (*(QModelIndex*)rhs);
	return __r;
}
int QModelIndex_OperatorEqual2(void* self, void* rhs) {
	auto&& __r = (*(QModelIndex*)self) == (*(QPersistentModelIndex*)rhs);
	return __r;
}
int QModelIndex_OperatorNotEqual2(void* self, void* rhs) {
	auto&& __r = (*(QModelIndex*)self) != (*(QPersistentModelIndex*)rhs);
	return __r;
}
int QModelIndex_OperatorLesser2(void* self, void* rhs) {
	auto&& __r = (*(QModelIndex*)self) < (*(QPersistentModelIndex*)rhs);
	return __r;
}
int QModelIndex_OperatorGreater2(void* self, void* rhs) {
	auto&& __r = (*(QModelIndex*)self) > (*(QPersistentModelIndex*)rhs);
	return __r;
}
int QModelIndex_OperatorLesserOrEqual2(void* self, void* rhs) {
	auto&& __r = (*(QModelIndex*)self) <= (*(QPersistentModelIndex*)rhs);
	return __r;
}
int QModelIndex_OperatorGreaterOrEqual2(void* self, void* rhs) {
	auto&& __r = (*(QModelIndex*)self) >= (*(QPersistentModelIndex*)rhs);
	return __r;
}
void QModelIndex_delete(void* self) { delete (QModelIndex*)self; }

}

#include <QtCore/qbytearraymatcher.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qbytearrayview.h>
#include "gen_qbytearraymatcher.h"

extern "C" {

void* QByteArrayMatcher_new() {
	return new QByteArrayMatcher();
}
void* QByteArrayMatcher_new2(void* pattern) {
	return new QByteArrayMatcher(*(QByteArray*)pattern);
}
void* QByteArrayMatcher_new3(void* pattern) {
	return new QByteArrayMatcher(*(QByteArrayView*)pattern);
}
void* QByteArrayMatcher_new4(const char* pattern) {
	return new QByteArrayMatcher(pattern);
}
void* QByteArrayMatcher_new5(void* other) {
	return new QByteArrayMatcher(*(QByteArrayMatcher*)other);
}
void* QByteArrayMatcher_new6(const char* pattern, long long length) {
	return new QByteArrayMatcher(pattern, length);
}
void* QByteArrayMatcher_OperatorAssign(void* self, void* other) {
	auto&& __r = ((QByteArrayMatcher*)self)->operator=(*(QByteArrayMatcher*)other);
	return (void*)&(__r);
}
void QByteArrayMatcher_setPattern(void* self, void* pattern) {
	((QByteArrayMatcher*)self)->setPattern(*(QByteArray*)pattern);
}
long long QByteArrayMatcher_indexIn(void* self, const char* str, long long len_) {
	auto&& __r = ((QByteArrayMatcher*)self)->indexIn(str, len_);
	return __r;
}
long long QByteArrayMatcher_indexIn2(void* self, void* data) {
	auto&& __r = ((QByteArrayMatcher*)self)->indexIn(*(QByteArrayView*)data);
	return __r;
}
void* QByteArrayMatcher_pattern(void* self) {
	auto&& __r = ((QByteArrayMatcher*)self)->pattern();
	return new QByteArray(__r);
}
long long QByteArrayMatcher_indexIn3(void* self, const char* str, long long len_, long long from) {
	auto&& __r = ((QByteArrayMatcher*)self)->indexIn(str, len_, from);
	return __r;
}
long long QByteArrayMatcher_indexIn4(void* self, void* data, long long from) {
	auto&& __r = ((QByteArrayMatcher*)self)->indexIn(*(QByteArrayView*)data, from);
	return __r;
}
void QByteArrayMatcher_delete(void* self) { delete (QByteArrayMatcher*)self; }

}

#include <QtCore/qpropertyprivate.h>
#include "gen_quntypedpropertydata.h"

extern "C" {

void QUntypedPropertyData_delete(void* self) { delete (QUntypedPropertyData*)self; }

}

#include <QtCore/qtimezone.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qstring.h>
#include <QtCore/qlocale.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qbytearrayview.h>
#include <QtCore/qdatetime.h>
#include "gen_qtimezone.h"

extern "C" {

void* QTimeZone_new() {
	return new QTimeZone();
}
void* QTimeZone_new2(int spec) {
	return new QTimeZone((QTimeZone::Initialization)(spec));
}
void* QTimeZone_new3(int offsetSeconds) {
	return new QTimeZone(offsetSeconds);
}
void* QTimeZone_new4(void* ianaId) {
	return new QTimeZone(*(QByteArray*)ianaId);
}
void* QTimeZone_new5(void* zoneId, int offsetSeconds, const char* name, long long name_len, const char* abbreviation, long long abbreviation_len) {
	QString name_qstr = QString::fromUtf8(name, name_len);
	QString abbreviation_qstr = QString::fromUtf8(abbreviation, abbreviation_len);
	return new QTimeZone(*(QByteArray*)zoneId, offsetSeconds, name_qstr, abbreviation_qstr);
}
void* QTimeZone_new6(void* other) {
	return new QTimeZone(*(QTimeZone*)other);
}
void* QTimeZone_new7(void* zoneId, int offsetSeconds, const char* name, long long name_len, const char* abbreviation, long long abbreviation_len, int territory) {
	QString name_qstr = QString::fromUtf8(name, name_len);
	QString abbreviation_qstr = QString::fromUtf8(abbreviation, abbreviation_len);
	return new QTimeZone(*(QByteArray*)zoneId, offsetSeconds, name_qstr, abbreviation_qstr, (QLocale::Country)(territory));
}
void* QTimeZone_new8(void* zoneId, int offsetSeconds, const char* name, long long name_len, const char* abbreviation, long long abbreviation_len, int territory, const char* comment, long long comment_len) {
	QString name_qstr = QString::fromUtf8(name, name_len);
	QString abbreviation_qstr = QString::fromUtf8(abbreviation, abbreviation_len);
	QString comment_qstr = QString::fromUtf8(comment, comment_len);
	return new QTimeZone(*(QByteArray*)zoneId, offsetSeconds, name_qstr, abbreviation_qstr, (QLocale::Country)(territory), comment_qstr);
}
void* QTimeZone_OperatorAssign(void* self, void* other) {
	auto&& __r = ((QTimeZone*)self)->operator=(*(QTimeZone*)other);
	return (void*)&(__r);
}
void QTimeZone_swap(void* self, void* other) {
	((QTimeZone*)self)->swap(*(QTimeZone*)other);
}
int QTimeZone_isValid(void* self) {
	auto&& __r = ((QTimeZone*)self)->isValid();
	return __r;
}
void* QTimeZone_fromDurationAheadOfUtc(long long offset) {
	auto&& __r = QTimeZone::fromDurationAheadOfUtc(std::chrono::seconds((long long)offset));
	return new QTimeZone(__r);
}
void* QTimeZone_fromSecondsAheadOfUtc(int offset) {
	auto&& __r = QTimeZone::fromSecondsAheadOfUtc(offset);
	return new QTimeZone(__r);
}
int QTimeZone_timeSpec(void* self) {
	auto&& __r = ((QTimeZone*)self)->timeSpec();
	return (int)(__r);
}
int QTimeZone_fixedSecondsAheadOfUtc(void* self) {
	auto&& __r = ((QTimeZone*)self)->fixedSecondsAheadOfUtc();
	return __r;
}
int QTimeZone_isUtcOrFixedOffset(int spec) {
	auto&& __r = QTimeZone::isUtcOrFixedOffset((Qt::TimeSpec)(spec));
	return __r;
}
int QTimeZone_isUtcOrFixedOffset2(void* self) {
	auto&& __r = ((QTimeZone*)self)->isUtcOrFixedOffset();
	return __r;
}
void* QTimeZone_asBackendZone(void* self) {
	auto&& __r = ((QTimeZone*)self)->asBackendZone();
	return new QTimeZone(__r);
}
int QTimeZone_hasAlternativeName(void* self, void* alias) {
	auto&& __r = ((QTimeZone*)self)->hasAlternativeName(*(QByteArrayView*)alias);
	return __r;
}
void* QTimeZone_id(void* self) {
	auto&& __r = ((QTimeZone*)self)->id();
	return new QByteArray(__r);
}
int QTimeZone_territory(void* self) {
	auto&& __r = ((QTimeZone*)self)->territory();
	return (int)(__r);
}
int QTimeZone_country(void* self) {
	auto&& __r = ((QTimeZone*)self)->country();
	return (int)(__r);
}
char* QTimeZone_comment(void* self, long long* __out_len) {
	auto&& __r = ((QTimeZone*)self)->comment();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QTimeZone_displayName(void* self, void* atDateTime, long long* __out_len) {
	auto&& __r = ((QTimeZone*)self)->displayName(*(QDateTime*)atDateTime);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QTimeZone_displayName2(void* self, int timeType, long long* __out_len) {
	auto&& __r = ((QTimeZone*)self)->displayName((QTimeZone::TimeType)(timeType));
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QTimeZone_abbreviation(void* self, void* atDateTime, long long* __out_len) {
	auto&& __r = ((QTimeZone*)self)->abbreviation(*(QDateTime*)atDateTime);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QTimeZone_offsetFromUtc(void* self, void* atDateTime) {
	auto&& __r = ((QTimeZone*)self)->offsetFromUtc(*(QDateTime*)atDateTime);
	return __r;
}
int QTimeZone_standardTimeOffset(void* self, void* atDateTime) {
	auto&& __r = ((QTimeZone*)self)->standardTimeOffset(*(QDateTime*)atDateTime);
	return __r;
}
int QTimeZone_daylightTimeOffset(void* self, void* atDateTime) {
	auto&& __r = ((QTimeZone*)self)->daylightTimeOffset(*(QDateTime*)atDateTime);
	return __r;
}
int QTimeZone_hasDaylightTime(void* self) {
	auto&& __r = ((QTimeZone*)self)->hasDaylightTime();
	return __r;
}
int QTimeZone_isDaylightTime(void* self, void* atDateTime) {
	auto&& __r = ((QTimeZone*)self)->isDaylightTime(*(QDateTime*)atDateTime);
	return __r;
}
void* QTimeZone_offsetData(void* self, void* forDateTime) {
	auto&& __r = ((QTimeZone*)self)->offsetData(*(QDateTime*)forDateTime);
	return new QTimeZone::OffsetData(__r);
}
int QTimeZone_hasTransitions(void* self) {
	auto&& __r = ((QTimeZone*)self)->hasTransitions();
	return __r;
}
void* QTimeZone_nextTransition(void* self, void* afterDateTime) {
	auto&& __r = ((QTimeZone*)self)->nextTransition(*(QDateTime*)afterDateTime);
	return new QTimeZone::OffsetData(__r);
}
void* QTimeZone_previousTransition(void* self, void* beforeDateTime) {
	auto&& __r = ((QTimeZone*)self)->previousTransition(*(QDateTime*)beforeDateTime);
	return new QTimeZone::OffsetData(__r);
}
void** QTimeZone_transitions(void* self, void* fromDateTime, void* toDateTime, long long* __out_len) {
	auto&& __r = ((QTimeZone*)self)->transitions(*(QDateTime*)fromDateTime, *(QDateTime*)toDateTime);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QTimeZone::OffsetData(__e);
	return __arr;
}
void* QTimeZone_systemTimeZoneId() {
	auto&& __r = QTimeZone::systemTimeZoneId();
	return new QByteArray(__r);
}
void* QTimeZone_systemTimeZone() {
	auto&& __r = QTimeZone::systemTimeZone();
	return new QTimeZone(__r);
}
void* QTimeZone_utc() {
	auto&& __r = QTimeZone::utc();
	return new QTimeZone(__r);
}
int QTimeZone_isTimeZoneIdAvailable(void* ianaId) {
	auto&& __r = QTimeZone::isTimeZoneIdAvailable(*(QByteArray*)ianaId);
	return __r;
}
void** QTimeZone_availableTimeZoneIds(long long* __out_len) {
	auto&& __r = QTimeZone::availableTimeZoneIds();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QByteArray(__e);
	return __arr;
}
void** QTimeZone_availableTimeZoneIds2(int territory, long long* __out_len) {
	auto&& __r = QTimeZone::availableTimeZoneIds((QLocale::Country)(territory));
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QByteArray(__e);
	return __arr;
}
void** QTimeZone_availableTimeZoneIds3(int offsetSeconds, long long* __out_len) {
	auto&& __r = QTimeZone::availableTimeZoneIds(offsetSeconds);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QByteArray(__e);
	return __arr;
}
void* QTimeZone_ianaIdToWindowsId(void* ianaId) {
	auto&& __r = QTimeZone::ianaIdToWindowsId(*(QByteArray*)ianaId);
	return new QByteArray(__r);
}
void* QTimeZone_windowsIdToDefaultIanaId(void* windowsId) {
	auto&& __r = QTimeZone::windowsIdToDefaultIanaId(*(QByteArray*)windowsId);
	return new QByteArray(__r);
}
void* QTimeZone_windowsIdToDefaultIanaId2(void* windowsId, int territory) {
	auto&& __r = QTimeZone::windowsIdToDefaultIanaId(*(QByteArray*)windowsId, (QLocale::Country)(territory));
	return new QByteArray(__r);
}
void** QTimeZone_windowsIdToIanaIds(void* windowsId, long long* __out_len) {
	auto&& __r = QTimeZone::windowsIdToIanaIds(*(QByteArray*)windowsId);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QByteArray(__e);
	return __arr;
}
void** QTimeZone_windowsIdToIanaIds2(void* windowsId, int territory, long long* __out_len) {
	auto&& __r = QTimeZone::windowsIdToIanaIds(*(QByteArray*)windowsId, (QLocale::Country)(territory));
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QByteArray(__e);
	return __arr;
}
char* QTimeZone_displayName3(void* self, void* atDateTime, int nameType, long long* __out_len) {
	auto&& __r = ((QTimeZone*)self)->displayName(*(QDateTime*)atDateTime, (QTimeZone::NameType)(nameType));
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QTimeZone_displayName4(void* self, void* atDateTime, int nameType, void* locale, long long* __out_len) {
	auto&& __r = ((QTimeZone*)self)->displayName(*(QDateTime*)atDateTime, (QTimeZone::NameType)(nameType), *(QLocale*)locale);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QTimeZone_displayName5(void* self, int timeType, int nameType, long long* __out_len) {
	auto&& __r = ((QTimeZone*)self)->displayName((QTimeZone::TimeType)(timeType), (QTimeZone::NameType)(nameType));
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QTimeZone_displayName6(void* self, int timeType, int nameType, void* locale, long long* __out_len) {
	auto&& __r = ((QTimeZone*)self)->displayName((QTimeZone::TimeType)(timeType), (QTimeZone::NameType)(nameType), *(QLocale*)locale);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QTimeZone_OperatorEqual(void* self, void* rhs) {
	auto&& __r = (*(QTimeZone*)self) == (*(QTimeZone*)rhs);
	return __r;
}
int QTimeZone_OperatorNotEqual(void* self, void* rhs) {
	auto&& __r = (*(QTimeZone*)self) != (*(QTimeZone*)rhs);
	return __r;
}
void QTimeZone_delete(void* self) { delete (QTimeZone*)self; }

}

#include <QtCore/qjsonvalue.h>
#include <QtCore/qvariant.h>
#include <QtCore/qstring.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qjsonarray.h>
#include <QtCore/qjsonobject.h>
#include <QtCore/qstringview.h>
#include <QtCore/qlatin1stringview.h>
#include "gen_qjsonvalueconstref.h"

extern "C" {

void* QJsonValueConstRef_new(void* arg1) {
	return new QJsonValueConstRef(*(QJsonValueConstRef*)arg1);
}
void* QJsonValueConstRef_ToQJsonValue(void* self) {
	auto&& __r = static_cast<QJsonValue>(*(QJsonValueConstRef*)self);
	return new QJsonValue(__r);
}
void* QJsonValueConstRef_toVariant(void* self) {
	auto&& __r = ((QJsonValueConstRef*)self)->toVariant();
	return new QVariant(__r);
}
int QJsonValueConstRef_type(void* self) {
	auto&& __r = ((QJsonValueConstRef*)self)->type();
	return (int)(__r);
}
int QJsonValueConstRef_isNull(void* self) {
	auto&& __r = ((QJsonValueConstRef*)self)->isNull();
	return __r;
}
int QJsonValueConstRef_isBool(void* self) {
	auto&& __r = ((QJsonValueConstRef*)self)->isBool();
	return __r;
}
int QJsonValueConstRef_isDouble(void* self) {
	auto&& __r = ((QJsonValueConstRef*)self)->isDouble();
	return __r;
}
int QJsonValueConstRef_isString(void* self) {
	auto&& __r = ((QJsonValueConstRef*)self)->isString();
	return __r;
}
int QJsonValueConstRef_isArray(void* self) {
	auto&& __r = ((QJsonValueConstRef*)self)->isArray();
	return __r;
}
int QJsonValueConstRef_isObject(void* self) {
	auto&& __r = ((QJsonValueConstRef*)self)->isObject();
	return __r;
}
int QJsonValueConstRef_isUndefined(void* self) {
	auto&& __r = ((QJsonValueConstRef*)self)->isUndefined();
	return __r;
}
int QJsonValueConstRef_toBool(void* self) {
	auto&& __r = ((QJsonValueConstRef*)self)->toBool();
	return __r;
}
int QJsonValueConstRef_toInt(void* self) {
	auto&& __r = ((QJsonValueConstRef*)self)->toInt();
	return __r;
}
long long QJsonValueConstRef_toInteger(void* self) {
	auto&& __r = ((QJsonValueConstRef*)self)->toInteger();
	return __r;
}
double QJsonValueConstRef_toDouble(void* self) {
	auto&& __r = ((QJsonValueConstRef*)self)->toDouble();
	return __r;
}
char* QJsonValueConstRef_toString(void* self, long long* __out_len) {
	auto&& __r = ((QJsonValueConstRef*)self)->toString();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QJsonValueConstRef_toStringView(void* self) {
	auto&& __r = ((QJsonValueConstRef*)self)->toStringView();
	return new QAnyStringView(__r);
}
void* QJsonValueConstRef_toArray(void* self) {
	auto&& __r = ((QJsonValueConstRef*)self)->toArray();
	return new QJsonArray(__r);
}
void* QJsonValueConstRef_toObject(void* self) {
	auto&& __r = ((QJsonValueConstRef*)self)->toObject();
	return new QJsonObject(__r);
}
void* QJsonValueConstRef_OperatorSubscript(void* self, void* key) {
	auto&& __r = ((QJsonValueConstRef*)self)->operator[](*(QStringView*)key);
	return new QJsonValue(__r);
}
void* QJsonValueConstRef_OperatorSubscript2(void* self, void* key) {
	auto&& __r = ((QJsonValueConstRef*)self)->operator[](*(QLatin1String*)key);
	return new QJsonValue(__r);
}
void* QJsonValueConstRef_OperatorSubscript3(void* self, long long i) {
	auto&& __r = ((QJsonValueConstRef*)self)->operator[](i);
	return new QJsonValue(__r);
}
int QJsonValueConstRef_toBool2(void* self, int defaultValue) {
	auto&& __r = ((QJsonValueConstRef*)self)->toBool(defaultValue);
	return __r;
}
int QJsonValueConstRef_toInt2(void* self, int defaultValue) {
	auto&& __r = ((QJsonValueConstRef*)self)->toInt(defaultValue);
	return __r;
}
long long QJsonValueConstRef_toInteger2(void* self, long long defaultValue) {
	auto&& __r = ((QJsonValueConstRef*)self)->toInteger(defaultValue);
	return __r;
}
double QJsonValueConstRef_toDouble2(void* self, double defaultValue) {
	auto&& __r = ((QJsonValueConstRef*)self)->toDouble(defaultValue);
	return __r;
}
char* QJsonValueConstRef_toString2(void* self, const char* defaultValue, long long defaultValue_len, long long* __out_len) {
	QString defaultValue_qstr = QString::fromUtf8(defaultValue, defaultValue_len);
	auto&& __r = ((QJsonValueConstRef*)self)->toString(defaultValue_qstr);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QJsonValueConstRef_toStringView2(void* self, const char* defaultValue, long long defaultValue_len) {
	QAnyStringView defaultValue_qstr = QAnyStringView(defaultValue, defaultValue_len);
	auto&& __r = ((QJsonValueConstRef*)self)->toStringView(defaultValue_qstr);
	return new QAnyStringView(__r);
}
int QJsonValueConstRef_OperatorEqual(void* self, void* rhs) {
	auto&& __r = (*(QJsonValueConstRef*)self) == (*(QJsonValueConstRef*)rhs);
	return __r;
}
int QJsonValueConstRef_OperatorNotEqual(void* self, void* rhs) {
	auto&& __r = (*(QJsonValueConstRef*)self) != (*(QJsonValueConstRef*)rhs);
	return __r;
}
int QJsonValueConstRef_OperatorEqual2(void* self, void* rhs) {
	auto&& __r = (*(QJsonValueConstRef*)self) == (*(QJsonValue*)rhs);
	return __r;
}
int QJsonValueConstRef_OperatorNotEqual2(void* self, void* rhs) {
	auto&& __r = (*(QJsonValueConstRef*)self) != (*(QJsonValue*)rhs);
	return __r;
}
int QJsonValueConstRef_OperatorEqual3(void* self, void* rhs) {
	auto&& __r = (*(QJsonValueConstRef*)self) == (*(QJsonValueRef*)rhs);
	return __r;
}
int QJsonValueConstRef_OperatorNotEqual3(void* self, void* rhs) {
	auto&& __r = (*(QJsonValueConstRef*)self) != (*(QJsonValueRef*)rhs);
	return __r;
}
int QJsonValueConstRef_OperatorEqual4(void* self, void* rhs) {
	auto&& __r = (*(QJsonValueConstRef*)self) == (*(QJsonObject*)rhs);
	return __r;
}
int QJsonValueConstRef_OperatorNotEqual4(void* self, void* rhs) {
	auto&& __r = (*(QJsonValueConstRef*)self) != (*(QJsonObject*)rhs);
	return __r;
}
void QJsonValueConstRef_delete(void* self) { delete (QJsonValueConstRef*)self; }

}

#include <QtCore/qmimedatabase.h>
#include <QtCore/qmimetype.h>
#include <QtCore/qstring.h>
#include <QtCore/qfileinfo.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qiodevice.h>
#include <QtCore/qurl.h>
#include "gen_qmimedatabase.h"

extern "C" {

void* QMimeDatabase_new() {
	return new QMimeDatabase();
}
void* QMimeDatabase_mimeTypeForName(void* self, const char* nameOrAlias, long long nameOrAlias_len) {
	QString nameOrAlias_qstr = QString::fromUtf8(nameOrAlias, nameOrAlias_len);
	auto&& __r = ((QMimeDatabase*)self)->mimeTypeForName(nameOrAlias_qstr);
	return new QMimeType(__r);
}
void* QMimeDatabase_mimeTypeForFile(void* self, const char* fileName, long long fileName_len) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	auto&& __r = ((QMimeDatabase*)self)->mimeTypeForFile(fileName_qstr);
	return new QMimeType(__r);
}
void* QMimeDatabase_mimeTypeForFile2(void* self, void* fileInfo) {
	auto&& __r = ((QMimeDatabase*)self)->mimeTypeForFile(*(QFileInfo*)fileInfo);
	return new QMimeType(__r);
}
void** QMimeDatabase_mimeTypesForFileName(void* self, const char* fileName, long long fileName_len, long long* __out_len) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	auto&& __r = ((QMimeDatabase*)self)->mimeTypesForFileName(fileName_qstr);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QMimeType(__e);
	return __arr;
}
void* QMimeDatabase_mimeTypeForData(void* self, void* data) {
	auto&& __r = ((QMimeDatabase*)self)->mimeTypeForData(*(QByteArray*)data);
	return new QMimeType(__r);
}
void* QMimeDatabase_mimeTypeForData2(void* self, void* device) {
	auto&& __r = ((QMimeDatabase*)self)->mimeTypeForData((QIODevice*)device);
	return new QMimeType(__r);
}
void* QMimeDatabase_mimeTypeForUrl(void* self, void* url) {
	auto&& __r = ((QMimeDatabase*)self)->mimeTypeForUrl(*(QUrl*)url);
	return new QMimeType(__r);
}
void* QMimeDatabase_mimeTypeForFileNameAndData(void* self, const char* fileName, long long fileName_len, void* device) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	auto&& __r = ((QMimeDatabase*)self)->mimeTypeForFileNameAndData(fileName_qstr, (QIODevice*)device);
	return new QMimeType(__r);
}
void* QMimeDatabase_mimeTypeForFileNameAndData2(void* self, const char* fileName, long long fileName_len, void* data) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	auto&& __r = ((QMimeDatabase*)self)->mimeTypeForFileNameAndData(fileName_qstr, *(QByteArray*)data);
	return new QMimeType(__r);
}
char* QMimeDatabase_suffixForFileName(void* self, const char* fileName, long long fileName_len, long long* __out_len) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	auto&& __r = ((QMimeDatabase*)self)->suffixForFileName(fileName_qstr);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void** QMimeDatabase_allMimeTypes(void* self, long long* __out_len) {
	auto&& __r = ((QMimeDatabase*)self)->allMimeTypes();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QMimeType(__e);
	return __arr;
}
void* QMimeDatabase_mimeTypeForFile3(void* self, const char* fileName, long long fileName_len, int mode) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	auto&& __r = ((QMimeDatabase*)self)->mimeTypeForFile(fileName_qstr, (QMimeDatabase::MatchMode)(mode));
	return new QMimeType(__r);
}
void* QMimeDatabase_mimeTypeForFile4(void* self, void* fileInfo, int mode) {
	auto&& __r = ((QMimeDatabase*)self)->mimeTypeForFile(*(QFileInfo*)fileInfo, (QMimeDatabase::MatchMode)(mode));
	return new QMimeType(__r);
}
void QMimeDatabase_delete(void* self) { delete (QMimeDatabase*)self; }

}

#include <QtCore/qpropertyanimation.h>
#include <QtCore/qvariantanimation.h>
#include <QtCore/qobject.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qvariant.h>
#include <QtCore/qabstractanimation.h>
#include <QtCore/qeasingcurve.h>
#include <QtCore/qanimationgroup.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qpropertyanimation.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QPropertyAnimation__Go : public QPropertyAnimation {
public:
	long long __goId;
	unsigned char __goOn[14] = {};
	template<typename... A> QPropertyAnimation__Go(long long id, A&&... a) : QPropertyAnimation(std::forward<A>(a)...), __goId(id) {}
	using QPropertyAnimation::sender;
	using QPropertyAnimation::senderSignalIndex;
	using QPropertyAnimation::receivers;
	using QPropertyAnimation::isSignalConnected;
	bool cppgoBase_event(QEvent* __a0) { return QPropertyAnimation::event(__a0); }
	void cppgoBase_updateCurrentValue(const QVariant& __a0) { QPropertyAnimation::updateCurrentValue(__a0); }
	void cppgoBase_updateState(int __a0, int __a1) { QPropertyAnimation::updateState((QAbstractAnimation::State)__a0, (QAbstractAnimation::State)__a1); }
	void cppgoBase_updateCurrentTime(int __a0) { QPropertyAnimation::updateCurrentTime(__a0); }
	QVariant cppgoBase_interpolated(const QVariant& __a0, const QVariant& __a1, qreal __a2) { return QPropertyAnimation::interpolated(__a0, __a1, __a2); }
	void cppgoBase_updateDirection(int __a0) { QPropertyAnimation::updateDirection((QAbstractAnimation::Direction)__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QPropertyAnimation::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QPropertyAnimation::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QPropertyAnimation::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QPropertyAnimation::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QPropertyAnimation::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QPropertyAnimation::metaObject();
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QPropertyAnimation::event(event);
	}
	void updateCurrentValue(const QVariant& value) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4717240701590708418LL, (void*)&value, &__h); if (__h) return; }
		QPropertyAnimation::updateCurrentValue(value);
	}
	void updateState(QAbstractAnimation::State newState, QAbstractAnimation::State oldState) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_int_int) { ((void(*)(long long, long long, int, int, int*))cppgo_virt_void_int_int)(__goId, 7692482089134926569LL, (int)(newState), (int)(oldState), &__h); if (__h) return; }
		QPropertyAnimation::updateState(newState, oldState);
	}
	int duration() const override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8434675112194400115LL, &__h); if (__h) return __r; }
		return QPropertyAnimation::duration();
	}
	void updateCurrentTime(int arg1) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, -1342418961870150968LL, arg1, &__h); if (__h) return; }
		QPropertyAnimation::updateCurrentTime(arg1);
	}
	QVariant interpolated(const QVariant& from, const QVariant& to, qreal progress) const override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_voidp_voidp_voidp_unsignedlonglong) { void* __r = ((void*(*)(long long, long long, void*, void*, unsigned long long, int*))cppgo_virt_voidp_voidp_voidp_unsignedlonglong)(__goId, 4185842742461511100LL, (void*)&from, (void*)&to, __builtin_bit_cast(unsigned long long, (double)(progress)), &__h); if (__h) return *(QVariant*)__r; }
		return QPropertyAnimation::interpolated(from, to, progress);
	}
	void updateDirection(QAbstractAnimation::Direction direction) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, -3231888725407197643LL, (int)(direction), &__h); if (__h) return; }
		QPropertyAnimation::updateDirection(direction);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QPropertyAnimation::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QPropertyAnimation::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QPropertyAnimation::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QPropertyAnimation::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QPropertyAnimation::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QPropertyAnimation::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QPropertyAnimation;
static std::mutex __connsMu_QPropertyAnimation;

extern "C" {

void* QPropertyAnimation_base(void* self) { return (void*)static_cast<QVariantAnimation*>((QPropertyAnimation*)self); }
void QPropertyAnimation_chain(void* self, void** out) {
	QPropertyAnimation* __p = (QPropertyAnimation*)self;
	QVariantAnimation* __b0 = static_cast<QVariantAnimation*>(__p); out[0] = (void*)__b0;
	QAbstractAnimation* __b1 = static_cast<QAbstractAnimation*>(__b0); out[1] = (void*)__b1;
	QObject* __b2 = static_cast<QObject*>(__b1); out[2] = (void*)__b2;
}
void QPropertyAnimation__set_override(void* self, int idx) { ((QPropertyAnimation__Go*)self)->__goOn[idx] = 1; }
void* QPropertyAnimation_fromQObject(void* o) { return (void*)qobject_cast<QPropertyAnimation*>((QObject*)o); }
void* QPropertyAnimation_new(long long __goid) {
	return new QPropertyAnimation__Go(__goid);
}
void* QPropertyAnimation_new2(long long __goid, void* target, void* propertyName) {
	return new QPropertyAnimation__Go(__goid, (QObject*)target, *(QByteArray*)propertyName);
}
void* QPropertyAnimation_new3(long long __goid, void* parent) {
	return new QPropertyAnimation__Go(__goid, (QObject*)parent);
}
void* QPropertyAnimation_new4(long long __goid, void* target, void* propertyName, void* parent) {
	return new QPropertyAnimation__Go(__goid, (QObject*)target, *(QByteArray*)propertyName, (QObject*)parent);
}
char* QPropertyAnimation_tr(const char* s, long long* __out_len) {
	auto&& __r = QPropertyAnimation::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QPropertyAnimation_targetObject(void* self) {
	auto&& __r = ((QPropertyAnimation*)self)->targetObject();
	return (void*)(__r);
}
void QPropertyAnimation_setTargetObject(void* self, void* target) {
	((QPropertyAnimation*)self)->setTargetObject((QObject*)target);
}
void* QPropertyAnimation_propertyName(void* self) {
	auto&& __r = ((QPropertyAnimation*)self)->propertyName();
	return new QByteArray(__r);
}
void QPropertyAnimation_setPropertyName(void* self, void* propertyName) {
	((QPropertyAnimation*)self)->setPropertyName(*(QByteArray*)propertyName);
}
void* QPropertyAnimation_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QPropertyAnimation::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QPropertyAnimation_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QPropertyAnimation::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QPropertyAnimation_sender(void* self) {
	auto&& __r = ((QPropertyAnimation__Go*)self)->sender();
	return (void*)(__r);
}
int QPropertyAnimation_senderSignalIndex(void* self) {
	auto&& __r = ((QPropertyAnimation__Go*)self)->senderSignalIndex();
	return __r;
}
int QPropertyAnimation_receivers(void* self, const char* signal) {
	auto&& __r = ((QPropertyAnimation__Go*)self)->receivers(signal);
	return __r;
}
int QPropertyAnimation_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QPropertyAnimation__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QPropertyAnimation_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QPropertyAnimation::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QPropertyAnimation_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QPropertyAnimation::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QPropertyAnimation_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QPropertyAnimation::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QPropertyAnimation_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QPropertyAnimation::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QPropertyAnimation_metaObject_default(void* self) {
	auto&& __r = ((QPropertyAnimation*)self)->QPropertyAnimation::metaObject();
	return (void*)(__r);
}
int QPropertyAnimation_event_default(void* self, void* event) {
	auto&& __r = ((QPropertyAnimation__Go*)self)->cppgoBase_event((QEvent*)event);
	return __r;
}
void QPropertyAnimation_updateCurrentValue_default(void* self, void* value) {
	((QPropertyAnimation__Go*)self)->cppgoBase_updateCurrentValue(*(QVariant*)value);
}
void QPropertyAnimation_updateState_default(void* self, int newState, int oldState) {
	((QPropertyAnimation__Go*)self)->cppgoBase_updateState(newState, oldState);
}
int QPropertyAnimation_duration_default(void* self) {
	auto&& __r = ((QPropertyAnimation*)self)->QPropertyAnimation::duration();
	return __r;
}
void QPropertyAnimation_updateCurrentTime_default(void* self, int arg1) {
	((QPropertyAnimation__Go*)self)->cppgoBase_updateCurrentTime(arg1);
}
void* QPropertyAnimation_interpolated_default(void* self, void* from, void* to, double progress) {
	auto&& __r = ((QPropertyAnimation__Go*)self)->cppgoBase_interpolated(*(QVariant*)from, *(QVariant*)to, progress);
	return new QVariant(__r);
}
void QPropertyAnimation_updateDirection_default(void* self, int direction) {
	((QPropertyAnimation__Go*)self)->cppgoBase_updateDirection(direction);
}
int QPropertyAnimation_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QPropertyAnimation*)self)->QPropertyAnimation::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QPropertyAnimation_timerEvent_default(void* self, void* event) {
	((QPropertyAnimation__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QPropertyAnimation_childEvent_default(void* self, void* event) {
	((QPropertyAnimation__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QPropertyAnimation_customEvent_default(void* self, void* event) {
	((QPropertyAnimation__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QPropertyAnimation_connectNotify_default(void* self, void* signal) {
	((QPropertyAnimation__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QPropertyAnimation_disconnectNotify_default(void* self, void* signal) {
	((QPropertyAnimation__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QPropertyAnimation_connect_valueChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QPropertyAnimation*)self, &QPropertyAnimation::valueChanged, (QPropertyAnimation*)self, [__id](const QVariant& value){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&value);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QPropertyAnimation);
	__conns_QPropertyAnimation[__id] = __c;
}
void QPropertyAnimation_connect_finished(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QPropertyAnimation*)self, &QPropertyAnimation::finished, (QPropertyAnimation*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QPropertyAnimation);
	__conns_QPropertyAnimation[__id] = __c;
}
void QPropertyAnimation_connect_stateChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QPropertyAnimation*)self, &QPropertyAnimation::stateChanged, (QPropertyAnimation*)self, [__id](QAbstractAnimation::State newState, QAbstractAnimation::State oldState){
		if (cppgo_sig_void_int_int) ((void(*)(long long, int, int))cppgo_sig_void_int_int)(__id, (int)(newState), (int)(oldState));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QPropertyAnimation);
	__conns_QPropertyAnimation[__id] = __c;
}
void QPropertyAnimation_connect_currentLoopChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QPropertyAnimation*)self, &QPropertyAnimation::currentLoopChanged, (QPropertyAnimation*)self, [__id](int currentLoop){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, currentLoop);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QPropertyAnimation);
	__conns_QPropertyAnimation[__id] = __c;
}
void QPropertyAnimation_connect_directionChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QPropertyAnimation*)self, &QPropertyAnimation::directionChanged, (QPropertyAnimation*)self, [__id](QAbstractAnimation::Direction arg1){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, (int)(arg1));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QPropertyAnimation);
	__conns_QPropertyAnimation[__id] = __c;
}
void QPropertyAnimation_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QPropertyAnimation*)self, &QPropertyAnimation::destroyed, (QPropertyAnimation*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QPropertyAnimation);
	__conns_QPropertyAnimation[__id] = __c;
}
void QPropertyAnimation_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QPropertyAnimation*)self, &QPropertyAnimation::destroyed, (QPropertyAnimation*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QPropertyAnimation);
	__conns_QPropertyAnimation[__id] = __c;
}
void QPropertyAnimation_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QPropertyAnimation*)self, &QPropertyAnimation::objectNameChanged, (QPropertyAnimation*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QPropertyAnimation);
	__conns_QPropertyAnimation[__id] = __c;
}
void QPropertyAnimation__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QPropertyAnimation);
		auto __it = __conns_QPropertyAnimation.find(id);
		if (__it == __conns_QPropertyAnimation.end()) return;
		__c = __it->second;
		__conns_QPropertyAnimation.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QPropertyAnimation__Go_new(long long id, void* parent) { return new QPropertyAnimation__Go(id, (QObject*)parent); }
void QPropertyAnimation_delete(void* self) { delete (QPropertyAnimation*)self; }

}

#include <QtCore/qsocketnotifier.h>
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
#include "gen_qsocketnotifier.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

static std::unordered_map<long long, QMetaObject::Connection> __conns_QSocketNotifier;
static std::mutex __connsMu_QSocketNotifier;

extern "C" {

void* QSocketNotifier_base(void* self) { return (void*)static_cast<QObject*>((QSocketNotifier*)self); }
void* QSocketNotifier_fromQObject(void* o) { return (void*)qobject_cast<QSocketNotifier*>((QObject*)o); }
void* QSocketNotifier_new(int arg1) {
	return new QSocketNotifier((QSocketNotifier::Type)(arg1));
}
void* QSocketNotifier_new2(long long socket, int arg2) {
	return new QSocketNotifier(socket, (QSocketNotifier::Type)(arg2));
}
void* QSocketNotifier_new3(int arg1, void* parent) {
	return new QSocketNotifier((QSocketNotifier::Type)(arg1), (QObject*)parent);
}
void* QSocketNotifier_new4(long long socket, int arg2, void* parent) {
	return new QSocketNotifier(socket, (QSocketNotifier::Type)(arg2), (QObject*)parent);
}
char* QSocketNotifier_tr(const char* s, long long* __out_len) {
	auto&& __r = QSocketNotifier::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QSocketNotifier_setSocket(void* self, long long socket) {
	((QSocketNotifier*)self)->setSocket(socket);
}
long long QSocketNotifier_socket(void* self) {
	auto&& __r = ((QSocketNotifier*)self)->socket();
	return __r;
}
int QSocketNotifier_type(void* self) {
	auto&& __r = ((QSocketNotifier*)self)->type();
	return (int)(__r);
}
int QSocketNotifier_isValid(void* self) {
	auto&& __r = ((QSocketNotifier*)self)->isValid();
	return __r;
}
int QSocketNotifier_isEnabled(void* self) {
	auto&& __r = ((QSocketNotifier*)self)->isEnabled();
	return __r;
}
void QSocketNotifier_setEnabled(void* self, int arg1) {
	((QSocketNotifier*)self)->setEnabled(arg1);
}
void* QSocketNotifier_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QSocketNotifier::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QSocketNotifier_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QSocketNotifier::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
char* QSocketNotifier_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QSocketNotifier::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QSocketNotifier_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QSocketNotifier::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QSocketNotifier_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QSocketNotifier::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QSocketNotifier_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QSocketNotifier::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void QSocketNotifier_connect_activated(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSocketNotifier*)self, &QSocketNotifier::activated, (QSocketNotifier*)self, [__id](QSocketDescriptor socket, QSocketNotifier::Type activationEvent){
		if (cppgo_sig_void_voidp_int) ((void(*)(long long, void*, int))cppgo_sig_void_voidp_int)(__id, (void*)&socket, (int)(activationEvent));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSocketNotifier);
	__conns_QSocketNotifier[__id] = __c;
}
void QSocketNotifier_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSocketNotifier*)self, &QSocketNotifier::destroyed, (QSocketNotifier*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSocketNotifier);
	__conns_QSocketNotifier[__id] = __c;
}
void QSocketNotifier_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSocketNotifier*)self, &QSocketNotifier::destroyed, (QSocketNotifier*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSocketNotifier);
	__conns_QSocketNotifier[__id] = __c;
}
void QSocketNotifier_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSocketNotifier*)self, &QSocketNotifier::objectNameChanged, (QSocketNotifier*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSocketNotifier);
	__conns_QSocketNotifier[__id] = __c;
}
void QSocketNotifier__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QSocketNotifier);
		auto __it = __conns_QSocketNotifier.find(id);
		if (__it == __conns_QSocketNotifier.end()) return;
		__c = __it->second;
		__conns_QSocketNotifier.erase(__it);
	}
	QObject::disconnect(__c);
}
void QSocketNotifier_delete(void* self) { delete (QSocketNotifier*)self; }

}

#include <QtCore/qxmlstream.h>
#include <QtCore/qiodevice.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qstring.h>
#include <QtCore/qanystringview.h>
#include "gen_qxmlstreamwriter.h"

extern "C" {

void* QXmlStreamWriter_new() {
	return new QXmlStreamWriter();
}
void* QXmlStreamWriter_new2(void* device) {
	return new QXmlStreamWriter((QIODevice*)device);
}
void* QXmlStreamWriter_new3(void* array) {
	return new QXmlStreamWriter((QByteArray*)array);
}
void* QXmlStreamWriter_new4(void* string_) {
	return new QXmlStreamWriter((QString*)string_);
}
void QXmlStreamWriter_setDevice(void* self, void* device) {
	((QXmlStreamWriter*)self)->setDevice((QIODevice*)device);
}
void* QXmlStreamWriter_device(void* self) {
	auto&& __r = ((QXmlStreamWriter*)self)->device();
	return (void*)(__r);
}
void QXmlStreamWriter_setAutoFormatting(void* self, int arg1) {
	((QXmlStreamWriter*)self)->setAutoFormatting(arg1);
}
int QXmlStreamWriter_autoFormatting(void* self) {
	auto&& __r = ((QXmlStreamWriter*)self)->autoFormatting();
	return __r;
}
void QXmlStreamWriter_setAutoFormattingIndent(void* self, int spacesOrTabs) {
	((QXmlStreamWriter*)self)->setAutoFormattingIndent(spacesOrTabs);
}
int QXmlStreamWriter_autoFormattingIndent(void* self) {
	auto&& __r = ((QXmlStreamWriter*)self)->autoFormattingIndent();
	return __r;
}
void QXmlStreamWriter_setStopWritingOnError(void* self, int stop) {
	((QXmlStreamWriter*)self)->setStopWritingOnError(stop);
}
int QXmlStreamWriter_stopWritingOnError(void* self) {
	auto&& __r = ((QXmlStreamWriter*)self)->stopWritingOnError();
	return __r;
}
void QXmlStreamWriter_writeAttribute(void* self, const char* qualifiedName, long long qualifiedName_len, const char* value, long long value_len) {
	QAnyStringView qualifiedName_qstr = QAnyStringView(qualifiedName, qualifiedName_len);
	QAnyStringView value_qstr = QAnyStringView(value, value_len);
	((QXmlStreamWriter*)self)->writeAttribute(qualifiedName_qstr, value_qstr);
}
void QXmlStreamWriter_writeAttribute2(void* self, const char* namespaceUri, long long namespaceUri_len, const char* name, long long name_len, const char* value, long long value_len) {
	QAnyStringView namespaceUri_qstr = QAnyStringView(namespaceUri, namespaceUri_len);
	QAnyStringView name_qstr = QAnyStringView(name, name_len);
	QAnyStringView value_qstr = QAnyStringView(value, value_len);
	((QXmlStreamWriter*)self)->writeAttribute(namespaceUri_qstr, name_qstr, value_qstr);
}
void QXmlStreamWriter_writeAttribute3(void* self, void* attribute) {
	((QXmlStreamWriter*)self)->writeAttribute(*(QXmlStreamAttribute*)attribute);
}
void QXmlStreamWriter_writeAttributes(void* self, void* attributes) {
	((QXmlStreamWriter*)self)->writeAttributes(*(QXmlStreamAttributes*)attributes);
}
void QXmlStreamWriter_writeCDATA(void* self, const char* text, long long text_len) {
	QAnyStringView text_qstr = QAnyStringView(text, text_len);
	((QXmlStreamWriter*)self)->writeCDATA(text_qstr);
}
void QXmlStreamWriter_writeCharacters(void* self, const char* text, long long text_len) {
	QAnyStringView text_qstr = QAnyStringView(text, text_len);
	((QXmlStreamWriter*)self)->writeCharacters(text_qstr);
}
void QXmlStreamWriter_writeComment(void* self, const char* text, long long text_len) {
	QAnyStringView text_qstr = QAnyStringView(text, text_len);
	((QXmlStreamWriter*)self)->writeComment(text_qstr);
}
void QXmlStreamWriter_writeDTD(void* self, const char* dtd, long long dtd_len) {
	QAnyStringView dtd_qstr = QAnyStringView(dtd, dtd_len);
	((QXmlStreamWriter*)self)->writeDTD(dtd_qstr);
}
void QXmlStreamWriter_writeEmptyElement(void* self, const char* qualifiedName, long long qualifiedName_len) {
	QAnyStringView qualifiedName_qstr = QAnyStringView(qualifiedName, qualifiedName_len);
	((QXmlStreamWriter*)self)->writeEmptyElement(qualifiedName_qstr);
}
void QXmlStreamWriter_writeEmptyElement2(void* self, const char* namespaceUri, long long namespaceUri_len, const char* name, long long name_len) {
	QAnyStringView namespaceUri_qstr = QAnyStringView(namespaceUri, namespaceUri_len);
	QAnyStringView name_qstr = QAnyStringView(name, name_len);
	((QXmlStreamWriter*)self)->writeEmptyElement(namespaceUri_qstr, name_qstr);
}
void QXmlStreamWriter_writeTextElement(void* self, const char* qualifiedName, long long qualifiedName_len, const char* text, long long text_len) {
	QAnyStringView qualifiedName_qstr = QAnyStringView(qualifiedName, qualifiedName_len);
	QAnyStringView text_qstr = QAnyStringView(text, text_len);
	((QXmlStreamWriter*)self)->writeTextElement(qualifiedName_qstr, text_qstr);
}
void QXmlStreamWriter_writeTextElement2(void* self, const char* namespaceUri, long long namespaceUri_len, const char* name, long long name_len, const char* text, long long text_len) {
	QAnyStringView namespaceUri_qstr = QAnyStringView(namespaceUri, namespaceUri_len);
	QAnyStringView name_qstr = QAnyStringView(name, name_len);
	QAnyStringView text_qstr = QAnyStringView(text, text_len);
	((QXmlStreamWriter*)self)->writeTextElement(namespaceUri_qstr, name_qstr, text_qstr);
}
void QXmlStreamWriter_writeEndDocument(void* self) {
	((QXmlStreamWriter*)self)->writeEndDocument();
}
void QXmlStreamWriter_writeEndElement(void* self) {
	((QXmlStreamWriter*)self)->writeEndElement();
}
void QXmlStreamWriter_writeEntityReference(void* self, const char* name, long long name_len) {
	QAnyStringView name_qstr = QAnyStringView(name, name_len);
	((QXmlStreamWriter*)self)->writeEntityReference(name_qstr);
}
void QXmlStreamWriter_writeNamespace(void* self, const char* namespaceUri, long long namespaceUri_len) {
	QAnyStringView namespaceUri_qstr = QAnyStringView(namespaceUri, namespaceUri_len);
	((QXmlStreamWriter*)self)->writeNamespace(namespaceUri_qstr);
}
void QXmlStreamWriter_writeDefaultNamespace(void* self, const char* namespaceUri, long long namespaceUri_len) {
	QAnyStringView namespaceUri_qstr = QAnyStringView(namespaceUri, namespaceUri_len);
	((QXmlStreamWriter*)self)->writeDefaultNamespace(namespaceUri_qstr);
}
void QXmlStreamWriter_writeProcessingInstruction(void* self, const char* target, long long target_len) {
	QAnyStringView target_qstr = QAnyStringView(target, target_len);
	((QXmlStreamWriter*)self)->writeProcessingInstruction(target_qstr);
}
void QXmlStreamWriter_writeStartDocument(void* self) {
	((QXmlStreamWriter*)self)->writeStartDocument();
}
void QXmlStreamWriter_writeStartDocument2(void* self, const char* version, long long version_len) {
	QAnyStringView version_qstr = QAnyStringView(version, version_len);
	((QXmlStreamWriter*)self)->writeStartDocument(version_qstr);
}
void QXmlStreamWriter_writeStartDocument3(void* self, const char* version, long long version_len, int standalone) {
	QAnyStringView version_qstr = QAnyStringView(version, version_len);
	((QXmlStreamWriter*)self)->writeStartDocument(version_qstr, standalone);
}
void QXmlStreamWriter_writeStartElement(void* self, const char* qualifiedName, long long qualifiedName_len) {
	QAnyStringView qualifiedName_qstr = QAnyStringView(qualifiedName, qualifiedName_len);
	((QXmlStreamWriter*)self)->writeStartElement(qualifiedName_qstr);
}
void QXmlStreamWriter_writeStartElement2(void* self, const char* namespaceUri, long long namespaceUri_len, const char* name, long long name_len) {
	QAnyStringView namespaceUri_qstr = QAnyStringView(namespaceUri, namespaceUri_len);
	QAnyStringView name_qstr = QAnyStringView(name, name_len);
	((QXmlStreamWriter*)self)->writeStartElement(namespaceUri_qstr, name_qstr);
}
void QXmlStreamWriter_writeCurrentToken(void* self, void* reader) {
	((QXmlStreamWriter*)self)->writeCurrentToken(*(QXmlStreamReader*)reader);
}
void QXmlStreamWriter_raiseError(void* self, const char* message, long long message_len) {
	QAnyStringView message_qstr = QAnyStringView(message, message_len);
	((QXmlStreamWriter*)self)->raiseError(message_qstr);
}
char* QXmlStreamWriter_errorString(void* self, long long* __out_len) {
	auto&& __r = ((QXmlStreamWriter*)self)->errorString();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QXmlStreamWriter_error(void* self) {
	auto&& __r = ((QXmlStreamWriter*)self)->error();
	return (int)(__r);
}
int QXmlStreamWriter_hasError(void* self) {
	auto&& __r = ((QXmlStreamWriter*)self)->hasError();
	return __r;
}
void QXmlStreamWriter_writeNamespace2(void* self, const char* namespaceUri, long long namespaceUri_len, const char* prefix, long long prefix_len) {
	QAnyStringView namespaceUri_qstr = QAnyStringView(namespaceUri, namespaceUri_len);
	QAnyStringView prefix_qstr = QAnyStringView(prefix, prefix_len);
	((QXmlStreamWriter*)self)->writeNamespace(namespaceUri_qstr, prefix_qstr);
}
void QXmlStreamWriter_writeProcessingInstruction2(void* self, const char* target, long long target_len, const char* data, long long data_len) {
	QAnyStringView target_qstr = QAnyStringView(target, target_len);
	QAnyStringView data_qstr = QAnyStringView(data, data_len);
	((QXmlStreamWriter*)self)->writeProcessingInstruction(target_qstr, data_qstr);
}
void QXmlStreamWriter_delete(void* self) { delete (QXmlStreamWriter*)self; }

}

#include <QtGui/qevent.h>
#include <QtCore/qcoreevent.h>
#include <QtGui/qinputdevice.h>
#include <QtCore/qnamespace.h>
#include "gen_qinputevent.h"

extern "C" {

void* QInputEvent_base(void* self) { return (void*)static_cast<QEvent*>((QInputEvent*)self); }
void* QInputEvent_new(int type_, void* m_dev) {
	return new QInputEvent((QEvent::Type)(type_), (QInputDevice*)m_dev);
}
void* QInputEvent_new2(int type_, void* m_dev, int modifiers) {
	return new QInputEvent((QEvent::Type)(type_), (QInputDevice*)m_dev, QFlags<Qt::KeyboardModifier>(modifiers));
}
void* QInputEvent_device(void* self) {
	auto&& __r = ((QInputEvent*)self)->device();
	return (void*)(__r);
}
int QInputEvent_deviceType(void* self) {
	auto&& __r = ((QInputEvent*)self)->deviceType();
	return (int)(__r);
}
int QInputEvent_modifiers(void* self) {
	auto&& __r = ((QInputEvent*)self)->modifiers();
	return (int)(__r);
}
void QInputEvent_setModifiers(void* self, int modifiers) {
	((QInputEvent*)self)->setModifiers(QFlags<Qt::KeyboardModifier>(modifiers));
}
unsigned long long QInputEvent_timestamp(void* self) {
	auto&& __r = ((QInputEvent*)self)->timestamp();
	return __r;
}
void QInputEvent_setTimestamp(void* self, unsigned long long timestamp) {
	((QInputEvent*)self)->setTimestamp(timestamp);
}
int QInputEvent_registerEventType() {
	auto&& __r = QInputEvent::registerEventType();
	return __r;
}
int QInputEvent_registerEventType2(int hint) {
	auto&& __r = QInputEvent::registerEventType(hint);
	return __r;
}
void QInputEvent_delete(void* self) { delete (QInputEvent*)self; }

}

#include <QtGui/qevent.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qpoint.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qmimedata.h>
#include <QtCore/qobject.h>
#include "gen_qdropevent.h"

extern "C" {

void* QDropEvent_base(void* self) { return (void*)static_cast<QEvent*>((QDropEvent*)self); }
void* QDropEvent_new(void* pos, int actions, void* data, int buttons, int modifiers) {
	return new QDropEvent(*(QPointF*)pos, QFlags<Qt::DropAction>(actions), (QMimeData*)data, QFlags<Qt::MouseButton>(buttons), QFlags<Qt::KeyboardModifier>(modifiers));
}
void* QDropEvent_new2(void* pos, int actions, void* data, int buttons, int modifiers, int type_) {
	return new QDropEvent(*(QPointF*)pos, QFlags<Qt::DropAction>(actions), (QMimeData*)data, QFlags<Qt::MouseButton>(buttons), QFlags<Qt::KeyboardModifier>(modifiers), (QEvent::Type)(type_));
}
void* QDropEvent_pos(void* self) {
	auto&& __r = ((QDropEvent*)self)->pos();
	return new QPoint(__r);
}
void* QDropEvent_posF(void* self) {
	auto&& __r = ((QDropEvent*)self)->posF();
	return new QPointF(__r);
}
int QDropEvent_mouseButtons(void* self) {
	auto&& __r = ((QDropEvent*)self)->mouseButtons();
	return (int)(__r);
}
int QDropEvent_keyboardModifiers(void* self) {
	auto&& __r = ((QDropEvent*)self)->keyboardModifiers();
	return (int)(__r);
}
void* QDropEvent_position(void* self) {
	auto&& __r = ((QDropEvent*)self)->position();
	return new QPointF(__r);
}
int QDropEvent_buttons(void* self) {
	auto&& __r = ((QDropEvent*)self)->buttons();
	return (int)(__r);
}
int QDropEvent_modifiers(void* self) {
	auto&& __r = ((QDropEvent*)self)->modifiers();
	return (int)(__r);
}
int QDropEvent_possibleActions(void* self) {
	auto&& __r = ((QDropEvent*)self)->possibleActions();
	return (int)(__r);
}
int QDropEvent_proposedAction(void* self) {
	auto&& __r = ((QDropEvent*)self)->proposedAction();
	return (int)(__r);
}
void QDropEvent_acceptProposedAction(void* self) {
	((QDropEvent*)self)->acceptProposedAction();
}
int QDropEvent_dropAction(void* self) {
	auto&& __r = ((QDropEvent*)self)->dropAction();
	return (int)(__r);
}
void QDropEvent_setDropAction(void* self, int action) {
	((QDropEvent*)self)->setDropAction((Qt::DropAction)(action));
}
void* QDropEvent_source(void* self) {
	auto&& __r = ((QDropEvent*)self)->source();
	return (void*)(__r);
}
void* QDropEvent_mimeData(void* self) {
	auto&& __r = ((QDropEvent*)self)->mimeData();
	return (void*)(__r);
}
int QDropEvent_registerEventType() {
	auto&& __r = QDropEvent::registerEventType();
	return __r;
}
int QDropEvent_registerEventType2(int hint) {
	auto&& __r = QDropEvent::registerEventType(hint);
	return __r;
}
void QDropEvent_delete(void* self) { delete (QDropEvent*)self; }

}

#include <QtGui/qtextdocument.h>
#include <QtCore/qobject.h>
#include <QtCore/qstring.h>
#include <QtCore/qobjectdefs.h>
#include <QtGui/qabstracttextdocumentlayout.h>
#include <QtCore/qchar.h>
#include <QtGui/qtextcursor.h>
#include <QtCore/qregularexpression.h>
#include <QtGui/qtextobject.h>
#include <QtGui/qtextformat.h>
#include <QtCore/qsize.h>
#include <QtGui/qfont.h>
#include <QtGui/qpagedpaintdevice.h>
#include <QtCore/qvariant.h>
#include <QtCore/qurl.h>
#include <QtGui/qpainter.h>
#include <QtCore/qrect.h>
#include <QtGui/qtextoption.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qtextdocument.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QTextDocument__Go : public QTextDocument {
public:
	long long __goId;
	unsigned char __goOn[11] = {};
	template<typename... A> QTextDocument__Go(long long id, A&&... a) : QTextDocument(std::forward<A>(a)...), __goId(id) {}
	using QTextDocument::sender;
	using QTextDocument::senderSignalIndex;
	using QTextDocument::receivers;
	using QTextDocument::isSignalConnected;
	QTextObject* cppgoBase_createObject(const QTextFormat& __a0) { return QTextDocument::createObject(__a0); }
	QVariant cppgoBase_loadResource(int __a0, const QUrl& __a1) { return QTextDocument::loadResource(__a0, __a1); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QTextDocument::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QTextDocument::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QTextDocument::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QTextDocument::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QTextDocument::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QTextDocument::metaObject();
	}
	void clear() override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, -778567940572639070LL, &__h); if (__h) return; }
		QTextDocument::clear();
	}
	QTextObject* createObject(const QTextFormat& f) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 7258893770181894642LL, (void*)&f, &__h); if (__h) return (QTextObject*)__r; }
		return QTextDocument::createObject(f);
	}
	QVariant loadResource(int type_, const QUrl& name) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_voidp_int_voidp) { void* __r = ((void*(*)(long long, long long, int, void*, int*))cppgo_virt_voidp_int_voidp)(__goId, -519616681208050567LL, type_, (void*)&name, &__h); if (__h) return *(QVariant*)__r; }
		return QTextDocument::loadResource(type_, name);
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QTextDocument::event(event);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QTextDocument::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QTextDocument::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QTextDocument::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QTextDocument::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QTextDocument::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QTextDocument::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QTextDocument;
static std::mutex __connsMu_QTextDocument;

extern "C" {

void* QTextDocument_base(void* self) { return (void*)static_cast<QObject*>((QTextDocument*)self); }
void QTextDocument__set_override(void* self, int idx) { ((QTextDocument__Go*)self)->__goOn[idx] = 1; }
void* QTextDocument_fromQObject(void* o) { return (void*)qobject_cast<QTextDocument*>((QObject*)o); }
void* QTextDocument_new(long long __goid) {
	return new QTextDocument__Go(__goid);
}
void* QTextDocument_new2(long long __goid, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	return new QTextDocument__Go(__goid, text_qstr);
}
void* QTextDocument_new3(long long __goid, void* parent) {
	return new QTextDocument__Go(__goid, (QObject*)parent);
}
void* QTextDocument_new4(long long __goid, const char* text, long long text_len, void* parent) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	return new QTextDocument__Go(__goid, text_qstr, (QObject*)parent);
}
char* QTextDocument_tr(const char* s, long long* __out_len) {
	auto&& __r = QTextDocument::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QTextDocument_clone(void* self) {
	auto&& __r = ((QTextDocument*)self)->clone();
	return (void*)(__r);
}
int QTextDocument_isEmpty(void* self) {
	auto&& __r = ((QTextDocument*)self)->isEmpty();
	return __r;
}
void QTextDocument_clear(void* self) {
	((QTextDocument*)self)->clear();
}
void QTextDocument_setUndoRedoEnabled(void* self, int enable) {
	((QTextDocument*)self)->setUndoRedoEnabled(enable);
}
int QTextDocument_isUndoRedoEnabled(void* self) {
	auto&& __r = ((QTextDocument*)self)->isUndoRedoEnabled();
	return __r;
}
int QTextDocument_isUndoAvailable(void* self) {
	auto&& __r = ((QTextDocument*)self)->isUndoAvailable();
	return __r;
}
int QTextDocument_isRedoAvailable(void* self) {
	auto&& __r = ((QTextDocument*)self)->isRedoAvailable();
	return __r;
}
int QTextDocument_availableUndoSteps(void* self) {
	auto&& __r = ((QTextDocument*)self)->availableUndoSteps();
	return __r;
}
int QTextDocument_availableRedoSteps(void* self) {
	auto&& __r = ((QTextDocument*)self)->availableRedoSteps();
	return __r;
}
int QTextDocument_revision(void* self) {
	auto&& __r = ((QTextDocument*)self)->revision();
	return __r;
}
void QTextDocument_setDocumentLayout(void* self, void* layout) {
	((QTextDocument*)self)->setDocumentLayout((QAbstractTextDocumentLayout*)layout);
}
void* QTextDocument_documentLayout(void* self) {
	auto&& __r = ((QTextDocument*)self)->documentLayout();
	return (void*)(__r);
}
void QTextDocument_setMetaInformation(void* self, int info, const char* arg2, long long arg2_len) {
	QString arg2_qstr = QString::fromUtf8(arg2, arg2_len);
	((QTextDocument*)self)->setMetaInformation((QTextDocument::MetaInformation)(info), arg2_qstr);
}
char* QTextDocument_metaInformation(void* self, int info, long long* __out_len) {
	auto&& __r = ((QTextDocument*)self)->metaInformation((QTextDocument::MetaInformation)(info));
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QTextDocument_toHtml(void* self, long long* __out_len) {
	auto&& __r = ((QTextDocument*)self)->toHtml();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QTextDocument_setHtml(void* self, const char* html, long long html_len) {
	QString html_qstr = QString::fromUtf8(html, html_len);
	((QTextDocument*)self)->setHtml(html_qstr);
}
char* QTextDocument_toMarkdown(void* self, long long* __out_len) {
	auto&& __r = ((QTextDocument*)self)->toMarkdown();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QTextDocument_setMarkdown(void* self, const char* markdown, long long markdown_len) {
	QString markdown_qstr = QString::fromUtf8(markdown, markdown_len);
	((QTextDocument*)self)->setMarkdown(markdown_qstr);
}
char* QTextDocument_toRawText(void* self, long long* __out_len) {
	auto&& __r = ((QTextDocument*)self)->toRawText();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QTextDocument_toPlainText(void* self, long long* __out_len) {
	auto&& __r = ((QTextDocument*)self)->toPlainText();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QTextDocument_setPlainText(void* self, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	((QTextDocument*)self)->setPlainText(text_qstr);
}
void* QTextDocument_characterAt(void* self, int pos) {
	auto&& __r = ((QTextDocument*)self)->characterAt(pos);
	return new QChar(__r);
}
void* QTextDocument_find(void* self, const char* subString, long long subString_len) {
	QString subString_qstr = QString::fromUtf8(subString, subString_len);
	auto&& __r = ((QTextDocument*)self)->find(subString_qstr);
	return new QTextCursor(__r);
}
void* QTextDocument_find2(void* self, const char* subString, long long subString_len, void* cursor) {
	QString subString_qstr = QString::fromUtf8(subString, subString_len);
	auto&& __r = ((QTextDocument*)self)->find(subString_qstr, *(QTextCursor*)cursor);
	return new QTextCursor(__r);
}
void* QTextDocument_find3(void* self, void* expr) {
	auto&& __r = ((QTextDocument*)self)->find(*(QRegularExpression*)expr);
	return new QTextCursor(__r);
}
void* QTextDocument_find4(void* self, void* expr, void* cursor) {
	auto&& __r = ((QTextDocument*)self)->find(*(QRegularExpression*)expr, *(QTextCursor*)cursor);
	return new QTextCursor(__r);
}
void* QTextDocument_frameAt(void* self, int pos) {
	auto&& __r = ((QTextDocument*)self)->frameAt(pos);
	return (void*)(__r);
}
void* QTextDocument_rootFrame(void* self) {
	auto&& __r = ((QTextDocument*)self)->rootFrame();
	return (void*)(__r);
}
void* QTextDocument_object(void* self, int objectIndex) {
	auto&& __r = ((QTextDocument*)self)->object(objectIndex);
	return (void*)(__r);
}
void* QTextDocument_objectForFormat(void* self, void* arg1) {
	auto&& __r = ((QTextDocument*)self)->objectForFormat(*(QTextFormat*)arg1);
	return (void*)(__r);
}
void* QTextDocument_findBlock(void* self, int pos) {
	auto&& __r = ((QTextDocument*)self)->findBlock(pos);
	return new QTextBlock(__r);
}
void* QTextDocument_findBlockByNumber(void* self, int blockNumber) {
	auto&& __r = ((QTextDocument*)self)->findBlockByNumber(blockNumber);
	return new QTextBlock(__r);
}
void* QTextDocument_findBlockByLineNumber(void* self, int blockNumber) {
	auto&& __r = ((QTextDocument*)self)->findBlockByLineNumber(blockNumber);
	return new QTextBlock(__r);
}
void* QTextDocument_begin(void* self) {
	auto&& __r = ((QTextDocument*)self)->begin();
	return new QTextBlock(__r);
}
void* QTextDocument_end(void* self) {
	auto&& __r = ((QTextDocument*)self)->end();
	return new QTextBlock(__r);
}
void* QTextDocument_firstBlock(void* self) {
	auto&& __r = ((QTextDocument*)self)->firstBlock();
	return new QTextBlock(__r);
}
void* QTextDocument_lastBlock(void* self) {
	auto&& __r = ((QTextDocument*)self)->lastBlock();
	return new QTextBlock(__r);
}
void QTextDocument_setPageSize(void* self, void* size) {
	((QTextDocument*)self)->setPageSize(*(QSizeF*)size);
}
void* QTextDocument_pageSize(void* self) {
	auto&& __r = ((QTextDocument*)self)->pageSize();
	return new QSizeF(__r);
}
void QTextDocument_setDefaultFont(void* self, void* font) {
	((QTextDocument*)self)->setDefaultFont(*(QFont*)font);
}
void* QTextDocument_defaultFont(void* self) {
	auto&& __r = ((QTextDocument*)self)->defaultFont();
	return new QFont(__r);
}
void QTextDocument_setSuperScriptBaseline(void* self, double baseline) {
	((QTextDocument*)self)->setSuperScriptBaseline(baseline);
}
double QTextDocument_superScriptBaseline(void* self) {
	auto&& __r = ((QTextDocument*)self)->superScriptBaseline();
	return __r;
}
void QTextDocument_setSubScriptBaseline(void* self, double baseline) {
	((QTextDocument*)self)->setSubScriptBaseline(baseline);
}
double QTextDocument_subScriptBaseline(void* self) {
	auto&& __r = ((QTextDocument*)self)->subScriptBaseline();
	return __r;
}
void QTextDocument_setBaselineOffset(void* self, double baseline) {
	((QTextDocument*)self)->setBaselineOffset(baseline);
}
double QTextDocument_baselineOffset(void* self) {
	auto&& __r = ((QTextDocument*)self)->baselineOffset();
	return __r;
}
int QTextDocument_pageCount(void* self) {
	auto&& __r = ((QTextDocument*)self)->pageCount();
	return __r;
}
int QTextDocument_isModified(void* self) {
	auto&& __r = ((QTextDocument*)self)->isModified();
	return __r;
}
void QTextDocument_print(void* self, void* printer) {
	((QTextDocument*)self)->print((QPagedPaintDevice*)printer);
}
void* QTextDocument_resource(void* self, int type_, void* name) {
	auto&& __r = ((QTextDocument*)self)->resource(type_, *(QUrl*)name);
	return new QVariant(__r);
}
void QTextDocument_addResource(void* self, int type_, void* name, void* resource) {
	((QTextDocument*)self)->addResource(type_, *(QUrl*)name, *(QVariant*)resource);
}
void** QTextDocument_allFormats(void* self, long long* __out_len) {
	auto&& __r = ((QTextDocument*)self)->allFormats();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QTextFormat(__e);
	return __arr;
}
void QTextDocument_markContentsDirty(void* self, int from, int length) {
	((QTextDocument*)self)->markContentsDirty(from, length);
}
void QTextDocument_setUseDesignMetrics(void* self, int b) {
	((QTextDocument*)self)->setUseDesignMetrics(b);
}
int QTextDocument_useDesignMetrics(void* self) {
	auto&& __r = ((QTextDocument*)self)->useDesignMetrics();
	return __r;
}
void QTextDocument_setLayoutEnabled(void* self, int b) {
	((QTextDocument*)self)->setLayoutEnabled(b);
}
int QTextDocument_isLayoutEnabled(void* self) {
	auto&& __r = ((QTextDocument*)self)->isLayoutEnabled();
	return __r;
}
void QTextDocument_drawContents(void* self, void* painter) {
	((QTextDocument*)self)->drawContents((QPainter*)painter);
}
void QTextDocument_setTextWidth(void* self, double width) {
	((QTextDocument*)self)->setTextWidth(width);
}
double QTextDocument_textWidth(void* self) {
	auto&& __r = ((QTextDocument*)self)->textWidth();
	return __r;
}
double QTextDocument_idealWidth(void* self) {
	auto&& __r = ((QTextDocument*)self)->idealWidth();
	return __r;
}
double QTextDocument_indentWidth(void* self) {
	auto&& __r = ((QTextDocument*)self)->indentWidth();
	return __r;
}
void QTextDocument_setIndentWidth(void* self, double width) {
	((QTextDocument*)self)->setIndentWidth(width);
}
double QTextDocument_documentMargin(void* self) {
	auto&& __r = ((QTextDocument*)self)->documentMargin();
	return __r;
}
void QTextDocument_setDocumentMargin(void* self, double margin) {
	((QTextDocument*)self)->setDocumentMargin(margin);
}
void QTextDocument_adjustSize(void* self) {
	((QTextDocument*)self)->adjustSize();
}
void* QTextDocument_size(void* self) {
	auto&& __r = ((QTextDocument*)self)->size();
	return new QSizeF(__r);
}
int QTextDocument_blockCount(void* self) {
	auto&& __r = ((QTextDocument*)self)->blockCount();
	return __r;
}
int QTextDocument_lineCount(void* self) {
	auto&& __r = ((QTextDocument*)self)->lineCount();
	return __r;
}
int QTextDocument_characterCount(void* self) {
	auto&& __r = ((QTextDocument*)self)->characterCount();
	return __r;
}
void QTextDocument_setDefaultStyleSheet(void* self, const char* sheet, long long sheet_len) {
	QString sheet_qstr = QString::fromUtf8(sheet, sheet_len);
	((QTextDocument*)self)->setDefaultStyleSheet(sheet_qstr);
}
char* QTextDocument_defaultStyleSheet(void* self, long long* __out_len) {
	auto&& __r = ((QTextDocument*)self)->defaultStyleSheet();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QTextDocument_undo(void* self, void* cursor) {
	((QTextDocument*)self)->undo((QTextCursor*)cursor);
}
void QTextDocument_redo(void* self, void* cursor) {
	((QTextDocument*)self)->redo((QTextCursor*)cursor);
}
void QTextDocument_clearUndoRedoStacks(void* self) {
	((QTextDocument*)self)->clearUndoRedoStacks();
}
int QTextDocument_maximumBlockCount(void* self) {
	auto&& __r = ((QTextDocument*)self)->maximumBlockCount();
	return __r;
}
void QTextDocument_setMaximumBlockCount(void* self, int maximum) {
	((QTextDocument*)self)->setMaximumBlockCount(maximum);
}
void* QTextDocument_defaultTextOption(void* self) {
	auto&& __r = ((QTextDocument*)self)->defaultTextOption();
	return new QTextOption(__r);
}
void QTextDocument_setDefaultTextOption(void* self, void* option) {
	((QTextDocument*)self)->setDefaultTextOption(*(QTextOption*)option);
}
void* QTextDocument_baseUrl(void* self) {
	auto&& __r = ((QTextDocument*)self)->baseUrl();
	return new QUrl(__r);
}
void QTextDocument_setBaseUrl(void* self, void* url) {
	((QTextDocument*)self)->setBaseUrl(*(QUrl*)url);
}
int QTextDocument_defaultCursorMoveStyle(void* self) {
	auto&& __r = ((QTextDocument*)self)->defaultCursorMoveStyle();
	return (int)(__r);
}
void QTextDocument_setDefaultCursorMoveStyle(void* self, int style) {
	((QTextDocument*)self)->setDefaultCursorMoveStyle((Qt::CursorMoveStyle)(style));
}
void QTextDocument_contentsChange(void* self, int from, int charsRemoved, int charsAdded) {
	((QTextDocument*)self)->contentsChange(from, charsRemoved, charsAdded);
}
void QTextDocument_contentsChanged(void* self) {
	((QTextDocument*)self)->contentsChanged();
}
void QTextDocument_undoAvailable(void* self, int arg1) {
	((QTextDocument*)self)->undoAvailable(arg1);
}
void QTextDocument_redoAvailable(void* self, int arg1) {
	((QTextDocument*)self)->redoAvailable(arg1);
}
void QTextDocument_undoCommandAdded(void* self) {
	((QTextDocument*)self)->undoCommandAdded();
}
void QTextDocument_modificationChanged(void* self, int m) {
	((QTextDocument*)self)->modificationChanged(m);
}
void QTextDocument_cursorPositionChanged(void* self, void* cursor) {
	((QTextDocument*)self)->cursorPositionChanged(*(QTextCursor*)cursor);
}
void QTextDocument_blockCountChanged(void* self, int newBlockCount) {
	((QTextDocument*)self)->blockCountChanged(newBlockCount);
}
void QTextDocument_baseUrlChanged(void* self, void* url) {
	((QTextDocument*)self)->baseUrlChanged(*(QUrl*)url);
}
void QTextDocument_documentLayoutChanged(void* self) {
	((QTextDocument*)self)->documentLayoutChanged();
}
void QTextDocument_undo2(void* self) {
	((QTextDocument*)self)->undo();
}
void QTextDocument_redo2(void* self) {
	((QTextDocument*)self)->redo();
}
void QTextDocument_appendUndoItem(void* self, void* arg1) {
	((QTextDocument*)self)->appendUndoItem((QAbstractUndoItem*)arg1);
}
void QTextDocument_setModified(void* self) {
	((QTextDocument*)self)->setModified();
}
void* QTextDocument_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QTextDocument::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QTextDocument_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QTextDocument::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QTextDocument_sender(void* self) {
	auto&& __r = ((QTextDocument__Go*)self)->sender();
	return (void*)(__r);
}
int QTextDocument_senderSignalIndex(void* self) {
	auto&& __r = ((QTextDocument__Go*)self)->senderSignalIndex();
	return __r;
}
int QTextDocument_receivers(void* self, const char* signal) {
	auto&& __r = ((QTextDocument__Go*)self)->receivers(signal);
	return __r;
}
int QTextDocument_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QTextDocument__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QTextDocument_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QTextDocument::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QTextDocument_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QTextDocument::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QTextDocument_clone2(void* self, void* parent) {
	auto&& __r = ((QTextDocument*)self)->clone((QObject*)parent);
	return (void*)(__r);
}
char* QTextDocument_toMarkdown2(void* self, int features, long long* __out_len) {
	auto&& __r = ((QTextDocument*)self)->toMarkdown(QFlags<QTextDocument::MarkdownFeature>(features));
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QTextDocument_setMarkdown2(void* self, const char* markdown, long long markdown_len, int features) {
	QString markdown_qstr = QString::fromUtf8(markdown, markdown_len);
	((QTextDocument*)self)->setMarkdown(markdown_qstr, QFlags<QTextDocument::MarkdownFeature>(features));
}
void* QTextDocument_find5(void* self, const char* subString, long long subString_len, int from) {
	QString subString_qstr = QString::fromUtf8(subString, subString_len);
	auto&& __r = ((QTextDocument*)self)->find(subString_qstr, from);
	return new QTextCursor(__r);
}
void* QTextDocument_find6(void* self, const char* subString, long long subString_len, int from, int options) {
	QString subString_qstr = QString::fromUtf8(subString, subString_len);
	auto&& __r = ((QTextDocument*)self)->find(subString_qstr, from, QFlags<QTextDocument::FindFlag>(options));
	return new QTextCursor(__r);
}
void* QTextDocument_find7(void* self, const char* subString, long long subString_len, void* cursor, int options) {
	QString subString_qstr = QString::fromUtf8(subString, subString_len);
	auto&& __r = ((QTextDocument*)self)->find(subString_qstr, *(QTextCursor*)cursor, QFlags<QTextDocument::FindFlag>(options));
	return new QTextCursor(__r);
}
void* QTextDocument_find8(void* self, void* expr, int from) {
	auto&& __r = ((QTextDocument*)self)->find(*(QRegularExpression*)expr, from);
	return new QTextCursor(__r);
}
void* QTextDocument_find9(void* self, void* expr, int from, int options) {
	auto&& __r = ((QTextDocument*)self)->find(*(QRegularExpression*)expr, from, QFlags<QTextDocument::FindFlag>(options));
	return new QTextCursor(__r);
}
void* QTextDocument_find10(void* self, void* expr, void* cursor, int options) {
	auto&& __r = ((QTextDocument*)self)->find(*(QRegularExpression*)expr, *(QTextCursor*)cursor, QFlags<QTextDocument::FindFlag>(options));
	return new QTextCursor(__r);
}
void QTextDocument_drawContents2(void* self, void* painter, void* rect) {
	((QTextDocument*)self)->drawContents((QPainter*)painter, *(QRectF*)rect);
}
void QTextDocument_clearUndoRedoStacks2(void* self, int historyToClear) {
	((QTextDocument*)self)->clearUndoRedoStacks((QTextDocument::Stacks)(historyToClear));
}
void QTextDocument_setModified2(void* self, int m) {
	((QTextDocument*)self)->setModified(m);
}
void* QTextDocument_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QTextDocument::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QTextDocument_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QTextDocument::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void QTextDocument_clear_default(void* self) {
	((QTextDocument*)self)->QTextDocument::clear();
}
void* QTextDocument_metaObject_default(void* self) {
	auto&& __r = ((QTextDocument*)self)->QTextDocument::metaObject();
	return (void*)(__r);
}
void* QTextDocument_createObject_default(void* self, void* f) {
	auto&& __r = ((QTextDocument__Go*)self)->cppgoBase_createObject(*(QTextFormat*)f);
	return (void*)(__r);
}
void* QTextDocument_loadResource_default(void* self, int type_, void* name) {
	auto&& __r = ((QTextDocument__Go*)self)->cppgoBase_loadResource(type_, *(QUrl*)name);
	return new QVariant(__r);
}
int QTextDocument_event_default(void* self, void* event) {
	auto&& __r = ((QTextDocument*)self)->QTextDocument::event((QEvent*)event);
	return __r;
}
int QTextDocument_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QTextDocument*)self)->QTextDocument::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QTextDocument_timerEvent_default(void* self, void* event) {
	((QTextDocument__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QTextDocument_childEvent_default(void* self, void* event) {
	((QTextDocument__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QTextDocument_customEvent_default(void* self, void* event) {
	((QTextDocument__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QTextDocument_connectNotify_default(void* self, void* signal) {
	((QTextDocument__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QTextDocument_disconnectNotify_default(void* self, void* signal) {
	((QTextDocument__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QTextDocument_connect_contentsChange(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTextDocument*)self, &QTextDocument::contentsChange, (QTextDocument*)self, [__id](int from, int charsRemoved, int charsAdded){
		if (cppgo_sig_void_int_int_int) ((void(*)(long long, int, int, int))cppgo_sig_void_int_int_int)(__id, from, charsRemoved, charsAdded);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTextDocument);
	__conns_QTextDocument[__id] = __c;
}
void QTextDocument_connect_contentsChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTextDocument*)self, &QTextDocument::contentsChanged, (QTextDocument*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTextDocument);
	__conns_QTextDocument[__id] = __c;
}
void QTextDocument_connect_undoAvailable(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTextDocument*)self, &QTextDocument::undoAvailable, (QTextDocument*)self, [__id](bool arg1){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTextDocument);
	__conns_QTextDocument[__id] = __c;
}
void QTextDocument_connect_redoAvailable(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTextDocument*)self, &QTextDocument::redoAvailable, (QTextDocument*)self, [__id](bool arg1){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTextDocument);
	__conns_QTextDocument[__id] = __c;
}
void QTextDocument_connect_undoCommandAdded(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTextDocument*)self, &QTextDocument::undoCommandAdded, (QTextDocument*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTextDocument);
	__conns_QTextDocument[__id] = __c;
}
void QTextDocument_connect_modificationChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTextDocument*)self, &QTextDocument::modificationChanged, (QTextDocument*)self, [__id](bool m){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, m);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTextDocument);
	__conns_QTextDocument[__id] = __c;
}
void QTextDocument_connect_cursorPositionChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTextDocument*)self, &QTextDocument::cursorPositionChanged, (QTextDocument*)self, [__id](const QTextCursor& cursor){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&cursor);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTextDocument);
	__conns_QTextDocument[__id] = __c;
}
void QTextDocument_connect_blockCountChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTextDocument*)self, &QTextDocument::blockCountChanged, (QTextDocument*)self, [__id](int newBlockCount){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, newBlockCount);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTextDocument);
	__conns_QTextDocument[__id] = __c;
}
void QTextDocument_connect_baseUrlChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTextDocument*)self, &QTextDocument::baseUrlChanged, (QTextDocument*)self, [__id](const QUrl& url){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&url);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTextDocument);
	__conns_QTextDocument[__id] = __c;
}
void QTextDocument_connect_documentLayoutChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTextDocument*)self, &QTextDocument::documentLayoutChanged, (QTextDocument*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTextDocument);
	__conns_QTextDocument[__id] = __c;
}
void QTextDocument_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTextDocument*)self, &QTextDocument::destroyed, (QTextDocument*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTextDocument);
	__conns_QTextDocument[__id] = __c;
}
void QTextDocument_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTextDocument*)self, &QTextDocument::destroyed, (QTextDocument*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTextDocument);
	__conns_QTextDocument[__id] = __c;
}
void QTextDocument_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTextDocument*)self, &QTextDocument::objectNameChanged, (QTextDocument*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTextDocument);
	__conns_QTextDocument[__id] = __c;
}
void QTextDocument__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QTextDocument);
		auto __it = __conns_QTextDocument.find(id);
		if (__it == __conns_QTextDocument.end()) return;
		__c = __it->second;
		__conns_QTextDocument.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QTextDocument__Go_new(long long id, void* parent) { return new QTextDocument__Go(id, (QObject*)parent); }
void QTextDocument_delete(void* self) { delete (QTextDocument*)self; }

}

#include <QtGui/qtextlayout.h>
#include <QtCore/qrect.h>
#include <QtCore/qpoint.h>
#include <QtGui/qpainter.h>
#include <QtGui/qglyphrun.h>
#include "gen_qtextline.h"

extern "C" {

void* QTextLine_new() {
	return new QTextLine();
}
int QTextLine_isValid(void* self) {
	auto&& __r = ((QTextLine*)self)->isValid();
	return __r;
}
void* QTextLine_rect(void* self) {
	auto&& __r = ((QTextLine*)self)->rect();
	return new QRectF(__r);
}
double QTextLine_x(void* self) {
	auto&& __r = ((QTextLine*)self)->x();
	return __r;
}
double QTextLine_y(void* self) {
	auto&& __r = ((QTextLine*)self)->y();
	return __r;
}
double QTextLine_width(void* self) {
	auto&& __r = ((QTextLine*)self)->width();
	return __r;
}
double QTextLine_ascent(void* self) {
	auto&& __r = ((QTextLine*)self)->ascent();
	return __r;
}
double QTextLine_descent(void* self) {
	auto&& __r = ((QTextLine*)self)->descent();
	return __r;
}
double QTextLine_height(void* self) {
	auto&& __r = ((QTextLine*)self)->height();
	return __r;
}
double QTextLine_leading(void* self) {
	auto&& __r = ((QTextLine*)self)->leading();
	return __r;
}
void QTextLine_setLeadingIncluded(void* self, int included) {
	((QTextLine*)self)->setLeadingIncluded(included);
}
int QTextLine_leadingIncluded(void* self) {
	auto&& __r = ((QTextLine*)self)->leadingIncluded();
	return __r;
}
double QTextLine_naturalTextWidth(void* self) {
	auto&& __r = ((QTextLine*)self)->naturalTextWidth();
	return __r;
}
double QTextLine_horizontalAdvance(void* self) {
	auto&& __r = ((QTextLine*)self)->horizontalAdvance();
	return __r;
}
void* QTextLine_naturalTextRect(void* self) {
	auto&& __r = ((QTextLine*)self)->naturalTextRect();
	return new QRectF(__r);
}
double QTextLine_cursorToX(void* self, int* cursorPos) {
	int cursorPos_local{};
	auto&& __r = ((QTextLine*)self)->cursorToX(&cursorPos_local);
	*cursorPos = cursorPos_local;
	return __r;
}
double QTextLine_cursorToX2(void* self, int cursorPos) {
	auto&& __r = ((QTextLine*)self)->cursorToX(cursorPos);
	return __r;
}
int QTextLine_xToCursor(void* self, double x) {
	auto&& __r = ((QTextLine*)self)->xToCursor(x);
	return __r;
}
void QTextLine_setLineWidth(void* self, double width) {
	((QTextLine*)self)->setLineWidth(width);
}
void QTextLine_setNumColumns(void* self, int columns) {
	((QTextLine*)self)->setNumColumns(columns);
}
void QTextLine_setNumColumns2(void* self, int columns, double alignmentWidth) {
	((QTextLine*)self)->setNumColumns(columns, alignmentWidth);
}
void QTextLine_setPosition(void* self, void* pos) {
	((QTextLine*)self)->setPosition(*(QPointF*)pos);
}
void* QTextLine_position(void* self) {
	auto&& __r = ((QTextLine*)self)->position();
	return new QPointF(__r);
}
int QTextLine_textStart(void* self) {
	auto&& __r = ((QTextLine*)self)->textStart();
	return __r;
}
int QTextLine_textLength(void* self) {
	auto&& __r = ((QTextLine*)self)->textLength();
	return __r;
}
int QTextLine_lineNumber(void* self) {
	auto&& __r = ((QTextLine*)self)->lineNumber();
	return __r;
}
void QTextLine_draw(void* self, void* painter, void* position) {
	((QTextLine*)self)->draw((QPainter*)painter, *(QPointF*)position);
}
void** QTextLine_glyphRuns(void* self, int from, int length, int flags, long long* __out_len) {
	auto&& __r = ((QTextLine*)self)->glyphRuns(from, length, QFlags<QTextLayout::GlyphRunRetrievalFlag>(flags));
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QGlyphRun(__e);
	return __arr;
}
void** QTextLine_glyphRuns2(void* self, long long* __out_len) {
	auto&& __r = ((QTextLine*)self)->glyphRuns();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QGlyphRun(__e);
	return __arr;
}
double QTextLine_cursorToX3(void* self, int* cursorPos, int edge) {
	int cursorPos_local{};
	auto&& __r = ((QTextLine*)self)->cursorToX(&cursorPos_local, (QTextLine::Edge)(edge));
	*cursorPos = cursorPos_local;
	return __r;
}
double QTextLine_cursorToX4(void* self, int cursorPos, int edge) {
	auto&& __r = ((QTextLine*)self)->cursorToX(cursorPos, (QTextLine::Edge)(edge));
	return __r;
}
int QTextLine_xToCursor2(void* self, double x, int arg2) {
	auto&& __r = ((QTextLine*)self)->xToCursor(x, (QTextLine::CursorPosition)(arg2));
	return __r;
}
void** QTextLine_glyphRuns3(void* self, int from, long long* __out_len) {
	auto&& __r = ((QTextLine*)self)->glyphRuns(from);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QGlyphRun(__e);
	return __arr;
}
void** QTextLine_glyphRuns4(void* self, int from, int length, long long* __out_len) {
	auto&& __r = ((QTextLine*)self)->glyphRuns(from, length);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QGlyphRun(__e);
	return __arr;
}
void QTextLine_delete(void* self) { delete (QTextLine*)self; }

}

#include <QtGui/qaccessible.h>
#include <QtCore/qobject.h>
#include <QtGui/qaccessible_base.h>
#include "gen_qaccessibletextselectionevent.h"

extern "C" {

void* QAccessibleTextSelectionEvent_base(void* self) { return (void*)static_cast<QAccessibleTextCursorEvent*>((QAccessibleTextSelectionEvent*)self); }
void QAccessibleTextSelectionEvent_chain(void* self, void** out) {
	QAccessibleTextSelectionEvent* __p = (QAccessibleTextSelectionEvent*)self;
	QAccessibleTextCursorEvent* __b0 = static_cast<QAccessibleTextCursorEvent*>(__p); out[0] = (void*)__b0;
	QAccessibleEvent* __b1 = static_cast<QAccessibleEvent*>(__b0); out[1] = (void*)__b1;
}
void* QAccessibleTextSelectionEvent_new(void* obj, int start, int end) {
	return new QAccessibleTextSelectionEvent((QObject*)obj, start, end);
}
void* QAccessibleTextSelectionEvent_new2(void* iface, int start, int end) {
	return new QAccessibleTextSelectionEvent((QAccessibleInterface*)iface, start, end);
}
void QAccessibleTextSelectionEvent_setSelection(void* self, int start, int end) {
	((QAccessibleTextSelectionEvent*)self)->setSelection(start, end);
}
int QAccessibleTextSelectionEvent_selectionStart(void* self) {
	auto&& __r = ((QAccessibleTextSelectionEvent*)self)->selectionStart();
	return __r;
}
int QAccessibleTextSelectionEvent_selectionEnd(void* self) {
	auto&& __r = ((QAccessibleTextSelectionEvent*)self)->selectionEnd();
	return __r;
}
void QAccessibleTextSelectionEvent_delete(void* self) { delete (QAccessibleTextSelectionEvent*)self; }

}

#include <QtGui/qdesktopservices.h>
#include <QtCore/qurl.h>
#include <QtCore/qstring.h>
#include <QtCore/qobject.h>
#include "gen_qdesktopservices.h"

extern "C" {

int QDesktopServices_openUrl(void* url) {
	auto&& __r = QDesktopServices::openUrl(*(QUrl*)url);
	return __r;
}
void QDesktopServices_setUrlHandler(const char* scheme, long long scheme_len, void* receiver, const char* method) {
	QString scheme_qstr = QString::fromUtf8(scheme, scheme_len);
	QDesktopServices::setUrlHandler(scheme_qstr, (QObject*)receiver, method);
}
void QDesktopServices_unsetUrlHandler(const char* scheme, long long scheme_len) {
	QString scheme_qstr = QString::fromUtf8(scheme, scheme_len);
	QDesktopServices::unsetUrlHandler(scheme_qstr);
}
void QDesktopServices_delete(void* self) { delete (QDesktopServices*)self; }

}

#include <QtGui/qpagelayout.h>
#include <QtGui/qpagesize.h>
#include <QtCore/qmargins.h>
#include <QtCore/qrect.h>
#include "gen_qpagelayout.h"

extern "C" {

void* QPageLayout_new() {
	return new QPageLayout();
}
void* QPageLayout_new2(void* pageSize, int orientation, void* margins) {
	return new QPageLayout(*(QPageSize*)pageSize, (QPageLayout::Orientation)(orientation), *(QMarginsF*)margins);
}
void* QPageLayout_new3(void* other) {
	return new QPageLayout(*(QPageLayout*)other);
}
void* QPageLayout_new4(void* pageSize, int orientation, void* margins, int units) {
	return new QPageLayout(*(QPageSize*)pageSize, (QPageLayout::Orientation)(orientation), *(QMarginsF*)margins, (QPageLayout::Unit)(units));
}
void* QPageLayout_new5(void* pageSize, int orientation, void* margins, int units, void* minMargins) {
	return new QPageLayout(*(QPageSize*)pageSize, (QPageLayout::Orientation)(orientation), *(QMarginsF*)margins, (QPageLayout::Unit)(units), *(QMarginsF*)minMargins);
}
void* QPageLayout_OperatorAssign(void* self, void* other) {
	auto&& __r = ((QPageLayout*)self)->operator=(*(QPageLayout*)other);
	return (void*)&(__r);
}
void QPageLayout_swap(void* self, void* other) {
	((QPageLayout*)self)->swap(*(QPageLayout*)other);
}
int QPageLayout_isEquivalentTo(void* self, void* other) {
	auto&& __r = ((QPageLayout*)self)->isEquivalentTo(*(QPageLayout*)other);
	return __r;
}
int QPageLayout_isValid(void* self) {
	auto&& __r = ((QPageLayout*)self)->isValid();
	return __r;
}
void QPageLayout_setMode(void* self, int mode) {
	((QPageLayout*)self)->setMode((QPageLayout::Mode)(mode));
}
int QPageLayout_mode(void* self) {
	auto&& __r = ((QPageLayout*)self)->mode();
	return (int)(__r);
}
void QPageLayout_setPageSize(void* self, void* pageSize) {
	((QPageLayout*)self)->setPageSize(*(QPageSize*)pageSize);
}
void* QPageLayout_pageSize(void* self) {
	auto&& __r = ((QPageLayout*)self)->pageSize();
	return new QPageSize(__r);
}
void QPageLayout_setOrientation(void* self, int orientation) {
	((QPageLayout*)self)->setOrientation((QPageLayout::Orientation)(orientation));
}
int QPageLayout_orientation(void* self) {
	auto&& __r = ((QPageLayout*)self)->orientation();
	return (int)(__r);
}
void QPageLayout_setUnits(void* self, int units) {
	((QPageLayout*)self)->setUnits((QPageLayout::Unit)(units));
}
int QPageLayout_units(void* self) {
	auto&& __r = ((QPageLayout*)self)->units();
	return (int)(__r);
}
int QPageLayout_setMargins(void* self, void* margins) {
	auto&& __r = ((QPageLayout*)self)->setMargins(*(QMarginsF*)margins);
	return __r;
}
int QPageLayout_setLeftMargin(void* self, double leftMargin) {
	auto&& __r = ((QPageLayout*)self)->setLeftMargin(leftMargin);
	return __r;
}
int QPageLayout_setRightMargin(void* self, double rightMargin) {
	auto&& __r = ((QPageLayout*)self)->setRightMargin(rightMargin);
	return __r;
}
int QPageLayout_setTopMargin(void* self, double topMargin) {
	auto&& __r = ((QPageLayout*)self)->setTopMargin(topMargin);
	return __r;
}
int QPageLayout_setBottomMargin(void* self, double bottomMargin) {
	auto&& __r = ((QPageLayout*)self)->setBottomMargin(bottomMargin);
	return __r;
}
void* QPageLayout_margins(void* self) {
	auto&& __r = ((QPageLayout*)self)->margins();
	return new QMarginsF(__r);
}
void* QPageLayout_margins2(void* self, int units) {
	auto&& __r = ((QPageLayout*)self)->margins((QPageLayout::Unit)(units));
	return new QMarginsF(__r);
}
void* QPageLayout_marginsPoints(void* self) {
	auto&& __r = ((QPageLayout*)self)->marginsPoints();
	return new QMargins(__r);
}
void* QPageLayout_marginsPixels(void* self, int resolution) {
	auto&& __r = ((QPageLayout*)self)->marginsPixels(resolution);
	return new QMargins(__r);
}
void QPageLayout_setMinimumMargins(void* self, void* minMargins) {
	((QPageLayout*)self)->setMinimumMargins(*(QMarginsF*)minMargins);
}
void* QPageLayout_minimumMargins(void* self) {
	auto&& __r = ((QPageLayout*)self)->minimumMargins();
	return new QMarginsF(__r);
}
void* QPageLayout_maximumMargins(void* self) {
	auto&& __r = ((QPageLayout*)self)->maximumMargins();
	return new QMarginsF(__r);
}
void* QPageLayout_fullRect(void* self) {
	auto&& __r = ((QPageLayout*)self)->fullRect();
	return new QRectF(__r);
}
void* QPageLayout_fullRect2(void* self, int units) {
	auto&& __r = ((QPageLayout*)self)->fullRect((QPageLayout::Unit)(units));
	return new QRectF(__r);
}
void* QPageLayout_fullRectPoints(void* self) {
	auto&& __r = ((QPageLayout*)self)->fullRectPoints();
	return new QRect(__r);
}
void* QPageLayout_fullRectPixels(void* self, int resolution) {
	auto&& __r = ((QPageLayout*)self)->fullRectPixels(resolution);
	return new QRect(__r);
}
void* QPageLayout_paintRect(void* self) {
	auto&& __r = ((QPageLayout*)self)->paintRect();
	return new QRectF(__r);
}
void* QPageLayout_paintRect2(void* self, int units) {
	auto&& __r = ((QPageLayout*)self)->paintRect((QPageLayout::Unit)(units));
	return new QRectF(__r);
}
void* QPageLayout_paintRectPoints(void* self) {
	auto&& __r = ((QPageLayout*)self)->paintRectPoints();
	return new QRect(__r);
}
void* QPageLayout_paintRectPixels(void* self, int resolution) {
	auto&& __r = ((QPageLayout*)self)->paintRectPixels(resolution);
	return new QRect(__r);
}
void QPageLayout_setPageSize2(void* self, void* pageSize, void* minMargins) {
	((QPageLayout*)self)->setPageSize(*(QPageSize*)pageSize, *(QMarginsF*)minMargins);
}
int QPageLayout_setMargins2(void* self, void* margins, int outOfBoundsPolicy) {
	auto&& __r = ((QPageLayout*)self)->setMargins(*(QMarginsF*)margins, (QPageLayout::OutOfBoundsPolicy)(outOfBoundsPolicy));
	return __r;
}
int QPageLayout_setLeftMargin2(void* self, double leftMargin, int outOfBoundsPolicy) {
	auto&& __r = ((QPageLayout*)self)->setLeftMargin(leftMargin, (QPageLayout::OutOfBoundsPolicy)(outOfBoundsPolicy));
	return __r;
}
int QPageLayout_setRightMargin2(void* self, double rightMargin, int outOfBoundsPolicy) {
	auto&& __r = ((QPageLayout*)self)->setRightMargin(rightMargin, (QPageLayout::OutOfBoundsPolicy)(outOfBoundsPolicy));
	return __r;
}
int QPageLayout_setTopMargin2(void* self, double topMargin, int outOfBoundsPolicy) {
	auto&& __r = ((QPageLayout*)self)->setTopMargin(topMargin, (QPageLayout::OutOfBoundsPolicy)(outOfBoundsPolicy));
	return __r;
}
int QPageLayout_setBottomMargin2(void* self, double bottomMargin, int outOfBoundsPolicy) {
	auto&& __r = ((QPageLayout*)self)->setBottomMargin(bottomMargin, (QPageLayout::OutOfBoundsPolicy)(outOfBoundsPolicy));
	return __r;
}
int QPageLayout_OperatorEqual(void* self, void* rhs) {
	auto&& __r = (*(QPageLayout*)self) == (*(QPageLayout*)rhs);
	return __r;
}
int QPageLayout_OperatorNotEqual(void* self, void* rhs) {
	auto&& __r = (*(QPageLayout*)self) != (*(QPageLayout*)rhs);
	return __r;
}
void QPageLayout_delete(void* self) { delete (QPageLayout*)self; }

}

#include <QtGui/qshortcut.h>
#include <QtCore/qobject.h>
#include <QtGui/qkeysequence.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qvariant.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qshortcut.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QShortcut__Go : public QShortcut {
public:
	long long __goId;
	unsigned char __goOn[8] = {};
	template<typename... A> QShortcut__Go(long long id, A&&... a) : QShortcut(std::forward<A>(a)...), __goId(id) {}
	using QShortcut::sender;
	using QShortcut::senderSignalIndex;
	using QShortcut::receivers;
	using QShortcut::isSignalConnected;
	bool cppgoBase_event(QEvent* __a0) { return QShortcut::event(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QShortcut::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QShortcut::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QShortcut::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QShortcut::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QShortcut::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QShortcut::metaObject();
	}
	bool event(QEvent* e) override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)e, &__h); if (__h) return __r; }
		return QShortcut::event(e);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QShortcut::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QShortcut::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QShortcut::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QShortcut::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QShortcut::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QShortcut::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QShortcut;
static std::mutex __connsMu_QShortcut;

extern "C" {

void* QShortcut_base(void* self) { return (void*)static_cast<QObject*>((QShortcut*)self); }
void QShortcut__set_override(void* self, int idx) { ((QShortcut__Go*)self)->__goOn[idx] = 1; }
void* QShortcut_fromQObject(void* o) { return (void*)qobject_cast<QShortcut*>((QObject*)o); }
void* QShortcut_new(long long __goid, void* parent) {
	return new QShortcut__Go(__goid, (QObject*)parent);
}
void* QShortcut_new2(long long __goid, void* key, void* parent) {
	return new QShortcut__Go(__goid, *(QKeySequence*)key, (QObject*)parent);
}
void* QShortcut_new3(long long __goid, int key, void* parent) {
	return new QShortcut__Go(__goid, (QKeySequence::StandardKey)(key), (QObject*)parent);
}
void* QShortcut_new4(long long __goid, void* key, void* parent, const char* member) {
	return new QShortcut__Go(__goid, *(QKeySequence*)key, (QObject*)parent, member);
}
void* QShortcut_new5(long long __goid, void* key, void* parent, const char* member, const char* ambiguousMember) {
	return new QShortcut__Go(__goid, *(QKeySequence*)key, (QObject*)parent, member, ambiguousMember);
}
void* QShortcut_new6(long long __goid, void* key, void* parent, const char* member, const char* ambiguousMember, int context) {
	return new QShortcut__Go(__goid, *(QKeySequence*)key, (QObject*)parent, member, ambiguousMember, (Qt::ShortcutContext)(context));
}
void* QShortcut_new7(long long __goid, int key, void* parent, const char* member) {
	return new QShortcut__Go(__goid, (QKeySequence::StandardKey)(key), (QObject*)parent, member);
}
void* QShortcut_new8(long long __goid, int key, void* parent, const char* member, const char* ambiguousMember) {
	return new QShortcut__Go(__goid, (QKeySequence::StandardKey)(key), (QObject*)parent, member, ambiguousMember);
}
void* QShortcut_new9(long long __goid, int key, void* parent, const char* member, const char* ambiguousMember, int context) {
	return new QShortcut__Go(__goid, (QKeySequence::StandardKey)(key), (QObject*)parent, member, ambiguousMember, (Qt::ShortcutContext)(context));
}
char* QShortcut_tr(const char* s, long long* __out_len) {
	auto&& __r = QShortcut::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QShortcut_setKey(void* self, void* key) {
	((QShortcut*)self)->setKey(*(QKeySequence*)key);
}
void* QShortcut_key(void* self) {
	auto&& __r = ((QShortcut*)self)->key();
	return new QKeySequence(__r);
}
void QShortcut_setKeys(void* self, int key) {
	((QShortcut*)self)->setKeys((QKeySequence::StandardKey)(key));
}
void QShortcut_setKeys2(void* self, void** keys, long long keys_len) {
	QList<QKeySequence> keys_list; keys_list.reserve(keys_len);
	for (long long __i = 0; __i < keys_len; __i++) keys_list.append(*(QKeySequence*)keys[__i]);
	((QShortcut*)self)->setKeys(keys_list);
}
void** QShortcut_keys(void* self, long long* __out_len) {
	auto&& __r = ((QShortcut*)self)->keys();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QKeySequence(__e);
	return __arr;
}
void QShortcut_setEnabled(void* self, int enable) {
	((QShortcut*)self)->setEnabled(enable);
}
int QShortcut_isEnabled(void* self) {
	auto&& __r = ((QShortcut*)self)->isEnabled();
	return __r;
}
void QShortcut_setContext(void* self, int context) {
	((QShortcut*)self)->setContext((Qt::ShortcutContext)(context));
}
int QShortcut_context(void* self) {
	auto&& __r = ((QShortcut*)self)->context();
	return (int)(__r);
}
void QShortcut_setAutoRepeat(void* self, int on) {
	((QShortcut*)self)->setAutoRepeat(on);
}
int QShortcut_autoRepeat(void* self) {
	auto&& __r = ((QShortcut*)self)->autoRepeat();
	return __r;
}
int QShortcut_id(void* self) {
	auto&& __r = ((QShortcut*)self)->id();
	return __r;
}
void QShortcut_setWhatsThis(void* self, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	((QShortcut*)self)->setWhatsThis(text_qstr);
}
char* QShortcut_whatsThis(void* self, long long* __out_len) {
	auto&& __r = ((QShortcut*)self)->whatsThis();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QShortcut_activated(void* self) {
	((QShortcut*)self)->activated();
}
void QShortcut_activatedAmbiguously(void* self) {
	((QShortcut*)self)->activatedAmbiguously();
}
void* QShortcut_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QShortcut::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QShortcut_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QShortcut::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QShortcut_sender(void* self) {
	auto&& __r = ((QShortcut__Go*)self)->sender();
	return (void*)(__r);
}
int QShortcut_senderSignalIndex(void* self) {
	auto&& __r = ((QShortcut__Go*)self)->senderSignalIndex();
	return __r;
}
int QShortcut_receivers(void* self, const char* signal) {
	auto&& __r = ((QShortcut__Go*)self)->receivers(signal);
	return __r;
}
int QShortcut_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QShortcut__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QShortcut_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QShortcut::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QShortcut_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QShortcut::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QShortcut_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QShortcut::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QShortcut_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QShortcut::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QShortcut_metaObject_default(void* self) {
	auto&& __r = ((QShortcut*)self)->QShortcut::metaObject();
	return (void*)(__r);
}
int QShortcut_event_default(void* self, void* e) {
	auto&& __r = ((QShortcut__Go*)self)->cppgoBase_event((QEvent*)e);
	return __r;
}
int QShortcut_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QShortcut*)self)->QShortcut::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QShortcut_timerEvent_default(void* self, void* event) {
	((QShortcut__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QShortcut_childEvent_default(void* self, void* event) {
	((QShortcut__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QShortcut_customEvent_default(void* self, void* event) {
	((QShortcut__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QShortcut_connectNotify_default(void* self, void* signal) {
	((QShortcut__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QShortcut_disconnectNotify_default(void* self, void* signal) {
	((QShortcut__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QShortcut_connect_activated(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QShortcut*)self, &QShortcut::activated, (QShortcut*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QShortcut);
	__conns_QShortcut[__id] = __c;
}
void QShortcut_connect_activatedAmbiguously(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QShortcut*)self, &QShortcut::activatedAmbiguously, (QShortcut*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QShortcut);
	__conns_QShortcut[__id] = __c;
}
void QShortcut_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QShortcut*)self, &QShortcut::destroyed, (QShortcut*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QShortcut);
	__conns_QShortcut[__id] = __c;
}
void QShortcut_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QShortcut*)self, &QShortcut::destroyed, (QShortcut*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QShortcut);
	__conns_QShortcut[__id] = __c;
}
void QShortcut_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QShortcut*)self, &QShortcut::objectNameChanged, (QShortcut*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QShortcut);
	__conns_QShortcut[__id] = __c;
}
void QShortcut__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QShortcut);
		auto __it = __conns_QShortcut.find(id);
		if (__it == __conns_QShortcut.end()) return;
		__c = __it->second;
		__conns_QShortcut.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QShortcut__Go_new(long long id, void* parent) { return new QShortcut__Go(id, (QObject*)parent); }
void QShortcut_delete(void* self) { delete (QShortcut*)self; }

}

#include <QtGui/qvalidator.h>
#include <QtCore/qobject.h>
#include <QtCore/qregularexpression.h>
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
#include "gen_qregularexpressionvalidator.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QRegularExpressionValidator__Go : public QRegularExpressionValidator {
public:
	long long __goId;
	unsigned char __goOn[10] = {};
	template<typename... A> QRegularExpressionValidator__Go(long long id, A&&... a) : QRegularExpressionValidator(std::forward<A>(a)...), __goId(id) {}
	using QRegularExpressionValidator::sender;
	using QRegularExpressionValidator::senderSignalIndex;
	using QRegularExpressionValidator::receivers;
	using QRegularExpressionValidator::isSignalConnected;
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QRegularExpressionValidator::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QRegularExpressionValidator::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QRegularExpressionValidator::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QRegularExpressionValidator::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QRegularExpressionValidator::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QRegularExpressionValidator::metaObject();
	}
	QValidator::State validate(QString& input, int& pos) const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_int_voidp_intp) { int __r = ((int(*)(long long, long long, void*, int*, int*))cppgo_virt_int_voidp_intp)(__goId, 6784179835071232555LL, (void*)&input, (int*)&pos, &__h); if (__h) return (QValidator::State)(__r); }
		return QRegularExpressionValidator::validate(input, pos);
	}
	void fixup(QString& arg1) const override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4855206404746966715LL, (void*)&arg1, &__h); if (__h) return; }
		QRegularExpressionValidator::fixup(arg1);
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QRegularExpressionValidator::event(event);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QRegularExpressionValidator::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QRegularExpressionValidator::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QRegularExpressionValidator::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QRegularExpressionValidator::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QRegularExpressionValidator::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QRegularExpressionValidator::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QRegularExpressionValidator;
static std::mutex __connsMu_QRegularExpressionValidator;

extern "C" {

void* QRegularExpressionValidator_base(void* self) { return (void*)static_cast<QValidator*>((QRegularExpressionValidator*)self); }
void QRegularExpressionValidator_chain(void* self, void** out) {
	QRegularExpressionValidator* __p = (QRegularExpressionValidator*)self;
	QValidator* __b0 = static_cast<QValidator*>(__p); out[0] = (void*)__b0;
	QObject* __b1 = static_cast<QObject*>(__b0); out[1] = (void*)__b1;
}
void QRegularExpressionValidator__set_override(void* self, int idx) { ((QRegularExpressionValidator__Go*)self)->__goOn[idx] = 1; }
void* QRegularExpressionValidator_fromQObject(void* o) { return (void*)qobject_cast<QRegularExpressionValidator*>((QObject*)o); }
void* QRegularExpressionValidator_new(long long __goid) {
	return new QRegularExpressionValidator__Go(__goid);
}
void* QRegularExpressionValidator_new2(long long __goid, void* re) {
	return new QRegularExpressionValidator__Go(__goid, *(QRegularExpression*)re);
}
void* QRegularExpressionValidator_new3(long long __goid, void* parent) {
	return new QRegularExpressionValidator__Go(__goid, (QObject*)parent);
}
void* QRegularExpressionValidator_new4(long long __goid, void* re, void* parent) {
	return new QRegularExpressionValidator__Go(__goid, *(QRegularExpression*)re, (QObject*)parent);
}
char* QRegularExpressionValidator_tr(const char* s, long long* __out_len) {
	auto&& __r = QRegularExpressionValidator::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QRegularExpressionValidator_regularExpression(void* self) {
	auto&& __r = ((QRegularExpressionValidator*)self)->regularExpression();
	return new QRegularExpression(__r);
}
void QRegularExpressionValidator_setRegularExpression(void* self, void* re) {
	((QRegularExpressionValidator*)self)->setRegularExpression(*(QRegularExpression*)re);
}
void QRegularExpressionValidator_regularExpressionChanged(void* self, void* re) {
	((QRegularExpressionValidator*)self)->regularExpressionChanged(*(QRegularExpression*)re);
}
void* QRegularExpressionValidator_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QRegularExpressionValidator::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QRegularExpressionValidator_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QRegularExpressionValidator::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QRegularExpressionValidator_sender(void* self) {
	auto&& __r = ((QRegularExpressionValidator__Go*)self)->sender();
	return (void*)(__r);
}
int QRegularExpressionValidator_senderSignalIndex(void* self) {
	auto&& __r = ((QRegularExpressionValidator__Go*)self)->senderSignalIndex();
	return __r;
}
int QRegularExpressionValidator_receivers(void* self, const char* signal) {
	auto&& __r = ((QRegularExpressionValidator__Go*)self)->receivers(signal);
	return __r;
}
int QRegularExpressionValidator_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QRegularExpressionValidator__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QRegularExpressionValidator_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QRegularExpressionValidator::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QRegularExpressionValidator_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QRegularExpressionValidator::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QRegularExpressionValidator_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QRegularExpressionValidator::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QRegularExpressionValidator_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QRegularExpressionValidator::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QRegularExpressionValidator_metaObject_default(void* self) {
	auto&& __r = ((QRegularExpressionValidator*)self)->QRegularExpressionValidator::metaObject();
	return (void*)(__r);
}
int QRegularExpressionValidator_validate_default(void* self, void* input, int* pos) {
	int pos_ref = *pos;
	auto&& __r = ((QRegularExpressionValidator*)self)->QRegularExpressionValidator::validate(*(QString*)input, pos_ref);
	*pos = pos_ref;
	return (int)(__r);
}
void QRegularExpressionValidator_fixup_default(void* self, void* arg1) {
	((QRegularExpressionValidator*)self)->QRegularExpressionValidator::fixup(*(QString*)arg1);
}
int QRegularExpressionValidator_event_default(void* self, void* event) {
	auto&& __r = ((QRegularExpressionValidator*)self)->QRegularExpressionValidator::event((QEvent*)event);
	return __r;
}
int QRegularExpressionValidator_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QRegularExpressionValidator*)self)->QRegularExpressionValidator::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QRegularExpressionValidator_timerEvent_default(void* self, void* event) {
	((QRegularExpressionValidator__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QRegularExpressionValidator_childEvent_default(void* self, void* event) {
	((QRegularExpressionValidator__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QRegularExpressionValidator_customEvent_default(void* self, void* event) {
	((QRegularExpressionValidator__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QRegularExpressionValidator_connectNotify_default(void* self, void* signal) {
	((QRegularExpressionValidator__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QRegularExpressionValidator_disconnectNotify_default(void* self, void* signal) {
	((QRegularExpressionValidator__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QRegularExpressionValidator_connect_regularExpressionChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QRegularExpressionValidator*)self, &QRegularExpressionValidator::regularExpressionChanged, (QRegularExpressionValidator*)self, [__id](const QRegularExpression& re){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&re);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QRegularExpressionValidator);
	__conns_QRegularExpressionValidator[__id] = __c;
}
void QRegularExpressionValidator_connect_changed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QRegularExpressionValidator*)self, &QRegularExpressionValidator::changed, (QRegularExpressionValidator*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QRegularExpressionValidator);
	__conns_QRegularExpressionValidator[__id] = __c;
}
void QRegularExpressionValidator_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QRegularExpressionValidator*)self, &QRegularExpressionValidator::destroyed, (QRegularExpressionValidator*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QRegularExpressionValidator);
	__conns_QRegularExpressionValidator[__id] = __c;
}
void QRegularExpressionValidator_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QRegularExpressionValidator*)self, &QRegularExpressionValidator::destroyed, (QRegularExpressionValidator*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QRegularExpressionValidator);
	__conns_QRegularExpressionValidator[__id] = __c;
}
void QRegularExpressionValidator_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QRegularExpressionValidator*)self, &QRegularExpressionValidator::objectNameChanged, (QRegularExpressionValidator*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QRegularExpressionValidator);
	__conns_QRegularExpressionValidator[__id] = __c;
}
void QRegularExpressionValidator__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QRegularExpressionValidator);
		auto __it = __conns_QRegularExpressionValidator.find(id);
		if (__it == __conns_QRegularExpressionValidator.end()) return;
		__c = __it->second;
		__conns_QRegularExpressionValidator.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QRegularExpressionValidator__Go_new(long long id, void* parent) { return new QRegularExpressionValidator__Go(id, (QObject*)parent); }
void QRegularExpressionValidator_delete(void* self) { delete (QRegularExpressionValidator*)self; }

}

#include <QtWidgets/qstyleoption.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qstring.h>
#include <QtCore/qnamespace.h>
#include "gen_qstyleoptionprogressbar.h"

extern "C" {

void* QStyleOptionProgressBar_base(void* self) { return (void*)static_cast<QStyleOption*>((QStyleOptionProgressBar*)self); }
void* QStyleOptionProgressBar_new() {
	return new QStyleOptionProgressBar();
}
void* QStyleOptionProgressBar_new2(void* other) {
	return new QStyleOptionProgressBar(*(QStyleOptionProgressBar*)other);
}
void* QStyleOptionProgressBar_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QStyleOptionProgressBar*)self)->operator=(*(QStyleOptionProgressBar*)arg1);
	return (void*)&(__r);
}
int QStyleOptionProgressBar_minimum_get(void* self) {
	auto __r = ((QStyleOptionProgressBar*)self)->minimum;
	return __r;
}
void QStyleOptionProgressBar_minimum_set(void* self, int minimum) {
	((QStyleOptionProgressBar*)self)->minimum = minimum;
}
int QStyleOptionProgressBar_maximum_get(void* self) {
	auto __r = ((QStyleOptionProgressBar*)self)->maximum;
	return __r;
}
void QStyleOptionProgressBar_maximum_set(void* self, int maximum) {
	((QStyleOptionProgressBar*)self)->maximum = maximum;
}
int QStyleOptionProgressBar_progress_get(void* self) {
	auto __r = ((QStyleOptionProgressBar*)self)->progress;
	return __r;
}
void QStyleOptionProgressBar_progress_set(void* self, int progress) {
	((QStyleOptionProgressBar*)self)->progress = progress;
}
char* QStyleOptionProgressBar_text_get(void* self, long long* __out_len) {
	auto __r = ((QStyleOptionProgressBar*)self)->text;
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QStyleOptionProgressBar_text_set(void* self, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	((QStyleOptionProgressBar*)self)->text = text_qstr;
}
int QStyleOptionProgressBar_textAlignment_get(void* self) {
	auto __r = ((QStyleOptionProgressBar*)self)->textAlignment;
	return (int)(__r);
}
void QStyleOptionProgressBar_textAlignment_set(void* self, int textAlignment) {
	((QStyleOptionProgressBar*)self)->textAlignment = QFlags<Qt::AlignmentFlag>(textAlignment);
}
int QStyleOptionProgressBar_textVisible_get(void* self) {
	auto __r = ((QStyleOptionProgressBar*)self)->textVisible;
	return __r;
}
void QStyleOptionProgressBar_textVisible_set(void* self, int textVisible) {
	((QStyleOptionProgressBar*)self)->textVisible = textVisible;
}
int QStyleOptionProgressBar_invertedAppearance_get(void* self) {
	auto __r = ((QStyleOptionProgressBar*)self)->invertedAppearance;
	return __r;
}
void QStyleOptionProgressBar_invertedAppearance_set(void* self, int invertedAppearance) {
	((QStyleOptionProgressBar*)self)->invertedAppearance = invertedAppearance;
}
int QStyleOptionProgressBar_bottomToTop_get(void* self) {
	auto __r = ((QStyleOptionProgressBar*)self)->bottomToTop;
	return __r;
}
void QStyleOptionProgressBar_bottomToTop_set(void* self, int bottomToTop) {
	((QStyleOptionProgressBar*)self)->bottomToTop = bottomToTop;
}
void QStyleOptionProgressBar_delete(void* self) { delete (QStyleOptionProgressBar*)self; }

}

#include <QtWidgets/qapplication.h>
#include <QtGui/qguiapplication.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtWidgets/qstyle.h>
#include <QtGui/qpalette.h>
#include <QtWidgets/qwidget.h>
#include <QtGui/qfont.h>
#include <QtGui/qfontmetrics.h>
#include <QtCore/qpoint.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qobject.h>
#include <QtCore/qcoreevent.h>
#include <QtGui/qwindow.h>
#include <QtGui/qicon.h>
#include <QtGui/qscreen.h>
#include <QtGui/qcursor.h>
#include <QtGui/qclipboard.h>
#include <QtGui/qstylehints.h>
#include <QtGui/qinputmethod.h>
#include <QtCore/qbytearray.h>
#include <QtGui/qsessionmanager.h>
#include <QtCore/qcoreapplication.h>
#include <QtCore/qeventloop.h>
#include <QtCore/qdeadlinetimer.h>
#include <QtCore/qabstracteventdispatcher.h>
#include <QtCore/qpermissions.h>
#include <QtCore/qtranslator.h>
#include <QtCore/qabstractnativeeventfilter.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qvariant.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qapplication.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

static std::unordered_map<long long, QMetaObject::Connection> __conns_QApplication;
static std::mutex __connsMu_QApplication;

extern "C" {

void* QApplication_base(void* self) { return (void*)static_cast<QGuiApplication*>((QApplication*)self); }
void QApplication_chain(void* self, void** out) {
	QApplication* __p = (QApplication*)self;
	QGuiApplication* __b0 = static_cast<QGuiApplication*>(__p); out[0] = (void*)__b0;
	QCoreApplication* __b1 = static_cast<QCoreApplication*>(__b0); out[1] = (void*)__b1;
	QObject* __b2 = static_cast<QObject*>(__b1); out[2] = (void*)__b2;
}
void* QApplication_fromQObject(void* o) { return (void*)qobject_cast<QApplication*>((QObject*)o); }
char* QApplication_tr(const char* s, long long* __out_len) {
	auto&& __r = QApplication::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QApplication_style() {
	auto&& __r = QApplication::style();
	return (void*)(__r);
}
void QApplication_setStyle(void* arg1) {
	QApplication::setStyle((QStyle*)arg1);
}
void* QApplication_setStyle2(const char* arg1, long long arg1_len) {
	QString arg1_qstr = QString::fromUtf8(arg1, arg1_len);
	auto&& __r = QApplication::setStyle(arg1_qstr);
	return (void*)(__r);
}
void* QApplication_palette(void* arg1) {
	auto&& __r = QApplication::palette((QWidget*)arg1);
	return new QPalette(__r);
}
void* QApplication_palette2(const char* className) {
	auto&& __r = QApplication::palette(className);
	return new QPalette(__r);
}
void QApplication_setPalette(void* arg1) {
	QApplication::setPalette(*(QPalette*)arg1);
}
void* QApplication_font() {
	auto&& __r = QApplication::font();
	return new QFont(__r);
}
void* QApplication_font2(void* arg1) {
	auto&& __r = QApplication::font((QWidget*)arg1);
	return new QFont(__r);
}
void* QApplication_font3(const char* className) {
	auto&& __r = QApplication::font(className);
	return new QFont(__r);
}
void QApplication_setFont(void* arg1) {
	QApplication::setFont(*(QFont*)arg1);
}
void* QApplication_fontMetrics() {
	auto&& __r = QApplication::fontMetrics();
	return new QFontMetrics(__r);
}
void** QApplication_allWidgets(long long* __out_len) {
	auto&& __r = QApplication::allWidgets();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
void** QApplication_topLevelWidgets(long long* __out_len) {
	auto&& __r = QApplication::topLevelWidgets();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
void* QApplication_activePopupWidget() {
	auto&& __r = QApplication::activePopupWidget();
	return (void*)(__r);
}
void* QApplication_activeModalWidget() {
	auto&& __r = QApplication::activeModalWidget();
	return (void*)(__r);
}
void* QApplication_focusWidget() {
	auto&& __r = QApplication::focusWidget();
	return (void*)(__r);
}
void* QApplication_activeWindow() {
	auto&& __r = QApplication::activeWindow();
	return (void*)(__r);
}
void QApplication_setActiveWindow(void* act) {
	QApplication::setActiveWindow((QWidget*)act);
}
void* QApplication_widgetAt(void* p) {
	auto&& __r = QApplication::widgetAt(*(QPoint*)p);
	return (void*)(__r);
}
void* QApplication_widgetAt2(int x, int y) {
	auto&& __r = QApplication::widgetAt(x, y);
	return (void*)(__r);
}
void* QApplication_topLevelAt(void* p) {
	auto&& __r = QApplication::topLevelAt(*(QPoint*)p);
	return (void*)(__r);
}
void* QApplication_topLevelAt2(int x, int y) {
	auto&& __r = QApplication::topLevelAt(x, y);
	return (void*)(__r);
}
void QApplication_beep() {
	QApplication::beep();
}
void QApplication_alert(void* widget) {
	QApplication::alert((QWidget*)widget);
}
void QApplication_setCursorFlashTime(int arg1) {
	QApplication::setCursorFlashTime(arg1);
}
int QApplication_cursorFlashTime() {
	auto&& __r = QApplication::cursorFlashTime();
	return __r;
}
void QApplication_setDoubleClickInterval(int arg1) {
	QApplication::setDoubleClickInterval(arg1);
}
int QApplication_doubleClickInterval() {
	auto&& __r = QApplication::doubleClickInterval();
	return __r;
}
void QApplication_setKeyboardInputInterval(int arg1) {
	QApplication::setKeyboardInputInterval(arg1);
}
int QApplication_keyboardInputInterval() {
	auto&& __r = QApplication::keyboardInputInterval();
	return __r;
}
void QApplication_setWheelScrollLines(int arg1) {
	QApplication::setWheelScrollLines(arg1);
}
int QApplication_wheelScrollLines() {
	auto&& __r = QApplication::wheelScrollLines();
	return __r;
}
void QApplication_setStartDragTime(int ms) {
	QApplication::setStartDragTime(ms);
}
int QApplication_startDragTime() {
	auto&& __r = QApplication::startDragTime();
	return __r;
}
void QApplication_setStartDragDistance(int l) {
	QApplication::setStartDragDistance(l);
}
int QApplication_startDragDistance() {
	auto&& __r = QApplication::startDragDistance();
	return __r;
}
int QApplication_isEffectEnabled(int arg1) {
	auto&& __r = QApplication::isEffectEnabled((Qt::UIEffect)(arg1));
	return __r;
}
void QApplication_setEffectEnabled(int arg1) {
	QApplication::setEffectEnabled((Qt::UIEffect)(arg1));
}
int QApplication_exec() {
	auto&& __r = QApplication::exec();
	return __r;
}
void QApplication_focusChanged(void* self, void* old, void* now) {
	((QApplication*)self)->focusChanged((QWidget*)old, (QWidget*)now);
}
char* QApplication_styleSheet(void* self, long long* __out_len) {
	auto&& __r = ((QApplication*)self)->styleSheet();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QApplication_autoSipEnabled(void* self) {
	auto&& __r = ((QApplication*)self)->autoSipEnabled();
	return __r;
}
void QApplication_setStyleSheet(void* self, const char* sheet, long long sheet_len) {
	QString sheet_qstr = QString::fromUtf8(sheet, sheet_len);
	((QApplication*)self)->setStyleSheet(sheet_qstr);
}
void QApplication_setAutoSipEnabled(void* self, int enabled) {
	((QApplication*)self)->setAutoSipEnabled(enabled);
}
void QApplication_closeAllWindows() {
	QApplication::closeAllWindows();
}
void QApplication_aboutQt() {
	QApplication::aboutQt();
}
void QApplication_setApplicationDisplayName(const char* name, long long name_len) {
	QString name_qstr = QString::fromUtf8(name, name_len);
	QApplication::setApplicationDisplayName(name_qstr);
}
char* QApplication_applicationDisplayName(long long* __out_len) {
	auto&& __r = QApplication::applicationDisplayName();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QApplication_setDesktopFileName(const char* name, long long name_len) {
	QString name_qstr = QString::fromUtf8(name, name_len);
	QApplication::setDesktopFileName(name_qstr);
}
char* QApplication_desktopFileName(long long* __out_len) {
	auto&& __r = QApplication::desktopFileName();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void** QApplication_allWindows(long long* __out_len) {
	auto&& __r = QApplication::allWindows();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
void** QApplication_topLevelWindows(long long* __out_len) {
	auto&& __r = QApplication::topLevelWindows();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
void QApplication_setWindowIcon(void* icon) {
	QApplication::setWindowIcon(*(QIcon*)icon);
}
void* QApplication_windowIcon() {
	auto&& __r = QApplication::windowIcon();
	return new QIcon(__r);
}
char* QApplication_platformName(long long* __out_len) {
	auto&& __r = QApplication::platformName();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QApplication_modalWindow() {
	auto&& __r = QApplication::modalWindow();
	return (void*)(__r);
}
void* QApplication_focusWindow() {
	auto&& __r = QApplication::focusWindow();
	return (void*)(__r);
}
void* QApplication_focusObject() {
	auto&& __r = QApplication::focusObject();
	return (void*)(__r);
}
void* QApplication_primaryScreen() {
	auto&& __r = QApplication::primaryScreen();
	return (void*)(__r);
}
void** QApplication_screens(long long* __out_len) {
	auto&& __r = QApplication::screens();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
void* QApplication_screenAt(void* point) {
	auto&& __r = QApplication::screenAt(*(QPoint*)point);
	return (void*)(__r);
}
void* QApplication_overrideCursor() {
	auto&& __r = QApplication::overrideCursor();
	return (void*)(__r);
}
void QApplication_setOverrideCursor(void* arg1) {
	QApplication::setOverrideCursor(*(QCursor*)arg1);
}
void QApplication_changeOverrideCursor(void* arg1) {
	QApplication::changeOverrideCursor(*(QCursor*)arg1);
}
void QApplication_restoreOverrideCursor() {
	QApplication::restoreOverrideCursor();
}
void* QApplication_clipboard() {
	auto&& __r = QApplication::clipboard();
	return (void*)(__r);
}
int QApplication_keyboardModifiers() {
	auto&& __r = QApplication::keyboardModifiers();
	return (int)(__r);
}
int QApplication_queryKeyboardModifiers() {
	auto&& __r = QApplication::queryKeyboardModifiers();
	return (int)(__r);
}
int QApplication_mouseButtons() {
	auto&& __r = QApplication::mouseButtons();
	return (int)(__r);
}
void QApplication_setLayoutDirection(int direction) {
	QApplication::setLayoutDirection((Qt::LayoutDirection)(direction));
}
int QApplication_layoutDirection() {
	auto&& __r = QApplication::layoutDirection();
	return (int)(__r);
}
int QApplication_isRightToLeft() {
	auto&& __r = QApplication::isRightToLeft();
	return __r;
}
int QApplication_isLeftToRight() {
	auto&& __r = QApplication::isLeftToRight();
	return __r;
}
void* QApplication_styleHints() {
	auto&& __r = QApplication::styleHints();
	return (void*)(__r);
}
void QApplication_setDesktopSettingsAware(int on) {
	QApplication::setDesktopSettingsAware(on);
}
int QApplication_desktopSettingsAware() {
	auto&& __r = QApplication::desktopSettingsAware();
	return __r;
}
void* QApplication_inputMethod() {
	auto&& __r = QApplication::inputMethod();
	return (void*)(__r);
}
void QApplication_setQuitOnLastWindowClosed(int quit) {
	QApplication::setQuitOnLastWindowClosed(quit);
}
int QApplication_quitOnLastWindowClosed() {
	auto&& __r = QApplication::quitOnLastWindowClosed();
	return __r;
}
int QApplication_applicationState() {
	auto&& __r = QApplication::applicationState();
	return (int)(__r);
}
void QApplication_setHighDpiScaleFactorRoundingPolicy(int policy) {
	QApplication::setHighDpiScaleFactorRoundingPolicy((Qt::HighDpiScaleFactorRoundingPolicy)(policy));
}
int QApplication_highDpiScaleFactorRoundingPolicy() {
	auto&& __r = QApplication::highDpiScaleFactorRoundingPolicy();
	return (int)(__r);
}
void QApplication_sync() {
	QApplication::sync();
}
char** QApplication_arguments(long long* __out_len) {
	auto&& __r = QApplication::arguments();
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
void QApplication_setAttribute(int attribute) {
	QApplication::setAttribute((Qt::ApplicationAttribute)(attribute));
}
int QApplication_testAttribute(int attribute) {
	auto&& __r = QApplication::testAttribute((Qt::ApplicationAttribute)(attribute));
	return __r;
}
void QApplication_setOrganizationDomain(const char* orgDomain, long long orgDomain_len) {
	QString orgDomain_qstr = QString::fromUtf8(orgDomain, orgDomain_len);
	QApplication::setOrganizationDomain(orgDomain_qstr);
}
char* QApplication_organizationDomain(long long* __out_len) {
	auto&& __r = QApplication::organizationDomain();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QApplication_setOrganizationName(const char* orgName, long long orgName_len) {
	QString orgName_qstr = QString::fromUtf8(orgName, orgName_len);
	QApplication::setOrganizationName(orgName_qstr);
}
char* QApplication_organizationName(long long* __out_len) {
	auto&& __r = QApplication::organizationName();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QApplication_setApplicationName(const char* application, long long application_len) {
	QString application_qstr = QString::fromUtf8(application, application_len);
	QApplication::setApplicationName(application_qstr);
}
char* QApplication_applicationName(long long* __out_len) {
	auto&& __r = QApplication::applicationName();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QApplication_setApplicationVersion(const char* version, long long version_len) {
	QString version_qstr = QString::fromUtf8(version, version_len);
	QApplication::setApplicationVersion(version_qstr);
}
char* QApplication_applicationVersion(long long* __out_len) {
	auto&& __r = QApplication::applicationVersion();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QApplication_setSetuidAllowed(int allow) {
	QApplication::setSetuidAllowed(allow);
}
int QApplication_isSetuidAllowed() {
	auto&& __r = QApplication::isSetuidAllowed();
	return __r;
}
void* QApplication_instance() {
	auto&& __r = QApplication::instance();
	return (void*)(__r);
}
int QApplication_instanceExists() {
	auto&& __r = QApplication::instanceExists();
	return __r;
}
void QApplication_processEvents() {
	QApplication::processEvents();
}
void QApplication_processEvents2(int flags, int maxtime) {
	QApplication::processEvents(QFlags<QEventLoop::ProcessEventsFlag>(flags), maxtime);
}
void QApplication_processEvents3(int flags, void* deadline) {
	QApplication::processEvents(QFlags<QEventLoop::ProcessEventsFlag>(flags), *(QDeadlineTimer*)deadline);
}
int QApplication_sendEvent(void* receiver, void* event) {
	auto&& __r = QApplication::sendEvent((QObject*)receiver, (QEvent*)event);
	return __r;
}
void QApplication_postEvent(void* receiver, void* event) {
	QApplication::postEvent((QObject*)receiver, (QEvent*)event);
}
void QApplication_sendPostedEvents() {
	QApplication::sendPostedEvents();
}
void QApplication_removePostedEvents(void* receiver) {
	QApplication::removePostedEvents((QObject*)receiver);
}
void* QApplication_eventDispatcher() {
	auto&& __r = QApplication::eventDispatcher();
	return (void*)(__r);
}
void QApplication_setEventDispatcher(void* eventDispatcher) {
	QApplication::setEventDispatcher((QAbstractEventDispatcher*)eventDispatcher);
}
int QApplication_startingUp() {
	auto&& __r = QApplication::startingUp();
	return __r;
}
int QApplication_closingDown() {
	auto&& __r = QApplication::closingDown();
	return __r;
}
char* QApplication_applicationDirPath(long long* __out_len) {
	auto&& __r = QApplication::applicationDirPath();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QApplication_applicationFilePath(long long* __out_len) {
	auto&& __r = QApplication::applicationFilePath();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
long long QApplication_applicationPid() {
	auto&& __r = QApplication::applicationPid();
	return __r;
}
void QApplication_setLibraryPaths(char** arg1, long long arg1_len) {
	QList<QString> arg1_list; arg1_list.reserve(arg1_len);
	for (long long __i = 0; __i < arg1_len; __i++) arg1_list.append(QString::fromUtf8(arg1[__i]));
	QApplication::setLibraryPaths(arg1_list);
}
char** QApplication_libraryPaths(long long* __out_len) {
	auto&& __r = QApplication::libraryPaths();
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
void QApplication_addLibraryPath(const char* arg1, long long arg1_len) {
	QString arg1_qstr = QString::fromUtf8(arg1, arg1_len);
	QApplication::addLibraryPath(arg1_qstr);
}
void QApplication_removeLibraryPath(const char* arg1, long long arg1_len) {
	QString arg1_qstr = QString::fromUtf8(arg1, arg1_len);
	QApplication::removeLibraryPath(arg1_qstr);
}
int QApplication_installTranslator(void* messageFile) {
	auto&& __r = QApplication::installTranslator((QTranslator*)messageFile);
	return __r;
}
int QApplication_removeTranslator(void* messageFile) {
	auto&& __r = QApplication::removeTranslator((QTranslator*)messageFile);
	return __r;
}
char* QApplication_translate(const char* context, const char* key, long long* __out_len) {
	auto&& __r = QApplication::translate(context, key);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QApplication_isQuitLockEnabled() {
	auto&& __r = QApplication::isQuitLockEnabled();
	return __r;
}
void QApplication_setQuitLockEnabled(int enabled) {
	QApplication::setQuitLockEnabled(enabled);
}
void QApplication_quit() {
	QApplication::quit();
}
void QApplication_exit() {
	QApplication::exit();
}
void* QApplication_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QApplication::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QApplication_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QApplication::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
char* QApplication_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QApplication::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QApplication_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QApplication::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QApplication_setPalette2(void* arg1, const char* className) {
	QApplication::setPalette(*(QPalette*)arg1, className);
}
void QApplication_setFont2(void* arg1, const char* className) {
	QApplication::setFont(*(QFont*)arg1, className);
}
void QApplication_alert2(void* widget, int duration) {
	QApplication::alert((QWidget*)widget, duration);
}
void QApplication_setEffectEnabled2(int arg1, int enable) {
	QApplication::setEffectEnabled((Qt::UIEffect)(arg1), enable);
}
void QApplication_setAttribute2(int attribute, int on) {
	QApplication::setAttribute((Qt::ApplicationAttribute)(attribute), on);
}
void QApplication_processEvents4(int flags) {
	QApplication::processEvents(QFlags<QEventLoop::ProcessEventsFlag>(flags));
}
void QApplication_postEvent2(void* receiver, void* event, int priority) {
	QApplication::postEvent((QObject*)receiver, (QEvent*)event, priority);
}
void QApplication_sendPostedEvents2(void* receiver) {
	QApplication::sendPostedEvents((QObject*)receiver);
}
void QApplication_sendPostedEvents3(void* receiver, int event_type) {
	QApplication::sendPostedEvents((QObject*)receiver, event_type);
}
void QApplication_removePostedEvents2(void* receiver, int eventType) {
	QApplication::removePostedEvents((QObject*)receiver, eventType);
}
char* QApplication_translate2(const char* context, const char* key, const char* disambiguation, long long* __out_len) {
	auto&& __r = QApplication::translate(context, key, disambiguation);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QApplication_translate3(const char* context, const char* key, const char* disambiguation, int n, long long* __out_len) {
	auto&& __r = QApplication::translate(context, key, disambiguation, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QApplication_exit2(int retcode) {
	QApplication::exit(retcode);
}
void* QApplication_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QApplication::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QApplication_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QApplication::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void QApplication_connect_focusChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QApplication*)self, &QApplication::focusChanged, (QApplication*)self, [__id](QWidget* old, QWidget* now){
		if (cppgo_sig_void_voidp_voidp) ((void(*)(long long, void*, void*))cppgo_sig_void_voidp_voidp)(__id, (void*)old, (void*)now);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QApplication);
	__conns_QApplication[__id] = __c;
}
void QApplication_connect_fontDatabaseChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QApplication*)self, &QApplication::fontDatabaseChanged, (QApplication*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QApplication);
	__conns_QApplication[__id] = __c;
}
void QApplication_connect_screenAdded(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QApplication*)self, &QApplication::screenAdded, (QApplication*)self, [__id](QScreen* screen){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)screen);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QApplication);
	__conns_QApplication[__id] = __c;
}
void QApplication_connect_screenRemoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QApplication*)self, &QApplication::screenRemoved, (QApplication*)self, [__id](QScreen* screen){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)screen);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QApplication);
	__conns_QApplication[__id] = __c;
}
void QApplication_connect_primaryScreenChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QApplication*)self, &QApplication::primaryScreenChanged, (QApplication*)self, [__id](QScreen* screen){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)screen);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QApplication);
	__conns_QApplication[__id] = __c;
}
void QApplication_connect_lastWindowClosed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QApplication*)self, &QApplication::lastWindowClosed, (QApplication*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QApplication);
	__conns_QApplication[__id] = __c;
}
void QApplication_connect_focusObjectChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QApplication*)self, &QApplication::focusObjectChanged, (QApplication*)self, [__id](QObject* focusObject){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)focusObject);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QApplication);
	__conns_QApplication[__id] = __c;
}
void QApplication_connect_focusWindowChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QApplication*)self, &QApplication::focusWindowChanged, (QApplication*)self, [__id](QWindow* focusWindow){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)focusWindow);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QApplication);
	__conns_QApplication[__id] = __c;
}
void QApplication_connect_applicationStateChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QApplication*)self, &QApplication::applicationStateChanged, (QApplication*)self, [__id](Qt::ApplicationState state){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, (int)(state));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QApplication);
	__conns_QApplication[__id] = __c;
}
void QApplication_connect_layoutDirectionChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QApplication*)self, &QApplication::layoutDirectionChanged, (QApplication*)self, [__id](Qt::LayoutDirection direction){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, (int)(direction));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QApplication);
	__conns_QApplication[__id] = __c;
}
void QApplication_connect_commitDataRequest(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QApplication*)self, &QApplication::commitDataRequest, (QApplication*)self, [__id](QSessionManager& sessionManager){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&sessionManager);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QApplication);
	__conns_QApplication[__id] = __c;
}
void QApplication_connect_saveStateRequest(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QApplication*)self, &QApplication::saveStateRequest, (QApplication*)self, [__id](QSessionManager& sessionManager){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&sessionManager);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QApplication);
	__conns_QApplication[__id] = __c;
}
void QApplication_connect_applicationDisplayNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QApplication*)self, &QApplication::applicationDisplayNameChanged, (QApplication*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QApplication);
	__conns_QApplication[__id] = __c;
}
void QApplication_connect_paletteChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QApplication*)self, &QApplication::paletteChanged, (QApplication*)self, [__id](const QPalette& pal){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&pal);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QApplication);
	__conns_QApplication[__id] = __c;
}
void QApplication_connect_fontChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QApplication*)self, &QApplication::fontChanged, (QApplication*)self, [__id](const QFont& font){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&font);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QApplication);
	__conns_QApplication[__id] = __c;
}
void QApplication_connect_aboutToQuit(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QApplication*)self, &QApplication::aboutToQuit, (QApplication*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QApplication);
	__conns_QApplication[__id] = __c;
}
void QApplication_connect_organizationNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QApplication*)self, &QApplication::organizationNameChanged, (QApplication*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QApplication);
	__conns_QApplication[__id] = __c;
}
void QApplication_connect_organizationDomainChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QApplication*)self, &QApplication::organizationDomainChanged, (QApplication*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QApplication);
	__conns_QApplication[__id] = __c;
}
void QApplication_connect_applicationNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QApplication*)self, &QApplication::applicationNameChanged, (QApplication*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QApplication);
	__conns_QApplication[__id] = __c;
}
void QApplication_connect_applicationVersionChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QApplication*)self, &QApplication::applicationVersionChanged, (QApplication*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QApplication);
	__conns_QApplication[__id] = __c;
}
void QApplication_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QApplication*)self, &QApplication::destroyed, (QApplication*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QApplication);
	__conns_QApplication[__id] = __c;
}
void QApplication_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QApplication*)self, &QApplication::destroyed, (QApplication*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QApplication);
	__conns_QApplication[__id] = __c;
}
void QApplication_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QApplication*)self, &QApplication::objectNameChanged, (QApplication*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QApplication);
	__conns_QApplication[__id] = __c;
}
void QApplication__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QApplication);
		auto __it = __conns_QApplication.find(id);
		if (__it == __conns_QApplication.end()) return;
		__c = __it->second;
		__conns_QApplication.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QApplication_new_args(int* argc, char** argv) { return new QApplication(*argc, argv); }
void QApplication_delete(void* self) { delete (QApplication*)self; }

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
#include "gen_qtimeedit.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QTimeEdit__Go : public QTimeEdit {
public:
	long long __goId;
	unsigned char __goOn[56] = {};
	template<typename... A> QTimeEdit__Go(long long id, A&&... a) : QTimeEdit(std::forward<A>(a)...), __goId(id) {}
	using QTimeEdit::lineEdit;
	using QTimeEdit::setLineEdit;
	using QTimeEdit::updateMicroFocus;
	using QTimeEdit::create;
	using QTimeEdit::destroy;
	using QTimeEdit::focusNextChild;
	using QTimeEdit::focusPreviousChild;
	using QTimeEdit::sender;
	using QTimeEdit::senderSignalIndex;
	using QTimeEdit::receivers;
	using QTimeEdit::isSignalConnected;
	using QTimeEdit::getDecodedMetricF;
	void cppgoBase_keyPressEvent(QKeyEvent* __a0) { QTimeEdit::keyPressEvent(__a0); }
	void cppgoBase_wheelEvent(QWheelEvent* __a0) { QTimeEdit::wheelEvent(__a0); }
	void cppgoBase_focusInEvent(QFocusEvent* __a0) { QTimeEdit::focusInEvent(__a0); }
	bool cppgoBase_focusNextPrevChild(bool __a0) { return QTimeEdit::focusNextPrevChild(__a0); }
	QValidator::State cppgoBase_validate(QString& __a0, int& __a1) { return QTimeEdit::validate(__a0, __a1); }
	void cppgoBase_fixup(QString& __a0) { QTimeEdit::fixup(__a0); }
	QDateTime cppgoBase_dateTimeFromText(const QString& __a0) { return QTimeEdit::dateTimeFromText(__a0); }
	QString cppgoBase_textFromDateTime(const QDateTime& __a0) { return QTimeEdit::textFromDateTime(__a0); }
	QFlags<QAbstractSpinBox::StepEnabledFlag> cppgoBase_stepEnabled() { return QTimeEdit::stepEnabled(); }
	void cppgoBase_mousePressEvent(QMouseEvent* __a0) { QTimeEdit::mousePressEvent(__a0); }
	void cppgoBase_paintEvent(QPaintEvent* __a0) { QTimeEdit::paintEvent(__a0); }
	void cppgoBase_initStyleOption(QStyleOptionSpinBox* __a0) { QTimeEdit::initStyleOption(__a0); }
	void cppgoBase_resizeEvent(QResizeEvent* __a0) { QTimeEdit::resizeEvent(__a0); }
	void cppgoBase_keyReleaseEvent(QKeyEvent* __a0) { QTimeEdit::keyReleaseEvent(__a0); }
	void cppgoBase_focusOutEvent(QFocusEvent* __a0) { QTimeEdit::focusOutEvent(__a0); }
	void cppgoBase_contextMenuEvent(QContextMenuEvent* __a0) { QTimeEdit::contextMenuEvent(__a0); }
	void cppgoBase_changeEvent(QEvent* __a0) { QTimeEdit::changeEvent(__a0); }
	void cppgoBase_closeEvent(QCloseEvent* __a0) { QTimeEdit::closeEvent(__a0); }
	void cppgoBase_hideEvent(QHideEvent* __a0) { QTimeEdit::hideEvent(__a0); }
	void cppgoBase_mouseReleaseEvent(QMouseEvent* __a0) { QTimeEdit::mouseReleaseEvent(__a0); }
	void cppgoBase_mouseMoveEvent(QMouseEvent* __a0) { QTimeEdit::mouseMoveEvent(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QTimeEdit::timerEvent(__a0); }
	void cppgoBase_showEvent(QShowEvent* __a0) { QTimeEdit::showEvent(__a0); }
	void cppgoBase_mouseDoubleClickEvent(QMouseEvent* __a0) { QTimeEdit::mouseDoubleClickEvent(__a0); }
	void cppgoBase_enterEvent(QEnterEvent* __a0) { QTimeEdit::enterEvent(__a0); }
	void cppgoBase_leaveEvent(QEvent* __a0) { QTimeEdit::leaveEvent(__a0); }
	void cppgoBase_moveEvent(QMoveEvent* __a0) { QTimeEdit::moveEvent(__a0); }
	void cppgoBase_tabletEvent(QTabletEvent* __a0) { QTimeEdit::tabletEvent(__a0); }
	void cppgoBase_actionEvent(QActionEvent* __a0) { QTimeEdit::actionEvent(__a0); }
	void cppgoBase_dragEnterEvent(QDragEnterEvent* __a0) { QTimeEdit::dragEnterEvent(__a0); }
	void cppgoBase_dragMoveEvent(QDragMoveEvent* __a0) { QTimeEdit::dragMoveEvent(__a0); }
	void cppgoBase_dragLeaveEvent(QDragLeaveEvent* __a0) { QTimeEdit::dragLeaveEvent(__a0); }
	void cppgoBase_dropEvent(QDropEvent* __a0) { QTimeEdit::dropEvent(__a0); }
	int cppgoBase_metric(int __a0) { return QTimeEdit::metric((QPaintDevice::PaintDeviceMetric)__a0); }
	void cppgoBase_initPainter(QPainter* __a0) { QTimeEdit::initPainter(__a0); }
	QPaintDevice* cppgoBase_redirected(QPoint* __a0) { return QTimeEdit::redirected(__a0); }
	QPainter* cppgoBase_sharedPainter() { return QTimeEdit::sharedPainter(); }
	void cppgoBase_inputMethodEvent(QInputMethodEvent* __a0) { QTimeEdit::inputMethodEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QTimeEdit::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QTimeEdit::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QTimeEdit::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QTimeEdit::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QTimeEdit::metaObject();
	}
	QSize sizeHint() const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 729690261457684101LL, &__h); if (__h) return *(QSize*)__r; }
		return QTimeEdit::sizeHint();
	}
	void clear() override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, -778567940572639070LL, &__h); if (__h) return; }
		QTimeEdit::clear();
	}
	void stepBy(int steps) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 7085718643589954634LL, steps, &__h); if (__h) return; }
		QTimeEdit::stepBy(steps);
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QTimeEdit::event(event);
	}
	void keyPressEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1023791977679793185LL, (void*)event, &__h); if (__h) return; }
		QTimeEdit::keyPressEvent(event);
	}
	void wheelEvent(QWheelEvent* event) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1919988340934511316LL, (void*)event, &__h); if (__h) return; }
		QTimeEdit::wheelEvent(event);
	}
	void focusInEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 975043767033552552LL, (void*)event, &__h); if (__h) return; }
		QTimeEdit::focusInEvent(event);
	}
	bool focusNextPrevChild(bool next) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 5898328517505784191LL, next, &__h); if (__h) return __r; }
		return QTimeEdit::focusNextPrevChild(next);
	}
	QValidator::State validate(QString& input, int& pos) const override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_int_voidp_intp) { int __r = ((int(*)(long long, long long, void*, int*, int*))cppgo_virt_int_voidp_intp)(__goId, 6784179835071232555LL, (void*)&input, (int*)&pos, &__h); if (__h) return (QValidator::State)(__r); }
		return QTimeEdit::validate(input, pos);
	}
	void fixup(QString& input) const override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4855206404746966715LL, (void*)&input, &__h); if (__h) return; }
		QTimeEdit::fixup(input);
	}
	QDateTime dateTimeFromText(const QString& text) const override {
		int __h = 0;
		QByteArray text_ba = text.toUtf8();
		if (__goOn[11] && cppgo_virt_voidp_constcharp_longlong) { void* __r = ((void*(*)(long long, long long, const char*, long long, int*))cppgo_virt_voidp_constcharp_longlong)(__goId, -50441425836614729LL, (char*)text_ba.constData(), text_ba.size(), &__h); if (__h) return *(QDateTime*)__r; }
		return QTimeEdit::dateTimeFromText(text);
	}
	QString textFromDateTime(const QDateTime& dt) const override {
		int __h = 0;
		if (__goOn[12] && cppgo_virtp_charp_voidp) {
			long long __len = 0;
			char* __s = ((char*(*)(long long, long long, void*, long long*, int*))cppgo_virtp_charp_voidp)(__goId, -3778148144965928937LL, (void*)&dt, &__len, &__h);
			if (__h) { QString __r = QString::fromUtf8(__s, __len); free(__s); return __r; }
		}
		return QTimeEdit::textFromDateTime(dt);
	}
	QFlags<QAbstractSpinBox::StepEnabledFlag> stepEnabled() const override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, 5257289574516412244LL, &__h); if (__h) return QFlags<QAbstractSpinBox::StepEnabledFlag>(__r); }
		return QTimeEdit::stepEnabled();
	}
	void mousePressEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -332090898078543319LL, (void*)event, &__h); if (__h) return; }
		QTimeEdit::mousePressEvent(event);
	}
	void paintEvent(QPaintEvent* event) override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5290442163167458465LL, (void*)event, &__h); if (__h) return; }
		QTimeEdit::paintEvent(event);
	}
	void initStyleOption(QStyleOptionSpinBox* option) const override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -488693769771308531LL, (void*)option, &__h); if (__h) return; }
		QTimeEdit::initStyleOption(option);
	}
	QSize minimumSizeHint() const override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 7049264599837532491LL, &__h); if (__h) return *(QSize*)__r; }
		return QTimeEdit::minimumSizeHint();
	}
	QVariant inputMethodQuery(Qt::InputMethodQuery arg1) const override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7345853097340709638LL, (int)(arg1), &__h); if (__h) return *(QVariant*)__r; }
		return QTimeEdit::inputMethodQuery(arg1);
	}
	void resizeEvent(QResizeEvent* event) override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4302917285655014661LL, (void*)event, &__h); if (__h) return; }
		QTimeEdit::resizeEvent(event);
	}
	void keyReleaseEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 772598475950050299LL, (void*)event, &__h); if (__h) return; }
		QTimeEdit::keyReleaseEvent(event);
	}
	void focusOutEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -385596010576048993LL, (void*)event, &__h); if (__h) return; }
		QTimeEdit::focusOutEvent(event);
	}
	void contextMenuEvent(QContextMenuEvent* event) override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6698043222227006015LL, (void*)event, &__h); if (__h) return; }
		QTimeEdit::contextMenuEvent(event);
	}
	void changeEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4959376063826165943LL, (void*)event, &__h); if (__h) return; }
		QTimeEdit::changeEvent(event);
	}
	void closeEvent(QCloseEvent* event) override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1989937050019464783LL, (void*)event, &__h); if (__h) return; }
		QTimeEdit::closeEvent(event);
	}
	void hideEvent(QHideEvent* event) override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -199718707876719129LL, (void*)event, &__h); if (__h) return; }
		QTimeEdit::hideEvent(event);
	}
	void mouseReleaseEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4415920049297795213LL, (void*)event, &__h); if (__h) return; }
		QTimeEdit::mouseReleaseEvent(event);
	}
	void mouseMoveEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4255494004099806581LL, (void*)event, &__h); if (__h) return; }
		QTimeEdit::mouseMoveEvent(event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QTimeEdit::timerEvent(event);
	}
	void showEvent(QShowEvent* event) override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6825369174751916LL, (void*)event, &__h); if (__h) return; }
		QTimeEdit::showEvent(event);
	}
	int devType() const override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8659679893868524044LL, &__h); if (__h) return __r; }
		return QTimeEdit::devType();
	}
	void setVisible(bool visible) override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 2110182719375245575LL, visible, &__h); if (__h) return; }
		QTimeEdit::setVisible(visible);
	}
	int heightForWidth(int arg1) const override {
		int __h = 0;
		if (__goOn[32] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 4311534060495737339LL, arg1, &__h); if (__h) return __r; }
		return QTimeEdit::heightForWidth(arg1);
	}
	bool hasHeightForWidth() const override {
		int __h = 0;
		if (__goOn[33] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8266762397995423915LL, &__h); if (__h) return __r; }
		return QTimeEdit::hasHeightForWidth();
	}
	QPaintEngine* paintEngine() const override {
		int __h = 0;
		if (__goOn[34] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 6398007404756346381LL, &__h); if (__h) return (QPaintEngine*)__r; }
		return QTimeEdit::paintEngine();
	}
	void mouseDoubleClickEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[35] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4029073905697023887LL, (void*)event, &__h); if (__h) return; }
		QTimeEdit::mouseDoubleClickEvent(event);
	}
	void enterEvent(QEnterEvent* event) override {
		int __h = 0;
		if (__goOn[36] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6316270346668237895LL, (void*)event, &__h); if (__h) return; }
		QTimeEdit::enterEvent(event);
	}
	void leaveEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[37] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -551070009623711640LL, (void*)event, &__h); if (__h) return; }
		QTimeEdit::leaveEvent(event);
	}
	void moveEvent(QMoveEvent* event) override {
		int __h = 0;
		if (__goOn[38] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -9161949150489525396LL, (void*)event, &__h); if (__h) return; }
		QTimeEdit::moveEvent(event);
	}
	void tabletEvent(QTabletEvent* event) override {
		int __h = 0;
		if (__goOn[39] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -212068866860701141LL, (void*)event, &__h); if (__h) return; }
		QTimeEdit::tabletEvent(event);
	}
	void actionEvent(QActionEvent* event) override {
		int __h = 0;
		if (__goOn[40] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1697444140088555581LL, (void*)event, &__h); if (__h) return; }
		QTimeEdit::actionEvent(event);
	}
	void dragEnterEvent(QDragEnterEvent* event) override {
		int __h = 0;
		if (__goOn[41] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 206122715919917589LL, (void*)event, &__h); if (__h) return; }
		QTimeEdit::dragEnterEvent(event);
	}
	void dragMoveEvent(QDragMoveEvent* event) override {
		int __h = 0;
		if (__goOn[42] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1781545137352745978LL, (void*)event, &__h); if (__h) return; }
		QTimeEdit::dragMoveEvent(event);
	}
	void dragLeaveEvent(QDragLeaveEvent* event) override {
		int __h = 0;
		if (__goOn[43] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -3139904158373725006LL, (void*)event, &__h); if (__h) return; }
		QTimeEdit::dragLeaveEvent(event);
	}
	void dropEvent(QDropEvent* event) override {
		int __h = 0;
		if (__goOn[44] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 9107058682754591716LL, (void*)event, &__h); if (__h) return; }
		QTimeEdit::dropEvent(event);
	}
	bool nativeEvent(const QByteArray& eventType, void* message, qintptr* result) override {
		int __h = 0;
		if (__goOn[45] && cppgo_virt_int_voidp_voidp_longlongp) { int __r = ((int(*)(long long, long long, void*, void*, long long*, int*))cppgo_virt_int_voidp_voidp_longlongp)(__goId, -2039166586246443706LL, (void*)&eventType, (void*)message, (long long*)result, &__h); if (__h) return __r; }
		return QTimeEdit::nativeEvent(eventType, message, result);
	}
	int metric(QPaintDevice::PaintDeviceMetric arg1) const override {
		int __h = 0;
		if (__goOn[46] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 427629649267311777LL, (int)(arg1), &__h); if (__h) return __r; }
		return QTimeEdit::metric(arg1);
	}
	void initPainter(QPainter* painter) const override {
		int __h = 0;
		if (__goOn[47] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6477500296980050210LL, (void*)painter, &__h); if (__h) return; }
		QTimeEdit::initPainter(painter);
	}
	QPaintDevice* redirected(QPoint* offset) const override {
		int __h = 0;
		if (__goOn[48] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 4127240633310655048LL, (void*)offset, &__h); if (__h) return (QPaintDevice*)__r; }
		return QTimeEdit::redirected(offset);
	}
	QPainter* sharedPainter() const override {
		int __h = 0;
		if (__goOn[49] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 5792771328741774991LL, &__h); if (__h) return (QPainter*)__r; }
		return QTimeEdit::sharedPainter();
	}
	void inputMethodEvent(QInputMethodEvent* arg1) override {
		int __h = 0;
		if (__goOn[50] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5438819661942843222LL, (void*)arg1, &__h); if (__h) return; }
		QTimeEdit::inputMethodEvent(arg1);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[51] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QTimeEdit::eventFilter(watched, event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[52] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QTimeEdit::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[53] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QTimeEdit::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[54] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QTimeEdit::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[55] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QTimeEdit::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QTimeEdit;
static std::mutex __connsMu_QTimeEdit;

extern "C" {

void* QTimeEdit_base(void* self) { return (void*)static_cast<QDateTimeEdit*>((QTimeEdit*)self); }
void QTimeEdit_chain(void* self, void** out) {
	QTimeEdit* __p = (QTimeEdit*)self;
	QDateTimeEdit* __b0 = static_cast<QDateTimeEdit*>(__p); out[0] = (void*)__b0;
	QAbstractSpinBox* __b1 = static_cast<QAbstractSpinBox*>(__b0); out[1] = (void*)__b1;
	QWidget* __b2 = static_cast<QWidget*>(__b1); out[2] = (void*)__b2;
	QObject* __b3 = static_cast<QObject*>(__b2); out[3] = (void*)__b3;
}
void QTimeEdit__set_override(void* self, int idx) { ((QTimeEdit__Go*)self)->__goOn[idx] = 1; }
void* QTimeEdit_fromQObject(void* o) { return (void*)qobject_cast<QTimeEdit*>((QObject*)o); }
void* QTimeEdit_new(long long __goid, void* parent) {
	return new QTimeEdit__Go(__goid, (QWidget*)parent);
}
void* QTimeEdit_new2(long long __goid) {
	return new QTimeEdit__Go(__goid);
}
void* QTimeEdit_new3(long long __goid, void* time) {
	return new QTimeEdit__Go(__goid, *(QTime*)time);
}
void* QTimeEdit_new4(long long __goid, void* time, void* parent) {
	return new QTimeEdit__Go(__goid, *(QTime*)time, (QWidget*)parent);
}
char* QTimeEdit_tr(const char* s, long long* __out_len) {
	auto&& __r = QTimeEdit::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QTimeEdit_userTimeChanged(void* self, void* time) {
	((QTimeEdit*)self)->userTimeChanged(*(QTime*)time);
}
void* QTimeEdit_lineEdit(void* self) {
	auto&& __r = ((QTimeEdit__Go*)self)->lineEdit();
	return (void*)(__r);
}
void QTimeEdit_setLineEdit(void* self, void* edit) {
	((QTimeEdit__Go*)self)->setLineEdit((QLineEdit*)edit);
}
void QTimeEdit_setTabOrder(void* arg1, void* arg2) {
	QTimeEdit::setTabOrder((QWidget*)arg1, (QWidget*)arg2);
}
void* QTimeEdit_mouseGrabber() {
	auto&& __r = QTimeEdit::mouseGrabber();
	return (void*)(__r);
}
void* QTimeEdit_keyboardGrabber() {
	auto&& __r = QTimeEdit::keyboardGrabber();
	return (void*)(__r);
}
void* QTimeEdit_find(unsigned long long arg1) {
	auto&& __r = QTimeEdit::find(arg1);
	return (void*)(__r);
}
void* QTimeEdit_createWindowContainer(void* window) {
	auto&& __r = QTimeEdit::createWindowContainer((QWindow*)window);
	return (void*)(__r);
}
void QTimeEdit_updateMicroFocus(void* self) {
	((QTimeEdit__Go*)self)->updateMicroFocus();
}
void QTimeEdit_create(void* self) {
	((QTimeEdit__Go*)self)->create();
}
void QTimeEdit_destroy(void* self) {
	((QTimeEdit__Go*)self)->destroy();
}
int QTimeEdit_focusNextChild(void* self) {
	auto&& __r = ((QTimeEdit__Go*)self)->focusNextChild();
	return __r;
}
int QTimeEdit_focusPreviousChild(void* self) {
	auto&& __r = ((QTimeEdit__Go*)self)->focusPreviousChild();
	return __r;
}
void* QTimeEdit_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QTimeEdit::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QTimeEdit_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QTimeEdit::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QTimeEdit_sender(void* self) {
	auto&& __r = ((QTimeEdit__Go*)self)->sender();
	return (void*)(__r);
}
int QTimeEdit_senderSignalIndex(void* self) {
	auto&& __r = ((QTimeEdit__Go*)self)->senderSignalIndex();
	return __r;
}
int QTimeEdit_receivers(void* self, const char* signal) {
	auto&& __r = ((QTimeEdit__Go*)self)->receivers(signal);
	return __r;
}
int QTimeEdit_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QTimeEdit__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
double QTimeEdit_devicePixelRatioFScale() {
	auto&& __r = QTimeEdit::devicePixelRatioFScale();
	return __r;
}
int QTimeEdit_encodeMetricF(int metric, double value) {
	auto&& __r = QTimeEdit::encodeMetricF((QPaintDevice::PaintDeviceMetric)(metric), value);
	return __r;
}
double QTimeEdit_getDecodedMetricF(void* self, int metricA, int metricB) {
	auto&& __r = ((QTimeEdit__Go*)self)->getDecodedMetricF((QPaintDevice::PaintDeviceMetric)(metricA), (QPaintDevice::PaintDeviceMetric)(metricB));
	return __r;
}
char* QTimeEdit_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QTimeEdit::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QTimeEdit_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QTimeEdit::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QTimeEdit_createWindowContainer2(void* window, void* parent) {
	auto&& __r = QTimeEdit::createWindowContainer((QWindow*)window, (QWidget*)parent);
	return (void*)(__r);
}
void* QTimeEdit_createWindowContainer3(void* window, void* parent, int flags) {
	auto&& __r = QTimeEdit::createWindowContainer((QWindow*)window, (QWidget*)parent, QFlags<Qt::WindowType>(flags));
	return (void*)(__r);
}
void QTimeEdit_updateMicroFocus2(void* self, int query) {
	((QTimeEdit__Go*)self)->updateMicroFocus((Qt::InputMethodQuery)(query));
}
void QTimeEdit_create2(void* self, unsigned long long arg1) {
	((QTimeEdit__Go*)self)->create(arg1);
}
void QTimeEdit_create3(void* self, unsigned long long arg1, int initializeWindow) {
	((QTimeEdit__Go*)self)->create(arg1, initializeWindow);
}
void QTimeEdit_create4(void* self, unsigned long long arg1, int initializeWindow, int destroyOldWindow) {
	((QTimeEdit__Go*)self)->create(arg1, initializeWindow, destroyOldWindow);
}
void QTimeEdit_destroy2(void* self, int destroyWindow) {
	((QTimeEdit__Go*)self)->destroy(destroyWindow);
}
void QTimeEdit_destroy3(void* self, int destroyWindow, int destroySubWindows) {
	((QTimeEdit__Go*)self)->destroy(destroyWindow, destroySubWindows);
}
void* QTimeEdit_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QTimeEdit::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QTimeEdit_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QTimeEdit::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QTimeEdit_metaObject_default(void* self) {
	auto&& __r = ((QTimeEdit*)self)->QTimeEdit::metaObject();
	return (void*)(__r);
}
void* QTimeEdit_sizeHint_default(void* self) {
	auto&& __r = ((QTimeEdit*)self)->QTimeEdit::sizeHint();
	return new QSize(__r);
}
void QTimeEdit_clear_default(void* self) {
	((QTimeEdit*)self)->QTimeEdit::clear();
}
void QTimeEdit_stepBy_default(void* self, int steps) {
	((QTimeEdit*)self)->QTimeEdit::stepBy(steps);
}
int QTimeEdit_event_default(void* self, void* event) {
	auto&& __r = ((QTimeEdit*)self)->QTimeEdit::event((QEvent*)event);
	return __r;
}
void QTimeEdit_keyPressEvent_default(void* self, void* event) {
	((QTimeEdit__Go*)self)->cppgoBase_keyPressEvent((QKeyEvent*)event);
}
void QTimeEdit_wheelEvent_default(void* self, void* event) {
	((QTimeEdit__Go*)self)->cppgoBase_wheelEvent((QWheelEvent*)event);
}
void QTimeEdit_focusInEvent_default(void* self, void* event) {
	((QTimeEdit__Go*)self)->cppgoBase_focusInEvent((QFocusEvent*)event);
}
int QTimeEdit_focusNextPrevChild_default(void* self, int next) {
	auto&& __r = ((QTimeEdit__Go*)self)->cppgoBase_focusNextPrevChild(next);
	return __r;
}
int QTimeEdit_validate_default(void* self, void* input, int* pos) {
	int pos_ref = *pos;
	auto&& __r = ((QTimeEdit__Go*)self)->cppgoBase_validate(*(QString*)input, pos_ref);
	*pos = pos_ref;
	return (int)(__r);
}
void QTimeEdit_fixup_default(void* self, void* input) {
	((QTimeEdit__Go*)self)->cppgoBase_fixup(*(QString*)input);
}
void* QTimeEdit_dateTimeFromText_default(void* self, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = ((QTimeEdit__Go*)self)->cppgoBase_dateTimeFromText(text_qstr);
	return new QDateTime(__r);
}
char* QTimeEdit_textFromDateTime_default(void* self, void* dt, long long* __out_len) {
	auto&& __r = ((QTimeEdit__Go*)self)->cppgoBase_textFromDateTime(*(QDateTime*)dt);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QTimeEdit_stepEnabled_default(void* self) {
	auto&& __r = ((QTimeEdit__Go*)self)->cppgoBase_stepEnabled();
	return (int)(__r);
}
void QTimeEdit_mousePressEvent_default(void* self, void* event) {
	((QTimeEdit__Go*)self)->cppgoBase_mousePressEvent((QMouseEvent*)event);
}
void QTimeEdit_paintEvent_default(void* self, void* event) {
	((QTimeEdit__Go*)self)->cppgoBase_paintEvent((QPaintEvent*)event);
}
void QTimeEdit_initStyleOption_default(void* self, void* option) {
	((QTimeEdit__Go*)self)->cppgoBase_initStyleOption((QStyleOptionSpinBox*)option);
}
void* QTimeEdit_minimumSizeHint_default(void* self) {
	auto&& __r = ((QTimeEdit*)self)->QTimeEdit::minimumSizeHint();
	return new QSize(__r);
}
void* QTimeEdit_inputMethodQuery_default(void* self, int arg1) {
	auto&& __r = ((QTimeEdit*)self)->QTimeEdit::inputMethodQuery((Qt::InputMethodQuery)(arg1));
	return new QVariant(__r);
}
void QTimeEdit_resizeEvent_default(void* self, void* event) {
	((QTimeEdit__Go*)self)->cppgoBase_resizeEvent((QResizeEvent*)event);
}
void QTimeEdit_keyReleaseEvent_default(void* self, void* event) {
	((QTimeEdit__Go*)self)->cppgoBase_keyReleaseEvent((QKeyEvent*)event);
}
void QTimeEdit_focusOutEvent_default(void* self, void* event) {
	((QTimeEdit__Go*)self)->cppgoBase_focusOutEvent((QFocusEvent*)event);
}
void QTimeEdit_contextMenuEvent_default(void* self, void* event) {
	((QTimeEdit__Go*)self)->cppgoBase_contextMenuEvent((QContextMenuEvent*)event);
}
void QTimeEdit_changeEvent_default(void* self, void* event) {
	((QTimeEdit__Go*)self)->cppgoBase_changeEvent((QEvent*)event);
}
void QTimeEdit_closeEvent_default(void* self, void* event) {
	((QTimeEdit__Go*)self)->cppgoBase_closeEvent((QCloseEvent*)event);
}
void QTimeEdit_hideEvent_default(void* self, void* event) {
	((QTimeEdit__Go*)self)->cppgoBase_hideEvent((QHideEvent*)event);
}
void QTimeEdit_mouseReleaseEvent_default(void* self, void* event) {
	((QTimeEdit__Go*)self)->cppgoBase_mouseReleaseEvent((QMouseEvent*)event);
}
void QTimeEdit_mouseMoveEvent_default(void* self, void* event) {
	((QTimeEdit__Go*)self)->cppgoBase_mouseMoveEvent((QMouseEvent*)event);
}
void QTimeEdit_timerEvent_default(void* self, void* event) {
	((QTimeEdit__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QTimeEdit_showEvent_default(void* self, void* event) {
	((QTimeEdit__Go*)self)->cppgoBase_showEvent((QShowEvent*)event);
}
int QTimeEdit_devType_default(void* self) {
	auto&& __r = ((QTimeEdit*)self)->QTimeEdit::devType();
	return __r;
}
void QTimeEdit_setVisible_default(void* self, int visible) {
	((QTimeEdit*)self)->QTimeEdit::setVisible(visible);
}
int QTimeEdit_heightForWidth_default(void* self, int arg1) {
	auto&& __r = ((QTimeEdit*)self)->QTimeEdit::heightForWidth(arg1);
	return __r;
}
int QTimeEdit_hasHeightForWidth_default(void* self) {
	auto&& __r = ((QTimeEdit*)self)->QTimeEdit::hasHeightForWidth();
	return __r;
}
void* QTimeEdit_paintEngine_default(void* self) {
	auto&& __r = ((QTimeEdit*)self)->QTimeEdit::paintEngine();
	return (void*)(__r);
}
void QTimeEdit_mouseDoubleClickEvent_default(void* self, void* event) {
	((QTimeEdit__Go*)self)->cppgoBase_mouseDoubleClickEvent((QMouseEvent*)event);
}
void QTimeEdit_enterEvent_default(void* self, void* event) {
	((QTimeEdit__Go*)self)->cppgoBase_enterEvent((QEnterEvent*)event);
}
void QTimeEdit_leaveEvent_default(void* self, void* event) {
	((QTimeEdit__Go*)self)->cppgoBase_leaveEvent((QEvent*)event);
}
void QTimeEdit_moveEvent_default(void* self, void* event) {
	((QTimeEdit__Go*)self)->cppgoBase_moveEvent((QMoveEvent*)event);
}
void QTimeEdit_tabletEvent_default(void* self, void* event) {
	((QTimeEdit__Go*)self)->cppgoBase_tabletEvent((QTabletEvent*)event);
}
void QTimeEdit_actionEvent_default(void* self, void* event) {
	((QTimeEdit__Go*)self)->cppgoBase_actionEvent((QActionEvent*)event);
}
void QTimeEdit_dragEnterEvent_default(void* self, void* event) {
	((QTimeEdit__Go*)self)->cppgoBase_dragEnterEvent((QDragEnterEvent*)event);
}
void QTimeEdit_dragMoveEvent_default(void* self, void* event) {
	((QTimeEdit__Go*)self)->cppgoBase_dragMoveEvent((QDragMoveEvent*)event);
}
void QTimeEdit_dragLeaveEvent_default(void* self, void* event) {
	((QTimeEdit__Go*)self)->cppgoBase_dragLeaveEvent((QDragLeaveEvent*)event);
}
void QTimeEdit_dropEvent_default(void* self, void* event) {
	((QTimeEdit__Go*)self)->cppgoBase_dropEvent((QDropEvent*)event);
}
int QTimeEdit_metric_default(void* self, int arg1) {
	auto&& __r = ((QTimeEdit__Go*)self)->cppgoBase_metric(arg1);
	return __r;
}
void QTimeEdit_initPainter_default(void* self, void* painter) {
	((QTimeEdit__Go*)self)->cppgoBase_initPainter((QPainter*)painter);
}
void* QTimeEdit_redirected_default(void* self, void* offset) {
	auto&& __r = ((QTimeEdit__Go*)self)->cppgoBase_redirected((QPoint*)offset);
	return (void*)(__r);
}
void* QTimeEdit_sharedPainter_default(void* self) {
	auto&& __r = ((QTimeEdit__Go*)self)->cppgoBase_sharedPainter();
	return (void*)(__r);
}
void QTimeEdit_inputMethodEvent_default(void* self, void* arg1) {
	((QTimeEdit__Go*)self)->cppgoBase_inputMethodEvent((QInputMethodEvent*)arg1);
}
int QTimeEdit_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QTimeEdit*)self)->QTimeEdit::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QTimeEdit_childEvent_default(void* self, void* event) {
	((QTimeEdit__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QTimeEdit_customEvent_default(void* self, void* event) {
	((QTimeEdit__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QTimeEdit_connectNotify_default(void* self, void* signal) {
	((QTimeEdit__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QTimeEdit_disconnectNotify_default(void* self, void* signal) {
	((QTimeEdit__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QTimeEdit_connect_userTimeChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTimeEdit*)self, &QTimeEdit::userTimeChanged, (QTimeEdit*)self, [__id](QTime time){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&time);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTimeEdit);
	__conns_QTimeEdit[__id] = __c;
}
void QTimeEdit_connect_dateTimeChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTimeEdit*)self, &QTimeEdit::dateTimeChanged, (QTimeEdit*)self, [__id](const QDateTime& dateTime){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&dateTime);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTimeEdit);
	__conns_QTimeEdit[__id] = __c;
}
void QTimeEdit_connect_timeChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTimeEdit*)self, &QTimeEdit::timeChanged, (QTimeEdit*)self, [__id](QTime time){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&time);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTimeEdit);
	__conns_QTimeEdit[__id] = __c;
}
void QTimeEdit_connect_dateChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTimeEdit*)self, &QTimeEdit::dateChanged, (QTimeEdit*)self, [__id](QDate date){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&date);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTimeEdit);
	__conns_QTimeEdit[__id] = __c;
}
void QTimeEdit_connect_editingFinished(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTimeEdit*)self, &QTimeEdit::editingFinished, (QTimeEdit*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTimeEdit);
	__conns_QTimeEdit[__id] = __c;
}
void QTimeEdit_connect_returnPressed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTimeEdit*)self, &QTimeEdit::returnPressed, (QTimeEdit*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTimeEdit);
	__conns_QTimeEdit[__id] = __c;
}
void QTimeEdit_connect_windowTitleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTimeEdit*)self, &QTimeEdit::windowTitleChanged, (QTimeEdit*)self, [__id](const QString& title){
		QByteArray title_ba = title.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)title_ba.constData(), title_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTimeEdit);
	__conns_QTimeEdit[__id] = __c;
}
void QTimeEdit_connect_windowIconChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTimeEdit*)self, &QTimeEdit::windowIconChanged, (QTimeEdit*)self, [__id](const QIcon& icon){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&icon);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTimeEdit);
	__conns_QTimeEdit[__id] = __c;
}
void QTimeEdit_connect_windowIconTextChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTimeEdit*)self, &QTimeEdit::windowIconTextChanged, (QTimeEdit*)self, [__id](const QString& iconText){
		QByteArray iconText_ba = iconText.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)iconText_ba.constData(), iconText_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTimeEdit);
	__conns_QTimeEdit[__id] = __c;
}
void QTimeEdit_connect_customContextMenuRequested(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTimeEdit*)self, &QTimeEdit::customContextMenuRequested, (QTimeEdit*)self, [__id](const QPoint& pos){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&pos);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTimeEdit);
	__conns_QTimeEdit[__id] = __c;
}
void QTimeEdit_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTimeEdit*)self, &QTimeEdit::destroyed, (QTimeEdit*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTimeEdit);
	__conns_QTimeEdit[__id] = __c;
}
void QTimeEdit_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTimeEdit*)self, &QTimeEdit::destroyed, (QTimeEdit*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTimeEdit);
	__conns_QTimeEdit[__id] = __c;
}
void QTimeEdit_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTimeEdit*)self, &QTimeEdit::objectNameChanged, (QTimeEdit*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTimeEdit);
	__conns_QTimeEdit[__id] = __c;
}
void QTimeEdit__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QTimeEdit);
		auto __it = __conns_QTimeEdit.find(id);
		if (__it == __conns_QTimeEdit.end()) return;
		__c = __it->second;
		__conns_QTimeEdit.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QTimeEdit__Go_new(long long id) { return new QTimeEdit__Go(id); }
void QTimeEdit_delete(void* self) { delete (QTimeEdit*)self; }

}

#include <QtWidgets/qgraphicsitem.h>
#include <QtGui/qpen.h>
#include <QtGui/qbrush.h>
#include <QtGui/qpainterpath.h>
#include <QtWidgets/qgraphicsscene.h>
#include <QtWidgets/qgraphicswidget.h>
#include <QtCore/qsize.h>
#include <QtCore/qstring.h>
#include <QtGui/qcursor.h>
#include <QtWidgets/qgraphicseffect.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qpoint.h>
#include <QtCore/qrect.h>
#include <QtGui/qtransform.h>
#include <QtWidgets/qgraphicstransform.h>
#include <QtGui/qregion.h>
#include <QtGui/qpainter.h>
#include <QtWidgets/qstyleoption.h>
#include <QtWidgets/qwidget.h>
#include <QtGui/qpolygon.h>
#include <QtCore/qvariant.h>
#include <QtCore/qcoreevent.h>
#include <QtWidgets/qgraphicssceneevent.h>
#include <QtGui/qevent.h>
#include "gen_qabstractgraphicsshapeitem.h"
#include "gen_purego_cb.h"

class QAbstractGraphicsShapeItem__Go : public QAbstractGraphicsShapeItem {
public:
	long long __goId;
	unsigned char __goOn[35] = {};
	template<typename... A> QAbstractGraphicsShapeItem__Go(long long id, A&&... a) : QAbstractGraphicsShapeItem(std::forward<A>(a)...), __goId(id) {}
	using QAbstractGraphicsShapeItem::updateMicroFocus;
	using QAbstractGraphicsShapeItem::addToIndex;
	using QAbstractGraphicsShapeItem::removeFromIndex;
	using QAbstractGraphicsShapeItem::prepareGeometryChange;
	bool cppgoBase_sceneEventFilter(QGraphicsItem* __a0, QEvent* __a1) { return QAbstractGraphicsShapeItem::sceneEventFilter(__a0, __a1); }
	bool cppgoBase_sceneEvent(QEvent* __a0) { return QAbstractGraphicsShapeItem::sceneEvent(__a0); }
	void cppgoBase_contextMenuEvent(QGraphicsSceneContextMenuEvent* __a0) { QAbstractGraphicsShapeItem::contextMenuEvent(__a0); }
	void cppgoBase_dragEnterEvent(QGraphicsSceneDragDropEvent* __a0) { QAbstractGraphicsShapeItem::dragEnterEvent(__a0); }
	void cppgoBase_dragLeaveEvent(QGraphicsSceneDragDropEvent* __a0) { QAbstractGraphicsShapeItem::dragLeaveEvent(__a0); }
	void cppgoBase_dragMoveEvent(QGraphicsSceneDragDropEvent* __a0) { QAbstractGraphicsShapeItem::dragMoveEvent(__a0); }
	void cppgoBase_dropEvent(QGraphicsSceneDragDropEvent* __a0) { QAbstractGraphicsShapeItem::dropEvent(__a0); }
	void cppgoBase_focusInEvent(QFocusEvent* __a0) { QAbstractGraphicsShapeItem::focusInEvent(__a0); }
	void cppgoBase_focusOutEvent(QFocusEvent* __a0) { QAbstractGraphicsShapeItem::focusOutEvent(__a0); }
	void cppgoBase_hoverEnterEvent(QGraphicsSceneHoverEvent* __a0) { QAbstractGraphicsShapeItem::hoverEnterEvent(__a0); }
	void cppgoBase_hoverMoveEvent(QGraphicsSceneHoverEvent* __a0) { QAbstractGraphicsShapeItem::hoverMoveEvent(__a0); }
	void cppgoBase_hoverLeaveEvent(QGraphicsSceneHoverEvent* __a0) { QAbstractGraphicsShapeItem::hoverLeaveEvent(__a0); }
	void cppgoBase_keyPressEvent(QKeyEvent* __a0) { QAbstractGraphicsShapeItem::keyPressEvent(__a0); }
	void cppgoBase_keyReleaseEvent(QKeyEvent* __a0) { QAbstractGraphicsShapeItem::keyReleaseEvent(__a0); }
	void cppgoBase_mousePressEvent(QGraphicsSceneMouseEvent* __a0) { QAbstractGraphicsShapeItem::mousePressEvent(__a0); }
	void cppgoBase_mouseMoveEvent(QGraphicsSceneMouseEvent* __a0) { QAbstractGraphicsShapeItem::mouseMoveEvent(__a0); }
	void cppgoBase_mouseReleaseEvent(QGraphicsSceneMouseEvent* __a0) { QAbstractGraphicsShapeItem::mouseReleaseEvent(__a0); }
	void cppgoBase_mouseDoubleClickEvent(QGraphicsSceneMouseEvent* __a0) { QAbstractGraphicsShapeItem::mouseDoubleClickEvent(__a0); }
	void cppgoBase_wheelEvent(QGraphicsSceneWheelEvent* __a0) { QAbstractGraphicsShapeItem::wheelEvent(__a0); }
	void cppgoBase_inputMethodEvent(QInputMethodEvent* __a0) { QAbstractGraphicsShapeItem::inputMethodEvent(__a0); }
	QVariant cppgoBase_inputMethodQuery(int __a0) { return QAbstractGraphicsShapeItem::inputMethodQuery((Qt::InputMethodQuery)__a0); }
	QVariant cppgoBase_itemChange(int __a0, const QVariant& __a1) { return QAbstractGraphicsShapeItem::itemChange((QGraphicsItem::GraphicsItemChange)__a0, __a1); }
	bool cppgoBase_supportsExtension(int __a0) { return QAbstractGraphicsShapeItem::supportsExtension((QGraphicsItem::Extension)__a0); }
	void cppgoBase_setExtension(int __a0, const QVariant& __a1) { QAbstractGraphicsShapeItem::setExtension((QGraphicsItem::Extension)__a0, __a1); }
	QVariant cppgoBase_extension(const QVariant& __a0) { return QAbstractGraphicsShapeItem::extension(__a0); }
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
	bool isObscuredBy(const QGraphicsItem* item) const override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, -7215030271914903671LL, (void*)item, &__h); if (__h) return __r; }
		return QAbstractGraphicsShapeItem::isObscuredBy(item);
	}
	QPainterPath opaqueArea() const override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 399370970023989435LL, &__h); if (__h) return *(QPainterPath*)__r; }
		return QAbstractGraphicsShapeItem::opaqueArea();
	}
	void advance(int phase) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, -4830836341480073475LL, phase, &__h); if (__h) return; }
		QAbstractGraphicsShapeItem::advance(phase);
	}
	QPainterPath shape() const override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 7938019069247444294LL, &__h); if (__h) return *(QPainterPath*)__r; }
		return QAbstractGraphicsShapeItem::shape();
	}
	bool contains(const QPointF& point) const override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 5645366511944179768LL, (void*)&point, &__h); if (__h) return __r; }
		return QAbstractGraphicsShapeItem::contains(point);
	}
	bool collidesWithItem(const QGraphicsItem* other, Qt::ItemSelectionMode mode) const override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_int_voidp_int) { int __r = ((int(*)(long long, long long, void*, int, int*))cppgo_virt_int_voidp_int)(__goId, 6469736184362593173LL, (void*)other, (int)(mode), &__h); if (__h) return __r; }
		return QAbstractGraphicsShapeItem::collidesWithItem(other, mode);
	}
	bool collidesWithPath(const QPainterPath& path, Qt::ItemSelectionMode mode) const override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_int_voidp_int) { int __r = ((int(*)(long long, long long, void*, int, int*))cppgo_virt_int_voidp_int)(__goId, 6636575182718753853LL, (void*)&path, (int)(mode), &__h); if (__h) return __r; }
		return QAbstractGraphicsShapeItem::collidesWithPath(path, mode);
	}
	int type() const override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -6371403872082420691LL, &__h); if (__h) return __r; }
		return QAbstractGraphicsShapeItem::type();
	}
	bool sceneEventFilter(QGraphicsItem* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, -4584203218241215725LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QAbstractGraphicsShapeItem::sceneEventFilter(watched, event);
	}
	bool sceneEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, -2846371507838587071LL, (void*)event, &__h); if (__h) return __r; }
		return QAbstractGraphicsShapeItem::sceneEvent(event);
	}
	void contextMenuEvent(QGraphicsSceneContextMenuEvent* event) override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6698043222227006015LL, (void*)event, &__h); if (__h) return; }
		QAbstractGraphicsShapeItem::contextMenuEvent(event);
	}
	void dragEnterEvent(QGraphicsSceneDragDropEvent* event) override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 206122715919917589LL, (void*)event, &__h); if (__h) return; }
		QAbstractGraphicsShapeItem::dragEnterEvent(event);
	}
	void dragLeaveEvent(QGraphicsSceneDragDropEvent* event) override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -3139904158373725006LL, (void*)event, &__h); if (__h) return; }
		QAbstractGraphicsShapeItem::dragLeaveEvent(event);
	}
	void dragMoveEvent(QGraphicsSceneDragDropEvent* event) override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1781545137352745978LL, (void*)event, &__h); if (__h) return; }
		QAbstractGraphicsShapeItem::dragMoveEvent(event);
	}
	void dropEvent(QGraphicsSceneDragDropEvent* event) override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 9107058682754591716LL, (void*)event, &__h); if (__h) return; }
		QAbstractGraphicsShapeItem::dropEvent(event);
	}
	void focusInEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 975043767033552552LL, (void*)event, &__h); if (__h) return; }
		QAbstractGraphicsShapeItem::focusInEvent(event);
	}
	void focusOutEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -385596010576048993LL, (void*)event, &__h); if (__h) return; }
		QAbstractGraphicsShapeItem::focusOutEvent(event);
	}
	void hoverEnterEvent(QGraphicsSceneHoverEvent* event) override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6932941943906681679LL, (void*)event, &__h); if (__h) return; }
		QAbstractGraphicsShapeItem::hoverEnterEvent(event);
	}
	void hoverMoveEvent(QGraphicsSceneHoverEvent* event) override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -2745191522144316396LL, (void*)event, &__h); if (__h) return; }
		QAbstractGraphicsShapeItem::hoverMoveEvent(event);
	}
	void hoverLeaveEvent(QGraphicsSceneHoverEvent* event) override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7944408482203603984LL, (void*)event, &__h); if (__h) return; }
		QAbstractGraphicsShapeItem::hoverLeaveEvent(event);
	}
	void keyPressEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1023791977679793185LL, (void*)event, &__h); if (__h) return; }
		QAbstractGraphicsShapeItem::keyPressEvent(event);
	}
	void keyReleaseEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 772598475950050299LL, (void*)event, &__h); if (__h) return; }
		QAbstractGraphicsShapeItem::keyReleaseEvent(event);
	}
	void mousePressEvent(QGraphicsSceneMouseEvent* event) override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -332090898078543319LL, (void*)event, &__h); if (__h) return; }
		QAbstractGraphicsShapeItem::mousePressEvent(event);
	}
	void mouseMoveEvent(QGraphicsSceneMouseEvent* event) override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4255494004099806581LL, (void*)event, &__h); if (__h) return; }
		QAbstractGraphicsShapeItem::mouseMoveEvent(event);
	}
	void mouseReleaseEvent(QGraphicsSceneMouseEvent* event) override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4415920049297795213LL, (void*)event, &__h); if (__h) return; }
		QAbstractGraphicsShapeItem::mouseReleaseEvent(event);
	}
	void mouseDoubleClickEvent(QGraphicsSceneMouseEvent* event) override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4029073905697023887LL, (void*)event, &__h); if (__h) return; }
		QAbstractGraphicsShapeItem::mouseDoubleClickEvent(event);
	}
	void wheelEvent(QGraphicsSceneWheelEvent* event) override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1919988340934511316LL, (void*)event, &__h); if (__h) return; }
		QAbstractGraphicsShapeItem::wheelEvent(event);
	}
	void inputMethodEvent(QInputMethodEvent* event) override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5438819661942843222LL, (void*)event, &__h); if (__h) return; }
		QAbstractGraphicsShapeItem::inputMethodEvent(event);
	}
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7345853097340709638LL, (int)(query), &__h); if (__h) return *(QVariant*)__r; }
		return QAbstractGraphicsShapeItem::inputMethodQuery(query);
	}
	QVariant itemChange(QGraphicsItem::GraphicsItemChange change, const QVariant& value) override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_voidp_int_voidp) { void* __r = ((void*(*)(long long, long long, int, void*, int*))cppgo_virt_voidp_int_voidp)(__goId, -1888572168794878830LL, (int)(change), (void*)&value, &__h); if (__h) return *(QVariant*)__r; }
		return QAbstractGraphicsShapeItem::itemChange(change, value);
	}
	bool supportsExtension(QGraphicsItem::Extension extension) const override {
		int __h = 0;
		if (__goOn[32] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, -5729166443651306418LL, (int)(extension), &__h); if (__h) return __r; }
		return QAbstractGraphicsShapeItem::supportsExtension(extension);
	}
	void setExtension(QGraphicsItem::Extension extension, const QVariant& variant) override {
		int __h = 0;
		if (__goOn[33] && cppgo_virt_void_int_voidp) { ((void(*)(long long, long long, int, void*, int*))cppgo_virt_void_int_voidp)(__goId, -6730181576106171580LL, (int)(extension), (void*)&variant, &__h); if (__h) return; }
		QAbstractGraphicsShapeItem::setExtension(extension, variant);
	}
	QVariant extension(const QVariant& variant) const override {
		int __h = 0;
		if (__goOn[34] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 2270616358835251226LL, (void*)&variant, &__h); if (__h) return *(QVariant*)__r; }
		return QAbstractGraphicsShapeItem::extension(variant);
	}
};

extern "C" {

void* QAbstractGraphicsShapeItem_base(void* self) { return (void*)static_cast<QGraphicsItem*>((QAbstractGraphicsShapeItem*)self); }
void QAbstractGraphicsShapeItem__set_override(void* self, int idx) { ((QAbstractGraphicsShapeItem__Go*)self)->__goOn[idx] = 1; }
void* QAbstractGraphicsShapeItem_pen(void* self) {
	auto&& __r = ((QAbstractGraphicsShapeItem*)self)->pen();
	return new QPen(__r);
}
void QAbstractGraphicsShapeItem_setPen(void* self, void* pen) {
	((QAbstractGraphicsShapeItem*)self)->setPen(*(QPen*)pen);
}
void* QAbstractGraphicsShapeItem_brush(void* self) {
	auto&& __r = ((QAbstractGraphicsShapeItem*)self)->brush();
	return new QBrush(__r);
}
void QAbstractGraphicsShapeItem_setBrush(void* self, void* brush) {
	((QAbstractGraphicsShapeItem*)self)->setBrush(*(QBrush*)brush);
}
void QAbstractGraphicsShapeItem_updateMicroFocus(void* self) {
	((QAbstractGraphicsShapeItem__Go*)self)->updateMicroFocus();
}
void QAbstractGraphicsShapeItem_addToIndex(void* self) {
	((QAbstractGraphicsShapeItem__Go*)self)->addToIndex();
}
void QAbstractGraphicsShapeItem_removeFromIndex(void* self) {
	((QAbstractGraphicsShapeItem__Go*)self)->removeFromIndex();
}
void QAbstractGraphicsShapeItem_prepareGeometryChange(void* self) {
	((QAbstractGraphicsShapeItem__Go*)self)->prepareGeometryChange();
}
int QAbstractGraphicsShapeItem_isObscuredBy_default(void* self, void* item) {
	auto&& __r = ((QAbstractGraphicsShapeItem*)self)->QAbstractGraphicsShapeItem::isObscuredBy((QGraphicsItem*)item);
	return __r;
}
void* QAbstractGraphicsShapeItem_opaqueArea_default(void* self) {
	auto&& __r = ((QAbstractGraphicsShapeItem*)self)->QAbstractGraphicsShapeItem::opaqueArea();
	return new QPainterPath(__r);
}
void QAbstractGraphicsShapeItem_advance_default(void* self, int phase) {
	((QAbstractGraphicsShapeItem*)self)->QAbstractGraphicsShapeItem::advance(phase);
}
void* QAbstractGraphicsShapeItem_shape_default(void* self) {
	auto&& __r = ((QAbstractGraphicsShapeItem*)self)->QAbstractGraphicsShapeItem::shape();
	return new QPainterPath(__r);
}
int QAbstractGraphicsShapeItem_contains_default(void* self, void* point) {
	auto&& __r = ((QAbstractGraphicsShapeItem*)self)->QAbstractGraphicsShapeItem::contains(*(QPointF*)point);
	return __r;
}
int QAbstractGraphicsShapeItem_collidesWithItem_default(void* self, void* other, int mode) {
	auto&& __r = ((QAbstractGraphicsShapeItem*)self)->QAbstractGraphicsShapeItem::collidesWithItem((QGraphicsItem*)other, (Qt::ItemSelectionMode)(mode));
	return __r;
}
int QAbstractGraphicsShapeItem_collidesWithPath_default(void* self, void* path, int mode) {
	auto&& __r = ((QAbstractGraphicsShapeItem*)self)->QAbstractGraphicsShapeItem::collidesWithPath(*(QPainterPath*)path, (Qt::ItemSelectionMode)(mode));
	return __r;
}
int QAbstractGraphicsShapeItem_type_default(void* self) {
	auto&& __r = ((QAbstractGraphicsShapeItem*)self)->QAbstractGraphicsShapeItem::type();
	return __r;
}
int QAbstractGraphicsShapeItem_sceneEventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QAbstractGraphicsShapeItem__Go*)self)->cppgoBase_sceneEventFilter((QGraphicsItem*)watched, (QEvent*)event);
	return __r;
}
int QAbstractGraphicsShapeItem_sceneEvent_default(void* self, void* event) {
	auto&& __r = ((QAbstractGraphicsShapeItem__Go*)self)->cppgoBase_sceneEvent((QEvent*)event);
	return __r;
}
void QAbstractGraphicsShapeItem_contextMenuEvent_default(void* self, void* event) {
	((QAbstractGraphicsShapeItem__Go*)self)->cppgoBase_contextMenuEvent((QGraphicsSceneContextMenuEvent*)event);
}
void QAbstractGraphicsShapeItem_dragEnterEvent_default(void* self, void* event) {
	((QAbstractGraphicsShapeItem__Go*)self)->cppgoBase_dragEnterEvent((QGraphicsSceneDragDropEvent*)event);
}
void QAbstractGraphicsShapeItem_dragLeaveEvent_default(void* self, void* event) {
	((QAbstractGraphicsShapeItem__Go*)self)->cppgoBase_dragLeaveEvent((QGraphicsSceneDragDropEvent*)event);
}
void QAbstractGraphicsShapeItem_dragMoveEvent_default(void* self, void* event) {
	((QAbstractGraphicsShapeItem__Go*)self)->cppgoBase_dragMoveEvent((QGraphicsSceneDragDropEvent*)event);
}
void QAbstractGraphicsShapeItem_dropEvent_default(void* self, void* event) {
	((QAbstractGraphicsShapeItem__Go*)self)->cppgoBase_dropEvent((QGraphicsSceneDragDropEvent*)event);
}
void QAbstractGraphicsShapeItem_focusInEvent_default(void* self, void* event) {
	((QAbstractGraphicsShapeItem__Go*)self)->cppgoBase_focusInEvent((QFocusEvent*)event);
}
void QAbstractGraphicsShapeItem_focusOutEvent_default(void* self, void* event) {
	((QAbstractGraphicsShapeItem__Go*)self)->cppgoBase_focusOutEvent((QFocusEvent*)event);
}
void QAbstractGraphicsShapeItem_hoverEnterEvent_default(void* self, void* event) {
	((QAbstractGraphicsShapeItem__Go*)self)->cppgoBase_hoverEnterEvent((QGraphicsSceneHoverEvent*)event);
}
void QAbstractGraphicsShapeItem_hoverMoveEvent_default(void* self, void* event) {
	((QAbstractGraphicsShapeItem__Go*)self)->cppgoBase_hoverMoveEvent((QGraphicsSceneHoverEvent*)event);
}
void QAbstractGraphicsShapeItem_hoverLeaveEvent_default(void* self, void* event) {
	((QAbstractGraphicsShapeItem__Go*)self)->cppgoBase_hoverLeaveEvent((QGraphicsSceneHoverEvent*)event);
}
void QAbstractGraphicsShapeItem_keyPressEvent_default(void* self, void* event) {
	((QAbstractGraphicsShapeItem__Go*)self)->cppgoBase_keyPressEvent((QKeyEvent*)event);
}
void QAbstractGraphicsShapeItem_keyReleaseEvent_default(void* self, void* event) {
	((QAbstractGraphicsShapeItem__Go*)self)->cppgoBase_keyReleaseEvent((QKeyEvent*)event);
}
void QAbstractGraphicsShapeItem_mousePressEvent_default(void* self, void* event) {
	((QAbstractGraphicsShapeItem__Go*)self)->cppgoBase_mousePressEvent((QGraphicsSceneMouseEvent*)event);
}
void QAbstractGraphicsShapeItem_mouseMoveEvent_default(void* self, void* event) {
	((QAbstractGraphicsShapeItem__Go*)self)->cppgoBase_mouseMoveEvent((QGraphicsSceneMouseEvent*)event);
}
void QAbstractGraphicsShapeItem_mouseReleaseEvent_default(void* self, void* event) {
	((QAbstractGraphicsShapeItem__Go*)self)->cppgoBase_mouseReleaseEvent((QGraphicsSceneMouseEvent*)event);
}
void QAbstractGraphicsShapeItem_mouseDoubleClickEvent_default(void* self, void* event) {
	((QAbstractGraphicsShapeItem__Go*)self)->cppgoBase_mouseDoubleClickEvent((QGraphicsSceneMouseEvent*)event);
}
void QAbstractGraphicsShapeItem_wheelEvent_default(void* self, void* event) {
	((QAbstractGraphicsShapeItem__Go*)self)->cppgoBase_wheelEvent((QGraphicsSceneWheelEvent*)event);
}
void QAbstractGraphicsShapeItem_inputMethodEvent_default(void* self, void* event) {
	((QAbstractGraphicsShapeItem__Go*)self)->cppgoBase_inputMethodEvent((QInputMethodEvent*)event);
}
void* QAbstractGraphicsShapeItem_inputMethodQuery_default(void* self, int query) {
	auto&& __r = ((QAbstractGraphicsShapeItem__Go*)self)->cppgoBase_inputMethodQuery(query);
	return new QVariant(__r);
}
void* QAbstractGraphicsShapeItem_itemChange_default(void* self, int change, void* value) {
	auto&& __r = ((QAbstractGraphicsShapeItem__Go*)self)->cppgoBase_itemChange(change, *(QVariant*)value);
	return new QVariant(__r);
}
int QAbstractGraphicsShapeItem_supportsExtension_default(void* self, int extension) {
	auto&& __r = ((QAbstractGraphicsShapeItem__Go*)self)->cppgoBase_supportsExtension(extension);
	return __r;
}
void QAbstractGraphicsShapeItem_setExtension_default(void* self, int extension, void* variant) {
	((QAbstractGraphicsShapeItem__Go*)self)->cppgoBase_setExtension(extension, *(QVariant*)variant);
}
void* QAbstractGraphicsShapeItem_extension_default(void* self, void* variant) {
	auto&& __r = ((QAbstractGraphicsShapeItem__Go*)self)->cppgoBase_extension(*(QVariant*)variant);
	return new QVariant(__r);
}
void* QAbstractGraphicsShapeItem__Go_new(long long id) { return new QAbstractGraphicsShapeItem__Go(id); }
void QAbstractGraphicsShapeItem_delete(void* self) { delete (QAbstractGraphicsShapeItem*)self; }

}

#include <QtWidgets/qgraphicsscene.h>
#include <QtCore/qobject.h>
#include <QtCore/qrect.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtGui/qpainter.h>
#include <QtCore/qnamespace.h>
#include <QtWidgets/qgraphicsitem.h>
#include <QtCore/qpoint.h>
#include <QtGui/qtransform.h>
#include <QtGui/qpolygon.h>
#include <QtGui/qpainterpath.h>
#include <QtGui/qpen.h>
#include <QtGui/qbrush.h>
#include <QtCore/qline.h>
#include <QtGui/qpixmap.h>
#include <QtGui/qfont.h>
#include <QtWidgets/qgraphicsproxywidget.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qvariant.h>
#include <QtWidgets/qgraphicsview.h>
#include <QtWidgets/qstyle.h>
#include <QtGui/qpalette.h>
#include <QtWidgets/qgraphicswidget.h>
#include <QtCore/qcoreevent.h>
#include <QtWidgets/qgraphicssceneevent.h>
#include <QtGui/qevent.h>
#include <QtWidgets/qstyleoption.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qgraphicsscene.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QGraphicsScene__Go : public QGraphicsScene {
public:
	long long __goId;
	unsigned char __goOn[28] = {};
	template<typename... A> QGraphicsScene__Go(long long id, A&&... a) : QGraphicsScene(std::forward<A>(a)...), __goId(id) {}
	using QGraphicsScene::sender;
	using QGraphicsScene::senderSignalIndex;
	using QGraphicsScene::receivers;
	using QGraphicsScene::isSignalConnected;
	bool cppgoBase_event(QEvent* __a0) { return QGraphicsScene::event(__a0); }
	bool cppgoBase_eventFilter(QObject* __a0, QEvent* __a1) { return QGraphicsScene::eventFilter(__a0, __a1); }
	void cppgoBase_contextMenuEvent(QGraphicsSceneContextMenuEvent* __a0) { QGraphicsScene::contextMenuEvent(__a0); }
	void cppgoBase_dragEnterEvent(QGraphicsSceneDragDropEvent* __a0) { QGraphicsScene::dragEnterEvent(__a0); }
	void cppgoBase_dragMoveEvent(QGraphicsSceneDragDropEvent* __a0) { QGraphicsScene::dragMoveEvent(__a0); }
	void cppgoBase_dragLeaveEvent(QGraphicsSceneDragDropEvent* __a0) { QGraphicsScene::dragLeaveEvent(__a0); }
	void cppgoBase_dropEvent(QGraphicsSceneDragDropEvent* __a0) { QGraphicsScene::dropEvent(__a0); }
	void cppgoBase_focusInEvent(QFocusEvent* __a0) { QGraphicsScene::focusInEvent(__a0); }
	void cppgoBase_focusOutEvent(QFocusEvent* __a0) { QGraphicsScene::focusOutEvent(__a0); }
	void cppgoBase_helpEvent(QGraphicsSceneHelpEvent* __a0) { QGraphicsScene::helpEvent(__a0); }
	void cppgoBase_keyPressEvent(QKeyEvent* __a0) { QGraphicsScene::keyPressEvent(__a0); }
	void cppgoBase_keyReleaseEvent(QKeyEvent* __a0) { QGraphicsScene::keyReleaseEvent(__a0); }
	void cppgoBase_mousePressEvent(QGraphicsSceneMouseEvent* __a0) { QGraphicsScene::mousePressEvent(__a0); }
	void cppgoBase_mouseMoveEvent(QGraphicsSceneMouseEvent* __a0) { QGraphicsScene::mouseMoveEvent(__a0); }
	void cppgoBase_mouseReleaseEvent(QGraphicsSceneMouseEvent* __a0) { QGraphicsScene::mouseReleaseEvent(__a0); }
	void cppgoBase_mouseDoubleClickEvent(QGraphicsSceneMouseEvent* __a0) { QGraphicsScene::mouseDoubleClickEvent(__a0); }
	void cppgoBase_wheelEvent(QGraphicsSceneWheelEvent* __a0) { QGraphicsScene::wheelEvent(__a0); }
	void cppgoBase_inputMethodEvent(QInputMethodEvent* __a0) { QGraphicsScene::inputMethodEvent(__a0); }
	void cppgoBase_drawBackground(QPainter* __a0, const QRectF& __a1) { QGraphicsScene::drawBackground(__a0, __a1); }
	void cppgoBase_drawForeground(QPainter* __a0, const QRectF& __a1) { QGraphicsScene::drawForeground(__a0, __a1); }
	bool cppgoBase_focusNextPrevChild(bool __a0) { return QGraphicsScene::focusNextPrevChild(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QGraphicsScene::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QGraphicsScene::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QGraphicsScene::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QGraphicsScene::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QGraphicsScene::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QGraphicsScene::metaObject();
	}
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7345853097340709638LL, (int)(query), &__h); if (__h) return *(QVariant*)__r; }
		return QGraphicsScene::inputMethodQuery(query);
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QGraphicsScene::event(event);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QGraphicsScene::eventFilter(watched, event);
	}
	void contextMenuEvent(QGraphicsSceneContextMenuEvent* event) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6698043222227006015LL, (void*)event, &__h); if (__h) return; }
		QGraphicsScene::contextMenuEvent(event);
	}
	void dragEnterEvent(QGraphicsSceneDragDropEvent* event) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 206122715919917589LL, (void*)event, &__h); if (__h) return; }
		QGraphicsScene::dragEnterEvent(event);
	}
	void dragMoveEvent(QGraphicsSceneDragDropEvent* event) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1781545137352745978LL, (void*)event, &__h); if (__h) return; }
		QGraphicsScene::dragMoveEvent(event);
	}
	void dragLeaveEvent(QGraphicsSceneDragDropEvent* event) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -3139904158373725006LL, (void*)event, &__h); if (__h) return; }
		QGraphicsScene::dragLeaveEvent(event);
	}
	void dropEvent(QGraphicsSceneDragDropEvent* event) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 9107058682754591716LL, (void*)event, &__h); if (__h) return; }
		QGraphicsScene::dropEvent(event);
	}
	void focusInEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 975043767033552552LL, (void*)event, &__h); if (__h) return; }
		QGraphicsScene::focusInEvent(event);
	}
	void focusOutEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -385596010576048993LL, (void*)event, &__h); if (__h) return; }
		QGraphicsScene::focusOutEvent(event);
	}
	void helpEvent(QGraphicsSceneHelpEvent* event) override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1696901855134526294LL, (void*)event, &__h); if (__h) return; }
		QGraphicsScene::helpEvent(event);
	}
	void keyPressEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1023791977679793185LL, (void*)event, &__h); if (__h) return; }
		QGraphicsScene::keyPressEvent(event);
	}
	void keyReleaseEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 772598475950050299LL, (void*)event, &__h); if (__h) return; }
		QGraphicsScene::keyReleaseEvent(event);
	}
	void mousePressEvent(QGraphicsSceneMouseEvent* event) override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -332090898078543319LL, (void*)event, &__h); if (__h) return; }
		QGraphicsScene::mousePressEvent(event);
	}
	void mouseMoveEvent(QGraphicsSceneMouseEvent* event) override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4255494004099806581LL, (void*)event, &__h); if (__h) return; }
		QGraphicsScene::mouseMoveEvent(event);
	}
	void mouseReleaseEvent(QGraphicsSceneMouseEvent* event) override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4415920049297795213LL, (void*)event, &__h); if (__h) return; }
		QGraphicsScene::mouseReleaseEvent(event);
	}
	void mouseDoubleClickEvent(QGraphicsSceneMouseEvent* event) override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4029073905697023887LL, (void*)event, &__h); if (__h) return; }
		QGraphicsScene::mouseDoubleClickEvent(event);
	}
	void wheelEvent(QGraphicsSceneWheelEvent* event) override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1919988340934511316LL, (void*)event, &__h); if (__h) return; }
		QGraphicsScene::wheelEvent(event);
	}
	void inputMethodEvent(QInputMethodEvent* event) override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5438819661942843222LL, (void*)event, &__h); if (__h) return; }
		QGraphicsScene::inputMethodEvent(event);
	}
	void drawBackground(QPainter* painter, const QRectF& rect) override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_void_voidp_voidp) { ((void(*)(long long, long long, void*, void*, int*))cppgo_virt_void_voidp_voidp)(__goId, -3076319991076358961LL, (void*)painter, (void*)&rect, &__h); if (__h) return; }
		QGraphicsScene::drawBackground(painter, rect);
	}
	void drawForeground(QPainter* painter, const QRectF& rect) override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_void_voidp_voidp) { ((void(*)(long long, long long, void*, void*, int*))cppgo_virt_void_voidp_voidp)(__goId, 8267095530709235866LL, (void*)painter, (void*)&rect, &__h); if (__h) return; }
		QGraphicsScene::drawForeground(painter, rect);
	}
	bool focusNextPrevChild(bool next) override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 5898328517505784191LL, next, &__h); if (__h) return __r; }
		return QGraphicsScene::focusNextPrevChild(next);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QGraphicsScene::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QGraphicsScene::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QGraphicsScene::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QGraphicsScene::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QGraphicsScene::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QGraphicsScene;
static std::mutex __connsMu_QGraphicsScene;

extern "C" {

void* QGraphicsScene_base(void* self) { return (void*)static_cast<QObject*>((QGraphicsScene*)self); }
void QGraphicsScene__set_override(void* self, int idx) { ((QGraphicsScene__Go*)self)->__goOn[idx] = 1; }
void* QGraphicsScene_fromQObject(void* o) { return (void*)qobject_cast<QGraphicsScene*>((QObject*)o); }
void* QGraphicsScene_new(long long __goid) {
	return new QGraphicsScene__Go(__goid);
}
void* QGraphicsScene_new2(long long __goid, void* sceneRect) {
	return new QGraphicsScene__Go(__goid, *(QRectF*)sceneRect);
}
void* QGraphicsScene_new3(long long __goid, double x, double y, double width, double height) {
	return new QGraphicsScene__Go(__goid, x, y, width, height);
}
void* QGraphicsScene_new4(long long __goid, void* parent) {
	return new QGraphicsScene__Go(__goid, (QObject*)parent);
}
void* QGraphicsScene_new5(long long __goid, void* sceneRect, void* parent) {
	return new QGraphicsScene__Go(__goid, *(QRectF*)sceneRect, (QObject*)parent);
}
void* QGraphicsScene_new6(long long __goid, double x, double y, double width, double height, void* parent) {
	return new QGraphicsScene__Go(__goid, x, y, width, height, (QObject*)parent);
}
char* QGraphicsScene_tr(const char* s, long long* __out_len) {
	auto&& __r = QGraphicsScene::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QGraphicsScene_sceneRect(void* self) {
	auto&& __r = ((QGraphicsScene*)self)->sceneRect();
	return new QRectF(__r);
}
double QGraphicsScene_width(void* self) {
	auto&& __r = ((QGraphicsScene*)self)->width();
	return __r;
}
double QGraphicsScene_height(void* self) {
	auto&& __r = ((QGraphicsScene*)self)->height();
	return __r;
}
void QGraphicsScene_setSceneRect(void* self, void* rect) {
	((QGraphicsScene*)self)->setSceneRect(*(QRectF*)rect);
}
void QGraphicsScene_setSceneRect2(void* self, double x, double y, double w, double h) {
	((QGraphicsScene*)self)->setSceneRect(x, y, w, h);
}
void QGraphicsScene_render(void* self, void* painter) {
	((QGraphicsScene*)self)->render((QPainter*)painter);
}
int QGraphicsScene_itemIndexMethod(void* self) {
	auto&& __r = ((QGraphicsScene*)self)->itemIndexMethod();
	return (int)(__r);
}
void QGraphicsScene_setItemIndexMethod(void* self, int method) {
	((QGraphicsScene*)self)->setItemIndexMethod((QGraphicsScene::ItemIndexMethod)(method));
}
int QGraphicsScene_bspTreeDepth(void* self) {
	auto&& __r = ((QGraphicsScene*)self)->bspTreeDepth();
	return __r;
}
void QGraphicsScene_setBspTreeDepth(void* self, int depth) {
	((QGraphicsScene*)self)->setBspTreeDepth(depth);
}
void* QGraphicsScene_itemsBoundingRect(void* self) {
	auto&& __r = ((QGraphicsScene*)self)->itemsBoundingRect();
	return new QRectF(__r);
}
void** QGraphicsScene_items(void* self, long long* __out_len) {
	auto&& __r = ((QGraphicsScene*)self)->items();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
void** QGraphicsScene_items2(void* self, void* pos, long long* __out_len) {
	auto&& __r = ((QGraphicsScene*)self)->items(*(QPointF*)pos);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
void** QGraphicsScene_items3(void* self, void* rect, long long* __out_len) {
	auto&& __r = ((QGraphicsScene*)self)->items(*(QRectF*)rect);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
void** QGraphicsScene_items4(void* self, void* polygon, long long* __out_len) {
	auto&& __r = ((QGraphicsScene*)self)->items(*(QPolygonF*)polygon);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
void** QGraphicsScene_items5(void* self, void* path, long long* __out_len) {
	auto&& __r = ((QGraphicsScene*)self)->items(*(QPainterPath*)path);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
void** QGraphicsScene_items6(void* self, double x, double y, double w, double h, int mode, int order, long long* __out_len) {
	auto&& __r = ((QGraphicsScene*)self)->items(x, y, w, h, (Qt::ItemSelectionMode)(mode), (Qt::SortOrder)(order));
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
void** QGraphicsScene_collidingItems(void* self, void* item, long long* __out_len) {
	auto&& __r = ((QGraphicsScene*)self)->collidingItems((QGraphicsItem*)item);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
void* QGraphicsScene_itemAt(void* self, void* pos, void* deviceTransform) {
	auto&& __r = ((QGraphicsScene*)self)->itemAt(*(QPointF*)pos, *(QTransform*)deviceTransform);
	return (void*)(__r);
}
void* QGraphicsScene_itemAt2(void* self, double x, double y, void* deviceTransform) {
	auto&& __r = ((QGraphicsScene*)self)->itemAt(x, y, *(QTransform*)deviceTransform);
	return (void*)(__r);
}
void** QGraphicsScene_selectedItems(void* self, long long* __out_len) {
	auto&& __r = ((QGraphicsScene*)self)->selectedItems();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
void* QGraphicsScene_selectionArea(void* self) {
	auto&& __r = ((QGraphicsScene*)self)->selectionArea();
	return new QPainterPath(__r);
}
void QGraphicsScene_setSelectionArea(void* self, void* path, void* deviceTransform) {
	((QGraphicsScene*)self)->setSelectionArea(*(QPainterPath*)path, *(QTransform*)deviceTransform);
}
void QGraphicsScene_setSelectionArea2(void* self, void* path) {
	((QGraphicsScene*)self)->setSelectionArea(*(QPainterPath*)path);
}
void* QGraphicsScene_createItemGroup(void* self, void** items, long long items_len) {
	QList<QGraphicsItem*> items_list; items_list.reserve(items_len);
	for (long long __i = 0; __i < items_len; __i++) items_list.append((QGraphicsItem*)items[__i]);
	auto&& __r = ((QGraphicsScene*)self)->createItemGroup(items_list);
	return (void*)(__r);
}
void QGraphicsScene_destroyItemGroup(void* self, void* group) {
	((QGraphicsScene*)self)->destroyItemGroup((QGraphicsItemGroup*)group);
}
void QGraphicsScene_addItem(void* self, void* item) {
	((QGraphicsScene*)self)->addItem((QGraphicsItem*)item);
}
void* QGraphicsScene_addEllipse(void* self, void* rect) {
	auto&& __r = ((QGraphicsScene*)self)->addEllipse(*(QRectF*)rect);
	return (void*)(__r);
}
void* QGraphicsScene_addLine(void* self, void* line) {
	auto&& __r = ((QGraphicsScene*)self)->addLine(*(QLineF*)line);
	return (void*)(__r);
}
void* QGraphicsScene_addPath(void* self, void* path) {
	auto&& __r = ((QGraphicsScene*)self)->addPath(*(QPainterPath*)path);
	return (void*)(__r);
}
void* QGraphicsScene_addPixmap(void* self, void* pixmap) {
	auto&& __r = ((QGraphicsScene*)self)->addPixmap(*(QPixmap*)pixmap);
	return (void*)(__r);
}
void* QGraphicsScene_addPolygon(void* self, void* polygon) {
	auto&& __r = ((QGraphicsScene*)self)->addPolygon(*(QPolygonF*)polygon);
	return (void*)(__r);
}
void* QGraphicsScene_addRect(void* self, void* rect) {
	auto&& __r = ((QGraphicsScene*)self)->addRect(*(QRectF*)rect);
	return (void*)(__r);
}
void* QGraphicsScene_addText(void* self, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = ((QGraphicsScene*)self)->addText(text_qstr);
	return (void*)(__r);
}
void* QGraphicsScene_addSimpleText(void* self, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = ((QGraphicsScene*)self)->addSimpleText(text_qstr);
	return (void*)(__r);
}
void* QGraphicsScene_addWidget(void* self, void* widget) {
	auto&& __r = ((QGraphicsScene*)self)->addWidget((QWidget*)widget);
	return (void*)(__r);
}
void* QGraphicsScene_addEllipse2(void* self, double x, double y, double w, double h) {
	auto&& __r = ((QGraphicsScene*)self)->addEllipse(x, y, w, h);
	return (void*)(__r);
}
void* QGraphicsScene_addLine2(void* self, double x1, double y1, double x2, double y2) {
	auto&& __r = ((QGraphicsScene*)self)->addLine(x1, y1, x2, y2);
	return (void*)(__r);
}
void* QGraphicsScene_addRect2(void* self, double x, double y, double w, double h) {
	auto&& __r = ((QGraphicsScene*)self)->addRect(x, y, w, h);
	return (void*)(__r);
}
void QGraphicsScene_removeItem(void* self, void* item) {
	((QGraphicsScene*)self)->removeItem((QGraphicsItem*)item);
}
void* QGraphicsScene_focusItem(void* self) {
	auto&& __r = ((QGraphicsScene*)self)->focusItem();
	return (void*)(__r);
}
void QGraphicsScene_setFocusItem(void* self, void* item) {
	((QGraphicsScene*)self)->setFocusItem((QGraphicsItem*)item);
}
int QGraphicsScene_hasFocus(void* self) {
	auto&& __r = ((QGraphicsScene*)self)->hasFocus();
	return __r;
}
void QGraphicsScene_setFocus(void* self) {
	((QGraphicsScene*)self)->setFocus();
}
void QGraphicsScene_clearFocus(void* self) {
	((QGraphicsScene*)self)->clearFocus();
}
void QGraphicsScene_setStickyFocus(void* self, int enabled) {
	((QGraphicsScene*)self)->setStickyFocus(enabled);
}
int QGraphicsScene_stickyFocus(void* self) {
	auto&& __r = ((QGraphicsScene*)self)->stickyFocus();
	return __r;
}
void* QGraphicsScene_mouseGrabberItem(void* self) {
	auto&& __r = ((QGraphicsScene*)self)->mouseGrabberItem();
	return (void*)(__r);
}
void* QGraphicsScene_backgroundBrush(void* self) {
	auto&& __r = ((QGraphicsScene*)self)->backgroundBrush();
	return new QBrush(__r);
}
void QGraphicsScene_setBackgroundBrush(void* self, void* brush) {
	((QGraphicsScene*)self)->setBackgroundBrush(*(QBrush*)brush);
}
void* QGraphicsScene_foregroundBrush(void* self) {
	auto&& __r = ((QGraphicsScene*)self)->foregroundBrush();
	return new QBrush(__r);
}
void QGraphicsScene_setForegroundBrush(void* self, void* brush) {
	((QGraphicsScene*)self)->setForegroundBrush(*(QBrush*)brush);
}
void* QGraphicsScene_inputMethodQuery(void* self, int query) {
	auto&& __r = ((QGraphicsScene*)self)->inputMethodQuery((Qt::InputMethodQuery)(query));
	return new QVariant(__r);
}
void** QGraphicsScene_views(void* self, long long* __out_len) {
	auto&& __r = ((QGraphicsScene*)self)->views();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
void QGraphicsScene_update(void* self, double x, double y, double w, double h) {
	((QGraphicsScene*)self)->update(x, y, w, h);
}
void QGraphicsScene_invalidate(void* self, double x, double y, double w, double h) {
	((QGraphicsScene*)self)->invalidate(x, y, w, h);
}
void* QGraphicsScene_style(void* self) {
	auto&& __r = ((QGraphicsScene*)self)->style();
	return (void*)(__r);
}
void QGraphicsScene_setStyle(void* self, void* style) {
	((QGraphicsScene*)self)->setStyle((QStyle*)style);
}
void* QGraphicsScene_font(void* self) {
	auto&& __r = ((QGraphicsScene*)self)->font();
	return new QFont(__r);
}
void QGraphicsScene_setFont(void* self, void* font) {
	((QGraphicsScene*)self)->setFont(*(QFont*)font);
}
void* QGraphicsScene_palette(void* self) {
	auto&& __r = ((QGraphicsScene*)self)->palette();
	return new QPalette(__r);
}
void QGraphicsScene_setPalette(void* self, void* palette) {
	((QGraphicsScene*)self)->setPalette(*(QPalette*)palette);
}
int QGraphicsScene_isActive(void* self) {
	auto&& __r = ((QGraphicsScene*)self)->isActive();
	return __r;
}
void* QGraphicsScene_activePanel(void* self) {
	auto&& __r = ((QGraphicsScene*)self)->activePanel();
	return (void*)(__r);
}
void QGraphicsScene_setActivePanel(void* self, void* item) {
	((QGraphicsScene*)self)->setActivePanel((QGraphicsItem*)item);
}
void* QGraphicsScene_activeWindow(void* self) {
	auto&& __r = ((QGraphicsScene*)self)->activeWindow();
	return (void*)(__r);
}
void QGraphicsScene_setActiveWindow(void* self, void* widget) {
	((QGraphicsScene*)self)->setActiveWindow((QGraphicsWidget*)widget);
}
int QGraphicsScene_sendEvent(void* self, void* item, void* event) {
	auto&& __r = ((QGraphicsScene*)self)->sendEvent((QGraphicsItem*)item, (QEvent*)event);
	return __r;
}
double QGraphicsScene_minimumRenderSize(void* self) {
	auto&& __r = ((QGraphicsScene*)self)->minimumRenderSize();
	return __r;
}
void QGraphicsScene_setMinimumRenderSize(void* self, double minSize) {
	((QGraphicsScene*)self)->setMinimumRenderSize(minSize);
}
int QGraphicsScene_focusOnTouch(void* self) {
	auto&& __r = ((QGraphicsScene*)self)->focusOnTouch();
	return __r;
}
void QGraphicsScene_setFocusOnTouch(void* self, int enabled) {
	((QGraphicsScene*)self)->setFocusOnTouch(enabled);
}
void QGraphicsScene_update2(void* self) {
	((QGraphicsScene*)self)->update();
}
void QGraphicsScene_invalidate2(void* self) {
	((QGraphicsScene*)self)->invalidate();
}
void QGraphicsScene_advance(void* self) {
	((QGraphicsScene*)self)->advance();
}
void QGraphicsScene_clearSelection(void* self) {
	((QGraphicsScene*)self)->clearSelection();
}
void QGraphicsScene_clear(void* self) {
	((QGraphicsScene*)self)->clear();
}
void QGraphicsScene_changed(void* self, void** region, long long region_len) {
	QList<QRectF> region_list; region_list.reserve(region_len);
	for (long long __i = 0; __i < region_len; __i++) region_list.append(*(QRectF*)region[__i]);
	((QGraphicsScene*)self)->changed(region_list);
}
void QGraphicsScene_sceneRectChanged(void* self, void* rect) {
	((QGraphicsScene*)self)->sceneRectChanged(*(QRectF*)rect);
}
void QGraphicsScene_selectionChanged(void* self) {
	((QGraphicsScene*)self)->selectionChanged();
}
void QGraphicsScene_focusItemChanged(void* self, void* newFocus, void* oldFocus, int reason) {
	((QGraphicsScene*)self)->focusItemChanged((QGraphicsItem*)newFocus, (QGraphicsItem*)oldFocus, (Qt::FocusReason)(reason));
}
void* QGraphicsScene_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QGraphicsScene::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QGraphicsScene_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QGraphicsScene::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QGraphicsScene_sender(void* self) {
	auto&& __r = ((QGraphicsScene__Go*)self)->sender();
	return (void*)(__r);
}
int QGraphicsScene_senderSignalIndex(void* self) {
	auto&& __r = ((QGraphicsScene__Go*)self)->senderSignalIndex();
	return __r;
}
int QGraphicsScene_receivers(void* self, const char* signal) {
	auto&& __r = ((QGraphicsScene__Go*)self)->receivers(signal);
	return __r;
}
int QGraphicsScene_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QGraphicsScene__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QGraphicsScene_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QGraphicsScene::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QGraphicsScene_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QGraphicsScene::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QGraphicsScene_render2(void* self, void* painter, void* target) {
	((QGraphicsScene*)self)->render((QPainter*)painter, *(QRectF*)target);
}
void QGraphicsScene_render3(void* self, void* painter, void* target, void* source) {
	((QGraphicsScene*)self)->render((QPainter*)painter, *(QRectF*)target, *(QRectF*)source);
}
void QGraphicsScene_render4(void* self, void* painter, void* target, void* source, int aspectRatioMode) {
	((QGraphicsScene*)self)->render((QPainter*)painter, *(QRectF*)target, *(QRectF*)source, (Qt::AspectRatioMode)(aspectRatioMode));
}
void** QGraphicsScene_items7(void* self, int order, long long* __out_len) {
	auto&& __r = ((QGraphicsScene*)self)->items((Qt::SortOrder)(order));
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
void** QGraphicsScene_items8(void* self, void* pos, int mode, long long* __out_len) {
	auto&& __r = ((QGraphicsScene*)self)->items(*(QPointF*)pos, (Qt::ItemSelectionMode)(mode));
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
void** QGraphicsScene_items9(void* self, void* pos, int mode, int order, long long* __out_len) {
	auto&& __r = ((QGraphicsScene*)self)->items(*(QPointF*)pos, (Qt::ItemSelectionMode)(mode), (Qt::SortOrder)(order));
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
void** QGraphicsScene_items10(void* self, void* pos, int mode, int order, void* deviceTransform, long long* __out_len) {
	auto&& __r = ((QGraphicsScene*)self)->items(*(QPointF*)pos, (Qt::ItemSelectionMode)(mode), (Qt::SortOrder)(order), *(QTransform*)deviceTransform);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
void** QGraphicsScene_items11(void* self, void* rect, int mode, long long* __out_len) {
	auto&& __r = ((QGraphicsScene*)self)->items(*(QRectF*)rect, (Qt::ItemSelectionMode)(mode));
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
void** QGraphicsScene_items12(void* self, void* rect, int mode, int order, long long* __out_len) {
	auto&& __r = ((QGraphicsScene*)self)->items(*(QRectF*)rect, (Qt::ItemSelectionMode)(mode), (Qt::SortOrder)(order));
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
void** QGraphicsScene_items13(void* self, void* rect, int mode, int order, void* deviceTransform, long long* __out_len) {
	auto&& __r = ((QGraphicsScene*)self)->items(*(QRectF*)rect, (Qt::ItemSelectionMode)(mode), (Qt::SortOrder)(order), *(QTransform*)deviceTransform);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
void** QGraphicsScene_items14(void* self, void* polygon, int mode, long long* __out_len) {
	auto&& __r = ((QGraphicsScene*)self)->items(*(QPolygonF*)polygon, (Qt::ItemSelectionMode)(mode));
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
void** QGraphicsScene_items15(void* self, void* polygon, int mode, int order, long long* __out_len) {
	auto&& __r = ((QGraphicsScene*)self)->items(*(QPolygonF*)polygon, (Qt::ItemSelectionMode)(mode), (Qt::SortOrder)(order));
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
void** QGraphicsScene_items16(void* self, void* polygon, int mode, int order, void* deviceTransform, long long* __out_len) {
	auto&& __r = ((QGraphicsScene*)self)->items(*(QPolygonF*)polygon, (Qt::ItemSelectionMode)(mode), (Qt::SortOrder)(order), *(QTransform*)deviceTransform);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
void** QGraphicsScene_items17(void* self, void* path, int mode, long long* __out_len) {
	auto&& __r = ((QGraphicsScene*)self)->items(*(QPainterPath*)path, (Qt::ItemSelectionMode)(mode));
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
void** QGraphicsScene_items18(void* self, void* path, int mode, int order, long long* __out_len) {
	auto&& __r = ((QGraphicsScene*)self)->items(*(QPainterPath*)path, (Qt::ItemSelectionMode)(mode), (Qt::SortOrder)(order));
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
void** QGraphicsScene_items19(void* self, void* path, int mode, int order, void* deviceTransform, long long* __out_len) {
	auto&& __r = ((QGraphicsScene*)self)->items(*(QPainterPath*)path, (Qt::ItemSelectionMode)(mode), (Qt::SortOrder)(order), *(QTransform*)deviceTransform);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
void** QGraphicsScene_items20(void* self, double x, double y, double w, double h, int mode, int order, void* deviceTransform, long long* __out_len) {
	auto&& __r = ((QGraphicsScene*)self)->items(x, y, w, h, (Qt::ItemSelectionMode)(mode), (Qt::SortOrder)(order), *(QTransform*)deviceTransform);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
void** QGraphicsScene_collidingItems2(void* self, void* item, int mode, long long* __out_len) {
	auto&& __r = ((QGraphicsScene*)self)->collidingItems((QGraphicsItem*)item, (Qt::ItemSelectionMode)(mode));
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
void QGraphicsScene_setSelectionArea3(void* self, void* path, int selectionOperation) {
	((QGraphicsScene*)self)->setSelectionArea(*(QPainterPath*)path, (Qt::ItemSelectionOperation)(selectionOperation));
}
void QGraphicsScene_setSelectionArea4(void* self, void* path, int selectionOperation, int mode) {
	((QGraphicsScene*)self)->setSelectionArea(*(QPainterPath*)path, (Qt::ItemSelectionOperation)(selectionOperation), (Qt::ItemSelectionMode)(mode));
}
void QGraphicsScene_setSelectionArea5(void* self, void* path, int selectionOperation, int mode, void* deviceTransform) {
	((QGraphicsScene*)self)->setSelectionArea(*(QPainterPath*)path, (Qt::ItemSelectionOperation)(selectionOperation), (Qt::ItemSelectionMode)(mode), *(QTransform*)deviceTransform);
}
void* QGraphicsScene_addEllipse3(void* self, void* rect, void* pen) {
	auto&& __r = ((QGraphicsScene*)self)->addEllipse(*(QRectF*)rect, *(QPen*)pen);
	return (void*)(__r);
}
void* QGraphicsScene_addEllipse4(void* self, void* rect, void* pen, void* brush) {
	auto&& __r = ((QGraphicsScene*)self)->addEllipse(*(QRectF*)rect, *(QPen*)pen, *(QBrush*)brush);
	return (void*)(__r);
}
void* QGraphicsScene_addLine3(void* self, void* line, void* pen) {
	auto&& __r = ((QGraphicsScene*)self)->addLine(*(QLineF*)line, *(QPen*)pen);
	return (void*)(__r);
}
void* QGraphicsScene_addPath2(void* self, void* path, void* pen) {
	auto&& __r = ((QGraphicsScene*)self)->addPath(*(QPainterPath*)path, *(QPen*)pen);
	return (void*)(__r);
}
void* QGraphicsScene_addPath3(void* self, void* path, void* pen, void* brush) {
	auto&& __r = ((QGraphicsScene*)self)->addPath(*(QPainterPath*)path, *(QPen*)pen, *(QBrush*)brush);
	return (void*)(__r);
}
void* QGraphicsScene_addPolygon2(void* self, void* polygon, void* pen) {
	auto&& __r = ((QGraphicsScene*)self)->addPolygon(*(QPolygonF*)polygon, *(QPen*)pen);
	return (void*)(__r);
}
void* QGraphicsScene_addPolygon3(void* self, void* polygon, void* pen, void* brush) {
	auto&& __r = ((QGraphicsScene*)self)->addPolygon(*(QPolygonF*)polygon, *(QPen*)pen, *(QBrush*)brush);
	return (void*)(__r);
}
void* QGraphicsScene_addRect3(void* self, void* rect, void* pen) {
	auto&& __r = ((QGraphicsScene*)self)->addRect(*(QRectF*)rect, *(QPen*)pen);
	return (void*)(__r);
}
void* QGraphicsScene_addRect4(void* self, void* rect, void* pen, void* brush) {
	auto&& __r = ((QGraphicsScene*)self)->addRect(*(QRectF*)rect, *(QPen*)pen, *(QBrush*)brush);
	return (void*)(__r);
}
void* QGraphicsScene_addText2(void* self, const char* text, long long text_len, void* font) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = ((QGraphicsScene*)self)->addText(text_qstr, *(QFont*)font);
	return (void*)(__r);
}
void* QGraphicsScene_addSimpleText2(void* self, const char* text, long long text_len, void* font) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = ((QGraphicsScene*)self)->addSimpleText(text_qstr, *(QFont*)font);
	return (void*)(__r);
}
void* QGraphicsScene_addWidget2(void* self, void* widget, int wFlags) {
	auto&& __r = ((QGraphicsScene*)self)->addWidget((QWidget*)widget, QFlags<Qt::WindowType>(wFlags));
	return (void*)(__r);
}
void* QGraphicsScene_addEllipse5(void* self, double x, double y, double w, double h, void* pen) {
	auto&& __r = ((QGraphicsScene*)self)->addEllipse(x, y, w, h, *(QPen*)pen);
	return (void*)(__r);
}
void* QGraphicsScene_addEllipse6(void* self, double x, double y, double w, double h, void* pen, void* brush) {
	auto&& __r = ((QGraphicsScene*)self)->addEllipse(x, y, w, h, *(QPen*)pen, *(QBrush*)brush);
	return (void*)(__r);
}
void* QGraphicsScene_addLine4(void* self, double x1, double y1, double x2, double y2, void* pen) {
	auto&& __r = ((QGraphicsScene*)self)->addLine(x1, y1, x2, y2, *(QPen*)pen);
	return (void*)(__r);
}
void* QGraphicsScene_addRect5(void* self, double x, double y, double w, double h, void* pen) {
	auto&& __r = ((QGraphicsScene*)self)->addRect(x, y, w, h, *(QPen*)pen);
	return (void*)(__r);
}
void* QGraphicsScene_addRect6(void* self, double x, double y, double w, double h, void* pen, void* brush) {
	auto&& __r = ((QGraphicsScene*)self)->addRect(x, y, w, h, *(QPen*)pen, *(QBrush*)brush);
	return (void*)(__r);
}
void QGraphicsScene_setFocusItem2(void* self, void* item, int focusReason) {
	((QGraphicsScene*)self)->setFocusItem((QGraphicsItem*)item, (Qt::FocusReason)(focusReason));
}
void QGraphicsScene_setFocus2(void* self, int focusReason) {
	((QGraphicsScene*)self)->setFocus((Qt::FocusReason)(focusReason));
}
void QGraphicsScene_invalidate3(void* self, double x, double y, double w, double h, int layers) {
	((QGraphicsScene*)self)->invalidate(x, y, w, h, QFlags<QGraphicsScene::SceneLayer>(layers));
}
void QGraphicsScene_update3(void* self, void* rect) {
	((QGraphicsScene*)self)->update(*(QRectF*)rect);
}
void QGraphicsScene_invalidate4(void* self, void* rect) {
	((QGraphicsScene*)self)->invalidate(*(QRectF*)rect);
}
void QGraphicsScene_invalidate5(void* self, void* rect, int layers) {
	((QGraphicsScene*)self)->invalidate(*(QRectF*)rect, QFlags<QGraphicsScene::SceneLayer>(layers));
}
void* QGraphicsScene_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QGraphicsScene::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QGraphicsScene_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QGraphicsScene::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QGraphicsScene_inputMethodQuery_default(void* self, int query) {
	auto&& __r = ((QGraphicsScene*)self)->QGraphicsScene::inputMethodQuery((Qt::InputMethodQuery)(query));
	return new QVariant(__r);
}
void* QGraphicsScene_metaObject_default(void* self) {
	auto&& __r = ((QGraphicsScene*)self)->QGraphicsScene::metaObject();
	return (void*)(__r);
}
int QGraphicsScene_event_default(void* self, void* event) {
	auto&& __r = ((QGraphicsScene__Go*)self)->cppgoBase_event((QEvent*)event);
	return __r;
}
int QGraphicsScene_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QGraphicsScene__Go*)self)->cppgoBase_eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QGraphicsScene_contextMenuEvent_default(void* self, void* event) {
	((QGraphicsScene__Go*)self)->cppgoBase_contextMenuEvent((QGraphicsSceneContextMenuEvent*)event);
}
void QGraphicsScene_dragEnterEvent_default(void* self, void* event) {
	((QGraphicsScene__Go*)self)->cppgoBase_dragEnterEvent((QGraphicsSceneDragDropEvent*)event);
}
void QGraphicsScene_dragMoveEvent_default(void* self, void* event) {
	((QGraphicsScene__Go*)self)->cppgoBase_dragMoveEvent((QGraphicsSceneDragDropEvent*)event);
}
void QGraphicsScene_dragLeaveEvent_default(void* self, void* event) {
	((QGraphicsScene__Go*)self)->cppgoBase_dragLeaveEvent((QGraphicsSceneDragDropEvent*)event);
}
void QGraphicsScene_dropEvent_default(void* self, void* event) {
	((QGraphicsScene__Go*)self)->cppgoBase_dropEvent((QGraphicsSceneDragDropEvent*)event);
}
void QGraphicsScene_focusInEvent_default(void* self, void* event) {
	((QGraphicsScene__Go*)self)->cppgoBase_focusInEvent((QFocusEvent*)event);
}
void QGraphicsScene_focusOutEvent_default(void* self, void* event) {
	((QGraphicsScene__Go*)self)->cppgoBase_focusOutEvent((QFocusEvent*)event);
}
void QGraphicsScene_helpEvent_default(void* self, void* event) {
	((QGraphicsScene__Go*)self)->cppgoBase_helpEvent((QGraphicsSceneHelpEvent*)event);
}
void QGraphicsScene_keyPressEvent_default(void* self, void* event) {
	((QGraphicsScene__Go*)self)->cppgoBase_keyPressEvent((QKeyEvent*)event);
}
void QGraphicsScene_keyReleaseEvent_default(void* self, void* event) {
	((QGraphicsScene__Go*)self)->cppgoBase_keyReleaseEvent((QKeyEvent*)event);
}
void QGraphicsScene_mousePressEvent_default(void* self, void* event) {
	((QGraphicsScene__Go*)self)->cppgoBase_mousePressEvent((QGraphicsSceneMouseEvent*)event);
}
void QGraphicsScene_mouseMoveEvent_default(void* self, void* event) {
	((QGraphicsScene__Go*)self)->cppgoBase_mouseMoveEvent((QGraphicsSceneMouseEvent*)event);
}
void QGraphicsScene_mouseReleaseEvent_default(void* self, void* event) {
	((QGraphicsScene__Go*)self)->cppgoBase_mouseReleaseEvent((QGraphicsSceneMouseEvent*)event);
}
void QGraphicsScene_mouseDoubleClickEvent_default(void* self, void* event) {
	((QGraphicsScene__Go*)self)->cppgoBase_mouseDoubleClickEvent((QGraphicsSceneMouseEvent*)event);
}
void QGraphicsScene_wheelEvent_default(void* self, void* event) {
	((QGraphicsScene__Go*)self)->cppgoBase_wheelEvent((QGraphicsSceneWheelEvent*)event);
}
void QGraphicsScene_inputMethodEvent_default(void* self, void* event) {
	((QGraphicsScene__Go*)self)->cppgoBase_inputMethodEvent((QInputMethodEvent*)event);
}
void QGraphicsScene_drawBackground_default(void* self, void* painter, void* rect) {
	((QGraphicsScene__Go*)self)->cppgoBase_drawBackground((QPainter*)painter, *(QRectF*)rect);
}
void QGraphicsScene_drawForeground_default(void* self, void* painter, void* rect) {
	((QGraphicsScene__Go*)self)->cppgoBase_drawForeground((QPainter*)painter, *(QRectF*)rect);
}
int QGraphicsScene_focusNextPrevChild_default(void* self, int next) {
	auto&& __r = ((QGraphicsScene__Go*)self)->cppgoBase_focusNextPrevChild(next);
	return __r;
}
void QGraphicsScene_timerEvent_default(void* self, void* event) {
	((QGraphicsScene__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QGraphicsScene_childEvent_default(void* self, void* event) {
	((QGraphicsScene__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QGraphicsScene_customEvent_default(void* self, void* event) {
	((QGraphicsScene__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QGraphicsScene_connectNotify_default(void* self, void* signal) {
	((QGraphicsScene__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QGraphicsScene_disconnectNotify_default(void* self, void* signal) {
	((QGraphicsScene__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QGraphicsScene_connect_changed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsScene*)self, &QGraphicsScene::changed, (QGraphicsScene*)self, [__id](const QList<QRectF>& region){
		std::vector<void*> region_v; region_v.reserve(region.size());
		for (long long __i = 0; __i < region.size(); __i++) region_v.push_back((void*)&region[__i]);
		if (cppgo_sig_void_voidpp_longlong) ((void(*)(long long, void**, long long))cppgo_sig_void_voidpp_longlong)(__id, region_v.data(), (long long)region.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsScene);
	__conns_QGraphicsScene[__id] = __c;
}
void QGraphicsScene_connect_sceneRectChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsScene*)self, &QGraphicsScene::sceneRectChanged, (QGraphicsScene*)self, [__id](const QRectF& rect){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&rect);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsScene);
	__conns_QGraphicsScene[__id] = __c;
}
void QGraphicsScene_connect_selectionChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsScene*)self, &QGraphicsScene::selectionChanged, (QGraphicsScene*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsScene);
	__conns_QGraphicsScene[__id] = __c;
}
void QGraphicsScene_connect_focusItemChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsScene*)self, &QGraphicsScene::focusItemChanged, (QGraphicsScene*)self, [__id](QGraphicsItem* newFocus, QGraphicsItem* oldFocus, Qt::FocusReason reason){
		if (cppgo_sig_void_voidp_voidp_int) ((void(*)(long long, void*, void*, int))cppgo_sig_void_voidp_voidp_int)(__id, (void*)newFocus, (void*)oldFocus, (int)(reason));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsScene);
	__conns_QGraphicsScene[__id] = __c;
}
void QGraphicsScene_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsScene*)self, &QGraphicsScene::destroyed, (QGraphicsScene*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsScene);
	__conns_QGraphicsScene[__id] = __c;
}
void QGraphicsScene_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsScene*)self, &QGraphicsScene::destroyed, (QGraphicsScene*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsScene);
	__conns_QGraphicsScene[__id] = __c;
}
void QGraphicsScene_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsScene*)self, &QGraphicsScene::objectNameChanged, (QGraphicsScene*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsScene);
	__conns_QGraphicsScene[__id] = __c;
}
void QGraphicsScene__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsScene);
		auto __it = __conns_QGraphicsScene.find(id);
		if (__it == __conns_QGraphicsScene.end()) return;
		__c = __it->second;
		__conns_QGraphicsScene.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QGraphicsScene__Go_new(long long id, void* parent) { return new QGraphicsScene__Go(id, (QObject*)parent); }
void QGraphicsScene_delete(void* self) { delete (QGraphicsScene*)self; }

}

#include <QtWidgets/qlcdnumber.h>
#include <QtWidgets/qframe.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qsize.h>
#include <QtCore/qcoreevent.h>
#include <QtGui/qevent.h>
#include <QtCore/qrect.h>
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
#include "gen_qlcdnumber.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QLCDNumber__Go : public QLCDNumber {
public:
	long long __goId;
	unsigned char __goOn[49] = {};
	template<typename... A> QLCDNumber__Go(long long id, A&&... a) : QLCDNumber(std::forward<A>(a)...), __goId(id) {}
	using QLCDNumber::drawFrame;
	using QLCDNumber::updateMicroFocus;
	using QLCDNumber::create;
	using QLCDNumber::destroy;
	using QLCDNumber::focusNextChild;
	using QLCDNumber::focusPreviousChild;
	using QLCDNumber::sender;
	using QLCDNumber::senderSignalIndex;
	using QLCDNumber::receivers;
	using QLCDNumber::isSignalConnected;
	using QLCDNumber::getDecodedMetricF;
	bool cppgoBase_event(QEvent* __a0) { return QLCDNumber::event(__a0); }
	void cppgoBase_paintEvent(QPaintEvent* __a0) { QLCDNumber::paintEvent(__a0); }
	void cppgoBase_changeEvent(QEvent* __a0) { QLCDNumber::changeEvent(__a0); }
	void cppgoBase_initStyleOption(QStyleOptionFrame* __a0) { QLCDNumber::initStyleOption(__a0); }
	void cppgoBase_mousePressEvent(QMouseEvent* __a0) { QLCDNumber::mousePressEvent(__a0); }
	void cppgoBase_mouseReleaseEvent(QMouseEvent* __a0) { QLCDNumber::mouseReleaseEvent(__a0); }
	void cppgoBase_mouseDoubleClickEvent(QMouseEvent* __a0) { QLCDNumber::mouseDoubleClickEvent(__a0); }
	void cppgoBase_mouseMoveEvent(QMouseEvent* __a0) { QLCDNumber::mouseMoveEvent(__a0); }
	void cppgoBase_wheelEvent(QWheelEvent* __a0) { QLCDNumber::wheelEvent(__a0); }
	void cppgoBase_keyPressEvent(QKeyEvent* __a0) { QLCDNumber::keyPressEvent(__a0); }
	void cppgoBase_keyReleaseEvent(QKeyEvent* __a0) { QLCDNumber::keyReleaseEvent(__a0); }
	void cppgoBase_focusInEvent(QFocusEvent* __a0) { QLCDNumber::focusInEvent(__a0); }
	void cppgoBase_focusOutEvent(QFocusEvent* __a0) { QLCDNumber::focusOutEvent(__a0); }
	void cppgoBase_enterEvent(QEnterEvent* __a0) { QLCDNumber::enterEvent(__a0); }
	void cppgoBase_leaveEvent(QEvent* __a0) { QLCDNumber::leaveEvent(__a0); }
	void cppgoBase_moveEvent(QMoveEvent* __a0) { QLCDNumber::moveEvent(__a0); }
	void cppgoBase_resizeEvent(QResizeEvent* __a0) { QLCDNumber::resizeEvent(__a0); }
	void cppgoBase_closeEvent(QCloseEvent* __a0) { QLCDNumber::closeEvent(__a0); }
	void cppgoBase_contextMenuEvent(QContextMenuEvent* __a0) { QLCDNumber::contextMenuEvent(__a0); }
	void cppgoBase_tabletEvent(QTabletEvent* __a0) { QLCDNumber::tabletEvent(__a0); }
	void cppgoBase_actionEvent(QActionEvent* __a0) { QLCDNumber::actionEvent(__a0); }
	void cppgoBase_dragEnterEvent(QDragEnterEvent* __a0) { QLCDNumber::dragEnterEvent(__a0); }
	void cppgoBase_dragMoveEvent(QDragMoveEvent* __a0) { QLCDNumber::dragMoveEvent(__a0); }
	void cppgoBase_dragLeaveEvent(QDragLeaveEvent* __a0) { QLCDNumber::dragLeaveEvent(__a0); }
	void cppgoBase_dropEvent(QDropEvent* __a0) { QLCDNumber::dropEvent(__a0); }
	void cppgoBase_showEvent(QShowEvent* __a0) { QLCDNumber::showEvent(__a0); }
	void cppgoBase_hideEvent(QHideEvent* __a0) { QLCDNumber::hideEvent(__a0); }
	int cppgoBase_metric(int __a0) { return QLCDNumber::metric((QPaintDevice::PaintDeviceMetric)__a0); }
	void cppgoBase_initPainter(QPainter* __a0) { QLCDNumber::initPainter(__a0); }
	QPaintDevice* cppgoBase_redirected(QPoint* __a0) { return QLCDNumber::redirected(__a0); }
	QPainter* cppgoBase_sharedPainter() { return QLCDNumber::sharedPainter(); }
	void cppgoBase_inputMethodEvent(QInputMethodEvent* __a0) { QLCDNumber::inputMethodEvent(__a0); }
	bool cppgoBase_focusNextPrevChild(bool __a0) { return QLCDNumber::focusNextPrevChild(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QLCDNumber::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QLCDNumber::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QLCDNumber::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QLCDNumber::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QLCDNumber::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QLCDNumber::metaObject();
	}
	QSize sizeHint() const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 729690261457684101LL, &__h); if (__h) return *(QSize*)__r; }
		return QLCDNumber::sizeHint();
	}
	bool event(QEvent* e) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)e, &__h); if (__h) return __r; }
		return QLCDNumber::event(e);
	}
	void paintEvent(QPaintEvent* arg1) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5290442163167458465LL, (void*)arg1, &__h); if (__h) return; }
		QLCDNumber::paintEvent(arg1);
	}
	void changeEvent(QEvent* arg1) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4959376063826165943LL, (void*)arg1, &__h); if (__h) return; }
		QLCDNumber::changeEvent(arg1);
	}
	void initStyleOption(QStyleOptionFrame* option) const override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -488693769771308531LL, (void*)option, &__h); if (__h) return; }
		QLCDNumber::initStyleOption(option);
	}
	int devType() const override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8659679893868524044LL, &__h); if (__h) return __r; }
		return QLCDNumber::devType();
	}
	void setVisible(bool visible) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 2110182719375245575LL, visible, &__h); if (__h) return; }
		QLCDNumber::setVisible(visible);
	}
	QSize minimumSizeHint() const override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 7049264599837532491LL, &__h); if (__h) return *(QSize*)__r; }
		return QLCDNumber::minimumSizeHint();
	}
	int heightForWidth(int arg1) const override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 4311534060495737339LL, arg1, &__h); if (__h) return __r; }
		return QLCDNumber::heightForWidth(arg1);
	}
	bool hasHeightForWidth() const override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8266762397995423915LL, &__h); if (__h) return __r; }
		return QLCDNumber::hasHeightForWidth();
	}
	QPaintEngine* paintEngine() const override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 6398007404756346381LL, &__h); if (__h) return (QPaintEngine*)__r; }
		return QLCDNumber::paintEngine();
	}
	void mousePressEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -332090898078543319LL, (void*)event, &__h); if (__h) return; }
		QLCDNumber::mousePressEvent(event);
	}
	void mouseReleaseEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4415920049297795213LL, (void*)event, &__h); if (__h) return; }
		QLCDNumber::mouseReleaseEvent(event);
	}
	void mouseDoubleClickEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4029073905697023887LL, (void*)event, &__h); if (__h) return; }
		QLCDNumber::mouseDoubleClickEvent(event);
	}
	void mouseMoveEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4255494004099806581LL, (void*)event, &__h); if (__h) return; }
		QLCDNumber::mouseMoveEvent(event);
	}
	void wheelEvent(QWheelEvent* event) override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1919988340934511316LL, (void*)event, &__h); if (__h) return; }
		QLCDNumber::wheelEvent(event);
	}
	void keyPressEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1023791977679793185LL, (void*)event, &__h); if (__h) return; }
		QLCDNumber::keyPressEvent(event);
	}
	void keyReleaseEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 772598475950050299LL, (void*)event, &__h); if (__h) return; }
		QLCDNumber::keyReleaseEvent(event);
	}
	void focusInEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 975043767033552552LL, (void*)event, &__h); if (__h) return; }
		QLCDNumber::focusInEvent(event);
	}
	void focusOutEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -385596010576048993LL, (void*)event, &__h); if (__h) return; }
		QLCDNumber::focusOutEvent(event);
	}
	void enterEvent(QEnterEvent* event) override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6316270346668237895LL, (void*)event, &__h); if (__h) return; }
		QLCDNumber::enterEvent(event);
	}
	void leaveEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -551070009623711640LL, (void*)event, &__h); if (__h) return; }
		QLCDNumber::leaveEvent(event);
	}
	void moveEvent(QMoveEvent* event) override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -9161949150489525396LL, (void*)event, &__h); if (__h) return; }
		QLCDNumber::moveEvent(event);
	}
	void resizeEvent(QResizeEvent* event) override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4302917285655014661LL, (void*)event, &__h); if (__h) return; }
		QLCDNumber::resizeEvent(event);
	}
	void closeEvent(QCloseEvent* event) override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1989937050019464783LL, (void*)event, &__h); if (__h) return; }
		QLCDNumber::closeEvent(event);
	}
	void contextMenuEvent(QContextMenuEvent* event) override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6698043222227006015LL, (void*)event, &__h); if (__h) return; }
		QLCDNumber::contextMenuEvent(event);
	}
	void tabletEvent(QTabletEvent* event) override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -212068866860701141LL, (void*)event, &__h); if (__h) return; }
		QLCDNumber::tabletEvent(event);
	}
	void actionEvent(QActionEvent* event) override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1697444140088555581LL, (void*)event, &__h); if (__h) return; }
		QLCDNumber::actionEvent(event);
	}
	void dragEnterEvent(QDragEnterEvent* event) override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 206122715919917589LL, (void*)event, &__h); if (__h) return; }
		QLCDNumber::dragEnterEvent(event);
	}
	void dragMoveEvent(QDragMoveEvent* event) override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1781545137352745978LL, (void*)event, &__h); if (__h) return; }
		QLCDNumber::dragMoveEvent(event);
	}
	void dragLeaveEvent(QDragLeaveEvent* event) override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -3139904158373725006LL, (void*)event, &__h); if (__h) return; }
		QLCDNumber::dragLeaveEvent(event);
	}
	void dropEvent(QDropEvent* event) override {
		int __h = 0;
		if (__goOn[32] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 9107058682754591716LL, (void*)event, &__h); if (__h) return; }
		QLCDNumber::dropEvent(event);
	}
	void showEvent(QShowEvent* event) override {
		int __h = 0;
		if (__goOn[33] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6825369174751916LL, (void*)event, &__h); if (__h) return; }
		QLCDNumber::showEvent(event);
	}
	void hideEvent(QHideEvent* event) override {
		int __h = 0;
		if (__goOn[34] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -199718707876719129LL, (void*)event, &__h); if (__h) return; }
		QLCDNumber::hideEvent(event);
	}
	bool nativeEvent(const QByteArray& eventType, void* message, qintptr* result) override {
		int __h = 0;
		if (__goOn[35] && cppgo_virt_int_voidp_voidp_longlongp) { int __r = ((int(*)(long long, long long, void*, void*, long long*, int*))cppgo_virt_int_voidp_voidp_longlongp)(__goId, -2039166586246443706LL, (void*)&eventType, (void*)message, (long long*)result, &__h); if (__h) return __r; }
		return QLCDNumber::nativeEvent(eventType, message, result);
	}
	int metric(QPaintDevice::PaintDeviceMetric arg1) const override {
		int __h = 0;
		if (__goOn[36] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 427629649267311777LL, (int)(arg1), &__h); if (__h) return __r; }
		return QLCDNumber::metric(arg1);
	}
	void initPainter(QPainter* painter) const override {
		int __h = 0;
		if (__goOn[37] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6477500296980050210LL, (void*)painter, &__h); if (__h) return; }
		QLCDNumber::initPainter(painter);
	}
	QPaintDevice* redirected(QPoint* offset) const override {
		int __h = 0;
		if (__goOn[38] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 4127240633310655048LL, (void*)offset, &__h); if (__h) return (QPaintDevice*)__r; }
		return QLCDNumber::redirected(offset);
	}
	QPainter* sharedPainter() const override {
		int __h = 0;
		if (__goOn[39] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 5792771328741774991LL, &__h); if (__h) return (QPainter*)__r; }
		return QLCDNumber::sharedPainter();
	}
	void inputMethodEvent(QInputMethodEvent* arg1) override {
		int __h = 0;
		if (__goOn[40] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5438819661942843222LL, (void*)arg1, &__h); if (__h) return; }
		QLCDNumber::inputMethodEvent(arg1);
	}
	QVariant inputMethodQuery(Qt::InputMethodQuery arg1) const override {
		int __h = 0;
		if (__goOn[41] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7345853097340709638LL, (int)(arg1), &__h); if (__h) return *(QVariant*)__r; }
		return QLCDNumber::inputMethodQuery(arg1);
	}
	bool focusNextPrevChild(bool next) override {
		int __h = 0;
		if (__goOn[42] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 5898328517505784191LL, next, &__h); if (__h) return __r; }
		return QLCDNumber::focusNextPrevChild(next);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[43] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QLCDNumber::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[44] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QLCDNumber::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[45] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QLCDNumber::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[46] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QLCDNumber::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[47] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QLCDNumber::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[48] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QLCDNumber::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QLCDNumber;
static std::mutex __connsMu_QLCDNumber;

extern "C" {

void* QLCDNumber_base(void* self) { return (void*)static_cast<QFrame*>((QLCDNumber*)self); }
void QLCDNumber_chain(void* self, void** out) {
	QLCDNumber* __p = (QLCDNumber*)self;
	QFrame* __b0 = static_cast<QFrame*>(__p); out[0] = (void*)__b0;
	QWidget* __b1 = static_cast<QWidget*>(__b0); out[1] = (void*)__b1;
	QObject* __b2 = static_cast<QObject*>(__b1); out[2] = (void*)__b2;
}
void QLCDNumber__set_override(void* self, int idx) { ((QLCDNumber__Go*)self)->__goOn[idx] = 1; }
void* QLCDNumber_fromQObject(void* o) { return (void*)qobject_cast<QLCDNumber*>((QObject*)o); }
void* QLCDNumber_new(long long __goid, void* parent) {
	return new QLCDNumber__Go(__goid, (QWidget*)parent);
}
void* QLCDNumber_new2(long long __goid) {
	return new QLCDNumber__Go(__goid);
}
void* QLCDNumber_new3(long long __goid, unsigned int numDigits) {
	return new QLCDNumber__Go(__goid, numDigits);
}
void* QLCDNumber_new4(long long __goid, unsigned int numDigits, void* parent) {
	return new QLCDNumber__Go(__goid, numDigits, (QWidget*)parent);
}
char* QLCDNumber_tr(const char* s, long long* __out_len) {
	auto&& __r = QLCDNumber::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QLCDNumber_smallDecimalPoint(void* self) {
	auto&& __r = ((QLCDNumber*)self)->smallDecimalPoint();
	return __r;
}
int QLCDNumber_digitCount(void* self) {
	auto&& __r = ((QLCDNumber*)self)->digitCount();
	return __r;
}
void QLCDNumber_setDigitCount(void* self, int nDigits) {
	((QLCDNumber*)self)->setDigitCount(nDigits);
}
int QLCDNumber_checkOverflow(void* self, double num) {
	auto&& __r = ((QLCDNumber*)self)->checkOverflow(num);
	return __r;
}
int QLCDNumber_checkOverflow2(void* self, int num) {
	auto&& __r = ((QLCDNumber*)self)->checkOverflow(num);
	return __r;
}
int QLCDNumber_mode(void* self) {
	auto&& __r = ((QLCDNumber*)self)->mode();
	return (int)(__r);
}
void QLCDNumber_setMode(void* self, int arg1) {
	((QLCDNumber*)self)->setMode((QLCDNumber::Mode)(arg1));
}
int QLCDNumber_segmentStyle(void* self) {
	auto&& __r = ((QLCDNumber*)self)->segmentStyle();
	return (int)(__r);
}
void QLCDNumber_setSegmentStyle(void* self, int arg1) {
	((QLCDNumber*)self)->setSegmentStyle((QLCDNumber::SegmentStyle)(arg1));
}
double QLCDNumber_value(void* self) {
	auto&& __r = ((QLCDNumber*)self)->value();
	return __r;
}
int QLCDNumber_intValue(void* self) {
	auto&& __r = ((QLCDNumber*)self)->intValue();
	return __r;
}
void QLCDNumber_display(void* self, const char* str, long long str_len) {
	QString str_qstr = QString::fromUtf8(str, str_len);
	((QLCDNumber*)self)->display(str_qstr);
}
void QLCDNumber_display2(void* self, int num) {
	((QLCDNumber*)self)->display(num);
}
void QLCDNumber_display3(void* self, double num) {
	((QLCDNumber*)self)->display(num);
}
void QLCDNumber_setHexMode(void* self) {
	((QLCDNumber*)self)->setHexMode();
}
void QLCDNumber_setDecMode(void* self) {
	((QLCDNumber*)self)->setDecMode();
}
void QLCDNumber_setOctMode(void* self) {
	((QLCDNumber*)self)->setOctMode();
}
void QLCDNumber_setBinMode(void* self) {
	((QLCDNumber*)self)->setBinMode();
}
void QLCDNumber_setSmallDecimalPoint(void* self, int arg1) {
	((QLCDNumber*)self)->setSmallDecimalPoint(arg1);
}
void QLCDNumber_overflow(void* self) {
	((QLCDNumber*)self)->overflow();
}
void QLCDNumber_drawFrame(void* self, void* arg1) {
	((QLCDNumber__Go*)self)->drawFrame((QPainter*)arg1);
}
void QLCDNumber_setTabOrder(void* arg1, void* arg2) {
	QLCDNumber::setTabOrder((QWidget*)arg1, (QWidget*)arg2);
}
void* QLCDNumber_mouseGrabber() {
	auto&& __r = QLCDNumber::mouseGrabber();
	return (void*)(__r);
}
void* QLCDNumber_keyboardGrabber() {
	auto&& __r = QLCDNumber::keyboardGrabber();
	return (void*)(__r);
}
void* QLCDNumber_find(unsigned long long arg1) {
	auto&& __r = QLCDNumber::find(arg1);
	return (void*)(__r);
}
void* QLCDNumber_createWindowContainer(void* window) {
	auto&& __r = QLCDNumber::createWindowContainer((QWindow*)window);
	return (void*)(__r);
}
void QLCDNumber_updateMicroFocus(void* self) {
	((QLCDNumber__Go*)self)->updateMicroFocus();
}
void QLCDNumber_create(void* self) {
	((QLCDNumber__Go*)self)->create();
}
void QLCDNumber_destroy(void* self) {
	((QLCDNumber__Go*)self)->destroy();
}
int QLCDNumber_focusNextChild(void* self) {
	auto&& __r = ((QLCDNumber__Go*)self)->focusNextChild();
	return __r;
}
int QLCDNumber_focusPreviousChild(void* self) {
	auto&& __r = ((QLCDNumber__Go*)self)->focusPreviousChild();
	return __r;
}
void* QLCDNumber_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QLCDNumber::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QLCDNumber_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QLCDNumber::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QLCDNumber_sender(void* self) {
	auto&& __r = ((QLCDNumber__Go*)self)->sender();
	return (void*)(__r);
}
int QLCDNumber_senderSignalIndex(void* self) {
	auto&& __r = ((QLCDNumber__Go*)self)->senderSignalIndex();
	return __r;
}
int QLCDNumber_receivers(void* self, const char* signal) {
	auto&& __r = ((QLCDNumber__Go*)self)->receivers(signal);
	return __r;
}
int QLCDNumber_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QLCDNumber__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
double QLCDNumber_devicePixelRatioFScale() {
	auto&& __r = QLCDNumber::devicePixelRatioFScale();
	return __r;
}
int QLCDNumber_encodeMetricF(int metric, double value) {
	auto&& __r = QLCDNumber::encodeMetricF((QPaintDevice::PaintDeviceMetric)(metric), value);
	return __r;
}
double QLCDNumber_getDecodedMetricF(void* self, int metricA, int metricB) {
	auto&& __r = ((QLCDNumber__Go*)self)->getDecodedMetricF((QPaintDevice::PaintDeviceMetric)(metricA), (QPaintDevice::PaintDeviceMetric)(metricB));
	return __r;
}
char* QLCDNumber_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QLCDNumber::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QLCDNumber_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QLCDNumber::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QLCDNumber_createWindowContainer2(void* window, void* parent) {
	auto&& __r = QLCDNumber::createWindowContainer((QWindow*)window, (QWidget*)parent);
	return (void*)(__r);
}
void* QLCDNumber_createWindowContainer3(void* window, void* parent, int flags) {
	auto&& __r = QLCDNumber::createWindowContainer((QWindow*)window, (QWidget*)parent, QFlags<Qt::WindowType>(flags));
	return (void*)(__r);
}
void QLCDNumber_updateMicroFocus2(void* self, int query) {
	((QLCDNumber__Go*)self)->updateMicroFocus((Qt::InputMethodQuery)(query));
}
void QLCDNumber_create2(void* self, unsigned long long arg1) {
	((QLCDNumber__Go*)self)->create(arg1);
}
void QLCDNumber_create3(void* self, unsigned long long arg1, int initializeWindow) {
	((QLCDNumber__Go*)self)->create(arg1, initializeWindow);
}
void QLCDNumber_create4(void* self, unsigned long long arg1, int initializeWindow, int destroyOldWindow) {
	((QLCDNumber__Go*)self)->create(arg1, initializeWindow, destroyOldWindow);
}
void QLCDNumber_destroy2(void* self, int destroyWindow) {
	((QLCDNumber__Go*)self)->destroy(destroyWindow);
}
void QLCDNumber_destroy3(void* self, int destroyWindow, int destroySubWindows) {
	((QLCDNumber__Go*)self)->destroy(destroyWindow, destroySubWindows);
}
void* QLCDNumber_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QLCDNumber::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QLCDNumber_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QLCDNumber::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QLCDNumber_metaObject_default(void* self) {
	auto&& __r = ((QLCDNumber*)self)->QLCDNumber::metaObject();
	return (void*)(__r);
}
void* QLCDNumber_sizeHint_default(void* self) {
	auto&& __r = ((QLCDNumber*)self)->QLCDNumber::sizeHint();
	return new QSize(__r);
}
int QLCDNumber_event_default(void* self, void* e) {
	auto&& __r = ((QLCDNumber__Go*)self)->cppgoBase_event((QEvent*)e);
	return __r;
}
void QLCDNumber_paintEvent_default(void* self, void* arg1) {
	((QLCDNumber__Go*)self)->cppgoBase_paintEvent((QPaintEvent*)arg1);
}
void QLCDNumber_changeEvent_default(void* self, void* arg1) {
	((QLCDNumber__Go*)self)->cppgoBase_changeEvent((QEvent*)arg1);
}
void QLCDNumber_initStyleOption_default(void* self, void* option) {
	((QLCDNumber__Go*)self)->cppgoBase_initStyleOption((QStyleOptionFrame*)option);
}
int QLCDNumber_devType_default(void* self) {
	auto&& __r = ((QLCDNumber*)self)->QLCDNumber::devType();
	return __r;
}
void QLCDNumber_setVisible_default(void* self, int visible) {
	((QLCDNumber*)self)->QLCDNumber::setVisible(visible);
}
void* QLCDNumber_minimumSizeHint_default(void* self) {
	auto&& __r = ((QLCDNumber*)self)->QLCDNumber::minimumSizeHint();
	return new QSize(__r);
}
int QLCDNumber_heightForWidth_default(void* self, int arg1) {
	auto&& __r = ((QLCDNumber*)self)->QLCDNumber::heightForWidth(arg1);
	return __r;
}
int QLCDNumber_hasHeightForWidth_default(void* self) {
	auto&& __r = ((QLCDNumber*)self)->QLCDNumber::hasHeightForWidth();
	return __r;
}
void* QLCDNumber_paintEngine_default(void* self) {
	auto&& __r = ((QLCDNumber*)self)->QLCDNumber::paintEngine();
	return (void*)(__r);
}
void QLCDNumber_mousePressEvent_default(void* self, void* event) {
	((QLCDNumber__Go*)self)->cppgoBase_mousePressEvent((QMouseEvent*)event);
}
void QLCDNumber_mouseReleaseEvent_default(void* self, void* event) {
	((QLCDNumber__Go*)self)->cppgoBase_mouseReleaseEvent((QMouseEvent*)event);
}
void QLCDNumber_mouseDoubleClickEvent_default(void* self, void* event) {
	((QLCDNumber__Go*)self)->cppgoBase_mouseDoubleClickEvent((QMouseEvent*)event);
}
void QLCDNumber_mouseMoveEvent_default(void* self, void* event) {
	((QLCDNumber__Go*)self)->cppgoBase_mouseMoveEvent((QMouseEvent*)event);
}
void QLCDNumber_wheelEvent_default(void* self, void* event) {
	((QLCDNumber__Go*)self)->cppgoBase_wheelEvent((QWheelEvent*)event);
}
void QLCDNumber_keyPressEvent_default(void* self, void* event) {
	((QLCDNumber__Go*)self)->cppgoBase_keyPressEvent((QKeyEvent*)event);
}
void QLCDNumber_keyReleaseEvent_default(void* self, void* event) {
	((QLCDNumber__Go*)self)->cppgoBase_keyReleaseEvent((QKeyEvent*)event);
}
void QLCDNumber_focusInEvent_default(void* self, void* event) {
	((QLCDNumber__Go*)self)->cppgoBase_focusInEvent((QFocusEvent*)event);
}
void QLCDNumber_focusOutEvent_default(void* self, void* event) {
	((QLCDNumber__Go*)self)->cppgoBase_focusOutEvent((QFocusEvent*)event);
}
void QLCDNumber_enterEvent_default(void* self, void* event) {
	((QLCDNumber__Go*)self)->cppgoBase_enterEvent((QEnterEvent*)event);
}
void QLCDNumber_leaveEvent_default(void* self, void* event) {
	((QLCDNumber__Go*)self)->cppgoBase_leaveEvent((QEvent*)event);
}
void QLCDNumber_moveEvent_default(void* self, void* event) {
	((QLCDNumber__Go*)self)->cppgoBase_moveEvent((QMoveEvent*)event);
}
void QLCDNumber_resizeEvent_default(void* self, void* event) {
	((QLCDNumber__Go*)self)->cppgoBase_resizeEvent((QResizeEvent*)event);
}
void QLCDNumber_closeEvent_default(void* self, void* event) {
	((QLCDNumber__Go*)self)->cppgoBase_closeEvent((QCloseEvent*)event);
}
void QLCDNumber_contextMenuEvent_default(void* self, void* event) {
	((QLCDNumber__Go*)self)->cppgoBase_contextMenuEvent((QContextMenuEvent*)event);
}
void QLCDNumber_tabletEvent_default(void* self, void* event) {
	((QLCDNumber__Go*)self)->cppgoBase_tabletEvent((QTabletEvent*)event);
}
void QLCDNumber_actionEvent_default(void* self, void* event) {
	((QLCDNumber__Go*)self)->cppgoBase_actionEvent((QActionEvent*)event);
}
void QLCDNumber_dragEnterEvent_default(void* self, void* event) {
	((QLCDNumber__Go*)self)->cppgoBase_dragEnterEvent((QDragEnterEvent*)event);
}
void QLCDNumber_dragMoveEvent_default(void* self, void* event) {
	((QLCDNumber__Go*)self)->cppgoBase_dragMoveEvent((QDragMoveEvent*)event);
}
void QLCDNumber_dragLeaveEvent_default(void* self, void* event) {
	((QLCDNumber__Go*)self)->cppgoBase_dragLeaveEvent((QDragLeaveEvent*)event);
}
void QLCDNumber_dropEvent_default(void* self, void* event) {
	((QLCDNumber__Go*)self)->cppgoBase_dropEvent((QDropEvent*)event);
}
void QLCDNumber_showEvent_default(void* self, void* event) {
	((QLCDNumber__Go*)self)->cppgoBase_showEvent((QShowEvent*)event);
}
void QLCDNumber_hideEvent_default(void* self, void* event) {
	((QLCDNumber__Go*)self)->cppgoBase_hideEvent((QHideEvent*)event);
}
int QLCDNumber_metric_default(void* self, int arg1) {
	auto&& __r = ((QLCDNumber__Go*)self)->cppgoBase_metric(arg1);
	return __r;
}
void QLCDNumber_initPainter_default(void* self, void* painter) {
	((QLCDNumber__Go*)self)->cppgoBase_initPainter((QPainter*)painter);
}
void* QLCDNumber_redirected_default(void* self, void* offset) {
	auto&& __r = ((QLCDNumber__Go*)self)->cppgoBase_redirected((QPoint*)offset);
	return (void*)(__r);
}
void* QLCDNumber_sharedPainter_default(void* self) {
	auto&& __r = ((QLCDNumber__Go*)self)->cppgoBase_sharedPainter();
	return (void*)(__r);
}
void QLCDNumber_inputMethodEvent_default(void* self, void* arg1) {
	((QLCDNumber__Go*)self)->cppgoBase_inputMethodEvent((QInputMethodEvent*)arg1);
}
void* QLCDNumber_inputMethodQuery_default(void* self, int arg1) {
	auto&& __r = ((QLCDNumber*)self)->QLCDNumber::inputMethodQuery((Qt::InputMethodQuery)(arg1));
	return new QVariant(__r);
}
int QLCDNumber_focusNextPrevChild_default(void* self, int next) {
	auto&& __r = ((QLCDNumber__Go*)self)->cppgoBase_focusNextPrevChild(next);
	return __r;
}
int QLCDNumber_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QLCDNumber*)self)->QLCDNumber::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QLCDNumber_timerEvent_default(void* self, void* event) {
	((QLCDNumber__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QLCDNumber_childEvent_default(void* self, void* event) {
	((QLCDNumber__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QLCDNumber_customEvent_default(void* self, void* event) {
	((QLCDNumber__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QLCDNumber_connectNotify_default(void* self, void* signal) {
	((QLCDNumber__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QLCDNumber_disconnectNotify_default(void* self, void* signal) {
	((QLCDNumber__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QLCDNumber_connect_overflow(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QLCDNumber*)self, &QLCDNumber::overflow, (QLCDNumber*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QLCDNumber);
	__conns_QLCDNumber[__id] = __c;
}
void QLCDNumber_connect_windowTitleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QLCDNumber*)self, &QLCDNumber::windowTitleChanged, (QLCDNumber*)self, [__id](const QString& title){
		QByteArray title_ba = title.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)title_ba.constData(), title_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QLCDNumber);
	__conns_QLCDNumber[__id] = __c;
}
void QLCDNumber_connect_windowIconChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QLCDNumber*)self, &QLCDNumber::windowIconChanged, (QLCDNumber*)self, [__id](const QIcon& icon){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&icon);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QLCDNumber);
	__conns_QLCDNumber[__id] = __c;
}
void QLCDNumber_connect_windowIconTextChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QLCDNumber*)self, &QLCDNumber::windowIconTextChanged, (QLCDNumber*)self, [__id](const QString& iconText){
		QByteArray iconText_ba = iconText.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)iconText_ba.constData(), iconText_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QLCDNumber);
	__conns_QLCDNumber[__id] = __c;
}
void QLCDNumber_connect_customContextMenuRequested(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QLCDNumber*)self, &QLCDNumber::customContextMenuRequested, (QLCDNumber*)self, [__id](const QPoint& pos){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&pos);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QLCDNumber);
	__conns_QLCDNumber[__id] = __c;
}
void QLCDNumber_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QLCDNumber*)self, &QLCDNumber::destroyed, (QLCDNumber*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QLCDNumber);
	__conns_QLCDNumber[__id] = __c;
}
void QLCDNumber_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QLCDNumber*)self, &QLCDNumber::destroyed, (QLCDNumber*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QLCDNumber);
	__conns_QLCDNumber[__id] = __c;
}
void QLCDNumber_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QLCDNumber*)self, &QLCDNumber::objectNameChanged, (QLCDNumber*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QLCDNumber);
	__conns_QLCDNumber[__id] = __c;
}
void QLCDNumber__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QLCDNumber);
		auto __it = __conns_QLCDNumber.find(id);
		if (__it == __conns_QLCDNumber.end()) return;
		__c = __it->second;
		__conns_QLCDNumber.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QLCDNumber__Go_new(long long id) { return new QLCDNumber__Go(id); }
void QLCDNumber_delete(void* self) { delete (QLCDNumber*)self; }

}

#include <QtWidgets/qspinbox.h>
#include <QtWidgets/qabstractspinbox.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtGui/qvalidator.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qsize.h>
#include <QtCore/qcoreevent.h>
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
#include "gen_qdoublespinbox.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QDoubleSpinBox__Go : public QDoubleSpinBox {
public:
	long long __goId;
	unsigned char __goOn[56] = {};
	template<typename... A> QDoubleSpinBox__Go(long long id, A&&... a) : QDoubleSpinBox(std::forward<A>(a)...), __goId(id) {}
	using QDoubleSpinBox::lineEdit;
	using QDoubleSpinBox::setLineEdit;
	using QDoubleSpinBox::updateMicroFocus;
	using QDoubleSpinBox::create;
	using QDoubleSpinBox::destroy;
	using QDoubleSpinBox::focusNextChild;
	using QDoubleSpinBox::focusPreviousChild;
	using QDoubleSpinBox::sender;
	using QDoubleSpinBox::senderSignalIndex;
	using QDoubleSpinBox::receivers;
	using QDoubleSpinBox::isSignalConnected;
	using QDoubleSpinBox::getDecodedMetricF;
	void cppgoBase_resizeEvent(QResizeEvent* __a0) { QDoubleSpinBox::resizeEvent(__a0); }
	void cppgoBase_keyPressEvent(QKeyEvent* __a0) { QDoubleSpinBox::keyPressEvent(__a0); }
	void cppgoBase_keyReleaseEvent(QKeyEvent* __a0) { QDoubleSpinBox::keyReleaseEvent(__a0); }
	void cppgoBase_wheelEvent(QWheelEvent* __a0) { QDoubleSpinBox::wheelEvent(__a0); }
	void cppgoBase_focusInEvent(QFocusEvent* __a0) { QDoubleSpinBox::focusInEvent(__a0); }
	void cppgoBase_focusOutEvent(QFocusEvent* __a0) { QDoubleSpinBox::focusOutEvent(__a0); }
	void cppgoBase_contextMenuEvent(QContextMenuEvent* __a0) { QDoubleSpinBox::contextMenuEvent(__a0); }
	void cppgoBase_changeEvent(QEvent* __a0) { QDoubleSpinBox::changeEvent(__a0); }
	void cppgoBase_closeEvent(QCloseEvent* __a0) { QDoubleSpinBox::closeEvent(__a0); }
	void cppgoBase_hideEvent(QHideEvent* __a0) { QDoubleSpinBox::hideEvent(__a0); }
	void cppgoBase_mousePressEvent(QMouseEvent* __a0) { QDoubleSpinBox::mousePressEvent(__a0); }
	void cppgoBase_mouseReleaseEvent(QMouseEvent* __a0) { QDoubleSpinBox::mouseReleaseEvent(__a0); }
	void cppgoBase_mouseMoveEvent(QMouseEvent* __a0) { QDoubleSpinBox::mouseMoveEvent(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QDoubleSpinBox::timerEvent(__a0); }
	void cppgoBase_paintEvent(QPaintEvent* __a0) { QDoubleSpinBox::paintEvent(__a0); }
	void cppgoBase_showEvent(QShowEvent* __a0) { QDoubleSpinBox::showEvent(__a0); }
	void cppgoBase_initStyleOption(QStyleOptionSpinBox* __a0) { QDoubleSpinBox::initStyleOption(__a0); }
	QFlags<QAbstractSpinBox::StepEnabledFlag> cppgoBase_stepEnabled() { return QDoubleSpinBox::stepEnabled(); }
	void cppgoBase_mouseDoubleClickEvent(QMouseEvent* __a0) { QDoubleSpinBox::mouseDoubleClickEvent(__a0); }
	void cppgoBase_enterEvent(QEnterEvent* __a0) { QDoubleSpinBox::enterEvent(__a0); }
	void cppgoBase_leaveEvent(QEvent* __a0) { QDoubleSpinBox::leaveEvent(__a0); }
	void cppgoBase_moveEvent(QMoveEvent* __a0) { QDoubleSpinBox::moveEvent(__a0); }
	void cppgoBase_tabletEvent(QTabletEvent* __a0) { QDoubleSpinBox::tabletEvent(__a0); }
	void cppgoBase_actionEvent(QActionEvent* __a0) { QDoubleSpinBox::actionEvent(__a0); }
	void cppgoBase_dragEnterEvent(QDragEnterEvent* __a0) { QDoubleSpinBox::dragEnterEvent(__a0); }
	void cppgoBase_dragMoveEvent(QDragMoveEvent* __a0) { QDoubleSpinBox::dragMoveEvent(__a0); }
	void cppgoBase_dragLeaveEvent(QDragLeaveEvent* __a0) { QDoubleSpinBox::dragLeaveEvent(__a0); }
	void cppgoBase_dropEvent(QDropEvent* __a0) { QDoubleSpinBox::dropEvent(__a0); }
	int cppgoBase_metric(int __a0) { return QDoubleSpinBox::metric((QPaintDevice::PaintDeviceMetric)__a0); }
	void cppgoBase_initPainter(QPainter* __a0) { QDoubleSpinBox::initPainter(__a0); }
	QPaintDevice* cppgoBase_redirected(QPoint* __a0) { return QDoubleSpinBox::redirected(__a0); }
	QPainter* cppgoBase_sharedPainter() { return QDoubleSpinBox::sharedPainter(); }
	void cppgoBase_inputMethodEvent(QInputMethodEvent* __a0) { QDoubleSpinBox::inputMethodEvent(__a0); }
	bool cppgoBase_focusNextPrevChild(bool __a0) { return QDoubleSpinBox::focusNextPrevChild(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QDoubleSpinBox::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QDoubleSpinBox::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QDoubleSpinBox::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QDoubleSpinBox::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QDoubleSpinBox::metaObject();
	}
	QValidator::State validate(QString& input, int& pos) const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_int_voidp_intp) { int __r = ((int(*)(long long, long long, void*, int*, int*))cppgo_virt_int_voidp_intp)(__goId, 6784179835071232555LL, (void*)&input, (int*)&pos, &__h); if (__h) return (QValidator::State)(__r); }
		return QDoubleSpinBox::validate(input, pos);
	}
	double valueFromText(const QString& text) const override {
		int __h = 0;
		QByteArray text_ba = text.toUtf8();
		if (__goOn[2] && cppgo_virt_double_constcharp_longlong) { unsigned long long __b = ((unsigned long long(*)(long long, long long, const char*, long long, int*))cppgo_virt_double_constcharp_longlong)(__goId, 5050482105304000557LL, (char*)text_ba.constData(), text_ba.size(), &__h); if (__h) { double __r; __builtin_memcpy(&__r, &__b, sizeof(__r)); return __r; } }
		return QDoubleSpinBox::valueFromText(text);
	}
	QString textFromValue(double val) const override {
		int __h = 0;
		if (__goOn[3] && cppgo_virtp_charp_unsignedlonglong) {
			long long __len = 0;
			char* __s = ((char*(*)(long long, long long, unsigned long long, long long*, int*))cppgo_virtp_charp_unsignedlonglong)(__goId, -9083693956578760701LL, __builtin_bit_cast(unsigned long long, (double)(val)), &__len, &__h);
			if (__h) { QString __r = QString::fromUtf8(__s, __len); free(__s); return __r; }
		}
		return QDoubleSpinBox::textFromValue(val);
	}
	void fixup(QString& str) const override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4855206404746966715LL, (void*)&str, &__h); if (__h) return; }
		QDoubleSpinBox::fixup(str);
	}
	QSize sizeHint() const override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 729690261457684101LL, &__h); if (__h) return *(QSize*)__r; }
		return QDoubleSpinBox::sizeHint();
	}
	QSize minimumSizeHint() const override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 7049264599837532491LL, &__h); if (__h) return *(QSize*)__r; }
		return QDoubleSpinBox::minimumSizeHint();
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QDoubleSpinBox::event(event);
	}
	QVariant inputMethodQuery(Qt::InputMethodQuery arg1) const override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7345853097340709638LL, (int)(arg1), &__h); if (__h) return *(QVariant*)__r; }
		return QDoubleSpinBox::inputMethodQuery(arg1);
	}
	void stepBy(int steps) override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 7085718643589954634LL, steps, &__h); if (__h) return; }
		QDoubleSpinBox::stepBy(steps);
	}
	void clear() override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, -778567940572639070LL, &__h); if (__h) return; }
		QDoubleSpinBox::clear();
	}
	void resizeEvent(QResizeEvent* event) override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4302917285655014661LL, (void*)event, &__h); if (__h) return; }
		QDoubleSpinBox::resizeEvent(event);
	}
	void keyPressEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1023791977679793185LL, (void*)event, &__h); if (__h) return; }
		QDoubleSpinBox::keyPressEvent(event);
	}
	void keyReleaseEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 772598475950050299LL, (void*)event, &__h); if (__h) return; }
		QDoubleSpinBox::keyReleaseEvent(event);
	}
	void wheelEvent(QWheelEvent* event) override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1919988340934511316LL, (void*)event, &__h); if (__h) return; }
		QDoubleSpinBox::wheelEvent(event);
	}
	void focusInEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 975043767033552552LL, (void*)event, &__h); if (__h) return; }
		QDoubleSpinBox::focusInEvent(event);
	}
	void focusOutEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -385596010576048993LL, (void*)event, &__h); if (__h) return; }
		QDoubleSpinBox::focusOutEvent(event);
	}
	void contextMenuEvent(QContextMenuEvent* event) override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6698043222227006015LL, (void*)event, &__h); if (__h) return; }
		QDoubleSpinBox::contextMenuEvent(event);
	}
	void changeEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4959376063826165943LL, (void*)event, &__h); if (__h) return; }
		QDoubleSpinBox::changeEvent(event);
	}
	void closeEvent(QCloseEvent* event) override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1989937050019464783LL, (void*)event, &__h); if (__h) return; }
		QDoubleSpinBox::closeEvent(event);
	}
	void hideEvent(QHideEvent* event) override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -199718707876719129LL, (void*)event, &__h); if (__h) return; }
		QDoubleSpinBox::hideEvent(event);
	}
	void mousePressEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -332090898078543319LL, (void*)event, &__h); if (__h) return; }
		QDoubleSpinBox::mousePressEvent(event);
	}
	void mouseReleaseEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4415920049297795213LL, (void*)event, &__h); if (__h) return; }
		QDoubleSpinBox::mouseReleaseEvent(event);
	}
	void mouseMoveEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4255494004099806581LL, (void*)event, &__h); if (__h) return; }
		QDoubleSpinBox::mouseMoveEvent(event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QDoubleSpinBox::timerEvent(event);
	}
	void paintEvent(QPaintEvent* event) override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5290442163167458465LL, (void*)event, &__h); if (__h) return; }
		QDoubleSpinBox::paintEvent(event);
	}
	void showEvent(QShowEvent* event) override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6825369174751916LL, (void*)event, &__h); if (__h) return; }
		QDoubleSpinBox::showEvent(event);
	}
	void initStyleOption(QStyleOptionSpinBox* option) const override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -488693769771308531LL, (void*)option, &__h); if (__h) return; }
		QDoubleSpinBox::initStyleOption(option);
	}
	QFlags<QAbstractSpinBox::StepEnabledFlag> stepEnabled() const override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, 5257289574516412244LL, &__h); if (__h) return QFlags<QAbstractSpinBox::StepEnabledFlag>(__r); }
		return QDoubleSpinBox::stepEnabled();
	}
	int devType() const override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8659679893868524044LL, &__h); if (__h) return __r; }
		return QDoubleSpinBox::devType();
	}
	void setVisible(bool visible) override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 2110182719375245575LL, visible, &__h); if (__h) return; }
		QDoubleSpinBox::setVisible(visible);
	}
	int heightForWidth(int arg1) const override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 4311534060495737339LL, arg1, &__h); if (__h) return __r; }
		return QDoubleSpinBox::heightForWidth(arg1);
	}
	bool hasHeightForWidth() const override {
		int __h = 0;
		if (__goOn[32] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8266762397995423915LL, &__h); if (__h) return __r; }
		return QDoubleSpinBox::hasHeightForWidth();
	}
	QPaintEngine* paintEngine() const override {
		int __h = 0;
		if (__goOn[33] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 6398007404756346381LL, &__h); if (__h) return (QPaintEngine*)__r; }
		return QDoubleSpinBox::paintEngine();
	}
	void mouseDoubleClickEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[34] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4029073905697023887LL, (void*)event, &__h); if (__h) return; }
		QDoubleSpinBox::mouseDoubleClickEvent(event);
	}
	void enterEvent(QEnterEvent* event) override {
		int __h = 0;
		if (__goOn[35] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6316270346668237895LL, (void*)event, &__h); if (__h) return; }
		QDoubleSpinBox::enterEvent(event);
	}
	void leaveEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[36] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -551070009623711640LL, (void*)event, &__h); if (__h) return; }
		QDoubleSpinBox::leaveEvent(event);
	}
	void moveEvent(QMoveEvent* event) override {
		int __h = 0;
		if (__goOn[37] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -9161949150489525396LL, (void*)event, &__h); if (__h) return; }
		QDoubleSpinBox::moveEvent(event);
	}
	void tabletEvent(QTabletEvent* event) override {
		int __h = 0;
		if (__goOn[38] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -212068866860701141LL, (void*)event, &__h); if (__h) return; }
		QDoubleSpinBox::tabletEvent(event);
	}
	void actionEvent(QActionEvent* event) override {
		int __h = 0;
		if (__goOn[39] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1697444140088555581LL, (void*)event, &__h); if (__h) return; }
		QDoubleSpinBox::actionEvent(event);
	}
	void dragEnterEvent(QDragEnterEvent* event) override {
		int __h = 0;
		if (__goOn[40] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 206122715919917589LL, (void*)event, &__h); if (__h) return; }
		QDoubleSpinBox::dragEnterEvent(event);
	}
	void dragMoveEvent(QDragMoveEvent* event) override {
		int __h = 0;
		if (__goOn[41] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1781545137352745978LL, (void*)event, &__h); if (__h) return; }
		QDoubleSpinBox::dragMoveEvent(event);
	}
	void dragLeaveEvent(QDragLeaveEvent* event) override {
		int __h = 0;
		if (__goOn[42] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -3139904158373725006LL, (void*)event, &__h); if (__h) return; }
		QDoubleSpinBox::dragLeaveEvent(event);
	}
	void dropEvent(QDropEvent* event) override {
		int __h = 0;
		if (__goOn[43] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 9107058682754591716LL, (void*)event, &__h); if (__h) return; }
		QDoubleSpinBox::dropEvent(event);
	}
	bool nativeEvent(const QByteArray& eventType, void* message, qintptr* result) override {
		int __h = 0;
		if (__goOn[44] && cppgo_virt_int_voidp_voidp_longlongp) { int __r = ((int(*)(long long, long long, void*, void*, long long*, int*))cppgo_virt_int_voidp_voidp_longlongp)(__goId, -2039166586246443706LL, (void*)&eventType, (void*)message, (long long*)result, &__h); if (__h) return __r; }
		return QDoubleSpinBox::nativeEvent(eventType, message, result);
	}
	int metric(QPaintDevice::PaintDeviceMetric arg1) const override {
		int __h = 0;
		if (__goOn[45] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 427629649267311777LL, (int)(arg1), &__h); if (__h) return __r; }
		return QDoubleSpinBox::metric(arg1);
	}
	void initPainter(QPainter* painter) const override {
		int __h = 0;
		if (__goOn[46] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6477500296980050210LL, (void*)painter, &__h); if (__h) return; }
		QDoubleSpinBox::initPainter(painter);
	}
	QPaintDevice* redirected(QPoint* offset) const override {
		int __h = 0;
		if (__goOn[47] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 4127240633310655048LL, (void*)offset, &__h); if (__h) return (QPaintDevice*)__r; }
		return QDoubleSpinBox::redirected(offset);
	}
	QPainter* sharedPainter() const override {
		int __h = 0;
		if (__goOn[48] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 5792771328741774991LL, &__h); if (__h) return (QPainter*)__r; }
		return QDoubleSpinBox::sharedPainter();
	}
	void inputMethodEvent(QInputMethodEvent* arg1) override {
		int __h = 0;
		if (__goOn[49] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5438819661942843222LL, (void*)arg1, &__h); if (__h) return; }
		QDoubleSpinBox::inputMethodEvent(arg1);
	}
	bool focusNextPrevChild(bool next) override {
		int __h = 0;
		if (__goOn[50] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 5898328517505784191LL, next, &__h); if (__h) return __r; }
		return QDoubleSpinBox::focusNextPrevChild(next);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[51] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QDoubleSpinBox::eventFilter(watched, event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[52] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QDoubleSpinBox::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[53] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QDoubleSpinBox::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[54] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QDoubleSpinBox::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[55] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QDoubleSpinBox::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QDoubleSpinBox;
static std::mutex __connsMu_QDoubleSpinBox;

extern "C" {

void* QDoubleSpinBox_base(void* self) { return (void*)static_cast<QAbstractSpinBox*>((QDoubleSpinBox*)self); }
void QDoubleSpinBox_chain(void* self, void** out) {
	QDoubleSpinBox* __p = (QDoubleSpinBox*)self;
	QAbstractSpinBox* __b0 = static_cast<QAbstractSpinBox*>(__p); out[0] = (void*)__b0;
	QWidget* __b1 = static_cast<QWidget*>(__b0); out[1] = (void*)__b1;
	QObject* __b2 = static_cast<QObject*>(__b1); out[2] = (void*)__b2;
}
void QDoubleSpinBox__set_override(void* self, int idx) { ((QDoubleSpinBox__Go*)self)->__goOn[idx] = 1; }
void* QDoubleSpinBox_fromQObject(void* o) { return (void*)qobject_cast<QDoubleSpinBox*>((QObject*)o); }
void* QDoubleSpinBox_new(long long __goid, void* parent) {
	return new QDoubleSpinBox__Go(__goid, (QWidget*)parent);
}
void* QDoubleSpinBox_new2(long long __goid) {
	return new QDoubleSpinBox__Go(__goid);
}
char* QDoubleSpinBox_tr(const char* s, long long* __out_len) {
	auto&& __r = QDoubleSpinBox::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
double QDoubleSpinBox_value(void* self) {
	auto&& __r = ((QDoubleSpinBox*)self)->value();
	return __r;
}
char* QDoubleSpinBox_prefix(void* self, long long* __out_len) {
	auto&& __r = ((QDoubleSpinBox*)self)->prefix();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QDoubleSpinBox_setPrefix(void* self, const char* prefix, long long prefix_len) {
	QString prefix_qstr = QString::fromUtf8(prefix, prefix_len);
	((QDoubleSpinBox*)self)->setPrefix(prefix_qstr);
}
char* QDoubleSpinBox_suffix(void* self, long long* __out_len) {
	auto&& __r = ((QDoubleSpinBox*)self)->suffix();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QDoubleSpinBox_setSuffix(void* self, const char* suffix, long long suffix_len) {
	QString suffix_qstr = QString::fromUtf8(suffix, suffix_len);
	((QDoubleSpinBox*)self)->setSuffix(suffix_qstr);
}
char* QDoubleSpinBox_cleanText(void* self, long long* __out_len) {
	auto&& __r = ((QDoubleSpinBox*)self)->cleanText();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
double QDoubleSpinBox_singleStep(void* self) {
	auto&& __r = ((QDoubleSpinBox*)self)->singleStep();
	return __r;
}
void QDoubleSpinBox_setSingleStep(void* self, double val) {
	((QDoubleSpinBox*)self)->setSingleStep(val);
}
double QDoubleSpinBox_minimum(void* self) {
	auto&& __r = ((QDoubleSpinBox*)self)->minimum();
	return __r;
}
void QDoubleSpinBox_setMinimum(void* self, double min) {
	((QDoubleSpinBox*)self)->setMinimum(min);
}
double QDoubleSpinBox_maximum(void* self) {
	auto&& __r = ((QDoubleSpinBox*)self)->maximum();
	return __r;
}
void QDoubleSpinBox_setMaximum(void* self, double max) {
	((QDoubleSpinBox*)self)->setMaximum(max);
}
void QDoubleSpinBox_setRange(void* self, double min, double max) {
	((QDoubleSpinBox*)self)->setRange(min, max);
}
int QDoubleSpinBox_stepType(void* self) {
	auto&& __r = ((QDoubleSpinBox*)self)->stepType();
	return (int)(__r);
}
void QDoubleSpinBox_setStepType(void* self, int stepType) {
	((QDoubleSpinBox*)self)->setStepType((QAbstractSpinBox::StepType)(stepType));
}
int QDoubleSpinBox_decimals(void* self) {
	auto&& __r = ((QDoubleSpinBox*)self)->decimals();
	return __r;
}
void QDoubleSpinBox_setDecimals(void* self, int prec) {
	((QDoubleSpinBox*)self)->setDecimals(prec);
}
double QDoubleSpinBox_valueFromText(void* self, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = ((QDoubleSpinBox*)self)->valueFromText(text_qstr);
	return __r;
}
char* QDoubleSpinBox_textFromValue(void* self, double val, long long* __out_len) {
	auto&& __r = ((QDoubleSpinBox*)self)->textFromValue(val);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QDoubleSpinBox_setValue(void* self, double val) {
	((QDoubleSpinBox*)self)->setValue(val);
}
void QDoubleSpinBox_valueChanged(void* self, double arg1) {
	((QDoubleSpinBox*)self)->valueChanged(arg1);
}
void QDoubleSpinBox_textChanged(void* self, const char* arg1, long long arg1_len) {
	QString arg1_qstr = QString::fromUtf8(arg1, arg1_len);
	((QDoubleSpinBox*)self)->textChanged(arg1_qstr);
}
void* QDoubleSpinBox_lineEdit(void* self) {
	auto&& __r = ((QDoubleSpinBox__Go*)self)->lineEdit();
	return (void*)(__r);
}
void QDoubleSpinBox_setLineEdit(void* self, void* edit) {
	((QDoubleSpinBox__Go*)self)->setLineEdit((QLineEdit*)edit);
}
void QDoubleSpinBox_setTabOrder(void* arg1, void* arg2) {
	QDoubleSpinBox::setTabOrder((QWidget*)arg1, (QWidget*)arg2);
}
void* QDoubleSpinBox_mouseGrabber() {
	auto&& __r = QDoubleSpinBox::mouseGrabber();
	return (void*)(__r);
}
void* QDoubleSpinBox_keyboardGrabber() {
	auto&& __r = QDoubleSpinBox::keyboardGrabber();
	return (void*)(__r);
}
void* QDoubleSpinBox_find(unsigned long long arg1) {
	auto&& __r = QDoubleSpinBox::find(arg1);
	return (void*)(__r);
}
void* QDoubleSpinBox_createWindowContainer(void* window) {
	auto&& __r = QDoubleSpinBox::createWindowContainer((QWindow*)window);
	return (void*)(__r);
}
void QDoubleSpinBox_updateMicroFocus(void* self) {
	((QDoubleSpinBox__Go*)self)->updateMicroFocus();
}
void QDoubleSpinBox_create(void* self) {
	((QDoubleSpinBox__Go*)self)->create();
}
void QDoubleSpinBox_destroy(void* self) {
	((QDoubleSpinBox__Go*)self)->destroy();
}
int QDoubleSpinBox_focusNextChild(void* self) {
	auto&& __r = ((QDoubleSpinBox__Go*)self)->focusNextChild();
	return __r;
}
int QDoubleSpinBox_focusPreviousChild(void* self) {
	auto&& __r = ((QDoubleSpinBox__Go*)self)->focusPreviousChild();
	return __r;
}
void* QDoubleSpinBox_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QDoubleSpinBox::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QDoubleSpinBox_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QDoubleSpinBox::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QDoubleSpinBox_sender(void* self) {
	auto&& __r = ((QDoubleSpinBox__Go*)self)->sender();
	return (void*)(__r);
}
int QDoubleSpinBox_senderSignalIndex(void* self) {
	auto&& __r = ((QDoubleSpinBox__Go*)self)->senderSignalIndex();
	return __r;
}
int QDoubleSpinBox_receivers(void* self, const char* signal) {
	auto&& __r = ((QDoubleSpinBox__Go*)self)->receivers(signal);
	return __r;
}
int QDoubleSpinBox_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QDoubleSpinBox__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
double QDoubleSpinBox_devicePixelRatioFScale() {
	auto&& __r = QDoubleSpinBox::devicePixelRatioFScale();
	return __r;
}
int QDoubleSpinBox_encodeMetricF(int metric, double value) {
	auto&& __r = QDoubleSpinBox::encodeMetricF((QPaintDevice::PaintDeviceMetric)(metric), value);
	return __r;
}
double QDoubleSpinBox_getDecodedMetricF(void* self, int metricA, int metricB) {
	auto&& __r = ((QDoubleSpinBox__Go*)self)->getDecodedMetricF((QPaintDevice::PaintDeviceMetric)(metricA), (QPaintDevice::PaintDeviceMetric)(metricB));
	return __r;
}
char* QDoubleSpinBox_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QDoubleSpinBox::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QDoubleSpinBox_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QDoubleSpinBox::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QDoubleSpinBox_createWindowContainer2(void* window, void* parent) {
	auto&& __r = QDoubleSpinBox::createWindowContainer((QWindow*)window, (QWidget*)parent);
	return (void*)(__r);
}
void* QDoubleSpinBox_createWindowContainer3(void* window, void* parent, int flags) {
	auto&& __r = QDoubleSpinBox::createWindowContainer((QWindow*)window, (QWidget*)parent, QFlags<Qt::WindowType>(flags));
	return (void*)(__r);
}
void QDoubleSpinBox_updateMicroFocus2(void* self, int query) {
	((QDoubleSpinBox__Go*)self)->updateMicroFocus((Qt::InputMethodQuery)(query));
}
void QDoubleSpinBox_create2(void* self, unsigned long long arg1) {
	((QDoubleSpinBox__Go*)self)->create(arg1);
}
void QDoubleSpinBox_create3(void* self, unsigned long long arg1, int initializeWindow) {
	((QDoubleSpinBox__Go*)self)->create(arg1, initializeWindow);
}
void QDoubleSpinBox_create4(void* self, unsigned long long arg1, int initializeWindow, int destroyOldWindow) {
	((QDoubleSpinBox__Go*)self)->create(arg1, initializeWindow, destroyOldWindow);
}
void QDoubleSpinBox_destroy2(void* self, int destroyWindow) {
	((QDoubleSpinBox__Go*)self)->destroy(destroyWindow);
}
void QDoubleSpinBox_destroy3(void* self, int destroyWindow, int destroySubWindows) {
	((QDoubleSpinBox__Go*)self)->destroy(destroyWindow, destroySubWindows);
}
void* QDoubleSpinBox_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QDoubleSpinBox::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QDoubleSpinBox_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QDoubleSpinBox::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
double QDoubleSpinBox_valueFromText_default(void* self, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = ((QDoubleSpinBox*)self)->QDoubleSpinBox::valueFromText(text_qstr);
	return __r;
}
char* QDoubleSpinBox_textFromValue_default(void* self, double val, long long* __out_len) {
	auto&& __r = ((QDoubleSpinBox*)self)->QDoubleSpinBox::textFromValue(val);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QDoubleSpinBox_metaObject_default(void* self) {
	auto&& __r = ((QDoubleSpinBox*)self)->QDoubleSpinBox::metaObject();
	return (void*)(__r);
}
int QDoubleSpinBox_validate_default(void* self, void* input, int* pos) {
	int pos_ref = *pos;
	auto&& __r = ((QDoubleSpinBox*)self)->QDoubleSpinBox::validate(*(QString*)input, pos_ref);
	*pos = pos_ref;
	return (int)(__r);
}
void QDoubleSpinBox_fixup_default(void* self, void* str) {
	((QDoubleSpinBox*)self)->QDoubleSpinBox::fixup(*(QString*)str);
}
void* QDoubleSpinBox_sizeHint_default(void* self) {
	auto&& __r = ((QDoubleSpinBox*)self)->QDoubleSpinBox::sizeHint();
	return new QSize(__r);
}
void* QDoubleSpinBox_minimumSizeHint_default(void* self) {
	auto&& __r = ((QDoubleSpinBox*)self)->QDoubleSpinBox::minimumSizeHint();
	return new QSize(__r);
}
int QDoubleSpinBox_event_default(void* self, void* event) {
	auto&& __r = ((QDoubleSpinBox*)self)->QDoubleSpinBox::event((QEvent*)event);
	return __r;
}
void* QDoubleSpinBox_inputMethodQuery_default(void* self, int arg1) {
	auto&& __r = ((QDoubleSpinBox*)self)->QDoubleSpinBox::inputMethodQuery((Qt::InputMethodQuery)(arg1));
	return new QVariant(__r);
}
void QDoubleSpinBox_stepBy_default(void* self, int steps) {
	((QDoubleSpinBox*)self)->QDoubleSpinBox::stepBy(steps);
}
void QDoubleSpinBox_clear_default(void* self) {
	((QDoubleSpinBox*)self)->QDoubleSpinBox::clear();
}
void QDoubleSpinBox_resizeEvent_default(void* self, void* event) {
	((QDoubleSpinBox__Go*)self)->cppgoBase_resizeEvent((QResizeEvent*)event);
}
void QDoubleSpinBox_keyPressEvent_default(void* self, void* event) {
	((QDoubleSpinBox__Go*)self)->cppgoBase_keyPressEvent((QKeyEvent*)event);
}
void QDoubleSpinBox_keyReleaseEvent_default(void* self, void* event) {
	((QDoubleSpinBox__Go*)self)->cppgoBase_keyReleaseEvent((QKeyEvent*)event);
}
void QDoubleSpinBox_wheelEvent_default(void* self, void* event) {
	((QDoubleSpinBox__Go*)self)->cppgoBase_wheelEvent((QWheelEvent*)event);
}
void QDoubleSpinBox_focusInEvent_default(void* self, void* event) {
	((QDoubleSpinBox__Go*)self)->cppgoBase_focusInEvent((QFocusEvent*)event);
}
void QDoubleSpinBox_focusOutEvent_default(void* self, void* event) {
	((QDoubleSpinBox__Go*)self)->cppgoBase_focusOutEvent((QFocusEvent*)event);
}
void QDoubleSpinBox_contextMenuEvent_default(void* self, void* event) {
	((QDoubleSpinBox__Go*)self)->cppgoBase_contextMenuEvent((QContextMenuEvent*)event);
}
void QDoubleSpinBox_changeEvent_default(void* self, void* event) {
	((QDoubleSpinBox__Go*)self)->cppgoBase_changeEvent((QEvent*)event);
}
void QDoubleSpinBox_closeEvent_default(void* self, void* event) {
	((QDoubleSpinBox__Go*)self)->cppgoBase_closeEvent((QCloseEvent*)event);
}
void QDoubleSpinBox_hideEvent_default(void* self, void* event) {
	((QDoubleSpinBox__Go*)self)->cppgoBase_hideEvent((QHideEvent*)event);
}
void QDoubleSpinBox_mousePressEvent_default(void* self, void* event) {
	((QDoubleSpinBox__Go*)self)->cppgoBase_mousePressEvent((QMouseEvent*)event);
}
void QDoubleSpinBox_mouseReleaseEvent_default(void* self, void* event) {
	((QDoubleSpinBox__Go*)self)->cppgoBase_mouseReleaseEvent((QMouseEvent*)event);
}
void QDoubleSpinBox_mouseMoveEvent_default(void* self, void* event) {
	((QDoubleSpinBox__Go*)self)->cppgoBase_mouseMoveEvent((QMouseEvent*)event);
}
void QDoubleSpinBox_timerEvent_default(void* self, void* event) {
	((QDoubleSpinBox__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QDoubleSpinBox_paintEvent_default(void* self, void* event) {
	((QDoubleSpinBox__Go*)self)->cppgoBase_paintEvent((QPaintEvent*)event);
}
void QDoubleSpinBox_showEvent_default(void* self, void* event) {
	((QDoubleSpinBox__Go*)self)->cppgoBase_showEvent((QShowEvent*)event);
}
void QDoubleSpinBox_initStyleOption_default(void* self, void* option) {
	((QDoubleSpinBox__Go*)self)->cppgoBase_initStyleOption((QStyleOptionSpinBox*)option);
}
int QDoubleSpinBox_stepEnabled_default(void* self) {
	auto&& __r = ((QDoubleSpinBox__Go*)self)->cppgoBase_stepEnabled();
	return (int)(__r);
}
int QDoubleSpinBox_devType_default(void* self) {
	auto&& __r = ((QDoubleSpinBox*)self)->QDoubleSpinBox::devType();
	return __r;
}
void QDoubleSpinBox_setVisible_default(void* self, int visible) {
	((QDoubleSpinBox*)self)->QDoubleSpinBox::setVisible(visible);
}
int QDoubleSpinBox_heightForWidth_default(void* self, int arg1) {
	auto&& __r = ((QDoubleSpinBox*)self)->QDoubleSpinBox::heightForWidth(arg1);
	return __r;
}
int QDoubleSpinBox_hasHeightForWidth_default(void* self) {
	auto&& __r = ((QDoubleSpinBox*)self)->QDoubleSpinBox::hasHeightForWidth();
	return __r;
}
void* QDoubleSpinBox_paintEngine_default(void* self) {
	auto&& __r = ((QDoubleSpinBox*)self)->QDoubleSpinBox::paintEngine();
	return (void*)(__r);
}
void QDoubleSpinBox_mouseDoubleClickEvent_default(void* self, void* event) {
	((QDoubleSpinBox__Go*)self)->cppgoBase_mouseDoubleClickEvent((QMouseEvent*)event);
}
void QDoubleSpinBox_enterEvent_default(void* self, void* event) {
	((QDoubleSpinBox__Go*)self)->cppgoBase_enterEvent((QEnterEvent*)event);
}
void QDoubleSpinBox_leaveEvent_default(void* self, void* event) {
	((QDoubleSpinBox__Go*)self)->cppgoBase_leaveEvent((QEvent*)event);
}
void QDoubleSpinBox_moveEvent_default(void* self, void* event) {
	((QDoubleSpinBox__Go*)self)->cppgoBase_moveEvent((QMoveEvent*)event);
}
void QDoubleSpinBox_tabletEvent_default(void* self, void* event) {
	((QDoubleSpinBox__Go*)self)->cppgoBase_tabletEvent((QTabletEvent*)event);
}
void QDoubleSpinBox_actionEvent_default(void* self, void* event) {
	((QDoubleSpinBox__Go*)self)->cppgoBase_actionEvent((QActionEvent*)event);
}
void QDoubleSpinBox_dragEnterEvent_default(void* self, void* event) {
	((QDoubleSpinBox__Go*)self)->cppgoBase_dragEnterEvent((QDragEnterEvent*)event);
}
void QDoubleSpinBox_dragMoveEvent_default(void* self, void* event) {
	((QDoubleSpinBox__Go*)self)->cppgoBase_dragMoveEvent((QDragMoveEvent*)event);
}
void QDoubleSpinBox_dragLeaveEvent_default(void* self, void* event) {
	((QDoubleSpinBox__Go*)self)->cppgoBase_dragLeaveEvent((QDragLeaveEvent*)event);
}
void QDoubleSpinBox_dropEvent_default(void* self, void* event) {
	((QDoubleSpinBox__Go*)self)->cppgoBase_dropEvent((QDropEvent*)event);
}
int QDoubleSpinBox_metric_default(void* self, int arg1) {
	auto&& __r = ((QDoubleSpinBox__Go*)self)->cppgoBase_metric(arg1);
	return __r;
}
void QDoubleSpinBox_initPainter_default(void* self, void* painter) {
	((QDoubleSpinBox__Go*)self)->cppgoBase_initPainter((QPainter*)painter);
}
void* QDoubleSpinBox_redirected_default(void* self, void* offset) {
	auto&& __r = ((QDoubleSpinBox__Go*)self)->cppgoBase_redirected((QPoint*)offset);
	return (void*)(__r);
}
void* QDoubleSpinBox_sharedPainter_default(void* self) {
	auto&& __r = ((QDoubleSpinBox__Go*)self)->cppgoBase_sharedPainter();
	return (void*)(__r);
}
void QDoubleSpinBox_inputMethodEvent_default(void* self, void* arg1) {
	((QDoubleSpinBox__Go*)self)->cppgoBase_inputMethodEvent((QInputMethodEvent*)arg1);
}
int QDoubleSpinBox_focusNextPrevChild_default(void* self, int next) {
	auto&& __r = ((QDoubleSpinBox__Go*)self)->cppgoBase_focusNextPrevChild(next);
	return __r;
}
int QDoubleSpinBox_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QDoubleSpinBox*)self)->QDoubleSpinBox::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QDoubleSpinBox_childEvent_default(void* self, void* event) {
	((QDoubleSpinBox__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QDoubleSpinBox_customEvent_default(void* self, void* event) {
	((QDoubleSpinBox__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QDoubleSpinBox_connectNotify_default(void* self, void* signal) {
	((QDoubleSpinBox__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QDoubleSpinBox_disconnectNotify_default(void* self, void* signal) {
	((QDoubleSpinBox__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QDoubleSpinBox_connect_valueChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDoubleSpinBox*)self, &QDoubleSpinBox::valueChanged, (QDoubleSpinBox*)self, [__id](double arg1){
		if (cppgo_sig_void_unsignedlonglong) ((void(*)(long long, unsigned long long))cppgo_sig_void_unsignedlonglong)(__id, __builtin_bit_cast(unsigned long long, (double)(arg1)));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDoubleSpinBox);
	__conns_QDoubleSpinBox[__id] = __c;
}
void QDoubleSpinBox_connect_textChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDoubleSpinBox*)self, &QDoubleSpinBox::textChanged, (QDoubleSpinBox*)self, [__id](const QString& arg1){
		QByteArray arg1_ba = arg1.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)arg1_ba.constData(), arg1_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDoubleSpinBox);
	__conns_QDoubleSpinBox[__id] = __c;
}
void QDoubleSpinBox_connect_editingFinished(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDoubleSpinBox*)self, &QDoubleSpinBox::editingFinished, (QDoubleSpinBox*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDoubleSpinBox);
	__conns_QDoubleSpinBox[__id] = __c;
}
void QDoubleSpinBox_connect_returnPressed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDoubleSpinBox*)self, &QDoubleSpinBox::returnPressed, (QDoubleSpinBox*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDoubleSpinBox);
	__conns_QDoubleSpinBox[__id] = __c;
}
void QDoubleSpinBox_connect_windowTitleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDoubleSpinBox*)self, &QDoubleSpinBox::windowTitleChanged, (QDoubleSpinBox*)self, [__id](const QString& title){
		QByteArray title_ba = title.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)title_ba.constData(), title_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDoubleSpinBox);
	__conns_QDoubleSpinBox[__id] = __c;
}
void QDoubleSpinBox_connect_windowIconChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDoubleSpinBox*)self, &QDoubleSpinBox::windowIconChanged, (QDoubleSpinBox*)self, [__id](const QIcon& icon){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&icon);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDoubleSpinBox);
	__conns_QDoubleSpinBox[__id] = __c;
}
void QDoubleSpinBox_connect_windowIconTextChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDoubleSpinBox*)self, &QDoubleSpinBox::windowIconTextChanged, (QDoubleSpinBox*)self, [__id](const QString& iconText){
		QByteArray iconText_ba = iconText.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)iconText_ba.constData(), iconText_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDoubleSpinBox);
	__conns_QDoubleSpinBox[__id] = __c;
}
void QDoubleSpinBox_connect_customContextMenuRequested(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDoubleSpinBox*)self, &QDoubleSpinBox::customContextMenuRequested, (QDoubleSpinBox*)self, [__id](const QPoint& pos){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&pos);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDoubleSpinBox);
	__conns_QDoubleSpinBox[__id] = __c;
}
void QDoubleSpinBox_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDoubleSpinBox*)self, &QDoubleSpinBox::destroyed, (QDoubleSpinBox*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDoubleSpinBox);
	__conns_QDoubleSpinBox[__id] = __c;
}
void QDoubleSpinBox_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDoubleSpinBox*)self, &QDoubleSpinBox::destroyed, (QDoubleSpinBox*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDoubleSpinBox);
	__conns_QDoubleSpinBox[__id] = __c;
}
void QDoubleSpinBox_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDoubleSpinBox*)self, &QDoubleSpinBox::objectNameChanged, (QDoubleSpinBox*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDoubleSpinBox);
	__conns_QDoubleSpinBox[__id] = __c;
}
void QDoubleSpinBox__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QDoubleSpinBox);
		auto __it = __conns_QDoubleSpinBox.find(id);
		if (__it == __conns_QDoubleSpinBox.end()) return;
		__c = __it->second;
		__conns_QDoubleSpinBox.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QDoubleSpinBox__Go_new(long long id) { return new QDoubleSpinBox__Go(id); }
void QDoubleSpinBox_delete(void* self) { delete (QDoubleSpinBox*)self; }

}

#include <QtWidgets/qtreewidget.h>
#include <QtWidgets/qtreeview.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qitemselectionmodel.h>
#include <QtCore/qpoint.h>
#include <QtCore/qrect.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qabstractitemmodel.h>
#include <QtWidgets/qabstractitemview.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qmimedata.h>
#include <QtGui/qevent.h>
#include <QtWidgets/qheaderview.h>
#include <QtGui/qregion.h>
#include <QtGui/qpainter.h>
#include <QtWidgets/qstyleoption.h>
#include <QtCore/qsize.h>
#include <QtWidgets/qabstractitemdelegate.h>
#include <QtCore/qvariant.h>
#include <QtCore/qobject.h>
#include <QtWidgets/qscrollbar.h>
#include <QtWidgets/qabstractscrollarea.h>
#include <QtCore/qmargins.h>
#include <QtWidgets/qframe.h>
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
#include "gen_qtreewidget.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QTreeWidget__Go : public QTreeWidget {
public:
	long long __goId;
	unsigned char __goOn[95] = {};
	template<typename... A> QTreeWidget__Go(long long id, A&&... a) : QTreeWidget(std::forward<A>(a)...), __goId(id) {}
	using QTreeWidget::columnResized;
	using QTreeWidget::columnCountChanged;
	using QTreeWidget::columnMoved;
	using QTreeWidget::reexpand;
	using QTreeWidget::rowsRemoved;
	using QTreeWidget::drawTree;
	using QTreeWidget::indexRowSizeHint;
	using QTreeWidget::rowHeight;
	using QTreeWidget::state;
	using QTreeWidget::scheduleDelayedItemsLayout;
	using QTreeWidget::executeDelayedItemsLayout;
	using QTreeWidget::setDirtyRegion;
	using QTreeWidget::scrollDirtyRegion;
	using QTreeWidget::dirtyRegionOffset;
	using QTreeWidget::startAutoScroll;
	using QTreeWidget::stopAutoScroll;
	using QTreeWidget::doAutoScroll;
	using QTreeWidget::dropIndicatorPosition;
	using QTreeWidget::setViewportMargins;
	using QTreeWidget::viewportMargins;
	using QTreeWidget::drawFrame;
	using QTreeWidget::updateMicroFocus;
	using QTreeWidget::create;
	using QTreeWidget::destroy;
	using QTreeWidget::focusNextChild;
	using QTreeWidget::focusPreviousChild;
	using QTreeWidget::sender;
	using QTreeWidget::senderSignalIndex;
	using QTreeWidget::receivers;
	using QTreeWidget::isSignalConnected;
	using QTreeWidget::getDecodedMetricF;
	bool cppgoBase_event(QEvent* __a0) { return QTreeWidget::event(__a0); }
	QList<QString> cppgoBase_mimeTypes() { return QTreeWidget::mimeTypes(); }
	bool cppgoBase_dropMimeData(QTreeWidgetItem* __a0, int __a1, const QMimeData* __a2, int __a3) { return QTreeWidget::dropMimeData(__a0, __a1, __a2, (Qt::DropAction)__a3); }
	QFlags<Qt::DropAction> cppgoBase_supportedDropActions() { return QTreeWidget::supportedDropActions(); }
	void cppgoBase_dropEvent(QDropEvent* __a0) { QTreeWidget::dropEvent(__a0); }
	void cppgoBase_verticalScrollbarValueChanged(int __a0) { QTreeWidget::verticalScrollbarValueChanged(__a0); }
	void cppgoBase_scrollContentsBy(int __a0, int __a1) { QTreeWidget::scrollContentsBy(__a0, __a1); }
	void cppgoBase_rowsInserted(const QModelIndex& __a0, int __a1, int __a2) { QTreeWidget::rowsInserted(__a0, __a1, __a2); }
	void cppgoBase_rowsAboutToBeRemoved(const QModelIndex& __a0, int __a1, int __a2) { QTreeWidget::rowsAboutToBeRemoved(__a0, __a1, __a2); }
	QModelIndex cppgoBase_moveCursor(int __a0, QFlags<Qt::KeyboardModifier> __a1) { return QTreeWidget::moveCursor((QAbstractItemView::CursorAction)__a0, __a1); }
	int cppgoBase_horizontalOffset() { return QTreeWidget::horizontalOffset(); }
	int cppgoBase_verticalOffset() { return QTreeWidget::verticalOffset(); }
	void cppgoBase_setSelection(const QRect& __a0, QFlags<QItemSelectionModel::SelectionFlag> __a1) { QTreeWidget::setSelection(__a0, __a1); }
	QRegion cppgoBase_visualRegionForSelection(const QItemSelection& __a0) { return QTreeWidget::visualRegionForSelection(__a0); }
	QList<QModelIndex> cppgoBase_selectedIndexes() { return QTreeWidget::selectedIndexes(); }
	void cppgoBase_changeEvent(QEvent* __a0) { QTreeWidget::changeEvent(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QTreeWidget::timerEvent(__a0); }
	void cppgoBase_paintEvent(QPaintEvent* __a0) { QTreeWidget::paintEvent(__a0); }
	void cppgoBase_drawRow(QPainter* __a0, const QStyleOptionViewItem& __a1, const QModelIndex& __a2) { QTreeWidget::drawRow(__a0, __a1, __a2); }
	void cppgoBase_drawBranches(QPainter* __a0, const QRect& __a1, const QModelIndex& __a2) { QTreeWidget::drawBranches(__a0, __a1, __a2); }
	void cppgoBase_mousePressEvent(QMouseEvent* __a0) { QTreeWidget::mousePressEvent(__a0); }
	void cppgoBase_mouseReleaseEvent(QMouseEvent* __a0) { QTreeWidget::mouseReleaseEvent(__a0); }
	void cppgoBase_mouseDoubleClickEvent(QMouseEvent* __a0) { QTreeWidget::mouseDoubleClickEvent(__a0); }
	void cppgoBase_mouseMoveEvent(QMouseEvent* __a0) { QTreeWidget::mouseMoveEvent(__a0); }
	void cppgoBase_keyPressEvent(QKeyEvent* __a0) { QTreeWidget::keyPressEvent(__a0); }
	void cppgoBase_dragMoveEvent(QDragMoveEvent* __a0) { QTreeWidget::dragMoveEvent(__a0); }
	bool cppgoBase_viewportEvent(QEvent* __a0) { return QTreeWidget::viewportEvent(__a0); }
	void cppgoBase_updateGeometries() { QTreeWidget::updateGeometries(); }
	QSize cppgoBase_viewportSizeHint() { return QTreeWidget::viewportSizeHint(); }
	int cppgoBase_sizeHintForColumn(int __a0) { return QTreeWidget::sizeHintForColumn(__a0); }
	void cppgoBase_horizontalScrollbarAction(int __a0) { QTreeWidget::horizontalScrollbarAction(__a0); }
	bool cppgoBase_isIndexHidden(const QModelIndex& __a0) { return QTreeWidget::isIndexHidden(__a0); }
	void cppgoBase_selectionChanged(const QItemSelection& __a0, const QItemSelection& __a1) { QTreeWidget::selectionChanged(__a0, __a1); }
	void cppgoBase_currentChanged(const QModelIndex& __a0, const QModelIndex& __a1) { QTreeWidget::currentChanged(__a0, __a1); }
	void cppgoBase_updateEditorData() { QTreeWidget::updateEditorData(); }
	void cppgoBase_updateEditorGeometries() { QTreeWidget::updateEditorGeometries(); }
	void cppgoBase_verticalScrollbarAction(int __a0) { QTreeWidget::verticalScrollbarAction(__a0); }
	void cppgoBase_horizontalScrollbarValueChanged(int __a0) { QTreeWidget::horizontalScrollbarValueChanged(__a0); }
	void cppgoBase_closeEditor(QWidget* __a0, int __a1) { QTreeWidget::closeEditor(__a0, (QAbstractItemDelegate::EndEditHint)__a1); }
	void cppgoBase_commitData(QWidget* __a0) { QTreeWidget::commitData(__a0); }
	void cppgoBase_editorDestroyed(QObject* __a0) { QTreeWidget::editorDestroyed(__a0); }
	bool cppgoBase_edit(const QModelIndex& __a0, int __a1, QEvent* __a2) { return QTreeWidget::edit(__a0, (QAbstractItemView::EditTrigger)__a1, __a2); }
	QFlags<QItemSelectionModel::SelectionFlag> cppgoBase_selectionCommand(const QModelIndex& __a0, const QEvent* __a1) { return QTreeWidget::selectionCommand(__a0, __a1); }
	void cppgoBase_startDrag(QFlags<Qt::DropAction> __a0) { QTreeWidget::startDrag(__a0); }
	void cppgoBase_initViewItemOption(QStyleOptionViewItem* __a0) { QTreeWidget::initViewItemOption(__a0); }
	bool cppgoBase_focusNextPrevChild(bool __a0) { return QTreeWidget::focusNextPrevChild(__a0); }
	void cppgoBase_dragEnterEvent(QDragEnterEvent* __a0) { QTreeWidget::dragEnterEvent(__a0); }
	void cppgoBase_dragLeaveEvent(QDragLeaveEvent* __a0) { QTreeWidget::dragLeaveEvent(__a0); }
	void cppgoBase_focusInEvent(QFocusEvent* __a0) { QTreeWidget::focusInEvent(__a0); }
	void cppgoBase_focusOutEvent(QFocusEvent* __a0) { QTreeWidget::focusOutEvent(__a0); }
	void cppgoBase_resizeEvent(QResizeEvent* __a0) { QTreeWidget::resizeEvent(__a0); }
	void cppgoBase_inputMethodEvent(QInputMethodEvent* __a0) { QTreeWidget::inputMethodEvent(__a0); }
	bool cppgoBase_eventFilter(QObject* __a0, QEvent* __a1) { return QTreeWidget::eventFilter(__a0, __a1); }
	void cppgoBase_wheelEvent(QWheelEvent* __a0) { QTreeWidget::wheelEvent(__a0); }
	void cppgoBase_contextMenuEvent(QContextMenuEvent* __a0) { QTreeWidget::contextMenuEvent(__a0); }
	void cppgoBase_initStyleOption(QStyleOptionFrame* __a0) { QTreeWidget::initStyleOption(__a0); }
	void cppgoBase_keyReleaseEvent(QKeyEvent* __a0) { QTreeWidget::keyReleaseEvent(__a0); }
	void cppgoBase_enterEvent(QEnterEvent* __a0) { QTreeWidget::enterEvent(__a0); }
	void cppgoBase_leaveEvent(QEvent* __a0) { QTreeWidget::leaveEvent(__a0); }
	void cppgoBase_moveEvent(QMoveEvent* __a0) { QTreeWidget::moveEvent(__a0); }
	void cppgoBase_closeEvent(QCloseEvent* __a0) { QTreeWidget::closeEvent(__a0); }
	void cppgoBase_tabletEvent(QTabletEvent* __a0) { QTreeWidget::tabletEvent(__a0); }
	void cppgoBase_actionEvent(QActionEvent* __a0) { QTreeWidget::actionEvent(__a0); }
	void cppgoBase_showEvent(QShowEvent* __a0) { QTreeWidget::showEvent(__a0); }
	void cppgoBase_hideEvent(QHideEvent* __a0) { QTreeWidget::hideEvent(__a0); }
	int cppgoBase_metric(int __a0) { return QTreeWidget::metric((QPaintDevice::PaintDeviceMetric)__a0); }
	void cppgoBase_initPainter(QPainter* __a0) { QTreeWidget::initPainter(__a0); }
	QPaintDevice* cppgoBase_redirected(QPoint* __a0) { return QTreeWidget::redirected(__a0); }
	QPainter* cppgoBase_sharedPainter() { return QTreeWidget::sharedPainter(); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QTreeWidget::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QTreeWidget::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QTreeWidget::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QTreeWidget::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QTreeWidget::metaObject();
	}
	void setSelectionModel(QItemSelectionModel* selectionModel) override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4689322347236366780LL, (void*)selectionModel, &__h); if (__h) return; }
		QTreeWidget::setSelectionModel(selectionModel);
	}
	bool event(QEvent* e) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)e, &__h); if (__h) return __r; }
		return QTreeWidget::event(e);
	}
	QList<QString> mimeTypes() const override {
		int __h = 0;
		if (__goOn[3] && cppgo_virtp_voidp) {
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
		return QTreeWidget::mimeTypes();
	}
	bool dropMimeData(QTreeWidgetItem* parent, int index, const QMimeData* data, Qt::DropAction action) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_int_voidp_int_voidp_int) { int __r = ((int(*)(long long, long long, void*, int, void*, int, int*))cppgo_virt_int_voidp_int_voidp_int)(__goId, -9049994684305682212LL, (void*)parent, index, (void*)data, (int)(action), &__h); if (__h) return __r; }
		return QTreeWidget::dropMimeData(parent, index, data, action);
	}
	QFlags<Qt::DropAction> supportedDropActions() const override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -7784642176186821205LL, &__h); if (__h) return QFlags<Qt::DropAction>(__r); }
		return QTreeWidget::supportedDropActions();
	}
	void dropEvent(QDropEvent* event) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 9107058682754591716LL, (void*)event, &__h); if (__h) return; }
		QTreeWidget::dropEvent(event);
	}
	void setRootIndex(const QModelIndex& index) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1574077111192992763LL, (void*)&index, &__h); if (__h) return; }
		QTreeWidget::setRootIndex(index);
	}
	void keyboardSearch(const QString& search) override {
		int __h = 0;
		QByteArray search_ba = search.toUtf8();
		if (__goOn[8] && cppgo_virt_void_constcharp_longlong) { ((void(*)(long long, long long, const char*, long long, int*))cppgo_virt_void_constcharp_longlong)(__goId, -5041452885022114200LL, (char*)search_ba.constData(), search_ba.size(), &__h); if (__h) return; }
		QTreeWidget::keyboardSearch(search);
	}
	QRect visualRect(const QModelIndex& index) const override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, -1267696851865372997LL, (void*)&index, &__h); if (__h) return *(QRect*)__r; }
		return QTreeWidget::visualRect(index);
	}
	void scrollTo(const QModelIndex& index, QAbstractItemView::ScrollHint hint) override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_void_voidp_int) { ((void(*)(long long, long long, void*, int, int*))cppgo_virt_void_voidp_int)(__goId, -4832340528568269947LL, (void*)&index, (int)(hint), &__h); if (__h) return; }
		QTreeWidget::scrollTo(index, hint);
	}
	QModelIndex indexAt(const QPoint& p) const override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, -7252823847285716802LL, (void*)&p, &__h); if (__h) return *(QModelIndex*)__r; }
		return QTreeWidget::indexAt(p);
	}
	void doItemsLayout() override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, -1888629738259754754LL, &__h); if (__h) return; }
		QTreeWidget::doItemsLayout();
	}
	void reset() override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 1086335023529244512LL, &__h); if (__h) return; }
		QTreeWidget::reset();
	}
	void selectAll() override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 8758113809152574300LL, &__h); if (__h) return; }
		QTreeWidget::selectAll();
	}
	void verticalScrollbarValueChanged(int value) override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 6366953527806415524LL, value, &__h); if (__h) return; }
		QTreeWidget::verticalScrollbarValueChanged(value);
	}
	void scrollContentsBy(int dx, int dy) override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_void_int_int) { ((void(*)(long long, long long, int, int, int*))cppgo_virt_void_int_int)(__goId, -6118001150126564745LL, dx, dy, &__h); if (__h) return; }
		QTreeWidget::scrollContentsBy(dx, dy);
	}
	void rowsInserted(const QModelIndex& parent, int start, int end) override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_void_voidp_int_int) { ((void(*)(long long, long long, void*, int, int, int*))cppgo_virt_void_voidp_int_int)(__goId, 4472434241583495620LL, (void*)&parent, start, end, &__h); if (__h) return; }
		QTreeWidget::rowsInserted(parent, start, end);
	}
	void rowsAboutToBeRemoved(const QModelIndex& parent, int start, int end) override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_void_voidp_int_int) { ((void(*)(long long, long long, void*, int, int, int*))cppgo_virt_void_voidp_int_int)(__goId, -8151049665126058157LL, (void*)&parent, start, end, &__h); if (__h) return; }
		QTreeWidget::rowsAboutToBeRemoved(parent, start, end);
	}
	QModelIndex moveCursor(QAbstractItemView::CursorAction cursorAction, QFlags<Qt::KeyboardModifier> modifiers) override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_voidp_int_int) { void* __r = ((void*(*)(long long, long long, int, int, int*))cppgo_virt_voidp_int_int)(__goId, -158523882194220890LL, (int)(cursorAction), (int)(modifiers), &__h); if (__h) return *(QModelIndex*)__r; }
		return QTreeWidget::moveCursor(cursorAction, modifiers);
	}
	int horizontalOffset() const override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -1500310024801917802LL, &__h); if (__h) return __r; }
		return QTreeWidget::horizontalOffset();
	}
	int verticalOffset() const override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, 5218707810677921016LL, &__h); if (__h) return __r; }
		return QTreeWidget::verticalOffset();
	}
	void setSelection(const QRect& rect, QFlags<QItemSelectionModel::SelectionFlag> command) override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_void_voidp_int) { ((void(*)(long long, long long, void*, int, int*))cppgo_virt_void_voidp_int)(__goId, -1539266838375037305LL, (void*)&rect, (int)(command), &__h); if (__h) return; }
		QTreeWidget::setSelection(rect, command);
	}
	QRegion visualRegionForSelection(const QItemSelection& selection) const override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 7063795720810552074LL, (void*)&selection, &__h); if (__h) return *(QRegion*)__r; }
		return QTreeWidget::visualRegionForSelection(selection);
	}
	QList<QModelIndex> selectedIndexes() const override {
		int __h = 0;
		if (__goOn[24] && cppgo_virtp_voidp) {
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
		return QTreeWidget::selectedIndexes();
	}
	void changeEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4959376063826165943LL, (void*)event, &__h); if (__h) return; }
		QTreeWidget::changeEvent(event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QTreeWidget::timerEvent(event);
	}
	void paintEvent(QPaintEvent* event) override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5290442163167458465LL, (void*)event, &__h); if (__h) return; }
		QTreeWidget::paintEvent(event);
	}
	void drawRow(QPainter* painter, const QStyleOptionViewItem& options, const QModelIndex& index) const override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_void_voidp_voidp_voidp) { ((void(*)(long long, long long, void*, void*, void*, int*))cppgo_virt_void_voidp_voidp_voidp)(__goId, -3950604704165452999LL, (void*)painter, (void*)&options, (void*)&index, &__h); if (__h) return; }
		QTreeWidget::drawRow(painter, options, index);
	}
	void drawBranches(QPainter* painter, const QRect& rect, const QModelIndex& index) const override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_void_voidp_voidp_voidp) { ((void(*)(long long, long long, void*, void*, void*, int*))cppgo_virt_void_voidp_voidp_voidp)(__goId, 6369455307610057755LL, (void*)painter, (void*)&rect, (void*)&index, &__h); if (__h) return; }
		QTreeWidget::drawBranches(painter, rect, index);
	}
	void mousePressEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -332090898078543319LL, (void*)event, &__h); if (__h) return; }
		QTreeWidget::mousePressEvent(event);
	}
	void mouseReleaseEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4415920049297795213LL, (void*)event, &__h); if (__h) return; }
		QTreeWidget::mouseReleaseEvent(event);
	}
	void mouseDoubleClickEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[32] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4029073905697023887LL, (void*)event, &__h); if (__h) return; }
		QTreeWidget::mouseDoubleClickEvent(event);
	}
	void mouseMoveEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[33] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4255494004099806581LL, (void*)event, &__h); if (__h) return; }
		QTreeWidget::mouseMoveEvent(event);
	}
	void keyPressEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[34] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1023791977679793185LL, (void*)event, &__h); if (__h) return; }
		QTreeWidget::keyPressEvent(event);
	}
	void dragMoveEvent(QDragMoveEvent* event) override {
		int __h = 0;
		if (__goOn[35] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1781545137352745978LL, (void*)event, &__h); if (__h) return; }
		QTreeWidget::dragMoveEvent(event);
	}
	bool viewportEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[36] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, -1971465653061861519LL, (void*)event, &__h); if (__h) return __r; }
		return QTreeWidget::viewportEvent(event);
	}
	void updateGeometries() override {
		int __h = 0;
		if (__goOn[37] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 1022771470898460576LL, &__h); if (__h) return; }
		QTreeWidget::updateGeometries();
	}
	QSize viewportSizeHint() const override {
		int __h = 0;
		if (__goOn[38] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 503579067252048239LL, &__h); if (__h) return *(QSize*)__r; }
		return QTreeWidget::viewportSizeHint();
	}
	int sizeHintForColumn(int column) const override {
		int __h = 0;
		if (__goOn[39] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, -3666792239577137638LL, column, &__h); if (__h) return __r; }
		return QTreeWidget::sizeHintForColumn(column);
	}
	void horizontalScrollbarAction(int action) override {
		int __h = 0;
		if (__goOn[40] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, -785551405282638793LL, action, &__h); if (__h) return; }
		QTreeWidget::horizontalScrollbarAction(action);
	}
	bool isIndexHidden(const QModelIndex& index) const override {
		int __h = 0;
		if (__goOn[41] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 6057989667089378859LL, (void*)&index, &__h); if (__h) return __r; }
		return QTreeWidget::isIndexHidden(index);
	}
	void selectionChanged(const QItemSelection& selected, const QItemSelection& deselected) override {
		int __h = 0;
		if (__goOn[42] && cppgo_virt_void_voidp_voidp) { ((void(*)(long long, long long, void*, void*, int*))cppgo_virt_void_voidp_voidp)(__goId, -5892205462145366901LL, (void*)&selected, (void*)&deselected, &__h); if (__h) return; }
		QTreeWidget::selectionChanged(selected, deselected);
	}
	void currentChanged(const QModelIndex& current, const QModelIndex& previous) override {
		int __h = 0;
		if (__goOn[43] && cppgo_virt_void_voidp_voidp) { ((void(*)(long long, long long, void*, void*, int*))cppgo_virt_void_voidp_voidp)(__goId, -6978583972112316034LL, (void*)&current, (void*)&previous, &__h); if (__h) return; }
		QTreeWidget::currentChanged(current, previous);
	}
	int sizeHintForRow(int row) const override {
		int __h = 0;
		if (__goOn[44] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, -4084029023708556876LL, row, &__h); if (__h) return __r; }
		return QTreeWidget::sizeHintForRow(row);
	}
	QAbstractItemDelegate* itemDelegateForIndex(const QModelIndex& index) const override {
		int __h = 0;
		if (__goOn[45] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, -3811438228839363824LL, (void*)&index, &__h); if (__h) return (QAbstractItemDelegate*)__r; }
		return QTreeWidget::itemDelegateForIndex(index);
	}
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const override {
		int __h = 0;
		if (__goOn[46] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7345853097340709638LL, (int)(query), &__h); if (__h) return *(QVariant*)__r; }
		return QTreeWidget::inputMethodQuery(query);
	}
	void updateEditorData() override {
		int __h = 0;
		if (__goOn[47] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 9029941737538200621LL, &__h); if (__h) return; }
		QTreeWidget::updateEditorData();
	}
	void updateEditorGeometries() override {
		int __h = 0;
		if (__goOn[48] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, -8726300139121449691LL, &__h); if (__h) return; }
		QTreeWidget::updateEditorGeometries();
	}
	void verticalScrollbarAction(int action) override {
		int __h = 0;
		if (__goOn[49] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, -1155555122391250623LL, action, &__h); if (__h) return; }
		QTreeWidget::verticalScrollbarAction(action);
	}
	void horizontalScrollbarValueChanged(int value) override {
		int __h = 0;
		if (__goOn[50] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 1786430177580376614LL, value, &__h); if (__h) return; }
		QTreeWidget::horizontalScrollbarValueChanged(value);
	}
	void closeEditor(QWidget* editor, QAbstractItemDelegate::EndEditHint hint) override {
		int __h = 0;
		if (__goOn[51] && cppgo_virt_void_voidp_int) { ((void(*)(long long, long long, void*, int, int*))cppgo_virt_void_voidp_int)(__goId, 8648484264628816806LL, (void*)editor, (int)(hint), &__h); if (__h) return; }
		QTreeWidget::closeEditor(editor, hint);
	}
	void commitData(QWidget* editor) override {
		int __h = 0;
		if (__goOn[52] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6021383853323633948LL, (void*)editor, &__h); if (__h) return; }
		QTreeWidget::commitData(editor);
	}
	void editorDestroyed(QObject* editor) override {
		int __h = 0;
		if (__goOn[53] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6979418566323745651LL, (void*)editor, &__h); if (__h) return; }
		QTreeWidget::editorDestroyed(editor);
	}
	bool edit(const QModelIndex& index, QAbstractItemView::EditTrigger trigger, QEvent* event) override {
		int __h = 0;
		if (__goOn[54] && cppgo_virt_int_voidp_int_voidp) { int __r = ((int(*)(long long, long long, void*, int, void*, int*))cppgo_virt_int_voidp_int_voidp)(__goId, 4206689019978932209LL, (void*)&index, (int)(trigger), (void*)event, &__h); if (__h) return __r; }
		return QTreeWidget::edit(index, trigger, event);
	}
	QFlags<QItemSelectionModel::SelectionFlag> selectionCommand(const QModelIndex& index, const QEvent* event) const override {
		int __h = 0;
		if (__goOn[55] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, -2447490284072398302LL, (void*)&index, (void*)event, &__h); if (__h) return QFlags<QItemSelectionModel::SelectionFlag>(__r); }
		return QTreeWidget::selectionCommand(index, event);
	}
	void startDrag(QFlags<Qt::DropAction> supportedActions) override {
		int __h = 0;
		if (__goOn[56] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, -8112366604639385491LL, (int)(supportedActions), &__h); if (__h) return; }
		QTreeWidget::startDrag(supportedActions);
	}
	void initViewItemOption(QStyleOptionViewItem* option) const override {
		int __h = 0;
		if (__goOn[57] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -7450787277202094556LL, (void*)option, &__h); if (__h) return; }
		QTreeWidget::initViewItemOption(option);
	}
	bool focusNextPrevChild(bool next) override {
		int __h = 0;
		if (__goOn[58] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 5898328517505784191LL, next, &__h); if (__h) return __r; }
		return QTreeWidget::focusNextPrevChild(next);
	}
	void dragEnterEvent(QDragEnterEvent* event) override {
		int __h = 0;
		if (__goOn[59] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 206122715919917589LL, (void*)event, &__h); if (__h) return; }
		QTreeWidget::dragEnterEvent(event);
	}
	void dragLeaveEvent(QDragLeaveEvent* event) override {
		int __h = 0;
		if (__goOn[60] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -3139904158373725006LL, (void*)event, &__h); if (__h) return; }
		QTreeWidget::dragLeaveEvent(event);
	}
	void focusInEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[61] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 975043767033552552LL, (void*)event, &__h); if (__h) return; }
		QTreeWidget::focusInEvent(event);
	}
	void focusOutEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[62] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -385596010576048993LL, (void*)event, &__h); if (__h) return; }
		QTreeWidget::focusOutEvent(event);
	}
	void resizeEvent(QResizeEvent* event) override {
		int __h = 0;
		if (__goOn[63] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4302917285655014661LL, (void*)event, &__h); if (__h) return; }
		QTreeWidget::resizeEvent(event);
	}
	void inputMethodEvent(QInputMethodEvent* event) override {
		int __h = 0;
		if (__goOn[64] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5438819661942843222LL, (void*)event, &__h); if (__h) return; }
		QTreeWidget::inputMethodEvent(event);
	}
	bool eventFilter(QObject* object, QEvent* event) override {
		int __h = 0;
		if (__goOn[65] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)object, (void*)event, &__h); if (__h) return __r; }
		return QTreeWidget::eventFilter(object, event);
	}
	QSize minimumSizeHint() const override {
		int __h = 0;
		if (__goOn[66] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 7049264599837532491LL, &__h); if (__h) return *(QSize*)__r; }
		return QTreeWidget::minimumSizeHint();
	}
	QSize sizeHint() const override {
		int __h = 0;
		if (__goOn[67] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 729690261457684101LL, &__h); if (__h) return *(QSize*)__r; }
		return QTreeWidget::sizeHint();
	}
	void setupViewport(QWidget* viewport) override {
		int __h = 0;
		if (__goOn[68] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4531996381736730688LL, (void*)viewport, &__h); if (__h) return; }
		QTreeWidget::setupViewport(viewport);
	}
	void wheelEvent(QWheelEvent* arg1) override {
		int __h = 0;
		if (__goOn[69] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1919988340934511316LL, (void*)arg1, &__h); if (__h) return; }
		QTreeWidget::wheelEvent(arg1);
	}
	void contextMenuEvent(QContextMenuEvent* arg1) override {
		int __h = 0;
		if (__goOn[70] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6698043222227006015LL, (void*)arg1, &__h); if (__h) return; }
		QTreeWidget::contextMenuEvent(arg1);
	}
	void initStyleOption(QStyleOptionFrame* option) const override {
		int __h = 0;
		if (__goOn[71] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -488693769771308531LL, (void*)option, &__h); if (__h) return; }
		QTreeWidget::initStyleOption(option);
	}
	int devType() const override {
		int __h = 0;
		if (__goOn[72] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8659679893868524044LL, &__h); if (__h) return __r; }
		return QTreeWidget::devType();
	}
	void setVisible(bool visible) override {
		int __h = 0;
		if (__goOn[73] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 2110182719375245575LL, visible, &__h); if (__h) return; }
		QTreeWidget::setVisible(visible);
	}
	int heightForWidth(int arg1) const override {
		int __h = 0;
		if (__goOn[74] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 4311534060495737339LL, arg1, &__h); if (__h) return __r; }
		return QTreeWidget::heightForWidth(arg1);
	}
	bool hasHeightForWidth() const override {
		int __h = 0;
		if (__goOn[75] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8266762397995423915LL, &__h); if (__h) return __r; }
		return QTreeWidget::hasHeightForWidth();
	}
	QPaintEngine* paintEngine() const override {
		int __h = 0;
		if (__goOn[76] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 6398007404756346381LL, &__h); if (__h) return (QPaintEngine*)__r; }
		return QTreeWidget::paintEngine();
	}
	void keyReleaseEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[77] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 772598475950050299LL, (void*)event, &__h); if (__h) return; }
		QTreeWidget::keyReleaseEvent(event);
	}
	void enterEvent(QEnterEvent* event) override {
		int __h = 0;
		if (__goOn[78] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6316270346668237895LL, (void*)event, &__h); if (__h) return; }
		QTreeWidget::enterEvent(event);
	}
	void leaveEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[79] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -551070009623711640LL, (void*)event, &__h); if (__h) return; }
		QTreeWidget::leaveEvent(event);
	}
	void moveEvent(QMoveEvent* event) override {
		int __h = 0;
		if (__goOn[80] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -9161949150489525396LL, (void*)event, &__h); if (__h) return; }
		QTreeWidget::moveEvent(event);
	}
	void closeEvent(QCloseEvent* event) override {
		int __h = 0;
		if (__goOn[81] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1989937050019464783LL, (void*)event, &__h); if (__h) return; }
		QTreeWidget::closeEvent(event);
	}
	void tabletEvent(QTabletEvent* event) override {
		int __h = 0;
		if (__goOn[82] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -212068866860701141LL, (void*)event, &__h); if (__h) return; }
		QTreeWidget::tabletEvent(event);
	}
	void actionEvent(QActionEvent* event) override {
		int __h = 0;
		if (__goOn[83] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1697444140088555581LL, (void*)event, &__h); if (__h) return; }
		QTreeWidget::actionEvent(event);
	}
	void showEvent(QShowEvent* event) override {
		int __h = 0;
		if (__goOn[84] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6825369174751916LL, (void*)event, &__h); if (__h) return; }
		QTreeWidget::showEvent(event);
	}
	void hideEvent(QHideEvent* event) override {
		int __h = 0;
		if (__goOn[85] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -199718707876719129LL, (void*)event, &__h); if (__h) return; }
		QTreeWidget::hideEvent(event);
	}
	bool nativeEvent(const QByteArray& eventType, void* message, qintptr* result) override {
		int __h = 0;
		if (__goOn[86] && cppgo_virt_int_voidp_voidp_longlongp) { int __r = ((int(*)(long long, long long, void*, void*, long long*, int*))cppgo_virt_int_voidp_voidp_longlongp)(__goId, -2039166586246443706LL, (void*)&eventType, (void*)message, (long long*)result, &__h); if (__h) return __r; }
		return QTreeWidget::nativeEvent(eventType, message, result);
	}
	int metric(QPaintDevice::PaintDeviceMetric arg1) const override {
		int __h = 0;
		if (__goOn[87] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 427629649267311777LL, (int)(arg1), &__h); if (__h) return __r; }
		return QTreeWidget::metric(arg1);
	}
	void initPainter(QPainter* painter) const override {
		int __h = 0;
		if (__goOn[88] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6477500296980050210LL, (void*)painter, &__h); if (__h) return; }
		QTreeWidget::initPainter(painter);
	}
	QPaintDevice* redirected(QPoint* offset) const override {
		int __h = 0;
		if (__goOn[89] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 4127240633310655048LL, (void*)offset, &__h); if (__h) return (QPaintDevice*)__r; }
		return QTreeWidget::redirected(offset);
	}
	QPainter* sharedPainter() const override {
		int __h = 0;
		if (__goOn[90] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 5792771328741774991LL, &__h); if (__h) return (QPainter*)__r; }
		return QTreeWidget::sharedPainter();
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[91] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QTreeWidget::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[92] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QTreeWidget::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[93] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QTreeWidget::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[94] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QTreeWidget::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QTreeWidget;
static std::mutex __connsMu_QTreeWidget;

extern "C" {

void* QTreeWidget_base(void* self) { return (void*)static_cast<QTreeView*>((QTreeWidget*)self); }
void QTreeWidget_chain(void* self, void** out) {
	QTreeWidget* __p = (QTreeWidget*)self;
	QTreeView* __b0 = static_cast<QTreeView*>(__p); out[0] = (void*)__b0;
	QAbstractItemView* __b1 = static_cast<QAbstractItemView*>(__b0); out[1] = (void*)__b1;
	QAbstractScrollArea* __b2 = static_cast<QAbstractScrollArea*>(__b1); out[2] = (void*)__b2;
	QFrame* __b3 = static_cast<QFrame*>(__b2); out[3] = (void*)__b3;
	QWidget* __b4 = static_cast<QWidget*>(__b3); out[4] = (void*)__b4;
	QObject* __b5 = static_cast<QObject*>(__b4); out[5] = (void*)__b5;
}
void QTreeWidget__set_override(void* self, int idx) { ((QTreeWidget__Go*)self)->__goOn[idx] = 1; }
void* QTreeWidget_fromQObject(void* o) { return (void*)qobject_cast<QTreeWidget*>((QObject*)o); }
void* QTreeWidget_new(long long __goid, void* parent) {
	return new QTreeWidget__Go(__goid, (QWidget*)parent);
}
void* QTreeWidget_new2(long long __goid) {
	return new QTreeWidget__Go(__goid);
}
char* QTreeWidget_tr(const char* s, long long* __out_len) {
	auto&& __r = QTreeWidget::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QTreeWidget_columnCount(void* self) {
	auto&& __r = ((QTreeWidget*)self)->columnCount();
	return __r;
}
void QTreeWidget_setColumnCount(void* self, int columns) {
	((QTreeWidget*)self)->setColumnCount(columns);
}
void* QTreeWidget_invisibleRootItem(void* self) {
	auto&& __r = ((QTreeWidget*)self)->invisibleRootItem();
	return (void*)(__r);
}
void* QTreeWidget_topLevelItem(void* self, int index) {
	auto&& __r = ((QTreeWidget*)self)->topLevelItem(index);
	return (void*)(__r);
}
int QTreeWidget_topLevelItemCount(void* self) {
	auto&& __r = ((QTreeWidget*)self)->topLevelItemCount();
	return __r;
}
void QTreeWidget_insertTopLevelItem(void* self, int index, void* item) {
	((QTreeWidget*)self)->insertTopLevelItem(index, (QTreeWidgetItem*)item);
}
void QTreeWidget_addTopLevelItem(void* self, void* item) {
	((QTreeWidget*)self)->addTopLevelItem((QTreeWidgetItem*)item);
}
void* QTreeWidget_takeTopLevelItem(void* self, int index) {
	auto&& __r = ((QTreeWidget*)self)->takeTopLevelItem(index);
	return (void*)(__r);
}
int QTreeWidget_indexOfTopLevelItem(void* self, void* item) {
	auto&& __r = ((QTreeWidget*)self)->indexOfTopLevelItem((QTreeWidgetItem*)item);
	return __r;
}
void QTreeWidget_insertTopLevelItems(void* self, int index, void** items, long long items_len) {
	QList<QTreeWidgetItem*> items_list; items_list.reserve(items_len);
	for (long long __i = 0; __i < items_len; __i++) items_list.append((QTreeWidgetItem*)items[__i]);
	((QTreeWidget*)self)->insertTopLevelItems(index, items_list);
}
void QTreeWidget_addTopLevelItems(void* self, void** items, long long items_len) {
	QList<QTreeWidgetItem*> items_list; items_list.reserve(items_len);
	for (long long __i = 0; __i < items_len; __i++) items_list.append((QTreeWidgetItem*)items[__i]);
	((QTreeWidget*)self)->addTopLevelItems(items_list);
}
void* QTreeWidget_headerItem(void* self) {
	auto&& __r = ((QTreeWidget*)self)->headerItem();
	return (void*)(__r);
}
void QTreeWidget_setHeaderItem(void* self, void* item) {
	((QTreeWidget*)self)->setHeaderItem((QTreeWidgetItem*)item);
}
void QTreeWidget_setHeaderLabels(void* self, char** labels, long long labels_len) {
	QList<QString> labels_list; labels_list.reserve(labels_len);
	for (long long __i = 0; __i < labels_len; __i++) labels_list.append(QString::fromUtf8(labels[__i]));
	((QTreeWidget*)self)->setHeaderLabels(labels_list);
}
void QTreeWidget_setHeaderLabel(void* self, const char* label, long long label_len) {
	QString label_qstr = QString::fromUtf8(label, label_len);
	((QTreeWidget*)self)->setHeaderLabel(label_qstr);
}
void* QTreeWidget_currentItem(void* self) {
	auto&& __r = ((QTreeWidget*)self)->currentItem();
	return (void*)(__r);
}
int QTreeWidget_currentColumn(void* self) {
	auto&& __r = ((QTreeWidget*)self)->currentColumn();
	return __r;
}
void QTreeWidget_setCurrentItem(void* self, void* item) {
	((QTreeWidget*)self)->setCurrentItem((QTreeWidgetItem*)item);
}
void QTreeWidget_setCurrentItem2(void* self, void* item, int column) {
	((QTreeWidget*)self)->setCurrentItem((QTreeWidgetItem*)item, column);
}
void QTreeWidget_setCurrentItem3(void* self, void* item, int column, int command) {
	((QTreeWidget*)self)->setCurrentItem((QTreeWidgetItem*)item, column, QFlags<QItemSelectionModel::SelectionFlag>(command));
}
void* QTreeWidget_itemAt(void* self, void* p) {
	auto&& __r = ((QTreeWidget*)self)->itemAt(*(QPoint*)p);
	return (void*)(__r);
}
void* QTreeWidget_itemAt2(void* self, int x, int y) {
	auto&& __r = ((QTreeWidget*)self)->itemAt(x, y);
	return (void*)(__r);
}
void* QTreeWidget_visualItemRect(void* self, void* item) {
	auto&& __r = ((QTreeWidget*)self)->visualItemRect((QTreeWidgetItem*)item);
	return new QRect(__r);
}
int QTreeWidget_sortColumn(void* self) {
	auto&& __r = ((QTreeWidget*)self)->sortColumn();
	return __r;
}
void QTreeWidget_sortItems(void* self, int column, int order) {
	((QTreeWidget*)self)->sortItems(column, (Qt::SortOrder)(order));
}
void QTreeWidget_editItem(void* self, void* item) {
	((QTreeWidget*)self)->editItem((QTreeWidgetItem*)item);
}
void QTreeWidget_openPersistentEditor(void* self, void* item) {
	((QTreeWidget*)self)->openPersistentEditor((QTreeWidgetItem*)item);
}
void QTreeWidget_closePersistentEditor(void* self, void* item) {
	((QTreeWidget*)self)->closePersistentEditor((QTreeWidgetItem*)item);
}
int QTreeWidget_isPersistentEditorOpen(void* self, void* item) {
	auto&& __r = ((QTreeWidget*)self)->isPersistentEditorOpen((QTreeWidgetItem*)item);
	return __r;
}
void* QTreeWidget_itemWidget(void* self, void* item, int column) {
	auto&& __r = ((QTreeWidget*)self)->itemWidget((QTreeWidgetItem*)item, column);
	return (void*)(__r);
}
void QTreeWidget_setItemWidget(void* self, void* item, int column, void* widget) {
	((QTreeWidget*)self)->setItemWidget((QTreeWidgetItem*)item, column, (QWidget*)widget);
}
void QTreeWidget_removeItemWidget(void* self, void* item, int column) {
	((QTreeWidget*)self)->removeItemWidget((QTreeWidgetItem*)item, column);
}
void** QTreeWidget_selectedItems(void* self, long long* __out_len) {
	auto&& __r = ((QTreeWidget*)self)->selectedItems();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
void** QTreeWidget_findItems(void* self, const char* text, long long text_len, int flags, long long* __out_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = ((QTreeWidget*)self)->findItems(text_qstr, QFlags<Qt::MatchFlag>(flags));
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
void* QTreeWidget_itemAbove(void* self, void* item) {
	auto&& __r = ((QTreeWidget*)self)->itemAbove((QTreeWidgetItem*)item);
	return (void*)(__r);
}
void* QTreeWidget_itemBelow(void* self, void* item) {
	auto&& __r = ((QTreeWidget*)self)->itemBelow((QTreeWidgetItem*)item);
	return (void*)(__r);
}
void* QTreeWidget_indexFromItem(void* self, void* item) {
	auto&& __r = ((QTreeWidget*)self)->indexFromItem((QTreeWidgetItem*)item);
	return new QModelIndex(__r);
}
void* QTreeWidget_itemFromIndex(void* self, void* index) {
	auto&& __r = ((QTreeWidget*)self)->itemFromIndex(*(QModelIndex*)index);
	return (void*)(__r);
}
int QTreeWidget_supportedDragActions(void* self) {
	auto&& __r = ((QTreeWidget*)self)->supportedDragActions();
	return (int)(__r);
}
void QTreeWidget_setSupportedDragActions(void* self, int actions) {
	((QTreeWidget*)self)->setSupportedDragActions(QFlags<Qt::DropAction>(actions));
}
void QTreeWidget_scrollToItem(void* self, void* item) {
	((QTreeWidget*)self)->scrollToItem((QTreeWidgetItem*)item);
}
void QTreeWidget_expandItem(void* self, void* item) {
	((QTreeWidget*)self)->expandItem((QTreeWidgetItem*)item);
}
void QTreeWidget_collapseItem(void* self, void* item) {
	((QTreeWidget*)self)->collapseItem((QTreeWidgetItem*)item);
}
void QTreeWidget_clear(void* self) {
	((QTreeWidget*)self)->clear();
}
void QTreeWidget_itemPressed(void* self, void* item, int column) {
	((QTreeWidget*)self)->itemPressed((QTreeWidgetItem*)item, column);
}
void QTreeWidget_itemClicked(void* self, void* item, int column) {
	((QTreeWidget*)self)->itemClicked((QTreeWidgetItem*)item, column);
}
void QTreeWidget_itemDoubleClicked(void* self, void* item, int column) {
	((QTreeWidget*)self)->itemDoubleClicked((QTreeWidgetItem*)item, column);
}
void QTreeWidget_itemActivated(void* self, void* item, int column) {
	((QTreeWidget*)self)->itemActivated((QTreeWidgetItem*)item, column);
}
void QTreeWidget_itemEntered(void* self, void* item, int column) {
	((QTreeWidget*)self)->itemEntered((QTreeWidgetItem*)item, column);
}
void QTreeWidget_itemChanged(void* self, void* item, int column) {
	((QTreeWidget*)self)->itemChanged((QTreeWidgetItem*)item, column);
}
void QTreeWidget_itemExpanded(void* self, void* item) {
	((QTreeWidget*)self)->itemExpanded((QTreeWidgetItem*)item);
}
void QTreeWidget_itemCollapsed(void* self, void* item) {
	((QTreeWidget*)self)->itemCollapsed((QTreeWidgetItem*)item);
}
void QTreeWidget_currentItemChanged(void* self, void* current, void* previous) {
	((QTreeWidget*)self)->currentItemChanged((QTreeWidgetItem*)current, (QTreeWidgetItem*)previous);
}
void QTreeWidget_itemSelectionChanged(void* self) {
	((QTreeWidget*)self)->itemSelectionChanged();
}
void QTreeWidget_columnResized(void* self, int column, int oldSize, int newSize) {
	((QTreeWidget__Go*)self)->columnResized(column, oldSize, newSize);
}
void QTreeWidget_columnCountChanged(void* self, int oldCount, int newCount) {
	((QTreeWidget__Go*)self)->columnCountChanged(oldCount, newCount);
}
void QTreeWidget_columnMoved(void* self) {
	((QTreeWidget__Go*)self)->columnMoved();
}
void QTreeWidget_reexpand(void* self) {
	((QTreeWidget__Go*)self)->reexpand();
}
void QTreeWidget_rowsRemoved(void* self, void* parent, int first, int last) {
	((QTreeWidget__Go*)self)->rowsRemoved(*(QModelIndex*)parent, first, last);
}
void QTreeWidget_drawTree(void* self, void* painter, void* region) {
	((QTreeWidget__Go*)self)->drawTree((QPainter*)painter, *(QRegion*)region);
}
int QTreeWidget_indexRowSizeHint(void* self, void* index) {
	auto&& __r = ((QTreeWidget__Go*)self)->indexRowSizeHint(*(QModelIndex*)index);
	return __r;
}
int QTreeWidget_rowHeight(void* self, void* index) {
	auto&& __r = ((QTreeWidget__Go*)self)->rowHeight(*(QModelIndex*)index);
	return __r;
}
int QTreeWidget_state(void* self) {
	auto&& __r = ((QTreeWidget__Go*)self)->state();
	return (int)(__r);
}
void QTreeWidget_scheduleDelayedItemsLayout(void* self) {
	((QTreeWidget__Go*)self)->scheduleDelayedItemsLayout();
}
void QTreeWidget_executeDelayedItemsLayout(void* self) {
	((QTreeWidget__Go*)self)->executeDelayedItemsLayout();
}
void QTreeWidget_setDirtyRegion(void* self, void* region) {
	((QTreeWidget__Go*)self)->setDirtyRegion(*(QRegion*)region);
}
void QTreeWidget_scrollDirtyRegion(void* self, int dx, int dy) {
	((QTreeWidget__Go*)self)->scrollDirtyRegion(dx, dy);
}
void* QTreeWidget_dirtyRegionOffset(void* self) {
	auto&& __r = ((QTreeWidget__Go*)self)->dirtyRegionOffset();
	return new QPoint(__r);
}
void QTreeWidget_startAutoScroll(void* self) {
	((QTreeWidget__Go*)self)->startAutoScroll();
}
void QTreeWidget_stopAutoScroll(void* self) {
	((QTreeWidget__Go*)self)->stopAutoScroll();
}
void QTreeWidget_doAutoScroll(void* self) {
	((QTreeWidget__Go*)self)->doAutoScroll();
}
int QTreeWidget_dropIndicatorPosition(void* self) {
	auto&& __r = ((QTreeWidget__Go*)self)->dropIndicatorPosition();
	return (int)(__r);
}
void QTreeWidget_setViewportMargins(void* self, int left, int top, int right, int bottom) {
	((QTreeWidget__Go*)self)->setViewportMargins(left, top, right, bottom);
}
void QTreeWidget_setViewportMargins2(void* self, void* margins) {
	((QTreeWidget__Go*)self)->setViewportMargins(*(QMargins*)margins);
}
void* QTreeWidget_viewportMargins(void* self) {
	auto&& __r = ((QTreeWidget__Go*)self)->viewportMargins();
	return new QMargins(__r);
}
void QTreeWidget_drawFrame(void* self, void* arg1) {
	((QTreeWidget__Go*)self)->drawFrame((QPainter*)arg1);
}
void QTreeWidget_setTabOrder(void* arg1, void* arg2) {
	QTreeWidget::setTabOrder((QWidget*)arg1, (QWidget*)arg2);
}
void* QTreeWidget_mouseGrabber() {
	auto&& __r = QTreeWidget::mouseGrabber();
	return (void*)(__r);
}
void* QTreeWidget_keyboardGrabber() {
	auto&& __r = QTreeWidget::keyboardGrabber();
	return (void*)(__r);
}
void* QTreeWidget_find(unsigned long long arg1) {
	auto&& __r = QTreeWidget::find(arg1);
	return (void*)(__r);
}
void* QTreeWidget_createWindowContainer(void* window) {
	auto&& __r = QTreeWidget::createWindowContainer((QWindow*)window);
	return (void*)(__r);
}
void QTreeWidget_updateMicroFocus(void* self) {
	((QTreeWidget__Go*)self)->updateMicroFocus();
}
void QTreeWidget_create(void* self) {
	((QTreeWidget__Go*)self)->create();
}
void QTreeWidget_destroy(void* self) {
	((QTreeWidget__Go*)self)->destroy();
}
int QTreeWidget_focusNextChild(void* self) {
	auto&& __r = ((QTreeWidget__Go*)self)->focusNextChild();
	return __r;
}
int QTreeWidget_focusPreviousChild(void* self) {
	auto&& __r = ((QTreeWidget__Go*)self)->focusPreviousChild();
	return __r;
}
void* QTreeWidget_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QTreeWidget::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QTreeWidget_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QTreeWidget::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QTreeWidget_sender(void* self) {
	auto&& __r = ((QTreeWidget__Go*)self)->sender();
	return (void*)(__r);
}
int QTreeWidget_senderSignalIndex(void* self) {
	auto&& __r = ((QTreeWidget__Go*)self)->senderSignalIndex();
	return __r;
}
int QTreeWidget_receivers(void* self, const char* signal) {
	auto&& __r = ((QTreeWidget__Go*)self)->receivers(signal);
	return __r;
}
int QTreeWidget_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QTreeWidget__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
double QTreeWidget_devicePixelRatioFScale() {
	auto&& __r = QTreeWidget::devicePixelRatioFScale();
	return __r;
}
int QTreeWidget_encodeMetricF(int metric, double value) {
	auto&& __r = QTreeWidget::encodeMetricF((QPaintDevice::PaintDeviceMetric)(metric), value);
	return __r;
}
double QTreeWidget_getDecodedMetricF(void* self, int metricA, int metricB) {
	auto&& __r = ((QTreeWidget__Go*)self)->getDecodedMetricF((QPaintDevice::PaintDeviceMetric)(metricA), (QPaintDevice::PaintDeviceMetric)(metricB));
	return __r;
}
char* QTreeWidget_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QTreeWidget::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QTreeWidget_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QTreeWidget::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QTreeWidget_editItem2(void* self, void* item, int column) {
	((QTreeWidget*)self)->editItem((QTreeWidgetItem*)item, column);
}
void QTreeWidget_openPersistentEditor2(void* self, void* item, int column) {
	((QTreeWidget*)self)->openPersistentEditor((QTreeWidgetItem*)item, column);
}
void QTreeWidget_closePersistentEditor2(void* self, void* item, int column) {
	((QTreeWidget*)self)->closePersistentEditor((QTreeWidgetItem*)item, column);
}
int QTreeWidget_isPersistentEditorOpen2(void* self, void* item, int column) {
	auto&& __r = ((QTreeWidget*)self)->isPersistentEditorOpen((QTreeWidgetItem*)item, column);
	return __r;
}
void** QTreeWidget_findItems2(void* self, const char* text, long long text_len, int flags, int column, long long* __out_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = ((QTreeWidget*)self)->findItems(text_qstr, QFlags<Qt::MatchFlag>(flags), column);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
void* QTreeWidget_indexFromItem2(void* self, void* item, int column) {
	auto&& __r = ((QTreeWidget*)self)->indexFromItem((QTreeWidgetItem*)item, column);
	return new QModelIndex(__r);
}
void QTreeWidget_scrollToItem2(void* self, void* item, int hint) {
	((QTreeWidget*)self)->scrollToItem((QTreeWidgetItem*)item, (QAbstractItemView::ScrollHint)(hint));
}
void* QTreeWidget_createWindowContainer2(void* window, void* parent) {
	auto&& __r = QTreeWidget::createWindowContainer((QWindow*)window, (QWidget*)parent);
	return (void*)(__r);
}
void* QTreeWidget_createWindowContainer3(void* window, void* parent, int flags) {
	auto&& __r = QTreeWidget::createWindowContainer((QWindow*)window, (QWidget*)parent, QFlags<Qt::WindowType>(flags));
	return (void*)(__r);
}
void QTreeWidget_updateMicroFocus2(void* self, int query) {
	((QTreeWidget__Go*)self)->updateMicroFocus((Qt::InputMethodQuery)(query));
}
void QTreeWidget_create2(void* self, unsigned long long arg1) {
	((QTreeWidget__Go*)self)->create(arg1);
}
void QTreeWidget_create3(void* self, unsigned long long arg1, int initializeWindow) {
	((QTreeWidget__Go*)self)->create(arg1, initializeWindow);
}
void QTreeWidget_create4(void* self, unsigned long long arg1, int initializeWindow, int destroyOldWindow) {
	((QTreeWidget__Go*)self)->create(arg1, initializeWindow, destroyOldWindow);
}
void QTreeWidget_destroy2(void* self, int destroyWindow) {
	((QTreeWidget__Go*)self)->destroy(destroyWindow);
}
void QTreeWidget_destroy3(void* self, int destroyWindow, int destroySubWindows) {
	((QTreeWidget__Go*)self)->destroy(destroyWindow, destroySubWindows);
}
void* QTreeWidget_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QTreeWidget::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QTreeWidget_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QTreeWidget::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QTreeWidget_metaObject_default(void* self) {
	auto&& __r = ((QTreeWidget*)self)->QTreeWidget::metaObject();
	return (void*)(__r);
}
void QTreeWidget_setSelectionModel_default(void* self, void* selectionModel) {
	((QTreeWidget*)self)->QTreeWidget::setSelectionModel((QItemSelectionModel*)selectionModel);
}
int QTreeWidget_event_default(void* self, void* e) {
	auto&& __r = ((QTreeWidget__Go*)self)->cppgoBase_event((QEvent*)e);
	return __r;
}
char** QTreeWidget_mimeTypes_default(void* self, long long* __out_len) {
	auto&& __r = ((QTreeWidget__Go*)self)->cppgoBase_mimeTypes();
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
int QTreeWidget_dropMimeData_default(void* self, void* parent, int index, void* data, int action) {
	auto&& __r = ((QTreeWidget__Go*)self)->cppgoBase_dropMimeData((QTreeWidgetItem*)parent, index, (QMimeData*)data, action);
	return __r;
}
int QTreeWidget_supportedDropActions_default(void* self) {
	auto&& __r = ((QTreeWidget__Go*)self)->cppgoBase_supportedDropActions();
	return (int)(__r);
}
void QTreeWidget_dropEvent_default(void* self, void* event) {
	((QTreeWidget__Go*)self)->cppgoBase_dropEvent((QDropEvent*)event);
}
void QTreeWidget_setRootIndex_default(void* self, void* index) {
	((QTreeWidget*)self)->QTreeWidget::setRootIndex(*(QModelIndex*)index);
}
void QTreeWidget_keyboardSearch_default(void* self, const char* search, long long search_len) {
	QString search_qstr = QString::fromUtf8(search, search_len);
	((QTreeWidget*)self)->QTreeWidget::keyboardSearch(search_qstr);
}
void* QTreeWidget_visualRect_default(void* self, void* index) {
	auto&& __r = ((QTreeWidget*)self)->QTreeWidget::visualRect(*(QModelIndex*)index);
	return new QRect(__r);
}
void QTreeWidget_scrollTo_default(void* self, void* index, int hint) {
	((QTreeWidget*)self)->QTreeWidget::scrollTo(*(QModelIndex*)index, (QAbstractItemView::ScrollHint)(hint));
}
void* QTreeWidget_indexAt_default(void* self, void* p) {
	auto&& __r = ((QTreeWidget*)self)->QTreeWidget::indexAt(*(QPoint*)p);
	return new QModelIndex(__r);
}
void QTreeWidget_doItemsLayout_default(void* self) {
	((QTreeWidget*)self)->QTreeWidget::doItemsLayout();
}
void QTreeWidget_reset_default(void* self) {
	((QTreeWidget*)self)->QTreeWidget::reset();
}
void QTreeWidget_selectAll_default(void* self) {
	((QTreeWidget*)self)->QTreeWidget::selectAll();
}
void QTreeWidget_verticalScrollbarValueChanged_default(void* self, int value) {
	((QTreeWidget__Go*)self)->cppgoBase_verticalScrollbarValueChanged(value);
}
void QTreeWidget_scrollContentsBy_default(void* self, int dx, int dy) {
	((QTreeWidget__Go*)self)->cppgoBase_scrollContentsBy(dx, dy);
}
void QTreeWidget_rowsInserted_default(void* self, void* parent, int start, int end) {
	((QTreeWidget__Go*)self)->cppgoBase_rowsInserted(*(QModelIndex*)parent, start, end);
}
void QTreeWidget_rowsAboutToBeRemoved_default(void* self, void* parent, int start, int end) {
	((QTreeWidget__Go*)self)->cppgoBase_rowsAboutToBeRemoved(*(QModelIndex*)parent, start, end);
}
void* QTreeWidget_moveCursor_default(void* self, int cursorAction, int modifiers) {
	auto&& __r = ((QTreeWidget__Go*)self)->cppgoBase_moveCursor(cursorAction, QFlags<Qt::KeyboardModifier>(modifiers));
	return new QModelIndex(__r);
}
int QTreeWidget_horizontalOffset_default(void* self) {
	auto&& __r = ((QTreeWidget__Go*)self)->cppgoBase_horizontalOffset();
	return __r;
}
int QTreeWidget_verticalOffset_default(void* self) {
	auto&& __r = ((QTreeWidget__Go*)self)->cppgoBase_verticalOffset();
	return __r;
}
void QTreeWidget_setSelection_default(void* self, void* rect, int command) {
	((QTreeWidget__Go*)self)->cppgoBase_setSelection(*(QRect*)rect, QFlags<QItemSelectionModel::SelectionFlag>(command));
}
void* QTreeWidget_visualRegionForSelection_default(void* self, void* selection) {
	auto&& __r = ((QTreeWidget__Go*)self)->cppgoBase_visualRegionForSelection(*(QItemSelection*)selection);
	return new QRegion(__r);
}
void** QTreeWidget_selectedIndexes_default(void* self, long long* __out_len) {
	auto&& __r = ((QTreeWidget__Go*)self)->cppgoBase_selectedIndexes();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QModelIndex(__e);
	return __arr;
}
void QTreeWidget_changeEvent_default(void* self, void* event) {
	((QTreeWidget__Go*)self)->cppgoBase_changeEvent((QEvent*)event);
}
void QTreeWidget_timerEvent_default(void* self, void* event) {
	((QTreeWidget__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QTreeWidget_paintEvent_default(void* self, void* event) {
	((QTreeWidget__Go*)self)->cppgoBase_paintEvent((QPaintEvent*)event);
}
void QTreeWidget_drawRow_default(void* self, void* painter, void* options, void* index) {
	((QTreeWidget__Go*)self)->cppgoBase_drawRow((QPainter*)painter, *(QStyleOptionViewItem*)options, *(QModelIndex*)index);
}
void QTreeWidget_drawBranches_default(void* self, void* painter, void* rect, void* index) {
	((QTreeWidget__Go*)self)->cppgoBase_drawBranches((QPainter*)painter, *(QRect*)rect, *(QModelIndex*)index);
}
void QTreeWidget_mousePressEvent_default(void* self, void* event) {
	((QTreeWidget__Go*)self)->cppgoBase_mousePressEvent((QMouseEvent*)event);
}
void QTreeWidget_mouseReleaseEvent_default(void* self, void* event) {
	((QTreeWidget__Go*)self)->cppgoBase_mouseReleaseEvent((QMouseEvent*)event);
}
void QTreeWidget_mouseDoubleClickEvent_default(void* self, void* event) {
	((QTreeWidget__Go*)self)->cppgoBase_mouseDoubleClickEvent((QMouseEvent*)event);
}
void QTreeWidget_mouseMoveEvent_default(void* self, void* event) {
	((QTreeWidget__Go*)self)->cppgoBase_mouseMoveEvent((QMouseEvent*)event);
}
void QTreeWidget_keyPressEvent_default(void* self, void* event) {
	((QTreeWidget__Go*)self)->cppgoBase_keyPressEvent((QKeyEvent*)event);
}
void QTreeWidget_dragMoveEvent_default(void* self, void* event) {
	((QTreeWidget__Go*)self)->cppgoBase_dragMoveEvent((QDragMoveEvent*)event);
}
int QTreeWidget_viewportEvent_default(void* self, void* event) {
	auto&& __r = ((QTreeWidget__Go*)self)->cppgoBase_viewportEvent((QEvent*)event);
	return __r;
}
void QTreeWidget_updateGeometries_default(void* self) {
	((QTreeWidget__Go*)self)->cppgoBase_updateGeometries();
}
void* QTreeWidget_viewportSizeHint_default(void* self) {
	auto&& __r = ((QTreeWidget__Go*)self)->cppgoBase_viewportSizeHint();
	return new QSize(__r);
}
int QTreeWidget_sizeHintForColumn_default(void* self, int column) {
	auto&& __r = ((QTreeWidget__Go*)self)->cppgoBase_sizeHintForColumn(column);
	return __r;
}
void QTreeWidget_horizontalScrollbarAction_default(void* self, int action) {
	((QTreeWidget__Go*)self)->cppgoBase_horizontalScrollbarAction(action);
}
int QTreeWidget_isIndexHidden_default(void* self, void* index) {
	auto&& __r = ((QTreeWidget__Go*)self)->cppgoBase_isIndexHidden(*(QModelIndex*)index);
	return __r;
}
void QTreeWidget_selectionChanged_default(void* self, void* selected, void* deselected) {
	((QTreeWidget__Go*)self)->cppgoBase_selectionChanged(*(QItemSelection*)selected, *(QItemSelection*)deselected);
}
void QTreeWidget_currentChanged_default(void* self, void* current, void* previous) {
	((QTreeWidget__Go*)self)->cppgoBase_currentChanged(*(QModelIndex*)current, *(QModelIndex*)previous);
}
int QTreeWidget_sizeHintForRow_default(void* self, int row) {
	auto&& __r = ((QTreeWidget*)self)->QTreeWidget::sizeHintForRow(row);
	return __r;
}
void* QTreeWidget_itemDelegateForIndex_default(void* self, void* index) {
	auto&& __r = ((QTreeWidget*)self)->QTreeWidget::itemDelegateForIndex(*(QModelIndex*)index);
	return (void*)(__r);
}
void* QTreeWidget_inputMethodQuery_default(void* self, int query) {
	auto&& __r = ((QTreeWidget*)self)->QTreeWidget::inputMethodQuery((Qt::InputMethodQuery)(query));
	return new QVariant(__r);
}
void QTreeWidget_updateEditorData_default(void* self) {
	((QTreeWidget__Go*)self)->cppgoBase_updateEditorData();
}
void QTreeWidget_updateEditorGeometries_default(void* self) {
	((QTreeWidget__Go*)self)->cppgoBase_updateEditorGeometries();
}
void QTreeWidget_verticalScrollbarAction_default(void* self, int action) {
	((QTreeWidget__Go*)self)->cppgoBase_verticalScrollbarAction(action);
}
void QTreeWidget_horizontalScrollbarValueChanged_default(void* self, int value) {
	((QTreeWidget__Go*)self)->cppgoBase_horizontalScrollbarValueChanged(value);
}
void QTreeWidget_closeEditor_default(void* self, void* editor, int hint) {
	((QTreeWidget__Go*)self)->cppgoBase_closeEditor((QWidget*)editor, hint);
}
void QTreeWidget_commitData_default(void* self, void* editor) {
	((QTreeWidget__Go*)self)->cppgoBase_commitData((QWidget*)editor);
}
void QTreeWidget_editorDestroyed_default(void* self, void* editor) {
	((QTreeWidget__Go*)self)->cppgoBase_editorDestroyed((QObject*)editor);
}
int QTreeWidget_edit_default(void* self, void* index, int trigger, void* event) {
	auto&& __r = ((QTreeWidget__Go*)self)->cppgoBase_edit(*(QModelIndex*)index, trigger, (QEvent*)event);
	return __r;
}
int QTreeWidget_selectionCommand_default(void* self, void* index, void* event) {
	auto&& __r = ((QTreeWidget__Go*)self)->cppgoBase_selectionCommand(*(QModelIndex*)index, (QEvent*)event);
	return (int)(__r);
}
void QTreeWidget_startDrag_default(void* self, int supportedActions) {
	((QTreeWidget__Go*)self)->cppgoBase_startDrag(QFlags<Qt::DropAction>(supportedActions));
}
void QTreeWidget_initViewItemOption_default(void* self, void* option) {
	((QTreeWidget__Go*)self)->cppgoBase_initViewItemOption((QStyleOptionViewItem*)option);
}
int QTreeWidget_focusNextPrevChild_default(void* self, int next) {
	auto&& __r = ((QTreeWidget__Go*)self)->cppgoBase_focusNextPrevChild(next);
	return __r;
}
void QTreeWidget_dragEnterEvent_default(void* self, void* event) {
	((QTreeWidget__Go*)self)->cppgoBase_dragEnterEvent((QDragEnterEvent*)event);
}
void QTreeWidget_dragLeaveEvent_default(void* self, void* event) {
	((QTreeWidget__Go*)self)->cppgoBase_dragLeaveEvent((QDragLeaveEvent*)event);
}
void QTreeWidget_focusInEvent_default(void* self, void* event) {
	((QTreeWidget__Go*)self)->cppgoBase_focusInEvent((QFocusEvent*)event);
}
void QTreeWidget_focusOutEvent_default(void* self, void* event) {
	((QTreeWidget__Go*)self)->cppgoBase_focusOutEvent((QFocusEvent*)event);
}
void QTreeWidget_resizeEvent_default(void* self, void* event) {
	((QTreeWidget__Go*)self)->cppgoBase_resizeEvent((QResizeEvent*)event);
}
void QTreeWidget_inputMethodEvent_default(void* self, void* event) {
	((QTreeWidget__Go*)self)->cppgoBase_inputMethodEvent((QInputMethodEvent*)event);
}
int QTreeWidget_eventFilter_default(void* self, void* object, void* event) {
	auto&& __r = ((QTreeWidget__Go*)self)->cppgoBase_eventFilter((QObject*)object, (QEvent*)event);
	return __r;
}
void* QTreeWidget_minimumSizeHint_default(void* self) {
	auto&& __r = ((QTreeWidget*)self)->QTreeWidget::minimumSizeHint();
	return new QSize(__r);
}
void* QTreeWidget_sizeHint_default(void* self) {
	auto&& __r = ((QTreeWidget*)self)->QTreeWidget::sizeHint();
	return new QSize(__r);
}
void QTreeWidget_setupViewport_default(void* self, void* viewport) {
	((QTreeWidget*)self)->QTreeWidget::setupViewport((QWidget*)viewport);
}
void QTreeWidget_wheelEvent_default(void* self, void* arg1) {
	((QTreeWidget__Go*)self)->cppgoBase_wheelEvent((QWheelEvent*)arg1);
}
void QTreeWidget_contextMenuEvent_default(void* self, void* arg1) {
	((QTreeWidget__Go*)self)->cppgoBase_contextMenuEvent((QContextMenuEvent*)arg1);
}
void QTreeWidget_initStyleOption_default(void* self, void* option) {
	((QTreeWidget__Go*)self)->cppgoBase_initStyleOption((QStyleOptionFrame*)option);
}
int QTreeWidget_devType_default(void* self) {
	auto&& __r = ((QTreeWidget*)self)->QTreeWidget::devType();
	return __r;
}
void QTreeWidget_setVisible_default(void* self, int visible) {
	((QTreeWidget*)self)->QTreeWidget::setVisible(visible);
}
int QTreeWidget_heightForWidth_default(void* self, int arg1) {
	auto&& __r = ((QTreeWidget*)self)->QTreeWidget::heightForWidth(arg1);
	return __r;
}
int QTreeWidget_hasHeightForWidth_default(void* self) {
	auto&& __r = ((QTreeWidget*)self)->QTreeWidget::hasHeightForWidth();
	return __r;
}
void* QTreeWidget_paintEngine_default(void* self) {
	auto&& __r = ((QTreeWidget*)self)->QTreeWidget::paintEngine();
	return (void*)(__r);
}
void QTreeWidget_keyReleaseEvent_default(void* self, void* event) {
	((QTreeWidget__Go*)self)->cppgoBase_keyReleaseEvent((QKeyEvent*)event);
}
void QTreeWidget_enterEvent_default(void* self, void* event) {
	((QTreeWidget__Go*)self)->cppgoBase_enterEvent((QEnterEvent*)event);
}
void QTreeWidget_leaveEvent_default(void* self, void* event) {
	((QTreeWidget__Go*)self)->cppgoBase_leaveEvent((QEvent*)event);
}
void QTreeWidget_moveEvent_default(void* self, void* event) {
	((QTreeWidget__Go*)self)->cppgoBase_moveEvent((QMoveEvent*)event);
}
void QTreeWidget_closeEvent_default(void* self, void* event) {
	((QTreeWidget__Go*)self)->cppgoBase_closeEvent((QCloseEvent*)event);
}
void QTreeWidget_tabletEvent_default(void* self, void* event) {
	((QTreeWidget__Go*)self)->cppgoBase_tabletEvent((QTabletEvent*)event);
}
void QTreeWidget_actionEvent_default(void* self, void* event) {
	((QTreeWidget__Go*)self)->cppgoBase_actionEvent((QActionEvent*)event);
}
void QTreeWidget_showEvent_default(void* self, void* event) {
	((QTreeWidget__Go*)self)->cppgoBase_showEvent((QShowEvent*)event);
}
void QTreeWidget_hideEvent_default(void* self, void* event) {
	((QTreeWidget__Go*)self)->cppgoBase_hideEvent((QHideEvent*)event);
}
int QTreeWidget_metric_default(void* self, int arg1) {
	auto&& __r = ((QTreeWidget__Go*)self)->cppgoBase_metric(arg1);
	return __r;
}
void QTreeWidget_initPainter_default(void* self, void* painter) {
	((QTreeWidget__Go*)self)->cppgoBase_initPainter((QPainter*)painter);
}
void* QTreeWidget_redirected_default(void* self, void* offset) {
	auto&& __r = ((QTreeWidget__Go*)self)->cppgoBase_redirected((QPoint*)offset);
	return (void*)(__r);
}
void* QTreeWidget_sharedPainter_default(void* self) {
	auto&& __r = ((QTreeWidget__Go*)self)->cppgoBase_sharedPainter();
	return (void*)(__r);
}
void QTreeWidget_childEvent_default(void* self, void* event) {
	((QTreeWidget__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QTreeWidget_customEvent_default(void* self, void* event) {
	((QTreeWidget__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QTreeWidget_connectNotify_default(void* self, void* signal) {
	((QTreeWidget__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QTreeWidget_disconnectNotify_default(void* self, void* signal) {
	((QTreeWidget__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QTreeWidget_connect_itemPressed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTreeWidget*)self, &QTreeWidget::itemPressed, (QTreeWidget*)self, [__id](QTreeWidgetItem* item, int column){
		if (cppgo_sig_void_voidp_int) ((void(*)(long long, void*, int))cppgo_sig_void_voidp_int)(__id, (void*)item, column);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTreeWidget);
	__conns_QTreeWidget[__id] = __c;
}
void QTreeWidget_connect_itemClicked(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTreeWidget*)self, &QTreeWidget::itemClicked, (QTreeWidget*)self, [__id](QTreeWidgetItem* item, int column){
		if (cppgo_sig_void_voidp_int) ((void(*)(long long, void*, int))cppgo_sig_void_voidp_int)(__id, (void*)item, column);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTreeWidget);
	__conns_QTreeWidget[__id] = __c;
}
void QTreeWidget_connect_itemDoubleClicked(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTreeWidget*)self, &QTreeWidget::itemDoubleClicked, (QTreeWidget*)self, [__id](QTreeWidgetItem* item, int column){
		if (cppgo_sig_void_voidp_int) ((void(*)(long long, void*, int))cppgo_sig_void_voidp_int)(__id, (void*)item, column);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTreeWidget);
	__conns_QTreeWidget[__id] = __c;
}
void QTreeWidget_connect_itemActivated(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTreeWidget*)self, &QTreeWidget::itemActivated, (QTreeWidget*)self, [__id](QTreeWidgetItem* item, int column){
		if (cppgo_sig_void_voidp_int) ((void(*)(long long, void*, int))cppgo_sig_void_voidp_int)(__id, (void*)item, column);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTreeWidget);
	__conns_QTreeWidget[__id] = __c;
}
void QTreeWidget_connect_itemEntered(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTreeWidget*)self, &QTreeWidget::itemEntered, (QTreeWidget*)self, [__id](QTreeWidgetItem* item, int column){
		if (cppgo_sig_void_voidp_int) ((void(*)(long long, void*, int))cppgo_sig_void_voidp_int)(__id, (void*)item, column);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTreeWidget);
	__conns_QTreeWidget[__id] = __c;
}
void QTreeWidget_connect_itemChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTreeWidget*)self, &QTreeWidget::itemChanged, (QTreeWidget*)self, [__id](QTreeWidgetItem* item, int column){
		if (cppgo_sig_void_voidp_int) ((void(*)(long long, void*, int))cppgo_sig_void_voidp_int)(__id, (void*)item, column);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTreeWidget);
	__conns_QTreeWidget[__id] = __c;
}
void QTreeWidget_connect_itemExpanded(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTreeWidget*)self, &QTreeWidget::itemExpanded, (QTreeWidget*)self, [__id](QTreeWidgetItem* item){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)item);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTreeWidget);
	__conns_QTreeWidget[__id] = __c;
}
void QTreeWidget_connect_itemCollapsed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTreeWidget*)self, &QTreeWidget::itemCollapsed, (QTreeWidget*)self, [__id](QTreeWidgetItem* item){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)item);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTreeWidget);
	__conns_QTreeWidget[__id] = __c;
}
void QTreeWidget_connect_currentItemChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTreeWidget*)self, &QTreeWidget::currentItemChanged, (QTreeWidget*)self, [__id](QTreeWidgetItem* current, QTreeWidgetItem* previous){
		if (cppgo_sig_void_voidp_voidp) ((void(*)(long long, void*, void*))cppgo_sig_void_voidp_voidp)(__id, (void*)current, (void*)previous);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTreeWidget);
	__conns_QTreeWidget[__id] = __c;
}
void QTreeWidget_connect_itemSelectionChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTreeWidget*)self, &QTreeWidget::itemSelectionChanged, (QTreeWidget*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTreeWidget);
	__conns_QTreeWidget[__id] = __c;
}
void QTreeWidget_connect_expanded(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTreeWidget*)self, &QTreeWidget::expanded, (QTreeWidget*)self, [__id](const QModelIndex& index){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&index);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTreeWidget);
	__conns_QTreeWidget[__id] = __c;
}
void QTreeWidget_connect_collapsed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTreeWidget*)self, &QTreeWidget::collapsed, (QTreeWidget*)self, [__id](const QModelIndex& index){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&index);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTreeWidget);
	__conns_QTreeWidget[__id] = __c;
}
void QTreeWidget_connect_pressed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTreeWidget*)self, &QTreeWidget::pressed, (QTreeWidget*)self, [__id](const QModelIndex& index){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&index);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTreeWidget);
	__conns_QTreeWidget[__id] = __c;
}
void QTreeWidget_connect_clicked(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTreeWidget*)self, &QTreeWidget::clicked, (QTreeWidget*)self, [__id](const QModelIndex& index){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&index);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTreeWidget);
	__conns_QTreeWidget[__id] = __c;
}
void QTreeWidget_connect_doubleClicked(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTreeWidget*)self, &QTreeWidget::doubleClicked, (QTreeWidget*)self, [__id](const QModelIndex& index){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&index);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTreeWidget);
	__conns_QTreeWidget[__id] = __c;
}
void QTreeWidget_connect_activated(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTreeWidget*)self, &QTreeWidget::activated, (QTreeWidget*)self, [__id](const QModelIndex& index){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&index);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTreeWidget);
	__conns_QTreeWidget[__id] = __c;
}
void QTreeWidget_connect_entered(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTreeWidget*)self, &QTreeWidget::entered, (QTreeWidget*)self, [__id](const QModelIndex& index){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&index);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTreeWidget);
	__conns_QTreeWidget[__id] = __c;
}
void QTreeWidget_connect_viewportEntered(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTreeWidget*)self, &QTreeWidget::viewportEntered, (QTreeWidget*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTreeWidget);
	__conns_QTreeWidget[__id] = __c;
}
void QTreeWidget_connect_iconSizeChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTreeWidget*)self, &QTreeWidget::iconSizeChanged, (QTreeWidget*)self, [__id](const QSize& size){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&size);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTreeWidget);
	__conns_QTreeWidget[__id] = __c;
}
void QTreeWidget_connect_windowTitleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTreeWidget*)self, &QTreeWidget::windowTitleChanged, (QTreeWidget*)self, [__id](const QString& title){
		QByteArray title_ba = title.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)title_ba.constData(), title_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTreeWidget);
	__conns_QTreeWidget[__id] = __c;
}
void QTreeWidget_connect_windowIconChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTreeWidget*)self, &QTreeWidget::windowIconChanged, (QTreeWidget*)self, [__id](const QIcon& icon){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&icon);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTreeWidget);
	__conns_QTreeWidget[__id] = __c;
}
void QTreeWidget_connect_windowIconTextChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTreeWidget*)self, &QTreeWidget::windowIconTextChanged, (QTreeWidget*)self, [__id](const QString& iconText){
		QByteArray iconText_ba = iconText.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)iconText_ba.constData(), iconText_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTreeWidget);
	__conns_QTreeWidget[__id] = __c;
}
void QTreeWidget_connect_customContextMenuRequested(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTreeWidget*)self, &QTreeWidget::customContextMenuRequested, (QTreeWidget*)self, [__id](const QPoint& pos){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&pos);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTreeWidget);
	__conns_QTreeWidget[__id] = __c;
}
void QTreeWidget_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTreeWidget*)self, &QTreeWidget::destroyed, (QTreeWidget*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTreeWidget);
	__conns_QTreeWidget[__id] = __c;
}
void QTreeWidget_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTreeWidget*)self, &QTreeWidget::destroyed, (QTreeWidget*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTreeWidget);
	__conns_QTreeWidget[__id] = __c;
}
void QTreeWidget_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTreeWidget*)self, &QTreeWidget::objectNameChanged, (QTreeWidget*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTreeWidget);
	__conns_QTreeWidget[__id] = __c;
}
void QTreeWidget__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QTreeWidget);
		auto __it = __conns_QTreeWidget.find(id);
		if (__it == __conns_QTreeWidget.end()) return;
		__c = __it->second;
		__conns_QTreeWidget.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QTreeWidget__Go_new(long long id) { return new QTreeWidget__Go(id); }
void QTreeWidget_delete(void* self) { delete (QTreeWidget*)self; }

}

