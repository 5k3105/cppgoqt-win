#include <QtCore/qobjectdefs.h>
#include "gen_qmetamethodargument.h"

extern "C" {

const char* QMetaMethodArgument_name_get(void* self) {
	auto __r = ((QMetaMethodArgument*)self)->name;
	return __r;
}
void QMetaMethodArgument_name_set(void* self, const char* name) {
	((QMetaMethodArgument*)self)->name = name;
}
void QMetaMethodArgument_delete(void* self) { delete (QMetaMethodArgument*)self; }

}

#include <QtCore/qiodevicebase.h>
#include "gen_qiodevicebase.h"

extern "C" {

void* QIODeviceBase_new(void* arg1) {
	return new QIODeviceBase(*(QIODeviceBase*)arg1);
}
void* QIODeviceBase_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QIODeviceBase*)self)->operator=(*(QIODeviceBase*)arg1);
	return (void*)&(__r);
}

}

#include <QtCore/qcontiguouscache.h>
#include "gen_qcontiguouscachedata.h"

extern "C" {

void* QContiguousCacheData_allocateData(long long size, long long alignment) {
	auto&& __r = QContiguousCacheData::allocateData(size, alignment);
	return (void*)(__r);
}
void QContiguousCacheData_freeData(void* data) {
	QContiguousCacheData::freeData((QContiguousCacheData*)data);
}
long long QContiguousCacheData_alloc_get(void* self) {
	auto __r = ((QContiguousCacheData*)self)->alloc;
	return __r;
}
void QContiguousCacheData_alloc_set(void* self, long long alloc) {
	((QContiguousCacheData*)self)->alloc = alloc;
}
long long QContiguousCacheData_count_get(void* self) {
	auto __r = ((QContiguousCacheData*)self)->count;
	return __r;
}
void QContiguousCacheData_count_set(void* self, long long count) {
	((QContiguousCacheData*)self)->count = count;
}
long long QContiguousCacheData_start_get(void* self) {
	auto __r = ((QContiguousCacheData*)self)->start;
	return __r;
}
void QContiguousCacheData_start_set(void* self, long long start) {
	((QContiguousCacheData*)self)->start = start;
}
long long QContiguousCacheData_offset_get(void* self) {
	auto __r = ((QContiguousCacheData*)self)->offset;
	return __r;
}
void QContiguousCacheData_offset_set(void* self, long long offset) {
	((QContiguousCacheData*)self)->offset = offset;
}
void QContiguousCacheData_delete(void* self) { delete (QContiguousCacheData*)self; }

}

#include <QtCore/qcoreevent.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qbasictimer.h>
#include "gen_qtimerevent.h"

extern "C" {

void* QTimerEvent_base(void* self) { return (void*)static_cast<QEvent*>((QTimerEvent*)self); }
void* QTimerEvent_new(int timerId) {
	return new QTimerEvent(timerId);
}
void* QTimerEvent_new2(int timerId) {
	return new QTimerEvent((Qt::TimerId)(timerId));
}
int QTimerEvent_timerId(void* self) {
	auto&& __r = ((QTimerEvent*)self)->timerId();
	return __r;
}
int QTimerEvent_id(void* self) {
	auto&& __r = ((QTimerEvent*)self)->id();
	return (int)(__r);
}
int QTimerEvent_matches(void* self, void* timer) {
	auto&& __r = ((QTimerEvent*)self)->matches(*(QBasicTimer*)timer);
	return __r;
}
int QTimerEvent_registerEventType() {
	auto&& __r = QTimerEvent::registerEventType();
	return __r;
}
int QTimerEvent_registerEventType2(int hint) {
	auto&& __r = QTimerEvent::registerEventType(hint);
	return __r;
}
void QTimerEvent_delete(void* self) { delete (QTimerEvent*)self; }

}

#include <QtCore/quuid.h>
#include "gen__guid.h"

extern "C" {

void* _GUID_new() {
	return new _GUID();
}
void* _GUID_new2(void* arg1) {
	return new _GUID(*(_GUID*)arg1);
}
unsigned long _GUID_Data1_get(void* self) {
	auto __r = ((_GUID*)self)->Data1;
	return __r;
}
void _GUID_Data1_set(void* self, unsigned long Data1) {
	((_GUID*)self)->Data1 = Data1;
}
unsigned short _GUID_Data2_get(void* self) {
	auto __r = ((_GUID*)self)->Data2;
	return __r;
}
void _GUID_Data2_set(void* self, unsigned short Data2) {
	((_GUID*)self)->Data2 = Data2;
}
unsigned short _GUID_Data3_get(void* self) {
	auto __r = ((_GUID*)self)->Data3;
	return __r;
}
void _GUID_Data3_set(void* self, unsigned short Data3) {
	((_GUID*)self)->Data3 = Data3;
}
void _GUID_delete(void* self) { delete (_GUID*)self; }

}

#include <QtCore/qcollator.h>
#include <QtCore/qlocale.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qstring.h>
#include <QtCore/qchar.h>
#include <QtCore/qstringview.h>
#include "gen_qcollator.h"

extern "C" {

void* QCollator_new() {
	return new QCollator();
}
void* QCollator_new2(void* locale) {
	return new QCollator(*(QLocale*)locale);
}
void* QCollator_new3(void* arg1) {
	return new QCollator(*(QCollator*)arg1);
}
void* QCollator_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QCollator*)self)->operator=(*(QCollator*)arg1);
	return (void*)&(__r);
}
void QCollator_swap(void* self, void* other) {
	((QCollator*)self)->swap(*(QCollator*)other);
}
void QCollator_setLocale(void* self, void* locale) {
	((QCollator*)self)->setLocale(*(QLocale*)locale);
}
void* QCollator_locale(void* self) {
	auto&& __r = ((QCollator*)self)->locale();
	return new QLocale(__r);
}
int QCollator_caseSensitivity(void* self) {
	auto&& __r = ((QCollator*)self)->caseSensitivity();
	return (int)(__r);
}
void QCollator_setCaseSensitivity(void* self, int cs) {
	((QCollator*)self)->setCaseSensitivity((Qt::CaseSensitivity)(cs));
}
void QCollator_setNumericMode(void* self, int on) {
	((QCollator*)self)->setNumericMode(on);
}
int QCollator_numericMode(void* self) {
	auto&& __r = ((QCollator*)self)->numericMode();
	return __r;
}
void QCollator_setIgnorePunctuation(void* self, int on) {
	((QCollator*)self)->setIgnorePunctuation(on);
}
int QCollator_ignorePunctuation(void* self) {
	auto&& __r = ((QCollator*)self)->ignorePunctuation();
	return __r;
}
int QCollator_compare(void* self, const char* s1, long long s1_len, const char* s2, long long s2_len) {
	QString s1_qstr = QString::fromUtf8(s1, s1_len);
	QString s2_qstr = QString::fromUtf8(s2, s2_len);
	auto&& __r = ((QCollator*)self)->compare(s1_qstr, s2_qstr);
	return __r;
}
int QCollator_compare2(void* self, void* s1, long long len1, void* s2, long long len2) {
	auto&& __r = ((QCollator*)self)->compare((QChar*)s1, len1, (QChar*)s2, len2);
	return __r;
}
int QCollator_OperatorCall(void* self, const char* s1, long long s1_len, const char* s2, long long s2_len) {
	QString s1_qstr = QString::fromUtf8(s1, s1_len);
	QString s2_qstr = QString::fromUtf8(s2, s2_len);
	auto&& __r = ((QCollator*)self)->operator()(s1_qstr, s2_qstr);
	return __r;
}
int QCollator_compare3(void* self, void* s1, void* s2) {
	auto&& __r = ((QCollator*)self)->compare(*(QStringView*)s1, *(QStringView*)s2);
	return __r;
}
int QCollator_OperatorCall2(void* self, void* s1, void* s2) {
	auto&& __r = ((QCollator*)self)->operator()(*(QStringView*)s1, *(QStringView*)s2);
	return __r;
}
void* QCollator_sortKey(void* self, const char* string_, long long string__len) {
	QString string__qstr = QString::fromUtf8(string_, string__len);
	auto&& __r = ((QCollator*)self)->sortKey(string__qstr);
	return new QCollatorSortKey(__r);
}
int QCollator_defaultCompare(void* s1, void* s2) {
	auto&& __r = QCollator::defaultCompare(*(QStringView*)s1, *(QStringView*)s2);
	return __r;
}
void* QCollator_defaultSortKey(void* key) {
	auto&& __r = QCollator::defaultSortKey(*(QStringView*)key);
	return new QCollatorSortKey(__r);
}
void QCollator_delete(void* self) { delete (QCollator*)self; }

}

#include <QtCore/qthread.h>
#include <QtCore/qobject.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qabstracteventdispatcher.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qdeadlinetimer.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qvariant.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qthread.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QThread__Go : public QThread {
public:
	long long __goId;
	unsigned char __goOn[9] = {};
	template<typename... A> QThread__Go(long long id, A&&... a) : QThread(std::forward<A>(a)...), __goId(id) {}
	using QThread::exec;
	using QThread::sender;
	using QThread::senderSignalIndex;
	using QThread::receivers;
	using QThread::isSignalConnected;
	void cppgoBase_run() { QThread::run(); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QThread::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QThread::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QThread::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QThread::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QThread::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QThread::metaObject();
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QThread::event(event);
	}
	void run() override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, -8524566317707732406LL, &__h); if (__h) return; }
		QThread::run();
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QThread::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QThread::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QThread::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QThread::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QThread::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QThread::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QThread;
static std::mutex __connsMu_QThread;

extern "C" {

void* QThread_base(void* self) { return (void*)static_cast<QObject*>((QThread*)self); }
void QThread__set_override(void* self, int idx) { ((QThread__Go*)self)->__goOn[idx] = 1; }
void* QThread_fromQObject(void* o) { return (void*)qobject_cast<QThread*>((QObject*)o); }
void* QThread_new(long long __goid) {
	return new QThread__Go(__goid);
}
void* QThread_new2(long long __goid, void* parent) {
	return new QThread__Go(__goid, (QObject*)parent);
}
char* QThread_tr(const char* s, long long* __out_len) {
	auto&& __r = QThread::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QThread_currentThread() {
	auto&& __r = QThread::currentThread();
	return (void*)(__r);
}
int QThread_isMainThread() {
	auto&& __r = QThread::isMainThread();
	return __r;
}
int QThread_idealThreadCount() {
	auto&& __r = QThread::idealThreadCount();
	return __r;
}
void QThread_yieldCurrentThread() {
	QThread::yieldCurrentThread();
}
void QThread_setPriority(void* self, int priority) {
	((QThread*)self)->setPriority((QThread::Priority)(priority));
}
int QThread_priority(void* self) {
	auto&& __r = ((QThread*)self)->priority();
	return (int)(__r);
}
int QThread_isFinished(void* self) {
	auto&& __r = ((QThread*)self)->isFinished();
	return __r;
}
int QThread_isRunning(void* self) {
	auto&& __r = ((QThread*)self)->isRunning();
	return __r;
}
void QThread_requestInterruption(void* self) {
	((QThread*)self)->requestInterruption();
}
int QThread_isInterruptionRequested(void* self) {
	auto&& __r = ((QThread*)self)->isInterruptionRequested();
	return __r;
}
void QThread_setStackSize(void* self, unsigned int stackSize) {
	((QThread*)self)->setStackSize(stackSize);
}
unsigned int QThread_stackSize(void* self) {
	auto&& __r = ((QThread*)self)->stackSize();
	return __r;
}
void* QThread_eventDispatcher(void* self) {
	auto&& __r = ((QThread*)self)->eventDispatcher();
	return (void*)(__r);
}
void QThread_setEventDispatcher(void* self, void* eventDispatcher) {
	((QThread*)self)->setEventDispatcher((QAbstractEventDispatcher*)eventDispatcher);
}
int QThread_loopLevel(void* self) {
	auto&& __r = ((QThread*)self)->loopLevel();
	return __r;
}
int QThread_isCurrentThread(void* self) {
	auto&& __r = ((QThread*)self)->isCurrentThread();
	return __r;
}
void QThread_setServiceLevel(void* self, int serviceLevel) {
	((QThread*)self)->setServiceLevel((QThread::QualityOfService)(serviceLevel));
}
int QThread_serviceLevel(void* self) {
	auto&& __r = ((QThread*)self)->serviceLevel();
	return (int)(__r);
}
void QThread_start(void* self) {
	((QThread*)self)->start();
}
void QThread_terminate(void* self) {
	((QThread*)self)->terminate();
}
void QThread_exit(void* self) {
	((QThread*)self)->exit();
}
void QThread_quit(void* self) {
	((QThread*)self)->quit();
}
int QThread_wait(void* self) {
	auto&& __r = ((QThread*)self)->wait();
	return __r;
}
int QThread_wait2(void* self, unsigned long time) {
	auto&& __r = ((QThread*)self)->wait(time);
	return __r;
}
void QThread_sleep(unsigned long arg1) {
	QThread::sleep(arg1);
}
void QThread_msleep(unsigned long arg1) {
	QThread::msleep(arg1);
}
void QThread_usleep(unsigned long arg1) {
	QThread::usleep(arg1);
}
void QThread_sleep2(long long nsec) {
	QThread::sleep(std::chrono::nanoseconds((long long)nsec));
}
int QThread_exec(void* self) {
	auto&& __r = ((QThread__Go*)self)->exec();
	return __r;
}
void* QThread_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QThread::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QThread_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QThread::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QThread_sender(void* self) {
	auto&& __r = ((QThread__Go*)self)->sender();
	return (void*)(__r);
}
int QThread_senderSignalIndex(void* self) {
	auto&& __r = ((QThread__Go*)self)->senderSignalIndex();
	return __r;
}
int QThread_receivers(void* self, const char* signal) {
	auto&& __r = ((QThread__Go*)self)->receivers(signal);
	return __r;
}
int QThread_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QThread__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QThread_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QThread::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QThread_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QThread::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QThread_start2(void* self, int arg1) {
	((QThread*)self)->start((QThread::Priority)(arg1));
}
void QThread_exit2(void* self, int retcode) {
	((QThread*)self)->exit(retcode);
}
int QThread_wait3(void* self, void* deadline) {
	auto&& __r = ((QThread*)self)->wait(*(QDeadlineTimer*)deadline);
	return __r;
}
void* QThread_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QThread::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QThread_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QThread::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QThread_metaObject_default(void* self) {
	auto&& __r = ((QThread*)self)->QThread::metaObject();
	return (void*)(__r);
}
int QThread_event_default(void* self, void* event) {
	auto&& __r = ((QThread*)self)->QThread::event((QEvent*)event);
	return __r;
}
void QThread_run_default(void* self) {
	((QThread__Go*)self)->cppgoBase_run();
}
int QThread_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QThread*)self)->QThread::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QThread_timerEvent_default(void* self, void* event) {
	((QThread__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QThread_childEvent_default(void* self, void* event) {
	((QThread__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QThread_customEvent_default(void* self, void* event) {
	((QThread__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QThread_connectNotify_default(void* self, void* signal) {
	((QThread__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QThread_disconnectNotify_default(void* self, void* signal) {
	((QThread__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QThread_connect_started(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QThread*)self, &QThread::started, (QThread*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QThread);
	__conns_QThread[__id] = __c;
}
void QThread_connect_finished(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QThread*)self, &QThread::finished, (QThread*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QThread);
	__conns_QThread[__id] = __c;
}
void QThread_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QThread*)self, &QThread::destroyed, (QThread*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QThread);
	__conns_QThread[__id] = __c;
}
void QThread_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QThread*)self, &QThread::destroyed, (QThread*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QThread);
	__conns_QThread[__id] = __c;
}
void QThread_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QThread*)self, &QThread::objectNameChanged, (QThread*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QThread);
	__conns_QThread[__id] = __c;
}
void QThread__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QThread);
		auto __it = __conns_QThread.find(id);
		if (__it == __conns_QThread.end()) return;
		__c = __it->second;
		__conns_QThread.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QThread__Go_new(long long id, void* parent) { return new QThread__Go(id, (QObject*)parent); }
void QThread_delete(void* self) { delete (QThread*)self; }

}

#include <QtCore/qlockfile.h>
#include <QtCore/qstring.h>
#include "gen_qlockfile.h"

extern "C" {

void* QLockFile_new(const char* fileName, long long fileName_len) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	return new QLockFile(fileName_qstr);
}
char* QLockFile_fileName(void* self, long long* __out_len) {
	auto&& __r = ((QLockFile*)self)->fileName();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QLockFile_lock(void* self) {
	auto&& __r = ((QLockFile*)self)->lock();
	return __r;
}
int QLockFile_tryLock(void* self, int timeout) {
	auto&& __r = ((QLockFile*)self)->tryLock(timeout);
	return __r;
}
void QLockFile_unlock(void* self) {
	((QLockFile*)self)->unlock();
}
void QLockFile_setStaleLockTime(void* self, int arg1) {
	((QLockFile*)self)->setStaleLockTime(arg1);
}
int QLockFile_staleLockTime(void* self) {
	auto&& __r = ((QLockFile*)self)->staleLockTime();
	return __r;
}
int QLockFile_tryLock2(void* self) {
	auto&& __r = ((QLockFile*)self)->tryLock();
	return __r;
}
void QLockFile_setStaleLockTime2(void* self, long long value) {
	((QLockFile*)self)->setStaleLockTime(std::chrono::milliseconds((long long)value));
}
long long QLockFile_staleLockTimeAsDuration(void* self) {
	auto&& __r = ((QLockFile*)self)->staleLockTimeAsDuration();
	return (long long)(__r).count();
}
int QLockFile_isLocked(void* self) {
	auto&& __r = ((QLockFile*)self)->isLocked();
	return __r;
}
int QLockFile_getLockInfo(void* self, long long* pid, void* hostname, void* appname) {
	qint64 pid_local{};
	auto&& __r = ((QLockFile*)self)->getLockInfo(&pid_local, (QString*)hostname, (QString*)appname);
	*pid = pid_local;
	return __r;
}
int QLockFile_removeStaleLockFile(void* self) {
	auto&& __r = ((QLockFile*)self)->removeStaleLockFile();
	return __r;
}
int QLockFile_error(void* self) {
	auto&& __r = ((QLockFile*)self)->error();
	return (int)(__r);
}
int QLockFile_tryLock3(void* self, long long timeout) {
	auto&& __r = ((QLockFile*)self)->tryLock(std::chrono::milliseconds((long long)timeout));
	return __r;
}
void QLockFile_delete(void* self) { delete (QLockFile*)self; }

}

#include <QtCore/qpermissions.h>
#include "gen_qmicrophonepermission.h"

extern "C" {

void* QMicrophonePermission_new() {
	return new QMicrophonePermission();
}
void* QMicrophonePermission_new2(void* other) {
	return new QMicrophonePermission(*(QMicrophonePermission*)other);
}
void* QMicrophonePermission_OperatorAssign(void* self, void* other) {
	auto&& __r = ((QMicrophonePermission*)self)->operator=(*(QMicrophonePermission*)other);
	return (void*)&(__r);
}
void QMicrophonePermission_swap(void* self, void* other) {
	((QMicrophonePermission*)self)->swap(*(QMicrophonePermission*)other);
}
void QMicrophonePermission_delete(void* self) { delete (QMicrophonePermission*)self; }

}

#include <QtCore/qsavefile.h>
#include <QtCore/qfiledevice.h>
#include <QtCore/qobject.h>
#include <QtCore/qstring.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qiodevicebase.h>
#include <QtCore/qdatetime.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qbytearrayview.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qvariant.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qsavefile.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QSaveFile__Go : public QSaveFile {
public:
	long long __goId;
	unsigned char __goOn[25] = {};
	template<typename... A> QSaveFile__Go(long long id, A&&... a) : QSaveFile(std::forward<A>(a)...), __goId(id) {}
	using QSaveFile::setOpenMode;
	using QSaveFile::setErrorString;
	using QSaveFile::sender;
	using QSaveFile::senderSignalIndex;
	using QSaveFile::receivers;
	using QSaveFile::isSignalConnected;
	qint64 cppgoBase_skipData(qint64 __a0) { return QSaveFile::skipData(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QSaveFile::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QSaveFile::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QSaveFile::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QSaveFile::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QSaveFile::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QSaveFile::metaObject();
	}
	QString fileName() const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virtp_charp) {
			long long __len = 0;
			char* __s = ((char*(*)(long long, long long, long long*, int*))cppgo_virtp_charp)(__goId, 8848495542063941256LL, &__len, &__h);
			if (__h) { QString __r = QString::fromUtf8(__s, __len); free(__s); return __r; }
		}
		return QSaveFile::fileName();
	}
	bool open(QFlags<QIODeviceBase::OpenModeFlag> flags) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, -554057428128861975LL, (int)(flags), &__h); if (__h) return __r; }
		return QSaveFile::open(flags);
	}
	bool isSequential() const override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, 8888098400087344862LL, &__h); if (__h) return __r; }
		return QSaveFile::isSequential();
	}
	qint64 pos() const override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_longlong) { long long __r = ((long long(*)(long long, long long, int*))cppgo_virt_longlong)(__goId, 8618582476283917417LL, &__h); if (__h) return __r; }
		return QSaveFile::pos();
	}
	bool seek(qint64 offset) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_int_longlong) { int __r = ((int(*)(long long, long long, long long, int*))cppgo_virt_int_longlong)(__goId, 3241140482992819217LL, offset, &__h); if (__h) return __r; }
		return QSaveFile::seek(offset);
	}
	bool atEnd() const override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -1487528724786208339LL, &__h); if (__h) return __r; }
		return QSaveFile::atEnd();
	}
	qint64 size() const override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_longlong) { long long __r = ((long long(*)(long long, long long, int*))cppgo_virt_longlong)(__goId, 5614464919154503228LL, &__h); if (__h) return __r; }
		return QSaveFile::size();
	}
	bool resize(qint64 sz) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_int_longlong) { int __r = ((int(*)(long long, long long, long long, int*))cppgo_virt_int_longlong)(__goId, -4890198374319873567LL, sz, &__h); if (__h) return __r; }
		return QSaveFile::resize(sz);
	}
	QFlags<QFileDevice::Permission> permissions() const override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -7517927236947282013LL, &__h); if (__h) return QFlags<QFileDevice::Permission>(__r); }
		return QSaveFile::permissions();
	}
	bool setPermissions(QFlags<QFileDevice::Permission> permissionSpec) override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 420407675554349589LL, (int)(permissionSpec), &__h); if (__h) return __r; }
		return QSaveFile::setPermissions(permissionSpec);
	}
	bool reset() override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, 1086335023529244512LL, &__h); if (__h) return __r; }
		return QSaveFile::reset();
	}
	qint64 bytesAvailable() const override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_longlong) { long long __r = ((long long(*)(long long, long long, int*))cppgo_virt_longlong)(__goId, 4814449630054007943LL, &__h); if (__h) return __r; }
		return QSaveFile::bytesAvailable();
	}
	qint64 bytesToWrite() const override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_longlong) { long long __r = ((long long(*)(long long, long long, int*))cppgo_virt_longlong)(__goId, 1024747515116322948LL, &__h); if (__h) return __r; }
		return QSaveFile::bytesToWrite();
	}
	bool canReadLine() const override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -1250682508954504761LL, &__h); if (__h) return __r; }
		return QSaveFile::canReadLine();
	}
	bool waitForReadyRead(int msecs) override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 8356976193102722774LL, msecs, &__h); if (__h) return __r; }
		return QSaveFile::waitForReadyRead(msecs);
	}
	bool waitForBytesWritten(int msecs) override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 7726833620054767889LL, msecs, &__h); if (__h) return __r; }
		return QSaveFile::waitForBytesWritten(msecs);
	}
	qint64 skipData(qint64 maxSize) override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_longlong_longlong) { long long __r = ((long long(*)(long long, long long, long long, int*))cppgo_virt_longlong_longlong)(__goId, 8178738247667652082LL, maxSize, &__h); if (__h) return __r; }
		return QSaveFile::skipData(maxSize);
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QSaveFile::event(event);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QSaveFile::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QSaveFile::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QSaveFile::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QSaveFile::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QSaveFile::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QSaveFile::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QSaveFile;
static std::mutex __connsMu_QSaveFile;

