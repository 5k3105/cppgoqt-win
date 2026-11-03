#include <QtCore/qnamespace.h>
#include "gen_qkeycombination.h"

extern "C" {

void* QKeyCombination_new() {
	return new QKeyCombination();
}
void* QKeyCombination_new2(int modifiers) {
	return new QKeyCombination(QFlags<Qt::Modifier>(modifiers));
}
void* QKeyCombination_new3(int modifiers) {
	return new QKeyCombination(QFlags<Qt::KeyboardModifier>(modifiers));
}
void* QKeyCombination_new4(void* arg1) {
	return new QKeyCombination(*(QKeyCombination*)arg1);
}
void* QKeyCombination_new5(int key) {
	return new QKeyCombination((Qt::Key)(key));
}
void* QKeyCombination_new6(int modifiers, int key) {
	return new QKeyCombination(QFlags<Qt::Modifier>(modifiers), (Qt::Key)(key));
}
void* QKeyCombination_new7(int modifiers, int key) {
	return new QKeyCombination(QFlags<Qt::KeyboardModifier>(modifiers), (Qt::Key)(key));
}
int QKeyCombination_keyboardModifiers(void* self) {
	auto&& __r = ((QKeyCombination*)self)->keyboardModifiers();
	return (int)(__r);
}
int QKeyCombination_key(void* self) {
	auto&& __r = ((QKeyCombination*)self)->key();
	return (int)(__r);
}
void* QKeyCombination_fromCombined(int combined) {
	auto&& __r = QKeyCombination::fromCombined(combined);
	return new QKeyCombination(__r);
}
int QKeyCombination_toCombined(void* self) {
	auto&& __r = ((QKeyCombination*)self)->toCombined();
	return __r;
}
int QKeyCombination_ToInt(void* self) {
	auto&& __r = static_cast<int>(*(QKeyCombination*)self);
	return __r;
}
int QKeyCombination_OperatorEqual(void* self, void* rhs) {
	auto&& __r = (*(QKeyCombination*)self) == (*(QKeyCombination*)rhs);
	return __r;
}
int QKeyCombination_OperatorNotEqual(void* self, void* rhs) {
	auto&& __r = (*(QKeyCombination*)self) != (*(QKeyCombination*)rhs);
	return __r;
}
void QKeyCombination_delete(void* self) { delete (QKeyCombination*)self; }

}

#include <QtCore/qstringconverter.h>
#include <QtCore/qstringconverter_base.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qbytearrayview.h>
#include <QtCore/qchar.h>
#include <QtCore/qstring.h>
#include "gen_qstringdecoder.h"

extern "C" {

void* QStringDecoder_base(void* self) { return (void*)static_cast<QStringConverter*>((QStringDecoder*)self); }
void QStringDecoder_chain(void* self, void** out) {
	QStringDecoder* __p = (QStringDecoder*)self;
	QStringConverter* __b0 = static_cast<QStringConverter*>(__p); out[0] = (void*)__b0;
	QStringConverterBase* __b1 = static_cast<QStringConverterBase*>(__b0); out[1] = (void*)__b1;
}
void* QStringDecoder_new(int encoding) {
	return new QStringDecoder((QStringConverter::Encoding)(encoding));
}
void* QStringDecoder_new2() {
	return new QStringDecoder();
}
void* QStringDecoder_new3(const char* name, long long name_len) {
	QAnyStringView name_qstr = QAnyStringView(name, name_len);
	return new QStringDecoder(name_qstr);
}
void* QStringDecoder_new4(int encoding, int flags) {
	return new QStringDecoder((QStringConverter::Encoding)(encoding), QFlags<QStringConverterBase::Flag>(flags));
}
void* QStringDecoder_new5(const char* name, long long name_len, int f) {
	QAnyStringView name_qstr = QAnyStringView(name, name_len);
	return new QStringDecoder(name_qstr, QFlags<QStringConverterBase::Flag>(f));
}
long long QStringDecoder_requiredSpace(void* self, long long inputLength) {
	auto&& __r = ((QStringDecoder*)self)->requiredSpace(inputLength);
	return __r;
}
void* QStringDecoder_appendToBuffer(void* self, void* out, void* ba) {
	auto&& __r = ((QStringDecoder*)self)->appendToBuffer((QChar*)out, *(QByteArrayView*)ba);
	return (void*)(__r);
}
int QStringDecoder_encodingForName(const char* name, long long name_len, int* __ok) {
	QAnyStringView name_qstr = QAnyStringView(name, name_len);
	auto&& __r = QStringDecoder::encodingForName(name_qstr);
	*__ok = __r.has_value() ? 1 : 0;
	return __r.has_value() ? (int)(*__r) : (int)0;
}
const char* QStringDecoder_nameForEncoding(int e) {
	auto&& __r = QStringDecoder::nameForEncoding((QStringConverter::Encoding)(e));
	return __r;
}
int QStringDecoder_encodingForData(void* data, int* __ok) {
	auto&& __r = QStringDecoder::encodingForData(*(QByteArrayView*)data);
	*__ok = __r.has_value() ? 1 : 0;
	return __r.has_value() ? (int)(*__r) : (int)0;
}
int QStringDecoder_encodingForHtml(void* data, int* __ok) {
	auto&& __r = QStringDecoder::encodingForHtml(*(QByteArrayView*)data);
	*__ok = __r.has_value() ? 1 : 0;
	return __r.has_value() ? (int)(*__r) : (int)0;
}
char** QStringDecoder_availableCodecs(long long* __out_len) {
	auto&& __r = QStringDecoder::availableCodecs();
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
int QStringDecoder_encodingForData2(void* data, unsigned short expectedFirstCharacter, int* __ok) {
	auto&& __r = QStringDecoder::encodingForData(*(QByteArrayView*)data, expectedFirstCharacter);
	*__ok = __r.has_value() ? 1 : 0;
	return __r.has_value() ? (int)(*__r) : (int)0;
}
void QStringDecoder_delete(void* self) { delete (QStringDecoder*)self; }

}

#include <QtCore/qhash.h>
#include "gen_qhashdummyvalue.h"

extern "C" {

void* QHashDummyValue_new() {
	return new QHashDummyValue();
}
int QHashDummyValue_OperatorEqual(void* self, void* arg1) {
	auto&& __r = (*(QHashDummyValue*)self) == (*(QHashDummyValue*)arg1);
	return __r;
}
int QHashDummyValue_OperatorNotEqual(void* self, void* arg1) {
	auto&& __r = (*(QHashDummyValue*)self) != (*(QHashDummyValue*)arg1);
	return __r;
}
void QHashDummyValue_delete(void* self) { delete (QHashDummyValue*)self; }

}

#include <QtCore/qmutex.h>
#include <QtCore/qdeadlinetimer.h>
#include "gen_qrecursivemutex.h"

extern "C" {

void* QRecursiveMutex_new() {
	return new QRecursiveMutex();
}
void QRecursiveMutex_lock(void* self) {
	((QRecursiveMutex*)self)->lock();
}
int QRecursiveMutex_tryLock(void* self, int timeout) {
	auto&& __r = ((QRecursiveMutex*)self)->tryLock(timeout);
	return __r;
}
int QRecursiveMutex_tryLock2(void* self) {
	auto&& __r = ((QRecursiveMutex*)self)->tryLock();
	return __r;
}
void QRecursiveMutex_unlock(void* self) {
	((QRecursiveMutex*)self)->unlock();
}
int QRecursiveMutex_try_lock(void* self) {
	auto&& __r = ((QRecursiveMutex*)self)->try_lock();
	return __r;
}
int QRecursiveMutex_tryLock3(void* self, void* timer) {
	auto&& __r = ((QRecursiveMutex*)self)->tryLock(*(QDeadlineTimer*)timer);
	return __r;
}
void QRecursiveMutex_delete(void* self) { delete (QRecursiveMutex*)self; }

}

#include <QtCore/qregularexpression.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qstring.h>
#include <QtCore/qstringview.h>
#include "gen_qregularexpressionmatch.h"

extern "C" {

void* QRegularExpressionMatch_new() {
	return new QRegularExpressionMatch();
}
void* QRegularExpressionMatch_new2(void* match) {
	return new QRegularExpressionMatch(*(QRegularExpressionMatch*)match);
}
void* QRegularExpressionMatch_OperatorAssign(void* self, void* match) {
	auto&& __r = ((QRegularExpressionMatch*)self)->operator=(*(QRegularExpressionMatch*)match);
	return (void*)&(__r);
}
void QRegularExpressionMatch_swap(void* self, void* other) {
	((QRegularExpressionMatch*)self)->swap(*(QRegularExpressionMatch*)other);
}
void* QRegularExpressionMatch_regularExpression(void* self) {
	auto&& __r = ((QRegularExpressionMatch*)self)->regularExpression();
	return new QRegularExpression(__r);
}
int QRegularExpressionMatch_matchType(void* self) {
	auto&& __r = ((QRegularExpressionMatch*)self)->matchType();
	return (int)(__r);
}
int QRegularExpressionMatch_matchOptions(void* self) {
	auto&& __r = ((QRegularExpressionMatch*)self)->matchOptions();
	return (int)(__r);
}
int QRegularExpressionMatch_hasMatch(void* self) {
	auto&& __r = ((QRegularExpressionMatch*)self)->hasMatch();
	return __r;
}
int QRegularExpressionMatch_hasPartialMatch(void* self) {
	auto&& __r = ((QRegularExpressionMatch*)self)->hasPartialMatch();
	return __r;
}
int QRegularExpressionMatch_isValid(void* self) {
	auto&& __r = ((QRegularExpressionMatch*)self)->isValid();
	return __r;
}
int QRegularExpressionMatch_lastCapturedIndex(void* self) {
	auto&& __r = ((QRegularExpressionMatch*)self)->lastCapturedIndex();
	return __r;
}
int QRegularExpressionMatch_hasCaptured(void* self, const char* name, long long name_len) {
	QAnyStringView name_qstr = QAnyStringView(name, name_len);
	auto&& __r = ((QRegularExpressionMatch*)self)->hasCaptured(name_qstr);
	return __r;
}
int QRegularExpressionMatch_hasCaptured2(void* self, int nth) {
	auto&& __r = ((QRegularExpressionMatch*)self)->hasCaptured(nth);
	return __r;
}
char* QRegularExpressionMatch_captured(void* self, long long* __out_len) {
	auto&& __r = ((QRegularExpressionMatch*)self)->captured();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QRegularExpressionMatch_capturedView(void* self) {
	auto&& __r = ((QRegularExpressionMatch*)self)->capturedView();
	return new QStringView(__r);
}
char* QRegularExpressionMatch_captured2(void* self, const char* name, long long name_len, long long* __out_len) {
	QAnyStringView name_qstr = QAnyStringView(name, name_len);
	auto&& __r = ((QRegularExpressionMatch*)self)->captured(name_qstr);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QRegularExpressionMatch_capturedView2(void* self, const char* name, long long name_len) {
	QAnyStringView name_qstr = QAnyStringView(name, name_len);
	auto&& __r = ((QRegularExpressionMatch*)self)->capturedView(name_qstr);
	return new QStringView(__r);
}
char** QRegularExpressionMatch_capturedTexts(void* self, long long* __out_len) {
	auto&& __r = ((QRegularExpressionMatch*)self)->capturedTexts();
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
long long QRegularExpressionMatch_capturedStart(void* self) {
	auto&& __r = ((QRegularExpressionMatch*)self)->capturedStart();
	return __r;
}
long long QRegularExpressionMatch_capturedLength(void* self) {
	auto&& __r = ((QRegularExpressionMatch*)self)->capturedLength();
	return __r;
}
long long QRegularExpressionMatch_capturedEnd(void* self) {
	auto&& __r = ((QRegularExpressionMatch*)self)->capturedEnd();
	return __r;
}
long long QRegularExpressionMatch_capturedStart2(void* self, const char* name, long long name_len) {
	QAnyStringView name_qstr = QAnyStringView(name, name_len);
	auto&& __r = ((QRegularExpressionMatch*)self)->capturedStart(name_qstr);
	return __r;
}
long long QRegularExpressionMatch_capturedLength2(void* self, const char* name, long long name_len) {
	QAnyStringView name_qstr = QAnyStringView(name, name_len);
	auto&& __r = ((QRegularExpressionMatch*)self)->capturedLength(name_qstr);
	return __r;
}
long long QRegularExpressionMatch_capturedEnd2(void* self, const char* name, long long name_len) {
	QAnyStringView name_qstr = QAnyStringView(name, name_len);
	auto&& __r = ((QRegularExpressionMatch*)self)->capturedEnd(name_qstr);
	return __r;
}
char* QRegularExpressionMatch_captured3(void* self, int nth, long long* __out_len) {
	auto&& __r = ((QRegularExpressionMatch*)self)->captured(nth);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QRegularExpressionMatch_capturedView3(void* self, int nth) {
	auto&& __r = ((QRegularExpressionMatch*)self)->capturedView(nth);
	return new QStringView(__r);
}
long long QRegularExpressionMatch_capturedStart3(void* self, int nth) {
	auto&& __r = ((QRegularExpressionMatch*)self)->capturedStart(nth);
	return __r;
}
long long QRegularExpressionMatch_capturedLength3(void* self, int nth) {
	auto&& __r = ((QRegularExpressionMatch*)self)->capturedLength(nth);
	return __r;
}
long long QRegularExpressionMatch_capturedEnd3(void* self, int nth) {
	auto&& __r = ((QRegularExpressionMatch*)self)->capturedEnd(nth);
	return __r;
}
void QRegularExpressionMatch_delete(void* self) { delete (QRegularExpressionMatch*)self; }

}

#include <QtCore/qtimer.h>
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
#include "gen_qtimer.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QTimer__Go : public QTimer {
public:
	long long __goId;
	unsigned char __goOn[8] = {};
	template<typename... A> QTimer__Go(long long id, A&&... a) : QTimer(std::forward<A>(a)...), __goId(id) {}
	using QTimer::sender;
	using QTimer::senderSignalIndex;
	using QTimer::receivers;
	using QTimer::isSignalConnected;
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QTimer::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QTimer::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QTimer::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QTimer::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QTimer::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QTimer::metaObject();
	}
	void timerEvent(QTimerEvent* arg1) override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)arg1, &__h); if (__h) return; }
		QTimer::timerEvent(arg1);
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QTimer::event(event);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QTimer::eventFilter(watched, event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QTimer::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QTimer::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QTimer::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QTimer::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QTimer;
static std::mutex __connsMu_QTimer;

extern "C" {

void* QTimer_base(void* self) { return (void*)static_cast<QObject*>((QTimer*)self); }
void QTimer__set_override(void* self, int idx) { ((QTimer__Go*)self)->__goOn[idx] = 1; }
void* QTimer_fromQObject(void* o) { return (void*)qobject_cast<QTimer*>((QObject*)o); }
void* QTimer_new(long long __goid) {
	return new QTimer__Go(__goid);
}
void* QTimer_new2(long long __goid, void* parent) {
	return new QTimer__Go(__goid, (QObject*)parent);
}
char* QTimer_tr(const char* s, long long* __out_len) {
	auto&& __r = QTimer::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QTimer_isActive(void* self) {
	auto&& __r = ((QTimer*)self)->isActive();
	return __r;
}
int QTimer_timerId(void* self) {
	auto&& __r = ((QTimer*)self)->timerId();
	return __r;
}
int QTimer_id(void* self) {
	auto&& __r = ((QTimer*)self)->id();
	return (int)(__r);
}
void QTimer_setInterval(void* self, int msec) {
	((QTimer*)self)->setInterval(msec);
}
int QTimer_interval(void* self) {
	auto&& __r = ((QTimer*)self)->interval();
	return __r;
}
int QTimer_remainingTime(void* self) {
	auto&& __r = ((QTimer*)self)->remainingTime();
	return __r;
}
void QTimer_setTimerType(void* self, int atype) {
	((QTimer*)self)->setTimerType((Qt::TimerType)(atype));
}
int QTimer_timerType(void* self) {
	auto&& __r = ((QTimer*)self)->timerType();
	return (int)(__r);
}
void QTimer_setSingleShot(void* self, int singleShot) {
	((QTimer*)self)->setSingleShot(singleShot);
}
int QTimer_isSingleShot(void* self) {
	auto&& __r = ((QTimer*)self)->isSingleShot();
	return __r;
}
void QTimer_singleShot(int msec, void* receiver, const char* member) {
	QTimer::singleShot(msec, (QObject*)receiver, member);
}
void QTimer_singleShot2(int msec, int timerType, void* receiver, const char* member) {
	QTimer::singleShot(msec, (Qt::TimerType)(timerType), (QObject*)receiver, member);
}
void QTimer_start(void* self, int msec) {
	((QTimer*)self)->start(msec);
}
void QTimer_start2(void* self) {
	((QTimer*)self)->start();
}
void QTimer_stop(void* self) {
	((QTimer*)self)->stop();
}
void QTimer_setInterval2(void* self, long long value) {
	((QTimer*)self)->setInterval(std::chrono::milliseconds((long long)value));
}
long long QTimer_intervalAsDuration(void* self) {
	auto&& __r = ((QTimer*)self)->intervalAsDuration();
	return (long long)(__r).count();
}
long long QTimer_remainingTimeAsDuration(void* self) {
	auto&& __r = ((QTimer*)self)->remainingTimeAsDuration();
	return (long long)(__r).count();
}
void QTimer_singleShot3(long long value, void* receiver, const char* member) {
	QTimer::singleShot(std::chrono::nanoseconds((long long)value), (QObject*)receiver, member);
}
void QTimer_singleShot4(long long interval, int timerType, void* receiver, const char* member) {
	QTimer::singleShot(std::chrono::nanoseconds((long long)interval), (Qt::TimerType)(timerType), (QObject*)receiver, member);
}
void QTimer_start3(void* self, long long value) {
	((QTimer*)self)->start(std::chrono::milliseconds((long long)value));
}
void* QTimer_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QTimer::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QTimer_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QTimer::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QTimer_sender(void* self) {
	auto&& __r = ((QTimer__Go*)self)->sender();
	return (void*)(__r);
}
int QTimer_senderSignalIndex(void* self) {
	auto&& __r = ((QTimer__Go*)self)->senderSignalIndex();
	return __r;
}
int QTimer_receivers(void* self, const char* signal) {
	auto&& __r = ((QTimer__Go*)self)->receivers(signal);
	return __r;
}
int QTimer_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QTimer__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QTimer_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QTimer::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QTimer_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QTimer::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QTimer_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QTimer::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QTimer_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QTimer::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QTimer_metaObject_default(void* self) {
	auto&& __r = ((QTimer*)self)->QTimer::metaObject();
	return (void*)(__r);
}
void QTimer_timerEvent_default(void* self, void* arg1) {
	((QTimer__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)arg1);
}
int QTimer_event_default(void* self, void* event) {
	auto&& __r = ((QTimer*)self)->QTimer::event((QEvent*)event);
	return __r;
}
int QTimer_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QTimer*)self)->QTimer::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QTimer_childEvent_default(void* self, void* event) {
	((QTimer__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QTimer_customEvent_default(void* self, void* event) {
	((QTimer__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QTimer_connectNotify_default(void* self, void* signal) {
	((QTimer__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QTimer_disconnectNotify_default(void* self, void* signal) {
	((QTimer__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QTimer_connect_timeout(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTimer*)self, &QTimer::timeout, (QTimer*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTimer);
	__conns_QTimer[__id] = __c;
}
void QTimer_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTimer*)self, &QTimer::destroyed, (QTimer*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTimer);
	__conns_QTimer[__id] = __c;
}
void QTimer_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTimer*)self, &QTimer::destroyed, (QTimer*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTimer);
	__conns_QTimer[__id] = __c;
}
void QTimer_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTimer*)self, &QTimer::objectNameChanged, (QTimer*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTimer);
	__conns_QTimer[__id] = __c;
}
void QTimer__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QTimer);
		auto __it = __conns_QTimer.find(id);
		if (__it == __conns_QTimer.end()) return;
		__c = __it->second;
		__conns_QTimer.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QTimer__Go_new(long long id, void* parent) { return new QTimer__Go(id, (QObject*)parent); }
void QTimer_delete(void* self) { delete (QTimer*)self; }

}

#include <QtCore/qfileselector.h>
#include <QtCore/qobject.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qurl.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qvariant.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qfileselector.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QFileSelector__Go : public QFileSelector {
public:
	long long __goId;
	unsigned char __goOn[8] = {};
	template<typename... A> QFileSelector__Go(long long id, A&&... a) : QFileSelector(std::forward<A>(a)...), __goId(id) {}
	using QFileSelector::sender;
	using QFileSelector::senderSignalIndex;
	using QFileSelector::receivers;
	using QFileSelector::isSignalConnected;
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QFileSelector::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QFileSelector::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QFileSelector::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QFileSelector::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QFileSelector::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QFileSelector::metaObject();
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QFileSelector::event(event);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QFileSelector::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QFileSelector::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QFileSelector::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QFileSelector::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QFileSelector::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QFileSelector::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QFileSelector;
static std::mutex __connsMu_QFileSelector;

extern "C" {

void* QFileSelector_base(void* self) { return (void*)static_cast<QObject*>((QFileSelector*)self); }
void QFileSelector__set_override(void* self, int idx) { ((QFileSelector__Go*)self)->__goOn[idx] = 1; }
void* QFileSelector_fromQObject(void* o) { return (void*)qobject_cast<QFileSelector*>((QObject*)o); }
void* QFileSelector_new(long long __goid) {
	return new QFileSelector__Go(__goid);
}
void* QFileSelector_new2(long long __goid, void* parent) {
	return new QFileSelector__Go(__goid, (QObject*)parent);
}
char* QFileSelector_tr(const char* s, long long* __out_len) {
	auto&& __r = QFileSelector::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QFileSelector_select(void* self, const char* filePath, long long filePath_len, long long* __out_len) {
	QString filePath_qstr = QString::fromUtf8(filePath, filePath_len);
	auto&& __r = ((QFileSelector*)self)->select(filePath_qstr);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QFileSelector_select2(void* self, void* filePath) {
	auto&& __r = ((QFileSelector*)self)->select(*(QUrl*)filePath);
	return new QUrl(__r);
}
char** QFileSelector_extraSelectors(void* self, long long* __out_len) {
	auto&& __r = ((QFileSelector*)self)->extraSelectors();
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
void QFileSelector_setExtraSelectors(void* self, char** list, long long list_len) {
	QList<QString> list_list; list_list.reserve(list_len);
	for (long long __i = 0; __i < list_len; __i++) list_list.append(QString::fromUtf8(list[__i]));
	((QFileSelector*)self)->setExtraSelectors(list_list);
}
char** QFileSelector_allSelectors(void* self, long long* __out_len) {
	auto&& __r = ((QFileSelector*)self)->allSelectors();
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
void* QFileSelector_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QFileSelector::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QFileSelector_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QFileSelector::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QFileSelector_sender(void* self) {
	auto&& __r = ((QFileSelector__Go*)self)->sender();
	return (void*)(__r);
}
int QFileSelector_senderSignalIndex(void* self) {
	auto&& __r = ((QFileSelector__Go*)self)->senderSignalIndex();
	return __r;
}
int QFileSelector_receivers(void* self, const char* signal) {
	auto&& __r = ((QFileSelector__Go*)self)->receivers(signal);
	return __r;
}
int QFileSelector_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QFileSelector__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QFileSelector_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QFileSelector::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QFileSelector_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QFileSelector::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QFileSelector_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QFileSelector::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QFileSelector_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QFileSelector::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QFileSelector_metaObject_default(void* self) {
	auto&& __r = ((QFileSelector*)self)->QFileSelector::metaObject();
	return (void*)(__r);
}
int QFileSelector_event_default(void* self, void* event) {
	auto&& __r = ((QFileSelector*)self)->QFileSelector::event((QEvent*)event);
	return __r;
}
int QFileSelector_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QFileSelector*)self)->QFileSelector::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QFileSelector_timerEvent_default(void* self, void* event) {
	((QFileSelector__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QFileSelector_childEvent_default(void* self, void* event) {
	((QFileSelector__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QFileSelector_customEvent_default(void* self, void* event) {
	((QFileSelector__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QFileSelector_connectNotify_default(void* self, void* signal) {
	((QFileSelector__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QFileSelector_disconnectNotify_default(void* self, void* signal) {
	((QFileSelector__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QFileSelector_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFileSelector*)self, &QFileSelector::destroyed, (QFileSelector*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFileSelector);
	__conns_QFileSelector[__id] = __c;
}
void QFileSelector_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFileSelector*)self, &QFileSelector::destroyed, (QFileSelector*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFileSelector);
	__conns_QFileSelector[__id] = __c;
}
void QFileSelector_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFileSelector*)self, &QFileSelector::objectNameChanged, (QFileSelector*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFileSelector);
	__conns_QFileSelector[__id] = __c;
}
void QFileSelector__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QFileSelector);
		auto __it = __conns_QFileSelector.find(id);
		if (__it == __conns_QFileSelector.end()) return;
		__c = __it->second;
		__conns_QFileSelector.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QFileSelector__Go_new(long long id, void* parent) { return new QFileSelector__Go(id, (QObject*)parent); }
void QFileSelector_delete(void* self) { delete (QFileSelector*)self; }

}

#include <QtCore/qpoint.h>
#include <QtGui/qtransform.h>
#include <QtGui/qmatrix4x4.h>
#include "gen_qpointf.h"

extern "C" {

void* QPointF_new() {
	return new QPointF();
}
void* QPointF_new2(void* p) {
	return new QPointF(*(QPoint*)p);
}
void* QPointF_new3(double xpos, double ypos) {
	return new QPointF(xpos, ypos);
}
void* QPointF_new4(void* arg1) {
	return new QPointF(*(QPointF*)arg1);
}
double QPointF_manhattanLength(void* self) {
	auto&& __r = ((QPointF*)self)->manhattanLength();
	return __r;
}
int QPointF_isNull(void* self) {
	auto&& __r = ((QPointF*)self)->isNull();
	return __r;
}
double QPointF_x(void* self) {
	auto&& __r = ((QPointF*)self)->x();
	return __r;
}
double QPointF_y(void* self) {
	auto&& __r = ((QPointF*)self)->y();
	return __r;
}
void QPointF_setX(void* self, double x) {
	((QPointF*)self)->setX(x);
}
void QPointF_setY(void* self, double y) {
	((QPointF*)self)->setY(y);
}
void* QPointF_transposed(void* self) {
	auto&& __r = ((QPointF*)self)->transposed();
	return new QPointF(__r);
}
void* QPointF_OperatorPlusAssign(void* self, void* p) {
	auto&& __r = ((QPointF*)self)->operator+=(*(QPointF*)p);
	return (void*)&(__r);
}
void* QPointF_OperatorMinusAssign(void* self, void* p) {
	auto&& __r = ((QPointF*)self)->operator-=(*(QPointF*)p);
	return (void*)&(__r);
}
void* QPointF_OperatorMultiplyAssign(void* self, double c) {
	auto&& __r = ((QPointF*)self)->operator*=(c);
	return (void*)&(__r);
}
void* QPointF_OperatorDivideAssign(void* self, double c) {
	auto&& __r = ((QPointF*)self)->operator/=(c);
	return (void*)&(__r);
}
double QPointF_dotProduct(void* p1, void* p2) {
	auto&& __r = QPointF::dotProduct(*(QPointF*)p1, *(QPointF*)p2);
	return __r;
}
void* QPointF_toPoint(void* self) {
	auto&& __r = ((QPointF*)self)->toPoint();
	return new QPoint(__r);
}
void* QPointF_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QPointF*)self)->operator=(*(QPointF*)arg1);
	return (void*)&(__r);
}
int QPointF_OperatorEqual(void* self, void* rhs) {
	auto&& __r = (*(QPointF*)self) == (*(QPointF*)rhs);
	return __r;
}
int QPointF_OperatorNotEqual(void* self, void* rhs) {
	auto&& __r = (*(QPointF*)self) != (*(QPointF*)rhs);
	return __r;
}
int QPointF_OperatorEqual2(void* self, void* rhs) {
	auto&& __r = (*(QPointF*)self) == (*(QPoint*)rhs);
	return __r;
}
int QPointF_OperatorNotEqual2(void* self, void* rhs) {
	auto&& __r = (*(QPointF*)self) != (*(QPoint*)rhs);
	return __r;
}
void* QPointF_OperatorPlus(void* self, void* p2) {
	auto&& __r = (*(QPointF*)self) + (*(QPointF*)p2);
	return new QPointF(__r);
}
void* QPointF_OperatorMinus(void* self, void* p2) {
	auto&& __r = (*(QPointF*)self) - (*(QPointF*)p2);
	return new QPointF(__r);
}
void* QPointF_OperatorMultiply(void* self, double c) {
	auto&& __r = (*(QPointF*)self) * (c);
	return new QPointF(__r);
}
void* QPointF_OperatorPlus2(void* self) {
	auto&& __r = +(*(QPointF*)self);
	return new QPointF(__r);
}
void* QPointF_OperatorMinus2(void* self) {
	auto&& __r = -(*(QPointF*)self);
	return new QPointF(__r);
}
void* QPointF_OperatorDivide(void* self, double divisor) {
	auto&& __r = (*(QPointF*)self) / (divisor);
	return new QPointF(__r);
}
void* QPointF_OperatorMultiply2(void* self, void* m) {
	auto&& __r = (*(QPointF*)self) * (*(QTransform*)m);
	return new QPointF(__r);
}
void* QPointF_OperatorMultiply3(void* self, void* matrix) {
	auto&& __r = (*(QPointF*)self) * (*(QMatrix4x4*)matrix);
	return new QPointF(__r);
}
void QPointF_delete(void* self) { delete (QPointF*)self; }

}

#include <QtCore/qpermissions.h>
#include "gen_qcontactspermission.h"

extern "C" {

void* QContactsPermission_new() {
	return new QContactsPermission();
}
void* QContactsPermission_new2(void* other) {
	return new QContactsPermission(*(QContactsPermission*)other);
}
void QContactsPermission_setAccessMode(void* self, int mode) {
	((QContactsPermission*)self)->setAccessMode((QContactsPermission::AccessMode)(mode));
}
int QContactsPermission_accessMode(void* self) {
	auto&& __r = ((QContactsPermission*)self)->accessMode();
	return (int)(__r);
}
void* QContactsPermission_OperatorAssign(void* self, void* other) {
	auto&& __r = ((QContactsPermission*)self)->operator=(*(QContactsPermission*)other);
	return (void*)&(__r);
}
void QContactsPermission_swap(void* self, void* other) {
	((QContactsPermission*)self)->swap(*(QContactsPermission*)other);
}
void QContactsPermission_delete(void* self) { delete (QContactsPermission*)self; }

}

#include <QtCore/qrect.h>
#include <QtCore/qpoint.h>
#include <QtCore/qsize.h>
#include <QtCore/qmargins.h>
#include "gen_qrect.h"

extern "C" {

void* QRect_new() {
	return new QRect();
}
void* QRect_new2(void* topleft, void* bottomright) {
	return new QRect(*(QPoint*)topleft, *(QPoint*)bottomright);
}
void* QRect_new3(void* topleft, void* size) {
	return new QRect(*(QPoint*)topleft, *(QSize*)size);
}
void* QRect_new4(int left, int top, int width, int height) {
	return new QRect(left, top, width, height);
}
void* QRect_new5(void* arg1) {
	return new QRect(*(QRect*)arg1);
}
int QRect_isNull(void* self) {
	auto&& __r = ((QRect*)self)->isNull();
	return __r;
}
int QRect_isEmpty(void* self) {
	auto&& __r = ((QRect*)self)->isEmpty();
	return __r;
}
int QRect_isValid(void* self) {
	auto&& __r = ((QRect*)self)->isValid();
	return __r;
}
int QRect_left(void* self) {
	auto&& __r = ((QRect*)self)->left();
	return __r;
}
int QRect_top(void* self) {
	auto&& __r = ((QRect*)self)->top();
	return __r;
}
int QRect_right(void* self) {
	auto&& __r = ((QRect*)self)->right();
	return __r;
}
int QRect_bottom(void* self) {
	auto&& __r = ((QRect*)self)->bottom();
	return __r;
}
void* QRect_normalized(void* self) {
	auto&& __r = ((QRect*)self)->normalized();
	return new QRect(__r);
}
int QRect_x(void* self) {
	auto&& __r = ((QRect*)self)->x();
	return __r;
}
int QRect_y(void* self) {
	auto&& __r = ((QRect*)self)->y();
	return __r;
}
void QRect_setLeft(void* self, int pos) {
	((QRect*)self)->setLeft(pos);
}
void QRect_setTop(void* self, int pos) {
	((QRect*)self)->setTop(pos);
}
void QRect_setRight(void* self, int pos) {
	((QRect*)self)->setRight(pos);
}
void QRect_setBottom(void* self, int pos) {
	((QRect*)self)->setBottom(pos);
}
void QRect_setX(void* self, int x) {
	((QRect*)self)->setX(x);
}
void QRect_setY(void* self, int y) {
	((QRect*)self)->setY(y);
}
void QRect_setTopLeft(void* self, void* p) {
	((QRect*)self)->setTopLeft(*(QPoint*)p);
}
void QRect_setBottomRight(void* self, void* p) {
	((QRect*)self)->setBottomRight(*(QPoint*)p);
}
void QRect_setTopRight(void* self, void* p) {
	((QRect*)self)->setTopRight(*(QPoint*)p);
}
void QRect_setBottomLeft(void* self, void* p) {
	((QRect*)self)->setBottomLeft(*(QPoint*)p);
}
void* QRect_topLeft(void* self) {
	auto&& __r = ((QRect*)self)->topLeft();
	return new QPoint(__r);
}
void* QRect_bottomRight(void* self) {
	auto&& __r = ((QRect*)self)->bottomRight();
	return new QPoint(__r);
}
void* QRect_topRight(void* self) {
	auto&& __r = ((QRect*)self)->topRight();
	return new QPoint(__r);
}
void* QRect_bottomLeft(void* self) {
	auto&& __r = ((QRect*)self)->bottomLeft();
	return new QPoint(__r);
}
void* QRect_center(void* self) {
	auto&& __r = ((QRect*)self)->center();
	return new QPoint(__r);
}
void QRect_moveLeft(void* self, int pos) {
	((QRect*)self)->moveLeft(pos);
}
void QRect_moveTop(void* self, int pos) {
	((QRect*)self)->moveTop(pos);
}
void QRect_moveRight(void* self, int pos) {
	((QRect*)self)->moveRight(pos);
}
void QRect_moveBottom(void* self, int pos) {
	((QRect*)self)->moveBottom(pos);
}
void QRect_moveTopLeft(void* self, void* p) {
	((QRect*)self)->moveTopLeft(*(QPoint*)p);
}
void QRect_moveBottomRight(void* self, void* p) {
	((QRect*)self)->moveBottomRight(*(QPoint*)p);
}
void QRect_moveTopRight(void* self, void* p) {
	((QRect*)self)->moveTopRight(*(QPoint*)p);
}
void QRect_moveBottomLeft(void* self, void* p) {
	((QRect*)self)->moveBottomLeft(*(QPoint*)p);
}
void QRect_moveCenter(void* self, void* p) {
	((QRect*)self)->moveCenter(*(QPoint*)p);
}
void QRect_translate(void* self, int dx, int dy) {
	((QRect*)self)->translate(dx, dy);
}
void QRect_translate2(void* self, void* p) {
	((QRect*)self)->translate(*(QPoint*)p);
}
void* QRect_translated(void* self, int dx, int dy) {
	auto&& __r = ((QRect*)self)->translated(dx, dy);
	return new QRect(__r);
}
void* QRect_translated2(void* self, void* p) {
	auto&& __r = ((QRect*)self)->translated(*(QPoint*)p);
	return new QRect(__r);
}
void* QRect_transposed(void* self) {
	auto&& __r = ((QRect*)self)->transposed();
	return new QRect(__r);
}
void QRect_moveTo(void* self, int x, int t) {
	((QRect*)self)->moveTo(x, t);
}
void QRect_moveTo2(void* self, void* p) {
	((QRect*)self)->moveTo(*(QPoint*)p);
}
void QRect_setRect(void* self, int x, int y, int w, int h) {
	((QRect*)self)->setRect(x, y, w, h);
}
void QRect_getRect(void* self, int* x, int* y, int* w, int* h) {
	int x_local{};
	int y_local{};
	int w_local{};
	int h_local{};
	((QRect*)self)->getRect(&x_local, &y_local, &w_local, &h_local);
	*x = x_local;
	*y = y_local;
	*w = w_local;
	*h = h_local;
}
void QRect_setCoords(void* self, int x1, int y1, int x2, int y2) {
	((QRect*)self)->setCoords(x1, y1, x2, y2);
}
void QRect_getCoords(void* self, int* x1, int* y1, int* x2, int* y2) {
	int x1_local{};
	int y1_local{};
	int x2_local{};
	int y2_local{};
	((QRect*)self)->getCoords(&x1_local, &y1_local, &x2_local, &y2_local);
	*x1 = x1_local;
	*y1 = y1_local;
	*x2 = x2_local;
	*y2 = y2_local;
}
void QRect_adjust(void* self, int x1, int y1, int x2, int y2) {
	((QRect*)self)->adjust(x1, y1, x2, y2);
}
void* QRect_adjusted(void* self, int x1, int y1, int x2, int y2) {
	auto&& __r = ((QRect*)self)->adjusted(x1, y1, x2, y2);
	return new QRect(__r);
}
void* QRect_size(void* self) {
	auto&& __r = ((QRect*)self)->size();
	return new QSize(__r);
}
int QRect_width(void* self) {
	auto&& __r = ((QRect*)self)->width();
	return __r;
}
int QRect_height(void* self) {
	auto&& __r = ((QRect*)self)->height();
	return __r;
}
void QRect_setWidth(void* self, int w) {
	((QRect*)self)->setWidth(w);
}
void QRect_setHeight(void* self, int h) {
	((QRect*)self)->setHeight(h);
}
void QRect_setSize(void* self, void* s) {
	((QRect*)self)->setSize(*(QSize*)s);
}
void* QRect_OperatorBitwiseOr(void* self, void* r) {
	auto&& __r = ((QRect*)self)->operator|(*(QRect*)r);
	return new QRect(__r);
}
void* QRect_OperatorBitwiseAnd(void* self, void* r) {
	auto&& __r = ((QRect*)self)->operator&(*(QRect*)r);
	return new QRect(__r);
}
void* QRect_OperatorBitwiseOrAssign(void* self, void* r) {
	auto&& __r = ((QRect*)self)->operator|=(*(QRect*)r);
	return (void*)&(__r);
}
void* QRect_OperatorBitwiseAndAssign(void* self, void* r) {
	auto&& __r = ((QRect*)self)->operator&=(*(QRect*)r);
	return (void*)&(__r);
}
int QRect_contains(void* self, void* r) {
	auto&& __r = ((QRect*)self)->contains(*(QRect*)r);
	return __r;
}
int QRect_contains2(void* self, void* p) {
	auto&& __r = ((QRect*)self)->contains(*(QPoint*)p);
	return __r;
}
int QRect_contains3(void* self, int x, int y) {
	auto&& __r = ((QRect*)self)->contains(x, y);
	return __r;
}
int QRect_contains4(void* self, int x, int y, int proper) {
	auto&& __r = ((QRect*)self)->contains(x, y, proper);
	return __r;
}
void* QRect_united(void* self, void* other) {
	auto&& __r = ((QRect*)self)->united(*(QRect*)other);
	return new QRect(__r);
}
void* QRect_intersected(void* self, void* other) {
	auto&& __r = ((QRect*)self)->intersected(*(QRect*)other);
	return new QRect(__r);
}
int QRect_intersects(void* self, void* r) {
	auto&& __r = ((QRect*)self)->intersects(*(QRect*)r);
	return __r;
}
void* QRect_marginsAdded(void* self, void* margins) {
	auto&& __r = ((QRect*)self)->marginsAdded(*(QMargins*)margins);
	return new QRect(__r);
}
void* QRect_marginsRemoved(void* self, void* margins) {
	auto&& __r = ((QRect*)self)->marginsRemoved(*(QMargins*)margins);
	return new QRect(__r);
}
void* QRect_OperatorPlusAssign(void* self, void* margins) {
	auto&& __r = ((QRect*)self)->operator+=(*(QMargins*)margins);
	return (void*)&(__r);
}
void* QRect_OperatorMinusAssign(void* self, void* margins) {
	auto&& __r = ((QRect*)self)->operator-=(*(QMargins*)margins);
	return (void*)&(__r);
}
void* QRect_span(void* p1, void* p2) {
	auto&& __r = QRect::span(*(QPoint*)p1, *(QPoint*)p2);
	return new QRect(__r);
}
void* QRect_toRectF(void* self) {
	auto&& __r = ((QRect*)self)->toRectF();
	return new QRectF(__r);
}
void* QRect_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QRect*)self)->operator=(*(QRect*)arg1);
	return (void*)&(__r);
}
int QRect_contains5(void* self, void* r, int proper) {
	auto&& __r = ((QRect*)self)->contains(*(QRect*)r, proper);
	return __r;
}
int QRect_contains6(void* self, void* p, int proper) {
	auto&& __r = ((QRect*)self)->contains(*(QPoint*)p, proper);
	return __r;
}
int QRect_OperatorEqual(void* self, void* rhs) {
	auto&& __r = (*(QRect*)self) == (*(QRect*)rhs);
	return __r;
}
int QRect_OperatorNotEqual(void* self, void* rhs) {
	auto&& __r = (*(QRect*)self) != (*(QRect*)rhs);
	return __r;
}
void* QRect_OperatorPlus(void* self, void* margins) {
	auto&& __r = (*(QRect*)self) + (*(QMargins*)margins);
	return new QRect(__r);
}
void* QRect_OperatorMinus(void* self, void* rhs) {
	auto&& __r = (*(QRect*)self) - (*(QMargins*)rhs);
	return new QRect(__r);
}
int QRect_OperatorEqual2(void* self, void* rhs) {
	auto&& __r = (*(QRect*)self) == (*(QRectF*)rhs);
	return __r;
}
int QRect_OperatorNotEqual2(void* self, void* rhs) {
	auto&& __r = (*(QRect*)self) != (*(QRectF*)rhs);
	return __r;
}
void QRect_delete(void* self) { delete (QRect*)self; }

}

#include <QtCore/qthreadstorage.h>
#include "gen_qthreadstoragedata.h"

extern "C" {

int QThreadStorageData_id_get(void* self) {
	auto __r = ((QThreadStorageData*)self)->id;
	return __r;
}
void QThreadStorageData_id_set(void* self, int id) {
	((QThreadStorageData*)self)->id = id;
}
void QThreadStorageData_delete(void* self) { delete (QThreadStorageData*)self; }

}

#include <QtGui/qpixmap.h>
#include <QtGui/qpaintdevice.h>
#include <QtCore/qsize.h>
#include <QtCore/qstring.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qvariant.h>
#include <QtCore/qrect.h>
#include <QtGui/qcolor.h>
#include <QtGui/qbitmap.h>
#include <QtGui/qtransform.h>
#include <QtGui/qimage.h>
#include <QtGui/qimagereader.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qiodevice.h>
#include <QtGui/qregion.h>
#include <QtGui/qpaintengine.h>
#include <QtGui/qpainter.h>
#include <QtCore/qpoint.h>
#include "gen_qpixmap.h"
#include "gen_purego_cb.h"

class QPixmap__Go : public QPixmap {
public:
	long long __goId;
	unsigned char __goOn[6] = {};
	template<typename... A> QPixmap__Go(long long id, A&&... a) : QPixmap(std::forward<A>(a)...), __goId(id) {}
	using QPixmap::getDecodedMetricF;
	int cppgoBase_metric(int __a0) { return QPixmap::metric((QPaintDevice::PaintDeviceMetric)__a0); }
	void cppgoBase_initPainter(QPainter* __a0) { QPixmap::initPainter(__a0); }
	QPaintDevice* cppgoBase_redirected(QPoint* __a0) { return QPixmap::redirected(__a0); }
	QPainter* cppgoBase_sharedPainter() { return QPixmap::sharedPainter(); }
	int devType() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8659679893868524044LL, &__h); if (__h) return __r; }
		return QPixmap::devType();
	}
	QPaintEngine* paintEngine() const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 6398007404756346381LL, &__h); if (__h) return (QPaintEngine*)__r; }
		return QPixmap::paintEngine();
	}
	int metric(QPaintDevice::PaintDeviceMetric arg1) const override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 427629649267311777LL, (int)(arg1), &__h); if (__h) return __r; }
		return QPixmap::metric(arg1);
	}
	void initPainter(QPainter* painter) const override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6477500296980050210LL, (void*)painter, &__h); if (__h) return; }
		QPixmap::initPainter(painter);
	}
	QPaintDevice* redirected(QPoint* offset) const override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 4127240633310655048LL, (void*)offset, &__h); if (__h) return (QPaintDevice*)__r; }
		return QPixmap::redirected(offset);
	}
	QPainter* sharedPainter() const override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 5792771328741774991LL, &__h); if (__h) return (QPainter*)__r; }
		return QPixmap::sharedPainter();
	}
};

