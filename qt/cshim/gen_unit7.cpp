#include <QtCore/qcompare.h>
#include "gen_qt__weak_ordering.h"

extern "C" {

void* Qt__weak_ordering_new(void* arg1) {
	return new Qt::weak_ordering(*(Qt::weak_ordering*)arg1);
}
void* Qt__weak_ordering_ToQtPartialOrdering(void* self) {
	auto&& __r = static_cast<Qt::partial_ordering>(*(Qt::weak_ordering*)self);
	return new Qt::partial_ordering(__r);
}
int Qt__weak_ordering_OperatorEqual(void* self, void* rhs) {
	auto&& __r = (*(Qt::weak_ordering*)self) == (*(Qt::weak_ordering*)rhs);
	return __r;
}
int Qt__weak_ordering_OperatorNotEqual(void* self, void* rhs) {
	auto&& __r = (*(Qt::weak_ordering*)self) != (*(Qt::weak_ordering*)rhs);
	return __r;
}
int Qt__weak_ordering_OperatorEqual2(void* self, void* rhs) {
	auto&& __r = (*(Qt::weak_ordering*)self) == (*(Qt::partial_ordering*)rhs);
	return __r;
}
int Qt__weak_ordering_OperatorNotEqual2(void* self, void* rhs) {
	auto&& __r = (*(Qt::weak_ordering*)self) != (*(Qt::partial_ordering*)rhs);
	return __r;
}
int Qt__weak_ordering_OperatorEqual3(void* self, void* rhs) {
	auto&& __r = (*(Qt::weak_ordering*)self) == (*(Qt::strong_ordering*)rhs);
	return __r;
}
int Qt__weak_ordering_OperatorNotEqual3(void* self, void* rhs) {
	auto&& __r = (*(Qt::weak_ordering*)self) != (*(Qt::strong_ordering*)rhs);
	return __r;
}
void Qt__weak_ordering_delete(void* self) { delete (Qt::weak_ordering*)self; }

}

#include <QtCore/qstringconverter_base.h>
#include "gen_qstringconverterbase__state.h"

extern "C" {

void* QStringConverterBase__State_new() {
	return new QStringConverterBase::State();
}
void* QStringConverterBase__State_new2(int f) {
	return new QStringConverterBase::State(QFlags<QStringConverterBase::Flag>(f));
}
void QStringConverterBase__State_clear(void* self) {
	((QStringConverterBase::State*)self)->clear();
}
void QStringConverterBase__State_reset(void* self) {
	((QStringConverterBase::State*)self)->reset();
}
int QStringConverterBase__State_flags_get(void* self) {
	auto __r = ((QStringConverterBase::State*)self)->flags;
	return (int)(__r);
}
void QStringConverterBase__State_flags_set(void* self, int flags) {
	((QStringConverterBase::State*)self)->flags = QFlags<QStringConverterBase::Flag>(flags);
}
int QStringConverterBase__State_internalState_get(void* self) {
	auto __r = ((QStringConverterBase::State*)self)->internalState;
	return __r;
}
void QStringConverterBase__State_internalState_set(void* self, int internalState) {
	((QStringConverterBase::State*)self)->internalState = internalState;
}
long long QStringConverterBase__State_remainingChars_get(void* self) {
	auto __r = ((QStringConverterBase::State*)self)->remainingChars;
	return __r;
}
void QStringConverterBase__State_remainingChars_set(void* self, long long remainingChars) {
	((QStringConverterBase::State*)self)->remainingChars = remainingChars;
}
long long QStringConverterBase__State_invalidChars_get(void* self) {
	auto __r = ((QStringConverterBase::State*)self)->invalidChars;
	return __r;
}
void QStringConverterBase__State_invalidChars_set(void* self, long long invalidChars) {
	((QStringConverterBase::State*)self)->invalidChars = invalidChars;
}
void QStringConverterBase__State_delete(void* self) { delete (QStringConverterBase::State*)self; }

}

#include <QtCore/qabstracteventdispatcher.h>
#include <QtCore/qnamespace.h>
#include "gen_qabstracteventdispatcher__timerinfo.h"

extern "C" {

void* QAbstractEventDispatcher__TimerInfo_new() {
	return new QAbstractEventDispatcher::TimerInfo();
}
void* QAbstractEventDispatcher__TimerInfo_new2(int id, int i, int t) {
	return new QAbstractEventDispatcher::TimerInfo(id, i, (Qt::TimerType)(t));
}
void* QAbstractEventDispatcher__TimerInfo_new3(void* arg1) {
	return new QAbstractEventDispatcher::TimerInfo(*(QAbstractEventDispatcher::TimerInfo*)arg1);
}
int QAbstractEventDispatcher__TimerInfo_timerId_get(void* self) {
	auto __r = ((QAbstractEventDispatcher::TimerInfo*)self)->timerId;
	return __r;
}
void QAbstractEventDispatcher__TimerInfo_timerId_set(void* self, int timerId) {
	((QAbstractEventDispatcher::TimerInfo*)self)->timerId = timerId;
}
int QAbstractEventDispatcher__TimerInfo_interval_get(void* self) {
	auto __r = ((QAbstractEventDispatcher::TimerInfo*)self)->interval;
	return __r;
}
void QAbstractEventDispatcher__TimerInfo_interval_set(void* self, int interval) {
	((QAbstractEventDispatcher::TimerInfo*)self)->interval = interval;
}
int QAbstractEventDispatcher__TimerInfo_timerType_get(void* self) {
	auto __r = ((QAbstractEventDispatcher::TimerInfo*)self)->timerType;
	return (int)(__r);
}
void QAbstractEventDispatcher__TimerInfo_timerType_set(void* self, int timerType) {
	((QAbstractEventDispatcher::TimerInfo*)self)->timerType = (Qt::TimerType)(timerType);
}
void QAbstractEventDispatcher__TimerInfo_delete(void* self) { delete (QAbstractEventDispatcher::TimerInfo*)self; }

}

#include <QtCore/qanimationgroup.h>
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
#include "gen_qanimationgroup.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QAnimationGroup__Go : public QAnimationGroup {
public:
	long long __goId;
	unsigned char __goOn[12] = {};
	template<typename... A> QAnimationGroup__Go(long long id, A&&... a) : QAnimationGroup(std::forward<A>(a)...), __goId(id) {}
	using QAnimationGroup::sender;
	using QAnimationGroup::senderSignalIndex;
	using QAnimationGroup::receivers;
	using QAnimationGroup::isSignalConnected;
	bool cppgoBase_event(QEvent* __a0) { return QAnimationGroup::event(__a0); }
	void cppgoBase_updateState(int __a0, int __a1) { QAnimationGroup::updateState((QAbstractAnimation::State)__a0, (QAbstractAnimation::State)__a1); }
	void cppgoBase_updateDirection(int __a0) { QAnimationGroup::updateDirection((QAbstractAnimation::Direction)__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QAnimationGroup::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QAnimationGroup::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QAnimationGroup::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QAnimationGroup::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QAnimationGroup::disconnectNotify(__a0); }
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
		return QAnimationGroup::metaObject();
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QAnimationGroup::event(event);
	}
	void updateState(QAbstractAnimation::State newState, QAbstractAnimation::State oldState) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_int_int) { ((void(*)(long long, long long, int, int, int*))cppgo_virt_void_int_int)(__goId, 7692482089134926569LL, (int)(newState), (int)(oldState), &__h); if (__h) return; }
		QAnimationGroup::updateState(newState, oldState);
	}
	void updateDirection(QAbstractAnimation::Direction direction) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, -3231888725407197643LL, (int)(direction), &__h); if (__h) return; }
		QAnimationGroup::updateDirection(direction);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QAnimationGroup::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QAnimationGroup::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QAnimationGroup::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QAnimationGroup::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QAnimationGroup::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QAnimationGroup::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QAnimationGroup;
static std::mutex __connsMu_QAnimationGroup;

extern "C" {

void* QAnimationGroup_base(void* self) { return (void*)static_cast<QAbstractAnimation*>((QAnimationGroup*)self); }
void QAnimationGroup_chain(void* self, void** out) {
	QAnimationGroup* __p = (QAnimationGroup*)self;
	QAbstractAnimation* __b0 = static_cast<QAbstractAnimation*>(__p); out[0] = (void*)__b0;
	QObject* __b1 = static_cast<QObject*>(__b0); out[1] = (void*)__b1;
}
void QAnimationGroup__set_override(void* self, int idx) { ((QAnimationGroup__Go*)self)->__goOn[idx] = 1; }
void* QAnimationGroup_fromQObject(void* o) { return (void*)qobject_cast<QAnimationGroup*>((QObject*)o); }
char* QAnimationGroup_tr(const char* s, long long* __out_len) {
	auto&& __r = QAnimationGroup::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QAnimationGroup_animationAt(void* self, int index) {
	auto&& __r = ((QAnimationGroup*)self)->animationAt(index);
	return (void*)(__r);
}
int QAnimationGroup_animationCount(void* self) {
	auto&& __r = ((QAnimationGroup*)self)->animationCount();
	return __r;
}
int QAnimationGroup_indexOfAnimation(void* self, void* animation) {
	auto&& __r = ((QAnimationGroup*)self)->indexOfAnimation((QAbstractAnimation*)animation);
	return __r;
}
void QAnimationGroup_addAnimation(void* self, void* animation) {
	((QAnimationGroup*)self)->addAnimation((QAbstractAnimation*)animation);
}
void QAnimationGroup_insertAnimation(void* self, int index, void* animation) {
	((QAnimationGroup*)self)->insertAnimation(index, (QAbstractAnimation*)animation);
}
void QAnimationGroup_removeAnimation(void* self, void* animation) {
	((QAnimationGroup*)self)->removeAnimation((QAbstractAnimation*)animation);
}
void* QAnimationGroup_takeAnimation(void* self, int index) {
	auto&& __r = ((QAnimationGroup*)self)->takeAnimation(index);
	return (void*)(__r);
}
void QAnimationGroup_clear(void* self) {
	((QAnimationGroup*)self)->clear();
}
void* QAnimationGroup_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QAnimationGroup::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QAnimationGroup_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QAnimationGroup::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QAnimationGroup_sender(void* self) {
	auto&& __r = ((QAnimationGroup__Go*)self)->sender();
	return (void*)(__r);
}
int QAnimationGroup_senderSignalIndex(void* self) {
	auto&& __r = ((QAnimationGroup__Go*)self)->senderSignalIndex();
	return __r;
}
int QAnimationGroup_receivers(void* self, const char* signal) {
	auto&& __r = ((QAnimationGroup__Go*)self)->receivers(signal);
	return __r;
}
int QAnimationGroup_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QAnimationGroup__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QAnimationGroup_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QAnimationGroup::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QAnimationGroup_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QAnimationGroup::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QAnimationGroup_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QAnimationGroup::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QAnimationGroup_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QAnimationGroup::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QAnimationGroup_metaObject_default(void* self) {
	auto&& __r = ((QAnimationGroup*)self)->QAnimationGroup::metaObject();
	return (void*)(__r);
}
int QAnimationGroup_event_default(void* self, void* event) {
	auto&& __r = ((QAnimationGroup__Go*)self)->cppgoBase_event((QEvent*)event);
	return __r;
}
void QAnimationGroup_updateState_default(void* self, int newState, int oldState) {
	((QAnimationGroup__Go*)self)->cppgoBase_updateState(newState, oldState);
}
void QAnimationGroup_updateDirection_default(void* self, int direction) {
	((QAnimationGroup__Go*)self)->cppgoBase_updateDirection(direction);
}
int QAnimationGroup_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QAnimationGroup*)self)->QAnimationGroup::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QAnimationGroup_timerEvent_default(void* self, void* event) {
	((QAnimationGroup__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QAnimationGroup_childEvent_default(void* self, void* event) {
	((QAnimationGroup__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QAnimationGroup_customEvent_default(void* self, void* event) {
	((QAnimationGroup__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QAnimationGroup_connectNotify_default(void* self, void* signal) {
	((QAnimationGroup__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QAnimationGroup_disconnectNotify_default(void* self, void* signal) {
	((QAnimationGroup__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QAnimationGroup_connect_finished(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAnimationGroup*)self, &QAnimationGroup::finished, (QAnimationGroup*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAnimationGroup);
	__conns_QAnimationGroup[__id] = __c;
}
void QAnimationGroup_connect_stateChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAnimationGroup*)self, &QAnimationGroup::stateChanged, (QAnimationGroup*)self, [__id](QAbstractAnimation::State newState, QAbstractAnimation::State oldState){
		if (cppgo_sig_void_int_int) ((void(*)(long long, int, int))cppgo_sig_void_int_int)(__id, (int)(newState), (int)(oldState));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAnimationGroup);
	__conns_QAnimationGroup[__id] = __c;
}
void QAnimationGroup_connect_currentLoopChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAnimationGroup*)self, &QAnimationGroup::currentLoopChanged, (QAnimationGroup*)self, [__id](int currentLoop){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, currentLoop);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAnimationGroup);
	__conns_QAnimationGroup[__id] = __c;
}
void QAnimationGroup_connect_directionChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAnimationGroup*)self, &QAnimationGroup::directionChanged, (QAnimationGroup*)self, [__id](QAbstractAnimation::Direction arg1){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, (int)(arg1));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAnimationGroup);
	__conns_QAnimationGroup[__id] = __c;
}
void QAnimationGroup_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAnimationGroup*)self, &QAnimationGroup::destroyed, (QAnimationGroup*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAnimationGroup);
	__conns_QAnimationGroup[__id] = __c;
}
void QAnimationGroup_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAnimationGroup*)self, &QAnimationGroup::destroyed, (QAnimationGroup*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAnimationGroup);
	__conns_QAnimationGroup[__id] = __c;
}
void QAnimationGroup_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAnimationGroup*)self, &QAnimationGroup::objectNameChanged, (QAnimationGroup*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAnimationGroup);
	__conns_QAnimationGroup[__id] = __c;
}
void QAnimationGroup__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QAnimationGroup);
		auto __it = __conns_QAnimationGroup.find(id);
		if (__it == __conns_QAnimationGroup.end()) return;
		__c = __it->second;
		__conns_QAnimationGroup.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QAnimationGroup__Go_new(long long id, void* parent) { return new QAnimationGroup__Go(id, (QObject*)parent); }
void QAnimationGroup_delete(void* self) { delete (QAnimationGroup*)self; }

}

#include <QtCore/qdatetime.h>
#include <QtCore/qstring.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qstringview.h>
#include "gen_qtime.h"

extern "C" {

void* QTime_new() {
	return new QTime();
}
void* QTime_new2(int h, int m) {
	return new QTime(h, m);
}
void* QTime_new3(void* arg1) {
	return new QTime(*(QTime*)arg1);
}
void* QTime_new4(int h, int m, int s) {
	return new QTime(h, m, s);
}
void* QTime_new5(int h, int m, int s, int ms) {
	return new QTime(h, m, s, ms);
}
int QTime_isNull(void* self) {
	auto&& __r = ((QTime*)self)->isNull();
	return __r;
}
int QTime_isValid(void* self) {
	auto&& __r = ((QTime*)self)->isValid();
	return __r;
}
int QTime_hour(void* self) {
	auto&& __r = ((QTime*)self)->hour();
	return __r;
}
int QTime_minute(void* self) {
	auto&& __r = ((QTime*)self)->minute();
	return __r;
}
int QTime_second(void* self) {
	auto&& __r = ((QTime*)self)->second();
	return __r;
}
int QTime_msec(void* self) {
	auto&& __r = ((QTime*)self)->msec();
	return __r;
}
char* QTime_toString(void* self, long long* __out_len) {
	auto&& __r = ((QTime*)self)->toString();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QTime_toString2(void* self, const char* format, long long format_len, long long* __out_len) {
	QString format_qstr = QString::fromUtf8(format, format_len);
	auto&& __r = ((QTime*)self)->toString(format_qstr);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QTime_toString3(void* self, void* format, long long* __out_len) {
	auto&& __r = ((QTime*)self)->toString(*(QStringView*)format);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QTime_setHMS(void* self, int h, int m, int s) {
	auto&& __r = ((QTime*)self)->setHMS(h, m, s);
	return __r;
}
void* QTime_addSecs(void* self, int secs) {
	auto&& __r = ((QTime*)self)->addSecs(secs);
	return new QTime(__r);
}
int QTime_secsTo(void* self, void* t) {
	auto&& __r = ((QTime*)self)->secsTo(*(QTime*)t);
	return __r;
}
void* QTime_addMSecs(void* self, int ms) {
	auto&& __r = ((QTime*)self)->addMSecs(ms);
	return new QTime(__r);
}
int QTime_msecsTo(void* self, void* t) {
	auto&& __r = ((QTime*)self)->msecsTo(*(QTime*)t);
	return __r;
}
void* QTime_fromMSecsSinceStartOfDay(int msecs) {
	auto&& __r = QTime::fromMSecsSinceStartOfDay(msecs);
	return new QTime(__r);
}
int QTime_msecsSinceStartOfDay(void* self) {
	auto&& __r = ((QTime*)self)->msecsSinceStartOfDay();
	return __r;
}
void* QTime_currentTime() {
	auto&& __r = QTime::currentTime();
	return new QTime(__r);
}
void* QTime_fromString(void* string_) {
	auto&& __r = QTime::fromString(*(QStringView*)string_);
	return new QTime(__r);
}
void* QTime_fromString2(void* string_, void* format) {
	auto&& __r = QTime::fromString(*(QStringView*)string_, *(QStringView*)format);
	return new QTime(__r);
}
void* QTime_fromString3(const char* string_, long long string__len, void* format) {
	QString string__qstr = QString::fromUtf8(string_, string__len);
	auto&& __r = QTime::fromString(string__qstr, *(QStringView*)format);
	return new QTime(__r);
}
void* QTime_fromString4(const char* string_, long long string__len) {
	QString string__qstr = QString::fromUtf8(string_, string__len);
	auto&& __r = QTime::fromString(string__qstr);
	return new QTime(__r);
}
void* QTime_fromString5(const char* string_, long long string__len, const char* format, long long format_len) {
	QString string__qstr = QString::fromUtf8(string_, string__len);
	QString format_qstr = QString::fromUtf8(format, format_len);
	auto&& __r = QTime::fromString(string__qstr, format_qstr);
	return new QTime(__r);
}
int QTime_isValid2(int h, int m, int s) {
	auto&& __r = QTime::isValid(h, m, s);
	return __r;
}
char* QTime_toString4(void* self, int f, long long* __out_len) {
	auto&& __r = ((QTime*)self)->toString((Qt::DateFormat)(f));
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QTime_setHMS2(void* self, int h, int m, int s, int ms) {
	auto&& __r = ((QTime*)self)->setHMS(h, m, s, ms);
	return __r;
}
void* QTime_fromString6(void* string_, int format) {
	auto&& __r = QTime::fromString(*(QStringView*)string_, (Qt::DateFormat)(format));
	return new QTime(__r);
}
void* QTime_fromString7(const char* string_, long long string__len, int format) {
	QString string__qstr = QString::fromUtf8(string_, string__len);
	auto&& __r = QTime::fromString(string__qstr, (Qt::DateFormat)(format));
	return new QTime(__r);
}
int QTime_isValid3(int h, int m, int s, int ms) {
	auto&& __r = QTime::isValid(h, m, s, ms);
	return __r;
}
int QTime_OperatorEqual(void* self, void* rhs) {
	auto&& __r = (*(QTime*)self) == (*(QTime*)rhs);
	return __r;
}
int QTime_OperatorNotEqual(void* self, void* rhs) {
	auto&& __r = (*(QTime*)self) != (*(QTime*)rhs);
	return __r;
}
int QTime_OperatorLesser(void* self, void* rhs) {
	auto&& __r = (*(QTime*)self) < (*(QTime*)rhs);
	return __r;
}
int QTime_OperatorGreater(void* self, void* rhs) {
	auto&& __r = (*(QTime*)self) > (*(QTime*)rhs);
	return __r;
}
int QTime_OperatorLesserOrEqual(void* self, void* rhs) {
	auto&& __r = (*(QTime*)self) <= (*(QTime*)rhs);
	return __r;
}
int QTime_OperatorGreaterOrEqual(void* self, void* rhs) {
	auto&& __r = (*(QTime*)self) >= (*(QTime*)rhs);
	return __r;
}
void QTime_delete(void* self) { delete (QTime*)self; }

}

#include <QtCore/qproperty.h>
#include <QtCore/qpropertyprivate.h>
#include "gen_qpropertyobserver.h"

extern "C" {

void* QPropertyObserver_base(void* self) { return (void*)static_cast<QPropertyObserverBase*>((QPropertyObserver*)self); }
void* QPropertyObserver_new() {
	return new QPropertyObserver();
}
void QPropertyObserver_delete(void* self) { delete (QPropertyObserver*)self; }

}

#include <QtCore/qexception.h>
#include "gen_qexception.h"
#include "gen_purego_cb.h"

class QException__Go : public QException {
public:
	long long __goId;
	unsigned char __goOn[2] = {};
	template<typename... A> QException__Go(long long id, A&&... a) : QException(std::forward<A>(a)...), __goId(id) {}
	void raise() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 8013398049691041699LL, &__h); if (__h) return; }
		QException::raise();
	}
	QException* clone() const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, -7032813281048616478LL, &__h); if (__h) return (QException*)__r; }
		return QException::clone();
	}
};

extern "C" {

void QException__set_override(void* self, int idx) { ((QException__Go*)self)->__goOn[idx] = 1; }
void* QException_new(long long __goid) {
	return new QException__Go(__goid);
}
void* QException_new2(long long __goid, void* arg1) {
	return new QException__Go(__goid, *(QException*)arg1);
}
void* QException_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QException*)self)->operator=(*(QException*)arg1);
	return (void*)&(__r);
}
void QException_raise(void* self) {
	((QException*)self)->raise();
}
void* QException_clone(void* self) {
	auto&& __r = ((QException*)self)->clone();
	return (void*)(__r);
}
const char* QException_what(void* self) {
	auto&& __r = ((QException*)self)->what();
	return __r;
}
void QException_raise_default(void* self) {
	((QException*)self)->QException::raise();
}
void* QException_clone_default(void* self) {
	auto&& __r = ((QException*)self)->QException::clone();
	return (void*)(__r);
}
void* QException__Go_new(long long id) { return new QException__Go(id); }
void QException_delete(void* self) { delete (QException*)self; }

}

#include <QtCore/qversionnumber.h>
#include <QtCore/qstring.h>
#include <QtCore/qanystringview.h>
#include "gen_qversionnumber.h"

extern "C" {

void* QVersionNumber_new() {
	return new QVersionNumber();
}
void* QVersionNumber_new2(int maj) {
	return new QVersionNumber(maj);
}
void* QVersionNumber_new3(int maj, int min) {
	return new QVersionNumber(maj, min);
}
void* QVersionNumber_new4(int maj, int min, int mic) {
	return new QVersionNumber(maj, min, mic);
}
void* QVersionNumber_new5(void* arg1) {
	return new QVersionNumber(*(QVersionNumber*)arg1);
}
int QVersionNumber_isNull(void* self) {
	auto&& __r = ((QVersionNumber*)self)->isNull();
	return __r;
}
int QVersionNumber_isNormalized(void* self) {
	auto&& __r = ((QVersionNumber*)self)->isNormalized();
	return __r;
}
int QVersionNumber_majorVersion(void* self) {
	auto&& __r = ((QVersionNumber*)self)->majorVersion();
	return __r;
}
int QVersionNumber_minorVersion(void* self) {
	auto&& __r = ((QVersionNumber*)self)->minorVersion();
	return __r;
}
int QVersionNumber_microVersion(void* self) {
	auto&& __r = ((QVersionNumber*)self)->microVersion();
	return __r;
}
void* QVersionNumber_normalized(void* self) {
	auto&& __r = ((QVersionNumber*)self)->normalized();
	return new QVersionNumber(__r);
}
int* QVersionNumber_segments(void* self, long long* __out_len) {
	auto&& __r = ((QVersionNumber*)self)->segments();
	*__out_len = __r.size();
	int* __arr = (int*)malloc(sizeof(int) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = __e;
	return __arr;
}
int QVersionNumber_segmentAt(void* self, long long index) {
	auto&& __r = ((QVersionNumber*)self)->segmentAt(index);
	return __r;
}
long long QVersionNumber_segmentCount(void* self) {
	auto&& __r = ((QVersionNumber*)self)->segmentCount();
	return __r;
}
int QVersionNumber_isPrefixOf(void* self, void* other) {
	auto&& __r = ((QVersionNumber*)self)->isPrefixOf(*(QVersionNumber*)other);
	return __r;
}
int QVersionNumber_compare(void* v1, void* v2) {
	auto&& __r = QVersionNumber::compare(*(QVersionNumber*)v1, *(QVersionNumber*)v2);
	return __r;
}
void* QVersionNumber_commonPrefix(void* v1, void* v2) {
	auto&& __r = QVersionNumber::commonPrefix(*(QVersionNumber*)v1, *(QVersionNumber*)v2);
	return new QVersionNumber(__r);
}
char* QVersionNumber_toString(void* self, long long* __out_len) {
	auto&& __r = ((QVersionNumber*)self)->toString();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QVersionNumber_fromString(const char* string_, long long string__len) {
	QAnyStringView string__qstr = QAnyStringView(string_, string__len);
	auto&& __r = QVersionNumber::fromString(string__qstr);
	return new QVersionNumber(__r);
}
void* QVersionNumber_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QVersionNumber*)self)->operator=(*(QVersionNumber*)arg1);
	return (void*)&(__r);
}
void* QVersionNumber_fromString2(const char* string_, long long string__len, long long* suffixIndex) {
	QAnyStringView string__qstr = QAnyStringView(string_, string__len);
	qsizetype suffixIndex_local{};
	auto&& __r = QVersionNumber::fromString(string__qstr, &suffixIndex_local);
	*suffixIndex = suffixIndex_local;
	return new QVersionNumber(__r);
}
int QVersionNumber_OperatorEqual(void* self, void* rhs) {
	auto&& __r = (*(QVersionNumber*)self) == (*(QVersionNumber*)rhs);
	return __r;
}
int QVersionNumber_OperatorNotEqual(void* self, void* rhs) {
	auto&& __r = (*(QVersionNumber*)self) != (*(QVersionNumber*)rhs);
	return __r;
}
int QVersionNumber_OperatorLesser(void* self, void* rhs) {
	auto&& __r = (*(QVersionNumber*)self) < (*(QVersionNumber*)rhs);
	return __r;
}
int QVersionNumber_OperatorGreater(void* self, void* rhs) {
	auto&& __r = (*(QVersionNumber*)self) > (*(QVersionNumber*)rhs);
	return __r;
}
int QVersionNumber_OperatorLesserOrEqual(void* self, void* rhs) {
	auto&& __r = (*(QVersionNumber*)self) <= (*(QVersionNumber*)rhs);
	return __r;
}
int QVersionNumber_OperatorGreaterOrEqual(void* self, void* rhs) {
	auto&& __r = (*(QVersionNumber*)self) >= (*(QVersionNumber*)rhs);
	return __r;
}
void QVersionNumber_delete(void* self) { delete (QVersionNumber*)self; }

}

#include <QtCore/qpermissions.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qmetatype.h>
#include "gen_qpermission.h"

extern "C" {

void* QPermission_new() {
	return new QPermission();
}
void* QPermission_new2(void* arg1) {
	return new QPermission(*(QPermission*)arg1);
}
int QPermission_status(void* self) {
	auto&& __r = ((QPermission*)self)->status();
	return (int)(__r);
}
void* QPermission_type(void* self) {
	auto&& __r = ((QPermission*)self)->type();
	return new QMetaType(__r);
}
void* QPermission_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QPermission*)self)->operator=(*(QPermission*)arg1);
	return (void*)&(__r);
}
void QPermission_delete(void* self) { delete (QPermission*)self; }

}

#include <QtCore/qreadwritelock.h>
#include "gen_qwritelocker.h"

extern "C" {

void* QWriteLocker_new(void* readWriteLock) {
	return new QWriteLocker((QReadWriteLock*)readWriteLock);
}
void QWriteLocker_unlock(void* self) {
	((QWriteLocker*)self)->unlock();
}
void QWriteLocker_relock(void* self) {
	((QWriteLocker*)self)->relock();
}
void* QWriteLocker_readWriteLock(void* self) {
	auto&& __r = ((QWriteLocker*)self)->readWriteLock();
	return (void*)(__r);
}
void QWriteLocker_delete(void* self) { delete (QWriteLocker*)self; }

}

#include <QtCore/qtemporarydir.h>
#include <QtCore/qstring.h>
#include "gen_qtemporarydir.h"