extern "C" {

void* QSaveFile_base(void* self) { return (void*)static_cast<QFileDevice*>((QSaveFile*)self); }
void QSaveFile_chain(void* self, void** out) {
	QSaveFile* __p = (QSaveFile*)self;
	QFileDevice* __b0 = static_cast<QFileDevice*>(__p); out[0] = (void*)__b0;
	QIODevice* __b1 = static_cast<QIODevice*>(__b0); out[1] = (void*)__b1;
	QObject* __b2 = static_cast<QObject*>(__b1); out[2] = (void*)__b2;
}
void QSaveFile__set_override(void* self, int idx) { ((QSaveFile__Go*)self)->__goOn[idx] = 1; }
void* QSaveFile_fromQObject(void* o) { return (void*)qobject_cast<QSaveFile*>((QObject*)o); }
void* QSaveFile_new(long long __goid) {
	return new QSaveFile__Go(__goid);
}
void* QSaveFile_new2(long long __goid, const char* name, long long name_len) {
	QString name_qstr = QString::fromUtf8(name, name_len);
	return new QSaveFile__Go(__goid, name_qstr);
}
void* QSaveFile_new3(long long __goid, void* parent) {
	return new QSaveFile__Go(__goid, (QObject*)parent);
}
void* QSaveFile_new4(long long __goid, const char* name, long long name_len, void* parent) {
	QString name_qstr = QString::fromUtf8(name, name_len);
	return new QSaveFile__Go(__goid, name_qstr, (QObject*)parent);
}
char* QSaveFile_tr(const char* s, long long* __out_len) {
	auto&& __r = QSaveFile::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QSaveFile_setFileName(void* self, const char* name, long long name_len) {
	QString name_qstr = QString::fromUtf8(name, name_len);
	((QSaveFile*)self)->setFileName(name_qstr);
}
int QSaveFile_commit(void* self) {
	auto&& __r = ((QSaveFile*)self)->commit();
	return __r;
}
void QSaveFile_cancelWriting(void* self) {
	((QSaveFile*)self)->cancelWriting();
}
void QSaveFile_setDirectWriteFallback(void* self, int enabled) {
	((QSaveFile*)self)->setDirectWriteFallback(enabled);
}
int QSaveFile_directWriteFallback(void* self) {
	auto&& __r = ((QSaveFile*)self)->directWriteFallback();
	return __r;
}
void QSaveFile_setOpenMode(void* self, int openMode) {
	((QSaveFile__Go*)self)->setOpenMode(QFlags<QIODeviceBase::OpenModeFlag>(openMode));
}
void QSaveFile_setErrorString(void* self, const char* errorString, long long errorString_len) {
	QString errorString_qstr = QString::fromUtf8(errorString, errorString_len);
	((QSaveFile__Go*)self)->setErrorString(errorString_qstr);
}
void* QSaveFile_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QSaveFile::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QSaveFile_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QSaveFile::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QSaveFile_sender(void* self) {
	auto&& __r = ((QSaveFile__Go*)self)->sender();
	return (void*)(__r);
}
int QSaveFile_senderSignalIndex(void* self) {
	auto&& __r = ((QSaveFile__Go*)self)->senderSignalIndex();
	return __r;
}
int QSaveFile_receivers(void* self, const char* signal) {
	auto&& __r = ((QSaveFile__Go*)self)->receivers(signal);
	return __r;
}
int QSaveFile_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QSaveFile__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QSaveFile_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QSaveFile::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QSaveFile_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QSaveFile::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QSaveFile_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QSaveFile::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QSaveFile_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QSaveFile::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QSaveFile_metaObject_default(void* self) {
	auto&& __r = ((QSaveFile*)self)->QSaveFile::metaObject();
	return (void*)(__r);
}
char* QSaveFile_fileName_default(void* self, long long* __out_len) {
	auto&& __r = ((QSaveFile*)self)->QSaveFile::fileName();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QSaveFile_open_default(void* self, int flags) {
	auto&& __r = ((QSaveFile*)self)->QSaveFile::open(QFlags<QIODeviceBase::OpenModeFlag>(flags));
	return __r;
}
int QSaveFile_isSequential_default(void* self) {
	auto&& __r = ((QSaveFile*)self)->QSaveFile::isSequential();
	return __r;
}
long long QSaveFile_pos_default(void* self) {
	auto&& __r = ((QSaveFile*)self)->QSaveFile::pos();
	return __r;
}
int QSaveFile_seek_default(void* self, long long offset) {
	auto&& __r = ((QSaveFile*)self)->QSaveFile::seek(offset);
	return __r;
}
int QSaveFile_atEnd_default(void* self) {
	auto&& __r = ((QSaveFile*)self)->QSaveFile::atEnd();
	return __r;
}
long long QSaveFile_size_default(void* self) {
	auto&& __r = ((QSaveFile*)self)->QSaveFile::size();
	return __r;
}
int QSaveFile_resize_default(void* self, long long sz) {
	auto&& __r = ((QSaveFile*)self)->QSaveFile::resize(sz);
	return __r;
}
int QSaveFile_permissions_default(void* self) {
	auto&& __r = ((QSaveFile*)self)->QSaveFile::permissions();
	return (int)(__r);
}
int QSaveFile_setPermissions_default(void* self, int permissionSpec) {
	auto&& __r = ((QSaveFile*)self)->QSaveFile::setPermissions(QFlags<QFileDevice::Permission>(permissionSpec));
	return __r;
}
int QSaveFile_reset_default(void* self) {
	auto&& __r = ((QSaveFile*)self)->QSaveFile::reset();
	return __r;
}
long long QSaveFile_bytesAvailable_default(void* self) {
	auto&& __r = ((QSaveFile*)self)->QSaveFile::bytesAvailable();
	return __r;
}
long long QSaveFile_bytesToWrite_default(void* self) {
	auto&& __r = ((QSaveFile*)self)->QSaveFile::bytesToWrite();
	return __r;
}
int QSaveFile_canReadLine_default(void* self) {
	auto&& __r = ((QSaveFile*)self)->QSaveFile::canReadLine();
	return __r;
}
int QSaveFile_waitForReadyRead_default(void* self, int msecs) {
	auto&& __r = ((QSaveFile*)self)->QSaveFile::waitForReadyRead(msecs);
	return __r;
}
int QSaveFile_waitForBytesWritten_default(void* self, int msecs) {
	auto&& __r = ((QSaveFile*)self)->QSaveFile::waitForBytesWritten(msecs);
	return __r;
}
long long QSaveFile_skipData_default(void* self, long long maxSize) {
	auto&& __r = ((QSaveFile__Go*)self)->cppgoBase_skipData(maxSize);
	return __r;
}
int QSaveFile_event_default(void* self, void* event) {
	auto&& __r = ((QSaveFile*)self)->QSaveFile::event((QEvent*)event);
	return __r;
}
int QSaveFile_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QSaveFile*)self)->QSaveFile::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QSaveFile_timerEvent_default(void* self, void* event) {
	((QSaveFile__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QSaveFile_childEvent_default(void* self, void* event) {
	((QSaveFile__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QSaveFile_customEvent_default(void* self, void* event) {
	((QSaveFile__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QSaveFile_connectNotify_default(void* self, void* signal) {
	((QSaveFile__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QSaveFile_disconnectNotify_default(void* self, void* signal) {
	((QSaveFile__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QSaveFile_connect_readyRead(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSaveFile*)self, &QSaveFile::readyRead, (QSaveFile*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSaveFile);
	__conns_QSaveFile[__id] = __c;
}
void QSaveFile_connect_channelReadyRead(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSaveFile*)self, &QSaveFile::channelReadyRead, (QSaveFile*)self, [__id](int channel){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, channel);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSaveFile);
	__conns_QSaveFile[__id] = __c;
}
void QSaveFile_connect_bytesWritten(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSaveFile*)self, &QSaveFile::bytesWritten, (QSaveFile*)self, [__id](qint64 bytes){
		if (cppgo_sig_void_longlong) ((void(*)(long long, long long))cppgo_sig_void_longlong)(__id, bytes);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSaveFile);
	__conns_QSaveFile[__id] = __c;
}
void QSaveFile_connect_channelBytesWritten(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSaveFile*)self, &QSaveFile::channelBytesWritten, (QSaveFile*)self, [__id](int channel, qint64 bytes){
		if (cppgo_sig_void_int_longlong) ((void(*)(long long, int, long long))cppgo_sig_void_int_longlong)(__id, channel, bytes);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSaveFile);
	__conns_QSaveFile[__id] = __c;
}
void QSaveFile_connect_aboutToClose(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSaveFile*)self, &QSaveFile::aboutToClose, (QSaveFile*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSaveFile);
	__conns_QSaveFile[__id] = __c;
}
void QSaveFile_connect_readChannelFinished(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSaveFile*)self, &QSaveFile::readChannelFinished, (QSaveFile*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSaveFile);
	__conns_QSaveFile[__id] = __c;
}
void QSaveFile_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSaveFile*)self, &QSaveFile::destroyed, (QSaveFile*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSaveFile);
	__conns_QSaveFile[__id] = __c;
}
void QSaveFile_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSaveFile*)self, &QSaveFile::destroyed, (QSaveFile*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSaveFile);
	__conns_QSaveFile[__id] = __c;
}
void QSaveFile_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSaveFile*)self, &QSaveFile::objectNameChanged, (QSaveFile*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSaveFile);
	__conns_QSaveFile[__id] = __c;
}
void QSaveFile__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QSaveFile);
		auto __it = __conns_QSaveFile.find(id);
		if (__it == __conns_QSaveFile.end()) return;
		__c = __it->second;
		__conns_QSaveFile.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QSaveFile__Go_new(long long id, void* parent) { return new QSaveFile__Go(id, (QObject*)parent); }
void QSaveFile_delete(void* self) { delete (QSaveFile*)self; }

}

#include <QtCore/qtransposeproxymodel.h>
#include <QtCore/qabstractproxymodel.h>
#include <QtCore/qobject.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qabstractitemmodel.h>
#include <QtCore/qvariant.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qsize.h>
#include <QtCore/qitemselectionmodel.h>
#include <QtCore/qmimedata.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qdatastream.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qtransposeproxymodel.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QTransposeProxyModel__Go : public QTransposeProxyModel {
public:
	long long __goId;
	unsigned char __goOn[47] = {};
	template<typename... A> QTransposeProxyModel__Go(long long id, A&&... a) : QTransposeProxyModel(std::forward<A>(a)...), __goId(id) {}
	using QTransposeProxyModel::createIndex;
	using QTransposeProxyModel::encodeData;
	using QTransposeProxyModel::decodeData;
	using QTransposeProxyModel::beginInsertRows;
	using QTransposeProxyModel::endInsertRows;
	using QTransposeProxyModel::beginRemoveRows;
	using QTransposeProxyModel::endRemoveRows;
	using QTransposeProxyModel::beginMoveRows;
	using QTransposeProxyModel::endMoveRows;
	using QTransposeProxyModel::beginInsertColumns;
	using QTransposeProxyModel::endInsertColumns;
	using QTransposeProxyModel::beginRemoveColumns;
	using QTransposeProxyModel::endRemoveColumns;
	using QTransposeProxyModel::beginMoveColumns;
	using QTransposeProxyModel::endMoveColumns;
	using QTransposeProxyModel::beginResetModel;
	using QTransposeProxyModel::endResetModel;
	using QTransposeProxyModel::changePersistentIndex;
	using QTransposeProxyModel::changePersistentIndexList;
	using QTransposeProxyModel::persistentIndexList;
	using QTransposeProxyModel::sender;
	using QTransposeProxyModel::senderSignalIndex;
	using QTransposeProxyModel::receivers;
	using QTransposeProxyModel::isSignalConnected;
	void cppgoBase_resetInternalData() { QTransposeProxyModel::resetInternalData(); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QTransposeProxyModel::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QTransposeProxyModel::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QTransposeProxyModel::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QTransposeProxyModel::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QTransposeProxyModel::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QTransposeProxyModel::metaObject();
	}
	void setSourceModel(QAbstractItemModel* newSourceModel) override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7818871925547456643LL, (void*)newSourceModel, &__h); if (__h) return; }
		QTransposeProxyModel::setSourceModel(newSourceModel);
	}
	int rowCount(const QModelIndex& parent) const override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1960960171956884010LL, (void*)&parent, &__h); if (__h) return __r; }
		return QTransposeProxyModel::rowCount(parent);
	}
	int columnCount(const QModelIndex& parent) const override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, -8256963024669584834LL, (void*)&parent, &__h); if (__h) return __r; }
		return QTransposeProxyModel::columnCount(parent);
	}
	QVariant headerData(int section, Qt::Orientation orientation, int role) const override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_voidp_int_int_int) { void* __r = ((void*(*)(long long, long long, int, int, int, int*))cppgo_virt_voidp_int_int_int)(__goId, -3297140021896693688LL, section, (int)(orientation), role, &__h); if (__h) return *(QVariant*)__r; }
		return QTransposeProxyModel::headerData(section, orientation, role);
	}
	bool setHeaderData(int section, Qt::Orientation orientation, const QVariant& value, int role) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_int_int_int_voidp_int) { int __r = ((int(*)(long long, long long, int, int, void*, int, int*))cppgo_virt_int_int_int_voidp_int)(__goId, 4852200306590054850LL, section, (int)(orientation), (void*)&value, role, &__h); if (__h) return __r; }
		return QTransposeProxyModel::setHeaderData(section, orientation, value, role);
	}
	QSize span(const QModelIndex& index) const override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, -8395342916587040543LL, (void*)&index, &__h); if (__h) return *(QSize*)__r; }
		return QTransposeProxyModel::span(index);
	}
	QModelIndex mapFromSource(const QModelIndex& sourceIndex) const override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, -1717009778236344012LL, (void*)&sourceIndex, &__h); if (__h) return *(QModelIndex*)__r; }
		return QTransposeProxyModel::mapFromSource(sourceIndex);
	}
	QModelIndex mapToSource(const QModelIndex& proxyIndex) const override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, -2333321292016234891LL, (void*)&proxyIndex, &__h); if (__h) return *(QModelIndex*)__r; }
		return QTransposeProxyModel::mapToSource(proxyIndex);
	}
	QModelIndex parent(const QModelIndex& index) const override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 7371339085820260189LL, (void*)&index, &__h); if (__h) return *(QModelIndex*)__r; }
		return QTransposeProxyModel::parent(index);
	}
	QModelIndex index(int row, int column, const QModelIndex& parent) const override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_voidp_int_int_voidp) { void* __r = ((void*(*)(long long, long long, int, int, void*, int*))cppgo_virt_voidp_int_int_voidp)(__goId, -8948777187306027381LL, row, column, (void*)&parent, &__h); if (__h) return *(QModelIndex*)__r; }
		return QTransposeProxyModel::index(row, column, parent);
	}
	bool insertRows(int row, int count, const QModelIndex& parent) override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_int_int_int_voidp) { int __r = ((int(*)(long long, long long, int, int, void*, int*))cppgo_virt_int_int_int_voidp)(__goId, -2775729825962395603LL, row, count, (void*)&parent, &__h); if (__h) return __r; }
		return QTransposeProxyModel::insertRows(row, count, parent);
	}
	bool removeRows(int row, int count, const QModelIndex& parent) override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_int_int_int_voidp) { int __r = ((int(*)(long long, long long, int, int, void*, int*))cppgo_virt_int_int_int_voidp)(__goId, 5104072549182222656LL, row, count, (void*)&parent, &__h); if (__h) return __r; }
		return QTransposeProxyModel::removeRows(row, count, parent);
	}
	bool moveRows(const QModelIndex& sourceParent, int sourceRow, int count, const QModelIndex& destinationParent, int destinationChild) override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_int_voidp_int_int_voidp_int) { int __r = ((int(*)(long long, long long, void*, int, int, void*, int, int*))cppgo_virt_int_voidp_int_int_voidp_int)(__goId, -6074819356947405415LL, (void*)&sourceParent, sourceRow, count, (void*)&destinationParent, destinationChild, &__h); if (__h) return __r; }
		return QTransposeProxyModel::moveRows(sourceParent, sourceRow, count, destinationParent, destinationChild);
	}
	bool insertColumns(int column, int count, const QModelIndex& parent) override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_int_int_int_voidp) { int __r = ((int(*)(long long, long long, int, int, void*, int*))cppgo_virt_int_int_int_voidp)(__goId, -8973293846230792237LL, column, count, (void*)&parent, &__h); if (__h) return __r; }
		return QTransposeProxyModel::insertColumns(column, count, parent);
	}
	bool removeColumns(int column, int count, const QModelIndex& parent) override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_int_int_int_voidp) { int __r = ((int(*)(long long, long long, int, int, void*, int*))cppgo_virt_int_int_int_voidp)(__goId, 2078548737056540420LL, column, count, (void*)&parent, &__h); if (__h) return __r; }
		return QTransposeProxyModel::removeColumns(column, count, parent);
	}
	bool moveColumns(const QModelIndex& sourceParent, int sourceColumn, int count, const QModelIndex& destinationParent, int destinationChild) override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_int_voidp_int_int_voidp_int) { int __r = ((int(*)(long long, long long, void*, int, int, void*, int, int*))cppgo_virt_int_voidp_int_int_voidp_int)(__goId, -6321007591388309905LL, (void*)&sourceParent, sourceColumn, count, (void*)&destinationParent, destinationChild, &__h); if (__h) return __r; }
		return QTransposeProxyModel::moveColumns(sourceParent, sourceColumn, count, destinationParent, destinationChild);
	}
	void sort(int column, Qt::SortOrder order) override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_void_int_int) { ((void(*)(long long, long long, int, int, int*))cppgo_virt_void_int_int)(__goId, 6912441694453382161LL, column, (int)(order), &__h); if (__h) return; }
		QTransposeProxyModel::sort(column, order);
	}
	QItemSelection mapSelectionToSource(const QItemSelection& selection) const override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 2788628866280405837LL, (void*)&selection, &__h); if (__h) return *(QItemSelection*)__r; }
		return QTransposeProxyModel::mapSelectionToSource(selection);
	}
	QItemSelection mapSelectionFromSource(const QItemSelection& selection) const override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 9067522407071207004LL, (void*)&selection, &__h); if (__h) return *(QItemSelection*)__r; }
		return QTransposeProxyModel::mapSelectionFromSource(selection);
	}
	bool submit() override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -3788331537442685179LL, &__h); if (__h) return __r; }
		return QTransposeProxyModel::submit();
	}
	void revert() override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, -1522951607536117149LL, &__h); if (__h) return; }
		QTransposeProxyModel::revert();
	}
	QVariant data(const QModelIndex& proxyIndex, int role) const override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_voidp_voidp_int) { void* __r = ((void*(*)(long long, long long, void*, int, int*))cppgo_virt_voidp_voidp_int)(__goId, -8837375942098466299LL, (void*)&proxyIndex, role, &__h); if (__h) return *(QVariant*)__r; }
		return QTransposeProxyModel::data(proxyIndex, role);
	}
	QFlags<Qt::ItemFlag> flags(const QModelIndex& index) const override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1703382833545304812LL, (void*)&index, &__h); if (__h) return QFlags<Qt::ItemFlag>(__r); }
		return QTransposeProxyModel::flags(index);
	}
	bool setData(const QModelIndex& index, const QVariant& value, int role) override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_int_voidp_voidp_int) { int __r = ((int(*)(long long, long long, void*, void*, int, int*))cppgo_virt_int_voidp_voidp_int)(__goId, 35191454575202219LL, (void*)&index, (void*)&value, role, &__h); if (__h) return __r; }
		return QTransposeProxyModel::setData(index, value, role);
	}
	bool clearItemData(const QModelIndex& index) override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 8914190500520542409LL, (void*)&index, &__h); if (__h) return __r; }
		return QTransposeProxyModel::clearItemData(index);
	}
	QModelIndex buddy(const QModelIndex& index) const override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 6112860112248786499LL, (void*)&index, &__h); if (__h) return *(QModelIndex*)__r; }
		return QTransposeProxyModel::buddy(index);
	}
	bool canFetchMore(const QModelIndex& parent) const override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 8639262066473964150LL, (void*)&parent, &__h); if (__h) return __r; }
		return QTransposeProxyModel::canFetchMore(parent);
	}
	void fetchMore(const QModelIndex& parent) override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3553294593611130806LL, (void*)&parent, &__h); if (__h) return; }
		QTransposeProxyModel::fetchMore(parent);
	}
	bool hasChildren(const QModelIndex& parent) const override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 3369452259144514380LL, (void*)&parent, &__h); if (__h) return __r; }
		return QTransposeProxyModel::hasChildren(parent);
	}
	QModelIndex sibling(int row, int column, const QModelIndex& idx) const override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_voidp_int_int_voidp) { void* __r = ((void*(*)(long long, long long, int, int, void*, int*))cppgo_virt_voidp_int_int_voidp)(__goId, 397329346813782521LL, row, column, (void*)&idx, &__h); if (__h) return *(QModelIndex*)__r; }
		return QTransposeProxyModel::sibling(row, column, idx);
	}
	QMimeData* mimeData(const QList<QModelIndex>& indexes) const override {
		int __h = 0;
		std::vector<void*> indexes_v; indexes_v.reserve(indexes.size());
		for (long long __i = 0; __i < indexes.size(); __i++) indexes_v.push_back((void*)&indexes[__i]);
		if (__goOn[31] && cppgo_virt_voidp_voidpp_longlong) { void* __r = ((void*(*)(long long, long long, void**, long long, int*))cppgo_virt_voidp_voidpp_longlong)(__goId, 5175297091024974749LL, indexes_v.data(), (long long)indexes.size(), &__h); if (__h) return (QMimeData*)__r; }
		return QTransposeProxyModel::mimeData(indexes);
	}
	bool canDropMimeData(const QMimeData* data, Qt::DropAction action, int row, int column, const QModelIndex& parent) const override {
		int __h = 0;
		if (__goOn[32] && cppgo_virt_int_voidp_int_int_int_voidp) { int __r = ((int(*)(long long, long long, void*, int, int, int, void*, int*))cppgo_virt_int_voidp_int_int_int_voidp)(__goId, 731539533644753052LL, (void*)data, (int)(action), row, column, (void*)&parent, &__h); if (__h) return __r; }
		return QTransposeProxyModel::canDropMimeData(data, action, row, column, parent);
	}
	bool dropMimeData(const QMimeData* data, Qt::DropAction action, int row, int column, const QModelIndex& parent) override {
		int __h = 0;
		if (__goOn[33] && cppgo_virt_int_voidp_int_int_int_voidp) { int __r = ((int(*)(long long, long long, void*, int, int, int, void*, int*))cppgo_virt_int_voidp_int_int_int_voidp)(__goId, -9049994684305682212LL, (void*)data, (int)(action), row, column, (void*)&parent, &__h); if (__h) return __r; }
		return QTransposeProxyModel::dropMimeData(data, action, row, column, parent);
	}
	QList<QString> mimeTypes() const override {
		int __h = 0;
		if (__goOn[34] && cppgo_virtp_voidp) {
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
		return QTransposeProxyModel::mimeTypes();
	}
	QFlags<Qt::DropAction> supportedDragActions() const override {
		int __h = 0;
		if (__goOn[35] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, 3421601386088570762LL, &__h); if (__h) return QFlags<Qt::DropAction>(__r); }
		return QTransposeProxyModel::supportedDragActions();
	}
	QFlags<Qt::DropAction> supportedDropActions() const override {
		int __h = 0;
		if (__goOn[36] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -7784642176186821205LL, &__h); if (__h) return QFlags<Qt::DropAction>(__r); }
		return QTransposeProxyModel::supportedDropActions();
	}
	QList<QModelIndex> match(const QModelIndex& start, int role, const QVariant& value, int hits, QFlags<Qt::MatchFlag> flags) const override {
		int __h = 0;
		if (__goOn[37] && cppgo_virtp_voidp_voidp_int_voidp_int_int) {
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
		return QTransposeProxyModel::match(start, role, value, hits, flags);
	}
	void multiData(const QModelIndex& index, QModelRoleDataSpan roleDataSpan) const override {
		int __h = 0;
		if (__goOn[38] && cppgo_virt_void_voidp_voidp) { ((void(*)(long long, long long, void*, void*, int*))cppgo_virt_void_voidp_voidp)(__goId, 731311060944096642LL, (void*)&index, (void*)&roleDataSpan, &__h); if (__h) return; }
		QTransposeProxyModel::multiData(index, roleDataSpan);
	}
	void resetInternalData() override {
		int __h = 0;
		if (__goOn[39] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, -4510675770784051183LL, &__h); if (__h) return; }
		QTransposeProxyModel::resetInternalData();
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[40] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QTransposeProxyModel::event(event);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[41] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QTransposeProxyModel::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[42] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QTransposeProxyModel::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[43] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QTransposeProxyModel::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[44] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QTransposeProxyModel::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[45] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QTransposeProxyModel::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[46] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QTransposeProxyModel::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QTransposeProxyModel;
static std::mutex __connsMu_QTransposeProxyModel;

extern "C" {

void* QTransposeProxyModel_base(void* self) { return (void*)static_cast<QAbstractProxyModel*>((QTransposeProxyModel*)self); }
void QTransposeProxyModel_chain(void* self, void** out) {
	QTransposeProxyModel* __p = (QTransposeProxyModel*)self;
	QAbstractProxyModel* __b0 = static_cast<QAbstractProxyModel*>(__p); out[0] = (void*)__b0;
	QAbstractItemModel* __b1 = static_cast<QAbstractItemModel*>(__b0); out[1] = (void*)__b1;
	QObject* __b2 = static_cast<QObject*>(__b1); out[2] = (void*)__b2;
}
void QTransposeProxyModel__set_override(void* self, int idx) { ((QTransposeProxyModel__Go*)self)->__goOn[idx] = 1; }
void* QTransposeProxyModel_fromQObject(void* o) { return (void*)qobject_cast<QTransposeProxyModel*>((QObject*)o); }
void* QTransposeProxyModel_new(long long __goid) {
	return new QTransposeProxyModel__Go(__goid);
}
void* QTransposeProxyModel_new2(long long __goid, void* parent) {
	return new QTransposeProxyModel__Go(__goid, (QObject*)parent);
}
char* QTransposeProxyModel_tr(const char* s, long long* __out_len) {
	auto&& __r = QTransposeProxyModel::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QTransposeProxyModel_createIndex(void* self, int row, int column) {
	auto&& __r = ((QTransposeProxyModel__Go*)self)->createIndex(row, column);
	return new QModelIndex(__r);
}
void* QTransposeProxyModel_createIndex2(void* self, int row, int column, unsigned long long id) {
	auto&& __r = ((QTransposeProxyModel__Go*)self)->createIndex(row, column, id);
	return new QModelIndex(__r);
}
void QTransposeProxyModel_encodeData(void* self, void** indexes, long long indexes_len, void* stream) {
	QList<QModelIndex> indexes_list; indexes_list.reserve(indexes_len);
	for (long long __i = 0; __i < indexes_len; __i++) indexes_list.append(*(QModelIndex*)indexes[__i]);
	((QTransposeProxyModel__Go*)self)->encodeData(indexes_list, *(QDataStream*)stream);
}
int QTransposeProxyModel_decodeData(void* self, int row, int column, void* parent, void* stream) {
	auto&& __r = ((QTransposeProxyModel__Go*)self)->decodeData(row, column, *(QModelIndex*)parent, *(QDataStream*)stream);
	return __r;
}
void QTransposeProxyModel_beginInsertRows(void* self, void* parent, int first, int last) {
	((QTransposeProxyModel__Go*)self)->beginInsertRows(*(QModelIndex*)parent, first, last);
}
void QTransposeProxyModel_endInsertRows(void* self) {
	((QTransposeProxyModel__Go*)self)->endInsertRows();
}
void QTransposeProxyModel_beginRemoveRows(void* self, void* parent, int first, int last) {
	((QTransposeProxyModel__Go*)self)->beginRemoveRows(*(QModelIndex*)parent, first, last);
}
void QTransposeProxyModel_endRemoveRows(void* self) {
	((QTransposeProxyModel__Go*)self)->endRemoveRows();
}
int QTransposeProxyModel_beginMoveRows(void* self, void* sourceParent, int sourceFirst, int sourceLast, void* destinationParent, int destinationRow) {
	auto&& __r = ((QTransposeProxyModel__Go*)self)->beginMoveRows(*(QModelIndex*)sourceParent, sourceFirst, sourceLast, *(QModelIndex*)destinationParent, destinationRow);
	return __r;
}
void QTransposeProxyModel_endMoveRows(void* self) {
	((QTransposeProxyModel__Go*)self)->endMoveRows();
}
void QTransposeProxyModel_beginInsertColumns(void* self, void* parent, int first, int last) {
	((QTransposeProxyModel__Go*)self)->beginInsertColumns(*(QModelIndex*)parent, first, last);
}
void QTransposeProxyModel_endInsertColumns(void* self) {
	((QTransposeProxyModel__Go*)self)->endInsertColumns();
}
void QTransposeProxyModel_beginRemoveColumns(void* self, void* parent, int first, int last) {
	((QTransposeProxyModel__Go*)self)->beginRemoveColumns(*(QModelIndex*)parent, first, last);
}
void QTransposeProxyModel_endRemoveColumns(void* self) {
	((QTransposeProxyModel__Go*)self)->endRemoveColumns();
}
int QTransposeProxyModel_beginMoveColumns(void* self, void* sourceParent, int sourceFirst, int sourceLast, void* destinationParent, int destinationColumn) {
	auto&& __r = ((QTransposeProxyModel__Go*)self)->beginMoveColumns(*(QModelIndex*)sourceParent, sourceFirst, sourceLast, *(QModelIndex*)destinationParent, destinationColumn);
	return __r;
}
void QTransposeProxyModel_endMoveColumns(void* self) {
	((QTransposeProxyModel__Go*)self)->endMoveColumns();
}
void QTransposeProxyModel_beginResetModel(void* self) {
	((QTransposeProxyModel__Go*)self)->beginResetModel();
}
void QTransposeProxyModel_endResetModel(void* self) {
	((QTransposeProxyModel__Go*)self)->endResetModel();
}
void QTransposeProxyModel_changePersistentIndex(void* self, void* from, void* to) {
	((QTransposeProxyModel__Go*)self)->changePersistentIndex(*(QModelIndex*)from, *(QModelIndex*)to);
}
void QTransposeProxyModel_changePersistentIndexList(void* self, void** from, long long from_len, void** to, long long to_len) {
	QList<QModelIndex> from_list; from_list.reserve(from_len);
	for (long long __i = 0; __i < from_len; __i++) from_list.append(*(QModelIndex*)from[__i]);
	QList<QModelIndex> to_list; to_list.reserve(to_len);
	for (long long __i = 0; __i < to_len; __i++) to_list.append(*(QModelIndex*)to[__i]);
	((QTransposeProxyModel__Go*)self)->changePersistentIndexList(from_list, to_list);
}
void** QTransposeProxyModel_persistentIndexList(void* self, long long* __out_len) {
	auto&& __r = ((QTransposeProxyModel__Go*)self)->persistentIndexList();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QModelIndex(__e);
	return __arr;
}
void* QTransposeProxyModel_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QTransposeProxyModel::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QTransposeProxyModel_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QTransposeProxyModel::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QTransposeProxyModel_sender(void* self) {
	auto&& __r = ((QTransposeProxyModel__Go*)self)->sender();
	return (void*)(__r);
}
int QTransposeProxyModel_senderSignalIndex(void* self) {
	auto&& __r = ((QTransposeProxyModel__Go*)self)->senderSignalIndex();
	return __r;
}
int QTransposeProxyModel_receivers(void* self, const char* signal) {
	auto&& __r = ((QTransposeProxyModel__Go*)self)->receivers(signal);
	return __r;
}
int QTransposeProxyModel_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QTransposeProxyModel__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QTransposeProxyModel_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QTransposeProxyModel::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QTransposeProxyModel_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QTransposeProxyModel::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QTransposeProxyModel_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QTransposeProxyModel::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QTransposeProxyModel_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QTransposeProxyModel::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QTransposeProxyModel_metaObject_default(void* self) {
	auto&& __r = ((QTransposeProxyModel*)self)->QTransposeProxyModel::metaObject();
	return (void*)(__r);
}
void QTransposeProxyModel_setSourceModel_default(void* self, void* newSourceModel) {
	((QTransposeProxyModel*)self)->QTransposeProxyModel::setSourceModel((QAbstractItemModel*)newSourceModel);
}
int QTransposeProxyModel_rowCount_default(void* self, void* parent) {
	auto&& __r = ((QTransposeProxyModel*)self)->QTransposeProxyModel::rowCount(*(QModelIndex*)parent);
	return __r;
}
int QTransposeProxyModel_columnCount_default(void* self, void* parent) {
	auto&& __r = ((QTransposeProxyModel*)self)->QTransposeProxyModel::columnCount(*(QModelIndex*)parent);
	return __r;
}
void* QTransposeProxyModel_headerData_default(void* self, int section, int orientation, int role) {
	auto&& __r = ((QTransposeProxyModel*)self)->QTransposeProxyModel::headerData(section, (Qt::Orientation)(orientation), role);
	return new QVariant(__r);
}
int QTransposeProxyModel_setHeaderData_default(void* self, int section, int orientation, void* value, int role) {
	auto&& __r = ((QTransposeProxyModel*)self)->QTransposeProxyModel::setHeaderData(section, (Qt::Orientation)(orientation), *(QVariant*)value, role);
	return __r;
}
void* QTransposeProxyModel_span_default(void* self, void* index) {
	auto&& __r = ((QTransposeProxyModel*)self)->QTransposeProxyModel::span(*(QModelIndex*)index);
	return new QSize(__r);
}
void* QTransposeProxyModel_mapFromSource_default(void* self, void* sourceIndex) {
	auto&& __r = ((QTransposeProxyModel*)self)->QTransposeProxyModel::mapFromSource(*(QModelIndex*)sourceIndex);
	return new QModelIndex(__r);
}
void* QTransposeProxyModel_mapToSource_default(void* self, void* proxyIndex) {
	auto&& __r = ((QTransposeProxyModel*)self)->QTransposeProxyModel::mapToSource(*(QModelIndex*)proxyIndex);
	return new QModelIndex(__r);
}
void* QTransposeProxyModel_parent_default(void* self, void* index) {
	auto&& __r = ((QTransposeProxyModel*)self)->QTransposeProxyModel::parent(*(QModelIndex*)index);
	return new QModelIndex(__r);
}
void* QTransposeProxyModel_index_default(void* self, int row, int column, void* parent) {
	auto&& __r = ((QTransposeProxyModel*)self)->QTransposeProxyModel::index(row, column, *(QModelIndex*)parent);
	return new QModelIndex(__r);
}
int QTransposeProxyModel_insertRows_default(void* self, int row, int count, void* parent) {
	auto&& __r = ((QTransposeProxyModel*)self)->QTransposeProxyModel::insertRows(row, count, *(QModelIndex*)parent);
	return __r;
}
int QTransposeProxyModel_removeRows_default(void* self, int row, int count, void* parent) {
	auto&& __r = ((QTransposeProxyModel*)self)->QTransposeProxyModel::removeRows(row, count, *(QModelIndex*)parent);
	return __r;
}
int QTransposeProxyModel_moveRows_default(void* self, void* sourceParent, int sourceRow, int count, void* destinationParent, int destinationChild) {
	auto&& __r = ((QTransposeProxyModel*)self)->QTransposeProxyModel::moveRows(*(QModelIndex*)sourceParent, sourceRow, count, *(QModelIndex*)destinationParent, destinationChild);
	return __r;
}
int QTransposeProxyModel_insertColumns_default(void* self, int column, int count, void* parent) {
	auto&& __r = ((QTransposeProxyModel*)self)->QTransposeProxyModel::insertColumns(column, count, *(QModelIndex*)parent);
	return __r;
}
int QTransposeProxyModel_removeColumns_default(void* self, int column, int count, void* parent) {
	auto&& __r = ((QTransposeProxyModel*)self)->QTransposeProxyModel::removeColumns(column, count, *(QModelIndex*)parent);
	return __r;
}
int QTransposeProxyModel_moveColumns_default(void* self, void* sourceParent, int sourceColumn, int count, void* destinationParent, int destinationChild) {
	auto&& __r = ((QTransposeProxyModel*)self)->QTransposeProxyModel::moveColumns(*(QModelIndex*)sourceParent, sourceColumn, count, *(QModelIndex*)destinationParent, destinationChild);
	return __r;
}
void QTransposeProxyModel_sort_default(void* self, int column, int order) {
	((QTransposeProxyModel*)self)->QTransposeProxyModel::sort(column, (Qt::SortOrder)(order));
}
void* QTransposeProxyModel_mapSelectionToSource_default(void* self, void* selection) {
	auto&& __r = ((QTransposeProxyModel*)self)->QTransposeProxyModel::mapSelectionToSource(*(QItemSelection*)selection);
	return new QItemSelection(__r);
}
void* QTransposeProxyModel_mapSelectionFromSource_default(void* self, void* selection) {
	auto&& __r = ((QTransposeProxyModel*)self)->QTransposeProxyModel::mapSelectionFromSource(*(QItemSelection*)selection);
	return new QItemSelection(__r);
}
int QTransposeProxyModel_submit_default(void* self) {
	auto&& __r = ((QTransposeProxyModel*)self)->QTransposeProxyModel::submit();
	return __r;
}
void QTransposeProxyModel_revert_default(void* self) {
	((QTransposeProxyModel*)self)->QTransposeProxyModel::revert();
}
void* QTransposeProxyModel_data_default(void* self, void* proxyIndex, int role) {
	auto&& __r = ((QTransposeProxyModel*)self)->QTransposeProxyModel::data(*(QModelIndex*)proxyIndex, role);
	return new QVariant(__r);
}
int QTransposeProxyModel_flags_default(void* self, void* index) {
	auto&& __r = ((QTransposeProxyModel*)self)->QTransposeProxyModel::flags(*(QModelIndex*)index);
	return (int)(__r);
}
int QTransposeProxyModel_setData_default(void* self, void* index, void* value, int role) {
	auto&& __r = ((QTransposeProxyModel*)self)->QTransposeProxyModel::setData(*(QModelIndex*)index, *(QVariant*)value, role);
	return __r;
}
int QTransposeProxyModel_clearItemData_default(void* self, void* index) {
	auto&& __r = ((QTransposeProxyModel*)self)->QTransposeProxyModel::clearItemData(*(QModelIndex*)index);
	return __r;
}
void* QTransposeProxyModel_buddy_default(void* self, void* index) {
	auto&& __r = ((QTransposeProxyModel*)self)->QTransposeProxyModel::buddy(*(QModelIndex*)index);
	return new QModelIndex(__r);
}
int QTransposeProxyModel_canFetchMore_default(void* self, void* parent) {
	auto&& __r = ((QTransposeProxyModel*)self)->QTransposeProxyModel::canFetchMore(*(QModelIndex*)parent);
	return __r;
}
void QTransposeProxyModel_fetchMore_default(void* self, void* parent) {
	((QTransposeProxyModel*)self)->QTransposeProxyModel::fetchMore(*(QModelIndex*)parent);
}
int QTransposeProxyModel_hasChildren_default(void* self, void* parent) {
	auto&& __r = ((QTransposeProxyModel*)self)->QTransposeProxyModel::hasChildren(*(QModelIndex*)parent);
	return __r;
}
void* QTransposeProxyModel_sibling_default(void* self, int row, int column, void* idx) {
	auto&& __r = ((QTransposeProxyModel*)self)->QTransposeProxyModel::sibling(row, column, *(QModelIndex*)idx);
	return new QModelIndex(__r);
}
void* QTransposeProxyModel_mimeData_default(void* self, void** indexes, long long indexes_len) {
	QList<QModelIndex> indexes_list; indexes_list.reserve(indexes_len);
	for (long long __i = 0; __i < indexes_len; __i++) indexes_list.append(*(QModelIndex*)indexes[__i]);
	auto&& __r = ((QTransposeProxyModel*)self)->QTransposeProxyModel::mimeData(indexes_list);
	return (void*)(__r);
}
int QTransposeProxyModel_canDropMimeData_default(void* self, void* data, int action, int row, int column, void* parent) {
	auto&& __r = ((QTransposeProxyModel*)self)->QTransposeProxyModel::canDropMimeData((QMimeData*)data, (Qt::DropAction)(action), row, column, *(QModelIndex*)parent);
	return __r;
}
int QTransposeProxyModel_dropMimeData_default(void* self, void* data, int action, int row, int column, void* parent) {
	auto&& __r = ((QTransposeProxyModel*)self)->QTransposeProxyModel::dropMimeData((QMimeData*)data, (Qt::DropAction)(action), row, column, *(QModelIndex*)parent);
	return __r;
}
char** QTransposeProxyModel_mimeTypes_default(void* self, long long* __out_len) {
	auto&& __r = ((QTransposeProxyModel*)self)->QTransposeProxyModel::mimeTypes();
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
int QTransposeProxyModel_supportedDragActions_default(void* self) {
	auto&& __r = ((QTransposeProxyModel*)self)->QTransposeProxyModel::supportedDragActions();
	return (int)(__r);
}
int QTransposeProxyModel_supportedDropActions_default(void* self) {
	auto&& __r = ((QTransposeProxyModel*)self)->QTransposeProxyModel::supportedDropActions();
	return (int)(__r);
}
void** QTransposeProxyModel_match_default(void* self, void* start, int role, void* value, int hits, int flags, long long* __out_len) {
	auto&& __r = ((QTransposeProxyModel*)self)->QTransposeProxyModel::match(*(QModelIndex*)start, role, *(QVariant*)value, hits, QFlags<Qt::MatchFlag>(flags));
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QModelIndex(__e);
	return __arr;
}
void QTransposeProxyModel_multiData_default(void* self, void* index, void* roleDataSpan) {
	((QTransposeProxyModel*)self)->QTransposeProxyModel::multiData(*(QModelIndex*)index, *(QModelRoleDataSpan*)roleDataSpan);
}
void QTransposeProxyModel_resetInternalData_default(void* self) {
	((QTransposeProxyModel__Go*)self)->cppgoBase_resetInternalData();
}
int QTransposeProxyModel_event_default(void* self, void* event) {
	auto&& __r = ((QTransposeProxyModel*)self)->QTransposeProxyModel::event((QEvent*)event);
	return __r;
}
int QTransposeProxyModel_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QTransposeProxyModel*)self)->QTransposeProxyModel::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QTransposeProxyModel_timerEvent_default(void* self, void* event) {
	((QTransposeProxyModel__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QTransposeProxyModel_childEvent_default(void* self, void* event) {
	((QTransposeProxyModel__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QTransposeProxyModel_customEvent_default(void* self, void* event) {
	((QTransposeProxyModel__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QTransposeProxyModel_connectNotify_default(void* self, void* signal) {
	((QTransposeProxyModel__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QTransposeProxyModel_disconnectNotify_default(void* self, void* signal) {
	((QTransposeProxyModel__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QTransposeProxyModel_connect_sourceModelChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTransposeProxyModel*)self, &QTransposeProxyModel::sourceModelChanged, (QTransposeProxyModel*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTransposeProxyModel);
	__conns_QTransposeProxyModel[__id] = __c;
}
void QTransposeProxyModel_connect_headerDataChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTransposeProxyModel*)self, &QTransposeProxyModel::headerDataChanged, (QTransposeProxyModel*)self, [__id](Qt::Orientation orientation, int first, int last){
		if (cppgo_sig_void_int_int_int) ((void(*)(long long, int, int, int))cppgo_sig_void_int_int_int)(__id, (int)(orientation), first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTransposeProxyModel);
	__conns_QTransposeProxyModel[__id] = __c;
}
void QTransposeProxyModel_connect_layoutChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTransposeProxyModel*)self, &QTransposeProxyModel::layoutChanged, (QTransposeProxyModel*)self, [__id](const QList<QPersistentModelIndex>& parents, QAbstractItemModel::LayoutChangeHint hint){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTransposeProxyModel);
	__conns_QTransposeProxyModel[__id] = __c;
}
void QTransposeProxyModel_connect_layoutChanged2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTransposeProxyModel*)self, &QTransposeProxyModel::layoutChanged, (QTransposeProxyModel*)self, [__id](const QList<QPersistentModelIndex>& parents, QAbstractItemModel::LayoutChangeHint hint){
		std::vector<void*> parents_v; parents_v.reserve(parents.size());
		for (long long __i = 0; __i < parents.size(); __i++) parents_v.push_back((void*)&parents[__i]);
		if (cppgo_sig_void_voidpp_longlong_int) ((void(*)(long long, void**, long long, int))cppgo_sig_void_voidpp_longlong_int)(__id, parents_v.data(), (long long)parents.size(), (int)(hint));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTransposeProxyModel);
	__conns_QTransposeProxyModel[__id] = __c;
}
void QTransposeProxyModel_connect_layoutAboutToBeChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTransposeProxyModel*)self, &QTransposeProxyModel::layoutAboutToBeChanged, (QTransposeProxyModel*)self, [__id](const QList<QPersistentModelIndex>& parents, QAbstractItemModel::LayoutChangeHint hint){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTransposeProxyModel);
	__conns_QTransposeProxyModel[__id] = __c;
}
void QTransposeProxyModel_connect_layoutAboutToBeChanged2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTransposeProxyModel*)self, &QTransposeProxyModel::layoutAboutToBeChanged, (QTransposeProxyModel*)self, [__id](const QList<QPersistentModelIndex>& parents, QAbstractItemModel::LayoutChangeHint hint){
		std::vector<void*> parents_v; parents_v.reserve(parents.size());
		for (long long __i = 0; __i < parents.size(); __i++) parents_v.push_back((void*)&parents[__i]);
		if (cppgo_sig_void_voidpp_longlong_int) ((void(*)(long long, void**, long long, int))cppgo_sig_void_voidpp_longlong_int)(__id, parents_v.data(), (long long)parents.size(), (int)(hint));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTransposeProxyModel);
	__conns_QTransposeProxyModel[__id] = __c;
}
void QTransposeProxyModel_connect_rowsAboutToBeInserted(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTransposeProxyModel*)self, &QTransposeProxyModel::rowsAboutToBeInserted, (QTransposeProxyModel*)self, [__id](const QModelIndex& parent, int first, int last){
		if (cppgo_sig_void_voidp_int_int) ((void(*)(long long, void*, int, int))cppgo_sig_void_voidp_int_int)(__id, (void*)&parent, first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTransposeProxyModel);
	__conns_QTransposeProxyModel[__id] = __c;
}
void QTransposeProxyModel_connect_rowsInserted(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTransposeProxyModel*)self, &QTransposeProxyModel::rowsInserted, (QTransposeProxyModel*)self, [__id](const QModelIndex& parent, int first, int last){
		if (cppgo_sig_void_voidp_int_int) ((void(*)(long long, void*, int, int))cppgo_sig_void_voidp_int_int)(__id, (void*)&parent, first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTransposeProxyModel);
	__conns_QTransposeProxyModel[__id] = __c;
}
void QTransposeProxyModel_connect_rowsAboutToBeRemoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTransposeProxyModel*)self, &QTransposeProxyModel::rowsAboutToBeRemoved, (QTransposeProxyModel*)self, [__id](const QModelIndex& parent, int first, int last){
		if (cppgo_sig_void_voidp_int_int) ((void(*)(long long, void*, int, int))cppgo_sig_void_voidp_int_int)(__id, (void*)&parent, first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTransposeProxyModel);
	__conns_QTransposeProxyModel[__id] = __c;
}
void QTransposeProxyModel_connect_rowsRemoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTransposeProxyModel*)self, &QTransposeProxyModel::rowsRemoved, (QTransposeProxyModel*)self, [__id](const QModelIndex& parent, int first, int last){
		if (cppgo_sig_void_voidp_int_int) ((void(*)(long long, void*, int, int))cppgo_sig_void_voidp_int_int)(__id, (void*)&parent, first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTransposeProxyModel);
	__conns_QTransposeProxyModel[__id] = __c;
}
void QTransposeProxyModel_connect_columnsAboutToBeInserted(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTransposeProxyModel*)self, &QTransposeProxyModel::columnsAboutToBeInserted, (QTransposeProxyModel*)self, [__id](const QModelIndex& parent, int first, int last){
		if (cppgo_sig_void_voidp_int_int) ((void(*)(long long, void*, int, int))cppgo_sig_void_voidp_int_int)(__id, (void*)&parent, first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTransposeProxyModel);
	__conns_QTransposeProxyModel[__id] = __c;
}
void QTransposeProxyModel_connect_columnsInserted(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTransposeProxyModel*)self, &QTransposeProxyModel::columnsInserted, (QTransposeProxyModel*)self, [__id](const QModelIndex& parent, int first, int last){
		if (cppgo_sig_void_voidp_int_int) ((void(*)(long long, void*, int, int))cppgo_sig_void_voidp_int_int)(__id, (void*)&parent, first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTransposeProxyModel);
	__conns_QTransposeProxyModel[__id] = __c;
}
void QTransposeProxyModel_connect_columnsAboutToBeRemoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTransposeProxyModel*)self, &QTransposeProxyModel::columnsAboutToBeRemoved, (QTransposeProxyModel*)self, [__id](const QModelIndex& parent, int first, int last){
		if (cppgo_sig_void_voidp_int_int) ((void(*)(long long, void*, int, int))cppgo_sig_void_voidp_int_int)(__id, (void*)&parent, first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTransposeProxyModel);
	__conns_QTransposeProxyModel[__id] = __c;
}
void QTransposeProxyModel_connect_columnsRemoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTransposeProxyModel*)self, &QTransposeProxyModel::columnsRemoved, (QTransposeProxyModel*)self, [__id](const QModelIndex& parent, int first, int last){
		if (cppgo_sig_void_voidp_int_int) ((void(*)(long long, void*, int, int))cppgo_sig_void_voidp_int_int)(__id, (void*)&parent, first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTransposeProxyModel);
	__conns_QTransposeProxyModel[__id] = __c;
}
void QTransposeProxyModel_connect_modelAboutToBeReset(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTransposeProxyModel*)self, &QTransposeProxyModel::modelAboutToBeReset, (QTransposeProxyModel*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTransposeProxyModel);
	__conns_QTransposeProxyModel[__id] = __c;
}
void QTransposeProxyModel_connect_modelReset(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTransposeProxyModel*)self, &QTransposeProxyModel::modelReset, (QTransposeProxyModel*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTransposeProxyModel);
	__conns_QTransposeProxyModel[__id] = __c;
}
void QTransposeProxyModel_connect_rowsAboutToBeMoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTransposeProxyModel*)self, &QTransposeProxyModel::rowsAboutToBeMoved, (QTransposeProxyModel*)self, [__id](const QModelIndex& sourceParent, int sourceStart, int sourceEnd, const QModelIndex& destinationParent, int destinationRow){
		if (cppgo_sig_void_voidp_int_int_voidp_int) ((void(*)(long long, void*, int, int, void*, int))cppgo_sig_void_voidp_int_int_voidp_int)(__id, (void*)&sourceParent, sourceStart, sourceEnd, (void*)&destinationParent, destinationRow);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTransposeProxyModel);
	__conns_QTransposeProxyModel[__id] = __c;
}
void QTransposeProxyModel_connect_rowsMoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTransposeProxyModel*)self, &QTransposeProxyModel::rowsMoved, (QTransposeProxyModel*)self, [__id](const QModelIndex& sourceParent, int sourceStart, int sourceEnd, const QModelIndex& destinationParent, int destinationRow){
		if (cppgo_sig_void_voidp_int_int_voidp_int) ((void(*)(long long, void*, int, int, void*, int))cppgo_sig_void_voidp_int_int_voidp_int)(__id, (void*)&sourceParent, sourceStart, sourceEnd, (void*)&destinationParent, destinationRow);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTransposeProxyModel);
	__conns_QTransposeProxyModel[__id] = __c;
}
void QTransposeProxyModel_connect_columnsAboutToBeMoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTransposeProxyModel*)self, &QTransposeProxyModel::columnsAboutToBeMoved, (QTransposeProxyModel*)self, [__id](const QModelIndex& sourceParent, int sourceStart, int sourceEnd, const QModelIndex& destinationParent, int destinationColumn){
		if (cppgo_sig_void_voidp_int_int_voidp_int) ((void(*)(long long, void*, int, int, void*, int))cppgo_sig_void_voidp_int_int_voidp_int)(__id, (void*)&sourceParent, sourceStart, sourceEnd, (void*)&destinationParent, destinationColumn);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTransposeProxyModel);
	__conns_QTransposeProxyModel[__id] = __c;
}
void QTransposeProxyModel_connect_columnsMoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTransposeProxyModel*)self, &QTransposeProxyModel::columnsMoved, (QTransposeProxyModel*)self, [__id](const QModelIndex& sourceParent, int sourceStart, int sourceEnd, const QModelIndex& destinationParent, int destinationColumn){
		if (cppgo_sig_void_voidp_int_int_voidp_int) ((void(*)(long long, void*, int, int, void*, int))cppgo_sig_void_voidp_int_int_voidp_int)(__id, (void*)&sourceParent, sourceStart, sourceEnd, (void*)&destinationParent, destinationColumn);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTransposeProxyModel);
	__conns_QTransposeProxyModel[__id] = __c;
}
void QTransposeProxyModel_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTransposeProxyModel*)self, &QTransposeProxyModel::destroyed, (QTransposeProxyModel*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTransposeProxyModel);
	__conns_QTransposeProxyModel[__id] = __c;
}
void QTransposeProxyModel_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTransposeProxyModel*)self, &QTransposeProxyModel::destroyed, (QTransposeProxyModel*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTransposeProxyModel);
	__conns_QTransposeProxyModel[__id] = __c;
}
void QTransposeProxyModel_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTransposeProxyModel*)self, &QTransposeProxyModel::objectNameChanged, (QTransposeProxyModel*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTransposeProxyModel);
	__conns_QTransposeProxyModel[__id] = __c;
}
void QTransposeProxyModel__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QTransposeProxyModel);
		auto __it = __conns_QTransposeProxyModel.find(id);
		if (__it == __conns_QTransposeProxyModel.end()) return;
		__c = __it->second;
		__conns_QTransposeProxyModel.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QTransposeProxyModel__Go_new(long long id, void* parent) { return new QTransposeProxyModel__Go(id, (QObject*)parent); }