extern "C" {

void* QPixmap_base(void* self) { return (void*)static_cast<QPaintDevice*>((QPixmap*)self); }
void QPixmap__set_override(void* self, int idx) { ((QPixmap__Go*)self)->__goOn[idx] = 1; }
void* QPixmap_new(long long __goid) {
	return new QPixmap__Go(__goid);
}
void* QPixmap_new2(long long __goid, int w, int h) {
	return new QPixmap__Go(__goid, w, h);
}
void* QPixmap_new3(long long __goid, void* arg1) {
	return new QPixmap__Go(__goid, *(QSize*)arg1);
}
void* QPixmap_new4(long long __goid, const char* fileName, long long fileName_len) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	return new QPixmap__Go(__goid, fileName_qstr);
}
void* QPixmap_new5(long long __goid, void* arg1) {
	return new QPixmap__Go(__goid, *(QPixmap*)arg1);
}
void* QPixmap_new6(long long __goid, const char* fileName, long long fileName_len, const char* format) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	return new QPixmap__Go(__goid, fileName_qstr, format);
}
void* QPixmap_new7(long long __goid, const char* fileName, long long fileName_len, const char* format, int flags) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	return new QPixmap__Go(__goid, fileName_qstr, format, QFlags<Qt::ImageConversionFlag>(flags));
}
void* QPixmap_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QPixmap*)self)->operator=(*(QPixmap*)arg1);
	return (void*)&(__r);
}
void QPixmap_swap(void* self, void* other) {
	((QPixmap*)self)->swap(*(QPixmap*)other);
}
void* QPixmap_ToQVariant(void* self) {
	auto&& __r = static_cast<QVariant>(*(QPixmap*)self);
	return new QVariant(__r);
}
int QPixmap_isNull(void* self) {
	auto&& __r = ((QPixmap*)self)->isNull();
	return __r;
}
int QPixmap_width(void* self) {
	auto&& __r = ((QPixmap*)self)->width();
	return __r;
}
int QPixmap_height(void* self) {
	auto&& __r = ((QPixmap*)self)->height();
	return __r;
}
void* QPixmap_size(void* self) {
	auto&& __r = ((QPixmap*)self)->size();
	return new QSize(__r);
}
void* QPixmap_rect(void* self) {
	auto&& __r = ((QPixmap*)self)->rect();
	return new QRect(__r);
}
int QPixmap_depth(void* self) {
	auto&& __r = ((QPixmap*)self)->depth();
	return __r;
}
int QPixmap_defaultDepth() {
	auto&& __r = QPixmap::defaultDepth();
	return __r;
}
void QPixmap_fill(void* self) {
	((QPixmap*)self)->fill();
}
void* QPixmap_mask(void* self) {
	auto&& __r = ((QPixmap*)self)->mask();
	return new QBitmap(__r);
}
void QPixmap_setMask(void* self, void* arg1) {
	((QPixmap*)self)->setMask(*(QBitmap*)arg1);
}
double QPixmap_devicePixelRatio(void* self) {
	auto&& __r = ((QPixmap*)self)->devicePixelRatio();
	return __r;
}
void QPixmap_setDevicePixelRatio(void* self, double scaleFactor) {
	((QPixmap*)self)->setDevicePixelRatio(scaleFactor);
}
void* QPixmap_deviceIndependentSize(void* self) {
	auto&& __r = ((QPixmap*)self)->deviceIndependentSize();
	return new QSizeF(__r);
}
int QPixmap_hasAlpha(void* self) {
	auto&& __r = ((QPixmap*)self)->hasAlpha();
	return __r;
}
int QPixmap_hasAlphaChannel(void* self) {
	auto&& __r = ((QPixmap*)self)->hasAlphaChannel();
	return __r;
}
void* QPixmap_createHeuristicMask(void* self) {
	auto&& __r = ((QPixmap*)self)->createHeuristicMask();
	return new QBitmap(__r);
}
void* QPixmap_createMaskFromColor(void* self, void* maskColor) {
	auto&& __r = ((QPixmap*)self)->createMaskFromColor(*(QColor*)maskColor);
	return new QBitmap(__r);
}
void* QPixmap_scaled(void* self, int w, int h) {
	auto&& __r = ((QPixmap*)self)->scaled(w, h);
	return new QPixmap(__r);
}
void* QPixmap_scaled2(void* self, void* s) {
	auto&& __r = ((QPixmap*)self)->scaled(*(QSize*)s);
	return new QPixmap(__r);
}
void* QPixmap_scaledToWidth(void* self, int w) {
	auto&& __r = ((QPixmap*)self)->scaledToWidth(w);
	return new QPixmap(__r);
}
void* QPixmap_scaledToHeight(void* self, int h) {
	auto&& __r = ((QPixmap*)self)->scaledToHeight(h);
	return new QPixmap(__r);
}
void* QPixmap_transformed(void* self, void* arg1) {
	auto&& __r = ((QPixmap*)self)->transformed(*(QTransform*)arg1);
	return new QPixmap(__r);
}
void* QPixmap_trueMatrix(void* m, int w, int h) {
	auto&& __r = QPixmap::trueMatrix(*(QTransform*)m, w, h);
	return new QTransform(__r);
}
void* QPixmap_toImage(void* self) {
	auto&& __r = ((QPixmap*)self)->toImage();
	return new QImage(__r);
}
void* QPixmap_fromImage(void* image) {
	auto&& __r = QPixmap::fromImage(*(QImage*)image);
	return new QPixmap(__r);
}
void* QPixmap_fromImageReader(void* imageReader) {
	auto&& __r = QPixmap::fromImageReader((QImageReader*)imageReader);
	return new QPixmap(__r);
}
int QPixmap_load(void* self, const char* fileName, long long fileName_len) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	auto&& __r = ((QPixmap*)self)->load(fileName_qstr);
	return __r;
}
int QPixmap_loadFromData(void* self, const unsigned char* buf, unsigned int len_) {
	auto&& __r = ((QPixmap*)self)->loadFromData(buf, len_);
	return __r;
}
int QPixmap_loadFromData2(void* self, void* data) {
	auto&& __r = ((QPixmap*)self)->loadFromData(*(QByteArray*)data);
	return __r;
}
int QPixmap_save(void* self, const char* fileName, long long fileName_len) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	auto&& __r = ((QPixmap*)self)->save(fileName_qstr);
	return __r;
}
int QPixmap_save2(void* self, void* device) {
	auto&& __r = ((QPixmap*)self)->save((QIODevice*)device);
	return __r;
}
int QPixmap_convertFromImage(void* self, void* img) {
	auto&& __r = ((QPixmap*)self)->convertFromImage(*(QImage*)img);
	return __r;
}
void* QPixmap_copy(void* self, int x, int y, int width, int height) {
	auto&& __r = ((QPixmap*)self)->copy(x, y, width, height);
	return new QPixmap(__r);
}
void* QPixmap_copy2(void* self) {
	auto&& __r = ((QPixmap*)self)->copy();
	return new QPixmap(__r);
}
void QPixmap_scroll(void* self, int dx, int dy, int x, int y, int width, int height) {
	((QPixmap*)self)->scroll(dx, dy, x, y, width, height);
}
void QPixmap_scroll2(void* self, int dx, int dy, void* rect) {
	((QPixmap*)self)->scroll(dx, dy, *(QRect*)rect);
}
long long QPixmap_cacheKey(void* self) {
	auto&& __r = ((QPixmap*)self)->cacheKey();
	return __r;
}
int QPixmap_isDetached(void* self) {
	auto&& __r = ((QPixmap*)self)->isDetached();
	return __r;
}
void QPixmap_detach(void* self) {
	((QPixmap*)self)->detach();
}
int QPixmap_isQBitmap(void* self) {
	auto&& __r = ((QPixmap*)self)->isQBitmap();
	return __r;
}
int QPixmap_OperatorNot(void* self) {
	auto&& __r = ((QPixmap*)self)->operator!();
	return __r;
}
double QPixmap_devicePixelRatioFScale() {
	auto&& __r = QPixmap::devicePixelRatioFScale();
	return __r;
}
int QPixmap_encodeMetricF(int metric, double value) {
	auto&& __r = QPixmap::encodeMetricF((QPaintDevice::PaintDeviceMetric)(metric), value);
	return __r;
}
double QPixmap_getDecodedMetricF(void* self, int metricA, int metricB) {
	auto&& __r = ((QPixmap__Go*)self)->getDecodedMetricF((QPaintDevice::PaintDeviceMetric)(metricA), (QPaintDevice::PaintDeviceMetric)(metricB));
	return __r;
}
void QPixmap_fill2(void* self, void* fillColor) {
	((QPixmap*)self)->fill(*(QColor*)fillColor);
}
void* QPixmap_createHeuristicMask2(void* self, int clipTight) {
	auto&& __r = ((QPixmap*)self)->createHeuristicMask(clipTight);
	return new QBitmap(__r);
}
void* QPixmap_createMaskFromColor2(void* self, void* maskColor, int mode) {
	auto&& __r = ((QPixmap*)self)->createMaskFromColor(*(QColor*)maskColor, (Qt::MaskMode)(mode));
	return new QBitmap(__r);
}
void* QPixmap_scaled3(void* self, int w, int h, int aspectMode) {
	auto&& __r = ((QPixmap*)self)->scaled(w, h, (Qt::AspectRatioMode)(aspectMode));
	return new QPixmap(__r);
}
void* QPixmap_scaled4(void* self, int w, int h, int aspectMode, int mode) {
	auto&& __r = ((QPixmap*)self)->scaled(w, h, (Qt::AspectRatioMode)(aspectMode), (Qt::TransformationMode)(mode));
	return new QPixmap(__r);
}
void* QPixmap_scaled5(void* self, void* s, int aspectMode) {
	auto&& __r = ((QPixmap*)self)->scaled(*(QSize*)s, (Qt::AspectRatioMode)(aspectMode));
	return new QPixmap(__r);
}
void* QPixmap_scaled6(void* self, void* s, int aspectMode, int mode) {
	auto&& __r = ((QPixmap*)self)->scaled(*(QSize*)s, (Qt::AspectRatioMode)(aspectMode), (Qt::TransformationMode)(mode));
	return new QPixmap(__r);
}
void* QPixmap_scaledToWidth2(void* self, int w, int mode) {
	auto&& __r = ((QPixmap*)self)->scaledToWidth(w, (Qt::TransformationMode)(mode));
	return new QPixmap(__r);
}
void* QPixmap_scaledToHeight2(void* self, int h, int mode) {
	auto&& __r = ((QPixmap*)self)->scaledToHeight(h, (Qt::TransformationMode)(mode));
	return new QPixmap(__r);
}
void* QPixmap_transformed2(void* self, void* arg1, int mode) {
	auto&& __r = ((QPixmap*)self)->transformed(*(QTransform*)arg1, (Qt::TransformationMode)(mode));
	return new QPixmap(__r);
}
void* QPixmap_fromImage2(void* image, int flags) {
	auto&& __r = QPixmap::fromImage(*(QImage*)image, QFlags<Qt::ImageConversionFlag>(flags));
	return new QPixmap(__r);
}
void* QPixmap_fromImageReader2(void* imageReader, int flags) {
	auto&& __r = QPixmap::fromImageReader((QImageReader*)imageReader, QFlags<Qt::ImageConversionFlag>(flags));
	return new QPixmap(__r);
}
int QPixmap_load2(void* self, const char* fileName, long long fileName_len, const char* format) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	auto&& __r = ((QPixmap*)self)->load(fileName_qstr, format);
	return __r;
}
int QPixmap_load3(void* self, const char* fileName, long long fileName_len, const char* format, int flags) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	auto&& __r = ((QPixmap*)self)->load(fileName_qstr, format, QFlags<Qt::ImageConversionFlag>(flags));
	return __r;
}
int QPixmap_loadFromData3(void* self, const unsigned char* buf, unsigned int len_, const char* format) {
	auto&& __r = ((QPixmap*)self)->loadFromData(buf, len_, format);
	return __r;
}
int QPixmap_loadFromData4(void* self, const unsigned char* buf, unsigned int len_, const char* format, int flags) {
	auto&& __r = ((QPixmap*)self)->loadFromData(buf, len_, format, QFlags<Qt::ImageConversionFlag>(flags));
	return __r;
}
int QPixmap_loadFromData5(void* self, void* data, const char* format) {
	auto&& __r = ((QPixmap*)self)->loadFromData(*(QByteArray*)data, format);
	return __r;
}
int QPixmap_loadFromData6(void* self, void* data, const char* format, int flags) {
	auto&& __r = ((QPixmap*)self)->loadFromData(*(QByteArray*)data, format, QFlags<Qt::ImageConversionFlag>(flags));
	return __r;
}
int QPixmap_save3(void* self, const char* fileName, long long fileName_len, const char* format) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	auto&& __r = ((QPixmap*)self)->save(fileName_qstr, format);
	return __r;
}
int QPixmap_save4(void* self, const char* fileName, long long fileName_len, const char* format, int quality) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	auto&& __r = ((QPixmap*)self)->save(fileName_qstr, format, quality);
	return __r;
}
int QPixmap_save5(void* self, void* device, const char* format) {
	auto&& __r = ((QPixmap*)self)->save((QIODevice*)device, format);
	return __r;
}
int QPixmap_save6(void* self, void* device, const char* format, int quality) {
	auto&& __r = ((QPixmap*)self)->save((QIODevice*)device, format, quality);
	return __r;
}
int QPixmap_convertFromImage2(void* self, void* img, int flags) {
	auto&& __r = ((QPixmap*)self)->convertFromImage(*(QImage*)img, QFlags<Qt::ImageConversionFlag>(flags));
	return __r;
}
void* QPixmap_copy3(void* self, void* rect) {
	auto&& __r = ((QPixmap*)self)->copy(*(QRect*)rect);
	return new QPixmap(__r);
}
void QPixmap_scroll3(void* self, int dx, int dy, int x, int y, int width, int height, void* exposed) {
	((QPixmap*)self)->scroll(dx, dy, x, y, width, height, (QRegion*)exposed);
}
void QPixmap_scroll4(void* self, int dx, int dy, void* rect, void* exposed) {
	((QPixmap*)self)->scroll(dx, dy, *(QRect*)rect, (QRegion*)exposed);
}
void* QPixmap_toQVariant(void* self) {
	auto&& __r = QVariant::fromValue(*(QPixmap*)self);
	return new QVariant(__r);
}
int QPixmap_devType_default(void* self) {
	auto&& __r = ((QPixmap*)self)->QPixmap::devType();
	return __r;
}
void* QPixmap_paintEngine_default(void* self) {
	auto&& __r = ((QPixmap*)self)->QPixmap::paintEngine();
	return (void*)(__r);
}
int QPixmap_metric_default(void* self, int arg1) {
	auto&& __r = ((QPixmap__Go*)self)->cppgoBase_metric(arg1);
	return __r;
}
void QPixmap_initPainter_default(void* self, void* painter) {
	((QPixmap__Go*)self)->cppgoBase_initPainter((QPainter*)painter);
}
void* QPixmap_redirected_default(void* self, void* offset) {
	auto&& __r = ((QPixmap__Go*)self)->cppgoBase_redirected((QPoint*)offset);
	return (void*)(__r);
}
void* QPixmap_sharedPainter_default(void* self) {
	auto&& __r = ((QPixmap__Go*)self)->cppgoBase_sharedPainter();
	return (void*)(__r);
}
void* QPixmap__Go_new(long long id) { return new QPixmap__Go(id); }
void QPixmap_delete(void* self) { delete (QPixmap*)self; }

}

#include <QtGui/qevent.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qnamespace.h>
#include "gen_qfocusevent.h"

extern "C" {

void* QFocusEvent_base(void* self) { return (void*)static_cast<QEvent*>((QFocusEvent*)self); }
void* QFocusEvent_new(int type_) {
	return new QFocusEvent((QEvent::Type)(type_));
}
void* QFocusEvent_new2(int type_, int reason) {
	return new QFocusEvent((QEvent::Type)(type_), (Qt::FocusReason)(reason));
}
int QFocusEvent_gotFocus(void* self) {
	auto&& __r = ((QFocusEvent*)self)->gotFocus();
	return __r;
}
int QFocusEvent_lostFocus(void* self) {
	auto&& __r = ((QFocusEvent*)self)->lostFocus();
	return __r;
}
int QFocusEvent_reason(void* self) {
	auto&& __r = ((QFocusEvent*)self)->reason();
	return (int)(__r);
}
int QFocusEvent_registerEventType() {
	auto&& __r = QFocusEvent::registerEventType();
	return __r;
}
int QFocusEvent_registerEventType2(int hint) {
	auto&& __r = QFocusEvent::registerEventType(hint);
	return __r;
}
void QFocusEvent_delete(void* self) { delete (QFocusEvent*)self; }

}

#include <QtGui/qevent.h>
#include <QtCore/qcoreevent.h>
#include <QtGui/qkeysequence.h>
#include <QtGui/qshortcut.h>
#include "gen_qshortcutevent.h"

extern "C" {

void* QShortcutEvent_base(void* self) { return (void*)static_cast<QEvent*>((QShortcutEvent*)self); }
void* QShortcutEvent_new(void* key, int id) {
	return new QShortcutEvent(*(QKeySequence*)key, id);
}
void* QShortcutEvent_new2(void* key) {
	return new QShortcutEvent(*(QKeySequence*)key);
}
void* QShortcutEvent_new3(void* key, int id, int ambiguous) {
	return new QShortcutEvent(*(QKeySequence*)key, id, ambiguous);
}
void* QShortcutEvent_new4(void* key, void* shortcut) {
	return new QShortcutEvent(*(QKeySequence*)key, (QShortcut*)shortcut);
}
void* QShortcutEvent_new5(void* key, void* shortcut, int ambiguous) {
	return new QShortcutEvent(*(QKeySequence*)key, (QShortcut*)shortcut, ambiguous);
}
void* QShortcutEvent_key(void* self) {
	auto&& __r = ((QShortcutEvent*)self)->key();
	return (void*)&(__r);
}
int QShortcutEvent_shortcutId(void* self) {
	auto&& __r = ((QShortcutEvent*)self)->shortcutId();
	return __r;
}
int QShortcutEvent_isAmbiguous(void* self) {
	auto&& __r = ((QShortcutEvent*)self)->isAmbiguous();
	return __r;
}
int QShortcutEvent_registerEventType() {
	auto&& __r = QShortcutEvent::registerEventType();
	return __r;
}
int QShortcutEvent_registerEventType2(int hint) {
	auto&& __r = QShortcutEvent::registerEventType(hint);
	return __r;
}
void QShortcutEvent_delete(void* self) { delete (QShortcutEvent*)self; }

}

#include <QtGui/qtextoption.h>
#include <QtCore/qchar.h>
#include "gen_qtextoption__tab.h"

extern "C" {

void* QTextOption__Tab_new() {
	return new QTextOption::Tab();
}
void* QTextOption__Tab_new2(double pos, int tabType) {
	return new QTextOption::Tab(pos, (QTextOption::TabType)(tabType));
}
void* QTextOption__Tab_new3(double pos, int tabType, void* delim) {
	return new QTextOption::Tab(pos, (QTextOption::TabType)(tabType), *(QChar*)delim);
}
int QTextOption__Tab_OperatorEqual(void* self, void* other) {
	auto&& __r = ((QTextOption::Tab*)self)->operator==(*(QTextOption::Tab*)other);
	return __r;
}
int QTextOption__Tab_OperatorNotEqual(void* self, void* other) {
	auto&& __r = ((QTextOption::Tab*)self)->operator!=(*(QTextOption::Tab*)other);
	return __r;
}
double QTextOption__Tab_position_get(void* self) {
	auto __r = ((QTextOption::Tab*)self)->position;
	return __r;
}
void QTextOption__Tab_position_set(void* self, double position) {
	((QTextOption::Tab*)self)->position = position;
}
int QTextOption__Tab_type_get(void* self) {
	auto __r = ((QTextOption::Tab*)self)->type;
	return (int)(__r);
}
void QTextOption__Tab_type_set(void* self, int type_) {
	((QTextOption::Tab*)self)->type = (QTextOption::TabType)(type_);
}
void* QTextOption__Tab_delimiter_get(void* self) {
	auto __r = ((QTextOption::Tab*)self)->delimiter;
	return new QChar(__r);
}
void QTextOption__Tab_delimiter_set(void* self, void* delimiter) {
	((QTextOption::Tab*)self)->delimiter = *(QChar*)delimiter;
}
void QTextOption__Tab_delete(void* self) { delete (QTextOption::Tab*)self; }

}

#include <QtGui/qaccessible.h>
#include <QtCore/qobject.h>
#include <QtGui/qwindow.h>
#include <QtGui/qaccessible_base.h>
#include <QtCore/qstring.h>
#include <QtCore/qrect.h>
#include <QtGui/qcolor.h>
#include "gen_qaccessibleinterface.h"