extern "C" {

void* QTemporaryDir_new() {
	return new QTemporaryDir();
}
void* QTemporaryDir_new2(const char* templateName, long long templateName_len) {
	QString templateName_qstr = QString::fromUtf8(templateName, templateName_len);
	return new QTemporaryDir(templateName_qstr);
}
void QTemporaryDir_swap(void* self, void* other) {
	((QTemporaryDir*)self)->swap(*(QTemporaryDir*)other);
}
int QTemporaryDir_isValid(void* self) {
	auto&& __r = ((QTemporaryDir*)self)->isValid();
	return __r;
}
char* QTemporaryDir_errorString(void* self, long long* __out_len) {
	auto&& __r = ((QTemporaryDir*)self)->errorString();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QTemporaryDir_autoRemove(void* self) {
	auto&& __r = ((QTemporaryDir*)self)->autoRemove();
	return __r;
}
void QTemporaryDir_setAutoRemove(void* self, int b) {
	((QTemporaryDir*)self)->setAutoRemove(b);
}
int QTemporaryDir_remove(void* self) {
	auto&& __r = ((QTemporaryDir*)self)->remove();
	return __r;
}
char* QTemporaryDir_path(void* self, long long* __out_len) {
	auto&& __r = ((QTemporaryDir*)self)->path();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QTemporaryDir_filePath(void* self, const char* fileName, long long fileName_len, long long* __out_len) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	auto&& __r = ((QTemporaryDir*)self)->filePath(fileName_qstr);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QTemporaryDir_delete(void* self) { delete (QTemporaryDir*)self; }

}

#include <QtGui/qregion.h>
#include <QtCore/qrect.h>
#include <QtGui/qpolygon.h>
#include <QtCore/qnamespace.h>
#include <QtGui/qbitmap.h>
#include <QtCore/qpoint.h>
#include <QtCore/qvariant.h>
#include <QtGui/qtransform.h>
#include "gen_qregion.h"

extern "C" {

void* QRegion_new() {
	return new QRegion();
}
void* QRegion_new2(int x, int y, int w, int h) {
	return new QRegion(x, y, w, h);
}
void* QRegion_new3(void* r) {
	return new QRegion(*(QRect*)r);
}
void* QRegion_new4(void* pa) {
	return new QRegion(*(QPolygon*)pa);
}
void* QRegion_new5(void* region) {
	return new QRegion(*(QRegion*)region);
}
void* QRegion_new6(void* bitmap) {
	return new QRegion(*(QBitmap*)bitmap);
}
void* QRegion_new7(int x, int y, int w, int h, int t) {
	return new QRegion(x, y, w, h, (QRegion::RegionType)(t));
}
void* QRegion_new8(void* r, int t) {
	return new QRegion(*(QRect*)r, (QRegion::RegionType)(t));
}
void* QRegion_new9(void* pa, int fillRule) {
	return new QRegion(*(QPolygon*)pa, (Qt::FillRule)(fillRule));
}
void* QRegion_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QRegion*)self)->operator=(*(QRegion*)arg1);
	return (void*)&(__r);
}
void QRegion_swap(void* self, void* other) {
	((QRegion*)self)->swap(*(QRegion*)other);
}
int QRegion_isEmpty(void* self) {
	auto&& __r = ((QRegion*)self)->isEmpty();
	return __r;
}
int QRegion_isNull(void* self) {
	auto&& __r = ((QRegion*)self)->isNull();
	return __r;
}
void* QRegion_begin(void* self) {
	auto&& __r = ((QRegion*)self)->begin();
	return (void*)(__r);
}
void* QRegion_cbegin(void* self) {
	auto&& __r = ((QRegion*)self)->cbegin();
	return (void*)(__r);
}
void* QRegion_end(void* self) {
	auto&& __r = ((QRegion*)self)->end();
	return (void*)(__r);
}
void* QRegion_cend(void* self) {
	auto&& __r = ((QRegion*)self)->cend();
	return (void*)(__r);
}
int QRegion_contains(void* self, void* p) {
	auto&& __r = ((QRegion*)self)->contains(*(QPoint*)p);
	return __r;
}
int QRegion_contains2(void* self, void* r) {
	auto&& __r = ((QRegion*)self)->contains(*(QRect*)r);
	return __r;
}
void QRegion_translate(void* self, int dx, int dy) {
	((QRegion*)self)->translate(dx, dy);
}
void QRegion_translate2(void* self, void* p) {
	((QRegion*)self)->translate(*(QPoint*)p);
}
void* QRegion_translated(void* self, int dx, int dy) {
	auto&& __r = ((QRegion*)self)->translated(dx, dy);
	return new QRegion(__r);
}
void* QRegion_translated2(void* self, void* p) {
	auto&& __r = ((QRegion*)self)->translated(*(QPoint*)p);
	return new QRegion(__r);
}
void* QRegion_united(void* self, void* r) {
	auto&& __r = ((QRegion*)self)->united(*(QRegion*)r);
	return new QRegion(__r);
}
void* QRegion_united2(void* self, void* r) {
	auto&& __r = ((QRegion*)self)->united(*(QRect*)r);
	return new QRegion(__r);
}
void* QRegion_intersected(void* self, void* r) {
	auto&& __r = ((QRegion*)self)->intersected(*(QRegion*)r);
	return new QRegion(__r);
}
void* QRegion_intersected2(void* self, void* r) {
	auto&& __r = ((QRegion*)self)->intersected(*(QRect*)r);
	return new QRegion(__r);
}
void* QRegion_subtracted(void* self, void* r) {
	auto&& __r = ((QRegion*)self)->subtracted(*(QRegion*)r);
	return new QRegion(__r);
}
void* QRegion_xored(void* self, void* r) {
	auto&& __r = ((QRegion*)self)->xored(*(QRegion*)r);
	return new QRegion(__r);
}
int QRegion_intersects(void* self, void* r) {
	auto&& __r = ((QRegion*)self)->intersects(*(QRegion*)r);
	return __r;
}
int QRegion_intersects2(void* self, void* r) {
	auto&& __r = ((QRegion*)self)->intersects(*(QRect*)r);
	return __r;
}
void* QRegion_boundingRect(void* self) {
	auto&& __r = ((QRegion*)self)->boundingRect();
	return new QRect(__r);
}
void QRegion_setRects(void* self, void* rect, int num) {
	((QRegion*)self)->setRects((QRect*)rect, num);
}
int QRegion_rectCount(void* self) {
	auto&& __r = ((QRegion*)self)->rectCount();
	return __r;
}
void* QRegion_OperatorBitwiseOr(void* self, void* r) {
	auto&& __r = ((QRegion*)self)->operator|(*(QRegion*)r);
	return new QRegion(__r);
}
void* QRegion_OperatorPlus(void* self, void* r) {
	auto&& __r = ((QRegion*)self)->operator+(*(QRegion*)r);
	return new QRegion(__r);
}
void* QRegion_OperatorPlus2(void* self, void* r) {
	auto&& __r = ((QRegion*)self)->operator+(*(QRect*)r);
	return new QRegion(__r);
}
void* QRegion_OperatorBitwiseAnd(void* self, void* r) {
	auto&& __r = ((QRegion*)self)->operator&(*(QRegion*)r);
	return new QRegion(__r);
}
void* QRegion_OperatorBitwiseAnd2(void* self, void* r) {
	auto&& __r = ((QRegion*)self)->operator&(*(QRect*)r);
	return new QRegion(__r);
}
void* QRegion_OperatorMinus(void* self, void* r) {
	auto&& __r = ((QRegion*)self)->operator-(*(QRegion*)r);
	return new QRegion(__r);
}
void* QRegion_OperatorBitwiseXor(void* self, void* r) {
	auto&& __r = ((QRegion*)self)->operator^(*(QRegion*)r);
	return new QRegion(__r);
}
void* QRegion_OperatorBitwiseOrAssign(void* self, void* r) {
	auto&& __r = ((QRegion*)self)->operator|=(*(QRegion*)r);
	return (void*)&(__r);
}
void* QRegion_OperatorPlusAssign(void* self, void* r) {
	auto&& __r = ((QRegion*)self)->operator+=(*(QRegion*)r);
	return (void*)&(__r);
}
void* QRegion_OperatorPlusAssign2(void* self, void* r) {
	auto&& __r = ((QRegion*)self)->operator+=(*(QRect*)r);
	return (void*)&(__r);
}
void* QRegion_OperatorBitwiseAndAssign(void* self, void* r) {
	auto&& __r = ((QRegion*)self)->operator&=(*(QRegion*)r);
	return (void*)&(__r);
}
void* QRegion_OperatorBitwiseAndAssign2(void* self, void* r) {
	auto&& __r = ((QRegion*)self)->operator&=(*(QRect*)r);
	return (void*)&(__r);
}
void* QRegion_OperatorMinusAssign(void* self, void* r) {
	auto&& __r = ((QRegion*)self)->operator-=(*(QRegion*)r);
	return (void*)&(__r);
}
void* QRegion_OperatorBitwiseXorAssign(void* self, void* r) {
	auto&& __r = ((QRegion*)self)->operator^=(*(QRegion*)r);
	return (void*)&(__r);
}
int QRegion_OperatorEqual(void* self, void* r) {
	auto&& __r = ((QRegion*)self)->operator==(*(QRegion*)r);
	return __r;
}
int QRegion_OperatorNotEqual(void* self, void* r) {
	auto&& __r = ((QRegion*)self)->operator!=(*(QRegion*)r);
	return __r;
}
void* QRegion_ToQVariant(void* self) {
	auto&& __r = static_cast<QVariant>(*(QRegion*)self);
	return new QVariant(__r);
}
void* QRegion_OperatorMultiply(void* self, void* m) {
	auto&& __r = (*(QRegion*)self) * (*(QTransform*)m);
	return new QRegion(__r);
}
void* QRegion_toQVariant(void* self) {
	auto&& __r = QVariant::fromValue(*(QRegion*)self);
	return new QVariant(__r);
}
void QRegion_delete(void* self) { delete (QRegion*)self; }

}

#include <QtGui/qevent.h>
#include <QtCore/qcoreevent.h>
#include <QtGui/qpointingdevice.h>
#include <QtCore/qpoint.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qobject.h>
#include <QtGui/qeventpoint.h>
#include <QtGui/qinputdevice.h>
#include "gen_qtabletevent.h"

extern "C" {

void* QTabletEvent_base(void* self) { return (void*)static_cast<QSinglePointEvent*>((QTabletEvent*)self); }
void QTabletEvent_chain(void* self, void** out) {
	QTabletEvent* __p = (QTabletEvent*)self;
	QSinglePointEvent* __b0 = static_cast<QSinglePointEvent*>(__p); out[0] = (void*)__b0;
	QPointerEvent* __b1 = static_cast<QPointerEvent*>(__b0); out[1] = (void*)__b1;
	QInputEvent* __b2 = static_cast<QInputEvent*>(__b1); out[2] = (void*)__b2;
	QEvent* __b3 = static_cast<QEvent*>(__b2); out[3] = (void*)__b3;
}
void* QTabletEvent_new(int t, void* device, void* pos, void* globalPos, double pressure, float xTilt, float yTilt, float tangentialPressure, double rotation, float z, int keyState, int button, int buttons) {
	return new QTabletEvent((QEvent::Type)(t), (QPointingDevice*)device, *(QPointF*)pos, *(QPointF*)globalPos, pressure, xTilt, yTilt, tangentialPressure, rotation, z, QFlags<Qt::KeyboardModifier>(keyState), (Qt::MouseButton)(button), QFlags<Qt::MouseButton>(buttons));
}
void* QTabletEvent_pos(void* self) {
	auto&& __r = ((QTabletEvent*)self)->pos();
	return new QPoint(__r);
}
void* QTabletEvent_globalPos(void* self) {
	auto&& __r = ((QTabletEvent*)self)->globalPos();
	return new QPoint(__r);
}
void* QTabletEvent_posF(void* self) {
	auto&& __r = ((QTabletEvent*)self)->posF();
	return new QPointF(__r);
}
void* QTabletEvent_globalPosF(void* self) {
	auto&& __r = ((QTabletEvent*)self)->globalPosF();
	return new QPointF(__r);
}
int QTabletEvent_x(void* self) {
	auto&& __r = ((QTabletEvent*)self)->x();
	return __r;
}
int QTabletEvent_y(void* self) {
	auto&& __r = ((QTabletEvent*)self)->y();
	return __r;
}
int QTabletEvent_globalX(void* self) {
	auto&& __r = ((QTabletEvent*)self)->globalX();
	return __r;
}
int QTabletEvent_globalY(void* self) {
	auto&& __r = ((QTabletEvent*)self)->globalY();
	return __r;
}
double QTabletEvent_hiResGlobalX(void* self) {
	auto&& __r = ((QTabletEvent*)self)->hiResGlobalX();
	return __r;
}
double QTabletEvent_hiResGlobalY(void* self) {
	auto&& __r = ((QTabletEvent*)self)->hiResGlobalY();
	return __r;
}
long long QTabletEvent_uniqueId(void* self) {
	auto&& __r = ((QTabletEvent*)self)->uniqueId();
	return __r;
}
double QTabletEvent_pressure(void* self) {
	auto&& __r = ((QTabletEvent*)self)->pressure();
	return __r;
}
double QTabletEvent_rotation(void* self) {
	auto&& __r = ((QTabletEvent*)self)->rotation();
	return __r;
}
double QTabletEvent_z(void* self) {
	auto&& __r = ((QTabletEvent*)self)->z();
	return __r;
}
double QTabletEvent_tangentialPressure(void* self) {
	auto&& __r = ((QTabletEvent*)self)->tangentialPressure();
	return __r;
}
double QTabletEvent_xTilt(void* self) {
	auto&& __r = ((QTabletEvent*)self)->xTilt();
	return __r;
}
double QTabletEvent_yTilt(void* self) {
	auto&& __r = ((QTabletEvent*)self)->yTilt();
	return __r;
}
int QTabletEvent_registerEventType() {
	auto&& __r = QTabletEvent::registerEventType();
	return __r;
}
int QTabletEvent_registerEventType2(int hint) {
	auto&& __r = QTabletEvent::registerEventType(hint);
	return __r;
}
void QTabletEvent_delete(void* self) { delete (QTabletEvent*)self; }

}

#include <QtGui/qevent.h>
#include <QtCore/qcoreevent.h>
#include <QtGui/qaction.h>
#include "gen_qactionevent.h"

extern "C" {

void* QActionEvent_base(void* self) { return (void*)static_cast<QEvent*>((QActionEvent*)self); }
void* QActionEvent_new(int type_, void* action) {
	return new QActionEvent(type_, (QAction*)action);
}
void* QActionEvent_new2(int type_, void* action, void* before) {
	return new QActionEvent(type_, (QAction*)action, (QAction*)before);
}
void* QActionEvent_action(void* self) {
	auto&& __r = ((QActionEvent*)self)->action();
	return (void*)(__r);
}
void* QActionEvent_before(void* self) {
	auto&& __r = ((QActionEvent*)self)->before();
	return (void*)(__r);
}
int QActionEvent_registerEventType() {
	auto&& __r = QActionEvent::registerEventType();
	return __r;
}
int QActionEvent_registerEventType2(int hint) {
	auto&& __r = QActionEvent::registerEventType(hint);
	return __r;
}
void QActionEvent_delete(void* self) { delete (QActionEvent*)self; }

}

#include <QtGui/qbrush.h>
#include <QtCore/qpoint.h>
#include <QtGui/qcolor.h>
#include "gen_qconicalgradient.h"

extern "C" {

void* QConicalGradient_base(void* self) { return (void*)static_cast<QGradient*>((QConicalGradient*)self); }
void* QConicalGradient_new() {
	return new QConicalGradient();
}
void* QConicalGradient_new2(void* center, double startAngle) {
	return new QConicalGradient(*(QPointF*)center, startAngle);
}
void* QConicalGradient_new3(double cx, double cy, double startAngle) {
	return new QConicalGradient(cx, cy, startAngle);
}
void* QConicalGradient_new4(void* arg1) {
	return new QConicalGradient(*(QConicalGradient*)arg1);
}
void* QConicalGradient_center(void* self) {
	auto&& __r = ((QConicalGradient*)self)->center();
	return new QPointF(__r);
}
void QConicalGradient_setCenter(void* self, void* center) {
	((QConicalGradient*)self)->setCenter(*(QPointF*)center);
}
void QConicalGradient_setCenter2(void* self, double x, double y) {
	((QConicalGradient*)self)->setCenter(x, y);
}
double QConicalGradient_angle(void* self) {
	auto&& __r = ((QConicalGradient*)self)->angle();
	return __r;
}
void QConicalGradient_setAngle(void* self, double angle) {
	((QConicalGradient*)self)->setAngle(angle);
}
void QConicalGradient_delete(void* self) { delete (QConicalGradient*)self; }

}

#include <QtGui/qaccessible_base.h>
#include <QtGui/qaccessible.h>
#include <QtCore/qobject.h>
#include <QtGui/qtextcursor.h>
#include "gen_qaccessible.h"

extern "C" {

void QAccessible_installActivationObserver(void* arg1) {
	QAccessible::installActivationObserver((QAccessible::ActivationObserver*)arg1);
}
void QAccessible_removeActivationObserver(void* arg1) {
	QAccessible::removeActivationObserver((QAccessible::ActivationObserver*)arg1);
}
void* QAccessible_queryAccessibleInterface(void* arg1) {
	auto&& __r = QAccessible::queryAccessibleInterface((QObject*)arg1);
	return (void*)(__r);
}
unsigned int QAccessible_uniqueId(void* iface) {
	auto&& __r = QAccessible::uniqueId((QAccessibleInterface*)iface);
	return __r;
}
void* QAccessible_accessibleInterface(unsigned int uniqueId) {
	auto&& __r = QAccessible::accessibleInterface(uniqueId);
	return (void*)(__r);
}
unsigned int QAccessible_registerAccessibleInterface(void* iface) {
	auto&& __r = QAccessible::registerAccessibleInterface((QAccessibleInterface*)iface);
	return __r;
}
void QAccessible_deleteAccessibleInterface(unsigned int uniqueId) {
	QAccessible::deleteAccessibleInterface(uniqueId);
}
void QAccessible_updateAccessibility(void* event) {
	QAccessible::updateAccessibility((QAccessibleEvent*)event);
}
int QAccessible_isActive() {
	auto&& __r = QAccessible::isActive();
	return __r;
}
void QAccessible_setActive(int active) {
	QAccessible::setActive(active);
}
void QAccessible_setRootObject(void* object) {
	QAccessible::setRootObject((QObject*)object);
}
void QAccessible_cleanup() {
	QAccessible::cleanup();
}
int QAccessible_qAccessibleTextBoundaryHelper(void* cursor, int boundaryType, int* __out_b) {
	auto&& __r = QAccessible::qAccessibleTextBoundaryHelper(*(QTextCursor*)cursor, (QAccessible::TextBoundaryType)(boundaryType));
	*__out_b = __r.second;
	return __r.first;
}
void QAccessible_delete(void* self) { delete (QAccessible*)self; }

}

#include <QtGui/qaccessiblebridge.h>
#include <QtGui/qaccessible.h>
#include "gen_qaccessiblebridge.h"
#include "gen_purego_cb.h"

class QAccessibleBridge__Go : public QAccessibleBridge {
public:
	long long __goId;
	unsigned char __goOn[2] = {};
	template<typename... A> QAccessibleBridge__Go(long long id, A&&... a) : QAccessibleBridge(std::forward<A>(a)...), __goId(id) {}
	void setRootObject(QAccessibleInterface* arg1) override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7245844568460288668LL, (void*)arg1, &__h); if (__h) return; }
		return;
	}
	void notifyAccessibilityUpdate(QAccessibleEvent* event) override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4087650821821605459LL, (void*)event, &__h); if (__h) return; }
		return;
	}
};

extern "C" {

void QAccessibleBridge__set_override(void* self, int idx) { ((QAccessibleBridge__Go*)self)->__goOn[idx] = 1; }
void QAccessibleBridge_setRootObject(void* self, void* arg1) {
	((QAccessibleBridge*)self)->setRootObject((QAccessibleInterface*)arg1);
}
void QAccessibleBridge_notifyAccessibilityUpdate(void* self, void* event) {
	((QAccessibleBridge*)self)->notifyAccessibilityUpdate((QAccessibleEvent*)event);
}
void* QAccessibleBridge_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QAccessibleBridge*)self)->operator=(*(QAccessibleBridge*)arg1);
	return (void*)&(__r);
}
void* QAccessibleBridge__Go_new(long long id) { return new QAccessibleBridge__Go(id); }
void QAccessibleBridge_delete(void* self) { delete (QAccessibleBridge*)self; }

}

#include <QtGui/qgenericplugin.h>
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
#include "gen_qgenericplugin.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QGenericPlugin__Go : public QGenericPlugin {
public:
	long long __goId;
	unsigned char __goOn[9] = {};
	template<typename... A> QGenericPlugin__Go(long long id, A&&... a) : QGenericPlugin(std::forward<A>(a)...), __goId(id) {}
	using QGenericPlugin::sender;
	using QGenericPlugin::senderSignalIndex;
	using QGenericPlugin::receivers;
	using QGenericPlugin::isSignalConnected;
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QGenericPlugin::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QGenericPlugin::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QGenericPlugin::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QGenericPlugin::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QGenericPlugin::disconnectNotify(__a0); }
	QObject* create(const QString& name, const QString& spec) override {
		int __h = 0;
		QByteArray name_ba = name.toUtf8();
		QByteArray spec_ba = spec.toUtf8();
		if (__goOn[0] && cppgo_virt_voidp_constcharp_longlong_constcharp_longlong) { void* __r = ((void*(*)(long long, long long, const char*, long long, const char*, long long, int*))cppgo_virt_voidp_constcharp_longlong_constcharp_longlong)(__goId, -626457835795418115LL, (char*)name_ba.constData(), name_ba.size(), (char*)spec_ba.constData(), spec_ba.size(), &__h); if (__h) return (QObject*)__r; }
		return nullptr;
	}
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QGenericPlugin::metaObject();
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QGenericPlugin::event(event);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QGenericPlugin::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QGenericPlugin::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QGenericPlugin::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QGenericPlugin::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QGenericPlugin::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QGenericPlugin::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QGenericPlugin;
static std::mutex __connsMu_QGenericPlugin;

