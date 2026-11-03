#include <QtCore/qcompare.h>
#include "gen_qpartialordering.h"

extern "C" {

void* QPartialOrdering_new(void* order) {
	return new QPartialOrdering(*(Qt::partial_ordering*)order);
}
void* QPartialOrdering_new2(void* stdorder) {
	return new QPartialOrdering(*(Qt::weak_ordering*)stdorder);
}
void* QPartialOrdering_new3(void* stdorder) {
	return new QPartialOrdering(*(Qt::strong_ordering*)stdorder);
}
void* QPartialOrdering_new4(void* arg1) {
	return new QPartialOrdering(*(QPartialOrdering*)arg1);
}
void* QPartialOrdering_ToQtPartialOrdering(void* self) {
	auto&& __r = static_cast<Qt::partial_ordering>(*(QPartialOrdering*)self);
	return new Qt::partial_ordering(__r);
}
int QPartialOrdering_OperatorEqual(void* self, void* rhs) {
	auto&& __r = (*(QPartialOrdering*)self) == (*(QPartialOrdering*)rhs);
	return __r;
}
int QPartialOrdering_OperatorNotEqual(void* self, void* rhs) {
	auto&& __r = (*(QPartialOrdering*)self) != (*(QPartialOrdering*)rhs);
	return __r;
}
int QPartialOrdering_OperatorEqual2(void* self, void* rhs) {
	auto&& __r = (*(QPartialOrdering*)self) == (*(Qt::partial_ordering*)rhs);
	return __r;
}
int QPartialOrdering_OperatorNotEqual2(void* self, void* rhs) {
	auto&& __r = (*(QPartialOrdering*)self) != (*(Qt::partial_ordering*)rhs);
	return __r;
}
void QPartialOrdering_delete(void* self) { delete (QPartialOrdering*)self; }

}

#include <QtCore/qstringconverter.h>
#include <QtCore/qstringconverter_base.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qstringview.h>
#include <QtCore/qbytearrayview.h>
#include <QtCore/qstring.h>
#include "gen_qstringencoder.h"

extern "C" {

void* QStringEncoder_base(void* self) { return (void*)static_cast<QStringConverter*>((QStringEncoder*)self); }
void QStringEncoder_chain(void* self, void** out) {
	QStringEncoder* __p = (QStringEncoder*)self;
	QStringConverter* __b0 = static_cast<QStringConverter*>(__p); out[0] = (void*)__b0;
	QStringConverterBase* __b1 = static_cast<QStringConverterBase*>(__b0); out[1] = (void*)__b1;
}
void* QStringEncoder_new() {
	return new QStringEncoder();
}
void* QStringEncoder_new2(int encoding) {
	return new QStringEncoder((QStringConverter::Encoding)(encoding));
}
void* QStringEncoder_new3(const char* name, long long name_len) {
	QAnyStringView name_qstr = QAnyStringView(name, name_len);
	return new QStringEncoder(name_qstr);
}
void* QStringEncoder_new4(int encoding, int flags) {
	return new QStringEncoder((QStringConverter::Encoding)(encoding), QFlags<QStringConverterBase::Flag>(flags));
}
void* QStringEncoder_new5(const char* name, long long name_len, int flags) {
	QAnyStringView name_qstr = QAnyStringView(name, name_len);
	return new QStringEncoder(name_qstr, QFlags<QStringConverterBase::Flag>(flags));
}
long long QStringEncoder_requiredSpace(void* self, long long inputLength) {
	auto&& __r = ((QStringEncoder*)self)->requiredSpace(inputLength);
	return __r;
}
int QStringEncoder_encodingForName(const char* name, long long name_len, int* __ok) {
	QAnyStringView name_qstr = QAnyStringView(name, name_len);
	auto&& __r = QStringEncoder::encodingForName(name_qstr);
	*__ok = __r.has_value() ? 1 : 0;
	return __r.has_value() ? (int)(*__r) : (int)0;
}
const char* QStringEncoder_nameForEncoding(int e) {
	auto&& __r = QStringEncoder::nameForEncoding((QStringConverter::Encoding)(e));
	return __r;
}
int QStringEncoder_encodingForData(void* data, int* __ok) {
	auto&& __r = QStringEncoder::encodingForData(*(QByteArrayView*)data);
	*__ok = __r.has_value() ? 1 : 0;
	return __r.has_value() ? (int)(*__r) : (int)0;
}
int QStringEncoder_encodingForHtml(void* data, int* __ok) {
	auto&& __r = QStringEncoder::encodingForHtml(*(QByteArrayView*)data);
	*__ok = __r.has_value() ? 1 : 0;
	return __r.has_value() ? (int)(*__r) : (int)0;
}
char** QStringEncoder_availableCodecs(long long* __out_len) {
	auto&& __r = QStringEncoder::availableCodecs();
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
int QStringEncoder_encodingForData2(void* data, unsigned short expectedFirstCharacter, int* __ok) {
	auto&& __r = QStringEncoder::encodingForData(*(QByteArrayView*)data, expectedFirstCharacter);
	*__ok = __r.has_value() ? 1 : 0;
	return __r.has_value() ? (int)(*__r) : (int)0;
}
void QStringEncoder_delete(void* self) { delete (QStringEncoder*)self; }

}

#include <QtCore/qabstracteventdispatcher.h>
#include <QtCore/qobject.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qeventloop.h>
#include <QtCore/qdeadlinetimer.h>
#include <QtCore/qthread.h>
#include <QtCore/qsocketnotifier.h>
#include <QtCore/qabstractnativeeventfilter.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qvariant.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qabstracteventdispatcherv2.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

static std::unordered_map<long long, QMetaObject::Connection> __conns_QAbstractEventDispatcherV2;
static std::mutex __connsMu_QAbstractEventDispatcherV2;

extern "C" {

void* QAbstractEventDispatcherV2_base(void* self) { return (void*)static_cast<QAbstractEventDispatcher*>((QAbstractEventDispatcherV2*)self); }
void QAbstractEventDispatcherV2_chain(void* self, void** out) {
	QAbstractEventDispatcherV2* __p = (QAbstractEventDispatcherV2*)self;
	QAbstractEventDispatcher* __b0 = static_cast<QAbstractEventDispatcher*>(__p); out[0] = (void*)__b0;
	QObject* __b1 = static_cast<QObject*>(__b0); out[1] = (void*)__b1;
}
void* QAbstractEventDispatcherV2_fromQObject(void* o) { return (void*)qobject_cast<QAbstractEventDispatcherV2*>((QObject*)o); }
char* QAbstractEventDispatcherV2_tr(const char* s, long long* __out_len) {
	auto&& __r = QAbstractEventDispatcherV2::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QAbstractEventDispatcherV2_processEventsWithDeadline(void* self, int flags, void* deadline) {
	auto&& __r = ((QAbstractEventDispatcherV2*)self)->processEventsWithDeadline(QFlags<QEventLoop::ProcessEventsFlag>(flags), *(QDeadlineTimer*)deadline);
	return __r;
}
void* QAbstractEventDispatcherV2_instance() {
	auto&& __r = QAbstractEventDispatcherV2::instance();
	return (void*)(__r);
}
void* QAbstractEventDispatcherV2_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QAbstractEventDispatcherV2::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QAbstractEventDispatcherV2_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QAbstractEventDispatcherV2::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
char* QAbstractEventDispatcherV2_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QAbstractEventDispatcherV2::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QAbstractEventDispatcherV2_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QAbstractEventDispatcherV2::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QAbstractEventDispatcherV2_instance2(void* thread) {
	auto&& __r = QAbstractEventDispatcherV2::instance((QThread*)thread);
	return (void*)(__r);
}
void* QAbstractEventDispatcherV2_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QAbstractEventDispatcherV2::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QAbstractEventDispatcherV2_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QAbstractEventDispatcherV2::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void QAbstractEventDispatcherV2_connect_aboutToBlock(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractEventDispatcherV2*)self, &QAbstractEventDispatcherV2::aboutToBlock, (QAbstractEventDispatcherV2*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractEventDispatcherV2);
	__conns_QAbstractEventDispatcherV2[__id] = __c;
}
void QAbstractEventDispatcherV2_connect_awake(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractEventDispatcherV2*)self, &QAbstractEventDispatcherV2::awake, (QAbstractEventDispatcherV2*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractEventDispatcherV2);
	__conns_QAbstractEventDispatcherV2[__id] = __c;
}
void QAbstractEventDispatcherV2_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractEventDispatcherV2*)self, &QAbstractEventDispatcherV2::destroyed, (QAbstractEventDispatcherV2*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractEventDispatcherV2);
	__conns_QAbstractEventDispatcherV2[__id] = __c;
}
void QAbstractEventDispatcherV2_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractEventDispatcherV2*)self, &QAbstractEventDispatcherV2::destroyed, (QAbstractEventDispatcherV2*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractEventDispatcherV2);
	__conns_QAbstractEventDispatcherV2[__id] = __c;
}
void QAbstractEventDispatcherV2_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractEventDispatcherV2*)self, &QAbstractEventDispatcherV2::objectNameChanged, (QAbstractEventDispatcherV2*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractEventDispatcherV2);
	__conns_QAbstractEventDispatcherV2[__id] = __c;
}
void QAbstractEventDispatcherV2__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QAbstractEventDispatcherV2);
		auto __it = __conns_QAbstractEventDispatcherV2.find(id);
		if (__it == __conns_QAbstractEventDispatcherV2.end()) return;
		__c = __it->second;
		__conns_QAbstractEventDispatcherV2.erase(__it);
	}
	QObject::disconnect(__c);
}
void QAbstractEventDispatcherV2_delete(void* self) { delete (QAbstractEventDispatcherV2*)self; }

}

#include <QtCore/qmutex.h>
#include <QtCore/qdeadlinetimer.h>
#include "gen_qmutex.h"

extern "C" {

void* QMutex_base(void* self) { return (void*)static_cast<QBasicMutex*>((QMutex*)self); }
void* QMutex_new() {
	return new QMutex();
}
int QMutex_try_lock(void* self) {
	auto&& __r = ((QMutex*)self)->try_lock();
	return __r;
}
int QMutex_tryLock(void* self, int timeout) {
	auto&& __r = ((QMutex*)self)->tryLock(timeout);
	return __r;
}
int QMutex_tryLock2(void* self, void* timeout) {
	auto&& __r = ((QMutex*)self)->tryLock(*(QDeadlineTimer*)timeout);
	return __r;
}
void QMutex_delete(void* self) { delete (QMutex*)self; }

}

#include <QtCore/qregularexpression.h>
#include <QtCore/qstring.h>
#include <QtCore/qstringview.h>
#include <QtCore/qnamespace.h>
#include "gen_qregularexpression.h"

extern "C" {

void* QRegularExpression_new() {
	return new QRegularExpression();
}
void* QRegularExpression_new2(const char* pattern, long long pattern_len) {
	QString pattern_qstr = QString::fromUtf8(pattern, pattern_len);
	return new QRegularExpression(pattern_qstr);
}
void* QRegularExpression_new3(void* re) {
	return new QRegularExpression(*(QRegularExpression*)re);
}
void* QRegularExpression_new4(const char* pattern, long long pattern_len, int options) {
	QString pattern_qstr = QString::fromUtf8(pattern, pattern_len);
	return new QRegularExpression(pattern_qstr, QFlags<QRegularExpression::PatternOption>(options));
}
int QRegularExpression_patternOptions(void* self) {
	auto&& __r = ((QRegularExpression*)self)->patternOptions();
	return (int)(__r);
}
void QRegularExpression_setPatternOptions(void* self, int options) {
	((QRegularExpression*)self)->setPatternOptions(QFlags<QRegularExpression::PatternOption>(options));
}
void* QRegularExpression_OperatorAssign(void* self, void* re) {
	auto&& __r = ((QRegularExpression*)self)->operator=(*(QRegularExpression*)re);
	return (void*)&(__r);
}
void QRegularExpression_swap(void* self, void* other) {
	((QRegularExpression*)self)->swap(*(QRegularExpression*)other);
}
char* QRegularExpression_pattern(void* self, long long* __out_len) {
	auto&& __r = ((QRegularExpression*)self)->pattern();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QRegularExpression_setPattern(void* self, const char* pattern, long long pattern_len) {
	QString pattern_qstr = QString::fromUtf8(pattern, pattern_len);
	((QRegularExpression*)self)->setPattern(pattern_qstr);
}
int QRegularExpression_isValid(void* self) {
	auto&& __r = ((QRegularExpression*)self)->isValid();
	return __r;
}
long long QRegularExpression_patternErrorOffset(void* self) {
	auto&& __r = ((QRegularExpression*)self)->patternErrorOffset();
	return __r;
}
char* QRegularExpression_errorString(void* self, long long* __out_len) {
	auto&& __r = ((QRegularExpression*)self)->errorString();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QRegularExpression_captureCount(void* self) {
	auto&& __r = ((QRegularExpression*)self)->captureCount();
	return __r;
}
char** QRegularExpression_namedCaptureGroups(void* self, long long* __out_len) {
	auto&& __r = ((QRegularExpression*)self)->namedCaptureGroups();
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
void* QRegularExpression_match(void* self, const char* subject, long long subject_len) {
	QString subject_qstr = QString::fromUtf8(subject, subject_len);
	auto&& __r = ((QRegularExpression*)self)->match(subject_qstr);
	return new QRegularExpressionMatch(__r);
}
void* QRegularExpression_match2(void* self, void* subjectView) {
	auto&& __r = ((QRegularExpression*)self)->match(*(QStringView*)subjectView);
	return new QRegularExpressionMatch(__r);
}
void* QRegularExpression_matchView(void* self, void* subjectView) {
	auto&& __r = ((QRegularExpression*)self)->matchView(*(QStringView*)subjectView);
	return new QRegularExpressionMatch(__r);
}
void* QRegularExpression_globalMatch(void* self, const char* subject, long long subject_len) {
	QString subject_qstr = QString::fromUtf8(subject, subject_len);
	auto&& __r = ((QRegularExpression*)self)->globalMatch(subject_qstr);
	return new QRegularExpressionMatchIterator(__r);
}
void* QRegularExpression_globalMatch2(void* self, void* subjectView) {
	auto&& __r = ((QRegularExpression*)self)->globalMatch(*(QStringView*)subjectView);
	return new QRegularExpressionMatchIterator(__r);
}
void* QRegularExpression_globalMatchView(void* self, void* subjectView) {
	auto&& __r = ((QRegularExpression*)self)->globalMatchView(*(QStringView*)subjectView);
	return new QRegularExpressionMatchIterator(__r);
}
void QRegularExpression_optimize(void* self) {
	((QRegularExpression*)self)->optimize();
}
char* QRegularExpression_escape(const char* str, long long str_len, long long* __out_len) {
	QString str_qstr = QString::fromUtf8(str, str_len);
	auto&& __r = QRegularExpression::escape(str_qstr);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QRegularExpression_wildcardToRegularExpression(const char* str, long long str_len, long long* __out_len) {
	QString str_qstr = QString::fromUtf8(str, str_len);
	auto&& __r = QRegularExpression::wildcardToRegularExpression(str_qstr);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QRegularExpression_anchoredPattern(const char* expression, long long expression_len, long long* __out_len) {
	QString expression_qstr = QString::fromUtf8(expression, expression_len);
	auto&& __r = QRegularExpression::anchoredPattern(expression_qstr);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QRegularExpression_escape2(void* str, long long* __out_len) {
	auto&& __r = QRegularExpression::escape(*(QStringView*)str);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QRegularExpression_wildcardToRegularExpression2(void* str, long long* __out_len) {
	auto&& __r = QRegularExpression::wildcardToRegularExpression(*(QStringView*)str);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QRegularExpression_anchoredPattern2(void* expression, long long* __out_len) {
	auto&& __r = QRegularExpression::anchoredPattern(*(QStringView*)expression);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QRegularExpression_fromWildcard(void* pattern) {
	auto&& __r = QRegularExpression::fromWildcard(*(QStringView*)pattern);
	return new QRegularExpression(__r);
}
void* QRegularExpression_match3(void* self, const char* subject, long long subject_len, long long offset) {
	QString subject_qstr = QString::fromUtf8(subject, subject_len);
	auto&& __r = ((QRegularExpression*)self)->match(subject_qstr, offset);
	return new QRegularExpressionMatch(__r);
}
void* QRegularExpression_match4(void* self, const char* subject, long long subject_len, long long offset, int matchType) {
	QString subject_qstr = QString::fromUtf8(subject, subject_len);
	auto&& __r = ((QRegularExpression*)self)->match(subject_qstr, offset, (QRegularExpression::MatchType)(matchType));
	return new QRegularExpressionMatch(__r);
}
void* QRegularExpression_match5(void* self, const char* subject, long long subject_len, long long offset, int matchType, int matchOptions) {
	QString subject_qstr = QString::fromUtf8(subject, subject_len);
	auto&& __r = ((QRegularExpression*)self)->match(subject_qstr, offset, (QRegularExpression::MatchType)(matchType), QFlags<QRegularExpression::MatchOption>(matchOptions));
	return new QRegularExpressionMatch(__r);
}
void* QRegularExpression_match6(void* self, void* subjectView, long long offset) {
	auto&& __r = ((QRegularExpression*)self)->match(*(QStringView*)subjectView, offset);
	return new QRegularExpressionMatch(__r);
}
void* QRegularExpression_match7(void* self, void* subjectView, long long offset, int matchType) {
	auto&& __r = ((QRegularExpression*)self)->match(*(QStringView*)subjectView, offset, (QRegularExpression::MatchType)(matchType));
	return new QRegularExpressionMatch(__r);
}
void* QRegularExpression_match8(void* self, void* subjectView, long long offset, int matchType, int matchOptions) {
	auto&& __r = ((QRegularExpression*)self)->match(*(QStringView*)subjectView, offset, (QRegularExpression::MatchType)(matchType), QFlags<QRegularExpression::MatchOption>(matchOptions));
	return new QRegularExpressionMatch(__r);
}
void* QRegularExpression_matchView2(void* self, void* subjectView, long long offset) {
	auto&& __r = ((QRegularExpression*)self)->matchView(*(QStringView*)subjectView, offset);
	return new QRegularExpressionMatch(__r);
}
void* QRegularExpression_matchView3(void* self, void* subjectView, long long offset, int matchType) {
	auto&& __r = ((QRegularExpression*)self)->matchView(*(QStringView*)subjectView, offset, (QRegularExpression::MatchType)(matchType));
	return new QRegularExpressionMatch(__r);
}
void* QRegularExpression_matchView4(void* self, void* subjectView, long long offset, int matchType, int matchOptions) {
	auto&& __r = ((QRegularExpression*)self)->matchView(*(QStringView*)subjectView, offset, (QRegularExpression::MatchType)(matchType), QFlags<QRegularExpression::MatchOption>(matchOptions));
	return new QRegularExpressionMatch(__r);
}
void* QRegularExpression_globalMatch3(void* self, const char* subject, long long subject_len, long long offset) {
	QString subject_qstr = QString::fromUtf8(subject, subject_len);
	auto&& __r = ((QRegularExpression*)self)->globalMatch(subject_qstr, offset);
	return new QRegularExpressionMatchIterator(__r);
}
void* QRegularExpression_globalMatch4(void* self, const char* subject, long long subject_len, long long offset, int matchType) {
	QString subject_qstr = QString::fromUtf8(subject, subject_len);
	auto&& __r = ((QRegularExpression*)self)->globalMatch(subject_qstr, offset, (QRegularExpression::MatchType)(matchType));
	return new QRegularExpressionMatchIterator(__r);
}
void* QRegularExpression_globalMatch5(void* self, const char* subject, long long subject_len, long long offset, int matchType, int matchOptions) {
	QString subject_qstr = QString::fromUtf8(subject, subject_len);
	auto&& __r = ((QRegularExpression*)self)->globalMatch(subject_qstr, offset, (QRegularExpression::MatchType)(matchType), QFlags<QRegularExpression::MatchOption>(matchOptions));
	return new QRegularExpressionMatchIterator(__r);
}
void* QRegularExpression_globalMatch6(void* self, void* subjectView, long long offset) {
	auto&& __r = ((QRegularExpression*)self)->globalMatch(*(QStringView*)subjectView, offset);
	return new QRegularExpressionMatchIterator(__r);
}
void* QRegularExpression_globalMatch7(void* self, void* subjectView, long long offset, int matchType) {
	auto&& __r = ((QRegularExpression*)self)->globalMatch(*(QStringView*)subjectView, offset, (QRegularExpression::MatchType)(matchType));
	return new QRegularExpressionMatchIterator(__r);
}
void* QRegularExpression_globalMatch8(void* self, void* subjectView, long long offset, int matchType, int matchOptions) {
	auto&& __r = ((QRegularExpression*)self)->globalMatch(*(QStringView*)subjectView, offset, (QRegularExpression::MatchType)(matchType), QFlags<QRegularExpression::MatchOption>(matchOptions));
	return new QRegularExpressionMatchIterator(__r);
}
void* QRegularExpression_globalMatchView2(void* self, void* subjectView, long long offset) {
	auto&& __r = ((QRegularExpression*)self)->globalMatchView(*(QStringView*)subjectView, offset);
	return new QRegularExpressionMatchIterator(__r);
}
void* QRegularExpression_globalMatchView3(void* self, void* subjectView, long long offset, int matchType) {
	auto&& __r = ((QRegularExpression*)self)->globalMatchView(*(QStringView*)subjectView, offset, (QRegularExpression::MatchType)(matchType));
	return new QRegularExpressionMatchIterator(__r);
}
void* QRegularExpression_globalMatchView4(void* self, void* subjectView, long long offset, int matchType, int matchOptions) {
	auto&& __r = ((QRegularExpression*)self)->globalMatchView(*(QStringView*)subjectView, offset, (QRegularExpression::MatchType)(matchType), QFlags<QRegularExpression::MatchOption>(matchOptions));
	return new QRegularExpressionMatchIterator(__r);
}
char* QRegularExpression_wildcardToRegularExpression3(const char* str, long long str_len, int options, long long* __out_len) {
	QString str_qstr = QString::fromUtf8(str, str_len);
	auto&& __r = QRegularExpression::wildcardToRegularExpression(str_qstr, QFlags<QRegularExpression::WildcardConversionOption>(options));
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QRegularExpression_wildcardToRegularExpression4(void* str, int options, long long* __out_len) {
	auto&& __r = QRegularExpression::wildcardToRegularExpression(*(QStringView*)str, QFlags<QRegularExpression::WildcardConversionOption>(options));
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QRegularExpression_fromWildcard2(void* pattern, int cs) {
	auto&& __r = QRegularExpression::fromWildcard(*(QStringView*)pattern, (Qt::CaseSensitivity)(cs));
	return new QRegularExpression(__r);
}
void* QRegularExpression_fromWildcard3(void* pattern, int cs, int options) {
	auto&& __r = QRegularExpression::fromWildcard(*(QStringView*)pattern, (Qt::CaseSensitivity)(cs), QFlags<QRegularExpression::WildcardConversionOption>(options));
	return new QRegularExpression(__r);
}
int QRegularExpression_OperatorEqual(void* self, void* rhs) {
	auto&& __r = (*(QRegularExpression*)self) == (*(QRegularExpression*)rhs);
	return __r;
}
int QRegularExpression_OperatorNotEqual(void* self, void* rhs) {
	auto&& __r = (*(QRegularExpression*)self) != (*(QRegularExpression*)rhs);
	return __r;
}
void QRegularExpression_delete(void* self) { delete (QRegularExpression*)self; }

}

#include <QtCore/qproperty.h>
#include <QtCore/qmetatype.h>
#include "gen_quntypedbindable.h"

extern "C" {

void* QUntypedBindable_new() {
	return new QUntypedBindable();
}
void* QUntypedBindable_new2(void* arg1) {
	return new QUntypedBindable(*(QUntypedBindable*)arg1);
}
int QUntypedBindable_isValid(void* self) {
	auto&& __r = ((QUntypedBindable*)self)->isValid();
	return __r;
}
int QUntypedBindable_isBindable(void* self) {
	auto&& __r = ((QUntypedBindable*)self)->isBindable();
	return __r;
}
int QUntypedBindable_isReadOnly(void* self) {
	auto&& __r = ((QUntypedBindable*)self)->isReadOnly();
	return __r;
}
void* QUntypedBindable_makeBinding(void* self) {
	auto&& __r = ((QUntypedBindable*)self)->makeBinding();
	return new QUntypedPropertyBinding(__r);
}
void* QUntypedBindable_takeBinding(void* self) {
	auto&& __r = ((QUntypedBindable*)self)->takeBinding();
	return new QUntypedPropertyBinding(__r);
}
void QUntypedBindable_observe(void* self, void* observer) {
	((QUntypedBindable*)self)->observe((QPropertyObserver*)observer);
}
void* QUntypedBindable_binding(void* self) {
	auto&& __r = ((QUntypedBindable*)self)->binding();
	return new QUntypedPropertyBinding(__r);
}
int QUntypedBindable_setBinding(void* self, void* binding) {
	auto&& __r = ((QUntypedBindable*)self)->setBinding(*(QUntypedPropertyBinding*)binding);
	return __r;
}
int QUntypedBindable_hasBinding(void* self) {
	auto&& __r = ((QUntypedBindable*)self)->hasBinding();
	return __r;
}
void* QUntypedBindable_metaType(void* self) {
	auto&& __r = ((QUntypedBindable*)self)->metaType();
	return new QMetaType(__r);
}
void* QUntypedBindable_makeBinding2(void* self, void* location) {
	auto&& __r = ((QUntypedBindable*)self)->makeBinding(*(QPropertyBindingSourceLocation*)location);
	return new QUntypedPropertyBinding(__r);
}
void QUntypedBindable_delete(void* self) { delete (QUntypedBindable*)self; }

}

#include <QtCore/qfactoryinterface.h>
#include <QtCore/qstring.h>
#include "gen_qfactoryinterface.h"
#include "gen_purego_cb.h"

class QFactoryInterface__Go : public QFactoryInterface {
public:
	long long __goId;
	unsigned char __goOn[1] = {};
	template<typename... A> QFactoryInterface__Go(long long id, A&&... a) : QFactoryInterface(std::forward<A>(a)...), __goId(id) {}
	QList<QString> keys() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virtp_voidp) {
			long long __len = 0;
			void* __a = ((void*(*)(long long, long long, long long*, int*))cppgo_virtp_voidp)(__goId, 6348254508799049261LL, &__len, &__h);
			if (__h) {
				QList<QString> __r; __r.reserve(__len);
				char** __aa = (char**)__a;
				for (long long __i = 0; __i < __len; __i++) { __r.append(QString::fromUtf8(__aa[__i])); free(__aa[__i]); }
				if (__a) free(__a);
				return __r;
			}
		}
		return QList<QString>{};
	}
};

extern "C" {

void QFactoryInterface__set_override(void* self, int idx) { ((QFactoryInterface__Go*)self)->__goOn[idx] = 1; }
char** QFactoryInterface_keys(void* self, long long* __out_len) {
	auto&& __r = ((QFactoryInterface*)self)->keys();
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
void* QFactoryInterface_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QFactoryInterface*)self)->operator=(*(QFactoryInterface*)arg1);
	return (void*)&(__r);
}
void* QFactoryInterface__Go_new(long long id) { return new QFactoryInterface__Go(id); }
void QFactoryInterface_delete(void* self) { delete (QFactoryInterface*)self; }

}

#include <QtCore/qpoint.h>
#include <QtGui/qtransform.h>
#include <QtGui/qmatrix4x4.h>
#include "gen_qpoint.h"

extern "C" {

void* QPoint_new() {
	return new QPoint();
}
void* QPoint_new2(int xpos, int ypos) {
	return new QPoint(xpos, ypos);
}
void* QPoint_new3(void* arg1) {
	return new QPoint(*(QPoint*)arg1);
}
int QPoint_isNull(void* self) {
	auto&& __r = ((QPoint*)self)->isNull();
	return __r;
}
int QPoint_x(void* self) {
	auto&& __r = ((QPoint*)self)->x();
	return __r;
}
int QPoint_y(void* self) {
	auto&& __r = ((QPoint*)self)->y();
	return __r;
}
void QPoint_setX(void* self, int x) {
	((QPoint*)self)->setX(x);
}
void QPoint_setY(void* self, int y) {
	((QPoint*)self)->setY(y);
}
int QPoint_manhattanLength(void* self) {
	auto&& __r = ((QPoint*)self)->manhattanLength();
	return __r;
}
void* QPoint_transposed(void* self) {
	auto&& __r = ((QPoint*)self)->transposed();
	return new QPoint(__r);
}
void* QPoint_OperatorPlusAssign(void* self, void* p) {
	auto&& __r = ((QPoint*)self)->operator+=(*(QPoint*)p);
	return (void*)&(__r);
}
void* QPoint_OperatorMinusAssign(void* self, void* p) {
	auto&& __r = ((QPoint*)self)->operator-=(*(QPoint*)p);
	return (void*)&(__r);
}
void* QPoint_OperatorMultiplyAssign(void* self, float factor) {
	auto&& __r = ((QPoint*)self)->operator*=(factor);
	return (void*)&(__r);
}
void* QPoint_OperatorMultiplyAssign2(void* self, double factor) {
	auto&& __r = ((QPoint*)self)->operator*=(factor);
	return (void*)&(__r);
}
void* QPoint_OperatorMultiplyAssign3(void* self, int factor) {
	auto&& __r = ((QPoint*)self)->operator*=(factor);
	return (void*)&(__r);
}
void* QPoint_OperatorDivideAssign(void* self, double divisor) {
	auto&& __r = ((QPoint*)self)->operator/=(divisor);
	return (void*)&(__r);
}
int QPoint_dotProduct(void* p1, void* p2) {
	auto&& __r = QPoint::dotProduct(*(QPoint*)p1, *(QPoint*)p2);
	return __r;
}
void* QPoint_toPointF(void* self) {
	auto&& __r = ((QPoint*)self)->toPointF();
	return new QPointF(__r);
}
void* QPoint_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QPoint*)self)->operator=(*(QPoint*)arg1);
	return (void*)&(__r);
}
int QPoint_OperatorEqual(void* self, void* rhs) {
	auto&& __r = (*(QPoint*)self) == (*(QPoint*)rhs);
	return __r;
}
int QPoint_OperatorNotEqual(void* self, void* rhs) {
	auto&& __r = (*(QPoint*)self) != (*(QPoint*)rhs);
	return __r;
}
void* QPoint_OperatorPlus(void* self, void* p2) {
	auto&& __r = (*(QPoint*)self) + (*(QPoint*)p2);
	return new QPoint(__r);
}
void* QPoint_OperatorMinus(void* self, void* p2) {
	auto&& __r = (*(QPoint*)self) - (*(QPoint*)p2);
	return new QPoint(__r);
}
void* QPoint_OperatorMultiply(void* self, float factor) {
	auto&& __r = (*(QPoint*)self) * (factor);
	return new QPoint(__r);
}
void* QPoint_OperatorMultiply2(void* self, double factor) {
	auto&& __r = (*(QPoint*)self) * (factor);
	return new QPoint(__r);
}
void* QPoint_OperatorMultiply3(void* self, int factor) {
	auto&& __r = (*(QPoint*)self) * (factor);
	return new QPoint(__r);
}
void* QPoint_OperatorPlus2(void* self) {
	auto&& __r = +(*(QPoint*)self);
	return new QPoint(__r);
}
void* QPoint_OperatorMinus2(void* self) {
	auto&& __r = -(*(QPoint*)self);
	return new QPoint(__r);
}
void* QPoint_OperatorDivide(void* self, double c) {
	auto&& __r = (*(QPoint*)self) / (c);
	return new QPoint(__r);
}
int QPoint_OperatorEqual2(void* self, void* rhs) {
	auto&& __r = (*(QPoint*)self) == (*(QPointF*)rhs);
	return __r;
}
int QPoint_OperatorNotEqual2(void* self, void* rhs) {
	auto&& __r = (*(QPoint*)self) != (*(QPointF*)rhs);
	return __r;
}
void* QPoint_OperatorMultiply4(void* self, void* m) {
	auto&& __r = (*(QPoint*)self) * (*(QTransform*)m);
	return new QPoint(__r);
}
void* QPoint_OperatorMultiply5(void* self, void* matrix) {
	auto&& __r = (*(QPoint*)self) * (*(QMatrix4x4*)matrix);
	return new QPoint(__r);
}
void QPoint_delete(void* self) { delete (QPoint*)self; }

}

#include <QtCore/qpermissions.h>
#include "gen_qcalendarpermission.h"

extern "C" {

void* QCalendarPermission_new() {
	return new QCalendarPermission();
}
void* QCalendarPermission_new2(void* other) {
	return new QCalendarPermission(*(QCalendarPermission*)other);
}
void QCalendarPermission_setAccessMode(void* self, int mode) {
	((QCalendarPermission*)self)->setAccessMode((QCalendarPermission::AccessMode)(mode));
}
int QCalendarPermission_accessMode(void* self) {
	auto&& __r = ((QCalendarPermission*)self)->accessMode();
	return (int)(__r);
}
void* QCalendarPermission_OperatorAssign(void* self, void* other) {
	auto&& __r = ((QCalendarPermission*)self)->operator=(*(QCalendarPermission*)other);
	return (void*)&(__r);
}
void QCalendarPermission_swap(void* self, void* other) {
	((QCalendarPermission*)self)->swap(*(QCalendarPermission*)other);
}
void QCalendarPermission_delete(void* self) { delete (QCalendarPermission*)self; }

}

#include <QtCore/qsize.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qmargins.h>
#include "gen_qsizef.h"

extern "C" {

void* QSizeF_new() {
	return new QSizeF();
}
void* QSizeF_new2(void* sz) {
	return new QSizeF(*(QSize*)sz);
}
void* QSizeF_new3(double w, double h) {
	return new QSizeF(w, h);
}
void* QSizeF_new4(void* arg1) {
	return new QSizeF(*(QSizeF*)arg1);
}
int QSizeF_isNull(void* self) {
	auto&& __r = ((QSizeF*)self)->isNull();
	return __r;
}
int QSizeF_isEmpty(void* self) {
	auto&& __r = ((QSizeF*)self)->isEmpty();
	return __r;
}
int QSizeF_isValid(void* self) {
	auto&& __r = ((QSizeF*)self)->isValid();
	return __r;
}
double QSizeF_width(void* self) {
	auto&& __r = ((QSizeF*)self)->width();
	return __r;
}
double QSizeF_height(void* self) {
	auto&& __r = ((QSizeF*)self)->height();
	return __r;
}
void QSizeF_setWidth(void* self, double w) {
	((QSizeF*)self)->setWidth(w);
}
void QSizeF_setHeight(void* self, double h) {
	((QSizeF*)self)->setHeight(h);
}
void QSizeF_transpose(void* self) {
	((QSizeF*)self)->transpose();
}
void* QSizeF_transposed(void* self) {
	auto&& __r = ((QSizeF*)self)->transposed();
	return new QSizeF(__r);
}
void QSizeF_scale(void* self, double w, double h, int mode) {
	((QSizeF*)self)->scale(w, h, (Qt::AspectRatioMode)(mode));
}
void QSizeF_scale2(void* self, void* s, int mode) {
	((QSizeF*)self)->scale(*(QSizeF*)s, (Qt::AspectRatioMode)(mode));
}
void* QSizeF_scaled(void* self, double w, double h, int mode) {
	auto&& __r = ((QSizeF*)self)->scaled(w, h, (Qt::AspectRatioMode)(mode));
	return new QSizeF(__r);
}
void* QSizeF_scaled2(void* self, void* s, int mode) {
	auto&& __r = ((QSizeF*)self)->scaled(*(QSizeF*)s, (Qt::AspectRatioMode)(mode));
	return new QSizeF(__r);
}
void* QSizeF_expandedTo(void* self, void* arg1) {
	auto&& __r = ((QSizeF*)self)->expandedTo(*(QSizeF*)arg1);
	return new QSizeF(__r);
}
void* QSizeF_boundedTo(void* self, void* arg1) {
	auto&& __r = ((QSizeF*)self)->boundedTo(*(QSizeF*)arg1);
	return new QSizeF(__r);
}
void* QSizeF_grownBy(void* self, void* m) {
	auto&& __r = ((QSizeF*)self)->grownBy(*(QMarginsF*)m);
	return new QSizeF(__r);
}
void* QSizeF_shrunkBy(void* self, void* m) {
	auto&& __r = ((QSizeF*)self)->shrunkBy(*(QMarginsF*)m);
	return new QSizeF(__r);
}
void* QSizeF_OperatorPlusAssign(void* self, void* arg1) {
	auto&& __r = ((QSizeF*)self)->operator+=(*(QSizeF*)arg1);
	return (void*)&(__r);
}
void* QSizeF_OperatorMinusAssign(void* self, void* arg1) {
	auto&& __r = ((QSizeF*)self)->operator-=(*(QSizeF*)arg1);
	return (void*)&(__r);
}
void* QSizeF_OperatorMultiplyAssign(void* self, double c) {
	auto&& __r = ((QSizeF*)self)->operator*=(c);
	return (void*)&(__r);
}
void* QSizeF_OperatorDivideAssign(void* self, double c) {
	auto&& __r = ((QSizeF*)self)->operator/=(c);
	return (void*)&(__r);
}
void* QSizeF_toSize(void* self) {
	auto&& __r = ((QSizeF*)self)->toSize();
	return new QSize(__r);
}
void* QSizeF_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QSizeF*)self)->operator=(*(QSizeF*)arg1);
	return (void*)&(__r);
}
int QSizeF_OperatorEqual(void* self, void* rhs) {
	auto&& __r = (*(QSizeF*)self) == (*(QSizeF*)rhs);
	return __r;
}
int QSizeF_OperatorNotEqual(void* self, void* rhs) {
	auto&& __r = (*(QSizeF*)self) != (*(QSizeF*)rhs);
	return __r;
}
int QSizeF_OperatorEqual2(void* self, void* rhs) {
	auto&& __r = (*(QSizeF*)self) == (*(QSize*)rhs);
	return __r;
}
int QSizeF_OperatorNotEqual2(void* self, void* rhs) {
	auto&& __r = (*(QSizeF*)self) != (*(QSize*)rhs);
	return __r;
}
void* QSizeF_OperatorPlus(void* self, void* s2) {
	auto&& __r = (*(QSizeF*)self) + (*(QSizeF*)s2);
	return new QSizeF(__r);
}
void* QSizeF_OperatorMinus(void* self, void* s2) {
	auto&& __r = (*(QSizeF*)self) - (*(QSizeF*)s2);
	return new QSizeF(__r);
}
void* QSizeF_OperatorMultiply(void* self, double c) {
	auto&& __r = (*(QSizeF*)self) * (c);
	return new QSizeF(__r);
}
void* QSizeF_OperatorDivide(void* self, double c) {
	auto&& __r = (*(QSizeF*)self) / (c);
	return new QSizeF(__r);
}
void QSizeF_delete(void* self) { delete (QSizeF*)self; }

}

#include <QtCore/qtextboundaryfinder.h>
#include <QtCore/qstring.h>
#include <QtCore/qchar.h>
#include <QtCore/qstringview.h>
#include "gen_qtextboundaryfinder.h"

extern "C" {

void* QTextBoundaryFinder_new() {
	return new QTextBoundaryFinder();
}
void* QTextBoundaryFinder_new2(void* other) {
	return new QTextBoundaryFinder(*(QTextBoundaryFinder*)other);
}
void* QTextBoundaryFinder_new3(int type_, const char* string_, long long string__len) {
	QString string__qstr = QString::fromUtf8(string_, string__len);
	return new QTextBoundaryFinder((QTextBoundaryFinder::BoundaryType)(type_), string__qstr);
}
void* QTextBoundaryFinder_new4(int type_, void* chars, long long length) {
	return new QTextBoundaryFinder((QTextBoundaryFinder::BoundaryType)(type_), (QChar*)chars, length);
}
void* QTextBoundaryFinder_new5(int type_, void* str) {
	return new QTextBoundaryFinder((QTextBoundaryFinder::BoundaryType)(type_), *(QStringView*)str);
}
void* QTextBoundaryFinder_new6(int type_, void* chars, long long length, char* buffer) {
	return new QTextBoundaryFinder((QTextBoundaryFinder::BoundaryType)(type_), (QChar*)chars, length, (unsigned char*)buffer);
}
void* QTextBoundaryFinder_new7(int type_, void* chars, long long length, char* buffer, long long bufferSize) {
	return new QTextBoundaryFinder((QTextBoundaryFinder::BoundaryType)(type_), (QChar*)chars, length, (unsigned char*)buffer, bufferSize);
}
void* QTextBoundaryFinder_new8(int type_, void* str, char* buffer) {
	return new QTextBoundaryFinder((QTextBoundaryFinder::BoundaryType)(type_), *(QStringView*)str, (unsigned char*)buffer);
}
void* QTextBoundaryFinder_new9(int type_, void* str, char* buffer, long long bufferSize) {
	return new QTextBoundaryFinder((QTextBoundaryFinder::BoundaryType)(type_), *(QStringView*)str, (unsigned char*)buffer, bufferSize);
}
void* QTextBoundaryFinder_OperatorAssign(void* self, void* other) {
	auto&& __r = ((QTextBoundaryFinder*)self)->operator=(*(QTextBoundaryFinder*)other);
	return (void*)&(__r);
}
void QTextBoundaryFinder_swap(void* self, void* other) {
	((QTextBoundaryFinder*)self)->swap(*(QTextBoundaryFinder*)other);
}
int QTextBoundaryFinder_isValid(void* self) {
	auto&& __r = ((QTextBoundaryFinder*)self)->isValid();
	return __r;
}
int QTextBoundaryFinder_type(void* self) {
	auto&& __r = ((QTextBoundaryFinder*)self)->type();
	return (int)(__r);
}
char* QTextBoundaryFinder_string(void* self, long long* __out_len) {
	auto&& __r = ((QTextBoundaryFinder*)self)->string();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QTextBoundaryFinder_toStart(void* self) {
	((QTextBoundaryFinder*)self)->toStart();
}
void QTextBoundaryFinder_toEnd(void* self) {
	((QTextBoundaryFinder*)self)->toEnd();
}
long long QTextBoundaryFinder_position(void* self) {
	auto&& __r = ((QTextBoundaryFinder*)self)->position();
	return __r;
}
void QTextBoundaryFinder_setPosition(void* self, long long position) {
	((QTextBoundaryFinder*)self)->setPosition(position);
}
long long QTextBoundaryFinder_toNextBoundary(void* self) {
	auto&& __r = ((QTextBoundaryFinder*)self)->toNextBoundary();
	return __r;
}
long long QTextBoundaryFinder_toPreviousBoundary(void* self) {
	auto&& __r = ((QTextBoundaryFinder*)self)->toPreviousBoundary();
	return __r;
}
int QTextBoundaryFinder_isAtBoundary(void* self) {
	auto&& __r = ((QTextBoundaryFinder*)self)->isAtBoundary();
	return __r;
}
int QTextBoundaryFinder_boundaryReasons(void* self) {
	auto&& __r = ((QTextBoundaryFinder*)self)->boundaryReasons();
	return (int)(__r);
}
void QTextBoundaryFinder_delete(void* self) { delete (QTextBoundaryFinder*)self; }

}

#include <QtGui/qimage.h>
#include <QtGui/qpaintdevice.h>
#include <QtCore/qsize.h>
#include <QtCore/qstring.h>
#include <QtCore/qvariant.h>
#include <QtCore/qrect.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qpoint.h>
#include <QtGui/qcolor.h>
#include <QtGui/qtransform.h>
#include <QtGui/qcolorspace.h>
#include <QtGui/qcolortransform.h>
#include <QtCore/qiodevice.h>
#include <QtCore/qbytearrayview.h>
#include <QtCore/qbytearray.h>
#include <QtGui/qpaintengine.h>
#include <QtGui/qpixelformat.h>
#include <QtGui/qpainter.h>
#include "gen_qimage.h"
#include "gen_purego_cb.h"

class QImage__Go : public QImage {
public:
	long long __goId;
	unsigned char __goOn[6] = {};
	template<typename... A> QImage__Go(long long id, A&&... a) : QImage(std::forward<A>(a)...), __goId(id) {}
	using QImage::mirrored_helper;
	using QImage::rgbSwapped_helper;
	using QImage::mirrored_inplace;
	using QImage::rgbSwapped_inplace;
	using QImage::convertToFormat_helper;
	using QImage::convertToFormat_inplace;
	using QImage::smoothScaled;
	using QImage::detachMetadata;
	using QImage::getDecodedMetricF;
	int cppgoBase_metric(int __a0) { return QImage::metric((QPaintDevice::PaintDeviceMetric)__a0); }
	void cppgoBase_initPainter(QPainter* __a0) { QImage::initPainter(__a0); }
	QPaintDevice* cppgoBase_redirected(QPoint* __a0) { return QImage::redirected(__a0); }
	QPainter* cppgoBase_sharedPainter() { return QImage::sharedPainter(); }
	int devType() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8659679893868524044LL, &__h); if (__h) return __r; }
		return QImage::devType();
	}
	QPaintEngine* paintEngine() const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 6398007404756346381LL, &__h); if (__h) return (QPaintEngine*)__r; }
		return QImage::paintEngine();
	}
	int metric(QPaintDevice::PaintDeviceMetric metric) const override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 427629649267311777LL, (int)(metric), &__h); if (__h) return __r; }
		return QImage::metric(metric);
	}
	void initPainter(QPainter* painter) const override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6477500296980050210LL, (void*)painter, &__h); if (__h) return; }
		QImage::initPainter(painter);
	}
	QPaintDevice* redirected(QPoint* offset) const override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 4127240633310655048LL, (void*)offset, &__h); if (__h) return (QPaintDevice*)__r; }
		return QImage::redirected(offset);
	}
	QPainter* sharedPainter() const override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 5792771328741774991LL, &__h); if (__h) return (QPainter*)__r; }
		return QImage::sharedPainter();
	}
};