void QTransposeProxyModel_delete(void* self) { delete (QTransposeProxyModel*)self; }

}

#include <QtGui/qvectornd.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qpoint.h>
#include <QtCore/qvariant.h>
#include "gen_qvector2d.h"

extern "C" {

void* QVector2D_new() {
	return new QVector2D();
}
void* QVector2D_new2(int arg1) {
	return new QVector2D((Qt::Initialization)(arg1));
}
void* QVector2D_new3(float xpos, float ypos) {
	return new QVector2D(xpos, ypos);
}
void* QVector2D_new4(void* point) {
	return new QVector2D(*(QPoint*)point);
}
void* QVector2D_new5(void* point) {
	return new QVector2D(*(QPointF*)point);
}
void* QVector2D_new6(void* vector) {
	return new QVector2D(*(QVector3D*)vector);
}
void* QVector2D_new7(void* vector) {
	return new QVector2D(*(QVector4D*)vector);
}
void* QVector2D_new8(void* arg1) {
	return new QVector2D(*(QVector2D*)arg1);
}
int QVector2D_isNull(void* self) {
	auto&& __r = ((QVector2D*)self)->isNull();
	return __r;
}
float QVector2D_x(void* self) {
	auto&& __r = ((QVector2D*)self)->x();
	return __r;
}
float QVector2D_y(void* self) {
	auto&& __r = ((QVector2D*)self)->y();
	return __r;
}
void QVector2D_setX(void* self, float x) {
	((QVector2D*)self)->setX(x);
}
void QVector2D_setY(void* self, float y) {
	((QVector2D*)self)->setY(y);
}
float QVector2D_OperatorSubscript(void* self, int i) {
	auto&& __r = ((QVector2D*)self)->operator[](i);
	return __r;
}
float QVector2D_length(void* self) {
	auto&& __r = ((QVector2D*)self)->length();
	return __r;
}
float QVector2D_lengthSquared(void* self) {
	auto&& __r = ((QVector2D*)self)->lengthSquared();
	return __r;
}
void* QVector2D_normalized(void* self) {
	auto&& __r = ((QVector2D*)self)->normalized();
	return new QVector2D(__r);
}
void QVector2D_normalize(void* self) {
	((QVector2D*)self)->normalize();
}
float QVector2D_distanceToPoint(void* self, void* point) {
	auto&& __r = ((QVector2D*)self)->distanceToPoint(*(QVector2D*)point);
	return __r;
}
float QVector2D_distanceToLine(void* self, void* point, void* direction) {
	auto&& __r = ((QVector2D*)self)->distanceToLine(*(QVector2D*)point, *(QVector2D*)direction);
	return __r;
}
void* QVector2D_OperatorPlusAssign(void* self, void* vector) {
	auto&& __r = ((QVector2D*)self)->operator+=(*(QVector2D*)vector);
	return (void*)&(__r);
}
void* QVector2D_OperatorMinusAssign(void* self, void* vector) {
	auto&& __r = ((QVector2D*)self)->operator-=(*(QVector2D*)vector);
	return (void*)&(__r);
}
void* QVector2D_OperatorMultiplyAssign(void* self, float factor) {
	auto&& __r = ((QVector2D*)self)->operator*=(factor);
	return (void*)&(__r);
}
void* QVector2D_OperatorMultiplyAssign2(void* self, void* vector) {
	auto&& __r = ((QVector2D*)self)->operator*=(*(QVector2D*)vector);
	return (void*)&(__r);
}
void* QVector2D_OperatorDivideAssign(void* self, float divisor) {
	auto&& __r = ((QVector2D*)self)->operator/=(divisor);
	return (void*)&(__r);
}
void* QVector2D_OperatorDivideAssign2(void* self, void* vector) {
	auto&& __r = ((QVector2D*)self)->operator/=(*(QVector2D*)vector);
	return (void*)&(__r);
}
float QVector2D_dotProduct(void* v1, void* v2) {
	auto&& __r = QVector2D::dotProduct(*(QVector2D*)v1, *(QVector2D*)v2);
	return __r;
}
void* QVector2D_toVector3D(void* self) {
	auto&& __r = ((QVector2D*)self)->toVector3D();
	return new QVector3D(__r);
}
void* QVector2D_toVector4D(void* self) {
	auto&& __r = ((QVector2D*)self)->toVector4D();
	return new QVector4D(__r);
}
void* QVector2D_toPoint(void* self) {
	auto&& __r = ((QVector2D*)self)->toPoint();
	return new QPoint(__r);
}
void* QVector2D_toPointF(void* self) {
	auto&& __r = ((QVector2D*)self)->toPointF();
	return new QPointF(__r);
}
void* QVector2D_ToQVariant(void* self) {
	auto&& __r = static_cast<QVariant>(*(QVector2D*)self);
	return new QVariant(__r);
}
void* QVector2D_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QVector2D*)self)->operator=(*(QVector2D*)arg1);
	return (void*)&(__r);
}
int QVector2D_OperatorEqual(void* self, void* v2) {
	auto&& __r = (*(QVector2D*)self) == (*(QVector2D*)v2);
	return __r;
}
int QVector2D_OperatorNotEqual(void* self, void* v2) {
	auto&& __r = (*(QVector2D*)self) != (*(QVector2D*)v2);
	return __r;
}
void* QVector2D_OperatorPlus(void* self, void* v2) {
	auto&& __r = (*(QVector2D*)self) + (*(QVector2D*)v2);
	return new QVector2D(__r);
}
void* QVector2D_OperatorMinus(void* self, void* v2) {
	auto&& __r = (*(QVector2D*)self) - (*(QVector2D*)v2);
	return new QVector2D(__r);
}
void* QVector2D_OperatorMultiply(void* self, float factor) {
	auto&& __r = (*(QVector2D*)self) * (factor);
	return new QVector2D(__r);
}
void* QVector2D_OperatorMultiply2(void* self, void* v2) {
	auto&& __r = (*(QVector2D*)self) * (*(QVector2D*)v2);
	return new QVector2D(__r);
}
void* QVector2D_OperatorMinus2(void* self) {
	auto&& __r = -(*(QVector2D*)self);
	return new QVector2D(__r);
}
void* QVector2D_OperatorDivide(void* self, float divisor) {
	auto&& __r = (*(QVector2D*)self) / (divisor);
	return new QVector2D(__r);
}
void* QVector2D_OperatorDivide2(void* self, void* divisor) {
	auto&& __r = (*(QVector2D*)self) / (*(QVector2D*)divisor);
	return new QVector2D(__r);
}
void QVector2D_delete(void* self) { delete (QVector2D*)self; }

}

#include <QtGui/qevent.h>
#include <QtCore/qcoreevent.h>
#include <QtGui/qregion.h>
#include "gen_qexposeevent.h"

extern "C" {

void* QExposeEvent_base(void* self) { return (void*)static_cast<QEvent*>((QExposeEvent*)self); }
void* QExposeEvent_new(void* m_region) {
	return new QExposeEvent(*(QRegion*)m_region);
}
void* QExposeEvent_region(void* self) {
	auto&& __r = ((QExposeEvent*)self)->region();
	return (void*)&(__r);
}
int QExposeEvent_registerEventType() {
	auto&& __r = QExposeEvent::registerEventType();
	return __r;
}
int QExposeEvent_registerEventType2(int hint) {
	auto&& __r = QExposeEvent::registerEventType(hint);
	return __r;
}
void QExposeEvent_delete(void* self) { delete (QExposeEvent*)self; }

}

#include <QtGui/qevent.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qpoint.h>
#include <QtCore/qsize.h>
#include <QtCore/qrect.h>
#include "gen_qscrollprepareevent.h"

extern "C" {

void* QScrollPrepareEvent_base(void* self) { return (void*)static_cast<QEvent*>((QScrollPrepareEvent*)self); }
void* QScrollPrepareEvent_new(void* startPos) {
	return new QScrollPrepareEvent(*(QPointF*)startPos);
}
void* QScrollPrepareEvent_startPos(void* self) {
	auto&& __r = ((QScrollPrepareEvent*)self)->startPos();
	return new QPointF(__r);
}
void* QScrollPrepareEvent_viewportSize(void* self) {
	auto&& __r = ((QScrollPrepareEvent*)self)->viewportSize();
	return new QSizeF(__r);
}
void* QScrollPrepareEvent_contentPosRange(void* self) {
	auto&& __r = ((QScrollPrepareEvent*)self)->contentPosRange();
	return new QRectF(__r);
}
void* QScrollPrepareEvent_contentPos(void* self) {
	auto&& __r = ((QScrollPrepareEvent*)self)->contentPos();
	return new QPointF(__r);
}
void QScrollPrepareEvent_setViewportSize(void* self, void* size) {
	((QScrollPrepareEvent*)self)->setViewportSize(*(QSizeF*)size);
}
void QScrollPrepareEvent_setContentPosRange(void* self, void* rect) {
	((QScrollPrepareEvent*)self)->setContentPosRange(*(QRectF*)rect);
}
void QScrollPrepareEvent_setContentPos(void* self, void* pos) {
	((QScrollPrepareEvent*)self)->setContentPos(*(QPointF*)pos);
}
int QScrollPrepareEvent_registerEventType() {
	auto&& __r = QScrollPrepareEvent::registerEventType();
	return __r;
}
int QScrollPrepareEvent_registerEventType2(int hint) {
	auto&& __r = QScrollPrepareEvent::registerEventType(hint);
	return __r;
}
void QScrollPrepareEvent_delete(void* self) { delete (QScrollPrepareEvent*)self; }

}

#include <QtGui/qtextformat.h>
#include <QtGui/qfont.h>
#include <QtCore/qstring.h>
#include <QtCore/qvariant.h>
#include <QtGui/qcolor.h>
#include <QtGui/qpen.h>
#include <QtGui/qbrush.h>
#include <QtCore/qnamespace.h>
#include "gen_qtextcharformat.h"

