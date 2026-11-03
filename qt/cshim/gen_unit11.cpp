#include <QtCore/qobjectdefs.h>
#include "gen_qgenericargument.h"

extern "C" {

void* QGenericArgument_new() {
	return new QGenericArgument();
}
void* QGenericArgument_new2(void* arg1) {
	return new QGenericArgument(*(QGenericArgument*)arg1);
}
void* QGenericArgument_new3(const char* aName) {
	return new QGenericArgument(aName);
}
const char* QGenericArgument_name(void* self) {
	auto&& __r = ((QGenericArgument*)self)->name();
	return __r;
}
void QGenericArgument_delete(void* self) { delete (QGenericArgument*)self; }

}

#include <QtCore/qhashfunctions.h>
#include "gen_qhashseed.h"

extern "C" {

void* QHashSeed_new() {
	return new QHashSeed();
}
void* QHashSeed_new2(unsigned long long d) {
	return new QHashSeed(d);
}
unsigned long long QHashSeed_ToSizeT(void* self) {
	auto&& __r = static_cast<unsigned long long>(*(QHashSeed*)self);
	return __r;
}
void* QHashSeed_globalSeed() {
	auto&& __r = QHashSeed::globalSeed();
	return new QHashSeed(__r);
}
void QHashSeed_setDeterministicGlobalSeed() {
	QHashSeed::setDeterministicGlobalSeed();
}
void QHashSeed_resetRandomGlobalSeed() {
	QHashSeed::resetRandomGlobalSeed();
}
void QHashSeed_delete(void* self) { delete (QHashSeed*)self; }

}

#include <QtCore/qtextstream.h>
#include <QtCore/qiodevicebase.h>
#include <QtCore/qiodevice.h>
#include <QtCore/qstring.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qstringconverter_base.h>
#include <QtCore/qlocale.h>
#include <QtCore/qchar.h>
#include <QtCore/qstringview.h>
#include <QtCore/qlatin1stringview.h>
#include <QtCore/qfloat16.h>
#include "gen_qtextstream.h"