extern "C" {

void* QImage_base(void* self) { return (void*)static_cast<QPaintDevice*>((QImage*)self); }
void QImage__set_override(void* self, int idx) { ((QImage__Go*)self)->__goOn[idx] = 1; }
void* QImage_new(long long __goid) {
	return new QImage__Go(__goid);
}
void* QImage_new2(long long __goid, void* size, int format) {
	return new QImage__Go(__goid, *(QSize*)size, (QImage::Format)(format));
}
void* QImage_new3(long long __goid, int width, int height, int format) {
	return new QImage__Go(__goid, width, height, (QImage::Format)(format));
}
void* QImage_new4(long long __goid, unsigned char* data, int width, int height, int format) {
	uchar data_local{};
	return new QImage__Go(__goid, &data_local, width, height, (QImage::Format)(format));
}
void* QImage_new5(long long __goid, const unsigned char* data, int width, int height, int format) {
	return new QImage__Go(__goid, data, width, height, (QImage::Format)(format));
}
void* QImage_new6(long long __goid, unsigned char* data, int width, int height, long long bytesPerLine, int format) {
	uchar data_local{};
	return new QImage__Go(__goid, &data_local, width, height, bytesPerLine, (QImage::Format)(format));
}
void* QImage_new7(long long __goid, const unsigned char* data, int width, int height, long long bytesPerLine, int format) {
	return new QImage__Go(__goid, data, width, height, bytesPerLine, (QImage::Format)(format));
}
void* QImage_new8(long long __goid, const char* fileName, long long fileName_len) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	return new QImage__Go(__goid, fileName_qstr);
}
void* QImage_new9(long long __goid, void* arg1) {
	return new QImage__Go(__goid, *(QImage*)arg1);
}
void* QImage_new10(long long __goid, const char* fileName, long long fileName_len, const char* format) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	return new QImage__Go(__goid, fileName_qstr, format);
}
void* QImage_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QImage*)self)->operator=(*(QImage*)arg1);
	return (void*)&(__r);
}
void QImage_swap(void* self, void* other) {
	((QImage*)self)->swap(*(QImage*)other);
}
int QImage_isNull(void* self) {
	auto&& __r = ((QImage*)self)->isNull();
	return __r;
}
int QImage_OperatorEqual(void* self, void* arg1) {
	auto&& __r = ((QImage*)self)->operator==(*(QImage*)arg1);
	return __r;
}
int QImage_OperatorNotEqual(void* self, void* arg1) {
	auto&& __r = ((QImage*)self)->operator!=(*(QImage*)arg1);
	return __r;
}
void* QImage_ToQVariant(void* self) {
	auto&& __r = static_cast<QVariant>(*(QImage*)self);
	return new QVariant(__r);
}
void QImage_detach(void* self) {
	((QImage*)self)->detach();
}
int QImage_isDetached(void* self) {
	auto&& __r = ((QImage*)self)->isDetached();
	return __r;
}
void* QImage_copy(void* self) {
	auto&& __r = ((QImage*)self)->copy();
	return new QImage(__r);
}
void* QImage_copy2(void* self, int x, int y, int w, int h) {
	auto&& __r = ((QImage*)self)->copy(x, y, w, h);
	return new QImage(__r);
}
int QImage_format(void* self) {
	auto&& __r = ((QImage*)self)->format();
	return (int)(__r);
}
void* QImage_convertToFormat(void* self, int f) {
	auto&& __r = ((QImage*)self)->convertToFormat((QImage::Format)(f));
	return new QImage(__r);
}
void* QImage_convertToFormat2(void* self, int f) {
	auto&& __r = ((QImage*)self)->convertToFormat((QImage::Format)(f));
	return new QImage(__r);
}
void* QImage_convertToFormat3(void* self, int f, unsigned int* colorTable, long long colorTable_len) {
	QList<unsigned int> colorTable_list; colorTable_list.reserve(colorTable_len);
	for (long long __i = 0; __i < colorTable_len; __i++) colorTable_list.append(colorTable[__i]);
	auto&& __r = ((QImage*)self)->convertToFormat((QImage::Format)(f), colorTable_list);
	return new QImage(__r);
}
int QImage_reinterpretAsFormat(void* self, int f) {
	auto&& __r = ((QImage*)self)->reinterpretAsFormat((QImage::Format)(f));
	return __r;
}
void* QImage_convertedTo(void* self, int f) {
	auto&& __r = ((QImage*)self)->convertedTo((QImage::Format)(f));
	return new QImage(__r);
}
void* QImage_convertedTo2(void* self, int f) {
	auto&& __r = ((QImage*)self)->convertedTo((QImage::Format)(f));
	return new QImage(__r);
}
void QImage_convertTo(void* self, int f) {
	((QImage*)self)->convertTo((QImage::Format)(f));
}
int QImage_width(void* self) {
	auto&& __r = ((QImage*)self)->width();
	return __r;
}
int QImage_height(void* self) {
	auto&& __r = ((QImage*)self)->height();
	return __r;
}
void* QImage_size(void* self) {
	auto&& __r = ((QImage*)self)->size();
	return new QSize(__r);
}
void* QImage_rect(void* self) {
	auto&& __r = ((QImage*)self)->rect();
	return new QRect(__r);
}
int QImage_depth(void* self) {
	auto&& __r = ((QImage*)self)->depth();
	return __r;
}
int QImage_colorCount(void* self) {
	auto&& __r = ((QImage*)self)->colorCount();
	return __r;
}
int QImage_bitPlaneCount(void* self) {
	auto&& __r = ((QImage*)self)->bitPlaneCount();
	return __r;
}
unsigned int QImage_color(void* self, int i) {
	auto&& __r = ((QImage*)self)->color(i);
	return __r;
}
void QImage_setColor(void* self, int i, unsigned int c) {
	((QImage*)self)->setColor(i, c);
}
void QImage_setColorCount(void* self, int arg1) {
	((QImage*)self)->setColorCount(arg1);
}
int QImage_allGray(void* self) {
	auto&& __r = ((QImage*)self)->allGray();
	return __r;
}
int QImage_isGrayscale(void* self) {
	auto&& __r = ((QImage*)self)->isGrayscale();
	return __r;
}
long long QImage_sizeInBytes(void* self) {
	auto&& __r = ((QImage*)self)->sizeInBytes();
	return __r;
}
long long QImage_bytesPerLine(void* self) {
	auto&& __r = ((QImage*)self)->bytesPerLine();
	return __r;
}
int QImage_valid(void* self, int x, int y) {
	auto&& __r = ((QImage*)self)->valid(x, y);
	return __r;
}
int QImage_valid2(void* self, void* pt) {
	auto&& __r = ((QImage*)self)->valid(*(QPoint*)pt);
	return __r;
}
int QImage_pixelIndex(void* self, int x, int y) {
	auto&& __r = ((QImage*)self)->pixelIndex(x, y);
	return __r;
}
int QImage_pixelIndex2(void* self, void* pt) {
	auto&& __r = ((QImage*)self)->pixelIndex(*(QPoint*)pt);
	return __r;
}
unsigned int QImage_pixel(void* self, int x, int y) {
	auto&& __r = ((QImage*)self)->pixel(x, y);
	return __r;
}
unsigned int QImage_pixel2(void* self, void* pt) {
	auto&& __r = ((QImage*)self)->pixel(*(QPoint*)pt);
	return __r;
}
void QImage_setPixel(void* self, int x, int y, unsigned int index_or_rgb) {
	((QImage*)self)->setPixel(x, y, index_or_rgb);
}
void QImage_setPixel2(void* self, void* pt, unsigned int index_or_rgb) {
	((QImage*)self)->setPixel(*(QPoint*)pt, index_or_rgb);
}
void* QImage_pixelColor(void* self, int x, int y) {
	auto&& __r = ((QImage*)self)->pixelColor(x, y);
	return new QColor(__r);
}
void* QImage_pixelColor2(void* self, void* pt) {
	auto&& __r = ((QImage*)self)->pixelColor(*(QPoint*)pt);
	return new QColor(__r);
}
void QImage_setPixelColor(void* self, int x, int y, void* c) {
	((QImage*)self)->setPixelColor(x, y, *(QColor*)c);
}
void QImage_setPixelColor2(void* self, void* pt, void* c) {
	((QImage*)self)->setPixelColor(*(QPoint*)pt, *(QColor*)c);
}
unsigned int* QImage_colorTable(void* self, long long* __out_len) {
	auto&& __r = ((QImage*)self)->colorTable();
	*__out_len = __r.size();
	unsigned int* __arr = (unsigned int*)malloc(sizeof(unsigned int) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = __e;
	return __arr;
}
void QImage_setColorTable(void* self, unsigned int* colors, long long colors_len) {
	QList<unsigned int> colors_list; colors_list.reserve(colors_len);
	for (long long __i = 0; __i < colors_len; __i++) colors_list.append(colors[__i]);
	((QImage*)self)->setColorTable(colors_list);
}
double QImage_devicePixelRatio(void* self) {
	auto&& __r = ((QImage*)self)->devicePixelRatio();
	return __r;
}
void QImage_setDevicePixelRatio(void* self, double scaleFactor) {
	((QImage*)self)->setDevicePixelRatio(scaleFactor);
}
void* QImage_deviceIndependentSize(void* self) {
	auto&& __r = ((QImage*)self)->deviceIndependentSize();
	return new QSizeF(__r);
}
void QImage_fill(void* self, unsigned int pixel) {
	((QImage*)self)->fill(pixel);
}
void QImage_fill2(void* self, void* color) {
	((QImage*)self)->fill(*(QColor*)color);
}
void QImage_fill3(void* self, int color) {
	((QImage*)self)->fill((Qt::GlobalColor)(color));
}
int QImage_hasAlphaChannel(void* self) {
	auto&& __r = ((QImage*)self)->hasAlphaChannel();
	return __r;
}
void QImage_setAlphaChannel(void* self, void* alphaChannel) {
	((QImage*)self)->setAlphaChannel(*(QImage*)alphaChannel);
}
void* QImage_createAlphaMask(void* self) {
	auto&& __r = ((QImage*)self)->createAlphaMask();
	return new QImage(__r);
}
void* QImage_createHeuristicMask(void* self) {
	auto&& __r = ((QImage*)self)->createHeuristicMask();
	return new QImage(__r);
}
void* QImage_createMaskFromColor(void* self, unsigned int color) {
	auto&& __r = ((QImage*)self)->createMaskFromColor(color);
	return new QImage(__r);
}
void* QImage_scaled(void* self, int w, int h) {
	auto&& __r = ((QImage*)self)->scaled(w, h);
	return new QImage(__r);
}
void* QImage_scaled2(void* self, void* s) {
	auto&& __r = ((QImage*)self)->scaled(*(QSize*)s);
	return new QImage(__r);
}
void* QImage_scaledToWidth(void* self, int w) {
	auto&& __r = ((QImage*)self)->scaledToWidth(w);
	return new QImage(__r);
}
void* QImage_scaledToHeight(void* self, int h) {
	auto&& __r = ((QImage*)self)->scaledToHeight(h);
	return new QImage(__r);
}
void* QImage_transformed(void* self, void* matrix) {
	auto&& __r = ((QImage*)self)->transformed(*(QTransform*)matrix);
	return new QImage(__r);
}
void* QImage_trueMatrix(void* arg1, int w, int h) {
	auto&& __r = QImage::trueMatrix(*(QTransform*)arg1, w, h);
	return new QTransform(__r);
}
void* QImage_mirrored(void* self) {
	auto&& __r = ((QImage*)self)->mirrored();
	return new QImage(__r);
}
void* QImage_mirrored2(void* self) {
	auto&& __r = ((QImage*)self)->mirrored();
	return new QImage(__r);
}
void QImage_mirror(void* self) {
	((QImage*)self)->mirror();
}
void* QImage_rgbSwapped(void* self) {
	auto&& __r = ((QImage*)self)->rgbSwapped();
	return new QImage(__r);
}
void* QImage_rgbSwapped2(void* self) {
	auto&& __r = ((QImage*)self)->rgbSwapped();
	return new QImage(__r);
}
void* QImage_flipped(void* self) {
	auto&& __r = ((QImage*)self)->flipped();
	return new QImage(__r);
}
void* QImage_flipped2(void* self) {
	auto&& __r = ((QImage*)self)->flipped();
	return new QImage(__r);
}
void QImage_flip(void* self) {
	((QImage*)self)->flip();
}
void QImage_rgbSwap(void* self) {
	((QImage*)self)->rgbSwap();
}
void QImage_invertPixels(void* self) {
	((QImage*)self)->invertPixels();
}
void* QImage_colorSpace(void* self) {
	auto&& __r = ((QImage*)self)->colorSpace();
	return new QColorSpace(__r);
}
void* QImage_convertedToColorSpace(void* self, void* colorSpace) {
	auto&& __r = ((QImage*)self)->convertedToColorSpace(*(QColorSpace*)colorSpace);
	return new QImage(__r);
}
void* QImage_convertedToColorSpace2(void* self, void* colorSpace, int format) {
	auto&& __r = ((QImage*)self)->convertedToColorSpace(*(QColorSpace*)colorSpace, (QImage::Format)(format));
	return new QImage(__r);
}
void* QImage_convertedToColorSpace3(void* self, void* colorSpace, int format) {
	auto&& __r = ((QImage*)self)->convertedToColorSpace(*(QColorSpace*)colorSpace, (QImage::Format)(format));
	return new QImage(__r);
}
void QImage_convertToColorSpace(void* self, void* colorSpace) {
	((QImage*)self)->convertToColorSpace(*(QColorSpace*)colorSpace);
}
void QImage_convertToColorSpace2(void* self, void* colorSpace, int format) {
	((QImage*)self)->convertToColorSpace(*(QColorSpace*)colorSpace, (QImage::Format)(format));
}
void QImage_setColorSpace(void* self, void* colorSpace) {
	((QImage*)self)->setColorSpace(*(QColorSpace*)colorSpace);
}
void* QImage_colorTransformed(void* self, void* transform) {
	auto&& __r = ((QImage*)self)->colorTransformed(*(QColorTransform*)transform);
	return new QImage(__r);
}
void* QImage_colorTransformed2(void* self, void* transform, int format) {
	auto&& __r = ((QImage*)self)->colorTransformed(*(QColorTransform*)transform, (QImage::Format)(format));
	return new QImage(__r);
}
void* QImage_colorTransformed3(void* self, void* transform) {
	auto&& __r = ((QImage*)self)->colorTransformed(*(QColorTransform*)transform);
	return new QImage(__r);
}
void* QImage_colorTransformed4(void* self, void* transform, int format) {
	auto&& __r = ((QImage*)self)->colorTransformed(*(QColorTransform*)transform, (QImage::Format)(format));
	return new QImage(__r);
}
void QImage_applyColorTransform(void* self, void* transform) {
	((QImage*)self)->applyColorTransform(*(QColorTransform*)transform);
}
void QImage_applyColorTransform2(void* self, void* transform, int format) {
	((QImage*)self)->applyColorTransform(*(QColorTransform*)transform, (QImage::Format)(format));
}
int QImage_load(void* self, void* device, const char* format) {
	auto&& __r = ((QImage*)self)->load((QIODevice*)device, format);
	return __r;
}
int QImage_load2(void* self, const char* fileName, long long fileName_len) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	auto&& __r = ((QImage*)self)->load(fileName_qstr);
	return __r;
}
int QImage_loadFromData(void* self, void* data) {
	auto&& __r = ((QImage*)self)->loadFromData(*(QByteArrayView*)data);
	return __r;
}
int QImage_loadFromData2(void* self, const unsigned char* buf, int len_) {
	auto&& __r = ((QImage*)self)->loadFromData(buf, len_);
	return __r;
}
int QImage_loadFromData3(void* self, void* data) {
	auto&& __r = ((QImage*)self)->loadFromData(*(QByteArray*)data);
	return __r;
}
int QImage_save(void* self, const char* fileName, long long fileName_len) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	auto&& __r = ((QImage*)self)->save(fileName_qstr);
	return __r;
}
int QImage_save2(void* self, void* device) {
	auto&& __r = ((QImage*)self)->save((QIODevice*)device);
	return __r;
}
void* QImage_fromData(void* data) {
	auto&& __r = QImage::fromData(*(QByteArrayView*)data);
	return new QImage(__r);
}
void* QImage_fromData2(const unsigned char* data, int size) {
	auto&& __r = QImage::fromData(data, size);
	return new QImage(__r);
}
void* QImage_fromData3(void* data) {
	auto&& __r = QImage::fromData(*(QByteArray*)data);
	return new QImage(__r);
}
long long QImage_cacheKey(void* self) {
	auto&& __r = ((QImage*)self)->cacheKey();
	return __r;
}
int QImage_dotsPerMeterX(void* self) {
	auto&& __r = ((QImage*)self)->dotsPerMeterX();
	return __r;
}
int QImage_dotsPerMeterY(void* self) {
	auto&& __r = ((QImage*)self)->dotsPerMeterY();
	return __r;
}
void QImage_setDotsPerMeterX(void* self, int arg1) {
	((QImage*)self)->setDotsPerMeterX(arg1);
}
void QImage_setDotsPerMeterY(void* self, int arg1) {
	((QImage*)self)->setDotsPerMeterY(arg1);
}
void* QImage_offset(void* self) {
	auto&& __r = ((QImage*)self)->offset();
	return new QPoint(__r);
}
void QImage_setOffset(void* self, void* arg1) {
	((QImage*)self)->setOffset(*(QPoint*)arg1);
}
char** QImage_textKeys(void* self, long long* __out_len) {
	auto&& __r = ((QImage*)self)->textKeys();
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
char* QImage_text(void* self, long long* __out_len) {
	auto&& __r = ((QImage*)self)->text();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QImage_setText(void* self, const char* key, long long key_len, const char* value, long long value_len) {
	QString key_qstr = QString::fromUtf8(key, key_len);
	QString value_qstr = QString::fromUtf8(value, value_len);
	((QImage*)self)->setText(key_qstr, value_qstr);
}
void* QImage_pixelFormat(void* self) {
	auto&& __r = ((QImage*)self)->pixelFormat();
	return new QPixelFormat(__r);
}
void* QImage_toPixelFormat(int format) {
	auto&& __r = QImage::toPixelFormat((QImage::Format)(format));
	return new QPixelFormat(__r);
}
int QImage_toImageFormat(void* format) {
	auto&& __r = QImage::toImageFormat(*(QPixelFormat*)format);
	return (int)(__r);
}
void* QImage_mirrored_helper(void* self, int horizontal, int vertical) {
	auto&& __r = ((QImage__Go*)self)->mirrored_helper(horizontal, vertical);
	return new QImage(__r);
}
void* QImage_rgbSwapped_helper(void* self) {
	auto&& __r = ((QImage__Go*)self)->rgbSwapped_helper();
	return new QImage(__r);
}
void QImage_mirrored_inplace(void* self, int horizontal, int vertical) {
	((QImage__Go*)self)->mirrored_inplace(horizontal, vertical);
}
void QImage_rgbSwapped_inplace(void* self) {
	((QImage__Go*)self)->rgbSwapped_inplace();
}
void* QImage_convertToFormat_helper(void* self, int format, int flags) {
	auto&& __r = ((QImage__Go*)self)->convertToFormat_helper((QImage::Format)(format), QFlags<Qt::ImageConversionFlag>(flags));
	return new QImage(__r);
}
int QImage_convertToFormat_inplace(void* self, int format, int flags) {
	auto&& __r = ((QImage__Go*)self)->convertToFormat_inplace((QImage::Format)(format), QFlags<Qt::ImageConversionFlag>(flags));
	return __r;
}
void* QImage_smoothScaled(void* self, int w, int h) {
	auto&& __r = ((QImage__Go*)self)->smoothScaled(w, h);
	return new QImage(__r);
}
void QImage_detachMetadata(void* self) {
	((QImage__Go*)self)->detachMetadata();
}
double QImage_devicePixelRatioFScale() {
	auto&& __r = QImage::devicePixelRatioFScale();
	return __r;
}
int QImage_encodeMetricF(int metric, double value) {
	auto&& __r = QImage::encodeMetricF((QPaintDevice::PaintDeviceMetric)(metric), value);
	return __r;
}
double QImage_getDecodedMetricF(void* self, int metricA, int metricB) {
	auto&& __r = ((QImage__Go*)self)->getDecodedMetricF((QPaintDevice::PaintDeviceMetric)(metricA), (QPaintDevice::PaintDeviceMetric)(metricB));
	return __r;
}
void* QImage_copy3(void* self, void* rect) {
	auto&& __r = ((QImage*)self)->copy(*(QRect*)rect);
	return new QImage(__r);
}
void* QImage_convertToFormat4(void* self, int f, int flags) {
	auto&& __r = ((QImage*)self)->convertToFormat((QImage::Format)(f), QFlags<Qt::ImageConversionFlag>(flags));
	return new QImage(__r);
}
void* QImage_convertToFormat5(void* self, int f, int flags) {
	auto&& __r = ((QImage*)self)->convertToFormat((QImage::Format)(f), QFlags<Qt::ImageConversionFlag>(flags));
	return new QImage(__r);
}
void* QImage_convertToFormat6(void* self, int f, unsigned int* colorTable, long long colorTable_len, int flags) {
	QList<unsigned int> colorTable_list; colorTable_list.reserve(colorTable_len);
	for (long long __i = 0; __i < colorTable_len; __i++) colorTable_list.append(colorTable[__i]);
	auto&& __r = ((QImage*)self)->convertToFormat((QImage::Format)(f), colorTable_list, QFlags<Qt::ImageConversionFlag>(flags));
	return new QImage(__r);
}
void* QImage_convertedTo3(void* self, int f, int flags) {
	auto&& __r = ((QImage*)self)->convertedTo((QImage::Format)(f), QFlags<Qt::ImageConversionFlag>(flags));
	return new QImage(__r);
}
void* QImage_convertedTo4(void* self, int f, int flags) {
	auto&& __r = ((QImage*)self)->convertedTo((QImage::Format)(f), QFlags<Qt::ImageConversionFlag>(flags));
	return new QImage(__r);
}
void QImage_convertTo2(void* self, int f, int flags) {
	((QImage*)self)->convertTo((QImage::Format)(f), QFlags<Qt::ImageConversionFlag>(flags));
}
void* QImage_createAlphaMask2(void* self, int flags) {
	auto&& __r = ((QImage*)self)->createAlphaMask(QFlags<Qt::ImageConversionFlag>(flags));
	return new QImage(__r);
}
void* QImage_createHeuristicMask2(void* self, int clipTight) {
	auto&& __r = ((QImage*)self)->createHeuristicMask(clipTight);
	return new QImage(__r);
}
void* QImage_createMaskFromColor2(void* self, unsigned int color, int mode) {
	auto&& __r = ((QImage*)self)->createMaskFromColor(color, (Qt::MaskMode)(mode));
	return new QImage(__r);
}
void* QImage_scaled3(void* self, int w, int h, int aspectMode) {
	auto&& __r = ((QImage*)self)->scaled(w, h, (Qt::AspectRatioMode)(aspectMode));
	return new QImage(__r);
}
void* QImage_scaled4(void* self, int w, int h, int aspectMode, int mode) {
	auto&& __r = ((QImage*)self)->scaled(w, h, (Qt::AspectRatioMode)(aspectMode), (Qt::TransformationMode)(mode));
	return new QImage(__r);
}
void* QImage_scaled5(void* self, void* s, int aspectMode) {
	auto&& __r = ((QImage*)self)->scaled(*(QSize*)s, (Qt::AspectRatioMode)(aspectMode));
	return new QImage(__r);
}
void* QImage_scaled6(void* self, void* s, int aspectMode, int mode) {
	auto&& __r = ((QImage*)self)->scaled(*(QSize*)s, (Qt::AspectRatioMode)(aspectMode), (Qt::TransformationMode)(mode));
	return new QImage(__r);
}
void* QImage_scaledToWidth2(void* self, int w, int mode) {
	auto&& __r = ((QImage*)self)->scaledToWidth(w, (Qt::TransformationMode)(mode));
	return new QImage(__r);
}
void* QImage_scaledToHeight2(void* self, int h, int mode) {
	auto&& __r = ((QImage*)self)->scaledToHeight(h, (Qt::TransformationMode)(mode));
	return new QImage(__r);
}
void* QImage_transformed2(void* self, void* matrix, int mode) {
	auto&& __r = ((QImage*)self)->transformed(*(QTransform*)matrix, (Qt::TransformationMode)(mode));
	return new QImage(__r);
}
void* QImage_mirrored3(void* self, int horizontally) {
	auto&& __r = ((QImage*)self)->mirrored(horizontally);
	return new QImage(__r);
}
void* QImage_mirrored4(void* self, int horizontally, int vertically) {
	auto&& __r = ((QImage*)self)->mirrored(horizontally, vertically);
	return new QImage(__r);
}
void* QImage_mirrored5(void* self, int horizontally) {
	auto&& __r = ((QImage*)self)->mirrored(horizontally);
	return new QImage(__r);
}
void* QImage_mirrored6(void* self, int horizontally, int vertically) {
	auto&& __r = ((QImage*)self)->mirrored(horizontally, vertically);
	return new QImage(__r);
}
void QImage_mirror2(void* self, int horizontally) {
	((QImage*)self)->mirror(horizontally);
}
void QImage_mirror3(void* self, int horizontally, int vertically) {
	((QImage*)self)->mirror(horizontally, vertically);
}
void* QImage_flipped3(void* self, int orient) {
	auto&& __r = ((QImage*)self)->flipped(QFlags<Qt::Orientation>(orient));
	return new QImage(__r);
}
void* QImage_flipped4(void* self, int orient) {
	auto&& __r = ((QImage*)self)->flipped(QFlags<Qt::Orientation>(orient));
	return new QImage(__r);
}
void QImage_flip2(void* self, int orient) {
	((QImage*)self)->flip(QFlags<Qt::Orientation>(orient));
}
void QImage_invertPixels2(void* self, int arg1) {
	((QImage*)self)->invertPixels((QImage::InvertMode)(arg1));
}
void* QImage_convertedToColorSpace4(void* self, void* colorSpace, int format, int flags) {
	auto&& __r = ((QImage*)self)->convertedToColorSpace(*(QColorSpace*)colorSpace, (QImage::Format)(format), QFlags<Qt::ImageConversionFlag>(flags));
	return new QImage(__r);
}
void* QImage_convertedToColorSpace5(void* self, void* colorSpace, int format, int flags) {
	auto&& __r = ((QImage*)self)->convertedToColorSpace(*(QColorSpace*)colorSpace, (QImage::Format)(format), QFlags<Qt::ImageConversionFlag>(flags));
	return new QImage(__r);
}
void QImage_convertToColorSpace3(void* self, void* colorSpace, int format, int flags) {
	((QImage*)self)->convertToColorSpace(*(QColorSpace*)colorSpace, (QImage::Format)(format), QFlags<Qt::ImageConversionFlag>(flags));
}
void* QImage_colorTransformed5(void* self, void* transform, int format, int flags) {
	auto&& __r = ((QImage*)self)->colorTransformed(*(QColorTransform*)transform, (QImage::Format)(format), QFlags<Qt::ImageConversionFlag>(flags));
	return new QImage(__r);
}
void* QImage_colorTransformed6(void* self, void* transform, int format, int flags) {
	auto&& __r = ((QImage*)self)->colorTransformed(*(QColorTransform*)transform, (QImage::Format)(format), QFlags<Qt::ImageConversionFlag>(flags));
	return new QImage(__r);
}
void QImage_applyColorTransform3(void* self, void* transform, int format, int flags) {
	((QImage*)self)->applyColorTransform(*(QColorTransform*)transform, (QImage::Format)(format), QFlags<Qt::ImageConversionFlag>(flags));
}
int QImage_load3(void* self, const char* fileName, long long fileName_len, const char* format) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	auto&& __r = ((QImage*)self)->load(fileName_qstr, format);
	return __r;
}
int QImage_loadFromData4(void* self, void* data, const char* format) {
	auto&& __r = ((QImage*)self)->loadFromData(*(QByteArrayView*)data, format);
	return __r;
}
int QImage_loadFromData5(void* self, const unsigned char* buf, int len_, const char* format) {
	auto&& __r = ((QImage*)self)->loadFromData(buf, len_, format);
	return __r;
}
int QImage_loadFromData6(void* self, void* data, const char* format) {
	auto&& __r = ((QImage*)self)->loadFromData(*(QByteArray*)data, format);
	return __r;
}
int QImage_save3(void* self, const char* fileName, long long fileName_len, const char* format) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	auto&& __r = ((QImage*)self)->save(fileName_qstr, format);
	return __r;
}
int QImage_save4(void* self, const char* fileName, long long fileName_len, const char* format, int quality) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	auto&& __r = ((QImage*)self)->save(fileName_qstr, format, quality);
	return __r;
}
int QImage_save5(void* self, void* device, const char* format) {
	auto&& __r = ((QImage*)self)->save((QIODevice*)device, format);
	return __r;
}
int QImage_save6(void* self, void* device, const char* format, int quality) {
	auto&& __r = ((QImage*)self)->save((QIODevice*)device, format, quality);
	return __r;
}
void* QImage_fromData4(void* data, const char* format) {
	auto&& __r = QImage::fromData(*(QByteArrayView*)data, format);
	return new QImage(__r);
}
void* QImage_fromData5(const unsigned char* data, int size, const char* format) {
	auto&& __r = QImage::fromData(data, size, format);
	return new QImage(__r);
}
void* QImage_fromData6(void* data, const char* format) {
	auto&& __r = QImage::fromData(*(QByteArray*)data, format);
	return new QImage(__r);
}
char* QImage_text2(void* self, const char* key, long long key_len, long long* __out_len) {
	QString key_qstr = QString::fromUtf8(key, key_len);
	auto&& __r = ((QImage*)self)->text(key_qstr);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QImage_detachMetadata2(void* self, int invalidateCache) {
	((QImage__Go*)self)->detachMetadata(invalidateCache);
}
void* QImage_toQVariant(void* self) {
	auto&& __r = QVariant::fromValue(*(QImage*)self);
	return new QVariant(__r);
}
int QImage_devType_default(void* self) {
	auto&& __r = ((QImage*)self)->QImage::devType();
	return __r;
}
void* QImage_paintEngine_default(void* self) {
	auto&& __r = ((QImage*)self)->QImage::paintEngine();
	return (void*)(__r);
}
int QImage_metric_default(void* self, int metric) {
	auto&& __r = ((QImage__Go*)self)->cppgoBase_metric(metric);
	return __r;
}
void QImage_initPainter_default(void* self, void* painter) {
	((QImage__Go*)self)->cppgoBase_initPainter((QPainter*)painter);
}
void* QImage_redirected_default(void* self, void* offset) {
	auto&& __r = ((QImage__Go*)self)->cppgoBase_redirected((QPoint*)offset);
	return (void*)(__r);
}
void* QImage_sharedPainter_default(void* self) {
	auto&& __r = ((QImage__Go*)self)->cppgoBase_sharedPainter();
	return (void*)(__r);
}
void* QImage__Go_new(long long id) { return new QImage__Go(id); }
void QImage_delete(void* self) { delete (QImage*)self; }

}