extern "C" {

int QAccessibleInterface_isValid(void* self) {
	auto&& __r = ((QAccessibleInterface*)self)->isValid();
	return __r;
}
void* QAccessibleInterface_object(void* self) {
	auto&& __r = ((QAccessibleInterface*)self)->object();
	return (void*)(__r);
}
void* QAccessibleInterface_window(void* self) {
	auto&& __r = ((QAccessibleInterface*)self)->window();
	return (void*)(__r);
}
void* QAccessibleInterface_focusChild(void* self) {
	auto&& __r = ((QAccessibleInterface*)self)->focusChild();
	return (void*)(__r);
}
void* QAccessibleInterface_childAt(void* self, int x, int y) {
	auto&& __r = ((QAccessibleInterface*)self)->childAt(x, y);
	return (void*)(__r);
}
void* QAccessibleInterface_parent(void* self) {
	auto&& __r = ((QAccessibleInterface*)self)->parent();
	return (void*)(__r);
}
void* QAccessibleInterface_child(void* self, int index) {
	auto&& __r = ((QAccessibleInterface*)self)->child(index);
	return (void*)(__r);
}
int QAccessibleInterface_childCount(void* self) {
	auto&& __r = ((QAccessibleInterface*)self)->childCount();
	return __r;
}
int QAccessibleInterface_indexOfChild(void* self, void* arg1) {
	auto&& __r = ((QAccessibleInterface*)self)->indexOfChild((QAccessibleInterface*)arg1);
	return __r;
}
char* QAccessibleInterface_text(void* self, int t, long long* __out_len) {
	auto&& __r = ((QAccessibleInterface*)self)->text((QAccessible::Text)(t));
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QAccessibleInterface_setText(void* self, int t, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	((QAccessibleInterface*)self)->setText((QAccessible::Text)(t), text_qstr);
}
void* QAccessibleInterface_rect(void* self) {
	auto&& __r = ((QAccessibleInterface*)self)->rect();
	return new QRect(__r);
}
int QAccessibleInterface_role(void* self) {
	auto&& __r = ((QAccessibleInterface*)self)->role();
	return (int)(__r);
}
void* QAccessibleInterface_state(void* self) {
	auto&& __r = ((QAccessibleInterface*)self)->state();
	return new QAccessible::State(__r);
}
void* QAccessibleInterface_foregroundColor(void* self) {
	auto&& __r = ((QAccessibleInterface*)self)->foregroundColor();
	return new QColor(__r);
}
void* QAccessibleInterface_backgroundColor(void* self) {
	auto&& __r = ((QAccessibleInterface*)self)->backgroundColor();
	return new QColor(__r);
}
void* QAccessibleInterface_textInterface(void* self) {
	auto&& __r = ((QAccessibleInterface*)self)->textInterface();
	return (void*)(__r);
}
void* QAccessibleInterface_editableTextInterface(void* self) {
	auto&& __r = ((QAccessibleInterface*)self)->editableTextInterface();
	return (void*)(__r);
}
void* QAccessibleInterface_valueInterface(void* self) {
	auto&& __r = ((QAccessibleInterface*)self)->valueInterface();
	return (void*)(__r);
}
void* QAccessibleInterface_actionInterface(void* self) {
	auto&& __r = ((QAccessibleInterface*)self)->actionInterface();
	return (void*)(__r);
}
void* QAccessibleInterface_imageInterface(void* self) {
	auto&& __r = ((QAccessibleInterface*)self)->imageInterface();
	return (void*)(__r);
}
void* QAccessibleInterface_tableInterface(void* self) {
	auto&& __r = ((QAccessibleInterface*)self)->tableInterface();
	return (void*)(__r);
}
void* QAccessibleInterface_tableCellInterface(void* self) {
	auto&& __r = ((QAccessibleInterface*)self)->tableCellInterface();
	return (void*)(__r);
}
void* QAccessibleInterface_hyperlinkInterface(void* self) {
	auto&& __r = ((QAccessibleInterface*)self)->hyperlinkInterface();
	return (void*)(__r);
}
void* QAccessibleInterface_selectionInterface(void* self) {
	auto&& __r = ((QAccessibleInterface*)self)->selectionInterface();
	return (void*)(__r);
}
void* QAccessibleInterface_attributesInterface(void* self) {
	auto&& __r = ((QAccessibleInterface*)self)->attributesInterface();
	return (void*)(__r);
}
void* QAccessibleInterface_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QAccessibleInterface*)self)->operator=(*(QAccessibleInterface*)arg1);
	return (void*)&(__r);
}

}

#include <QtGui/qaccessibleobject.h>
#include <QtGui/qwindow.h>
#include <QtGui/qaccessible.h>
#include <QtCore/qstring.h>
#include <QtGui/qaccessible_base.h>
#include <QtCore/qobject.h>
#include <QtCore/qrect.h>
#include <QtGui/qcolor.h>
#include "gen_qaccessibleapplication.h"
#include "gen_purego_cb.h"

class QAccessibleApplication__Go : public QAccessibleApplication {
public:
	long long __goId;
	unsigned char __goOn[16] = {};
	template<typename... A> QAccessibleApplication__Go(long long id, A&&... a) : QAccessibleApplication(std::forward<A>(a)...), __goId(id) {}
	QWindow* window() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 8210911950352161981LL, &__h); if (__h) return (QWindow*)__r; }
		return QAccessibleApplication::window();
	}
	int childCount() const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, 1383730954482307804LL, &__h); if (__h) return __r; }
		return QAccessibleApplication::childCount();
	}
	int indexOfChild(const QAccessibleInterface* arg1) const override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, -9019360399188768292LL, (void*)arg1, &__h); if (__h) return __r; }
		return QAccessibleApplication::indexOfChild(arg1);
	}
	QAccessibleInterface* focusChild() const override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, -6961835385172491989LL, &__h); if (__h) return (QAccessibleInterface*)__r; }
		return QAccessibleApplication::focusChild();
	}
	QAccessibleInterface* parent() const override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 7371339085820260189LL, &__h); if (__h) return (QAccessibleInterface*)__r; }
		return QAccessibleApplication::parent();
	}
	QAccessibleInterface* child(int index) const override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7071834996910082205LL, index, &__h); if (__h) return (QAccessibleInterface*)__r; }
		return QAccessibleApplication::child(index);
	}
	QString text(QAccessible::Text t) const override {
		int __h = 0;
		if (__goOn[6] && cppgo_virtp_charp_int) {
			long long __len = 0;
			char* __s = ((char*(*)(long long, long long, int, long long*, int*))cppgo_virtp_charp_int)(__goId, -430950356557152130LL, (int)(t), &__len, &__h);
			if (__h) { QString __r = QString::fromUtf8(__s, __len); free(__s); return __r; }
		}
		return QAccessibleApplication::text(t);
	}
	QAccessible::Role role() const override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -6676326825648966727LL, &__h); if (__h) return (QAccessible::Role)(__r); }
		return QAccessibleApplication::role();
	}
	QAccessible::State state() const override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, -1268983008813010666LL, &__h); if (__h) return *(QAccessible::State*)__r; }
		return QAccessibleApplication::state();
	}
	bool isValid() const override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, 7231157875507698241LL, &__h); if (__h) return __r; }
		return QAccessibleApplication::isValid();
	}
	QObject* object() const override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, -8214935597255553030LL, &__h); if (__h) return (QObject*)__r; }
		return QAccessibleApplication::object();
	}
	QRect rect() const override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 5542927423540195655LL, &__h); if (__h) return *(QRect*)__r; }
		return QAccessibleApplication::rect();
	}
	void setText(QAccessible::Text t, const QString& text) override {
		int __h = 0;
		QByteArray text_ba = text.toUtf8();
		if (__goOn[12] && cppgo_virt_void_int_constcharp_longlong) { ((void(*)(long long, long long, int, const char*, long long, int*))cppgo_virt_void_int_constcharp_longlong)(__goId, -5557558535206160508LL, (int)(t), (char*)text_ba.constData(), text_ba.size(), &__h); if (__h) return; }
		QAccessibleApplication::setText(t, text);
	}
	QAccessibleInterface* childAt(int x, int y) const override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_voidp_int_int) { void* __r = ((void*(*)(long long, long long, int, int, int*))cppgo_virt_voidp_int_int)(__goId, 3663933090421434080LL, x, y, &__h); if (__h) return (QAccessibleInterface*)__r; }
		return QAccessibleApplication::childAt(x, y);
	}
	QColor foregroundColor() const override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, -3317306646323195305LL, &__h); if (__h) return *(QColor*)__r; }
		return QAccessibleApplication::foregroundColor();
	}
	QColor backgroundColor() const override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 2962551866395559440LL, &__h); if (__h) return *(QColor*)__r; }
		return QAccessibleApplication::backgroundColor();
	}
};

extern "C" {

void* QAccessibleApplication_base(void* self) { return (void*)static_cast<QAccessibleObject*>((QAccessibleApplication*)self); }
void QAccessibleApplication_chain(void* self, void** out) {
	QAccessibleApplication* __p = (QAccessibleApplication*)self;
	QAccessibleObject* __b0 = static_cast<QAccessibleObject*>(__p); out[0] = (void*)__b0;
	QAccessibleInterface* __b1 = static_cast<QAccessibleInterface*>(__b0); out[1] = (void*)__b1;
}
void QAccessibleApplication__set_override(void* self, int idx) { ((QAccessibleApplication__Go*)self)->__goOn[idx] = 1; }
void* QAccessibleApplication_new(long long __goid) {
	return new QAccessibleApplication__Go(__goid);
}
void* QAccessibleApplication_window_default(void* self) {
	auto&& __r = ((QAccessibleApplication*)self)->QAccessibleApplication::window();
	return (void*)(__r);
}
int QAccessibleApplication_childCount_default(void* self) {
	auto&& __r = ((QAccessibleApplication*)self)->QAccessibleApplication::childCount();
	return __r;
}
int QAccessibleApplication_indexOfChild_default(void* self, void* arg1) {
	auto&& __r = ((QAccessibleApplication*)self)->QAccessibleApplication::indexOfChild((QAccessibleInterface*)arg1);
	return __r;
}
void* QAccessibleApplication_focusChild_default(void* self) {
	auto&& __r = ((QAccessibleApplication*)self)->QAccessibleApplication::focusChild();
	return (void*)(__r);
}
void* QAccessibleApplication_parent_default(void* self) {
	auto&& __r = ((QAccessibleApplication*)self)->QAccessibleApplication::parent();
	return (void*)(__r);
}
void* QAccessibleApplication_child_default(void* self, int index) {
	auto&& __r = ((QAccessibleApplication*)self)->QAccessibleApplication::child(index);
	return (void*)(__r);
}
char* QAccessibleApplication_text_default(void* self, int t, long long* __out_len) {
	auto&& __r = ((QAccessibleApplication*)self)->QAccessibleApplication::text((QAccessible::Text)(t));
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QAccessibleApplication_role_default(void* self) {
	auto&& __r = ((QAccessibleApplication*)self)->QAccessibleApplication::role();
	return (int)(__r);
}
void* QAccessibleApplication_state_default(void* self) {
	auto&& __r = ((QAccessibleApplication*)self)->QAccessibleApplication::state();
	return new QAccessible::State(__r);
}
int QAccessibleApplication_isValid_default(void* self) {
	auto&& __r = ((QAccessibleApplication*)self)->QAccessibleApplication::isValid();
	return __r;
}
void* QAccessibleApplication_object_default(void* self) {
	auto&& __r = ((QAccessibleApplication*)self)->QAccessibleApplication::object();
	return (void*)(__r);
}
void* QAccessibleApplication_rect_default(void* self) {
	auto&& __r = ((QAccessibleApplication*)self)->QAccessibleApplication::rect();
	return new QRect(__r);
}
void QAccessibleApplication_setText_default(void* self, int t, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	((QAccessibleApplication*)self)->QAccessibleApplication::setText((QAccessible::Text)(t), text_qstr);
}
void* QAccessibleApplication_childAt_default(void* self, int x, int y) {
	auto&& __r = ((QAccessibleApplication*)self)->QAccessibleApplication::childAt(x, y);
	return (void*)(__r);
}
void* QAccessibleApplication_foregroundColor_default(void* self) {
	auto&& __r = ((QAccessibleApplication*)self)->QAccessibleApplication::foregroundColor();
	return new QColor(__r);
}
void* QAccessibleApplication_backgroundColor_default(void* self) {
	auto&& __r = ((QAccessibleApplication*)self)->QAccessibleApplication::backgroundColor();
	return new QColor(__r);
}
void* QAccessibleApplication__Go_new(long long id) { return new QAccessibleApplication__Go(id); }
void QAccessibleApplication_delete(void* self) { delete (QAccessibleApplication*)self; }

}

#include <QtGui/qiconengine.h>
#include <QtCore/qsize.h>
#include <QtGui/qicon.h>
#include <QtGui/qpixmap.h>
#include "gen_qiconengine__scaledpixmapargument.h"

extern "C" {

void* QIconEngine__ScaledPixmapArgument_new(void* arg1) {
	return new QIconEngine::ScaledPixmapArgument(*(QIconEngine::ScaledPixmapArgument*)arg1);
}
void* QIconEngine__ScaledPixmapArgument_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QIconEngine::ScaledPixmapArgument*)self)->operator=(*(QIconEngine::ScaledPixmapArgument*)arg1);
	return (void*)&(__r);
}
void* QIconEngine__ScaledPixmapArgument_size_get(void* self) {
	auto __r = ((QIconEngine::ScaledPixmapArgument*)self)->size;
	return new QSize(__r);
}
void QIconEngine__ScaledPixmapArgument_size_set(void* self, void* size) {
	((QIconEngine::ScaledPixmapArgument*)self)->size = *(QSize*)size;
}
int QIconEngine__ScaledPixmapArgument_mode_get(void* self) {
	auto __r = ((QIconEngine::ScaledPixmapArgument*)self)->mode;
	return (int)(__r);
}
void QIconEngine__ScaledPixmapArgument_mode_set(void* self, int mode) {
	((QIconEngine::ScaledPixmapArgument*)self)->mode = (QIcon::Mode)(mode);
}
int QIconEngine__ScaledPixmapArgument_state_get(void* self) {
	auto __r = ((QIconEngine::ScaledPixmapArgument*)self)->state;
	return (int)(__r);
}
void QIconEngine__ScaledPixmapArgument_state_set(void* self, int state) {
	((QIconEngine::ScaledPixmapArgument*)self)->state = (QIcon::State)(state);
}
double QIconEngine__ScaledPixmapArgument_scale_get(void* self) {
	auto __r = ((QIconEngine::ScaledPixmapArgument*)self)->scale;
	return __r;
}
void QIconEngine__ScaledPixmapArgument_scale_set(void* self, double scale) {
	((QIconEngine::ScaledPixmapArgument*)self)->scale = scale;
}
void* QIconEngine__ScaledPixmapArgument_pixmap_get(void* self) {
	auto __r = ((QIconEngine::ScaledPixmapArgument*)self)->pixmap;
	return new QPixmap(__r);
}
void QIconEngine__ScaledPixmapArgument_pixmap_set(void* self, void* pixmap) {
	((QIconEngine::ScaledPixmapArgument*)self)->pixmap = *(QPixmap*)pixmap;
}
void QIconEngine__ScaledPixmapArgument_delete(void* self) { delete (QIconEngine::ScaledPixmapArgument*)self; }

}

#include <QtGui/qpainterpath.h>
#include <QtCore/qpoint.h>
#include <QtCore/qrect.h>
#include <QtGui/qpolygon.h>
#include <QtGui/qfont.h>
#include <QtCore/qstring.h>
#include <QtGui/qregion.h>
#include <QtCore/qnamespace.h>
#include <QtGui/qtransform.h>
#include "gen_qpainterpath.h"

extern "C" {

void* QPainterPath_new() {
	return new QPainterPath();
}
void* QPainterPath_new2(void* startPoint) {
	return new QPainterPath(*(QPointF*)startPoint);
}
void* QPainterPath_new3(void* other) {
	return new QPainterPath(*(QPainterPath*)other);
}
void* QPainterPath_OperatorAssign(void* self, void* other) {
	auto&& __r = ((QPainterPath*)self)->operator=(*(QPainterPath*)other);
	return (void*)&(__r);
}
void QPainterPath_swap(void* self, void* other) {
	((QPainterPath*)self)->swap(*(QPainterPath*)other);
}
void QPainterPath_clear(void* self) {
	((QPainterPath*)self)->clear();
}
void QPainterPath_reserve(void* self, int size) {
	((QPainterPath*)self)->reserve(size);
}
int QPainterPath_capacity(void* self) {
	auto&& __r = ((QPainterPath*)self)->capacity();
	return __r;
}
void QPainterPath_closeSubpath(void* self) {
	((QPainterPath*)self)->closeSubpath();
}
void QPainterPath_moveTo(void* self, void* p) {
	((QPainterPath*)self)->moveTo(*(QPointF*)p);
}
void QPainterPath_moveTo2(void* self, double x, double y) {
	((QPainterPath*)self)->moveTo(x, y);
}
void QPainterPath_lineTo(void* self, void* p) {
	((QPainterPath*)self)->lineTo(*(QPointF*)p);
}
void QPainterPath_lineTo2(void* self, double x, double y) {
	((QPainterPath*)self)->lineTo(x, y);
}
void QPainterPath_arcMoveTo(void* self, void* rect, double angle) {
	((QPainterPath*)self)->arcMoveTo(*(QRectF*)rect, angle);
}
void QPainterPath_arcMoveTo2(void* self, double x, double y, double w, double h, double angle) {
	((QPainterPath*)self)->arcMoveTo(x, y, w, h, angle);
}
void QPainterPath_arcTo(void* self, void* rect, double startAngle, double arcLength) {
	((QPainterPath*)self)->arcTo(*(QRectF*)rect, startAngle, arcLength);
}
void QPainterPath_arcTo2(void* self, double x, double y, double w, double h, double startAngle, double arcLength) {
	((QPainterPath*)self)->arcTo(x, y, w, h, startAngle, arcLength);
}
void QPainterPath_cubicTo(void* self, void* ctrlPt1, void* ctrlPt2, void* endPt) {
	((QPainterPath*)self)->cubicTo(*(QPointF*)ctrlPt1, *(QPointF*)ctrlPt2, *(QPointF*)endPt);
}
void QPainterPath_cubicTo2(void* self, double ctrlPt1x, double ctrlPt1y, double ctrlPt2x, double ctrlPt2y, double endPtx, double endPty) {
	((QPainterPath*)self)->cubicTo(ctrlPt1x, ctrlPt1y, ctrlPt2x, ctrlPt2y, endPtx, endPty);
}
void QPainterPath_quadTo(void* self, void* ctrlPt, void* endPt) {
	((QPainterPath*)self)->quadTo(*(QPointF*)ctrlPt, *(QPointF*)endPt);
}
void QPainterPath_quadTo2(void* self, double ctrlPtx, double ctrlPty, double endPtx, double endPty) {
	((QPainterPath*)self)->quadTo(ctrlPtx, ctrlPty, endPtx, endPty);
}
void* QPainterPath_currentPosition(void* self) {
	auto&& __r = ((QPainterPath*)self)->currentPosition();
	return new QPointF(__r);
}
void QPainterPath_addRect(void* self, void* rect) {
	((QPainterPath*)self)->addRect(*(QRectF*)rect);
}
void QPainterPath_addRect2(void* self, double x, double y, double w, double h) {
	((QPainterPath*)self)->addRect(x, y, w, h);
}
void QPainterPath_addEllipse(void* self, void* rect) {
	((QPainterPath*)self)->addEllipse(*(QRectF*)rect);
}
void QPainterPath_addEllipse2(void* self, double x, double y, double w, double h) {
	((QPainterPath*)self)->addEllipse(x, y, w, h);
}
void QPainterPath_addEllipse3(void* self, void* center, double rx, double ry) {
	((QPainterPath*)self)->addEllipse(*(QPointF*)center, rx, ry);
}
void QPainterPath_addPolygon(void* self, void* polygon) {
	((QPainterPath*)self)->addPolygon(*(QPolygonF*)polygon);
}
void QPainterPath_addText(void* self, void* point, void* f, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	((QPainterPath*)self)->addText(*(QPointF*)point, *(QFont*)f, text_qstr);
}
void QPainterPath_addText2(void* self, double x, double y, void* f, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	((QPainterPath*)self)->addText(x, y, *(QFont*)f, text_qstr);
}
void QPainterPath_addPath(void* self, void* path) {
	((QPainterPath*)self)->addPath(*(QPainterPath*)path);
}
void QPainterPath_addRegion(void* self, void* region) {
	((QPainterPath*)self)->addRegion(*(QRegion*)region);
}
void QPainterPath_addRoundedRect(void* self, void* rect, double xRadius, double yRadius) {
	((QPainterPath*)self)->addRoundedRect(*(QRectF*)rect, xRadius, yRadius);
}
void QPainterPath_addRoundedRect2(void* self, double x, double y, double w, double h, double xRadius, double yRadius) {
	((QPainterPath*)self)->addRoundedRect(x, y, w, h, xRadius, yRadius);
}
void QPainterPath_connectPath(void* self, void* path) {
	((QPainterPath*)self)->connectPath(*(QPainterPath*)path);
}
int QPainterPath_contains(void* self, void* pt) {
	auto&& __r = ((QPainterPath*)self)->contains(*(QPointF*)pt);
	return __r;
}
int QPainterPath_contains2(void* self, void* rect) {
	auto&& __r = ((QPainterPath*)self)->contains(*(QRectF*)rect);
	return __r;
}
int QPainterPath_intersects(void* self, void* rect) {
	auto&& __r = ((QPainterPath*)self)->intersects(*(QRectF*)rect);
	return __r;
}
void QPainterPath_translate(void* self, double dx, double dy) {
	((QPainterPath*)self)->translate(dx, dy);
}
void QPainterPath_translate2(void* self, void* offset) {
	((QPainterPath*)self)->translate(*(QPointF*)offset);
}
void* QPainterPath_translated(void* self, double dx, double dy) {
	auto&& __r = ((QPainterPath*)self)->translated(dx, dy);
	return new QPainterPath(__r);
}
void* QPainterPath_translated2(void* self, void* offset) {
	auto&& __r = ((QPainterPath*)self)->translated(*(QPointF*)offset);
	return new QPainterPath(__r);
}
void* QPainterPath_boundingRect(void* self) {
	auto&& __r = ((QPainterPath*)self)->boundingRect();
	return new QRectF(__r);
}
void* QPainterPath_controlPointRect(void* self) {
	auto&& __r = ((QPainterPath*)self)->controlPointRect();
	return new QRectF(__r);
}
int QPainterPath_fillRule(void* self) {
	auto&& __r = ((QPainterPath*)self)->fillRule();
	return (int)(__r);
}
void QPainterPath_setFillRule(void* self, int fillRule) {
	((QPainterPath*)self)->setFillRule((Qt::FillRule)(fillRule));
}
int QPainterPath_isEmpty(void* self) {
	auto&& __r = ((QPainterPath*)self)->isEmpty();
	return __r;
}
void* QPainterPath_toReversed(void* self) {
	auto&& __r = ((QPainterPath*)self)->toReversed();
	return new QPainterPath(__r);
}
void** QPainterPath_toSubpathPolygons(void* self, long long* __out_len) {
	auto&& __r = ((QPainterPath*)self)->toSubpathPolygons();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QPolygonF(__e);
	return __arr;
}
void** QPainterPath_toFillPolygons(void* self, long long* __out_len) {
	auto&& __r = ((QPainterPath*)self)->toFillPolygons();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QPolygonF(__e);
	return __arr;
}
void* QPainterPath_toFillPolygon(void* self) {
	auto&& __r = ((QPainterPath*)self)->toFillPolygon();
	return new QPolygonF(__r);
}
int QPainterPath_elementCount(void* self) {
	auto&& __r = ((QPainterPath*)self)->elementCount();
	return __r;
}
void* QPainterPath_elementAt(void* self, int i) {
	auto&& __r = ((QPainterPath*)self)->elementAt(i);
	return new QPainterPath::Element(__r);
}
void QPainterPath_setElementPositionAt(void* self, int i, double x, double y) {
	((QPainterPath*)self)->setElementPositionAt(i, x, y);
}
int QPainterPath_isCachingEnabled(void* self) {
	auto&& __r = ((QPainterPath*)self)->isCachingEnabled();
	return __r;
}
void QPainterPath_setCachingEnabled(void* self, int enabled) {
	((QPainterPath*)self)->setCachingEnabled(enabled);
}
double QPainterPath_length(void* self) {
	auto&& __r = ((QPainterPath*)self)->length();
	return __r;
}
double QPainterPath_percentAtLength(void* self, double len_) {
	auto&& __r = ((QPainterPath*)self)->percentAtLength(len_);
	return __r;
}
void* QPainterPath_pointAtPercent(void* self, double t) {
	auto&& __r = ((QPainterPath*)self)->pointAtPercent(t);
	return new QPointF(__r);
}
double QPainterPath_angleAtPercent(void* self, double t) {
	auto&& __r = ((QPainterPath*)self)->angleAtPercent(t);
	return __r;
}
double QPainterPath_slopeAtPercent(void* self, double t) {
	auto&& __r = ((QPainterPath*)self)->slopeAtPercent(t);
	return __r;
}
void* QPainterPath_trimmed(void* self, double fromFraction, double toFraction) {
	auto&& __r = ((QPainterPath*)self)->trimmed(fromFraction, toFraction);
	return new QPainterPath(__r);
}
int QPainterPath_intersects2(void* self, void* p) {
	auto&& __r = ((QPainterPath*)self)->intersects(*(QPainterPath*)p);
	return __r;
}
int QPainterPath_contains3(void* self, void* p) {
	auto&& __r = ((QPainterPath*)self)->contains(*(QPainterPath*)p);
	return __r;
}
void* QPainterPath_united(void* self, void* r) {
	auto&& __r = ((QPainterPath*)self)->united(*(QPainterPath*)r);
	return new QPainterPath(__r);
}
void* QPainterPath_intersected(void* self, void* r) {
	auto&& __r = ((QPainterPath*)self)->intersected(*(QPainterPath*)r);
	return new QPainterPath(__r);
}
void* QPainterPath_subtracted(void* self, void* r) {
	auto&& __r = ((QPainterPath*)self)->subtracted(*(QPainterPath*)r);
	return new QPainterPath(__r);
}
void* QPainterPath_simplified(void* self) {
	auto&& __r = ((QPainterPath*)self)->simplified();
	return new QPainterPath(__r);
}
int QPainterPath_OperatorEqual(void* self, void* other) {
	auto&& __r = ((QPainterPath*)self)->operator==(*(QPainterPath*)other);
	return __r;
}
int QPainterPath_OperatorNotEqual(void* self, void* other) {
	auto&& __r = ((QPainterPath*)self)->operator!=(*(QPainterPath*)other);
	return __r;
}
void* QPainterPath_OperatorBitwiseAnd(void* self, void* other) {
	auto&& __r = ((QPainterPath*)self)->operator&(*(QPainterPath*)other);
	return new QPainterPath(__r);
}
void* QPainterPath_OperatorBitwiseOr(void* self, void* other) {
	auto&& __r = ((QPainterPath*)self)->operator|(*(QPainterPath*)other);
	return new QPainterPath(__r);
}
void* QPainterPath_OperatorPlus(void* self, void* other) {
	auto&& __r = ((QPainterPath*)self)->operator+(*(QPainterPath*)other);
	return new QPainterPath(__r);
}
void* QPainterPath_OperatorMinus(void* self, void* other) {
	auto&& __r = ((QPainterPath*)self)->operator-(*(QPainterPath*)other);
	return new QPainterPath(__r);
}
void* QPainterPath_OperatorBitwiseAndAssign(void* self, void* other) {
	auto&& __r = ((QPainterPath*)self)->operator&=(*(QPainterPath*)other);
	return (void*)&(__r);
}
void* QPainterPath_OperatorBitwiseOrAssign(void* self, void* other) {
	auto&& __r = ((QPainterPath*)self)->operator|=(*(QPainterPath*)other);
	return (void*)&(__r);
}
void* QPainterPath_OperatorPlusAssign(void* self, void* other) {
	auto&& __r = ((QPainterPath*)self)->operator+=(*(QPainterPath*)other);
	return (void*)&(__r);
}
void* QPainterPath_OperatorMinusAssign(void* self, void* other) {
	auto&& __r = ((QPainterPath*)self)->operator-=(*(QPainterPath*)other);
	return (void*)&(__r);
}
void QPainterPath_addRoundedRect3(void* self, void* rect, double xRadius, double yRadius, int mode) {
	((QPainterPath*)self)->addRoundedRect(*(QRectF*)rect, xRadius, yRadius, (Qt::SizeMode)(mode));
}
void QPainterPath_addRoundedRect4(void* self, double x, double y, double w, double h, double xRadius, double yRadius, int mode) {
	((QPainterPath*)self)->addRoundedRect(x, y, w, h, xRadius, yRadius, (Qt::SizeMode)(mode));
}
void** QPainterPath_toSubpathPolygons2(void* self, void* matrix, long long* __out_len) {
	auto&& __r = ((QPainterPath*)self)->toSubpathPolygons(*(QTransform*)matrix);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QPolygonF(__e);
	return __arr;
}
void** QPainterPath_toFillPolygons2(void* self, void* matrix, long long* __out_len) {
	auto&& __r = ((QPainterPath*)self)->toFillPolygons(*(QTransform*)matrix);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QPolygonF(__e);
	return __arr;
}
void* QPainterPath_toFillPolygon2(void* self, void* matrix) {
	auto&& __r = ((QPainterPath*)self)->toFillPolygon(*(QTransform*)matrix);
	return new QPolygonF(__r);
}
void* QPainterPath_trimmed2(void* self, double fromFraction, double toFraction, double offset) {
	auto&& __r = ((QPainterPath*)self)->trimmed(fromFraction, toFraction, offset);
	return new QPainterPath(__r);
}
void* QPainterPath_OperatorMultiply(void* self, void* m) {
	auto&& __r = (*(QPainterPath*)self) * (*(QTransform*)m);
	return new QPainterPath(__r);
}
void QPainterPath_delete(void* self) { delete (QPainterPath*)self; }

}

#include <QtGui/qtextobject.h>
#include <QtGui/qtextlayout.h>
#include <QtGui/qtextformat.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qstring.h>
#include <QtGui/qtextdocument.h>
#include <QtGui/qtextlist.h>
#include "gen_qtextblock.h"