extern "C" {

void* QTextStream_base(void* self) { return (void*)static_cast<QIODeviceBase*>((QTextStream*)self); }
void* QTextStream_new() {
	return new QTextStream();
}
void* QTextStream_new2(void* device) {
	return new QTextStream((QIODevice*)device);
}
void* QTextStream_new3(void* string_) {
	return new QTextStream((QString*)string_);
}
void* QTextStream_new4(void* array) {
	return new QTextStream((QByteArray*)array);
}
void* QTextStream_new5(void* array) {
	return new QTextStream(*(QByteArray*)array);
}
void* QTextStream_new6(void* string_, int openMode) {
	return new QTextStream((QString*)string_, QFlags<QIODeviceBase::OpenModeFlag>(openMode));
}
void* QTextStream_new7(void* array, int openMode) {
	return new QTextStream((QByteArray*)array, QFlags<QIODeviceBase::OpenModeFlag>(openMode));
}
void* QTextStream_new8(void* array, int openMode) {
	return new QTextStream(*(QByteArray*)array, QFlags<QIODeviceBase::OpenModeFlag>(openMode));
}
void QTextStream_setEncoding(void* self, int encoding) {
	((QTextStream*)self)->setEncoding((QStringConverter::Encoding)(encoding));
}
int QTextStream_encoding(void* self) {
	auto&& __r = ((QTextStream*)self)->encoding();
	return (int)(__r);
}
void QTextStream_setAutoDetectUnicode(void* self, int enabled) {
	((QTextStream*)self)->setAutoDetectUnicode(enabled);
}
int QTextStream_autoDetectUnicode(void* self) {
	auto&& __r = ((QTextStream*)self)->autoDetectUnicode();
	return __r;
}
void QTextStream_setGenerateByteOrderMark(void* self, int generate) {
	((QTextStream*)self)->setGenerateByteOrderMark(generate);
}
int QTextStream_generateByteOrderMark(void* self) {
	auto&& __r = ((QTextStream*)self)->generateByteOrderMark();
	return __r;
}
void QTextStream_setLocale(void* self, void* locale) {
	((QTextStream*)self)->setLocale(*(QLocale*)locale);
}
void* QTextStream_locale(void* self) {
	auto&& __r = ((QTextStream*)self)->locale();
	return new QLocale(__r);
}
void QTextStream_setDevice(void* self, void* device) {
	((QTextStream*)self)->setDevice((QIODevice*)device);
}
void* QTextStream_device(void* self) {
	auto&& __r = ((QTextStream*)self)->device();
	return (void*)(__r);
}
void QTextStream_setString(void* self, void* string_) {
	((QTextStream*)self)->setString((QString*)string_);
}
void* QTextStream_string(void* self) {
	auto&& __r = ((QTextStream*)self)->string();
	return (void*)(__r);
}
int QTextStream_status(void* self) {
	auto&& __r = ((QTextStream*)self)->status();
	return (int)(__r);
}
void QTextStream_setStatus(void* self, int status) {
	((QTextStream*)self)->setStatus((QTextStream::Status)(status));
}
void QTextStream_resetStatus(void* self) {
	((QTextStream*)self)->resetStatus();
}
int QTextStream_atEnd(void* self) {
	auto&& __r = ((QTextStream*)self)->atEnd();
	return __r;
}
void QTextStream_reset(void* self) {
	((QTextStream*)self)->reset();
}
void QTextStream_flush(void* self) {
	((QTextStream*)self)->flush();
}
int QTextStream_seek(void* self, long long pos) {
	auto&& __r = ((QTextStream*)self)->seek(pos);
	return __r;
}
long long QTextStream_pos(void* self) {
	auto&& __r = ((QTextStream*)self)->pos();
	return __r;
}
void QTextStream_skipWhiteSpace(void* self) {
	((QTextStream*)self)->skipWhiteSpace();
}
char* QTextStream_readLine(void* self, long long* __out_len) {
	auto&& __r = ((QTextStream*)self)->readLine();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QTextStream_readLineInto(void* self, void* line) {
	auto&& __r = ((QTextStream*)self)->readLineInto((QString*)line);
	return __r;
}
char* QTextStream_readAll(void* self, long long* __out_len) {
	auto&& __r = ((QTextStream*)self)->readAll();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QTextStream_read(void* self, long long maxlen, long long* __out_len) {
	auto&& __r = ((QTextStream*)self)->read(maxlen);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QTextStream_setFieldAlignment(void* self, int alignment) {
	((QTextStream*)self)->setFieldAlignment((QTextStream::FieldAlignment)(alignment));
}
int QTextStream_fieldAlignment(void* self) {
	auto&& __r = ((QTextStream*)self)->fieldAlignment();
	return (int)(__r);
}
void QTextStream_setPadChar(void* self, void* ch) {
	((QTextStream*)self)->setPadChar(*(QChar*)ch);
}
void* QTextStream_padChar(void* self) {
	auto&& __r = ((QTextStream*)self)->padChar();
	return new QChar(__r);
}
void QTextStream_setFieldWidth(void* self, int width) {
	((QTextStream*)self)->setFieldWidth(width);
}
int QTextStream_fieldWidth(void* self) {
	auto&& __r = ((QTextStream*)self)->fieldWidth();
	return __r;
}
void QTextStream_setNumberFlags(void* self, int flags) {
	((QTextStream*)self)->setNumberFlags(QFlags<QTextStream::NumberFlag>(flags));
}
int QTextStream_numberFlags(void* self) {
	auto&& __r = ((QTextStream*)self)->numberFlags();
	return (int)(__r);
}
void QTextStream_setIntegerBase(void* self, int base) {
	((QTextStream*)self)->setIntegerBase(base);
}
int QTextStream_integerBase(void* self) {
	auto&& __r = ((QTextStream*)self)->integerBase();
	return __r;
}
void QTextStream_setRealNumberNotation(void* self, int notation) {
	((QTextStream*)self)->setRealNumberNotation((QTextStream::RealNumberNotation)(notation));
}
int QTextStream_realNumberNotation(void* self) {
	auto&& __r = ((QTextStream*)self)->realNumberNotation();
	return (int)(__r);
}
void QTextStream_setRealNumberPrecision(void* self, int precision) {
	((QTextStream*)self)->setRealNumberPrecision(precision);
}
int QTextStream_realNumberPrecision(void* self) {
	auto&& __r = ((QTextStream*)self)->realNumberPrecision();
	return __r;
}
void* QTextStream_OperatorShiftRight(void* self, void* ch) {
	auto&& __r = ((QTextStream*)self)->operator>>(*(QChar*)ch);
	return (void*)&(__r);
}
void* QTextStream_OperatorShiftRight2(void* self, char* ch) {
	char ch_ref = *ch;
	auto&& __r = ((QTextStream*)self)->operator>>(ch_ref);
	*ch = ch_ref;
	return (void*)&(__r);
}
void* QTextStream_OperatorShiftRight3(void* self, unsigned short* ch) {
	char16_t ch_ref = *ch;
	auto&& __r = ((QTextStream*)self)->operator>>(ch_ref);
	*ch = ch_ref;
	return (void*)&(__r);
}
void* QTextStream_OperatorShiftRight4(void* self, short* i) {
	short i_ref = *i;
	auto&& __r = ((QTextStream*)self)->operator>>(i_ref);
	*i = i_ref;
	return (void*)&(__r);
}
void* QTextStream_OperatorShiftRight5(void* self, unsigned short* i) {
	unsigned short i_ref = *i;
	auto&& __r = ((QTextStream*)self)->operator>>(i_ref);
	*i = i_ref;
	return (void*)&(__r);
}
void* QTextStream_OperatorShiftRight6(void* self, int* i) {
	int i_ref = *i;
	auto&& __r = ((QTextStream*)self)->operator>>(i_ref);
	*i = i_ref;
	return (void*)&(__r);
}
void* QTextStream_OperatorShiftRight7(void* self, unsigned int* i) {
	unsigned int i_ref = *i;
	auto&& __r = ((QTextStream*)self)->operator>>(i_ref);
	*i = i_ref;
	return (void*)&(__r);
}
void* QTextStream_OperatorShiftRight8(void* self, long* i) {
	long i_ref = *i;
	auto&& __r = ((QTextStream*)self)->operator>>(i_ref);
	*i = i_ref;
	return (void*)&(__r);
}
void* QTextStream_OperatorShiftRight9(void* self, unsigned long* i) {
	unsigned long i_ref = *i;
	auto&& __r = ((QTextStream*)self)->operator>>(i_ref);
	*i = i_ref;
	return (void*)&(__r);
}
void* QTextStream_OperatorShiftRight10(void* self, long long* i) {
	qlonglong i_ref = *i;
	auto&& __r = ((QTextStream*)self)->operator>>(i_ref);
	*i = i_ref;
	return (void*)&(__r);
}
void* QTextStream_OperatorShiftRight11(void* self, unsigned long long* i) {
	qulonglong i_ref = *i;
	auto&& __r = ((QTextStream*)self)->operator>>(i_ref);
	*i = i_ref;
	return (void*)&(__r);
}
void* QTextStream_OperatorShiftRight12(void* self, float* f) {
	float f_ref = *f;
	auto&& __r = ((QTextStream*)self)->operator>>(f_ref);
	*f = f_ref;
	return (void*)&(__r);
}
void* QTextStream_OperatorShiftRight13(void* self, double* f) {
	double f_ref = *f;
	auto&& __r = ((QTextStream*)self)->operator>>(f_ref);
	*f = f_ref;
	return (void*)&(__r);
}
void* QTextStream_OperatorShiftRight14(void* self, void* s) {
	auto&& __r = ((QTextStream*)self)->operator>>(*(QString*)s);
	return (void*)&(__r);
}
void* QTextStream_OperatorShiftRight15(void* self, void* array) {
	auto&& __r = ((QTextStream*)self)->operator>>(*(QByteArray*)array);
	return (void*)&(__r);
}
void* QTextStream_OperatorShiftRight16(void* self, char* c) {
	auto&& __r = ((QTextStream*)self)->operator>>((char*)c);
	return (void*)&(__r);
}
void* QTextStream_OperatorShiftLeft(void* self, void* ch) {
	auto&& __r = ((QTextStream*)self)->operator<<(*(QChar*)ch);
	return (void*)&(__r);
}
void* QTextStream_OperatorShiftLeft2(void* self, char ch) {
	auto&& __r = ((QTextStream*)self)->operator<<(ch);
	return (void*)&(__r);
}
void* QTextStream_OperatorShiftLeft3(void* self, unsigned short ch) {
	auto&& __r = ((QTextStream*)self)->operator<<(ch);
	return (void*)&(__r);
}
void* QTextStream_OperatorShiftLeft4(void* self, short i) {
	auto&& __r = ((QTextStream*)self)->operator<<(i);
	return (void*)&(__r);
}
void* QTextStream_OperatorShiftLeft5(void* self, unsigned short i) {
	auto&& __r = ((QTextStream*)self)->operator<<(i);
	return (void*)&(__r);
}
void* QTextStream_OperatorShiftLeft6(void* self, int i) {
	auto&& __r = ((QTextStream*)self)->operator<<(i);
	return (void*)&(__r);
}
void* QTextStream_OperatorShiftLeft7(void* self, unsigned int i) {
	auto&& __r = ((QTextStream*)self)->operator<<(i);
	return (void*)&(__r);
}
void* QTextStream_OperatorShiftLeft8(void* self, long i) {
	auto&& __r = ((QTextStream*)self)->operator<<(i);
	return (void*)&(__r);
}
void* QTextStream_OperatorShiftLeft9(void* self, unsigned long i) {
	auto&& __r = ((QTextStream*)self)->operator<<(i);
	return (void*)&(__r);
}
void* QTextStream_OperatorShiftLeft10(void* self, long long i) {
	auto&& __r = ((QTextStream*)self)->operator<<(i);
	return (void*)&(__r);
}
void* QTextStream_OperatorShiftLeft11(void* self, unsigned long long i) {
	auto&& __r = ((QTextStream*)self)->operator<<(i);
	return (void*)&(__r);
}
void* QTextStream_OperatorShiftLeft12(void* self, float f) {
	auto&& __r = ((QTextStream*)self)->operator<<(f);
	return (void*)&(__r);
}
void* QTextStream_OperatorShiftLeft13(void* self, double f) {
	auto&& __r = ((QTextStream*)self)->operator<<(f);
	return (void*)&(__r);
}
void* QTextStream_OperatorShiftLeft14(void* self, const char* s, long long s_len) {
	QString s_qstr = QString::fromUtf8(s, s_len);
	auto&& __r = ((QTextStream*)self)->operator<<(s_qstr);
	return (void*)&(__r);
}
void* QTextStream_OperatorShiftLeft15(void* self, void* s) {
	auto&& __r = ((QTextStream*)self)->operator<<(*(QStringView*)s);
	return (void*)&(__r);
}
void* QTextStream_OperatorShiftLeft16(void* self, void* s) {
	auto&& __r = ((QTextStream*)self)->operator<<(*(QLatin1String*)s);
	return (void*)&(__r);
}
void* QTextStream_OperatorShiftLeft17(void* self, void* array) {
	auto&& __r = ((QTextStream*)self)->operator<<(*(QByteArray*)array);
	return (void*)&(__r);
}
void* QTextStream_OperatorShiftLeft18(void* self, const char* c) {
	auto&& __r = ((QTextStream*)self)->operator<<(c);
	return (void*)&(__r);
}
int QTextStream_ToBool(void* self) {
	auto&& __r = static_cast<bool>(*(QTextStream*)self);
	return __r;
}
void QTextStream_setString2(void* self, void* string_, int openMode) {
	((QTextStream*)self)->setString((QString*)string_, QFlags<QIODeviceBase::OpenModeFlag>(openMode));
}
char* QTextStream_readLine2(void* self, long long maxlen, long long* __out_len) {
	auto&& __r = ((QTextStream*)self)->readLine(maxlen);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QTextStream_readLineInto2(void* self, void* line, long long maxlen) {
	auto&& __r = ((QTextStream*)self)->readLineInto((QString*)line, maxlen);
	return __r;
}
void* QTextStream_OperatorShiftLeft19(void* self, void* f) {
	auto&& __r = (*(QTextStream*)self) << (*(qfloat16*)f);
	return (void*)&(__r);
}
void* QTextStream_OperatorShiftRight17(void* self, void* f) {
	auto&& __r = (*(QTextStream*)self) >> (*(qfloat16*)f);
	return (void*)&(__r);
}
void* QTextStream_OperatorShiftLeft20(void* self, void* m) {
	auto&& __r = (*(QTextStream*)self) << (*(QTextStreamManipulator*)m);
	return (void*)&(__r);
}
void QTextStream_delete(void* self) { delete (QTextStream*)self; }

}

#include <QtCore/qbasictimer.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qobject.h>
#include "gen_qbasictimer.h"

extern "C" {

void* QBasicTimer_new() {
	return new QBasicTimer();
}
void QBasicTimer_swap(void* self, void* other) {
	((QBasicTimer*)self)->swap(*(QBasicTimer*)other);
}
int QBasicTimer_isActive(void* self) {
	auto&& __r = ((QBasicTimer*)self)->isActive();
	return __r;
}
int QBasicTimer_timerId(void* self) {
	auto&& __r = ((QBasicTimer*)self)->timerId();
	return __r;
}
int QBasicTimer_id(void* self) {
	auto&& __r = ((QBasicTimer*)self)->id();
	return (int)(__r);
}
void QBasicTimer_start(void* self, int msec, void* obj) {
	((QBasicTimer*)self)->start(msec, (QObject*)obj);
}
void QBasicTimer_start2(void* self, int msec, int timerType, void* obj) {
	((QBasicTimer*)self)->start(msec, (Qt::TimerType)(timerType), (QObject*)obj);
}
void QBasicTimer_stop(void* self) {
	((QBasicTimer*)self)->stop();
}
void QBasicTimer_delete(void* self) { delete (QBasicTimer*)self; }

}

#include <QtCore/qregularexpression.h>
#include "gen_qregularexpressionmatchiterator.h"

extern "C" {

void* QRegularExpressionMatchIterator_new() {
	return new QRegularExpressionMatchIterator();
}
void* QRegularExpressionMatchIterator_new2(void* iterator) {
	return new QRegularExpressionMatchIterator(*(QRegularExpressionMatchIterator*)iterator);
}
void* QRegularExpressionMatchIterator_OperatorAssign(void* self, void* iterator) {
	auto&& __r = ((QRegularExpressionMatchIterator*)self)->operator=(*(QRegularExpressionMatchIterator*)iterator);
	return (void*)&(__r);
}
void QRegularExpressionMatchIterator_swap(void* self, void* other) {
	((QRegularExpressionMatchIterator*)self)->swap(*(QRegularExpressionMatchIterator*)other);
}
int QRegularExpressionMatchIterator_isValid(void* self) {
	auto&& __r = ((QRegularExpressionMatchIterator*)self)->isValid();
	return __r;
}
int QRegularExpressionMatchIterator_hasNext(void* self) {
	auto&& __r = ((QRegularExpressionMatchIterator*)self)->hasNext();
	return __r;
}
void* QRegularExpressionMatchIterator_next(void* self) {
	auto&& __r = ((QRegularExpressionMatchIterator*)self)->next();
	return new QRegularExpressionMatch(__r);
}
void* QRegularExpressionMatchIterator_peekNext(void* self) {
	auto&& __r = ((QRegularExpressionMatchIterator*)self)->peekNext();
	return new QRegularExpressionMatch(__r);
}
void* QRegularExpressionMatchIterator_regularExpression(void* self) {
	auto&& __r = ((QRegularExpressionMatchIterator*)self)->regularExpression();
	return new QRegularExpression(__r);
}
int QRegularExpressionMatchIterator_matchType(void* self) {
	auto&& __r = ((QRegularExpressionMatchIterator*)self)->matchType();
	return (int)(__r);
}
int QRegularExpressionMatchIterator_matchOptions(void* self) {
	auto&& __r = ((QRegularExpressionMatchIterator*)self)->matchOptions();
	return (int)(__r);
}
void QRegularExpressionMatchIterator_delete(void* self) { delete (QRegularExpressionMatchIterator*)self; }

}

#include <QtCore/qchronotimer.h>
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
#include "gen_qchronotimer.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QChronoTimer__Go : public QChronoTimer {
public:
	long long __goId;
	unsigned char __goOn[8] = {};
	template<typename... A> QChronoTimer__Go(long long id, A&&... a) : QChronoTimer(std::forward<A>(a)...), __goId(id) {}
	using QChronoTimer::sender;
	using QChronoTimer::senderSignalIndex;
	using QChronoTimer::receivers;
	using QChronoTimer::isSignalConnected;
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QChronoTimer::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QChronoTimer::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QChronoTimer::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QChronoTimer::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QChronoTimer::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QChronoTimer::metaObject();
	}
	void timerEvent(QTimerEvent* arg1) override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)arg1, &__h); if (__h) return; }
		QChronoTimer::timerEvent(arg1);
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QChronoTimer::event(event);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QChronoTimer::eventFilter(watched, event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QChronoTimer::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QChronoTimer::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QChronoTimer::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QChronoTimer::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QChronoTimer;
static std::mutex __connsMu_QChronoTimer;

extern "C" {

void* QChronoTimer_base(void* self) { return (void*)static_cast<QObject*>((QChronoTimer*)self); }
void QChronoTimer__set_override(void* self, int idx) { ((QChronoTimer__Go*)self)->__goOn[idx] = 1; }
void* QChronoTimer_fromQObject(void* o) { return (void*)qobject_cast<QChronoTimer*>((QObject*)o); }
void* QChronoTimer_new(long long __goid, long long nsec) {
	return new QChronoTimer__Go(__goid, std::chrono::nanoseconds((long long)nsec));
}
void* QChronoTimer_new2(long long __goid) {
	return new QChronoTimer__Go(__goid);
}
void* QChronoTimer_new3(long long __goid, long long nsec, void* parent) {
	return new QChronoTimer__Go(__goid, std::chrono::nanoseconds((long long)nsec), (QObject*)parent);
}
void* QChronoTimer_new4(long long __goid, void* parent) {
	return new QChronoTimer__Go(__goid, (QObject*)parent);
}
char* QChronoTimer_tr(const char* s, long long* __out_len) {
	auto&& __r = QChronoTimer::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QChronoTimer_isActive(void* self) {
	auto&& __r = ((QChronoTimer*)self)->isActive();
	return __r;
}
int QChronoTimer_id(void* self) {
	auto&& __r = ((QChronoTimer*)self)->id();
	return (int)(__r);
}
void QChronoTimer_setInterval(void* self, long long nsec) {
	((QChronoTimer*)self)->setInterval(std::chrono::nanoseconds((long long)nsec));
}
long long QChronoTimer_interval(void* self) {
	auto&& __r = ((QChronoTimer*)self)->interval();
	return (long long)(__r).count();
}
long long QChronoTimer_remainingTime(void* self) {
	auto&& __r = ((QChronoTimer*)self)->remainingTime();
	return (long long)(__r).count();
}
void QChronoTimer_setTimerType(void* self, int atype) {
	((QChronoTimer*)self)->setTimerType((Qt::TimerType)(atype));
}
int QChronoTimer_timerType(void* self) {
	auto&& __r = ((QChronoTimer*)self)->timerType();
	return (int)(__r);
}
void QChronoTimer_setSingleShot(void* self, int singleShot) {
	((QChronoTimer*)self)->setSingleShot(singleShot);
}
int QChronoTimer_isSingleShot(void* self) {
	auto&& __r = ((QChronoTimer*)self)->isSingleShot();
	return __r;
}
void QChronoTimer_start(void* self) {
	((QChronoTimer*)self)->start();
}
void QChronoTimer_stop(void* self) {
	((QChronoTimer*)self)->stop();
}
void* QChronoTimer_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QChronoTimer::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QChronoTimer_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QChronoTimer::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QChronoTimer_sender(void* self) {
	auto&& __r = ((QChronoTimer__Go*)self)->sender();
	return (void*)(__r);
}
int QChronoTimer_senderSignalIndex(void* self) {
	auto&& __r = ((QChronoTimer__Go*)self)->senderSignalIndex();
	return __r;
}
int QChronoTimer_receivers(void* self, const char* signal) {
	auto&& __r = ((QChronoTimer__Go*)self)->receivers(signal);
	return __r;
}
int QChronoTimer_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QChronoTimer__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QChronoTimer_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QChronoTimer::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QChronoTimer_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QChronoTimer::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QChronoTimer_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QChronoTimer::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QChronoTimer_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QChronoTimer::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QChronoTimer_metaObject_default(void* self) {
	auto&& __r = ((QChronoTimer*)self)->QChronoTimer::metaObject();
	return (void*)(__r);
}
void QChronoTimer_timerEvent_default(void* self, void* arg1) {
	((QChronoTimer__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)arg1);
}
int QChronoTimer_event_default(void* self, void* event) {
	auto&& __r = ((QChronoTimer*)self)->QChronoTimer::event((QEvent*)event);
	return __r;
}
int QChronoTimer_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QChronoTimer*)self)->QChronoTimer::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QChronoTimer_childEvent_default(void* self, void* event) {
	((QChronoTimer__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QChronoTimer_customEvent_default(void* self, void* event) {
	((QChronoTimer__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QChronoTimer_connectNotify_default(void* self, void* signal) {
	((QChronoTimer__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QChronoTimer_disconnectNotify_default(void* self, void* signal) {
	((QChronoTimer__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QChronoTimer_connect_timeout(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QChronoTimer*)self, &QChronoTimer::timeout, (QChronoTimer*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QChronoTimer);
	__conns_QChronoTimer[__id] = __c;
}
void QChronoTimer_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QChronoTimer*)self, &QChronoTimer::destroyed, (QChronoTimer*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QChronoTimer);
	__conns_QChronoTimer[__id] = __c;
}
void QChronoTimer_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QChronoTimer*)self, &QChronoTimer::destroyed, (QChronoTimer*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QChronoTimer);
	__conns_QChronoTimer[__id] = __c;
}
void QChronoTimer_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QChronoTimer*)self, &QChronoTimer::objectNameChanged, (QChronoTimer*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QChronoTimer);
	__conns_QChronoTimer[__id] = __c;
}
void QChronoTimer__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QChronoTimer);
		auto __it = __conns_QChronoTimer.find(id);
		if (__it == __conns_QChronoTimer.end()) return;
		__c = __it->second;
		__conns_QChronoTimer.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QChronoTimer__Go_new(long long id, void* parent) { return new QChronoTimer__Go(id, (QObject*)parent); }
void QChronoTimer_delete(void* self) { delete (QChronoTimer*)self; }

}

#include <QtCore/qfilesystemwatcher.h>
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
#include "gen_qfilesystemwatcher.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QFileSystemWatcher__Go : public QFileSystemWatcher {
public:
	long long __goId;
	unsigned char __goOn[8] = {};
	template<typename... A> QFileSystemWatcher__Go(long long id, A&&... a) : QFileSystemWatcher(std::forward<A>(a)...), __goId(id) {}
	using QFileSystemWatcher::sender;
	using QFileSystemWatcher::senderSignalIndex;
	using QFileSystemWatcher::receivers;
	using QFileSystemWatcher::isSignalConnected;
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QFileSystemWatcher::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QFileSystemWatcher::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QFileSystemWatcher::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QFileSystemWatcher::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QFileSystemWatcher::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QFileSystemWatcher::metaObject();
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QFileSystemWatcher::event(event);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QFileSystemWatcher::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QFileSystemWatcher::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QFileSystemWatcher::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QFileSystemWatcher::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QFileSystemWatcher::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QFileSystemWatcher::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QFileSystemWatcher;
static std::mutex __connsMu_QFileSystemWatcher;

extern "C" {

void* QFileSystemWatcher_base(void* self) { return (void*)static_cast<QObject*>((QFileSystemWatcher*)self); }
void QFileSystemWatcher__set_override(void* self, int idx) { ((QFileSystemWatcher__Go*)self)->__goOn[idx] = 1; }
void* QFileSystemWatcher_fromQObject(void* o) { return (void*)qobject_cast<QFileSystemWatcher*>((QObject*)o); }
void* QFileSystemWatcher_new(long long __goid) {
	return new QFileSystemWatcher__Go(__goid);
}
void* QFileSystemWatcher_new2(long long __goid, char** paths, long long paths_len) {
	QList<QString> paths_list; paths_list.reserve(paths_len);
	for (long long __i = 0; __i < paths_len; __i++) paths_list.append(QString::fromUtf8(paths[__i]));
	return new QFileSystemWatcher__Go(__goid, paths_list);
}
void* QFileSystemWatcher_new3(long long __goid, void* parent) {
	return new QFileSystemWatcher__Go(__goid, (QObject*)parent);
}
void* QFileSystemWatcher_new4(long long __goid, char** paths, long long paths_len, void* parent) {
	QList<QString> paths_list; paths_list.reserve(paths_len);
	for (long long __i = 0; __i < paths_len; __i++) paths_list.append(QString::fromUtf8(paths[__i]));
	return new QFileSystemWatcher__Go(__goid, paths_list, (QObject*)parent);
}
char* QFileSystemWatcher_tr(const char* s, long long* __out_len) {
	auto&& __r = QFileSystemWatcher::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QFileSystemWatcher_addPath(void* self, const char* file, long long file_len) {
	QString file_qstr = QString::fromUtf8(file, file_len);
	auto&& __r = ((QFileSystemWatcher*)self)->addPath(file_qstr);
	return __r;
}
char** QFileSystemWatcher_addPaths(void* self, char** files, long long files_len, long long* __out_len) {
	QList<QString> files_list; files_list.reserve(files_len);
	for (long long __i = 0; __i < files_len; __i++) files_list.append(QString::fromUtf8(files[__i]));
	auto&& __r = ((QFileSystemWatcher*)self)->addPaths(files_list);
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
int QFileSystemWatcher_removePath(void* self, const char* file, long long file_len) {
	QString file_qstr = QString::fromUtf8(file, file_len);
	auto&& __r = ((QFileSystemWatcher*)self)->removePath(file_qstr);
	return __r;
}
char** QFileSystemWatcher_removePaths(void* self, char** files, long long files_len, long long* __out_len) {
	QList<QString> files_list; files_list.reserve(files_len);
	for (long long __i = 0; __i < files_len; __i++) files_list.append(QString::fromUtf8(files[__i]));
	auto&& __r = ((QFileSystemWatcher*)self)->removePaths(files_list);
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
char** QFileSystemWatcher_files(void* self, long long* __out_len) {
	auto&& __r = ((QFileSystemWatcher*)self)->files();
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
char** QFileSystemWatcher_directories(void* self, long long* __out_len) {
	auto&& __r = ((QFileSystemWatcher*)self)->directories();
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
void* QFileSystemWatcher_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QFileSystemWatcher::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QFileSystemWatcher_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QFileSystemWatcher::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QFileSystemWatcher_sender(void* self) {
	auto&& __r = ((QFileSystemWatcher__Go*)self)->sender();
	return (void*)(__r);
}
int QFileSystemWatcher_senderSignalIndex(void* self) {
	auto&& __r = ((QFileSystemWatcher__Go*)self)->senderSignalIndex();
	return __r;
}
int QFileSystemWatcher_receivers(void* self, const char* signal) {
	auto&& __r = ((QFileSystemWatcher__Go*)self)->receivers(signal);
	return __r;
}
int QFileSystemWatcher_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QFileSystemWatcher__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QFileSystemWatcher_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QFileSystemWatcher::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QFileSystemWatcher_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QFileSystemWatcher::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QFileSystemWatcher_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QFileSystemWatcher::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QFileSystemWatcher_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QFileSystemWatcher::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QFileSystemWatcher_metaObject_default(void* self) {
	auto&& __r = ((QFileSystemWatcher*)self)->QFileSystemWatcher::metaObject();
	return (void*)(__r);
}
int QFileSystemWatcher_event_default(void* self, void* event) {
	auto&& __r = ((QFileSystemWatcher*)self)->QFileSystemWatcher::event((QEvent*)event);
	return __r;
}
int QFileSystemWatcher_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QFileSystemWatcher*)self)->QFileSystemWatcher::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QFileSystemWatcher_timerEvent_default(void* self, void* event) {
	((QFileSystemWatcher__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QFileSystemWatcher_childEvent_default(void* self, void* event) {
	((QFileSystemWatcher__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QFileSystemWatcher_customEvent_default(void* self, void* event) {
	((QFileSystemWatcher__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QFileSystemWatcher_connectNotify_default(void* self, void* signal) {
	((QFileSystemWatcher__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QFileSystemWatcher_disconnectNotify_default(void* self, void* signal) {
	((QFileSystemWatcher__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QFileSystemWatcher_connect_fileChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFileSystemWatcher*)self, &QFileSystemWatcher::fileChanged, (QFileSystemWatcher*)self, [__id](const QString& path){
		QByteArray path_ba = path.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)path_ba.constData(), path_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFileSystemWatcher);
	__conns_QFileSystemWatcher[__id] = __c;
}
void QFileSystemWatcher_connect_directoryChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFileSystemWatcher*)self, &QFileSystemWatcher::directoryChanged, (QFileSystemWatcher*)self, [__id](const QString& path){
		QByteArray path_ba = path.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)path_ba.constData(), path_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFileSystemWatcher);
	__conns_QFileSystemWatcher[__id] = __c;
}
void QFileSystemWatcher_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFileSystemWatcher*)self, &QFileSystemWatcher::destroyed, (QFileSystemWatcher*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFileSystemWatcher);
	__conns_QFileSystemWatcher[__id] = __c;
}
void QFileSystemWatcher_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFileSystemWatcher*)self, &QFileSystemWatcher::destroyed, (QFileSystemWatcher*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFileSystemWatcher);
	__conns_QFileSystemWatcher[__id] = __c;
}
void QFileSystemWatcher_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFileSystemWatcher*)self, &QFileSystemWatcher::objectNameChanged, (QFileSystemWatcher*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFileSystemWatcher);
	__conns_QFileSystemWatcher[__id] = __c;
}
void QFileSystemWatcher__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QFileSystemWatcher);
		auto __it = __conns_QFileSystemWatcher.find(id);
		if (__it == __conns_QFileSystemWatcher.end()) return;
		__c = __it->second;
		__conns_QFileSystemWatcher.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QFileSystemWatcher__Go_new(long long id, void* parent) { return new QFileSystemWatcher__Go(id, (QObject*)parent); }
void QFileSystemWatcher_delete(void* self) { delete (QFileSystemWatcher*)self; }

}

#include <QtCore/qline.h>
#include <QtCore/qpoint.h>
#include <QtGui/qtransform.h>
#include "gen_qline.h"

extern "C" {

void* QLine_new() {
	return new QLine();
}
void* QLine_new2(void* pt1, void* pt2) {
	return new QLine(*(QPoint*)pt1, *(QPoint*)pt2);
}
void* QLine_new3(int x1, int y1, int x2, int y2) {
	return new QLine(x1, y1, x2, y2);
}
void* QLine_new4(void* arg1) {
	return new QLine(*(QLine*)arg1);
}
int QLine_isNull(void* self) {
	auto&& __r = ((QLine*)self)->isNull();
	return __r;
}
void* QLine_p1(void* self) {
	auto&& __r = ((QLine*)self)->p1();
	return new QPoint(__r);
}
void* QLine_p2(void* self) {
	auto&& __r = ((QLine*)self)->p2();
	return new QPoint(__r);
}
int QLine_x1(void* self) {
	auto&& __r = ((QLine*)self)->x1();
	return __r;
}
int QLine_y1(void* self) {
	auto&& __r = ((QLine*)self)->y1();
	return __r;
}
int QLine_x2(void* self) {
	auto&& __r = ((QLine*)self)->x2();
	return __r;
}
int QLine_y2(void* self) {
	auto&& __r = ((QLine*)self)->y2();
	return __r;
}
int QLine_dx(void* self) {
	auto&& __r = ((QLine*)self)->dx();
	return __r;
}
int QLine_dy(void* self) {
	auto&& __r = ((QLine*)self)->dy();
	return __r;
}
void QLine_translate(void* self, void* p) {
	((QLine*)self)->translate(*(QPoint*)p);
}
void QLine_translate2(void* self, int dx, int dy) {
	((QLine*)self)->translate(dx, dy);
}
void* QLine_translated(void* self, void* p) {
	auto&& __r = ((QLine*)self)->translated(*(QPoint*)p);
	return new QLine(__r);
}
void* QLine_translated2(void* self, int dx, int dy) {
	auto&& __r = ((QLine*)self)->translated(dx, dy);
	return new QLine(__r);
}
void* QLine_center(void* self) {
	auto&& __r = ((QLine*)self)->center();
	return new QPoint(__r);
}
void QLine_setP1(void* self, void* p1) {
	((QLine*)self)->setP1(*(QPoint*)p1);
}
void QLine_setP2(void* self, void* p2) {
	((QLine*)self)->setP2(*(QPoint*)p2);
}
void QLine_setPoints(void* self, void* p1, void* p2) {
	((QLine*)self)->setPoints(*(QPoint*)p1, *(QPoint*)p2);
}
void QLine_setLine(void* self, int x1, int y1, int x2, int y2) {
	((QLine*)self)->setLine(x1, y1, x2, y2);
}
void* QLine_toLineF(void* self) {
	auto&& __r = ((QLine*)self)->toLineF();
	return new QLineF(__r);
}
int QLine_OperatorEqual(void* self, void* rhs) {
	auto&& __r = (*(QLine*)self) == (*(QLine*)rhs);
	return __r;
}
int QLine_OperatorNotEqual(void* self, void* rhs) {
	auto&& __r = (*(QLine*)self) != (*(QLine*)rhs);
	return __r;
}
int QLine_OperatorEqual2(void* self, void* rhs) {
	auto&& __r = (*(QLine*)self) == (*(QLineF*)rhs);
	return __r;
}
int QLine_OperatorNotEqual2(void* self, void* rhs) {
	auto&& __r = (*(QLine*)self) != (*(QLineF*)rhs);
	return __r;
}
void* QLine_OperatorMultiply(void* self, void* m) {
	auto&& __r = (*(QLine*)self) * (*(QTransform*)m);
	return new QLine(__r);
}
void QLine_delete(void* self) { delete (QLine*)self; }

}

#include <QtCore/qpermissions.h>
#include "gen_qbluetoothpermission.h"

extern "C" {

void* QBluetoothPermission_new() {
	return new QBluetoothPermission();
}
void* QBluetoothPermission_new2(void* other) {
	return new QBluetoothPermission(*(QBluetoothPermission*)other);
}
void QBluetoothPermission_setCommunicationModes(void* self, int modes) {
	((QBluetoothPermission*)self)->setCommunicationModes(QFlags<QBluetoothPermission::CommunicationMode>(modes));
}
int QBluetoothPermission_communicationModes(void* self) {
	auto&& __r = ((QBluetoothPermission*)self)->communicationModes();
	return (int)(__r);
}
void* QBluetoothPermission_OperatorAssign(void* self, void* other) {
	auto&& __r = ((QBluetoothPermission*)self)->operator=(*(QBluetoothPermission*)other);
	return (void*)&(__r);
}
void QBluetoothPermission_swap(void* self, void* other) {
	((QBluetoothPermission*)self)->swap(*(QBluetoothPermission*)other);
}
void QBluetoothPermission_delete(void* self) { delete (QBluetoothPermission*)self; }

}

#include <QtCore/qrect.h>
#include <QtCore/qpoint.h>
#include <QtCore/qsize.h>
#include <QtCore/qmargins.h>
#include "gen_qrectf.h"

extern "C" {

void* QRectF_new() {
	return new QRectF();
}
void* QRectF_new2(void* topleft, void* size) {
	return new QRectF(*(QPointF*)topleft, *(QSizeF*)size);
}
void* QRectF_new3(void* topleft, void* bottomRight) {
	return new QRectF(*(QPointF*)topleft, *(QPointF*)bottomRight);
}
void* QRectF_new4(double left, double top, double width, double height) {
	return new QRectF(left, top, width, height);
}
void* QRectF_new5(void* rect) {
	return new QRectF(*(QRect*)rect);
}
void* QRectF_new6(void* arg1) {
	return new QRectF(*(QRectF*)arg1);
}
int QRectF_isNull(void* self) {
	auto&& __r = ((QRectF*)self)->isNull();
	return __r;
}
int QRectF_isEmpty(void* self) {
	auto&& __r = ((QRectF*)self)->isEmpty();
	return __r;
}
int QRectF_isValid(void* self) {
	auto&& __r = ((QRectF*)self)->isValid();
	return __r;
}
void* QRectF_normalized(void* self) {
	auto&& __r = ((QRectF*)self)->normalized();
	return new QRectF(__r);
}
double QRectF_left(void* self) {
	auto&& __r = ((QRectF*)self)->left();
	return __r;
}
double QRectF_top(void* self) {
	auto&& __r = ((QRectF*)self)->top();
	return __r;
}
double QRectF_right(void* self) {
	auto&& __r = ((QRectF*)self)->right();
	return __r;
}
double QRectF_bottom(void* self) {
	auto&& __r = ((QRectF*)self)->bottom();
	return __r;
}
double QRectF_x(void* self) {
	auto&& __r = ((QRectF*)self)->x();
	return __r;
}
double QRectF_y(void* self) {
	auto&& __r = ((QRectF*)self)->y();
	return __r;
}
void QRectF_setLeft(void* self, double pos) {
	((QRectF*)self)->setLeft(pos);
}
void QRectF_setTop(void* self, double pos) {
	((QRectF*)self)->setTop(pos);
}
void QRectF_setRight(void* self, double pos) {
	((QRectF*)self)->setRight(pos);
}
void QRectF_setBottom(void* self, double pos) {
	((QRectF*)self)->setBottom(pos);
}
void QRectF_setX(void* self, double pos) {
	((QRectF*)self)->setX(pos);
}
void QRectF_setY(void* self, double pos) {
	((QRectF*)self)->setY(pos);
}
void* QRectF_topLeft(void* self) {
	auto&& __r = ((QRectF*)self)->topLeft();
	return new QPointF(__r);
}
void* QRectF_bottomRight(void* self) {
	auto&& __r = ((QRectF*)self)->bottomRight();
	return new QPointF(__r);
}
void* QRectF_topRight(void* self) {
	auto&& __r = ((QRectF*)self)->topRight();
	return new QPointF(__r);
}
void* QRectF_bottomLeft(void* self) {
	auto&& __r = ((QRectF*)self)->bottomLeft();
	return new QPointF(__r);
}
void* QRectF_center(void* self) {
	auto&& __r = ((QRectF*)self)->center();
	return new QPointF(__r);
}
void QRectF_setTopLeft(void* self, void* p) {
	((QRectF*)self)->setTopLeft(*(QPointF*)p);
}
void QRectF_setBottomRight(void* self, void* p) {
	((QRectF*)self)->setBottomRight(*(QPointF*)p);
}
void QRectF_setTopRight(void* self, void* p) {
	((QRectF*)self)->setTopRight(*(QPointF*)p);
}
void QRectF_setBottomLeft(void* self, void* p) {
	((QRectF*)self)->setBottomLeft(*(QPointF*)p);
}
void QRectF_moveLeft(void* self, double pos) {
	((QRectF*)self)->moveLeft(pos);
}
void QRectF_moveTop(void* self, double pos) {
	((QRectF*)self)->moveTop(pos);
}
void QRectF_moveRight(void* self, double pos) {
	((QRectF*)self)->moveRight(pos);
}
void QRectF_moveBottom(void* self, double pos) {
	((QRectF*)self)->moveBottom(pos);
}
void QRectF_moveTopLeft(void* self, void* p) {
	((QRectF*)self)->moveTopLeft(*(QPointF*)p);
}
void QRectF_moveBottomRight(void* self, void* p) {
	((QRectF*)self)->moveBottomRight(*(QPointF*)p);
}
void QRectF_moveTopRight(void* self, void* p) {
	((QRectF*)self)->moveTopRight(*(QPointF*)p);
}
void QRectF_moveBottomLeft(void* self, void* p) {
	((QRectF*)self)->moveBottomLeft(*(QPointF*)p);
}
void QRectF_moveCenter(void* self, void* p) {
	((QRectF*)self)->moveCenter(*(QPointF*)p);
}
void QRectF_translate(void* self, double dx, double dy) {
	((QRectF*)self)->translate(dx, dy);
}
void QRectF_translate2(void* self, void* p) {
	((QRectF*)self)->translate(*(QPointF*)p);
}
void* QRectF_translated(void* self, double dx, double dy) {
	auto&& __r = ((QRectF*)self)->translated(dx, dy);
	return new QRectF(__r);
}
void* QRectF_translated2(void* self, void* p) {
	auto&& __r = ((QRectF*)self)->translated(*(QPointF*)p);
	return new QRectF(__r);
}
void* QRectF_transposed(void* self) {
	auto&& __r = ((QRectF*)self)->transposed();
	return new QRectF(__r);
}
void QRectF_moveTo(void* self, double x, double y) {
	((QRectF*)self)->moveTo(x, y);
}
void QRectF_moveTo2(void* self, void* p) {
	((QRectF*)self)->moveTo(*(QPointF*)p);
}
void QRectF_setRect(void* self, double x, double y, double w, double h) {
	((QRectF*)self)->setRect(x, y, w, h);
}
void QRectF_getRect(void* self, double* x, double* y, double* w, double* h) {
	qreal x_local{};
	qreal y_local{};
	qreal w_local{};
	qreal h_local{};
	((QRectF*)self)->getRect(&x_local, &y_local, &w_local, &h_local);
	*x = x_local;
	*y = y_local;
	*w = w_local;
	*h = h_local;
}
void QRectF_setCoords(void* self, double x1, double y1, double x2, double y2) {
	((QRectF*)self)->setCoords(x1, y1, x2, y2);
}
void QRectF_getCoords(void* self, double* x1, double* y1, double* x2, double* y2) {
	qreal x1_local{};
	qreal y1_local{};
	qreal x2_local{};
	qreal y2_local{};
	((QRectF*)self)->getCoords(&x1_local, &y1_local, &x2_local, &y2_local);
	*x1 = x1_local;
	*y1 = y1_local;
	*x2 = x2_local;
	*y2 = y2_local;
}
void QRectF_adjust(void* self, double x1, double y1, double x2, double y2) {
	((QRectF*)self)->adjust(x1, y1, x2, y2);
}
void* QRectF_adjusted(void* self, double x1, double y1, double x2, double y2) {
	auto&& __r = ((QRectF*)self)->adjusted(x1, y1, x2, y2);
	return new QRectF(__r);
}
void* QRectF_size(void* self) {
	auto&& __r = ((QRectF*)self)->size();
	return new QSizeF(__r);
}
double QRectF_width(void* self) {
	auto&& __r = ((QRectF*)self)->width();
	return __r;
}
double QRectF_height(void* self) {
	auto&& __r = ((QRectF*)self)->height();
	return __r;
}
void QRectF_setWidth(void* self, double w) {
	((QRectF*)self)->setWidth(w);
}
void QRectF_setHeight(void* self, double h) {
	((QRectF*)self)->setHeight(h);
}
void QRectF_setSize(void* self, void* s) {
	((QRectF*)self)->setSize(*(QSizeF*)s);
}
void* QRectF_OperatorBitwiseOr(void* self, void* r) {
	auto&& __r = ((QRectF*)self)->operator|(*(QRectF*)r);
	return new QRectF(__r);
}
void* QRectF_OperatorBitwiseAnd(void* self, void* r) {
	auto&& __r = ((QRectF*)self)->operator&(*(QRectF*)r);
	return new QRectF(__r);
}
void* QRectF_OperatorBitwiseOrAssign(void* self, void* r) {
	auto&& __r = ((QRectF*)self)->operator|=(*(QRectF*)r);
	return (void*)&(__r);
}
void* QRectF_OperatorBitwiseAndAssign(void* self, void* r) {
	auto&& __r = ((QRectF*)self)->operator&=(*(QRectF*)r);
	return (void*)&(__r);
}
int QRectF_contains(void* self, void* r) {
	auto&& __r = ((QRectF*)self)->contains(*(QRectF*)r);
	return __r;
}
int QRectF_contains2(void* self, void* p) {
	auto&& __r = ((QRectF*)self)->contains(*(QPointF*)p);
	return __r;
}
int QRectF_contains3(void* self, double x, double y) {
	auto&& __r = ((QRectF*)self)->contains(x, y);
	return __r;
}
void* QRectF_united(void* self, void* other) {
	auto&& __r = ((QRectF*)self)->united(*(QRectF*)other);
	return new QRectF(__r);
}
void* QRectF_intersected(void* self, void* other) {
	auto&& __r = ((QRectF*)self)->intersected(*(QRectF*)other);
	return new QRectF(__r);
}
int QRectF_intersects(void* self, void* r) {
	auto&& __r = ((QRectF*)self)->intersects(*(QRectF*)r);
	return __r;
}
void* QRectF_marginsAdded(void* self, void* margins) {
	auto&& __r = ((QRectF*)self)->marginsAdded(*(QMarginsF*)margins);
	return new QRectF(__r);
}
void* QRectF_marginsRemoved(void* self, void* margins) {
	auto&& __r = ((QRectF*)self)->marginsRemoved(*(QMarginsF*)margins);
	return new QRectF(__r);
}
void* QRectF_OperatorPlusAssign(void* self, void* margins) {
	auto&& __r = ((QRectF*)self)->operator+=(*(QMarginsF*)margins);
	return (void*)&(__r);
}
void* QRectF_OperatorMinusAssign(void* self, void* margins) {
	auto&& __r = ((QRectF*)self)->operator-=(*(QMarginsF*)margins);
	return (void*)&(__r);
}
void* QRectF_toRect(void* self) {
	auto&& __r = ((QRectF*)self)->toRect();
	return new QRect(__r);
}
void* QRectF_toAlignedRect(void* self) {
	auto&& __r = ((QRectF*)self)->toAlignedRect();
	return new QRect(__r);
}
void* QRectF_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QRectF*)self)->operator=(*(QRectF*)arg1);
	return (void*)&(__r);
}
int QRectF_OperatorEqual(void* self, void* rhs) {
	auto&& __r = (*(QRectF*)self) == (*(QRectF*)rhs);
	return __r;
}
int QRectF_OperatorNotEqual(void* self, void* rhs) {
	auto&& __r = (*(QRectF*)self) != (*(QRectF*)rhs);
	return __r;
}
int QRectF_OperatorEqual2(void* self, void* rhs) {
	auto&& __r = (*(QRectF*)self) == (*(QRect*)rhs);
	return __r;
}
int QRectF_OperatorNotEqual2(void* self, void* rhs) {
	auto&& __r = (*(QRectF*)self) != (*(QRect*)rhs);
	return __r;
}
void* QRectF_OperatorPlus(void* self, void* rhs) {
	auto&& __r = (*(QRectF*)self) + (*(QMarginsF*)rhs);
	return new QRectF(__r);
}
void* QRectF_OperatorMinus(void* self, void* rhs) {
	auto&& __r = (*(QRectF*)self) - (*(QMarginsF*)rhs);
	return new QRectF(__r);
}
void QRectF_delete(void* self) { delete (QRectF*)self; }

}

#include <QtCore/qtimeline.h>
#include <QtCore/qobject.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qeasingcurve.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qvariant.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qtimeline.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QTimeLine__Go : public QTimeLine {
public:
	long long __goId;
	unsigned char __goOn[9] = {};
	template<typename... A> QTimeLine__Go(long long id, A&&... a) : QTimeLine(std::forward<A>(a)...), __goId(id) {}
	using QTimeLine::sender;
	using QTimeLine::senderSignalIndex;
	using QTimeLine::receivers;
	using QTimeLine::isSignalConnected;
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QTimeLine::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QTimeLine::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QTimeLine::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QTimeLine::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QTimeLine::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QTimeLine::metaObject();
	}
	qreal valueForTime(int msec) const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_double_int) { unsigned long long __b = ((unsigned long long(*)(long long, long long, int, int*))cppgo_virt_double_int)(__goId, 3548770768539434384LL, msec, &__h); if (__h) { double __r; __builtin_memcpy(&__r, &__b, sizeof(__r)); return __r; } }
		return QTimeLine::valueForTime(msec);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QTimeLine::timerEvent(event);
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QTimeLine::event(event);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QTimeLine::eventFilter(watched, event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QTimeLine::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QTimeLine::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QTimeLine::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QTimeLine::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QTimeLine;
static std::mutex __connsMu_QTimeLine;

extern "C" {

void* QTimeLine_base(void* self) { return (void*)static_cast<QObject*>((QTimeLine*)self); }
void QTimeLine__set_override(void* self, int idx) { ((QTimeLine__Go*)self)->__goOn[idx] = 1; }
void* QTimeLine_fromQObject(void* o) { return (void*)qobject_cast<QTimeLine*>((QObject*)o); }
void* QTimeLine_new(long long __goid) {
	return new QTimeLine__Go(__goid);
}
void* QTimeLine_new2(long long __goid, int duration) {
	return new QTimeLine__Go(__goid, duration);
}
void* QTimeLine_new3(long long __goid, int duration, void* parent) {
	return new QTimeLine__Go(__goid, duration, (QObject*)parent);
}
char* QTimeLine_tr(const char* s, long long* __out_len) {
	auto&& __r = QTimeLine::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QTimeLine_state(void* self) {
	auto&& __r = ((QTimeLine*)self)->state();
	return (int)(__r);
}
int QTimeLine_loopCount(void* self) {
	auto&& __r = ((QTimeLine*)self)->loopCount();
	return __r;
}
void QTimeLine_setLoopCount(void* self, int count) {
	((QTimeLine*)self)->setLoopCount(count);
}
int QTimeLine_direction(void* self) {
	auto&& __r = ((QTimeLine*)self)->direction();
	return (int)(__r);
}
void QTimeLine_setDirection(void* self, int direction) {
	((QTimeLine*)self)->setDirection((QTimeLine::Direction)(direction));
}
int QTimeLine_duration(void* self) {
	auto&& __r = ((QTimeLine*)self)->duration();
	return __r;
}
void QTimeLine_setDuration(void* self, int duration) {
	((QTimeLine*)self)->setDuration(duration);
}
int QTimeLine_startFrame(void* self) {
	auto&& __r = ((QTimeLine*)self)->startFrame();
	return __r;
}
void QTimeLine_setStartFrame(void* self, int frame) {
	((QTimeLine*)self)->setStartFrame(frame);
}
int QTimeLine_endFrame(void* self) {
	auto&& __r = ((QTimeLine*)self)->endFrame();
	return __r;
}
void QTimeLine_setEndFrame(void* self, int frame) {
	((QTimeLine*)self)->setEndFrame(frame);
}
void QTimeLine_setFrameRange(void* self, int startFrame, int endFrame) {
	((QTimeLine*)self)->setFrameRange(startFrame, endFrame);
}
int QTimeLine_updateInterval(void* self) {
	auto&& __r = ((QTimeLine*)self)->updateInterval();
	return __r;
}
void QTimeLine_setUpdateInterval(void* self, int interval) {
	((QTimeLine*)self)->setUpdateInterval(interval);
}
void* QTimeLine_easingCurve(void* self) {
	auto&& __r = ((QTimeLine*)self)->easingCurve();
	return new QEasingCurve(__r);
}
void QTimeLine_setEasingCurve(void* self, void* curve) {
	((QTimeLine*)self)->setEasingCurve(*(QEasingCurve*)curve);
}
int QTimeLine_currentTime(void* self) {
	auto&& __r = ((QTimeLine*)self)->currentTime();
	return __r;
}
int QTimeLine_currentFrame(void* self) {
	auto&& __r = ((QTimeLine*)self)->currentFrame();
	return __r;
}
double QTimeLine_currentValue(void* self) {
	auto&& __r = ((QTimeLine*)self)->currentValue();
	return __r;
}
int QTimeLine_frameForTime(void* self, int msec) {
	auto&& __r = ((QTimeLine*)self)->frameForTime(msec);
	return __r;
}
double QTimeLine_valueForTime(void* self, int msec) {
	auto&& __r = ((QTimeLine*)self)->valueForTime(msec);
	return __r;
}
void QTimeLine_start(void* self) {
	((QTimeLine*)self)->start();
}
void QTimeLine_resume(void* self) {
	((QTimeLine*)self)->resume();
}
void QTimeLine_stop(void* self) {
	((QTimeLine*)self)->stop();
}
void QTimeLine_setPaused(void* self, int paused) {
	((QTimeLine*)self)->setPaused(paused);
}
void QTimeLine_setCurrentTime(void* self, int msec) {
	((QTimeLine*)self)->setCurrentTime(msec);
}
void QTimeLine_toggleDirection(void* self) {
	((QTimeLine*)self)->toggleDirection();
}
void* QTimeLine_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QTimeLine::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QTimeLine_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QTimeLine::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QTimeLine_sender(void* self) {
	auto&& __r = ((QTimeLine__Go*)self)->sender();
	return (void*)(__r);
}
int QTimeLine_senderSignalIndex(void* self) {
	auto&& __r = ((QTimeLine__Go*)self)->senderSignalIndex();
	return __r;
}
int QTimeLine_receivers(void* self, const char* signal) {
	auto&& __r = ((QTimeLine__Go*)self)->receivers(signal);
	return __r;
}
int QTimeLine_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QTimeLine__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QTimeLine_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QTimeLine::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QTimeLine_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QTimeLine::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QTimeLine_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QTimeLine::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QTimeLine_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QTimeLine::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
double QTimeLine_valueForTime_default(void* self, int msec) {
	auto&& __r = ((QTimeLine*)self)->QTimeLine::valueForTime(msec);
	return __r;
}
void* QTimeLine_metaObject_default(void* self) {
	auto&& __r = ((QTimeLine*)self)->QTimeLine::metaObject();
	return (void*)(__r);
}
void QTimeLine_timerEvent_default(void* self, void* event) {
	((QTimeLine__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
int QTimeLine_event_default(void* self, void* event) {
	auto&& __r = ((QTimeLine*)self)->QTimeLine::event((QEvent*)event);
	return __r;
}
int QTimeLine_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QTimeLine*)self)->QTimeLine::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QTimeLine_childEvent_default(void* self, void* event) {
	((QTimeLine__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QTimeLine_customEvent_default(void* self, void* event) {
	((QTimeLine__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QTimeLine_connectNotify_default(void* self, void* signal) {
	((QTimeLine__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QTimeLine_disconnectNotify_default(void* self, void* signal) {
	((QTimeLine__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QTimeLine_connect_valueChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTimeLine*)self, &QTimeLine::valueChanged, (QTimeLine*)self, [__id](qreal x){
		if (cppgo_sig_void_unsignedlonglong) ((void(*)(long long, unsigned long long))cppgo_sig_void_unsignedlonglong)(__id, __builtin_bit_cast(unsigned long long, (double)(x)));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTimeLine);
	__conns_QTimeLine[__id] = __c;
}
void QTimeLine_connect_frameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTimeLine*)self, &QTimeLine::frameChanged, (QTimeLine*)self, [__id](int arg1){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTimeLine);
	__conns_QTimeLine[__id] = __c;
}
void QTimeLine_connect_stateChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTimeLine*)self, &QTimeLine::stateChanged, (QTimeLine*)self, [__id](QTimeLine::State newState){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, (int)(newState));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTimeLine);
	__conns_QTimeLine[__id] = __c;
}
void QTimeLine_connect_finished(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTimeLine*)self, &QTimeLine::finished, (QTimeLine*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTimeLine);
	__conns_QTimeLine[__id] = __c;
}
void QTimeLine_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTimeLine*)self, &QTimeLine::destroyed, (QTimeLine*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTimeLine);
	__conns_QTimeLine[__id] = __c;
}
void QTimeLine_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTimeLine*)self, &QTimeLine::destroyed, (QTimeLine*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTimeLine);
	__conns_QTimeLine[__id] = __c;
}
void QTimeLine_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTimeLine*)self, &QTimeLine::objectNameChanged, (QTimeLine*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTimeLine);
	__conns_QTimeLine[__id] = __c;
}
void QTimeLine__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QTimeLine);
		auto __it = __conns_QTimeLine.find(id);
		if (__it == __conns_QTimeLine.end()) return;
		__c = __it->second;
		__conns_QTimeLine.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QTimeLine__Go_new(long long id) { return new QTimeLine__Go(id); }
void QTimeLine_delete(void* self) { delete (QTimeLine*)self; }

}

#include <QtGui/qicon.h>
#include <QtGui/qpixmap.h>
#include <QtCore/qstring.h>
#include <QtGui/qiconengine.h>
#include <QtCore/qvariant.h>
#include <QtCore/qsize.h>
#include <QtGui/qwindow.h>
#include <QtGui/qpainter.h>
#include <QtCore/qrect.h>
#include <QtCore/qnamespace.h>
#include "gen_qicon.h"

extern "C" {

void* QIcon_new() {
	return new QIcon();
}
void* QIcon_new2(void* pixmap) {
	return new QIcon(*(QPixmap*)pixmap);
}
void* QIcon_new3(void* other) {
	return new QIcon(*(QIcon*)other);
}
void* QIcon_new4(const char* fileName, long long fileName_len) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	return new QIcon(fileName_qstr);
}
void* QIcon_new5(void* engine) {
	return new QIcon((QIconEngine*)engine);
}
void* QIcon_OperatorAssign(void* self, void* other) {
	auto&& __r = ((QIcon*)self)->operator=(*(QIcon*)other);
	return (void*)&(__r);
}
void QIcon_swap(void* self, void* other) {
	((QIcon*)self)->swap(*(QIcon*)other);
}
void* QIcon_ToQVariant(void* self) {
	auto&& __r = static_cast<QVariant>(*(QIcon*)self);
	return new QVariant(__r);
}
void* QIcon_pixmap(void* self, void* size) {
	auto&& __r = ((QIcon*)self)->pixmap(*(QSize*)size);
	return new QPixmap(__r);
}
void* QIcon_pixmap2(void* self, int w, int h) {
	auto&& __r = ((QIcon*)self)->pixmap(w, h);
	return new QPixmap(__r);
}
void* QIcon_pixmap3(void* self, int extent) {
	auto&& __r = ((QIcon*)self)->pixmap(extent);
	return new QPixmap(__r);
}
void* QIcon_pixmap4(void* self, void* size, double devicePixelRatio) {
	auto&& __r = ((QIcon*)self)->pixmap(*(QSize*)size, devicePixelRatio);
	return new QPixmap(__r);
}
void* QIcon_pixmap5(void* self, void* window, void* size) {
	auto&& __r = ((QIcon*)self)->pixmap((QWindow*)window, *(QSize*)size);
	return new QPixmap(__r);
}
void* QIcon_actualSize(void* self, void* size) {
	auto&& __r = ((QIcon*)self)->actualSize(*(QSize*)size);
	return new QSize(__r);
}
void* QIcon_actualSize2(void* self, void* window, void* size) {
	auto&& __r = ((QIcon*)self)->actualSize((QWindow*)window, *(QSize*)size);
	return new QSize(__r);
}
char* QIcon_name(void* self, long long* __out_len) {
	auto&& __r = ((QIcon*)self)->name();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QIcon_paint(void* self, void* painter, void* rect) {
	((QIcon*)self)->paint((QPainter*)painter, *(QRect*)rect);
}
void QIcon_paint2(void* self, void* painter, int x, int y, int w, int h) {
	((QIcon*)self)->paint((QPainter*)painter, x, y, w, h);
}
int QIcon_isNull(void* self) {
	auto&& __r = ((QIcon*)self)->isNull();
	return __r;
}
int QIcon_isDetached(void* self) {
	auto&& __r = ((QIcon*)self)->isDetached();
	return __r;
}
void QIcon_detach(void* self) {
	((QIcon*)self)->detach();
}
long long QIcon_cacheKey(void* self) {
	auto&& __r = ((QIcon*)self)->cacheKey();
	return __r;
}
void QIcon_addPixmap(void* self, void* pixmap) {
	((QIcon*)self)->addPixmap(*(QPixmap*)pixmap);
}
void QIcon_addFile(void* self, const char* fileName, long long fileName_len) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	((QIcon*)self)->addFile(fileName_qstr);
}
void** QIcon_availableSizes(void* self, long long* __out_len) {
	auto&& __r = ((QIcon*)self)->availableSizes();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QSize(__e);
	return __arr;
}
void QIcon_setIsMask(void* self, int isMask) {
	((QIcon*)self)->setIsMask(isMask);
}
int QIcon_isMask(void* self) {
	auto&& __r = ((QIcon*)self)->isMask();
	return __r;
}
void* QIcon_fromTheme(const char* name, long long name_len) {
	QString name_qstr = QString::fromUtf8(name, name_len);
	auto&& __r = QIcon::fromTheme(name_qstr);
	return new QIcon(__r);
}
void* QIcon_fromTheme2(const char* name, long long name_len, void* fallback) {
	QString name_qstr = QString::fromUtf8(name, name_len);
	auto&& __r = QIcon::fromTheme(name_qstr, *(QIcon*)fallback);
	return new QIcon(__r);
}
int QIcon_hasThemeIcon(const char* name, long long name_len) {
	QString name_qstr = QString::fromUtf8(name, name_len);
	auto&& __r = QIcon::hasThemeIcon(name_qstr);
	return __r;
}
void* QIcon_fromTheme3(int icon) {
	auto&& __r = QIcon::fromTheme((QIcon::ThemeIcon)(icon));
	return new QIcon(__r);
}
void* QIcon_fromTheme4(int icon, void* fallback) {
	auto&& __r = QIcon::fromTheme((QIcon::ThemeIcon)(icon), *(QIcon*)fallback);
	return new QIcon(__r);
}
int QIcon_hasThemeIcon2(int icon) {
	auto&& __r = QIcon::hasThemeIcon((QIcon::ThemeIcon)(icon));
	return __r;
}
char** QIcon_themeSearchPaths(long long* __out_len) {
	auto&& __r = QIcon::themeSearchPaths();
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
void QIcon_setThemeSearchPaths(char** searchpath, long long searchpath_len) {
	QList<QString> searchpath_list; searchpath_list.reserve(searchpath_len);
	for (long long __i = 0; __i < searchpath_len; __i++) searchpath_list.append(QString::fromUtf8(searchpath[__i]));
	QIcon::setThemeSearchPaths(searchpath_list);
}
char** QIcon_fallbackSearchPaths(long long* __out_len) {
	auto&& __r = QIcon::fallbackSearchPaths();
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
void QIcon_setFallbackSearchPaths(char** paths, long long paths_len) {
	QList<QString> paths_list; paths_list.reserve(paths_len);
	for (long long __i = 0; __i < paths_len; __i++) paths_list.append(QString::fromUtf8(paths[__i]));
	QIcon::setFallbackSearchPaths(paths_list);
}
char* QIcon_themeName(long long* __out_len) {
	auto&& __r = QIcon::themeName();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QIcon_setThemeName(const char* path, long long path_len) {
	QString path_qstr = QString::fromUtf8(path, path_len);
	QIcon::setThemeName(path_qstr);
}
char* QIcon_fallbackThemeName(long long* __out_len) {
	auto&& __r = QIcon::fallbackThemeName();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QIcon_setFallbackThemeName(const char* name, long long name_len) {
	QString name_qstr = QString::fromUtf8(name, name_len);
	QIcon::setFallbackThemeName(name_qstr);
}
void* QIcon_pixmap6(void* self, void* size, int mode) {
	auto&& __r = ((QIcon*)self)->pixmap(*(QSize*)size, (QIcon::Mode)(mode));
	return new QPixmap(__r);
}
void* QIcon_pixmap7(void* self, void* size, int mode, int state) {
	auto&& __r = ((QIcon*)self)->pixmap(*(QSize*)size, (QIcon::Mode)(mode), (QIcon::State)(state));
	return new QPixmap(__r);
}
void* QIcon_pixmap8(void* self, int w, int h, int mode) {
	auto&& __r = ((QIcon*)self)->pixmap(w, h, (QIcon::Mode)(mode));
	return new QPixmap(__r);
}
void* QIcon_pixmap9(void* self, int w, int h, int mode, int state) {
	auto&& __r = ((QIcon*)self)->pixmap(w, h, (QIcon::Mode)(mode), (QIcon::State)(state));
	return new QPixmap(__r);
}
void* QIcon_pixmap10(void* self, int extent, int mode) {
	auto&& __r = ((QIcon*)self)->pixmap(extent, (QIcon::Mode)(mode));
	return new QPixmap(__r);
}
void* QIcon_pixmap11(void* self, int extent, int mode, int state) {
	auto&& __r = ((QIcon*)self)->pixmap(extent, (QIcon::Mode)(mode), (QIcon::State)(state));
	return new QPixmap(__r);
}
void* QIcon_pixmap12(void* self, void* size, double devicePixelRatio, int mode) {
	auto&& __r = ((QIcon*)self)->pixmap(*(QSize*)size, devicePixelRatio, (QIcon::Mode)(mode));
	return new QPixmap(__r);
}
void* QIcon_pixmap13(void* self, void* size, double devicePixelRatio, int mode, int state) {
	auto&& __r = ((QIcon*)self)->pixmap(*(QSize*)size, devicePixelRatio, (QIcon::Mode)(mode), (QIcon::State)(state));
	return new QPixmap(__r);
}
void* QIcon_pixmap14(void* self, void* window, void* size, int mode) {
	auto&& __r = ((QIcon*)self)->pixmap((QWindow*)window, *(QSize*)size, (QIcon::Mode)(mode));
	return new QPixmap(__r);
}
void* QIcon_pixmap15(void* self, void* window, void* size, int mode, int state) {
	auto&& __r = ((QIcon*)self)->pixmap((QWindow*)window, *(QSize*)size, (QIcon::Mode)(mode), (QIcon::State)(state));
	return new QPixmap(__r);
}
void* QIcon_actualSize3(void* self, void* size, int mode) {
	auto&& __r = ((QIcon*)self)->actualSize(*(QSize*)size, (QIcon::Mode)(mode));
	return new QSize(__r);
}
void* QIcon_actualSize4(void* self, void* size, int mode, int state) {
	auto&& __r = ((QIcon*)self)->actualSize(*(QSize*)size, (QIcon::Mode)(mode), (QIcon::State)(state));
	return new QSize(__r);
}
void* QIcon_actualSize5(void* self, void* window, void* size, int mode) {
	auto&& __r = ((QIcon*)self)->actualSize((QWindow*)window, *(QSize*)size, (QIcon::Mode)(mode));
	return new QSize(__r);
}
void* QIcon_actualSize6(void* self, void* window, void* size, int mode, int state) {
	auto&& __r = ((QIcon*)self)->actualSize((QWindow*)window, *(QSize*)size, (QIcon::Mode)(mode), (QIcon::State)(state));
	return new QSize(__r);
}
void QIcon_paint3(void* self, void* painter, void* rect, int alignment) {
	((QIcon*)self)->paint((QPainter*)painter, *(QRect*)rect, QFlags<Qt::AlignmentFlag>(alignment));
}
void QIcon_paint4(void* self, void* painter, void* rect, int alignment, int mode) {
	((QIcon*)self)->paint((QPainter*)painter, *(QRect*)rect, QFlags<Qt::AlignmentFlag>(alignment), (QIcon::Mode)(mode));
}
void QIcon_paint5(void* self, void* painter, void* rect, int alignment, int mode, int state) {
	((QIcon*)self)->paint((QPainter*)painter, *(QRect*)rect, QFlags<Qt::AlignmentFlag>(alignment), (QIcon::Mode)(mode), (QIcon::State)(state));
}
void QIcon_paint6(void* self, void* painter, int x, int y, int w, int h, int alignment) {
	((QIcon*)self)->paint((QPainter*)painter, x, y, w, h, QFlags<Qt::AlignmentFlag>(alignment));
}
void QIcon_paint7(void* self, void* painter, int x, int y, int w, int h, int alignment, int mode) {
	((QIcon*)self)->paint((QPainter*)painter, x, y, w, h, QFlags<Qt::AlignmentFlag>(alignment), (QIcon::Mode)(mode));
}
void QIcon_paint8(void* self, void* painter, int x, int y, int w, int h, int alignment, int mode, int state) {
	((QIcon*)self)->paint((QPainter*)painter, x, y, w, h, QFlags<Qt::AlignmentFlag>(alignment), (QIcon::Mode)(mode), (QIcon::State)(state));
}
void QIcon_addPixmap2(void* self, void* pixmap, int mode) {
	((QIcon*)self)->addPixmap(*(QPixmap*)pixmap, (QIcon::Mode)(mode));
}
void QIcon_addPixmap3(void* self, void* pixmap, int mode, int state) {
	((QIcon*)self)->addPixmap(*(QPixmap*)pixmap, (QIcon::Mode)(mode), (QIcon::State)(state));
}
void QIcon_addFile2(void* self, const char* fileName, long long fileName_len, void* size) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	((QIcon*)self)->addFile(fileName_qstr, *(QSize*)size);
}
void QIcon_addFile3(void* self, const char* fileName, long long fileName_len, void* size, int mode) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	((QIcon*)self)->addFile(fileName_qstr, *(QSize*)size, (QIcon::Mode)(mode));
}
void QIcon_addFile4(void* self, const char* fileName, long long fileName_len, void* size, int mode, int state) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	((QIcon*)self)->addFile(fileName_qstr, *(QSize*)size, (QIcon::Mode)(mode), (QIcon::State)(state));
}
void** QIcon_availableSizes2(void* self, int mode, long long* __out_len) {
	auto&& __r = ((QIcon*)self)->availableSizes((QIcon::Mode)(mode));
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QSize(__e);
	return __arr;
}
void** QIcon_availableSizes3(void* self, int mode, int state, long long* __out_len) {
	auto&& __r = ((QIcon*)self)->availableSizes((QIcon::Mode)(mode), (QIcon::State)(state));
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QSize(__e);
	return __arr;
}
void* QIcon_toQVariant(void* self) {
	auto&& __r = QVariant::fromValue(*(QIcon*)self);
	return new QVariant(__r);
}
void QIcon_delete(void* self) { delete (QIcon*)self; }

}

#include <QtGui/qevent.h>
#include <QtCore/qcoreevent.h>
#include <QtGui/qregion.h>
#include <QtCore/qrect.h>
#include "gen_qpaintevent.h"

extern "C" {

void* QPaintEvent_base(void* self) { return (void*)static_cast<QEvent*>((QPaintEvent*)self); }
void* QPaintEvent_new(void* paintRegion) {
	return new QPaintEvent(*(QRegion*)paintRegion);
}
void* QPaintEvent_new2(void* paintRect) {
	return new QPaintEvent(*(QRect*)paintRect);
}
void* QPaintEvent_rect(void* self) {
	auto&& __r = ((QPaintEvent*)self)->rect();
	return (void*)&(__r);
}
void* QPaintEvent_region(void* self) {
	auto&& __r = ((QPaintEvent*)self)->region();
	return (void*)&(__r);
}
int QPaintEvent_registerEventType() {
	auto&& __r = QPaintEvent::registerEventType();
	return __r;
}
int QPaintEvent_registerEventType2(int hint) {
	auto&& __r = QPaintEvent::registerEventType(hint);
	return __r;
}
void QPaintEvent_delete(void* self) { delete (QPaintEvent*)self; }

}

#include <QtGui/qevent.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qnamespace.h>
#include "gen_qwindowstatechangeevent.h"

extern "C" {

void* QWindowStateChangeEvent_base(void* self) { return (void*)static_cast<QEvent*>((QWindowStateChangeEvent*)self); }
void* QWindowStateChangeEvent_new(int oldState) {
	return new QWindowStateChangeEvent(QFlags<Qt::WindowState>(oldState));
}
void* QWindowStateChangeEvent_new2(int oldState, int isOverride) {
	return new QWindowStateChangeEvent(QFlags<Qt::WindowState>(oldState), isOverride);
}
int QWindowStateChangeEvent_oldState(void* self) {
	auto&& __r = ((QWindowStateChangeEvent*)self)->oldState();
	return (int)(__r);
}
int QWindowStateChangeEvent_isOverride(void* self) {
	auto&& __r = ((QWindowStateChangeEvent*)self)->isOverride();
	return __r;
}
int QWindowStateChangeEvent_registerEventType() {
	auto&& __r = QWindowStateChangeEvent::registerEventType();
	return __r;
}
int QWindowStateChangeEvent_registerEventType2(int hint) {
	auto&& __r = QWindowStateChangeEvent::registerEventType(hint);
	return __r;
}
void QWindowStateChangeEvent_delete(void* self) { delete (QWindowStateChangeEvent*)self; }

}

#include <QtGui/qtextformat.h>
#include <QtCore/qvariant.h>
#include "gen_qtextlength.h"

extern "C" {

void* QTextLength_new() {
	return new QTextLength();
}
void* QTextLength_new2(int type_, double value) {
	return new QTextLength((QTextLength::Type)(type_), value);
}
void* QTextLength_new3(void* arg1) {
	return new QTextLength(*(QTextLength*)arg1);
}
int QTextLength_type(void* self) {
	auto&& __r = ((QTextLength*)self)->type();
	return (int)(__r);
}
double QTextLength_value(void* self, double maximumLength) {
	auto&& __r = ((QTextLength*)self)->value(maximumLength);
	return __r;
}
double QTextLength_rawValue(void* self) {
	auto&& __r = ((QTextLength*)self)->rawValue();
	return __r;
}
int QTextLength_OperatorEqual(void* self, void* other) {
	auto&& __r = ((QTextLength*)self)->operator==(*(QTextLength*)other);
	return __r;
}
int QTextLength_OperatorNotEqual(void* self, void* other) {
	auto&& __r = ((QTextLength*)self)->operator!=(*(QTextLength*)other);
	return __r;
}
void* QTextLength_ToQVariant(void* self) {
	auto&& __r = static_cast<QVariant>(*(QTextLength*)self);
	return new QVariant(__r);
}
void* QTextLength_toQVariant(void* self) {
	auto&& __r = QVariant::fromValue(*(QTextLength*)self);
	return new QVariant(__r);
}
void QTextLength_delete(void* self) { delete (QTextLength*)self; }

}

#include <QtGui/qaccessible.h>
#include <QtCore/qstring.h>
#include <QtGui/qaccessible_base.h>
#include <QtCore/qrect.h>
#include <QtCore/qpoint.h>
#include "gen_qaccessibletextinterface.h"
#include "gen_purego_cb.h"

class QAccessibleTextInterface__Go : public QAccessibleTextInterface {
public:
	long long __goId;
	unsigned char __goOn[16] = {};
	template<typename... A> QAccessibleTextInterface__Go(long long id, A&&... a) : QAccessibleTextInterface(std::forward<A>(a)...), __goId(id) {}
	void selection(int selectionIndex, int* startOffset, int* endOffset) const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_void_int_intp_intp) { ((void(*)(long long, long long, int, int*, int*, int*))cppgo_virt_void_int_intp_intp)(__goId, 5156601901670015669LL, selectionIndex, (int*)startOffset, (int*)endOffset, &__h); if (__h) return; }
		return;
	}
	int selectionCount() const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -4199621612085967932LL, &__h); if (__h) return __r; }
		return int{};
	}
	void addSelection(int startOffset, int endOffset) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_void_int_int) { ((void(*)(long long, long long, int, int, int*))cppgo_virt_void_int_int)(__goId, 1701418371756208442LL, startOffset, endOffset, &__h); if (__h) return; }
		return;
	}
	void removeSelection(int selectionIndex) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 4093222286646031805LL, selectionIndex, &__h); if (__h) return; }
		return;
	}
	void setSelection(int selectionIndex, int startOffset, int endOffset) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_int_int_int) { ((void(*)(long long, long long, int, int, int, int*))cppgo_virt_void_int_int_int)(__goId, -1539266838375037305LL, selectionIndex, startOffset, endOffset, &__h); if (__h) return; }
		return;
	}
	int cursorPosition() const override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -4273823765597644148LL, &__h); if (__h) return __r; }
		return int{};
	}
	void setCursorPosition(int position) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, -6512451105733458522LL, position, &__h); if (__h) return; }
		return;
	}
	QString text(int startOffset, int endOffset) const override {
		int __h = 0;
		if (__goOn[7] && cppgo_virtp_charp_int_int) {
			long long __len = 0;
			char* __s = ((char*(*)(long long, long long, int, int, long long*, int*))cppgo_virtp_charp_int_int)(__goId, -430950356557152130LL, startOffset, endOffset, &__len, &__h);
			if (__h) { QString __r = QString::fromUtf8(__s, __len); free(__s); return __r; }
		}
		return QString{};
	}
	int characterCount() const override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -6653538151610911041LL, &__h); if (__h) return __r; }
		return int{};
	}
	QRect characterRect(int offset) const override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 3334036117981237742LL, offset, &__h); if (__h) return *(QRect*)__r; }
		return QRect{};
	}
	int offsetAtPoint(const QPoint& point) const override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, -4703424646947788745LL, (void*)&point, &__h); if (__h) return __r; }
		return int{};
	}
	void scrollToSubstring(int startIndex, int endIndex) override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_void_int_int) { ((void(*)(long long, long long, int, int, int*))cppgo_virt_void_int_int)(__goId, 6934233772508985160LL, startIndex, endIndex, &__h); if (__h) return; }
		return;
	}
	QString attributes(int offset, int* startOffset, int* endOffset) const override {
		int __h = 0;
		if (__goOn[12] && cppgo_virtp_charp_int_intp_intp) {
			long long __len = 0;
			char* __s = ((char*(*)(long long, long long, int, int*, int*, long long*, int*))cppgo_virtp_charp_int_intp_intp)(__goId, 6181953782893286932LL, offset, (int*)startOffset, (int*)endOffset, &__len, &__h);
			if (__h) { QString __r = QString::fromUtf8(__s, __len); free(__s); return __r; }
		}
		return QString{};
	}
	QString textBeforeOffset(int offset, QAccessible::TextBoundaryType boundaryType, int* startOffset, int* endOffset) const override {
		int __h = 0;
		if (__goOn[13] && cppgo_virtp_charp_int_int_intp_intp) {
			long long __len = 0;
			char* __s = ((char*(*)(long long, long long, int, int, int*, int*, long long*, int*))cppgo_virtp_charp_int_int_intp_intp)(__goId, 2244918558822102534LL, offset, (int)(boundaryType), (int*)startOffset, (int*)endOffset, &__len, &__h);
			if (__h) { QString __r = QString::fromUtf8(__s, __len); free(__s); return __r; }
		}
		return QAccessibleTextInterface::textBeforeOffset(offset, boundaryType, startOffset, endOffset);
	}
	QString textAfterOffset(int offset, QAccessible::TextBoundaryType boundaryType, int* startOffset, int* endOffset) const override {
		int __h = 0;
		if (__goOn[14] && cppgo_virtp_charp_int_int_intp_intp) {
			long long __len = 0;
			char* __s = ((char*(*)(long long, long long, int, int, int*, int*, long long*, int*))cppgo_virtp_charp_int_int_intp_intp)(__goId, 3240586353052927855LL, offset, (int)(boundaryType), (int*)startOffset, (int*)endOffset, &__len, &__h);
			if (__h) { QString __r = QString::fromUtf8(__s, __len); free(__s); return __r; }
		}
		return QAccessibleTextInterface::textAfterOffset(offset, boundaryType, startOffset, endOffset);
	}
	QString textAtOffset(int offset, QAccessible::TextBoundaryType boundaryType, int* startOffset, int* endOffset) const override {
		int __h = 0;
		if (__goOn[15] && cppgo_virtp_charp_int_int_intp_intp) {
			long long __len = 0;
			char* __s = ((char*(*)(long long, long long, int, int, int*, int*, long long*, int*))cppgo_virtp_charp_int_int_intp_intp)(__goId, -6929277707497019272LL, offset, (int)(boundaryType), (int*)startOffset, (int*)endOffset, &__len, &__h);
			if (__h) { QString __r = QString::fromUtf8(__s, __len); free(__s); return __r; }
		}
		return QAccessibleTextInterface::textAtOffset(offset, boundaryType, startOffset, endOffset);
	}
};