#include <QtGui/qevent.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qstring.h>
#include <QtGui/qinputdevice.h>
#include <QtGui/qkeysequence.h>
#include "gen_qkeyevent.h"

extern "C" {

void* QKeyEvent_base(void* self) { return (void*)static_cast<QInputEvent*>((QKeyEvent*)self); }
void QKeyEvent_chain(void* self, void** out) {
	QKeyEvent* __p = (QKeyEvent*)self;
	QInputEvent* __b0 = static_cast<QInputEvent*>(__p); out[0] = (void*)__b0;
	QEvent* __b1 = static_cast<QEvent*>(__b0); out[1] = (void*)__b1;
}
void* QKeyEvent_new(int type_, int key, int modifiers) {
	return new QKeyEvent((QEvent::Type)(type_), key, QFlags<Qt::KeyboardModifier>(modifiers));
}
void* QKeyEvent_new2(int type_, int key, int modifiers, unsigned int nativeScanCode, unsigned int nativeVirtualKey, unsigned int nativeModifiers) {
	return new QKeyEvent((QEvent::Type)(type_), key, QFlags<Qt::KeyboardModifier>(modifiers), nativeScanCode, nativeVirtualKey, nativeModifiers);
}
void* QKeyEvent_new3(int type_, int key, int modifiers, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	return new QKeyEvent((QEvent::Type)(type_), key, QFlags<Qt::KeyboardModifier>(modifiers), text_qstr);
}
void* QKeyEvent_new4(int type_, int key, int modifiers, const char* text, long long text_len, int autorep) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	return new QKeyEvent((QEvent::Type)(type_), key, QFlags<Qt::KeyboardModifier>(modifiers), text_qstr, autorep);
}
void* QKeyEvent_new5(int type_, int key, int modifiers, const char* text, long long text_len, int autorep, unsigned short count) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	return new QKeyEvent((QEvent::Type)(type_), key, QFlags<Qt::KeyboardModifier>(modifiers), text_qstr, autorep, count);
}
void* QKeyEvent_new6(int type_, int key, int modifiers, unsigned int nativeScanCode, unsigned int nativeVirtualKey, unsigned int nativeModifiers, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	return new QKeyEvent((QEvent::Type)(type_), key, QFlags<Qt::KeyboardModifier>(modifiers), nativeScanCode, nativeVirtualKey, nativeModifiers, text_qstr);
}
void* QKeyEvent_new7(int type_, int key, int modifiers, unsigned int nativeScanCode, unsigned int nativeVirtualKey, unsigned int nativeModifiers, const char* text, long long text_len, int autorep) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	return new QKeyEvent((QEvent::Type)(type_), key, QFlags<Qt::KeyboardModifier>(modifiers), nativeScanCode, nativeVirtualKey, nativeModifiers, text_qstr, autorep);
}
void* QKeyEvent_new8(int type_, int key, int modifiers, unsigned int nativeScanCode, unsigned int nativeVirtualKey, unsigned int nativeModifiers, const char* text, long long text_len, int autorep, unsigned short count) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	return new QKeyEvent((QEvent::Type)(type_), key, QFlags<Qt::KeyboardModifier>(modifiers), nativeScanCode, nativeVirtualKey, nativeModifiers, text_qstr, autorep, count);
}
void* QKeyEvent_new9(int type_, int key, int modifiers, unsigned int nativeScanCode, unsigned int nativeVirtualKey, unsigned int nativeModifiers, const char* text, long long text_len, int autorep, unsigned short count, void* device) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	return new QKeyEvent((QEvent::Type)(type_), key, QFlags<Qt::KeyboardModifier>(modifiers), nativeScanCode, nativeVirtualKey, nativeModifiers, text_qstr, autorep, count, (QInputDevice*)device);
}
int QKeyEvent_key(void* self) {
	auto&& __r = ((QKeyEvent*)self)->key();
	return __r;
}
int QKeyEvent_matches(void* self, int key) {
	auto&& __r = ((QKeyEvent*)self)->matches((QKeySequence::StandardKey)(key));
	return __r;
}
int QKeyEvent_modifiers(void* self) {
	auto&& __r = ((QKeyEvent*)self)->modifiers();
	return (int)(__r);
}
void* QKeyEvent_keyCombination(void* self) {
	auto&& __r = ((QKeyEvent*)self)->keyCombination();
	return new QKeyCombination(__r);
}
char* QKeyEvent_text(void* self, long long* __out_len) {
	auto&& __r = ((QKeyEvent*)self)->text();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QKeyEvent_isAutoRepeat(void* self) {
	auto&& __r = ((QKeyEvent*)self)->isAutoRepeat();
	return __r;
}
int QKeyEvent_count(void* self) {
	auto&& __r = ((QKeyEvent*)self)->count();
	return __r;
}
unsigned int QKeyEvent_nativeScanCode(void* self) {
	auto&& __r = ((QKeyEvent*)self)->nativeScanCode();
	return __r;
}
unsigned int QKeyEvent_nativeVirtualKey(void* self) {
	auto&& __r = ((QKeyEvent*)self)->nativeVirtualKey();
	return __r;
}
unsigned int QKeyEvent_nativeModifiers(void* self) {
	auto&& __r = ((QKeyEvent*)self)->nativeModifiers();
	return __r;
}
int QKeyEvent_registerEventType() {
	auto&& __r = QKeyEvent::registerEventType();
	return __r;
}
int QKeyEvent_registerEventType2(int hint) {
	auto&& __r = QKeyEvent::registerEventType(hint);
	return __r;
}
void QKeyEvent_delete(void* self) { delete (QKeyEvent*)self; }

}

#include <QtGui/qevent.h>
#include <QtCore/qcoreevent.h>
#include "gen_qtoolbarchangeevent.h"

extern "C" {

void* QToolBarChangeEvent_base(void* self) { return (void*)static_cast<QEvent*>((QToolBarChangeEvent*)self); }
void* QToolBarChangeEvent_new(int t) {
	return new QToolBarChangeEvent(t);
}
int QToolBarChangeEvent_toggle(void* self) {
	auto&& __r = ((QToolBarChangeEvent*)self)->toggle();
	return __r;
}
int QToolBarChangeEvent_registerEventType() {
	auto&& __r = QToolBarChangeEvent::registerEventType();
	return __r;
}
int QToolBarChangeEvent_registerEventType2(int hint) {
	auto&& __r = QToolBarChangeEvent::registerEventType(hint);
	return __r;
}
void QToolBarChangeEvent_delete(void* self) { delete (QToolBarChangeEvent*)self; }

}

#include <QtGui/qtextoption.h>
#include <QtCore/qnamespace.h>
#include "gen_qtextoption.h"

extern "C" {

void* QTextOption_new() {
	return new QTextOption();
}
void* QTextOption_new2(int alignment) {
	return new QTextOption(QFlags<Qt::AlignmentFlag>(alignment));
}
void* QTextOption_new3(void* o) {
	return new QTextOption(*(QTextOption*)o);
}
void* QTextOption_OperatorAssign(void* self, void* o) {
	auto&& __r = ((QTextOption*)self)->operator=(*(QTextOption*)o);
	return (void*)&(__r);
}
void QTextOption_setAlignment(void* self, int alignment) {
	((QTextOption*)self)->setAlignment(QFlags<Qt::AlignmentFlag>(alignment));
}
int QTextOption_alignment(void* self) {
	auto&& __r = ((QTextOption*)self)->alignment();
	return (int)(__r);
}
void QTextOption_setTextDirection(void* self, int aDirection) {
	((QTextOption*)self)->setTextDirection((Qt::LayoutDirection)(aDirection));
}
int QTextOption_textDirection(void* self) {
	auto&& __r = ((QTextOption*)self)->textDirection();
	return (int)(__r);
}
void QTextOption_setWrapMode(void* self, int wrap) {
	((QTextOption*)self)->setWrapMode((QTextOption::WrapMode)(wrap));
}
int QTextOption_wrapMode(void* self) {
	auto&& __r = ((QTextOption*)self)->wrapMode();
	return (int)(__r);
}
void QTextOption_setFlags(void* self, int flags) {
	((QTextOption*)self)->setFlags(QFlags<QTextOption::Flag>(flags));
}
int QTextOption_flags(void* self) {
	auto&& __r = ((QTextOption*)self)->flags();
	return (int)(__r);
}
void QTextOption_setTabStopDistance(void* self, double tabStopDistance) {
	((QTextOption*)self)->setTabStopDistance(tabStopDistance);
}
double QTextOption_tabStopDistance(void* self) {
	auto&& __r = ((QTextOption*)self)->tabStopDistance();
	return __r;
}
void QTextOption_setTabArray(void* self, double* tabStops, long long tabStops_len) {
	QList<double> tabStops_list; tabStops_list.reserve(tabStops_len);
	for (long long __i = 0; __i < tabStops_len; __i++) tabStops_list.append(tabStops[__i]);
	((QTextOption*)self)->setTabArray(tabStops_list);
}
double* QTextOption_tabArray(void* self, long long* __out_len) {
	auto&& __r = ((QTextOption*)self)->tabArray();
	*__out_len = __r.size();
	double* __arr = (double*)malloc(sizeof(double) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = __e;
	return __arr;
}
void QTextOption_setTabs(void* self, void** tabStops, long long tabStops_len) {
	QList<QTextOption::Tab> tabStops_list; tabStops_list.reserve(tabStops_len);
	for (long long __i = 0; __i < tabStops_len; __i++) tabStops_list.append(*(QTextOption::Tab*)tabStops[__i]);
	((QTextOption*)self)->setTabs(tabStops_list);
}
void** QTextOption_tabs(void* self, long long* __out_len) {
	auto&& __r = ((QTextOption*)self)->tabs();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QTextOption::Tab(__e);
	return __arr;
}
void QTextOption_setUseDesignMetrics(void* self, int b) {
	((QTextOption*)self)->setUseDesignMetrics(b);
}
int QTextOption_useDesignMetrics(void* self) {
	auto&& __r = ((QTextOption*)self)->useDesignMetrics();
	return __r;
}
void QTextOption_delete(void* self) { delete (QTextOption*)self; }

}

#include <QtGui/qaccessible_base.h>
#include "gen_qaccessible__activationobserver.h"
#include "gen_purego_cb.h"

class QAccessible__ActivationObserver__Go : public QAccessible::ActivationObserver {
public:
	long long __goId;
	unsigned char __goOn[1] = {};
	template<typename... A> QAccessible__ActivationObserver__Go(long long id, A&&... a) : QAccessible::ActivationObserver(std::forward<A>(a)...), __goId(id) {}
	void accessibilityActiveChanged(bool active) override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 5505922000441301119LL, active, &__h); if (__h) return; }
		return;
	}
};

extern "C" {

void QAccessible__ActivationObserver__set_override(void* self, int idx) { ((QAccessible__ActivationObserver__Go*)self)->__goOn[idx] = 1; }
void QAccessible__ActivationObserver_accessibilityActiveChanged(void* self, int active) {
	((QAccessible::ActivationObserver*)self)->accessibilityActiveChanged(active);
}
void* QAccessible__ActivationObserver_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QAccessible::ActivationObserver*)self)->operator=(*(QAccessible::ActivationObserver*)arg1);
	return (void*)&(__r);
}
void* QAccessible__ActivationObserver__Go_new(long long id) { return new QAccessible__ActivationObserver__Go(id); }
void QAccessible__ActivationObserver_delete(void* self) { delete (QAccessible::ActivationObserver*)self; }

}

#include <QtGui/qaccessibleobject.h>
#include <QtGui/qaccessible.h>
#include <QtCore/qobject.h>
#include <QtCore/qrect.h>
#include <QtGui/qaccessible_base.h>
#include <QtCore/qstring.h>
#include <QtGui/qwindow.h>
#include <QtGui/qcolor.h>
#include "gen_qaccessibleobject.h"

extern "C" {

void* QAccessibleObject_base(void* self) { return (void*)static_cast<QAccessibleInterface*>((QAccessibleObject*)self); }

}

#include <QtGui/qiconengine.h>
#include <QtGui/qpainter.h>
#include <QtCore/qrect.h>
#include <QtGui/qicon.h>
#include <QtCore/qsize.h>
#include <QtGui/qpixmap.h>
#include <QtCore/qstring.h>
#include <QtCore/qdatastream.h>
#include "gen_qiconengine.h"
#include "gen_purego_cb.h"

class QIconEngine__Go : public QIconEngine {
public:
	long long __goId;
	unsigned char __goOn[13] = {};
	template<typename... A> QIconEngine__Go(long long id, A&&... a) : QIconEngine(std::forward<A>(a)...), __goId(id) {}
	void paint(QPainter* painter, const QRect& rect, QIcon::Mode mode, QIcon::State state) override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_void_voidp_voidp_int_int) { ((void(*)(long long, long long, void*, void*, int, int, int*))cppgo_virt_void_voidp_voidp_int_int)(__goId, 6083584261053556819LL, (void*)painter, (void*)&rect, (int)(mode), (int)(state), &__h); if (__h) return; }
		return;
	}
	QIconEngine* clone() const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, -7032813281048616478LL, &__h); if (__h) return (QIconEngine*)__r; }
		return nullptr;
	}
	QSize actualSize(const QSize& size, QIcon::Mode mode, QIcon::State state) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_voidp_voidp_int_int) { void* __r = ((void*(*)(long long, long long, void*, int, int, int*))cppgo_virt_voidp_voidp_int_int)(__goId, 6795227186376262862LL, (void*)&size, (int)(mode), (int)(state), &__h); if (__h) return *(QSize*)__r; }
		return QIconEngine::actualSize(size, mode, state);
	}
	QPixmap pixmap(const QSize& size, QIcon::Mode mode, QIcon::State state) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_voidp_voidp_int_int) { void* __r = ((void*(*)(long long, long long, void*, int, int, int*))cppgo_virt_voidp_voidp_int_int)(__goId, -5258308936133250472LL, (void*)&size, (int)(mode), (int)(state), &__h); if (__h) return *(QPixmap*)__r; }
		return QIconEngine::pixmap(size, mode, state);
	}
	void addPixmap(const QPixmap& pixmap, QIcon::Mode mode, QIcon::State state) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_voidp_int_int) { ((void(*)(long long, long long, void*, int, int, int*))cppgo_virt_void_voidp_int_int)(__goId, -3986256266470558159LL, (void*)&pixmap, (int)(mode), (int)(state), &__h); if (__h) return; }
		QIconEngine::addPixmap(pixmap, mode, state);
	}
	void addFile(const QString& fileName, const QSize& size, QIcon::Mode mode, QIcon::State state) override {
		int __h = 0;
		QByteArray fileName_ba = fileName.toUtf8();
		if (__goOn[5] && cppgo_virt_void_constcharp_longlong_voidp_int_int) { ((void(*)(long long, long long, const char*, long long, void*, int, int, int*))cppgo_virt_void_constcharp_longlong_voidp_int_int)(__goId, -5081758061965960830LL, (char*)fileName_ba.constData(), fileName_ba.size(), (void*)&size, (int)(mode), (int)(state), &__h); if (__h) return; }
		QIconEngine::addFile(fileName, size, mode, state);
	}
	QString key() const override {
		int __h = 0;
		if (__goOn[6] && cppgo_virtp_charp) {
			long long __len = 0;
			char* __s = ((char*(*)(long long, long long, long long*, int*))cppgo_virtp_charp)(__goId, 4452171178779021548LL, &__len, &__h);
			if (__h) { QString __r = QString::fromUtf8(__s, __len); free(__s); return __r; }
		}
		return QIconEngine::key();
	}
	bool read(QDataStream& in) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 5541207787354052101LL, (void*)&in, &__h); if (__h) return __r; }
		return QIconEngine::read(in);
	}
	bool write(QDataStream& out) const override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, -5099743077423337732LL, (void*)&out, &__h); if (__h) return __r; }
		return QIconEngine::write(out);
	}
	QList<QSize> availableSizes(QIcon::Mode mode, QIcon::State state) override {
		int __h = 0;
		if (__goOn[9] && cppgo_virtp_voidp_int_int) {
			long long __len = 0;
			void* __a = ((void*(*)(long long, long long, int, int, long long*, int*))cppgo_virtp_voidp_int_int)(__goId, 7223428027511831406LL, (int)(mode), (int)(state), &__len, &__h);
			if (__h) {
				QList<QSize> __r; __r.reserve(__len);
				void** __aa = (void**)__a;
				for (long long __i = 0; __i < __len; __i++) __r.append(*(QSize*)__aa[__i]);
				if (__a) free(__a);
				return __r;
			}
		}
		return QIconEngine::availableSizes(mode, state);
	}
	QString iconName() override {
		int __h = 0;
		if (__goOn[10] && cppgo_virtp_charp) {
			long long __len = 0;
			char* __s = ((char*(*)(long long, long long, long long*, int*))cppgo_virtp_charp)(__goId, -3060016504017260269LL, &__len, &__h);
			if (__h) { QString __r = QString::fromUtf8(__s, __len); free(__s); return __r; }
		}
		return QIconEngine::iconName();
	}
	bool isNull() override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, 2193725045587114196LL, &__h); if (__h) return __r; }
		return QIconEngine::isNull();
	}
	QPixmap scaledPixmap(const QSize& size, QIcon::Mode mode, QIcon::State state, qreal scale) override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_voidp_voidp_int_int_unsignedlonglong) { void* __r = ((void*(*)(long long, long long, void*, int, int, unsigned long long, int*))cppgo_virt_voidp_voidp_int_int_unsignedlonglong)(__goId, -1544658507073703166LL, (void*)&size, (int)(mode), (int)(state), __builtin_bit_cast(unsigned long long, (double)(scale)), &__h); if (__h) return *(QPixmap*)__r; }
		return QIconEngine::scaledPixmap(size, mode, state, scale);
	}
};

extern "C" {

void QIconEngine__set_override(void* self, int idx) { ((QIconEngine__Go*)self)->__goOn[idx] = 1; }
void QIconEngine_paint(void* self, void* painter, void* rect, int mode, int state) {
	((QIconEngine*)self)->paint((QPainter*)painter, *(QRect*)rect, (QIcon::Mode)(mode), (QIcon::State)(state));
}
void* QIconEngine_actualSize(void* self, void* size, int mode, int state) {
	auto&& __r = ((QIconEngine*)self)->actualSize(*(QSize*)size, (QIcon::Mode)(mode), (QIcon::State)(state));
	return new QSize(__r);
}
void* QIconEngine_pixmap(void* self, void* size, int mode, int state) {
	auto&& __r = ((QIconEngine*)self)->pixmap(*(QSize*)size, (QIcon::Mode)(mode), (QIcon::State)(state));
	return new QPixmap(__r);
}
void QIconEngine_addPixmap(void* self, void* pixmap, int mode, int state) {
	((QIconEngine*)self)->addPixmap(*(QPixmap*)pixmap, (QIcon::Mode)(mode), (QIcon::State)(state));
}
void QIconEngine_addFile(void* self, const char* fileName, long long fileName_len, void* size, int mode, int state) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	((QIconEngine*)self)->addFile(fileName_qstr, *(QSize*)size, (QIcon::Mode)(mode), (QIcon::State)(state));
}
char* QIconEngine_key(void* self, long long* __out_len) {
	auto&& __r = ((QIconEngine*)self)->key();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QIconEngine_clone(void* self) {
	auto&& __r = ((QIconEngine*)self)->clone();
	return (void*)(__r);
}
int QIconEngine_read(void* self, void* in) {
	auto&& __r = ((QIconEngine*)self)->read(*(QDataStream*)in);
	return __r;
}
int QIconEngine_write(void* self, void* out) {
	auto&& __r = ((QIconEngine*)self)->write(*(QDataStream*)out);
	return __r;
}
void** QIconEngine_availableSizes(void* self, int mode, int state, long long* __out_len) {
	auto&& __r = ((QIconEngine*)self)->availableSizes((QIcon::Mode)(mode), (QIcon::State)(state));
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QSize(__e);
	return __arr;
}
char* QIconEngine_iconName(void* self, long long* __out_len) {
	auto&& __r = ((QIconEngine*)self)->iconName();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QIconEngine_isNull(void* self) {
	auto&& __r = ((QIconEngine*)self)->isNull();
	return __r;
}
void* QIconEngine_scaledPixmap(void* self, void* size, int mode, int state, double scale) {
	auto&& __r = ((QIconEngine*)self)->scaledPixmap(*(QSize*)size, (QIcon::Mode)(mode), (QIcon::State)(state), scale);
	return new QPixmap(__r);
}
void* QIconEngine_actualSize_default(void* self, void* size, int mode, int state) {
	auto&& __r = ((QIconEngine*)self)->QIconEngine::actualSize(*(QSize*)size, (QIcon::Mode)(mode), (QIcon::State)(state));
	return new QSize(__r);
}
void* QIconEngine_pixmap_default(void* self, void* size, int mode, int state) {
	auto&& __r = ((QIconEngine*)self)->QIconEngine::pixmap(*(QSize*)size, (QIcon::Mode)(mode), (QIcon::State)(state));
	return new QPixmap(__r);
}
void QIconEngine_addPixmap_default(void* self, void* pixmap, int mode, int state) {
	((QIconEngine*)self)->QIconEngine::addPixmap(*(QPixmap*)pixmap, (QIcon::Mode)(mode), (QIcon::State)(state));
}
void QIconEngine_addFile_default(void* self, const char* fileName, long long fileName_len, void* size, int mode, int state) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	((QIconEngine*)self)->QIconEngine::addFile(fileName_qstr, *(QSize*)size, (QIcon::Mode)(mode), (QIcon::State)(state));
}
char* QIconEngine_key_default(void* self, long long* __out_len) {
	auto&& __r = ((QIconEngine*)self)->QIconEngine::key();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QIconEngine_read_default(void* self, void* in) {
	auto&& __r = ((QIconEngine*)self)->QIconEngine::read(*(QDataStream*)in);
	return __r;
}
int QIconEngine_write_default(void* self, void* out) {
	auto&& __r = ((QIconEngine*)self)->QIconEngine::write(*(QDataStream*)out);
	return __r;
}
void** QIconEngine_availableSizes_default(void* self, int mode, int state, long long* __out_len) {
	auto&& __r = ((QIconEngine*)self)->QIconEngine::availableSizes((QIcon::Mode)(mode), (QIcon::State)(state));
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QSize(__e);
	return __arr;
}
char* QIconEngine_iconName_default(void* self, long long* __out_len) {
	auto&& __r = ((QIconEngine*)self)->QIconEngine::iconName();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QIconEngine_isNull_default(void* self) {
	auto&& __r = ((QIconEngine*)self)->QIconEngine::isNull();
	return __r;
}
void* QIconEngine_scaledPixmap_default(void* self, void* size, int mode, int state, double scale) {
	auto&& __r = ((QIconEngine*)self)->QIconEngine::scaledPixmap(*(QSize*)size, (QIcon::Mode)(mode), (QIcon::State)(state), scale);
	return new QPixmap(__r);
}
void* QIconEngine__Go_new(long long id) { return new QIconEngine__Go(id); }
void QIconEngine_delete(void* self) { delete (QIconEngine*)self; }

}

#include <QtGui/qpaintengine.h>
#include <QtGui/qpen.h>
#include <QtGui/qbrush.h>
#include <QtCore/qpoint.h>
#include <QtCore/qnamespace.h>
#include <QtGui/qfont.h>
#include <QtGui/qtransform.h>
#include <QtGui/qregion.h>
#include <QtGui/qpainterpath.h>
#include <QtGui/qpainter.h>
#include "gen_qpaintenginestate.h"

extern "C" {

int QPaintEngineState_state(void* self) {
	auto&& __r = ((QPaintEngineState*)self)->state();
	return (int)(__r);
}
void* QPaintEngineState_pen(void* self) {
	auto&& __r = ((QPaintEngineState*)self)->pen();
	return new QPen(__r);
}
void* QPaintEngineState_brush(void* self) {
	auto&& __r = ((QPaintEngineState*)self)->brush();
	return new QBrush(__r);
}
void* QPaintEngineState_brushOrigin(void* self) {
	auto&& __r = ((QPaintEngineState*)self)->brushOrigin();
	return new QPointF(__r);
}
void* QPaintEngineState_backgroundBrush(void* self) {
	auto&& __r = ((QPaintEngineState*)self)->backgroundBrush();
	return new QBrush(__r);
}
int QPaintEngineState_backgroundMode(void* self) {
	auto&& __r = ((QPaintEngineState*)self)->backgroundMode();
	return (int)(__r);
}
void* QPaintEngineState_font(void* self) {
	auto&& __r = ((QPaintEngineState*)self)->font();
	return new QFont(__r);
}
void* QPaintEngineState_transform(void* self) {
	auto&& __r = ((QPaintEngineState*)self)->transform();
	return new QTransform(__r);
}
int QPaintEngineState_clipOperation(void* self) {
	auto&& __r = ((QPaintEngineState*)self)->clipOperation();
	return (int)(__r);
}
void* QPaintEngineState_clipRegion(void* self) {
	auto&& __r = ((QPaintEngineState*)self)->clipRegion();
	return new QRegion(__r);
}
void* QPaintEngineState_clipPath(void* self) {
	auto&& __r = ((QPaintEngineState*)self)->clipPath();
	return new QPainterPath(__r);
}
int QPaintEngineState_isClipEnabled(void* self) {
	auto&& __r = ((QPaintEngineState*)self)->isClipEnabled();
	return __r;
}
int QPaintEngineState_renderHints(void* self) {
	auto&& __r = ((QPaintEngineState*)self)->renderHints();
	return (int)(__r);
}
int QPaintEngineState_compositionMode(void* self) {
	auto&& __r = ((QPaintEngineState*)self)->compositionMode();
	return (int)(__r);
}
double QPaintEngineState_opacity(void* self) {
	auto&& __r = ((QPaintEngineState*)self)->opacity();
	return __r;
}
void* QPaintEngineState_painter(void* self) {
	auto&& __r = ((QPaintEngineState*)self)->painter();
	return (void*)(__r);
}
int QPaintEngineState_brushNeedsResolving(void* self) {
	auto&& __r = ((QPaintEngineState*)self)->brushNeedsResolving();
	return __r;
}
int QPaintEngineState_penNeedsResolving(void* self) {
	auto&& __r = ((QPaintEngineState*)self)->penNeedsResolving();
	return __r;
}
void QPaintEngineState_delete(void* self) { delete (QPaintEngineState*)self; }

}

#include <QtGui/qtextobject.h>
#include "gen_qtextblockuserdata.h"

extern "C" {

void* QTextBlockUserData_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QTextBlockUserData*)self)->operator=(*(QTextBlockUserData*)arg1);
	return (void*)&(__r);
}
void QTextBlockUserData_delete(void* self) { delete (QTextBlockUserData*)self; }

}