extern "C" {

void* QTextBlock_new() {
	return new QTextBlock();
}
void* QTextBlock_new2(void* o) {
	return new QTextBlock(*(QTextBlock*)o);
}
void* QTextBlock_OperatorAssign(void* self, void* o) {
	auto&& __r = ((QTextBlock*)self)->operator=(*(QTextBlock*)o);
	return (void*)&(__r);
}
int QTextBlock_isValid(void* self) {
	auto&& __r = ((QTextBlock*)self)->isValid();
	return __r;
}
int QTextBlock_OperatorEqual(void* self, void* o) {
	auto&& __r = ((QTextBlock*)self)->operator==(*(QTextBlock*)o);
	return __r;
}
int QTextBlock_OperatorNotEqual(void* self, void* o) {
	auto&& __r = ((QTextBlock*)self)->operator!=(*(QTextBlock*)o);
	return __r;
}
int QTextBlock_OperatorLesser(void* self, void* o) {
	auto&& __r = ((QTextBlock*)self)->operator<(*(QTextBlock*)o);
	return __r;
}
int QTextBlock_position(void* self) {
	auto&& __r = ((QTextBlock*)self)->position();
	return __r;
}
int QTextBlock_length(void* self) {
	auto&& __r = ((QTextBlock*)self)->length();
	return __r;
}
int QTextBlock_contains(void* self, int position) {
	auto&& __r = ((QTextBlock*)self)->contains(position);
	return __r;
}
void* QTextBlock_layout(void* self) {
	auto&& __r = ((QTextBlock*)self)->layout();
	return (void*)(__r);
}
void QTextBlock_clearLayout(void* self) {
	((QTextBlock*)self)->clearLayout();
}
void* QTextBlock_blockFormat(void* self) {
	auto&& __r = ((QTextBlock*)self)->blockFormat();
	return new QTextBlockFormat(__r);
}
int QTextBlock_blockFormatIndex(void* self) {
	auto&& __r = ((QTextBlock*)self)->blockFormatIndex();
	return __r;
}
void* QTextBlock_charFormat(void* self) {
	auto&& __r = ((QTextBlock*)self)->charFormat();
	return new QTextCharFormat(__r);
}
int QTextBlock_charFormatIndex(void* self) {
	auto&& __r = ((QTextBlock*)self)->charFormatIndex();
	return __r;
}
int QTextBlock_textDirection(void* self) {
	auto&& __r = ((QTextBlock*)self)->textDirection();
	return (int)(__r);
}
char* QTextBlock_text(void* self, long long* __out_len) {
	auto&& __r = ((QTextBlock*)self)->text();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void** QTextBlock_textFormats(void* self, long long* __out_len) {
	auto&& __r = ((QTextBlock*)self)->textFormats();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QTextLayout::FormatRange(__e);
	return __arr;
}
void* QTextBlock_document(void* self) {
	auto&& __r = ((QTextBlock*)self)->document();
	return (void*)(__r);
}
void* QTextBlock_textList(void* self) {
	auto&& __r = ((QTextBlock*)self)->textList();
	return (void*)(__r);
}
void* QTextBlock_userData(void* self) {
	auto&& __r = ((QTextBlock*)self)->userData();
	return (void*)(__r);
}
void QTextBlock_setUserData(void* self, void* data) {
	((QTextBlock*)self)->setUserData((QTextBlockUserData*)data);
}
int QTextBlock_userState(void* self) {
	auto&& __r = ((QTextBlock*)self)->userState();
	return __r;
}
void QTextBlock_setUserState(void* self, int state) {
	((QTextBlock*)self)->setUserState(state);
}
int QTextBlock_revision(void* self) {
	auto&& __r = ((QTextBlock*)self)->revision();
	return __r;
}
void QTextBlock_setRevision(void* self, int rev) {
	((QTextBlock*)self)->setRevision(rev);
}
int QTextBlock_isVisible(void* self) {
	auto&& __r = ((QTextBlock*)self)->isVisible();
	return __r;
}
void QTextBlock_setVisible(void* self, int visible) {
	((QTextBlock*)self)->setVisible(visible);
}
int QTextBlock_blockNumber(void* self) {
	auto&& __r = ((QTextBlock*)self)->blockNumber();
	return __r;
}
int QTextBlock_firstLineNumber(void* self) {
	auto&& __r = ((QTextBlock*)self)->firstLineNumber();
	return __r;
}
void QTextBlock_setLineCount(void* self, int count) {
	((QTextBlock*)self)->setLineCount(count);
}
int QTextBlock_lineCount(void* self) {
	auto&& __r = ((QTextBlock*)self)->lineCount();
	return __r;
}
void* QTextBlock_next(void* self) {
	auto&& __r = ((QTextBlock*)self)->next();
	return new QTextBlock(__r);
}
void* QTextBlock_previous(void* self) {
	auto&& __r = ((QTextBlock*)self)->previous();
	return new QTextBlock(__r);
}
int QTextBlock_fragmentIndex(void* self) {
	auto&& __r = ((QTextBlock*)self)->fragmentIndex();
	return __r;
}
void QTextBlock_delete(void* self) { delete (QTextBlock*)self; }

}

#include <QtWidgets/qtabbar.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtGui/qicon.h>
#include <QtGui/qcolor.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qvariant.h>
#include <QtCore/qrect.h>
#include <QtCore/qpoint.h>
#include <QtCore/qsize.h>
#include <QtCore/qcoreevent.h>
#include <QtGui/qevent.h>
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
#include "gen_qtabbar.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QTabBar__Go : public QTabBar {
public:
	long long __goId;
	unsigned char __goOn[54] = {};
	template<typename... A> QTabBar__Go(long long id, A&&... a) : QTabBar(std::forward<A>(a)...), __goId(id) {}
	using QTabBar::updateMicroFocus;
	using QTabBar::create;
	using QTabBar::destroy;
	using QTabBar::focusNextChild;
	using QTabBar::focusPreviousChild;
	using QTabBar::sender;
	using QTabBar::senderSignalIndex;
	using QTabBar::receivers;
	using QTabBar::isSignalConnected;
	using QTabBar::getDecodedMetricF;
	QSize cppgoBase_tabSizeHint(int __a0) { return QTabBar::tabSizeHint(__a0); }
	QSize cppgoBase_minimumTabSizeHint(int __a0) { return QTabBar::minimumTabSizeHint(__a0); }
	void cppgoBase_tabInserted(int __a0) { QTabBar::tabInserted(__a0); }
	void cppgoBase_tabRemoved(int __a0) { QTabBar::tabRemoved(__a0); }
	void cppgoBase_tabLayoutChange() { QTabBar::tabLayoutChange(); }
	bool cppgoBase_event(QEvent* __a0) { return QTabBar::event(__a0); }
	void cppgoBase_resizeEvent(QResizeEvent* __a0) { QTabBar::resizeEvent(__a0); }
	void cppgoBase_showEvent(QShowEvent* __a0) { QTabBar::showEvent(__a0); }
	void cppgoBase_hideEvent(QHideEvent* __a0) { QTabBar::hideEvent(__a0); }
	void cppgoBase_paintEvent(QPaintEvent* __a0) { QTabBar::paintEvent(__a0); }
	void cppgoBase_mousePressEvent(QMouseEvent* __a0) { QTabBar::mousePressEvent(__a0); }
	void cppgoBase_mouseMoveEvent(QMouseEvent* __a0) { QTabBar::mouseMoveEvent(__a0); }
	void cppgoBase_mouseReleaseEvent(QMouseEvent* __a0) { QTabBar::mouseReleaseEvent(__a0); }
	void cppgoBase_mouseDoubleClickEvent(QMouseEvent* __a0) { QTabBar::mouseDoubleClickEvent(__a0); }
	void cppgoBase_wheelEvent(QWheelEvent* __a0) { QTabBar::wheelEvent(__a0); }
	void cppgoBase_keyPressEvent(QKeyEvent* __a0) { QTabBar::keyPressEvent(__a0); }
	void cppgoBase_changeEvent(QEvent* __a0) { QTabBar::changeEvent(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QTabBar::timerEvent(__a0); }
	void cppgoBase_initStyleOption(QStyleOptionTab* __a0, int __a1) { QTabBar::initStyleOption(__a0, __a1); }
	void cppgoBase_keyReleaseEvent(QKeyEvent* __a0) { QTabBar::keyReleaseEvent(__a0); }
	void cppgoBase_focusInEvent(QFocusEvent* __a0) { QTabBar::focusInEvent(__a0); }
	void cppgoBase_focusOutEvent(QFocusEvent* __a0) { QTabBar::focusOutEvent(__a0); }
	void cppgoBase_enterEvent(QEnterEvent* __a0) { QTabBar::enterEvent(__a0); }
	void cppgoBase_leaveEvent(QEvent* __a0) { QTabBar::leaveEvent(__a0); }
	void cppgoBase_moveEvent(QMoveEvent* __a0) { QTabBar::moveEvent(__a0); }
	void cppgoBase_closeEvent(QCloseEvent* __a0) { QTabBar::closeEvent(__a0); }
	void cppgoBase_contextMenuEvent(QContextMenuEvent* __a0) { QTabBar::contextMenuEvent(__a0); }
	void cppgoBase_tabletEvent(QTabletEvent* __a0) { QTabBar::tabletEvent(__a0); }
	void cppgoBase_actionEvent(QActionEvent* __a0) { QTabBar::actionEvent(__a0); }
	void cppgoBase_dragEnterEvent(QDragEnterEvent* __a0) { QTabBar::dragEnterEvent(__a0); }
	void cppgoBase_dragMoveEvent(QDragMoveEvent* __a0) { QTabBar::dragMoveEvent(__a0); }
	void cppgoBase_dragLeaveEvent(QDragLeaveEvent* __a0) { QTabBar::dragLeaveEvent(__a0); }
	void cppgoBase_dropEvent(QDropEvent* __a0) { QTabBar::dropEvent(__a0); }
	int cppgoBase_metric(int __a0) { return QTabBar::metric((QPaintDevice::PaintDeviceMetric)__a0); }
	void cppgoBase_initPainter(QPainter* __a0) { QTabBar::initPainter(__a0); }
	QPaintDevice* cppgoBase_redirected(QPoint* __a0) { return QTabBar::redirected(__a0); }
	QPainter* cppgoBase_sharedPainter() { return QTabBar::sharedPainter(); }
	void cppgoBase_inputMethodEvent(QInputMethodEvent* __a0) { QTabBar::inputMethodEvent(__a0); }
	bool cppgoBase_focusNextPrevChild(bool __a0) { return QTabBar::focusNextPrevChild(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QTabBar::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QTabBar::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QTabBar::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QTabBar::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QTabBar::metaObject();
	}
	QSize sizeHint() const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 729690261457684101LL, &__h); if (__h) return *(QSize*)__r; }
		return QTabBar::sizeHint();
	}
	QSize minimumSizeHint() const override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 7049264599837532491LL, &__h); if (__h) return *(QSize*)__r; }
		return QTabBar::minimumSizeHint();
	}
	QSize tabSizeHint(int index) const override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, -5008018812442754120LL, index, &__h); if (__h) return *(QSize*)__r; }
		return QTabBar::tabSizeHint(index);
	}
	QSize minimumTabSizeHint(int index) const override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, -5553900436438809362LL, index, &__h); if (__h) return *(QSize*)__r; }
		return QTabBar::minimumTabSizeHint(index);
	}
	void tabInserted(int index) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, -2734286415018532480LL, index, &__h); if (__h) return; }
		QTabBar::tabInserted(index);
	}
	void tabRemoved(int index) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, -3307029722643392856LL, index, &__h); if (__h) return; }
		QTabBar::tabRemoved(index);
	}
	void tabLayoutChange() override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 9115710866053612342LL, &__h); if (__h) return; }
		QTabBar::tabLayoutChange();
	}
	bool event(QEvent* arg1) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)arg1, &__h); if (__h) return __r; }
		return QTabBar::event(arg1);
	}
	void resizeEvent(QResizeEvent* arg1) override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4302917285655014661LL, (void*)arg1, &__h); if (__h) return; }
		QTabBar::resizeEvent(arg1);
	}
	void showEvent(QShowEvent* arg1) override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6825369174751916LL, (void*)arg1, &__h); if (__h) return; }
		QTabBar::showEvent(arg1);
	}
	void hideEvent(QHideEvent* arg1) override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -199718707876719129LL, (void*)arg1, &__h); if (__h) return; }
		QTabBar::hideEvent(arg1);
	}
	void paintEvent(QPaintEvent* arg1) override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5290442163167458465LL, (void*)arg1, &__h); if (__h) return; }
		QTabBar::paintEvent(arg1);
	}
	void mousePressEvent(QMouseEvent* arg1) override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -332090898078543319LL, (void*)arg1, &__h); if (__h) return; }
		QTabBar::mousePressEvent(arg1);
	}
	void mouseMoveEvent(QMouseEvent* arg1) override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4255494004099806581LL, (void*)arg1, &__h); if (__h) return; }
		QTabBar::mouseMoveEvent(arg1);
	}
	void mouseReleaseEvent(QMouseEvent* arg1) override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4415920049297795213LL, (void*)arg1, &__h); if (__h) return; }
		QTabBar::mouseReleaseEvent(arg1);
	}
	void mouseDoubleClickEvent(QMouseEvent* arg1) override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4029073905697023887LL, (void*)arg1, &__h); if (__h) return; }
		QTabBar::mouseDoubleClickEvent(arg1);
	}
	void wheelEvent(QWheelEvent* event) override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1919988340934511316LL, (void*)event, &__h); if (__h) return; }
		QTabBar::wheelEvent(event);
	}
	void keyPressEvent(QKeyEvent* arg1) override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1023791977679793185LL, (void*)arg1, &__h); if (__h) return; }
		QTabBar::keyPressEvent(arg1);
	}
	void changeEvent(QEvent* arg1) override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4959376063826165943LL, (void*)arg1, &__h); if (__h) return; }
		QTabBar::changeEvent(arg1);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QTabBar::timerEvent(event);
	}
	void initStyleOption(QStyleOptionTab* option, int tabIndex) const override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_void_voidp_int) { ((void(*)(long long, long long, void*, int, int*))cppgo_virt_void_voidp_int)(__goId, -488693769771308531LL, (void*)option, tabIndex, &__h); if (__h) return; }
		QTabBar::initStyleOption(option, tabIndex);
	}
	int devType() const override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8659679893868524044LL, &__h); if (__h) return __r; }
		return QTabBar::devType();
	}
	void setVisible(bool visible) override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 2110182719375245575LL, visible, &__h); if (__h) return; }
		QTabBar::setVisible(visible);
	}
	int heightForWidth(int arg1) const override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 4311534060495737339LL, arg1, &__h); if (__h) return __r; }
		return QTabBar::heightForWidth(arg1);
	}
	bool hasHeightForWidth() const override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8266762397995423915LL, &__h); if (__h) return __r; }
		return QTabBar::hasHeightForWidth();
	}
	QPaintEngine* paintEngine() const override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 6398007404756346381LL, &__h); if (__h) return (QPaintEngine*)__r; }
		return QTabBar::paintEngine();
	}
	void keyReleaseEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 772598475950050299LL, (void*)event, &__h); if (__h) return; }
		QTabBar::keyReleaseEvent(event);
	}
	void focusInEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 975043767033552552LL, (void*)event, &__h); if (__h) return; }
		QTabBar::focusInEvent(event);
	}
	void focusOutEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -385596010576048993LL, (void*)event, &__h); if (__h) return; }
		QTabBar::focusOutEvent(event);
	}
	void enterEvent(QEnterEvent* event) override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6316270346668237895LL, (void*)event, &__h); if (__h) return; }
		QTabBar::enterEvent(event);
	}
	void leaveEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -551070009623711640LL, (void*)event, &__h); if (__h) return; }
		QTabBar::leaveEvent(event);
	}
	void moveEvent(QMoveEvent* event) override {
		int __h = 0;
		if (__goOn[32] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -9161949150489525396LL, (void*)event, &__h); if (__h) return; }
		QTabBar::moveEvent(event);
	}
	void closeEvent(QCloseEvent* event) override {
		int __h = 0;
		if (__goOn[33] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1989937050019464783LL, (void*)event, &__h); if (__h) return; }
		QTabBar::closeEvent(event);
	}
	void contextMenuEvent(QContextMenuEvent* event) override {
		int __h = 0;
		if (__goOn[34] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6698043222227006015LL, (void*)event, &__h); if (__h) return; }
		QTabBar::contextMenuEvent(event);
	}
	void tabletEvent(QTabletEvent* event) override {
		int __h = 0;
		if (__goOn[35] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -212068866860701141LL, (void*)event, &__h); if (__h) return; }
		QTabBar::tabletEvent(event);
	}
	void actionEvent(QActionEvent* event) override {
		int __h = 0;
		if (__goOn[36] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1697444140088555581LL, (void*)event, &__h); if (__h) return; }
		QTabBar::actionEvent(event);
	}
	void dragEnterEvent(QDragEnterEvent* event) override {
		int __h = 0;
		if (__goOn[37] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 206122715919917589LL, (void*)event, &__h); if (__h) return; }
		QTabBar::dragEnterEvent(event);
	}
	void dragMoveEvent(QDragMoveEvent* event) override {
		int __h = 0;
		if (__goOn[38] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1781545137352745978LL, (void*)event, &__h); if (__h) return; }
		QTabBar::dragMoveEvent(event);
	}
	void dragLeaveEvent(QDragLeaveEvent* event) override {
		int __h = 0;
		if (__goOn[39] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -3139904158373725006LL, (void*)event, &__h); if (__h) return; }
		QTabBar::dragLeaveEvent(event);
	}
	void dropEvent(QDropEvent* event) override {
		int __h = 0;
		if (__goOn[40] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 9107058682754591716LL, (void*)event, &__h); if (__h) return; }
		QTabBar::dropEvent(event);
	}
	bool nativeEvent(const QByteArray& eventType, void* message, qintptr* result) override {
		int __h = 0;
		if (__goOn[41] && cppgo_virt_int_voidp_voidp_longlongp) { int __r = ((int(*)(long long, long long, void*, void*, long long*, int*))cppgo_virt_int_voidp_voidp_longlongp)(__goId, -2039166586246443706LL, (void*)&eventType, (void*)message, (long long*)result, &__h); if (__h) return __r; }
		return QTabBar::nativeEvent(eventType, message, result);
	}
	int metric(QPaintDevice::PaintDeviceMetric arg1) const override {
		int __h = 0;
		if (__goOn[42] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 427629649267311777LL, (int)(arg1), &__h); if (__h) return __r; }
		return QTabBar::metric(arg1);
	}
	void initPainter(QPainter* painter) const override {
		int __h = 0;
		if (__goOn[43] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6477500296980050210LL, (void*)painter, &__h); if (__h) return; }
		QTabBar::initPainter(painter);
	}
	QPaintDevice* redirected(QPoint* offset) const override {
		int __h = 0;
		if (__goOn[44] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 4127240633310655048LL, (void*)offset, &__h); if (__h) return (QPaintDevice*)__r; }
		return QTabBar::redirected(offset);
	}
	QPainter* sharedPainter() const override {
		int __h = 0;
		if (__goOn[45] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 5792771328741774991LL, &__h); if (__h) return (QPainter*)__r; }
		return QTabBar::sharedPainter();
	}
	void inputMethodEvent(QInputMethodEvent* arg1) override {
		int __h = 0;
		if (__goOn[46] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5438819661942843222LL, (void*)arg1, &__h); if (__h) return; }
		QTabBar::inputMethodEvent(arg1);
	}
	QVariant inputMethodQuery(Qt::InputMethodQuery arg1) const override {
		int __h = 0;
		if (__goOn[47] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7345853097340709638LL, (int)(arg1), &__h); if (__h) return *(QVariant*)__r; }
		return QTabBar::inputMethodQuery(arg1);
	}
	bool focusNextPrevChild(bool next) override {
		int __h = 0;
		if (__goOn[48] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 5898328517505784191LL, next, &__h); if (__h) return __r; }
		return QTabBar::focusNextPrevChild(next);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[49] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QTabBar::eventFilter(watched, event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[50] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QTabBar::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[51] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QTabBar::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[52] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QTabBar::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[53] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QTabBar::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QTabBar;
static std::mutex __connsMu_QTabBar;

extern "C" {

void* QTabBar_base(void* self) { return (void*)static_cast<QWidget*>((QTabBar*)self); }
void QTabBar_chain(void* self, void** out) {
	QTabBar* __p = (QTabBar*)self;
	QWidget* __b0 = static_cast<QWidget*>(__p); out[0] = (void*)__b0;
	QObject* __b1 = static_cast<QObject*>(__b0); out[1] = (void*)__b1;
}
void QTabBar__set_override(void* self, int idx) { ((QTabBar__Go*)self)->__goOn[idx] = 1; }
void* QTabBar_fromQObject(void* o) { return (void*)qobject_cast<QTabBar*>((QObject*)o); }
void* QTabBar_new(long long __goid, void* parent) {
	return new QTabBar__Go(__goid, (QWidget*)parent);
}
void* QTabBar_new2(long long __goid) {
	return new QTabBar__Go(__goid);
}
char* QTabBar_tr(const char* s, long long* __out_len) {
	auto&& __r = QTabBar::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QTabBar_shape(void* self) {
	auto&& __r = ((QTabBar*)self)->shape();
	return (int)(__r);
}
void QTabBar_setShape(void* self, int shape) {
	((QTabBar*)self)->setShape((QTabBar::Shape)(shape));
}
int QTabBar_addTab(void* self, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = ((QTabBar*)self)->addTab(text_qstr);
	return __r;
}
int QTabBar_addTab2(void* self, void* icon, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = ((QTabBar*)self)->addTab(*(QIcon*)icon, text_qstr);
	return __r;
}
int QTabBar_insertTab(void* self, int index, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = ((QTabBar*)self)->insertTab(index, text_qstr);
	return __r;
}
int QTabBar_insertTab2(void* self, int index, void* icon, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = ((QTabBar*)self)->insertTab(index, *(QIcon*)icon, text_qstr);
	return __r;
}
void QTabBar_removeTab(void* self, int index) {
	((QTabBar*)self)->removeTab(index);
}
void QTabBar_moveTab(void* self, int from, int to) {
	((QTabBar*)self)->moveTab(from, to);
}
int QTabBar_isTabEnabled(void* self, int index) {
	auto&& __r = ((QTabBar*)self)->isTabEnabled(index);
	return __r;
}
void QTabBar_setTabEnabled(void* self, int index, int enabled) {
	((QTabBar*)self)->setTabEnabled(index, enabled);
}
int QTabBar_isTabVisible(void* self, int index) {
	auto&& __r = ((QTabBar*)self)->isTabVisible(index);
	return __r;
}
void QTabBar_setTabVisible(void* self, int index, int visible) {
	((QTabBar*)self)->setTabVisible(index, visible);
}
char* QTabBar_tabText(void* self, int index, long long* __out_len) {
	auto&& __r = ((QTabBar*)self)->tabText(index);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QTabBar_setTabText(void* self, int index, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	((QTabBar*)self)->setTabText(index, text_qstr);
}
void* QTabBar_tabTextColor(void* self, int index) {
	auto&& __r = ((QTabBar*)self)->tabTextColor(index);
	return new QColor(__r);
}
void QTabBar_setTabTextColor(void* self, int index, void* color) {
	((QTabBar*)self)->setTabTextColor(index, *(QColor*)color);
}
void* QTabBar_tabIcon(void* self, int index) {
	auto&& __r = ((QTabBar*)self)->tabIcon(index);
	return new QIcon(__r);
}
void QTabBar_setTabIcon(void* self, int index, void* icon) {
	((QTabBar*)self)->setTabIcon(index, *(QIcon*)icon);
}
int QTabBar_elideMode(void* self) {
	auto&& __r = ((QTabBar*)self)->elideMode();
	return (int)(__r);
}
void QTabBar_setElideMode(void* self, int mode) {
	((QTabBar*)self)->setElideMode((Qt::TextElideMode)(mode));
}
void QTabBar_setTabToolTip(void* self, int index, const char* tip, long long tip_len) {
	QString tip_qstr = QString::fromUtf8(tip, tip_len);
	((QTabBar*)self)->setTabToolTip(index, tip_qstr);
}
char* QTabBar_tabToolTip(void* self, int index, long long* __out_len) {
	auto&& __r = ((QTabBar*)self)->tabToolTip(index);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QTabBar_setTabWhatsThis(void* self, int index, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	((QTabBar*)self)->setTabWhatsThis(index, text_qstr);
}
char* QTabBar_tabWhatsThis(void* self, int index, long long* __out_len) {
	auto&& __r = ((QTabBar*)self)->tabWhatsThis(index);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QTabBar_setTabData(void* self, int index, void* data) {
	((QTabBar*)self)->setTabData(index, *(QVariant*)data);
}
void* QTabBar_tabData(void* self, int index) {
	auto&& __r = ((QTabBar*)self)->tabData(index);
	return new QVariant(__r);
}
void* QTabBar_tabRect(void* self, int index) {
	auto&& __r = ((QTabBar*)self)->tabRect(index);
	return new QRect(__r);
}
int QTabBar_tabAt(void* self, void* pos) {
	auto&& __r = ((QTabBar*)self)->tabAt(*(QPoint*)pos);
	return __r;
}
int QTabBar_currentIndex(void* self) {
	auto&& __r = ((QTabBar*)self)->currentIndex();
	return __r;
}
int QTabBar_count(void* self) {
	auto&& __r = ((QTabBar*)self)->count();
	return __r;
}
void QTabBar_setDrawBase(void* self, int drawTheBase) {
	((QTabBar*)self)->setDrawBase(drawTheBase);
}
int QTabBar_drawBase(void* self) {
	auto&& __r = ((QTabBar*)self)->drawBase();
	return __r;
}
void* QTabBar_iconSize(void* self) {
	auto&& __r = ((QTabBar*)self)->iconSize();
	return new QSize(__r);
}
void QTabBar_setIconSize(void* self, void* size) {
	((QTabBar*)self)->setIconSize(*(QSize*)size);
}
int QTabBar_usesScrollButtons(void* self) {
	auto&& __r = ((QTabBar*)self)->usesScrollButtons();
	return __r;
}
void QTabBar_setUsesScrollButtons(void* self, int useButtons) {
	((QTabBar*)self)->setUsesScrollButtons(useButtons);
}
int QTabBar_tabsClosable(void* self) {
	auto&& __r = ((QTabBar*)self)->tabsClosable();
	return __r;
}
void QTabBar_setTabsClosable(void* self, int closable) {
	((QTabBar*)self)->setTabsClosable(closable);
}
void QTabBar_setTabButton(void* self, int index, int position, void* widget) {
	((QTabBar*)self)->setTabButton(index, (QTabBar::ButtonPosition)(position), (QWidget*)widget);
}
void* QTabBar_tabButton(void* self, int index, int position) {
	auto&& __r = ((QTabBar*)self)->tabButton(index, (QTabBar::ButtonPosition)(position));
	return (void*)(__r);
}
int QTabBar_selectionBehaviorOnRemove(void* self) {
	auto&& __r = ((QTabBar*)self)->selectionBehaviorOnRemove();
	return (int)(__r);
}
void QTabBar_setSelectionBehaviorOnRemove(void* self, int behavior) {
	((QTabBar*)self)->setSelectionBehaviorOnRemove((QTabBar::SelectionBehavior)(behavior));
}
int QTabBar_expanding(void* self) {
	auto&& __r = ((QTabBar*)self)->expanding();
	return __r;
}
void QTabBar_setExpanding(void* self, int enabled) {
	((QTabBar*)self)->setExpanding(enabled);
}
int QTabBar_isMovable(void* self) {
	auto&& __r = ((QTabBar*)self)->isMovable();
	return __r;
}
void QTabBar_setMovable(void* self, int movable) {
	((QTabBar*)self)->setMovable(movable);
}
int QTabBar_documentMode(void* self) {
	auto&& __r = ((QTabBar*)self)->documentMode();
	return __r;
}
void QTabBar_setDocumentMode(void* self, int set) {
	((QTabBar*)self)->setDocumentMode(set);
}
int QTabBar_autoHide(void* self) {
	auto&& __r = ((QTabBar*)self)->autoHide();
	return __r;
}
void QTabBar_setAutoHide(void* self, int hide) {
	((QTabBar*)self)->setAutoHide(hide);
}
int QTabBar_changeCurrentOnDrag(void* self) {
	auto&& __r = ((QTabBar*)self)->changeCurrentOnDrag();
	return __r;
}
void QTabBar_setChangeCurrentOnDrag(void* self, int change) {
	((QTabBar*)self)->setChangeCurrentOnDrag(change);
}
char* QTabBar_accessibleTabName(void* self, int index, long long* __out_len) {
	auto&& __r = ((QTabBar*)self)->accessibleTabName(index);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QTabBar_setAccessibleTabName(void* self, int index, const char* name, long long name_len) {
	QString name_qstr = QString::fromUtf8(name, name_len);
	((QTabBar*)self)->setAccessibleTabName(index, name_qstr);
}
void QTabBar_setCurrentIndex(void* self, int index) {
	((QTabBar*)self)->setCurrentIndex(index);
}
void QTabBar_currentChanged(void* self, int index) {
	((QTabBar*)self)->currentChanged(index);
}
void QTabBar_tabCloseRequested(void* self, int index) {
	((QTabBar*)self)->tabCloseRequested(index);
}
void QTabBar_tabMoved(void* self, int from, int to) {
	((QTabBar*)self)->tabMoved(from, to);
}
void QTabBar_tabBarClicked(void* self, int index) {
	((QTabBar*)self)->tabBarClicked(index);
}
void QTabBar_tabBarDoubleClicked(void* self, int index) {
	((QTabBar*)self)->tabBarDoubleClicked(index);
}
void QTabBar_setTabOrder(void* arg1, void* arg2) {
	QTabBar::setTabOrder((QWidget*)arg1, (QWidget*)arg2);
}
void* QTabBar_mouseGrabber() {
	auto&& __r = QTabBar::mouseGrabber();
	return (void*)(__r);
}
void* QTabBar_keyboardGrabber() {
	auto&& __r = QTabBar::keyboardGrabber();
	return (void*)(__r);
}
void* QTabBar_find(unsigned long long arg1) {
	auto&& __r = QTabBar::find(arg1);
	return (void*)(__r);
}
void* QTabBar_createWindowContainer(void* window) {
	auto&& __r = QTabBar::createWindowContainer((QWindow*)window);
	return (void*)(__r);
}
void QTabBar_updateMicroFocus(void* self) {
	((QTabBar__Go*)self)->updateMicroFocus();
}
void QTabBar_create(void* self) {
	((QTabBar__Go*)self)->create();
}
void QTabBar_destroy(void* self) {
	((QTabBar__Go*)self)->destroy();
}
int QTabBar_focusNextChild(void* self) {
	auto&& __r = ((QTabBar__Go*)self)->focusNextChild();
	return __r;
}
int QTabBar_focusPreviousChild(void* self) {
	auto&& __r = ((QTabBar__Go*)self)->focusPreviousChild();
	return __r;
}
void* QTabBar_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QTabBar::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QTabBar_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QTabBar::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QTabBar_sender(void* self) {
	auto&& __r = ((QTabBar__Go*)self)->sender();
	return (void*)(__r);
}
int QTabBar_senderSignalIndex(void* self) {
	auto&& __r = ((QTabBar__Go*)self)->senderSignalIndex();
	return __r;
}
int QTabBar_receivers(void* self, const char* signal) {
	auto&& __r = ((QTabBar__Go*)self)->receivers(signal);
	return __r;
}
int QTabBar_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QTabBar__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
double QTabBar_devicePixelRatioFScale() {
	auto&& __r = QTabBar::devicePixelRatioFScale();
	return __r;
}
int QTabBar_encodeMetricF(int metric, double value) {
	auto&& __r = QTabBar::encodeMetricF((QPaintDevice::PaintDeviceMetric)(metric), value);
	return __r;
}
double QTabBar_getDecodedMetricF(void* self, int metricA, int metricB) {
	auto&& __r = ((QTabBar__Go*)self)->getDecodedMetricF((QPaintDevice::PaintDeviceMetric)(metricA), (QPaintDevice::PaintDeviceMetric)(metricB));
	return __r;
}
char* QTabBar_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QTabBar::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QTabBar_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QTabBar::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QTabBar_createWindowContainer2(void* window, void* parent) {
	auto&& __r = QTabBar::createWindowContainer((QWindow*)window, (QWidget*)parent);
	return (void*)(__r);
}
void* QTabBar_createWindowContainer3(void* window, void* parent, int flags) {
	auto&& __r = QTabBar::createWindowContainer((QWindow*)window, (QWidget*)parent, QFlags<Qt::WindowType>(flags));
	return (void*)(__r);
}
void QTabBar_updateMicroFocus2(void* self, int query) {
	((QTabBar__Go*)self)->updateMicroFocus((Qt::InputMethodQuery)(query));
}
void QTabBar_create2(void* self, unsigned long long arg1) {
	((QTabBar__Go*)self)->create(arg1);
}
void QTabBar_create3(void* self, unsigned long long arg1, int initializeWindow) {
	((QTabBar__Go*)self)->create(arg1, initializeWindow);
}
void QTabBar_create4(void* self, unsigned long long arg1, int initializeWindow, int destroyOldWindow) {
	((QTabBar__Go*)self)->create(arg1, initializeWindow, destroyOldWindow);
}
void QTabBar_destroy2(void* self, int destroyWindow) {
	((QTabBar__Go*)self)->destroy(destroyWindow);
}
void QTabBar_destroy3(void* self, int destroyWindow, int destroySubWindows) {
	((QTabBar__Go*)self)->destroy(destroyWindow, destroySubWindows);
}
void* QTabBar_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QTabBar::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QTabBar_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QTabBar::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QTabBar_metaObject_default(void* self) {
	auto&& __r = ((QTabBar*)self)->QTabBar::metaObject();
	return (void*)(__r);
}
void* QTabBar_sizeHint_default(void* self) {
	auto&& __r = ((QTabBar*)self)->QTabBar::sizeHint();
	return new QSize(__r);
}
void* QTabBar_minimumSizeHint_default(void* self) {
	auto&& __r = ((QTabBar*)self)->QTabBar::minimumSizeHint();
	return new QSize(__r);
}
void* QTabBar_tabSizeHint_default(void* self, int index) {
	auto&& __r = ((QTabBar__Go*)self)->cppgoBase_tabSizeHint(index);
	return new QSize(__r);
}
void* QTabBar_minimumTabSizeHint_default(void* self, int index) {
	auto&& __r = ((QTabBar__Go*)self)->cppgoBase_minimumTabSizeHint(index);
	return new QSize(__r);
}
void QTabBar_tabInserted_default(void* self, int index) {
	((QTabBar__Go*)self)->cppgoBase_tabInserted(index);
}
void QTabBar_tabRemoved_default(void* self, int index) {
	((QTabBar__Go*)self)->cppgoBase_tabRemoved(index);
}
void QTabBar_tabLayoutChange_default(void* self) {
	((QTabBar__Go*)self)->cppgoBase_tabLayoutChange();
}
int QTabBar_event_default(void* self, void* arg1) {
	auto&& __r = ((QTabBar__Go*)self)->cppgoBase_event((QEvent*)arg1);
	return __r;
}
void QTabBar_resizeEvent_default(void* self, void* arg1) {
	((QTabBar__Go*)self)->cppgoBase_resizeEvent((QResizeEvent*)arg1);
}
void QTabBar_showEvent_default(void* self, void* arg1) {
	((QTabBar__Go*)self)->cppgoBase_showEvent((QShowEvent*)arg1);
}
void QTabBar_hideEvent_default(void* self, void* arg1) {
	((QTabBar__Go*)self)->cppgoBase_hideEvent((QHideEvent*)arg1);
}
void QTabBar_paintEvent_default(void* self, void* arg1) {
	((QTabBar__Go*)self)->cppgoBase_paintEvent((QPaintEvent*)arg1);
}
void QTabBar_mousePressEvent_default(void* self, void* arg1) {
	((QTabBar__Go*)self)->cppgoBase_mousePressEvent((QMouseEvent*)arg1);
}
void QTabBar_mouseMoveEvent_default(void* self, void* arg1) {
	((QTabBar__Go*)self)->cppgoBase_mouseMoveEvent((QMouseEvent*)arg1);
}
void QTabBar_mouseReleaseEvent_default(void* self, void* arg1) {
	((QTabBar__Go*)self)->cppgoBase_mouseReleaseEvent((QMouseEvent*)arg1);
}
void QTabBar_mouseDoubleClickEvent_default(void* self, void* arg1) {
	((QTabBar__Go*)self)->cppgoBase_mouseDoubleClickEvent((QMouseEvent*)arg1);
}
void QTabBar_wheelEvent_default(void* self, void* event) {
	((QTabBar__Go*)self)->cppgoBase_wheelEvent((QWheelEvent*)event);
}
void QTabBar_keyPressEvent_default(void* self, void* arg1) {
	((QTabBar__Go*)self)->cppgoBase_keyPressEvent((QKeyEvent*)arg1);
}
void QTabBar_changeEvent_default(void* self, void* arg1) {
	((QTabBar__Go*)self)->cppgoBase_changeEvent((QEvent*)arg1);
}
void QTabBar_timerEvent_default(void* self, void* event) {
	((QTabBar__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QTabBar_initStyleOption_default(void* self, void* option, int tabIndex) {
	((QTabBar__Go*)self)->cppgoBase_initStyleOption((QStyleOptionTab*)option, tabIndex);
}
int QTabBar_devType_default(void* self) {
	auto&& __r = ((QTabBar*)self)->QTabBar::devType();
	return __r;
}
void QTabBar_setVisible_default(void* self, int visible) {
	((QTabBar*)self)->QTabBar::setVisible(visible);
}
int QTabBar_heightForWidth_default(void* self, int arg1) {
	auto&& __r = ((QTabBar*)self)->QTabBar::heightForWidth(arg1);
	return __r;
}
int QTabBar_hasHeightForWidth_default(void* self) {
	auto&& __r = ((QTabBar*)self)->QTabBar::hasHeightForWidth();
	return __r;
}
void* QTabBar_paintEngine_default(void* self) {
	auto&& __r = ((QTabBar*)self)->QTabBar::paintEngine();
	return (void*)(__r);
}
void QTabBar_keyReleaseEvent_default(void* self, void* event) {
	((QTabBar__Go*)self)->cppgoBase_keyReleaseEvent((QKeyEvent*)event);
}
void QTabBar_focusInEvent_default(void* self, void* event) {
	((QTabBar__Go*)self)->cppgoBase_focusInEvent((QFocusEvent*)event);
}
void QTabBar_focusOutEvent_default(void* self, void* event) {
	((QTabBar__Go*)self)->cppgoBase_focusOutEvent((QFocusEvent*)event);
}
void QTabBar_enterEvent_default(void* self, void* event) {
	((QTabBar__Go*)self)->cppgoBase_enterEvent((QEnterEvent*)event);
}
void QTabBar_leaveEvent_default(void* self, void* event) {
	((QTabBar__Go*)self)->cppgoBase_leaveEvent((QEvent*)event);
}
void QTabBar_moveEvent_default(void* self, void* event) {
	((QTabBar__Go*)self)->cppgoBase_moveEvent((QMoveEvent*)event);
}
void QTabBar_closeEvent_default(void* self, void* event) {
	((QTabBar__Go*)self)->cppgoBase_closeEvent((QCloseEvent*)event);
}
void QTabBar_contextMenuEvent_default(void* self, void* event) {
	((QTabBar__Go*)self)->cppgoBase_contextMenuEvent((QContextMenuEvent*)event);
}
void QTabBar_tabletEvent_default(void* self, void* event) {
	((QTabBar__Go*)self)->cppgoBase_tabletEvent((QTabletEvent*)event);
}
void QTabBar_actionEvent_default(void* self, void* event) {
	((QTabBar__Go*)self)->cppgoBase_actionEvent((QActionEvent*)event);
}
void QTabBar_dragEnterEvent_default(void* self, void* event) {
	((QTabBar__Go*)self)->cppgoBase_dragEnterEvent((QDragEnterEvent*)event);
}
void QTabBar_dragMoveEvent_default(void* self, void* event) {
	((QTabBar__Go*)self)->cppgoBase_dragMoveEvent((QDragMoveEvent*)event);
}
void QTabBar_dragLeaveEvent_default(void* self, void* event) {
	((QTabBar__Go*)self)->cppgoBase_dragLeaveEvent((QDragLeaveEvent*)event);
}
void QTabBar_dropEvent_default(void* self, void* event) {
	((QTabBar__Go*)self)->cppgoBase_dropEvent((QDropEvent*)event);
}
int QTabBar_metric_default(void* self, int arg1) {
	auto&& __r = ((QTabBar__Go*)self)->cppgoBase_metric(arg1);
	return __r;
}
void QTabBar_initPainter_default(void* self, void* painter) {
	((QTabBar__Go*)self)->cppgoBase_initPainter((QPainter*)painter);
}
void* QTabBar_redirected_default(void* self, void* offset) {
	auto&& __r = ((QTabBar__Go*)self)->cppgoBase_redirected((QPoint*)offset);
	return (void*)(__r);
}
void* QTabBar_sharedPainter_default(void* self) {
	auto&& __r = ((QTabBar__Go*)self)->cppgoBase_sharedPainter();
	return (void*)(__r);
}
void QTabBar_inputMethodEvent_default(void* self, void* arg1) {
	((QTabBar__Go*)self)->cppgoBase_inputMethodEvent((QInputMethodEvent*)arg1);
}
void* QTabBar_inputMethodQuery_default(void* self, int arg1) {
	auto&& __r = ((QTabBar*)self)->QTabBar::inputMethodQuery((Qt::InputMethodQuery)(arg1));
	return new QVariant(__r);
}
int QTabBar_focusNextPrevChild_default(void* self, int next) {
	auto&& __r = ((QTabBar__Go*)self)->cppgoBase_focusNextPrevChild(next);
	return __r;
}
int QTabBar_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QTabBar*)self)->QTabBar::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QTabBar_childEvent_default(void* self, void* event) {
	((QTabBar__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QTabBar_customEvent_default(void* self, void* event) {
	((QTabBar__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QTabBar_connectNotify_default(void* self, void* signal) {
	((QTabBar__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QTabBar_disconnectNotify_default(void* self, void* signal) {
	((QTabBar__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QTabBar_connect_currentChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTabBar*)self, &QTabBar::currentChanged, (QTabBar*)self, [__id](int index){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, index);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTabBar);
	__conns_QTabBar[__id] = __c;
}
void QTabBar_connect_tabCloseRequested(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTabBar*)self, &QTabBar::tabCloseRequested, (QTabBar*)self, [__id](int index){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, index);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTabBar);
	__conns_QTabBar[__id] = __c;
}
void QTabBar_connect_tabMoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTabBar*)self, &QTabBar::tabMoved, (QTabBar*)self, [__id](int from, int to){
		if (cppgo_sig_void_int_int) ((void(*)(long long, int, int))cppgo_sig_void_int_int)(__id, from, to);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTabBar);
	__conns_QTabBar[__id] = __c;
}
void QTabBar_connect_tabBarClicked(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTabBar*)self, &QTabBar::tabBarClicked, (QTabBar*)self, [__id](int index){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, index);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTabBar);
	__conns_QTabBar[__id] = __c;
}
void QTabBar_connect_tabBarDoubleClicked(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTabBar*)self, &QTabBar::tabBarDoubleClicked, (QTabBar*)self, [__id](int index){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, index);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTabBar);
	__conns_QTabBar[__id] = __c;
}
void QTabBar_connect_windowTitleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTabBar*)self, &QTabBar::windowTitleChanged, (QTabBar*)self, [__id](const QString& title){
		QByteArray title_ba = title.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)title_ba.constData(), title_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTabBar);
	__conns_QTabBar[__id] = __c;
}
void QTabBar_connect_windowIconChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTabBar*)self, &QTabBar::windowIconChanged, (QTabBar*)self, [__id](const QIcon& icon){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&icon);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTabBar);
	__conns_QTabBar[__id] = __c;
}
void QTabBar_connect_windowIconTextChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTabBar*)self, &QTabBar::windowIconTextChanged, (QTabBar*)self, [__id](const QString& iconText){
		QByteArray iconText_ba = iconText.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)iconText_ba.constData(), iconText_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTabBar);
	__conns_QTabBar[__id] = __c;
}
void QTabBar_connect_customContextMenuRequested(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTabBar*)self, &QTabBar::customContextMenuRequested, (QTabBar*)self, [__id](const QPoint& pos){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&pos);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTabBar);
	__conns_QTabBar[__id] = __c;
}
void QTabBar_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTabBar*)self, &QTabBar::destroyed, (QTabBar*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTabBar);
	__conns_QTabBar[__id] = __c;
}
void QTabBar_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTabBar*)self, &QTabBar::destroyed, (QTabBar*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTabBar);
	__conns_QTabBar[__id] = __c;
}
void QTabBar_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTabBar*)self, &QTabBar::objectNameChanged, (QTabBar*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTabBar);
	__conns_QTabBar[__id] = __c;
}
void QTabBar__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QTabBar);
		auto __it = __conns_QTabBar.find(id);
		if (__it == __conns_QTabBar.end()) return;
		__c = __it->second;
		__conns_QTabBar.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QTabBar__Go_new(long long id) { return new QTabBar__Go(id); }
