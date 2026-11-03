#include <QtCore/qflags.h>
#include "gen_qincompatibleflag.h"

extern "C" {

void* QIncompatibleFlag_new(int i) {
	return new QIncompatibleFlag(i);
}
void* QIncompatibleFlag_new2(void* arg1) {
	return new QIncompatibleFlag(*(QIncompatibleFlag*)arg1);
}
int QIncompatibleFlag_ToInt(void* self) {
	auto&& __r = static_cast<int>(*(QIncompatibleFlag*)self);
	return __r;
}
void QIncompatibleFlag_delete(void* self) { delete (QIncompatibleFlag*)self; }

}

#include <QtCore/qstringtokenizer.h>
#include "gen_qstringtokenizerbasebase.h"

extern "C" {

void* QStringTokenizerBaseBase_new(void* arg1) {
	return new QStringTokenizerBaseBase(*(QStringTokenizerBaseBase*)arg1);
}

}

#include <QtCore/qeventloop.h>
#include <QtCore/qobject.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qdeadlinetimer.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qvariant.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qeventloop.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QEventLoop__Go : public QEventLoop {
public:
	long long __goId;
	unsigned char __goOn[8] = {};
	template<typename... A> QEventLoop__Go(long long id, A&&... a) : QEventLoop(std::forward<A>(a)...), __goId(id) {}
	using QEventLoop::sender;
	using QEventLoop::senderSignalIndex;
	using QEventLoop::receivers;
	using QEventLoop::isSignalConnected;
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QEventLoop::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QEventLoop::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QEventLoop::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QEventLoop::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QEventLoop::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QEventLoop::metaObject();
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QEventLoop::event(event);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QEventLoop::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QEventLoop::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QEventLoop::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QEventLoop::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QEventLoop::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QEventLoop::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QEventLoop;
static std::mutex __connsMu_QEventLoop;

extern "C" {

void* QEventLoop_base(void* self) { return (void*)static_cast<QObject*>((QEventLoop*)self); }
void QEventLoop__set_override(void* self, int idx) { ((QEventLoop__Go*)self)->__goOn[idx] = 1; }
void* QEventLoop_fromQObject(void* o) { return (void*)qobject_cast<QEventLoop*>((QObject*)o); }
void* QEventLoop_new(long long __goid) {
	return new QEventLoop__Go(__goid);
}
void* QEventLoop_new2(long long __goid, void* parent) {
	return new QEventLoop__Go(__goid, (QObject*)parent);
}
char* QEventLoop_tr(const char* s, long long* __out_len) {
	auto&& __r = QEventLoop::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QEventLoop_processEvents(void* self) {
	auto&& __r = ((QEventLoop*)self)->processEvents();
	return __r;
}
void QEventLoop_processEvents2(void* self, int flags, int maximumTime) {
	((QEventLoop*)self)->processEvents(QFlags<QEventLoop::ProcessEventsFlag>(flags), maximumTime);
}
void QEventLoop_processEvents3(void* self, int flags, void* deadline) {
	((QEventLoop*)self)->processEvents(QFlags<QEventLoop::ProcessEventsFlag>(flags), *(QDeadlineTimer*)deadline);
}
int QEventLoop_exec(void* self) {
	auto&& __r = ((QEventLoop*)self)->exec();
	return __r;
}
int QEventLoop_isRunning(void* self) {
	auto&& __r = ((QEventLoop*)self)->isRunning();
	return __r;
}
void QEventLoop_wakeUp(void* self) {
	((QEventLoop*)self)->wakeUp();
}
void QEventLoop_exit(void* self) {
	((QEventLoop*)self)->exit();
}
void QEventLoop_quit(void* self) {
	((QEventLoop*)self)->quit();
}
void* QEventLoop_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QEventLoop::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QEventLoop_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QEventLoop::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QEventLoop_sender(void* self) {
	auto&& __r = ((QEventLoop__Go*)self)->sender();
	return (void*)(__r);
}
int QEventLoop_senderSignalIndex(void* self) {
	auto&& __r = ((QEventLoop__Go*)self)->senderSignalIndex();
	return __r;
}
int QEventLoop_receivers(void* self, const char* signal) {
	auto&& __r = ((QEventLoop__Go*)self)->receivers(signal);
	return __r;
}
int QEventLoop_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QEventLoop__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QEventLoop_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QEventLoop::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QEventLoop_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QEventLoop::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QEventLoop_processEvents4(void* self, int flags) {
	auto&& __r = ((QEventLoop*)self)->processEvents(QFlags<QEventLoop::ProcessEventsFlag>(flags));
	return __r;
}
int QEventLoop_exec2(void* self, int flags) {
	auto&& __r = ((QEventLoop*)self)->exec(QFlags<QEventLoop::ProcessEventsFlag>(flags));
	return __r;
}
void QEventLoop_exit2(void* self, int returnCode) {
	((QEventLoop*)self)->exit(returnCode);
}
void* QEventLoop_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QEventLoop::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QEventLoop_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QEventLoop::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QEventLoop_metaObject_default(void* self) {
	auto&& __r = ((QEventLoop*)self)->QEventLoop::metaObject();
	return (void*)(__r);
}
int QEventLoop_event_default(void* self, void* event) {
	auto&& __r = ((QEventLoop*)self)->QEventLoop::event((QEvent*)event);
	return __r;
}
int QEventLoop_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QEventLoop*)self)->QEventLoop::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QEventLoop_timerEvent_default(void* self, void* event) {
	((QEventLoop__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QEventLoop_childEvent_default(void* self, void* event) {
	((QEventLoop__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QEventLoop_customEvent_default(void* self, void* event) {
	((QEventLoop__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QEventLoop_connectNotify_default(void* self, void* signal) {
	((QEventLoop__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QEventLoop_disconnectNotify_default(void* self, void* signal) {
	((QEventLoop__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QEventLoop_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QEventLoop*)self, &QEventLoop::destroyed, (QEventLoop*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QEventLoop);
	__conns_QEventLoop[__id] = __c;
}
void QEventLoop_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QEventLoop*)self, &QEventLoop::destroyed, (QEventLoop*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QEventLoop);
	__conns_QEventLoop[__id] = __c;
}
void QEventLoop_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QEventLoop*)self, &QEventLoop::objectNameChanged, (QEventLoop*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QEventLoop);
	__conns_QEventLoop[__id] = __c;
}
void QEventLoop__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QEventLoop);
		auto __it = __conns_QEventLoop.find(id);
		if (__it == __conns_QEventLoop.end()) return;
		__c = __it->second;
		__conns_QEventLoop.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QEventLoop__Go_new(long long id, void* parent) { return new QEventLoop__Go(id, (QObject*)parent); }
void QEventLoop_delete(void* self) { delete (QEventLoop*)self; }

}

#include <QtCore/qabstractitemmodel.h>
#include <QtCore/qobject.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qmimedata.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qvariant.h>
#include <QtCore/qsize.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qdatastream.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qabstractlistmodel.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QAbstractListModel__Go : public QAbstractListModel {
public:
	long long __goId;
	unsigned char __goOn[39] = {};
	template<typename... A> QAbstractListModel__Go(long long id, A&&... a) : QAbstractListModel(std::forward<A>(a)...), __goId(id) {}
	using QAbstractListModel::createIndex;
	using QAbstractListModel::encodeData;
	using QAbstractListModel::decodeData;
	using QAbstractListModel::beginInsertRows;
	using QAbstractListModel::endInsertRows;
	using QAbstractListModel::beginRemoveRows;
	using QAbstractListModel::endRemoveRows;
	using QAbstractListModel::beginMoveRows;
	using QAbstractListModel::endMoveRows;
	using QAbstractListModel::beginInsertColumns;
	using QAbstractListModel::endInsertColumns;
	using QAbstractListModel::beginRemoveColumns;
	using QAbstractListModel::endRemoveColumns;
	using QAbstractListModel::beginMoveColumns;
	using QAbstractListModel::endMoveColumns;
	using QAbstractListModel::beginResetModel;
	using QAbstractListModel::endResetModel;
	using QAbstractListModel::changePersistentIndex;
	using QAbstractListModel::changePersistentIndexList;
	using QAbstractListModel::persistentIndexList;
	using QAbstractListModel::sender;
	using QAbstractListModel::senderSignalIndex;
	using QAbstractListModel::receivers;
	using QAbstractListModel::isSignalConnected;
	void cppgoBase_resetInternalData() { QAbstractListModel::resetInternalData(); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QAbstractListModel::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QAbstractListModel::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QAbstractListModel::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QAbstractListModel::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QAbstractListModel::disconnectNotify(__a0); }
	int rowCount(const QModelIndex& parent) const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1960960171956884010LL, (void*)&parent, &__h); if (__h) return __r; }
		return int{};
	}
	QVariant data(const QModelIndex& index, int role) const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_voidp_voidp_int) { void* __r = ((void*(*)(long long, long long, void*, int, int*))cppgo_virt_voidp_voidp_int)(__goId, -8837375942098466299LL, (void*)&index, role, &__h); if (__h) return *(QVariant*)__r; }
		return QVariant{};
	}
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QAbstractListModel::metaObject();
	}
	QModelIndex index(int row, int column, const QModelIndex& parent) const override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_voidp_int_int_voidp) { void* __r = ((void*(*)(long long, long long, int, int, void*, int*))cppgo_virt_voidp_int_int_voidp)(__goId, -8948777187306027381LL, row, column, (void*)&parent, &__h); if (__h) return *(QModelIndex*)__r; }
		return QAbstractListModel::index(row, column, parent);
	}
	QModelIndex sibling(int row, int column, const QModelIndex& idx) const override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_voidp_int_int_voidp) { void* __r = ((void*(*)(long long, long long, int, int, void*, int*))cppgo_virt_voidp_int_int_voidp)(__goId, 397329346813782521LL, row, column, (void*)&idx, &__h); if (__h) return *(QModelIndex*)__r; }
		return QAbstractListModel::sibling(row, column, idx);
	}
	bool dropMimeData(const QMimeData* data, Qt::DropAction action, int row, int column, const QModelIndex& parent) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_int_voidp_int_int_int_voidp) { int __r = ((int(*)(long long, long long, void*, int, int, int, void*, int*))cppgo_virt_int_voidp_int_int_int_voidp)(__goId, -9049994684305682212LL, (void*)data, (int)(action), row, column, (void*)&parent, &__h); if (__h) return __r; }
		return QAbstractListModel::dropMimeData(data, action, row, column, parent);
	}
	QFlags<Qt::ItemFlag> flags(const QModelIndex& index) const override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1703382833545304812LL, (void*)&index, &__h); if (__h) return QFlags<Qt::ItemFlag>(__r); }
		return QAbstractListModel::flags(index);
	}
	bool setData(const QModelIndex& index, const QVariant& value, int role) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_int_voidp_voidp_int) { int __r = ((int(*)(long long, long long, void*, void*, int, int*))cppgo_virt_int_voidp_voidp_int)(__goId, 35191454575202219LL, (void*)&index, (void*)&value, role, &__h); if (__h) return __r; }
		return QAbstractListModel::setData(index, value, role);
	}
	QVariant headerData(int section, Qt::Orientation orientation, int role) const override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_voidp_int_int_int) { void* __r = ((void*(*)(long long, long long, int, int, int, int*))cppgo_virt_voidp_int_int_int)(__goId, -3297140021896693688LL, section, (int)(orientation), role, &__h); if (__h) return *(QVariant*)__r; }
		return QAbstractListModel::headerData(section, orientation, role);
	}
	bool setHeaderData(int section, Qt::Orientation orientation, const QVariant& value, int role) override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_int_int_int_voidp_int) { int __r = ((int(*)(long long, long long, int, int, void*, int, int*))cppgo_virt_int_int_int_voidp_int)(__goId, 4852200306590054850LL, section, (int)(orientation), (void*)&value, role, &__h); if (__h) return __r; }
		return QAbstractListModel::setHeaderData(section, orientation, value, role);
	}
	bool clearItemData(const QModelIndex& index) override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 8914190500520542409LL, (void*)&index, &__h); if (__h) return __r; }
		return QAbstractListModel::clearItemData(index);
	}
	QList<QString> mimeTypes() const override {
		int __h = 0;
		if (__goOn[11] && cppgo_virtp_voidp) {
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
		return QAbstractListModel::mimeTypes();
	}
	QMimeData* mimeData(const QList<QModelIndex>& indexes) const override {
		int __h = 0;
		std::vector<void*> indexes_v; indexes_v.reserve(indexes.size());
		for (long long __i = 0; __i < indexes.size(); __i++) indexes_v.push_back((void*)&indexes[__i]);
		if (__goOn[12] && cppgo_virt_voidp_voidpp_longlong) { void* __r = ((void*(*)(long long, long long, void**, long long, int*))cppgo_virt_voidp_voidpp_longlong)(__goId, 5175297091024974749LL, indexes_v.data(), (long long)indexes.size(), &__h); if (__h) return (QMimeData*)__r; }
		return QAbstractListModel::mimeData(indexes);
	}
	bool canDropMimeData(const QMimeData* data, Qt::DropAction action, int row, int column, const QModelIndex& parent) const override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_int_voidp_int_int_int_voidp) { int __r = ((int(*)(long long, long long, void*, int, int, int, void*, int*))cppgo_virt_int_voidp_int_int_int_voidp)(__goId, 731539533644753052LL, (void*)data, (int)(action), row, column, (void*)&parent, &__h); if (__h) return __r; }
		return QAbstractListModel::canDropMimeData(data, action, row, column, parent);
	}
	QFlags<Qt::DropAction> supportedDropActions() const override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -7784642176186821205LL, &__h); if (__h) return QFlags<Qt::DropAction>(__r); }
		return QAbstractListModel::supportedDropActions();
	}
	QFlags<Qt::DropAction> supportedDragActions() const override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, 3421601386088570762LL, &__h); if (__h) return QFlags<Qt::DropAction>(__r); }
		return QAbstractListModel::supportedDragActions();
	}
	bool insertRows(int row, int count, const QModelIndex& parent) override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_int_int_int_voidp) { int __r = ((int(*)(long long, long long, int, int, void*, int*))cppgo_virt_int_int_int_voidp)(__goId, -2775729825962395603LL, row, count, (void*)&parent, &__h); if (__h) return __r; }
		return QAbstractListModel::insertRows(row, count, parent);
	}
	bool insertColumns(int column, int count, const QModelIndex& parent) override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_int_int_int_voidp) { int __r = ((int(*)(long long, long long, int, int, void*, int*))cppgo_virt_int_int_int_voidp)(__goId, -8973293846230792237LL, column, count, (void*)&parent, &__h); if (__h) return __r; }
		return QAbstractListModel::insertColumns(column, count, parent);
	}
	bool removeRows(int row, int count, const QModelIndex& parent) override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_int_int_int_voidp) { int __r = ((int(*)(long long, long long, int, int, void*, int*))cppgo_virt_int_int_int_voidp)(__goId, 5104072549182222656LL, row, count, (void*)&parent, &__h); if (__h) return __r; }
		return QAbstractListModel::removeRows(row, count, parent);
	}
	bool removeColumns(int column, int count, const QModelIndex& parent) override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_int_int_int_voidp) { int __r = ((int(*)(long long, long long, int, int, void*, int*))cppgo_virt_int_int_int_voidp)(__goId, 2078548737056540420LL, column, count, (void*)&parent, &__h); if (__h) return __r; }
		return QAbstractListModel::removeColumns(column, count, parent);
	}
	bool moveRows(const QModelIndex& sourceParent, int sourceRow, int count, const QModelIndex& destinationParent, int destinationChild) override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_int_voidp_int_int_voidp_int) { int __r = ((int(*)(long long, long long, void*, int, int, void*, int, int*))cppgo_virt_int_voidp_int_int_voidp_int)(__goId, -6074819356947405415LL, (void*)&sourceParent, sourceRow, count, (void*)&destinationParent, destinationChild, &__h); if (__h) return __r; }
		return QAbstractListModel::moveRows(sourceParent, sourceRow, count, destinationParent, destinationChild);
	}
	bool moveColumns(const QModelIndex& sourceParent, int sourceColumn, int count, const QModelIndex& destinationParent, int destinationChild) override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_int_voidp_int_int_voidp_int) { int __r = ((int(*)(long long, long long, void*, int, int, void*, int, int*))cppgo_virt_int_voidp_int_int_voidp_int)(__goId, -6321007591388309905LL, (void*)&sourceParent, sourceColumn, count, (void*)&destinationParent, destinationChild, &__h); if (__h) return __r; }
		return QAbstractListModel::moveColumns(sourceParent, sourceColumn, count, destinationParent, destinationChild);
	}
	void fetchMore(const QModelIndex& parent) override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3553294593611130806LL, (void*)&parent, &__h); if (__h) return; }
		QAbstractListModel::fetchMore(parent);
	}
	bool canFetchMore(const QModelIndex& parent) const override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 8639262066473964150LL, (void*)&parent, &__h); if (__h) return __r; }
		return QAbstractListModel::canFetchMore(parent);
	}
	void sort(int column, Qt::SortOrder order) override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_void_int_int) { ((void(*)(long long, long long, int, int, int*))cppgo_virt_void_int_int)(__goId, 6912441694453382161LL, column, (int)(order), &__h); if (__h) return; }
		QAbstractListModel::sort(column, order);
	}
	QModelIndex buddy(const QModelIndex& index) const override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 6112860112248786499LL, (void*)&index, &__h); if (__h) return *(QModelIndex*)__r; }
		return QAbstractListModel::buddy(index);
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
		return QAbstractListModel::match(start, role, value, hits, flags);
	}
	QSize span(const QModelIndex& index) const override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, -8395342916587040543LL, (void*)&index, &__h); if (__h) return *(QSize*)__r; }
		return QAbstractListModel::span(index);
	}
	void multiData(const QModelIndex& index, QModelRoleDataSpan roleDataSpan) const override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_void_voidp_voidp) { ((void(*)(long long, long long, void*, void*, int*))cppgo_virt_void_voidp_voidp)(__goId, 731311060944096642LL, (void*)&index, (void*)&roleDataSpan, &__h); if (__h) return; }
		QAbstractListModel::multiData(index, roleDataSpan);
	}
	bool submit() override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -3788331537442685179LL, &__h); if (__h) return __r; }
		return QAbstractListModel::submit();
	}
	void revert() override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, -1522951607536117149LL, &__h); if (__h) return; }
		QAbstractListModel::revert();
	}
	void resetInternalData() override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, -4510675770784051183LL, &__h); if (__h) return; }
		QAbstractListModel::resetInternalData();
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[32] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QAbstractListModel::event(event);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[33] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QAbstractListModel::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[34] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QAbstractListModel::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[35] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QAbstractListModel::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[36] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QAbstractListModel::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[37] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QAbstractListModel::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[38] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QAbstractListModel::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QAbstractListModel;
static std::mutex __connsMu_QAbstractListModel;

extern "C" {

void* QAbstractListModel_base(void* self) { return (void*)static_cast<QAbstractItemModel*>((QAbstractListModel*)self); }
void QAbstractListModel_chain(void* self, void** out) {
	QAbstractListModel* __p = (QAbstractListModel*)self;
	QAbstractItemModel* __b0 = static_cast<QAbstractItemModel*>(__p); out[0] = (void*)__b0;
	QObject* __b1 = static_cast<QObject*>(__b0); out[1] = (void*)__b1;
}
void QAbstractListModel__set_override(void* self, int idx) { ((QAbstractListModel__Go*)self)->__goOn[idx] = 1; }
void* QAbstractListModel_fromQObject(void* o) { return (void*)qobject_cast<QAbstractListModel*>((QObject*)o); }
char* QAbstractListModel_tr(const char* s, long long* __out_len) {
	auto&& __r = QAbstractListModel::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QAbstractListModel_createIndex(void* self, int row, int column) {
	auto&& __r = ((QAbstractListModel__Go*)self)->createIndex(row, column);
	return new QModelIndex(__r);
}
void* QAbstractListModel_createIndex2(void* self, int row, int column, unsigned long long id) {
	auto&& __r = ((QAbstractListModel__Go*)self)->createIndex(row, column, id);
	return new QModelIndex(__r);
}
void QAbstractListModel_encodeData(void* self, void** indexes, long long indexes_len, void* stream) {
	QList<QModelIndex> indexes_list; indexes_list.reserve(indexes_len);
	for (long long __i = 0; __i < indexes_len; __i++) indexes_list.append(*(QModelIndex*)indexes[__i]);
	((QAbstractListModel__Go*)self)->encodeData(indexes_list, *(QDataStream*)stream);
}
int QAbstractListModel_decodeData(void* self, int row, int column, void* parent, void* stream) {
	auto&& __r = ((QAbstractListModel__Go*)self)->decodeData(row, column, *(QModelIndex*)parent, *(QDataStream*)stream);
	return __r;
}
void QAbstractListModel_beginInsertRows(void* self, void* parent, int first, int last) {
	((QAbstractListModel__Go*)self)->beginInsertRows(*(QModelIndex*)parent, first, last);
}
void QAbstractListModel_endInsertRows(void* self) {
	((QAbstractListModel__Go*)self)->endInsertRows();
}
void QAbstractListModel_beginRemoveRows(void* self, void* parent, int first, int last) {
	((QAbstractListModel__Go*)self)->beginRemoveRows(*(QModelIndex*)parent, first, last);
}
void QAbstractListModel_endRemoveRows(void* self) {
	((QAbstractListModel__Go*)self)->endRemoveRows();
}
int QAbstractListModel_beginMoveRows(void* self, void* sourceParent, int sourceFirst, int sourceLast, void* destinationParent, int destinationRow) {
	auto&& __r = ((QAbstractListModel__Go*)self)->beginMoveRows(*(QModelIndex*)sourceParent, sourceFirst, sourceLast, *(QModelIndex*)destinationParent, destinationRow);
	return __r;
}
void QAbstractListModel_endMoveRows(void* self) {
	((QAbstractListModel__Go*)self)->endMoveRows();
}
void QAbstractListModel_beginInsertColumns(void* self, void* parent, int first, int last) {
	((QAbstractListModel__Go*)self)->beginInsertColumns(*(QModelIndex*)parent, first, last);
}
void QAbstractListModel_endInsertColumns(void* self) {
	((QAbstractListModel__Go*)self)->endInsertColumns();
}
void QAbstractListModel_beginRemoveColumns(void* self, void* parent, int first, int last) {
	((QAbstractListModel__Go*)self)->beginRemoveColumns(*(QModelIndex*)parent, first, last);
}
void QAbstractListModel_endRemoveColumns(void* self) {
	((QAbstractListModel__Go*)self)->endRemoveColumns();
}
int QAbstractListModel_beginMoveColumns(void* self, void* sourceParent, int sourceFirst, int sourceLast, void* destinationParent, int destinationColumn) {
	auto&& __r = ((QAbstractListModel__Go*)self)->beginMoveColumns(*(QModelIndex*)sourceParent, sourceFirst, sourceLast, *(QModelIndex*)destinationParent, destinationColumn);
	return __r;
}
void QAbstractListModel_endMoveColumns(void* self) {
	((QAbstractListModel__Go*)self)->endMoveColumns();
}
void QAbstractListModel_beginResetModel(void* self) {
	((QAbstractListModel__Go*)self)->beginResetModel();
}
void QAbstractListModel_endResetModel(void* self) {
	((QAbstractListModel__Go*)self)->endResetModel();
}
void QAbstractListModel_changePersistentIndex(void* self, void* from, void* to) {
	((QAbstractListModel__Go*)self)->changePersistentIndex(*(QModelIndex*)from, *(QModelIndex*)to);
}
void QAbstractListModel_changePersistentIndexList(void* self, void** from, long long from_len, void** to, long long to_len) {
	QList<QModelIndex> from_list; from_list.reserve(from_len);
	for (long long __i = 0; __i < from_len; __i++) from_list.append(*(QModelIndex*)from[__i]);
	QList<QModelIndex> to_list; to_list.reserve(to_len);
	for (long long __i = 0; __i < to_len; __i++) to_list.append(*(QModelIndex*)to[__i]);
	((QAbstractListModel__Go*)self)->changePersistentIndexList(from_list, to_list);
}
void** QAbstractListModel_persistentIndexList(void* self, long long* __out_len) {
	auto&& __r = ((QAbstractListModel__Go*)self)->persistentIndexList();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QModelIndex(__e);
	return __arr;
}
void* QAbstractListModel_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QAbstractListModel::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QAbstractListModel_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QAbstractListModel::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QAbstractListModel_sender(void* self) {
	auto&& __r = ((QAbstractListModel__Go*)self)->sender();
	return (void*)(__r);
}
int QAbstractListModel_senderSignalIndex(void* self) {
	auto&& __r = ((QAbstractListModel__Go*)self)->senderSignalIndex();
	return __r;
}
int QAbstractListModel_receivers(void* self, const char* signal) {
	auto&& __r = ((QAbstractListModel__Go*)self)->receivers(signal);
	return __r;
}
int QAbstractListModel_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QAbstractListModel__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QAbstractListModel_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QAbstractListModel::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QAbstractListModel_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QAbstractListModel::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QAbstractListModel_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QAbstractListModel::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QAbstractListModel_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QAbstractListModel::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QAbstractListModel_metaObject_default(void* self) {
	auto&& __r = ((QAbstractListModel*)self)->QAbstractListModel::metaObject();
	return (void*)(__r);
}
void* QAbstractListModel_index_default(void* self, int row, int column, void* parent) {
	auto&& __r = ((QAbstractListModel*)self)->QAbstractListModel::index(row, column, *(QModelIndex*)parent);
	return new QModelIndex(__r);
}
void* QAbstractListModel_sibling_default(void* self, int row, int column, void* idx) {
	auto&& __r = ((QAbstractListModel*)self)->QAbstractListModel::sibling(row, column, *(QModelIndex*)idx);
	return new QModelIndex(__r);
}
int QAbstractListModel_dropMimeData_default(void* self, void* data, int action, int row, int column, void* parent) {
	auto&& __r = ((QAbstractListModel*)self)->QAbstractListModel::dropMimeData((QMimeData*)data, (Qt::DropAction)(action), row, column, *(QModelIndex*)parent);
	return __r;
}
int QAbstractListModel_flags_default(void* self, void* index) {
	auto&& __r = ((QAbstractListModel*)self)->QAbstractListModel::flags(*(QModelIndex*)index);
	return (int)(__r);
}
int QAbstractListModel_setData_default(void* self, void* index, void* value, int role) {
	auto&& __r = ((QAbstractListModel*)self)->QAbstractListModel::setData(*(QModelIndex*)index, *(QVariant*)value, role);
	return __r;
}
void* QAbstractListModel_headerData_default(void* self, int section, int orientation, int role) {
	auto&& __r = ((QAbstractListModel*)self)->QAbstractListModel::headerData(section, (Qt::Orientation)(orientation), role);
	return new QVariant(__r);
}
int QAbstractListModel_setHeaderData_default(void* self, int section, int orientation, void* value, int role) {
	auto&& __r = ((QAbstractListModel*)self)->QAbstractListModel::setHeaderData(section, (Qt::Orientation)(orientation), *(QVariant*)value, role);
	return __r;
}
int QAbstractListModel_clearItemData_default(void* self, void* index) {
	auto&& __r = ((QAbstractListModel*)self)->QAbstractListModel::clearItemData(*(QModelIndex*)index);
	return __r;
}
char** QAbstractListModel_mimeTypes_default(void* self, long long* __out_len) {
	auto&& __r = ((QAbstractListModel*)self)->QAbstractListModel::mimeTypes();
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
void* QAbstractListModel_mimeData_default(void* self, void** indexes, long long indexes_len) {
	QList<QModelIndex> indexes_list; indexes_list.reserve(indexes_len);
	for (long long __i = 0; __i < indexes_len; __i++) indexes_list.append(*(QModelIndex*)indexes[__i]);
	auto&& __r = ((QAbstractListModel*)self)->QAbstractListModel::mimeData(indexes_list);
	return (void*)(__r);
}
int QAbstractListModel_canDropMimeData_default(void* self, void* data, int action, int row, int column, void* parent) {
	auto&& __r = ((QAbstractListModel*)self)->QAbstractListModel::canDropMimeData((QMimeData*)data, (Qt::DropAction)(action), row, column, *(QModelIndex*)parent);
	return __r;
}
int QAbstractListModel_supportedDropActions_default(void* self) {
	auto&& __r = ((QAbstractListModel*)self)->QAbstractListModel::supportedDropActions();
	return (int)(__r);
}
int QAbstractListModel_supportedDragActions_default(void* self) {
	auto&& __r = ((QAbstractListModel*)self)->QAbstractListModel::supportedDragActions();
	return (int)(__r);
}
int QAbstractListModel_insertRows_default(void* self, int row, int count, void* parent) {
	auto&& __r = ((QAbstractListModel*)self)->QAbstractListModel::insertRows(row, count, *(QModelIndex*)parent);
	return __r;
}
int QAbstractListModel_insertColumns_default(void* self, int column, int count, void* parent) {
	auto&& __r = ((QAbstractListModel*)self)->QAbstractListModel::insertColumns(column, count, *(QModelIndex*)parent);
	return __r;
}
int QAbstractListModel_removeRows_default(void* self, int row, int count, void* parent) {
	auto&& __r = ((QAbstractListModel*)self)->QAbstractListModel::removeRows(row, count, *(QModelIndex*)parent);
	return __r;
}
int QAbstractListModel_removeColumns_default(void* self, int column, int count, void* parent) {
	auto&& __r = ((QAbstractListModel*)self)->QAbstractListModel::removeColumns(column, count, *(QModelIndex*)parent);
	return __r;
}
int QAbstractListModel_moveRows_default(void* self, void* sourceParent, int sourceRow, int count, void* destinationParent, int destinationChild) {
	auto&& __r = ((QAbstractListModel*)self)->QAbstractListModel::moveRows(*(QModelIndex*)sourceParent, sourceRow, count, *(QModelIndex*)destinationParent, destinationChild);
	return __r;
}
int QAbstractListModel_moveColumns_default(void* self, void* sourceParent, int sourceColumn, int count, void* destinationParent, int destinationChild) {
	auto&& __r = ((QAbstractListModel*)self)->QAbstractListModel::moveColumns(*(QModelIndex*)sourceParent, sourceColumn, count, *(QModelIndex*)destinationParent, destinationChild);
	return __r;
}
void QAbstractListModel_fetchMore_default(void* self, void* parent) {
	((QAbstractListModel*)self)->QAbstractListModel::fetchMore(*(QModelIndex*)parent);
}
int QAbstractListModel_canFetchMore_default(void* self, void* parent) {
	auto&& __r = ((QAbstractListModel*)self)->QAbstractListModel::canFetchMore(*(QModelIndex*)parent);
	return __r;
}
void QAbstractListModel_sort_default(void* self, int column, int order) {
	((QAbstractListModel*)self)->QAbstractListModel::sort(column, (Qt::SortOrder)(order));
}
void* QAbstractListModel_buddy_default(void* self, void* index) {
	auto&& __r = ((QAbstractListModel*)self)->QAbstractListModel::buddy(*(QModelIndex*)index);
	return new QModelIndex(__r);
}
void** QAbstractListModel_match_default(void* self, void* start, int role, void* value, int hits, int flags, long long* __out_len) {
	auto&& __r = ((QAbstractListModel*)self)->QAbstractListModel::match(*(QModelIndex*)start, role, *(QVariant*)value, hits, QFlags<Qt::MatchFlag>(flags));
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QModelIndex(__e);
	return __arr;
}
void* QAbstractListModel_span_default(void* self, void* index) {
	auto&& __r = ((QAbstractListModel*)self)->QAbstractListModel::span(*(QModelIndex*)index);
	return new QSize(__r);
}
void QAbstractListModel_multiData_default(void* self, void* index, void* roleDataSpan) {
	((QAbstractListModel*)self)->QAbstractListModel::multiData(*(QModelIndex*)index, *(QModelRoleDataSpan*)roleDataSpan);
}
int QAbstractListModel_submit_default(void* self) {
	auto&& __r = ((QAbstractListModel*)self)->QAbstractListModel::submit();
	return __r;
}
void QAbstractListModel_revert_default(void* self) {
	((QAbstractListModel*)self)->QAbstractListModel::revert();
}
void QAbstractListModel_resetInternalData_default(void* self) {
	((QAbstractListModel__Go*)self)->cppgoBase_resetInternalData();
}
int QAbstractListModel_event_default(void* self, void* event) {
	auto&& __r = ((QAbstractListModel*)self)->QAbstractListModel::event((QEvent*)event);
	return __r;
}
int QAbstractListModel_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QAbstractListModel*)self)->QAbstractListModel::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QAbstractListModel_timerEvent_default(void* self, void* event) {
	((QAbstractListModel__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QAbstractListModel_childEvent_default(void* self, void* event) {
	((QAbstractListModel__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QAbstractListModel_customEvent_default(void* self, void* event) {
	((QAbstractListModel__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QAbstractListModel_connectNotify_default(void* self, void* signal) {
	((QAbstractListModel__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QAbstractListModel_disconnectNotify_default(void* self, void* signal) {
	((QAbstractListModel__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QAbstractListModel_connect_headerDataChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractListModel*)self, &QAbstractListModel::headerDataChanged, (QAbstractListModel*)self, [__id](Qt::Orientation orientation, int first, int last){
		if (cppgo_sig_void_int_int_int) ((void(*)(long long, int, int, int))cppgo_sig_void_int_int_int)(__id, (int)(orientation), first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractListModel);
	__conns_QAbstractListModel[__id] = __c;
}
void QAbstractListModel_connect_layoutChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractListModel*)self, &QAbstractListModel::layoutChanged, (QAbstractListModel*)self, [__id](const QList<QPersistentModelIndex>& parents, QAbstractItemModel::LayoutChangeHint hint){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractListModel);
	__conns_QAbstractListModel[__id] = __c;
}
void QAbstractListModel_connect_layoutChanged2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractListModel*)self, &QAbstractListModel::layoutChanged, (QAbstractListModel*)self, [__id](const QList<QPersistentModelIndex>& parents, QAbstractItemModel::LayoutChangeHint hint){
		std::vector<void*> parents_v; parents_v.reserve(parents.size());
		for (long long __i = 0; __i < parents.size(); __i++) parents_v.push_back((void*)&parents[__i]);
		if (cppgo_sig_void_voidpp_longlong_int) ((void(*)(long long, void**, long long, int))cppgo_sig_void_voidpp_longlong_int)(__id, parents_v.data(), (long long)parents.size(), (int)(hint));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractListModel);
	__conns_QAbstractListModel[__id] = __c;
}
void QAbstractListModel_connect_layoutAboutToBeChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractListModel*)self, &QAbstractListModel::layoutAboutToBeChanged, (QAbstractListModel*)self, [__id](const QList<QPersistentModelIndex>& parents, QAbstractItemModel::LayoutChangeHint hint){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractListModel);
	__conns_QAbstractListModel[__id] = __c;
}
void QAbstractListModel_connect_layoutAboutToBeChanged2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractListModel*)self, &QAbstractListModel::layoutAboutToBeChanged, (QAbstractListModel*)self, [__id](const QList<QPersistentModelIndex>& parents, QAbstractItemModel::LayoutChangeHint hint){
		std::vector<void*> parents_v; parents_v.reserve(parents.size());
		for (long long __i = 0; __i < parents.size(); __i++) parents_v.push_back((void*)&parents[__i]);
		if (cppgo_sig_void_voidpp_longlong_int) ((void(*)(long long, void**, long long, int))cppgo_sig_void_voidpp_longlong_int)(__id, parents_v.data(), (long long)parents.size(), (int)(hint));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractListModel);
	__conns_QAbstractListModel[__id] = __c;
}
void QAbstractListModel_connect_rowsAboutToBeInserted(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractListModel*)self, &QAbstractListModel::rowsAboutToBeInserted, (QAbstractListModel*)self, [__id](const QModelIndex& parent, int first, int last){
		if (cppgo_sig_void_voidp_int_int) ((void(*)(long long, void*, int, int))cppgo_sig_void_voidp_int_int)(__id, (void*)&parent, first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractListModel);
	__conns_QAbstractListModel[__id] = __c;
}
void QAbstractListModel_connect_rowsInserted(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractListModel*)self, &QAbstractListModel::rowsInserted, (QAbstractListModel*)self, [__id](const QModelIndex& parent, int first, int last){
		if (cppgo_sig_void_voidp_int_int) ((void(*)(long long, void*, int, int))cppgo_sig_void_voidp_int_int)(__id, (void*)&parent, first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractListModel);
	__conns_QAbstractListModel[__id] = __c;
}
void QAbstractListModel_connect_rowsAboutToBeRemoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractListModel*)self, &QAbstractListModel::rowsAboutToBeRemoved, (QAbstractListModel*)self, [__id](const QModelIndex& parent, int first, int last){
		if (cppgo_sig_void_voidp_int_int) ((void(*)(long long, void*, int, int))cppgo_sig_void_voidp_int_int)(__id, (void*)&parent, first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractListModel);
	__conns_QAbstractListModel[__id] = __c;
}
void QAbstractListModel_connect_rowsRemoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractListModel*)self, &QAbstractListModel::rowsRemoved, (QAbstractListModel*)self, [__id](const QModelIndex& parent, int first, int last){
		if (cppgo_sig_void_voidp_int_int) ((void(*)(long long, void*, int, int))cppgo_sig_void_voidp_int_int)(__id, (void*)&parent, first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractListModel);
	__conns_QAbstractListModel[__id] = __c;
}
void QAbstractListModel_connect_columnsAboutToBeInserted(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractListModel*)self, &QAbstractListModel::columnsAboutToBeInserted, (QAbstractListModel*)self, [__id](const QModelIndex& parent, int first, int last){
		if (cppgo_sig_void_voidp_int_int) ((void(*)(long long, void*, int, int))cppgo_sig_void_voidp_int_int)(__id, (void*)&parent, first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractListModel);
	__conns_QAbstractListModel[__id] = __c;
}
void QAbstractListModel_connect_columnsInserted(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractListModel*)self, &QAbstractListModel::columnsInserted, (QAbstractListModel*)self, [__id](const QModelIndex& parent, int first, int last){
		if (cppgo_sig_void_voidp_int_int) ((void(*)(long long, void*, int, int))cppgo_sig_void_voidp_int_int)(__id, (void*)&parent, first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractListModel);
	__conns_QAbstractListModel[__id] = __c;
}
void QAbstractListModel_connect_columnsAboutToBeRemoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractListModel*)self, &QAbstractListModel::columnsAboutToBeRemoved, (QAbstractListModel*)self, [__id](const QModelIndex& parent, int first, int last){
		if (cppgo_sig_void_voidp_int_int) ((void(*)(long long, void*, int, int))cppgo_sig_void_voidp_int_int)(__id, (void*)&parent, first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractListModel);
	__conns_QAbstractListModel[__id] = __c;
}
void QAbstractListModel_connect_columnsRemoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractListModel*)self, &QAbstractListModel::columnsRemoved, (QAbstractListModel*)self, [__id](const QModelIndex& parent, int first, int last){
		if (cppgo_sig_void_voidp_int_int) ((void(*)(long long, void*, int, int))cppgo_sig_void_voidp_int_int)(__id, (void*)&parent, first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractListModel);
	__conns_QAbstractListModel[__id] = __c;
}
void QAbstractListModel_connect_modelAboutToBeReset(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractListModel*)self, &QAbstractListModel::modelAboutToBeReset, (QAbstractListModel*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractListModel);
	__conns_QAbstractListModel[__id] = __c;
}
void QAbstractListModel_connect_modelReset(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractListModel*)self, &QAbstractListModel::modelReset, (QAbstractListModel*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractListModel);
	__conns_QAbstractListModel[__id] = __c;
}
void QAbstractListModel_connect_rowsAboutToBeMoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractListModel*)self, &QAbstractListModel::rowsAboutToBeMoved, (QAbstractListModel*)self, [__id](const QModelIndex& sourceParent, int sourceStart, int sourceEnd, const QModelIndex& destinationParent, int destinationRow){
		if (cppgo_sig_void_voidp_int_int_voidp_int) ((void(*)(long long, void*, int, int, void*, int))cppgo_sig_void_voidp_int_int_voidp_int)(__id, (void*)&sourceParent, sourceStart, sourceEnd, (void*)&destinationParent, destinationRow);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractListModel);
	__conns_QAbstractListModel[__id] = __c;
}
void QAbstractListModel_connect_rowsMoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractListModel*)self, &QAbstractListModel::rowsMoved, (QAbstractListModel*)self, [__id](const QModelIndex& sourceParent, int sourceStart, int sourceEnd, const QModelIndex& destinationParent, int destinationRow){
		if (cppgo_sig_void_voidp_int_int_voidp_int) ((void(*)(long long, void*, int, int, void*, int))cppgo_sig_void_voidp_int_int_voidp_int)(__id, (void*)&sourceParent, sourceStart, sourceEnd, (void*)&destinationParent, destinationRow);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractListModel);
	__conns_QAbstractListModel[__id] = __c;
}
void QAbstractListModel_connect_columnsAboutToBeMoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractListModel*)self, &QAbstractListModel::columnsAboutToBeMoved, (QAbstractListModel*)self, [__id](const QModelIndex& sourceParent, int sourceStart, int sourceEnd, const QModelIndex& destinationParent, int destinationColumn){
		if (cppgo_sig_void_voidp_int_int_voidp_int) ((void(*)(long long, void*, int, int, void*, int))cppgo_sig_void_voidp_int_int_voidp_int)(__id, (void*)&sourceParent, sourceStart, sourceEnd, (void*)&destinationParent, destinationColumn);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractListModel);
	__conns_QAbstractListModel[__id] = __c;
}
void QAbstractListModel_connect_columnsMoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractListModel*)self, &QAbstractListModel::columnsMoved, (QAbstractListModel*)self, [__id](const QModelIndex& sourceParent, int sourceStart, int sourceEnd, const QModelIndex& destinationParent, int destinationColumn){
		if (cppgo_sig_void_voidp_int_int_voidp_int) ((void(*)(long long, void*, int, int, void*, int))cppgo_sig_void_voidp_int_int_voidp_int)(__id, (void*)&sourceParent, sourceStart, sourceEnd, (void*)&destinationParent, destinationColumn);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractListModel);
	__conns_QAbstractListModel[__id] = __c;
}
void QAbstractListModel_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractListModel*)self, &QAbstractListModel::destroyed, (QAbstractListModel*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractListModel);
	__conns_QAbstractListModel[__id] = __c;
}
void QAbstractListModel_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractListModel*)self, &QAbstractListModel::destroyed, (QAbstractListModel*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractListModel);
	__conns_QAbstractListModel[__id] = __c;
}
void QAbstractListModel_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QAbstractListModel*)self, &QAbstractListModel::objectNameChanged, (QAbstractListModel*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QAbstractListModel);
	__conns_QAbstractListModel[__id] = __c;
}
void QAbstractListModel__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QAbstractListModel);
		auto __it = __conns_QAbstractListModel.find(id);
		if (__it == __conns_QAbstractListModel.end()) return;
		__c = __it->second;
		__conns_QAbstractListModel.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QAbstractListModel__Go_new(long long id, void* parent) { return new QAbstractListModel__Go(id, (QObject*)parent); }
void QAbstractListModel_delete(void* self) { delete (QAbstractListModel*)self; }

}

#include <QtCore/qcalendar.h>
#include "gen_qcalendar__systemid.h"

extern "C" {

void* QCalendar__SystemId_new() {
	return new QCalendar::SystemId();
}
unsigned long long QCalendar__SystemId_index(void* self) {
	auto&& __r = ((QCalendar::SystemId*)self)->index();
	return __r;
}
int QCalendar__SystemId_isValid(void* self) {
	auto&& __r = ((QCalendar::SystemId*)self)->isValid();
	return __r;
}
void QCalendar__SystemId_delete(void* self) { delete (QCalendar::SystemId*)self; }

}

#include <QtCore/qproperty.h>
#include <QtCore/qstring.h>
#include "gen_qpropertybindingerror.h"

extern "C" {

void* QPropertyBindingError_new() {
	return new QPropertyBindingError();
}
void* QPropertyBindingError_new2(int type_) {
	return new QPropertyBindingError((QPropertyBindingError::Type)(type_));
}
void* QPropertyBindingError_new3(void* other) {
	return new QPropertyBindingError(*(QPropertyBindingError*)other);
}
void* QPropertyBindingError_new4(int type_, const char* description, long long description_len) {
	QString description_qstr = QString::fromUtf8(description, description_len);
	return new QPropertyBindingError((QPropertyBindingError::Type)(type_), description_qstr);
}
void* QPropertyBindingError_OperatorAssign(void* self, void* other) {
	auto&& __r = ((QPropertyBindingError*)self)->operator=(*(QPropertyBindingError*)other);
	return (void*)&(__r);
}
int QPropertyBindingError_hasError(void* self) {
	auto&& __r = ((QPropertyBindingError*)self)->hasError();
	return __r;
}
int QPropertyBindingError_type(void* self) {
	auto&& __r = ((QPropertyBindingError*)self)->type();
	return (int)(__r);
}
char* QPropertyBindingError_description(void* self, long long* __out_len) {
	auto&& __r = ((QPropertyBindingError*)self)->description();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QPropertyBindingError_delete(void* self) { delete (QPropertyBindingError*)self; }

}

#include <QtCore/qdiriterator.h>
#include <QtCore/qdir.h>
#include <QtCore/qstring.h>
#include <QtCore/qfileinfo.h>
#include "gen_qdiriterator.h"

extern "C" {

void* QDirIterator_new(void* dir) {
	return new QDirIterator(*(QDir*)dir);
}
void* QDirIterator_new2(const char* path, long long path_len) {
	QString path_qstr = QString::fromUtf8(path, path_len);
	return new QDirIterator(path_qstr);
}
void* QDirIterator_new3(const char* path, long long path_len, int filter) {
	QString path_qstr = QString::fromUtf8(path, path_len);
	return new QDirIterator(path_qstr, QFlags<QDir::Filter>(filter));
}
void* QDirIterator_new4(const char* path, long long path_len, char** nameFilters, long long nameFilters_len) {
	QString path_qstr = QString::fromUtf8(path, path_len);
	QList<QString> nameFilters_list; nameFilters_list.reserve(nameFilters_len);
	for (long long __i = 0; __i < nameFilters_len; __i++) nameFilters_list.append(QString::fromUtf8(nameFilters[__i]));
	return new QDirIterator(path_qstr, nameFilters_list);
}
void* QDirIterator_new5(void* dir, int flags) {
	return new QDirIterator(*(QDir*)dir, QFlags<QDirIterator::IteratorFlag>(flags));
}
void* QDirIterator_new6(const char* path, long long path_len, int flags) {
	QString path_qstr = QString::fromUtf8(path, path_len);
	return new QDirIterator(path_qstr, QFlags<QDirIterator::IteratorFlag>(flags));
}
void* QDirIterator_new7(const char* path, long long path_len, int filter, int flags) {
	QString path_qstr = QString::fromUtf8(path, path_len);
	return new QDirIterator(path_qstr, QFlags<QDir::Filter>(filter), QFlags<QDirIterator::IteratorFlag>(flags));
}
void* QDirIterator_new8(const char* path, long long path_len, char** nameFilters, long long nameFilters_len, int filters) {
	QString path_qstr = QString::fromUtf8(path, path_len);
	QList<QString> nameFilters_list; nameFilters_list.reserve(nameFilters_len);
	for (long long __i = 0; __i < nameFilters_len; __i++) nameFilters_list.append(QString::fromUtf8(nameFilters[__i]));
	return new QDirIterator(path_qstr, nameFilters_list, QFlags<QDir::Filter>(filters));
}
void* QDirIterator_new9(const char* path, long long path_len, char** nameFilters, long long nameFilters_len, int filters, int flags) {
	QString path_qstr = QString::fromUtf8(path, path_len);
	QList<QString> nameFilters_list; nameFilters_list.reserve(nameFilters_len);
	for (long long __i = 0; __i < nameFilters_len; __i++) nameFilters_list.append(QString::fromUtf8(nameFilters[__i]));
	return new QDirIterator(path_qstr, nameFilters_list, QFlags<QDir::Filter>(filters), QFlags<QDirIterator::IteratorFlag>(flags));
}
char* QDirIterator_next(void* self, long long* __out_len) {
	auto&& __r = ((QDirIterator*)self)->next();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QDirIterator_nextFileInfo(void* self) {
	auto&& __r = ((QDirIterator*)self)->nextFileInfo();
	return new QFileInfo(__r);
}
int QDirIterator_hasNext(void* self) {
	auto&& __r = ((QDirIterator*)self)->hasNext();
	return __r;
}
char* QDirIterator_fileName(void* self, long long* __out_len) {
	auto&& __r = ((QDirIterator*)self)->fileName();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QDirIterator_filePath(void* self, long long* __out_len) {
	auto&& __r = ((QDirIterator*)self)->filePath();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QDirIterator_fileInfo(void* self) {
	auto&& __r = ((QDirIterator*)self)->fileInfo();
	return new QFileInfo(__r);
}
char* QDirIterator_path(void* self, long long* __out_len) {
	auto&& __r = ((QDirIterator*)self)->path();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QDirIterator_delete(void* self) { delete (QDirIterator*)self; }

}

#include <QtCore/qlatin1stringmatcher.h>
#include <QtCore/qlatin1stringview.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qstringview.h>
#include "gen_qlatin1stringmatcher.h"

extern "C" {

void* QLatin1StringMatcher_new() {
	return new QLatin1StringMatcher();
}
void* QLatin1StringMatcher_new2(void* pattern) {
	return new QLatin1StringMatcher(*(QLatin1String*)pattern);
}
void* QLatin1StringMatcher_new3(void* pattern, int cs) {
	return new QLatin1StringMatcher(*(QLatin1String*)pattern, (Qt::CaseSensitivity)(cs));
}
void QLatin1StringMatcher_setPattern(void* self, void* pattern) {
	((QLatin1StringMatcher*)self)->setPattern(*(QLatin1String*)pattern);
}
void* QLatin1StringMatcher_pattern(void* self) {
	auto&& __r = ((QLatin1StringMatcher*)self)->pattern();
	return new QLatin1String(__r);
}
void QLatin1StringMatcher_setCaseSensitivity(void* self, int cs) {
	((QLatin1StringMatcher*)self)->setCaseSensitivity((Qt::CaseSensitivity)(cs));
}
int QLatin1StringMatcher_caseSensitivity(void* self) {
	auto&& __r = ((QLatin1StringMatcher*)self)->caseSensitivity();
	return (int)(__r);
}
long long QLatin1StringMatcher_indexIn(void* self, void* haystack) {
	auto&& __r = ((QLatin1StringMatcher*)self)->indexIn(*(QLatin1String*)haystack);
	return __r;
}
long long QLatin1StringMatcher_indexIn2(void* self, void* haystack) {
	auto&& __r = ((QLatin1StringMatcher*)self)->indexIn(*(QStringView*)haystack);
	return __r;
}
long long QLatin1StringMatcher_indexIn3(void* self, void* haystack, long long from) {
	auto&& __r = ((QLatin1StringMatcher*)self)->indexIn(*(QLatin1String*)haystack, from);
	return __r;
}
long long QLatin1StringMatcher_indexIn4(void* self, void* haystack, long long from) {
	auto&& __r = ((QLatin1StringMatcher*)self)->indexIn(*(QStringView*)haystack, from);
	return __r;
}
void QLatin1StringMatcher_delete(void* self) { delete (QLatin1StringMatcher*)self; }

}

#include <QtCore/qoperatingsystemversion.h>
#include <QtCore/qstring.h>
#include <QtCore/qversionnumber.h>
#include "gen_qoperatingsystemversion.h"

extern "C" {

void* QOperatingSystemVersion_base(void* self) { return (void*)static_cast<QOperatingSystemVersionUnexported*>((QOperatingSystemVersion*)self); }
void QOperatingSystemVersion_chain(void* self, void** out) {
	QOperatingSystemVersion* __p = (QOperatingSystemVersion*)self;
	QOperatingSystemVersionUnexported* __b0 = static_cast<QOperatingSystemVersionUnexported*>(__p); out[0] = (void*)__b0;
	QOperatingSystemVersionBase* __b1 = static_cast<QOperatingSystemVersionBase*>(__b0); out[1] = (void*)__b1;
}
void* QOperatingSystemVersion_new(void* osversion) {
	return new QOperatingSystemVersion(*(QOperatingSystemVersionBase*)osversion);
}
void* QOperatingSystemVersion_new2(int osType, int vmajor) {
	return new QOperatingSystemVersion((QOperatingSystemVersion::OSType)(osType), vmajor);
}
void* QOperatingSystemVersion_new3(void* arg1) {
	return new QOperatingSystemVersion(*(QOperatingSystemVersion*)arg1);
}
void* QOperatingSystemVersion_new4(int osType, int vmajor, int vminor) {
	return new QOperatingSystemVersion((QOperatingSystemVersion::OSType)(osType), vmajor, vminor);
}
void* QOperatingSystemVersion_new5(int osType, int vmajor, int vminor, int vmicro) {
	return new QOperatingSystemVersion((QOperatingSystemVersion::OSType)(osType), vmajor, vminor, vmicro);
}
void* QOperatingSystemVersion_current() {
	auto&& __r = QOperatingSystemVersion::current();
	return new QOperatingSystemVersion(__r);
}
int QOperatingSystemVersion_currentType() {
	auto&& __r = QOperatingSystemVersion::currentType();
	return (int)(__r);
}
int QOperatingSystemVersion_type(void* self) {
	auto&& __r = ((QOperatingSystemVersion*)self)->type();
	return (int)(__r);
}
char* QOperatingSystemVersion_name(void* osversion, long long* __out_len) {
	auto&& __r = QOperatingSystemVersion::name(*(QOperatingSystemVersionBase*)osversion);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QOperatingSystemVersion_delete(void* self) { delete (QOperatingSystemVersion*)self; }

}

#include <QtCore/qreadwritelock.h>
#include <QtCore/qdeadlinetimer.h>
#include "gen_qbasicreadwritelock.h"

extern "C" {

void* QBasicReadWriteLock_new() {
	return new QBasicReadWriteLock();
}
void QBasicReadWriteLock_lockForRead(void* self) {
	((QBasicReadWriteLock*)self)->lockForRead();
}
int QBasicReadWriteLock_tryLockForRead(void* self) {
	auto&& __r = ((QBasicReadWriteLock*)self)->tryLockForRead();
	return __r;
}
int QBasicReadWriteLock_tryLockForRead2(void* self, void* timeout) {
	auto&& __r = ((QBasicReadWriteLock*)self)->tryLockForRead(*(QDeadlineTimer*)timeout);
	return __r;
}
void QBasicReadWriteLock_lockForWrite(void* self) {
	((QBasicReadWriteLock*)self)->lockForWrite();
}
int QBasicReadWriteLock_tryLockForWrite(void* self) {
	auto&& __r = ((QBasicReadWriteLock*)self)->tryLockForWrite();
	return __r;
}
int QBasicReadWriteLock_tryLockForWrite2(void* self, void* timeout) {
	auto&& __r = ((QBasicReadWriteLock*)self)->tryLockForWrite(*(QDeadlineTimer*)timeout);
	return __r;
}
void QBasicReadWriteLock_unlock(void* self) {
	((QBasicReadWriteLock*)self)->unlock();
}
void QBasicReadWriteLock_lock(void* self) {
	((QBasicReadWriteLock*)self)->lock();
}
void QBasicReadWriteLock_lock_shared(void* self) {
	((QBasicReadWriteLock*)self)->lock_shared();
}
int QBasicReadWriteLock_try_lock(void* self) {
	auto&& __r = ((QBasicReadWriteLock*)self)->try_lock();
	return __r;
}
int QBasicReadWriteLock_try_lock_shared(void* self) {
	auto&& __r = ((QBasicReadWriteLock*)self)->try_lock_shared();
	return __r;
}
void QBasicReadWriteLock_unlock_shared(void* self) {
	((QBasicReadWriteLock*)self)->unlock_shared();
}
void QBasicReadWriteLock_delete(void* self) { delete (QBasicReadWriteLock*)self; }

}

#include <QtCore/qstorageinfo.h>
#include <QtCore/qstring.h>
#include <QtCore/qdir.h>
#include <QtCore/qbytearray.h>
#include "gen_qstorageinfo.h"

extern "C" {

void* QStorageInfo_new() {
	return new QStorageInfo();
}
void* QStorageInfo_new2(const char* path, long long path_len) {
	QString path_qstr = QString::fromUtf8(path, path_len);
	return new QStorageInfo(path_qstr);
}
void* QStorageInfo_new3(void* dir) {
	return new QStorageInfo(*(QDir*)dir);
}
void* QStorageInfo_new4(void* other) {
	return new QStorageInfo(*(QStorageInfo*)other);
}
void* QStorageInfo_OperatorAssign(void* self, void* other) {
	auto&& __r = ((QStorageInfo*)self)->operator=(*(QStorageInfo*)other);
	return (void*)&(__r);
}
void QStorageInfo_swap(void* self, void* other) {
	((QStorageInfo*)self)->swap(*(QStorageInfo*)other);
}
void QStorageInfo_setPath(void* self, const char* path, long long path_len) {
	QString path_qstr = QString::fromUtf8(path, path_len);
	((QStorageInfo*)self)->setPath(path_qstr);
}
char* QStorageInfo_rootPath(void* self, long long* __out_len) {
	auto&& __r = ((QStorageInfo*)self)->rootPath();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QStorageInfo_device(void* self) {
	auto&& __r = ((QStorageInfo*)self)->device();
	return new QByteArray(__r);
}
void* QStorageInfo_subvolume(void* self) {
	auto&& __r = ((QStorageInfo*)self)->subvolume();
	return new QByteArray(__r);
}
void* QStorageInfo_fileSystemType(void* self) {
	auto&& __r = ((QStorageInfo*)self)->fileSystemType();
	return new QByteArray(__r);
}
char* QStorageInfo_name(void* self, long long* __out_len) {
	auto&& __r = ((QStorageInfo*)self)->name();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QStorageInfo_displayName(void* self, long long* __out_len) {
	auto&& __r = ((QStorageInfo*)self)->displayName();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
long long QStorageInfo_bytesTotal(void* self) {
	auto&& __r = ((QStorageInfo*)self)->bytesTotal();
	return __r;
}
long long QStorageInfo_bytesFree(void* self) {
	auto&& __r = ((QStorageInfo*)self)->bytesFree();
	return __r;
}
long long QStorageInfo_bytesAvailable(void* self) {
	auto&& __r = ((QStorageInfo*)self)->bytesAvailable();
	return __r;
}
int QStorageInfo_blockSize(void* self) {
	auto&& __r = ((QStorageInfo*)self)->blockSize();
	return __r;
}
int QStorageInfo_isRoot(void* self) {
	auto&& __r = ((QStorageInfo*)self)->isRoot();
	return __r;
}
int QStorageInfo_isReadOnly(void* self) {
	auto&& __r = ((QStorageInfo*)self)->isReadOnly();
	return __r;
}
int QStorageInfo_isReady(void* self) {
	auto&& __r = ((QStorageInfo*)self)->isReady();
	return __r;
}
int QStorageInfo_isValid(void* self) {
	auto&& __r = ((QStorageInfo*)self)->isValid();
	return __r;
}
void QStorageInfo_refresh(void* self) {
	((QStorageInfo*)self)->refresh();
}
void** QStorageInfo_mountedVolumes(long long* __out_len) {
	auto&& __r = QStorageInfo::mountedVolumes();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QStorageInfo(__e);
	return __arr;
}
void* QStorageInfo_root() {
	auto&& __r = QStorageInfo::root();
	return new QStorageInfo(__r);
}
int QStorageInfo_OperatorEqual(void* self, void* rhs) {
	auto&& __r = (*(QStorageInfo*)self) == (*(QStorageInfo*)rhs);
	return __r;
}
int QStorageInfo_OperatorNotEqual(void* self, void* rhs) {
	auto&& __r = (*(QStorageInfo*)self) != (*(QStorageInfo*)rhs);
	return __r;
}
void QStorageInfo_delete(void* self) { delete (QStorageInfo*)self; }

}

#include <QtGui/qpixelformat.h>
#include "gen_qpixelformat.h"

extern "C" {

void* QPixelFormat_new() {
	return new QPixelFormat();
}
void* QPixelFormat_new2(int colorModel, unsigned char firstSize, unsigned char secondSize, unsigned char thirdSize, unsigned char fourthSize, unsigned char fifthSize, unsigned char alphaSize, int alphaUsage, int alphaPosition, int premultiplied, int typeInterpretation) {
	return new QPixelFormat((QPixelFormat::ColorModel)(colorModel), firstSize, secondSize, thirdSize, fourthSize, fifthSize, alphaSize, (QPixelFormat::AlphaUsage)(alphaUsage), (QPixelFormat::AlphaPosition)(alphaPosition), (QPixelFormat::AlphaPremultiplied)(premultiplied), (QPixelFormat::TypeInterpretation)(typeInterpretation));
}
void* QPixelFormat_new3(void* arg1) {
	return new QPixelFormat(*(QPixelFormat*)arg1);
}
void* QPixelFormat_new4(int colorModel, unsigned char firstSize, unsigned char secondSize, unsigned char thirdSize, unsigned char fourthSize, unsigned char fifthSize, unsigned char alphaSize, int alphaUsage, int alphaPosition, int premultiplied, int typeInterpretation, int byteOrder) {
	return new QPixelFormat((QPixelFormat::ColorModel)(colorModel), firstSize, secondSize, thirdSize, fourthSize, fifthSize, alphaSize, (QPixelFormat::AlphaUsage)(alphaUsage), (QPixelFormat::AlphaPosition)(alphaPosition), (QPixelFormat::AlphaPremultiplied)(premultiplied), (QPixelFormat::TypeInterpretation)(typeInterpretation), (QPixelFormat::ByteOrder)(byteOrder));
}
void* QPixelFormat_new5(int colorModel, unsigned char firstSize, unsigned char secondSize, unsigned char thirdSize, unsigned char fourthSize, unsigned char fifthSize, unsigned char alphaSize, int alphaUsage, int alphaPosition, int premultiplied, int typeInterpretation, int byteOrder, unsigned char subEnum) {
	return new QPixelFormat((QPixelFormat::ColorModel)(colorModel), firstSize, secondSize, thirdSize, fourthSize, fifthSize, alphaSize, (QPixelFormat::AlphaUsage)(alphaUsage), (QPixelFormat::AlphaPosition)(alphaPosition), (QPixelFormat::AlphaPremultiplied)(premultiplied), (QPixelFormat::TypeInterpretation)(typeInterpretation), (QPixelFormat::ByteOrder)(byteOrder), subEnum);
}
int QPixelFormat_colorModel(void* self) {
	auto&& __r = ((QPixelFormat*)self)->colorModel();
	return (int)(__r);
}
unsigned char QPixelFormat_channelCount(void* self) {
	auto&& __r = ((QPixelFormat*)self)->channelCount();
	return __r;
}
unsigned char QPixelFormat_redSize(void* self) {
	auto&& __r = ((QPixelFormat*)self)->redSize();
	return __r;
}
unsigned char QPixelFormat_greenSize(void* self) {
	auto&& __r = ((QPixelFormat*)self)->greenSize();
	return __r;
}
unsigned char QPixelFormat_blueSize(void* self) {
	auto&& __r = ((QPixelFormat*)self)->blueSize();
	return __r;
}
unsigned char QPixelFormat_cyanSize(void* self) {
	auto&& __r = ((QPixelFormat*)self)->cyanSize();
	return __r;
}
unsigned char QPixelFormat_magentaSize(void* self) {
	auto&& __r = ((QPixelFormat*)self)->magentaSize();
	return __r;
}
unsigned char QPixelFormat_yellowSize(void* self) {
	auto&& __r = ((QPixelFormat*)self)->yellowSize();
	return __r;
}
unsigned char QPixelFormat_blackSize(void* self) {
	auto&& __r = ((QPixelFormat*)self)->blackSize();
	return __r;
}
unsigned char QPixelFormat_hueSize(void* self) {
	auto&& __r = ((QPixelFormat*)self)->hueSize();
	return __r;
}
unsigned char QPixelFormat_saturationSize(void* self) {
	auto&& __r = ((QPixelFormat*)self)->saturationSize();
	return __r;
}
unsigned char QPixelFormat_lightnessSize(void* self) {
	auto&& __r = ((QPixelFormat*)self)->lightnessSize();
	return __r;
}
unsigned char QPixelFormat_brightnessSize(void* self) {
	auto&& __r = ((QPixelFormat*)self)->brightnessSize();
	return __r;
}
unsigned char QPixelFormat_alphaSize(void* self) {
	auto&& __r = ((QPixelFormat*)self)->alphaSize();
	return __r;
}
unsigned char QPixelFormat_bitsPerPixel(void* self) {
	auto&& __r = ((QPixelFormat*)self)->bitsPerPixel();
	return __r;
}
int QPixelFormat_alphaUsage(void* self) {
	auto&& __r = ((QPixelFormat*)self)->alphaUsage();
	return (int)(__r);
}
int QPixelFormat_alphaPosition(void* self) {
	auto&& __r = ((QPixelFormat*)self)->alphaPosition();
	return (int)(__r);
}
int QPixelFormat_premultiplied(void* self) {
	auto&& __r = ((QPixelFormat*)self)->premultiplied();
	return (int)(__r);
}
int QPixelFormat_typeInterpretation(void* self) {
	auto&& __r = ((QPixelFormat*)self)->typeInterpretation();
	return (int)(__r);
}
int QPixelFormat_byteOrder(void* self) {
	auto&& __r = ((QPixelFormat*)self)->byteOrder();
	return (int)(__r);
}
int QPixelFormat_yuvLayout(void* self) {
	auto&& __r = ((QPixelFormat*)self)->yuvLayout();
	return (int)(__r);
}
unsigned char QPixelFormat_subEnum(void* self) {
	auto&& __r = ((QPixelFormat*)self)->subEnum();
	return __r;
}
int QPixelFormat_OperatorEqual(void* self, void* fmt2) {
	auto&& __r = (*(QPixelFormat*)self) == (*(QPixelFormat*)fmt2);
	return __r;
}
int QPixelFormat_OperatorNotEqual(void* self, void* fmt2) {
	auto&& __r = (*(QPixelFormat*)self) != (*(QPixelFormat*)fmt2);
	return __r;
}
void QPixelFormat_delete(void* self) { delete (QPixelFormat*)self; }

}

#include <QtGui/qevent.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qpoint.h>
#include <QtCore/qnamespace.h>
#include <QtGui/qpointingdevice.h>
#include <QtCore/qobject.h>
#include <QtGui/qeventpoint.h>
#include <QtGui/qinputdevice.h>
#include "gen_qmouseevent.h"

extern "C" {

void* QMouseEvent_base(void* self) { return (void*)static_cast<QSinglePointEvent*>((QMouseEvent*)self); }
void QMouseEvent_chain(void* self, void** out) {
	QMouseEvent* __p = (QMouseEvent*)self;
	QSinglePointEvent* __b0 = static_cast<QSinglePointEvent*>(__p); out[0] = (void*)__b0;
	QPointerEvent* __b1 = static_cast<QPointerEvent*>(__b0); out[1] = (void*)__b1;
	QInputEvent* __b2 = static_cast<QInputEvent*>(__b1); out[2] = (void*)__b2;
	QEvent* __b3 = static_cast<QEvent*>(__b2); out[3] = (void*)__b3;
}
void* QMouseEvent_new(int type_, void* localPos, int button, int buttons, int modifiers) {
	return new QMouseEvent((QEvent::Type)(type_), *(QPointF*)localPos, (Qt::MouseButton)(button), QFlags<Qt::MouseButton>(buttons), QFlags<Qt::KeyboardModifier>(modifiers));
}
void* QMouseEvent_new2(int type_, void* localPos, void* globalPos, int button, int buttons, int modifiers) {
	return new QMouseEvent((QEvent::Type)(type_), *(QPointF*)localPos, *(QPointF*)globalPos, (Qt::MouseButton)(button), QFlags<Qt::MouseButton>(buttons), QFlags<Qt::KeyboardModifier>(modifiers));
}
void* QMouseEvent_new3(int type_, void* localPos, void* scenePos, void* globalPos, int button, int buttons, int modifiers) {
	return new QMouseEvent((QEvent::Type)(type_), *(QPointF*)localPos, *(QPointF*)scenePos, *(QPointF*)globalPos, (Qt::MouseButton)(button), QFlags<Qt::MouseButton>(buttons), QFlags<Qt::KeyboardModifier>(modifiers));
}
void* QMouseEvent_new4(int type_, void* localPos, void* scenePos, void* globalPos, int button, int buttons, int modifiers, int source) {
	return new QMouseEvent((QEvent::Type)(type_), *(QPointF*)localPos, *(QPointF*)scenePos, *(QPointF*)globalPos, (Qt::MouseButton)(button), QFlags<Qt::MouseButton>(buttons), QFlags<Qt::KeyboardModifier>(modifiers), (Qt::MouseEventSource)(source));
}
void* QMouseEvent_new5(int type_, void* localPos, int button, int buttons, int modifiers, void* device) {
	return new QMouseEvent((QEvent::Type)(type_), *(QPointF*)localPos, (Qt::MouseButton)(button), QFlags<Qt::MouseButton>(buttons), QFlags<Qt::KeyboardModifier>(modifiers), (QPointingDevice*)device);
}
void* QMouseEvent_new6(int type_, void* localPos, void* globalPos, int button, int buttons, int modifiers, void* device) {
	return new QMouseEvent((QEvent::Type)(type_), *(QPointF*)localPos, *(QPointF*)globalPos, (Qt::MouseButton)(button), QFlags<Qt::MouseButton>(buttons), QFlags<Qt::KeyboardModifier>(modifiers), (QPointingDevice*)device);
}
void* QMouseEvent_new7(int type_, void* localPos, void* scenePos, void* globalPos, int button, int buttons, int modifiers, void* device) {
	return new QMouseEvent((QEvent::Type)(type_), *(QPointF*)localPos, *(QPointF*)scenePos, *(QPointF*)globalPos, (Qt::MouseButton)(button), QFlags<Qt::MouseButton>(buttons), QFlags<Qt::KeyboardModifier>(modifiers), (QPointingDevice*)device);
}
void* QMouseEvent_new8(int type_, void* localPos, void* scenePos, void* globalPos, int button, int buttons, int modifiers, int source, void* device) {
	return new QMouseEvent((QEvent::Type)(type_), *(QPointF*)localPos, *(QPointF*)scenePos, *(QPointF*)globalPos, (Qt::MouseButton)(button), QFlags<Qt::MouseButton>(buttons), QFlags<Qt::KeyboardModifier>(modifiers), (Qt::MouseEventSource)(source), (QPointingDevice*)device);
}
void* QMouseEvent_pos(void* self) {
	auto&& __r = ((QMouseEvent*)self)->pos();
	return new QPoint(__r);
}
void* QMouseEvent_globalPos(void* self) {
	auto&& __r = ((QMouseEvent*)self)->globalPos();
	return new QPoint(__r);
}
int QMouseEvent_x(void* self) {
	auto&& __r = ((QMouseEvent*)self)->x();
	return __r;
}
int QMouseEvent_y(void* self) {
	auto&& __r = ((QMouseEvent*)self)->y();
	return __r;
}
int QMouseEvent_globalX(void* self) {
	auto&& __r = ((QMouseEvent*)self)->globalX();
	return __r;
}
int QMouseEvent_globalY(void* self) {
	auto&& __r = ((QMouseEvent*)self)->globalY();
	return __r;
}
void* QMouseEvent_localPos(void* self) {
	auto&& __r = ((QMouseEvent*)self)->localPos();
	return new QPointF(__r);
}
void* QMouseEvent_windowPos(void* self) {
	auto&& __r = ((QMouseEvent*)self)->windowPos();
	return new QPointF(__r);
}
void* QMouseEvent_screenPos(void* self) {
	auto&& __r = ((QMouseEvent*)self)->screenPos();
	return new QPointF(__r);
}
int QMouseEvent_source(void* self) {
	auto&& __r = ((QMouseEvent*)self)->source();
	return (int)(__r);
}
int QMouseEvent_flags(void* self) {
	auto&& __r = ((QMouseEvent*)self)->flags();
	return (int)(__r);
}
int QMouseEvent_registerEventType() {
	auto&& __r = QMouseEvent::registerEventType();
	return __r;
}
int QMouseEvent_registerEventType2(int hint) {
	auto&& __r = QMouseEvent::registerEventType(hint);
	return __r;
}
void QMouseEvent_delete(void* self) { delete (QMouseEvent*)self; }

}

#include <QtGui/qevent.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qpoint.h>
#include "gen_qhelpevent.h"

extern "C" {

void* QHelpEvent_base(void* self) { return (void*)static_cast<QEvent*>((QHelpEvent*)self); }
void* QHelpEvent_new(int type_, void* pos, void* globalPos) {
	return new QHelpEvent((QEvent::Type)(type_), *(QPoint*)pos, *(QPoint*)globalPos);
}
int QHelpEvent_x(void* self) {
	auto&& __r = ((QHelpEvent*)self)->x();
	return __r;
}
int QHelpEvent_y(void* self) {
	auto&& __r = ((QHelpEvent*)self)->y();
	return __r;
}
int QHelpEvent_globalX(void* self) {
	auto&& __r = ((QHelpEvent*)self)->globalX();
	return __r;
}
int QHelpEvent_globalY(void* self) {
	auto&& __r = ((QHelpEvent*)self)->globalY();
	return __r;
}
void* QHelpEvent_pos(void* self) {
	auto&& __r = ((QHelpEvent*)self)->pos();
	return (void*)&(__r);
}
void* QHelpEvent_globalPos(void* self) {
	auto&& __r = ((QHelpEvent*)self)->globalPos();
	return (void*)&(__r);
}
int QHelpEvent_registerEventType() {
	auto&& __r = QHelpEvent::registerEventType();
	return __r;
}
int QHelpEvent_registerEventType2(int hint) {
	auto&& __r = QHelpEvent::registerEventType(hint);
	return __r;
}
void QHelpEvent_delete(void* self) { delete (QHelpEvent*)self; }

}

#include <QtGui/qbrush.h>
#include "gen_qgradient__qgradientdata.h"

extern "C" {

void* QGradient__QGradientData_new(void* arg1) {
	return new QGradient::QGradientData(*(QGradient::QGradientData*)arg1);
}
void QGradient__QGradientData_delete(void* self) { delete (QGradient::QGradientData*)self; }

}

#include <QtGui/qabstracttextdocumentlayout.h>
#include <QtGui/qpalette.h>
#include <QtCore/qrect.h>
#include "gen_qabstracttextdocumentlayout__paintcontext.h"

extern "C" {

void* QAbstractTextDocumentLayout__PaintContext_new() {
	return new QAbstractTextDocumentLayout::PaintContext();
}
void* QAbstractTextDocumentLayout__PaintContext_new2(void* arg1) {
	return new QAbstractTextDocumentLayout::PaintContext(*(QAbstractTextDocumentLayout::PaintContext*)arg1);
}
void* QAbstractTextDocumentLayout__PaintContext_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QAbstractTextDocumentLayout::PaintContext*)self)->operator=(*(QAbstractTextDocumentLayout::PaintContext*)arg1);
	return (void*)&(__r);
}
int QAbstractTextDocumentLayout__PaintContext_cursorPosition_get(void* self) {
	auto __r = ((QAbstractTextDocumentLayout::PaintContext*)self)->cursorPosition;
	return __r;
}
void QAbstractTextDocumentLayout__PaintContext_cursorPosition_set(void* self, int cursorPosition) {
	((QAbstractTextDocumentLayout::PaintContext*)self)->cursorPosition = cursorPosition;
}
void* QAbstractTextDocumentLayout__PaintContext_palette_get(void* self) {
	auto __r = ((QAbstractTextDocumentLayout::PaintContext*)self)->palette;
	return new QPalette(__r);
}
void QAbstractTextDocumentLayout__PaintContext_palette_set(void* self, void* palette) {
	((QAbstractTextDocumentLayout::PaintContext*)self)->palette = *(QPalette*)palette;
}
void* QAbstractTextDocumentLayout__PaintContext_clip_get(void* self) {
	auto __r = ((QAbstractTextDocumentLayout::PaintContext*)self)->clip;
	return new QRectF(__r);
}
void QAbstractTextDocumentLayout__PaintContext_clip_set(void* self, void* clip) {
	((QAbstractTextDocumentLayout::PaintContext*)self)->clip = *(QRectF*)clip;
}
void** QAbstractTextDocumentLayout__PaintContext_selections_get(void* self, long long* __out_len) {
	auto __r = ((QAbstractTextDocumentLayout::PaintContext*)self)->selections;
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QAbstractTextDocumentLayout::Selection(__e);
	return __arr;
}
void QAbstractTextDocumentLayout__PaintContext_selections_set(void* self, void** selections, long long selections_len) {
	QList<QAbstractTextDocumentLayout::Selection> selections_list; selections_list.reserve(selections_len);
	for (long long __i = 0; __i < selections_len; __i++) selections_list.append(*(QAbstractTextDocumentLayout::Selection*)selections[__i]);
	((QAbstractTextDocumentLayout::PaintContext*)self)->selections = selections_list;
}
void QAbstractTextDocumentLayout__PaintContext_delete(void* self) { delete (QAbstractTextDocumentLayout::PaintContext*)self; }

}

#include <QtGui/qaccessible.h>
#include <QtCore/qobject.h>
#include <QtCore/qvariant.h>
#include <QtGui/qaccessible_base.h>
#include "gen_qaccessiblevaluechangeevent.h"

extern "C" {

void* QAccessibleValueChangeEvent_base(void* self) { return (void*)static_cast<QAccessibleEvent*>((QAccessibleValueChangeEvent*)self); }
void* QAccessibleValueChangeEvent_new(void* obj, void* val) {
	return new QAccessibleValueChangeEvent((QObject*)obj, *(QVariant*)val);
}
void* QAccessibleValueChangeEvent_new2(void* iface, void* val) {
	return new QAccessibleValueChangeEvent((QAccessibleInterface*)iface, *(QVariant*)val);
}
void QAccessibleValueChangeEvent_setValue(void* self, void* val) {
	((QAccessibleValueChangeEvent*)self)->setValue(*(QVariant*)val);
}
void* QAccessibleValueChangeEvent_value(void* self) {
	auto&& __r = ((QAccessibleValueChangeEvent*)self)->value();
	return new QVariant(__r);
}
void QAccessibleValueChangeEvent_delete(void* self) { delete (QAccessibleValueChangeEvent*)self; }

}

#include <QtGui/qfontinfo.h>
#include <QtGui/qfont.h>
#include <QtCore/qstring.h>
#include <QtGui/qfontvariableaxis.h>
#include "gen_qfontinfo.h"

extern "C" {

void* QFontInfo_new(void* arg1) {
	return new QFontInfo(*(QFont*)arg1);
}
void* QFontInfo_new2(void* arg1) {
	return new QFontInfo(*(QFontInfo*)arg1);
}
void* QFontInfo_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QFontInfo*)self)->operator=(*(QFontInfo*)arg1);
	return (void*)&(__r);
}
void QFontInfo_swap(void* self, void* other) {
	((QFontInfo*)self)->swap(*(QFontInfo*)other);
}
char* QFontInfo_family(void* self, long long* __out_len) {
	auto&& __r = ((QFontInfo*)self)->family();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QFontInfo_styleName(void* self, long long* __out_len) {
	auto&& __r = ((QFontInfo*)self)->styleName();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QFontInfo_pixelSize(void* self) {
	auto&& __r = ((QFontInfo*)self)->pixelSize();
	return __r;
}
int QFontInfo_pointSize(void* self) {
	auto&& __r = ((QFontInfo*)self)->pointSize();
	return __r;
}
double QFontInfo_pointSizeF(void* self) {
	auto&& __r = ((QFontInfo*)self)->pointSizeF();
	return __r;
}
int QFontInfo_italic(void* self) {
	auto&& __r = ((QFontInfo*)self)->italic();
	return __r;
}
int QFontInfo_style(void* self) {
	auto&& __r = ((QFontInfo*)self)->style();
	return (int)(__r);
}
int QFontInfo_weight(void* self) {
	auto&& __r = ((QFontInfo*)self)->weight();
	return __r;
}
int QFontInfo_bold(void* self) {
	auto&& __r = ((QFontInfo*)self)->bold();
	return __r;
}
int QFontInfo_underline(void* self) {
	auto&& __r = ((QFontInfo*)self)->underline();
	return __r;
}
int QFontInfo_overline(void* self) {
	auto&& __r = ((QFontInfo*)self)->overline();
	return __r;
}
int QFontInfo_strikeOut(void* self) {
	auto&& __r = ((QFontInfo*)self)->strikeOut();
	return __r;
}
int QFontInfo_fixedPitch(void* self) {
	auto&& __r = ((QFontInfo*)self)->fixedPitch();
	return __r;
}
int QFontInfo_styleHint(void* self) {
	auto&& __r = ((QFontInfo*)self)->styleHint();
	return (int)(__r);
}
void** QFontInfo_variableAxes(void* self, long long* __out_len) {
	auto&& __r = ((QFontInfo*)self)->variableAxes();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QFontVariableAxis(__e);
	return __arr;
}
int QFontInfo_legacyWeight(void* self) {
	auto&& __r = ((QFontInfo*)self)->legacyWeight();
	return __r;
}
int QFontInfo_exactMatch(void* self) {
	auto&& __r = ((QFontInfo*)self)->exactMatch();
	return __r;
}
void QFontInfo_delete(void* self) { delete (QFontInfo*)self; }

}

#include <QtGui/qpaintdevicewindow.h>
#include <QtGui/qwindow.h>
#include <QtGui/qpaintdevice.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qrect.h>
#include <QtGui/qregion.h>
#include <QtGui/qevent.h>
#include <QtCore/qcoreevent.h>
#include <QtGui/qsurface.h>
#include <QtCore/qnamespace.h>
#include <QtGui/qsurfaceformat.h>
#include <QtCore/qsize.h>
#include <QtCore/qmargins.h>
#include <QtCore/qpoint.h>
#include <QtGui/qicon.h>
#include <QtGui/qscreen.h>
#include <QtGui/qaccessible.h>
#include <QtCore/qobject.h>
#include <QtGui/qcursor.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qvariant.h>
#include <QtCore/qbindingstorage.h>
#include <QtGui/qpainter.h>
#include "gen_qpaintdevicewindow.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

static std::unordered_map<long long, QMetaObject::Connection> __conns_QPaintDeviceWindow;
static std::mutex __connsMu_QPaintDeviceWindow;

extern "C" {

void* QPaintDeviceWindow_base(void* self) { return (void*)static_cast<QWindow*>((QPaintDeviceWindow*)self); }
void QPaintDeviceWindow_chain(void* self, void** out) {
	QPaintDeviceWindow* __p = (QPaintDeviceWindow*)self;
	QWindow* __b0 = static_cast<QWindow*>(__p); out[0] = (void*)__b0;
	QObject* __b1 = static_cast<QObject*>(__b0); out[1] = (void*)__b1;
}
void* QPaintDeviceWindow_fromQObject(void* o) { return (void*)qobject_cast<QPaintDeviceWindow*>((QObject*)o); }
char* QPaintDeviceWindow_tr(const char* s, long long* __out_len) {
	auto&& __r = QPaintDeviceWindow::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QPaintDeviceWindow_update(void* self, void* rect) {
	((QPaintDeviceWindow*)self)->update(*(QRect*)rect);
}
void QPaintDeviceWindow_update2(void* self, void* region) {
	((QPaintDeviceWindow*)self)->update(*(QRegion*)region);
}
void QPaintDeviceWindow_update3(void* self) {
	((QPaintDeviceWindow*)self)->update();
}
void* QPaintDeviceWindow_fromWinId(unsigned long long id) {
	auto&& __r = QPaintDeviceWindow::fromWinId(id);
	return (void*)(__r);
}
void* QPaintDeviceWindow_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QPaintDeviceWindow::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QPaintDeviceWindow_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QPaintDeviceWindow::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
int QPaintDeviceWindow_devType(void* self) {
	auto&& __r = ((QPaintDeviceWindow*)self)->devType();
	return __r;
}
int QPaintDeviceWindow_paintingActive(void* self) {
	auto&& __r = ((QPaintDeviceWindow*)self)->paintingActive();
	return __r;
}
int QPaintDeviceWindow_widthMM(void* self) {
	auto&& __r = ((QPaintDeviceWindow*)self)->widthMM();
	return __r;
}
int QPaintDeviceWindow_heightMM(void* self) {
	auto&& __r = ((QPaintDeviceWindow*)self)->heightMM();
	return __r;
}
int QPaintDeviceWindow_logicalDpiX(void* self) {
	auto&& __r = ((QPaintDeviceWindow*)self)->logicalDpiX();
	return __r;
}
int QPaintDeviceWindow_logicalDpiY(void* self) {
	auto&& __r = ((QPaintDeviceWindow*)self)->logicalDpiY();
	return __r;
}
int QPaintDeviceWindow_physicalDpiX(void* self) {
	auto&& __r = ((QPaintDeviceWindow*)self)->physicalDpiX();
	return __r;
}
int QPaintDeviceWindow_physicalDpiY(void* self) {
	auto&& __r = ((QPaintDeviceWindow*)self)->physicalDpiY();
	return __r;
}
double QPaintDeviceWindow_devicePixelRatioF(void* self) {
	auto&& __r = ((QPaintDeviceWindow*)self)->devicePixelRatioF();
	return __r;
}
int QPaintDeviceWindow_colorCount(void* self) {
	auto&& __r = ((QPaintDeviceWindow*)self)->colorCount();
	return __r;
}
int QPaintDeviceWindow_depth(void* self) {
	auto&& __r = ((QPaintDeviceWindow*)self)->depth();
	return __r;
}
double QPaintDeviceWindow_devicePixelRatioFScale() {
	auto&& __r = QPaintDeviceWindow::devicePixelRatioFScale();
	return __r;
}
int QPaintDeviceWindow_encodeMetricF(int metric, double value) {
	auto&& __r = QPaintDeviceWindow::encodeMetricF((QPaintDevice::PaintDeviceMetric)(metric), value);
	return __r;
}
char* QPaintDeviceWindow_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QPaintDeviceWindow::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QPaintDeviceWindow_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QPaintDeviceWindow::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QPaintDeviceWindow_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QPaintDeviceWindow::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QPaintDeviceWindow_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QPaintDeviceWindow::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QPaintDeviceWindow_asQPaintDevice(void* self) {
	auto&& __r = static_cast<QPaintDevice*>((QPaintDeviceWindow*)self);
	return (void*)(__r);
}
void QPaintDeviceWindow_connect_screenChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QPaintDeviceWindow*)self, &QPaintDeviceWindow::screenChanged, (QPaintDeviceWindow*)self, [__id](QScreen* screen){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)screen);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QPaintDeviceWindow);
	__conns_QPaintDeviceWindow[__id] = __c;
}
void QPaintDeviceWindow_connect_modalityChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QPaintDeviceWindow*)self, &QPaintDeviceWindow::modalityChanged, (QPaintDeviceWindow*)self, [__id](Qt::WindowModality modality){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, (int)(modality));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QPaintDeviceWindow);
	__conns_QPaintDeviceWindow[__id] = __c;
}
void QPaintDeviceWindow_connect_flagsChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QPaintDeviceWindow*)self, &QPaintDeviceWindow::flagsChanged, (QPaintDeviceWindow*)self, [__id](QFlags<Qt::WindowType> flags){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, (int)(flags));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QPaintDeviceWindow);
	__conns_QPaintDeviceWindow[__id] = __c;
}
void QPaintDeviceWindow_connect_windowStateChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QPaintDeviceWindow*)self, &QPaintDeviceWindow::windowStateChanged, (QPaintDeviceWindow*)self, [__id](Qt::WindowState windowState){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, (int)(windowState));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QPaintDeviceWindow);
	__conns_QPaintDeviceWindow[__id] = __c;
}
void QPaintDeviceWindow_connect_windowTitleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QPaintDeviceWindow*)self, &QPaintDeviceWindow::windowTitleChanged, (QPaintDeviceWindow*)self, [__id](const QString& title){
		QByteArray title_ba = title.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)title_ba.constData(), title_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QPaintDeviceWindow);
	__conns_QPaintDeviceWindow[__id] = __c;
}
void QPaintDeviceWindow_connect_xChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QPaintDeviceWindow*)self, &QPaintDeviceWindow::xChanged, (QPaintDeviceWindow*)self, [__id](int arg){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, arg);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QPaintDeviceWindow);
	__conns_QPaintDeviceWindow[__id] = __c;
}
void QPaintDeviceWindow_connect_yChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QPaintDeviceWindow*)self, &QPaintDeviceWindow::yChanged, (QPaintDeviceWindow*)self, [__id](int arg){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, arg);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QPaintDeviceWindow);
	__conns_QPaintDeviceWindow[__id] = __c;
}
void QPaintDeviceWindow_connect_widthChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QPaintDeviceWindow*)self, &QPaintDeviceWindow::widthChanged, (QPaintDeviceWindow*)self, [__id](int arg){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, arg);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QPaintDeviceWindow);
	__conns_QPaintDeviceWindow[__id] = __c;
}
void QPaintDeviceWindow_connect_heightChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QPaintDeviceWindow*)self, &QPaintDeviceWindow::heightChanged, (QPaintDeviceWindow*)self, [__id](int arg){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, arg);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QPaintDeviceWindow);
	__conns_QPaintDeviceWindow[__id] = __c;
}
void QPaintDeviceWindow_connect_minimumWidthChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QPaintDeviceWindow*)self, &QPaintDeviceWindow::minimumWidthChanged, (QPaintDeviceWindow*)self, [__id](int arg){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, arg);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QPaintDeviceWindow);
	__conns_QPaintDeviceWindow[__id] = __c;
}
void QPaintDeviceWindow_connect_minimumHeightChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QPaintDeviceWindow*)self, &QPaintDeviceWindow::minimumHeightChanged, (QPaintDeviceWindow*)self, [__id](int arg){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, arg);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QPaintDeviceWindow);
	__conns_QPaintDeviceWindow[__id] = __c;
}
void QPaintDeviceWindow_connect_maximumWidthChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QPaintDeviceWindow*)self, &QPaintDeviceWindow::maximumWidthChanged, (QPaintDeviceWindow*)self, [__id](int arg){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, arg);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QPaintDeviceWindow);
	__conns_QPaintDeviceWindow[__id] = __c;
}
void QPaintDeviceWindow_connect_maximumHeightChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QPaintDeviceWindow*)self, &QPaintDeviceWindow::maximumHeightChanged, (QPaintDeviceWindow*)self, [__id](int arg){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, arg);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QPaintDeviceWindow);
	__conns_QPaintDeviceWindow[__id] = __c;
}
void QPaintDeviceWindow_connect_safeAreaMarginsChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QPaintDeviceWindow*)self, &QPaintDeviceWindow::safeAreaMarginsChanged, (QPaintDeviceWindow*)self, [__id](QMargins arg){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&arg);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QPaintDeviceWindow);
	__conns_QPaintDeviceWindow[__id] = __c;
}
void QPaintDeviceWindow_connect_visibleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QPaintDeviceWindow*)self, &QPaintDeviceWindow::visibleChanged, (QPaintDeviceWindow*)self, [__id](bool arg){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, arg);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QPaintDeviceWindow);
	__conns_QPaintDeviceWindow[__id] = __c;
}
void QPaintDeviceWindow_connect_visibilityChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QPaintDeviceWindow*)self, &QPaintDeviceWindow::visibilityChanged, (QPaintDeviceWindow*)self, [__id](QWindow::Visibility visibility){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, (int)(visibility));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QPaintDeviceWindow);
	__conns_QPaintDeviceWindow[__id] = __c;
}
void QPaintDeviceWindow_connect_activeChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QPaintDeviceWindow*)self, &QPaintDeviceWindow::activeChanged, (QPaintDeviceWindow*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QPaintDeviceWindow);
	__conns_QPaintDeviceWindow[__id] = __c;
}
void QPaintDeviceWindow_connect_contentOrientationChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QPaintDeviceWindow*)self, &QPaintDeviceWindow::contentOrientationChanged, (QPaintDeviceWindow*)self, [__id](Qt::ScreenOrientation orientation){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, (int)(orientation));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QPaintDeviceWindow);
	__conns_QPaintDeviceWindow[__id] = __c;
}
void QPaintDeviceWindow_connect_focusObjectChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QPaintDeviceWindow*)self, &QPaintDeviceWindow::focusObjectChanged, (QPaintDeviceWindow*)self, [__id](QObject* object){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)object);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QPaintDeviceWindow);
	__conns_QPaintDeviceWindow[__id] = __c;
}
void QPaintDeviceWindow_connect_opacityChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QPaintDeviceWindow*)self, &QPaintDeviceWindow::opacityChanged, (QPaintDeviceWindow*)self, [__id](qreal opacity){
		if (cppgo_sig_void_unsignedlonglong) ((void(*)(long long, unsigned long long))cppgo_sig_void_unsignedlonglong)(__id, __builtin_bit_cast(unsigned long long, (double)(opacity)));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QPaintDeviceWindow);
	__conns_QPaintDeviceWindow[__id] = __c;
}
void QPaintDeviceWindow_connect_transientParentChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QPaintDeviceWindow*)self, &QPaintDeviceWindow::transientParentChanged, (QPaintDeviceWindow*)self, [__id](QWindow* transientParent){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)transientParent);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QPaintDeviceWindow);
	__conns_QPaintDeviceWindow[__id] = __c;
}
void QPaintDeviceWindow_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QPaintDeviceWindow*)self, &QPaintDeviceWindow::destroyed, (QPaintDeviceWindow*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QPaintDeviceWindow);
	__conns_QPaintDeviceWindow[__id] = __c;
}
void QPaintDeviceWindow_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QPaintDeviceWindow*)self, &QPaintDeviceWindow::destroyed, (QPaintDeviceWindow*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QPaintDeviceWindow);
	__conns_QPaintDeviceWindow[__id] = __c;
}
void QPaintDeviceWindow_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QPaintDeviceWindow*)self, &QPaintDeviceWindow::objectNameChanged, (QPaintDeviceWindow*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QPaintDeviceWindow);
	__conns_QPaintDeviceWindow[__id] = __c;
}
void QPaintDeviceWindow__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QPaintDeviceWindow);
		auto __it = __conns_QPaintDeviceWindow.find(id);
		if (__it == __conns_QPaintDeviceWindow.end()) return;
		__c = __it->second;
		__conns_QPaintDeviceWindow.erase(__it);
	}
	QObject::disconnect(__c);
}
void QPaintDeviceWindow_delete(void* self) { delete (QPaintDeviceWindow*)self; }

}

#include <QtGui/qstylehints.h>
#include <QtCore/qobject.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qchar.h>
#include <QtGui/qaccessibilityhints.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qvariant.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qstylehints.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

static std::unordered_map<long long, QMetaObject::Connection> __conns_QStyleHints;
static std::mutex __connsMu_QStyleHints;

extern "C" {

void* QStyleHints_base(void* self) { return (void*)static_cast<QObject*>((QStyleHints*)self); }
void* QStyleHints_fromQObject(void* o) { return (void*)qobject_cast<QStyleHints*>((QObject*)o); }
char* QStyleHints_tr(const char* s, long long* __out_len) {
	auto&& __r = QStyleHints::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QStyleHints_setMouseDoubleClickInterval(void* self, int mouseDoubleClickInterval) {
	((QStyleHints*)self)->setMouseDoubleClickInterval(mouseDoubleClickInterval);
}
int QStyleHints_mouseDoubleClickInterval(void* self) {
	auto&& __r = ((QStyleHints*)self)->mouseDoubleClickInterval();
	return __r;
}
int QStyleHints_mouseDoubleClickDistance(void* self) {
	auto&& __r = ((QStyleHints*)self)->mouseDoubleClickDistance();
	return __r;
}
int QStyleHints_touchDoubleTapDistance(void* self) {
	auto&& __r = ((QStyleHints*)self)->touchDoubleTapDistance();
	return __r;
}
void QStyleHints_setMousePressAndHoldInterval(void* self, int mousePressAndHoldInterval) {
	((QStyleHints*)self)->setMousePressAndHoldInterval(mousePressAndHoldInterval);
}
int QStyleHints_mousePressAndHoldInterval(void* self) {
	auto&& __r = ((QStyleHints*)self)->mousePressAndHoldInterval();
	return __r;
}
void QStyleHints_setStartDragDistance(void* self, int startDragDistance) {
	((QStyleHints*)self)->setStartDragDistance(startDragDistance);
}
int QStyleHints_startDragDistance(void* self) {
	auto&& __r = ((QStyleHints*)self)->startDragDistance();
	return __r;
}
void QStyleHints_setStartDragTime(void* self, int startDragTime) {
	((QStyleHints*)self)->setStartDragTime(startDragTime);
}
int QStyleHints_startDragTime(void* self) {
	auto&& __r = ((QStyleHints*)self)->startDragTime();
	return __r;
}
int QStyleHints_startDragVelocity(void* self) {
	auto&& __r = ((QStyleHints*)self)->startDragVelocity();
	return __r;
}
void QStyleHints_setKeyboardInputInterval(void* self, int keyboardInputInterval) {
	((QStyleHints*)self)->setKeyboardInputInterval(keyboardInputInterval);
}
int QStyleHints_keyboardInputInterval(void* self) {
	auto&& __r = ((QStyleHints*)self)->keyboardInputInterval();
	return __r;
}
int QStyleHints_keyboardAutoRepeatRate(void* self) {
	auto&& __r = ((QStyleHints*)self)->keyboardAutoRepeatRate();
	return __r;
}
double QStyleHints_keyboardAutoRepeatRateF(void* self) {
	auto&& __r = ((QStyleHints*)self)->keyboardAutoRepeatRateF();
	return __r;
}
void QStyleHints_setCursorFlashTime(void* self, int cursorFlashTime) {
	((QStyleHints*)self)->setCursorFlashTime(cursorFlashTime);
}
int QStyleHints_cursorFlashTime(void* self) {
	auto&& __r = ((QStyleHints*)self)->cursorFlashTime();
	return __r;
}
int QStyleHints_showIsFullScreen(void* self) {
	auto&& __r = ((QStyleHints*)self)->showIsFullScreen();
	return __r;
}
int QStyleHints_showIsMaximized(void* self) {
	auto&& __r = ((QStyleHints*)self)->showIsMaximized();
	return __r;
}
int QStyleHints_showShortcutsInContextMenus(void* self) {
	auto&& __r = ((QStyleHints*)self)->showShortcutsInContextMenus();
	return __r;
}
void QStyleHints_setShowShortcutsInContextMenus(void* self, int showShortcutsInContextMenus) {
	((QStyleHints*)self)->setShowShortcutsInContextMenus(showShortcutsInContextMenus);
}
int QStyleHints_contextMenuTrigger(void* self) {
	auto&& __r = ((QStyleHints*)self)->contextMenuTrigger();
	return (int)(__r);
}
void QStyleHints_setContextMenuTrigger(void* self, int contextMenuTrigger) {
	((QStyleHints*)self)->setContextMenuTrigger((Qt::ContextMenuTrigger)(contextMenuTrigger));
}
int QStyleHints_menuSelectionWraps(void* self) {
	auto&& __r = ((QStyleHints*)self)->menuSelectionWraps();
	return __r;
}
int QStyleHints_passwordMaskDelay(void* self) {
	auto&& __r = ((QStyleHints*)self)->passwordMaskDelay();
	return __r;
}
void* QStyleHints_passwordMaskCharacter(void* self) {
	auto&& __r = ((QStyleHints*)self)->passwordMaskCharacter();
	return new QChar(__r);
}
double QStyleHints_fontSmoothingGamma(void* self) {
	auto&& __r = ((QStyleHints*)self)->fontSmoothingGamma();
	return __r;
}
int QStyleHints_useRtlExtensions(void* self) {
	auto&& __r = ((QStyleHints*)self)->useRtlExtensions();
	return __r;
}
int QStyleHints_setFocusOnTouchRelease(void* self) {
	auto&& __r = ((QStyleHints*)self)->setFocusOnTouchRelease();
	return __r;
}
int QStyleHints_tabFocusBehavior(void* self) {
	auto&& __r = ((QStyleHints*)self)->tabFocusBehavior();
	return (int)(__r);
}
void QStyleHints_setTabFocusBehavior(void* self, int tabFocusBehavior) {
	((QStyleHints*)self)->setTabFocusBehavior((Qt::TabFocusBehavior)(tabFocusBehavior));
}
int QStyleHints_singleClickActivation(void* self) {
	auto&& __r = ((QStyleHints*)self)->singleClickActivation();
	return __r;
}
int QStyleHints_useHoverEffects(void* self) {
	auto&& __r = ((QStyleHints*)self)->useHoverEffects();
	return __r;
}
void QStyleHints_setUseHoverEffects(void* self, int useHoverEffects) {
	((QStyleHints*)self)->setUseHoverEffects(useHoverEffects);
}
int QStyleHints_wheelScrollLines(void* self) {
	auto&& __r = ((QStyleHints*)self)->wheelScrollLines();
	return __r;
}
void QStyleHints_setWheelScrollLines(void* self, int scrollLines) {
	((QStyleHints*)self)->setWheelScrollLines(scrollLines);
}
void QStyleHints_setMouseQuickSelectionThreshold(void* self, int threshold) {
	((QStyleHints*)self)->setMouseQuickSelectionThreshold(threshold);
}
int QStyleHints_mouseQuickSelectionThreshold(void* self) {
	auto&& __r = ((QStyleHints*)self)->mouseQuickSelectionThreshold();
	return __r;
}
int QStyleHints_colorScheme(void* self) {
	auto&& __r = ((QStyleHints*)self)->colorScheme();
	return (int)(__r);
}
void QStyleHints_setColorScheme(void* self, int scheme) {
	((QStyleHints*)self)->setColorScheme((Qt::ColorScheme)(scheme));
}
void QStyleHints_unsetColorScheme(void* self) {
	((QStyleHints*)self)->unsetColorScheme();
}
void* QStyleHints_accessibility(void* self) {
	auto&& __r = ((QStyleHints*)self)->accessibility();
	return (void*)(__r);
}
void QStyleHints_cursorFlashTimeChanged(void* self, int cursorFlashTime) {
	((QStyleHints*)self)->cursorFlashTimeChanged(cursorFlashTime);
}
void QStyleHints_keyboardInputIntervalChanged(void* self, int keyboardInputInterval) {
	((QStyleHints*)self)->keyboardInputIntervalChanged(keyboardInputInterval);
}
void QStyleHints_mouseDoubleClickIntervalChanged(void* self, int mouseDoubleClickInterval) {
	((QStyleHints*)self)->mouseDoubleClickIntervalChanged(mouseDoubleClickInterval);
}
void QStyleHints_mousePressAndHoldIntervalChanged(void* self, int mousePressAndHoldInterval) {
	((QStyleHints*)self)->mousePressAndHoldIntervalChanged(mousePressAndHoldInterval);
}
void QStyleHints_startDragDistanceChanged(void* self, int startDragDistance) {
	((QStyleHints*)self)->startDragDistanceChanged(startDragDistance);
}
void QStyleHints_startDragTimeChanged(void* self, int startDragTime) {
	((QStyleHints*)self)->startDragTimeChanged(startDragTime);
}
void QStyleHints_tabFocusBehaviorChanged(void* self, int tabFocusBehavior) {
	((QStyleHints*)self)->tabFocusBehaviorChanged((Qt::TabFocusBehavior)(tabFocusBehavior));
}
void QStyleHints_useHoverEffectsChanged(void* self, int useHoverEffects) {
	((QStyleHints*)self)->useHoverEffectsChanged(useHoverEffects);
}
void QStyleHints_showShortcutsInContextMenusChanged(void* self, int arg1) {
	((QStyleHints*)self)->showShortcutsInContextMenusChanged(arg1);
}
void QStyleHints_contextMenuTriggerChanged(void* self, int contextMenuTrigger) {
	((QStyleHints*)self)->contextMenuTriggerChanged((Qt::ContextMenuTrigger)(contextMenuTrigger));
}
void QStyleHints_wheelScrollLinesChanged(void* self, int scrollLines) {
	((QStyleHints*)self)->wheelScrollLinesChanged(scrollLines);
}
void QStyleHints_mouseQuickSelectionThresholdChanged(void* self, int threshold) {
	((QStyleHints*)self)->mouseQuickSelectionThresholdChanged(threshold);
}
void QStyleHints_colorSchemeChanged(void* self, int colorScheme) {
	((QStyleHints*)self)->colorSchemeChanged((Qt::ColorScheme)(colorScheme));
}
void* QStyleHints_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QStyleHints::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QStyleHints_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QStyleHints::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
char* QStyleHints_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QStyleHints::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QStyleHints_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QStyleHints::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QStyleHints_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QStyleHints::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QStyleHints_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QStyleHints::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void QStyleHints_connect_cursorFlashTimeChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStyleHints*)self, &QStyleHints::cursorFlashTimeChanged, (QStyleHints*)self, [__id](int cursorFlashTime){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, cursorFlashTime);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStyleHints);
	__conns_QStyleHints[__id] = __c;
}
void QStyleHints_connect_keyboardInputIntervalChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStyleHints*)self, &QStyleHints::keyboardInputIntervalChanged, (QStyleHints*)self, [__id](int keyboardInputInterval){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, keyboardInputInterval);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStyleHints);
	__conns_QStyleHints[__id] = __c;
}
void QStyleHints_connect_mouseDoubleClickIntervalChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStyleHints*)self, &QStyleHints::mouseDoubleClickIntervalChanged, (QStyleHints*)self, [__id](int mouseDoubleClickInterval){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, mouseDoubleClickInterval);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStyleHints);
	__conns_QStyleHints[__id] = __c;
}
void QStyleHints_connect_mousePressAndHoldIntervalChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStyleHints*)self, &QStyleHints::mousePressAndHoldIntervalChanged, (QStyleHints*)self, [__id](int mousePressAndHoldInterval){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, mousePressAndHoldInterval);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStyleHints);
	__conns_QStyleHints[__id] = __c;
}
void QStyleHints_connect_startDragDistanceChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStyleHints*)self, &QStyleHints::startDragDistanceChanged, (QStyleHints*)self, [__id](int startDragDistance){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, startDragDistance);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStyleHints);
	__conns_QStyleHints[__id] = __c;
}
void QStyleHints_connect_startDragTimeChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStyleHints*)self, &QStyleHints::startDragTimeChanged, (QStyleHints*)self, [__id](int startDragTime){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, startDragTime);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStyleHints);
	__conns_QStyleHints[__id] = __c;
}
void QStyleHints_connect_tabFocusBehaviorChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStyleHints*)self, &QStyleHints::tabFocusBehaviorChanged, (QStyleHints*)self, [__id](Qt::TabFocusBehavior tabFocusBehavior){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, (int)(tabFocusBehavior));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStyleHints);
	__conns_QStyleHints[__id] = __c;
}
void QStyleHints_connect_useHoverEffectsChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStyleHints*)self, &QStyleHints::useHoverEffectsChanged, (QStyleHints*)self, [__id](bool useHoverEffects){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, useHoverEffects);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStyleHints);
	__conns_QStyleHints[__id] = __c;
}
void QStyleHints_connect_showShortcutsInContextMenusChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStyleHints*)self, &QStyleHints::showShortcutsInContextMenusChanged, (QStyleHints*)self, [__id](bool arg1){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStyleHints);
	__conns_QStyleHints[__id] = __c;
}
void QStyleHints_connect_contextMenuTriggerChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStyleHints*)self, &QStyleHints::contextMenuTriggerChanged, (QStyleHints*)self, [__id](Qt::ContextMenuTrigger contextMenuTrigger){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, (int)(contextMenuTrigger));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStyleHints);
	__conns_QStyleHints[__id] = __c;
}
void QStyleHints_connect_wheelScrollLinesChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStyleHints*)self, &QStyleHints::wheelScrollLinesChanged, (QStyleHints*)self, [__id](int scrollLines){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, scrollLines);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStyleHints);
	__conns_QStyleHints[__id] = __c;
}
void QStyleHints_connect_mouseQuickSelectionThresholdChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStyleHints*)self, &QStyleHints::mouseQuickSelectionThresholdChanged, (QStyleHints*)self, [__id](int threshold){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, threshold);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStyleHints);
	__conns_QStyleHints[__id] = __c;
}
void QStyleHints_connect_colorSchemeChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStyleHints*)self, &QStyleHints::colorSchemeChanged, (QStyleHints*)self, [__id](Qt::ColorScheme colorScheme){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, (int)(colorScheme));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStyleHints);
	__conns_QStyleHints[__id] = __c;
}
void QStyleHints_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStyleHints*)self, &QStyleHints::destroyed, (QStyleHints*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStyleHints);
	__conns_QStyleHints[__id] = __c;
}
void QStyleHints_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStyleHints*)self, &QStyleHints::destroyed, (QStyleHints*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStyleHints);
	__conns_QStyleHints[__id] = __c;
}
void QStyleHints_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStyleHints*)self, &QStyleHints::objectNameChanged, (QStyleHints*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStyleHints);
	__conns_QStyleHints[__id] = __c;
}
void QStyleHints__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QStyleHints);
		auto __it = __conns_QStyleHints.find(id);
		if (__it == __conns_QStyleHints.end()) return;
		__c = __it->second;
		__conns_QStyleHints.erase(__it);
	}
	QObject::disconnect(__c);
}
void QStyleHints_delete(void* self) { delete (QStyleHints*)self; }

}

#include <QtWidgets/qwidget.h>
#include <QtCore/qobject.h>
#include <QtGui/qpaintdevice.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
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
#include <QtCore/qcoreevent.h>
#include <QtGui/qevent.h>
#include <QtCore/qvariant.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qwidget.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QWidget__Go : public QWidget {
public:
	long long __goId;
	unsigned char __goOn[48] = {};
	template<typename... A> QWidget__Go(long long id, A&&... a) : QWidget(std::forward<A>(a)...), __goId(id) {}
	using QWidget::updateMicroFocus;
	using QWidget::create;
	using QWidget::destroy;
	using QWidget::focusNextChild;
	using QWidget::focusPreviousChild;
	using QWidget::sender;
	using QWidget::senderSignalIndex;
	using QWidget::receivers;
	using QWidget::isSignalConnected;
	using QWidget::getDecodedMetricF;
	bool cppgoBase_event(QEvent* __a0) { return QWidget::event(__a0); }
	void cppgoBase_mousePressEvent(QMouseEvent* __a0) { QWidget::mousePressEvent(__a0); }
	void cppgoBase_mouseReleaseEvent(QMouseEvent* __a0) { QWidget::mouseReleaseEvent(__a0); }
	void cppgoBase_mouseDoubleClickEvent(QMouseEvent* __a0) { QWidget::mouseDoubleClickEvent(__a0); }
	void cppgoBase_mouseMoveEvent(QMouseEvent* __a0) { QWidget::mouseMoveEvent(__a0); }
	void cppgoBase_wheelEvent(QWheelEvent* __a0) { QWidget::wheelEvent(__a0); }
	void cppgoBase_keyPressEvent(QKeyEvent* __a0) { QWidget::keyPressEvent(__a0); }
	void cppgoBase_keyReleaseEvent(QKeyEvent* __a0) { QWidget::keyReleaseEvent(__a0); }
	void cppgoBase_focusInEvent(QFocusEvent* __a0) { QWidget::focusInEvent(__a0); }
	void cppgoBase_focusOutEvent(QFocusEvent* __a0) { QWidget::focusOutEvent(__a0); }
	void cppgoBase_enterEvent(QEnterEvent* __a0) { QWidget::enterEvent(__a0); }
	void cppgoBase_leaveEvent(QEvent* __a0) { QWidget::leaveEvent(__a0); }
	void cppgoBase_paintEvent(QPaintEvent* __a0) { QWidget::paintEvent(__a0); }
	void cppgoBase_moveEvent(QMoveEvent* __a0) { QWidget::moveEvent(__a0); }
	void cppgoBase_resizeEvent(QResizeEvent* __a0) { QWidget::resizeEvent(__a0); }
	void cppgoBase_closeEvent(QCloseEvent* __a0) { QWidget::closeEvent(__a0); }
	void cppgoBase_contextMenuEvent(QContextMenuEvent* __a0) { QWidget::contextMenuEvent(__a0); }
	void cppgoBase_tabletEvent(QTabletEvent* __a0) { QWidget::tabletEvent(__a0); }
	void cppgoBase_actionEvent(QActionEvent* __a0) { QWidget::actionEvent(__a0); }
	void cppgoBase_dragEnterEvent(QDragEnterEvent* __a0) { QWidget::dragEnterEvent(__a0); }
	void cppgoBase_dragMoveEvent(QDragMoveEvent* __a0) { QWidget::dragMoveEvent(__a0); }
	void cppgoBase_dragLeaveEvent(QDragLeaveEvent* __a0) { QWidget::dragLeaveEvent(__a0); }
	void cppgoBase_dropEvent(QDropEvent* __a0) { QWidget::dropEvent(__a0); }
	void cppgoBase_showEvent(QShowEvent* __a0) { QWidget::showEvent(__a0); }
	void cppgoBase_hideEvent(QHideEvent* __a0) { QWidget::hideEvent(__a0); }
	void cppgoBase_changeEvent(QEvent* __a0) { QWidget::changeEvent(__a0); }
	int cppgoBase_metric(int __a0) { return QWidget::metric((QPaintDevice::PaintDeviceMetric)__a0); }
	void cppgoBase_initPainter(QPainter* __a0) { QWidget::initPainter(__a0); }
	QPaintDevice* cppgoBase_redirected(QPoint* __a0) { return QWidget::redirected(__a0); }
	QPainter* cppgoBase_sharedPainter() { return QWidget::sharedPainter(); }
	void cppgoBase_inputMethodEvent(QInputMethodEvent* __a0) { QWidget::inputMethodEvent(__a0); }
	bool cppgoBase_focusNextPrevChild(bool __a0) { return QWidget::focusNextPrevChild(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QWidget::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QWidget::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QWidget::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QWidget::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QWidget::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QWidget::metaObject();
	}
	int devType() const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8659679893868524044LL, &__h); if (__h) return __r; }
		return QWidget::devType();
	}
	void setVisible(bool visible) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 2110182719375245575LL, visible, &__h); if (__h) return; }
		QWidget::setVisible(visible);
	}
	QSize sizeHint() const override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 729690261457684101LL, &__h); if (__h) return *(QSize*)__r; }
		return QWidget::sizeHint();
	}
	QSize minimumSizeHint() const override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 7049264599837532491LL, &__h); if (__h) return *(QSize*)__r; }
		return QWidget::minimumSizeHint();
	}
	int heightForWidth(int arg1) const override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 4311534060495737339LL, arg1, &__h); if (__h) return __r; }
		return QWidget::heightForWidth(arg1);
	}
	bool hasHeightForWidth() const override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8266762397995423915LL, &__h); if (__h) return __r; }
		return QWidget::hasHeightForWidth();
	}
	QPaintEngine* paintEngine() const override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 6398007404756346381LL, &__h); if (__h) return (QPaintEngine*)__r; }
		return QWidget::paintEngine();
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QWidget::event(event);
	}
	void mousePressEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -332090898078543319LL, (void*)event, &__h); if (__h) return; }
		QWidget::mousePressEvent(event);
	}
	void mouseReleaseEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4415920049297795213LL, (void*)event, &__h); if (__h) return; }
		QWidget::mouseReleaseEvent(event);
	}
	void mouseDoubleClickEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4029073905697023887LL, (void*)event, &__h); if (__h) return; }
		QWidget::mouseDoubleClickEvent(event);
	}
	void mouseMoveEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4255494004099806581LL, (void*)event, &__h); if (__h) return; }
		QWidget::mouseMoveEvent(event);
	}
	void wheelEvent(QWheelEvent* event) override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1919988340934511316LL, (void*)event, &__h); if (__h) return; }
		QWidget::wheelEvent(event);
	}
	void keyPressEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1023791977679793185LL, (void*)event, &__h); if (__h) return; }
		QWidget::keyPressEvent(event);
	}
	void keyReleaseEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 772598475950050299LL, (void*)event, &__h); if (__h) return; }
		QWidget::keyReleaseEvent(event);
	}
	void focusInEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 975043767033552552LL, (void*)event, &__h); if (__h) return; }
		QWidget::focusInEvent(event);
	}
	void focusOutEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -385596010576048993LL, (void*)event, &__h); if (__h) return; }
		QWidget::focusOutEvent(event);
	}
	void enterEvent(QEnterEvent* event) override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6316270346668237895LL, (void*)event, &__h); if (__h) return; }
		QWidget::enterEvent(event);
	}
	void leaveEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -551070009623711640LL, (void*)event, &__h); if (__h) return; }
		QWidget::leaveEvent(event);
	}
	void paintEvent(QPaintEvent* event) override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5290442163167458465LL, (void*)event, &__h); if (__h) return; }
		QWidget::paintEvent(event);
	}
	void moveEvent(QMoveEvent* event) override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -9161949150489525396LL, (void*)event, &__h); if (__h) return; }
		QWidget::moveEvent(event);
	}
	void resizeEvent(QResizeEvent* event) override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4302917285655014661LL, (void*)event, &__h); if (__h) return; }
		QWidget::resizeEvent(event);
	}
	void closeEvent(QCloseEvent* event) override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1989937050019464783LL, (void*)event, &__h); if (__h) return; }
		QWidget::closeEvent(event);
	}
	void contextMenuEvent(QContextMenuEvent* event) override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6698043222227006015LL, (void*)event, &__h); if (__h) return; }
		QWidget::contextMenuEvent(event);
	}
	void tabletEvent(QTabletEvent* event) override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -212068866860701141LL, (void*)event, &__h); if (__h) return; }
		QWidget::tabletEvent(event);
	}
	void actionEvent(QActionEvent* event) override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1697444140088555581LL, (void*)event, &__h); if (__h) return; }
		QWidget::actionEvent(event);
	}
	void dragEnterEvent(QDragEnterEvent* event) override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 206122715919917589LL, (void*)event, &__h); if (__h) return; }
		QWidget::dragEnterEvent(event);
	}
	void dragMoveEvent(QDragMoveEvent* event) override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1781545137352745978LL, (void*)event, &__h); if (__h) return; }
		QWidget::dragMoveEvent(event);
	}
	void dragLeaveEvent(QDragLeaveEvent* event) override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -3139904158373725006LL, (void*)event, &__h); if (__h) return; }
		QWidget::dragLeaveEvent(event);
	}
	void dropEvent(QDropEvent* event) override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 9107058682754591716LL, (void*)event, &__h); if (__h) return; }
		QWidget::dropEvent(event);
	}
	void showEvent(QShowEvent* event) override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6825369174751916LL, (void*)event, &__h); if (__h) return; }
		QWidget::showEvent(event);
	}
	void hideEvent(QHideEvent* event) override {
		int __h = 0;
		if (__goOn[32] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -199718707876719129LL, (void*)event, &__h); if (__h) return; }
		QWidget::hideEvent(event);
	}
	bool nativeEvent(const QByteArray& eventType, void* message, qintptr* result) override {
		int __h = 0;
		if (__goOn[33] && cppgo_virt_int_voidp_voidp_longlongp) { int __r = ((int(*)(long long, long long, void*, void*, long long*, int*))cppgo_virt_int_voidp_voidp_longlongp)(__goId, -2039166586246443706LL, (void*)&eventType, (void*)message, (long long*)result, &__h); if (__h) return __r; }
		return QWidget::nativeEvent(eventType, message, result);
	}
	void changeEvent(QEvent* arg1) override {
		int __h = 0;
		if (__goOn[34] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4959376063826165943LL, (void*)arg1, &__h); if (__h) return; }
		QWidget::changeEvent(arg1);
	}
	int metric(QPaintDevice::PaintDeviceMetric arg1) const override {
		int __h = 0;
		if (__goOn[35] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 427629649267311777LL, (int)(arg1), &__h); if (__h) return __r; }
		return QWidget::metric(arg1);
	}
	void initPainter(QPainter* painter) const override {
		int __h = 0;
		if (__goOn[36] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6477500296980050210LL, (void*)painter, &__h); if (__h) return; }
		QWidget::initPainter(painter);
	}
	QPaintDevice* redirected(QPoint* offset) const override {
		int __h = 0;
		if (__goOn[37] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 4127240633310655048LL, (void*)offset, &__h); if (__h) return (QPaintDevice*)__r; }
		return QWidget::redirected(offset);
	}
	QPainter* sharedPainter() const override {
		int __h = 0;
		if (__goOn[38] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 5792771328741774991LL, &__h); if (__h) return (QPainter*)__r; }
		return QWidget::sharedPainter();
	}
	void inputMethodEvent(QInputMethodEvent* arg1) override {
		int __h = 0;
		if (__goOn[39] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5438819661942843222LL, (void*)arg1, &__h); if (__h) return; }
		QWidget::inputMethodEvent(arg1);
	}
	QVariant inputMethodQuery(Qt::InputMethodQuery arg1) const override {
		int __h = 0;
		if (__goOn[40] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7345853097340709638LL, (int)(arg1), &__h); if (__h) return *(QVariant*)__r; }
		return QWidget::inputMethodQuery(arg1);
	}
	bool focusNextPrevChild(bool next) override {
		int __h = 0;
		if (__goOn[41] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 5898328517505784191LL, next, &__h); if (__h) return __r; }
		return QWidget::focusNextPrevChild(next);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[42] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QWidget::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[43] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QWidget::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[44] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QWidget::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[45] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QWidget::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[46] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QWidget::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[47] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QWidget::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QWidget;
static std::mutex __connsMu_QWidget;

extern "C" {

void* QWidget_base(void* self) { return (void*)static_cast<QObject*>((QWidget*)self); }
void QWidget__set_override(void* self, int idx) { ((QWidget__Go*)self)->__goOn[idx] = 1; }
void* QWidget_fromQObject(void* o) { return (void*)qobject_cast<QWidget*>((QObject*)o); }
void* QWidget_new(long long __goid, void* parent) {
	return new QWidget__Go(__goid, (QWidget*)parent);
}
void* QWidget_new2(long long __goid) {
	return new QWidget__Go(__goid);
}
void* QWidget_new3(long long __goid, void* parent, int f) {
	return new QWidget__Go(__goid, (QWidget*)parent, QFlags<Qt::WindowType>(f));
}
char* QWidget_tr(const char* s, long long* __out_len) {
	auto&& __r = QWidget::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QWidget_devType(void* self) {
	auto&& __r = ((QWidget*)self)->devType();
	return __r;
}
unsigned long long QWidget_winId(void* self) {
	auto&& __r = ((QWidget*)self)->winId();
	return __r;
}
void QWidget_createWinId(void* self) {
	((QWidget*)self)->createWinId();
}
unsigned long long QWidget_internalWinId(void* self) {
	auto&& __r = ((QWidget*)self)->internalWinId();
	return __r;
}
unsigned long long QWidget_effectiveWinId(void* self) {
	auto&& __r = ((QWidget*)self)->effectiveWinId();
	return __r;
}
void* QWidget_style(void* self) {
	auto&& __r = ((QWidget*)self)->style();
	return (void*)(__r);
}
void QWidget_setStyle(void* self, void* arg1) {
	((QWidget*)self)->setStyle((QStyle*)arg1);
}
int QWidget_isTopLevel(void* self) {
	auto&& __r = ((QWidget*)self)->isTopLevel();
	return __r;
}
int QWidget_isWindow(void* self) {
	auto&& __r = ((QWidget*)self)->isWindow();
	return __r;
}
int QWidget_isModal(void* self) {
	auto&& __r = ((QWidget*)self)->isModal();
	return __r;
}
int QWidget_windowModality(void* self) {
	auto&& __r = ((QWidget*)self)->windowModality();
	return (int)(__r);
}
void QWidget_setWindowModality(void* self, int windowModality) {
	((QWidget*)self)->setWindowModality((Qt::WindowModality)(windowModality));
}
int QWidget_isEnabled(void* self) {
	auto&& __r = ((QWidget*)self)->isEnabled();
	return __r;
}
int QWidget_isEnabledTo(void* self, void* arg1) {
	auto&& __r = ((QWidget*)self)->isEnabledTo((QWidget*)arg1);
	return __r;
}
void QWidget_setEnabled(void* self, int arg1) {
	((QWidget*)self)->setEnabled(arg1);
}
void QWidget_setDisabled(void* self, int arg1) {
	((QWidget*)self)->setDisabled(arg1);
}
void QWidget_setWindowModified(void* self, int arg1) {
	((QWidget*)self)->setWindowModified(arg1);
}
void* QWidget_frameGeometry(void* self) {
	auto&& __r = ((QWidget*)self)->frameGeometry();
	return new QRect(__r);
}
void* QWidget_geometry(void* self) {
	auto&& __r = ((QWidget*)self)->geometry();
	return (void*)&(__r);
}
void* QWidget_normalGeometry(void* self) {
	auto&& __r = ((QWidget*)self)->normalGeometry();
	return new QRect(__r);
}
int QWidget_x(void* self) {
	auto&& __r = ((QWidget*)self)->x();
	return __r;
}
int QWidget_y(void* self) {
	auto&& __r = ((QWidget*)self)->y();
	return __r;
}
void* QWidget_pos(void* self) {
	auto&& __r = ((QWidget*)self)->pos();
	return new QPoint(__r);
}
void* QWidget_frameSize(void* self) {
	auto&& __r = ((QWidget*)self)->frameSize();
	return new QSize(__r);
}
void* QWidget_size(void* self) {
	auto&& __r = ((QWidget*)self)->size();
	return new QSize(__r);
}
int QWidget_width(void* self) {
	auto&& __r = ((QWidget*)self)->width();
	return __r;
}
int QWidget_height(void* self) {
	auto&& __r = ((QWidget*)self)->height();
	return __r;
}
void* QWidget_rect(void* self) {
	auto&& __r = ((QWidget*)self)->rect();
	return new QRect(__r);
}
void* QWidget_childrenRect(void* self) {
	auto&& __r = ((QWidget*)self)->childrenRect();
	return new QRect(__r);
}
void* QWidget_childrenRegion(void* self) {
	auto&& __r = ((QWidget*)self)->childrenRegion();
	return new QRegion(__r);
}
void* QWidget_minimumSize(void* self) {
	auto&& __r = ((QWidget*)self)->minimumSize();
	return new QSize(__r);
}
void* QWidget_maximumSize(void* self) {
	auto&& __r = ((QWidget*)self)->maximumSize();
	return new QSize(__r);
}
int QWidget_minimumWidth(void* self) {
	auto&& __r = ((QWidget*)self)->minimumWidth();
	return __r;
}
int QWidget_minimumHeight(void* self) {
	auto&& __r = ((QWidget*)self)->minimumHeight();
	return __r;
}
int QWidget_maximumWidth(void* self) {
	auto&& __r = ((QWidget*)self)->maximumWidth();
	return __r;
}
int QWidget_maximumHeight(void* self) {
	auto&& __r = ((QWidget*)self)->maximumHeight();
	return __r;
}
void QWidget_setMinimumSize(void* self, void* arg1) {
	((QWidget*)self)->setMinimumSize(*(QSize*)arg1);
}
void QWidget_setMinimumSize2(void* self, int minw, int minh) {
	((QWidget*)self)->setMinimumSize(minw, minh);
}
void QWidget_setMaximumSize(void* self, void* arg1) {
	((QWidget*)self)->setMaximumSize(*(QSize*)arg1);
}
void QWidget_setMaximumSize2(void* self, int maxw, int maxh) {
	((QWidget*)self)->setMaximumSize(maxw, maxh);
}
void QWidget_setMinimumWidth(void* self, int minw) {
	((QWidget*)self)->setMinimumWidth(minw);
}
void QWidget_setMinimumHeight(void* self, int minh) {
	((QWidget*)self)->setMinimumHeight(minh);
}
void QWidget_setMaximumWidth(void* self, int maxw) {
	((QWidget*)self)->setMaximumWidth(maxw);
}
void QWidget_setMaximumHeight(void* self, int maxh) {
	((QWidget*)self)->setMaximumHeight(maxh);
}
void* QWidget_sizeIncrement(void* self) {
	auto&& __r = ((QWidget*)self)->sizeIncrement();
	return new QSize(__r);
}
void QWidget_setSizeIncrement(void* self, void* arg1) {
	((QWidget*)self)->setSizeIncrement(*(QSize*)arg1);
}
void QWidget_setSizeIncrement2(void* self, int w, int h) {
	((QWidget*)self)->setSizeIncrement(w, h);
}
void* QWidget_baseSize(void* self) {
	auto&& __r = ((QWidget*)self)->baseSize();
	return new QSize(__r);
}
void QWidget_setBaseSize(void* self, void* arg1) {
	((QWidget*)self)->setBaseSize(*(QSize*)arg1);
}
void QWidget_setBaseSize2(void* self, int basew, int baseh) {
	((QWidget*)self)->setBaseSize(basew, baseh);
}
void QWidget_setFixedSize(void* self, void* arg1) {
	((QWidget*)self)->setFixedSize(*(QSize*)arg1);
}
void QWidget_setFixedSize2(void* self, int w, int h) {
	((QWidget*)self)->setFixedSize(w, h);
}
void QWidget_setFixedWidth(void* self, int w) {
	((QWidget*)self)->setFixedWidth(w);
}
void QWidget_setFixedHeight(void* self, int h) {
	((QWidget*)self)->setFixedHeight(h);
}
void* QWidget_mapToGlobal(void* self, void* arg1) {
	auto&& __r = ((QWidget*)self)->mapToGlobal(*(QPointF*)arg1);
	return new QPointF(__r);
}
void* QWidget_mapToGlobal2(void* self, void* arg1) {
	auto&& __r = ((QWidget*)self)->mapToGlobal(*(QPoint*)arg1);
	return new QPoint(__r);
}
void* QWidget_mapFromGlobal(void* self, void* arg1) {
	auto&& __r = ((QWidget*)self)->mapFromGlobal(*(QPointF*)arg1);
	return new QPointF(__r);
}
void* QWidget_mapFromGlobal2(void* self, void* arg1) {
	auto&& __r = ((QWidget*)self)->mapFromGlobal(*(QPoint*)arg1);
	return new QPoint(__r);
}
void* QWidget_mapToParent(void* self, void* arg1) {
	auto&& __r = ((QWidget*)self)->mapToParent(*(QPointF*)arg1);
	return new QPointF(__r);
}
void* QWidget_mapToParent2(void* self, void* arg1) {
	auto&& __r = ((QWidget*)self)->mapToParent(*(QPoint*)arg1);
	return new QPoint(__r);
}
void* QWidget_mapFromParent(void* self, void* arg1) {
	auto&& __r = ((QWidget*)self)->mapFromParent(*(QPointF*)arg1);
	return new QPointF(__r);
}
void* QWidget_mapFromParent2(void* self, void* arg1) {
	auto&& __r = ((QWidget*)self)->mapFromParent(*(QPoint*)arg1);
	return new QPoint(__r);
}
void* QWidget_mapTo(void* self, void* arg1, void* arg2) {
	auto&& __r = ((QWidget*)self)->mapTo((QWidget*)arg1, *(QPointF*)arg2);
	return new QPointF(__r);
}
void* QWidget_mapTo2(void* self, void* arg1, void* arg2) {
	auto&& __r = ((QWidget*)self)->mapTo((QWidget*)arg1, *(QPoint*)arg2);
	return new QPoint(__r);
}
void* QWidget_mapFrom(void* self, void* arg1, void* arg2) {
	auto&& __r = ((QWidget*)self)->mapFrom((QWidget*)arg1, *(QPointF*)arg2);
	return new QPointF(__r);
}
void* QWidget_mapFrom2(void* self, void* arg1, void* arg2) {
	auto&& __r = ((QWidget*)self)->mapFrom((QWidget*)arg1, *(QPoint*)arg2);
	return new QPoint(__r);
}
void* QWidget_window(void* self) {
	auto&& __r = ((QWidget*)self)->window();
	return (void*)(__r);
}
void* QWidget_nativeParentWidget(void* self) {
	auto&& __r = ((QWidget*)self)->nativeParentWidget();
	return (void*)(__r);
}
void* QWidget_topLevelWidget(void* self) {
	auto&& __r = ((QWidget*)self)->topLevelWidget();
	return (void*)(__r);
}
void* QWidget_palette(void* self) {
	auto&& __r = ((QWidget*)self)->palette();
	return (void*)&(__r);
}
void QWidget_setPalette(void* self, void* arg1) {
	((QWidget*)self)->setPalette(*(QPalette*)arg1);
}
void QWidget_setBackgroundRole(void* self, int arg1) {
	((QWidget*)self)->setBackgroundRole((QPalette::ColorRole)(arg1));
}
int QWidget_backgroundRole(void* self) {
	auto&& __r = ((QWidget*)self)->backgroundRole();
	return (int)(__r);
}
void QWidget_setForegroundRole(void* self, int arg1) {
	((QWidget*)self)->setForegroundRole((QPalette::ColorRole)(arg1));
}
int QWidget_foregroundRole(void* self) {
	auto&& __r = ((QWidget*)self)->foregroundRole();
	return (int)(__r);
}
void* QWidget_font(void* self) {
	auto&& __r = ((QWidget*)self)->font();
	return (void*)&(__r);
}
void QWidget_setFont(void* self, void* arg1) {
	((QWidget*)self)->setFont(*(QFont*)arg1);
}
void* QWidget_fontMetrics(void* self) {
	auto&& __r = ((QWidget*)self)->fontMetrics();
	return new QFontMetrics(__r);
}
void* QWidget_fontInfo(void* self) {
	auto&& __r = ((QWidget*)self)->fontInfo();
	return new QFontInfo(__r);
}
void* QWidget_cursor(void* self) {
	auto&& __r = ((QWidget*)self)->cursor();
	return new QCursor(__r);
}
void QWidget_setCursor(void* self, void* arg1) {
	((QWidget*)self)->setCursor(*(QCursor*)arg1);
}
void QWidget_unsetCursor(void* self) {
	((QWidget*)self)->unsetCursor();
}
void QWidget_setMouseTracking(void* self, int enable) {
	((QWidget*)self)->setMouseTracking(enable);
}
int QWidget_hasMouseTracking(void* self) {
	auto&& __r = ((QWidget*)self)->hasMouseTracking();
	return __r;
}
int QWidget_underMouse(void* self) {
	auto&& __r = ((QWidget*)self)->underMouse();
	return __r;
}
void QWidget_setTabletTracking(void* self, int enable) {
	((QWidget*)self)->setTabletTracking(enable);
}
int QWidget_hasTabletTracking(void* self) {
	auto&& __r = ((QWidget*)self)->hasTabletTracking();
	return __r;
}
void QWidget_setMask(void* self, void* arg1) {
	((QWidget*)self)->setMask(*(QBitmap*)arg1);
}
void QWidget_setMask2(void* self, void* arg1) {
	((QWidget*)self)->setMask(*(QRegion*)arg1);
}
void* QWidget_mask(void* self) {
	auto&& __r = ((QWidget*)self)->mask();
	return new QRegion(__r);
}
void QWidget_clearMask(void* self) {
	((QWidget*)self)->clearMask();
}
void QWidget_render(void* self, void* target) {
	((QWidget*)self)->render((QPaintDevice*)target);
}
void QWidget_render2(void* self, void* painter) {
	((QWidget*)self)->render((QPainter*)painter);
}
void* QWidget_grab(void* self) {
	auto&& __r = ((QWidget*)self)->grab();
	return new QPixmap(__r);
}
void* QWidget_graphicsEffect(void* self) {
	auto&& __r = ((QWidget*)self)->graphicsEffect();
	return (void*)(__r);
}
void QWidget_setGraphicsEffect(void* self, void* effect) {
	((QWidget*)self)->setGraphicsEffect((QGraphicsEffect*)effect);
}
void QWidget_grabGesture(void* self, int type_) {
	((QWidget*)self)->grabGesture((Qt::GestureType)(type_));
}
void QWidget_ungrabGesture(void* self, int type_) {
	((QWidget*)self)->ungrabGesture((Qt::GestureType)(type_));
}
void QWidget_setWindowTitle(void* self, const char* arg1, long long arg1_len) {
	QString arg1_qstr = QString::fromUtf8(arg1, arg1_len);
	((QWidget*)self)->setWindowTitle(arg1_qstr);
}
void QWidget_setStyleSheet(void* self, const char* styleSheet, long long styleSheet_len) {
	QString styleSheet_qstr = QString::fromUtf8(styleSheet, styleSheet_len);
	((QWidget*)self)->setStyleSheet(styleSheet_qstr);
}
char* QWidget_styleSheet(void* self, long long* __out_len) {
	auto&& __r = ((QWidget*)self)->styleSheet();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QWidget_windowTitle(void* self, long long* __out_len) {
	auto&& __r = ((QWidget*)self)->windowTitle();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QWidget_setWindowIcon(void* self, void* icon) {
	((QWidget*)self)->setWindowIcon(*(QIcon*)icon);
}
void* QWidget_windowIcon(void* self) {
	auto&& __r = ((QWidget*)self)->windowIcon();
	return new QIcon(__r);
}
void QWidget_setWindowIconText(void* self, const char* arg1, long long arg1_len) {
	QString arg1_qstr = QString::fromUtf8(arg1, arg1_len);
	((QWidget*)self)->setWindowIconText(arg1_qstr);
}
char* QWidget_windowIconText(void* self, long long* __out_len) {
	auto&& __r = ((QWidget*)self)->windowIconText();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QWidget_setWindowRole(void* self, const char* arg1, long long arg1_len) {
	QString arg1_qstr = QString::fromUtf8(arg1, arg1_len);
	((QWidget*)self)->setWindowRole(arg1_qstr);
}
char* QWidget_windowRole(void* self, long long* __out_len) {
	auto&& __r = ((QWidget*)self)->windowRole();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QWidget_setWindowFilePath(void* self, const char* filePath, long long filePath_len) {
	QString filePath_qstr = QString::fromUtf8(filePath, filePath_len);
	((QWidget*)self)->setWindowFilePath(filePath_qstr);
}
char* QWidget_windowFilePath(void* self, long long* __out_len) {
	auto&& __r = ((QWidget*)self)->windowFilePath();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QWidget_setWindowOpacity(void* self, double level) {
	((QWidget*)self)->setWindowOpacity(level);
}
double QWidget_windowOpacity(void* self) {
	auto&& __r = ((QWidget*)self)->windowOpacity();
	return __r;
}
int QWidget_isWindowModified(void* self) {
	auto&& __r = ((QWidget*)self)->isWindowModified();
	return __r;
}
void QWidget_setToolTip(void* self, const char* arg1, long long arg1_len) {
	QString arg1_qstr = QString::fromUtf8(arg1, arg1_len);
	((QWidget*)self)->setToolTip(arg1_qstr);
}
char* QWidget_toolTip(void* self, long long* __out_len) {
	auto&& __r = ((QWidget*)self)->toolTip();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QWidget_setToolTipDuration(void* self, int msec) {
	((QWidget*)self)->setToolTipDuration(msec);
}
int QWidget_toolTipDuration(void* self) {
	auto&& __r = ((QWidget*)self)->toolTipDuration();
	return __r;
}
void QWidget_setStatusTip(void* self, const char* arg1, long long arg1_len) {
	QString arg1_qstr = QString::fromUtf8(arg1, arg1_len);
	((QWidget*)self)->setStatusTip(arg1_qstr);
}
char* QWidget_statusTip(void* self, long long* __out_len) {
	auto&& __r = ((QWidget*)self)->statusTip();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QWidget_setWhatsThis(void* self, const char* arg1, long long arg1_len) {
	QString arg1_qstr = QString::fromUtf8(arg1, arg1_len);
	((QWidget*)self)->setWhatsThis(arg1_qstr);
}
char* QWidget_whatsThis(void* self, long long* __out_len) {
	auto&& __r = ((QWidget*)self)->whatsThis();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QWidget_accessibleName(void* self, long long* __out_len) {
	auto&& __r = ((QWidget*)self)->accessibleName();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QWidget_setAccessibleName(void* self, const char* name, long long name_len) {
	QString name_qstr = QString::fromUtf8(name, name_len);
	((QWidget*)self)->setAccessibleName(name_qstr);
}
char* QWidget_accessibleDescription(void* self, long long* __out_len) {
	auto&& __r = ((QWidget*)self)->accessibleDescription();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QWidget_setAccessibleDescription(void* self, const char* description, long long description_len) {
	QString description_qstr = QString::fromUtf8(description, description_len);
	((QWidget*)self)->setAccessibleDescription(description_qstr);
}
char* QWidget_accessibleIdentifier(void* self, long long* __out_len) {
	auto&& __r = ((QWidget*)self)->accessibleIdentifier();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QWidget_setAccessibleIdentifier(void* self, const char* identifier, long long identifier_len) {
	QString identifier_qstr = QString::fromUtf8(identifier, identifier_len);
	((QWidget*)self)->setAccessibleIdentifier(identifier_qstr);
}
void QWidget_setLayoutDirection(void* self, int direction) {
	((QWidget*)self)->setLayoutDirection((Qt::LayoutDirection)(direction));
}
int QWidget_layoutDirection(void* self) {
	auto&& __r = ((QWidget*)self)->layoutDirection();
	return (int)(__r);
}
void QWidget_unsetLayoutDirection(void* self) {
	((QWidget*)self)->unsetLayoutDirection();
}
void QWidget_setLocale(void* self, void* locale) {
	((QWidget*)self)->setLocale(*(QLocale*)locale);
}
void* QWidget_locale(void* self) {
	auto&& __r = ((QWidget*)self)->locale();
	return new QLocale(__r);
}
void QWidget_unsetLocale(void* self) {
	((QWidget*)self)->unsetLocale();
}
int QWidget_isRightToLeft(void* self) {
	auto&& __r = ((QWidget*)self)->isRightToLeft();
	return __r;
}
int QWidget_isLeftToRight(void* self) {
	auto&& __r = ((QWidget*)self)->isLeftToRight();
	return __r;
}
void QWidget_setFocus(void* self) {
	((QWidget*)self)->setFocus();
}
int QWidget_isActiveWindow(void* self) {
	auto&& __r = ((QWidget*)self)->isActiveWindow();
	return __r;
}
void QWidget_activateWindow(void* self) {
	((QWidget*)self)->activateWindow();
}
void QWidget_clearFocus(void* self) {
	((QWidget*)self)->clearFocus();
}
void QWidget_setFocus2(void* self, int reason) {
	((QWidget*)self)->setFocus((Qt::FocusReason)(reason));
}
int QWidget_focusPolicy(void* self) {
	auto&& __r = ((QWidget*)self)->focusPolicy();
	return (int)(__r);
}
void QWidget_setFocusPolicy(void* self, int policy) {
	((QWidget*)self)->setFocusPolicy((Qt::FocusPolicy)(policy));
}
int QWidget_hasFocus(void* self) {
	auto&& __r = ((QWidget*)self)->hasFocus();
	return __r;
}
void QWidget_setTabOrder(void* arg1, void* arg2) {
	QWidget::setTabOrder((QWidget*)arg1, (QWidget*)arg2);
}
void QWidget_setFocusProxy(void* self, void* arg1) {
	((QWidget*)self)->setFocusProxy((QWidget*)arg1);
}
void* QWidget_focusProxy(void* self) {
	auto&& __r = ((QWidget*)self)->focusProxy();
	return (void*)(__r);
}
int QWidget_contextMenuPolicy(void* self) {
	auto&& __r = ((QWidget*)self)->contextMenuPolicy();
	return (int)(__r);
}
void QWidget_setContextMenuPolicy(void* self, int policy) {
	((QWidget*)self)->setContextMenuPolicy((Qt::ContextMenuPolicy)(policy));
}
void QWidget_grabMouse(void* self) {
	((QWidget*)self)->grabMouse();
}
void QWidget_grabMouse2(void* self, void* arg1) {
	((QWidget*)self)->grabMouse(*(QCursor*)arg1);
}
void QWidget_releaseMouse(void* self) {
	((QWidget*)self)->releaseMouse();
}
void QWidget_grabKeyboard(void* self) {
	((QWidget*)self)->grabKeyboard();
}
void QWidget_releaseKeyboard(void* self) {
	((QWidget*)self)->releaseKeyboard();
}
int QWidget_grabShortcut(void* self, void* key) {
	auto&& __r = ((QWidget*)self)->grabShortcut(*(QKeySequence*)key);
	return __r;
}
void QWidget_releaseShortcut(void* self, int id) {
	((QWidget*)self)->releaseShortcut(id);
}
void QWidget_setShortcutEnabled(void* self, int id) {
	((QWidget*)self)->setShortcutEnabled(id);
}
void QWidget_setShortcutAutoRepeat(void* self, int id) {
	((QWidget*)self)->setShortcutAutoRepeat(id);
}
void* QWidget_mouseGrabber() {
	auto&& __r = QWidget::mouseGrabber();
	return (void*)(__r);
}
void* QWidget_keyboardGrabber() {
	auto&& __r = QWidget::keyboardGrabber();
	return (void*)(__r);
}
int QWidget_updatesEnabled(void* self) {
	auto&& __r = ((QWidget*)self)->updatesEnabled();
	return __r;
}
void QWidget_setUpdatesEnabled(void* self, int enable) {
	((QWidget*)self)->setUpdatesEnabled(enable);
}
void* QWidget_graphicsProxyWidget(void* self) {
	auto&& __r = ((QWidget*)self)->graphicsProxyWidget();
	return (void*)(__r);
}
void QWidget_update(void* self) {
	((QWidget*)self)->update();
}
void QWidget_repaint(void* self) {
	((QWidget*)self)->repaint();
}
void QWidget_update2(void* self, int x, int y, int w, int h) {
	((QWidget*)self)->update(x, y, w, h);
}
void QWidget_update3(void* self, void* arg1) {
	((QWidget*)self)->update(*(QRect*)arg1);
}
void QWidget_update4(void* self, void* arg1) {
	((QWidget*)self)->update(*(QRegion*)arg1);
}
void QWidget_repaint2(void* self, int x, int y, int w, int h) {
	((QWidget*)self)->repaint(x, y, w, h);
}
void QWidget_repaint3(void* self, void* arg1) {
	((QWidget*)self)->repaint(*(QRect*)arg1);
}
void QWidget_repaint4(void* self, void* arg1) {
	((QWidget*)self)->repaint(*(QRegion*)arg1);
}
void QWidget_setVisible(void* self, int visible) {
	((QWidget*)self)->setVisible(visible);
}
void QWidget_setHidden(void* self, int hidden) {
	((QWidget*)self)->setHidden(hidden);
}
void QWidget_show(void* self) {
	((QWidget*)self)->show();
}
void QWidget_hide(void* self) {
	((QWidget*)self)->hide();
}
void QWidget_showMinimized(void* self) {
	((QWidget*)self)->showMinimized();
}
void QWidget_showMaximized(void* self) {
	((QWidget*)self)->showMaximized();
}
void QWidget_showFullScreen(void* self) {
	((QWidget*)self)->showFullScreen();
}
void QWidget_showNormal(void* self) {
	((QWidget*)self)->showNormal();
}
int QWidget_close(void* self) {
	auto&& __r = ((QWidget*)self)->close();
	return __r;
}
void QWidget_raise(void* self) {
	((QWidget*)self)->raise();
}
void QWidget_lower(void* self) {
	((QWidget*)self)->lower();
}
void QWidget_stackUnder(void* self, void* arg1) {
	((QWidget*)self)->stackUnder((QWidget*)arg1);
}
void QWidget_move(void* self, int x, int y) {
	((QWidget*)self)->move(x, y);
}
void QWidget_move2(void* self, void* arg1) {
	((QWidget*)self)->move(*(QPoint*)arg1);
}
void QWidget_resize(void* self, int w, int h) {
	((QWidget*)self)->resize(w, h);
}
void QWidget_resize2(void* self, void* arg1) {
	((QWidget*)self)->resize(*(QSize*)arg1);
}
void QWidget_setGeometry(void* self, int x, int y, int w, int h) {
	((QWidget*)self)->setGeometry(x, y, w, h);
}
void QWidget_setGeometry2(void* self, void* arg1) {
	((QWidget*)self)->setGeometry(*(QRect*)arg1);
}
void* QWidget_saveGeometry(void* self) {
	auto&& __r = ((QWidget*)self)->saveGeometry();
	return new QByteArray(__r);
}
int QWidget_restoreGeometry(void* self, void* geometry) {
	auto&& __r = ((QWidget*)self)->restoreGeometry(*(QByteArray*)geometry);
	return __r;
}
void QWidget_adjustSize(void* self) {
	((QWidget*)self)->adjustSize();
}
int QWidget_isVisible(void* self) {
	auto&& __r = ((QWidget*)self)->isVisible();
	return __r;
}
int QWidget_isVisibleTo(void* self, void* arg1) {
	auto&& __r = ((QWidget*)self)->isVisibleTo((QWidget*)arg1);
	return __r;
}
int QWidget_isHidden(void* self) {
	auto&& __r = ((QWidget*)self)->isHidden();
	return __r;
}
int QWidget_isMinimized(void* self) {
	auto&& __r = ((QWidget*)self)->isMinimized();
	return __r;
}
int QWidget_isMaximized(void* self) {
	auto&& __r = ((QWidget*)self)->isMaximized();
	return __r;
}
int QWidget_isFullScreen(void* self) {
	auto&& __r = ((QWidget*)self)->isFullScreen();
	return __r;
}
int QWidget_windowState(void* self) {
	auto&& __r = ((QWidget*)self)->windowState();
	return (int)(__r);
}
void QWidget_setWindowState(void* self, int state) {
	((QWidget*)self)->setWindowState(QFlags<Qt::WindowState>(state));
}
void QWidget_overrideWindowState(void* self, int state) {
	((QWidget*)self)->overrideWindowState(QFlags<Qt::WindowState>(state));
}
void* QWidget_sizeHint(void* self) {
	auto&& __r = ((QWidget*)self)->sizeHint();
	return new QSize(__r);
}
void* QWidget_minimumSizeHint(void* self) {
	auto&& __r = ((QWidget*)self)->minimumSizeHint();
	return new QSize(__r);
}
void* QWidget_sizePolicy(void* self) {
	auto&& __r = ((QWidget*)self)->sizePolicy();
	return new QSizePolicy(__r);
}
void QWidget_setSizePolicy(void* self, void* arg1) {
	((QWidget*)self)->setSizePolicy(*(QSizePolicy*)arg1);
}
void QWidget_setSizePolicy2(void* self, int horizontal, int vertical) {
	((QWidget*)self)->setSizePolicy((QSizePolicy::Policy)(horizontal), (QSizePolicy::Policy)(vertical));
}
int QWidget_heightForWidth(void* self, int arg1) {
	auto&& __r = ((QWidget*)self)->heightForWidth(arg1);
	return __r;
}
int QWidget_hasHeightForWidth(void* self) {
	auto&& __r = ((QWidget*)self)->hasHeightForWidth();
	return __r;
}
void* QWidget_visibleRegion(void* self) {
	auto&& __r = ((QWidget*)self)->visibleRegion();
	return new QRegion(__r);
}
void QWidget_setContentsMargins(void* self, int left, int top, int right, int bottom) {
	((QWidget*)self)->setContentsMargins(left, top, right, bottom);
}
void QWidget_setContentsMargins2(void* self, void* margins) {
	((QWidget*)self)->setContentsMargins(*(QMargins*)margins);
}
void* QWidget_contentsMargins(void* self) {
	auto&& __r = ((QWidget*)self)->contentsMargins();
	return new QMargins(__r);
}
void* QWidget_contentsRect(void* self) {
	auto&& __r = ((QWidget*)self)->contentsRect();
	return new QRect(__r);
}
void* QWidget_layout(void* self) {
	auto&& __r = ((QWidget*)self)->layout();
	return (void*)(__r);
}
void QWidget_setLayout(void* self, void* arg1) {
	((QWidget*)self)->setLayout((QLayout*)arg1);
}
void QWidget_updateGeometry(void* self) {
	((QWidget*)self)->updateGeometry();
}
void QWidget_setParent(void* self, void* parent) {
	((QWidget*)self)->setParent((QWidget*)parent);
}
void QWidget_setParent2(void* self, void* parent, int f) {
	((QWidget*)self)->setParent((QWidget*)parent, QFlags<Qt::WindowType>(f));
}
void QWidget_scroll(void* self, int dx, int dy) {
	((QWidget*)self)->scroll(dx, dy);
}
void QWidget_scroll2(void* self, int dx, int dy, void* arg3) {
	((QWidget*)self)->scroll(dx, dy, *(QRect*)arg3);
}
void* QWidget_focusWidget(void* self) {
	auto&& __r = ((QWidget*)self)->focusWidget();
	return (void*)(__r);
}
void* QWidget_nextInFocusChain(void* self) {
	auto&& __r = ((QWidget*)self)->nextInFocusChain();
	return (void*)(__r);
}
void* QWidget_previousInFocusChain(void* self) {
	auto&& __r = ((QWidget*)self)->previousInFocusChain();
	return (void*)(__r);
}
int QWidget_acceptDrops(void* self) {
	auto&& __r = ((QWidget*)self)->acceptDrops();
	return __r;
}
void QWidget_setAcceptDrops(void* self, int on) {
	((QWidget*)self)->setAcceptDrops(on);
}
void QWidget_addAction(void* self, void* action) {
	((QWidget*)self)->addAction((QAction*)action);
}
void QWidget_addActions(void* self, void** actions, long long actions_len) {
	QList<QAction*> actions_list; actions_list.reserve(actions_len);
	for (long long __i = 0; __i < actions_len; __i++) actions_list.append((QAction*)actions[__i]);
	((QWidget*)self)->addActions(actions_list);
}
void QWidget_insertActions(void* self, void* before, void** actions, long long actions_len) {
	QList<QAction*> actions_list; actions_list.reserve(actions_len);
	for (long long __i = 0; __i < actions_len; __i++) actions_list.append((QAction*)actions[__i]);
	((QWidget*)self)->insertActions((QAction*)before, actions_list);
}
void QWidget_insertAction(void* self, void* before, void* action) {
	((QWidget*)self)->insertAction((QAction*)before, (QAction*)action);
}
void QWidget_removeAction(void* self, void* action) {
	((QWidget*)self)->removeAction((QAction*)action);
}
void** QWidget_actions(void* self, long long* __out_len) {
	auto&& __r = ((QWidget*)self)->actions();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
void* QWidget_addAction2(void* self, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = ((QWidget*)self)->addAction(text_qstr);
	return (void*)(__r);
}
void* QWidget_addAction3(void* self, void* icon, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = ((QWidget*)self)->addAction(*(QIcon*)icon, text_qstr);
	return (void*)(__r);
}
void* QWidget_addAction4(void* self, const char* text, long long text_len, void* receiver, const char* member) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = ((QWidget*)self)->addAction(text_qstr, (QObject*)receiver, member);
	return (void*)(__r);
}
void* QWidget_addAction5(void* self, void* icon, const char* text, long long text_len, void* receiver, const char* member) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = ((QWidget*)self)->addAction(*(QIcon*)icon, text_qstr, (QObject*)receiver, member);
	return (void*)(__r);
}
void* QWidget_addAction6(void* self, const char* text, long long text_len, void* shortcut) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = ((QWidget*)self)->addAction(text_qstr, *(QKeySequence*)shortcut);
	return (void*)(__r);
}
void* QWidget_addAction7(void* self, void* icon, const char* text, long long text_len, void* shortcut) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = ((QWidget*)self)->addAction(*(QIcon*)icon, text_qstr, *(QKeySequence*)shortcut);
	return (void*)(__r);
}
void* QWidget_addAction8(void* self, const char* text, long long text_len, void* shortcut, void* receiver, const char* member) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = ((QWidget*)self)->addAction(text_qstr, *(QKeySequence*)shortcut, (QObject*)receiver, member);
	return (void*)(__r);
}
void* QWidget_addAction9(void* self, void* icon, const char* text, long long text_len, void* shortcut, void* receiver, const char* member) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = ((QWidget*)self)->addAction(*(QIcon*)icon, text_qstr, *(QKeySequence*)shortcut, (QObject*)receiver, member);
	return (void*)(__r);
}
void* QWidget_parentWidget(void* self) {
	auto&& __r = ((QWidget*)self)->parentWidget();
	return (void*)(__r);
}
void QWidget_setWindowFlags(void* self, int type_) {
	((QWidget*)self)->setWindowFlags(QFlags<Qt::WindowType>(type_));
}
int QWidget_windowFlags(void* self) {
	auto&& __r = ((QWidget*)self)->windowFlags();
	return (int)(__r);
}
void QWidget_setWindowFlag(void* self, int arg1) {
	((QWidget*)self)->setWindowFlag((Qt::WindowType)(arg1));
}
void QWidget_overrideWindowFlags(void* self, int type_) {
	((QWidget*)self)->overrideWindowFlags(QFlags<Qt::WindowType>(type_));
}
int QWidget_windowType(void* self) {
	auto&& __r = ((QWidget*)self)->windowType();
	return (int)(__r);
}
void* QWidget_find(unsigned long long arg1) {
	auto&& __r = QWidget::find(arg1);
	return (void*)(__r);
}
void* QWidget_childAt(void* self, int x, int y) {
	auto&& __r = ((QWidget*)self)->childAt(x, y);
	return (void*)(__r);
}
void* QWidget_childAt2(void* self, void* p) {
	auto&& __r = ((QWidget*)self)->childAt(*(QPoint*)p);
	return (void*)(__r);
}
void* QWidget_childAt3(void* self, void* p) {
	auto&& __r = ((QWidget*)self)->childAt(*(QPointF*)p);
	return (void*)(__r);
}
void QWidget_setAttribute(void* self, int arg1) {
	((QWidget*)self)->setAttribute((Qt::WidgetAttribute)(arg1));
}
int QWidget_testAttribute(void* self, int arg1) {
	auto&& __r = ((QWidget*)self)->testAttribute((Qt::WidgetAttribute)(arg1));
	return __r;
}
void* QWidget_paintEngine(void* self) {
	auto&& __r = ((QWidget*)self)->paintEngine();
	return (void*)(__r);
}
void QWidget_ensurePolished(void* self) {
	((QWidget*)self)->ensurePolished();
}
int QWidget_isAncestorOf(void* self, void* child) {
	auto&& __r = ((QWidget*)self)->isAncestorOf((QWidget*)child);
	return __r;
}
int QWidget_autoFillBackground(void* self) {
	auto&& __r = ((QWidget*)self)->autoFillBackground();
	return __r;
}
void QWidget_setAutoFillBackground(void* self, int enabled) {
	((QWidget*)self)->setAutoFillBackground(enabled);
}
void* QWidget_backingStore(void* self) {
	auto&& __r = ((QWidget*)self)->backingStore();
	return (void*)(__r);
}
void* QWidget_windowHandle(void* self) {
	auto&& __r = ((QWidget*)self)->windowHandle();
	return (void*)(__r);
}
void* QWidget_screen(void* self) {
	auto&& __r = ((QWidget*)self)->screen();
	return (void*)(__r);
}
void QWidget_setScreen(void* self, void* arg1) {
	((QWidget*)self)->setScreen((QScreen*)arg1);
}
void* QWidget_createWindowContainer(void* window) {
	auto&& __r = QWidget::createWindowContainer((QWindow*)window);
	return (void*)(__r);
}
void QWidget_windowTitleChanged(void* self, const char* title, long long title_len) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	((QWidget*)self)->windowTitleChanged(title_qstr);
}
void QWidget_windowIconChanged(void* self, void* icon) {
	((QWidget*)self)->windowIconChanged(*(QIcon*)icon);
}
void QWidget_windowIconTextChanged(void* self, const char* iconText, long long iconText_len) {
	QString iconText_qstr = QString::fromUtf8(iconText, iconText_len);
	((QWidget*)self)->windowIconTextChanged(iconText_qstr);
}
void QWidget_customContextMenuRequested(void* self, void* pos) {
	((QWidget*)self)->customContextMenuRequested(*(QPoint*)pos);
}
void* QWidget_inputMethodQuery(void* self, int arg1) {
	auto&& __r = ((QWidget*)self)->inputMethodQuery((Qt::InputMethodQuery)(arg1));
	return new QVariant(__r);
}
int QWidget_inputMethodHints(void* self) {
	auto&& __r = ((QWidget*)self)->inputMethodHints();
	return (int)(__r);
}
void QWidget_setInputMethodHints(void* self, int hints) {
	((QWidget*)self)->setInputMethodHints(QFlags<Qt::InputMethodHint>(hints));
}
void QWidget_updateMicroFocus(void* self) {
	((QWidget__Go*)self)->updateMicroFocus();
}
void QWidget_create(void* self) {
	((QWidget__Go*)self)->create();
}
void QWidget_destroy(void* self) {
	((QWidget__Go*)self)->destroy();
}
int QWidget_focusNextChild(void* self) {
	auto&& __r = ((QWidget__Go*)self)->focusNextChild();
	return __r;
}
int QWidget_focusPreviousChild(void* self) {
	auto&& __r = ((QWidget__Go*)self)->focusPreviousChild();
	return __r;
}
void* QWidget_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QWidget::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QWidget_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QWidget::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QWidget_sender(void* self) {
	auto&& __r = ((QWidget__Go*)self)->sender();
	return (void*)(__r);
}
int QWidget_senderSignalIndex(void* self) {
	auto&& __r = ((QWidget__Go*)self)->senderSignalIndex();
	return __r;
}
int QWidget_receivers(void* self, const char* signal) {
	auto&& __r = ((QWidget__Go*)self)->receivers(signal);
	return __r;
}
int QWidget_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QWidget__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
int QWidget_paintingActive(void* self) {
	auto&& __r = ((QWidget*)self)->paintingActive();
	return __r;
}
int QWidget_widthMM(void* self) {
	auto&& __r = ((QWidget*)self)->widthMM();
	return __r;
}
int QWidget_heightMM(void* self) {
	auto&& __r = ((QWidget*)self)->heightMM();
	return __r;
}
int QWidget_logicalDpiX(void* self) {
	auto&& __r = ((QWidget*)self)->logicalDpiX();
	return __r;
}
int QWidget_logicalDpiY(void* self) {
	auto&& __r = ((QWidget*)self)->logicalDpiY();
	return __r;
}
int QWidget_physicalDpiX(void* self) {
	auto&& __r = ((QWidget*)self)->physicalDpiX();
	return __r;
}
int QWidget_physicalDpiY(void* self) {
	auto&& __r = ((QWidget*)self)->physicalDpiY();
	return __r;
}
double QWidget_devicePixelRatio(void* self) {
	auto&& __r = ((QWidget*)self)->devicePixelRatio();
	return __r;
}
double QWidget_devicePixelRatioF(void* self) {
	auto&& __r = ((QWidget*)self)->devicePixelRatioF();
	return __r;
}
int QWidget_colorCount(void* self) {
	auto&& __r = ((QWidget*)self)->colorCount();
	return __r;
}
int QWidget_depth(void* self) {
	auto&& __r = ((QWidget*)self)->depth();
	return __r;
}
double QWidget_devicePixelRatioFScale() {
	auto&& __r = QWidget::devicePixelRatioFScale();
	return __r;
}
int QWidget_encodeMetricF(int metric, double value) {
	auto&& __r = QWidget::encodeMetricF((QPaintDevice::PaintDeviceMetric)(metric), value);
	return __r;
}
double QWidget_getDecodedMetricF(void* self, int metricA, int metricB) {
	auto&& __r = ((QWidget__Go*)self)->getDecodedMetricF((QPaintDevice::PaintDeviceMetric)(metricA), (QPaintDevice::PaintDeviceMetric)(metricB));
	return __r;
}
char* QWidget_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QWidget::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QWidget_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QWidget::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QWidget_render3(void* self, void* target, void* targetOffset) {
	((QWidget*)self)->render((QPaintDevice*)target, *(QPoint*)targetOffset);
}
void QWidget_render4(void* self, void* target, void* targetOffset, void* sourceRegion) {
	((QWidget*)self)->render((QPaintDevice*)target, *(QPoint*)targetOffset, *(QRegion*)sourceRegion);
}
void QWidget_render5(void* self, void* target, void* targetOffset, void* sourceRegion, int renderFlags) {
	((QWidget*)self)->render((QPaintDevice*)target, *(QPoint*)targetOffset, *(QRegion*)sourceRegion, QFlags<QWidget::RenderFlag>(renderFlags));
}
void QWidget_render6(void* self, void* painter, void* targetOffset) {
	((QWidget*)self)->render((QPainter*)painter, *(QPoint*)targetOffset);
}
void QWidget_render7(void* self, void* painter, void* targetOffset, void* sourceRegion) {
	((QWidget*)self)->render((QPainter*)painter, *(QPoint*)targetOffset, *(QRegion*)sourceRegion);
}
void QWidget_render8(void* self, void* painter, void* targetOffset, void* sourceRegion, int renderFlags) {
	((QWidget*)self)->render((QPainter*)painter, *(QPoint*)targetOffset, *(QRegion*)sourceRegion, QFlags<QWidget::RenderFlag>(renderFlags));
}
void* QWidget_grab2(void* self, void* rectangle) {
	auto&& __r = ((QWidget*)self)->grab(*(QRect*)rectangle);
	return new QPixmap(__r);
}
void QWidget_grabGesture2(void* self, int type_, int flags) {
	((QWidget*)self)->grabGesture((Qt::GestureType)(type_), QFlags<Qt::GestureFlag>(flags));
}
int QWidget_grabShortcut2(void* self, void* key, int context) {
	auto&& __r = ((QWidget*)self)->grabShortcut(*(QKeySequence*)key, (Qt::ShortcutContext)(context));
	return __r;
}
void QWidget_setShortcutEnabled2(void* self, int id, int enable) {
	((QWidget*)self)->setShortcutEnabled(id, enable);
}
void QWidget_setShortcutAutoRepeat2(void* self, int id, int enable) {
	((QWidget*)self)->setShortcutAutoRepeat(id, enable);
}
void* QWidget_addAction10(void* self, const char* text, long long text_len, void* receiver, const char* member, int type_) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = ((QWidget*)self)->addAction(text_qstr, (QObject*)receiver, member, (Qt::ConnectionType)(type_));
	return (void*)(__r);
}
void* QWidget_addAction11(void* self, void* icon, const char* text, long long text_len, void* receiver, const char* member, int type_) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = ((QWidget*)self)->addAction(*(QIcon*)icon, text_qstr, (QObject*)receiver, member, (Qt::ConnectionType)(type_));
	return (void*)(__r);
}
void* QWidget_addAction12(void* self, const char* text, long long text_len, void* shortcut, void* receiver, const char* member, int type_) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = ((QWidget*)self)->addAction(text_qstr, *(QKeySequence*)shortcut, (QObject*)receiver, member, (Qt::ConnectionType)(type_));
	return (void*)(__r);
}
void* QWidget_addAction13(void* self, void* icon, const char* text, long long text_len, void* shortcut, void* receiver, const char* member, int type_) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = ((QWidget*)self)->addAction(*(QIcon*)icon, text_qstr, *(QKeySequence*)shortcut, (QObject*)receiver, member, (Qt::ConnectionType)(type_));
	return (void*)(__r);
}
void QWidget_setWindowFlag2(void* self, int arg1, int on) {
	((QWidget*)self)->setWindowFlag((Qt::WindowType)(arg1), on);
}
void QWidget_setAttribute2(void* self, int arg1, int on) {
	((QWidget*)self)->setAttribute((Qt::WidgetAttribute)(arg1), on);
}
void* QWidget_createWindowContainer2(void* window, void* parent) {
	auto&& __r = QWidget::createWindowContainer((QWindow*)window, (QWidget*)parent);
	return (void*)(__r);
}
void* QWidget_createWindowContainer3(void* window, void* parent, int flags) {
	auto&& __r = QWidget::createWindowContainer((QWindow*)window, (QWidget*)parent, QFlags<Qt::WindowType>(flags));
	return (void*)(__r);
}
void QWidget_updateMicroFocus2(void* self, int query) {
	((QWidget__Go*)self)->updateMicroFocus((Qt::InputMethodQuery)(query));
}
void QWidget_create2(void* self, unsigned long long arg1) {
	((QWidget__Go*)self)->create(arg1);
}
void QWidget_create3(void* self, unsigned long long arg1, int initializeWindow) {
	((QWidget__Go*)self)->create(arg1, initializeWindow);
}
void QWidget_create4(void* self, unsigned long long arg1, int initializeWindow, int destroyOldWindow) {
	((QWidget__Go*)self)->create(arg1, initializeWindow, destroyOldWindow);
}
void QWidget_destroy2(void* self, int destroyWindow) {
	((QWidget__Go*)self)->destroy(destroyWindow);
}
void QWidget_destroy3(void* self, int destroyWindow, int destroySubWindows) {
	((QWidget__Go*)self)->destroy(destroyWindow, destroySubWindows);
}
void* QWidget_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QWidget::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QWidget_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QWidget::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QWidget_asQPaintDevice(void* self) {
	auto&& __r = static_cast<QPaintDevice*>((QWidget*)self);
	return (void*)(__r);
}
int QWidget_devType_default(void* self) {
	auto&& __r = ((QWidget*)self)->QWidget::devType();
	return __r;
}
void QWidget_setVisible_default(void* self, int visible) {
	((QWidget*)self)->QWidget::setVisible(visible);
}
void* QWidget_sizeHint_default(void* self) {
	auto&& __r = ((QWidget*)self)->QWidget::sizeHint();
	return new QSize(__r);
}
void* QWidget_minimumSizeHint_default(void* self) {
	auto&& __r = ((QWidget*)self)->QWidget::minimumSizeHint();
	return new QSize(__r);
}
int QWidget_heightForWidth_default(void* self, int arg1) {
	auto&& __r = ((QWidget*)self)->QWidget::heightForWidth(arg1);
	return __r;
}
int QWidget_hasHeightForWidth_default(void* self) {
	auto&& __r = ((QWidget*)self)->QWidget::hasHeightForWidth();
	return __r;
}
void* QWidget_paintEngine_default(void* self) {
	auto&& __r = ((QWidget*)self)->QWidget::paintEngine();
	return (void*)(__r);
}
void* QWidget_inputMethodQuery_default(void* self, int arg1) {
	auto&& __r = ((QWidget*)self)->QWidget::inputMethodQuery((Qt::InputMethodQuery)(arg1));
	return new QVariant(__r);
}
void* QWidget_metaObject_default(void* self) {
	auto&& __r = ((QWidget*)self)->QWidget::metaObject();
	return (void*)(__r);
}
int QWidget_event_default(void* self, void* event) {
	auto&& __r = ((QWidget__Go*)self)->cppgoBase_event((QEvent*)event);
	return __r;
}
void QWidget_mousePressEvent_default(void* self, void* event) {
	((QWidget__Go*)self)->cppgoBase_mousePressEvent((QMouseEvent*)event);
}
void QWidget_mouseReleaseEvent_default(void* self, void* event) {
	((QWidget__Go*)self)->cppgoBase_mouseReleaseEvent((QMouseEvent*)event);
}
void QWidget_mouseDoubleClickEvent_default(void* self, void* event) {
	((QWidget__Go*)self)->cppgoBase_mouseDoubleClickEvent((QMouseEvent*)event);
}
void QWidget_mouseMoveEvent_default(void* self, void* event) {
	((QWidget__Go*)self)->cppgoBase_mouseMoveEvent((QMouseEvent*)event);
}
void QWidget_wheelEvent_default(void* self, void* event) {
	((QWidget__Go*)self)->cppgoBase_wheelEvent((QWheelEvent*)event);
}
void QWidget_keyPressEvent_default(void* self, void* event) {
	((QWidget__Go*)self)->cppgoBase_keyPressEvent((QKeyEvent*)event);
}
void QWidget_keyReleaseEvent_default(void* self, void* event) {
	((QWidget__Go*)self)->cppgoBase_keyReleaseEvent((QKeyEvent*)event);
}
void QWidget_focusInEvent_default(void* self, void* event) {
	((QWidget__Go*)self)->cppgoBase_focusInEvent((QFocusEvent*)event);
}
void QWidget_focusOutEvent_default(void* self, void* event) {
	((QWidget__Go*)self)->cppgoBase_focusOutEvent((QFocusEvent*)event);
}
void QWidget_enterEvent_default(void* self, void* event) {
	((QWidget__Go*)self)->cppgoBase_enterEvent((QEnterEvent*)event);
}
void QWidget_leaveEvent_default(void* self, void* event) {
	((QWidget__Go*)self)->cppgoBase_leaveEvent((QEvent*)event);
}
void QWidget_paintEvent_default(void* self, void* event) {
	((QWidget__Go*)self)->cppgoBase_paintEvent((QPaintEvent*)event);
}
void QWidget_moveEvent_default(void* self, void* event) {
	((QWidget__Go*)self)->cppgoBase_moveEvent((QMoveEvent*)event);
}
void QWidget_resizeEvent_default(void* self, void* event) {
	((QWidget__Go*)self)->cppgoBase_resizeEvent((QResizeEvent*)event);
}
void QWidget_closeEvent_default(void* self, void* event) {
	((QWidget__Go*)self)->cppgoBase_closeEvent((QCloseEvent*)event);
}
void QWidget_contextMenuEvent_default(void* self, void* event) {
	((QWidget__Go*)self)->cppgoBase_contextMenuEvent((QContextMenuEvent*)event);
}
void QWidget_tabletEvent_default(void* self, void* event) {
	((QWidget__Go*)self)->cppgoBase_tabletEvent((QTabletEvent*)event);
}
void QWidget_actionEvent_default(void* self, void* event) {
	((QWidget__Go*)self)->cppgoBase_actionEvent((QActionEvent*)event);
}
void QWidget_dragEnterEvent_default(void* self, void* event) {
	((QWidget__Go*)self)->cppgoBase_dragEnterEvent((QDragEnterEvent*)event);
}
void QWidget_dragMoveEvent_default(void* self, void* event) {
	((QWidget__Go*)self)->cppgoBase_dragMoveEvent((QDragMoveEvent*)event);
}
void QWidget_dragLeaveEvent_default(void* self, void* event) {
	((QWidget__Go*)self)->cppgoBase_dragLeaveEvent((QDragLeaveEvent*)event);
}
void QWidget_dropEvent_default(void* self, void* event) {
	((QWidget__Go*)self)->cppgoBase_dropEvent((QDropEvent*)event);
}
void QWidget_showEvent_default(void* self, void* event) {
	((QWidget__Go*)self)->cppgoBase_showEvent((QShowEvent*)event);
}
void QWidget_hideEvent_default(void* self, void* event) {
	((QWidget__Go*)self)->cppgoBase_hideEvent((QHideEvent*)event);
}
void QWidget_changeEvent_default(void* self, void* arg1) {
	((QWidget__Go*)self)->cppgoBase_changeEvent((QEvent*)arg1);
}
int QWidget_metric_default(void* self, int arg1) {
	auto&& __r = ((QWidget__Go*)self)->cppgoBase_metric(arg1);
	return __r;
}
void QWidget_initPainter_default(void* self, void* painter) {
	((QWidget__Go*)self)->cppgoBase_initPainter((QPainter*)painter);
}
void* QWidget_redirected_default(void* self, void* offset) {
	auto&& __r = ((QWidget__Go*)self)->cppgoBase_redirected((QPoint*)offset);
	return (void*)(__r);
}
void* QWidget_sharedPainter_default(void* self) {
	auto&& __r = ((QWidget__Go*)self)->cppgoBase_sharedPainter();
	return (void*)(__r);
}
void QWidget_inputMethodEvent_default(void* self, void* arg1) {
	((QWidget__Go*)self)->cppgoBase_inputMethodEvent((QInputMethodEvent*)arg1);
}
int QWidget_focusNextPrevChild_default(void* self, int next) {
	auto&& __r = ((QWidget__Go*)self)->cppgoBase_focusNextPrevChild(next);
	return __r;
}
int QWidget_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QWidget*)self)->QWidget::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QWidget_timerEvent_default(void* self, void* event) {
	((QWidget__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QWidget_childEvent_default(void* self, void* event) {
	((QWidget__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QWidget_customEvent_default(void* self, void* event) {
	((QWidget__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QWidget_connectNotify_default(void* self, void* signal) {
	((QWidget__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QWidget_disconnectNotify_default(void* self, void* signal) {
	((QWidget__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QWidget_connect_windowTitleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QWidget*)self, &QWidget::windowTitleChanged, (QWidget*)self, [__id](const QString& title){
		QByteArray title_ba = title.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)title_ba.constData(), title_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QWidget);
	__conns_QWidget[__id] = __c;
}
void QWidget_connect_windowIconChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QWidget*)self, &QWidget::windowIconChanged, (QWidget*)self, [__id](const QIcon& icon){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&icon);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QWidget);
	__conns_QWidget[__id] = __c;
}
void QWidget_connect_windowIconTextChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QWidget*)self, &QWidget::windowIconTextChanged, (QWidget*)self, [__id](const QString& iconText){
		QByteArray iconText_ba = iconText.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)iconText_ba.constData(), iconText_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QWidget);
	__conns_QWidget[__id] = __c;
}
void QWidget_connect_customContextMenuRequested(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QWidget*)self, &QWidget::customContextMenuRequested, (QWidget*)self, [__id](const QPoint& pos){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&pos);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QWidget);
	__conns_QWidget[__id] = __c;
}
void QWidget_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QWidget*)self, &QWidget::destroyed, (QWidget*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QWidget);
	__conns_QWidget[__id] = __c;
}
void QWidget_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QWidget*)self, &QWidget::destroyed, (QWidget*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QWidget);
	__conns_QWidget[__id] = __c;
}
void QWidget_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QWidget*)self, &QWidget::objectNameChanged, (QWidget*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QWidget);
	__conns_QWidget[__id] = __c;
}
void QWidget__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QWidget);
		auto __it = __conns_QWidget.find(id);
		if (__it == __conns_QWidget.end()) return;
		__c = __it->second;
		__conns_QWidget.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QWidget__Go_new(long long id) { return new QWidget__Go(id); }
void QWidget_delete(void* self) { delete (QWidget*)self; }

}

#include <QtWidgets/qstyleoption.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qsize.h>
#include <QtGui/qfont.h>
#include <QtCore/qlocale.h>
#include <QtCore/qabstractitemmodel.h>
#include <QtGui/qicon.h>
#include <QtCore/qstring.h>
#include <QtGui/qbrush.h>
#include "gen_qstyleoptionviewitem.h"

extern "C" {

void* QStyleOptionViewItem_base(void* self) { return (void*)static_cast<QStyleOption*>((QStyleOptionViewItem*)self); }
void* QStyleOptionViewItem_new() {
	return new QStyleOptionViewItem();
}
void* QStyleOptionViewItem_new2(void* other) {
	return new QStyleOptionViewItem(*(QStyleOptionViewItem*)other);
}
void* QStyleOptionViewItem_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QStyleOptionViewItem*)self)->operator=(*(QStyleOptionViewItem*)arg1);
	return (void*)&(__r);
}
int QStyleOptionViewItem_displayAlignment_get(void* self) {
	auto __r = ((QStyleOptionViewItem*)self)->displayAlignment;
	return (int)(__r);
}
void QStyleOptionViewItem_displayAlignment_set(void* self, int displayAlignment) {
	((QStyleOptionViewItem*)self)->displayAlignment = QFlags<Qt::AlignmentFlag>(displayAlignment);
}
int QStyleOptionViewItem_decorationAlignment_get(void* self) {
	auto __r = ((QStyleOptionViewItem*)self)->decorationAlignment;
	return (int)(__r);
}
void QStyleOptionViewItem_decorationAlignment_set(void* self, int decorationAlignment) {
	((QStyleOptionViewItem*)self)->decorationAlignment = QFlags<Qt::AlignmentFlag>(decorationAlignment);
}
int QStyleOptionViewItem_textElideMode_get(void* self) {
	auto __r = ((QStyleOptionViewItem*)self)->textElideMode;
	return (int)(__r);
}
void QStyleOptionViewItem_textElideMode_set(void* self, int textElideMode) {
	((QStyleOptionViewItem*)self)->textElideMode = (Qt::TextElideMode)(textElideMode);
}
int QStyleOptionViewItem_decorationPosition_get(void* self) {
	auto __r = ((QStyleOptionViewItem*)self)->decorationPosition;
	return (int)(__r);
}
void QStyleOptionViewItem_decorationPosition_set(void* self, int decorationPosition) {
	((QStyleOptionViewItem*)self)->decorationPosition = (QStyleOptionViewItem::Position)(decorationPosition);
}
void* QStyleOptionViewItem_decorationSize_get(void* self) {
	auto __r = ((QStyleOptionViewItem*)self)->decorationSize;
	return new QSize(__r);
}
void QStyleOptionViewItem_decorationSize_set(void* self, void* decorationSize) {
	((QStyleOptionViewItem*)self)->decorationSize = *(QSize*)decorationSize;
}
void* QStyleOptionViewItem_font_get(void* self) {
	auto __r = ((QStyleOptionViewItem*)self)->font;
	return new QFont(__r);
}
void QStyleOptionViewItem_font_set(void* self, void* font) {
	((QStyleOptionViewItem*)self)->font = *(QFont*)font;
}
int QStyleOptionViewItem_showDecorationSelected_get(void* self) {
	auto __r = ((QStyleOptionViewItem*)self)->showDecorationSelected;
	return __r;
}
void QStyleOptionViewItem_showDecorationSelected_set(void* self, int showDecorationSelected) {
	((QStyleOptionViewItem*)self)->showDecorationSelected = showDecorationSelected;
}
int QStyleOptionViewItem_features_get(void* self) {
	auto __r = ((QStyleOptionViewItem*)self)->features;
	return (int)(__r);
}
void QStyleOptionViewItem_features_set(void* self, int features) {
	((QStyleOptionViewItem*)self)->features = QFlags<QStyleOptionViewItem::ViewItemFeature>(features);
}
void* QStyleOptionViewItem_locale_get(void* self) {
	auto __r = ((QStyleOptionViewItem*)self)->locale;
	return new QLocale(__r);
}
void QStyleOptionViewItem_locale_set(void* self, void* locale) {
	((QStyleOptionViewItem*)self)->locale = *(QLocale*)locale;
}
void* QStyleOptionViewItem_widget_get(void* self) {
	auto __r = ((QStyleOptionViewItem*)self)->widget;
	return (void*)(__r);
}
void QStyleOptionViewItem_widget_set(void* self, void* widget) {
	((QStyleOptionViewItem*)self)->widget = (QWidget*)widget;
}
void* QStyleOptionViewItem_index_get(void* self) {
	auto __r = ((QStyleOptionViewItem*)self)->index;
	return new QModelIndex(__r);
}
void QStyleOptionViewItem_index_set(void* self, void* index) {
	((QStyleOptionViewItem*)self)->index = *(QModelIndex*)index;
}
int QStyleOptionViewItem_checkState_get(void* self) {
	auto __r = ((QStyleOptionViewItem*)self)->checkState;
	return (int)(__r);
}
void QStyleOptionViewItem_checkState_set(void* self, int checkState) {
	((QStyleOptionViewItem*)self)->checkState = (Qt::CheckState)(checkState);
}
void* QStyleOptionViewItem_icon_get(void* self) {
	auto __r = ((QStyleOptionViewItem*)self)->icon;
	return new QIcon(__r);
}
void QStyleOptionViewItem_icon_set(void* self, void* icon) {
	((QStyleOptionViewItem*)self)->icon = *(QIcon*)icon;
}
char* QStyleOptionViewItem_text_get(void* self, long long* __out_len) {
	auto __r = ((QStyleOptionViewItem*)self)->text;
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QStyleOptionViewItem_text_set(void* self, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	((QStyleOptionViewItem*)self)->text = text_qstr;
}
int QStyleOptionViewItem_viewItemPosition_get(void* self) {
	auto __r = ((QStyleOptionViewItem*)self)->viewItemPosition;
	return (int)(__r);
}
void QStyleOptionViewItem_viewItemPosition_set(void* self, int viewItemPosition) {
	((QStyleOptionViewItem*)self)->viewItemPosition = (QStyleOptionViewItem::ViewItemPosition)(viewItemPosition);
}
void* QStyleOptionViewItem_backgroundBrush_get(void* self) {
	auto __r = ((QStyleOptionViewItem*)self)->backgroundBrush;
	return new QBrush(__r);
}
void QStyleOptionViewItem_backgroundBrush_set(void* self, void* backgroundBrush) {
	((QStyleOptionViewItem*)self)->backgroundBrush = *(QBrush*)backgroundBrush;
}
void QStyleOptionViewItem_delete(void* self) { delete (QStyleOptionViewItem*)self; }

}

#include <QtWidgets/qlayoutitem.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qsize.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qrect.h>
#include <QtWidgets/qsizepolicy.h>
#include <QtWidgets/qlayout.h>
#include "gen_qwidgetitemv2.h"

extern "C" {

void* QWidgetItemV2_base(void* self) { return (void*)static_cast<QWidgetItem*>((QWidgetItemV2*)self); }
void QWidgetItemV2_chain(void* self, void** out) {
	QWidgetItemV2* __p = (QWidgetItemV2*)self;
	QWidgetItem* __b0 = static_cast<QWidgetItem*>(__p); out[0] = (void*)__b0;
	QLayoutItem* __b1 = static_cast<QLayoutItem*>(__b0); out[1] = (void*)__b1;
}
void* QWidgetItemV2_new(void* widget) {
	return new QWidgetItemV2((QWidget*)widget);
}
void QWidgetItemV2_delete(void* self) { delete (QWidgetItemV2*)self; }

}

#include <QtWidgets/qdockwidget.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qstring.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qobjectdefs.h>
#include <QtGui/qaction.h>
#include <QtCore/qcoreevent.h>
#include <QtGui/qevent.h>
#include <QtWidgets/qstyleoption.h>
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
#include "gen_qdockwidget.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QDockWidget__Go : public QDockWidget {
public:
	long long __goId;
	unsigned char __goOn[49] = {};
	template<typename... A> QDockWidget__Go(long long id, A&&... a) : QDockWidget(std::forward<A>(a)...), __goId(id) {}
	using QDockWidget::updateMicroFocus;
	using QDockWidget::create;
	using QDockWidget::destroy;
	using QDockWidget::focusNextChild;
	using QDockWidget::focusPreviousChild;
	using QDockWidget::sender;
	using QDockWidget::senderSignalIndex;
	using QDockWidget::receivers;
	using QDockWidget::isSignalConnected;
	using QDockWidget::getDecodedMetricF;
	void cppgoBase_changeEvent(QEvent* __a0) { QDockWidget::changeEvent(__a0); }
	void cppgoBase_closeEvent(QCloseEvent* __a0) { QDockWidget::closeEvent(__a0); }
	void cppgoBase_paintEvent(QPaintEvent* __a0) { QDockWidget::paintEvent(__a0); }
	bool cppgoBase_event(QEvent* __a0) { return QDockWidget::event(__a0); }
	void cppgoBase_initStyleOption(QStyleOptionDockWidget* __a0) { QDockWidget::initStyleOption(__a0); }
	void cppgoBase_mousePressEvent(QMouseEvent* __a0) { QDockWidget::mousePressEvent(__a0); }
	void cppgoBase_mouseReleaseEvent(QMouseEvent* __a0) { QDockWidget::mouseReleaseEvent(__a0); }
	void cppgoBase_mouseDoubleClickEvent(QMouseEvent* __a0) { QDockWidget::mouseDoubleClickEvent(__a0); }
	void cppgoBase_mouseMoveEvent(QMouseEvent* __a0) { QDockWidget::mouseMoveEvent(__a0); }
	void cppgoBase_wheelEvent(QWheelEvent* __a0) { QDockWidget::wheelEvent(__a0); }
	void cppgoBase_keyPressEvent(QKeyEvent* __a0) { QDockWidget::keyPressEvent(__a0); }
	void cppgoBase_keyReleaseEvent(QKeyEvent* __a0) { QDockWidget::keyReleaseEvent(__a0); }
	void cppgoBase_focusInEvent(QFocusEvent* __a0) { QDockWidget::focusInEvent(__a0); }
	void cppgoBase_focusOutEvent(QFocusEvent* __a0) { QDockWidget::focusOutEvent(__a0); }
	void cppgoBase_enterEvent(QEnterEvent* __a0) { QDockWidget::enterEvent(__a0); }
	void cppgoBase_leaveEvent(QEvent* __a0) { QDockWidget::leaveEvent(__a0); }
	void cppgoBase_moveEvent(QMoveEvent* __a0) { QDockWidget::moveEvent(__a0); }
	void cppgoBase_resizeEvent(QResizeEvent* __a0) { QDockWidget::resizeEvent(__a0); }
	void cppgoBase_contextMenuEvent(QContextMenuEvent* __a0) { QDockWidget::contextMenuEvent(__a0); }
	void cppgoBase_tabletEvent(QTabletEvent* __a0) { QDockWidget::tabletEvent(__a0); }
	void cppgoBase_actionEvent(QActionEvent* __a0) { QDockWidget::actionEvent(__a0); }
	void cppgoBase_dragEnterEvent(QDragEnterEvent* __a0) { QDockWidget::dragEnterEvent(__a0); }
	void cppgoBase_dragMoveEvent(QDragMoveEvent* __a0) { QDockWidget::dragMoveEvent(__a0); }
	void cppgoBase_dragLeaveEvent(QDragLeaveEvent* __a0) { QDockWidget::dragLeaveEvent(__a0); }
	void cppgoBase_dropEvent(QDropEvent* __a0) { QDockWidget::dropEvent(__a0); }
	void cppgoBase_showEvent(QShowEvent* __a0) { QDockWidget::showEvent(__a0); }
	void cppgoBase_hideEvent(QHideEvent* __a0) { QDockWidget::hideEvent(__a0); }
	int cppgoBase_metric(int __a0) { return QDockWidget::metric((QPaintDevice::PaintDeviceMetric)__a0); }
	void cppgoBase_initPainter(QPainter* __a0) { QDockWidget::initPainter(__a0); }
	QPaintDevice* cppgoBase_redirected(QPoint* __a0) { return QDockWidget::redirected(__a0); }
	QPainter* cppgoBase_sharedPainter() { return QDockWidget::sharedPainter(); }
	void cppgoBase_inputMethodEvent(QInputMethodEvent* __a0) { QDockWidget::inputMethodEvent(__a0); }
	bool cppgoBase_focusNextPrevChild(bool __a0) { return QDockWidget::focusNextPrevChild(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QDockWidget::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QDockWidget::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QDockWidget::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QDockWidget::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QDockWidget::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QDockWidget::metaObject();
	}
	void changeEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4959376063826165943LL, (void*)event, &__h); if (__h) return; }
		QDockWidget::changeEvent(event);
	}
	void closeEvent(QCloseEvent* event) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1989937050019464783LL, (void*)event, &__h); if (__h) return; }
		QDockWidget::closeEvent(event);
	}
	void paintEvent(QPaintEvent* event) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5290442163167458465LL, (void*)event, &__h); if (__h) return; }
		QDockWidget::paintEvent(event);
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QDockWidget::event(event);
	}
	void initStyleOption(QStyleOptionDockWidget* option) const override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -488693769771308531LL, (void*)option, &__h); if (__h) return; }
		QDockWidget::initStyleOption(option);
	}
	int devType() const override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8659679893868524044LL, &__h); if (__h) return __r; }
		return QDockWidget::devType();
	}
	void setVisible(bool visible) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 2110182719375245575LL, visible, &__h); if (__h) return; }
		QDockWidget::setVisible(visible);
	}
	QSize sizeHint() const override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 729690261457684101LL, &__h); if (__h) return *(QSize*)__r; }
		return QDockWidget::sizeHint();
	}
	QSize minimumSizeHint() const override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 7049264599837532491LL, &__h); if (__h) return *(QSize*)__r; }
		return QDockWidget::minimumSizeHint();
	}
	int heightForWidth(int arg1) const override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 4311534060495737339LL, arg1, &__h); if (__h) return __r; }
		return QDockWidget::heightForWidth(arg1);
	}
	bool hasHeightForWidth() const override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8266762397995423915LL, &__h); if (__h) return __r; }
		return QDockWidget::hasHeightForWidth();
	}
	QPaintEngine* paintEngine() const override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 6398007404756346381LL, &__h); if (__h) return (QPaintEngine*)__r; }
		return QDockWidget::paintEngine();
	}
	void mousePressEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -332090898078543319LL, (void*)event, &__h); if (__h) return; }
		QDockWidget::mousePressEvent(event);
	}
	void mouseReleaseEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4415920049297795213LL, (void*)event, &__h); if (__h) return; }
		QDockWidget::mouseReleaseEvent(event);
	}
	void mouseDoubleClickEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4029073905697023887LL, (void*)event, &__h); if (__h) return; }
		QDockWidget::mouseDoubleClickEvent(event);
	}
	void mouseMoveEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4255494004099806581LL, (void*)event, &__h); if (__h) return; }
		QDockWidget::mouseMoveEvent(event);
	}
	void wheelEvent(QWheelEvent* event) override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1919988340934511316LL, (void*)event, &__h); if (__h) return; }
		QDockWidget::wheelEvent(event);
	}
	void keyPressEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1023791977679793185LL, (void*)event, &__h); if (__h) return; }
		QDockWidget::keyPressEvent(event);
	}
	void keyReleaseEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 772598475950050299LL, (void*)event, &__h); if (__h) return; }
		QDockWidget::keyReleaseEvent(event);
	}
	void focusInEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 975043767033552552LL, (void*)event, &__h); if (__h) return; }
		QDockWidget::focusInEvent(event);
	}
	void focusOutEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -385596010576048993LL, (void*)event, &__h); if (__h) return; }
		QDockWidget::focusOutEvent(event);
	}
	void enterEvent(QEnterEvent* event) override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6316270346668237895LL, (void*)event, &__h); if (__h) return; }
		QDockWidget::enterEvent(event);
	}
	void leaveEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -551070009623711640LL, (void*)event, &__h); if (__h) return; }
		QDockWidget::leaveEvent(event);
	}
	void moveEvent(QMoveEvent* event) override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -9161949150489525396LL, (void*)event, &__h); if (__h) return; }
		QDockWidget::moveEvent(event);
	}
	void resizeEvent(QResizeEvent* event) override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4302917285655014661LL, (void*)event, &__h); if (__h) return; }
		QDockWidget::resizeEvent(event);
	}
	void contextMenuEvent(QContextMenuEvent* event) override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6698043222227006015LL, (void*)event, &__h); if (__h) return; }
		QDockWidget::contextMenuEvent(event);
	}
	void tabletEvent(QTabletEvent* event) override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -212068866860701141LL, (void*)event, &__h); if (__h) return; }
		QDockWidget::tabletEvent(event);
	}
	void actionEvent(QActionEvent* event) override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1697444140088555581LL, (void*)event, &__h); if (__h) return; }
		QDockWidget::actionEvent(event);
	}
	void dragEnterEvent(QDragEnterEvent* event) override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 206122715919917589LL, (void*)event, &__h); if (__h) return; }
		QDockWidget::dragEnterEvent(event);
	}
	void dragMoveEvent(QDragMoveEvent* event) override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1781545137352745978LL, (void*)event, &__h); if (__h) return; }
		QDockWidget::dragMoveEvent(event);
	}
	void dragLeaveEvent(QDragLeaveEvent* event) override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -3139904158373725006LL, (void*)event, &__h); if (__h) return; }
		QDockWidget::dragLeaveEvent(event);
	}
	void dropEvent(QDropEvent* event) override {
		int __h = 0;
		if (__goOn[32] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 9107058682754591716LL, (void*)event, &__h); if (__h) return; }
		QDockWidget::dropEvent(event);
	}
	void showEvent(QShowEvent* event) override {
		int __h = 0;
		if (__goOn[33] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6825369174751916LL, (void*)event, &__h); if (__h) return; }
		QDockWidget::showEvent(event);
	}
	void hideEvent(QHideEvent* event) override {
		int __h = 0;
		if (__goOn[34] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -199718707876719129LL, (void*)event, &__h); if (__h) return; }
		QDockWidget::hideEvent(event);
	}
	bool nativeEvent(const QByteArray& eventType, void* message, qintptr* result) override {
		int __h = 0;
		if (__goOn[35] && cppgo_virt_int_voidp_voidp_longlongp) { int __r = ((int(*)(long long, long long, void*, void*, long long*, int*))cppgo_virt_int_voidp_voidp_longlongp)(__goId, -2039166586246443706LL, (void*)&eventType, (void*)message, (long long*)result, &__h); if (__h) return __r; }
		return QDockWidget::nativeEvent(eventType, message, result);
	}
	int metric(QPaintDevice::PaintDeviceMetric arg1) const override {
		int __h = 0;
		if (__goOn[36] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 427629649267311777LL, (int)(arg1), &__h); if (__h) return __r; }
		return QDockWidget::metric(arg1);
	}
	void initPainter(QPainter* painter) const override {
		int __h = 0;
		if (__goOn[37] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6477500296980050210LL, (void*)painter, &__h); if (__h) return; }
		QDockWidget::initPainter(painter);
	}
	QPaintDevice* redirected(QPoint* offset) const override {
		int __h = 0;
		if (__goOn[38] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 4127240633310655048LL, (void*)offset, &__h); if (__h) return (QPaintDevice*)__r; }
		return QDockWidget::redirected(offset);
	}
	QPainter* sharedPainter() const override {
		int __h = 0;
		if (__goOn[39] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 5792771328741774991LL, &__h); if (__h) return (QPainter*)__r; }
		return QDockWidget::sharedPainter();
	}
	void inputMethodEvent(QInputMethodEvent* arg1) override {
		int __h = 0;
		if (__goOn[40] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5438819661942843222LL, (void*)arg1, &__h); if (__h) return; }
		QDockWidget::inputMethodEvent(arg1);
	}
	QVariant inputMethodQuery(Qt::InputMethodQuery arg1) const override {
		int __h = 0;
		if (__goOn[41] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7345853097340709638LL, (int)(arg1), &__h); if (__h) return *(QVariant*)__r; }
		return QDockWidget::inputMethodQuery(arg1);
	}
	bool focusNextPrevChild(bool next) override {
		int __h = 0;
		if (__goOn[42] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 5898328517505784191LL, next, &__h); if (__h) return __r; }
		return QDockWidget::focusNextPrevChild(next);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[43] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QDockWidget::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[44] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QDockWidget::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[45] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QDockWidget::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[46] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QDockWidget::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[47] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QDockWidget::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[48] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QDockWidget::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QDockWidget;
static std::mutex __connsMu_QDockWidget;

extern "C" {

void* QDockWidget_base(void* self) { return (void*)static_cast<QWidget*>((QDockWidget*)self); }
void QDockWidget_chain(void* self, void** out) {
	QDockWidget* __p = (QDockWidget*)self;
	QWidget* __b0 = static_cast<QWidget*>(__p); out[0] = (void*)__b0;
	QObject* __b1 = static_cast<QObject*>(__b0); out[1] = (void*)__b1;
}
void QDockWidget__set_override(void* self, int idx) { ((QDockWidget__Go*)self)->__goOn[idx] = 1; }
void* QDockWidget_fromQObject(void* o) { return (void*)qobject_cast<QDockWidget*>((QObject*)o); }
void* QDockWidget_new(long long __goid, void* parent) {
	return new QDockWidget__Go(__goid, (QWidget*)parent);
}
void* QDockWidget_new2(long long __goid, const char* title, long long title_len) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	return new QDockWidget__Go(__goid, title_qstr);
}
void* QDockWidget_new3(long long __goid) {
	return new QDockWidget__Go(__goid);
}
void* QDockWidget_new4(long long __goid, const char* title, long long title_len, void* parent) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	return new QDockWidget__Go(__goid, title_qstr, (QWidget*)parent);
}
void* QDockWidget_new5(long long __goid, const char* title, long long title_len, void* parent, int flags) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	return new QDockWidget__Go(__goid, title_qstr, (QWidget*)parent, QFlags<Qt::WindowType>(flags));
}
void* QDockWidget_new6(long long __goid, void* parent, int flags) {
	return new QDockWidget__Go(__goid, (QWidget*)parent, QFlags<Qt::WindowType>(flags));
}
char* QDockWidget_tr(const char* s, long long* __out_len) {
	auto&& __r = QDockWidget::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QDockWidget_widget(void* self) {
	auto&& __r = ((QDockWidget*)self)->widget();
	return (void*)(__r);
}
void QDockWidget_setWidget(void* self, void* widget) {
	((QDockWidget*)self)->setWidget((QWidget*)widget);
}
void QDockWidget_setFeatures(void* self, int features) {
	((QDockWidget*)self)->setFeatures(QFlags<QDockWidget::DockWidgetFeature>(features));
}
int QDockWidget_features(void* self) {
	auto&& __r = ((QDockWidget*)self)->features();
	return (int)(__r);
}
void QDockWidget_setFloating(void* self, int floating) {
	((QDockWidget*)self)->setFloating(floating);
}
int QDockWidget_isFloating(void* self) {
	auto&& __r = ((QDockWidget*)self)->isFloating();
	return __r;
}
void QDockWidget_setAllowedAreas(void* self, int areas) {
	((QDockWidget*)self)->setAllowedAreas(QFlags<Qt::DockWidgetArea>(areas));
}
int QDockWidget_allowedAreas(void* self) {
	auto&& __r = ((QDockWidget*)self)->allowedAreas();
	return (int)(__r);
}
void QDockWidget_setTitleBarWidget(void* self, void* widget) {
	((QDockWidget*)self)->setTitleBarWidget((QWidget*)widget);
}
void* QDockWidget_titleBarWidget(void* self) {
	auto&& __r = ((QDockWidget*)self)->titleBarWidget();
	return (void*)(__r);
}
void QDockWidget_setDockLocation(void* self, int area) {
	((QDockWidget*)self)->setDockLocation((Qt::DockWidgetArea)(area));
}
int QDockWidget_dockLocation(void* self) {
	auto&& __r = ((QDockWidget*)self)->dockLocation();
	return (int)(__r);
}
int QDockWidget_isAreaAllowed(void* self, int area) {
	auto&& __r = ((QDockWidget*)self)->isAreaAllowed((Qt::DockWidgetArea)(area));
	return __r;
}
void* QDockWidget_toggleViewAction(void* self) {
	auto&& __r = ((QDockWidget*)self)->toggleViewAction();
	return (void*)(__r);
}
void QDockWidget_featuresChanged(void* self, int features) {
	((QDockWidget*)self)->featuresChanged(QFlags<QDockWidget::DockWidgetFeature>(features));
}
void QDockWidget_topLevelChanged(void* self, int topLevel) {
	((QDockWidget*)self)->topLevelChanged(topLevel);
}
void QDockWidget_allowedAreasChanged(void* self, int allowedAreas) {
	((QDockWidget*)self)->allowedAreasChanged(QFlags<Qt::DockWidgetArea>(allowedAreas));
}
void QDockWidget_visibilityChanged(void* self, int visible) {
	((QDockWidget*)self)->visibilityChanged(visible);
}
void QDockWidget_dockLocationChanged(void* self, int area) {
	((QDockWidget*)self)->dockLocationChanged((Qt::DockWidgetArea)(area));
}
void QDockWidget_setTabOrder(void* arg1, void* arg2) {
	QDockWidget::setTabOrder((QWidget*)arg1, (QWidget*)arg2);
}
void* QDockWidget_mouseGrabber() {
	auto&& __r = QDockWidget::mouseGrabber();
	return (void*)(__r);
}
void* QDockWidget_keyboardGrabber() {
	auto&& __r = QDockWidget::keyboardGrabber();
	return (void*)(__r);
}
void* QDockWidget_find(unsigned long long arg1) {
	auto&& __r = QDockWidget::find(arg1);
	return (void*)(__r);
}
void* QDockWidget_createWindowContainer(void* window) {
	auto&& __r = QDockWidget::createWindowContainer((QWindow*)window);
	return (void*)(__r);
}
void QDockWidget_updateMicroFocus(void* self) {
	((QDockWidget__Go*)self)->updateMicroFocus();
}
void QDockWidget_create(void* self) {
	((QDockWidget__Go*)self)->create();
}
void QDockWidget_destroy(void* self) {
	((QDockWidget__Go*)self)->destroy();
}
int QDockWidget_focusNextChild(void* self) {
	auto&& __r = ((QDockWidget__Go*)self)->focusNextChild();
	return __r;
}
int QDockWidget_focusPreviousChild(void* self) {
	auto&& __r = ((QDockWidget__Go*)self)->focusPreviousChild();
	return __r;
}
void* QDockWidget_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QDockWidget::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QDockWidget_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QDockWidget::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QDockWidget_sender(void* self) {
	auto&& __r = ((QDockWidget__Go*)self)->sender();
	return (void*)(__r);
}
int QDockWidget_senderSignalIndex(void* self) {
	auto&& __r = ((QDockWidget__Go*)self)->senderSignalIndex();
	return __r;
}
int QDockWidget_receivers(void* self, const char* signal) {
	auto&& __r = ((QDockWidget__Go*)self)->receivers(signal);
	return __r;
}
int QDockWidget_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QDockWidget__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
double QDockWidget_devicePixelRatioFScale() {
	auto&& __r = QDockWidget::devicePixelRatioFScale();
	return __r;
}
int QDockWidget_encodeMetricF(int metric, double value) {
	auto&& __r = QDockWidget::encodeMetricF((QPaintDevice::PaintDeviceMetric)(metric), value);
	return __r;
}
double QDockWidget_getDecodedMetricF(void* self, int metricA, int metricB) {
	auto&& __r = ((QDockWidget__Go*)self)->getDecodedMetricF((QPaintDevice::PaintDeviceMetric)(metricA), (QPaintDevice::PaintDeviceMetric)(metricB));
	return __r;
}
char* QDockWidget_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QDockWidget::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QDockWidget_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QDockWidget::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QDockWidget_createWindowContainer2(void* window, void* parent) {
	auto&& __r = QDockWidget::createWindowContainer((QWindow*)window, (QWidget*)parent);
	return (void*)(__r);
}
void* QDockWidget_createWindowContainer3(void* window, void* parent, int flags) {
	auto&& __r = QDockWidget::createWindowContainer((QWindow*)window, (QWidget*)parent, QFlags<Qt::WindowType>(flags));
	return (void*)(__r);
}
void QDockWidget_updateMicroFocus2(void* self, int query) {
	((QDockWidget__Go*)self)->updateMicroFocus((Qt::InputMethodQuery)(query));
}
void QDockWidget_create2(void* self, unsigned long long arg1) {
	((QDockWidget__Go*)self)->create(arg1);
}
void QDockWidget_create3(void* self, unsigned long long arg1, int initializeWindow) {
	((QDockWidget__Go*)self)->create(arg1, initializeWindow);
}
void QDockWidget_create4(void* self, unsigned long long arg1, int initializeWindow, int destroyOldWindow) {
	((QDockWidget__Go*)self)->create(arg1, initializeWindow, destroyOldWindow);
}
void QDockWidget_destroy2(void* self, int destroyWindow) {
	((QDockWidget__Go*)self)->destroy(destroyWindow);
}
void QDockWidget_destroy3(void* self, int destroyWindow, int destroySubWindows) {
	((QDockWidget__Go*)self)->destroy(destroyWindow, destroySubWindows);
}
void* QDockWidget_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QDockWidget::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QDockWidget_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QDockWidget::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QDockWidget_metaObject_default(void* self) {
	auto&& __r = ((QDockWidget*)self)->QDockWidget::metaObject();
	return (void*)(__r);
}
void QDockWidget_changeEvent_default(void* self, void* event) {
	((QDockWidget__Go*)self)->cppgoBase_changeEvent((QEvent*)event);
}
void QDockWidget_closeEvent_default(void* self, void* event) {
	((QDockWidget__Go*)self)->cppgoBase_closeEvent((QCloseEvent*)event);
}
void QDockWidget_paintEvent_default(void* self, void* event) {
	((QDockWidget__Go*)self)->cppgoBase_paintEvent((QPaintEvent*)event);
}
int QDockWidget_event_default(void* self, void* event) {
	auto&& __r = ((QDockWidget__Go*)self)->cppgoBase_event((QEvent*)event);
	return __r;
}
void QDockWidget_initStyleOption_default(void* self, void* option) {
	((QDockWidget__Go*)self)->cppgoBase_initStyleOption((QStyleOptionDockWidget*)option);
}
int QDockWidget_devType_default(void* self) {
	auto&& __r = ((QDockWidget*)self)->QDockWidget::devType();
	return __r;
}
void QDockWidget_setVisible_default(void* self, int visible) {
	((QDockWidget*)self)->QDockWidget::setVisible(visible);
}
void* QDockWidget_sizeHint_default(void* self) {
	auto&& __r = ((QDockWidget*)self)->QDockWidget::sizeHint();
	return new QSize(__r);
}
void* QDockWidget_minimumSizeHint_default(void* self) {
	auto&& __r = ((QDockWidget*)self)->QDockWidget::minimumSizeHint();
	return new QSize(__r);
}
int QDockWidget_heightForWidth_default(void* self, int arg1) {
	auto&& __r = ((QDockWidget*)self)->QDockWidget::heightForWidth(arg1);
	return __r;
}
int QDockWidget_hasHeightForWidth_default(void* self) {
	auto&& __r = ((QDockWidget*)self)->QDockWidget::hasHeightForWidth();
	return __r;
}
void* QDockWidget_paintEngine_default(void* self) {
	auto&& __r = ((QDockWidget*)self)->QDockWidget::paintEngine();
	return (void*)(__r);
}
void QDockWidget_mousePressEvent_default(void* self, void* event) {
	((QDockWidget__Go*)self)->cppgoBase_mousePressEvent((QMouseEvent*)event);
}
void QDockWidget_mouseReleaseEvent_default(void* self, void* event) {
	((QDockWidget__Go*)self)->cppgoBase_mouseReleaseEvent((QMouseEvent*)event);
}
void QDockWidget_mouseDoubleClickEvent_default(void* self, void* event) {
	((QDockWidget__Go*)self)->cppgoBase_mouseDoubleClickEvent((QMouseEvent*)event);
}
void QDockWidget_mouseMoveEvent_default(void* self, void* event) {
	((QDockWidget__Go*)self)->cppgoBase_mouseMoveEvent((QMouseEvent*)event);
}
void QDockWidget_wheelEvent_default(void* self, void* event) {
	((QDockWidget__Go*)self)->cppgoBase_wheelEvent((QWheelEvent*)event);
}
void QDockWidget_keyPressEvent_default(void* self, void* event) {
	((QDockWidget__Go*)self)->cppgoBase_keyPressEvent((QKeyEvent*)event);
}
void QDockWidget_keyReleaseEvent_default(void* self, void* event) {
	((QDockWidget__Go*)self)->cppgoBase_keyReleaseEvent((QKeyEvent*)event);
}
void QDockWidget_focusInEvent_default(void* self, void* event) {
	((QDockWidget__Go*)self)->cppgoBase_focusInEvent((QFocusEvent*)event);
}
void QDockWidget_focusOutEvent_default(void* self, void* event) {
	((QDockWidget__Go*)self)->cppgoBase_focusOutEvent((QFocusEvent*)event);
}
void QDockWidget_enterEvent_default(void* self, void* event) {
	((QDockWidget__Go*)self)->cppgoBase_enterEvent((QEnterEvent*)event);
}
void QDockWidget_leaveEvent_default(void* self, void* event) {
	((QDockWidget__Go*)self)->cppgoBase_leaveEvent((QEvent*)event);
}
void QDockWidget_moveEvent_default(void* self, void* event) {
	((QDockWidget__Go*)self)->cppgoBase_moveEvent((QMoveEvent*)event);
}
void QDockWidget_resizeEvent_default(void* self, void* event) {
	((QDockWidget__Go*)self)->cppgoBase_resizeEvent((QResizeEvent*)event);
}
void QDockWidget_contextMenuEvent_default(void* self, void* event) {
	((QDockWidget__Go*)self)->cppgoBase_contextMenuEvent((QContextMenuEvent*)event);
}
void QDockWidget_tabletEvent_default(void* self, void* event) {
	((QDockWidget__Go*)self)->cppgoBase_tabletEvent((QTabletEvent*)event);
}
void QDockWidget_actionEvent_default(void* self, void* event) {
	((QDockWidget__Go*)self)->cppgoBase_actionEvent((QActionEvent*)event);
}
void QDockWidget_dragEnterEvent_default(void* self, void* event) {
	((QDockWidget__Go*)self)->cppgoBase_dragEnterEvent((QDragEnterEvent*)event);
}
void QDockWidget_dragMoveEvent_default(void* self, void* event) {
	((QDockWidget__Go*)self)->cppgoBase_dragMoveEvent((QDragMoveEvent*)event);
}
void QDockWidget_dragLeaveEvent_default(void* self, void* event) {
	((QDockWidget__Go*)self)->cppgoBase_dragLeaveEvent((QDragLeaveEvent*)event);
}
void QDockWidget_dropEvent_default(void* self, void* event) {
	((QDockWidget__Go*)self)->cppgoBase_dropEvent((QDropEvent*)event);
}
void QDockWidget_showEvent_default(void* self, void* event) {
	((QDockWidget__Go*)self)->cppgoBase_showEvent((QShowEvent*)event);
}
void QDockWidget_hideEvent_default(void* self, void* event) {
	((QDockWidget__Go*)self)->cppgoBase_hideEvent((QHideEvent*)event);
}
int QDockWidget_metric_default(void* self, int arg1) {
	auto&& __r = ((QDockWidget__Go*)self)->cppgoBase_metric(arg1);
	return __r;
}
void QDockWidget_initPainter_default(void* self, void* painter) {
	((QDockWidget__Go*)self)->cppgoBase_initPainter((QPainter*)painter);
}
void* QDockWidget_redirected_default(void* self, void* offset) {
	auto&& __r = ((QDockWidget__Go*)self)->cppgoBase_redirected((QPoint*)offset);
	return (void*)(__r);
}
void* QDockWidget_sharedPainter_default(void* self) {
	auto&& __r = ((QDockWidget__Go*)self)->cppgoBase_sharedPainter();
	return (void*)(__r);
}
void QDockWidget_inputMethodEvent_default(void* self, void* arg1) {
	((QDockWidget__Go*)self)->cppgoBase_inputMethodEvent((QInputMethodEvent*)arg1);
}
void* QDockWidget_inputMethodQuery_default(void* self, int arg1) {
	auto&& __r = ((QDockWidget*)self)->QDockWidget::inputMethodQuery((Qt::InputMethodQuery)(arg1));
	return new QVariant(__r);
}
int QDockWidget_focusNextPrevChild_default(void* self, int next) {
	auto&& __r = ((QDockWidget__Go*)self)->cppgoBase_focusNextPrevChild(next);
	return __r;
}
int QDockWidget_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QDockWidget*)self)->QDockWidget::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QDockWidget_timerEvent_default(void* self, void* event) {
	((QDockWidget__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QDockWidget_childEvent_default(void* self, void* event) {
	((QDockWidget__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QDockWidget_customEvent_default(void* self, void* event) {
	((QDockWidget__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QDockWidget_connectNotify_default(void* self, void* signal) {
	((QDockWidget__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QDockWidget_disconnectNotify_default(void* self, void* signal) {
	((QDockWidget__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QDockWidget_connect_featuresChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDockWidget*)self, &QDockWidget::featuresChanged, (QDockWidget*)self, [__id](QFlags<QDockWidget::DockWidgetFeature> features){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, (int)(features));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDockWidget);
	__conns_QDockWidget[__id] = __c;
}
void QDockWidget_connect_topLevelChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDockWidget*)self, &QDockWidget::topLevelChanged, (QDockWidget*)self, [__id](bool topLevel){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, topLevel);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDockWidget);
	__conns_QDockWidget[__id] = __c;
}
void QDockWidget_connect_allowedAreasChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDockWidget*)self, &QDockWidget::allowedAreasChanged, (QDockWidget*)self, [__id](QFlags<Qt::DockWidgetArea> allowedAreas){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, (int)(allowedAreas));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDockWidget);
	__conns_QDockWidget[__id] = __c;
}
void QDockWidget_connect_visibilityChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDockWidget*)self, &QDockWidget::visibilityChanged, (QDockWidget*)self, [__id](bool visible){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, visible);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDockWidget);
	__conns_QDockWidget[__id] = __c;
}
void QDockWidget_connect_dockLocationChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDockWidget*)self, &QDockWidget::dockLocationChanged, (QDockWidget*)self, [__id](Qt::DockWidgetArea area){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, (int)(area));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDockWidget);
	__conns_QDockWidget[__id] = __c;
}
void QDockWidget_connect_windowTitleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDockWidget*)self, &QDockWidget::windowTitleChanged, (QDockWidget*)self, [__id](const QString& title){
		QByteArray title_ba = title.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)title_ba.constData(), title_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDockWidget);
	__conns_QDockWidget[__id] = __c;
}
void QDockWidget_connect_windowIconChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDockWidget*)self, &QDockWidget::windowIconChanged, (QDockWidget*)self, [__id](const QIcon& icon){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&icon);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDockWidget);
	__conns_QDockWidget[__id] = __c;
}
void QDockWidget_connect_windowIconTextChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDockWidget*)self, &QDockWidget::windowIconTextChanged, (QDockWidget*)self, [__id](const QString& iconText){
		QByteArray iconText_ba = iconText.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)iconText_ba.constData(), iconText_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDockWidget);
	__conns_QDockWidget[__id] = __c;
}
void QDockWidget_connect_customContextMenuRequested(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDockWidget*)self, &QDockWidget::customContextMenuRequested, (QDockWidget*)self, [__id](const QPoint& pos){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&pos);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDockWidget);
	__conns_QDockWidget[__id] = __c;
}
void QDockWidget_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDockWidget*)self, &QDockWidget::destroyed, (QDockWidget*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDockWidget);
	__conns_QDockWidget[__id] = __c;
}
void QDockWidget_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDockWidget*)self, &QDockWidget::destroyed, (QDockWidget*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDockWidget);
	__conns_QDockWidget[__id] = __c;
}
void QDockWidget_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QDockWidget*)self, &QDockWidget::objectNameChanged, (QDockWidget*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QDockWidget);
	__conns_QDockWidget[__id] = __c;
}
void QDockWidget__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QDockWidget);
		auto __it = __conns_QDockWidget.find(id);
		if (__it == __conns_QDockWidget.end()) return;
		__c = __it->second;
		__conns_QDockWidget.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QDockWidget__Go_new(long long id) { return new QDockWidget__Go(id); }
void QDockWidget_delete(void* self) { delete (QDockWidget*)self; }

}

#include <QtWidgets/qgraphicsitem.h>
#include <QtGui/qpolygon.h>
#include <QtCore/qnamespace.h>
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
#include <QtGui/qtransform.h>
#include <QtWidgets/qgraphicstransform.h>
#include <QtGui/qregion.h>
#include <QtCore/qcoreevent.h>
#include <QtWidgets/qgraphicssceneevent.h>
#include <QtGui/qevent.h>
#include "gen_qgraphicspolygonitem.h"
#include "gen_purego_cb.h"

class QGraphicsPolygonItem__Go : public QGraphicsPolygonItem {
public:
	long long __goId;
	unsigned char __goOn[35] = {};
	template<typename... A> QGraphicsPolygonItem__Go(long long id, A&&... a) : QGraphicsPolygonItem(std::forward<A>(a)...), __goId(id) {}
	using QGraphicsPolygonItem::updateMicroFocus;
	using QGraphicsPolygonItem::addToIndex;
	using QGraphicsPolygonItem::removeFromIndex;
	using QGraphicsPolygonItem::prepareGeometryChange;
	bool cppgoBase_supportsExtension(int __a0) { return QGraphicsPolygonItem::supportsExtension((QGraphicsItem::Extension)__a0); }
	void cppgoBase_setExtension(int __a0, const QVariant& __a1) { QGraphicsPolygonItem::setExtension((QGraphicsItem::Extension)__a0, __a1); }
	QVariant cppgoBase_extension(const QVariant& __a0) { return QGraphicsPolygonItem::extension(__a0); }
	bool cppgoBase_sceneEventFilter(QGraphicsItem* __a0, QEvent* __a1) { return QGraphicsPolygonItem::sceneEventFilter(__a0, __a1); }
	bool cppgoBase_sceneEvent(QEvent* __a0) { return QGraphicsPolygonItem::sceneEvent(__a0); }
	void cppgoBase_contextMenuEvent(QGraphicsSceneContextMenuEvent* __a0) { QGraphicsPolygonItem::contextMenuEvent(__a0); }
	void cppgoBase_dragEnterEvent(QGraphicsSceneDragDropEvent* __a0) { QGraphicsPolygonItem::dragEnterEvent(__a0); }
	void cppgoBase_dragLeaveEvent(QGraphicsSceneDragDropEvent* __a0) { QGraphicsPolygonItem::dragLeaveEvent(__a0); }
	void cppgoBase_dragMoveEvent(QGraphicsSceneDragDropEvent* __a0) { QGraphicsPolygonItem::dragMoveEvent(__a0); }
	void cppgoBase_dropEvent(QGraphicsSceneDragDropEvent* __a0) { QGraphicsPolygonItem::dropEvent(__a0); }
	void cppgoBase_focusInEvent(QFocusEvent* __a0) { QGraphicsPolygonItem::focusInEvent(__a0); }
	void cppgoBase_focusOutEvent(QFocusEvent* __a0) { QGraphicsPolygonItem::focusOutEvent(__a0); }
	void cppgoBase_hoverEnterEvent(QGraphicsSceneHoverEvent* __a0) { QGraphicsPolygonItem::hoverEnterEvent(__a0); }
	void cppgoBase_hoverMoveEvent(QGraphicsSceneHoverEvent* __a0) { QGraphicsPolygonItem::hoverMoveEvent(__a0); }
	void cppgoBase_hoverLeaveEvent(QGraphicsSceneHoverEvent* __a0) { QGraphicsPolygonItem::hoverLeaveEvent(__a0); }
	void cppgoBase_keyPressEvent(QKeyEvent* __a0) { QGraphicsPolygonItem::keyPressEvent(__a0); }
	void cppgoBase_keyReleaseEvent(QKeyEvent* __a0) { QGraphicsPolygonItem::keyReleaseEvent(__a0); }
	void cppgoBase_mousePressEvent(QGraphicsSceneMouseEvent* __a0) { QGraphicsPolygonItem::mousePressEvent(__a0); }
	void cppgoBase_mouseMoveEvent(QGraphicsSceneMouseEvent* __a0) { QGraphicsPolygonItem::mouseMoveEvent(__a0); }
	void cppgoBase_mouseReleaseEvent(QGraphicsSceneMouseEvent* __a0) { QGraphicsPolygonItem::mouseReleaseEvent(__a0); }
	void cppgoBase_mouseDoubleClickEvent(QGraphicsSceneMouseEvent* __a0) { QGraphicsPolygonItem::mouseDoubleClickEvent(__a0); }
	void cppgoBase_wheelEvent(QGraphicsSceneWheelEvent* __a0) { QGraphicsPolygonItem::wheelEvent(__a0); }
	void cppgoBase_inputMethodEvent(QInputMethodEvent* __a0) { QGraphicsPolygonItem::inputMethodEvent(__a0); }
	QVariant cppgoBase_inputMethodQuery(int __a0) { return QGraphicsPolygonItem::inputMethodQuery((Qt::InputMethodQuery)__a0); }
	QVariant cppgoBase_itemChange(int __a0, const QVariant& __a1) { return QGraphicsPolygonItem::itemChange((QGraphicsItem::GraphicsItemChange)__a0, __a1); }
	QRectF boundingRect() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 2510008782785038831LL, &__h); if (__h) return *(QRectF*)__r; }
		return QGraphicsPolygonItem::boundingRect();
	}
	QPainterPath shape() const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 7938019069247444294LL, &__h); if (__h) return *(QPainterPath*)__r; }
		return QGraphicsPolygonItem::shape();
	}
	bool contains(const QPointF& point) const override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 5645366511944179768LL, (void*)&point, &__h); if (__h) return __r; }
		return QGraphicsPolygonItem::contains(point);
	}
	void paint(QPainter* painter, const QStyleOptionGraphicsItem* option, QWidget* widget) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_voidp_voidp_voidp) { ((void(*)(long long, long long, void*, void*, void*, int*))cppgo_virt_void_voidp_voidp_voidp)(__goId, 6083584261053556819LL, (void*)painter, (void*)option, (void*)widget, &__h); if (__h) return; }
		QGraphicsPolygonItem::paint(painter, option, widget);
	}
	bool isObscuredBy(const QGraphicsItem* item) const override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, -7215030271914903671LL, (void*)item, &__h); if (__h) return __r; }
		return QGraphicsPolygonItem::isObscuredBy(item);
	}
	QPainterPath opaqueArea() const override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 399370970023989435LL, &__h); if (__h) return *(QPainterPath*)__r; }
		return QGraphicsPolygonItem::opaqueArea();
	}
	int type() const override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -6371403872082420691LL, &__h); if (__h) return __r; }
		return QGraphicsPolygonItem::type();
	}
	bool supportsExtension(QGraphicsItem::Extension extension) const override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, -5729166443651306418LL, (int)(extension), &__h); if (__h) return __r; }
		return QGraphicsPolygonItem::supportsExtension(extension);
	}
	void setExtension(QGraphicsItem::Extension extension, const QVariant& variant) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void_int_voidp) { ((void(*)(long long, long long, int, void*, int*))cppgo_virt_void_int_voidp)(__goId, -6730181576106171580LL, (int)(extension), (void*)&variant, &__h); if (__h) return; }
		QGraphicsPolygonItem::setExtension(extension, variant);
	}
	QVariant extension(const QVariant& variant) const override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 2270616358835251226LL, (void*)&variant, &__h); if (__h) return *(QVariant*)__r; }
		return QGraphicsPolygonItem::extension(variant);
	}
	void advance(int phase) override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, -4830836341480073475LL, phase, &__h); if (__h) return; }
		QGraphicsPolygonItem::advance(phase);
	}
	bool collidesWithItem(const QGraphicsItem* other, Qt::ItemSelectionMode mode) const override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_int_voidp_int) { int __r = ((int(*)(long long, long long, void*, int, int*))cppgo_virt_int_voidp_int)(__goId, 6469736184362593173LL, (void*)other, (int)(mode), &__h); if (__h) return __r; }
		return QGraphicsPolygonItem::collidesWithItem(other, mode);
	}
	bool collidesWithPath(const QPainterPath& path, Qt::ItemSelectionMode mode) const override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_int_voidp_int) { int __r = ((int(*)(long long, long long, void*, int, int*))cppgo_virt_int_voidp_int)(__goId, 6636575182718753853LL, (void*)&path, (int)(mode), &__h); if (__h) return __r; }
		return QGraphicsPolygonItem::collidesWithPath(path, mode);
	}
	bool sceneEventFilter(QGraphicsItem* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, -4584203218241215725LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QGraphicsPolygonItem::sceneEventFilter(watched, event);
	}
	bool sceneEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, -2846371507838587071LL, (void*)event, &__h); if (__h) return __r; }
		return QGraphicsPolygonItem::sceneEvent(event);
	}
	void contextMenuEvent(QGraphicsSceneContextMenuEvent* event) override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6698043222227006015LL, (void*)event, &__h); if (__h) return; }
		QGraphicsPolygonItem::contextMenuEvent(event);
	}
	void dragEnterEvent(QGraphicsSceneDragDropEvent* event) override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 206122715919917589LL, (void*)event, &__h); if (__h) return; }
		QGraphicsPolygonItem::dragEnterEvent(event);
	}
	void dragLeaveEvent(QGraphicsSceneDragDropEvent* event) override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -3139904158373725006LL, (void*)event, &__h); if (__h) return; }
		QGraphicsPolygonItem::dragLeaveEvent(event);
	}
	void dragMoveEvent(QGraphicsSceneDragDropEvent* event) override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1781545137352745978LL, (void*)event, &__h); if (__h) return; }
		QGraphicsPolygonItem::dragMoveEvent(event);
	}
	void dropEvent(QGraphicsSceneDragDropEvent* event) override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 9107058682754591716LL, (void*)event, &__h); if (__h) return; }
		QGraphicsPolygonItem::dropEvent(event);
	}
	void focusInEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 975043767033552552LL, (void*)event, &__h); if (__h) return; }
		QGraphicsPolygonItem::focusInEvent(event);
	}
	void focusOutEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -385596010576048993LL, (void*)event, &__h); if (__h) return; }
		QGraphicsPolygonItem::focusOutEvent(event);
	}
	void hoverEnterEvent(QGraphicsSceneHoverEvent* event) override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6932941943906681679LL, (void*)event, &__h); if (__h) return; }
		QGraphicsPolygonItem::hoverEnterEvent(event);
	}
	void hoverMoveEvent(QGraphicsSceneHoverEvent* event) override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -2745191522144316396LL, (void*)event, &__h); if (__h) return; }
		QGraphicsPolygonItem::hoverMoveEvent(event);
	}
	void hoverLeaveEvent(QGraphicsSceneHoverEvent* event) override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7944408482203603984LL, (void*)event, &__h); if (__h) return; }
		QGraphicsPolygonItem::hoverLeaveEvent(event);
	}
	void keyPressEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1023791977679793185LL, (void*)event, &__h); if (__h) return; }
		QGraphicsPolygonItem::keyPressEvent(event);
	}
	void keyReleaseEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 772598475950050299LL, (void*)event, &__h); if (__h) return; }
		QGraphicsPolygonItem::keyReleaseEvent(event);
	}
	void mousePressEvent(QGraphicsSceneMouseEvent* event) override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -332090898078543319LL, (void*)event, &__h); if (__h) return; }
		QGraphicsPolygonItem::mousePressEvent(event);
	}
	void mouseMoveEvent(QGraphicsSceneMouseEvent* event) override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4255494004099806581LL, (void*)event, &__h); if (__h) return; }
		QGraphicsPolygonItem::mouseMoveEvent(event);
	}
	void mouseReleaseEvent(QGraphicsSceneMouseEvent* event) override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4415920049297795213LL, (void*)event, &__h); if (__h) return; }
		QGraphicsPolygonItem::mouseReleaseEvent(event);
	}
	void mouseDoubleClickEvent(QGraphicsSceneMouseEvent* event) override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4029073905697023887LL, (void*)event, &__h); if (__h) return; }
		QGraphicsPolygonItem::mouseDoubleClickEvent(event);
	}
	void wheelEvent(QGraphicsSceneWheelEvent* event) override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1919988340934511316LL, (void*)event, &__h); if (__h) return; }
		QGraphicsPolygonItem::wheelEvent(event);
	}
	void inputMethodEvent(QInputMethodEvent* event) override {
		int __h = 0;
		if (__goOn[32] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5438819661942843222LL, (void*)event, &__h); if (__h) return; }
		QGraphicsPolygonItem::inputMethodEvent(event);
	}
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const override {
		int __h = 0;
		if (__goOn[33] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7345853097340709638LL, (int)(query), &__h); if (__h) return *(QVariant*)__r; }
		return QGraphicsPolygonItem::inputMethodQuery(query);
	}
	QVariant itemChange(QGraphicsItem::GraphicsItemChange change, const QVariant& value) override {
		int __h = 0;
		if (__goOn[34] && cppgo_virt_voidp_int_voidp) { void* __r = ((void*(*)(long long, long long, int, void*, int*))cppgo_virt_voidp_int_voidp)(__goId, -1888572168794878830LL, (int)(change), (void*)&value, &__h); if (__h) return *(QVariant*)__r; }
		return QGraphicsPolygonItem::itemChange(change, value);
	}
};

extern "C" {

void* QGraphicsPolygonItem_base(void* self) { return (void*)static_cast<QAbstractGraphicsShapeItem*>((QGraphicsPolygonItem*)self); }
void QGraphicsPolygonItem_chain(void* self, void** out) {
	QGraphicsPolygonItem* __p = (QGraphicsPolygonItem*)self;
	QAbstractGraphicsShapeItem* __b0 = static_cast<QAbstractGraphicsShapeItem*>(__p); out[0] = (void*)__b0;
	QGraphicsItem* __b1 = static_cast<QGraphicsItem*>(__b0); out[1] = (void*)__b1;
}
void QGraphicsPolygonItem__set_override(void* self, int idx) { ((QGraphicsPolygonItem__Go*)self)->__goOn[idx] = 1; }
void* QGraphicsPolygonItem_new(long long __goid) {
	return new QGraphicsPolygonItem__Go(__goid);
}
void* QGraphicsPolygonItem_new2(long long __goid, void* polygon) {
	return new QGraphicsPolygonItem__Go(__goid, *(QPolygonF*)polygon);
}
void* QGraphicsPolygonItem_new3(long long __goid, void* parent) {
	return new QGraphicsPolygonItem__Go(__goid, (QGraphicsItem*)parent);
}
void* QGraphicsPolygonItem_new4(long long __goid, void* polygon, void* parent) {
	return new QGraphicsPolygonItem__Go(__goid, *(QPolygonF*)polygon, (QGraphicsItem*)parent);
}
void* QGraphicsPolygonItem_polygon(void* self) {
	auto&& __r = ((QGraphicsPolygonItem*)self)->polygon();
	return new QPolygonF(__r);
}
void QGraphicsPolygonItem_setPolygon(void* self, void* polygon) {
	((QGraphicsPolygonItem*)self)->setPolygon(*(QPolygonF*)polygon);
}
int QGraphicsPolygonItem_fillRule(void* self) {
	auto&& __r = ((QGraphicsPolygonItem*)self)->fillRule();
	return (int)(__r);
}
void QGraphicsPolygonItem_setFillRule(void* self, int rule) {
	((QGraphicsPolygonItem*)self)->setFillRule((Qt::FillRule)(rule));
}
void QGraphicsPolygonItem_updateMicroFocus(void* self) {
	((QGraphicsPolygonItem__Go*)self)->updateMicroFocus();
}
void QGraphicsPolygonItem_addToIndex(void* self) {
	((QGraphicsPolygonItem__Go*)self)->addToIndex();
}
void QGraphicsPolygonItem_removeFromIndex(void* self) {
	((QGraphicsPolygonItem__Go*)self)->removeFromIndex();
}
void QGraphicsPolygonItem_prepareGeometryChange(void* self) {
	((QGraphicsPolygonItem__Go*)self)->prepareGeometryChange();
}
void* QGraphicsPolygonItem_boundingRect_default(void* self) {
	auto&& __r = ((QGraphicsPolygonItem*)self)->QGraphicsPolygonItem::boundingRect();
	return new QRectF(__r);
}
void* QGraphicsPolygonItem_shape_default(void* self) {
	auto&& __r = ((QGraphicsPolygonItem*)self)->QGraphicsPolygonItem::shape();
	return new QPainterPath(__r);
}
int QGraphicsPolygonItem_contains_default(void* self, void* point) {
	auto&& __r = ((QGraphicsPolygonItem*)self)->QGraphicsPolygonItem::contains(*(QPointF*)point);
	return __r;
}
void QGraphicsPolygonItem_paint_default(void* self, void* painter, void* option, void* widget) {
	((QGraphicsPolygonItem*)self)->QGraphicsPolygonItem::paint((QPainter*)painter, (QStyleOptionGraphicsItem*)option, (QWidget*)widget);
}
int QGraphicsPolygonItem_isObscuredBy_default(void* self, void* item) {
	auto&& __r = ((QGraphicsPolygonItem*)self)->QGraphicsPolygonItem::isObscuredBy((QGraphicsItem*)item);
	return __r;
}
void* QGraphicsPolygonItem_opaqueArea_default(void* self) {
	auto&& __r = ((QGraphicsPolygonItem*)self)->QGraphicsPolygonItem::opaqueArea();
	return new QPainterPath(__r);
}
int QGraphicsPolygonItem_type_default(void* self) {
	auto&& __r = ((QGraphicsPolygonItem*)self)->QGraphicsPolygonItem::type();
	return __r;
}
int QGraphicsPolygonItem_supportsExtension_default(void* self, int extension) {
	auto&& __r = ((QGraphicsPolygonItem__Go*)self)->cppgoBase_supportsExtension(extension);
	return __r;
}
void QGraphicsPolygonItem_setExtension_default(void* self, int extension, void* variant) {
	((QGraphicsPolygonItem__Go*)self)->cppgoBase_setExtension(extension, *(QVariant*)variant);
}
void* QGraphicsPolygonItem_extension_default(void* self, void* variant) {
	auto&& __r = ((QGraphicsPolygonItem__Go*)self)->cppgoBase_extension(*(QVariant*)variant);
	return new QVariant(__r);
}
void QGraphicsPolygonItem_advance_default(void* self, int phase) {
	((QGraphicsPolygonItem*)self)->QGraphicsPolygonItem::advance(phase);
}
int QGraphicsPolygonItem_collidesWithItem_default(void* self, void* other, int mode) {
	auto&& __r = ((QGraphicsPolygonItem*)self)->QGraphicsPolygonItem::collidesWithItem((QGraphicsItem*)other, (Qt::ItemSelectionMode)(mode));
	return __r;
}
int QGraphicsPolygonItem_collidesWithPath_default(void* self, void* path, int mode) {
	auto&& __r = ((QGraphicsPolygonItem*)self)->QGraphicsPolygonItem::collidesWithPath(*(QPainterPath*)path, (Qt::ItemSelectionMode)(mode));
	return __r;
}
int QGraphicsPolygonItem_sceneEventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QGraphicsPolygonItem__Go*)self)->cppgoBase_sceneEventFilter((QGraphicsItem*)watched, (QEvent*)event);
	return __r;
}
int QGraphicsPolygonItem_sceneEvent_default(void* self, void* event) {
	auto&& __r = ((QGraphicsPolygonItem__Go*)self)->cppgoBase_sceneEvent((QEvent*)event);
	return __r;
}
void QGraphicsPolygonItem_contextMenuEvent_default(void* self, void* event) {
	((QGraphicsPolygonItem__Go*)self)->cppgoBase_contextMenuEvent((QGraphicsSceneContextMenuEvent*)event);
}
void QGraphicsPolygonItem_dragEnterEvent_default(void* self, void* event) {
	((QGraphicsPolygonItem__Go*)self)->cppgoBase_dragEnterEvent((QGraphicsSceneDragDropEvent*)event);
}
void QGraphicsPolygonItem_dragLeaveEvent_default(void* self, void* event) {
	((QGraphicsPolygonItem__Go*)self)->cppgoBase_dragLeaveEvent((QGraphicsSceneDragDropEvent*)event);
}
void QGraphicsPolygonItem_dragMoveEvent_default(void* self, void* event) {
	((QGraphicsPolygonItem__Go*)self)->cppgoBase_dragMoveEvent((QGraphicsSceneDragDropEvent*)event);
}
void QGraphicsPolygonItem_dropEvent_default(void* self, void* event) {
	((QGraphicsPolygonItem__Go*)self)->cppgoBase_dropEvent((QGraphicsSceneDragDropEvent*)event);
}
void QGraphicsPolygonItem_focusInEvent_default(void* self, void* event) {
	((QGraphicsPolygonItem__Go*)self)->cppgoBase_focusInEvent((QFocusEvent*)event);
}
void QGraphicsPolygonItem_focusOutEvent_default(void* self, void* event) {
	((QGraphicsPolygonItem__Go*)self)->cppgoBase_focusOutEvent((QFocusEvent*)event);
}
void QGraphicsPolygonItem_hoverEnterEvent_default(void* self, void* event) {
	((QGraphicsPolygonItem__Go*)self)->cppgoBase_hoverEnterEvent((QGraphicsSceneHoverEvent*)event);
}
void QGraphicsPolygonItem_hoverMoveEvent_default(void* self, void* event) {
	((QGraphicsPolygonItem__Go*)self)->cppgoBase_hoverMoveEvent((QGraphicsSceneHoverEvent*)event);
}
void QGraphicsPolygonItem_hoverLeaveEvent_default(void* self, void* event) {
	((QGraphicsPolygonItem__Go*)self)->cppgoBase_hoverLeaveEvent((QGraphicsSceneHoverEvent*)event);
}
void QGraphicsPolygonItem_keyPressEvent_default(void* self, void* event) {
	((QGraphicsPolygonItem__Go*)self)->cppgoBase_keyPressEvent((QKeyEvent*)event);
}
void QGraphicsPolygonItem_keyReleaseEvent_default(void* self, void* event) {
	((QGraphicsPolygonItem__Go*)self)->cppgoBase_keyReleaseEvent((QKeyEvent*)event);
}
void QGraphicsPolygonItem_mousePressEvent_default(void* self, void* event) {
	((QGraphicsPolygonItem__Go*)self)->cppgoBase_mousePressEvent((QGraphicsSceneMouseEvent*)event);
}
void QGraphicsPolygonItem_mouseMoveEvent_default(void* self, void* event) {
	((QGraphicsPolygonItem__Go*)self)->cppgoBase_mouseMoveEvent((QGraphicsSceneMouseEvent*)event);
}
void QGraphicsPolygonItem_mouseReleaseEvent_default(void* self, void* event) {
	((QGraphicsPolygonItem__Go*)self)->cppgoBase_mouseReleaseEvent((QGraphicsSceneMouseEvent*)event);
}
void QGraphicsPolygonItem_mouseDoubleClickEvent_default(void* self, void* event) {
	((QGraphicsPolygonItem__Go*)self)->cppgoBase_mouseDoubleClickEvent((QGraphicsSceneMouseEvent*)event);
}
void QGraphicsPolygonItem_wheelEvent_default(void* self, void* event) {
	((QGraphicsPolygonItem__Go*)self)->cppgoBase_wheelEvent((QGraphicsSceneWheelEvent*)event);
}
void QGraphicsPolygonItem_inputMethodEvent_default(void* self, void* event) {
	((QGraphicsPolygonItem__Go*)self)->cppgoBase_inputMethodEvent((QInputMethodEvent*)event);
}
void* QGraphicsPolygonItem_inputMethodQuery_default(void* self, int query) {
	auto&& __r = ((QGraphicsPolygonItem__Go*)self)->cppgoBase_inputMethodQuery(query);
	return new QVariant(__r);
}
void* QGraphicsPolygonItem_itemChange_default(void* self, int change, void* value) {
	auto&& __r = ((QGraphicsPolygonItem__Go*)self)->cppgoBase_itemChange(change, *(QVariant*)value);
	return new QVariant(__r);
}
void* QGraphicsPolygonItem__Go_new(long long id) { return new QGraphicsPolygonItem__Go(id); }
void QGraphicsPolygonItem_delete(void* self) { delete (QGraphicsPolygonItem*)self; }

}

#include <QtWidgets/qgraphicssceneevent.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qpoint.h>
#include <QtCore/qnamespace.h>
#include <QtWidgets/qwidget.h>
#include "gen_qgraphicsscenecontextmenuevent.h"
#include "gen_purego_cb.h"

class QGraphicsSceneContextMenuEvent__Go : public QGraphicsSceneContextMenuEvent {
public:
	long long __goId;
	unsigned char __goOn[2] = {};
	template<typename... A> QGraphicsSceneContextMenuEvent__Go(long long id, A&&... a) : QGraphicsSceneContextMenuEvent(std::forward<A>(a)...), __goId(id) {}
	void setAccepted(bool accepted) override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 6367896784077321574LL, accepted, &__h); if (__h) return; }
		QGraphicsSceneContextMenuEvent::setAccepted(accepted);
	}
	QEvent* clone() const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, -7032813281048616478LL, &__h); if (__h) return (QEvent*)__r; }
		return QGraphicsSceneContextMenuEvent::clone();
	}
};

extern "C" {

void* QGraphicsSceneContextMenuEvent_base(void* self) { return (void*)static_cast<QGraphicsSceneEvent*>((QGraphicsSceneContextMenuEvent*)self); }
void QGraphicsSceneContextMenuEvent_chain(void* self, void** out) {
	QGraphicsSceneContextMenuEvent* __p = (QGraphicsSceneContextMenuEvent*)self;
	QGraphicsSceneEvent* __b0 = static_cast<QGraphicsSceneEvent*>(__p); out[0] = (void*)__b0;
	QEvent* __b1 = static_cast<QEvent*>(__b0); out[1] = (void*)__b1;
}
void QGraphicsSceneContextMenuEvent__set_override(void* self, int idx) { ((QGraphicsSceneContextMenuEvent__Go*)self)->__goOn[idx] = 1; }
void* QGraphicsSceneContextMenuEvent_new(long long __goid) {
	return new QGraphicsSceneContextMenuEvent__Go(__goid);
}
void* QGraphicsSceneContextMenuEvent_new2(long long __goid, int type_) {
	return new QGraphicsSceneContextMenuEvent__Go(__goid, (QEvent::Type)(type_));
}
void* QGraphicsSceneContextMenuEvent_pos(void* self) {
	auto&& __r = ((QGraphicsSceneContextMenuEvent*)self)->pos();
	return new QPointF(__r);
}
void QGraphicsSceneContextMenuEvent_setPos(void* self, void* pos) {
	((QGraphicsSceneContextMenuEvent*)self)->setPos(*(QPointF*)pos);
}
void* QGraphicsSceneContextMenuEvent_scenePos(void* self) {
	auto&& __r = ((QGraphicsSceneContextMenuEvent*)self)->scenePos();
	return new QPointF(__r);
}
void QGraphicsSceneContextMenuEvent_setScenePos(void* self, void* pos) {
	((QGraphicsSceneContextMenuEvent*)self)->setScenePos(*(QPointF*)pos);
}
void* QGraphicsSceneContextMenuEvent_screenPos(void* self) {
	auto&& __r = ((QGraphicsSceneContextMenuEvent*)self)->screenPos();
	return new QPoint(__r);
}
void QGraphicsSceneContextMenuEvent_setScreenPos(void* self, void* pos) {
	((QGraphicsSceneContextMenuEvent*)self)->setScreenPos(*(QPoint*)pos);
}
int QGraphicsSceneContextMenuEvent_modifiers(void* self) {
	auto&& __r = ((QGraphicsSceneContextMenuEvent*)self)->modifiers();
	return (int)(__r);
}
void QGraphicsSceneContextMenuEvent_setModifiers(void* self, int modifiers) {
	((QGraphicsSceneContextMenuEvent*)self)->setModifiers(QFlags<Qt::KeyboardModifier>(modifiers));
}
int QGraphicsSceneContextMenuEvent_reason(void* self) {
	auto&& __r = ((QGraphicsSceneContextMenuEvent*)self)->reason();
	return (int)(__r);
}
void QGraphicsSceneContextMenuEvent_setReason(void* self, int reason) {
	((QGraphicsSceneContextMenuEvent*)self)->setReason((QGraphicsSceneContextMenuEvent::Reason)(reason));
}
int QGraphicsSceneContextMenuEvent_registerEventType() {
	auto&& __r = QGraphicsSceneContextMenuEvent::registerEventType();
	return __r;
}
int QGraphicsSceneContextMenuEvent_registerEventType2(int hint) {
	auto&& __r = QGraphicsSceneContextMenuEvent::registerEventType(hint);
	return __r;
}
void QGraphicsSceneContextMenuEvent_setAccepted_default(void* self, int accepted) {
	((QGraphicsSceneContextMenuEvent*)self)->QGraphicsSceneContextMenuEvent::setAccepted(accepted);
}
void* QGraphicsSceneContextMenuEvent_clone_default(void* self) {
	auto&& __r = ((QGraphicsSceneContextMenuEvent*)self)->QGraphicsSceneContextMenuEvent::clone();
	return (void*)(__r);
}
void* QGraphicsSceneContextMenuEvent__Go_new(long long id) { return new QGraphicsSceneContextMenuEvent__Go(id); }
void QGraphicsSceneContextMenuEvent_delete(void* self) { delete (QGraphicsSceneContextMenuEvent*)self; }

}

#include <QtWidgets/qmainwindow.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qsize.h>
#include <QtWidgets/qtabwidget.h>
#include <QtCore/qpoint.h>
#include <QtWidgets/qmenubar.h>
#include <QtWidgets/qstatusbar.h>
#include <QtWidgets/qtoolbar.h>
#include <QtWidgets/qdockwidget.h>
#include <QtCore/qbytearray.h>
#include <QtWidgets/qmenu.h>
#include <QtGui/qevent.h>
#include <QtCore/qcoreevent.h>
#include <QtWidgets/qstyle.h>
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
#include "gen_qmainwindow.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QMainWindow__Go : public QMainWindow {
public:
	long long __goId;
	unsigned char __goOn[49] = {};
	template<typename... A> QMainWindow__Go(long long id, A&&... a) : QMainWindow(std::forward<A>(a)...), __goId(id) {}
	using QMainWindow::updateMicroFocus;
	using QMainWindow::create;
	using QMainWindow::destroy;
	using QMainWindow::focusNextChild;
	using QMainWindow::focusPreviousChild;
	using QMainWindow::sender;
	using QMainWindow::senderSignalIndex;
	using QMainWindow::receivers;
	using QMainWindow::isSignalConnected;
	using QMainWindow::getDecodedMetricF;
	void cppgoBase_contextMenuEvent(QContextMenuEvent* __a0) { QMainWindow::contextMenuEvent(__a0); }
	bool cppgoBase_event(QEvent* __a0) { return QMainWindow::event(__a0); }
	void cppgoBase_mousePressEvent(QMouseEvent* __a0) { QMainWindow::mousePressEvent(__a0); }
	void cppgoBase_mouseReleaseEvent(QMouseEvent* __a0) { QMainWindow::mouseReleaseEvent(__a0); }
	void cppgoBase_mouseDoubleClickEvent(QMouseEvent* __a0) { QMainWindow::mouseDoubleClickEvent(__a0); }
	void cppgoBase_mouseMoveEvent(QMouseEvent* __a0) { QMainWindow::mouseMoveEvent(__a0); }
	void cppgoBase_wheelEvent(QWheelEvent* __a0) { QMainWindow::wheelEvent(__a0); }
	void cppgoBase_keyPressEvent(QKeyEvent* __a0) { QMainWindow::keyPressEvent(__a0); }
	void cppgoBase_keyReleaseEvent(QKeyEvent* __a0) { QMainWindow::keyReleaseEvent(__a0); }
	void cppgoBase_focusInEvent(QFocusEvent* __a0) { QMainWindow::focusInEvent(__a0); }
	void cppgoBase_focusOutEvent(QFocusEvent* __a0) { QMainWindow::focusOutEvent(__a0); }
	void cppgoBase_enterEvent(QEnterEvent* __a0) { QMainWindow::enterEvent(__a0); }
	void cppgoBase_leaveEvent(QEvent* __a0) { QMainWindow::leaveEvent(__a0); }
	void cppgoBase_paintEvent(QPaintEvent* __a0) { QMainWindow::paintEvent(__a0); }
	void cppgoBase_moveEvent(QMoveEvent* __a0) { QMainWindow::moveEvent(__a0); }
	void cppgoBase_resizeEvent(QResizeEvent* __a0) { QMainWindow::resizeEvent(__a0); }
	void cppgoBase_closeEvent(QCloseEvent* __a0) { QMainWindow::closeEvent(__a0); }
	void cppgoBase_tabletEvent(QTabletEvent* __a0) { QMainWindow::tabletEvent(__a0); }
	void cppgoBase_actionEvent(QActionEvent* __a0) { QMainWindow::actionEvent(__a0); }
	void cppgoBase_dragEnterEvent(QDragEnterEvent* __a0) { QMainWindow::dragEnterEvent(__a0); }
	void cppgoBase_dragMoveEvent(QDragMoveEvent* __a0) { QMainWindow::dragMoveEvent(__a0); }
	void cppgoBase_dragLeaveEvent(QDragLeaveEvent* __a0) { QMainWindow::dragLeaveEvent(__a0); }
	void cppgoBase_dropEvent(QDropEvent* __a0) { QMainWindow::dropEvent(__a0); }
	void cppgoBase_showEvent(QShowEvent* __a0) { QMainWindow::showEvent(__a0); }
	void cppgoBase_hideEvent(QHideEvent* __a0) { QMainWindow::hideEvent(__a0); }
	void cppgoBase_changeEvent(QEvent* __a0) { QMainWindow::changeEvent(__a0); }
	int cppgoBase_metric(int __a0) { return QMainWindow::metric((QPaintDevice::PaintDeviceMetric)__a0); }
	void cppgoBase_initPainter(QPainter* __a0) { QMainWindow::initPainter(__a0); }
	QPaintDevice* cppgoBase_redirected(QPoint* __a0) { return QMainWindow::redirected(__a0); }
	QPainter* cppgoBase_sharedPainter() { return QMainWindow::sharedPainter(); }
	void cppgoBase_inputMethodEvent(QInputMethodEvent* __a0) { QMainWindow::inputMethodEvent(__a0); }
	bool cppgoBase_focusNextPrevChild(bool __a0) { return QMainWindow::focusNextPrevChild(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QMainWindow::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QMainWindow::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QMainWindow::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QMainWindow::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QMainWindow::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QMainWindow::metaObject();
	}
	QMenu* createPopupMenu() override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1078400071543754474LL, &__h); if (__h) return (QMenu*)__r; }
		return QMainWindow::createPopupMenu();
	}
	void contextMenuEvent(QContextMenuEvent* event) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6698043222227006015LL, (void*)event, &__h); if (__h) return; }
		QMainWindow::contextMenuEvent(event);
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QMainWindow::event(event);
	}
	int devType() const override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8659679893868524044LL, &__h); if (__h) return __r; }
		return QMainWindow::devType();
	}
	void setVisible(bool visible) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 2110182719375245575LL, visible, &__h); if (__h) return; }
		QMainWindow::setVisible(visible);
	}
	QSize sizeHint() const override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 729690261457684101LL, &__h); if (__h) return *(QSize*)__r; }
		return QMainWindow::sizeHint();
	}
	QSize minimumSizeHint() const override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 7049264599837532491LL, &__h); if (__h) return *(QSize*)__r; }
		return QMainWindow::minimumSizeHint();
	}
	int heightForWidth(int arg1) const override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 4311534060495737339LL, arg1, &__h); if (__h) return __r; }
		return QMainWindow::heightForWidth(arg1);
	}
	bool hasHeightForWidth() const override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8266762397995423915LL, &__h); if (__h) return __r; }
		return QMainWindow::hasHeightForWidth();
	}
	QPaintEngine* paintEngine() const override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 6398007404756346381LL, &__h); if (__h) return (QPaintEngine*)__r; }
		return QMainWindow::paintEngine();
	}
	void mousePressEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -332090898078543319LL, (void*)event, &__h); if (__h) return; }
		QMainWindow::mousePressEvent(event);
	}
	void mouseReleaseEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4415920049297795213LL, (void*)event, &__h); if (__h) return; }
		QMainWindow::mouseReleaseEvent(event);
	}
	void mouseDoubleClickEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4029073905697023887LL, (void*)event, &__h); if (__h) return; }
		QMainWindow::mouseDoubleClickEvent(event);
	}
	void mouseMoveEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4255494004099806581LL, (void*)event, &__h); if (__h) return; }
		QMainWindow::mouseMoveEvent(event);
	}
	void wheelEvent(QWheelEvent* event) override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1919988340934511316LL, (void*)event, &__h); if (__h) return; }
		QMainWindow::wheelEvent(event);
	}
	void keyPressEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1023791977679793185LL, (void*)event, &__h); if (__h) return; }
		QMainWindow::keyPressEvent(event);
	}
	void keyReleaseEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 772598475950050299LL, (void*)event, &__h); if (__h) return; }
		QMainWindow::keyReleaseEvent(event);
	}
	void focusInEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 975043767033552552LL, (void*)event, &__h); if (__h) return; }
		QMainWindow::focusInEvent(event);
	}
	void focusOutEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -385596010576048993LL, (void*)event, &__h); if (__h) return; }
		QMainWindow::focusOutEvent(event);
	}
	void enterEvent(QEnterEvent* event) override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6316270346668237895LL, (void*)event, &__h); if (__h) return; }
		QMainWindow::enterEvent(event);
	}
	void leaveEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -551070009623711640LL, (void*)event, &__h); if (__h) return; }
		QMainWindow::leaveEvent(event);
	}
	void paintEvent(QPaintEvent* event) override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5290442163167458465LL, (void*)event, &__h); if (__h) return; }
		QMainWindow::paintEvent(event);
	}
	void moveEvent(QMoveEvent* event) override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -9161949150489525396LL, (void*)event, &__h); if (__h) return; }
		QMainWindow::moveEvent(event);
	}
	void resizeEvent(QResizeEvent* event) override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4302917285655014661LL, (void*)event, &__h); if (__h) return; }
		QMainWindow::resizeEvent(event);
	}
	void closeEvent(QCloseEvent* event) override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1989937050019464783LL, (void*)event, &__h); if (__h) return; }
		QMainWindow::closeEvent(event);
	}
	void tabletEvent(QTabletEvent* event) override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -212068866860701141LL, (void*)event, &__h); if (__h) return; }
		QMainWindow::tabletEvent(event);
	}
	void actionEvent(QActionEvent* event) override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1697444140088555581LL, (void*)event, &__h); if (__h) return; }
		QMainWindow::actionEvent(event);
	}
	void dragEnterEvent(QDragEnterEvent* event) override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 206122715919917589LL, (void*)event, &__h); if (__h) return; }
		QMainWindow::dragEnterEvent(event);
	}
	void dragMoveEvent(QDragMoveEvent* event) override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1781545137352745978LL, (void*)event, &__h); if (__h) return; }
		QMainWindow::dragMoveEvent(event);
	}
	void dragLeaveEvent(QDragLeaveEvent* event) override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -3139904158373725006LL, (void*)event, &__h); if (__h) return; }
		QMainWindow::dragLeaveEvent(event);
	}
	void dropEvent(QDropEvent* event) override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 9107058682754591716LL, (void*)event, &__h); if (__h) return; }
		QMainWindow::dropEvent(event);
	}
	void showEvent(QShowEvent* event) override {
		int __h = 0;
		if (__goOn[32] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6825369174751916LL, (void*)event, &__h); if (__h) return; }
		QMainWindow::showEvent(event);
	}
	void hideEvent(QHideEvent* event) override {
		int __h = 0;
		if (__goOn[33] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -199718707876719129LL, (void*)event, &__h); if (__h) return; }
		QMainWindow::hideEvent(event);
	}
	bool nativeEvent(const QByteArray& eventType, void* message, qintptr* result) override {
		int __h = 0;
		if (__goOn[34] && cppgo_virt_int_voidp_voidp_longlongp) { int __r = ((int(*)(long long, long long, void*, void*, long long*, int*))cppgo_virt_int_voidp_voidp_longlongp)(__goId, -2039166586246443706LL, (void*)&eventType, (void*)message, (long long*)result, &__h); if (__h) return __r; }
		return QMainWindow::nativeEvent(eventType, message, result);
	}
	void changeEvent(QEvent* arg1) override {
		int __h = 0;
		if (__goOn[35] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4959376063826165943LL, (void*)arg1, &__h); if (__h) return; }
		QMainWindow::changeEvent(arg1);
	}
	int metric(QPaintDevice::PaintDeviceMetric arg1) const override {
		int __h = 0;
		if (__goOn[36] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 427629649267311777LL, (int)(arg1), &__h); if (__h) return __r; }
		return QMainWindow::metric(arg1);
	}
	void initPainter(QPainter* painter) const override {
		int __h = 0;
		if (__goOn[37] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6477500296980050210LL, (void*)painter, &__h); if (__h) return; }
		QMainWindow::initPainter(painter);
	}
	QPaintDevice* redirected(QPoint* offset) const override {
		int __h = 0;
		if (__goOn[38] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 4127240633310655048LL, (void*)offset, &__h); if (__h) return (QPaintDevice*)__r; }
		return QMainWindow::redirected(offset);
	}
	QPainter* sharedPainter() const override {
		int __h = 0;
		if (__goOn[39] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 5792771328741774991LL, &__h); if (__h) return (QPainter*)__r; }
		return QMainWindow::sharedPainter();
	}
	void inputMethodEvent(QInputMethodEvent* arg1) override {
		int __h = 0;
		if (__goOn[40] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5438819661942843222LL, (void*)arg1, &__h); if (__h) return; }
		QMainWindow::inputMethodEvent(arg1);
	}
	QVariant inputMethodQuery(Qt::InputMethodQuery arg1) const override {
		int __h = 0;
		if (__goOn[41] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7345853097340709638LL, (int)(arg1), &__h); if (__h) return *(QVariant*)__r; }
		return QMainWindow::inputMethodQuery(arg1);
	}
	bool focusNextPrevChild(bool next) override {
		int __h = 0;
		if (__goOn[42] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 5898328517505784191LL, next, &__h); if (__h) return __r; }
		return QMainWindow::focusNextPrevChild(next);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[43] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QMainWindow::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[44] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QMainWindow::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[45] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QMainWindow::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[46] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QMainWindow::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[47] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QMainWindow::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[48] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QMainWindow::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QMainWindow;
static std::mutex __connsMu_QMainWindow;

extern "C" {

void* QMainWindow_base(void* self) { return (void*)static_cast<QWidget*>((QMainWindow*)self); }
void QMainWindow_chain(void* self, void** out) {
	QMainWindow* __p = (QMainWindow*)self;
	QWidget* __b0 = static_cast<QWidget*>(__p); out[0] = (void*)__b0;
	QObject* __b1 = static_cast<QObject*>(__b0); out[1] = (void*)__b1;
}
void QMainWindow__set_override(void* self, int idx) { ((QMainWindow__Go*)self)->__goOn[idx] = 1; }
void* QMainWindow_fromQObject(void* o) { return (void*)qobject_cast<QMainWindow*>((QObject*)o); }
void* QMainWindow_new(long long __goid, void* parent) {
	return new QMainWindow__Go(__goid, (QWidget*)parent);
}
void* QMainWindow_new2(long long __goid) {
	return new QMainWindow__Go(__goid);
}
void* QMainWindow_new3(long long __goid, void* parent, int flags) {
	return new QMainWindow__Go(__goid, (QWidget*)parent, QFlags<Qt::WindowType>(flags));
}
char* QMainWindow_tr(const char* s, long long* __out_len) {
	auto&& __r = QMainWindow::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QMainWindow_iconSize(void* self) {
	auto&& __r = ((QMainWindow*)self)->iconSize();
	return new QSize(__r);
}
void QMainWindow_setIconSize(void* self, void* iconSize) {
	((QMainWindow*)self)->setIconSize(*(QSize*)iconSize);
}
int QMainWindow_toolButtonStyle(void* self) {
	auto&& __r = ((QMainWindow*)self)->toolButtonStyle();
	return (int)(__r);
}
void QMainWindow_setToolButtonStyle(void* self, int toolButtonStyle) {
	((QMainWindow*)self)->setToolButtonStyle((Qt::ToolButtonStyle)(toolButtonStyle));
}
int QMainWindow_isAnimated(void* self) {
	auto&& __r = ((QMainWindow*)self)->isAnimated();
	return __r;
}
int QMainWindow_isDockNestingEnabled(void* self) {
	auto&& __r = ((QMainWindow*)self)->isDockNestingEnabled();
	return __r;
}
int QMainWindow_documentMode(void* self) {
	auto&& __r = ((QMainWindow*)self)->documentMode();
	return __r;
}
void QMainWindow_setDocumentMode(void* self, int enabled) {
	((QMainWindow*)self)->setDocumentMode(enabled);
}
int QMainWindow_tabShape(void* self) {
	auto&& __r = ((QMainWindow*)self)->tabShape();
	return (int)(__r);
}
void QMainWindow_setTabShape(void* self, int tabShape) {
	((QMainWindow*)self)->setTabShape((QTabWidget::TabShape)(tabShape));
}
int QMainWindow_tabPosition(void* self, int area) {
	auto&& __r = ((QMainWindow*)self)->tabPosition((Qt::DockWidgetArea)(area));
	return (int)(__r);
}
void QMainWindow_setTabPosition(void* self, int areas, int tabPosition) {
	((QMainWindow*)self)->setTabPosition(QFlags<Qt::DockWidgetArea>(areas), (QTabWidget::TabPosition)(tabPosition));
}
void QMainWindow_setDockOptions(void* self, int options) {
	((QMainWindow*)self)->setDockOptions(QFlags<QMainWindow::DockOption>(options));
}
int QMainWindow_dockOptions(void* self) {
	auto&& __r = ((QMainWindow*)self)->dockOptions();
	return (int)(__r);
}
int QMainWindow_isSeparator(void* self, void* pos) {
	auto&& __r = ((QMainWindow*)self)->isSeparator(*(QPoint*)pos);
	return __r;
}
void* QMainWindow_menuBar(void* self) {
	auto&& __r = ((QMainWindow*)self)->menuBar();
	return (void*)(__r);
}
void QMainWindow_setMenuBar(void* self, void* menubar) {
	((QMainWindow*)self)->setMenuBar((QMenuBar*)menubar);
}
void* QMainWindow_menuWidget(void* self) {
	auto&& __r = ((QMainWindow*)self)->menuWidget();
	return (void*)(__r);
}
void QMainWindow_setMenuWidget(void* self, void* menubar) {
	((QMainWindow*)self)->setMenuWidget((QWidget*)menubar);
}
void* QMainWindow_statusBar(void* self) {
	auto&& __r = ((QMainWindow*)self)->statusBar();
	return (void*)(__r);
}
void QMainWindow_setStatusBar(void* self, void* statusbar) {
	((QMainWindow*)self)->setStatusBar((QStatusBar*)statusbar);
}
void* QMainWindow_centralWidget(void* self) {
	auto&& __r = ((QMainWindow*)self)->centralWidget();
	return (void*)(__r);
}
void QMainWindow_setCentralWidget(void* self, void* widget) {
	((QMainWindow*)self)->setCentralWidget((QWidget*)widget);
}
void* QMainWindow_takeCentralWidget(void* self) {
	auto&& __r = ((QMainWindow*)self)->takeCentralWidget();
	return (void*)(__r);
}
void QMainWindow_setCorner(void* self, int corner, int area) {
	((QMainWindow*)self)->setCorner((Qt::Corner)(corner), (Qt::DockWidgetArea)(area));
}
int QMainWindow_corner(void* self, int corner) {
	auto&& __r = ((QMainWindow*)self)->corner((Qt::Corner)(corner));
	return (int)(__r);
}
void QMainWindow_addToolBarBreak(void* self) {
	((QMainWindow*)self)->addToolBarBreak();
}
void QMainWindow_insertToolBarBreak(void* self, void* before) {
	((QMainWindow*)self)->insertToolBarBreak((QToolBar*)before);
}
void QMainWindow_addToolBar(void* self, int area, void* toolbar) {
	((QMainWindow*)self)->addToolBar((Qt::ToolBarArea)(area), (QToolBar*)toolbar);
}
void QMainWindow_addToolBar2(void* self, void* toolbar) {
	((QMainWindow*)self)->addToolBar((QToolBar*)toolbar);
}
void* QMainWindow_addToolBar3(void* self, const char* title, long long title_len) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	auto&& __r = ((QMainWindow*)self)->addToolBar(title_qstr);
	return (void*)(__r);
}
void QMainWindow_insertToolBar(void* self, void* before, void* toolbar) {
	((QMainWindow*)self)->insertToolBar((QToolBar*)before, (QToolBar*)toolbar);
}
void QMainWindow_removeToolBar(void* self, void* toolbar) {
	((QMainWindow*)self)->removeToolBar((QToolBar*)toolbar);
}
void QMainWindow_removeToolBarBreak(void* self, void* before) {
	((QMainWindow*)self)->removeToolBarBreak((QToolBar*)before);
}
int QMainWindow_unifiedTitleAndToolBarOnMac(void* self) {
	auto&& __r = ((QMainWindow*)self)->unifiedTitleAndToolBarOnMac();
	return __r;
}
int QMainWindow_toolBarArea(void* self, void* toolbar) {
	auto&& __r = ((QMainWindow*)self)->toolBarArea((QToolBar*)toolbar);
	return (int)(__r);
}
int QMainWindow_toolBarBreak(void* self, void* toolbar) {
	auto&& __r = ((QMainWindow*)self)->toolBarBreak((QToolBar*)toolbar);
	return __r;
}
void QMainWindow_addDockWidget(void* self, int area, void* dockwidget) {
	((QMainWindow*)self)->addDockWidget((Qt::DockWidgetArea)(area), (QDockWidget*)dockwidget);
}
void QMainWindow_addDockWidget2(void* self, int area, void* dockwidget, int orientation) {
	((QMainWindow*)self)->addDockWidget((Qt::DockWidgetArea)(area), (QDockWidget*)dockwidget, (Qt::Orientation)(orientation));
}
void QMainWindow_splitDockWidget(void* self, void* after, void* dockwidget, int orientation) {
	((QMainWindow*)self)->splitDockWidget((QDockWidget*)after, (QDockWidget*)dockwidget, (Qt::Orientation)(orientation));
}
void QMainWindow_tabifyDockWidget(void* self, void* first, void* second) {
	((QMainWindow*)self)->tabifyDockWidget((QDockWidget*)first, (QDockWidget*)second);
}
void** QMainWindow_tabifiedDockWidgets(void* self, void* dockwidget, long long* __out_len) {
	auto&& __r = ((QMainWindow*)self)->tabifiedDockWidgets((QDockWidget*)dockwidget);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
void QMainWindow_removeDockWidget(void* self, void* dockwidget) {
	((QMainWindow*)self)->removeDockWidget((QDockWidget*)dockwidget);
}
int QMainWindow_restoreDockWidget(void* self, void* dockwidget) {
	auto&& __r = ((QMainWindow*)self)->restoreDockWidget((QDockWidget*)dockwidget);
	return __r;
}
int QMainWindow_dockWidgetArea(void* self, void* dockwidget) {
	auto&& __r = ((QMainWindow*)self)->dockWidgetArea((QDockWidget*)dockwidget);
	return (int)(__r);
}
void QMainWindow_resizeDocks(void* self, void** docks, long long docks_len, int* sizes, long long sizes_len, int orientation) {
	QList<QDockWidget*> docks_list; docks_list.reserve(docks_len);
	for (long long __i = 0; __i < docks_len; __i++) docks_list.append((QDockWidget*)docks[__i]);
	QList<int> sizes_list; sizes_list.reserve(sizes_len);
	for (long long __i = 0; __i < sizes_len; __i++) sizes_list.append(sizes[__i]);
	((QMainWindow*)self)->resizeDocks(docks_list, sizes_list, (Qt::Orientation)(orientation));
}
void* QMainWindow_saveState(void* self) {
	auto&& __r = ((QMainWindow*)self)->saveState();
	return new QByteArray(__r);
}
int QMainWindow_restoreState(void* self, void* state) {
	auto&& __r = ((QMainWindow*)self)->restoreState(*(QByteArray*)state);
	return __r;
}
void* QMainWindow_createPopupMenu(void* self) {
	auto&& __r = ((QMainWindow*)self)->createPopupMenu();
	return (void*)(__r);
}
void QMainWindow_setAnimated(void* self, int enabled) {
	((QMainWindow*)self)->setAnimated(enabled);
}
void QMainWindow_setDockNestingEnabled(void* self, int enabled) {
	((QMainWindow*)self)->setDockNestingEnabled(enabled);
}
void QMainWindow_setUnifiedTitleAndToolBarOnMac(void* self, int set) {
	((QMainWindow*)self)->setUnifiedTitleAndToolBarOnMac(set);
}
void QMainWindow_iconSizeChanged(void* self, void* iconSize) {
	((QMainWindow*)self)->iconSizeChanged(*(QSize*)iconSize);
}
void QMainWindow_toolButtonStyleChanged(void* self, int toolButtonStyle) {
	((QMainWindow*)self)->toolButtonStyleChanged((Qt::ToolButtonStyle)(toolButtonStyle));
}
void QMainWindow_tabifiedDockWidgetActivated(void* self, void* dockWidget) {
	((QMainWindow*)self)->tabifiedDockWidgetActivated((QDockWidget*)dockWidget);
}
void QMainWindow_setTabOrder(void* arg1, void* arg2) {
	QMainWindow::setTabOrder((QWidget*)arg1, (QWidget*)arg2);
}
void* QMainWindow_mouseGrabber() {
	auto&& __r = QMainWindow::mouseGrabber();
	return (void*)(__r);
}
void* QMainWindow_keyboardGrabber() {
	auto&& __r = QMainWindow::keyboardGrabber();
	return (void*)(__r);
}
void* QMainWindow_find(unsigned long long arg1) {
	auto&& __r = QMainWindow::find(arg1);
	return (void*)(__r);
}
void* QMainWindow_createWindowContainer(void* window) {
	auto&& __r = QMainWindow::createWindowContainer((QWindow*)window);
	return (void*)(__r);
}
void QMainWindow_updateMicroFocus(void* self) {
	((QMainWindow__Go*)self)->updateMicroFocus();
}
void QMainWindow_create(void* self) {
	((QMainWindow__Go*)self)->create();
}
void QMainWindow_destroy(void* self) {
	((QMainWindow__Go*)self)->destroy();
}
int QMainWindow_focusNextChild(void* self) {
	auto&& __r = ((QMainWindow__Go*)self)->focusNextChild();
	return __r;
}
int QMainWindow_focusPreviousChild(void* self) {
	auto&& __r = ((QMainWindow__Go*)self)->focusPreviousChild();
	return __r;
}
void* QMainWindow_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QMainWindow::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QMainWindow_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QMainWindow::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QMainWindow_sender(void* self) {
	auto&& __r = ((QMainWindow__Go*)self)->sender();
	return (void*)(__r);
}
int QMainWindow_senderSignalIndex(void* self) {
	auto&& __r = ((QMainWindow__Go*)self)->senderSignalIndex();
	return __r;
}
int QMainWindow_receivers(void* self, const char* signal) {
	auto&& __r = ((QMainWindow__Go*)self)->receivers(signal);
	return __r;
}
int QMainWindow_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QMainWindow__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
double QMainWindow_devicePixelRatioFScale() {
	auto&& __r = QMainWindow::devicePixelRatioFScale();
	return __r;
}
int QMainWindow_encodeMetricF(int metric, double value) {
	auto&& __r = QMainWindow::encodeMetricF((QPaintDevice::PaintDeviceMetric)(metric), value);
	return __r;
}
double QMainWindow_getDecodedMetricF(void* self, int metricA, int metricB) {
	auto&& __r = ((QMainWindow__Go*)self)->getDecodedMetricF((QPaintDevice::PaintDeviceMetric)(metricA), (QPaintDevice::PaintDeviceMetric)(metricB));
	return __r;
}
char* QMainWindow_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QMainWindow::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QMainWindow_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QMainWindow::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QMainWindow_addToolBarBreak2(void* self, int area) {
	((QMainWindow*)self)->addToolBarBreak((Qt::ToolBarArea)(area));
}
void* QMainWindow_saveState2(void* self, int version) {
	auto&& __r = ((QMainWindow*)self)->saveState(version);
	return new QByteArray(__r);
}
int QMainWindow_restoreState2(void* self, void* state, int version) {
	auto&& __r = ((QMainWindow*)self)->restoreState(*(QByteArray*)state, version);
	return __r;
}
void* QMainWindow_createWindowContainer2(void* window, void* parent) {
	auto&& __r = QMainWindow::createWindowContainer((QWindow*)window, (QWidget*)parent);
	return (void*)(__r);
}
void* QMainWindow_createWindowContainer3(void* window, void* parent, int flags) {
	auto&& __r = QMainWindow::createWindowContainer((QWindow*)window, (QWidget*)parent, QFlags<Qt::WindowType>(flags));
	return (void*)(__r);
}
void QMainWindow_updateMicroFocus2(void* self, int query) {
	((QMainWindow__Go*)self)->updateMicroFocus((Qt::InputMethodQuery)(query));
}
void QMainWindow_create2(void* self, unsigned long long arg1) {
	((QMainWindow__Go*)self)->create(arg1);
}
void QMainWindow_create3(void* self, unsigned long long arg1, int initializeWindow) {
	((QMainWindow__Go*)self)->create(arg1, initializeWindow);
}
void QMainWindow_create4(void* self, unsigned long long arg1, int initializeWindow, int destroyOldWindow) {
	((QMainWindow__Go*)self)->create(arg1, initializeWindow, destroyOldWindow);
}
void QMainWindow_destroy2(void* self, int destroyWindow) {
	((QMainWindow__Go*)self)->destroy(destroyWindow);
}
void QMainWindow_destroy3(void* self, int destroyWindow, int destroySubWindows) {
	((QMainWindow__Go*)self)->destroy(destroyWindow, destroySubWindows);
}
void* QMainWindow_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QMainWindow::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QMainWindow_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QMainWindow::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QMainWindow_createPopupMenu_default(void* self) {
	auto&& __r = ((QMainWindow*)self)->QMainWindow::createPopupMenu();
	return (void*)(__r);
}
void* QMainWindow_metaObject_default(void* self) {
	auto&& __r = ((QMainWindow*)self)->QMainWindow::metaObject();
	return (void*)(__r);
}
void QMainWindow_contextMenuEvent_default(void* self, void* event) {
	((QMainWindow__Go*)self)->cppgoBase_contextMenuEvent((QContextMenuEvent*)event);
}
int QMainWindow_event_default(void* self, void* event) {
	auto&& __r = ((QMainWindow__Go*)self)->cppgoBase_event((QEvent*)event);
	return __r;
}
int QMainWindow_devType_default(void* self) {
	auto&& __r = ((QMainWindow*)self)->QMainWindow::devType();
	return __r;
}
void QMainWindow_setVisible_default(void* self, int visible) {
	((QMainWindow*)self)->QMainWindow::setVisible(visible);
}
void* QMainWindow_sizeHint_default(void* self) {
	auto&& __r = ((QMainWindow*)self)->QMainWindow::sizeHint();
	return new QSize(__r);
}
void* QMainWindow_minimumSizeHint_default(void* self) {
	auto&& __r = ((QMainWindow*)self)->QMainWindow::minimumSizeHint();
	return new QSize(__r);
}
int QMainWindow_heightForWidth_default(void* self, int arg1) {
	auto&& __r = ((QMainWindow*)self)->QMainWindow::heightForWidth(arg1);
	return __r;
}
int QMainWindow_hasHeightForWidth_default(void* self) {
	auto&& __r = ((QMainWindow*)self)->QMainWindow::hasHeightForWidth();
	return __r;
}
void* QMainWindow_paintEngine_default(void* self) {
	auto&& __r = ((QMainWindow*)self)->QMainWindow::paintEngine();
	return (void*)(__r);
}
void QMainWindow_mousePressEvent_default(void* self, void* event) {
	((QMainWindow__Go*)self)->cppgoBase_mousePressEvent((QMouseEvent*)event);
}
void QMainWindow_mouseReleaseEvent_default(void* self, void* event) {
	((QMainWindow__Go*)self)->cppgoBase_mouseReleaseEvent((QMouseEvent*)event);
}
void QMainWindow_mouseDoubleClickEvent_default(void* self, void* event) {
	((QMainWindow__Go*)self)->cppgoBase_mouseDoubleClickEvent((QMouseEvent*)event);
}
void QMainWindow_mouseMoveEvent_default(void* self, void* event) {
	((QMainWindow__Go*)self)->cppgoBase_mouseMoveEvent((QMouseEvent*)event);
}
void QMainWindow_wheelEvent_default(void* self, void* event) {
	((QMainWindow__Go*)self)->cppgoBase_wheelEvent((QWheelEvent*)event);
}
void QMainWindow_keyPressEvent_default(void* self, void* event) {
	((QMainWindow__Go*)self)->cppgoBase_keyPressEvent((QKeyEvent*)event);
}
void QMainWindow_keyReleaseEvent_default(void* self, void* event) {
	((QMainWindow__Go*)self)->cppgoBase_keyReleaseEvent((QKeyEvent*)event);
}
void QMainWindow_focusInEvent_default(void* self, void* event) {
	((QMainWindow__Go*)self)->cppgoBase_focusInEvent((QFocusEvent*)event);
}
void QMainWindow_focusOutEvent_default(void* self, void* event) {
	((QMainWindow__Go*)self)->cppgoBase_focusOutEvent((QFocusEvent*)event);
}
void QMainWindow_enterEvent_default(void* self, void* event) {
	((QMainWindow__Go*)self)->cppgoBase_enterEvent((QEnterEvent*)event);
}
void QMainWindow_leaveEvent_default(void* self, void* event) {
	((QMainWindow__Go*)self)->cppgoBase_leaveEvent((QEvent*)event);
}
void QMainWindow_paintEvent_default(void* self, void* event) {
	((QMainWindow__Go*)self)->cppgoBase_paintEvent((QPaintEvent*)event);
}
void QMainWindow_moveEvent_default(void* self, void* event) {
	((QMainWindow__Go*)self)->cppgoBase_moveEvent((QMoveEvent*)event);
}
void QMainWindow_resizeEvent_default(void* self, void* event) {
	((QMainWindow__Go*)self)->cppgoBase_resizeEvent((QResizeEvent*)event);
}
void QMainWindow_closeEvent_default(void* self, void* event) {
	((QMainWindow__Go*)self)->cppgoBase_closeEvent((QCloseEvent*)event);
}
void QMainWindow_tabletEvent_default(void* self, void* event) {
	((QMainWindow__Go*)self)->cppgoBase_tabletEvent((QTabletEvent*)event);
}
void QMainWindow_actionEvent_default(void* self, void* event) {
	((QMainWindow__Go*)self)->cppgoBase_actionEvent((QActionEvent*)event);
}
void QMainWindow_dragEnterEvent_default(void* self, void* event) {
	((QMainWindow__Go*)self)->cppgoBase_dragEnterEvent((QDragEnterEvent*)event);
}
void QMainWindow_dragMoveEvent_default(void* self, void* event) {
	((QMainWindow__Go*)self)->cppgoBase_dragMoveEvent((QDragMoveEvent*)event);
}
void QMainWindow_dragLeaveEvent_default(void* self, void* event) {
	((QMainWindow__Go*)self)->cppgoBase_dragLeaveEvent((QDragLeaveEvent*)event);
}
void QMainWindow_dropEvent_default(void* self, void* event) {
	((QMainWindow__Go*)self)->cppgoBase_dropEvent((QDropEvent*)event);
}
void QMainWindow_showEvent_default(void* self, void* event) {
	((QMainWindow__Go*)self)->cppgoBase_showEvent((QShowEvent*)event);
}
void QMainWindow_hideEvent_default(void* self, void* event) {
	((QMainWindow__Go*)self)->cppgoBase_hideEvent((QHideEvent*)event);
}
void QMainWindow_changeEvent_default(void* self, void* arg1) {
	((QMainWindow__Go*)self)->cppgoBase_changeEvent((QEvent*)arg1);
}
int QMainWindow_metric_default(void* self, int arg1) {
	auto&& __r = ((QMainWindow__Go*)self)->cppgoBase_metric(arg1);
	return __r;
}
void QMainWindow_initPainter_default(void* self, void* painter) {
	((QMainWindow__Go*)self)->cppgoBase_initPainter((QPainter*)painter);
}
void* QMainWindow_redirected_default(void* self, void* offset) {
	auto&& __r = ((QMainWindow__Go*)self)->cppgoBase_redirected((QPoint*)offset);
	return (void*)(__r);
}
void* QMainWindow_sharedPainter_default(void* self) {
	auto&& __r = ((QMainWindow__Go*)self)->cppgoBase_sharedPainter();
	return (void*)(__r);
}
void QMainWindow_inputMethodEvent_default(void* self, void* arg1) {
	((QMainWindow__Go*)self)->cppgoBase_inputMethodEvent((QInputMethodEvent*)arg1);
}
void* QMainWindow_inputMethodQuery_default(void* self, int arg1) {
	auto&& __r = ((QMainWindow*)self)->QMainWindow::inputMethodQuery((Qt::InputMethodQuery)(arg1));
	return new QVariant(__r);
}
int QMainWindow_focusNextPrevChild_default(void* self, int next) {
	auto&& __r = ((QMainWindow__Go*)self)->cppgoBase_focusNextPrevChild(next);
	return __r;
}
int QMainWindow_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QMainWindow*)self)->QMainWindow::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QMainWindow_timerEvent_default(void* self, void* event) {
	((QMainWindow__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QMainWindow_childEvent_default(void* self, void* event) {
	((QMainWindow__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QMainWindow_customEvent_default(void* self, void* event) {
	((QMainWindow__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QMainWindow_connectNotify_default(void* self, void* signal) {
	((QMainWindow__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QMainWindow_disconnectNotify_default(void* self, void* signal) {
	((QMainWindow__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QMainWindow_connect_iconSizeChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QMainWindow*)self, &QMainWindow::iconSizeChanged, (QMainWindow*)self, [__id](const QSize& iconSize){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&iconSize);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QMainWindow);
	__conns_QMainWindow[__id] = __c;
}
void QMainWindow_connect_toolButtonStyleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QMainWindow*)self, &QMainWindow::toolButtonStyleChanged, (QMainWindow*)self, [__id](Qt::ToolButtonStyle toolButtonStyle){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, (int)(toolButtonStyle));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QMainWindow);
	__conns_QMainWindow[__id] = __c;
}
void QMainWindow_connect_tabifiedDockWidgetActivated(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QMainWindow*)self, &QMainWindow::tabifiedDockWidgetActivated, (QMainWindow*)self, [__id](QDockWidget* dockWidget){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)dockWidget);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QMainWindow);
	__conns_QMainWindow[__id] = __c;
}
void QMainWindow_connect_windowTitleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QMainWindow*)self, &QMainWindow::windowTitleChanged, (QMainWindow*)self, [__id](const QString& title){
		QByteArray title_ba = title.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)title_ba.constData(), title_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QMainWindow);
	__conns_QMainWindow[__id] = __c;
}
void QMainWindow_connect_windowIconChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QMainWindow*)self, &QMainWindow::windowIconChanged, (QMainWindow*)self, [__id](const QIcon& icon){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&icon);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QMainWindow);
	__conns_QMainWindow[__id] = __c;
}
void QMainWindow_connect_windowIconTextChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QMainWindow*)self, &QMainWindow::windowIconTextChanged, (QMainWindow*)self, [__id](const QString& iconText){
		QByteArray iconText_ba = iconText.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)iconText_ba.constData(), iconText_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QMainWindow);
	__conns_QMainWindow[__id] = __c;
}
void QMainWindow_connect_customContextMenuRequested(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QMainWindow*)self, &QMainWindow::customContextMenuRequested, (QMainWindow*)self, [__id](const QPoint& pos){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&pos);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QMainWindow);
	__conns_QMainWindow[__id] = __c;
}
void QMainWindow_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QMainWindow*)self, &QMainWindow::destroyed, (QMainWindow*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QMainWindow);
	__conns_QMainWindow[__id] = __c;
}
void QMainWindow_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QMainWindow*)self, &QMainWindow::destroyed, (QMainWindow*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QMainWindow);
	__conns_QMainWindow[__id] = __c;
}
void QMainWindow_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QMainWindow*)self, &QMainWindow::objectNameChanged, (QMainWindow*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QMainWindow);
	__conns_QMainWindow[__id] = __c;
}
void QMainWindow__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QMainWindow);
		auto __it = __conns_QMainWindow.find(id);
		if (__it == __conns_QMainWindow.end()) return;
		__c = __it->second;
		__conns_QMainWindow.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QMainWindow__Go_new(long long id) { return new QMainWindow__Go(id); }
void QMainWindow_delete(void* self) { delete (QMainWindow*)self; }

}

#include <QtWidgets/qstackedlayout.h>
#include <QtWidgets/qlayout.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtWidgets/qlayoutitem.h>
#include <QtCore/qsize.h>
#include <QtCore/qrect.h>
#include <QtCore/qmargins.h>
#include <QtCore/qnamespace.h>
#include <QtWidgets/qsizepolicy.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qobject.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qvariant.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qstackedlayout.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QStackedLayout__Go : public QStackedLayout {
public:
	long long __goId;
	unsigned char __goOn[31] = {};
	template<typename... A> QStackedLayout__Go(long long id, A&&... a) : QStackedLayout(std::forward<A>(a)...), __goId(id) {}
	using QStackedLayout::widgetEvent;
	using QStackedLayout::addChildLayout;
	using QStackedLayout::addChildWidget;
	using QStackedLayout::adoptLayout;
	using QStackedLayout::alignmentRect;
	using QStackedLayout::sender;
	using QStackedLayout::senderSignalIndex;
	using QStackedLayout::receivers;
	using QStackedLayout::isSignalConnected;
	void cppgoBase_childEvent(QChildEvent* __a0) { QStackedLayout::childEvent(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QStackedLayout::timerEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QStackedLayout::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QStackedLayout::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QStackedLayout::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QStackedLayout::metaObject();
	}
	int count() const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -5627843058667511180LL, &__h); if (__h) return __r; }
		return QStackedLayout::count();
	}
	void addItem(QLayoutItem* item) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6884465960736393867LL, (void*)item, &__h); if (__h) return; }
		QStackedLayout::addItem(item);
	}
	QSize sizeHint() const override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 729690261457684101LL, &__h); if (__h) return *(QSize*)__r; }
		return QStackedLayout::sizeHint();
	}
	QSize minimumSize() const override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, -7246748916541192690LL, &__h); if (__h) return *(QSize*)__r; }
		return QStackedLayout::minimumSize();
	}
	QLayoutItem* itemAt(int arg1) const override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7508512214013391107LL, arg1, &__h); if (__h) return (QLayoutItem*)__r; }
		return QStackedLayout::itemAt(arg1);
	}
	QLayoutItem* takeAt(int arg1) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7549889108533486643LL, arg1, &__h); if (__h) return (QLayoutItem*)__r; }
		return QStackedLayout::takeAt(arg1);
	}
	void setGeometry(const QRect& rect) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5684820848726969227LL, (void*)&rect, &__h); if (__h) return; }
		QStackedLayout::setGeometry(rect);
	}
	bool hasHeightForWidth() const override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8266762397995423915LL, &__h); if (__h) return __r; }
		return QStackedLayout::hasHeightForWidth();
	}
	int heightForWidth(int width) const override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 4311534060495737339LL, width, &__h); if (__h) return __r; }
		return QStackedLayout::heightForWidth(width);
	}
	int spacing() const override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -4587105490415799682LL, &__h); if (__h) return __r; }
		return QStackedLayout::spacing();
	}
	void setSpacing(int arg1) override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 4039479028406881496LL, arg1, &__h); if (__h) return; }
		QStackedLayout::setSpacing(arg1);
	}
	void invalidate() override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 6991341317990632368LL, &__h); if (__h) return; }
		QStackedLayout::invalidate();
	}
	QRect geometry() const override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1050490530495513297LL, &__h); if (__h) return *(QRect*)__r; }
		return QStackedLayout::geometry();
	}
	QFlags<Qt::Orientation> expandingDirections() const override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, 415665469015095569LL, &__h); if (__h) return QFlags<Qt::Orientation>(__r); }
		return QStackedLayout::expandingDirections();
	}
	QSize maximumSize() const override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 4158099606088301136LL, &__h); if (__h) return *(QSize*)__r; }
		return QStackedLayout::maximumSize();
	}
	int indexOf(const QWidget* arg1) const override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, -7258473138030329570LL, (void*)arg1, &__h); if (__h) return __r; }
		return QStackedLayout::indexOf(arg1);
	}
	int indexOf(const QLayoutItem* arg1) const override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, -3276435765994963516LL, (void*)arg1, &__h); if (__h) return __r; }
		return QStackedLayout::indexOf(arg1);
	}
	bool isEmpty() const override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -626822430251098498LL, &__h); if (__h) return __r; }
		return QStackedLayout::isEmpty();
	}
	QFlags<QSizePolicy::ControlType> controlTypes() const override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, 3278947275714896887LL, &__h); if (__h) return QFlags<QSizePolicy::ControlType>(__r); }
		return QStackedLayout::controlTypes();
	}
	QLayoutItem* replaceWidget(QWidget* from, QWidget* to, QFlags<Qt::FindChildOption> options) override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_voidp_voidp_voidp_int) { void* __r = ((void*(*)(long long, long long, void*, void*, int, int*))cppgo_virt_voidp_voidp_voidp_int)(__goId, -712862742159708061LL, (void*)from, (void*)to, (int)(options), &__h); if (__h) return (QLayoutItem*)__r; }
		return QStackedLayout::replaceWidget(from, to, options);
	}
	QLayout* layout() override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, -7183390618627992737LL, &__h); if (__h) return (QLayout*)__r; }
		return QStackedLayout::layout();
	}
	void childEvent(QChildEvent* e) override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)e, &__h); if (__h) return; }
		QStackedLayout::childEvent(e);
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QStackedLayout::event(event);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QStackedLayout::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QStackedLayout::timerEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QStackedLayout::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QStackedLayout::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QStackedLayout::disconnectNotify(signal);
	}
	int minimumHeightForWidth(int arg1) const override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 5037837229528203129LL, arg1, &__h); if (__h) return __r; }
		return QStackedLayout::minimumHeightForWidth(arg1);
	}
	QSpacerItem* spacerItem() override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 2272240669359776278LL, &__h); if (__h) return (QSpacerItem*)__r; }
		return QStackedLayout::spacerItem();
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QStackedLayout;
static std::mutex __connsMu_QStackedLayout;

extern "C" {

void* QStackedLayout_base(void* self) { return (void*)static_cast<QLayout*>((QStackedLayout*)self); }
void QStackedLayout_chain(void* self, void** out) {
	QStackedLayout* __p = (QStackedLayout*)self;
	QLayout* __b0 = static_cast<QLayout*>(__p); out[0] = (void*)__b0;
	QObject* __b1 = static_cast<QObject*>(__b0); out[1] = (void*)__b1;
}
void QStackedLayout__set_override(void* self, int idx) { ((QStackedLayout__Go*)self)->__goOn[idx] = 1; }
void* QStackedLayout_fromQObject(void* o) { return (void*)qobject_cast<QStackedLayout*>((QObject*)o); }
void* QStackedLayout_new(long long __goid, void* parent) {
	return new QStackedLayout__Go(__goid, (QWidget*)parent);
}
void* QStackedLayout_new2(long long __goid) {
	return new QStackedLayout__Go(__goid);
}
void* QStackedLayout_new3(long long __goid, void* parentLayout) {
	return new QStackedLayout__Go(__goid, (QLayout*)parentLayout);
}
char* QStackedLayout_tr(const char* s, long long* __out_len) {
	auto&& __r = QStackedLayout::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QStackedLayout_addWidget(void* self, void* w) {
	auto&& __r = ((QStackedLayout*)self)->addWidget((QWidget*)w);
	return __r;
}
int QStackedLayout_insertWidget(void* self, int index, void* w) {
	auto&& __r = ((QStackedLayout*)self)->insertWidget(index, (QWidget*)w);
	return __r;
}
void* QStackedLayout_currentWidget(void* self) {
	auto&& __r = ((QStackedLayout*)self)->currentWidget();
	return (void*)(__r);
}
int QStackedLayout_currentIndex(void* self) {
	auto&& __r = ((QStackedLayout*)self)->currentIndex();
	return __r;
}
void* QStackedLayout_widget(void* self, int arg1) {
	auto&& __r = ((QStackedLayout*)self)->widget(arg1);
	return (void*)(__r);
}
int QStackedLayout_stackingMode(void* self) {
	auto&& __r = ((QStackedLayout*)self)->stackingMode();
	return (int)(__r);
}
void QStackedLayout_setStackingMode(void* self, int stackingMode) {
	((QStackedLayout*)self)->setStackingMode((QStackedLayout::StackingMode)(stackingMode));
}
void QStackedLayout_widgetRemoved(void* self, int index) {
	((QStackedLayout*)self)->widgetRemoved(index);
}
void QStackedLayout_currentChanged(void* self, int index) {
	((QStackedLayout*)self)->currentChanged(index);
}
void QStackedLayout_widgetAdded(void* self, int index) {
	((QStackedLayout*)self)->widgetAdded(index);
}
void QStackedLayout_setCurrentIndex(void* self, int index) {
	((QStackedLayout*)self)->setCurrentIndex(index);
}
void QStackedLayout_setCurrentWidget(void* self, void* w) {
	((QStackedLayout*)self)->setCurrentWidget((QWidget*)w);
}
void* QStackedLayout_closestAcceptableSize(void* w, void* s) {
	auto&& __r = QStackedLayout::closestAcceptableSize((QWidget*)w, *(QSize*)s);
	return new QSize(__r);
}
void QStackedLayout_widgetEvent(void* self, void* arg1) {
	((QStackedLayout__Go*)self)->widgetEvent((QEvent*)arg1);
}
void QStackedLayout_addChildLayout(void* self, void* l) {
	((QStackedLayout__Go*)self)->addChildLayout((QLayout*)l);
}
void QStackedLayout_addChildWidget(void* self, void* w) {
	((QStackedLayout__Go*)self)->addChildWidget((QWidget*)w);
}
int QStackedLayout_adoptLayout(void* self, void* layout) {
	auto&& __r = ((QStackedLayout__Go*)self)->adoptLayout((QLayout*)layout);
	return __r;
}
void* QStackedLayout_alignmentRect(void* self, void* arg1) {
	auto&& __r = ((QStackedLayout__Go*)self)->alignmentRect(*(QRect*)arg1);
	return new QRect(__r);
}
void* QStackedLayout_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QStackedLayout::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QStackedLayout_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QStackedLayout::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QStackedLayout_sender(void* self) {
	auto&& __r = ((QStackedLayout__Go*)self)->sender();
	return (void*)(__r);
}
int QStackedLayout_senderSignalIndex(void* self) {
	auto&& __r = ((QStackedLayout__Go*)self)->senderSignalIndex();
	return __r;
}
int QStackedLayout_receivers(void* self, const char* signal) {
	auto&& __r = ((QStackedLayout__Go*)self)->receivers(signal);
	return __r;
}
int QStackedLayout_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QStackedLayout__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QStackedLayout_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QStackedLayout::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QStackedLayout_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QStackedLayout::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QStackedLayout_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QStackedLayout::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QStackedLayout_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QStackedLayout::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QStackedLayout_metaObject_default(void* self) {
	auto&& __r = ((QStackedLayout*)self)->QStackedLayout::metaObject();
	return (void*)(__r);
}
int QStackedLayout_count_default(void* self) {
	auto&& __r = ((QStackedLayout*)self)->QStackedLayout::count();
	return __r;
}
void QStackedLayout_addItem_default(void* self, void* item) {
	((QStackedLayout*)self)->QStackedLayout::addItem((QLayoutItem*)item);
}
void* QStackedLayout_sizeHint_default(void* self) {
	auto&& __r = ((QStackedLayout*)self)->QStackedLayout::sizeHint();
	return new QSize(__r);
}
void* QStackedLayout_minimumSize_default(void* self) {
	auto&& __r = ((QStackedLayout*)self)->QStackedLayout::minimumSize();
	return new QSize(__r);
}
void* QStackedLayout_itemAt_default(void* self, int arg1) {
	auto&& __r = ((QStackedLayout*)self)->QStackedLayout::itemAt(arg1);
	return (void*)(__r);
}
void* QStackedLayout_takeAt_default(void* self, int arg1) {
	auto&& __r = ((QStackedLayout*)self)->QStackedLayout::takeAt(arg1);
	return (void*)(__r);
}
void QStackedLayout_setGeometry_default(void* self, void* rect) {
	((QStackedLayout*)self)->QStackedLayout::setGeometry(*(QRect*)rect);
}
int QStackedLayout_hasHeightForWidth_default(void* self) {
	auto&& __r = ((QStackedLayout*)self)->QStackedLayout::hasHeightForWidth();
	return __r;
}
int QStackedLayout_heightForWidth_default(void* self, int width) {
	auto&& __r = ((QStackedLayout*)self)->QStackedLayout::heightForWidth(width);
	return __r;
}
int QStackedLayout_spacing_default(void* self) {
	auto&& __r = ((QStackedLayout*)self)->QStackedLayout::spacing();
	return __r;
}
void QStackedLayout_setSpacing_default(void* self, int arg1) {
	((QStackedLayout*)self)->QStackedLayout::setSpacing(arg1);
}
void QStackedLayout_invalidate_default(void* self) {
	((QStackedLayout*)self)->QStackedLayout::invalidate();
}
void* QStackedLayout_geometry_default(void* self) {
	auto&& __r = ((QStackedLayout*)self)->QStackedLayout::geometry();
	return new QRect(__r);
}
int QStackedLayout_expandingDirections_default(void* self) {
	auto&& __r = ((QStackedLayout*)self)->QStackedLayout::expandingDirections();
	return (int)(__r);
}
void* QStackedLayout_maximumSize_default(void* self) {
	auto&& __r = ((QStackedLayout*)self)->QStackedLayout::maximumSize();
	return new QSize(__r);
}
int QStackedLayout_indexOf_default(void* self, void* arg1) {
	auto&& __r = ((QStackedLayout*)self)->QStackedLayout::indexOf((QWidget*)arg1);
	return __r;
}
int QStackedLayout_isEmpty_default(void* self) {
	auto&& __r = ((QStackedLayout*)self)->QStackedLayout::isEmpty();
	return __r;
}
int QStackedLayout_controlTypes_default(void* self) {
	auto&& __r = ((QStackedLayout*)self)->QStackedLayout::controlTypes();
	return (int)(__r);
}
void* QStackedLayout_replaceWidget_default(void* self, void* from, void* to, int options) {
	auto&& __r = ((QStackedLayout*)self)->QStackedLayout::replaceWidget((QWidget*)from, (QWidget*)to, QFlags<Qt::FindChildOption>(options));
	return (void*)(__r);
}
void* QStackedLayout_layout_default(void* self) {
	auto&& __r = ((QStackedLayout*)self)->QStackedLayout::layout();
	return (void*)(__r);
}
void QStackedLayout_childEvent_default(void* self, void* e) {
	((QStackedLayout__Go*)self)->cppgoBase_childEvent((QChildEvent*)e);
}
int QStackedLayout_event_default(void* self, void* event) {
	auto&& __r = ((QStackedLayout*)self)->QStackedLayout::event((QEvent*)event);
	return __r;
}
int QStackedLayout_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QStackedLayout*)self)->QStackedLayout::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QStackedLayout_timerEvent_default(void* self, void* event) {
	((QStackedLayout__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QStackedLayout_customEvent_default(void* self, void* event) {
	((QStackedLayout__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QStackedLayout_connectNotify_default(void* self, void* signal) {
	((QStackedLayout__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QStackedLayout_disconnectNotify_default(void* self, void* signal) {
	((QStackedLayout__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
int QStackedLayout_minimumHeightForWidth_default(void* self, int arg1) {
	auto&& __r = ((QStackedLayout*)self)->QStackedLayout::minimumHeightForWidth(arg1);
	return __r;
}
void* QStackedLayout_spacerItem_default(void* self) {
	auto&& __r = ((QStackedLayout*)self)->QStackedLayout::spacerItem();
	return (void*)(__r);
}
void QStackedLayout_connect_widgetRemoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStackedLayout*)self, &QStackedLayout::widgetRemoved, (QStackedLayout*)self, [__id](int index){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, index);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStackedLayout);
	__conns_QStackedLayout[__id] = __c;
}
void QStackedLayout_connect_currentChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStackedLayout*)self, &QStackedLayout::currentChanged, (QStackedLayout*)self, [__id](int index){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, index);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStackedLayout);
	__conns_QStackedLayout[__id] = __c;
}
void QStackedLayout_connect_widgetAdded(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStackedLayout*)self, &QStackedLayout::widgetAdded, (QStackedLayout*)self, [__id](int index){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, index);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStackedLayout);
	__conns_QStackedLayout[__id] = __c;
}
void QStackedLayout_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStackedLayout*)self, &QStackedLayout::destroyed, (QStackedLayout*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStackedLayout);
	__conns_QStackedLayout[__id] = __c;
}
void QStackedLayout_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStackedLayout*)self, &QStackedLayout::destroyed, (QStackedLayout*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStackedLayout);
	__conns_QStackedLayout[__id] = __c;
}
void QStackedLayout_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QStackedLayout*)self, &QStackedLayout::objectNameChanged, (QStackedLayout*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QStackedLayout);
	__conns_QStackedLayout[__id] = __c;
}
void QStackedLayout__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QStackedLayout);
		auto __it = __conns_QStackedLayout.find(id);
		if (__it == __conns_QStackedLayout.end()) return;
		__c = __it->second;
		__conns_QStackedLayout.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QStackedLayout__Go_new(long long id) { return new QStackedLayout__Go(id); }
void QStackedLayout_delete(void* self) { delete (QStackedLayout*)self; }

}

#include <QtWidgets/qwizard.h>
#include <QtWidgets/qdialog.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qvariant.h>
#include <QtWidgets/qabstractbutton.h>
#include <QtGui/qpixmap.h>
#include <QtCore/qsize.h>
#include <QtCore/qcoreevent.h>
#include <QtGui/qevent.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qobject.h>
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
#include <QtWidgets/qsizepolicy.h>
#include <QtCore/qmargins.h>
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
#include "gen_qwizard.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QWizard__Go : public QWizard {
public:
	long long __goId;
	unsigned char __goOn[57] = {};
	template<typename... A> QWizard__Go(long long id, A&&... a) : QWizard(std::forward<A>(a)...), __goId(id) {}
	using QWizard::adjustPosition;
	using QWizard::updateMicroFocus;
	using QWizard::create;
	using QWizard::destroy;
	using QWizard::focusNextChild;
	using QWizard::focusPreviousChild;
	using QWizard::sender;
	using QWizard::senderSignalIndex;
	using QWizard::receivers;
	using QWizard::isSignalConnected;
	using QWizard::getDecodedMetricF;
	bool cppgoBase_event(QEvent* __a0) { return QWizard::event(__a0); }
	void cppgoBase_resizeEvent(QResizeEvent* __a0) { QWizard::resizeEvent(__a0); }
	void cppgoBase_paintEvent(QPaintEvent* __a0) { QWizard::paintEvent(__a0); }
	void cppgoBase_done(int __a0) { QWizard::done(__a0); }
	void cppgoBase_initializePage(int __a0) { QWizard::initializePage(__a0); }
	void cppgoBase_cleanupPage(int __a0) { QWizard::cleanupPage(__a0); }
	void cppgoBase_keyPressEvent(QKeyEvent* __a0) { QWizard::keyPressEvent(__a0); }
	void cppgoBase_closeEvent(QCloseEvent* __a0) { QWizard::closeEvent(__a0); }
	void cppgoBase_showEvent(QShowEvent* __a0) { QWizard::showEvent(__a0); }
	void cppgoBase_contextMenuEvent(QContextMenuEvent* __a0) { QWizard::contextMenuEvent(__a0); }
	bool cppgoBase_eventFilter(QObject* __a0, QEvent* __a1) { return QWizard::eventFilter(__a0, __a1); }
	void cppgoBase_mousePressEvent(QMouseEvent* __a0) { QWizard::mousePressEvent(__a0); }
	void cppgoBase_mouseReleaseEvent(QMouseEvent* __a0) { QWizard::mouseReleaseEvent(__a0); }
	void cppgoBase_mouseDoubleClickEvent(QMouseEvent* __a0) { QWizard::mouseDoubleClickEvent(__a0); }
	void cppgoBase_mouseMoveEvent(QMouseEvent* __a0) { QWizard::mouseMoveEvent(__a0); }
	void cppgoBase_wheelEvent(QWheelEvent* __a0) { QWizard::wheelEvent(__a0); }
	void cppgoBase_keyReleaseEvent(QKeyEvent* __a0) { QWizard::keyReleaseEvent(__a0); }
	void cppgoBase_focusInEvent(QFocusEvent* __a0) { QWizard::focusInEvent(__a0); }
	void cppgoBase_focusOutEvent(QFocusEvent* __a0) { QWizard::focusOutEvent(__a0); }
	void cppgoBase_enterEvent(QEnterEvent* __a0) { QWizard::enterEvent(__a0); }
	void cppgoBase_leaveEvent(QEvent* __a0) { QWizard::leaveEvent(__a0); }
	void cppgoBase_moveEvent(QMoveEvent* __a0) { QWizard::moveEvent(__a0); }
	void cppgoBase_tabletEvent(QTabletEvent* __a0) { QWizard::tabletEvent(__a0); }
	void cppgoBase_actionEvent(QActionEvent* __a0) { QWizard::actionEvent(__a0); }
	void cppgoBase_dragEnterEvent(QDragEnterEvent* __a0) { QWizard::dragEnterEvent(__a0); }
	void cppgoBase_dragMoveEvent(QDragMoveEvent* __a0) { QWizard::dragMoveEvent(__a0); }
	void cppgoBase_dragLeaveEvent(QDragLeaveEvent* __a0) { QWizard::dragLeaveEvent(__a0); }
	void cppgoBase_dropEvent(QDropEvent* __a0) { QWizard::dropEvent(__a0); }
	void cppgoBase_hideEvent(QHideEvent* __a0) { QWizard::hideEvent(__a0); }
	void cppgoBase_changeEvent(QEvent* __a0) { QWizard::changeEvent(__a0); }
	int cppgoBase_metric(int __a0) { return QWizard::metric((QPaintDevice::PaintDeviceMetric)__a0); }
	void cppgoBase_initPainter(QPainter* __a0) { QWizard::initPainter(__a0); }
	QPaintDevice* cppgoBase_redirected(QPoint* __a0) { return QWizard::redirected(__a0); }
	QPainter* cppgoBase_sharedPainter() { return QWizard::sharedPainter(); }
	void cppgoBase_inputMethodEvent(QInputMethodEvent* __a0) { QWizard::inputMethodEvent(__a0); }
	bool cppgoBase_focusNextPrevChild(bool __a0) { return QWizard::focusNextPrevChild(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QWizard::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QWizard::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QWizard::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QWizard::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QWizard::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QWizard::metaObject();
	}
	bool validateCurrentPage() override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, 7851200041405997743LL, &__h); if (__h) return __r; }
		return QWizard::validateCurrentPage();
	}
	int nextId() const override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, 5814903794109731573LL, &__h); if (__h) return __r; }
		return QWizard::nextId();
	}
	void setVisible(bool visible) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 2110182719375245575LL, visible, &__h); if (__h) return; }
		QWizard::setVisible(visible);
	}
	QSize sizeHint() const override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 729690261457684101LL, &__h); if (__h) return *(QSize*)__r; }
		return QWizard::sizeHint();
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QWizard::event(event);
	}
	void resizeEvent(QResizeEvent* event) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4302917285655014661LL, (void*)event, &__h); if (__h) return; }
		QWizard::resizeEvent(event);
	}
	void paintEvent(QPaintEvent* event) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5290442163167458465LL, (void*)event, &__h); if (__h) return; }
		QWizard::paintEvent(event);
	}
	bool nativeEvent(const QByteArray& eventType, void* message, qintptr* result) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_int_voidp_voidp_longlongp) { int __r = ((int(*)(long long, long long, void*, void*, long long*, int*))cppgo_virt_int_voidp_voidp_longlongp)(__goId, -2039166586246443706LL, (void*)&eventType, (void*)message, (long long*)result, &__h); if (__h) return __r; }
		return QWizard::nativeEvent(eventType, message, result);
	}
	void done(int result) override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, -2570997490807640431LL, result, &__h); if (__h) return; }
		QWizard::done(result);
	}
	void initializePage(int id) override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 3074997362940195296LL, id, &__h); if (__h) return; }
		QWizard::initializePage(id);
	}
	void cleanupPage(int id) override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 5913490871578514792LL, id, &__h); if (__h) return; }
		QWizard::cleanupPage(id);
	}
	QSize minimumSizeHint() const override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 7049264599837532491LL, &__h); if (__h) return *(QSize*)__r; }
		return QWizard::minimumSizeHint();
	}
	void open() override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, -554057428128861975LL, &__h); if (__h) return; }
		QWizard::open();
	}
	int exec() override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -1745322561779819684LL, &__h); if (__h) return __r; }
		return QWizard::exec();
	}
	void accept() override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 9106131928185084681LL, &__h); if (__h) return; }
		QWizard::accept();
	}
	void reject() override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 2875220963396170150LL, &__h); if (__h) return; }
		QWizard::reject();
	}
	void keyPressEvent(QKeyEvent* arg1) override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1023791977679793185LL, (void*)arg1, &__h); if (__h) return; }
		QWizard::keyPressEvent(arg1);
	}
	void closeEvent(QCloseEvent* arg1) override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1989937050019464783LL, (void*)arg1, &__h); if (__h) return; }
		QWizard::closeEvent(arg1);
	}
	void showEvent(QShowEvent* arg1) override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6825369174751916LL, (void*)arg1, &__h); if (__h) return; }
		QWizard::showEvent(arg1);
	}
	void contextMenuEvent(QContextMenuEvent* arg1) override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6698043222227006015LL, (void*)arg1, &__h); if (__h) return; }
		QWizard::contextMenuEvent(arg1);
	}
	bool eventFilter(QObject* arg1, QEvent* arg2) override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)arg1, (void*)arg2, &__h); if (__h) return __r; }
		return QWizard::eventFilter(arg1, arg2);
	}
	int devType() const override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8659679893868524044LL, &__h); if (__h) return __r; }
		return QWizard::devType();
	}
	int heightForWidth(int arg1) const override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 4311534060495737339LL, arg1, &__h); if (__h) return __r; }
		return QWizard::heightForWidth(arg1);
	}
	bool hasHeightForWidth() const override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8266762397995423915LL, &__h); if (__h) return __r; }
		return QWizard::hasHeightForWidth();
	}
	QPaintEngine* paintEngine() const override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 6398007404756346381LL, &__h); if (__h) return (QPaintEngine*)__r; }
		return QWizard::paintEngine();
	}
	void mousePressEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -332090898078543319LL, (void*)event, &__h); if (__h) return; }
		QWizard::mousePressEvent(event);
	}
	void mouseReleaseEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4415920049297795213LL, (void*)event, &__h); if (__h) return; }
		QWizard::mouseReleaseEvent(event);
	}
	void mouseDoubleClickEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4029073905697023887LL, (void*)event, &__h); if (__h) return; }
		QWizard::mouseDoubleClickEvent(event);
	}
	void mouseMoveEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4255494004099806581LL, (void*)event, &__h); if (__h) return; }
		QWizard::mouseMoveEvent(event);
	}
	void wheelEvent(QWheelEvent* event) override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1919988340934511316LL, (void*)event, &__h); if (__h) return; }
		QWizard::wheelEvent(event);
	}
	void keyReleaseEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 772598475950050299LL, (void*)event, &__h); if (__h) return; }
		QWizard::keyReleaseEvent(event);
	}
	void focusInEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[32] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 975043767033552552LL, (void*)event, &__h); if (__h) return; }
		QWizard::focusInEvent(event);
	}
	void focusOutEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[33] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -385596010576048993LL, (void*)event, &__h); if (__h) return; }
		QWizard::focusOutEvent(event);
	}
	void enterEvent(QEnterEvent* event) override {
		int __h = 0;
		if (__goOn[34] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6316270346668237895LL, (void*)event, &__h); if (__h) return; }
		QWizard::enterEvent(event);
	}
	void leaveEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[35] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -551070009623711640LL, (void*)event, &__h); if (__h) return; }
		QWizard::leaveEvent(event);
	}
	void moveEvent(QMoveEvent* event) override {
		int __h = 0;
		if (__goOn[36] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -9161949150489525396LL, (void*)event, &__h); if (__h) return; }
		QWizard::moveEvent(event);
	}
	void tabletEvent(QTabletEvent* event) override {
		int __h = 0;
		if (__goOn[37] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -212068866860701141LL, (void*)event, &__h); if (__h) return; }
		QWizard::tabletEvent(event);
	}
	void actionEvent(QActionEvent* event) override {
		int __h = 0;
		if (__goOn[38] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1697444140088555581LL, (void*)event, &__h); if (__h) return; }
		QWizard::actionEvent(event);
	}
	void dragEnterEvent(QDragEnterEvent* event) override {
		int __h = 0;
		if (__goOn[39] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 206122715919917589LL, (void*)event, &__h); if (__h) return; }
		QWizard::dragEnterEvent(event);
	}
	void dragMoveEvent(QDragMoveEvent* event) override {
		int __h = 0;
		if (__goOn[40] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1781545137352745978LL, (void*)event, &__h); if (__h) return; }
		QWizard::dragMoveEvent(event);
	}
	void dragLeaveEvent(QDragLeaveEvent* event) override {
		int __h = 0;
		if (__goOn[41] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -3139904158373725006LL, (void*)event, &__h); if (__h) return; }
		QWizard::dragLeaveEvent(event);
	}
	void dropEvent(QDropEvent* event) override {
		int __h = 0;
		if (__goOn[42] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 9107058682754591716LL, (void*)event, &__h); if (__h) return; }
		QWizard::dropEvent(event);
	}
	void hideEvent(QHideEvent* event) override {
		int __h = 0;
		if (__goOn[43] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -199718707876719129LL, (void*)event, &__h); if (__h) return; }
		QWizard::hideEvent(event);
	}
	void changeEvent(QEvent* arg1) override {
		int __h = 0;
		if (__goOn[44] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4959376063826165943LL, (void*)arg1, &__h); if (__h) return; }
		QWizard::changeEvent(arg1);
	}
	int metric(QPaintDevice::PaintDeviceMetric arg1) const override {
		int __h = 0;
		if (__goOn[45] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 427629649267311777LL, (int)(arg1), &__h); if (__h) return __r; }
		return QWizard::metric(arg1);
	}
	void initPainter(QPainter* painter) const override {
		int __h = 0;
		if (__goOn[46] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6477500296980050210LL, (void*)painter, &__h); if (__h) return; }
		QWizard::initPainter(painter);
	}
	QPaintDevice* redirected(QPoint* offset) const override {
		int __h = 0;
		if (__goOn[47] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 4127240633310655048LL, (void*)offset, &__h); if (__h) return (QPaintDevice*)__r; }
		return QWizard::redirected(offset);
	}
	QPainter* sharedPainter() const override {
		int __h = 0;
		if (__goOn[48] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 5792771328741774991LL, &__h); if (__h) return (QPainter*)__r; }
		return QWizard::sharedPainter();
	}
	void inputMethodEvent(QInputMethodEvent* arg1) override {
		int __h = 0;
		if (__goOn[49] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5438819661942843222LL, (void*)arg1, &__h); if (__h) return; }
		QWizard::inputMethodEvent(arg1);
	}
	QVariant inputMethodQuery(Qt::InputMethodQuery arg1) const override {
		int __h = 0;
		if (__goOn[50] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7345853097340709638LL, (int)(arg1), &__h); if (__h) return *(QVariant*)__r; }
		return QWizard::inputMethodQuery(arg1);
	}
	bool focusNextPrevChild(bool next) override {
		int __h = 0;
		if (__goOn[51] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 5898328517505784191LL, next, &__h); if (__h) return __r; }
		return QWizard::focusNextPrevChild(next);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[52] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QWizard::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[53] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QWizard::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[54] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QWizard::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[55] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QWizard::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[56] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QWizard::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QWizard;
static std::mutex __connsMu_QWizard;

extern "C" {

void* QWizard_base(void* self) { return (void*)static_cast<QDialog*>((QWizard*)self); }
void QWizard_chain(void* self, void** out) {
	QWizard* __p = (QWizard*)self;
	QDialog* __b0 = static_cast<QDialog*>(__p); out[0] = (void*)__b0;
	QWidget* __b1 = static_cast<QWidget*>(__b0); out[1] = (void*)__b1;
	QObject* __b2 = static_cast<QObject*>(__b1); out[2] = (void*)__b2;
}
void QWizard__set_override(void* self, int idx) { ((QWizard__Go*)self)->__goOn[idx] = 1; }
void* QWizard_fromQObject(void* o) { return (void*)qobject_cast<QWizard*>((QObject*)o); }
void* QWizard_new(long long __goid, void* parent) {
	return new QWizard__Go(__goid, (QWidget*)parent);
}
void* QWizard_new2(long long __goid) {
	return new QWizard__Go(__goid);
}
void* QWizard_new3(long long __goid, void* parent, int flags) {
	return new QWizard__Go(__goid, (QWidget*)parent, QFlags<Qt::WindowType>(flags));
}
char* QWizard_tr(const char* s, long long* __out_len) {
	auto&& __r = QWizard::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QWizard_addPage(void* self, void* page) {
	auto&& __r = ((QWizard*)self)->addPage((QWizardPage*)page);
	return __r;
}
void QWizard_setPage(void* self, int id, void* page) {
	((QWizard*)self)->setPage(id, (QWizardPage*)page);
}
void QWizard_removePage(void* self, int id) {
	((QWizard*)self)->removePage(id);
}
void* QWizard_page(void* self, int id) {
	auto&& __r = ((QWizard*)self)->page(id);
	return (void*)(__r);
}
int QWizard_hasVisitedPage(void* self, int id) {
	auto&& __r = ((QWizard*)self)->hasVisitedPage(id);
	return __r;
}
int* QWizard_visitedIds(void* self, long long* __out_len) {
	auto&& __r = ((QWizard*)self)->visitedIds();
	*__out_len = __r.size();
	int* __arr = (int*)malloc(sizeof(int) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = __e;
	return __arr;
}
int* QWizard_pageIds(void* self, long long* __out_len) {
	auto&& __r = ((QWizard*)self)->pageIds();
	*__out_len = __r.size();
	int* __arr = (int*)malloc(sizeof(int) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = __e;
	return __arr;
}
void QWizard_setStartId(void* self, int id) {
	((QWizard*)self)->setStartId(id);
}
int QWizard_startId(void* self) {
	auto&& __r = ((QWizard*)self)->startId();
	return __r;
}
void* QWizard_currentPage(void* self) {
	auto&& __r = ((QWizard*)self)->currentPage();
	return (void*)(__r);
}
int QWizard_currentId(void* self) {
	auto&& __r = ((QWizard*)self)->currentId();
	return __r;
}
int QWizard_validateCurrentPage(void* self) {
	auto&& __r = ((QWizard*)self)->validateCurrentPage();
	return __r;
}
int QWizard_nextId(void* self) {
	auto&& __r = ((QWizard*)self)->nextId();
	return __r;
}
void QWizard_setField(void* self, const char* name, long long name_len, void* value) {
	QString name_qstr = QString::fromUtf8(name, name_len);
	((QWizard*)self)->setField(name_qstr, *(QVariant*)value);
}
void* QWizard_field(void* self, const char* name, long long name_len) {
	QString name_qstr = QString::fromUtf8(name, name_len);
	auto&& __r = ((QWizard*)self)->field(name_qstr);
	return new QVariant(__r);
}
void QWizard_setWizardStyle(void* self, int style) {
	((QWizard*)self)->setWizardStyle((QWizard::WizardStyle)(style));
}
int QWizard_wizardStyle(void* self) {
	auto&& __r = ((QWizard*)self)->wizardStyle();
	return (int)(__r);
}
void QWizard_setOption(void* self, int option) {
	((QWizard*)self)->setOption((QWizard::WizardOption)(option));
}
int QWizard_testOption(void* self, int option) {
	auto&& __r = ((QWizard*)self)->testOption((QWizard::WizardOption)(option));
	return __r;
}
void QWizard_setOptions(void* self, int options) {
	((QWizard*)self)->setOptions(QFlags<QWizard::WizardOption>(options));
}
int QWizard_options(void* self) {
	auto&& __r = ((QWizard*)self)->options();
	return (int)(__r);
}
void QWizard_setButtonText(void* self, int which, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	((QWizard*)self)->setButtonText((QWizard::WizardButton)(which), text_qstr);
}
char* QWizard_buttonText(void* self, int which, long long* __out_len) {
	auto&& __r = ((QWizard*)self)->buttonText((QWizard::WizardButton)(which));
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QWizard_setButtonLayout(void* self, int* layout, long long layout_len) {
	QList<QWizard::WizardButton> layout_list; layout_list.reserve(layout_len);
	for (long long __i = 0; __i < layout_len; __i++) layout_list.append((QWizard::WizardButton)layout[__i]);
	((QWizard*)self)->setButtonLayout(layout_list);
}
void QWizard_setButton(void* self, int which, void* button) {
	((QWizard*)self)->setButton((QWizard::WizardButton)(which), (QAbstractButton*)button);
}
void* QWizard_button(void* self, int which) {
	auto&& __r = ((QWizard*)self)->button((QWizard::WizardButton)(which));
	return (void*)(__r);
}
void QWizard_setTitleFormat(void* self, int format) {
	((QWizard*)self)->setTitleFormat((Qt::TextFormat)(format));
}
int QWizard_titleFormat(void* self) {
	auto&& __r = ((QWizard*)self)->titleFormat();
	return (int)(__r);
}
void QWizard_setSubTitleFormat(void* self, int format) {
	((QWizard*)self)->setSubTitleFormat((Qt::TextFormat)(format));
}
int QWizard_subTitleFormat(void* self) {
	auto&& __r = ((QWizard*)self)->subTitleFormat();
	return (int)(__r);
}
void QWizard_setPixmap(void* self, int which, void* pixmap) {
	((QWizard*)self)->setPixmap((QWizard::WizardPixmap)(which), *(QPixmap*)pixmap);
}
void* QWizard_pixmap(void* self, int which) {
	auto&& __r = ((QWizard*)self)->pixmap((QWizard::WizardPixmap)(which));
	return new QPixmap(__r);
}
void QWizard_setSideWidget(void* self, void* widget) {
	((QWizard*)self)->setSideWidget((QWidget*)widget);
}
void* QWizard_sideWidget(void* self) {
	auto&& __r = ((QWizard*)self)->sideWidget();
	return (void*)(__r);
}
void QWizard_setDefaultProperty(void* self, const char* className, const char* property, const char* changedSignal) {
	((QWizard*)self)->setDefaultProperty(className, property, changedSignal);
}
void QWizard_currentIdChanged(void* self, int id) {
	((QWizard*)self)->currentIdChanged(id);
}
void QWizard_helpRequested(void* self) {
	((QWizard*)self)->helpRequested();
}
void QWizard_customButtonClicked(void* self, int which) {
	((QWizard*)self)->customButtonClicked(which);
}
void QWizard_pageAdded(void* self, int id) {
	((QWizard*)self)->pageAdded(id);
}
void QWizard_pageRemoved(void* self, int id) {
	((QWizard*)self)->pageRemoved(id);
}
void QWizard_back(void* self) {
	((QWizard*)self)->back();
}
void QWizard_next(void* self) {
	((QWizard*)self)->next();
}
void QWizard_setCurrentId(void* self, int id) {
	((QWizard*)self)->setCurrentId(id);
}
void QWizard_restart(void* self) {
	((QWizard*)self)->restart();
}
void QWizard_adjustPosition(void* self, void* arg1) {
	((QWizard__Go*)self)->adjustPosition((QWidget*)arg1);
}
void QWizard_setTabOrder(void* arg1, void* arg2) {
	QWizard::setTabOrder((QWidget*)arg1, (QWidget*)arg2);
}
void* QWizard_mouseGrabber() {
	auto&& __r = QWizard::mouseGrabber();
	return (void*)(__r);
}
void* QWizard_keyboardGrabber() {
	auto&& __r = QWizard::keyboardGrabber();
	return (void*)(__r);
}
void* QWizard_find(unsigned long long arg1) {
	auto&& __r = QWizard::find(arg1);
	return (void*)(__r);
}
void* QWizard_createWindowContainer(void* window) {
	auto&& __r = QWizard::createWindowContainer((QWindow*)window);
	return (void*)(__r);
}
void QWizard_updateMicroFocus(void* self) {
	((QWizard__Go*)self)->updateMicroFocus();
}
void QWizard_create(void* self) {
	((QWizard__Go*)self)->create();
}
void QWizard_destroy(void* self) {
	((QWizard__Go*)self)->destroy();
}
int QWizard_focusNextChild(void* self) {
	auto&& __r = ((QWizard__Go*)self)->focusNextChild();
	return __r;
}
int QWizard_focusPreviousChild(void* self) {
	auto&& __r = ((QWizard__Go*)self)->focusPreviousChild();
	return __r;
}
void* QWizard_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QWizard::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QWizard_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QWizard::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QWizard_sender(void* self) {
	auto&& __r = ((QWizard__Go*)self)->sender();
	return (void*)(__r);
}
int QWizard_senderSignalIndex(void* self) {
	auto&& __r = ((QWizard__Go*)self)->senderSignalIndex();
	return __r;
}
int QWizard_receivers(void* self, const char* signal) {
	auto&& __r = ((QWizard__Go*)self)->receivers(signal);
	return __r;
}
int QWizard_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QWizard__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
double QWizard_devicePixelRatioFScale() {
	auto&& __r = QWizard::devicePixelRatioFScale();
	return __r;
}
int QWizard_encodeMetricF(int metric, double value) {
	auto&& __r = QWizard::encodeMetricF((QPaintDevice::PaintDeviceMetric)(metric), value);
	return __r;
}
double QWizard_getDecodedMetricF(void* self, int metricA, int metricB) {
	auto&& __r = ((QWizard__Go*)self)->getDecodedMetricF((QPaintDevice::PaintDeviceMetric)(metricA), (QPaintDevice::PaintDeviceMetric)(metricB));
	return __r;
}
char* QWizard_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QWizard::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QWizard_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QWizard::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QWizard_setOption2(void* self, int option, int on) {
	((QWizard*)self)->setOption((QWizard::WizardOption)(option), on);
}
void* QWizard_createWindowContainer2(void* window, void* parent) {
	auto&& __r = QWizard::createWindowContainer((QWindow*)window, (QWidget*)parent);
	return (void*)(__r);
}
void* QWizard_createWindowContainer3(void* window, void* parent, int flags) {
	auto&& __r = QWizard::createWindowContainer((QWindow*)window, (QWidget*)parent, QFlags<Qt::WindowType>(flags));
	return (void*)(__r);
}
void QWizard_updateMicroFocus2(void* self, int query) {
	((QWizard__Go*)self)->updateMicroFocus((Qt::InputMethodQuery)(query));
}
void QWizard_create2(void* self, unsigned long long arg1) {
	((QWizard__Go*)self)->create(arg1);
}
void QWizard_create3(void* self, unsigned long long arg1, int initializeWindow) {
	((QWizard__Go*)self)->create(arg1, initializeWindow);
}
void QWizard_create4(void* self, unsigned long long arg1, int initializeWindow, int destroyOldWindow) {
	((QWizard__Go*)self)->create(arg1, initializeWindow, destroyOldWindow);
}
void QWizard_destroy2(void* self, int destroyWindow) {
	((QWizard__Go*)self)->destroy(destroyWindow);
}
void QWizard_destroy3(void* self, int destroyWindow, int destroySubWindows) {
	((QWizard__Go*)self)->destroy(destroyWindow, destroySubWindows);
}
void* QWizard_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QWizard::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QWizard_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QWizard::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
int QWizard_validateCurrentPage_default(void* self) {
	auto&& __r = ((QWizard*)self)->QWizard::validateCurrentPage();
	return __r;
}
int QWizard_nextId_default(void* self) {
	auto&& __r = ((QWizard*)self)->QWizard::nextId();
	return __r;
}
void* QWizard_metaObject_default(void* self) {
	auto&& __r = ((QWizard*)self)->QWizard::metaObject();
	return (void*)(__r);
}
void QWizard_setVisible_default(void* self, int visible) {
	((QWizard*)self)->QWizard::setVisible(visible);
}
void* QWizard_sizeHint_default(void* self) {
	auto&& __r = ((QWizard*)self)->QWizard::sizeHint();
	return new QSize(__r);
}
int QWizard_event_default(void* self, void* event) {
	auto&& __r = ((QWizard__Go*)self)->cppgoBase_event((QEvent*)event);
	return __r;
}
void QWizard_resizeEvent_default(void* self, void* event) {
	((QWizard__Go*)self)->cppgoBase_resizeEvent((QResizeEvent*)event);
}
void QWizard_paintEvent_default(void* self, void* event) {
	((QWizard__Go*)self)->cppgoBase_paintEvent((QPaintEvent*)event);
}
void QWizard_done_default(void* self, int result) {
	((QWizard__Go*)self)->cppgoBase_done(result);
}
void QWizard_initializePage_default(void* self, int id) {
	((QWizard__Go*)self)->cppgoBase_initializePage(id);
}
void QWizard_cleanupPage_default(void* self, int id) {
	((QWizard__Go*)self)->cppgoBase_cleanupPage(id);
}
void* QWizard_minimumSizeHint_default(void* self) {
	auto&& __r = ((QWizard*)self)->QWizard::minimumSizeHint();
	return new QSize(__r);
}
void QWizard_open_default(void* self) {
	((QWizard*)self)->QWizard::open();
}
int QWizard_exec_default(void* self) {
	auto&& __r = ((QWizard*)self)->QWizard::exec();
	return __r;
}
void QWizard_accept_default(void* self) {
	((QWizard*)self)->QWizard::accept();
}
void QWizard_reject_default(void* self) {
	((QWizard*)self)->QWizard::reject();
}
void QWizard_keyPressEvent_default(void* self, void* arg1) {
	((QWizard__Go*)self)->cppgoBase_keyPressEvent((QKeyEvent*)arg1);
}
void QWizard_closeEvent_default(void* self, void* arg1) {
	((QWizard__Go*)self)->cppgoBase_closeEvent((QCloseEvent*)arg1);
}
void QWizard_showEvent_default(void* self, void* arg1) {
	((QWizard__Go*)self)->cppgoBase_showEvent((QShowEvent*)arg1);
}
void QWizard_contextMenuEvent_default(void* self, void* arg1) {
	((QWizard__Go*)self)->cppgoBase_contextMenuEvent((QContextMenuEvent*)arg1);
}
int QWizard_eventFilter_default(void* self, void* arg1, void* arg2) {
	auto&& __r = ((QWizard__Go*)self)->cppgoBase_eventFilter((QObject*)arg1, (QEvent*)arg2);
	return __r;
}
int QWizard_devType_default(void* self) {
	auto&& __r = ((QWizard*)self)->QWizard::devType();
	return __r;
}
int QWizard_heightForWidth_default(void* self, int arg1) {
	auto&& __r = ((QWizard*)self)->QWizard::heightForWidth(arg1);
	return __r;
}
int QWizard_hasHeightForWidth_default(void* self) {
	auto&& __r = ((QWizard*)self)->QWizard::hasHeightForWidth();
	return __r;
}
void* QWizard_paintEngine_default(void* self) {
	auto&& __r = ((QWizard*)self)->QWizard::paintEngine();
	return (void*)(__r);
}
void QWizard_mousePressEvent_default(void* self, void* event) {
	((QWizard__Go*)self)->cppgoBase_mousePressEvent((QMouseEvent*)event);
}
void QWizard_mouseReleaseEvent_default(void* self, void* event) {
	((QWizard__Go*)self)->cppgoBase_mouseReleaseEvent((QMouseEvent*)event);
}
void QWizard_mouseDoubleClickEvent_default(void* self, void* event) {
	((QWizard__Go*)self)->cppgoBase_mouseDoubleClickEvent((QMouseEvent*)event);
}
void QWizard_mouseMoveEvent_default(void* self, void* event) {
	((QWizard__Go*)self)->cppgoBase_mouseMoveEvent((QMouseEvent*)event);
}
void QWizard_wheelEvent_default(void* self, void* event) {
	((QWizard__Go*)self)->cppgoBase_wheelEvent((QWheelEvent*)event);
}
void QWizard_keyReleaseEvent_default(void* self, void* event) {
	((QWizard__Go*)self)->cppgoBase_keyReleaseEvent((QKeyEvent*)event);
}
void QWizard_focusInEvent_default(void* self, void* event) {
	((QWizard__Go*)self)->cppgoBase_focusInEvent((QFocusEvent*)event);
}
void QWizard_focusOutEvent_default(void* self, void* event) {
	((QWizard__Go*)self)->cppgoBase_focusOutEvent((QFocusEvent*)event);
}
void QWizard_enterEvent_default(void* self, void* event) {
	((QWizard__Go*)self)->cppgoBase_enterEvent((QEnterEvent*)event);
}
void QWizard_leaveEvent_default(void* self, void* event) {
	((QWizard__Go*)self)->cppgoBase_leaveEvent((QEvent*)event);
}
void QWizard_moveEvent_default(void* self, void* event) {
	((QWizard__Go*)self)->cppgoBase_moveEvent((QMoveEvent*)event);
}
void QWizard_tabletEvent_default(void* self, void* event) {
	((QWizard__Go*)self)->cppgoBase_tabletEvent((QTabletEvent*)event);
}
void QWizard_actionEvent_default(void* self, void* event) {
	((QWizard__Go*)self)->cppgoBase_actionEvent((QActionEvent*)event);
}
void QWizard_dragEnterEvent_default(void* self, void* event) {
	((QWizard__Go*)self)->cppgoBase_dragEnterEvent((QDragEnterEvent*)event);
}
void QWizard_dragMoveEvent_default(void* self, void* event) {
	((QWizard__Go*)self)->cppgoBase_dragMoveEvent((QDragMoveEvent*)event);
}
void QWizard_dragLeaveEvent_default(void* self, void* event) {
	((QWizard__Go*)self)->cppgoBase_dragLeaveEvent((QDragLeaveEvent*)event);
}
void QWizard_dropEvent_default(void* self, void* event) {
	((QWizard__Go*)self)->cppgoBase_dropEvent((QDropEvent*)event);
}
void QWizard_hideEvent_default(void* self, void* event) {
	((QWizard__Go*)self)->cppgoBase_hideEvent((QHideEvent*)event);
}
void QWizard_changeEvent_default(void* self, void* arg1) {
	((QWizard__Go*)self)->cppgoBase_changeEvent((QEvent*)arg1);
}
int QWizard_metric_default(void* self, int arg1) {
	auto&& __r = ((QWizard__Go*)self)->cppgoBase_metric(arg1);
	return __r;
}
void QWizard_initPainter_default(void* self, void* painter) {
	((QWizard__Go*)self)->cppgoBase_initPainter((QPainter*)painter);
}
void* QWizard_redirected_default(void* self, void* offset) {
	auto&& __r = ((QWizard__Go*)self)->cppgoBase_redirected((QPoint*)offset);
	return (void*)(__r);
}
void* QWizard_sharedPainter_default(void* self) {
	auto&& __r = ((QWizard__Go*)self)->cppgoBase_sharedPainter();
	return (void*)(__r);
}
void QWizard_inputMethodEvent_default(void* self, void* arg1) {
	((QWizard__Go*)self)->cppgoBase_inputMethodEvent((QInputMethodEvent*)arg1);
}
void* QWizard_inputMethodQuery_default(void* self, int arg1) {
	auto&& __r = ((QWizard*)self)->QWizard::inputMethodQuery((Qt::InputMethodQuery)(arg1));
	return new QVariant(__r);
}
int QWizard_focusNextPrevChild_default(void* self, int next) {
	auto&& __r = ((QWizard__Go*)self)->cppgoBase_focusNextPrevChild(next);
	return __r;
}
void QWizard_timerEvent_default(void* self, void* event) {
	((QWizard__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QWizard_childEvent_default(void* self, void* event) {
	((QWizard__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QWizard_customEvent_default(void* self, void* event) {
	((QWizard__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QWizard_connectNotify_default(void* self, void* signal) {
	((QWizard__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QWizard_disconnectNotify_default(void* self, void* signal) {
	((QWizard__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QWizard_connect_currentIdChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QWizard*)self, &QWizard::currentIdChanged, (QWizard*)self, [__id](int id){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QWizard);
	__conns_QWizard[__id] = __c;
}
void QWizard_connect_helpRequested(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QWizard*)self, &QWizard::helpRequested, (QWizard*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QWizard);
	__conns_QWizard[__id] = __c;
}
void QWizard_connect_customButtonClicked(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QWizard*)self, &QWizard::customButtonClicked, (QWizard*)self, [__id](int which){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, which);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QWizard);
	__conns_QWizard[__id] = __c;
}
void QWizard_connect_pageAdded(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QWizard*)self, &QWizard::pageAdded, (QWizard*)self, [__id](int id){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QWizard);
	__conns_QWizard[__id] = __c;
}
void QWizard_connect_pageRemoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QWizard*)self, &QWizard::pageRemoved, (QWizard*)self, [__id](int id){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QWizard);
	__conns_QWizard[__id] = __c;
}
void QWizard_connect_finished(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QWizard*)self, &QWizard::finished, (QWizard*)self, [__id](int result){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, result);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QWizard);
	__conns_QWizard[__id] = __c;
}
void QWizard_connect_accepted(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QWizard*)self, &QWizard::accepted, (QWizard*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QWizard);
	__conns_QWizard[__id] = __c;
}
void QWizard_connect_rejected(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QWizard*)self, &QWizard::rejected, (QWizard*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QWizard);
	__conns_QWizard[__id] = __c;
}
void QWizard_connect_windowTitleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QWizard*)self, &QWizard::windowTitleChanged, (QWizard*)self, [__id](const QString& title){
		QByteArray title_ba = title.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)title_ba.constData(), title_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QWizard);
	__conns_QWizard[__id] = __c;
}
void QWizard_connect_windowIconChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QWizard*)self, &QWizard::windowIconChanged, (QWizard*)self, [__id](const QIcon& icon){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&icon);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QWizard);
	__conns_QWizard[__id] = __c;
}
void QWizard_connect_windowIconTextChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QWizard*)self, &QWizard::windowIconTextChanged, (QWizard*)self, [__id](const QString& iconText){
		QByteArray iconText_ba = iconText.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)iconText_ba.constData(), iconText_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QWizard);
	__conns_QWizard[__id] = __c;
}
void QWizard_connect_customContextMenuRequested(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QWizard*)self, &QWizard::customContextMenuRequested, (QWizard*)self, [__id](const QPoint& pos){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&pos);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QWizard);
	__conns_QWizard[__id] = __c;
}
void QWizard_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QWizard*)self, &QWizard::destroyed, (QWizard*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QWizard);
	__conns_QWizard[__id] = __c;
}
void QWizard_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QWizard*)self, &QWizard::destroyed, (QWizard*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QWizard);
	__conns_QWizard[__id] = __c;
}
void QWizard_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QWizard*)self, &QWizard::objectNameChanged, (QWizard*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QWizard);
	__conns_QWizard[__id] = __c;
}
void QWizard__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QWizard);
		auto __it = __conns_QWizard.find(id);
		if (__it == __conns_QWizard.end()) return;
		__c = __it->second;
		__conns_QWizard.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QWizard__Go_new(long long id) { return new QWizard__Go(id); }
void QWizard_delete(void* self) { delete (QWizard*)self; }

}