extern "C" {

void* QTextCharFormat_base(void* self) { return (void*)static_cast<QTextFormat*>((QTextCharFormat*)self); }
void* QTextCharFormat_new() {
	return new QTextCharFormat();
}
void* QTextCharFormat_new2(void* arg1) {
	return new QTextCharFormat(*(QTextCharFormat*)arg1);
}
int QTextCharFormat_isValid(void* self) {
	auto&& __r = ((QTextCharFormat*)self)->isValid();
	return __r;
}
void QTextCharFormat_setFont(void* self, void* font) {
	((QTextCharFormat*)self)->setFont(*(QFont*)font);
}
void* QTextCharFormat_font(void* self) {
	auto&& __r = ((QTextCharFormat*)self)->font();
	return new QFont(__r);
}
void QTextCharFormat_setFontFamily(void* self, const char* family, long long family_len) {
	QString family_qstr = QString::fromUtf8(family, family_len);
	((QTextCharFormat*)self)->setFontFamily(family_qstr);
}
char* QTextCharFormat_fontFamily(void* self, long long* __out_len) {
	auto&& __r = ((QTextCharFormat*)self)->fontFamily();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QTextCharFormat_setFontFamilies(void* self, char** families, long long families_len) {
	QList<QString> families_list; families_list.reserve(families_len);
	for (long long __i = 0; __i < families_len; __i++) families_list.append(QString::fromUtf8(families[__i]));
	((QTextCharFormat*)self)->setFontFamilies(families_list);
}
void* QTextCharFormat_fontFamilies(void* self) {
	auto&& __r = ((QTextCharFormat*)self)->fontFamilies();
	return new QVariant(__r);
}
void QTextCharFormat_setFontStyleName(void* self, const char* styleName, long long styleName_len) {
	QString styleName_qstr = QString::fromUtf8(styleName, styleName_len);
	((QTextCharFormat*)self)->setFontStyleName(styleName_qstr);
}
void* QTextCharFormat_fontStyleName(void* self) {
	auto&& __r = ((QTextCharFormat*)self)->fontStyleName();
	return new QVariant(__r);
}
void QTextCharFormat_setFontPointSize(void* self, double size) {
	((QTextCharFormat*)self)->setFontPointSize(size);
}
double QTextCharFormat_fontPointSize(void* self) {
	auto&& __r = ((QTextCharFormat*)self)->fontPointSize();
	return __r;
}
void QTextCharFormat_setFontWeight(void* self, int weight) {
	((QTextCharFormat*)self)->setFontWeight(weight);
}
int QTextCharFormat_fontWeight(void* self) {
	auto&& __r = ((QTextCharFormat*)self)->fontWeight();
	return __r;
}
void QTextCharFormat_setFontItalic(void* self, int italic) {
	((QTextCharFormat*)self)->setFontItalic(italic);
}
int QTextCharFormat_fontItalic(void* self) {
	auto&& __r = ((QTextCharFormat*)self)->fontItalic();
	return __r;
}
void QTextCharFormat_setFontCapitalization(void* self, int capitalization) {
	((QTextCharFormat*)self)->setFontCapitalization((QFont::Capitalization)(capitalization));
}
int QTextCharFormat_fontCapitalization(void* self) {
	auto&& __r = ((QTextCharFormat*)self)->fontCapitalization();
	return (int)(__r);
}
void QTextCharFormat_setFontLetterSpacingType(void* self, int letterSpacingType) {
	((QTextCharFormat*)self)->setFontLetterSpacingType((QFont::SpacingType)(letterSpacingType));
}
int QTextCharFormat_fontLetterSpacingType(void* self) {
	auto&& __r = ((QTextCharFormat*)self)->fontLetterSpacingType();
	return (int)(__r);
}
void QTextCharFormat_setFontLetterSpacing(void* self, double spacing) {
	((QTextCharFormat*)self)->setFontLetterSpacing(spacing);
}
double QTextCharFormat_fontLetterSpacing(void* self) {
	auto&& __r = ((QTextCharFormat*)self)->fontLetterSpacing();
	return __r;
}
void QTextCharFormat_setFontWordSpacing(void* self, double spacing) {
	((QTextCharFormat*)self)->setFontWordSpacing(spacing);
}
double QTextCharFormat_fontWordSpacing(void* self) {
	auto&& __r = ((QTextCharFormat*)self)->fontWordSpacing();
	return __r;
}
void QTextCharFormat_setFontUnderline(void* self, int underline) {
	((QTextCharFormat*)self)->setFontUnderline(underline);
}
int QTextCharFormat_fontUnderline(void* self) {
	auto&& __r = ((QTextCharFormat*)self)->fontUnderline();
	return __r;
}
void QTextCharFormat_setFontOverline(void* self, int overline) {
	((QTextCharFormat*)self)->setFontOverline(overline);
}
int QTextCharFormat_fontOverline(void* self) {
	auto&& __r = ((QTextCharFormat*)self)->fontOverline();
	return __r;
}
void QTextCharFormat_setFontStrikeOut(void* self, int strikeOut) {
	((QTextCharFormat*)self)->setFontStrikeOut(strikeOut);
}
int QTextCharFormat_fontStrikeOut(void* self) {
	auto&& __r = ((QTextCharFormat*)self)->fontStrikeOut();
	return __r;
}
void QTextCharFormat_setUnderlineColor(void* self, void* color) {
	((QTextCharFormat*)self)->setUnderlineColor(*(QColor*)color);
}
void* QTextCharFormat_underlineColor(void* self) {
	auto&& __r = ((QTextCharFormat*)self)->underlineColor();
	return new QColor(__r);
}
void QTextCharFormat_setFontFixedPitch(void* self, int fixedPitch) {
	((QTextCharFormat*)self)->setFontFixedPitch(fixedPitch);
}
int QTextCharFormat_fontFixedPitch(void* self) {
	auto&& __r = ((QTextCharFormat*)self)->fontFixedPitch();
	return __r;
}
void QTextCharFormat_setFontStretch(void* self, int factor) {
	((QTextCharFormat*)self)->setFontStretch(factor);
}
int QTextCharFormat_fontStretch(void* self) {
	auto&& __r = ((QTextCharFormat*)self)->fontStretch();
	return __r;
}
void QTextCharFormat_setFontStyleHint(void* self, int hint) {
	((QTextCharFormat*)self)->setFontStyleHint((QFont::StyleHint)(hint));
}
void QTextCharFormat_setFontStyleStrategy(void* self, int strategy) {
	((QTextCharFormat*)self)->setFontStyleStrategy((QFont::StyleStrategy)(strategy));
}
int QTextCharFormat_fontStyleHint(void* self) {
	auto&& __r = ((QTextCharFormat*)self)->fontStyleHint();
	return (int)(__r);
}
int QTextCharFormat_fontStyleStrategy(void* self) {
	auto&& __r = ((QTextCharFormat*)self)->fontStyleStrategy();
	return (int)(__r);
}
void QTextCharFormat_setFontHintingPreference(void* self, int hintingPreference) {
	((QTextCharFormat*)self)->setFontHintingPreference((QFont::HintingPreference)(hintingPreference));
}
int QTextCharFormat_fontHintingPreference(void* self) {
	auto&& __r = ((QTextCharFormat*)self)->fontHintingPreference();
	return (int)(__r);
}
void QTextCharFormat_setFontKerning(void* self, int enable) {
	((QTextCharFormat*)self)->setFontKerning(enable);
}
int QTextCharFormat_fontKerning(void* self) {
	auto&& __r = ((QTextCharFormat*)self)->fontKerning();
	return __r;
}
void QTextCharFormat_setUnderlineStyle(void* self, int style) {
	((QTextCharFormat*)self)->setUnderlineStyle((QTextCharFormat::UnderlineStyle)(style));
}
int QTextCharFormat_underlineStyle(void* self) {
	auto&& __r = ((QTextCharFormat*)self)->underlineStyle();
	return (int)(__r);
}
void QTextCharFormat_setVerticalAlignment(void* self, int alignment) {
	((QTextCharFormat*)self)->setVerticalAlignment((QTextCharFormat::VerticalAlignment)(alignment));
}
int QTextCharFormat_verticalAlignment(void* self) {
	auto&& __r = ((QTextCharFormat*)self)->verticalAlignment();
	return (int)(__r);
}
void QTextCharFormat_setTextOutline(void* self, void* pen) {
	((QTextCharFormat*)self)->setTextOutline(*(QPen*)pen);
}
void* QTextCharFormat_textOutline(void* self) {
	auto&& __r = ((QTextCharFormat*)self)->textOutline();
	return new QPen(__r);
}
void QTextCharFormat_setToolTip(void* self, const char* tip, long long tip_len) {
	QString tip_qstr = QString::fromUtf8(tip, tip_len);
	((QTextCharFormat*)self)->setToolTip(tip_qstr);
}
char* QTextCharFormat_toolTip(void* self, long long* __out_len) {
	auto&& __r = ((QTextCharFormat*)self)->toolTip();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QTextCharFormat_setSuperScriptBaseline(void* self, double baseline) {
	((QTextCharFormat*)self)->setSuperScriptBaseline(baseline);
}
double QTextCharFormat_superScriptBaseline(void* self) {
	auto&& __r = ((QTextCharFormat*)self)->superScriptBaseline();
	return __r;
}
void QTextCharFormat_setSubScriptBaseline(void* self, double baseline) {
	((QTextCharFormat*)self)->setSubScriptBaseline(baseline);
}
double QTextCharFormat_subScriptBaseline(void* self) {
	auto&& __r = ((QTextCharFormat*)self)->subScriptBaseline();
	return __r;
}
void QTextCharFormat_setBaselineOffset(void* self, double baseline) {
	((QTextCharFormat*)self)->setBaselineOffset(baseline);
}
double QTextCharFormat_baselineOffset(void* self) {
	auto&& __r = ((QTextCharFormat*)self)->baselineOffset();
	return __r;
}
void QTextCharFormat_setAnchor(void* self, int anchor) {
	((QTextCharFormat*)self)->setAnchor(anchor);
}
int QTextCharFormat_isAnchor(void* self) {
	auto&& __r = ((QTextCharFormat*)self)->isAnchor();
	return __r;
}
void QTextCharFormat_setAnchorHref(void* self, const char* value, long long value_len) {
	QString value_qstr = QString::fromUtf8(value, value_len);
	((QTextCharFormat*)self)->setAnchorHref(value_qstr);
}
char* QTextCharFormat_anchorHref(void* self, long long* __out_len) {
	auto&& __r = ((QTextCharFormat*)self)->anchorHref();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QTextCharFormat_setAnchorNames(void* self, char** names, long long names_len) {
	QList<QString> names_list; names_list.reserve(names_len);
	for (long long __i = 0; __i < names_len; __i++) names_list.append(QString::fromUtf8(names[__i]));
	((QTextCharFormat*)self)->setAnchorNames(names_list);
}
char** QTextCharFormat_anchorNames(void* self, long long* __out_len) {
	auto&& __r = ((QTextCharFormat*)self)->anchorNames();
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
void QTextCharFormat_setTableCellRowSpan(void* self, int tableCellRowSpan) {
	((QTextCharFormat*)self)->setTableCellRowSpan(tableCellRowSpan);
}
int QTextCharFormat_tableCellRowSpan(void* self) {
	auto&& __r = ((QTextCharFormat*)self)->tableCellRowSpan();
	return __r;
}
void QTextCharFormat_setTableCellColumnSpan(void* self, int tableCellColumnSpan) {
	((QTextCharFormat*)self)->setTableCellColumnSpan(tableCellColumnSpan);
}
int QTextCharFormat_tableCellColumnSpan(void* self) {
	auto&& __r = ((QTextCharFormat*)self)->tableCellColumnSpan();
	return __r;
}
void* QTextCharFormat_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QTextCharFormat*)self)->operator=(*(QTextCharFormat*)arg1);
	return (void*)&(__r);
}
void QTextCharFormat_setFont2(void* self, void* font, int behavior) {
	((QTextCharFormat*)self)->setFont(*(QFont*)font, (QTextCharFormat::FontPropertiesInheritanceBehavior)(behavior));
}
void QTextCharFormat_setFontStyleHint2(void* self, int hint, int strategy) {
	((QTextCharFormat*)self)->setFontStyleHint((QFont::StyleHint)(hint), (QFont::StyleStrategy)(strategy));
}
void QTextCharFormat_delete(void* self) { delete (QTextCharFormat*)self; }

}

#include <QtGui/qaccessible.h>
#include <QtCore/qvariant.h>
#include "gen_qaccessiblevalueinterface.h"
#include "gen_purego_cb.h"

class QAccessibleValueInterface__Go : public QAccessibleValueInterface {
public:
	long long __goId;
	unsigned char __goOn[5] = {};
	template<typename... A> QAccessibleValueInterface__Go(long long id, A&&... a) : QAccessibleValueInterface(std::forward<A>(a)...), __goId(id) {}
	QVariant currentValue() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 9131170551987814307LL, &__h); if (__h) return *(QVariant*)__r; }
		return QVariant{};
	}
	void setCurrentValue(const QVariant& value) override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3595292167320780357LL, (void*)&value, &__h); if (__h) return; }
		return;
	}
	QVariant maximumValue() const override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 3939360635140333998LL, &__h); if (__h) return *(QVariant*)__r; }
		return QVariant{};
	}
	QVariant minimumValue() const override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, -2855573697932810260LL, &__h); if (__h) return *(QVariant*)__r; }
		return QVariant{};
	}
	QVariant minimumStepSize() const override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, -373090913906449796LL, &__h); if (__h) return *(QVariant*)__r; }
		return QVariant{};
	}
};

extern "C" {

void QAccessibleValueInterface__set_override(void* self, int idx) { ((QAccessibleValueInterface__Go*)self)->__goOn[idx] = 1; }
void* QAccessibleValueInterface_currentValue(void* self) {
	auto&& __r = ((QAccessibleValueInterface*)self)->currentValue();
	return new QVariant(__r);
}
void QAccessibleValueInterface_setCurrentValue(void* self, void* value) {
	((QAccessibleValueInterface*)self)->setCurrentValue(*(QVariant*)value);
}
void* QAccessibleValueInterface_maximumValue(void* self) {
	auto&& __r = ((QAccessibleValueInterface*)self)->maximumValue();
	return new QVariant(__r);
}
void* QAccessibleValueInterface_minimumValue(void* self) {
	auto&& __r = ((QAccessibleValueInterface*)self)->minimumValue();
	return new QVariant(__r);
}
void* QAccessibleValueInterface_minimumStepSize(void* self) {
	auto&& __r = ((QAccessibleValueInterface*)self)->minimumStepSize();
	return new QVariant(__r);
}
void* QAccessibleValueInterface_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QAccessibleValueInterface*)self)->operator=(*(QAccessibleValueInterface*)arg1);
	return (void*)&(__r);
}
void* QAccessibleValueInterface__Go_new(long long id) { return new QAccessibleValueInterface__Go(id); }
void QAccessibleValueInterface_delete(void* self) { delete (QAccessibleValueInterface*)self; }

}

#include <QtGui/qactiongroup.h>
#include <QtCore/qobject.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtGui/qaction.h>
#include <QtGui/qicon.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qvariant.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qactiongroup.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QActionGroup__Go : public QActionGroup {
public:
	long long __goId;
	unsigned char __goOn[8] = {};
	template<typename... A> QActionGroup__Go(long long id, A&&... a) : QActionGroup(std::forward<A>(a)...), __goId(id) {}
	using QActionGroup::sender;
	using QActionGroup::senderSignalIndex;
	using QActionGroup::receivers;
	using QActionGroup::isSignalConnected;
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QActionGroup::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QActionGroup::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QActionGroup::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QActionGroup::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QActionGroup::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QActionGroup::metaObject();
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QActionGroup::event(event);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QActionGroup::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QActionGroup::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QActionGroup::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QActionGroup::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QActionGroup::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QActionGroup::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QActionGroup;
static std::mutex __connsMu_QActionGroup;

extern "C" {

void* QActionGroup_base(void* self) { return (void*)static_cast<QObject*>((QActionGroup*)self); }
void QActionGroup__set_override(void* self, int idx) { ((QActionGroup__Go*)self)->__goOn[idx] = 1; }
void* QActionGroup_fromQObject(void* o) { return (void*)qobject_cast<QActionGroup*>((QObject*)o); }
void* QActionGroup_new(long long __goid, void* parent) {
	return new QActionGroup__Go(__goid, (QObject*)parent);
}
char* QActionGroup_tr(const char* s, long long* __out_len) {
	auto&& __r = QActionGroup::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QActionGroup_addAction(void* self, void* a) {
	auto&& __r = ((QActionGroup*)self)->addAction((QAction*)a);
	return (void*)(__r);
}
void* QActionGroup_addAction2(void* self, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = ((QActionGroup*)self)->addAction(text_qstr);
	return (void*)(__r);
}
void* QActionGroup_addAction3(void* self, void* icon, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = ((QActionGroup*)self)->addAction(*(QIcon*)icon, text_qstr);
	return (void*)(__r);
}
void QActionGroup_removeAction(void* self, void* a) {
	((QActionGroup*)self)->removeAction((QAction*)a);
}
void** QActionGroup_actions(void* self, long long* __out_len) {
	auto&& __r = ((QActionGroup*)self)->actions();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
void* QActionGroup_checkedAction(void* self) {
	auto&& __r = ((QActionGroup*)self)->checkedAction();
	return (void*)(__r);
}
int QActionGroup_isExclusive(void* self) {
	auto&& __r = ((QActionGroup*)self)->isExclusive();
	return __r;
}
int QActionGroup_isEnabled(void* self) {
	auto&& __r = ((QActionGroup*)self)->isEnabled();
	return __r;
}
int QActionGroup_isVisible(void* self) {
	auto&& __r = ((QActionGroup*)self)->isVisible();
	return __r;
}
int QActionGroup_exclusionPolicy(void* self) {
	auto&& __r = ((QActionGroup*)self)->exclusionPolicy();
	return (int)(__r);
}
void QActionGroup_setEnabled(void* self, int arg1) {
	((QActionGroup*)self)->setEnabled(arg1);
}
void QActionGroup_setDisabled(void* self, int b) {
	((QActionGroup*)self)->setDisabled(b);
}
void QActionGroup_setVisible(void* self, int arg1) {
	((QActionGroup*)self)->setVisible(arg1);
}
void QActionGroup_setExclusive(void* self, int arg1) {
	((QActionGroup*)self)->setExclusive(arg1);
}
void QActionGroup_setExclusionPolicy(void* self, int policy) {
	((QActionGroup*)self)->setExclusionPolicy((QActionGroup::ExclusionPolicy)(policy));
}
void QActionGroup_triggered(void* self, void* arg1) {
	((QActionGroup*)self)->triggered((QAction*)arg1);
}
void QActionGroup_hovered(void* self, void* arg1) {
	((QActionGroup*)self)->hovered((QAction*)arg1);
}
void* QActionGroup_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QActionGroup::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QActionGroup_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QActionGroup::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QActionGroup_sender(void* self) {
	auto&& __r = ((QActionGroup__Go*)self)->sender();
	return (void*)(__r);
}
int QActionGroup_senderSignalIndex(void* self) {
	auto&& __r = ((QActionGroup__Go*)self)->senderSignalIndex();
	return __r;
}
int QActionGroup_receivers(void* self, const char* signal) {
	auto&& __r = ((QActionGroup__Go*)self)->receivers(signal);
	return __r;
}
int QActionGroup_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QActionGroup__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QActionGroup_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QActionGroup::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QActionGroup_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QActionGroup::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QActionGroup_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QActionGroup::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QActionGroup_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QActionGroup::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QActionGroup_metaObject_default(void* self) {
	auto&& __r = ((QActionGroup*)self)->QActionGroup::metaObject();
	return (void*)(__r);
}
int QActionGroup_event_default(void* self, void* event) {
	auto&& __r = ((QActionGroup*)self)->QActionGroup::event((QEvent*)event);
	return __r;
}
int QActionGroup_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QActionGroup*)self)->QActionGroup::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QActionGroup_timerEvent_default(void* self, void* event) {
	((QActionGroup__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QActionGroup_childEvent_default(void* self, void* event) {
	((QActionGroup__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QActionGroup_customEvent_default(void* self, void* event) {
	((QActionGroup__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QActionGroup_connectNotify_default(void* self, void* signal) {
	((QActionGroup__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QActionGroup_disconnectNotify_default(void* self, void* signal) {
	((QActionGroup__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QActionGroup_connect_triggered(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QActionGroup*)self, &QActionGroup::triggered, (QActionGroup*)self, [__id](QAction* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QActionGroup);
	__conns_QActionGroup[__id] = __c;
}
void QActionGroup_connect_hovered(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QActionGroup*)self, &QActionGroup::hovered, (QActionGroup*)self, [__id](QAction* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QActionGroup);
	__conns_QActionGroup[__id] = __c;
}
void QActionGroup_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QActionGroup*)self, &QActionGroup::destroyed, (QActionGroup*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QActionGroup);
	__conns_QActionGroup[__id] = __c;
}
void QActionGroup_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QActionGroup*)self, &QActionGroup::destroyed, (QActionGroup*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QActionGroup);
	__conns_QActionGroup[__id] = __c;
}
void QActionGroup_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QActionGroup*)self, &QActionGroup::objectNameChanged, (QActionGroup*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QActionGroup);
	__conns_QActionGroup[__id] = __c;
}
void QActionGroup__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QActionGroup);
		auto __it = __conns_QActionGroup.find(id);
		if (__it == __conns_QActionGroup.end()) return;
		__c = __it->second;
		__conns_QActionGroup.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QActionGroup__Go_new(long long id, void* parent) { return new QActionGroup__Go(id, (QObject*)parent); }
void QActionGroup_delete(void* self) { delete (QActionGroup*)self; }

}

#include <QtGui/qimageiohandler.h>
#include <QtCore/qobject.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qiodevice.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qvariant.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qimageioplugin.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QImageIOPlugin__Go : public QImageIOPlugin {
public:
	long long __goId;
	unsigned char __goOn[10] = {};
	template<typename... A> QImageIOPlugin__Go(long long id, A&&... a) : QImageIOPlugin(std::forward<A>(a)...), __goId(id) {}
	using QImageIOPlugin::sender;
	using QImageIOPlugin::senderSignalIndex;
	using QImageIOPlugin::receivers;
	using QImageIOPlugin::isSignalConnected;
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QImageIOPlugin::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QImageIOPlugin::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QImageIOPlugin::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QImageIOPlugin::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QImageIOPlugin::disconnectNotify(__a0); }
	QFlags<QImageIOPlugin::Capability> capabilities(QIODevice* device, const QByteArray& format) const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5572960902616693279LL, (void*)device, (void*)&format, &__h); if (__h) return QFlags<QImageIOPlugin::Capability>(__r); }
		return QFlags<QImageIOPlugin::Capability>{};
	}
	QImageIOHandler* create(QIODevice* device, const QByteArray& format) const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_voidp_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, void*, int*))cppgo_virt_voidp_voidp_voidp)(__goId, -626457835795418115LL, (void*)device, (void*)&format, &__h); if (__h) return (QImageIOHandler*)__r; }
		return nullptr;
	}
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QImageIOPlugin::metaObject();
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QImageIOPlugin::event(event);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QImageIOPlugin::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QImageIOPlugin::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QImageIOPlugin::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QImageIOPlugin::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QImageIOPlugin::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QImageIOPlugin::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QImageIOPlugin;
static std::mutex __connsMu_QImageIOPlugin;