extern "C" {

void* QGenericPlugin_base(void* self) { return (void*)static_cast<QObject*>((QGenericPlugin*)self); }
void QGenericPlugin__set_override(void* self, int idx) { ((QGenericPlugin__Go*)self)->__goOn[idx] = 1; }
void* QGenericPlugin_fromQObject(void* o) { return (void*)qobject_cast<QGenericPlugin*>((QObject*)o); }
char* QGenericPlugin_tr(const char* s, long long* __out_len) {
	auto&& __r = QGenericPlugin::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QGenericPlugin_create(void* self, const char* name, long long name_len, const char* spec, long long spec_len) {
	QString name_qstr = QString::fromUtf8(name, name_len);
	QString spec_qstr = QString::fromUtf8(spec, spec_len);
	auto&& __r = ((QGenericPlugin*)self)->create(name_qstr, spec_qstr);
	return (void*)(__r);
}
void* QGenericPlugin_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QGenericPlugin::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QGenericPlugin_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QGenericPlugin::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QGenericPlugin_sender(void* self) {
	auto&& __r = ((QGenericPlugin__Go*)self)->sender();
	return (void*)(__r);
}
int QGenericPlugin_senderSignalIndex(void* self) {
	auto&& __r = ((QGenericPlugin__Go*)self)->senderSignalIndex();
	return __r;
}
int QGenericPlugin_receivers(void* self, const char* signal) {
	auto&& __r = ((QGenericPlugin__Go*)self)->receivers(signal);
	return __r;
}
int QGenericPlugin_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QGenericPlugin__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QGenericPlugin_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QGenericPlugin::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QGenericPlugin_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QGenericPlugin::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QGenericPlugin_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QGenericPlugin::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QGenericPlugin_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QGenericPlugin::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QGenericPlugin_metaObject_default(void* self) {
	auto&& __r = ((QGenericPlugin*)self)->QGenericPlugin::metaObject();
	return (void*)(__r);
}
int QGenericPlugin_event_default(void* self, void* event) {
	auto&& __r = ((QGenericPlugin*)self)->QGenericPlugin::event((QEvent*)event);
	return __r;
}
int QGenericPlugin_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QGenericPlugin*)self)->QGenericPlugin::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QGenericPlugin_timerEvent_default(void* self, void* event) {
	((QGenericPlugin__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QGenericPlugin_childEvent_default(void* self, void* event) {
	((QGenericPlugin__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QGenericPlugin_customEvent_default(void* self, void* event) {
	((QGenericPlugin__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QGenericPlugin_connectNotify_default(void* self, void* signal) {
	((QGenericPlugin__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QGenericPlugin_disconnectNotify_default(void* self, void* signal) {
	((QGenericPlugin__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QGenericPlugin_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGenericPlugin*)self, &QGenericPlugin::destroyed, (QGenericPlugin*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGenericPlugin);
	__conns_QGenericPlugin[__id] = __c;
}
void QGenericPlugin_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGenericPlugin*)self, &QGenericPlugin::destroyed, (QGenericPlugin*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGenericPlugin);
	__conns_QGenericPlugin[__id] = __c;
}
void QGenericPlugin_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGenericPlugin*)self, &QGenericPlugin::objectNameChanged, (QGenericPlugin*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGenericPlugin);
	__conns_QGenericPlugin[__id] = __c;
}
void QGenericPlugin__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QGenericPlugin);
		auto __it = __conns_QGenericPlugin.find(id);
		if (__it == __conns_QGenericPlugin.end()) return;
		__c = __it->second;
		__conns_QGenericPlugin.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QGenericPlugin__Go_new(long long id, void* parent) { return new QGenericPlugin__Go(id, (QObject*)parent); }
void QGenericPlugin_delete(void* self) { delete (QGenericPlugin*)self; }

}

#include <QtGui/qpaintengine.h>
#include <QtCore/qstring.h>
#include <QtGui/qfont.h>
#include "gen_qtextitem.h"

extern "C" {

void* QTextItem_new() {
	return new QTextItem();
}
void* QTextItem_new2(void* arg1) {
	return new QTextItem(*(QTextItem*)arg1);
}
double QTextItem_descent(void* self) {
	auto&& __r = ((QTextItem*)self)->descent();
	return __r;
}
double QTextItem_ascent(void* self) {
	auto&& __r = ((QTextItem*)self)->ascent();
	return __r;
}
double QTextItem_width(void* self) {
	auto&& __r = ((QTextItem*)self)->width();
	return __r;
}
int QTextItem_renderFlags(void* self) {
	auto&& __r = ((QTextItem*)self)->renderFlags();
	return (int)(__r);
}
char* QTextItem_text(void* self, long long* __out_len) {
	auto&& __r = ((QTextItem*)self)->text();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QTextItem_font(void* self) {
	auto&& __r = ((QTextItem*)self)->font();
	return new QFont(__r);
}
void QTextItem_delete(void* self) { delete (QTextItem*)self; }

}

#include <QtGui/qtextobject.h>
#include "gen_qtextframelayoutdata.h"

extern "C" {

void* QTextFrameLayoutData_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QTextFrameLayoutData*)self)->operator=(*(QTextFrameLayoutData*)arg1);
	return (void*)&(__r);
}
void QTextFrameLayoutData_delete(void* self) { delete (QTextFrameLayoutData*)self; }

}

#include <QtWidgets/qabstractslider.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qcoreevent.h>
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
#include "gen_qabstractslider.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QAbstractSlider__Go : public QAbstractSlider {
public:
	long long __goId;
	unsigned char __goOn[49] = {};
	template<typename... A> QAbstractSlider__Go(long long id, A&&... a) : QAbstractSlider(std::forward<A>(a)...), __goId(id) {}
	using QAbstractSlider::setRepeatAction;
	using QAbstractSlider::repeatAction;
	using QAbstractSlider::updateMicroFocus;
	using QAbstractSlider::create;
	using QAbstractSlider::destroy;
	using QAbstractSlider::focusNextChild;
	using QAbstractSlider::focusPreviousChild;
	using QAbstractSlider::sender;
	using QAbstractSlider::senderSignalIndex;
	using QAbstractSlider::receivers;
	using QAbstractSlider::isSignalConnected;
	using QAbstractSlider::getDecodedMetricF;
	bool cppgoBase_event(QEvent* __a0) { return QAbstractSlider::event(__a0); }
	void cppgoBase_sliderChange(int __a0) { QAbstractSlider::sliderChange((QAbstractSlider::SliderChange)__a0); }
	void cppgoBase_keyPressEvent(QKeyEvent* __a0) { QAbstractSlider::keyPressEvent(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QAbstractSlider::timerEvent(__a0); }
	void cppgoBase_wheelEvent(QWheelEvent* __a0) { QAbstractSlider::wheelEvent(__a0); }
	void cppgoBase_changeEvent(QEvent* __a0) { QAbstractSlider::changeEvent(__a0); }
	void cppgoBase_mousePressEvent(QMouseEvent* __a0) { QAbstractSlider::mousePressEvent(__a0); }
	void cppgoBase_mouseReleaseEvent(QMouseEvent* __a0) { QAbstractSlider::mouseReleaseEvent(__a0); }
	void cppgoBase_mouseDoubleClickEvent(QMouseEvent* __a0) { QAbstractSlider::mouseDoubleClickEvent(__a0); }
	void cppgoBase_mouseMoveEvent(QMouseEvent* __a0) { QAbstractSlider::mouseMoveEvent(__a0); }
	void cppgoBase_keyReleaseEvent(QKeyEvent* __a0) { QAbstractSlider::keyReleaseEvent(__a0); }
	void cppgoBase_focusInEvent(QFocusEvent* __a0) { QAbstractSlider::focusInEvent(__a0); }
	void cppgoBase_focusOutEvent(QFocusEvent* __a0) { QAbstractSlider::focusOutEvent(__a0); }
	void cppgoBase_enterEvent(QEnterEvent* __a0) { QAbstractSlider::enterEvent(__a0); }
	void cppgoBase_leaveEvent(QEvent* __a0) { QAbstractSlider::leaveEvent(__a0); }
	void cppgoBase_paintEvent(QPaintEvent* __a0) { QAbstractSlider::paintEvent(__a0); }
	void cppgoBase_moveEvent(QMoveEvent* __a0) { QAbstractSlider::moveEvent(__a0); }
	void cppgoBase_resizeEvent(QResizeEvent* __a0) { QAbstractSlider::resizeEvent(__a0); }
	void cppgoBase_closeEvent(QCloseEvent* __a0) { QAbstractSlider::closeEvent(__a0); }
	void cppgoBase_contextMenuEvent(QContextMenuEvent* __a0) { QAbstractSlider::contextMenuEvent(__a0); }
	void cppgoBase_tabletEvent(QTabletEvent* __a0) { QAbstractSlider::tabletEvent(__a0); }
	void cppgoBase_actionEvent(QActionEvent* __a0) { QAbstractSlider::actionEvent(__a0); }
	void cppgoBase_dragEnterEvent(QDragEnterEvent* __a0) { QAbstractSlider::dragEnterEvent(__a0); }
	void cppgoBase_dragMoveEvent(QDragMoveEvent* __a0) { QAbstractSlider::dragMoveEvent(__a0); }
	void cppgoBase_dragLeaveEvent(QDragLeaveEvent* __a0) { QAbstractSlider::dragLeaveEvent(__a0); }
	void cppgoBase_dropEvent(QDropEvent* __a0) { QAbstractSlider::dropEvent(__a0); }
	void cppgoBase_showEvent(QShowEvent* __a0) { QAbstractSlider::showEvent(__a0); }
	void cppgoBase_hideEvent(QHideEvent* __a0) { QAbstractSlider::hideEvent(__a0); }
	int cppgoBase_metric(int __a0) { return QAbstractSlider::metric((QPaintDevice::PaintDeviceMetric)__a0); }
	void cppgoBase_initPainter(QPainter* __a0) { QAbstractSlider::initPainter(__a0); }
	QPaintDevice* cppgoBase_redirected(QPoint* __a0) { return QAbstractSlider::redirected(__a0); }
	QPainter* cppgoBase_sharedPainter() { return QAbstractSlider::sharedPainter(); }
	void cppgoBase_inputMethodEvent(QInputMethodEvent* __a0) { QAbstractSlider::inputMethodEvent(__a0); }
	bool cppgoBase_focusNextPrevChild(bool __a0) { return QAbstractSlider::focusNextPrevChild(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QAbstractSlider::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QAbstractSlider::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QAbstractSlider::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QAbstractSlider::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QAbstractSlider::metaObject();
	}
	bool event(QEvent* e) override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)e, &__h); if (__h) return __r; }
		return QAbstractSlider::event(e);
	}
	void sliderChange(QAbstractSlider::SliderChange change) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 4658628703300230988LL, (int)(change), &__h); if (__h) return; }
		QAbstractSlider::sliderChange(change);
	}
	void keyPressEvent(QKeyEvent* ev) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1023791977679793185LL, (void*)ev, &__h); if (__h) return; }
		QAbstractSlider::keyPressEvent(ev);
	}
	void timerEvent(QTimerEvent* arg1) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)arg1, &__h); if (__h) return; }
		QAbstractSlider::timerEvent(arg1);
	}
	void wheelEvent(QWheelEvent* e) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1919988340934511316LL, (void*)e, &__h); if (__h) return; }
		QAbstractSlider::wheelEvent(e);
	}
	void changeEvent(QEvent* e) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4959376063826165943LL, (void*)e, &__h); if (__h) return; }
		QAbstractSlider::changeEvent(e);
	}
	int devType() const override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8659679893868524044LL, &__h); if (__h) return __r; }
		return QAbstractSlider::devType();
	}
	void setVisible(bool visible) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 2110182719375245575LL, visible, &__h); if (__h) return; }
		QAbstractSlider::setVisible(visible);
	}
	QSize sizeHint() const override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 729690261457684101LL, &__h); if (__h) return *(QSize*)__r; }
		return QAbstractSlider::sizeHint();
	}
	QSize minimumSizeHint() const override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 7049264599837532491LL, &__h); if (__h) return *(QSize*)__r; }
		return QAbstractSlider::minimumSizeHint();
	}
	int heightForWidth(int arg1) const override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 4311534060495737339LL, arg1, &__h); if (__h) return __r; }
		return QAbstractSlider::heightForWidth(arg1);
	}
	bool hasHeightForWidth() const override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8266762397995423915LL, &__h); if (__h) return __r; }
		return QAbstractSlider::hasHeightForWidth();
	}
	QPaintEngine* paintEngine() const override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 6398007404756346381LL, &__h); if (__h) return (QPaintEngine*)__r; }
		return QAbstractSlider::paintEngine();
	}
	void mousePressEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -332090898078543319LL, (void*)event, &__h); if (__h) return; }
		QAbstractSlider::mousePressEvent(event);
	}
	void mouseReleaseEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4415920049297795213LL, (void*)event, &__h); if (__h) return; }
		QAbstractSlider::mouseReleaseEvent(event);
	}
	void mouseDoubleClickEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4029073905697023887LL, (void*)event, &__h); if (__h) return; }
		QAbstractSlider::mouseDoubleClickEvent(event);
	}
	void mouseMoveEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4255494004099806581LL, (void*)event, &__h); if (__h) return; }
		QAbstractSlider::mouseMoveEvent(event);
	}
	void keyReleaseEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 772598475950050299LL, (void*)event, &__h); if (__h) return; }
		QAbstractSlider::keyReleaseEvent(event);
	}
	void focusInEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 975043767033552552LL, (void*)event, &__h); if (__h) return; }
		QAbstractSlider::focusInEvent(event);
	}
	void focusOutEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -385596010576048993LL, (void*)event, &__h); if (__h) return; }
		QAbstractSlider::focusOutEvent(event);
	}
	void enterEvent(QEnterEvent* event) override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6316270346668237895LL, (void*)event, &__h); if (__h) return; }
		QAbstractSlider::enterEvent(event);
	}
	void leaveEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -551070009623711640LL, (void*)event, &__h); if (__h) return; }
		QAbstractSlider::leaveEvent(event);
	}
	void paintEvent(QPaintEvent* event) override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5290442163167458465LL, (void*)event, &__h); if (__h) return; }
		QAbstractSlider::paintEvent(event);
	}
	void moveEvent(QMoveEvent* event) override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -9161949150489525396LL, (void*)event, &__h); if (__h) return; }
		QAbstractSlider::moveEvent(event);
	}
	void resizeEvent(QResizeEvent* event) override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4302917285655014661LL, (void*)event, &__h); if (__h) return; }
		QAbstractSlider::resizeEvent(event);
	}
	void closeEvent(QCloseEvent* event) override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1989937050019464783LL, (void*)event, &__h); if (__h) return; }
		QAbstractSlider::closeEvent(event);
	}
	void contextMenuEvent(QContextMenuEvent* event) override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6698043222227006015LL, (void*)event, &__h); if (__h) return; }
		QAbstractSlider::contextMenuEvent(event);
	}
	void tabletEvent(QTabletEvent* event) override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -212068866860701141LL, (void*)event, &__h); if (__h) return; }
		QAbstractSlider::tabletEvent(event);
	}
	void actionEvent(QActionEvent* event) override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1697444140088555581LL, (void*)event, &__h); if (__h) return; }
		QAbstractSlider::actionEvent(event);
	}
	void dragEnterEvent(QDragEnterEvent* event) override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 206122715919917589LL, (void*)event, &__h); if (__h) return; }
		QAbstractSlider::dragEnterEvent(event);
	}
	void dragMoveEvent(QDragMoveEvent* event) override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1781545137352745978LL, (void*)event, &__h); if (__h) return; }
		QAbstractSlider::dragMoveEvent(event);
	}
	void dragLeaveEvent(QDragLeaveEvent* event) override {
		int __h = 0;
		if (__goOn[32] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -3139904158373725006LL, (void*)event, &__h); if (__h) return; }
		QAbstractSlider::dragLeaveEvent(event);
	}
	void dropEvent(QDropEvent* event) override {
		int __h = 0;
		if (__goOn[33] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 9107058682754591716LL, (void*)event, &__h); if (__h) return; }
		QAbstractSlider::dropEvent(event);
	}
	void showEvent(QShowEvent* event) override {
		int __h = 0;
		if (__goOn[34] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6825369174751916LL, (void*)event, &__h); if (__h) return; }
		QAbstractSlider::showEvent(event);
	}
	void hideEvent(QHideEvent* event) override {
		int __h = 0;
		if (__goOn[35] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -199718707876719129LL, (void*)event, &__h); if (__h) return; }
		QAbstractSlider::hideEvent(event);
	}
	bool nativeEvent(const QByteArray& eventType, void* message, qintptr* result) override {
		int __h = 0;
		if (__goOn[36] && cppgo_virt_int_voidp_voidp_longlongp) { int __r = ((int(*)(long long, long long, void*, void*, long long*, int*))cppgo_virt_int_voidp_voidp_longlongp)(__goId, -2039166586246443706LL, (void*)&eventType, (void*)message, (long long*)result, &__h); if (__h) return __r; }
		return QAbstractSlider::nativeEvent(eventType, message, result);
	}
	int metric(QPaintDevice::PaintDeviceMetric arg1) const override {
		int __h = 0;
		if (__goOn[37] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 427629649267311777LL, (int)(arg1), &__h); if (__h) return __r; }
		return QAbstractSlider::metric(arg1);
	}
	void initPainter(QPainter* painter) const override {
		int __h = 0;
		if (__goOn[38] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6477500296980050210LL, (void*)painter, &__h); if (__h) return; }
		QAbstractSlider::initPainter(painter);
	}
	QPaintDevice* redirected(QPoint* offset) const override {
		int __h = 0;
		if (__goOn[39] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 4127240633310655048LL, (void*)offset, &__h); if (__h) return (QPaintDevice*)__r; }
		return QAbstractSlider::redirected(offset);
	}
	QPainter* sharedPainter() const override {
		int __h = 0;
		if (__goOn[40] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 5792771328741774991LL, &__h); if (__h) return (QPainter*)__r; }
		return QAbstractSlider::sharedPainter();
	}
	void inputMethodEvent(QInputMethodEvent* arg1) override {
		int __h = 0;
		if (__goOn[41] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5438819661942843222LL, (void*)arg1, &__h); if (__h) return; }
		QAbstractSlider::inputMethodEvent(arg1);
	}
	QVariant inputMethodQuery(Qt::InputMethodQuery arg1) const override {
		int __h = 0;
		if (__goOn[42] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7345853097340709638LL, (int)(arg1), &__h); if (__h) return *(QVariant*)__r; }
		return QAbstractSlider::inputMethodQuery(arg1);
	}
	bool focusNextPrevChild(bool next) override {
		int __h = 0;
		if (__goOn[43] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 5898328517505784191LL, next, &__h); if (__h) return __r; }
		return QAbstractSlider::focusNextPrevChild(next);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[44] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QAbstractSlider::eventFilter(watched, event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[45] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QAbstractSlider::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[46] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QAbstractSlider::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[47] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QAbstractSlider::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[48] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QAbstractSlider::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QAbstractSlider;
static std::mutex __connsMu_QAbstractSlider;

extern "C" {

void* QAbstractSlider_base(void* self) { return (void*)static_cast<QWidget*>((QAbstractSlider*)self); }
void QAbstractSlider_chain(void* self, void** out) {
	QAbstractSlider* __p = (QAbstractSlider*)self;
	QWidget* __b0 = static_cast<QWidget*>(__p); out[0] = (void*)__b0;
	QObject* __b1 = static_cast<QObject*>(__b0); out[1] = (void*)__b1;
}
void QAbstractSlider__set_override(void* self, int idx) { ((QAbstractSlider__Go*)self)->__goOn[idx] = 1; }
void* QAbstractSlider_fromQObject(void* o) { return (void*)qobject_cast<QAbstractSlider*>((QObject*)o); }
void* QAbstractSlider_new(long long __goid, void* parent) {
	return new QAbstractSlider__Go(__goid, (QWidget*)parent);
}
void* QAbstractSlider_new2(long long __goid) {
	return new QAbstractSlider__Go(__goid);
}
char* QAbstractSlider_tr(const char* s, long long* __out_len) {
	auto&& __r = QAbstractSlider::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QAbstractSlider_orientation(void* self) {
	auto&& __r = ((QAbstractSlider*)self)->orientation();
	return (int)(__r);
}
void QAbstractSlider_setMinimum(void* self, int arg1) {
	((QAbstractSlider*)self)->setMinimum(arg1);
}
int QAbstractSlider_minimum(void* self) {
	auto&& __r = ((QAbstractSlider*)self)->minimum();
	return __r;
}
void QAbstractSlider_setMaximum(void* self, int arg1) {
	((QAbstractSlider*)self)->setMaximum(arg1);
}
int QAbstractSlider_maximum(void* self) {
	auto&& __r = ((QAbstractSlider*)self)->maximum();
	return __r;
}
void QAbstractSlider_setSingleStep(void* self, int arg1) {
	((QAbstractSlider*)self)->setSingleStep(arg1);
}
int QAbstractSlider_singleStep(void* self) {
	auto&& __r = ((QAbstractSlider*)self)->singleStep();
	return __r;
}
void QAbstractSlider_setPageStep(void* self, int arg1) {
	((QAbstractSlider*)self)->setPageStep(arg1);
}
int QAbstractSlider_pageStep(void* self) {
	auto&& __r = ((QAbstractSlider*)self)->pageStep();
	return __r;
}
void QAbstractSlider_setTracking(void* self, int enable) {
	((QAbstractSlider*)self)->setTracking(enable);
}
int QAbstractSlider_hasTracking(void* self) {
	auto&& __r = ((QAbstractSlider*)self)->hasTracking();
	return __r;
}
void QAbstractSlider_setSliderDown(void* self, int arg1) {
	((QAbstractSlider*)self)->setSliderDown(arg1);
}
int QAbstractSlider_isSliderDown(void* self) {
	auto&& __r = ((QAbstractSlider*)self)->isSliderDown();
	return __r;
}
void QAbstractSlider_setSliderPosition(void* self, int arg1) {
	((QAbstractSlider*)self)->setSliderPosition(arg1);
}
int QAbstractSlider_sliderPosition(void* self) {
	auto&& __r = ((QAbstractSlider*)self)->sliderPosition();
	return __r;
}
void QAbstractSlider_setInvertedAppearance(void* self, int arg1) {
	((QAbstractSlider*)self)->setInvertedAppearance(arg1);
}
int QAbstractSlider_invertedAppearance(void* self) {
	auto&& __r = ((QAbstractSlider*)self)->invertedAppearance();
	return __r;
}
void QAbstractSlider_setInvertedControls(void* self, int arg1) {
	((QAbstractSlider*)self)->setInvertedControls(arg1);
}
int QAbstractSlider_invertedControls(void* self) {
	auto&& __r = ((QAbstractSlider*)self)->invertedControls();
	return __r;
}
int QAbstractSlider_value(void* self) {
	auto&& __r = ((QAbstractSlider*)self)->value();
	return __r;
}
void QAbstractSlider_triggerAction(void* self, int action) {
	((QAbstractSlider*)self)->triggerAction((QAbstractSlider::SliderAction)(action));
}
void QAbstractSlider_setValue(void* self, int arg1) {
	((QAbstractSlider*)self)->setValue(arg1);
}
void QAbstractSlider_setOrientation(void* self, int arg1) {
	((QAbstractSlider*)self)->setOrientation((Qt::Orientation)(arg1));
}
void QAbstractSlider_setRange(void* self, int min, int max) {
	((QAbstractSlider*)self)->setRange(min, max);
}
void QAbstractSlider_valueChanged(void* self, int value) {
	((QAbstractSlider*)self)->valueChanged(value);
}
void QAbstractSlider_sliderPressed(void* self) {
	((QAbstractSlider*)self)->sliderPressed();
}
void QAbstractSlider_sliderMoved(void* self, int position) {
	((QAbstractSlider*)self)->sliderMoved(position);
}
void QAbstractSlider_sliderReleased(void* self) {
	((QAbstractSlider*)self)->sliderReleased();
}
void QAbstractSlider_rangeChanged(void* self, int min, int max) {
	((QAbstractSlider*)self)->rangeChanged(min, max);
}
void QAbstractSlider_actionTriggered(void* self, int action) {
	((QAbstractSlider*)self)->actionTriggered(action);
}
void QAbstractSlider_setRepeatAction(void* self, int action) {
	((QAbstractSlider__Go*)self)->setRepeatAction((QAbstractSlider::SliderAction)(action));
}
int QAbstractSlider_repeatAction(void* self) {
	auto&& __r = ((QAbstractSlider__Go*)self)->repeatAction();
	return (int)(__r);
}
void QAbstractSlider_setTabOrder(void* arg1, void* arg2) {
	QAbstractSlider::setTabOrder((QWidget*)arg1, (QWidget*)arg2);
}
void* QAbstractSlider_mouseGrabber() {
	auto&& __r = QAbstractSlider::mouseGrabber();
	return (void*)(__r);
}
void* QAbstractSlider_keyboardGrabber() {
	auto&& __r = QAbstractSlider::keyboardGrabber();
	return (void*)(__r);
}
void* QAbstractSlider_find(unsigned long long arg1) {
	auto&& __r = QAbstractSlider::find(arg1);
	return (void*)(__r);
}
void* QAbstractSlider_createWindowContainer(void* window) {
	auto&& __r = QAbstractSlider::createWindowContainer((QWindow*)window);
	return (void*)(__r);
}
void QAbstractSlider_updateMicroFocus(void* self) {
	((QAbstractSlider__Go*)self)->updateMicroFocus();
}
void QAbstractSlider_create(void* self) {
	((QAbstractSlider__Go*)self)->create();
}
void QAbstractSlider_destroy(void* self) {
	((QAbstractSlider__Go*)self)->destroy();
}
int QAbstractSlider_focusNextChild(void* self) {
	auto&& __r = ((QAbstractSlider__Go*)self)->focusNextChild();
	return __r;
}
int QAbstractSlider_focusPreviousChild(void* self) {
	auto&& __r = ((QAbstractSlider__Go*)self)->focusPreviousChild();
	return __r;
}
void* QAbstractSlider_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QAbstractSlider::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QAbstractSlider_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QAbstractSlider::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QAbstractSlider_sender(void* self) {
	auto&& __r = ((QAbstractSlider__Go*)self)->sender();
	return (void*)(__r);
}
int QAbstractSlider_senderSignalIndex(void* self) {
	auto&& __r = ((QAbstractSlider__Go*)self)->senderSignalIndex();
	return __r;
}
int QAbstractSlider_receivers(void* self, const char* signal) {
	auto&& __r = ((QAbstractSlider__Go*)self)->receivers(signal);
	return __r;
}
int QAbstractSlider_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QAbstractSlider__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
double QAbstractSlider_devicePixelRatioFScale() {
	auto&& __r = QAbstractSlider::devicePixelRatioFScale();
	return __r;
}
int QAbstractSlider_encodeMetricF(int metric, double value) {
	auto&& __r = QAbstractSlider::encodeMetricF((QPaintDevice::PaintDeviceMetric)(metric), value);
	return __r;
}
double QAbstractSlider_getDecodedMetricF(void* self, int metricA, int metricB) {
	auto&& __r = ((QAbstractSlider__Go*)self)->getDecodedMetricF((QPaintDevice::PaintDeviceMetric)(metricA), (QPaintDevice::PaintDeviceMetric)(metricB));
	return __r;
}
char* QAbstractSlider_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QAbstractSlider::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QAbstractSlider_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QAbstractSlider::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QAbstractSlider_setRepeatAction2(void* self, int action, int thresholdTime) {
	((QAbstractSlider__Go*)self)->setRepeatAction((QAbstractSlider::SliderAction)(action), thresholdTime);
}
void QAbstractSlider_setRepeatAction3(void* self, int action, int thresholdTime, int repeatTime) {
	((QAbstractSlider__Go*)self)->setRepeatAction((QAbstractSlider::SliderAction)(action), thresholdTime, repeatTime);
}
void* QAbstractSlider_createWindowContainer2(void* window, void* parent) {
	auto&& __r = QAbstractSlider::createWindowContainer((QWindow*)window, (QWidget*)parent);
	return (void*)(__r);
}
void* QAbstractSlider_createWindowContainer3(void* window, void* parent, int flags) {
	auto&& __r = QAbstractSlider::createWindowContainer((QWindow*)window, (QWidget*)parent, QFlags<Qt::WindowType>(flags));
	return (void*)(__r);
}
void QAbstractSlider_updateMicroFocus2(void* self, int query) {
	((QAbstractSlider__Go*)self)->updateMicroFocus((Qt::InputMethodQuery)(query));
}
void QAbstractSlider_create2(void* self, unsigned long long arg1) {
	((QAbstractSlider__Go*)self)->create(arg1);
}
void QAbstractSlider_create3(void* self, unsigned long long arg1, int initializeWindow) {
	((QAbstractSlider__Go*)self)->create(arg1, initializeWindow);
}
void QAbstractSlider_create4(void* self, unsigned long long arg1, int initializeWindow, int destroyOldWindow) {
	((QAbstractSlider__Go*)self)->create(arg1, initializeWindow, destroyOldWindow);
}
void QAbstractSlider_destroy2(void* self, int destroyWindow) {
	((QAbstractSlider__Go*)self)->destroy(destroyWindow);
}
void QAbstractSlider_destroy3(void* self, int destroyWindow, int destroySubWindows) {
	((QAbstractSlider__Go*)self)->destroy(destroyWindow, destroySubWindows);
}
void* QAbstractSlider_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QAbstractSlider::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QAbstractSlider_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QAbstractSlider::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QAbstractSlider_metaObject_default(void* self) {
	auto&& __r = ((QAbstractSlider*)self)->QAbstractSlider::metaObject();
	return (void*)(__r);
}
int QAbstractSlider_event_default(void* self, void* e) {
	auto&& __r = ((QAbstractSlider__Go*)self)->cppgoBase_event((QEvent*)e);
	return __r;
}
void QAbstractSlider_sliderChange_default(void* self, int change) {
	((QAbstractSlider__Go*)self)->cppgoBase_sliderChange(change);
}
void QAbstractSlider_keyPressEvent_default(void* self, void* ev) {
	((QAbstractSlider__Go*)self)->cppgoBase_keyPressEvent((QKeyEvent*)ev);
}
void QAbstractSlider_timerEvent_default(void* self, void* arg1) {
	((QAbstractSlider__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)arg1);
}
void QAbstractSlider_wheelEvent_default(void* self, void* e) {
	((QAbstractSlider__Go*)self)->cppgoBase_wheelEvent((QWheelEvent*)e);
}
void QAbstractSlider_changeEvent_default(void* self, void* e) {
	((QAbstractSlider__Go*)self)->cppgoBase_changeEvent((QEvent*)e);
}
int QAbstractSlider_devType_default(void* self) {
	auto&& __r = ((QAbstractSlider*)self)->QAbstractSlider::devType();
	return __r;
}
void QAbstractSlider_setVisible_default(void* self, int visible) {
	((QAbstractSlider*)self)->QAbstractSlider::setVisible(visible);
}
void* QAbstractSlider_sizeHint_default(void* self) {
	auto&& __r = ((QAbstractSlider*)self)->QAbstractSlider::sizeHint();
	return new QSize(__r);
}
void* QAbstractSlider_minimumSizeHint_default(void* self) {
	auto&& __r = ((QAbstractSlider*)self)->QAbstractSlider::minimumSizeHint();
	return new QSize(__r);
}
int QAbstractSlider_heightForWidth_default(void* self, int arg1) {
	auto&& __r = ((QAbstractSlider*)self)->QAbstractSlider::heightForWidth(arg1);
	return __r;
}
int QAbstractSlider_hasHeightForWidth_default(void* self) {
	auto&& __r = ((QAbstractSlider*)self)->QAbstractSlider::hasHeightForWidth();
	return __r;
}
void* QAbstractSlider_paintEngine_default(void* self) {
	auto&& __r = ((QAbstractSlider*)self)->QAbstractSlider::paintEngine();
	return (void*)(__r);
}
void QAbstractSlider_mousePressEvent_default(void* self, void* event) {
	((QAbstractSlider__Go*)self)->cppgoBase_mousePressEvent((QMouseEvent*)event);
}
void QAbstractSlider_mouseReleaseEvent_default(void* self, void* event) {
	((QAbstractSlider__Go*)self)->cppgoBase_mouseReleaseEvent((QMouseEvent*)event);
}
void QAbstractSlider_mouseDoubleClickEvent_default(void* self, void* event) {
	((QAbstractSlider__Go*)self)->cppgoBase_mouseDoubleClickEvent((QMouseEvent*)event);
}
void QAbstractSlider_mouseMoveEvent_default(void* self, void* event) {
	((QAbstractSlider__Go*)self)->cppgoBase_mouseMoveEvent((QMouseEvent*)event);
}
void QAbstractSlider_keyReleaseEvent_default(void* self, void* event) {
	((QAbstractSlider__Go*)self)->cppgoBase_keyReleaseEvent((QKeyEvent*)event);
}
void QAbstractSlider_focusInEvent_default(void* self, void* event) {
	((QAbstractSlider__Go*)self)->cppgoBase_focusInEvent((QFocusEvent*)event);
}
void QAbstractSlider_focusOutEvent_default(void* self, void* event) {
	((QAbstractSlider__Go*)self)->cppgoBase_focusOutEvent((QFocusEvent*)event);
}
void QAbstractSlider_enterEvent_default(void* self, void* event) {
	((QAbstractSlider__Go*)self)->cppgoBase_enterEvent((QEnterEvent*)event);
}
void QAbstractSlider_leaveEvent_default(void* self, void* event) {
	((QAbstractSlider__Go*)self)->cppgoBase_leaveEvent((QEvent*)event);
}
void QAbstractSlider_paintEvent_default(void* self, void* event) {
	((QAbstractSlider__Go*)self)->cppgoBase_paintEvent((QPaintEvent*)event);
}
void QAbstractSlider_moveEvent_default(void* self, void* event) {
	((QAbstractSlider__Go*)self)->cppgoBase_moveEvent((QMoveEvent*)event);
}
void QAbstractSlider_resizeEvent_default(void* self, void* event) {
	((QAbstractSlider__Go*)self)->cppgoBase_resizeEvent((QResizeEvent*)event);
}
void QAbstractSlider_closeEvent_default(void* self, void* event) {
	((QAbstractSlider__Go*)self)->cppgoBase_closeEvent((QCloseEvent*)event);
}
void QAbstractSlider_contextMenuEvent_default(void* self, void* event) {
	((QAbstractSlider__Go*)self)->cppgoBase_contextMenuEvent((QContextMenuEvent*)event);
}
void QAbstractSlider_tabletEvent_default(void* self, void* event) {
	((QAbstractSlider__Go*)self)->cppgoBase_tabletEvent((QTabletEvent*)event);
}
void QAbstractSlider_actionEvent_default(void* self, void* event) {
	((QAbstractSlider__Go*)self)->cppgoBase_actionEvent((QActionEvent*)event);
}
void QAbstractSlider_dragEnterEvent_default(void* self, void* event) {
	((QAbstractSlider__Go*)self)->cppgoBase_dragEnterEvent((QDragEnterEvent*)event);
}
void QAbstractSlider_dragMoveEvent_default(void* self, void* event) {
	((QAbstractSlider__Go*)self)->cppgoBase_dragMoveEvent((QDragMoveEvent*)event);
}
void QAbstractSlider_dragLeaveEvent_default(void* self, void* event) {
	((QAbstractSlider__Go*)self)->cppgoBase_dragLeaveEvent((QDragLeaveEvent*)event);
}
void QAbstractSlider_dropEvent_default(void* self, void* event) {
	((QAbstractSlider__Go*)self)->cppgoBase_dropEvent((QDropEvent*)event);
}
void QAbstractSlider_showEvent_default(void* self, void* event) {
	((QAbstractSlider__Go*)self)->cppgoBase_showEvent((QShowEvent*)event);
}
void QAbstractSlider_hideEvent_default(void* self, void* event) {
	((QAbstractSlider__Go*)self)->cppgoBase_hideEvent((QHideEvent*)event);
}
int QAbstractSlider_metric_default(void* self, int arg1) {
	auto&& __r = ((QAbstractSlider__Go*)self)->cppgoBase_metric(arg1);
	return __r;
}
void QAbstractSlider_initPainter_default(void* self, void* painter) {
	((QAbstractSlider__Go*)self)->cppgoBase_initPainter((QPainter*)painter);
}
void* QAbstractSlider_redirected_default(void* self, void* offset) {
	auto&& __r = ((QAbstractSlider__Go*)self)->cppgoBase_redirected((QPoint*)offset);
	return (void*)(__r);
}
void* QAbstractSlider_sharedPainter_default(void* self) {
	auto&& __r = ((QAbstractSlider__Go*)self)->cppgoBase_sharedPainter();
	return (void*)(__r);
}
void QAbstractSlider_inputMethodEvent_default(void* self, void* arg1) {
	((QAbstractSlider__Go*)self)->cppgoBase_inputMethodEvent((QInputMethodEvent*)arg1);
}
void* QAbstractSlider_inputMethodQuery_default(void* self, int arg1) {
	auto&& __r = ((QAbstractSlider*)self)->QAbstractSlider::inputMethodQuery((Qt::InputMethodQuery)(arg1));
	return new QVariant(__r);
}
int QAbstractSlider_focusNextPrevChild_default(void* self, int next) {
	auto&& __r = ((QAbstractSlider__Go*)self)->cppgoBase_focusNextPrevChild(next);
	return __r;
}
int QAbstractSlider_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QAbstractSlider*)self)->QAbstractSlider::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QAbstractSlider_childEvent_default(void* self, void* event) {
	((QAbstractSlider__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QAbstractSlider_customEvent_default(void* self, void* event) {
	((QAbstractSlider__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QAbstractSlider_connectNotify_default(void* self, void* signal) {
	((QAbstractSlider__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QAbstractSlider_disconnectNotify_default(void* self, void* signal) {
	((QAbstractSlider__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QAbstractSlider_connect_valueChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractSlider*)self, &QAbstractSlider::valueChanged, (QAbstractSlider*)self, [__id](int value){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, value);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractSlider);
	__conns_QAbstractSlider[__id] = __c;
}
void QAbstractSlider_connect_sliderPressed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractSlider*)self, &QAbstractSlider::sliderPressed, (QAbstractSlider*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractSlider);
	__conns_QAbstractSlider[__id] = __c;
}
void QAbstractSlider_connect_sliderMoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractSlider*)self, &QAbstractSlider::sliderMoved, (QAbstractSlider*)self, [__id](int position){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, position);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractSlider);
	__conns_QAbstractSlider[__id] = __c;
}
void QAbstractSlider_connect_sliderReleased(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractSlider*)self, &QAbstractSlider::sliderReleased, (QAbstractSlider*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractSlider);
	__conns_QAbstractSlider[__id] = __c;
}
void QAbstractSlider_connect_rangeChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractSlider*)self, &QAbstractSlider::rangeChanged, (QAbstractSlider*)self, [__id](int min, int max){
		if (cppgo_sig_void_int_int) ((void(*)(long long, int, int))cppgo_sig_void_int_int)(__id, min, max);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractSlider);
	__conns_QAbstractSlider[__id] = __c;
}
void QAbstractSlider_connect_actionTriggered(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractSlider*)self, &QAbstractSlider::actionTriggered, (QAbstractSlider*)self, [__id](int action){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, action);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractSlider);
	__conns_QAbstractSlider[__id] = __c;
}
void QAbstractSlider_connect_windowTitleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractSlider*)self, &QAbstractSlider::windowTitleChanged, (QAbstractSlider*)self, [__id](const QString& title){
		QByteArray title_ba = title.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)title_ba.constData(), title_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractSlider);
	__conns_QAbstractSlider[__id] = __c;
}
void QAbstractSlider_connect_windowIconChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractSlider*)self, &QAbstractSlider::windowIconChanged, (QAbstractSlider*)self, [__id](const QIcon& icon){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&icon);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractSlider);
	__conns_QAbstractSlider[__id] = __c;
}
void QAbstractSlider_connect_windowIconTextChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractSlider*)self, &QAbstractSlider::windowIconTextChanged, (QAbstractSlider*)self, [__id](const QString& iconText){
		QByteArray iconText_ba = iconText.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)iconText_ba.constData(), iconText_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractSlider);
	__conns_QAbstractSlider[__id] = __c;
}
void QAbstractSlider_connect_customContextMenuRequested(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractSlider*)self, &QAbstractSlider::customContextMenuRequested, (QAbstractSlider*)self, [__id](const QPoint& pos){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&pos);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractSlider);
	__conns_QAbstractSlider[__id] = __c;
}
void QAbstractSlider_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractSlider*)self, &QAbstractSlider::destroyed, (QAbstractSlider*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractSlider);
	__conns_QAbstractSlider[__id] = __c;
}
void QAbstractSlider_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractSlider*)self, &QAbstractSlider::destroyed, (QAbstractSlider*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractSlider);
	__conns_QAbstractSlider[__id] = __c;
}
void QAbstractSlider_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractSlider*)self, &QAbstractSlider::objectNameChanged, (QAbstractSlider*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractSlider);
	__conns_QAbstractSlider[__id] = __c;
}
void QAbstractSlider__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QAbstractSlider);
		auto __it = __conns_QAbstractSlider.find(id);
		if (__it == __conns_QAbstractSlider.end()) return;
		__c = __it->second;
		__conns_QAbstractSlider.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QAbstractSlider__Go_new(long long id) { return new QAbstractSlider__Go(id); }
void QAbstractSlider_delete(void* self) { delete (QAbstractSlider*)self; }

}

#include <QtWidgets/qstyleoption.h>
#include <QtWidgets/qwidget.h>
#include <QtWidgets/qstyle.h>
#include "gen_qstyleoptioncomplex.h"

extern "C" {

void* QStyleOptionComplex_base(void* self) { return (void*)static_cast<QStyleOption*>((QStyleOptionComplex*)self); }
void* QStyleOptionComplex_new() {
	return new QStyleOptionComplex();
}
void* QStyleOptionComplex_new2(void* other) {
	return new QStyleOptionComplex(*(QStyleOptionComplex*)other);
}
void* QStyleOptionComplex_new3(int version) {
	return new QStyleOptionComplex(version);
}
void* QStyleOptionComplex_new4(int version, int type_) {
	return new QStyleOptionComplex(version, type_);
}
void* QStyleOptionComplex_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QStyleOptionComplex*)self)->operator=(*(QStyleOptionComplex*)arg1);
	return (void*)&(__r);
}
int QStyleOptionComplex_subControls_get(void* self) {
	auto __r = ((QStyleOptionComplex*)self)->subControls;
	return (int)(__r);
}
void QStyleOptionComplex_subControls_set(void* self, int subControls) {
	((QStyleOptionComplex*)self)->subControls = QFlags<QStyle::SubControl>(subControls);
}
int QStyleOptionComplex_activeSubControls_get(void* self) {
	auto __r = ((QStyleOptionComplex*)self)->activeSubControls;
	return (int)(__r);
}
void QStyleOptionComplex_activeSubControls_set(void* self, int activeSubControls) {
	((QStyleOptionComplex*)self)->activeSubControls = QFlags<QStyle::SubControl>(activeSubControls);
}
void QStyleOptionComplex_delete(void* self) { delete (QStyleOptionComplex*)self; }

}

#include <QtWidgets/qboxlayout.h>
#include <QtWidgets/qlayout.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtWidgets/qlayoutitem.h>
#include <QtCore/qnamespace.h>
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
#include "gen_qboxlayout.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

static std::unordered_map<long long, QMetaObject::Connection> __conns_QBoxLayout;
static std::mutex __connsMu_QBoxLayout;

extern "C" {

void* QBoxLayout_base(void* self) { return (void*)static_cast<QLayout*>((QBoxLayout*)self); }
void QBoxLayout_chain(void* self, void** out) {
	QBoxLayout* __p = (QBoxLayout*)self;
	QLayout* __b0 = static_cast<QLayout*>(__p); out[0] = (void*)__b0;
	QObject* __b1 = static_cast<QObject*>(__b0); out[1] = (void*)__b1;
}
void* QBoxLayout_fromQObject(void* o) { return (void*)qobject_cast<QBoxLayout*>((QObject*)o); }
void* QBoxLayout_new(int arg1) {
	return new QBoxLayout((QBoxLayout::Direction)(arg1));
}
void* QBoxLayout_new2(int arg1, void* parent) {
	return new QBoxLayout((QBoxLayout::Direction)(arg1), (QWidget*)parent);
}
char* QBoxLayout_tr(const char* s, long long* __out_len) {
	auto&& __r = QBoxLayout::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QBoxLayout_direction(void* self) {
	auto&& __r = ((QBoxLayout*)self)->direction();
	return (int)(__r);
}
void QBoxLayout_setDirection(void* self, int arg1) {
	((QBoxLayout*)self)->setDirection((QBoxLayout::Direction)(arg1));
}
void QBoxLayout_addSpacing(void* self, int size) {
	((QBoxLayout*)self)->addSpacing(size);
}
void QBoxLayout_addStretch(void* self) {
	((QBoxLayout*)self)->addStretch();
}
void QBoxLayout_addSpacerItem(void* self, void* spacerItem) {
	((QBoxLayout*)self)->addSpacerItem((QSpacerItem*)spacerItem);
}
void QBoxLayout_addWidget(void* self, void* arg1) {
	((QBoxLayout*)self)->addWidget((QWidget*)arg1);
}
void QBoxLayout_addLayout(void* self, void* layout) {
	((QBoxLayout*)self)->addLayout((QLayout*)layout);
}
void QBoxLayout_addStrut(void* self, int arg1) {
	((QBoxLayout*)self)->addStrut(arg1);
}
void QBoxLayout_insertSpacing(void* self, int index, int size) {
	((QBoxLayout*)self)->insertSpacing(index, size);
}
void QBoxLayout_insertStretch(void* self, int index) {
	((QBoxLayout*)self)->insertStretch(index);
}
void QBoxLayout_insertSpacerItem(void* self, int index, void* spacerItem) {
	((QBoxLayout*)self)->insertSpacerItem(index, (QSpacerItem*)spacerItem);
}
void QBoxLayout_insertWidget(void* self, int index, void* widget) {
	((QBoxLayout*)self)->insertWidget(index, (QWidget*)widget);
}
void QBoxLayout_insertLayout(void* self, int index, void* layout) {
	((QBoxLayout*)self)->insertLayout(index, (QLayout*)layout);
}
void QBoxLayout_insertItem(void* self, int index, void* arg2) {
	((QBoxLayout*)self)->insertItem(index, (QLayoutItem*)arg2);
}
int QBoxLayout_setStretchFactor(void* self, void* w, int stretch) {
	auto&& __r = ((QBoxLayout*)self)->setStretchFactor((QWidget*)w, stretch);
	return __r;
}
int QBoxLayout_setStretchFactor2(void* self, void* l, int stretch) {
	auto&& __r = ((QBoxLayout*)self)->setStretchFactor((QLayout*)l, stretch);
	return __r;
}
void QBoxLayout_setStretch(void* self, int index, int stretch) {
	((QBoxLayout*)self)->setStretch(index, stretch);
}
int QBoxLayout_stretch(void* self, int index) {
	auto&& __r = ((QBoxLayout*)self)->stretch(index);
	return __r;
}
void* QBoxLayout_closestAcceptableSize(void* w, void* s) {
	auto&& __r = QBoxLayout::closestAcceptableSize((QWidget*)w, *(QSize*)s);
	return new QSize(__r);
}
void* QBoxLayout_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QBoxLayout::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QBoxLayout_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QBoxLayout::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
char* QBoxLayout_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QBoxLayout::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QBoxLayout_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QBoxLayout::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QBoxLayout_addStretch2(void* self, int stretch) {
	((QBoxLayout*)self)->addStretch(stretch);
}
void QBoxLayout_addWidget2(void* self, void* arg1, int stretch) {
	((QBoxLayout*)self)->addWidget((QWidget*)arg1, stretch);
}
void QBoxLayout_addWidget3(void* self, void* arg1, int stretch, int alignment) {
	((QBoxLayout*)self)->addWidget((QWidget*)arg1, stretch, QFlags<Qt::AlignmentFlag>(alignment));
}
void QBoxLayout_addLayout2(void* self, void* layout, int stretch) {
	((QBoxLayout*)self)->addLayout((QLayout*)layout, stretch);
}
void QBoxLayout_insertStretch2(void* self, int index, int stretch) {
	((QBoxLayout*)self)->insertStretch(index, stretch);
}
void QBoxLayout_insertWidget2(void* self, int index, void* widget, int stretch) {
	((QBoxLayout*)self)->insertWidget(index, (QWidget*)widget, stretch);
}
void QBoxLayout_insertWidget3(void* self, int index, void* widget, int stretch, int alignment) {
	((QBoxLayout*)self)->insertWidget(index, (QWidget*)widget, stretch, QFlags<Qt::AlignmentFlag>(alignment));
}
void QBoxLayout_insertLayout2(void* self, int index, void* layout, int stretch) {
	((QBoxLayout*)self)->insertLayout(index, (QLayout*)layout, stretch);
}
void* QBoxLayout_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QBoxLayout::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QBoxLayout_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QBoxLayout::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void QBoxLayout_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QBoxLayout*)self, &QBoxLayout::destroyed, (QBoxLayout*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QBoxLayout);
	__conns_QBoxLayout[__id] = __c;
}
void QBoxLayout_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QBoxLayout*)self, &QBoxLayout::destroyed, (QBoxLayout*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QBoxLayout);
	__conns_QBoxLayout[__id] = __c;
}
void QBoxLayout_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QBoxLayout*)self, &QBoxLayout::objectNameChanged, (QBoxLayout*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QBoxLayout);
	__conns_QBoxLayout[__id] = __c;
}
void QBoxLayout__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QBoxLayout);
		auto __it = __conns_QBoxLayout.find(id);
		if (__it == __conns_QBoxLayout.end()) return;
		__c = __it->second;
		__conns_QBoxLayout.erase(__it);
	}
	QObject::disconnect(__c);
}
void QBoxLayout_delete(void* self) { delete (QBoxLayout*)self; }

}