extern "C" {

void QAccessibleTextInterface__set_override(void* self, int idx) { ((QAccessibleTextInterface__Go*)self)->__goOn[idx] = 1; }
void QAccessibleTextInterface_selection(void* self, int selectionIndex, int* startOffset, int* endOffset) {
	int startOffset_local{};
	int endOffset_local{};
	((QAccessibleTextInterface*)self)->selection(selectionIndex, &startOffset_local, &endOffset_local);
	*startOffset = startOffset_local;
	*endOffset = endOffset_local;
}
int QAccessibleTextInterface_selectionCount(void* self) {
	auto&& __r = ((QAccessibleTextInterface*)self)->selectionCount();
	return __r;
}
void QAccessibleTextInterface_addSelection(void* self, int startOffset, int endOffset) {
	((QAccessibleTextInterface*)self)->addSelection(startOffset, endOffset);
}
void QAccessibleTextInterface_removeSelection(void* self, int selectionIndex) {
	((QAccessibleTextInterface*)self)->removeSelection(selectionIndex);
}
void QAccessibleTextInterface_setSelection(void* self, int selectionIndex, int startOffset, int endOffset) {
	((QAccessibleTextInterface*)self)->setSelection(selectionIndex, startOffset, endOffset);
}
int QAccessibleTextInterface_cursorPosition(void* self) {
	auto&& __r = ((QAccessibleTextInterface*)self)->cursorPosition();
	return __r;
}
void QAccessibleTextInterface_setCursorPosition(void* self, int position) {
	((QAccessibleTextInterface*)self)->setCursorPosition(position);
}
char* QAccessibleTextInterface_text(void* self, int startOffset, int endOffset, long long* __out_len) {
	auto&& __r = ((QAccessibleTextInterface*)self)->text(startOffset, endOffset);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QAccessibleTextInterface_textBeforeOffset(void* self, int offset, int boundaryType, int* startOffset, int* endOffset, long long* __out_len) {
	int startOffset_local{};
	int endOffset_local{};
	auto&& __r = ((QAccessibleTextInterface*)self)->textBeforeOffset(offset, (QAccessible::TextBoundaryType)(boundaryType), &startOffset_local, &endOffset_local);
	*startOffset = startOffset_local;
	*endOffset = endOffset_local;
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QAccessibleTextInterface_textAfterOffset(void* self, int offset, int boundaryType, int* startOffset, int* endOffset, long long* __out_len) {
	int startOffset_local{};
	int endOffset_local{};
	auto&& __r = ((QAccessibleTextInterface*)self)->textAfterOffset(offset, (QAccessible::TextBoundaryType)(boundaryType), &startOffset_local, &endOffset_local);
	*startOffset = startOffset_local;
	*endOffset = endOffset_local;
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QAccessibleTextInterface_textAtOffset(void* self, int offset, int boundaryType, int* startOffset, int* endOffset, long long* __out_len) {
	int startOffset_local{};
	int endOffset_local{};
	auto&& __r = ((QAccessibleTextInterface*)self)->textAtOffset(offset, (QAccessible::TextBoundaryType)(boundaryType), &startOffset_local, &endOffset_local);
	*startOffset = startOffset_local;
	*endOffset = endOffset_local;
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QAccessibleTextInterface_characterCount(void* self) {
	auto&& __r = ((QAccessibleTextInterface*)self)->characterCount();
	return __r;
}
void* QAccessibleTextInterface_characterRect(void* self, int offset) {
	auto&& __r = ((QAccessibleTextInterface*)self)->characterRect(offset);
	return new QRect(__r);
}
int QAccessibleTextInterface_offsetAtPoint(void* self, void* point) {
	auto&& __r = ((QAccessibleTextInterface*)self)->offsetAtPoint(*(QPoint*)point);
	return __r;
}
void QAccessibleTextInterface_scrollToSubstring(void* self, int startIndex, int endIndex) {
	((QAccessibleTextInterface*)self)->scrollToSubstring(startIndex, endIndex);
}
char* QAccessibleTextInterface_attributes(void* self, int offset, int* startOffset, int* endOffset, long long* __out_len) {
	int startOffset_local{};
	int endOffset_local{};
	auto&& __r = ((QAccessibleTextInterface*)self)->attributes(offset, &startOffset_local, &endOffset_local);
	*startOffset = startOffset_local;
	*endOffset = endOffset_local;
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QAccessibleTextInterface_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QAccessibleTextInterface*)self)->operator=(*(QAccessibleTextInterface*)arg1);
	return (void*)&(__r);
}
char* QAccessibleTextInterface_textBeforeOffset_default(void* self, int offset, int boundaryType, int* startOffset, int* endOffset, long long* __out_len) {
	int startOffset_local{};
	int endOffset_local{};
	auto&& __r = ((QAccessibleTextInterface*)self)->QAccessibleTextInterface::textBeforeOffset(offset, (QAccessible::TextBoundaryType)(boundaryType), &startOffset_local, &endOffset_local);
	*startOffset = startOffset_local;
	*endOffset = endOffset_local;
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QAccessibleTextInterface_textAfterOffset_default(void* self, int offset, int boundaryType, int* startOffset, int* endOffset, long long* __out_len) {
	int startOffset_local{};
	int endOffset_local{};
	auto&& __r = ((QAccessibleTextInterface*)self)->QAccessibleTextInterface::textAfterOffset(offset, (QAccessible::TextBoundaryType)(boundaryType), &startOffset_local, &endOffset_local);
	*startOffset = startOffset_local;
	*endOffset = endOffset_local;
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QAccessibleTextInterface_textAtOffset_default(void* self, int offset, int boundaryType, int* startOffset, int* endOffset, long long* __out_len) {
	int startOffset_local{};
	int endOffset_local{};
	auto&& __r = ((QAccessibleTextInterface*)self)->QAccessibleTextInterface::textAtOffset(offset, (QAccessible::TextBoundaryType)(boundaryType), &startOffset_local, &endOffset_local);
	*startOffset = startOffset_local;
	*endOffset = endOffset_local;
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QAccessibleTextInterface__Go_new(long long id) { return new QAccessibleTextInterface__Go(id); }
void QAccessibleTextInterface_delete(void* self) { delete (QAccessibleTextInterface*)self; }

}

#include <QtGui/qaccessibleplugin.h>
#include <QtCore/qobject.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtGui/qaccessible.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qvariant.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qaccessibleplugin.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QAccessiblePlugin__Go : public QAccessiblePlugin {
public:
	long long __goId;
	unsigned char __goOn[9] = {};
	template<typename... A> QAccessiblePlugin__Go(long long id, A&&... a) : QAccessiblePlugin(std::forward<A>(a)...), __goId(id) {}
	using QAccessiblePlugin::sender;
	using QAccessiblePlugin::senderSignalIndex;
	using QAccessiblePlugin::receivers;
	using QAccessiblePlugin::isSignalConnected;
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QAccessiblePlugin::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QAccessiblePlugin::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QAccessiblePlugin::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QAccessiblePlugin::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QAccessiblePlugin::disconnectNotify(__a0); }
	QAccessibleInterface* create(const QString& key, QObject* object) override {
		int __h = 0;
		QByteArray key_ba = key.toUtf8();
		if (__goOn[0] && cppgo_virt_voidp_constcharp_longlong_voidp) { void* __r = ((void*(*)(long long, long long, const char*, long long, void*, int*))cppgo_virt_voidp_constcharp_longlong_voidp)(__goId, -626457835795418115LL, (char*)key_ba.constData(), key_ba.size(), (void*)object, &__h); if (__h) return (QAccessibleInterface*)__r; }
		return nullptr;
	}
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QAccessiblePlugin::metaObject();
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QAccessiblePlugin::event(event);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QAccessiblePlugin::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QAccessiblePlugin::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QAccessiblePlugin::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QAccessiblePlugin::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QAccessiblePlugin::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QAccessiblePlugin::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QAccessiblePlugin;
static std::mutex __connsMu_QAccessiblePlugin;

extern "C" {

void* QAccessiblePlugin_base(void* self) { return (void*)static_cast<QObject*>((QAccessiblePlugin*)self); }
void QAccessiblePlugin__set_override(void* self, int idx) { ((QAccessiblePlugin__Go*)self)->__goOn[idx] = 1; }
void* QAccessiblePlugin_fromQObject(void* o) { return (void*)qobject_cast<QAccessiblePlugin*>((QObject*)o); }
char* QAccessiblePlugin_tr(const char* s, long long* __out_len) {
	auto&& __r = QAccessiblePlugin::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QAccessiblePlugin_create(void* self, const char* key, long long key_len, void* object) {
	QString key_qstr = QString::fromUtf8(key, key_len);
	auto&& __r = ((QAccessiblePlugin*)self)->create(key_qstr, (QObject*)object);
	return (void*)(__r);
}
void* QAccessiblePlugin_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QAccessiblePlugin::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QAccessiblePlugin_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QAccessiblePlugin::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QAccessiblePlugin_sender(void* self) {
	auto&& __r = ((QAccessiblePlugin__Go*)self)->sender();
	return (void*)(__r);
}
int QAccessiblePlugin_senderSignalIndex(void* self) {
	auto&& __r = ((QAccessiblePlugin__Go*)self)->senderSignalIndex();
	return __r;
}
int QAccessiblePlugin_receivers(void* self, const char* signal) {
	auto&& __r = ((QAccessiblePlugin__Go*)self)->receivers(signal);
	return __r;
}
int QAccessiblePlugin_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QAccessiblePlugin__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QAccessiblePlugin_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QAccessiblePlugin::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QAccessiblePlugin_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QAccessiblePlugin::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QAccessiblePlugin_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QAccessiblePlugin::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QAccessiblePlugin_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QAccessiblePlugin::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QAccessiblePlugin_metaObject_default(void* self) {
	auto&& __r = ((QAccessiblePlugin*)self)->QAccessiblePlugin::metaObject();
	return (void*)(__r);
}
int QAccessiblePlugin_event_default(void* self, void* event) {
	auto&& __r = ((QAccessiblePlugin*)self)->QAccessiblePlugin::event((QEvent*)event);
	return __r;
}
int QAccessiblePlugin_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QAccessiblePlugin*)self)->QAccessiblePlugin::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QAccessiblePlugin_timerEvent_default(void* self, void* event) {
	((QAccessiblePlugin__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QAccessiblePlugin_childEvent_default(void* self, void* event) {
	((QAccessiblePlugin__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QAccessiblePlugin_customEvent_default(void* self, void* event) {
	((QAccessiblePlugin__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QAccessiblePlugin_connectNotify_default(void* self, void* signal) {
	((QAccessiblePlugin__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QAccessiblePlugin_disconnectNotify_default(void* self, void* signal) {
	((QAccessiblePlugin__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QAccessiblePlugin_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAccessiblePlugin*)self, &QAccessiblePlugin::destroyed, (QAccessiblePlugin*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAccessiblePlugin);
	__conns_QAccessiblePlugin[__id] = __c;
}
void QAccessiblePlugin_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAccessiblePlugin*)self, &QAccessiblePlugin::destroyed, (QAccessiblePlugin*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAccessiblePlugin);
	__conns_QAccessiblePlugin[__id] = __c;
}
void QAccessiblePlugin_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAccessiblePlugin*)self, &QAccessiblePlugin::objectNameChanged, (QAccessiblePlugin*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAccessiblePlugin);
	__conns_QAccessiblePlugin[__id] = __c;
}
void QAccessiblePlugin__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QAccessiblePlugin);
		auto __it = __conns_QAccessiblePlugin.find(id);
		if (__it == __conns_QAccessiblePlugin.end()) return;
		__c = __it->second;
		__conns_QAccessiblePlugin.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QAccessiblePlugin__Go_new(long long id, void* parent) { return new QAccessiblePlugin__Go(id, (QObject*)parent); }
void QAccessiblePlugin_delete(void* self) { delete (QAccessiblePlugin*)self; }

}

#include <QtGui/qiconengineplugin.h>
#include <QtCore/qobject.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtGui/qiconengine.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qvariant.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qiconengineplugin.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QIconEnginePlugin__Go : public QIconEnginePlugin {
public:
	long long __goId;
	unsigned char __goOn[9] = {};
	template<typename... A> QIconEnginePlugin__Go(long long id, A&&... a) : QIconEnginePlugin(std::forward<A>(a)...), __goId(id) {}
	using QIconEnginePlugin::sender;
	using QIconEnginePlugin::senderSignalIndex;
	using QIconEnginePlugin::receivers;
	using QIconEnginePlugin::isSignalConnected;
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QIconEnginePlugin::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QIconEnginePlugin::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QIconEnginePlugin::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QIconEnginePlugin::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QIconEnginePlugin::disconnectNotify(__a0); }
	QIconEngine* create(const QString& filename) override {
		int __h = 0;
		QByteArray filename_ba = filename.toUtf8();
		if (__goOn[0] && cppgo_virt_voidp_constcharp_longlong) { void* __r = ((void*(*)(long long, long long, const char*, long long, int*))cppgo_virt_voidp_constcharp_longlong)(__goId, -626457835795418115LL, (char*)filename_ba.constData(), filename_ba.size(), &__h); if (__h) return (QIconEngine*)__r; }
		return nullptr;
	}
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QIconEnginePlugin::metaObject();
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QIconEnginePlugin::event(event);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QIconEnginePlugin::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QIconEnginePlugin::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QIconEnginePlugin::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QIconEnginePlugin::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QIconEnginePlugin::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QIconEnginePlugin::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QIconEnginePlugin;
static std::mutex __connsMu_QIconEnginePlugin;

extern "C" {

void* QIconEnginePlugin_base(void* self) { return (void*)static_cast<QObject*>((QIconEnginePlugin*)self); }
void QIconEnginePlugin__set_override(void* self, int idx) { ((QIconEnginePlugin__Go*)self)->__goOn[idx] = 1; }
void* QIconEnginePlugin_fromQObject(void* o) { return (void*)qobject_cast<QIconEnginePlugin*>((QObject*)o); }
char* QIconEnginePlugin_tr(const char* s, long long* __out_len) {
	auto&& __r = QIconEnginePlugin::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QIconEnginePlugin_create(void* self, const char* filename, long long filename_len) {
	QString filename_qstr = QString::fromUtf8(filename, filename_len);
	auto&& __r = ((QIconEnginePlugin*)self)->create(filename_qstr);
	return (void*)(__r);
}
void* QIconEnginePlugin_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QIconEnginePlugin::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QIconEnginePlugin_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QIconEnginePlugin::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QIconEnginePlugin_sender(void* self) {
	auto&& __r = ((QIconEnginePlugin__Go*)self)->sender();
	return (void*)(__r);
}
int QIconEnginePlugin_senderSignalIndex(void* self) {
	auto&& __r = ((QIconEnginePlugin__Go*)self)->senderSignalIndex();
	return __r;
}
int QIconEnginePlugin_receivers(void* self, const char* signal) {
	auto&& __r = ((QIconEnginePlugin__Go*)self)->receivers(signal);
	return __r;
}
int QIconEnginePlugin_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QIconEnginePlugin__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QIconEnginePlugin_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QIconEnginePlugin::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QIconEnginePlugin_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QIconEnginePlugin::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QIconEnginePlugin_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QIconEnginePlugin::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QIconEnginePlugin_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QIconEnginePlugin::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QIconEnginePlugin_metaObject_default(void* self) {
	auto&& __r = ((QIconEnginePlugin*)self)->QIconEnginePlugin::metaObject();
	return (void*)(__r);
}
int QIconEnginePlugin_event_default(void* self, void* event) {
	auto&& __r = ((QIconEnginePlugin*)self)->QIconEnginePlugin::event((QEvent*)event);
	return __r;
}
int QIconEnginePlugin_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QIconEnginePlugin*)self)->QIconEnginePlugin::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QIconEnginePlugin_timerEvent_default(void* self, void* event) {
	((QIconEnginePlugin__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QIconEnginePlugin_childEvent_default(void* self, void* event) {
	((QIconEnginePlugin__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QIconEnginePlugin_customEvent_default(void* self, void* event) {
	((QIconEnginePlugin__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QIconEnginePlugin_connectNotify_default(void* self, void* signal) {
	((QIconEnginePlugin__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QIconEnginePlugin_disconnectNotify_default(void* self, void* signal) {
	((QIconEnginePlugin__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QIconEnginePlugin_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QIconEnginePlugin*)self, &QIconEnginePlugin::destroyed, (QIconEnginePlugin*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QIconEnginePlugin);
	__conns_QIconEnginePlugin[__id] = __c;
}
void QIconEnginePlugin_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QIconEnginePlugin*)self, &QIconEnginePlugin::destroyed, (QIconEnginePlugin*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QIconEnginePlugin);
	__conns_QIconEnginePlugin[__id] = __c;
}
void QIconEnginePlugin_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QIconEnginePlugin*)self, &QIconEnginePlugin::objectNameChanged, (QIconEnginePlugin*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QIconEnginePlugin);
	__conns_QIconEnginePlugin[__id] = __c;
}
void QIconEnginePlugin__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QIconEnginePlugin);
		auto __it = __conns_QIconEnginePlugin.find(id);
		if (__it == __conns_QIconEnginePlugin.end()) return;
		__c = __it->second;
		__conns_QIconEnginePlugin.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QIconEnginePlugin__Go_new(long long id, void* parent) { return new QIconEnginePlugin__Go(id, (QObject*)parent); }
void QIconEnginePlugin_delete(void* self) { delete (QIconEnginePlugin*)self; }

}

#include <QtGui/qpainterpath.h>
#include <QtCore/qpoint.h>
#include "gen_qpainterpath__element.h"

extern "C" {

void* QPainterPath__Element_new() {
	return new QPainterPath::Element();
}
void* QPainterPath__Element_new2(void* arg1) {
	return new QPainterPath::Element(*(QPainterPath::Element*)arg1);
}
int QPainterPath__Element_isMoveTo(void* self) {
	auto&& __r = ((QPainterPath::Element*)self)->isMoveTo();
	return __r;
}
int QPainterPath__Element_isLineTo(void* self) {
	auto&& __r = ((QPainterPath::Element*)self)->isLineTo();
	return __r;
}
int QPainterPath__Element_isCurveTo(void* self) {
	auto&& __r = ((QPainterPath::Element*)self)->isCurveTo();
	return __r;
}
void* QPainterPath__Element_ToQPointF(void* self) {
	auto&& __r = static_cast<QPointF>(*(QPainterPath::Element*)self);
	return new QPointF(__r);
}
int QPainterPath__Element_OperatorEqual(void* self, void* e) {
	auto&& __r = ((QPainterPath::Element*)self)->operator==(*(QPainterPath::Element*)e);
	return __r;
}
int QPainterPath__Element_OperatorNotEqual(void* self, void* e) {
	auto&& __r = ((QPainterPath::Element*)self)->operator!=(*(QPainterPath::Element*)e);
	return __r;
}
double QPainterPath__Element_x_get(void* self) {
	auto __r = ((QPainterPath::Element*)self)->x;
	return __r;
}
void QPainterPath__Element_x_set(void* self, double x) {
	((QPainterPath::Element*)self)->x = x;
}
double QPainterPath__Element_y_get(void* self) {
	auto __r = ((QPainterPath::Element*)self)->y;
	return __r;
}
void QPainterPath__Element_y_set(void* self, double y) {
	((QPainterPath::Element*)self)->y = y;
}
int QPainterPath__Element_type_get(void* self) {
	auto __r = ((QPainterPath::Element*)self)->type;
	return (int)(__r);
}
void QPainterPath__Element_type_set(void* self, int type_) {
	((QPainterPath::Element*)self)->type = (QPainterPath::ElementType)(type_);
}
void QPainterPath__Element_delete(void* self) { delete (QPainterPath::Element*)self; }

}

#include <QtGui/qtextobject.h>
#include <QtGui/qtextformat.h>
#include <QtCore/qstring.h>
#include <QtGui/qglyphrun.h>
#include "gen_qtextfragment.h"

extern "C" {

void* QTextFragment_new() {
	return new QTextFragment();
}
void* QTextFragment_new2(void* o) {
	return new QTextFragment(*(QTextFragment*)o);
}
void* QTextFragment_OperatorAssign(void* self, void* o) {
	auto&& __r = ((QTextFragment*)self)->operator=(*(QTextFragment*)o);
	return (void*)&(__r);
}
int QTextFragment_isValid(void* self) {
	auto&& __r = ((QTextFragment*)self)->isValid();
	return __r;
}
int QTextFragment_OperatorEqual(void* self, void* o) {
	auto&& __r = ((QTextFragment*)self)->operator==(*(QTextFragment*)o);
	return __r;
}
int QTextFragment_OperatorNotEqual(void* self, void* o) {
	auto&& __r = ((QTextFragment*)self)->operator!=(*(QTextFragment*)o);
	return __r;
}
int QTextFragment_OperatorLesser(void* self, void* o) {
	auto&& __r = ((QTextFragment*)self)->operator<(*(QTextFragment*)o);
	return __r;
}
int QTextFragment_position(void* self) {
	auto&& __r = ((QTextFragment*)self)->position();
	return __r;
}
int QTextFragment_length(void* self) {
	auto&& __r = ((QTextFragment*)self)->length();
	return __r;
}
int QTextFragment_contains(void* self, int position) {
	auto&& __r = ((QTextFragment*)self)->contains(position);
	return __r;
}
void* QTextFragment_charFormat(void* self) {
	auto&& __r = ((QTextFragment*)self)->charFormat();
	return new QTextCharFormat(__r);
}
int QTextFragment_charFormatIndex(void* self) {
	auto&& __r = ((QTextFragment*)self)->charFormatIndex();
	return __r;
}
char* QTextFragment_text(void* self, long long* __out_len) {
	auto&& __r = ((QTextFragment*)self)->text();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void** QTextFragment_glyphRuns(void* self, long long* __out_len) {
	auto&& __r = ((QTextFragment*)self)->glyphRuns();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QGlyphRun(__e);
	return __arr;
}
void** QTextFragment_glyphRuns2(void* self, int from, long long* __out_len) {
	auto&& __r = ((QTextFragment*)self)->glyphRuns(from);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QGlyphRun(__e);
	return __arr;
}
void** QTextFragment_glyphRuns3(void* self, int from, int length, long long* __out_len) {
	auto&& __r = ((QTextFragment*)self)->glyphRuns(from, length);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QGlyphRun(__e);
	return __arr;
}
void QTextFragment_delete(void* self) { delete (QTextFragment*)self; }

}

#include <QtWidgets/qtabwidget.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtGui/qicon.h>
#include <QtCore/qsize.h>
#include <QtCore/qnamespace.h>
#include <QtWidgets/qtabbar.h>
#include <QtGui/qevent.h>
#include <QtCore/qcoreevent.h>
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
#include "gen_qtabwidget.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QTabWidget__Go : public QTabWidget {
public:
	long long __goId;
	unsigned char __goOn[51] = {};
	template<typename... A> QTabWidget__Go(long long id, A&&... a) : QTabWidget(std::forward<A>(a)...), __goId(id) {}
	using QTabWidget::setTabBar;
	using QTabWidget::updateMicroFocus;
	using QTabWidget::create;
	using QTabWidget::destroy;
	using QTabWidget::focusNextChild;
	using QTabWidget::focusPreviousChild;
	using QTabWidget::sender;
	using QTabWidget::senderSignalIndex;
	using QTabWidget::receivers;
	using QTabWidget::isSignalConnected;
	using QTabWidget::getDecodedMetricF;
	void cppgoBase_tabInserted(int __a0) { QTabWidget::tabInserted(__a0); }
	void cppgoBase_tabRemoved(int __a0) { QTabWidget::tabRemoved(__a0); }
	void cppgoBase_showEvent(QShowEvent* __a0) { QTabWidget::showEvent(__a0); }
	void cppgoBase_resizeEvent(QResizeEvent* __a0) { QTabWidget::resizeEvent(__a0); }
	void cppgoBase_keyPressEvent(QKeyEvent* __a0) { QTabWidget::keyPressEvent(__a0); }
	void cppgoBase_paintEvent(QPaintEvent* __a0) { QTabWidget::paintEvent(__a0); }
	void cppgoBase_changeEvent(QEvent* __a0) { QTabWidget::changeEvent(__a0); }
	bool cppgoBase_event(QEvent* __a0) { return QTabWidget::event(__a0); }
	void cppgoBase_initStyleOption(QStyleOptionTabWidgetFrame* __a0) { QTabWidget::initStyleOption(__a0); }
	void cppgoBase_mousePressEvent(QMouseEvent* __a0) { QTabWidget::mousePressEvent(__a0); }
	void cppgoBase_mouseReleaseEvent(QMouseEvent* __a0) { QTabWidget::mouseReleaseEvent(__a0); }
	void cppgoBase_mouseDoubleClickEvent(QMouseEvent* __a0) { QTabWidget::mouseDoubleClickEvent(__a0); }
	void cppgoBase_mouseMoveEvent(QMouseEvent* __a0) { QTabWidget::mouseMoveEvent(__a0); }
	void cppgoBase_wheelEvent(QWheelEvent* __a0) { QTabWidget::wheelEvent(__a0); }
	void cppgoBase_keyReleaseEvent(QKeyEvent* __a0) { QTabWidget::keyReleaseEvent(__a0); }
	void cppgoBase_focusInEvent(QFocusEvent* __a0) { QTabWidget::focusInEvent(__a0); }
	void cppgoBase_focusOutEvent(QFocusEvent* __a0) { QTabWidget::focusOutEvent(__a0); }
	void cppgoBase_enterEvent(QEnterEvent* __a0) { QTabWidget::enterEvent(__a0); }
	void cppgoBase_leaveEvent(QEvent* __a0) { QTabWidget::leaveEvent(__a0); }
	void cppgoBase_moveEvent(QMoveEvent* __a0) { QTabWidget::moveEvent(__a0); }
	void cppgoBase_closeEvent(QCloseEvent* __a0) { QTabWidget::closeEvent(__a0); }
	void cppgoBase_contextMenuEvent(QContextMenuEvent* __a0) { QTabWidget::contextMenuEvent(__a0); }
	void cppgoBase_tabletEvent(QTabletEvent* __a0) { QTabWidget::tabletEvent(__a0); }
	void cppgoBase_actionEvent(QActionEvent* __a0) { QTabWidget::actionEvent(__a0); }
	void cppgoBase_dragEnterEvent(QDragEnterEvent* __a0) { QTabWidget::dragEnterEvent(__a0); }
	void cppgoBase_dragMoveEvent(QDragMoveEvent* __a0) { QTabWidget::dragMoveEvent(__a0); }
	void cppgoBase_dragLeaveEvent(QDragLeaveEvent* __a0) { QTabWidget::dragLeaveEvent(__a0); }
	void cppgoBase_dropEvent(QDropEvent* __a0) { QTabWidget::dropEvent(__a0); }
	void cppgoBase_hideEvent(QHideEvent* __a0) { QTabWidget::hideEvent(__a0); }
	int cppgoBase_metric(int __a0) { return QTabWidget::metric((QPaintDevice::PaintDeviceMetric)__a0); }
	void cppgoBase_initPainter(QPainter* __a0) { QTabWidget::initPainter(__a0); }
	QPaintDevice* cppgoBase_redirected(QPoint* __a0) { return QTabWidget::redirected(__a0); }
	QPainter* cppgoBase_sharedPainter() { return QTabWidget::sharedPainter(); }
	void cppgoBase_inputMethodEvent(QInputMethodEvent* __a0) { QTabWidget::inputMethodEvent(__a0); }
	bool cppgoBase_focusNextPrevChild(bool __a0) { return QTabWidget::focusNextPrevChild(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QTabWidget::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QTabWidget::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QTabWidget::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QTabWidget::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QTabWidget::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QTabWidget::metaObject();
	}
	QSize sizeHint() const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 729690261457684101LL, &__h); if (__h) return *(QSize*)__r; }
		return QTabWidget::sizeHint();
	}
	QSize minimumSizeHint() const override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 7049264599837532491LL, &__h); if (__h) return *(QSize*)__r; }
		return QTabWidget::minimumSizeHint();
	}
	int heightForWidth(int width) const override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 4311534060495737339LL, width, &__h); if (__h) return __r; }
		return QTabWidget::heightForWidth(width);
	}
	bool hasHeightForWidth() const override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8266762397995423915LL, &__h); if (__h) return __r; }
		return QTabWidget::hasHeightForWidth();
	}
	void tabInserted(int index) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, -2734286415018532480LL, index, &__h); if (__h) return; }
		QTabWidget::tabInserted(index);
	}
	void tabRemoved(int index) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, -3307029722643392856LL, index, &__h); if (__h) return; }
		QTabWidget::tabRemoved(index);
	}
	void showEvent(QShowEvent* arg1) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6825369174751916LL, (void*)arg1, &__h); if (__h) return; }
		QTabWidget::showEvent(arg1);
	}
	void resizeEvent(QResizeEvent* arg1) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4302917285655014661LL, (void*)arg1, &__h); if (__h) return; }
		QTabWidget::resizeEvent(arg1);
	}
	void keyPressEvent(QKeyEvent* arg1) override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1023791977679793185LL, (void*)arg1, &__h); if (__h) return; }
		QTabWidget::keyPressEvent(arg1);
	}
	void paintEvent(QPaintEvent* arg1) override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5290442163167458465LL, (void*)arg1, &__h); if (__h) return; }
		QTabWidget::paintEvent(arg1);
	}
	void changeEvent(QEvent* arg1) override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4959376063826165943LL, (void*)arg1, &__h); if (__h) return; }
		QTabWidget::changeEvent(arg1);
	}
	bool event(QEvent* arg1) override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)arg1, &__h); if (__h) return __r; }
		return QTabWidget::event(arg1);
	}
	void initStyleOption(QStyleOptionTabWidgetFrame* option) const override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -488693769771308531LL, (void*)option, &__h); if (__h) return; }
		QTabWidget::initStyleOption(option);
	}
	int devType() const override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8659679893868524044LL, &__h); if (__h) return __r; }
		return QTabWidget::devType();
	}
	void setVisible(bool visible) override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 2110182719375245575LL, visible, &__h); if (__h) return; }
		QTabWidget::setVisible(visible);
	}
	QPaintEngine* paintEngine() const override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 6398007404756346381LL, &__h); if (__h) return (QPaintEngine*)__r; }
		return QTabWidget::paintEngine();
	}
	void mousePressEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -332090898078543319LL, (void*)event, &__h); if (__h) return; }
		QTabWidget::mousePressEvent(event);
	}
	void mouseReleaseEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4415920049297795213LL, (void*)event, &__h); if (__h) return; }
		QTabWidget::mouseReleaseEvent(event);
	}
	void mouseDoubleClickEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4029073905697023887LL, (void*)event, &__h); if (__h) return; }
		QTabWidget::mouseDoubleClickEvent(event);
	}
	void mouseMoveEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4255494004099806581LL, (void*)event, &__h); if (__h) return; }
		QTabWidget::mouseMoveEvent(event);
	}
	void wheelEvent(QWheelEvent* event) override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1919988340934511316LL, (void*)event, &__h); if (__h) return; }
		QTabWidget::wheelEvent(event);
	}
	void keyReleaseEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 772598475950050299LL, (void*)event, &__h); if (__h) return; }
		QTabWidget::keyReleaseEvent(event);
	}
	void focusInEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 975043767033552552LL, (void*)event, &__h); if (__h) return; }
		QTabWidget::focusInEvent(event);
	}
	void focusOutEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -385596010576048993LL, (void*)event, &__h); if (__h) return; }
		QTabWidget::focusOutEvent(event);
	}
	void enterEvent(QEnterEvent* event) override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6316270346668237895LL, (void*)event, &__h); if (__h) return; }
		QTabWidget::enterEvent(event);
	}
	void leaveEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -551070009623711640LL, (void*)event, &__h); if (__h) return; }
		QTabWidget::leaveEvent(event);
	}
	void moveEvent(QMoveEvent* event) override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -9161949150489525396LL, (void*)event, &__h); if (__h) return; }
		QTabWidget::moveEvent(event);
	}
	void closeEvent(QCloseEvent* event) override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1989937050019464783LL, (void*)event, &__h); if (__h) return; }
		QTabWidget::closeEvent(event);
	}
	void contextMenuEvent(QContextMenuEvent* event) override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6698043222227006015LL, (void*)event, &__h); if (__h) return; }
		QTabWidget::contextMenuEvent(event);
	}
	void tabletEvent(QTabletEvent* event) override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -212068866860701141LL, (void*)event, &__h); if (__h) return; }
		QTabWidget::tabletEvent(event);
	}
	void actionEvent(QActionEvent* event) override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1697444140088555581LL, (void*)event, &__h); if (__h) return; }
		QTabWidget::actionEvent(event);
	}
	void dragEnterEvent(QDragEnterEvent* event) override {
		int __h = 0;
		if (__goOn[32] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 206122715919917589LL, (void*)event, &__h); if (__h) return; }
		QTabWidget::dragEnterEvent(event);
	}
	void dragMoveEvent(QDragMoveEvent* event) override {
		int __h = 0;
		if (__goOn[33] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1781545137352745978LL, (void*)event, &__h); if (__h) return; }
		QTabWidget::dragMoveEvent(event);
	}
	void dragLeaveEvent(QDragLeaveEvent* event) override {
		int __h = 0;
		if (__goOn[34] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -3139904158373725006LL, (void*)event, &__h); if (__h) return; }
		QTabWidget::dragLeaveEvent(event);
	}
	void dropEvent(QDropEvent* event) override {
		int __h = 0;
		if (__goOn[35] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 9107058682754591716LL, (void*)event, &__h); if (__h) return; }
		QTabWidget::dropEvent(event);
	}
	void hideEvent(QHideEvent* event) override {
		int __h = 0;
		if (__goOn[36] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -199718707876719129LL, (void*)event, &__h); if (__h) return; }
		QTabWidget::hideEvent(event);
	}
	bool nativeEvent(const QByteArray& eventType, void* message, qintptr* result) override {
		int __h = 0;
		if (__goOn[37] && cppgo_virt_int_voidp_voidp_longlongp) { int __r = ((int(*)(long long, long long, void*, void*, long long*, int*))cppgo_virt_int_voidp_voidp_longlongp)(__goId, -2039166586246443706LL, (void*)&eventType, (void*)message, (long long*)result, &__h); if (__h) return __r; }
		return QTabWidget::nativeEvent(eventType, message, result);
	}
	int metric(QPaintDevice::PaintDeviceMetric arg1) const override {
		int __h = 0;
		if (__goOn[38] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 427629649267311777LL, (int)(arg1), &__h); if (__h) return __r; }
		return QTabWidget::metric(arg1);
	}
	void initPainter(QPainter* painter) const override {
		int __h = 0;
		if (__goOn[39] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6477500296980050210LL, (void*)painter, &__h); if (__h) return; }
		QTabWidget::initPainter(painter);
	}
	QPaintDevice* redirected(QPoint* offset) const override {
		int __h = 0;
		if (__goOn[40] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 4127240633310655048LL, (void*)offset, &__h); if (__h) return (QPaintDevice*)__r; }
		return QTabWidget::redirected(offset);
	}
	QPainter* sharedPainter() const override {
		int __h = 0;
		if (__goOn[41] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 5792771328741774991LL, &__h); if (__h) return (QPainter*)__r; }
		return QTabWidget::sharedPainter();
	}
	void inputMethodEvent(QInputMethodEvent* arg1) override {
		int __h = 0;
		if (__goOn[42] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5438819661942843222LL, (void*)arg1, &__h); if (__h) return; }
		QTabWidget::inputMethodEvent(arg1);
	}
	QVariant inputMethodQuery(Qt::InputMethodQuery arg1) const override {
		int __h = 0;
		if (__goOn[43] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7345853097340709638LL, (int)(arg1), &__h); if (__h) return *(QVariant*)__r; }
		return QTabWidget::inputMethodQuery(arg1);
	}
	bool focusNextPrevChild(bool next) override {
		int __h = 0;
		if (__goOn[44] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 5898328517505784191LL, next, &__h); if (__h) return __r; }
		return QTabWidget::focusNextPrevChild(next);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[45] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QTabWidget::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[46] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QTabWidget::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[47] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QTabWidget::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[48] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QTabWidget::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[49] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QTabWidget::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[50] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QTabWidget::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QTabWidget;