extern "C" {

void* QImageIOPlugin_base(void* self) { return (void*)static_cast<QObject*>((QImageIOPlugin*)self); }
void QImageIOPlugin__set_override(void* self, int idx) { ((QImageIOPlugin__Go*)self)->__goOn[idx] = 1; }
void* QImageIOPlugin_fromQObject(void* o) { return (void*)qobject_cast<QImageIOPlugin*>((QObject*)o); }
char* QImageIOPlugin_tr(const char* s, long long* __out_len) {
	auto&& __r = QImageIOPlugin::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QImageIOPlugin_capabilities(void* self, void* device, void* format) {
	auto&& __r = ((QImageIOPlugin*)self)->capabilities((QIODevice*)device, *(QByteArray*)format);
	return (int)(__r);
}
void* QImageIOPlugin_create(void* self, void* device, void* format) {
	auto&& __r = ((QImageIOPlugin*)self)->create((QIODevice*)device, *(QByteArray*)format);
	return (void*)(__r);
}
void* QImageIOPlugin_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QImageIOPlugin::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QImageIOPlugin_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QImageIOPlugin::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QImageIOPlugin_sender(void* self) {
	auto&& __r = ((QImageIOPlugin__Go*)self)->sender();
	return (void*)(__r);
}
int QImageIOPlugin_senderSignalIndex(void* self) {
	auto&& __r = ((QImageIOPlugin__Go*)self)->senderSignalIndex();
	return __r;
}
int QImageIOPlugin_receivers(void* self, const char* signal) {
	auto&& __r = ((QImageIOPlugin__Go*)self)->receivers(signal);
	return __r;
}
int QImageIOPlugin_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QImageIOPlugin__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QImageIOPlugin_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QImageIOPlugin::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QImageIOPlugin_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QImageIOPlugin::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QImageIOPlugin_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QImageIOPlugin::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QImageIOPlugin_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QImageIOPlugin::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QImageIOPlugin_metaObject_default(void* self) {
	auto&& __r = ((QImageIOPlugin*)self)->QImageIOPlugin::metaObject();
	return (void*)(__r);
}
int QImageIOPlugin_event_default(void* self, void* event) {
	auto&& __r = ((QImageIOPlugin*)self)->QImageIOPlugin::event((QEvent*)event);
	return __r;
}
int QImageIOPlugin_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QImageIOPlugin*)self)->QImageIOPlugin::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QImageIOPlugin_timerEvent_default(void* self, void* event) {
	((QImageIOPlugin__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QImageIOPlugin_childEvent_default(void* self, void* event) {
	((QImageIOPlugin__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QImageIOPlugin_customEvent_default(void* self, void* event) {
	((QImageIOPlugin__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QImageIOPlugin_connectNotify_default(void* self, void* signal) {
	((QImageIOPlugin__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QImageIOPlugin_disconnectNotify_default(void* self, void* signal) {
	((QImageIOPlugin__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QImageIOPlugin_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QImageIOPlugin*)self, &QImageIOPlugin::destroyed, (QImageIOPlugin*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QImageIOPlugin);
	__conns_QImageIOPlugin[__id] = __c;
}
void QImageIOPlugin_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QImageIOPlugin*)self, &QImageIOPlugin::destroyed, (QImageIOPlugin*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QImageIOPlugin);
	__conns_QImageIOPlugin[__id] = __c;
}
void QImageIOPlugin_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QImageIOPlugin*)self, &QImageIOPlugin::objectNameChanged, (QImageIOPlugin*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QImageIOPlugin);
	__conns_QImageIOPlugin[__id] = __c;
}
void QImageIOPlugin__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QImageIOPlugin);
		auto __it = __conns_QImageIOPlugin.find(id);
		if (__it == __conns_QImageIOPlugin.end()) return;
		__c = __it->second;
		__conns_QImageIOPlugin.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QImageIOPlugin__Go_new(long long id, void* parent) { return new QImageIOPlugin__Go(id, (QObject*)parent); }
void QImageIOPlugin_delete(void* self) { delete (QImageIOPlugin*)self; }

}

#include <QtGui/qpainterstateguard.h>
#include <QtGui/qpainter.h>
#include "gen_qpainterstateguard.h"

extern "C" {

void* QPainterStateGuard_new(void* painter) {
	return new QPainterStateGuard((QPainter*)painter);
}
void* QPainterStateGuard_new2(void* painter, int state) {
	return new QPainterStateGuard((QPainter*)painter, (QPainterStateGuard::InitialState)(state));
}
void QPainterStateGuard_swap(void* self, void* other) {
	((QPainterStateGuard*)self)->swap(*(QPainterStateGuard*)other);
}
void QPainterStateGuard_save(void* self) {
	((QPainterStateGuard*)self)->save();
}
void QPainterStateGuard_restore(void* self) {
	((QPainterStateGuard*)self)->restore();
}
void QPainterStateGuard_delete(void* self) { delete (QPainterStateGuard*)self; }

}

#include <QtGui/qtextdocumentfragment.h>
#include <QtGui/qtextdocument.h>
#include <QtGui/qtextcursor.h>
#include <QtCore/qstring.h>
#include "gen_qtextdocumentfragment.h"

extern "C" {

void* QTextDocumentFragment_new() {
	return new QTextDocumentFragment();
}
void* QTextDocumentFragment_new2(void* document) {
	return new QTextDocumentFragment((QTextDocument*)document);
}
void* QTextDocumentFragment_new3(void* range_) {
	return new QTextDocumentFragment(*(QTextCursor*)range_);
}
void* QTextDocumentFragment_new4(void* rhs) {
	return new QTextDocumentFragment(*(QTextDocumentFragment*)rhs);
}
void* QTextDocumentFragment_OperatorAssign(void* self, void* rhs) {
	auto&& __r = ((QTextDocumentFragment*)self)->operator=(*(QTextDocumentFragment*)rhs);
	return (void*)&(__r);
}
int QTextDocumentFragment_isEmpty(void* self) {
	auto&& __r = ((QTextDocumentFragment*)self)->isEmpty();
	return __r;
}
char* QTextDocumentFragment_toPlainText(void* self, long long* __out_len) {
	auto&& __r = ((QTextDocumentFragment*)self)->toPlainText();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QTextDocumentFragment_toRawText(void* self, long long* __out_len) {
	auto&& __r = ((QTextDocumentFragment*)self)->toRawText();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QTextDocumentFragment_toHtml(void* self, long long* __out_len) {
	auto&& __r = ((QTextDocumentFragment*)self)->toHtml();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QTextDocumentFragment_toMarkdown(void* self, long long* __out_len) {
	auto&& __r = ((QTextDocumentFragment*)self)->toMarkdown();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QTextDocumentFragment_fromPlainText(const char* plainText, long long plainText_len) {
	QString plainText_qstr = QString::fromUtf8(plainText, plainText_len);
	auto&& __r = QTextDocumentFragment::fromPlainText(plainText_qstr);
	return new QTextDocumentFragment(__r);
}
void* QTextDocumentFragment_fromHtml(const char* html, long long html_len) {
	QString html_qstr = QString::fromUtf8(html, html_len);
	auto&& __r = QTextDocumentFragment::fromHtml(html_qstr);
	return new QTextDocumentFragment(__r);
}
void* QTextDocumentFragment_fromMarkdown(const char* markdown, long long markdown_len) {
	QString markdown_qstr = QString::fromUtf8(markdown, markdown_len);
	auto&& __r = QTextDocumentFragment::fromMarkdown(markdown_qstr);
	return new QTextDocumentFragment(__r);
}
char* QTextDocumentFragment_toMarkdown2(void* self, int features, long long* __out_len) {
	auto&& __r = ((QTextDocumentFragment*)self)->toMarkdown(QFlags<QTextDocument::MarkdownFeature>(features));
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QTextDocumentFragment_fromHtml2(const char* html, long long html_len, void* resourceProvider) {
	QString html_qstr = QString::fromUtf8(html, html_len);
	auto&& __r = QTextDocumentFragment::fromHtml(html_qstr, (QTextDocument*)resourceProvider);
	return new QTextDocumentFragment(__r);
}
void* QTextDocumentFragment_fromMarkdown2(const char* markdown, long long markdown_len, int features) {
	QString markdown_qstr = QString::fromUtf8(markdown, markdown_len);
	auto&& __r = QTextDocumentFragment::fromMarkdown(markdown_qstr, QFlags<QTextDocument::MarkdownFeature>(features));
	return new QTextDocumentFragment(__r);
}
void QTextDocumentFragment_delete(void* self) { delete (QTextDocumentFragment*)self; }

}

#include <QtWidgets/qframe.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qsize.h>
#include <QtCore/qrect.h>
#include <QtCore/qcoreevent.h>
#include <QtGui/qevent.h>
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
#include "gen_qframe.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QFrame__Go : public QFrame {
public:
	long long __goId;
	unsigned char __goOn[49] = {};
	template<typename... A> QFrame__Go(long long id, A&&... a) : QFrame(std::forward<A>(a)...), __goId(id) {}
	using QFrame::drawFrame;
	using QFrame::updateMicroFocus;
	using QFrame::create;
	using QFrame::destroy;
	using QFrame::focusNextChild;
	using QFrame::focusPreviousChild;
	using QFrame::sender;
	using QFrame::senderSignalIndex;
	using QFrame::receivers;
	using QFrame::isSignalConnected;
	using QFrame::getDecodedMetricF;
	bool cppgoBase_event(QEvent* __a0) { return QFrame::event(__a0); }
	void cppgoBase_paintEvent(QPaintEvent* __a0) { QFrame::paintEvent(__a0); }
	void cppgoBase_changeEvent(QEvent* __a0) { QFrame::changeEvent(__a0); }
	void cppgoBase_initStyleOption(QStyleOptionFrame* __a0) { QFrame::initStyleOption(__a0); }
	void cppgoBase_mousePressEvent(QMouseEvent* __a0) { QFrame::mousePressEvent(__a0); }
	void cppgoBase_mouseReleaseEvent(QMouseEvent* __a0) { QFrame::mouseReleaseEvent(__a0); }
	void cppgoBase_mouseDoubleClickEvent(QMouseEvent* __a0) { QFrame::mouseDoubleClickEvent(__a0); }
	void cppgoBase_mouseMoveEvent(QMouseEvent* __a0) { QFrame::mouseMoveEvent(__a0); }
	void cppgoBase_wheelEvent(QWheelEvent* __a0) { QFrame::wheelEvent(__a0); }
	void cppgoBase_keyPressEvent(QKeyEvent* __a0) { QFrame::keyPressEvent(__a0); }
	void cppgoBase_keyReleaseEvent(QKeyEvent* __a0) { QFrame::keyReleaseEvent(__a0); }
	void cppgoBase_focusInEvent(QFocusEvent* __a0) { QFrame::focusInEvent(__a0); }
	void cppgoBase_focusOutEvent(QFocusEvent* __a0) { QFrame::focusOutEvent(__a0); }
	void cppgoBase_enterEvent(QEnterEvent* __a0) { QFrame::enterEvent(__a0); }
	void cppgoBase_leaveEvent(QEvent* __a0) { QFrame::leaveEvent(__a0); }
	void cppgoBase_moveEvent(QMoveEvent* __a0) { QFrame::moveEvent(__a0); }
	void cppgoBase_resizeEvent(QResizeEvent* __a0) { QFrame::resizeEvent(__a0); }
	void cppgoBase_closeEvent(QCloseEvent* __a0) { QFrame::closeEvent(__a0); }
	void cppgoBase_contextMenuEvent(QContextMenuEvent* __a0) { QFrame::contextMenuEvent(__a0); }
	void cppgoBase_tabletEvent(QTabletEvent* __a0) { QFrame::tabletEvent(__a0); }
	void cppgoBase_actionEvent(QActionEvent* __a0) { QFrame::actionEvent(__a0); }
	void cppgoBase_dragEnterEvent(QDragEnterEvent* __a0) { QFrame::dragEnterEvent(__a0); }
	void cppgoBase_dragMoveEvent(QDragMoveEvent* __a0) { QFrame::dragMoveEvent(__a0); }
	void cppgoBase_dragLeaveEvent(QDragLeaveEvent* __a0) { QFrame::dragLeaveEvent(__a0); }
	void cppgoBase_dropEvent(QDropEvent* __a0) { QFrame::dropEvent(__a0); }
	void cppgoBase_showEvent(QShowEvent* __a0) { QFrame::showEvent(__a0); }
	void cppgoBase_hideEvent(QHideEvent* __a0) { QFrame::hideEvent(__a0); }
	int cppgoBase_metric(int __a0) { return QFrame::metric((QPaintDevice::PaintDeviceMetric)__a0); }
	void cppgoBase_initPainter(QPainter* __a0) { QFrame::initPainter(__a0); }
	QPaintDevice* cppgoBase_redirected(QPoint* __a0) { return QFrame::redirected(__a0); }
	QPainter* cppgoBase_sharedPainter() { return QFrame::sharedPainter(); }
	void cppgoBase_inputMethodEvent(QInputMethodEvent* __a0) { QFrame::inputMethodEvent(__a0); }
	bool cppgoBase_focusNextPrevChild(bool __a0) { return QFrame::focusNextPrevChild(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QFrame::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QFrame::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QFrame::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QFrame::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QFrame::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QFrame::metaObject();
	}
	QSize sizeHint() const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 729690261457684101LL, &__h); if (__h) return *(QSize*)__r; }
		return QFrame::sizeHint();
	}
	bool event(QEvent* e) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)e, &__h); if (__h) return __r; }
		return QFrame::event(e);
	}
	void paintEvent(QPaintEvent* arg1) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5290442163167458465LL, (void*)arg1, &__h); if (__h) return; }
		QFrame::paintEvent(arg1);
	}
	void changeEvent(QEvent* arg1) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4959376063826165943LL, (void*)arg1, &__h); if (__h) return; }
		QFrame::changeEvent(arg1);
	}
	void initStyleOption(QStyleOptionFrame* option) const override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -488693769771308531LL, (void*)option, &__h); if (__h) return; }
		QFrame::initStyleOption(option);
	}
	int devType() const override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8659679893868524044LL, &__h); if (__h) return __r; }
		return QFrame::devType();
	}
	void setVisible(bool visible) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 2110182719375245575LL, visible, &__h); if (__h) return; }
		QFrame::setVisible(visible);
	}
	QSize minimumSizeHint() const override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 7049264599837532491LL, &__h); if (__h) return *(QSize*)__r; }
		return QFrame::minimumSizeHint();
	}
	int heightForWidth(int arg1) const override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 4311534060495737339LL, arg1, &__h); if (__h) return __r; }
		return QFrame::heightForWidth(arg1);
	}
	bool hasHeightForWidth() const override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8266762397995423915LL, &__h); if (__h) return __r; }
		return QFrame::hasHeightForWidth();
	}
	QPaintEngine* paintEngine() const override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 6398007404756346381LL, &__h); if (__h) return (QPaintEngine*)__r; }
		return QFrame::paintEngine();
	}
	void mousePressEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -332090898078543319LL, (void*)event, &__h); if (__h) return; }
		QFrame::mousePressEvent(event);
	}
	void mouseReleaseEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4415920049297795213LL, (void*)event, &__h); if (__h) return; }
		QFrame::mouseReleaseEvent(event);
	}
	void mouseDoubleClickEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4029073905697023887LL, (void*)event, &__h); if (__h) return; }
		QFrame::mouseDoubleClickEvent(event);
	}
	void mouseMoveEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4255494004099806581LL, (void*)event, &__h); if (__h) return; }
		QFrame::mouseMoveEvent(event);
	}
	void wheelEvent(QWheelEvent* event) override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1919988340934511316LL, (void*)event, &__h); if (__h) return; }
		QFrame::wheelEvent(event);
	}
	void keyPressEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1023791977679793185LL, (void*)event, &__h); if (__h) return; }
		QFrame::keyPressEvent(event);
	}
	void keyReleaseEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 772598475950050299LL, (void*)event, &__h); if (__h) return; }
		QFrame::keyReleaseEvent(event);
	}
	void focusInEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 975043767033552552LL, (void*)event, &__h); if (__h) return; }
		QFrame::focusInEvent(event);
	}
	void focusOutEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -385596010576048993LL, (void*)event, &__h); if (__h) return; }
		QFrame::focusOutEvent(event);
	}
	void enterEvent(QEnterEvent* event) override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6316270346668237895LL, (void*)event, &__h); if (__h) return; }
		QFrame::enterEvent(event);
	}
	void leaveEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -551070009623711640LL, (void*)event, &__h); if (__h) return; }
		QFrame::leaveEvent(event);
	}
	void moveEvent(QMoveEvent* event) override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -9161949150489525396LL, (void*)event, &__h); if (__h) return; }
		QFrame::moveEvent(event);
	}
	void resizeEvent(QResizeEvent* event) override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4302917285655014661LL, (void*)event, &__h); if (__h) return; }
		QFrame::resizeEvent(event);
	}
	void closeEvent(QCloseEvent* event) override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1989937050019464783LL, (void*)event, &__h); if (__h) return; }
		QFrame::closeEvent(event);
	}
	void contextMenuEvent(QContextMenuEvent* event) override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6698043222227006015LL, (void*)event, &__h); if (__h) return; }
		QFrame::contextMenuEvent(event);
	}
	void tabletEvent(QTabletEvent* event) override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -212068866860701141LL, (void*)event, &__h); if (__h) return; }
		QFrame::tabletEvent(event);
	}
	void actionEvent(QActionEvent* event) override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1697444140088555581LL, (void*)event, &__h); if (__h) return; }
		QFrame::actionEvent(event);
	}
	void dragEnterEvent(QDragEnterEvent* event) override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 206122715919917589LL, (void*)event, &__h); if (__h) return; }
		QFrame::dragEnterEvent(event);
	}
	void dragMoveEvent(QDragMoveEvent* event) override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1781545137352745978LL, (void*)event, &__h); if (__h) return; }
		QFrame::dragMoveEvent(event);
	}
	void dragLeaveEvent(QDragLeaveEvent* event) override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -3139904158373725006LL, (void*)event, &__h); if (__h) return; }
		QFrame::dragLeaveEvent(event);
	}
	void dropEvent(QDropEvent* event) override {
		int __h = 0;
		if (__goOn[32] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 9107058682754591716LL, (void*)event, &__h); if (__h) return; }
		QFrame::dropEvent(event);
	}
	void showEvent(QShowEvent* event) override {
		int __h = 0;
		if (__goOn[33] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6825369174751916LL, (void*)event, &__h); if (__h) return; }
		QFrame::showEvent(event);
	}
	void hideEvent(QHideEvent* event) override {
		int __h = 0;
		if (__goOn[34] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -199718707876719129LL, (void*)event, &__h); if (__h) return; }
		QFrame::hideEvent(event);
	}
	bool nativeEvent(const QByteArray& eventType, void* message, qintptr* result) override {
		int __h = 0;
		if (__goOn[35] && cppgo_virt_int_voidp_voidp_longlongp) { int __r = ((int(*)(long long, long long, void*, void*, long long*, int*))cppgo_virt_int_voidp_voidp_longlongp)(__goId, -2039166586246443706LL, (void*)&eventType, (void*)message, (long long*)result, &__h); if (__h) return __r; }
		return QFrame::nativeEvent(eventType, message, result);
	}
	int metric(QPaintDevice::PaintDeviceMetric arg1) const override {
		int __h = 0;
		if (__goOn[36] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 427629649267311777LL, (int)(arg1), &__h); if (__h) return __r; }
		return QFrame::metric(arg1);
	}
	void initPainter(QPainter* painter) const override {
		int __h = 0;
		if (__goOn[37] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6477500296980050210LL, (void*)painter, &__h); if (__h) return; }
		QFrame::initPainter(painter);
	}
	QPaintDevice* redirected(QPoint* offset) const override {
		int __h = 0;
		if (__goOn[38] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 4127240633310655048LL, (void*)offset, &__h); if (__h) return (QPaintDevice*)__r; }
		return QFrame::redirected(offset);
	}
	QPainter* sharedPainter() const override {
		int __h = 0;
		if (__goOn[39] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 5792771328741774991LL, &__h); if (__h) return (QPainter*)__r; }
		return QFrame::sharedPainter();
	}
	void inputMethodEvent(QInputMethodEvent* arg1) override {
		int __h = 0;
		if (__goOn[40] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5438819661942843222LL, (void*)arg1, &__h); if (__h) return; }
		QFrame::inputMethodEvent(arg1);
	}
	QVariant inputMethodQuery(Qt::InputMethodQuery arg1) const override {
		int __h = 0;
		if (__goOn[41] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7345853097340709638LL, (int)(arg1), &__h); if (__h) return *(QVariant*)__r; }
		return QFrame::inputMethodQuery(arg1);
	}
	bool focusNextPrevChild(bool next) override {
		int __h = 0;
		if (__goOn[42] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 5898328517505784191LL, next, &__h); if (__h) return __r; }
		return QFrame::focusNextPrevChild(next);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[43] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QFrame::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[44] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QFrame::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[45] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QFrame::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[46] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QFrame::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[47] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QFrame::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[48] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QFrame::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QFrame;
static std::mutex __connsMu_QFrame;

extern "C" {

void* QFrame_base(void* self) { return (void*)static_cast<QWidget*>((QFrame*)self); }
void QFrame_chain(void* self, void** out) {
	QFrame* __p = (QFrame*)self;
	QWidget* __b0 = static_cast<QWidget*>(__p); out[0] = (void*)__b0;
	QObject* __b1 = static_cast<QObject*>(__b0); out[1] = (void*)__b1;
}
void QFrame__set_override(void* self, int idx) { ((QFrame__Go*)self)->__goOn[idx] = 1; }
void* QFrame_fromQObject(void* o) { return (void*)qobject_cast<QFrame*>((QObject*)o); }
void* QFrame_new(long long __goid, void* parent) {
	return new QFrame__Go(__goid, (QWidget*)parent);
}
void* QFrame_new2(long long __goid) {
	return new QFrame__Go(__goid);
}
void* QFrame_new3(long long __goid, void* parent, int f) {
	return new QFrame__Go(__goid, (QWidget*)parent, QFlags<Qt::WindowType>(f));
}
char* QFrame_tr(const char* s, long long* __out_len) {
	auto&& __r = QFrame::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QFrame_frameStyle(void* self) {
	auto&& __r = ((QFrame*)self)->frameStyle();
	return __r;
}
void QFrame_setFrameStyle(void* self, int arg1) {
	((QFrame*)self)->setFrameStyle(arg1);
}
int QFrame_frameWidth(void* self) {
	auto&& __r = ((QFrame*)self)->frameWidth();
	return __r;
}
int QFrame_frameShape(void* self) {
	auto&& __r = ((QFrame*)self)->frameShape();
	return (int)(__r);
}
void QFrame_setFrameShape(void* self, int arg1) {
	((QFrame*)self)->setFrameShape((QFrame::Shape)(arg1));
}
int QFrame_frameShadow(void* self) {
	auto&& __r = ((QFrame*)self)->frameShadow();
	return (int)(__r);
}
void QFrame_setFrameShadow(void* self, int arg1) {
	((QFrame*)self)->setFrameShadow((QFrame::Shadow)(arg1));
}
int QFrame_lineWidth(void* self) {
	auto&& __r = ((QFrame*)self)->lineWidth();
	return __r;
}
void QFrame_setLineWidth(void* self, int arg1) {
	((QFrame*)self)->setLineWidth(arg1);
}
int QFrame_midLineWidth(void* self) {
	auto&& __r = ((QFrame*)self)->midLineWidth();
	return __r;
}
void QFrame_setMidLineWidth(void* self, int arg1) {
	((QFrame*)self)->setMidLineWidth(arg1);
}
void* QFrame_frameRect(void* self) {
	auto&& __r = ((QFrame*)self)->frameRect();
	return new QRect(__r);
}
void QFrame_setFrameRect(void* self, void* arg1) {
	((QFrame*)self)->setFrameRect(*(QRect*)arg1);
}
void QFrame_drawFrame(void* self, void* arg1) {
	((QFrame__Go*)self)->drawFrame((QPainter*)arg1);
}
void QFrame_setTabOrder(void* arg1, void* arg2) {
	QFrame::setTabOrder((QWidget*)arg1, (QWidget*)arg2);
}
void* QFrame_mouseGrabber() {
	auto&& __r = QFrame::mouseGrabber();
	return (void*)(__r);
}
void* QFrame_keyboardGrabber() {
	auto&& __r = QFrame::keyboardGrabber();
	return (void*)(__r);
}
void* QFrame_find(unsigned long long arg1) {
	auto&& __r = QFrame::find(arg1);
	return (void*)(__r);
}
void* QFrame_createWindowContainer(void* window) {
	auto&& __r = QFrame::createWindowContainer((QWindow*)window);
	return (void*)(__r);
}
void QFrame_updateMicroFocus(void* self) {
	((QFrame__Go*)self)->updateMicroFocus();
}
void QFrame_create(void* self) {
	((QFrame__Go*)self)->create();
}
void QFrame_destroy(void* self) {
	((QFrame__Go*)self)->destroy();
}
int QFrame_focusNextChild(void* self) {
	auto&& __r = ((QFrame__Go*)self)->focusNextChild();
	return __r;
}
int QFrame_focusPreviousChild(void* self) {
	auto&& __r = ((QFrame__Go*)self)->focusPreviousChild();
	return __r;
}
void* QFrame_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QFrame::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QFrame_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QFrame::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QFrame_sender(void* self) {
	auto&& __r = ((QFrame__Go*)self)->sender();
	return (void*)(__r);
}
int QFrame_senderSignalIndex(void* self) {
	auto&& __r = ((QFrame__Go*)self)->senderSignalIndex();
	return __r;
}
int QFrame_receivers(void* self, const char* signal) {
	auto&& __r = ((QFrame__Go*)self)->receivers(signal);
	return __r;
}
int QFrame_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QFrame__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
double QFrame_devicePixelRatioFScale() {
	auto&& __r = QFrame::devicePixelRatioFScale();
	return __r;
}
int QFrame_encodeMetricF(int metric, double value) {
	auto&& __r = QFrame::encodeMetricF((QPaintDevice::PaintDeviceMetric)(metric), value);
	return __r;
}
double QFrame_getDecodedMetricF(void* self, int metricA, int metricB) {
	auto&& __r = ((QFrame__Go*)self)->getDecodedMetricF((QPaintDevice::PaintDeviceMetric)(metricA), (QPaintDevice::PaintDeviceMetric)(metricB));
	return __r;
}
char* QFrame_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QFrame::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QFrame_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QFrame::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QFrame_createWindowContainer2(void* window, void* parent) {
	auto&& __r = QFrame::createWindowContainer((QWindow*)window, (QWidget*)parent);
	return (void*)(__r);
}
void* QFrame_createWindowContainer3(void* window, void* parent, int flags) {
	auto&& __r = QFrame::createWindowContainer((QWindow*)window, (QWidget*)parent, QFlags<Qt::WindowType>(flags));
	return (void*)(__r);
}
void QFrame_updateMicroFocus2(void* self, int query) {
	((QFrame__Go*)self)->updateMicroFocus((Qt::InputMethodQuery)(query));
}
void QFrame_create2(void* self, unsigned long long arg1) {
	((QFrame__Go*)self)->create(arg1);
}
void QFrame_create3(void* self, unsigned long long arg1, int initializeWindow) {
	((QFrame__Go*)self)->create(arg1, initializeWindow);
}
void QFrame_create4(void* self, unsigned long long arg1, int initializeWindow, int destroyOldWindow) {
	((QFrame__Go*)self)->create(arg1, initializeWindow, destroyOldWindow);
}
void QFrame_destroy2(void* self, int destroyWindow) {
	((QFrame__Go*)self)->destroy(destroyWindow);
}
void QFrame_destroy3(void* self, int destroyWindow, int destroySubWindows) {
	((QFrame__Go*)self)->destroy(destroyWindow, destroySubWindows);
}
void* QFrame_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QFrame::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QFrame_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QFrame::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QFrame_metaObject_default(void* self) {
	auto&& __r = ((QFrame*)self)->QFrame::metaObject();
	return (void*)(__r);
}
void* QFrame_sizeHint_default(void* self) {
	auto&& __r = ((QFrame*)self)->QFrame::sizeHint();
	return new QSize(__r);
}
int QFrame_event_default(void* self, void* e) {
	auto&& __r = ((QFrame__Go*)self)->cppgoBase_event((QEvent*)e);
	return __r;
}
void QFrame_paintEvent_default(void* self, void* arg1) {
	((QFrame__Go*)self)->cppgoBase_paintEvent((QPaintEvent*)arg1);
}
void QFrame_changeEvent_default(void* self, void* arg1) {
	((QFrame__Go*)self)->cppgoBase_changeEvent((QEvent*)arg1);
}
void QFrame_initStyleOption_default(void* self, void* option) {
	((QFrame__Go*)self)->cppgoBase_initStyleOption((QStyleOptionFrame*)option);
}
int QFrame_devType_default(void* self) {
	auto&& __r = ((QFrame*)self)->QFrame::devType();
	return __r;
}
void QFrame_setVisible_default(void* self, int visible) {
	((QFrame*)self)->QFrame::setVisible(visible);
}
void* QFrame_minimumSizeHint_default(void* self) {
	auto&& __r = ((QFrame*)self)->QFrame::minimumSizeHint();
	return new QSize(__r);
}
int QFrame_heightForWidth_default(void* self, int arg1) {
	auto&& __r = ((QFrame*)self)->QFrame::heightForWidth(arg1);
	return __r;
}
int QFrame_hasHeightForWidth_default(void* self) {
	auto&& __r = ((QFrame*)self)->QFrame::hasHeightForWidth();
	return __r;
}
void* QFrame_paintEngine_default(void* self) {
	auto&& __r = ((QFrame*)self)->QFrame::paintEngine();
	return (void*)(__r);
}
void QFrame_mousePressEvent_default(void* self, void* event) {
	((QFrame__Go*)self)->cppgoBase_mousePressEvent((QMouseEvent*)event);
}
void QFrame_mouseReleaseEvent_default(void* self, void* event) {
	((QFrame__Go*)self)->cppgoBase_mouseReleaseEvent((QMouseEvent*)event);
}
void QFrame_mouseDoubleClickEvent_default(void* self, void* event) {
	((QFrame__Go*)self)->cppgoBase_mouseDoubleClickEvent((QMouseEvent*)event);
}
void QFrame_mouseMoveEvent_default(void* self, void* event) {
	((QFrame__Go*)self)->cppgoBase_mouseMoveEvent((QMouseEvent*)event);
}
void QFrame_wheelEvent_default(void* self, void* event) {
	((QFrame__Go*)self)->cppgoBase_wheelEvent((QWheelEvent*)event);
}
void QFrame_keyPressEvent_default(void* self, void* event) {
	((QFrame__Go*)self)->cppgoBase_keyPressEvent((QKeyEvent*)event);
}
void QFrame_keyReleaseEvent_default(void* self, void* event) {
	((QFrame__Go*)self)->cppgoBase_keyReleaseEvent((QKeyEvent*)event);
}
void QFrame_focusInEvent_default(void* self, void* event) {
	((QFrame__Go*)self)->cppgoBase_focusInEvent((QFocusEvent*)event);
}
void QFrame_focusOutEvent_default(void* self, void* event) {
	((QFrame__Go*)self)->cppgoBase_focusOutEvent((QFocusEvent*)event);
}
void QFrame_enterEvent_default(void* self, void* event) {
	((QFrame__Go*)self)->cppgoBase_enterEvent((QEnterEvent*)event);
}
void QFrame_leaveEvent_default(void* self, void* event) {
	((QFrame__Go*)self)->cppgoBase_leaveEvent((QEvent*)event);
}
void QFrame_moveEvent_default(void* self, void* event) {
	((QFrame__Go*)self)->cppgoBase_moveEvent((QMoveEvent*)event);
}
void QFrame_resizeEvent_default(void* self, void* event) {
	((QFrame__Go*)self)->cppgoBase_resizeEvent((QResizeEvent*)event);
}
void QFrame_closeEvent_default(void* self, void* event) {
	((QFrame__Go*)self)->cppgoBase_closeEvent((QCloseEvent*)event);
}
void QFrame_contextMenuEvent_default(void* self, void* event) {
	((QFrame__Go*)self)->cppgoBase_contextMenuEvent((QContextMenuEvent*)event);
}
void QFrame_tabletEvent_default(void* self, void* event) {
	((QFrame__Go*)self)->cppgoBase_tabletEvent((QTabletEvent*)event);
}
void QFrame_actionEvent_default(void* self, void* event) {
	((QFrame__Go*)self)->cppgoBase_actionEvent((QActionEvent*)event);
}
void QFrame_dragEnterEvent_default(void* self, void* event) {
	((QFrame__Go*)self)->cppgoBase_dragEnterEvent((QDragEnterEvent*)event);
}
void QFrame_dragMoveEvent_default(void* self, void* event) {
	((QFrame__Go*)self)->cppgoBase_dragMoveEvent((QDragMoveEvent*)event);
}
void QFrame_dragLeaveEvent_default(void* self, void* event) {
	((QFrame__Go*)self)->cppgoBase_dragLeaveEvent((QDragLeaveEvent*)event);
}
void QFrame_dropEvent_default(void* self, void* event) {
	((QFrame__Go*)self)->cppgoBase_dropEvent((QDropEvent*)event);
}
void QFrame_showEvent_default(void* self, void* event) {
	((QFrame__Go*)self)->cppgoBase_showEvent((QShowEvent*)event);
}
void QFrame_hideEvent_default(void* self, void* event) {
	((QFrame__Go*)self)->cppgoBase_hideEvent((QHideEvent*)event);
}
int QFrame_metric_default(void* self, int arg1) {
	auto&& __r = ((QFrame__Go*)self)->cppgoBase_metric(arg1);
	return __r;
}
void QFrame_initPainter_default(void* self, void* painter) {
	((QFrame__Go*)self)->cppgoBase_initPainter((QPainter*)painter);
}
void* QFrame_redirected_default(void* self, void* offset) {
	auto&& __r = ((QFrame__Go*)self)->cppgoBase_redirected((QPoint*)offset);
	return (void*)(__r);
}
void* QFrame_sharedPainter_default(void* self) {
	auto&& __r = ((QFrame__Go*)self)->cppgoBase_sharedPainter();
	return (void*)(__r);
}
void QFrame_inputMethodEvent_default(void* self, void* arg1) {
	((QFrame__Go*)self)->cppgoBase_inputMethodEvent((QInputMethodEvent*)arg1);
}
void* QFrame_inputMethodQuery_default(void* self, int arg1) {
	auto&& __r = ((QFrame*)self)->QFrame::inputMethodQuery((Qt::InputMethodQuery)(arg1));
	return new QVariant(__r);
}
int QFrame_focusNextPrevChild_default(void* self, int next) {
	auto&& __r = ((QFrame__Go*)self)->cppgoBase_focusNextPrevChild(next);
	return __r;
}
int QFrame_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QFrame*)self)->QFrame::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QFrame_timerEvent_default(void* self, void* event) {
	((QFrame__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QFrame_childEvent_default(void* self, void* event) {
	((QFrame__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QFrame_customEvent_default(void* self, void* event) {
	((QFrame__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QFrame_connectNotify_default(void* self, void* signal) {
	((QFrame__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QFrame_disconnectNotify_default(void* self, void* signal) {
	((QFrame__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QFrame_connect_windowTitleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFrame*)self, &QFrame::windowTitleChanged, (QFrame*)self, [__id](const QString& title){
		QByteArray title_ba = title.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)title_ba.constData(), title_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFrame);
	__conns_QFrame[__id] = __c;
}
void QFrame_connect_windowIconChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFrame*)self, &QFrame::windowIconChanged, (QFrame*)self, [__id](const QIcon& icon){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&icon);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFrame);
	__conns_QFrame[__id] = __c;
}
void QFrame_connect_windowIconTextChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFrame*)self, &QFrame::windowIconTextChanged, (QFrame*)self, [__id](const QString& iconText){
		QByteArray iconText_ba = iconText.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)iconText_ba.constData(), iconText_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFrame);
	__conns_QFrame[__id] = __c;
}
void QFrame_connect_customContextMenuRequested(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFrame*)self, &QFrame::customContextMenuRequested, (QFrame*)self, [__id](const QPoint& pos){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&pos);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFrame);
	__conns_QFrame[__id] = __c;
}
void QFrame_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFrame*)self, &QFrame::destroyed, (QFrame*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFrame);
	__conns_QFrame[__id] = __c;
}
void QFrame_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFrame*)self, &QFrame::destroyed, (QFrame*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFrame);
	__conns_QFrame[__id] = __c;
}
void QFrame_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFrame*)self, &QFrame::objectNameChanged, (QFrame*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFrame);
	__conns_QFrame[__id] = __c;
}
void QFrame__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QFrame);
		auto __it = __conns_QFrame.find(id);
		if (__it == __conns_QFrame.end()) return;
		__c = __it->second;
		__conns_QFrame.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QFrame__Go_new(long long id) { return new QFrame__Go(id); }