#include <QtWidgets/qfiledialog.h>
#include <QtWidgets/qdialog.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qstring.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qdir.h>
#include <QtCore/qurl.h>
#include <QtCore/qbytearray.h>
#include <QtWidgets/qabstractitemdelegate.h>
#include <QtGui/qabstractfileiconprovider.h>
#include <QtCore/qabstractproxymodel.h>
#include <QtCore/qobject.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qsize.h>
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
#include "gen_qfiledialog.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QFileDialog__Go : public QFileDialog {
public:
	long long __goId;
	unsigned char __goOn[52] = {};
	template<typename... A> QFileDialog__Go(long long id, A&&... a) : QFileDialog(std::forward<A>(a)...), __goId(id) {}
	using QFileDialog::adjustPosition;
	using QFileDialog::updateMicroFocus;
	using QFileDialog::create;
	using QFileDialog::destroy;
	using QFileDialog::focusNextChild;
	using QFileDialog::focusPreviousChild;
	using QFileDialog::sender;
	using QFileDialog::senderSignalIndex;
	using QFileDialog::receivers;
	using QFileDialog::isSignalConnected;
	using QFileDialog::getDecodedMetricF;
	void cppgoBase_done(int __a0) { QFileDialog::done(__a0); }
	void cppgoBase_accept() { QFileDialog::accept(); }
	void cppgoBase_changeEvent(QEvent* __a0) { QFileDialog::changeEvent(__a0); }
	void cppgoBase_keyPressEvent(QKeyEvent* __a0) { QFileDialog::keyPressEvent(__a0); }
	void cppgoBase_closeEvent(QCloseEvent* __a0) { QFileDialog::closeEvent(__a0); }
	void cppgoBase_showEvent(QShowEvent* __a0) { QFileDialog::showEvent(__a0); }
	void cppgoBase_resizeEvent(QResizeEvent* __a0) { QFileDialog::resizeEvent(__a0); }
	void cppgoBase_contextMenuEvent(QContextMenuEvent* __a0) { QFileDialog::contextMenuEvent(__a0); }
	bool cppgoBase_eventFilter(QObject* __a0, QEvent* __a1) { return QFileDialog::eventFilter(__a0, __a1); }
	bool cppgoBase_event(QEvent* __a0) { return QFileDialog::event(__a0); }
	void cppgoBase_mousePressEvent(QMouseEvent* __a0) { QFileDialog::mousePressEvent(__a0); }
	void cppgoBase_mouseReleaseEvent(QMouseEvent* __a0) { QFileDialog::mouseReleaseEvent(__a0); }
	void cppgoBase_mouseDoubleClickEvent(QMouseEvent* __a0) { QFileDialog::mouseDoubleClickEvent(__a0); }
	void cppgoBase_mouseMoveEvent(QMouseEvent* __a0) { QFileDialog::mouseMoveEvent(__a0); }
	void cppgoBase_wheelEvent(QWheelEvent* __a0) { QFileDialog::wheelEvent(__a0); }
	void cppgoBase_keyReleaseEvent(QKeyEvent* __a0) { QFileDialog::keyReleaseEvent(__a0); }
	void cppgoBase_focusInEvent(QFocusEvent* __a0) { QFileDialog::focusInEvent(__a0); }
	void cppgoBase_focusOutEvent(QFocusEvent* __a0) { QFileDialog::focusOutEvent(__a0); }
	void cppgoBase_enterEvent(QEnterEvent* __a0) { QFileDialog::enterEvent(__a0); }
	void cppgoBase_leaveEvent(QEvent* __a0) { QFileDialog::leaveEvent(__a0); }
	void cppgoBase_paintEvent(QPaintEvent* __a0) { QFileDialog::paintEvent(__a0); }
	void cppgoBase_moveEvent(QMoveEvent* __a0) { QFileDialog::moveEvent(__a0); }
	void cppgoBase_tabletEvent(QTabletEvent* __a0) { QFileDialog::tabletEvent(__a0); }
	void cppgoBase_actionEvent(QActionEvent* __a0) { QFileDialog::actionEvent(__a0); }
	void cppgoBase_dragEnterEvent(QDragEnterEvent* __a0) { QFileDialog::dragEnterEvent(__a0); }
	void cppgoBase_dragMoveEvent(QDragMoveEvent* __a0) { QFileDialog::dragMoveEvent(__a0); }
	void cppgoBase_dragLeaveEvent(QDragLeaveEvent* __a0) { QFileDialog::dragLeaveEvent(__a0); }
	void cppgoBase_dropEvent(QDropEvent* __a0) { QFileDialog::dropEvent(__a0); }
	void cppgoBase_hideEvent(QHideEvent* __a0) { QFileDialog::hideEvent(__a0); }
	int cppgoBase_metric(int __a0) { return QFileDialog::metric((QPaintDevice::PaintDeviceMetric)__a0); }
	void cppgoBase_initPainter(QPainter* __a0) { QFileDialog::initPainter(__a0); }
	QPaintDevice* cppgoBase_redirected(QPoint* __a0) { return QFileDialog::redirected(__a0); }
	QPainter* cppgoBase_sharedPainter() { return QFileDialog::sharedPainter(); }
	void cppgoBase_inputMethodEvent(QInputMethodEvent* __a0) { QFileDialog::inputMethodEvent(__a0); }
	bool cppgoBase_focusNextPrevChild(bool __a0) { return QFileDialog::focusNextPrevChild(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QFileDialog::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QFileDialog::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QFileDialog::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QFileDialog::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QFileDialog::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QFileDialog::metaObject();
	}
	void setVisible(bool visible) override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 2110182719375245575LL, visible, &__h); if (__h) return; }
		QFileDialog::setVisible(visible);
	}
	void done(int result) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, -2570997490807640431LL, result, &__h); if (__h) return; }
		QFileDialog::done(result);
	}
	void accept() override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 9106131928185084681LL, &__h); if (__h) return; }
		QFileDialog::accept();
	}
	void changeEvent(QEvent* e) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4959376063826165943LL, (void*)e, &__h); if (__h) return; }
		QFileDialog::changeEvent(e);
	}
	QSize sizeHint() const override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 729690261457684101LL, &__h); if (__h) return *(QSize*)__r; }
		return QFileDialog::sizeHint();
	}
	QSize minimumSizeHint() const override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 7049264599837532491LL, &__h); if (__h) return *(QSize*)__r; }
		return QFileDialog::minimumSizeHint();
	}
	int exec() override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -1745322561779819684LL, &__h); if (__h) return __r; }
		return QFileDialog::exec();
	}
	void reject() override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 2875220963396170150LL, &__h); if (__h) return; }
		QFileDialog::reject();
	}
	void keyPressEvent(QKeyEvent* arg1) override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1023791977679793185LL, (void*)arg1, &__h); if (__h) return; }
		QFileDialog::keyPressEvent(arg1);
	}
	void closeEvent(QCloseEvent* arg1) override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1989937050019464783LL, (void*)arg1, &__h); if (__h) return; }
		QFileDialog::closeEvent(arg1);
	}
	void showEvent(QShowEvent* arg1) override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6825369174751916LL, (void*)arg1, &__h); if (__h) return; }
		QFileDialog::showEvent(arg1);
	}
	void resizeEvent(QResizeEvent* arg1) override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4302917285655014661LL, (void*)arg1, &__h); if (__h) return; }
		QFileDialog::resizeEvent(arg1);
	}
	void contextMenuEvent(QContextMenuEvent* arg1) override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6698043222227006015LL, (void*)arg1, &__h); if (__h) return; }
		QFileDialog::contextMenuEvent(arg1);
	}
	bool eventFilter(QObject* arg1, QEvent* arg2) override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)arg1, (void*)arg2, &__h); if (__h) return __r; }
		return QFileDialog::eventFilter(arg1, arg2);
	}
	int devType() const override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8659679893868524044LL, &__h); if (__h) return __r; }
		return QFileDialog::devType();
	}
	int heightForWidth(int arg1) const override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 4311534060495737339LL, arg1, &__h); if (__h) return __r; }
		return QFileDialog::heightForWidth(arg1);
	}
	bool hasHeightForWidth() const override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8266762397995423915LL, &__h); if (__h) return __r; }
		return QFileDialog::hasHeightForWidth();
	}
	QPaintEngine* paintEngine() const override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 6398007404756346381LL, &__h); if (__h) return (QPaintEngine*)__r; }
		return QFileDialog::paintEngine();
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QFileDialog::event(event);
	}
	void mousePressEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -332090898078543319LL, (void*)event, &__h); if (__h) return; }
		QFileDialog::mousePressEvent(event);
	}
	void mouseReleaseEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4415920049297795213LL, (void*)event, &__h); if (__h) return; }
		QFileDialog::mouseReleaseEvent(event);
	}
	void mouseDoubleClickEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4029073905697023887LL, (void*)event, &__h); if (__h) return; }
		QFileDialog::mouseDoubleClickEvent(event);
	}
	void mouseMoveEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4255494004099806581LL, (void*)event, &__h); if (__h) return; }
		QFileDialog::mouseMoveEvent(event);
	}
	void wheelEvent(QWheelEvent* event) override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1919988340934511316LL, (void*)event, &__h); if (__h) return; }
		QFileDialog::wheelEvent(event);
	}
	void keyReleaseEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 772598475950050299LL, (void*)event, &__h); if (__h) return; }
		QFileDialog::keyReleaseEvent(event);
	}
	void focusInEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 975043767033552552LL, (void*)event, &__h); if (__h) return; }
		QFileDialog::focusInEvent(event);
	}
	void focusOutEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -385596010576048993LL, (void*)event, &__h); if (__h) return; }
		QFileDialog::focusOutEvent(event);
	}
	void enterEvent(QEnterEvent* event) override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6316270346668237895LL, (void*)event, &__h); if (__h) return; }
		QFileDialog::enterEvent(event);
	}
	void leaveEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -551070009623711640LL, (void*)event, &__h); if (__h) return; }
		QFileDialog::leaveEvent(event);
	}
	void paintEvent(QPaintEvent* event) override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5290442163167458465LL, (void*)event, &__h); if (__h) return; }
		QFileDialog::paintEvent(event);
	}
	void moveEvent(QMoveEvent* event) override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -9161949150489525396LL, (void*)event, &__h); if (__h) return; }
		QFileDialog::moveEvent(event);
	}
	void tabletEvent(QTabletEvent* event) override {
		int __h = 0;
		if (__goOn[32] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -212068866860701141LL, (void*)event, &__h); if (__h) return; }
		QFileDialog::tabletEvent(event);
	}
	void actionEvent(QActionEvent* event) override {
		int __h = 0;
		if (__goOn[33] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1697444140088555581LL, (void*)event, &__h); if (__h) return; }
		QFileDialog::actionEvent(event);
	}
	void dragEnterEvent(QDragEnterEvent* event) override {
		int __h = 0;
		if (__goOn[34] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 206122715919917589LL, (void*)event, &__h); if (__h) return; }
		QFileDialog::dragEnterEvent(event);
	}
	void dragMoveEvent(QDragMoveEvent* event) override {
		int __h = 0;
		if (__goOn[35] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1781545137352745978LL, (void*)event, &__h); if (__h) return; }
		QFileDialog::dragMoveEvent(event);
	}
	void dragLeaveEvent(QDragLeaveEvent* event) override {
		int __h = 0;
		if (__goOn[36] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -3139904158373725006LL, (void*)event, &__h); if (__h) return; }
		QFileDialog::dragLeaveEvent(event);
	}
	void dropEvent(QDropEvent* event) override {
		int __h = 0;
		if (__goOn[37] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 9107058682754591716LL, (void*)event, &__h); if (__h) return; }
		QFileDialog::dropEvent(event);
	}
	void hideEvent(QHideEvent* event) override {
		int __h = 0;
		if (__goOn[38] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -199718707876719129LL, (void*)event, &__h); if (__h) return; }
		QFileDialog::hideEvent(event);
	}
	bool nativeEvent(const QByteArray& eventType, void* message, qintptr* result) override {
		int __h = 0;
		if (__goOn[39] && cppgo_virt_int_voidp_voidp_longlongp) { int __r = ((int(*)(long long, long long, void*, void*, long long*, int*))cppgo_virt_int_voidp_voidp_longlongp)(__goId, -2039166586246443706LL, (void*)&eventType, (void*)message, (long long*)result, &__h); if (__h) return __r; }
		return QFileDialog::nativeEvent(eventType, message, result);
	}
	int metric(QPaintDevice::PaintDeviceMetric arg1) const override {
		int __h = 0;
		if (__goOn[40] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 427629649267311777LL, (int)(arg1), &__h); if (__h) return __r; }
		return QFileDialog::metric(arg1);
	}
	void initPainter(QPainter* painter) const override {
		int __h = 0;
		if (__goOn[41] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6477500296980050210LL, (void*)painter, &__h); if (__h) return; }
		QFileDialog::initPainter(painter);
	}
	QPaintDevice* redirected(QPoint* offset) const override {
		int __h = 0;
		if (__goOn[42] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 4127240633310655048LL, (void*)offset, &__h); if (__h) return (QPaintDevice*)__r; }
		return QFileDialog::redirected(offset);
	}
	QPainter* sharedPainter() const override {
		int __h = 0;
		if (__goOn[43] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 5792771328741774991LL, &__h); if (__h) return (QPainter*)__r; }
		return QFileDialog::sharedPainter();
	}
	void inputMethodEvent(QInputMethodEvent* arg1) override {
		int __h = 0;
		if (__goOn[44] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5438819661942843222LL, (void*)arg1, &__h); if (__h) return; }
		QFileDialog::inputMethodEvent(arg1);
	}
	QVariant inputMethodQuery(Qt::InputMethodQuery arg1) const override {
		int __h = 0;
		if (__goOn[45] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7345853097340709638LL, (int)(arg1), &__h); if (__h) return *(QVariant*)__r; }
		return QFileDialog::inputMethodQuery(arg1);
	}
	bool focusNextPrevChild(bool next) override {
		int __h = 0;
		if (__goOn[46] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 5898328517505784191LL, next, &__h); if (__h) return __r; }
		return QFileDialog::focusNextPrevChild(next);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[47] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QFileDialog::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[48] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QFileDialog::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[49] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QFileDialog::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[50] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QFileDialog::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[51] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QFileDialog::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QFileDialog;
static std::mutex __connsMu_QFileDialog;