void QTabBar_delete(void* self) { delete (QTabBar*)self; }

}

#include <QtWidgets/qstyleoption.h>
#include <QtWidgets/qwidget.h>
#include <QtGui/qicon.h>
#include <QtCore/qsize.h>
#include <QtCore/qstring.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qpoint.h>
#include <QtGui/qfont.h>
#include "gen_qstyleoptiontoolbutton.h"

extern "C" {

void* QStyleOptionToolButton_base(void* self) { return (void*)static_cast<QStyleOptionComplex*>((QStyleOptionToolButton*)self); }
void QStyleOptionToolButton_chain(void* self, void** out) {
	QStyleOptionToolButton* __p = (QStyleOptionToolButton*)self;
	QStyleOptionComplex* __b0 = static_cast<QStyleOptionComplex*>(__p); out[0] = (void*)__b0;
	QStyleOption* __b1 = static_cast<QStyleOption*>(__b0); out[1] = (void*)__b1;
}
void* QStyleOptionToolButton_new() {
	return new QStyleOptionToolButton();
}
void* QStyleOptionToolButton_new2(void* other) {
	return new QStyleOptionToolButton(*(QStyleOptionToolButton*)other);
}
void* QStyleOptionToolButton_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QStyleOptionToolButton*)self)->operator=(*(QStyleOptionToolButton*)arg1);
	return (void*)&(__r);
}
int QStyleOptionToolButton_features_get(void* self) {
	auto __r = ((QStyleOptionToolButton*)self)->features;
	return (int)(__r);
}
void QStyleOptionToolButton_features_set(void* self, int features) {
	((QStyleOptionToolButton*)self)->features = QFlags<QStyleOptionToolButton::ToolButtonFeature>(features);
}
void* QStyleOptionToolButton_icon_get(void* self) {
	auto __r = ((QStyleOptionToolButton*)self)->icon;
	return new QIcon(__r);
}
void QStyleOptionToolButton_icon_set(void* self, void* icon) {
	((QStyleOptionToolButton*)self)->icon = *(QIcon*)icon;
}
void* QStyleOptionToolButton_iconSize_get(void* self) {
	auto __r = ((QStyleOptionToolButton*)self)->iconSize;
	return new QSize(__r);
}
void QStyleOptionToolButton_iconSize_set(void* self, void* iconSize) {
	((QStyleOptionToolButton*)self)->iconSize = *(QSize*)iconSize;
}
char* QStyleOptionToolButton_text_get(void* self, long long* __out_len) {
	auto __r = ((QStyleOptionToolButton*)self)->text;
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QStyleOptionToolButton_text_set(void* self, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	((QStyleOptionToolButton*)self)->text = text_qstr;
}
int QStyleOptionToolButton_arrowType_get(void* self) {
	auto __r = ((QStyleOptionToolButton*)self)->arrowType;
	return (int)(__r);
}
void QStyleOptionToolButton_arrowType_set(void* self, int arrowType) {
	((QStyleOptionToolButton*)self)->arrowType = (Qt::ArrowType)(arrowType);
}
int QStyleOptionToolButton_toolButtonStyle_get(void* self) {
	auto __r = ((QStyleOptionToolButton*)self)->toolButtonStyle;
	return (int)(__r);
}
void QStyleOptionToolButton_toolButtonStyle_set(void* self, int toolButtonStyle) {
	((QStyleOptionToolButton*)self)->toolButtonStyle = (Qt::ToolButtonStyle)(toolButtonStyle);
}
void* QStyleOptionToolButton_pos_get(void* self) {
	auto __r = ((QStyleOptionToolButton*)self)->pos;
	return new QPoint(__r);
}
void QStyleOptionToolButton_pos_set(void* self, void* pos) {
	((QStyleOptionToolButton*)self)->pos = *(QPoint*)pos;
}
void* QStyleOptionToolButton_font_get(void* self) {
	auto __r = ((QStyleOptionToolButton*)self)->font;
	return new QFont(__r);
}
void QStyleOptionToolButton_font_set(void* self, void* font) {
	((QStyleOptionToolButton*)self)->font = *(QFont*)font;
}
void QStyleOptionToolButton_delete(void* self) { delete (QStyleOptionToolButton*)self; }

}

#include <QtWidgets/qbuttongroup.h>
#include <QtCore/qobject.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtWidgets/qabstractbutton.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qvariant.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qbuttongroup.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QButtonGroup__Go : public QButtonGroup {
public:
	long long __goId;
	unsigned char __goOn[8] = {};
	template<typename... A> QButtonGroup__Go(long long id, A&&... a) : QButtonGroup(std::forward<A>(a)...), __goId(id) {}
	using QButtonGroup::sender;
	using QButtonGroup::senderSignalIndex;
	using QButtonGroup::receivers;
	using QButtonGroup::isSignalConnected;
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QButtonGroup::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QButtonGroup::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QButtonGroup::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QButtonGroup::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QButtonGroup::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QButtonGroup::metaObject();
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QButtonGroup::event(event);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QButtonGroup::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QButtonGroup::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QButtonGroup::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QButtonGroup::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QButtonGroup::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QButtonGroup::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QButtonGroup;
static std::mutex __connsMu_QButtonGroup;

extern "C" {

void* QButtonGroup_base(void* self) { return (void*)static_cast<QObject*>((QButtonGroup*)self); }
void QButtonGroup__set_override(void* self, int idx) { ((QButtonGroup__Go*)self)->__goOn[idx] = 1; }
void* QButtonGroup_fromQObject(void* o) { return (void*)qobject_cast<QButtonGroup*>((QObject*)o); }
void* QButtonGroup_new(long long __goid) {
	return new QButtonGroup__Go(__goid);
}
void* QButtonGroup_new2(long long __goid, void* parent) {
	return new QButtonGroup__Go(__goid, (QObject*)parent);
}
char* QButtonGroup_tr(const char* s, long long* __out_len) {
	auto&& __r = QButtonGroup::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QButtonGroup_setExclusive(void* self, int arg1) {
	((QButtonGroup*)self)->setExclusive(arg1);
}
int QButtonGroup_exclusive(void* self) {
	auto&& __r = ((QButtonGroup*)self)->exclusive();
	return __r;
}
void QButtonGroup_addButton(void* self, void* arg1) {
	((QButtonGroup*)self)->addButton((QAbstractButton*)arg1);
}
void QButtonGroup_removeButton(void* self, void* arg1) {
	((QButtonGroup*)self)->removeButton((QAbstractButton*)arg1);
}
void** QButtonGroup_buttons(void* self, long long* __out_len) {
	auto&& __r = ((QButtonGroup*)self)->buttons();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
void* QButtonGroup_checkedButton(void* self) {
	auto&& __r = ((QButtonGroup*)self)->checkedButton();
	return (void*)(__r);
}
void* QButtonGroup_button(void* self, int id) {
	auto&& __r = ((QButtonGroup*)self)->button(id);
	return (void*)(__r);
}
void QButtonGroup_setId(void* self, void* button, int id) {
	((QButtonGroup*)self)->setId((QAbstractButton*)button, id);
}
int QButtonGroup_id(void* self, void* button) {
	auto&& __r = ((QButtonGroup*)self)->id((QAbstractButton*)button);
	return __r;
}
int QButtonGroup_checkedId(void* self) {
	auto&& __r = ((QButtonGroup*)self)->checkedId();
	return __r;
}
void QButtonGroup_buttonClicked(void* self, void* arg1) {
	((QButtonGroup*)self)->buttonClicked((QAbstractButton*)arg1);
}
void QButtonGroup_buttonPressed(void* self, void* arg1) {
	((QButtonGroup*)self)->buttonPressed((QAbstractButton*)arg1);
}
void QButtonGroup_buttonReleased(void* self, void* arg1) {
	((QButtonGroup*)self)->buttonReleased((QAbstractButton*)arg1);
}
void QButtonGroup_buttonToggled(void* self, void* arg1, int arg2) {
	((QButtonGroup*)self)->buttonToggled((QAbstractButton*)arg1, arg2);
}
void QButtonGroup_idClicked(void* self, int arg1) {
	((QButtonGroup*)self)->idClicked(arg1);
}
void QButtonGroup_idPressed(void* self, int arg1) {
	((QButtonGroup*)self)->idPressed(arg1);
}
void QButtonGroup_idReleased(void* self, int arg1) {
	((QButtonGroup*)self)->idReleased(arg1);
}
void QButtonGroup_idToggled(void* self, int arg1, int arg2) {
	((QButtonGroup*)self)->idToggled(arg1, arg2);
}
void* QButtonGroup_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QButtonGroup::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QButtonGroup_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QButtonGroup::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QButtonGroup_sender(void* self) {
	auto&& __r = ((QButtonGroup__Go*)self)->sender();
	return (void*)(__r);
}
int QButtonGroup_senderSignalIndex(void* self) {
	auto&& __r = ((QButtonGroup__Go*)self)->senderSignalIndex();
	return __r;
}
int QButtonGroup_receivers(void* self, const char* signal) {
	auto&& __r = ((QButtonGroup__Go*)self)->receivers(signal);
	return __r;
}
int QButtonGroup_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QButtonGroup__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QButtonGroup_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QButtonGroup::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QButtonGroup_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QButtonGroup::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QButtonGroup_addButton2(void* self, void* arg1, int id) {
	((QButtonGroup*)self)->addButton((QAbstractButton*)arg1, id);
}
void* QButtonGroup_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QButtonGroup::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QButtonGroup_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QButtonGroup::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QButtonGroup_metaObject_default(void* self) {
	auto&& __r = ((QButtonGroup*)self)->QButtonGroup::metaObject();
	return (void*)(__r);
}
int QButtonGroup_event_default(void* self, void* event) {
	auto&& __r = ((QButtonGroup*)self)->QButtonGroup::event((QEvent*)event);
	return __r;
}
int QButtonGroup_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QButtonGroup*)self)->QButtonGroup::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QButtonGroup_timerEvent_default(void* self, void* event) {
	((QButtonGroup__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QButtonGroup_childEvent_default(void* self, void* event) {
	((QButtonGroup__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QButtonGroup_customEvent_default(void* self, void* event) {
	((QButtonGroup__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QButtonGroup_connectNotify_default(void* self, void* signal) {
	((QButtonGroup__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QButtonGroup_disconnectNotify_default(void* self, void* signal) {
	((QButtonGroup__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QButtonGroup_connect_buttonClicked(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QButtonGroup*)self, &QButtonGroup::buttonClicked, (QButtonGroup*)self, [__id](QAbstractButton* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QButtonGroup);
	__conns_QButtonGroup[__id] = __c;
}
void QButtonGroup_connect_buttonPressed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QButtonGroup*)self, &QButtonGroup::buttonPressed, (QButtonGroup*)self, [__id](QAbstractButton* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QButtonGroup);
	__conns_QButtonGroup[__id] = __c;
}
void QButtonGroup_connect_buttonReleased(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QButtonGroup*)self, &QButtonGroup::buttonReleased, (QButtonGroup*)self, [__id](QAbstractButton* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QButtonGroup);
	__conns_QButtonGroup[__id] = __c;
}
void QButtonGroup_connect_buttonToggled(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QButtonGroup*)self, &QButtonGroup::buttonToggled, (QButtonGroup*)self, [__id](QAbstractButton* arg1, bool arg2){
		if (cppgo_sig_void_voidp_int) ((void(*)(long long, void*, int))cppgo_sig_void_voidp_int)(__id, (void*)arg1, arg2);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QButtonGroup);
	__conns_QButtonGroup[__id] = __c;
}
void QButtonGroup_connect_idClicked(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QButtonGroup*)self, &QButtonGroup::idClicked, (QButtonGroup*)self, [__id](int arg1){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QButtonGroup);
	__conns_QButtonGroup[__id] = __c;
}
void QButtonGroup_connect_idPressed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QButtonGroup*)self, &QButtonGroup::idPressed, (QButtonGroup*)self, [__id](int arg1){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QButtonGroup);
	__conns_QButtonGroup[__id] = __c;
}
void QButtonGroup_connect_idReleased(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QButtonGroup*)self, &QButtonGroup::idReleased, (QButtonGroup*)self, [__id](int arg1){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QButtonGroup);
	__conns_QButtonGroup[__id] = __c;
}
void QButtonGroup_connect_idToggled(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QButtonGroup*)self, &QButtonGroup::idToggled, (QButtonGroup*)self, [__id](int arg1, bool arg2){
		if (cppgo_sig_void_int_int) ((void(*)(long long, int, int))cppgo_sig_void_int_int)(__id, arg1, arg2);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QButtonGroup);
	__conns_QButtonGroup[__id] = __c;
}
void QButtonGroup_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QButtonGroup*)self, &QButtonGroup::destroyed, (QButtonGroup*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QButtonGroup);
	__conns_QButtonGroup[__id] = __c;
}
void QButtonGroup_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QButtonGroup*)self, &QButtonGroup::destroyed, (QButtonGroup*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QButtonGroup);
	__conns_QButtonGroup[__id] = __c;
}
void QButtonGroup_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QButtonGroup*)self, &QButtonGroup::objectNameChanged, (QButtonGroup*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QButtonGroup);
	__conns_QButtonGroup[__id] = __c;
}
void QButtonGroup__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QButtonGroup);
		auto __it = __conns_QButtonGroup.find(id);
		if (__it == __conns_QButtonGroup.end()) return;
		__c = __it->second;
		__conns_QButtonGroup.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QButtonGroup__Go_new(long long id, void* parent) { return new QButtonGroup__Go(id, (QObject*)parent); }
void QButtonGroup_delete(void* self) { delete (QButtonGroup*)self; }

}

#include <QtWidgets/qfontcombobox.h>
#include <QtWidgets/qcombobox.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtGui/qfontdatabase.h>
#include <QtGui/qfont.h>
#include <QtCore/qsize.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qvariant.h>
#include <QtWidgets/qlineedit.h>
#include <QtGui/qvalidator.h>
#include <QtWidgets/qcompleter.h>
#include <QtWidgets/qabstractitemdelegate.h>
#include <QtCore/qabstractitemmodel.h>
#include <QtGui/qicon.h>
#include <QtWidgets/qabstractitemview.h>
#include <QtGui/qevent.h>
#include <QtWidgets/qstyleoption.h>
#include <QtWidgets/qstyle.h>
#include <QtCore/qrect.h>
#include <QtCore/qpoint.h>
#include <QtGui/qregion.h>
#include <QtGui/qpalette.h>
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
#include "gen_qfontcombobox.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QFontComboBox__Go : public QFontComboBox {
public:
	long long __goId;
	unsigned char __goOn[52] = {};
	template<typename... A> QFontComboBox__Go(long long id, A&&... a) : QFontComboBox(std::forward<A>(a)...), __goId(id) {}
	using QFontComboBox::updateMicroFocus;
	using QFontComboBox::create;
	using QFontComboBox::destroy;
	using QFontComboBox::focusNextChild;
	using QFontComboBox::focusPreviousChild;
	using QFontComboBox::sender;
	using QFontComboBox::senderSignalIndex;
	using QFontComboBox::receivers;
	using QFontComboBox::isSignalConnected;
	using QFontComboBox::getDecodedMetricF;
	bool cppgoBase_event(QEvent* __a0) { return QFontComboBox::event(__a0); }
	void cppgoBase_focusInEvent(QFocusEvent* __a0) { QFontComboBox::focusInEvent(__a0); }
	void cppgoBase_focusOutEvent(QFocusEvent* __a0) { QFontComboBox::focusOutEvent(__a0); }
	void cppgoBase_changeEvent(QEvent* __a0) { QFontComboBox::changeEvent(__a0); }
	void cppgoBase_resizeEvent(QResizeEvent* __a0) { QFontComboBox::resizeEvent(__a0); }
	void cppgoBase_paintEvent(QPaintEvent* __a0) { QFontComboBox::paintEvent(__a0); }
	void cppgoBase_showEvent(QShowEvent* __a0) { QFontComboBox::showEvent(__a0); }
	void cppgoBase_hideEvent(QHideEvent* __a0) { QFontComboBox::hideEvent(__a0); }
	void cppgoBase_mousePressEvent(QMouseEvent* __a0) { QFontComboBox::mousePressEvent(__a0); }
	void cppgoBase_mouseReleaseEvent(QMouseEvent* __a0) { QFontComboBox::mouseReleaseEvent(__a0); }
	void cppgoBase_keyPressEvent(QKeyEvent* __a0) { QFontComboBox::keyPressEvent(__a0); }
	void cppgoBase_keyReleaseEvent(QKeyEvent* __a0) { QFontComboBox::keyReleaseEvent(__a0); }
	void cppgoBase_wheelEvent(QWheelEvent* __a0) { QFontComboBox::wheelEvent(__a0); }
	void cppgoBase_contextMenuEvent(QContextMenuEvent* __a0) { QFontComboBox::contextMenuEvent(__a0); }
	void cppgoBase_inputMethodEvent(QInputMethodEvent* __a0) { QFontComboBox::inputMethodEvent(__a0); }
	void cppgoBase_initStyleOption(QStyleOptionComboBox* __a0) { QFontComboBox::initStyleOption(__a0); }
	void cppgoBase_mouseDoubleClickEvent(QMouseEvent* __a0) { QFontComboBox::mouseDoubleClickEvent(__a0); }
	void cppgoBase_mouseMoveEvent(QMouseEvent* __a0) { QFontComboBox::mouseMoveEvent(__a0); }
	void cppgoBase_enterEvent(QEnterEvent* __a0) { QFontComboBox::enterEvent(__a0); }
	void cppgoBase_leaveEvent(QEvent* __a0) { QFontComboBox::leaveEvent(__a0); }
	void cppgoBase_moveEvent(QMoveEvent* __a0) { QFontComboBox::moveEvent(__a0); }
	void cppgoBase_closeEvent(QCloseEvent* __a0) { QFontComboBox::closeEvent(__a0); }
	void cppgoBase_tabletEvent(QTabletEvent* __a0) { QFontComboBox::tabletEvent(__a0); }
	void cppgoBase_actionEvent(QActionEvent* __a0) { QFontComboBox::actionEvent(__a0); }
	void cppgoBase_dragEnterEvent(QDragEnterEvent* __a0) { QFontComboBox::dragEnterEvent(__a0); }
	void cppgoBase_dragMoveEvent(QDragMoveEvent* __a0) { QFontComboBox::dragMoveEvent(__a0); }
	void cppgoBase_dragLeaveEvent(QDragLeaveEvent* __a0) { QFontComboBox::dragLeaveEvent(__a0); }
	void cppgoBase_dropEvent(QDropEvent* __a0) { QFontComboBox::dropEvent(__a0); }
	int cppgoBase_metric(int __a0) { return QFontComboBox::metric((QPaintDevice::PaintDeviceMetric)__a0); }
	void cppgoBase_initPainter(QPainter* __a0) { QFontComboBox::initPainter(__a0); }
	QPaintDevice* cppgoBase_redirected(QPoint* __a0) { return QFontComboBox::redirected(__a0); }
	QPainter* cppgoBase_sharedPainter() { return QFontComboBox::sharedPainter(); }
	bool cppgoBase_focusNextPrevChild(bool __a0) { return QFontComboBox::focusNextPrevChild(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QFontComboBox::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QFontComboBox::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QFontComboBox::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QFontComboBox::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QFontComboBox::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QFontComboBox::metaObject();
	}
	QSize sizeHint() const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 729690261457684101LL, &__h); if (__h) return *(QSize*)__r; }
		return QFontComboBox::sizeHint();
	}
	bool event(QEvent* e) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)e, &__h); if (__h) return __r; }
		return QFontComboBox::event(e);
	}
	void setModel(QAbstractItemModel* model) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4283922507359547360LL, (void*)model, &__h); if (__h) return; }
		QFontComboBox::setModel(model);
	}
	QSize minimumSizeHint() const override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 7049264599837532491LL, &__h); if (__h) return *(QSize*)__r; }
		return QFontComboBox::minimumSizeHint();
	}
	void showPopup() override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, -977583612921719482LL, &__h); if (__h) return; }
		QFontComboBox::showPopup();
	}
	void hidePopup() override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 3308350117881191909LL, &__h); if (__h) return; }
		QFontComboBox::hidePopup();
	}
	QVariant inputMethodQuery(Qt::InputMethodQuery arg1) const override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7345853097340709638LL, (int)(arg1), &__h); if (__h) return *(QVariant*)__r; }
		return QFontComboBox::inputMethodQuery(arg1);
	}
	void focusInEvent(QFocusEvent* e) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 975043767033552552LL, (void*)e, &__h); if (__h) return; }
		QFontComboBox::focusInEvent(e);
	}
	void focusOutEvent(QFocusEvent* e) override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -385596010576048993LL, (void*)e, &__h); if (__h) return; }
		QFontComboBox::focusOutEvent(e);
	}
	void changeEvent(QEvent* e) override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4959376063826165943LL, (void*)e, &__h); if (__h) return; }
		QFontComboBox::changeEvent(e);
	}
	void resizeEvent(QResizeEvent* e) override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4302917285655014661LL, (void*)e, &__h); if (__h) return; }
		QFontComboBox::resizeEvent(e);
	}
	void paintEvent(QPaintEvent* e) override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5290442163167458465LL, (void*)e, &__h); if (__h) return; }
		QFontComboBox::paintEvent(e);
	}
	void showEvent(QShowEvent* e) override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6825369174751916LL, (void*)e, &__h); if (__h) return; }
		QFontComboBox::showEvent(e);
	}
	void hideEvent(QHideEvent* e) override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -199718707876719129LL, (void*)e, &__h); if (__h) return; }
		QFontComboBox::hideEvent(e);
	}
	void mousePressEvent(QMouseEvent* e) override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -332090898078543319LL, (void*)e, &__h); if (__h) return; }
		QFontComboBox::mousePressEvent(e);
	}
	void mouseReleaseEvent(QMouseEvent* e) override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4415920049297795213LL, (void*)e, &__h); if (__h) return; }
		QFontComboBox::mouseReleaseEvent(e);
	}
	void keyPressEvent(QKeyEvent* e) override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1023791977679793185LL, (void*)e, &__h); if (__h) return; }
		QFontComboBox::keyPressEvent(e);
	}
	void keyReleaseEvent(QKeyEvent* e) override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 772598475950050299LL, (void*)e, &__h); if (__h) return; }
		QFontComboBox::keyReleaseEvent(e);
	}
	void wheelEvent(QWheelEvent* e) override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1919988340934511316LL, (void*)e, &__h); if (__h) return; }
		QFontComboBox::wheelEvent(e);
	}
	void contextMenuEvent(QContextMenuEvent* e) override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6698043222227006015LL, (void*)e, &__h); if (__h) return; }
		QFontComboBox::contextMenuEvent(e);
	}
	void inputMethodEvent(QInputMethodEvent* arg1) override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5438819661942843222LL, (void*)arg1, &__h); if (__h) return; }
		QFontComboBox::inputMethodEvent(arg1);
	}
	void initStyleOption(QStyleOptionComboBox* option) const override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -488693769771308531LL, (void*)option, &__h); if (__h) return; }
		QFontComboBox::initStyleOption(option);
	}
	int devType() const override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8659679893868524044LL, &__h); if (__h) return __r; }
		return QFontComboBox::devType();
	}
	void setVisible(bool visible) override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 2110182719375245575LL, visible, &__h); if (__h) return; }
		QFontComboBox::setVisible(visible);
	}
	int heightForWidth(int arg1) const override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 4311534060495737339LL, arg1, &__h); if (__h) return __r; }
		return QFontComboBox::heightForWidth(arg1);
	}
	bool hasHeightForWidth() const override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8266762397995423915LL, &__h); if (__h) return __r; }
		return QFontComboBox::hasHeightForWidth();
	}
	QPaintEngine* paintEngine() const override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 6398007404756346381LL, &__h); if (__h) return (QPaintEngine*)__r; }
		return QFontComboBox::paintEngine();
	}
	void mouseDoubleClickEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4029073905697023887LL, (void*)event, &__h); if (__h) return; }
		QFontComboBox::mouseDoubleClickEvent(event);
	}
	void mouseMoveEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4255494004099806581LL, (void*)event, &__h); if (__h) return; }
		QFontComboBox::mouseMoveEvent(event);
	}
	void enterEvent(QEnterEvent* event) override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6316270346668237895LL, (void*)event, &__h); if (__h) return; }
		QFontComboBox::enterEvent(event);
	}
	void leaveEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -551070009623711640LL, (void*)event, &__h); if (__h) return; }
		QFontComboBox::leaveEvent(event);
	}
	void moveEvent(QMoveEvent* event) override {
		int __h = 0;
		if (__goOn[32] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -9161949150489525396LL, (void*)event, &__h); if (__h) return; }
		QFontComboBox::moveEvent(event);
	}
	void closeEvent(QCloseEvent* event) override {
		int __h = 0;
		if (__goOn[33] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1989937050019464783LL, (void*)event, &__h); if (__h) return; }
		QFontComboBox::closeEvent(event);
	}
	void tabletEvent(QTabletEvent* event) override {
		int __h = 0;
		if (__goOn[34] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -212068866860701141LL, (void*)event, &__h); if (__h) return; }
		QFontComboBox::tabletEvent(event);
	}
	void actionEvent(QActionEvent* event) override {
		int __h = 0;
		if (__goOn[35] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1697444140088555581LL, (void*)event, &__h); if (__h) return; }
		QFontComboBox::actionEvent(event);
	}
	void dragEnterEvent(QDragEnterEvent* event) override {
		int __h = 0;
		if (__goOn[36] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 206122715919917589LL, (void*)event, &__h); if (__h) return; }
		QFontComboBox::dragEnterEvent(event);
	}
	void dragMoveEvent(QDragMoveEvent* event) override {
		int __h = 0;
		if (__goOn[37] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1781545137352745978LL, (void*)event, &__h); if (__h) return; }
		QFontComboBox::dragMoveEvent(event);
	}
	void dragLeaveEvent(QDragLeaveEvent* event) override {
		int __h = 0;
		if (__goOn[38] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -3139904158373725006LL, (void*)event, &__h); if (__h) return; }
		QFontComboBox::dragLeaveEvent(event);
	}
	void dropEvent(QDropEvent* event) override {
		int __h = 0;
		if (__goOn[39] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 9107058682754591716LL, (void*)event, &__h); if (__h) return; }
		QFontComboBox::dropEvent(event);
	}
	bool nativeEvent(const QByteArray& eventType, void* message, qintptr* result) override {
		int __h = 0;
		if (__goOn[40] && cppgo_virt_int_voidp_voidp_longlongp) { int __r = ((int(*)(long long, long long, void*, void*, long long*, int*))cppgo_virt_int_voidp_voidp_longlongp)(__goId, -2039166586246443706LL, (void*)&eventType, (void*)message, (long long*)result, &__h); if (__h) return __r; }
		return QFontComboBox::nativeEvent(eventType, message, result);
	}
	int metric(QPaintDevice::PaintDeviceMetric arg1) const override {
		int __h = 0;
		if (__goOn[41] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 427629649267311777LL, (int)(arg1), &__h); if (__h) return __r; }
		return QFontComboBox::metric(arg1);
	}
	void initPainter(QPainter* painter) const override {
		int __h = 0;
		if (__goOn[42] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6477500296980050210LL, (void*)painter, &__h); if (__h) return; }
		QFontComboBox::initPainter(painter);
	}
	QPaintDevice* redirected(QPoint* offset) const override {
		int __h = 0;
		if (__goOn[43] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 4127240633310655048LL, (void*)offset, &__h); if (__h) return (QPaintDevice*)__r; }
		return QFontComboBox::redirected(offset);
	}
	QPainter* sharedPainter() const override {
		int __h = 0;
		if (__goOn[44] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 5792771328741774991LL, &__h); if (__h) return (QPainter*)__r; }
		return QFontComboBox::sharedPainter();
	}
	bool focusNextPrevChild(bool next) override {
		int __h = 0;
		if (__goOn[45] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 5898328517505784191LL, next, &__h); if (__h) return __r; }
		return QFontComboBox::focusNextPrevChild(next);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[46] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QFontComboBox::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[47] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QFontComboBox::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[48] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QFontComboBox::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[49] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QFontComboBox::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[50] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QFontComboBox::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[51] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QFontComboBox::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QFontComboBox;