static std::mutex __connsMu_QTabWidget;

extern "C" {

void* QTabWidget_base(void* self) { return (void*)static_cast<QWidget*>((QTabWidget*)self); }
void QTabWidget_chain(void* self, void** out) {
	QTabWidget* __p = (QTabWidget*)self;
	QWidget* __b0 = static_cast<QWidget*>(__p); out[0] = (void*)__b0;
	QObject* __b1 = static_cast<QObject*>(__b0); out[1] = (void*)__b1;
}
void QTabWidget__set_override(void* self, int idx) { ((QTabWidget__Go*)self)->__goOn[idx] = 1; }
void* QTabWidget_fromQObject(void* o) { return (void*)qobject_cast<QTabWidget*>((QObject*)o); }
void* QTabWidget_new(long long __goid, void* parent) {
	return new QTabWidget__Go(__goid, (QWidget*)parent);
}
void* QTabWidget_new2(long long __goid) {
	return new QTabWidget__Go(__goid);
}
char* QTabWidget_tr(const char* s, long long* __out_len) {
	auto&& __r = QTabWidget::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QTabWidget_addTab(void* self, void* widget, const char* arg2, long long arg2_len) {
	QString arg2_qstr = QString::fromUtf8(arg2, arg2_len);
	auto&& __r = ((QTabWidget*)self)->addTab((QWidget*)widget, arg2_qstr);
	return __r;
}
int QTabWidget_addTab2(void* self, void* widget, void* icon, const char* label, long long label_len) {
	QString label_qstr = QString::fromUtf8(label, label_len);
	auto&& __r = ((QTabWidget*)self)->addTab((QWidget*)widget, *(QIcon*)icon, label_qstr);
	return __r;
}
int QTabWidget_insertTab(void* self, int index, void* widget, const char* arg3, long long arg3_len) {
	QString arg3_qstr = QString::fromUtf8(arg3, arg3_len);
	auto&& __r = ((QTabWidget*)self)->insertTab(index, (QWidget*)widget, arg3_qstr);
	return __r;
}
int QTabWidget_insertTab2(void* self, int index, void* widget, void* icon, const char* label, long long label_len) {
	QString label_qstr = QString::fromUtf8(label, label_len);
	auto&& __r = ((QTabWidget*)self)->insertTab(index, (QWidget*)widget, *(QIcon*)icon, label_qstr);
	return __r;
}
void QTabWidget_removeTab(void* self, int index) {
	((QTabWidget*)self)->removeTab(index);
}
int QTabWidget_isTabEnabled(void* self, int index) {
	auto&& __r = ((QTabWidget*)self)->isTabEnabled(index);
	return __r;
}
void QTabWidget_setTabEnabled(void* self, int index, int enabled) {
	((QTabWidget*)self)->setTabEnabled(index, enabled);
}
int QTabWidget_isTabVisible(void* self, int index) {
	auto&& __r = ((QTabWidget*)self)->isTabVisible(index);
	return __r;
}
void QTabWidget_setTabVisible(void* self, int index, int visible) {
	((QTabWidget*)self)->setTabVisible(index, visible);
}
char* QTabWidget_tabText(void* self, int index, long long* __out_len) {
	auto&& __r = ((QTabWidget*)self)->tabText(index);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QTabWidget_setTabText(void* self, int index, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	((QTabWidget*)self)->setTabText(index, text_qstr);
}
void* QTabWidget_tabIcon(void* self, int index) {
	auto&& __r = ((QTabWidget*)self)->tabIcon(index);
	return new QIcon(__r);
}
void QTabWidget_setTabIcon(void* self, int index, void* icon) {
	((QTabWidget*)self)->setTabIcon(index, *(QIcon*)icon);
}
void QTabWidget_setTabToolTip(void* self, int index, const char* tip, long long tip_len) {
	QString tip_qstr = QString::fromUtf8(tip, tip_len);
	((QTabWidget*)self)->setTabToolTip(index, tip_qstr);
}
char* QTabWidget_tabToolTip(void* self, int index, long long* __out_len) {
	auto&& __r = ((QTabWidget*)self)->tabToolTip(index);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QTabWidget_setTabWhatsThis(void* self, int index, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	((QTabWidget*)self)->setTabWhatsThis(index, text_qstr);
}
char* QTabWidget_tabWhatsThis(void* self, int index, long long* __out_len) {
	auto&& __r = ((QTabWidget*)self)->tabWhatsThis(index);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QTabWidget_currentIndex(void* self) {
	auto&& __r = ((QTabWidget*)self)->currentIndex();
	return __r;
}
void* QTabWidget_currentWidget(void* self) {
	auto&& __r = ((QTabWidget*)self)->currentWidget();
	return (void*)(__r);
}
void* QTabWidget_widget(void* self, int index) {
	auto&& __r = ((QTabWidget*)self)->widget(index);
	return (void*)(__r);
}
int QTabWidget_indexOf(void* self, void* widget) {
	auto&& __r = ((QTabWidget*)self)->indexOf((QWidget*)widget);
	return __r;
}
int QTabWidget_count(void* self) {
	auto&& __r = ((QTabWidget*)self)->count();
	return __r;
}
int QTabWidget_tabPosition(void* self) {
	auto&& __r = ((QTabWidget*)self)->tabPosition();
	return (int)(__r);
}
void QTabWidget_setTabPosition(void* self, int position) {
	((QTabWidget*)self)->setTabPosition((QTabWidget::TabPosition)(position));
}
int QTabWidget_tabsClosable(void* self) {
	auto&& __r = ((QTabWidget*)self)->tabsClosable();
	return __r;
}
void QTabWidget_setTabsClosable(void* self, int closeable) {
	((QTabWidget*)self)->setTabsClosable(closeable);
}
int QTabWidget_isMovable(void* self) {
	auto&& __r = ((QTabWidget*)self)->isMovable();
	return __r;
}
void QTabWidget_setMovable(void* self, int movable) {
	((QTabWidget*)self)->setMovable(movable);
}
int QTabWidget_tabShape(void* self) {
	auto&& __r = ((QTabWidget*)self)->tabShape();
	return (int)(__r);
}
void QTabWidget_setTabShape(void* self, int s) {
	((QTabWidget*)self)->setTabShape((QTabWidget::TabShape)(s));
}
void QTabWidget_setCornerWidget(void* self, void* w) {
	((QTabWidget*)self)->setCornerWidget((QWidget*)w);
}
void* QTabWidget_cornerWidget(void* self) {
	auto&& __r = ((QTabWidget*)self)->cornerWidget();
	return (void*)(__r);
}
int QTabWidget_elideMode(void* self) {
	auto&& __r = ((QTabWidget*)self)->elideMode();
	return (int)(__r);
}
void QTabWidget_setElideMode(void* self, int mode) {
	((QTabWidget*)self)->setElideMode((Qt::TextElideMode)(mode));
}
void* QTabWidget_iconSize(void* self) {
	auto&& __r = ((QTabWidget*)self)->iconSize();
	return new QSize(__r);
}
void QTabWidget_setIconSize(void* self, void* size) {
	((QTabWidget*)self)->setIconSize(*(QSize*)size);
}
int QTabWidget_usesScrollButtons(void* self) {
	auto&& __r = ((QTabWidget*)self)->usesScrollButtons();
	return __r;
}
void QTabWidget_setUsesScrollButtons(void* self, int useButtons) {
	((QTabWidget*)self)->setUsesScrollButtons(useButtons);
}
int QTabWidget_documentMode(void* self) {
	auto&& __r = ((QTabWidget*)self)->documentMode();
	return __r;
}
void QTabWidget_setDocumentMode(void* self, int set) {
	((QTabWidget*)self)->setDocumentMode(set);
}
int QTabWidget_tabBarAutoHide(void* self) {
	auto&& __r = ((QTabWidget*)self)->tabBarAutoHide();
	return __r;
}
void QTabWidget_setTabBarAutoHide(void* self, int enabled) {
	((QTabWidget*)self)->setTabBarAutoHide(enabled);
}
void QTabWidget_clear(void* self) {
	((QTabWidget*)self)->clear();
}
void* QTabWidget_tabBar(void* self) {
	auto&& __r = ((QTabWidget*)self)->tabBar();
	return (void*)(__r);
}
void QTabWidget_setCurrentIndex(void* self, int index) {
	((QTabWidget*)self)->setCurrentIndex(index);
}
void QTabWidget_setCurrentWidget(void* self, void* widget) {
	((QTabWidget*)self)->setCurrentWidget((QWidget*)widget);
}
void QTabWidget_currentChanged(void* self, int index) {
	((QTabWidget*)self)->currentChanged(index);
}
void QTabWidget_tabCloseRequested(void* self, int index) {
	((QTabWidget*)self)->tabCloseRequested(index);
}
void QTabWidget_tabBarClicked(void* self, int index) {
	((QTabWidget*)self)->tabBarClicked(index);
}
void QTabWidget_tabBarDoubleClicked(void* self, int index) {
	((QTabWidget*)self)->tabBarDoubleClicked(index);
}
void QTabWidget_setTabBar(void* self, void* arg1) {
	((QTabWidget__Go*)self)->setTabBar((QTabBar*)arg1);
}
void QTabWidget_setTabOrder(void* arg1, void* arg2) {
	QTabWidget::setTabOrder((QWidget*)arg1, (QWidget*)arg2);
}
void* QTabWidget_mouseGrabber() {
	auto&& __r = QTabWidget::mouseGrabber();
	return (void*)(__r);
}
void* QTabWidget_keyboardGrabber() {
	auto&& __r = QTabWidget::keyboardGrabber();
	return (void*)(__r);
}
void* QTabWidget_find(unsigned long long arg1) {
	auto&& __r = QTabWidget::find(arg1);
	return (void*)(__r);
}
void* QTabWidget_createWindowContainer(void* window) {
	auto&& __r = QTabWidget::createWindowContainer((QWindow*)window);
	return (void*)(__r);
}
void QTabWidget_updateMicroFocus(void* self) {
	((QTabWidget__Go*)self)->updateMicroFocus();
}
void QTabWidget_create(void* self) {
	((QTabWidget__Go*)self)->create();
}
void QTabWidget_destroy(void* self) {
	((QTabWidget__Go*)self)->destroy();
}
int QTabWidget_focusNextChild(void* self) {
	auto&& __r = ((QTabWidget__Go*)self)->focusNextChild();
	return __r;
}
int QTabWidget_focusPreviousChild(void* self) {
	auto&& __r = ((QTabWidget__Go*)self)->focusPreviousChild();
	return __r;
}
void* QTabWidget_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QTabWidget::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QTabWidget_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QTabWidget::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QTabWidget_sender(void* self) {
	auto&& __r = ((QTabWidget__Go*)self)->sender();
	return (void*)(__r);
}
int QTabWidget_senderSignalIndex(void* self) {
	auto&& __r = ((QTabWidget__Go*)self)->senderSignalIndex();
	return __r;
}
int QTabWidget_receivers(void* self, const char* signal) {
	auto&& __r = ((QTabWidget__Go*)self)->receivers(signal);
	return __r;
}
int QTabWidget_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QTabWidget__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
double QTabWidget_devicePixelRatioFScale() {
	auto&& __r = QTabWidget::devicePixelRatioFScale();
	return __r;
}
int QTabWidget_encodeMetricF(int metric, double value) {
	auto&& __r = QTabWidget::encodeMetricF((QPaintDevice::PaintDeviceMetric)(metric), value);
	return __r;
}
double QTabWidget_getDecodedMetricF(void* self, int metricA, int metricB) {
	auto&& __r = ((QTabWidget__Go*)self)->getDecodedMetricF((QPaintDevice::PaintDeviceMetric)(metricA), (QPaintDevice::PaintDeviceMetric)(metricB));
	return __r;
}
char* QTabWidget_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QTabWidget::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QTabWidget_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QTabWidget::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QTabWidget_setCornerWidget2(void* self, void* w, int corner) {
	((QTabWidget*)self)->setCornerWidget((QWidget*)w, (Qt::Corner)(corner));
}
void* QTabWidget_cornerWidget2(void* self, int corner) {
	auto&& __r = ((QTabWidget*)self)->cornerWidget((Qt::Corner)(corner));
	return (void*)(__r);
}
void* QTabWidget_createWindowContainer2(void* window, void* parent) {
	auto&& __r = QTabWidget::createWindowContainer((QWindow*)window, (QWidget*)parent);
	return (void*)(__r);
}
void* QTabWidget_createWindowContainer3(void* window, void* parent, int flags) {
	auto&& __r = QTabWidget::createWindowContainer((QWindow*)window, (QWidget*)parent, QFlags<Qt::WindowType>(flags));
	return (void*)(__r);
}
void QTabWidget_updateMicroFocus2(void* self, int query) {
	((QTabWidget__Go*)self)->updateMicroFocus((Qt::InputMethodQuery)(query));
}
void QTabWidget_create2(void* self, unsigned long long arg1) {
	((QTabWidget__Go*)self)->create(arg1);
}
void QTabWidget_create3(void* self, unsigned long long arg1, int initializeWindow) {
	((QTabWidget__Go*)self)->create(arg1, initializeWindow);
}
void QTabWidget_create4(void* self, unsigned long long arg1, int initializeWindow, int destroyOldWindow) {
	((QTabWidget__Go*)self)->create(arg1, initializeWindow, destroyOldWindow);
}
void QTabWidget_destroy2(void* self, int destroyWindow) {
	((QTabWidget__Go*)self)->destroy(destroyWindow);
}
void QTabWidget_destroy3(void* self, int destroyWindow, int destroySubWindows) {
	((QTabWidget__Go*)self)->destroy(destroyWindow, destroySubWindows);
}
void* QTabWidget_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QTabWidget::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QTabWidget_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QTabWidget::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QTabWidget_metaObject_default(void* self) {
	auto&& __r = ((QTabWidget*)self)->QTabWidget::metaObject();
	return (void*)(__r);
}
void* QTabWidget_sizeHint_default(void* self) {
	auto&& __r = ((QTabWidget*)self)->QTabWidget::sizeHint();
	return new QSize(__r);
}
void* QTabWidget_minimumSizeHint_default(void* self) {
	auto&& __r = ((QTabWidget*)self)->QTabWidget::minimumSizeHint();
	return new QSize(__r);
}
int QTabWidget_heightForWidth_default(void* self, int width) {
	auto&& __r = ((QTabWidget*)self)->QTabWidget::heightForWidth(width);
	return __r;
}
int QTabWidget_hasHeightForWidth_default(void* self) {
	auto&& __r = ((QTabWidget*)self)->QTabWidget::hasHeightForWidth();
	return __r;
}
void QTabWidget_tabInserted_default(void* self, int index) {
	((QTabWidget__Go*)self)->cppgoBase_tabInserted(index);
}
void QTabWidget_tabRemoved_default(void* self, int index) {
	((QTabWidget__Go*)self)->cppgoBase_tabRemoved(index);
}
void QTabWidget_showEvent_default(void* self, void* arg1) {
	((QTabWidget__Go*)self)->cppgoBase_showEvent((QShowEvent*)arg1);
}
void QTabWidget_resizeEvent_default(void* self, void* arg1) {
	((QTabWidget__Go*)self)->cppgoBase_resizeEvent((QResizeEvent*)arg1);
}
void QTabWidget_keyPressEvent_default(void* self, void* arg1) {
	((QTabWidget__Go*)self)->cppgoBase_keyPressEvent((QKeyEvent*)arg1);
}
void QTabWidget_paintEvent_default(void* self, void* arg1) {
	((QTabWidget__Go*)self)->cppgoBase_paintEvent((QPaintEvent*)arg1);
}
void QTabWidget_changeEvent_default(void* self, void* arg1) {
	((QTabWidget__Go*)self)->cppgoBase_changeEvent((QEvent*)arg1);
}
int QTabWidget_event_default(void* self, void* arg1) {
	auto&& __r = ((QTabWidget__Go*)self)->cppgoBase_event((QEvent*)arg1);
	return __r;
}
void QTabWidget_initStyleOption_default(void* self, void* option) {
	((QTabWidget__Go*)self)->cppgoBase_initStyleOption((QStyleOptionTabWidgetFrame*)option);
}
int QTabWidget_devType_default(void* self) {
	auto&& __r = ((QTabWidget*)self)->QTabWidget::devType();
	return __r;
}
void QTabWidget_setVisible_default(void* self, int visible) {
	((QTabWidget*)self)->QTabWidget::setVisible(visible);
}
void* QTabWidget_paintEngine_default(void* self) {
	auto&& __r = ((QTabWidget*)self)->QTabWidget::paintEngine();
	return (void*)(__r);
}
void QTabWidget_mousePressEvent_default(void* self, void* event) {
	((QTabWidget__Go*)self)->cppgoBase_mousePressEvent((QMouseEvent*)event);
}
void QTabWidget_mouseReleaseEvent_default(void* self, void* event) {
	((QTabWidget__Go*)self)->cppgoBase_mouseReleaseEvent((QMouseEvent*)event);
}
void QTabWidget_mouseDoubleClickEvent_default(void* self, void* event) {
	((QTabWidget__Go*)self)->cppgoBase_mouseDoubleClickEvent((QMouseEvent*)event);
}
void QTabWidget_mouseMoveEvent_default(void* self, void* event) {
	((QTabWidget__Go*)self)->cppgoBase_mouseMoveEvent((QMouseEvent*)event);
}
void QTabWidget_wheelEvent_default(void* self, void* event) {
	((QTabWidget__Go*)self)->cppgoBase_wheelEvent((QWheelEvent*)event);
}
void QTabWidget_keyReleaseEvent_default(void* self, void* event) {
	((QTabWidget__Go*)self)->cppgoBase_keyReleaseEvent((QKeyEvent*)event);
}
void QTabWidget_focusInEvent_default(void* self, void* event) {
	((QTabWidget__Go*)self)->cppgoBase_focusInEvent((QFocusEvent*)event);
}
void QTabWidget_focusOutEvent_default(void* self, void* event) {
	((QTabWidget__Go*)self)->cppgoBase_focusOutEvent((QFocusEvent*)event);
}
void QTabWidget_enterEvent_default(void* self, void* event) {
	((QTabWidget__Go*)self)->cppgoBase_enterEvent((QEnterEvent*)event);
}
void QTabWidget_leaveEvent_default(void* self, void* event) {
	((QTabWidget__Go*)self)->cppgoBase_leaveEvent((QEvent*)event);
}
void QTabWidget_moveEvent_default(void* self, void* event) {
	((QTabWidget__Go*)self)->cppgoBase_moveEvent((QMoveEvent*)event);
}
void QTabWidget_closeEvent_default(void* self, void* event) {
	((QTabWidget__Go*)self)->cppgoBase_closeEvent((QCloseEvent*)event);
}
void QTabWidget_contextMenuEvent_default(void* self, void* event) {
	((QTabWidget__Go*)self)->cppgoBase_contextMenuEvent((QContextMenuEvent*)event);
}
void QTabWidget_tabletEvent_default(void* self, void* event) {
	((QTabWidget__Go*)self)->cppgoBase_tabletEvent((QTabletEvent*)event);
}
void QTabWidget_actionEvent_default(void* self, void* event) {
	((QTabWidget__Go*)self)->cppgoBase_actionEvent((QActionEvent*)event);
}
void QTabWidget_dragEnterEvent_default(void* self, void* event) {
	((QTabWidget__Go*)self)->cppgoBase_dragEnterEvent((QDragEnterEvent*)event);
}
void QTabWidget_dragMoveEvent_default(void* self, void* event) {
	((QTabWidget__Go*)self)->cppgoBase_dragMoveEvent((QDragMoveEvent*)event);
}
void QTabWidget_dragLeaveEvent_default(void* self, void* event) {
	((QTabWidget__Go*)self)->cppgoBase_dragLeaveEvent((QDragLeaveEvent*)event);
}
void QTabWidget_dropEvent_default(void* self, void* event) {
	((QTabWidget__Go*)self)->cppgoBase_dropEvent((QDropEvent*)event);
}
void QTabWidget_hideEvent_default(void* self, void* event) {
	((QTabWidget__Go*)self)->cppgoBase_hideEvent((QHideEvent*)event);
}
int QTabWidget_metric_default(void* self, int arg1) {
	auto&& __r = ((QTabWidget__Go*)self)->cppgoBase_metric(arg1);
	return __r;
}
void QTabWidget_initPainter_default(void* self, void* painter) {
	((QTabWidget__Go*)self)->cppgoBase_initPainter((QPainter*)painter);
}
void* QTabWidget_redirected_default(void* self, void* offset) {
	auto&& __r = ((QTabWidget__Go*)self)->cppgoBase_redirected((QPoint*)offset);
	return (void*)(__r);
}
void* QTabWidget_sharedPainter_default(void* self) {
	auto&& __r = ((QTabWidget__Go*)self)->cppgoBase_sharedPainter();
	return (void*)(__r);
}
void QTabWidget_inputMethodEvent_default(void* self, void* arg1) {
	((QTabWidget__Go*)self)->cppgoBase_inputMethodEvent((QInputMethodEvent*)arg1);
}
void* QTabWidget_inputMethodQuery_default(void* self, int arg1) {
	auto&& __r = ((QTabWidget*)self)->QTabWidget::inputMethodQuery((Qt::InputMethodQuery)(arg1));
	return new QVariant(__r);
}
int QTabWidget_focusNextPrevChild_default(void* self, int next) {
	auto&& __r = ((QTabWidget__Go*)self)->cppgoBase_focusNextPrevChild(next);
	return __r;
}
int QTabWidget_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QTabWidget*)self)->QTabWidget::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QTabWidget_timerEvent_default(void* self, void* event) {
	((QTabWidget__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QTabWidget_childEvent_default(void* self, void* event) {
	((QTabWidget__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QTabWidget_customEvent_default(void* self, void* event) {
	((QTabWidget__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QTabWidget_connectNotify_default(void* self, void* signal) {
	((QTabWidget__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QTabWidget_disconnectNotify_default(void* self, void* signal) {
	((QTabWidget__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QTabWidget_connect_currentChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTabWidget*)self, &QTabWidget::currentChanged, (QTabWidget*)self, [__id](int index){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, index);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTabWidget);
	__conns_QTabWidget[__id] = __c;
}
void QTabWidget_connect_tabCloseRequested(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTabWidget*)self, &QTabWidget::tabCloseRequested, (QTabWidget*)self, [__id](int index){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, index);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTabWidget);
	__conns_QTabWidget[__id] = __c;
}
void QTabWidget_connect_tabBarClicked(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTabWidget*)self, &QTabWidget::tabBarClicked, (QTabWidget*)self, [__id](int index){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, index);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTabWidget);
	__conns_QTabWidget[__id] = __c;
}
void QTabWidget_connect_tabBarDoubleClicked(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTabWidget*)self, &QTabWidget::tabBarDoubleClicked, (QTabWidget*)self, [__id](int index){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, index);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTabWidget);
	__conns_QTabWidget[__id] = __c;
}
void QTabWidget_connect_windowTitleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTabWidget*)self, &QTabWidget::windowTitleChanged, (QTabWidget*)self, [__id](const QString& title){
		QByteArray title_ba = title.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)title_ba.constData(), title_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTabWidget);
	__conns_QTabWidget[__id] = __c;
}
void QTabWidget_connect_windowIconChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTabWidget*)self, &QTabWidget::windowIconChanged, (QTabWidget*)self, [__id](const QIcon& icon){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&icon);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTabWidget);
	__conns_QTabWidget[__id] = __c;
}
void QTabWidget_connect_windowIconTextChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTabWidget*)self, &QTabWidget::windowIconTextChanged, (QTabWidget*)self, [__id](const QString& iconText){
		QByteArray iconText_ba = iconText.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)iconText_ba.constData(), iconText_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTabWidget);
	__conns_QTabWidget[__id] = __c;
}
void QTabWidget_connect_customContextMenuRequested(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTabWidget*)self, &QTabWidget::customContextMenuRequested, (QTabWidget*)self, [__id](const QPoint& pos){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&pos);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTabWidget);
	__conns_QTabWidget[__id] = __c;
}
void QTabWidget_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTabWidget*)self, &QTabWidget::destroyed, (QTabWidget*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTabWidget);
	__conns_QTabWidget[__id] = __c;
}
void QTabWidget_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTabWidget*)self, &QTabWidget::destroyed, (QTabWidget*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTabWidget);
	__conns_QTabWidget[__id] = __c;
}
void QTabWidget_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTabWidget*)self, &QTabWidget::objectNameChanged, (QTabWidget*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTabWidget);
	__conns_QTabWidget[__id] = __c;
}
void QTabWidget__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QTabWidget);
		auto __it = __conns_QTabWidget.find(id);
		if (__it == __conns_QTabWidget.end()) return;
		__c = __it->second;
		__conns_QTabWidget.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QTabWidget__Go_new(long long id) { return new QTabWidget__Go(id); }