void QFrame_delete(void* self) { delete (QFrame*)self; }

}

#include <QtWidgets/qstyleoption.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qstring.h>
#include <QtCore/qnamespace.h>
#include <QtGui/qcolor.h>
#include "gen_qstyleoptiongroupbox.h"

extern "C" {

void* QStyleOptionGroupBox_base(void* self) { return (void*)static_cast<QStyleOptionComplex*>((QStyleOptionGroupBox*)self); }
void QStyleOptionGroupBox_chain(void* self, void** out) {
	QStyleOptionGroupBox* __p = (QStyleOptionGroupBox*)self;
	QStyleOptionComplex* __b0 = static_cast<QStyleOptionComplex*>(__p); out[0] = (void*)__b0;
	QStyleOption* __b1 = static_cast<QStyleOption*>(__b0); out[1] = (void*)__b1;
}
void* QStyleOptionGroupBox_new() {
	return new QStyleOptionGroupBox();
}
void* QStyleOptionGroupBox_new2(void* other) {
	return new QStyleOptionGroupBox(*(QStyleOptionGroupBox*)other);
}
void* QStyleOptionGroupBox_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QStyleOptionGroupBox*)self)->operator=(*(QStyleOptionGroupBox*)arg1);
	return (void*)&(__r);
}
int QStyleOptionGroupBox_features_get(void* self) {
	auto __r = ((QStyleOptionGroupBox*)self)->features;
	return (int)(__r);
}
void QStyleOptionGroupBox_features_set(void* self, int features) {
	((QStyleOptionGroupBox*)self)->features = QFlags<QStyleOptionFrame::FrameFeature>(features);
}
char* QStyleOptionGroupBox_text_get(void* self, long long* __out_len) {
	auto __r = ((QStyleOptionGroupBox*)self)->text;
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QStyleOptionGroupBox_text_set(void* self, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	((QStyleOptionGroupBox*)self)->text = text_qstr;
}
int QStyleOptionGroupBox_textAlignment_get(void* self) {
	auto __r = ((QStyleOptionGroupBox*)self)->textAlignment;
	return (int)(__r);
}
void QStyleOptionGroupBox_textAlignment_set(void* self, int textAlignment) {
	((QStyleOptionGroupBox*)self)->textAlignment = QFlags<Qt::AlignmentFlag>(textAlignment);
}
void* QStyleOptionGroupBox_textColor_get(void* self) {
	auto __r = ((QStyleOptionGroupBox*)self)->textColor;
	return new QColor(__r);
}
void QStyleOptionGroupBox_textColor_set(void* self, void* textColor) {
	((QStyleOptionGroupBox*)self)->textColor = *(QColor*)textColor;
}
int QStyleOptionGroupBox_lineWidth_get(void* self) {
	auto __r = ((QStyleOptionGroupBox*)self)->lineWidth;
	return __r;
}
void QStyleOptionGroupBox_lineWidth_set(void* self, int lineWidth) {
	((QStyleOptionGroupBox*)self)->lineWidth = lineWidth;
}
int QStyleOptionGroupBox_midLineWidth_get(void* self) {
	auto __r = ((QStyleOptionGroupBox*)self)->midLineWidth;
	return __r;
}
void QStyleOptionGroupBox_midLineWidth_set(void* self, int midLineWidth) {
	((QStyleOptionGroupBox*)self)->midLineWidth = midLineWidth;
}
void QStyleOptionGroupBox_delete(void* self) { delete (QStyleOptionGroupBox*)self; }

}

#include <QtWidgets/qdialog.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qsize.h>
#include <QtGui/qevent.h>
#include <QtCore/qobject.h>
#include <QtCore/qcoreevent.h>
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
#include <QtGui/qpaintengine.h>
#include <QtGui/qbackingstore.h>
#include <QtGui/qwindow.h>
#include <QtGui/qscreen.h>
#include <QtCore/qvariant.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qdialog.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QDialog__Go : public QDialog {
public:
	long long __goId;
	unsigned char __goOn[53] = {};
	template<typename... A> QDialog__Go(long long id, A&&... a) : QDialog(std::forward<A>(a)...), __goId(id) {}
	using QDialog::adjustPosition;
	using QDialog::updateMicroFocus;
	using QDialog::create;
	using QDialog::destroy;
	using QDialog::focusNextChild;
	using QDialog::focusPreviousChild;
	using QDialog::sender;
	using QDialog::senderSignalIndex;
	using QDialog::receivers;
	using QDialog::isSignalConnected;
	using QDialog::getDecodedMetricF;
	void cppgoBase_keyPressEvent(QKeyEvent* __a0) { QDialog::keyPressEvent(__a0); }
	void cppgoBase_closeEvent(QCloseEvent* __a0) { QDialog::closeEvent(__a0); }
	void cppgoBase_showEvent(QShowEvent* __a0) { QDialog::showEvent(__a0); }
	void cppgoBase_resizeEvent(QResizeEvent* __a0) { QDialog::resizeEvent(__a0); }
	void cppgoBase_contextMenuEvent(QContextMenuEvent* __a0) { QDialog::contextMenuEvent(__a0); }
	bool cppgoBase_eventFilter(QObject* __a0, QEvent* __a1) { return QDialog::eventFilter(__a0, __a1); }
	bool cppgoBase_event(QEvent* __a0) { return QDialog::event(__a0); }
	void cppgoBase_mousePressEvent(QMouseEvent* __a0) { QDialog::mousePressEvent(__a0); }
	void cppgoBase_mouseReleaseEvent(QMouseEvent* __a0) { QDialog::mouseReleaseEvent(__a0); }
	void cppgoBase_mouseDoubleClickEvent(QMouseEvent* __a0) { QDialog::mouseDoubleClickEvent(__a0); }
	void cppgoBase_mouseMoveEvent(QMouseEvent* __a0) { QDialog::mouseMoveEvent(__a0); }
	void cppgoBase_wheelEvent(QWheelEvent* __a0) { QDialog::wheelEvent(__a0); }
	void cppgoBase_keyReleaseEvent(QKeyEvent* __a0) { QDialog::keyReleaseEvent(__a0); }
	void cppgoBase_focusInEvent(QFocusEvent* __a0) { QDialog::focusInEvent(__a0); }
	void cppgoBase_focusOutEvent(QFocusEvent* __a0) { QDialog::focusOutEvent(__a0); }
	void cppgoBase_enterEvent(QEnterEvent* __a0) { QDialog::enterEvent(__a0); }
	void cppgoBase_leaveEvent(QEvent* __a0) { QDialog::leaveEvent(__a0); }
	void cppgoBase_paintEvent(QPaintEvent* __a0) { QDialog::paintEvent(__a0); }
	void cppgoBase_moveEvent(QMoveEvent* __a0) { QDialog::moveEvent(__a0); }
	void cppgoBase_tabletEvent(QTabletEvent* __a0) { QDialog::tabletEvent(__a0); }
	void cppgoBase_actionEvent(QActionEvent* __a0) { QDialog::actionEvent(__a0); }
	void cppgoBase_dragEnterEvent(QDragEnterEvent* __a0) { QDialog::dragEnterEvent(__a0); }
	void cppgoBase_dragMoveEvent(QDragMoveEvent* __a0) { QDialog::dragMoveEvent(__a0); }
	void cppgoBase_dragLeaveEvent(QDragLeaveEvent* __a0) { QDialog::dragLeaveEvent(__a0); }
	void cppgoBase_dropEvent(QDropEvent* __a0) { QDialog::dropEvent(__a0); }
	void cppgoBase_hideEvent(QHideEvent* __a0) { QDialog::hideEvent(__a0); }
	void cppgoBase_changeEvent(QEvent* __a0) { QDialog::changeEvent(__a0); }
	int cppgoBase_metric(int __a0) { return QDialog::metric((QPaintDevice::PaintDeviceMetric)__a0); }
	void cppgoBase_initPainter(QPainter* __a0) { QDialog::initPainter(__a0); }
	QPaintDevice* cppgoBase_redirected(QPoint* __a0) { return QDialog::redirected(__a0); }
	QPainter* cppgoBase_sharedPainter() { return QDialog::sharedPainter(); }
	void cppgoBase_inputMethodEvent(QInputMethodEvent* __a0) { QDialog::inputMethodEvent(__a0); }
	bool cppgoBase_focusNextPrevChild(bool __a0) { return QDialog::focusNextPrevChild(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QDialog::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QDialog::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QDialog::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QDialog::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QDialog::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QDialog::metaObject();
	}
	void setVisible(bool visible) override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 2110182719375245575LL, visible, &__h); if (__h) return; }
		QDialog::setVisible(visible);
	}
	QSize sizeHint() const override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 729690261457684101LL, &__h); if (__h) return *(QSize*)__r; }
		return QDialog::sizeHint();
	}
	QSize minimumSizeHint() const override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 7049264599837532491LL, &__h); if (__h) return *(QSize*)__r; }
		return QDialog::minimumSizeHint();
	}
	void open() override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, -554057428128861975LL, &__h); if (__h) return; }
		QDialog::open();
	}
	int exec() override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -1745322561779819684LL, &__h); if (__h) return __r; }
		return QDialog::exec();
	}
	void done(int arg1) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, -2570997490807640431LL, arg1, &__h); if (__h) return; }
		QDialog::done(arg1);
	}
	void accept() override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 9106131928185084681LL, &__h); if (__h) return; }
		QDialog::accept();
	}
	void reject() override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 2875220963396170150LL, &__h); if (__h) return; }
		QDialog::reject();
	}
	void keyPressEvent(QKeyEvent* arg1) override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1023791977679793185LL, (void*)arg1, &__h); if (__h) return; }
		QDialog::keyPressEvent(arg1);
	}
	void closeEvent(QCloseEvent* arg1) override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1989937050019464783LL, (void*)arg1, &__h); if (__h) return; }
		QDialog::closeEvent(arg1);
	}
	void showEvent(QShowEvent* arg1) override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6825369174751916LL, (void*)arg1, &__h); if (__h) return; }
		QDialog::showEvent(arg1);
	}
	void resizeEvent(QResizeEvent* arg1) override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4302917285655014661LL, (void*)arg1, &__h); if (__h) return; }
		QDialog::resizeEvent(arg1);
	}
	void contextMenuEvent(QContextMenuEvent* arg1) override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6698043222227006015LL, (void*)arg1, &__h); if (__h) return; }
		QDialog::contextMenuEvent(arg1);
	}
	bool eventFilter(QObject* arg1, QEvent* arg2) override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)arg1, (void*)arg2, &__h); if (__h) return __r; }
		return QDialog::eventFilter(arg1, arg2);
	}
	int devType() const override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8659679893868524044LL, &__h); if (__h) return __r; }
		return QDialog::devType();
	}
	int heightForWidth(int arg1) const override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 4311534060495737339LL, arg1, &__h); if (__h) return __r; }
		return QDialog::heightForWidth(arg1);
	}
	bool hasHeightForWidth() const override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8266762397995423915LL, &__h); if (__h) return __r; }
		return QDialog::hasHeightForWidth();
	}
	QPaintEngine* paintEngine() const override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 6398007404756346381LL, &__h); if (__h) return (QPaintEngine*)__r; }
		return QDialog::paintEngine();
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QDialog::event(event);
	}
	void mousePressEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -332090898078543319LL, (void*)event, &__h); if (__h) return; }
		QDialog::mousePressEvent(event);
	}
	void mouseReleaseEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4415920049297795213LL, (void*)event, &__h); if (__h) return; }
		QDialog::mouseReleaseEvent(event);
	}
	void mouseDoubleClickEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4029073905697023887LL, (void*)event, &__h); if (__h) return; }
		QDialog::mouseDoubleClickEvent(event);
	}
	void mouseMoveEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4255494004099806581LL, (void*)event, &__h); if (__h) return; }
		QDialog::mouseMoveEvent(event);
	}
	void wheelEvent(QWheelEvent* event) override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1919988340934511316LL, (void*)event, &__h); if (__h) return; }
		QDialog::wheelEvent(event);
	}
	void keyReleaseEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 772598475950050299LL, (void*)event, &__h); if (__h) return; }
		QDialog::keyReleaseEvent(event);
	}
	void focusInEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 975043767033552552LL, (void*)event, &__h); if (__h) return; }
		QDialog::focusInEvent(event);
	}
	void focusOutEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -385596010576048993LL, (void*)event, &__h); if (__h) return; }
		QDialog::focusOutEvent(event);
	}
	void enterEvent(QEnterEvent* event) override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6316270346668237895LL, (void*)event, &__h); if (__h) return; }
		QDialog::enterEvent(event);
	}
	void leaveEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -551070009623711640LL, (void*)event, &__h); if (__h) return; }
		QDialog::leaveEvent(event);
	}
	void paintEvent(QPaintEvent* event) override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5290442163167458465LL, (void*)event, &__h); if (__h) return; }
		QDialog::paintEvent(event);
	}
	void moveEvent(QMoveEvent* event) override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -9161949150489525396LL, (void*)event, &__h); if (__h) return; }
		QDialog::moveEvent(event);
	}
	void tabletEvent(QTabletEvent* event) override {
		int __h = 0;
		if (__goOn[32] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -212068866860701141LL, (void*)event, &__h); if (__h) return; }
		QDialog::tabletEvent(event);
	}
	void actionEvent(QActionEvent* event) override {
		int __h = 0;
		if (__goOn[33] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1697444140088555581LL, (void*)event, &__h); if (__h) return; }
		QDialog::actionEvent(event);
	}
	void dragEnterEvent(QDragEnterEvent* event) override {
		int __h = 0;
		if (__goOn[34] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 206122715919917589LL, (void*)event, &__h); if (__h) return; }
		QDialog::dragEnterEvent(event);
	}
	void dragMoveEvent(QDragMoveEvent* event) override {
		int __h = 0;
		if (__goOn[35] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1781545137352745978LL, (void*)event, &__h); if (__h) return; }
		QDialog::dragMoveEvent(event);
	}
	void dragLeaveEvent(QDragLeaveEvent* event) override {
		int __h = 0;
		if (__goOn[36] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -3139904158373725006LL, (void*)event, &__h); if (__h) return; }
		QDialog::dragLeaveEvent(event);
	}
	void dropEvent(QDropEvent* event) override {
		int __h = 0;
		if (__goOn[37] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 9107058682754591716LL, (void*)event, &__h); if (__h) return; }
		QDialog::dropEvent(event);
	}
	void hideEvent(QHideEvent* event) override {
		int __h = 0;
		if (__goOn[38] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -199718707876719129LL, (void*)event, &__h); if (__h) return; }
		QDialog::hideEvent(event);
	}
	bool nativeEvent(const QByteArray& eventType, void* message, qintptr* result) override {
		int __h = 0;
		if (__goOn[39] && cppgo_virt_int_voidp_voidp_longlongp) { int __r = ((int(*)(long long, long long, void*, void*, long long*, int*))cppgo_virt_int_voidp_voidp_longlongp)(__goId, -2039166586246443706LL, (void*)&eventType, (void*)message, (long long*)result, &__h); if (__h) return __r; }
		return QDialog::nativeEvent(eventType, message, result);
	}
	void changeEvent(QEvent* arg1) override {
		int __h = 0;
		if (__goOn[40] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4959376063826165943LL, (void*)arg1, &__h); if (__h) return; }
		QDialog::changeEvent(arg1);
	}
	int metric(QPaintDevice::PaintDeviceMetric arg1) const override {
		int __h = 0;
		if (__goOn[41] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 427629649267311777LL, (int)(arg1), &__h); if (__h) return __r; }
		return QDialog::metric(arg1);
	}
	void initPainter(QPainter* painter) const override {
		int __h = 0;
		if (__goOn[42] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6477500296980050210LL, (void*)painter, &__h); if (__h) return; }
		QDialog::initPainter(painter);
	}
	QPaintDevice* redirected(QPoint* offset) const override {
		int __h = 0;
		if (__goOn[43] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 4127240633310655048LL, (void*)offset, &__h); if (__h) return (QPaintDevice*)__r; }
		return QDialog::redirected(offset);
	}
	QPainter* sharedPainter() const override {
		int __h = 0;
		if (__goOn[44] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 5792771328741774991LL, &__h); if (__h) return (QPainter*)__r; }
		return QDialog::sharedPainter();
	}
	void inputMethodEvent(QInputMethodEvent* arg1) override {
		int __h = 0;
		if (__goOn[45] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5438819661942843222LL, (void*)arg1, &__h); if (__h) return; }
		QDialog::inputMethodEvent(arg1);
	}
	QVariant inputMethodQuery(Qt::InputMethodQuery arg1) const override {
		int __h = 0;
		if (__goOn[46] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7345853097340709638LL, (int)(arg1), &__h); if (__h) return *(QVariant*)__r; }
		return QDialog::inputMethodQuery(arg1);
	}
	bool focusNextPrevChild(bool next) override {
		int __h = 0;
		if (__goOn[47] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 5898328517505784191LL, next, &__h); if (__h) return __r; }
		return QDialog::focusNextPrevChild(next);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[48] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QDialog::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[49] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QDialog::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[50] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QDialog::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[51] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QDialog::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[52] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QDialog::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QDialog;
static std::mutex __connsMu_QDialog;

extern "C" {

void* QDialog_base(void* self) { return (void*)static_cast<QWidget*>((QDialog*)self); }
void QDialog_chain(void* self, void** out) {
	QDialog* __p = (QDialog*)self;
	QWidget* __b0 = static_cast<QWidget*>(__p); out[0] = (void*)__b0;
	QObject* __b1 = static_cast<QObject*>(__b0); out[1] = (void*)__b1;
}
void QDialog__set_override(void* self, int idx) { ((QDialog__Go*)self)->__goOn[idx] = 1; }
void* QDialog_fromQObject(void* o) { return (void*)qobject_cast<QDialog*>((QObject*)o); }
void* QDialog_new(long long __goid, void* parent) {
	return new QDialog__Go(__goid, (QWidget*)parent);
}
void* QDialog_new2(long long __goid) {
	return new QDialog__Go(__goid);
}
void* QDialog_new3(long long __goid, void* parent, int f) {
	return new QDialog__Go(__goid, (QWidget*)parent, QFlags<Qt::WindowType>(f));
}
char* QDialog_tr(const char* s, long long* __out_len) {
	auto&& __r = QDialog::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QDialog_result(void* self) {
	auto&& __r = ((QDialog*)self)->result();
	return __r;
}
void QDialog_setSizeGripEnabled(void* self, int arg1) {
	((QDialog*)self)->setSizeGripEnabled(arg1);
}
int QDialog_isSizeGripEnabled(void* self) {
	auto&& __r = ((QDialog*)self)->isSizeGripEnabled();
	return __r;
}
void QDialog_setModal(void* self, int modal) {
	((QDialog*)self)->setModal(modal);
}
void QDialog_setResult(void* self, int r) {
	((QDialog*)self)->setResult(r);
}
void QDialog_finished(void* self, int result) {
	((QDialog*)self)->finished(result);
}
void QDialog_accepted(void* self) {
	((QDialog*)self)->accepted();
}
void QDialog_rejected(void* self) {
	((QDialog*)self)->rejected();
}
void QDialog_open(void* self) {
	((QDialog*)self)->open();
}
int QDialog_exec(void* self) {
	auto&& __r = ((QDialog*)self)->exec();
	return __r;
}
void QDialog_done(void* self, int arg1) {
	((QDialog*)self)->done(arg1);
}
void QDialog_accept(void* self) {
	((QDialog*)self)->accept();
}
void QDialog_reject(void* self) {
	((QDialog*)self)->reject();
}
void QDialog_adjustPosition(void* self, void* arg1) {
	((QDialog__Go*)self)->adjustPosition((QWidget*)arg1);
}
void QDialog_setTabOrder(void* arg1, void* arg2) {
	QDialog::setTabOrder((QWidget*)arg1, (QWidget*)arg2);
}
void* QDialog_mouseGrabber() {
	auto&& __r = QDialog::mouseGrabber();
	return (void*)(__r);
}
void* QDialog_keyboardGrabber() {
	auto&& __r = QDialog::keyboardGrabber();
	return (void*)(__r);
}
void* QDialog_find(unsigned long long arg1) {
	auto&& __r = QDialog::find(arg1);
	return (void*)(__r);
}
void* QDialog_createWindowContainer(void* window) {
	auto&& __r = QDialog::createWindowContainer((QWindow*)window);
	return (void*)(__r);
}
void QDialog_updateMicroFocus(void* self) {
	((QDialog__Go*)self)->updateMicroFocus();
}
void QDialog_create(void* self) {
	((QDialog__Go*)self)->create();
}
void QDialog_destroy(void* self) {
	((QDialog__Go*)self)->destroy();
}
int QDialog_focusNextChild(void* self) {
	auto&& __r = ((QDialog__Go*)self)->focusNextChild();
	return __r;
}
int QDialog_focusPreviousChild(void* self) {
	auto&& __r = ((QDialog__Go*)self)->focusPreviousChild();
	return __r;
}
void* QDialog_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QDialog::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QDialog_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QDialog::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QDialog_sender(void* self) {
	auto&& __r = ((QDialog__Go*)self)->sender();
	return (void*)(__r);
}
int QDialog_senderSignalIndex(void* self) {
	auto&& __r = ((QDialog__Go*)self)->senderSignalIndex();
	return __r;
}
int QDialog_receivers(void* self, const char* signal) {
	auto&& __r = ((QDialog__Go*)self)->receivers(signal);
	return __r;
}
int QDialog_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QDialog__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
double QDialog_devicePixelRatioFScale() {
	auto&& __r = QDialog::devicePixelRatioFScale();
	return __r;
}
int QDialog_encodeMetricF(int metric, double value) {
	auto&& __r = QDialog::encodeMetricF((QPaintDevice::PaintDeviceMetric)(metric), value);
	return __r;
}
double QDialog_getDecodedMetricF(void* self, int metricA, int metricB) {
	auto&& __r = ((QDialog__Go*)self)->getDecodedMetricF((QPaintDevice::PaintDeviceMetric)(metricA), (QPaintDevice::PaintDeviceMetric)(metricB));
	return __r;
}
char* QDialog_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QDialog::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QDialog_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QDialog::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QDialog_createWindowContainer2(void* window, void* parent) {
	auto&& __r = QDialog::createWindowContainer((QWindow*)window, (QWidget*)parent);
	return (void*)(__r);
}
void* QDialog_createWindowContainer3(void* window, void* parent, int flags) {
	auto&& __r = QDialog::createWindowContainer((QWindow*)window, (QWidget*)parent, QFlags<Qt::WindowType>(flags));
	return (void*)(__r);
}
void QDialog_updateMicroFocus2(void* self, int query) {
	((QDialog__Go*)self)->updateMicroFocus((Qt::InputMethodQuery)(query));
}
void QDialog_create2(void* self, unsigned long long arg1) {
	((QDialog__Go*)self)->create(arg1);
}
void QDialog_create3(void* self, unsigned long long arg1, int initializeWindow) {
	((QDialog__Go*)self)->create(arg1, initializeWindow);
}
void QDialog_create4(void* self, unsigned long long arg1, int initializeWindow, int destroyOldWindow) {
	((QDialog__Go*)self)->create(arg1, initializeWindow, destroyOldWindow);
}
void QDialog_destroy2(void* self, int destroyWindow) {
	((QDialog__Go*)self)->destroy(destroyWindow);
}
void QDialog_destroy3(void* self, int destroyWindow, int destroySubWindows) {
	((QDialog__Go*)self)->destroy(destroyWindow, destroySubWindows);
}
void* QDialog_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QDialog::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QDialog_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QDialog::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void QDialog_open_default(void* self) {
	((QDialog*)self)->QDialog::open();
}
int QDialog_exec_default(void* self) {
	auto&& __r = ((QDialog*)self)->QDialog::exec();
	return __r;
}
void QDialog_done_default(void* self, int arg1) {
	((QDialog*)self)->QDialog::done(arg1);
}
void QDialog_accept_default(void* self) {
	((QDialog*)self)->QDialog::accept();
}
void QDialog_reject_default(void* self) {
	((QDialog*)self)->QDialog::reject();
}
void* QDialog_metaObject_default(void* self) {
	auto&& __r = ((QDialog*)self)->QDialog::metaObject();
	return (void*)(__r);
}
void QDialog_setVisible_default(void* self, int visible) {
	((QDialog*)self)->QDialog::setVisible(visible);
}
void* QDialog_sizeHint_default(void* self) {
	auto&& __r = ((QDialog*)self)->QDialog::sizeHint();
	return new QSize(__r);
}
void* QDialog_minimumSizeHint_default(void* self) {
	auto&& __r = ((QDialog*)self)->QDialog::minimumSizeHint();
	return new QSize(__r);
}
void QDialog_keyPressEvent_default(void* self, void* arg1) {
	((QDialog__Go*)self)->cppgoBase_keyPressEvent((QKeyEvent*)arg1);
}
void QDialog_closeEvent_default(void* self, void* arg1) {
	((QDialog__Go*)self)->cppgoBase_closeEvent((QCloseEvent*)arg1);
}
void QDialog_showEvent_default(void* self, void* arg1) {
	((QDialog__Go*)self)->cppgoBase_showEvent((QShowEvent*)arg1);
}
void QDialog_resizeEvent_default(void* self, void* arg1) {
	((QDialog__Go*)self)->cppgoBase_resizeEvent((QResizeEvent*)arg1);
}
void QDialog_contextMenuEvent_default(void* self, void* arg1) {
	((QDialog__Go*)self)->cppgoBase_contextMenuEvent((QContextMenuEvent*)arg1);
}
int QDialog_eventFilter_default(void* self, void* arg1, void* arg2) {
	auto&& __r = ((QDialog__Go*)self)->cppgoBase_eventFilter((QObject*)arg1, (QEvent*)arg2);
	return __r;
}
int QDialog_devType_default(void* self) {
	auto&& __r = ((QDialog*)self)->QDialog::devType();
	return __r;
}
int QDialog_heightForWidth_default(void* self, int arg1) {
	auto&& __r = ((QDialog*)self)->QDialog::heightForWidth(arg1);
	return __r;
}
int QDialog_hasHeightForWidth_default(void* self) {
	auto&& __r = ((QDialog*)self)->QDialog::hasHeightForWidth();
	return __r;
}
void* QDialog_paintEngine_default(void* self) {
	auto&& __r = ((QDialog*)self)->QDialog::paintEngine();
	return (void*)(__r);
}
int QDialog_event_default(void* self, void* event) {
	auto&& __r = ((QDialog__Go*)self)->cppgoBase_event((QEvent*)event);
	return __r;
}
void QDialog_mousePressEvent_default(void* self, void* event) {
	((QDialog__Go*)self)->cppgoBase_mousePressEvent((QMouseEvent*)event);
}
void QDialog_mouseReleaseEvent_default(void* self, void* event) {
	((QDialog__Go*)self)->cppgoBase_mouseReleaseEvent((QMouseEvent*)event);
}
void QDialog_mouseDoubleClickEvent_default(void* self, void* event) {
	((QDialog__Go*)self)->cppgoBase_mouseDoubleClickEvent((QMouseEvent*)event);
}
void QDialog_mouseMoveEvent_default(void* self, void* event) {
	((QDialog__Go*)self)->cppgoBase_mouseMoveEvent((QMouseEvent*)event);
}
void QDialog_wheelEvent_default(void* self, void* event) {
	((QDialog__Go*)self)->cppgoBase_wheelEvent((QWheelEvent*)event);
}
void QDialog_keyReleaseEvent_default(void* self, void* event) {
	((QDialog__Go*)self)->cppgoBase_keyReleaseEvent((QKeyEvent*)event);
}
void QDialog_focusInEvent_default(void* self, void* event) {
	((QDialog__Go*)self)->cppgoBase_focusInEvent((QFocusEvent*)event);
}
void QDialog_focusOutEvent_default(void* self, void* event) {
	((QDialog__Go*)self)->cppgoBase_focusOutEvent((QFocusEvent*)event);
}
void QDialog_enterEvent_default(void* self, void* event) {
	((QDialog__Go*)self)->cppgoBase_enterEvent((QEnterEvent*)event);
}
void QDialog_leaveEvent_default(void* self, void* event) {
	((QDialog__Go*)self)->cppgoBase_leaveEvent((QEvent*)event);
}
void QDialog_paintEvent_default(void* self, void* event) {
	((QDialog__Go*)self)->cppgoBase_paintEvent((QPaintEvent*)event);
}
void QDialog_moveEvent_default(void* self, void* event) {
	((QDialog__Go*)self)->cppgoBase_moveEvent((QMoveEvent*)event);
}
void QDialog_tabletEvent_default(void* self, void* event) {
	((QDialog__Go*)self)->cppgoBase_tabletEvent((QTabletEvent*)event);
}
void QDialog_actionEvent_default(void* self, void* event) {
	((QDialog__Go*)self)->cppgoBase_actionEvent((QActionEvent*)event);
}
void QDialog_dragEnterEvent_default(void* self, void* event) {
	((QDialog__Go*)self)->cppgoBase_dragEnterEvent((QDragEnterEvent*)event);
}
void QDialog_dragMoveEvent_default(void* self, void* event) {
	((QDialog__Go*)self)->cppgoBase_dragMoveEvent((QDragMoveEvent*)event);
}
void QDialog_dragLeaveEvent_default(void* self, void* event) {
	((QDialog__Go*)self)->cppgoBase_dragLeaveEvent((QDragLeaveEvent*)event);
}
void QDialog_dropEvent_default(void* self, void* event) {
	((QDialog__Go*)self)->cppgoBase_dropEvent((QDropEvent*)event);
}
void QDialog_hideEvent_default(void* self, void* event) {
	((QDialog__Go*)self)->cppgoBase_hideEvent((QHideEvent*)event);
}
void QDialog_changeEvent_default(void* self, void* arg1) {
	((QDialog__Go*)self)->cppgoBase_changeEvent((QEvent*)arg1);
}
int QDialog_metric_default(void* self, int arg1) {
	auto&& __r = ((QDialog__Go*)self)->cppgoBase_metric(arg1);
	return __r;
}
void QDialog_initPainter_default(void* self, void* painter) {
	((QDialog__Go*)self)->cppgoBase_initPainter((QPainter*)painter);
}
void* QDialog_redirected_default(void* self, void* offset) {
	auto&& __r = ((QDialog__Go*)self)->cppgoBase_redirected((QPoint*)offset);
	return (void*)(__r);
}
void* QDialog_sharedPainter_default(void* self) {
	auto&& __r = ((QDialog__Go*)self)->cppgoBase_sharedPainter();
	return (void*)(__r);
}
void QDialog_inputMethodEvent_default(void* self, void* arg1) {
	((QDialog__Go*)self)->cppgoBase_inputMethodEvent((QInputMethodEvent*)arg1);
}
void* QDialog_inputMethodQuery_default(void* self, int arg1) {
	auto&& __r = ((QDialog*)self)->QDialog::inputMethodQuery((Qt::InputMethodQuery)(arg1));
	return new QVariant(__r);
}
int QDialog_focusNextPrevChild_default(void* self, int next) {
	auto&& __r = ((QDialog__Go*)self)->cppgoBase_focusNextPrevChild(next);
	return __r;
}
void QDialog_timerEvent_default(void* self, void* event) {
	((QDialog__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QDialog_childEvent_default(void* self, void* event) {
	((QDialog__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QDialog_customEvent_default(void* self, void* event) {
	((QDialog__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QDialog_connectNotify_default(void* self, void* signal) {
	((QDialog__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QDialog_disconnectNotify_default(void* self, void* signal) {
	((QDialog__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QDialog_connect_finished(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDialog*)self, &QDialog::finished, (QDialog*)self, [__id](int result){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, result);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDialog);
	__conns_QDialog[__id] = __c;
}
void QDialog_connect_accepted(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDialog*)self, &QDialog::accepted, (QDialog*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDialog);
	__conns_QDialog[__id] = __c;
}
void QDialog_connect_rejected(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDialog*)self, &QDialog::rejected, (QDialog*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDialog);
	__conns_QDialog[__id] = __c;
}
void QDialog_connect_windowTitleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDialog*)self, &QDialog::windowTitleChanged, (QDialog*)self, [__id](const QString& title){
		QByteArray title_ba = title.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)title_ba.constData(), title_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDialog);
	__conns_QDialog[__id] = __c;
}
void QDialog_connect_windowIconChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDialog*)self, &QDialog::windowIconChanged, (QDialog*)self, [__id](const QIcon& icon){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&icon);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDialog);
	__conns_QDialog[__id] = __c;
}
void QDialog_connect_windowIconTextChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDialog*)self, &QDialog::windowIconTextChanged, (QDialog*)self, [__id](const QString& iconText){
		QByteArray iconText_ba = iconText.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)iconText_ba.constData(), iconText_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDialog);
	__conns_QDialog[__id] = __c;
}
void QDialog_connect_customContextMenuRequested(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDialog*)self, &QDialog::customContextMenuRequested, (QDialog*)self, [__id](const QPoint& pos){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&pos);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDialog);
	__conns_QDialog[__id] = __c;
}
void QDialog_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDialog*)self, &QDialog::destroyed, (QDialog*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDialog);
	__conns_QDialog[__id] = __c;
}
void QDialog_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDialog*)self, &QDialog::destroyed, (QDialog*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDialog);
	__conns_QDialog[__id] = __c;
}
void QDialog_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDialog*)self, &QDialog::objectNameChanged, (QDialog*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDialog);
	__conns_QDialog[__id] = __c;
}
void QDialog__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QDialog);
		auto __it = __conns_QDialog.find(id);
		if (__it == __conns_QDialog.end()) return;
		__c = __it->second;
		__conns_QDialog.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QDialog__Go_new(long long id) { return new QDialog__Go(id); }