#include <QtWidgets/qstyle.h>
#include <QtCore/qobject.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtWidgets/qwidget.h>
#include <QtWidgets/qapplication.h>
#include <QtGui/qpalette.h>
#include <QtCore/qrect.h>
#include <QtGui/qfontmetrics.h>
#include <QtGui/qpixmap.h>
#include <QtGui/qpainter.h>
#include <QtWidgets/qstyleoption.h>
#include <QtCore/qpoint.h>
#include <QtCore/qsize.h>
#include <QtGui/qicon.h>
#include <QtCore/qnamespace.h>
#include <QtWidgets/qsizepolicy.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qvariant.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qstyle.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QStyle__Go : public QStyle {
public:
	long long __goId;
	unsigned char __goOn[31] = {};
	template<typename... A> QStyle__Go(long long id, A&&... a) : QStyle(std::forward<A>(a)...), __goId(id) {}
	using QStyle::sender;
	using QStyle::senderSignalIndex;
	using QStyle::receivers;
	using QStyle::isSignalConnected;
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QStyle::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QStyle::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QStyle::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QStyle::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QStyle::disconnectNotify(__a0); }
	void drawPrimitive(QStyle::PrimitiveElement pe, const QStyleOption* opt, QPainter* p, const QWidget* w) const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_void_int_voidp_voidp_voidp) { ((void(*)(long long, long long, int, void*, void*, void*, int*))cppgo_virt_void_int_voidp_voidp_voidp)(__goId, 6098275085281943360LL, (int)(pe), (void*)opt, (void*)p, (void*)w, &__h); if (__h) return; }
		return;
	}
	void drawControl(QStyle::ControlElement element, const QStyleOption* opt, QPainter* p, const QWidget* w) const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_void_int_voidp_voidp_voidp) { ((void(*)(long long, long long, int, void*, void*, void*, int*))cppgo_virt_void_int_voidp_voidp_voidp)(__goId, -8697913287309583788LL, (int)(element), (void*)opt, (void*)p, (void*)w, &__h); if (__h) return; }
		return;
	}
	QRect subElementRect(QStyle::SubElement subElement, const QStyleOption* option, const QWidget* widget) const override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_voidp_int_voidp_voidp) { void* __r = ((void*(*)(long long, long long, int, void*, void*, int*))cppgo_virt_voidp_int_voidp_voidp)(__goId, -1580195784239416471LL, (int)(subElement), (void*)option, (void*)widget, &__h); if (__h) return *(QRect*)__r; }
		return QRect{};
	}
	void drawComplexControl(QStyle::ComplexControl cc, const QStyleOptionComplex* opt, QPainter* p, const QWidget* widget) const override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_int_voidp_voidp_voidp) { ((void(*)(long long, long long, int, void*, void*, void*, int*))cppgo_virt_void_int_voidp_voidp_voidp)(__goId, -778167327566211030LL, (int)(cc), (void*)opt, (void*)p, (void*)widget, &__h); if (__h) return; }
		return;
	}
	QStyle::SubControl hitTestComplexControl(QStyle::ComplexControl cc, const QStyleOptionComplex* opt, const QPoint& pt, const QWidget* widget) const override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_int_int_voidp_voidp_voidp) { int __r = ((int(*)(long long, long long, int, void*, void*, void*, int*))cppgo_virt_int_int_voidp_voidp_voidp)(__goId, 7058885155526130727LL, (int)(cc), (void*)opt, (void*)&pt, (void*)widget, &__h); if (__h) return (QStyle::SubControl)(__r); }
		return QStyle::SubControl{};
	}
	QRect subControlRect(QStyle::ComplexControl cc, const QStyleOptionComplex* opt, QStyle::SubControl sc, const QWidget* widget) const override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_voidp_int_voidp_int_voidp) { void* __r = ((void*(*)(long long, long long, int, void*, int, void*, int*))cppgo_virt_voidp_int_voidp_int_voidp)(__goId, -8244607819117461060LL, (int)(cc), (void*)opt, (int)(sc), (void*)widget, &__h); if (__h) return *(QRect*)__r; }
		return QRect{};
	}
	int pixelMetric(QStyle::PixelMetric metric, const QStyleOption* option, const QWidget* widget) const override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_int_int_voidp_voidp) { int __r = ((int(*)(long long, long long, int, void*, void*, int*))cppgo_virt_int_int_voidp_voidp)(__goId, -7621686327863150821LL, (int)(metric), (void*)option, (void*)widget, &__h); if (__h) return __r; }
		return int{};
	}
	QSize sizeFromContents(QStyle::ContentsType ct, const QStyleOption* opt, const QSize& contentsSize, const QWidget* w) const override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_voidp_int_voidp_voidp_voidp) { void* __r = ((void*(*)(long long, long long, int, void*, void*, void*, int*))cppgo_virt_voidp_int_voidp_voidp_voidp)(__goId, -3316981301817172730LL, (int)(ct), (void*)opt, (void*)&contentsSize, (void*)w, &__h); if (__h) return *(QSize*)__r; }
		return QSize{};
	}
	int styleHint(QStyle::StyleHint stylehint, const QStyleOption* opt, const QWidget* widget, QStyleHintReturn* returnData) const override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_int_int_voidp_voidp_voidp) { int __r = ((int(*)(long long, long long, int, void*, void*, void*, int*))cppgo_virt_int_int_voidp_voidp_voidp)(__goId, 2017029407462380579LL, (int)(stylehint), (void*)opt, (void*)widget, (void*)returnData, &__h); if (__h) return __r; }
		return int{};
	}
	QPixmap standardPixmap(QStyle::StandardPixmap standardPixmap, const QStyleOption* opt, const QWidget* widget) const override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_voidp_int_voidp_voidp) { void* __r = ((void*(*)(long long, long long, int, void*, void*, int*))cppgo_virt_voidp_int_voidp_voidp)(__goId, -5288619860977043091LL, (int)(standardPixmap), (void*)opt, (void*)widget, &__h); if (__h) return *(QPixmap*)__r; }
		return QPixmap{};
	}
	QIcon standardIcon(QStyle::StandardPixmap standardIcon, const QStyleOption* option, const QWidget* widget) const override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_voidp_int_voidp_voidp) { void* __r = ((void*(*)(long long, long long, int, void*, void*, int*))cppgo_virt_voidp_int_voidp_voidp)(__goId, -1870022100809264807LL, (int)(standardIcon), (void*)option, (void*)widget, &__h); if (__h) return *(QIcon*)__r; }
		return QIcon{};
	}
	QPixmap generatedIconPixmap(QIcon::Mode iconMode, const QPixmap& pixmap, const QStyleOption* opt) const override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_voidp_int_voidp_voidp) { void* __r = ((void*(*)(long long, long long, int, void*, void*, int*))cppgo_virt_voidp_int_voidp_voidp)(__goId, -4387834948099211678LL, (int)(iconMode), (void*)&pixmap, (void*)opt, &__h); if (__h) return *(QPixmap*)__r; }
		return QPixmap{};
	}
	int layoutSpacing(QSizePolicy::ControlType control1, QSizePolicy::ControlType control2, Qt::Orientation orientation, const QStyleOption* option, const QWidget* widget) const override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_int_int_int_int_voidp_voidp) { int __r = ((int(*)(long long, long long, int, int, int, void*, void*, int*))cppgo_virt_int_int_int_int_voidp_voidp)(__goId, -4874522738816254324LL, (int)(control1), (int)(control2), (int)(orientation), (void*)option, (void*)widget, &__h); if (__h) return __r; }
		return int{};
	}
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QStyle::metaObject();
	}
	void polish(QWidget* widget) override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4633126903937721620LL, (void*)widget, &__h); if (__h) return; }
		QStyle::polish(widget);
	}
	void unpolish(QWidget* widget) override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5547720835653541585LL, (void*)widget, &__h); if (__h) return; }
		QStyle::unpolish(widget);
	}
	void polish(QApplication* application) override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4263633720130086714LL, (void*)application, &__h); if (__h) return; }
		QStyle::polish(application);
	}
	void unpolish(QApplication* application) override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 265826155645830215LL, (void*)application, &__h); if (__h) return; }
		QStyle::unpolish(application);
	}
	void polish(QPalette& palette) override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4263634819641714925LL, (void*)&palette, &__h); if (__h) return; }
		QStyle::polish(palette);
	}
	QRect itemTextRect(const QFontMetrics& fm, const QRect& r, int flags, bool enabled, const QString& text) const override {
		int __h = 0;
		QByteArray text_ba = text.toUtf8();
		if (__goOn[19] && cppgo_virt_voidp_voidp_voidp_int_int_constcharp_longlong) { void* __r = ((void*(*)(long long, long long, void*, void*, int, int, const char*, long long, int*))cppgo_virt_voidp_voidp_voidp_int_int_constcharp_longlong)(__goId, 8974830244642632823LL, (void*)&fm, (void*)&r, flags, enabled, (char*)text_ba.constData(), text_ba.size(), &__h); if (__h) return *(QRect*)__r; }
		return QStyle::itemTextRect(fm, r, flags, enabled, text);
	}
	QRect itemPixmapRect(const QRect& r, int flags, const QPixmap& pixmap) const override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_voidp_voidp_int_voidp) { void* __r = ((void*(*)(long long, long long, void*, int, void*, int*))cppgo_virt_voidp_voidp_int_voidp)(__goId, 3847657773274941773LL, (void*)&r, flags, (void*)&pixmap, &__h); if (__h) return *(QRect*)__r; }
		return QStyle::itemPixmapRect(r, flags, pixmap);
	}
	void drawItemText(QPainter* painter, const QRect& rect, int flags, const QPalette& pal, bool enabled, const QString& text, QPalette::ColorRole textRole) const override {
		int __h = 0;
		QByteArray text_ba = text.toUtf8();
		if (__goOn[21] && cppgo_virt_void_voidp_voidp_int_voidp_int_constcharp_longlong_int) { ((void(*)(long long, long long, void*, void*, int, void*, int, const char*, long long, int, int*))cppgo_virt_void_voidp_voidp_int_voidp_int_constcharp_longlong_int)(__goId, 2315857538243927747LL, (void*)painter, (void*)&rect, flags, (void*)&pal, enabled, (char*)text_ba.constData(), text_ba.size(), (int)(textRole), &__h); if (__h) return; }
		QStyle::drawItemText(painter, rect, flags, pal, enabled, text, textRole);
	}
	void drawItemPixmap(QPainter* painter, const QRect& rect, int alignment, const QPixmap& pixmap) const override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_void_voidp_voidp_int_voidp) { ((void(*)(long long, long long, void*, void*, int, void*, int*))cppgo_virt_void_voidp_voidp_int_voidp)(__goId, -3740415202673381007LL, (void*)painter, (void*)&rect, alignment, (void*)&pixmap, &__h); if (__h) return; }
		QStyle::drawItemPixmap(painter, rect, alignment, pixmap);
	}
	QPalette standardPalette() const override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 7278964765692882119LL, &__h); if (__h) return *(QPalette*)__r; }
		return QStyle::standardPalette();
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QStyle::event(event);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QStyle::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QStyle::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QStyle::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QStyle::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QStyle::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QStyle::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QStyle;
static std::mutex __connsMu_QStyle;

extern "C" {

void* QStyle_base(void* self) { return (void*)static_cast<QObject*>((QStyle*)self); }
void QStyle__set_override(void* self, int idx) { ((QStyle__Go*)self)->__goOn[idx] = 1; }
void* QStyle_fromQObject(void* o) { return (void*)qobject_cast<QStyle*>((QObject*)o); }
char* QStyle_tr(const char* s, long long* __out_len) {
	auto&& __r = QStyle::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QStyle_name(void* self, long long* __out_len) {
	auto&& __r = ((QStyle*)self)->name();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QStyle_polish(void* self, void* widget) {
	((QStyle*)self)->polish((QWidget*)widget);
}
void QStyle_unpolish(void* self, void* widget) {
	((QStyle*)self)->unpolish((QWidget*)widget);
}
void QStyle_polish2(void* self, void* application) {
	((QStyle*)self)->polish((QApplication*)application);
}
void QStyle_unpolish2(void* self, void* application) {
	((QStyle*)self)->unpolish((QApplication*)application);
}
void QStyle_polish3(void* self, void* palette) {
	((QStyle*)self)->polish(*(QPalette*)palette);
}
void* QStyle_itemTextRect(void* self, void* fm, void* r, int flags, int enabled, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = ((QStyle*)self)->itemTextRect(*(QFontMetrics*)fm, *(QRect*)r, flags, enabled, text_qstr);
	return new QRect(__r);
}
void* QStyle_itemPixmapRect(void* self, void* r, int flags, void* pixmap) {
	auto&& __r = ((QStyle*)self)->itemPixmapRect(*(QRect*)r, flags, *(QPixmap*)pixmap);
	return new QRect(__r);
}
void QStyle_drawItemText(void* self, void* painter, void* rect, int flags, void* pal, int enabled, const char* text, long long text_len, int textRole) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	((QStyle*)self)->drawItemText((QPainter*)painter, *(QRect*)rect, flags, *(QPalette*)pal, enabled, text_qstr, (QPalette::ColorRole)(textRole));
}
void QStyle_drawItemPixmap(void* self, void* painter, void* rect, int alignment, void* pixmap) {
	((QStyle*)self)->drawItemPixmap((QPainter*)painter, *(QRect*)rect, alignment, *(QPixmap*)pixmap);
}
void* QStyle_standardPalette(void* self) {
	auto&& __r = ((QStyle*)self)->standardPalette();
	return new QPalette(__r);
}
void QStyle_drawPrimitive(void* self, int pe, void* opt, void* p, void* w) {
	((QStyle*)self)->drawPrimitive((QStyle::PrimitiveElement)(pe), (QStyleOption*)opt, (QPainter*)p, (QWidget*)w);
}
void QStyle_drawControl(void* self, int element, void* opt, void* p, void* w) {
	((QStyle*)self)->drawControl((QStyle::ControlElement)(element), (QStyleOption*)opt, (QPainter*)p, (QWidget*)w);
}
void* QStyle_subElementRect(void* self, int subElement, void* option, void* widget) {
	auto&& __r = ((QStyle*)self)->subElementRect((QStyle::SubElement)(subElement), (QStyleOption*)option, (QWidget*)widget);
	return new QRect(__r);
}
void QStyle_drawComplexControl(void* self, int cc, void* opt, void* p, void* widget) {
	((QStyle*)self)->drawComplexControl((QStyle::ComplexControl)(cc), (QStyleOptionComplex*)opt, (QPainter*)p, (QWidget*)widget);
}
int QStyle_hitTestComplexControl(void* self, int cc, void* opt, void* pt, void* widget) {
	auto&& __r = ((QStyle*)self)->hitTestComplexControl((QStyle::ComplexControl)(cc), (QStyleOptionComplex*)opt, *(QPoint*)pt, (QWidget*)widget);
	return (int)(__r);
}
void* QStyle_subControlRect(void* self, int cc, void* opt, int sc, void* widget) {
	auto&& __r = ((QStyle*)self)->subControlRect((QStyle::ComplexControl)(cc), (QStyleOptionComplex*)opt, (QStyle::SubControl)(sc), (QWidget*)widget);
	return new QRect(__r);
}
int QStyle_pixelMetric(void* self, int metric, void* option, void* widget) {
	auto&& __r = ((QStyle*)self)->pixelMetric((QStyle::PixelMetric)(metric), (QStyleOption*)option, (QWidget*)widget);
	return __r;
}
void* QStyle_sizeFromContents(void* self, int ct, void* opt, void* contentsSize, void* w) {
	auto&& __r = ((QStyle*)self)->sizeFromContents((QStyle::ContentsType)(ct), (QStyleOption*)opt, *(QSize*)contentsSize, (QWidget*)w);
	return new QSize(__r);
}
int QStyle_styleHint(void* self, int stylehint, void* opt, void* widget, void* returnData) {
	auto&& __r = ((QStyle*)self)->styleHint((QStyle::StyleHint)(stylehint), (QStyleOption*)opt, (QWidget*)widget, (QStyleHintReturn*)returnData);
	return __r;
}
void* QStyle_standardPixmap(void* self, int standardPixmap, void* opt, void* widget) {
	auto&& __r = ((QStyle*)self)->standardPixmap((QStyle::StandardPixmap)(standardPixmap), (QStyleOption*)opt, (QWidget*)widget);
	return new QPixmap(__r);
}
void* QStyle_standardIcon(void* self, int standardIcon, void* option, void* widget) {
	auto&& __r = ((QStyle*)self)->standardIcon((QStyle::StandardPixmap)(standardIcon), (QStyleOption*)option, (QWidget*)widget);
	return new QIcon(__r);
}
void* QStyle_generatedIconPixmap(void* self, int iconMode, void* pixmap, void* opt) {
	auto&& __r = ((QStyle*)self)->generatedIconPixmap((QIcon::Mode)(iconMode), *(QPixmap*)pixmap, (QStyleOption*)opt);
	return new QPixmap(__r);
}
void* QStyle_visualRect(int direction, void* boundingRect, void* logicalRect) {
	auto&& __r = QStyle::visualRect((Qt::LayoutDirection)(direction), *(QRect*)boundingRect, *(QRect*)logicalRect);
	return new QRect(__r);
}
void* QStyle_visualPos(int direction, void* boundingRect, void* logicalPos) {
	auto&& __r = QStyle::visualPos((Qt::LayoutDirection)(direction), *(QRect*)boundingRect, *(QPoint*)logicalPos);
	return new QPoint(__r);
}
int QStyle_sliderPositionFromValue(int min, int max, int val, int space) {
	auto&& __r = QStyle::sliderPositionFromValue(min, max, val, space);
	return __r;
}
int QStyle_sliderValueFromPosition(int min, int max, int pos, int space) {
	auto&& __r = QStyle::sliderValueFromPosition(min, max, pos, space);
	return __r;
}
int QStyle_visualAlignment(int direction, int alignment) {
	auto&& __r = QStyle::visualAlignment((Qt::LayoutDirection)(direction), QFlags<Qt::AlignmentFlag>(alignment));
	return (int)(__r);
}
void* QStyle_alignedRect(int direction, int alignment, void* size, void* rectangle) {
	auto&& __r = QStyle::alignedRect((Qt::LayoutDirection)(direction), QFlags<Qt::AlignmentFlag>(alignment), *(QSize*)size, *(QRect*)rectangle);
	return new QRect(__r);
}
int QStyle_layoutSpacing(void* self, int control1, int control2, int orientation, void* option, void* widget) {
	auto&& __r = ((QStyle*)self)->layoutSpacing((QSizePolicy::ControlType)(control1), (QSizePolicy::ControlType)(control2), (Qt::Orientation)(orientation), (QStyleOption*)option, (QWidget*)widget);
	return __r;
}
int QStyle_combinedLayoutSpacing(void* self, int controls1, int controls2, int orientation) {
	auto&& __r = ((QStyle*)self)->combinedLayoutSpacing(QFlags<QSizePolicy::ControlType>(controls1), QFlags<QSizePolicy::ControlType>(controls2), (Qt::Orientation)(orientation));
	return __r;
}
void* QStyle_proxy(void* self) {
	auto&& __r = ((QStyle*)self)->proxy();
	return (void*)(__r);
}
void* QStyle_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QStyle::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QStyle_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QStyle::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QStyle_sender(void* self) {
	auto&& __r = ((QStyle__Go*)self)->sender();
	return (void*)(__r);
}
int QStyle_senderSignalIndex(void* self) {
	auto&& __r = ((QStyle__Go*)self)->senderSignalIndex();
	return __r;
}
int QStyle_receivers(void* self, const char* signal) {
	auto&& __r = ((QStyle__Go*)self)->receivers(signal);
	return __r;
}
int QStyle_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QStyle__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QStyle_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QStyle::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QStyle_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QStyle::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QStyle_sliderPositionFromValue2(int min, int max, int val, int space, int upsideDown) {
	auto&& __r = QStyle::sliderPositionFromValue(min, max, val, space, upsideDown);
	return __r;
}
int QStyle_sliderValueFromPosition2(int min, int max, int pos, int space, int upsideDown) {
	auto&& __r = QStyle::sliderValueFromPosition(min, max, pos, space, upsideDown);
	return __r;
}
int QStyle_combinedLayoutSpacing2(void* self, int controls1, int controls2, int orientation, void* option) {
	auto&& __r = ((QStyle*)self)->combinedLayoutSpacing(QFlags<QSizePolicy::ControlType>(controls1), QFlags<QSizePolicy::ControlType>(controls2), (Qt::Orientation)(orientation), (QStyleOption*)option);
	return __r;
}
int QStyle_combinedLayoutSpacing3(void* self, int controls1, int controls2, int orientation, void* option, void* widget) {
	auto&& __r = ((QStyle*)self)->combinedLayoutSpacing(QFlags<QSizePolicy::ControlType>(controls1), QFlags<QSizePolicy::ControlType>(controls2), (Qt::Orientation)(orientation), (QStyleOption*)option, (QWidget*)widget);
	return __r;
}
void* QStyle_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QStyle::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QStyle_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QStyle::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void QStyle_polish_default(void* self, void* widget) {
	((QStyle*)self)->QStyle::polish((QWidget*)widget);
}
void QStyle_unpolish_default(void* self, void* widget) {
	((QStyle*)self)->QStyle::unpolish((QWidget*)widget);
}
void QStyle_polish_default2(void* self, void* application) {
	((QStyle*)self)->QStyle::polish((QApplication*)application);
}
void QStyle_unpolish_default2(void* self, void* application) {
	((QStyle*)self)->QStyle::unpolish((QApplication*)application);
}
void QStyle_polish_default3(void* self, void* palette) {
	((QStyle*)self)->QStyle::polish(*(QPalette*)palette);
}
void* QStyle_itemTextRect_default(void* self, void* fm, void* r, int flags, int enabled, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = ((QStyle*)self)->QStyle::itemTextRect(*(QFontMetrics*)fm, *(QRect*)r, flags, enabled, text_qstr);
	return new QRect(__r);
}
void* QStyle_itemPixmapRect_default(void* self, void* r, int flags, void* pixmap) {
	auto&& __r = ((QStyle*)self)->QStyle::itemPixmapRect(*(QRect*)r, flags, *(QPixmap*)pixmap);
	return new QRect(__r);
}
void QStyle_drawItemText_default(void* self, void* painter, void* rect, int flags, void* pal, int enabled, const char* text, long long text_len, int textRole) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	((QStyle*)self)->QStyle::drawItemText((QPainter*)painter, *(QRect*)rect, flags, *(QPalette*)pal, enabled, text_qstr, (QPalette::ColorRole)(textRole));
}
void QStyle_drawItemPixmap_default(void* self, void* painter, void* rect, int alignment, void* pixmap) {
	((QStyle*)self)->QStyle::drawItemPixmap((QPainter*)painter, *(QRect*)rect, alignment, *(QPixmap*)pixmap);
}
void* QStyle_standardPalette_default(void* self) {
	auto&& __r = ((QStyle*)self)->QStyle::standardPalette();
	return new QPalette(__r);
}
void* QStyle_metaObject_default(void* self) {
	auto&& __r = ((QStyle*)self)->QStyle::metaObject();
	return (void*)(__r);
}
int QStyle_event_default(void* self, void* event) {
	auto&& __r = ((QStyle*)self)->QStyle::event((QEvent*)event);
	return __r;
}
int QStyle_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QStyle*)self)->QStyle::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QStyle_timerEvent_default(void* self, void* event) {
	((QStyle__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QStyle_childEvent_default(void* self, void* event) {
	((QStyle__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QStyle_customEvent_default(void* self, void* event) {
	((QStyle__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QStyle_connectNotify_default(void* self, void* signal) {
	((QStyle__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QStyle_disconnectNotify_default(void* self, void* signal) {
	((QStyle__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QStyle_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStyle*)self, &QStyle::destroyed, (QStyle*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStyle);
	__conns_QStyle[__id] = __c;
}
void QStyle_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStyle*)self, &QStyle::destroyed, (QStyle*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStyle);
	__conns_QStyle[__id] = __c;
}
void QStyle_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStyle*)self, &QStyle::objectNameChanged, (QStyle*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStyle);
	__conns_QStyle[__id] = __c;
}
void QStyle__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QStyle);
		auto __it = __conns_QStyle.find(id);
		if (__it == __conns_QStyle.end()) return;
		__c = __it->second;
		__conns_QStyle.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QStyle__Go_new(long long id) { return new QStyle__Go(id); }
void QStyle_delete(void* self) { delete (QStyle*)self; }

}

#include <QtWidgets/qstyleoption.h>
#include <QtWidgets/qwidget.h>
#include <QtWidgets/qabstractspinbox.h>
#include "gen_qstyleoptionspinbox.h"

extern "C" {

void* QStyleOptionSpinBox_base(void* self) { return (void*)static_cast<QStyleOptionComplex*>((QStyleOptionSpinBox*)self); }
void QStyleOptionSpinBox_chain(void* self, void** out) {
	QStyleOptionSpinBox* __p = (QStyleOptionSpinBox*)self;
	QStyleOptionComplex* __b0 = static_cast<QStyleOptionComplex*>(__p); out[0] = (void*)__b0;
	QStyleOption* __b1 = static_cast<QStyleOption*>(__b0); out[1] = (void*)__b1;
}
void* QStyleOptionSpinBox_new() {
	return new QStyleOptionSpinBox();
}
void* QStyleOptionSpinBox_new2(void* other) {
	return new QStyleOptionSpinBox(*(QStyleOptionSpinBox*)other);
}
void* QStyleOptionSpinBox_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QStyleOptionSpinBox*)self)->operator=(*(QStyleOptionSpinBox*)arg1);
	return (void*)&(__r);
}
int QStyleOptionSpinBox_buttonSymbols_get(void* self) {
	auto __r = ((QStyleOptionSpinBox*)self)->buttonSymbols;
	return (int)(__r);
}
void QStyleOptionSpinBox_buttonSymbols_set(void* self, int buttonSymbols) {
	((QStyleOptionSpinBox*)self)->buttonSymbols = (QAbstractSpinBox::ButtonSymbols)(buttonSymbols);
}
int QStyleOptionSpinBox_stepEnabled_get(void* self) {
	auto __r = ((QStyleOptionSpinBox*)self)->stepEnabled;
	return (int)(__r);
}
void QStyleOptionSpinBox_stepEnabled_set(void* self, int stepEnabled) {
	((QStyleOptionSpinBox*)self)->stepEnabled = QFlags<QAbstractSpinBox::StepEnabledFlag>(stepEnabled);
}
int QStyleOptionSpinBox_frame_get(void* self) {
	auto __r = ((QStyleOptionSpinBox*)self)->frame;
	return __r;
}
void QStyleOptionSpinBox_frame_set(void* self, int frame) {
	((QStyleOptionSpinBox*)self)->frame = frame;
}
void QStyleOptionSpinBox_delete(void* self) { delete (QStyleOptionSpinBox*)self; }

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
#include "gen_qvboxlayout.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QVBoxLayout__Go : public QVBoxLayout {
public:
	long long __goId;
	unsigned char __goOn[32] = {};
	template<typename... A> QVBoxLayout__Go(long long id, A&&... a) : QVBoxLayout(std::forward<A>(a)...), __goId(id) {}
	using QVBoxLayout::widgetEvent;
	using QVBoxLayout::addChildLayout;
	using QVBoxLayout::addChildWidget;
	using QVBoxLayout::adoptLayout;
	using QVBoxLayout::alignmentRect;
	using QVBoxLayout::sender;
	using QVBoxLayout::senderSignalIndex;
	using QVBoxLayout::receivers;
	using QVBoxLayout::isSignalConnected;
	void cppgoBase_childEvent(QChildEvent* __a0) { QVBoxLayout::childEvent(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QVBoxLayout::timerEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QVBoxLayout::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QVBoxLayout::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QVBoxLayout::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QVBoxLayout::metaObject();
	}
	void addItem(QLayoutItem* arg1) override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6884465960736393867LL, (void*)arg1, &__h); if (__h) return; }
		QVBoxLayout::addItem(arg1);
	}
	int spacing() const override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -4587105490415799682LL, &__h); if (__h) return __r; }
		return QVBoxLayout::spacing();
	}
	void setSpacing(int spacing) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 4039479028406881496LL, spacing, &__h); if (__h) return; }
		QVBoxLayout::setSpacing(spacing);
	}
	QSize sizeHint() const override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 729690261457684101LL, &__h); if (__h) return *(QSize*)__r; }
		return QVBoxLayout::sizeHint();
	}
	QSize minimumSize() const override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, -7246748916541192690LL, &__h); if (__h) return *(QSize*)__r; }
		return QVBoxLayout::minimumSize();
	}
	QSize maximumSize() const override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 4158099606088301136LL, &__h); if (__h) return *(QSize*)__r; }
		return QVBoxLayout::maximumSize();
	}
	bool hasHeightForWidth() const override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8266762397995423915LL, &__h); if (__h) return __r; }
		return QVBoxLayout::hasHeightForWidth();
	}
	int heightForWidth(int arg1) const override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 4311534060495737339LL, arg1, &__h); if (__h) return __r; }
		return QVBoxLayout::heightForWidth(arg1);
	}
	int minimumHeightForWidth(int arg1) const override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 5037837229528203129LL, arg1, &__h); if (__h) return __r; }
		return QVBoxLayout::minimumHeightForWidth(arg1);
	}
	QFlags<Qt::Orientation> expandingDirections() const override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, 415665469015095569LL, &__h); if (__h) return QFlags<Qt::Orientation>(__r); }
		return QVBoxLayout::expandingDirections();
	}
	void invalidate() override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 6991341317990632368LL, &__h); if (__h) return; }
		QVBoxLayout::invalidate();
	}
	QLayoutItem* itemAt(int arg1) const override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7508512214013391107LL, arg1, &__h); if (__h) return (QLayoutItem*)__r; }
		return QVBoxLayout::itemAt(arg1);
	}
	QLayoutItem* takeAt(int arg1) override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7549889108533486643LL, arg1, &__h); if (__h) return (QLayoutItem*)__r; }
		return QVBoxLayout::takeAt(arg1);
	}
	int count() const override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -5627843058667511180LL, &__h); if (__h) return __r; }
		return QVBoxLayout::count();
	}
	void setGeometry(const QRect& arg1) override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5684820848726969227LL, (void*)&arg1, &__h); if (__h) return; }
		QVBoxLayout::setGeometry(arg1);
	}
	QRect geometry() const override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1050490530495513297LL, &__h); if (__h) return *(QRect*)__r; }
		return QVBoxLayout::geometry();
	}
	int indexOf(const QWidget* arg1) const override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, -7258473138030329570LL, (void*)arg1, &__h); if (__h) return __r; }
		return QVBoxLayout::indexOf(arg1);
	}
	int indexOf(const QLayoutItem* arg1) const override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, -3276435765994963516LL, (void*)arg1, &__h); if (__h) return __r; }
		return QVBoxLayout::indexOf(arg1);
	}
	bool isEmpty() const override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -626822430251098498LL, &__h); if (__h) return __r; }
		return QVBoxLayout::isEmpty();
	}
	QFlags<QSizePolicy::ControlType> controlTypes() const override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, 3278947275714896887LL, &__h); if (__h) return QFlags<QSizePolicy::ControlType>(__r); }
		return QVBoxLayout::controlTypes();
	}
	QLayoutItem* replaceWidget(QWidget* from, QWidget* to, QFlags<Qt::FindChildOption> options) override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_voidp_voidp_voidp_int) { void* __r = ((void*(*)(long long, long long, void*, void*, int, int*))cppgo_virt_voidp_voidp_voidp_int)(__goId, -712862742159708061LL, (void*)from, (void*)to, (int)(options), &__h); if (__h) return (QLayoutItem*)__r; }
		return QVBoxLayout::replaceWidget(from, to, options);
	}
	QLayout* layout() override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, -7183390618627992737LL, &__h); if (__h) return (QLayout*)__r; }
		return QVBoxLayout::layout();
	}
	void childEvent(QChildEvent* e) override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)e, &__h); if (__h) return; }
		QVBoxLayout::childEvent(e);
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QVBoxLayout::event(event);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QVBoxLayout::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QVBoxLayout::timerEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QVBoxLayout::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QVBoxLayout::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QVBoxLayout::disconnectNotify(signal);
	}
	QWidget* widget() const override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, -4575937485689502171LL, &__h); if (__h) return (QWidget*)__r; }
		return QVBoxLayout::widget();
	}
	QSpacerItem* spacerItem() override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 2272240669359776278LL, &__h); if (__h) return (QSpacerItem*)__r; }
		return QVBoxLayout::spacerItem();
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QVBoxLayout;
static std::mutex __connsMu_QVBoxLayout;