void QTabWidget_delete(void* self) { delete (QTabWidget*)self; }

}

#include <QtWidgets/qstyleoption.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qrect.h>
#include <QtCore/qstring.h>
#include <QtGui/qicon.h>
#include <QtCore/qsize.h>
#include <QtCore/qnamespace.h>
#include "gen_qstyleoptioncombobox.h"

extern "C" {

void* QStyleOptionComboBox_base(void* self) { return (void*)static_cast<QStyleOptionComplex*>((QStyleOptionComboBox*)self); }
void QStyleOptionComboBox_chain(void* self, void** out) {
	QStyleOptionComboBox* __p = (QStyleOptionComboBox*)self;
	QStyleOptionComplex* __b0 = static_cast<QStyleOptionComplex*>(__p); out[0] = (void*)__b0;
	QStyleOption* __b1 = static_cast<QStyleOption*>(__b0); out[1] = (void*)__b1;
}
void* QStyleOptionComboBox_new() {
	return new QStyleOptionComboBox();
}
void* QStyleOptionComboBox_new2(void* other) {
	return new QStyleOptionComboBox(*(QStyleOptionComboBox*)other);
}
void* QStyleOptionComboBox_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QStyleOptionComboBox*)self)->operator=(*(QStyleOptionComboBox*)arg1);
	return (void*)&(__r);
}
int QStyleOptionComboBox_editable_get(void* self) {
	auto __r = ((QStyleOptionComboBox*)self)->editable;
	return __r;
}
void QStyleOptionComboBox_editable_set(void* self, int editable) {
	((QStyleOptionComboBox*)self)->editable = editable;
}
void* QStyleOptionComboBox_popupRect_get(void* self) {
	auto __r = ((QStyleOptionComboBox*)self)->popupRect;
	return new QRect(__r);
}
void QStyleOptionComboBox_popupRect_set(void* self, void* popupRect) {
	((QStyleOptionComboBox*)self)->popupRect = *(QRect*)popupRect;
}
int QStyleOptionComboBox_frame_get(void* self) {
	auto __r = ((QStyleOptionComboBox*)self)->frame;
	return __r;
}
void QStyleOptionComboBox_frame_set(void* self, int frame) {
	((QStyleOptionComboBox*)self)->frame = frame;
}
char* QStyleOptionComboBox_currentText_get(void* self, long long* __out_len) {
	auto __r = ((QStyleOptionComboBox*)self)->currentText;
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QStyleOptionComboBox_currentText_set(void* self, const char* currentText, long long currentText_len) {
	QString currentText_qstr = QString::fromUtf8(currentText, currentText_len);
	((QStyleOptionComboBox*)self)->currentText = currentText_qstr;
}
void* QStyleOptionComboBox_currentIcon_get(void* self) {
	auto __r = ((QStyleOptionComboBox*)self)->currentIcon;
	return new QIcon(__r);
}
void QStyleOptionComboBox_currentIcon_set(void* self, void* currentIcon) {
	((QStyleOptionComboBox*)self)->currentIcon = *(QIcon*)currentIcon;
}
void* QStyleOptionComboBox_iconSize_get(void* self) {
	auto __r = ((QStyleOptionComboBox*)self)->iconSize;
	return new QSize(__r);
}
void QStyleOptionComboBox_iconSize_set(void* self, void* iconSize) {
	((QStyleOptionComboBox*)self)->iconSize = *(QSize*)iconSize;
}
int QStyleOptionComboBox_textAlignment_get(void* self) {
	auto __r = ((QStyleOptionComboBox*)self)->textAlignment;
	return (int)(__r);
}
void QStyleOptionComboBox_textAlignment_set(void* self, int textAlignment) {
	((QStyleOptionComboBox*)self)->textAlignment = QFlags<Qt::AlignmentFlag>(textAlignment);
}
void QStyleOptionComboBox_delete(void* self) { delete (QStyleOptionComboBox*)self; }

}