void QDialog_delete(void* self) { delete (QDialog*)self; }

}

#include <QtWidgets/qformlayout.h>
#include <QtWidgets/qlayoutitem.h>
#include "gen_qformlayout__takerowresult.h"

extern "C" {

void* QFormLayout__TakeRowResult_new() {
	return new QFormLayout::TakeRowResult();
}
void* QFormLayout__TakeRowResult_new2(void* arg1) {
	return new QFormLayout::TakeRowResult(*(QFormLayout::TakeRowResult*)arg1);
}
void* QFormLayout__TakeRowResult_labelItem_get(void* self) {
	auto __r = ((QFormLayout::TakeRowResult*)self)->labelItem;
	return (void*)(__r);
}
void QFormLayout__TakeRowResult_labelItem_set(void* self, void* labelItem) {
	((QFormLayout::TakeRowResult*)self)->labelItem = (QLayoutItem*)labelItem;
}
void* QFormLayout__TakeRowResult_fieldItem_get(void* self) {
	auto __r = ((QFormLayout::TakeRowResult*)self)->fieldItem;
	return (void*)(__r);
}
void QFormLayout__TakeRowResult_fieldItem_set(void* self, void* fieldItem) {
	((QFormLayout::TakeRowResult*)self)->fieldItem = (QLayoutItem*)fieldItem;
}
void QFormLayout__TakeRowResult_delete(void* self) { delete (QFormLayout::TakeRowResult*)self; }

}

#include <QtWidgets/qgraphicsanchorlayout.h>
#include <QtWidgets/qgraphicslayout.h>
#include <QtWidgets/qgraphicslayoutitem.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qrect.h>
#include <QtCore/qsize.h>
#include <QtCore/qcoreevent.h>
#include <QtWidgets/qsizepolicy.h>
#include <QtWidgets/qgraphicsitem.h>
#include "gen_qgraphicsanchorlayout.h"
#include "gen_purego_cb.h"

class QGraphicsAnchorLayout__Go : public QGraphicsAnchorLayout {
public:
	long long __goId;
	unsigned char __goOn[10] = {};
	template<typename... A> QGraphicsAnchorLayout__Go(long long id, A&&... a) : QGraphicsAnchorLayout(std::forward<A>(a)...), __goId(id) {}
	using QGraphicsAnchorLayout::addChildLayoutItem;
	using QGraphicsAnchorLayout::setGraphicsItem;
	using QGraphicsAnchorLayout::setOwnedByLayout;
	QSizeF cppgoBase_sizeHint(int __a0, const QSizeF& __a1) { return QGraphicsAnchorLayout::sizeHint((Qt::SizeHint)__a0, __a1); }
	void removeAt(int index) override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 1230346564330554752LL, index, &__h); if (__h) return; }
		QGraphicsAnchorLayout::removeAt(index);
	}
	void setGeometry(const QRectF& rect) override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5684820848726969227LL, (void*)&rect, &__h); if (__h) return; }
		QGraphicsAnchorLayout::setGeometry(rect);
	}
	int count() const override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -5627843058667511180LL, &__h); if (__h) return __r; }
		return QGraphicsAnchorLayout::count();
	}
	QGraphicsLayoutItem* itemAt(int index) const override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7508512214013391107LL, index, &__h); if (__h) return (QGraphicsLayoutItem*)__r; }
		return QGraphicsAnchorLayout::itemAt(index);
	}
	void invalidate() override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 6991341317990632368LL, &__h); if (__h) return; }
		QGraphicsAnchorLayout::invalidate();
	}
	QSizeF sizeHint(Qt::SizeHint which, const QSizeF& constraint) const override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_voidp_int_voidp) { void* __r = ((void*(*)(long long, long long, int, void*, int*))cppgo_virt_voidp_int_voidp)(__goId, 729690261457684101LL, (int)(which), (void*)&constraint, &__h); if (__h) return *(QSizeF*)__r; }
		return QGraphicsAnchorLayout::sizeHint(which, constraint);
	}
	void getContentsMargins(qreal* left, qreal* top, qreal* right, qreal* bottom) const override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_doublep_doublep_doublep_doublep) { ((void(*)(long long, long long, double*, double*, double*, double*, int*))cppgo_virt_void_doublep_doublep_doublep_doublep)(__goId, 164760286202646972LL, (double*)left, (double*)top, (double*)right, (double*)bottom, &__h); if (__h) return; }
		QGraphicsAnchorLayout::getContentsMargins(left, top, right, bottom);
	}
	void updateGeometry() override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 8082455443689927668LL, &__h); if (__h) return; }
		QGraphicsAnchorLayout::updateGeometry();
	}
	void widgetEvent(QEvent* e) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 8535746785413991679LL, (void*)e, &__h); if (__h) return; }
		QGraphicsAnchorLayout::widgetEvent(e);
	}
	bool isEmpty() const override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -626822430251098498LL, &__h); if (__h) return __r; }
		return QGraphicsAnchorLayout::isEmpty();
	}
};

extern "C" {

void* QGraphicsAnchorLayout_base(void* self) { return (void*)static_cast<QGraphicsLayout*>((QGraphicsAnchorLayout*)self); }
void QGraphicsAnchorLayout_chain(void* self, void** out) {
	QGraphicsAnchorLayout* __p = (QGraphicsAnchorLayout*)self;
	QGraphicsLayout* __b0 = static_cast<QGraphicsLayout*>(__p); out[0] = (void*)__b0;
	QGraphicsLayoutItem* __b1 = static_cast<QGraphicsLayoutItem*>(__b0); out[1] = (void*)__b1;
}
void QGraphicsAnchorLayout__set_override(void* self, int idx) { ((QGraphicsAnchorLayout__Go*)self)->__goOn[idx] = 1; }
void* QGraphicsAnchorLayout_new(long long __goid) {
	return new QGraphicsAnchorLayout__Go(__goid);
}
void* QGraphicsAnchorLayout_new2(long long __goid, void* parent) {
	return new QGraphicsAnchorLayout__Go(__goid, (QGraphicsLayoutItem*)parent);
}
void* QGraphicsAnchorLayout_addAnchor(void* self, void* firstItem, int firstEdge, void* secondItem, int secondEdge) {
	auto&& __r = ((QGraphicsAnchorLayout*)self)->addAnchor((QGraphicsLayoutItem*)firstItem, (Qt::AnchorPoint)(firstEdge), (QGraphicsLayoutItem*)secondItem, (Qt::AnchorPoint)(secondEdge));
	return (void*)(__r);
}
void* QGraphicsAnchorLayout_anchor(void* self, void* firstItem, int firstEdge, void* secondItem, int secondEdge) {
	auto&& __r = ((QGraphicsAnchorLayout*)self)->anchor((QGraphicsLayoutItem*)firstItem, (Qt::AnchorPoint)(firstEdge), (QGraphicsLayoutItem*)secondItem, (Qt::AnchorPoint)(secondEdge));
	return (void*)(__r);
}
void QGraphicsAnchorLayout_addCornerAnchors(void* self, void* firstItem, int firstCorner, void* secondItem, int secondCorner) {
	((QGraphicsAnchorLayout*)self)->addCornerAnchors((QGraphicsLayoutItem*)firstItem, (Qt::Corner)(firstCorner), (QGraphicsLayoutItem*)secondItem, (Qt::Corner)(secondCorner));
}
void QGraphicsAnchorLayout_addAnchors(void* self, void* firstItem, void* secondItem) {
	((QGraphicsAnchorLayout*)self)->addAnchors((QGraphicsLayoutItem*)firstItem, (QGraphicsLayoutItem*)secondItem);
}
void QGraphicsAnchorLayout_setHorizontalSpacing(void* self, double spacing) {
	((QGraphicsAnchorLayout*)self)->setHorizontalSpacing(spacing);
}
void QGraphicsAnchorLayout_setVerticalSpacing(void* self, double spacing) {
	((QGraphicsAnchorLayout*)self)->setVerticalSpacing(spacing);
}
void QGraphicsAnchorLayout_setSpacing(void* self, double spacing) {
	((QGraphicsAnchorLayout*)self)->setSpacing(spacing);
}
double QGraphicsAnchorLayout_horizontalSpacing(void* self) {
	auto&& __r = ((QGraphicsAnchorLayout*)self)->horizontalSpacing();
	return __r;
}
double QGraphicsAnchorLayout_verticalSpacing(void* self) {
	auto&& __r = ((QGraphicsAnchorLayout*)self)->verticalSpacing();
	return __r;
}
void QGraphicsAnchorLayout_setInstantInvalidatePropagation(int enable) {
	QGraphicsAnchorLayout::setInstantInvalidatePropagation(enable);
}
int QGraphicsAnchorLayout_instantInvalidatePropagation() {
	auto&& __r = QGraphicsAnchorLayout::instantInvalidatePropagation();
	return __r;
}
void QGraphicsAnchorLayout_addChildLayoutItem(void* self, void* layoutItem) {
	((QGraphicsAnchorLayout__Go*)self)->addChildLayoutItem((QGraphicsLayoutItem*)layoutItem);
}
void QGraphicsAnchorLayout_setGraphicsItem(void* self, void* item) {
	((QGraphicsAnchorLayout__Go*)self)->setGraphicsItem((QGraphicsItem*)item);
}
void QGraphicsAnchorLayout_setOwnedByLayout(void* self, int ownedByLayout) {
	((QGraphicsAnchorLayout__Go*)self)->setOwnedByLayout(ownedByLayout);
}
void QGraphicsAnchorLayout_addAnchors2(void* self, void* firstItem, void* secondItem, int orientations) {
	((QGraphicsAnchorLayout*)self)->addAnchors((QGraphicsLayoutItem*)firstItem, (QGraphicsLayoutItem*)secondItem, QFlags<Qt::Orientation>(orientations));
}
void QGraphicsAnchorLayout_removeAt_default(void* self, int index) {
	((QGraphicsAnchorLayout*)self)->QGraphicsAnchorLayout::removeAt(index);
}
void QGraphicsAnchorLayout_setGeometry_default(void* self, void* rect) {
	((QGraphicsAnchorLayout*)self)->QGraphicsAnchorLayout::setGeometry(*(QRectF*)rect);
}
int QGraphicsAnchorLayout_count_default(void* self) {
	auto&& __r = ((QGraphicsAnchorLayout*)self)->QGraphicsAnchorLayout::count();
	return __r;
}
void* QGraphicsAnchorLayout_itemAt_default(void* self, int index) {
	auto&& __r = ((QGraphicsAnchorLayout*)self)->QGraphicsAnchorLayout::itemAt(index);
	return (void*)(__r);
}
void QGraphicsAnchorLayout_invalidate_default(void* self) {
	((QGraphicsAnchorLayout*)self)->QGraphicsAnchorLayout::invalidate();
}
void* QGraphicsAnchorLayout_sizeHint_default(void* self, int which, void* constraint) {
	auto&& __r = ((QGraphicsAnchorLayout__Go*)self)->cppgoBase_sizeHint(which, *(QSizeF*)constraint);
	return new QSizeF(__r);
}
void QGraphicsAnchorLayout_getContentsMargins_default(void* self, double* left, double* top, double* right, double* bottom) {
	qreal left_local{};
	qreal top_local{};
	qreal right_local{};
	qreal bottom_local{};
	((QGraphicsAnchorLayout*)self)->QGraphicsAnchorLayout::getContentsMargins(&left_local, &top_local, &right_local, &bottom_local);
	*left = left_local;
	*top = top_local;
	*right = right_local;
	*bottom = bottom_local;
}
void QGraphicsAnchorLayout_updateGeometry_default(void* self) {
	((QGraphicsAnchorLayout*)self)->QGraphicsAnchorLayout::updateGeometry();
}
void QGraphicsAnchorLayout_widgetEvent_default(void* self, void* e) {
	((QGraphicsAnchorLayout*)self)->QGraphicsAnchorLayout::widgetEvent((QEvent*)e);
}
int QGraphicsAnchorLayout_isEmpty_default(void* self) {
	auto&& __r = ((QGraphicsAnchorLayout*)self)->QGraphicsAnchorLayout::isEmpty();
	return __r;
}
void* QGraphicsAnchorLayout__Go_new(long long id) { return new QGraphicsAnchorLayout__Go(id); }
void QGraphicsAnchorLayout_delete(void* self) { delete (QGraphicsAnchorLayout*)self; }

}

#include <QtWidgets/qscrollarea.h>
#include <QtWidgets/qabstractscrollarea.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qsize.h>
#include <QtCore/qnamespace.h>
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
#include "gen_qscrollarea.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QScrollArea__Go : public QScrollArea {
public:
	long long __goId;
	unsigned char __goOn[53] = {};
	template<typename... A> QScrollArea__Go(long long id, A&&... a) : QScrollArea(std::forward<A>(a)...), __goId(id) {}
	using QScrollArea::setViewportMargins;
	using QScrollArea::viewportMargins;
	using QScrollArea::drawFrame;
	using QScrollArea::updateMicroFocus;
	using QScrollArea::create;
	using QScrollArea::destroy;
	using QScrollArea::focusNextChild;
	using QScrollArea::focusPreviousChild;
	using QScrollArea::sender;
	using QScrollArea::senderSignalIndex;
	using QScrollArea::receivers;
	using QScrollArea::isSignalConnected;
	using QScrollArea::getDecodedMetricF;
	bool cppgoBase_event(QEvent* __a0) { return QScrollArea::event(__a0); }
	bool cppgoBase_eventFilter(QObject* __a0, QEvent* __a1) { return QScrollArea::eventFilter(__a0, __a1); }
	void cppgoBase_resizeEvent(QResizeEvent* __a0) { QScrollArea::resizeEvent(__a0); }
	void cppgoBase_scrollContentsBy(int __a0, int __a1) { QScrollArea::scrollContentsBy(__a0, __a1); }
	QSize cppgoBase_viewportSizeHint() { return QScrollArea::viewportSizeHint(); }
	bool cppgoBase_viewportEvent(QEvent* __a0) { return QScrollArea::viewportEvent(__a0); }
	void cppgoBase_paintEvent(QPaintEvent* __a0) { QScrollArea::paintEvent(__a0); }
	void cppgoBase_mousePressEvent(QMouseEvent* __a0) { QScrollArea::mousePressEvent(__a0); }
	void cppgoBase_mouseReleaseEvent(QMouseEvent* __a0) { QScrollArea::mouseReleaseEvent(__a0); }
	void cppgoBase_mouseDoubleClickEvent(QMouseEvent* __a0) { QScrollArea::mouseDoubleClickEvent(__a0); }
	void cppgoBase_mouseMoveEvent(QMouseEvent* __a0) { QScrollArea::mouseMoveEvent(__a0); }
	void cppgoBase_wheelEvent(QWheelEvent* __a0) { QScrollArea::wheelEvent(__a0); }
	void cppgoBase_contextMenuEvent(QContextMenuEvent* __a0) { QScrollArea::contextMenuEvent(__a0); }
	void cppgoBase_dragEnterEvent(QDragEnterEvent* __a0) { QScrollArea::dragEnterEvent(__a0); }
	void cppgoBase_dragMoveEvent(QDragMoveEvent* __a0) { QScrollArea::dragMoveEvent(__a0); }
	void cppgoBase_dragLeaveEvent(QDragLeaveEvent* __a0) { QScrollArea::dragLeaveEvent(__a0); }
	void cppgoBase_dropEvent(QDropEvent* __a0) { QScrollArea::dropEvent(__a0); }
	void cppgoBase_keyPressEvent(QKeyEvent* __a0) { QScrollArea::keyPressEvent(__a0); }
	void cppgoBase_changeEvent(QEvent* __a0) { QScrollArea::changeEvent(__a0); }
	void cppgoBase_initStyleOption(QStyleOptionFrame* __a0) { QScrollArea::initStyleOption(__a0); }
	void cppgoBase_keyReleaseEvent(QKeyEvent* __a0) { QScrollArea::keyReleaseEvent(__a0); }
	void cppgoBase_focusInEvent(QFocusEvent* __a0) { QScrollArea::focusInEvent(__a0); }
	void cppgoBase_focusOutEvent(QFocusEvent* __a0) { QScrollArea::focusOutEvent(__a0); }
	void cppgoBase_enterEvent(QEnterEvent* __a0) { QScrollArea::enterEvent(__a0); }
	void cppgoBase_leaveEvent(QEvent* __a0) { QScrollArea::leaveEvent(__a0); }
	void cppgoBase_moveEvent(QMoveEvent* __a0) { QScrollArea::moveEvent(__a0); }
	void cppgoBase_closeEvent(QCloseEvent* __a0) { QScrollArea::closeEvent(__a0); }
	void cppgoBase_tabletEvent(QTabletEvent* __a0) { QScrollArea::tabletEvent(__a0); }
	void cppgoBase_actionEvent(QActionEvent* __a0) { QScrollArea::actionEvent(__a0); }
	void cppgoBase_showEvent(QShowEvent* __a0) { QScrollArea::showEvent(__a0); }
	void cppgoBase_hideEvent(QHideEvent* __a0) { QScrollArea::hideEvent(__a0); }
	int cppgoBase_metric(int __a0) { return QScrollArea::metric((QPaintDevice::PaintDeviceMetric)__a0); }
	void cppgoBase_initPainter(QPainter* __a0) { QScrollArea::initPainter(__a0); }
	QPaintDevice* cppgoBase_redirected(QPoint* __a0) { return QScrollArea::redirected(__a0); }
	QPainter* cppgoBase_sharedPainter() { return QScrollArea::sharedPainter(); }
	void cppgoBase_inputMethodEvent(QInputMethodEvent* __a0) { QScrollArea::inputMethodEvent(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QScrollArea::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QScrollArea::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QScrollArea::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QScrollArea::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QScrollArea::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QScrollArea::metaObject();
	}
	QSize sizeHint() const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 729690261457684101LL, &__h); if (__h) return *(QSize*)__r; }
		return QScrollArea::sizeHint();
	}
	bool focusNextPrevChild(bool next) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 5898328517505784191LL, next, &__h); if (__h) return __r; }
		return QScrollArea::focusNextPrevChild(next);
	}
	bool event(QEvent* arg1) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)arg1, &__h); if (__h) return __r; }
		return QScrollArea::event(arg1);
	}
	bool eventFilter(QObject* arg1, QEvent* arg2) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)arg1, (void*)arg2, &__h); if (__h) return __r; }
		return QScrollArea::eventFilter(arg1, arg2);
	}
	void resizeEvent(QResizeEvent* arg1) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4302917285655014661LL, (void*)arg1, &__h); if (__h) return; }
		QScrollArea::resizeEvent(arg1);
	}
	void scrollContentsBy(int dx, int dy) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_int_int) { ((void(*)(long long, long long, int, int, int*))cppgo_virt_void_int_int)(__goId, -6118001150126564745LL, dx, dy, &__h); if (__h) return; }
		QScrollArea::scrollContentsBy(dx, dy);
	}
	QSize viewportSizeHint() const override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 503579067252048239LL, &__h); if (__h) return *(QSize*)__r; }
		return QScrollArea::viewportSizeHint();
	}
	QSize minimumSizeHint() const override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 7049264599837532491LL, &__h); if (__h) return *(QSize*)__r; }
		return QScrollArea::minimumSizeHint();
	}
	void setupViewport(QWidget* viewport) override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4531996381736730688LL, (void*)viewport, &__h); if (__h) return; }
		QScrollArea::setupViewport(viewport);
	}
	bool viewportEvent(QEvent* arg1) override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, -1971465653061861519LL, (void*)arg1, &__h); if (__h) return __r; }
		return QScrollArea::viewportEvent(arg1);
	}
	void paintEvent(QPaintEvent* arg1) override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5290442163167458465LL, (void*)arg1, &__h); if (__h) return; }
		QScrollArea::paintEvent(arg1);
	}
	void mousePressEvent(QMouseEvent* arg1) override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -332090898078543319LL, (void*)arg1, &__h); if (__h) return; }
		QScrollArea::mousePressEvent(arg1);
	}
	void mouseReleaseEvent(QMouseEvent* arg1) override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4415920049297795213LL, (void*)arg1, &__h); if (__h) return; }
		QScrollArea::mouseReleaseEvent(arg1);
	}
	void mouseDoubleClickEvent(QMouseEvent* arg1) override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4029073905697023887LL, (void*)arg1, &__h); if (__h) return; }
		QScrollArea::mouseDoubleClickEvent(arg1);
	}
	void mouseMoveEvent(QMouseEvent* arg1) override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4255494004099806581LL, (void*)arg1, &__h); if (__h) return; }
		QScrollArea::mouseMoveEvent(arg1);
	}
	void wheelEvent(QWheelEvent* arg1) override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1919988340934511316LL, (void*)arg1, &__h); if (__h) return; }
		QScrollArea::wheelEvent(arg1);
	}
	void contextMenuEvent(QContextMenuEvent* arg1) override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6698043222227006015LL, (void*)arg1, &__h); if (__h) return; }
		QScrollArea::contextMenuEvent(arg1);
	}
	void dragEnterEvent(QDragEnterEvent* arg1) override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 206122715919917589LL, (void*)arg1, &__h); if (__h) return; }
		QScrollArea::dragEnterEvent(arg1);
	}
	void dragMoveEvent(QDragMoveEvent* arg1) override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1781545137352745978LL, (void*)arg1, &__h); if (__h) return; }
		QScrollArea::dragMoveEvent(arg1);
	}
	void dragLeaveEvent(QDragLeaveEvent* arg1) override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -3139904158373725006LL, (void*)arg1, &__h); if (__h) return; }
		QScrollArea::dragLeaveEvent(arg1);
	}
	void dropEvent(QDropEvent* arg1) override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 9107058682754591716LL, (void*)arg1, &__h); if (__h) return; }
		QScrollArea::dropEvent(arg1);
	}
	void keyPressEvent(QKeyEvent* arg1) override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1023791977679793185LL, (void*)arg1, &__h); if (__h) return; }
		QScrollArea::keyPressEvent(arg1);
	}
	void changeEvent(QEvent* arg1) override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4959376063826165943LL, (void*)arg1, &__h); if (__h) return; }
		QScrollArea::changeEvent(arg1);
	}
	void initStyleOption(QStyleOptionFrame* option) const override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -488693769771308531LL, (void*)option, &__h); if (__h) return; }
		QScrollArea::initStyleOption(option);
	}
	int devType() const override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8659679893868524044LL, &__h); if (__h) return __r; }
		return QScrollArea::devType();
	}
	void setVisible(bool visible) override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 2110182719375245575LL, visible, &__h); if (__h) return; }
		QScrollArea::setVisible(visible);
	}
	int heightForWidth(int arg1) const override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 4311534060495737339LL, arg1, &__h); if (__h) return __r; }
		return QScrollArea::heightForWidth(arg1);
	}
	bool hasHeightForWidth() const override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8266762397995423915LL, &__h); if (__h) return __r; }
		return QScrollArea::hasHeightForWidth();
	}
	QPaintEngine* paintEngine() const override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 6398007404756346381LL, &__h); if (__h) return (QPaintEngine*)__r; }
		return QScrollArea::paintEngine();
	}
	void keyReleaseEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 772598475950050299LL, (void*)event, &__h); if (__h) return; }
		QScrollArea::keyReleaseEvent(event);
	}
	void focusInEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 975043767033552552LL, (void*)event, &__h); if (__h) return; }
		QScrollArea::focusInEvent(event);
	}
	void focusOutEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[32] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -385596010576048993LL, (void*)event, &__h); if (__h) return; }
		QScrollArea::focusOutEvent(event);
	}
	void enterEvent(QEnterEvent* event) override {
		int __h = 0;
		if (__goOn[33] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6316270346668237895LL, (void*)event, &__h); if (__h) return; }
		QScrollArea::enterEvent(event);
	}
	void leaveEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[34] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -551070009623711640LL, (void*)event, &__h); if (__h) return; }
		QScrollArea::leaveEvent(event);
	}
	void moveEvent(QMoveEvent* event) override {
		int __h = 0;
		if (__goOn[35] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -9161949150489525396LL, (void*)event, &__h); if (__h) return; }
		QScrollArea::moveEvent(event);
	}
	void closeEvent(QCloseEvent* event) override {
		int __h = 0;
		if (__goOn[36] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1989937050019464783LL, (void*)event, &__h); if (__h) return; }
		QScrollArea::closeEvent(event);
	}
	void tabletEvent(QTabletEvent* event) override {
		int __h = 0;
		if (__goOn[37] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -212068866860701141LL, (void*)event, &__h); if (__h) return; }
		QScrollArea::tabletEvent(event);
	}
	void actionEvent(QActionEvent* event) override {
		int __h = 0;
		if (__goOn[38] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1697444140088555581LL, (void*)event, &__h); if (__h) return; }
		QScrollArea::actionEvent(event);
	}
	void showEvent(QShowEvent* event) override {
		int __h = 0;
		if (__goOn[39] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6825369174751916LL, (void*)event, &__h); if (__h) return; }
		QScrollArea::showEvent(event);
	}
	void hideEvent(QHideEvent* event) override {
		int __h = 0;
		if (__goOn[40] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -199718707876719129LL, (void*)event, &__h); if (__h) return; }
		QScrollArea::hideEvent(event);
	}
	bool nativeEvent(const QByteArray& eventType, void* message, qintptr* result) override {
		int __h = 0;
		if (__goOn[41] && cppgo_virt_int_voidp_voidp_longlongp) { int __r = ((int(*)(long long, long long, void*, void*, long long*, int*))cppgo_virt_int_voidp_voidp_longlongp)(__goId, -2039166586246443706LL, (void*)&eventType, (void*)message, (long long*)result, &__h); if (__h) return __r; }
		return QScrollArea::nativeEvent(eventType, message, result);
	}
	int metric(QPaintDevice::PaintDeviceMetric arg1) const override {
		int __h = 0;
		if (__goOn[42] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 427629649267311777LL, (int)(arg1), &__h); if (__h) return __r; }
		return QScrollArea::metric(arg1);
	}
	void initPainter(QPainter* painter) const override {
		int __h = 0;
		if (__goOn[43] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6477500296980050210LL, (void*)painter, &__h); if (__h) return; }
		QScrollArea::initPainter(painter);
	}
	QPaintDevice* redirected(QPoint* offset) const override {
		int __h = 0;
		if (__goOn[44] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 4127240633310655048LL, (void*)offset, &__h); if (__h) return (QPaintDevice*)__r; }
		return QScrollArea::redirected(offset);
	}
	QPainter* sharedPainter() const override {
		int __h = 0;
		if (__goOn[45] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 5792771328741774991LL, &__h); if (__h) return (QPainter*)__r; }
		return QScrollArea::sharedPainter();
	}
	void inputMethodEvent(QInputMethodEvent* arg1) override {
		int __h = 0;
		if (__goOn[46] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5438819661942843222LL, (void*)arg1, &__h); if (__h) return; }
		QScrollArea::inputMethodEvent(arg1);
	}
	QVariant inputMethodQuery(Qt::InputMethodQuery arg1) const override {
		int __h = 0;
		if (__goOn[47] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7345853097340709638LL, (int)(arg1), &__h); if (__h) return *(QVariant*)__r; }
		return QScrollArea::inputMethodQuery(arg1);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[48] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QScrollArea::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[49] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QScrollArea::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[50] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QScrollArea::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[51] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QScrollArea::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[52] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QScrollArea::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QScrollArea;