static std::mutex __connsMu_QFontComboBox;

extern "C" {

void* QFontComboBox_base(void* self) { return (void*)static_cast<QComboBox*>((QFontComboBox*)self); }
void QFontComboBox_chain(void* self, void** out) {
	QFontComboBox* __p = (QFontComboBox*)self;
	QComboBox* __b0 = static_cast<QComboBox*>(__p); out[0] = (void*)__b0;
	QWidget* __b1 = static_cast<QWidget*>(__b0); out[1] = (void*)__b1;
	QObject* __b2 = static_cast<QObject*>(__b1); out[2] = (void*)__b2;
}
void QFontComboBox__set_override(void* self, int idx) { ((QFontComboBox__Go*)self)->__goOn[idx] = 1; }
void* QFontComboBox_fromQObject(void* o) { return (void*)qobject_cast<QFontComboBox*>((QObject*)o); }
void* QFontComboBox_new(long long __goid, void* parent) {
	return new QFontComboBox__Go(__goid, (QWidget*)parent);
}
void* QFontComboBox_new2(long long __goid) {
	return new QFontComboBox__Go(__goid);
}
char* QFontComboBox_tr(const char* s, long long* __out_len) {
	auto&& __r = QFontComboBox::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QFontComboBox_setWritingSystem(void* self, int arg1) {
	((QFontComboBox*)self)->setWritingSystem((QFontDatabase::WritingSystem)(arg1));
}
int QFontComboBox_writingSystem(void* self) {
	auto&& __r = ((QFontComboBox*)self)->writingSystem();
	return (int)(__r);
}
void QFontComboBox_setFontFilters(void* self, int filters) {
	((QFontComboBox*)self)->setFontFilters(QFlags<QFontComboBox::FontFilter>(filters));
}
int QFontComboBox_fontFilters(void* self) {
	auto&& __r = ((QFontComboBox*)self)->fontFilters();
	return (int)(__r);
}
void* QFontComboBox_currentFont(void* self) {
	auto&& __r = ((QFontComboBox*)self)->currentFont();
	return new QFont(__r);
}
void QFontComboBox_setSampleTextForSystem(void* self, int writingSystem, const char* sampleText, long long sampleText_len) {
	QString sampleText_qstr = QString::fromUtf8(sampleText, sampleText_len);
	((QFontComboBox*)self)->setSampleTextForSystem((QFontDatabase::WritingSystem)(writingSystem), sampleText_qstr);
}
char* QFontComboBox_sampleTextForSystem(void* self, int writingSystem, long long* __out_len) {
	auto&& __r = ((QFontComboBox*)self)->sampleTextForSystem((QFontDatabase::WritingSystem)(writingSystem));
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QFontComboBox_setSampleTextForFont(void* self, const char* fontFamily, long long fontFamily_len, const char* sampleText, long long sampleText_len) {
	QString fontFamily_qstr = QString::fromUtf8(fontFamily, fontFamily_len);
	QString sampleText_qstr = QString::fromUtf8(sampleText, sampleText_len);
	((QFontComboBox*)self)->setSampleTextForFont(fontFamily_qstr, sampleText_qstr);
}
char* QFontComboBox_sampleTextForFont(void* self, const char* fontFamily, long long fontFamily_len, long long* __out_len) {
	QString fontFamily_qstr = QString::fromUtf8(fontFamily, fontFamily_len);
	auto&& __r = ((QFontComboBox*)self)->sampleTextForFont(fontFamily_qstr);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QFontComboBox_setDisplayFont(void* self, const char* fontFamily, long long fontFamily_len, void* font) {
	QString fontFamily_qstr = QString::fromUtf8(fontFamily, fontFamily_len);
	((QFontComboBox*)self)->setDisplayFont(fontFamily_qstr, *(QFont*)font);
}
void* QFontComboBox_displayFont(void* self, const char* fontFamily, long long fontFamily_len, int* __ok) {
	QString fontFamily_qstr = QString::fromUtf8(fontFamily, fontFamily_len);
	auto&& __r = ((QFontComboBox*)self)->displayFont(fontFamily_qstr);
	*__ok = __r.has_value() ? 1 : 0;
	return __r.has_value() ? new QFont(*__r) : (void*)0;
}
void QFontComboBox_setCurrentFont(void* self, void* f) {
	((QFontComboBox*)self)->setCurrentFont(*(QFont*)f);
}
void QFontComboBox_currentFontChanged(void* self, void* f) {
	((QFontComboBox*)self)->currentFontChanged(*(QFont*)f);
}
void QFontComboBox_setTabOrder(void* arg1, void* arg2) {
	QFontComboBox::setTabOrder((QWidget*)arg1, (QWidget*)arg2);
}
void* QFontComboBox_mouseGrabber() {
	auto&& __r = QFontComboBox::mouseGrabber();
	return (void*)(__r);
}
void* QFontComboBox_keyboardGrabber() {
	auto&& __r = QFontComboBox::keyboardGrabber();
	return (void*)(__r);
}
void* QFontComboBox_find(unsigned long long arg1) {
	auto&& __r = QFontComboBox::find(arg1);
	return (void*)(__r);
}
void* QFontComboBox_createWindowContainer(void* window) {
	auto&& __r = QFontComboBox::createWindowContainer((QWindow*)window);
	return (void*)(__r);
}
void QFontComboBox_updateMicroFocus(void* self) {
	((QFontComboBox__Go*)self)->updateMicroFocus();
}
void QFontComboBox_create(void* self) {
	((QFontComboBox__Go*)self)->create();
}
void QFontComboBox_destroy(void* self) {
	((QFontComboBox__Go*)self)->destroy();
}
int QFontComboBox_focusNextChild(void* self) {
	auto&& __r = ((QFontComboBox__Go*)self)->focusNextChild();
	return __r;
}
int QFontComboBox_focusPreviousChild(void* self) {
	auto&& __r = ((QFontComboBox__Go*)self)->focusPreviousChild();
	return __r;
}
void* QFontComboBox_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QFontComboBox::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QFontComboBox_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QFontComboBox::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QFontComboBox_sender(void* self) {
	auto&& __r = ((QFontComboBox__Go*)self)->sender();
	return (void*)(__r);
}
int QFontComboBox_senderSignalIndex(void* self) {
	auto&& __r = ((QFontComboBox__Go*)self)->senderSignalIndex();
	return __r;
}
int QFontComboBox_receivers(void* self, const char* signal) {
	auto&& __r = ((QFontComboBox__Go*)self)->receivers(signal);
	return __r;
}
int QFontComboBox_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QFontComboBox__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
double QFontComboBox_devicePixelRatioFScale() {
	auto&& __r = QFontComboBox::devicePixelRatioFScale();
	return __r;
}
int QFontComboBox_encodeMetricF(int metric, double value) {
	auto&& __r = QFontComboBox::encodeMetricF((QPaintDevice::PaintDeviceMetric)(metric), value);
	return __r;
}
double QFontComboBox_getDecodedMetricF(void* self, int metricA, int metricB) {
	auto&& __r = ((QFontComboBox__Go*)self)->getDecodedMetricF((QPaintDevice::PaintDeviceMetric)(metricA), (QPaintDevice::PaintDeviceMetric)(metricB));
	return __r;
}
char* QFontComboBox_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QFontComboBox::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QFontComboBox_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QFontComboBox::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QFontComboBox_createWindowContainer2(void* window, void* parent) {
	auto&& __r = QFontComboBox::createWindowContainer((QWindow*)window, (QWidget*)parent);
	return (void*)(__r);
}
void* QFontComboBox_createWindowContainer3(void* window, void* parent, int flags) {
	auto&& __r = QFontComboBox::createWindowContainer((QWindow*)window, (QWidget*)parent, QFlags<Qt::WindowType>(flags));
	return (void*)(__r);
}
void QFontComboBox_updateMicroFocus2(void* self, int query) {
	((QFontComboBox__Go*)self)->updateMicroFocus((Qt::InputMethodQuery)(query));
}
void QFontComboBox_create2(void* self, unsigned long long arg1) {
	((QFontComboBox__Go*)self)->create(arg1);
}
void QFontComboBox_create3(void* self, unsigned long long arg1, int initializeWindow) {
	((QFontComboBox__Go*)self)->create(arg1, initializeWindow);
}
void QFontComboBox_create4(void* self, unsigned long long arg1, int initializeWindow, int destroyOldWindow) {
	((QFontComboBox__Go*)self)->create(arg1, initializeWindow, destroyOldWindow);
}
void QFontComboBox_destroy2(void* self, int destroyWindow) {
	((QFontComboBox__Go*)self)->destroy(destroyWindow);
}
void QFontComboBox_destroy3(void* self, int destroyWindow, int destroySubWindows) {
	((QFontComboBox__Go*)self)->destroy(destroyWindow, destroySubWindows);
}
void* QFontComboBox_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QFontComboBox::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QFontComboBox_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QFontComboBox::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QFontComboBox_metaObject_default(void* self) {
	auto&& __r = ((QFontComboBox*)self)->QFontComboBox::metaObject();
	return (void*)(__r);
}
void* QFontComboBox_sizeHint_default(void* self) {
	auto&& __r = ((QFontComboBox*)self)->QFontComboBox::sizeHint();
	return new QSize(__r);
}
int QFontComboBox_event_default(void* self, void* e) {
	auto&& __r = ((QFontComboBox__Go*)self)->cppgoBase_event((QEvent*)e);
	return __r;
}
void QFontComboBox_setModel_default(void* self, void* model) {
	((QFontComboBox*)self)->QFontComboBox::setModel((QAbstractItemModel*)model);
}
void* QFontComboBox_minimumSizeHint_default(void* self) {
	auto&& __r = ((QFontComboBox*)self)->QFontComboBox::minimumSizeHint();
	return new QSize(__r);
}
void QFontComboBox_showPopup_default(void* self) {
	((QFontComboBox*)self)->QFontComboBox::showPopup();
}
void QFontComboBox_hidePopup_default(void* self) {
	((QFontComboBox*)self)->QFontComboBox::hidePopup();
}
void* QFontComboBox_inputMethodQuery_default(void* self, int arg1) {
	auto&& __r = ((QFontComboBox*)self)->QFontComboBox::inputMethodQuery((Qt::InputMethodQuery)(arg1));
	return new QVariant(__r);
}
void QFontComboBox_focusInEvent_default(void* self, void* e) {
	((QFontComboBox__Go*)self)->cppgoBase_focusInEvent((QFocusEvent*)e);
}
void QFontComboBox_focusOutEvent_default(void* self, void* e) {
	((QFontComboBox__Go*)self)->cppgoBase_focusOutEvent((QFocusEvent*)e);
}
void QFontComboBox_changeEvent_default(void* self, void* e) {
	((QFontComboBox__Go*)self)->cppgoBase_changeEvent((QEvent*)e);
}
void QFontComboBox_resizeEvent_default(void* self, void* e) {
	((QFontComboBox__Go*)self)->cppgoBase_resizeEvent((QResizeEvent*)e);
}
void QFontComboBox_paintEvent_default(void* self, void* e) {
	((QFontComboBox__Go*)self)->cppgoBase_paintEvent((QPaintEvent*)e);
}
void QFontComboBox_showEvent_default(void* self, void* e) {
	((QFontComboBox__Go*)self)->cppgoBase_showEvent((QShowEvent*)e);
}
void QFontComboBox_hideEvent_default(void* self, void* e) {
	((QFontComboBox__Go*)self)->cppgoBase_hideEvent((QHideEvent*)e);
}
void QFontComboBox_mousePressEvent_default(void* self, void* e) {
	((QFontComboBox__Go*)self)->cppgoBase_mousePressEvent((QMouseEvent*)e);
}
void QFontComboBox_mouseReleaseEvent_default(void* self, void* e) {
	((QFontComboBox__Go*)self)->cppgoBase_mouseReleaseEvent((QMouseEvent*)e);
}
void QFontComboBox_keyPressEvent_default(void* self, void* e) {
	((QFontComboBox__Go*)self)->cppgoBase_keyPressEvent((QKeyEvent*)e);
}
void QFontComboBox_keyReleaseEvent_default(void* self, void* e) {
	((QFontComboBox__Go*)self)->cppgoBase_keyReleaseEvent((QKeyEvent*)e);
}
void QFontComboBox_wheelEvent_default(void* self, void* e) {
	((QFontComboBox__Go*)self)->cppgoBase_wheelEvent((QWheelEvent*)e);
}
void QFontComboBox_contextMenuEvent_default(void* self, void* e) {
	((QFontComboBox__Go*)self)->cppgoBase_contextMenuEvent((QContextMenuEvent*)e);
}
void QFontComboBox_inputMethodEvent_default(void* self, void* arg1) {
	((QFontComboBox__Go*)self)->cppgoBase_inputMethodEvent((QInputMethodEvent*)arg1);
}
void QFontComboBox_initStyleOption_default(void* self, void* option) {
	((QFontComboBox__Go*)self)->cppgoBase_initStyleOption((QStyleOptionComboBox*)option);
}
int QFontComboBox_devType_default(void* self) {
	auto&& __r = ((QFontComboBox*)self)->QFontComboBox::devType();
	return __r;
}
void QFontComboBox_setVisible_default(void* self, int visible) {
	((QFontComboBox*)self)->QFontComboBox::setVisible(visible);
}
int QFontComboBox_heightForWidth_default(void* self, int arg1) {
	auto&& __r = ((QFontComboBox*)self)->QFontComboBox::heightForWidth(arg1);
	return __r;
}
int QFontComboBox_hasHeightForWidth_default(void* self) {
	auto&& __r = ((QFontComboBox*)self)->QFontComboBox::hasHeightForWidth();
	return __r;
}
void* QFontComboBox_paintEngine_default(void* self) {
	auto&& __r = ((QFontComboBox*)self)->QFontComboBox::paintEngine();
	return (void*)(__r);
}
void QFontComboBox_mouseDoubleClickEvent_default(void* self, void* event) {
	((QFontComboBox__Go*)self)->cppgoBase_mouseDoubleClickEvent((QMouseEvent*)event);
}
void QFontComboBox_mouseMoveEvent_default(void* self, void* event) {
	((QFontComboBox__Go*)self)->cppgoBase_mouseMoveEvent((QMouseEvent*)event);
}
void QFontComboBox_enterEvent_default(void* self, void* event) {
	((QFontComboBox__Go*)self)->cppgoBase_enterEvent((QEnterEvent*)event);
}
void QFontComboBox_leaveEvent_default(void* self, void* event) {
	((QFontComboBox__Go*)self)->cppgoBase_leaveEvent((QEvent*)event);
}
void QFontComboBox_moveEvent_default(void* self, void* event) {
	((QFontComboBox__Go*)self)->cppgoBase_moveEvent((QMoveEvent*)event);
}
void QFontComboBox_closeEvent_default(void* self, void* event) {
	((QFontComboBox__Go*)self)->cppgoBase_closeEvent((QCloseEvent*)event);
}
void QFontComboBox_tabletEvent_default(void* self, void* event) {
	((QFontComboBox__Go*)self)->cppgoBase_tabletEvent((QTabletEvent*)event);
}
void QFontComboBox_actionEvent_default(void* self, void* event) {
	((QFontComboBox__Go*)self)->cppgoBase_actionEvent((QActionEvent*)event);
}
void QFontComboBox_dragEnterEvent_default(void* self, void* event) {
	((QFontComboBox__Go*)self)->cppgoBase_dragEnterEvent((QDragEnterEvent*)event);
}
void QFontComboBox_dragMoveEvent_default(void* self, void* event) {
	((QFontComboBox__Go*)self)->cppgoBase_dragMoveEvent((QDragMoveEvent*)event);
}
void QFontComboBox_dragLeaveEvent_default(void* self, void* event) {
	((QFontComboBox__Go*)self)->cppgoBase_dragLeaveEvent((QDragLeaveEvent*)event);
}
void QFontComboBox_dropEvent_default(void* self, void* event) {
	((QFontComboBox__Go*)self)->cppgoBase_dropEvent((QDropEvent*)event);
}
int QFontComboBox_metric_default(void* self, int arg1) {
	auto&& __r = ((QFontComboBox__Go*)self)->cppgoBase_metric(arg1);
	return __r;
}
void QFontComboBox_initPainter_default(void* self, void* painter) {
	((QFontComboBox__Go*)self)->cppgoBase_initPainter((QPainter*)painter);
}
void* QFontComboBox_redirected_default(void* self, void* offset) {
	auto&& __r = ((QFontComboBox__Go*)self)->cppgoBase_redirected((QPoint*)offset);
	return (void*)(__r);
}
void* QFontComboBox_sharedPainter_default(void* self) {
	auto&& __r = ((QFontComboBox__Go*)self)->cppgoBase_sharedPainter();
	return (void*)(__r);
}
int QFontComboBox_focusNextPrevChild_default(void* self, int next) {
	auto&& __r = ((QFontComboBox__Go*)self)->cppgoBase_focusNextPrevChild(next);
	return __r;
}
int QFontComboBox_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QFontComboBox*)self)->QFontComboBox::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QFontComboBox_timerEvent_default(void* self, void* event) {
	((QFontComboBox__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QFontComboBox_childEvent_default(void* self, void* event) {
	((QFontComboBox__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QFontComboBox_customEvent_default(void* self, void* event) {
	((QFontComboBox__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QFontComboBox_connectNotify_default(void* self, void* signal) {
	((QFontComboBox__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QFontComboBox_disconnectNotify_default(void* self, void* signal) {
	((QFontComboBox__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QFontComboBox_connect_currentFontChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFontComboBox*)self, &QFontComboBox::currentFontChanged, (QFontComboBox*)self, [__id](const QFont& f){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&f);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFontComboBox);
	__conns_QFontComboBox[__id] = __c;
}
void QFontComboBox_connect_editTextChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFontComboBox*)self, &QFontComboBox::editTextChanged, (QFontComboBox*)self, [__id](const QString& arg1){
		QByteArray arg1_ba = arg1.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)arg1_ba.constData(), arg1_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFontComboBox);
	__conns_QFontComboBox[__id] = __c;
}
void QFontComboBox_connect_activated(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFontComboBox*)self, &QFontComboBox::activated, (QFontComboBox*)self, [__id](int index){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, index);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFontComboBox);
	__conns_QFontComboBox[__id] = __c;
}
void QFontComboBox_connect_textActivated(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFontComboBox*)self, &QFontComboBox::textActivated, (QFontComboBox*)self, [__id](const QString& arg1){
		QByteArray arg1_ba = arg1.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)arg1_ba.constData(), arg1_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFontComboBox);
	__conns_QFontComboBox[__id] = __c;
}
void QFontComboBox_connect_highlighted(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFontComboBox*)self, &QFontComboBox::highlighted, (QFontComboBox*)self, [__id](int index){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, index);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFontComboBox);
	__conns_QFontComboBox[__id] = __c;
}
void QFontComboBox_connect_textHighlighted(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFontComboBox*)self, &QFontComboBox::textHighlighted, (QFontComboBox*)self, [__id](const QString& arg1){
		QByteArray arg1_ba = arg1.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)arg1_ba.constData(), arg1_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFontComboBox);
	__conns_QFontComboBox[__id] = __c;
}
void QFontComboBox_connect_currentIndexChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFontComboBox*)self, &QFontComboBox::currentIndexChanged, (QFontComboBox*)self, [__id](int index){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, index);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFontComboBox);
	__conns_QFontComboBox[__id] = __c;
}
void QFontComboBox_connect_currentTextChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFontComboBox*)self, &QFontComboBox::currentTextChanged, (QFontComboBox*)self, [__id](const QString& arg1){
		QByteArray arg1_ba = arg1.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)arg1_ba.constData(), arg1_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFontComboBox);
	__conns_QFontComboBox[__id] = __c;
}
void QFontComboBox_connect_windowTitleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFontComboBox*)self, &QFontComboBox::windowTitleChanged, (QFontComboBox*)self, [__id](const QString& title){
		QByteArray title_ba = title.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)title_ba.constData(), title_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFontComboBox);
	__conns_QFontComboBox[__id] = __c;
}
void QFontComboBox_connect_windowIconChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFontComboBox*)self, &QFontComboBox::windowIconChanged, (QFontComboBox*)self, [__id](const QIcon& icon){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&icon);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFontComboBox);
	__conns_QFontComboBox[__id] = __c;
}
void QFontComboBox_connect_windowIconTextChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFontComboBox*)self, &QFontComboBox::windowIconTextChanged, (QFontComboBox*)self, [__id](const QString& iconText){
		QByteArray iconText_ba = iconText.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)iconText_ba.constData(), iconText_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFontComboBox);
	__conns_QFontComboBox[__id] = __c;
}
void QFontComboBox_connect_customContextMenuRequested(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFontComboBox*)self, &QFontComboBox::customContextMenuRequested, (QFontComboBox*)self, [__id](const QPoint& pos){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&pos);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFontComboBox);
	__conns_QFontComboBox[__id] = __c;
}
void QFontComboBox_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFontComboBox*)self, &QFontComboBox::destroyed, (QFontComboBox*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFontComboBox);
	__conns_QFontComboBox[__id] = __c;
}
void QFontComboBox_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFontComboBox*)self, &QFontComboBox::destroyed, (QFontComboBox*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFontComboBox);
	__conns_QFontComboBox[__id] = __c;
}
void QFontComboBox_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFontComboBox*)self, &QFontComboBox::objectNameChanged, (QFontComboBox*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFontComboBox);
	__conns_QFontComboBox[__id] = __c;
}
void QFontComboBox__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QFontComboBox);
		auto __it = __conns_QFontComboBox.find(id);
		if (__it == __conns_QFontComboBox.end()) return;
		__c = __it->second;
		__conns_QFontComboBox.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QFontComboBox__Go_new(long long id) { return new QFontComboBox__Go(id); }
void QFontComboBox_delete(void* self) { delete (QFontComboBox*)self; }

}

#include <QtWidgets/qgraphicslayoutitem.h>
#include <QtWidgets/qsizepolicy.h>
#include <QtCore/qsize.h>
#include <QtCore/qrect.h>
#include <QtCore/qnamespace.h>
#include <QtWidgets/qgraphicsitem.h>
#include "gen_qgraphicslayoutitem.h"
#include "gen_purego_cb.h"

class QGraphicsLayoutItem__Go : public QGraphicsLayoutItem {
public:
	long long __goId;
	unsigned char __goOn[5] = {};
	template<typename... A> QGraphicsLayoutItem__Go(long long id, A&&... a) : QGraphicsLayoutItem(std::forward<A>(a)...), __goId(id) {}
	using QGraphicsLayoutItem::setGraphicsItem;
	using QGraphicsLayoutItem::setOwnedByLayout;
	QSizeF sizeHint(Qt::SizeHint which, const QSizeF& constraint) const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp_int_voidp) { void* __r = ((void*(*)(long long, long long, int, void*, int*))cppgo_virt_voidp_int_voidp)(__goId, 729690261457684101LL, (int)(which), (void*)&constraint, &__h); if (__h) return *(QSizeF*)__r; }
		return QSizeF{};
	}
	void setGeometry(const QRectF& rect) override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5684820848726969227LL, (void*)&rect, &__h); if (__h) return; }
		QGraphicsLayoutItem::setGeometry(rect);
	}
	void getContentsMargins(qreal* left, qreal* top, qreal* right, qreal* bottom) const override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_void_doublep_doublep_doublep_doublep) { ((void(*)(long long, long long, double*, double*, double*, double*, int*))cppgo_virt_void_doublep_doublep_doublep_doublep)(__goId, 164760286202646972LL, (double*)left, (double*)top, (double*)right, (double*)bottom, &__h); if (__h) return; }
		QGraphicsLayoutItem::getContentsMargins(left, top, right, bottom);
	}
	void updateGeometry() override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 8082455443689927668LL, &__h); if (__h) return; }
		QGraphicsLayoutItem::updateGeometry();
	}
	bool isEmpty() const override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -626822430251098498LL, &__h); if (__h) return __r; }
		return QGraphicsLayoutItem::isEmpty();
	}
};

extern "C" {

void QGraphicsLayoutItem__set_override(void* self, int idx) { ((QGraphicsLayoutItem__Go*)self)->__goOn[idx] = 1; }
void QGraphicsLayoutItem_setSizePolicy(void* self, void* policy) {
	((QGraphicsLayoutItem*)self)->setSizePolicy(*(QSizePolicy*)policy);
}
void QGraphicsLayoutItem_setSizePolicy2(void* self, int hPolicy, int vPolicy) {
	((QGraphicsLayoutItem*)self)->setSizePolicy((QSizePolicy::Policy)(hPolicy), (QSizePolicy::Policy)(vPolicy));
}
void* QGraphicsLayoutItem_sizePolicy(void* self) {
	auto&& __r = ((QGraphicsLayoutItem*)self)->sizePolicy();
	return new QSizePolicy(__r);
}
void QGraphicsLayoutItem_setMinimumSize(void* self, void* size) {
	((QGraphicsLayoutItem*)self)->setMinimumSize(*(QSizeF*)size);
}
void QGraphicsLayoutItem_setMinimumSize2(void* self, double w, double h) {
	((QGraphicsLayoutItem*)self)->setMinimumSize(w, h);
}
void* QGraphicsLayoutItem_minimumSize(void* self) {
	auto&& __r = ((QGraphicsLayoutItem*)self)->minimumSize();
	return new QSizeF(__r);
}
void QGraphicsLayoutItem_setMinimumWidth(void* self, double width) {
	((QGraphicsLayoutItem*)self)->setMinimumWidth(width);
}
double QGraphicsLayoutItem_minimumWidth(void* self) {
	auto&& __r = ((QGraphicsLayoutItem*)self)->minimumWidth();
	return __r;
}
void QGraphicsLayoutItem_setMinimumHeight(void* self, double height) {
	((QGraphicsLayoutItem*)self)->setMinimumHeight(height);
}
double QGraphicsLayoutItem_minimumHeight(void* self) {
	auto&& __r = ((QGraphicsLayoutItem*)self)->minimumHeight();
	return __r;
}
void QGraphicsLayoutItem_setPreferredSize(void* self, void* size) {
	((QGraphicsLayoutItem*)self)->setPreferredSize(*(QSizeF*)size);
}
void QGraphicsLayoutItem_setPreferredSize2(void* self, double w, double h) {
	((QGraphicsLayoutItem*)self)->setPreferredSize(w, h);
}
void* QGraphicsLayoutItem_preferredSize(void* self) {
	auto&& __r = ((QGraphicsLayoutItem*)self)->preferredSize();
	return new QSizeF(__r);
}
void QGraphicsLayoutItem_setPreferredWidth(void* self, double width) {
	((QGraphicsLayoutItem*)self)->setPreferredWidth(width);
}
double QGraphicsLayoutItem_preferredWidth(void* self) {
	auto&& __r = ((QGraphicsLayoutItem*)self)->preferredWidth();
	return __r;
}
void QGraphicsLayoutItem_setPreferredHeight(void* self, double height) {
	((QGraphicsLayoutItem*)self)->setPreferredHeight(height);
}
double QGraphicsLayoutItem_preferredHeight(void* self) {
	auto&& __r = ((QGraphicsLayoutItem*)self)->preferredHeight();
	return __r;
}
void QGraphicsLayoutItem_setMaximumSize(void* self, void* size) {
	((QGraphicsLayoutItem*)self)->setMaximumSize(*(QSizeF*)size);
}
void QGraphicsLayoutItem_setMaximumSize2(void* self, double w, double h) {
	((QGraphicsLayoutItem*)self)->setMaximumSize(w, h);
}
void* QGraphicsLayoutItem_maximumSize(void* self) {
	auto&& __r = ((QGraphicsLayoutItem*)self)->maximumSize();
	return new QSizeF(__r);
}
void QGraphicsLayoutItem_setMaximumWidth(void* self, double width) {
	((QGraphicsLayoutItem*)self)->setMaximumWidth(width);
}
double QGraphicsLayoutItem_maximumWidth(void* self) {
	auto&& __r = ((QGraphicsLayoutItem*)self)->maximumWidth();
	return __r;
}
void QGraphicsLayoutItem_setMaximumHeight(void* self, double height) {
	((QGraphicsLayoutItem*)self)->setMaximumHeight(height);
}
double QGraphicsLayoutItem_maximumHeight(void* self) {
	auto&& __r = ((QGraphicsLayoutItem*)self)->maximumHeight();
	return __r;
}
void QGraphicsLayoutItem_setGeometry(void* self, void* rect) {
	((QGraphicsLayoutItem*)self)->setGeometry(*(QRectF*)rect);
}
void* QGraphicsLayoutItem_geometry(void* self) {
	auto&& __r = ((QGraphicsLayoutItem*)self)->geometry();
	return new QRectF(__r);
}
void QGraphicsLayoutItem_getContentsMargins(void* self, double* left, double* top, double* right, double* bottom) {
	qreal left_local{};
	qreal top_local{};
	qreal right_local{};
	qreal bottom_local{};
	((QGraphicsLayoutItem*)self)->getContentsMargins(&left_local, &top_local, &right_local, &bottom_local);
	*left = left_local;
	*top = top_local;
	*right = right_local;
	*bottom = bottom_local;
}
void* QGraphicsLayoutItem_contentsRect(void* self) {
	auto&& __r = ((QGraphicsLayoutItem*)self)->contentsRect();
	return new QRectF(__r);
}
void* QGraphicsLayoutItem_effectiveSizeHint(void* self, int which) {
	auto&& __r = ((QGraphicsLayoutItem*)self)->effectiveSizeHint((Qt::SizeHint)(which));
	return new QSizeF(__r);
}
void QGraphicsLayoutItem_updateGeometry(void* self) {
	((QGraphicsLayoutItem*)self)->updateGeometry();
}
int QGraphicsLayoutItem_isEmpty(void* self) {
	auto&& __r = ((QGraphicsLayoutItem*)self)->isEmpty();
	return __r;
}
void* QGraphicsLayoutItem_parentLayoutItem(void* self) {
	auto&& __r = ((QGraphicsLayoutItem*)self)->parentLayoutItem();
	return (void*)(__r);
}
void QGraphicsLayoutItem_setParentLayoutItem(void* self, void* parent) {
	((QGraphicsLayoutItem*)self)->setParentLayoutItem((QGraphicsLayoutItem*)parent);
}
int QGraphicsLayoutItem_isLayout(void* self) {
	auto&& __r = ((QGraphicsLayoutItem*)self)->isLayout();
	return __r;
}
void* QGraphicsLayoutItem_graphicsItem(void* self) {
	auto&& __r = ((QGraphicsLayoutItem*)self)->graphicsItem();
	return (void*)(__r);
}
int QGraphicsLayoutItem_ownedByLayout(void* self) {
	auto&& __r = ((QGraphicsLayoutItem*)self)->ownedByLayout();
	return __r;
}
void QGraphicsLayoutItem_setGraphicsItem(void* self, void* item) {
	((QGraphicsLayoutItem__Go*)self)->setGraphicsItem((QGraphicsItem*)item);
}
void QGraphicsLayoutItem_setOwnedByLayout(void* self, int ownedByLayout) {
	((QGraphicsLayoutItem__Go*)self)->setOwnedByLayout(ownedByLayout);
}
void QGraphicsLayoutItem_setSizePolicy3(void* self, int hPolicy, int vPolicy, int controlType) {
	((QGraphicsLayoutItem*)self)->setSizePolicy((QSizePolicy::Policy)(hPolicy), (QSizePolicy::Policy)(vPolicy), (QSizePolicy::ControlType)(controlType));
}
void* QGraphicsLayoutItem_effectiveSizeHint2(void* self, int which, void* constraint) {
	auto&& __r = ((QGraphicsLayoutItem*)self)->effectiveSizeHint((Qt::SizeHint)(which), *(QSizeF*)constraint);
	return new QSizeF(__r);
}
void QGraphicsLayoutItem_setGeometry_default(void* self, void* rect) {
	((QGraphicsLayoutItem*)self)->QGraphicsLayoutItem::setGeometry(*(QRectF*)rect);
}
void QGraphicsLayoutItem_getContentsMargins_default(void* self, double* left, double* top, double* right, double* bottom) {
	qreal left_local{};
	qreal top_local{};
	qreal right_local{};
	qreal bottom_local{};
	((QGraphicsLayoutItem*)self)->QGraphicsLayoutItem::getContentsMargins(&left_local, &top_local, &right_local, &bottom_local);
	*left = left_local;
	*top = top_local;
	*right = right_local;
	*bottom = bottom_local;
}
void QGraphicsLayoutItem_updateGeometry_default(void* self) {
	((QGraphicsLayoutItem*)self)->QGraphicsLayoutItem::updateGeometry();
}
int QGraphicsLayoutItem_isEmpty_default(void* self) {
	auto&& __r = ((QGraphicsLayoutItem*)self)->QGraphicsLayoutItem::isEmpty();
	return __r;
}
void* QGraphicsLayoutItem__Go_new(long long id) { return new QGraphicsLayoutItem__Go(id); }
void QGraphicsLayoutItem_delete(void* self) { delete (QGraphicsLayoutItem*)self; }

}