#include <QtWidgets/qcalendarwidget.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qsize.h>
#include <QtCore/qdatetime.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qcalendar.h>
#include <QtGui/qtextformat.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qobject.h>
#include <QtGui/qevent.h>
#include <QtGui/qpainter.h>
#include <QtCore/qrect.h>
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
#include "gen_qcalendarwidget.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QCalendarWidget__Go : public QCalendarWidget {
public:
	long long __goId;
	unsigned char __goOn[49] = {};
	template<typename... A> QCalendarWidget__Go(long long id, A&&... a) : QCalendarWidget(std::forward<A>(a)...), __goId(id) {}
	using QCalendarWidget::updateCell;
	using QCalendarWidget::updateCells;
	using QCalendarWidget::updateMicroFocus;
	using QCalendarWidget::create;
	using QCalendarWidget::destroy;
	using QCalendarWidget::focusNextChild;
	using QCalendarWidget::focusPreviousChild;
	using QCalendarWidget::sender;
	using QCalendarWidget::senderSignalIndex;
	using QCalendarWidget::receivers;
	using QCalendarWidget::isSignalConnected;
	using QCalendarWidget::getDecodedMetricF;
	bool cppgoBase_event(QEvent* __a0) { return QCalendarWidget::event(__a0); }
	bool cppgoBase_eventFilter(QObject* __a0, QEvent* __a1) { return QCalendarWidget::eventFilter(__a0, __a1); }
	void cppgoBase_mousePressEvent(QMouseEvent* __a0) { QCalendarWidget::mousePressEvent(__a0); }
	void cppgoBase_resizeEvent(QResizeEvent* __a0) { QCalendarWidget::resizeEvent(__a0); }
	void cppgoBase_keyPressEvent(QKeyEvent* __a0) { QCalendarWidget::keyPressEvent(__a0); }
	void cppgoBase_paintCell(QPainter* __a0, const QRect& __a1, QDate __a2) { QCalendarWidget::paintCell(__a0, __a1, __a2); }
	void cppgoBase_mouseReleaseEvent(QMouseEvent* __a0) { QCalendarWidget::mouseReleaseEvent(__a0); }
	void cppgoBase_mouseDoubleClickEvent(QMouseEvent* __a0) { QCalendarWidget::mouseDoubleClickEvent(__a0); }
	void cppgoBase_mouseMoveEvent(QMouseEvent* __a0) { QCalendarWidget::mouseMoveEvent(__a0); }
	void cppgoBase_wheelEvent(QWheelEvent* __a0) { QCalendarWidget::wheelEvent(__a0); }
	void cppgoBase_keyReleaseEvent(QKeyEvent* __a0) { QCalendarWidget::keyReleaseEvent(__a0); }
	void cppgoBase_focusInEvent(QFocusEvent* __a0) { QCalendarWidget::focusInEvent(__a0); }
	void cppgoBase_focusOutEvent(QFocusEvent* __a0) { QCalendarWidget::focusOutEvent(__a0); }
	void cppgoBase_enterEvent(QEnterEvent* __a0) { QCalendarWidget::enterEvent(__a0); }
	void cppgoBase_leaveEvent(QEvent* __a0) { QCalendarWidget::leaveEvent(__a0); }
	void cppgoBase_paintEvent(QPaintEvent* __a0) { QCalendarWidget::paintEvent(__a0); }
	void cppgoBase_moveEvent(QMoveEvent* __a0) { QCalendarWidget::moveEvent(__a0); }
	void cppgoBase_closeEvent(QCloseEvent* __a0) { QCalendarWidget::closeEvent(__a0); }
	void cppgoBase_contextMenuEvent(QContextMenuEvent* __a0) { QCalendarWidget::contextMenuEvent(__a0); }
	void cppgoBase_tabletEvent(QTabletEvent* __a0) { QCalendarWidget::tabletEvent(__a0); }
	void cppgoBase_actionEvent(QActionEvent* __a0) { QCalendarWidget::actionEvent(__a0); }
	void cppgoBase_dragEnterEvent(QDragEnterEvent* __a0) { QCalendarWidget::dragEnterEvent(__a0); }
	void cppgoBase_dragMoveEvent(QDragMoveEvent* __a0) { QCalendarWidget::dragMoveEvent(__a0); }
	void cppgoBase_dragLeaveEvent(QDragLeaveEvent* __a0) { QCalendarWidget::dragLeaveEvent(__a0); }
	void cppgoBase_dropEvent(QDropEvent* __a0) { QCalendarWidget::dropEvent(__a0); }
	void cppgoBase_showEvent(QShowEvent* __a0) { QCalendarWidget::showEvent(__a0); }
	void cppgoBase_hideEvent(QHideEvent* __a0) { QCalendarWidget::hideEvent(__a0); }
	void cppgoBase_changeEvent(QEvent* __a0) { QCalendarWidget::changeEvent(__a0); }
	int cppgoBase_metric(int __a0) { return QCalendarWidget::metric((QPaintDevice::PaintDeviceMetric)__a0); }
	void cppgoBase_initPainter(QPainter* __a0) { QCalendarWidget::initPainter(__a0); }
	QPaintDevice* cppgoBase_redirected(QPoint* __a0) { return QCalendarWidget::redirected(__a0); }
	QPainter* cppgoBase_sharedPainter() { return QCalendarWidget::sharedPainter(); }
	void cppgoBase_inputMethodEvent(QInputMethodEvent* __a0) { QCalendarWidget::inputMethodEvent(__a0); }
	bool cppgoBase_focusNextPrevChild(bool __a0) { return QCalendarWidget::focusNextPrevChild(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QCalendarWidget::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QCalendarWidget::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QCalendarWidget::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QCalendarWidget::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QCalendarWidget::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QCalendarWidget::metaObject();
	}
	QSize sizeHint() const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 729690261457684101LL, &__h); if (__h) return *(QSize*)__r; }
		return QCalendarWidget::sizeHint();
	}
	QSize minimumSizeHint() const override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 7049264599837532491LL, &__h); if (__h) return *(QSize*)__r; }
		return QCalendarWidget::minimumSizeHint();
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QCalendarWidget::event(event);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QCalendarWidget::eventFilter(watched, event);
	}
	void mousePressEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -332090898078543319LL, (void*)event, &__h); if (__h) return; }
		QCalendarWidget::mousePressEvent(event);
	}
	void resizeEvent(QResizeEvent* event) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4302917285655014661LL, (void*)event, &__h); if (__h) return; }
		QCalendarWidget::resizeEvent(event);
	}
	void keyPressEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1023791977679793185LL, (void*)event, &__h); if (__h) return; }
		QCalendarWidget::keyPressEvent(event);
	}
	void paintCell(QPainter* painter, const QRect& rect, QDate date) const override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void_voidp_voidp_voidp) { ((void(*)(long long, long long, void*, void*, void*, int*))cppgo_virt_void_voidp_voidp_voidp)(__goId, 8082959910766380751LL, (void*)painter, (void*)&rect, (void*)&date, &__h); if (__h) return; }
		QCalendarWidget::paintCell(painter, rect, date);
	}
	int devType() const override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8659679893868524044LL, &__h); if (__h) return __r; }
		return QCalendarWidget::devType();
	}
	void setVisible(bool visible) override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 2110182719375245575LL, visible, &__h); if (__h) return; }
		QCalendarWidget::setVisible(visible);
	}
	int heightForWidth(int arg1) const override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 4311534060495737339LL, arg1, &__h); if (__h) return __r; }
		return QCalendarWidget::heightForWidth(arg1);
	}
	bool hasHeightForWidth() const override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8266762397995423915LL, &__h); if (__h) return __r; }
		return QCalendarWidget::hasHeightForWidth();
	}
	QPaintEngine* paintEngine() const override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 6398007404756346381LL, &__h); if (__h) return (QPaintEngine*)__r; }
		return QCalendarWidget::paintEngine();
	}
	void mouseReleaseEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4415920049297795213LL, (void*)event, &__h); if (__h) return; }
		QCalendarWidget::mouseReleaseEvent(event);
	}
	void mouseDoubleClickEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4029073905697023887LL, (void*)event, &__h); if (__h) return; }
		QCalendarWidget::mouseDoubleClickEvent(event);
	}
	void mouseMoveEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4255494004099806581LL, (void*)event, &__h); if (__h) return; }
		QCalendarWidget::mouseMoveEvent(event);
	}
	void wheelEvent(QWheelEvent* event) override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1919988340934511316LL, (void*)event, &__h); if (__h) return; }
		QCalendarWidget::wheelEvent(event);
	}
	void keyReleaseEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 772598475950050299LL, (void*)event, &__h); if (__h) return; }
		QCalendarWidget::keyReleaseEvent(event);
	}
	void focusInEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 975043767033552552LL, (void*)event, &__h); if (__h) return; }
		QCalendarWidget::focusInEvent(event);
	}
	void focusOutEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -385596010576048993LL, (void*)event, &__h); if (__h) return; }
		QCalendarWidget::focusOutEvent(event);
	}
	void enterEvent(QEnterEvent* event) override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6316270346668237895LL, (void*)event, &__h); if (__h) return; }
		QCalendarWidget::enterEvent(event);
	}
	void leaveEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -551070009623711640LL, (void*)event, &__h); if (__h) return; }
		QCalendarWidget::leaveEvent(event);
	}
	void paintEvent(QPaintEvent* event) override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5290442163167458465LL, (void*)event, &__h); if (__h) return; }
		QCalendarWidget::paintEvent(event);
	}
	void moveEvent(QMoveEvent* event) override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -9161949150489525396LL, (void*)event, &__h); if (__h) return; }
		QCalendarWidget::moveEvent(event);
	}
	void closeEvent(QCloseEvent* event) override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1989937050019464783LL, (void*)event, &__h); if (__h) return; }
		QCalendarWidget::closeEvent(event);
	}
	void contextMenuEvent(QContextMenuEvent* event) override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6698043222227006015LL, (void*)event, &__h); if (__h) return; }
		QCalendarWidget::contextMenuEvent(event);
	}
	void tabletEvent(QTabletEvent* event) override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -212068866860701141LL, (void*)event, &__h); if (__h) return; }
		QCalendarWidget::tabletEvent(event);
	}
	void actionEvent(QActionEvent* event) override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1697444140088555581LL, (void*)event, &__h); if (__h) return; }
		QCalendarWidget::actionEvent(event);
	}
	void dragEnterEvent(QDragEnterEvent* event) override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 206122715919917589LL, (void*)event, &__h); if (__h) return; }
		QCalendarWidget::dragEnterEvent(event);
	}
	void dragMoveEvent(QDragMoveEvent* event) override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1781545137352745978LL, (void*)event, &__h); if (__h) return; }
		QCalendarWidget::dragMoveEvent(event);
	}
	void dragLeaveEvent(QDragLeaveEvent* event) override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -3139904158373725006LL, (void*)event, &__h); if (__h) return; }
		QCalendarWidget::dragLeaveEvent(event);
	}
	void dropEvent(QDropEvent* event) override {
		int __h = 0;
		if (__goOn[32] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 9107058682754591716LL, (void*)event, &__h); if (__h) return; }
		QCalendarWidget::dropEvent(event);
	}
	void showEvent(QShowEvent* event) override {
		int __h = 0;
		if (__goOn[33] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6825369174751916LL, (void*)event, &__h); if (__h) return; }
		QCalendarWidget::showEvent(event);
	}
	void hideEvent(QHideEvent* event) override {
		int __h = 0;
		if (__goOn[34] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -199718707876719129LL, (void*)event, &__h); if (__h) return; }
		QCalendarWidget::hideEvent(event);
	}
	bool nativeEvent(const QByteArray& eventType, void* message, qintptr* result) override {
		int __h = 0;
		if (__goOn[35] && cppgo_virt_int_voidp_voidp_longlongp) { int __r = ((int(*)(long long, long long, void*, void*, long long*, int*))cppgo_virt_int_voidp_voidp_longlongp)(__goId, -2039166586246443706LL, (void*)&eventType, (void*)message, (long long*)result, &__h); if (__h) return __r; }
		return QCalendarWidget::nativeEvent(eventType, message, result);
	}
	void changeEvent(QEvent* arg1) override {
		int __h = 0;
		if (__goOn[36] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4959376063826165943LL, (void*)arg1, &__h); if (__h) return; }
		QCalendarWidget::changeEvent(arg1);
	}
	int metric(QPaintDevice::PaintDeviceMetric arg1) const override {
		int __h = 0;
		if (__goOn[37] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 427629649267311777LL, (int)(arg1), &__h); if (__h) return __r; }
		return QCalendarWidget::metric(arg1);
	}
	void initPainter(QPainter* painter) const override {
		int __h = 0;
		if (__goOn[38] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6477500296980050210LL, (void*)painter, &__h); if (__h) return; }
		QCalendarWidget::initPainter(painter);
	}
	QPaintDevice* redirected(QPoint* offset) const override {
		int __h = 0;
		if (__goOn[39] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 4127240633310655048LL, (void*)offset, &__h); if (__h) return (QPaintDevice*)__r; }
		return QCalendarWidget::redirected(offset);
	}
	QPainter* sharedPainter() const override {
		int __h = 0;
		if (__goOn[40] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 5792771328741774991LL, &__h); if (__h) return (QPainter*)__r; }
		return QCalendarWidget::sharedPainter();
	}
	void inputMethodEvent(QInputMethodEvent* arg1) override {
		int __h = 0;
		if (__goOn[41] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5438819661942843222LL, (void*)arg1, &__h); if (__h) return; }
		QCalendarWidget::inputMethodEvent(arg1);
	}
	QVariant inputMethodQuery(Qt::InputMethodQuery arg1) const override {
		int __h = 0;
		if (__goOn[42] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7345853097340709638LL, (int)(arg1), &__h); if (__h) return *(QVariant*)__r; }
		return QCalendarWidget::inputMethodQuery(arg1);
	}
	bool focusNextPrevChild(bool next) override {
		int __h = 0;
		if (__goOn[43] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 5898328517505784191LL, next, &__h); if (__h) return __r; }
		return QCalendarWidget::focusNextPrevChild(next);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[44] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QCalendarWidget::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[45] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QCalendarWidget::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[46] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QCalendarWidget::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[47] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QCalendarWidget::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[48] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QCalendarWidget::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QCalendarWidget;
static std::mutex __connsMu_QCalendarWidget;

extern "C" {

void* QCalendarWidget_base(void* self) { return (void*)static_cast<QWidget*>((QCalendarWidget*)self); }
void QCalendarWidget_chain(void* self, void** out) {
	QCalendarWidget* __p = (QCalendarWidget*)self;
	QWidget* __b0 = static_cast<QWidget*>(__p); out[0] = (void*)__b0;
	QObject* __b1 = static_cast<QObject*>(__b0); out[1] = (void*)__b1;
}
void QCalendarWidget__set_override(void* self, int idx) { ((QCalendarWidget__Go*)self)->__goOn[idx] = 1; }
void* QCalendarWidget_fromQObject(void* o) { return (void*)qobject_cast<QCalendarWidget*>((QObject*)o); }
void* QCalendarWidget_new(long long __goid, void* parent) {
	return new QCalendarWidget__Go(__goid, (QWidget*)parent);
}
void* QCalendarWidget_new2(long long __goid) {
	return new QCalendarWidget__Go(__goid);
}
char* QCalendarWidget_tr(const char* s, long long* __out_len) {
	auto&& __r = QCalendarWidget::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QCalendarWidget_selectedDate(void* self) {
	auto&& __r = ((QCalendarWidget*)self)->selectedDate();
	return new QDate(__r);
}
int QCalendarWidget_yearShown(void* self) {
	auto&& __r = ((QCalendarWidget*)self)->yearShown();
	return __r;
}
int QCalendarWidget_monthShown(void* self) {
	auto&& __r = ((QCalendarWidget*)self)->monthShown();
	return __r;
}
void* QCalendarWidget_minimumDate(void* self) {
	auto&& __r = ((QCalendarWidget*)self)->minimumDate();
	return new QDate(__r);
}
void QCalendarWidget_setMinimumDate(void* self, void* date) {
	((QCalendarWidget*)self)->setMinimumDate(*(QDate*)date);
}
void QCalendarWidget_clearMinimumDate(void* self) {
	((QCalendarWidget*)self)->clearMinimumDate();
}
void* QCalendarWidget_maximumDate(void* self) {
	auto&& __r = ((QCalendarWidget*)self)->maximumDate();
	return new QDate(__r);
}
void QCalendarWidget_setMaximumDate(void* self, void* date) {
	((QCalendarWidget*)self)->setMaximumDate(*(QDate*)date);
}
void QCalendarWidget_clearMaximumDate(void* self) {
	((QCalendarWidget*)self)->clearMaximumDate();
}
int QCalendarWidget_firstDayOfWeek(void* self) {
	auto&& __r = ((QCalendarWidget*)self)->firstDayOfWeek();
	return (int)(__r);
}
void QCalendarWidget_setFirstDayOfWeek(void* self, int dayOfWeek) {
	((QCalendarWidget*)self)->setFirstDayOfWeek((Qt::DayOfWeek)(dayOfWeek));
}
int QCalendarWidget_isNavigationBarVisible(void* self) {
	auto&& __r = ((QCalendarWidget*)self)->isNavigationBarVisible();
	return __r;
}
int QCalendarWidget_isGridVisible(void* self) {
	auto&& __r = ((QCalendarWidget*)self)->isGridVisible();
	return __r;
}
void* QCalendarWidget_calendar(void* self) {
	auto&& __r = ((QCalendarWidget*)self)->calendar();
	return new QCalendar(__r);
}
void QCalendarWidget_setCalendar(void* self, void* calendar) {
	((QCalendarWidget*)self)->setCalendar(*(QCalendar*)calendar);
}
int QCalendarWidget_selectionMode(void* self) {
	auto&& __r = ((QCalendarWidget*)self)->selectionMode();
	return (int)(__r);
}
void QCalendarWidget_setSelectionMode(void* self, int mode) {
	((QCalendarWidget*)self)->setSelectionMode((QCalendarWidget::SelectionMode)(mode));
}
int QCalendarWidget_horizontalHeaderFormat(void* self) {
	auto&& __r = ((QCalendarWidget*)self)->horizontalHeaderFormat();
	return (int)(__r);
}
void QCalendarWidget_setHorizontalHeaderFormat(void* self, int format) {
	((QCalendarWidget*)self)->setHorizontalHeaderFormat((QCalendarWidget::HorizontalHeaderFormat)(format));
}
int QCalendarWidget_verticalHeaderFormat(void* self) {
	auto&& __r = ((QCalendarWidget*)self)->verticalHeaderFormat();
	return (int)(__r);
}
void QCalendarWidget_setVerticalHeaderFormat(void* self, int format) {
	((QCalendarWidget*)self)->setVerticalHeaderFormat((QCalendarWidget::VerticalHeaderFormat)(format));
}
void* QCalendarWidget_headerTextFormat(void* self) {
	auto&& __r = ((QCalendarWidget*)self)->headerTextFormat();
	return new QTextCharFormat(__r);
}
void QCalendarWidget_setHeaderTextFormat(void* self, void* format) {
	((QCalendarWidget*)self)->setHeaderTextFormat(*(QTextCharFormat*)format);
}
void* QCalendarWidget_weekdayTextFormat(void* self, int dayOfWeek) {
	auto&& __r = ((QCalendarWidget*)self)->weekdayTextFormat((Qt::DayOfWeek)(dayOfWeek));
	return new QTextCharFormat(__r);
}
void QCalendarWidget_setWeekdayTextFormat(void* self, int dayOfWeek, void* format) {
	((QCalendarWidget*)self)->setWeekdayTextFormat((Qt::DayOfWeek)(dayOfWeek), *(QTextCharFormat*)format);
}
void* QCalendarWidget_dateTextFormat(void* self, void* date) {
	auto&& __r = ((QCalendarWidget*)self)->dateTextFormat(*(QDate*)date);
	return new QTextCharFormat(__r);
}
void QCalendarWidget_setDateTextFormat(void* self, void* date, void* format) {
	((QCalendarWidget*)self)->setDateTextFormat(*(QDate*)date, *(QTextCharFormat*)format);
}
int QCalendarWidget_isDateEditEnabled(void* self) {
	auto&& __r = ((QCalendarWidget*)self)->isDateEditEnabled();
	return __r;
}
void QCalendarWidget_setDateEditEnabled(void* self, int enable) {
	((QCalendarWidget*)self)->setDateEditEnabled(enable);
}
int QCalendarWidget_dateEditAcceptDelay(void* self) {
	auto&& __r = ((QCalendarWidget*)self)->dateEditAcceptDelay();
	return __r;
}
void QCalendarWidget_setDateEditAcceptDelay(void* self, int delay) {
	((QCalendarWidget*)self)->setDateEditAcceptDelay(delay);
}
void QCalendarWidget_updateCell(void* self, void* date) {
	((QCalendarWidget__Go*)self)->updateCell(*(QDate*)date);
}
void QCalendarWidget_updateCells(void* self) {
	((QCalendarWidget__Go*)self)->updateCells();
}
void QCalendarWidget_setSelectedDate(void* self, void* date) {
	((QCalendarWidget*)self)->setSelectedDate(*(QDate*)date);
}
void QCalendarWidget_setDateRange(void* self, void* min, void* max) {
	((QCalendarWidget*)self)->setDateRange(*(QDate*)min, *(QDate*)max);
}
void QCalendarWidget_setCurrentPage(void* self, int year, int month) {
	((QCalendarWidget*)self)->setCurrentPage(year, month);
}
void QCalendarWidget_setGridVisible(void* self, int show) {
	((QCalendarWidget*)self)->setGridVisible(show);
}
void QCalendarWidget_setNavigationBarVisible(void* self, int visible) {
	((QCalendarWidget*)self)->setNavigationBarVisible(visible);
}
void QCalendarWidget_showNextMonth(void* self) {
	((QCalendarWidget*)self)->showNextMonth();
}
void QCalendarWidget_showPreviousMonth(void* self) {
	((QCalendarWidget*)self)->showPreviousMonth();
}
void QCalendarWidget_showNextYear(void* self) {
	((QCalendarWidget*)self)->showNextYear();
}
void QCalendarWidget_showPreviousYear(void* self) {
	((QCalendarWidget*)self)->showPreviousYear();
}
void QCalendarWidget_showSelectedDate(void* self) {
	((QCalendarWidget*)self)->showSelectedDate();
}
void QCalendarWidget_showToday(void* self) {
	((QCalendarWidget*)self)->showToday();
}
void QCalendarWidget_selectionChanged(void* self) {
	((QCalendarWidget*)self)->selectionChanged();
}
void QCalendarWidget_clicked(void* self, void* date) {
	((QCalendarWidget*)self)->clicked(*(QDate*)date);
}
void QCalendarWidget_activated(void* self, void* date) {
	((QCalendarWidget*)self)->activated(*(QDate*)date);
}
void QCalendarWidget_currentPageChanged(void* self, int year, int month) {
	((QCalendarWidget*)self)->currentPageChanged(year, month);
}
void QCalendarWidget_setTabOrder(void* arg1, void* arg2) {
	QCalendarWidget::setTabOrder((QWidget*)arg1, (QWidget*)arg2);
}
void* QCalendarWidget_mouseGrabber() {
	auto&& __r = QCalendarWidget::mouseGrabber();
	return (void*)(__r);
}
void* QCalendarWidget_keyboardGrabber() {
	auto&& __r = QCalendarWidget::keyboardGrabber();
	return (void*)(__r);
}
void* QCalendarWidget_find(unsigned long long arg1) {
	auto&& __r = QCalendarWidget::find(arg1);
	return (void*)(__r);
}
void* QCalendarWidget_createWindowContainer(void* window) {
	auto&& __r = QCalendarWidget::createWindowContainer((QWindow*)window);
	return (void*)(__r);
}
void QCalendarWidget_updateMicroFocus(void* self) {
	((QCalendarWidget__Go*)self)->updateMicroFocus();
}
void QCalendarWidget_create(void* self) {
	((QCalendarWidget__Go*)self)->create();
}
void QCalendarWidget_destroy(void* self) {
	((QCalendarWidget__Go*)self)->destroy();
}
int QCalendarWidget_focusNextChild(void* self) {
	auto&& __r = ((QCalendarWidget__Go*)self)->focusNextChild();
	return __r;
}
int QCalendarWidget_focusPreviousChild(void* self) {
	auto&& __r = ((QCalendarWidget__Go*)self)->focusPreviousChild();
	return __r;
}
void* QCalendarWidget_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QCalendarWidget::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QCalendarWidget_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QCalendarWidget::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QCalendarWidget_sender(void* self) {
	auto&& __r = ((QCalendarWidget__Go*)self)->sender();
	return (void*)(__r);
}
int QCalendarWidget_senderSignalIndex(void* self) {
	auto&& __r = ((QCalendarWidget__Go*)self)->senderSignalIndex();
	return __r;
}
int QCalendarWidget_receivers(void* self, const char* signal) {
	auto&& __r = ((QCalendarWidget__Go*)self)->receivers(signal);
	return __r;
}
int QCalendarWidget_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QCalendarWidget__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
double QCalendarWidget_devicePixelRatioFScale() {
	auto&& __r = QCalendarWidget::devicePixelRatioFScale();
	return __r;
}
int QCalendarWidget_encodeMetricF(int metric, double value) {
	auto&& __r = QCalendarWidget::encodeMetricF((QPaintDevice::PaintDeviceMetric)(metric), value);
	return __r;
}
double QCalendarWidget_getDecodedMetricF(void* self, int metricA, int metricB) {
	auto&& __r = ((QCalendarWidget__Go*)self)->getDecodedMetricF((QPaintDevice::PaintDeviceMetric)(metricA), (QPaintDevice::PaintDeviceMetric)(metricB));
	return __r;
}
char* QCalendarWidget_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QCalendarWidget::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QCalendarWidget_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QCalendarWidget::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QCalendarWidget_createWindowContainer2(void* window, void* parent) {
	auto&& __r = QCalendarWidget::createWindowContainer((QWindow*)window, (QWidget*)parent);
	return (void*)(__r);
}
void* QCalendarWidget_createWindowContainer3(void* window, void* parent, int flags) {
	auto&& __r = QCalendarWidget::createWindowContainer((QWindow*)window, (QWidget*)parent, QFlags<Qt::WindowType>(flags));
	return (void*)(__r);
}
void QCalendarWidget_updateMicroFocus2(void* self, int query) {
	((QCalendarWidget__Go*)self)->updateMicroFocus((Qt::InputMethodQuery)(query));
}
void QCalendarWidget_create2(void* self, unsigned long long arg1) {
	((QCalendarWidget__Go*)self)->create(arg1);
}
void QCalendarWidget_create3(void* self, unsigned long long arg1, int initializeWindow) {
	((QCalendarWidget__Go*)self)->create(arg1, initializeWindow);
}
void QCalendarWidget_create4(void* self, unsigned long long arg1, int initializeWindow, int destroyOldWindow) {
	((QCalendarWidget__Go*)self)->create(arg1, initializeWindow, destroyOldWindow);
}
void QCalendarWidget_destroy2(void* self, int destroyWindow) {
	((QCalendarWidget__Go*)self)->destroy(destroyWindow);
}
void QCalendarWidget_destroy3(void* self, int destroyWindow, int destroySubWindows) {
	((QCalendarWidget__Go*)self)->destroy(destroyWindow, destroySubWindows);
}
void* QCalendarWidget_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QCalendarWidget::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QCalendarWidget_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QCalendarWidget::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QCalendarWidget_metaObject_default(void* self) {
	auto&& __r = ((QCalendarWidget*)self)->QCalendarWidget::metaObject();
	return (void*)(__r);
}
void* QCalendarWidget_sizeHint_default(void* self) {
	auto&& __r = ((QCalendarWidget*)self)->QCalendarWidget::sizeHint();
	return new QSize(__r);
}
void* QCalendarWidget_minimumSizeHint_default(void* self) {
	auto&& __r = ((QCalendarWidget*)self)->QCalendarWidget::minimumSizeHint();
	return new QSize(__r);
}
int QCalendarWidget_event_default(void* self, void* event) {
	auto&& __r = ((QCalendarWidget__Go*)self)->cppgoBase_event((QEvent*)event);
	return __r;
}
int QCalendarWidget_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QCalendarWidget__Go*)self)->cppgoBase_eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QCalendarWidget_mousePressEvent_default(void* self, void* event) {
	((QCalendarWidget__Go*)self)->cppgoBase_mousePressEvent((QMouseEvent*)event);
}
void QCalendarWidget_resizeEvent_default(void* self, void* event) {
	((QCalendarWidget__Go*)self)->cppgoBase_resizeEvent((QResizeEvent*)event);
}
void QCalendarWidget_keyPressEvent_default(void* self, void* event) {
	((QCalendarWidget__Go*)self)->cppgoBase_keyPressEvent((QKeyEvent*)event);
}
void QCalendarWidget_paintCell_default(void* self, void* painter, void* rect, void* date) {
	((QCalendarWidget__Go*)self)->cppgoBase_paintCell((QPainter*)painter, *(QRect*)rect, *(QDate*)date);
}
int QCalendarWidget_devType_default(void* self) {
	auto&& __r = ((QCalendarWidget*)self)->QCalendarWidget::devType();
	return __r;
}
void QCalendarWidget_setVisible_default(void* self, int visible) {
	((QCalendarWidget*)self)->QCalendarWidget::setVisible(visible);
}
int QCalendarWidget_heightForWidth_default(void* self, int arg1) {
	auto&& __r = ((QCalendarWidget*)self)->QCalendarWidget::heightForWidth(arg1);
	return __r;
}
int QCalendarWidget_hasHeightForWidth_default(void* self) {
	auto&& __r = ((QCalendarWidget*)self)->QCalendarWidget::hasHeightForWidth();
	return __r;
}
void* QCalendarWidget_paintEngine_default(void* self) {
	auto&& __r = ((QCalendarWidget*)self)->QCalendarWidget::paintEngine();
	return (void*)(__r);
}
void QCalendarWidget_mouseReleaseEvent_default(void* self, void* event) {
	((QCalendarWidget__Go*)self)->cppgoBase_mouseReleaseEvent((QMouseEvent*)event);
}
void QCalendarWidget_mouseDoubleClickEvent_default(void* self, void* event) {
	((QCalendarWidget__Go*)self)->cppgoBase_mouseDoubleClickEvent((QMouseEvent*)event);
}
void QCalendarWidget_mouseMoveEvent_default(void* self, void* event) {
	((QCalendarWidget__Go*)self)->cppgoBase_mouseMoveEvent((QMouseEvent*)event);
}
void QCalendarWidget_wheelEvent_default(void* self, void* event) {
	((QCalendarWidget__Go*)self)->cppgoBase_wheelEvent((QWheelEvent*)event);
}
void QCalendarWidget_keyReleaseEvent_default(void* self, void* event) {
	((QCalendarWidget__Go*)self)->cppgoBase_keyReleaseEvent((QKeyEvent*)event);
}
void QCalendarWidget_focusInEvent_default(void* self, void* event) {
	((QCalendarWidget__Go*)self)->cppgoBase_focusInEvent((QFocusEvent*)event);
}
void QCalendarWidget_focusOutEvent_default(void* self, void* event) {
	((QCalendarWidget__Go*)self)->cppgoBase_focusOutEvent((QFocusEvent*)event);
}
void QCalendarWidget_enterEvent_default(void* self, void* event) {
	((QCalendarWidget__Go*)self)->cppgoBase_enterEvent((QEnterEvent*)event);
}
void QCalendarWidget_leaveEvent_default(void* self, void* event) {
	((QCalendarWidget__Go*)self)->cppgoBase_leaveEvent((QEvent*)event);
}
void QCalendarWidget_paintEvent_default(void* self, void* event) {
	((QCalendarWidget__Go*)self)->cppgoBase_paintEvent((QPaintEvent*)event);
}
void QCalendarWidget_moveEvent_default(void* self, void* event) {
	((QCalendarWidget__Go*)self)->cppgoBase_moveEvent((QMoveEvent*)event);
}
void QCalendarWidget_closeEvent_default(void* self, void* event) {
	((QCalendarWidget__Go*)self)->cppgoBase_closeEvent((QCloseEvent*)event);
}
void QCalendarWidget_contextMenuEvent_default(void* self, void* event) {
	((QCalendarWidget__Go*)self)->cppgoBase_contextMenuEvent((QContextMenuEvent*)event);
}
void QCalendarWidget_tabletEvent_default(void* self, void* event) {
	((QCalendarWidget__Go*)self)->cppgoBase_tabletEvent((QTabletEvent*)event);
}
void QCalendarWidget_actionEvent_default(void* self, void* event) {
	((QCalendarWidget__Go*)self)->cppgoBase_actionEvent((QActionEvent*)event);
}
void QCalendarWidget_dragEnterEvent_default(void* self, void* event) {
	((QCalendarWidget__Go*)self)->cppgoBase_dragEnterEvent((QDragEnterEvent*)event);
}
void QCalendarWidget_dragMoveEvent_default(void* self, void* event) {
	((QCalendarWidget__Go*)self)->cppgoBase_dragMoveEvent((QDragMoveEvent*)event);
}
void QCalendarWidget_dragLeaveEvent_default(void* self, void* event) {
	((QCalendarWidget__Go*)self)->cppgoBase_dragLeaveEvent((QDragLeaveEvent*)event);
}
void QCalendarWidget_dropEvent_default(void* self, void* event) {
	((QCalendarWidget__Go*)self)->cppgoBase_dropEvent((QDropEvent*)event);
}
void QCalendarWidget_showEvent_default(void* self, void* event) {
	((QCalendarWidget__Go*)self)->cppgoBase_showEvent((QShowEvent*)event);
}
void QCalendarWidget_hideEvent_default(void* self, void* event) {
	((QCalendarWidget__Go*)self)->cppgoBase_hideEvent((QHideEvent*)event);
}
void QCalendarWidget_changeEvent_default(void* self, void* arg1) {
	((QCalendarWidget__Go*)self)->cppgoBase_changeEvent((QEvent*)arg1);
}
int QCalendarWidget_metric_default(void* self, int arg1) {
	auto&& __r = ((QCalendarWidget__Go*)self)->cppgoBase_metric(arg1);
	return __r;
}
void QCalendarWidget_initPainter_default(void* self, void* painter) {
	((QCalendarWidget__Go*)self)->cppgoBase_initPainter((QPainter*)painter);
}
void* QCalendarWidget_redirected_default(void* self, void* offset) {
	auto&& __r = ((QCalendarWidget__Go*)self)->cppgoBase_redirected((QPoint*)offset);
	return (void*)(__r);
}
void* QCalendarWidget_sharedPainter_default(void* self) {
	auto&& __r = ((QCalendarWidget__Go*)self)->cppgoBase_sharedPainter();
	return (void*)(__r);
}
void QCalendarWidget_inputMethodEvent_default(void* self, void* arg1) {
	((QCalendarWidget__Go*)self)->cppgoBase_inputMethodEvent((QInputMethodEvent*)arg1);
}
void* QCalendarWidget_inputMethodQuery_default(void* self, int arg1) {
	auto&& __r = ((QCalendarWidget*)self)->QCalendarWidget::inputMethodQuery((Qt::InputMethodQuery)(arg1));
	return new QVariant(__r);
}
int QCalendarWidget_focusNextPrevChild_default(void* self, int next) {
	auto&& __r = ((QCalendarWidget__Go*)self)->cppgoBase_focusNextPrevChild(next);
	return __r;
}
void QCalendarWidget_timerEvent_default(void* self, void* event) {
	((QCalendarWidget__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QCalendarWidget_childEvent_default(void* self, void* event) {
	((QCalendarWidget__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QCalendarWidget_customEvent_default(void* self, void* event) {
	((QCalendarWidget__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QCalendarWidget_connectNotify_default(void* self, void* signal) {
	((QCalendarWidget__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QCalendarWidget_disconnectNotify_default(void* self, void* signal) {
	((QCalendarWidget__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QCalendarWidget_connect_selectionChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QCalendarWidget*)self, &QCalendarWidget::selectionChanged, (QCalendarWidget*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QCalendarWidget);
	__conns_QCalendarWidget[__id] = __c;
}
void QCalendarWidget_connect_clicked(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QCalendarWidget*)self, &QCalendarWidget::clicked, (QCalendarWidget*)self, [__id](QDate date){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&date);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QCalendarWidget);
	__conns_QCalendarWidget[__id] = __c;
}
void QCalendarWidget_connect_activated(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QCalendarWidget*)self, &QCalendarWidget::activated, (QCalendarWidget*)self, [__id](QDate date){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&date);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QCalendarWidget);
	__conns_QCalendarWidget[__id] = __c;
}
void QCalendarWidget_connect_currentPageChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QCalendarWidget*)self, &QCalendarWidget::currentPageChanged, (QCalendarWidget*)self, [__id](int year, int month){
		if (cppgo_sig_void_int_int) ((void(*)(long long, int, int))cppgo_sig_void_int_int)(__id, year, month);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QCalendarWidget);
	__conns_QCalendarWidget[__id] = __c;
}
void QCalendarWidget_connect_windowTitleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QCalendarWidget*)self, &QCalendarWidget::windowTitleChanged, (QCalendarWidget*)self, [__id](const QString& title){
		QByteArray title_ba = title.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)title_ba.constData(), title_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QCalendarWidget);
	__conns_QCalendarWidget[__id] = __c;
}
void QCalendarWidget_connect_windowIconChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QCalendarWidget*)self, &QCalendarWidget::windowIconChanged, (QCalendarWidget*)self, [__id](const QIcon& icon){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&icon);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QCalendarWidget);
	__conns_QCalendarWidget[__id] = __c;
}
void QCalendarWidget_connect_windowIconTextChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QCalendarWidget*)self, &QCalendarWidget::windowIconTextChanged, (QCalendarWidget*)self, [__id](const QString& iconText){
		QByteArray iconText_ba = iconText.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)iconText_ba.constData(), iconText_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QCalendarWidget);
	__conns_QCalendarWidget[__id] = __c;
}
void QCalendarWidget_connect_customContextMenuRequested(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QCalendarWidget*)self, &QCalendarWidget::customContextMenuRequested, (QCalendarWidget*)self, [__id](const QPoint& pos){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&pos);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QCalendarWidget);
	__conns_QCalendarWidget[__id] = __c;
}
void QCalendarWidget_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QCalendarWidget*)self, &QCalendarWidget::destroyed, (QCalendarWidget*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QCalendarWidget);
	__conns_QCalendarWidget[__id] = __c;
}
void QCalendarWidget_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QCalendarWidget*)self, &QCalendarWidget::destroyed, (QCalendarWidget*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QCalendarWidget);
	__conns_QCalendarWidget[__id] = __c;
}
void QCalendarWidget_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QCalendarWidget*)self, &QCalendarWidget::objectNameChanged, (QCalendarWidget*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QCalendarWidget);
	__conns_QCalendarWidget[__id] = __c;
}
void QCalendarWidget__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QCalendarWidget);
		auto __it = __conns_QCalendarWidget.find(id);
		if (__it == __conns_QCalendarWidget.end()) return;
		__c = __it->second;
		__conns_QCalendarWidget.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QCalendarWidget__Go_new(long long id) { return new QCalendarWidget__Go(id); }