extern "C" {

void* QFileDialog_base(void* self) { return (void*)static_cast<QDialog*>((QFileDialog*)self); }
void QFileDialog_chain(void* self, void** out) {
	QFileDialog* __p = (QFileDialog*)self;
	QDialog* __b0 = static_cast<QDialog*>(__p); out[0] = (void*)__b0;
	QWidget* __b1 = static_cast<QWidget*>(__b0); out[1] = (void*)__b1;
	QObject* __b2 = static_cast<QObject*>(__b1); out[2] = (void*)__b2;
}
void QFileDialog__set_override(void* self, int idx) { ((QFileDialog__Go*)self)->__goOn[idx] = 1; }
void* QFileDialog_fromQObject(void* o) { return (void*)qobject_cast<QFileDialog*>((QObject*)o); }
void* QFileDialog_new(long long __goid, void* parent) {
	return new QFileDialog__Go(__goid, (QWidget*)parent);
}
void* QFileDialog_new2(long long __goid, void* parent, int f) {
	return new QFileDialog__Go(__goid, (QWidget*)parent, QFlags<Qt::WindowType>(f));
}
void* QFileDialog_new3(long long __goid) {
	return new QFileDialog__Go(__goid);
}
void* QFileDialog_new4(long long __goid, void* parent, const char* caption, long long caption_len) {
	QString caption_qstr = QString::fromUtf8(caption, caption_len);
	return new QFileDialog__Go(__goid, (QWidget*)parent, caption_qstr);
}
void* QFileDialog_new5(long long __goid, void* parent, const char* caption, long long caption_len, const char* directory, long long directory_len) {
	QString caption_qstr = QString::fromUtf8(caption, caption_len);
	QString directory_qstr = QString::fromUtf8(directory, directory_len);
	return new QFileDialog__Go(__goid, (QWidget*)parent, caption_qstr, directory_qstr);
}
void* QFileDialog_new6(long long __goid, void* parent, const char* caption, long long caption_len, const char* directory, long long directory_len, const char* filter, long long filter_len) {
	QString caption_qstr = QString::fromUtf8(caption, caption_len);
	QString directory_qstr = QString::fromUtf8(directory, directory_len);
	QString filter_qstr = QString::fromUtf8(filter, filter_len);
	return new QFileDialog__Go(__goid, (QWidget*)parent, caption_qstr, directory_qstr, filter_qstr);
}
char* QFileDialog_tr(const char* s, long long* __out_len) {
	auto&& __r = QFileDialog::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QFileDialog_setDirectory(void* self, const char* directory, long long directory_len) {
	QString directory_qstr = QString::fromUtf8(directory, directory_len);
	((QFileDialog*)self)->setDirectory(directory_qstr);
}
void QFileDialog_setDirectory2(void* self, void* directory) {
	((QFileDialog*)self)->setDirectory(*(QDir*)directory);
}
void* QFileDialog_directory(void* self) {
	auto&& __r = ((QFileDialog*)self)->directory();
	return new QDir(__r);
}
void QFileDialog_setDirectoryUrl(void* self, void* directory) {
	((QFileDialog*)self)->setDirectoryUrl(*(QUrl*)directory);
}
void* QFileDialog_directoryUrl(void* self) {
	auto&& __r = ((QFileDialog*)self)->directoryUrl();
	return new QUrl(__r);
}
void QFileDialog_selectFile(void* self, const char* filename, long long filename_len) {
	QString filename_qstr = QString::fromUtf8(filename, filename_len);
	((QFileDialog*)self)->selectFile(filename_qstr);
}
char** QFileDialog_selectedFiles(void* self, long long* __out_len) {
	auto&& __r = ((QFileDialog*)self)->selectedFiles();
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
void QFileDialog_selectUrl(void* self, void* url) {
	((QFileDialog*)self)->selectUrl(*(QUrl*)url);
}
void** QFileDialog_selectedUrls(void* self, long long* __out_len) {
	auto&& __r = ((QFileDialog*)self)->selectedUrls();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QUrl(__e);
	return __arr;
}
void QFileDialog_setNameFilter(void* self, const char* filter, long long filter_len) {
	QString filter_qstr = QString::fromUtf8(filter, filter_len);
	((QFileDialog*)self)->setNameFilter(filter_qstr);
}
void QFileDialog_setNameFilters(void* self, char** filters, long long filters_len) {
	QList<QString> filters_list; filters_list.reserve(filters_len);
	for (long long __i = 0; __i < filters_len; __i++) filters_list.append(QString::fromUtf8(filters[__i]));
	((QFileDialog*)self)->setNameFilters(filters_list);
}
char** QFileDialog_nameFilters(void* self, long long* __out_len) {
	auto&& __r = ((QFileDialog*)self)->nameFilters();
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
void QFileDialog_selectNameFilter(void* self, const char* filter, long long filter_len) {
	QString filter_qstr = QString::fromUtf8(filter, filter_len);
	((QFileDialog*)self)->selectNameFilter(filter_qstr);
}
char* QFileDialog_selectedMimeTypeFilter(void* self, long long* __out_len) {
	auto&& __r = ((QFileDialog*)self)->selectedMimeTypeFilter();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QFileDialog_selectedNameFilter(void* self, long long* __out_len) {
	auto&& __r = ((QFileDialog*)self)->selectedNameFilter();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QFileDialog_setMimeTypeFilters(void* self, char** filters, long long filters_len) {
	QList<QString> filters_list; filters_list.reserve(filters_len);
	for (long long __i = 0; __i < filters_len; __i++) filters_list.append(QString::fromUtf8(filters[__i]));
	((QFileDialog*)self)->setMimeTypeFilters(filters_list);
}
char** QFileDialog_mimeTypeFilters(void* self, long long* __out_len) {
	auto&& __r = ((QFileDialog*)self)->mimeTypeFilters();
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
void QFileDialog_selectMimeTypeFilter(void* self, const char* filter, long long filter_len) {
	QString filter_qstr = QString::fromUtf8(filter, filter_len);
	((QFileDialog*)self)->selectMimeTypeFilter(filter_qstr);
}
int QFileDialog_filter(void* self) {
	auto&& __r = ((QFileDialog*)self)->filter();
	return (int)(__r);
}
void QFileDialog_setFilter(void* self, int filters) {
	((QFileDialog*)self)->setFilter(QFlags<QDir::Filter>(filters));
}
void QFileDialog_setViewMode(void* self, int mode) {
	((QFileDialog*)self)->setViewMode((QFileDialog::ViewMode)(mode));
}
int QFileDialog_viewMode(void* self) {
	auto&& __r = ((QFileDialog*)self)->viewMode();
	return (int)(__r);
}
void QFileDialog_setFileMode(void* self, int mode) {
	((QFileDialog*)self)->setFileMode((QFileDialog::FileMode)(mode));
}
int QFileDialog_fileMode(void* self) {
	auto&& __r = ((QFileDialog*)self)->fileMode();
	return (int)(__r);
}
void QFileDialog_setAcceptMode(void* self, int mode) {
	((QFileDialog*)self)->setAcceptMode((QFileDialog::AcceptMode)(mode));
}
int QFileDialog_acceptMode(void* self) {
	auto&& __r = ((QFileDialog*)self)->acceptMode();
	return (int)(__r);
}
void QFileDialog_setSidebarUrls(void* self, void** urls, long long urls_len) {
	QList<QUrl> urls_list; urls_list.reserve(urls_len);
	for (long long __i = 0; __i < urls_len; __i++) urls_list.append(*(QUrl*)urls[__i]);
	((QFileDialog*)self)->setSidebarUrls(urls_list);
}
void** QFileDialog_sidebarUrls(void* self, long long* __out_len) {
	auto&& __r = ((QFileDialog*)self)->sidebarUrls();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QUrl(__e);
	return __arr;
}
void* QFileDialog_saveState(void* self) {
	auto&& __r = ((QFileDialog*)self)->saveState();
	return new QByteArray(__r);
}
int QFileDialog_restoreState(void* self, void* state) {
	auto&& __r = ((QFileDialog*)self)->restoreState(*(QByteArray*)state);
	return __r;
}
void QFileDialog_setDefaultSuffix(void* self, const char* suffix, long long suffix_len) {
	QString suffix_qstr = QString::fromUtf8(suffix, suffix_len);
	((QFileDialog*)self)->setDefaultSuffix(suffix_qstr);
}
char* QFileDialog_defaultSuffix(void* self, long long* __out_len) {
	auto&& __r = ((QFileDialog*)self)->defaultSuffix();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QFileDialog_setHistory(void* self, char** paths, long long paths_len) {
	QList<QString> paths_list; paths_list.reserve(paths_len);
	for (long long __i = 0; __i < paths_len; __i++) paths_list.append(QString::fromUtf8(paths[__i]));
	((QFileDialog*)self)->setHistory(paths_list);
}
char** QFileDialog_history(void* self, long long* __out_len) {
	auto&& __r = ((QFileDialog*)self)->history();
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
void QFileDialog_setItemDelegate(void* self, void* delegate) {
	((QFileDialog*)self)->setItemDelegate((QAbstractItemDelegate*)delegate);
}
void* QFileDialog_itemDelegate(void* self) {
	auto&& __r = ((QFileDialog*)self)->itemDelegate();
	return (void*)(__r);
}
void QFileDialog_setIconProvider(void* self, void* provider) {
	((QFileDialog*)self)->setIconProvider((QAbstractFileIconProvider*)provider);
}
void* QFileDialog_iconProvider(void* self) {
	auto&& __r = ((QFileDialog*)self)->iconProvider();
	return (void*)(__r);
}
void QFileDialog_setLabelText(void* self, int label, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	((QFileDialog*)self)->setLabelText((QFileDialog::DialogLabel)(label), text_qstr);
}
char* QFileDialog_labelText(void* self, int label, long long* __out_len) {
	auto&& __r = ((QFileDialog*)self)->labelText((QFileDialog::DialogLabel)(label));
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QFileDialog_setSupportedSchemes(void* self, char** schemes, long long schemes_len) {
	QList<QString> schemes_list; schemes_list.reserve(schemes_len);
	for (long long __i = 0; __i < schemes_len; __i++) schemes_list.append(QString::fromUtf8(schemes[__i]));
	((QFileDialog*)self)->setSupportedSchemes(schemes_list);
}
char** QFileDialog_supportedSchemes(void* self, long long* __out_len) {
	auto&& __r = ((QFileDialog*)self)->supportedSchemes();
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
void QFileDialog_setProxyModel(void* self, void* model) {
	((QFileDialog*)self)->setProxyModel((QAbstractProxyModel*)model);
}
void* QFileDialog_proxyModel(void* self) {
	auto&& __r = ((QFileDialog*)self)->proxyModel();
	return (void*)(__r);
}
void QFileDialog_setOption(void* self, int option) {
	((QFileDialog*)self)->setOption((QFileDialog::Option)(option));
}
int QFileDialog_testOption(void* self, int option) {
	auto&& __r = ((QFileDialog*)self)->testOption((QFileDialog::Option)(option));
	return __r;
}
void QFileDialog_setOptions(void* self, int options) {
	((QFileDialog*)self)->setOptions(QFlags<QFileDialog::Option>(options));
}
int QFileDialog_options(void* self) {
	auto&& __r = ((QFileDialog*)self)->options();
	return (int)(__r);
}
void QFileDialog_open(void* self, void* receiver, const char* member) {
	((QFileDialog*)self)->open((QObject*)receiver, member);
}
void QFileDialog_fileSelected(void* self, const char* file, long long file_len) {
	QString file_qstr = QString::fromUtf8(file, file_len);
	((QFileDialog*)self)->fileSelected(file_qstr);
}
void QFileDialog_filesSelected(void* self, char** files, long long files_len) {
	QList<QString> files_list; files_list.reserve(files_len);
	for (long long __i = 0; __i < files_len; __i++) files_list.append(QString::fromUtf8(files[__i]));
	((QFileDialog*)self)->filesSelected(files_list);
}
void QFileDialog_currentChanged(void* self, const char* path, long long path_len) {
	QString path_qstr = QString::fromUtf8(path, path_len);
	((QFileDialog*)self)->currentChanged(path_qstr);
}
void QFileDialog_directoryEntered(void* self, const char* directory, long long directory_len) {
	QString directory_qstr = QString::fromUtf8(directory, directory_len);
	((QFileDialog*)self)->directoryEntered(directory_qstr);
}
void QFileDialog_urlSelected(void* self, void* url) {
	((QFileDialog*)self)->urlSelected(*(QUrl*)url);
}
void QFileDialog_urlsSelected(void* self, void** urls, long long urls_len) {
	QList<QUrl> urls_list; urls_list.reserve(urls_len);
	for (long long __i = 0; __i < urls_len; __i++) urls_list.append(*(QUrl*)urls[__i]);
	((QFileDialog*)self)->urlsSelected(urls_list);
}
void QFileDialog_currentUrlChanged(void* self, void* url) {
	((QFileDialog*)self)->currentUrlChanged(*(QUrl*)url);
}
void QFileDialog_directoryUrlEntered(void* self, void* directory) {
	((QFileDialog*)self)->directoryUrlEntered(*(QUrl*)directory);
}
void QFileDialog_filterSelected(void* self, const char* filter, long long filter_len) {
	QString filter_qstr = QString::fromUtf8(filter, filter_len);
	((QFileDialog*)self)->filterSelected(filter_qstr);
}
char* QFileDialog_getOpenFileName(long long* __out_len) {
	auto&& __r = QFileDialog::getOpenFileName();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QFileDialog_getOpenFileUrl() {
	auto&& __r = QFileDialog::getOpenFileUrl();
	return new QUrl(__r);
}
char* QFileDialog_getSaveFileName(long long* __out_len) {
	auto&& __r = QFileDialog::getSaveFileName();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QFileDialog_getSaveFileUrl() {
	auto&& __r = QFileDialog::getSaveFileUrl();
	return new QUrl(__r);
}
char* QFileDialog_getExistingDirectory(long long* __out_len) {
	auto&& __r = QFileDialog::getExistingDirectory();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QFileDialog_getExistingDirectoryUrl() {
	auto&& __r = QFileDialog::getExistingDirectoryUrl();
	return new QUrl(__r);
}
char** QFileDialog_getOpenFileNames(long long* __out_len) {
	auto&& __r = QFileDialog::getOpenFileNames();
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
void** QFileDialog_getOpenFileUrls(long long* __out_len) {
	auto&& __r = QFileDialog::getOpenFileUrls();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QUrl(__e);
	return __arr;
}
void QFileDialog_saveFileContent(void* fileContent, const char* fileNameHint, long long fileNameHint_len) {
	QString fileNameHint_qstr = QString::fromUtf8(fileNameHint, fileNameHint_len);
	QFileDialog::saveFileContent(*(QByteArray*)fileContent, fileNameHint_qstr);
}
void QFileDialog_adjustPosition(void* self, void* arg1) {
	((QFileDialog__Go*)self)->adjustPosition((QWidget*)arg1);
}
void QFileDialog_setTabOrder(void* arg1, void* arg2) {
	QFileDialog::setTabOrder((QWidget*)arg1, (QWidget*)arg2);
}
void* QFileDialog_mouseGrabber() {
	auto&& __r = QFileDialog::mouseGrabber();
	return (void*)(__r);
}
void* QFileDialog_keyboardGrabber() {
	auto&& __r = QFileDialog::keyboardGrabber();
	return (void*)(__r);
}
void* QFileDialog_find(unsigned long long arg1) {
	auto&& __r = QFileDialog::find(arg1);
	return (void*)(__r);
}
void* QFileDialog_createWindowContainer(void* window) {
	auto&& __r = QFileDialog::createWindowContainer((QWindow*)window);
	return (void*)(__r);
}
void QFileDialog_updateMicroFocus(void* self) {
	((QFileDialog__Go*)self)->updateMicroFocus();
}
void QFileDialog_create(void* self) {
	((QFileDialog__Go*)self)->create();
}
void QFileDialog_destroy(void* self) {
	((QFileDialog__Go*)self)->destroy();
}
int QFileDialog_focusNextChild(void* self) {
	auto&& __r = ((QFileDialog__Go*)self)->focusNextChild();
	return __r;
}
int QFileDialog_focusPreviousChild(void* self) {
	auto&& __r = ((QFileDialog__Go*)self)->focusPreviousChild();
	return __r;
}
void* QFileDialog_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QFileDialog::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QFileDialog_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QFileDialog::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QFileDialog_sender(void* self) {
	auto&& __r = ((QFileDialog__Go*)self)->sender();
	return (void*)(__r);
}
int QFileDialog_senderSignalIndex(void* self) {
	auto&& __r = ((QFileDialog__Go*)self)->senderSignalIndex();
	return __r;
}
int QFileDialog_receivers(void* self, const char* signal) {
	auto&& __r = ((QFileDialog__Go*)self)->receivers(signal);
	return __r;
}
int QFileDialog_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QFileDialog__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
double QFileDialog_devicePixelRatioFScale() {
	auto&& __r = QFileDialog::devicePixelRatioFScale();
	return __r;
}
int QFileDialog_encodeMetricF(int metric, double value) {
	auto&& __r = QFileDialog::encodeMetricF((QPaintDevice::PaintDeviceMetric)(metric), value);
	return __r;
}
double QFileDialog_getDecodedMetricF(void* self, int metricA, int metricB) {
	auto&& __r = ((QFileDialog__Go*)self)->getDecodedMetricF((QPaintDevice::PaintDeviceMetric)(metricA), (QPaintDevice::PaintDeviceMetric)(metricB));
	return __r;
}
char* QFileDialog_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QFileDialog::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QFileDialog_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QFileDialog::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QFileDialog_setOption2(void* self, int option, int on) {
	((QFileDialog*)self)->setOption((QFileDialog::Option)(option), on);
}
char* QFileDialog_getOpenFileName2(void* parent, long long* __out_len) {
	auto&& __r = QFileDialog::getOpenFileName((QWidget*)parent);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QFileDialog_getOpenFileName3(void* parent, const char* caption, long long caption_len, long long* __out_len) {
	QString caption_qstr = QString::fromUtf8(caption, caption_len);
	auto&& __r = QFileDialog::getOpenFileName((QWidget*)parent, caption_qstr);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QFileDialog_getOpenFileName4(void* parent, const char* caption, long long caption_len, const char* dir, long long dir_len, long long* __out_len) {
	QString caption_qstr = QString::fromUtf8(caption, caption_len);
	QString dir_qstr = QString::fromUtf8(dir, dir_len);
	auto&& __r = QFileDialog::getOpenFileName((QWidget*)parent, caption_qstr, dir_qstr);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QFileDialog_getOpenFileName5(void* parent, const char* caption, long long caption_len, const char* dir, long long dir_len, const char* filter, long long filter_len, long long* __out_len) {
	QString caption_qstr = QString::fromUtf8(caption, caption_len);
	QString dir_qstr = QString::fromUtf8(dir, dir_len);
	QString filter_qstr = QString::fromUtf8(filter, filter_len);
	auto&& __r = QFileDialog::getOpenFileName((QWidget*)parent, caption_qstr, dir_qstr, filter_qstr);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QFileDialog_getOpenFileName6(void* parent, const char* caption, long long caption_len, const char* dir, long long dir_len, const char* filter, long long filter_len, void* selectedFilter, long long* __out_len) {
	QString caption_qstr = QString::fromUtf8(caption, caption_len);
	QString dir_qstr = QString::fromUtf8(dir, dir_len);
	QString filter_qstr = QString::fromUtf8(filter, filter_len);
	auto&& __r = QFileDialog::getOpenFileName((QWidget*)parent, caption_qstr, dir_qstr, filter_qstr, (QString*)selectedFilter);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QFileDialog_getOpenFileName7(void* parent, const char* caption, long long caption_len, const char* dir, long long dir_len, const char* filter, long long filter_len, void* selectedFilter, int options, long long* __out_len) {
	QString caption_qstr = QString::fromUtf8(caption, caption_len);
	QString dir_qstr = QString::fromUtf8(dir, dir_len);
	QString filter_qstr = QString::fromUtf8(filter, filter_len);
	auto&& __r = QFileDialog::getOpenFileName((QWidget*)parent, caption_qstr, dir_qstr, filter_qstr, (QString*)selectedFilter, QFlags<QFileDialog::Option>(options));
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QFileDialog_getOpenFileUrl2(void* parent) {
	auto&& __r = QFileDialog::getOpenFileUrl((QWidget*)parent);
	return new QUrl(__r);
}
void* QFileDialog_getOpenFileUrl3(void* parent, const char* caption, long long caption_len) {
	QString caption_qstr = QString::fromUtf8(caption, caption_len);
	auto&& __r = QFileDialog::getOpenFileUrl((QWidget*)parent, caption_qstr);
	return new QUrl(__r);
}
void* QFileDialog_getOpenFileUrl4(void* parent, const char* caption, long long caption_len, void* dir) {
	QString caption_qstr = QString::fromUtf8(caption, caption_len);
	auto&& __r = QFileDialog::getOpenFileUrl((QWidget*)parent, caption_qstr, *(QUrl*)dir);
	return new QUrl(__r);
}
void* QFileDialog_getOpenFileUrl5(void* parent, const char* caption, long long caption_len, void* dir, const char* filter, long long filter_len) {
	QString caption_qstr = QString::fromUtf8(caption, caption_len);
	QString filter_qstr = QString::fromUtf8(filter, filter_len);
	auto&& __r = QFileDialog::getOpenFileUrl((QWidget*)parent, caption_qstr, *(QUrl*)dir, filter_qstr);
	return new QUrl(__r);
}
void* QFileDialog_getOpenFileUrl6(void* parent, const char* caption, long long caption_len, void* dir, const char* filter, long long filter_len, void* selectedFilter) {
	QString caption_qstr = QString::fromUtf8(caption, caption_len);
	QString filter_qstr = QString::fromUtf8(filter, filter_len);
	auto&& __r = QFileDialog::getOpenFileUrl((QWidget*)parent, caption_qstr, *(QUrl*)dir, filter_qstr, (QString*)selectedFilter);
	return new QUrl(__r);
}
void* QFileDialog_getOpenFileUrl7(void* parent, const char* caption, long long caption_len, void* dir, const char* filter, long long filter_len, void* selectedFilter, int options) {
	QString caption_qstr = QString::fromUtf8(caption, caption_len);
	QString filter_qstr = QString::fromUtf8(filter, filter_len);
	auto&& __r = QFileDialog::getOpenFileUrl((QWidget*)parent, caption_qstr, *(QUrl*)dir, filter_qstr, (QString*)selectedFilter, QFlags<QFileDialog::Option>(options));
	return new QUrl(__r);
}
void* QFileDialog_getOpenFileUrl8(void* parent, const char* caption, long long caption_len, void* dir, const char* filter, long long filter_len, void* selectedFilter, int options, char** supportedSchemes, long long supportedSchemes_len) {
	QString caption_qstr = QString::fromUtf8(caption, caption_len);
	QString filter_qstr = QString::fromUtf8(filter, filter_len);
	QList<QString> supportedSchemes_list; supportedSchemes_list.reserve(supportedSchemes_len);
	for (long long __i = 0; __i < supportedSchemes_len; __i++) supportedSchemes_list.append(QString::fromUtf8(supportedSchemes[__i]));
	auto&& __r = QFileDialog::getOpenFileUrl((QWidget*)parent, caption_qstr, *(QUrl*)dir, filter_qstr, (QString*)selectedFilter, QFlags<QFileDialog::Option>(options), supportedSchemes_list);
	return new QUrl(__r);
}
char* QFileDialog_getSaveFileName2(void* parent, long long* __out_len) {
	auto&& __r = QFileDialog::getSaveFileName((QWidget*)parent);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QFileDialog_getSaveFileName3(void* parent, const char* caption, long long caption_len, long long* __out_len) {
	QString caption_qstr = QString::fromUtf8(caption, caption_len);
	auto&& __r = QFileDialog::getSaveFileName((QWidget*)parent, caption_qstr);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QFileDialog_getSaveFileName4(void* parent, const char* caption, long long caption_len, const char* dir, long long dir_len, long long* __out_len) {
	QString caption_qstr = QString::fromUtf8(caption, caption_len);
	QString dir_qstr = QString::fromUtf8(dir, dir_len);
	auto&& __r = QFileDialog::getSaveFileName((QWidget*)parent, caption_qstr, dir_qstr);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QFileDialog_getSaveFileName5(void* parent, const char* caption, long long caption_len, const char* dir, long long dir_len, const char* filter, long long filter_len, long long* __out_len) {
	QString caption_qstr = QString::fromUtf8(caption, caption_len);
	QString dir_qstr = QString::fromUtf8(dir, dir_len);
	QString filter_qstr = QString::fromUtf8(filter, filter_len);
	auto&& __r = QFileDialog::getSaveFileName((QWidget*)parent, caption_qstr, dir_qstr, filter_qstr);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QFileDialog_getSaveFileName6(void* parent, const char* caption, long long caption_len, const char* dir, long long dir_len, const char* filter, long long filter_len, void* selectedFilter, long long* __out_len) {
	QString caption_qstr = QString::fromUtf8(caption, caption_len);
	QString dir_qstr = QString::fromUtf8(dir, dir_len);
	QString filter_qstr = QString::fromUtf8(filter, filter_len);
	auto&& __r = QFileDialog::getSaveFileName((QWidget*)parent, caption_qstr, dir_qstr, filter_qstr, (QString*)selectedFilter);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QFileDialog_getSaveFileName7(void* parent, const char* caption, long long caption_len, const char* dir, long long dir_len, const char* filter, long long filter_len, void* selectedFilter, int options, long long* __out_len) {
	QString caption_qstr = QString::fromUtf8(caption, caption_len);
	QString dir_qstr = QString::fromUtf8(dir, dir_len);
	QString filter_qstr = QString::fromUtf8(filter, filter_len);
	auto&& __r = QFileDialog::getSaveFileName((QWidget*)parent, caption_qstr, dir_qstr, filter_qstr, (QString*)selectedFilter, QFlags<QFileDialog::Option>(options));
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QFileDialog_getSaveFileUrl2(void* parent) {
	auto&& __r = QFileDialog::getSaveFileUrl((QWidget*)parent);
	return new QUrl(__r);
}
void* QFileDialog_getSaveFileUrl3(void* parent, const char* caption, long long caption_len) {
	QString caption_qstr = QString::fromUtf8(caption, caption_len);
	auto&& __r = QFileDialog::getSaveFileUrl((QWidget*)parent, caption_qstr);
	return new QUrl(__r);
}
void* QFileDialog_getSaveFileUrl4(void* parent, const char* caption, long long caption_len, void* dir) {
	QString caption_qstr = QString::fromUtf8(caption, caption_len);
	auto&& __r = QFileDialog::getSaveFileUrl((QWidget*)parent, caption_qstr, *(QUrl*)dir);
	return new QUrl(__r);
}
void* QFileDialog_getSaveFileUrl5(void* parent, const char* caption, long long caption_len, void* dir, const char* filter, long long filter_len) {
	QString caption_qstr = QString::fromUtf8(caption, caption_len);
	QString filter_qstr = QString::fromUtf8(filter, filter_len);
	auto&& __r = QFileDialog::getSaveFileUrl((QWidget*)parent, caption_qstr, *(QUrl*)dir, filter_qstr);
	return new QUrl(__r);
}
void* QFileDialog_getSaveFileUrl6(void* parent, const char* caption, long long caption_len, void* dir, const char* filter, long long filter_len, void* selectedFilter) {
	QString caption_qstr = QString::fromUtf8(caption, caption_len);
	QString filter_qstr = QString::fromUtf8(filter, filter_len);
	auto&& __r = QFileDialog::getSaveFileUrl((QWidget*)parent, caption_qstr, *(QUrl*)dir, filter_qstr, (QString*)selectedFilter);
	return new QUrl(__r);
}
void* QFileDialog_getSaveFileUrl7(void* parent, const char* caption, long long caption_len, void* dir, const char* filter, long long filter_len, void* selectedFilter, int options) {
	QString caption_qstr = QString::fromUtf8(caption, caption_len);
	QString filter_qstr = QString::fromUtf8(filter, filter_len);
	auto&& __r = QFileDialog::getSaveFileUrl((QWidget*)parent, caption_qstr, *(QUrl*)dir, filter_qstr, (QString*)selectedFilter, QFlags<QFileDialog::Option>(options));
	return new QUrl(__r);
}
void* QFileDialog_getSaveFileUrl8(void* parent, const char* caption, long long caption_len, void* dir, const char* filter, long long filter_len, void* selectedFilter, int options, char** supportedSchemes, long long supportedSchemes_len) {
	QString caption_qstr = QString::fromUtf8(caption, caption_len);
	QString filter_qstr = QString::fromUtf8(filter, filter_len);
	QList<QString> supportedSchemes_list; supportedSchemes_list.reserve(supportedSchemes_len);
	for (long long __i = 0; __i < supportedSchemes_len; __i++) supportedSchemes_list.append(QString::fromUtf8(supportedSchemes[__i]));
	auto&& __r = QFileDialog::getSaveFileUrl((QWidget*)parent, caption_qstr, *(QUrl*)dir, filter_qstr, (QString*)selectedFilter, QFlags<QFileDialog::Option>(options), supportedSchemes_list);
	return new QUrl(__r);
}
char* QFileDialog_getExistingDirectory2(void* parent, long long* __out_len) {
	auto&& __r = QFileDialog::getExistingDirectory((QWidget*)parent);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QFileDialog_getExistingDirectory3(void* parent, const char* caption, long long caption_len, long long* __out_len) {
	QString caption_qstr = QString::fromUtf8(caption, caption_len);
	auto&& __r = QFileDialog::getExistingDirectory((QWidget*)parent, caption_qstr);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QFileDialog_getExistingDirectory4(void* parent, const char* caption, long long caption_len, const char* dir, long long dir_len, long long* __out_len) {
	QString caption_qstr = QString::fromUtf8(caption, caption_len);
	QString dir_qstr = QString::fromUtf8(dir, dir_len);
	auto&& __r = QFileDialog::getExistingDirectory((QWidget*)parent, caption_qstr, dir_qstr);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QFileDialog_getExistingDirectory5(void* parent, const char* caption, long long caption_len, const char* dir, long long dir_len, int options, long long* __out_len) {
	QString caption_qstr = QString::fromUtf8(caption, caption_len);
	QString dir_qstr = QString::fromUtf8(dir, dir_len);
	auto&& __r = QFileDialog::getExistingDirectory((QWidget*)parent, caption_qstr, dir_qstr, QFlags<QFileDialog::Option>(options));
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QFileDialog_getExistingDirectoryUrl2(void* parent) {
	auto&& __r = QFileDialog::getExistingDirectoryUrl((QWidget*)parent);
	return new QUrl(__r);
}
void* QFileDialog_getExistingDirectoryUrl3(void* parent, const char* caption, long long caption_len) {
	QString caption_qstr = QString::fromUtf8(caption, caption_len);
	auto&& __r = QFileDialog::getExistingDirectoryUrl((QWidget*)parent, caption_qstr);
	return new QUrl(__r);
}
void* QFileDialog_getExistingDirectoryUrl4(void* parent, const char* caption, long long caption_len, void* dir) {
	QString caption_qstr = QString::fromUtf8(caption, caption_len);
	auto&& __r = QFileDialog::getExistingDirectoryUrl((QWidget*)parent, caption_qstr, *(QUrl*)dir);
	return new QUrl(__r);
}
void* QFileDialog_getExistingDirectoryUrl5(void* parent, const char* caption, long long caption_len, void* dir, int options) {
	QString caption_qstr = QString::fromUtf8(caption, caption_len);
	auto&& __r = QFileDialog::getExistingDirectoryUrl((QWidget*)parent, caption_qstr, *(QUrl*)dir, QFlags<QFileDialog::Option>(options));
	return new QUrl(__r);
}
void* QFileDialog_getExistingDirectoryUrl6(void* parent, const char* caption, long long caption_len, void* dir, int options, char** supportedSchemes, long long supportedSchemes_len) {
	QString caption_qstr = QString::fromUtf8(caption, caption_len);
	QList<QString> supportedSchemes_list; supportedSchemes_list.reserve(supportedSchemes_len);
	for (long long __i = 0; __i < supportedSchemes_len; __i++) supportedSchemes_list.append(QString::fromUtf8(supportedSchemes[__i]));
	auto&& __r = QFileDialog::getExistingDirectoryUrl((QWidget*)parent, caption_qstr, *(QUrl*)dir, QFlags<QFileDialog::Option>(options), supportedSchemes_list);
	return new QUrl(__r);
}
char** QFileDialog_getOpenFileNames2(void* parent, long long* __out_len) {
	auto&& __r = QFileDialog::getOpenFileNames((QWidget*)parent);
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
char** QFileDialog_getOpenFileNames3(void* parent, const char* caption, long long caption_len, long long* __out_len) {
	QString caption_qstr = QString::fromUtf8(caption, caption_len);
	auto&& __r = QFileDialog::getOpenFileNames((QWidget*)parent, caption_qstr);
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
char** QFileDialog_getOpenFileNames4(void* parent, const char* caption, long long caption_len, const char* dir, long long dir_len, long long* __out_len) {
	QString caption_qstr = QString::fromUtf8(caption, caption_len);
	QString dir_qstr = QString::fromUtf8(dir, dir_len);
	auto&& __r = QFileDialog::getOpenFileNames((QWidget*)parent, caption_qstr, dir_qstr);
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
char** QFileDialog_getOpenFileNames5(void* parent, const char* caption, long long caption_len, const char* dir, long long dir_len, const char* filter, long long filter_len, long long* __out_len) {
	QString caption_qstr = QString::fromUtf8(caption, caption_len);
	QString dir_qstr = QString::fromUtf8(dir, dir_len);
	QString filter_qstr = QString::fromUtf8(filter, filter_len);
	auto&& __r = QFileDialog::getOpenFileNames((QWidget*)parent, caption_qstr, dir_qstr, filter_qstr);
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
char** QFileDialog_getOpenFileNames6(void* parent, const char* caption, long long caption_len, const char* dir, long long dir_len, const char* filter, long long filter_len, void* selectedFilter, long long* __out_len) {
	QString caption_qstr = QString::fromUtf8(caption, caption_len);
	QString dir_qstr = QString::fromUtf8(dir, dir_len);
	QString filter_qstr = QString::fromUtf8(filter, filter_len);
	auto&& __r = QFileDialog::getOpenFileNames((QWidget*)parent, caption_qstr, dir_qstr, filter_qstr, (QString*)selectedFilter);
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
char** QFileDialog_getOpenFileNames7(void* parent, const char* caption, long long caption_len, const char* dir, long long dir_len, const char* filter, long long filter_len, void* selectedFilter, int options, long long* __out_len) {
	QString caption_qstr = QString::fromUtf8(caption, caption_len);
	QString dir_qstr = QString::fromUtf8(dir, dir_len);
	QString filter_qstr = QString::fromUtf8(filter, filter_len);
	auto&& __r = QFileDialog::getOpenFileNames((QWidget*)parent, caption_qstr, dir_qstr, filter_qstr, (QString*)selectedFilter, QFlags<QFileDialog::Option>(options));
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
void** QFileDialog_getOpenFileUrls2(void* parent, long long* __out_len) {
	auto&& __r = QFileDialog::getOpenFileUrls((QWidget*)parent);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QUrl(__e);
	return __arr;
}
void** QFileDialog_getOpenFileUrls3(void* parent, const char* caption, long long caption_len, long long* __out_len) {
	QString caption_qstr = QString::fromUtf8(caption, caption_len);
	auto&& __r = QFileDialog::getOpenFileUrls((QWidget*)parent, caption_qstr);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QUrl(__e);
	return __arr;
}
void** QFileDialog_getOpenFileUrls4(void* parent, const char* caption, long long caption_len, void* dir, long long* __out_len) {
	QString caption_qstr = QString::fromUtf8(caption, caption_len);
	auto&& __r = QFileDialog::getOpenFileUrls((QWidget*)parent, caption_qstr, *(QUrl*)dir);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QUrl(__e);
	return __arr;
}
void** QFileDialog_getOpenFileUrls5(void* parent, const char* caption, long long caption_len, void* dir, const char* filter, long long filter_len, long long* __out_len) {
	QString caption_qstr = QString::fromUtf8(caption, caption_len);
	QString filter_qstr = QString::fromUtf8(filter, filter_len);
	auto&& __r = QFileDialog::getOpenFileUrls((QWidget*)parent, caption_qstr, *(QUrl*)dir, filter_qstr);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QUrl(__e);
	return __arr;
}
void** QFileDialog_getOpenFileUrls6(void* parent, const char* caption, long long caption_len, void* dir, const char* filter, long long filter_len, void* selectedFilter, long long* __out_len) {
	QString caption_qstr = QString::fromUtf8(caption, caption_len);
	QString filter_qstr = QString::fromUtf8(filter, filter_len);
	auto&& __r = QFileDialog::getOpenFileUrls((QWidget*)parent, caption_qstr, *(QUrl*)dir, filter_qstr, (QString*)selectedFilter);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QUrl(__e);
	return __arr;
}
void** QFileDialog_getOpenFileUrls7(void* parent, const char* caption, long long caption_len, void* dir, const char* filter, long long filter_len, void* selectedFilter, int options, long long* __out_len) {
	QString caption_qstr = QString::fromUtf8(caption, caption_len);
	QString filter_qstr = QString::fromUtf8(filter, filter_len);
	auto&& __r = QFileDialog::getOpenFileUrls((QWidget*)parent, caption_qstr, *(QUrl*)dir, filter_qstr, (QString*)selectedFilter, QFlags<QFileDialog::Option>(options));
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QUrl(__e);
	return __arr;
}
void** QFileDialog_getOpenFileUrls8(void* parent, const char* caption, long long caption_len, void* dir, const char* filter, long long filter_len, void* selectedFilter, int options, char** supportedSchemes, long long supportedSchemes_len, long long* __out_len) {
	QString caption_qstr = QString::fromUtf8(caption, caption_len);
	QString filter_qstr = QString::fromUtf8(filter, filter_len);
	QList<QString> supportedSchemes_list; supportedSchemes_list.reserve(supportedSchemes_len);
	for (long long __i = 0; __i < supportedSchemes_len; __i++) supportedSchemes_list.append(QString::fromUtf8(supportedSchemes[__i]));
	auto&& __r = QFileDialog::getOpenFileUrls((QWidget*)parent, caption_qstr, *(QUrl*)dir, filter_qstr, (QString*)selectedFilter, QFlags<QFileDialog::Option>(options), supportedSchemes_list);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QUrl(__e);
	return __arr;
}
void QFileDialog_saveFileContent2(void* fileContent, const char* fileNameHint, long long fileNameHint_len, void* parent) {
	QString fileNameHint_qstr = QString::fromUtf8(fileNameHint, fileNameHint_len);
	QFileDialog::saveFileContent(*(QByteArray*)fileContent, fileNameHint_qstr, (QWidget*)parent);
}
void* QFileDialog_createWindowContainer2(void* window, void* parent) {
	auto&& __r = QFileDialog::createWindowContainer((QWindow*)window, (QWidget*)parent);
	return (void*)(__r);
}
void* QFileDialog_createWindowContainer3(void* window, void* parent, int flags) {
	auto&& __r = QFileDialog::createWindowContainer((QWindow*)window, (QWidget*)parent, QFlags<Qt::WindowType>(flags));
	return (void*)(__r);
}
void QFileDialog_updateMicroFocus2(void* self, int query) {
	((QFileDialog__Go*)self)->updateMicroFocus((Qt::InputMethodQuery)(query));
}
void QFileDialog_create2(void* self, unsigned long long arg1) {
	((QFileDialog__Go*)self)->create(arg1);
}
void QFileDialog_create3(void* self, unsigned long long arg1, int initializeWindow) {
	((QFileDialog__Go*)self)->create(arg1, initializeWindow);
}
void QFileDialog_create4(void* self, unsigned long long arg1, int initializeWindow, int destroyOldWindow) {
	((QFileDialog__Go*)self)->create(arg1, initializeWindow, destroyOldWindow);
}
void QFileDialog_destroy2(void* self, int destroyWindow) {
	((QFileDialog__Go*)self)->destroy(destroyWindow);
}
void QFileDialog_destroy3(void* self, int destroyWindow, int destroySubWindows) {
	((QFileDialog__Go*)self)->destroy(destroyWindow, destroySubWindows);
}
void* QFileDialog_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QFileDialog::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QFileDialog_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QFileDialog::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QFileDialog_metaObject_default(void* self) {
	auto&& __r = ((QFileDialog*)self)->QFileDialog::metaObject();
	return (void*)(__r);
}
void QFileDialog_setVisible_default(void* self, int visible) {
	((QFileDialog*)self)->QFileDialog::setVisible(visible);
}
void QFileDialog_done_default(void* self, int result) {
	((QFileDialog__Go*)self)->cppgoBase_done(result);
}
void QFileDialog_accept_default(void* self) {
	((QFileDialog__Go*)self)->cppgoBase_accept();
}
void QFileDialog_changeEvent_default(void* self, void* e) {
	((QFileDialog__Go*)self)->cppgoBase_changeEvent((QEvent*)e);
}
void* QFileDialog_sizeHint_default(void* self) {
	auto&& __r = ((QFileDialog*)self)->QFileDialog::sizeHint();
	return new QSize(__r);
}
void* QFileDialog_minimumSizeHint_default(void* self) {
	auto&& __r = ((QFileDialog*)self)->QFileDialog::minimumSizeHint();
	return new QSize(__r);
}
int QFileDialog_exec_default(void* self) {
	auto&& __r = ((QFileDialog*)self)->QFileDialog::exec();
	return __r;
}
void QFileDialog_reject_default(void* self) {
	((QFileDialog*)self)->QFileDialog::reject();
}
void QFileDialog_keyPressEvent_default(void* self, void* arg1) {
	((QFileDialog__Go*)self)->cppgoBase_keyPressEvent((QKeyEvent*)arg1);
}
void QFileDialog_closeEvent_default(void* self, void* arg1) {
	((QFileDialog__Go*)self)->cppgoBase_closeEvent((QCloseEvent*)arg1);
}
void QFileDialog_showEvent_default(void* self, void* arg1) {
	((QFileDialog__Go*)self)->cppgoBase_showEvent((QShowEvent*)arg1);
}
void QFileDialog_resizeEvent_default(void* self, void* arg1) {
	((QFileDialog__Go*)self)->cppgoBase_resizeEvent((QResizeEvent*)arg1);
}
void QFileDialog_contextMenuEvent_default(void* self, void* arg1) {
	((QFileDialog__Go*)self)->cppgoBase_contextMenuEvent((QContextMenuEvent*)arg1);
}
int QFileDialog_eventFilter_default(void* self, void* arg1, void* arg2) {
	auto&& __r = ((QFileDialog__Go*)self)->cppgoBase_eventFilter((QObject*)arg1, (QEvent*)arg2);
	return __r;
}
int QFileDialog_devType_default(void* self) {
	auto&& __r = ((QFileDialog*)self)->QFileDialog::devType();
	return __r;
}
int QFileDialog_heightForWidth_default(void* self, int arg1) {
	auto&& __r = ((QFileDialog*)self)->QFileDialog::heightForWidth(arg1);
	return __r;
}
int QFileDialog_hasHeightForWidth_default(void* self) {
	auto&& __r = ((QFileDialog*)self)->QFileDialog::hasHeightForWidth();
	return __r;
}
void* QFileDialog_paintEngine_default(void* self) {
	auto&& __r = ((QFileDialog*)self)->QFileDialog::paintEngine();
	return (void*)(__r);
}
int QFileDialog_event_default(void* self, void* event) {
	auto&& __r = ((QFileDialog__Go*)self)->cppgoBase_event((QEvent*)event);
	return __r;
}
void QFileDialog_mousePressEvent_default(void* self, void* event) {
	((QFileDialog__Go*)self)->cppgoBase_mousePressEvent((QMouseEvent*)event);
}
void QFileDialog_mouseReleaseEvent_default(void* self, void* event) {
	((QFileDialog__Go*)self)->cppgoBase_mouseReleaseEvent((QMouseEvent*)event);
}
void QFileDialog_mouseDoubleClickEvent_default(void* self, void* event) {
	((QFileDialog__Go*)self)->cppgoBase_mouseDoubleClickEvent((QMouseEvent*)event);
}
void QFileDialog_mouseMoveEvent_default(void* self, void* event) {
	((QFileDialog__Go*)self)->cppgoBase_mouseMoveEvent((QMouseEvent*)event);
}
void QFileDialog_wheelEvent_default(void* self, void* event) {
	((QFileDialog__Go*)self)->cppgoBase_wheelEvent((QWheelEvent*)event);
}
void QFileDialog_keyReleaseEvent_default(void* self, void* event) {
	((QFileDialog__Go*)self)->cppgoBase_keyReleaseEvent((QKeyEvent*)event);
}
void QFileDialog_focusInEvent_default(void* self, void* event) {
	((QFileDialog__Go*)self)->cppgoBase_focusInEvent((QFocusEvent*)event);
}
void QFileDialog_focusOutEvent_default(void* self, void* event) {
	((QFileDialog__Go*)self)->cppgoBase_focusOutEvent((QFocusEvent*)event);
}
void QFileDialog_enterEvent_default(void* self, void* event) {
	((QFileDialog__Go*)self)->cppgoBase_enterEvent((QEnterEvent*)event);
}
void QFileDialog_leaveEvent_default(void* self, void* event) {
	((QFileDialog__Go*)self)->cppgoBase_leaveEvent((QEvent*)event);
}
void QFileDialog_paintEvent_default(void* self, void* event) {
	((QFileDialog__Go*)self)->cppgoBase_paintEvent((QPaintEvent*)event);
}
void QFileDialog_moveEvent_default(void* self, void* event) {
	((QFileDialog__Go*)self)->cppgoBase_moveEvent((QMoveEvent*)event);
}
void QFileDialog_tabletEvent_default(void* self, void* event) {
	((QFileDialog__Go*)self)->cppgoBase_tabletEvent((QTabletEvent*)event);
}
void QFileDialog_actionEvent_default(void* self, void* event) {
	((QFileDialog__Go*)self)->cppgoBase_actionEvent((QActionEvent*)event);
}
void QFileDialog_dragEnterEvent_default(void* self, void* event) {
	((QFileDialog__Go*)self)->cppgoBase_dragEnterEvent((QDragEnterEvent*)event);
}
void QFileDialog_dragMoveEvent_default(void* self, void* event) {
	((QFileDialog__Go*)self)->cppgoBase_dragMoveEvent((QDragMoveEvent*)event);
}
void QFileDialog_dragLeaveEvent_default(void* self, void* event) {
	((QFileDialog__Go*)self)->cppgoBase_dragLeaveEvent((QDragLeaveEvent*)event);
}
void QFileDialog_dropEvent_default(void* self, void* event) {
	((QFileDialog__Go*)self)->cppgoBase_dropEvent((QDropEvent*)event);
}
void QFileDialog_hideEvent_default(void* self, void* event) {
	((QFileDialog__Go*)self)->cppgoBase_hideEvent((QHideEvent*)event);
}
int QFileDialog_metric_default(void* self, int arg1) {
	auto&& __r = ((QFileDialog__Go*)self)->cppgoBase_metric(arg1);
	return __r;
}
void QFileDialog_initPainter_default(void* self, void* painter) {
	((QFileDialog__Go*)self)->cppgoBase_initPainter((QPainter*)painter);
}
void* QFileDialog_redirected_default(void* self, void* offset) {
	auto&& __r = ((QFileDialog__Go*)self)->cppgoBase_redirected((QPoint*)offset);
	return (void*)(__r);
}
void* QFileDialog_sharedPainter_default(void* self) {
	auto&& __r = ((QFileDialog__Go*)self)->cppgoBase_sharedPainter();
	return (void*)(__r);
}
void QFileDialog_inputMethodEvent_default(void* self, void* arg1) {
	((QFileDialog__Go*)self)->cppgoBase_inputMethodEvent((QInputMethodEvent*)arg1);
}
void* QFileDialog_inputMethodQuery_default(void* self, int arg1) {
	auto&& __r = ((QFileDialog*)self)->QFileDialog::inputMethodQuery((Qt::InputMethodQuery)(arg1));
	return new QVariant(__r);
}
int QFileDialog_focusNextPrevChild_default(void* self, int next) {
	auto&& __r = ((QFileDialog__Go*)self)->cppgoBase_focusNextPrevChild(next);
	return __r;
}
void QFileDialog_timerEvent_default(void* self, void* event) {
	((QFileDialog__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QFileDialog_childEvent_default(void* self, void* event) {
	((QFileDialog__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QFileDialog_customEvent_default(void* self, void* event) {
	((QFileDialog__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QFileDialog_connectNotify_default(void* self, void* signal) {
	((QFileDialog__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QFileDialog_disconnectNotify_default(void* self, void* signal) {
	((QFileDialog__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QFileDialog_connect_fileSelected(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFileDialog*)self, &QFileDialog::fileSelected, (QFileDialog*)self, [__id](const QString& file){
		QByteArray file_ba = file.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)file_ba.constData(), file_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFileDialog);
	__conns_QFileDialog[__id] = __c;
}
void QFileDialog_connect_filesSelected(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFileDialog*)self, &QFileDialog::filesSelected, (QFileDialog*)self, [__id](const QList<QString>& files){
		std::vector<void*> files_v; files_v.reserve(files.size());
		for (long long __i = 0; __i < files.size(); __i++) files_v.push_back((void*)&files[__i]);
		if (cppgo_sig_void_voidpp_longlong) ((void(*)(long long, void**, long long))cppgo_sig_void_voidpp_longlong)(__id, files_v.data(), (long long)files.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFileDialog);
	__conns_QFileDialog[__id] = __c;
}
void QFileDialog_connect_currentChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFileDialog*)self, &QFileDialog::currentChanged, (QFileDialog*)self, [__id](const QString& path){
		QByteArray path_ba = path.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)path_ba.constData(), path_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFileDialog);
	__conns_QFileDialog[__id] = __c;
}
void QFileDialog_connect_directoryEntered(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFileDialog*)self, &QFileDialog::directoryEntered, (QFileDialog*)self, [__id](const QString& directory){
		QByteArray directory_ba = directory.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)directory_ba.constData(), directory_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFileDialog);
	__conns_QFileDialog[__id] = __c;
}
void QFileDialog_connect_urlSelected(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFileDialog*)self, &QFileDialog::urlSelected, (QFileDialog*)self, [__id](const QUrl& url){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&url);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFileDialog);
	__conns_QFileDialog[__id] = __c;
}
void QFileDialog_connect_urlsSelected(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFileDialog*)self, &QFileDialog::urlsSelected, (QFileDialog*)self, [__id](const QList<QUrl>& urls){
		std::vector<void*> urls_v; urls_v.reserve(urls.size());
		for (long long __i = 0; __i < urls.size(); __i++) urls_v.push_back((void*)&urls[__i]);
		if (cppgo_sig_void_voidpp_longlong) ((void(*)(long long, void**, long long))cppgo_sig_void_voidpp_longlong)(__id, urls_v.data(), (long long)urls.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFileDialog);
	__conns_QFileDialog[__id] = __c;
}
void QFileDialog_connect_currentUrlChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFileDialog*)self, &QFileDialog::currentUrlChanged, (QFileDialog*)self, [__id](const QUrl& url){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&url);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFileDialog);
	__conns_QFileDialog[__id] = __c;
}
void QFileDialog_connect_directoryUrlEntered(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFileDialog*)self, &QFileDialog::directoryUrlEntered, (QFileDialog*)self, [__id](const QUrl& directory){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&directory);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFileDialog);
	__conns_QFileDialog[__id] = __c;
}
void QFileDialog_connect_filterSelected(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFileDialog*)self, &QFileDialog::filterSelected, (QFileDialog*)self, [__id](const QString& filter){
		QByteArray filter_ba = filter.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)filter_ba.constData(), filter_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFileDialog);
	__conns_QFileDialog[__id] = __c;
}
void QFileDialog_connect_finished(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFileDialog*)self, &QFileDialog::finished, (QFileDialog*)self, [__id](int result){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, result);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFileDialog);
	__conns_QFileDialog[__id] = __c;
}
void QFileDialog_connect_accepted(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFileDialog*)self, &QFileDialog::accepted, (QFileDialog*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFileDialog);
	__conns_QFileDialog[__id] = __c;
}
void QFileDialog_connect_rejected(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFileDialog*)self, &QFileDialog::rejected, (QFileDialog*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFileDialog);
	__conns_QFileDialog[__id] = __c;
}
void QFileDialog_connect_windowTitleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFileDialog*)self, &QFileDialog::windowTitleChanged, (QFileDialog*)self, [__id](const QString& title){
		QByteArray title_ba = title.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)title_ba.constData(), title_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFileDialog);
	__conns_QFileDialog[__id] = __c;
}
void QFileDialog_connect_windowIconChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFileDialog*)self, &QFileDialog::windowIconChanged, (QFileDialog*)self, [__id](const QIcon& icon){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&icon);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFileDialog);
	__conns_QFileDialog[__id] = __c;
}
void QFileDialog_connect_windowIconTextChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFileDialog*)self, &QFileDialog::windowIconTextChanged, (QFileDialog*)self, [__id](const QString& iconText){
		QByteArray iconText_ba = iconText.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)iconText_ba.constData(), iconText_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFileDialog);
	__conns_QFileDialog[__id] = __c;
}
void QFileDialog_connect_customContextMenuRequested(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFileDialog*)self, &QFileDialog::customContextMenuRequested, (QFileDialog*)self, [__id](const QPoint& pos){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&pos);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFileDialog);
	__conns_QFileDialog[__id] = __c;
}
void QFileDialog_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFileDialog*)self, &QFileDialog::destroyed, (QFileDialog*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFileDialog);
	__conns_QFileDialog[__id] = __c;
}
void QFileDialog_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFileDialog*)self, &QFileDialog::destroyed, (QFileDialog*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFileDialog);
	__conns_QFileDialog[__id] = __c;
}
void QFileDialog_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFileDialog*)self, &QFileDialog::objectNameChanged, (QFileDialog*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFileDialog);
	__conns_QFileDialog[__id] = __c;
}
void QFileDialog__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QFileDialog);
		auto __it = __conns_QFileDialog.find(id);
		if (__it == __conns_QFileDialog.end()) return;
		__c = __it->second;
		__conns_QFileDialog.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QFileDialog__Go_new(long long id) { return new QFileDialog__Go(id); }
void QFileDialog_delete(void* self) { delete (QFileDialog*)self; }

}

#include <QtWidgets/qgraphicsitem.h>
#include <QtCore/qstring.h>
#include <QtCore/qobjectdefs.h>
#include <QtGui/qfont.h>
#include <QtGui/qcolor.h>
#include <QtCore/qrect.h>
#include <QtGui/qpainterpath.h>
#include <QtCore/qpoint.h>
#include <QtGui/qpainter.h>
#include <QtWidgets/qstyleoption.h>
#include <QtWidgets/qwidget.h>
#include <QtGui/qtextdocument.h>
#include <QtCore/qnamespace.h>
#include <QtGui/qtextcursor.h>
#include <QtCore/qcoreevent.h>
#include <QtWidgets/qgraphicssceneevent.h>
#include <QtGui/qevent.h>
#include <QtCore/qvariant.h>
#include <QtCore/qobject.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qbindingstorage.h>
#include <QtWidgets/qgraphicsscene.h>
#include <QtWidgets/qgraphicswidget.h>
#include <QtCore/qsize.h>
#include <QtGui/qcursor.h>
#include <QtWidgets/qgraphicseffect.h>
#include <QtGui/qtransform.h>
#include <QtWidgets/qgraphicstransform.h>
#include <QtGui/qregion.h>
#include <QtGui/qpolygon.h>
#include "gen_qgraphicstextitem.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QGraphicsTextItem__Go : public QGraphicsTextItem {
public:
	long long __goId;
	unsigned char __goOn[43] = {};
	template<typename... A> QGraphicsTextItem__Go(long long id, A&&... a) : QGraphicsTextItem(std::forward<A>(a)...), __goId(id) {}
	using QGraphicsTextItem::updateMicroFocus;
	using QGraphicsTextItem::sender;
	using QGraphicsTextItem::senderSignalIndex;
	using QGraphicsTextItem::receivers;
	using QGraphicsTextItem::isSignalConnected;
	using QGraphicsTextItem::addToIndex;
	using QGraphicsTextItem::removeFromIndex;
	using QGraphicsTextItem::prepareGeometryChange;
	bool cppgoBase_sceneEvent(QEvent* __a0) { return QGraphicsTextItem::sceneEvent(__a0); }
	void cppgoBase_mousePressEvent(QGraphicsSceneMouseEvent* __a0) { QGraphicsTextItem::mousePressEvent(__a0); }
	void cppgoBase_mouseMoveEvent(QGraphicsSceneMouseEvent* __a0) { QGraphicsTextItem::mouseMoveEvent(__a0); }
	void cppgoBase_mouseReleaseEvent(QGraphicsSceneMouseEvent* __a0) { QGraphicsTextItem::mouseReleaseEvent(__a0); }
	void cppgoBase_mouseDoubleClickEvent(QGraphicsSceneMouseEvent* __a0) { QGraphicsTextItem::mouseDoubleClickEvent(__a0); }
	void cppgoBase_contextMenuEvent(QGraphicsSceneContextMenuEvent* __a0) { QGraphicsTextItem::contextMenuEvent(__a0); }
	void cppgoBase_keyPressEvent(QKeyEvent* __a0) { QGraphicsTextItem::keyPressEvent(__a0); }
	void cppgoBase_keyReleaseEvent(QKeyEvent* __a0) { QGraphicsTextItem::keyReleaseEvent(__a0); }
	void cppgoBase_focusInEvent(QFocusEvent* __a0) { QGraphicsTextItem::focusInEvent(__a0); }
	void cppgoBase_focusOutEvent(QFocusEvent* __a0) { QGraphicsTextItem::focusOutEvent(__a0); }
	void cppgoBase_dragEnterEvent(QGraphicsSceneDragDropEvent* __a0) { QGraphicsTextItem::dragEnterEvent(__a0); }
	void cppgoBase_dragLeaveEvent(QGraphicsSceneDragDropEvent* __a0) { QGraphicsTextItem::dragLeaveEvent(__a0); }
	void cppgoBase_dragMoveEvent(QGraphicsSceneDragDropEvent* __a0) { QGraphicsTextItem::dragMoveEvent(__a0); }
	void cppgoBase_dropEvent(QGraphicsSceneDragDropEvent* __a0) { QGraphicsTextItem::dropEvent(__a0); }
	void cppgoBase_inputMethodEvent(QInputMethodEvent* __a0) { QGraphicsTextItem::inputMethodEvent(__a0); }
	void cppgoBase_hoverEnterEvent(QGraphicsSceneHoverEvent* __a0) { QGraphicsTextItem::hoverEnterEvent(__a0); }
	void cppgoBase_hoverMoveEvent(QGraphicsSceneHoverEvent* __a0) { QGraphicsTextItem::hoverMoveEvent(__a0); }
	void cppgoBase_hoverLeaveEvent(QGraphicsSceneHoverEvent* __a0) { QGraphicsTextItem::hoverLeaveEvent(__a0); }
	QVariant cppgoBase_inputMethodQuery(int __a0) { return QGraphicsTextItem::inputMethodQuery((Qt::InputMethodQuery)__a0); }
	bool cppgoBase_supportsExtension(int __a0) { return QGraphicsTextItem::supportsExtension((QGraphicsItem::Extension)__a0); }
	void cppgoBase_setExtension(int __a0, const QVariant& __a1) { QGraphicsTextItem::setExtension((QGraphicsItem::Extension)__a0, __a1); }
	QVariant cppgoBase_extension(const QVariant& __a0) { return QGraphicsTextItem::extension(__a0); }
	bool cppgoBase_event(QEvent* __a0) { return QGraphicsTextItem::event(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QGraphicsTextItem::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QGraphicsTextItem::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QGraphicsTextItem::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QGraphicsTextItem::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QGraphicsTextItem::disconnectNotify(__a0); }
	bool cppgoBase_sceneEventFilter(QGraphicsItem* __a0, QEvent* __a1) { return QGraphicsTextItem::sceneEventFilter(__a0, __a1); }
	void cppgoBase_wheelEvent(QGraphicsSceneWheelEvent* __a0) { QGraphicsTextItem::wheelEvent(__a0); }
	QVariant cppgoBase_itemChange(int __a0, const QVariant& __a1) { return QGraphicsTextItem::itemChange((QGraphicsItem::GraphicsItemChange)__a0, __a1); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QGraphicsTextItem::metaObject();
	}
	QRectF boundingRect() const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 2510008782785038831LL, &__h); if (__h) return *(QRectF*)__r; }
		return QGraphicsTextItem::boundingRect();
	}
	QPainterPath shape() const override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 7938019069247444294LL, &__h); if (__h) return *(QPainterPath*)__r; }
		return QGraphicsTextItem::shape();
	}
	bool contains(const QPointF& point) const override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 5645366511944179768LL, (void*)&point, &__h); if (__h) return __r; }
		return QGraphicsTextItem::contains(point);
	}
	void paint(QPainter* painter, const QStyleOptionGraphicsItem* option, QWidget* widget) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_voidp_voidp_voidp) { ((void(*)(long long, long long, void*, void*, void*, int*))cppgo_virt_void_voidp_voidp_voidp)(__goId, 6083584261053556819LL, (void*)painter, (void*)option, (void*)widget, &__h); if (__h) return; }
		QGraphicsTextItem::paint(painter, option, widget);
	}
	bool isObscuredBy(const QGraphicsItem* item) const override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, -7215030271914903671LL, (void*)item, &__h); if (__h) return __r; }
		return QGraphicsTextItem::isObscuredBy(item);
	}
	QPainterPath opaqueArea() const override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 399370970023989435LL, &__h); if (__h) return *(QPainterPath*)__r; }
		return QGraphicsTextItem::opaqueArea();
	}
	int type() const override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -6371403872082420691LL, &__h); if (__h) return __r; }
		return QGraphicsTextItem::type();
	}
	bool sceneEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, -2846371507838587071LL, (void*)event, &__h); if (__h) return __r; }
		return QGraphicsTextItem::sceneEvent(event);
	}
	void mousePressEvent(QGraphicsSceneMouseEvent* event) override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -332090898078543319LL, (void*)event, &__h); if (__h) return; }
		QGraphicsTextItem::mousePressEvent(event);
	}
	void mouseMoveEvent(QGraphicsSceneMouseEvent* event) override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4255494004099806581LL, (void*)event, &__h); if (__h) return; }
		QGraphicsTextItem::mouseMoveEvent(event);
	}
	void mouseReleaseEvent(QGraphicsSceneMouseEvent* event) override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4415920049297795213LL, (void*)event, &__h); if (__h) return; }
		QGraphicsTextItem::mouseReleaseEvent(event);
	}
	void mouseDoubleClickEvent(QGraphicsSceneMouseEvent* event) override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4029073905697023887LL, (void*)event, &__h); if (__h) return; }
		QGraphicsTextItem::mouseDoubleClickEvent(event);
	}
	void contextMenuEvent(QGraphicsSceneContextMenuEvent* event) override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6698043222227006015LL, (void*)event, &__h); if (__h) return; }
		QGraphicsTextItem::contextMenuEvent(event);
	}
	void keyPressEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1023791977679793185LL, (void*)event, &__h); if (__h) return; }
		QGraphicsTextItem::keyPressEvent(event);
	}
	void keyReleaseEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 772598475950050299LL, (void*)event, &__h); if (__h) return; }
		QGraphicsTextItem::keyReleaseEvent(event);
	}
	void focusInEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 975043767033552552LL, (void*)event, &__h); if (__h) return; }
		QGraphicsTextItem::focusInEvent(event);
	}
	void focusOutEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -385596010576048993LL, (void*)event, &__h); if (__h) return; }
		QGraphicsTextItem::focusOutEvent(event);
	}
	void dragEnterEvent(QGraphicsSceneDragDropEvent* event) override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 206122715919917589LL, (void*)event, &__h); if (__h) return; }
		QGraphicsTextItem::dragEnterEvent(event);
	}
	void dragLeaveEvent(QGraphicsSceneDragDropEvent* event) override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -3139904158373725006LL, (void*)event, &__h); if (__h) return; }
		QGraphicsTextItem::dragLeaveEvent(event);
	}
	void dragMoveEvent(QGraphicsSceneDragDropEvent* event) override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1781545137352745978LL, (void*)event, &__h); if (__h) return; }
		QGraphicsTextItem::dragMoveEvent(event);
	}
	void dropEvent(QGraphicsSceneDragDropEvent* event) override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 9107058682754591716LL, (void*)event, &__h); if (__h) return; }
		QGraphicsTextItem::dropEvent(event);
	}
	void inputMethodEvent(QInputMethodEvent* event) override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5438819661942843222LL, (void*)event, &__h); if (__h) return; }
		QGraphicsTextItem::inputMethodEvent(event);
	}
	void hoverEnterEvent(QGraphicsSceneHoverEvent* event) override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6932941943906681679LL, (void*)event, &__h); if (__h) return; }
		QGraphicsTextItem::hoverEnterEvent(event);
	}
	void hoverMoveEvent(QGraphicsSceneHoverEvent* event) override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -2745191522144316396LL, (void*)event, &__h); if (__h) return; }
		QGraphicsTextItem::hoverMoveEvent(event);
	}
	void hoverLeaveEvent(QGraphicsSceneHoverEvent* event) override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7944408482203603984LL, (void*)event, &__h); if (__h) return; }
		QGraphicsTextItem::hoverLeaveEvent(event);
	}
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7345853097340709638LL, (int)(query), &__h); if (__h) return *(QVariant*)__r; }
		return QGraphicsTextItem::inputMethodQuery(query);
	}
	bool supportsExtension(QGraphicsItem::Extension extension) const override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, -5729166443651306418LL, (int)(extension), &__h); if (__h) return __r; }
		return QGraphicsTextItem::supportsExtension(extension);
	}
	void setExtension(QGraphicsItem::Extension extension, const QVariant& variant) override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_void_int_voidp) { ((void(*)(long long, long long, int, void*, int*))cppgo_virt_void_int_voidp)(__goId, -6730181576106171580LL, (int)(extension), (void*)&variant, &__h); if (__h) return; }
		QGraphicsTextItem::setExtension(extension, variant);
	}
	QVariant extension(const QVariant& variant) const override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 2270616358835251226LL, (void*)&variant, &__h); if (__h) return *(QVariant*)__r; }
		return QGraphicsTextItem::extension(variant);
	}
	bool event(QEvent* ev) override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)ev, &__h); if (__h) return __r; }
		return QGraphicsTextItem::event(ev);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QGraphicsTextItem::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[32] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QGraphicsTextItem::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[33] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QGraphicsTextItem::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[34] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QGraphicsTextItem::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[35] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QGraphicsTextItem::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[36] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QGraphicsTextItem::disconnectNotify(signal);
	}
	void advance(int phase) override {
		int __h = 0;
		if (__goOn[37] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, -4830836341480073475LL, phase, &__h); if (__h) return; }
		QGraphicsTextItem::advance(phase);
	}
	bool collidesWithItem(const QGraphicsItem* other, Qt::ItemSelectionMode mode) const override {
		int __h = 0;
		if (__goOn[38] && cppgo_virt_int_voidp_int) { int __r = ((int(*)(long long, long long, void*, int, int*))cppgo_virt_int_voidp_int)(__goId, 6469736184362593173LL, (void*)other, (int)(mode), &__h); if (__h) return __r; }
		return QGraphicsTextItem::collidesWithItem(other, mode);
	}
	bool collidesWithPath(const QPainterPath& path, Qt::ItemSelectionMode mode) const override {
		int __h = 0;
		if (__goOn[39] && cppgo_virt_int_voidp_int) { int __r = ((int(*)(long long, long long, void*, int, int*))cppgo_virt_int_voidp_int)(__goId, 6636575182718753853LL, (void*)&path, (int)(mode), &__h); if (__h) return __r; }
		return QGraphicsTextItem::collidesWithPath(path, mode);
	}
	bool sceneEventFilter(QGraphicsItem* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[40] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, -4584203218241215725LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QGraphicsTextItem::sceneEventFilter(watched, event);
	}
	void wheelEvent(QGraphicsSceneWheelEvent* event) override {
		int __h = 0;
		if (__goOn[41] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1919988340934511316LL, (void*)event, &__h); if (__h) return; }
		QGraphicsTextItem::wheelEvent(event);
	}
	QVariant itemChange(QGraphicsItem::GraphicsItemChange change, const QVariant& value) override {
		int __h = 0;
		if (__goOn[42] && cppgo_virt_voidp_int_voidp) { void* __r = ((void*(*)(long long, long long, int, void*, int*))cppgo_virt_voidp_int_voidp)(__goId, -1888572168794878830LL, (int)(change), (void*)&value, &__h); if (__h) return *(QVariant*)__r; }
		return QGraphicsTextItem::itemChange(change, value);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QGraphicsTextItem;
static std::mutex __connsMu_QGraphicsTextItem;

extern "C" {

void* QGraphicsTextItem_base(void* self) { return (void*)static_cast<QGraphicsObject*>((QGraphicsTextItem*)self); }
void QGraphicsTextItem_chain(void* self, void** out) {
	QGraphicsTextItem* __p = (QGraphicsTextItem*)self;
	QGraphicsObject* __b0 = static_cast<QGraphicsObject*>(__p); out[0] = (void*)__b0;
	QObject* __b1 = static_cast<QObject*>(__b0); out[1] = (void*)__b1;
}
void QGraphicsTextItem__set_override(void* self, int idx) { ((QGraphicsTextItem__Go*)self)->__goOn[idx] = 1; }
void* QGraphicsTextItem_fromQObject(void* o) { return (void*)qobject_cast<QGraphicsTextItem*>((QObject*)o); }
void* QGraphicsTextItem_new(long long __goid) {
	return new QGraphicsTextItem__Go(__goid);
}
void* QGraphicsTextItem_new2(long long __goid, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	return new QGraphicsTextItem__Go(__goid, text_qstr);
}
void* QGraphicsTextItem_new3(long long __goid, void* parent) {
	return new QGraphicsTextItem__Go(__goid, (QGraphicsItem*)parent);
}
void* QGraphicsTextItem_new4(long long __goid, const char* text, long long text_len, void* parent) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	return new QGraphicsTextItem__Go(__goid, text_qstr, (QGraphicsItem*)parent);
}
char* QGraphicsTextItem_tr(const char* s, long long* __out_len) {
	auto&& __r = QGraphicsTextItem::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QGraphicsTextItem_toHtml(void* self, long long* __out_len) {
	auto&& __r = ((QGraphicsTextItem*)self)->toHtml();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QGraphicsTextItem_setHtml(void* self, const char* html, long long html_len) {
	QString html_qstr = QString::fromUtf8(html, html_len);
	((QGraphicsTextItem*)self)->setHtml(html_qstr);
}
char* QGraphicsTextItem_toPlainText(void* self, long long* __out_len) {
	auto&& __r = ((QGraphicsTextItem*)self)->toPlainText();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QGraphicsTextItem_setPlainText(void* self, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	((QGraphicsTextItem*)self)->setPlainText(text_qstr);
}
void* QGraphicsTextItem_font(void* self) {
	auto&& __r = ((QGraphicsTextItem*)self)->font();
	return new QFont(__r);
}
void QGraphicsTextItem_setFont(void* self, void* font) {
	((QGraphicsTextItem*)self)->setFont(*(QFont*)font);
}
void QGraphicsTextItem_setDefaultTextColor(void* self, void* c) {
	((QGraphicsTextItem*)self)->setDefaultTextColor(*(QColor*)c);
}
void* QGraphicsTextItem_defaultTextColor(void* self) {
	auto&& __r = ((QGraphicsTextItem*)self)->defaultTextColor();
	return new QColor(__r);
}
void QGraphicsTextItem_setTextWidth(void* self, double width) {
	((QGraphicsTextItem*)self)->setTextWidth(width);
}
double QGraphicsTextItem_textWidth(void* self) {
	auto&& __r = ((QGraphicsTextItem*)self)->textWidth();
	return __r;
}
void QGraphicsTextItem_adjustSize(void* self) {
	((QGraphicsTextItem*)self)->adjustSize();
}
void QGraphicsTextItem_setDocument(void* self, void* document) {
	((QGraphicsTextItem*)self)->setDocument((QTextDocument*)document);
}
void* QGraphicsTextItem_document(void* self) {
	auto&& __r = ((QGraphicsTextItem*)self)->document();
	return (void*)(__r);
}
void QGraphicsTextItem_setTextInteractionFlags(void* self, int flags) {
	((QGraphicsTextItem*)self)->setTextInteractionFlags(QFlags<Qt::TextInteractionFlag>(flags));
}
int QGraphicsTextItem_textInteractionFlags(void* self) {
	auto&& __r = ((QGraphicsTextItem*)self)->textInteractionFlags();
	return (int)(__r);
}
void QGraphicsTextItem_setTabChangesFocus(void* self, int b) {
	((QGraphicsTextItem*)self)->setTabChangesFocus(b);
}
int QGraphicsTextItem_tabChangesFocus(void* self) {
	auto&& __r = ((QGraphicsTextItem*)self)->tabChangesFocus();
	return __r;
}
void QGraphicsTextItem_setOpenExternalLinks(void* self, int open) {
	((QGraphicsTextItem*)self)->setOpenExternalLinks(open);
}
int QGraphicsTextItem_openExternalLinks(void* self) {
	auto&& __r = ((QGraphicsTextItem*)self)->openExternalLinks();
	return __r;
}
void QGraphicsTextItem_setTextCursor(void* self, void* cursor) {
	((QGraphicsTextItem*)self)->setTextCursor(*(QTextCursor*)cursor);
}
void* QGraphicsTextItem_textCursor(void* self) {
	auto&& __r = ((QGraphicsTextItem*)self)->textCursor();
	return new QTextCursor(__r);
}
void QGraphicsTextItem_linkActivated(void* self, const char* arg1, long long arg1_len) {
	QString arg1_qstr = QString::fromUtf8(arg1, arg1_len);
	((QGraphicsTextItem*)self)->linkActivated(arg1_qstr);
}
void QGraphicsTextItem_linkHovered(void* self, const char* arg1, long long arg1_len) {
	QString arg1_qstr = QString::fromUtf8(arg1, arg1_len);
	((QGraphicsTextItem*)self)->linkHovered(arg1_qstr);
}
void QGraphicsTextItem_updateMicroFocus(void* self) {
	((QGraphicsTextItem__Go*)self)->updateMicroFocus();
}
void* QGraphicsTextItem_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QGraphicsTextItem::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QGraphicsTextItem_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QGraphicsTextItem::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QGraphicsTextItem_sender(void* self) {
	auto&& __r = ((QGraphicsTextItem__Go*)self)->sender();
	return (void*)(__r);
}
int QGraphicsTextItem_senderSignalIndex(void* self) {
	auto&& __r = ((QGraphicsTextItem__Go*)self)->senderSignalIndex();
	return __r;
}
int QGraphicsTextItem_receivers(void* self, const char* signal) {
	auto&& __r = ((QGraphicsTextItem__Go*)self)->receivers(signal);
	return __r;
}
int QGraphicsTextItem_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QGraphicsTextItem__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
void QGraphicsTextItem_addToIndex(void* self) {
	((QGraphicsTextItem__Go*)self)->addToIndex();
}
void QGraphicsTextItem_removeFromIndex(void* self) {
	((QGraphicsTextItem__Go*)self)->removeFromIndex();
}
void QGraphicsTextItem_prepareGeometryChange(void* self) {
	((QGraphicsTextItem__Go*)self)->prepareGeometryChange();
}
char* QGraphicsTextItem_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QGraphicsTextItem::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QGraphicsTextItem_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QGraphicsTextItem::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QGraphicsTextItem_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QGraphicsTextItem::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QGraphicsTextItem_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QGraphicsTextItem::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QGraphicsTextItem_metaObject_default(void* self) {
	auto&& __r = ((QGraphicsTextItem*)self)->QGraphicsTextItem::metaObject();
	return (void*)(__r);
}
void* QGraphicsTextItem_boundingRect_default(void* self) {
	auto&& __r = ((QGraphicsTextItem*)self)->QGraphicsTextItem::boundingRect();
	return new QRectF(__r);
}
void* QGraphicsTextItem_shape_default(void* self) {
	auto&& __r = ((QGraphicsTextItem*)self)->QGraphicsTextItem::shape();
	return new QPainterPath(__r);
}
int QGraphicsTextItem_contains_default(void* self, void* point) {
	auto&& __r = ((QGraphicsTextItem*)self)->QGraphicsTextItem::contains(*(QPointF*)point);
	return __r;
}
void QGraphicsTextItem_paint_default(void* self, void* painter, void* option, void* widget) {
	((QGraphicsTextItem*)self)->QGraphicsTextItem::paint((QPainter*)painter, (QStyleOptionGraphicsItem*)option, (QWidget*)widget);
}
int QGraphicsTextItem_isObscuredBy_default(void* self, void* item) {
	auto&& __r = ((QGraphicsTextItem*)self)->QGraphicsTextItem::isObscuredBy((QGraphicsItem*)item);
	return __r;
}
void* QGraphicsTextItem_opaqueArea_default(void* self) {
	auto&& __r = ((QGraphicsTextItem*)self)->QGraphicsTextItem::opaqueArea();
	return new QPainterPath(__r);
}
int QGraphicsTextItem_type_default(void* self) {
	auto&& __r = ((QGraphicsTextItem*)self)->QGraphicsTextItem::type();
	return __r;
}
int QGraphicsTextItem_sceneEvent_default(void* self, void* event) {
	auto&& __r = ((QGraphicsTextItem__Go*)self)->cppgoBase_sceneEvent((QEvent*)event);
	return __r;
}
void QGraphicsTextItem_mousePressEvent_default(void* self, void* event) {
	((QGraphicsTextItem__Go*)self)->cppgoBase_mousePressEvent((QGraphicsSceneMouseEvent*)event);
}
void QGraphicsTextItem_mouseMoveEvent_default(void* self, void* event) {
	((QGraphicsTextItem__Go*)self)->cppgoBase_mouseMoveEvent((QGraphicsSceneMouseEvent*)event);
}
void QGraphicsTextItem_mouseReleaseEvent_default(void* self, void* event) {
	((QGraphicsTextItem__Go*)self)->cppgoBase_mouseReleaseEvent((QGraphicsSceneMouseEvent*)event);
}
void QGraphicsTextItem_mouseDoubleClickEvent_default(void* self, void* event) {
	((QGraphicsTextItem__Go*)self)->cppgoBase_mouseDoubleClickEvent((QGraphicsSceneMouseEvent*)event);
}
void QGraphicsTextItem_contextMenuEvent_default(void* self, void* event) {
	((QGraphicsTextItem__Go*)self)->cppgoBase_contextMenuEvent((QGraphicsSceneContextMenuEvent*)event);
}
void QGraphicsTextItem_keyPressEvent_default(void* self, void* event) {
	((QGraphicsTextItem__Go*)self)->cppgoBase_keyPressEvent((QKeyEvent*)event);
}
void QGraphicsTextItem_keyReleaseEvent_default(void* self, void* event) {
	((QGraphicsTextItem__Go*)self)->cppgoBase_keyReleaseEvent((QKeyEvent*)event);
}
void QGraphicsTextItem_focusInEvent_default(void* self, void* event) {
	((QGraphicsTextItem__Go*)self)->cppgoBase_focusInEvent((QFocusEvent*)event);
}
void QGraphicsTextItem_focusOutEvent_default(void* self, void* event) {
	((QGraphicsTextItem__Go*)self)->cppgoBase_focusOutEvent((QFocusEvent*)event);
}
void QGraphicsTextItem_dragEnterEvent_default(void* self, void* event) {
	((QGraphicsTextItem__Go*)self)->cppgoBase_dragEnterEvent((QGraphicsSceneDragDropEvent*)event);
}
void QGraphicsTextItem_dragLeaveEvent_default(void* self, void* event) {
	((QGraphicsTextItem__Go*)self)->cppgoBase_dragLeaveEvent((QGraphicsSceneDragDropEvent*)event);
}
void QGraphicsTextItem_dragMoveEvent_default(void* self, void* event) {
	((QGraphicsTextItem__Go*)self)->cppgoBase_dragMoveEvent((QGraphicsSceneDragDropEvent*)event);
}
void QGraphicsTextItem_dropEvent_default(void* self, void* event) {
	((QGraphicsTextItem__Go*)self)->cppgoBase_dropEvent((QGraphicsSceneDragDropEvent*)event);
}
void QGraphicsTextItem_inputMethodEvent_default(void* self, void* event) {
	((QGraphicsTextItem__Go*)self)->cppgoBase_inputMethodEvent((QInputMethodEvent*)event);
}
void QGraphicsTextItem_hoverEnterEvent_default(void* self, void* event) {
	((QGraphicsTextItem__Go*)self)->cppgoBase_hoverEnterEvent((QGraphicsSceneHoverEvent*)event);
}
void QGraphicsTextItem_hoverMoveEvent_default(void* self, void* event) {
	((QGraphicsTextItem__Go*)self)->cppgoBase_hoverMoveEvent((QGraphicsSceneHoverEvent*)event);
}
void QGraphicsTextItem_hoverLeaveEvent_default(void* self, void* event) {
	((QGraphicsTextItem__Go*)self)->cppgoBase_hoverLeaveEvent((QGraphicsSceneHoverEvent*)event);
}
void* QGraphicsTextItem_inputMethodQuery_default(void* self, int query) {
	auto&& __r = ((QGraphicsTextItem__Go*)self)->cppgoBase_inputMethodQuery(query);
	return new QVariant(__r);
}
int QGraphicsTextItem_supportsExtension_default(void* self, int extension) {
	auto&& __r = ((QGraphicsTextItem__Go*)self)->cppgoBase_supportsExtension(extension);
	return __r;
}
void QGraphicsTextItem_setExtension_default(void* self, int extension, void* variant) {
	((QGraphicsTextItem__Go*)self)->cppgoBase_setExtension(extension, *(QVariant*)variant);
}
void* QGraphicsTextItem_extension_default(void* self, void* variant) {
	auto&& __r = ((QGraphicsTextItem__Go*)self)->cppgoBase_extension(*(QVariant*)variant);
	return new QVariant(__r);
}
int QGraphicsTextItem_event_default(void* self, void* ev) {
	auto&& __r = ((QGraphicsTextItem__Go*)self)->cppgoBase_event((QEvent*)ev);
	return __r;
}
int QGraphicsTextItem_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QGraphicsTextItem*)self)->QGraphicsTextItem::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QGraphicsTextItem_timerEvent_default(void* self, void* event) {
	((QGraphicsTextItem__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QGraphicsTextItem_childEvent_default(void* self, void* event) {
	((QGraphicsTextItem__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QGraphicsTextItem_customEvent_default(void* self, void* event) {
	((QGraphicsTextItem__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QGraphicsTextItem_connectNotify_default(void* self, void* signal) {
	((QGraphicsTextItem__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QGraphicsTextItem_disconnectNotify_default(void* self, void* signal) {
	((QGraphicsTextItem__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QGraphicsTextItem_advance_default(void* self, int phase) {
	((QGraphicsTextItem*)self)->QGraphicsTextItem::advance(phase);
}
int QGraphicsTextItem_collidesWithItem_default(void* self, void* other, int mode) {
	auto&& __r = ((QGraphicsTextItem*)self)->QGraphicsTextItem::collidesWithItem((QGraphicsItem*)other, (Qt::ItemSelectionMode)(mode));
	return __r;
}
int QGraphicsTextItem_collidesWithPath_default(void* self, void* path, int mode) {
	auto&& __r = ((QGraphicsTextItem*)self)->QGraphicsTextItem::collidesWithPath(*(QPainterPath*)path, (Qt::ItemSelectionMode)(mode));
	return __r;
}
int QGraphicsTextItem_sceneEventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QGraphicsTextItem__Go*)self)->cppgoBase_sceneEventFilter((QGraphicsItem*)watched, (QEvent*)event);
	return __r;
}
void QGraphicsTextItem_wheelEvent_default(void* self, void* event) {
	((QGraphicsTextItem__Go*)self)->cppgoBase_wheelEvent((QGraphicsSceneWheelEvent*)event);
}
void* QGraphicsTextItem_itemChange_default(void* self, int change, void* value) {
	auto&& __r = ((QGraphicsTextItem__Go*)self)->cppgoBase_itemChange(change, *(QVariant*)value);
	return new QVariant(__r);
}
void QGraphicsTextItem_connect_linkActivated(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsTextItem*)self, &QGraphicsTextItem::linkActivated, (QGraphicsTextItem*)self, [__id](const QString& arg1){
		QByteArray arg1_ba = arg1.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)arg1_ba.constData(), arg1_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsTextItem);
	__conns_QGraphicsTextItem[__id] = __c;
}
void QGraphicsTextItem_connect_linkHovered(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsTextItem*)self, &QGraphicsTextItem::linkHovered, (QGraphicsTextItem*)self, [__id](const QString& arg1){
		QByteArray arg1_ba = arg1.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)arg1_ba.constData(), arg1_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsTextItem);
	__conns_QGraphicsTextItem[__id] = __c;
}
void QGraphicsTextItem_connect_parentChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsTextItem*)self, &QGraphicsTextItem::parentChanged, (QGraphicsTextItem*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsTextItem);
	__conns_QGraphicsTextItem[__id] = __c;
}
void QGraphicsTextItem_connect_opacityChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsTextItem*)self, &QGraphicsTextItem::opacityChanged, (QGraphicsTextItem*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsTextItem);
	__conns_QGraphicsTextItem[__id] = __c;
}
void QGraphicsTextItem_connect_visibleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsTextItem*)self, &QGraphicsTextItem::visibleChanged, (QGraphicsTextItem*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsTextItem);
	__conns_QGraphicsTextItem[__id] = __c;
}
void QGraphicsTextItem_connect_enabledChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsTextItem*)self, &QGraphicsTextItem::enabledChanged, (QGraphicsTextItem*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsTextItem);
	__conns_QGraphicsTextItem[__id] = __c;
}
void QGraphicsTextItem_connect_xChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsTextItem*)self, &QGraphicsTextItem::xChanged, (QGraphicsTextItem*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsTextItem);
	__conns_QGraphicsTextItem[__id] = __c;
}
void QGraphicsTextItem_connect_yChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsTextItem*)self, &QGraphicsTextItem::yChanged, (QGraphicsTextItem*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsTextItem);
	__conns_QGraphicsTextItem[__id] = __c;
}
void QGraphicsTextItem_connect_zChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsTextItem*)self, &QGraphicsTextItem::zChanged, (QGraphicsTextItem*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsTextItem);
	__conns_QGraphicsTextItem[__id] = __c;
}
void QGraphicsTextItem_connect_rotationChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsTextItem*)self, &QGraphicsTextItem::rotationChanged, (QGraphicsTextItem*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsTextItem);
	__conns_QGraphicsTextItem[__id] = __c;
}
void QGraphicsTextItem_connect_scaleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsTextItem*)self, &QGraphicsTextItem::scaleChanged, (QGraphicsTextItem*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsTextItem);
	__conns_QGraphicsTextItem[__id] = __c;
}
void QGraphicsTextItem_connect_childrenChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsTextItem*)self, &QGraphicsTextItem::childrenChanged, (QGraphicsTextItem*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsTextItem);
	__conns_QGraphicsTextItem[__id] = __c;
}
void QGraphicsTextItem_connect_widthChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsTextItem*)self, &QGraphicsTextItem::widthChanged, (QGraphicsTextItem*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsTextItem);
	__conns_QGraphicsTextItem[__id] = __c;
}
void QGraphicsTextItem_connect_heightChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsTextItem*)self, &QGraphicsTextItem::heightChanged, (QGraphicsTextItem*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsTextItem);
	__conns_QGraphicsTextItem[__id] = __c;
}
void QGraphicsTextItem_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsTextItem*)self, &QGraphicsTextItem::destroyed, (QGraphicsTextItem*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsTextItem);
	__conns_QGraphicsTextItem[__id] = __c;
}
void QGraphicsTextItem_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsTextItem*)self, &QGraphicsTextItem::destroyed, (QGraphicsTextItem*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsTextItem);
	__conns_QGraphicsTextItem[__id] = __c;
}
void QGraphicsTextItem_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsTextItem*)self, &QGraphicsTextItem::objectNameChanged, (QGraphicsTextItem*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsTextItem);
	__conns_QGraphicsTextItem[__id] = __c;
}
void QGraphicsTextItem__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsTextItem);
		auto __it = __conns_QGraphicsTextItem.find(id);
		if (__it == __conns_QGraphicsTextItem.end()) return;
		__c = __it->second;
		__conns_QGraphicsTextItem.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QGraphicsTextItem__Go_new(long long id) { return new QGraphicsTextItem__Go(id); }