#include <QtWidgets/qgraphicstransform.h>
#include <QtCore/qobject.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtGui/qmatrix4x4.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qvariant.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qgraphicstransform.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QGraphicsTransform__Go : public QGraphicsTransform {
public:
	long long __goId;
	unsigned char __goOn[9] = {};
	template<typename... A> QGraphicsTransform__Go(long long id, A&&... a) : QGraphicsTransform(std::forward<A>(a)...), __goId(id) {}
	using QGraphicsTransform::update;
	using QGraphicsTransform::sender;
	using QGraphicsTransform::senderSignalIndex;
	using QGraphicsTransform::receivers;
	using QGraphicsTransform::isSignalConnected;
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QGraphicsTransform::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QGraphicsTransform::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QGraphicsTransform::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QGraphicsTransform::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QGraphicsTransform::disconnectNotify(__a0); }
	void applyTo(QMatrix4x4* matrix) const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1580758529722784438LL, (void*)matrix, &__h); if (__h) return; }
		return;
	}
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QGraphicsTransform::metaObject();
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QGraphicsTransform::event(event);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QGraphicsTransform::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QGraphicsTransform::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QGraphicsTransform::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QGraphicsTransform::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QGraphicsTransform::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QGraphicsTransform::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QGraphicsTransform;
static std::mutex __connsMu_QGraphicsTransform;

extern "C" {

void* QGraphicsTransform_base(void* self) { return (void*)static_cast<QObject*>((QGraphicsTransform*)self); }
void QGraphicsTransform__set_override(void* self, int idx) { ((QGraphicsTransform__Go*)self)->__goOn[idx] = 1; }
void* QGraphicsTransform_fromQObject(void* o) { return (void*)qobject_cast<QGraphicsTransform*>((QObject*)o); }
char* QGraphicsTransform_tr(const char* s, long long* __out_len) {
	auto&& __r = QGraphicsTransform::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QGraphicsTransform_applyTo(void* self, void* matrix) {
	((QGraphicsTransform*)self)->applyTo((QMatrix4x4*)matrix);
}
void QGraphicsTransform_update(void* self) {
	((QGraphicsTransform__Go*)self)->update();
}
void* QGraphicsTransform_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QGraphicsTransform::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QGraphicsTransform_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QGraphicsTransform::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QGraphicsTransform_sender(void* self) {
	auto&& __r = ((QGraphicsTransform__Go*)self)->sender();
	return (void*)(__r);
}
int QGraphicsTransform_senderSignalIndex(void* self) {
	auto&& __r = ((QGraphicsTransform__Go*)self)->senderSignalIndex();
	return __r;
}
int QGraphicsTransform_receivers(void* self, const char* signal) {
	auto&& __r = ((QGraphicsTransform__Go*)self)->receivers(signal);
	return __r;
}
int QGraphicsTransform_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QGraphicsTransform__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QGraphicsTransform_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QGraphicsTransform::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QGraphicsTransform_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QGraphicsTransform::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QGraphicsTransform_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QGraphicsTransform::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QGraphicsTransform_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QGraphicsTransform::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QGraphicsTransform_metaObject_default(void* self) {
	auto&& __r = ((QGraphicsTransform*)self)->QGraphicsTransform::metaObject();
	return (void*)(__r);
}
int QGraphicsTransform_event_default(void* self, void* event) {
	auto&& __r = ((QGraphicsTransform*)self)->QGraphicsTransform::event((QEvent*)event);
	return __r;
}
int QGraphicsTransform_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QGraphicsTransform*)self)->QGraphicsTransform::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QGraphicsTransform_timerEvent_default(void* self, void* event) {
	((QGraphicsTransform__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QGraphicsTransform_childEvent_default(void* self, void* event) {
	((QGraphicsTransform__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QGraphicsTransform_customEvent_default(void* self, void* event) {
	((QGraphicsTransform__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QGraphicsTransform_connectNotify_default(void* self, void* signal) {
	((QGraphicsTransform__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QGraphicsTransform_disconnectNotify_default(void* self, void* signal) {
	((QGraphicsTransform__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QGraphicsTransform_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsTransform*)self, &QGraphicsTransform::destroyed, (QGraphicsTransform*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsTransform);
	__conns_QGraphicsTransform[__id] = __c;
}
void QGraphicsTransform_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsTransform*)self, &QGraphicsTransform::destroyed, (QGraphicsTransform*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsTransform);
	__conns_QGraphicsTransform[__id] = __c;
}
void QGraphicsTransform_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsTransform*)self, &QGraphicsTransform::objectNameChanged, (QGraphicsTransform*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsTransform);
	__conns_QGraphicsTransform[__id] = __c;
}
void QGraphicsTransform__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsTransform);
		auto __it = __conns_QGraphicsTransform.find(id);
		if (__it == __conns_QGraphicsTransform.end()) return;
		__c = __it->second;
		__conns_QGraphicsTransform.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QGraphicsTransform__Go_new(long long id, void* parent) { return new QGraphicsTransform__Go(id, (QObject*)parent); }
void QGraphicsTransform_delete(void* self) { delete (QGraphicsTransform*)self; }

}

#include <QtWidgets/qtextedit.h>
#include <QtWidgets/qabstractscrollarea.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qstring.h>
#include <QtCore/qobjectdefs.h>
#include <QtGui/qtextdocument.h>
#include <QtGui/qtextcursor.h>
#include <QtCore/qnamespace.h>
#include <QtGui/qcolor.h>
#include <QtGui/qfont.h>
#include <QtGui/qtextformat.h>
#include <QtGui/qtextoption.h>
#include <QtCore/qregularexpression.h>
#include <QtCore/qvariant.h>
#include <QtCore/qurl.h>
#include <QtWidgets/qmenu.h>
#include <QtCore/qpoint.h>
#include <QtCore/qrect.h>
#include <QtGui/qpagedpaintdevice.h>
#include <QtCore/qcoreevent.h>
#include <QtGui/qevent.h>
#include <QtCore/qmimedata.h>
#include <QtWidgets/qscrollbar.h>
#include <QtCore/qsize.h>
#include <QtCore/qmargins.h>
#include <QtCore/qobject.h>
#include <QtWidgets/qframe.h>
#include <QtGui/qpainter.h>
#include <QtWidgets/qstyleoption.h>
#include <QtWidgets/qstyle.h>
#include <QtGui/qregion.h>
#include <QtGui/qpalette.h>
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
#include "gen_qtextedit.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QTextEdit__Go : public QTextEdit {
public:
	long long __goId;
	unsigned char __goOn[58] = {};
	template<typename... A> QTextEdit__Go(long long id, A&&... a) : QTextEdit(std::forward<A>(a)...), __goId(id) {}
	using QTextEdit::zoomInF;
	using QTextEdit::setViewportMargins;
	using QTextEdit::viewportMargins;
	using QTextEdit::drawFrame;
	using QTextEdit::updateMicroFocus;
	using QTextEdit::create;
	using QTextEdit::destroy;
	using QTextEdit::focusNextChild;
	using QTextEdit::focusPreviousChild;
	using QTextEdit::sender;
	using QTextEdit::senderSignalIndex;
	using QTextEdit::receivers;
	using QTextEdit::isSignalConnected;
	using QTextEdit::getDecodedMetricF;
	bool cppgoBase_event(QEvent* __a0) { return QTextEdit::event(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QTextEdit::timerEvent(__a0); }
	void cppgoBase_keyPressEvent(QKeyEvent* __a0) { QTextEdit::keyPressEvent(__a0); }
	void cppgoBase_keyReleaseEvent(QKeyEvent* __a0) { QTextEdit::keyReleaseEvent(__a0); }
	void cppgoBase_resizeEvent(QResizeEvent* __a0) { QTextEdit::resizeEvent(__a0); }
	void cppgoBase_paintEvent(QPaintEvent* __a0) { QTextEdit::paintEvent(__a0); }
	void cppgoBase_mousePressEvent(QMouseEvent* __a0) { QTextEdit::mousePressEvent(__a0); }
	void cppgoBase_mouseMoveEvent(QMouseEvent* __a0) { QTextEdit::mouseMoveEvent(__a0); }
	void cppgoBase_mouseReleaseEvent(QMouseEvent* __a0) { QTextEdit::mouseReleaseEvent(__a0); }
	void cppgoBase_mouseDoubleClickEvent(QMouseEvent* __a0) { QTextEdit::mouseDoubleClickEvent(__a0); }
	bool cppgoBase_focusNextPrevChild(bool __a0) { return QTextEdit::focusNextPrevChild(__a0); }
	void cppgoBase_contextMenuEvent(QContextMenuEvent* __a0) { QTextEdit::contextMenuEvent(__a0); }
	void cppgoBase_dragEnterEvent(QDragEnterEvent* __a0) { QTextEdit::dragEnterEvent(__a0); }
	void cppgoBase_dragLeaveEvent(QDragLeaveEvent* __a0) { QTextEdit::dragLeaveEvent(__a0); }
	void cppgoBase_dragMoveEvent(QDragMoveEvent* __a0) { QTextEdit::dragMoveEvent(__a0); }
	void cppgoBase_dropEvent(QDropEvent* __a0) { QTextEdit::dropEvent(__a0); }
	void cppgoBase_focusInEvent(QFocusEvent* __a0) { QTextEdit::focusInEvent(__a0); }
	void cppgoBase_focusOutEvent(QFocusEvent* __a0) { QTextEdit::focusOutEvent(__a0); }
	void cppgoBase_showEvent(QShowEvent* __a0) { QTextEdit::showEvent(__a0); }
	void cppgoBase_changeEvent(QEvent* __a0) { QTextEdit::changeEvent(__a0); }
	void cppgoBase_wheelEvent(QWheelEvent* __a0) { QTextEdit::wheelEvent(__a0); }
	QMimeData* cppgoBase_createMimeDataFromSelection() { return QTextEdit::createMimeDataFromSelection(); }
	bool cppgoBase_canInsertFromMimeData(const QMimeData* __a0) { return QTextEdit::canInsertFromMimeData(__a0); }
	void cppgoBase_insertFromMimeData(const QMimeData* __a0) { QTextEdit::insertFromMimeData(__a0); }
	void cppgoBase_inputMethodEvent(QInputMethodEvent* __a0) { QTextEdit::inputMethodEvent(__a0); }
	void cppgoBase_scrollContentsBy(int __a0, int __a1) { QTextEdit::scrollContentsBy(__a0, __a1); }
	void cppgoBase_doSetTextCursor(const QTextCursor& __a0) { QTextEdit::doSetTextCursor(__a0); }
	bool cppgoBase_eventFilter(QObject* __a0, QEvent* __a1) { return QTextEdit::eventFilter(__a0, __a1); }
	bool cppgoBase_viewportEvent(QEvent* __a0) { return QTextEdit::viewportEvent(__a0); }
	QSize cppgoBase_viewportSizeHint() { return QTextEdit::viewportSizeHint(); }
	void cppgoBase_initStyleOption(QStyleOptionFrame* __a0) { QTextEdit::initStyleOption(__a0); }
	void cppgoBase_enterEvent(QEnterEvent* __a0) { QTextEdit::enterEvent(__a0); }
	void cppgoBase_leaveEvent(QEvent* __a0) { QTextEdit::leaveEvent(__a0); }
	void cppgoBase_moveEvent(QMoveEvent* __a0) { QTextEdit::moveEvent(__a0); }
	void cppgoBase_closeEvent(QCloseEvent* __a0) { QTextEdit::closeEvent(__a0); }
	void cppgoBase_tabletEvent(QTabletEvent* __a0) { QTextEdit::tabletEvent(__a0); }
	void cppgoBase_actionEvent(QActionEvent* __a0) { QTextEdit::actionEvent(__a0); }
	void cppgoBase_hideEvent(QHideEvent* __a0) { QTextEdit::hideEvent(__a0); }
	int cppgoBase_metric(int __a0) { return QTextEdit::metric((QPaintDevice::PaintDeviceMetric)__a0); }
	void cppgoBase_initPainter(QPainter* __a0) { QTextEdit::initPainter(__a0); }
	QPaintDevice* cppgoBase_redirected(QPoint* __a0) { return QTextEdit::redirected(__a0); }
	QPainter* cppgoBase_sharedPainter() { return QTextEdit::sharedPainter(); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QTextEdit::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QTextEdit::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QTextEdit::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QTextEdit::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QTextEdit::metaObject();
	}
	QVariant loadResource(int type_, const QUrl& name) override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_voidp_int_voidp) { void* __r = ((void*(*)(long long, long long, int, void*, int*))cppgo_virt_voidp_int_voidp)(__goId, -519616681208050567LL, type_, (void*)&name, &__h); if (__h) return *(QVariant*)__r; }
		return QTextEdit::loadResource(type_, name);
	}
	QVariant inputMethodQuery(Qt::InputMethodQuery property) const override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7345853097340709638LL, (int)(property), &__h); if (__h) return *(QVariant*)__r; }
		return QTextEdit::inputMethodQuery(property);
	}
	bool event(QEvent* e) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)e, &__h); if (__h) return __r; }
		return QTextEdit::event(e);
	}
	void timerEvent(QTimerEvent* e) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)e, &__h); if (__h) return; }
		QTextEdit::timerEvent(e);
	}
	void keyPressEvent(QKeyEvent* e) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1023791977679793185LL, (void*)e, &__h); if (__h) return; }
		QTextEdit::keyPressEvent(e);
	}
	void keyReleaseEvent(QKeyEvent* e) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 772598475950050299LL, (void*)e, &__h); if (__h) return; }
		QTextEdit::keyReleaseEvent(e);
	}
	void resizeEvent(QResizeEvent* e) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4302917285655014661LL, (void*)e, &__h); if (__h) return; }
		QTextEdit::resizeEvent(e);
	}
	void paintEvent(QPaintEvent* e) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5290442163167458465LL, (void*)e, &__h); if (__h) return; }
		QTextEdit::paintEvent(e);
	}
	void mousePressEvent(QMouseEvent* e) override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -332090898078543319LL, (void*)e, &__h); if (__h) return; }
		QTextEdit::mousePressEvent(e);
	}
	void mouseMoveEvent(QMouseEvent* e) override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4255494004099806581LL, (void*)e, &__h); if (__h) return; }
		QTextEdit::mouseMoveEvent(e);
	}
	void mouseReleaseEvent(QMouseEvent* e) override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4415920049297795213LL, (void*)e, &__h); if (__h) return; }
		QTextEdit::mouseReleaseEvent(e);
	}
	void mouseDoubleClickEvent(QMouseEvent* e) override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4029073905697023887LL, (void*)e, &__h); if (__h) return; }
		QTextEdit::mouseDoubleClickEvent(e);
	}
	bool focusNextPrevChild(bool next) override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 5898328517505784191LL, next, &__h); if (__h) return __r; }
		return QTextEdit::focusNextPrevChild(next);
	}
	void contextMenuEvent(QContextMenuEvent* e) override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6698043222227006015LL, (void*)e, &__h); if (__h) return; }
		QTextEdit::contextMenuEvent(e);
	}
	void dragEnterEvent(QDragEnterEvent* e) override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 206122715919917589LL, (void*)e, &__h); if (__h) return; }
		QTextEdit::dragEnterEvent(e);
	}
	void dragLeaveEvent(QDragLeaveEvent* e) override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -3139904158373725006LL, (void*)e, &__h); if (__h) return; }
		QTextEdit::dragLeaveEvent(e);
	}
	void dragMoveEvent(QDragMoveEvent* e) override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1781545137352745978LL, (void*)e, &__h); if (__h) return; }
		QTextEdit::dragMoveEvent(e);
	}
	void dropEvent(QDropEvent* e) override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 9107058682754591716LL, (void*)e, &__h); if (__h) return; }
		QTextEdit::dropEvent(e);
	}
	void focusInEvent(QFocusEvent* e) override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 975043767033552552LL, (void*)e, &__h); if (__h) return; }
		QTextEdit::focusInEvent(e);
	}
	void focusOutEvent(QFocusEvent* e) override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -385596010576048993LL, (void*)e, &__h); if (__h) return; }
		QTextEdit::focusOutEvent(e);
	}
	void showEvent(QShowEvent* arg1) override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6825369174751916LL, (void*)arg1, &__h); if (__h) return; }
		QTextEdit::showEvent(arg1);
	}
	void changeEvent(QEvent* e) override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4959376063826165943LL, (void*)e, &__h); if (__h) return; }
		QTextEdit::changeEvent(e);
	}
	void wheelEvent(QWheelEvent* e) override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1919988340934511316LL, (void*)e, &__h); if (__h) return; }
		QTextEdit::wheelEvent(e);
	}
	QMimeData* createMimeDataFromSelection() const override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 3493558106800165317LL, &__h); if (__h) return (QMimeData*)__r; }
		return QTextEdit::createMimeDataFromSelection();
	}
	bool canInsertFromMimeData(const QMimeData* source) const override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, -2605205276815256896LL, (void*)source, &__h); if (__h) return __r; }
		return QTextEdit::canInsertFromMimeData(source);
	}
	void insertFromMimeData(const QMimeData* source) override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -466553387102125056LL, (void*)source, &__h); if (__h) return; }
		QTextEdit::insertFromMimeData(source);
	}
	void inputMethodEvent(QInputMethodEvent* arg1) override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5438819661942843222LL, (void*)arg1, &__h); if (__h) return; }
		QTextEdit::inputMethodEvent(arg1);
	}
	void scrollContentsBy(int dx, int dy) override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_void_int_int) { ((void(*)(long long, long long, int, int, int*))cppgo_virt_void_int_int)(__goId, -6118001150126564745LL, dx, dy, &__h); if (__h) return; }
		QTextEdit::scrollContentsBy(dx, dy);
	}
	void doSetTextCursor(const QTextCursor& cursor) override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4090661491944147209LL, (void*)&cursor, &__h); if (__h) return; }
		QTextEdit::doSetTextCursor(cursor);
	}
	QSize minimumSizeHint() const override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 7049264599837532491LL, &__h); if (__h) return *(QSize*)__r; }
		return QTextEdit::minimumSizeHint();
	}
	QSize sizeHint() const override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 729690261457684101LL, &__h); if (__h) return *(QSize*)__r; }
		return QTextEdit::sizeHint();
	}
	void setupViewport(QWidget* viewport) override {
		int __h = 0;
		if (__goOn[32] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4531996381736730688LL, (void*)viewport, &__h); if (__h) return; }
		QTextEdit::setupViewport(viewport);
	}
	bool eventFilter(QObject* arg1, QEvent* arg2) override {
		int __h = 0;
		if (__goOn[33] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)arg1, (void*)arg2, &__h); if (__h) return __r; }
		return QTextEdit::eventFilter(arg1, arg2);
	}
	bool viewportEvent(QEvent* arg1) override {
		int __h = 0;
		if (__goOn[34] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, -1971465653061861519LL, (void*)arg1, &__h); if (__h) return __r; }
		return QTextEdit::viewportEvent(arg1);
	}
	QSize viewportSizeHint() const override {
		int __h = 0;
		if (__goOn[35] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 503579067252048239LL, &__h); if (__h) return *(QSize*)__r; }
		return QTextEdit::viewportSizeHint();
	}
	void initStyleOption(QStyleOptionFrame* option) const override {
		int __h = 0;
		if (__goOn[36] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -488693769771308531LL, (void*)option, &__h); if (__h) return; }
		QTextEdit::initStyleOption(option);
	}
	int devType() const override {
		int __h = 0;
		if (__goOn[37] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8659679893868524044LL, &__h); if (__h) return __r; }
		return QTextEdit::devType();
	}
	void setVisible(bool visible) override {
		int __h = 0;
		if (__goOn[38] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 2110182719375245575LL, visible, &__h); if (__h) return; }
		QTextEdit::setVisible(visible);
	}
	int heightForWidth(int arg1) const override {
		int __h = 0;
		if (__goOn[39] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 4311534060495737339LL, arg1, &__h); if (__h) return __r; }
		return QTextEdit::heightForWidth(arg1);
	}
	bool hasHeightForWidth() const override {
		int __h = 0;
		if (__goOn[40] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8266762397995423915LL, &__h); if (__h) return __r; }
		return QTextEdit::hasHeightForWidth();
	}
	QPaintEngine* paintEngine() const override {
		int __h = 0;
		if (__goOn[41] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 6398007404756346381LL, &__h); if (__h) return (QPaintEngine*)__r; }
		return QTextEdit::paintEngine();
	}
	void enterEvent(QEnterEvent* event) override {
		int __h = 0;
		if (__goOn[42] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6316270346668237895LL, (void*)event, &__h); if (__h) return; }
		QTextEdit::enterEvent(event);
	}
	void leaveEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[43] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -551070009623711640LL, (void*)event, &__h); if (__h) return; }
		QTextEdit::leaveEvent(event);
	}
	void moveEvent(QMoveEvent* event) override {
		int __h = 0;
		if (__goOn[44] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -9161949150489525396LL, (void*)event, &__h); if (__h) return; }
		QTextEdit::moveEvent(event);
	}
	void closeEvent(QCloseEvent* event) override {
		int __h = 0;
		if (__goOn[45] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1989937050019464783LL, (void*)event, &__h); if (__h) return; }
		QTextEdit::closeEvent(event);
	}
	void tabletEvent(QTabletEvent* event) override {
		int __h = 0;
		if (__goOn[46] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -212068866860701141LL, (void*)event, &__h); if (__h) return; }
		QTextEdit::tabletEvent(event);
	}
	void actionEvent(QActionEvent* event) override {
		int __h = 0;
		if (__goOn[47] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1697444140088555581LL, (void*)event, &__h); if (__h) return; }
		QTextEdit::actionEvent(event);
	}
	void hideEvent(QHideEvent* event) override {
		int __h = 0;
		if (__goOn[48] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -199718707876719129LL, (void*)event, &__h); if (__h) return; }
		QTextEdit::hideEvent(event);
	}
	bool nativeEvent(const QByteArray& eventType, void* message, qintptr* result) override {
		int __h = 0;
		if (__goOn[49] && cppgo_virt_int_voidp_voidp_longlongp) { int __r = ((int(*)(long long, long long, void*, void*, long long*, int*))cppgo_virt_int_voidp_voidp_longlongp)(__goId, -2039166586246443706LL, (void*)&eventType, (void*)message, (long long*)result, &__h); if (__h) return __r; }
		return QTextEdit::nativeEvent(eventType, message, result);
	}
	int metric(QPaintDevice::PaintDeviceMetric arg1) const override {
		int __h = 0;
		if (__goOn[50] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 427629649267311777LL, (int)(arg1), &__h); if (__h) return __r; }
		return QTextEdit::metric(arg1);
	}
	void initPainter(QPainter* painter) const override {
		int __h = 0;
		if (__goOn[51] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6477500296980050210LL, (void*)painter, &__h); if (__h) return; }
		QTextEdit::initPainter(painter);
	}
	QPaintDevice* redirected(QPoint* offset) const override {
		int __h = 0;
		if (__goOn[52] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 4127240633310655048LL, (void*)offset, &__h); if (__h) return (QPaintDevice*)__r; }
		return QTextEdit::redirected(offset);
	}
	QPainter* sharedPainter() const override {
		int __h = 0;
		if (__goOn[53] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 5792771328741774991LL, &__h); if (__h) return (QPainter*)__r; }
		return QTextEdit::sharedPainter();
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[54] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QTextEdit::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[55] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QTextEdit::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[56] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QTextEdit::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[57] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QTextEdit::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QTextEdit;
static std::mutex __connsMu_QTextEdit;