extern "C" {

void* QVBoxLayout_base(void* self) { return (void*)static_cast<QBoxLayout*>((QVBoxLayout*)self); }
void QVBoxLayout_chain(void* self, void** out) {
	QVBoxLayout* __p = (QVBoxLayout*)self;
	QBoxLayout* __b0 = static_cast<QBoxLayout*>(__p); out[0] = (void*)__b0;
	QLayout* __b1 = static_cast<QLayout*>(__b0); out[1] = (void*)__b1;
	QObject* __b2 = static_cast<QObject*>(__b1); out[2] = (void*)__b2;
}
void QVBoxLayout__set_override(void* self, int idx) { ((QVBoxLayout__Go*)self)->__goOn[idx] = 1; }
void* QVBoxLayout_fromQObject(void* o) { return (void*)qobject_cast<QVBoxLayout*>((QObject*)o); }
void* QVBoxLayout_new(long long __goid, void* parent) {
	return new QVBoxLayout__Go(__goid, (QWidget*)parent);
}
void* QVBoxLayout_new2(long long __goid) {
	return new QVBoxLayout__Go(__goid);
}
char* QVBoxLayout_tr(const char* s, long long* __out_len) {
	auto&& __r = QVBoxLayout::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QVBoxLayout_closestAcceptableSize(void* w, void* s) {
	auto&& __r = QVBoxLayout::closestAcceptableSize((QWidget*)w, *(QSize*)s);
	return new QSize(__r);
}
void QVBoxLayout_widgetEvent(void* self, void* arg1) {
	((QVBoxLayout__Go*)self)->widgetEvent((QEvent*)arg1);
}
void QVBoxLayout_addChildLayout(void* self, void* l) {
	((QVBoxLayout__Go*)self)->addChildLayout((QLayout*)l);
}
void QVBoxLayout_addChildWidget(void* self, void* w) {
	((QVBoxLayout__Go*)self)->addChildWidget((QWidget*)w);
}
int QVBoxLayout_adoptLayout(void* self, void* layout) {
	auto&& __r = ((QVBoxLayout__Go*)self)->adoptLayout((QLayout*)layout);
	return __r;
}
void* QVBoxLayout_alignmentRect(void* self, void* arg1) {
	auto&& __r = ((QVBoxLayout__Go*)self)->alignmentRect(*(QRect*)arg1);
	return new QRect(__r);
}
void* QVBoxLayout_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QVBoxLayout::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QVBoxLayout_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QVBoxLayout::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QVBoxLayout_sender(void* self) {
	auto&& __r = ((QVBoxLayout__Go*)self)->sender();
	return (void*)(__r);
}
int QVBoxLayout_senderSignalIndex(void* self) {
	auto&& __r = ((QVBoxLayout__Go*)self)->senderSignalIndex();
	return __r;
}
int QVBoxLayout_receivers(void* self, const char* signal) {
	auto&& __r = ((QVBoxLayout__Go*)self)->receivers(signal);
	return __r;
}
int QVBoxLayout_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QVBoxLayout__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QVBoxLayout_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QVBoxLayout::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QVBoxLayout_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QVBoxLayout::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QVBoxLayout_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QVBoxLayout::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QVBoxLayout_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QVBoxLayout::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QVBoxLayout_metaObject_default(void* self) {
	auto&& __r = ((QVBoxLayout*)self)->QVBoxLayout::metaObject();
	return (void*)(__r);
}
void QVBoxLayout_addItem_default(void* self, void* arg1) {
	((QVBoxLayout*)self)->QVBoxLayout::addItem((QLayoutItem*)arg1);
}
int QVBoxLayout_spacing_default(void* self) {
	auto&& __r = ((QVBoxLayout*)self)->QVBoxLayout::spacing();
	return __r;
}
void QVBoxLayout_setSpacing_default(void* self, int spacing) {
	((QVBoxLayout*)self)->QVBoxLayout::setSpacing(spacing);
}
void* QVBoxLayout_sizeHint_default(void* self) {
	auto&& __r = ((QVBoxLayout*)self)->QVBoxLayout::sizeHint();
	return new QSize(__r);
}
void* QVBoxLayout_minimumSize_default(void* self) {
	auto&& __r = ((QVBoxLayout*)self)->QVBoxLayout::minimumSize();
	return new QSize(__r);
}
void* QVBoxLayout_maximumSize_default(void* self) {
	auto&& __r = ((QVBoxLayout*)self)->QVBoxLayout::maximumSize();
	return new QSize(__r);
}
int QVBoxLayout_hasHeightForWidth_default(void* self) {
	auto&& __r = ((QVBoxLayout*)self)->QVBoxLayout::hasHeightForWidth();
	return __r;
}
int QVBoxLayout_heightForWidth_default(void* self, int arg1) {
	auto&& __r = ((QVBoxLayout*)self)->QVBoxLayout::heightForWidth(arg1);
	return __r;
}
int QVBoxLayout_minimumHeightForWidth_default(void* self, int arg1) {
	auto&& __r = ((QVBoxLayout*)self)->QVBoxLayout::minimumHeightForWidth(arg1);
	return __r;
}
int QVBoxLayout_expandingDirections_default(void* self) {
	auto&& __r = ((QVBoxLayout*)self)->QVBoxLayout::expandingDirections();
	return (int)(__r);
}
void QVBoxLayout_invalidate_default(void* self) {
	((QVBoxLayout*)self)->QVBoxLayout::invalidate();
}
void* QVBoxLayout_itemAt_default(void* self, int arg1) {
	auto&& __r = ((QVBoxLayout*)self)->QVBoxLayout::itemAt(arg1);
	return (void*)(__r);
}
void* QVBoxLayout_takeAt_default(void* self, int arg1) {
	auto&& __r = ((QVBoxLayout*)self)->QVBoxLayout::takeAt(arg1);
	return (void*)(__r);
}
int QVBoxLayout_count_default(void* self) {
	auto&& __r = ((QVBoxLayout*)self)->QVBoxLayout::count();
	return __r;
}
void QVBoxLayout_setGeometry_default(void* self, void* arg1) {
	((QVBoxLayout*)self)->QVBoxLayout::setGeometry(*(QRect*)arg1);
}
void* QVBoxLayout_geometry_default(void* self) {
	auto&& __r = ((QVBoxLayout*)self)->QVBoxLayout::geometry();
	return new QRect(__r);
}
int QVBoxLayout_indexOf_default(void* self, void* arg1) {
	auto&& __r = ((QVBoxLayout*)self)->QVBoxLayout::indexOf((QWidget*)arg1);
	return __r;
}
int QVBoxLayout_isEmpty_default(void* self) {
	auto&& __r = ((QVBoxLayout*)self)->QVBoxLayout::isEmpty();
	return __r;
}
int QVBoxLayout_controlTypes_default(void* self) {
	auto&& __r = ((QVBoxLayout*)self)->QVBoxLayout::controlTypes();
	return (int)(__r);
}
void* QVBoxLayout_replaceWidget_default(void* self, void* from, void* to, int options) {
	auto&& __r = ((QVBoxLayout*)self)->QVBoxLayout::replaceWidget((QWidget*)from, (QWidget*)to, QFlags<Qt::FindChildOption>(options));
	return (void*)(__r);
}
void* QVBoxLayout_layout_default(void* self) {
	auto&& __r = ((QVBoxLayout*)self)->QVBoxLayout::layout();
	return (void*)(__r);
}
void QVBoxLayout_childEvent_default(void* self, void* e) {
	((QVBoxLayout__Go*)self)->cppgoBase_childEvent((QChildEvent*)e);
}
int QVBoxLayout_event_default(void* self, void* event) {
	auto&& __r = ((QVBoxLayout*)self)->QVBoxLayout::event((QEvent*)event);
	return __r;
}
int QVBoxLayout_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QVBoxLayout*)self)->QVBoxLayout::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QVBoxLayout_timerEvent_default(void* self, void* event) {
	((QVBoxLayout__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QVBoxLayout_customEvent_default(void* self, void* event) {
	((QVBoxLayout__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QVBoxLayout_connectNotify_default(void* self, void* signal) {
	((QVBoxLayout__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QVBoxLayout_disconnectNotify_default(void* self, void* signal) {
	((QVBoxLayout__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void* QVBoxLayout_widget_default(void* self) {
	auto&& __r = ((QVBoxLayout*)self)->QVBoxLayout::widget();
	return (void*)(__r);
}
void* QVBoxLayout_spacerItem_default(void* self) {
	auto&& __r = ((QVBoxLayout*)self)->QVBoxLayout::spacerItem();
	return (void*)(__r);
}
void QVBoxLayout_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QVBoxLayout*)self, &QVBoxLayout::destroyed, (QVBoxLayout*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QVBoxLayout);
	__conns_QVBoxLayout[__id] = __c;
}
void QVBoxLayout_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QVBoxLayout*)self, &QVBoxLayout::destroyed, (QVBoxLayout*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QVBoxLayout);
	__conns_QVBoxLayout[__id] = __c;
}
void QVBoxLayout_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QVBoxLayout*)self, &QVBoxLayout::objectNameChanged, (QVBoxLayout*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QVBoxLayout);
	__conns_QVBoxLayout[__id] = __c;
}
void QVBoxLayout__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QVBoxLayout);
		auto __it = __conns_QVBoxLayout.find(id);
		if (__it == __conns_QVBoxLayout.end()) return;
		__c = __it->second;
		__conns_QVBoxLayout.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QVBoxLayout__Go_new(long long id) { return new QVBoxLayout__Go(id); }
void QVBoxLayout_delete(void* self) { delete (QVBoxLayout*)self; }

}

#include <QtWidgets/qfocusframe.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qobject.h>
#include <QtGui/qevent.h>
#include <QtWidgets/qstyleoption.h>
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
#include <QtGui/qpaintengine.h>
#include <QtGui/qbackingstore.h>
#include <QtGui/qwindow.h>
#include <QtGui/qscreen.h>
#include <QtCore/qvariant.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qfocusframe.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QFocusFrame__Go : public QFocusFrame {
public:
	long long __goId;
	unsigned char __goOn[49] = {};
	template<typename... A> QFocusFrame__Go(long long id, A&&... a) : QFocusFrame(std::forward<A>(a)...), __goId(id) {}
	using QFocusFrame::updateMicroFocus;
	using QFocusFrame::create;
	using QFocusFrame::destroy;
	using QFocusFrame::focusNextChild;
	using QFocusFrame::focusPreviousChild;
	using QFocusFrame::sender;
	using QFocusFrame::senderSignalIndex;
	using QFocusFrame::receivers;
	using QFocusFrame::isSignalConnected;
	using QFocusFrame::getDecodedMetricF;
	bool cppgoBase_event(QEvent* __a0) { return QFocusFrame::event(__a0); }
	bool cppgoBase_eventFilter(QObject* __a0, QEvent* __a1) { return QFocusFrame::eventFilter(__a0, __a1); }
	void cppgoBase_paintEvent(QPaintEvent* __a0) { QFocusFrame::paintEvent(__a0); }
	void cppgoBase_initStyleOption(QStyleOption* __a0) { QFocusFrame::initStyleOption(__a0); }
	void cppgoBase_mousePressEvent(QMouseEvent* __a0) { QFocusFrame::mousePressEvent(__a0); }
	void cppgoBase_mouseReleaseEvent(QMouseEvent* __a0) { QFocusFrame::mouseReleaseEvent(__a0); }
	void cppgoBase_mouseDoubleClickEvent(QMouseEvent* __a0) { QFocusFrame::mouseDoubleClickEvent(__a0); }
	void cppgoBase_mouseMoveEvent(QMouseEvent* __a0) { QFocusFrame::mouseMoveEvent(__a0); }
	void cppgoBase_wheelEvent(QWheelEvent* __a0) { QFocusFrame::wheelEvent(__a0); }
	void cppgoBase_keyPressEvent(QKeyEvent* __a0) { QFocusFrame::keyPressEvent(__a0); }
	void cppgoBase_keyReleaseEvent(QKeyEvent* __a0) { QFocusFrame::keyReleaseEvent(__a0); }
	void cppgoBase_focusInEvent(QFocusEvent* __a0) { QFocusFrame::focusInEvent(__a0); }
	void cppgoBase_focusOutEvent(QFocusEvent* __a0) { QFocusFrame::focusOutEvent(__a0); }
	void cppgoBase_enterEvent(QEnterEvent* __a0) { QFocusFrame::enterEvent(__a0); }
	void cppgoBase_leaveEvent(QEvent* __a0) { QFocusFrame::leaveEvent(__a0); }
	void cppgoBase_moveEvent(QMoveEvent* __a0) { QFocusFrame::moveEvent(__a0); }
	void cppgoBase_resizeEvent(QResizeEvent* __a0) { QFocusFrame::resizeEvent(__a0); }
	void cppgoBase_closeEvent(QCloseEvent* __a0) { QFocusFrame::closeEvent(__a0); }
	void cppgoBase_contextMenuEvent(QContextMenuEvent* __a0) { QFocusFrame::contextMenuEvent(__a0); }
	void cppgoBase_tabletEvent(QTabletEvent* __a0) { QFocusFrame::tabletEvent(__a0); }
	void cppgoBase_actionEvent(QActionEvent* __a0) { QFocusFrame::actionEvent(__a0); }
	void cppgoBase_dragEnterEvent(QDragEnterEvent* __a0) { QFocusFrame::dragEnterEvent(__a0); }
	void cppgoBase_dragMoveEvent(QDragMoveEvent* __a0) { QFocusFrame::dragMoveEvent(__a0); }
	void cppgoBase_dragLeaveEvent(QDragLeaveEvent* __a0) { QFocusFrame::dragLeaveEvent(__a0); }
	void cppgoBase_dropEvent(QDropEvent* __a0) { QFocusFrame::dropEvent(__a0); }
	void cppgoBase_showEvent(QShowEvent* __a0) { QFocusFrame::showEvent(__a0); }
	void cppgoBase_hideEvent(QHideEvent* __a0) { QFocusFrame::hideEvent(__a0); }
	void cppgoBase_changeEvent(QEvent* __a0) { QFocusFrame::changeEvent(__a0); }
	int cppgoBase_metric(int __a0) { return QFocusFrame::metric((QPaintDevice::PaintDeviceMetric)__a0); }
	void cppgoBase_initPainter(QPainter* __a0) { QFocusFrame::initPainter(__a0); }
	QPaintDevice* cppgoBase_redirected(QPoint* __a0) { return QFocusFrame::redirected(__a0); }
	QPainter* cppgoBase_sharedPainter() { return QFocusFrame::sharedPainter(); }
	void cppgoBase_inputMethodEvent(QInputMethodEvent* __a0) { QFocusFrame::inputMethodEvent(__a0); }
	bool cppgoBase_focusNextPrevChild(bool __a0) { return QFocusFrame::focusNextPrevChild(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QFocusFrame::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QFocusFrame::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QFocusFrame::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QFocusFrame::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QFocusFrame::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QFocusFrame::metaObject();
	}
	bool event(QEvent* e) override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)e, &__h); if (__h) return __r; }
		return QFocusFrame::event(e);
	}
	bool eventFilter(QObject* arg1, QEvent* arg2) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)arg1, (void*)arg2, &__h); if (__h) return __r; }
		return QFocusFrame::eventFilter(arg1, arg2);
	}
	void paintEvent(QPaintEvent* arg1) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5290442163167458465LL, (void*)arg1, &__h); if (__h) return; }
		QFocusFrame::paintEvent(arg1);
	}
	void initStyleOption(QStyleOption* option) const override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -488693769771308531LL, (void*)option, &__h); if (__h) return; }
		QFocusFrame::initStyleOption(option);
	}
	int devType() const override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8659679893868524044LL, &__h); if (__h) return __r; }
		return QFocusFrame::devType();
	}
	void setVisible(bool visible) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 2110182719375245575LL, visible, &__h); if (__h) return; }
		QFocusFrame::setVisible(visible);
	}
	QSize sizeHint() const override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 729690261457684101LL, &__h); if (__h) return *(QSize*)__r; }
		return QFocusFrame::sizeHint();
	}
	QSize minimumSizeHint() const override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 7049264599837532491LL, &__h); if (__h) return *(QSize*)__r; }
		return QFocusFrame::minimumSizeHint();
	}
	int heightForWidth(int arg1) const override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 4311534060495737339LL, arg1, &__h); if (__h) return __r; }
		return QFocusFrame::heightForWidth(arg1);
	}
	bool hasHeightForWidth() const override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8266762397995423915LL, &__h); if (__h) return __r; }
		return QFocusFrame::hasHeightForWidth();
	}
	QPaintEngine* paintEngine() const override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 6398007404756346381LL, &__h); if (__h) return (QPaintEngine*)__r; }
		return QFocusFrame::paintEngine();
	}
	void mousePressEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -332090898078543319LL, (void*)event, &__h); if (__h) return; }
		QFocusFrame::mousePressEvent(event);
	}
	void mouseReleaseEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4415920049297795213LL, (void*)event, &__h); if (__h) return; }
		QFocusFrame::mouseReleaseEvent(event);
	}
	void mouseDoubleClickEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4029073905697023887LL, (void*)event, &__h); if (__h) return; }
		QFocusFrame::mouseDoubleClickEvent(event);
	}
	void mouseMoveEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4255494004099806581LL, (void*)event, &__h); if (__h) return; }
		QFocusFrame::mouseMoveEvent(event);
	}
	void wheelEvent(QWheelEvent* event) override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1919988340934511316LL, (void*)event, &__h); if (__h) return; }
		QFocusFrame::wheelEvent(event);
	}
	void keyPressEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1023791977679793185LL, (void*)event, &__h); if (__h) return; }
		QFocusFrame::keyPressEvent(event);
	}
	void keyReleaseEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 772598475950050299LL, (void*)event, &__h); if (__h) return; }
		QFocusFrame::keyReleaseEvent(event);
	}
	void focusInEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 975043767033552552LL, (void*)event, &__h); if (__h) return; }
		QFocusFrame::focusInEvent(event);
	}
	void focusOutEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -385596010576048993LL, (void*)event, &__h); if (__h) return; }
		QFocusFrame::focusOutEvent(event);
	}
	void enterEvent(QEnterEvent* event) override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6316270346668237895LL, (void*)event, &__h); if (__h) return; }
		QFocusFrame::enterEvent(event);
	}
	void leaveEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -551070009623711640LL, (void*)event, &__h); if (__h) return; }
		QFocusFrame::leaveEvent(event);
	}
	void moveEvent(QMoveEvent* event) override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -9161949150489525396LL, (void*)event, &__h); if (__h) return; }
		QFocusFrame::moveEvent(event);
	}
	void resizeEvent(QResizeEvent* event) override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4302917285655014661LL, (void*)event, &__h); if (__h) return; }
		QFocusFrame::resizeEvent(event);
	}
	void closeEvent(QCloseEvent* event) override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1989937050019464783LL, (void*)event, &__h); if (__h) return; }
		QFocusFrame::closeEvent(event);
	}
	void contextMenuEvent(QContextMenuEvent* event) override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6698043222227006015LL, (void*)event, &__h); if (__h) return; }
		QFocusFrame::contextMenuEvent(event);
	}
	void tabletEvent(QTabletEvent* event) override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -212068866860701141LL, (void*)event, &__h); if (__h) return; }
		QFocusFrame::tabletEvent(event);
	}
	void actionEvent(QActionEvent* event) override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1697444140088555581LL, (void*)event, &__h); if (__h) return; }
		QFocusFrame::actionEvent(event);
	}
	void dragEnterEvent(QDragEnterEvent* event) override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 206122715919917589LL, (void*)event, &__h); if (__h) return; }
		QFocusFrame::dragEnterEvent(event);
	}
	void dragMoveEvent(QDragMoveEvent* event) override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1781545137352745978LL, (void*)event, &__h); if (__h) return; }
		QFocusFrame::dragMoveEvent(event);
	}
	void dragLeaveEvent(QDragLeaveEvent* event) override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -3139904158373725006LL, (void*)event, &__h); if (__h) return; }
		QFocusFrame::dragLeaveEvent(event);
	}
	void dropEvent(QDropEvent* event) override {
		int __h = 0;
		if (__goOn[32] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 9107058682754591716LL, (void*)event, &__h); if (__h) return; }
		QFocusFrame::dropEvent(event);
	}
	void showEvent(QShowEvent* event) override {
		int __h = 0;
		if (__goOn[33] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6825369174751916LL, (void*)event, &__h); if (__h) return; }
		QFocusFrame::showEvent(event);
	}
	void hideEvent(QHideEvent* event) override {
		int __h = 0;
		if (__goOn[34] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -199718707876719129LL, (void*)event, &__h); if (__h) return; }
		QFocusFrame::hideEvent(event);
	}
	bool nativeEvent(const QByteArray& eventType, void* message, qintptr* result) override {
		int __h = 0;
		if (__goOn[35] && cppgo_virt_int_voidp_voidp_longlongp) { int __r = ((int(*)(long long, long long, void*, void*, long long*, int*))cppgo_virt_int_voidp_voidp_longlongp)(__goId, -2039166586246443706LL, (void*)&eventType, (void*)message, (long long*)result, &__h); if (__h) return __r; }
		return QFocusFrame::nativeEvent(eventType, message, result);
	}
	void changeEvent(QEvent* arg1) override {
		int __h = 0;
		if (__goOn[36] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4959376063826165943LL, (void*)arg1, &__h); if (__h) return; }
		QFocusFrame::changeEvent(arg1);
	}
	int metric(QPaintDevice::PaintDeviceMetric arg1) const override {
		int __h = 0;
		if (__goOn[37] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 427629649267311777LL, (int)(arg1), &__h); if (__h) return __r; }
		return QFocusFrame::metric(arg1);
	}
	void initPainter(QPainter* painter) const override {
		int __h = 0;
		if (__goOn[38] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6477500296980050210LL, (void*)painter, &__h); if (__h) return; }
		QFocusFrame::initPainter(painter);
	}
	QPaintDevice* redirected(QPoint* offset) const override {
		int __h = 0;
		if (__goOn[39] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 4127240633310655048LL, (void*)offset, &__h); if (__h) return (QPaintDevice*)__r; }
		return QFocusFrame::redirected(offset);
	}
	QPainter* sharedPainter() const override {
		int __h = 0;
		if (__goOn[40] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 5792771328741774991LL, &__h); if (__h) return (QPainter*)__r; }
		return QFocusFrame::sharedPainter();
	}
	void inputMethodEvent(QInputMethodEvent* arg1) override {
		int __h = 0;
		if (__goOn[41] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5438819661942843222LL, (void*)arg1, &__h); if (__h) return; }
		QFocusFrame::inputMethodEvent(arg1);
	}
	QVariant inputMethodQuery(Qt::InputMethodQuery arg1) const override {
		int __h = 0;
		if (__goOn[42] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7345853097340709638LL, (int)(arg1), &__h); if (__h) return *(QVariant*)__r; }
		return QFocusFrame::inputMethodQuery(arg1);
	}
	bool focusNextPrevChild(bool next) override {
		int __h = 0;
		if (__goOn[43] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 5898328517505784191LL, next, &__h); if (__h) return __r; }
		return QFocusFrame::focusNextPrevChild(next);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[44] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QFocusFrame::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[45] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QFocusFrame::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[46] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QFocusFrame::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[47] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QFocusFrame::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[48] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QFocusFrame::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QFocusFrame;
static std::mutex __connsMu_QFocusFrame;