void QGraphicsTextItem_delete(void* self) { delete (QGraphicsTextItem*)self; }

}

#include <QtWidgets/qgraphicssceneevent.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qpoint.h>
#include <QtCore/qnamespace.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qmimedata.h>
#include "gen_qgraphicsscenedragdropevent.h"
#include "gen_purego_cb.h"

class QGraphicsSceneDragDropEvent__Go : public QGraphicsSceneDragDropEvent {
public:
	long long __goId;
	unsigned char __goOn[2] = {};
	template<typename... A> QGraphicsSceneDragDropEvent__Go(long long id, A&&... a) : QGraphicsSceneDragDropEvent(std::forward<A>(a)...), __goId(id) {}
	void setAccepted(bool accepted) override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 6367896784077321574LL, accepted, &__h); if (__h) return; }
		QGraphicsSceneDragDropEvent::setAccepted(accepted);
	}
	QEvent* clone() const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, -7032813281048616478LL, &__h); if (__h) return (QEvent*)__r; }
		return QGraphicsSceneDragDropEvent::clone();
	}
};

extern "C" {

void* QGraphicsSceneDragDropEvent_base(void* self) { return (void*)static_cast<QGraphicsSceneEvent*>((QGraphicsSceneDragDropEvent*)self); }
void QGraphicsSceneDragDropEvent_chain(void* self, void** out) {
	QGraphicsSceneDragDropEvent* __p = (QGraphicsSceneDragDropEvent*)self;
	QGraphicsSceneEvent* __b0 = static_cast<QGraphicsSceneEvent*>(__p); out[0] = (void*)__b0;
	QEvent* __b1 = static_cast<QEvent*>(__b0); out[1] = (void*)__b1;
}
void QGraphicsSceneDragDropEvent__set_override(void* self, int idx) { ((QGraphicsSceneDragDropEvent__Go*)self)->__goOn[idx] = 1; }
void* QGraphicsSceneDragDropEvent_new(long long __goid) {
	return new QGraphicsSceneDragDropEvent__Go(__goid);
}
void* QGraphicsSceneDragDropEvent_new2(long long __goid, int type_) {
	return new QGraphicsSceneDragDropEvent__Go(__goid, (QEvent::Type)(type_));
}
void* QGraphicsSceneDragDropEvent_pos(void* self) {
	auto&& __r = ((QGraphicsSceneDragDropEvent*)self)->pos();
	return new QPointF(__r);
}
void QGraphicsSceneDragDropEvent_setPos(void* self, void* pos) {
	((QGraphicsSceneDragDropEvent*)self)->setPos(*(QPointF*)pos);
}
void* QGraphicsSceneDragDropEvent_scenePos(void* self) {
	auto&& __r = ((QGraphicsSceneDragDropEvent*)self)->scenePos();
	return new QPointF(__r);
}
void QGraphicsSceneDragDropEvent_setScenePos(void* self, void* pos) {
	((QGraphicsSceneDragDropEvent*)self)->setScenePos(*(QPointF*)pos);
}
void* QGraphicsSceneDragDropEvent_screenPos(void* self) {
	auto&& __r = ((QGraphicsSceneDragDropEvent*)self)->screenPos();
	return new QPoint(__r);
}
void QGraphicsSceneDragDropEvent_setScreenPos(void* self, void* pos) {
	((QGraphicsSceneDragDropEvent*)self)->setScreenPos(*(QPoint*)pos);
}
int QGraphicsSceneDragDropEvent_buttons(void* self) {
	auto&& __r = ((QGraphicsSceneDragDropEvent*)self)->buttons();
	return (int)(__r);
}
void QGraphicsSceneDragDropEvent_setButtons(void* self, int buttons) {
	((QGraphicsSceneDragDropEvent*)self)->setButtons(QFlags<Qt::MouseButton>(buttons));
}
int QGraphicsSceneDragDropEvent_modifiers(void* self) {
	auto&& __r = ((QGraphicsSceneDragDropEvent*)self)->modifiers();
	return (int)(__r);
}
void QGraphicsSceneDragDropEvent_setModifiers(void* self, int modifiers) {
	((QGraphicsSceneDragDropEvent*)self)->setModifiers(QFlags<Qt::KeyboardModifier>(modifiers));
}
int QGraphicsSceneDragDropEvent_possibleActions(void* self) {
	auto&& __r = ((QGraphicsSceneDragDropEvent*)self)->possibleActions();
	return (int)(__r);
}
void QGraphicsSceneDragDropEvent_setPossibleActions(void* self, int actions) {
	((QGraphicsSceneDragDropEvent*)self)->setPossibleActions(QFlags<Qt::DropAction>(actions));
}
int QGraphicsSceneDragDropEvent_proposedAction(void* self) {
	auto&& __r = ((QGraphicsSceneDragDropEvent*)self)->proposedAction();
	return (int)(__r);
}
void QGraphicsSceneDragDropEvent_setProposedAction(void* self, int action) {
	((QGraphicsSceneDragDropEvent*)self)->setProposedAction((Qt::DropAction)(action));
}
void QGraphicsSceneDragDropEvent_acceptProposedAction(void* self) {
	((QGraphicsSceneDragDropEvent*)self)->acceptProposedAction();
}
int QGraphicsSceneDragDropEvent_dropAction(void* self) {
	auto&& __r = ((QGraphicsSceneDragDropEvent*)self)->dropAction();
	return (int)(__r);
}
void QGraphicsSceneDragDropEvent_setDropAction(void* self, int action) {
	((QGraphicsSceneDragDropEvent*)self)->setDropAction((Qt::DropAction)(action));
}
void* QGraphicsSceneDragDropEvent_source(void* self) {
	auto&& __r = ((QGraphicsSceneDragDropEvent*)self)->source();
	return (void*)(__r);
}
void QGraphicsSceneDragDropEvent_setSource(void* self, void* source) {
	((QGraphicsSceneDragDropEvent*)self)->setSource((QWidget*)source);
}
void* QGraphicsSceneDragDropEvent_mimeData(void* self) {
	auto&& __r = ((QGraphicsSceneDragDropEvent*)self)->mimeData();
	return (void*)(__r);
}
void QGraphicsSceneDragDropEvent_setMimeData(void* self, void* data) {
	((QGraphicsSceneDragDropEvent*)self)->setMimeData((QMimeData*)data);
}
int QGraphicsSceneDragDropEvent_registerEventType() {
	auto&& __r = QGraphicsSceneDragDropEvent::registerEventType();
	return __r;
}
int QGraphicsSceneDragDropEvent_registerEventType2(int hint) {
	auto&& __r = QGraphicsSceneDragDropEvent::registerEventType(hint);
	return __r;
}
void QGraphicsSceneDragDropEvent_setAccepted_default(void* self, int accepted) {
	((QGraphicsSceneDragDropEvent*)self)->QGraphicsSceneDragDropEvent::setAccepted(accepted);
}
void* QGraphicsSceneDragDropEvent_clone_default(void* self) {
	auto&& __r = ((QGraphicsSceneDragDropEvent*)self)->QGraphicsSceneDragDropEvent::clone();
	return (void*)(__r);
}
void* QGraphicsSceneDragDropEvent__Go_new(long long id) { return new QGraphicsSceneDragDropEvent__Go(id); }
void QGraphicsSceneDragDropEvent_delete(void* self) { delete (QGraphicsSceneDragDropEvent*)self; }

}