extern "C" {

void* QTextEdit_base(void* self) { return (void*)static_cast<QAbstractScrollArea*>((QTextEdit*)self); }
void QTextEdit_chain(void* self, void** out) {
	QTextEdit* __p = (QTextEdit*)self;
	QAbstractScrollArea* __b0 = static_cast<QAbstractScrollArea*>(__p); out[0] = (void*)__b0;
	QFrame* __b1 = static_cast<QFrame*>(__b0); out[1] = (void*)__b1;
	QWidget* __b2 = static_cast<QWidget*>(__b1); out[2] = (void*)__b2;
	QObject* __b3 = static_cast<QObject*>(__b2); out[3] = (void*)__b3;
}
void QTextEdit__set_override(void* self, int idx) { ((QTextEdit__Go*)self)->__goOn[idx] = 1; }
void* QTextEdit_fromQObject(void* o) { return (void*)qobject_cast<QTextEdit*>((QObject*)o); }
void* QTextEdit_new(long long __goid, void* parent) {
	return new QTextEdit__Go(__goid, (QWidget*)parent);
}
void* QTextEdit_new2(long long __goid) {
	return new QTextEdit__Go(__goid);
}
void* QTextEdit_new3(long long __goid, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	return new QTextEdit__Go(__goid, text_qstr);
}
void* QTextEdit_new4(long long __goid, const char* text, long long text_len, void* parent) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	return new QTextEdit__Go(__goid, text_qstr, (QWidget*)parent);
}
char* QTextEdit_tr(const char* s, long long* __out_len) {
	auto&& __r = QTextEdit::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QTextEdit_setDocument(void* self, void* document) {
	((QTextEdit*)self)->setDocument((QTextDocument*)document);
}
void* QTextEdit_document(void* self) {
	auto&& __r = ((QTextEdit*)self)->document();
	return (void*)(__r);
}
void QTextEdit_setPlaceholderText(void* self, const char* placeholderText, long long placeholderText_len) {
	QString placeholderText_qstr = QString::fromUtf8(placeholderText, placeholderText_len);
	((QTextEdit*)self)->setPlaceholderText(placeholderText_qstr);
}
char* QTextEdit_placeholderText(void* self, long long* __out_len) {
	auto&& __r = ((QTextEdit*)self)->placeholderText();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QTextEdit_setTextCursor(void* self, void* cursor) {
	((QTextEdit*)self)->setTextCursor(*(QTextCursor*)cursor);
}
void* QTextEdit_textCursor(void* self) {
	auto&& __r = ((QTextEdit*)self)->textCursor();
	return new QTextCursor(__r);
}
int QTextEdit_isReadOnly(void* self) {
	auto&& __r = ((QTextEdit*)self)->isReadOnly();
	return __r;
}
void QTextEdit_setReadOnly(void* self, int ro) {
	((QTextEdit*)self)->setReadOnly(ro);
}
void QTextEdit_setTextInteractionFlags(void* self, int flags) {
	((QTextEdit*)self)->setTextInteractionFlags(QFlags<Qt::TextInteractionFlag>(flags));
}
int QTextEdit_textInteractionFlags(void* self) {
	auto&& __r = ((QTextEdit*)self)->textInteractionFlags();
	return (int)(__r);
}
double QTextEdit_fontPointSize(void* self) {
	auto&& __r = ((QTextEdit*)self)->fontPointSize();
	return __r;
}
char* QTextEdit_fontFamily(void* self, long long* __out_len) {
	auto&& __r = ((QTextEdit*)self)->fontFamily();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QTextEdit_fontWeight(void* self) {
	auto&& __r = ((QTextEdit*)self)->fontWeight();
	return __r;
}
int QTextEdit_fontUnderline(void* self) {
	auto&& __r = ((QTextEdit*)self)->fontUnderline();
	return __r;
}
int QTextEdit_fontItalic(void* self) {
	auto&& __r = ((QTextEdit*)self)->fontItalic();
	return __r;
}
void* QTextEdit_textColor(void* self) {
	auto&& __r = ((QTextEdit*)self)->textColor();
	return new QColor(__r);
}
void* QTextEdit_textBackgroundColor(void* self) {
	auto&& __r = ((QTextEdit*)self)->textBackgroundColor();
	return new QColor(__r);
}
void* QTextEdit_currentFont(void* self) {
	auto&& __r = ((QTextEdit*)self)->currentFont();
	return new QFont(__r);
}
int QTextEdit_alignment(void* self) {
	auto&& __r = ((QTextEdit*)self)->alignment();
	return (int)(__r);
}
void QTextEdit_mergeCurrentCharFormat(void* self, void* modifier) {
	((QTextEdit*)self)->mergeCurrentCharFormat(*(QTextCharFormat*)modifier);
}
void QTextEdit_setCurrentCharFormat(void* self, void* format) {
	((QTextEdit*)self)->setCurrentCharFormat(*(QTextCharFormat*)format);
}
void* QTextEdit_currentCharFormat(void* self) {
	auto&& __r = ((QTextEdit*)self)->currentCharFormat();
	return new QTextCharFormat(__r);
}
int QTextEdit_autoFormatting(void* self) {
	auto&& __r = ((QTextEdit*)self)->autoFormatting();
	return (int)(__r);
}
void QTextEdit_setAutoFormatting(void* self, int features) {
	((QTextEdit*)self)->setAutoFormatting(QFlags<QTextEdit::AutoFormattingFlag>(features));
}
int QTextEdit_tabChangesFocus(void* self) {
	auto&& __r = ((QTextEdit*)self)->tabChangesFocus();
	return __r;
}
void QTextEdit_setTabChangesFocus(void* self, int b) {
	((QTextEdit*)self)->setTabChangesFocus(b);
}
void QTextEdit_setDocumentTitle(void* self, const char* title, long long title_len) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	((QTextEdit*)self)->setDocumentTitle(title_qstr);
}
char* QTextEdit_documentTitle(void* self, long long* __out_len) {
	auto&& __r = ((QTextEdit*)self)->documentTitle();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QTextEdit_isUndoRedoEnabled(void* self) {
	auto&& __r = ((QTextEdit*)self)->isUndoRedoEnabled();
	return __r;
}
void QTextEdit_setUndoRedoEnabled(void* self, int enable) {
	((QTextEdit*)self)->setUndoRedoEnabled(enable);
}
int QTextEdit_lineWrapMode(void* self) {
	auto&& __r = ((QTextEdit*)self)->lineWrapMode();
	return (int)(__r);
}
void QTextEdit_setLineWrapMode(void* self, int mode) {
	((QTextEdit*)self)->setLineWrapMode((QTextEdit::LineWrapMode)(mode));
}
int QTextEdit_lineWrapColumnOrWidth(void* self) {
	auto&& __r = ((QTextEdit*)self)->lineWrapColumnOrWidth();
	return __r;
}
void QTextEdit_setLineWrapColumnOrWidth(void* self, int w) {
	((QTextEdit*)self)->setLineWrapColumnOrWidth(w);
}
int QTextEdit_wordWrapMode(void* self) {
	auto&& __r = ((QTextEdit*)self)->wordWrapMode();
	return (int)(__r);
}
void QTextEdit_setWordWrapMode(void* self, int policy) {
	((QTextEdit*)self)->setWordWrapMode((QTextOption::WrapMode)(policy));
}
int QTextEdit_find(void* self, const char* exp, long long exp_len) {
	QString exp_qstr = QString::fromUtf8(exp, exp_len);
	auto&& __r = ((QTextEdit*)self)->find(exp_qstr);
	return __r;
}
int QTextEdit_find2(void* self, void* exp) {
	auto&& __r = ((QTextEdit*)self)->find(*(QRegularExpression*)exp);
	return __r;
}
char* QTextEdit_toPlainText(void* self, long long* __out_len) {
	auto&& __r = ((QTextEdit*)self)->toPlainText();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QTextEdit_toHtml(void* self, long long* __out_len) {
	auto&& __r = ((QTextEdit*)self)->toHtml();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QTextEdit_toMarkdown(void* self, long long* __out_len) {
	auto&& __r = ((QTextEdit*)self)->toMarkdown();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QTextEdit_ensureCursorVisible(void* self) {
	((QTextEdit*)self)->ensureCursorVisible();
}
void* QTextEdit_loadResource(void* self, int type_, void* name) {
	auto&& __r = ((QTextEdit*)self)->loadResource(type_, *(QUrl*)name);
	return new QVariant(__r);
}
void* QTextEdit_createStandardContextMenu(void* self) {
	auto&& __r = ((QTextEdit*)self)->createStandardContextMenu();
	return (void*)(__r);
}
void* QTextEdit_createStandardContextMenu2(void* self, void* position) {
	auto&& __r = ((QTextEdit*)self)->createStandardContextMenu(*(QPoint*)position);
	return (void*)(__r);
}
void* QTextEdit_cursorForPosition(void* self, void* pos) {
	auto&& __r = ((QTextEdit*)self)->cursorForPosition(*(QPoint*)pos);
	return new QTextCursor(__r);
}
void* QTextEdit_cursorRect(void* self, void* cursor) {
	auto&& __r = ((QTextEdit*)self)->cursorRect(*(QTextCursor*)cursor);
	return new QRect(__r);
}
void* QTextEdit_cursorRect2(void* self) {
	auto&& __r = ((QTextEdit*)self)->cursorRect();
	return new QRect(__r);
}
char* QTextEdit_anchorAt(void* self, void* pos, long long* __out_len) {
	auto&& __r = ((QTextEdit*)self)->anchorAt(*(QPoint*)pos);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QTextEdit_overwriteMode(void* self) {
	auto&& __r = ((QTextEdit*)self)->overwriteMode();
	return __r;
}
void QTextEdit_setOverwriteMode(void* self, int overwrite) {
	((QTextEdit*)self)->setOverwriteMode(overwrite);
}
double QTextEdit_tabStopDistance(void* self) {
	auto&& __r = ((QTextEdit*)self)->tabStopDistance();
	return __r;
}
void QTextEdit_setTabStopDistance(void* self, double distance) {
	((QTextEdit*)self)->setTabStopDistance(distance);
}
int QTextEdit_cursorWidth(void* self) {
	auto&& __r = ((QTextEdit*)self)->cursorWidth();
	return __r;
}
void QTextEdit_setCursorWidth(void* self, int width) {
	((QTextEdit*)self)->setCursorWidth(width);
}
int QTextEdit_acceptRichText(void* self) {
	auto&& __r = ((QTextEdit*)self)->acceptRichText();
	return __r;
}
void QTextEdit_setAcceptRichText(void* self, int accept) {
	((QTextEdit*)self)->setAcceptRichText(accept);
}
void QTextEdit_setExtraSelections(void* self, void** selections, long long selections_len) {
	QList<QTextEdit::ExtraSelection> selections_list; selections_list.reserve(selections_len);
	for (long long __i = 0; __i < selections_len; __i++) selections_list.append(*(QTextEdit::ExtraSelection*)selections[__i]);
	((QTextEdit*)self)->setExtraSelections(selections_list);
}
void** QTextEdit_extraSelections(void* self, long long* __out_len) {
	auto&& __r = ((QTextEdit*)self)->extraSelections();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QTextEdit::ExtraSelection(__e);
	return __arr;
}
void QTextEdit_moveCursor(void* self, int operation) {
	((QTextEdit*)self)->moveCursor((QTextCursor::MoveOperation)(operation));
}
int QTextEdit_canPaste(void* self) {
	auto&& __r = ((QTextEdit*)self)->canPaste();
	return __r;
}
void QTextEdit_print(void* self, void* printer) {
	((QTextEdit*)self)->print((QPagedPaintDevice*)printer);
}
void* QTextEdit_inputMethodQuery(void* self, int property) {
	auto&& __r = ((QTextEdit*)self)->inputMethodQuery((Qt::InputMethodQuery)(property));
	return new QVariant(__r);
}
void* QTextEdit_inputMethodQuery2(void* self, int query, void* argument) {
	auto&& __r = ((QTextEdit*)self)->inputMethodQuery((Qt::InputMethodQuery)(query), *(QVariant*)argument);
	return new QVariant(__r);
}
void QTextEdit_setFontPointSize(void* self, double s) {
	((QTextEdit*)self)->setFontPointSize(s);
}
void QTextEdit_setFontFamily(void* self, const char* fontFamily, long long fontFamily_len) {
	QString fontFamily_qstr = QString::fromUtf8(fontFamily, fontFamily_len);
	((QTextEdit*)self)->setFontFamily(fontFamily_qstr);
}
void QTextEdit_setFontWeight(void* self, int w) {
	((QTextEdit*)self)->setFontWeight(w);
}
void QTextEdit_setFontUnderline(void* self, int b) {
	((QTextEdit*)self)->setFontUnderline(b);
}
void QTextEdit_setFontItalic(void* self, int b) {
	((QTextEdit*)self)->setFontItalic(b);
}
void QTextEdit_setTextColor(void* self, void* c) {
	((QTextEdit*)self)->setTextColor(*(QColor*)c);
}
void QTextEdit_setTextBackgroundColor(void* self, void* c) {
	((QTextEdit*)self)->setTextBackgroundColor(*(QColor*)c);
}
void QTextEdit_setCurrentFont(void* self, void* f) {
	((QTextEdit*)self)->setCurrentFont(*(QFont*)f);
}
void QTextEdit_setAlignment(void* self, int a) {
	((QTextEdit*)self)->setAlignment(QFlags<Qt::AlignmentFlag>(a));
}
void QTextEdit_setPlainText(void* self, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	((QTextEdit*)self)->setPlainText(text_qstr);
}
void QTextEdit_setHtml(void* self, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	((QTextEdit*)self)->setHtml(text_qstr);
}
void QTextEdit_setMarkdown(void* self, const char* markdown, long long markdown_len) {
	QString markdown_qstr = QString::fromUtf8(markdown, markdown_len);
	((QTextEdit*)self)->setMarkdown(markdown_qstr);
}
void QTextEdit_setText(void* self, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	((QTextEdit*)self)->setText(text_qstr);
}
void QTextEdit_cut(void* self) {
	((QTextEdit*)self)->cut();
}
void QTextEdit_copy(void* self) {
	((QTextEdit*)self)->copy();
}
void QTextEdit_paste(void* self) {
	((QTextEdit*)self)->paste();
}
void QTextEdit_undo(void* self) {
	((QTextEdit*)self)->undo();
}
void QTextEdit_redo(void* self) {
	((QTextEdit*)self)->redo();
}
void QTextEdit_clear(void* self) {
	((QTextEdit*)self)->clear();
}
void QTextEdit_selectAll(void* self) {
	((QTextEdit*)self)->selectAll();
}
void QTextEdit_insertPlainText(void* self, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	((QTextEdit*)self)->insertPlainText(text_qstr);
}
void QTextEdit_insertHtml(void* self, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	((QTextEdit*)self)->insertHtml(text_qstr);
}
void QTextEdit_append(void* self, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	((QTextEdit*)self)->append(text_qstr);
}
void QTextEdit_scrollToAnchor(void* self, const char* name, long long name_len) {
	QString name_qstr = QString::fromUtf8(name, name_len);
	((QTextEdit*)self)->scrollToAnchor(name_qstr);
}
void QTextEdit_zoomIn(void* self) {
	((QTextEdit*)self)->zoomIn();
}
void QTextEdit_zoomOut(void* self) {
	((QTextEdit*)self)->zoomOut();
}
void QTextEdit_textChanged(void* self) {
	((QTextEdit*)self)->textChanged();
}
void QTextEdit_undoAvailable(void* self, int b) {
	((QTextEdit*)self)->undoAvailable(b);
}
void QTextEdit_redoAvailable(void* self, int b) {
	((QTextEdit*)self)->redoAvailable(b);
}
void QTextEdit_currentCharFormatChanged(void* self, void* format) {
	((QTextEdit*)self)->currentCharFormatChanged(*(QTextCharFormat*)format);
}
void QTextEdit_copyAvailable(void* self, int b) {
	((QTextEdit*)self)->copyAvailable(b);
}
void QTextEdit_selectionChanged(void* self) {
	((QTextEdit*)self)->selectionChanged();
}
void QTextEdit_cursorPositionChanged(void* self) {
	((QTextEdit*)self)->cursorPositionChanged();
}
void QTextEdit_zoomInF(void* self, float range_) {
	((QTextEdit__Go*)self)->zoomInF(range_);
}
void QTextEdit_setViewportMargins(void* self, int left, int top, int right, int bottom) {
	((QTextEdit__Go*)self)->setViewportMargins(left, top, right, bottom);
}
void QTextEdit_setViewportMargins2(void* self, void* margins) {
	((QTextEdit__Go*)self)->setViewportMargins(*(QMargins*)margins);
}
void* QTextEdit_viewportMargins(void* self) {
	auto&& __r = ((QTextEdit__Go*)self)->viewportMargins();
	return new QMargins(__r);
}
void QTextEdit_drawFrame(void* self, void* arg1) {
	((QTextEdit__Go*)self)->drawFrame((QPainter*)arg1);
}
void QTextEdit_setTabOrder(void* arg1, void* arg2) {
	QTextEdit::setTabOrder((QWidget*)arg1, (QWidget*)arg2);
}
void* QTextEdit_mouseGrabber() {
	auto&& __r = QTextEdit::mouseGrabber();
	return (void*)(__r);
}
void* QTextEdit_keyboardGrabber() {
	auto&& __r = QTextEdit::keyboardGrabber();
	return (void*)(__r);
}
void* QTextEdit_createWindowContainer(void* window) {
	auto&& __r = QTextEdit::createWindowContainer((QWindow*)window);
	return (void*)(__r);
}
void QTextEdit_updateMicroFocus(void* self) {
	((QTextEdit__Go*)self)->updateMicroFocus();
}
void QTextEdit_create(void* self) {
	((QTextEdit__Go*)self)->create();
}
void QTextEdit_destroy(void* self) {
	((QTextEdit__Go*)self)->destroy();
}
int QTextEdit_focusNextChild(void* self) {
	auto&& __r = ((QTextEdit__Go*)self)->focusNextChild();
	return __r;
}
int QTextEdit_focusPreviousChild(void* self) {
	auto&& __r = ((QTextEdit__Go*)self)->focusPreviousChild();
	return __r;
}
void* QTextEdit_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QTextEdit::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QTextEdit_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QTextEdit::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QTextEdit_sender(void* self) {
	auto&& __r = ((QTextEdit__Go*)self)->sender();
	return (void*)(__r);
}
int QTextEdit_senderSignalIndex(void* self) {
	auto&& __r = ((QTextEdit__Go*)self)->senderSignalIndex();
	return __r;
}
int QTextEdit_receivers(void* self, const char* signal) {
	auto&& __r = ((QTextEdit__Go*)self)->receivers(signal);
	return __r;
}
int QTextEdit_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QTextEdit__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
double QTextEdit_devicePixelRatioFScale() {
	auto&& __r = QTextEdit::devicePixelRatioFScale();
	return __r;
}
int QTextEdit_encodeMetricF(int metric, double value) {
	auto&& __r = QTextEdit::encodeMetricF((QPaintDevice::PaintDeviceMetric)(metric), value);
	return __r;
}
double QTextEdit_getDecodedMetricF(void* self, int metricA, int metricB) {
	auto&& __r = ((QTextEdit__Go*)self)->getDecodedMetricF((QPaintDevice::PaintDeviceMetric)(metricA), (QPaintDevice::PaintDeviceMetric)(metricB));
	return __r;
}
char* QTextEdit_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QTextEdit::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QTextEdit_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QTextEdit::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QTextEdit_find3(void* self, const char* exp, long long exp_len, int options) {
	QString exp_qstr = QString::fromUtf8(exp, exp_len);
	auto&& __r = ((QTextEdit*)self)->find(exp_qstr, QFlags<QTextDocument::FindFlag>(options));
	return __r;
}
int QTextEdit_find4(void* self, void* exp, int options) {
	auto&& __r = ((QTextEdit*)self)->find(*(QRegularExpression*)exp, QFlags<QTextDocument::FindFlag>(options));
	return __r;
}
char* QTextEdit_toMarkdown2(void* self, int features, long long* __out_len) {
	auto&& __r = ((QTextEdit*)self)->toMarkdown(QFlags<QTextDocument::MarkdownFeature>(features));
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QTextEdit_moveCursor2(void* self, int operation, int mode) {
	((QTextEdit*)self)->moveCursor((QTextCursor::MoveOperation)(operation), (QTextCursor::MoveMode)(mode));
}
void QTextEdit_zoomIn2(void* self, int range_) {
	((QTextEdit*)self)->zoomIn(range_);
}
void QTextEdit_zoomOut2(void* self, int range_) {
	((QTextEdit*)self)->zoomOut(range_);
}
void* QTextEdit_createWindowContainer2(void* window, void* parent) {
	auto&& __r = QTextEdit::createWindowContainer((QWindow*)window, (QWidget*)parent);
	return (void*)(__r);
}
void* QTextEdit_createWindowContainer3(void* window, void* parent, int flags) {
	auto&& __r = QTextEdit::createWindowContainer((QWindow*)window, (QWidget*)parent, QFlags<Qt::WindowType>(flags));
	return (void*)(__r);
}
void QTextEdit_updateMicroFocus2(void* self, int query) {
	((QTextEdit__Go*)self)->updateMicroFocus((Qt::InputMethodQuery)(query));
}
void QTextEdit_create2(void* self, unsigned long long arg1) {
	((QTextEdit__Go*)self)->create(arg1);
}
void QTextEdit_create3(void* self, unsigned long long arg1, int initializeWindow) {
	((QTextEdit__Go*)self)->create(arg1, initializeWindow);
}
void QTextEdit_create4(void* self, unsigned long long arg1, int initializeWindow, int destroyOldWindow) {
	((QTextEdit__Go*)self)->create(arg1, initializeWindow, destroyOldWindow);
}
void QTextEdit_destroy2(void* self, int destroyWindow) {
	((QTextEdit__Go*)self)->destroy(destroyWindow);
}
void QTextEdit_destroy3(void* self, int destroyWindow, int destroySubWindows) {
	((QTextEdit__Go*)self)->destroy(destroyWindow, destroySubWindows);
}
void* QTextEdit_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QTextEdit::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QTextEdit_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QTextEdit::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QTextEdit_loadResource_default(void* self, int type_, void* name) {
	auto&& __r = ((QTextEdit*)self)->QTextEdit::loadResource(type_, *(QUrl*)name);
	return new QVariant(__r);
}
void* QTextEdit_inputMethodQuery_default(void* self, int property) {
	auto&& __r = ((QTextEdit*)self)->QTextEdit::inputMethodQuery((Qt::InputMethodQuery)(property));
	return new QVariant(__r);
}
void* QTextEdit_inputMethodQuery_default2(void* self, int query, void* argument) {
	auto&& __r = ((QTextEdit*)self)->QTextEdit::inputMethodQuery((Qt::InputMethodQuery)(query), *(QVariant*)argument);
	return new QVariant(__r);
}
void* QTextEdit_metaObject_default(void* self) {
	auto&& __r = ((QTextEdit*)self)->QTextEdit::metaObject();
	return (void*)(__r);
}
int QTextEdit_event_default(void* self, void* e) {
	auto&& __r = ((QTextEdit__Go*)self)->cppgoBase_event((QEvent*)e);
	return __r;
}
void QTextEdit_timerEvent_default(void* self, void* e) {
	((QTextEdit__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)e);
}
void QTextEdit_keyPressEvent_default(void* self, void* e) {
	((QTextEdit__Go*)self)->cppgoBase_keyPressEvent((QKeyEvent*)e);
}
void QTextEdit_keyReleaseEvent_default(void* self, void* e) {
	((QTextEdit__Go*)self)->cppgoBase_keyReleaseEvent((QKeyEvent*)e);
}
void QTextEdit_resizeEvent_default(void* self, void* e) {
	((QTextEdit__Go*)self)->cppgoBase_resizeEvent((QResizeEvent*)e);
}
void QTextEdit_paintEvent_default(void* self, void* e) {
	((QTextEdit__Go*)self)->cppgoBase_paintEvent((QPaintEvent*)e);
}
void QTextEdit_mousePressEvent_default(void* self, void* e) {
	((QTextEdit__Go*)self)->cppgoBase_mousePressEvent((QMouseEvent*)e);
}
void QTextEdit_mouseMoveEvent_default(void* self, void* e) {
	((QTextEdit__Go*)self)->cppgoBase_mouseMoveEvent((QMouseEvent*)e);
}
void QTextEdit_mouseReleaseEvent_default(void* self, void* e) {
	((QTextEdit__Go*)self)->cppgoBase_mouseReleaseEvent((QMouseEvent*)e);
}
void QTextEdit_mouseDoubleClickEvent_default(void* self, void* e) {
	((QTextEdit__Go*)self)->cppgoBase_mouseDoubleClickEvent((QMouseEvent*)e);
}
int QTextEdit_focusNextPrevChild_default(void* self, int next) {
	auto&& __r = ((QTextEdit__Go*)self)->cppgoBase_focusNextPrevChild(next);
	return __r;
}
void QTextEdit_contextMenuEvent_default(void* self, void* e) {
	((QTextEdit__Go*)self)->cppgoBase_contextMenuEvent((QContextMenuEvent*)e);
}
void QTextEdit_dragEnterEvent_default(void* self, void* e) {
	((QTextEdit__Go*)self)->cppgoBase_dragEnterEvent((QDragEnterEvent*)e);
}
void QTextEdit_dragLeaveEvent_default(void* self, void* e) {
	((QTextEdit__Go*)self)->cppgoBase_dragLeaveEvent((QDragLeaveEvent*)e);
}
void QTextEdit_dragMoveEvent_default(void* self, void* e) {
	((QTextEdit__Go*)self)->cppgoBase_dragMoveEvent((QDragMoveEvent*)e);
}
void QTextEdit_dropEvent_default(void* self, void* e) {
	((QTextEdit__Go*)self)->cppgoBase_dropEvent((QDropEvent*)e);
}
void QTextEdit_focusInEvent_default(void* self, void* e) {
	((QTextEdit__Go*)self)->cppgoBase_focusInEvent((QFocusEvent*)e);
}
void QTextEdit_focusOutEvent_default(void* self, void* e) {
	((QTextEdit__Go*)self)->cppgoBase_focusOutEvent((QFocusEvent*)e);
}
void QTextEdit_showEvent_default(void* self, void* arg1) {
	((QTextEdit__Go*)self)->cppgoBase_showEvent((QShowEvent*)arg1);
}
void QTextEdit_changeEvent_default(void* self, void* e) {
	((QTextEdit__Go*)self)->cppgoBase_changeEvent((QEvent*)e);
}
void QTextEdit_wheelEvent_default(void* self, void* e) {
	((QTextEdit__Go*)self)->cppgoBase_wheelEvent((QWheelEvent*)e);
}
void* QTextEdit_createMimeDataFromSelection_default(void* self) {
	auto&& __r = ((QTextEdit__Go*)self)->cppgoBase_createMimeDataFromSelection();
	return (void*)(__r);
}
int QTextEdit_canInsertFromMimeData_default(void* self, void* source) {
	auto&& __r = ((QTextEdit__Go*)self)->cppgoBase_canInsertFromMimeData((QMimeData*)source);
	return __r;
}
void QTextEdit_insertFromMimeData_default(void* self, void* source) {
	((QTextEdit__Go*)self)->cppgoBase_insertFromMimeData((QMimeData*)source);
}
void QTextEdit_inputMethodEvent_default(void* self, void* arg1) {
	((QTextEdit__Go*)self)->cppgoBase_inputMethodEvent((QInputMethodEvent*)arg1);
}
void QTextEdit_scrollContentsBy_default(void* self, int dx, int dy) {
	((QTextEdit__Go*)self)->cppgoBase_scrollContentsBy(dx, dy);
}
void QTextEdit_doSetTextCursor_default(void* self, void* cursor) {
	((QTextEdit__Go*)self)->cppgoBase_doSetTextCursor(*(QTextCursor*)cursor);
}
void* QTextEdit_minimumSizeHint_default(void* self) {
	auto&& __r = ((QTextEdit*)self)->QTextEdit::minimumSizeHint();
	return new QSize(__r);
}
void* QTextEdit_sizeHint_default(void* self) {
	auto&& __r = ((QTextEdit*)self)->QTextEdit::sizeHint();
	return new QSize(__r);
}
void QTextEdit_setupViewport_default(void* self, void* viewport) {
	((QTextEdit*)self)->QTextEdit::setupViewport((QWidget*)viewport);
}
int QTextEdit_eventFilter_default(void* self, void* arg1, void* arg2) {
	auto&& __r = ((QTextEdit__Go*)self)->cppgoBase_eventFilter((QObject*)arg1, (QEvent*)arg2);
	return __r;
}
int QTextEdit_viewportEvent_default(void* self, void* arg1) {
	auto&& __r = ((QTextEdit__Go*)self)->cppgoBase_viewportEvent((QEvent*)arg1);
	return __r;
}
void* QTextEdit_viewportSizeHint_default(void* self) {
	auto&& __r = ((QTextEdit__Go*)self)->cppgoBase_viewportSizeHint();
	return new QSize(__r);
}
void QTextEdit_initStyleOption_default(void* self, void* option) {
	((QTextEdit__Go*)self)->cppgoBase_initStyleOption((QStyleOptionFrame*)option);
}
int QTextEdit_devType_default(void* self) {
	auto&& __r = ((QTextEdit*)self)->QTextEdit::devType();
	return __r;
}
void QTextEdit_setVisible_default(void* self, int visible) {
	((QTextEdit*)self)->QTextEdit::setVisible(visible);
}
int QTextEdit_heightForWidth_default(void* self, int arg1) {
	auto&& __r = ((QTextEdit*)self)->QTextEdit::heightForWidth(arg1);
	return __r;
}
int QTextEdit_hasHeightForWidth_default(void* self) {
	auto&& __r = ((QTextEdit*)self)->QTextEdit::hasHeightForWidth();
	return __r;
}
void* QTextEdit_paintEngine_default(void* self) {
	auto&& __r = ((QTextEdit*)self)->QTextEdit::paintEngine();
	return (void*)(__r);
}
void QTextEdit_enterEvent_default(void* self, void* event) {
	((QTextEdit__Go*)self)->cppgoBase_enterEvent((QEnterEvent*)event);
}
void QTextEdit_leaveEvent_default(void* self, void* event) {
	((QTextEdit__Go*)self)->cppgoBase_leaveEvent((QEvent*)event);
}
void QTextEdit_moveEvent_default(void* self, void* event) {
	((QTextEdit__Go*)self)->cppgoBase_moveEvent((QMoveEvent*)event);
}
void QTextEdit_closeEvent_default(void* self, void* event) {
	((QTextEdit__Go*)self)->cppgoBase_closeEvent((QCloseEvent*)event);
}
void QTextEdit_tabletEvent_default(void* self, void* event) {
	((QTextEdit__Go*)self)->cppgoBase_tabletEvent((QTabletEvent*)event);
}
void QTextEdit_actionEvent_default(void* self, void* event) {
	((QTextEdit__Go*)self)->cppgoBase_actionEvent((QActionEvent*)event);
}
void QTextEdit_hideEvent_default(void* self, void* event) {
	((QTextEdit__Go*)self)->cppgoBase_hideEvent((QHideEvent*)event);
}
int QTextEdit_metric_default(void* self, int arg1) {
	auto&& __r = ((QTextEdit__Go*)self)->cppgoBase_metric(arg1);
	return __r;
}
void QTextEdit_initPainter_default(void* self, void* painter) {
	((QTextEdit__Go*)self)->cppgoBase_initPainter((QPainter*)painter);
}
void* QTextEdit_redirected_default(void* self, void* offset) {
	auto&& __r = ((QTextEdit__Go*)self)->cppgoBase_redirected((QPoint*)offset);
	return (void*)(__r);
}
void* QTextEdit_sharedPainter_default(void* self) {
	auto&& __r = ((QTextEdit__Go*)self)->cppgoBase_sharedPainter();
	return (void*)(__r);
}
void QTextEdit_childEvent_default(void* self, void* event) {
	((QTextEdit__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QTextEdit_customEvent_default(void* self, void* event) {
	((QTextEdit__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QTextEdit_connectNotify_default(void* self, void* signal) {
	((QTextEdit__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QTextEdit_disconnectNotify_default(void* self, void* signal) {
	((QTextEdit__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QTextEdit_connect_textChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTextEdit*)self, &QTextEdit::textChanged, (QTextEdit*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTextEdit);
	__conns_QTextEdit[__id] = __c;
}
void QTextEdit_connect_undoAvailable(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTextEdit*)self, &QTextEdit::undoAvailable, (QTextEdit*)self, [__id](bool b){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, b);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTextEdit);
	__conns_QTextEdit[__id] = __c;
}
void QTextEdit_connect_redoAvailable(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTextEdit*)self, &QTextEdit::redoAvailable, (QTextEdit*)self, [__id](bool b){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, b);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTextEdit);
	__conns_QTextEdit[__id] = __c;
}
void QTextEdit_connect_currentCharFormatChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTextEdit*)self, &QTextEdit::currentCharFormatChanged, (QTextEdit*)self, [__id](const QTextCharFormat& format){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&format);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTextEdit);
	__conns_QTextEdit[__id] = __c;
}
void QTextEdit_connect_copyAvailable(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTextEdit*)self, &QTextEdit::copyAvailable, (QTextEdit*)self, [__id](bool b){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, b);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTextEdit);
	__conns_QTextEdit[__id] = __c;
}
void QTextEdit_connect_selectionChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTextEdit*)self, &QTextEdit::selectionChanged, (QTextEdit*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTextEdit);
	__conns_QTextEdit[__id] = __c;
}
void QTextEdit_connect_cursorPositionChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTextEdit*)self, &QTextEdit::cursorPositionChanged, (QTextEdit*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTextEdit);
	__conns_QTextEdit[__id] = __c;
}
void QTextEdit_connect_windowTitleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTextEdit*)self, &QTextEdit::windowTitleChanged, (QTextEdit*)self, [__id](const QString& title){
		QByteArray title_ba = title.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)title_ba.constData(), title_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTextEdit);
	__conns_QTextEdit[__id] = __c;
}
void QTextEdit_connect_windowIconChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTextEdit*)self, &QTextEdit::windowIconChanged, (QTextEdit*)self, [__id](const QIcon& icon){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&icon);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTextEdit);
	__conns_QTextEdit[__id] = __c;
}
void QTextEdit_connect_windowIconTextChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTextEdit*)self, &QTextEdit::windowIconTextChanged, (QTextEdit*)self, [__id](const QString& iconText){
		QByteArray iconText_ba = iconText.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)iconText_ba.constData(), iconText_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTextEdit);
	__conns_QTextEdit[__id] = __c;
}
void QTextEdit_connect_customContextMenuRequested(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTextEdit*)self, &QTextEdit::customContextMenuRequested, (QTextEdit*)self, [__id](const QPoint& pos){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&pos);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTextEdit);
	__conns_QTextEdit[__id] = __c;
}
void QTextEdit_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTextEdit*)self, &QTextEdit::destroyed, (QTextEdit*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTextEdit);
	__conns_QTextEdit[__id] = __c;
}
void QTextEdit_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTextEdit*)self, &QTextEdit::destroyed, (QTextEdit*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTextEdit);
	__conns_QTextEdit[__id] = __c;
}
void QTextEdit_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTextEdit*)self, &QTextEdit::objectNameChanged, (QTextEdit*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTextEdit);
	__conns_QTextEdit[__id] = __c;
}
void QTextEdit__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QTextEdit);
		auto __it = __conns_QTextEdit.find(id);
		if (__it == __conns_QTextEdit.end()) return;
		__c = __it->second;
		__conns_QTextEdit.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QTextEdit__Go_new(long long id) { return new QTextEdit__Go(id); }
void QTextEdit_delete(void* self) { delete (QTextEdit*)self; }

}

#include <QtWidgets/qstyleplugin.h>
#include <QtCore/qobject.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtWidgets/qstyle.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qvariant.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qstyleplugin.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QStylePlugin__Go : public QStylePlugin {
public:
	long long __goId;
	unsigned char __goOn[9] = {};
	template<typename... A> QStylePlugin__Go(long long id, A&&... a) : QStylePlugin(std::forward<A>(a)...), __goId(id) {}
	using QStylePlugin::sender;
	using QStylePlugin::senderSignalIndex;
	using QStylePlugin::receivers;
	using QStylePlugin::isSignalConnected;
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QStylePlugin::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QStylePlugin::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QStylePlugin::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QStylePlugin::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QStylePlugin::disconnectNotify(__a0); }
	QStyle* create(const QString& key) override {
		int __h = 0;
		QByteArray key_ba = key.toUtf8();
		if (__goOn[0] && cppgo_virt_voidp_constcharp_longlong) { void* __r = ((void*(*)(long long, long long, const char*, long long, int*))cppgo_virt_voidp_constcharp_longlong)(__goId, -626457835795418115LL, (char*)key_ba.constData(), key_ba.size(), &__h); if (__h) return (QStyle*)__r; }
		return nullptr;
	}
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QStylePlugin::metaObject();
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QStylePlugin::event(event);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QStylePlugin::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QStylePlugin::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QStylePlugin::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QStylePlugin::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QStylePlugin::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QStylePlugin::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QStylePlugin;
static std::mutex __connsMu_QStylePlugin;

extern "C" {

void* QStylePlugin_base(void* self) { return (void*)static_cast<QObject*>((QStylePlugin*)self); }
void QStylePlugin__set_override(void* self, int idx) { ((QStylePlugin__Go*)self)->__goOn[idx] = 1; }
void* QStylePlugin_fromQObject(void* o) { return (void*)qobject_cast<QStylePlugin*>((QObject*)o); }
char* QStylePlugin_tr(const char* s, long long* __out_len) {
	auto&& __r = QStylePlugin::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QStylePlugin_create(void* self, const char* key, long long key_len) {
	QString key_qstr = QString::fromUtf8(key, key_len);
	auto&& __r = ((QStylePlugin*)self)->create(key_qstr);
	return (void*)(__r);
}
void* QStylePlugin_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QStylePlugin::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QStylePlugin_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QStylePlugin::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QStylePlugin_sender(void* self) {
	auto&& __r = ((QStylePlugin__Go*)self)->sender();
	return (void*)(__r);
}
int QStylePlugin_senderSignalIndex(void* self) {
	auto&& __r = ((QStylePlugin__Go*)self)->senderSignalIndex();
	return __r;
}
int QStylePlugin_receivers(void* self, const char* signal) {
	auto&& __r = ((QStylePlugin__Go*)self)->receivers(signal);
	return __r;
}
int QStylePlugin_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QStylePlugin__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QStylePlugin_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QStylePlugin::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QStylePlugin_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QStylePlugin::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QStylePlugin_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QStylePlugin::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QStylePlugin_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QStylePlugin::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QStylePlugin_metaObject_default(void* self) {
	auto&& __r = ((QStylePlugin*)self)->QStylePlugin::metaObject();
	return (void*)(__r);
}
int QStylePlugin_event_default(void* self, void* event) {
	auto&& __r = ((QStylePlugin*)self)->QStylePlugin::event((QEvent*)event);
	return __r;
}
int QStylePlugin_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QStylePlugin*)self)->QStylePlugin::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QStylePlugin_timerEvent_default(void* self, void* event) {
	((QStylePlugin__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QStylePlugin_childEvent_default(void* self, void* event) {
	((QStylePlugin__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QStylePlugin_customEvent_default(void* self, void* event) {
	((QStylePlugin__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QStylePlugin_connectNotify_default(void* self, void* signal) {
	((QStylePlugin__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QStylePlugin_disconnectNotify_default(void* self, void* signal) {
	((QStylePlugin__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QStylePlugin_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStylePlugin*)self, &QStylePlugin::destroyed, (QStylePlugin*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStylePlugin);
	__conns_QStylePlugin[__id] = __c;
}
void QStylePlugin_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStylePlugin*)self, &QStylePlugin::destroyed, (QStylePlugin*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStylePlugin);
	__conns_QStylePlugin[__id] = __c;
}
void QStylePlugin_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStylePlugin*)self, &QStylePlugin::objectNameChanged, (QStylePlugin*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStylePlugin);
	__conns_QStylePlugin[__id] = __c;
}
void QStylePlugin__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QStylePlugin);
		auto __it = __conns_QStylePlugin.find(id);
		if (__it == __conns_QStylePlugin.end()) return;
		__c = __it->second;
		__conns_QStylePlugin.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QStylePlugin__Go_new(long long id, void* parent) { return new QStylePlugin__Go(id, (QObject*)parent); }
void QStylePlugin_delete(void* self) { delete (QStylePlugin*)self; }

}