extern "C" {

void* QFocusFrame_base(void* self) { return (void*)static_cast<QWidget*>((QFocusFrame*)self); }
void QFocusFrame_chain(void* self, void** out) {
	QFocusFrame* __p = (QFocusFrame*)self;
	QWidget* __b0 = static_cast<QWidget*>(__p); out[0] = (void*)__b0;
	QObject* __b1 = static_cast<QObject*>(__b0); out[1] = (void*)__b1;
}
void QFocusFrame__set_override(void* self, int idx) { ((QFocusFrame__Go*)self)->__goOn[idx] = 1; }
void* QFocusFrame_fromQObject(void* o) { return (void*)qobject_cast<QFocusFrame*>((QObject*)o); }
void* QFocusFrame_new(long long __goid, void* parent) {
	return new QFocusFrame__Go(__goid, (QWidget*)parent);
}
void* QFocusFrame_new2(long long __goid) {
	return new QFocusFrame__Go(__goid);
}
char* QFocusFrame_tr(const char* s, long long* __out_len) {
	auto&& __r = QFocusFrame::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QFocusFrame_setWidget(void* self, void* widget) {
	((QFocusFrame*)self)->setWidget((QWidget*)widget);
}
void* QFocusFrame_widget(void* self) {
	auto&& __r = ((QFocusFrame*)self)->widget();
	return (void*)(__r);
}
void QFocusFrame_setTabOrder(void* arg1, void* arg2) {
	QFocusFrame::setTabOrder((QWidget*)arg1, (QWidget*)arg2);
}
void* QFocusFrame_mouseGrabber() {
	auto&& __r = QFocusFrame::mouseGrabber();
	return (void*)(__r);
}
void* QFocusFrame_keyboardGrabber() {
	auto&& __r = QFocusFrame::keyboardGrabber();
	return (void*)(__r);
}
void* QFocusFrame_find(unsigned long long arg1) {
	auto&& __r = QFocusFrame::find(arg1);
	return (void*)(__r);
}
void* QFocusFrame_createWindowContainer(void* window) {
	auto&& __r = QFocusFrame::createWindowContainer((QWindow*)window);
	return (void*)(__r);
}
void QFocusFrame_updateMicroFocus(void* self) {
	((QFocusFrame__Go*)self)->updateMicroFocus();
}
void QFocusFrame_create(void* self) {
	((QFocusFrame__Go*)self)->create();
}
void QFocusFrame_destroy(void* self) {
	((QFocusFrame__Go*)self)->destroy();
}
int QFocusFrame_focusNextChild(void* self) {
	auto&& __r = ((QFocusFrame__Go*)self)->focusNextChild();
	return __r;
}
int QFocusFrame_focusPreviousChild(void* self) {
	auto&& __r = ((QFocusFrame__Go*)self)->focusPreviousChild();
	return __r;
}
void* QFocusFrame_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QFocusFrame::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QFocusFrame_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QFocusFrame::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QFocusFrame_sender(void* self) {
	auto&& __r = ((QFocusFrame__Go*)self)->sender();
	return (void*)(__r);
}
int QFocusFrame_senderSignalIndex(void* self) {
	auto&& __r = ((QFocusFrame__Go*)self)->senderSignalIndex();
	return __r;
}
int QFocusFrame_receivers(void* self, const char* signal) {
	auto&& __r = ((QFocusFrame__Go*)self)->receivers(signal);
	return __r;
}
int QFocusFrame_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QFocusFrame__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
double QFocusFrame_devicePixelRatioFScale() {
	auto&& __r = QFocusFrame::devicePixelRatioFScale();
	return __r;
}
int QFocusFrame_encodeMetricF(int metric, double value) {
	auto&& __r = QFocusFrame::encodeMetricF((QPaintDevice::PaintDeviceMetric)(metric), value);
	return __r;
}
double QFocusFrame_getDecodedMetricF(void* self, int metricA, int metricB) {
	auto&& __r = ((QFocusFrame__Go*)self)->getDecodedMetricF((QPaintDevice::PaintDeviceMetric)(metricA), (QPaintDevice::PaintDeviceMetric)(metricB));
	return __r;
}
char* QFocusFrame_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QFocusFrame::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QFocusFrame_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QFocusFrame::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QFocusFrame_createWindowContainer2(void* window, void* parent) {
	auto&& __r = QFocusFrame::createWindowContainer((QWindow*)window, (QWidget*)parent);
	return (void*)(__r);
}
void* QFocusFrame_createWindowContainer3(void* window, void* parent, int flags) {
	auto&& __r = QFocusFrame::createWindowContainer((QWindow*)window, (QWidget*)parent, QFlags<Qt::WindowType>(flags));
	return (void*)(__r);
}
void QFocusFrame_updateMicroFocus2(void* self, int query) {
	((QFocusFrame__Go*)self)->updateMicroFocus((Qt::InputMethodQuery)(query));
}
void QFocusFrame_create2(void* self, unsigned long long arg1) {
	((QFocusFrame__Go*)self)->create(arg1);
}
void QFocusFrame_create3(void* self, unsigned long long arg1, int initializeWindow) {
	((QFocusFrame__Go*)self)->create(arg1, initializeWindow);
}
void QFocusFrame_create4(void* self, unsigned long long arg1, int initializeWindow, int destroyOldWindow) {
	((QFocusFrame__Go*)self)->create(arg1, initializeWindow, destroyOldWindow);
}
void QFocusFrame_destroy2(void* self, int destroyWindow) {
	((QFocusFrame__Go*)self)->destroy(destroyWindow);
}
void QFocusFrame_destroy3(void* self, int destroyWindow, int destroySubWindows) {
	((QFocusFrame__Go*)self)->destroy(destroyWindow, destroySubWindows);
}
void* QFocusFrame_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QFocusFrame::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QFocusFrame_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QFocusFrame::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QFocusFrame_metaObject_default(void* self) {
	auto&& __r = ((QFocusFrame*)self)->QFocusFrame::metaObject();
	return (void*)(__r);
}
int QFocusFrame_event_default(void* self, void* e) {
	auto&& __r = ((QFocusFrame__Go*)self)->cppgoBase_event((QEvent*)e);
	return __r;
}
int QFocusFrame_eventFilter_default(void* self, void* arg1, void* arg2) {
	auto&& __r = ((QFocusFrame__Go*)self)->cppgoBase_eventFilter((QObject*)arg1, (QEvent*)arg2);
	return __r;
}
void QFocusFrame_paintEvent_default(void* self, void* arg1) {
	((QFocusFrame__Go*)self)->cppgoBase_paintEvent((QPaintEvent*)arg1);
}
void QFocusFrame_initStyleOption_default(void* self, void* option) {
	((QFocusFrame__Go*)self)->cppgoBase_initStyleOption((QStyleOption*)option);
}
int QFocusFrame_devType_default(void* self) {
	auto&& __r = ((QFocusFrame*)self)->QFocusFrame::devType();
	return __r;
}
void QFocusFrame_setVisible_default(void* self, int visible) {
	((QFocusFrame*)self)->QFocusFrame::setVisible(visible);
}
void* QFocusFrame_sizeHint_default(void* self) {
	auto&& __r = ((QFocusFrame*)self)->QFocusFrame::sizeHint();
	return new QSize(__r);
}
void* QFocusFrame_minimumSizeHint_default(void* self) {
	auto&& __r = ((QFocusFrame*)self)->QFocusFrame::minimumSizeHint();
	return new QSize(__r);
}
int QFocusFrame_heightForWidth_default(void* self, int arg1) {
	auto&& __r = ((QFocusFrame*)self)->QFocusFrame::heightForWidth(arg1);
	return __r;
}
int QFocusFrame_hasHeightForWidth_default(void* self) {
	auto&& __r = ((QFocusFrame*)self)->QFocusFrame::hasHeightForWidth();
	return __r;
}
void* QFocusFrame_paintEngine_default(void* self) {
	auto&& __r = ((QFocusFrame*)self)->QFocusFrame::paintEngine();
	return (void*)(__r);
}
void QFocusFrame_mousePressEvent_default(void* self, void* event) {
	((QFocusFrame__Go*)self)->cppgoBase_mousePressEvent((QMouseEvent*)event);
}
void QFocusFrame_mouseReleaseEvent_default(void* self, void* event) {
	((QFocusFrame__Go*)self)->cppgoBase_mouseReleaseEvent((QMouseEvent*)event);
}
void QFocusFrame_mouseDoubleClickEvent_default(void* self, void* event) {
	((QFocusFrame__Go*)self)->cppgoBase_mouseDoubleClickEvent((QMouseEvent*)event);
}
void QFocusFrame_mouseMoveEvent_default(void* self, void* event) {
	((QFocusFrame__Go*)self)->cppgoBase_mouseMoveEvent((QMouseEvent*)event);
}
void QFocusFrame_wheelEvent_default(void* self, void* event) {
	((QFocusFrame__Go*)self)->cppgoBase_wheelEvent((QWheelEvent*)event);
}
void QFocusFrame_keyPressEvent_default(void* self, void* event) {
	((QFocusFrame__Go*)self)->cppgoBase_keyPressEvent((QKeyEvent*)event);
}
void QFocusFrame_keyReleaseEvent_default(void* self, void* event) {
	((QFocusFrame__Go*)self)->cppgoBase_keyReleaseEvent((QKeyEvent*)event);
}
void QFocusFrame_focusInEvent_default(void* self, void* event) {
	((QFocusFrame__Go*)self)->cppgoBase_focusInEvent((QFocusEvent*)event);
}
void QFocusFrame_focusOutEvent_default(void* self, void* event) {
	((QFocusFrame__Go*)self)->cppgoBase_focusOutEvent((QFocusEvent*)event);
}
void QFocusFrame_enterEvent_default(void* self, void* event) {
	((QFocusFrame__Go*)self)->cppgoBase_enterEvent((QEnterEvent*)event);
}
void QFocusFrame_leaveEvent_default(void* self, void* event) {
	((QFocusFrame__Go*)self)->cppgoBase_leaveEvent((QEvent*)event);
}
void QFocusFrame_moveEvent_default(void* self, void* event) {
	((QFocusFrame__Go*)self)->cppgoBase_moveEvent((QMoveEvent*)event);
}
void QFocusFrame_resizeEvent_default(void* self, void* event) {
	((QFocusFrame__Go*)self)->cppgoBase_resizeEvent((QResizeEvent*)event);
}
void QFocusFrame_closeEvent_default(void* self, void* event) {
	((QFocusFrame__Go*)self)->cppgoBase_closeEvent((QCloseEvent*)event);
}
void QFocusFrame_contextMenuEvent_default(void* self, void* event) {
	((QFocusFrame__Go*)self)->cppgoBase_contextMenuEvent((QContextMenuEvent*)event);
}
void QFocusFrame_tabletEvent_default(void* self, void* event) {
	((QFocusFrame__Go*)self)->cppgoBase_tabletEvent((QTabletEvent*)event);
}
void QFocusFrame_actionEvent_default(void* self, void* event) {
	((QFocusFrame__Go*)self)->cppgoBase_actionEvent((QActionEvent*)event);
}
void QFocusFrame_dragEnterEvent_default(void* self, void* event) {
	((QFocusFrame__Go*)self)->cppgoBase_dragEnterEvent((QDragEnterEvent*)event);
}
void QFocusFrame_dragMoveEvent_default(void* self, void* event) {
	((QFocusFrame__Go*)self)->cppgoBase_dragMoveEvent((QDragMoveEvent*)event);
}
void QFocusFrame_dragLeaveEvent_default(void* self, void* event) {
	((QFocusFrame__Go*)self)->cppgoBase_dragLeaveEvent((QDragLeaveEvent*)event);
}
void QFocusFrame_dropEvent_default(void* self, void* event) {
	((QFocusFrame__Go*)self)->cppgoBase_dropEvent((QDropEvent*)event);
}
void QFocusFrame_showEvent_default(void* self, void* event) {
	((QFocusFrame__Go*)self)->cppgoBase_showEvent((QShowEvent*)event);
}
void QFocusFrame_hideEvent_default(void* self, void* event) {
	((QFocusFrame__Go*)self)->cppgoBase_hideEvent((QHideEvent*)event);
}
void QFocusFrame_changeEvent_default(void* self, void* arg1) {
	((QFocusFrame__Go*)self)->cppgoBase_changeEvent((QEvent*)arg1);
}
int QFocusFrame_metric_default(void* self, int arg1) {
	auto&& __r = ((QFocusFrame__Go*)self)->cppgoBase_metric(arg1);
	return __r;
}
void QFocusFrame_initPainter_default(void* self, void* painter) {
	((QFocusFrame__Go*)self)->cppgoBase_initPainter((QPainter*)painter);
}
void* QFocusFrame_redirected_default(void* self, void* offset) {
	auto&& __r = ((QFocusFrame__Go*)self)->cppgoBase_redirected((QPoint*)offset);
	return (void*)(__r);
}
void* QFocusFrame_sharedPainter_default(void* self) {
	auto&& __r = ((QFocusFrame__Go*)self)->cppgoBase_sharedPainter();
	return (void*)(__r);
}
void QFocusFrame_inputMethodEvent_default(void* self, void* arg1) {
	((QFocusFrame__Go*)self)->cppgoBase_inputMethodEvent((QInputMethodEvent*)arg1);
}
void* QFocusFrame_inputMethodQuery_default(void* self, int arg1) {
	auto&& __r = ((QFocusFrame*)self)->QFocusFrame::inputMethodQuery((Qt::InputMethodQuery)(arg1));
	return new QVariant(__r);
}
int QFocusFrame_focusNextPrevChild_default(void* self, int next) {
	auto&& __r = ((QFocusFrame__Go*)self)->cppgoBase_focusNextPrevChild(next);
	return __r;
}
void QFocusFrame_timerEvent_default(void* self, void* event) {
	((QFocusFrame__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QFocusFrame_childEvent_default(void* self, void* event) {
	((QFocusFrame__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QFocusFrame_customEvent_default(void* self, void* event) {
	((QFocusFrame__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QFocusFrame_connectNotify_default(void* self, void* signal) {
	((QFocusFrame__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QFocusFrame_disconnectNotify_default(void* self, void* signal) {
	((QFocusFrame__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QFocusFrame_connect_windowTitleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFocusFrame*)self, &QFocusFrame::windowTitleChanged, (QFocusFrame*)self, [__id](const QString& title){
		QByteArray title_ba = title.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)title_ba.constData(), title_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFocusFrame);
	__conns_QFocusFrame[__id] = __c;
}
void QFocusFrame_connect_windowIconChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFocusFrame*)self, &QFocusFrame::windowIconChanged, (QFocusFrame*)self, [__id](const QIcon& icon){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&icon);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFocusFrame);
	__conns_QFocusFrame[__id] = __c;
}
void QFocusFrame_connect_windowIconTextChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFocusFrame*)self, &QFocusFrame::windowIconTextChanged, (QFocusFrame*)self, [__id](const QString& iconText){
		QByteArray iconText_ba = iconText.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)iconText_ba.constData(), iconText_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFocusFrame);
	__conns_QFocusFrame[__id] = __c;
}
void QFocusFrame_connect_customContextMenuRequested(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFocusFrame*)self, &QFocusFrame::customContextMenuRequested, (QFocusFrame*)self, [__id](const QPoint& pos){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&pos);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFocusFrame);
	__conns_QFocusFrame[__id] = __c;
}
void QFocusFrame_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFocusFrame*)self, &QFocusFrame::destroyed, (QFocusFrame*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFocusFrame);
	__conns_QFocusFrame[__id] = __c;
}
void QFocusFrame_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFocusFrame*)self, &QFocusFrame::destroyed, (QFocusFrame*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFocusFrame);
	__conns_QFocusFrame[__id] = __c;
}
void QFocusFrame_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFocusFrame*)self, &QFocusFrame::objectNameChanged, (QFocusFrame*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFocusFrame);
	__conns_QFocusFrame[__id] = __c;
}
void QFocusFrame__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QFocusFrame);
		auto __it = __conns_QFocusFrame.find(id);
		if (__it == __conns_QFocusFrame.end()) return;
		__c = __it->second;
		__conns_QFocusFrame.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QFocusFrame__Go_new(long long id) { return new QFocusFrame__Go(id); }
void QFocusFrame_delete(void* self) { delete (QFocusFrame*)self; }

}

#include <QtWidgets/qgraphicsitem.h>
#include <QtCore/qrect.h>
#include <QtGui/qpainter.h>
#include <QtWidgets/qstyleoption.h>
#include <QtWidgets/qwidget.h>
#include <QtGui/qpainterpath.h>
#include <QtWidgets/qgraphicsscene.h>
#include <QtWidgets/qgraphicswidget.h>
#include <QtCore/qsize.h>
#include <QtCore/qstring.h>
#include <QtGui/qcursor.h>
#include <QtWidgets/qgraphicseffect.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qpoint.h>
#include <QtGui/qtransform.h>
#include <QtWidgets/qgraphicstransform.h>
#include <QtGui/qregion.h>
#include <QtGui/qpolygon.h>
#include <QtCore/qvariant.h>
#include <QtCore/qcoreevent.h>
#include <QtWidgets/qgraphicssceneevent.h>
#include <QtGui/qevent.h>
#include "gen_qgraphicsitemgroup.h"
#include "gen_purego_cb.h"

class QGraphicsItemGroup__Go : public QGraphicsItemGroup {
public:
	long long __goId;
	unsigned char __goOn[35] = {};
	template<typename... A> QGraphicsItemGroup__Go(long long id, A&&... a) : QGraphicsItemGroup(std::forward<A>(a)...), __goId(id) {}
	using QGraphicsItemGroup::updateMicroFocus;
	using QGraphicsItemGroup::addToIndex;
	using QGraphicsItemGroup::removeFromIndex;
	using QGraphicsItemGroup::prepareGeometryChange;
	bool cppgoBase_sceneEventFilter(QGraphicsItem* __a0, QEvent* __a1) { return QGraphicsItemGroup::sceneEventFilter(__a0, __a1); }
	bool cppgoBase_sceneEvent(QEvent* __a0) { return QGraphicsItemGroup::sceneEvent(__a0); }
	void cppgoBase_contextMenuEvent(QGraphicsSceneContextMenuEvent* __a0) { QGraphicsItemGroup::contextMenuEvent(__a0); }
	void cppgoBase_dragEnterEvent(QGraphicsSceneDragDropEvent* __a0) { QGraphicsItemGroup::dragEnterEvent(__a0); }
	void cppgoBase_dragLeaveEvent(QGraphicsSceneDragDropEvent* __a0) { QGraphicsItemGroup::dragLeaveEvent(__a0); }
	void cppgoBase_dragMoveEvent(QGraphicsSceneDragDropEvent* __a0) { QGraphicsItemGroup::dragMoveEvent(__a0); }
	void cppgoBase_dropEvent(QGraphicsSceneDragDropEvent* __a0) { QGraphicsItemGroup::dropEvent(__a0); }
	void cppgoBase_focusInEvent(QFocusEvent* __a0) { QGraphicsItemGroup::focusInEvent(__a0); }
	void cppgoBase_focusOutEvent(QFocusEvent* __a0) { QGraphicsItemGroup::focusOutEvent(__a0); }
	void cppgoBase_hoverEnterEvent(QGraphicsSceneHoverEvent* __a0) { QGraphicsItemGroup::hoverEnterEvent(__a0); }
	void cppgoBase_hoverMoveEvent(QGraphicsSceneHoverEvent* __a0) { QGraphicsItemGroup::hoverMoveEvent(__a0); }
	void cppgoBase_hoverLeaveEvent(QGraphicsSceneHoverEvent* __a0) { QGraphicsItemGroup::hoverLeaveEvent(__a0); }
	void cppgoBase_keyPressEvent(QKeyEvent* __a0) { QGraphicsItemGroup::keyPressEvent(__a0); }
	void cppgoBase_keyReleaseEvent(QKeyEvent* __a0) { QGraphicsItemGroup::keyReleaseEvent(__a0); }
	void cppgoBase_mousePressEvent(QGraphicsSceneMouseEvent* __a0) { QGraphicsItemGroup::mousePressEvent(__a0); }
	void cppgoBase_mouseMoveEvent(QGraphicsSceneMouseEvent* __a0) { QGraphicsItemGroup::mouseMoveEvent(__a0); }
	void cppgoBase_mouseReleaseEvent(QGraphicsSceneMouseEvent* __a0) { QGraphicsItemGroup::mouseReleaseEvent(__a0); }
	void cppgoBase_mouseDoubleClickEvent(QGraphicsSceneMouseEvent* __a0) { QGraphicsItemGroup::mouseDoubleClickEvent(__a0); }
	void cppgoBase_wheelEvent(QGraphicsSceneWheelEvent* __a0) { QGraphicsItemGroup::wheelEvent(__a0); }
	void cppgoBase_inputMethodEvent(QInputMethodEvent* __a0) { QGraphicsItemGroup::inputMethodEvent(__a0); }
	QVariant cppgoBase_inputMethodQuery(int __a0) { return QGraphicsItemGroup::inputMethodQuery((Qt::InputMethodQuery)__a0); }
	QVariant cppgoBase_itemChange(int __a0, const QVariant& __a1) { return QGraphicsItemGroup::itemChange((QGraphicsItem::GraphicsItemChange)__a0, __a1); }
	bool cppgoBase_supportsExtension(int __a0) { return QGraphicsItemGroup::supportsExtension((QGraphicsItem::Extension)__a0); }
	void cppgoBase_setExtension(int __a0, const QVariant& __a1) { QGraphicsItemGroup::setExtension((QGraphicsItem::Extension)__a0, __a1); }
	QVariant cppgoBase_extension(const QVariant& __a0) { return QGraphicsItemGroup::extension(__a0); }
	QRectF boundingRect() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 2510008782785038831LL, &__h); if (__h) return *(QRectF*)__r; }
		return QGraphicsItemGroup::boundingRect();
	}
	void paint(QPainter* painter, const QStyleOptionGraphicsItem* option, QWidget* widget) override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_void_voidp_voidp_voidp) { ((void(*)(long long, long long, void*, void*, void*, int*))cppgo_virt_void_voidp_voidp_voidp)(__goId, 6083584261053556819LL, (void*)painter, (void*)option, (void*)widget, &__h); if (__h) return; }
		QGraphicsItemGroup::paint(painter, option, widget);
	}
	bool isObscuredBy(const QGraphicsItem* item) const override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, -7215030271914903671LL, (void*)item, &__h); if (__h) return __r; }
		return QGraphicsItemGroup::isObscuredBy(item);
	}
	QPainterPath opaqueArea() const override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 399370970023989435LL, &__h); if (__h) return *(QPainterPath*)__r; }
		return QGraphicsItemGroup::opaqueArea();
	}
	int type() const override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -6371403872082420691LL, &__h); if (__h) return __r; }
		return QGraphicsItemGroup::type();
	}
	void advance(int phase) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, -4830836341480073475LL, phase, &__h); if (__h) return; }
		QGraphicsItemGroup::advance(phase);
	}
	QPainterPath shape() const override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 7938019069247444294LL, &__h); if (__h) return *(QPainterPath*)__r; }
		return QGraphicsItemGroup::shape();
	}
	bool contains(const QPointF& point) const override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 5645366511944179768LL, (void*)&point, &__h); if (__h) return __r; }
		return QGraphicsItemGroup::contains(point);
	}
	bool collidesWithItem(const QGraphicsItem* other, Qt::ItemSelectionMode mode) const override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_int_voidp_int) { int __r = ((int(*)(long long, long long, void*, int, int*))cppgo_virt_int_voidp_int)(__goId, 6469736184362593173LL, (void*)other, (int)(mode), &__h); if (__h) return __r; }
		return QGraphicsItemGroup::collidesWithItem(other, mode);
	}
	bool collidesWithPath(const QPainterPath& path, Qt::ItemSelectionMode mode) const override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_int_voidp_int) { int __r = ((int(*)(long long, long long, void*, int, int*))cppgo_virt_int_voidp_int)(__goId, 6636575182718753853LL, (void*)&path, (int)(mode), &__h); if (__h) return __r; }
		return QGraphicsItemGroup::collidesWithPath(path, mode);
	}
	bool sceneEventFilter(QGraphicsItem* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, -4584203218241215725LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QGraphicsItemGroup::sceneEventFilter(watched, event);
	}
	bool sceneEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, -2846371507838587071LL, (void*)event, &__h); if (__h) return __r; }
		return QGraphicsItemGroup::sceneEvent(event);
	}
	void contextMenuEvent(QGraphicsSceneContextMenuEvent* event) override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6698043222227006015LL, (void*)event, &__h); if (__h) return; }
		QGraphicsItemGroup::contextMenuEvent(event);
	}
	void dragEnterEvent(QGraphicsSceneDragDropEvent* event) override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 206122715919917589LL, (void*)event, &__h); if (__h) return; }
		QGraphicsItemGroup::dragEnterEvent(event);
	}
	void dragLeaveEvent(QGraphicsSceneDragDropEvent* event) override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -3139904158373725006LL, (void*)event, &__h); if (__h) return; }
		QGraphicsItemGroup::dragLeaveEvent(event);
	}
	void dragMoveEvent(QGraphicsSceneDragDropEvent* event) override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1781545137352745978LL, (void*)event, &__h); if (__h) return; }
		QGraphicsItemGroup::dragMoveEvent(event);
	}
	void dropEvent(QGraphicsSceneDragDropEvent* event) override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 9107058682754591716LL, (void*)event, &__h); if (__h) return; }
		QGraphicsItemGroup::dropEvent(event);
	}
	void focusInEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 975043767033552552LL, (void*)event, &__h); if (__h) return; }
		QGraphicsItemGroup::focusInEvent(event);
	}
	void focusOutEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -385596010576048993LL, (void*)event, &__h); if (__h) return; }
		QGraphicsItemGroup::focusOutEvent(event);
	}
	void hoverEnterEvent(QGraphicsSceneHoverEvent* event) override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6932941943906681679LL, (void*)event, &__h); if (__h) return; }
		QGraphicsItemGroup::hoverEnterEvent(event);
	}
	void hoverMoveEvent(QGraphicsSceneHoverEvent* event) override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -2745191522144316396LL, (void*)event, &__h); if (__h) return; }
		QGraphicsItemGroup::hoverMoveEvent(event);
	}
	void hoverLeaveEvent(QGraphicsSceneHoverEvent* event) override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7944408482203603984LL, (void*)event, &__h); if (__h) return; }
		QGraphicsItemGroup::hoverLeaveEvent(event);
	}
	void keyPressEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1023791977679793185LL, (void*)event, &__h); if (__h) return; }
		QGraphicsItemGroup::keyPressEvent(event);
	}
	void keyReleaseEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 772598475950050299LL, (void*)event, &__h); if (__h) return; }
		QGraphicsItemGroup::keyReleaseEvent(event);
	}
	void mousePressEvent(QGraphicsSceneMouseEvent* event) override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -332090898078543319LL, (void*)event, &__h); if (__h) return; }
		QGraphicsItemGroup::mousePressEvent(event);
	}
	void mouseMoveEvent(QGraphicsSceneMouseEvent* event) override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4255494004099806581LL, (void*)event, &__h); if (__h) return; }
		QGraphicsItemGroup::mouseMoveEvent(event);
	}
	void mouseReleaseEvent(QGraphicsSceneMouseEvent* event) override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4415920049297795213LL, (void*)event, &__h); if (__h) return; }
		QGraphicsItemGroup::mouseReleaseEvent(event);
	}
	void mouseDoubleClickEvent(QGraphicsSceneMouseEvent* event) override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4029073905697023887LL, (void*)event, &__h); if (__h) return; }
		QGraphicsItemGroup::mouseDoubleClickEvent(event);
	}
	void wheelEvent(QGraphicsSceneWheelEvent* event) override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1919988340934511316LL, (void*)event, &__h); if (__h) return; }
		QGraphicsItemGroup::wheelEvent(event);
	}
	void inputMethodEvent(QInputMethodEvent* event) override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5438819661942843222LL, (void*)event, &__h); if (__h) return; }
		QGraphicsItemGroup::inputMethodEvent(event);
	}
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7345853097340709638LL, (int)(query), &__h); if (__h) return *(QVariant*)__r; }
		return QGraphicsItemGroup::inputMethodQuery(query);
	}
	QVariant itemChange(QGraphicsItem::GraphicsItemChange change, const QVariant& value) override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_voidp_int_voidp) { void* __r = ((void*(*)(long long, long long, int, void*, int*))cppgo_virt_voidp_int_voidp)(__goId, -1888572168794878830LL, (int)(change), (void*)&value, &__h); if (__h) return *(QVariant*)__r; }
		return QGraphicsItemGroup::itemChange(change, value);
	}
	bool supportsExtension(QGraphicsItem::Extension extension) const override {
		int __h = 0;
		if (__goOn[32] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, -5729166443651306418LL, (int)(extension), &__h); if (__h) return __r; }
		return QGraphicsItemGroup::supportsExtension(extension);
	}
	void setExtension(QGraphicsItem::Extension extension, const QVariant& variant) override {
		int __h = 0;
		if (__goOn[33] && cppgo_virt_void_int_voidp) { ((void(*)(long long, long long, int, void*, int*))cppgo_virt_void_int_voidp)(__goId, -6730181576106171580LL, (int)(extension), (void*)&variant, &__h); if (__h) return; }
		QGraphicsItemGroup::setExtension(extension, variant);
	}
	QVariant extension(const QVariant& variant) const override {
		int __h = 0;
		if (__goOn[34] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 2270616358835251226LL, (void*)&variant, &__h); if (__h) return *(QVariant*)__r; }
		return QGraphicsItemGroup::extension(variant);
	}
};

extern "C" {

void* QGraphicsItemGroup_base(void* self) { return (void*)static_cast<QGraphicsItem*>((QGraphicsItemGroup*)self); }
void QGraphicsItemGroup__set_override(void* self, int idx) { ((QGraphicsItemGroup__Go*)self)->__goOn[idx] = 1; }
void* QGraphicsItemGroup_new(long long __goid) {
	return new QGraphicsItemGroup__Go(__goid);
}
void* QGraphicsItemGroup_new2(long long __goid, void* parent) {
	return new QGraphicsItemGroup__Go(__goid, (QGraphicsItem*)parent);
}
void QGraphicsItemGroup_addToGroup(void* self, void* item) {
	((QGraphicsItemGroup*)self)->addToGroup((QGraphicsItem*)item);
}
void QGraphicsItemGroup_removeFromGroup(void* self, void* item) {
	((QGraphicsItemGroup*)self)->removeFromGroup((QGraphicsItem*)item);
}
void QGraphicsItemGroup_updateMicroFocus(void* self) {
	((QGraphicsItemGroup__Go*)self)->updateMicroFocus();
}
void QGraphicsItemGroup_addToIndex(void* self) {
	((QGraphicsItemGroup__Go*)self)->addToIndex();
}
void QGraphicsItemGroup_removeFromIndex(void* self) {
	((QGraphicsItemGroup__Go*)self)->removeFromIndex();
}
void QGraphicsItemGroup_prepareGeometryChange(void* self) {
	((QGraphicsItemGroup__Go*)self)->prepareGeometryChange();
}
void* QGraphicsItemGroup_boundingRect_default(void* self) {
	auto&& __r = ((QGraphicsItemGroup*)self)->QGraphicsItemGroup::boundingRect();
	return new QRectF(__r);
}
void QGraphicsItemGroup_paint_default(void* self, void* painter, void* option, void* widget) {
	((QGraphicsItemGroup*)self)->QGraphicsItemGroup::paint((QPainter*)painter, (QStyleOptionGraphicsItem*)option, (QWidget*)widget);
}
int QGraphicsItemGroup_isObscuredBy_default(void* self, void* item) {
	auto&& __r = ((QGraphicsItemGroup*)self)->QGraphicsItemGroup::isObscuredBy((QGraphicsItem*)item);
	return __r;
}
void* QGraphicsItemGroup_opaqueArea_default(void* self) {
	auto&& __r = ((QGraphicsItemGroup*)self)->QGraphicsItemGroup::opaqueArea();
	return new QPainterPath(__r);
}
int QGraphicsItemGroup_type_default(void* self) {
	auto&& __r = ((QGraphicsItemGroup*)self)->QGraphicsItemGroup::type();
	return __r;
}
void QGraphicsItemGroup_advance_default(void* self, int phase) {
	((QGraphicsItemGroup*)self)->QGraphicsItemGroup::advance(phase);
}
void* QGraphicsItemGroup_shape_default(void* self) {
	auto&& __r = ((QGraphicsItemGroup*)self)->QGraphicsItemGroup::shape();
	return new QPainterPath(__r);
}
int QGraphicsItemGroup_contains_default(void* self, void* point) {
	auto&& __r = ((QGraphicsItemGroup*)self)->QGraphicsItemGroup::contains(*(QPointF*)point);
	return __r;
}
int QGraphicsItemGroup_collidesWithItem_default(void* self, void* other, int mode) {
	auto&& __r = ((QGraphicsItemGroup*)self)->QGraphicsItemGroup::collidesWithItem((QGraphicsItem*)other, (Qt::ItemSelectionMode)(mode));
	return __r;
}
int QGraphicsItemGroup_collidesWithPath_default(void* self, void* path, int mode) {
	auto&& __r = ((QGraphicsItemGroup*)self)->QGraphicsItemGroup::collidesWithPath(*(QPainterPath*)path, (Qt::ItemSelectionMode)(mode));
	return __r;
}
int QGraphicsItemGroup_sceneEventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QGraphicsItemGroup__Go*)self)->cppgoBase_sceneEventFilter((QGraphicsItem*)watched, (QEvent*)event);
	return __r;
}
int QGraphicsItemGroup_sceneEvent_default(void* self, void* event) {
	auto&& __r = ((QGraphicsItemGroup__Go*)self)->cppgoBase_sceneEvent((QEvent*)event);
	return __r;
}
void QGraphicsItemGroup_contextMenuEvent_default(void* self, void* event) {
	((QGraphicsItemGroup__Go*)self)->cppgoBase_contextMenuEvent((QGraphicsSceneContextMenuEvent*)event);
}
void QGraphicsItemGroup_dragEnterEvent_default(void* self, void* event) {
	((QGraphicsItemGroup__Go*)self)->cppgoBase_dragEnterEvent((QGraphicsSceneDragDropEvent*)event);
}
void QGraphicsItemGroup_dragLeaveEvent_default(void* self, void* event) {
	((QGraphicsItemGroup__Go*)self)->cppgoBase_dragLeaveEvent((QGraphicsSceneDragDropEvent*)event);
}
void QGraphicsItemGroup_dragMoveEvent_default(void* self, void* event) {
	((QGraphicsItemGroup__Go*)self)->cppgoBase_dragMoveEvent((QGraphicsSceneDragDropEvent*)event);
}
void QGraphicsItemGroup_dropEvent_default(void* self, void* event) {
	((QGraphicsItemGroup__Go*)self)->cppgoBase_dropEvent((QGraphicsSceneDragDropEvent*)event);
}
void QGraphicsItemGroup_focusInEvent_default(void* self, void* event) {
	((QGraphicsItemGroup__Go*)self)->cppgoBase_focusInEvent((QFocusEvent*)event);
}
void QGraphicsItemGroup_focusOutEvent_default(void* self, void* event) {
	((QGraphicsItemGroup__Go*)self)->cppgoBase_focusOutEvent((QFocusEvent*)event);
}
void QGraphicsItemGroup_hoverEnterEvent_default(void* self, void* event) {
	((QGraphicsItemGroup__Go*)self)->cppgoBase_hoverEnterEvent((QGraphicsSceneHoverEvent*)event);
}
void QGraphicsItemGroup_hoverMoveEvent_default(void* self, void* event) {
	((QGraphicsItemGroup__Go*)self)->cppgoBase_hoverMoveEvent((QGraphicsSceneHoverEvent*)event);
}
void QGraphicsItemGroup_hoverLeaveEvent_default(void* self, void* event) {
	((QGraphicsItemGroup__Go*)self)->cppgoBase_hoverLeaveEvent((QGraphicsSceneHoverEvent*)event);
}
void QGraphicsItemGroup_keyPressEvent_default(void* self, void* event) {
	((QGraphicsItemGroup__Go*)self)->cppgoBase_keyPressEvent((QKeyEvent*)event);
}
void QGraphicsItemGroup_keyReleaseEvent_default(void* self, void* event) {
	((QGraphicsItemGroup__Go*)self)->cppgoBase_keyReleaseEvent((QKeyEvent*)event);
}
void QGraphicsItemGroup_mousePressEvent_default(void* self, void* event) {
	((QGraphicsItemGroup__Go*)self)->cppgoBase_mousePressEvent((QGraphicsSceneMouseEvent*)event);
}
void QGraphicsItemGroup_mouseMoveEvent_default(void* self, void* event) {
	((QGraphicsItemGroup__Go*)self)->cppgoBase_mouseMoveEvent((QGraphicsSceneMouseEvent*)event);
}
void QGraphicsItemGroup_mouseReleaseEvent_default(void* self, void* event) {
	((QGraphicsItemGroup__Go*)self)->cppgoBase_mouseReleaseEvent((QGraphicsSceneMouseEvent*)event);
}
void QGraphicsItemGroup_mouseDoubleClickEvent_default(void* self, void* event) {
	((QGraphicsItemGroup__Go*)self)->cppgoBase_mouseDoubleClickEvent((QGraphicsSceneMouseEvent*)event);
}
void QGraphicsItemGroup_wheelEvent_default(void* self, void* event) {
	((QGraphicsItemGroup__Go*)self)->cppgoBase_wheelEvent((QGraphicsSceneWheelEvent*)event);
}
void QGraphicsItemGroup_inputMethodEvent_default(void* self, void* event) {
	((QGraphicsItemGroup__Go*)self)->cppgoBase_inputMethodEvent((QInputMethodEvent*)event);
}
void* QGraphicsItemGroup_inputMethodQuery_default(void* self, int query) {
	auto&& __r = ((QGraphicsItemGroup__Go*)self)->cppgoBase_inputMethodQuery(query);
	return new QVariant(__r);
}
void* QGraphicsItemGroup_itemChange_default(void* self, int change, void* value) {
	auto&& __r = ((QGraphicsItemGroup__Go*)self)->cppgoBase_itemChange(change, *(QVariant*)value);
	return new QVariant(__r);
}
int QGraphicsItemGroup_supportsExtension_default(void* self, int extension) {
	auto&& __r = ((QGraphicsItemGroup__Go*)self)->cppgoBase_supportsExtension(extension);
	return __r;
}
void QGraphicsItemGroup_setExtension_default(void* self, int extension, void* variant) {
	((QGraphicsItemGroup__Go*)self)->cppgoBase_setExtension(extension, *(QVariant*)variant);
}
void* QGraphicsItemGroup_extension_default(void* self, void* variant) {
	auto&& __r = ((QGraphicsItemGroup__Go*)self)->cppgoBase_extension(*(QVariant*)variant);
	return new QVariant(__r);
}
void* QGraphicsItemGroup__Go_new(long long id) { return new QGraphicsItemGroup__Go(id); }
void QGraphicsItemGroup_delete(void* self) { delete (QGraphicsItemGroup*)self; }

}

#include <QtWidgets/qgraphicssceneevent.h>
#include <QtCore/qpoint.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qcoreevent.h>
#include "gen_qgraphicsscenemoveevent.h"
#include "gen_purego_cb.h"

class QGraphicsSceneMoveEvent__Go : public QGraphicsSceneMoveEvent {
public:
	long long __goId;
	unsigned char __goOn[2] = {};
	template<typename... A> QGraphicsSceneMoveEvent__Go(long long id, A&&... a) : QGraphicsSceneMoveEvent(std::forward<A>(a)...), __goId(id) {}
	void setAccepted(bool accepted) override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 6367896784077321574LL, accepted, &__h); if (__h) return; }
		QGraphicsSceneMoveEvent::setAccepted(accepted);
	}
	QEvent* clone() const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, -7032813281048616478LL, &__h); if (__h) return (QEvent*)__r; }
		return QGraphicsSceneMoveEvent::clone();
	}
};