static std::mutex __connsMu_QScrollArea;

extern "C" {

void* QScrollArea_base(void* self) { return (void*)static_cast<QAbstractScrollArea*>((QScrollArea*)self); }
void QScrollArea_chain(void* self, void** out) {
	QScrollArea* __p = (QScrollArea*)self;
	QAbstractScrollArea* __b0 = static_cast<QAbstractScrollArea*>(__p); out[0] = (void*)__b0;
	QFrame* __b1 = static_cast<QFrame*>(__b0); out[1] = (void*)__b1;
	QWidget* __b2 = static_cast<QWidget*>(__b1); out[2] = (void*)__b2;
	QObject* __b3 = static_cast<QObject*>(__b2); out[3] = (void*)__b3;
}
void QScrollArea__set_override(void* self, int idx) { ((QScrollArea__Go*)self)->__goOn[idx] = 1; }
void* QScrollArea_fromQObject(void* o) { return (void*)qobject_cast<QScrollArea*>((QObject*)o); }
void* QScrollArea_new(long long __goid, void* parent) {
	return new QScrollArea__Go(__goid, (QWidget*)parent);
}
void* QScrollArea_new2(long long __goid) {
	return new QScrollArea__Go(__goid);
}
char* QScrollArea_tr(const char* s, long long* __out_len) {
	auto&& __r = QScrollArea::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QScrollArea_widget(void* self) {
	auto&& __r = ((QScrollArea*)self)->widget();
	return (void*)(__r);
}
void QScrollArea_setWidget(void* self, void* widget) {
	((QScrollArea*)self)->setWidget((QWidget*)widget);
}
void* QScrollArea_takeWidget(void* self) {
	auto&& __r = ((QScrollArea*)self)->takeWidget();
	return (void*)(__r);
}
int QScrollArea_widgetResizable(void* self) {
	auto&& __r = ((QScrollArea*)self)->widgetResizable();
	return __r;
}
void QScrollArea_setWidgetResizable(void* self, int resizable) {
	((QScrollArea*)self)->setWidgetResizable(resizable);
}
int QScrollArea_alignment(void* self) {
	auto&& __r = ((QScrollArea*)self)->alignment();
	return (int)(__r);
}
void QScrollArea_setAlignment(void* self, int arg1) {
	((QScrollArea*)self)->setAlignment(QFlags<Qt::AlignmentFlag>(arg1));
}
void QScrollArea_ensureVisible(void* self, int x, int y) {
	((QScrollArea*)self)->ensureVisible(x, y);
}
void QScrollArea_ensureWidgetVisible(void* self, void* childWidget) {
	((QScrollArea*)self)->ensureWidgetVisible((QWidget*)childWidget);
}
void QScrollArea_setViewportMargins(void* self, int left, int top, int right, int bottom) {
	((QScrollArea__Go*)self)->setViewportMargins(left, top, right, bottom);
}
void QScrollArea_setViewportMargins2(void* self, void* margins) {
	((QScrollArea__Go*)self)->setViewportMargins(*(QMargins*)margins);
}
void* QScrollArea_viewportMargins(void* self) {
	auto&& __r = ((QScrollArea__Go*)self)->viewportMargins();
	return new QMargins(__r);
}
void QScrollArea_drawFrame(void* self, void* arg1) {
	((QScrollArea__Go*)self)->drawFrame((QPainter*)arg1);
}
void QScrollArea_setTabOrder(void* arg1, void* arg2) {
	QScrollArea::setTabOrder((QWidget*)arg1, (QWidget*)arg2);
}
void* QScrollArea_mouseGrabber() {
	auto&& __r = QScrollArea::mouseGrabber();
	return (void*)(__r);
}
void* QScrollArea_keyboardGrabber() {
	auto&& __r = QScrollArea::keyboardGrabber();
	return (void*)(__r);
}
void* QScrollArea_find(unsigned long long arg1) {
	auto&& __r = QScrollArea::find(arg1);
	return (void*)(__r);
}
void* QScrollArea_createWindowContainer(void* window) {
	auto&& __r = QScrollArea::createWindowContainer((QWindow*)window);
	return (void*)(__r);
}
void QScrollArea_updateMicroFocus(void* self) {
	((QScrollArea__Go*)self)->updateMicroFocus();
}
void QScrollArea_create(void* self) {
	((QScrollArea__Go*)self)->create();
}
void QScrollArea_destroy(void* self) {
	((QScrollArea__Go*)self)->destroy();
}
int QScrollArea_focusNextChild(void* self) {
	auto&& __r = ((QScrollArea__Go*)self)->focusNextChild();
	return __r;
}
int QScrollArea_focusPreviousChild(void* self) {
	auto&& __r = ((QScrollArea__Go*)self)->focusPreviousChild();
	return __r;
}
void* QScrollArea_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QScrollArea::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QScrollArea_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QScrollArea::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QScrollArea_sender(void* self) {
	auto&& __r = ((QScrollArea__Go*)self)->sender();
	return (void*)(__r);
}
int QScrollArea_senderSignalIndex(void* self) {
	auto&& __r = ((QScrollArea__Go*)self)->senderSignalIndex();
	return __r;
}
int QScrollArea_receivers(void* self, const char* signal) {
	auto&& __r = ((QScrollArea__Go*)self)->receivers(signal);
	return __r;
}
int QScrollArea_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QScrollArea__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
double QScrollArea_devicePixelRatioFScale() {
	auto&& __r = QScrollArea::devicePixelRatioFScale();
	return __r;
}
int QScrollArea_encodeMetricF(int metric, double value) {
	auto&& __r = QScrollArea::encodeMetricF((QPaintDevice::PaintDeviceMetric)(metric), value);
	return __r;
}
double QScrollArea_getDecodedMetricF(void* self, int metricA, int metricB) {
	auto&& __r = ((QScrollArea__Go*)self)->getDecodedMetricF((QPaintDevice::PaintDeviceMetric)(metricA), (QPaintDevice::PaintDeviceMetric)(metricB));
	return __r;
}
char* QScrollArea_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QScrollArea::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QScrollArea_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QScrollArea::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QScrollArea_ensureVisible2(void* self, int x, int y, int xmargin) {
	((QScrollArea*)self)->ensureVisible(x, y, xmargin);
}
void QScrollArea_ensureVisible3(void* self, int x, int y, int xmargin, int ymargin) {
	((QScrollArea*)self)->ensureVisible(x, y, xmargin, ymargin);
}
void QScrollArea_ensureWidgetVisible2(void* self, void* childWidget, int xmargin) {
	((QScrollArea*)self)->ensureWidgetVisible((QWidget*)childWidget, xmargin);
}
void QScrollArea_ensureWidgetVisible3(void* self, void* childWidget, int xmargin, int ymargin) {
	((QScrollArea*)self)->ensureWidgetVisible((QWidget*)childWidget, xmargin, ymargin);
}
void* QScrollArea_createWindowContainer2(void* window, void* parent) {
	auto&& __r = QScrollArea::createWindowContainer((QWindow*)window, (QWidget*)parent);
	return (void*)(__r);
}
void* QScrollArea_createWindowContainer3(void* window, void* parent, int flags) {
	auto&& __r = QScrollArea::createWindowContainer((QWindow*)window, (QWidget*)parent, QFlags<Qt::WindowType>(flags));
	return (void*)(__r);
}
void QScrollArea_updateMicroFocus2(void* self, int query) {
	((QScrollArea__Go*)self)->updateMicroFocus((Qt::InputMethodQuery)(query));
}
void QScrollArea_create2(void* self, unsigned long long arg1) {
	((QScrollArea__Go*)self)->create(arg1);
}
void QScrollArea_create3(void* self, unsigned long long arg1, int initializeWindow) {
	((QScrollArea__Go*)self)->create(arg1, initializeWindow);
}
void QScrollArea_create4(void* self, unsigned long long arg1, int initializeWindow, int destroyOldWindow) {
	((QScrollArea__Go*)self)->create(arg1, initializeWindow, destroyOldWindow);
}
void QScrollArea_destroy2(void* self, int destroyWindow) {
	((QScrollArea__Go*)self)->destroy(destroyWindow);
}
void QScrollArea_destroy3(void* self, int destroyWindow, int destroySubWindows) {
	((QScrollArea__Go*)self)->destroy(destroyWindow, destroySubWindows);
}
void* QScrollArea_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QScrollArea::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QScrollArea_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QScrollArea::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QScrollArea_metaObject_default(void* self) {
	auto&& __r = ((QScrollArea*)self)->QScrollArea::metaObject();
	return (void*)(__r);
}
void* QScrollArea_sizeHint_default(void* self) {
	auto&& __r = ((QScrollArea*)self)->QScrollArea::sizeHint();
	return new QSize(__r);
}
int QScrollArea_focusNextPrevChild_default(void* self, int next) {
	auto&& __r = ((QScrollArea*)self)->QScrollArea::focusNextPrevChild(next);
	return __r;
}
int QScrollArea_event_default(void* self, void* arg1) {
	auto&& __r = ((QScrollArea__Go*)self)->cppgoBase_event((QEvent*)arg1);
	return __r;
}
int QScrollArea_eventFilter_default(void* self, void* arg1, void* arg2) {
	auto&& __r = ((QScrollArea__Go*)self)->cppgoBase_eventFilter((QObject*)arg1, (QEvent*)arg2);
	return __r;
}
void QScrollArea_resizeEvent_default(void* self, void* arg1) {
	((QScrollArea__Go*)self)->cppgoBase_resizeEvent((QResizeEvent*)arg1);
}
void QScrollArea_scrollContentsBy_default(void* self, int dx, int dy) {
	((QScrollArea__Go*)self)->cppgoBase_scrollContentsBy(dx, dy);
}
void* QScrollArea_viewportSizeHint_default(void* self) {
	auto&& __r = ((QScrollArea__Go*)self)->cppgoBase_viewportSizeHint();
	return new QSize(__r);
}
void* QScrollArea_minimumSizeHint_default(void* self) {
	auto&& __r = ((QScrollArea*)self)->QScrollArea::minimumSizeHint();
	return new QSize(__r);
}
void QScrollArea_setupViewport_default(void* self, void* viewport) {
	((QScrollArea*)self)->QScrollArea::setupViewport((QWidget*)viewport);
}
int QScrollArea_viewportEvent_default(void* self, void* arg1) {
	auto&& __r = ((QScrollArea__Go*)self)->cppgoBase_viewportEvent((QEvent*)arg1);
	return __r;
}
void QScrollArea_paintEvent_default(void* self, void* arg1) {
	((QScrollArea__Go*)self)->cppgoBase_paintEvent((QPaintEvent*)arg1);
}
void QScrollArea_mousePressEvent_default(void* self, void* arg1) {
	((QScrollArea__Go*)self)->cppgoBase_mousePressEvent((QMouseEvent*)arg1);
}
void QScrollArea_mouseReleaseEvent_default(void* self, void* arg1) {
	((QScrollArea__Go*)self)->cppgoBase_mouseReleaseEvent((QMouseEvent*)arg1);
}
void QScrollArea_mouseDoubleClickEvent_default(void* self, void* arg1) {
	((QScrollArea__Go*)self)->cppgoBase_mouseDoubleClickEvent((QMouseEvent*)arg1);
}
void QScrollArea_mouseMoveEvent_default(void* self, void* arg1) {
	((QScrollArea__Go*)self)->cppgoBase_mouseMoveEvent((QMouseEvent*)arg1);
}
void QScrollArea_wheelEvent_default(void* self, void* arg1) {
	((QScrollArea__Go*)self)->cppgoBase_wheelEvent((QWheelEvent*)arg1);
}
void QScrollArea_contextMenuEvent_default(void* self, void* arg1) {
	((QScrollArea__Go*)self)->cppgoBase_contextMenuEvent((QContextMenuEvent*)arg1);
}
void QScrollArea_dragEnterEvent_default(void* self, void* arg1) {
	((QScrollArea__Go*)self)->cppgoBase_dragEnterEvent((QDragEnterEvent*)arg1);
}
void QScrollArea_dragMoveEvent_default(void* self, void* arg1) {
	((QScrollArea__Go*)self)->cppgoBase_dragMoveEvent((QDragMoveEvent*)arg1);
}
void QScrollArea_dragLeaveEvent_default(void* self, void* arg1) {
	((QScrollArea__Go*)self)->cppgoBase_dragLeaveEvent((QDragLeaveEvent*)arg1);
}
void QScrollArea_dropEvent_default(void* self, void* arg1) {
	((QScrollArea__Go*)self)->cppgoBase_dropEvent((QDropEvent*)arg1);
}
void QScrollArea_keyPressEvent_default(void* self, void* arg1) {
	((QScrollArea__Go*)self)->cppgoBase_keyPressEvent((QKeyEvent*)arg1);
}
void QScrollArea_changeEvent_default(void* self, void* arg1) {
	((QScrollArea__Go*)self)->cppgoBase_changeEvent((QEvent*)arg1);
}
void QScrollArea_initStyleOption_default(void* self, void* option) {
	((QScrollArea__Go*)self)->cppgoBase_initStyleOption((QStyleOptionFrame*)option);
}
int QScrollArea_devType_default(void* self) {
	auto&& __r = ((QScrollArea*)self)->QScrollArea::devType();
	return __r;
}
void QScrollArea_setVisible_default(void* self, int visible) {
	((QScrollArea*)self)->QScrollArea::setVisible(visible);
}
int QScrollArea_heightForWidth_default(void* self, int arg1) {
	auto&& __r = ((QScrollArea*)self)->QScrollArea::heightForWidth(arg1);
	return __r;
}
int QScrollArea_hasHeightForWidth_default(void* self) {
	auto&& __r = ((QScrollArea*)self)->QScrollArea::hasHeightForWidth();
	return __r;
}
void* QScrollArea_paintEngine_default(void* self) {
	auto&& __r = ((QScrollArea*)self)->QScrollArea::paintEngine();
	return (void*)(__r);
}
void QScrollArea_keyReleaseEvent_default(void* self, void* event) {
	((QScrollArea__Go*)self)->cppgoBase_keyReleaseEvent((QKeyEvent*)event);
}
void QScrollArea_focusInEvent_default(void* self, void* event) {
	((QScrollArea__Go*)self)->cppgoBase_focusInEvent((QFocusEvent*)event);
}
void QScrollArea_focusOutEvent_default(void* self, void* event) {
	((QScrollArea__Go*)self)->cppgoBase_focusOutEvent((QFocusEvent*)event);
}
void QScrollArea_enterEvent_default(void* self, void* event) {
	((QScrollArea__Go*)self)->cppgoBase_enterEvent((QEnterEvent*)event);
}
void QScrollArea_leaveEvent_default(void* self, void* event) {
	((QScrollArea__Go*)self)->cppgoBase_leaveEvent((QEvent*)event);
}
void QScrollArea_moveEvent_default(void* self, void* event) {
	((QScrollArea__Go*)self)->cppgoBase_moveEvent((QMoveEvent*)event);
}
void QScrollArea_closeEvent_default(void* self, void* event) {
	((QScrollArea__Go*)self)->cppgoBase_closeEvent((QCloseEvent*)event);
}
void QScrollArea_tabletEvent_default(void* self, void* event) {
	((QScrollArea__Go*)self)->cppgoBase_tabletEvent((QTabletEvent*)event);
}
void QScrollArea_actionEvent_default(void* self, void* event) {
	((QScrollArea__Go*)self)->cppgoBase_actionEvent((QActionEvent*)event);
}
void QScrollArea_showEvent_default(void* self, void* event) {
	((QScrollArea__Go*)self)->cppgoBase_showEvent((QShowEvent*)event);
}
void QScrollArea_hideEvent_default(void* self, void* event) {
	((QScrollArea__Go*)self)->cppgoBase_hideEvent((QHideEvent*)event);
}
int QScrollArea_metric_default(void* self, int arg1) {
	auto&& __r = ((QScrollArea__Go*)self)->cppgoBase_metric(arg1);
	return __r;
}
void QScrollArea_initPainter_default(void* self, void* painter) {
	((QScrollArea__Go*)self)->cppgoBase_initPainter((QPainter*)painter);
}
void* QScrollArea_redirected_default(void* self, void* offset) {
	auto&& __r = ((QScrollArea__Go*)self)->cppgoBase_redirected((QPoint*)offset);
	return (void*)(__r);
}
void* QScrollArea_sharedPainter_default(void* self) {
	auto&& __r = ((QScrollArea__Go*)self)->cppgoBase_sharedPainter();
	return (void*)(__r);
}
void QScrollArea_inputMethodEvent_default(void* self, void* arg1) {
	((QScrollArea__Go*)self)->cppgoBase_inputMethodEvent((QInputMethodEvent*)arg1);
}
void* QScrollArea_inputMethodQuery_default(void* self, int arg1) {
	auto&& __r = ((QScrollArea*)self)->QScrollArea::inputMethodQuery((Qt::InputMethodQuery)(arg1));
	return new QVariant(__r);
}
void QScrollArea_timerEvent_default(void* self, void* event) {
	((QScrollArea__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QScrollArea_childEvent_default(void* self, void* event) {
	((QScrollArea__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QScrollArea_customEvent_default(void* self, void* event) {
	((QScrollArea__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QScrollArea_connectNotify_default(void* self, void* signal) {
	((QScrollArea__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QScrollArea_disconnectNotify_default(void* self, void* signal) {
	((QScrollArea__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QScrollArea_connect_windowTitleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QScrollArea*)self, &QScrollArea::windowTitleChanged, (QScrollArea*)self, [__id](const QString& title){
		QByteArray title_ba = title.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)title_ba.constData(), title_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QScrollArea);
	__conns_QScrollArea[__id] = __c;
}
void QScrollArea_connect_windowIconChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QScrollArea*)self, &QScrollArea::windowIconChanged, (QScrollArea*)self, [__id](const QIcon& icon){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&icon);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QScrollArea);
	__conns_QScrollArea[__id] = __c;
}
void QScrollArea_connect_windowIconTextChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QScrollArea*)self, &QScrollArea::windowIconTextChanged, (QScrollArea*)self, [__id](const QString& iconText){
		QByteArray iconText_ba = iconText.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)iconText_ba.constData(), iconText_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QScrollArea);
	__conns_QScrollArea[__id] = __c;
}
void QScrollArea_connect_customContextMenuRequested(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QScrollArea*)self, &QScrollArea::customContextMenuRequested, (QScrollArea*)self, [__id](const QPoint& pos){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&pos);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QScrollArea);
	__conns_QScrollArea[__id] = __c;
}
void QScrollArea_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QScrollArea*)self, &QScrollArea::destroyed, (QScrollArea*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QScrollArea);
	__conns_QScrollArea[__id] = __c;
}
void QScrollArea_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QScrollArea*)self, &QScrollArea::destroyed, (QScrollArea*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QScrollArea);
	__conns_QScrollArea[__id] = __c;
}
void QScrollArea_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QScrollArea*)self, &QScrollArea::objectNameChanged, (QScrollArea*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QScrollArea);
	__conns_QScrollArea[__id] = __c;
}
void QScrollArea__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QScrollArea);
		auto __it = __conns_QScrollArea.find(id);
		if (__it == __conns_QScrollArea.end()) return;
		__c = __it->second;
		__conns_QScrollArea.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QScrollArea__Go_new(long long id) { return new QScrollArea__Go(id); }
void QScrollArea_delete(void* self) { delete (QScrollArea*)self; }

}

#include <QtWidgets/qplaintextedit.h>
#include <QtGui/qabstracttextdocumentlayout.h>
#include <QtGui/qtextdocument.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtGui/qpainter.h>
#include <QtCore/qpoint.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qsize.h>
#include <QtCore/qrect.h>
#include <QtGui/qtextobject.h>
#include <QtGui/qtextformat.h>
#include <QtGui/qpaintdevice.h>
#include <QtCore/qobject.h>
#include <QtGui/qtextlayout.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qvariant.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qplaintextdocumentlayout.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

static std::unordered_map<long long, QMetaObject::Connection> __conns_QPlainTextDocumentLayout;
static std::mutex __connsMu_QPlainTextDocumentLayout;

extern "C" {

void* QPlainTextDocumentLayout_base(void* self) { return (void*)static_cast<QAbstractTextDocumentLayout*>((QPlainTextDocumentLayout*)self); }
void QPlainTextDocumentLayout_chain(void* self, void** out) {
	QPlainTextDocumentLayout* __p = (QPlainTextDocumentLayout*)self;
	QAbstractTextDocumentLayout* __b0 = static_cast<QAbstractTextDocumentLayout*>(__p); out[0] = (void*)__b0;
	QObject* __b1 = static_cast<QObject*>(__b0); out[1] = (void*)__b1;
}
void* QPlainTextDocumentLayout_fromQObject(void* o) { return (void*)qobject_cast<QPlainTextDocumentLayout*>((QObject*)o); }
void* QPlainTextDocumentLayout_new(void* document) {
	return new QPlainTextDocumentLayout((QTextDocument*)document);
}
char* QPlainTextDocumentLayout_tr(const char* s, long long* __out_len) {
	auto&& __r = QPlainTextDocumentLayout::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QPlainTextDocumentLayout_ensureBlockLayout(void* self, void* block) {
	((QPlainTextDocumentLayout*)self)->ensureBlockLayout(*(QTextBlock*)block);
}
void QPlainTextDocumentLayout_setCursorWidth(void* self, int width) {
	((QPlainTextDocumentLayout*)self)->setCursorWidth(width);
}
int QPlainTextDocumentLayout_cursorWidth(void* self) {
	auto&& __r = ((QPlainTextDocumentLayout*)self)->cursorWidth();
	return __r;
}
void QPlainTextDocumentLayout_requestUpdate(void* self) {
	((QPlainTextDocumentLayout*)self)->requestUpdate();
}
void* QPlainTextDocumentLayout_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QPlainTextDocumentLayout::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QPlainTextDocumentLayout_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QPlainTextDocumentLayout::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
char* QPlainTextDocumentLayout_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QPlainTextDocumentLayout::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QPlainTextDocumentLayout_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QPlainTextDocumentLayout::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QPlainTextDocumentLayout_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QPlainTextDocumentLayout::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QPlainTextDocumentLayout_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QPlainTextDocumentLayout::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void QPlainTextDocumentLayout_connect_update(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QPlainTextDocumentLayout*)self, &QPlainTextDocumentLayout::update, (QPlainTextDocumentLayout*)self, [__id](const QRectF& arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QPlainTextDocumentLayout);
	__conns_QPlainTextDocumentLayout[__id] = __c;
}
void QPlainTextDocumentLayout_connect_update2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QPlainTextDocumentLayout*)self, &QPlainTextDocumentLayout::update, (QPlainTextDocumentLayout*)self, [__id](const QRectF& arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QPlainTextDocumentLayout);
	__conns_QPlainTextDocumentLayout[__id] = __c;
}
void QPlainTextDocumentLayout_connect_updateBlock(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QPlainTextDocumentLayout*)self, &QPlainTextDocumentLayout::updateBlock, (QPlainTextDocumentLayout*)self, [__id](const QTextBlock& block){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&block);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QPlainTextDocumentLayout);
	__conns_QPlainTextDocumentLayout[__id] = __c;
}
void QPlainTextDocumentLayout_connect_documentSizeChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QPlainTextDocumentLayout*)self, &QPlainTextDocumentLayout::documentSizeChanged, (QPlainTextDocumentLayout*)self, [__id](const QSizeF& newSize){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&newSize);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QPlainTextDocumentLayout);
	__conns_QPlainTextDocumentLayout[__id] = __c;
}
void QPlainTextDocumentLayout_connect_pageCountChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QPlainTextDocumentLayout*)self, &QPlainTextDocumentLayout::pageCountChanged, (QPlainTextDocumentLayout*)self, [__id](int newPages){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, newPages);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QPlainTextDocumentLayout);
	__conns_QPlainTextDocumentLayout[__id] = __c;
}
void QPlainTextDocumentLayout_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QPlainTextDocumentLayout*)self, &QPlainTextDocumentLayout::destroyed, (QPlainTextDocumentLayout*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QPlainTextDocumentLayout);
	__conns_QPlainTextDocumentLayout[__id] = __c;
}
void QPlainTextDocumentLayout_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QPlainTextDocumentLayout*)self, &QPlainTextDocumentLayout::destroyed, (QPlainTextDocumentLayout*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QPlainTextDocumentLayout);
	__conns_QPlainTextDocumentLayout[__id] = __c;
}
void QPlainTextDocumentLayout_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QPlainTextDocumentLayout*)self, &QPlainTextDocumentLayout::objectNameChanged, (QPlainTextDocumentLayout*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QPlainTextDocumentLayout);
	__conns_QPlainTextDocumentLayout[__id] = __c;
}
void QPlainTextDocumentLayout__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QPlainTextDocumentLayout);
		auto __it = __conns_QPlainTextDocumentLayout.find(id);
		if (__it == __conns_QPlainTextDocumentLayout.end()) return;
		__c = __it->second;
		__conns_QPlainTextDocumentLayout.erase(__it);
	}
	QObject::disconnect(__c);
}
void QPlainTextDocumentLayout_delete(void* self) { delete (QPlainTextDocumentLayout*)self; }

}

#include <QtWidgets/qtablewidget.h>
#include "gen_qtablewidgetselectionrange.h"

extern "C" {

void* QTableWidgetSelectionRange_new() {
	return new QTableWidgetSelectionRange();
}
void* QTableWidgetSelectionRange_new2(int top, int left, int bottom, int right) {
	return new QTableWidgetSelectionRange(top, left, bottom, right);
}
int QTableWidgetSelectionRange_topRow(void* self) {
	auto&& __r = ((QTableWidgetSelectionRange*)self)->topRow();
	return __r;
}
int QTableWidgetSelectionRange_bottomRow(void* self) {
	auto&& __r = ((QTableWidgetSelectionRange*)self)->bottomRow();
	return __r;
}
int QTableWidgetSelectionRange_leftColumn(void* self) {
	auto&& __r = ((QTableWidgetSelectionRange*)self)->leftColumn();
	return __r;
}
int QTableWidgetSelectionRange_rightColumn(void* self) {
	auto&& __r = ((QTableWidgetSelectionRange*)self)->rightColumn();
	return __r;
}
int QTableWidgetSelectionRange_rowCount(void* self) {
	auto&& __r = ((QTableWidgetSelectionRange*)self)->rowCount();
	return __r;
}
int QTableWidgetSelectionRange_columnCount(void* self) {
	auto&& __r = ((QTableWidgetSelectionRange*)self)->columnCount();
	return __r;
}
int QTableWidgetSelectionRange_OperatorEqual(void* self, void* rhs) {
	auto&& __r = (*(QTableWidgetSelectionRange*)self) == (*(QTableWidgetSelectionRange*)rhs);
	return __r;
}
int QTableWidgetSelectionRange_OperatorNotEqual(void* self, void* rhs) {
	auto&& __r = (*(QTableWidgetSelectionRange*)self) != (*(QTableWidgetSelectionRange*)rhs);
	return __r;
}
void QTableWidgetSelectionRange_delete(void* self) { delete (QTableWidgetSelectionRange*)self; }

}