#include <QtWidgets/qmenu.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qstring.h>
#include <QtCore/qobjectdefs.h>
#include <QtGui/qaction.h>
#include <QtCore/qobject.h>
#include <QtGui/qkeysequence.h>
#include <QtGui/qicon.h>
#include <QtCore/qpoint.h>
#include <QtCore/qsize.h>
#include <QtCore/qrect.h>
#include <QtCore/qcoreevent.h>
#include <QtGui/qevent.h>
#include <QtWidgets/qstyleoption.h>
#include <QtWidgets/qstyle.h>
#include <QtCore/qnamespace.h>
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
#include <QtGui/qpaintengine.h>
#include <QtGui/qbackingstore.h>
#include <QtGui/qwindow.h>
#include <QtGui/qscreen.h>
#include <QtCore/qvariant.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qmenu.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QMenu__Go : public QMenu {
public:
	long long __goId;
	unsigned char __goOn[49] = {};
	template<typename... A> QMenu__Go(long long id, A&&... a) : QMenu(std::forward<A>(a)...), __goId(id) {}
	using QMenu::columnCount;
	using QMenu::updateMicroFocus;
	using QMenu::create;
	using QMenu::destroy;
	using QMenu::focusNextChild;
	using QMenu::focusPreviousChild;
	using QMenu::sender;
	using QMenu::senderSignalIndex;
	using QMenu::receivers;
	using QMenu::isSignalConnected;
	using QMenu::getDecodedMetricF;
	void cppgoBase_changeEvent(QEvent* __a0) { QMenu::changeEvent(__a0); }
	void cppgoBase_keyPressEvent(QKeyEvent* __a0) { QMenu::keyPressEvent(__a0); }
	void cppgoBase_mouseReleaseEvent(QMouseEvent* __a0) { QMenu::mouseReleaseEvent(__a0); }
	void cppgoBase_mousePressEvent(QMouseEvent* __a0) { QMenu::mousePressEvent(__a0); }
	void cppgoBase_mouseMoveEvent(QMouseEvent* __a0) { QMenu::mouseMoveEvent(__a0); }
	void cppgoBase_wheelEvent(QWheelEvent* __a0) { QMenu::wheelEvent(__a0); }
	void cppgoBase_enterEvent(QEnterEvent* __a0) { QMenu::enterEvent(__a0); }
	void cppgoBase_leaveEvent(QEvent* __a0) { QMenu::leaveEvent(__a0); }
	void cppgoBase_hideEvent(QHideEvent* __a0) { QMenu::hideEvent(__a0); }
	void cppgoBase_paintEvent(QPaintEvent* __a0) { QMenu::paintEvent(__a0); }
	void cppgoBase_actionEvent(QActionEvent* __a0) { QMenu::actionEvent(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QMenu::timerEvent(__a0); }
	bool cppgoBase_event(QEvent* __a0) { return QMenu::event(__a0); }
	bool cppgoBase_focusNextPrevChild(bool __a0) { return QMenu::focusNextPrevChild(__a0); }
	void cppgoBase_initStyleOption(QStyleOptionMenuItem* __a0, const QAction* __a1) { QMenu::initStyleOption(__a0, __a1); }
	void cppgoBase_mouseDoubleClickEvent(QMouseEvent* __a0) { QMenu::mouseDoubleClickEvent(__a0); }
	void cppgoBase_keyReleaseEvent(QKeyEvent* __a0) { QMenu::keyReleaseEvent(__a0); }
	void cppgoBase_focusInEvent(QFocusEvent* __a0) { QMenu::focusInEvent(__a0); }
	void cppgoBase_focusOutEvent(QFocusEvent* __a0) { QMenu::focusOutEvent(__a0); }
	void cppgoBase_moveEvent(QMoveEvent* __a0) { QMenu::moveEvent(__a0); }
	void cppgoBase_resizeEvent(QResizeEvent* __a0) { QMenu::resizeEvent(__a0); }
	void cppgoBase_closeEvent(QCloseEvent* __a0) { QMenu::closeEvent(__a0); }
	void cppgoBase_contextMenuEvent(QContextMenuEvent* __a0) { QMenu::contextMenuEvent(__a0); }
	void cppgoBase_tabletEvent(QTabletEvent* __a0) { QMenu::tabletEvent(__a0); }
	void cppgoBase_dragEnterEvent(QDragEnterEvent* __a0) { QMenu::dragEnterEvent(__a0); }
	void cppgoBase_dragMoveEvent(QDragMoveEvent* __a0) { QMenu::dragMoveEvent(__a0); }
	void cppgoBase_dragLeaveEvent(QDragLeaveEvent* __a0) { QMenu::dragLeaveEvent(__a0); }
	void cppgoBase_dropEvent(QDropEvent* __a0) { QMenu::dropEvent(__a0); }
	void cppgoBase_showEvent(QShowEvent* __a0) { QMenu::showEvent(__a0); }
	int cppgoBase_metric(int __a0) { return QMenu::metric((QPaintDevice::PaintDeviceMetric)__a0); }
	void cppgoBase_initPainter(QPainter* __a0) { QMenu::initPainter(__a0); }
	QPaintDevice* cppgoBase_redirected(QPoint* __a0) { return QMenu::redirected(__a0); }
	QPainter* cppgoBase_sharedPainter() { return QMenu::sharedPainter(); }
	void cppgoBase_inputMethodEvent(QInputMethodEvent* __a0) { QMenu::inputMethodEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QMenu::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QMenu::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QMenu::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QMenu::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QMenu::metaObject();
	}
	QSize sizeHint() const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 729690261457684101LL, &__h); if (__h) return *(QSize*)__r; }
		return QMenu::sizeHint();
	}
	void changeEvent(QEvent* arg1) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4959376063826165943LL, (void*)arg1, &__h); if (__h) return; }
		QMenu::changeEvent(arg1);
	}
	void keyPressEvent(QKeyEvent* arg1) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1023791977679793185LL, (void*)arg1, &__h); if (__h) return; }
		QMenu::keyPressEvent(arg1);
	}
	void mouseReleaseEvent(QMouseEvent* arg1) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4415920049297795213LL, (void*)arg1, &__h); if (__h) return; }
		QMenu::mouseReleaseEvent(arg1);
	}
	void mousePressEvent(QMouseEvent* arg1) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -332090898078543319LL, (void*)arg1, &__h); if (__h) return; }
		QMenu::mousePressEvent(arg1);
	}
	void mouseMoveEvent(QMouseEvent* arg1) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4255494004099806581LL, (void*)arg1, &__h); if (__h) return; }
		QMenu::mouseMoveEvent(arg1);
	}
	void wheelEvent(QWheelEvent* arg1) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1919988340934511316LL, (void*)arg1, &__h); if (__h) return; }
		QMenu::wheelEvent(arg1);
	}
	void enterEvent(QEnterEvent* arg1) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6316270346668237895LL, (void*)arg1, &__h); if (__h) return; }
		QMenu::enterEvent(arg1);
	}
	void leaveEvent(QEvent* arg1) override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -551070009623711640LL, (void*)arg1, &__h); if (__h) return; }
		QMenu::leaveEvent(arg1);
	}
	void hideEvent(QHideEvent* arg1) override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -199718707876719129LL, (void*)arg1, &__h); if (__h) return; }
		QMenu::hideEvent(arg1);
	}
	void paintEvent(QPaintEvent* arg1) override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5290442163167458465LL, (void*)arg1, &__h); if (__h) return; }
		QMenu::paintEvent(arg1);
	}
	void actionEvent(QActionEvent* arg1) override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1697444140088555581LL, (void*)arg1, &__h); if (__h) return; }
		QMenu::actionEvent(arg1);
	}
	void timerEvent(QTimerEvent* arg1) override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)arg1, &__h); if (__h) return; }
		QMenu::timerEvent(arg1);
	}
	bool event(QEvent* arg1) override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)arg1, &__h); if (__h) return __r; }
		return QMenu::event(arg1);
	}
	bool focusNextPrevChild(bool next) override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 5898328517505784191LL, next, &__h); if (__h) return __r; }
		return QMenu::focusNextPrevChild(next);
	}
	void initStyleOption(QStyleOptionMenuItem* option, const QAction* action) const override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_void_voidp_voidp) { ((void(*)(long long, long long, void*, void*, int*))cppgo_virt_void_voidp_voidp)(__goId, -488693769771308531LL, (void*)option, (void*)action, &__h); if (__h) return; }
		QMenu::initStyleOption(option, action);
	}
	int devType() const override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8659679893868524044LL, &__h); if (__h) return __r; }
		return QMenu::devType();
	}
	void setVisible(bool visible) override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 2110182719375245575LL, visible, &__h); if (__h) return; }
		QMenu::setVisible(visible);
	}
	QSize minimumSizeHint() const override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 7049264599837532491LL, &__h); if (__h) return *(QSize*)__r; }
		return QMenu::minimumSizeHint();
	}
	int heightForWidth(int arg1) const override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 4311534060495737339LL, arg1, &__h); if (__h) return __r; }
		return QMenu::heightForWidth(arg1);
	}
	bool hasHeightForWidth() const override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8266762397995423915LL, &__h); if (__h) return __r; }
		return QMenu::hasHeightForWidth();
	}
	QPaintEngine* paintEngine() const override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 6398007404756346381LL, &__h); if (__h) return (QPaintEngine*)__r; }
		return QMenu::paintEngine();
	}
	void mouseDoubleClickEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4029073905697023887LL, (void*)event, &__h); if (__h) return; }
		QMenu::mouseDoubleClickEvent(event);
	}
	void keyReleaseEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 772598475950050299LL, (void*)event, &__h); if (__h) return; }
		QMenu::keyReleaseEvent(event);
	}
	void focusInEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 975043767033552552LL, (void*)event, &__h); if (__h) return; }
		QMenu::focusInEvent(event);
	}
	void focusOutEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -385596010576048993LL, (void*)event, &__h); if (__h) return; }
		QMenu::focusOutEvent(event);
	}
	void moveEvent(QMoveEvent* event) override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -9161949150489525396LL, (void*)event, &__h); if (__h) return; }
		QMenu::moveEvent(event);
	}
	void resizeEvent(QResizeEvent* event) override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4302917285655014661LL, (void*)event, &__h); if (__h) return; }
		QMenu::resizeEvent(event);
	}
	void closeEvent(QCloseEvent* event) override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1989937050019464783LL, (void*)event, &__h); if (__h) return; }
		QMenu::closeEvent(event);
	}
	void contextMenuEvent(QContextMenuEvent* event) override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6698043222227006015LL, (void*)event, &__h); if (__h) return; }
		QMenu::contextMenuEvent(event);
	}
	void tabletEvent(QTabletEvent* event) override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -212068866860701141LL, (void*)event, &__h); if (__h) return; }
		QMenu::tabletEvent(event);
	}
	void dragEnterEvent(QDragEnterEvent* event) override {
		int __h = 0;
		if (__goOn[32] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 206122715919917589LL, (void*)event, &__h); if (__h) return; }
		QMenu::dragEnterEvent(event);
	}
	void dragMoveEvent(QDragMoveEvent* event) override {
		int __h = 0;
		if (__goOn[33] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1781545137352745978LL, (void*)event, &__h); if (__h) return; }
		QMenu::dragMoveEvent(event);
	}
	void dragLeaveEvent(QDragLeaveEvent* event) override {
		int __h = 0;
		if (__goOn[34] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -3139904158373725006LL, (void*)event, &__h); if (__h) return; }
		QMenu::dragLeaveEvent(event);
	}
	void dropEvent(QDropEvent* event) override {
		int __h = 0;
		if (__goOn[35] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 9107058682754591716LL, (void*)event, &__h); if (__h) return; }
		QMenu::dropEvent(event);
	}
	void showEvent(QShowEvent* event) override {
		int __h = 0;
		if (__goOn[36] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6825369174751916LL, (void*)event, &__h); if (__h) return; }
		QMenu::showEvent(event);
	}
	bool nativeEvent(const QByteArray& eventType, void* message, qintptr* result) override {
		int __h = 0;
		if (__goOn[37] && cppgo_virt_int_voidp_voidp_longlongp) { int __r = ((int(*)(long long, long long, void*, void*, long long*, int*))cppgo_virt_int_voidp_voidp_longlongp)(__goId, -2039166586246443706LL, (void*)&eventType, (void*)message, (long long*)result, &__h); if (__h) return __r; }
		return QMenu::nativeEvent(eventType, message, result);
	}
	int metric(QPaintDevice::PaintDeviceMetric arg1) const override {
		int __h = 0;
		if (__goOn[38] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 427629649267311777LL, (int)(arg1), &__h); if (__h) return __r; }
		return QMenu::metric(arg1);
	}
	void initPainter(QPainter* painter) const override {
		int __h = 0;
		if (__goOn[39] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6477500296980050210LL, (void*)painter, &__h); if (__h) return; }
		QMenu::initPainter(painter);
	}
	QPaintDevice* redirected(QPoint* offset) const override {
		int __h = 0;
		if (__goOn[40] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 4127240633310655048LL, (void*)offset, &__h); if (__h) return (QPaintDevice*)__r; }
		return QMenu::redirected(offset);
	}
	QPainter* sharedPainter() const override {
		int __h = 0;
		if (__goOn[41] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 5792771328741774991LL, &__h); if (__h) return (QPainter*)__r; }
		return QMenu::sharedPainter();
	}
	void inputMethodEvent(QInputMethodEvent* arg1) override {
		int __h = 0;
		if (__goOn[42] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5438819661942843222LL, (void*)arg1, &__h); if (__h) return; }
		QMenu::inputMethodEvent(arg1);
	}
	QVariant inputMethodQuery(Qt::InputMethodQuery arg1) const override {
		int __h = 0;
		if (__goOn[43] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7345853097340709638LL, (int)(arg1), &__h); if (__h) return *(QVariant*)__r; }
		return QMenu::inputMethodQuery(arg1);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[44] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QMenu::eventFilter(watched, event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[45] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QMenu::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[46] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QMenu::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[47] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QMenu::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[48] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QMenu::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QMenu;
static std::mutex __connsMu_QMenu;

extern "C" {

void* QMenu_base(void* self) { return (void*)static_cast<QWidget*>((QMenu*)self); }
void QMenu_chain(void* self, void** out) {
	QMenu* __p = (QMenu*)self;
	QWidget* __b0 = static_cast<QWidget*>(__p); out[0] = (void*)__b0;
	QObject* __b1 = static_cast<QObject*>(__b0); out[1] = (void*)__b1;
}
void QMenu__set_override(void* self, int idx) { ((QMenu__Go*)self)->__goOn[idx] = 1; }
void* QMenu_fromQObject(void* o) { return (void*)qobject_cast<QMenu*>((QObject*)o); }
void* QMenu_new(long long __goid, void* parent) {
	return new QMenu__Go(__goid, (QWidget*)parent);
}
void* QMenu_new2(long long __goid) {
	return new QMenu__Go(__goid);
}
void* QMenu_new3(long long __goid, const char* title, long long title_len) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	return new QMenu__Go(__goid, title_qstr);
}
void* QMenu_new4(long long __goid, const char* title, long long title_len, void* parent) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	return new QMenu__Go(__goid, title_qstr, (QWidget*)parent);
}
char* QMenu_tr(const char* s, long long* __out_len) {
	auto&& __r = QMenu::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QMenu_addAction(void* self, const char* text, long long text_len, void* receiver, const char* member, void* shortcut) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = ((QMenu*)self)->addAction(text_qstr, (QObject*)receiver, member, *(QKeySequence*)shortcut);
	return (void*)(__r);
}
void* QMenu_addAction2(void* self, void* icon, const char* text, long long text_len, void* receiver, const char* member, void* shortcut) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = ((QMenu*)self)->addAction(*(QIcon*)icon, text_qstr, (QObject*)receiver, member, *(QKeySequence*)shortcut);
	return (void*)(__r);
}
void* QMenu_addMenu(void* self, void* menu) {
	auto&& __r = ((QMenu*)self)->addMenu((QMenu*)menu);
	return (void*)(__r);
}
void* QMenu_addMenu2(void* self, const char* title, long long title_len) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	auto&& __r = ((QMenu*)self)->addMenu(title_qstr);
	return (void*)(__r);
}
void* QMenu_addMenu3(void* self, void* icon, const char* title, long long title_len) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	auto&& __r = ((QMenu*)self)->addMenu(*(QIcon*)icon, title_qstr);
	return (void*)(__r);
}
void* QMenu_addSeparator(void* self) {
	auto&& __r = ((QMenu*)self)->addSeparator();
	return (void*)(__r);
}
void* QMenu_addSection(void* self, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = ((QMenu*)self)->addSection(text_qstr);
	return (void*)(__r);
}
void* QMenu_addSection2(void* self, void* icon, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = ((QMenu*)self)->addSection(*(QIcon*)icon, text_qstr);
	return (void*)(__r);
}
void* QMenu_insertMenu(void* self, void* before, void* menu) {
	auto&& __r = ((QMenu*)self)->insertMenu((QAction*)before, (QMenu*)menu);
	return (void*)(__r);
}
void* QMenu_insertSeparator(void* self, void* before) {
	auto&& __r = ((QMenu*)self)->insertSeparator((QAction*)before);
	return (void*)(__r);
}
void* QMenu_insertSection(void* self, void* before, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = ((QMenu*)self)->insertSection((QAction*)before, text_qstr);
	return (void*)(__r);
}
void* QMenu_insertSection2(void* self, void* before, void* icon, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = ((QMenu*)self)->insertSection((QAction*)before, *(QIcon*)icon, text_qstr);
	return (void*)(__r);
}
int QMenu_isEmpty(void* self) {
	auto&& __r = ((QMenu*)self)->isEmpty();
	return __r;
}
void QMenu_clear(void* self) {
	((QMenu*)self)->clear();
}
void QMenu_setTearOffEnabled(void* self, int arg1) {
	((QMenu*)self)->setTearOffEnabled(arg1);
}
int QMenu_isTearOffEnabled(void* self) {
	auto&& __r = ((QMenu*)self)->isTearOffEnabled();
	return __r;
}
int QMenu_isTearOffMenuVisible(void* self) {
	auto&& __r = ((QMenu*)self)->isTearOffMenuVisible();
	return __r;
}
void QMenu_showTearOffMenu(void* self) {
	((QMenu*)self)->showTearOffMenu();
}
void QMenu_showTearOffMenu2(void* self, void* pos) {
	((QMenu*)self)->showTearOffMenu(*(QPoint*)pos);
}
void QMenu_hideTearOffMenu(void* self) {
	((QMenu*)self)->hideTearOffMenu();
}
void QMenu_setDefaultAction(void* self, void* arg1) {
	((QMenu*)self)->setDefaultAction((QAction*)arg1);
}
void* QMenu_defaultAction(void* self) {
	auto&& __r = ((QMenu*)self)->defaultAction();
	return (void*)(__r);
}
void QMenu_setActiveAction(void* self, void* act) {
	((QMenu*)self)->setActiveAction((QAction*)act);
}
void* QMenu_activeAction(void* self) {
	auto&& __r = ((QMenu*)self)->activeAction();
	return (void*)(__r);
}
void QMenu_popup(void* self, void* pos) {
	((QMenu*)self)->popup(*(QPoint*)pos);
}
void* QMenu_exec(void* self) {
	auto&& __r = ((QMenu*)self)->exec();
	return (void*)(__r);
}
void* QMenu_exec2(void* self, void* pos) {
	auto&& __r = ((QMenu*)self)->exec(*(QPoint*)pos);
	return (void*)(__r);
}
void* QMenu_exec3(void** actions, long long actions_len, void* pos) {
	QList<QAction*> actions_list; actions_list.reserve(actions_len);
	for (long long __i = 0; __i < actions_len; __i++) actions_list.append((QAction*)actions[__i]);
	auto&& __r = QMenu::exec(actions_list, *(QPoint*)pos);
	return (void*)(__r);
}
void* QMenu_actionGeometry(void* self, void* arg1) {
	auto&& __r = ((QMenu*)self)->actionGeometry((QAction*)arg1);
	return new QRect(__r);
}
void* QMenu_actionAt(void* self, void* arg1) {
	auto&& __r = ((QMenu*)self)->actionAt(*(QPoint*)arg1);
	return (void*)(__r);
}
void* QMenu_menuAction(void* self) {
	auto&& __r = ((QMenu*)self)->menuAction();
	return (void*)(__r);
}
void* QMenu_menuInAction(void* action) {
	auto&& __r = QMenu::menuInAction((QAction*)action);
	return (void*)(__r);
}
char* QMenu_title(void* self, long long* __out_len) {
	auto&& __r = ((QMenu*)self)->title();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QMenu_setTitle(void* self, const char* title, long long title_len) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	((QMenu*)self)->setTitle(title_qstr);
}
void* QMenu_icon(void* self) {
	auto&& __r = ((QMenu*)self)->icon();
	return new QIcon(__r);
}
void QMenu_setIcon(void* self, void* icon) {
	((QMenu*)self)->setIcon(*(QIcon*)icon);
}
void QMenu_setNoReplayFor(void* self, void* widget) {
	((QMenu*)self)->setNoReplayFor((QWidget*)widget);
}
int QMenu_separatorsCollapsible(void* self) {
	auto&& __r = ((QMenu*)self)->separatorsCollapsible();
	return __r;
}
void QMenu_setSeparatorsCollapsible(void* self, int collapse) {
	((QMenu*)self)->setSeparatorsCollapsible(collapse);
}
int QMenu_toolTipsVisible(void* self) {
	auto&& __r = ((QMenu*)self)->toolTipsVisible();
	return __r;
}
void QMenu_setToolTipsVisible(void* self, int visible) {
	((QMenu*)self)->setToolTipsVisible(visible);
}
void QMenu_aboutToShow(void* self) {
	((QMenu*)self)->aboutToShow();
}
void QMenu_aboutToHide(void* self) {
	((QMenu*)self)->aboutToHide();
}
void QMenu_triggered(void* self, void* action) {
	((QMenu*)self)->triggered((QAction*)action);
}
void QMenu_hovered(void* self, void* action) {
	((QMenu*)self)->hovered((QAction*)action);
}
int QMenu_columnCount(void* self) {
	auto&& __r = ((QMenu__Go*)self)->columnCount();
	return __r;
}
void QMenu_setTabOrder(void* arg1, void* arg2) {
	QMenu::setTabOrder((QWidget*)arg1, (QWidget*)arg2);
}
void* QMenu_mouseGrabber() {
	auto&& __r = QMenu::mouseGrabber();
	return (void*)(__r);
}
void* QMenu_keyboardGrabber() {
	auto&& __r = QMenu::keyboardGrabber();
	return (void*)(__r);
}
void* QMenu_find(unsigned long long arg1) {
	auto&& __r = QMenu::find(arg1);
	return (void*)(__r);
}
void* QMenu_createWindowContainer(void* window) {
	auto&& __r = QMenu::createWindowContainer((QWindow*)window);
	return (void*)(__r);
}
void QMenu_updateMicroFocus(void* self) {
	((QMenu__Go*)self)->updateMicroFocus();
}
void QMenu_create(void* self) {
	((QMenu__Go*)self)->create();
}
void QMenu_destroy(void* self) {
	((QMenu__Go*)self)->destroy();
}
int QMenu_focusNextChild(void* self) {
	auto&& __r = ((QMenu__Go*)self)->focusNextChild();
	return __r;
}
int QMenu_focusPreviousChild(void* self) {
	auto&& __r = ((QMenu__Go*)self)->focusPreviousChild();
	return __r;
}
void* QMenu_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QMenu::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QMenu_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QMenu::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QMenu_sender(void* self) {
	auto&& __r = ((QMenu__Go*)self)->sender();
	return (void*)(__r);
}
int QMenu_senderSignalIndex(void* self) {
	auto&& __r = ((QMenu__Go*)self)->senderSignalIndex();
	return __r;
}
int QMenu_receivers(void* self, const char* signal) {
	auto&& __r = ((QMenu__Go*)self)->receivers(signal);
	return __r;
}
int QMenu_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QMenu__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
double QMenu_devicePixelRatioFScale() {
	auto&& __r = QMenu::devicePixelRatioFScale();
	return __r;
}
int QMenu_encodeMetricF(int metric, double value) {
	auto&& __r = QMenu::encodeMetricF((QPaintDevice::PaintDeviceMetric)(metric), value);
	return __r;
}
double QMenu_getDecodedMetricF(void* self, int metricA, int metricB) {
	auto&& __r = ((QMenu__Go*)self)->getDecodedMetricF((QPaintDevice::PaintDeviceMetric)(metricA), (QPaintDevice::PaintDeviceMetric)(metricB));
	return __r;
}
char* QMenu_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QMenu::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QMenu_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QMenu::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QMenu_popup2(void* self, void* pos, void* at) {
	((QMenu*)self)->popup(*(QPoint*)pos, (QAction*)at);
}
void* QMenu_exec4(void* self, void* pos, void* at) {
	auto&& __r = ((QMenu*)self)->exec(*(QPoint*)pos, (QAction*)at);
	return (void*)(__r);
}
void* QMenu_exec5(void** actions, long long actions_len, void* pos, void* at) {
	QList<QAction*> actions_list; actions_list.reserve(actions_len);
	for (long long __i = 0; __i < actions_len; __i++) actions_list.append((QAction*)actions[__i]);
	auto&& __r = QMenu::exec(actions_list, *(QPoint*)pos, (QAction*)at);
	return (void*)(__r);
}
void* QMenu_exec6(void** actions, long long actions_len, void* pos, void* at, void* parent) {
	QList<QAction*> actions_list; actions_list.reserve(actions_len);
	for (long long __i = 0; __i < actions_len; __i++) actions_list.append((QAction*)actions[__i]);
	auto&& __r = QMenu::exec(actions_list, *(QPoint*)pos, (QAction*)at, (QWidget*)parent);
	return (void*)(__r);
}
void* QMenu_createWindowContainer2(void* window, void* parent) {
	auto&& __r = QMenu::createWindowContainer((QWindow*)window, (QWidget*)parent);
	return (void*)(__r);
}
void* QMenu_createWindowContainer3(void* window, void* parent, int flags) {
	auto&& __r = QMenu::createWindowContainer((QWindow*)window, (QWidget*)parent, QFlags<Qt::WindowType>(flags));
	return (void*)(__r);
}
void QMenu_updateMicroFocus2(void* self, int query) {
	((QMenu__Go*)self)->updateMicroFocus((Qt::InputMethodQuery)(query));
}
void QMenu_create2(void* self, unsigned long long arg1) {
	((QMenu__Go*)self)->create(arg1);
}
void QMenu_create3(void* self, unsigned long long arg1, int initializeWindow) {
	((QMenu__Go*)self)->create(arg1, initializeWindow);
}
void QMenu_create4(void* self, unsigned long long arg1, int initializeWindow, int destroyOldWindow) {
	((QMenu__Go*)self)->create(arg1, initializeWindow, destroyOldWindow);
}
void QMenu_destroy2(void* self, int destroyWindow) {
	((QMenu__Go*)self)->destroy(destroyWindow);
}
void QMenu_destroy3(void* self, int destroyWindow, int destroySubWindows) {
	((QMenu__Go*)self)->destroy(destroyWindow, destroySubWindows);
}
void* QMenu_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QMenu::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QMenu_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QMenu::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QMenu_metaObject_default(void* self) {
	auto&& __r = ((QMenu*)self)->QMenu::metaObject();
	return (void*)(__r);
}
void* QMenu_sizeHint_default(void* self) {
	auto&& __r = ((QMenu*)self)->QMenu::sizeHint();
	return new QSize(__r);
}
void QMenu_changeEvent_default(void* self, void* arg1) {
	((QMenu__Go*)self)->cppgoBase_changeEvent((QEvent*)arg1);
}
void QMenu_keyPressEvent_default(void* self, void* arg1) {
	((QMenu__Go*)self)->cppgoBase_keyPressEvent((QKeyEvent*)arg1);
}
void QMenu_mouseReleaseEvent_default(void* self, void* arg1) {
	((QMenu__Go*)self)->cppgoBase_mouseReleaseEvent((QMouseEvent*)arg1);
}
void QMenu_mousePressEvent_default(void* self, void* arg1) {
	((QMenu__Go*)self)->cppgoBase_mousePressEvent((QMouseEvent*)arg1);
}
void QMenu_mouseMoveEvent_default(void* self, void* arg1) {
	((QMenu__Go*)self)->cppgoBase_mouseMoveEvent((QMouseEvent*)arg1);
}
void QMenu_wheelEvent_default(void* self, void* arg1) {
	((QMenu__Go*)self)->cppgoBase_wheelEvent((QWheelEvent*)arg1);
}
void QMenu_enterEvent_default(void* self, void* arg1) {
	((QMenu__Go*)self)->cppgoBase_enterEvent((QEnterEvent*)arg1);
}
void QMenu_leaveEvent_default(void* self, void* arg1) {
	((QMenu__Go*)self)->cppgoBase_leaveEvent((QEvent*)arg1);
}
void QMenu_hideEvent_default(void* self, void* arg1) {
	((QMenu__Go*)self)->cppgoBase_hideEvent((QHideEvent*)arg1);
}
void QMenu_paintEvent_default(void* self, void* arg1) {
	((QMenu__Go*)self)->cppgoBase_paintEvent((QPaintEvent*)arg1);
}
void QMenu_actionEvent_default(void* self, void* arg1) {
	((QMenu__Go*)self)->cppgoBase_actionEvent((QActionEvent*)arg1);
}
void QMenu_timerEvent_default(void* self, void* arg1) {
	((QMenu__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)arg1);
}
int QMenu_event_default(void* self, void* arg1) {
	auto&& __r = ((QMenu__Go*)self)->cppgoBase_event((QEvent*)arg1);
	return __r;
}
int QMenu_focusNextPrevChild_default(void* self, int next) {
	auto&& __r = ((QMenu__Go*)self)->cppgoBase_focusNextPrevChild(next);
	return __r;
}
void QMenu_initStyleOption_default(void* self, void* option, void* action) {
	((QMenu__Go*)self)->cppgoBase_initStyleOption((QStyleOptionMenuItem*)option, (QAction*)action);
}
int QMenu_devType_default(void* self) {
	auto&& __r = ((QMenu*)self)->QMenu::devType();
	return __r;
}
void QMenu_setVisible_default(void* self, int visible) {
	((QMenu*)self)->QMenu::setVisible(visible);
}
void* QMenu_minimumSizeHint_default(void* self) {
	auto&& __r = ((QMenu*)self)->QMenu::minimumSizeHint();
	return new QSize(__r);
}
int QMenu_heightForWidth_default(void* self, int arg1) {
	auto&& __r = ((QMenu*)self)->QMenu::heightForWidth(arg1);
	return __r;
}
int QMenu_hasHeightForWidth_default(void* self) {
	auto&& __r = ((QMenu*)self)->QMenu::hasHeightForWidth();
	return __r;
}
void* QMenu_paintEngine_default(void* self) {
	auto&& __r = ((QMenu*)self)->QMenu::paintEngine();
	return (void*)(__r);
}
void QMenu_mouseDoubleClickEvent_default(void* self, void* event) {
	((QMenu__Go*)self)->cppgoBase_mouseDoubleClickEvent((QMouseEvent*)event);
}
void QMenu_keyReleaseEvent_default(void* self, void* event) {
	((QMenu__Go*)self)->cppgoBase_keyReleaseEvent((QKeyEvent*)event);
}
void QMenu_focusInEvent_default(void* self, void* event) {
	((QMenu__Go*)self)->cppgoBase_focusInEvent((QFocusEvent*)event);
}
void QMenu_focusOutEvent_default(void* self, void* event) {
	((QMenu__Go*)self)->cppgoBase_focusOutEvent((QFocusEvent*)event);
}
void QMenu_moveEvent_default(void* self, void* event) {
	((QMenu__Go*)self)->cppgoBase_moveEvent((QMoveEvent*)event);
}
void QMenu_resizeEvent_default(void* self, void* event) {
	((QMenu__Go*)self)->cppgoBase_resizeEvent((QResizeEvent*)event);
}
void QMenu_closeEvent_default(void* self, void* event) {
	((QMenu__Go*)self)->cppgoBase_closeEvent((QCloseEvent*)event);
}
void QMenu_contextMenuEvent_default(void* self, void* event) {
	((QMenu__Go*)self)->cppgoBase_contextMenuEvent((QContextMenuEvent*)event);
}
void QMenu_tabletEvent_default(void* self, void* event) {
	((QMenu__Go*)self)->cppgoBase_tabletEvent((QTabletEvent*)event);
}
void QMenu_dragEnterEvent_default(void* self, void* event) {
	((QMenu__Go*)self)->cppgoBase_dragEnterEvent((QDragEnterEvent*)event);
}
void QMenu_dragMoveEvent_default(void* self, void* event) {
	((QMenu__Go*)self)->cppgoBase_dragMoveEvent((QDragMoveEvent*)event);
}
void QMenu_dragLeaveEvent_default(void* self, void* event) {
	((QMenu__Go*)self)->cppgoBase_dragLeaveEvent((QDragLeaveEvent*)event);
}
void QMenu_dropEvent_default(void* self, void* event) {
	((QMenu__Go*)self)->cppgoBase_dropEvent((QDropEvent*)event);
}
void QMenu_showEvent_default(void* self, void* event) {
	((QMenu__Go*)self)->cppgoBase_showEvent((QShowEvent*)event);
}
int QMenu_metric_default(void* self, int arg1) {
	auto&& __r = ((QMenu__Go*)self)->cppgoBase_metric(arg1);
	return __r;
}
void QMenu_initPainter_default(void* self, void* painter) {
	((QMenu__Go*)self)->cppgoBase_initPainter((QPainter*)painter);
}
void* QMenu_redirected_default(void* self, void* offset) {
	auto&& __r = ((QMenu__Go*)self)->cppgoBase_redirected((QPoint*)offset);
	return (void*)(__r);
}
void* QMenu_sharedPainter_default(void* self) {
	auto&& __r = ((QMenu__Go*)self)->cppgoBase_sharedPainter();
	return (void*)(__r);
}
void QMenu_inputMethodEvent_default(void* self, void* arg1) {
	((QMenu__Go*)self)->cppgoBase_inputMethodEvent((QInputMethodEvent*)arg1);
}
void* QMenu_inputMethodQuery_default(void* self, int arg1) {
	auto&& __r = ((QMenu*)self)->QMenu::inputMethodQuery((Qt::InputMethodQuery)(arg1));
	return new QVariant(__r);
}
int QMenu_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QMenu*)self)->QMenu::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QMenu_childEvent_default(void* self, void* event) {
	((QMenu__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QMenu_customEvent_default(void* self, void* event) {
	((QMenu__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QMenu_connectNotify_default(void* self, void* signal) {
	((QMenu__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QMenu_disconnectNotify_default(void* self, void* signal) {
	((QMenu__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QMenu_connect_aboutToShow(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QMenu*)self, &QMenu::aboutToShow, (QMenu*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QMenu);
	__conns_QMenu[__id] = __c;
}
void QMenu_connect_aboutToHide(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QMenu*)self, &QMenu::aboutToHide, (QMenu*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QMenu);
	__conns_QMenu[__id] = __c;
}
void QMenu_connect_triggered(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QMenu*)self, &QMenu::triggered, (QMenu*)self, [__id](QAction* action){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)action);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QMenu);
	__conns_QMenu[__id] = __c;
}
void QMenu_connect_hovered(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QMenu*)self, &QMenu::hovered, (QMenu*)self, [__id](QAction* action){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)action);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QMenu);
	__conns_QMenu[__id] = __c;
}
void QMenu_connect_windowTitleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QMenu*)self, &QMenu::windowTitleChanged, (QMenu*)self, [__id](const QString& title){
		QByteArray title_ba = title.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)title_ba.constData(), title_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QMenu);
	__conns_QMenu[__id] = __c;
}
void QMenu_connect_windowIconChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QMenu*)self, &QMenu::windowIconChanged, (QMenu*)self, [__id](const QIcon& icon){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&icon);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QMenu);
	__conns_QMenu[__id] = __c;
}
void QMenu_connect_windowIconTextChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QMenu*)self, &QMenu::windowIconTextChanged, (QMenu*)self, [__id](const QString& iconText){
		QByteArray iconText_ba = iconText.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)iconText_ba.constData(), iconText_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QMenu);
	__conns_QMenu[__id] = __c;
}
void QMenu_connect_customContextMenuRequested(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QMenu*)self, &QMenu::customContextMenuRequested, (QMenu*)self, [__id](const QPoint& pos){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&pos);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QMenu);
	__conns_QMenu[__id] = __c;
}
void QMenu_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QMenu*)self, &QMenu::destroyed, (QMenu*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QMenu);
	__conns_QMenu[__id] = __c;
}
void QMenu_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QMenu*)self, &QMenu::destroyed, (QMenu*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QMenu);
	__conns_QMenu[__id] = __c;
}
void QMenu_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QMenu*)self, &QMenu::objectNameChanged, (QMenu*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QMenu);
	__conns_QMenu[__id] = __c;
}
void QMenu__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QMenu);
		auto __it = __conns_QMenu.find(id);
		if (__it == __conns_QMenu.end()) return;
		__c = __it->second;
		__conns_QMenu.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QMenu__Go_new(long long id) { return new QMenu__Go(id); }