void QCalendarWidget_delete(void* self) { delete (QCalendarWidget*)self; }

}

#include <QtWidgets/qfontdialog.h>
#include <QtWidgets/qdialog.h>
#include <QtWidgets/qwidget.h>
#include <QtGui/qfont.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qobject.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qsize.h>
#include <QtGui/qevent.h>
#include <QtWidgets/qstyle.h>
#include <QtCore/qnamespace.h>
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
#include "gen_qfontdialog.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QFontDialog__Go : public QFontDialog {
public:
	long long __goId;
	unsigned char __goOn[52] = {};
	template<typename... A> QFontDialog__Go(long long id, A&&... a) : QFontDialog(std::forward<A>(a)...), __goId(id) {}
	using QFontDialog::adjustPosition;
	using QFontDialog::updateMicroFocus;
	using QFontDialog::create;
	using QFontDialog::destroy;
	using QFontDialog::focusNextChild;
	using QFontDialog::focusPreviousChild;
	using QFontDialog::sender;
	using QFontDialog::senderSignalIndex;
	using QFontDialog::receivers;
	using QFontDialog::isSignalConnected;
	using QFontDialog::getDecodedMetricF;
	void cppgoBase_changeEvent(QEvent* __a0) { QFontDialog::changeEvent(__a0); }
	void cppgoBase_done(int __a0) { QFontDialog::done(__a0); }
	bool cppgoBase_eventFilter(QObject* __a0, QEvent* __a1) { return QFontDialog::eventFilter(__a0, __a1); }
	void cppgoBase_keyPressEvent(QKeyEvent* __a0) { QFontDialog::keyPressEvent(__a0); }
	void cppgoBase_closeEvent(QCloseEvent* __a0) { QFontDialog::closeEvent(__a0); }
	void cppgoBase_showEvent(QShowEvent* __a0) { QFontDialog::showEvent(__a0); }
	void cppgoBase_resizeEvent(QResizeEvent* __a0) { QFontDialog::resizeEvent(__a0); }
	void cppgoBase_contextMenuEvent(QContextMenuEvent* __a0) { QFontDialog::contextMenuEvent(__a0); }
	bool cppgoBase_event(QEvent* __a0) { return QFontDialog::event(__a0); }
	void cppgoBase_mousePressEvent(QMouseEvent* __a0) { QFontDialog::mousePressEvent(__a0); }
	void cppgoBase_mouseReleaseEvent(QMouseEvent* __a0) { QFontDialog::mouseReleaseEvent(__a0); }
	void cppgoBase_mouseDoubleClickEvent(QMouseEvent* __a0) { QFontDialog::mouseDoubleClickEvent(__a0); }
	void cppgoBase_mouseMoveEvent(QMouseEvent* __a0) { QFontDialog::mouseMoveEvent(__a0); }
	void cppgoBase_wheelEvent(QWheelEvent* __a0) { QFontDialog::wheelEvent(__a0); }
	void cppgoBase_keyReleaseEvent(QKeyEvent* __a0) { QFontDialog::keyReleaseEvent(__a0); }
	void cppgoBase_focusInEvent(QFocusEvent* __a0) { QFontDialog::focusInEvent(__a0); }
	void cppgoBase_focusOutEvent(QFocusEvent* __a0) { QFontDialog::focusOutEvent(__a0); }
	void cppgoBase_enterEvent(QEnterEvent* __a0) { QFontDialog::enterEvent(__a0); }
	void cppgoBase_leaveEvent(QEvent* __a0) { QFontDialog::leaveEvent(__a0); }
	void cppgoBase_paintEvent(QPaintEvent* __a0) { QFontDialog::paintEvent(__a0); }
	void cppgoBase_moveEvent(QMoveEvent* __a0) { QFontDialog::moveEvent(__a0); }
	void cppgoBase_tabletEvent(QTabletEvent* __a0) { QFontDialog::tabletEvent(__a0); }
	void cppgoBase_actionEvent(QActionEvent* __a0) { QFontDialog::actionEvent(__a0); }
	void cppgoBase_dragEnterEvent(QDragEnterEvent* __a0) { QFontDialog::dragEnterEvent(__a0); }
	void cppgoBase_dragMoveEvent(QDragMoveEvent* __a0) { QFontDialog::dragMoveEvent(__a0); }
	void cppgoBase_dragLeaveEvent(QDragLeaveEvent* __a0) { QFontDialog::dragLeaveEvent(__a0); }
	void cppgoBase_dropEvent(QDropEvent* __a0) { QFontDialog::dropEvent(__a0); }
	void cppgoBase_hideEvent(QHideEvent* __a0) { QFontDialog::hideEvent(__a0); }
	int cppgoBase_metric(int __a0) { return QFontDialog::metric((QPaintDevice::PaintDeviceMetric)__a0); }
	void cppgoBase_initPainter(QPainter* __a0) { QFontDialog::initPainter(__a0); }
	QPaintDevice* cppgoBase_redirected(QPoint* __a0) { return QFontDialog::redirected(__a0); }
	QPainter* cppgoBase_sharedPainter() { return QFontDialog::sharedPainter(); }
	void cppgoBase_inputMethodEvent(QInputMethodEvent* __a0) { QFontDialog::inputMethodEvent(__a0); }
	bool cppgoBase_focusNextPrevChild(bool __a0) { return QFontDialog::focusNextPrevChild(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QFontDialog::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QFontDialog::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QFontDialog::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QFontDialog::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QFontDialog::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QFontDialog::metaObject();
	}
	void setVisible(bool visible) override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 2110182719375245575LL, visible, &__h); if (__h) return; }
		QFontDialog::setVisible(visible);
	}
	void changeEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4959376063826165943LL, (void*)event, &__h); if (__h) return; }
		QFontDialog::changeEvent(event);
	}
	void done(int result) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, -2570997490807640431LL, result, &__h); if (__h) return; }
		QFontDialog::done(result);
	}
	bool eventFilter(QObject* object, QEvent* event) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)object, (void*)event, &__h); if (__h) return __r; }
		return QFontDialog::eventFilter(object, event);
	}
	QSize sizeHint() const override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 729690261457684101LL, &__h); if (__h) return *(QSize*)__r; }
		return QFontDialog::sizeHint();
	}
	QSize minimumSizeHint() const override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 7049264599837532491LL, &__h); if (__h) return *(QSize*)__r; }
		return QFontDialog::minimumSizeHint();
	}
	int exec() override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -1745322561779819684LL, &__h); if (__h) return __r; }
		return QFontDialog::exec();
	}
	void accept() override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 9106131928185084681LL, &__h); if (__h) return; }
		QFontDialog::accept();
	}
	void reject() override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 2875220963396170150LL, &__h); if (__h) return; }
		QFontDialog::reject();
	}
	void keyPressEvent(QKeyEvent* arg1) override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1023791977679793185LL, (void*)arg1, &__h); if (__h) return; }
		QFontDialog::keyPressEvent(arg1);
	}
	void closeEvent(QCloseEvent* arg1) override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1989937050019464783LL, (void*)arg1, &__h); if (__h) return; }
		QFontDialog::closeEvent(arg1);
	}
	void showEvent(QShowEvent* arg1) override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6825369174751916LL, (void*)arg1, &__h); if (__h) return; }
		QFontDialog::showEvent(arg1);
	}
	void resizeEvent(QResizeEvent* arg1) override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4302917285655014661LL, (void*)arg1, &__h); if (__h) return; }
		QFontDialog::resizeEvent(arg1);
	}
	void contextMenuEvent(QContextMenuEvent* arg1) override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6698043222227006015LL, (void*)arg1, &__h); if (__h) return; }
		QFontDialog::contextMenuEvent(arg1);
	}
	int devType() const override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8659679893868524044LL, &__h); if (__h) return __r; }
		return QFontDialog::devType();
	}
	int heightForWidth(int arg1) const override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 4311534060495737339LL, arg1, &__h); if (__h) return __r; }
		return QFontDialog::heightForWidth(arg1);
	}
	bool hasHeightForWidth() const override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8266762397995423915LL, &__h); if (__h) return __r; }
		return QFontDialog::hasHeightForWidth();
	}
	QPaintEngine* paintEngine() const override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 6398007404756346381LL, &__h); if (__h) return (QPaintEngine*)__r; }
		return QFontDialog::paintEngine();
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QFontDialog::event(event);
	}
	void mousePressEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -332090898078543319LL, (void*)event, &__h); if (__h) return; }
		QFontDialog::mousePressEvent(event);
	}
	void mouseReleaseEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4415920049297795213LL, (void*)event, &__h); if (__h) return; }
		QFontDialog::mouseReleaseEvent(event);
	}
	void mouseDoubleClickEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4029073905697023887LL, (void*)event, &__h); if (__h) return; }
		QFontDialog::mouseDoubleClickEvent(event);
	}
	void mouseMoveEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4255494004099806581LL, (void*)event, &__h); if (__h) return; }
		QFontDialog::mouseMoveEvent(event);
	}
	void wheelEvent(QWheelEvent* event) override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1919988340934511316LL, (void*)event, &__h); if (__h) return; }
		QFontDialog::wheelEvent(event);
	}
	void keyReleaseEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 772598475950050299LL, (void*)event, &__h); if (__h) return; }
		QFontDialog::keyReleaseEvent(event);
	}
	void focusInEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 975043767033552552LL, (void*)event, &__h); if (__h) return; }
		QFontDialog::focusInEvent(event);
	}
	void focusOutEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -385596010576048993LL, (void*)event, &__h); if (__h) return; }
		QFontDialog::focusOutEvent(event);
	}
	void enterEvent(QEnterEvent* event) override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6316270346668237895LL, (void*)event, &__h); if (__h) return; }
		QFontDialog::enterEvent(event);
	}
	void leaveEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -551070009623711640LL, (void*)event, &__h); if (__h) return; }
		QFontDialog::leaveEvent(event);
	}
	void paintEvent(QPaintEvent* event) override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5290442163167458465LL, (void*)event, &__h); if (__h) return; }
		QFontDialog::paintEvent(event);
	}
	void moveEvent(QMoveEvent* event) override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -9161949150489525396LL, (void*)event, &__h); if (__h) return; }
		QFontDialog::moveEvent(event);
	}
	void tabletEvent(QTabletEvent* event) override {
		int __h = 0;
		if (__goOn[32] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -212068866860701141LL, (void*)event, &__h); if (__h) return; }
		QFontDialog::tabletEvent(event);
	}
	void actionEvent(QActionEvent* event) override {
		int __h = 0;
		if (__goOn[33] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1697444140088555581LL, (void*)event, &__h); if (__h) return; }
		QFontDialog::actionEvent(event);
	}
	void dragEnterEvent(QDragEnterEvent* event) override {
		int __h = 0;
		if (__goOn[34] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 206122715919917589LL, (void*)event, &__h); if (__h) return; }
		QFontDialog::dragEnterEvent(event);
	}
	void dragMoveEvent(QDragMoveEvent* event) override {
		int __h = 0;
		if (__goOn[35] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1781545137352745978LL, (void*)event, &__h); if (__h) return; }
		QFontDialog::dragMoveEvent(event);
	}
	void dragLeaveEvent(QDragLeaveEvent* event) override {
		int __h = 0;
		if (__goOn[36] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -3139904158373725006LL, (void*)event, &__h); if (__h) return; }
		QFontDialog::dragLeaveEvent(event);
	}
	void dropEvent(QDropEvent* event) override {
		int __h = 0;
		if (__goOn[37] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 9107058682754591716LL, (void*)event, &__h); if (__h) return; }
		QFontDialog::dropEvent(event);
	}
	void hideEvent(QHideEvent* event) override {
		int __h = 0;
		if (__goOn[38] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -199718707876719129LL, (void*)event, &__h); if (__h) return; }
		QFontDialog::hideEvent(event);
	}
	bool nativeEvent(const QByteArray& eventType, void* message, qintptr* result) override {
		int __h = 0;
		if (__goOn[39] && cppgo_virt_int_voidp_voidp_longlongp) { int __r = ((int(*)(long long, long long, void*, void*, long long*, int*))cppgo_virt_int_voidp_voidp_longlongp)(__goId, -2039166586246443706LL, (void*)&eventType, (void*)message, (long long*)result, &__h); if (__h) return __r; }
		return QFontDialog::nativeEvent(eventType, message, result);
	}
	int metric(QPaintDevice::PaintDeviceMetric arg1) const override {
		int __h = 0;
		if (__goOn[40] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 427629649267311777LL, (int)(arg1), &__h); if (__h) return __r; }
		return QFontDialog::metric(arg1);
	}
	void initPainter(QPainter* painter) const override {
		int __h = 0;
		if (__goOn[41] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6477500296980050210LL, (void*)painter, &__h); if (__h) return; }
		QFontDialog::initPainter(painter);
	}
	QPaintDevice* redirected(QPoint* offset) const override {
		int __h = 0;
		if (__goOn[42] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 4127240633310655048LL, (void*)offset, &__h); if (__h) return (QPaintDevice*)__r; }
		return QFontDialog::redirected(offset);
	}
	QPainter* sharedPainter() const override {
		int __h = 0;
		if (__goOn[43] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 5792771328741774991LL, &__h); if (__h) return (QPainter*)__r; }
		return QFontDialog::sharedPainter();
	}
	void inputMethodEvent(QInputMethodEvent* arg1) override {
		int __h = 0;
		if (__goOn[44] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5438819661942843222LL, (void*)arg1, &__h); if (__h) return; }
		QFontDialog::inputMethodEvent(arg1);
	}
	QVariant inputMethodQuery(Qt::InputMethodQuery arg1) const override {
		int __h = 0;
		if (__goOn[45] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7345853097340709638LL, (int)(arg1), &__h); if (__h) return *(QVariant*)__r; }
		return QFontDialog::inputMethodQuery(arg1);
	}
	bool focusNextPrevChild(bool next) override {
		int __h = 0;
		if (__goOn[46] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 5898328517505784191LL, next, &__h); if (__h) return __r; }
		return QFontDialog::focusNextPrevChild(next);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[47] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QFontDialog::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[48] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QFontDialog::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[49] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QFontDialog::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[50] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QFontDialog::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[51] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QFontDialog::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QFontDialog;
static std::mutex __connsMu_QFontDialog;

extern "C" {

void* QFontDialog_base(void* self) { return (void*)static_cast<QDialog*>((QFontDialog*)self); }
void QFontDialog_chain(void* self, void** out) {
	QFontDialog* __p = (QFontDialog*)self;
	QDialog* __b0 = static_cast<QDialog*>(__p); out[0] = (void*)__b0;
	QWidget* __b1 = static_cast<QWidget*>(__b0); out[1] = (void*)__b1;
	QObject* __b2 = static_cast<QObject*>(__b1); out[2] = (void*)__b2;
}
void QFontDialog__set_override(void* self, int idx) { ((QFontDialog__Go*)self)->__goOn[idx] = 1; }
void* QFontDialog_fromQObject(void* o) { return (void*)qobject_cast<QFontDialog*>((QObject*)o); }
void* QFontDialog_new(long long __goid, void* parent) {
	return new QFontDialog__Go(__goid, (QWidget*)parent);
}
void* QFontDialog_new2(long long __goid) {
	return new QFontDialog__Go(__goid);
}
void* QFontDialog_new3(long long __goid, void* initial) {
	return new QFontDialog__Go(__goid, *(QFont*)initial);
}
void* QFontDialog_new4(long long __goid, void* initial, void* parent) {
	return new QFontDialog__Go(__goid, *(QFont*)initial, (QWidget*)parent);
}
char* QFontDialog_tr(const char* s, long long* __out_len) {
	auto&& __r = QFontDialog::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QFontDialog_setCurrentFont(void* self, void* font) {
	((QFontDialog*)self)->setCurrentFont(*(QFont*)font);
}
void* QFontDialog_currentFont(void* self) {
	auto&& __r = ((QFontDialog*)self)->currentFont();
	return new QFont(__r);
}
void* QFontDialog_selectedFont(void* self) {
	auto&& __r = ((QFontDialog*)self)->selectedFont();
	return new QFont(__r);
}
void QFontDialog_setOption(void* self, int option) {
	((QFontDialog*)self)->setOption((QFontDialog::FontDialogOption)(option));
}
int QFontDialog_testOption(void* self, int option) {
	auto&& __r = ((QFontDialog*)self)->testOption((QFontDialog::FontDialogOption)(option));
	return __r;
}
void QFontDialog_setOptions(void* self, int options) {
	((QFontDialog*)self)->setOptions(QFlags<QFontDialog::FontDialogOption>(options));
}
int QFontDialog_options(void* self) {
	auto&& __r = ((QFontDialog*)self)->options();
	return (int)(__r);
}
void QFontDialog_open(void* self, void* receiver, const char* member) {
	((QFontDialog*)self)->open((QObject*)receiver, member);
}
void* QFontDialog_getFont(int* ok) {
	bool ok_local{};
	auto&& __r = QFontDialog::getFont(&ok_local);
	*ok = ok_local;
	return new QFont(__r);
}
void* QFontDialog_getFont2(int* ok, void* initial) {
	bool ok_local{};
	auto&& __r = QFontDialog::getFont(&ok_local, *(QFont*)initial);
	*ok = ok_local;
	return new QFont(__r);
}
void QFontDialog_currentFontChanged(void* self, void* font) {
	((QFontDialog*)self)->currentFontChanged(*(QFont*)font);
}
void QFontDialog_fontSelected(void* self, void* font) {
	((QFontDialog*)self)->fontSelected(*(QFont*)font);
}
void QFontDialog_adjustPosition(void* self, void* arg1) {
	((QFontDialog__Go*)self)->adjustPosition((QWidget*)arg1);
}
void QFontDialog_setTabOrder(void* arg1, void* arg2) {
	QFontDialog::setTabOrder((QWidget*)arg1, (QWidget*)arg2);
}
void* QFontDialog_mouseGrabber() {
	auto&& __r = QFontDialog::mouseGrabber();
	return (void*)(__r);
}
void* QFontDialog_keyboardGrabber() {
	auto&& __r = QFontDialog::keyboardGrabber();
	return (void*)(__r);
}
void* QFontDialog_find(unsigned long long arg1) {
	auto&& __r = QFontDialog::find(arg1);
	return (void*)(__r);
}
void* QFontDialog_createWindowContainer(void* window) {
	auto&& __r = QFontDialog::createWindowContainer((QWindow*)window);
	return (void*)(__r);
}
void QFontDialog_updateMicroFocus(void* self) {
	((QFontDialog__Go*)self)->updateMicroFocus();
}
void QFontDialog_create(void* self) {
	((QFontDialog__Go*)self)->create();
}
void QFontDialog_destroy(void* self) {
	((QFontDialog__Go*)self)->destroy();
}
int QFontDialog_focusNextChild(void* self) {
	auto&& __r = ((QFontDialog__Go*)self)->focusNextChild();
	return __r;
}
int QFontDialog_focusPreviousChild(void* self) {
	auto&& __r = ((QFontDialog__Go*)self)->focusPreviousChild();
	return __r;
}
void* QFontDialog_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QFontDialog::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QFontDialog_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QFontDialog::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QFontDialog_sender(void* self) {
	auto&& __r = ((QFontDialog__Go*)self)->sender();
	return (void*)(__r);
}
int QFontDialog_senderSignalIndex(void* self) {
	auto&& __r = ((QFontDialog__Go*)self)->senderSignalIndex();
	return __r;
}
int QFontDialog_receivers(void* self, const char* signal) {
	auto&& __r = ((QFontDialog__Go*)self)->receivers(signal);
	return __r;
}
int QFontDialog_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QFontDialog__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
double QFontDialog_devicePixelRatioFScale() {
	auto&& __r = QFontDialog::devicePixelRatioFScale();
	return __r;
}
int QFontDialog_encodeMetricF(int metric, double value) {
	auto&& __r = QFontDialog::encodeMetricF((QPaintDevice::PaintDeviceMetric)(metric), value);
	return __r;
}
double QFontDialog_getDecodedMetricF(void* self, int metricA, int metricB) {
	auto&& __r = ((QFontDialog__Go*)self)->getDecodedMetricF((QPaintDevice::PaintDeviceMetric)(metricA), (QPaintDevice::PaintDeviceMetric)(metricB));
	return __r;
}
char* QFontDialog_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QFontDialog::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QFontDialog_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QFontDialog::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QFontDialog_setOption2(void* self, int option, int on) {
	((QFontDialog*)self)->setOption((QFontDialog::FontDialogOption)(option), on);
}
void* QFontDialog_getFont3(int* ok, void* parent) {
	bool ok_local{};
	auto&& __r = QFontDialog::getFont(&ok_local, (QWidget*)parent);
	*ok = ok_local;
	return new QFont(__r);
}
void* QFontDialog_getFont4(int* ok, void* initial, void* parent) {
	bool ok_local{};
	auto&& __r = QFontDialog::getFont(&ok_local, *(QFont*)initial, (QWidget*)parent);
	*ok = ok_local;
	return new QFont(__r);
}
void* QFontDialog_getFont5(int* ok, void* initial, void* parent, const char* title, long long title_len) {
	bool ok_local{};
	QString title_qstr = QString::fromUtf8(title, title_len);
	auto&& __r = QFontDialog::getFont(&ok_local, *(QFont*)initial, (QWidget*)parent, title_qstr);
	*ok = ok_local;
	return new QFont(__r);
}
void* QFontDialog_getFont6(int* ok, void* initial, void* parent, const char* title, long long title_len, int options) {
	bool ok_local{};
	QString title_qstr = QString::fromUtf8(title, title_len);
	auto&& __r = QFontDialog::getFont(&ok_local, *(QFont*)initial, (QWidget*)parent, title_qstr, QFlags<QFontDialog::FontDialogOption>(options));
	*ok = ok_local;
	return new QFont(__r);
}
void* QFontDialog_createWindowContainer2(void* window, void* parent) {
	auto&& __r = QFontDialog::createWindowContainer((QWindow*)window, (QWidget*)parent);
	return (void*)(__r);
}
void* QFontDialog_createWindowContainer3(void* window, void* parent, int flags) {
	auto&& __r = QFontDialog::createWindowContainer((QWindow*)window, (QWidget*)parent, QFlags<Qt::WindowType>(flags));
	return (void*)(__r);
}
void QFontDialog_updateMicroFocus2(void* self, int query) {
	((QFontDialog__Go*)self)->updateMicroFocus((Qt::InputMethodQuery)(query));
}
void QFontDialog_create2(void* self, unsigned long long arg1) {
	((QFontDialog__Go*)self)->create(arg1);
}
void QFontDialog_create3(void* self, unsigned long long arg1, int initializeWindow) {
	((QFontDialog__Go*)self)->create(arg1, initializeWindow);
}
void QFontDialog_create4(void* self, unsigned long long arg1, int initializeWindow, int destroyOldWindow) {
	((QFontDialog__Go*)self)->create(arg1, initializeWindow, destroyOldWindow);
}
void QFontDialog_destroy2(void* self, int destroyWindow) {
	((QFontDialog__Go*)self)->destroy(destroyWindow);
}
void QFontDialog_destroy3(void* self, int destroyWindow, int destroySubWindows) {
	((QFontDialog__Go*)self)->destroy(destroyWindow, destroySubWindows);
}
void* QFontDialog_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QFontDialog::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QFontDialog_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QFontDialog::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QFontDialog_metaObject_default(void* self) {
	auto&& __r = ((QFontDialog*)self)->QFontDialog::metaObject();
	return (void*)(__r);
}
void QFontDialog_setVisible_default(void* self, int visible) {
	((QFontDialog*)self)->QFontDialog::setVisible(visible);
}
void QFontDialog_changeEvent_default(void* self, void* event) {
	((QFontDialog__Go*)self)->cppgoBase_changeEvent((QEvent*)event);
}
void QFontDialog_done_default(void* self, int result) {
	((QFontDialog__Go*)self)->cppgoBase_done(result);
}
int QFontDialog_eventFilter_default(void* self, void* object, void* event) {
	auto&& __r = ((QFontDialog__Go*)self)->cppgoBase_eventFilter((QObject*)object, (QEvent*)event);
	return __r;
}
void* QFontDialog_sizeHint_default(void* self) {
	auto&& __r = ((QFontDialog*)self)->QFontDialog::sizeHint();
	return new QSize(__r);
}
void* QFontDialog_minimumSizeHint_default(void* self) {
	auto&& __r = ((QFontDialog*)self)->QFontDialog::minimumSizeHint();
	return new QSize(__r);
}
int QFontDialog_exec_default(void* self) {
	auto&& __r = ((QFontDialog*)self)->QFontDialog::exec();
	return __r;
}
void QFontDialog_accept_default(void* self) {
	((QFontDialog*)self)->QFontDialog::accept();
}
void QFontDialog_reject_default(void* self) {
	((QFontDialog*)self)->QFontDialog::reject();
}
void QFontDialog_keyPressEvent_default(void* self, void* arg1) {
	((QFontDialog__Go*)self)->cppgoBase_keyPressEvent((QKeyEvent*)arg1);
}
void QFontDialog_closeEvent_default(void* self, void* arg1) {
	((QFontDialog__Go*)self)->cppgoBase_closeEvent((QCloseEvent*)arg1);
}
void QFontDialog_showEvent_default(void* self, void* arg1) {
	((QFontDialog__Go*)self)->cppgoBase_showEvent((QShowEvent*)arg1);
}
void QFontDialog_resizeEvent_default(void* self, void* arg1) {
	((QFontDialog__Go*)self)->cppgoBase_resizeEvent((QResizeEvent*)arg1);
}
void QFontDialog_contextMenuEvent_default(void* self, void* arg1) {
	((QFontDialog__Go*)self)->cppgoBase_contextMenuEvent((QContextMenuEvent*)arg1);
}
int QFontDialog_devType_default(void* self) {
	auto&& __r = ((QFontDialog*)self)->QFontDialog::devType();
	return __r;
}
int QFontDialog_heightForWidth_default(void* self, int arg1) {
	auto&& __r = ((QFontDialog*)self)->QFontDialog::heightForWidth(arg1);
	return __r;
}
int QFontDialog_hasHeightForWidth_default(void* self) {
	auto&& __r = ((QFontDialog*)self)->QFontDialog::hasHeightForWidth();
	return __r;
}
void* QFontDialog_paintEngine_default(void* self) {
	auto&& __r = ((QFontDialog*)self)->QFontDialog::paintEngine();
	return (void*)(__r);
}
int QFontDialog_event_default(void* self, void* event) {
	auto&& __r = ((QFontDialog__Go*)self)->cppgoBase_event((QEvent*)event);
	return __r;
}
void QFontDialog_mousePressEvent_default(void* self, void* event) {
	((QFontDialog__Go*)self)->cppgoBase_mousePressEvent((QMouseEvent*)event);
}
void QFontDialog_mouseReleaseEvent_default(void* self, void* event) {
	((QFontDialog__Go*)self)->cppgoBase_mouseReleaseEvent((QMouseEvent*)event);
}
void QFontDialog_mouseDoubleClickEvent_default(void* self, void* event) {
	((QFontDialog__Go*)self)->cppgoBase_mouseDoubleClickEvent((QMouseEvent*)event);
}
void QFontDialog_mouseMoveEvent_default(void* self, void* event) {
	((QFontDialog__Go*)self)->cppgoBase_mouseMoveEvent((QMouseEvent*)event);
}
void QFontDialog_wheelEvent_default(void* self, void* event) {
	((QFontDialog__Go*)self)->cppgoBase_wheelEvent((QWheelEvent*)event);
}
void QFontDialog_keyReleaseEvent_default(void* self, void* event) {
	((QFontDialog__Go*)self)->cppgoBase_keyReleaseEvent((QKeyEvent*)event);
}
void QFontDialog_focusInEvent_default(void* self, void* event) {
	((QFontDialog__Go*)self)->cppgoBase_focusInEvent((QFocusEvent*)event);
}
void QFontDialog_focusOutEvent_default(void* self, void* event) {
	((QFontDialog__Go*)self)->cppgoBase_focusOutEvent((QFocusEvent*)event);
}
void QFontDialog_enterEvent_default(void* self, void* event) {
	((QFontDialog__Go*)self)->cppgoBase_enterEvent((QEnterEvent*)event);
}
void QFontDialog_leaveEvent_default(void* self, void* event) {
	((QFontDialog__Go*)self)->cppgoBase_leaveEvent((QEvent*)event);
}
void QFontDialog_paintEvent_default(void* self, void* event) {
	((QFontDialog__Go*)self)->cppgoBase_paintEvent((QPaintEvent*)event);
}
void QFontDialog_moveEvent_default(void* self, void* event) {
	((QFontDialog__Go*)self)->cppgoBase_moveEvent((QMoveEvent*)event);
}
void QFontDialog_tabletEvent_default(void* self, void* event) {
	((QFontDialog__Go*)self)->cppgoBase_tabletEvent((QTabletEvent*)event);
}
void QFontDialog_actionEvent_default(void* self, void* event) {
	((QFontDialog__Go*)self)->cppgoBase_actionEvent((QActionEvent*)event);
}
void QFontDialog_dragEnterEvent_default(void* self, void* event) {
	((QFontDialog__Go*)self)->cppgoBase_dragEnterEvent((QDragEnterEvent*)event);
}
void QFontDialog_dragMoveEvent_default(void* self, void* event) {
	((QFontDialog__Go*)self)->cppgoBase_dragMoveEvent((QDragMoveEvent*)event);
}
void QFontDialog_dragLeaveEvent_default(void* self, void* event) {
	((QFontDialog__Go*)self)->cppgoBase_dragLeaveEvent((QDragLeaveEvent*)event);
}
void QFontDialog_dropEvent_default(void* self, void* event) {
	((QFontDialog__Go*)self)->cppgoBase_dropEvent((QDropEvent*)event);
}
void QFontDialog_hideEvent_default(void* self, void* event) {
	((QFontDialog__Go*)self)->cppgoBase_hideEvent((QHideEvent*)event);
}
int QFontDialog_metric_default(void* self, int arg1) {
	auto&& __r = ((QFontDialog__Go*)self)->cppgoBase_metric(arg1);
	return __r;
}
void QFontDialog_initPainter_default(void* self, void* painter) {
	((QFontDialog__Go*)self)->cppgoBase_initPainter((QPainter*)painter);
}
void* QFontDialog_redirected_default(void* self, void* offset) {
	auto&& __r = ((QFontDialog__Go*)self)->cppgoBase_redirected((QPoint*)offset);
	return (void*)(__r);
}
void* QFontDialog_sharedPainter_default(void* self) {
	auto&& __r = ((QFontDialog__Go*)self)->cppgoBase_sharedPainter();
	return (void*)(__r);
}
void QFontDialog_inputMethodEvent_default(void* self, void* arg1) {
	((QFontDialog__Go*)self)->cppgoBase_inputMethodEvent((QInputMethodEvent*)arg1);
}
void* QFontDialog_inputMethodQuery_default(void* self, int arg1) {
	auto&& __r = ((QFontDialog*)self)->QFontDialog::inputMethodQuery((Qt::InputMethodQuery)(arg1));
	return new QVariant(__r);
}
int QFontDialog_focusNextPrevChild_default(void* self, int next) {
	auto&& __r = ((QFontDialog__Go*)self)->cppgoBase_focusNextPrevChild(next);
	return __r;
}
void QFontDialog_timerEvent_default(void* self, void* event) {
	((QFontDialog__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QFontDialog_childEvent_default(void* self, void* event) {
	((QFontDialog__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QFontDialog_customEvent_default(void* self, void* event) {
	((QFontDialog__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QFontDialog_connectNotify_default(void* self, void* signal) {
	((QFontDialog__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QFontDialog_disconnectNotify_default(void* self, void* signal) {
	((QFontDialog__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QFontDialog_connect_currentFontChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFontDialog*)self, &QFontDialog::currentFontChanged, (QFontDialog*)self, [__id](const QFont& font){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&font);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFontDialog);
	__conns_QFontDialog[__id] = __c;
}
void QFontDialog_connect_fontSelected(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFontDialog*)self, &QFontDialog::fontSelected, (QFontDialog*)self, [__id](const QFont& font){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&font);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFontDialog);
	__conns_QFontDialog[__id] = __c;
}
void QFontDialog_connect_finished(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFontDialog*)self, &QFontDialog::finished, (QFontDialog*)self, [__id](int result){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, result);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFontDialog);
	__conns_QFontDialog[__id] = __c;
}
void QFontDialog_connect_accepted(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFontDialog*)self, &QFontDialog::accepted, (QFontDialog*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFontDialog);
	__conns_QFontDialog[__id] = __c;
}
void QFontDialog_connect_rejected(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFontDialog*)self, &QFontDialog::rejected, (QFontDialog*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFontDialog);
	__conns_QFontDialog[__id] = __c;
}
void QFontDialog_connect_windowTitleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFontDialog*)self, &QFontDialog::windowTitleChanged, (QFontDialog*)self, [__id](const QString& title){
		QByteArray title_ba = title.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)title_ba.constData(), title_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFontDialog);
	__conns_QFontDialog[__id] = __c;
}
void QFontDialog_connect_windowIconChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFontDialog*)self, &QFontDialog::windowIconChanged, (QFontDialog*)self, [__id](const QIcon& icon){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&icon);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFontDialog);
	__conns_QFontDialog[__id] = __c;
}
void QFontDialog_connect_windowIconTextChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFontDialog*)self, &QFontDialog::windowIconTextChanged, (QFontDialog*)self, [__id](const QString& iconText){
		QByteArray iconText_ba = iconText.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)iconText_ba.constData(), iconText_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFontDialog);
	__conns_QFontDialog[__id] = __c;
}
void QFontDialog_connect_customContextMenuRequested(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFontDialog*)self, &QFontDialog::customContextMenuRequested, (QFontDialog*)self, [__id](const QPoint& pos){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&pos);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFontDialog);
	__conns_QFontDialog[__id] = __c;
}
void QFontDialog_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFontDialog*)self, &QFontDialog::destroyed, (QFontDialog*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFontDialog);
	__conns_QFontDialog[__id] = __c;
}
void QFontDialog_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFontDialog*)self, &QFontDialog::destroyed, (QFontDialog*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFontDialog);
	__conns_QFontDialog[__id] = __c;
}
void QFontDialog_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFontDialog*)self, &QFontDialog::objectNameChanged, (QFontDialog*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFontDialog);
	__conns_QFontDialog[__id] = __c;
}
void QFontDialog__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QFontDialog);
		auto __it = __conns_QFontDialog.find(id);
		if (__it == __conns_QFontDialog.end()) return;
		__c = __it->second;
		__conns_QFontDialog.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QFontDialog__Go_new(long long id) { return new QFontDialog__Go(id); }