extern "C" {

void* QGraphicsSceneMoveEvent_base(void* self) { return (void*)static_cast<QGraphicsSceneEvent*>((QGraphicsSceneMoveEvent*)self); }
void QGraphicsSceneMoveEvent_chain(void* self, void** out) {
	QGraphicsSceneMoveEvent* __p = (QGraphicsSceneMoveEvent*)self;
	QGraphicsSceneEvent* __b0 = static_cast<QGraphicsSceneEvent*>(__p); out[0] = (void*)__b0;
	QEvent* __b1 = static_cast<QEvent*>(__b0); out[1] = (void*)__b1;
}
void QGraphicsSceneMoveEvent__set_override(void* self, int idx) { ((QGraphicsSceneMoveEvent__Go*)self)->__goOn[idx] = 1; }
void* QGraphicsSceneMoveEvent_new(long long __goid) {
	return new QGraphicsSceneMoveEvent__Go(__goid);
}
void* QGraphicsSceneMoveEvent_oldPos(void* self) {
	auto&& __r = ((QGraphicsSceneMoveEvent*)self)->oldPos();
	return new QPointF(__r);
}
void QGraphicsSceneMoveEvent_setOldPos(void* self, void* pos) {
	((QGraphicsSceneMoveEvent*)self)->setOldPos(*(QPointF*)pos);
}
void* QGraphicsSceneMoveEvent_newPos(void* self) {
	auto&& __r = ((QGraphicsSceneMoveEvent*)self)->newPos();
	return new QPointF(__r);
}
void QGraphicsSceneMoveEvent_setNewPos(void* self, void* pos) {
	((QGraphicsSceneMoveEvent*)self)->setNewPos(*(QPointF*)pos);
}
int QGraphicsSceneMoveEvent_registerEventType() {
	auto&& __r = QGraphicsSceneMoveEvent::registerEventType();
	return __r;
}
int QGraphicsSceneMoveEvent_registerEventType2(int hint) {
	auto&& __r = QGraphicsSceneMoveEvent::registerEventType(hint);
	return __r;
}
void QGraphicsSceneMoveEvent_setAccepted_default(void* self, int accepted) {
	((QGraphicsSceneMoveEvent*)self)->QGraphicsSceneMoveEvent::setAccepted(accepted);
}
void* QGraphicsSceneMoveEvent_clone_default(void* self) {
	auto&& __r = ((QGraphicsSceneMoveEvent*)self)->QGraphicsSceneMoveEvent::clone();
	return (void*)(__r);
}
void* QGraphicsSceneMoveEvent__Go_new(long long id) { return new QGraphicsSceneMoveEvent__Go(id); }
void QGraphicsSceneMoveEvent_delete(void* self) { delete (QGraphicsSceneMoveEvent*)self; }

}

#include <QtWidgets/qmessagebox.h>
#include <QtWidgets/qdialog.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qstring.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qobjectdefs.h>
#include <QtWidgets/qabstractbutton.h>
#include <QtWidgets/qpushbutton.h>
#include <QtCore/qobject.h>
#include <QtGui/qpixmap.h>
#include <QtWidgets/qcheckbox.h>
#include <QtCore/qcoreevent.h>
#include <QtGui/qevent.h>
#include <QtCore/qsize.h>
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
#include <QtGui/qpaintengine.h>
#include <QtGui/qbackingstore.h>
#include <QtGui/qwindow.h>
#include <QtGui/qscreen.h>
#include <QtCore/qvariant.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qmessagebox.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QMessageBox__Go : public QMessageBox {
public:
	long long __goId;
	unsigned char __goOn[52] = {};
	template<typename... A> QMessageBox__Go(long long id, A&&... a) : QMessageBox(std::forward<A>(a)...), __goId(id) {}
	using QMessageBox::adjustPosition;
	using QMessageBox::updateMicroFocus;
	using QMessageBox::create;
	using QMessageBox::destroy;
	using QMessageBox::focusNextChild;
	using QMessageBox::focusPreviousChild;
	using QMessageBox::sender;
	using QMessageBox::senderSignalIndex;
	using QMessageBox::receivers;
	using QMessageBox::isSignalConnected;
	using QMessageBox::getDecodedMetricF;
	bool cppgoBase_event(QEvent* __a0) { return QMessageBox::event(__a0); }
	void cppgoBase_resizeEvent(QResizeEvent* __a0) { QMessageBox::resizeEvent(__a0); }
	void cppgoBase_showEvent(QShowEvent* __a0) { QMessageBox::showEvent(__a0); }
	void cppgoBase_closeEvent(QCloseEvent* __a0) { QMessageBox::closeEvent(__a0); }
	void cppgoBase_keyPressEvent(QKeyEvent* __a0) { QMessageBox::keyPressEvent(__a0); }
	void cppgoBase_changeEvent(QEvent* __a0) { QMessageBox::changeEvent(__a0); }
	void cppgoBase_contextMenuEvent(QContextMenuEvent* __a0) { QMessageBox::contextMenuEvent(__a0); }
	bool cppgoBase_eventFilter(QObject* __a0, QEvent* __a1) { return QMessageBox::eventFilter(__a0, __a1); }
	void cppgoBase_mousePressEvent(QMouseEvent* __a0) { QMessageBox::mousePressEvent(__a0); }
	void cppgoBase_mouseReleaseEvent(QMouseEvent* __a0) { QMessageBox::mouseReleaseEvent(__a0); }
	void cppgoBase_mouseDoubleClickEvent(QMouseEvent* __a0) { QMessageBox::mouseDoubleClickEvent(__a0); }
	void cppgoBase_mouseMoveEvent(QMouseEvent* __a0) { QMessageBox::mouseMoveEvent(__a0); }
	void cppgoBase_wheelEvent(QWheelEvent* __a0) { QMessageBox::wheelEvent(__a0); }
	void cppgoBase_keyReleaseEvent(QKeyEvent* __a0) { QMessageBox::keyReleaseEvent(__a0); }
	void cppgoBase_focusInEvent(QFocusEvent* __a0) { QMessageBox::focusInEvent(__a0); }
	void cppgoBase_focusOutEvent(QFocusEvent* __a0) { QMessageBox::focusOutEvent(__a0); }
	void cppgoBase_enterEvent(QEnterEvent* __a0) { QMessageBox::enterEvent(__a0); }
	void cppgoBase_leaveEvent(QEvent* __a0) { QMessageBox::leaveEvent(__a0); }
	void cppgoBase_paintEvent(QPaintEvent* __a0) { QMessageBox::paintEvent(__a0); }
	void cppgoBase_moveEvent(QMoveEvent* __a0) { QMessageBox::moveEvent(__a0); }
	void cppgoBase_tabletEvent(QTabletEvent* __a0) { QMessageBox::tabletEvent(__a0); }
	void cppgoBase_actionEvent(QActionEvent* __a0) { QMessageBox::actionEvent(__a0); }
	void cppgoBase_dragEnterEvent(QDragEnterEvent* __a0) { QMessageBox::dragEnterEvent(__a0); }
	void cppgoBase_dragMoveEvent(QDragMoveEvent* __a0) { QMessageBox::dragMoveEvent(__a0); }
	void cppgoBase_dragLeaveEvent(QDragLeaveEvent* __a0) { QMessageBox::dragLeaveEvent(__a0); }
	void cppgoBase_dropEvent(QDropEvent* __a0) { QMessageBox::dropEvent(__a0); }
	void cppgoBase_hideEvent(QHideEvent* __a0) { QMessageBox::hideEvent(__a0); }
	int cppgoBase_metric(int __a0) { return QMessageBox::metric((QPaintDevice::PaintDeviceMetric)__a0); }
	void cppgoBase_initPainter(QPainter* __a0) { QMessageBox::initPainter(__a0); }
	QPaintDevice* cppgoBase_redirected(QPoint* __a0) { return QMessageBox::redirected(__a0); }
	QPainter* cppgoBase_sharedPainter() { return QMessageBox::sharedPainter(); }
	void cppgoBase_inputMethodEvent(QInputMethodEvent* __a0) { QMessageBox::inputMethodEvent(__a0); }
	bool cppgoBase_focusNextPrevChild(bool __a0) { return QMessageBox::focusNextPrevChild(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QMessageBox::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QMessageBox::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QMessageBox::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QMessageBox::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QMessageBox::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QMessageBox::metaObject();
	}
	bool event(QEvent* e) override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)e, &__h); if (__h) return __r; }
		return QMessageBox::event(e);
	}
	void resizeEvent(QResizeEvent* event) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4302917285655014661LL, (void*)event, &__h); if (__h) return; }
		QMessageBox::resizeEvent(event);
	}
	void showEvent(QShowEvent* event) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6825369174751916LL, (void*)event, &__h); if (__h) return; }
		QMessageBox::showEvent(event);
	}
	void closeEvent(QCloseEvent* event) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1989937050019464783LL, (void*)event, &__h); if (__h) return; }
		QMessageBox::closeEvent(event);
	}
	void keyPressEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1023791977679793185LL, (void*)event, &__h); if (__h) return; }
		QMessageBox::keyPressEvent(event);
	}
	void changeEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4959376063826165943LL, (void*)event, &__h); if (__h) return; }
		QMessageBox::changeEvent(event);
	}
	void setVisible(bool visible) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 2110182719375245575LL, visible, &__h); if (__h) return; }
		QMessageBox::setVisible(visible);
	}
	QSize sizeHint() const override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 729690261457684101LL, &__h); if (__h) return *(QSize*)__r; }
		return QMessageBox::sizeHint();
	}
	QSize minimumSizeHint() const override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 7049264599837532491LL, &__h); if (__h) return *(QSize*)__r; }
		return QMessageBox::minimumSizeHint();
	}
	int exec() override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -1745322561779819684LL, &__h); if (__h) return __r; }
		return QMessageBox::exec();
	}
	void done(int arg1) override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, -2570997490807640431LL, arg1, &__h); if (__h) return; }
		QMessageBox::done(arg1);
	}
	void accept() override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 9106131928185084681LL, &__h); if (__h) return; }
		QMessageBox::accept();
	}
	void reject() override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 2875220963396170150LL, &__h); if (__h) return; }
		QMessageBox::reject();
	}
	void contextMenuEvent(QContextMenuEvent* arg1) override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6698043222227006015LL, (void*)arg1, &__h); if (__h) return; }
		QMessageBox::contextMenuEvent(arg1);
	}
	bool eventFilter(QObject* arg1, QEvent* arg2) override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)arg1, (void*)arg2, &__h); if (__h) return __r; }
		return QMessageBox::eventFilter(arg1, arg2);
	}
	int devType() const override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8659679893868524044LL, &__h); if (__h) return __r; }
		return QMessageBox::devType();
	}
	int heightForWidth(int arg1) const override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 4311534060495737339LL, arg1, &__h); if (__h) return __r; }
		return QMessageBox::heightForWidth(arg1);
	}
	bool hasHeightForWidth() const override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8266762397995423915LL, &__h); if (__h) return __r; }
		return QMessageBox::hasHeightForWidth();
	}
	QPaintEngine* paintEngine() const override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 6398007404756346381LL, &__h); if (__h) return (QPaintEngine*)__r; }
		return QMessageBox::paintEngine();
	}
	void mousePressEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -332090898078543319LL, (void*)event, &__h); if (__h) return; }
		QMessageBox::mousePressEvent(event);
	}
	void mouseReleaseEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4415920049297795213LL, (void*)event, &__h); if (__h) return; }
		QMessageBox::mouseReleaseEvent(event);
	}
	void mouseDoubleClickEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4029073905697023887LL, (void*)event, &__h); if (__h) return; }
		QMessageBox::mouseDoubleClickEvent(event);
	}
	void mouseMoveEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4255494004099806581LL, (void*)event, &__h); if (__h) return; }
		QMessageBox::mouseMoveEvent(event);
	}
	void wheelEvent(QWheelEvent* event) override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1919988340934511316LL, (void*)event, &__h); if (__h) return; }
		QMessageBox::wheelEvent(event);
	}
	void keyReleaseEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 772598475950050299LL, (void*)event, &__h); if (__h) return; }
		QMessageBox::keyReleaseEvent(event);
	}
	void focusInEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 975043767033552552LL, (void*)event, &__h); if (__h) return; }
		QMessageBox::focusInEvent(event);
	}
	void focusOutEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -385596010576048993LL, (void*)event, &__h); if (__h) return; }
		QMessageBox::focusOutEvent(event);
	}
	void enterEvent(QEnterEvent* event) override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6316270346668237895LL, (void*)event, &__h); if (__h) return; }
		QMessageBox::enterEvent(event);
	}
	void leaveEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -551070009623711640LL, (void*)event, &__h); if (__h) return; }
		QMessageBox::leaveEvent(event);
	}
	void paintEvent(QPaintEvent* event) override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5290442163167458465LL, (void*)event, &__h); if (__h) return; }
		QMessageBox::paintEvent(event);
	}
	void moveEvent(QMoveEvent* event) override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -9161949150489525396LL, (void*)event, &__h); if (__h) return; }
		QMessageBox::moveEvent(event);
	}
	void tabletEvent(QTabletEvent* event) override {
		int __h = 0;
		if (__goOn[32] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -212068866860701141LL, (void*)event, &__h); if (__h) return; }
		QMessageBox::tabletEvent(event);
	}
	void actionEvent(QActionEvent* event) override {
		int __h = 0;
		if (__goOn[33] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1697444140088555581LL, (void*)event, &__h); if (__h) return; }
		QMessageBox::actionEvent(event);
	}
	void dragEnterEvent(QDragEnterEvent* event) override {
		int __h = 0;
		if (__goOn[34] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 206122715919917589LL, (void*)event, &__h); if (__h) return; }
		QMessageBox::dragEnterEvent(event);
	}
	void dragMoveEvent(QDragMoveEvent* event) override {
		int __h = 0;
		if (__goOn[35] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1781545137352745978LL, (void*)event, &__h); if (__h) return; }
		QMessageBox::dragMoveEvent(event);
	}
	void dragLeaveEvent(QDragLeaveEvent* event) override {
		int __h = 0;
		if (__goOn[36] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -3139904158373725006LL, (void*)event, &__h); if (__h) return; }
		QMessageBox::dragLeaveEvent(event);
	}
	void dropEvent(QDropEvent* event) override {
		int __h = 0;
		if (__goOn[37] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 9107058682754591716LL, (void*)event, &__h); if (__h) return; }
		QMessageBox::dropEvent(event);
	}
	void hideEvent(QHideEvent* event) override {
		int __h = 0;
		if (__goOn[38] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -199718707876719129LL, (void*)event, &__h); if (__h) return; }
		QMessageBox::hideEvent(event);
	}
	bool nativeEvent(const QByteArray& eventType, void* message, qintptr* result) override {
		int __h = 0;
		if (__goOn[39] && cppgo_virt_int_voidp_voidp_longlongp) { int __r = ((int(*)(long long, long long, void*, void*, long long*, int*))cppgo_virt_int_voidp_voidp_longlongp)(__goId, -2039166586246443706LL, (void*)&eventType, (void*)message, (long long*)result, &__h); if (__h) return __r; }
		return QMessageBox::nativeEvent(eventType, message, result);
	}
	int metric(QPaintDevice::PaintDeviceMetric arg1) const override {
		int __h = 0;
		if (__goOn[40] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 427629649267311777LL, (int)(arg1), &__h); if (__h) return __r; }
		return QMessageBox::metric(arg1);
	}
	void initPainter(QPainter* painter) const override {
		int __h = 0;
		if (__goOn[41] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6477500296980050210LL, (void*)painter, &__h); if (__h) return; }
		QMessageBox::initPainter(painter);
	}
	QPaintDevice* redirected(QPoint* offset) const override {
		int __h = 0;
		if (__goOn[42] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 4127240633310655048LL, (void*)offset, &__h); if (__h) return (QPaintDevice*)__r; }
		return QMessageBox::redirected(offset);
	}
	QPainter* sharedPainter() const override {
		int __h = 0;
		if (__goOn[43] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 5792771328741774991LL, &__h); if (__h) return (QPainter*)__r; }
		return QMessageBox::sharedPainter();
	}
	void inputMethodEvent(QInputMethodEvent* arg1) override {
		int __h = 0;
		if (__goOn[44] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5438819661942843222LL, (void*)arg1, &__h); if (__h) return; }
		QMessageBox::inputMethodEvent(arg1);
	}
	QVariant inputMethodQuery(Qt::InputMethodQuery arg1) const override {
		int __h = 0;
		if (__goOn[45] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7345853097340709638LL, (int)(arg1), &__h); if (__h) return *(QVariant*)__r; }
		return QMessageBox::inputMethodQuery(arg1);
	}
	bool focusNextPrevChild(bool next) override {
		int __h = 0;
		if (__goOn[46] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 5898328517505784191LL, next, &__h); if (__h) return __r; }
		return QMessageBox::focusNextPrevChild(next);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[47] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QMessageBox::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[48] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QMessageBox::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[49] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QMessageBox::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[50] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QMessageBox::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[51] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QMessageBox::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QMessageBox;
static std::mutex __connsMu_QMessageBox;