void QMenu_delete(void* self) { delete (QMenu*)self; }

}

#include <QtWidgets/qstyleditemdelegate.h>
#include <QtWidgets/qabstractitemdelegate.h>
#include <QtCore/qobject.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtGui/qpainter.h>
#include <QtWidgets/qstyleoption.h>
#include <QtCore/qabstractitemmodel.h>
#include <QtCore/qsize.h>
#include <QtWidgets/qwidget.h>
#include <QtWidgets/qitemeditorfactory.h>
#include <QtCore/qvariant.h>
#include <QtCore/qlocale.h>
#include <QtCore/qcoreevent.h>
#include <QtGui/qevent.h>
#include <QtWidgets/qabstractitemview.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qstyleditemdelegate.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QStyledItemDelegate__Go : public QStyledItemDelegate {
public:
	long long __goId;
	unsigned char __goOn[20] = {};
	template<typename... A> QStyledItemDelegate__Go(long long id, A&&... a) : QStyledItemDelegate(std::forward<A>(a)...), __goId(id) {}
	using QStyledItemDelegate::sender;
	using QStyledItemDelegate::senderSignalIndex;
	using QStyledItemDelegate::receivers;
	using QStyledItemDelegate::isSignalConnected;
	void cppgoBase_initStyleOption(QStyleOptionViewItem* __a0, const QModelIndex& __a1) { QStyledItemDelegate::initStyleOption(__a0, __a1); }
	bool cppgoBase_eventFilter(QObject* __a0, QEvent* __a1) { return QStyledItemDelegate::eventFilter(__a0, __a1); }
	bool cppgoBase_editorEvent(QEvent* __a0, QAbstractItemModel* __a1, const QStyleOptionViewItem& __a2, const QModelIndex& __a3) { return QStyledItemDelegate::editorEvent(__a0, __a1, __a2, __a3); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QStyledItemDelegate::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QStyledItemDelegate::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QStyledItemDelegate::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QStyledItemDelegate::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QStyledItemDelegate::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QStyledItemDelegate::metaObject();
	}
	void paint(QPainter* painter, const QStyleOptionViewItem& option, const QModelIndex& index) const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_void_voidp_voidp_voidp) { ((void(*)(long long, long long, void*, void*, void*, int*))cppgo_virt_void_voidp_voidp_voidp)(__goId, 6083584261053556819LL, (void*)painter, (void*)&option, (void*)&index, &__h); if (__h) return; }
		QStyledItemDelegate::paint(painter, option, index);
	}
	QSize sizeHint(const QStyleOptionViewItem& option, const QModelIndex& index) const override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_voidp_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, void*, int*))cppgo_virt_voidp_voidp_voidp)(__goId, 729690261457684101LL, (void*)&option, (void*)&index, &__h); if (__h) return *(QSize*)__r; }
		return QStyledItemDelegate::sizeHint(option, index);
	}
	QWidget* createEditor(QWidget* parent, const QStyleOptionViewItem& option, const QModelIndex& index) const override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_voidp_voidp_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, void*, void*, int*))cppgo_virt_voidp_voidp_voidp_voidp)(__goId, -2118980437705941456LL, (void*)parent, (void*)&option, (void*)&index, &__h); if (__h) return (QWidget*)__r; }
		return QStyledItemDelegate::createEditor(parent, option, index);
	}
	void setEditorData(QWidget* editor, const QModelIndex& index) const override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_voidp_voidp) { ((void(*)(long long, long long, void*, void*, int*))cppgo_virt_void_voidp_voidp)(__goId, -858237687273351194LL, (void*)editor, (void*)&index, &__h); if (__h) return; }
		QStyledItemDelegate::setEditorData(editor, index);
	}
	void setModelData(QWidget* editor, QAbstractItemModel* model, const QModelIndex& index) const override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp_voidp_voidp) { ((void(*)(long long, long long, void*, void*, void*, int*))cppgo_virt_void_voidp_voidp_voidp)(__goId, -4740900545199884600LL, (void*)editor, (void*)model, (void*)&index, &__h); if (__h) return; }
		QStyledItemDelegate::setModelData(editor, model, index);
	}
	void updateEditorGeometry(QWidget* editor, const QStyleOptionViewItem& option, const QModelIndex& index) const override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp_voidp_voidp) { ((void(*)(long long, long long, void*, void*, void*, int*))cppgo_virt_void_voidp_voidp_voidp)(__goId, -1005009386014896055LL, (void*)editor, (void*)&option, (void*)&index, &__h); if (__h) return; }
		QStyledItemDelegate::updateEditorGeometry(editor, option, index);
	}
	QString displayText(const QVariant& value, const QLocale& locale) const override {
		int __h = 0;
		if (__goOn[7] && cppgo_virtp_charp_voidp_voidp) {
			long long __len = 0;
			char* __s = ((char*(*)(long long, long long, void*, void*, long long*, int*))cppgo_virtp_charp_voidp_voidp)(__goId, -6300405565848239506LL, (void*)&value, (void*)&locale, &__len, &__h);
			if (__h) { QString __r = QString::fromUtf8(__s, __len); free(__s); return __r; }
		}
		return QStyledItemDelegate::displayText(value, locale);
	}
	void initStyleOption(QStyleOptionViewItem* option, const QModelIndex& index) const override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void_voidp_voidp) { ((void(*)(long long, long long, void*, void*, int*))cppgo_virt_void_voidp_voidp)(__goId, -488693769771308531LL, (void*)option, (void*)&index, &__h); if (__h) return; }
		QStyledItemDelegate::initStyleOption(option, index);
	}
	bool eventFilter(QObject* object, QEvent* event) override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)object, (void*)event, &__h); if (__h) return __r; }
		return QStyledItemDelegate::eventFilter(object, event);
	}
	bool editorEvent(QEvent* event, QAbstractItemModel* model, const QStyleOptionViewItem& option, const QModelIndex& index) override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_int_voidp_voidp_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, void*, void*, int*))cppgo_virt_int_voidp_voidp_voidp_voidp)(__goId, -2261279149879782144LL, (void*)event, (void*)model, (void*)&option, (void*)&index, &__h); if (__h) return __r; }
		return QStyledItemDelegate::editorEvent(event, model, option, index);
	}
	void destroyEditor(QWidget* editor, const QModelIndex& index) const override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_void_voidp_voidp) { ((void(*)(long long, long long, void*, void*, int*))cppgo_virt_void_voidp_voidp)(__goId, -7699250917160078568LL, (void*)editor, (void*)&index, &__h); if (__h) return; }
		QStyledItemDelegate::destroyEditor(editor, index);
	}
	bool helpEvent(QHelpEvent* event, QAbstractItemView* view, const QStyleOptionViewItem& option, const QModelIndex& index) override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_int_voidp_voidp_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, void*, void*, int*))cppgo_virt_int_voidp_voidp_voidp_voidp)(__goId, 1696901855134526294LL, (void*)event, (void*)view, (void*)&option, (void*)&index, &__h); if (__h) return __r; }
		return QStyledItemDelegate::helpEvent(event, view, option, index);
	}
	QList<int> paintingRoles() const override {
		int __h = 0;
		if (__goOn[13] && cppgo_virtp_voidp) {
			long long __len = 0;
			void* __a = ((void*(*)(long long, long long, long long*, int*))cppgo_virtp_voidp)(__goId, -4420854320663935586LL, &__len, &__h);
			if (__h) {
				QList<int> __r; __r.reserve(__len);
				int* __aa = (int*)__a;
				for (long long __i = 0; __i < __len; __i++) __r.append(__aa[__i]);
				if (__a) free(__a);
				return __r;
			}
		}
		return QStyledItemDelegate::paintingRoles();
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QStyledItemDelegate::event(event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QStyledItemDelegate::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QStyledItemDelegate::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QStyledItemDelegate::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QStyledItemDelegate::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QStyledItemDelegate::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QStyledItemDelegate;
static std::mutex __connsMu_QStyledItemDelegate;

extern "C" {

void* QStyledItemDelegate_base(void* self) { return (void*)static_cast<QAbstractItemDelegate*>((QStyledItemDelegate*)self); }
void QStyledItemDelegate_chain(void* self, void** out) {
	QStyledItemDelegate* __p = (QStyledItemDelegate*)self;
	QAbstractItemDelegate* __b0 = static_cast<QAbstractItemDelegate*>(__p); out[0] = (void*)__b0;
	QObject* __b1 = static_cast<QObject*>(__b0); out[1] = (void*)__b1;
}
void QStyledItemDelegate__set_override(void* self, int idx) { ((QStyledItemDelegate__Go*)self)->__goOn[idx] = 1; }
void* QStyledItemDelegate_fromQObject(void* o) { return (void*)qobject_cast<QStyledItemDelegate*>((QObject*)o); }
void* QStyledItemDelegate_new(long long __goid) {
	return new QStyledItemDelegate__Go(__goid);
}
void* QStyledItemDelegate_new2(long long __goid, void* parent) {
	return new QStyledItemDelegate__Go(__goid, (QObject*)parent);
}
char* QStyledItemDelegate_tr(const char* s, long long* __out_len) {
	auto&& __r = QStyledItemDelegate::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QStyledItemDelegate_itemEditorFactory(void* self) {
	auto&& __r = ((QStyledItemDelegate*)self)->itemEditorFactory();
	return (void*)(__r);
}
void QStyledItemDelegate_setItemEditorFactory(void* self, void* factory) {
	((QStyledItemDelegate*)self)->setItemEditorFactory((QItemEditorFactory*)factory);
}
char* QStyledItemDelegate_displayText(void* self, void* value, void* locale, long long* __out_len) {
	auto&& __r = ((QStyledItemDelegate*)self)->displayText(*(QVariant*)value, *(QLocale*)locale);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QStyledItemDelegate_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QStyledItemDelegate::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QStyledItemDelegate_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QStyledItemDelegate::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QStyledItemDelegate_sender(void* self) {
	auto&& __r = ((QStyledItemDelegate__Go*)self)->sender();
	return (void*)(__r);
}
int QStyledItemDelegate_senderSignalIndex(void* self) {
	auto&& __r = ((QStyledItemDelegate__Go*)self)->senderSignalIndex();
	return __r;
}
int QStyledItemDelegate_receivers(void* self, const char* signal) {
	auto&& __r = ((QStyledItemDelegate__Go*)self)->receivers(signal);
	return __r;
}
int QStyledItemDelegate_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QStyledItemDelegate__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QStyledItemDelegate_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QStyledItemDelegate::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QStyledItemDelegate_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QStyledItemDelegate::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QStyledItemDelegate_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QStyledItemDelegate::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QStyledItemDelegate_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QStyledItemDelegate::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
char* QStyledItemDelegate_displayText_default(void* self, void* value, void* locale, long long* __out_len) {
	auto&& __r = ((QStyledItemDelegate*)self)->QStyledItemDelegate::displayText(*(QVariant*)value, *(QLocale*)locale);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QStyledItemDelegate_metaObject_default(void* self) {
	auto&& __r = ((QStyledItemDelegate*)self)->QStyledItemDelegate::metaObject();
	return (void*)(__r);
}
void QStyledItemDelegate_paint_default(void* self, void* painter, void* option, void* index) {
	((QStyledItemDelegate*)self)->QStyledItemDelegate::paint((QPainter*)painter, *(QStyleOptionViewItem*)option, *(QModelIndex*)index);
}
void* QStyledItemDelegate_sizeHint_default(void* self, void* option, void* index) {
	auto&& __r = ((QStyledItemDelegate*)self)->QStyledItemDelegate::sizeHint(*(QStyleOptionViewItem*)option, *(QModelIndex*)index);
	return new QSize(__r);
}
void* QStyledItemDelegate_createEditor_default(void* self, void* parent, void* option, void* index) {
	auto&& __r = ((QStyledItemDelegate*)self)->QStyledItemDelegate::createEditor((QWidget*)parent, *(QStyleOptionViewItem*)option, *(QModelIndex*)index);
	return (void*)(__r);
}
void QStyledItemDelegate_setEditorData_default(void* self, void* editor, void* index) {
	((QStyledItemDelegate*)self)->QStyledItemDelegate::setEditorData((QWidget*)editor, *(QModelIndex*)index);
}
void QStyledItemDelegate_setModelData_default(void* self, void* editor, void* model, void* index) {
	((QStyledItemDelegate*)self)->QStyledItemDelegate::setModelData((QWidget*)editor, (QAbstractItemModel*)model, *(QModelIndex*)index);
}
void QStyledItemDelegate_updateEditorGeometry_default(void* self, void* editor, void* option, void* index) {
	((QStyledItemDelegate*)self)->QStyledItemDelegate::updateEditorGeometry((QWidget*)editor, *(QStyleOptionViewItem*)option, *(QModelIndex*)index);
}
void QStyledItemDelegate_initStyleOption_default(void* self, void* option, void* index) {
	((QStyledItemDelegate__Go*)self)->cppgoBase_initStyleOption((QStyleOptionViewItem*)option, *(QModelIndex*)index);
}
int QStyledItemDelegate_eventFilter_default(void* self, void* object, void* event) {
	auto&& __r = ((QStyledItemDelegate__Go*)self)->cppgoBase_eventFilter((QObject*)object, (QEvent*)event);
	return __r;
}
int QStyledItemDelegate_editorEvent_default(void* self, void* event, void* model, void* option, void* index) {
	auto&& __r = ((QStyledItemDelegate__Go*)self)->cppgoBase_editorEvent((QEvent*)event, (QAbstractItemModel*)model, *(QStyleOptionViewItem*)option, *(QModelIndex*)index);
	return __r;
}
void QStyledItemDelegate_destroyEditor_default(void* self, void* editor, void* index) {
	((QStyledItemDelegate*)self)->QStyledItemDelegate::destroyEditor((QWidget*)editor, *(QModelIndex*)index);
}
int QStyledItemDelegate_helpEvent_default(void* self, void* event, void* view, void* option, void* index) {
	auto&& __r = ((QStyledItemDelegate*)self)->QStyledItemDelegate::helpEvent((QHelpEvent*)event, (QAbstractItemView*)view, *(QStyleOptionViewItem*)option, *(QModelIndex*)index);
	return __r;
}
int* QStyledItemDelegate_paintingRoles_default(void* self, long long* __out_len) {
	auto&& __r = ((QStyledItemDelegate*)self)->QStyledItemDelegate::paintingRoles();
	*__out_len = __r.size();
	int* __arr = (int*)malloc(sizeof(int) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = __e;
	return __arr;
}
int QStyledItemDelegate_event_default(void* self, void* event) {
	auto&& __r = ((QStyledItemDelegate*)self)->QStyledItemDelegate::event((QEvent*)event);
	return __r;
}
void QStyledItemDelegate_timerEvent_default(void* self, void* event) {
	((QStyledItemDelegate__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QStyledItemDelegate_childEvent_default(void* self, void* event) {
	((QStyledItemDelegate__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QStyledItemDelegate_customEvent_default(void* self, void* event) {
	((QStyledItemDelegate__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QStyledItemDelegate_connectNotify_default(void* self, void* signal) {
	((QStyledItemDelegate__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QStyledItemDelegate_disconnectNotify_default(void* self, void* signal) {
	((QStyledItemDelegate__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QStyledItemDelegate_connect_commitData(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStyledItemDelegate*)self, &QStyledItemDelegate::commitData, (QStyledItemDelegate*)self, [__id](QWidget* editor){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)editor);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStyledItemDelegate);
	__conns_QStyledItemDelegate[__id] = __c;
}
void QStyledItemDelegate_connect_closeEditor(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStyledItemDelegate*)self, &QStyledItemDelegate::closeEditor, (QStyledItemDelegate*)self, [__id](QWidget* editor, QAbstractItemDelegate::EndEditHint hint){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)editor);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStyledItemDelegate);
	__conns_QStyledItemDelegate[__id] = __c;
}
void QStyledItemDelegate_connect_closeEditor2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStyledItemDelegate*)self, &QStyledItemDelegate::closeEditor, (QStyledItemDelegate*)self, [__id](QWidget* editor, QAbstractItemDelegate::EndEditHint hint){
		if (cppgo_sig_void_voidp_int) ((void(*)(long long, void*, int))cppgo_sig_void_voidp_int)(__id, (void*)editor, (int)(hint));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStyledItemDelegate);
	__conns_QStyledItemDelegate[__id] = __c;
}
void QStyledItemDelegate_connect_sizeHintChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStyledItemDelegate*)self, &QStyledItemDelegate::sizeHintChanged, (QStyledItemDelegate*)self, [__id](const QModelIndex& arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStyledItemDelegate);
	__conns_QStyledItemDelegate[__id] = __c;
}
void QStyledItemDelegate_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStyledItemDelegate*)self, &QStyledItemDelegate::destroyed, (QStyledItemDelegate*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStyledItemDelegate);
	__conns_QStyledItemDelegate[__id] = __c;
}
void QStyledItemDelegate_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStyledItemDelegate*)self, &QStyledItemDelegate::destroyed, (QStyledItemDelegate*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStyledItemDelegate);
	__conns_QStyledItemDelegate[__id] = __c;
}
void QStyledItemDelegate_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStyledItemDelegate*)self, &QStyledItemDelegate::objectNameChanged, (QStyledItemDelegate*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStyledItemDelegate);
	__conns_QStyledItemDelegate[__id] = __c;
}
void QStyledItemDelegate__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QStyledItemDelegate);
		auto __it = __conns_QStyledItemDelegate.find(id);
		if (__it == __conns_QStyledItemDelegate.end()) return;
		__c = __it->second;
		__conns_QStyledItemDelegate.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QStyledItemDelegate__Go_new(long long id, void* parent) { return new QStyledItemDelegate__Go(id, (QObject*)parent); }
void QStyledItemDelegate_delete(void* self) { delete (QStyledItemDelegate*)self; }

}