void QFontDialog_delete(void* self) { delete (QFontDialog*)self; }

}

#include <QtWidgets/qgraphicslayout.h>
#include <QtWidgets/qgraphicslayoutitem.h>
#include <QtCore/qcoreevent.h>
#include <QtWidgets/qsizepolicy.h>
#include <QtCore/qsize.h>
#include <QtCore/qrect.h>
#include <QtCore/qnamespace.h>
#include <QtWidgets/qgraphicsitem.h>
#include "gen_qgraphicslayout.h"
#include "gen_purego_cb.h"

class QGraphicsLayout__Go : public QGraphicsLayout {
public:
	long long __goId;
	unsigned char __goOn[10] = {};
	template<typename... A> QGraphicsLayout__Go(long long id, A&&... a) : QGraphicsLayout(std::forward<A>(a)...), __goId(id) {}
	using QGraphicsLayout::addChildLayoutItem;
	using QGraphicsLayout::setGraphicsItem;
	using QGraphicsLayout::setOwnedByLayout;
	QSizeF sizeHint(Qt::SizeHint which, const QSizeF& constraint) const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp_int_voidp) { void* __r = ((void*(*)(long long, long long, int, void*, int*))cppgo_virt_voidp_int_voidp)(__goId, 729690261457684101LL, (int)(which), (void*)&constraint, &__h); if (__h) return *(QSizeF*)__r; }
		return QSizeF{};
	}
	int count() const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -5627843058667511180LL, &__h); if (__h) return __r; }
		return int{};
	}
	QGraphicsLayoutItem* itemAt(int i) const override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7508512214013391107LL, i, &__h); if (__h) return (QGraphicsLayoutItem*)__r; }
		return nullptr;
	}
	void removeAt(int index) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 1230346564330554752LL, index, &__h); if (__h) return; }
		return;
	}
	void getContentsMargins(qreal* left, qreal* top, qreal* right, qreal* bottom) const override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_doublep_doublep_doublep_doublep) { ((void(*)(long long, long long, double*, double*, double*, double*, int*))cppgo_virt_void_doublep_doublep_doublep_doublep)(__goId, 164760286202646972LL, (double*)left, (double*)top, (double*)right, (double*)bottom, &__h); if (__h) return; }
		QGraphicsLayout::getContentsMargins(left, top, right, bottom);
	}
	void invalidate() override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 6991341317990632368LL, &__h); if (__h) return; }
		QGraphicsLayout::invalidate();
	}
	void updateGeometry() override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 8082455443689927668LL, &__h); if (__h) return; }
		QGraphicsLayout::updateGeometry();
	}
	void widgetEvent(QEvent* e) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 8535746785413991679LL, (void*)e, &__h); if (__h) return; }
		QGraphicsLayout::widgetEvent(e);
	}
	void setGeometry(const QRectF& rect) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5684820848726969227LL, (void*)&rect, &__h); if (__h) return; }
		QGraphicsLayout::setGeometry(rect);
	}
	bool isEmpty() const override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -626822430251098498LL, &__h); if (__h) return __r; }
		return QGraphicsLayout::isEmpty();
	}
};

extern "C" {

void* QGraphicsLayout_base(void* self) { return (void*)static_cast<QGraphicsLayoutItem*>((QGraphicsLayout*)self); }
void QGraphicsLayout__set_override(void* self, int idx) { ((QGraphicsLayout__Go*)self)->__goOn[idx] = 1; }
void QGraphicsLayout_setContentsMargins(void* self, double left, double top, double right, double bottom) {
	((QGraphicsLayout*)self)->setContentsMargins(left, top, right, bottom);
}
void QGraphicsLayout_activate(void* self) {
	((QGraphicsLayout*)self)->activate();
}
int QGraphicsLayout_isActivated(void* self) {
	auto&& __r = ((QGraphicsLayout*)self)->isActivated();
	return __r;
}
void QGraphicsLayout_invalidate(void* self) {
	((QGraphicsLayout*)self)->invalidate();
}
void QGraphicsLayout_widgetEvent(void* self, void* e) {
	((QGraphicsLayout*)self)->widgetEvent((QEvent*)e);
}
int QGraphicsLayout_count(void* self) {
	auto&& __r = ((QGraphicsLayout*)self)->count();
	return __r;
}
void* QGraphicsLayout_itemAt(void* self, int i) {
	auto&& __r = ((QGraphicsLayout*)self)->itemAt(i);
	return (void*)(__r);
}
void QGraphicsLayout_removeAt(void* self, int index) {
	((QGraphicsLayout*)self)->removeAt(index);
}
void QGraphicsLayout_setInstantInvalidatePropagation(int enable) {
	QGraphicsLayout::setInstantInvalidatePropagation(enable);
}
int QGraphicsLayout_instantInvalidatePropagation() {
	auto&& __r = QGraphicsLayout::instantInvalidatePropagation();
	return __r;
}
void QGraphicsLayout_addChildLayoutItem(void* self, void* layoutItem) {
	((QGraphicsLayout__Go*)self)->addChildLayoutItem((QGraphicsLayoutItem*)layoutItem);
}
void QGraphicsLayout_setGraphicsItem(void* self, void* item) {
	((QGraphicsLayout__Go*)self)->setGraphicsItem((QGraphicsItem*)item);
}
void QGraphicsLayout_setOwnedByLayout(void* self, int ownedByLayout) {
	((QGraphicsLayout__Go*)self)->setOwnedByLayout(ownedByLayout);
}
void QGraphicsLayout_invalidate_default(void* self) {
	((QGraphicsLayout*)self)->QGraphicsLayout::invalidate();
}
void QGraphicsLayout_widgetEvent_default(void* self, void* e) {
	((QGraphicsLayout*)self)->QGraphicsLayout::widgetEvent((QEvent*)e);
}
void QGraphicsLayout_getContentsMargins_default(void* self, double* left, double* top, double* right, double* bottom) {
	qreal left_local{};
	qreal top_local{};
	qreal right_local{};
	qreal bottom_local{};
	((QGraphicsLayout*)self)->QGraphicsLayout::getContentsMargins(&left_local, &top_local, &right_local, &bottom_local);
	*left = left_local;
	*top = top_local;
	*right = right_local;
	*bottom = bottom_local;
}
void QGraphicsLayout_updateGeometry_default(void* self) {
	((QGraphicsLayout*)self)->QGraphicsLayout::updateGeometry();
}
void QGraphicsLayout_setGeometry_default(void* self, void* rect) {
	((QGraphicsLayout*)self)->QGraphicsLayout::setGeometry(*(QRectF*)rect);
}
int QGraphicsLayout_isEmpty_default(void* self) {
	auto&& __r = ((QGraphicsLayout*)self)->QGraphicsLayout::isEmpty();
	return __r;
}
void* QGraphicsLayout__Go_new(long long id) { return new QGraphicsLayout__Go(id); }
void QGraphicsLayout_delete(void* self) { delete (QGraphicsLayout*)self; }

}

#include <QtWidgets/qgraphicstransform.h>
#include <QtCore/qobject.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtGui/qvectornd.h>
#include <QtGui/qmatrix4x4.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qvariant.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qgraphicsscale.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QGraphicsScale__Go : public QGraphicsScale {
public:
	long long __goId;
	unsigned char __goOn[9] = {};
	template<typename... A> QGraphicsScale__Go(long long id, A&&... a) : QGraphicsScale(std::forward<A>(a)...), __goId(id) {}
	using QGraphicsScale::update;
	using QGraphicsScale::sender;
	using QGraphicsScale::senderSignalIndex;
	using QGraphicsScale::receivers;
	using QGraphicsScale::isSignalConnected;
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QGraphicsScale::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QGraphicsScale::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QGraphicsScale::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QGraphicsScale::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QGraphicsScale::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QGraphicsScale::metaObject();
	}
	void applyTo(QMatrix4x4* matrix) const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1580758529722784438LL, (void*)matrix, &__h); if (__h) return; }
		QGraphicsScale::applyTo(matrix);
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QGraphicsScale::event(event);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QGraphicsScale::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QGraphicsScale::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QGraphicsScale::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QGraphicsScale::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QGraphicsScale::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QGraphicsScale::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QGraphicsScale;
static std::mutex __connsMu_QGraphicsScale;

extern "C" {

void* QGraphicsScale_base(void* self) { return (void*)static_cast<QGraphicsTransform*>((QGraphicsScale*)self); }
void QGraphicsScale_chain(void* self, void** out) {
	QGraphicsScale* __p = (QGraphicsScale*)self;
	QGraphicsTransform* __b0 = static_cast<QGraphicsTransform*>(__p); out[0] = (void*)__b0;
	QObject* __b1 = static_cast<QObject*>(__b0); out[1] = (void*)__b1;
}
void QGraphicsScale__set_override(void* self, int idx) { ((QGraphicsScale__Go*)self)->__goOn[idx] = 1; }
void* QGraphicsScale_fromQObject(void* o) { return (void*)qobject_cast<QGraphicsScale*>((QObject*)o); }
void* QGraphicsScale_new(long long __goid) {
	return new QGraphicsScale__Go(__goid);
}
void* QGraphicsScale_new2(long long __goid, void* parent) {
	return new QGraphicsScale__Go(__goid, (QObject*)parent);
}
char* QGraphicsScale_tr(const char* s, long long* __out_len) {
	auto&& __r = QGraphicsScale::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QGraphicsScale_origin(void* self) {
	auto&& __r = ((QGraphicsScale*)self)->origin();
	return new QVector3D(__r);
}
void QGraphicsScale_setOrigin(void* self, void* point) {
	((QGraphicsScale*)self)->setOrigin(*(QVector3D*)point);
}
double QGraphicsScale_xScale(void* self) {
	auto&& __r = ((QGraphicsScale*)self)->xScale();
	return __r;
}
void QGraphicsScale_setXScale(void* self, double arg1) {
	((QGraphicsScale*)self)->setXScale(arg1);
}
double QGraphicsScale_yScale(void* self) {
	auto&& __r = ((QGraphicsScale*)self)->yScale();
	return __r;
}
void QGraphicsScale_setYScale(void* self, double arg1) {
	((QGraphicsScale*)self)->setYScale(arg1);
}
double QGraphicsScale_zScale(void* self) {
	auto&& __r = ((QGraphicsScale*)self)->zScale();
	return __r;
}
void QGraphicsScale_setZScale(void* self, double arg1) {
	((QGraphicsScale*)self)->setZScale(arg1);
}
void QGraphicsScale_originChanged(void* self) {
	((QGraphicsScale*)self)->originChanged();
}
void QGraphicsScale_xScaleChanged(void* self) {
	((QGraphicsScale*)self)->xScaleChanged();
}
void QGraphicsScale_yScaleChanged(void* self) {
	((QGraphicsScale*)self)->yScaleChanged();
}
void QGraphicsScale_zScaleChanged(void* self) {
	((QGraphicsScale*)self)->zScaleChanged();
}
void QGraphicsScale_scaleChanged(void* self) {
	((QGraphicsScale*)self)->scaleChanged();
}
void QGraphicsScale_update(void* self) {
	((QGraphicsScale__Go*)self)->update();
}
void* QGraphicsScale_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QGraphicsScale::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QGraphicsScale_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QGraphicsScale::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QGraphicsScale_sender(void* self) {
	auto&& __r = ((QGraphicsScale__Go*)self)->sender();
	return (void*)(__r);
}
int QGraphicsScale_senderSignalIndex(void* self) {
	auto&& __r = ((QGraphicsScale__Go*)self)->senderSignalIndex();
	return __r;
}
int QGraphicsScale_receivers(void* self, const char* signal) {
	auto&& __r = ((QGraphicsScale__Go*)self)->receivers(signal);
	return __r;
}
int QGraphicsScale_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QGraphicsScale__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QGraphicsScale_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QGraphicsScale::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QGraphicsScale_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QGraphicsScale::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QGraphicsScale_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QGraphicsScale::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QGraphicsScale_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QGraphicsScale::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QGraphicsScale_metaObject_default(void* self) {
	auto&& __r = ((QGraphicsScale*)self)->QGraphicsScale::metaObject();
	return (void*)(__r);
}
void QGraphicsScale_applyTo_default(void* self, void* matrix) {
	((QGraphicsScale*)self)->QGraphicsScale::applyTo((QMatrix4x4*)matrix);
}
int QGraphicsScale_event_default(void* self, void* event) {
	auto&& __r = ((QGraphicsScale*)self)->QGraphicsScale::event((QEvent*)event);
	return __r;
}
int QGraphicsScale_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QGraphicsScale*)self)->QGraphicsScale::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QGraphicsScale_timerEvent_default(void* self, void* event) {
	((QGraphicsScale__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QGraphicsScale_childEvent_default(void* self, void* event) {
	((QGraphicsScale__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QGraphicsScale_customEvent_default(void* self, void* event) {
	((QGraphicsScale__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QGraphicsScale_connectNotify_default(void* self, void* signal) {
	((QGraphicsScale__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QGraphicsScale_disconnectNotify_default(void* self, void* signal) {
	((QGraphicsScale__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QGraphicsScale_connect_originChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsScale*)self, &QGraphicsScale::originChanged, (QGraphicsScale*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsScale);
	__conns_QGraphicsScale[__id] = __c;
}
void QGraphicsScale_connect_xScaleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsScale*)self, &QGraphicsScale::xScaleChanged, (QGraphicsScale*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsScale);
	__conns_QGraphicsScale[__id] = __c;
}
void QGraphicsScale_connect_yScaleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsScale*)self, &QGraphicsScale::yScaleChanged, (QGraphicsScale*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsScale);
	__conns_QGraphicsScale[__id] = __c;
}
void QGraphicsScale_connect_zScaleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsScale*)self, &QGraphicsScale::zScaleChanged, (QGraphicsScale*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsScale);
	__conns_QGraphicsScale[__id] = __c;
}
void QGraphicsScale_connect_scaleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsScale*)self, &QGraphicsScale::scaleChanged, (QGraphicsScale*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsScale);
	__conns_QGraphicsScale[__id] = __c;
}
void QGraphicsScale_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsScale*)self, &QGraphicsScale::destroyed, (QGraphicsScale*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsScale);
	__conns_QGraphicsScale[__id] = __c;
}
void QGraphicsScale_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsScale*)self, &QGraphicsScale::destroyed, (QGraphicsScale*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsScale);
	__conns_QGraphicsScale[__id] = __c;
}
void QGraphicsScale_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsScale*)self, &QGraphicsScale::objectNameChanged, (QGraphicsScale*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsScale);
	__conns_QGraphicsScale[__id] = __c;
}
void QGraphicsScale__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsScale);
		auto __it = __conns_QGraphicsScale.find(id);
		if (__it == __conns_QGraphicsScale.end()) return;
		__c = __it->second;
		__conns_QGraphicsScale.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QGraphicsScale__Go_new(long long id, void* parent) { return new QGraphicsScale__Go(id, (QObject*)parent); }
void QGraphicsScale_delete(void* self) { delete (QGraphicsScale*)self; }

}

#include <QtWidgets/qtextedit.h>
#include <QtGui/qtextcursor.h>
#include <QtGui/qtextformat.h>
#include "gen_qtextedit__extraselection.h"

extern "C" {

void* QTextEdit__ExtraSelection_new(void* arg1) {
	return new QTextEdit::ExtraSelection(*(QTextEdit::ExtraSelection*)arg1);
}
void* QTextEdit__ExtraSelection_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QTextEdit::ExtraSelection*)self)->operator=(*(QTextEdit::ExtraSelection*)arg1);
	return (void*)&(__r);
}
void* QTextEdit__ExtraSelection_cursor_get(void* self) {
	auto __r = ((QTextEdit::ExtraSelection*)self)->cursor;
	return new QTextCursor(__r);
}
void QTextEdit__ExtraSelection_cursor_set(void* self, void* cursor) {
	((QTextEdit::ExtraSelection*)self)->cursor = *(QTextCursor*)cursor;
}
void* QTextEdit__ExtraSelection_format_get(void* self) {
	auto __r = ((QTextEdit::ExtraSelection*)self)->format;
	return new QTextCharFormat(__r);
}
void QTextEdit__ExtraSelection_format_set(void* self, void* format) {
	((QTextEdit::ExtraSelection*)self)->format = *(QTextCharFormat*)format;
}
void QTextEdit__ExtraSelection_delete(void* self) { delete (QTextEdit::ExtraSelection*)self; }

}

#include <QtWidgets/qsystemtrayicon.h>
#include <QtCore/qobject.h>
#include <QtGui/qicon.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtWidgets/qmenu.h>
#include <QtCore/qrect.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qvariant.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qsystemtrayicon.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QSystemTrayIcon__Go : public QSystemTrayIcon {
public:
	long long __goId;
	unsigned char __goOn[8] = {};
	template<typename... A> QSystemTrayIcon__Go(long long id, A&&... a) : QSystemTrayIcon(std::forward<A>(a)...), __goId(id) {}
	using QSystemTrayIcon::sender;
	using QSystemTrayIcon::senderSignalIndex;
	using QSystemTrayIcon::receivers;
	using QSystemTrayIcon::isSignalConnected;
	bool cppgoBase_event(QEvent* __a0) { return QSystemTrayIcon::event(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QSystemTrayIcon::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QSystemTrayIcon::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QSystemTrayIcon::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QSystemTrayIcon::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QSystemTrayIcon::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QSystemTrayIcon::metaObject();
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QSystemTrayIcon::event(event);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QSystemTrayIcon::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QSystemTrayIcon::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QSystemTrayIcon::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QSystemTrayIcon::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QSystemTrayIcon::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QSystemTrayIcon::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QSystemTrayIcon;
static std::mutex __connsMu_QSystemTrayIcon;

extern "C" {

void* QSystemTrayIcon_base(void* self) { return (void*)static_cast<QObject*>((QSystemTrayIcon*)self); }
void QSystemTrayIcon__set_override(void* self, int idx) { ((QSystemTrayIcon__Go*)self)->__goOn[idx] = 1; }
void* QSystemTrayIcon_fromQObject(void* o) { return (void*)qobject_cast<QSystemTrayIcon*>((QObject*)o); }
void* QSystemTrayIcon_new(long long __goid) {
	return new QSystemTrayIcon__Go(__goid);
}
void* QSystemTrayIcon_new2(long long __goid, void* icon) {
	return new QSystemTrayIcon__Go(__goid, *(QIcon*)icon);
}
void* QSystemTrayIcon_new3(long long __goid, void* parent) {
	return new QSystemTrayIcon__Go(__goid, (QObject*)parent);
}
void* QSystemTrayIcon_new4(long long __goid, void* icon, void* parent) {
	return new QSystemTrayIcon__Go(__goid, *(QIcon*)icon, (QObject*)parent);
}
char* QSystemTrayIcon_tr(const char* s, long long* __out_len) {
	auto&& __r = QSystemTrayIcon::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QSystemTrayIcon_setContextMenu(void* self, void* menu) {
	((QSystemTrayIcon*)self)->setContextMenu((QMenu*)menu);
}
void* QSystemTrayIcon_contextMenu(void* self) {
	auto&& __r = ((QSystemTrayIcon*)self)->contextMenu();
	return (void*)(__r);
}
void* QSystemTrayIcon_icon(void* self) {
	auto&& __r = ((QSystemTrayIcon*)self)->icon();
	return new QIcon(__r);
}
void QSystemTrayIcon_setIcon(void* self, void* icon) {
	((QSystemTrayIcon*)self)->setIcon(*(QIcon*)icon);
}
char* QSystemTrayIcon_toolTip(void* self, long long* __out_len) {
	auto&& __r = ((QSystemTrayIcon*)self)->toolTip();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QSystemTrayIcon_setToolTip(void* self, const char* tip, long long tip_len) {
	QString tip_qstr = QString::fromUtf8(tip, tip_len);
	((QSystemTrayIcon*)self)->setToolTip(tip_qstr);
}
int QSystemTrayIcon_isSystemTrayAvailable() {
	auto&& __r = QSystemTrayIcon::isSystemTrayAvailable();
	return __r;
}
int QSystemTrayIcon_supportsMessages() {
	auto&& __r = QSystemTrayIcon::supportsMessages();
	return __r;
}
void* QSystemTrayIcon_geometry(void* self) {
	auto&& __r = ((QSystemTrayIcon*)self)->geometry();
	return new QRect(__r);
}
int QSystemTrayIcon_isVisible(void* self) {
	auto&& __r = ((QSystemTrayIcon*)self)->isVisible();
	return __r;
}
void QSystemTrayIcon_setVisible(void* self, int visible) {
	((QSystemTrayIcon*)self)->setVisible(visible);
}
void QSystemTrayIcon_show(void* self) {
	((QSystemTrayIcon*)self)->show();
}
void QSystemTrayIcon_hide(void* self) {
	((QSystemTrayIcon*)self)->hide();
}
void QSystemTrayIcon_showMessage(void* self, const char* title, long long title_len, const char* msg, long long msg_len, void* icon) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	QString msg_qstr = QString::fromUtf8(msg, msg_len);
	((QSystemTrayIcon*)self)->showMessage(title_qstr, msg_qstr, *(QIcon*)icon);
}
void QSystemTrayIcon_showMessage2(void* self, const char* title, long long title_len, const char* msg, long long msg_len) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	QString msg_qstr = QString::fromUtf8(msg, msg_len);
	((QSystemTrayIcon*)self)->showMessage(title_qstr, msg_qstr);
}
void QSystemTrayIcon_activated(void* self, int reason) {
	((QSystemTrayIcon*)self)->activated((QSystemTrayIcon::ActivationReason)(reason));
}
void QSystemTrayIcon_messageClicked(void* self) {
	((QSystemTrayIcon*)self)->messageClicked();
}
void* QSystemTrayIcon_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QSystemTrayIcon::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QSystemTrayIcon_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QSystemTrayIcon::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QSystemTrayIcon_sender(void* self) {
	auto&& __r = ((QSystemTrayIcon__Go*)self)->sender();
	return (void*)(__r);
}
int QSystemTrayIcon_senderSignalIndex(void* self) {
	auto&& __r = ((QSystemTrayIcon__Go*)self)->senderSignalIndex();
	return __r;
}
int QSystemTrayIcon_receivers(void* self, const char* signal) {
	auto&& __r = ((QSystemTrayIcon__Go*)self)->receivers(signal);
	return __r;
}
int QSystemTrayIcon_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QSystemTrayIcon__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QSystemTrayIcon_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QSystemTrayIcon::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QSystemTrayIcon_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QSystemTrayIcon::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QSystemTrayIcon_showMessage3(void* self, const char* title, long long title_len, const char* msg, long long msg_len, void* icon, int msecs) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	QString msg_qstr = QString::fromUtf8(msg, msg_len);
	((QSystemTrayIcon*)self)->showMessage(title_qstr, msg_qstr, *(QIcon*)icon, msecs);
}
void QSystemTrayIcon_showMessage4(void* self, const char* title, long long title_len, const char* msg, long long msg_len, int icon) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	QString msg_qstr = QString::fromUtf8(msg, msg_len);
	((QSystemTrayIcon*)self)->showMessage(title_qstr, msg_qstr, (QSystemTrayIcon::MessageIcon)(icon));
}
void QSystemTrayIcon_showMessage5(void* self, const char* title, long long title_len, const char* msg, long long msg_len, int icon, int msecs) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	QString msg_qstr = QString::fromUtf8(msg, msg_len);
	((QSystemTrayIcon*)self)->showMessage(title_qstr, msg_qstr, (QSystemTrayIcon::MessageIcon)(icon), msecs);
}
void* QSystemTrayIcon_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QSystemTrayIcon::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QSystemTrayIcon_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QSystemTrayIcon::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QSystemTrayIcon_metaObject_default(void* self) {
	auto&& __r = ((QSystemTrayIcon*)self)->QSystemTrayIcon::metaObject();
	return (void*)(__r);
}
int QSystemTrayIcon_event_default(void* self, void* event) {
	auto&& __r = ((QSystemTrayIcon__Go*)self)->cppgoBase_event((QEvent*)event);
	return __r;
}
int QSystemTrayIcon_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QSystemTrayIcon*)self)->QSystemTrayIcon::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QSystemTrayIcon_timerEvent_default(void* self, void* event) {
	((QSystemTrayIcon__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QSystemTrayIcon_childEvent_default(void* self, void* event) {
	((QSystemTrayIcon__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QSystemTrayIcon_customEvent_default(void* self, void* event) {
	((QSystemTrayIcon__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QSystemTrayIcon_connectNotify_default(void* self, void* signal) {
	((QSystemTrayIcon__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QSystemTrayIcon_disconnectNotify_default(void* self, void* signal) {
	((QSystemTrayIcon__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QSystemTrayIcon_connect_activated(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSystemTrayIcon*)self, &QSystemTrayIcon::activated, (QSystemTrayIcon*)self, [__id](QSystemTrayIcon::ActivationReason reason){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, (int)(reason));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSystemTrayIcon);
	__conns_QSystemTrayIcon[__id] = __c;
}
void QSystemTrayIcon_connect_messageClicked(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSystemTrayIcon*)self, &QSystemTrayIcon::messageClicked, (QSystemTrayIcon*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSystemTrayIcon);
	__conns_QSystemTrayIcon[__id] = __c;
}
void QSystemTrayIcon_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSystemTrayIcon*)self, &QSystemTrayIcon::destroyed, (QSystemTrayIcon*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSystemTrayIcon);
	__conns_QSystemTrayIcon[__id] = __c;
}
void QSystemTrayIcon_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSystemTrayIcon*)self, &QSystemTrayIcon::destroyed, (QSystemTrayIcon*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSystemTrayIcon);
	__conns_QSystemTrayIcon[__id] = __c;
}
void QSystemTrayIcon_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSystemTrayIcon*)self, &QSystemTrayIcon::objectNameChanged, (QSystemTrayIcon*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSystemTrayIcon);
	__conns_QSystemTrayIcon[__id] = __c;
}
void QSystemTrayIcon__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QSystemTrayIcon);
		auto __it = __conns_QSystemTrayIcon.find(id);
		if (__it == __conns_QSystemTrayIcon.end()) return;
		__c = __it->second;
		__conns_QSystemTrayIcon.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QSystemTrayIcon__Go_new(long long id, void* parent) { return new QSystemTrayIcon__Go(id, (QObject*)parent); }
void QSystemTrayIcon_delete(void* self) { delete (QSystemTrayIcon*)self; }

}

