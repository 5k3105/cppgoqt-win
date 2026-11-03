#include <QtCore/qcompare.h>
#include "gen_qt__partial_ordering.h"

extern "C" {

void* Qt__partial_ordering_new(void* arg1) {
	return new Qt::partial_ordering(*(Qt::partial_ordering*)arg1);
}
int Qt__partial_ordering_OperatorEqual(void* self, void* rhs) {
	auto&& __r = (*(Qt::partial_ordering*)self) == (*(Qt::partial_ordering*)rhs);
	return __r;
}
int Qt__partial_ordering_OperatorNotEqual(void* self, void* rhs) {
	auto&& __r = (*(Qt::partial_ordering*)self) != (*(Qt::partial_ordering*)rhs);
	return __r;
}
int Qt__partial_ordering_OperatorEqual2(void* self, void* rhs) {
	auto&& __r = (*(Qt::partial_ordering*)self) == (*(Qt::weak_ordering*)rhs);
	return __r;
}
int Qt__partial_ordering_OperatorNotEqual2(void* self, void* rhs) {
	auto&& __r = (*(Qt::partial_ordering*)self) != (*(Qt::weak_ordering*)rhs);
	return __r;
}
int Qt__partial_ordering_OperatorEqual3(void* self, void* rhs) {
	auto&& __r = (*(Qt::partial_ordering*)self) == (*(Qt::strong_ordering*)rhs);
	return __r;
}
int Qt__partial_ordering_OperatorNotEqual3(void* self, void* rhs) {
	auto&& __r = (*(Qt::partial_ordering*)self) != (*(Qt::strong_ordering*)rhs);
	return __r;
}
int Qt__partial_ordering_OperatorEqual4(void* self, void* rhs) {
	auto&& __r = (*(Qt::partial_ordering*)self) == (*(QPartialOrdering*)rhs);
	return __r;
}
int Qt__partial_ordering_OperatorNotEqual4(void* self, void* rhs) {
	auto&& __r = (*(Qt::partial_ordering*)self) != (*(QPartialOrdering*)rhs);
	return __r;
}
void Qt__partial_ordering_delete(void* self) { delete (Qt::partial_ordering*)self; }

}

#include <QtCore/qstringconverter_base.h>
#include "gen_qstringconverterbase.h"

extern "C" {


}

#include <QtCore/qabstracteventdispatcher.h>
#include <QtCore/qobject.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qthread.h>
#include <QtCore/qeventloop.h>
#include <QtCore/qsocketnotifier.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qabstractnativeeventfilter.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qvariant.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qabstracteventdispatcher.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QAbstractEventDispatcher__Go : public QAbstractEventDispatcher {
public:
	long long __goId;
	unsigned char __goOn[20] = {};
	template<typename... A> QAbstractEventDispatcher__Go(long long id, A&&... a) : QAbstractEventDispatcher(std::forward<A>(a)...), __goId(id) {}
	using QAbstractEventDispatcher::sender;
	using QAbstractEventDispatcher::senderSignalIndex;
	using QAbstractEventDispatcher::receivers;
	using QAbstractEventDispatcher::isSignalConnected;
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QAbstractEventDispatcher::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QAbstractEventDispatcher::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QAbstractEventDispatcher::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QAbstractEventDispatcher::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QAbstractEventDispatcher::disconnectNotify(__a0); }
	bool processEvents(QFlags<QEventLoop::ProcessEventsFlag> flags) override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, -2075143971860206337LL, (int)(flags), &__h); if (__h) return __r; }
		return bool{};
	}
	void registerSocketNotifier(QSocketNotifier* notifier) override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 849885779796930167LL, (void*)notifier, &__h); if (__h) return; }
		return;
	}
	void unregisterSocketNotifier(QSocketNotifier* notifier) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5438220604201114580LL, (void*)notifier, &__h); if (__h) return; }
		return;
	}
	void registerTimer(int timerId, qint64 interval, Qt::TimerType timerType, QObject* object) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_int_longlong_int_voidp) { ((void(*)(long long, long long, int, long long, int, void*, int*))cppgo_virt_void_int_longlong_int_voidp)(__goId, 5345355196729094449LL, timerId, interval, (int)(timerType), (void*)object, &__h); if (__h) return; }
		return;
	}
	bool unregisterTimer(int timerId) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, -85727209796773320LL, timerId, &__h); if (__h) return __r; }
		return bool{};
	}
	bool unregisterTimers(QObject* object) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 8418593569725967729LL, (void*)object, &__h); if (__h) return __r; }
		return bool{};
	}
	QList<QAbstractEventDispatcher::TimerInfo> registeredTimers(QObject* object) const override {
		int __h = 0;
		if (__goOn[6] && cppgo_virtp_voidp_voidp) {
			long long __len = 0;
			void* __a = ((void*(*)(long long, long long, void*, long long*, int*))cppgo_virtp_voidp_voidp)(__goId, 776599310866639255LL, (void*)object, &__len, &__h);
			if (__h) {
				QList<QAbstractEventDispatcher::TimerInfo> __r; __r.reserve(__len);
				void** __aa = (void**)__a;
				for (long long __i = 0; __i < __len; __i++) __r.append(*(QAbstractEventDispatcher::TimerInfo*)__aa[__i]);
				if (__a) free(__a);
				return __r;
			}
		}
		return QList<QAbstractEventDispatcher::TimerInfo>{};
	}
	int remainingTime(int timerId) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, -7776574712936801734LL, timerId, &__h); if (__h) return __r; }
		return int{};
	}
	void wakeUp() override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, -6649993720439713720LL, &__h); if (__h) return; }
		return;
	}
	void interrupt() override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, -3579312661142054576LL, &__h); if (__h) return; }
		return;
	}
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QAbstractEventDispatcher::metaObject();
	}
	void startingUp() override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, -4418351852319543908LL, &__h); if (__h) return; }
		QAbstractEventDispatcher::startingUp();
	}
	void closingDown() override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, -8213485174827668208LL, &__h); if (__h) return; }
		QAbstractEventDispatcher::closingDown();
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QAbstractEventDispatcher::event(event);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QAbstractEventDispatcher::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QAbstractEventDispatcher::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QAbstractEventDispatcher::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QAbstractEventDispatcher::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QAbstractEventDispatcher::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QAbstractEventDispatcher::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QAbstractEventDispatcher;
static std::mutex __connsMu_QAbstractEventDispatcher;

extern "C" {

void* QAbstractEventDispatcher_base(void* self) { return (void*)static_cast<QObject*>((QAbstractEventDispatcher*)self); }
void QAbstractEventDispatcher__set_override(void* self, int idx) { ((QAbstractEventDispatcher__Go*)self)->__goOn[idx] = 1; }
void* QAbstractEventDispatcher_fromQObject(void* o) { return (void*)qobject_cast<QAbstractEventDispatcher*>((QObject*)o); }
char* QAbstractEventDispatcher_tr(const char* s, long long* __out_len) {
	auto&& __r = QAbstractEventDispatcher::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QAbstractEventDispatcher_instance() {
	auto&& __r = QAbstractEventDispatcher::instance();
	return (void*)(__r);
}
int QAbstractEventDispatcher_processEvents(void* self, int flags) {
	auto&& __r = ((QAbstractEventDispatcher*)self)->processEvents(QFlags<QEventLoop::ProcessEventsFlag>(flags));
	return __r;
}
void QAbstractEventDispatcher_registerSocketNotifier(void* self, void* notifier) {
	((QAbstractEventDispatcher*)self)->registerSocketNotifier((QSocketNotifier*)notifier);
}
void QAbstractEventDispatcher_unregisterSocketNotifier(void* self, void* notifier) {
	((QAbstractEventDispatcher*)self)->unregisterSocketNotifier((QSocketNotifier*)notifier);
}
int QAbstractEventDispatcher_registerTimer(void* self, long long interval, int timerType, void* object) {
	auto&& __r = ((QAbstractEventDispatcher*)self)->registerTimer(interval, (Qt::TimerType)(timerType), (QObject*)object);
	return __r;
}
void QAbstractEventDispatcher_registerTimer2(void* self, int timerId, long long interval, int timerType, void* object) {
	((QAbstractEventDispatcher*)self)->registerTimer(timerId, interval, (Qt::TimerType)(timerType), (QObject*)object);
}
int QAbstractEventDispatcher_unregisterTimer(void* self, int timerId) {
	auto&& __r = ((QAbstractEventDispatcher*)self)->unregisterTimer(timerId);
	return __r;
}
int QAbstractEventDispatcher_unregisterTimers(void* self, void* object) {
	auto&& __r = ((QAbstractEventDispatcher*)self)->unregisterTimers((QObject*)object);
	return __r;
}
void** QAbstractEventDispatcher_registeredTimers(void* self, void* object, long long* __out_len) {
	auto&& __r = ((QAbstractEventDispatcher*)self)->registeredTimers((QObject*)object);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QAbstractEventDispatcher::TimerInfo(__e);
	return __arr;
}
int QAbstractEventDispatcher_remainingTime(void* self, int timerId) {
	auto&& __r = ((QAbstractEventDispatcher*)self)->remainingTime(timerId);
	return __r;
}
int QAbstractEventDispatcher_unregisterTimer2(void* self, int timerId) {
	auto&& __r = ((QAbstractEventDispatcher*)self)->unregisterTimer((Qt::TimerId)(timerId));
	return __r;
}
void** QAbstractEventDispatcher_timersForObject(void* self, void* object, long long* __out_len) {
	auto&& __r = ((QAbstractEventDispatcher*)self)->timersForObject((QObject*)object);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QAbstractEventDispatcher::TimerInfoV2(__e);
	return __arr;
}
void QAbstractEventDispatcher_wakeUp(void* self) {
	((QAbstractEventDispatcher*)self)->wakeUp();
}
void QAbstractEventDispatcher_interrupt(void* self) {
	((QAbstractEventDispatcher*)self)->interrupt();
}
void QAbstractEventDispatcher_startingUp(void* self) {
	((QAbstractEventDispatcher*)self)->startingUp();
}
void QAbstractEventDispatcher_closingDown(void* self) {
	((QAbstractEventDispatcher*)self)->closingDown();
}
void QAbstractEventDispatcher_installNativeEventFilter(void* self, void* filterObj) {
	((QAbstractEventDispatcher*)self)->installNativeEventFilter((QAbstractNativeEventFilter*)filterObj);
}
void QAbstractEventDispatcher_removeNativeEventFilter(void* self, void* filterObj) {
	((QAbstractEventDispatcher*)self)->removeNativeEventFilter((QAbstractNativeEventFilter*)filterObj);
}
void QAbstractEventDispatcher_aboutToBlock(void* self) {
	((QAbstractEventDispatcher*)self)->aboutToBlock();
}
void QAbstractEventDispatcher_awake(void* self) {
	((QAbstractEventDispatcher*)self)->awake();
}
void* QAbstractEventDispatcher_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QAbstractEventDispatcher::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QAbstractEventDispatcher_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QAbstractEventDispatcher::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QAbstractEventDispatcher_sender(void* self) {
	auto&& __r = ((QAbstractEventDispatcher__Go*)self)->sender();
	return (void*)(__r);
}
int QAbstractEventDispatcher_senderSignalIndex(void* self) {
	auto&& __r = ((QAbstractEventDispatcher__Go*)self)->senderSignalIndex();
	return __r;
}
int QAbstractEventDispatcher_receivers(void* self, const char* signal) {
	auto&& __r = ((QAbstractEventDispatcher__Go*)self)->receivers(signal);
	return __r;
}
int QAbstractEventDispatcher_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QAbstractEventDispatcher__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QAbstractEventDispatcher_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QAbstractEventDispatcher::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QAbstractEventDispatcher_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QAbstractEventDispatcher::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QAbstractEventDispatcher_instance2(void* thread) {
	auto&& __r = QAbstractEventDispatcher::instance((QThread*)thread);
	return (void*)(__r);
}
void* QAbstractEventDispatcher_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QAbstractEventDispatcher::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QAbstractEventDispatcher_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QAbstractEventDispatcher::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void QAbstractEventDispatcher_startingUp_default(void* self) {
	((QAbstractEventDispatcher*)self)->QAbstractEventDispatcher::startingUp();
}
void QAbstractEventDispatcher_closingDown_default(void* self) {
	((QAbstractEventDispatcher*)self)->QAbstractEventDispatcher::closingDown();
}
void* QAbstractEventDispatcher_metaObject_default(void* self) {
	auto&& __r = ((QAbstractEventDispatcher*)self)->QAbstractEventDispatcher::metaObject();
	return (void*)(__r);
}
int QAbstractEventDispatcher_event_default(void* self, void* event) {
	auto&& __r = ((QAbstractEventDispatcher*)self)->QAbstractEventDispatcher::event((QEvent*)event);
	return __r;
}
int QAbstractEventDispatcher_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QAbstractEventDispatcher*)self)->QAbstractEventDispatcher::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QAbstractEventDispatcher_timerEvent_default(void* self, void* event) {
	((QAbstractEventDispatcher__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QAbstractEventDispatcher_childEvent_default(void* self, void* event) {
	((QAbstractEventDispatcher__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QAbstractEventDispatcher_customEvent_default(void* self, void* event) {
	((QAbstractEventDispatcher__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QAbstractEventDispatcher_connectNotify_default(void* self, void* signal) {
	((QAbstractEventDispatcher__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QAbstractEventDispatcher_disconnectNotify_default(void* self, void* signal) {
	((QAbstractEventDispatcher__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QAbstractEventDispatcher_connect_aboutToBlock(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractEventDispatcher*)self, &QAbstractEventDispatcher::aboutToBlock, (QAbstractEventDispatcher*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractEventDispatcher);
	__conns_QAbstractEventDispatcher[__id] = __c;
}
void QAbstractEventDispatcher_connect_awake(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractEventDispatcher*)self, &QAbstractEventDispatcher::awake, (QAbstractEventDispatcher*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractEventDispatcher);
	__conns_QAbstractEventDispatcher[__id] = __c;
}
void QAbstractEventDispatcher_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractEventDispatcher*)self, &QAbstractEventDispatcher::destroyed, (QAbstractEventDispatcher*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractEventDispatcher);
	__conns_QAbstractEventDispatcher[__id] = __c;
}
void QAbstractEventDispatcher_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractEventDispatcher*)self, &QAbstractEventDispatcher::destroyed, (QAbstractEventDispatcher*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractEventDispatcher);
	__conns_QAbstractEventDispatcher[__id] = __c;
}
void QAbstractEventDispatcher_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractEventDispatcher*)self, &QAbstractEventDispatcher::objectNameChanged, (QAbstractEventDispatcher*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractEventDispatcher);
	__conns_QAbstractEventDispatcher[__id] = __c;
}
void QAbstractEventDispatcher__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QAbstractEventDispatcher);
		auto __it = __conns_QAbstractEventDispatcher.find(id);
		if (__it == __conns_QAbstractEventDispatcher.end()) return;
		__c = __it->second;
		__conns_QAbstractEventDispatcher.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QAbstractEventDispatcher__Go_new(long long id, void* parent) { return new QAbstractEventDispatcher__Go(id, (QObject*)parent); }
void QAbstractEventDispatcher_delete(void* self) { delete (QAbstractEventDispatcher*)self; }

}

#include <QtCore/qabstractproxymodel.h>
#include <QtCore/qabstractitemmodel.h>
#include <QtCore/qobject.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qitemselectionmodel.h>
#include <QtCore/qvariant.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qsize.h>
#include <QtCore/qmimedata.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qdatastream.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qabstractproxymodel.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QAbstractProxyModel__Go : public QAbstractProxyModel {
public:
	long long __goId;
	unsigned char __goOn[47] = {};
	template<typename... A> QAbstractProxyModel__Go(long long id, A&&... a) : QAbstractProxyModel(std::forward<A>(a)...), __goId(id) {}
	using QAbstractProxyModel::createIndex;
	using QAbstractProxyModel::encodeData;
	using QAbstractProxyModel::decodeData;
	using QAbstractProxyModel::beginInsertRows;
	using QAbstractProxyModel::endInsertRows;
	using QAbstractProxyModel::beginRemoveRows;
	using QAbstractProxyModel::endRemoveRows;
	using QAbstractProxyModel::beginMoveRows;
	using QAbstractProxyModel::endMoveRows;
	using QAbstractProxyModel::beginInsertColumns;
	using QAbstractProxyModel::endInsertColumns;
	using QAbstractProxyModel::beginRemoveColumns;
	using QAbstractProxyModel::endRemoveColumns;
	using QAbstractProxyModel::beginMoveColumns;
	using QAbstractProxyModel::endMoveColumns;
	using QAbstractProxyModel::beginResetModel;
	using QAbstractProxyModel::endResetModel;
	using QAbstractProxyModel::changePersistentIndex;
	using QAbstractProxyModel::changePersistentIndexList;
	using QAbstractProxyModel::persistentIndexList;
	using QAbstractProxyModel::sender;
	using QAbstractProxyModel::senderSignalIndex;
	using QAbstractProxyModel::receivers;
	using QAbstractProxyModel::isSignalConnected;
	void cppgoBase_resetInternalData() { QAbstractProxyModel::resetInternalData(); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QAbstractProxyModel::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QAbstractProxyModel::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QAbstractProxyModel::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QAbstractProxyModel::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QAbstractProxyModel::disconnectNotify(__a0); }
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
	QModelIndex mapToSource(const QModelIndex& proxyIndex) const override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, -2333321292016234891LL, (void*)&proxyIndex, &__h); if (__h) return *(QModelIndex*)__r; }
		return QModelIndex{};
	}
	QModelIndex mapFromSource(const QModelIndex& sourceIndex) const override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, -1717009778236344012LL, (void*)&sourceIndex, &__h); if (__h) return *(QModelIndex*)__r; }
		return QModelIndex{};
	}
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QAbstractProxyModel::metaObject();
	}
	void setSourceModel(QAbstractItemModel* sourceModel) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7818871925547456643LL, (void*)sourceModel, &__h); if (__h) return; }
		QAbstractProxyModel::setSourceModel(sourceModel);
	}
	QItemSelection mapSelectionToSource(const QItemSelection& selection) const override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 2788628866280405837LL, (void*)&selection, &__h); if (__h) return *(QItemSelection*)__r; }
		return QAbstractProxyModel::mapSelectionToSource(selection);
	}
	QItemSelection mapSelectionFromSource(const QItemSelection& selection) const override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 9067522407071207004LL, (void*)&selection, &__h); if (__h) return *(QItemSelection*)__r; }
		return QAbstractProxyModel::mapSelectionFromSource(selection);
	}
	bool submit() override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -3788331537442685179LL, &__h); if (__h) return __r; }
		return QAbstractProxyModel::submit();
	}
	void revert() override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, -1522951607536117149LL, &__h); if (__h) return; }
		QAbstractProxyModel::revert();
	}
	QVariant data(const QModelIndex& proxyIndex, int role) const override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_voidp_voidp_int) { void* __r = ((void*(*)(long long, long long, void*, int, int*))cppgo_virt_voidp_voidp_int)(__goId, -8837375942098466299LL, (void*)&proxyIndex, role, &__h); if (__h) return *(QVariant*)__r; }
		return QAbstractProxyModel::data(proxyIndex, role);
	}
	QVariant headerData(int section, Qt::Orientation orientation, int role) const override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_voidp_int_int_int) { void* __r = ((void*(*)(long long, long long, int, int, int, int*))cppgo_virt_voidp_int_int_int)(__goId, -3297140021896693688LL, section, (int)(orientation), role, &__h); if (__h) return *(QVariant*)__r; }
		return QAbstractProxyModel::headerData(section, orientation, role);
	}
	QFlags<Qt::ItemFlag> flags(const QModelIndex& index) const override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1703382833545304812LL, (void*)&index, &__h); if (__h) return QFlags<Qt::ItemFlag>(__r); }
		return QAbstractProxyModel::flags(index);
	}
	bool setData(const QModelIndex& index, const QVariant& value, int role) override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_int_voidp_voidp_int) { int __r = ((int(*)(long long, long long, void*, void*, int, int*))cppgo_virt_int_voidp_voidp_int)(__goId, 35191454575202219LL, (void*)&index, (void*)&value, role, &__h); if (__h) return __r; }
		return QAbstractProxyModel::setData(index, value, role);
	}
	bool setHeaderData(int section, Qt::Orientation orientation, const QVariant& value, int role) override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_int_int_int_voidp_int) { int __r = ((int(*)(long long, long long, int, int, void*, int, int*))cppgo_virt_int_int_int_voidp_int)(__goId, 4852200306590054850LL, section, (int)(orientation), (void*)&value, role, &__h); if (__h) return __r; }
		return QAbstractProxyModel::setHeaderData(section, orientation, value, role);
	}
	bool clearItemData(const QModelIndex& index) override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 8914190500520542409LL, (void*)&index, &__h); if (__h) return __r; }
		return QAbstractProxyModel::clearItemData(index);
	}
	QModelIndex buddy(const QModelIndex& index) const override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 6112860112248786499LL, (void*)&index, &__h); if (__h) return *(QModelIndex*)__r; }
		return QAbstractProxyModel::buddy(index);
	}
	bool canFetchMore(const QModelIndex& parent) const override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 8639262066473964150LL, (void*)&parent, &__h); if (__h) return __r; }
		return QAbstractProxyModel::canFetchMore(parent);
	}
	void fetchMore(const QModelIndex& parent) override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3553294593611130806LL, (void*)&parent, &__h); if (__h) return; }
		QAbstractProxyModel::fetchMore(parent);
	}
	void sort(int column, Qt::SortOrder order) override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_void_int_int) { ((void(*)(long long, long long, int, int, int*))cppgo_virt_void_int_int)(__goId, 6912441694453382161LL, column, (int)(order), &__h); if (__h) return; }
		QAbstractProxyModel::sort(column, order);
	}
	QSize span(const QModelIndex& index) const override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, -8395342916587040543LL, (void*)&index, &__h); if (__h) return *(QSize*)__r; }
		return QAbstractProxyModel::span(index);
	}
	bool hasChildren(const QModelIndex& parent) const override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 3369452259144514380LL, (void*)&parent, &__h); if (__h) return __r; }
		return QAbstractProxyModel::hasChildren(parent);
	}
	QModelIndex sibling(int row, int column, const QModelIndex& idx) const override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_voidp_int_int_voidp) { void* __r = ((void*(*)(long long, long long, int, int, void*, int*))cppgo_virt_voidp_int_int_voidp)(__goId, 397329346813782521LL, row, column, (void*)&idx, &__h); if (__h) return *(QModelIndex*)__r; }
		return QAbstractProxyModel::sibling(row, column, idx);
	}
	QMimeData* mimeData(const QList<QModelIndex>& indexes) const override {
		int __h = 0;
		std::vector<void*> indexes_v; indexes_v.reserve(indexes.size());
		for (long long __i = 0; __i < indexes.size(); __i++) indexes_v.push_back((void*)&indexes[__i]);
		if (__goOn[25] && cppgo_virt_voidp_voidpp_longlong) { void* __r = ((void*(*)(long long, long long, void**, long long, int*))cppgo_virt_voidp_voidpp_longlong)(__goId, 5175297091024974749LL, indexes_v.data(), (long long)indexes.size(), &__h); if (__h) return (QMimeData*)__r; }
		return QAbstractProxyModel::mimeData(indexes);
	}
	bool canDropMimeData(const QMimeData* data, Qt::DropAction action, int row, int column, const QModelIndex& parent) const override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_int_voidp_int_int_int_voidp) { int __r = ((int(*)(long long, long long, void*, int, int, int, void*, int*))cppgo_virt_int_voidp_int_int_int_voidp)(__goId, 731539533644753052LL, (void*)data, (int)(action), row, column, (void*)&parent, &__h); if (__h) return __r; }
		return QAbstractProxyModel::canDropMimeData(data, action, row, column, parent);
	}
	bool dropMimeData(const QMimeData* data, Qt::DropAction action, int row, int column, const QModelIndex& parent) override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_int_voidp_int_int_int_voidp) { int __r = ((int(*)(long long, long long, void*, int, int, int, void*, int*))cppgo_virt_int_voidp_int_int_int_voidp)(__goId, -9049994684305682212LL, (void*)data, (int)(action), row, column, (void*)&parent, &__h); if (__h) return __r; }
		return QAbstractProxyModel::dropMimeData(data, action, row, column, parent);
	}
	QList<QString> mimeTypes() const override {
		int __h = 0;
		if (__goOn[28] && cppgo_virtp_voidp) {
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
		return QAbstractProxyModel::mimeTypes();
	}
	QFlags<Qt::DropAction> supportedDragActions() const override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, 3421601386088570762LL, &__h); if (__h) return QFlags<Qt::DropAction>(__r); }
		return QAbstractProxyModel::supportedDragActions();
	}
	QFlags<Qt::DropAction> supportedDropActions() const override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -7784642176186821205LL, &__h); if (__h) return QFlags<Qt::DropAction>(__r); }
		return QAbstractProxyModel::supportedDropActions();
	}
	bool insertRows(int row, int count, const QModelIndex& parent) override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_int_int_int_voidp) { int __r = ((int(*)(long long, long long, int, int, void*, int*))cppgo_virt_int_int_int_voidp)(__goId, -2775729825962395603LL, row, count, (void*)&parent, &__h); if (__h) return __r; }
		return QAbstractProxyModel::insertRows(row, count, parent);
	}
	bool insertColumns(int column, int count, const QModelIndex& parent) override {
		int __h = 0;
		if (__goOn[32] && cppgo_virt_int_int_int_voidp) { int __r = ((int(*)(long long, long long, int, int, void*, int*))cppgo_virt_int_int_int_voidp)(__goId, -8973293846230792237LL, column, count, (void*)&parent, &__h); if (__h) return __r; }
		return QAbstractProxyModel::insertColumns(column, count, parent);
	}
	bool removeRows(int row, int count, const QModelIndex& parent) override {
		int __h = 0;
		if (__goOn[33] && cppgo_virt_int_int_int_voidp) { int __r = ((int(*)(long long, long long, int, int, void*, int*))cppgo_virt_int_int_int_voidp)(__goId, 5104072549182222656LL, row, count, (void*)&parent, &__h); if (__h) return __r; }
		return QAbstractProxyModel::removeRows(row, count, parent);
	}
	bool removeColumns(int column, int count, const QModelIndex& parent) override {
		int __h = 0;
		if (__goOn[34] && cppgo_virt_int_int_int_voidp) { int __r = ((int(*)(long long, long long, int, int, void*, int*))cppgo_virt_int_int_int_voidp)(__goId, 2078548737056540420LL, column, count, (void*)&parent, &__h); if (__h) return __r; }
		return QAbstractProxyModel::removeColumns(column, count, parent);
	}
	bool moveRows(const QModelIndex& sourceParent, int sourceRow, int count, const QModelIndex& destinationParent, int destinationChild) override {
		int __h = 0;
		if (__goOn[35] && cppgo_virt_int_voidp_int_int_voidp_int) { int __r = ((int(*)(long long, long long, void*, int, int, void*, int, int*))cppgo_virt_int_voidp_int_int_voidp_int)(__goId, -6074819356947405415LL, (void*)&sourceParent, sourceRow, count, (void*)&destinationParent, destinationChild, &__h); if (__h) return __r; }
		return QAbstractProxyModel::moveRows(sourceParent, sourceRow, count, destinationParent, destinationChild);
	}
	bool moveColumns(const QModelIndex& sourceParent, int sourceColumn, int count, const QModelIndex& destinationParent, int destinationChild) override {
		int __h = 0;
		if (__goOn[36] && cppgo_virt_int_voidp_int_int_voidp_int) { int __r = ((int(*)(long long, long long, void*, int, int, void*, int, int*))cppgo_virt_int_voidp_int_int_voidp_int)(__goId, -6321007591388309905LL, (void*)&sourceParent, sourceColumn, count, (void*)&destinationParent, destinationChild, &__h); if (__h) return __r; }
		return QAbstractProxyModel::moveColumns(sourceParent, sourceColumn, count, destinationParent, destinationChild);
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
		return QAbstractProxyModel::match(start, role, value, hits, flags);
	}
	void multiData(const QModelIndex& index, QModelRoleDataSpan roleDataSpan) const override {
		int __h = 0;
		if (__goOn[38] && cppgo_virt_void_voidp_voidp) { ((void(*)(long long, long long, void*, void*, int*))cppgo_virt_void_voidp_voidp)(__goId, 731311060944096642LL, (void*)&index, (void*)&roleDataSpan, &__h); if (__h) return; }
		QAbstractProxyModel::multiData(index, roleDataSpan);
	}
	void resetInternalData() override {
		int __h = 0;
		if (__goOn[39] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, -4510675770784051183LL, &__h); if (__h) return; }
		QAbstractProxyModel::resetInternalData();
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[40] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QAbstractProxyModel::event(event);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[41] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QAbstractProxyModel::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[42] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QAbstractProxyModel::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[43] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QAbstractProxyModel::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[44] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QAbstractProxyModel::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[45] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QAbstractProxyModel::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[46] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QAbstractProxyModel::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QAbstractProxyModel;
static std::mutex __connsMu_QAbstractProxyModel;

extern "C" {

void* QAbstractProxyModel_base(void* self) { return (void*)static_cast<QAbstractItemModel*>((QAbstractProxyModel*)self); }
void QAbstractProxyModel_chain(void* self, void** out) {
	QAbstractProxyModel* __p = (QAbstractProxyModel*)self;
	QAbstractItemModel* __b0 = static_cast<QAbstractItemModel*>(__p); out[0] = (void*)__b0;
	QObject* __b1 = static_cast<QObject*>(__b0); out[1] = (void*)__b1;
}
void QAbstractProxyModel__set_override(void* self, int idx) { ((QAbstractProxyModel__Go*)self)->__goOn[idx] = 1; }
void* QAbstractProxyModel_fromQObject(void* o) { return (void*)qobject_cast<QAbstractProxyModel*>((QObject*)o); }
char* QAbstractProxyModel_tr(const char* s, long long* __out_len) {
	auto&& __r = QAbstractProxyModel::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QAbstractProxyModel_setSourceModel(void* self, void* sourceModel) {
	((QAbstractProxyModel*)self)->setSourceModel((QAbstractItemModel*)sourceModel);
}
void* QAbstractProxyModel_sourceModel(void* self) {
	auto&& __r = ((QAbstractProxyModel*)self)->sourceModel();
	return (void*)(__r);
}
void* QAbstractProxyModel_mapToSource(void* self, void* proxyIndex) {
	auto&& __r = ((QAbstractProxyModel*)self)->mapToSource(*(QModelIndex*)proxyIndex);
	return new QModelIndex(__r);
}
void* QAbstractProxyModel_mapFromSource(void* self, void* sourceIndex) {
	auto&& __r = ((QAbstractProxyModel*)self)->mapFromSource(*(QModelIndex*)sourceIndex);
	return new QModelIndex(__r);
}
void* QAbstractProxyModel_mapSelectionToSource(void* self, void* selection) {
	auto&& __r = ((QAbstractProxyModel*)self)->mapSelectionToSource(*(QItemSelection*)selection);
	return new QItemSelection(__r);
}
void* QAbstractProxyModel_mapSelectionFromSource(void* self, void* selection) {
	auto&& __r = ((QAbstractProxyModel*)self)->mapSelectionFromSource(*(QItemSelection*)selection);
	return new QItemSelection(__r);
}
void* QAbstractProxyModel_createIndex(void* self, int row, int column) {
	auto&& __r = ((QAbstractProxyModel__Go*)self)->createIndex(row, column);
	return new QModelIndex(__r);
}
void* QAbstractProxyModel_createIndex2(void* self, int row, int column, unsigned long long id) {
	auto&& __r = ((QAbstractProxyModel__Go*)self)->createIndex(row, column, id);
	return new QModelIndex(__r);
}
void QAbstractProxyModel_encodeData(void* self, void** indexes, long long indexes_len, void* stream) {
	QList<QModelIndex> indexes_list; indexes_list.reserve(indexes_len);
	for (long long __i = 0; __i < indexes_len; __i++) indexes_list.append(*(QModelIndex*)indexes[__i]);
	((QAbstractProxyModel__Go*)self)->encodeData(indexes_list, *(QDataStream*)stream);
}
int QAbstractProxyModel_decodeData(void* self, int row, int column, void* parent, void* stream) {
	auto&& __r = ((QAbstractProxyModel__Go*)self)->decodeData(row, column, *(QModelIndex*)parent, *(QDataStream*)stream);
	return __r;
}
void QAbstractProxyModel_beginInsertRows(void* self, void* parent, int first, int last) {
	((QAbstractProxyModel__Go*)self)->beginInsertRows(*(QModelIndex*)parent, first, last);
}
void QAbstractProxyModel_endInsertRows(void* self) {
	((QAbstractProxyModel__Go*)self)->endInsertRows();
}
void QAbstractProxyModel_beginRemoveRows(void* self, void* parent, int first, int last) {
	((QAbstractProxyModel__Go*)self)->beginRemoveRows(*(QModelIndex*)parent, first, last);
}
void QAbstractProxyModel_endRemoveRows(void* self) {
	((QAbstractProxyModel__Go*)self)->endRemoveRows();
}
int QAbstractProxyModel_beginMoveRows(void* self, void* sourceParent, int sourceFirst, int sourceLast, void* destinationParent, int destinationRow) {
	auto&& __r = ((QAbstractProxyModel__Go*)self)->beginMoveRows(*(QModelIndex*)sourceParent, sourceFirst, sourceLast, *(QModelIndex*)destinationParent, destinationRow);
	return __r;
}
void QAbstractProxyModel_endMoveRows(void* self) {
	((QAbstractProxyModel__Go*)self)->endMoveRows();
}
void QAbstractProxyModel_beginInsertColumns(void* self, void* parent, int first, int last) {
	((QAbstractProxyModel__Go*)self)->beginInsertColumns(*(QModelIndex*)parent, first, last);
}
void QAbstractProxyModel_endInsertColumns(void* self) {
	((QAbstractProxyModel__Go*)self)->endInsertColumns();
}
void QAbstractProxyModel_beginRemoveColumns(void* self, void* parent, int first, int last) {
	((QAbstractProxyModel__Go*)self)->beginRemoveColumns(*(QModelIndex*)parent, first, last);
}
void QAbstractProxyModel_endRemoveColumns(void* self) {
	((QAbstractProxyModel__Go*)self)->endRemoveColumns();
}
int QAbstractProxyModel_beginMoveColumns(void* self, void* sourceParent, int sourceFirst, int sourceLast, void* destinationParent, int destinationColumn) {
	auto&& __r = ((QAbstractProxyModel__Go*)self)->beginMoveColumns(*(QModelIndex*)sourceParent, sourceFirst, sourceLast, *(QModelIndex*)destinationParent, destinationColumn);
	return __r;
}
void QAbstractProxyModel_endMoveColumns(void* self) {
	((QAbstractProxyModel__Go*)self)->endMoveColumns();
}
void QAbstractProxyModel_beginResetModel(void* self) {
	((QAbstractProxyModel__Go*)self)->beginResetModel();
}
void QAbstractProxyModel_endResetModel(void* self) {
	((QAbstractProxyModel__Go*)self)->endResetModel();
}
void QAbstractProxyModel_changePersistentIndex(void* self, void* from, void* to) {
	((QAbstractProxyModel__Go*)self)->changePersistentIndex(*(QModelIndex*)from, *(QModelIndex*)to);
}
void QAbstractProxyModel_changePersistentIndexList(void* self, void** from, long long from_len, void** to, long long to_len) {
	QList<QModelIndex> from_list; from_list.reserve(from_len);
	for (long long __i = 0; __i < from_len; __i++) from_list.append(*(QModelIndex*)from[__i]);
	QList<QModelIndex> to_list; to_list.reserve(to_len);
	for (long long __i = 0; __i < to_len; __i++) to_list.append(*(QModelIndex*)to[__i]);
	((QAbstractProxyModel__Go*)self)->changePersistentIndexList(from_list, to_list);
}
void** QAbstractProxyModel_persistentIndexList(void* self, long long* __out_len) {
	auto&& __r = ((QAbstractProxyModel__Go*)self)->persistentIndexList();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QModelIndex(__e);
	return __arr;
}
void* QAbstractProxyModel_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QAbstractProxyModel::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QAbstractProxyModel_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QAbstractProxyModel::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QAbstractProxyModel_sender(void* self) {
	auto&& __r = ((QAbstractProxyModel__Go*)self)->sender();
	return (void*)(__r);
}
int QAbstractProxyModel_senderSignalIndex(void* self) {
	auto&& __r = ((QAbstractProxyModel__Go*)self)->senderSignalIndex();
	return __r;
}
int QAbstractProxyModel_receivers(void* self, const char* signal) {
	auto&& __r = ((QAbstractProxyModel__Go*)self)->receivers(signal);
	return __r;
}
int QAbstractProxyModel_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QAbstractProxyModel__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QAbstractProxyModel_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QAbstractProxyModel::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QAbstractProxyModel_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QAbstractProxyModel::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QAbstractProxyModel_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QAbstractProxyModel::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QAbstractProxyModel_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QAbstractProxyModel::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void QAbstractProxyModel_setSourceModel_default(void* self, void* sourceModel) {
	((QAbstractProxyModel*)self)->QAbstractProxyModel::setSourceModel((QAbstractItemModel*)sourceModel);
}
void* QAbstractProxyModel_mapSelectionToSource_default(void* self, void* selection) {
	auto&& __r = ((QAbstractProxyModel*)self)->QAbstractProxyModel::mapSelectionToSource(*(QItemSelection*)selection);
	return new QItemSelection(__r);
}
void* QAbstractProxyModel_mapSelectionFromSource_default(void* self, void* selection) {
	auto&& __r = ((QAbstractProxyModel*)self)->QAbstractProxyModel::mapSelectionFromSource(*(QItemSelection*)selection);
	return new QItemSelection(__r);
}
void* QAbstractProxyModel_metaObject_default(void* self) {
	auto&& __r = ((QAbstractProxyModel*)self)->QAbstractProxyModel::metaObject();
	return (void*)(__r);
}
int QAbstractProxyModel_submit_default(void* self) {
	auto&& __r = ((QAbstractProxyModel*)self)->QAbstractProxyModel::submit();
	return __r;
}
void QAbstractProxyModel_revert_default(void* self) {
	((QAbstractProxyModel*)self)->QAbstractProxyModel::revert();
}
void* QAbstractProxyModel_data_default(void* self, void* proxyIndex, int role) {
	auto&& __r = ((QAbstractProxyModel*)self)->QAbstractProxyModel::data(*(QModelIndex*)proxyIndex, role);
	return new QVariant(__r);
}
void* QAbstractProxyModel_headerData_default(void* self, int section, int orientation, int role) {
	auto&& __r = ((QAbstractProxyModel*)self)->QAbstractProxyModel::headerData(section, (Qt::Orientation)(orientation), role);
	return new QVariant(__r);
}
int QAbstractProxyModel_flags_default(void* self, void* index) {
	auto&& __r = ((QAbstractProxyModel*)self)->QAbstractProxyModel::flags(*(QModelIndex*)index);
	return (int)(__r);
}
int QAbstractProxyModel_setData_default(void* self, void* index, void* value, int role) {
	auto&& __r = ((QAbstractProxyModel*)self)->QAbstractProxyModel::setData(*(QModelIndex*)index, *(QVariant*)value, role);
	return __r;
}
int QAbstractProxyModel_setHeaderData_default(void* self, int section, int orientation, void* value, int role) {
	auto&& __r = ((QAbstractProxyModel*)self)->QAbstractProxyModel::setHeaderData(section, (Qt::Orientation)(orientation), *(QVariant*)value, role);
	return __r;
}
int QAbstractProxyModel_clearItemData_default(void* self, void* index) {
	auto&& __r = ((QAbstractProxyModel*)self)->QAbstractProxyModel::clearItemData(*(QModelIndex*)index);
	return __r;
}
void* QAbstractProxyModel_buddy_default(void* self, void* index) {
	auto&& __r = ((QAbstractProxyModel*)self)->QAbstractProxyModel::buddy(*(QModelIndex*)index);
	return new QModelIndex(__r);
}
int QAbstractProxyModel_canFetchMore_default(void* self, void* parent) {
	auto&& __r = ((QAbstractProxyModel*)self)->QAbstractProxyModel::canFetchMore(*(QModelIndex*)parent);
	return __r;
}
void QAbstractProxyModel_fetchMore_default(void* self, void* parent) {
	((QAbstractProxyModel*)self)->QAbstractProxyModel::fetchMore(*(QModelIndex*)parent);
}
void QAbstractProxyModel_sort_default(void* self, int column, int order) {
	((QAbstractProxyModel*)self)->QAbstractProxyModel::sort(column, (Qt::SortOrder)(order));
}
void* QAbstractProxyModel_span_default(void* self, void* index) {
	auto&& __r = ((QAbstractProxyModel*)self)->QAbstractProxyModel::span(*(QModelIndex*)index);
	return new QSize(__r);
}
int QAbstractProxyModel_hasChildren_default(void* self, void* parent) {
	auto&& __r = ((QAbstractProxyModel*)self)->QAbstractProxyModel::hasChildren(*(QModelIndex*)parent);
	return __r;
}
void* QAbstractProxyModel_sibling_default(void* self, int row, int column, void* idx) {
	auto&& __r = ((QAbstractProxyModel*)self)->QAbstractProxyModel::sibling(row, column, *(QModelIndex*)idx);
	return new QModelIndex(__r);
}
void* QAbstractProxyModel_mimeData_default(void* self, void** indexes, long long indexes_len) {
	QList<QModelIndex> indexes_list; indexes_list.reserve(indexes_len);
	for (long long __i = 0; __i < indexes_len; __i++) indexes_list.append(*(QModelIndex*)indexes[__i]);
	auto&& __r = ((QAbstractProxyModel*)self)->QAbstractProxyModel::mimeData(indexes_list);
	return (void*)(__r);
}
int QAbstractProxyModel_canDropMimeData_default(void* self, void* data, int action, int row, int column, void* parent) {
	auto&& __r = ((QAbstractProxyModel*)self)->QAbstractProxyModel::canDropMimeData((QMimeData*)data, (Qt::DropAction)(action), row, column, *(QModelIndex*)parent);
	return __r;
}
int QAbstractProxyModel_dropMimeData_default(void* self, void* data, int action, int row, int column, void* parent) {
	auto&& __r = ((QAbstractProxyModel*)self)->QAbstractProxyModel::dropMimeData((QMimeData*)data, (Qt::DropAction)(action), row, column, *(QModelIndex*)parent);
	return __r;
}
char** QAbstractProxyModel_mimeTypes_default(void* self, long long* __out_len) {
	auto&& __r = ((QAbstractProxyModel*)self)->QAbstractProxyModel::mimeTypes();
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
int QAbstractProxyModel_supportedDragActions_default(void* self) {
	auto&& __r = ((QAbstractProxyModel*)self)->QAbstractProxyModel::supportedDragActions();
	return (int)(__r);
}
int QAbstractProxyModel_supportedDropActions_default(void* self) {
	auto&& __r = ((QAbstractProxyModel*)self)->QAbstractProxyModel::supportedDropActions();
	return (int)(__r);
}
int QAbstractProxyModel_insertRows_default(void* self, int row, int count, void* parent) {
	auto&& __r = ((QAbstractProxyModel*)self)->QAbstractProxyModel::insertRows(row, count, *(QModelIndex*)parent);
	return __r;
}
int QAbstractProxyModel_insertColumns_default(void* self, int column, int count, void* parent) {
	auto&& __r = ((QAbstractProxyModel*)self)->QAbstractProxyModel::insertColumns(column, count, *(QModelIndex*)parent);
	return __r;
}
int QAbstractProxyModel_removeRows_default(void* self, int row, int count, void* parent) {
	auto&& __r = ((QAbstractProxyModel*)self)->QAbstractProxyModel::removeRows(row, count, *(QModelIndex*)parent);
	return __r;
}
int QAbstractProxyModel_removeColumns_default(void* self, int column, int count, void* parent) {
	auto&& __r = ((QAbstractProxyModel*)self)->QAbstractProxyModel::removeColumns(column, count, *(QModelIndex*)parent);
	return __r;
}
int QAbstractProxyModel_moveRows_default(void* self, void* sourceParent, int sourceRow, int count, void* destinationParent, int destinationChild) {
	auto&& __r = ((QAbstractProxyModel*)self)->QAbstractProxyModel::moveRows(*(QModelIndex*)sourceParent, sourceRow, count, *(QModelIndex*)destinationParent, destinationChild);
	return __r;
}
int QAbstractProxyModel_moveColumns_default(void* self, void* sourceParent, int sourceColumn, int count, void* destinationParent, int destinationChild) {
	auto&& __r = ((QAbstractProxyModel*)self)->QAbstractProxyModel::moveColumns(*(QModelIndex*)sourceParent, sourceColumn, count, *(QModelIndex*)destinationParent, destinationChild);
	return __r;
}
void** QAbstractProxyModel_match_default(void* self, void* start, int role, void* value, int hits, int flags, long long* __out_len) {
	auto&& __r = ((QAbstractProxyModel*)self)->QAbstractProxyModel::match(*(QModelIndex*)start, role, *(QVariant*)value, hits, QFlags<Qt::MatchFlag>(flags));
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QModelIndex(__e);
	return __arr;
}
void QAbstractProxyModel_multiData_default(void* self, void* index, void* roleDataSpan) {
	((QAbstractProxyModel*)self)->QAbstractProxyModel::multiData(*(QModelIndex*)index, *(QModelRoleDataSpan*)roleDataSpan);
}
void QAbstractProxyModel_resetInternalData_default(void* self) {
	((QAbstractProxyModel__Go*)self)->cppgoBase_resetInternalData();
}
int QAbstractProxyModel_event_default(void* self, void* event) {
	auto&& __r = ((QAbstractProxyModel*)self)->QAbstractProxyModel::event((QEvent*)event);
	return __r;
}
int QAbstractProxyModel_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QAbstractProxyModel*)self)->QAbstractProxyModel::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QAbstractProxyModel_timerEvent_default(void* self, void* event) {
	((QAbstractProxyModel__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QAbstractProxyModel_childEvent_default(void* self, void* event) {
	((QAbstractProxyModel__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QAbstractProxyModel_customEvent_default(void* self, void* event) {
	((QAbstractProxyModel__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QAbstractProxyModel_connectNotify_default(void* self, void* signal) {
	((QAbstractProxyModel__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QAbstractProxyModel_disconnectNotify_default(void* self, void* signal) {
	((QAbstractProxyModel__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QAbstractProxyModel_connect_sourceModelChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractProxyModel*)self, &QAbstractProxyModel::sourceModelChanged, (QAbstractProxyModel*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractProxyModel);
	__conns_QAbstractProxyModel[__id] = __c;
}
void QAbstractProxyModel_connect_headerDataChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractProxyModel*)self, &QAbstractProxyModel::headerDataChanged, (QAbstractProxyModel*)self, [__id](Qt::Orientation orientation, int first, int last){
		if (cppgo_sig_void_int_int_int) ((void(*)(long long, int, int, int))cppgo_sig_void_int_int_int)(__id, (int)(orientation), first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractProxyModel);
	__conns_QAbstractProxyModel[__id] = __c;
}
void QAbstractProxyModel_connect_layoutChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractProxyModel*)self, &QAbstractProxyModel::layoutChanged, (QAbstractProxyModel*)self, [__id](const QList<QPersistentModelIndex>& parents, QAbstractItemModel::LayoutChangeHint hint){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractProxyModel);
	__conns_QAbstractProxyModel[__id] = __c;
}
void QAbstractProxyModel_connect_layoutChanged2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractProxyModel*)self, &QAbstractProxyModel::layoutChanged, (QAbstractProxyModel*)self, [__id](const QList<QPersistentModelIndex>& parents, QAbstractItemModel::LayoutChangeHint hint){
		std::vector<void*> parents_v; parents_v.reserve(parents.size());
		for (long long __i = 0; __i < parents.size(); __i++) parents_v.push_back((void*)&parents[__i]);
		if (cppgo_sig_void_voidpp_longlong_int) ((void(*)(long long, void**, long long, int))cppgo_sig_void_voidpp_longlong_int)(__id, parents_v.data(), (long long)parents.size(), (int)(hint));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractProxyModel);
	__conns_QAbstractProxyModel[__id] = __c;
}
void QAbstractProxyModel_connect_layoutAboutToBeChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractProxyModel*)self, &QAbstractProxyModel::layoutAboutToBeChanged, (QAbstractProxyModel*)self, [__id](const QList<QPersistentModelIndex>& parents, QAbstractItemModel::LayoutChangeHint hint){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractProxyModel);
	__conns_QAbstractProxyModel[__id] = __c;
}
void QAbstractProxyModel_connect_layoutAboutToBeChanged2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractProxyModel*)self, &QAbstractProxyModel::layoutAboutToBeChanged, (QAbstractProxyModel*)self, [__id](const QList<QPersistentModelIndex>& parents, QAbstractItemModel::LayoutChangeHint hint){
		std::vector<void*> parents_v; parents_v.reserve(parents.size());
		for (long long __i = 0; __i < parents.size(); __i++) parents_v.push_back((void*)&parents[__i]);
		if (cppgo_sig_void_voidpp_longlong_int) ((void(*)(long long, void**, long long, int))cppgo_sig_void_voidpp_longlong_int)(__id, parents_v.data(), (long long)parents.size(), (int)(hint));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractProxyModel);
	__conns_QAbstractProxyModel[__id] = __c;
}
void QAbstractProxyModel_connect_rowsAboutToBeInserted(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractProxyModel*)self, &QAbstractProxyModel::rowsAboutToBeInserted, (QAbstractProxyModel*)self, [__id](const QModelIndex& parent, int first, int last){
		if (cppgo_sig_void_voidp_int_int) ((void(*)(long long, void*, int, int))cppgo_sig_void_voidp_int_int)(__id, (void*)&parent, first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractProxyModel);
	__conns_QAbstractProxyModel[__id] = __c;
}
void QAbstractProxyModel_connect_rowsInserted(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractProxyModel*)self, &QAbstractProxyModel::rowsInserted, (QAbstractProxyModel*)self, [__id](const QModelIndex& parent, int first, int last){
		if (cppgo_sig_void_voidp_int_int) ((void(*)(long long, void*, int, int))cppgo_sig_void_voidp_int_int)(__id, (void*)&parent, first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractProxyModel);
	__conns_QAbstractProxyModel[__id] = __c;
}
void QAbstractProxyModel_connect_rowsAboutToBeRemoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractProxyModel*)self, &QAbstractProxyModel::rowsAboutToBeRemoved, (QAbstractProxyModel*)self, [__id](const QModelIndex& parent, int first, int last){
		if (cppgo_sig_void_voidp_int_int) ((void(*)(long long, void*, int, int))cppgo_sig_void_voidp_int_int)(__id, (void*)&parent, first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractProxyModel);
	__conns_QAbstractProxyModel[__id] = __c;
}
void QAbstractProxyModel_connect_rowsRemoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractProxyModel*)self, &QAbstractProxyModel::rowsRemoved, (QAbstractProxyModel*)self, [__id](const QModelIndex& parent, int first, int last){
		if (cppgo_sig_void_voidp_int_int) ((void(*)(long long, void*, int, int))cppgo_sig_void_voidp_int_int)(__id, (void*)&parent, first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractProxyModel);
	__conns_QAbstractProxyModel[__id] = __c;
}
void QAbstractProxyModel_connect_columnsAboutToBeInserted(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractProxyModel*)self, &QAbstractProxyModel::columnsAboutToBeInserted, (QAbstractProxyModel*)self, [__id](const QModelIndex& parent, int first, int last){
		if (cppgo_sig_void_voidp_int_int) ((void(*)(long long, void*, int, int))cppgo_sig_void_voidp_int_int)(__id, (void*)&parent, first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractProxyModel);
	__conns_QAbstractProxyModel[__id] = __c;
}
void QAbstractProxyModel_connect_columnsInserted(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractProxyModel*)self, &QAbstractProxyModel::columnsInserted, (QAbstractProxyModel*)self, [__id](const QModelIndex& parent, int first, int last){
		if (cppgo_sig_void_voidp_int_int) ((void(*)(long long, void*, int, int))cppgo_sig_void_voidp_int_int)(__id, (void*)&parent, first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractProxyModel);
	__conns_QAbstractProxyModel[__id] = __c;
}
void QAbstractProxyModel_connect_columnsAboutToBeRemoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractProxyModel*)self, &QAbstractProxyModel::columnsAboutToBeRemoved, (QAbstractProxyModel*)self, [__id](const QModelIndex& parent, int first, int last){
		if (cppgo_sig_void_voidp_int_int) ((void(*)(long long, void*, int, int))cppgo_sig_void_voidp_int_int)(__id, (void*)&parent, first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractProxyModel);
	__conns_QAbstractProxyModel[__id] = __c;
}
void QAbstractProxyModel_connect_columnsRemoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractProxyModel*)self, &QAbstractProxyModel::columnsRemoved, (QAbstractProxyModel*)self, [__id](const QModelIndex& parent, int first, int last){
		if (cppgo_sig_void_voidp_int_int) ((void(*)(long long, void*, int, int))cppgo_sig_void_voidp_int_int)(__id, (void*)&parent, first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractProxyModel);
	__conns_QAbstractProxyModel[__id] = __c;
}
void QAbstractProxyModel_connect_modelAboutToBeReset(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractProxyModel*)self, &QAbstractProxyModel::modelAboutToBeReset, (QAbstractProxyModel*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractProxyModel);
	__conns_QAbstractProxyModel[__id] = __c;
}
void QAbstractProxyModel_connect_modelReset(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractProxyModel*)self, &QAbstractProxyModel::modelReset, (QAbstractProxyModel*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractProxyModel);
	__conns_QAbstractProxyModel[__id] = __c;
}
void QAbstractProxyModel_connect_rowsAboutToBeMoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractProxyModel*)self, &QAbstractProxyModel::rowsAboutToBeMoved, (QAbstractProxyModel*)self, [__id](const QModelIndex& sourceParent, int sourceStart, int sourceEnd, const QModelIndex& destinationParent, int destinationRow){
		if (cppgo_sig_void_voidp_int_int_voidp_int) ((void(*)(long long, void*, int, int, void*, int))cppgo_sig_void_voidp_int_int_voidp_int)(__id, (void*)&sourceParent, sourceStart, sourceEnd, (void*)&destinationParent, destinationRow);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractProxyModel);
	__conns_QAbstractProxyModel[__id] = __c;
}
void QAbstractProxyModel_connect_rowsMoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractProxyModel*)self, &QAbstractProxyModel::rowsMoved, (QAbstractProxyModel*)self, [__id](const QModelIndex& sourceParent, int sourceStart, int sourceEnd, const QModelIndex& destinationParent, int destinationRow){
		if (cppgo_sig_void_voidp_int_int_voidp_int) ((void(*)(long long, void*, int, int, void*, int))cppgo_sig_void_voidp_int_int_voidp_int)(__id, (void*)&sourceParent, sourceStart, sourceEnd, (void*)&destinationParent, destinationRow);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractProxyModel);
	__conns_QAbstractProxyModel[__id] = __c;
}
void QAbstractProxyModel_connect_columnsAboutToBeMoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractProxyModel*)self, &QAbstractProxyModel::columnsAboutToBeMoved, (QAbstractProxyModel*)self, [__id](const QModelIndex& sourceParent, int sourceStart, int sourceEnd, const QModelIndex& destinationParent, int destinationColumn){
		if (cppgo_sig_void_voidp_int_int_voidp_int) ((void(*)(long long, void*, int, int, void*, int))cppgo_sig_void_voidp_int_int_voidp_int)(__id, (void*)&sourceParent, sourceStart, sourceEnd, (void*)&destinationParent, destinationColumn);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractProxyModel);
	__conns_QAbstractProxyModel[__id] = __c;
}
void QAbstractProxyModel_connect_columnsMoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractProxyModel*)self, &QAbstractProxyModel::columnsMoved, (QAbstractProxyModel*)self, [__id](const QModelIndex& sourceParent, int sourceStart, int sourceEnd, const QModelIndex& destinationParent, int destinationColumn){
		if (cppgo_sig_void_voidp_int_int_voidp_int) ((void(*)(long long, void*, int, int, void*, int))cppgo_sig_void_voidp_int_int_voidp_int)(__id, (void*)&sourceParent, sourceStart, sourceEnd, (void*)&destinationParent, destinationColumn);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractProxyModel);
	__conns_QAbstractProxyModel[__id] = __c;
}
void QAbstractProxyModel_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractProxyModel*)self, &QAbstractProxyModel::destroyed, (QAbstractProxyModel*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractProxyModel);
	__conns_QAbstractProxyModel[__id] = __c;
}
void QAbstractProxyModel_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractProxyModel*)self, &QAbstractProxyModel::destroyed, (QAbstractProxyModel*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractProxyModel);
	__conns_QAbstractProxyModel[__id] = __c;
}
void QAbstractProxyModel_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractProxyModel*)self, &QAbstractProxyModel::objectNameChanged, (QAbstractProxyModel*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractProxyModel);
	__conns_QAbstractProxyModel[__id] = __c;
}
void QAbstractProxyModel__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QAbstractProxyModel);
		auto __it = __conns_QAbstractProxyModel.find(id);
		if (__it == __conns_QAbstractProxyModel.end()) return;
		__c = __it->second;
		__conns_QAbstractProxyModel.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QAbstractProxyModel__Go_new(long long id, void* parent) { return new QAbstractProxyModel__Go(id, (QObject*)parent); }
void QAbstractProxyModel_delete(void* self) { delete (QAbstractProxyModel*)self; }

}

#include <QtCore/qdatetime.h>
#include <QtCore/qcalendar.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qtimezone.h>
#include <QtCore/qstring.h>
#include <QtCore/qstringview.h>
#include "gen_qdate.h"

extern "C" {

void* QDate_new() {
	return new QDate();
}
void* QDate_new2(int y, int m, int d) {
	return new QDate(y, m, d);
}
void* QDate_new3(int y, int m, int d, void* cal) {
	return new QDate(y, m, d, *(QCalendar*)cal);
}
void* QDate_new4(void* arg1) {
	return new QDate(*(QDate*)arg1);
}
int QDate_isNull(void* self) {
	auto&& __r = ((QDate*)self)->isNull();
	return __r;
}
int QDate_isValid(void* self) {
	auto&& __r = ((QDate*)self)->isValid();
	return __r;
}
int QDate_year(void* self) {
	auto&& __r = ((QDate*)self)->year();
	return __r;
}
int QDate_month(void* self) {
	auto&& __r = ((QDate*)self)->month();
	return __r;
}
int QDate_day(void* self) {
	auto&& __r = ((QDate*)self)->day();
	return __r;
}
int QDate_dayOfWeek(void* self) {
	auto&& __r = ((QDate*)self)->dayOfWeek();
	return __r;
}
int QDate_dayOfYear(void* self) {
	auto&& __r = ((QDate*)self)->dayOfYear();
	return __r;
}
int QDate_daysInMonth(void* self) {
	auto&& __r = ((QDate*)self)->daysInMonth();
	return __r;
}
int QDate_daysInYear(void* self) {
	auto&& __r = ((QDate*)self)->daysInYear();
	return __r;
}
int QDate_weekNumber(void* self) {
	auto&& __r = ((QDate*)self)->weekNumber();
	return __r;
}
int QDate_year2(void* self, void* cal) {
	auto&& __r = ((QDate*)self)->year(*(QCalendar*)cal);
	return __r;
}
int QDate_month2(void* self, void* cal) {
	auto&& __r = ((QDate*)self)->month(*(QCalendar*)cal);
	return __r;
}
int QDate_day2(void* self, void* cal) {
	auto&& __r = ((QDate*)self)->day(*(QCalendar*)cal);
	return __r;
}
int QDate_dayOfWeek2(void* self, void* cal) {
	auto&& __r = ((QDate*)self)->dayOfWeek(*(QCalendar*)cal);
	return __r;
}
int QDate_dayOfYear2(void* self, void* cal) {
	auto&& __r = ((QDate*)self)->dayOfYear(*(QCalendar*)cal);
	return __r;
}
int QDate_daysInMonth2(void* self, void* cal) {
	auto&& __r = ((QDate*)self)->daysInMonth(*(QCalendar*)cal);
	return __r;
}
int QDate_daysInYear2(void* self, void* cal) {
	auto&& __r = ((QDate*)self)->daysInYear(*(QCalendar*)cal);
	return __r;
}
void* QDate_startOfDay(void* self, int spec) {
	auto&& __r = ((QDate*)self)->startOfDay((Qt::TimeSpec)(spec));
	return new QDateTime(__r);
}
void* QDate_endOfDay(void* self, int spec) {
	auto&& __r = ((QDate*)self)->endOfDay((Qt::TimeSpec)(spec));
	return new QDateTime(__r);
}
void* QDate_startOfDay2(void* self, void* zone) {
	auto&& __r = ((QDate*)self)->startOfDay(*(QTimeZone*)zone);
	return new QDateTime(__r);
}
void* QDate_endOfDay2(void* self, void* zone) {
	auto&& __r = ((QDate*)self)->endOfDay(*(QTimeZone*)zone);
	return new QDateTime(__r);
}
void* QDate_startOfDay3(void* self) {
	auto&& __r = ((QDate*)self)->startOfDay();
	return new QDateTime(__r);
}
void* QDate_endOfDay3(void* self) {
	auto&& __r = ((QDate*)self)->endOfDay();
	return new QDateTime(__r);
}
char* QDate_toString(void* self, long long* __out_len) {
	auto&& __r = ((QDate*)self)->toString();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QDate_toString2(void* self, const char* format, long long format_len, long long* __out_len) {
	QString format_qstr = QString::fromUtf8(format, format_len);
	auto&& __r = ((QDate*)self)->toString(format_qstr);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QDate_toString3(void* self, const char* format, long long format_len, void* cal, long long* __out_len) {
	QString format_qstr = QString::fromUtf8(format, format_len);
	auto&& __r = ((QDate*)self)->toString(format_qstr, *(QCalendar*)cal);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QDate_toString4(void* self, void* format, long long* __out_len) {
	auto&& __r = ((QDate*)self)->toString(*(QStringView*)format);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QDate_toString5(void* self, void* format, void* cal, long long* __out_len) {
	auto&& __r = ((QDate*)self)->toString(*(QStringView*)format, *(QCalendar*)cal);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QDate_setDate(void* self, int year, int month, int day) {
	auto&& __r = ((QDate*)self)->setDate(year, month, day);
	return __r;
}
int QDate_setDate2(void* self, int year, int month, int day, void* cal) {
	auto&& __r = ((QDate*)self)->setDate(year, month, day, *(QCalendar*)cal);
	return __r;
}
void QDate_getDate(void* self, int* year, int* month, int* day) {
	int year_local{};
	int month_local{};
	int day_local{};
	((QDate*)self)->getDate(&year_local, &month_local, &day_local);
	*year = year_local;
	*month = month_local;
	*day = day_local;
}
void* QDate_addDays(void* self, long long days) {
	auto&& __r = ((QDate*)self)->addDays(days);
	return new QDate(__r);
}
void* QDate_addMonths(void* self, int months) {
	auto&& __r = ((QDate*)self)->addMonths(months);
	return new QDate(__r);
}
void* QDate_addYears(void* self, int years) {
	auto&& __r = ((QDate*)self)->addYears(years);
	return new QDate(__r);
}
void* QDate_addMonths2(void* self, int months, void* cal) {
	auto&& __r = ((QDate*)self)->addMonths(months, *(QCalendar*)cal);
	return new QDate(__r);
}
void* QDate_addYears2(void* self, int years, void* cal) {
	auto&& __r = ((QDate*)self)->addYears(years, *(QCalendar*)cal);
	return new QDate(__r);
}
long long QDate_daysTo(void* self, void* d) {
	auto&& __r = ((QDate*)self)->daysTo(*(QDate*)d);
	return __r;
}
void* QDate_currentDate() {
	auto&& __r = QDate::currentDate();
	return new QDate(__r);
}
void* QDate_fromString(void* string_) {
	auto&& __r = QDate::fromString(*(QStringView*)string_);
	return new QDate(__r);
}
void* QDate_fromString2(const char* string_, long long string__len) {
	QString string__qstr = QString::fromUtf8(string_, string__len);
	auto&& __r = QDate::fromString(string__qstr);
	return new QDate(__r);
}
void* QDate_fromString3(void* string_, void* format, void* cal) {
	auto&& __r = QDate::fromString(*(QStringView*)string_, *(QStringView*)format, *(QCalendar*)cal);
	return new QDate(__r);
}
void* QDate_fromString4(const char* string_, long long string__len, void* format, void* cal) {
	QString string__qstr = QString::fromUtf8(string_, string__len);
	auto&& __r = QDate::fromString(string__qstr, *(QStringView*)format, *(QCalendar*)cal);
	return new QDate(__r);
}
void* QDate_fromString5(const char* string_, long long string__len, const char* format, long long format_len, void* cal) {
	QString string__qstr = QString::fromUtf8(string_, string__len);
	QString format_qstr = QString::fromUtf8(format, format_len);
	auto&& __r = QDate::fromString(string__qstr, format_qstr, *(QCalendar*)cal);
	return new QDate(__r);
}
void* QDate_fromString6(void* string_, void* format) {
	auto&& __r = QDate::fromString(*(QStringView*)string_, *(QStringView*)format);
	return new QDate(__r);
}
void* QDate_fromString7(void* string_, void* format, int baseYear, void* cal) {
	auto&& __r = QDate::fromString(*(QStringView*)string_, *(QStringView*)format, baseYear, *(QCalendar*)cal);
	return new QDate(__r);
}
void* QDate_fromString8(const char* string_, long long string__len, void* format) {
	QString string__qstr = QString::fromUtf8(string_, string__len);
	auto&& __r = QDate::fromString(string__qstr, *(QStringView*)format);
	return new QDate(__r);
}
void* QDate_fromString9(const char* string_, long long string__len, void* format, int baseYear, void* cal) {
	QString string__qstr = QString::fromUtf8(string_, string__len);
	auto&& __r = QDate::fromString(string__qstr, *(QStringView*)format, baseYear, *(QCalendar*)cal);
	return new QDate(__r);
}
void* QDate_fromString10(const char* string_, long long string__len, const char* format, long long format_len) {
	QString string__qstr = QString::fromUtf8(string_, string__len);
	QString format_qstr = QString::fromUtf8(format, format_len);
	auto&& __r = QDate::fromString(string__qstr, format_qstr);
	return new QDate(__r);
}
void* QDate_fromString11(const char* string_, long long string__len, const char* format, long long format_len, int baseYear, void* cal) {
	QString string__qstr = QString::fromUtf8(string_, string__len);
	QString format_qstr = QString::fromUtf8(format, format_len);
	auto&& __r = QDate::fromString(string__qstr, format_qstr, baseYear, *(QCalendar*)cal);
	return new QDate(__r);
}
int QDate_isValid2(int y, int m, int d) {
	auto&& __r = QDate::isValid(y, m, d);
	return __r;
}
int QDate_isLeapYear(int year) {
	auto&& __r = QDate::isLeapYear(year);
	return __r;
}
void* QDate_fromJulianDay(long long jd_) {
	auto&& __r = QDate::fromJulianDay(jd_);
	return new QDate(__r);
}
long long QDate_toJulianDay(void* self) {
	auto&& __r = ((QDate*)self)->toJulianDay();
	return __r;
}
int QDate_weekNumber2(void* self, int* yearNum) {
	int yearNum_local{};
	auto&& __r = ((QDate*)self)->weekNumber(&yearNum_local);
	*yearNum = yearNum_local;
	return __r;
}
void* QDate_startOfDay4(void* self, int spec, int offsetSeconds) {
	auto&& __r = ((QDate*)self)->startOfDay((Qt::TimeSpec)(spec), offsetSeconds);
	return new QDateTime(__r);
}
void* QDate_endOfDay4(void* self, int spec, int offsetSeconds) {
	auto&& __r = ((QDate*)self)->endOfDay((Qt::TimeSpec)(spec), offsetSeconds);
	return new QDateTime(__r);
}
char* QDate_toString6(void* self, int format, long long* __out_len) {
	auto&& __r = ((QDate*)self)->toString((Qt::DateFormat)(format));
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QDate_fromString12(void* string_, int format) {
	auto&& __r = QDate::fromString(*(QStringView*)string_, (Qt::DateFormat)(format));
	return new QDate(__r);
}
void* QDate_fromString13(const char* string_, long long string__len, int format) {
	QString string__qstr = QString::fromUtf8(string_, string__len);
	auto&& __r = QDate::fromString(string__qstr, (Qt::DateFormat)(format));
	return new QDate(__r);
}
void* QDate_fromString14(void* string_, void* format, int baseYear) {
	auto&& __r = QDate::fromString(*(QStringView*)string_, *(QStringView*)format, baseYear);
	return new QDate(__r);
}
void* QDate_fromString15(const char* string_, long long string__len, void* format, int baseYear) {
	QString string__qstr = QString::fromUtf8(string_, string__len);
	auto&& __r = QDate::fromString(string__qstr, *(QStringView*)format, baseYear);
	return new QDate(__r);
}
void* QDate_fromString16(const char* string_, long long string__len, const char* format, long long format_len, int baseYear) {
	QString string__qstr = QString::fromUtf8(string_, string__len);
	QString format_qstr = QString::fromUtf8(format, format_len);
	auto&& __r = QDate::fromString(string__qstr, format_qstr, baseYear);
	return new QDate(__r);
}
int QDate_OperatorEqual(void* self, void* rhs) {
	auto&& __r = (*(QDate*)self) == (*(QDate*)rhs);
	return __r;
}
int QDate_OperatorNotEqual(void* self, void* rhs) {
	auto&& __r = (*(QDate*)self) != (*(QDate*)rhs);
	return __r;
}
int QDate_OperatorLesser(void* self, void* rhs) {
	auto&& __r = (*(QDate*)self) < (*(QDate*)rhs);
	return __r;
}
int QDate_OperatorGreater(void* self, void* rhs) {
	auto&& __r = (*(QDate*)self) > (*(QDate*)rhs);
	return __r;
}
int QDate_OperatorLesserOrEqual(void* self, void* rhs) {
	auto&& __r = (*(QDate*)self) <= (*(QDate*)rhs);
	return __r;
}
int QDate_OperatorGreaterOrEqual(void* self, void* rhs) {
	auto&& __r = (*(QDate*)self) >= (*(QDate*)rhs);
	return __r;
}
void* QDate_OperatorPlusPlus(void* self) {
	auto&& __r = ++(*(QDate*)self);
	return (void*)&(__r);
}
void* QDate_OperatorMinusMinus(void* self) {
	auto&& __r = --(*(QDate*)self);
	return (void*)&(__r);
}
void* QDate_OperatorPlusPlus2(void* self, int arg1) {
	auto&& __r = (*(QDate*)self)++;
	return new QDate(__r);
}
void* QDate_OperatorMinusMinus2(void* self, int arg1) {
	auto&& __r = (*(QDate*)self)--;
	return new QDate(__r);
}
void QDate_delete(void* self) { delete (QDate*)self; }

}

#include <QtCore/qproperty.h>
#include "gen_qpropertyobserverbase.h"

extern "C" {

void* QPropertyObserverBase_new() {
	return new QPropertyObserverBase();
}
void* QPropertyObserverBase_new2(void* arg1) {
	return new QPropertyObserverBase(*(QPropertyObserverBase*)arg1);
}
void* QPropertyObserverBase_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QPropertyObserverBase*)self)->operator=(*(QPropertyObserverBase*)arg1);
	return (void*)&(__r);
}
void QPropertyObserverBase_delete(void* self) { delete (QPropertyObserverBase*)self; }

}

#include <QtCore/qelapsedtimer.h>
#include "gen_qelapsedtimer.h"

extern "C" {

void* QElapsedTimer_new() {
	return new QElapsedTimer();
}
int QElapsedTimer_clockType() {
	auto&& __r = QElapsedTimer::clockType();
	return (int)(__r);
}
int QElapsedTimer_isMonotonic() {
	auto&& __r = QElapsedTimer::isMonotonic();
	return __r;
}
void QElapsedTimer_start(void* self) {
	((QElapsedTimer*)self)->start();
}
long long QElapsedTimer_restart(void* self) {
	auto&& __r = ((QElapsedTimer*)self)->restart();
	return __r;
}
void QElapsedTimer_invalidate(void* self) {
	((QElapsedTimer*)self)->invalidate();
}
int QElapsedTimer_isValid(void* self) {
	auto&& __r = ((QElapsedTimer*)self)->isValid();
	return __r;
}
long long QElapsedTimer_nsecsElapsed(void* self) {
	auto&& __r = ((QElapsedTimer*)self)->nsecsElapsed();
	return __r;
}
long long QElapsedTimer_elapsed(void* self) {
	auto&& __r = ((QElapsedTimer*)self)->elapsed();
	return __r;
}
int QElapsedTimer_hasExpired(void* self, long long timeout) {
	auto&& __r = ((QElapsedTimer*)self)->hasExpired(timeout);
	return __r;
}
long long QElapsedTimer_msecsSinceReference(void* self) {
	auto&& __r = ((QElapsedTimer*)self)->msecsSinceReference();
	return __r;
}
long long QElapsedTimer_msecsTo(void* self, void* other) {
	auto&& __r = ((QElapsedTimer*)self)->msecsTo(*(QElapsedTimer*)other);
	return __r;
}
long long QElapsedTimer_secsTo(void* self, void* other) {
	auto&& __r = ((QElapsedTimer*)self)->secsTo(*(QElapsedTimer*)other);
	return __r;
}
int QElapsedTimer_OperatorLesser(void* self, void* rhs) {
	auto&& __r = (*(QElapsedTimer*)self) < (*(QElapsedTimer*)rhs);
	return __r;
}
int QElapsedTimer_OperatorEqual(void* self, void* rhs) {
	auto&& __r = (*(QElapsedTimer*)self) == (*(QElapsedTimer*)rhs);
	return __r;
}
int QElapsedTimer_OperatorNotEqual(void* self, void* rhs) {
	auto&& __r = (*(QElapsedTimer*)self) != (*(QElapsedTimer*)rhs);
	return __r;
}
int QElapsedTimer_OperatorGreater(void* self, void* rhs) {
	auto&& __r = (*(QElapsedTimer*)self) > (*(QElapsedTimer*)rhs);
	return __r;
}
int QElapsedTimer_OperatorLesserOrEqual(void* self, void* rhs) {
	auto&& __r = (*(QElapsedTimer*)self) <= (*(QElapsedTimer*)rhs);
	return __r;
}
int QElapsedTimer_OperatorGreaterOrEqual(void* self, void* rhs) {
	auto&& __r = (*(QElapsedTimer*)self) >= (*(QElapsedTimer*)rhs);
	return __r;
}
void QElapsedTimer_delete(void* self) { delete (QElapsedTimer*)self; }

}

#include <QtCore/qtyperevision.h>
#include "gen_qtyperevision.h"

extern "C" {

void* QTypeRevision_new() {
	return new QTypeRevision();
}
void* QTypeRevision_new2(void* arg1) {
	return new QTypeRevision(*(QTypeRevision*)arg1);
}
void* QTypeRevision_zero() {
	auto&& __r = QTypeRevision::zero();
	return new QTypeRevision(__r);
}
int QTypeRevision_hasMajorVersion(void* self) {
	auto&& __r = ((QTypeRevision*)self)->hasMajorVersion();
	return __r;
}
unsigned char QTypeRevision_majorVersion(void* self) {
	auto&& __r = ((QTypeRevision*)self)->majorVersion();
	return __r;
}
int QTypeRevision_hasMinorVersion(void* self) {
	auto&& __r = ((QTypeRevision*)self)->hasMinorVersion();
	return __r;
}
unsigned char QTypeRevision_minorVersion(void* self) {
	auto&& __r = ((QTypeRevision*)self)->minorVersion();
	return __r;
}
int QTypeRevision_isValid(void* self) {
	auto&& __r = ((QTypeRevision*)self)->isValid();
	return __r;
}
int QTypeRevision_OperatorEqual(void* self, void* rhs) {
	auto&& __r = (*(QTypeRevision*)self) == (*(QTypeRevision*)rhs);
	return __r;
}
int QTypeRevision_OperatorNotEqual(void* self, void* rhs) {
	auto&& __r = (*(QTypeRevision*)self) != (*(QTypeRevision*)rhs);
	return __r;
}
int QTypeRevision_OperatorLesser(void* self, void* rhs) {
	auto&& __r = (*(QTypeRevision*)self) < (*(QTypeRevision*)rhs);
	return __r;
}
int QTypeRevision_OperatorGreater(void* self, void* rhs) {
	auto&& __r = (*(QTypeRevision*)self) > (*(QTypeRevision*)rhs);
	return __r;
}
int QTypeRevision_OperatorLesserOrEqual(void* self, void* rhs) {
	auto&& __r = (*(QTypeRevision*)self) <= (*(QTypeRevision*)rhs);
	return __r;
}
int QTypeRevision_OperatorGreaterOrEqual(void* self, void* rhs) {
	auto&& __r = (*(QTypeRevision*)self) >= (*(QTypeRevision*)rhs);
	return __r;
}
void QTypeRevision_delete(void* self) { delete (QTypeRevision*)self; }

}

#include <QtCore/qpauseanimation.h>
#include <QtCore/qabstractanimation.h>
#include <QtCore/qobject.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qanimationgroup.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qvariant.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qpauseanimation.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QPauseAnimation__Go : public QPauseAnimation {
public:
	long long __goId;
	unsigned char __goOn[12] = {};
	template<typename... A> QPauseAnimation__Go(long long id, A&&... a) : QPauseAnimation(std::forward<A>(a)...), __goId(id) {}
	using QPauseAnimation::sender;
	using QPauseAnimation::senderSignalIndex;
	using QPauseAnimation::receivers;
	using QPauseAnimation::isSignalConnected;
	bool cppgoBase_event(QEvent* __a0) { return QPauseAnimation::event(__a0); }
	void cppgoBase_updateCurrentTime(int __a0) { QPauseAnimation::updateCurrentTime(__a0); }
	void cppgoBase_updateState(int __a0, int __a1) { QPauseAnimation::updateState((QAbstractAnimation::State)__a0, (QAbstractAnimation::State)__a1); }
	void cppgoBase_updateDirection(int __a0) { QPauseAnimation::updateDirection((QAbstractAnimation::Direction)__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QPauseAnimation::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QPauseAnimation::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QPauseAnimation::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QPauseAnimation::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QPauseAnimation::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QPauseAnimation::metaObject();
	}
	int duration() const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8434675112194400115LL, &__h); if (__h) return __r; }
		return QPauseAnimation::duration();
	}
	bool event(QEvent* e) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)e, &__h); if (__h) return __r; }
		return QPauseAnimation::event(e);
	}
	void updateCurrentTime(int arg1) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, -1342418961870150968LL, arg1, &__h); if (__h) return; }
		QPauseAnimation::updateCurrentTime(arg1);
	}
	void updateState(QAbstractAnimation::State newState, QAbstractAnimation::State oldState) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_int_int) { ((void(*)(long long, long long, int, int, int*))cppgo_virt_void_int_int)(__goId, 7692482089134926569LL, (int)(newState), (int)(oldState), &__h); if (__h) return; }
		QPauseAnimation::updateState(newState, oldState);
	}
	void updateDirection(QAbstractAnimation::Direction direction) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, -3231888725407197643LL, (int)(direction), &__h); if (__h) return; }
		QPauseAnimation::updateDirection(direction);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QPauseAnimation::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QPauseAnimation::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QPauseAnimation::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QPauseAnimation::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QPauseAnimation::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QPauseAnimation::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QPauseAnimation;
static std::mutex __connsMu_QPauseAnimation;

extern "C" {

void* QPauseAnimation_base(void* self) { return (void*)static_cast<QAbstractAnimation*>((QPauseAnimation*)self); }
void QPauseAnimation_chain(void* self, void** out) {
	QPauseAnimation* __p = (QPauseAnimation*)self;
	QAbstractAnimation* __b0 = static_cast<QAbstractAnimation*>(__p); out[0] = (void*)__b0;
	QObject* __b1 = static_cast<QObject*>(__b0); out[1] = (void*)__b1;
}
void QPauseAnimation__set_override(void* self, int idx) { ((QPauseAnimation__Go*)self)->__goOn[idx] = 1; }
void* QPauseAnimation_fromQObject(void* o) { return (void*)qobject_cast<QPauseAnimation*>((QObject*)o); }
void* QPauseAnimation_new(long long __goid) {
	return new QPauseAnimation__Go(__goid);
}
void* QPauseAnimation_new2(long long __goid, int msecs) {
	return new QPauseAnimation__Go(__goid, msecs);
}
void* QPauseAnimation_new3(long long __goid, void* parent) {
	return new QPauseAnimation__Go(__goid, (QObject*)parent);
}
void* QPauseAnimation_new4(long long __goid, int msecs, void* parent) {
	return new QPauseAnimation__Go(__goid, msecs, (QObject*)parent);
}
char* QPauseAnimation_tr(const char* s, long long* __out_len) {
	auto&& __r = QPauseAnimation::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QPauseAnimation_setDuration(void* self, int msecs) {
	((QPauseAnimation*)self)->setDuration(msecs);
}
void* QPauseAnimation_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QPauseAnimation::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QPauseAnimation_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QPauseAnimation::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QPauseAnimation_sender(void* self) {
	auto&& __r = ((QPauseAnimation__Go*)self)->sender();
	return (void*)(__r);
}
int QPauseAnimation_senderSignalIndex(void* self) {
	auto&& __r = ((QPauseAnimation__Go*)self)->senderSignalIndex();
	return __r;
}
int QPauseAnimation_receivers(void* self, const char* signal) {
	auto&& __r = ((QPauseAnimation__Go*)self)->receivers(signal);
	return __r;
}
int QPauseAnimation_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QPauseAnimation__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QPauseAnimation_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QPauseAnimation::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QPauseAnimation_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QPauseAnimation::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QPauseAnimation_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QPauseAnimation::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QPauseAnimation_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QPauseAnimation::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QPauseAnimation_metaObject_default(void* self) {
	auto&& __r = ((QPauseAnimation*)self)->QPauseAnimation::metaObject();
	return (void*)(__r);
}
int QPauseAnimation_duration_default(void* self) {
	auto&& __r = ((QPauseAnimation*)self)->QPauseAnimation::duration();
	return __r;
}
int QPauseAnimation_event_default(void* self, void* e) {
	auto&& __r = ((QPauseAnimation__Go*)self)->cppgoBase_event((QEvent*)e);
	return __r;
}
void QPauseAnimation_updateCurrentTime_default(void* self, int arg1) {
	((QPauseAnimation__Go*)self)->cppgoBase_updateCurrentTime(arg1);
}
void QPauseAnimation_updateState_default(void* self, int newState, int oldState) {
	((QPauseAnimation__Go*)self)->cppgoBase_updateState(newState, oldState);
}
void QPauseAnimation_updateDirection_default(void* self, int direction) {
	((QPauseAnimation__Go*)self)->cppgoBase_updateDirection(direction);
}
int QPauseAnimation_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QPauseAnimation*)self)->QPauseAnimation::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QPauseAnimation_timerEvent_default(void* self, void* event) {
	((QPauseAnimation__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QPauseAnimation_childEvent_default(void* self, void* event) {
	((QPauseAnimation__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QPauseAnimation_customEvent_default(void* self, void* event) {
	((QPauseAnimation__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QPauseAnimation_connectNotify_default(void* self, void* signal) {
	((QPauseAnimation__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QPauseAnimation_disconnectNotify_default(void* self, void* signal) {
	((QPauseAnimation__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QPauseAnimation_connect_finished(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QPauseAnimation*)self, &QPauseAnimation::finished, (QPauseAnimation*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QPauseAnimation);
	__conns_QPauseAnimation[__id] = __c;
}
void QPauseAnimation_connect_stateChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QPauseAnimation*)self, &QPauseAnimation::stateChanged, (QPauseAnimation*)self, [__id](QAbstractAnimation::State newState, QAbstractAnimation::State oldState){
		if (cppgo_sig_void_int_int) ((void(*)(long long, int, int))cppgo_sig_void_int_int)(__id, (int)(newState), (int)(oldState));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QPauseAnimation);
	__conns_QPauseAnimation[__id] = __c;
}
void QPauseAnimation_connect_currentLoopChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QPauseAnimation*)self, &QPauseAnimation::currentLoopChanged, (QPauseAnimation*)self, [__id](int currentLoop){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, currentLoop);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QPauseAnimation);
	__conns_QPauseAnimation[__id] = __c;
}
void QPauseAnimation_connect_directionChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QPauseAnimation*)self, &QPauseAnimation::directionChanged, (QPauseAnimation*)self, [__id](QAbstractAnimation::Direction arg1){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, (int)(arg1));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QPauseAnimation);
	__conns_QPauseAnimation[__id] = __c;
}
void QPauseAnimation_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QPauseAnimation*)self, &QPauseAnimation::destroyed, (QPauseAnimation*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QPauseAnimation);
	__conns_QPauseAnimation[__id] = __c;
}
void QPauseAnimation_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QPauseAnimation*)self, &QPauseAnimation::destroyed, (QPauseAnimation*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QPauseAnimation);
	__conns_QPauseAnimation[__id] = __c;
}
void QPauseAnimation_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QPauseAnimation*)self, &QPauseAnimation::objectNameChanged, (QPauseAnimation*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QPauseAnimation);
	__conns_QPauseAnimation[__id] = __c;
}
void QPauseAnimation__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QPauseAnimation);
		auto __it = __conns_QPauseAnimation.find(id);
		if (__it == __conns_QPauseAnimation.end()) return;
		__c = __it->second;
		__conns_QPauseAnimation.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QPauseAnimation__Go_new(long long id, void* parent) { return new QPauseAnimation__Go(id, (QObject*)parent); }
void QPauseAnimation_delete(void* self) { delete (QPauseAnimation*)self; }

}

#include <QtCore/qreadwritelock.h>
#include "gen_qreadlocker.h"

extern "C" {

void* QReadLocker_new(void* readWriteLock) {
	return new QReadLocker((QReadWriteLock*)readWriteLock);
}
void QReadLocker_unlock(void* self) {
	((QReadLocker*)self)->unlock();
}
void QReadLocker_relock(void* self) {
	((QReadLocker*)self)->relock();
}
void* QReadLocker_readWriteLock(void* self) {
	auto&& __r = ((QReadLocker*)self)->readWriteLock();
	return (void*)(__r);
}
void QReadLocker_delete(void* self) { delete (QReadLocker*)self; }

}

#include <QtCore/qsystemsemaphore.h>
#include <QtCore/qtipccommon.h>
#include <QtCore/qstring.h>
#include "gen_qsystemsemaphore.h"

extern "C" {

void* QSystemSemaphore_new(void* key) {
	return new QSystemSemaphore(*(QNativeIpcKey*)key);
}
void* QSystemSemaphore_new2(const char* key, long long key_len) {
	QString key_qstr = QString::fromUtf8(key, key_len);
	return new QSystemSemaphore(key_qstr);
}
void* QSystemSemaphore_new3(void* key, int initialValue) {
	return new QSystemSemaphore(*(QNativeIpcKey*)key, initialValue);
}
void* QSystemSemaphore_new4(void* key, int initialValue, int arg3) {
	return new QSystemSemaphore(*(QNativeIpcKey*)key, initialValue, (QSystemSemaphore::AccessMode)(arg3));
}
void* QSystemSemaphore_new5(const char* key, long long key_len, int initialValue) {
	QString key_qstr = QString::fromUtf8(key, key_len);
	return new QSystemSemaphore(key_qstr, initialValue);
}
void* QSystemSemaphore_new6(const char* key, long long key_len, int initialValue, int mode) {
	QString key_qstr = QString::fromUtf8(key, key_len);
	return new QSystemSemaphore(key_qstr, initialValue, (QSystemSemaphore::AccessMode)(mode));
}
char* QSystemSemaphore_tr(const char* sourceText, long long* __out_len) {
	auto&& __r = QSystemSemaphore::tr(sourceText);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QSystemSemaphore_setNativeKey(void* self, void* key) {
	((QSystemSemaphore*)self)->setNativeKey(*(QNativeIpcKey*)key);
}
void QSystemSemaphore_setNativeKey2(void* self, const char* key, long long key_len) {
	QString key_qstr = QString::fromUtf8(key, key_len);
	((QSystemSemaphore*)self)->setNativeKey(key_qstr);
}
void* QSystemSemaphore_nativeIpcKey(void* self) {
	auto&& __r = ((QSystemSemaphore*)self)->nativeIpcKey();
	return new QNativeIpcKey(__r);
}
void QSystemSemaphore_setKey(void* self, const char* key, long long key_len) {
	QString key_qstr = QString::fromUtf8(key, key_len);
	((QSystemSemaphore*)self)->setKey(key_qstr);
}
char* QSystemSemaphore_key(void* self, long long* __out_len) {
	auto&& __r = ((QSystemSemaphore*)self)->key();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QSystemSemaphore_acquire(void* self) {
	auto&& __r = ((QSystemSemaphore*)self)->acquire();
	return __r;
}
int QSystemSemaphore_release(void* self) {
	auto&& __r = ((QSystemSemaphore*)self)->release();
	return __r;
}
int QSystemSemaphore_error(void* self) {
	auto&& __r = ((QSystemSemaphore*)self)->error();
	return (int)(__r);
}
char* QSystemSemaphore_errorString(void* self, long long* __out_len) {
	auto&& __r = ((QSystemSemaphore*)self)->errorString();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QSystemSemaphore_isKeyTypeSupported(int type_) {
	auto&& __r = QSystemSemaphore::isKeyTypeSupported((QNativeIpcKey::Type)(type_));
	return __r;
}
void* QSystemSemaphore_platformSafeKey(const char* key, long long key_len) {
	QString key_qstr = QString::fromUtf8(key, key_len);
	auto&& __r = QSystemSemaphore::platformSafeKey(key_qstr);
	return new QNativeIpcKey(__r);
}
void* QSystemSemaphore_legacyNativeKey(const char* key, long long key_len) {
	QString key_qstr = QString::fromUtf8(key, key_len);
	auto&& __r = QSystemSemaphore::legacyNativeKey(key_qstr);
	return new QNativeIpcKey(__r);
}
char* QSystemSemaphore_tr2(const char* sourceText, const char* disambiguation, long long* __out_len) {
	auto&& __r = QSystemSemaphore::tr(sourceText, disambiguation);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QSystemSemaphore_tr3(const char* sourceText, const char* disambiguation, int n, long long* __out_len) {
	auto&& __r = QSystemSemaphore::tr(sourceText, disambiguation, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QSystemSemaphore_setNativeKey3(void* self, void* key, int initialValue) {
	((QSystemSemaphore*)self)->setNativeKey(*(QNativeIpcKey*)key, initialValue);
}
void QSystemSemaphore_setNativeKey4(void* self, void* key, int initialValue, int arg3) {
	((QSystemSemaphore*)self)->setNativeKey(*(QNativeIpcKey*)key, initialValue, (QSystemSemaphore::AccessMode)(arg3));
}
void QSystemSemaphore_setNativeKey5(void* self, const char* key, long long key_len, int initialValue) {
	QString key_qstr = QString::fromUtf8(key, key_len);
	((QSystemSemaphore*)self)->setNativeKey(key_qstr, initialValue);
}
void QSystemSemaphore_setNativeKey6(void* self, const char* key, long long key_len, int initialValue, int mode) {
	QString key_qstr = QString::fromUtf8(key, key_len);
	((QSystemSemaphore*)self)->setNativeKey(key_qstr, initialValue, (QSystemSemaphore::AccessMode)(mode));
}
void QSystemSemaphore_setNativeKey7(void* self, const char* key, long long key_len, int initialValue, int mode, int type_) {
	QString key_qstr = QString::fromUtf8(key, key_len);
	((QSystemSemaphore*)self)->setNativeKey(key_qstr, initialValue, (QSystemSemaphore::AccessMode)(mode), (QNativeIpcKey::Type)(type_));
}
void QSystemSemaphore_setKey2(void* self, const char* key, long long key_len, int initialValue) {
	QString key_qstr = QString::fromUtf8(key, key_len);
	((QSystemSemaphore*)self)->setKey(key_qstr, initialValue);
}
void QSystemSemaphore_setKey3(void* self, const char* key, long long key_len, int initialValue, int mode) {
	QString key_qstr = QString::fromUtf8(key, key_len);
	((QSystemSemaphore*)self)->setKey(key_qstr, initialValue, (QSystemSemaphore::AccessMode)(mode));
}
int QSystemSemaphore_release2(void* self, int n) {
	auto&& __r = ((QSystemSemaphore*)self)->release(n);
	return __r;
}
void* QSystemSemaphore_platformSafeKey2(const char* key, long long key_len, int type_) {
	QString key_qstr = QString::fromUtf8(key, key_len);
	auto&& __r = QSystemSemaphore::platformSafeKey(key_qstr, (QNativeIpcKey::Type)(type_));
	return new QNativeIpcKey(__r);
}
void* QSystemSemaphore_legacyNativeKey2(const char* key, long long key_len, int type_) {
	QString key_qstr = QString::fromUtf8(key, key_len);
	auto&& __r = QSystemSemaphore::legacyNativeKey(key_qstr, (QNativeIpcKey::Type)(type_));
	return new QNativeIpcKey(__r);
}
void QSystemSemaphore_delete(void* self) { delete (QSystemSemaphore*)self; }

}

#include <QtGui/qpolygon.h>
#include <QtCore/qpoint.h>
#include <QtCore/qrect.h>
#include <QtCore/qvariant.h>
#include <QtCore/qnamespace.h>
#include <QtGui/qtransform.h>
#include "gen_qpolygonf.h"

extern "C" {

void* QPolygonF_new() {
	return new QPolygonF();
}
void* QPolygonF_new2(void** v, long long v_len) {
	QList<QPointF> v_list; v_list.reserve(v_len);
	for (long long __i = 0; __i < v_len; __i++) v_list.append(*(QPointF*)v[__i]);
	return new QPolygonF(v_list);
}
void* QPolygonF_new3(void* r) {
	return new QPolygonF(*(QRectF*)r);
}
void* QPolygonF_new4(void* a) {
	return new QPolygonF(*(QPolygon*)a);
}
void* QPolygonF_new5(void* arg1) {
	return new QPolygonF(*(QPolygonF*)arg1);
}
void QPolygonF_swap(void* self, void* other) {
	((QPolygonF*)self)->swap(*(QPolygonF*)other);
}
void* QPolygonF_ToQVariant(void* self) {
	auto&& __r = static_cast<QVariant>(*(QPolygonF*)self);
	return new QVariant(__r);
}
void QPolygonF_translate(void* self, double dx, double dy) {
	((QPolygonF*)self)->translate(dx, dy);
}
void QPolygonF_translate2(void* self, void* offset) {
	((QPolygonF*)self)->translate(*(QPointF*)offset);
}
void* QPolygonF_translated(void* self, double dx, double dy) {
	auto&& __r = ((QPolygonF*)self)->translated(dx, dy);
	return new QPolygonF(__r);
}
void* QPolygonF_translated2(void* self, void* offset) {
	auto&& __r = ((QPolygonF*)self)->translated(*(QPointF*)offset);
	return new QPolygonF(__r);
}
void* QPolygonF_toPolygon(void* self) {
	auto&& __r = ((QPolygonF*)self)->toPolygon();
	return new QPolygon(__r);
}
int QPolygonF_isClosed(void* self) {
	auto&& __r = ((QPolygonF*)self)->isClosed();
	return __r;
}
void* QPolygonF_boundingRect(void* self) {
	auto&& __r = ((QPolygonF*)self)->boundingRect();
	return new QRectF(__r);
}
int QPolygonF_containsPoint(void* self, void* pt, int fillRule) {
	auto&& __r = ((QPolygonF*)self)->containsPoint(*(QPointF*)pt, (Qt::FillRule)(fillRule));
	return __r;
}
void* QPolygonF_united(void* self, void* r) {
	auto&& __r = ((QPolygonF*)self)->united(*(QPolygonF*)r);
	return new QPolygonF(__r);
}
void* QPolygonF_intersected(void* self, void* r) {
	auto&& __r = ((QPolygonF*)self)->intersected(*(QPolygonF*)r);
	return new QPolygonF(__r);
}
void* QPolygonF_subtracted(void* self, void* r) {
	auto&& __r = ((QPolygonF*)self)->subtracted(*(QPolygonF*)r);
	return new QPolygonF(__r);
}
int QPolygonF_intersects(void* self, void* r) {
	auto&& __r = ((QPolygonF*)self)->intersects(*(QPolygonF*)r);
	return __r;
}
long long QPolygonF_maxSize() {
	auto&& __r = QPolygonF::maxSize();
	return __r;
}
long long QPolygonF_size(void* self) {
	auto&& __r = ((QPolygonF*)self)->size();
	return __r;
}
long long QPolygonF_count(void* self) {
	auto&& __r = ((QPolygonF*)self)->count();
	return __r;
}
long long QPolygonF_length(void* self) {
	auto&& __r = ((QPolygonF*)self)->length();
	return __r;
}
int QPolygonF_isEmpty(void* self) {
	auto&& __r = ((QPolygonF*)self)->isEmpty();
	return __r;
}
void QPolygonF_resize(void* self, long long size) {
	((QPolygonF*)self)->resize(size);
}
void QPolygonF_resize2(void* self, long long size, void* c) {
	((QPolygonF*)self)->resize(size, *(QPointF*)c);
}
void QPolygonF_resizeForOverwrite(void* self, long long size) {
	((QPolygonF*)self)->resizeForOverwrite(size);
}
long long QPolygonF_capacity(void* self) {
	auto&& __r = ((QPolygonF*)self)->capacity();
	return __r;
}
void QPolygonF_reserve(void* self, long long size) {
	((QPolygonF*)self)->reserve(size);
}
void QPolygonF_squeeze(void* self) {
	((QPolygonF*)self)->squeeze();
}
void QPolygonF_detach(void* self) {
	((QPolygonF*)self)->detach();
}
int QPolygonF_isDetached(void* self) {
	auto&& __r = ((QPolygonF*)self)->isDetached();
	return __r;
}
int QPolygonF_isSharedWith(void* self, void** other, long long other_len) {
	QList<QPointF> other_list; other_list.reserve(other_len);
	for (long long __i = 0; __i < other_len; __i++) other_list.append(*(QPointF*)other[__i]);
	auto&& __r = ((QPolygonF*)self)->isSharedWith(other_list);
	return __r;
}
void* QPolygonF_data(void* self) {
	auto&& __r = ((QPolygonF*)self)->data();
	return (void*)(__r);
}
void* QPolygonF_data2(void* self) {
	auto&& __r = ((QPolygonF*)self)->data();
	return (void*)(__r);
}
void* QPolygonF_constData(void* self) {
	auto&& __r = ((QPolygonF*)self)->constData();
	return (void*)(__r);
}
void QPolygonF_clear(void* self) {
	((QPolygonF*)self)->clear();
}
void* QPolygonF_at(void* self, long long i) {
	auto&& __r = ((QPolygonF*)self)->at(i);
	return (void*)&(__r);
}
void* QPolygonF_OperatorSubscript(void* self, long long i) {
	auto&& __r = ((QPolygonF*)self)->operator[](i);
	return (void*)&(__r);
}
void* QPolygonF_OperatorSubscript2(void* self, long long i) {
	auto&& __r = ((QPolygonF*)self)->operator[](i);
	return (void*)&(__r);
}
void QPolygonF_append(void* self, void* t) {
	((QPolygonF*)self)->append(*(QPointF*)t);
}
void QPolygonF_append2(void* self, void** l, long long l_len) {
	QList<QPointF> l_list; l_list.reserve(l_len);
	for (long long __i = 0; __i < l_len; __i++) l_list.append(*(QPointF*)l[__i]);
	((QPolygonF*)self)->append(l_list);
}
void QPolygonF_prepend(void* self, void* t) {
	((QPolygonF*)self)->prepend(*(QPointF*)t);
}
void QPolygonF_replace(void* self, long long i, void* t) {
	((QPolygonF*)self)->replace(i, *(QPointF*)t);
}
void QPolygonF_remove(void* self, long long i) {
	((QPolygonF*)self)->remove(i);
}
void QPolygonF_removeFirst(void* self) {
	((QPolygonF*)self)->removeFirst();
}
void QPolygonF_removeLast(void* self) {
	((QPolygonF*)self)->removeLast();
}
void* QPolygonF_takeFirst(void* self) {
	auto&& __r = ((QPolygonF*)self)->takeFirst();
	return new QPointF(__r);
}
void* QPolygonF_takeLast(void* self) {
	auto&& __r = ((QPolygonF*)self)->takeLast();
	return new QPointF(__r);
}
void QPolygonF_removeAt(void* self, long long i) {
	((QPolygonF*)self)->removeAt(i);
}
void* QPolygonF_takeAt(void* self, long long i) {
	auto&& __r = ((QPolygonF*)self)->takeAt(i);
	return new QPointF(__r);
}
void QPolygonF_move(void* self, long long from, long long to) {
	((QPolygonF*)self)->move(from, to);
}
void* QPolygonF_first(void* self) {
	auto&& __r = ((QPolygonF*)self)->first();
	return (void*)&(__r);
}
void* QPolygonF_first2(void* self) {
	auto&& __r = ((QPolygonF*)self)->first();
	return (void*)&(__r);
}
void* QPolygonF_constFirst(void* self) {
	auto&& __r = ((QPolygonF*)self)->constFirst();
	return (void*)&(__r);
}
void* QPolygonF_last(void* self) {
	auto&& __r = ((QPolygonF*)self)->last();
	return (void*)&(__r);
}
void* QPolygonF_last2(void* self) {
	auto&& __r = ((QPolygonF*)self)->last();
	return (void*)&(__r);
}
void* QPolygonF_constLast(void* self) {
	auto&& __r = ((QPolygonF*)self)->constLast();
	return (void*)&(__r);
}
int QPolygonF_startsWith(void* self, void* t) {
	auto&& __r = ((QPolygonF*)self)->startsWith(*(QPointF*)t);
	return __r;
}
int QPolygonF_endsWith(void* self, void* t) {
	auto&& __r = ((QPolygonF*)self)->endsWith(*(QPointF*)t);
	return __r;
}
void** QPolygonF_mid(void* self, long long pos, long long* __out_len) {
	auto&& __r = ((QPolygonF*)self)->mid(pos);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QPointF(__e);
	return __arr;
}
void** QPolygonF_first3(void* self, long long n, long long* __out_len) {
	auto&& __r = ((QPolygonF*)self)->first(n);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QPointF(__e);
	return __arr;
}
void** QPolygonF_last3(void* self, long long n, long long* __out_len) {
	auto&& __r = ((QPolygonF*)self)->last(n);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QPointF(__e);
	return __arr;
}
void** QPolygonF_sliced(void* self, long long pos, long long* __out_len) {
	auto&& __r = ((QPolygonF*)self)->sliced(pos);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QPointF(__e);
	return __arr;
}
void** QPolygonF_sliced2(void* self, long long pos, long long n, long long* __out_len) {
	auto&& __r = ((QPolygonF*)self)->sliced(pos, n);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QPointF(__e);
	return __arr;
}
void* QPolygonF_value(void* self, long long i) {
	auto&& __r = ((QPolygonF*)self)->value(i);
	return new QPointF(__r);
}
void* QPolygonF_value2(void* self, long long i, void* defaultValue) {
	auto&& __r = ((QPolygonF*)self)->value(i, *(QPointF*)defaultValue);
	return new QPointF(__r);
}
void QPolygonF_swapItemsAt(void* self, long long i, long long j) {
	((QPolygonF*)self)->swapItemsAt(i, j);
}
void QPolygonF_push_back(void* self, void* t) {
	((QPolygonF*)self)->push_back(*(QPointF*)t);
}
void QPolygonF_push_front(void* self, void* t) {
	((QPolygonF*)self)->push_front(*(QPointF*)t);
}
void QPolygonF_pop_back(void* self) {
	((QPolygonF*)self)->pop_back();
}
void QPolygonF_pop_front(void* self) {
	((QPolygonF*)self)->pop_front();
}
int QPolygonF_empty(void* self) {
	auto&& __r = ((QPolygonF*)self)->empty();
	return __r;
}
void* QPolygonF_front(void* self) {
	auto&& __r = ((QPolygonF*)self)->front();
	return (void*)&(__r);
}
void* QPolygonF_front2(void* self) {
	auto&& __r = ((QPolygonF*)self)->front();
	return (void*)&(__r);
}
void* QPolygonF_back(void* self) {
	auto&& __r = ((QPolygonF*)self)->back();
	return (void*)&(__r);
}
void* QPolygonF_back2(void* self) {
	auto&& __r = ((QPolygonF*)self)->back();
	return (void*)&(__r);
}
void QPolygonF_shrink_to_fit(void* self) {
	((QPolygonF*)self)->shrink_to_fit();
}
long long QPolygonF_max_size(void* self) {
	auto&& __r = ((QPolygonF*)self)->max_size();
	return __r;
}
void** QPolygonF_OperatorPlus(void* self, void** l, long long l_len, long long* __out_len) {
	QList<QPointF> l_list; l_list.reserve(l_len);
	for (long long __i = 0; __i < l_len; __i++) l_list.append(*(QPointF*)l[__i]);
	auto&& __r = ((QPolygonF*)self)->operator+(l_list);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QPointF(__e);
	return __arr;
}
void** QPolygonF_OperatorPlus2(void* self, void** l, long long l_len, long long* __out_len) {
	QList<QPointF> l_list; l_list.reserve(l_len);
	for (long long __i = 0; __i < l_len; __i++) l_list.append(*(QPointF*)l[__i]);
	auto&& __r = ((QPolygonF*)self)->operator+(l_list);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QPointF(__e);
	return __arr;
}
void** QPolygonF_fromList(void** list, long long list_len, long long* __out_len) {
	QList<QPointF> list_list; list_list.reserve(list_len);
	for (long long __i = 0; __i < list_len; __i++) list_list.append(*(QPointF*)list[__i]);
	auto&& __r = QPolygonF::fromList(list_list);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QPointF(__e);
	return __arr;
}
void** QPolygonF_toList(void* self, long long* __out_len) {
	auto&& __r = ((QPolygonF*)self)->toList();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QPointF(__e);
	return __arr;
}
void** QPolygonF_fromVector(void** vector, long long vector_len, long long* __out_len) {
	QList<QPointF> vector_list; vector_list.reserve(vector_len);
	for (long long __i = 0; __i < vector_len; __i++) vector_list.append(*(QPointF*)vector[__i]);
	auto&& __r = QPolygonF::fromVector(vector_list);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QPointF(__e);
	return __arr;
}
void** QPolygonF_toVector(void* self, long long* __out_len) {
	auto&& __r = ((QPolygonF*)self)->toVector();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QPointF(__e);
	return __arr;
}
void QPolygonF_remove2(void* self, long long i, long long n) {
	((QPolygonF*)self)->remove(i, n);
}
void** QPolygonF_mid2(void* self, long long pos, long long len_, long long* __out_len) {
	auto&& __r = ((QPolygonF*)self)->mid(pos, len_);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QPointF(__e);
	return __arr;
}
void* QPolygonF_OperatorMultiply(void* self, void* m) {
	auto&& __r = (*(QPolygonF*)self) * (*(QTransform*)m);
	return new QPolygonF(__r);
}
void* QPolygonF_toQVariant(void* self) {
	auto&& __r = QVariant::fromValue(*(QPolygonF*)self);
	return new QVariant(__r);
}
void QPolygonF_delete(void* self) { delete (QPolygonF*)self; }

}

#include <QtGui/qevent.h>
#include <QtCore/qpoint.h>
#include <QtCore/qnamespace.h>
#include <QtGui/qpointingdevice.h>
#include <QtCore/qobject.h>
#include <QtGui/qeventpoint.h>
#include <QtGui/qinputdevice.h>
#include <QtCore/qcoreevent.h>
#include "gen_qwheelevent.h"

extern "C" {

void* QWheelEvent_base(void* self) { return (void*)static_cast<QSinglePointEvent*>((QWheelEvent*)self); }
void QWheelEvent_chain(void* self, void** out) {
	QWheelEvent* __p = (QWheelEvent*)self;
	QSinglePointEvent* __b0 = static_cast<QSinglePointEvent*>(__p); out[0] = (void*)__b0;
	QPointerEvent* __b1 = static_cast<QPointerEvent*>(__b0); out[1] = (void*)__b1;
	QInputEvent* __b2 = static_cast<QInputEvent*>(__b1); out[2] = (void*)__b2;
	QEvent* __b3 = static_cast<QEvent*>(__b2); out[3] = (void*)__b3;
}
void* QWheelEvent_new(void* pos, void* globalPos, void* pixelDelta, void* angleDelta, int buttons, int modifiers, int phase, int inverted) {
	return new QWheelEvent(*(QPointF*)pos, *(QPointF*)globalPos, *(QPoint*)pixelDelta, *(QPoint*)angleDelta, QFlags<Qt::MouseButton>(buttons), QFlags<Qt::KeyboardModifier>(modifiers), (Qt::ScrollPhase)(phase), inverted);
}
void* QWheelEvent_new2(void* pos, void* globalPos, void* pixelDelta, void* angleDelta, int buttons, int modifiers, int phase, int inverted, int source) {
	return new QWheelEvent(*(QPointF*)pos, *(QPointF*)globalPos, *(QPoint*)pixelDelta, *(QPoint*)angleDelta, QFlags<Qt::MouseButton>(buttons), QFlags<Qt::KeyboardModifier>(modifiers), (Qt::ScrollPhase)(phase), inverted, (Qt::MouseEventSource)(source));
}
void* QWheelEvent_new3(void* pos, void* globalPos, void* pixelDelta, void* angleDelta, int buttons, int modifiers, int phase, int inverted, int source, void* device) {
	return new QWheelEvent(*(QPointF*)pos, *(QPointF*)globalPos, *(QPoint*)pixelDelta, *(QPoint*)angleDelta, QFlags<Qt::MouseButton>(buttons), QFlags<Qt::KeyboardModifier>(modifiers), (Qt::ScrollPhase)(phase), inverted, (Qt::MouseEventSource)(source), (QPointingDevice*)device);
}
void* QWheelEvent_pixelDelta(void* self) {
	auto&& __r = ((QWheelEvent*)self)->pixelDelta();
	return new QPoint(__r);
}
void* QWheelEvent_angleDelta(void* self) {
	auto&& __r = ((QWheelEvent*)self)->angleDelta();
	return new QPoint(__r);
}
int QWheelEvent_phase(void* self) {
	auto&& __r = ((QWheelEvent*)self)->phase();
	return (int)(__r);
}
int QWheelEvent_inverted(void* self) {
	auto&& __r = ((QWheelEvent*)self)->inverted();
	return __r;
}
int QWheelEvent_isInverted(void* self) {
	auto&& __r = ((QWheelEvent*)self)->isInverted();
	return __r;
}
int QWheelEvent_hasPixelDelta(void* self) {
	auto&& __r = ((QWheelEvent*)self)->hasPixelDelta();
	return __r;
}
int QWheelEvent_source(void* self) {
	auto&& __r = ((QWheelEvent*)self)->source();
	return (int)(__r);
}
int QWheelEvent_registerEventType() {
	auto&& __r = QWheelEvent::registerEventType();
	return __r;
}
int QWheelEvent_registerEventType2(int hint) {
	auto&& __r = QWheelEvent::registerEventType(hint);
	return __r;
}
void QWheelEvent_delete(void* self) { delete (QWheelEvent*)self; }

}

#include <QtGui/qevent.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qstring.h>
#include "gen_qwhatsthisclickedevent.h"

extern "C" {

void* QWhatsThisClickedEvent_base(void* self) { return (void*)static_cast<QEvent*>((QWhatsThisClickedEvent*)self); }
void* QWhatsThisClickedEvent_new(const char* href, long long href_len) {
	QString href_qstr = QString::fromUtf8(href, href_len);
	return new QWhatsThisClickedEvent(href_qstr);
}
char* QWhatsThisClickedEvent_href(void* self, long long* __out_len) {
	auto&& __r = ((QWhatsThisClickedEvent*)self)->href();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QWhatsThisClickedEvent_registerEventType() {
	auto&& __r = QWhatsThisClickedEvent::registerEventType();
	return __r;
}
int QWhatsThisClickedEvent_registerEventType2(int hint) {
	auto&& __r = QWhatsThisClickedEvent::registerEventType(hint);
	return __r;
}
void QWhatsThisClickedEvent_delete(void* self) { delete (QWhatsThisClickedEvent*)self; }

}

#include <QtGui/qbrush.h>
#include <QtCore/qpoint.h>
#include <QtGui/qcolor.h>
#include "gen_qradialgradient.h"

extern "C" {

void* QRadialGradient_base(void* self) { return (void*)static_cast<QGradient*>((QRadialGradient*)self); }
void* QRadialGradient_new() {
	return new QRadialGradient();
}
void* QRadialGradient_new2(void* center, double radius, void* focalPoint) {
	return new QRadialGradient(*(QPointF*)center, radius, *(QPointF*)focalPoint);
}
void* QRadialGradient_new3(double cx, double cy, double radius, double fx, double fy) {
	return new QRadialGradient(cx, cy, radius, fx, fy);
}
void* QRadialGradient_new4(void* center, double radius) {
	return new QRadialGradient(*(QPointF*)center, radius);
}
void* QRadialGradient_new5(double cx, double cy, double radius) {
	return new QRadialGradient(cx, cy, radius);
}
void* QRadialGradient_new6(void* center, double centerRadius, void* focalPoint, double focalRadius) {
	return new QRadialGradient(*(QPointF*)center, centerRadius, *(QPointF*)focalPoint, focalRadius);
}
void* QRadialGradient_new7(double cx, double cy, double centerRadius, double fx, double fy, double focalRadius) {
	return new QRadialGradient(cx, cy, centerRadius, fx, fy, focalRadius);
}
void* QRadialGradient_new8(void* arg1) {
	return new QRadialGradient(*(QRadialGradient*)arg1);
}
void* QRadialGradient_center(void* self) {
	auto&& __r = ((QRadialGradient*)self)->center();
	return new QPointF(__r);
}
void QRadialGradient_setCenter(void* self, void* center) {
	((QRadialGradient*)self)->setCenter(*(QPointF*)center);
}
void QRadialGradient_setCenter2(void* self, double x, double y) {
	((QRadialGradient*)self)->setCenter(x, y);
}
void* QRadialGradient_focalPoint(void* self) {
	auto&& __r = ((QRadialGradient*)self)->focalPoint();
	return new QPointF(__r);
}
void QRadialGradient_setFocalPoint(void* self, void* focalPoint) {
	((QRadialGradient*)self)->setFocalPoint(*(QPointF*)focalPoint);
}
void QRadialGradient_setFocalPoint2(void* self, double x, double y) {
	((QRadialGradient*)self)->setFocalPoint(x, y);
}
double QRadialGradient_radius(void* self) {
	auto&& __r = ((QRadialGradient*)self)->radius();
	return __r;
}
void QRadialGradient_setRadius(void* self, double radius) {
	((QRadialGradient*)self)->setRadius(radius);
}
double QRadialGradient_centerRadius(void* self) {
	auto&& __r = ((QRadialGradient*)self)->centerRadius();
	return __r;
}
void QRadialGradient_setCenterRadius(void* self, double radius) {
	((QRadialGradient*)self)->setCenterRadius(radius);
}
double QRadialGradient_focalRadius(void* self) {
	auto&& __r = ((QRadialGradient*)self)->focalRadius();
	return __r;
}
void QRadialGradient_setFocalRadius(void* self, double radius) {
	((QRadialGradient*)self)->setFocalRadius(radius);
}
void QRadialGradient_delete(void* self) { delete (QRadialGradient*)self; }

}

#include <QtGui/qaccessibilityhints.h>
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
#include "gen_qaccessibilityhints.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QAccessibilityHints__Go : public QAccessibilityHints {
public:
	long long __goId;
	unsigned char __goOn[8] = {};
	template<typename... A> QAccessibilityHints__Go(long long id, A&&... a) : QAccessibilityHints(std::forward<A>(a)...), __goId(id) {}
	using QAccessibilityHints::sender;
	using QAccessibilityHints::senderSignalIndex;
	using QAccessibilityHints::receivers;
	using QAccessibilityHints::isSignalConnected;
	bool cppgoBase_event(QEvent* __a0) { return QAccessibilityHints::event(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QAccessibilityHints::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QAccessibilityHints::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QAccessibilityHints::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QAccessibilityHints::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QAccessibilityHints::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QAccessibilityHints::metaObject();
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QAccessibilityHints::event(event);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QAccessibilityHints::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QAccessibilityHints::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QAccessibilityHints::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QAccessibilityHints::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QAccessibilityHints::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QAccessibilityHints::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QAccessibilityHints;
static std::mutex __connsMu_QAccessibilityHints;

extern "C" {

void* QAccessibilityHints_base(void* self) { return (void*)static_cast<QObject*>((QAccessibilityHints*)self); }
void QAccessibilityHints__set_override(void* self, int idx) { ((QAccessibilityHints__Go*)self)->__goOn[idx] = 1; }
void* QAccessibilityHints_fromQObject(void* o) { return (void*)qobject_cast<QAccessibilityHints*>((QObject*)o); }
void* QAccessibilityHints_new(long long __goid) {
	return new QAccessibilityHints__Go(__goid);
}
void* QAccessibilityHints_new2(long long __goid, void* parent) {
	return new QAccessibilityHints__Go(__goid, (QObject*)parent);
}
char* QAccessibilityHints_tr(const char* s, long long* __out_len) {
	auto&& __r = QAccessibilityHints::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QAccessibilityHints_contrastPreference(void* self) {
	auto&& __r = ((QAccessibilityHints*)self)->contrastPreference();
	return (int)(__r);
}
void QAccessibilityHints_contrastPreferenceChanged(void* self, int contrastPreference) {
	((QAccessibilityHints*)self)->contrastPreferenceChanged((Qt::ContrastPreference)(contrastPreference));
}
void* QAccessibilityHints_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QAccessibilityHints::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QAccessibilityHints_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QAccessibilityHints::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QAccessibilityHints_sender(void* self) {
	auto&& __r = ((QAccessibilityHints__Go*)self)->sender();
	return (void*)(__r);
}
int QAccessibilityHints_senderSignalIndex(void* self) {
	auto&& __r = ((QAccessibilityHints__Go*)self)->senderSignalIndex();
	return __r;
}
int QAccessibilityHints_receivers(void* self, const char* signal) {
	auto&& __r = ((QAccessibilityHints__Go*)self)->receivers(signal);
	return __r;
}
int QAccessibilityHints_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QAccessibilityHints__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QAccessibilityHints_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QAccessibilityHints::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QAccessibilityHints_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QAccessibilityHints::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QAccessibilityHints_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QAccessibilityHints::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QAccessibilityHints_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QAccessibilityHints::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QAccessibilityHints_metaObject_default(void* self) {
	auto&& __r = ((QAccessibilityHints*)self)->QAccessibilityHints::metaObject();
	return (void*)(__r);
}
int QAccessibilityHints_event_default(void* self, void* event) {
	auto&& __r = ((QAccessibilityHints__Go*)self)->cppgoBase_event((QEvent*)event);
	return __r;
}
int QAccessibilityHints_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QAccessibilityHints*)self)->QAccessibilityHints::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QAccessibilityHints_timerEvent_default(void* self, void* event) {
	((QAccessibilityHints__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QAccessibilityHints_childEvent_default(void* self, void* event) {
	((QAccessibilityHints__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QAccessibilityHints_customEvent_default(void* self, void* event) {
	((QAccessibilityHints__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QAccessibilityHints_connectNotify_default(void* self, void* signal) {
	((QAccessibilityHints__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QAccessibilityHints_disconnectNotify_default(void* self, void* signal) {
	((QAccessibilityHints__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QAccessibilityHints_connect_contrastPreferenceChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAccessibilityHints*)self, &QAccessibilityHints::contrastPreferenceChanged, (QAccessibilityHints*)self, [__id](Qt::ContrastPreference contrastPreference){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, (int)(contrastPreference));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAccessibilityHints);
	__conns_QAccessibilityHints[__id] = __c;
}
void QAccessibilityHints_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAccessibilityHints*)self, &QAccessibilityHints::destroyed, (QAccessibilityHints*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAccessibilityHints);
	__conns_QAccessibilityHints[__id] = __c;
}
void QAccessibilityHints_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAccessibilityHints*)self, &QAccessibilityHints::destroyed, (QAccessibilityHints*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAccessibilityHints);
	__conns_QAccessibilityHints[__id] = __c;
}
void QAccessibilityHints_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAccessibilityHints*)self, &QAccessibilityHints::objectNameChanged, (QAccessibilityHints*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAccessibilityHints);
	__conns_QAccessibilityHints[__id] = __c;
}
void QAccessibilityHints__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QAccessibilityHints);
		auto __it = __conns_QAccessibilityHints.find(id);
		if (__it == __conns_QAccessibilityHints.end()) return;
		__c = __it->second;
		__conns_QAccessibilityHints.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QAccessibilityHints__Go_new(long long id, void* parent) { return new QAccessibilityHints__Go(id, (QObject*)parent); }
void QAccessibilityHints_delete(void* self) { delete (QAccessibilityHints*)self; }

}

#include <QtGui/qaccessible.h>
#include <QtCore/qobject.h>
#include <QtCore/qstring.h>
#include <QtGui/qaccessible_base.h>
#include "gen_qaccessibleannouncementevent.h"

extern "C" {

void* QAccessibleAnnouncementEvent_base(void* self) { return (void*)static_cast<QAccessibleEvent*>((QAccessibleAnnouncementEvent*)self); }
void* QAccessibleAnnouncementEvent_new(void* object, const char* message, long long message_len) {
	QString message_qstr = QString::fromUtf8(message, message_len);
	return new QAccessibleAnnouncementEvent((QObject*)object, message_qstr);
}
void* QAccessibleAnnouncementEvent_new2(void* iface, const char* message, long long message_len) {
	QString message_qstr = QString::fromUtf8(message, message_len);
	return new QAccessibleAnnouncementEvent((QAccessibleInterface*)iface, message_qstr);
}
char* QAccessibleAnnouncementEvent_message(void* self, long long* __out_len) {
	auto&& __r = ((QAccessibleAnnouncementEvent*)self)->message();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QAccessibleAnnouncementEvent_politeness(void* self) {
	auto&& __r = ((QAccessibleAnnouncementEvent*)self)->politeness();
	return (int)(__r);
}
void QAccessibleAnnouncementEvent_setPoliteness(void* self, int politeness) {
	((QAccessibleAnnouncementEvent*)self)->setPoliteness((QAccessible::AnnouncementPoliteness)(politeness));
}
void QAccessibleAnnouncementEvent_delete(void* self) { delete (QAccessibleAnnouncementEvent*)self; }

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
#include "gen_qfontmetricsf.h"

extern "C" {

void* QFontMetricsF_new(void* font) {
	return new QFontMetricsF(*(QFont*)font);
}
void* QFontMetricsF_new2(void* font, void* pd) {
	return new QFontMetricsF(*(QFont*)font, (QPaintDevice*)pd);
}
void* QFontMetricsF_new3(void* arg1) {
	return new QFontMetricsF(*(QFontMetrics*)arg1);
}
void* QFontMetricsF_new4(void* arg1) {
	return new QFontMetricsF(*(QFontMetricsF*)arg1);
}
void* QFontMetricsF_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QFontMetricsF*)self)->operator=(*(QFontMetricsF*)arg1);
	return (void*)&(__r);
}
void* QFontMetricsF_OperatorAssign2(void* self, void* arg1) {
	auto&& __r = ((QFontMetricsF*)self)->operator=(*(QFontMetrics*)arg1);
	return (void*)&(__r);
}
void QFontMetricsF_swap(void* self, void* other) {
	((QFontMetricsF*)self)->swap(*(QFontMetricsF*)other);
}
double QFontMetricsF_ascent(void* self) {
	auto&& __r = ((QFontMetricsF*)self)->ascent();
	return __r;
}
double QFontMetricsF_capHeight(void* self) {
	auto&& __r = ((QFontMetricsF*)self)->capHeight();
	return __r;
}
double QFontMetricsF_descent(void* self) {
	auto&& __r = ((QFontMetricsF*)self)->descent();
	return __r;
}
double QFontMetricsF_height(void* self) {
	auto&& __r = ((QFontMetricsF*)self)->height();
	return __r;
}
double QFontMetricsF_leading(void* self) {
	auto&& __r = ((QFontMetricsF*)self)->leading();
	return __r;
}
double QFontMetricsF_lineSpacing(void* self) {
	auto&& __r = ((QFontMetricsF*)self)->lineSpacing();
	return __r;
}
double QFontMetricsF_minLeftBearing(void* self) {
	auto&& __r = ((QFontMetricsF*)self)->minLeftBearing();
	return __r;
}
double QFontMetricsF_minRightBearing(void* self) {
	auto&& __r = ((QFontMetricsF*)self)->minRightBearing();
	return __r;
}
double QFontMetricsF_maxWidth(void* self) {
	auto&& __r = ((QFontMetricsF*)self)->maxWidth();
	return __r;
}
double QFontMetricsF_xHeight(void* self) {
	auto&& __r = ((QFontMetricsF*)self)->xHeight();
	return __r;
}
double QFontMetricsF_averageCharWidth(void* self) {
	auto&& __r = ((QFontMetricsF*)self)->averageCharWidth();
	return __r;
}
int QFontMetricsF_inFont(void* self, void* arg1) {
	auto&& __r = ((QFontMetricsF*)self)->inFont(*(QChar*)arg1);
	return __r;
}
int QFontMetricsF_inFontUcs4(void* self, unsigned int ucs4) {
	auto&& __r = ((QFontMetricsF*)self)->inFontUcs4(ucs4);
	return __r;
}
double QFontMetricsF_leftBearing(void* self, void* arg1) {
	auto&& __r = ((QFontMetricsF*)self)->leftBearing(*(QChar*)arg1);
	return __r;
}
double QFontMetricsF_rightBearing(void* self, void* arg1) {
	auto&& __r = ((QFontMetricsF*)self)->rightBearing(*(QChar*)arg1);
	return __r;
}
double QFontMetricsF_horizontalAdvance(void* self, const char* string_, long long string__len) {
	QString string__qstr = QString::fromUtf8(string_, string__len);
	auto&& __r = ((QFontMetricsF*)self)->horizontalAdvance(string__qstr);
	return __r;
}
double QFontMetricsF_horizontalAdvance2(void* self, void* arg1) {
	auto&& __r = ((QFontMetricsF*)self)->horizontalAdvance(*(QChar*)arg1);
	return __r;
}
double QFontMetricsF_horizontalAdvance3(void* self, const char* string_, long long string__len, void* textOption) {
	QString string__qstr = QString::fromUtf8(string_, string__len);
	auto&& __r = ((QFontMetricsF*)self)->horizontalAdvance(string__qstr, *(QTextOption*)textOption);
	return __r;
}
void* QFontMetricsF_boundingRect(void* self, const char* string_, long long string__len) {
	QString string__qstr = QString::fromUtf8(string_, string__len);
	auto&& __r = ((QFontMetricsF*)self)->boundingRect(string__qstr);
	return new QRectF(__r);
}
void* QFontMetricsF_boundingRect2(void* self, const char* text, long long text_len, void* textOption) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = ((QFontMetricsF*)self)->boundingRect(text_qstr, *(QTextOption*)textOption);
	return new QRectF(__r);
}
void* QFontMetricsF_boundingRect3(void* self, void* arg1) {
	auto&& __r = ((QFontMetricsF*)self)->boundingRect(*(QChar*)arg1);
	return new QRectF(__r);
}
void* QFontMetricsF_boundingRect4(void* self, void* r, int flags, const char* string_, long long string__len) {
	QString string__qstr = QString::fromUtf8(string_, string__len);
	auto&& __r = ((QFontMetricsF*)self)->boundingRect(*(QRectF*)r, flags, string__qstr);
	return new QRectF(__r);
}
void* QFontMetricsF_size(void* self, int flags, const char* str, long long str_len) {
	QString str_qstr = QString::fromUtf8(str, str_len);
	auto&& __r = ((QFontMetricsF*)self)->size(flags, str_qstr);
	return new QSizeF(__r);
}
void* QFontMetricsF_tightBoundingRect(void* self, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = ((QFontMetricsF*)self)->tightBoundingRect(text_qstr);
	return new QRectF(__r);
}
void* QFontMetricsF_tightBoundingRect2(void* self, const char* text, long long text_len, void* textOption) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = ((QFontMetricsF*)self)->tightBoundingRect(text_qstr, *(QTextOption*)textOption);
	return new QRectF(__r);
}
char* QFontMetricsF_elidedText(void* self, const char* text, long long text_len, int mode, double width, long long* __out_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = ((QFontMetricsF*)self)->elidedText(text_qstr, (Qt::TextElideMode)(mode), width);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
double QFontMetricsF_underlinePos(void* self) {
	auto&& __r = ((QFontMetricsF*)self)->underlinePos();
	return __r;
}
double QFontMetricsF_overlinePos(void* self) {
	auto&& __r = ((QFontMetricsF*)self)->overlinePos();
	return __r;
}
double QFontMetricsF_strikeOutPos(void* self) {
	auto&& __r = ((QFontMetricsF*)self)->strikeOutPos();
	return __r;
}
double QFontMetricsF_lineWidth(void* self) {
	auto&& __r = ((QFontMetricsF*)self)->lineWidth();
	return __r;
}
double QFontMetricsF_fontDpi(void* self) {
	auto&& __r = ((QFontMetricsF*)self)->fontDpi();
	return __r;
}
int QFontMetricsF_OperatorEqual(void* self, void* other) {
	auto&& __r = ((QFontMetricsF*)self)->operator==(*(QFontMetricsF*)other);
	return __r;
}
int QFontMetricsF_OperatorNotEqual(void* self, void* other) {
	auto&& __r = ((QFontMetricsF*)self)->operator!=(*(QFontMetricsF*)other);
	return __r;
}
double QFontMetricsF_horizontalAdvance4(void* self, const char* string_, long long string__len, int length) {
	QString string__qstr = QString::fromUtf8(string_, string__len);
	auto&& __r = ((QFontMetricsF*)self)->horizontalAdvance(string__qstr, length);
	return __r;
}
void* QFontMetricsF_boundingRect5(void* self, void* r, int flags, const char* string_, long long string__len, int tabstops) {
	QString string__qstr = QString::fromUtf8(string_, string__len);
	auto&& __r = ((QFontMetricsF*)self)->boundingRect(*(QRectF*)r, flags, string__qstr, tabstops);
	return new QRectF(__r);
}
void* QFontMetricsF_boundingRect6(void* self, void* r, int flags, const char* string_, long long string__len, int tabstops, int* tabarray) {
	QString string__qstr = QString::fromUtf8(string_, string__len);
	int tabarray_local{};
	auto&& __r = ((QFontMetricsF*)self)->boundingRect(*(QRectF*)r, flags, string__qstr, tabstops, &tabarray_local);
	*tabarray = tabarray_local;
	return new QRectF(__r);
}
void* QFontMetricsF_size2(void* self, int flags, const char* str, long long str_len, int tabstops) {
	QString str_qstr = QString::fromUtf8(str, str_len);
	auto&& __r = ((QFontMetricsF*)self)->size(flags, str_qstr, tabstops);
	return new QSizeF(__r);
}
void* QFontMetricsF_size3(void* self, int flags, const char* str, long long str_len, int tabstops, int* tabarray) {
	QString str_qstr = QString::fromUtf8(str, str_len);
	int tabarray_local{};
	auto&& __r = ((QFontMetricsF*)self)->size(flags, str_qstr, tabstops, &tabarray_local);
	*tabarray = tabarray_local;
	return new QSizeF(__r);
}
char* QFontMetricsF_elidedText2(void* self, const char* text, long long text_len, int mode, double width, int flags, long long* __out_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = ((QFontMetricsF*)self)->elidedText(text_qstr, (Qt::TextElideMode)(mode), width, flags);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QFontMetricsF_delete(void* self) { delete (QFontMetricsF*)self; }

}

#include <QtGui/qpainter.h>
#include <QtCore/qpoint.h>
#include <QtCore/qrect.h>
#include "gen_qpainter__pixmapfragment.h"

extern "C" {

void* QPainter__PixmapFragment_new() {
	return new QPainter::PixmapFragment();
}
void* QPainter__PixmapFragment_new2(void* arg1) {
	return new QPainter::PixmapFragment(*(QPainter::PixmapFragment*)arg1);
}
void* QPainter__PixmapFragment_create(void* pos, void* sourceRect) {
	auto&& __r = QPainter::PixmapFragment::create(*(QPointF*)pos, *(QRectF*)sourceRect);
	return new QPainter::PixmapFragment(__r);
}
void* QPainter__PixmapFragment_create2(void* pos, void* sourceRect, double scaleX) {
	auto&& __r = QPainter::PixmapFragment::create(*(QPointF*)pos, *(QRectF*)sourceRect, scaleX);
	return new QPainter::PixmapFragment(__r);
}
void* QPainter__PixmapFragment_create3(void* pos, void* sourceRect, double scaleX, double scaleY) {
	auto&& __r = QPainter::PixmapFragment::create(*(QPointF*)pos, *(QRectF*)sourceRect, scaleX, scaleY);
	return new QPainter::PixmapFragment(__r);
}
void* QPainter__PixmapFragment_create4(void* pos, void* sourceRect, double scaleX, double scaleY, double rotation) {
	auto&& __r = QPainter::PixmapFragment::create(*(QPointF*)pos, *(QRectF*)sourceRect, scaleX, scaleY, rotation);
	return new QPainter::PixmapFragment(__r);
}
void* QPainter__PixmapFragment_create5(void* pos, void* sourceRect, double scaleX, double scaleY, double rotation, double opacity) {
	auto&& __r = QPainter::PixmapFragment::create(*(QPointF*)pos, *(QRectF*)sourceRect, scaleX, scaleY, rotation, opacity);
	return new QPainter::PixmapFragment(__r);
}
double QPainter__PixmapFragment_x_get(void* self) {
	auto __r = ((QPainter::PixmapFragment*)self)->x;
	return __r;
}
void QPainter__PixmapFragment_x_set(void* self, double x) {
	((QPainter::PixmapFragment*)self)->x = x;
}
double QPainter__PixmapFragment_y_get(void* self) {
	auto __r = ((QPainter::PixmapFragment*)self)->y;
	return __r;
}
void QPainter__PixmapFragment_y_set(void* self, double y) {
	((QPainter::PixmapFragment*)self)->y = y;
}
double QPainter__PixmapFragment_sourceLeft_get(void* self) {
	auto __r = ((QPainter::PixmapFragment*)self)->sourceLeft;
	return __r;
}
void QPainter__PixmapFragment_sourceLeft_set(void* self, double sourceLeft) {
	((QPainter::PixmapFragment*)self)->sourceLeft = sourceLeft;
}
double QPainter__PixmapFragment_sourceTop_get(void* self) {
	auto __r = ((QPainter::PixmapFragment*)self)->sourceTop;
	return __r;
}
void QPainter__PixmapFragment_sourceTop_set(void* self, double sourceTop) {
	((QPainter::PixmapFragment*)self)->sourceTop = sourceTop;
}
double QPainter__PixmapFragment_width_get(void* self) {
	auto __r = ((QPainter::PixmapFragment*)self)->width;
	return __r;
}
void QPainter__PixmapFragment_width_set(void* self, double width) {
	((QPainter::PixmapFragment*)self)->width = width;
}
double QPainter__PixmapFragment_height_get(void* self) {
	auto __r = ((QPainter::PixmapFragment*)self)->height;
	return __r;
}
void QPainter__PixmapFragment_height_set(void* self, double height) {
	((QPainter::PixmapFragment*)self)->height = height;
}
double QPainter__PixmapFragment_scaleX_get(void* self) {
	auto __r = ((QPainter::PixmapFragment*)self)->scaleX;
	return __r;
}
void QPainter__PixmapFragment_scaleX_set(void* self, double scaleX) {
	((QPainter::PixmapFragment*)self)->scaleX = scaleX;
}
double QPainter__PixmapFragment_scaleY_get(void* self) {
	auto __r = ((QPainter::PixmapFragment*)self)->scaleY;
	return __r;
}
void QPainter__PixmapFragment_scaleY_set(void* self, double scaleY) {
	((QPainter::PixmapFragment*)self)->scaleY = scaleY;
}
double QPainter__PixmapFragment_rotation_get(void* self) {
	auto __r = ((QPainter::PixmapFragment*)self)->rotation;
	return __r;
}
void QPainter__PixmapFragment_rotation_set(void* self, double rotation) {
	((QPainter::PixmapFragment*)self)->rotation = rotation;
}
double QPainter__PixmapFragment_opacity_get(void* self) {
	auto __r = ((QPainter::PixmapFragment*)self)->opacity;
	return __r;
}
void QPainter__PixmapFragment_opacity_set(void* self, double opacity) {
	((QPainter::PixmapFragment*)self)->opacity = opacity;
}
void QPainter__PixmapFragment_delete(void* self) { delete (QPainter::PixmapFragment*)self; }

}

#include <QtGui/qtextobject.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtGui/qtextformat.h>
#include <QtGui/qtextdocument.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qobject.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qvariant.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qtextblockgroup.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

static std::unordered_map<long long, QMetaObject::Connection> __conns_QTextBlockGroup;
static std::mutex __connsMu_QTextBlockGroup;

extern "C" {

void* QTextBlockGroup_base(void* self) { return (void*)static_cast<QTextObject*>((QTextBlockGroup*)self); }
void QTextBlockGroup_chain(void* self, void** out) {
	QTextBlockGroup* __p = (QTextBlockGroup*)self;
	QTextObject* __b0 = static_cast<QTextObject*>(__p); out[0] = (void*)__b0;
	QObject* __b1 = static_cast<QObject*>(__b0); out[1] = (void*)__b1;
}
void* QTextBlockGroup_fromQObject(void* o) { return (void*)qobject_cast<QTextBlockGroup*>((QObject*)o); }
char* QTextBlockGroup_tr(const char* s, long long* __out_len) {
	auto&& __r = QTextBlockGroup::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QTextBlockGroup_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QTextBlockGroup::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QTextBlockGroup_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QTextBlockGroup::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
char* QTextBlockGroup_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QTextBlockGroup::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QTextBlockGroup_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QTextBlockGroup::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QTextBlockGroup_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QTextBlockGroup::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QTextBlockGroup_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QTextBlockGroup::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void QTextBlockGroup_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTextBlockGroup*)self, &QTextBlockGroup::destroyed, (QTextBlockGroup*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTextBlockGroup);
	__conns_QTextBlockGroup[__id] = __c;
}
void QTextBlockGroup_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTextBlockGroup*)self, &QTextBlockGroup::destroyed, (QTextBlockGroup*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTextBlockGroup);
	__conns_QTextBlockGroup[__id] = __c;
}
void QTextBlockGroup_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTextBlockGroup*)self, &QTextBlockGroup::objectNameChanged, (QTextBlockGroup*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTextBlockGroup);
	__conns_QTextBlockGroup[__id] = __c;
}
void QTextBlockGroup__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QTextBlockGroup);
		auto __it = __conns_QTextBlockGroup.find(id);
		if (__it == __conns_QTextBlockGroup.end()) return;
		__c = __it->second;
		__conns_QTextBlockGroup.erase(__it);
	}
	QObject::disconnect(__c);
}

}

#include <QtWidgets/qabstractspinbox.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qsize.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qvariant.h>
#include <QtGui/qvalidator.h>
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
#include "gen_qabstractspinbox.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QAbstractSpinBox__Go : public QAbstractSpinBox {
public:
	long long __goId;
	unsigned char __goOn[54] = {};
	template<typename... A> QAbstractSpinBox__Go(long long id, A&&... a) : QAbstractSpinBox(std::forward<A>(a)...), __goId(id) {}
	using QAbstractSpinBox::lineEdit;
	using QAbstractSpinBox::setLineEdit;
	using QAbstractSpinBox::updateMicroFocus;
	using QAbstractSpinBox::create;
	using QAbstractSpinBox::destroy;
	using QAbstractSpinBox::focusNextChild;
	using QAbstractSpinBox::focusPreviousChild;
	using QAbstractSpinBox::sender;
	using QAbstractSpinBox::senderSignalIndex;
	using QAbstractSpinBox::receivers;
	using QAbstractSpinBox::isSignalConnected;
	using QAbstractSpinBox::getDecodedMetricF;
	void cppgoBase_resizeEvent(QResizeEvent* __a0) { QAbstractSpinBox::resizeEvent(__a0); }
	void cppgoBase_keyPressEvent(QKeyEvent* __a0) { QAbstractSpinBox::keyPressEvent(__a0); }
	void cppgoBase_keyReleaseEvent(QKeyEvent* __a0) { QAbstractSpinBox::keyReleaseEvent(__a0); }
	void cppgoBase_wheelEvent(QWheelEvent* __a0) { QAbstractSpinBox::wheelEvent(__a0); }
	void cppgoBase_focusInEvent(QFocusEvent* __a0) { QAbstractSpinBox::focusInEvent(__a0); }
	void cppgoBase_focusOutEvent(QFocusEvent* __a0) { QAbstractSpinBox::focusOutEvent(__a0); }
	void cppgoBase_contextMenuEvent(QContextMenuEvent* __a0) { QAbstractSpinBox::contextMenuEvent(__a0); }
	void cppgoBase_changeEvent(QEvent* __a0) { QAbstractSpinBox::changeEvent(__a0); }
	void cppgoBase_closeEvent(QCloseEvent* __a0) { QAbstractSpinBox::closeEvent(__a0); }
	void cppgoBase_hideEvent(QHideEvent* __a0) { QAbstractSpinBox::hideEvent(__a0); }
	void cppgoBase_mousePressEvent(QMouseEvent* __a0) { QAbstractSpinBox::mousePressEvent(__a0); }
	void cppgoBase_mouseReleaseEvent(QMouseEvent* __a0) { QAbstractSpinBox::mouseReleaseEvent(__a0); }
	void cppgoBase_mouseMoveEvent(QMouseEvent* __a0) { QAbstractSpinBox::mouseMoveEvent(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QAbstractSpinBox::timerEvent(__a0); }
	void cppgoBase_paintEvent(QPaintEvent* __a0) { QAbstractSpinBox::paintEvent(__a0); }
	void cppgoBase_showEvent(QShowEvent* __a0) { QAbstractSpinBox::showEvent(__a0); }
	void cppgoBase_initStyleOption(QStyleOptionSpinBox* __a0) { QAbstractSpinBox::initStyleOption(__a0); }
	QFlags<QAbstractSpinBox::StepEnabledFlag> cppgoBase_stepEnabled() { return QAbstractSpinBox::stepEnabled(); }
	void cppgoBase_mouseDoubleClickEvent(QMouseEvent* __a0) { QAbstractSpinBox::mouseDoubleClickEvent(__a0); }
	void cppgoBase_enterEvent(QEnterEvent* __a0) { QAbstractSpinBox::enterEvent(__a0); }
	void cppgoBase_leaveEvent(QEvent* __a0) { QAbstractSpinBox::leaveEvent(__a0); }
	void cppgoBase_moveEvent(QMoveEvent* __a0) { QAbstractSpinBox::moveEvent(__a0); }
	void cppgoBase_tabletEvent(QTabletEvent* __a0) { QAbstractSpinBox::tabletEvent(__a0); }
	void cppgoBase_actionEvent(QActionEvent* __a0) { QAbstractSpinBox::actionEvent(__a0); }
	void cppgoBase_dragEnterEvent(QDragEnterEvent* __a0) { QAbstractSpinBox::dragEnterEvent(__a0); }
	void cppgoBase_dragMoveEvent(QDragMoveEvent* __a0) { QAbstractSpinBox::dragMoveEvent(__a0); }
	void cppgoBase_dragLeaveEvent(QDragLeaveEvent* __a0) { QAbstractSpinBox::dragLeaveEvent(__a0); }
	void cppgoBase_dropEvent(QDropEvent* __a0) { QAbstractSpinBox::dropEvent(__a0); }
	int cppgoBase_metric(int __a0) { return QAbstractSpinBox::metric((QPaintDevice::PaintDeviceMetric)__a0); }
	void cppgoBase_initPainter(QPainter* __a0) { QAbstractSpinBox::initPainter(__a0); }
	QPaintDevice* cppgoBase_redirected(QPoint* __a0) { return QAbstractSpinBox::redirected(__a0); }
	QPainter* cppgoBase_sharedPainter() { return QAbstractSpinBox::sharedPainter(); }
	void cppgoBase_inputMethodEvent(QInputMethodEvent* __a0) { QAbstractSpinBox::inputMethodEvent(__a0); }
	bool cppgoBase_focusNextPrevChild(bool __a0) { return QAbstractSpinBox::focusNextPrevChild(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QAbstractSpinBox::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QAbstractSpinBox::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QAbstractSpinBox::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QAbstractSpinBox::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QAbstractSpinBox::metaObject();
	}
	QSize sizeHint() const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 729690261457684101LL, &__h); if (__h) return *(QSize*)__r; }
		return QAbstractSpinBox::sizeHint();
	}
	QSize minimumSizeHint() const override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 7049264599837532491LL, &__h); if (__h) return *(QSize*)__r; }
		return QAbstractSpinBox::minimumSizeHint();
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QAbstractSpinBox::event(event);
	}
	QVariant inputMethodQuery(Qt::InputMethodQuery arg1) const override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7345853097340709638LL, (int)(arg1), &__h); if (__h) return *(QVariant*)__r; }
		return QAbstractSpinBox::inputMethodQuery(arg1);
	}
	QValidator::State validate(QString& input, int& pos) const override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_int_voidp_intp) { int __r = ((int(*)(long long, long long, void*, int*, int*))cppgo_virt_int_voidp_intp)(__goId, 6784179835071232555LL, (void*)&input, (int*)&pos, &__h); if (__h) return (QValidator::State)(__r); }
		return QAbstractSpinBox::validate(input, pos);
	}
	void fixup(QString& input) const override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4855206404746966715LL, (void*)&input, &__h); if (__h) return; }
		QAbstractSpinBox::fixup(input);
	}
	void stepBy(int steps) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 7085718643589954634LL, steps, &__h); if (__h) return; }
		QAbstractSpinBox::stepBy(steps);
	}
	void clear() override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, -778567940572639070LL, &__h); if (__h) return; }
		QAbstractSpinBox::clear();
	}
	void resizeEvent(QResizeEvent* event) override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4302917285655014661LL, (void*)event, &__h); if (__h) return; }
		QAbstractSpinBox::resizeEvent(event);
	}
	void keyPressEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1023791977679793185LL, (void*)event, &__h); if (__h) return; }
		QAbstractSpinBox::keyPressEvent(event);
	}
	void keyReleaseEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 772598475950050299LL, (void*)event, &__h); if (__h) return; }
		QAbstractSpinBox::keyReleaseEvent(event);
	}
	void wheelEvent(QWheelEvent* event) override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1919988340934511316LL, (void*)event, &__h); if (__h) return; }
		QAbstractSpinBox::wheelEvent(event);
	}
	void focusInEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 975043767033552552LL, (void*)event, &__h); if (__h) return; }
		QAbstractSpinBox::focusInEvent(event);
	}
	void focusOutEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -385596010576048993LL, (void*)event, &__h); if (__h) return; }
		QAbstractSpinBox::focusOutEvent(event);
	}
	void contextMenuEvent(QContextMenuEvent* event) override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6698043222227006015LL, (void*)event, &__h); if (__h) return; }
		QAbstractSpinBox::contextMenuEvent(event);
	}
	void changeEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4959376063826165943LL, (void*)event, &__h); if (__h) return; }
		QAbstractSpinBox::changeEvent(event);
	}
	void closeEvent(QCloseEvent* event) override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1989937050019464783LL, (void*)event, &__h); if (__h) return; }
		QAbstractSpinBox::closeEvent(event);
	}
	void hideEvent(QHideEvent* event) override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -199718707876719129LL, (void*)event, &__h); if (__h) return; }
		QAbstractSpinBox::hideEvent(event);
	}
	void mousePressEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -332090898078543319LL, (void*)event, &__h); if (__h) return; }
		QAbstractSpinBox::mousePressEvent(event);
	}
	void mouseReleaseEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4415920049297795213LL, (void*)event, &__h); if (__h) return; }
		QAbstractSpinBox::mouseReleaseEvent(event);
	}
	void mouseMoveEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4255494004099806581LL, (void*)event, &__h); if (__h) return; }
		QAbstractSpinBox::mouseMoveEvent(event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QAbstractSpinBox::timerEvent(event);
	}
	void paintEvent(QPaintEvent* event) override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5290442163167458465LL, (void*)event, &__h); if (__h) return; }
		QAbstractSpinBox::paintEvent(event);
	}
	void showEvent(QShowEvent* event) override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6825369174751916LL, (void*)event, &__h); if (__h) return; }
		QAbstractSpinBox::showEvent(event);
	}
	void initStyleOption(QStyleOptionSpinBox* option) const override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -488693769771308531LL, (void*)option, &__h); if (__h) return; }
		QAbstractSpinBox::initStyleOption(option);
	}
	QFlags<QAbstractSpinBox::StepEnabledFlag> stepEnabled() const override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, 5257289574516412244LL, &__h); if (__h) return QFlags<QAbstractSpinBox::StepEnabledFlag>(__r); }
		return QAbstractSpinBox::stepEnabled();
	}
	int devType() const override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8659679893868524044LL, &__h); if (__h) return __r; }
		return QAbstractSpinBox::devType();
	}
	void setVisible(bool visible) override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 2110182719375245575LL, visible, &__h); if (__h) return; }
		QAbstractSpinBox::setVisible(visible);
	}
	int heightForWidth(int arg1) const override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 4311534060495737339LL, arg1, &__h); if (__h) return __r; }
		return QAbstractSpinBox::heightForWidth(arg1);
	}
	bool hasHeightForWidth() const override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8266762397995423915LL, &__h); if (__h) return __r; }
		return QAbstractSpinBox::hasHeightForWidth();
	}
	QPaintEngine* paintEngine() const override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 6398007404756346381LL, &__h); if (__h) return (QPaintEngine*)__r; }
		return QAbstractSpinBox::paintEngine();
	}
	void mouseDoubleClickEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[32] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4029073905697023887LL, (void*)event, &__h); if (__h) return; }
		QAbstractSpinBox::mouseDoubleClickEvent(event);
	}
	void enterEvent(QEnterEvent* event) override {
		int __h = 0;
		if (__goOn[33] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6316270346668237895LL, (void*)event, &__h); if (__h) return; }
		QAbstractSpinBox::enterEvent(event);
	}
	void leaveEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[34] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -551070009623711640LL, (void*)event, &__h); if (__h) return; }
		QAbstractSpinBox::leaveEvent(event);
	}
	void moveEvent(QMoveEvent* event) override {
		int __h = 0;
		if (__goOn[35] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -9161949150489525396LL, (void*)event, &__h); if (__h) return; }
		QAbstractSpinBox::moveEvent(event);
	}
	void tabletEvent(QTabletEvent* event) override {
		int __h = 0;
		if (__goOn[36] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -212068866860701141LL, (void*)event, &__h); if (__h) return; }
		QAbstractSpinBox::tabletEvent(event);
	}
	void actionEvent(QActionEvent* event) override {
		int __h = 0;
		if (__goOn[37] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1697444140088555581LL, (void*)event, &__h); if (__h) return; }
		QAbstractSpinBox::actionEvent(event);
	}
	void dragEnterEvent(QDragEnterEvent* event) override {
		int __h = 0;
		if (__goOn[38] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 206122715919917589LL, (void*)event, &__h); if (__h) return; }
		QAbstractSpinBox::dragEnterEvent(event);
	}
	void dragMoveEvent(QDragMoveEvent* event) override {
		int __h = 0;
		if (__goOn[39] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1781545137352745978LL, (void*)event, &__h); if (__h) return; }
		QAbstractSpinBox::dragMoveEvent(event);
	}
	void dragLeaveEvent(QDragLeaveEvent* event) override {
		int __h = 0;
		if (__goOn[40] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -3139904158373725006LL, (void*)event, &__h); if (__h) return; }
		QAbstractSpinBox::dragLeaveEvent(event);
	}
	void dropEvent(QDropEvent* event) override {
		int __h = 0;
		if (__goOn[41] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 9107058682754591716LL, (void*)event, &__h); if (__h) return; }
		QAbstractSpinBox::dropEvent(event);
	}
	bool nativeEvent(const QByteArray& eventType, void* message, qintptr* result) override {
		int __h = 0;
		if (__goOn[42] && cppgo_virt_int_voidp_voidp_longlongp) { int __r = ((int(*)(long long, long long, void*, void*, long long*, int*))cppgo_virt_int_voidp_voidp_longlongp)(__goId, -2039166586246443706LL, (void*)&eventType, (void*)message, (long long*)result, &__h); if (__h) return __r; }
		return QAbstractSpinBox::nativeEvent(eventType, message, result);
	}
	int metric(QPaintDevice::PaintDeviceMetric arg1) const override {
		int __h = 0;
		if (__goOn[43] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 427629649267311777LL, (int)(arg1), &__h); if (__h) return __r; }
		return QAbstractSpinBox::metric(arg1);
	}
	void initPainter(QPainter* painter) const override {
		int __h = 0;
		if (__goOn[44] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6477500296980050210LL, (void*)painter, &__h); if (__h) return; }
		QAbstractSpinBox::initPainter(painter);
	}
	QPaintDevice* redirected(QPoint* offset) const override {
		int __h = 0;
		if (__goOn[45] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 4127240633310655048LL, (void*)offset, &__h); if (__h) return (QPaintDevice*)__r; }
		return QAbstractSpinBox::redirected(offset);
	}
	QPainter* sharedPainter() const override {
		int __h = 0;
		if (__goOn[46] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 5792771328741774991LL, &__h); if (__h) return (QPainter*)__r; }
		return QAbstractSpinBox::sharedPainter();
	}
	void inputMethodEvent(QInputMethodEvent* arg1) override {
		int __h = 0;
		if (__goOn[47] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5438819661942843222LL, (void*)arg1, &__h); if (__h) return; }
		QAbstractSpinBox::inputMethodEvent(arg1);
	}
	bool focusNextPrevChild(bool next) override {
		int __h = 0;
		if (__goOn[48] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 5898328517505784191LL, next, &__h); if (__h) return __r; }
		return QAbstractSpinBox::focusNextPrevChild(next);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[49] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QAbstractSpinBox::eventFilter(watched, event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[50] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QAbstractSpinBox::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[51] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QAbstractSpinBox::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[52] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QAbstractSpinBox::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[53] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QAbstractSpinBox::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QAbstractSpinBox;
static std::mutex __connsMu_QAbstractSpinBox;

extern "C" {

void* QAbstractSpinBox_base(void* self) { return (void*)static_cast<QWidget*>((QAbstractSpinBox*)self); }
void QAbstractSpinBox_chain(void* self, void** out) {
	QAbstractSpinBox* __p = (QAbstractSpinBox*)self;
	QWidget* __b0 = static_cast<QWidget*>(__p); out[0] = (void*)__b0;
	QObject* __b1 = static_cast<QObject*>(__b0); out[1] = (void*)__b1;
}
void QAbstractSpinBox__set_override(void* self, int idx) { ((QAbstractSpinBox__Go*)self)->__goOn[idx] = 1; }
void* QAbstractSpinBox_fromQObject(void* o) { return (void*)qobject_cast<QAbstractSpinBox*>((QObject*)o); }
void* QAbstractSpinBox_new(long long __goid, void* parent) {
	return new QAbstractSpinBox__Go(__goid, (QWidget*)parent);
}
void* QAbstractSpinBox_new2(long long __goid) {
	return new QAbstractSpinBox__Go(__goid);
}
char* QAbstractSpinBox_tr(const char* s, long long* __out_len) {
	auto&& __r = QAbstractSpinBox::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QAbstractSpinBox_buttonSymbols(void* self) {
	auto&& __r = ((QAbstractSpinBox*)self)->buttonSymbols();
	return (int)(__r);
}
void QAbstractSpinBox_setButtonSymbols(void* self, int bs) {
	((QAbstractSpinBox*)self)->setButtonSymbols((QAbstractSpinBox::ButtonSymbols)(bs));
}
void QAbstractSpinBox_setCorrectionMode(void* self, int cm) {
	((QAbstractSpinBox*)self)->setCorrectionMode((QAbstractSpinBox::CorrectionMode)(cm));
}
int QAbstractSpinBox_correctionMode(void* self) {
	auto&& __r = ((QAbstractSpinBox*)self)->correctionMode();
	return (int)(__r);
}
int QAbstractSpinBox_hasAcceptableInput(void* self) {
	auto&& __r = ((QAbstractSpinBox*)self)->hasAcceptableInput();
	return __r;
}
char* QAbstractSpinBox_text(void* self, long long* __out_len) {
	auto&& __r = ((QAbstractSpinBox*)self)->text();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QAbstractSpinBox_specialValueText(void* self, long long* __out_len) {
	auto&& __r = ((QAbstractSpinBox*)self)->specialValueText();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QAbstractSpinBox_setSpecialValueText(void* self, const char* txt, long long txt_len) {
	QString txt_qstr = QString::fromUtf8(txt, txt_len);
	((QAbstractSpinBox*)self)->setSpecialValueText(txt_qstr);
}
int QAbstractSpinBox_wrapping(void* self) {
	auto&& __r = ((QAbstractSpinBox*)self)->wrapping();
	return __r;
}
void QAbstractSpinBox_setWrapping(void* self, int w) {
	((QAbstractSpinBox*)self)->setWrapping(w);
}
void QAbstractSpinBox_setReadOnly(void* self, int r) {
	((QAbstractSpinBox*)self)->setReadOnly(r);
}
int QAbstractSpinBox_isReadOnly(void* self) {
	auto&& __r = ((QAbstractSpinBox*)self)->isReadOnly();
	return __r;
}
void QAbstractSpinBox_setKeyboardTracking(void* self, int kt) {
	((QAbstractSpinBox*)self)->setKeyboardTracking(kt);
}
int QAbstractSpinBox_keyboardTracking(void* self) {
	auto&& __r = ((QAbstractSpinBox*)self)->keyboardTracking();
	return __r;
}
void QAbstractSpinBox_setAlignment(void* self, int flag) {
	((QAbstractSpinBox*)self)->setAlignment(QFlags<Qt::AlignmentFlag>(flag));
}
int QAbstractSpinBox_alignment(void* self) {
	auto&& __r = ((QAbstractSpinBox*)self)->alignment();
	return (int)(__r);
}
void QAbstractSpinBox_setFrame(void* self, int arg1) {
	((QAbstractSpinBox*)self)->setFrame(arg1);
}
int QAbstractSpinBox_hasFrame(void* self) {
	auto&& __r = ((QAbstractSpinBox*)self)->hasFrame();
	return __r;
}
void QAbstractSpinBox_setAccelerated(void* self, int on) {
	((QAbstractSpinBox*)self)->setAccelerated(on);
}
int QAbstractSpinBox_isAccelerated(void* self) {
	auto&& __r = ((QAbstractSpinBox*)self)->isAccelerated();
	return __r;
}
void QAbstractSpinBox_setGroupSeparatorShown(void* self, int shown) {
	((QAbstractSpinBox*)self)->setGroupSeparatorShown(shown);
}
int QAbstractSpinBox_isGroupSeparatorShown(void* self) {
	auto&& __r = ((QAbstractSpinBox*)self)->isGroupSeparatorShown();
	return __r;
}
void QAbstractSpinBox_interpretText(void* self) {
	((QAbstractSpinBox*)self)->interpretText();
}
int QAbstractSpinBox_validate(void* self, void* input, int* pos) {
	int pos_ref = *pos;
	auto&& __r = ((QAbstractSpinBox*)self)->validate(*(QString*)input, pos_ref);
	*pos = pos_ref;
	return (int)(__r);
}
void QAbstractSpinBox_fixup(void* self, void* input) {
	((QAbstractSpinBox*)self)->fixup(*(QString*)input);
}
void QAbstractSpinBox_stepBy(void* self, int steps) {
	((QAbstractSpinBox*)self)->stepBy(steps);
}
void QAbstractSpinBox_stepUp(void* self) {
	((QAbstractSpinBox*)self)->stepUp();
}
void QAbstractSpinBox_stepDown(void* self) {
	((QAbstractSpinBox*)self)->stepDown();
}
void QAbstractSpinBox_selectAll(void* self) {
	((QAbstractSpinBox*)self)->selectAll();
}
void QAbstractSpinBox_clear(void* self) {
	((QAbstractSpinBox*)self)->clear();
}
void* QAbstractSpinBox_lineEdit(void* self) {
	auto&& __r = ((QAbstractSpinBox__Go*)self)->lineEdit();
	return (void*)(__r);
}
void QAbstractSpinBox_setLineEdit(void* self, void* edit) {
	((QAbstractSpinBox__Go*)self)->setLineEdit((QLineEdit*)edit);
}
void QAbstractSpinBox_editingFinished(void* self) {
	((QAbstractSpinBox*)self)->editingFinished();
}
void QAbstractSpinBox_returnPressed(void* self) {
	((QAbstractSpinBox*)self)->returnPressed();
}
void QAbstractSpinBox_setTabOrder(void* arg1, void* arg2) {
	QAbstractSpinBox::setTabOrder((QWidget*)arg1, (QWidget*)arg2);
}
void* QAbstractSpinBox_mouseGrabber() {
	auto&& __r = QAbstractSpinBox::mouseGrabber();
	return (void*)(__r);
}
void* QAbstractSpinBox_keyboardGrabber() {
	auto&& __r = QAbstractSpinBox::keyboardGrabber();
	return (void*)(__r);
}
void* QAbstractSpinBox_find(unsigned long long arg1) {
	auto&& __r = QAbstractSpinBox::find(arg1);
	return (void*)(__r);
}
void* QAbstractSpinBox_createWindowContainer(void* window) {
	auto&& __r = QAbstractSpinBox::createWindowContainer((QWindow*)window);
	return (void*)(__r);
}
void QAbstractSpinBox_updateMicroFocus(void* self) {
	((QAbstractSpinBox__Go*)self)->updateMicroFocus();
}
void QAbstractSpinBox_create(void* self) {
	((QAbstractSpinBox__Go*)self)->create();
}
void QAbstractSpinBox_destroy(void* self) {
	((QAbstractSpinBox__Go*)self)->destroy();
}
int QAbstractSpinBox_focusNextChild(void* self) {
	auto&& __r = ((QAbstractSpinBox__Go*)self)->focusNextChild();
	return __r;
}
int QAbstractSpinBox_focusPreviousChild(void* self) {
	auto&& __r = ((QAbstractSpinBox__Go*)self)->focusPreviousChild();
	return __r;
}
void* QAbstractSpinBox_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QAbstractSpinBox::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QAbstractSpinBox_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QAbstractSpinBox::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QAbstractSpinBox_sender(void* self) {
	auto&& __r = ((QAbstractSpinBox__Go*)self)->sender();
	return (void*)(__r);
}
int QAbstractSpinBox_senderSignalIndex(void* self) {
	auto&& __r = ((QAbstractSpinBox__Go*)self)->senderSignalIndex();
	return __r;
}
int QAbstractSpinBox_receivers(void* self, const char* signal) {
	auto&& __r = ((QAbstractSpinBox__Go*)self)->receivers(signal);
	return __r;
}
int QAbstractSpinBox_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QAbstractSpinBox__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
double QAbstractSpinBox_devicePixelRatioFScale() {
	auto&& __r = QAbstractSpinBox::devicePixelRatioFScale();
	return __r;
}
int QAbstractSpinBox_encodeMetricF(int metric, double value) {
	auto&& __r = QAbstractSpinBox::encodeMetricF((QPaintDevice::PaintDeviceMetric)(metric), value);
	return __r;
}
double QAbstractSpinBox_getDecodedMetricF(void* self, int metricA, int metricB) {
	auto&& __r = ((QAbstractSpinBox__Go*)self)->getDecodedMetricF((QPaintDevice::PaintDeviceMetric)(metricA), (QPaintDevice::PaintDeviceMetric)(metricB));
	return __r;
}
char* QAbstractSpinBox_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QAbstractSpinBox::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QAbstractSpinBox_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QAbstractSpinBox::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QAbstractSpinBox_createWindowContainer2(void* window, void* parent) {
	auto&& __r = QAbstractSpinBox::createWindowContainer((QWindow*)window, (QWidget*)parent);
	return (void*)(__r);
}
void* QAbstractSpinBox_createWindowContainer3(void* window, void* parent, int flags) {
	auto&& __r = QAbstractSpinBox::createWindowContainer((QWindow*)window, (QWidget*)parent, QFlags<Qt::WindowType>(flags));
	return (void*)(__r);
}
void QAbstractSpinBox_updateMicroFocus2(void* self, int query) {
	((QAbstractSpinBox__Go*)self)->updateMicroFocus((Qt::InputMethodQuery)(query));
}
void QAbstractSpinBox_create2(void* self, unsigned long long arg1) {
	((QAbstractSpinBox__Go*)self)->create(arg1);
}
void QAbstractSpinBox_create3(void* self, unsigned long long arg1, int initializeWindow) {
	((QAbstractSpinBox__Go*)self)->create(arg1, initializeWindow);
}
void QAbstractSpinBox_create4(void* self, unsigned long long arg1, int initializeWindow, int destroyOldWindow) {
	((QAbstractSpinBox__Go*)self)->create(arg1, initializeWindow, destroyOldWindow);
}
void QAbstractSpinBox_destroy2(void* self, int destroyWindow) {
	((QAbstractSpinBox__Go*)self)->destroy(destroyWindow);
}
void QAbstractSpinBox_destroy3(void* self, int destroyWindow, int destroySubWindows) {
	((QAbstractSpinBox__Go*)self)->destroy(destroyWindow, destroySubWindows);
}
void* QAbstractSpinBox_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QAbstractSpinBox::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QAbstractSpinBox_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QAbstractSpinBox::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
int QAbstractSpinBox_validate_default(void* self, void* input, int* pos) {
	int pos_ref = *pos;
	auto&& __r = ((QAbstractSpinBox*)self)->QAbstractSpinBox::validate(*(QString*)input, pos_ref);
	*pos = pos_ref;
	return (int)(__r);
}
void QAbstractSpinBox_fixup_default(void* self, void* input) {
	((QAbstractSpinBox*)self)->QAbstractSpinBox::fixup(*(QString*)input);
}
void QAbstractSpinBox_stepBy_default(void* self, int steps) {
	((QAbstractSpinBox*)self)->QAbstractSpinBox::stepBy(steps);
}
void QAbstractSpinBox_clear_default(void* self) {
	((QAbstractSpinBox*)self)->QAbstractSpinBox::clear();
}
void* QAbstractSpinBox_metaObject_default(void* self) {
	auto&& __r = ((QAbstractSpinBox*)self)->QAbstractSpinBox::metaObject();
	return (void*)(__r);
}
void* QAbstractSpinBox_sizeHint_default(void* self) {
	auto&& __r = ((QAbstractSpinBox*)self)->QAbstractSpinBox::sizeHint();
	return new QSize(__r);
}
void* QAbstractSpinBox_minimumSizeHint_default(void* self) {
	auto&& __r = ((QAbstractSpinBox*)self)->QAbstractSpinBox::minimumSizeHint();
	return new QSize(__r);
}
int QAbstractSpinBox_event_default(void* self, void* event) {
	auto&& __r = ((QAbstractSpinBox*)self)->QAbstractSpinBox::event((QEvent*)event);
	return __r;
}
void* QAbstractSpinBox_inputMethodQuery_default(void* self, int arg1) {
	auto&& __r = ((QAbstractSpinBox*)self)->QAbstractSpinBox::inputMethodQuery((Qt::InputMethodQuery)(arg1));
	return new QVariant(__r);
}
void QAbstractSpinBox_resizeEvent_default(void* self, void* event) {
	((QAbstractSpinBox__Go*)self)->cppgoBase_resizeEvent((QResizeEvent*)event);
}
void QAbstractSpinBox_keyPressEvent_default(void* self, void* event) {
	((QAbstractSpinBox__Go*)self)->cppgoBase_keyPressEvent((QKeyEvent*)event);
}
void QAbstractSpinBox_keyReleaseEvent_default(void* self, void* event) {
	((QAbstractSpinBox__Go*)self)->cppgoBase_keyReleaseEvent((QKeyEvent*)event);
}
void QAbstractSpinBox_wheelEvent_default(void* self, void* event) {
	((QAbstractSpinBox__Go*)self)->cppgoBase_wheelEvent((QWheelEvent*)event);
}
void QAbstractSpinBox_focusInEvent_default(void* self, void* event) {
	((QAbstractSpinBox__Go*)self)->cppgoBase_focusInEvent((QFocusEvent*)event);
}
void QAbstractSpinBox_focusOutEvent_default(void* self, void* event) {
	((QAbstractSpinBox__Go*)self)->cppgoBase_focusOutEvent((QFocusEvent*)event);
}
void QAbstractSpinBox_contextMenuEvent_default(void* self, void* event) {
	((QAbstractSpinBox__Go*)self)->cppgoBase_contextMenuEvent((QContextMenuEvent*)event);
}
void QAbstractSpinBox_changeEvent_default(void* self, void* event) {
	((QAbstractSpinBox__Go*)self)->cppgoBase_changeEvent((QEvent*)event);
}
void QAbstractSpinBox_closeEvent_default(void* self, void* event) {
	((QAbstractSpinBox__Go*)self)->cppgoBase_closeEvent((QCloseEvent*)event);
}
void QAbstractSpinBox_hideEvent_default(void* self, void* event) {
	((QAbstractSpinBox__Go*)self)->cppgoBase_hideEvent((QHideEvent*)event);
}
void QAbstractSpinBox_mousePressEvent_default(void* self, void* event) {
	((QAbstractSpinBox__Go*)self)->cppgoBase_mousePressEvent((QMouseEvent*)event);
}
void QAbstractSpinBox_mouseReleaseEvent_default(void* self, void* event) {
	((QAbstractSpinBox__Go*)self)->cppgoBase_mouseReleaseEvent((QMouseEvent*)event);
}
void QAbstractSpinBox_mouseMoveEvent_default(void* self, void* event) {
	((QAbstractSpinBox__Go*)self)->cppgoBase_mouseMoveEvent((QMouseEvent*)event);
}
void QAbstractSpinBox_timerEvent_default(void* self, void* event) {
	((QAbstractSpinBox__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QAbstractSpinBox_paintEvent_default(void* self, void* event) {
	((QAbstractSpinBox__Go*)self)->cppgoBase_paintEvent((QPaintEvent*)event);
}
void QAbstractSpinBox_showEvent_default(void* self, void* event) {
	((QAbstractSpinBox__Go*)self)->cppgoBase_showEvent((QShowEvent*)event);
}
void QAbstractSpinBox_initStyleOption_default(void* self, void* option) {
	((QAbstractSpinBox__Go*)self)->cppgoBase_initStyleOption((QStyleOptionSpinBox*)option);
}
int QAbstractSpinBox_stepEnabled_default(void* self) {
	auto&& __r = ((QAbstractSpinBox__Go*)self)->cppgoBase_stepEnabled();
	return (int)(__r);
}
int QAbstractSpinBox_devType_default(void* self) {
	auto&& __r = ((QAbstractSpinBox*)self)->QAbstractSpinBox::devType();
	return __r;
}
void QAbstractSpinBox_setVisible_default(void* self, int visible) {
	((QAbstractSpinBox*)self)->QAbstractSpinBox::setVisible(visible);
}
int QAbstractSpinBox_heightForWidth_default(void* self, int arg1) {
	auto&& __r = ((QAbstractSpinBox*)self)->QAbstractSpinBox::heightForWidth(arg1);
	return __r;
}
int QAbstractSpinBox_hasHeightForWidth_default(void* self) {
	auto&& __r = ((QAbstractSpinBox*)self)->QAbstractSpinBox::hasHeightForWidth();
	return __r;
}
void* QAbstractSpinBox_paintEngine_default(void* self) {
	auto&& __r = ((QAbstractSpinBox*)self)->QAbstractSpinBox::paintEngine();
	return (void*)(__r);
}
void QAbstractSpinBox_mouseDoubleClickEvent_default(void* self, void* event) {
	((QAbstractSpinBox__Go*)self)->cppgoBase_mouseDoubleClickEvent((QMouseEvent*)event);
}
void QAbstractSpinBox_enterEvent_default(void* self, void* event) {
	((QAbstractSpinBox__Go*)self)->cppgoBase_enterEvent((QEnterEvent*)event);
}
void QAbstractSpinBox_leaveEvent_default(void* self, void* event) {
	((QAbstractSpinBox__Go*)self)->cppgoBase_leaveEvent((QEvent*)event);
}
void QAbstractSpinBox_moveEvent_default(void* self, void* event) {
	((QAbstractSpinBox__Go*)self)->cppgoBase_moveEvent((QMoveEvent*)event);
}
void QAbstractSpinBox_tabletEvent_default(void* self, void* event) {
	((QAbstractSpinBox__Go*)self)->cppgoBase_tabletEvent((QTabletEvent*)event);
}
void QAbstractSpinBox_actionEvent_default(void* self, void* event) {
	((QAbstractSpinBox__Go*)self)->cppgoBase_actionEvent((QActionEvent*)event);
}
void QAbstractSpinBox_dragEnterEvent_default(void* self, void* event) {
	((QAbstractSpinBox__Go*)self)->cppgoBase_dragEnterEvent((QDragEnterEvent*)event);
}
void QAbstractSpinBox_dragMoveEvent_default(void* self, void* event) {
	((QAbstractSpinBox__Go*)self)->cppgoBase_dragMoveEvent((QDragMoveEvent*)event);
}
void QAbstractSpinBox_dragLeaveEvent_default(void* self, void* event) {
	((QAbstractSpinBox__Go*)self)->cppgoBase_dragLeaveEvent((QDragLeaveEvent*)event);
}
void QAbstractSpinBox_dropEvent_default(void* self, void* event) {
	((QAbstractSpinBox__Go*)self)->cppgoBase_dropEvent((QDropEvent*)event);
}
int QAbstractSpinBox_metric_default(void* self, int arg1) {
	auto&& __r = ((QAbstractSpinBox__Go*)self)->cppgoBase_metric(arg1);
	return __r;
}
void QAbstractSpinBox_initPainter_default(void* self, void* painter) {
	((QAbstractSpinBox__Go*)self)->cppgoBase_initPainter((QPainter*)painter);
}
void* QAbstractSpinBox_redirected_default(void* self, void* offset) {
	auto&& __r = ((QAbstractSpinBox__Go*)self)->cppgoBase_redirected((QPoint*)offset);
	return (void*)(__r);
}
void* QAbstractSpinBox_sharedPainter_default(void* self) {
	auto&& __r = ((QAbstractSpinBox__Go*)self)->cppgoBase_sharedPainter();
	return (void*)(__r);
}
void QAbstractSpinBox_inputMethodEvent_default(void* self, void* arg1) {
	((QAbstractSpinBox__Go*)self)->cppgoBase_inputMethodEvent((QInputMethodEvent*)arg1);
}
int QAbstractSpinBox_focusNextPrevChild_default(void* self, int next) {
	auto&& __r = ((QAbstractSpinBox__Go*)self)->cppgoBase_focusNextPrevChild(next);
	return __r;
}
int QAbstractSpinBox_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QAbstractSpinBox*)self)->QAbstractSpinBox::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QAbstractSpinBox_childEvent_default(void* self, void* event) {
	((QAbstractSpinBox__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QAbstractSpinBox_customEvent_default(void* self, void* event) {
	((QAbstractSpinBox__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QAbstractSpinBox_connectNotify_default(void* self, void* signal) {
	((QAbstractSpinBox__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QAbstractSpinBox_disconnectNotify_default(void* self, void* signal) {
	((QAbstractSpinBox__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QAbstractSpinBox_connect_editingFinished(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractSpinBox*)self, &QAbstractSpinBox::editingFinished, (QAbstractSpinBox*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractSpinBox);
	__conns_QAbstractSpinBox[__id] = __c;
}
void QAbstractSpinBox_connect_returnPressed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractSpinBox*)self, &QAbstractSpinBox::returnPressed, (QAbstractSpinBox*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractSpinBox);
	__conns_QAbstractSpinBox[__id] = __c;
}
void QAbstractSpinBox_connect_windowTitleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractSpinBox*)self, &QAbstractSpinBox::windowTitleChanged, (QAbstractSpinBox*)self, [__id](const QString& title){
		QByteArray title_ba = title.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)title_ba.constData(), title_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractSpinBox);
	__conns_QAbstractSpinBox[__id] = __c;
}
void QAbstractSpinBox_connect_windowIconChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractSpinBox*)self, &QAbstractSpinBox::windowIconChanged, (QAbstractSpinBox*)self, [__id](const QIcon& icon){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&icon);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractSpinBox);
	__conns_QAbstractSpinBox[__id] = __c;
}
void QAbstractSpinBox_connect_windowIconTextChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractSpinBox*)self, &QAbstractSpinBox::windowIconTextChanged, (QAbstractSpinBox*)self, [__id](const QString& iconText){
		QByteArray iconText_ba = iconText.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)iconText_ba.constData(), iconText_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractSpinBox);
	__conns_QAbstractSpinBox[__id] = __c;
}
void QAbstractSpinBox_connect_customContextMenuRequested(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractSpinBox*)self, &QAbstractSpinBox::customContextMenuRequested, (QAbstractSpinBox*)self, [__id](const QPoint& pos){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&pos);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractSpinBox);
	__conns_QAbstractSpinBox[__id] = __c;
}
void QAbstractSpinBox_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractSpinBox*)self, &QAbstractSpinBox::destroyed, (QAbstractSpinBox*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractSpinBox);
	__conns_QAbstractSpinBox[__id] = __c;
}
void QAbstractSpinBox_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractSpinBox*)self, &QAbstractSpinBox::destroyed, (QAbstractSpinBox*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractSpinBox);
	__conns_QAbstractSpinBox[__id] = __c;
}
void QAbstractSpinBox_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractSpinBox*)self, &QAbstractSpinBox::objectNameChanged, (QAbstractSpinBox*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractSpinBox);
	__conns_QAbstractSpinBox[__id] = __c;
}
void QAbstractSpinBox__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QAbstractSpinBox);
		auto __it = __conns_QAbstractSpinBox.find(id);
		if (__it == __conns_QAbstractSpinBox.end()) return;
		__c = __it->second;
		__conns_QAbstractSpinBox.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QAbstractSpinBox__Go_new(long long id) { return new QAbstractSpinBox__Go(id); }
void QAbstractSpinBox_delete(void* self) { delete (QAbstractSpinBox*)self; }

}

#include <QtWidgets/qstyleoption.h>
#include <QtWidgets/qwidget.h>
#include <QtWidgets/qrubberband.h>
#include "gen_qstyleoptionrubberband.h"

extern "C" {

void* QStyleOptionRubberBand_base(void* self) { return (void*)static_cast<QStyleOption*>((QStyleOptionRubberBand*)self); }
void* QStyleOptionRubberBand_new() {
	return new QStyleOptionRubberBand();
}
void* QStyleOptionRubberBand_new2(void* other) {
	return new QStyleOptionRubberBand(*(QStyleOptionRubberBand*)other);
}
void* QStyleOptionRubberBand_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QStyleOptionRubberBand*)self)->operator=(*(QStyleOptionRubberBand*)arg1);
	return (void*)&(__r);
}
int QStyleOptionRubberBand_shape_get(void* self) {
	auto __r = ((QStyleOptionRubberBand*)self)->shape;
	return (int)(__r);
}
void QStyleOptionRubberBand_shape_set(void* self, int shape) {
	((QStyleOptionRubberBand*)self)->shape = (QRubberBand::Shape)(shape);
}
int QStyleOptionRubberBand_opaque_get(void* self) {
	auto __r = ((QStyleOptionRubberBand*)self)->opaque;
	return __r;
}
void QStyleOptionRubberBand_opaque_set(void* self, int opaque) {
	((QStyleOptionRubberBand*)self)->opaque = opaque;
}
void QStyleOptionRubberBand_delete(void* self) { delete (QStyleOptionRubberBand*)self; }

}

#include <QtWidgets/qgridlayout.h>
#include <QtWidgets/qlayout.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qsize.h>
#include <QtCore/qrect.h>
#include <QtCore/qnamespace.h>
#include <QtWidgets/qlayoutitem.h>
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
#include "gen_qgridlayout.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QGridLayout__Go : public QGridLayout {
public:
	long long __goId;
	unsigned char __goOn[32] = {};
	template<typename... A> QGridLayout__Go(long long id, A&&... a) : QGridLayout(std::forward<A>(a)...), __goId(id) {}
	using QGridLayout::widgetEvent;
	using QGridLayout::addChildLayout;
	using QGridLayout::addChildWidget;
	using QGridLayout::adoptLayout;
	using QGridLayout::alignmentRect;
	using QGridLayout::sender;
	using QGridLayout::senderSignalIndex;
	using QGridLayout::receivers;
	using QGridLayout::isSignalConnected;
	void cppgoBase_childEvent(QChildEvent* __a0) { QGridLayout::childEvent(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QGridLayout::timerEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QGridLayout::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QGridLayout::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QGridLayout::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QGridLayout::metaObject();
	}
	QSize sizeHint() const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 729690261457684101LL, &__h); if (__h) return *(QSize*)__r; }
		return QGridLayout::sizeHint();
	}
	QSize minimumSize() const override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, -7246748916541192690LL, &__h); if (__h) return *(QSize*)__r; }
		return QGridLayout::minimumSize();
	}
	QSize maximumSize() const override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 4158099606088301136LL, &__h); if (__h) return *(QSize*)__r; }
		return QGridLayout::maximumSize();
	}
	void setSpacing(int spacing) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 4039479028406881496LL, spacing, &__h); if (__h) return; }
		QGridLayout::setSpacing(spacing);
	}
	int spacing() const override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -4587105490415799682LL, &__h); if (__h) return __r; }
		return QGridLayout::spacing();
	}
	bool hasHeightForWidth() const override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8266762397995423915LL, &__h); if (__h) return __r; }
		return QGridLayout::hasHeightForWidth();
	}
	int heightForWidth(int arg1) const override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 4311534060495737339LL, arg1, &__h); if (__h) return __r; }
		return QGridLayout::heightForWidth(arg1);
	}
	int minimumHeightForWidth(int arg1) const override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 5037837229528203129LL, arg1, &__h); if (__h) return __r; }
		return QGridLayout::minimumHeightForWidth(arg1);
	}
	QFlags<Qt::Orientation> expandingDirections() const override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, 415665469015095569LL, &__h); if (__h) return QFlags<Qt::Orientation>(__r); }
		return QGridLayout::expandingDirections();
	}
	void invalidate() override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 6991341317990632368LL, &__h); if (__h) return; }
		QGridLayout::invalidate();
	}
	QLayoutItem* itemAt(int index) const override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7508512214013391107LL, index, &__h); if (__h) return (QLayoutItem*)__r; }
		return QGridLayout::itemAt(index);
	}
	QLayoutItem* takeAt(int index) override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7549889108533486643LL, index, &__h); if (__h) return (QLayoutItem*)__r; }
		return QGridLayout::takeAt(index);
	}
	int count() const override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -5627843058667511180LL, &__h); if (__h) return __r; }
		return QGridLayout::count();
	}
	void setGeometry(const QRect& arg1) override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5684820848726969227LL, (void*)&arg1, &__h); if (__h) return; }
		QGridLayout::setGeometry(arg1);
	}
	void addItem(QLayoutItem* arg1) override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6884465960736393867LL, (void*)arg1, &__h); if (__h) return; }
		QGridLayout::addItem(arg1);
	}
	QRect geometry() const override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1050490530495513297LL, &__h); if (__h) return *(QRect*)__r; }
		return QGridLayout::geometry();
	}
	int indexOf(const QWidget* arg1) const override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, -7258473138030329570LL, (void*)arg1, &__h); if (__h) return __r; }
		return QGridLayout::indexOf(arg1);
	}
	int indexOf(const QLayoutItem* arg1) const override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, -3276435765994963516LL, (void*)arg1, &__h); if (__h) return __r; }
		return QGridLayout::indexOf(arg1);
	}
	bool isEmpty() const override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -626822430251098498LL, &__h); if (__h) return __r; }
		return QGridLayout::isEmpty();
	}
	QFlags<QSizePolicy::ControlType> controlTypes() const override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, 3278947275714896887LL, &__h); if (__h) return QFlags<QSizePolicy::ControlType>(__r); }
		return QGridLayout::controlTypes();
	}
	QLayoutItem* replaceWidget(QWidget* from, QWidget* to, QFlags<Qt::FindChildOption> options) override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_voidp_voidp_voidp_int) { void* __r = ((void*(*)(long long, long long, void*, void*, int, int*))cppgo_virt_voidp_voidp_voidp_int)(__goId, -712862742159708061LL, (void*)from, (void*)to, (int)(options), &__h); if (__h) return (QLayoutItem*)__r; }
		return QGridLayout::replaceWidget(from, to, options);
	}
	QLayout* layout() override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, -7183390618627992737LL, &__h); if (__h) return (QLayout*)__r; }
		return QGridLayout::layout();
	}
	void childEvent(QChildEvent* e) override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)e, &__h); if (__h) return; }
		QGridLayout::childEvent(e);
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QGridLayout::event(event);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QGridLayout::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QGridLayout::timerEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QGridLayout::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QGridLayout::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QGridLayout::disconnectNotify(signal);
	}
	QWidget* widget() const override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, -4575937485689502171LL, &__h); if (__h) return (QWidget*)__r; }
		return QGridLayout::widget();
	}
	QSpacerItem* spacerItem() override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 2272240669359776278LL, &__h); if (__h) return (QSpacerItem*)__r; }
		return QGridLayout::spacerItem();
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QGridLayout;
static std::mutex __connsMu_QGridLayout;

extern "C" {

void* QGridLayout_base(void* self) { return (void*)static_cast<QLayout*>((QGridLayout*)self); }
void QGridLayout_chain(void* self, void** out) {
	QGridLayout* __p = (QGridLayout*)self;
	QLayout* __b0 = static_cast<QLayout*>(__p); out[0] = (void*)__b0;
	QObject* __b1 = static_cast<QObject*>(__b0); out[1] = (void*)__b1;
}
void QGridLayout__set_override(void* self, int idx) { ((QGridLayout__Go*)self)->__goOn[idx] = 1; }
void* QGridLayout_fromQObject(void* o) { return (void*)qobject_cast<QGridLayout*>((QObject*)o); }
void* QGridLayout_new(long long __goid, void* parent) {
	return new QGridLayout__Go(__goid, (QWidget*)parent);
}
void* QGridLayout_new2(long long __goid) {
	return new QGridLayout__Go(__goid);
}
char* QGridLayout_tr(const char* s, long long* __out_len) {
	auto&& __r = QGridLayout::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QGridLayout_setHorizontalSpacing(void* self, int spacing) {
	((QGridLayout*)self)->setHorizontalSpacing(spacing);
}
int QGridLayout_horizontalSpacing(void* self) {
	auto&& __r = ((QGridLayout*)self)->horizontalSpacing();
	return __r;
}
void QGridLayout_setVerticalSpacing(void* self, int spacing) {
	((QGridLayout*)self)->setVerticalSpacing(spacing);
}
int QGridLayout_verticalSpacing(void* self) {
	auto&& __r = ((QGridLayout*)self)->verticalSpacing();
	return __r;
}
void QGridLayout_setRowStretch(void* self, int row, int stretch) {
	((QGridLayout*)self)->setRowStretch(row, stretch);
}
void QGridLayout_setColumnStretch(void* self, int column, int stretch) {
	((QGridLayout*)self)->setColumnStretch(column, stretch);
}
int QGridLayout_rowStretch(void* self, int row) {
	auto&& __r = ((QGridLayout*)self)->rowStretch(row);
	return __r;
}
int QGridLayout_columnStretch(void* self, int column) {
	auto&& __r = ((QGridLayout*)self)->columnStretch(column);
	return __r;
}
void QGridLayout_setRowMinimumHeight(void* self, int row, int minSize) {
	((QGridLayout*)self)->setRowMinimumHeight(row, minSize);
}
void QGridLayout_setColumnMinimumWidth(void* self, int column, int minSize) {
	((QGridLayout*)self)->setColumnMinimumWidth(column, minSize);
}
int QGridLayout_rowMinimumHeight(void* self, int row) {
	auto&& __r = ((QGridLayout*)self)->rowMinimumHeight(row);
	return __r;
}
int QGridLayout_columnMinimumWidth(void* self, int column) {
	auto&& __r = ((QGridLayout*)self)->columnMinimumWidth(column);
	return __r;
}
int QGridLayout_columnCount(void* self) {
	auto&& __r = ((QGridLayout*)self)->columnCount();
	return __r;
}
int QGridLayout_rowCount(void* self) {
	auto&& __r = ((QGridLayout*)self)->rowCount();
	return __r;
}
void* QGridLayout_cellRect(void* self, int row, int column) {
	auto&& __r = ((QGridLayout*)self)->cellRect(row, column);
	return new QRect(__r);
}
void QGridLayout_addWidget(void* self, void* w) {
	((QGridLayout*)self)->addWidget((QWidget*)w);
}
void QGridLayout_addWidget2(void* self, void* arg1, int row, int column) {
	((QGridLayout*)self)->addWidget((QWidget*)arg1, row, column);
}
void QGridLayout_addWidget3(void* self, void* arg1, int row, int column, int rowSpan, int columnSpan) {
	((QGridLayout*)self)->addWidget((QWidget*)arg1, row, column, rowSpan, columnSpan);
}
void QGridLayout_addLayout(void* self, void* arg1, int row, int column) {
	((QGridLayout*)self)->addLayout((QLayout*)arg1, row, column);
}
void QGridLayout_addLayout2(void* self, void* arg1, int row, int column, int rowSpan, int columnSpan) {
	((QGridLayout*)self)->addLayout((QLayout*)arg1, row, column, rowSpan, columnSpan);
}
void QGridLayout_setOriginCorner(void* self, int arg1) {
	((QGridLayout*)self)->setOriginCorner((Qt::Corner)(arg1));
}
int QGridLayout_originCorner(void* self) {
	auto&& __r = ((QGridLayout*)self)->originCorner();
	return (int)(__r);
}
void* QGridLayout_itemAtPosition(void* self, int row, int column) {
	auto&& __r = ((QGridLayout*)self)->itemAtPosition(row, column);
	return (void*)(__r);
}
void QGridLayout_addItem(void* self, void* item, int row, int column) {
	((QGridLayout*)self)->addItem((QLayoutItem*)item, row, column);
}
void QGridLayout_setDefaultPositioning(void* self, int n, int orient) {
	((QGridLayout*)self)->setDefaultPositioning(n, (Qt::Orientation)(orient));
}
void QGridLayout_getItemPosition(void* self, int idx, int* row, int* column, int* rowSpan, int* columnSpan) {
	int row_local{};
	int column_local{};
	int rowSpan_local{};
	int columnSpan_local{};
	((QGridLayout*)self)->getItemPosition(idx, &row_local, &column_local, &rowSpan_local, &columnSpan_local);
	*row = row_local;
	*column = column_local;
	*rowSpan = rowSpan_local;
	*columnSpan = columnSpan_local;
}
void* QGridLayout_closestAcceptableSize(void* w, void* s) {
	auto&& __r = QGridLayout::closestAcceptableSize((QWidget*)w, *(QSize*)s);
	return new QSize(__r);
}
void QGridLayout_widgetEvent(void* self, void* arg1) {
	((QGridLayout__Go*)self)->widgetEvent((QEvent*)arg1);
}
void QGridLayout_addChildLayout(void* self, void* l) {
	((QGridLayout__Go*)self)->addChildLayout((QLayout*)l);
}
void QGridLayout_addChildWidget(void* self, void* w) {
	((QGridLayout__Go*)self)->addChildWidget((QWidget*)w);
}
int QGridLayout_adoptLayout(void* self, void* layout) {
	auto&& __r = ((QGridLayout__Go*)self)->adoptLayout((QLayout*)layout);
	return __r;
}
void* QGridLayout_alignmentRect(void* self, void* arg1) {
	auto&& __r = ((QGridLayout__Go*)self)->alignmentRect(*(QRect*)arg1);
	return new QRect(__r);
}
void* QGridLayout_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QGridLayout::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QGridLayout_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QGridLayout::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QGridLayout_sender(void* self) {
	auto&& __r = ((QGridLayout__Go*)self)->sender();
	return (void*)(__r);
}
int QGridLayout_senderSignalIndex(void* self) {
	auto&& __r = ((QGridLayout__Go*)self)->senderSignalIndex();
	return __r;
}
int QGridLayout_receivers(void* self, const char* signal) {
	auto&& __r = ((QGridLayout__Go*)self)->receivers(signal);
	return __r;
}
int QGridLayout_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QGridLayout__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QGridLayout_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QGridLayout::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QGridLayout_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QGridLayout::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QGridLayout_addWidget4(void* self, void* arg1, int row, int column, int arg4) {
	((QGridLayout*)self)->addWidget((QWidget*)arg1, row, column, QFlags<Qt::AlignmentFlag>(arg4));
}
void QGridLayout_addWidget5(void* self, void* arg1, int row, int column, int rowSpan, int columnSpan, int arg6) {
	((QGridLayout*)self)->addWidget((QWidget*)arg1, row, column, rowSpan, columnSpan, QFlags<Qt::AlignmentFlag>(arg6));
}
void QGridLayout_addLayout3(void* self, void* arg1, int row, int column, int arg4) {
	((QGridLayout*)self)->addLayout((QLayout*)arg1, row, column, QFlags<Qt::AlignmentFlag>(arg4));
}
void QGridLayout_addLayout4(void* self, void* arg1, int row, int column, int rowSpan, int columnSpan, int arg6) {
	((QGridLayout*)self)->addLayout((QLayout*)arg1, row, column, rowSpan, columnSpan, QFlags<Qt::AlignmentFlag>(arg6));
}
void QGridLayout_addItem2(void* self, void* item, int row, int column, int rowSpan) {
	((QGridLayout*)self)->addItem((QLayoutItem*)item, row, column, rowSpan);
}
void QGridLayout_addItem3(void* self, void* item, int row, int column, int rowSpan, int columnSpan) {
	((QGridLayout*)self)->addItem((QLayoutItem*)item, row, column, rowSpan, columnSpan);
}
void QGridLayout_addItem4(void* self, void* item, int row, int column, int rowSpan, int columnSpan, int arg6) {
	((QGridLayout*)self)->addItem((QLayoutItem*)item, row, column, rowSpan, columnSpan, QFlags<Qt::AlignmentFlag>(arg6));
}
void* QGridLayout_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QGridLayout::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QGridLayout_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QGridLayout::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void QGridLayout_addItem_default(void* self, void* item, int row, int column) {
	((QGridLayout*)self)->QGridLayout::addItem((QLayoutItem*)item, row, column);
}
void QGridLayout_addItem_default2(void* self, void* item, int row, int column, int rowSpan) {
	((QGridLayout*)self)->QGridLayout::addItem((QLayoutItem*)item, row, column, rowSpan);
}
void QGridLayout_addItem_default3(void* self, void* item, int row, int column, int rowSpan, int columnSpan) {
	((QGridLayout*)self)->QGridLayout::addItem((QLayoutItem*)item, row, column, rowSpan, columnSpan);
}
void QGridLayout_addItem_default4(void* self, void* item, int row, int column, int rowSpan, int columnSpan, int arg6) {
	((QGridLayout*)self)->QGridLayout::addItem((QLayoutItem*)item, row, column, rowSpan, columnSpan, QFlags<Qt::AlignmentFlag>(arg6));
}
void* QGridLayout_metaObject_default(void* self) {
	auto&& __r = ((QGridLayout*)self)->QGridLayout::metaObject();
	return (void*)(__r);
}
void* QGridLayout_sizeHint_default(void* self) {
	auto&& __r = ((QGridLayout*)self)->QGridLayout::sizeHint();
	return new QSize(__r);
}
void* QGridLayout_minimumSize_default(void* self) {
	auto&& __r = ((QGridLayout*)self)->QGridLayout::minimumSize();
	return new QSize(__r);
}
void* QGridLayout_maximumSize_default(void* self) {
	auto&& __r = ((QGridLayout*)self)->QGridLayout::maximumSize();
	return new QSize(__r);
}
void QGridLayout_setSpacing_default(void* self, int spacing) {
	((QGridLayout*)self)->QGridLayout::setSpacing(spacing);
}
int QGridLayout_spacing_default(void* self) {
	auto&& __r = ((QGridLayout*)self)->QGridLayout::spacing();
	return __r;
}
int QGridLayout_hasHeightForWidth_default(void* self) {
	auto&& __r = ((QGridLayout*)self)->QGridLayout::hasHeightForWidth();
	return __r;
}
int QGridLayout_heightForWidth_default(void* self, int arg1) {
	auto&& __r = ((QGridLayout*)self)->QGridLayout::heightForWidth(arg1);
	return __r;
}
int QGridLayout_minimumHeightForWidth_default(void* self, int arg1) {
	auto&& __r = ((QGridLayout*)self)->QGridLayout::minimumHeightForWidth(arg1);
	return __r;
}
int QGridLayout_expandingDirections_default(void* self) {
	auto&& __r = ((QGridLayout*)self)->QGridLayout::expandingDirections();
	return (int)(__r);
}
void QGridLayout_invalidate_default(void* self) {
	((QGridLayout*)self)->QGridLayout::invalidate();
}
void* QGridLayout_itemAt_default(void* self, int index) {
	auto&& __r = ((QGridLayout*)self)->QGridLayout::itemAt(index);
	return (void*)(__r);
}
void* QGridLayout_takeAt_default(void* self, int index) {
	auto&& __r = ((QGridLayout*)self)->QGridLayout::takeAt(index);
	return (void*)(__r);
}
int QGridLayout_count_default(void* self) {
	auto&& __r = ((QGridLayout*)self)->QGridLayout::count();
	return __r;
}
void QGridLayout_setGeometry_default(void* self, void* arg1) {
	((QGridLayout*)self)->QGridLayout::setGeometry(*(QRect*)arg1);
}
void* QGridLayout_geometry_default(void* self) {
	auto&& __r = ((QGridLayout*)self)->QGridLayout::geometry();
	return new QRect(__r);
}
int QGridLayout_indexOf_default(void* self, void* arg1) {
	auto&& __r = ((QGridLayout*)self)->QGridLayout::indexOf((QWidget*)arg1);
	return __r;
}
int QGridLayout_isEmpty_default(void* self) {
	auto&& __r = ((QGridLayout*)self)->QGridLayout::isEmpty();
	return __r;
}
int QGridLayout_controlTypes_default(void* self) {
	auto&& __r = ((QGridLayout*)self)->QGridLayout::controlTypes();
	return (int)(__r);
}
void* QGridLayout_replaceWidget_default(void* self, void* from, void* to, int options) {
	auto&& __r = ((QGridLayout*)self)->QGridLayout::replaceWidget((QWidget*)from, (QWidget*)to, QFlags<Qt::FindChildOption>(options));
	return (void*)(__r);
}
void* QGridLayout_layout_default(void* self) {
	auto&& __r = ((QGridLayout*)self)->QGridLayout::layout();
	return (void*)(__r);
}
void QGridLayout_childEvent_default(void* self, void* e) {
	((QGridLayout__Go*)self)->cppgoBase_childEvent((QChildEvent*)e);
}
int QGridLayout_event_default(void* self, void* event) {
	auto&& __r = ((QGridLayout*)self)->QGridLayout::event((QEvent*)event);
	return __r;
}
int QGridLayout_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QGridLayout*)self)->QGridLayout::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QGridLayout_timerEvent_default(void* self, void* event) {
	((QGridLayout__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QGridLayout_customEvent_default(void* self, void* event) {
	((QGridLayout__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QGridLayout_connectNotify_default(void* self, void* signal) {
	((QGridLayout__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QGridLayout_disconnectNotify_default(void* self, void* signal) {
	((QGridLayout__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void* QGridLayout_widget_default(void* self) {
	auto&& __r = ((QGridLayout*)self)->QGridLayout::widget();
	return (void*)(__r);
}
void* QGridLayout_spacerItem_default(void* self) {
	auto&& __r = ((QGridLayout*)self)->QGridLayout::spacerItem();
	return (void*)(__r);
}
void QGridLayout_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGridLayout*)self, &QGridLayout::destroyed, (QGridLayout*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGridLayout);
	__conns_QGridLayout[__id] = __c;
}
void QGridLayout_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGridLayout*)self, &QGridLayout::destroyed, (QGridLayout*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGridLayout);
	__conns_QGridLayout[__id] = __c;
}
void QGridLayout_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGridLayout*)self, &QGridLayout::objectNameChanged, (QGridLayout*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGridLayout);
	__conns_QGridLayout[__id] = __c;
}
void QGridLayout__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QGridLayout);
		auto __it = __conns_QGridLayout.find(id);
		if (__it == __conns_QGridLayout.end()) return;
		__c = __it->second;
		__conns_QGridLayout.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QGridLayout__Go_new(long long id) { return new QGridLayout__Go(id); }
void QGridLayout_delete(void* self) { delete (QGridLayout*)self; }

}

#include <QtWidgets/qerrormessage.h>
#include <QtWidgets/qdialog.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qsize.h>
#include <QtGui/qevent.h>
#include <QtCore/qobject.h>
#include <QtWidgets/qstyle.h>
#include <QtCore/qnamespace.h>
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
#include "gen_qerrormessage.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QErrorMessage__Go : public QErrorMessage {
public:
	long long __goId;
	unsigned char __goOn[53] = {};
	template<typename... A> QErrorMessage__Go(long long id, A&&... a) : QErrorMessage(std::forward<A>(a)...), __goId(id) {}
	using QErrorMessage::adjustPosition;
	using QErrorMessage::updateMicroFocus;
	using QErrorMessage::create;
	using QErrorMessage::destroy;
	using QErrorMessage::focusNextChild;
	using QErrorMessage::focusPreviousChild;
	using QErrorMessage::sender;
	using QErrorMessage::senderSignalIndex;
	using QErrorMessage::receivers;
	using QErrorMessage::isSignalConnected;
	using QErrorMessage::getDecodedMetricF;
	void cppgoBase_done(int __a0) { QErrorMessage::done(__a0); }
	void cppgoBase_changeEvent(QEvent* __a0) { QErrorMessage::changeEvent(__a0); }
	void cppgoBase_keyPressEvent(QKeyEvent* __a0) { QErrorMessage::keyPressEvent(__a0); }
	void cppgoBase_closeEvent(QCloseEvent* __a0) { QErrorMessage::closeEvent(__a0); }
	void cppgoBase_showEvent(QShowEvent* __a0) { QErrorMessage::showEvent(__a0); }
	void cppgoBase_resizeEvent(QResizeEvent* __a0) { QErrorMessage::resizeEvent(__a0); }
	void cppgoBase_contextMenuEvent(QContextMenuEvent* __a0) { QErrorMessage::contextMenuEvent(__a0); }
	bool cppgoBase_eventFilter(QObject* __a0, QEvent* __a1) { return QErrorMessage::eventFilter(__a0, __a1); }
	bool cppgoBase_event(QEvent* __a0) { return QErrorMessage::event(__a0); }
	void cppgoBase_mousePressEvent(QMouseEvent* __a0) { QErrorMessage::mousePressEvent(__a0); }
	void cppgoBase_mouseReleaseEvent(QMouseEvent* __a0) { QErrorMessage::mouseReleaseEvent(__a0); }
	void cppgoBase_mouseDoubleClickEvent(QMouseEvent* __a0) { QErrorMessage::mouseDoubleClickEvent(__a0); }
	void cppgoBase_mouseMoveEvent(QMouseEvent* __a0) { QErrorMessage::mouseMoveEvent(__a0); }
	void cppgoBase_wheelEvent(QWheelEvent* __a0) { QErrorMessage::wheelEvent(__a0); }
	void cppgoBase_keyReleaseEvent(QKeyEvent* __a0) { QErrorMessage::keyReleaseEvent(__a0); }
	void cppgoBase_focusInEvent(QFocusEvent* __a0) { QErrorMessage::focusInEvent(__a0); }
	void cppgoBase_focusOutEvent(QFocusEvent* __a0) { QErrorMessage::focusOutEvent(__a0); }
	void cppgoBase_enterEvent(QEnterEvent* __a0) { QErrorMessage::enterEvent(__a0); }
	void cppgoBase_leaveEvent(QEvent* __a0) { QErrorMessage::leaveEvent(__a0); }
	void cppgoBase_paintEvent(QPaintEvent* __a0) { QErrorMessage::paintEvent(__a0); }
	void cppgoBase_moveEvent(QMoveEvent* __a0) { QErrorMessage::moveEvent(__a0); }
	void cppgoBase_tabletEvent(QTabletEvent* __a0) { QErrorMessage::tabletEvent(__a0); }
	void cppgoBase_actionEvent(QActionEvent* __a0) { QErrorMessage::actionEvent(__a0); }
	void cppgoBase_dragEnterEvent(QDragEnterEvent* __a0) { QErrorMessage::dragEnterEvent(__a0); }
	void cppgoBase_dragMoveEvent(QDragMoveEvent* __a0) { QErrorMessage::dragMoveEvent(__a0); }
	void cppgoBase_dragLeaveEvent(QDragLeaveEvent* __a0) { QErrorMessage::dragLeaveEvent(__a0); }
	void cppgoBase_dropEvent(QDropEvent* __a0) { QErrorMessage::dropEvent(__a0); }
	void cppgoBase_hideEvent(QHideEvent* __a0) { QErrorMessage::hideEvent(__a0); }
	int cppgoBase_metric(int __a0) { return QErrorMessage::metric((QPaintDevice::PaintDeviceMetric)__a0); }
	void cppgoBase_initPainter(QPainter* __a0) { QErrorMessage::initPainter(__a0); }
	QPaintDevice* cppgoBase_redirected(QPoint* __a0) { return QErrorMessage::redirected(__a0); }
	QPainter* cppgoBase_sharedPainter() { return QErrorMessage::sharedPainter(); }
	void cppgoBase_inputMethodEvent(QInputMethodEvent* __a0) { QErrorMessage::inputMethodEvent(__a0); }
	bool cppgoBase_focusNextPrevChild(bool __a0) { return QErrorMessage::focusNextPrevChild(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QErrorMessage::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QErrorMessage::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QErrorMessage::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QErrorMessage::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QErrorMessage::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QErrorMessage::metaObject();
	}
	void done(int arg1) override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, -2570997490807640431LL, arg1, &__h); if (__h) return; }
		QErrorMessage::done(arg1);
	}
	void changeEvent(QEvent* e) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4959376063826165943LL, (void*)e, &__h); if (__h) return; }
		QErrorMessage::changeEvent(e);
	}
	void setVisible(bool visible) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 2110182719375245575LL, visible, &__h); if (__h) return; }
		QErrorMessage::setVisible(visible);
	}
	QSize sizeHint() const override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 729690261457684101LL, &__h); if (__h) return *(QSize*)__r; }
		return QErrorMessage::sizeHint();
	}
	QSize minimumSizeHint() const override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 7049264599837532491LL, &__h); if (__h) return *(QSize*)__r; }
		return QErrorMessage::minimumSizeHint();
	}
	void open() override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, -554057428128861975LL, &__h); if (__h) return; }
		QErrorMessage::open();
	}
	int exec() override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -1745322561779819684LL, &__h); if (__h) return __r; }
		return QErrorMessage::exec();
	}
	void accept() override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 9106131928185084681LL, &__h); if (__h) return; }
		QErrorMessage::accept();
	}
	void reject() override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 2875220963396170150LL, &__h); if (__h) return; }
		QErrorMessage::reject();
	}
	void keyPressEvent(QKeyEvent* arg1) override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1023791977679793185LL, (void*)arg1, &__h); if (__h) return; }
		QErrorMessage::keyPressEvent(arg1);
	}
	void closeEvent(QCloseEvent* arg1) override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1989937050019464783LL, (void*)arg1, &__h); if (__h) return; }
		QErrorMessage::closeEvent(arg1);
	}
	void showEvent(QShowEvent* arg1) override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6825369174751916LL, (void*)arg1, &__h); if (__h) return; }
		QErrorMessage::showEvent(arg1);
	}
	void resizeEvent(QResizeEvent* arg1) override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4302917285655014661LL, (void*)arg1, &__h); if (__h) return; }
		QErrorMessage::resizeEvent(arg1);
	}
	void contextMenuEvent(QContextMenuEvent* arg1) override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6698043222227006015LL, (void*)arg1, &__h); if (__h) return; }
		QErrorMessage::contextMenuEvent(arg1);
	}
	bool eventFilter(QObject* arg1, QEvent* arg2) override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)arg1, (void*)arg2, &__h); if (__h) return __r; }
		return QErrorMessage::eventFilter(arg1, arg2);
	}
	int devType() const override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8659679893868524044LL, &__h); if (__h) return __r; }
		return QErrorMessage::devType();
	}
	int heightForWidth(int arg1) const override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 4311534060495737339LL, arg1, &__h); if (__h) return __r; }
		return QErrorMessage::heightForWidth(arg1);
	}
	bool hasHeightForWidth() const override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8266762397995423915LL, &__h); if (__h) return __r; }
		return QErrorMessage::hasHeightForWidth();
	}
	QPaintEngine* paintEngine() const override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 6398007404756346381LL, &__h); if (__h) return (QPaintEngine*)__r; }
		return QErrorMessage::paintEngine();
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QErrorMessage::event(event);
	}
	void mousePressEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -332090898078543319LL, (void*)event, &__h); if (__h) return; }
		QErrorMessage::mousePressEvent(event);
	}
	void mouseReleaseEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4415920049297795213LL, (void*)event, &__h); if (__h) return; }
		QErrorMessage::mouseReleaseEvent(event);
	}
	void mouseDoubleClickEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4029073905697023887LL, (void*)event, &__h); if (__h) return; }
		QErrorMessage::mouseDoubleClickEvent(event);
	}
	void mouseMoveEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4255494004099806581LL, (void*)event, &__h); if (__h) return; }
		QErrorMessage::mouseMoveEvent(event);
	}
	void wheelEvent(QWheelEvent* event) override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1919988340934511316LL, (void*)event, &__h); if (__h) return; }
		QErrorMessage::wheelEvent(event);
	}
	void keyReleaseEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 772598475950050299LL, (void*)event, &__h); if (__h) return; }
		QErrorMessage::keyReleaseEvent(event);
	}
	void focusInEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 975043767033552552LL, (void*)event, &__h); if (__h) return; }
		QErrorMessage::focusInEvent(event);
	}
	void focusOutEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -385596010576048993LL, (void*)event, &__h); if (__h) return; }
		QErrorMessage::focusOutEvent(event);
	}
	void enterEvent(QEnterEvent* event) override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6316270346668237895LL, (void*)event, &__h); if (__h) return; }
		QErrorMessage::enterEvent(event);
	}
	void leaveEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -551070009623711640LL, (void*)event, &__h); if (__h) return; }
		QErrorMessage::leaveEvent(event);
	}
	void paintEvent(QPaintEvent* event) override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5290442163167458465LL, (void*)event, &__h); if (__h) return; }
		QErrorMessage::paintEvent(event);
	}
	void moveEvent(QMoveEvent* event) override {
		int __h = 0;
		if (__goOn[32] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -9161949150489525396LL, (void*)event, &__h); if (__h) return; }
		QErrorMessage::moveEvent(event);
	}
	void tabletEvent(QTabletEvent* event) override {
		int __h = 0;
		if (__goOn[33] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -212068866860701141LL, (void*)event, &__h); if (__h) return; }
		QErrorMessage::tabletEvent(event);
	}
	void actionEvent(QActionEvent* event) override {
		int __h = 0;
		if (__goOn[34] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1697444140088555581LL, (void*)event, &__h); if (__h) return; }
		QErrorMessage::actionEvent(event);
	}
	void dragEnterEvent(QDragEnterEvent* event) override {
		int __h = 0;
		if (__goOn[35] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 206122715919917589LL, (void*)event, &__h); if (__h) return; }
		QErrorMessage::dragEnterEvent(event);
	}
	void dragMoveEvent(QDragMoveEvent* event) override {
		int __h = 0;
		if (__goOn[36] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1781545137352745978LL, (void*)event, &__h); if (__h) return; }
		QErrorMessage::dragMoveEvent(event);
	}
	void dragLeaveEvent(QDragLeaveEvent* event) override {
		int __h = 0;
		if (__goOn[37] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -3139904158373725006LL, (void*)event, &__h); if (__h) return; }
		QErrorMessage::dragLeaveEvent(event);
	}
	void dropEvent(QDropEvent* event) override {
		int __h = 0;
		if (__goOn[38] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 9107058682754591716LL, (void*)event, &__h); if (__h) return; }
		QErrorMessage::dropEvent(event);
	}
	void hideEvent(QHideEvent* event) override {
		int __h = 0;
		if (__goOn[39] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -199718707876719129LL, (void*)event, &__h); if (__h) return; }
		QErrorMessage::hideEvent(event);
	}
	bool nativeEvent(const QByteArray& eventType, void* message, qintptr* result) override {
		int __h = 0;
		if (__goOn[40] && cppgo_virt_int_voidp_voidp_longlongp) { int __r = ((int(*)(long long, long long, void*, void*, long long*, int*))cppgo_virt_int_voidp_voidp_longlongp)(__goId, -2039166586246443706LL, (void*)&eventType, (void*)message, (long long*)result, &__h); if (__h) return __r; }
		return QErrorMessage::nativeEvent(eventType, message, result);
	}
	int metric(QPaintDevice::PaintDeviceMetric arg1) const override {
		int __h = 0;
		if (__goOn[41] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 427629649267311777LL, (int)(arg1), &__h); if (__h) return __r; }
		return QErrorMessage::metric(arg1);
	}
	void initPainter(QPainter* painter) const override {
		int __h = 0;
		if (__goOn[42] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6477500296980050210LL, (void*)painter, &__h); if (__h) return; }
		QErrorMessage::initPainter(painter);
	}
	QPaintDevice* redirected(QPoint* offset) const override {
		int __h = 0;
		if (__goOn[43] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 4127240633310655048LL, (void*)offset, &__h); if (__h) return (QPaintDevice*)__r; }
		return QErrorMessage::redirected(offset);
	}
	QPainter* sharedPainter() const override {
		int __h = 0;
		if (__goOn[44] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 5792771328741774991LL, &__h); if (__h) return (QPainter*)__r; }
		return QErrorMessage::sharedPainter();
	}
	void inputMethodEvent(QInputMethodEvent* arg1) override {
		int __h = 0;
		if (__goOn[45] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5438819661942843222LL, (void*)arg1, &__h); if (__h) return; }
		QErrorMessage::inputMethodEvent(arg1);
	}
	QVariant inputMethodQuery(Qt::InputMethodQuery arg1) const override {
		int __h = 0;
		if (__goOn[46] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7345853097340709638LL, (int)(arg1), &__h); if (__h) return *(QVariant*)__r; }
		return QErrorMessage::inputMethodQuery(arg1);
	}
	bool focusNextPrevChild(bool next) override {
		int __h = 0;
		if (__goOn[47] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 5898328517505784191LL, next, &__h); if (__h) return __r; }
		return QErrorMessage::focusNextPrevChild(next);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[48] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QErrorMessage::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[49] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QErrorMessage::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[50] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QErrorMessage::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[51] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QErrorMessage::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[52] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QErrorMessage::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QErrorMessage;
static std::mutex __connsMu_QErrorMessage;

extern "C" {

void* QErrorMessage_base(void* self) { return (void*)static_cast<QDialog*>((QErrorMessage*)self); }
void QErrorMessage_chain(void* self, void** out) {
	QErrorMessage* __p = (QErrorMessage*)self;
	QDialog* __b0 = static_cast<QDialog*>(__p); out[0] = (void*)__b0;
	QWidget* __b1 = static_cast<QWidget*>(__b0); out[1] = (void*)__b1;
	QObject* __b2 = static_cast<QObject*>(__b1); out[2] = (void*)__b2;
}
void QErrorMessage__set_override(void* self, int idx) { ((QErrorMessage__Go*)self)->__goOn[idx] = 1; }
void* QErrorMessage_fromQObject(void* o) { return (void*)qobject_cast<QErrorMessage*>((QObject*)o); }
void* QErrorMessage_new(long long __goid, void* parent) {
	return new QErrorMessage__Go(__goid, (QWidget*)parent);
}
void* QErrorMessage_new2(long long __goid) {
	return new QErrorMessage__Go(__goid);
}
char* QErrorMessage_tr(const char* s, long long* __out_len) {
	auto&& __r = QErrorMessage::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QErrorMessage_qtHandler() {
	auto&& __r = QErrorMessage::qtHandler();
	return (void*)(__r);
}
void QErrorMessage_showMessage(void* self, const char* message, long long message_len) {
	QString message_qstr = QString::fromUtf8(message, message_len);
	((QErrorMessage*)self)->showMessage(message_qstr);
}
void QErrorMessage_showMessage2(void* self, const char* message, long long message_len, const char* type_, long long type__len) {
	QString message_qstr = QString::fromUtf8(message, message_len);
	QString type__qstr = QString::fromUtf8(type_, type__len);
	((QErrorMessage*)self)->showMessage(message_qstr, type__qstr);
}
void QErrorMessage_adjustPosition(void* self, void* arg1) {
	((QErrorMessage__Go*)self)->adjustPosition((QWidget*)arg1);
}
void QErrorMessage_setTabOrder(void* arg1, void* arg2) {
	QErrorMessage::setTabOrder((QWidget*)arg1, (QWidget*)arg2);
}
void* QErrorMessage_mouseGrabber() {
	auto&& __r = QErrorMessage::mouseGrabber();
	return (void*)(__r);
}
void* QErrorMessage_keyboardGrabber() {
	auto&& __r = QErrorMessage::keyboardGrabber();
	return (void*)(__r);
}
void* QErrorMessage_find(unsigned long long arg1) {
	auto&& __r = QErrorMessage::find(arg1);
	return (void*)(__r);
}
void* QErrorMessage_createWindowContainer(void* window) {
	auto&& __r = QErrorMessage::createWindowContainer((QWindow*)window);
	return (void*)(__r);
}
void QErrorMessage_updateMicroFocus(void* self) {
	((QErrorMessage__Go*)self)->updateMicroFocus();
}
void QErrorMessage_create(void* self) {
	((QErrorMessage__Go*)self)->create();
}
void QErrorMessage_destroy(void* self) {
	((QErrorMessage__Go*)self)->destroy();
}
int QErrorMessage_focusNextChild(void* self) {
	auto&& __r = ((QErrorMessage__Go*)self)->focusNextChild();
	return __r;
}
int QErrorMessage_focusPreviousChild(void* self) {
	auto&& __r = ((QErrorMessage__Go*)self)->focusPreviousChild();
	return __r;
}
void* QErrorMessage_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QErrorMessage::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QErrorMessage_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QErrorMessage::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QErrorMessage_sender(void* self) {
	auto&& __r = ((QErrorMessage__Go*)self)->sender();
	return (void*)(__r);
}
int QErrorMessage_senderSignalIndex(void* self) {
	auto&& __r = ((QErrorMessage__Go*)self)->senderSignalIndex();
	return __r;
}
int QErrorMessage_receivers(void* self, const char* signal) {
	auto&& __r = ((QErrorMessage__Go*)self)->receivers(signal);
	return __r;
}
int QErrorMessage_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QErrorMessage__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
double QErrorMessage_devicePixelRatioFScale() {
	auto&& __r = QErrorMessage::devicePixelRatioFScale();
	return __r;
}
int QErrorMessage_encodeMetricF(int metric, double value) {
	auto&& __r = QErrorMessage::encodeMetricF((QPaintDevice::PaintDeviceMetric)(metric), value);
	return __r;
}
double QErrorMessage_getDecodedMetricF(void* self, int metricA, int metricB) {
	auto&& __r = ((QErrorMessage__Go*)self)->getDecodedMetricF((QPaintDevice::PaintDeviceMetric)(metricA), (QPaintDevice::PaintDeviceMetric)(metricB));
	return __r;
}
char* QErrorMessage_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QErrorMessage::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QErrorMessage_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QErrorMessage::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QErrorMessage_createWindowContainer2(void* window, void* parent) {
	auto&& __r = QErrorMessage::createWindowContainer((QWindow*)window, (QWidget*)parent);
	return (void*)(__r);
}
void* QErrorMessage_createWindowContainer3(void* window, void* parent, int flags) {
	auto&& __r = QErrorMessage::createWindowContainer((QWindow*)window, (QWidget*)parent, QFlags<Qt::WindowType>(flags));
	return (void*)(__r);
}
void QErrorMessage_updateMicroFocus2(void* self, int query) {
	((QErrorMessage__Go*)self)->updateMicroFocus((Qt::InputMethodQuery)(query));
}
void QErrorMessage_create2(void* self, unsigned long long arg1) {
	((QErrorMessage__Go*)self)->create(arg1);
}
void QErrorMessage_create3(void* self, unsigned long long arg1, int initializeWindow) {
	((QErrorMessage__Go*)self)->create(arg1, initializeWindow);
}
void QErrorMessage_create4(void* self, unsigned long long arg1, int initializeWindow, int destroyOldWindow) {
	((QErrorMessage__Go*)self)->create(arg1, initializeWindow, destroyOldWindow);
}
void QErrorMessage_destroy2(void* self, int destroyWindow) {
	((QErrorMessage__Go*)self)->destroy(destroyWindow);
}
void QErrorMessage_destroy3(void* self, int destroyWindow, int destroySubWindows) {
	((QErrorMessage__Go*)self)->destroy(destroyWindow, destroySubWindows);
}
void* QErrorMessage_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QErrorMessage::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QErrorMessage_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QErrorMessage::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QErrorMessage_metaObject_default(void* self) {
	auto&& __r = ((QErrorMessage*)self)->QErrorMessage::metaObject();
	return (void*)(__r);
}
void QErrorMessage_done_default(void* self, int arg1) {
	((QErrorMessage__Go*)self)->cppgoBase_done(arg1);
}
void QErrorMessage_changeEvent_default(void* self, void* e) {
	((QErrorMessage__Go*)self)->cppgoBase_changeEvent((QEvent*)e);
}
void QErrorMessage_setVisible_default(void* self, int visible) {
	((QErrorMessage*)self)->QErrorMessage::setVisible(visible);
}
void* QErrorMessage_sizeHint_default(void* self) {
	auto&& __r = ((QErrorMessage*)self)->QErrorMessage::sizeHint();
	return new QSize(__r);
}
void* QErrorMessage_minimumSizeHint_default(void* self) {
	auto&& __r = ((QErrorMessage*)self)->QErrorMessage::minimumSizeHint();
	return new QSize(__r);
}
void QErrorMessage_open_default(void* self) {
	((QErrorMessage*)self)->QErrorMessage::open();
}
int QErrorMessage_exec_default(void* self) {
	auto&& __r = ((QErrorMessage*)self)->QErrorMessage::exec();
	return __r;
}
void QErrorMessage_accept_default(void* self) {
	((QErrorMessage*)self)->QErrorMessage::accept();
}
void QErrorMessage_reject_default(void* self) {
	((QErrorMessage*)self)->QErrorMessage::reject();
}
void QErrorMessage_keyPressEvent_default(void* self, void* arg1) {
	((QErrorMessage__Go*)self)->cppgoBase_keyPressEvent((QKeyEvent*)arg1);
}
void QErrorMessage_closeEvent_default(void* self, void* arg1) {
	((QErrorMessage__Go*)self)->cppgoBase_closeEvent((QCloseEvent*)arg1);
}
void QErrorMessage_showEvent_default(void* self, void* arg1) {
	((QErrorMessage__Go*)self)->cppgoBase_showEvent((QShowEvent*)arg1);
}
void QErrorMessage_resizeEvent_default(void* self, void* arg1) {
	((QErrorMessage__Go*)self)->cppgoBase_resizeEvent((QResizeEvent*)arg1);
}
void QErrorMessage_contextMenuEvent_default(void* self, void* arg1) {
	((QErrorMessage__Go*)self)->cppgoBase_contextMenuEvent((QContextMenuEvent*)arg1);
}
int QErrorMessage_eventFilter_default(void* self, void* arg1, void* arg2) {
	auto&& __r = ((QErrorMessage__Go*)self)->cppgoBase_eventFilter((QObject*)arg1, (QEvent*)arg2);
	return __r;
}
int QErrorMessage_devType_default(void* self) {
	auto&& __r = ((QErrorMessage*)self)->QErrorMessage::devType();
	return __r;
}
int QErrorMessage_heightForWidth_default(void* self, int arg1) {
	auto&& __r = ((QErrorMessage*)self)->QErrorMessage::heightForWidth(arg1);
	return __r;
}
int QErrorMessage_hasHeightForWidth_default(void* self) {
	auto&& __r = ((QErrorMessage*)self)->QErrorMessage::hasHeightForWidth();
	return __r;
}
void* QErrorMessage_paintEngine_default(void* self) {
	auto&& __r = ((QErrorMessage*)self)->QErrorMessage::paintEngine();
	return (void*)(__r);
}
int QErrorMessage_event_default(void* self, void* event) {
	auto&& __r = ((QErrorMessage__Go*)self)->cppgoBase_event((QEvent*)event);
	return __r;
}
void QErrorMessage_mousePressEvent_default(void* self, void* event) {
	((QErrorMessage__Go*)self)->cppgoBase_mousePressEvent((QMouseEvent*)event);
}
void QErrorMessage_mouseReleaseEvent_default(void* self, void* event) {
	((QErrorMessage__Go*)self)->cppgoBase_mouseReleaseEvent((QMouseEvent*)event);
}
void QErrorMessage_mouseDoubleClickEvent_default(void* self, void* event) {
	((QErrorMessage__Go*)self)->cppgoBase_mouseDoubleClickEvent((QMouseEvent*)event);
}
void QErrorMessage_mouseMoveEvent_default(void* self, void* event) {
	((QErrorMessage__Go*)self)->cppgoBase_mouseMoveEvent((QMouseEvent*)event);
}
void QErrorMessage_wheelEvent_default(void* self, void* event) {
	((QErrorMessage__Go*)self)->cppgoBase_wheelEvent((QWheelEvent*)event);
}
void QErrorMessage_keyReleaseEvent_default(void* self, void* event) {
	((QErrorMessage__Go*)self)->cppgoBase_keyReleaseEvent((QKeyEvent*)event);
}
void QErrorMessage_focusInEvent_default(void* self, void* event) {
	((QErrorMessage__Go*)self)->cppgoBase_focusInEvent((QFocusEvent*)event);
}
void QErrorMessage_focusOutEvent_default(void* self, void* event) {
	((QErrorMessage__Go*)self)->cppgoBase_focusOutEvent((QFocusEvent*)event);
}
void QErrorMessage_enterEvent_default(void* self, void* event) {
	((QErrorMessage__Go*)self)->cppgoBase_enterEvent((QEnterEvent*)event);
}
void QErrorMessage_leaveEvent_default(void* self, void* event) {
	((QErrorMessage__Go*)self)->cppgoBase_leaveEvent((QEvent*)event);
}
void QErrorMessage_paintEvent_default(void* self, void* event) {
	((QErrorMessage__Go*)self)->cppgoBase_paintEvent((QPaintEvent*)event);
}
void QErrorMessage_moveEvent_default(void* self, void* event) {
	((QErrorMessage__Go*)self)->cppgoBase_moveEvent((QMoveEvent*)event);
}
void QErrorMessage_tabletEvent_default(void* self, void* event) {
	((QErrorMessage__Go*)self)->cppgoBase_tabletEvent((QTabletEvent*)event);
}
void QErrorMessage_actionEvent_default(void* self, void* event) {
	((QErrorMessage__Go*)self)->cppgoBase_actionEvent((QActionEvent*)event);
}
void QErrorMessage_dragEnterEvent_default(void* self, void* event) {
	((QErrorMessage__Go*)self)->cppgoBase_dragEnterEvent((QDragEnterEvent*)event);
}
void QErrorMessage_dragMoveEvent_default(void* self, void* event) {
	((QErrorMessage__Go*)self)->cppgoBase_dragMoveEvent((QDragMoveEvent*)event);
}
void QErrorMessage_dragLeaveEvent_default(void* self, void* event) {
	((QErrorMessage__Go*)self)->cppgoBase_dragLeaveEvent((QDragLeaveEvent*)event);
}
void QErrorMessage_dropEvent_default(void* self, void* event) {
	((QErrorMessage__Go*)self)->cppgoBase_dropEvent((QDropEvent*)event);
}
void QErrorMessage_hideEvent_default(void* self, void* event) {
	((QErrorMessage__Go*)self)->cppgoBase_hideEvent((QHideEvent*)event);
}
int QErrorMessage_metric_default(void* self, int arg1) {
	auto&& __r = ((QErrorMessage__Go*)self)->cppgoBase_metric(arg1);
	return __r;
}
void QErrorMessage_initPainter_default(void* self, void* painter) {
	((QErrorMessage__Go*)self)->cppgoBase_initPainter((QPainter*)painter);
}
void* QErrorMessage_redirected_default(void* self, void* offset) {
	auto&& __r = ((QErrorMessage__Go*)self)->cppgoBase_redirected((QPoint*)offset);
	return (void*)(__r);
}
void* QErrorMessage_sharedPainter_default(void* self) {
	auto&& __r = ((QErrorMessage__Go*)self)->cppgoBase_sharedPainter();
	return (void*)(__r);
}
void QErrorMessage_inputMethodEvent_default(void* self, void* arg1) {
	((QErrorMessage__Go*)self)->cppgoBase_inputMethodEvent((QInputMethodEvent*)arg1);
}
void* QErrorMessage_inputMethodQuery_default(void* self, int arg1) {
	auto&& __r = ((QErrorMessage*)self)->QErrorMessage::inputMethodQuery((Qt::InputMethodQuery)(arg1));
	return new QVariant(__r);
}
int QErrorMessage_focusNextPrevChild_default(void* self, int next) {
	auto&& __r = ((QErrorMessage__Go*)self)->cppgoBase_focusNextPrevChild(next);
	return __r;
}
void QErrorMessage_timerEvent_default(void* self, void* event) {
	((QErrorMessage__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QErrorMessage_childEvent_default(void* self, void* event) {
	((QErrorMessage__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QErrorMessage_customEvent_default(void* self, void* event) {
	((QErrorMessage__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QErrorMessage_connectNotify_default(void* self, void* signal) {
	((QErrorMessage__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QErrorMessage_disconnectNotify_default(void* self, void* signal) {
	((QErrorMessage__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QErrorMessage_connect_finished(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QErrorMessage*)self, &QErrorMessage::finished, (QErrorMessage*)self, [__id](int result){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, result);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QErrorMessage);
	__conns_QErrorMessage[__id] = __c;
}
void QErrorMessage_connect_accepted(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QErrorMessage*)self, &QErrorMessage::accepted, (QErrorMessage*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QErrorMessage);
	__conns_QErrorMessage[__id] = __c;
}
void QErrorMessage_connect_rejected(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QErrorMessage*)self, &QErrorMessage::rejected, (QErrorMessage*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QErrorMessage);
	__conns_QErrorMessage[__id] = __c;
}
void QErrorMessage_connect_windowTitleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QErrorMessage*)self, &QErrorMessage::windowTitleChanged, (QErrorMessage*)self, [__id](const QString& title){
		QByteArray title_ba = title.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)title_ba.constData(), title_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QErrorMessage);
	__conns_QErrorMessage[__id] = __c;
}
void QErrorMessage_connect_windowIconChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QErrorMessage*)self, &QErrorMessage::windowIconChanged, (QErrorMessage*)self, [__id](const QIcon& icon){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&icon);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QErrorMessage);
	__conns_QErrorMessage[__id] = __c;
}
void QErrorMessage_connect_windowIconTextChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QErrorMessage*)self, &QErrorMessage::windowIconTextChanged, (QErrorMessage*)self, [__id](const QString& iconText){
		QByteArray iconText_ba = iconText.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)iconText_ba.constData(), iconText_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QErrorMessage);
	__conns_QErrorMessage[__id] = __c;
}
void QErrorMessage_connect_customContextMenuRequested(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QErrorMessage*)self, &QErrorMessage::customContextMenuRequested, (QErrorMessage*)self, [__id](const QPoint& pos){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&pos);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QErrorMessage);
	__conns_QErrorMessage[__id] = __c;
}
void QErrorMessage_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QErrorMessage*)self, &QErrorMessage::destroyed, (QErrorMessage*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QErrorMessage);
	__conns_QErrorMessage[__id] = __c;
}
void QErrorMessage_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QErrorMessage*)self, &QErrorMessage::destroyed, (QErrorMessage*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QErrorMessage);
	__conns_QErrorMessage[__id] = __c;
}
void QErrorMessage_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QErrorMessage*)self, &QErrorMessage::objectNameChanged, (QErrorMessage*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QErrorMessage);
	__conns_QErrorMessage[__id] = __c;
}
void QErrorMessage__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QErrorMessage);
		auto __it = __conns_QErrorMessage.find(id);
		if (__it == __conns_QErrorMessage.end()) return;
		__c = __it->second;
		__conns_QErrorMessage.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QErrorMessage__Go_new(long long id) { return new QErrorMessage__Go(id); }
void QErrorMessage_delete(void* self) { delete (QErrorMessage*)self; }

}

#include <QtWidgets/qgraphicsitem.h>
#include <QtGui/qpixmap.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qpoint.h>
#include <QtCore/qrect.h>
#include <QtGui/qpainterpath.h>
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
#include <QtGui/qtransform.h>
#include <QtWidgets/qgraphicstransform.h>
#include <QtGui/qregion.h>
#include <QtGui/qpolygon.h>
#include <QtCore/qcoreevent.h>
#include <QtWidgets/qgraphicssceneevent.h>
#include <QtGui/qevent.h>
#include "gen_qgraphicspixmapitem.h"
#include "gen_purego_cb.h"

class QGraphicsPixmapItem__Go : public QGraphicsPixmapItem {
public:
	long long __goId;
	unsigned char __goOn[35] = {};
	template<typename... A> QGraphicsPixmapItem__Go(long long id, A&&... a) : QGraphicsPixmapItem(std::forward<A>(a)...), __goId(id) {}
	using QGraphicsPixmapItem::updateMicroFocus;
	using QGraphicsPixmapItem::addToIndex;
	using QGraphicsPixmapItem::removeFromIndex;
	using QGraphicsPixmapItem::prepareGeometryChange;
	bool cppgoBase_supportsExtension(int __a0) { return QGraphicsPixmapItem::supportsExtension((QGraphicsItem::Extension)__a0); }
	void cppgoBase_setExtension(int __a0, const QVariant& __a1) { QGraphicsPixmapItem::setExtension((QGraphicsItem::Extension)__a0, __a1); }
	QVariant cppgoBase_extension(const QVariant& __a0) { return QGraphicsPixmapItem::extension(__a0); }
	bool cppgoBase_sceneEventFilter(QGraphicsItem* __a0, QEvent* __a1) { return QGraphicsPixmapItem::sceneEventFilter(__a0, __a1); }
	bool cppgoBase_sceneEvent(QEvent* __a0) { return QGraphicsPixmapItem::sceneEvent(__a0); }
	void cppgoBase_contextMenuEvent(QGraphicsSceneContextMenuEvent* __a0) { QGraphicsPixmapItem::contextMenuEvent(__a0); }
	void cppgoBase_dragEnterEvent(QGraphicsSceneDragDropEvent* __a0) { QGraphicsPixmapItem::dragEnterEvent(__a0); }
	void cppgoBase_dragLeaveEvent(QGraphicsSceneDragDropEvent* __a0) { QGraphicsPixmapItem::dragLeaveEvent(__a0); }
	void cppgoBase_dragMoveEvent(QGraphicsSceneDragDropEvent* __a0) { QGraphicsPixmapItem::dragMoveEvent(__a0); }
	void cppgoBase_dropEvent(QGraphicsSceneDragDropEvent* __a0) { QGraphicsPixmapItem::dropEvent(__a0); }
	void cppgoBase_focusInEvent(QFocusEvent* __a0) { QGraphicsPixmapItem::focusInEvent(__a0); }
	void cppgoBase_focusOutEvent(QFocusEvent* __a0) { QGraphicsPixmapItem::focusOutEvent(__a0); }
	void cppgoBase_hoverEnterEvent(QGraphicsSceneHoverEvent* __a0) { QGraphicsPixmapItem::hoverEnterEvent(__a0); }
	void cppgoBase_hoverMoveEvent(QGraphicsSceneHoverEvent* __a0) { QGraphicsPixmapItem::hoverMoveEvent(__a0); }
	void cppgoBase_hoverLeaveEvent(QGraphicsSceneHoverEvent* __a0) { QGraphicsPixmapItem::hoverLeaveEvent(__a0); }
	void cppgoBase_keyPressEvent(QKeyEvent* __a0) { QGraphicsPixmapItem::keyPressEvent(__a0); }
	void cppgoBase_keyReleaseEvent(QKeyEvent* __a0) { QGraphicsPixmapItem::keyReleaseEvent(__a0); }
	void cppgoBase_mousePressEvent(QGraphicsSceneMouseEvent* __a0) { QGraphicsPixmapItem::mousePressEvent(__a0); }
	void cppgoBase_mouseMoveEvent(QGraphicsSceneMouseEvent* __a0) { QGraphicsPixmapItem::mouseMoveEvent(__a0); }
	void cppgoBase_mouseReleaseEvent(QGraphicsSceneMouseEvent* __a0) { QGraphicsPixmapItem::mouseReleaseEvent(__a0); }
	void cppgoBase_mouseDoubleClickEvent(QGraphicsSceneMouseEvent* __a0) { QGraphicsPixmapItem::mouseDoubleClickEvent(__a0); }
	void cppgoBase_wheelEvent(QGraphicsSceneWheelEvent* __a0) { QGraphicsPixmapItem::wheelEvent(__a0); }
	void cppgoBase_inputMethodEvent(QInputMethodEvent* __a0) { QGraphicsPixmapItem::inputMethodEvent(__a0); }
	QVariant cppgoBase_inputMethodQuery(int __a0) { return QGraphicsPixmapItem::inputMethodQuery((Qt::InputMethodQuery)__a0); }
	QVariant cppgoBase_itemChange(int __a0, const QVariant& __a1) { return QGraphicsPixmapItem::itemChange((QGraphicsItem::GraphicsItemChange)__a0, __a1); }
	QRectF boundingRect() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 2510008782785038831LL, &__h); if (__h) return *(QRectF*)__r; }
		return QGraphicsPixmapItem::boundingRect();
	}
	QPainterPath shape() const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 7938019069247444294LL, &__h); if (__h) return *(QPainterPath*)__r; }
		return QGraphicsPixmapItem::shape();
	}
	bool contains(const QPointF& point) const override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 5645366511944179768LL, (void*)&point, &__h); if (__h) return __r; }
		return QGraphicsPixmapItem::contains(point);
	}
	void paint(QPainter* painter, const QStyleOptionGraphicsItem* option, QWidget* widget) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_voidp_voidp_voidp) { ((void(*)(long long, long long, void*, void*, void*, int*))cppgo_virt_void_voidp_voidp_voidp)(__goId, 6083584261053556819LL, (void*)painter, (void*)option, (void*)widget, &__h); if (__h) return; }
		QGraphicsPixmapItem::paint(painter, option, widget);
	}
	bool isObscuredBy(const QGraphicsItem* item) const override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, -7215030271914903671LL, (void*)item, &__h); if (__h) return __r; }
		return QGraphicsPixmapItem::isObscuredBy(item);
	}
	QPainterPath opaqueArea() const override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 399370970023989435LL, &__h); if (__h) return *(QPainterPath*)__r; }
		return QGraphicsPixmapItem::opaqueArea();
	}
	int type() const override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -6371403872082420691LL, &__h); if (__h) return __r; }
		return QGraphicsPixmapItem::type();
	}
	bool supportsExtension(QGraphicsItem::Extension extension) const override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, -5729166443651306418LL, (int)(extension), &__h); if (__h) return __r; }
		return QGraphicsPixmapItem::supportsExtension(extension);
	}
	void setExtension(QGraphicsItem::Extension extension, const QVariant& variant) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void_int_voidp) { ((void(*)(long long, long long, int, void*, int*))cppgo_virt_void_int_voidp)(__goId, -6730181576106171580LL, (int)(extension), (void*)&variant, &__h); if (__h) return; }
		QGraphicsPixmapItem::setExtension(extension, variant);
	}
	QVariant extension(const QVariant& variant) const override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 2270616358835251226LL, (void*)&variant, &__h); if (__h) return *(QVariant*)__r; }
		return QGraphicsPixmapItem::extension(variant);
	}
	void advance(int phase) override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, -4830836341480073475LL, phase, &__h); if (__h) return; }
		QGraphicsPixmapItem::advance(phase);
	}
	bool collidesWithItem(const QGraphicsItem* other, Qt::ItemSelectionMode mode) const override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_int_voidp_int) { int __r = ((int(*)(long long, long long, void*, int, int*))cppgo_virt_int_voidp_int)(__goId, 6469736184362593173LL, (void*)other, (int)(mode), &__h); if (__h) return __r; }
		return QGraphicsPixmapItem::collidesWithItem(other, mode);
	}
	bool collidesWithPath(const QPainterPath& path, Qt::ItemSelectionMode mode) const override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_int_voidp_int) { int __r = ((int(*)(long long, long long, void*, int, int*))cppgo_virt_int_voidp_int)(__goId, 6636575182718753853LL, (void*)&path, (int)(mode), &__h); if (__h) return __r; }
		return QGraphicsPixmapItem::collidesWithPath(path, mode);
	}
	bool sceneEventFilter(QGraphicsItem* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, -4584203218241215725LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QGraphicsPixmapItem::sceneEventFilter(watched, event);
	}
	bool sceneEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, -2846371507838587071LL, (void*)event, &__h); if (__h) return __r; }
		return QGraphicsPixmapItem::sceneEvent(event);
	}
	void contextMenuEvent(QGraphicsSceneContextMenuEvent* event) override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6698043222227006015LL, (void*)event, &__h); if (__h) return; }
		QGraphicsPixmapItem::contextMenuEvent(event);
	}
	void dragEnterEvent(QGraphicsSceneDragDropEvent* event) override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 206122715919917589LL, (void*)event, &__h); if (__h) return; }
		QGraphicsPixmapItem::dragEnterEvent(event);
	}
	void dragLeaveEvent(QGraphicsSceneDragDropEvent* event) override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -3139904158373725006LL, (void*)event, &__h); if (__h) return; }
		QGraphicsPixmapItem::dragLeaveEvent(event);
	}
	void dragMoveEvent(QGraphicsSceneDragDropEvent* event) override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1781545137352745978LL, (void*)event, &__h); if (__h) return; }
		QGraphicsPixmapItem::dragMoveEvent(event);
	}
	void dropEvent(QGraphicsSceneDragDropEvent* event) override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 9107058682754591716LL, (void*)event, &__h); if (__h) return; }
		QGraphicsPixmapItem::dropEvent(event);
	}
	void focusInEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 975043767033552552LL, (void*)event, &__h); if (__h) return; }
		QGraphicsPixmapItem::focusInEvent(event);
	}
	void focusOutEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -385596010576048993LL, (void*)event, &__h); if (__h) return; }
		QGraphicsPixmapItem::focusOutEvent(event);
	}
	void hoverEnterEvent(QGraphicsSceneHoverEvent* event) override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6932941943906681679LL, (void*)event, &__h); if (__h) return; }
		QGraphicsPixmapItem::hoverEnterEvent(event);
	}
	void hoverMoveEvent(QGraphicsSceneHoverEvent* event) override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -2745191522144316396LL, (void*)event, &__h); if (__h) return; }
		QGraphicsPixmapItem::hoverMoveEvent(event);
	}
	void hoverLeaveEvent(QGraphicsSceneHoverEvent* event) override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7944408482203603984LL, (void*)event, &__h); if (__h) return; }
		QGraphicsPixmapItem::hoverLeaveEvent(event);
	}
	void keyPressEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1023791977679793185LL, (void*)event, &__h); if (__h) return; }
		QGraphicsPixmapItem::keyPressEvent(event);
	}
	void keyReleaseEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 772598475950050299LL, (void*)event, &__h); if (__h) return; }
		QGraphicsPixmapItem::keyReleaseEvent(event);
	}
	void mousePressEvent(QGraphicsSceneMouseEvent* event) override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -332090898078543319LL, (void*)event, &__h); if (__h) return; }
		QGraphicsPixmapItem::mousePressEvent(event);
	}
	void mouseMoveEvent(QGraphicsSceneMouseEvent* event) override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4255494004099806581LL, (void*)event, &__h); if (__h) return; }
		QGraphicsPixmapItem::mouseMoveEvent(event);
	}
	void mouseReleaseEvent(QGraphicsSceneMouseEvent* event) override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4415920049297795213LL, (void*)event, &__h); if (__h) return; }
		QGraphicsPixmapItem::mouseReleaseEvent(event);
	}
	void mouseDoubleClickEvent(QGraphicsSceneMouseEvent* event) override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4029073905697023887LL, (void*)event, &__h); if (__h) return; }
		QGraphicsPixmapItem::mouseDoubleClickEvent(event);
	}
	void wheelEvent(QGraphicsSceneWheelEvent* event) override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1919988340934511316LL, (void*)event, &__h); if (__h) return; }
		QGraphicsPixmapItem::wheelEvent(event);
	}
	void inputMethodEvent(QInputMethodEvent* event) override {
		int __h = 0;
		if (__goOn[32] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5438819661942843222LL, (void*)event, &__h); if (__h) return; }
		QGraphicsPixmapItem::inputMethodEvent(event);
	}
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const override {
		int __h = 0;
		if (__goOn[33] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7345853097340709638LL, (int)(query), &__h); if (__h) return *(QVariant*)__r; }
		return QGraphicsPixmapItem::inputMethodQuery(query);
	}
	QVariant itemChange(QGraphicsItem::GraphicsItemChange change, const QVariant& value) override {
		int __h = 0;
		if (__goOn[34] && cppgo_virt_voidp_int_voidp) { void* __r = ((void*(*)(long long, long long, int, void*, int*))cppgo_virt_voidp_int_voidp)(__goId, -1888572168794878830LL, (int)(change), (void*)&value, &__h); if (__h) return *(QVariant*)__r; }
		return QGraphicsPixmapItem::itemChange(change, value);
	}
};

extern "C" {

void* QGraphicsPixmapItem_base(void* self) { return (void*)static_cast<QGraphicsItem*>((QGraphicsPixmapItem*)self); }
void QGraphicsPixmapItem__set_override(void* self, int idx) { ((QGraphicsPixmapItem__Go*)self)->__goOn[idx] = 1; }
void* QGraphicsPixmapItem_new(long long __goid) {
	return new QGraphicsPixmapItem__Go(__goid);
}
void* QGraphicsPixmapItem_new2(long long __goid, void* pixmap) {
	return new QGraphicsPixmapItem__Go(__goid, *(QPixmap*)pixmap);
}
void* QGraphicsPixmapItem_new3(long long __goid, void* parent) {
	return new QGraphicsPixmapItem__Go(__goid, (QGraphicsItem*)parent);
}
void* QGraphicsPixmapItem_new4(long long __goid, void* pixmap, void* parent) {
	return new QGraphicsPixmapItem__Go(__goid, *(QPixmap*)pixmap, (QGraphicsItem*)parent);
}
void* QGraphicsPixmapItem_pixmap(void* self) {
	auto&& __r = ((QGraphicsPixmapItem*)self)->pixmap();
	return new QPixmap(__r);
}
void QGraphicsPixmapItem_setPixmap(void* self, void* pixmap) {
	((QGraphicsPixmapItem*)self)->setPixmap(*(QPixmap*)pixmap);
}
int QGraphicsPixmapItem_transformationMode(void* self) {
	auto&& __r = ((QGraphicsPixmapItem*)self)->transformationMode();
	return (int)(__r);
}
void QGraphicsPixmapItem_setTransformationMode(void* self, int mode) {
	((QGraphicsPixmapItem*)self)->setTransformationMode((Qt::TransformationMode)(mode));
}
void* QGraphicsPixmapItem_offset(void* self) {
	auto&& __r = ((QGraphicsPixmapItem*)self)->offset();
	return new QPointF(__r);
}
void QGraphicsPixmapItem_setOffset(void* self, void* offset) {
	((QGraphicsPixmapItem*)self)->setOffset(*(QPointF*)offset);
}
void QGraphicsPixmapItem_setOffset2(void* self, double x, double y) {
	((QGraphicsPixmapItem*)self)->setOffset(x, y);
}
int QGraphicsPixmapItem_shapeMode(void* self) {
	auto&& __r = ((QGraphicsPixmapItem*)self)->shapeMode();
	return (int)(__r);
}
void QGraphicsPixmapItem_setShapeMode(void* self, int mode) {
	((QGraphicsPixmapItem*)self)->setShapeMode((QGraphicsPixmapItem::ShapeMode)(mode));
}
void QGraphicsPixmapItem_updateMicroFocus(void* self) {
	((QGraphicsPixmapItem__Go*)self)->updateMicroFocus();
}
void QGraphicsPixmapItem_addToIndex(void* self) {
	((QGraphicsPixmapItem__Go*)self)->addToIndex();
}
void QGraphicsPixmapItem_removeFromIndex(void* self) {
	((QGraphicsPixmapItem__Go*)self)->removeFromIndex();
}
void QGraphicsPixmapItem_prepareGeometryChange(void* self) {
	((QGraphicsPixmapItem__Go*)self)->prepareGeometryChange();
}
void* QGraphicsPixmapItem_boundingRect_default(void* self) {
	auto&& __r = ((QGraphicsPixmapItem*)self)->QGraphicsPixmapItem::boundingRect();
	return new QRectF(__r);
}
void* QGraphicsPixmapItem_shape_default(void* self) {
	auto&& __r = ((QGraphicsPixmapItem*)self)->QGraphicsPixmapItem::shape();
	return new QPainterPath(__r);
}
int QGraphicsPixmapItem_contains_default(void* self, void* point) {
	auto&& __r = ((QGraphicsPixmapItem*)self)->QGraphicsPixmapItem::contains(*(QPointF*)point);
	return __r;
}
void QGraphicsPixmapItem_paint_default(void* self, void* painter, void* option, void* widget) {
	((QGraphicsPixmapItem*)self)->QGraphicsPixmapItem::paint((QPainter*)painter, (QStyleOptionGraphicsItem*)option, (QWidget*)widget);
}
int QGraphicsPixmapItem_isObscuredBy_default(void* self, void* item) {
	auto&& __r = ((QGraphicsPixmapItem*)self)->QGraphicsPixmapItem::isObscuredBy((QGraphicsItem*)item);
	return __r;
}
void* QGraphicsPixmapItem_opaqueArea_default(void* self) {
	auto&& __r = ((QGraphicsPixmapItem*)self)->QGraphicsPixmapItem::opaqueArea();
	return new QPainterPath(__r);
}
int QGraphicsPixmapItem_type_default(void* self) {
	auto&& __r = ((QGraphicsPixmapItem*)self)->QGraphicsPixmapItem::type();
	return __r;
}
int QGraphicsPixmapItem_supportsExtension_default(void* self, int extension) {
	auto&& __r = ((QGraphicsPixmapItem__Go*)self)->cppgoBase_supportsExtension(extension);
	return __r;
}
void QGraphicsPixmapItem_setExtension_default(void* self, int extension, void* variant) {
	((QGraphicsPixmapItem__Go*)self)->cppgoBase_setExtension(extension, *(QVariant*)variant);
}
void* QGraphicsPixmapItem_extension_default(void* self, void* variant) {
	auto&& __r = ((QGraphicsPixmapItem__Go*)self)->cppgoBase_extension(*(QVariant*)variant);
	return new QVariant(__r);
}
void QGraphicsPixmapItem_advance_default(void* self, int phase) {
	((QGraphicsPixmapItem*)self)->QGraphicsPixmapItem::advance(phase);
}
int QGraphicsPixmapItem_collidesWithItem_default(void* self, void* other, int mode) {
	auto&& __r = ((QGraphicsPixmapItem*)self)->QGraphicsPixmapItem::collidesWithItem((QGraphicsItem*)other, (Qt::ItemSelectionMode)(mode));
	return __r;
}
int QGraphicsPixmapItem_collidesWithPath_default(void* self, void* path, int mode) {
	auto&& __r = ((QGraphicsPixmapItem*)self)->QGraphicsPixmapItem::collidesWithPath(*(QPainterPath*)path, (Qt::ItemSelectionMode)(mode));
	return __r;
}
int QGraphicsPixmapItem_sceneEventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QGraphicsPixmapItem__Go*)self)->cppgoBase_sceneEventFilter((QGraphicsItem*)watched, (QEvent*)event);
	return __r;
}
int QGraphicsPixmapItem_sceneEvent_default(void* self, void* event) {
	auto&& __r = ((QGraphicsPixmapItem__Go*)self)->cppgoBase_sceneEvent((QEvent*)event);
	return __r;
}
void QGraphicsPixmapItem_contextMenuEvent_default(void* self, void* event) {
	((QGraphicsPixmapItem__Go*)self)->cppgoBase_contextMenuEvent((QGraphicsSceneContextMenuEvent*)event);
}
void QGraphicsPixmapItem_dragEnterEvent_default(void* self, void* event) {
	((QGraphicsPixmapItem__Go*)self)->cppgoBase_dragEnterEvent((QGraphicsSceneDragDropEvent*)event);
}
void QGraphicsPixmapItem_dragLeaveEvent_default(void* self, void* event) {
	((QGraphicsPixmapItem__Go*)self)->cppgoBase_dragLeaveEvent((QGraphicsSceneDragDropEvent*)event);
}
void QGraphicsPixmapItem_dragMoveEvent_default(void* self, void* event) {
	((QGraphicsPixmapItem__Go*)self)->cppgoBase_dragMoveEvent((QGraphicsSceneDragDropEvent*)event);
}
void QGraphicsPixmapItem_dropEvent_default(void* self, void* event) {
	((QGraphicsPixmapItem__Go*)self)->cppgoBase_dropEvent((QGraphicsSceneDragDropEvent*)event);
}
void QGraphicsPixmapItem_focusInEvent_default(void* self, void* event) {
	((QGraphicsPixmapItem__Go*)self)->cppgoBase_focusInEvent((QFocusEvent*)event);
}
void QGraphicsPixmapItem_focusOutEvent_default(void* self, void* event) {
	((QGraphicsPixmapItem__Go*)self)->cppgoBase_focusOutEvent((QFocusEvent*)event);
}
void QGraphicsPixmapItem_hoverEnterEvent_default(void* self, void* event) {
	((QGraphicsPixmapItem__Go*)self)->cppgoBase_hoverEnterEvent((QGraphicsSceneHoverEvent*)event);
}
void QGraphicsPixmapItem_hoverMoveEvent_default(void* self, void* event) {
	((QGraphicsPixmapItem__Go*)self)->cppgoBase_hoverMoveEvent((QGraphicsSceneHoverEvent*)event);
}
void QGraphicsPixmapItem_hoverLeaveEvent_default(void* self, void* event) {
	((QGraphicsPixmapItem__Go*)self)->cppgoBase_hoverLeaveEvent((QGraphicsSceneHoverEvent*)event);
}
void QGraphicsPixmapItem_keyPressEvent_default(void* self, void* event) {
	((QGraphicsPixmapItem__Go*)self)->cppgoBase_keyPressEvent((QKeyEvent*)event);
}
void QGraphicsPixmapItem_keyReleaseEvent_default(void* self, void* event) {
	((QGraphicsPixmapItem__Go*)self)->cppgoBase_keyReleaseEvent((QKeyEvent*)event);
}
void QGraphicsPixmapItem_mousePressEvent_default(void* self, void* event) {
	((QGraphicsPixmapItem__Go*)self)->cppgoBase_mousePressEvent((QGraphicsSceneMouseEvent*)event);
}
void QGraphicsPixmapItem_mouseMoveEvent_default(void* self, void* event) {
	((QGraphicsPixmapItem__Go*)self)->cppgoBase_mouseMoveEvent((QGraphicsSceneMouseEvent*)event);
}
void QGraphicsPixmapItem_mouseReleaseEvent_default(void* self, void* event) {
	((QGraphicsPixmapItem__Go*)self)->cppgoBase_mouseReleaseEvent((QGraphicsSceneMouseEvent*)event);
}
void QGraphicsPixmapItem_mouseDoubleClickEvent_default(void* self, void* event) {
	((QGraphicsPixmapItem__Go*)self)->cppgoBase_mouseDoubleClickEvent((QGraphicsSceneMouseEvent*)event);
}
void QGraphicsPixmapItem_wheelEvent_default(void* self, void* event) {
	((QGraphicsPixmapItem__Go*)self)->cppgoBase_wheelEvent((QGraphicsSceneWheelEvent*)event);
}
void QGraphicsPixmapItem_inputMethodEvent_default(void* self, void* event) {
	((QGraphicsPixmapItem__Go*)self)->cppgoBase_inputMethodEvent((QInputMethodEvent*)event);
}
void* QGraphicsPixmapItem_inputMethodQuery_default(void* self, int query) {
	auto&& __r = ((QGraphicsPixmapItem__Go*)self)->cppgoBase_inputMethodQuery(query);
	return new QVariant(__r);
}
void* QGraphicsPixmapItem_itemChange_default(void* self, int change, void* value) {
	auto&& __r = ((QGraphicsPixmapItem__Go*)self)->cppgoBase_itemChange(change, *(QVariant*)value);
	return new QVariant(__r);
}
void* QGraphicsPixmapItem__Go_new(long long id) { return new QGraphicsPixmapItem__Go(id); }
void QGraphicsPixmapItem_delete(void* self) { delete (QGraphicsPixmapItem*)self; }

}

#include <QtWidgets/qgraphicssceneevent.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qpoint.h>
#include <QtWidgets/qwidget.h>
#include "gen_qgraphicsscenehelpevent.h"
#include "gen_purego_cb.h"

class QGraphicsSceneHelpEvent__Go : public QGraphicsSceneHelpEvent {
public:
	long long __goId;
	unsigned char __goOn[2] = {};
	template<typename... A> QGraphicsSceneHelpEvent__Go(long long id, A&&... a) : QGraphicsSceneHelpEvent(std::forward<A>(a)...), __goId(id) {}
	void setAccepted(bool accepted) override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 6367896784077321574LL, accepted, &__h); if (__h) return; }
		QGraphicsSceneHelpEvent::setAccepted(accepted);
	}
	QEvent* clone() const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, -7032813281048616478LL, &__h); if (__h) return (QEvent*)__r; }
		return QGraphicsSceneHelpEvent::clone();
	}
};

extern "C" {

void* QGraphicsSceneHelpEvent_base(void* self) { return (void*)static_cast<QGraphicsSceneEvent*>((QGraphicsSceneHelpEvent*)self); }
void QGraphicsSceneHelpEvent_chain(void* self, void** out) {
	QGraphicsSceneHelpEvent* __p = (QGraphicsSceneHelpEvent*)self;
	QGraphicsSceneEvent* __b0 = static_cast<QGraphicsSceneEvent*>(__p); out[0] = (void*)__b0;
	QEvent* __b1 = static_cast<QEvent*>(__b0); out[1] = (void*)__b1;
}
void QGraphicsSceneHelpEvent__set_override(void* self, int idx) { ((QGraphicsSceneHelpEvent__Go*)self)->__goOn[idx] = 1; }
void* QGraphicsSceneHelpEvent_new(long long __goid) {
	return new QGraphicsSceneHelpEvent__Go(__goid);
}
void* QGraphicsSceneHelpEvent_new2(long long __goid, int type_) {
	return new QGraphicsSceneHelpEvent__Go(__goid, (QEvent::Type)(type_));
}
void* QGraphicsSceneHelpEvent_scenePos(void* self) {
	auto&& __r = ((QGraphicsSceneHelpEvent*)self)->scenePos();
	return new QPointF(__r);
}
void QGraphicsSceneHelpEvent_setScenePos(void* self, void* pos) {
	((QGraphicsSceneHelpEvent*)self)->setScenePos(*(QPointF*)pos);
}
void* QGraphicsSceneHelpEvent_screenPos(void* self) {
	auto&& __r = ((QGraphicsSceneHelpEvent*)self)->screenPos();
	return new QPoint(__r);
}
void QGraphicsSceneHelpEvent_setScreenPos(void* self, void* pos) {
	((QGraphicsSceneHelpEvent*)self)->setScreenPos(*(QPoint*)pos);
}
int QGraphicsSceneHelpEvent_registerEventType() {
	auto&& __r = QGraphicsSceneHelpEvent::registerEventType();
	return __r;
}
int QGraphicsSceneHelpEvent_registerEventType2(int hint) {
	auto&& __r = QGraphicsSceneHelpEvent::registerEventType(hint);
	return __r;
}
void QGraphicsSceneHelpEvent_setAccepted_default(void* self, int accepted) {
	((QGraphicsSceneHelpEvent*)self)->QGraphicsSceneHelpEvent::setAccepted(accepted);
}
void* QGraphicsSceneHelpEvent_clone_default(void* self) {
	auto&& __r = ((QGraphicsSceneHelpEvent*)self)->QGraphicsSceneHelpEvent::clone();
	return (void*)(__r);
}
void* QGraphicsSceneHelpEvent__Go_new(long long id) { return new QGraphicsSceneHelpEvent__Go(id); }
void QGraphicsSceneHelpEvent_delete(void* self) { delete (QGraphicsSceneHelpEvent*)self; }

}

#include <QtWidgets/qmdisubwindow.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qsize.h>
#include <QtWidgets/qmenu.h>
#include <QtWidgets/qmdiarea.h>
#include <QtCore/qobject.h>
#include <QtCore/qcoreevent.h>
#include <QtGui/qevent.h>
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
#include "gen_qmdisubwindow.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QMdiSubWindow__Go : public QMdiSubWindow {
public:
	long long __goId;
	unsigned char __goOn[48] = {};
	template<typename... A> QMdiSubWindow__Go(long long id, A&&... a) : QMdiSubWindow(std::forward<A>(a)...), __goId(id) {}
	using QMdiSubWindow::updateMicroFocus;
	using QMdiSubWindow::create;
	using QMdiSubWindow::destroy;
	using QMdiSubWindow::focusNextChild;
	using QMdiSubWindow::focusPreviousChild;
	using QMdiSubWindow::sender;
	using QMdiSubWindow::senderSignalIndex;
	using QMdiSubWindow::receivers;
	using QMdiSubWindow::isSignalConnected;
	using QMdiSubWindow::getDecodedMetricF;
	bool cppgoBase_eventFilter(QObject* __a0, QEvent* __a1) { return QMdiSubWindow::eventFilter(__a0, __a1); }
	bool cppgoBase_event(QEvent* __a0) { return QMdiSubWindow::event(__a0); }
	void cppgoBase_showEvent(QShowEvent* __a0) { QMdiSubWindow::showEvent(__a0); }
	void cppgoBase_hideEvent(QHideEvent* __a0) { QMdiSubWindow::hideEvent(__a0); }
	void cppgoBase_changeEvent(QEvent* __a0) { QMdiSubWindow::changeEvent(__a0); }
	void cppgoBase_closeEvent(QCloseEvent* __a0) { QMdiSubWindow::closeEvent(__a0); }
	void cppgoBase_leaveEvent(QEvent* __a0) { QMdiSubWindow::leaveEvent(__a0); }
	void cppgoBase_resizeEvent(QResizeEvent* __a0) { QMdiSubWindow::resizeEvent(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QMdiSubWindow::timerEvent(__a0); }
	void cppgoBase_moveEvent(QMoveEvent* __a0) { QMdiSubWindow::moveEvent(__a0); }
	void cppgoBase_paintEvent(QPaintEvent* __a0) { QMdiSubWindow::paintEvent(__a0); }
	void cppgoBase_mousePressEvent(QMouseEvent* __a0) { QMdiSubWindow::mousePressEvent(__a0); }
	void cppgoBase_mouseDoubleClickEvent(QMouseEvent* __a0) { QMdiSubWindow::mouseDoubleClickEvent(__a0); }
	void cppgoBase_mouseReleaseEvent(QMouseEvent* __a0) { QMdiSubWindow::mouseReleaseEvent(__a0); }
	void cppgoBase_mouseMoveEvent(QMouseEvent* __a0) { QMdiSubWindow::mouseMoveEvent(__a0); }
	void cppgoBase_keyPressEvent(QKeyEvent* __a0) { QMdiSubWindow::keyPressEvent(__a0); }
	void cppgoBase_contextMenuEvent(QContextMenuEvent* __a0) { QMdiSubWindow::contextMenuEvent(__a0); }
	void cppgoBase_focusInEvent(QFocusEvent* __a0) { QMdiSubWindow::focusInEvent(__a0); }
	void cppgoBase_focusOutEvent(QFocusEvent* __a0) { QMdiSubWindow::focusOutEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QMdiSubWindow::childEvent(__a0); }
	void cppgoBase_wheelEvent(QWheelEvent* __a0) { QMdiSubWindow::wheelEvent(__a0); }
	void cppgoBase_keyReleaseEvent(QKeyEvent* __a0) { QMdiSubWindow::keyReleaseEvent(__a0); }
	void cppgoBase_enterEvent(QEnterEvent* __a0) { QMdiSubWindow::enterEvent(__a0); }
	void cppgoBase_tabletEvent(QTabletEvent* __a0) { QMdiSubWindow::tabletEvent(__a0); }
	void cppgoBase_actionEvent(QActionEvent* __a0) { QMdiSubWindow::actionEvent(__a0); }
	void cppgoBase_dragEnterEvent(QDragEnterEvent* __a0) { QMdiSubWindow::dragEnterEvent(__a0); }
	void cppgoBase_dragMoveEvent(QDragMoveEvent* __a0) { QMdiSubWindow::dragMoveEvent(__a0); }
	void cppgoBase_dragLeaveEvent(QDragLeaveEvent* __a0) { QMdiSubWindow::dragLeaveEvent(__a0); }
	void cppgoBase_dropEvent(QDropEvent* __a0) { QMdiSubWindow::dropEvent(__a0); }
	int cppgoBase_metric(int __a0) { return QMdiSubWindow::metric((QPaintDevice::PaintDeviceMetric)__a0); }
	void cppgoBase_initPainter(QPainter* __a0) { QMdiSubWindow::initPainter(__a0); }
	QPaintDevice* cppgoBase_redirected(QPoint* __a0) { return QMdiSubWindow::redirected(__a0); }
	QPainter* cppgoBase_sharedPainter() { return QMdiSubWindow::sharedPainter(); }
	void cppgoBase_inputMethodEvent(QInputMethodEvent* __a0) { QMdiSubWindow::inputMethodEvent(__a0); }
	bool cppgoBase_focusNextPrevChild(bool __a0) { return QMdiSubWindow::focusNextPrevChild(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QMdiSubWindow::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QMdiSubWindow::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QMdiSubWindow::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QMdiSubWindow::metaObject();
	}
	QSize sizeHint() const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 729690261457684101LL, &__h); if (__h) return *(QSize*)__r; }
		return QMdiSubWindow::sizeHint();
	}
	QSize minimumSizeHint() const override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 7049264599837532491LL, &__h); if (__h) return *(QSize*)__r; }
		return QMdiSubWindow::minimumSizeHint();
	}
	bool eventFilter(QObject* object, QEvent* event) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)object, (void*)event, &__h); if (__h) return __r; }
		return QMdiSubWindow::eventFilter(object, event);
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QMdiSubWindow::event(event);
	}
	void showEvent(QShowEvent* showEvent) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6825369174751916LL, (void*)showEvent, &__h); if (__h) return; }
		QMdiSubWindow::showEvent(showEvent);
	}
	void hideEvent(QHideEvent* hideEvent) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -199718707876719129LL, (void*)hideEvent, &__h); if (__h) return; }
		QMdiSubWindow::hideEvent(hideEvent);
	}
	void changeEvent(QEvent* changeEvent) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4959376063826165943LL, (void*)changeEvent, &__h); if (__h) return; }
		QMdiSubWindow::changeEvent(changeEvent);
	}
	void closeEvent(QCloseEvent* closeEvent) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1989937050019464783LL, (void*)closeEvent, &__h); if (__h) return; }
		QMdiSubWindow::closeEvent(closeEvent);
	}
	void leaveEvent(QEvent* leaveEvent) override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -551070009623711640LL, (void*)leaveEvent, &__h); if (__h) return; }
		QMdiSubWindow::leaveEvent(leaveEvent);
	}
	void resizeEvent(QResizeEvent* resizeEvent) override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4302917285655014661LL, (void*)resizeEvent, &__h); if (__h) return; }
		QMdiSubWindow::resizeEvent(resizeEvent);
	}
	void timerEvent(QTimerEvent* timerEvent) override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)timerEvent, &__h); if (__h) return; }
		QMdiSubWindow::timerEvent(timerEvent);
	}
	void moveEvent(QMoveEvent* moveEvent) override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -9161949150489525396LL, (void*)moveEvent, &__h); if (__h) return; }
		QMdiSubWindow::moveEvent(moveEvent);
	}
	void paintEvent(QPaintEvent* paintEvent) override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5290442163167458465LL, (void*)paintEvent, &__h); if (__h) return; }
		QMdiSubWindow::paintEvent(paintEvent);
	}
	void mousePressEvent(QMouseEvent* mouseEvent) override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -332090898078543319LL, (void*)mouseEvent, &__h); if (__h) return; }
		QMdiSubWindow::mousePressEvent(mouseEvent);
	}
	void mouseDoubleClickEvent(QMouseEvent* mouseEvent) override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4029073905697023887LL, (void*)mouseEvent, &__h); if (__h) return; }
		QMdiSubWindow::mouseDoubleClickEvent(mouseEvent);
	}
	void mouseReleaseEvent(QMouseEvent* mouseEvent) override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4415920049297795213LL, (void*)mouseEvent, &__h); if (__h) return; }
		QMdiSubWindow::mouseReleaseEvent(mouseEvent);
	}
	void mouseMoveEvent(QMouseEvent* mouseEvent) override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4255494004099806581LL, (void*)mouseEvent, &__h); if (__h) return; }
		QMdiSubWindow::mouseMoveEvent(mouseEvent);
	}
	void keyPressEvent(QKeyEvent* keyEvent) override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1023791977679793185LL, (void*)keyEvent, &__h); if (__h) return; }
		QMdiSubWindow::keyPressEvent(keyEvent);
	}
	void contextMenuEvent(QContextMenuEvent* contextMenuEvent) override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6698043222227006015LL, (void*)contextMenuEvent, &__h); if (__h) return; }
		QMdiSubWindow::contextMenuEvent(contextMenuEvent);
	}
	void focusInEvent(QFocusEvent* focusInEvent) override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 975043767033552552LL, (void*)focusInEvent, &__h); if (__h) return; }
		QMdiSubWindow::focusInEvent(focusInEvent);
	}
	void focusOutEvent(QFocusEvent* focusOutEvent) override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -385596010576048993LL, (void*)focusOutEvent, &__h); if (__h) return; }
		QMdiSubWindow::focusOutEvent(focusOutEvent);
	}
	void childEvent(QChildEvent* childEvent) override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)childEvent, &__h); if (__h) return; }
		QMdiSubWindow::childEvent(childEvent);
	}
	int devType() const override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8659679893868524044LL, &__h); if (__h) return __r; }
		return QMdiSubWindow::devType();
	}
	void setVisible(bool visible) override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 2110182719375245575LL, visible, &__h); if (__h) return; }
		QMdiSubWindow::setVisible(visible);
	}
	int heightForWidth(int arg1) const override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 4311534060495737339LL, arg1, &__h); if (__h) return __r; }
		return QMdiSubWindow::heightForWidth(arg1);
	}
	bool hasHeightForWidth() const override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8266762397995423915LL, &__h); if (__h) return __r; }
		return QMdiSubWindow::hasHeightForWidth();
	}
	QPaintEngine* paintEngine() const override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 6398007404756346381LL, &__h); if (__h) return (QPaintEngine*)__r; }
		return QMdiSubWindow::paintEngine();
	}
	void wheelEvent(QWheelEvent* event) override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1919988340934511316LL, (void*)event, &__h); if (__h) return; }
		QMdiSubWindow::wheelEvent(event);
	}
	void keyReleaseEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 772598475950050299LL, (void*)event, &__h); if (__h) return; }
		QMdiSubWindow::keyReleaseEvent(event);
	}
	void enterEvent(QEnterEvent* event) override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6316270346668237895LL, (void*)event, &__h); if (__h) return; }
		QMdiSubWindow::enterEvent(event);
	}
	void tabletEvent(QTabletEvent* event) override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -212068866860701141LL, (void*)event, &__h); if (__h) return; }
		QMdiSubWindow::tabletEvent(event);
	}
	void actionEvent(QActionEvent* event) override {
		int __h = 0;
		if (__goOn[32] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1697444140088555581LL, (void*)event, &__h); if (__h) return; }
		QMdiSubWindow::actionEvent(event);
	}
	void dragEnterEvent(QDragEnterEvent* event) override {
		int __h = 0;
		if (__goOn[33] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 206122715919917589LL, (void*)event, &__h); if (__h) return; }
		QMdiSubWindow::dragEnterEvent(event);
	}
	void dragMoveEvent(QDragMoveEvent* event) override {
		int __h = 0;
		if (__goOn[34] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1781545137352745978LL, (void*)event, &__h); if (__h) return; }
		QMdiSubWindow::dragMoveEvent(event);
	}
	void dragLeaveEvent(QDragLeaveEvent* event) override {
		int __h = 0;
		if (__goOn[35] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -3139904158373725006LL, (void*)event, &__h); if (__h) return; }
		QMdiSubWindow::dragLeaveEvent(event);
	}
	void dropEvent(QDropEvent* event) override {
		int __h = 0;
		if (__goOn[36] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 9107058682754591716LL, (void*)event, &__h); if (__h) return; }
		QMdiSubWindow::dropEvent(event);
	}
	bool nativeEvent(const QByteArray& eventType, void* message, qintptr* result) override {
		int __h = 0;
		if (__goOn[37] && cppgo_virt_int_voidp_voidp_longlongp) { int __r = ((int(*)(long long, long long, void*, void*, long long*, int*))cppgo_virt_int_voidp_voidp_longlongp)(__goId, -2039166586246443706LL, (void*)&eventType, (void*)message, (long long*)result, &__h); if (__h) return __r; }
		return QMdiSubWindow::nativeEvent(eventType, message, result);
	}
	int metric(QPaintDevice::PaintDeviceMetric arg1) const override {
		int __h = 0;
		if (__goOn[38] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 427629649267311777LL, (int)(arg1), &__h); if (__h) return __r; }
		return QMdiSubWindow::metric(arg1);
	}
	void initPainter(QPainter* painter) const override {
		int __h = 0;
		if (__goOn[39] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6477500296980050210LL, (void*)painter, &__h); if (__h) return; }
		QMdiSubWindow::initPainter(painter);
	}
	QPaintDevice* redirected(QPoint* offset) const override {
		int __h = 0;
		if (__goOn[40] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 4127240633310655048LL, (void*)offset, &__h); if (__h) return (QPaintDevice*)__r; }
		return QMdiSubWindow::redirected(offset);
	}
	QPainter* sharedPainter() const override {
		int __h = 0;
		if (__goOn[41] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 5792771328741774991LL, &__h); if (__h) return (QPainter*)__r; }
		return QMdiSubWindow::sharedPainter();
	}
	void inputMethodEvent(QInputMethodEvent* arg1) override {
		int __h = 0;
		if (__goOn[42] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5438819661942843222LL, (void*)arg1, &__h); if (__h) return; }
		QMdiSubWindow::inputMethodEvent(arg1);
	}
	QVariant inputMethodQuery(Qt::InputMethodQuery arg1) const override {
		int __h = 0;
		if (__goOn[43] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7345853097340709638LL, (int)(arg1), &__h); if (__h) return *(QVariant*)__r; }
		return QMdiSubWindow::inputMethodQuery(arg1);
	}
	bool focusNextPrevChild(bool next) override {
		int __h = 0;
		if (__goOn[44] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 5898328517505784191LL, next, &__h); if (__h) return __r; }
		return QMdiSubWindow::focusNextPrevChild(next);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[45] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QMdiSubWindow::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[46] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QMdiSubWindow::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[47] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QMdiSubWindow::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QMdiSubWindow;
static std::mutex __connsMu_QMdiSubWindow;

extern "C" {

void* QMdiSubWindow_base(void* self) { return (void*)static_cast<QWidget*>((QMdiSubWindow*)self); }
void QMdiSubWindow_chain(void* self, void** out) {
	QMdiSubWindow* __p = (QMdiSubWindow*)self;
	QWidget* __b0 = static_cast<QWidget*>(__p); out[0] = (void*)__b0;
	QObject* __b1 = static_cast<QObject*>(__b0); out[1] = (void*)__b1;
}
void QMdiSubWindow__set_override(void* self, int idx) { ((QMdiSubWindow__Go*)self)->__goOn[idx] = 1; }
void* QMdiSubWindow_fromQObject(void* o) { return (void*)qobject_cast<QMdiSubWindow*>((QObject*)o); }
void* QMdiSubWindow_new(long long __goid, void* parent) {
	return new QMdiSubWindow__Go(__goid, (QWidget*)parent);
}
void* QMdiSubWindow_new2(long long __goid) {
	return new QMdiSubWindow__Go(__goid);
}
void* QMdiSubWindow_new3(long long __goid, void* parent, int flags) {
	return new QMdiSubWindow__Go(__goid, (QWidget*)parent, QFlags<Qt::WindowType>(flags));
}
char* QMdiSubWindow_tr(const char* s, long long* __out_len) {
	auto&& __r = QMdiSubWindow::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QMdiSubWindow_setWidget(void* self, void* widget) {
	((QMdiSubWindow*)self)->setWidget((QWidget*)widget);
}
void* QMdiSubWindow_widget(void* self) {
	auto&& __r = ((QMdiSubWindow*)self)->widget();
	return (void*)(__r);
}
void* QMdiSubWindow_maximizedButtonsWidget(void* self) {
	auto&& __r = ((QMdiSubWindow*)self)->maximizedButtonsWidget();
	return (void*)(__r);
}
void* QMdiSubWindow_maximizedSystemMenuIconWidget(void* self) {
	auto&& __r = ((QMdiSubWindow*)self)->maximizedSystemMenuIconWidget();
	return (void*)(__r);
}
int QMdiSubWindow_isShaded(void* self) {
	auto&& __r = ((QMdiSubWindow*)self)->isShaded();
	return __r;
}
void QMdiSubWindow_setOption(void* self, int option) {
	((QMdiSubWindow*)self)->setOption((QMdiSubWindow::SubWindowOption)(option));
}
int QMdiSubWindow_testOption(void* self, int arg1) {
	auto&& __r = ((QMdiSubWindow*)self)->testOption((QMdiSubWindow::SubWindowOption)(arg1));
	return __r;
}
void QMdiSubWindow_setKeyboardSingleStep(void* self, int step) {
	((QMdiSubWindow*)self)->setKeyboardSingleStep(step);
}
int QMdiSubWindow_keyboardSingleStep(void* self) {
	auto&& __r = ((QMdiSubWindow*)self)->keyboardSingleStep();
	return __r;
}
void QMdiSubWindow_setKeyboardPageStep(void* self, int step) {
	((QMdiSubWindow*)self)->setKeyboardPageStep(step);
}
int QMdiSubWindow_keyboardPageStep(void* self) {
	auto&& __r = ((QMdiSubWindow*)self)->keyboardPageStep();
	return __r;
}
void QMdiSubWindow_setSystemMenu(void* self, void* systemMenu) {
	((QMdiSubWindow*)self)->setSystemMenu((QMenu*)systemMenu);
}
void* QMdiSubWindow_systemMenu(void* self) {
	auto&& __r = ((QMdiSubWindow*)self)->systemMenu();
	return (void*)(__r);
}
void* QMdiSubWindow_mdiArea(void* self) {
	auto&& __r = ((QMdiSubWindow*)self)->mdiArea();
	return (void*)(__r);
}
void QMdiSubWindow_windowStateChanged(void* self, int oldState, int newState) {
	((QMdiSubWindow*)self)->windowStateChanged(QFlags<Qt::WindowState>(oldState), QFlags<Qt::WindowState>(newState));
}
void QMdiSubWindow_aboutToActivate(void* self) {
	((QMdiSubWindow*)self)->aboutToActivate();
}
void QMdiSubWindow_showSystemMenu(void* self) {
	((QMdiSubWindow*)self)->showSystemMenu();
}
void QMdiSubWindow_showShaded(void* self) {
	((QMdiSubWindow*)self)->showShaded();
}
void QMdiSubWindow_setTabOrder(void* arg1, void* arg2) {
	QMdiSubWindow::setTabOrder((QWidget*)arg1, (QWidget*)arg2);
}
void* QMdiSubWindow_mouseGrabber() {
	auto&& __r = QMdiSubWindow::mouseGrabber();
	return (void*)(__r);
}
void* QMdiSubWindow_keyboardGrabber() {
	auto&& __r = QMdiSubWindow::keyboardGrabber();
	return (void*)(__r);
}
void* QMdiSubWindow_find(unsigned long long arg1) {
	auto&& __r = QMdiSubWindow::find(arg1);
	return (void*)(__r);
}
void* QMdiSubWindow_createWindowContainer(void* window) {
	auto&& __r = QMdiSubWindow::createWindowContainer((QWindow*)window);
	return (void*)(__r);
}
void QMdiSubWindow_updateMicroFocus(void* self) {
	((QMdiSubWindow__Go*)self)->updateMicroFocus();
}
void QMdiSubWindow_create(void* self) {
	((QMdiSubWindow__Go*)self)->create();
}
void QMdiSubWindow_destroy(void* self) {
	((QMdiSubWindow__Go*)self)->destroy();
}
int QMdiSubWindow_focusNextChild(void* self) {
	auto&& __r = ((QMdiSubWindow__Go*)self)->focusNextChild();
	return __r;
}
int QMdiSubWindow_focusPreviousChild(void* self) {
	auto&& __r = ((QMdiSubWindow__Go*)self)->focusPreviousChild();
	return __r;
}
void* QMdiSubWindow_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QMdiSubWindow::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QMdiSubWindow_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QMdiSubWindow::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QMdiSubWindow_sender(void* self) {
	auto&& __r = ((QMdiSubWindow__Go*)self)->sender();
	return (void*)(__r);
}
int QMdiSubWindow_senderSignalIndex(void* self) {
	auto&& __r = ((QMdiSubWindow__Go*)self)->senderSignalIndex();
	return __r;
}
int QMdiSubWindow_receivers(void* self, const char* signal) {
	auto&& __r = ((QMdiSubWindow__Go*)self)->receivers(signal);
	return __r;
}
int QMdiSubWindow_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QMdiSubWindow__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
double QMdiSubWindow_devicePixelRatioFScale() {
	auto&& __r = QMdiSubWindow::devicePixelRatioFScale();
	return __r;
}
int QMdiSubWindow_encodeMetricF(int metric, double value) {
	auto&& __r = QMdiSubWindow::encodeMetricF((QPaintDevice::PaintDeviceMetric)(metric), value);
	return __r;
}
double QMdiSubWindow_getDecodedMetricF(void* self, int metricA, int metricB) {
	auto&& __r = ((QMdiSubWindow__Go*)self)->getDecodedMetricF((QPaintDevice::PaintDeviceMetric)(metricA), (QPaintDevice::PaintDeviceMetric)(metricB));
	return __r;
}
char* QMdiSubWindow_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QMdiSubWindow::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QMdiSubWindow_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QMdiSubWindow::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QMdiSubWindow_setOption2(void* self, int option, int on) {
	((QMdiSubWindow*)self)->setOption((QMdiSubWindow::SubWindowOption)(option), on);
}
void* QMdiSubWindow_createWindowContainer2(void* window, void* parent) {
	auto&& __r = QMdiSubWindow::createWindowContainer((QWindow*)window, (QWidget*)parent);
	return (void*)(__r);
}
void* QMdiSubWindow_createWindowContainer3(void* window, void* parent, int flags) {
	auto&& __r = QMdiSubWindow::createWindowContainer((QWindow*)window, (QWidget*)parent, QFlags<Qt::WindowType>(flags));
	return (void*)(__r);
}
void QMdiSubWindow_updateMicroFocus2(void* self, int query) {
	((QMdiSubWindow__Go*)self)->updateMicroFocus((Qt::InputMethodQuery)(query));
}
void QMdiSubWindow_create2(void* self, unsigned long long arg1) {
	((QMdiSubWindow__Go*)self)->create(arg1);
}
void QMdiSubWindow_create3(void* self, unsigned long long arg1, int initializeWindow) {
	((QMdiSubWindow__Go*)self)->create(arg1, initializeWindow);
}
void QMdiSubWindow_create4(void* self, unsigned long long arg1, int initializeWindow, int destroyOldWindow) {
	((QMdiSubWindow__Go*)self)->create(arg1, initializeWindow, destroyOldWindow);
}
void QMdiSubWindow_destroy2(void* self, int destroyWindow) {
	((QMdiSubWindow__Go*)self)->destroy(destroyWindow);
}
void QMdiSubWindow_destroy3(void* self, int destroyWindow, int destroySubWindows) {
	((QMdiSubWindow__Go*)self)->destroy(destroyWindow, destroySubWindows);
}
void* QMdiSubWindow_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QMdiSubWindow::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QMdiSubWindow_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QMdiSubWindow::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QMdiSubWindow_metaObject_default(void* self) {
	auto&& __r = ((QMdiSubWindow*)self)->QMdiSubWindow::metaObject();
	return (void*)(__r);
}
void* QMdiSubWindow_sizeHint_default(void* self) {
	auto&& __r = ((QMdiSubWindow*)self)->QMdiSubWindow::sizeHint();
	return new QSize(__r);
}
void* QMdiSubWindow_minimumSizeHint_default(void* self) {
	auto&& __r = ((QMdiSubWindow*)self)->QMdiSubWindow::minimumSizeHint();
	return new QSize(__r);
}
int QMdiSubWindow_eventFilter_default(void* self, void* object, void* event) {
	auto&& __r = ((QMdiSubWindow__Go*)self)->cppgoBase_eventFilter((QObject*)object, (QEvent*)event);
	return __r;
}
int QMdiSubWindow_event_default(void* self, void* event) {
	auto&& __r = ((QMdiSubWindow__Go*)self)->cppgoBase_event((QEvent*)event);
	return __r;
}
void QMdiSubWindow_showEvent_default(void* self, void* showEvent) {
	((QMdiSubWindow__Go*)self)->cppgoBase_showEvent((QShowEvent*)showEvent);
}
void QMdiSubWindow_hideEvent_default(void* self, void* hideEvent) {
	((QMdiSubWindow__Go*)self)->cppgoBase_hideEvent((QHideEvent*)hideEvent);
}
void QMdiSubWindow_changeEvent_default(void* self, void* changeEvent) {
	((QMdiSubWindow__Go*)self)->cppgoBase_changeEvent((QEvent*)changeEvent);
}
void QMdiSubWindow_closeEvent_default(void* self, void* closeEvent) {
	((QMdiSubWindow__Go*)self)->cppgoBase_closeEvent((QCloseEvent*)closeEvent);
}
void QMdiSubWindow_leaveEvent_default(void* self, void* leaveEvent) {
	((QMdiSubWindow__Go*)self)->cppgoBase_leaveEvent((QEvent*)leaveEvent);
}
void QMdiSubWindow_resizeEvent_default(void* self, void* resizeEvent) {
	((QMdiSubWindow__Go*)self)->cppgoBase_resizeEvent((QResizeEvent*)resizeEvent);
}
void QMdiSubWindow_timerEvent_default(void* self, void* timerEvent) {
	((QMdiSubWindow__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)timerEvent);
}
void QMdiSubWindow_moveEvent_default(void* self, void* moveEvent) {
	((QMdiSubWindow__Go*)self)->cppgoBase_moveEvent((QMoveEvent*)moveEvent);
}
void QMdiSubWindow_paintEvent_default(void* self, void* paintEvent) {
	((QMdiSubWindow__Go*)self)->cppgoBase_paintEvent((QPaintEvent*)paintEvent);
}
void QMdiSubWindow_mousePressEvent_default(void* self, void* mouseEvent) {
	((QMdiSubWindow__Go*)self)->cppgoBase_mousePressEvent((QMouseEvent*)mouseEvent);
}
void QMdiSubWindow_mouseDoubleClickEvent_default(void* self, void* mouseEvent) {
	((QMdiSubWindow__Go*)self)->cppgoBase_mouseDoubleClickEvent((QMouseEvent*)mouseEvent);
}
void QMdiSubWindow_mouseReleaseEvent_default(void* self, void* mouseEvent) {
	((QMdiSubWindow__Go*)self)->cppgoBase_mouseReleaseEvent((QMouseEvent*)mouseEvent);
}
void QMdiSubWindow_mouseMoveEvent_default(void* self, void* mouseEvent) {
	((QMdiSubWindow__Go*)self)->cppgoBase_mouseMoveEvent((QMouseEvent*)mouseEvent);
}
void QMdiSubWindow_keyPressEvent_default(void* self, void* keyEvent) {
	((QMdiSubWindow__Go*)self)->cppgoBase_keyPressEvent((QKeyEvent*)keyEvent);
}
void QMdiSubWindow_contextMenuEvent_default(void* self, void* contextMenuEvent) {
	((QMdiSubWindow__Go*)self)->cppgoBase_contextMenuEvent((QContextMenuEvent*)contextMenuEvent);
}
void QMdiSubWindow_focusInEvent_default(void* self, void* focusInEvent) {
	((QMdiSubWindow__Go*)self)->cppgoBase_focusInEvent((QFocusEvent*)focusInEvent);
}
void QMdiSubWindow_focusOutEvent_default(void* self, void* focusOutEvent) {
	((QMdiSubWindow__Go*)self)->cppgoBase_focusOutEvent((QFocusEvent*)focusOutEvent);
}
void QMdiSubWindow_childEvent_default(void* self, void* childEvent) {
	((QMdiSubWindow__Go*)self)->cppgoBase_childEvent((QChildEvent*)childEvent);
}
int QMdiSubWindow_devType_default(void* self) {
	auto&& __r = ((QMdiSubWindow*)self)->QMdiSubWindow::devType();
	return __r;
}
void QMdiSubWindow_setVisible_default(void* self, int visible) {
	((QMdiSubWindow*)self)->QMdiSubWindow::setVisible(visible);
}
int QMdiSubWindow_heightForWidth_default(void* self, int arg1) {
	auto&& __r = ((QMdiSubWindow*)self)->QMdiSubWindow::heightForWidth(arg1);
	return __r;
}
int QMdiSubWindow_hasHeightForWidth_default(void* self) {
	auto&& __r = ((QMdiSubWindow*)self)->QMdiSubWindow::hasHeightForWidth();
	return __r;
}
void* QMdiSubWindow_paintEngine_default(void* self) {
	auto&& __r = ((QMdiSubWindow*)self)->QMdiSubWindow::paintEngine();
	return (void*)(__r);
}
void QMdiSubWindow_wheelEvent_default(void* self, void* event) {
	((QMdiSubWindow__Go*)self)->cppgoBase_wheelEvent((QWheelEvent*)event);
}
void QMdiSubWindow_keyReleaseEvent_default(void* self, void* event) {
	((QMdiSubWindow__Go*)self)->cppgoBase_keyReleaseEvent((QKeyEvent*)event);
}
void QMdiSubWindow_enterEvent_default(void* self, void* event) {
	((QMdiSubWindow__Go*)self)->cppgoBase_enterEvent((QEnterEvent*)event);
}
void QMdiSubWindow_tabletEvent_default(void* self, void* event) {
	((QMdiSubWindow__Go*)self)->cppgoBase_tabletEvent((QTabletEvent*)event);
}
void QMdiSubWindow_actionEvent_default(void* self, void* event) {
	((QMdiSubWindow__Go*)self)->cppgoBase_actionEvent((QActionEvent*)event);
}
void QMdiSubWindow_dragEnterEvent_default(void* self, void* event) {
	((QMdiSubWindow__Go*)self)->cppgoBase_dragEnterEvent((QDragEnterEvent*)event);
}
void QMdiSubWindow_dragMoveEvent_default(void* self, void* event) {
	((QMdiSubWindow__Go*)self)->cppgoBase_dragMoveEvent((QDragMoveEvent*)event);
}
void QMdiSubWindow_dragLeaveEvent_default(void* self, void* event) {
	((QMdiSubWindow__Go*)self)->cppgoBase_dragLeaveEvent((QDragLeaveEvent*)event);
}
void QMdiSubWindow_dropEvent_default(void* self, void* event) {
	((QMdiSubWindow__Go*)self)->cppgoBase_dropEvent((QDropEvent*)event);
}
int QMdiSubWindow_metric_default(void* self, int arg1) {
	auto&& __r = ((QMdiSubWindow__Go*)self)->cppgoBase_metric(arg1);
	return __r;
}
void QMdiSubWindow_initPainter_default(void* self, void* painter) {
	((QMdiSubWindow__Go*)self)->cppgoBase_initPainter((QPainter*)painter);
}
void* QMdiSubWindow_redirected_default(void* self, void* offset) {
	auto&& __r = ((QMdiSubWindow__Go*)self)->cppgoBase_redirected((QPoint*)offset);
	return (void*)(__r);
}
void* QMdiSubWindow_sharedPainter_default(void* self) {
	auto&& __r = ((QMdiSubWindow__Go*)self)->cppgoBase_sharedPainter();
	return (void*)(__r);
}
void QMdiSubWindow_inputMethodEvent_default(void* self, void* arg1) {
	((QMdiSubWindow__Go*)self)->cppgoBase_inputMethodEvent((QInputMethodEvent*)arg1);
}
void* QMdiSubWindow_inputMethodQuery_default(void* self, int arg1) {
	auto&& __r = ((QMdiSubWindow*)self)->QMdiSubWindow::inputMethodQuery((Qt::InputMethodQuery)(arg1));
	return new QVariant(__r);
}
int QMdiSubWindow_focusNextPrevChild_default(void* self, int next) {
	auto&& __r = ((QMdiSubWindow__Go*)self)->cppgoBase_focusNextPrevChild(next);
	return __r;
}
void QMdiSubWindow_customEvent_default(void* self, void* event) {
	((QMdiSubWindow__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QMdiSubWindow_connectNotify_default(void* self, void* signal) {
	((QMdiSubWindow__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QMdiSubWindow_disconnectNotify_default(void* self, void* signal) {
	((QMdiSubWindow__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QMdiSubWindow_connect_windowStateChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QMdiSubWindow*)self, &QMdiSubWindow::windowStateChanged, (QMdiSubWindow*)self, [__id](QFlags<Qt::WindowState> oldState, QFlags<Qt::WindowState> newState){
		if (cppgo_sig_void_int_int) ((void(*)(long long, int, int))cppgo_sig_void_int_int)(__id, (int)(oldState), (int)(newState));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QMdiSubWindow);
	__conns_QMdiSubWindow[__id] = __c;
}
void QMdiSubWindow_connect_aboutToActivate(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QMdiSubWindow*)self, &QMdiSubWindow::aboutToActivate, (QMdiSubWindow*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QMdiSubWindow);
	__conns_QMdiSubWindow[__id] = __c;
}
void QMdiSubWindow_connect_windowTitleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QMdiSubWindow*)self, &QMdiSubWindow::windowTitleChanged, (QMdiSubWindow*)self, [__id](const QString& title){
		QByteArray title_ba = title.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)title_ba.constData(), title_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QMdiSubWindow);
	__conns_QMdiSubWindow[__id] = __c;
}
void QMdiSubWindow_connect_windowIconChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QMdiSubWindow*)self, &QMdiSubWindow::windowIconChanged, (QMdiSubWindow*)self, [__id](const QIcon& icon){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&icon);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QMdiSubWindow);
	__conns_QMdiSubWindow[__id] = __c;
}
void QMdiSubWindow_connect_windowIconTextChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QMdiSubWindow*)self, &QMdiSubWindow::windowIconTextChanged, (QMdiSubWindow*)self, [__id](const QString& iconText){
		QByteArray iconText_ba = iconText.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)iconText_ba.constData(), iconText_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QMdiSubWindow);
	__conns_QMdiSubWindow[__id] = __c;
}
void QMdiSubWindow_connect_customContextMenuRequested(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QMdiSubWindow*)self, &QMdiSubWindow::customContextMenuRequested, (QMdiSubWindow*)self, [__id](const QPoint& pos){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&pos);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QMdiSubWindow);
	__conns_QMdiSubWindow[__id] = __c;
}
void QMdiSubWindow_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QMdiSubWindow*)self, &QMdiSubWindow::destroyed, (QMdiSubWindow*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QMdiSubWindow);
	__conns_QMdiSubWindow[__id] = __c;
}
void QMdiSubWindow_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QMdiSubWindow*)self, &QMdiSubWindow::destroyed, (QMdiSubWindow*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QMdiSubWindow);
	__conns_QMdiSubWindow[__id] = __c;
}
void QMdiSubWindow_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QMdiSubWindow*)self, &QMdiSubWindow::objectNameChanged, (QMdiSubWindow*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QMdiSubWindow);
	__conns_QMdiSubWindow[__id] = __c;
}
void QMdiSubWindow__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QMdiSubWindow);
		auto __it = __conns_QMdiSubWindow.find(id);
		if (__it == __conns_QMdiSubWindow.end()) return;
		__c = __it->second;
		__conns_QMdiSubWindow.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QMdiSubWindow__Go_new(long long id) { return new QMdiSubWindow__Go(id); }
void QMdiSubWindow_delete(void* self) { delete (QMdiSubWindow*)self; }

}

#include <QtWidgets/qstatusbar.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtGui/qevent.h>
#include <QtCore/qcoreevent.h>
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
#include "gen_qstatusbar.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QStatusBar__Go : public QStatusBar {
public:
	long long __goId;
	unsigned char __goOn[48] = {};
	template<typename... A> QStatusBar__Go(long long id, A&&... a) : QStatusBar(std::forward<A>(a)...), __goId(id) {}
	using QStatusBar::reformat;
	using QStatusBar::hideOrShow;
	using QStatusBar::updateMicroFocus;
	using QStatusBar::create;
	using QStatusBar::destroy;
	using QStatusBar::focusNextChild;
	using QStatusBar::focusPreviousChild;
	using QStatusBar::sender;
	using QStatusBar::senderSignalIndex;
	using QStatusBar::receivers;
	using QStatusBar::isSignalConnected;
	using QStatusBar::getDecodedMetricF;
	void cppgoBase_showEvent(QShowEvent* __a0) { QStatusBar::showEvent(__a0); }
	void cppgoBase_paintEvent(QPaintEvent* __a0) { QStatusBar::paintEvent(__a0); }
	void cppgoBase_resizeEvent(QResizeEvent* __a0) { QStatusBar::resizeEvent(__a0); }
	bool cppgoBase_event(QEvent* __a0) { return QStatusBar::event(__a0); }
	void cppgoBase_mousePressEvent(QMouseEvent* __a0) { QStatusBar::mousePressEvent(__a0); }
	void cppgoBase_mouseReleaseEvent(QMouseEvent* __a0) { QStatusBar::mouseReleaseEvent(__a0); }
	void cppgoBase_mouseDoubleClickEvent(QMouseEvent* __a0) { QStatusBar::mouseDoubleClickEvent(__a0); }
	void cppgoBase_mouseMoveEvent(QMouseEvent* __a0) { QStatusBar::mouseMoveEvent(__a0); }
	void cppgoBase_wheelEvent(QWheelEvent* __a0) { QStatusBar::wheelEvent(__a0); }
	void cppgoBase_keyPressEvent(QKeyEvent* __a0) { QStatusBar::keyPressEvent(__a0); }
	void cppgoBase_keyReleaseEvent(QKeyEvent* __a0) { QStatusBar::keyReleaseEvent(__a0); }
	void cppgoBase_focusInEvent(QFocusEvent* __a0) { QStatusBar::focusInEvent(__a0); }
	void cppgoBase_focusOutEvent(QFocusEvent* __a0) { QStatusBar::focusOutEvent(__a0); }
	void cppgoBase_enterEvent(QEnterEvent* __a0) { QStatusBar::enterEvent(__a0); }
	void cppgoBase_leaveEvent(QEvent* __a0) { QStatusBar::leaveEvent(__a0); }
	void cppgoBase_moveEvent(QMoveEvent* __a0) { QStatusBar::moveEvent(__a0); }
	void cppgoBase_closeEvent(QCloseEvent* __a0) { QStatusBar::closeEvent(__a0); }
	void cppgoBase_contextMenuEvent(QContextMenuEvent* __a0) { QStatusBar::contextMenuEvent(__a0); }
	void cppgoBase_tabletEvent(QTabletEvent* __a0) { QStatusBar::tabletEvent(__a0); }
	void cppgoBase_actionEvent(QActionEvent* __a0) { QStatusBar::actionEvent(__a0); }
	void cppgoBase_dragEnterEvent(QDragEnterEvent* __a0) { QStatusBar::dragEnterEvent(__a0); }
	void cppgoBase_dragMoveEvent(QDragMoveEvent* __a0) { QStatusBar::dragMoveEvent(__a0); }
	void cppgoBase_dragLeaveEvent(QDragLeaveEvent* __a0) { QStatusBar::dragLeaveEvent(__a0); }
	void cppgoBase_dropEvent(QDropEvent* __a0) { QStatusBar::dropEvent(__a0); }
	void cppgoBase_hideEvent(QHideEvent* __a0) { QStatusBar::hideEvent(__a0); }
	void cppgoBase_changeEvent(QEvent* __a0) { QStatusBar::changeEvent(__a0); }
	int cppgoBase_metric(int __a0) { return QStatusBar::metric((QPaintDevice::PaintDeviceMetric)__a0); }
	void cppgoBase_initPainter(QPainter* __a0) { QStatusBar::initPainter(__a0); }
	QPaintDevice* cppgoBase_redirected(QPoint* __a0) { return QStatusBar::redirected(__a0); }
	QPainter* cppgoBase_sharedPainter() { return QStatusBar::sharedPainter(); }
	void cppgoBase_inputMethodEvent(QInputMethodEvent* __a0) { QStatusBar::inputMethodEvent(__a0); }
	bool cppgoBase_focusNextPrevChild(bool __a0) { return QStatusBar::focusNextPrevChild(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QStatusBar::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QStatusBar::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QStatusBar::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QStatusBar::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QStatusBar::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QStatusBar::metaObject();
	}
	void showEvent(QShowEvent* arg1) override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6825369174751916LL, (void*)arg1, &__h); if (__h) return; }
		QStatusBar::showEvent(arg1);
	}
	void paintEvent(QPaintEvent* arg1) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5290442163167458465LL, (void*)arg1, &__h); if (__h) return; }
		QStatusBar::paintEvent(arg1);
	}
	void resizeEvent(QResizeEvent* arg1) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4302917285655014661LL, (void*)arg1, &__h); if (__h) return; }
		QStatusBar::resizeEvent(arg1);
	}
	bool event(QEvent* arg1) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)arg1, &__h); if (__h) return __r; }
		return QStatusBar::event(arg1);
	}
	int devType() const override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8659679893868524044LL, &__h); if (__h) return __r; }
		return QStatusBar::devType();
	}
	void setVisible(bool visible) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 2110182719375245575LL, visible, &__h); if (__h) return; }
		QStatusBar::setVisible(visible);
	}
	QSize sizeHint() const override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 729690261457684101LL, &__h); if (__h) return *(QSize*)__r; }
		return QStatusBar::sizeHint();
	}
	QSize minimumSizeHint() const override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 7049264599837532491LL, &__h); if (__h) return *(QSize*)__r; }
		return QStatusBar::minimumSizeHint();
	}
	int heightForWidth(int arg1) const override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 4311534060495737339LL, arg1, &__h); if (__h) return __r; }
		return QStatusBar::heightForWidth(arg1);
	}
	bool hasHeightForWidth() const override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8266762397995423915LL, &__h); if (__h) return __r; }
		return QStatusBar::hasHeightForWidth();
	}
	QPaintEngine* paintEngine() const override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 6398007404756346381LL, &__h); if (__h) return (QPaintEngine*)__r; }
		return QStatusBar::paintEngine();
	}
	void mousePressEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -332090898078543319LL, (void*)event, &__h); if (__h) return; }
		QStatusBar::mousePressEvent(event);
	}
	void mouseReleaseEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4415920049297795213LL, (void*)event, &__h); if (__h) return; }
		QStatusBar::mouseReleaseEvent(event);
	}
	void mouseDoubleClickEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4029073905697023887LL, (void*)event, &__h); if (__h) return; }
		QStatusBar::mouseDoubleClickEvent(event);
	}
	void mouseMoveEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4255494004099806581LL, (void*)event, &__h); if (__h) return; }
		QStatusBar::mouseMoveEvent(event);
	}
	void wheelEvent(QWheelEvent* event) override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1919988340934511316LL, (void*)event, &__h); if (__h) return; }
		QStatusBar::wheelEvent(event);
	}
	void keyPressEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1023791977679793185LL, (void*)event, &__h); if (__h) return; }
		QStatusBar::keyPressEvent(event);
	}
	void keyReleaseEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 772598475950050299LL, (void*)event, &__h); if (__h) return; }
		QStatusBar::keyReleaseEvent(event);
	}
	void focusInEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 975043767033552552LL, (void*)event, &__h); if (__h) return; }
		QStatusBar::focusInEvent(event);
	}
	void focusOutEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -385596010576048993LL, (void*)event, &__h); if (__h) return; }
		QStatusBar::focusOutEvent(event);
	}
	void enterEvent(QEnterEvent* event) override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6316270346668237895LL, (void*)event, &__h); if (__h) return; }
		QStatusBar::enterEvent(event);
	}
	void leaveEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -551070009623711640LL, (void*)event, &__h); if (__h) return; }
		QStatusBar::leaveEvent(event);
	}
	void moveEvent(QMoveEvent* event) override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -9161949150489525396LL, (void*)event, &__h); if (__h) return; }
		QStatusBar::moveEvent(event);
	}
	void closeEvent(QCloseEvent* event) override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1989937050019464783LL, (void*)event, &__h); if (__h) return; }
		QStatusBar::closeEvent(event);
	}
	void contextMenuEvent(QContextMenuEvent* event) override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6698043222227006015LL, (void*)event, &__h); if (__h) return; }
		QStatusBar::contextMenuEvent(event);
	}
	void tabletEvent(QTabletEvent* event) override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -212068866860701141LL, (void*)event, &__h); if (__h) return; }
		QStatusBar::tabletEvent(event);
	}
	void actionEvent(QActionEvent* event) override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1697444140088555581LL, (void*)event, &__h); if (__h) return; }
		QStatusBar::actionEvent(event);
	}
	void dragEnterEvent(QDragEnterEvent* event) override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 206122715919917589LL, (void*)event, &__h); if (__h) return; }
		QStatusBar::dragEnterEvent(event);
	}
	void dragMoveEvent(QDragMoveEvent* event) override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1781545137352745978LL, (void*)event, &__h); if (__h) return; }
		QStatusBar::dragMoveEvent(event);
	}
	void dragLeaveEvent(QDragLeaveEvent* event) override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -3139904158373725006LL, (void*)event, &__h); if (__h) return; }
		QStatusBar::dragLeaveEvent(event);
	}
	void dropEvent(QDropEvent* event) override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 9107058682754591716LL, (void*)event, &__h); if (__h) return; }
		QStatusBar::dropEvent(event);
	}
	void hideEvent(QHideEvent* event) override {
		int __h = 0;
		if (__goOn[32] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -199718707876719129LL, (void*)event, &__h); if (__h) return; }
		QStatusBar::hideEvent(event);
	}
	bool nativeEvent(const QByteArray& eventType, void* message, qintptr* result) override {
		int __h = 0;
		if (__goOn[33] && cppgo_virt_int_voidp_voidp_longlongp) { int __r = ((int(*)(long long, long long, void*, void*, long long*, int*))cppgo_virt_int_voidp_voidp_longlongp)(__goId, -2039166586246443706LL, (void*)&eventType, (void*)message, (long long*)result, &__h); if (__h) return __r; }
		return QStatusBar::nativeEvent(eventType, message, result);
	}
	void changeEvent(QEvent* arg1) override {
		int __h = 0;
		if (__goOn[34] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4959376063826165943LL, (void*)arg1, &__h); if (__h) return; }
		QStatusBar::changeEvent(arg1);
	}
	int metric(QPaintDevice::PaintDeviceMetric arg1) const override {
		int __h = 0;
		if (__goOn[35] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 427629649267311777LL, (int)(arg1), &__h); if (__h) return __r; }
		return QStatusBar::metric(arg1);
	}
	void initPainter(QPainter* painter) const override {
		int __h = 0;
		if (__goOn[36] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6477500296980050210LL, (void*)painter, &__h); if (__h) return; }
		QStatusBar::initPainter(painter);
	}
	QPaintDevice* redirected(QPoint* offset) const override {
		int __h = 0;
		if (__goOn[37] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 4127240633310655048LL, (void*)offset, &__h); if (__h) return (QPaintDevice*)__r; }
		return QStatusBar::redirected(offset);
	}
	QPainter* sharedPainter() const override {
		int __h = 0;
		if (__goOn[38] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 5792771328741774991LL, &__h); if (__h) return (QPainter*)__r; }
		return QStatusBar::sharedPainter();
	}
	void inputMethodEvent(QInputMethodEvent* arg1) override {
		int __h = 0;
		if (__goOn[39] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5438819661942843222LL, (void*)arg1, &__h); if (__h) return; }
		QStatusBar::inputMethodEvent(arg1);
	}
	QVariant inputMethodQuery(Qt::InputMethodQuery arg1) const override {
		int __h = 0;
		if (__goOn[40] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7345853097340709638LL, (int)(arg1), &__h); if (__h) return *(QVariant*)__r; }
		return QStatusBar::inputMethodQuery(arg1);
	}
	bool focusNextPrevChild(bool next) override {
		int __h = 0;
		if (__goOn[41] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 5898328517505784191LL, next, &__h); if (__h) return __r; }
		return QStatusBar::focusNextPrevChild(next);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[42] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QStatusBar::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[43] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QStatusBar::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[44] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QStatusBar::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[45] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QStatusBar::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[46] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QStatusBar::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[47] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QStatusBar::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QStatusBar;
static std::mutex __connsMu_QStatusBar;

extern "C" {

void* QStatusBar_base(void* self) { return (void*)static_cast<QWidget*>((QStatusBar*)self); }
void QStatusBar_chain(void* self, void** out) {
	QStatusBar* __p = (QStatusBar*)self;
	QWidget* __b0 = static_cast<QWidget*>(__p); out[0] = (void*)__b0;
	QObject* __b1 = static_cast<QObject*>(__b0); out[1] = (void*)__b1;
}
void QStatusBar__set_override(void* self, int idx) { ((QStatusBar__Go*)self)->__goOn[idx] = 1; }
void* QStatusBar_fromQObject(void* o) { return (void*)qobject_cast<QStatusBar*>((QObject*)o); }
void* QStatusBar_new(long long __goid, void* parent) {
	return new QStatusBar__Go(__goid, (QWidget*)parent);
}
void* QStatusBar_new2(long long __goid) {
	return new QStatusBar__Go(__goid);
}
char* QStatusBar_tr(const char* s, long long* __out_len) {
	auto&& __r = QStatusBar::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QStatusBar_addWidget(void* self, void* widget) {
	((QStatusBar*)self)->addWidget((QWidget*)widget);
}
int QStatusBar_insertWidget(void* self, int index, void* widget) {
	auto&& __r = ((QStatusBar*)self)->insertWidget(index, (QWidget*)widget);
	return __r;
}
void QStatusBar_addPermanentWidget(void* self, void* widget) {
	((QStatusBar*)self)->addPermanentWidget((QWidget*)widget);
}
int QStatusBar_insertPermanentWidget(void* self, int index, void* widget) {
	auto&& __r = ((QStatusBar*)self)->insertPermanentWidget(index, (QWidget*)widget);
	return __r;
}
void QStatusBar_removeWidget(void* self, void* widget) {
	((QStatusBar*)self)->removeWidget((QWidget*)widget);
}
void QStatusBar_setSizeGripEnabled(void* self, int arg1) {
	((QStatusBar*)self)->setSizeGripEnabled(arg1);
}
int QStatusBar_isSizeGripEnabled(void* self) {
	auto&& __r = ((QStatusBar*)self)->isSizeGripEnabled();
	return __r;
}
char* QStatusBar_currentMessage(void* self, long long* __out_len) {
	auto&& __r = ((QStatusBar*)self)->currentMessage();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QStatusBar_showMessage(void* self, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	((QStatusBar*)self)->showMessage(text_qstr);
}
void QStatusBar_clearMessage(void* self) {
	((QStatusBar*)self)->clearMessage();
}
void QStatusBar_messageChanged(void* self, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	((QStatusBar*)self)->messageChanged(text_qstr);
}
void QStatusBar_reformat(void* self) {
	((QStatusBar__Go*)self)->reformat();
}
void QStatusBar_hideOrShow(void* self) {
	((QStatusBar__Go*)self)->hideOrShow();
}
void QStatusBar_setTabOrder(void* arg1, void* arg2) {
	QStatusBar::setTabOrder((QWidget*)arg1, (QWidget*)arg2);
}
void* QStatusBar_mouseGrabber() {
	auto&& __r = QStatusBar::mouseGrabber();
	return (void*)(__r);
}
void* QStatusBar_keyboardGrabber() {
	auto&& __r = QStatusBar::keyboardGrabber();
	return (void*)(__r);
}
void* QStatusBar_find(unsigned long long arg1) {
	auto&& __r = QStatusBar::find(arg1);
	return (void*)(__r);
}
void* QStatusBar_createWindowContainer(void* window) {
	auto&& __r = QStatusBar::createWindowContainer((QWindow*)window);
	return (void*)(__r);
}
void QStatusBar_updateMicroFocus(void* self) {
	((QStatusBar__Go*)self)->updateMicroFocus();
}
void QStatusBar_create(void* self) {
	((QStatusBar__Go*)self)->create();
}
void QStatusBar_destroy(void* self) {
	((QStatusBar__Go*)self)->destroy();
}
int QStatusBar_focusNextChild(void* self) {
	auto&& __r = ((QStatusBar__Go*)self)->focusNextChild();
	return __r;
}
int QStatusBar_focusPreviousChild(void* self) {
	auto&& __r = ((QStatusBar__Go*)self)->focusPreviousChild();
	return __r;
}
void* QStatusBar_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QStatusBar::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QStatusBar_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QStatusBar::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QStatusBar_sender(void* self) {
	auto&& __r = ((QStatusBar__Go*)self)->sender();
	return (void*)(__r);
}
int QStatusBar_senderSignalIndex(void* self) {
	auto&& __r = ((QStatusBar__Go*)self)->senderSignalIndex();
	return __r;
}
int QStatusBar_receivers(void* self, const char* signal) {
	auto&& __r = ((QStatusBar__Go*)self)->receivers(signal);
	return __r;
}
int QStatusBar_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QStatusBar__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
double QStatusBar_devicePixelRatioFScale() {
	auto&& __r = QStatusBar::devicePixelRatioFScale();
	return __r;
}
int QStatusBar_encodeMetricF(int metric, double value) {
	auto&& __r = QStatusBar::encodeMetricF((QPaintDevice::PaintDeviceMetric)(metric), value);
	return __r;
}
double QStatusBar_getDecodedMetricF(void* self, int metricA, int metricB) {
	auto&& __r = ((QStatusBar__Go*)self)->getDecodedMetricF((QPaintDevice::PaintDeviceMetric)(metricA), (QPaintDevice::PaintDeviceMetric)(metricB));
	return __r;
}
char* QStatusBar_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QStatusBar::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QStatusBar_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QStatusBar::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QStatusBar_addWidget2(void* self, void* widget, int stretch) {
	((QStatusBar*)self)->addWidget((QWidget*)widget, stretch);
}
int QStatusBar_insertWidget2(void* self, int index, void* widget, int stretch) {
	auto&& __r = ((QStatusBar*)self)->insertWidget(index, (QWidget*)widget, stretch);
	return __r;
}
void QStatusBar_addPermanentWidget2(void* self, void* widget, int stretch) {
	((QStatusBar*)self)->addPermanentWidget((QWidget*)widget, stretch);
}
int QStatusBar_insertPermanentWidget2(void* self, int index, void* widget, int stretch) {
	auto&& __r = ((QStatusBar*)self)->insertPermanentWidget(index, (QWidget*)widget, stretch);
	return __r;
}
void QStatusBar_showMessage2(void* self, const char* text, long long text_len, int timeout) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	((QStatusBar*)self)->showMessage(text_qstr, timeout);
}
void* QStatusBar_createWindowContainer2(void* window, void* parent) {
	auto&& __r = QStatusBar::createWindowContainer((QWindow*)window, (QWidget*)parent);
	return (void*)(__r);
}
void* QStatusBar_createWindowContainer3(void* window, void* parent, int flags) {
	auto&& __r = QStatusBar::createWindowContainer((QWindow*)window, (QWidget*)parent, QFlags<Qt::WindowType>(flags));
	return (void*)(__r);
}
void QStatusBar_updateMicroFocus2(void* self, int query) {
	((QStatusBar__Go*)self)->updateMicroFocus((Qt::InputMethodQuery)(query));
}
void QStatusBar_create2(void* self, unsigned long long arg1) {
	((QStatusBar__Go*)self)->create(arg1);
}
void QStatusBar_create3(void* self, unsigned long long arg1, int initializeWindow) {
	((QStatusBar__Go*)self)->create(arg1, initializeWindow);
}
void QStatusBar_create4(void* self, unsigned long long arg1, int initializeWindow, int destroyOldWindow) {
	((QStatusBar__Go*)self)->create(arg1, initializeWindow, destroyOldWindow);
}
void QStatusBar_destroy2(void* self, int destroyWindow) {
	((QStatusBar__Go*)self)->destroy(destroyWindow);
}
void QStatusBar_destroy3(void* self, int destroyWindow, int destroySubWindows) {
	((QStatusBar__Go*)self)->destroy(destroyWindow, destroySubWindows);
}
void* QStatusBar_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QStatusBar::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QStatusBar_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QStatusBar::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QStatusBar_metaObject_default(void* self) {
	auto&& __r = ((QStatusBar*)self)->QStatusBar::metaObject();
	return (void*)(__r);
}
void QStatusBar_showEvent_default(void* self, void* arg1) {
	((QStatusBar__Go*)self)->cppgoBase_showEvent((QShowEvent*)arg1);
}
void QStatusBar_paintEvent_default(void* self, void* arg1) {
	((QStatusBar__Go*)self)->cppgoBase_paintEvent((QPaintEvent*)arg1);
}
void QStatusBar_resizeEvent_default(void* self, void* arg1) {
	((QStatusBar__Go*)self)->cppgoBase_resizeEvent((QResizeEvent*)arg1);
}
int QStatusBar_event_default(void* self, void* arg1) {
	auto&& __r = ((QStatusBar__Go*)self)->cppgoBase_event((QEvent*)arg1);
	return __r;
}
int QStatusBar_devType_default(void* self) {
	auto&& __r = ((QStatusBar*)self)->QStatusBar::devType();
	return __r;
}
void QStatusBar_setVisible_default(void* self, int visible) {
	((QStatusBar*)self)->QStatusBar::setVisible(visible);
}
void* QStatusBar_sizeHint_default(void* self) {
	auto&& __r = ((QStatusBar*)self)->QStatusBar::sizeHint();
	return new QSize(__r);
}
void* QStatusBar_minimumSizeHint_default(void* self) {
	auto&& __r = ((QStatusBar*)self)->QStatusBar::minimumSizeHint();
	return new QSize(__r);
}
int QStatusBar_heightForWidth_default(void* self, int arg1) {
	auto&& __r = ((QStatusBar*)self)->QStatusBar::heightForWidth(arg1);
	return __r;
}
int QStatusBar_hasHeightForWidth_default(void* self) {
	auto&& __r = ((QStatusBar*)self)->QStatusBar::hasHeightForWidth();
	return __r;
}
void* QStatusBar_paintEngine_default(void* self) {
	auto&& __r = ((QStatusBar*)self)->QStatusBar::paintEngine();
	return (void*)(__r);
}
void QStatusBar_mousePressEvent_default(void* self, void* event) {
	((QStatusBar__Go*)self)->cppgoBase_mousePressEvent((QMouseEvent*)event);
}
void QStatusBar_mouseReleaseEvent_default(void* self, void* event) {
	((QStatusBar__Go*)self)->cppgoBase_mouseReleaseEvent((QMouseEvent*)event);
}
void QStatusBar_mouseDoubleClickEvent_default(void* self, void* event) {
	((QStatusBar__Go*)self)->cppgoBase_mouseDoubleClickEvent((QMouseEvent*)event);
}
void QStatusBar_mouseMoveEvent_default(void* self, void* event) {
	((QStatusBar__Go*)self)->cppgoBase_mouseMoveEvent((QMouseEvent*)event);
}
void QStatusBar_wheelEvent_default(void* self, void* event) {
	((QStatusBar__Go*)self)->cppgoBase_wheelEvent((QWheelEvent*)event);
}
void QStatusBar_keyPressEvent_default(void* self, void* event) {
	((QStatusBar__Go*)self)->cppgoBase_keyPressEvent((QKeyEvent*)event);
}
void QStatusBar_keyReleaseEvent_default(void* self, void* event) {
	((QStatusBar__Go*)self)->cppgoBase_keyReleaseEvent((QKeyEvent*)event);
}
void QStatusBar_focusInEvent_default(void* self, void* event) {
	((QStatusBar__Go*)self)->cppgoBase_focusInEvent((QFocusEvent*)event);
}
void QStatusBar_focusOutEvent_default(void* self, void* event) {
	((QStatusBar__Go*)self)->cppgoBase_focusOutEvent((QFocusEvent*)event);
}
void QStatusBar_enterEvent_default(void* self, void* event) {
	((QStatusBar__Go*)self)->cppgoBase_enterEvent((QEnterEvent*)event);
}
void QStatusBar_leaveEvent_default(void* self, void* event) {
	((QStatusBar__Go*)self)->cppgoBase_leaveEvent((QEvent*)event);
}
void QStatusBar_moveEvent_default(void* self, void* event) {
	((QStatusBar__Go*)self)->cppgoBase_moveEvent((QMoveEvent*)event);
}
void QStatusBar_closeEvent_default(void* self, void* event) {
	((QStatusBar__Go*)self)->cppgoBase_closeEvent((QCloseEvent*)event);
}
void QStatusBar_contextMenuEvent_default(void* self, void* event) {
	((QStatusBar__Go*)self)->cppgoBase_contextMenuEvent((QContextMenuEvent*)event);
}
void QStatusBar_tabletEvent_default(void* self, void* event) {
	((QStatusBar__Go*)self)->cppgoBase_tabletEvent((QTabletEvent*)event);
}
void QStatusBar_actionEvent_default(void* self, void* event) {
	((QStatusBar__Go*)self)->cppgoBase_actionEvent((QActionEvent*)event);
}
void QStatusBar_dragEnterEvent_default(void* self, void* event) {
	((QStatusBar__Go*)self)->cppgoBase_dragEnterEvent((QDragEnterEvent*)event);
}
void QStatusBar_dragMoveEvent_default(void* self, void* event) {
	((QStatusBar__Go*)self)->cppgoBase_dragMoveEvent((QDragMoveEvent*)event);
}
void QStatusBar_dragLeaveEvent_default(void* self, void* event) {
	((QStatusBar__Go*)self)->cppgoBase_dragLeaveEvent((QDragLeaveEvent*)event);
}
void QStatusBar_dropEvent_default(void* self, void* event) {
	((QStatusBar__Go*)self)->cppgoBase_dropEvent((QDropEvent*)event);
}
void QStatusBar_hideEvent_default(void* self, void* event) {
	((QStatusBar__Go*)self)->cppgoBase_hideEvent((QHideEvent*)event);
}
void QStatusBar_changeEvent_default(void* self, void* arg1) {
	((QStatusBar__Go*)self)->cppgoBase_changeEvent((QEvent*)arg1);
}
int QStatusBar_metric_default(void* self, int arg1) {
	auto&& __r = ((QStatusBar__Go*)self)->cppgoBase_metric(arg1);
	return __r;
}
void QStatusBar_initPainter_default(void* self, void* painter) {
	((QStatusBar__Go*)self)->cppgoBase_initPainter((QPainter*)painter);
}
void* QStatusBar_redirected_default(void* self, void* offset) {
	auto&& __r = ((QStatusBar__Go*)self)->cppgoBase_redirected((QPoint*)offset);
	return (void*)(__r);
}
void* QStatusBar_sharedPainter_default(void* self) {
	auto&& __r = ((QStatusBar__Go*)self)->cppgoBase_sharedPainter();
	return (void*)(__r);
}
void QStatusBar_inputMethodEvent_default(void* self, void* arg1) {
	((QStatusBar__Go*)self)->cppgoBase_inputMethodEvent((QInputMethodEvent*)arg1);
}
void* QStatusBar_inputMethodQuery_default(void* self, int arg1) {
	auto&& __r = ((QStatusBar*)self)->QStatusBar::inputMethodQuery((Qt::InputMethodQuery)(arg1));
	return new QVariant(__r);
}
int QStatusBar_focusNextPrevChild_default(void* self, int next) {
	auto&& __r = ((QStatusBar__Go*)self)->cppgoBase_focusNextPrevChild(next);
	return __r;
}
int QStatusBar_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QStatusBar*)self)->QStatusBar::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QStatusBar_timerEvent_default(void* self, void* event) {
	((QStatusBar__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QStatusBar_childEvent_default(void* self, void* event) {
	((QStatusBar__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QStatusBar_customEvent_default(void* self, void* event) {
	((QStatusBar__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QStatusBar_connectNotify_default(void* self, void* signal) {
	((QStatusBar__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QStatusBar_disconnectNotify_default(void* self, void* signal) {
	((QStatusBar__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QStatusBar_connect_messageChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStatusBar*)self, &QStatusBar::messageChanged, (QStatusBar*)self, [__id](const QString& text){
		QByteArray text_ba = text.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)text_ba.constData(), text_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStatusBar);
	__conns_QStatusBar[__id] = __c;
}
void QStatusBar_connect_windowTitleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStatusBar*)self, &QStatusBar::windowTitleChanged, (QStatusBar*)self, [__id](const QString& title){
		QByteArray title_ba = title.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)title_ba.constData(), title_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStatusBar);
	__conns_QStatusBar[__id] = __c;
}
void QStatusBar_connect_windowIconChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStatusBar*)self, &QStatusBar::windowIconChanged, (QStatusBar*)self, [__id](const QIcon& icon){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&icon);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStatusBar);
	__conns_QStatusBar[__id] = __c;
}
void QStatusBar_connect_windowIconTextChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStatusBar*)self, &QStatusBar::windowIconTextChanged, (QStatusBar*)self, [__id](const QString& iconText){
		QByteArray iconText_ba = iconText.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)iconText_ba.constData(), iconText_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStatusBar);
	__conns_QStatusBar[__id] = __c;
}
void QStatusBar_connect_customContextMenuRequested(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStatusBar*)self, &QStatusBar::customContextMenuRequested, (QStatusBar*)self, [__id](const QPoint& pos){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&pos);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStatusBar);
	__conns_QStatusBar[__id] = __c;
}
void QStatusBar_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStatusBar*)self, &QStatusBar::destroyed, (QStatusBar*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStatusBar);
	__conns_QStatusBar[__id] = __c;
}
void QStatusBar_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStatusBar*)self, &QStatusBar::destroyed, (QStatusBar*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStatusBar);
	__conns_QStatusBar[__id] = __c;
}
void QStatusBar_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStatusBar*)self, &QStatusBar::objectNameChanged, (QStatusBar*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStatusBar);
	__conns_QStatusBar[__id] = __c;
}
void QStatusBar__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QStatusBar);
		auto __it = __conns_QStatusBar.find(id);
		if (__it == __conns_QStatusBar.end()) return;
		__c = __it->second;
		__conns_QStatusBar.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QStatusBar__Go_new(long long id) { return new QStatusBar__Go(id); }
void QStatusBar_delete(void* self) { delete (QStatusBar*)self; }

}