extern "C" {

void* QMessageBox_base(void* self) { return (void*)static_cast<QDialog*>((QMessageBox*)self); }
void QMessageBox_chain(void* self, void** out) {
	QMessageBox* __p = (QMessageBox*)self;
	QDialog* __b0 = static_cast<QDialog*>(__p); out[0] = (void*)__b0;
	QWidget* __b1 = static_cast<QWidget*>(__b0); out[1] = (void*)__b1;
	QObject* __b2 = static_cast<QObject*>(__b1); out[2] = (void*)__b2;
}
void QMessageBox__set_override(void* self, int idx) { ((QMessageBox__Go*)self)->__goOn[idx] = 1; }
void* QMessageBox_fromQObject(void* o) { return (void*)qobject_cast<QMessageBox*>((QObject*)o); }
void* QMessageBox_new(long long __goid, void* parent) {
	return new QMessageBox__Go(__goid, (QWidget*)parent);
}
void* QMessageBox_new2(long long __goid) {
	return new QMessageBox__Go(__goid);
}
void* QMessageBox_new3(long long __goid, int icon, const char* title, long long title_len, const char* text, long long text_len) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	QString text_qstr = QString::fromUtf8(text, text_len);
	return new QMessageBox__Go(__goid, (QMessageBox::Icon)(icon), title_qstr, text_qstr);
}
void* QMessageBox_new4(long long __goid, const char* title, long long title_len, const char* text, long long text_len, int icon, int button0, int button1, int button2) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	QString text_qstr = QString::fromUtf8(text, text_len);
	return new QMessageBox__Go(__goid, title_qstr, text_qstr, (QMessageBox::Icon)(icon), button0, button1, button2);
}
void* QMessageBox_new5(long long __goid, int icon, const char* title, long long title_len, const char* text, long long text_len, int buttons) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	QString text_qstr = QString::fromUtf8(text, text_len);
	return new QMessageBox__Go(__goid, (QMessageBox::Icon)(icon), title_qstr, text_qstr, QFlags<QMessageBox::StandardButton>(buttons));
}
void* QMessageBox_new6(long long __goid, int icon, const char* title, long long title_len, const char* text, long long text_len, int buttons, void* parent) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	QString text_qstr = QString::fromUtf8(text, text_len);
	return new QMessageBox__Go(__goid, (QMessageBox::Icon)(icon), title_qstr, text_qstr, QFlags<QMessageBox::StandardButton>(buttons), (QWidget*)parent);
}
void* QMessageBox_new7(long long __goid, int icon, const char* title, long long title_len, const char* text, long long text_len, int buttons, void* parent, int flags) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	QString text_qstr = QString::fromUtf8(text, text_len);
	return new QMessageBox__Go(__goid, (QMessageBox::Icon)(icon), title_qstr, text_qstr, QFlags<QMessageBox::StandardButton>(buttons), (QWidget*)parent, QFlags<Qt::WindowType>(flags));
}
void* QMessageBox_new8(long long __goid, const char* title, long long title_len, const char* text, long long text_len, int icon, int button0, int button1, int button2, void* parent) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	QString text_qstr = QString::fromUtf8(text, text_len);
	return new QMessageBox__Go(__goid, title_qstr, text_qstr, (QMessageBox::Icon)(icon), button0, button1, button2, (QWidget*)parent);
}
void* QMessageBox_new9(long long __goid, const char* title, long long title_len, const char* text, long long text_len, int icon, int button0, int button1, int button2, void* parent, int f) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	QString text_qstr = QString::fromUtf8(text, text_len);
	return new QMessageBox__Go(__goid, title_qstr, text_qstr, (QMessageBox::Icon)(icon), button0, button1, button2, (QWidget*)parent, QFlags<Qt::WindowType>(f));
}
char* QMessageBox_tr(const char* s, long long* __out_len) {
	auto&& __r = QMessageBox::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QMessageBox_addButton(void* self, void* button, int role) {
	((QMessageBox*)self)->addButton((QAbstractButton*)button, (QMessageBox::ButtonRole)(role));
}
void* QMessageBox_addButton2(void* self, const char* text, long long text_len, int role) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = ((QMessageBox*)self)->addButton(text_qstr, (QMessageBox::ButtonRole)(role));
	return (void*)(__r);
}
void* QMessageBox_addButton3(void* self, int button) {
	auto&& __r = ((QMessageBox*)self)->addButton((QMessageBox::StandardButton)(button));
	return (void*)(__r);
}
void QMessageBox_removeButton(void* self, void* button) {
	((QMessageBox*)self)->removeButton((QAbstractButton*)button);
}
void QMessageBox_open(void* self, void* receiver, const char* member) {
	((QMessageBox*)self)->open((QObject*)receiver, member);
}
void** QMessageBox_buttons(void* self, long long* __out_len) {
	auto&& __r = ((QMessageBox*)self)->buttons();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
int QMessageBox_buttonRole(void* self, void* button) {
	auto&& __r = ((QMessageBox*)self)->buttonRole((QAbstractButton*)button);
	return (int)(__r);
}
void QMessageBox_setStandardButtons(void* self, int buttons) {
	((QMessageBox*)self)->setStandardButtons(QFlags<QMessageBox::StandardButton>(buttons));
}
int QMessageBox_standardButtons(void* self) {
	auto&& __r = ((QMessageBox*)self)->standardButtons();
	return (int)(__r);
}
int QMessageBox_standardButton(void* self, void* button) {
	auto&& __r = ((QMessageBox*)self)->standardButton((QAbstractButton*)button);
	return (int)(__r);
}
void* QMessageBox_button(void* self, int which) {
	auto&& __r = ((QMessageBox*)self)->button((QMessageBox::StandardButton)(which));
	return (void*)(__r);
}
void* QMessageBox_defaultButton(void* self) {
	auto&& __r = ((QMessageBox*)self)->defaultButton();
	return (void*)(__r);
}
void QMessageBox_setDefaultButton(void* self, void* button) {
	((QMessageBox*)self)->setDefaultButton((QPushButton*)button);
}
void QMessageBox_setDefaultButton2(void* self, int button) {
	((QMessageBox*)self)->setDefaultButton((QMessageBox::StandardButton)(button));
}
void* QMessageBox_escapeButton(void* self) {
	auto&& __r = ((QMessageBox*)self)->escapeButton();
	return (void*)(__r);
}
void QMessageBox_setEscapeButton(void* self, void* button) {
	((QMessageBox*)self)->setEscapeButton((QAbstractButton*)button);
}
void QMessageBox_setEscapeButton2(void* self, int button) {
	((QMessageBox*)self)->setEscapeButton((QMessageBox::StandardButton)(button));
}
void* QMessageBox_clickedButton(void* self) {
	auto&& __r = ((QMessageBox*)self)->clickedButton();
	return (void*)(__r);
}
char* QMessageBox_text(void* self, long long* __out_len) {
	auto&& __r = ((QMessageBox*)self)->text();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QMessageBox_setText(void* self, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	((QMessageBox*)self)->setText(text_qstr);
}
int QMessageBox_icon(void* self) {
	auto&& __r = ((QMessageBox*)self)->icon();
	return (int)(__r);
}
void QMessageBox_setIcon(void* self, int arg1) {
	((QMessageBox*)self)->setIcon((QMessageBox::Icon)(arg1));
}
void* QMessageBox_iconPixmap(void* self) {
	auto&& __r = ((QMessageBox*)self)->iconPixmap();
	return new QPixmap(__r);
}
void QMessageBox_setIconPixmap(void* self, void* pixmap) {
	((QMessageBox*)self)->setIconPixmap(*(QPixmap*)pixmap);
}
int QMessageBox_textFormat(void* self) {
	auto&& __r = ((QMessageBox*)self)->textFormat();
	return (int)(__r);
}
void QMessageBox_setTextFormat(void* self, int format) {
	((QMessageBox*)self)->setTextFormat((Qt::TextFormat)(format));
}
void QMessageBox_setTextInteractionFlags(void* self, int flags) {
	((QMessageBox*)self)->setTextInteractionFlags(QFlags<Qt::TextInteractionFlag>(flags));
}
int QMessageBox_textInteractionFlags(void* self) {
	auto&& __r = ((QMessageBox*)self)->textInteractionFlags();
	return (int)(__r);
}
void QMessageBox_setCheckBox(void* self, void* cb) {
	((QMessageBox*)self)->setCheckBox((QCheckBox*)cb);
}
void* QMessageBox_checkBox(void* self) {
	auto&& __r = ((QMessageBox*)self)->checkBox();
	return (void*)(__r);
}
void QMessageBox_setOption(void* self, int option) {
	((QMessageBox*)self)->setOption((QMessageBox::Option)(option));
}
int QMessageBox_testOption(void* self, int option) {
	auto&& __r = ((QMessageBox*)self)->testOption((QMessageBox::Option)(option));
	return __r;
}
void QMessageBox_setOptions(void* self, int options) {
	((QMessageBox*)self)->setOptions(QFlags<QMessageBox::Option>(options));
}
int QMessageBox_options(void* self) {
	auto&& __r = ((QMessageBox*)self)->options();
	return (int)(__r);
}
int QMessageBox_information(void* parent, const char* title, long long title_len, const char* text, long long text_len) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = QMessageBox::information((QWidget*)parent, title_qstr, text_qstr);
	return (int)(__r);
}
int QMessageBox_information2(void* parent, const char* title, long long title_len, const char* text, long long text_len, int button0) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = QMessageBox::information((QWidget*)parent, title_qstr, text_qstr, (QMessageBox::StandardButton)(button0));
	return (int)(__r);
}
int QMessageBox_question(void* parent, const char* title, long long title_len, const char* text, long long text_len) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = QMessageBox::question((QWidget*)parent, title_qstr, text_qstr);
	return (int)(__r);
}
int QMessageBox_question2(void* parent, const char* title, long long title_len, const char* text, long long text_len, int button0, int button1) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = QMessageBox::question((QWidget*)parent, title_qstr, text_qstr, (QMessageBox::StandardButton)(button0), (QMessageBox::StandardButton)(button1));
	return __r;
}
int QMessageBox_warning(void* parent, const char* title, long long title_len, const char* text, long long text_len) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = QMessageBox::warning((QWidget*)parent, title_qstr, text_qstr);
	return (int)(__r);
}
int QMessageBox_warning2(void* parent, const char* title, long long title_len, const char* text, long long text_len, int button0, int button1) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = QMessageBox::warning((QWidget*)parent, title_qstr, text_qstr, (QMessageBox::StandardButton)(button0), (QMessageBox::StandardButton)(button1));
	return __r;
}
int QMessageBox_critical(void* parent, const char* title, long long title_len, const char* text, long long text_len) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = QMessageBox::critical((QWidget*)parent, title_qstr, text_qstr);
	return (int)(__r);
}
int QMessageBox_critical2(void* parent, const char* title, long long title_len, const char* text, long long text_len, int button0, int button1) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = QMessageBox::critical((QWidget*)parent, title_qstr, text_qstr, (QMessageBox::StandardButton)(button0), (QMessageBox::StandardButton)(button1));
	return __r;
}
void QMessageBox_about(void* parent, const char* title, long long title_len, const char* text, long long text_len) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	QString text_qstr = QString::fromUtf8(text, text_len);
	QMessageBox::about((QWidget*)parent, title_qstr, text_qstr);
}
void QMessageBox_aboutQt(void* parent) {
	QMessageBox::aboutQt((QWidget*)parent);
}
int QMessageBox_information3(void* parent, const char* title, long long title_len, const char* text, long long text_len, int button0) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = QMessageBox::information((QWidget*)parent, title_qstr, text_qstr, button0);
	return __r;
}
int QMessageBox_information4(void* parent, const char* title, long long title_len, const char* text, long long text_len, const char* button0Text, long long button0Text_len) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	QString text_qstr = QString::fromUtf8(text, text_len);
	QString button0Text_qstr = QString::fromUtf8(button0Text, button0Text_len);
	auto&& __r = QMessageBox::information((QWidget*)parent, title_qstr, text_qstr, button0Text_qstr);
	return __r;
}
int QMessageBox_question3(void* parent, const char* title, long long title_len, const char* text, long long text_len, int button0) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = QMessageBox::question((QWidget*)parent, title_qstr, text_qstr, button0);
	return __r;
}
int QMessageBox_question4(void* parent, const char* title, long long title_len, const char* text, long long text_len, const char* button0Text, long long button0Text_len) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	QString text_qstr = QString::fromUtf8(text, text_len);
	QString button0Text_qstr = QString::fromUtf8(button0Text, button0Text_len);
	auto&& __r = QMessageBox::question((QWidget*)parent, title_qstr, text_qstr, button0Text_qstr);
	return __r;
}
int QMessageBox_warning3(void* parent, const char* title, long long title_len, const char* text, long long text_len, int button0, int button1) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = QMessageBox::warning((QWidget*)parent, title_qstr, text_qstr, button0, button1);
	return __r;
}
int QMessageBox_warning4(void* parent, const char* title, long long title_len, const char* text, long long text_len, const char* button0Text, long long button0Text_len) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	QString text_qstr = QString::fromUtf8(text, text_len);
	QString button0Text_qstr = QString::fromUtf8(button0Text, button0Text_len);
	auto&& __r = QMessageBox::warning((QWidget*)parent, title_qstr, text_qstr, button0Text_qstr);
	return __r;
}
int QMessageBox_critical3(void* parent, const char* title, long long title_len, const char* text, long long text_len, int button0, int button1) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = QMessageBox::critical((QWidget*)parent, title_qstr, text_qstr, button0, button1);
	return __r;
}
int QMessageBox_critical4(void* parent, const char* title, long long title_len, const char* text, long long text_len, const char* button0Text, long long button0Text_len) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	QString text_qstr = QString::fromUtf8(text, text_len);
	QString button0Text_qstr = QString::fromUtf8(button0Text, button0Text_len);
	auto&& __r = QMessageBox::critical((QWidget*)parent, title_qstr, text_qstr, button0Text_qstr);
	return __r;
}
char* QMessageBox_buttonText(void* self, int button, long long* __out_len) {
	auto&& __r = ((QMessageBox*)self)->buttonText(button);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QMessageBox_setButtonText(void* self, int button, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	((QMessageBox*)self)->setButtonText(button, text_qstr);
}
char* QMessageBox_informativeText(void* self, long long* __out_len) {
	auto&& __r = ((QMessageBox*)self)->informativeText();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QMessageBox_setInformativeText(void* self, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	((QMessageBox*)self)->setInformativeText(text_qstr);
}
char* QMessageBox_detailedText(void* self, long long* __out_len) {
	auto&& __r = ((QMessageBox*)self)->detailedText();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QMessageBox_setDetailedText(void* self, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	((QMessageBox*)self)->setDetailedText(text_qstr);
}
void QMessageBox_setWindowTitle(void* self, const char* title, long long title_len) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	((QMessageBox*)self)->setWindowTitle(title_qstr);
}
void QMessageBox_setWindowModality(void* self, int windowModality) {
	((QMessageBox*)self)->setWindowModality((Qt::WindowModality)(windowModality));
}
void* QMessageBox_standardIcon(int icon) {
	auto&& __r = QMessageBox::standardIcon((QMessageBox::Icon)(icon));
	return new QPixmap(__r);
}
void QMessageBox_buttonClicked(void* self, void* button) {
	((QMessageBox*)self)->buttonClicked((QAbstractButton*)button);
}
void QMessageBox_adjustPosition(void* self, void* arg1) {
	((QMessageBox__Go*)self)->adjustPosition((QWidget*)arg1);
}
void QMessageBox_setTabOrder(void* arg1, void* arg2) {
	QMessageBox::setTabOrder((QWidget*)arg1, (QWidget*)arg2);
}
void* QMessageBox_mouseGrabber() {
	auto&& __r = QMessageBox::mouseGrabber();
	return (void*)(__r);
}
void* QMessageBox_keyboardGrabber() {
	auto&& __r = QMessageBox::keyboardGrabber();
	return (void*)(__r);
}
void* QMessageBox_find(unsigned long long arg1) {
	auto&& __r = QMessageBox::find(arg1);
	return (void*)(__r);
}
void* QMessageBox_createWindowContainer(void* window) {
	auto&& __r = QMessageBox::createWindowContainer((QWindow*)window);
	return (void*)(__r);
}
void QMessageBox_updateMicroFocus(void* self) {
	((QMessageBox__Go*)self)->updateMicroFocus();
}
void QMessageBox_create(void* self) {
	((QMessageBox__Go*)self)->create();
}
void QMessageBox_destroy(void* self) {
	((QMessageBox__Go*)self)->destroy();
}
int QMessageBox_focusNextChild(void* self) {
	auto&& __r = ((QMessageBox__Go*)self)->focusNextChild();
	return __r;
}
int QMessageBox_focusPreviousChild(void* self) {
	auto&& __r = ((QMessageBox__Go*)self)->focusPreviousChild();
	return __r;
}
void* QMessageBox_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QMessageBox::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QMessageBox_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QMessageBox::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QMessageBox_sender(void* self) {
	auto&& __r = ((QMessageBox__Go*)self)->sender();
	return (void*)(__r);
}
int QMessageBox_senderSignalIndex(void* self) {
	auto&& __r = ((QMessageBox__Go*)self)->senderSignalIndex();
	return __r;
}
int QMessageBox_receivers(void* self, const char* signal) {
	auto&& __r = ((QMessageBox__Go*)self)->receivers(signal);
	return __r;
}
int QMessageBox_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QMessageBox__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
double QMessageBox_devicePixelRatioFScale() {
	auto&& __r = QMessageBox::devicePixelRatioFScale();
	return __r;
}
int QMessageBox_encodeMetricF(int metric, double value) {
	auto&& __r = QMessageBox::encodeMetricF((QPaintDevice::PaintDeviceMetric)(metric), value);
	return __r;
}
double QMessageBox_getDecodedMetricF(void* self, int metricA, int metricB) {
	auto&& __r = ((QMessageBox__Go*)self)->getDecodedMetricF((QPaintDevice::PaintDeviceMetric)(metricA), (QPaintDevice::PaintDeviceMetric)(metricB));
	return __r;
}
char* QMessageBox_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QMessageBox::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QMessageBox_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QMessageBox::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QMessageBox_setOption2(void* self, int option, int on) {
	((QMessageBox*)self)->setOption((QMessageBox::Option)(option), on);
}
int QMessageBox_information5(void* parent, const char* title, long long title_len, const char* text, long long text_len, int buttons) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = QMessageBox::information((QWidget*)parent, title_qstr, text_qstr, QFlags<QMessageBox::StandardButton>(buttons));
	return (int)(__r);
}
int QMessageBox_information6(void* parent, const char* title, long long title_len, const char* text, long long text_len, int buttons, int defaultButton) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = QMessageBox::information((QWidget*)parent, title_qstr, text_qstr, QFlags<QMessageBox::StandardButton>(buttons), (QMessageBox::StandardButton)(defaultButton));
	return (int)(__r);
}
int QMessageBox_information7(void* parent, const char* title, long long title_len, const char* text, long long text_len, int button0, int button1) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = QMessageBox::information((QWidget*)parent, title_qstr, text_qstr, (QMessageBox::StandardButton)(button0), (QMessageBox::StandardButton)(button1));
	return (int)(__r);
}
int QMessageBox_question5(void* parent, const char* title, long long title_len, const char* text, long long text_len, int buttons) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = QMessageBox::question((QWidget*)parent, title_qstr, text_qstr, QFlags<QMessageBox::StandardButton>(buttons));
	return (int)(__r);
}
int QMessageBox_question6(void* parent, const char* title, long long title_len, const char* text, long long text_len, int buttons, int defaultButton) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = QMessageBox::question((QWidget*)parent, title_qstr, text_qstr, QFlags<QMessageBox::StandardButton>(buttons), (QMessageBox::StandardButton)(defaultButton));
	return (int)(__r);
}
int QMessageBox_warning5(void* parent, const char* title, long long title_len, const char* text, long long text_len, int buttons) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = QMessageBox::warning((QWidget*)parent, title_qstr, text_qstr, QFlags<QMessageBox::StandardButton>(buttons));
	return (int)(__r);
}
int QMessageBox_warning6(void* parent, const char* title, long long title_len, const char* text, long long text_len, int buttons, int defaultButton) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = QMessageBox::warning((QWidget*)parent, title_qstr, text_qstr, QFlags<QMessageBox::StandardButton>(buttons), (QMessageBox::StandardButton)(defaultButton));
	return (int)(__r);
}
int QMessageBox_critical5(void* parent, const char* title, long long title_len, const char* text, long long text_len, int buttons) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = QMessageBox::critical((QWidget*)parent, title_qstr, text_qstr, QFlags<QMessageBox::StandardButton>(buttons));
	return (int)(__r);
}
int QMessageBox_critical6(void* parent, const char* title, long long title_len, const char* text, long long text_len, int buttons, int defaultButton) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = QMessageBox::critical((QWidget*)parent, title_qstr, text_qstr, QFlags<QMessageBox::StandardButton>(buttons), (QMessageBox::StandardButton)(defaultButton));
	return (int)(__r);
}
void QMessageBox_aboutQt2(void* parent, const char* title, long long title_len) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	QMessageBox::aboutQt((QWidget*)parent, title_qstr);
}
int QMessageBox_information8(void* parent, const char* title, long long title_len, const char* text, long long text_len, int button0, int button1) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = QMessageBox::information((QWidget*)parent, title_qstr, text_qstr, button0, button1);
	return __r;
}
int QMessageBox_information9(void* parent, const char* title, long long title_len, const char* text, long long text_len, int button0, int button1, int button2) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = QMessageBox::information((QWidget*)parent, title_qstr, text_qstr, button0, button1, button2);
	return __r;
}
int QMessageBox_information10(void* parent, const char* title, long long title_len, const char* text, long long text_len, const char* button0Text, long long button0Text_len, const char* button1Text, long long button1Text_len) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	QString text_qstr = QString::fromUtf8(text, text_len);
	QString button0Text_qstr = QString::fromUtf8(button0Text, button0Text_len);
	QString button1Text_qstr = QString::fromUtf8(button1Text, button1Text_len);
	auto&& __r = QMessageBox::information((QWidget*)parent, title_qstr, text_qstr, button0Text_qstr, button1Text_qstr);
	return __r;
}
int QMessageBox_information11(void* parent, const char* title, long long title_len, const char* text, long long text_len, const char* button0Text, long long button0Text_len, const char* button1Text, long long button1Text_len, const char* button2Text, long long button2Text_len) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	QString text_qstr = QString::fromUtf8(text, text_len);
	QString button0Text_qstr = QString::fromUtf8(button0Text, button0Text_len);
	QString button1Text_qstr = QString::fromUtf8(button1Text, button1Text_len);
	QString button2Text_qstr = QString::fromUtf8(button2Text, button2Text_len);
	auto&& __r = QMessageBox::information((QWidget*)parent, title_qstr, text_qstr, button0Text_qstr, button1Text_qstr, button2Text_qstr);
	return __r;
}
int QMessageBox_information12(void* parent, const char* title, long long title_len, const char* text, long long text_len, const char* button0Text, long long button0Text_len, const char* button1Text, long long button1Text_len, const char* button2Text, long long button2Text_len, int defaultButtonNumber) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	QString text_qstr = QString::fromUtf8(text, text_len);
	QString button0Text_qstr = QString::fromUtf8(button0Text, button0Text_len);
	QString button1Text_qstr = QString::fromUtf8(button1Text, button1Text_len);
	QString button2Text_qstr = QString::fromUtf8(button2Text, button2Text_len);
	auto&& __r = QMessageBox::information((QWidget*)parent, title_qstr, text_qstr, button0Text_qstr, button1Text_qstr, button2Text_qstr, defaultButtonNumber);
	return __r;
}
int QMessageBox_information13(void* parent, const char* title, long long title_len, const char* text, long long text_len, const char* button0Text, long long button0Text_len, const char* button1Text, long long button1Text_len, const char* button2Text, long long button2Text_len, int defaultButtonNumber, int escapeButtonNumber) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	QString text_qstr = QString::fromUtf8(text, text_len);
	QString button0Text_qstr = QString::fromUtf8(button0Text, button0Text_len);
	QString button1Text_qstr = QString::fromUtf8(button1Text, button1Text_len);
	QString button2Text_qstr = QString::fromUtf8(button2Text, button2Text_len);
	auto&& __r = QMessageBox::information((QWidget*)parent, title_qstr, text_qstr, button0Text_qstr, button1Text_qstr, button2Text_qstr, defaultButtonNumber, escapeButtonNumber);
	return __r;
}
int QMessageBox_question7(void* parent, const char* title, long long title_len, const char* text, long long text_len, int button0, int button1) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = QMessageBox::question((QWidget*)parent, title_qstr, text_qstr, button0, button1);
	return __r;
}
int QMessageBox_question8(void* parent, const char* title, long long title_len, const char* text, long long text_len, int button0, int button1, int button2) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = QMessageBox::question((QWidget*)parent, title_qstr, text_qstr, button0, button1, button2);
	return __r;
}
int QMessageBox_question9(void* parent, const char* title, long long title_len, const char* text, long long text_len, const char* button0Text, long long button0Text_len, const char* button1Text, long long button1Text_len) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	QString text_qstr = QString::fromUtf8(text, text_len);
	QString button0Text_qstr = QString::fromUtf8(button0Text, button0Text_len);
	QString button1Text_qstr = QString::fromUtf8(button1Text, button1Text_len);
	auto&& __r = QMessageBox::question((QWidget*)parent, title_qstr, text_qstr, button0Text_qstr, button1Text_qstr);
	return __r;
}
int QMessageBox_question10(void* parent, const char* title, long long title_len, const char* text, long long text_len, const char* button0Text, long long button0Text_len, const char* button1Text, long long button1Text_len, const char* button2Text, long long button2Text_len) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	QString text_qstr = QString::fromUtf8(text, text_len);
	QString button0Text_qstr = QString::fromUtf8(button0Text, button0Text_len);
	QString button1Text_qstr = QString::fromUtf8(button1Text, button1Text_len);
	QString button2Text_qstr = QString::fromUtf8(button2Text, button2Text_len);
	auto&& __r = QMessageBox::question((QWidget*)parent, title_qstr, text_qstr, button0Text_qstr, button1Text_qstr, button2Text_qstr);
	return __r;
}
int QMessageBox_question11(void* parent, const char* title, long long title_len, const char* text, long long text_len, const char* button0Text, long long button0Text_len, const char* button1Text, long long button1Text_len, const char* button2Text, long long button2Text_len, int defaultButtonNumber) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	QString text_qstr = QString::fromUtf8(text, text_len);
	QString button0Text_qstr = QString::fromUtf8(button0Text, button0Text_len);
	QString button1Text_qstr = QString::fromUtf8(button1Text, button1Text_len);
	QString button2Text_qstr = QString::fromUtf8(button2Text, button2Text_len);
	auto&& __r = QMessageBox::question((QWidget*)parent, title_qstr, text_qstr, button0Text_qstr, button1Text_qstr, button2Text_qstr, defaultButtonNumber);
	return __r;
}
int QMessageBox_question12(void* parent, const char* title, long long title_len, const char* text, long long text_len, const char* button0Text, long long button0Text_len, const char* button1Text, long long button1Text_len, const char* button2Text, long long button2Text_len, int defaultButtonNumber, int escapeButtonNumber) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	QString text_qstr = QString::fromUtf8(text, text_len);
	QString button0Text_qstr = QString::fromUtf8(button0Text, button0Text_len);
	QString button1Text_qstr = QString::fromUtf8(button1Text, button1Text_len);
	QString button2Text_qstr = QString::fromUtf8(button2Text, button2Text_len);
	auto&& __r = QMessageBox::question((QWidget*)parent, title_qstr, text_qstr, button0Text_qstr, button1Text_qstr, button2Text_qstr, defaultButtonNumber, escapeButtonNumber);
	return __r;
}
int QMessageBox_warning7(void* parent, const char* title, long long title_len, const char* text, long long text_len, int button0, int button1, int button2) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = QMessageBox::warning((QWidget*)parent, title_qstr, text_qstr, button0, button1, button2);
	return __r;
}
int QMessageBox_warning8(void* parent, const char* title, long long title_len, const char* text, long long text_len, const char* button0Text, long long button0Text_len, const char* button1Text, long long button1Text_len) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	QString text_qstr = QString::fromUtf8(text, text_len);
	QString button0Text_qstr = QString::fromUtf8(button0Text, button0Text_len);
	QString button1Text_qstr = QString::fromUtf8(button1Text, button1Text_len);
	auto&& __r = QMessageBox::warning((QWidget*)parent, title_qstr, text_qstr, button0Text_qstr, button1Text_qstr);
	return __r;
}
int QMessageBox_warning9(void* parent, const char* title, long long title_len, const char* text, long long text_len, const char* button0Text, long long button0Text_len, const char* button1Text, long long button1Text_len, const char* button2Text, long long button2Text_len) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	QString text_qstr = QString::fromUtf8(text, text_len);
	QString button0Text_qstr = QString::fromUtf8(button0Text, button0Text_len);
	QString button1Text_qstr = QString::fromUtf8(button1Text, button1Text_len);
	QString button2Text_qstr = QString::fromUtf8(button2Text, button2Text_len);
	auto&& __r = QMessageBox::warning((QWidget*)parent, title_qstr, text_qstr, button0Text_qstr, button1Text_qstr, button2Text_qstr);
	return __r;
}
int QMessageBox_warning10(void* parent, const char* title, long long title_len, const char* text, long long text_len, const char* button0Text, long long button0Text_len, const char* button1Text, long long button1Text_len, const char* button2Text, long long button2Text_len, int defaultButtonNumber) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	QString text_qstr = QString::fromUtf8(text, text_len);
	QString button0Text_qstr = QString::fromUtf8(button0Text, button0Text_len);
	QString button1Text_qstr = QString::fromUtf8(button1Text, button1Text_len);
	QString button2Text_qstr = QString::fromUtf8(button2Text, button2Text_len);
	auto&& __r = QMessageBox::warning((QWidget*)parent, title_qstr, text_qstr, button0Text_qstr, button1Text_qstr, button2Text_qstr, defaultButtonNumber);
	return __r;
}
int QMessageBox_warning11(void* parent, const char* title, long long title_len, const char* text, long long text_len, const char* button0Text, long long button0Text_len, const char* button1Text, long long button1Text_len, const char* button2Text, long long button2Text_len, int defaultButtonNumber, int escapeButtonNumber) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	QString text_qstr = QString::fromUtf8(text, text_len);
	QString button0Text_qstr = QString::fromUtf8(button0Text, button0Text_len);
	QString button1Text_qstr = QString::fromUtf8(button1Text, button1Text_len);
	QString button2Text_qstr = QString::fromUtf8(button2Text, button2Text_len);
	auto&& __r = QMessageBox::warning((QWidget*)parent, title_qstr, text_qstr, button0Text_qstr, button1Text_qstr, button2Text_qstr, defaultButtonNumber, escapeButtonNumber);
	return __r;
}
int QMessageBox_critical7(void* parent, const char* title, long long title_len, const char* text, long long text_len, int button0, int button1, int button2) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = QMessageBox::critical((QWidget*)parent, title_qstr, text_qstr, button0, button1, button2);
	return __r;
}
int QMessageBox_critical8(void* parent, const char* title, long long title_len, const char* text, long long text_len, const char* button0Text, long long button0Text_len, const char* button1Text, long long button1Text_len) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	QString text_qstr = QString::fromUtf8(text, text_len);
	QString button0Text_qstr = QString::fromUtf8(button0Text, button0Text_len);
	QString button1Text_qstr = QString::fromUtf8(button1Text, button1Text_len);
	auto&& __r = QMessageBox::critical((QWidget*)parent, title_qstr, text_qstr, button0Text_qstr, button1Text_qstr);
	return __r;
}
int QMessageBox_critical9(void* parent, const char* title, long long title_len, const char* text, long long text_len, const char* button0Text, long long button0Text_len, const char* button1Text, long long button1Text_len, const char* button2Text, long long button2Text_len) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	QString text_qstr = QString::fromUtf8(text, text_len);
	QString button0Text_qstr = QString::fromUtf8(button0Text, button0Text_len);
	QString button1Text_qstr = QString::fromUtf8(button1Text, button1Text_len);
	QString button2Text_qstr = QString::fromUtf8(button2Text, button2Text_len);
	auto&& __r = QMessageBox::critical((QWidget*)parent, title_qstr, text_qstr, button0Text_qstr, button1Text_qstr, button2Text_qstr);
	return __r;
}
int QMessageBox_critical10(void* parent, const char* title, long long title_len, const char* text, long long text_len, const char* button0Text, long long button0Text_len, const char* button1Text, long long button1Text_len, const char* button2Text, long long button2Text_len, int defaultButtonNumber) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	QString text_qstr = QString::fromUtf8(text, text_len);
	QString button0Text_qstr = QString::fromUtf8(button0Text, button0Text_len);
	QString button1Text_qstr = QString::fromUtf8(button1Text, button1Text_len);
	QString button2Text_qstr = QString::fromUtf8(button2Text, button2Text_len);
	auto&& __r = QMessageBox::critical((QWidget*)parent, title_qstr, text_qstr, button0Text_qstr, button1Text_qstr, button2Text_qstr, defaultButtonNumber);
	return __r;
}
int QMessageBox_critical11(void* parent, const char* title, long long title_len, const char* text, long long text_len, const char* button0Text, long long button0Text_len, const char* button1Text, long long button1Text_len, const char* button2Text, long long button2Text_len, int defaultButtonNumber, int escapeButtonNumber) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	QString text_qstr = QString::fromUtf8(text, text_len);
	QString button0Text_qstr = QString::fromUtf8(button0Text, button0Text_len);
	QString button1Text_qstr = QString::fromUtf8(button1Text, button1Text_len);
	QString button2Text_qstr = QString::fromUtf8(button2Text, button2Text_len);
	auto&& __r = QMessageBox::critical((QWidget*)parent, title_qstr, text_qstr, button0Text_qstr, button1Text_qstr, button2Text_qstr, defaultButtonNumber, escapeButtonNumber);
	return __r;
}
void* QMessageBox_createWindowContainer2(void* window, void* parent) {
	auto&& __r = QMessageBox::createWindowContainer((QWindow*)window, (QWidget*)parent);
	return (void*)(__r);
}
void* QMessageBox_createWindowContainer3(void* window, void* parent, int flags) {
	auto&& __r = QMessageBox::createWindowContainer((QWindow*)window, (QWidget*)parent, QFlags<Qt::WindowType>(flags));
	return (void*)(__r);
}
void QMessageBox_updateMicroFocus2(void* self, int query) {
	((QMessageBox__Go*)self)->updateMicroFocus((Qt::InputMethodQuery)(query));
}
void QMessageBox_create2(void* self, unsigned long long arg1) {
	((QMessageBox__Go*)self)->create(arg1);
}
void QMessageBox_create3(void* self, unsigned long long arg1, int initializeWindow) {
	((QMessageBox__Go*)self)->create(arg1, initializeWindow);
}
void QMessageBox_create4(void* self, unsigned long long arg1, int initializeWindow, int destroyOldWindow) {
	((QMessageBox__Go*)self)->create(arg1, initializeWindow, destroyOldWindow);
}
void QMessageBox_destroy2(void* self, int destroyWindow) {
	((QMessageBox__Go*)self)->destroy(destroyWindow);
}
void QMessageBox_destroy3(void* self, int destroyWindow, int destroySubWindows) {
	((QMessageBox__Go*)self)->destroy(destroyWindow, destroySubWindows);
}
void* QMessageBox_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QMessageBox::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QMessageBox_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QMessageBox::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QMessageBox_metaObject_default(void* self) {
	auto&& __r = ((QMessageBox*)self)->QMessageBox::metaObject();
	return (void*)(__r);
}
int QMessageBox_event_default(void* self, void* e) {
	auto&& __r = ((QMessageBox__Go*)self)->cppgoBase_event((QEvent*)e);
	return __r;
}
void QMessageBox_resizeEvent_default(void* self, void* event) {
	((QMessageBox__Go*)self)->cppgoBase_resizeEvent((QResizeEvent*)event);
}
void QMessageBox_showEvent_default(void* self, void* event) {
	((QMessageBox__Go*)self)->cppgoBase_showEvent((QShowEvent*)event);
}
void QMessageBox_closeEvent_default(void* self, void* event) {
	((QMessageBox__Go*)self)->cppgoBase_closeEvent((QCloseEvent*)event);
}
void QMessageBox_keyPressEvent_default(void* self, void* event) {
	((QMessageBox__Go*)self)->cppgoBase_keyPressEvent((QKeyEvent*)event);
}
void QMessageBox_changeEvent_default(void* self, void* event) {
	((QMessageBox__Go*)self)->cppgoBase_changeEvent((QEvent*)event);
}
void QMessageBox_setVisible_default(void* self, int visible) {
	((QMessageBox*)self)->QMessageBox::setVisible(visible);
}
void* QMessageBox_sizeHint_default(void* self) {
	auto&& __r = ((QMessageBox*)self)->QMessageBox::sizeHint();
	return new QSize(__r);
}
void* QMessageBox_minimumSizeHint_default(void* self) {
	auto&& __r = ((QMessageBox*)self)->QMessageBox::minimumSizeHint();
	return new QSize(__r);
}
int QMessageBox_exec_default(void* self) {
	auto&& __r = ((QMessageBox*)self)->QMessageBox::exec();
	return __r;
}
void QMessageBox_done_default(void* self, int arg1) {
	((QMessageBox*)self)->QMessageBox::done(arg1);
}
void QMessageBox_accept_default(void* self) {
	((QMessageBox*)self)->QMessageBox::accept();
}
void QMessageBox_reject_default(void* self) {
	((QMessageBox*)self)->QMessageBox::reject();
}
void QMessageBox_contextMenuEvent_default(void* self, void* arg1) {
	((QMessageBox__Go*)self)->cppgoBase_contextMenuEvent((QContextMenuEvent*)arg1);
}
int QMessageBox_eventFilter_default(void* self, void* arg1, void* arg2) {
	auto&& __r = ((QMessageBox__Go*)self)->cppgoBase_eventFilter((QObject*)arg1, (QEvent*)arg2);
	return __r;
}
int QMessageBox_devType_default(void* self) {
	auto&& __r = ((QMessageBox*)self)->QMessageBox::devType();
	return __r;
}
int QMessageBox_heightForWidth_default(void* self, int arg1) {
	auto&& __r = ((QMessageBox*)self)->QMessageBox::heightForWidth(arg1);
	return __r;
}
int QMessageBox_hasHeightForWidth_default(void* self) {
	auto&& __r = ((QMessageBox*)self)->QMessageBox::hasHeightForWidth();
	return __r;
}
void* QMessageBox_paintEngine_default(void* self) {
	auto&& __r = ((QMessageBox*)self)->QMessageBox::paintEngine();
	return (void*)(__r);
}
void QMessageBox_mousePressEvent_default(void* self, void* event) {
	((QMessageBox__Go*)self)->cppgoBase_mousePressEvent((QMouseEvent*)event);
}
void QMessageBox_mouseReleaseEvent_default(void* self, void* event) {
	((QMessageBox__Go*)self)->cppgoBase_mouseReleaseEvent((QMouseEvent*)event);
}
void QMessageBox_mouseDoubleClickEvent_default(void* self, void* event) {
	((QMessageBox__Go*)self)->cppgoBase_mouseDoubleClickEvent((QMouseEvent*)event);
}
void QMessageBox_mouseMoveEvent_default(void* self, void* event) {
	((QMessageBox__Go*)self)->cppgoBase_mouseMoveEvent((QMouseEvent*)event);
}
void QMessageBox_wheelEvent_default(void* self, void* event) {
	((QMessageBox__Go*)self)->cppgoBase_wheelEvent((QWheelEvent*)event);
}
void QMessageBox_keyReleaseEvent_default(void* self, void* event) {
	((QMessageBox__Go*)self)->cppgoBase_keyReleaseEvent((QKeyEvent*)event);
}
void QMessageBox_focusInEvent_default(void* self, void* event) {
	((QMessageBox__Go*)self)->cppgoBase_focusInEvent((QFocusEvent*)event);
}
void QMessageBox_focusOutEvent_default(void* self, void* event) {
	((QMessageBox__Go*)self)->cppgoBase_focusOutEvent((QFocusEvent*)event);
}
void QMessageBox_enterEvent_default(void* self, void* event) {
	((QMessageBox__Go*)self)->cppgoBase_enterEvent((QEnterEvent*)event);
}
void QMessageBox_leaveEvent_default(void* self, void* event) {
	((QMessageBox__Go*)self)->cppgoBase_leaveEvent((QEvent*)event);
}
void QMessageBox_paintEvent_default(void* self, void* event) {
	((QMessageBox__Go*)self)->cppgoBase_paintEvent((QPaintEvent*)event);
}
void QMessageBox_moveEvent_default(void* self, void* event) {
	((QMessageBox__Go*)self)->cppgoBase_moveEvent((QMoveEvent*)event);
}
void QMessageBox_tabletEvent_default(void* self, void* event) {
	((QMessageBox__Go*)self)->cppgoBase_tabletEvent((QTabletEvent*)event);
}
void QMessageBox_actionEvent_default(void* self, void* event) {
	((QMessageBox__Go*)self)->cppgoBase_actionEvent((QActionEvent*)event);
}
void QMessageBox_dragEnterEvent_default(void* self, void* event) {
	((QMessageBox__Go*)self)->cppgoBase_dragEnterEvent((QDragEnterEvent*)event);
}
void QMessageBox_dragMoveEvent_default(void* self, void* event) {
	((QMessageBox__Go*)self)->cppgoBase_dragMoveEvent((QDragMoveEvent*)event);
}
void QMessageBox_dragLeaveEvent_default(void* self, void* event) {
	((QMessageBox__Go*)self)->cppgoBase_dragLeaveEvent((QDragLeaveEvent*)event);
}
void QMessageBox_dropEvent_default(void* self, void* event) {
	((QMessageBox__Go*)self)->cppgoBase_dropEvent((QDropEvent*)event);
}
void QMessageBox_hideEvent_default(void* self, void* event) {
	((QMessageBox__Go*)self)->cppgoBase_hideEvent((QHideEvent*)event);
}
int QMessageBox_metric_default(void* self, int arg1) {
	auto&& __r = ((QMessageBox__Go*)self)->cppgoBase_metric(arg1);
	return __r;
}
void QMessageBox_initPainter_default(void* self, void* painter) {
	((QMessageBox__Go*)self)->cppgoBase_initPainter((QPainter*)painter);
}
void* QMessageBox_redirected_default(void* self, void* offset) {
	auto&& __r = ((QMessageBox__Go*)self)->cppgoBase_redirected((QPoint*)offset);
	return (void*)(__r);
}
void* QMessageBox_sharedPainter_default(void* self) {
	auto&& __r = ((QMessageBox__Go*)self)->cppgoBase_sharedPainter();
	return (void*)(__r);
}
void QMessageBox_inputMethodEvent_default(void* self, void* arg1) {
	((QMessageBox__Go*)self)->cppgoBase_inputMethodEvent((QInputMethodEvent*)arg1);
}
void* QMessageBox_inputMethodQuery_default(void* self, int arg1) {
	auto&& __r = ((QMessageBox*)self)->QMessageBox::inputMethodQuery((Qt::InputMethodQuery)(arg1));
	return new QVariant(__r);
}
int QMessageBox_focusNextPrevChild_default(void* self, int next) {
	auto&& __r = ((QMessageBox__Go*)self)->cppgoBase_focusNextPrevChild(next);
	return __r;
}
void QMessageBox_timerEvent_default(void* self, void* event) {
	((QMessageBox__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QMessageBox_childEvent_default(void* self, void* event) {
	((QMessageBox__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QMessageBox_customEvent_default(void* self, void* event) {
	((QMessageBox__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QMessageBox_connectNotify_default(void* self, void* signal) {
	((QMessageBox__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QMessageBox_disconnectNotify_default(void* self, void* signal) {
	((QMessageBox__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QMessageBox_connect_buttonClicked(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QMessageBox*)self, &QMessageBox::buttonClicked, (QMessageBox*)self, [__id](QAbstractButton* button){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)button);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QMessageBox);
	__conns_QMessageBox[__id] = __c;
}
void QMessageBox_connect_finished(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QMessageBox*)self, &QMessageBox::finished, (QMessageBox*)self, [__id](int result){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, result);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QMessageBox);
	__conns_QMessageBox[__id] = __c;
}
void QMessageBox_connect_accepted(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QMessageBox*)self, &QMessageBox::accepted, (QMessageBox*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QMessageBox);
	__conns_QMessageBox[__id] = __c;
}
void QMessageBox_connect_rejected(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QMessageBox*)self, &QMessageBox::rejected, (QMessageBox*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QMessageBox);
	__conns_QMessageBox[__id] = __c;
}
void QMessageBox_connect_windowTitleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QMessageBox*)self, &QMessageBox::windowTitleChanged, (QMessageBox*)self, [__id](const QString& title){
		QByteArray title_ba = title.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)title_ba.constData(), title_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QMessageBox);
	__conns_QMessageBox[__id] = __c;
}
void QMessageBox_connect_windowIconChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QMessageBox*)self, &QMessageBox::windowIconChanged, (QMessageBox*)self, [__id](const QIcon& icon){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&icon);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QMessageBox);
	__conns_QMessageBox[__id] = __c;
}
void QMessageBox_connect_windowIconTextChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QMessageBox*)self, &QMessageBox::windowIconTextChanged, (QMessageBox*)self, [__id](const QString& iconText){
		QByteArray iconText_ba = iconText.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)iconText_ba.constData(), iconText_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QMessageBox);
	__conns_QMessageBox[__id] = __c;
}
void QMessageBox_connect_customContextMenuRequested(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QMessageBox*)self, &QMessageBox::customContextMenuRequested, (QMessageBox*)self, [__id](const QPoint& pos){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&pos);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QMessageBox);
	__conns_QMessageBox[__id] = __c;
}
void QMessageBox_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QMessageBox*)self, &QMessageBox::destroyed, (QMessageBox*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QMessageBox);
	__conns_QMessageBox[__id] = __c;
}
void QMessageBox_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QMessageBox*)self, &QMessageBox::destroyed, (QMessageBox*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QMessageBox);
	__conns_QMessageBox[__id] = __c;
}
void QMessageBox_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QMessageBox*)self, &QMessageBox::objectNameChanged, (QMessageBox*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QMessageBox);
	__conns_QMessageBox[__id] = __c;
}
void QMessageBox__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QMessageBox);
		auto __it = __conns_QMessageBox.find(id);
		if (__it == __conns_QMessageBox.end()) return;
		__c = __it->second;
		__conns_QMessageBox.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QMessageBox__Go_new(long long id) { return new QMessageBox__Go(id); }
void QMessageBox_delete(void* self) { delete (QMessageBox*)self; }

}

#include <QtWidgets/qstylepainter.h>
#include <QtGui/qpainter.h>
#include <QtWidgets/qwidget.h>
#include <QtGui/qpaintdevice.h>
#include <QtWidgets/qstyle.h>
#include <QtWidgets/qstyleoption.h>
#include <QtCore/qrect.h>
#include <QtGui/qpalette.h>
#include <QtCore/qstring.h>
#include <QtGui/qpixmap.h>
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
#include <QtGui/qtransform.h>
#include <QtGui/qpolygon.h>
#include <QtCore/qline.h>
#include <QtGui/qpicture.h>
#include <QtGui/qimage.h>
#include <QtGui/qglyphrun.h>
#include <QtGui/qstatictext.h>
#include <QtGui/qtextoption.h>
#include <QtGui/qpaintengine.h>
#include "gen_qstylepainter.h"

extern "C" {

void* QStylePainter_base(void* self) { return (void*)static_cast<QPainter*>((QStylePainter*)self); }
void* QStylePainter_new(void* w) {
	return new QStylePainter((QWidget*)w);
}
void* QStylePainter_new2() {
	return new QStylePainter();
}
void* QStylePainter_new3(void* pd, void* w) {
	return new QStylePainter((QPaintDevice*)pd, (QWidget*)w);
}
int QStylePainter_begin(void* self, void* w) {
	auto&& __r = ((QStylePainter*)self)->begin((QWidget*)w);
	return __r;
}
int QStylePainter_begin2(void* self, void* pd, void* w) {
	auto&& __r = ((QStylePainter*)self)->begin((QPaintDevice*)pd, (QWidget*)w);
	return __r;
}
void QStylePainter_drawPrimitive(void* self, int pe, void* opt) {
	((QStylePainter*)self)->drawPrimitive((QStyle::PrimitiveElement)(pe), *(QStyleOption*)opt);
}
void QStylePainter_drawControl(void* self, int ce, void* opt) {
	((QStylePainter*)self)->drawControl((QStyle::ControlElement)(ce), *(QStyleOption*)opt);
}
void QStylePainter_drawComplexControl(void* self, int cc, void* opt) {
	((QStylePainter*)self)->drawComplexControl((QStyle::ComplexControl)(cc), *(QStyleOptionComplex*)opt);
}
void QStylePainter_drawItemText(void* self, void* r, int flags, void* pal, int enabled, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	((QStylePainter*)self)->drawItemText(*(QRect*)r, flags, *(QPalette*)pal, enabled, text_qstr);
}
void QStylePainter_drawItemPixmap(void* self, void* r, int flags, void* pixmap) {
	((QStylePainter*)self)->drawItemPixmap(*(QRect*)r, flags, *(QPixmap*)pixmap);
}
void* QStylePainter_style(void* self) {
	auto&& __r = ((QStylePainter*)self)->style();
	return (void*)(__r);
}
void QStylePainter_drawItemText2(void* self, void* r, int flags, void* pal, int enabled, const char* text, long long text_len, int textRole) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	((QStylePainter*)self)->drawItemText(*(QRect*)r, flags, *(QPalette*)pal, enabled, text_qstr, (QPalette::ColorRole)(textRole));
}
void QStylePainter_delete(void* self) { delete (QStylePainter*)self; }

}

